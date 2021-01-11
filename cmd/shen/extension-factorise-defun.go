package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionFactoriseDefunMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2012-2019 Bruno Deferrari.  All rights reserved.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")
	gen37522 := MakeNative(func(__e *ControlFlow) {
		V4931 := __e.Get(1)
		_ = V4931
		gen37519 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen37520 := Call(__e, gen37519, V4931)

		var gen37521 Obj

		if True == gen37520 {
			gen37514 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen37515 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37516 := Call(__e, gen37515, V4931)

			gen37517 := Call(__e, gen37514, symdefun, gen37516)

			var gen37518 Obj

			if True == gen37517 {
				gen37509 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37510 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37511 := Call(__e, gen37510, V4931)

				gen37512 := Call(__e, gen37509, gen37511)

				var gen37513 Obj

				if True == gen37512 {
					gen37502 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37503 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37504 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37505 := Call(__e, gen37504, V4931)

					gen37506 := Call(__e, gen37503, gen37505)

					gen37507 := Call(__e, gen37502, gen37506)

					var gen37508 Obj

					if True == gen37507 {
						gen37493 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37494 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37495 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37496 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37497 := Call(__e, gen37496, V4931)

						gen37498 := Call(__e, gen37495, gen37497)

						gen37499 := Call(__e, gen37494, gen37498)

						gen37500 := Call(__e, gen37493, gen37499)

						var gen37501 Obj

						if True == gen37500 {
							gen37482 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37483 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen37484 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37486 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37487 := Call(__e, gen37486, V4931)

							gen37488 := Call(__e, gen37485, gen37487)

							gen37489 := Call(__e, gen37484, gen37488)

							gen37490 := Call(__e, gen37483, gen37489)

							gen37491 := Call(__e, gen37482, gen37490)

							var gen37492 Obj

							if True == gen37491 {
								gen37469 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen37470 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37471 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37472 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37473 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37474 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37475 := Call(__e, gen37474, V4931)

								gen37476 := Call(__e, gen37473, gen37475)

								gen37477 := Call(__e, gen37472, gen37476)

								gen37478 := Call(__e, gen37471, gen37477)

								gen37479 := Call(__e, gen37470, gen37478)

								gen37480 := Call(__e, gen37469, symcond, gen37479)

								var gen37481 Obj

								if True == gen37480 {
									gen37459 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen37460 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37461 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37462 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37464 := Call(__e, gen37463, V4931)

									gen37465 := Call(__e, gen37462, gen37464)

									gen37466 := Call(__e, gen37461, gen37465)

									gen37467 := Call(__e, gen37460, gen37466)

									gen37468 := Call(__e, gen37459, Nil, gen37467)

									if True == gen37468 {
										gen37481 = True
									} else {
										gen37481 = False
									}

								} else {
									gen37481 = False
								}

								if True == gen37481 {
									gen37492 = True
								} else {
									gen37492 = False
								}

							} else {
								gen37492 = False
							}

							if True == gen37492 {
								gen37501 = True
							} else {
								gen37501 = False
							}

						} else {
							gen37501 = False
						}

						if True == gen37501 {
							gen37508 = True
						} else {
							gen37508 = False
						}

					} else {
						gen37508 = False
					}

					if True == gen37508 {
						gen37513 = True
					} else {
						gen37513 = False
					}

				} else {
					gen37513 = False
				}

				if True == gen37513 {
					gen37518 = True
				} else {
					gen37518 = False
				}

			} else {
				gen37518 = False
			}

			if True == gen37518 {
				gen37521 = True
			} else {
				gen37521 = False
			}

		} else {
			gen37521 = False
		}

		if True == gen37521 {
			gen37418 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen37419 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen37420 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37421 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37422 := Call(__e, gen37421, V4931)

			gen37423 := Call(__e, gen37420, gen37422)

			gen37424 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen37425 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37426 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37427 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37428 := Call(__e, gen37427, V4931)

			gen37429 := Call(__e, gen37426, gen37428)

			gen37430 := Call(__e, gen37425, gen37429)

			gen37431 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen37432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4factorise_1cond)

			gen37433 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37434 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37436 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37437 := Call(__e, gen37436, V4931)

			gen37438 := Call(__e, gen37435, gen37437)

			gen37439 := Call(__e, gen37434, gen37438)

			gen37440 := Call(__e, gen37433, gen37439)

			gen37441 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen37442 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen37443 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37444 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37445 := Call(__e, gen37444, V4931)

			gen37446 := Call(__e, gen37443, gen37445)

			gen37447 := Call(__e, gen37442, gen37446, Nil)

			gen37448 := Call(__e, gen37441, symshen_4f__error, gen37447)

			gen37449 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37450 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37451 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37452 := Call(__e, gen37451, V4931)

			gen37453 := Call(__e, gen37450, gen37452)

			gen37454 := Call(__e, gen37449, gen37453)

			gen37455 := Call(__e, gen37432, gen37440, gen37448, gen37454)

			gen37456 := Call(__e, gen37431, gen37455, Nil)

			gen37457 := Call(__e, gen37424, gen37430, gen37456)

			gen37458 := Call(__e, gen37419, gen37423, gen37457)

			__e.TailApply(gen37418, symdefun, gen37458)

			return

		} else {
			if True == True {
				__e.Return(V4931)
				return
			} else {
				gen37417 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen37417, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4factorise_1defun, gen37522)

	gen37538 := MakeNative(func(__e *ControlFlow) {
		V4943 := __e.Get(1)
		_ = V4943
		V4944 := __e.Get(2)
		_ = V4944
		V4945 := __e.Get(3)
		_ = V4945
		gen37535 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen37536 := Call(__e, gen37535, V4943)

		var gen37537 Obj

		if True == gen37536 {
			gen37531 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen37532 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37533 := Call(__e, gen37532, V4943)

			gen37534 := Call(__e, gen37531, symcond, gen37533)

			if True == gen37534 {
				gen37537 = True
			} else {
				gen37537 = False
			}

		} else {
			gen37537 = False
		}

		if True == gen37537 {
			gen37524 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

			gen37525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch)

			gen37526 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4add_1returns)

			gen37527 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37528 := Call(__e, gen37527, V4943)

			gen37529 := Call(__e, gen37526, gen37528)

			gen37530 := Call(__e, gen37525, gen37529, V4944)

			__e.TailApply(gen37524, gen37530, V4945)

			return

		} else {
			if True == True {
				__e.Return(V4943)
				return
			} else {
				gen37523 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen37523, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4factorise_1cond, gen37538)

	gen37585 := MakeNative(func(__e *ControlFlow) {
		V4947 := __e.Get(1)
		_ = V4947
		gen37583 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen37584 := Call(__e, gen37583, Nil, V4947)

		if True == gen37584 {
			__e.Return(Nil)
			return
		} else {
			gen37580 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen37581 := Call(__e, gen37580, V4947)

			var gen37582 Obj

			if True == gen37581 {
				gen37575 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37576 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37577 := Call(__e, gen37576, V4947)

				gen37578 := Call(__e, gen37575, gen37577)

				var gen37579 Obj

				if True == gen37578 {
					gen37568 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37569 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37570 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37571 := Call(__e, gen37570, V4947)

					gen37572 := Call(__e, gen37569, gen37571)

					gen37573 := Call(__e, gen37568, gen37572)

					var gen37574 Obj

					if True == gen37573 {
						gen37560 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen37561 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37562 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37563 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37564 := Call(__e, gen37563, V4947)

						gen37565 := Call(__e, gen37562, gen37564)

						gen37566 := Call(__e, gen37561, gen37565)

						gen37567 := Call(__e, gen37560, Nil, gen37566)

						if True == gen37567 {
							gen37574 = True
						} else {
							gen37574 = False
						}

					} else {
						gen37574 = False
					}

					if True == gen37574 {
						gen37579 = True
					} else {
						gen37579 = False
					}

				} else {
					gen37579 = False
				}

				if True == gen37579 {
					gen37582 = True
				} else {
					gen37582 = False
				}

			} else {
				gen37582 = False
			}

			if True == gen37582 {
				gen37541 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37542 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37543 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37544 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37545 := Call(__e, gen37544, V4947)

				gen37546 := Call(__e, gen37543, gen37545)

				gen37547 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37548 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37549 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37550 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37551 := Call(__e, gen37550, V4947)

				gen37552 := Call(__e, gen37549, gen37551)

				gen37553 := Call(__e, gen37548, sym_f_freturn, gen37552)

				gen37554 := Call(__e, gen37547, gen37553, Nil)

				gen37555 := Call(__e, gen37542, gen37546, gen37554)

				gen37556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4add_1returns)

				gen37557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37558 := Call(__e, gen37557, V4947)

				gen37559 := Call(__e, gen37556, gen37558)

				__e.TailApply(gen37541, gen37555, gen37559)

				return

			} else {
				if True == True {
					gen37540 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen37540, symshen_4x_4factorise_1defun_4add_1returns)

					return

				} else {
					gen37539 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen37539, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4add_1returns, gen37585)

	gen37587 := MakeNative(func(__e *ControlFlow) {
		gen37586 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		__e.TailApply(gen37586, sym_f_flabel)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4generate_1label, gen37587)

	gen37591 := MakeNative(func(__e *ControlFlow) {
		V4950 := __e.Get(1)
		_ = V4950
		V4951 := __e.Get(2)
		_ = V4951
		gen37588 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

		gen37589 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

		gen37590 := Call(__e, gen37589, V4950, V4951, Nil)

		__e.TailApply(gen37588, gen37590)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4free_1variables, gen37591)

	gen37708 := MakeNative(func(__e *ControlFlow) {
		V4963 := __e.Get(1)
		_ = V4963
		V4964 := __e.Get(2)
		_ = V4964
		V4965 := __e.Get(3)
		_ = V4965
		gen37705 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen37706 := Call(__e, gen37705, V4963)

		var gen37707 Obj

		if True == gen37706 {
			gen37700 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen37701 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37702 := Call(__e, gen37701, V4963)

			gen37703 := Call(__e, gen37700, symlet, gen37702)

			var gen37704 Obj

			if True == gen37703 {
				gen37695 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37696 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37697 := Call(__e, gen37696, V4963)

				gen37698 := Call(__e, gen37695, gen37697)

				var gen37699 Obj

				if True == gen37698 {
					gen37688 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37689 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37690 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37691 := Call(__e, gen37690, V4963)

					gen37692 := Call(__e, gen37689, gen37691)

					gen37693 := Call(__e, gen37688, gen37692)

					var gen37694 Obj

					if True == gen37693 {
						gen37679 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37680 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37681 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37682 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37683 := Call(__e, gen37682, V4963)

						gen37684 := Call(__e, gen37681, gen37683)

						gen37685 := Call(__e, gen37680, gen37684)

						gen37686 := Call(__e, gen37679, gen37685)

						var gen37687 Obj

						if True == gen37686 {
							gen37669 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37670 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37671 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37672 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37673 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37674 := Call(__e, gen37673, V4963)

							gen37675 := Call(__e, gen37672, gen37674)

							gen37676 := Call(__e, gen37671, gen37675)

							gen37677 := Call(__e, gen37670, gen37676)

							gen37678 := Call(__e, gen37669, Nil, gen37677)

							if True == gen37678 {
								gen37687 = True
							} else {
								gen37687 = False
							}

						} else {
							gen37687 = False
						}

						if True == gen37687 {
							gen37694 = True
						} else {
							gen37694 = False
						}

					} else {
						gen37694 = False
					}

					if True == gen37694 {
						gen37699 = True
					} else {
						gen37699 = False
					}

				} else {
					gen37699 = False
				}

				if True == gen37699 {
					gen37704 = True
				} else {
					gen37704 = False
				}

			} else {
				gen37704 = False
			}

			if True == gen37704 {
				gen37707 = True
			} else {
				gen37707 = False
			}

		} else {
			gen37707 = False
		}

		if True == gen37707 {
			gen37646 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

			gen37647 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37648 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37649 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37651 := Call(__e, gen37650, V4963)

			gen37652 := Call(__e, gen37649, gen37651)

			gen37653 := Call(__e, gen37648, gen37652)

			gen37654 := Call(__e, gen37647, gen37653)

			gen37655 := Call(__e, PrimNS1Value(symns2_1value), symremove)

			gen37656 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37657 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37658 := Call(__e, gen37657, V4963)

			gen37659 := Call(__e, gen37656, gen37658)

			gen37660 := Call(__e, gen37655, gen37659, V4964)

			gen37661 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

			gen37662 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37663 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37664 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37665 := Call(__e, gen37664, V4963)

			gen37666 := Call(__e, gen37663, gen37665)

			gen37667 := Call(__e, gen37662, gen37666)

			gen37668 := Call(__e, gen37661, gen37667, V4964, V4965)

			__e.TailApply(gen37646, gen37654, gen37660, gen37668)

			return

		} else {
			gen37643 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen37644 := Call(__e, gen37643, V4963)

			var gen37645 Obj

			if True == gen37644 {
				gen37638 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen37639 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37640 := Call(__e, gen37639, V4963)

				gen37641 := Call(__e, gen37638, symlambda, gen37640)

				var gen37642 Obj

				if True == gen37641 {
					gen37633 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37634 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37635 := Call(__e, gen37634, V4963)

					gen37636 := Call(__e, gen37633, gen37635)

					var gen37637 Obj

					if True == gen37636 {
						gen37626 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37627 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37628 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37629 := Call(__e, gen37628, V4963)

						gen37630 := Call(__e, gen37627, gen37629)

						gen37631 := Call(__e, gen37626, gen37630)

						var gen37632 Obj

						if True == gen37631 {
							gen37618 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37619 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37620 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37621 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37622 := Call(__e, gen37621, V4963)

							gen37623 := Call(__e, gen37620, gen37622)

							gen37624 := Call(__e, gen37619, gen37623)

							gen37625 := Call(__e, gen37618, Nil, gen37624)

							if True == gen37625 {
								gen37632 = True
							} else {
								gen37632 = False
							}

						} else {
							gen37632 = False
						}

						if True == gen37632 {
							gen37637 = True
						} else {
							gen37637 = False
						}

					} else {
						gen37637 = False
					}

					if True == gen37637 {
						gen37642 = True
					} else {
						gen37642 = False
					}

				} else {
					gen37642 = False
				}

				if True == gen37642 {
					gen37645 = True
				} else {
					gen37645 = False
				}

			} else {
				gen37645 = False
			}

			if True == gen37645 {
				gen37605 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

				gen37606 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37607 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37608 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37609 := Call(__e, gen37608, V4963)

				gen37610 := Call(__e, gen37607, gen37609)

				gen37611 := Call(__e, gen37606, gen37610)

				gen37612 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				gen37613 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37614 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37615 := Call(__e, gen37614, V4963)

				gen37616 := Call(__e, gen37613, gen37615)

				gen37617 := Call(__e, gen37612, gen37616, V4964)

				__e.TailApply(gen37605, gen37611, gen37617, V4965)

				return

			} else {
				gen37603 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37604 := Call(__e, gen37603, V4963)

				if True == gen37604 {
					gen37596 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

					gen37597 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37598 := Call(__e, gen37597, V4963)

					gen37599 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

					gen37600 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37601 := Call(__e, gen37600, V4963)

					gen37602 := Call(__e, gen37599, gen37601, V4964, V4965)

					__e.TailApply(gen37596, gen37598, V4964, gen37602)

					return

				} else {
					gen37594 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					gen37595 := Call(__e, gen37594, V4963, V4964)

					if True == gen37595 {
						gen37593 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

						__e.TailApply(gen37593, V4963, V4965)

						return

					} else {
						if True == True {
							__e.Return(V4965)
							return
						} else {
							gen37592 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen37592, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4free_1variables_1h, gen37708)

	gen37816 := MakeNative(func(__e *ControlFlow) {
		V4968 := __e.Get(1)
		_ = V4968
		V4969 := __e.Get(2)
		_ = V4969
		gen37813 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen37814 := Call(__e, gen37813, V4968)

		var gen37815 Obj

		if True == gen37814 {
			gen37808 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen37809 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37810 := Call(__e, gen37809, V4968)

			gen37811 := Call(__e, gen37808, sym_f_flet_1label, gen37810)

			var gen37812 Obj

			if True == gen37811 {
				gen37803 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37804 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37805 := Call(__e, gen37804, V4968)

				gen37806 := Call(__e, gen37803, gen37805)

				var gen37807 Obj

				if True == gen37806 {
					gen37796 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37799 := Call(__e, gen37798, V4968)

					gen37800 := Call(__e, gen37797, gen37799)

					gen37801 := Call(__e, gen37796, gen37800)

					var gen37802 Obj

					if True == gen37801 {
						gen37787 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37788 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37789 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37790 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37791 := Call(__e, gen37790, V4968)

						gen37792 := Call(__e, gen37789, gen37791)

						gen37793 := Call(__e, gen37788, gen37792)

						gen37794 := Call(__e, gen37787, gen37793)

						var gen37795 Obj

						if True == gen37794 {
							gen37777 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37779 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37781 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37782 := Call(__e, gen37781, V4968)

							gen37783 := Call(__e, gen37780, gen37782)

							gen37784 := Call(__e, gen37779, gen37783)

							gen37785 := Call(__e, gen37778, gen37784)

							gen37786 := Call(__e, gen37777, Nil, gen37785)

							if True == gen37786 {
								gen37795 = True
							} else {
								gen37795 = False
							}

						} else {
							gen37795 = False
						}

						if True == gen37795 {
							gen37802 = True
						} else {
							gen37802 = False
						}

					} else {
						gen37802 = False
					}

					if True == gen37802 {
						gen37807 = True
					} else {
						gen37807 = False
					}

				} else {
					gen37807 = False
				}

				if True == gen37807 {
					gen37812 = True
				} else {
					gen37812 = False
				}

			} else {
				gen37812 = False
			}

			if True == gen37812 {
				gen37815 = True
			} else {
				gen37815 = False
			}

		} else {
			gen37815 = False
		}

		if True == gen37815 {
			gen37768 := MakeNative(func(__e *ControlFlow) {
				FreeVars := __e.Get(1)
				_ = FreeVars
				gen37732 := MakeNative(func(__e *ControlFlow) {
					NewBody := __e.Get(1)
					_ = NewBody
					gen37711 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37712 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37713 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37714 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37715 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37716 := Call(__e, gen37715, V4968)

					gen37717 := Call(__e, gen37714, gen37716)

					gen37718 := Call(__e, gen37713, gen37717, FreeVars)

					gen37719 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37720 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37722 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37723 := Call(__e, gen37722, V4968)

					gen37724 := Call(__e, gen37721, gen37723)

					gen37725 := Call(__e, gen37720, gen37724)

					gen37726 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

					gen37728 := Call(__e, gen37727, NewBody, V4969)

					gen37729 := Call(__e, gen37726, gen37728, Nil)

					gen37730 := Call(__e, gen37719, gen37725, gen37729)

					gen37731 := Call(__e, gen37712, gen37718, gen37730)

					__e.TailApply(gen37711, sym_f_flet_1label, gen37731)

					return

				}, 1)

				gen37765 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen37766 := Call(__e, gen37765, Nil, FreeVars)

				var gen37767 Obj

				if True == gen37766 {
					gen37758 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37760 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37761 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37762 := Call(__e, gen37761, V4968)

					gen37763 := Call(__e, gen37760, gen37762)

					gen37764 := Call(__e, gen37759, gen37763)

					gen37767 = Call(__e, gen37758, gen37764)

				} else {
					gen37733 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					gen37734 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37736 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37737 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37738 := Call(__e, gen37737, V4968)

					gen37739 := Call(__e, gen37736, gen37738)

					gen37740 := Call(__e, gen37735, gen37739, FreeVars)

					gen37741 := Call(__e, gen37734, sym_f_fgoto_1label, gen37740)

					gen37742 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37743 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37744 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37745 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37746 := Call(__e, gen37745, V4968)

					gen37747 := Call(__e, gen37744, gen37746)

					gen37748 := Call(__e, gen37743, gen37747, Nil)

					gen37749 := Call(__e, gen37742, sym_f_fgoto_1label, gen37748)

					gen37750 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37751 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37752 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37753 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37754 := Call(__e, gen37753, V4968)

					gen37755 := Call(__e, gen37752, gen37754)

					gen37756 := Call(__e, gen37751, gen37755)

					gen37757 := Call(__e, gen37750, gen37756)

					gen37767 = Call(__e, gen37733, gen37741, gen37749, gen37757)

				}

				__e.TailApply(gen37732, gen37767)

				return

			}, 1)

			gen37769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables)

			gen37770 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37771 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37772 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37773 := Call(__e, gen37772, V4968)

			gen37774 := Call(__e, gen37771, gen37773)

			gen37775 := Call(__e, gen37770, gen37774)

			gen37776 := Call(__e, gen37769, gen37775, V4969)

			__e.TailApply(gen37768, gen37776)

			return

		} else {
			if True == True {
				gen37710 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen37710, symshen_4x_4factorise_1defun_4attach_1free_1variables)

				return

			} else {
				gen37709 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen37709, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4attach_1free_1variables, gen37816)

	gen38109 := MakeNative(func(__e *ControlFlow) {
		V4976 := __e.Get(1)
		_ = V4976
		V4977 := __e.Get(2)
		_ = V4977
		gen38106 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen38107 := Call(__e, gen38106, V4976)

		var gen38108 Obj

		if True == gen38107 {
			gen38101 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen38102 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38103 := Call(__e, gen38102, V4976)

			gen38104 := Call(__e, gen38101, sym_f_flet_1label, gen38103)

			var gen38105 Obj

			if True == gen38104 {
				gen38096 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38097 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38098 := Call(__e, gen38097, V4976)

				gen38099 := Call(__e, gen38096, gen38098)

				var gen38100 Obj

				if True == gen38099 {
					gen38089 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38090 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38091 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38092 := Call(__e, gen38091, V4976)

					gen38093 := Call(__e, gen38090, gen38092)

					gen38094 := Call(__e, gen38089, gen38093)

					var gen38095 Obj

					if True == gen38094 {
						gen38080 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38081 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38082 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38083 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38084 := Call(__e, gen38083, V4976)

						gen38085 := Call(__e, gen38082, gen38084)

						gen38086 := Call(__e, gen38081, gen38085)

						gen38087 := Call(__e, gen38080, gen38086)

						var gen38088 Obj

						if True == gen38087 {
							gen38069 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38072 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38073 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38074 := Call(__e, gen38073, V4976)

							gen38075 := Call(__e, gen38072, gen38074)

							gen38076 := Call(__e, gen38071, gen38075)

							gen38077 := Call(__e, gen38070, gen38076)

							gen38078 := Call(__e, gen38069, Nil, gen38077)

							var gen38079 Obj

							if True == gen38078 {
								gen38049 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

								gen38050 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

								gen38051 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen38052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen38053 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38054 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38055 := Call(__e, gen38054, V4976)

								gen38056 := Call(__e, gen38053, gen38055)

								gen38057 := Call(__e, gen38052, gen38056, Nil)

								gen38058 := Call(__e, gen38051, sym_f_fgoto_1label, gen38057)

								gen38059 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38060 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38061 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38062 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38063 := Call(__e, gen38062, V4976)

								gen38064 := Call(__e, gen38061, gen38063)

								gen38065 := Call(__e, gen38060, gen38064)

								gen38066 := Call(__e, gen38059, gen38065)

								gen38067 := Call(__e, gen38050, gen38058, gen38066)

								gen38068 := Call(__e, gen38049, gen38067, MakeNumber(1))

								if True == gen38068 {
									gen38079 = True
								} else {
									gen38079 = False
								}

							} else {
								gen38079 = False
							}

							if True == gen38079 {
								gen38088 = True
							} else {
								gen38088 = False
							}

						} else {
							gen38088 = False
						}

						if True == gen38088 {
							gen38095 = True
						} else {
							gen38095 = False
						}

					} else {
						gen38095 = False
					}

					if True == gen38095 {
						gen38100 = True
					} else {
						gen38100 = False
					}

				} else {
					gen38100 = False
				}

				if True == gen38100 {
					gen38105 = True
				} else {
					gen38105 = False
				}

			} else {
				gen38105 = False
			}

			if True == gen38105 {
				gen38108 = True
			} else {
				gen38108 = False
			}

		} else {
			gen38108 = False
		}

		if True == gen38108 {
			gen38024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4attach_1free_1variables)

			gen38025 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38027 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38028 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38029 := Call(__e, gen38028, V4976)

			gen38030 := Call(__e, gen38027, gen38029)

			gen38031 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38032 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

			gen38033 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38035 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38036 := Call(__e, gen38035, V4976)

			gen38037 := Call(__e, gen38034, gen38036)

			gen38038 := Call(__e, gen38033, gen38037)

			gen38039 := Call(__e, gen38032, gen38038, V4977)

			gen38040 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38041 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38043 := Call(__e, gen38042, V4976)

			gen38044 := Call(__e, gen38041, gen38043)

			gen38045 := Call(__e, gen38040, gen38044)

			gen38046 := Call(__e, gen38031, gen38039, gen38045)

			gen38047 := Call(__e, gen38026, gen38030, gen38046)

			gen38048 := Call(__e, gen38025, sym_f_flet_1label, gen38047)

			__e.TailApply(gen38024, gen38048, V4977)

			return

		} else {
			gen38021 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38022 := Call(__e, gen38021, V4976)

			var gen38023 Obj

			if True == gen38022 {
				gen38016 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen38017 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38018 := Call(__e, gen38017, V4976)

				gen38019 := Call(__e, gen38016, sym_f_flet_1label, gen38018)

				var gen38020 Obj

				if True == gen38019 {
					gen38011 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38012 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38013 := Call(__e, gen38012, V4976)

					gen38014 := Call(__e, gen38011, gen38013)

					var gen38015 Obj

					if True == gen38014 {
						gen38004 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38005 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38006 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38007 := Call(__e, gen38006, V4976)

						gen38008 := Call(__e, gen38005, gen38007)

						gen38009 := Call(__e, gen38004, gen38008)

						var gen38010 Obj

						if True == gen38009 {
							gen37995 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37996 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37997 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37998 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37999 := Call(__e, gen37998, V4976)

							gen38000 := Call(__e, gen37997, gen37999)

							gen38001 := Call(__e, gen37996, gen38000)

							gen38002 := Call(__e, gen37995, gen38001)

							var gen38003 Obj

							if True == gen38002 {
								gen37985 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen37986 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37988 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37989 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37990 := Call(__e, gen37989, V4976)

								gen37991 := Call(__e, gen37988, gen37990)

								gen37992 := Call(__e, gen37987, gen37991)

								gen37993 := Call(__e, gen37986, gen37992)

								gen37994 := Call(__e, gen37985, Nil, gen37993)

								if True == gen37994 {
									gen38003 = True
								} else {
									gen38003 = False
								}

							} else {
								gen38003 = False
							}

							if True == gen38003 {
								gen38010 = True
							} else {
								gen38010 = False
							}

						} else {
							gen38010 = False
						}

						if True == gen38010 {
							gen38015 = True
						} else {
							gen38015 = False
						}

					} else {
						gen38015 = False
					}

					if True == gen38015 {
						gen38020 = True
					} else {
						gen38020 = False
					}

				} else {
					gen38020 = False
				}

				if True == gen38020 {
					gen38023 = True
				} else {
					gen38023 = False
				}

			} else {
				gen38023 = False
			}

			if True == gen38023 {
				gen37958 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

				gen37959 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

				gen37960 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37961 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37962 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37963 := Call(__e, gen37962, V4976)

				gen37964 := Call(__e, gen37961, gen37963)

				gen37965 := Call(__e, gen37960, gen37964)

				gen37966 := Call(__e, gen37959, gen37965, V4977)

				gen37967 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37968 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37969 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37970 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37971 := Call(__e, gen37970, V4976)

				gen37972 := Call(__e, gen37969, gen37971)

				gen37973 := Call(__e, gen37968, gen37972, Nil)

				gen37974 := Call(__e, gen37967, sym_f_fgoto_1label, gen37973)

				gen37975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

				gen37976 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37977 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37978 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37979 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37980 := Call(__e, gen37979, V4976)

				gen37981 := Call(__e, gen37978, gen37980)

				gen37982 := Call(__e, gen37977, gen37981)

				gen37983 := Call(__e, gen37976, gen37982)

				gen37984 := Call(__e, gen37975, gen37983, V4977)

				__e.TailApply(gen37958, gen37966, gen37974, gen37984)

				return

			} else {
				gen37955 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37956 := Call(__e, gen37955, V4976)

				var gen37957 Obj

				if True == gen37956 {
					gen37950 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen37951 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37952 := Call(__e, gen37951, V4976)

					gen37953 := Call(__e, gen37950, symif, gen37952)

					var gen37954 Obj

					if True == gen37953 {
						gen37945 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37946 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37947 := Call(__e, gen37946, V4976)

						gen37948 := Call(__e, gen37945, gen37947)

						var gen37949 Obj

						if True == gen37948 {
							gen37938 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37941 := Call(__e, gen37940, V4976)

							gen37942 := Call(__e, gen37939, gen37941)

							gen37943 := Call(__e, gen37938, gen37942)

							var gen37944 Obj

							if True == gen37943 {
								gen37929 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen37930 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37931 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37932 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37933 := Call(__e, gen37932, V4976)

								gen37934 := Call(__e, gen37931, gen37933)

								gen37935 := Call(__e, gen37930, gen37934)

								gen37936 := Call(__e, gen37929, gen37935)

								var gen37937 Obj

								if True == gen37936 {
									gen37919 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen37920 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37921 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37922 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37923 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37924 := Call(__e, gen37923, V4976)

									gen37925 := Call(__e, gen37922, gen37924)

									gen37926 := Call(__e, gen37921, gen37925)

									gen37927 := Call(__e, gen37920, gen37926)

									gen37928 := Call(__e, gen37919, Nil, gen37927)

									if True == gen37928 {
										gen37937 = True
									} else {
										gen37937 = False
									}

								} else {
									gen37937 = False
								}

								if True == gen37937 {
									gen37944 = True
								} else {
									gen37944 = False
								}

							} else {
								gen37944 = False
							}

							if True == gen37944 {
								gen37949 = True
							} else {
								gen37949 = False
							}

						} else {
							gen37949 = False
						}

						if True == gen37949 {
							gen37954 = True
						} else {
							gen37954 = False
						}

					} else {
						gen37954 = False
					}

					if True == gen37954 {
						gen37957 = True
					} else {
						gen37957 = False
					}

				} else {
					gen37957 = False
				}

				if True == gen37957 {
					gen37890 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37891 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37892 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37893 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37894 := Call(__e, gen37893, V4976)

					gen37895 := Call(__e, gen37892, gen37894)

					gen37896 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37897 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

					gen37898 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37899 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37900 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37901 := Call(__e, gen37900, V4976)

					gen37902 := Call(__e, gen37899, gen37901)

					gen37903 := Call(__e, gen37898, gen37902)

					gen37904 := Call(__e, gen37897, gen37903, V4977)

					gen37905 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

					gen37907 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37908 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37909 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37910 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37911 := Call(__e, gen37910, V4976)

					gen37912 := Call(__e, gen37909, gen37911)

					gen37913 := Call(__e, gen37908, gen37912)

					gen37914 := Call(__e, gen37907, gen37913)

					gen37915 := Call(__e, gen37906, gen37914, V4977)

					gen37916 := Call(__e, gen37905, gen37915, Nil)

					gen37917 := Call(__e, gen37896, gen37904, gen37916)

					gen37918 := Call(__e, gen37891, gen37895, gen37917)

					__e.TailApply(gen37890, symif, gen37918)

					return

				} else {
					gen37887 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37888 := Call(__e, gen37887, V4976)

					var gen37889 Obj

					if True == gen37888 {
						gen37882 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen37883 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37884 := Call(__e, gen37883, V4976)

						gen37885 := Call(__e, gen37882, symlet, gen37884)

						var gen37886 Obj

						if True == gen37885 {
							gen37877 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37878 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37879 := Call(__e, gen37878, V4976)

							gen37880 := Call(__e, gen37877, gen37879)

							var gen37881 Obj

							if True == gen37880 {
								gen37870 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen37871 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37873 := Call(__e, gen37872, V4976)

								gen37874 := Call(__e, gen37871, gen37873)

								gen37875 := Call(__e, gen37870, gen37874)

								var gen37876 Obj

								if True == gen37875 {
									gen37861 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen37862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37863 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37864 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37865 := Call(__e, gen37864, V4976)

									gen37866 := Call(__e, gen37863, gen37865)

									gen37867 := Call(__e, gen37862, gen37866)

									gen37868 := Call(__e, gen37861, gen37867)

									var gen37869 Obj

									if True == gen37868 {
										gen37851 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen37852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen37853 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen37854 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen37855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen37856 := Call(__e, gen37855, V4976)

										gen37857 := Call(__e, gen37854, gen37856)

										gen37858 := Call(__e, gen37853, gen37857)

										gen37859 := Call(__e, gen37852, gen37858)

										gen37860 := Call(__e, gen37851, Nil, gen37859)

										if True == gen37860 {
											gen37869 = True
										} else {
											gen37869 = False
										}

									} else {
										gen37869 = False
									}

									if True == gen37869 {
										gen37876 = True
									} else {
										gen37876 = False
									}

								} else {
									gen37876 = False
								}

								if True == gen37876 {
									gen37881 = True
								} else {
									gen37881 = False
								}

							} else {
								gen37881 = False
							}

							if True == gen37881 {
								gen37886 = True
							} else {
								gen37886 = False
							}

						} else {
							gen37886 = False
						}

						if True == gen37886 {
							gen37889 = True
						} else {
							gen37889 = False
						}

					} else {
						gen37889 = False
					}

					if True == gen37889 {
						gen37818 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37819 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37820 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37822 := Call(__e, gen37821, V4976)

						gen37823 := Call(__e, gen37820, gen37822)

						gen37824 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37825 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37826 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37827 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37828 := Call(__e, gen37827, V4976)

						gen37829 := Call(__e, gen37826, gen37828)

						gen37830 := Call(__e, gen37825, gen37829)

						gen37831 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

						gen37833 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37836 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37837 := Call(__e, gen37836, V4976)

						gen37838 := Call(__e, gen37835, gen37837)

						gen37839 := Call(__e, gen37834, gen37838)

						gen37840 := Call(__e, gen37833, gen37839)

						gen37841 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37842 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37843 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37844 := Call(__e, gen37843, V4976)

						gen37845 := Call(__e, gen37842, gen37844)

						gen37846 := Call(__e, gen37841, gen37845, V4977)

						gen37847 := Call(__e, gen37832, gen37840, gen37846)

						gen37848 := Call(__e, gen37831, gen37847, Nil)

						gen37849 := Call(__e, gen37824, gen37830, gen37848)

						gen37850 := Call(__e, gen37819, gen37823, gen37849)

						__e.TailApply(gen37818, symlet, gen37850)

						return

					} else {
						if True == True {
							__e.Return(V4976)
							return
						} else {
							gen37817 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen37817, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4inline_1mono_1labels, gen38109)

	gen38294 := MakeNative(func(__e *ControlFlow) {
		V4984 := __e.Get(1)
		_ = V4984
		V4985 := __e.Get(2)
		_ = V4985
		gen38292 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen38293 := Call(__e, gen38292, Nil, V4984)

		if True == gen38293 {
			__e.Return(V4985)
			return
		} else {
			gen38289 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38290 := Call(__e, gen38289, V4984)

			var gen38291 Obj

			if True == gen38290 {
				gen38284 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38285 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38286 := Call(__e, gen38285, V4984)

				gen38287 := Call(__e, gen38284, gen38286)

				var gen38288 Obj

				if True == gen38287 {
					gen38277 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38278 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38279 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38280 := Call(__e, gen38279, V4984)

					gen38281 := Call(__e, gen38278, gen38280)

					gen38282 := Call(__e, gen38277, True, gen38281)

					var gen38283 Obj

					if True == gen38282 {
						gen38270 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38271 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38272 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38273 := Call(__e, gen38272, V4984)

						gen38274 := Call(__e, gen38271, gen38273)

						gen38275 := Call(__e, gen38270, gen38274)

						var gen38276 Obj

						if True == gen38275 {
							gen38262 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38263 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38264 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38265 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38266 := Call(__e, gen38265, V4984)

							gen38267 := Call(__e, gen38264, gen38266)

							gen38268 := Call(__e, gen38263, gen38267)

							gen38269 := Call(__e, gen38262, Nil, gen38268)

							if True == gen38269 {
								gen38276 = True
							} else {
								gen38276 = False
							}

						} else {
							gen38276 = False
						}

						if True == gen38276 {
							gen38283 = True
						} else {
							gen38283 = False
						}

					} else {
						gen38283 = False
					}

					if True == gen38283 {
						gen38288 = True
					} else {
						gen38288 = False
					}

				} else {
					gen38288 = False
				}

				if True == gen38288 {
					gen38291 = True
				} else {
					gen38291 = False
				}

			} else {
				gen38291 = False
			}

			if True == gen38291 {
				gen38257 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38259 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38260 := Call(__e, gen38259, V4984)

				gen38261 := Call(__e, gen38258, gen38260)

				__e.TailApply(gen38257, gen38261)

				return

			} else {
				gen38254 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38255 := Call(__e, gen38254, V4984)

				var gen38256 Obj

				if True == gen38255 {
					gen38249 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38250 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38251 := Call(__e, gen38250, V4984)

					gen38252 := Call(__e, gen38249, gen38251)

					var gen38253 Obj

					if True == gen38252 {
						gen38242 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38243 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38244 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38245 := Call(__e, gen38244, V4984)

						gen38246 := Call(__e, gen38243, gen38245)

						gen38247 := Call(__e, gen38242, gen38246)

						var gen38248 Obj

						if True == gen38247 {
							gen38233 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38234 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38235 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38236 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38237 := Call(__e, gen38236, V4984)

							gen38238 := Call(__e, gen38235, gen38237)

							gen38239 := Call(__e, gen38234, gen38238)

							gen38240 := Call(__e, gen38233, symand, gen38239)

							var gen38241 Obj

							if True == gen38240 {
								gen38224 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen38225 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38226 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38227 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38228 := Call(__e, gen38227, V4984)

								gen38229 := Call(__e, gen38226, gen38228)

								gen38230 := Call(__e, gen38225, gen38229)

								gen38231 := Call(__e, gen38224, gen38230)

								var gen38232 Obj

								if True == gen38231 {
									gen38213 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen38214 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38215 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38216 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen38217 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen38218 := Call(__e, gen38217, V4984)

									gen38219 := Call(__e, gen38216, gen38218)

									gen38220 := Call(__e, gen38215, gen38219)

									gen38221 := Call(__e, gen38214, gen38220)

									gen38222 := Call(__e, gen38213, gen38221)

									var gen38223 Obj

									if True == gen38222 {
										gen38200 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen38201 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38202 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38203 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38204 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen38205 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen38206 := Call(__e, gen38205, V4984)

										gen38207 := Call(__e, gen38204, gen38206)

										gen38208 := Call(__e, gen38203, gen38207)

										gen38209 := Call(__e, gen38202, gen38208)

										gen38210 := Call(__e, gen38201, gen38209)

										gen38211 := Call(__e, gen38200, Nil, gen38210)

										var gen38212 Obj

										if True == gen38211 {
											gen38193 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen38194 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38195 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38196 := Call(__e, gen38195, V4984)

											gen38197 := Call(__e, gen38194, gen38196)

											gen38198 := Call(__e, gen38193, gen38197)

											var gen38199 Obj

											if True == gen38198 {
												gen38185 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen38186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen38187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen38188 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen38189 := Call(__e, gen38188, V4984)

												gen38190 := Call(__e, gen38187, gen38189)

												gen38191 := Call(__e, gen38186, gen38190)

												gen38192 := Call(__e, gen38185, Nil, gen38191)

												if True == gen38192 {
													gen38199 = True
												} else {
													gen38199 = False
												}

											} else {
												gen38199 = False
											}

											if True == gen38199 {
												gen38212 = True
											} else {
												gen38212 = False
											}

										} else {
											gen38212 = False
										}

										if True == gen38212 {
											gen38223 = True
										} else {
											gen38223 = False
										}

									} else {
										gen38223 = False
									}

									if True == gen38223 {
										gen38232 = True
									} else {
										gen38232 = False
									}

								} else {
									gen38232 = False
								}

								if True == gen38232 {
									gen38241 = True
								} else {
									gen38241 = False
								}

							} else {
								gen38241 = False
							}

							if True == gen38241 {
								gen38248 = True
							} else {
								gen38248 = False
							}

						} else {
							gen38248 = False
						}

						if True == gen38248 {
							gen38253 = True
						} else {
							gen38253 = False
						}

					} else {
						gen38253 = False
					}

					if True == gen38253 {
						gen38256 = True
					} else {
						gen38256 = False
					}

				} else {
					gen38256 = False
				}

				if True == gen38256 {
					gen38174 := MakeNative(func(__e *ControlFlow) {
						TrueBranch := __e.Get(1)
						_ = TrueBranch
						gen38163 := MakeNative(func(__e *ControlFlow) {
							FalseBranch := __e.Get(1)
							_ = FalseBranch
							gen38154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch_1h)

							gen38155 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38156 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38157 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38158 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38159 := Call(__e, gen38158, V4984)

							gen38160 := Call(__e, gen38157, gen38159)

							gen38161 := Call(__e, gen38156, gen38160)

							gen38162 := Call(__e, gen38155, gen38161)

							__e.TailApply(gen38154, gen38162, TrueBranch, FalseBranch, V4985)

							return

						}, 1)

						gen38164 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

						gen38165 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38166 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38167 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38168 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38169 := Call(__e, gen38168, V4984)

						gen38170 := Call(__e, gen38167, gen38169)

						gen38171 := Call(__e, gen38166, gen38170)

						gen38172 := Call(__e, gen38165, gen38171)

						gen38173 := Call(__e, gen38164, gen38172, V4984)

						__e.TailApply(gen38163, gen38173)

						return

					}, 1)

					gen38175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4true_1branch)

					gen38176 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38177 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38178 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38179 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38180 := Call(__e, gen38179, V4984)

					gen38181 := Call(__e, gen38178, gen38180)

					gen38182 := Call(__e, gen38177, gen38181)

					gen38183 := Call(__e, gen38176, gen38182)

					gen38184 := Call(__e, gen38175, gen38183, V4984)

					__e.TailApply(gen38174, gen38184)

					return

				} else {
					gen38151 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38152 := Call(__e, gen38151, V4984)

					var gen38153 Obj

					if True == gen38152 {
						gen38146 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38147 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38148 := Call(__e, gen38147, V4984)

						gen38149 := Call(__e, gen38146, gen38148)

						var gen38150 Obj

						if True == gen38149 {
							gen38139 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen38140 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38141 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38142 := Call(__e, gen38141, V4984)

							gen38143 := Call(__e, gen38140, gen38142)

							gen38144 := Call(__e, gen38139, gen38143)

							var gen38145 Obj

							if True == gen38144 {
								gen38131 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen38132 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38133 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38134 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38135 := Call(__e, gen38134, V4984)

								gen38136 := Call(__e, gen38133, gen38135)

								gen38137 := Call(__e, gen38132, gen38136)

								gen38138 := Call(__e, gen38131, Nil, gen38137)

								if True == gen38138 {
									gen38145 = True
								} else {
									gen38145 = False
								}

							} else {
								gen38145 = False
							}

							if True == gen38145 {
								gen38150 = True
							} else {
								gen38150 = False
							}

						} else {
							gen38150 = False
						}

						if True == gen38150 {
							gen38153 = True
						} else {
							gen38153 = False
						}

					} else {
						gen38153 = False
					}

					if True == gen38153 {
						gen38124 := MakeNative(func(__e *ControlFlow) {
							TrueBranch := __e.Get(1)
							_ = TrueBranch
							gen38117 := MakeNative(func(__e *ControlFlow) {
								FalseBranch := __e.Get(1)
								_ = FalseBranch
								gen38112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch_1h)

								gen38113 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38114 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38115 := Call(__e, gen38114, V4984)

								gen38116 := Call(__e, gen38113, gen38115)

								__e.TailApply(gen38112, gen38116, TrueBranch, FalseBranch, V4985)

								return

							}, 1)

							gen38118 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

							gen38119 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38120 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38121 := Call(__e, gen38120, V4984)

							gen38122 := Call(__e, gen38119, gen38121)

							gen38123 := Call(__e, gen38118, gen38122, V4984)

							__e.TailApply(gen38117, gen38123)

							return

						}, 1)

						gen38125 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4true_1branch)

						gen38126 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38127 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38128 := Call(__e, gen38127, V4984)

						gen38129 := Call(__e, gen38126, gen38128)

						gen38130 := Call(__e, gen38125, gen38129, V4984)

						__e.TailApply(gen38124, gen38130)

						return

					} else {
						if True == True {
							gen38111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen38111, symshen_4x_4factorise_1defun_4rebranch)

							return

						} else {
							gen38110 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen38110, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4rebranch, gen38294)

	gen38313 := MakeNative(func(__e *ControlFlow) {
		V4990 := __e.Get(1)
		_ = V4990
		V4991 := __e.Get(2)
		_ = V4991
		V4992 := __e.Get(3)
		_ = V4992
		V4993 := __e.Get(4)
		_ = V4993
		gen38310 := MakeNative(func(__e *ControlFlow) {
			NewElse := __e.Get(1)
			_ = NewElse
			gen38295 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4with_1labelled_1else)

			gen38309 := MakeNative(func(__e *ControlFlow) {
				GotoElse := __e.Get(1)
				_ = GotoElse
				gen38296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs)

				gen38297 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38298 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38299 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38300 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4optimize_1selectors)

				gen38301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch)

				gen38302 := Call(__e, gen38301, V4991, GotoElse)

				gen38303 := Call(__e, gen38300, V4990, gen38302)

				gen38304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38305 := Call(__e, gen38304, GotoElse, Nil)

				gen38306 := Call(__e, gen38299, gen38303, gen38305)

				gen38307 := Call(__e, gen38298, V4990, gen38306)

				gen38308 := Call(__e, gen38297, symif, gen38307)

				__e.TailApply(gen38296, gen38308)

				return

			}, 1)

			__e.TailApply(gen38295, NewElse, gen38309)

			return

		}, 1)

		gen38311 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch)

		gen38312 := Call(__e, gen38311, V4992, V4993)

		__e.TailApply(gen38310, gen38312)

		return

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4rebranch_1h, gen38313)

	gen38464 := MakeNative(func(__e *ControlFlow) {
		V5006 := __e.Get(1)
		_ = V5006
		V5007 := __e.Get(2)
		_ = V5007
		gen38461 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen38462 := Call(__e, gen38461, V5007)

		var gen38463 Obj

		if True == gen38462 {
			gen38456 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38457 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38458 := Call(__e, gen38457, V5007)

			gen38459 := Call(__e, gen38456, gen38458)

			var gen38460 Obj

			if True == gen38459 {
				gen38449 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38450 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38451 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38452 := Call(__e, gen38451, V5007)

				gen38453 := Call(__e, gen38450, gen38452)

				gen38454 := Call(__e, gen38449, gen38453)

				var gen38455 Obj

				if True == gen38454 {
					gen38440 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38441 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38442 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38443 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38444 := Call(__e, gen38443, V5007)

					gen38445 := Call(__e, gen38442, gen38444)

					gen38446 := Call(__e, gen38441, gen38445)

					gen38447 := Call(__e, gen38440, symand, gen38446)

					var gen38448 Obj

					if True == gen38447 {
						gen38431 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38432 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38433 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38434 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38435 := Call(__e, gen38434, V5007)

						gen38436 := Call(__e, gen38433, gen38435)

						gen38437 := Call(__e, gen38432, gen38436)

						gen38438 := Call(__e, gen38431, gen38437)

						var gen38439 Obj

						if True == gen38438 {
							gen38420 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen38421 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38422 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38423 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38424 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38425 := Call(__e, gen38424, V5007)

							gen38426 := Call(__e, gen38423, gen38425)

							gen38427 := Call(__e, gen38422, gen38426)

							gen38428 := Call(__e, gen38421, gen38427)

							gen38429 := Call(__e, gen38420, gen38428)

							var gen38430 Obj

							if True == gen38429 {
								gen38407 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen38408 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38409 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38410 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38411 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38412 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38413 := Call(__e, gen38412, V5007)

								gen38414 := Call(__e, gen38411, gen38413)

								gen38415 := Call(__e, gen38410, gen38414)

								gen38416 := Call(__e, gen38409, gen38415)

								gen38417 := Call(__e, gen38408, gen38416)

								gen38418 := Call(__e, gen38407, Nil, gen38417)

								var gen38419 Obj

								if True == gen38418 {
									gen38400 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen38401 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38402 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen38403 := Call(__e, gen38402, V5007)

									gen38404 := Call(__e, gen38401, gen38403)

									gen38405 := Call(__e, gen38400, gen38404)

									var gen38406 Obj

									if True == gen38405 {
										gen38391 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen38392 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38393 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38394 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen38395 := Call(__e, gen38394, V5007)

										gen38396 := Call(__e, gen38393, gen38395)

										gen38397 := Call(__e, gen38392, gen38396)

										gen38398 := Call(__e, gen38391, Nil, gen38397)

										var gen38399 Obj

										if True == gen38398 {
											gen38381 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen38382 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38383 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38384 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38385 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38386 := Call(__e, gen38385, V5007)

											gen38387 := Call(__e, gen38384, gen38386)

											gen38388 := Call(__e, gen38383, gen38387)

											gen38389 := Call(__e, gen38382, gen38388)

											gen38390 := Call(__e, gen38381, gen38389, V5006)

											if True == gen38390 {
												gen38399 = True
											} else {
												gen38399 = False
											}

										} else {
											gen38399 = False
										}

										if True == gen38399 {
											gen38406 = True
										} else {
											gen38406 = False
										}

									} else {
										gen38406 = False
									}

									if True == gen38406 {
										gen38419 = True
									} else {
										gen38419 = False
									}

								} else {
									gen38419 = False
								}

								if True == gen38419 {
									gen38430 = True
								} else {
									gen38430 = False
								}

							} else {
								gen38430 = False
							}

							if True == gen38430 {
								gen38439 = True
							} else {
								gen38439 = False
							}

						} else {
							gen38439 = False
						}

						if True == gen38439 {
							gen38448 = True
						} else {
							gen38448 = False
						}

					} else {
						gen38448 = False
					}

					if True == gen38448 {
						gen38455 = True
					} else {
						gen38455 = False
					}

				} else {
					gen38455 = False
				}

				if True == gen38455 {
					gen38460 = True
				} else {
					gen38460 = False
				}

			} else {
				gen38460 = False
			}

			if True == gen38460 {
				gen38463 = True
			} else {
				gen38463 = False
			}

		} else {
			gen38463 = False
		}

		if True == gen38463 {
			gen38352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38354 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38356 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38357 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38358 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38359 := Call(__e, gen38358, V5007)

			gen38360 := Call(__e, gen38357, gen38359)

			gen38361 := Call(__e, gen38356, gen38360)

			gen38362 := Call(__e, gen38355, gen38361)

			gen38363 := Call(__e, gen38354, gen38362)

			gen38364 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38366 := Call(__e, gen38365, V5007)

			gen38367 := Call(__e, gen38364, gen38366)

			gen38368 := Call(__e, gen38353, gen38363, gen38367)

			gen38369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4true_1branch)

			gen38370 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38371 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38372 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38373 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38374 := Call(__e, gen38373, V5007)

			gen38375 := Call(__e, gen38372, gen38374)

			gen38376 := Call(__e, gen38371, gen38375)

			gen38377 := Call(__e, gen38370, gen38376)

			gen38378 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38379 := Call(__e, gen38378, V5007)

			gen38380 := Call(__e, gen38369, gen38377, gen38379)

			__e.TailApply(gen38352, gen38368, gen38380)

			return

		} else {
			gen38349 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38350 := Call(__e, gen38349, V5007)

			var gen38351 Obj

			if True == gen38350 {
				gen38344 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38345 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38346 := Call(__e, gen38345, V5007)

				gen38347 := Call(__e, gen38344, gen38346)

				var gen38348 Obj

				if True == gen38347 {
					gen38337 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38339 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38340 := Call(__e, gen38339, V5007)

					gen38341 := Call(__e, gen38338, gen38340)

					gen38342 := Call(__e, gen38337, gen38341)

					var gen38343 Obj

					if True == gen38342 {
						gen38328 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen38329 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38330 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38331 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38332 := Call(__e, gen38331, V5007)

						gen38333 := Call(__e, gen38330, gen38332)

						gen38334 := Call(__e, gen38329, gen38333)

						gen38335 := Call(__e, gen38328, Nil, gen38334)

						var gen38336 Obj

						if True == gen38335 {
							gen38322 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38323 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38324 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38325 := Call(__e, gen38324, V5007)

							gen38326 := Call(__e, gen38323, gen38325)

							gen38327 := Call(__e, gen38322, gen38326, V5006)

							if True == gen38327 {
								gen38336 = True
							} else {
								gen38336 = False
							}

						} else {
							gen38336 = False
						}

						if True == gen38336 {
							gen38343 = True
						} else {
							gen38343 = False
						}

					} else {
						gen38343 = False
					}

					if True == gen38343 {
						gen38348 = True
					} else {
						gen38348 = False
					}

				} else {
					gen38348 = False
				}

				if True == gen38348 {
					gen38351 = True
				} else {
					gen38351 = False
				}

			} else {
				gen38351 = False
			}

			if True == gen38351 {
				gen38315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38317 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38318 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38319 := Call(__e, gen38318, V5007)

				gen38320 := Call(__e, gen38317, gen38319)

				gen38321 := Call(__e, gen38316, True, gen38320)

				__e.TailApply(gen38315, gen38321, Nil)

				return

			} else {
				if True == True {
					__e.Return(Nil)
					return
				} else {
					gen38314 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen38314, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4true_1branch, gen38464)

	gen38597 := MakeNative(func(__e *ControlFlow) {
		V5016 := __e.Get(1)
		_ = V5016
		V5017 := __e.Get(2)
		_ = V5017
		gen38594 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen38595 := Call(__e, gen38594, V5017)

		var gen38596 Obj

		if True == gen38595 {
			gen38589 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38590 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38591 := Call(__e, gen38590, V5017)

			gen38592 := Call(__e, gen38589, gen38591)

			var gen38593 Obj

			if True == gen38592 {
				gen38582 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38583 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38584 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38585 := Call(__e, gen38584, V5017)

				gen38586 := Call(__e, gen38583, gen38585)

				gen38587 := Call(__e, gen38582, gen38586)

				var gen38588 Obj

				if True == gen38587 {
					gen38573 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38575 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38576 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38577 := Call(__e, gen38576, V5017)

					gen38578 := Call(__e, gen38575, gen38577)

					gen38579 := Call(__e, gen38574, gen38578)

					gen38580 := Call(__e, gen38573, symand, gen38579)

					var gen38581 Obj

					if True == gen38580 {
						gen38564 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38565 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38566 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38567 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38568 := Call(__e, gen38567, V5017)

						gen38569 := Call(__e, gen38566, gen38568)

						gen38570 := Call(__e, gen38565, gen38569)

						gen38571 := Call(__e, gen38564, gen38570)

						var gen38572 Obj

						if True == gen38571 {
							gen38553 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen38554 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38555 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38556 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38557 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38558 := Call(__e, gen38557, V5017)

							gen38559 := Call(__e, gen38556, gen38558)

							gen38560 := Call(__e, gen38555, gen38559)

							gen38561 := Call(__e, gen38554, gen38560)

							gen38562 := Call(__e, gen38553, gen38561)

							var gen38563 Obj

							if True == gen38562 {
								gen38540 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen38541 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38542 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38543 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38544 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38545 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38546 := Call(__e, gen38545, V5017)

								gen38547 := Call(__e, gen38544, gen38546)

								gen38548 := Call(__e, gen38543, gen38547)

								gen38549 := Call(__e, gen38542, gen38548)

								gen38550 := Call(__e, gen38541, gen38549)

								gen38551 := Call(__e, gen38540, Nil, gen38550)

								var gen38552 Obj

								if True == gen38551 {
									gen38533 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen38534 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38535 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen38536 := Call(__e, gen38535, V5017)

									gen38537 := Call(__e, gen38534, gen38536)

									gen38538 := Call(__e, gen38533, gen38537)

									var gen38539 Obj

									if True == gen38538 {
										gen38524 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen38525 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38526 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38527 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen38528 := Call(__e, gen38527, V5017)

										gen38529 := Call(__e, gen38526, gen38528)

										gen38530 := Call(__e, gen38525, gen38529)

										gen38531 := Call(__e, gen38524, Nil, gen38530)

										var gen38532 Obj

										if True == gen38531 {
											gen38514 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen38515 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38516 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38517 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38518 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38519 := Call(__e, gen38518, V5017)

											gen38520 := Call(__e, gen38517, gen38519)

											gen38521 := Call(__e, gen38516, gen38520)

											gen38522 := Call(__e, gen38515, gen38521)

											gen38523 := Call(__e, gen38514, gen38522, V5016)

											if True == gen38523 {
												gen38532 = True
											} else {
												gen38532 = False
											}

										} else {
											gen38532 = False
										}

										if True == gen38532 {
											gen38539 = True
										} else {
											gen38539 = False
										}

									} else {
										gen38539 = False
									}

									if True == gen38539 {
										gen38552 = True
									} else {
										gen38552 = False
									}

								} else {
									gen38552 = False
								}

								if True == gen38552 {
									gen38563 = True
								} else {
									gen38563 = False
								}

							} else {
								gen38563 = False
							}

							if True == gen38563 {
								gen38572 = True
							} else {
								gen38572 = False
							}

						} else {
							gen38572 = False
						}

						if True == gen38572 {
							gen38581 = True
						} else {
							gen38581 = False
						}

					} else {
						gen38581 = False
					}

					if True == gen38581 {
						gen38588 = True
					} else {
						gen38588 = False
					}

				} else {
					gen38588 = False
				}

				if True == gen38588 {
					gen38593 = True
				} else {
					gen38593 = False
				}

			} else {
				gen38593 = False
			}

			if True == gen38593 {
				gen38596 = True
			} else {
				gen38596 = False
			}

		} else {
			gen38596 = False
		}

		if True == gen38596 {
			gen38503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

			gen38504 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38505 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38506 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38507 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38508 := Call(__e, gen38507, V5017)

			gen38509 := Call(__e, gen38506, gen38508)

			gen38510 := Call(__e, gen38505, gen38509)

			gen38511 := Call(__e, gen38504, gen38510)

			gen38512 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38513 := Call(__e, gen38512, V5017)

			__e.TailApply(gen38503, gen38511, gen38513)

			return

		} else {
			gen38500 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38501 := Call(__e, gen38500, V5017)

			var gen38502 Obj

			if True == gen38501 {
				gen38495 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38496 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38497 := Call(__e, gen38496, V5017)

				gen38498 := Call(__e, gen38495, gen38497)

				var gen38499 Obj

				if True == gen38498 {
					gen38488 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38489 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38490 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38491 := Call(__e, gen38490, V5017)

					gen38492 := Call(__e, gen38489, gen38491)

					gen38493 := Call(__e, gen38488, gen38492)

					var gen38494 Obj

					if True == gen38493 {
						gen38479 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen38480 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38481 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38482 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38483 := Call(__e, gen38482, V5017)

						gen38484 := Call(__e, gen38481, gen38483)

						gen38485 := Call(__e, gen38480, gen38484)

						gen38486 := Call(__e, gen38479, Nil, gen38485)

						var gen38487 Obj

						if True == gen38486 {
							gen38473 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38474 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38475 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38476 := Call(__e, gen38475, V5017)

							gen38477 := Call(__e, gen38474, gen38476)

							gen38478 := Call(__e, gen38473, gen38477, V5016)

							if True == gen38478 {
								gen38487 = True
							} else {
								gen38487 = False
							}

						} else {
							gen38487 = False
						}

						if True == gen38487 {
							gen38494 = True
						} else {
							gen38494 = False
						}

					} else {
						gen38494 = False
					}

					if True == gen38494 {
						gen38499 = True
					} else {
						gen38499 = False
					}

				} else {
					gen38499 = False
				}

				if True == gen38499 {
					gen38502 = True
				} else {
					gen38502 = False
				}

			} else {
				gen38502 = False
			}

			if True == gen38502 {
				gen38466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

				gen38467 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38468 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38469 := Call(__e, gen38468, V5017)

				gen38470 := Call(__e, gen38467, gen38469)

				gen38471 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38472 := Call(__e, gen38471, V5017)

				__e.TailApply(gen38466, gen38470, gen38472)

				return

			} else {
				if True == True {
					__e.Return(V5017)
					return
				} else {
					gen38465 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen38465, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4false_1branch, gen38597)

	gen38673 := MakeNative(func(__e *ControlFlow) {
		V5020 := __e.Get(1)
		_ = V5020
		V5021 := __e.Get(2)
		_ = V5021
		gen38670 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen38671 := Call(__e, gen38670, V5020)

		var gen38672 Obj

		if True == gen38671 {
			gen38665 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen38666 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38667 := Call(__e, gen38666, V5020)

			gen38668 := Call(__e, gen38665, sym_f_freturn, gen38667)

			var gen38669 Obj

			if True == gen38668 {
				gen38660 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38661 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38662 := Call(__e, gen38661, V5020)

				gen38663 := Call(__e, gen38660, gen38662)

				var gen38664 Obj

				if True == gen38663 {
					gen38653 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38654 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38655 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38656 := Call(__e, gen38655, V5020)

					gen38657 := Call(__e, gen38654, gen38656)

					gen38658 := Call(__e, gen38653, Nil, gen38657)

					var gen38659 Obj

					if True == gen38658 {
						gen38645 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						gen38646 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38647 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38648 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38649 := Call(__e, gen38648, V5020)

						gen38650 := Call(__e, gen38647, gen38649)

						gen38651 := Call(__e, gen38646, gen38650)

						gen38652 := Call(__e, gen38645, gen38651)

						if True == gen38652 {
							gen38659 = True
						} else {
							gen38659 = False
						}

					} else {
						gen38659 = False
					}

					if True == gen38659 {
						gen38664 = True
					} else {
						gen38664 = False
					}

				} else {
					gen38664 = False
				}

				if True == gen38664 {
					gen38669 = True
				} else {
					gen38669 = False
				}

			} else {
				gen38669 = False
			}

			if True == gen38669 {
				gen38672 = True
			} else {
				gen38672 = False
			}

		} else {
			gen38672 = False
		}

		if True == gen38672 {
			__e.TailApply(V5021, V5020)

			return
		} else {
			gen38642 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38643 := Call(__e, gen38642, V5020)

			var gen38644 Obj

			if True == gen38643 {
				gen38637 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen38638 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38639 := Call(__e, gen38638, V5020)

				gen38640 := Call(__e, gen38637, symfail, gen38639)

				var gen38641 Obj

				if True == gen38640 {
					gen38633 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38634 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38635 := Call(__e, gen38634, V5020)

					gen38636 := Call(__e, gen38633, Nil, gen38635)

					if True == gen38636 {
						gen38641 = True
					} else {
						gen38641 = False
					}

				} else {
					gen38641 = False
				}

				if True == gen38641 {
					gen38644 = True
				} else {
					gen38644 = False
				}

			} else {
				gen38644 = False
			}

			if True == gen38644 {
				__e.TailApply(V5021, V5020)

				return
			} else {
				gen38630 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38631 := Call(__e, gen38630, V5020)

				var gen38632 Obj

				if True == gen38631 {
					gen38625 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38626 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38627 := Call(__e, gen38626, V5020)

					gen38628 := Call(__e, gen38625, sym_f_fgoto_1label, gen38627)

					var gen38629 Obj

					if True == gen38628 {
						gen38620 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38621 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38622 := Call(__e, gen38621, V5020)

						gen38623 := Call(__e, gen38620, gen38622)

						var gen38624 Obj

						if True == gen38623 {
							gen38614 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38615 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38616 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38617 := Call(__e, gen38616, V5020)

							gen38618 := Call(__e, gen38615, gen38617)

							gen38619 := Call(__e, gen38614, Nil, gen38618)

							if True == gen38619 {
								gen38624 = True
							} else {
								gen38624 = False
							}

						} else {
							gen38624 = False
						}

						if True == gen38624 {
							gen38629 = True
						} else {
							gen38629 = False
						}

					} else {
						gen38629 = False
					}

					if True == gen38629 {
						gen38632 = True
					} else {
						gen38632 = False
					}

				} else {
					gen38632 = False
				}

				if True == gen38632 {
					__e.TailApply(V5021, V5020)

					return
				} else {
					if True == True {
						gen38611 := MakeNative(func(__e *ControlFlow) {
							Label := __e.Get(1)
							_ = Label
							gen38599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen38600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen38601 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen38602 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen38603 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen38604 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen38605 := Call(__e, gen38604, Label, Nil)

							gen38606 := Call(__e, gen38603, sym_f_fgoto_1label, gen38605)

							gen38607 := Call(__e, V5021, gen38606)

							gen38608 := Call(__e, gen38602, gen38607, Nil)

							gen38609 := Call(__e, gen38601, V5020, gen38608)

							gen38610 := Call(__e, gen38600, Label, gen38609)

							__e.TailApply(gen38599, sym_f_flet_1label, gen38610)

							return

						}, 1)

						gen38612 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4generate_1label)

						gen38613 := Call(__e, gen38612)

						__e.TailApply(gen38611, gen38613)

						return

					} else {
						gen38598 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen38598, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4with_1labelled_1else, gen38673)

	gen38858 := MakeNative(func(__e *ControlFlow) {
		V5024 := __e.Get(1)
		_ = V5024
		gen38855 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen38856 := Call(__e, gen38855, V5024)

		var gen38857 Obj

		if True == gen38856 {
			gen38850 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen38851 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38852 := Call(__e, gen38851, V5024)

			gen38853 := Call(__e, gen38850, symif, gen38852)

			var gen38854 Obj

			if True == gen38853 {
				gen38845 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38846 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38847 := Call(__e, gen38846, V5024)

				gen38848 := Call(__e, gen38845, gen38847)

				var gen38849 Obj

				if True == gen38848 {
					gen38838 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38839 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38840 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38841 := Call(__e, gen38840, V5024)

					gen38842 := Call(__e, gen38839, gen38841)

					gen38843 := Call(__e, gen38838, gen38842)

					var gen38844 Obj

					if True == gen38843 {
						gen38829 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38830 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38831 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38832 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38833 := Call(__e, gen38832, V5024)

						gen38834 := Call(__e, gen38831, gen38833)

						gen38835 := Call(__e, gen38830, gen38834)

						gen38836 := Call(__e, gen38829, gen38835)

						var gen38837 Obj

						if True == gen38836 {
							gen38818 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38819 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38820 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen38821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38822 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38823 := Call(__e, gen38822, V5024)

							gen38824 := Call(__e, gen38821, gen38823)

							gen38825 := Call(__e, gen38820, gen38824)

							gen38826 := Call(__e, gen38819, gen38825)

							gen38827 := Call(__e, gen38818, symif, gen38826)

							var gen38828 Obj

							if True == gen38827 {
								gen38807 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen38808 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38809 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen38810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38811 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38812 := Call(__e, gen38811, V5024)

								gen38813 := Call(__e, gen38810, gen38812)

								gen38814 := Call(__e, gen38809, gen38813)

								gen38815 := Call(__e, gen38808, gen38814)

								gen38816 := Call(__e, gen38807, gen38815)

								var gen38817 Obj

								if True == gen38816 {
									gen38794 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen38795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38797 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen38798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38799 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen38800 := Call(__e, gen38799, V5024)

									gen38801 := Call(__e, gen38798, gen38800)

									gen38802 := Call(__e, gen38797, gen38801)

									gen38803 := Call(__e, gen38796, gen38802)

									gen38804 := Call(__e, gen38795, gen38803)

									gen38805 := Call(__e, gen38794, gen38804)

									var gen38806 Obj

									if True == gen38805 {
										gen38779 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen38780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38781 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38782 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38783 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen38784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen38786 := Call(__e, gen38785, V5024)

										gen38787 := Call(__e, gen38784, gen38786)

										gen38788 := Call(__e, gen38783, gen38787)

										gen38789 := Call(__e, gen38782, gen38788)

										gen38790 := Call(__e, gen38781, gen38789)

										gen38791 := Call(__e, gen38780, gen38790)

										gen38792 := Call(__e, gen38779, gen38791)

										var gen38793 Obj

										if True == gen38792 {
											gen38762 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen38763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38764 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38765 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38766 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38767 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen38768 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38769 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen38770 := Call(__e, gen38769, V5024)

											gen38771 := Call(__e, gen38768, gen38770)

											gen38772 := Call(__e, gen38767, gen38771)

											gen38773 := Call(__e, gen38766, gen38772)

											gen38774 := Call(__e, gen38765, gen38773)

											gen38775 := Call(__e, gen38764, gen38774)

											gen38776 := Call(__e, gen38763, gen38775)

											gen38777 := Call(__e, gen38762, Nil, gen38776)

											var gen38778 Obj

											if True == gen38777 {
												gen38753 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												gen38754 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen38755 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen38756 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen38757 := Call(__e, gen38756, V5024)

												gen38758 := Call(__e, gen38755, gen38757)

												gen38759 := Call(__e, gen38754, gen38758)

												gen38760 := Call(__e, gen38753, gen38759)

												var gen38761 Obj

												if True == gen38760 {
													gen38742 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen38743 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen38744 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen38745 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen38746 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen38747 := Call(__e, gen38746, V5024)

													gen38748 := Call(__e, gen38745, gen38747)

													gen38749 := Call(__e, gen38744, gen38748)

													gen38750 := Call(__e, gen38743, gen38749)

													gen38751 := Call(__e, gen38742, Nil, gen38750)

													var gen38752 Obj

													if True == gen38751 {
														gen38718 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														gen38719 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen38720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38722 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38723 := Call(__e, gen38722, V5024)

														gen38724 := Call(__e, gen38721, gen38723)

														gen38725 := Call(__e, gen38720, gen38724)

														gen38726 := Call(__e, gen38719, gen38725)

														gen38727 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen38728 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38730 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38731 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen38732 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38733 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen38734 := Call(__e, gen38733, V5024)

														gen38735 := Call(__e, gen38732, gen38734)

														gen38736 := Call(__e, gen38731, gen38735)

														gen38737 := Call(__e, gen38730, gen38736)

														gen38738 := Call(__e, gen38729, gen38737)

														gen38739 := Call(__e, gen38728, gen38738)

														gen38740 := Call(__e, gen38727, gen38739)

														gen38741 := Call(__e, gen38718, gen38726, gen38740)

														if True == gen38741 {
															gen38752 = True
														} else {
															gen38752 = False
														}

													} else {
														gen38752 = False
													}

													if True == gen38752 {
														gen38761 = True
													} else {
														gen38761 = False
													}

												} else {
													gen38761 = False
												}

												if True == gen38761 {
													gen38778 = True
												} else {
													gen38778 = False
												}

											} else {
												gen38778 = False
											}

											if True == gen38778 {
												gen38793 = True
											} else {
												gen38793 = False
											}

										} else {
											gen38793 = False
										}

										if True == gen38793 {
											gen38806 = True
										} else {
											gen38806 = False
										}

									} else {
										gen38806 = False
									}

									if True == gen38806 {
										gen38817 = True
									} else {
										gen38817 = False
									}

								} else {
									gen38817 = False
								}

								if True == gen38817 {
									gen38828 = True
								} else {
									gen38828 = False
								}

							} else {
								gen38828 = False
							}

							if True == gen38828 {
								gen38837 = True
							} else {
								gen38837 = False
							}

						} else {
							gen38837 = False
						}

						if True == gen38837 {
							gen38844 = True
						} else {
							gen38844 = False
						}

					} else {
						gen38844 = False
					}

					if True == gen38844 {
						gen38849 = True
					} else {
						gen38849 = False
					}

				} else {
					gen38849 = False
				}

				if True == gen38849 {
					gen38854 = True
				} else {
					gen38854 = False
				}

			} else {
				gen38854 = False
			}

			if True == gen38854 {
				gen38857 = True
			} else {
				gen38857 = False
			}

		} else {
			gen38857 = False
		}

		if True == gen38857 {
			gen38675 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38679 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38680 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38681 := Call(__e, gen38680, V5024)

			gen38682 := Call(__e, gen38679, gen38681)

			gen38683 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38684 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38685 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38686 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38688 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38689 := Call(__e, gen38688, V5024)

			gen38690 := Call(__e, gen38687, gen38689)

			gen38691 := Call(__e, gen38686, gen38690)

			gen38692 := Call(__e, gen38685, gen38691)

			gen38693 := Call(__e, gen38684, gen38692)

			gen38694 := Call(__e, gen38683, gen38693, Nil)

			gen38695 := Call(__e, gen38678, gen38682, gen38694)

			gen38696 := Call(__e, gen38677, symand, gen38695)

			gen38697 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen38698 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38699 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38701 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38702 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38703 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38704 := Call(__e, gen38703, V5024)

			gen38705 := Call(__e, gen38702, gen38704)

			gen38706 := Call(__e, gen38701, gen38705)

			gen38707 := Call(__e, gen38700, gen38706)

			gen38708 := Call(__e, gen38699, gen38707)

			gen38709 := Call(__e, gen38698, gen38708)

			gen38710 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38711 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38712 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38713 := Call(__e, gen38712, V5024)

			gen38714 := Call(__e, gen38711, gen38713)

			gen38715 := Call(__e, gen38710, gen38714)

			gen38716 := Call(__e, gen38697, gen38709, gen38715)

			gen38717 := Call(__e, gen38676, gen38696, gen38716)

			__e.TailApply(gen38675, symif, gen38717)

			return

		} else {
			if True == True {
				__e.Return(V5024)
				return
			} else {
				gen38674 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen38674, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs, gen38858)

	gen38862 := MakeNative(func(__e *ControlFlow) {
		V5027 := __e.Get(1)
		_ = V5027
		V5028 := __e.Get(2)
		_ = V5028
		gen38859 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		gen38860 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		gen38861 := Call(__e, gen38860, sym_c, V5028)

		__e.TailApply(gen38859, V5027, gen38861)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4concat_c, gen38862)

	gen38897 := MakeNative(func(__e *ControlFlow) {
		V5032 := __e.Get(1)
		_ = V5032
		gen38894 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen38895 := Call(__e, gen38894, V5032)

		var gen38896 Obj

		if True == gen38895 {
			gen38889 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38891 := Call(__e, gen38890, V5032)

			gen38892 := Call(__e, gen38889, gen38891)

			var gen38893 Obj

			if True == gen38892 {
				gen38882 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen38883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38884 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38885 := Call(__e, gen38884, V5032)

				gen38886 := Call(__e, gen38883, gen38885)

				gen38887 := Call(__e, gen38882, Nil, gen38886)

				var gen38888 Obj

				if True == gen38887 {
					gen38878 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

					gen38879 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38880 := Call(__e, gen38879, V5032)

					gen38881 := Call(__e, gen38878, gen38880)

					if True == gen38881 {
						gen38888 = True
					} else {
						gen38888 = False
					}

				} else {
					gen38888 = False
				}

				if True == gen38888 {
					gen38893 = True
				} else {
					gen38893 = False
				}

			} else {
				gen38893 = False
			}

			if True == gen38893 {
				gen38896 = True
			} else {
				gen38896 = False
			}

		} else {
			gen38896 = False
		}

		if True == gen38896 {
			gen38869 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4concat_c)

			gen38870 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4exp_1var)

			gen38871 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen38873 := Call(__e, gen38872, V5032)

			gen38874 := Call(__e, gen38871, gen38873)

			gen38875 := Call(__e, gen38870, gen38874)

			gen38876 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen38877 := Call(__e, gen38876, V5032)

			__e.TailApply(gen38869, gen38875, gen38877)

			return

		} else {
			gen38867 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen38868 := Call(__e, gen38867, V5032)

			if True == gen38868 {
				gen38864 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				gen38865 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38866 := Call(__e, gen38865, V5032)

				__e.TailApply(gen38864, gen38866)

				return

			} else {
				if True == True {
					__e.Return(V5032)
					return
				} else {
					gen38863 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen38863, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4exp_1var, gen38897)

	gen38901 := MakeNative(func(__e *ControlFlow) {
		V5035 := __e.Get(1)
		_ = V5035
		V5036 := __e.Get(2)
		_ = V5036
		gen38898 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)

		gen38899 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4test_1_6selectors)

		gen38900 := Call(__e, gen38899, V5035)

		__e.TailApply(gen38898, gen38900, V5036)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4optimize_1selectors, gen38901)

	gen39032 := MakeNative(func(__e *ControlFlow) {
		V5042 := __e.Get(1)
		_ = V5042
		gen39029 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen39030 := Call(__e, gen39029, V5042)

		var gen39031 Obj

		if True == gen39030 {
			gen39024 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen39025 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen39026 := Call(__e, gen39025, V5042)

			gen39027 := Call(__e, gen39024, symcons_2, gen39026)

			var gen39028 Obj

			if True == gen39027 {
				gen39019 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen39020 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen39021 := Call(__e, gen39020, V5042)

				gen39022 := Call(__e, gen39019, gen39021)

				var gen39023 Obj

				if True == gen39022 {
					gen39013 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen39014 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen39015 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen39016 := Call(__e, gen39015, V5042)

					gen39017 := Call(__e, gen39014, gen39016)

					gen39018 := Call(__e, gen39013, Nil, gen39017)

					if True == gen39018 {
						gen39023 = True
					} else {
						gen39023 = False
					}

				} else {
					gen39023 = False
				}

				if True == gen39023 {
					gen39028 = True
				} else {
					gen39028 = False
				}

			} else {
				gen39028 = False
			}

			if True == gen39028 {
				gen39031 = True
			} else {
				gen39031 = False
			}

		} else {
			gen39031 = False
		}

		if True == gen39031 {
			gen39002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen39003 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen39004 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen39005 := Call(__e, gen39004, V5042)

			gen39006 := Call(__e, gen39003, symhd, gen39005)

			gen39007 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen39008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen39009 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen39010 := Call(__e, gen39009, V5042)

			gen39011 := Call(__e, gen39008, symtl, gen39010)

			gen39012 := Call(__e, gen39007, gen39011, Nil)

			__e.TailApply(gen39002, gen39006, gen39012)

			return

		} else {
			gen38999 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen39000 := Call(__e, gen38999, V5042)

			var gen39001 Obj

			if True == gen39000 {
				gen38994 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen38995 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen38996 := Call(__e, gen38995, V5042)

				gen38997 := Call(__e, gen38994, symtuple_2, gen38996)

				var gen38998 Obj

				if True == gen38997 {
					gen38989 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38990 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38991 := Call(__e, gen38990, V5042)

					gen38992 := Call(__e, gen38989, gen38991)

					var gen38993 Obj

					if True == gen38992 {
						gen38983 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen38984 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38985 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38986 := Call(__e, gen38985, V5042)

						gen38987 := Call(__e, gen38984, gen38986)

						gen38988 := Call(__e, gen38983, Nil, gen38987)

						if True == gen38988 {
							gen38993 = True
						} else {
							gen38993 = False
						}

					} else {
						gen38993 = False
					}

					if True == gen38993 {
						gen38998 = True
					} else {
						gen38998 = False
					}

				} else {
					gen38998 = False
				}

				if True == gen38998 {
					gen39001 = True
				} else {
					gen39001 = False
				}

			} else {
				gen39001 = False
			}

			if True == gen39001 {
				gen38972 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38973 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38974 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38975 := Call(__e, gen38974, V5042)

				gen38976 := Call(__e, gen38973, symfst, gen38975)

				gen38977 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen38979 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen38980 := Call(__e, gen38979, V5042)

				gen38981 := Call(__e, gen38978, symsnd, gen38980)

				gen38982 := Call(__e, gen38977, gen38981, Nil)

				__e.TailApply(gen38972, gen38976, gen38982)

				return

			} else {
				gen38969 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen38970 := Call(__e, gen38969, V5042)

				var gen38971 Obj

				if True == gen38970 {
					gen38964 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen38965 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen38966 := Call(__e, gen38965, V5042)

					gen38967 := Call(__e, gen38964, symshen_4_7string_2, gen38966)

					var gen38968 Obj

					if True == gen38967 {
						gen38959 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen38960 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38961 := Call(__e, gen38960, V5042)

						gen38962 := Call(__e, gen38959, gen38961)

						var gen38963 Obj

						if True == gen38962 {
							gen38953 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen38954 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38955 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38956 := Call(__e, gen38955, V5042)

							gen38957 := Call(__e, gen38954, gen38956)

							gen38958 := Call(__e, gen38953, Nil, gen38957)

							if True == gen38958 {
								gen38963 = True
							} else {
								gen38963 = False
							}

						} else {
							gen38963 = False
						}

						if True == gen38963 {
							gen38968 = True
						} else {
							gen38968 = False
						}

					} else {
						gen38968 = False
					}

					if True == gen38968 {
						gen38971 = True
					} else {
						gen38971 = False
					}

				} else {
					gen38971 = False
				}

				if True == gen38971 {
					gen38942 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen38943 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen38944 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38945 := Call(__e, gen38944, V5042)

					gen38946 := Call(__e, gen38943, symhdstr, gen38945)

					gen38947 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen38948 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen38949 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen38950 := Call(__e, gen38949, V5042)

					gen38951 := Call(__e, gen38948, symtlstr, gen38950)

					gen38952 := Call(__e, gen38947, gen38951, Nil)

					__e.TailApply(gen38942, gen38946, gen38952)

					return

				} else {
					gen38939 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen38940 := Call(__e, gen38939, V5042)

					var gen38941 Obj

					if True == gen38940 {
						gen38934 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen38935 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen38936 := Call(__e, gen38935, V5042)

						gen38937 := Call(__e, gen38934, symshen_4_7vector_2, gen38936)

						var gen38938 Obj

						if True == gen38937 {
							gen38929 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen38930 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen38931 := Call(__e, gen38930, V5042)

							gen38932 := Call(__e, gen38929, gen38931)

							var gen38933 Obj

							if True == gen38932 {
								gen38923 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen38924 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38925 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen38926 := Call(__e, gen38925, V5042)

								gen38927 := Call(__e, gen38924, gen38926)

								gen38928 := Call(__e, gen38923, Nil, gen38927)

								if True == gen38928 {
									gen38933 = True
								} else {
									gen38933 = False
								}

							} else {
								gen38933 = False
							}

							if True == gen38933 {
								gen38938 = True
							} else {
								gen38938 = False
							}

						} else {
							gen38938 = False
						}

						if True == gen38938 {
							gen38941 = True
						} else {
							gen38941 = False
						}

					} else {
						gen38941 = False
					}

					if True == gen38941 {
						gen38912 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen38913 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen38914 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38915 := Call(__e, gen38914, V5042)

						gen38916 := Call(__e, gen38913, symhdv, gen38915)

						gen38917 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen38918 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen38919 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen38920 := Call(__e, gen38919, V5042)

						gen38921 := Call(__e, gen38918, symtlv, gen38920)

						gen38922 := Call(__e, gen38917, gen38921, Nil)

						__e.TailApply(gen38912, gen38916, gen38922)

						return

					} else {
						if True == True {
							gen38907 := MakeNative(func(__e *ControlFlow) {
								Result := __e.Get(1)
								_ = Result
								gen38903 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen38904 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								gen38905 := Call(__e, gen38904)

								gen38906 := Call(__e, gen38903, Result, gen38905)

								if True == gen38906 {
									__e.Return(Nil)
									return
								} else {
									__e.Return(Result)
									return
								}

							}, 1)

							gen38908 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4apply_1selector_1handlers)

							gen38909 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

							gen38910 := Call(__e, gen38909, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

							gen38911 := Call(__e, gen38908, gen38910, V5042)

							__e.TailApply(gen38907, gen38911)

							return

						} else {
							gen38902 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen38902, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4test_1_6selectors, gen39032)

	gen39046 := MakeNative(func(__e *ControlFlow) {
		V5045 := __e.Get(1)
		_ = V5045
		V5046 := __e.Get(2)
		_ = V5046
		gen39044 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen39045 := Call(__e, gen39044, V5045)

		if True == gen39045 {
			gen39037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4bind_1selector)

			gen39038 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen39039 := Call(__e, gen39038, V5045)

			gen39040 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)

			gen39041 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen39042 := Call(__e, gen39041, V5045)

			gen39043 := Call(__e, gen39040, gen39042, V5046)

			__e.TailApply(gen39037, gen39039, gen39043)

			return

		} else {
			gen39035 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen39036 := Call(__e, gen39035, Nil, V5045)

			if True == gen39036 {
				__e.Return(V5046)
				return
			} else {
				if True == True {
					gen39034 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen39034, symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)

					return

				} else {
					gen39033 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen39033, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors, gen39046)

	gen39064 := MakeNative(func(__e *ControlFlow) {
		V5053 := __e.Get(1)
		_ = V5053
		V5054 := __e.Get(2)
		_ = V5054
		gen39060 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		gen39061 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

		gen39062 := Call(__e, gen39061, V5053, V5054)

		gen39063 := Call(__e, gen39060, gen39062, MakeNumber(1))

		if True == gen39063 {
			gen39057 := MakeNative(func(__e *ControlFlow) {
				Var := __e.Get(1)
				_ = Var
				gen39048 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen39049 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen39050 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen39051 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen39052 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

				gen39053 := Call(__e, gen39052, Var, V5053, V5054)

				gen39054 := Call(__e, gen39051, gen39053, Nil)

				gen39055 := Call(__e, gen39050, V5053, gen39054)

				gen39056 := Call(__e, gen39049, Var, gen39055)

				__e.TailApply(gen39048, symlet, gen39056)

				return

			}, 1)

			gen39058 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4exp_1var)

			gen39059 := Call(__e, gen39058, V5053)

			__e.TailApply(gen39057, gen39059)

			return

		} else {
			if True == True {
				__e.Return(V5054)
				return
			} else {
				gen39047 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen39047, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4bind_1selector, gen39064)

	gen39090 := MakeNative(func(__e *ControlFlow) {
		V5067 := __e.Get(1)
		_ = V5067
		V5068 := __e.Get(2)
		_ = V5068
		gen39088 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen39089 := Call(__e, gen39088, Nil, V5067)

		if True == gen39089 {
			gen39087 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen39087)

			return

		} else {
			if True == True {
				gen39078 := MakeNative(func(__e *ControlFlow) {
					Freeze := __e.Get(1)
					_ = Freeze
					gen39076 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen39077 := Call(__e, gen39076, V5067)

					if True == gen39077 {
						gen39072 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							gen39068 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen39069 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen39070 := Call(__e, gen39069)

							gen39071 := Call(__e, gen39068, Result, gen39070)

							if True == gen39071 {
								gen39067 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

								__e.TailApply(gen39067, Freeze)

								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						gen39073 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen39074 := Call(__e, gen39073, V5067)

						gen39075 := Call(__e, gen39074, V5068)

						__e.TailApply(gen39072, gen39075)

						return

					} else {
						gen39066 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

						__e.TailApply(gen39066, Freeze)

						return

					}

				}, 1)

				gen39086 := MakeNative(func(__e *ControlFlow) {
					gen39084 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen39085 := Call(__e, gen39084, V5067)

					if True == gen39085 {
						gen39081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4apply_1selector_1handlers)

						gen39082 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen39083 := Call(__e, gen39082, V5067)

						__e.TailApply(gen39081, gen39083, V5068)

						return

					} else {
						if True == True {
							gen39080 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen39080, symshen_4x_4factorise_1defun_4apply_1selector_1handlers)

							return

						} else {
							gen39079 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen39079, MakeString("no cond match"))

							return

						}
					}

				}, 0)

				__e.TailApply(gen39078, gen39086)

				return

			} else {
				gen39065 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen39065, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4apply_1selector_1handlers, gen39090)

	gen39093 := MakeNative(func(__e *ControlFlow) {
		gen39091 := Call(__e, PrimNS1Value(symns2_1value), symset)

		Call(__e, gen39091, symshen_4x_4factorise_1defun_4_dselector_1handlers_d, Nil)

		gen39092 := Call(__e, PrimNS1Value(symns2_1value), symset)

		Call(__e, gen39092, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, Nil)

		__e.Return(symshen_4x_4factorise_1defun_4done)
		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4initialise, gen39093)

	gen39111 := MakeNative(func(__e *ControlFlow) {
		V5070 := __e.Get(1)
		_ = V5070
		gen39107 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen39108 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen39109 := Call(__e, gen39108, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

		gen39110 := Call(__e, gen39107, V5070, gen39109)

		if True == gen39110 {
			__e.Return(V5070)
			return
		} else {
			if True == True {
				gen39095 := Call(__e, PrimNS1Value(symns2_1value), symset)

				gen39096 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen39097 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				gen39098 := Call(__e, gen39097, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

				gen39099 := Call(__e, gen39096, V5070, gen39098)

				Call(__e, gen39095, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, gen39099)

				gen39100 := Call(__e, PrimNS1Value(symns2_1value), symset)

				gen39101 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen39102 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

				gen39103 := Call(__e, gen39102, V5070)

				gen39104 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				gen39105 := Call(__e, gen39104, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

				gen39106 := Call(__e, gen39101, gen39103, gen39105)

				Call(__e, gen39100, symshen_4x_4factorise_1defun_4_dselector_1handlers_d, gen39106)

				__e.Return(V5070)
				return

			} else {
				gen39094 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen39094, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4register_1selector_1handler, gen39111)

	gen39118 := MakeNative(func(__e *ControlFlow) {
		V5073 := __e.Get(1)
		_ = V5073
		V5074 := __e.Get(2)
		_ = V5074
		gen39113 := MakeNative(func(__e *ControlFlow) {
			gen39112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findpos)

			__e.TailApply(gen39112, V5073, V5074)

			return

		}, 0)

		gen39117 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			gen39114 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen39115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen39116 := Call(__e, gen39115, V5073, MakeString(" is not a selector handler\n"), symshen_4a)

			__e.TailApply(gen39114, gen39116)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen39113, gen39117)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4findpos, gen39118)

	gen39137 := MakeNative(func(__e *ControlFlow) {
		V5076 := __e.Get(1)
		_ = V5076
		gen39134 := MakeNative(func(__e *ControlFlow) {
			Reg := __e.Get(1)
			_ = Reg
			gen39131 := MakeNative(func(__e *ControlFlow) {
				Pos := __e.Get(1)
				_ = Pos
				gen39126 := MakeNative(func(__e *ControlFlow) {
					RemoveReg := __e.Get(1)
					_ = RemoveReg
					gen39119 := MakeNative(func(__e *ControlFlow) {
						RemoveFun := __e.Get(1)
						_ = RemoveFun
						__e.Return(V5076)
						return
					}, 1)

					gen39120 := Call(__e, PrimNS1Value(symns2_1value), symset)

					gen39121 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1nth)

					gen39122 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					gen39123 := Call(__e, gen39122, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

					gen39124 := Call(__e, gen39121, Pos, gen39123)

					gen39125 := Call(__e, gen39120, symshen_4x_4factorise_1defun_4_dselector_1handlers_d, gen39124)

					__e.TailApply(gen39119, gen39125)

					return

				}, 1)

				gen39127 := Call(__e, PrimNS1Value(symns2_1value), symset)

				gen39128 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				gen39129 := Call(__e, gen39128, V5076, Reg)

				gen39130 := Call(__e, gen39127, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, gen39129)

				__e.TailApply(gen39126, gen39130)

				return

			}, 1)

			gen39132 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4findpos)

			gen39133 := Call(__e, gen39132, V5076, Reg)

			__e.TailApply(gen39131, gen39133)

			return

		}, 1)

		gen39135 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen39136 := Call(__e, gen39135, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d)

		__e.TailApply(gen39134, gen39136)

		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4unregister_1selector_1handler, gen39137)

	return

}, 0)
var symshen_4initialise_1prolog = MakeSymbol("shen.initialise-prolog")
var symshen_4errormaxinfs = MakeSymbol("shen.errormaxinfs")
var symtlv = MakeSymbol("tlv")
var symshen_4sys_1error = MakeSymbol("shen.sys-error")
var symread = MakeSymbol("read")
var symshen_4construct_1search_1literals = MakeSymbol("shen.construct-search-literals")
var symshen_4include_1h = MakeSymbol("shen.include-h")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4initialise_1signedfunc_1lambda_1forms = MakeSymbol("shen.initialise-signedfunc-lambda-forms")
var symshen_4space = MakeSymbol("shen.space")
var symshen_4free__variable__check = MakeSymbol("shen.free_variable_check")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4type_1signature_1of_1element_2 = MakeSymbol("shen.type-signature-of-element?")
var symns2_1set = MakeSymbol("ns2-set")
var symtc = MakeSymbol("tc")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4lineread_1loop = MakeSymbol("shen.lineread-loop")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4type_1signature_1of_1spy = MakeSymbol("shen.type-signature-of-spy")
var symstoutput = MakeSymbol("stoutput")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4yacc = MakeSymbol("shen.yacc")
var sympackage_2 = MakeSymbol("package?")
var symshen_4variancy_1test = MakeSymbol("shen.variancy-test")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4x_4factorise_1defun_4unregister_1selector_1handler = MakeSymbol("shen.x.factorise-defun.unregister-selector-handler")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var syminput = MakeSymbol("input")
var symshen_4decimalise = MakeSymbol("shen.decimalise")
var symshen_4of = MakeSymbol("shen.of")
var symshen_4unbindv = MakeSymbol("shen.unbindv")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symboolean = MakeSymbol("boolean")
var symshen_4update_1demodulation_1function = MakeSymbol("shen.update-demodulation-function")
var symshen_4_a_a_6 = MakeSymbol("shen.==>")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var sym_f_flet_1label = MakeSymbol("%%let-label")
var symhdv = MakeSymbol("hdv")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symunspecialise = MakeSymbol("unspecialise")
var symT = MakeSymbol("T")
var sym_dargv_d = MakeSymbol("*argv*")
var sym_6_a = MakeSymbol(">=")
var symAssumptions__1957 = MakeSymbol("Assumptions_1957")
var symshen_4type_1signature_1of_1gensym = MakeSymbol("shen.type-signature-of-gensym")
var symshen_4curry_1type_1h = MakeSymbol("shen.curry-type-h")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symshen_4jump__stream = MakeSymbol("shen.jump_stream")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4type_1signature_1of_1pos = MakeSymbol("shen.type-signature-of-pos")
var symshen_4trim_1gubbins = MakeSymbol("shen.trim-gubbins")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4recursive__descent = MakeSymbol("shen.recursive_descent")
var symshen_4type_1signature_1of_1_6_a = MakeSymbol("shen.type-signature-of->=")
var symassoc = MakeSymbol("assoc")
var symshen_4_dmacroreg_d = MakeSymbol("shen.*macroreg*")
var symshen_4abs = MakeSymbol("shen.abs")
var symwarn = MakeSymbol("warn")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4pop = MakeSymbol("shen.pop")
var symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d = MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*")
var symshen_4_5premises_6 = MakeSymbol("shen.<premises>")
var symshen_4record_1it_1h = MakeSymbol("shen.record-it-h")
var symshen_4group__clauses = MakeSymbol("shen.group_clauses")
var symshen_4catchpoint_b = MakeSymbol("shen.catchpoint!")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4ue_1sig = MakeSymbol("shen.ue-sig")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4terminator_2 = MakeSymbol("shen.terminator?")
var symmap = MakeSymbol("map")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4compile_1macro = MakeSymbol("shen.compile-macro")
var symshen_4type_1signature_1of_1trap_1error = MakeSymbol("shen.type-signature-of-trap-error")
var sym_1 = MakeSymbol("-")
var symshen_4compile__to__machine__code = MakeSymbol("shen.compile_to_machine_code")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symreturn = MakeSymbol("return")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symloaded = MakeSymbol("loaded")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4encode_1choices = MakeSymbol("shen.encode-choices")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4_5byte_6 = MakeSymbol("shen.<byte>")
var symshen_4same__predicate_2 = MakeSymbol("shen.same_predicate?")
var symFinish = MakeSymbol("Finish")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4insert_1predicate = MakeSymbol("shen.insert-predicate")
var symshen_4type_1signature_1of_1occurrences = MakeSymbol("shen.type-signature-of-occurrences")
var symshen_4type_1signature_1of_1unprofile = MakeSymbol("shen.type-signature-of-unprofile")
var sym__ = MakeSymbol("_")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symnumber_2 = MakeSymbol("number?")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4cn_1all = MakeSymbol("shen.cn-all")
var symshen_4type_1signature_1of_1map = MakeSymbol("shen.type-signature-of-map")
var symYaccParse = MakeSymbol("YaccParse")
var symshen_4patthyps = MakeSymbol("shen.patthyps")
var symadjoin = MakeSymbol("adjoin")
var symconcat = MakeSymbol("concat")
var symshen_4synonyms_1help = MakeSymbol("shen.synonyms-help")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symshen_4read_1file_1as_1charlist = MakeSymbol("shen.read-file-as-charlist")
var symshen_4ues = MakeSymbol("shen.ues")
var symshen_4newhyps = MakeSymbol("shen.newhyps")
var symshen_4type_1signature_1of_1_5_a = MakeSymbol("shen.type-signature-of-<=")
var symshen_4multiple_1set = MakeSymbol("shen.multiple-set")
var symshen_4copy_1vector_1stage_11 = MakeSymbol("shen.copy-vector-stage-1")
var symis = MakeSymbol("is")
var symshen_4make__mu__application = MakeSymbol("shen.make_mu_application")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symput = MakeSymbol("put")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4rename = MakeSymbol("shen.rename")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4type_1signature_1of_1_5 = MakeSymbol("shen.type-signature-of-<")
var symshen_4type_1signature_1of_1_6 = MakeSymbol("shen.type-signature-of->")
var symshen_4interror = MakeSymbol("shen.interror")
var symshen_4extract__free__vars = MakeSymbol("shen.extract_free_vars")
var sym_k = MakeSymbol(";")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4type_1signature_1of_1assoc = MakeSymbol("shen.type-signature-of-assoc")
var symshen_4type_1signature_1of_1function = MakeSymbol("shen.type-signature-of-function")
var symshen_4type_1signature_1of_1get_1time = MakeSymbol("shen.type-signature-of-get-time")
var symshen_4type_1signature_1of_1_1 = MakeSymbol("shen.type-signature-of--")
var symshen_4atom_1type = MakeSymbol("shen.atom-type")
var symshen_4intprolog_1help_1help = MakeSymbol("shen.intprolog-help-help")
var symshen_4udefs_d = MakeSymbol("shen.udefs*")
var symshen_4type_1signature_1of_1tlstr = MakeSymbol("shen.type-signature-of-tlstr")
var symshen_4type_1signature_1of_1_c = MakeSymbol("shen.type-signature-of-/")
var symlineread = MakeSymbol("lineread")
var symshen_4type_1signature_1of_1untrack = MakeSymbol("shen.type-signature-of-untrack")
var symM = MakeSymbol("M")
var symnot = MakeSymbol("not")
var symboolean_2 = MakeSymbol("boolean?")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4compress_150 = MakeSymbol("shen.compress-50")
var symD = MakeSymbol("D")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symE = MakeSymbol("E")
var symshen_4explicit__modes = MakeSymbol("shen.explicit_modes")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symFreeze = MakeSymbol("Freeze")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var sym_3 = MakeSymbol("$")
var symstep = MakeSymbol("step")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4get_1type = MakeSymbol("shen.get-type")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4type_1signature_1of_1systemf = MakeSymbol("shen.type-signature-of-systemf")
var sym_7 = MakeSymbol("+")
var symshen_4prolog_1failure = MakeSymbol("shen.prolog-failure")
var symshen_4type_1signature_1of_1not = MakeSymbol("shen.type-signature-of-not")
var sym_h_a = MakeSymbol(":=")
var symdefcc = MakeSymbol("defcc")
var symshen_4cutpoint = MakeSymbol("shen.cutpoint")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4x_4factorise_1defun_4optimize_1selectors = MakeSymbol("shen.x.factorise-defun.optimize-selectors")
var symsimple_1error = MakeSymbol("simple-error")
var symlimit = MakeSymbol("limit")
var symshen_4_5side_1conditions_6 = MakeSymbol("shen.<side-conditions>")
var symshen_4_5side_1condition_6 = MakeSymbol("shen.<side-condition>")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4copy_1vector = MakeSymbol("shen.copy-vector")
var symshen_4read_1file_1as_1Xlist_1help = MakeSymbol("shen.read-file-as-Xlist-help")
var symfwhen = MakeSymbol("fwhen")
var symshen_4mu__reduction = MakeSymbol("shen.mu_reduction")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4type_1signature_1of_1read_1from_1string = MakeSymbol("shen.type-signature-of-read-from-string")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symnumber = MakeSymbol("number")
var sympos = MakeSymbol("pos")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var symshen_4typedf_2 = MakeSymbol("shen.typedf?")
var symshen_4type_1signature_1of_1and = MakeSymbol("shen.type-signature-of-and")
var sym_l = MakeSymbol(",")
var symshen_4elim_1def = MakeSymbol("shen.elim-def")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var symshen_4function_1abstraction = MakeSymbol("shen.function-abstraction")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4type_1signature_1of_1profile = MakeSymbol("shen.type-signature-of-profile")
var symshen_4reassemble = MakeSymbol("shen.reassemble")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symG = MakeSymbol("G")
var symshen_4_5head_d_6 = MakeSymbol("shen.<head*>")
var symshen_4mu = MakeSymbol("shen.mu")
var symshen_4pair = MakeSymbol("shen.pair")
var symstr = MakeSymbol("str")
var symshen_4embed_1and = MakeSymbol("shen.embed-and")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4cl = MakeSymbol("shen.cl")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4type_1signature_1of_1unspecialise = MakeSymbol("shen.type-signature-of-unspecialise")
var symshen_4prefix_2 = MakeSymbol("shen.prefix?")
var symshen_4_dcustom_1pattern_1compiler_d = MakeSymbol("shen.*custom-pattern-compiler*")
var symmapcan = MakeSymbol("mapcan")
var symshen_4_5conclusion_6 = MakeSymbol("shen.<conclusion>")
var symshen_4initialise__arity__table = MakeSymbol("shen.initialise_arity_table")
var symshen_4th_d = MakeSymbol("shen.th*")
var sym_i = MakeSymbol("{")
var symbound_2 = MakeSymbol("bound?")
var symshen_4_5singleunderline_6 = MakeSymbol("shen.<singleunderline>")
var symshen_4check__stream = MakeSymbol("shen.check_stream")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4x_4factorise_1defun_4apply_1selector_1handlers = MakeSymbol("shen.x.factorise-defun.apply-selector-handlers")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4insert_1deref = MakeSymbol("shen.insert-deref")
var symshen_4to = MakeSymbol("shen.to")
var symshen_4type_1signature_1of_1_7 = MakeSymbol("shen.type-signature-of-+")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symshen_4cond_1expression = MakeSymbol("shen.cond-expression")
var symit = MakeSymbol("it")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symset = MakeSymbol("set")
var symshen_4control_1chars = MakeSymbol("shen.control-chars")
var symshen_4collect = MakeSymbol("shen.collect")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4x_4factorise_1defun_4with_1labelled_1else = MakeSymbol("shen.x.factorise-defun.with-labelled-else")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symContinuation = MakeSymbol("Continuation")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symshen_4type_1signature_1of_1error_1to_1string = MakeSymbol("shen.type-signature-of-error-to-string")
var symN = MakeSymbol("N")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symmaxinferences = MakeSymbol("maxinferences")
var symreceive = MakeSymbol("receive")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4type_1signature_1of_1simple_1error = MakeSymbol("shen.type-signature-of-simple-error")
var symF = MakeSymbol("F")
var sym_dport_d = MakeSymbol("*port*")
var symshen_4update__history = MakeSymbol("shen.update_history")
var symshen_4shen_1syntax_1error = MakeSymbol("shen.shen-syntax-error")
var symshen_4start_1new_1prolog_1process = MakeSymbol("shen.start-new-prolog-process")
var symshen_4type_1signature_1of_1close = MakeSymbol("shen.type-signature-of-close")
var symshen_4code_1point = MakeSymbol("shen.code-point")
var symshen_4type_1signature_1of_1preclude = MakeSymbol("shen.type-signature-of-preclude")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var sym_8s = MakeSymbol("@s")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4double = MakeSymbol("shen.double")
var symshen_4ok = MakeSymbol("shen.ok")
var symabsvector = MakeSymbol("absvector")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4strip_1pathname = MakeSymbol("shen.strip-pathname")
var symshen_4symbol_1code_2 = MakeSymbol("shen.symbol-code?")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4complexity__head = MakeSymbol("shen.complexity_head")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symshen_4clauses_1to_1shen = MakeSymbol("shen.clauses-to-shen")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symK = MakeSymbol("K")
var symshen_4result_1type = MakeSymbol("shen.result-type")
var symlazy = MakeSymbol("lazy")
var symbar_b = MakeSymbol("bar!")
var symshen_4fail__if = MakeSymbol("shen.fail_if")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4remove_1bar = MakeSymbol("shen.remove-bar")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4type_1signature_1of_1load = MakeSymbol("shen.type-signature-of-load")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4package_1macro = MakeSymbol("shen.package-macro")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4assign_1types = MakeSymbol("shen.assign-types")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4t_d_1defh = MakeSymbol("shen.t*-defh")
var symshen_4x_4features_4initialise = MakeSymbol("shen.x.features.initialise")
var symshen_4type_1signature_1of_1bound_2 = MakeSymbol("shen.type-signature-of-bound?")
var symshen_4type_1signature_1of_1enable_1type_1theory = MakeSymbol("shen.type-signature-of-enable-type-theory")
var symelement_2 = MakeSymbol("element?")
var symshen_4_dteststack_d = MakeSymbol("shen.*teststack*")
var symshen_4doubleunderline_2 = MakeSymbol("shen.doubleunderline?")
var sym_b = MakeSymbol("!")
var symshen_4compile__prolog__procedure = MakeSymbol("shen.compile_prolog_procedure")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4variable_1match = MakeSymbol("shen.variable-match")
var symshen_4_5non_1return_6 = MakeSymbol("shen.<non-return>")
var sympr = MakeSymbol("pr")
var symshen_4rule_1_6horn_1clause_1head = MakeSymbol("shen.rule->horn-clause-head")
var symshen_4type_1signature_1of_1thaw = MakeSymbol("shen.type-signature-of-thaw")
var symfix = MakeSymbol("fix")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var sym_e_e_e = MakeSymbol("&&&")
var symshen_4typextable = MakeSymbol("shen.typextable")
var symdo = MakeSymbol("do")
var symshen_4type_1signature_1of_1it = MakeSymbol("shen.type-signature-of-it")
var symshen_4type_1signature_1of_1number_2 = MakeSymbol("shen.type-signature-of-number?")
var symshen_4x_4factorise_1defun_4free_1variables_1h = MakeSymbol("shen.x.factorise-defun.free-variables-h")
var symshen_4dh_2 = MakeSymbol("shen.dh?")
var symshen_4remove_1nth = MakeSymbol("shen.remove-nth")
var symshen_4type_1signature_1of_1preclude_1all_1but = MakeSymbol("shen.type-signature-of-preclude-all-but")
var symshen_4type_1signature_1of_1stoutput = MakeSymbol("shen.type-signature-of-stoutput")
var symshen_4_5signature_1help_6 = MakeSymbol("shen.<signature-help>")
var symshen_4the = MakeSymbol("shen.the")
var symshen_4else = MakeSymbol("shen.else")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var sym_6_6 = MakeSymbol(">>")
var symshen_4variables = MakeSymbol("shen.variables")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4void = MakeSymbol("shen.void")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4x_4factorise_1defun_4add_1returns = MakeSymbol("shen.x.factorise-defun.add-returns")
var symshen_4resizeprocessvector = MakeSymbol("shen.resizeprocessvector")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symcond = MakeSymbol("cond")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4s_1prolog__clause = MakeSymbol("shen.s-prolog_clause")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4cf__help = MakeSymbol("shen.cf_help")
var symshen_4cc__help = MakeSymbol("shen.cc_help")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4t_d_1action = MakeSymbol("shen.t*-action")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symempty_2 = MakeSymbol("empty?")
var symshen_4_5simple__pattern_6 = MakeSymbol("shen.<simple_pattern>")
var symshen_4insert_1runon = MakeSymbol("shen.insert-runon")
var symshen_4type_1signature_1of_1hdstr = MakeSymbol("shen.type-signature-of-hdstr")
var symshen_4type_1signature_1of_1step = MakeSymbol("shen.type-signature-of-step")
var symshen_4construct_1base_1search_1clause = MakeSymbol("shen.construct-base-search-clause")
var symverified = MakeSymbol("verified")
var symshen_4remember = MakeSymbol("shen.remember")
var symshen_4type_1signature_1of_1string_1_6n = MakeSymbol("shen.type-signature-of-string->n")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var symshen_4_dcustom_1pattern_1reducer_d = MakeSymbol("shen.*custom-pattern-reducer*")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symContext__1957 = MakeSymbol("Context_1957")
var symshen_4insert__modes = MakeSymbol("shen.insert_modes")
var symshen_4call_1rest = MakeSymbol("shen.call-rest")
var symshen_4type_1signature_1of_1include_1all_1but = MakeSymbol("shen.type-signature-of-include-all-but")
var symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs = MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs")
var symV = MakeSymbol("V")
var symshen_4_5doubleunderline_6 = MakeSymbol("shen.<doubleunderline>")
var symshen_4base = MakeSymbol("shen.base")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4left_1round = MakeSymbol("shen.left-round")
var symshen_4flatten = MakeSymbol("shen.flatten")
var symshen_4_dsignedfuncs_d = MakeSymbol("shen.*signedfuncs*")
var symlambda = MakeSymbol("lambda")
var symstream = MakeSymbol("stream")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symstring_2 = MakeSymbol("string?")
var syminferences = MakeSymbol("inferences")
var symshen_4x_4factorise_1defun_4attach_1free_1variables = MakeSymbol("shen.x.factorise-defun.attach-free-variables")
var sym_f_flabel = MakeSymbol("%%label")
var symcompile = MakeSymbol("compile")
var symshen_4typecheck_1and_1evaluate = MakeSymbol("shen.typecheck-and-evaluate")
var symget = MakeSymbol("get")
var sym_1_6 = MakeSymbol("->")
var symshen_4type_1signature_1of_1inferences = MakeSymbol("shen.type-signature-of-inferences")
var sym_f_freturn = MakeSymbol("%%return")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4_5sig_7rules_6 = MakeSymbol("shen.<sig+rules>")
var symshen_4_5st__input1_6 = MakeSymbol("shen.<st_input1>")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symshen_4type_1signature_1of_1_5_1vector = MakeSymbol("shen.type-signature-of-<-vector")
var symshen_4require = MakeSymbol("shen.require")
var sym_5 = MakeSymbol("<")
var symshen_4_5clause_d_6 = MakeSymbol("shen.<clause*>")
var symshen_4left_1rule = MakeSymbol("shen.left-rule")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4rule_1_6horn_1clause = MakeSymbol("shen.rule->horn-clause")
var symshen_4write_1char_1and_1inc = MakeSymbol("shen.write-char-and-inc")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4type_1signature_1of_1read_1file_1as_1string = MakeSymbol("shen.type-signature-of-read-file-as-string")
var symshen_4x_4factorise_1defun_4test_1_6selectors = MakeSymbol("shen.x.factorise-defun.test->selectors")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4type_1signature_1of_1read_1byte = MakeSymbol("shen.type-signature-of-read-byte")
var symcn = MakeSymbol("cn")
var symshen_4rules_1_6horn_1clauses = MakeSymbol("shen.rules->horn-clauses")
var symshen_4bld_1prolog_1call = MakeSymbol("shen.bld-prolog-call")
var symshen_4type_1signature_1of_1print = MakeSymbol("shen.type-signature-of-print")
var symR = MakeSymbol("R")
var symP = MakeSymbol("P")
var sym_6 = MakeSymbol(">")
var symIn__1957 = MakeSymbol("In_1957")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4toplevel = MakeSymbol("shen.toplevel")
var symshen_4alphanum_2 = MakeSymbol("shen.alphanum?")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symexception = MakeSymbol("exception")
var symshen_4type_1signature_1of_1tc = MakeSymbol("shen.type-signature-of-tc")
var symshow_1help = MakeSymbol("show-help")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symreverse = MakeSymbol("reverse")
var symshen_4digits_1_6integers = MakeSymbol("shen.digits->integers")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4continuation__call = MakeSymbol("shen.continuation_call")
var symshen_4_dvarcounter_d = MakeSymbol("shen.*varcounter*")
var symshen_4toplevel__evaluate = MakeSymbol("shen.toplevel_evaluate")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symdefine = MakeSymbol("define")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4variable = MakeSymbol("shen.variable")
var symshen_4ue_2 = MakeSymbol("shen.ue?")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4type_1signature_1of_1empty_2 = MakeSymbol("shen.type-signature-of-empty?")
var symshen_4nextline = MakeSymbol("shen.nextline")
var symvariable_2 = MakeSymbol("variable?")
var symAssumption__1957 = MakeSymbol("Assumption_1957")
var symThrowcontrol = MakeSymbol("Throwcontrol")
var symspy = MakeSymbol("spy")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4catchpoint = MakeSymbol("shen.catchpoint")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4app = MakeSymbol("shen.app")
var symtype = MakeSymbol("type")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4sequent = MakeSymbol("shen.sequent")
var symshen_4post = MakeSymbol("shen.post")
var symshen_4type_1signature_1of_1occurs_1check = MakeSymbol("shen.type-signature-of-occurs-check")
var symthaw = MakeSymbol("thaw")
var symshen_4mod = MakeSymbol("shen.mod")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4type_1signature_1of_1if = MakeSymbol("shen.type-signature-of-if")
var symshen_4type_1signature_1of_1shen_4insert = MakeSymbol("shen.type-signature-of-shen.insert")
var sym_a_a = MakeSymbol("==")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symshen_4type_1signature_1of_1fail = MakeSymbol("shen.type-signature-of-fail")
var symshen_4type_1signature_1of_1y_1or_1n_2 = MakeSymbol("shen.type-signature-of-y-or-n?")
var symI = MakeSymbol("I")
var symshen_4read_1char_1code = MakeSymbol("shen.read-char-code")
var symshen_4type_1signature_1of_1sterror = MakeSymbol("shen.type-signature-of-sterror")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symResult = MakeSymbol("Result")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4safe_1multiply = MakeSymbol("shen.safe-multiply")
var symshen_4compose = MakeSymbol("shen.compose")
var symshen_4type_1signature_1of_1 = MakeSymbol("shen.type-signature-of-")
var symshen_4active_1cons = MakeSymbol("shen.active-cons")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4type_1signature_1of_1string_1_6symbol = MakeSymbol("shen.type-signature-of-string->symbol")
var symshen_4x_4factorise_1defun_4true_1branch = MakeSymbol("shen.x.factorise-defun.true-branch")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symunit = MakeSymbol("unit")
var symshen_4type_1signature_1of_1package_2 = MakeSymbol("shen.type-signature-of-package?")
var symshen_4prbytes = MakeSymbol("shen.prbytes")
var symQv = MakeSymbol("Qv")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4next_150 = MakeSymbol("shen.next-50")
var symshen_4_5action_6 = MakeSymbol("shen.<action>")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4package_1contents = MakeSymbol("shen.package-contents")
var symL = MakeSymbol("L")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4placeholders = MakeSymbol("shen.placeholders")
var symshen_4x_4factorise_1defun_4findpos = MakeSymbol("shen.x.factorise-defun.findpos")
var symshen_4type_1signature_1of_1include = MakeSymbol("shen.type-signature-of-include")
var symshen_4x_4factorise_1defun_4factorise_1cond = MakeSymbol("shen.x.factorise-defun.factorise-cond")
var symunion = MakeSymbol("union")
var symshen_4semantics = MakeSymbol("shen.semantics")
var symshen_4eval_1cons = MakeSymbol("shen.eval-cons")
var symshen_4aum__to__shen = MakeSymbol("shen.aum_to_shen")
var symshen_4prolog_1aritycheck = MakeSymbol("shen.prolog-aritycheck")
var symshen_4then = MakeSymbol("shen.then")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symtail = MakeSymbol("tail")
var symshen_4type_1signature_1of_1length = MakeSymbol("shen.type-signature-of-length")
var symshen_4hdtl = MakeSymbol("shen.hdtl")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symporters = MakeSymbol("porters")
var symshen_4_5anymulti_6 = MakeSymbol("shen.<anymulti>")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4type_1signature_1of_1adjoin = MakeSymbol("shen.type-signature-of-adjoin")
var symStream = MakeSymbol("Stream")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symerror = MakeSymbol("error")
var symQ = MakeSymbol("Q")
var sym_5_1vector = MakeSymbol("<-vector")
var symshen_4_5end_d_6 = MakeSymbol("shen.<end*>")
var symhd = MakeSymbol("hd")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4list__variables = MakeSymbol("shen.list_variables")
var symand = MakeSymbol("and")
var symshen_4carriage_1return = MakeSymbol("shen.carriage-return")
var symshen_4lookup_1func = MakeSymbol("shen.lookup-func")
var symshen_4_5premise_6 = MakeSymbol("shen.<premise>")
var symshen_4construct_1recursive_1search_1clause = MakeSymbol("shen.construct-recursive-search-clause")
var symOut__1957 = MakeSymbol("Out_1957")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4measure_ereturn = MakeSymbol("shen.measure&return")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symout = MakeSymbol("out")
var symshen_4type_1signature_1of_1hdv = MakeSymbol("shen.type-signature-of-hdv")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var sympackage = MakeSymbol("package")
var symvector_2 = MakeSymbol("vector?")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4x_4features_4cond_1expand = MakeSymbol("shen.x.features.cond-expand")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symopen = MakeSymbol("open")
var symshen_4nextbyte = MakeSymbol("shen.nextbyte")
var symshen_4intprolog_1help = MakeSymbol("shen.intprolog-help")
var symshen_4demod_1rule = MakeSymbol("shen.demod-rule")
var symunify_b = MakeSymbol("unify!")
var symshen_4type_1signature_1of_1version = MakeSymbol("shen.type-signature-of-version")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var sym_h_1 = MakeSymbol(":-")
var symshen_4_5non_1ll_1rules_6 = MakeSymbol("shen.<non-ll-rules>")
var symshen_4type_1signature_1of_1explode = MakeSymbol("shen.type-signature-of-explode")
var symshen_4type_1signature_1of_1internal = MakeSymbol("shen.type-signature-of-internal")
var symshen_4clash_2 = MakeSymbol("shen.clash?")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4right_1_6left = MakeSymbol("shen.right->left")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4type_1signature_1of_1arity = MakeSymbol("shen.type-signature-of-arity")
var symshen_4application__build = MakeSymbol("shen.application_build")
var symshen_4read_7 = MakeSymbol("shen.read+")
var symin = MakeSymbol("in")
var symshen_4nest_1disjunct = MakeSymbol("shen.nest-disjunct")
var symshen_4insert_1prolog_1variables_1help = MakeSymbol("shen.insert-prolog-variables-help")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symtime = MakeSymbol("time")
var symshen_4reduce__help = MakeSymbol("shen.reduce_help")
var symshen_4add__test = MakeSymbol("shen.add_test")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4snd_1or_1fail = MakeSymbol("shen.snd-or-fail")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symremove = MakeSymbol("remove")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var symshen_4head__abstraction = MakeSymbol("shen.head_abstraction")
var symshen_4catch_1cut = MakeSymbol("shen.catch-cut")
var symshen_4tab = MakeSymbol("shen.tab")
var symshen_4compile__to__lambda_7 = MakeSymbol("shen.compile_to_lambda+")
var symshen_4extract__vars = MakeSymbol("shen.extract_vars")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4type_1signature_1of_1_a = MakeSymbol("shen.type-signature-of-=")
var symtc_2 = MakeSymbol("tc?")
var symshen_4singleunderline_2 = MakeSymbol("shen.singleunderline?")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symProcessN = MakeSymbol("ProcessN")
var symshen_4ephemeral__variable_2 = MakeSymbol("shen.ephemeral_variable?")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4type_1signature_1of_1head = MakeSymbol("shen.type-signature-of-head")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4datatype_1error = MakeSymbol("shen.datatype-error")
var symshen_4rule_1_6horn_1clause_1body = MakeSymbol("shen.rule->horn-clause-body")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4aritycheck = MakeSymbol("shen.aritycheck")
var symprotect = MakeSymbol("protect")
var symshen_4valvector = MakeSymbol("shen.valvector")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4t_d_1patterns = MakeSymbol("shen.t*-patterns")
var sym_5_a = MakeSymbol("<=")
var symshen_4type_1signature_1of_1track = MakeSymbol("shen.type-signature-of-track")
var symshen_4toplineread = MakeSymbol("shen.toplineread")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4place = MakeSymbol("shen.place")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4dereferencing = MakeSymbol("shen.dereferencing")
var symshen_4type_1signature_1of_1fst = MakeSymbol("shen.type-signature-of-fst")
var symcd = MakeSymbol("cd")
var symshen_4sh_2 = MakeSymbol("shen.sh?")
var symshen_4construct_1premiss_1literal = MakeSymbol("shen.construct-premiss-literal")
var symshen_4clause__form = MakeSymbol("shen.clause_form")
var symshen_4stack = MakeSymbol("shen.stack")
var symshen_4x_4factorise_1defun_4initialise = MakeSymbol("shen.x.factorise-defun.initialise")
var symcut = MakeSymbol("cut")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4assumetype = MakeSymbol("shen.assumetype")
var symshen_4_5predigits_6 = MakeSymbol("shen.<predigits>")
var symunprofile = MakeSymbol("unprofile")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4exclamation = MakeSymbol("shen.exclamation")
var symshen_4abstract__rule = MakeSymbol("shen.abstract_rule")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var sympreclude = MakeSymbol("preclude")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4type_1signature_1of_1language = MakeSymbol("shen.type-signature-of-language")
var symeval = MakeSymbol("eval")
var symshen_4x_4features_4current = MakeSymbol("shen.x.features.current")
var symshen_4x_4factorise_1defun_4concat_c = MakeSymbol("shen.x.factorise-defun.concat/")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4_dprologvectors_d = MakeSymbol("shen.*prologvectors*")
var symshen_4function_1abstraction_1help = MakeSymbol("shen.function-abstraction-help")
var symshen_4csl_1help = MakeSymbol("shen.csl-help")
var symshen_4split__cc__rule = MakeSymbol("shen.split_cc_rule")
var symclose = MakeSymbol("close")
var symshen_4call__the__continuation = MakeSymbol("shen.call_the_continuation")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symMessage = MakeSymbol("Message")
var symshen_4findpos = MakeSymbol("shen.findpos")
var symC = MakeSymbol("C")
var symshen_4type_1signature_1of_1intersection = MakeSymbol("shen.type-signature-of-intersection")
var symhead = MakeSymbol("head")
var symshen_4update_1symbol_1table = MakeSymbol("shen.update-symbol-table")
var sym_c_4 = MakeSymbol("/.")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4f__error = MakeSymbol("shen.f_error")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4type_1signature_1of_1boolean_2 = MakeSymbol("shen.type-signature-of-boolean?")
var symshen_4x_4factorise_1defun_4generate_1label = MakeSymbol("shen.x.factorise-defun.generate-label")
var sym_dsterror_d = MakeSymbol("*sterror*")
var symshen_4terminal_2 = MakeSymbol("shen.terminal?")
var symshen_4complexity = MakeSymbol("shen.complexity")
var symJ = MakeSymbol("J")
var symshen_4reduce = MakeSymbol("shen.reduce")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4remove__modes = MakeSymbol("shen.remove_modes")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4findallhelp = MakeSymbol("shen.findallhelp")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symshen_4type_1signature_1of_1external = MakeSymbol("shen.type-signature-of-external")
var symshen_4retrieve_1from_1history_1if_1needed = MakeSymbol("shen.retrieve-from-history-if-needed")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4construct_1side_1literals = MakeSymbol("shen.construct-side-literals")
var symshen_4_5body_d_6 = MakeSymbol("shen.<body*>")
var symshen_4procedure__name = MakeSymbol("shen.procedure_name")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var sym_dlanguage_d = MakeSymbol("*language*")
var symlength = MakeSymbol("length")
var symoptimise = MakeSymbol("optimise")
var symshen_4trim_1whitespace = MakeSymbol("shen.trim-whitespace")
var symS = MakeSymbol("S")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4type_1signature_1of_1snd = MakeSymbol("shen.type-signature-of-snd")
var symshen_4type_1signature_1of_1stinput = MakeSymbol("shen.type-signature-of-stinput")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symversion = MakeSymbol("version")
var symshen_4after_1codepoint = MakeSymbol("shen.after-codepoint")
var symshen_4typecheck_1and_1load = MakeSymbol("shen.typecheck-and-load")
var symshen_4type_1signature_1of_1cons_2 = MakeSymbol("shen.type-signature-of-cons?")
var symshen_4type_1signature_1of_1_d = MakeSymbol("shen.type-signature-of-*")
var symshen_4loop = MakeSymbol("shen.loop")
var symunput = MakeSymbol("unput")
var symshen_4_5anysingle_6 = MakeSymbol("shen.<anysingle>")
var symshen_4s_1prolog__literal = MakeSymbol("shen.s-prolog_literal")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symvector = MakeSymbol("vector")
var symshen_4prolog_1error = MakeSymbol("shen.prolog-error")
var symshen_4record_1exceptions = MakeSymbol("shen.record-exceptions")
var symget_1time = MakeSymbol("get-time")
var symsnd = MakeSymbol("snd")
var symshen_4newline = MakeSymbol("shen.newline")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4s_1prolog = MakeSymbol("shen.s-prolog")
var symContext = MakeSymbol("Context")
var symshen_4type_1signature_1of_1variable_2 = MakeSymbol("shen.type-signature-of-variable?")
var symrun = MakeSymbol("run")
var symshen_4type_1signature_1of_1implementation = MakeSymbol("shen.type-signature-of-implementation")
var symtuple_2 = MakeSymbol("tuple?")
var sym_8p = MakeSymbol("@p")
var symgensym = MakeSymbol("gensym")
var symuntrack = MakeSymbol("untrack")
var symspecialise = MakeSymbol("specialise")
var symshen_4construct_1search_1clause = MakeSymbol("shen.construct-search-clause")
var symshen_4type_1signature_1of_1destroy = MakeSymbol("shen.type-signature-of-destroy")
var symshen_4type_1signature_1of_1_a_a = MakeSymbol("shen.type-signature-of-==")
var symshen_4x_4factorise_1defun_4bind_1repeating_1selectors = MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors")
var symshen_4x_4factorise_1defun_4bind_1selector = MakeSymbol("shen.x.factorise-defun.bind-selector")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4aum = MakeSymbol("shen.aum")
var symshen_4ue = MakeSymbol("shen.ue")
var symO = MakeSymbol("O")
var symshen_4safe_1product = MakeSymbol("shen.safe-product")
var symshen_4lzy_a_a = MakeSymbol("shen.lzy==")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4_5sig_7rest_6 = MakeSymbol("shen.<sig+rest>")
var symshen_4x_4features_4cond_1expand_1macro = MakeSymbol("shen.x.features.cond-expand-macro")
var symsterror = MakeSymbol("sterror")
var symshen_4chwild = MakeSymbol("shen.chwild")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symwhen = MakeSymbol("when")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4analyse_1kill = MakeSymbol("shen.analyse-kill")
var symshen_4curry_1synonyms = MakeSymbol("shen.curry-synonyms")
var symshen_4type_1signature_1of_1read_1file_1as_1bytelist = MakeSymbol("shen.type-signature-of-read-file-as-bytelist")
var symstinput = MakeSymbol("stinput")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4grammar__symbol_2 = MakeSymbol("shen.grammar_symbol?")
var symshen_4intprolog = MakeSymbol("shen.intprolog")
var symshen_4case_1form = MakeSymbol("shen.case-form")
var symshen_4default__semantics = MakeSymbol("shen.default_semantics")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var sym_dporters_d = MakeSymbol("*porters*")
var symstring = MakeSymbol("string")
var symshen_4jump__stream_2 = MakeSymbol("shen.jump_stream?")
var symshen_4semantic_1completion_1warning = MakeSymbol("shen.semantic-completion-warning")
var symshen_4copy_1vector_1stage_12 = MakeSymbol("shen.copy-vector-stage-2")
var symshen_4type_1signature_1of_1vector_2 = MakeSymbol("shen.type-signature-of-vector?")
var symshen_4x_4features_4_dfeatures_d = MakeSymbol("shen.x.features.*features*")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4non_1empty = MakeSymbol("shen.non-empty")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4find = MakeSymbol("shen.find")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symshen_4single = MakeSymbol("shen.single")
var symshen_4_5num_6 = MakeSymbol("shen.<num>")
var symshen_4type_1signature_1of_1tail = MakeSymbol("shen.type-signature-of-tail")
var symshen_4x_4factorise_1defun_4free_1variables = MakeSymbol("shen.x.factorise-defun.free-variables")
var symdeclare = MakeSymbol("declare")
var symrelease = MakeSymbol("release")
var symshen_4type_1signature_1of_1profile_1results = MakeSymbol("shen.type-signature-of-profile-results")
var symshen_4_dmaxcomplexity_d = MakeSymbol("shen.*maxcomplexity*")
var symshen_4insert_1lazyderef = MakeSymbol("shen.insert-lazyderef")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4type_1signature_1of_1_5e_6 = MakeSymbol("shen.type-signature-of-<e>")
var symunify = MakeSymbol("unify")
var symidentical = MakeSymbol("identical")
var symshen_4type_1signature_1of_1tlv = MakeSymbol("shen.type-signature-of-tlv")
var symshen_4toplineread__loop = MakeSymbol("shen.toplineread_loop")
var symCase = MakeSymbol("Case")
var symshen_4remove_1synonyms = MakeSymbol("shen.remove-synonyms")
var symshen_4type_1signature_1of_1compile = MakeSymbol("shen.type-signature-of-compile")
var symshen_4type_1signature_1of_1nl = MakeSymbol("shen.type-signature-of-nl")
var symfile = MakeSymbol("file")
var symshen_4mult__subst = MakeSymbol("shen.mult_subst")
var symshen_4t_d_1hyps = MakeSymbol("shen.t*-hyps")
var symshen_4eval_1without_1macros = MakeSymbol("shen.eval-without-macros")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4x_4factorise_1defun_4factorise_1defun = MakeSymbol("shen.x.factorise-defun.factorise-defun")
var symshen_4x_4factorise_1defun_4inline_1mono_1labels = MakeSymbol("shen.x.factorise-defun.inline-mono-labels")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symsave = MakeSymbol("save")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4locally_1bind_1prolog_1vs = MakeSymbol("shen.locally-bind-prolog-vs")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symshen_4type_1signature_1of_1_5_b_6 = MakeSymbol("shen.type-signature-of-<!>")
var symshen_4_dcatch_d = MakeSymbol("shen.*catch*")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4type_1signature_1of_1porters = MakeSymbol("shen.type-signature-of-porters")
var symshen_4tlhd = MakeSymbol("shen.tlhd")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4type_1signature_1of_1symbol_2 = MakeSymbol("shen.type-signature-of-symbol?")
var symtl = MakeSymbol("tl")
var symshen_4find_1past_1inputs = MakeSymbol("shen.find-past-inputs")
var symshen_4double_1_6singles = MakeSymbol("shen.double->singles")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4variant_2 = MakeSymbol("shen.variant?")
var symshen_4type_1signature_1of_1do = MakeSymbol("shen.type-signature-of-do")
var symhdstr = MakeSymbol("hdstr")
var symport = MakeSymbol("port")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4type_1signature_1of_1protect = MakeSymbol("shen.type-signature-of-protect")
var symshen_4hdhd = MakeSymbol("shen.hdhd")
var symshen_4aritycheck_1name = MakeSymbol("shen.aritycheck-name")
var symshen_4_5variable_2_6 = MakeSymbol("shen.<variable?>")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symshen_4type_1signature_1of_1absvector_2 = MakeSymbol("shen.type-signature-of-absvector?")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4kill_1code = MakeSymbol("shen.kill-code")
var symshen_4_5st__input_6 = MakeSymbol("shen.<st_input>")
var symtlstr = MakeSymbol("tlstr")
var symimplementation = MakeSymbol("implementation")
var symshen_4decons_1string = MakeSymbol("shen.decons-string")
var symshen_4type_1signature_1of_1read = MakeSymbol("shen.type-signature-of-read")
var sym_dstinput_d = MakeSymbol("*stinput*")
var sym_dos_d = MakeSymbol("*os*")
var sym_f_fgoto_1label = MakeSymbol("%%goto-label")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5postdigits_6 = MakeSymbol("shen.<postdigits>")
var symshen_4newcontinuation = MakeSymbol("shen.newcontinuation")
var symshen_4type_1signature_1of_1shen_4prhush = MakeSymbol("shen.type-signature-of-shen.prhush")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symwhere = MakeSymbol("where")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4removetype = MakeSymbol("shen.removetype")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symdestroy = MakeSymbol("destroy")
var sym_8v = MakeSymbol("@v")
var symshen_4unix = MakeSymbol("shen.unix")
var symshen_4x_4factorise_1defun_4false_1branch = MakeSymbol("shen.x.factorise-defun.false-branch")
var symhash = MakeSymbol("hash")
var symshen_4split__cc__rules = MakeSymbol("shen.split_cc_rules")
var symread_1file = MakeSymbol("read-file")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4type_1signature_1of_1mapcan = MakeSymbol("shen.type-signature-of-mapcan")
var symX = MakeSymbol("X")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4_5clauses_d_6 = MakeSymbol("shen.<clauses*>")
var symshen_4_5literal_d_6 = MakeSymbol("shen.<literal*>")
var symB = MakeSymbol("B")
var symshen_4_dalphabet_d = MakeSymbol("shen.*alphabet*")
var symContextOut__1957 = MakeSymbol("ContextOut_1957")
var symlist = MakeSymbol("list")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4make_1key = MakeSymbol("shen.make-key")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4cons__form = MakeSymbol("shen.cons_form")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4strip_1protect = MakeSymbol("shen.strip-protect")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4memo = MakeSymbol("shen.memo")
var symshen_4type_1signature_1of_1shen_4proc_1nl = MakeSymbol("shen.type-signature-of-shen.proc-nl")
var symshen_4type_1signature_1of_1sum = MakeSymbol("shen.type-signature-of-sum")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symos = MakeSymbol("os")
var symshen_4externally_1bound = MakeSymbol("shen.externally-bound")
var symor = MakeSymbol("or")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4free__variable__warnings = MakeSymbol("shen.free_variable_warnings")
var symshen_4function_1macro = MakeSymbol("shen.function-macro")
var symshen_4x_4factorise_1defun_4rebranch = MakeSymbol("shen.x.factorise-defun.rebranch")
var symfunction = MakeSymbol("function")
var symshen_4type_1signature_1of_1cn = MakeSymbol("shen.type-signature-of-cn")
var symshen_4type_1signature_1of_1release = MakeSymbol("shen.type-signature-of-release")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var syminput_7 = MakeSymbol("input+")
var symlanguage = MakeSymbol("language")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4type_1signature_1of_1vector = MakeSymbol("shen.type-signature-of-vector")
var symshen_4type_1signature_1of_1optimise = MakeSymbol("shen.type-signature-of-optimise")
var symshen_4type_1signature_1of_1write_1to_1file = MakeSymbol("shen.type-signature-of-write-to-file")
var symoutput = MakeSymbol("output")
var symprolog_2 = MakeSymbol("prolog?")
var sym_j = MakeSymbol("}")
var symlet = MakeSymbol("let")
var symarity = MakeSymbol("arity")
var symexplode = MakeSymbol("explode")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symload = MakeSymbol("load")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symshen_4ebr = MakeSymbol("shen.ebr")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4result = MakeSymbol("shen.result")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symcases = MakeSymbol("cases")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4linearise__help = MakeSymbol("shen.linearise_help")
var symsum = MakeSymbol("sum")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4extract_1pvars = MakeSymbol("shen.extract-pvars")
var symshen_4aritycheck_1action = MakeSymbol("shen.aritycheck-action")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4type_1signature_1of_1n_1_6string = MakeSymbol("shen.type-signature-of-n->string")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symmake_1string = MakeSymbol("make-string")
var sym_a_b = MakeSymbol("=!")
var symshen_4x_4factorise_1defun_4done = MakeSymbol("shen.x.factorise-defun.done")
var symabort = MakeSymbol("abort")
var symshen_4aah = MakeSymbol("shen.aah")
var symshen_4prh = MakeSymbol("shen.prh")
var symshen_4type_1signature_1of_1shen_4app = MakeSymbol("shen.type-signature-of-shen.app")
var symshen_4type_1signature_1of_1read_1file = MakeSymbol("shen.type-signature-of-read-file")
var symshen_4type_1signature_1of_1undefmacro = MakeSymbol("shen.type-signature-of-undefmacro")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4type_1signature_1of_1ps = MakeSymbol("shen.type-signature-of-ps")
var symshen_4x_4features_4add = MakeSymbol("shen.x.features.add")
var symvalue = MakeSymbol("value")
var symshen_4_5comma_1symbol_6 = MakeSymbol("shen.<comma-symbol>")
var symshen_4prolog_1_6shen = MakeSymbol("shen.prolog->shen")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4x_4factorise_1defun_4rebranch_1h = MakeSymbol("shen.x.factorise-defun.rebranch-h")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4type_1signature_1of_1reverse = MakeSymbol("shen.type-signature-of-reverse")
var symappend = MakeSymbol("append")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4err_1condition = MakeSymbol("shen.err-condition")
var symtrack = MakeSymbol("track")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var syminclude = MakeSymbol("include")
var symshen_4list_1stream = MakeSymbol("shen.list-stream")
var sym_e_e = MakeSymbol("&&")
var symshen_4_5E_6 = MakeSymbol("shen.<E>")
var symshen_4_dprocess_1counter_d = MakeSymbol("shen.*process-counter*")
var symshen_4print_1past_1inputs = MakeSymbol("shen.print-past-inputs")
var symsubst = MakeSymbol("subst")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4_5guard_6 = MakeSymbol("shen.<guard>")
var symshen_4cc__body = MakeSymbol("shen.cc_body")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4prolog__constant_2 = MakeSymbol("shen.prolog_constant?")
var symshen_4type_1signature_1of_1kill = MakeSymbol("shen.type-signature-of-kill")
var symshen_4x_4factorise_1defun_4exp_1var = MakeSymbol("shen.x.factorise-defun.exp-var")
var symfail = MakeSymbol("fail")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symfreeze = MakeSymbol("freeze")
var symexternal = MakeSymbol("external")
var symStart = MakeSymbol("Start")
var symshen_4t_d_1def = MakeSymbol("shen.t*-def")
var symParse__Y = MakeSymbol("Parse_Y")
var symdifference = MakeSymbol("difference")
var symshen_4remtype = MakeSymbol("shen.remtype")
var symshen_4type_1signature_1of_1fix = MakeSymbol("shen.type-signature-of-fix")
var symtry_1catch = MakeSymbol("try-catch")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4mode_1ify = MakeSymbol("shen.mode-ify")
var symContext1__1957 = MakeSymbol("Context1_1957")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4failed_b = MakeSymbol("shen.failed!")
var symshen_4type_1signature_1of_1specialise = MakeSymbol("shen.type-signature-of-specialise")
var symshen_4percent = MakeSymbol("shen.percent")
var symintern = MakeSymbol("intern")
var symcall = MakeSymbol("call")
var symshen_4type_1signature_1of_1hash = MakeSymbol("shen.type-signature-of-hash")
var symshen_4type_1signature_1of_1remove = MakeSymbol("shen.type-signature-of-remove")
var symshen_4type_1signature_1of_1union = MakeSymbol("shen.type-signature-of-union")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4type_1signature_1of_1freeze = MakeSymbol("shen.type-signature-of-freeze")
var symdefun = MakeSymbol("defun")
var symshen_4_5st__input2_6 = MakeSymbol("shen.<st_input2>")
var symshen_4pre = MakeSymbol("shen.pre")
var symshen_4f = MakeSymbol("shen.f")
var symRecord = MakeSymbol("Record")
var symshen_4succeeds_2 = MakeSymbol("shen.succeeds?")
var symshen_4add_1macro = MakeSymbol("shen.add-macro")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4call_1help = MakeSymbol("shen.call-help")
var symprofile = MakeSymbol("profile")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symmacroexpand = MakeSymbol("macroexpand")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4type_1signature_1of_1limit = MakeSymbol("shen.type-signature-of-limit")
var symshen_4type_1signature_1of_1string_2 = MakeSymbol("shen.type-signature-of-string?")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symNewStream = MakeSymbol("NewStream")
var symshen_4_5term_d_6 = MakeSymbol("shen.<term*>")
var symdatatype = MakeSymbol("datatype")
var symshen_4type_1signature_1of_1cd = MakeSymbol("shen.type-signature-of-cd")
var symcons_2 = MakeSymbol("cons?")
var symshen_4tests = MakeSymbol("shen.tests")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4recursive__cons__form = MakeSymbol("shen.recursive_cons_form")
var symshen_4_5predicate_d_6 = MakeSymbol("shen.<predicate*>")
var symW = MakeSymbol("W")
var sym_a = MakeSymbol("=")
var symshen_4legitimate_1term_2 = MakeSymbol("shen.legitimate-term?")
var symshen_4packageh = MakeSymbol("shen.packageh")
var symshen_4yacc__cases = MakeSymbol("shen.yacc_cases")
var symsymbol = MakeSymbol("symbol")
var symnl = MakeSymbol("nl")
var symif = MakeSymbol("if")
var symshen_4proc_1input_7 = MakeSymbol("shen.proc-input+")
var syminternal = MakeSymbol("internal")
var symnull = MakeSymbol("null")
var symshen_4lisp_1or = MakeSymbol("shen.lisp-or")
var symNPP = MakeSymbol("NPP")
var symshen_4type_1signature_1of_1vector_1_6 = MakeSymbol("shen.type-signature-of-vector->")
var symshen_4be = MakeSymbol("shen.be")
var symshen_4insert = MakeSymbol("shen.insert")
var symns2_1value = MakeSymbol("ns2-value")
var symshen_4hat = MakeSymbol("shen.hat")
var symshen_4default_1rule = MakeSymbol("shen.default-rule")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var symshen_4type_1signature_1of_1tuple_2 = MakeSymbol("shen.type-signature-of-tuple?")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4type_1signature_1of_1integer_2 = MakeSymbol("shen.type-signature-of-integer?")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symsystemf = MakeSymbol("systemf")
var symshen_4type_1signature_1of_1or = MakeSymbol("shen.type-signature-of-or")
var symshen_4record_1source = MakeSymbol("shen.record-source")
var symshen_4right_1rule = MakeSymbol("shen.right-rule")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var symshen_4placeholder = MakeSymbol("shen.placeholder")
var symU = MakeSymbol("U")
var symshen_4compile__to__kl = MakeSymbol("shen.compile_to_kl")
var symshen_4read_1error = MakeSymbol("shen.read-error")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4type_1signature_1of_1append = MakeSymbol("shen.type-signature-of-append")
var symshen_4type_1signature_1of_1nth = MakeSymbol("shen.type-signature-of-nth")
var symshen_4type_1signature_1of_1port = MakeSymbol("shen.type-signature-of-port")
var symY = MakeSymbol("Y")
var symshen_4repl = MakeSymbol("shen.repl")
var sym_h = MakeSymbol(":")
var symshen_4abstraction__build = MakeSymbol("shen.abstraction_build")
var sym_c = MakeSymbol("/")
var symshen_4resize_1vector = MakeSymbol("shen.resize-vector")
var symshen_4type_1signature_1of_1fail_1if = MakeSymbol("shen.type-signature-of-fail-if")
var symshen_4continuation = MakeSymbol("shen.continuation")
var symshen_4by__hypothesis = MakeSymbol("shen.by_hypothesis")
var symshen_4add_1end = MakeSymbol("shen.add-end")
var symshen_4x_4factorise_1defun_4_dselector_1handlers_d = MakeSymbol("shen.x.factorise-defun.*selector-handlers*")
var symfst = MakeSymbol("fst")
var symintersection = MakeSymbol("intersection")
var symshen_4preclude_1h = MakeSymbol("shen.preclude-h")
var symshen_4mk_1pvar = MakeSymbol("shen.mk-pvar")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4lambda_1form_1entry = MakeSymbol("shen.lambda-form-entry")
var symcons = MakeSymbol("cons")
var sym_d = MakeSymbol("*")
var symbind = MakeSymbol("bind")
var symshen_4numbyte_2 = MakeSymbol("shen.numbyte?")
var symshen_4ue_1h_2 = MakeSymbol("shen.ue-h?")
var symnth = MakeSymbol("nth")
var symshen_4pushnew = MakeSymbol("shen.pushnew")
var symshen_4t_d_1defhh = MakeSymbol("shen.t*-defhh")
var symshen_4type_1signature_1of_1str = MakeSymbol("shen.type-signature-of-str")
var symH = MakeSymbol("H")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4syntax = MakeSymbol("shen.syntax")
var symshen_4read_1file_1as_1Xlist = MakeSymbol("shen.read-file-as-Xlist")
var symTime = MakeSymbol("Time")
var symshen_4type_1signature_1of_1maxinferences = MakeSymbol("shen.type-signature-of-maxinferences")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symParse__ = MakeSymbol("Parse_")
var symshen_4type_1signature_1of_1difference = MakeSymbol("shen.type-signature-of-difference")
var symshen_4_dinstalling_1kl_d = MakeSymbol("shen.*installing-kl*")
var symshen_4em__help = MakeSymbol("shen.em_help")
var sym_5_1 = MakeSymbol("<-")
var symshen_4linearise__X = MakeSymbol("shen.linearise_X")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4type_1signature_1of_1pr = MakeSymbol("shen.type-signature-of-pr")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4construct_1search_1clauses = MakeSymbol("shen.construct-search-clauses")
var symshen_4type_1signature_1of_1write_1byte = MakeSymbol("shen.type-signature-of-write-byte")
var symZ = MakeSymbol("Z")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symshen_4type_1signature_1of_1os = MakeSymbol("shen.type-signature-of-os")
var symshen_4type_1signature_1of_1tc_2 = MakeSymbol("shen.type-signature-of-tc?")
var symA = MakeSymbol("A")
var symshen_4reverse__help = MakeSymbol("shen.reverse_help")
var symmode = MakeSymbol("mode")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symfindall = MakeSymbol("findall")
var symshen_4x_4factorise_1defun_4register_1selector_1handler = MakeSymbol("shen.x.factorise-defun.register-selector-handler")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symsuccess = MakeSymbol("success")
var symprint = MakeSymbol("print")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4skip = MakeSymbol("shen.skip")
var symps = MakeSymbol("ps")
var symkill = MakeSymbol("kill")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symshen_4_5semicolon_1symbol_6 = MakeSymbol("shen.<semicolon-symbol>")
var symshen_4line = MakeSymbol("shen.line")
