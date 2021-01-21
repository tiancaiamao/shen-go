package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp79421 := MakeNative(func(__e *ControlFlow) {
		V867 := __e.Get(1)
		_ = V867
		tmp79422 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5predicate_d_6 := __e.Get(1)
			_ = Parse__shen_4_5predicate_d_6
			tmp79451 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp79452 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79453 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp79454 := Call(__e, tmp79453)

			tmp79455 := Call(__e, tmp79452, tmp79454, Parse__shen_4_5predicate_d_6)

			tmp79456 := Call(__e, tmp79451, tmp79455)

			if True == tmp79456 {
				tmp79425 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5clauses_d_6 := __e.Get(1)
					_ = Parse__shen_4_5clauses_d_6
					tmp79443 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp79444 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79445 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp79446 := Call(__e, tmp79445)

					tmp79447 := Call(__e, tmp79444, tmp79446, Parse__shen_4_5clauses_d_6)

					tmp79448 := Call(__e, tmp79443, tmp79447)

					if True == tmp79448 {
						tmp79428 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp79429 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp79430 := Call(__e, tmp79429, Parse__shen_4_5clauses_d_6)

						tmp79431 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp79432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1_6shen)

						tmp79433 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						tmp79434 := MakeNative(func(__e *ControlFlow) {
							Parse__X := __e.Get(1)
							_ = Parse__X
							tmp79435 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1predicate)

							tmp79436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp79437 := Call(__e, tmp79436, Parse__shen_4_5predicate_d_6)

							__e.TailApply(tmp79435, tmp79437, Parse__X)
							return

						}, 1)

						tmp79438 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp79439 := Call(__e, tmp79438, Parse__shen_4_5clauses_d_6)

						tmp79440 := Call(__e, tmp79433, tmp79434, tmp79439)

						tmp79441 := Call(__e, tmp79432, tmp79440)

						tmp79442 := Call(__e, tmp79431, tmp79441)

						__e.TailApply(tmp79428, tmp79430, tmp79442)
						return

					} else {
						tmp79427 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp79427)
						return

					}

				}, 1)

				tmp79449 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5clauses_d_6)

				tmp79450 := Call(__e, tmp79449, Parse__shen_4_5predicate_d_6)

				__e.TailApply(tmp79425, tmp79450)
				return

			} else {
				tmp79424 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp79424)
				return

			}

		}, 1)

		tmp79457 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5predicate_d_6)

		tmp79458 := Call(__e, tmp79457, V867)

		__e.TailApply(tmp79422, tmp79458)
		return

	}, 1)

	tmp79459 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5defprolog_6, tmp79421)

	_ = tmp79459

	tmp79460 := MakeNative(func(__e *ControlFlow) {
		V876 := __e.Get(1)
		_ = V876
		V877 := __e.Get(2)
		_ = V877
		tmp79494 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79495 := Call(__e, tmp79494, V877)

		var ifres79480 Obj

		if True == tmp79495 {
			tmp79490 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79491 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79492 := Call(__e, tmp79491, V877)

			tmp79493 := Call(__e, tmp79490, tmp79492)

			var ifres79482 Obj

			if True == tmp79493 {
				tmp79484 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp79485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79486 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79487 := Call(__e, tmp79486, V877)

				tmp79488 := Call(__e, tmp79485, tmp79487)

				tmp79489 := Call(__e, tmp79484, Nil, tmp79488)

				var ifres79483 Obj

				if True == tmp79489 {
					ifres79483 = True

				} else {
					ifres79483 = False

				}

				ifres79482 = ifres79483

			} else {
				ifres79482 = False

			}

			var ifres79481 Obj

			if True == ifres79482 {
				ifres79481 = True

			} else {
				ifres79481 = False

			}

			ifres79480 = ifres79481

		} else {
			ifres79480 = False

		}

		if True == ifres79480 {
			tmp79467 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp79468 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp79469 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp79470 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp79471 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp79472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4next_150)

			tmp79473 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79474 := Call(__e, tmp79473, V877)

			tmp79475 := Call(__e, tmp79472, MakeNumber(50), tmp79474)

			tmp79476 := Call(__e, tmp79471, tmp79475, MakeString("\n"), symshen_4a)

			tmp79477 := Call(__e, tmp79470, MakeString(" here:\n\n "), tmp79476)

			tmp79478 := Call(__e, tmp79469, V876, tmp79477, symshen_4a)

			tmp79479 := Call(__e, tmp79468, MakeString("prolog syntax error in "), tmp79478)

			__e.TailApply(tmp79467, tmp79479)
			return

		} else {
			tmp79462 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp79463 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp79464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp79465 := Call(__e, tmp79464, V876, MakeString("\n"), symshen_4a)

			tmp79466 := Call(__e, tmp79463, MakeString("prolog syntax error in "), tmp79465)

			__e.TailApply(tmp79462, tmp79466)
			return

		}

	}, 2)

	tmp79496 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1error, tmp79460)

	_ = tmp79496

	tmp79497 := MakeNative(func(__e *ControlFlow) {
		V884 := __e.Get(1)
		_ = V884
		V885 := __e.Get(2)
		_ = V885
		tmp79517 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp79518 := Call(__e, tmp79517, Nil, V885)

		if True == tmp79518 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp79515 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79516 := Call(__e, tmp79515, MakeNumber(0), V884)

			if True == tmp79516 {
				__e.Return(MakeString(""))
				return
			} else {
				tmp79513 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79514 := Call(__e, tmp79513, V885)

				if True == tmp79514 {
					tmp79502 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp79503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decons_1string)

					tmp79504 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79505 := Call(__e, tmp79504, V885)

					tmp79506 := Call(__e, tmp79503, tmp79505)

					tmp79507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4next_150)

					tmp79508 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					tmp79509 := Call(__e, tmp79508, V884, MakeNumber(1))

					tmp79510 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79511 := Call(__e, tmp79510, V885)

					tmp79512 := Call(__e, tmp79507, tmp79509, tmp79511)

					__e.TailApply(tmp79502, tmp79506, tmp79512)
					return

				} else {
					tmp79501 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp79501, symshen_4next_150)
					return

				}

			}

		}

	}, 2)

	tmp79519 := Call(__e, PrimNS1Value(symns2_1set), symshen_4next_150, tmp79497)

	_ = tmp79519

	tmp79520 := MakeNative(func(__e *ControlFlow) {
		V887 := __e.Get(1)
		_ = V887
		tmp79556 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79557 := Call(__e, tmp79556, V887)

		var ifres79526 Obj

		if True == tmp79557 {
			tmp79552 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79553 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79554 := Call(__e, tmp79553, V887)

			tmp79555 := Call(__e, tmp79552, symcons, tmp79554)

			var ifres79528 Obj

			if True == tmp79555 {
				tmp79548 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79549 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79550 := Call(__e, tmp79549, V887)

				tmp79551 := Call(__e, tmp79548, tmp79550)

				var ifres79530 Obj

				if True == tmp79551 {
					tmp79542 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp79543 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79544 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79545 := Call(__e, tmp79544, V887)

					tmp79546 := Call(__e, tmp79543, tmp79545)

					tmp79547 := Call(__e, tmp79542, tmp79546)

					var ifres79532 Obj

					if True == tmp79547 {
						tmp79534 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp79535 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79536 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79537 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79538 := Call(__e, tmp79537, V887)

						tmp79539 := Call(__e, tmp79536, tmp79538)

						tmp79540 := Call(__e, tmp79535, tmp79539)

						tmp79541 := Call(__e, tmp79534, Nil, tmp79540)

						var ifres79533 Obj

						if True == tmp79541 {
							ifres79533 = True

						} else {
							ifres79533 = False

						}

						ifres79532 = ifres79533

					} else {
						ifres79532 = False

					}

					var ifres79531 Obj

					if True == ifres79532 {
						ifres79531 = True

					} else {
						ifres79531 = False

					}

					ifres79530 = ifres79531

				} else {
					ifres79530 = False

				}

				var ifres79529 Obj

				if True == ifres79530 {
					ifres79529 = True

				} else {
					ifres79529 = False

				}

				ifres79528 = ifres79529

			} else {
				ifres79528 = False

			}

			var ifres79527 Obj

			if True == ifres79528 {
				ifres79527 = True

			} else {
				ifres79527 = False

			}

			ifres79526 = ifres79527

		} else {
			ifres79526 = False

		}

		if True == ifres79526 {
			tmp79523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp79524 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1cons)

			tmp79525 := Call(__e, tmp79524, V887)

			__e.TailApply(tmp79523, tmp79525, MakeString(" "), symshen_4s)
			return

		} else {
			tmp79522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			__e.TailApply(tmp79522, V887, MakeString(" "), symshen_4r)
			return

		}

	}, 1)

	tmp79558 := Call(__e, PrimNS1Value(symns2_1set), symshen_4decons_1string, tmp79520)

	_ = tmp79558

	tmp79559 := MakeNative(func(__e *ControlFlow) {
		V890 := __e.Get(1)
		_ = V890
		V891 := __e.Get(2)
		_ = V891
		tmp79585 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79586 := Call(__e, tmp79585, V891)

		var ifres79571 Obj

		if True == tmp79586 {
			tmp79581 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79582 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79583 := Call(__e, tmp79582, V891)

			tmp79584 := Call(__e, tmp79581, tmp79583)

			var ifres79573 Obj

			if True == tmp79584 {
				tmp79575 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp79576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79577 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79578 := Call(__e, tmp79577, V891)

				tmp79579 := Call(__e, tmp79576, tmp79578)

				tmp79580 := Call(__e, tmp79575, Nil, tmp79579)

				var ifres79574 Obj

				if True == tmp79580 {
					ifres79574 = True

				} else {
					ifres79574 = False

				}

				ifres79573 = ifres79574

			} else {
				ifres79573 = False

			}

			var ifres79572 Obj

			if True == ifres79573 {
				ifres79572 = True

			} else {
				ifres79572 = False

			}

			ifres79571 = ifres79572

		} else {
			ifres79571 = False

		}

		if True == ifres79571 {
			tmp79562 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79563 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79564 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79565 := Call(__e, tmp79564, V891)

			tmp79566 := Call(__e, tmp79563, V890, tmp79565)

			tmp79567 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79569 := Call(__e, tmp79568, V891)

			tmp79570 := Call(__e, tmp79567, sym_h_1, tmp79569)

			__e.TailApply(tmp79562, tmp79566, tmp79570)
			return

		} else {
			tmp79561 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp79561, symshen_4insert_1predicate)
			return

		}

	}, 2)

	tmp79587 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1predicate, tmp79559)

	_ = tmp79587

	tmp79588 := MakeNative(func(__e *ControlFlow) {
		V893 := __e.Get(1)
		_ = V893
		tmp79603 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79604 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp79605 := Call(__e, tmp79604, V893)

		tmp79606 := Call(__e, tmp79603, tmp79605)

		if True == tmp79606 {
			tmp79591 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp79592 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp79593 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp79595 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp79596 := Call(__e, tmp79595, V893)

				tmp79597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp79598 := Call(__e, tmp79597, V893)

				tmp79599 := Call(__e, tmp79594, tmp79596, tmp79598)

				tmp79600 := Call(__e, tmp79593, tmp79599)

				__e.TailApply(tmp79592, tmp79600, Parse__X)
				return

			}, 1)

			tmp79601 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp79602 := Call(__e, tmp79601, V893)

			__e.TailApply(tmp79591, tmp79602)
			return

		} else {
			tmp79590 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp79590)
			return

		}

	}, 1)

	tmp79607 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5predicate_d_6, tmp79588)

	_ = tmp79607

	tmp79608 := MakeNative(func(__e *ControlFlow) {
		V895 := __e.Get(1)
		_ = V895
		tmp79609 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp79625 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79626 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp79627 := Call(__e, tmp79626)

			tmp79628 := Call(__e, tmp79625, YaccParse, tmp79627)

			if True == tmp79628 {
				tmp79611 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp79617 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp79618 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79619 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp79620 := Call(__e, tmp79619)

					tmp79621 := Call(__e, tmp79618, tmp79620, Parse___5e_6)

					tmp79622 := Call(__e, tmp79617, tmp79621)

					if True == tmp79622 {
						tmp79614 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp79615 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp79616 := Call(__e, tmp79615, Parse___5e_6)

						__e.TailApply(tmp79614, tmp79616, Nil)
						return

					} else {
						tmp79613 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp79613)
						return

					}

				}, 1)

				tmp79623 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp79624 := Call(__e, tmp79623, V895)

				__e.TailApply(tmp79611, tmp79624)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp79629 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5clause_d_6 := __e.Get(1)
			_ = Parse__shen_4_5clause_d_6
			tmp79652 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp79653 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79654 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp79655 := Call(__e, tmp79654)

			tmp79656 := Call(__e, tmp79653, tmp79655, Parse__shen_4_5clause_d_6)

			tmp79657 := Call(__e, tmp79652, tmp79656)

			if True == tmp79657 {
				tmp79632 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5clauses_d_6 := __e.Get(1)
					_ = Parse__shen_4_5clauses_d_6
					tmp79644 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp79645 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79646 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp79647 := Call(__e, tmp79646)

					tmp79648 := Call(__e, tmp79645, tmp79647, Parse__shen_4_5clauses_d_6)

					tmp79649 := Call(__e, tmp79644, tmp79648)

					if True == tmp79649 {
						tmp79635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp79636 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp79637 := Call(__e, tmp79636, Parse__shen_4_5clauses_d_6)

						tmp79638 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp79639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp79640 := Call(__e, tmp79639, Parse__shen_4_5clause_d_6)

						tmp79641 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp79642 := Call(__e, tmp79641, Parse__shen_4_5clauses_d_6)

						tmp79643 := Call(__e, tmp79638, tmp79640, tmp79642)

						__e.TailApply(tmp79635, tmp79637, tmp79643)
						return

					} else {
						tmp79634 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp79634)
						return

					}

				}, 1)

				tmp79650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5clauses_d_6)

				tmp79651 := Call(__e, tmp79650, Parse__shen_4_5clause_d_6)

				__e.TailApply(tmp79632, tmp79651)
				return

			} else {
				tmp79631 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp79631)
				return

			}

		}, 1)

		tmp79658 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5clause_d_6)

		tmp79659 := Call(__e, tmp79658, V895)

		tmp79660 := Call(__e, tmp79629, tmp79659)

		__e.TailApply(tmp79609, tmp79660)
		return

	}, 1)

	tmp79661 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5clauses_d_6, tmp79608)

	_ = tmp79661

	tmp79662 := MakeNative(func(__e *ControlFlow) {
		V898 := __e.Get(1)
		_ = V898
		tmp79663 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5head_d_6 := __e.Get(1)
			_ = Parse__shen_4_5head_d_6
			tmp79718 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp79719 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79720 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp79721 := Call(__e, tmp79720)

			tmp79722 := Call(__e, tmp79719, tmp79721, Parse__shen_4_5head_d_6)

			tmp79723 := Call(__e, tmp79718, tmp79722)

			if True == tmp79723 {
				tmp79714 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79715 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79716 := Call(__e, tmp79715, Parse__shen_4_5head_d_6)

				tmp79717 := Call(__e, tmp79714, tmp79716)

				var ifres79708 Obj

				if True == tmp79717 {
					tmp79710 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79711 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp79712 := Call(__e, tmp79711, Parse__shen_4_5head_d_6)

					tmp79713 := Call(__e, tmp79710, sym_5_1_1, tmp79712)

					var ifres79709 Obj

					if True == tmp79713 {
						ifres79709 = True

					} else {
						ifres79709 = False

					}

					ifres79708 = ifres79709

				} else {
					ifres79708 = False

				}

				if True == ifres79708 {
					tmp79668 := MakeNative(func(__e *ControlFlow) {
						NewStream896 := __e.Get(1)
						_ = NewStream896
						tmp79669 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5body_d_6 := __e.Get(1)
							_ = Parse__shen_4_5body_d_6
							tmp79694 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp79695 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp79696 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp79697 := Call(__e, tmp79696)

							tmp79698 := Call(__e, tmp79695, tmp79697, Parse__shen_4_5body_d_6)

							tmp79699 := Call(__e, tmp79694, tmp79698)

							if True == tmp79699 {
								tmp79672 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5end_d_6 := __e.Get(1)
									_ = Parse__shen_4_5end_d_6
									tmp79686 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp79687 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp79688 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp79689 := Call(__e, tmp79688)

									tmp79690 := Call(__e, tmp79687, tmp79689, Parse__shen_4_5end_d_6)

									tmp79691 := Call(__e, tmp79686, tmp79690)

									if True == tmp79691 {
										tmp79675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp79676 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp79677 := Call(__e, tmp79676, Parse__shen_4_5end_d_6)

										tmp79678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp79679 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp79680 := Call(__e, tmp79679, Parse__shen_4_5head_d_6)

										tmp79681 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp79682 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp79683 := Call(__e, tmp79682, Parse__shen_4_5body_d_6)

										tmp79684 := Call(__e, tmp79681, tmp79683, Nil)

										tmp79685 := Call(__e, tmp79678, tmp79680, tmp79684)

										__e.TailApply(tmp79675, tmp79677, tmp79685)
										return

									} else {
										tmp79674 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp79674)
										return

									}

								}, 1)

								tmp79692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5end_d_6)

								tmp79693 := Call(__e, tmp79692, Parse__shen_4_5body_d_6)

								__e.TailApply(tmp79672, tmp79693)
								return

							} else {
								tmp79671 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp79671)
								return

							}

						}, 1)

						tmp79700 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5body_d_6)

						tmp79701 := Call(__e, tmp79700, NewStream896)

						__e.TailApply(tmp79669, tmp79701)
						return

					}, 1)

					tmp79702 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp79703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp79704 := Call(__e, tmp79703, Parse__shen_4_5head_d_6)

					tmp79705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp79706 := Call(__e, tmp79705, Parse__shen_4_5head_d_6)

					tmp79707 := Call(__e, tmp79702, tmp79704, tmp79706)

					__e.TailApply(tmp79668, tmp79707)
					return

				} else {
					tmp79667 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp79667)
					return

				}

			} else {
				tmp79665 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp79665)
				return

			}

		}, 1)

		tmp79724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5head_d_6)

		tmp79725 := Call(__e, tmp79724, V898)

		__e.TailApply(tmp79663, tmp79725)
		return

	}, 1)

	tmp79726 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5clause_d_6, tmp79662)

	_ = tmp79726

	tmp79727 := MakeNative(func(__e *ControlFlow) {
		V900 := __e.Get(1)
		_ = V900
		tmp79728 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp79744 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79745 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp79746 := Call(__e, tmp79745)

			tmp79747 := Call(__e, tmp79744, YaccParse, tmp79746)

			if True == tmp79747 {
				tmp79730 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp79736 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp79737 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79738 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp79739 := Call(__e, tmp79738)

					tmp79740 := Call(__e, tmp79737, tmp79739, Parse___5e_6)

					tmp79741 := Call(__e, tmp79736, tmp79740)

					if True == tmp79741 {
						tmp79733 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp79734 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp79735 := Call(__e, tmp79734, Parse___5e_6)

						__e.TailApply(tmp79733, tmp79735, Nil)
						return

					} else {
						tmp79732 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp79732)
						return

					}

				}, 1)

				tmp79742 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp79743 := Call(__e, tmp79742, V900)

				__e.TailApply(tmp79730, tmp79743)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp79748 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5term_d_6 := __e.Get(1)
			_ = Parse__shen_4_5term_d_6
			tmp79771 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp79772 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79773 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp79774 := Call(__e, tmp79773)

			tmp79775 := Call(__e, tmp79772, tmp79774, Parse__shen_4_5term_d_6)

			tmp79776 := Call(__e, tmp79771, tmp79775)

			if True == tmp79776 {
				tmp79751 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5head_d_6 := __e.Get(1)
					_ = Parse__shen_4_5head_d_6
					tmp79763 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp79764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79765 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp79766 := Call(__e, tmp79765)

					tmp79767 := Call(__e, tmp79764, tmp79766, Parse__shen_4_5head_d_6)

					tmp79768 := Call(__e, tmp79763, tmp79767)

					if True == tmp79768 {
						tmp79754 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp79755 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp79756 := Call(__e, tmp79755, Parse__shen_4_5head_d_6)

						tmp79757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp79758 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp79759 := Call(__e, tmp79758, Parse__shen_4_5term_d_6)

						tmp79760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp79761 := Call(__e, tmp79760, Parse__shen_4_5head_d_6)

						tmp79762 := Call(__e, tmp79757, tmp79759, tmp79761)

						__e.TailApply(tmp79754, tmp79756, tmp79762)
						return

					} else {
						tmp79753 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp79753)
						return

					}

				}, 1)

				tmp79769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5head_d_6)

				tmp79770 := Call(__e, tmp79769, Parse__shen_4_5term_d_6)

				__e.TailApply(tmp79751, tmp79770)
				return

			} else {
				tmp79750 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp79750)
				return

			}

		}, 1)

		tmp79777 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5term_d_6)

		tmp79778 := Call(__e, tmp79777, V900)

		tmp79779 := Call(__e, tmp79748, tmp79778)

		__e.TailApply(tmp79728, tmp79779)
		return

	}, 1)

	tmp79780 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5head_d_6, tmp79727)

	_ = tmp79780

	tmp79781 := MakeNative(func(__e *ControlFlow) {
		V902 := __e.Get(1)
		_ = V902
		tmp79808 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79809 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp79810 := Call(__e, tmp79809, V902)

		tmp79811 := Call(__e, tmp79808, tmp79810)

		if True == tmp79811 {
			tmp79784 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp79802 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp79803 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp79804 := Call(__e, tmp79803, sym_5_1_1, Parse__X)

				tmp79805 := Call(__e, tmp79802, tmp79804)

				var ifres79798 Obj

				if True == tmp79805 {
					tmp79800 := Call(__e, PrimNS1Value(symns2_1value), symshen_4legitimate_1term_2)

					tmp79801 := Call(__e, tmp79800, Parse__X)

					var ifres79799 Obj

					if True == tmp79801 {
						ifres79799 = True

					} else {
						ifres79799 = False

					}

					ifres79798 = ifres79799

				} else {
					ifres79798 = False

				}

				if True == ifres79798 {
					tmp79787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp79788 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp79790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp79791 := Call(__e, tmp79790, V902)

					tmp79792 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp79793 := Call(__e, tmp79792, V902)

					tmp79794 := Call(__e, tmp79789, tmp79791, tmp79793)

					tmp79795 := Call(__e, tmp79788, tmp79794)

					tmp79796 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1cons)

					tmp79797 := Call(__e, tmp79796, Parse__X)

					__e.TailApply(tmp79787, tmp79795, tmp79797)
					return

				} else {
					tmp79786 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp79786)
					return

				}

			}, 1)

			tmp79806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp79807 := Call(__e, tmp79806, V902)

			__e.TailApply(tmp79784, tmp79807)
			return

		} else {
			tmp79783 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp79783)
			return

		}

	}, 1)

	tmp79812 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5term_d_6, tmp79781)

	_ = tmp79812

	tmp79813 := MakeNative(func(__e *ControlFlow) {
		V908 := __e.Get(1)
		_ = V908
		tmp79960 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79961 := Call(__e, tmp79960, V908)

		var ifres79930 Obj

		if True == tmp79961 {
			tmp79956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79957 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79958 := Call(__e, tmp79957, V908)

			tmp79959 := Call(__e, tmp79956, symcons, tmp79958)

			var ifres79932 Obj

			if True == tmp79959 {
				tmp79952 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79953 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79954 := Call(__e, tmp79953, V908)

				tmp79955 := Call(__e, tmp79952, tmp79954)

				var ifres79934 Obj

				if True == tmp79955 {
					tmp79946 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp79947 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79948 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79949 := Call(__e, tmp79948, V908)

					tmp79950 := Call(__e, tmp79947, tmp79949)

					tmp79951 := Call(__e, tmp79946, tmp79950)

					var ifres79936 Obj

					if True == tmp79951 {
						tmp79938 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp79939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79941 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79942 := Call(__e, tmp79941, V908)

						tmp79943 := Call(__e, tmp79940, tmp79942)

						tmp79944 := Call(__e, tmp79939, tmp79943)

						tmp79945 := Call(__e, tmp79938, Nil, tmp79944)

						var ifres79937 Obj

						if True == tmp79945 {
							ifres79937 = True

						} else {
							ifres79937 = False

						}

						ifres79936 = ifres79937

					} else {
						ifres79936 = False

					}

					var ifres79935 Obj

					if True == ifres79936 {
						ifres79935 = True

					} else {
						ifres79935 = False

					}

					ifres79934 = ifres79935

				} else {
					ifres79934 = False

				}

				var ifres79933 Obj

				if True == ifres79934 {
					ifres79933 = True

				} else {
					ifres79933 = False

				}

				ifres79932 = ifres79933

			} else {
				ifres79932 = False

			}

			var ifres79931 Obj

			if True == ifres79932 {
				ifres79931 = True

			} else {
				ifres79931 = False

			}

			ifres79930 = ifres79931

		} else {
			ifres79930 = False

		}

		if True == ifres79930 {
			tmp79924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4legitimate_1term_2)

			tmp79925 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79926 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79927 := Call(__e, tmp79926, V908)

			tmp79928 := Call(__e, tmp79925, tmp79927)

			tmp79929 := Call(__e, tmp79924, tmp79928)

			if True == tmp79929 {
				tmp79916 := Call(__e, PrimNS1Value(symns2_1value), symshen_4legitimate_1term_2)

				tmp79917 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79918 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79919 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79920 := Call(__e, tmp79919, V908)

				tmp79921 := Call(__e, tmp79918, tmp79920)

				tmp79922 := Call(__e, tmp79917, tmp79921)

				tmp79923 := Call(__e, tmp79916, tmp79922)

				if True == tmp79923 {
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
			tmp79912 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79913 := Call(__e, tmp79912, V908)

			var ifres79872 Obj

			if True == tmp79913 {
				tmp79908 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp79909 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79910 := Call(__e, tmp79909, V908)

				tmp79911 := Call(__e, tmp79908, symmode, tmp79910)

				var ifres79874 Obj

				if True == tmp79911 {
					tmp79904 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp79905 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79906 := Call(__e, tmp79905, V908)

					tmp79907 := Call(__e, tmp79904, tmp79906)

					var ifres79876 Obj

					if True == tmp79907 {
						tmp79898 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp79899 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79900 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79901 := Call(__e, tmp79900, V908)

						tmp79902 := Call(__e, tmp79899, tmp79901)

						tmp79903 := Call(__e, tmp79898, tmp79902)

						var ifres79878 Obj

						if True == tmp79903 {
							tmp79890 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp79891 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp79892 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79893 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79894 := Call(__e, tmp79893, V908)

							tmp79895 := Call(__e, tmp79892, tmp79894)

							tmp79896 := Call(__e, tmp79891, tmp79895)

							tmp79897 := Call(__e, tmp79890, sym_7, tmp79896)

							var ifres79880 Obj

							if True == tmp79897 {
								tmp79882 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp79883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79884 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79885 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79886 := Call(__e, tmp79885, V908)

								tmp79887 := Call(__e, tmp79884, tmp79886)

								tmp79888 := Call(__e, tmp79883, tmp79887)

								tmp79889 := Call(__e, tmp79882, Nil, tmp79888)

								var ifres79881 Obj

								if True == tmp79889 {
									ifres79881 = True

								} else {
									ifres79881 = False

								}

								ifres79880 = ifres79881

							} else {
								ifres79880 = False

							}

							var ifres79879 Obj

							if True == ifres79880 {
								ifres79879 = True

							} else {
								ifres79879 = False

							}

							ifres79878 = ifres79879

						} else {
							ifres79878 = False

						}

						var ifres79877 Obj

						if True == ifres79878 {
							ifres79877 = True

						} else {
							ifres79877 = False

						}

						ifres79876 = ifres79877

					} else {
						ifres79876 = False

					}

					var ifres79875 Obj

					if True == ifres79876 {
						ifres79875 = True

					} else {
						ifres79875 = False

					}

					ifres79874 = ifres79875

				} else {
					ifres79874 = False

				}

				var ifres79873 Obj

				if True == ifres79874 {
					ifres79873 = True

				} else {
					ifres79873 = False

				}

				ifres79872 = ifres79873

			} else {
				ifres79872 = False

			}

			if True == ifres79872 {
				tmp79867 := Call(__e, PrimNS1Value(symns2_1value), symshen_4legitimate_1term_2)

				tmp79868 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79869 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79870 := Call(__e, tmp79869, V908)

				tmp79871 := Call(__e, tmp79868, tmp79870)

				__e.TailApply(tmp79867, tmp79871)
				return

			} else {
				tmp79865 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79866 := Call(__e, tmp79865, V908)

				var ifres79825 Obj

				if True == tmp79866 {
					tmp79861 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79862 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79863 := Call(__e, tmp79862, V908)

					tmp79864 := Call(__e, tmp79861, symmode, tmp79863)

					var ifres79827 Obj

					if True == tmp79864 {
						tmp79857 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp79858 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79859 := Call(__e, tmp79858, V908)

						tmp79860 := Call(__e, tmp79857, tmp79859)

						var ifres79829 Obj

						if True == tmp79860 {
							tmp79851 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp79852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79853 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79854 := Call(__e, tmp79853, V908)

							tmp79855 := Call(__e, tmp79852, tmp79854)

							tmp79856 := Call(__e, tmp79851, tmp79855)

							var ifres79831 Obj

							if True == tmp79856 {
								tmp79843 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp79844 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp79845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79846 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79847 := Call(__e, tmp79846, V908)

								tmp79848 := Call(__e, tmp79845, tmp79847)

								tmp79849 := Call(__e, tmp79844, tmp79848)

								tmp79850 := Call(__e, tmp79843, sym_1, tmp79849)

								var ifres79833 Obj

								if True == tmp79850 {
									tmp79835 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp79836 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp79837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp79838 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp79839 := Call(__e, tmp79838, V908)

									tmp79840 := Call(__e, tmp79837, tmp79839)

									tmp79841 := Call(__e, tmp79836, tmp79840)

									tmp79842 := Call(__e, tmp79835, Nil, tmp79841)

									var ifres79834 Obj

									if True == tmp79842 {
										ifres79834 = True

									} else {
										ifres79834 = False

									}

									ifres79833 = ifres79834

								} else {
									ifres79833 = False

								}

								var ifres79832 Obj

								if True == ifres79833 {
									ifres79832 = True

								} else {
									ifres79832 = False

								}

								ifres79831 = ifres79832

							} else {
								ifres79831 = False

							}

							var ifres79830 Obj

							if True == ifres79831 {
								ifres79830 = True

							} else {
								ifres79830 = False

							}

							ifres79829 = ifres79830

						} else {
							ifres79829 = False

						}

						var ifres79828 Obj

						if True == ifres79829 {
							ifres79828 = True

						} else {
							ifres79828 = False

						}

						ifres79827 = ifres79828

					} else {
						ifres79827 = False

					}

					var ifres79826 Obj

					if True == ifres79827 {
						ifres79826 = True

					} else {
						ifres79826 = False

					}

					ifres79825 = ifres79826

				} else {
					ifres79825 = False

				}

				if True == ifres79825 {
					tmp79820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4legitimate_1term_2)

					tmp79821 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79822 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79823 := Call(__e, tmp79822, V908)

					tmp79824 := Call(__e, tmp79821, tmp79823)

					__e.TailApply(tmp79820, tmp79824)
					return

				} else {
					tmp79818 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp79819 := Call(__e, tmp79818, V908)

					if True == tmp79819 {
						__e.Return(False)
						return
					} else {
						__e.Return(True)
						return
					}

				}

			}

		}

	}, 1)

	tmp79962 := Call(__e, PrimNS1Value(symns2_1set), symshen_4legitimate_1term_2, tmp79813)

	_ = tmp79962

	tmp79963 := MakeNative(func(__e *ControlFlow) {
		V910 := __e.Get(1)
		_ = V910
		tmp80056 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80057 := Call(__e, tmp80056, V910)

		var ifres80026 Obj

		if True == tmp80057 {
			tmp80052 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80053 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80054 := Call(__e, tmp80053, V910)

			tmp80055 := Call(__e, tmp80052, symcons, tmp80054)

			var ifres80028 Obj

			if True == tmp80055 {
				tmp80048 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80049 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80050 := Call(__e, tmp80049, V910)

				tmp80051 := Call(__e, tmp80048, tmp80050)

				var ifres80030 Obj

				if True == tmp80051 {
					tmp80042 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80043 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80044 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80045 := Call(__e, tmp80044, V910)

					tmp80046 := Call(__e, tmp80043, tmp80045)

					tmp80047 := Call(__e, tmp80042, tmp80046)

					var ifres80032 Obj

					if True == tmp80047 {
						tmp80034 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80035 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80036 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80037 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80038 := Call(__e, tmp80037, V910)

						tmp80039 := Call(__e, tmp80036, tmp80038)

						tmp80040 := Call(__e, tmp80035, tmp80039)

						tmp80041 := Call(__e, tmp80034, Nil, tmp80040)

						var ifres80033 Obj

						if True == tmp80041 {
							ifres80033 = True

						} else {
							ifres80033 = False

						}

						ifres80032 = ifres80033

					} else {
						ifres80032 = False

					}

					var ifres80031 Obj

					if True == ifres80032 {
						ifres80031 = True

					} else {
						ifres80031 = False

					}

					ifres80030 = ifres80031

				} else {
					ifres80030 = False

				}

				var ifres80029 Obj

				if True == ifres80030 {
					ifres80029 = True

				} else {
					ifres80029 = False

				}

				ifres80028 = ifres80029

			} else {
				ifres80028 = False

			}

			var ifres80027 Obj

			if True == ifres80028 {
				ifres80027 = True

			} else {
				ifres80027 = False

			}

			ifres80026 = ifres80027

		} else {
			ifres80026 = False

		}

		if True == ifres80026 {
			tmp80011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp80012 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1cons)

			tmp80013 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80014 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80015 := Call(__e, tmp80014, V910)

			tmp80016 := Call(__e, tmp80013, tmp80015)

			tmp80017 := Call(__e, tmp80012, tmp80016)

			tmp80018 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1cons)

			tmp80019 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80020 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80021 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80022 := Call(__e, tmp80021, V910)

			tmp80023 := Call(__e, tmp80020, tmp80022)

			tmp80024 := Call(__e, tmp80019, tmp80023)

			tmp80025 := Call(__e, tmp80018, tmp80024)

			__e.TailApply(tmp80011, tmp80017, tmp80025)
			return

		} else {
			tmp80009 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp80010 := Call(__e, tmp80009, V910)

			var ifres79979 Obj

			if True == tmp80010 {
				tmp80005 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp80006 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80007 := Call(__e, tmp80006, V910)

				tmp80008 := Call(__e, tmp80005, symmode, tmp80007)

				var ifres79981 Obj

				if True == tmp80008 {
					tmp80001 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80002 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80003 := Call(__e, tmp80002, V910)

					tmp80004 := Call(__e, tmp80001, tmp80003)

					var ifres79983 Obj

					if True == tmp80004 {
						tmp79995 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp79996 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79997 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79998 := Call(__e, tmp79997, V910)

						tmp79999 := Call(__e, tmp79996, tmp79998)

						tmp80000 := Call(__e, tmp79995, tmp79999)

						var ifres79985 Obj

						if True == tmp80000 {
							tmp79987 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp79988 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79989 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79990 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79991 := Call(__e, tmp79990, V910)

							tmp79992 := Call(__e, tmp79989, tmp79991)

							tmp79993 := Call(__e, tmp79988, tmp79992)

							tmp79994 := Call(__e, tmp79987, Nil, tmp79993)

							var ifres79986 Obj

							if True == tmp79994 {
								ifres79986 = True

							} else {
								ifres79986 = False

							}

							ifres79985 = ifres79986

						} else {
							ifres79985 = False

						}

						var ifres79984 Obj

						if True == ifres79985 {
							ifres79984 = True

						} else {
							ifres79984 = False

						}

						ifres79983 = ifres79984

					} else {
						ifres79983 = False

					}

					var ifres79982 Obj

					if True == ifres79983 {
						ifres79982 = True

					} else {
						ifres79982 = False

					}

					ifres79981 = ifres79982

				} else {
					ifres79981 = False

				}

				var ifres79980 Obj

				if True == ifres79981 {
					ifres79980 = True

				} else {
					ifres79980 = False

				}

				ifres79979 = ifres79980

			} else {
				ifres79979 = False

			}

			if True == ifres79979 {
				tmp79966 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79967 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79968 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1cons)

				tmp79969 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79970 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79971 := Call(__e, tmp79970, V910)

				tmp79972 := Call(__e, tmp79969, tmp79971)

				tmp79973 := Call(__e, tmp79968, tmp79972)

				tmp79974 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79975 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79976 := Call(__e, tmp79975, V910)

				tmp79977 := Call(__e, tmp79974, tmp79976)

				tmp79978 := Call(__e, tmp79967, tmp79973, tmp79977)

				__e.TailApply(tmp79966, symmode, tmp79978)
				return

			} else {
				__e.Return(V910)
				return
			}

		}

	}, 1)

	tmp80058 := Call(__e, PrimNS1Value(symns2_1set), symshen_4eval_1cons, tmp79963)

	_ = tmp80058

	tmp80059 := MakeNative(func(__e *ControlFlow) {
		V912 := __e.Get(1)
		_ = V912
		tmp80060 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp80076 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80077 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp80078 := Call(__e, tmp80077)

			tmp80079 := Call(__e, tmp80076, YaccParse, tmp80078)

			if True == tmp80079 {
				tmp80062 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp80068 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp80069 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp80070 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp80071 := Call(__e, tmp80070)

					tmp80072 := Call(__e, tmp80069, tmp80071, Parse___5e_6)

					tmp80073 := Call(__e, tmp80068, tmp80072)

					if True == tmp80073 {
						tmp80065 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp80066 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80067 := Call(__e, tmp80066, Parse___5e_6)

						__e.TailApply(tmp80065, tmp80067, Nil)
						return

					} else {
						tmp80064 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp80064)
						return

					}

				}, 1)

				tmp80074 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp80075 := Call(__e, tmp80074, V912)

				__e.TailApply(tmp80062, tmp80075)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp80080 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5literal_d_6 := __e.Get(1)
			_ = Parse__shen_4_5literal_d_6
			tmp80103 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp80104 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80105 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp80106 := Call(__e, tmp80105)

			tmp80107 := Call(__e, tmp80104, tmp80106, Parse__shen_4_5literal_d_6)

			tmp80108 := Call(__e, tmp80103, tmp80107)

			if True == tmp80108 {
				tmp80083 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5body_d_6 := __e.Get(1)
					_ = Parse__shen_4_5body_d_6
					tmp80095 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp80096 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp80097 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp80098 := Call(__e, tmp80097)

					tmp80099 := Call(__e, tmp80096, tmp80098, Parse__shen_4_5body_d_6)

					tmp80100 := Call(__e, tmp80095, tmp80099)

					if True == tmp80100 {
						tmp80086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp80087 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80088 := Call(__e, tmp80087, Parse__shen_4_5body_d_6)

						tmp80089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp80090 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp80091 := Call(__e, tmp80090, Parse__shen_4_5literal_d_6)

						tmp80092 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp80093 := Call(__e, tmp80092, Parse__shen_4_5body_d_6)

						tmp80094 := Call(__e, tmp80089, tmp80091, tmp80093)

						__e.TailApply(tmp80086, tmp80088, tmp80094)
						return

					} else {
						tmp80085 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp80085)
						return

					}

				}, 1)

				tmp80101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5body_d_6)

				tmp80102 := Call(__e, tmp80101, Parse__shen_4_5literal_d_6)

				__e.TailApply(tmp80083, tmp80102)
				return

			} else {
				tmp80082 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp80082)
				return

			}

		}, 1)

		tmp80109 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5literal_d_6)

		tmp80110 := Call(__e, tmp80109, V912)

		tmp80111 := Call(__e, tmp80080, tmp80110)

		__e.TailApply(tmp80060, tmp80111)
		return

	}, 1)

	tmp80112 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5body_d_6, tmp80059)

	_ = tmp80112

	tmp80113 := MakeNative(func(__e *ControlFlow) {
		V915 := __e.Get(1)
		_ = V915
		tmp80114 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp80138 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80139 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp80140 := Call(__e, tmp80139)

			tmp80141 := Call(__e, tmp80138, YaccParse, tmp80140)

			if True == tmp80141 {
				tmp80134 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80135 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80136 := Call(__e, tmp80135, V915)

				tmp80137 := Call(__e, tmp80134, tmp80136)

				if True == tmp80137 {
					tmp80118 := MakeNative(func(__e *ControlFlow) {
						Parse__X := __e.Get(1)
						_ = Parse__X
						tmp80130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp80131 := Call(__e, tmp80130, Parse__X)

						if True == tmp80131 {
							tmp80121 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp80122 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp80124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

							tmp80125 := Call(__e, tmp80124, V915)

							tmp80126 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp80127 := Call(__e, tmp80126, V915)

							tmp80128 := Call(__e, tmp80123, tmp80125, tmp80127)

							tmp80129 := Call(__e, tmp80122, tmp80128)

							__e.TailApply(tmp80121, tmp80129, Parse__X)
							return

						} else {
							tmp80120 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp80120)
							return

						}

					}, 1)

					tmp80132 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp80133 := Call(__e, tmp80132, V915)

					__e.TailApply(tmp80118, tmp80133)
					return

				} else {
					tmp80117 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp80117)
					return

				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp80168 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80169 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp80170 := Call(__e, tmp80169, V915)

		tmp80171 := Call(__e, tmp80168, tmp80170)

		var ifres80162 Obj

		if True == tmp80171 {
			tmp80164 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80165 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp80166 := Call(__e, tmp80165, V915)

			tmp80167 := Call(__e, tmp80164, sym_b, tmp80166)

			var ifres80163 Obj

			if True == tmp80167 {
				ifres80163 = True

			} else {
				ifres80163 = False

			}

			ifres80162 = ifres80163

		} else {
			ifres80162 = False

		}

		var ifres80142 Obj

		if True == ifres80162 {
			tmp80145 := MakeNative(func(__e *ControlFlow) {
				NewStream913 := __e.Get(1)
				_ = NewStream913
				tmp80146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp80147 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80148 := Call(__e, tmp80147, NewStream913)

				tmp80149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp80150 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp80151 := Call(__e, PrimNS1Value(symns2_1value), symintern)

				tmp80152 := Call(__e, tmp80151, MakeString("Throwcontrol"))

				tmp80153 := Call(__e, tmp80150, tmp80152, Nil)

				tmp80154 := Call(__e, tmp80149, symcut, tmp80153)

				__e.TailApply(tmp80146, tmp80148, tmp80154)
				return

			}, 1)

			tmp80155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp80156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp80157 := Call(__e, tmp80156, V915)

			tmp80158 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp80159 := Call(__e, tmp80158, V915)

			tmp80160 := Call(__e, tmp80155, tmp80157, tmp80159)

			tmp80161 := Call(__e, tmp80145, tmp80160)

			ifres80142 = tmp80161

		} else {
			tmp80143 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp80144 := Call(__e, tmp80143)

			ifres80142 = tmp80144

		}

		__e.TailApply(tmp80114, ifres80142)
		return

	}, 1)

	tmp80172 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5literal_d_6, tmp80113)

	_ = tmp80172

	tmp80173 := MakeNative(func(__e *ControlFlow) {
		V917 := __e.Get(1)
		_ = V917
		tmp80192 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80193 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp80194 := Call(__e, tmp80193, V917)

		tmp80195 := Call(__e, tmp80192, tmp80194)

		if True == tmp80195 {
			tmp80176 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp80188 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp80189 := Call(__e, tmp80188, Parse__X, sym_k)

				if True == tmp80189 {
					tmp80179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp80180 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80181 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp80182 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp80183 := Call(__e, tmp80182, V917)

					tmp80184 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp80185 := Call(__e, tmp80184, V917)

					tmp80186 := Call(__e, tmp80181, tmp80183, tmp80185)

					tmp80187 := Call(__e, tmp80180, tmp80186)

					__e.TailApply(tmp80179, tmp80187, Parse__X)
					return

				} else {
					tmp80178 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp80178)
					return

				}

			}, 1)

			tmp80190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp80191 := Call(__e, tmp80190, V917)

			__e.TailApply(tmp80176, tmp80191)
			return

		} else {
			tmp80175 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp80175)
			return

		}

	}, 1)

	tmp80196 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5end_d_6, tmp80173)

	_ = tmp80196

	tmp80197 := MakeNative(func(__e *ControlFlow) {
		V921 := __e.Get(1)
		_ = V921
		V922 := __e.Get(2)
		_ = V922
		V923 := __e.Get(3)
		_ = V923
		tmp80198 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp80200 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80201 := Call(__e, tmp80200, Result, False)

			if True == tmp80201 {
				__e.Return(V921)
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp80202 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

		tmp80203 := Call(__e, tmp80202, V923)

		__e.TailApply(tmp80198, tmp80203)
		return

	}, 3)

	tmp80204 := Call(__e, PrimNS1Value(symns2_1set), symcut, tmp80197)

	_ = tmp80204

	tmp80205 := MakeNative(func(__e *ControlFlow) {
		V925 := __e.Get(1)
		_ = V925
		tmp80262 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80263 := Call(__e, tmp80262, V925)

		var ifres80232 Obj

		if True == tmp80263 {
			tmp80258 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80259 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80260 := Call(__e, tmp80259, V925)

			tmp80261 := Call(__e, tmp80258, symmode, tmp80260)

			var ifres80234 Obj

			if True == tmp80261 {
				tmp80254 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80255 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80256 := Call(__e, tmp80255, V925)

				tmp80257 := Call(__e, tmp80254, tmp80256)

				var ifres80236 Obj

				if True == tmp80257 {
					tmp80248 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80251 := Call(__e, tmp80250, V925)

					tmp80252 := Call(__e, tmp80249, tmp80251)

					tmp80253 := Call(__e, tmp80248, tmp80252)

					var ifres80238 Obj

					if True == tmp80253 {
						tmp80240 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80241 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80242 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80243 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80244 := Call(__e, tmp80243, V925)

						tmp80245 := Call(__e, tmp80242, tmp80244)

						tmp80246 := Call(__e, tmp80241, tmp80245)

						tmp80247 := Call(__e, tmp80240, Nil, tmp80246)

						var ifres80239 Obj

						if True == tmp80247 {
							ifres80239 = True

						} else {
							ifres80239 = False

						}

						ifres80238 = ifres80239

					} else {
						ifres80238 = False

					}

					var ifres80237 Obj

					if True == ifres80238 {
						ifres80237 = True

					} else {
						ifres80237 = False

					}

					ifres80236 = ifres80237

				} else {
					ifres80236 = False

				}

				var ifres80235 Obj

				if True == ifres80236 {
					ifres80235 = True

				} else {
					ifres80235 = False

				}

				ifres80234 = ifres80235

			} else {
				ifres80234 = False

			}

			var ifres80233 Obj

			if True == ifres80234 {
				ifres80233 = True

			} else {
				ifres80233 = False

			}

			ifres80232 = ifres80233

		} else {
			ifres80232 = False

		}

		if True == ifres80232 {
			__e.Return(V925)
			return
		} else {
			tmp80230 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80231 := Call(__e, tmp80230, Nil, V925)

			if True == tmp80231 {
				__e.Return(Nil)
				return
			} else {
				tmp80228 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80229 := Call(__e, tmp80228, V925)

				if True == tmp80229 {
					tmp80209 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80210 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80211 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80212 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80213 := Call(__e, tmp80212, V925)

					tmp80214 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80215 := Call(__e, tmp80214, sym_7, Nil)

					tmp80216 := Call(__e, tmp80211, tmp80213, tmp80215)

					tmp80217 := Call(__e, tmp80210, symmode, tmp80216)

					tmp80218 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80219 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert__modes)

					tmp80221 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80222 := Call(__e, tmp80221, V925)

					tmp80223 := Call(__e, tmp80220, tmp80222)

					tmp80224 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80225 := Call(__e, tmp80224, sym_1, Nil)

					tmp80226 := Call(__e, tmp80219, tmp80223, tmp80225)

					tmp80227 := Call(__e, tmp80218, symmode, tmp80226)

					__e.TailApply(tmp80209, tmp80217, tmp80227)
					return

				} else {
					__e.Return(V925)
					return
				}

			}

		}

	}, 1)

	tmp80264 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert__modes, tmp80205)

	_ = tmp80264

	tmp80265 := MakeNative(func(__e *ControlFlow) {
		V927 := __e.Get(1)
		_ = V927
		tmp80266 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp80267 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp80268 := Call(__e, PrimNS1Value(symns2_1value), symeval)

			__e.TailApply(tmp80268, X)
			return

		}, 1)

		tmp80269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1_6shen)

		tmp80270 := Call(__e, tmp80269, V927)

		__e.TailApply(tmp80266, tmp80267, tmp80270)
		return

	}, 1)

	tmp80271 := Call(__e, PrimNS1Value(symns2_1set), symshen_4s_1prolog, tmp80265)

	_ = tmp80271

	tmp80272 := MakeNative(func(__e *ControlFlow) {
		V929 := __e.Get(1)
		_ = V929
		tmp80273 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp80274 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp80275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compile__prolog__procedure)

			__e.TailApply(tmp80275, X)
			return

		}, 1)

		tmp80276 := Call(__e, PrimNS1Value(symns2_1value), symshen_4group__clauses)

		tmp80277 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp80278 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp80279 := Call(__e, PrimNS1Value(symns2_1value), symshen_4s_1prolog__clause)

			__e.TailApply(tmp80279, X)
			return

		}, 1)

		tmp80280 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

		tmp80281 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp80282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4head__abstraction)

			__e.TailApply(tmp80282, X)
			return

		}, 1)

		tmp80283 := Call(__e, tmp80280, tmp80281, V929)

		tmp80284 := Call(__e, tmp80277, tmp80278, tmp80283)

		tmp80285 := Call(__e, tmp80276, tmp80284)

		__e.TailApply(tmp80273, tmp80274, tmp80285)
		return

	}, 1)

	tmp80286 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1_6shen, tmp80272)

	_ = tmp80286

	tmp80287 := MakeNative(func(__e *ControlFlow) {
		V931 := __e.Get(1)
		_ = V931
		tmp80339 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80340 := Call(__e, tmp80339, V931)

		var ifres80307 Obj

		if True == tmp80340 {
			tmp80335 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp80336 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80337 := Call(__e, tmp80336, V931)

			tmp80338 := Call(__e, tmp80335, tmp80337)

			var ifres80309 Obj

			if True == tmp80338 {
				tmp80329 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp80330 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80331 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80332 := Call(__e, tmp80331, V931)

				tmp80333 := Call(__e, tmp80330, tmp80332)

				tmp80334 := Call(__e, tmp80329, sym_h_1, tmp80333)

				var ifres80311 Obj

				if True == tmp80334 {
					tmp80323 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80324 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80325 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80326 := Call(__e, tmp80325, V931)

					tmp80327 := Call(__e, tmp80324, tmp80326)

					tmp80328 := Call(__e, tmp80323, tmp80327)

					var ifres80313 Obj

					if True == tmp80328 {
						tmp80315 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80316 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80317 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80318 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80319 := Call(__e, tmp80318, V931)

						tmp80320 := Call(__e, tmp80317, tmp80319)

						tmp80321 := Call(__e, tmp80316, tmp80320)

						tmp80322 := Call(__e, tmp80315, Nil, tmp80321)

						var ifres80314 Obj

						if True == tmp80322 {
							ifres80314 = True

						} else {
							ifres80314 = False

						}

						ifres80313 = ifres80314

					} else {
						ifres80313 = False

					}

					var ifres80312 Obj

					if True == ifres80313 {
						ifres80312 = True

					} else {
						ifres80312 = False

					}

					ifres80311 = ifres80312

				} else {
					ifres80311 = False

				}

				var ifres80310 Obj

				if True == ifres80311 {
					ifres80310 = True

				} else {
					ifres80310 = False

				}

				ifres80309 = ifres80310

			} else {
				ifres80309 = False

			}

			var ifres80308 Obj

			if True == ifres80309 {
				ifres80308 = True

			} else {
				ifres80308 = False

			}

			ifres80307 = ifres80308

		} else {
			ifres80307 = False

		}

		if True == ifres80307 {
			tmp80290 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp80291 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80292 := Call(__e, tmp80291, V931)

			tmp80293 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp80294 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp80295 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp80296 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp80297 := Call(__e, PrimNS1Value(symns2_1value), symshen_4s_1prolog__literal)

				__e.TailApply(tmp80297, X)
				return

			}, 1)

			tmp80298 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80299 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80300 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80301 := Call(__e, tmp80300, V931)

			tmp80302 := Call(__e, tmp80299, tmp80301)

			tmp80303 := Call(__e, tmp80298, tmp80302)

			tmp80304 := Call(__e, tmp80295, tmp80296, tmp80303)

			tmp80305 := Call(__e, tmp80294, tmp80304, Nil)

			tmp80306 := Call(__e, tmp80293, sym_h_1, tmp80305)

			__e.TailApply(tmp80290, tmp80292, tmp80306)
			return

		} else {
			tmp80289 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp80289, symshen_4s_1prolog__clause)
			return

		}

	}, 1)

	tmp80341 := Call(__e, PrimNS1Value(symns2_1set), symshen_4s_1prolog__clause, tmp80287)

	_ = tmp80341

	tmp80342 := MakeNative(func(__e *ControlFlow) {
		V933 := __e.Get(1)
		_ = V933
		tmp80480 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80481 := Call(__e, tmp80480, V933)

		var ifres80436 Obj

		if True == tmp80481 {
			tmp80476 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp80477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80478 := Call(__e, tmp80477, V933)

			tmp80479 := Call(__e, tmp80476, tmp80478)

			var ifres80438 Obj

			if True == tmp80479 {
				tmp80470 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp80471 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80472 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80473 := Call(__e, tmp80472, V933)

				tmp80474 := Call(__e, tmp80471, tmp80473)

				tmp80475 := Call(__e, tmp80470, sym_h_1, tmp80474)

				var ifres80440 Obj

				if True == tmp80475 {
					tmp80464 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80465 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80466 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80467 := Call(__e, tmp80466, V933)

					tmp80468 := Call(__e, tmp80465, tmp80467)

					tmp80469 := Call(__e, tmp80464, tmp80468)

					var ifres80442 Obj

					if True == tmp80469 {
						tmp80456 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80457 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80458 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80459 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80460 := Call(__e, tmp80459, V933)

						tmp80461 := Call(__e, tmp80458, tmp80460)

						tmp80462 := Call(__e, tmp80457, tmp80461)

						tmp80463 := Call(__e, tmp80456, Nil, tmp80462)

						var ifres80444 Obj

						if True == tmp80463 {
							tmp80446 := MakeNative(func(__e *ControlFlow) {
								tmp80447 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

								tmp80448 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity__head)

								tmp80449 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp80450 := Call(__e, tmp80449, V933)

								tmp80451 := Call(__e, tmp80448, tmp80450)

								tmp80452 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

								tmp80453 := Call(__e, tmp80452, symshen_4_dmaxcomplexity_d)

								__e.TailApply(tmp80447, tmp80451, tmp80453)
								return

							}, 0)

							tmp80454 := MakeNative(func(__e *ControlFlow) {
								__ := __e.Get(1)
								_ = __
								__e.Return(False)
								return
							}, 1)

							tmp80455 := Call(__e, PrimNS1Value(symtry_1catch), tmp80446, tmp80454)

							var ifres80445 Obj

							if True == tmp80455 {
								ifres80445 = True

							} else {
								ifres80445 = False

							}

							ifres80444 = ifres80445

						} else {
							ifres80444 = False

						}

						var ifres80443 Obj

						if True == ifres80444 {
							ifres80443 = True

						} else {
							ifres80443 = False

						}

						ifres80442 = ifres80443

					} else {
						ifres80442 = False

					}

					var ifres80441 Obj

					if True == ifres80442 {
						ifres80441 = True

					} else {
						ifres80441 = False

					}

					ifres80440 = ifres80441

				} else {
					ifres80440 = False

				}

				var ifres80439 Obj

				if True == ifres80440 {
					ifres80439 = True

				} else {
					ifres80439 = False

				}

				ifres80438 = ifres80439

			} else {
				ifres80438 = False

			}

			var ifres80437 Obj

			if True == ifres80438 {
				ifres80437 = True

			} else {
				ifres80437 = False

			}

			ifres80436 = ifres80437

		} else {
			ifres80436 = False

		}

		if True == ifres80436 {
			tmp80435 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp80435, V933, Nil)
			return

		} else {
			tmp80433 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp80434 := Call(__e, tmp80433, V933)

			var ifres80395 Obj

			if True == tmp80434 {
				tmp80429 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80430 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80431 := Call(__e, tmp80430, V933)

				tmp80432 := Call(__e, tmp80429, tmp80431)

				var ifres80397 Obj

				if True == tmp80432 {
					tmp80425 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80426 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80427 := Call(__e, tmp80426, V933)

					tmp80428 := Call(__e, tmp80425, tmp80427)

					var ifres80399 Obj

					if True == tmp80428 {
						tmp80419 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80420 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80421 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80422 := Call(__e, tmp80421, V933)

						tmp80423 := Call(__e, tmp80420, tmp80422)

						tmp80424 := Call(__e, tmp80419, sym_h_1, tmp80423)

						var ifres80401 Obj

						if True == tmp80424 {
							tmp80413 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80414 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80415 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80416 := Call(__e, tmp80415, V933)

							tmp80417 := Call(__e, tmp80414, tmp80416)

							tmp80418 := Call(__e, tmp80413, tmp80417)

							var ifres80403 Obj

							if True == tmp80418 {
								tmp80405 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp80406 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80407 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80408 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80409 := Call(__e, tmp80408, V933)

								tmp80410 := Call(__e, tmp80407, tmp80409)

								tmp80411 := Call(__e, tmp80406, tmp80410)

								tmp80412 := Call(__e, tmp80405, Nil, tmp80411)

								var ifres80404 Obj

								if True == tmp80412 {
									ifres80404 = True

								} else {
									ifres80404 = False

								}

								ifres80403 = ifres80404

							} else {
								ifres80403 = False

							}

							var ifres80402 Obj

							if True == ifres80403 {
								ifres80402 = True

							} else {
								ifres80402 = False

							}

							ifres80401 = ifres80402

						} else {
							ifres80401 = False

						}

						var ifres80400 Obj

						if True == ifres80401 {
							ifres80400 = True

						} else {
							ifres80400 = False

						}

						ifres80399 = ifres80400

					} else {
						ifres80399 = False

					}

					var ifres80398 Obj

					if True == ifres80399 {
						ifres80398 = True

					} else {
						ifres80398 = False

					}

					ifres80397 = ifres80398

				} else {
					ifres80397 = False

				}

				var ifres80396 Obj

				if True == ifres80397 {
					ifres80396 = True

				} else {
					ifres80396 = False

				}

				ifres80395 = ifres80396

			} else {
				ifres80395 = False

			}

			if True == ifres80395 {
				tmp80346 := MakeNative(func(__e *ControlFlow) {
					Terms := __e.Get(1)
					_ = Terms
					tmp80347 := MakeNative(func(__e *ControlFlow) {
						XTerms := __e.Get(1)
						_ = XTerms
						tmp80348 := MakeNative(func(__e *ControlFlow) {
							Literal := __e.Get(1)
							_ = Literal
							tmp80349 := MakeNative(func(__e *ControlFlow) {
								Clause := __e.Get(1)
								_ = Clause
								tmp80350 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								__e.TailApply(tmp80350, Clause, Nil)
								return

							}, 1)

							tmp80351 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp80352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp80353 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80354 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80355 := Call(__e, tmp80354, V933)

							tmp80356 := Call(__e, tmp80353, tmp80355)

							tmp80357 := Call(__e, tmp80352, tmp80356, Terms)

							tmp80358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp80359 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp80360 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp80361 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80362 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80363 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80364 := Call(__e, tmp80363, V933)

							tmp80365 := Call(__e, tmp80362, tmp80364)

							tmp80366 := Call(__e, tmp80361, tmp80365)

							tmp80367 := Call(__e, tmp80360, Literal, tmp80366)

							tmp80368 := Call(__e, tmp80359, tmp80367, Nil)

							tmp80369 := Call(__e, tmp80358, sym_h_1, tmp80368)

							tmp80370 := Call(__e, tmp80351, tmp80357, tmp80369)

							__e.TailApply(tmp80349, tmp80370)
							return

						}, 1)

						tmp80371 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp80372 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp80373 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cons__form)

						tmp80374 := Call(__e, tmp80373, Terms)

						tmp80375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp80376 := Call(__e, tmp80375, XTerms, Nil)

						tmp80377 := Call(__e, tmp80372, tmp80374, tmp80376)

						tmp80378 := Call(__e, tmp80371, symunify, tmp80377)

						__e.TailApply(tmp80348, tmp80378)
						return

					}, 1)

					tmp80379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

					tmp80380 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove__modes)

					tmp80381 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80382 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80383 := Call(__e, tmp80382, V933)

					tmp80384 := Call(__e, tmp80381, tmp80383)

					tmp80385 := Call(__e, tmp80380, tmp80384)

					tmp80386 := Call(__e, tmp80379, tmp80385)

					__e.TailApply(tmp80347, tmp80386)
					return

				}, 1)

				tmp80387 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp80388 := MakeNative(func(__e *ControlFlow) {
					Y := __e.Get(1)
					_ = Y
					tmp80389 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

					__e.TailApply(tmp80389, symV)
					return

				}, 1)

				tmp80390 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80391 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80392 := Call(__e, tmp80391, V933)

				tmp80393 := Call(__e, tmp80390, tmp80392)

				tmp80394 := Call(__e, tmp80387, tmp80388, tmp80393)

				__e.TailApply(tmp80346, tmp80394)
				return

			} else {
				tmp80345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp80345, symshen_4head__abstraction)
				return

			}

		}

	}, 1)

	tmp80482 := Call(__e, PrimNS1Value(symns2_1set), symshen_4head__abstraction, tmp80342)

	_ = tmp80482

	tmp80483 := MakeNative(func(__e *ControlFlow) {
		V939 := __e.Get(1)
		_ = V939
		tmp80493 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80494 := Call(__e, tmp80493, V939)

		if True == tmp80494 {
			tmp80486 := Call(__e, PrimNS1Value(symns2_1value), symshen_4safe_1product)

			tmp80487 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp80488 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp80489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

				__e.TailApply(tmp80489, X)
				return

			}, 1)

			tmp80490 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80491 := Call(__e, tmp80490, V939)

			tmp80492 := Call(__e, tmp80487, tmp80488, tmp80491)

			__e.TailApply(tmp80486, tmp80492)
			return

		} else {
			tmp80485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp80485, symshen_4complexity__head)
			return

		}

	}, 1)

	tmp80495 := Call(__e, PrimNS1Value(symns2_1set), symshen_4complexity__head, tmp80483)

	_ = tmp80495

	tmp80496 := MakeNative(func(__e *ControlFlow) {
		V942 := __e.Get(1)
		_ = V942
		V943 := __e.Get(2)
		_ = V943
		tmp80497 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

		__e.TailApply(tmp80497, V942, V943)
		return

	}, 2)

	tmp80498 := Call(__e, PrimNS1Value(symns2_1set), symshen_4safe_1multiply, tmp80496)

	_ = tmp80498

	tmp80499 := MakeNative(func(__e *ControlFlow) {
		V952 := __e.Get(1)
		_ = V952
		tmp80894 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp80895 := Call(__e, tmp80894, V952)

		var ifres80810 Obj

		if True == tmp80895 {
			tmp80890 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp80891 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80892 := Call(__e, tmp80891, V952)

			tmp80893 := Call(__e, tmp80890, symmode, tmp80892)

			var ifres80812 Obj

			if True == tmp80893 {
				tmp80886 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80887 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80888 := Call(__e, tmp80887, V952)

				tmp80889 := Call(__e, tmp80886, tmp80888)

				var ifres80814 Obj

				if True == tmp80889 {
					tmp80880 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80881 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80882 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80883 := Call(__e, tmp80882, V952)

					tmp80884 := Call(__e, tmp80881, tmp80883)

					tmp80885 := Call(__e, tmp80880, tmp80884)

					var ifres80816 Obj

					if True == tmp80885 {
						tmp80872 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80873 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80874 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80875 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80876 := Call(__e, tmp80875, V952)

						tmp80877 := Call(__e, tmp80874, tmp80876)

						tmp80878 := Call(__e, tmp80873, tmp80877)

						tmp80879 := Call(__e, tmp80872, symmode, tmp80878)

						var ifres80818 Obj

						if True == tmp80879 {
							tmp80864 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80865 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80866 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80868 := Call(__e, tmp80867, V952)

							tmp80869 := Call(__e, tmp80866, tmp80868)

							tmp80870 := Call(__e, tmp80865, tmp80869)

							tmp80871 := Call(__e, tmp80864, tmp80870)

							var ifres80820 Obj

							if True == tmp80871 {
								tmp80854 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp80855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80856 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80857 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp80858 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80859 := Call(__e, tmp80858, V952)

								tmp80860 := Call(__e, tmp80857, tmp80859)

								tmp80861 := Call(__e, tmp80856, tmp80860)

								tmp80862 := Call(__e, tmp80855, tmp80861)

								tmp80863 := Call(__e, tmp80854, tmp80862)

								var ifres80822 Obj

								if True == tmp80863 {
									tmp80842 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp80843 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80846 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp80847 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80848 := Call(__e, tmp80847, V952)

									tmp80849 := Call(__e, tmp80846, tmp80848)

									tmp80850 := Call(__e, tmp80845, tmp80849)

									tmp80851 := Call(__e, tmp80844, tmp80850)

									tmp80852 := Call(__e, tmp80843, tmp80851)

									tmp80853 := Call(__e, tmp80842, Nil, tmp80852)

									var ifres80824 Obj

									if True == tmp80853 {
										tmp80836 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp80837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80838 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80839 := Call(__e, tmp80838, V952)

										tmp80840 := Call(__e, tmp80837, tmp80839)

										tmp80841 := Call(__e, tmp80836, tmp80840)

										var ifres80826 Obj

										if True == tmp80841 {
											tmp80828 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp80829 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80830 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80831 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80832 := Call(__e, tmp80831, V952)

											tmp80833 := Call(__e, tmp80830, tmp80832)

											tmp80834 := Call(__e, tmp80829, tmp80833)

											tmp80835 := Call(__e, tmp80828, Nil, tmp80834)

											var ifres80827 Obj

											if True == tmp80835 {
												ifres80827 = True

											} else {
												ifres80827 = False

											}

											ifres80826 = ifres80827

										} else {
											ifres80826 = False

										}

										var ifres80825 Obj

										if True == ifres80826 {
											ifres80825 = True

										} else {
											ifres80825 = False

										}

										ifres80824 = ifres80825

									} else {
										ifres80824 = False

									}

									var ifres80823 Obj

									if True == ifres80824 {
										ifres80823 = True

									} else {
										ifres80823 = False

									}

									ifres80822 = ifres80823

								} else {
									ifres80822 = False

								}

								var ifres80821 Obj

								if True == ifres80822 {
									ifres80821 = True

								} else {
									ifres80821 = False

								}

								ifres80820 = ifres80821

							} else {
								ifres80820 = False

							}

							var ifres80819 Obj

							if True == ifres80820 {
								ifres80819 = True

							} else {
								ifres80819 = False

							}

							ifres80818 = ifres80819

						} else {
							ifres80818 = False

						}

						var ifres80817 Obj

						if True == ifres80818 {
							ifres80817 = True

						} else {
							ifres80817 = False

						}

						ifres80816 = ifres80817

					} else {
						ifres80816 = False

					}

					var ifres80815 Obj

					if True == ifres80816 {
						ifres80815 = True

					} else {
						ifres80815 = False

					}

					ifres80814 = ifres80815

				} else {
					ifres80814 = False

				}

				var ifres80813 Obj

				if True == ifres80814 {
					ifres80813 = True

				} else {
					ifres80813 = False

				}

				ifres80812 = ifres80813

			} else {
				ifres80812 = False

			}

			var ifres80811 Obj

			if True == ifres80812 {
				ifres80811 = True

			} else {
				ifres80811 = False

			}

			ifres80810 = ifres80811

		} else {
			ifres80810 = False

		}

		if True == ifres80810 {
			tmp80805 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

			tmp80806 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp80807 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp80808 := Call(__e, tmp80807, V952)

			tmp80809 := Call(__e, tmp80806, tmp80808)

			__e.TailApply(tmp80805, tmp80809)
			return

		} else {
			tmp80803 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp80804 := Call(__e, tmp80803, V952)

			var ifres80755 Obj

			if True == tmp80804 {
				tmp80799 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp80800 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80801 := Call(__e, tmp80800, V952)

				tmp80802 := Call(__e, tmp80799, symmode, tmp80801)

				var ifres80757 Obj

				if True == tmp80802 {
					tmp80795 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80797 := Call(__e, tmp80796, V952)

					tmp80798 := Call(__e, tmp80795, tmp80797)

					var ifres80759 Obj

					if True == tmp80798 {
						tmp80789 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp80790 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80791 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80792 := Call(__e, tmp80791, V952)

						tmp80793 := Call(__e, tmp80790, tmp80792)

						tmp80794 := Call(__e, tmp80789, tmp80793)

						var ifres80761 Obj

						if True == tmp80794 {
							tmp80783 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80786 := Call(__e, tmp80785, V952)

							tmp80787 := Call(__e, tmp80784, tmp80786)

							tmp80788 := Call(__e, tmp80783, tmp80787)

							var ifres80763 Obj

							if True == tmp80788 {
								tmp80775 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp80776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp80777 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80779 := Call(__e, tmp80778, V952)

								tmp80780 := Call(__e, tmp80777, tmp80779)

								tmp80781 := Call(__e, tmp80776, tmp80780)

								tmp80782 := Call(__e, tmp80775, sym_7, tmp80781)

								var ifres80765 Obj

								if True == tmp80782 {
									tmp80767 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp80768 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80769 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80770 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80771 := Call(__e, tmp80770, V952)

									tmp80772 := Call(__e, tmp80769, tmp80771)

									tmp80773 := Call(__e, tmp80768, tmp80772)

									tmp80774 := Call(__e, tmp80767, Nil, tmp80773)

									var ifres80766 Obj

									if True == tmp80774 {
										ifres80766 = True

									} else {
										ifres80766 = False

									}

									ifres80765 = ifres80766

								} else {
									ifres80765 = False

								}

								var ifres80764 Obj

								if True == ifres80765 {
									ifres80764 = True

								} else {
									ifres80764 = False

								}

								ifres80763 = ifres80764

							} else {
								ifres80763 = False

							}

							var ifres80762 Obj

							if True == ifres80763 {
								ifres80762 = True

							} else {
								ifres80762 = False

							}

							ifres80761 = ifres80762

						} else {
							ifres80761 = False

						}

						var ifres80760 Obj

						if True == ifres80761 {
							ifres80760 = True

						} else {
							ifres80760 = False

						}

						ifres80759 = ifres80760

					} else {
						ifres80759 = False

					}

					var ifres80758 Obj

					if True == ifres80759 {
						ifres80758 = True

					} else {
						ifres80758 = False

					}

					ifres80757 = ifres80758

				} else {
					ifres80757 = False

				}

				var ifres80756 Obj

				if True == ifres80757 {
					ifres80756 = True

				} else {
					ifres80756 = False

				}

				ifres80755 = ifres80756

			} else {
				ifres80755 = False

			}

			if True == ifres80755 {
				tmp80720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4safe_1multiply)

				tmp80721 := Call(__e, PrimNS1Value(symns2_1value), symshen_4safe_1multiply)

				tmp80722 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

				tmp80723 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp80724 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp80725 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80726 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80727 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80728 := Call(__e, tmp80727, V952)

				tmp80729 := Call(__e, tmp80726, tmp80728)

				tmp80730 := Call(__e, tmp80725, tmp80729)

				tmp80731 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80732 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80733 := Call(__e, tmp80732, V952)

				tmp80734 := Call(__e, tmp80731, tmp80733)

				tmp80735 := Call(__e, tmp80724, tmp80730, tmp80734)

				tmp80736 := Call(__e, tmp80723, symmode, tmp80735)

				tmp80737 := Call(__e, tmp80722, tmp80736)

				tmp80738 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

				tmp80739 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp80740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp80741 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80742 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80743 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80744 := Call(__e, tmp80743, V952)

				tmp80745 := Call(__e, tmp80742, tmp80744)

				tmp80746 := Call(__e, tmp80741, tmp80745)

				tmp80747 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80748 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80749 := Call(__e, tmp80748, V952)

				tmp80750 := Call(__e, tmp80747, tmp80749)

				tmp80751 := Call(__e, tmp80740, tmp80746, tmp80750)

				tmp80752 := Call(__e, tmp80739, symmode, tmp80751)

				tmp80753 := Call(__e, tmp80738, tmp80752)

				tmp80754 := Call(__e, tmp80721, tmp80737, tmp80753)

				__e.TailApply(tmp80720, MakeNumber(2), tmp80754)
				return

			} else {
				tmp80718 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp80719 := Call(__e, tmp80718, V952)

				var ifres80670 Obj

				if True == tmp80719 {
					tmp80714 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp80715 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80716 := Call(__e, tmp80715, V952)

					tmp80717 := Call(__e, tmp80714, symmode, tmp80716)

					var ifres80672 Obj

					if True == tmp80717 {
						tmp80710 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp80711 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80712 := Call(__e, tmp80711, V952)

						tmp80713 := Call(__e, tmp80710, tmp80712)

						var ifres80674 Obj

						if True == tmp80713 {
							tmp80704 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80705 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80706 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80707 := Call(__e, tmp80706, V952)

							tmp80708 := Call(__e, tmp80705, tmp80707)

							tmp80709 := Call(__e, tmp80704, tmp80708)

							var ifres80676 Obj

							if True == tmp80709 {
								tmp80698 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp80699 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80701 := Call(__e, tmp80700, V952)

								tmp80702 := Call(__e, tmp80699, tmp80701)

								tmp80703 := Call(__e, tmp80698, tmp80702)

								var ifres80678 Obj

								if True == tmp80703 {
									tmp80690 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp80691 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp80692 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80693 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80694 := Call(__e, tmp80693, V952)

									tmp80695 := Call(__e, tmp80692, tmp80694)

									tmp80696 := Call(__e, tmp80691, tmp80695)

									tmp80697 := Call(__e, tmp80690, sym_1, tmp80696)

									var ifres80680 Obj

									if True == tmp80697 {
										tmp80682 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp80683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80684 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80685 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80686 := Call(__e, tmp80685, V952)

										tmp80687 := Call(__e, tmp80684, tmp80686)

										tmp80688 := Call(__e, tmp80683, tmp80687)

										tmp80689 := Call(__e, tmp80682, Nil, tmp80688)

										var ifres80681 Obj

										if True == tmp80689 {
											ifres80681 = True

										} else {
											ifres80681 = False

										}

										ifres80680 = ifres80681

									} else {
										ifres80680 = False

									}

									var ifres80679 Obj

									if True == ifres80680 {
										ifres80679 = True

									} else {
										ifres80679 = False

									}

									ifres80678 = ifres80679

								} else {
									ifres80678 = False

								}

								var ifres80677 Obj

								if True == ifres80678 {
									ifres80677 = True

								} else {
									ifres80677 = False

								}

								ifres80676 = ifres80677

							} else {
								ifres80676 = False

							}

							var ifres80675 Obj

							if True == ifres80676 {
								ifres80675 = True

							} else {
								ifres80675 = False

							}

							ifres80674 = ifres80675

						} else {
							ifres80674 = False

						}

						var ifres80673 Obj

						if True == ifres80674 {
							ifres80673 = True

						} else {
							ifres80673 = False

						}

						ifres80672 = ifres80673

					} else {
						ifres80672 = False

					}

					var ifres80671 Obj

					if True == ifres80672 {
						ifres80671 = True

					} else {
						ifres80671 = False

					}

					ifres80670 = ifres80671

				} else {
					ifres80670 = False

				}

				if True == ifres80670 {
					tmp80637 := Call(__e, PrimNS1Value(symns2_1value), symshen_4safe_1multiply)

					tmp80638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

					tmp80639 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80640 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80641 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80642 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80643 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80644 := Call(__e, tmp80643, V952)

					tmp80645 := Call(__e, tmp80642, tmp80644)

					tmp80646 := Call(__e, tmp80641, tmp80645)

					tmp80647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80648 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80649 := Call(__e, tmp80648, V952)

					tmp80650 := Call(__e, tmp80647, tmp80649)

					tmp80651 := Call(__e, tmp80640, tmp80646, tmp80650)

					tmp80652 := Call(__e, tmp80639, symmode, tmp80651)

					tmp80653 := Call(__e, tmp80638, tmp80652)

					tmp80654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

					tmp80655 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80656 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80657 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80658 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80660 := Call(__e, tmp80659, V952)

					tmp80661 := Call(__e, tmp80658, tmp80660)

					tmp80662 := Call(__e, tmp80657, tmp80661)

					tmp80663 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80664 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80665 := Call(__e, tmp80664, V952)

					tmp80666 := Call(__e, tmp80663, tmp80665)

					tmp80667 := Call(__e, tmp80656, tmp80662, tmp80666)

					tmp80668 := Call(__e, tmp80655, symmode, tmp80667)

					tmp80669 := Call(__e, tmp80654, tmp80668)

					__e.TailApply(tmp80637, tmp80653, tmp80669)
					return

				} else {
					tmp80635 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80636 := Call(__e, tmp80635, V952)

					var ifres80597 Obj

					if True == tmp80636 {
						tmp80631 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80632 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80633 := Call(__e, tmp80632, V952)

						tmp80634 := Call(__e, tmp80631, symmode, tmp80633)

						var ifres80599 Obj

						if True == tmp80634 {
							tmp80627 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80628 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80629 := Call(__e, tmp80628, V952)

							tmp80630 := Call(__e, tmp80627, tmp80629)

							var ifres80601 Obj

							if True == tmp80630 {
								tmp80621 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp80622 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80624 := Call(__e, tmp80623, V952)

								tmp80625 := Call(__e, tmp80622, tmp80624)

								tmp80626 := Call(__e, tmp80621, tmp80625)

								var ifres80603 Obj

								if True == tmp80626 {
									tmp80613 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp80614 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80615 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80616 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80617 := Call(__e, tmp80616, V952)

									tmp80618 := Call(__e, tmp80615, tmp80617)

									tmp80619 := Call(__e, tmp80614, tmp80618)

									tmp80620 := Call(__e, tmp80613, Nil, tmp80619)

									var ifres80605 Obj

									if True == tmp80620 {
										tmp80607 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

										tmp80608 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp80609 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80610 := Call(__e, tmp80609, V952)

										tmp80611 := Call(__e, tmp80608, tmp80610)

										tmp80612 := Call(__e, tmp80607, tmp80611)

										var ifres80606 Obj

										if True == tmp80612 {
											ifres80606 = True

										} else {
											ifres80606 = False

										}

										ifres80605 = ifres80606

									} else {
										ifres80605 = False

									}

									var ifres80604 Obj

									if True == ifres80605 {
										ifres80604 = True

									} else {
										ifres80604 = False

									}

									ifres80603 = ifres80604

								} else {
									ifres80603 = False

								}

								var ifres80602 Obj

								if True == ifres80603 {
									ifres80602 = True

								} else {
									ifres80602 = False

								}

								ifres80601 = ifres80602

							} else {
								ifres80601 = False

							}

							var ifres80600 Obj

							if True == ifres80601 {
								ifres80600 = True

							} else {
								ifres80600 = False

							}

							ifres80599 = ifres80600

						} else {
							ifres80599 = False

						}

						var ifres80598 Obj

						if True == ifres80599 {
							ifres80598 = True

						} else {
							ifres80598 = False

						}

						ifres80597 = ifres80598

					} else {
						ifres80597 = False

					}

					if True == ifres80597 {
						__e.Return(MakeNumber(1))
						return
					} else {
						tmp80595 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp80596 := Call(__e, tmp80595, V952)

						var ifres80555 Obj

						if True == tmp80596 {
							tmp80591 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp80592 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp80593 := Call(__e, tmp80592, V952)

							tmp80594 := Call(__e, tmp80591, symmode, tmp80593)

							var ifres80557 Obj

							if True == tmp80594 {
								tmp80587 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp80588 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80589 := Call(__e, tmp80588, V952)

								tmp80590 := Call(__e, tmp80587, tmp80589)

								var ifres80559 Obj

								if True == tmp80590 {
									tmp80581 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp80582 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80584 := Call(__e, tmp80583, V952)

									tmp80585 := Call(__e, tmp80582, tmp80584)

									tmp80586 := Call(__e, tmp80581, tmp80585)

									var ifres80561 Obj

									if True == tmp80586 {
										tmp80573 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp80574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp80575 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80577 := Call(__e, tmp80576, V952)

										tmp80578 := Call(__e, tmp80575, tmp80577)

										tmp80579 := Call(__e, tmp80574, tmp80578)

										tmp80580 := Call(__e, tmp80573, sym_7, tmp80579)

										var ifres80563 Obj

										if True == tmp80580 {
											tmp80565 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp80566 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80567 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80569 := Call(__e, tmp80568, V952)

											tmp80570 := Call(__e, tmp80567, tmp80569)

											tmp80571 := Call(__e, tmp80566, tmp80570)

											tmp80572 := Call(__e, tmp80565, Nil, tmp80571)

											var ifres80564 Obj

											if True == tmp80572 {
												ifres80564 = True

											} else {
												ifres80564 = False

											}

											ifres80563 = ifres80564

										} else {
											ifres80563 = False

										}

										var ifres80562 Obj

										if True == ifres80563 {
											ifres80562 = True

										} else {
											ifres80562 = False

										}

										ifres80561 = ifres80562

									} else {
										ifres80561 = False

									}

									var ifres80560 Obj

									if True == ifres80561 {
										ifres80560 = True

									} else {
										ifres80560 = False

									}

									ifres80559 = ifres80560

								} else {
									ifres80559 = False

								}

								var ifres80558 Obj

								if True == ifres80559 {
									ifres80558 = True

								} else {
									ifres80558 = False

								}

								ifres80557 = ifres80558

							} else {
								ifres80557 = False

							}

							var ifres80556 Obj

							if True == ifres80557 {
								ifres80556 = True

							} else {
								ifres80556 = False

							}

							ifres80555 = ifres80556

						} else {
							ifres80555 = False

						}

						if True == ifres80555 {
							__e.Return(MakeNumber(2))
							return
						} else {
							tmp80553 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80554 := Call(__e, tmp80553, V952)

							var ifres80513 Obj

							if True == tmp80554 {
								tmp80549 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp80550 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp80551 := Call(__e, tmp80550, V952)

								tmp80552 := Call(__e, tmp80549, symmode, tmp80551)

								var ifres80515 Obj

								if True == tmp80552 {
									tmp80545 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp80546 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp80547 := Call(__e, tmp80546, V952)

									tmp80548 := Call(__e, tmp80545, tmp80547)

									var ifres80517 Obj

									if True == tmp80548 {
										tmp80539 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp80540 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80541 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp80542 := Call(__e, tmp80541, V952)

										tmp80543 := Call(__e, tmp80540, tmp80542)

										tmp80544 := Call(__e, tmp80539, tmp80543)

										var ifres80519 Obj

										if True == tmp80544 {
											tmp80531 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp80532 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp80533 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80534 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp80535 := Call(__e, tmp80534, V952)

											tmp80536 := Call(__e, tmp80533, tmp80535)

											tmp80537 := Call(__e, tmp80532, tmp80536)

											tmp80538 := Call(__e, tmp80531, sym_1, tmp80537)

											var ifres80521 Obj

											if True == tmp80538 {
												tmp80523 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp80524 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp80525 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp80526 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp80527 := Call(__e, tmp80526, V952)

												tmp80528 := Call(__e, tmp80525, tmp80527)

												tmp80529 := Call(__e, tmp80524, tmp80528)

												tmp80530 := Call(__e, tmp80523, Nil, tmp80529)

												var ifres80522 Obj

												if True == tmp80530 {
													ifres80522 = True

												} else {
													ifres80522 = False

												}

												ifres80521 = ifres80522

											} else {
												ifres80521 = False

											}

											var ifres80520 Obj

											if True == ifres80521 {
												ifres80520 = True

											} else {
												ifres80520 = False

											}

											ifres80519 = ifres80520

										} else {
											ifres80519 = False

										}

										var ifres80518 Obj

										if True == ifres80519 {
											ifres80518 = True

										} else {
											ifres80518 = False

										}

										ifres80517 = ifres80518

									} else {
										ifres80517 = False

									}

									var ifres80516 Obj

									if True == ifres80517 {
										ifres80516 = True

									} else {
										ifres80516 = False

									}

									ifres80515 = ifres80516

								} else {
									ifres80515 = False

								}

								var ifres80514 Obj

								if True == ifres80515 {
									ifres80514 = True

								} else {
									ifres80514 = False

								}

								ifres80513 = ifres80514

							} else {
								ifres80513 = False

							}

							if True == ifres80513 {
								__e.Return(MakeNumber(1))
								return
							} else {
								tmp80506 := Call(__e, PrimNS1Value(symns2_1value), symshen_4complexity)

								tmp80507 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp80508 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp80509 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp80510 := Call(__e, tmp80509, sym_7, Nil)

								tmp80511 := Call(__e, tmp80508, V952, tmp80510)

								tmp80512 := Call(__e, tmp80507, symmode, tmp80511)

								__e.TailApply(tmp80506, tmp80512)
								return

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp80896 := Call(__e, PrimNS1Value(symns2_1set), symshen_4complexity, tmp80499)

	_ = tmp80896

	tmp80897 := MakeNative(func(__e *ControlFlow) {
		V954 := __e.Get(1)
		_ = V954
		tmp80910 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp80911 := Call(__e, tmp80910, Nil, V954)

		if True == tmp80911 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp80908 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp80909 := Call(__e, tmp80908, V954)

			if True == tmp80909 {
				tmp80901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4safe_1multiply)

				tmp80902 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp80903 := Call(__e, tmp80902, V954)

				tmp80904 := Call(__e, PrimNS1Value(symns2_1value), symshen_4safe_1product)

				tmp80905 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp80906 := Call(__e, tmp80905, V954)

				tmp80907 := Call(__e, tmp80904, tmp80906)

				__e.TailApply(tmp80901, tmp80903, tmp80907)
				return

			} else {
				tmp80900 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp80900, symshen_4safe_1product)
				return

			}

		}

	}, 1)

	tmp80912 := Call(__e, PrimNS1Value(symns2_1set), symshen_4safe_1product, tmp80897)

	_ = tmp80912

	tmp80913 := MakeNative(func(__e *ControlFlow) {
		V956 := __e.Get(1)
		_ = V956
		tmp81080 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81081 := Call(__e, tmp81080, V956)

		var ifres81050 Obj

		if True == tmp81081 {
			tmp81076 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81077 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81078 := Call(__e, tmp81077, V956)

			tmp81079 := Call(__e, tmp81076, symis, tmp81078)

			var ifres81052 Obj

			if True == tmp81079 {
				tmp81072 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81073 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81074 := Call(__e, tmp81073, V956)

				tmp81075 := Call(__e, tmp81072, tmp81074)

				var ifres81054 Obj

				if True == tmp81075 {
					tmp81066 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81067 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81069 := Call(__e, tmp81068, V956)

					tmp81070 := Call(__e, tmp81067, tmp81069)

					tmp81071 := Call(__e, tmp81066, tmp81070)

					var ifres81056 Obj

					if True == tmp81071 {
						tmp81058 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp81059 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81060 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81061 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81062 := Call(__e, tmp81061, V956)

						tmp81063 := Call(__e, tmp81060, tmp81062)

						tmp81064 := Call(__e, tmp81059, tmp81063)

						tmp81065 := Call(__e, tmp81058, Nil, tmp81064)

						var ifres81057 Obj

						if True == tmp81065 {
							ifres81057 = True

						} else {
							ifres81057 = False

						}

						ifres81056 = ifres81057

					} else {
						ifres81056 = False

					}

					var ifres81055 Obj

					if True == ifres81056 {
						ifres81055 = True

					} else {
						ifres81055 = False

					}

					ifres81054 = ifres81055

				} else {
					ifres81054 = False

				}

				var ifres81053 Obj

				if True == ifres81054 {
					ifres81053 = True

				} else {
					ifres81053 = False

				}

				ifres81052 = ifres81053

			} else {
				ifres81052 = False

			}

			var ifres81051 Obj

			if True == ifres81052 {
				ifres81051 = True

			} else {
				ifres81051 = False

			}

			ifres81050 = ifres81051

		} else {
			ifres81050 = False

		}

		if True == ifres81050 {
			tmp81033 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81034 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81035 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81036 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81037 := Call(__e, tmp81036, V956)

			tmp81038 := Call(__e, tmp81035, tmp81037)

			tmp81039 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81040 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

			tmp81041 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81043 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81044 := Call(__e, tmp81043, V956)

			tmp81045 := Call(__e, tmp81042, tmp81044)

			tmp81046 := Call(__e, tmp81041, tmp81045)

			tmp81047 := Call(__e, tmp81040, tmp81046, symProcessN)

			tmp81048 := Call(__e, tmp81039, tmp81047, Nil)

			tmp81049 := Call(__e, tmp81034, tmp81038, tmp81048)

			__e.TailApply(tmp81033, symbind, tmp81049)
			return

		} else {
			tmp81031 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81032 := Call(__e, tmp81031, V956)

			var ifres81011 Obj

			if True == tmp81032 {
				tmp81027 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp81028 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81029 := Call(__e, tmp81028, V956)

				tmp81030 := Call(__e, tmp81027, symwhen, tmp81029)

				var ifres81013 Obj

				if True == tmp81030 {
					tmp81023 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81024 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81025 := Call(__e, tmp81024, V956)

					tmp81026 := Call(__e, tmp81023, tmp81025)

					var ifres81015 Obj

					if True == tmp81026 {
						tmp81017 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp81018 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81019 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81020 := Call(__e, tmp81019, V956)

						tmp81021 := Call(__e, tmp81018, tmp81020)

						tmp81022 := Call(__e, tmp81017, Nil, tmp81021)

						var ifres81016 Obj

						if True == tmp81022 {
							ifres81016 = True

						} else {
							ifres81016 = False

						}

						ifres81015 = ifres81016

					} else {
						ifres81015 = False

					}

					var ifres81014 Obj

					if True == ifres81015 {
						ifres81014 = True

					} else {
						ifres81014 = False

					}

					ifres81013 = ifres81014

				} else {
					ifres81013 = False

				}

				var ifres81012 Obj

				if True == ifres81013 {
					ifres81012 = True

				} else {
					ifres81012 = False

				}

				ifres81011 = ifres81012

			} else {
				ifres81011 = False

			}

			if True == ifres81011 {
				tmp81002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81003 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

				tmp81005 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81006 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81007 := Call(__e, tmp81006, V956)

				tmp81008 := Call(__e, tmp81005, tmp81007)

				tmp81009 := Call(__e, tmp81004, tmp81008, symProcessN)

				tmp81010 := Call(__e, tmp81003, tmp81009, Nil)

				__e.TailApply(tmp81002, symfwhen, tmp81010)
				return

			} else {
				tmp81000 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81001 := Call(__e, tmp81000, V956)

				var ifres80970 Obj

				if True == tmp81001 {
					tmp80996 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp80997 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80998 := Call(__e, tmp80997, V956)

					tmp80999 := Call(__e, tmp80996, symbind, tmp80998)

					var ifres80972 Obj

					if True == tmp80999 {
						tmp80992 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp80993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80994 := Call(__e, tmp80993, V956)

						tmp80995 := Call(__e, tmp80992, tmp80994)

						var ifres80974 Obj

						if True == tmp80995 {
							tmp80986 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80988 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80989 := Call(__e, tmp80988, V956)

							tmp80990 := Call(__e, tmp80987, tmp80989)

							tmp80991 := Call(__e, tmp80986, tmp80990)

							var ifres80976 Obj

							if True == tmp80991 {
								tmp80978 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp80979 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80980 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80981 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80982 := Call(__e, tmp80981, V956)

								tmp80983 := Call(__e, tmp80980, tmp80982)

								tmp80984 := Call(__e, tmp80979, tmp80983)

								tmp80985 := Call(__e, tmp80978, Nil, tmp80984)

								var ifres80977 Obj

								if True == tmp80985 {
									ifres80977 = True

								} else {
									ifres80977 = False

								}

								ifres80976 = ifres80977

							} else {
								ifres80976 = False

							}

							var ifres80975 Obj

							if True == ifres80976 {
								ifres80975 = True

							} else {
								ifres80975 = False

							}

							ifres80974 = ifres80975

						} else {
							ifres80974 = False

						}

						var ifres80973 Obj

						if True == ifres80974 {
							ifres80973 = True

						} else {
							ifres80973 = False

						}

						ifres80972 = ifres80973

					} else {
						ifres80972 = False

					}

					var ifres80971 Obj

					if True == ifres80972 {
						ifres80971 = True

					} else {
						ifres80971 = False

					}

					ifres80970 = ifres80971

				} else {
					ifres80970 = False

				}

				if True == ifres80970 {
					tmp80953 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80954 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80955 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80956 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80957 := Call(__e, tmp80956, V956)

					tmp80958 := Call(__e, tmp80955, tmp80957)

					tmp80959 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp80960 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

					tmp80961 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp80962 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80963 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp80964 := Call(__e, tmp80963, V956)

					tmp80965 := Call(__e, tmp80962, tmp80964)

					tmp80966 := Call(__e, tmp80961, tmp80965)

					tmp80967 := Call(__e, tmp80960, tmp80966, symProcessN)

					tmp80968 := Call(__e, tmp80959, tmp80967, Nil)

					tmp80969 := Call(__e, tmp80954, tmp80958, tmp80968)

					__e.TailApply(tmp80953, symbind, tmp80969)
					return

				} else {
					tmp80951 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp80952 := Call(__e, tmp80951, V956)

					var ifres80931 Obj

					if True == tmp80952 {
						tmp80947 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp80948 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80949 := Call(__e, tmp80948, V956)

						tmp80950 := Call(__e, tmp80947, symfwhen, tmp80949)

						var ifres80933 Obj

						if True == tmp80950 {
							tmp80943 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp80944 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp80945 := Call(__e, tmp80944, V956)

							tmp80946 := Call(__e, tmp80943, tmp80945)

							var ifres80935 Obj

							if True == tmp80946 {
								tmp80937 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp80938 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp80940 := Call(__e, tmp80939, V956)

								tmp80941 := Call(__e, tmp80938, tmp80940)

								tmp80942 := Call(__e, tmp80937, Nil, tmp80941)

								var ifres80936 Obj

								if True == tmp80942 {
									ifres80936 = True

								} else {
									ifres80936 = False

								}

								ifres80935 = ifres80936

							} else {
								ifres80935 = False

							}

							var ifres80934 Obj

							if True == ifres80935 {
								ifres80934 = True

							} else {
								ifres80934 = False

							}

							ifres80933 = ifres80934

						} else {
							ifres80933 = False

						}

						var ifres80932 Obj

						if True == ifres80933 {
							ifres80932 = True

						} else {
							ifres80932 = False

						}

						ifres80931 = ifres80932

					} else {
						ifres80931 = False

					}

					if True == ifres80931 {
						tmp80922 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp80923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp80924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

						tmp80925 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp80926 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp80927 := Call(__e, tmp80926, V956)

						tmp80928 := Call(__e, tmp80925, tmp80927)

						tmp80929 := Call(__e, tmp80924, tmp80928, symProcessN)

						tmp80930 := Call(__e, tmp80923, tmp80929, Nil)

						__e.TailApply(tmp80922, symfwhen, tmp80930)
						return

					} else {
						tmp80920 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp80921 := Call(__e, tmp80920, V956)

						if True == tmp80921 {
							__e.Return(V956)
							return
						} else {
							tmp80919 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(tmp80919, symshen_4s_1prolog__literal)
							return

						}

					}

				}

			}

		}

	}, 1)

	tmp81082 := Call(__e, PrimNS1Value(symns2_1set), symshen_4s_1prolog__literal, tmp80913)

	_ = tmp81082

	tmp81083 := MakeNative(func(__e *ControlFlow) {
		V963 := __e.Get(1)
		_ = V963
		V964 := __e.Get(2)
		_ = V964
		tmp81226 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

		tmp81227 := Call(__e, tmp81226, V963)

		if True == tmp81227 {
			tmp81221 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81222 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81223 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81224 := Call(__e, tmp81223, V964, Nil)

			tmp81225 := Call(__e, tmp81222, V963, tmp81224)

			__e.TailApply(tmp81221, symshen_4deref, tmp81225)
			return

		} else {
			tmp81219 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81220 := Call(__e, tmp81219, V963)

			var ifres81189 Obj

			if True == tmp81220 {
				tmp81215 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp81216 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81217 := Call(__e, tmp81216, V963)

				tmp81218 := Call(__e, tmp81215, symlambda, tmp81217)

				var ifres81191 Obj

				if True == tmp81218 {
					tmp81211 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81212 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81213 := Call(__e, tmp81212, V963)

					tmp81214 := Call(__e, tmp81211, tmp81213)

					var ifres81193 Obj

					if True == tmp81214 {
						tmp81205 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp81206 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81208 := Call(__e, tmp81207, V963)

						tmp81209 := Call(__e, tmp81206, tmp81208)

						tmp81210 := Call(__e, tmp81205, tmp81209)

						var ifres81195 Obj

						if True == tmp81210 {
							tmp81197 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp81198 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81199 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81200 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81201 := Call(__e, tmp81200, V963)

							tmp81202 := Call(__e, tmp81199, tmp81201)

							tmp81203 := Call(__e, tmp81198, tmp81202)

							tmp81204 := Call(__e, tmp81197, Nil, tmp81203)

							var ifres81196 Obj

							if True == tmp81204 {
								ifres81196 = True

							} else {
								ifres81196 = False

							}

							ifres81195 = ifres81196

						} else {
							ifres81195 = False

						}

						var ifres81194 Obj

						if True == ifres81195 {
							ifres81194 = True

						} else {
							ifres81194 = False

						}

						ifres81193 = ifres81194

					} else {
						ifres81193 = False

					}

					var ifres81192 Obj

					if True == ifres81193 {
						ifres81192 = True

					} else {
						ifres81192 = False

					}

					ifres81191 = ifres81192

				} else {
					ifres81191 = False

				}

				var ifres81190 Obj

				if True == ifres81191 {
					ifres81190 = True

				} else {
					ifres81190 = False

				}

				ifres81189 = ifres81190

			} else {
				ifres81189 = False

			}

			if True == ifres81189 {
				tmp81172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81174 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81175 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81176 := Call(__e, tmp81175, V963)

				tmp81177 := Call(__e, tmp81174, tmp81176)

				tmp81178 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

				tmp81180 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81181 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81182 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81183 := Call(__e, tmp81182, V963)

				tmp81184 := Call(__e, tmp81181, tmp81183)

				tmp81185 := Call(__e, tmp81180, tmp81184)

				tmp81186 := Call(__e, tmp81179, tmp81185, V964)

				tmp81187 := Call(__e, tmp81178, tmp81186, Nil)

				tmp81188 := Call(__e, tmp81173, tmp81177, tmp81187)

				__e.TailApply(tmp81172, symlambda, tmp81188)
				return

			} else {
				tmp81170 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81171 := Call(__e, tmp81170, V963)

				var ifres81128 Obj

				if True == tmp81171 {
					tmp81166 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp81167 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81168 := Call(__e, tmp81167, V963)

					tmp81169 := Call(__e, tmp81166, symlet, tmp81168)

					var ifres81130 Obj

					if True == tmp81169 {
						tmp81162 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp81163 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81164 := Call(__e, tmp81163, V963)

						tmp81165 := Call(__e, tmp81162, tmp81164)

						var ifres81132 Obj

						if True == tmp81165 {
							tmp81156 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp81157 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81158 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81159 := Call(__e, tmp81158, V963)

							tmp81160 := Call(__e, tmp81157, tmp81159)

							tmp81161 := Call(__e, tmp81156, tmp81160)

							var ifres81134 Obj

							if True == tmp81161 {
								tmp81148 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp81149 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81150 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81151 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81152 := Call(__e, tmp81151, V963)

								tmp81153 := Call(__e, tmp81150, tmp81152)

								tmp81154 := Call(__e, tmp81149, tmp81153)

								tmp81155 := Call(__e, tmp81148, tmp81154)

								var ifres81136 Obj

								if True == tmp81155 {
									tmp81138 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp81139 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81140 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81141 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81143 := Call(__e, tmp81142, V963)

									tmp81144 := Call(__e, tmp81141, tmp81143)

									tmp81145 := Call(__e, tmp81140, tmp81144)

									tmp81146 := Call(__e, tmp81139, tmp81145)

									tmp81147 := Call(__e, tmp81138, Nil, tmp81146)

									var ifres81137 Obj

									if True == tmp81147 {
										ifres81137 = True

									} else {
										ifres81137 = False

									}

									ifres81136 = ifres81137

								} else {
									ifres81136 = False

								}

								var ifres81135 Obj

								if True == ifres81136 {
									ifres81135 = True

								} else {
									ifres81135 = False

								}

								ifres81134 = ifres81135

							} else {
								ifres81134 = False

							}

							var ifres81133 Obj

							if True == ifres81134 {
								ifres81133 = True

							} else {
								ifres81133 = False

							}

							ifres81132 = ifres81133

						} else {
							ifres81132 = False

						}

						var ifres81131 Obj

						if True == ifres81132 {
							ifres81131 = True

						} else {
							ifres81131 = False

						}

						ifres81130 = ifres81131

					} else {
						ifres81130 = False

					}

					var ifres81129 Obj

					if True == ifres81130 {
						ifres81129 = True

					} else {
						ifres81129 = False

					}

					ifres81128 = ifres81129

				} else {
					ifres81128 = False

				}

				if True == ifres81128 {
					tmp81099 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81100 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81101 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81103 := Call(__e, tmp81102, V963)

					tmp81104 := Call(__e, tmp81101, tmp81103)

					tmp81105 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

					tmp81107 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81108 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81109 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81110 := Call(__e, tmp81109, V963)

					tmp81111 := Call(__e, tmp81108, tmp81110)

					tmp81112 := Call(__e, tmp81107, tmp81111)

					tmp81113 := Call(__e, tmp81106, tmp81112, V964)

					tmp81114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

					tmp81116 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81117 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81118 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81119 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81120 := Call(__e, tmp81119, V963)

					tmp81121 := Call(__e, tmp81118, tmp81120)

					tmp81122 := Call(__e, tmp81117, tmp81121)

					tmp81123 := Call(__e, tmp81116, tmp81122)

					tmp81124 := Call(__e, tmp81115, tmp81123, V964)

					tmp81125 := Call(__e, tmp81114, tmp81124, Nil)

					tmp81126 := Call(__e, tmp81105, tmp81113, tmp81125)

					tmp81127 := Call(__e, tmp81100, tmp81104, tmp81126)

					__e.TailApply(tmp81099, symlet, tmp81127)
					return

				} else {
					tmp81097 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81098 := Call(__e, tmp81097, V963)

					if True == tmp81098 {
						tmp81088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp81089 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

						tmp81090 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp81091 := Call(__e, tmp81090, V963)

						tmp81092 := Call(__e, tmp81089, tmp81091, V964)

						tmp81093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

						tmp81094 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81095 := Call(__e, tmp81094, V963)

						tmp81096 := Call(__e, tmp81093, tmp81095, V964)

						__e.TailApply(tmp81088, tmp81092, tmp81096)
						return

					} else {
						__e.Return(V963)
						return
					}

				}

			}

		}

	}, 2)

	tmp81228 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1deref, tmp81083)

	_ = tmp81228

	tmp81229 := MakeNative(func(__e *ControlFlow) {
		V971 := __e.Get(1)
		_ = V971
		V972 := __e.Get(2)
		_ = V972
		tmp81372 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

		tmp81373 := Call(__e, tmp81372, V971)

		if True == tmp81373 {
			tmp81367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81368 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81370 := Call(__e, tmp81369, V972, Nil)

			tmp81371 := Call(__e, tmp81368, V971, tmp81370)

			__e.TailApply(tmp81367, symshen_4lazyderef, tmp81371)
			return

		} else {
			tmp81365 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81366 := Call(__e, tmp81365, V971)

			var ifres81335 Obj

			if True == tmp81366 {
				tmp81361 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp81362 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81363 := Call(__e, tmp81362, V971)

				tmp81364 := Call(__e, tmp81361, symlambda, tmp81363)

				var ifres81337 Obj

				if True == tmp81364 {
					tmp81357 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81358 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81359 := Call(__e, tmp81358, V971)

					tmp81360 := Call(__e, tmp81357, tmp81359)

					var ifres81339 Obj

					if True == tmp81360 {
						tmp81351 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp81352 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81353 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81354 := Call(__e, tmp81353, V971)

						tmp81355 := Call(__e, tmp81352, tmp81354)

						tmp81356 := Call(__e, tmp81351, tmp81355)

						var ifres81341 Obj

						if True == tmp81356 {
							tmp81343 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp81344 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81345 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81347 := Call(__e, tmp81346, V971)

							tmp81348 := Call(__e, tmp81345, tmp81347)

							tmp81349 := Call(__e, tmp81344, tmp81348)

							tmp81350 := Call(__e, tmp81343, Nil, tmp81349)

							var ifres81342 Obj

							if True == tmp81350 {
								ifres81342 = True

							} else {
								ifres81342 = False

							}

							ifres81341 = ifres81342

						} else {
							ifres81341 = False

						}

						var ifres81340 Obj

						if True == ifres81341 {
							ifres81340 = True

						} else {
							ifres81340 = False

						}

						ifres81339 = ifres81340

					} else {
						ifres81339 = False

					}

					var ifres81338 Obj

					if True == ifres81339 {
						ifres81338 = True

					} else {
						ifres81338 = False

					}

					ifres81337 = ifres81338

				} else {
					ifres81337 = False

				}

				var ifres81336 Obj

				if True == ifres81337 {
					ifres81336 = True

				} else {
					ifres81336 = False

				}

				ifres81335 = ifres81336

			} else {
				ifres81335 = False

			}

			if True == ifres81335 {
				tmp81318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81320 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81321 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81322 := Call(__e, tmp81321, V971)

				tmp81323 := Call(__e, tmp81320, tmp81322)

				tmp81324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp81325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

				tmp81326 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81327 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81328 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81329 := Call(__e, tmp81328, V971)

				tmp81330 := Call(__e, tmp81327, tmp81329)

				tmp81331 := Call(__e, tmp81326, tmp81330)

				tmp81332 := Call(__e, tmp81325, tmp81331, V972)

				tmp81333 := Call(__e, tmp81324, tmp81332, Nil)

				tmp81334 := Call(__e, tmp81319, tmp81323, tmp81333)

				__e.TailApply(tmp81318, symlambda, tmp81334)
				return

			} else {
				tmp81316 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81317 := Call(__e, tmp81316, V971)

				var ifres81274 Obj

				if True == tmp81317 {
					tmp81312 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp81313 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81314 := Call(__e, tmp81313, V971)

					tmp81315 := Call(__e, tmp81312, symlet, tmp81314)

					var ifres81276 Obj

					if True == tmp81315 {
						tmp81308 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp81309 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81310 := Call(__e, tmp81309, V971)

						tmp81311 := Call(__e, tmp81308, tmp81310)

						var ifres81278 Obj

						if True == tmp81311 {
							tmp81302 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp81303 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81304 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81305 := Call(__e, tmp81304, V971)

							tmp81306 := Call(__e, tmp81303, tmp81305)

							tmp81307 := Call(__e, tmp81302, tmp81306)

							var ifres81280 Obj

							if True == tmp81307 {
								tmp81294 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp81295 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81296 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81297 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81298 := Call(__e, tmp81297, V971)

								tmp81299 := Call(__e, tmp81296, tmp81298)

								tmp81300 := Call(__e, tmp81295, tmp81299)

								tmp81301 := Call(__e, tmp81294, tmp81300)

								var ifres81282 Obj

								if True == tmp81301 {
									tmp81284 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp81285 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81286 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81287 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81289 := Call(__e, tmp81288, V971)

									tmp81290 := Call(__e, tmp81287, tmp81289)

									tmp81291 := Call(__e, tmp81286, tmp81290)

									tmp81292 := Call(__e, tmp81285, tmp81291)

									tmp81293 := Call(__e, tmp81284, Nil, tmp81292)

									var ifres81283 Obj

									if True == tmp81293 {
										ifres81283 = True

									} else {
										ifres81283 = False

									}

									ifres81282 = ifres81283

								} else {
									ifres81282 = False

								}

								var ifres81281 Obj

								if True == ifres81282 {
									ifres81281 = True

								} else {
									ifres81281 = False

								}

								ifres81280 = ifres81281

							} else {
								ifres81280 = False

							}

							var ifres81279 Obj

							if True == ifres81280 {
								ifres81279 = True

							} else {
								ifres81279 = False

							}

							ifres81278 = ifres81279

						} else {
							ifres81278 = False

						}

						var ifres81277 Obj

						if True == ifres81278 {
							ifres81277 = True

						} else {
							ifres81277 = False

						}

						ifres81276 = ifres81277

					} else {
						ifres81276 = False

					}

					var ifres81275 Obj

					if True == ifres81276 {
						ifres81275 = True

					} else {
						ifres81275 = False

					}

					ifres81274 = ifres81275

				} else {
					ifres81274 = False

				}

				if True == ifres81274 {
					tmp81245 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81246 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81247 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81248 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81249 := Call(__e, tmp81248, V971)

					tmp81250 := Call(__e, tmp81247, tmp81249)

					tmp81251 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

					tmp81253 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81254 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81255 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81256 := Call(__e, tmp81255, V971)

					tmp81257 := Call(__e, tmp81254, tmp81256)

					tmp81258 := Call(__e, tmp81253, tmp81257)

					tmp81259 := Call(__e, tmp81252, tmp81258, V972)

					tmp81260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81261 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

					tmp81262 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81263 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81264 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81265 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81266 := Call(__e, tmp81265, V971)

					tmp81267 := Call(__e, tmp81264, tmp81266)

					tmp81268 := Call(__e, tmp81263, tmp81267)

					tmp81269 := Call(__e, tmp81262, tmp81268)

					tmp81270 := Call(__e, tmp81261, tmp81269, V972)

					tmp81271 := Call(__e, tmp81260, tmp81270, Nil)

					tmp81272 := Call(__e, tmp81251, tmp81259, tmp81271)

					tmp81273 := Call(__e, tmp81246, tmp81250, tmp81272)

					__e.TailApply(tmp81245, symlet, tmp81273)
					return

				} else {
					tmp81243 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81244 := Call(__e, tmp81243, V971)

					if True == tmp81244 {
						tmp81234 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp81235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

						tmp81236 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp81237 := Call(__e, tmp81236, V971)

						tmp81238 := Call(__e, tmp81235, tmp81237, V972)

						tmp81239 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

						tmp81240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81241 := Call(__e, tmp81240, V971)

						tmp81242 := Call(__e, tmp81239, tmp81241, V972)

						__e.TailApply(tmp81234, tmp81238, tmp81242)
						return

					} else {
						__e.Return(V971)
						return
					}

				}

			}

		}

	}, 2)

	tmp81374 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1lazyderef, tmp81229)

	_ = tmp81374

	tmp81375 := MakeNative(func(__e *ControlFlow) {
		V974 := __e.Get(1)
		_ = V974
		tmp81394 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp81395 := Call(__e, tmp81394, Nil, V974)

		if True == tmp81395 {
			__e.Return(Nil)
			return
		} else {
			tmp81392 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81393 := Call(__e, tmp81392, V974)

			if True == tmp81393 {
				tmp81379 := MakeNative(func(__e *ControlFlow) {
					Group := __e.Get(1)
					_ = Group
					tmp81380 := MakeNative(func(__e *ControlFlow) {
						Rest := __e.Get(1)
						_ = Rest
						tmp81381 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp81382 := Call(__e, PrimNS1Value(symns2_1value), symshen_4group__clauses)

						tmp81383 := Call(__e, tmp81382, Rest)

						__e.TailApply(tmp81381, Group, tmp81383)
						return

					}, 1)

					tmp81384 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

					tmp81385 := Call(__e, tmp81384, V974, Group)

					__e.TailApply(tmp81380, tmp81385)
					return

				}, 1)

				tmp81386 := Call(__e, PrimNS1Value(symns2_1value), symshen_4collect)

				tmp81387 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp81388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4same__predicate_2)

					tmp81389 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81390 := Call(__e, tmp81389, V974)

					__e.TailApply(tmp81388, tmp81390, X)
					return

				}, 1)

				tmp81391 := Call(__e, tmp81386, tmp81387, V974)

				__e.TailApply(tmp81379, tmp81391)
				return

			} else {
				tmp81378 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp81378, symshen_4group__clauses)
				return

			}

		}

	}, 1)

	tmp81396 := Call(__e, PrimNS1Value(symns2_1set), symshen_4group__clauses, tmp81375)

	_ = tmp81396

	tmp81397 := MakeNative(func(__e *ControlFlow) {
		V979 := __e.Get(1)
		_ = V979
		V980 := __e.Get(2)
		_ = V980
		tmp81417 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp81418 := Call(__e, tmp81417, Nil, V980)

		if True == tmp81418 {
			__e.Return(Nil)
			return
		} else {
			tmp81415 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81416 := Call(__e, tmp81415, V980)

			if True == tmp81416 {
				tmp81412 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81413 := Call(__e, tmp81412, V980)

				tmp81414 := Call(__e, V979, tmp81413)

				if True == tmp81414 {
					tmp81405 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81406 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81407 := Call(__e, tmp81406, V980)

					tmp81408 := Call(__e, PrimNS1Value(symns2_1value), symshen_4collect)

					tmp81409 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81410 := Call(__e, tmp81409, V980)

					tmp81411 := Call(__e, tmp81408, V979, tmp81410)

					__e.TailApply(tmp81405, tmp81407, tmp81411)
					return

				} else {
					tmp81402 := Call(__e, PrimNS1Value(symns2_1value), symshen_4collect)

					tmp81403 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81404 := Call(__e, tmp81403, V980)

					__e.TailApply(tmp81402, V979, tmp81404)
					return

				}

			} else {
				tmp81400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp81400, symshen_4collect)
				return

			}

		}

	}, 2)

	tmp81419 := Call(__e, PrimNS1Value(symns2_1set), symshen_4collect, tmp81397)

	_ = tmp81419

	tmp81420 := MakeNative(func(__e *ControlFlow) {
		V999 := __e.Get(1)
		_ = V999
		V1000 := __e.Get(2)
		_ = V1000
		tmp81448 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81449 := Call(__e, tmp81448, V999)

		var ifres81432 Obj

		if True == tmp81449 {
			tmp81444 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81445 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81446 := Call(__e, tmp81445, V999)

			tmp81447 := Call(__e, tmp81444, tmp81446)

			var ifres81434 Obj

			if True == tmp81447 {
				tmp81442 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81443 := Call(__e, tmp81442, V1000)

				var ifres81436 Obj

				if True == tmp81443 {
					tmp81438 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81439 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81440 := Call(__e, tmp81439, V1000)

					tmp81441 := Call(__e, tmp81438, tmp81440)

					var ifres81437 Obj

					if True == tmp81441 {
						ifres81437 = True

					} else {
						ifres81437 = False

					}

					ifres81436 = ifres81437

				} else {
					ifres81436 = False

				}

				var ifres81435 Obj

				if True == ifres81436 {
					ifres81435 = True

				} else {
					ifres81435 = False

				}

				ifres81434 = ifres81435

			} else {
				ifres81434 = False

			}

			var ifres81433 Obj

			if True == ifres81434 {
				ifres81433 = True

			} else {
				ifres81433 = False

			}

			ifres81432 = ifres81433

		} else {
			ifres81432 = False

		}

		if True == ifres81432 {
			tmp81423 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81424 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81425 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81426 := Call(__e, tmp81425, V999)

			tmp81427 := Call(__e, tmp81424, tmp81426)

			tmp81428 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81429 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81430 := Call(__e, tmp81429, V1000)

			tmp81431 := Call(__e, tmp81428, tmp81430)

			__e.TailApply(tmp81423, tmp81427, tmp81431)
			return

		} else {
			tmp81422 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp81422, symshen_4same__predicate_2)
			return

		}

	}, 2)

	tmp81450 := Call(__e, PrimNS1Value(symns2_1set), symshen_4same__predicate_2, tmp81420)

	_ = tmp81450

	tmp81451 := MakeNative(func(__e *ControlFlow) {
		V1002 := __e.Get(1)
		_ = V1002
		tmp81452 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			tmp81453 := MakeNative(func(__e *ControlFlow) {
				Shen := __e.Get(1)
				_ = Shen
				__e.Return(Shen)
				return
			}, 1)

			tmp81454 := Call(__e, PrimNS1Value(symns2_1value), symshen_4clauses_1to_1shen)

			tmp81455 := Call(__e, tmp81454, F, V1002)

			__e.TailApply(tmp81453, tmp81455)
			return

		}, 1)

		tmp81456 := Call(__e, PrimNS1Value(symns2_1value), symshen_4procedure__name)

		tmp81457 := Call(__e, tmp81456, V1002)

		__e.TailApply(tmp81452, tmp81457)
		return

	}, 1)

	tmp81458 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__prolog__procedure, tmp81451)

	_ = tmp81458

	tmp81459 := MakeNative(func(__e *ControlFlow) {
		V1016 := __e.Get(1)
		_ = V1016
		tmp81481 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81482 := Call(__e, tmp81481, V1016)

		var ifres81467 Obj

		if True == tmp81482 {
			tmp81477 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81478 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81479 := Call(__e, tmp81478, V1016)

			tmp81480 := Call(__e, tmp81477, tmp81479)

			var ifres81469 Obj

			if True == tmp81480 {
				tmp81471 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81472 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81473 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81474 := Call(__e, tmp81473, V1016)

				tmp81475 := Call(__e, tmp81472, tmp81474)

				tmp81476 := Call(__e, tmp81471, tmp81475)

				var ifres81470 Obj

				if True == tmp81476 {
					ifres81470 = True

				} else {
					ifres81470 = False

				}

				ifres81469 = ifres81470

			} else {
				ifres81469 = False

			}

			var ifres81468 Obj

			if True == ifres81469 {
				ifres81468 = True

			} else {
				ifres81468 = False

			}

			ifres81467 = ifres81468

		} else {
			ifres81467 = False

		}

		if True == ifres81467 {
			tmp81462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81463 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81464 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81465 := Call(__e, tmp81464, V1016)

			tmp81466 := Call(__e, tmp81463, tmp81465)

			__e.TailApply(tmp81462, tmp81466)
			return

		} else {
			tmp81461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp81461, symshen_4procedure__name)
			return

		}

	}, 1)

	tmp81483 := Call(__e, PrimNS1Value(symns2_1set), symshen_4procedure__name, tmp81459)

	_ = tmp81483

	tmp81484 := MakeNative(func(__e *ControlFlow) {
		V1019 := __e.Get(1)
		_ = V1019
		V1020 := __e.Get(2)
		_ = V1020
		tmp81485 := MakeNative(func(__e *ControlFlow) {
			Linear := __e.Get(1)
			_ = Linear
			tmp81486 := MakeNative(func(__e *ControlFlow) {
				Arity := __e.Get(1)
				_ = Arity
				tmp81487 := MakeNative(func(__e *ControlFlow) {
					Parameters := __e.Get(1)
					_ = Parameters
					tmp81488 := MakeNative(func(__e *ControlFlow) {
						AUM__instructions := __e.Get(1)
						_ = AUM__instructions
						tmp81489 := MakeNative(func(__e *ControlFlow) {
							Code := __e.Get(1)
							_ = Code
							tmp81490 := MakeNative(func(__e *ControlFlow) {
								ShenDef := __e.Get(1)
								_ = ShenDef
								__e.Return(ShenDef)
								return
							}, 1)

							tmp81491 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp81492 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp81493 := Call(__e, PrimNS1Value(symns2_1value), symappend)

							tmp81494 := Call(__e, PrimNS1Value(symns2_1value), symappend)

							tmp81495 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp81496 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp81497 := Call(__e, tmp81496, symContinuation, Nil)

							tmp81498 := Call(__e, tmp81495, symProcessN, tmp81497)

							tmp81499 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp81500 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp81501 := Call(__e, tmp81500, Code, Nil)

							tmp81502 := Call(__e, tmp81499, sym_1_6, tmp81501)

							tmp81503 := Call(__e, tmp81494, tmp81498, tmp81502)

							tmp81504 := Call(__e, tmp81493, Parameters, tmp81503)

							tmp81505 := Call(__e, tmp81492, V1019, tmp81504)

							tmp81506 := Call(__e, tmp81491, symdefine, tmp81505)

							__e.TailApply(tmp81490, tmp81506)
							return

						}, 1)

						tmp81507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4catch_1cut)

						tmp81508 := Call(__e, PrimNS1Value(symns2_1value), symshen_4nest_1disjunct)

						tmp81509 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						tmp81510 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							tmp81511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

							__e.TailApply(tmp81511, X)
							return

						}, 1)

						tmp81512 := Call(__e, tmp81509, tmp81510, AUM__instructions)

						tmp81513 := Call(__e, tmp81508, tmp81512)

						tmp81514 := Call(__e, tmp81507, tmp81513)

						__e.TailApply(tmp81489, tmp81514)
						return

					}, 1)

					tmp81515 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					tmp81516 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp81517 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum)

						__e.TailApply(tmp81517, X, Parameters)
						return

					}, 1)

					tmp81518 := Call(__e, tmp81515, tmp81516, Linear)

					__e.TailApply(tmp81488, tmp81518)
					return

				}, 1)

				tmp81519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4parameters)

				tmp81520 := Call(__e, tmp81519, Arity)

				__e.TailApply(tmp81487, tmp81520)
				return

			}, 1)

			tmp81521 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1aritycheck)

			tmp81522 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp81523 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp81524 := Call(__e, PrimNS1Value(symns2_1value), symhead)

				__e.TailApply(tmp81524, X)
				return

			}, 1)

			tmp81525 := Call(__e, tmp81522, tmp81523, V1020)

			tmp81526 := Call(__e, tmp81521, V1019, tmp81525)

			__e.TailApply(tmp81486, tmp81526)
			return

		}, 1)

		tmp81527 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp81528 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp81529 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise_1clause)

			__e.TailApply(tmp81529, X)
			return

		}, 1)

		tmp81530 := Call(__e, tmp81527, tmp81528, V1020)

		__e.TailApply(tmp81485, tmp81530)
		return

	}, 2)

	tmp81531 := Call(__e, PrimNS1Value(symns2_1set), symshen_4clauses_1to_1shen, tmp81484)

	_ = tmp81531

	tmp81532 := MakeNative(func(__e *ControlFlow) {
		V1022 := __e.Get(1)
		_ = V1022
		tmp81549 := Call(__e, PrimNS1Value(symns2_1value), symnot)

		tmp81550 := Call(__e, PrimNS1Value(symns2_1value), symshen_4occurs_2)

		tmp81551 := Call(__e, tmp81550, symcut, V1022)

		tmp81552 := Call(__e, tmp81549, tmp81551)

		if True == tmp81552 {
			__e.Return(V1022)
			return
		} else {
			tmp81534 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81535 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81536 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81537 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81538 := Call(__e, tmp81537, symshen_4catchpoint, Nil)

			tmp81539 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81541 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81542 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81543 := Call(__e, tmp81542, V1022, Nil)

			tmp81544 := Call(__e, tmp81541, symThrowcontrol, tmp81543)

			tmp81545 := Call(__e, tmp81540, symshen_4cutpoint, tmp81544)

			tmp81546 := Call(__e, tmp81539, tmp81545, Nil)

			tmp81547 := Call(__e, tmp81536, tmp81538, tmp81546)

			tmp81548 := Call(__e, tmp81535, symThrowcontrol, tmp81547)

			__e.TailApply(tmp81534, symlet, tmp81548)
			return

		}

	}, 1)

	tmp81553 := Call(__e, PrimNS1Value(symns2_1set), symshen_4catch_1cut, tmp81532)

	_ = tmp81553

	tmp81554 := MakeNative(func(__e *ControlFlow) {
		tmp81555 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81556 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp81557 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp81558 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp81559 := Call(__e, tmp81558, symshen_4_dcatch_d)

		tmp81560 := Call(__e, tmp81557, MakeNumber(1), tmp81559)

		tmp81561 := Call(__e, tmp81556, symshen_4_dcatch_d, tmp81560)

		__e.TailApply(tmp81555, symshen_4catchpoint_b, tmp81561)
		return

	}, 0)

	tmp81562 := Call(__e, PrimNS1Value(symns2_1set), symshen_4catchpoint, tmp81554)

	_ = tmp81562

	tmp81563 := MakeNative(func(__e *ControlFlow) {
		V1030 := __e.Get(1)
		_ = V1030
		V1031 := __e.Get(2)
		_ = V1031
		tmp81565 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp81566 := Call(__e, tmp81565, V1031, V1030)

		if True == tmp81566 {
			__e.Return(False)
			return
		} else {
			__e.Return(V1031)
			return
		}

	}, 2)

	tmp81567 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cutpoint, tmp81563)

	_ = tmp81567

	tmp81568 := MakeNative(func(__e *ControlFlow) {
		V1033 := __e.Get(1)
		_ = V1033
		tmp81588 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81589 := Call(__e, tmp81588, V1033)

		var ifres81582 Obj

		if True == tmp81589 {
			tmp81584 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81585 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81586 := Call(__e, tmp81585, V1033)

			tmp81587 := Call(__e, tmp81584, Nil, tmp81586)

			var ifres81583 Obj

			if True == tmp81587 {
				ifres81583 = True

			} else {
				ifres81583 = False

			}

			ifres81582 = ifres81583

		} else {
			ifres81582 = False

		}

		if True == ifres81582 {
			tmp81581 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(tmp81581, V1033)
			return

		} else {
			tmp81579 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81580 := Call(__e, tmp81579, V1033)

			if True == tmp81580 {
				tmp81572 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lisp_1or)

				tmp81573 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81574 := Call(__e, tmp81573, V1033)

				tmp81575 := Call(__e, PrimNS1Value(symns2_1value), symshen_4nest_1disjunct)

				tmp81576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81577 := Call(__e, tmp81576, V1033)

				tmp81578 := Call(__e, tmp81575, tmp81577)

				__e.TailApply(tmp81572, tmp81574, tmp81578)
				return

			} else {
				tmp81571 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp81571, symshen_4nest_1disjunct)
				return

			}

		}

	}, 1)

	tmp81590 := Call(__e, PrimNS1Value(symns2_1set), symshen_4nest_1disjunct, tmp81568)

	_ = tmp81590

	tmp81591 := MakeNative(func(__e *ControlFlow) {
		V1036 := __e.Get(1)
		_ = V1036
		V1037 := __e.Get(2)
		_ = V1037
		tmp81592 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81593 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81594 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81595 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81596 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81601 := Call(__e, tmp81600, False, Nil)

		tmp81602 := Call(__e, tmp81599, symCase, tmp81601)

		tmp81603 := Call(__e, tmp81598, sym_a, tmp81602)

		tmp81604 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81605 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp81606 := Call(__e, tmp81605, symCase, Nil)

		tmp81607 := Call(__e, tmp81604, V1037, tmp81606)

		tmp81608 := Call(__e, tmp81597, tmp81603, tmp81607)

		tmp81609 := Call(__e, tmp81596, symif, tmp81608)

		tmp81610 := Call(__e, tmp81595, tmp81609, Nil)

		tmp81611 := Call(__e, tmp81594, V1036, tmp81610)

		tmp81612 := Call(__e, tmp81593, symCase, tmp81611)

		__e.TailApply(tmp81592, symlet, tmp81612)
		return

	}, 2)

	tmp81613 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lisp_1or, tmp81591)

	_ = tmp81613

	tmp81614 := MakeNative(func(__e *ControlFlow) {
		V1042 := __e.Get(1)
		_ = V1042
		V1043 := __e.Get(2)
		_ = V1043
		tmp81660 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81661 := Call(__e, tmp81660, V1043)

		var ifres81654 Obj

		if True == tmp81661 {
			tmp81656 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81657 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81658 := Call(__e, tmp81657, V1043)

			tmp81659 := Call(__e, tmp81656, Nil, tmp81658)

			var ifres81655 Obj

			if True == tmp81659 {
				ifres81655 = True

			} else {
				ifres81655 = False

			}

			ifres81654 = ifres81655

		} else {
			ifres81654 = False

		}

		if True == ifres81654 {
			tmp81649 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			tmp81650 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			tmp81651 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81652 := Call(__e, tmp81651, V1043)

			tmp81653 := Call(__e, tmp81650, tmp81652)

			__e.TailApply(tmp81649, tmp81653, MakeNumber(1))
			return

		} else {
			tmp81647 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81648 := Call(__e, tmp81647, V1043)

			var ifres81641 Obj

			if True == tmp81648 {
				tmp81643 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81644 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81645 := Call(__e, tmp81644, V1043)

				tmp81646 := Call(__e, tmp81643, tmp81645)

				var ifres81642 Obj

				if True == tmp81646 {
					ifres81642 = True

				} else {
					ifres81642 = False

				}

				ifres81641 = ifres81642

			} else {
				ifres81641 = False

			}

			if True == ifres81641 {
				tmp81629 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp81630 := Call(__e, PrimNS1Value(symns2_1value), symlength)

				tmp81631 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81632 := Call(__e, tmp81631, V1043)

				tmp81633 := Call(__e, tmp81630, tmp81632)

				tmp81634 := Call(__e, PrimNS1Value(symns2_1value), symlength)

				tmp81635 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81636 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81637 := Call(__e, tmp81636, V1043)

				tmp81638 := Call(__e, tmp81635, tmp81637)

				tmp81639 := Call(__e, tmp81634, tmp81638)

				tmp81640 := Call(__e, tmp81629, tmp81633, tmp81639)

				if True == tmp81640 {
					tmp81626 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1aritycheck)

					tmp81627 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81628 := Call(__e, tmp81627, V1043)

					__e.TailApply(tmp81626, V1042, tmp81628)
					return

				} else {
					tmp81619 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					tmp81620 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp81621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp81622 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp81623 := Call(__e, tmp81622, V1042, Nil)

					tmp81624 := Call(__e, tmp81621, tmp81623, MakeString("\n"), symshen_4a)

					tmp81625 := Call(__e, tmp81620, MakeString("arity error in prolog procedure "), tmp81624)

					__e.TailApply(tmp81619, tmp81625)
					return

				}

			} else {
				tmp81617 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp81617, symshen_4prolog_1aritycheck)
				return

			}

		}

	}, 2)

	tmp81662 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1aritycheck, tmp81614)

	_ = tmp81662

	tmp81663 := MakeNative(func(__e *ControlFlow) {
		V1045 := __e.Get(1)
		_ = V1045
		tmp81710 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81711 := Call(__e, tmp81710, V1045)

		var ifres81678 Obj

		if True == tmp81711 {
			tmp81706 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81707 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81708 := Call(__e, tmp81707, V1045)

			tmp81709 := Call(__e, tmp81706, tmp81708)

			var ifres81680 Obj

			if True == tmp81709 {
				tmp81700 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp81701 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp81702 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81703 := Call(__e, tmp81702, V1045)

				tmp81704 := Call(__e, tmp81701, tmp81703)

				tmp81705 := Call(__e, tmp81700, sym_h_1, tmp81704)

				var ifres81682 Obj

				if True == tmp81705 {
					tmp81694 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81695 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81696 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81697 := Call(__e, tmp81696, V1045)

					tmp81698 := Call(__e, tmp81695, tmp81697)

					tmp81699 := Call(__e, tmp81694, tmp81698)

					var ifres81684 Obj

					if True == tmp81699 {
						tmp81686 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp81687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81688 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81689 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81690 := Call(__e, tmp81689, V1045)

						tmp81691 := Call(__e, tmp81688, tmp81690)

						tmp81692 := Call(__e, tmp81687, tmp81691)

						tmp81693 := Call(__e, tmp81686, Nil, tmp81692)

						var ifres81685 Obj

						if True == tmp81693 {
							ifres81685 = True

						} else {
							ifres81685 = False

						}

						ifres81684 = ifres81685

					} else {
						ifres81684 = False

					}

					var ifres81683 Obj

					if True == ifres81684 {
						ifres81683 = True

					} else {
						ifres81683 = False

					}

					ifres81682 = ifres81683

				} else {
					ifres81682 = False

				}

				var ifres81681 Obj

				if True == ifres81682 {
					ifres81681 = True

				} else {
					ifres81681 = False

				}

				ifres81680 = ifres81681

			} else {
				ifres81680 = False

			}

			var ifres81679 Obj

			if True == ifres81680 {
				ifres81679 = True

			} else {
				ifres81679 = False

			}

			ifres81678 = ifres81679

		} else {
			ifres81678 = False

		}

		if True == ifres81678 {
			tmp81666 := MakeNative(func(__e *ControlFlow) {
				Linear := __e.Get(1)
				_ = Linear
				tmp81667 := Call(__e, PrimNS1Value(symns2_1value), symshen_4clause__form)

				__e.TailApply(tmp81667, Linear)
				return

			}, 1)

			tmp81668 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise)

			tmp81669 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81670 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81671 := Call(__e, tmp81670, V1045)

			tmp81672 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81673 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81674 := Call(__e, tmp81673, V1045)

			tmp81675 := Call(__e, tmp81672, tmp81674)

			tmp81676 := Call(__e, tmp81669, tmp81671, tmp81675)

			tmp81677 := Call(__e, tmp81668, tmp81676)

			__e.TailApply(tmp81666, tmp81677)
			return

		} else {
			tmp81665 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp81665, symshen_4linearise_1clause)
			return

		}

	}, 1)

	tmp81712 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise_1clause, tmp81663)

	_ = tmp81712

	tmp81713 := MakeNative(func(__e *ControlFlow) {
		V1047 := __e.Get(1)
		_ = V1047
		tmp81745 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81746 := Call(__e, tmp81745, V1047)

		var ifres81731 Obj

		if True == tmp81746 {
			tmp81741 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81742 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81743 := Call(__e, tmp81742, V1047)

			tmp81744 := Call(__e, tmp81741, tmp81743)

			var ifres81733 Obj

			if True == tmp81744 {
				tmp81735 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp81736 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81737 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81738 := Call(__e, tmp81737, V1047)

				tmp81739 := Call(__e, tmp81736, tmp81738)

				tmp81740 := Call(__e, tmp81735, Nil, tmp81739)

				var ifres81734 Obj

				if True == tmp81740 {
					ifres81734 = True

				} else {
					ifres81734 = False

				}

				ifres81733 = ifres81734

			} else {
				ifres81733 = False

			}

			var ifres81732 Obj

			if True == ifres81733 {
				ifres81732 = True

			} else {
				ifres81732 = False

			}

			ifres81731 = ifres81732

		} else {
			ifres81731 = False

		}

		if True == ifres81731 {
			tmp81716 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81717 := Call(__e, PrimNS1Value(symns2_1value), symshen_4explicit__modes)

			tmp81718 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81719 := Call(__e, tmp81718, V1047)

			tmp81720 := Call(__e, tmp81717, tmp81719)

			tmp81721 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81722 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81723 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cf__help)

			tmp81724 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81725 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81726 := Call(__e, tmp81725, V1047)

			tmp81727 := Call(__e, tmp81724, tmp81726)

			tmp81728 := Call(__e, tmp81723, tmp81727)

			tmp81729 := Call(__e, tmp81722, tmp81728, Nil)

			tmp81730 := Call(__e, tmp81721, sym_h_1, tmp81729)

			__e.TailApply(tmp81716, tmp81720, tmp81730)
			return

		} else {
			tmp81715 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp81715, symshen_4clause__form)
			return

		}

	}, 1)

	tmp81747 := Call(__e, PrimNS1Value(symns2_1set), symshen_4clause__form, tmp81713)

	_ = tmp81747

	tmp81748 := MakeNative(func(__e *ControlFlow) {
		V1049 := __e.Get(1)
		_ = V1049
		tmp81760 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81761 := Call(__e, tmp81760, V1049)

		if True == tmp81761 {
			tmp81751 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81752 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81753 := Call(__e, tmp81752, V1049)

			tmp81754 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp81755 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp81756 := Call(__e, PrimNS1Value(symns2_1value), symshen_4em__help)

				__e.TailApply(tmp81756, X)
				return

			}, 1)

			tmp81757 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81758 := Call(__e, tmp81757, V1049)

			tmp81759 := Call(__e, tmp81754, tmp81755, tmp81758)

			__e.TailApply(tmp81751, tmp81753, tmp81759)
			return

		} else {
			tmp81750 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp81750, symshen_4explicit__modes)
			return

		}

	}, 1)

	tmp81762 := Call(__e, PrimNS1Value(symns2_1set), symshen_4explicit__modes, tmp81748)

	_ = tmp81762

	tmp81763 := MakeNative(func(__e *ControlFlow) {
		V1051 := __e.Get(1)
		_ = V1051
		tmp81800 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81801 := Call(__e, tmp81800, V1051)

		var ifres81770 Obj

		if True == tmp81801 {
			tmp81796 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81797 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81798 := Call(__e, tmp81797, V1051)

			tmp81799 := Call(__e, tmp81796, symmode, tmp81798)

			var ifres81772 Obj

			if True == tmp81799 {
				tmp81792 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81793 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81794 := Call(__e, tmp81793, V1051)

				tmp81795 := Call(__e, tmp81792, tmp81794)

				var ifres81774 Obj

				if True == tmp81795 {
					tmp81786 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81787 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81788 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81789 := Call(__e, tmp81788, V1051)

					tmp81790 := Call(__e, tmp81787, tmp81789)

					tmp81791 := Call(__e, tmp81786, tmp81790)

					var ifres81776 Obj

					if True == tmp81791 {
						tmp81778 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp81779 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81781 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81782 := Call(__e, tmp81781, V1051)

						tmp81783 := Call(__e, tmp81780, tmp81782)

						tmp81784 := Call(__e, tmp81779, tmp81783)

						tmp81785 := Call(__e, tmp81778, Nil, tmp81784)

						var ifres81777 Obj

						if True == tmp81785 {
							ifres81777 = True

						} else {
							ifres81777 = False

						}

						ifres81776 = ifres81777

					} else {
						ifres81776 = False

					}

					var ifres81775 Obj

					if True == ifres81776 {
						ifres81775 = True

					} else {
						ifres81775 = False

					}

					ifres81774 = ifres81775

				} else {
					ifres81774 = False

				}

				var ifres81773 Obj

				if True == ifres81774 {
					ifres81773 = True

				} else {
					ifres81773 = False

				}

				ifres81772 = ifres81773

			} else {
				ifres81772 = False

			}

			var ifres81771 Obj

			if True == ifres81772 {
				ifres81771 = True

			} else {
				ifres81771 = False

			}

			ifres81770 = ifres81771

		} else {
			ifres81770 = False

		}

		if True == ifres81770 {
			__e.Return(V1051)
			return
		} else {
			tmp81765 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81766 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81767 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81768 := Call(__e, tmp81767, sym_7, Nil)

			tmp81769 := Call(__e, tmp81766, V1051, tmp81768)

			__e.TailApply(tmp81765, symmode, tmp81769)
			return

		}

	}, 1)

	tmp81802 := Call(__e, PrimNS1Value(symns2_1set), symshen_4em__help, tmp81763)

	_ = tmp81802

	tmp81803 := MakeNative(func(__e *ControlFlow) {
		V1053 := __e.Get(1)
		_ = V1053
		tmp81909 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81910 := Call(__e, tmp81909, V1053)

		var ifres81825 Obj

		if True == tmp81910 {
			tmp81905 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81906 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81907 := Call(__e, tmp81906, V1053)

			tmp81908 := Call(__e, tmp81905, symwhere, tmp81907)

			var ifres81827 Obj

			if True == tmp81908 {
				tmp81901 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81902 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81903 := Call(__e, tmp81902, V1053)

				tmp81904 := Call(__e, tmp81901, tmp81903)

				var ifres81829 Obj

				if True == tmp81904 {
					tmp81895 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp81896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81897 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81898 := Call(__e, tmp81897, V1053)

					tmp81899 := Call(__e, tmp81896, tmp81898)

					tmp81900 := Call(__e, tmp81895, tmp81899)

					var ifres81831 Obj

					if True == tmp81900 {
						tmp81887 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp81888 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp81889 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp81890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81891 := Call(__e, tmp81890, V1053)

						tmp81892 := Call(__e, tmp81889, tmp81891)

						tmp81893 := Call(__e, tmp81888, tmp81892)

						tmp81894 := Call(__e, tmp81887, sym_a, tmp81893)

						var ifres81833 Obj

						if True == tmp81894 {
							tmp81879 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp81880 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81881 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp81882 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81883 := Call(__e, tmp81882, V1053)

							tmp81884 := Call(__e, tmp81881, tmp81883)

							tmp81885 := Call(__e, tmp81880, tmp81884)

							tmp81886 := Call(__e, tmp81879, tmp81885)

							var ifres81835 Obj

							if True == tmp81886 {
								tmp81869 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp81870 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81871 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81872 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp81873 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp81874 := Call(__e, tmp81873, V1053)

								tmp81875 := Call(__e, tmp81872, tmp81874)

								tmp81876 := Call(__e, tmp81871, tmp81875)

								tmp81877 := Call(__e, tmp81870, tmp81876)

								tmp81878 := Call(__e, tmp81869, tmp81877)

								var ifres81837 Obj

								if True == tmp81878 {
									tmp81857 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp81858 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81859 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81860 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81861 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp81862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp81863 := Call(__e, tmp81862, V1053)

									tmp81864 := Call(__e, tmp81861, tmp81863)

									tmp81865 := Call(__e, tmp81860, tmp81864)

									tmp81866 := Call(__e, tmp81859, tmp81865)

									tmp81867 := Call(__e, tmp81858, tmp81866)

									tmp81868 := Call(__e, tmp81857, Nil, tmp81867)

									var ifres81839 Obj

									if True == tmp81868 {
										tmp81851 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp81852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp81853 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp81854 := Call(__e, tmp81853, V1053)

										tmp81855 := Call(__e, tmp81852, tmp81854)

										tmp81856 := Call(__e, tmp81851, tmp81855)

										var ifres81841 Obj

										if True == tmp81856 {
											tmp81843 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp81844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp81845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp81846 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp81847 := Call(__e, tmp81846, V1053)

											tmp81848 := Call(__e, tmp81845, tmp81847)

											tmp81849 := Call(__e, tmp81844, tmp81848)

											tmp81850 := Call(__e, tmp81843, Nil, tmp81849)

											var ifres81842 Obj

											if True == tmp81850 {
												ifres81842 = True

											} else {
												ifres81842 = False

											}

											ifres81841 = ifres81842

										} else {
											ifres81841 = False

										}

										var ifres81840 Obj

										if True == ifres81841 {
											ifres81840 = True

										} else {
											ifres81840 = False

										}

										ifres81839 = ifres81840

									} else {
										ifres81839 = False

									}

									var ifres81838 Obj

									if True == ifres81839 {
										ifres81838 = True

									} else {
										ifres81838 = False

									}

									ifres81837 = ifres81838

								} else {
									ifres81837 = False

								}

								var ifres81836 Obj

								if True == ifres81837 {
									ifres81836 = True

								} else {
									ifres81836 = False

								}

								ifres81835 = ifres81836

							} else {
								ifres81835 = False

							}

							var ifres81834 Obj

							if True == ifres81835 {
								ifres81834 = True

							} else {
								ifres81834 = False

							}

							ifres81833 = ifres81834

						} else {
							ifres81833 = False

						}

						var ifres81832 Obj

						if True == ifres81833 {
							ifres81832 = True

						} else {
							ifres81832 = False

						}

						ifres81831 = ifres81832

					} else {
						ifres81831 = False

					}

					var ifres81830 Obj

					if True == ifres81831 {
						ifres81830 = True

					} else {
						ifres81830 = False

					}

					ifres81829 = ifres81830

				} else {
					ifres81829 = False

				}

				var ifres81828 Obj

				if True == ifres81829 {
					ifres81828 = True

				} else {
					ifres81828 = False

				}

				ifres81827 = ifres81828

			} else {
				ifres81827 = False

			}

			var ifres81826 Obj

			if True == ifres81827 {
				ifres81826 = True

			} else {
				ifres81826 = False

			}

			ifres81825 = ifres81826

		} else {
			ifres81825 = False

		}

		if True == ifres81825 {
			tmp81805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81808 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp81809 := Call(__e, tmp81808, symshen_4_doccurs_d)

			var ifres81807 Obj

			if True == tmp81809 {
				ifres81807 = symunify_b

			} else {
				ifres81807 = symunify

			}

			tmp81810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81811 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81812 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81813 := Call(__e, tmp81812, V1053)

			tmp81814 := Call(__e, tmp81811, tmp81813)

			tmp81815 := Call(__e, tmp81810, tmp81814)

			tmp81816 := Call(__e, tmp81806, ifres81807, tmp81815)

			tmp81817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cf__help)

			tmp81818 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81819 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81820 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81821 := Call(__e, tmp81820, V1053)

			tmp81822 := Call(__e, tmp81819, tmp81821)

			tmp81823 := Call(__e, tmp81818, tmp81822)

			tmp81824 := Call(__e, tmp81817, tmp81823)

			__e.TailApply(tmp81805, tmp81816, tmp81824)
			return

		} else {
			__e.Return(V1053)
			return
		}

	}, 1)

	tmp81911 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cf__help, tmp81803)

	_ = tmp81911

	tmp81912 := MakeNative(func(__e *ControlFlow) {
		V1059 := __e.Get(1)
		_ = V1059
		tmp81920 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp81921 := Call(__e, tmp81920, sym_7, V1059)

		if True == tmp81921 {
			tmp81919 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(tmp81919, symshen_4_doccurs_d, True)
			return

		} else {
			tmp81917 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp81918 := Call(__e, tmp81917, sym_1, V1059)

			if True == tmp81918 {
				tmp81916 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(tmp81916, symshen_4_doccurs_d, False)
				return

			} else {
				tmp81915 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp81915, MakeString("occurs-check expects + or -\n"))
				return

			}

		}

	}, 1)

	tmp81922 := Call(__e, PrimNS1Value(symns2_1set), symoccurs_1check, tmp81912)

	_ = tmp81922

	tmp81923 := MakeNative(func(__e *ControlFlow) {
		V1062 := __e.Get(1)
		_ = V1062
		V1063 := __e.Get(2)
		_ = V1063
		tmp81990 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp81991 := Call(__e, tmp81990, V1062)

		var ifres81952 Obj

		if True == tmp81991 {
			tmp81986 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp81987 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81988 := Call(__e, tmp81987, V1062)

			tmp81989 := Call(__e, tmp81986, tmp81988)

			var ifres81954 Obj

			if True == tmp81989 {
				tmp81982 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp81983 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp81984 := Call(__e, tmp81983, V1062)

				tmp81985 := Call(__e, tmp81982, tmp81984)

				var ifres81956 Obj

				if True == tmp81985 {
					tmp81976 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp81977 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp81978 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp81979 := Call(__e, tmp81978, V1062)

					tmp81980 := Call(__e, tmp81977, tmp81979)

					tmp81981 := Call(__e, tmp81976, sym_h_1, tmp81980)

					var ifres81958 Obj

					if True == tmp81981 {
						tmp81970 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp81971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81972 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp81973 := Call(__e, tmp81972, V1062)

						tmp81974 := Call(__e, tmp81971, tmp81973)

						tmp81975 := Call(__e, tmp81970, tmp81974)

						var ifres81960 Obj

						if True == tmp81975 {
							tmp81962 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp81963 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81964 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp81966 := Call(__e, tmp81965, V1062)

							tmp81967 := Call(__e, tmp81964, tmp81966)

							tmp81968 := Call(__e, tmp81963, tmp81967)

							tmp81969 := Call(__e, tmp81962, Nil, tmp81968)

							var ifres81961 Obj

							if True == tmp81969 {
								ifres81961 = True

							} else {
								ifres81961 = False

							}

							ifres81960 = ifres81961

						} else {
							ifres81960 = False

						}

						var ifres81959 Obj

						if True == ifres81960 {
							ifres81959 = True

						} else {
							ifres81959 = False

						}

						ifres81958 = ifres81959

					} else {
						ifres81958 = False

					}

					var ifres81957 Obj

					if True == ifres81958 {
						ifres81957 = True

					} else {
						ifres81957 = False

					}

					ifres81956 = ifres81957

				} else {
					ifres81956 = False

				}

				var ifres81955 Obj

				if True == ifres81956 {
					ifres81955 = True

				} else {
					ifres81955 = False

				}

				ifres81954 = ifres81955

			} else {
				ifres81954 = False

			}

			var ifres81953 Obj

			if True == ifres81954 {
				ifres81953 = True

			} else {
				ifres81953 = False

			}

			ifres81952 = ifres81953

		} else {
			ifres81952 = False

		}

		if True == ifres81952 {
			tmp81926 := MakeNative(func(__e *ControlFlow) {
				MuApplication := __e.Get(1)
				_ = MuApplication
				tmp81927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

				__e.TailApply(tmp81927, MuApplication, sym_7)
				return

			}, 1)

			tmp81928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4make__mu__application)

			tmp81929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81930 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81931 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81932 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81933 := Call(__e, tmp81932, V1062)

			tmp81934 := Call(__e, tmp81931, tmp81933)

			tmp81935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp81936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4continuation__call)

			tmp81937 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81938 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81939 := Call(__e, tmp81938, V1062)

			tmp81940 := Call(__e, tmp81937, tmp81939)

			tmp81941 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp81942 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81943 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp81944 := Call(__e, tmp81943, V1062)

			tmp81945 := Call(__e, tmp81942, tmp81944)

			tmp81946 := Call(__e, tmp81941, tmp81945)

			tmp81947 := Call(__e, tmp81936, tmp81940, tmp81946)

			tmp81948 := Call(__e, tmp81935, tmp81947, Nil)

			tmp81949 := Call(__e, tmp81930, tmp81934, tmp81948)

			tmp81950 := Call(__e, tmp81929, symshen_4mu, tmp81949)

			tmp81951 := Call(__e, tmp81928, tmp81950, V1063)

			__e.TailApply(tmp81926, tmp81951)
			return

		} else {
			tmp81925 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp81925, symshen_4aum)
			return

		}

	}, 2)

	tmp81992 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aum, tmp81923)

	_ = tmp81992

	tmp81993 := MakeNative(func(__e *ControlFlow) {
		V1066 := __e.Get(1)
		_ = V1066
		V1067 := __e.Get(2)
		_ = V1067
		tmp81994 := MakeNative(func(__e *ControlFlow) {
			VTerms := __e.Get(1)
			_ = VTerms
			tmp81995 := MakeNative(func(__e *ControlFlow) {
				VBody := __e.Get(1)
				_ = VBody
				tmp81996 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					tmp81997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cc__help)

					__e.TailApply(tmp81997, Free, V1067)
					return

				}, 1)

				tmp81998 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				tmp81999 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

				tmp82000 := Call(__e, tmp81999, VBody, VTerms)

				tmp82001 := Call(__e, tmp81998, symThrowcontrol, tmp82000)

				__e.TailApply(tmp81996, tmp82001)
				return

			}, 1)

			tmp82002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

			tmp82003 := Call(__e, tmp82002, V1067)

			__e.TailApply(tmp81995, tmp82003)
			return

		}, 1)

		tmp82004 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp82005 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

		tmp82006 := Call(__e, tmp82005, V1066)

		tmp82007 := Call(__e, tmp82004, symProcessN, tmp82006)

		__e.TailApply(tmp81994, tmp82007)
		return

	}, 2)

	tmp82008 := Call(__e, PrimNS1Value(symns2_1set), symshen_4continuation__call, tmp81993)

	_ = tmp82008

	tmp82009 := MakeNative(func(__e *ControlFlow) {
		V1070 := __e.Get(1)
		_ = V1070
		V1071 := __e.Get(2)
		_ = V1071
		tmp82010 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1h)

		__e.TailApply(tmp82010, V1070, V1071, Nil)
		return

	}, 2)

	tmp82011 := Call(__e, PrimNS1Value(symns2_1set), symremove, tmp82009)

	_ = tmp82011

	tmp82012 := MakeNative(func(__e *ControlFlow) {
		V1078 := __e.Get(1)
		_ = V1078
		V1079 := __e.Get(2)
		_ = V1079
		V1080 := __e.Get(3)
		_ = V1080
		tmp82040 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp82041 := Call(__e, tmp82040, Nil, V1079)

		if True == tmp82041 {
			tmp82039 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			__e.TailApply(tmp82039, V1080)
			return

		} else {
			tmp82037 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp82038 := Call(__e, tmp82037, V1079)

			var ifres82031 Obj

			if True == tmp82038 {
				tmp82033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp82034 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82035 := Call(__e, tmp82034, V1079)

				tmp82036 := Call(__e, tmp82033, tmp82035, V1078)

				var ifres82032 Obj

				if True == tmp82036 {
					ifres82032 = True

				} else {
					ifres82032 = False

				}

				ifres82031 = ifres82032

			} else {
				ifres82031 = False

			}

			if True == ifres82031 {
				tmp82026 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1h)

				tmp82027 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82028 := Call(__e, tmp82027, V1079)

				tmp82029 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82030 := Call(__e, tmp82029, V1079)

				__e.TailApply(tmp82026, tmp82028, tmp82030, V1080)
				return

			} else {
				tmp82024 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp82025 := Call(__e, tmp82024, V1079)

				if True == tmp82025 {
					tmp82017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1h)

					tmp82018 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp82019 := Call(__e, tmp82018, V1079)

					tmp82020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82021 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp82022 := Call(__e, tmp82021, V1079)

					tmp82023 := Call(__e, tmp82020, tmp82022, V1080)

					__e.TailApply(tmp82017, V1078, tmp82019, tmp82023)
					return

				} else {
					tmp82016 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp82016, symshen_4remove_1h)
					return

				}

			}

		}

	}, 3)

	tmp82042 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1h, tmp82012)

	_ = tmp82042

	tmp82043 := MakeNative(func(__e *ControlFlow) {
		V1083 := __e.Get(1)
		_ = V1083
		V1084 := __e.Get(2)
		_ = V1084
		tmp82111 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp82112 := Call(__e, tmp82111, Nil, V1083)

		var ifres82107 Obj

		if True == tmp82112 {
			tmp82109 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp82110 := Call(__e, tmp82109, Nil, V1084)

			var ifres82108 Obj

			if True == tmp82110 {
				ifres82108 = True

			} else {
				ifres82108 = False

			}

			ifres82107 = ifres82108

		} else {
			ifres82107 = False

		}

		if True == ifres82107 {
			tmp82102 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp82103 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp82104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp82105 := Call(__e, tmp82104, symshen_4stack, Nil)

			tmp82106 := Call(__e, tmp82103, symshen_4the, tmp82105)

			__e.TailApply(tmp82102, symshen_4pop, tmp82106)
			return

		} else {
			tmp82100 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp82101 := Call(__e, tmp82100, Nil, V1084)

			if True == tmp82101 {
				tmp82079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82081 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82090 := Call(__e, tmp82089, symshen_4stack, Nil)

				tmp82091 := Call(__e, tmp82088, symshen_4the, tmp82090)

				tmp82092 := Call(__e, tmp82087, symshen_4pop, tmp82091)

				tmp82093 := Call(__e, tmp82086, tmp82092, Nil)

				tmp82094 := Call(__e, tmp82085, symshen_4then, tmp82093)

				tmp82095 := Call(__e, tmp82084, symand, tmp82094)

				tmp82096 := Call(__e, tmp82083, V1083, tmp82095)

				tmp82097 := Call(__e, tmp82082, symin, tmp82096)

				tmp82098 := Call(__e, tmp82081, symshen_4variables, tmp82097)

				tmp82099 := Call(__e, tmp82080, symshen_4the, tmp82098)

				__e.TailApply(tmp82079, symshen_4rename, tmp82099)
				return

			} else {
				tmp82077 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp82078 := Call(__e, tmp82077, Nil, V1083)

				if True == tmp82078 {
					tmp82070 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82071 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82072 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82073 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82074 := Call(__e, tmp82073, V1084, Nil)

					tmp82075 := Call(__e, tmp82072, symshen_4continuation, tmp82074)

					tmp82076 := Call(__e, tmp82071, symshen_4the, tmp82075)

					__e.TailApply(tmp82070, symcall, tmp82076)
					return

				} else {
					tmp82047 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82048 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82049 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82050 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82051 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82054 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82055 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82057 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp82059 := Call(__e, tmp82058, V1084, Nil)

					tmp82060 := Call(__e, tmp82057, symshen_4continuation, tmp82059)

					tmp82061 := Call(__e, tmp82056, symshen_4the, tmp82060)

					tmp82062 := Call(__e, tmp82055, symcall, tmp82061)

					tmp82063 := Call(__e, tmp82054, tmp82062, Nil)

					tmp82064 := Call(__e, tmp82053, symshen_4then, tmp82063)

					tmp82065 := Call(__e, tmp82052, symand, tmp82064)

					tmp82066 := Call(__e, tmp82051, V1083, tmp82065)

					tmp82067 := Call(__e, tmp82050, symin, tmp82066)

					tmp82068 := Call(__e, tmp82049, symshen_4variables, tmp82067)

					tmp82069 := Call(__e, tmp82048, symshen_4the, tmp82068)

					__e.TailApply(tmp82047, symshen_4rename, tmp82069)
					return

				}

			}

		}

	}, 2)

	tmp82113 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cc__help, tmp82043)

	_ = tmp82113

	tmp82114 := MakeNative(func(__e *ControlFlow) {
		V1087 := __e.Get(1)
		_ = V1087
		V1088 := __e.Get(2)
		_ = V1088
		tmp82244 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp82245 := Call(__e, tmp82244, V1087)

		var ifres82202 Obj

		if True == tmp82245 {
			tmp82240 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp82241 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp82242 := Call(__e, tmp82241, V1087)

			tmp82243 := Call(__e, tmp82240, symshen_4mu, tmp82242)

			var ifres82204 Obj

			if True == tmp82243 {
				tmp82236 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp82237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82238 := Call(__e, tmp82237, V1087)

				tmp82239 := Call(__e, tmp82236, tmp82238)

				var ifres82206 Obj

				if True == tmp82239 {
					tmp82230 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp82231 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp82232 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp82233 := Call(__e, tmp82232, V1087)

					tmp82234 := Call(__e, tmp82231, tmp82233)

					tmp82235 := Call(__e, tmp82230, Nil, tmp82234)

					var ifres82208 Obj

					if True == tmp82235 {
						tmp82224 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp82225 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp82226 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp82227 := Call(__e, tmp82226, V1087)

						tmp82228 := Call(__e, tmp82225, tmp82227)

						tmp82229 := Call(__e, tmp82224, tmp82228)

						var ifres82210 Obj

						if True == tmp82229 {
							tmp82216 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp82217 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp82218 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp82219 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp82220 := Call(__e, tmp82219, V1087)

							tmp82221 := Call(__e, tmp82218, tmp82220)

							tmp82222 := Call(__e, tmp82217, tmp82221)

							tmp82223 := Call(__e, tmp82216, Nil, tmp82222)

							var ifres82212 Obj

							if True == tmp82223 {
								tmp82214 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp82215 := Call(__e, tmp82214, Nil, V1088)

								var ifres82213 Obj

								if True == tmp82215 {
									ifres82213 = True

								} else {
									ifres82213 = False

								}

								ifres82212 = ifres82213

							} else {
								ifres82212 = False

							}

							var ifres82211 Obj

							if True == ifres82212 {
								ifres82211 = True

							} else {
								ifres82211 = False

							}

							ifres82210 = ifres82211

						} else {
							ifres82210 = False

						}

						var ifres82209 Obj

						if True == ifres82210 {
							ifres82209 = True

						} else {
							ifres82209 = False

						}

						ifres82208 = ifres82209

					} else {
						ifres82208 = False

					}

					var ifres82207 Obj

					if True == ifres82208 {
						ifres82207 = True

					} else {
						ifres82207 = False

					}

					ifres82206 = ifres82207

				} else {
					ifres82206 = False

				}

				var ifres82205 Obj

				if True == ifres82206 {
					ifres82205 = True

				} else {
					ifres82205 = False

				}

				ifres82204 = ifres82205

			} else {
				ifres82204 = False

			}

			var ifres82203 Obj

			if True == ifres82204 {
				ifres82203 = True

			} else {
				ifres82203 = False

			}

			ifres82202 = ifres82203

		} else {
			ifres82202 = False

		}

		if True == ifres82202 {
			tmp82197 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp82198 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp82199 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp82200 := Call(__e, tmp82199, V1087)

			tmp82201 := Call(__e, tmp82198, tmp82200)

			__e.TailApply(tmp82197, tmp82201)
			return

		} else {
			tmp82195 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp82196 := Call(__e, tmp82195, V1087)

			var ifres82153 Obj

			if True == tmp82196 {
				tmp82191 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp82192 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82193 := Call(__e, tmp82192, V1087)

				tmp82194 := Call(__e, tmp82191, symshen_4mu, tmp82193)

				var ifres82155 Obj

				if True == tmp82194 {
					tmp82187 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp82188 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp82189 := Call(__e, tmp82188, V1087)

					tmp82190 := Call(__e, tmp82187, tmp82189)

					var ifres82157 Obj

					if True == tmp82190 {
						tmp82181 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp82182 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp82183 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp82184 := Call(__e, tmp82183, V1087)

						tmp82185 := Call(__e, tmp82182, tmp82184)

						tmp82186 := Call(__e, tmp82181, tmp82185)

						var ifres82159 Obj

						if True == tmp82186 {
							tmp82175 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp82176 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp82177 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp82178 := Call(__e, tmp82177, V1087)

							tmp82179 := Call(__e, tmp82176, tmp82178)

							tmp82180 := Call(__e, tmp82175, tmp82179)

							var ifres82161 Obj

							if True == tmp82180 {
								tmp82167 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp82168 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82171 := Call(__e, tmp82170, V1087)

								tmp82172 := Call(__e, tmp82169, tmp82171)

								tmp82173 := Call(__e, tmp82168, tmp82172)

								tmp82174 := Call(__e, tmp82167, Nil, tmp82173)

								var ifres82163 Obj

								if True == tmp82174 {
									tmp82165 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp82166 := Call(__e, tmp82165, V1088)

									var ifres82164 Obj

									if True == tmp82166 {
										ifres82164 = True

									} else {
										ifres82164 = False

									}

									ifres82163 = ifres82164

								} else {
									ifres82163 = False

								}

								var ifres82162 Obj

								if True == ifres82163 {
									ifres82162 = True

								} else {
									ifres82162 = False

								}

								ifres82161 = ifres82162

							} else {
								ifres82161 = False

							}

							var ifres82160 Obj

							if True == ifres82161 {
								ifres82160 = True

							} else {
								ifres82160 = False

							}

							ifres82159 = ifres82160

						} else {
							ifres82159 = False

						}

						var ifres82158 Obj

						if True == ifres82159 {
							ifres82158 = True

						} else {
							ifres82158 = False

						}

						ifres82157 = ifres82158

					} else {
						ifres82157 = False

					}

					var ifres82156 Obj

					if True == ifres82157 {
						ifres82156 = True

					} else {
						ifres82156 = False

					}

					ifres82155 = ifres82156

				} else {
					ifres82155 = False

				}

				var ifres82154 Obj

				if True == ifres82155 {
					ifres82154 = True

				} else {
					ifres82154 = False

				}

				ifres82153 = ifres82154

			} else {
				ifres82153 = False

			}

			if True == ifres82153 {
				tmp82118 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82120 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82121 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82122 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82124 := Call(__e, tmp82123, V1087)

				tmp82125 := Call(__e, tmp82122, tmp82124)

				tmp82126 := Call(__e, tmp82121, tmp82125)

				tmp82127 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82128 := Call(__e, PrimNS1Value(symns2_1value), symshen_4make__mu__application)

				tmp82129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82132 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82133 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82134 := Call(__e, tmp82133, V1087)

				tmp82135 := Call(__e, tmp82132, tmp82134)

				tmp82136 := Call(__e, tmp82131, tmp82135)

				tmp82137 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82138 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82139 := Call(__e, tmp82138, V1087)

				tmp82140 := Call(__e, tmp82137, tmp82139)

				tmp82141 := Call(__e, tmp82130, tmp82136, tmp82140)

				tmp82142 := Call(__e, tmp82129, symshen_4mu, tmp82141)

				tmp82143 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp82144 := Call(__e, tmp82143, V1088)

				tmp82145 := Call(__e, tmp82128, tmp82142, tmp82144)

				tmp82146 := Call(__e, tmp82127, tmp82145, Nil)

				tmp82147 := Call(__e, tmp82120, tmp82126, tmp82146)

				tmp82148 := Call(__e, tmp82119, symshen_4mu, tmp82147)

				tmp82149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp82150 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp82151 := Call(__e, tmp82150, V1088)

				tmp82152 := Call(__e, tmp82149, tmp82151, Nil)

				__e.TailApply(tmp82118, tmp82148, tmp82152)
				return

			} else {
				tmp82117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp82117, symshen_4make__mu__application)
				return

			}

		}

	}, 2)

	tmp82246 := Call(__e, PrimNS1Value(symns2_1set), symshen_4make__mu__application, tmp82114)

	_ = tmp82246

	tmp82247 := MakeNative(func(__e *ControlFlow) {
		V1097 := __e.Get(1)
		_ = V1097
		V1098 := __e.Get(2)
		_ = V1098
		tmp83442 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp83443 := Call(__e, tmp83442, V1097)

		var ifres83320 Obj

		if True == tmp83443 {
			tmp83438 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp83439 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83440 := Call(__e, tmp83439, V1097)

			tmp83441 := Call(__e, tmp83438, tmp83440)

			var ifres83322 Obj

			if True == tmp83441 {
				tmp83432 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp83433 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83434 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83435 := Call(__e, tmp83434, V1097)

				tmp83436 := Call(__e, tmp83433, tmp83435)

				tmp83437 := Call(__e, tmp83432, symshen_4mu, tmp83436)

				var ifres83324 Obj

				if True == tmp83437 {
					tmp83426 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp83427 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83428 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83429 := Call(__e, tmp83428, V1097)

					tmp83430 := Call(__e, tmp83427, tmp83429)

					tmp83431 := Call(__e, tmp83426, tmp83430)

					var ifres83326 Obj

					if True == tmp83431 {
						tmp83418 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp83419 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83420 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83421 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83422 := Call(__e, tmp83421, V1097)

						tmp83423 := Call(__e, tmp83420, tmp83422)

						tmp83424 := Call(__e, tmp83419, tmp83423)

						tmp83425 := Call(__e, tmp83418, tmp83424)

						var ifres83328 Obj

						if True == tmp83425 {
							tmp83408 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp83409 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83410 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83411 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83412 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83413 := Call(__e, tmp83412, V1097)

							tmp83414 := Call(__e, tmp83411, tmp83413)

							tmp83415 := Call(__e, tmp83410, tmp83414)

							tmp83416 := Call(__e, tmp83409, tmp83415)

							tmp83417 := Call(__e, tmp83408, symmode, tmp83416)

							var ifres83330 Obj

							if True == tmp83417 {
								tmp83398 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp83399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83400 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp83401 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83402 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp83403 := Call(__e, tmp83402, V1097)

								tmp83404 := Call(__e, tmp83401, tmp83403)

								tmp83405 := Call(__e, tmp83400, tmp83404)

								tmp83406 := Call(__e, tmp83399, tmp83405)

								tmp83407 := Call(__e, tmp83398, tmp83406)

								var ifres83332 Obj

								if True == tmp83407 {
									tmp83386 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp83387 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83388 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83389 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp83390 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83391 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp83392 := Call(__e, tmp83391, V1097)

									tmp83393 := Call(__e, tmp83390, tmp83392)

									tmp83394 := Call(__e, tmp83389, tmp83393)

									tmp83395 := Call(__e, tmp83388, tmp83394)

									tmp83396 := Call(__e, tmp83387, tmp83395)

									tmp83397 := Call(__e, tmp83386, tmp83396)

									var ifres83334 Obj

									if True == tmp83397 {
										tmp83372 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp83373 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83374 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83375 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83376 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp83377 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83378 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp83379 := Call(__e, tmp83378, V1097)

										tmp83380 := Call(__e, tmp83377, tmp83379)

										tmp83381 := Call(__e, tmp83376, tmp83380)

										tmp83382 := Call(__e, tmp83375, tmp83381)

										tmp83383 := Call(__e, tmp83374, tmp83382)

										tmp83384 := Call(__e, tmp83373, tmp83383)

										tmp83385 := Call(__e, tmp83372, Nil, tmp83384)

										var ifres83336 Obj

										if True == tmp83385 {
											tmp83364 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp83365 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83366 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83367 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp83368 := Call(__e, tmp83367, V1097)

											tmp83369 := Call(__e, tmp83366, tmp83368)

											tmp83370 := Call(__e, tmp83365, tmp83369)

											tmp83371 := Call(__e, tmp83364, tmp83370)

											var ifres83338 Obj

											if True == tmp83371 {
												tmp83354 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp83355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83356 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83357 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83358 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp83359 := Call(__e, tmp83358, V1097)

												tmp83360 := Call(__e, tmp83357, tmp83359)

												tmp83361 := Call(__e, tmp83356, tmp83360)

												tmp83362 := Call(__e, tmp83355, tmp83361)

												tmp83363 := Call(__e, tmp83354, Nil, tmp83362)

												var ifres83340 Obj

												if True == tmp83363 {
													tmp83350 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp83351 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83352 := Call(__e, tmp83351, V1097)

													tmp83353 := Call(__e, tmp83350, tmp83352)

													var ifres83342 Obj

													if True == tmp83353 {
														tmp83344 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp83345 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83347 := Call(__e, tmp83346, V1097)

														tmp83348 := Call(__e, tmp83345, tmp83347)

														tmp83349 := Call(__e, tmp83344, Nil, tmp83348)

														var ifres83343 Obj

														if True == tmp83349 {
															ifres83343 = True

														} else {
															ifres83343 = False

														}

														ifres83342 = ifres83343

													} else {
														ifres83342 = False

													}

													var ifres83341 Obj

													if True == ifres83342 {
														ifres83341 = True

													} else {
														ifres83341 = False

													}

													ifres83340 = ifres83341

												} else {
													ifres83340 = False

												}

												var ifres83339 Obj

												if True == ifres83340 {
													ifres83339 = True

												} else {
													ifres83339 = False

												}

												ifres83338 = ifres83339

											} else {
												ifres83338 = False

											}

											var ifres83337 Obj

											if True == ifres83338 {
												ifres83337 = True

											} else {
												ifres83337 = False

											}

											ifres83336 = ifres83337

										} else {
											ifres83336 = False

										}

										var ifres83335 Obj

										if True == ifres83336 {
											ifres83335 = True

										} else {
											ifres83335 = False

										}

										ifres83334 = ifres83335

									} else {
										ifres83334 = False

									}

									var ifres83333 Obj

									if True == ifres83334 {
										ifres83333 = True

									} else {
										ifres83333 = False

									}

									ifres83332 = ifres83333

								} else {
									ifres83332 = False

								}

								var ifres83331 Obj

								if True == ifres83332 {
									ifres83331 = True

								} else {
									ifres83331 = False

								}

								ifres83330 = ifres83331

							} else {
								ifres83330 = False

							}

							var ifres83329 Obj

							if True == ifres83330 {
								ifres83329 = True

							} else {
								ifres83329 = False

							}

							ifres83328 = ifres83329

						} else {
							ifres83328 = False

						}

						var ifres83327 Obj

						if True == ifres83328 {
							ifres83327 = True

						} else {
							ifres83327 = False

						}

						ifres83326 = ifres83327

					} else {
						ifres83326 = False

					}

					var ifres83325 Obj

					if True == ifres83326 {
						ifres83325 = True

					} else {
						ifres83325 = False

					}

					ifres83324 = ifres83325

				} else {
					ifres83324 = False

				}

				var ifres83323 Obj

				if True == ifres83324 {
					ifres83323 = True

				} else {
					ifres83323 = False

				}

				ifres83322 = ifres83323

			} else {
				ifres83322 = False

			}

			var ifres83321 Obj

			if True == ifres83322 {
				ifres83321 = True

			} else {
				ifres83321 = False

			}

			ifres83320 = ifres83321

		} else {
			ifres83320 = False

		}

		if True == ifres83320 {
			tmp83283 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

			tmp83284 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp83285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp83286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp83287 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83289 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83290 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83291 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83292 := Call(__e, tmp83291, V1097)

			tmp83293 := Call(__e, tmp83290, tmp83292)

			tmp83294 := Call(__e, tmp83289, tmp83293)

			tmp83295 := Call(__e, tmp83288, tmp83294)

			tmp83296 := Call(__e, tmp83287, tmp83295)

			tmp83297 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83299 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83300 := Call(__e, tmp83299, V1097)

			tmp83301 := Call(__e, tmp83298, tmp83300)

			tmp83302 := Call(__e, tmp83297, tmp83301)

			tmp83303 := Call(__e, tmp83286, tmp83296, tmp83302)

			tmp83304 := Call(__e, tmp83285, symshen_4mu, tmp83303)

			tmp83305 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83306 := Call(__e, tmp83305, V1097)

			tmp83307 := Call(__e, tmp83284, tmp83304, tmp83306)

			tmp83308 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83309 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83310 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83311 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83312 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83313 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83314 := Call(__e, tmp83313, V1097)

			tmp83315 := Call(__e, tmp83312, tmp83314)

			tmp83316 := Call(__e, tmp83311, tmp83315)

			tmp83317 := Call(__e, tmp83310, tmp83316)

			tmp83318 := Call(__e, tmp83309, tmp83317)

			tmp83319 := Call(__e, tmp83308, tmp83318)

			__e.TailApply(tmp83283, tmp83307, tmp83319)
			return

		} else {
			tmp83281 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp83282 := Call(__e, tmp83281, V1097)

			var ifres83213 Obj

			if True == tmp83282 {
				tmp83277 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp83278 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83279 := Call(__e, tmp83278, V1097)

				tmp83280 := Call(__e, tmp83277, tmp83279)

				var ifres83215 Obj

				if True == tmp83280 {
					tmp83271 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp83272 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83273 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83274 := Call(__e, tmp83273, V1097)

					tmp83275 := Call(__e, tmp83272, tmp83274)

					tmp83276 := Call(__e, tmp83271, symshen_4mu, tmp83275)

					var ifres83217 Obj

					if True == tmp83276 {
						tmp83265 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp83266 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83267 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83268 := Call(__e, tmp83267, V1097)

						tmp83269 := Call(__e, tmp83266, tmp83268)

						tmp83270 := Call(__e, tmp83265, tmp83269)

						var ifres83219 Obj

						if True == tmp83270 {
							tmp83257 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp83258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83259 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83260 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83261 := Call(__e, tmp83260, V1097)

							tmp83262 := Call(__e, tmp83259, tmp83261)

							tmp83263 := Call(__e, tmp83258, tmp83262)

							tmp83264 := Call(__e, tmp83257, tmp83263)

							var ifres83221 Obj

							if True == tmp83264 {
								tmp83247 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp83248 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83251 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp83252 := Call(__e, tmp83251, V1097)

								tmp83253 := Call(__e, tmp83250, tmp83252)

								tmp83254 := Call(__e, tmp83249, tmp83253)

								tmp83255 := Call(__e, tmp83248, tmp83254)

								tmp83256 := Call(__e, tmp83247, Nil, tmp83255)

								var ifres83223 Obj

								if True == tmp83256 {
									tmp83243 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp83244 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83245 := Call(__e, tmp83244, V1097)

									tmp83246 := Call(__e, tmp83243, tmp83245)

									var ifres83225 Obj

									if True == tmp83246 {
										tmp83237 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp83238 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83239 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83240 := Call(__e, tmp83239, V1097)

										tmp83241 := Call(__e, tmp83238, tmp83240)

										tmp83242 := Call(__e, tmp83237, Nil, tmp83241)

										var ifres83227 Obj

										if True == tmp83242 {
											tmp83229 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp83230 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp83231 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83232 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp83233 := Call(__e, tmp83232, V1097)

											tmp83234 := Call(__e, tmp83231, tmp83233)

											tmp83235 := Call(__e, tmp83230, tmp83234)

											tmp83236 := Call(__e, tmp83229, sym__, tmp83235)

											var ifres83228 Obj

											if True == tmp83236 {
												ifres83228 = True

											} else {
												ifres83228 = False

											}

											ifres83227 = ifres83228

										} else {
											ifres83227 = False

										}

										var ifres83226 Obj

										if True == ifres83227 {
											ifres83226 = True

										} else {
											ifres83226 = False

										}

										ifres83225 = ifres83226

									} else {
										ifres83225 = False

									}

									var ifres83224 Obj

									if True == ifres83225 {
										ifres83224 = True

									} else {
										ifres83224 = False

									}

									ifres83223 = ifres83224

								} else {
									ifres83223 = False

								}

								var ifres83222 Obj

								if True == ifres83223 {
									ifres83222 = True

								} else {
									ifres83222 = False

								}

								ifres83221 = ifres83222

							} else {
								ifres83221 = False

							}

							var ifres83220 Obj

							if True == ifres83221 {
								ifres83220 = True

							} else {
								ifres83220 = False

							}

							ifres83219 = ifres83220

						} else {
							ifres83219 = False

						}

						var ifres83218 Obj

						if True == ifres83219 {
							ifres83218 = True

						} else {
							ifres83218 = False

						}

						ifres83217 = ifres83218

					} else {
						ifres83217 = False

					}

					var ifres83216 Obj

					if True == ifres83217 {
						ifres83216 = True

					} else {
						ifres83216 = False

					}

					ifres83215 = ifres83216

				} else {
					ifres83215 = False

				}

				var ifres83214 Obj

				if True == ifres83215 {
					ifres83214 = True

				} else {
					ifres83214 = False

				}

				ifres83213 = ifres83214

			} else {
				ifres83213 = False

			}

			if True == ifres83213 {
				tmp83204 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

				tmp83205 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83206 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp83207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp83208 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83209 := Call(__e, tmp83208, V1097)

				tmp83210 := Call(__e, tmp83207, tmp83209)

				tmp83211 := Call(__e, tmp83206, tmp83210)

				tmp83212 := Call(__e, tmp83205, tmp83211)

				__e.TailApply(tmp83204, tmp83212, V1098)
				return

			} else {
				tmp83202 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp83203 := Call(__e, tmp83202, V1097)

				var ifres83130 Obj

				if True == tmp83203 {
					tmp83198 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp83199 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83200 := Call(__e, tmp83199, V1097)

					tmp83201 := Call(__e, tmp83198, tmp83200)

					var ifres83132 Obj

					if True == tmp83201 {
						tmp83192 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp83193 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83194 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83195 := Call(__e, tmp83194, V1097)

						tmp83196 := Call(__e, tmp83193, tmp83195)

						tmp83197 := Call(__e, tmp83192, symshen_4mu, tmp83196)

						var ifres83134 Obj

						if True == tmp83197 {
							tmp83186 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp83187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83188 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83189 := Call(__e, tmp83188, V1097)

							tmp83190 := Call(__e, tmp83187, tmp83189)

							tmp83191 := Call(__e, tmp83186, tmp83190)

							var ifres83136 Obj

							if True == tmp83191 {
								tmp83178 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp83179 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83180 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83181 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp83182 := Call(__e, tmp83181, V1097)

								tmp83183 := Call(__e, tmp83180, tmp83182)

								tmp83184 := Call(__e, tmp83179, tmp83183)

								tmp83185 := Call(__e, tmp83178, tmp83184)

								var ifres83138 Obj

								if True == tmp83185 {
									tmp83168 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp83169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83172 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp83173 := Call(__e, tmp83172, V1097)

									tmp83174 := Call(__e, tmp83171, tmp83173)

									tmp83175 := Call(__e, tmp83170, tmp83174)

									tmp83176 := Call(__e, tmp83169, tmp83175)

									tmp83177 := Call(__e, tmp83168, Nil, tmp83176)

									var ifres83140 Obj

									if True == tmp83177 {
										tmp83164 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp83165 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83166 := Call(__e, tmp83165, V1097)

										tmp83167 := Call(__e, tmp83164, tmp83166)

										var ifres83142 Obj

										if True == tmp83167 {
											tmp83158 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp83159 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83160 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83161 := Call(__e, tmp83160, V1097)

											tmp83162 := Call(__e, tmp83159, tmp83161)

											tmp83163 := Call(__e, tmp83158, Nil, tmp83162)

											var ifres83144 Obj

											if True == tmp83163 {
												tmp83146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ephemeral__variable_2)

												tmp83147 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp83148 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83149 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp83150 := Call(__e, tmp83149, V1097)

												tmp83151 := Call(__e, tmp83148, tmp83150)

												tmp83152 := Call(__e, tmp83147, tmp83151)

												tmp83153 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp83154 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83155 := Call(__e, tmp83154, V1097)

												tmp83156 := Call(__e, tmp83153, tmp83155)

												tmp83157 := Call(__e, tmp83146, tmp83152, tmp83156)

												var ifres83145 Obj

												if True == tmp83157 {
													ifres83145 = True

												} else {
													ifres83145 = False

												}

												ifres83144 = ifres83145

											} else {
												ifres83144 = False

											}

											var ifres83143 Obj

											if True == ifres83144 {
												ifres83143 = True

											} else {
												ifres83143 = False

											}

											ifres83142 = ifres83143

										} else {
											ifres83142 = False

										}

										var ifres83141 Obj

										if True == ifres83142 {
											ifres83141 = True

										} else {
											ifres83141 = False

										}

										ifres83140 = ifres83141

									} else {
										ifres83140 = False

									}

									var ifres83139 Obj

									if True == ifres83140 {
										ifres83139 = True

									} else {
										ifres83139 = False

									}

									ifres83138 = ifres83139

								} else {
									ifres83138 = False

								}

								var ifres83137 Obj

								if True == ifres83138 {
									ifres83137 = True

								} else {
									ifres83137 = False

								}

								ifres83136 = ifres83137

							} else {
								ifres83136 = False

							}

							var ifres83135 Obj

							if True == ifres83136 {
								ifres83135 = True

							} else {
								ifres83135 = False

							}

							ifres83134 = ifres83135

						} else {
							ifres83134 = False

						}

						var ifres83133 Obj

						if True == ifres83134 {
							ifres83133 = True

						} else {
							ifres83133 = False

						}

						ifres83132 = ifres83133

					} else {
						ifres83132 = False

					}

					var ifres83131 Obj

					if True == ifres83132 {
						ifres83131 = True

					} else {
						ifres83131 = False

					}

					ifres83130 = ifres83131

				} else {
					ifres83130 = False

				}

				if True == ifres83130 {
					tmp83109 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp83110 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83111 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83112 := Call(__e, tmp83111, V1097)

					tmp83113 := Call(__e, tmp83110, tmp83112)

					tmp83114 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83115 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83116 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83117 := Call(__e, tmp83116, V1097)

					tmp83118 := Call(__e, tmp83115, tmp83117)

					tmp83119 := Call(__e, tmp83114, tmp83118)

					tmp83120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

					tmp83121 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83124 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83125 := Call(__e, tmp83124, V1097)

					tmp83126 := Call(__e, tmp83123, tmp83125)

					tmp83127 := Call(__e, tmp83122, tmp83126)

					tmp83128 := Call(__e, tmp83121, tmp83127)

					tmp83129 := Call(__e, tmp83120, tmp83128, V1098)

					__e.TailApply(tmp83109, tmp83113, tmp83119, tmp83129)
					return

				} else {
					tmp83107 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp83108 := Call(__e, tmp83107, V1097)

					var ifres83039 Obj

					if True == tmp83108 {
						tmp83103 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp83104 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83105 := Call(__e, tmp83104, V1097)

						tmp83106 := Call(__e, tmp83103, tmp83105)

						var ifres83041 Obj

						if True == tmp83106 {
							tmp83097 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp83098 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83099 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83100 := Call(__e, tmp83099, V1097)

							tmp83101 := Call(__e, tmp83098, tmp83100)

							tmp83102 := Call(__e, tmp83097, symshen_4mu, tmp83101)

							var ifres83043 Obj

							if True == tmp83102 {
								tmp83091 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp83092 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83093 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp83094 := Call(__e, tmp83093, V1097)

								tmp83095 := Call(__e, tmp83092, tmp83094)

								tmp83096 := Call(__e, tmp83091, tmp83095)

								var ifres83045 Obj

								if True == tmp83096 {
									tmp83083 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp83084 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83085 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp83086 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp83087 := Call(__e, tmp83086, V1097)

									tmp83088 := Call(__e, tmp83085, tmp83087)

									tmp83089 := Call(__e, tmp83084, tmp83088)

									tmp83090 := Call(__e, tmp83083, tmp83089)

									var ifres83047 Obj

									if True == tmp83090 {
										tmp83073 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp83074 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83075 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83076 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83077 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp83078 := Call(__e, tmp83077, V1097)

										tmp83079 := Call(__e, tmp83076, tmp83078)

										tmp83080 := Call(__e, tmp83075, tmp83079)

										tmp83081 := Call(__e, tmp83074, tmp83080)

										tmp83082 := Call(__e, tmp83073, Nil, tmp83081)

										var ifres83049 Obj

										if True == tmp83082 {
											tmp83069 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp83070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83071 := Call(__e, tmp83070, V1097)

											tmp83072 := Call(__e, tmp83069, tmp83071)

											var ifres83051 Obj

											if True == tmp83072 {
												tmp83063 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp83064 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83065 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83066 := Call(__e, tmp83065, V1097)

												tmp83067 := Call(__e, tmp83064, tmp83066)

												tmp83068 := Call(__e, tmp83063, Nil, tmp83067)

												var ifres83053 Obj

												if True == tmp83068 {
													tmp83055 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

													tmp83056 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp83057 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83058 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp83059 := Call(__e, tmp83058, V1097)

													tmp83060 := Call(__e, tmp83057, tmp83059)

													tmp83061 := Call(__e, tmp83056, tmp83060)

													tmp83062 := Call(__e, tmp83055, tmp83061)

													var ifres83054 Obj

													if True == tmp83062 {
														ifres83054 = True

													} else {
														ifres83054 = False

													}

													ifres83053 = ifres83054

												} else {
													ifres83053 = False

												}

												var ifres83052 Obj

												if True == ifres83053 {
													ifres83052 = True

												} else {
													ifres83052 = False

												}

												ifres83051 = ifres83052

											} else {
												ifres83051 = False

											}

											var ifres83050 Obj

											if True == ifres83051 {
												ifres83050 = True

											} else {
												ifres83050 = False

											}

											ifres83049 = ifres83050

										} else {
											ifres83049 = False

										}

										var ifres83048 Obj

										if True == ifres83049 {
											ifres83048 = True

										} else {
											ifres83048 = False

										}

										ifres83047 = ifres83048

									} else {
										ifres83047 = False

									}

									var ifres83046 Obj

									if True == ifres83047 {
										ifres83046 = True

									} else {
										ifres83046 = False

									}

									ifres83045 = ifres83046

								} else {
									ifres83045 = False

								}

								var ifres83044 Obj

								if True == ifres83045 {
									ifres83044 = True

								} else {
									ifres83044 = False

								}

								ifres83043 = ifres83044

							} else {
								ifres83043 = False

							}

							var ifres83042 Obj

							if True == ifres83043 {
								ifres83042 = True

							} else {
								ifres83042 = False

							}

							ifres83041 = ifres83042

						} else {
							ifres83041 = False

						}

						var ifres83040 Obj

						if True == ifres83041 {
							ifres83040 = True

						} else {
							ifres83040 = False

						}

						ifres83039 = ifres83040

					} else {
						ifres83039 = False

					}

					if True == ifres83039 {
						tmp83008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp83009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp83010 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83011 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83012 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83013 := Call(__e, tmp83012, V1097)

						tmp83014 := Call(__e, tmp83011, tmp83013)

						tmp83015 := Call(__e, tmp83010, tmp83014)

						tmp83016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp83017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp83018 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83019 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83020 := Call(__e, tmp83019, V1097)

						tmp83021 := Call(__e, tmp83018, tmp83020)

						tmp83022 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp83023 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp83024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

						tmp83025 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83026 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83027 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83028 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83029 := Call(__e, tmp83028, V1097)

						tmp83030 := Call(__e, tmp83027, tmp83029)

						tmp83031 := Call(__e, tmp83026, tmp83030)

						tmp83032 := Call(__e, tmp83025, tmp83031)

						tmp83033 := Call(__e, tmp83024, tmp83032, V1098)

						tmp83034 := Call(__e, tmp83023, tmp83033, Nil)

						tmp83035 := Call(__e, tmp83022, symin, tmp83034)

						tmp83036 := Call(__e, tmp83017, tmp83021, tmp83035)

						tmp83037 := Call(__e, tmp83016, symshen_4be, tmp83036)

						tmp83038 := Call(__e, tmp83009, tmp83015, tmp83037)

						__e.TailApply(tmp83008, symlet, tmp83038)
						return

					} else {
						tmp83006 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp83007 := Call(__e, tmp83006, V1097)

						var ifres82934 Obj

						if True == tmp83007 {
							tmp83002 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp83003 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83004 := Call(__e, tmp83003, V1097)

							tmp83005 := Call(__e, tmp83002, tmp83004)

							var ifres82936 Obj

							if True == tmp83005 {
								tmp82996 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp82997 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82998 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82999 := Call(__e, tmp82998, V1097)

								tmp83000 := Call(__e, tmp82997, tmp82999)

								tmp83001 := Call(__e, tmp82996, symshen_4mu, tmp83000)

								var ifres82938 Obj

								if True == tmp83001 {
									tmp82990 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp82991 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82993 := Call(__e, tmp82992, V1097)

									tmp82994 := Call(__e, tmp82991, tmp82993)

									tmp82995 := Call(__e, tmp82990, tmp82994)

									var ifres82940 Obj

									if True == tmp82995 {
										tmp82982 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp82983 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82984 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82985 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82986 := Call(__e, tmp82985, V1097)

										tmp82987 := Call(__e, tmp82984, tmp82986)

										tmp82988 := Call(__e, tmp82983, tmp82987)

										tmp82989 := Call(__e, tmp82982, tmp82988)

										var ifres82942 Obj

										if True == tmp82989 {
											tmp82972 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp82973 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82974 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82975 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82976 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82977 := Call(__e, tmp82976, V1097)

											tmp82978 := Call(__e, tmp82975, tmp82977)

											tmp82979 := Call(__e, tmp82974, tmp82978)

											tmp82980 := Call(__e, tmp82973, tmp82979)

											tmp82981 := Call(__e, tmp82972, Nil, tmp82980)

											var ifres82944 Obj

											if True == tmp82981 {
												tmp82968 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp82969 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp82970 := Call(__e, tmp82969, V1097)

												tmp82971 := Call(__e, tmp82968, tmp82970)

												var ifres82946 Obj

												if True == tmp82971 {
													tmp82962 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp82963 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp82964 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp82965 := Call(__e, tmp82964, V1097)

													tmp82966 := Call(__e, tmp82963, tmp82965)

													tmp82967 := Call(__e, tmp82962, Nil, tmp82966)

													var ifres82948 Obj

													if True == tmp82967 {
														tmp82960 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp82961 := Call(__e, tmp82960, sym_1, V1098)

														var ifres82950 Obj

														if True == tmp82961 {
															tmp82952 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog__constant_2)

															tmp82953 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp82954 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp82955 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp82956 := Call(__e, tmp82955, V1097)

															tmp82957 := Call(__e, tmp82954, tmp82956)

															tmp82958 := Call(__e, tmp82953, tmp82957)

															tmp82959 := Call(__e, tmp82952, tmp82958)

															var ifres82951 Obj

															if True == tmp82959 {
																ifres82951 = True

															} else {
																ifres82951 = False

															}

															ifres82950 = ifres82951

														} else {
															ifres82950 = False

														}

														var ifres82949 Obj

														if True == ifres82950 {
															ifres82949 = True

														} else {
															ifres82949 = False

														}

														ifres82948 = ifres82949

													} else {
														ifres82948 = False

													}

													var ifres82947 Obj

													if True == ifres82948 {
														ifres82947 = True

													} else {
														ifres82947 = False

													}

													ifres82946 = ifres82947

												} else {
													ifres82946 = False

												}

												var ifres82945 Obj

												if True == ifres82946 {
													ifres82945 = True

												} else {
													ifres82945 = False

												}

												ifres82944 = ifres82945

											} else {
												ifres82944 = False

											}

											var ifres82943 Obj

											if True == ifres82944 {
												ifres82943 = True

											} else {
												ifres82943 = False

											}

											ifres82942 = ifres82943

										} else {
											ifres82942 = False

										}

										var ifres82941 Obj

										if True == ifres82942 {
											ifres82941 = True

										} else {
											ifres82941 = False

										}

										ifres82940 = ifres82941

									} else {
										ifres82940 = False

									}

									var ifres82939 Obj

									if True == ifres82940 {
										ifres82939 = True

									} else {
										ifres82939 = False

									}

									ifres82938 = ifres82939

								} else {
									ifres82938 = False

								}

								var ifres82937 Obj

								if True == ifres82938 {
									ifres82937 = True

								} else {
									ifres82937 = False

								}

								ifres82936 = ifres82937

							} else {
								ifres82936 = False

							}

							var ifres82935 Obj

							if True == ifres82936 {
								ifres82935 = True

							} else {
								ifres82935 = False

							}

							ifres82934 = ifres82935

						} else {
							ifres82934 = False

						}

						if True == ifres82934 {
							tmp82872 := MakeNative(func(__e *ControlFlow) {
								Z := __e.Get(1)
								_ = Z
								tmp82873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82874 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82876 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82877 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82878 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82879 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82880 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82881 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82882 := Call(__e, tmp82881, V1097)

								tmp82883 := Call(__e, tmp82880, symshen_4dereferencing, tmp82882)

								tmp82884 := Call(__e, tmp82879, symshen_4of, tmp82883)

								tmp82885 := Call(__e, tmp82878, symshen_4result, tmp82884)

								tmp82886 := Call(__e, tmp82877, symshen_4the, tmp82885)

								tmp82887 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82888 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82890 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82891 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82892 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82897 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82898 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82899 := Call(__e, tmp82898, V1097)

								tmp82900 := Call(__e, tmp82897, tmp82899)

								tmp82901 := Call(__e, tmp82896, tmp82900)

								tmp82902 := Call(__e, tmp82895, tmp82901, Nil)

								tmp82903 := Call(__e, tmp82894, symshen_4to, tmp82902)

								tmp82904 := Call(__e, tmp82893, symidentical, tmp82903)

								tmp82905 := Call(__e, tmp82892, symis, tmp82904)

								tmp82906 := Call(__e, tmp82891, Z, tmp82905)

								tmp82907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82909 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

								tmp82910 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82911 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82912 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp82913 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82914 := Call(__e, tmp82913, V1097)

								tmp82915 := Call(__e, tmp82912, tmp82914)

								tmp82916 := Call(__e, tmp82911, tmp82915)

								tmp82917 := Call(__e, tmp82910, tmp82916)

								tmp82918 := Call(__e, tmp82909, tmp82917, sym_1)

								tmp82919 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82920 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp82921 := Call(__e, tmp82920, symshen_4failed_b, Nil)

								tmp82922 := Call(__e, tmp82919, symshen_4else, tmp82921)

								tmp82923 := Call(__e, tmp82908, tmp82918, tmp82922)

								tmp82924 := Call(__e, tmp82907, symshen_4then, tmp82923)

								tmp82925 := Call(__e, tmp82890, tmp82906, tmp82924)

								tmp82926 := Call(__e, tmp82889, symif, tmp82925)

								tmp82927 := Call(__e, tmp82888, tmp82926, Nil)

								tmp82928 := Call(__e, tmp82887, symin, tmp82927)

								tmp82929 := Call(__e, tmp82876, tmp82886, tmp82928)

								tmp82930 := Call(__e, tmp82875, symshen_4be, tmp82929)

								tmp82931 := Call(__e, tmp82874, Z, tmp82930)

								__e.TailApply(tmp82873, symlet, tmp82931)
								return

							}, 1)

							tmp82932 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

							tmp82933 := Call(__e, tmp82932, symV)

							__e.TailApply(tmp82872, tmp82933)
							return

						} else {
							tmp82870 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp82871 := Call(__e, tmp82870, V1097)

							var ifres82798 Obj

							if True == tmp82871 {
								tmp82866 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp82867 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp82868 := Call(__e, tmp82867, V1097)

								tmp82869 := Call(__e, tmp82866, tmp82868)

								var ifres82800 Obj

								if True == tmp82869 {
									tmp82860 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp82861 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82862 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82863 := Call(__e, tmp82862, V1097)

									tmp82864 := Call(__e, tmp82861, tmp82863)

									tmp82865 := Call(__e, tmp82860, symshen_4mu, tmp82864)

									var ifres82802 Obj

									if True == tmp82865 {
										tmp82854 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp82855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82856 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82857 := Call(__e, tmp82856, V1097)

										tmp82858 := Call(__e, tmp82855, tmp82857)

										tmp82859 := Call(__e, tmp82854, tmp82858)

										var ifres82804 Obj

										if True == tmp82859 {
											tmp82846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp82847 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82848 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82849 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82850 := Call(__e, tmp82849, V1097)

											tmp82851 := Call(__e, tmp82848, tmp82850)

											tmp82852 := Call(__e, tmp82847, tmp82851)

											tmp82853 := Call(__e, tmp82846, tmp82852)

											var ifres82806 Obj

											if True == tmp82853 {
												tmp82836 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp82837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp82838 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp82839 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp82840 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp82841 := Call(__e, tmp82840, V1097)

												tmp82842 := Call(__e, tmp82839, tmp82841)

												tmp82843 := Call(__e, tmp82838, tmp82842)

												tmp82844 := Call(__e, tmp82837, tmp82843)

												tmp82845 := Call(__e, tmp82836, Nil, tmp82844)

												var ifres82808 Obj

												if True == tmp82845 {
													tmp82832 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp82833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp82834 := Call(__e, tmp82833, V1097)

													tmp82835 := Call(__e, tmp82832, tmp82834)

													var ifres82810 Obj

													if True == tmp82835 {
														tmp82826 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp82827 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82828 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82829 := Call(__e, tmp82828, V1097)

														tmp82830 := Call(__e, tmp82827, tmp82829)

														tmp82831 := Call(__e, tmp82826, Nil, tmp82830)

														var ifres82812 Obj

														if True == tmp82831 {
															tmp82824 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp82825 := Call(__e, tmp82824, sym_7, V1098)

															var ifres82814 Obj

															if True == tmp82825 {
																tmp82816 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog__constant_2)

																tmp82817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp82818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp82819 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp82820 := Call(__e, tmp82819, V1097)

																tmp82821 := Call(__e, tmp82818, tmp82820)

																tmp82822 := Call(__e, tmp82817, tmp82821)

																tmp82823 := Call(__e, tmp82816, tmp82822)

																var ifres82815 Obj

																if True == tmp82823 {
																	ifres82815 = True

																} else {
																	ifres82815 = False

																}

																ifres82814 = ifres82815

															} else {
																ifres82814 = False

															}

															var ifres82813 Obj

															if True == ifres82814 {
																ifres82813 = True

															} else {
																ifres82813 = False

															}

															ifres82812 = ifres82813

														} else {
															ifres82812 = False

														}

														var ifres82811 Obj

														if True == ifres82812 {
															ifres82811 = True

														} else {
															ifres82811 = False

														}

														ifres82810 = ifres82811

													} else {
														ifres82810 = False

													}

													var ifres82809 Obj

													if True == ifres82810 {
														ifres82809 = True

													} else {
														ifres82809 = False

													}

													ifres82808 = ifres82809

												} else {
													ifres82808 = False

												}

												var ifres82807 Obj

												if True == ifres82808 {
													ifres82807 = True

												} else {
													ifres82807 = False

												}

												ifres82806 = ifres82807

											} else {
												ifres82806 = False

											}

											var ifres82805 Obj

											if True == ifres82806 {
												ifres82805 = True

											} else {
												ifres82805 = False

											}

											ifres82804 = ifres82805

										} else {
											ifres82804 = False

										}

										var ifres82803 Obj

										if True == ifres82804 {
											ifres82803 = True

										} else {
											ifres82803 = False

										}

										ifres82802 = ifres82803

									} else {
										ifres82802 = False

									}

									var ifres82801 Obj

									if True == ifres82802 {
										ifres82801 = True

									} else {
										ifres82801 = False

									}

									ifres82800 = ifres82801

								} else {
									ifres82800 = False

								}

								var ifres82799 Obj

								if True == ifres82800 {
									ifres82799 = True

								} else {
									ifres82799 = False

								}

								ifres82798 = ifres82799

							} else {
								ifres82798 = False

							}

							if True == ifres82798 {
								tmp82688 := MakeNative(func(__e *ControlFlow) {
									Z := __e.Get(1)
									_ = Z
									tmp82689 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82690 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82693 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82694 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82695 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82696 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82697 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82698 := Call(__e, tmp82697, V1097)

									tmp82699 := Call(__e, tmp82696, symshen_4dereferencing, tmp82698)

									tmp82700 := Call(__e, tmp82695, symshen_4of, tmp82699)

									tmp82701 := Call(__e, tmp82694, symshen_4result, tmp82700)

									tmp82702 := Call(__e, tmp82693, symshen_4the, tmp82701)

									tmp82703 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82704 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82705 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82707 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82709 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82710 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82711 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82712 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82713 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82714 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82715 := Call(__e, tmp82714, V1097)

									tmp82716 := Call(__e, tmp82713, tmp82715)

									tmp82717 := Call(__e, tmp82712, tmp82716)

									tmp82718 := Call(__e, tmp82711, tmp82717, Nil)

									tmp82719 := Call(__e, tmp82710, symshen_4to, tmp82718)

									tmp82720 := Call(__e, tmp82709, symidentical, tmp82719)

									tmp82721 := Call(__e, tmp82708, symis, tmp82720)

									tmp82722 := Call(__e, tmp82707, Z, tmp82721)

									tmp82723 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82724 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82725 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

									tmp82726 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82727 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82728 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82729 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82730 := Call(__e, tmp82729, V1097)

									tmp82731 := Call(__e, tmp82728, tmp82730)

									tmp82732 := Call(__e, tmp82727, tmp82731)

									tmp82733 := Call(__e, tmp82726, tmp82732)

									tmp82734 := Call(__e, tmp82725, tmp82733, sym_7)

									tmp82735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82736 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82737 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82739 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82742 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82743 := Call(__e, tmp82742, symshen_4variable, Nil)

									tmp82744 := Call(__e, tmp82741, symshen_4a, tmp82743)

									tmp82745 := Call(__e, tmp82740, symis, tmp82744)

									tmp82746 := Call(__e, tmp82739, Z, tmp82745)

									tmp82747 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82748 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82749 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82750 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82751 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82752 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82753 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82754 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82755 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82756 := Call(__e, tmp82755, V1097)

									tmp82757 := Call(__e, tmp82754, tmp82756)

									tmp82758 := Call(__e, tmp82753, tmp82757)

									tmp82759 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82760 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

									tmp82762 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82764 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp82765 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82766 := Call(__e, tmp82765, V1097)

									tmp82767 := Call(__e, tmp82764, tmp82766)

									tmp82768 := Call(__e, tmp82763, tmp82767)

									tmp82769 := Call(__e, tmp82762, tmp82768)

									tmp82770 := Call(__e, tmp82761, tmp82769, sym_7)

									tmp82771 := Call(__e, tmp82760, tmp82770, Nil)

									tmp82772 := Call(__e, tmp82759, symin, tmp82771)

									tmp82773 := Call(__e, tmp82752, tmp82758, tmp82772)

									tmp82774 := Call(__e, tmp82751, symshen_4to, tmp82773)

									tmp82775 := Call(__e, tmp82750, Z, tmp82774)

									tmp82776 := Call(__e, tmp82749, symbind, tmp82775)

									tmp82777 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82778 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp82779 := Call(__e, tmp82778, symshen_4failed_b, Nil)

									tmp82780 := Call(__e, tmp82777, symshen_4else, tmp82779)

									tmp82781 := Call(__e, tmp82748, tmp82776, tmp82780)

									tmp82782 := Call(__e, tmp82747, symshen_4then, tmp82781)

									tmp82783 := Call(__e, tmp82738, tmp82746, tmp82782)

									tmp82784 := Call(__e, tmp82737, symif, tmp82783)

									tmp82785 := Call(__e, tmp82736, tmp82784, Nil)

									tmp82786 := Call(__e, tmp82735, symshen_4else, tmp82785)

									tmp82787 := Call(__e, tmp82724, tmp82734, tmp82786)

									tmp82788 := Call(__e, tmp82723, symshen_4then, tmp82787)

									tmp82789 := Call(__e, tmp82706, tmp82722, tmp82788)

									tmp82790 := Call(__e, tmp82705, symif, tmp82789)

									tmp82791 := Call(__e, tmp82704, tmp82790, Nil)

									tmp82792 := Call(__e, tmp82703, symin, tmp82791)

									tmp82793 := Call(__e, tmp82692, tmp82702, tmp82792)

									tmp82794 := Call(__e, tmp82691, symshen_4be, tmp82793)

									tmp82795 := Call(__e, tmp82690, Z, tmp82794)

									__e.TailApply(tmp82689, symlet, tmp82795)
									return

								}, 1)

								tmp82796 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

								tmp82797 := Call(__e, tmp82796, symV)

								__e.TailApply(tmp82688, tmp82797)
								return

							} else {
								tmp82686 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp82687 := Call(__e, tmp82686, V1097)

								var ifres82614 Obj

								if True == tmp82687 {
									tmp82682 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp82683 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp82684 := Call(__e, tmp82683, V1097)

									tmp82685 := Call(__e, tmp82682, tmp82684)

									var ifres82616 Obj

									if True == tmp82685 {
										tmp82676 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp82677 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82678 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82679 := Call(__e, tmp82678, V1097)

										tmp82680 := Call(__e, tmp82677, tmp82679)

										tmp82681 := Call(__e, tmp82676, symshen_4mu, tmp82680)

										var ifres82618 Obj

										if True == tmp82681 {
											tmp82670 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp82671 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82672 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82673 := Call(__e, tmp82672, V1097)

											tmp82674 := Call(__e, tmp82671, tmp82673)

											tmp82675 := Call(__e, tmp82670, tmp82674)

											var ifres82620 Obj

											if True == tmp82675 {
												tmp82662 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp82663 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp82664 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp82665 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp82666 := Call(__e, tmp82665, V1097)

												tmp82667 := Call(__e, tmp82664, tmp82666)

												tmp82668 := Call(__e, tmp82663, tmp82667)

												tmp82669 := Call(__e, tmp82662, tmp82668)

												var ifres82622 Obj

												if True == tmp82669 {
													tmp82654 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp82655 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp82656 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp82657 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp82658 := Call(__e, tmp82657, V1097)

													tmp82659 := Call(__e, tmp82656, tmp82658)

													tmp82660 := Call(__e, tmp82655, tmp82659)

													tmp82661 := Call(__e, tmp82654, tmp82660)

													var ifres82624 Obj

													if True == tmp82661 {
														tmp82644 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp82645 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82646 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82648 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp82649 := Call(__e, tmp82648, V1097)

														tmp82650 := Call(__e, tmp82647, tmp82649)

														tmp82651 := Call(__e, tmp82646, tmp82650)

														tmp82652 := Call(__e, tmp82645, tmp82651)

														tmp82653 := Call(__e, tmp82644, Nil, tmp82652)

														var ifres82626 Obj

														if True == tmp82653 {
															tmp82640 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp82641 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp82642 := Call(__e, tmp82641, V1097)

															tmp82643 := Call(__e, tmp82640, tmp82642)

															var ifres82628 Obj

															if True == tmp82643 {
																tmp82634 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp82635 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp82636 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp82637 := Call(__e, tmp82636, V1097)

																tmp82638 := Call(__e, tmp82635, tmp82637)

																tmp82639 := Call(__e, tmp82634, Nil, tmp82638)

																var ifres82630 Obj

																if True == tmp82639 {
																	tmp82632 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp82633 := Call(__e, tmp82632, sym_1, V1098)

																	var ifres82631 Obj

																	if True == tmp82633 {
																		ifres82631 = True

																	} else {
																		ifres82631 = False

																	}

																	ifres82630 = ifres82631

																} else {
																	ifres82630 = False

																}

																var ifres82629 Obj

																if True == ifres82630 {
																	ifres82629 = True

																} else {
																	ifres82629 = False

																}

																ifres82628 = ifres82629

															} else {
																ifres82628 = False

															}

															var ifres82627 Obj

															if True == ifres82628 {
																ifres82627 = True

															} else {
																ifres82627 = False

															}

															ifres82626 = ifres82627

														} else {
															ifres82626 = False

														}

														var ifres82625 Obj

														if True == ifres82626 {
															ifres82625 = True

														} else {
															ifres82625 = False

														}

														ifres82624 = ifres82625

													} else {
														ifres82624 = False

													}

													var ifres82623 Obj

													if True == ifres82624 {
														ifres82623 = True

													} else {
														ifres82623 = False

													}

													ifres82622 = ifres82623

												} else {
													ifres82622 = False

												}

												var ifres82621 Obj

												if True == ifres82622 {
													ifres82621 = True

												} else {
													ifres82621 = False

												}

												ifres82620 = ifres82621

											} else {
												ifres82620 = False

											}

											var ifres82619 Obj

											if True == ifres82620 {
												ifres82619 = True

											} else {
												ifres82619 = False

											}

											ifres82618 = ifres82619

										} else {
											ifres82618 = False

										}

										var ifres82617 Obj

										if True == ifres82618 {
											ifres82617 = True

										} else {
											ifres82617 = False

										}

										ifres82616 = ifres82617

									} else {
										ifres82616 = False

									}

									var ifres82615 Obj

									if True == ifres82616 {
										ifres82615 = True

									} else {
										ifres82615 = False

									}

									ifres82614 = ifres82615

								} else {
									ifres82614 = False

								}

								if True == ifres82614 {
									tmp82510 := MakeNative(func(__e *ControlFlow) {
										Z := __e.Get(1)
										_ = Z
										tmp82511 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82512 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82513 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82514 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82515 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82516 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82517 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82518 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82519 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82520 := Call(__e, tmp82519, V1097)

										tmp82521 := Call(__e, tmp82518, symshen_4dereferencing, tmp82520)

										tmp82522 := Call(__e, tmp82517, symshen_4of, tmp82521)

										tmp82523 := Call(__e, tmp82516, symshen_4result, tmp82522)

										tmp82524 := Call(__e, tmp82515, symshen_4the, tmp82523)

										tmp82525 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82526 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82527 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82528 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82529 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82530 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82531 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82532 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82533 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82534 := Call(__e, tmp82533, symlist, Nil)

										tmp82535 := Call(__e, tmp82532, symshen_4non_1empty, tmp82534)

										tmp82536 := Call(__e, tmp82531, symshen_4a, tmp82535)

										tmp82537 := Call(__e, tmp82530, symis, tmp82536)

										tmp82538 := Call(__e, tmp82529, Z, tmp82537)

										tmp82539 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82541 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

										tmp82542 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82543 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82544 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82545 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82546 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82547 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82548 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82549 := Call(__e, tmp82548, V1097)

										tmp82550 := Call(__e, tmp82547, tmp82549)

										tmp82551 := Call(__e, tmp82546, tmp82550)

										tmp82552 := Call(__e, tmp82545, tmp82551)

										tmp82553 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82554 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82555 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82556 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82558 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82559 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82560 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82561 := Call(__e, tmp82560, V1097)

										tmp82562 := Call(__e, tmp82559, tmp82561)

										tmp82563 := Call(__e, tmp82558, tmp82562)

										tmp82564 := Call(__e, tmp82557, tmp82563)

										tmp82565 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82566 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp82567 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82568 := Call(__e, tmp82567, V1097)

										tmp82569 := Call(__e, tmp82566, tmp82568)

										tmp82570 := Call(__e, tmp82565, tmp82569)

										tmp82571 := Call(__e, tmp82556, tmp82564, tmp82570)

										tmp82572 := Call(__e, tmp82555, symshen_4mu, tmp82571)

										tmp82573 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82574 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82575 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82576 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82577 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82578 := Call(__e, tmp82577, Z, Nil)

										tmp82579 := Call(__e, tmp82576, symshen_4of, tmp82578)

										tmp82580 := Call(__e, tmp82575, symtail, tmp82579)

										tmp82581 := Call(__e, tmp82574, symshen_4the, tmp82580)

										tmp82582 := Call(__e, tmp82573, tmp82581, Nil)

										tmp82583 := Call(__e, tmp82554, tmp82572, tmp82582)

										tmp82584 := Call(__e, tmp82553, tmp82583, Nil)

										tmp82585 := Call(__e, tmp82544, tmp82552, tmp82584)

										tmp82586 := Call(__e, tmp82543, symshen_4mu, tmp82585)

										tmp82587 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82588 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82589 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82590 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82591 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82592 := Call(__e, tmp82591, Z, Nil)

										tmp82593 := Call(__e, tmp82590, symshen_4of, tmp82592)

										tmp82594 := Call(__e, tmp82589, symhead, tmp82593)

										tmp82595 := Call(__e, tmp82588, symshen_4the, tmp82594)

										tmp82596 := Call(__e, tmp82587, tmp82595, Nil)

										tmp82597 := Call(__e, tmp82542, tmp82586, tmp82596)

										tmp82598 := Call(__e, tmp82541, tmp82597, sym_1)

										tmp82599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp82601 := Call(__e, tmp82600, symshen_4failed_b, Nil)

										tmp82602 := Call(__e, tmp82599, symshen_4else, tmp82601)

										tmp82603 := Call(__e, tmp82540, tmp82598, tmp82602)

										tmp82604 := Call(__e, tmp82539, symshen_4then, tmp82603)

										tmp82605 := Call(__e, tmp82528, tmp82538, tmp82604)

										tmp82606 := Call(__e, tmp82527, symif, tmp82605)

										tmp82607 := Call(__e, tmp82526, tmp82606, Nil)

										tmp82608 := Call(__e, tmp82525, symin, tmp82607)

										tmp82609 := Call(__e, tmp82514, tmp82524, tmp82608)

										tmp82610 := Call(__e, tmp82513, symshen_4be, tmp82609)

										tmp82611 := Call(__e, tmp82512, Z, tmp82610)

										__e.TailApply(tmp82511, symlet, tmp82611)
										return

									}, 1)

									tmp82612 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

									tmp82613 := Call(__e, tmp82612, symV)

									__e.TailApply(tmp82510, tmp82613)
									return

								} else {
									tmp82508 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp82509 := Call(__e, tmp82508, V1097)

									var ifres82436 Obj

									if True == tmp82509 {
										tmp82504 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp82505 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp82506 := Call(__e, tmp82505, V1097)

										tmp82507 := Call(__e, tmp82504, tmp82506)

										var ifres82438 Obj

										if True == tmp82507 {
											tmp82498 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp82499 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82500 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82501 := Call(__e, tmp82500, V1097)

											tmp82502 := Call(__e, tmp82499, tmp82501)

											tmp82503 := Call(__e, tmp82498, symshen_4mu, tmp82502)

											var ifres82440 Obj

											if True == tmp82503 {
												tmp82492 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp82493 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp82494 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp82495 := Call(__e, tmp82494, V1097)

												tmp82496 := Call(__e, tmp82493, tmp82495)

												tmp82497 := Call(__e, tmp82492, tmp82496)

												var ifres82442 Obj

												if True == tmp82497 {
													tmp82484 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp82485 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp82486 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp82487 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp82488 := Call(__e, tmp82487, V1097)

													tmp82489 := Call(__e, tmp82486, tmp82488)

													tmp82490 := Call(__e, tmp82485, tmp82489)

													tmp82491 := Call(__e, tmp82484, tmp82490)

													var ifres82444 Obj

													if True == tmp82491 {
														tmp82476 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp82477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82478 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp82479 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp82480 := Call(__e, tmp82479, V1097)

														tmp82481 := Call(__e, tmp82478, tmp82480)

														tmp82482 := Call(__e, tmp82477, tmp82481)

														tmp82483 := Call(__e, tmp82476, tmp82482)

														var ifres82446 Obj

														if True == tmp82483 {
															tmp82466 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp82467 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp82468 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp82469 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp82470 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp82471 := Call(__e, tmp82470, V1097)

															tmp82472 := Call(__e, tmp82469, tmp82471)

															tmp82473 := Call(__e, tmp82468, tmp82472)

															tmp82474 := Call(__e, tmp82467, tmp82473)

															tmp82475 := Call(__e, tmp82466, Nil, tmp82474)

															var ifres82448 Obj

															if True == tmp82475 {
																tmp82462 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp82463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp82464 := Call(__e, tmp82463, V1097)

																tmp82465 := Call(__e, tmp82462, tmp82464)

																var ifres82450 Obj

																if True == tmp82465 {
																	tmp82456 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp82457 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp82458 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp82459 := Call(__e, tmp82458, V1097)

																	tmp82460 := Call(__e, tmp82457, tmp82459)

																	tmp82461 := Call(__e, tmp82456, Nil, tmp82460)

																	var ifres82452 Obj

																	if True == tmp82461 {
																		tmp82454 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp82455 := Call(__e, tmp82454, sym_7, V1098)

																		var ifres82453 Obj

																		if True == tmp82455 {
																			ifres82453 = True

																		} else {
																			ifres82453 = False

																		}

																		ifres82452 = ifres82453

																	} else {
																		ifres82452 = False

																	}

																	var ifres82451 Obj

																	if True == ifres82452 {
																		ifres82451 = True

																	} else {
																		ifres82451 = False

																	}

																	ifres82450 = ifres82451

																} else {
																	ifres82450 = False

																}

																var ifres82449 Obj

																if True == ifres82450 {
																	ifres82449 = True

																} else {
																	ifres82449 = False

																}

																ifres82448 = ifres82449

															} else {
																ifres82448 = False

															}

															var ifres82447 Obj

															if True == ifres82448 {
																ifres82447 = True

															} else {
																ifres82447 = False

															}

															ifres82446 = ifres82447

														} else {
															ifres82446 = False

														}

														var ifres82445 Obj

														if True == ifres82446 {
															ifres82445 = True

														} else {
															ifres82445 = False

														}

														ifres82444 = ifres82445

													} else {
														ifres82444 = False

													}

													var ifres82443 Obj

													if True == ifres82444 {
														ifres82443 = True

													} else {
														ifres82443 = False

													}

													ifres82442 = ifres82443

												} else {
													ifres82442 = False

												}

												var ifres82441 Obj

												if True == ifres82442 {
													ifres82441 = True

												} else {
													ifres82441 = False

												}

												ifres82440 = ifres82441

											} else {
												ifres82440 = False

											}

											var ifres82439 Obj

											if True == ifres82440 {
												ifres82439 = True

											} else {
												ifres82439 = False

											}

											ifres82438 = ifres82439

										} else {
											ifres82438 = False

										}

										var ifres82437 Obj

										if True == ifres82438 {
											ifres82437 = True

										} else {
											ifres82437 = False

										}

										ifres82436 = ifres82437

									} else {
										ifres82436 = False

									}

									if True == ifres82436 {
										tmp82256 := MakeNative(func(__e *ControlFlow) {
											Z := __e.Get(1)
											_ = Z
											tmp82257 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82258 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82261 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82265 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82266 := Call(__e, tmp82265, V1097)

											tmp82267 := Call(__e, tmp82264, symshen_4dereferencing, tmp82266)

											tmp82268 := Call(__e, tmp82263, symshen_4of, tmp82267)

											tmp82269 := Call(__e, tmp82262, symshen_4result, tmp82268)

											tmp82270 := Call(__e, tmp82261, symshen_4the, tmp82269)

											tmp82271 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82272 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82273 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82274 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82275 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82276 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82279 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82280 := Call(__e, tmp82279, symlist, Nil)

											tmp82281 := Call(__e, tmp82278, symshen_4non_1empty, tmp82280)

											tmp82282 := Call(__e, tmp82277, symshen_4a, tmp82281)

											tmp82283 := Call(__e, tmp82276, symis, tmp82282)

											tmp82284 := Call(__e, tmp82275, Z, tmp82283)

											tmp82285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82287 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

											tmp82288 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82289 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82290 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82291 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82292 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82293 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82294 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82295 := Call(__e, tmp82294, V1097)

											tmp82296 := Call(__e, tmp82293, tmp82295)

											tmp82297 := Call(__e, tmp82292, tmp82296)

											tmp82298 := Call(__e, tmp82291, tmp82297)

											tmp82299 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82300 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82301 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82302 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82303 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82304 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82305 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82306 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82307 := Call(__e, tmp82306, V1097)

											tmp82308 := Call(__e, tmp82305, tmp82307)

											tmp82309 := Call(__e, tmp82304, tmp82308)

											tmp82310 := Call(__e, tmp82303, tmp82309)

											tmp82311 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82312 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82313 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82314 := Call(__e, tmp82313, V1097)

											tmp82315 := Call(__e, tmp82312, tmp82314)

											tmp82316 := Call(__e, tmp82311, tmp82315)

											tmp82317 := Call(__e, tmp82302, tmp82310, tmp82316)

											tmp82318 := Call(__e, tmp82301, symshen_4mu, tmp82317)

											tmp82319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82324 := Call(__e, tmp82323, Z, Nil)

											tmp82325 := Call(__e, tmp82322, symshen_4of, tmp82324)

											tmp82326 := Call(__e, tmp82321, symtail, tmp82325)

											tmp82327 := Call(__e, tmp82320, symshen_4the, tmp82326)

											tmp82328 := Call(__e, tmp82319, tmp82327, Nil)

											tmp82329 := Call(__e, tmp82300, tmp82318, tmp82328)

											tmp82330 := Call(__e, tmp82299, tmp82329, Nil)

											tmp82331 := Call(__e, tmp82290, tmp82298, tmp82330)

											tmp82332 := Call(__e, tmp82289, symshen_4mu, tmp82331)

											tmp82333 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82338 := Call(__e, tmp82337, Z, Nil)

											tmp82339 := Call(__e, tmp82336, symshen_4of, tmp82338)

											tmp82340 := Call(__e, tmp82335, symhead, tmp82339)

											tmp82341 := Call(__e, tmp82334, symshen_4the, tmp82340)

											tmp82342 := Call(__e, tmp82333, tmp82341, Nil)

											tmp82343 := Call(__e, tmp82288, tmp82332, tmp82342)

											tmp82344 := Call(__e, tmp82287, tmp82343, sym_7)

											tmp82345 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82349 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82350 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82351 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82353 := Call(__e, tmp82352, symshen_4variable, Nil)

											tmp82354 := Call(__e, tmp82351, symshen_4a, tmp82353)

											tmp82355 := Call(__e, tmp82350, symis, tmp82354)

											tmp82356 := Call(__e, tmp82349, Z, tmp82355)

											tmp82357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82359 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82360 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82361 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82362 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82363 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

											tmp82365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82366 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82367 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82368 := Call(__e, tmp82367, V1097)

											tmp82369 := Call(__e, tmp82366, tmp82368)

											tmp82370 := Call(__e, tmp82365, tmp82369)

											tmp82371 := Call(__e, tmp82364, tmp82370)

											tmp82372 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82373 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82376 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82377 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

											tmp82380 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove__modes)

											tmp82381 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82382 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82383 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82384 := Call(__e, tmp82383, V1097)

											tmp82385 := Call(__e, tmp82382, tmp82384)

											tmp82386 := Call(__e, tmp82381, tmp82385)

											tmp82387 := Call(__e, tmp82380, tmp82386)

											tmp82388 := Call(__e, tmp82379, tmp82387)

											tmp82389 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82391 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mu__reduction)

											tmp82392 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82393 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82394 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp82395 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp82396 := Call(__e, tmp82395, V1097)

											tmp82397 := Call(__e, tmp82394, tmp82396)

											tmp82398 := Call(__e, tmp82393, tmp82397)

											tmp82399 := Call(__e, tmp82392, tmp82398)

											tmp82400 := Call(__e, tmp82391, tmp82399, sym_7)

											tmp82401 := Call(__e, tmp82390, tmp82400, Nil)

											tmp82402 := Call(__e, tmp82389, symin, tmp82401)

											tmp82403 := Call(__e, tmp82378, tmp82388, tmp82402)

											tmp82404 := Call(__e, tmp82377, symshen_4to, tmp82403)

											tmp82405 := Call(__e, tmp82376, Z, tmp82404)

											tmp82406 := Call(__e, tmp82375, symbind, tmp82405)

											tmp82407 := Call(__e, tmp82374, tmp82406, Nil)

											tmp82408 := Call(__e, tmp82373, symshen_4then, tmp82407)

											tmp82409 := Call(__e, tmp82372, symand, tmp82408)

											tmp82410 := Call(__e, tmp82363, tmp82371, tmp82409)

											tmp82411 := Call(__e, tmp82362, symin, tmp82410)

											tmp82412 := Call(__e, tmp82361, symshen_4variables, tmp82411)

											tmp82413 := Call(__e, tmp82360, symshen_4the, tmp82412)

											tmp82414 := Call(__e, tmp82359, symshen_4rename, tmp82413)

											tmp82415 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82416 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp82417 := Call(__e, tmp82416, symshen_4failed_b, Nil)

											tmp82418 := Call(__e, tmp82415, symshen_4else, tmp82417)

											tmp82419 := Call(__e, tmp82358, tmp82414, tmp82418)

											tmp82420 := Call(__e, tmp82357, symshen_4then, tmp82419)

											tmp82421 := Call(__e, tmp82348, tmp82356, tmp82420)

											tmp82422 := Call(__e, tmp82347, symif, tmp82421)

											tmp82423 := Call(__e, tmp82346, tmp82422, Nil)

											tmp82424 := Call(__e, tmp82345, symshen_4else, tmp82423)

											tmp82425 := Call(__e, tmp82286, tmp82344, tmp82424)

											tmp82426 := Call(__e, tmp82285, symshen_4then, tmp82425)

											tmp82427 := Call(__e, tmp82274, tmp82284, tmp82426)

											tmp82428 := Call(__e, tmp82273, symif, tmp82427)

											tmp82429 := Call(__e, tmp82272, tmp82428, Nil)

											tmp82430 := Call(__e, tmp82271, symin, tmp82429)

											tmp82431 := Call(__e, tmp82260, tmp82270, tmp82430)

											tmp82432 := Call(__e, tmp82259, symshen_4be, tmp82431)

											tmp82433 := Call(__e, tmp82258, Z, tmp82432)

											__e.TailApply(tmp82257, symlet, tmp82433)
											return

										}, 1)

										tmp82434 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

										tmp82435 := Call(__e, tmp82434, symV)

										__e.TailApply(tmp82256, tmp82435)
										return

									} else {
										__e.Return(V1097)
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

	tmp83444 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mu__reduction, tmp82247)

	_ = tmp83444

	tmp83445 := MakeNative(func(__e *ControlFlow) {
		V1100 := __e.Get(1)
		_ = V1100
		tmp83460 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp83461 := Call(__e, tmp83460, V1100)

		if True == tmp83461 {
			tmp83447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp83448 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp83449 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp83450 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83451 := Call(__e, tmp83450, V1100)

			tmp83452 := Call(__e, tmp83449, tmp83451)

			tmp83453 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp83454 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp83455 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83456 := Call(__e, tmp83455, V1100)

			tmp83457 := Call(__e, tmp83454, tmp83456)

			tmp83458 := Call(__e, tmp83453, tmp83457, Nil)

			tmp83459 := Call(__e, tmp83448, tmp83452, tmp83458)

			__e.TailApply(tmp83447, symcons, tmp83459)
			return

		} else {
			__e.Return(V1100)
			return
		}

	}, 1)

	tmp83462 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rcons__form, tmp83445)

	_ = tmp83462

	tmp83463 := MakeNative(func(__e *ControlFlow) {
		V1102 := __e.Get(1)
		_ = V1102
		tmp83570 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp83571 := Call(__e, tmp83570, V1102)

		var ifres83530 Obj

		if True == tmp83571 {
			tmp83566 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp83567 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83568 := Call(__e, tmp83567, V1102)

			tmp83569 := Call(__e, tmp83566, symmode, tmp83568)

			var ifres83532 Obj

			if True == tmp83569 {
				tmp83562 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp83563 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp83564 := Call(__e, tmp83563, V1102)

				tmp83565 := Call(__e, tmp83562, tmp83564)

				var ifres83534 Obj

				if True == tmp83565 {
					tmp83556 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp83557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83558 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83559 := Call(__e, tmp83558, V1102)

					tmp83560 := Call(__e, tmp83557, tmp83559)

					tmp83561 := Call(__e, tmp83556, tmp83560)

					var ifres83536 Obj

					if True == tmp83561 {
						tmp83548 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp83549 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp83550 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83551 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83552 := Call(__e, tmp83551, V1102)

						tmp83553 := Call(__e, tmp83550, tmp83552)

						tmp83554 := Call(__e, tmp83549, tmp83553)

						tmp83555 := Call(__e, tmp83548, sym_7, tmp83554)

						var ifres83538 Obj

						if True == tmp83555 {
							tmp83540 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp83541 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83542 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83543 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83544 := Call(__e, tmp83543, V1102)

							tmp83545 := Call(__e, tmp83542, tmp83544)

							tmp83546 := Call(__e, tmp83541, tmp83545)

							tmp83547 := Call(__e, tmp83540, Nil, tmp83546)

							var ifres83539 Obj

							if True == tmp83547 {
								ifres83539 = True

							} else {
								ifres83539 = False

							}

							ifres83538 = ifres83539

						} else {
							ifres83538 = False

						}

						var ifres83537 Obj

						if True == ifres83538 {
							ifres83537 = True

						} else {
							ifres83537 = False

						}

						ifres83536 = ifres83537

					} else {
						ifres83536 = False

					}

					var ifres83535 Obj

					if True == ifres83536 {
						ifres83535 = True

					} else {
						ifres83535 = False

					}

					ifres83534 = ifres83535

				} else {
					ifres83534 = False

				}

				var ifres83533 Obj

				if True == ifres83534 {
					ifres83533 = True

				} else {
					ifres83533 = False

				}

				ifres83532 = ifres83533

			} else {
				ifres83532 = False

			}

			var ifres83531 Obj

			if True == ifres83532 {
				ifres83531 = True

			} else {
				ifres83531 = False

			}

			ifres83530 = ifres83531

		} else {
			ifres83530 = False

		}

		if True == ifres83530 {
			tmp83525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove__modes)

			tmp83526 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp83527 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp83528 := Call(__e, tmp83527, V1102)

			tmp83529 := Call(__e, tmp83526, tmp83528)

			__e.TailApply(tmp83525, tmp83529)
			return

		} else {
			tmp83523 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp83524 := Call(__e, tmp83523, V1102)

			var ifres83483 Obj

			if True == tmp83524 {
				tmp83519 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp83520 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83521 := Call(__e, tmp83520, V1102)

				tmp83522 := Call(__e, tmp83519, symmode, tmp83521)

				var ifres83485 Obj

				if True == tmp83522 {
					tmp83515 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp83516 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83517 := Call(__e, tmp83516, V1102)

					tmp83518 := Call(__e, tmp83515, tmp83517)

					var ifres83487 Obj

					if True == tmp83518 {
						tmp83509 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp83510 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83511 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp83512 := Call(__e, tmp83511, V1102)

						tmp83513 := Call(__e, tmp83510, tmp83512)

						tmp83514 := Call(__e, tmp83509, tmp83513)

						var ifres83489 Obj

						if True == tmp83514 {
							tmp83501 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp83502 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp83503 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83504 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp83505 := Call(__e, tmp83504, V1102)

							tmp83506 := Call(__e, tmp83503, tmp83505)

							tmp83507 := Call(__e, tmp83502, tmp83506)

							tmp83508 := Call(__e, tmp83501, sym_1, tmp83507)

							var ifres83491 Obj

							if True == tmp83508 {
								tmp83493 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp83494 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83495 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83496 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp83497 := Call(__e, tmp83496, V1102)

								tmp83498 := Call(__e, tmp83495, tmp83497)

								tmp83499 := Call(__e, tmp83494, tmp83498)

								tmp83500 := Call(__e, tmp83493, Nil, tmp83499)

								var ifres83492 Obj

								if True == tmp83500 {
									ifres83492 = True

								} else {
									ifres83492 = False

								}

								ifres83491 = ifres83492

							} else {
								ifres83491 = False

							}

							var ifres83490 Obj

							if True == ifres83491 {
								ifres83490 = True

							} else {
								ifres83490 = False

							}

							ifres83489 = ifres83490

						} else {
							ifres83489 = False

						}

						var ifres83488 Obj

						if True == ifres83489 {
							ifres83488 = True

						} else {
							ifres83488 = False

						}

						ifres83487 = ifres83488

					} else {
						ifres83487 = False

					}

					var ifres83486 Obj

					if True == ifres83487 {
						ifres83486 = True

					} else {
						ifres83486 = False

					}

					ifres83485 = ifres83486

				} else {
					ifres83485 = False

				}

				var ifres83484 Obj

				if True == ifres83485 {
					ifres83484 = True

				} else {
					ifres83484 = False

				}

				ifres83483 = ifres83484

			} else {
				ifres83483 = False

			}

			if True == ifres83483 {
				tmp83478 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove__modes)

				tmp83479 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp83480 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp83481 := Call(__e, tmp83480, V1102)

				tmp83482 := Call(__e, tmp83479, tmp83481)

				__e.TailApply(tmp83478, tmp83482)
				return

			} else {
				tmp83476 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp83477 := Call(__e, tmp83476, V1102)

				if True == tmp83477 {
					tmp83467 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp83468 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove__modes)

					tmp83469 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp83470 := Call(__e, tmp83469, V1102)

					tmp83471 := Call(__e, tmp83468, tmp83470)

					tmp83472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove__modes)

					tmp83473 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp83474 := Call(__e, tmp83473, V1102)

					tmp83475 := Call(__e, tmp83472, tmp83474)

					__e.TailApply(tmp83467, tmp83471, tmp83475)
					return

				} else {
					__e.Return(V1102)
					return
				}

			}

		}

	}, 1)

	tmp83572 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove__modes, tmp83463)

	_ = tmp83572

	tmp83573 := MakeNative(func(__e *ControlFlow) {
		V1105 := __e.Get(1)
		_ = V1105
		V1106 := __e.Get(2)
		_ = V1106
		tmp83578 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

		tmp83579 := Call(__e, tmp83578, V1105)

		if True == tmp83579 {
			tmp83576 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			tmp83577 := Call(__e, tmp83576, V1106)

			if True == tmp83577 {
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

	}, 2)

	tmp83580 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ephemeral__variable_2, tmp83573)

	_ = tmp83580

	tmp83581 := MakeNative(func(__e *ControlFlow) {
		V1116 := __e.Get(1)
		_ = V1116
		tmp83583 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp83584 := Call(__e, tmp83583, V1116)

		if True == tmp83584 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp83585 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog__constant_2, tmp83581)

	_ = tmp83585

	tmp83586 := MakeNative(func(__e *ControlFlow) {
		V1118 := __e.Get(1)
		_ = V1118
		tmp85144 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85145 := Call(__e, tmp85144, V1118)

		var ifres85048 Obj

		if True == tmp85145 {
			tmp85140 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp85141 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85142 := Call(__e, tmp85141, V1118)

			tmp85143 := Call(__e, tmp85140, symlet, tmp85142)

			var ifres85050 Obj

			if True == tmp85143 {
				tmp85136 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85137 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85138 := Call(__e, tmp85137, V1118)

				tmp85139 := Call(__e, tmp85136, tmp85138)

				var ifres85052 Obj

				if True == tmp85139 {
					tmp85130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp85131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85132 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85133 := Call(__e, tmp85132, V1118)

					tmp85134 := Call(__e, tmp85131, tmp85133)

					tmp85135 := Call(__e, tmp85130, tmp85134)

					var ifres85054 Obj

					if True == tmp85135 {
						tmp85122 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp85123 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85124 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp85125 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp85126 := Call(__e, tmp85125, V1118)

						tmp85127 := Call(__e, tmp85124, tmp85126)

						tmp85128 := Call(__e, tmp85123, tmp85127)

						tmp85129 := Call(__e, tmp85122, symshen_4be, tmp85128)

						var ifres85056 Obj

						if True == tmp85129 {
							tmp85114 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp85115 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85116 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85117 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85118 := Call(__e, tmp85117, V1118)

							tmp85119 := Call(__e, tmp85116, tmp85118)

							tmp85120 := Call(__e, tmp85115, tmp85119)

							tmp85121 := Call(__e, tmp85114, tmp85120)

							var ifres85058 Obj

							if True == tmp85121 {
								tmp85104 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp85105 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp85106 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp85107 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp85108 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp85109 := Call(__e, tmp85108, V1118)

								tmp85110 := Call(__e, tmp85107, tmp85109)

								tmp85111 := Call(__e, tmp85106, tmp85110)

								tmp85112 := Call(__e, tmp85105, tmp85111)

								tmp85113 := Call(__e, tmp85104, tmp85112)

								var ifres85060 Obj

								if True == tmp85113 {
									tmp85092 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp85093 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp85094 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp85095 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp85096 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp85097 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp85098 := Call(__e, tmp85097, V1118)

									tmp85099 := Call(__e, tmp85096, tmp85098)

									tmp85100 := Call(__e, tmp85095, tmp85099)

									tmp85101 := Call(__e, tmp85094, tmp85100)

									tmp85102 := Call(__e, tmp85093, tmp85101)

									tmp85103 := Call(__e, tmp85092, symin, tmp85102)

									var ifres85062 Obj

									if True == tmp85103 {
										tmp85080 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp85081 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp85082 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp85083 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp85084 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp85085 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp85086 := Call(__e, tmp85085, V1118)

										tmp85087 := Call(__e, tmp85084, tmp85086)

										tmp85088 := Call(__e, tmp85083, tmp85087)

										tmp85089 := Call(__e, tmp85082, tmp85088)

										tmp85090 := Call(__e, tmp85081, tmp85089)

										tmp85091 := Call(__e, tmp85080, tmp85090)

										var ifres85064 Obj

										if True == tmp85091 {
											tmp85066 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp85067 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp85068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp85069 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp85070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp85071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp85072 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp85073 := Call(__e, tmp85072, V1118)

											tmp85074 := Call(__e, tmp85071, tmp85073)

											tmp85075 := Call(__e, tmp85070, tmp85074)

											tmp85076 := Call(__e, tmp85069, tmp85075)

											tmp85077 := Call(__e, tmp85068, tmp85076)

											tmp85078 := Call(__e, tmp85067, tmp85077)

											tmp85079 := Call(__e, tmp85066, Nil, tmp85078)

											var ifres85065 Obj

											if True == tmp85079 {
												ifres85065 = True

											} else {
												ifres85065 = False

											}

											ifres85064 = ifres85065

										} else {
											ifres85064 = False

										}

										var ifres85063 Obj

										if True == ifres85064 {
											ifres85063 = True

										} else {
											ifres85063 = False

										}

										ifres85062 = ifres85063

									} else {
										ifres85062 = False

									}

									var ifres85061 Obj

									if True == ifres85062 {
										ifres85061 = True

									} else {
										ifres85061 = False

									}

									ifres85060 = ifres85061

								} else {
									ifres85060 = False

								}

								var ifres85059 Obj

								if True == ifres85060 {
									ifres85059 = True

								} else {
									ifres85059 = False

								}

								ifres85058 = ifres85059

							} else {
								ifres85058 = False

							}

							var ifres85057 Obj

							if True == ifres85058 {
								ifres85057 = True

							} else {
								ifres85057 = False

							}

							ifres85056 = ifres85057

						} else {
							ifres85056 = False

						}

						var ifres85055 Obj

						if True == ifres85056 {
							ifres85055 = True

						} else {
							ifres85055 = False

						}

						ifres85054 = ifres85055

					} else {
						ifres85054 = False

					}

					var ifres85053 Obj

					if True == ifres85054 {
						ifres85053 = True

					} else {
						ifres85053 = False

					}

					ifres85052 = ifres85053

				} else {
					ifres85052 = False

				}

				var ifres85051 Obj

				if True == ifres85052 {
					ifres85051 = True

				} else {
					ifres85051 = False

				}

				ifres85050 = ifres85051

			} else {
				ifres85050 = False

			}

			var ifres85049 Obj

			if True == ifres85050 {
				ifres85049 = True

			} else {
				ifres85049 = False

			}

			ifres85048 = ifres85049

		} else {
			ifres85048 = False

		}

		if True == ifres85048 {
			tmp85013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85015 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85016 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85017 := Call(__e, tmp85016, V1118)

			tmp85018 := Call(__e, tmp85015, tmp85017)

			tmp85019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

			tmp85021 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85022 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85023 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85024 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85025 := Call(__e, tmp85024, V1118)

			tmp85026 := Call(__e, tmp85023, tmp85025)

			tmp85027 := Call(__e, tmp85022, tmp85026)

			tmp85028 := Call(__e, tmp85021, tmp85027)

			tmp85029 := Call(__e, tmp85020, tmp85028)

			tmp85030 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85031 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

			tmp85032 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85033 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85035 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85036 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85037 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85038 := Call(__e, tmp85037, V1118)

			tmp85039 := Call(__e, tmp85036, tmp85038)

			tmp85040 := Call(__e, tmp85035, tmp85039)

			tmp85041 := Call(__e, tmp85034, tmp85040)

			tmp85042 := Call(__e, tmp85033, tmp85041)

			tmp85043 := Call(__e, tmp85032, tmp85042)

			tmp85044 := Call(__e, tmp85031, tmp85043)

			tmp85045 := Call(__e, tmp85030, tmp85044, Nil)

			tmp85046 := Call(__e, tmp85019, tmp85029, tmp85045)

			tmp85047 := Call(__e, tmp85014, tmp85018, tmp85046)

			__e.TailApply(tmp85013, symlet, tmp85047)
			return

		} else {
			tmp85011 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85012 := Call(__e, tmp85011, V1118)

			var ifres84925 Obj

			if True == tmp85012 {
				tmp85007 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp85008 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85009 := Call(__e, tmp85008, V1118)

				tmp85010 := Call(__e, tmp85007, symshen_4the, tmp85009)

				var ifres84927 Obj

				if True == tmp85010 {
					tmp85003 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp85004 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85005 := Call(__e, tmp85004, V1118)

					tmp85006 := Call(__e, tmp85003, tmp85005)

					var ifres84929 Obj

					if True == tmp85006 {
						tmp84997 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp84998 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp84999 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp85000 := Call(__e, tmp84999, V1118)

						tmp85001 := Call(__e, tmp84998, tmp85000)

						tmp85002 := Call(__e, tmp84997, symshen_4result, tmp85001)

						var ifres84931 Obj

						if True == tmp85002 {
							tmp84991 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp84992 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp84993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp84994 := Call(__e, tmp84993, V1118)

							tmp84995 := Call(__e, tmp84992, tmp84994)

							tmp84996 := Call(__e, tmp84991, tmp84995)

							var ifres84933 Obj

							if True == tmp84996 {
								tmp84983 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp84984 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp84985 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84986 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84987 := Call(__e, tmp84986, V1118)

								tmp84988 := Call(__e, tmp84985, tmp84987)

								tmp84989 := Call(__e, tmp84984, tmp84988)

								tmp84990 := Call(__e, tmp84983, symshen_4of, tmp84989)

								var ifres84935 Obj

								if True == tmp84990 {
									tmp84975 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp84976 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84977 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84978 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84979 := Call(__e, tmp84978, V1118)

									tmp84980 := Call(__e, tmp84977, tmp84979)

									tmp84981 := Call(__e, tmp84976, tmp84980)

									tmp84982 := Call(__e, tmp84975, tmp84981)

									var ifres84937 Obj

									if True == tmp84982 {
										tmp84965 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp84966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp84967 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84968 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84969 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84970 := Call(__e, tmp84969, V1118)

										tmp84971 := Call(__e, tmp84968, tmp84970)

										tmp84972 := Call(__e, tmp84967, tmp84971)

										tmp84973 := Call(__e, tmp84966, tmp84972)

										tmp84974 := Call(__e, tmp84965, symshen_4dereferencing, tmp84973)

										var ifres84939 Obj

										if True == tmp84974 {
											tmp84955 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp84956 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84960 := Call(__e, tmp84959, V1118)

											tmp84961 := Call(__e, tmp84958, tmp84960)

											tmp84962 := Call(__e, tmp84957, tmp84961)

											tmp84963 := Call(__e, tmp84956, tmp84962)

											tmp84964 := Call(__e, tmp84955, tmp84963)

											var ifres84941 Obj

											if True == tmp84964 {
												tmp84943 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp84944 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84945 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84946 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84947 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84948 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84949 := Call(__e, tmp84948, V1118)

												tmp84950 := Call(__e, tmp84947, tmp84949)

												tmp84951 := Call(__e, tmp84946, tmp84950)

												tmp84952 := Call(__e, tmp84945, tmp84951)

												tmp84953 := Call(__e, tmp84944, tmp84952)

												tmp84954 := Call(__e, tmp84943, Nil, tmp84953)

												var ifres84942 Obj

												if True == tmp84954 {
													ifres84942 = True

												} else {
													ifres84942 = False

												}

												ifres84941 = ifres84942

											} else {
												ifres84941 = False

											}

											var ifres84940 Obj

											if True == ifres84941 {
												ifres84940 = True

											} else {
												ifres84940 = False

											}

											ifres84939 = ifres84940

										} else {
											ifres84939 = False

										}

										var ifres84938 Obj

										if True == ifres84939 {
											ifres84938 = True

										} else {
											ifres84938 = False

										}

										ifres84937 = ifres84938

									} else {
										ifres84937 = False

									}

									var ifres84936 Obj

									if True == ifres84937 {
										ifres84936 = True

									} else {
										ifres84936 = False

									}

									ifres84935 = ifres84936

								} else {
									ifres84935 = False

								}

								var ifres84934 Obj

								if True == ifres84935 {
									ifres84934 = True

								} else {
									ifres84934 = False

								}

								ifres84933 = ifres84934

							} else {
								ifres84933 = False

							}

							var ifres84932 Obj

							if True == ifres84933 {
								ifres84932 = True

							} else {
								ifres84932 = False

							}

							ifres84931 = ifres84932

						} else {
							ifres84931 = False

						}

						var ifres84930 Obj

						if True == ifres84931 {
							ifres84930 = True

						} else {
							ifres84930 = False

						}

						ifres84929 = ifres84930

					} else {
						ifres84929 = False

					}

					var ifres84928 Obj

					if True == ifres84929 {
						ifres84928 = True

					} else {
						ifres84928 = False

					}

					ifres84927 = ifres84928

				} else {
					ifres84927 = False

				}

				var ifres84926 Obj

				if True == ifres84927 {
					ifres84926 = True

				} else {
					ifres84926 = False

				}

				ifres84925 = ifres84926

			} else {
				ifres84925 = False

			}

			if True == ifres84925 {
				tmp84908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp84909 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp84910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

				tmp84911 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp84912 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp84913 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp84914 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp84915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp84916 := Call(__e, tmp84915, V1118)

				tmp84917 := Call(__e, tmp84914, tmp84916)

				tmp84918 := Call(__e, tmp84913, tmp84917)

				tmp84919 := Call(__e, tmp84912, tmp84918)

				tmp84920 := Call(__e, tmp84911, tmp84919)

				tmp84921 := Call(__e, tmp84910, tmp84920)

				tmp84922 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp84923 := Call(__e, tmp84922, symProcessN, Nil)

				tmp84924 := Call(__e, tmp84909, tmp84921, tmp84923)

				__e.TailApply(tmp84908, symshen_4lazyderef, tmp84924)
				return

			} else {
				tmp84906 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp84907 := Call(__e, tmp84906, V1118)

				var ifres84810 Obj

				if True == tmp84907 {
					tmp84902 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp84903 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp84904 := Call(__e, tmp84903, V1118)

					tmp84905 := Call(__e, tmp84902, symif, tmp84904)

					var ifres84812 Obj

					if True == tmp84905 {
						tmp84898 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp84899 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp84900 := Call(__e, tmp84899, V1118)

						tmp84901 := Call(__e, tmp84898, tmp84900)

						var ifres84814 Obj

						if True == tmp84901 {
							tmp84892 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp84893 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp84894 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp84895 := Call(__e, tmp84894, V1118)

							tmp84896 := Call(__e, tmp84893, tmp84895)

							tmp84897 := Call(__e, tmp84892, tmp84896)

							var ifres84816 Obj

							if True == tmp84897 {
								tmp84884 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp84885 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp84886 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84887 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84888 := Call(__e, tmp84887, V1118)

								tmp84889 := Call(__e, tmp84886, tmp84888)

								tmp84890 := Call(__e, tmp84885, tmp84889)

								tmp84891 := Call(__e, tmp84884, symshen_4then, tmp84890)

								var ifres84818 Obj

								if True == tmp84891 {
									tmp84876 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp84877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84878 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84879 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84880 := Call(__e, tmp84879, V1118)

									tmp84881 := Call(__e, tmp84878, tmp84880)

									tmp84882 := Call(__e, tmp84877, tmp84881)

									tmp84883 := Call(__e, tmp84876, tmp84882)

									var ifres84820 Obj

									if True == tmp84883 {
										tmp84866 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp84867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84868 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84869 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84870 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84871 := Call(__e, tmp84870, V1118)

										tmp84872 := Call(__e, tmp84869, tmp84871)

										tmp84873 := Call(__e, tmp84868, tmp84872)

										tmp84874 := Call(__e, tmp84867, tmp84873)

										tmp84875 := Call(__e, tmp84866, tmp84874)

										var ifres84822 Obj

										if True == tmp84875 {
											tmp84854 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp84855 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp84856 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84857 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84858 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84859 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84860 := Call(__e, tmp84859, V1118)

											tmp84861 := Call(__e, tmp84858, tmp84860)

											tmp84862 := Call(__e, tmp84857, tmp84861)

											tmp84863 := Call(__e, tmp84856, tmp84862)

											tmp84864 := Call(__e, tmp84855, tmp84863)

											tmp84865 := Call(__e, tmp84854, symshen_4else, tmp84864)

											var ifres84824 Obj

											if True == tmp84865 {
												tmp84842 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp84843 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84846 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84847 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84848 := Call(__e, tmp84847, V1118)

												tmp84849 := Call(__e, tmp84846, tmp84848)

												tmp84850 := Call(__e, tmp84845, tmp84849)

												tmp84851 := Call(__e, tmp84844, tmp84850)

												tmp84852 := Call(__e, tmp84843, tmp84851)

												tmp84853 := Call(__e, tmp84842, tmp84852)

												var ifres84826 Obj

												if True == tmp84853 {
													tmp84828 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp84829 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84830 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84831 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84832 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84835 := Call(__e, tmp84834, V1118)

													tmp84836 := Call(__e, tmp84833, tmp84835)

													tmp84837 := Call(__e, tmp84832, tmp84836)

													tmp84838 := Call(__e, tmp84831, tmp84837)

													tmp84839 := Call(__e, tmp84830, tmp84838)

													tmp84840 := Call(__e, tmp84829, tmp84839)

													tmp84841 := Call(__e, tmp84828, Nil, tmp84840)

													var ifres84827 Obj

													if True == tmp84841 {
														ifres84827 = True

													} else {
														ifres84827 = False

													}

													ifres84826 = ifres84827

												} else {
													ifres84826 = False

												}

												var ifres84825 Obj

												if True == ifres84826 {
													ifres84825 = True

												} else {
													ifres84825 = False

												}

												ifres84824 = ifres84825

											} else {
												ifres84824 = False

											}

											var ifres84823 Obj

											if True == ifres84824 {
												ifres84823 = True

											} else {
												ifres84823 = False

											}

											ifres84822 = ifres84823

										} else {
											ifres84822 = False

										}

										var ifres84821 Obj

										if True == ifres84822 {
											ifres84821 = True

										} else {
											ifres84821 = False

										}

										ifres84820 = ifres84821

									} else {
										ifres84820 = False

									}

									var ifres84819 Obj

									if True == ifres84820 {
										ifres84819 = True

									} else {
										ifres84819 = False

									}

									ifres84818 = ifres84819

								} else {
									ifres84818 = False

								}

								var ifres84817 Obj

								if True == ifres84818 {
									ifres84817 = True

								} else {
									ifres84817 = False

								}

								ifres84816 = ifres84817

							} else {
								ifres84816 = False

							}

							var ifres84815 Obj

							if True == ifres84816 {
								ifres84815 = True

							} else {
								ifres84815 = False

							}

							ifres84814 = ifres84815

						} else {
							ifres84814 = False

						}

						var ifres84813 Obj

						if True == ifres84814 {
							ifres84813 = True

						} else {
							ifres84813 = False

						}

						ifres84812 = ifres84813

					} else {
						ifres84812 = False

					}

					var ifres84811 Obj

					if True == ifres84812 {
						ifres84811 = True

					} else {
						ifres84811 = False

					}

					ifres84810 = ifres84811

				} else {
					ifres84810 = False

				}

				if True == ifres84810 {
					tmp84773 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp84774 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp84775 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

					tmp84776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp84777 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84778 := Call(__e, tmp84777, V1118)

					tmp84779 := Call(__e, tmp84776, tmp84778)

					tmp84780 := Call(__e, tmp84775, tmp84779)

					tmp84781 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp84782 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

					tmp84783 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp84784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84786 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84787 := Call(__e, tmp84786, V1118)

					tmp84788 := Call(__e, tmp84785, tmp84787)

					tmp84789 := Call(__e, tmp84784, tmp84788)

					tmp84790 := Call(__e, tmp84783, tmp84789)

					tmp84791 := Call(__e, tmp84782, tmp84790)

					tmp84792 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp84793 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

					tmp84794 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp84795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84799 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp84800 := Call(__e, tmp84799, V1118)

					tmp84801 := Call(__e, tmp84798, tmp84800)

					tmp84802 := Call(__e, tmp84797, tmp84801)

					tmp84803 := Call(__e, tmp84796, tmp84802)

					tmp84804 := Call(__e, tmp84795, tmp84803)

					tmp84805 := Call(__e, tmp84794, tmp84804)

					tmp84806 := Call(__e, tmp84793, tmp84805)

					tmp84807 := Call(__e, tmp84792, tmp84806, Nil)

					tmp84808 := Call(__e, tmp84781, tmp84791, tmp84807)

					tmp84809 := Call(__e, tmp84774, tmp84780, tmp84808)

					__e.TailApply(tmp84773, symif, tmp84809)
					return

				} else {
					tmp84771 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp84772 := Call(__e, tmp84771, V1118)

					var ifres84705 Obj

					if True == tmp84772 {
						tmp84767 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp84768 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp84769 := Call(__e, tmp84768, V1118)

						tmp84770 := Call(__e, tmp84767, tmp84769)

						var ifres84707 Obj

						if True == tmp84770 {
							tmp84761 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp84762 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp84763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp84764 := Call(__e, tmp84763, V1118)

							tmp84765 := Call(__e, tmp84762, tmp84764)

							tmp84766 := Call(__e, tmp84761, symis, tmp84765)

							var ifres84709 Obj

							if True == tmp84766 {
								tmp84755 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp84756 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84757 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84758 := Call(__e, tmp84757, V1118)

								tmp84759 := Call(__e, tmp84756, tmp84758)

								tmp84760 := Call(__e, tmp84755, tmp84759)

								var ifres84711 Obj

								if True == tmp84760 {
									tmp84747 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp84748 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp84749 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84750 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84751 := Call(__e, tmp84750, V1118)

									tmp84752 := Call(__e, tmp84749, tmp84751)

									tmp84753 := Call(__e, tmp84748, tmp84752)

									tmp84754 := Call(__e, tmp84747, symshen_4a, tmp84753)

									var ifres84713 Obj

									if True == tmp84754 {
										tmp84739 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp84740 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84741 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84742 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84743 := Call(__e, tmp84742, V1118)

										tmp84744 := Call(__e, tmp84741, tmp84743)

										tmp84745 := Call(__e, tmp84740, tmp84744)

										tmp84746 := Call(__e, tmp84739, tmp84745)

										var ifres84715 Obj

										if True == tmp84746 {
											tmp84729 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp84730 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp84731 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84732 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84733 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84734 := Call(__e, tmp84733, V1118)

											tmp84735 := Call(__e, tmp84732, tmp84734)

											tmp84736 := Call(__e, tmp84731, tmp84735)

											tmp84737 := Call(__e, tmp84730, tmp84736)

											tmp84738 := Call(__e, tmp84729, symshen_4variable, tmp84737)

											var ifres84717 Obj

											if True == tmp84738 {
												tmp84719 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp84720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84722 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84723 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84724 := Call(__e, tmp84723, V1118)

												tmp84725 := Call(__e, tmp84722, tmp84724)

												tmp84726 := Call(__e, tmp84721, tmp84725)

												tmp84727 := Call(__e, tmp84720, tmp84726)

												tmp84728 := Call(__e, tmp84719, Nil, tmp84727)

												var ifres84718 Obj

												if True == tmp84728 {
													ifres84718 = True

												} else {
													ifres84718 = False

												}

												ifres84717 = ifres84718

											} else {
												ifres84717 = False

											}

											var ifres84716 Obj

											if True == ifres84717 {
												ifres84716 = True

											} else {
												ifres84716 = False

											}

											ifres84715 = ifres84716

										} else {
											ifres84715 = False

										}

										var ifres84714 Obj

										if True == ifres84715 {
											ifres84714 = True

										} else {
											ifres84714 = False

										}

										ifres84713 = ifres84714

									} else {
										ifres84713 = False

									}

									var ifres84712 Obj

									if True == ifres84713 {
										ifres84712 = True

									} else {
										ifres84712 = False

									}

									ifres84711 = ifres84712

								} else {
									ifres84711 = False

								}

								var ifres84710 Obj

								if True == ifres84711 {
									ifres84710 = True

								} else {
									ifres84710 = False

								}

								ifres84709 = ifres84710

							} else {
								ifres84709 = False

							}

							var ifres84708 Obj

							if True == ifres84709 {
								ifres84708 = True

							} else {
								ifres84708 = False

							}

							ifres84707 = ifres84708

						} else {
							ifres84707 = False

						}

						var ifres84706 Obj

						if True == ifres84707 {
							ifres84706 = True

						} else {
							ifres84706 = False

						}

						ifres84705 = ifres84706

					} else {
						ifres84705 = False

					}

					if True == ifres84705 {
						tmp84700 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp84701 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp84702 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp84703 := Call(__e, tmp84702, V1118)

						tmp84704 := Call(__e, tmp84701, tmp84703, Nil)

						__e.TailApply(tmp84700, symshen_4pvar_2, tmp84704)
						return

					} else {
						tmp84698 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp84699 := Call(__e, tmp84698, V1118)

						var ifres84604 Obj

						if True == tmp84699 {
							tmp84694 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp84695 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp84696 := Call(__e, tmp84695, V1118)

							tmp84697 := Call(__e, tmp84694, tmp84696)

							var ifres84606 Obj

							if True == tmp84697 {
								tmp84688 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp84689 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp84690 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84691 := Call(__e, tmp84690, V1118)

								tmp84692 := Call(__e, tmp84689, tmp84691)

								tmp84693 := Call(__e, tmp84688, symis, tmp84692)

								var ifres84608 Obj

								if True == tmp84693 {
									tmp84682 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp84683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84684 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84685 := Call(__e, tmp84684, V1118)

									tmp84686 := Call(__e, tmp84683, tmp84685)

									tmp84687 := Call(__e, tmp84682, tmp84686)

									var ifres84610 Obj

									if True == tmp84687 {
										tmp84674 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp84675 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp84676 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84677 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84678 := Call(__e, tmp84677, V1118)

										tmp84679 := Call(__e, tmp84676, tmp84678)

										tmp84680 := Call(__e, tmp84675, tmp84679)

										tmp84681 := Call(__e, tmp84674, symshen_4a, tmp84680)

										var ifres84612 Obj

										if True == tmp84681 {
											tmp84666 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp84667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84668 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84669 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84670 := Call(__e, tmp84669, V1118)

											tmp84671 := Call(__e, tmp84668, tmp84670)

											tmp84672 := Call(__e, tmp84667, tmp84671)

											tmp84673 := Call(__e, tmp84666, tmp84672)

											var ifres84614 Obj

											if True == tmp84673 {
												tmp84656 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp84657 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp84658 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84661 := Call(__e, tmp84660, V1118)

												tmp84662 := Call(__e, tmp84659, tmp84661)

												tmp84663 := Call(__e, tmp84658, tmp84662)

												tmp84664 := Call(__e, tmp84657, tmp84663)

												tmp84665 := Call(__e, tmp84656, symshen_4non_1empty, tmp84664)

												var ifres84616 Obj

												if True == tmp84665 {
													tmp84646 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp84647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84648 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84649 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84651 := Call(__e, tmp84650, V1118)

													tmp84652 := Call(__e, tmp84649, tmp84651)

													tmp84653 := Call(__e, tmp84648, tmp84652)

													tmp84654 := Call(__e, tmp84647, tmp84653)

													tmp84655 := Call(__e, tmp84646, tmp84654)

													var ifres84618 Obj

													if True == tmp84655 {
														tmp84634 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp84635 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp84636 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84637 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84638 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84639 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84640 := Call(__e, tmp84639, V1118)

														tmp84641 := Call(__e, tmp84638, tmp84640)

														tmp84642 := Call(__e, tmp84637, tmp84641)

														tmp84643 := Call(__e, tmp84636, tmp84642)

														tmp84644 := Call(__e, tmp84635, tmp84643)

														tmp84645 := Call(__e, tmp84634, symlist, tmp84644)

														var ifres84620 Obj

														if True == tmp84645 {
															tmp84622 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp84623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84624 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84625 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84626 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84627 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84628 := Call(__e, tmp84627, V1118)

															tmp84629 := Call(__e, tmp84626, tmp84628)

															tmp84630 := Call(__e, tmp84625, tmp84629)

															tmp84631 := Call(__e, tmp84624, tmp84630)

															tmp84632 := Call(__e, tmp84623, tmp84631)

															tmp84633 := Call(__e, tmp84622, Nil, tmp84632)

															var ifres84621 Obj

															if True == tmp84633 {
																ifres84621 = True

															} else {
																ifres84621 = False

															}

															ifres84620 = ifres84621

														} else {
															ifres84620 = False

														}

														var ifres84619 Obj

														if True == ifres84620 {
															ifres84619 = True

														} else {
															ifres84619 = False

														}

														ifres84618 = ifres84619

													} else {
														ifres84618 = False

													}

													var ifres84617 Obj

													if True == ifres84618 {
														ifres84617 = True

													} else {
														ifres84617 = False

													}

													ifres84616 = ifres84617

												} else {
													ifres84616 = False

												}

												var ifres84615 Obj

												if True == ifres84616 {
													ifres84615 = True

												} else {
													ifres84615 = False

												}

												ifres84614 = ifres84615

											} else {
												ifres84614 = False

											}

											var ifres84613 Obj

											if True == ifres84614 {
												ifres84613 = True

											} else {
												ifres84613 = False

											}

											ifres84612 = ifres84613

										} else {
											ifres84612 = False

										}

										var ifres84611 Obj

										if True == ifres84612 {
											ifres84611 = True

										} else {
											ifres84611 = False

										}

										ifres84610 = ifres84611

									} else {
										ifres84610 = False

									}

									var ifres84609 Obj

									if True == ifres84610 {
										ifres84609 = True

									} else {
										ifres84609 = False

									}

									ifres84608 = ifres84609

								} else {
									ifres84608 = False

								}

								var ifres84607 Obj

								if True == ifres84608 {
									ifres84607 = True

								} else {
									ifres84607 = False

								}

								ifres84606 = ifres84607

							} else {
								ifres84606 = False

							}

							var ifres84605 Obj

							if True == ifres84606 {
								ifres84605 = True

							} else {
								ifres84605 = False

							}

							ifres84604 = ifres84605

						} else {
							ifres84604 = False

						}

						if True == ifres84604 {
							tmp84599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp84600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp84601 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp84602 := Call(__e, tmp84601, V1118)

							tmp84603 := Call(__e, tmp84600, tmp84602, Nil)

							__e.TailApply(tmp84599, symcons_2, tmp84603)
							return

						} else {
							tmp84597 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp84598 := Call(__e, tmp84597, V1118)

							var ifres84409 Obj

							if True == tmp84598 {
								tmp84593 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp84594 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp84595 := Call(__e, tmp84594, V1118)

								tmp84596 := Call(__e, tmp84593, symshen_4rename, tmp84595)

								var ifres84411 Obj

								if True == tmp84596 {
									tmp84589 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp84590 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84591 := Call(__e, tmp84590, V1118)

									tmp84592 := Call(__e, tmp84589, tmp84591)

									var ifres84413 Obj

									if True == tmp84592 {
										tmp84583 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp84584 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp84585 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84586 := Call(__e, tmp84585, V1118)

										tmp84587 := Call(__e, tmp84584, tmp84586)

										tmp84588 := Call(__e, tmp84583, symshen_4the, tmp84587)

										var ifres84415 Obj

										if True == tmp84588 {
											tmp84577 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp84578 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84579 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84580 := Call(__e, tmp84579, V1118)

											tmp84581 := Call(__e, tmp84578, tmp84580)

											tmp84582 := Call(__e, tmp84577, tmp84581)

											var ifres84417 Obj

											if True == tmp84582 {
												tmp84569 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp84570 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp84571 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84572 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84573 := Call(__e, tmp84572, V1118)

												tmp84574 := Call(__e, tmp84571, tmp84573)

												tmp84575 := Call(__e, tmp84570, tmp84574)

												tmp84576 := Call(__e, tmp84569, symshen_4variables, tmp84575)

												var ifres84419 Obj

												if True == tmp84576 {
													tmp84561 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp84562 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84563 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84564 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84565 := Call(__e, tmp84564, V1118)

													tmp84566 := Call(__e, tmp84563, tmp84565)

													tmp84567 := Call(__e, tmp84562, tmp84566)

													tmp84568 := Call(__e, tmp84561, tmp84567)

													var ifres84421 Obj

													if True == tmp84568 {
														tmp84551 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp84552 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp84553 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84554 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84555 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84556 := Call(__e, tmp84555, V1118)

														tmp84557 := Call(__e, tmp84554, tmp84556)

														tmp84558 := Call(__e, tmp84553, tmp84557)

														tmp84559 := Call(__e, tmp84552, tmp84558)

														tmp84560 := Call(__e, tmp84551, symin, tmp84559)

														var ifres84423 Obj

														if True == tmp84560 {
															tmp84541 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp84542 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84543 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84544 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84545 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84546 := Call(__e, tmp84545, V1118)

															tmp84547 := Call(__e, tmp84544, tmp84546)

															tmp84548 := Call(__e, tmp84543, tmp84547)

															tmp84549 := Call(__e, tmp84542, tmp84548)

															tmp84550 := Call(__e, tmp84541, tmp84549)

															var ifres84425 Obj

															if True == tmp84550 {
																tmp84529 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp84530 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp84531 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84532 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84533 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84534 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84535 := Call(__e, tmp84534, V1118)

																tmp84536 := Call(__e, tmp84533, tmp84535)

																tmp84537 := Call(__e, tmp84532, tmp84536)

																tmp84538 := Call(__e, tmp84531, tmp84537)

																tmp84539 := Call(__e, tmp84530, tmp84538)

																tmp84540 := Call(__e, tmp84529, Nil, tmp84539)

																var ifres84427 Obj

																if True == tmp84540 {
																	tmp84517 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp84518 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84519 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84520 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84521 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84522 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84523 := Call(__e, tmp84522, V1118)

																	tmp84524 := Call(__e, tmp84521, tmp84523)

																	tmp84525 := Call(__e, tmp84520, tmp84524)

																	tmp84526 := Call(__e, tmp84519, tmp84525)

																	tmp84527 := Call(__e, tmp84518, tmp84526)

																	tmp84528 := Call(__e, tmp84517, tmp84527)

																	var ifres84429 Obj

																	if True == tmp84528 {
																		tmp84503 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp84504 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																		tmp84505 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84506 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84507 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84508 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84509 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84510 := Call(__e, tmp84509, V1118)

																		tmp84511 := Call(__e, tmp84508, tmp84510)

																		tmp84512 := Call(__e, tmp84507, tmp84511)

																		tmp84513 := Call(__e, tmp84506, tmp84512)

																		tmp84514 := Call(__e, tmp84505, tmp84513)

																		tmp84515 := Call(__e, tmp84504, tmp84514)

																		tmp84516 := Call(__e, tmp84503, symand, tmp84515)

																		var ifres84431 Obj

																		if True == tmp84516 {
																			tmp84489 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																			tmp84490 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84491 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84492 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84493 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84494 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84495 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84496 := Call(__e, tmp84495, V1118)

																			tmp84497 := Call(__e, tmp84494, tmp84496)

																			tmp84498 := Call(__e, tmp84493, tmp84497)

																			tmp84499 := Call(__e, tmp84492, tmp84498)

																			tmp84500 := Call(__e, tmp84491, tmp84499)

																			tmp84501 := Call(__e, tmp84490, tmp84500)

																			tmp84502 := Call(__e, tmp84489, tmp84501)

																			var ifres84433 Obj

																			if True == tmp84502 {
																				tmp84473 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp84474 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																				tmp84475 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84478 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84479 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84480 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84481 := Call(__e, tmp84480, V1118)

																				tmp84482 := Call(__e, tmp84479, tmp84481)

																				tmp84483 := Call(__e, tmp84478, tmp84482)

																				tmp84484 := Call(__e, tmp84477, tmp84483)

																				tmp84485 := Call(__e, tmp84476, tmp84484)

																				tmp84486 := Call(__e, tmp84475, tmp84485)

																				tmp84487 := Call(__e, tmp84474, tmp84486)

																				tmp84488 := Call(__e, tmp84473, symshen_4then, tmp84487)

																				var ifres84435 Obj

																				if True == tmp84488 {
																					tmp84457 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																					tmp84458 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84459 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84460 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84461 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84462 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84464 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84465 := Call(__e, tmp84464, V1118)

																					tmp84466 := Call(__e, tmp84463, tmp84465)

																					tmp84467 := Call(__e, tmp84462, tmp84466)

																					tmp84468 := Call(__e, tmp84461, tmp84467)

																					tmp84469 := Call(__e, tmp84460, tmp84468)

																					tmp84470 := Call(__e, tmp84459, tmp84469)

																					tmp84471 := Call(__e, tmp84458, tmp84470)

																					tmp84472 := Call(__e, tmp84457, tmp84471)

																					var ifres84437 Obj

																					if True == tmp84472 {
																						tmp84439 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																						tmp84440 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84441 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84442 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84443 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84444 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84445 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84446 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84448 := Call(__e, tmp84447, V1118)

																						tmp84449 := Call(__e, tmp84446, tmp84448)

																						tmp84450 := Call(__e, tmp84445, tmp84449)

																						tmp84451 := Call(__e, tmp84444, tmp84450)

																						tmp84452 := Call(__e, tmp84443, tmp84451)

																						tmp84453 := Call(__e, tmp84442, tmp84452)

																						tmp84454 := Call(__e, tmp84441, tmp84453)

																						tmp84455 := Call(__e, tmp84440, tmp84454)

																						tmp84456 := Call(__e, tmp84439, Nil, tmp84455)

																						var ifres84438 Obj

																						if True == tmp84456 {
																							ifres84438 = True

																						} else {
																							ifres84438 = False

																						}

																						ifres84437 = ifres84438

																					} else {
																						ifres84437 = False

																					}

																					var ifres84436 Obj

																					if True == ifres84437 {
																						ifres84436 = True

																					} else {
																						ifres84436 = False

																					}

																					ifres84435 = ifres84436

																				} else {
																					ifres84435 = False

																				}

																				var ifres84434 Obj

																				if True == ifres84435 {
																					ifres84434 = True

																				} else {
																					ifres84434 = False

																				}

																				ifres84433 = ifres84434

																			} else {
																				ifres84433 = False

																			}

																			var ifres84432 Obj

																			if True == ifres84433 {
																				ifres84432 = True

																			} else {
																				ifres84432 = False

																			}

																			ifres84431 = ifres84432

																		} else {
																			ifres84431 = False

																		}

																		var ifres84430 Obj

																		if True == ifres84431 {
																			ifres84430 = True

																		} else {
																			ifres84430 = False

																		}

																		ifres84429 = ifres84430

																	} else {
																		ifres84429 = False

																	}

																	var ifres84428 Obj

																	if True == ifres84429 {
																		ifres84428 = True

																	} else {
																		ifres84428 = False

																	}

																	ifres84427 = ifres84428

																} else {
																	ifres84427 = False

																}

																var ifres84426 Obj

																if True == ifres84427 {
																	ifres84426 = True

																} else {
																	ifres84426 = False

																}

																ifres84425 = ifres84426

															} else {
																ifres84425 = False

															}

															var ifres84424 Obj

															if True == ifres84425 {
																ifres84424 = True

															} else {
																ifres84424 = False

															}

															ifres84423 = ifres84424

														} else {
															ifres84423 = False

														}

														var ifres84422 Obj

														if True == ifres84423 {
															ifres84422 = True

														} else {
															ifres84422 = False

														}

														ifres84421 = ifres84422

													} else {
														ifres84421 = False

													}

													var ifres84420 Obj

													if True == ifres84421 {
														ifres84420 = True

													} else {
														ifres84420 = False

													}

													ifres84419 = ifres84420

												} else {
													ifres84419 = False

												}

												var ifres84418 Obj

												if True == ifres84419 {
													ifres84418 = True

												} else {
													ifres84418 = False

												}

												ifres84417 = ifres84418

											} else {
												ifres84417 = False

											}

											var ifres84416 Obj

											if True == ifres84417 {
												ifres84416 = True

											} else {
												ifres84416 = False

											}

											ifres84415 = ifres84416

										} else {
											ifres84415 = False

										}

										var ifres84414 Obj

										if True == ifres84415 {
											ifres84414 = True

										} else {
											ifres84414 = False

										}

										ifres84413 = ifres84414

									} else {
										ifres84413 = False

									}

									var ifres84412 Obj

									if True == ifres84413 {
										ifres84412 = True

									} else {
										ifres84412 = False

									}

									ifres84411 = ifres84412

								} else {
									ifres84411 = False

								}

								var ifres84410 Obj

								if True == ifres84411 {
									ifres84410 = True

								} else {
									ifres84410 = False

								}

								ifres84409 = ifres84410

							} else {
								ifres84409 = False

							}

							if True == ifres84409 {
								tmp84392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

								tmp84393 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp84394 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84395 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84396 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84397 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84398 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84400 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp84401 := Call(__e, tmp84400, V1118)

								tmp84402 := Call(__e, tmp84399, tmp84401)

								tmp84403 := Call(__e, tmp84398, tmp84402)

								tmp84404 := Call(__e, tmp84397, tmp84403)

								tmp84405 := Call(__e, tmp84396, tmp84404)

								tmp84406 := Call(__e, tmp84395, tmp84405)

								tmp84407 := Call(__e, tmp84394, tmp84406)

								tmp84408 := Call(__e, tmp84393, tmp84407)

								__e.TailApply(tmp84392, tmp84408)
								return

							} else {
								tmp84390 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp84391 := Call(__e, tmp84390, V1118)

								var ifres84202 Obj

								if True == tmp84391 {
									tmp84386 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp84387 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp84388 := Call(__e, tmp84387, V1118)

									tmp84389 := Call(__e, tmp84386, symshen_4rename, tmp84388)

									var ifres84204 Obj

									if True == tmp84389 {
										tmp84382 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp84383 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84384 := Call(__e, tmp84383, V1118)

										tmp84385 := Call(__e, tmp84382, tmp84384)

										var ifres84206 Obj

										if True == tmp84385 {
											tmp84376 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp84377 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp84378 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84379 := Call(__e, tmp84378, V1118)

											tmp84380 := Call(__e, tmp84377, tmp84379)

											tmp84381 := Call(__e, tmp84376, symshen_4the, tmp84380)

											var ifres84208 Obj

											if True == tmp84381 {
												tmp84370 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp84371 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84372 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84373 := Call(__e, tmp84372, V1118)

												tmp84374 := Call(__e, tmp84371, tmp84373)

												tmp84375 := Call(__e, tmp84370, tmp84374)

												var ifres84210 Obj

												if True == tmp84375 {
													tmp84362 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp84363 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp84364 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84365 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84366 := Call(__e, tmp84365, V1118)

													tmp84367 := Call(__e, tmp84364, tmp84366)

													tmp84368 := Call(__e, tmp84363, tmp84367)

													tmp84369 := Call(__e, tmp84362, symshen_4variables, tmp84368)

													var ifres84212 Obj

													if True == tmp84369 {
														tmp84354 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp84355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84356 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84357 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84358 := Call(__e, tmp84357, V1118)

														tmp84359 := Call(__e, tmp84356, tmp84358)

														tmp84360 := Call(__e, tmp84355, tmp84359)

														tmp84361 := Call(__e, tmp84354, tmp84360)

														var ifres84214 Obj

														if True == tmp84361 {
															tmp84344 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp84345 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp84346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84347 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84348 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84349 := Call(__e, tmp84348, V1118)

															tmp84350 := Call(__e, tmp84347, tmp84349)

															tmp84351 := Call(__e, tmp84346, tmp84350)

															tmp84352 := Call(__e, tmp84345, tmp84351)

															tmp84353 := Call(__e, tmp84344, symin, tmp84352)

															var ifres84216 Obj

															if True == tmp84353 {
																tmp84334 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp84335 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84336 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84337 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84339 := Call(__e, tmp84338, V1118)

																tmp84340 := Call(__e, tmp84337, tmp84339)

																tmp84341 := Call(__e, tmp84336, tmp84340)

																tmp84342 := Call(__e, tmp84335, tmp84341)

																tmp84343 := Call(__e, tmp84334, tmp84342)

																var ifres84218 Obj

																if True == tmp84343 {
																	tmp84322 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp84323 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																	tmp84324 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84325 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84326 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84327 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84328 := Call(__e, tmp84327, V1118)

																	tmp84329 := Call(__e, tmp84326, tmp84328)

																	tmp84330 := Call(__e, tmp84325, tmp84329)

																	tmp84331 := Call(__e, tmp84324, tmp84330)

																	tmp84332 := Call(__e, tmp84323, tmp84331)

																	tmp84333 := Call(__e, tmp84322, tmp84332)

																	var ifres84220 Obj

																	if True == tmp84333 {
																		tmp84310 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp84311 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84312 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84313 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84314 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84315 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84316 := Call(__e, tmp84315, V1118)

																		tmp84317 := Call(__e, tmp84314, tmp84316)

																		tmp84318 := Call(__e, tmp84313, tmp84317)

																		tmp84319 := Call(__e, tmp84312, tmp84318)

																		tmp84320 := Call(__e, tmp84311, tmp84319)

																		tmp84321 := Call(__e, tmp84310, tmp84320)

																		var ifres84222 Obj

																		if True == tmp84321 {
																			tmp84296 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp84297 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp84298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84299 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84300 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84301 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84302 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp84303 := Call(__e, tmp84302, V1118)

																			tmp84304 := Call(__e, tmp84301, tmp84303)

																			tmp84305 := Call(__e, tmp84300, tmp84304)

																			tmp84306 := Call(__e, tmp84299, tmp84305)

																			tmp84307 := Call(__e, tmp84298, tmp84306)

																			tmp84308 := Call(__e, tmp84297, tmp84307)

																			tmp84309 := Call(__e, tmp84296, symand, tmp84308)

																			var ifres84224 Obj

																			if True == tmp84309 {
																				tmp84282 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																				tmp84283 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84284 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84285 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84286 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84287 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp84289 := Call(__e, tmp84288, V1118)

																				tmp84290 := Call(__e, tmp84287, tmp84289)

																				tmp84291 := Call(__e, tmp84286, tmp84290)

																				tmp84292 := Call(__e, tmp84285, tmp84291)

																				tmp84293 := Call(__e, tmp84284, tmp84292)

																				tmp84294 := Call(__e, tmp84283, tmp84293)

																				tmp84295 := Call(__e, tmp84282, tmp84294)

																				var ifres84226 Obj

																				if True == tmp84295 {
																					tmp84266 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp84267 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																					tmp84268 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84269 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84270 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84271 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84272 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84273 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp84274 := Call(__e, tmp84273, V1118)

																					tmp84275 := Call(__e, tmp84272, tmp84274)

																					tmp84276 := Call(__e, tmp84271, tmp84275)

																					tmp84277 := Call(__e, tmp84270, tmp84276)

																					tmp84278 := Call(__e, tmp84269, tmp84277)

																					tmp84279 := Call(__e, tmp84268, tmp84278)

																					tmp84280 := Call(__e, tmp84267, tmp84279)

																					tmp84281 := Call(__e, tmp84266, symshen_4then, tmp84280)

																					var ifres84228 Obj

																					if True == tmp84281 {
																						tmp84250 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																						tmp84251 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84252 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84253 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84254 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84255 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84256 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84257 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp84258 := Call(__e, tmp84257, V1118)

																						tmp84259 := Call(__e, tmp84256, tmp84258)

																						tmp84260 := Call(__e, tmp84255, tmp84259)

																						tmp84261 := Call(__e, tmp84254, tmp84260)

																						tmp84262 := Call(__e, tmp84253, tmp84261)

																						tmp84263 := Call(__e, tmp84252, tmp84262)

																						tmp84264 := Call(__e, tmp84251, tmp84263)

																						tmp84265 := Call(__e, tmp84250, tmp84264)

																						var ifres84230 Obj

																						if True == tmp84265 {
																							tmp84232 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							tmp84233 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84234 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84235 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84236 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84238 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84239 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp84241 := Call(__e, tmp84240, V1118)

																							tmp84242 := Call(__e, tmp84239, tmp84241)

																							tmp84243 := Call(__e, tmp84238, tmp84242)

																							tmp84244 := Call(__e, tmp84237, tmp84243)

																							tmp84245 := Call(__e, tmp84236, tmp84244)

																							tmp84246 := Call(__e, tmp84235, tmp84245)

																							tmp84247 := Call(__e, tmp84234, tmp84246)

																							tmp84248 := Call(__e, tmp84233, tmp84247)

																							tmp84249 := Call(__e, tmp84232, Nil, tmp84248)

																							var ifres84231 Obj

																							if True == tmp84249 {
																								ifres84231 = True

																							} else {
																								ifres84231 = False

																							}

																							ifres84230 = ifres84231

																						} else {
																							ifres84230 = False

																						}

																						var ifres84229 Obj

																						if True == ifres84230 {
																							ifres84229 = True

																						} else {
																							ifres84229 = False

																						}

																						ifres84228 = ifres84229

																					} else {
																						ifres84228 = False

																					}

																					var ifres84227 Obj

																					if True == ifres84228 {
																						ifres84227 = True

																					} else {
																						ifres84227 = False

																					}

																					ifres84226 = ifres84227

																				} else {
																					ifres84226 = False

																				}

																				var ifres84225 Obj

																				if True == ifres84226 {
																					ifres84225 = True

																				} else {
																					ifres84225 = False

																				}

																				ifres84224 = ifres84225

																			} else {
																				ifres84224 = False

																			}

																			var ifres84223 Obj

																			if True == ifres84224 {
																				ifres84223 = True

																			} else {
																				ifres84223 = False

																			}

																			ifres84222 = ifres84223

																		} else {
																			ifres84222 = False

																		}

																		var ifres84221 Obj

																		if True == ifres84222 {
																			ifres84221 = True

																		} else {
																			ifres84221 = False

																		}

																		ifres84220 = ifres84221

																	} else {
																		ifres84220 = False

																	}

																	var ifres84219 Obj

																	if True == ifres84220 {
																		ifres84219 = True

																	} else {
																		ifres84219 = False

																	}

																	ifres84218 = ifres84219

																} else {
																	ifres84218 = False

																}

																var ifres84217 Obj

																if True == ifres84218 {
																	ifres84217 = True

																} else {
																	ifres84217 = False

																}

																ifres84216 = ifres84217

															} else {
																ifres84216 = False

															}

															var ifres84215 Obj

															if True == ifres84216 {
																ifres84215 = True

															} else {
																ifres84215 = False

															}

															ifres84214 = ifres84215

														} else {
															ifres84214 = False

														}

														var ifres84213 Obj

														if True == ifres84214 {
															ifres84213 = True

														} else {
															ifres84213 = False

														}

														ifres84212 = ifres84213

													} else {
														ifres84212 = False

													}

													var ifres84211 Obj

													if True == ifres84212 {
														ifres84211 = True

													} else {
														ifres84211 = False

													}

													ifres84210 = ifres84211

												} else {
													ifres84210 = False

												}

												var ifres84209 Obj

												if True == ifres84210 {
													ifres84209 = True

												} else {
													ifres84209 = False

												}

												ifres84208 = ifres84209

											} else {
												ifres84208 = False

											}

											var ifres84207 Obj

											if True == ifres84208 {
												ifres84207 = True

											} else {
												ifres84207 = False

											}

											ifres84206 = ifres84207

										} else {
											ifres84206 = False

										}

										var ifres84205 Obj

										if True == ifres84206 {
											ifres84205 = True

										} else {
											ifres84205 = False

										}

										ifres84204 = ifres84205

									} else {
										ifres84204 = False

									}

									var ifres84203 Obj

									if True == ifres84204 {
										ifres84203 = True

									} else {
										ifres84203 = False

									}

									ifres84202 = ifres84203

								} else {
									ifres84202 = False

								}

								if True == ifres84202 {
									tmp84145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84146 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84147 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp84148 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp84149 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84150 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84151 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84152 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84153 := Call(__e, tmp84152, V1118)

									tmp84154 := Call(__e, tmp84151, tmp84153)

									tmp84155 := Call(__e, tmp84150, tmp84154)

									tmp84156 := Call(__e, tmp84149, tmp84155)

									tmp84157 := Call(__e, tmp84148, tmp84156)

									tmp84158 := Call(__e, tmp84147, tmp84157)

									tmp84159 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84162 := Call(__e, tmp84161, symProcessN, Nil)

									tmp84163 := Call(__e, tmp84160, symshen_4newpv, tmp84162)

									tmp84164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84165 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

									tmp84166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84167 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84168 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84169 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84170 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp84171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84172 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp84173 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84174 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84175 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84176 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84177 := Call(__e, tmp84176, V1118)

									tmp84178 := Call(__e, tmp84175, tmp84177)

									tmp84179 := Call(__e, tmp84174, tmp84178)

									tmp84180 := Call(__e, tmp84173, tmp84179)

									tmp84181 := Call(__e, tmp84172, tmp84180)

									tmp84182 := Call(__e, tmp84171, tmp84181)

									tmp84183 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84184 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84185 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp84188 := Call(__e, tmp84187, V1118)

									tmp84189 := Call(__e, tmp84186, tmp84188)

									tmp84190 := Call(__e, tmp84185, tmp84189)

									tmp84191 := Call(__e, tmp84184, tmp84190)

									tmp84192 := Call(__e, tmp84183, tmp84191)

									tmp84193 := Call(__e, tmp84170, tmp84182, tmp84192)

									tmp84194 := Call(__e, tmp84169, symin, tmp84193)

									tmp84195 := Call(__e, tmp84168, symshen_4variables, tmp84194)

									tmp84196 := Call(__e, tmp84167, symshen_4the, tmp84195)

									tmp84197 := Call(__e, tmp84166, symshen_4rename, tmp84196)

									tmp84198 := Call(__e, tmp84165, tmp84197)

									tmp84199 := Call(__e, tmp84164, tmp84198, Nil)

									tmp84200 := Call(__e, tmp84159, tmp84163, tmp84199)

									tmp84201 := Call(__e, tmp84146, tmp84158, tmp84200)

									__e.TailApply(tmp84145, symlet, tmp84201)
									return

								} else {
									tmp84143 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp84144 := Call(__e, tmp84143, V1118)

									var ifres84047 Obj

									if True == tmp84144 {
										tmp84139 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp84140 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp84141 := Call(__e, tmp84140, V1118)

										tmp84142 := Call(__e, tmp84139, symbind, tmp84141)

										var ifres84049 Obj

										if True == tmp84142 {
											tmp84135 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp84136 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp84137 := Call(__e, tmp84136, V1118)

											tmp84138 := Call(__e, tmp84135, tmp84137)

											var ifres84051 Obj

											if True == tmp84138 {
												tmp84129 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp84130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp84132 := Call(__e, tmp84131, V1118)

												tmp84133 := Call(__e, tmp84130, tmp84132)

												tmp84134 := Call(__e, tmp84129, tmp84133)

												var ifres84053 Obj

												if True == tmp84134 {
													tmp84121 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp84122 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp84123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84124 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp84125 := Call(__e, tmp84124, V1118)

													tmp84126 := Call(__e, tmp84123, tmp84125)

													tmp84127 := Call(__e, tmp84122, tmp84126)

													tmp84128 := Call(__e, tmp84121, symshen_4to, tmp84127)

													var ifres84055 Obj

													if True == tmp84128 {
														tmp84113 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp84114 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84115 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84116 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp84117 := Call(__e, tmp84116, V1118)

														tmp84118 := Call(__e, tmp84115, tmp84117)

														tmp84119 := Call(__e, tmp84114, tmp84118)

														tmp84120 := Call(__e, tmp84113, tmp84119)

														var ifres84057 Obj

														if True == tmp84120 {
															tmp84103 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp84104 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84105 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84106 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84107 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp84108 := Call(__e, tmp84107, V1118)

															tmp84109 := Call(__e, tmp84106, tmp84108)

															tmp84110 := Call(__e, tmp84105, tmp84109)

															tmp84111 := Call(__e, tmp84104, tmp84110)

															tmp84112 := Call(__e, tmp84103, tmp84111)

															var ifres84059 Obj

															if True == tmp84112 {
																tmp84091 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp84092 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp84093 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84094 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84095 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84096 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp84097 := Call(__e, tmp84096, V1118)

																tmp84098 := Call(__e, tmp84095, tmp84097)

																tmp84099 := Call(__e, tmp84094, tmp84098)

																tmp84100 := Call(__e, tmp84093, tmp84099)

																tmp84101 := Call(__e, tmp84092, tmp84100)

																tmp84102 := Call(__e, tmp84091, symin, tmp84101)

																var ifres84061 Obj

																if True == tmp84102 {
																	tmp84079 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp84080 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84081 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84082 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84083 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84084 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp84085 := Call(__e, tmp84084, V1118)

																	tmp84086 := Call(__e, tmp84083, tmp84085)

																	tmp84087 := Call(__e, tmp84082, tmp84086)

																	tmp84088 := Call(__e, tmp84081, tmp84087)

																	tmp84089 := Call(__e, tmp84080, tmp84088)

																	tmp84090 := Call(__e, tmp84079, tmp84089)

																	var ifres84063 Obj

																	if True == tmp84090 {
																		tmp84065 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp84066 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84067 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84069 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp84072 := Call(__e, tmp84071, V1118)

																		tmp84073 := Call(__e, tmp84070, tmp84072)

																		tmp84074 := Call(__e, tmp84069, tmp84073)

																		tmp84075 := Call(__e, tmp84068, tmp84074)

																		tmp84076 := Call(__e, tmp84067, tmp84075)

																		tmp84077 := Call(__e, tmp84066, tmp84076)

																		tmp84078 := Call(__e, tmp84065, Nil, tmp84077)

																		var ifres84064 Obj

																		if True == tmp84078 {
																			ifres84064 = True

																		} else {
																			ifres84064 = False

																		}

																		ifres84063 = ifres84064

																	} else {
																		ifres84063 = False

																	}

																	var ifres84062 Obj

																	if True == ifres84063 {
																		ifres84062 = True

																	} else {
																		ifres84062 = False

																	}

																	ifres84061 = ifres84062

																} else {
																	ifres84061 = False

																}

																var ifres84060 Obj

																if True == ifres84061 {
																	ifres84060 = True

																} else {
																	ifres84060 = False

																}

																ifres84059 = ifres84060

															} else {
																ifres84059 = False

															}

															var ifres84058 Obj

															if True == ifres84059 {
																ifres84058 = True

															} else {
																ifres84058 = False

															}

															ifres84057 = ifres84058

														} else {
															ifres84057 = False

														}

														var ifres84056 Obj

														if True == ifres84057 {
															ifres84056 = True

														} else {
															ifres84056 = False

														}

														ifres84055 = ifres84056

													} else {
														ifres84055 = False

													}

													var ifres84054 Obj

													if True == ifres84055 {
														ifres84054 = True

													} else {
														ifres84054 = False

													}

													ifres84053 = ifres84054

												} else {
													ifres84053 = False

												}

												var ifres84052 Obj

												if True == ifres84053 {
													ifres84052 = True

												} else {
													ifres84052 = False

												}

												ifres84051 = ifres84052

											} else {
												ifres84051 = False

											}

											var ifres84050 Obj

											if True == ifres84051 {
												ifres84050 = True

											} else {
												ifres84050 = False

											}

											ifres84049 = ifres84050

										} else {
											ifres84049 = False

										}

										var ifres84048 Obj

										if True == ifres84049 {
											ifres84048 = True

										} else {
											ifres84048 = False

										}

										ifres84047 = ifres84048

									} else {
										ifres84047 = False

									}

									if True == ifres84047 {
										tmp83982 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp83983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp83984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp83985 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp83986 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp83987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83988 := Call(__e, tmp83987, V1118)

										tmp83989 := Call(__e, tmp83986, tmp83988)

										tmp83990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp83991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4chwild)

										tmp83992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp83993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83994 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83995 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp83996 := Call(__e, tmp83995, V1118)

										tmp83997 := Call(__e, tmp83994, tmp83996)

										tmp83998 := Call(__e, tmp83993, tmp83997)

										tmp83999 := Call(__e, tmp83992, tmp83998)

										tmp84000 := Call(__e, tmp83991, tmp83999)

										tmp84001 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84002 := Call(__e, tmp84001, symProcessN, Nil)

										tmp84003 := Call(__e, tmp83990, tmp84000, tmp84002)

										tmp84004 := Call(__e, tmp83985, tmp83989, tmp84003)

										tmp84005 := Call(__e, tmp83984, symshen_4bindv, tmp84004)

										tmp84006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84007 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84010 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

										tmp84011 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp84012 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84013 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84014 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84015 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84016 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84017 := Call(__e, tmp84016, V1118)

										tmp84018 := Call(__e, tmp84015, tmp84017)

										tmp84019 := Call(__e, tmp84014, tmp84018)

										tmp84020 := Call(__e, tmp84013, tmp84019)

										tmp84021 := Call(__e, tmp84012, tmp84020)

										tmp84022 := Call(__e, tmp84011, tmp84021)

										tmp84023 := Call(__e, tmp84010, tmp84022)

										tmp84024 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84025 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84027 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84028 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84029 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp84030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp84031 := Call(__e, tmp84030, V1118)

										tmp84032 := Call(__e, tmp84029, tmp84031)

										tmp84033 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84034 := Call(__e, tmp84033, symProcessN, Nil)

										tmp84035 := Call(__e, tmp84028, tmp84032, tmp84034)

										tmp84036 := Call(__e, tmp84027, symshen_4unbindv, tmp84035)

										tmp84037 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp84038 := Call(__e, tmp84037, symResult, Nil)

										tmp84039 := Call(__e, tmp84026, tmp84036, tmp84038)

										tmp84040 := Call(__e, tmp84025, symdo, tmp84039)

										tmp84041 := Call(__e, tmp84024, tmp84040, Nil)

										tmp84042 := Call(__e, tmp84009, tmp84023, tmp84041)

										tmp84043 := Call(__e, tmp84008, symResult, tmp84042)

										tmp84044 := Call(__e, tmp84007, symlet, tmp84043)

										tmp84045 := Call(__e, tmp84006, tmp84044, Nil)

										tmp84046 := Call(__e, tmp83983, tmp84005, tmp84045)

										__e.TailApply(tmp83982, symdo, tmp84046)
										return

									} else {
										tmp83980 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp83981 := Call(__e, tmp83980, V1118)

										var ifres83900 Obj

										if True == tmp83981 {
											tmp83976 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp83977 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83978 := Call(__e, tmp83977, V1118)

											tmp83979 := Call(__e, tmp83976, tmp83978)

											var ifres83902 Obj

											if True == tmp83979 {
												tmp83970 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp83971 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp83972 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp83973 := Call(__e, tmp83972, V1118)

												tmp83974 := Call(__e, tmp83971, tmp83973)

												tmp83975 := Call(__e, tmp83970, symis, tmp83974)

												var ifres83904 Obj

												if True == tmp83975 {
													tmp83964 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp83965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83966 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83967 := Call(__e, tmp83966, V1118)

													tmp83968 := Call(__e, tmp83965, tmp83967)

													tmp83969 := Call(__e, tmp83964, tmp83968)

													var ifres83906 Obj

													if True == tmp83969 {
														tmp83956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp83957 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp83958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83960 := Call(__e, tmp83959, V1118)

														tmp83961 := Call(__e, tmp83958, tmp83960)

														tmp83962 := Call(__e, tmp83957, tmp83961)

														tmp83963 := Call(__e, tmp83956, symidentical, tmp83962)

														var ifres83908 Obj

														if True == tmp83963 {
															tmp83948 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp83949 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp83950 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp83951 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp83952 := Call(__e, tmp83951, V1118)

															tmp83953 := Call(__e, tmp83950, tmp83952)

															tmp83954 := Call(__e, tmp83949, tmp83953)

															tmp83955 := Call(__e, tmp83948, tmp83954)

															var ifres83910 Obj

															if True == tmp83955 {
																tmp83938 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp83939 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp83940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83941 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83942 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83943 := Call(__e, tmp83942, V1118)

																tmp83944 := Call(__e, tmp83941, tmp83943)

																tmp83945 := Call(__e, tmp83940, tmp83944)

																tmp83946 := Call(__e, tmp83939, tmp83945)

																tmp83947 := Call(__e, tmp83938, symshen_4to, tmp83946)

																var ifres83912 Obj

																if True == tmp83947 {
																	tmp83928 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp83929 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83930 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83931 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83932 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83933 := Call(__e, tmp83932, V1118)

																	tmp83934 := Call(__e, tmp83931, tmp83933)

																	tmp83935 := Call(__e, tmp83930, tmp83934)

																	tmp83936 := Call(__e, tmp83929, tmp83935)

																	tmp83937 := Call(__e, tmp83928, tmp83936)

																	var ifres83914 Obj

																	if True == tmp83937 {
																		tmp83916 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp83917 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83918 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83919 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83920 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83921 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83922 := Call(__e, tmp83921, V1118)

																		tmp83923 := Call(__e, tmp83920, tmp83922)

																		tmp83924 := Call(__e, tmp83919, tmp83923)

																		tmp83925 := Call(__e, tmp83918, tmp83924)

																		tmp83926 := Call(__e, tmp83917, tmp83925)

																		tmp83927 := Call(__e, tmp83916, Nil, tmp83926)

																		var ifres83915 Obj

																		if True == tmp83927 {
																			ifres83915 = True

																		} else {
																			ifres83915 = False

																		}

																		ifres83914 = ifres83915

																	} else {
																		ifres83914 = False

																	}

																	var ifres83913 Obj

																	if True == ifres83914 {
																		ifres83913 = True

																	} else {
																		ifres83913 = False

																	}

																	ifres83912 = ifres83913

																} else {
																	ifres83912 = False

																}

																var ifres83911 Obj

																if True == ifres83912 {
																	ifres83911 = True

																} else {
																	ifres83911 = False

																}

																ifres83910 = ifres83911

															} else {
																ifres83910 = False

															}

															var ifres83909 Obj

															if True == ifres83910 {
																ifres83909 = True

															} else {
																ifres83909 = False

															}

															ifres83908 = ifres83909

														} else {
															ifres83908 = False

														}

														var ifres83907 Obj

														if True == ifres83908 {
															ifres83907 = True

														} else {
															ifres83907 = False

														}

														ifres83906 = ifres83907

													} else {
														ifres83906 = False

													}

													var ifres83905 Obj

													if True == ifres83906 {
														ifres83905 = True

													} else {
														ifres83905 = False

													}

													ifres83904 = ifres83905

												} else {
													ifres83904 = False

												}

												var ifres83903 Obj

												if True == ifres83904 {
													ifres83903 = True

												} else {
													ifres83903 = False

												}

												ifres83902 = ifres83903

											} else {
												ifres83902 = False

											}

											var ifres83901 Obj

											if True == ifres83902 {
												ifres83901 = True

											} else {
												ifres83901 = False

											}

											ifres83900 = ifres83901

										} else {
											ifres83900 = False

										}

										if True == ifres83900 {
											tmp83883 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp83884 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp83885 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp83886 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83887 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83888 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83889 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp83890 := Call(__e, tmp83889, V1118)

											tmp83891 := Call(__e, tmp83888, tmp83890)

											tmp83892 := Call(__e, tmp83887, tmp83891)

											tmp83893 := Call(__e, tmp83886, tmp83892)

											tmp83894 := Call(__e, tmp83885, tmp83893)

											tmp83895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp83896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp83897 := Call(__e, tmp83896, V1118)

											tmp83898 := Call(__e, tmp83895, tmp83897, Nil)

											tmp83899 := Call(__e, tmp83884, tmp83894, tmp83898)

											__e.TailApply(tmp83883, sym_a, tmp83899)
											return

										} else {
											tmp83881 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp83882 := Call(__e, tmp83881, symshen_4failed_b, V1118)

											if True == tmp83882 {
												__e.Return(False)
												return
											} else {
												tmp83879 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp83880 := Call(__e, tmp83879, V1118)

												var ifres83819 Obj

												if True == tmp83880 {
													tmp83875 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp83876 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp83877 := Call(__e, tmp83876, V1118)

													tmp83878 := Call(__e, tmp83875, symshen_4the, tmp83877)

													var ifres83821 Obj

													if True == tmp83878 {
														tmp83871 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp83872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83873 := Call(__e, tmp83872, V1118)

														tmp83874 := Call(__e, tmp83871, tmp83873)

														var ifres83823 Obj

														if True == tmp83874 {
															tmp83865 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp83866 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp83867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp83868 := Call(__e, tmp83867, V1118)

															tmp83869 := Call(__e, tmp83866, tmp83868)

															tmp83870 := Call(__e, tmp83865, symhead, tmp83869)

															var ifres83825 Obj

															if True == tmp83870 {
																tmp83859 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp83860 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83861 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83862 := Call(__e, tmp83861, V1118)

																tmp83863 := Call(__e, tmp83860, tmp83862)

																tmp83864 := Call(__e, tmp83859, tmp83863)

																var ifres83827 Obj

																if True == tmp83864 {
																	tmp83851 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp83852 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																	tmp83853 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83854 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83855 := Call(__e, tmp83854, V1118)

																	tmp83856 := Call(__e, tmp83853, tmp83855)

																	tmp83857 := Call(__e, tmp83852, tmp83856)

																	tmp83858 := Call(__e, tmp83851, symshen_4of, tmp83857)

																	var ifres83829 Obj

																	if True == tmp83858 {
																		tmp83843 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp83844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83846 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83847 := Call(__e, tmp83846, V1118)

																		tmp83848 := Call(__e, tmp83845, tmp83847)

																		tmp83849 := Call(__e, tmp83844, tmp83848)

																		tmp83850 := Call(__e, tmp83843, tmp83849)

																		var ifres83831 Obj

																		if True == tmp83850 {
																			tmp83833 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp83834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83836 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83838 := Call(__e, tmp83837, V1118)

																			tmp83839 := Call(__e, tmp83836, tmp83838)

																			tmp83840 := Call(__e, tmp83835, tmp83839)

																			tmp83841 := Call(__e, tmp83834, tmp83840)

																			tmp83842 := Call(__e, tmp83833, Nil, tmp83841)

																			var ifres83832 Obj

																			if True == tmp83842 {
																				ifres83832 = True

																			} else {
																				ifres83832 = False

																			}

																			ifres83831 = ifres83832

																		} else {
																			ifres83831 = False

																		}

																		var ifres83830 Obj

																		if True == ifres83831 {
																			ifres83830 = True

																		} else {
																			ifres83830 = False

																		}

																		ifres83829 = ifres83830

																	} else {
																		ifres83829 = False

																	}

																	var ifres83828 Obj

																	if True == ifres83829 {
																		ifres83828 = True

																	} else {
																		ifres83828 = False

																	}

																	ifres83827 = ifres83828

																} else {
																	ifres83827 = False

																}

																var ifres83826 Obj

																if True == ifres83827 {
																	ifres83826 = True

																} else {
																	ifres83826 = False

																}

																ifres83825 = ifres83826

															} else {
																ifres83825 = False

															}

															var ifres83824 Obj

															if True == ifres83825 {
																ifres83824 = True

															} else {
																ifres83824 = False

															}

															ifres83823 = ifres83824

														} else {
															ifres83823 = False

														}

														var ifres83822 Obj

														if True == ifres83823 {
															ifres83822 = True

														} else {
															ifres83822 = False

														}

														ifres83821 = ifres83822

													} else {
														ifres83821 = False

													}

													var ifres83820 Obj

													if True == ifres83821 {
														ifres83820 = True

													} else {
														ifres83820 = False

													}

													ifres83819 = ifres83820

												} else {
													ifres83819 = False

												}

												if True == ifres83819 {
													tmp83812 := Call(__e, PrimNS1Value(symns2_1value), symcons)

													tmp83813 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83814 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83815 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp83816 := Call(__e, tmp83815, V1118)

													tmp83817 := Call(__e, tmp83814, tmp83816)

													tmp83818 := Call(__e, tmp83813, tmp83817)

													__e.TailApply(tmp83812, symhd, tmp83818)
													return

												} else {
													tmp83810 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp83811 := Call(__e, tmp83810, V1118)

													var ifres83750 Obj

													if True == tmp83811 {
														tmp83806 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp83807 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp83808 := Call(__e, tmp83807, V1118)

														tmp83809 := Call(__e, tmp83806, symshen_4the, tmp83808)

														var ifres83752 Obj

														if True == tmp83809 {
															tmp83802 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp83803 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp83804 := Call(__e, tmp83803, V1118)

															tmp83805 := Call(__e, tmp83802, tmp83804)

															var ifres83754 Obj

															if True == tmp83805 {
																tmp83796 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp83797 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp83798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83799 := Call(__e, tmp83798, V1118)

																tmp83800 := Call(__e, tmp83797, tmp83799)

																tmp83801 := Call(__e, tmp83796, symtail, tmp83800)

																var ifres83756 Obj

																if True == tmp83801 {
																	tmp83790 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp83791 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83792 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83793 := Call(__e, tmp83792, V1118)

																	tmp83794 := Call(__e, tmp83791, tmp83793)

																	tmp83795 := Call(__e, tmp83790, tmp83794)

																	var ifres83758 Obj

																	if True == tmp83795 {
																		tmp83782 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp83783 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																		tmp83784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83786 := Call(__e, tmp83785, V1118)

																		tmp83787 := Call(__e, tmp83784, tmp83786)

																		tmp83788 := Call(__e, tmp83783, tmp83787)

																		tmp83789 := Call(__e, tmp83782, symshen_4of, tmp83788)

																		var ifres83760 Obj

																		if True == tmp83789 {
																			tmp83774 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																			tmp83775 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83776 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83777 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83778 := Call(__e, tmp83777, V1118)

																			tmp83779 := Call(__e, tmp83776, tmp83778)

																			tmp83780 := Call(__e, tmp83775, tmp83779)

																			tmp83781 := Call(__e, tmp83774, tmp83780)

																			var ifres83762 Obj

																			if True == tmp83781 {
																				tmp83764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp83765 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83766 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83767 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83768 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83769 := Call(__e, tmp83768, V1118)

																				tmp83770 := Call(__e, tmp83767, tmp83769)

																				tmp83771 := Call(__e, tmp83766, tmp83770)

																				tmp83772 := Call(__e, tmp83765, tmp83771)

																				tmp83773 := Call(__e, tmp83764, Nil, tmp83772)

																				var ifres83763 Obj

																				if True == tmp83773 {
																					ifres83763 = True

																				} else {
																					ifres83763 = False

																				}

																				ifres83762 = ifres83763

																			} else {
																				ifres83762 = False

																			}

																			var ifres83761 Obj

																			if True == ifres83762 {
																				ifres83761 = True

																			} else {
																				ifres83761 = False

																			}

																			ifres83760 = ifres83761

																		} else {
																			ifres83760 = False

																		}

																		var ifres83759 Obj

																		if True == ifres83760 {
																			ifres83759 = True

																		} else {
																			ifres83759 = False

																		}

																		ifres83758 = ifres83759

																	} else {
																		ifres83758 = False

																	}

																	var ifres83757 Obj

																	if True == ifres83758 {
																		ifres83757 = True

																	} else {
																		ifres83757 = False

																	}

																	ifres83756 = ifres83757

																} else {
																	ifres83756 = False

																}

																var ifres83755 Obj

																if True == ifres83756 {
																	ifres83755 = True

																} else {
																	ifres83755 = False

																}

																ifres83754 = ifres83755

															} else {
																ifres83754 = False

															}

															var ifres83753 Obj

															if True == ifres83754 {
																ifres83753 = True

															} else {
																ifres83753 = False

															}

															ifres83752 = ifres83753

														} else {
															ifres83752 = False

														}

														var ifres83751 Obj

														if True == ifres83752 {
															ifres83751 = True

														} else {
															ifres83751 = False

														}

														ifres83750 = ifres83751

													} else {
														ifres83750 = False

													}

													if True == ifres83750 {
														tmp83743 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp83744 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83745 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83746 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp83747 := Call(__e, tmp83746, V1118)

														tmp83748 := Call(__e, tmp83745, tmp83747)

														tmp83749 := Call(__e, tmp83744, tmp83748)

														__e.TailApply(tmp83743, symtl, tmp83749)
														return

													} else {
														tmp83741 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp83742 := Call(__e, tmp83741, V1118)

														var ifres83693 Obj

														if True == tmp83742 {
															tmp83737 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp83738 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp83739 := Call(__e, tmp83738, V1118)

															tmp83740 := Call(__e, tmp83737, symshen_4pop, tmp83739)

															var ifres83695 Obj

															if True == tmp83740 {
																tmp83733 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp83734 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83735 := Call(__e, tmp83734, V1118)

																tmp83736 := Call(__e, tmp83733, tmp83735)

																var ifres83697 Obj

																if True == tmp83736 {
																	tmp83727 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp83728 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																	tmp83729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83730 := Call(__e, tmp83729, V1118)

																	tmp83731 := Call(__e, tmp83728, tmp83730)

																	tmp83732 := Call(__e, tmp83727, symshen_4the, tmp83731)

																	var ifres83699 Obj

																	if True == tmp83732 {
																		tmp83721 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp83722 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83723 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83724 := Call(__e, tmp83723, V1118)

																		tmp83725 := Call(__e, tmp83722, tmp83724)

																		tmp83726 := Call(__e, tmp83721, tmp83725)

																		var ifres83701 Obj

																		if True == tmp83726 {
																			tmp83713 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp83714 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp83715 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83716 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83717 := Call(__e, tmp83716, V1118)

																			tmp83718 := Call(__e, tmp83715, tmp83717)

																			tmp83719 := Call(__e, tmp83714, tmp83718)

																			tmp83720 := Call(__e, tmp83713, symshen_4stack, tmp83719)

																			var ifres83703 Obj

																			if True == tmp83720 {
																				tmp83705 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp83706 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83707 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83708 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83709 := Call(__e, tmp83708, V1118)

																				tmp83710 := Call(__e, tmp83707, tmp83709)

																				tmp83711 := Call(__e, tmp83706, tmp83710)

																				tmp83712 := Call(__e, tmp83705, Nil, tmp83711)

																				var ifres83704 Obj

																				if True == tmp83712 {
																					ifres83704 = True

																				} else {
																					ifres83704 = False

																				}

																				ifres83703 = ifres83704

																			} else {
																				ifres83703 = False

																			}

																			var ifres83702 Obj

																			if True == ifres83703 {
																				ifres83702 = True

																			} else {
																				ifres83702 = False

																			}

																			ifres83701 = ifres83702

																		} else {
																			ifres83701 = False

																		}

																		var ifres83700 Obj

																		if True == ifres83701 {
																			ifres83700 = True

																		} else {
																			ifres83700 = False

																		}

																		ifres83699 = ifres83700

																	} else {
																		ifres83699 = False

																	}

																	var ifres83698 Obj

																	if True == ifres83699 {
																		ifres83698 = True

																	} else {
																		ifres83698 = False

																	}

																	ifres83697 = ifres83698

																} else {
																	ifres83697 = False

																}

																var ifres83696 Obj

																if True == ifres83697 {
																	ifres83696 = True

																} else {
																	ifres83696 = False

																}

																ifres83695 = ifres83696

															} else {
																ifres83695 = False

															}

															var ifres83694 Obj

															if True == ifres83695 {
																ifres83694 = True

															} else {
																ifres83694 = False

															}

															ifres83693 = ifres83694

														} else {
															ifres83693 = False

														}

														if True == ifres83693 {
															tmp83682 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp83683 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp83684 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp83685 := Call(__e, tmp83684, symshen_4incinfs, Nil)

															tmp83686 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp83687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp83688 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp83689 := Call(__e, tmp83688, symContinuation, Nil)

															tmp83690 := Call(__e, tmp83687, symthaw, tmp83689)

															tmp83691 := Call(__e, tmp83686, tmp83690, Nil)

															tmp83692 := Call(__e, tmp83683, tmp83685, tmp83691)

															__e.TailApply(tmp83682, symdo, tmp83692)
															return

														} else {
															tmp83680 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp83681 := Call(__e, tmp83680, V1118)

															var ifres83620 Obj

															if True == tmp83681 {
																tmp83676 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp83677 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp83678 := Call(__e, tmp83677, V1118)

																tmp83679 := Call(__e, tmp83676, symcall, tmp83678)

																var ifres83622 Obj

																if True == tmp83679 {
																	tmp83672 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp83673 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp83674 := Call(__e, tmp83673, V1118)

																	tmp83675 := Call(__e, tmp83672, tmp83674)

																	var ifres83624 Obj

																	if True == tmp83675 {
																		tmp83666 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp83667 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																		tmp83668 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp83669 := Call(__e, tmp83668, V1118)

																		tmp83670 := Call(__e, tmp83667, tmp83669)

																		tmp83671 := Call(__e, tmp83666, symshen_4the, tmp83670)

																		var ifres83626 Obj

																		if True == tmp83671 {
																			tmp83660 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																			tmp83661 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83662 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp83663 := Call(__e, tmp83662, V1118)

																			tmp83664 := Call(__e, tmp83661, tmp83663)

																			tmp83665 := Call(__e, tmp83660, tmp83664)

																			var ifres83628 Obj

																			if True == tmp83665 {
																				tmp83652 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp83653 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																				tmp83654 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83655 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp83656 := Call(__e, tmp83655, V1118)

																				tmp83657 := Call(__e, tmp83654, tmp83656)

																				tmp83658 := Call(__e, tmp83653, tmp83657)

																				tmp83659 := Call(__e, tmp83652, symshen_4continuation, tmp83658)

																				var ifres83630 Obj

																				if True == tmp83659 {
																					tmp83644 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																					tmp83645 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp83646 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp83647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp83648 := Call(__e, tmp83647, V1118)

																					tmp83649 := Call(__e, tmp83646, tmp83648)

																					tmp83650 := Call(__e, tmp83645, tmp83649)

																					tmp83651 := Call(__e, tmp83644, tmp83650)

																					var ifres83632 Obj

																					if True == tmp83651 {
																						tmp83634 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																						tmp83635 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp83636 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp83637 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp83638 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp83639 := Call(__e, tmp83638, V1118)

																						tmp83640 := Call(__e, tmp83637, tmp83639)

																						tmp83641 := Call(__e, tmp83636, tmp83640)

																						tmp83642 := Call(__e, tmp83635, tmp83641)

																						tmp83643 := Call(__e, tmp83634, Nil, tmp83642)

																						var ifres83633 Obj

																						if True == tmp83643 {
																							ifres83633 = True

																						} else {
																							ifres83633 = False

																						}

																						ifres83632 = ifres83633

																					} else {
																						ifres83632 = False

																					}

																					var ifres83631 Obj

																					if True == ifres83632 {
																						ifres83631 = True

																					} else {
																						ifres83631 = False

																					}

																					ifres83630 = ifres83631

																				} else {
																					ifres83630 = False

																				}

																				var ifres83629 Obj

																				if True == ifres83630 {
																					ifres83629 = True

																				} else {
																					ifres83629 = False

																				}

																				ifres83628 = ifres83629

																			} else {
																				ifres83628 = False

																			}

																			var ifres83627 Obj

																			if True == ifres83628 {
																				ifres83627 = True

																			} else {
																				ifres83627 = False

																			}

																			ifres83626 = ifres83627

																		} else {
																			ifres83626 = False

																		}

																		var ifres83625 Obj

																		if True == ifres83626 {
																			ifres83625 = True

																		} else {
																			ifres83625 = False

																		}

																		ifres83624 = ifres83625

																	} else {
																		ifres83624 = False

																	}

																	var ifres83623 Obj

																	if True == ifres83624 {
																		ifres83623 = True

																	} else {
																		ifres83623 = False

																	}

																	ifres83622 = ifres83623

																} else {
																	ifres83622 = False

																}

																var ifres83621 Obj

																if True == ifres83622 {
																	ifres83621 = True

																} else {
																	ifres83621 = False

																}

																ifres83620 = ifres83621

															} else {
																ifres83620 = False

															}

															if True == ifres83620 {
																tmp83601 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp83602 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp83603 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp83604 := Call(__e, tmp83603, symshen_4incinfs, Nil)

																tmp83605 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp83606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4call__the__continuation)

																tmp83607 := Call(__e, PrimNS1Value(symns2_1value), symshen_4chwild)

																tmp83608 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp83609 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83610 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83611 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp83612 := Call(__e, tmp83611, V1118)

																tmp83613 := Call(__e, tmp83610, tmp83612)

																tmp83614 := Call(__e, tmp83609, tmp83613)

																tmp83615 := Call(__e, tmp83608, tmp83614)

																tmp83616 := Call(__e, tmp83607, tmp83615)

																tmp83617 := Call(__e, tmp83606, tmp83616, symProcessN, symContinuation)

																tmp83618 := Call(__e, tmp83605, tmp83617, Nil)

																tmp83619 := Call(__e, tmp83602, tmp83604, tmp83618)

																__e.TailApply(tmp83601, symdo, tmp83619)
																return

															} else {
																__e.Return(V1118)
																return
															}

														}

													}

												}

											}

										}

									}

								}

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp85146 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aum__to__shen, tmp83586)

	_ = tmp85146

	tmp85147 := MakeNative(func(__e *ControlFlow) {
		V1120 := __e.Get(1)
		_ = V1120
		tmp85158 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85159 := Call(__e, tmp85158, V1120, sym__)

		if True == tmp85159 {
			tmp85155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85157 := Call(__e, tmp85156, symProcessN, Nil)

			__e.TailApply(tmp85155, symshen_4newpv, tmp85157)
			return

		} else {
			tmp85153 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85154 := Call(__e, tmp85153, V1120)

			if True == tmp85154 {
				tmp85150 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp85151 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					tmp85152 := Call(__e, PrimNS1Value(symns2_1value), symshen_4chwild)

					__e.TailApply(tmp85152, Z)
					return

				}, 1)

				__e.TailApply(tmp85150, tmp85151, V1120)
				return

			} else {
				__e.Return(V1120)
				return
			}

		}

	}, 1)

	tmp85160 := Call(__e, PrimNS1Value(symns2_1set), symshen_4chwild, tmp85147)

	_ = tmp85160

	tmp85161 := MakeNative(func(__e *ControlFlow) {
		V1122 := __e.Get(1)
		_ = V1122
		tmp85162 := MakeNative(func(__e *ControlFlow) {
			Count_71 := __e.Get(1)
			_ = Count_71
			tmp85163 := MakeNative(func(__e *ControlFlow) {
				IncVar := __e.Get(1)
				_ = IncVar
				tmp85164 := MakeNative(func(__e *ControlFlow) {
					Vector := __e.Get(1)
					_ = Vector
					tmp85165 := MakeNative(func(__e *ControlFlow) {
						ResizeVectorIfNeeded := __e.Get(1)
						_ = ResizeVectorIfNeeded
						tmp85166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mk_1pvar)

						__e.TailApply(tmp85166, Count_71)
						return

					}, 1)

					tmp85170 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp85171 := Call(__e, PrimNS1Value(symns2_1value), symlimit)

					tmp85172 := Call(__e, tmp85171, Vector)

					tmp85173 := Call(__e, tmp85170, Count_71, tmp85172)

					var ifres85167 Obj

					if True == tmp85173 {
						tmp85168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4resizeprocessvector)

						tmp85169 := Call(__e, tmp85168, V1122, Count_71)

						ifres85167 = tmp85169

					} else {
						ifres85167 = symshen_4skip

					}

					__e.TailApply(tmp85165, ifres85167)
					return

				}, 1)

				tmp85174 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				tmp85175 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp85176 := Call(__e, tmp85175, symshen_4_dprologvectors_d)

				tmp85177 := Call(__e, tmp85174, tmp85176, V1122)

				__e.TailApply(tmp85164, tmp85177)
				return

			}, 1)

			tmp85178 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp85179 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp85180 := Call(__e, tmp85179, symshen_4_dvarcounter_d)

			tmp85181 := Call(__e, tmp85178, tmp85180, V1122, Count_71)

			__e.TailApply(tmp85163, tmp85181)
			return

		}, 1)

		tmp85182 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85183 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85184 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85185 := Call(__e, tmp85184, symshen_4_dvarcounter_d)

		tmp85186 := Call(__e, tmp85183, tmp85185, V1122)

		tmp85187 := Call(__e, tmp85182, tmp85186, MakeNumber(1))

		__e.TailApply(tmp85162, tmp85187)
		return

	}, 1)

	tmp85188 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newpv, tmp85161)

	_ = tmp85188

	tmp85189 := MakeNative(func(__e *ControlFlow) {
		V1125 := __e.Get(1)
		_ = V1125
		V1126 := __e.Get(2)
		_ = V1126
		tmp85190 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp85191 := MakeNative(func(__e *ControlFlow) {
				BigVector := __e.Get(1)
				_ = BigVector
				tmp85192 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

				tmp85193 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp85194 := Call(__e, tmp85193, symshen_4_dprologvectors_d)

				__e.TailApply(tmp85192, tmp85194, V1125, BigVector)
				return

			}, 1)

			tmp85195 := Call(__e, PrimNS1Value(symns2_1value), symshen_4resize_1vector)

			tmp85196 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp85197 := Call(__e, tmp85196, V1126, V1126)

			tmp85198 := Call(__e, tmp85195, Vector, tmp85197, symshen_4_1null_1)

			__e.TailApply(tmp85191, tmp85198)
			return

		}, 1)

		tmp85199 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85200 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85201 := Call(__e, tmp85200, symshen_4_dprologvectors_d)

		tmp85202 := Call(__e, tmp85199, tmp85201, V1125)

		__e.TailApply(tmp85190, tmp85202)
		return

	}, 2)

	tmp85203 := Call(__e, PrimNS1Value(symns2_1set), symshen_4resizeprocessvector, tmp85189)

	_ = tmp85203

	tmp85204 := MakeNative(func(__e *ControlFlow) {
		V1130 := __e.Get(1)
		_ = V1130
		V1131 := __e.Get(2)
		_ = V1131
		V1132 := __e.Get(3)
		_ = V1132
		tmp85205 := MakeNative(func(__e *ControlFlow) {
			BigVector := __e.Get(1)
			_ = BigVector
			tmp85206 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copy_1vector)

			tmp85207 := Call(__e, PrimNS1Value(symns2_1value), symlimit)

			tmp85208 := Call(__e, tmp85207, V1130)

			__e.TailApply(tmp85206, V1130, BigVector, tmp85208, V1131, V1132)
			return

		}, 1)

		tmp85209 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		tmp85210 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

		tmp85211 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85212 := Call(__e, tmp85211, MakeNumber(1), V1131)

		tmp85213 := Call(__e, tmp85210, tmp85212)

		tmp85214 := Call(__e, tmp85209, tmp85213, MakeNumber(0), V1131)

		__e.TailApply(tmp85205, tmp85214)
		return

	}, 3)

	tmp85215 := Call(__e, PrimNS1Value(symns2_1set), symshen_4resize_1vector, tmp85204)

	_ = tmp85215

	tmp85216 := MakeNative(func(__e *ControlFlow) {
		V1138 := __e.Get(1)
		_ = V1138
		V1139 := __e.Get(2)
		_ = V1139
		V1140 := __e.Get(3)
		_ = V1140
		V1141 := __e.Get(4)
		_ = V1141
		V1142 := __e.Get(5)
		_ = V1142
		tmp85217 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copy_1vector_1stage_12)

		tmp85218 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85219 := Call(__e, tmp85218, MakeNumber(1), V1140)

		tmp85220 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85221 := Call(__e, tmp85220, V1141, MakeNumber(1))

		tmp85222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copy_1vector_1stage_11)

		tmp85223 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85224 := Call(__e, tmp85223, MakeNumber(1), V1140)

		tmp85225 := Call(__e, tmp85222, MakeNumber(1), V1138, V1139, tmp85224)

		__e.TailApply(tmp85217, tmp85219, tmp85221, V1142, tmp85225)
		return

	}, 5)

	tmp85226 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copy_1vector, tmp85216)

	_ = tmp85226

	tmp85227 := MakeNative(func(__e *ControlFlow) {
		V1150 := __e.Get(1)
		_ = V1150
		V1151 := __e.Get(2)
		_ = V1151
		V1152 := __e.Get(3)
		_ = V1152
		V1153 := __e.Get(4)
		_ = V1153
		tmp85236 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85237 := Call(__e, tmp85236, V1153, V1150)

		if True == tmp85237 {
			__e.Return(V1152)
			return
		} else {
			tmp85229 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copy_1vector_1stage_11)

			tmp85230 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp85231 := Call(__e, tmp85230, MakeNumber(1), V1150)

			tmp85232 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp85233 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp85234 := Call(__e, tmp85233, V1151, V1150)

			tmp85235 := Call(__e, tmp85232, V1152, V1150, tmp85234)

			__e.TailApply(tmp85229, tmp85231, V1151, tmp85235, V1153)
			return

		}

	}, 4)

	tmp85238 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copy_1vector_1stage_11, tmp85227)

	_ = tmp85238

	tmp85239 := MakeNative(func(__e *ControlFlow) {
		V1161 := __e.Get(1)
		_ = V1161
		V1162 := __e.Get(2)
		_ = V1162
		V1163 := __e.Get(3)
		_ = V1163
		V1164 := __e.Get(4)
		_ = V1164
		tmp85246 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85247 := Call(__e, tmp85246, V1162, V1161)

		if True == tmp85247 {
			__e.Return(V1164)
			return
		} else {
			tmp85241 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copy_1vector_1stage_12)

			tmp85242 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp85243 := Call(__e, tmp85242, V1161, MakeNumber(1))

			tmp85244 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp85245 := Call(__e, tmp85244, V1164, V1161, V1163)

			__e.TailApply(tmp85241, tmp85243, V1162, V1163, tmp85245)
			return

		}

	}, 4)

	tmp85248 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copy_1vector_1stage_12, tmp85239)

	_ = tmp85248

	tmp85249 := MakeNative(func(__e *ControlFlow) {
		V1166 := __e.Get(1)
		_ = V1166
		tmp85250 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		tmp85251 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		tmp85252 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

		tmp85253 := Call(__e, tmp85252, MakeNumber(2))

		tmp85254 := Call(__e, tmp85251, tmp85253, MakeNumber(0), symshen_4pvar)

		__e.TailApply(tmp85250, tmp85254, MakeNumber(1), V1166)
		return

	}, 1)

	tmp85255 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mk_1pvar, tmp85249)

	_ = tmp85255

	tmp85256 := MakeNative(func(__e *ControlFlow) {
		V1168 := __e.Get(1)
		_ = V1168
		tmp85265 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		tmp85266 := Call(__e, tmp85265, V1168)

		if True == tmp85266 {
			tmp85259 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp85260 := MakeNative(func(__e *ControlFlow) {
				tmp85261 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(tmp85261, V1168, MakeNumber(0))
				return

			}, 0)

			tmp85262 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1pvar)
				return
			}, 1)

			tmp85263 := Call(__e, PrimNS1Value(symtry_1catch), tmp85260, tmp85262)

			tmp85264 := Call(__e, tmp85259, tmp85263, symshen_4pvar)

			if True == tmp85264 {
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

	tmp85267 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pvar_2, tmp85256)

	_ = tmp85267

	tmp85268 := MakeNative(func(__e *ControlFlow) {
		V1172 := __e.Get(1)
		_ = V1172
		V1173 := __e.Get(2)
		_ = V1173
		V1174 := __e.Get(3)
		_ = V1174
		tmp85269 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp85270 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp85271 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp85272 := Call(__e, tmp85271, V1172, MakeNumber(1))

			__e.TailApply(tmp85270, Vector, tmp85272, V1173)
			return

		}, 1)

		tmp85273 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85274 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85275 := Call(__e, tmp85274, symshen_4_dprologvectors_d)

		tmp85276 := Call(__e, tmp85273, tmp85275, V1174)

		__e.TailApply(tmp85269, tmp85276)
		return

	}, 3)

	tmp85277 := Call(__e, PrimNS1Value(symns2_1set), symshen_4bindv, tmp85268)

	_ = tmp85277

	tmp85278 := MakeNative(func(__e *ControlFlow) {
		V1177 := __e.Get(1)
		_ = V1177
		V1178 := __e.Get(2)
		_ = V1178
		tmp85279 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp85280 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp85281 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp85282 := Call(__e, tmp85281, V1177, MakeNumber(1))

			__e.TailApply(tmp85280, Vector, tmp85282, symshen_4_1null_1)
			return

		}, 1)

		tmp85283 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85284 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85285 := Call(__e, tmp85284, symshen_4_dprologvectors_d)

		tmp85286 := Call(__e, tmp85283, tmp85285, V1178)

		__e.TailApply(tmp85279, tmp85286)
		return

	}, 2)

	tmp85287 := Call(__e, PrimNS1Value(symns2_1set), symshen_4unbindv, tmp85278)

	_ = tmp85287

	tmp85288 := MakeNative(func(__e *ControlFlow) {
		tmp85289 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp85290 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85291 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85292 := Call(__e, tmp85291, symshen_4_dinfs_d)

		tmp85293 := Call(__e, tmp85290, MakeNumber(1), tmp85292)

		__e.TailApply(tmp85289, symshen_4_dinfs_d, tmp85293)
		return

	}, 0)

	tmp85294 := Call(__e, PrimNS1Value(symns2_1set), symshen_4incinfs, tmp85288)

	_ = tmp85294

	tmp85295 := MakeNative(func(__e *ControlFlow) {
		V1182 := __e.Get(1)
		_ = V1182
		V1183 := __e.Get(2)
		_ = V1183
		V1184 := __e.Get(3)
		_ = V1184
		tmp85354 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85355 := Call(__e, tmp85354, V1182)

		var ifres85342 Obj

		if True == tmp85355 {
			tmp85350 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85351 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85352 := Call(__e, tmp85351, V1182)

			tmp85353 := Call(__e, tmp85350, tmp85352)

			var ifres85344 Obj

			if True == tmp85353 {
				tmp85346 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp85347 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85348 := Call(__e, tmp85347, V1182)

				tmp85349 := Call(__e, tmp85346, Nil, tmp85348)

				var ifres85345 Obj

				if True == tmp85349 {
					ifres85345 = True

				} else {
					ifres85345 = False

				}

				ifres85344 = ifres85345

			} else {
				ifres85344 = False

			}

			var ifres85343 Obj

			if True == ifres85344 {
				ifres85343 = True

			} else {
				ifres85343 = False

			}

			ifres85342 = ifres85343

		} else {
			ifres85342 = False

		}

		if True == ifres85342 {
			tmp85327 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85328 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85329 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85330 := Call(__e, tmp85329, V1182)

			tmp85331 := Call(__e, tmp85328, tmp85330)

			tmp85332 := Call(__e, PrimNS1Value(symns2_1value), symappend)

			tmp85333 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85334 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85335 := Call(__e, tmp85334, V1182)

			tmp85336 := Call(__e, tmp85333, tmp85335)

			tmp85337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85339 := Call(__e, tmp85338, V1184, Nil)

			tmp85340 := Call(__e, tmp85337, V1183, tmp85339)

			tmp85341 := Call(__e, tmp85332, tmp85336, tmp85340)

			__e.TailApply(tmp85327, tmp85331, tmp85341)
			return

		} else {
			tmp85325 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85326 := Call(__e, tmp85325, V1182)

			var ifres85319 Obj

			if True == tmp85326 {
				tmp85321 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85322 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85323 := Call(__e, tmp85322, V1182)

				tmp85324 := Call(__e, tmp85321, tmp85323)

				var ifres85320 Obj

				if True == tmp85324 {
					ifres85320 = True

				} else {
					ifres85320 = False

				}

				ifres85319 = ifres85320

			} else {
				ifres85319 = False

			}

			if True == ifres85319 {
				tmp85299 := MakeNative(func(__e *ControlFlow) {
					NewContinuation := __e.Get(1)
					_ = NewContinuation
					tmp85300 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp85301 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85302 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85303 := Call(__e, tmp85302, V1182)

					tmp85304 := Call(__e, tmp85301, tmp85303)

					tmp85305 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp85306 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85307 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85308 := Call(__e, tmp85307, V1182)

					tmp85309 := Call(__e, tmp85306, tmp85308)

					tmp85310 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp85311 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp85312 := Call(__e, tmp85311, NewContinuation, Nil)

					tmp85313 := Call(__e, tmp85310, V1183, tmp85312)

					tmp85314 := Call(__e, tmp85305, tmp85309, tmp85313)

					__e.TailApply(tmp85300, tmp85304, tmp85314)
					return

				}, 1)

				tmp85315 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newcontinuation)

				tmp85316 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85317 := Call(__e, tmp85316, V1182)

				tmp85318 := Call(__e, tmp85315, tmp85317, V1183, V1184)

				__e.TailApply(tmp85299, tmp85318)
				return

			} else {
				tmp85298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp85298, symshen_4call__the__continuation)
				return

			}

		}

	}, 3)

	tmp85356 := Call(__e, PrimNS1Value(symns2_1set), symshen_4call__the__continuation, tmp85295)

	_ = tmp85356

	tmp85357 := MakeNative(func(__e *ControlFlow) {
		V1188 := __e.Get(1)
		_ = V1188
		V1189 := __e.Get(2)
		_ = V1189
		V1190 := __e.Get(3)
		_ = V1190
		tmp85392 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85393 := Call(__e, tmp85392, Nil, V1188)

		if True == tmp85393 {
			__e.Return(V1190)
			return
		} else {
			tmp85390 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85391 := Call(__e, tmp85390, V1188)

			var ifres85384 Obj

			if True == tmp85391 {
				tmp85386 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85387 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85388 := Call(__e, tmp85387, V1188)

				tmp85389 := Call(__e, tmp85386, tmp85388)

				var ifres85385 Obj

				if True == tmp85389 {
					ifres85385 = True

				} else {
					ifres85385 = False

				}

				ifres85384 = ifres85385

			} else {
				ifres85384 = False

			}

			if True == ifres85384 {
				tmp85361 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85362 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85363 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85364 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85366 := Call(__e, tmp85365, V1188)

				tmp85367 := Call(__e, tmp85364, tmp85366)

				tmp85368 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				tmp85369 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85370 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85371 := Call(__e, tmp85370, V1188)

				tmp85372 := Call(__e, tmp85369, tmp85371)

				tmp85373 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85375 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newcontinuation)

				tmp85376 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85377 := Call(__e, tmp85376, V1188)

				tmp85378 := Call(__e, tmp85375, tmp85377, V1189, V1190)

				tmp85379 := Call(__e, tmp85374, tmp85378, Nil)

				tmp85380 := Call(__e, tmp85373, V1189, tmp85379)

				tmp85381 := Call(__e, tmp85368, tmp85372, tmp85380)

				tmp85382 := Call(__e, tmp85363, tmp85367, tmp85381)

				tmp85383 := Call(__e, tmp85362, tmp85382, Nil)

				__e.TailApply(tmp85361, symfreeze, tmp85383)
				return

			} else {
				tmp85360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp85360, symshen_4newcontinuation)
				return

			}

		}

	}, 3)

	tmp85394 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newcontinuation, tmp85357)

	_ = tmp85394

	tmp85395 := MakeNative(func(__e *ControlFlow) {
		V1198 := __e.Get(1)
		_ = V1198
		V1199 := __e.Get(2)
		_ = V1199
		V1200 := __e.Get(3)
		_ = V1200
		tmp85396 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

		__e.TailApply(tmp85396, V1198, V1199)
		return

	}, 3)

	tmp85397 := Call(__e, PrimNS1Value(symns2_1set), symreturn, tmp85395)

	_ = tmp85397

	tmp85398 := MakeNative(func(__e *ControlFlow) {
		V1208 := __e.Get(1)
		_ = V1208
		V1209 := __e.Get(2)
		_ = V1209
		V1210 := __e.Get(3)
		_ = V1210
		tmp85399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp85400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp85401 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85402 := Call(__e, tmp85401, symshen_4_dinfs_d)

		tmp85403 := Call(__e, tmp85400, tmp85402, MakeString(" inferences\n"), symshen_4a)

		tmp85404 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp85405 := Call(__e, tmp85404)

		tmp85406 := Call(__e, tmp85399, tmp85403, tmp85405)

		_ = tmp85406

		tmp85407 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

		__e.TailApply(tmp85407, V1208, V1209)
		return

	}, 3)

	tmp85408 := Call(__e, PrimNS1Value(symns2_1set), symshen_4measure_ereturn, tmp85398)

	_ = tmp85408

	tmp85409 := MakeNative(func(__e *ControlFlow) {
		V1215 := __e.Get(1)
		_ = V1215
		V1216 := __e.Get(2)
		_ = V1216
		V1217 := __e.Get(3)
		_ = V1217
		V1218 := __e.Get(4)
		_ = V1218
		tmp85410 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a)

		tmp85411 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp85412 := Call(__e, tmp85411, V1215, V1217)

		tmp85413 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp85414 := Call(__e, tmp85413, V1216, V1217)

		__e.TailApply(tmp85410, tmp85412, tmp85414, V1217, V1218)
		return

	}, 4)

	tmp85415 := Call(__e, PrimNS1Value(symns2_1set), symunify, tmp85409)

	_ = tmp85415

	tmp85416 := MakeNative(func(__e *ControlFlow) {
		V1240 := __e.Get(1)
		_ = V1240
		V1241 := __e.Get(2)
		_ = V1241
		V1242 := __e.Get(3)
		_ = V1242
		V1243 := __e.Get(4)
		_ = V1243
		tmp85453 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85454 := Call(__e, tmp85453, V1241, V1240)

		if True == tmp85454 {
			tmp85452 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

			__e.TailApply(tmp85452, V1243)
			return

		} else {
			tmp85450 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

			tmp85451 := Call(__e, tmp85450, V1240)

			if True == tmp85451 {
				tmp85449 := Call(__e, PrimNS1Value(symns2_1value), symbind)

				__e.TailApply(tmp85449, V1240, V1241, V1242, V1243)
				return

			} else {
				tmp85447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

				tmp85448 := Call(__e, tmp85447, V1241)

				if True == tmp85448 {
					tmp85446 := Call(__e, PrimNS1Value(symns2_1value), symbind)

					__e.TailApply(tmp85446, V1241, V1240, V1242, V1243)
					return

				} else {
					tmp85444 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp85445 := Call(__e, tmp85444, V1240)

					var ifres85440 Obj

					if True == tmp85445 {
						tmp85442 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp85443 := Call(__e, tmp85442, V1241)

						var ifres85441 Obj

						if True == tmp85443 {
							ifres85441 = True

						} else {
							ifres85441 = False

						}

						ifres85440 = ifres85441

					} else {
						ifres85440 = False

					}

					if True == ifres85440 {
						tmp85421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a)

						tmp85422 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp85423 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85424 := Call(__e, tmp85423, V1240)

						tmp85425 := Call(__e, tmp85422, tmp85424, V1242)

						tmp85426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp85427 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85428 := Call(__e, tmp85427, V1241)

						tmp85429 := Call(__e, tmp85426, tmp85428, V1242)

						tmp85430 := MakeNative(func(__e *ControlFlow) {
							tmp85431 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a)

							tmp85432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp85433 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85434 := Call(__e, tmp85433, V1240)

							tmp85435 := Call(__e, tmp85432, tmp85434, V1242)

							tmp85436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp85437 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85438 := Call(__e, tmp85437, V1241)

							tmp85439 := Call(__e, tmp85436, tmp85438, V1242)

							__e.TailApply(tmp85431, tmp85435, tmp85439, V1242, V1243)
							return

						}, 0)

						__e.TailApply(tmp85421, tmp85425, tmp85429, V1242, tmp85430)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp85455 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lzy_a, tmp85416)

	_ = tmp85455

	tmp85456 := MakeNative(func(__e *ControlFlow) {
		V1246 := __e.Get(1)
		_ = V1246
		V1247 := __e.Get(2)
		_ = V1247
		tmp85477 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85478 := Call(__e, tmp85477, V1246)

		if True == tmp85478 {
			tmp85468 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85469 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp85470 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85471 := Call(__e, tmp85470, V1246)

			tmp85472 := Call(__e, tmp85469, tmp85471, V1247)

			tmp85473 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp85474 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85475 := Call(__e, tmp85474, V1246)

			tmp85476 := Call(__e, tmp85473, tmp85475, V1247)

			__e.TailApply(tmp85468, tmp85472, tmp85476)
			return

		} else {
			tmp85466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

			tmp85467 := Call(__e, tmp85466, V1246)

			if True == tmp85467 {
				tmp85459 := MakeNative(func(__e *ControlFlow) {
					Value := __e.Get(1)
					_ = Value
					tmp85462 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp85463 := Call(__e, tmp85462, Value, symshen_4_1null_1)

					if True == tmp85463 {
						__e.Return(V1246)
						return
					} else {
						tmp85461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

						__e.TailApply(tmp85461, Value, V1247)
						return

					}

				}, 1)

				tmp85464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4valvector)

				tmp85465 := Call(__e, tmp85464, V1246, V1247)

				__e.TailApply(tmp85459, tmp85465)
				return

			} else {
				__e.Return(V1246)
				return
			}

		}

	}, 2)

	tmp85479 := Call(__e, PrimNS1Value(symns2_1set), symshen_4deref, tmp85456)

	_ = tmp85479

	tmp85480 := MakeNative(func(__e *ControlFlow) {
		V1250 := __e.Get(1)
		_ = V1250
		V1251 := __e.Get(2)
		_ = V1251
		tmp85489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

		tmp85490 := Call(__e, tmp85489, V1250)

		if True == tmp85490 {
			tmp85482 := MakeNative(func(__e *ControlFlow) {
				Value := __e.Get(1)
				_ = Value
				tmp85485 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp85486 := Call(__e, tmp85485, Value, symshen_4_1null_1)

				if True == tmp85486 {
					__e.Return(V1250)
					return
				} else {
					tmp85484 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					__e.TailApply(tmp85484, Value, V1251)
					return

				}

			}, 1)

			tmp85487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4valvector)

			tmp85488 := Call(__e, tmp85487, V1250, V1251)

			__e.TailApply(tmp85482, tmp85488)
			return

		} else {
			__e.Return(V1250)
			return
		}

	}, 2)

	tmp85491 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lazyderef, tmp85480)

	_ = tmp85491

	tmp85492 := MakeNative(func(__e *ControlFlow) {
		V1254 := __e.Get(1)
		_ = V1254
		V1255 := __e.Get(2)
		_ = V1255
		tmp85493 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85494 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85495 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85496 := Call(__e, tmp85495, symshen_4_dprologvectors_d)

		tmp85497 := Call(__e, tmp85494, tmp85496, V1255)

		tmp85498 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85499 := Call(__e, tmp85498, V1254, MakeNumber(1))

		__e.TailApply(tmp85493, tmp85497, tmp85499)
		return

	}, 2)

	tmp85500 := Call(__e, PrimNS1Value(symns2_1set), symshen_4valvector, tmp85492)

	_ = tmp85500

	tmp85501 := MakeNative(func(__e *ControlFlow) {
		V1260 := __e.Get(1)
		_ = V1260
		V1261 := __e.Get(2)
		_ = V1261
		V1262 := __e.Get(3)
		_ = V1262
		V1263 := __e.Get(4)
		_ = V1263
		tmp85502 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a_b)

		tmp85503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp85504 := Call(__e, tmp85503, V1260, V1262)

		tmp85505 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp85506 := Call(__e, tmp85505, V1261, V1262)

		__e.TailApply(tmp85502, tmp85504, tmp85506, V1262, V1263)
		return

	}, 4)

	tmp85507 := Call(__e, PrimNS1Value(symns2_1set), symunify_b, tmp85501)

	_ = tmp85507

	tmp85508 := MakeNative(func(__e *ControlFlow) {
		V1285 := __e.Get(1)
		_ = V1285
		V1286 := __e.Get(2)
		_ = V1286
		V1287 := __e.Get(3)
		_ = V1287
		V1288 := __e.Get(4)
		_ = V1288
		tmp85561 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85562 := Call(__e, tmp85561, V1286, V1285)

		if True == tmp85562 {
			tmp85560 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

			__e.TailApply(tmp85560, V1288)
			return

		} else {
			tmp85558 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

			tmp85559 := Call(__e, tmp85558, V1285)

			var ifres85550 Obj

			if True == tmp85559 {
				tmp85552 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp85553 := Call(__e, PrimNS1Value(symns2_1value), symshen_4occurs_2)

				tmp85554 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

				tmp85555 := Call(__e, tmp85554, V1286, V1287)

				tmp85556 := Call(__e, tmp85553, V1285, tmp85555)

				tmp85557 := Call(__e, tmp85552, tmp85556)

				var ifres85551 Obj

				if True == tmp85557 {
					ifres85551 = True

				} else {
					ifres85551 = False

				}

				ifres85550 = ifres85551

			} else {
				ifres85550 = False

			}

			if True == ifres85550 {
				tmp85549 := Call(__e, PrimNS1Value(symns2_1value), symbind)

				__e.TailApply(tmp85549, V1285, V1286, V1287, V1288)
				return

			} else {
				tmp85547 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

				tmp85548 := Call(__e, tmp85547, V1286)

				var ifres85539 Obj

				if True == tmp85548 {
					tmp85541 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp85542 := Call(__e, PrimNS1Value(symns2_1value), symshen_4occurs_2)

					tmp85543 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

					tmp85544 := Call(__e, tmp85543, V1285, V1287)

					tmp85545 := Call(__e, tmp85542, V1286, tmp85544)

					tmp85546 := Call(__e, tmp85541, tmp85545)

					var ifres85540 Obj

					if True == tmp85546 {
						ifres85540 = True

					} else {
						ifres85540 = False

					}

					ifres85539 = ifres85540

				} else {
					ifres85539 = False

				}

				if True == ifres85539 {
					tmp85538 := Call(__e, PrimNS1Value(symns2_1value), symbind)

					__e.TailApply(tmp85538, V1286, V1285, V1287, V1288)
					return

				} else {
					tmp85536 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp85537 := Call(__e, tmp85536, V1285)

					var ifres85532 Obj

					if True == tmp85537 {
						tmp85534 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp85535 := Call(__e, tmp85534, V1286)

						var ifres85533 Obj

						if True == tmp85535 {
							ifres85533 = True

						} else {
							ifres85533 = False

						}

						ifres85532 = ifres85533

					} else {
						ifres85532 = False

					}

					if True == ifres85532 {
						tmp85513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a_b)

						tmp85514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp85515 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85516 := Call(__e, tmp85515, V1285)

						tmp85517 := Call(__e, tmp85514, tmp85516, V1287)

						tmp85518 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp85519 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85520 := Call(__e, tmp85519, V1286)

						tmp85521 := Call(__e, tmp85518, tmp85520, V1287)

						tmp85522 := MakeNative(func(__e *ControlFlow) {
							tmp85523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a_b)

							tmp85524 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp85525 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85526 := Call(__e, tmp85525, V1285)

							tmp85527 := Call(__e, tmp85524, tmp85526, V1287)

							tmp85528 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp85529 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp85530 := Call(__e, tmp85529, V1286)

							tmp85531 := Call(__e, tmp85528, tmp85530, V1287)

							__e.TailApply(tmp85523, tmp85527, tmp85531, V1287, V1288)
							return

						}, 0)

						__e.TailApply(tmp85513, tmp85517, tmp85521, V1287, tmp85522)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp85563 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lzy_a_b, tmp85508)

	_ = tmp85563

	tmp85564 := MakeNative(func(__e *ControlFlow) {
		V1300 := __e.Get(1)
		_ = V1300
		V1301 := __e.Get(2)
		_ = V1301
		tmp85579 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85580 := Call(__e, tmp85579, V1301, V1300)

		if True == tmp85580 {
			__e.Return(True)
			return
		} else {
			tmp85577 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85578 := Call(__e, tmp85577, V1301)

			if True == tmp85578 {
				tmp85573 := Call(__e, PrimNS1Value(symns2_1value), symshen_4occurs_2)

				tmp85574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85575 := Call(__e, tmp85574, V1301)

				tmp85576 := Call(__e, tmp85573, V1300, tmp85575)

				if True == tmp85576 {
					__e.Return(True)
					return
				} else {
					tmp85569 := Call(__e, PrimNS1Value(symns2_1value), symshen_4occurs_2)

					tmp85570 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85571 := Call(__e, tmp85570, V1301)

					tmp85572 := Call(__e, tmp85569, V1300, tmp85571)

					if True == tmp85572 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp85581 := Call(__e, PrimNS1Value(symns2_1set), symshen_4occurs_2, tmp85564)

	_ = tmp85581

	tmp85582 := MakeNative(func(__e *ControlFlow) {
		V1306 := __e.Get(1)
		_ = V1306
		V1307 := __e.Get(2)
		_ = V1307
		V1308 := __e.Get(3)
		_ = V1308
		V1309 := __e.Get(4)
		_ = V1309
		tmp85583 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a_a)

		tmp85584 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp85585 := Call(__e, tmp85584, V1306, V1308)

		tmp85586 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp85587 := Call(__e, tmp85586, V1307, V1308)

		__e.TailApply(tmp85583, tmp85585, tmp85587, V1308, V1309)
		return

	}, 4)

	tmp85588 := Call(__e, PrimNS1Value(symns2_1set), symidentical, tmp85582)

	_ = tmp85588

	tmp85589 := MakeNative(func(__e *ControlFlow) {
		V1331 := __e.Get(1)
		_ = V1331
		V1332 := __e.Get(2)
		_ = V1332
		V1333 := __e.Get(3)
		_ = V1333
		V1334 := __e.Get(4)
		_ = V1334
		tmp85614 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85615 := Call(__e, tmp85614, V1332, V1331)

		if True == tmp85615 {
			tmp85613 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

			__e.TailApply(tmp85613, V1334)
			return

		} else {
			tmp85611 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85612 := Call(__e, tmp85611, V1331)

			var ifres85607 Obj

			if True == tmp85612 {
				tmp85609 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85610 := Call(__e, tmp85609, V1332)

				var ifres85608 Obj

				if True == tmp85610 {
					ifres85608 = True

				} else {
					ifres85608 = False

				}

				ifres85607 = ifres85608

			} else {
				ifres85607 = False

			}

			if True == ifres85607 {
				tmp85592 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a_a)

				tmp85593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp85594 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85595 := Call(__e, tmp85594, V1331)

				tmp85596 := Call(__e, tmp85593, tmp85595, V1333)

				tmp85597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp85598 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85599 := Call(__e, tmp85598, V1332)

				tmp85600 := Call(__e, tmp85597, tmp85599, V1333)

				tmp85601 := MakeNative(func(__e *ControlFlow) {
					tmp85602 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lzy_a_a)

					tmp85603 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85604 := Call(__e, tmp85603, V1331)

					tmp85605 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85606 := Call(__e, tmp85605, V1332)

					__e.TailApply(tmp85602, tmp85604, tmp85606, V1333, V1334)
					return

				}, 0)

				__e.TailApply(tmp85592, tmp85596, tmp85600, V1333, tmp85601)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 4)

	tmp85616 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lzy_a_a, tmp85589)

	_ = tmp85616

	tmp85617 := MakeNative(func(__e *ControlFlow) {
		V1336 := __e.Get(1)
		_ = V1336
		tmp85618 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp85619 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp85620 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp85621 := Call(__e, tmp85620, V1336, MakeNumber(1))

		tmp85622 := Call(__e, tmp85619, tmp85621, MakeString(""), symshen_4a)

		__e.TailApply(tmp85618, MakeString("Var"), tmp85622)
		return

	}, 1)

	tmp85623 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pvar, tmp85617)

	_ = tmp85623

	tmp85624 := MakeNative(func(__e *ControlFlow) {
		V1341 := __e.Get(1)
		_ = V1341
		V1342 := __e.Get(2)
		_ = V1342
		V1343 := __e.Get(3)
		_ = V1343
		V1344 := __e.Get(4)
		_ = V1344
		tmp85625 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

		tmp85626 := Call(__e, tmp85625, V1341, V1342, V1343)

		_ = tmp85626

		tmp85627 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp85628 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

			tmp85629 := Call(__e, tmp85628, V1341, V1343)

			_ = tmp85629

			__e.Return(Result)
			return

		}, 1)

		tmp85630 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

		tmp85631 := Call(__e, tmp85630, V1344)

		__e.TailApply(tmp85627, tmp85631)
		return

	}, 4)

	tmp85632 := Call(__e, PrimNS1Value(symns2_1set), symbind, tmp85624)

	_ = tmp85632

	tmp85633 := MakeNative(func(__e *ControlFlow) {
		V1362 := __e.Get(1)
		_ = V1362
		V1363 := __e.Get(2)
		_ = V1363
		V1364 := __e.Get(3)
		_ = V1364
		tmp85644 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85645 := Call(__e, tmp85644, True, V1362)

		if True == tmp85645 {
			tmp85643 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

			__e.TailApply(tmp85643, V1364)
			return

		} else {
			tmp85641 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp85642 := Call(__e, tmp85641, False, V1362)

			if True == tmp85642 {
				__e.Return(False)
				return
			} else {
				tmp85636 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp85637 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp85638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp85639 := Call(__e, tmp85638, V1362, MakeString("%"), symshen_4s)

				tmp85640 := Call(__e, tmp85637, MakeString("fwhen expects a boolean: not "), tmp85639)

				__e.TailApply(tmp85636, tmp85640)
				return

			}

		}

	}, 3)

	tmp85646 := Call(__e, PrimNS1Value(symns2_1set), symfwhen, tmp85633)

	_ = tmp85646

	tmp85647 := MakeNative(func(__e *ControlFlow) {
		V1380 := __e.Get(1)
		_ = V1380
		V1381 := __e.Get(2)
		_ = V1381
		V1382 := __e.Get(3)
		_ = V1382
		tmp85664 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85665 := Call(__e, tmp85664, V1380)

		if True == tmp85665 {
			tmp85655 := Call(__e, PrimNS1Value(symns2_1value), symshen_4call_1help)

			tmp85656 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

			tmp85657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

			tmp85658 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85659 := Call(__e, tmp85658, V1380)

			tmp85660 := Call(__e, tmp85657, tmp85659, V1381)

			tmp85661 := Call(__e, tmp85656, tmp85660)

			tmp85662 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85663 := Call(__e, tmp85662, V1380)

			__e.TailApply(tmp85655, tmp85661, tmp85663, V1381, V1382)
			return

		} else {
			tmp85653 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

			tmp85654 := Call(__e, tmp85653, V1380)

			if True == tmp85654 {
				tmp85650 := Call(__e, PrimNS1Value(symns2_1value), symcall)

				tmp85651 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp85652 := Call(__e, tmp85651, V1380, V1381)

				__e.TailApply(tmp85650, tmp85652, V1381, V1382)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 3)

	tmp85666 := Call(__e, PrimNS1Value(symns2_1set), symcall, tmp85647)

	_ = tmp85666

	tmp85667 := MakeNative(func(__e *ControlFlow) {
		V1387 := __e.Get(1)
		_ = V1387
		V1388 := __e.Get(2)
		_ = V1388
		V1389 := __e.Get(3)
		_ = V1389
		V1390 := __e.Get(4)
		_ = V1390
		tmp85679 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85680 := Call(__e, tmp85679, Nil, V1388)

		if True == tmp85680 {
			__e.TailApply(V1387, V1389, V1390)
			return
		} else {
			tmp85677 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85678 := Call(__e, tmp85677, V1388)

			if True == tmp85678 {
				tmp85671 := Call(__e, PrimNS1Value(symns2_1value), symshen_4call_1help)

				tmp85672 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85673 := Call(__e, tmp85672, V1388)

				tmp85674 := Call(__e, V1387, tmp85673)

				tmp85675 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85676 := Call(__e, tmp85675, V1388)

				__e.TailApply(tmp85671, tmp85674, tmp85676, V1389, V1390)
				return

			} else {
				tmp85670 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp85670, symshen_4call_1help)
				return

			}

		}

	}, 4)

	tmp85681 := Call(__e, PrimNS1Value(symns2_1set), symshen_4call_1help, tmp85667)

	_ = tmp85681

	tmp85682 := MakeNative(func(__e *ControlFlow) {
		V1392 := __e.Get(1)
		_ = V1392
		tmp85711 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85712 := Call(__e, tmp85711, V1392)

		var ifres85705 Obj

		if True == tmp85712 {
			tmp85707 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85708 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85709 := Call(__e, tmp85708, V1392)

			tmp85710 := Call(__e, tmp85707, tmp85709)

			var ifres85706 Obj

			if True == tmp85710 {
				ifres85706 = True

			} else {
				ifres85706 = False

			}

			ifres85705 = ifres85706

		} else {
			ifres85705 = False

		}

		if True == ifres85705 {
			tmp85685 := MakeNative(func(__e *ControlFlow) {
				ProcessN := __e.Get(1)
				_ = ProcessN
				tmp85686 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intprolog_1help)

				tmp85687 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85688 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85689 := Call(__e, tmp85688, V1392)

				tmp85690 := Call(__e, tmp85687, tmp85689)

				tmp85691 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1prolog_1variables)

				tmp85692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85693 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85694 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85695 := Call(__e, tmp85694, V1392)

				tmp85696 := Call(__e, tmp85693, tmp85695)

				tmp85697 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85698 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85699 := Call(__e, tmp85698, V1392)

				tmp85700 := Call(__e, tmp85697, tmp85699, Nil)

				tmp85701 := Call(__e, tmp85692, tmp85696, tmp85700)

				tmp85702 := Call(__e, tmp85691, tmp85701, ProcessN)

				__e.TailApply(tmp85686, tmp85690, tmp85702, ProcessN)
				return

			}, 1)

			tmp85703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4start_1new_1prolog_1process)

			tmp85704 := Call(__e, tmp85703)

			__e.TailApply(tmp85685, tmp85704)
			return

		} else {
			tmp85684 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp85684, symshen_4intprolog)
			return

		}

	}, 1)

	tmp85713 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intprolog, tmp85682)

	_ = tmp85713

	tmp85714 := MakeNative(func(__e *ControlFlow) {
		V1396 := __e.Get(1)
		_ = V1396
		V1397 := __e.Get(2)
		_ = V1397
		V1398 := __e.Get(3)
		_ = V1398
		tmp85738 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85739 := Call(__e, tmp85738, V1397)

		var ifres85724 Obj

		if True == tmp85739 {
			tmp85734 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85735 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85736 := Call(__e, tmp85735, V1397)

			tmp85737 := Call(__e, tmp85734, tmp85736)

			var ifres85726 Obj

			if True == tmp85737 {
				tmp85728 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp85729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85730 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85731 := Call(__e, tmp85730, V1397)

				tmp85732 := Call(__e, tmp85729, tmp85731)

				tmp85733 := Call(__e, tmp85728, Nil, tmp85732)

				var ifres85727 Obj

				if True == tmp85733 {
					ifres85727 = True

				} else {
					ifres85727 = False

				}

				ifres85726 = ifres85727

			} else {
				ifres85726 = False

			}

			var ifres85725 Obj

			if True == ifres85726 {
				ifres85725 = True

			} else {
				ifres85725 = False

			}

			ifres85724 = ifres85725

		} else {
			ifres85724 = False

		}

		if True == ifres85724 {
			tmp85717 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intprolog_1help_1help)

			tmp85718 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85719 := Call(__e, tmp85718, V1397)

			tmp85720 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp85721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp85722 := Call(__e, tmp85721, V1397)

			tmp85723 := Call(__e, tmp85720, tmp85722)

			__e.TailApply(tmp85717, V1396, tmp85719, tmp85723, V1398)
			return

		} else {
			tmp85716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp85716, symshen_4intprolog_1help)
			return

		}

	}, 3)

	tmp85740 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intprolog_1help, tmp85714)

	_ = tmp85740

	tmp85741 := MakeNative(func(__e *ControlFlow) {
		V1403 := __e.Get(1)
		_ = V1403
		V1404 := __e.Get(2)
		_ = V1404
		V1405 := __e.Get(3)
		_ = V1405
		V1406 := __e.Get(4)
		_ = V1406
		tmp85755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85756 := Call(__e, tmp85755, Nil, V1404)

		if True == tmp85756 {
			tmp85753 := MakeNative(func(__e *ControlFlow) {
				tmp85754 := Call(__e, PrimNS1Value(symns2_1value), symshen_4call_1rest)

				__e.TailApply(tmp85754, V1405, V1406)
				return

			}, 0)

			__e.TailApply(V1403, V1406, tmp85753)
			return

		} else {
			tmp85751 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85752 := Call(__e, tmp85751, V1404)

			if True == tmp85752 {
				tmp85745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intprolog_1help_1help)

				tmp85746 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85747 := Call(__e, tmp85746, V1404)

				tmp85748 := Call(__e, V1403, tmp85747)

				tmp85749 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85750 := Call(__e, tmp85749, V1404)

				__e.TailApply(tmp85745, tmp85748, tmp85750, V1405, V1406)
				return

			} else {
				tmp85744 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp85744, symshen_4intprolog_1help_1help)
				return

			}

		}

	}, 4)

	tmp85757 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intprolog_1help_1help, tmp85741)

	_ = tmp85757

	tmp85758 := MakeNative(func(__e *ControlFlow) {
		V1411 := __e.Get(1)
		_ = V1411
		V1412 := __e.Get(2)
		_ = V1412
		tmp85827 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85828 := Call(__e, tmp85827, Nil, V1411)

		if True == tmp85828 {
			__e.Return(True)
			return
		} else {
			tmp85825 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85826 := Call(__e, tmp85825, V1411)

			var ifres85811 Obj

			if True == tmp85826 {
				tmp85821 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85822 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85823 := Call(__e, tmp85822, V1411)

				tmp85824 := Call(__e, tmp85821, tmp85823)

				var ifres85813 Obj

				if True == tmp85824 {
					tmp85815 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp85816 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85818 := Call(__e, tmp85817, V1411)

					tmp85819 := Call(__e, tmp85816, tmp85818)

					tmp85820 := Call(__e, tmp85815, tmp85819)

					var ifres85814 Obj

					if True == tmp85820 {
						ifres85814 = True

					} else {
						ifres85814 = False

					}

					ifres85813 = ifres85814

				} else {
					ifres85813 = False

				}

				var ifres85812 Obj

				if True == ifres85813 {
					ifres85812 = True

				} else {
					ifres85812 = False

				}

				ifres85811 = ifres85812

			} else {
				ifres85811 = False

			}

			if True == ifres85811 {
				tmp85787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4call_1rest)

				tmp85788 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85789 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp85790 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85791 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85792 := Call(__e, tmp85791, V1411)

				tmp85793 := Call(__e, tmp85790, tmp85792)

				tmp85794 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85796 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85797 := Call(__e, tmp85796, V1411)

				tmp85798 := Call(__e, tmp85795, tmp85797)

				tmp85799 := Call(__e, tmp85794, tmp85798)

				tmp85800 := Call(__e, tmp85793, tmp85799)

				tmp85801 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85802 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85803 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85804 := Call(__e, tmp85803, V1411)

				tmp85805 := Call(__e, tmp85802, tmp85804)

				tmp85806 := Call(__e, tmp85801, tmp85805)

				tmp85807 := Call(__e, tmp85789, tmp85800, tmp85806)

				tmp85808 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp85809 := Call(__e, tmp85808, V1411)

				tmp85810 := Call(__e, tmp85788, tmp85807, tmp85809)

				__e.TailApply(tmp85787, tmp85810, V1412)
				return

			} else {
				tmp85785 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85786 := Call(__e, tmp85785, V1411)

				var ifres85771 Obj

				if True == tmp85786 {
					tmp85781 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp85782 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85783 := Call(__e, tmp85782, V1411)

					tmp85784 := Call(__e, tmp85781, tmp85783)

					var ifres85773 Obj

					if True == tmp85784 {
						tmp85775 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp85776 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp85777 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85778 := Call(__e, tmp85777, V1411)

						tmp85779 := Call(__e, tmp85776, tmp85778)

						tmp85780 := Call(__e, tmp85775, Nil, tmp85779)

						var ifres85774 Obj

						if True == tmp85780 {
							ifres85774 = True

						} else {
							ifres85774 = False

						}

						ifres85773 = ifres85774

					} else {
						ifres85773 = False

					}

					var ifres85772 Obj

					if True == ifres85773 {
						ifres85772 = True

					} else {
						ifres85772 = False

					}

					ifres85771 = ifres85772

				} else {
					ifres85771 = False

				}

				if True == ifres85771 {
					tmp85763 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85764 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85765 := Call(__e, tmp85764, V1411)

					tmp85766 := Call(__e, tmp85763, tmp85765)

					tmp85767 := MakeNative(func(__e *ControlFlow) {
						tmp85768 := Call(__e, PrimNS1Value(symns2_1value), symshen_4call_1rest)

						tmp85769 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp85770 := Call(__e, tmp85769, V1411)

						__e.TailApply(tmp85768, tmp85770, V1412)
						return

					}, 0)

					__e.TailApply(tmp85766, V1412, tmp85767)
					return

				} else {
					tmp85762 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp85762, symshen_4call_1rest)
					return

				}

			}

		}

	}, 2)

	tmp85829 := Call(__e, PrimNS1Value(symns2_1set), symshen_4call_1rest, tmp85758)

	_ = tmp85829

	tmp85830 := MakeNative(func(__e *ControlFlow) {
		tmp85831 := MakeNative(func(__e *ControlFlow) {
			IncrementProcessCounter := __e.Get(1)
			_ = IncrementProcessCounter
			tmp85832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4initialise_1prolog)

			__e.TailApply(tmp85832, IncrementProcessCounter)
			return

		}, 1)

		tmp85833 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp85834 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp85835 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85836 := Call(__e, tmp85835, symshen_4_dprocess_1counter_d)

		tmp85837 := Call(__e, tmp85834, MakeNumber(1), tmp85836)

		tmp85838 := Call(__e, tmp85833, symshen_4_dprocess_1counter_d, tmp85837)

		__e.TailApply(tmp85831, tmp85838)
		return

	}, 0)

	tmp85839 := Call(__e, PrimNS1Value(symns2_1set), symshen_4start_1new_1prolog_1process, tmp85830)

	_ = tmp85839

	tmp85840 := MakeNative(func(__e *ControlFlow) {
		V1415 := __e.Get(1)
		_ = V1415
		V1416 := __e.Get(2)
		_ = V1416
		tmp85841 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1prolog_1variables_1help)

		tmp85842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4flatten)

		tmp85843 := Call(__e, tmp85842, V1415)

		__e.TailApply(tmp85841, V1415, tmp85843, V1416)
		return

	}, 2)

	tmp85844 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1prolog_1variables, tmp85840)

	_ = tmp85844

	tmp85845 := MakeNative(func(__e *ControlFlow) {
		V1424 := __e.Get(1)
		_ = V1424
		V1425 := __e.Get(2)
		_ = V1425
		V1426 := __e.Get(3)
		_ = V1426
		tmp85879 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85880 := Call(__e, tmp85879, Nil, V1425)

		if True == tmp85880 {
			__e.Return(V1424)
			return
		} else {
			tmp85877 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp85878 := Call(__e, tmp85877, V1425)

			var ifres85871 Obj

			if True == tmp85878 {
				tmp85873 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				tmp85874 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp85875 := Call(__e, tmp85874, V1425)

				tmp85876 := Call(__e, tmp85873, tmp85875)

				var ifres85872 Obj

				if True == tmp85876 {
					ifres85872 = True

				} else {
					ifres85872 = False

				}

				ifres85871 = ifres85872

			} else {
				ifres85871 = False

			}

			if True == ifres85871 {
				tmp85855 := MakeNative(func(__e *ControlFlow) {
					V := __e.Get(1)
					_ = V
					tmp85856 := MakeNative(func(__e *ControlFlow) {
						XV_cY := __e.Get(1)
						_ = XV_cY
						tmp85857 := MakeNative(func(__e *ControlFlow) {
							Z_1Y := __e.Get(1)
							_ = Z_1Y
							tmp85858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1prolog_1variables_1help)

							__e.TailApply(tmp85858, XV_cY, Z_1Y, V1426)
							return

						}, 1)

						tmp85859 := Call(__e, PrimNS1Value(symns2_1value), symremove)

						tmp85860 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp85861 := Call(__e, tmp85860, V1425)

						tmp85862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp85863 := Call(__e, tmp85862, V1425)

						tmp85864 := Call(__e, tmp85859, tmp85861, tmp85863)

						__e.TailApply(tmp85857, tmp85864)
						return

					}, 1)

					tmp85865 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp85866 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp85867 := Call(__e, tmp85866, V1425)

					tmp85868 := Call(__e, tmp85865, V, tmp85867, V1424)

					__e.TailApply(tmp85856, tmp85868)
					return

				}, 1)

				tmp85869 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

				tmp85870 := Call(__e, tmp85869, V1426)

				__e.TailApply(tmp85855, tmp85870)
				return

			} else {
				tmp85853 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp85854 := Call(__e, tmp85853, V1425)

				if True == tmp85854 {
					tmp85850 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1prolog_1variables_1help)

					tmp85851 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp85852 := Call(__e, tmp85851, V1425)

					__e.TailApply(tmp85850, V1424, tmp85852, V1426)
					return

				} else {
					tmp85849 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp85849, symshen_4insert_1prolog_1variables_1help)
					return

				}

			}

		}

	}, 3)

	tmp85881 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1prolog_1variables_1help, tmp85845)

	_ = tmp85881

	tmp85882 := MakeNative(func(__e *ControlFlow) {
		V1428 := __e.Get(1)
		_ = V1428
		tmp85883 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp85884 := MakeNative(func(__e *ControlFlow) {
				Counter := __e.Get(1)
				_ = Counter
				__e.Return(V1428)
				return
			}, 1)

			tmp85885 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp85886 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp85887 := Call(__e, tmp85886, symshen_4_dvarcounter_d)

			tmp85888 := Call(__e, tmp85885, tmp85887, V1428, MakeNumber(1))

			__e.TailApply(tmp85884, tmp85888)
			return

		}, 1)

		tmp85889 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		tmp85890 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp85891 := Call(__e, tmp85890, symshen_4_dprologvectors_d)

		tmp85892 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

		tmp85893 := Call(__e, PrimNS1Value(symns2_1value), symvector)

		tmp85894 := Call(__e, tmp85893, MakeNumber(10))

		tmp85895 := Call(__e, tmp85892, tmp85894, MakeNumber(1), MakeNumber(10), symshen_4_1null_1)

		tmp85896 := Call(__e, tmp85889, tmp85891, V1428, tmp85895)

		__e.TailApply(tmp85883, tmp85896)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4initialise_1prolog, tmp85882)
	return

}, 0)
