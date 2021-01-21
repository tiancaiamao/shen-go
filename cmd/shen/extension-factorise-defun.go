package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionFactoriseDefunMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2012-2019 Bruno Deferrari.  All rights reserved.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

	tmp68429 := MakeNative(func(__e *ControlFlow) {
		V4931 := __e.Get(1)
		_ = V4931
		tmp68540 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68541 := Call(__e, tmp68540, V4931)

		var ifres68472 Obj

		if True == tmp68541 {
			tmp68536 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68537 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68538 := Call(__e, tmp68537, V4931)

			tmp68539 := Call(__e, tmp68536, symdefun, tmp68538)

			var ifres68474 Obj

			if True == tmp68539 {
				tmp68532 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68533 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68534 := Call(__e, tmp68533, V4931)

				tmp68535 := Call(__e, tmp68532, tmp68534)

				var ifres68476 Obj

				if True == tmp68535 {
					tmp68526 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68527 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68528 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68529 := Call(__e, tmp68528, V4931)

					tmp68530 := Call(__e, tmp68527, tmp68529)

					tmp68531 := Call(__e, tmp68526, tmp68530)

					var ifres68478 Obj

					if True == tmp68531 {
						tmp68518 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp68519 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68520 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68521 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68522 := Call(__e, tmp68521, V4931)

						tmp68523 := Call(__e, tmp68520, tmp68522)

						tmp68524 := Call(__e, tmp68519, tmp68523)

						tmp68525 := Call(__e, tmp68518, tmp68524)

						var ifres68480 Obj

						if True == tmp68525 {
							tmp68508 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp68509 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp68510 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68511 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68512 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68513 := Call(__e, tmp68512, V4931)

							tmp68514 := Call(__e, tmp68511, tmp68513)

							tmp68515 := Call(__e, tmp68510, tmp68514)

							tmp68516 := Call(__e, tmp68509, tmp68515)

							tmp68517 := Call(__e, tmp68508, tmp68516)

							var ifres68482 Obj

							if True == tmp68517 {
								tmp68496 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp68497 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp68498 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp68499 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp68500 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp68501 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp68502 := Call(__e, tmp68501, V4931)

								tmp68503 := Call(__e, tmp68500, tmp68502)

								tmp68504 := Call(__e, tmp68499, tmp68503)

								tmp68505 := Call(__e, tmp68498, tmp68504)

								tmp68506 := Call(__e, tmp68497, tmp68505)

								tmp68507 := Call(__e, tmp68496, symcond, tmp68506)

								var ifres68484 Obj

								if True == tmp68507 {
									tmp68486 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp68487 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68488 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68489 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68490 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68491 := Call(__e, tmp68490, V4931)

									tmp68492 := Call(__e, tmp68489, tmp68491)

									tmp68493 := Call(__e, tmp68488, tmp68492)

									tmp68494 := Call(__e, tmp68487, tmp68493)

									tmp68495 := Call(__e, tmp68486, Nil, tmp68494)

									var ifres68485 Obj

									if True == tmp68495 {
										ifres68485 = True

									} else {
										ifres68485 = False

									}

									ifres68484 = ifres68485

								} else {
									ifres68484 = False

								}

								var ifres68483 Obj

								if True == ifres68484 {
									ifres68483 = True

								} else {
									ifres68483 = False

								}

								ifres68482 = ifres68483

							} else {
								ifres68482 = False

							}

							var ifres68481 Obj

							if True == ifres68482 {
								ifres68481 = True

							} else {
								ifres68481 = False

							}

							ifres68480 = ifres68481

						} else {
							ifres68480 = False

						}

						var ifres68479 Obj

						if True == ifres68480 {
							ifres68479 = True

						} else {
							ifres68479 = False

						}

						ifres68478 = ifres68479

					} else {
						ifres68478 = False

					}

					var ifres68477 Obj

					if True == ifres68478 {
						ifres68477 = True

					} else {
						ifres68477 = False

					}

					ifres68476 = ifres68477

				} else {
					ifres68476 = False

				}

				var ifres68475 Obj

				if True == ifres68476 {
					ifres68475 = True

				} else {
					ifres68475 = False

				}

				ifres68474 = ifres68475

			} else {
				ifres68474 = False

			}

			var ifres68473 Obj

			if True == ifres68474 {
				ifres68473 = True

			} else {
				ifres68473 = False

			}

			ifres68472 = ifres68473

		} else {
			ifres68472 = False

		}

		if True == ifres68472 {
			tmp68431 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68432 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68433 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68434 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68435 := Call(__e, tmp68434, V4931)

			tmp68436 := Call(__e, tmp68433, tmp68435)

			tmp68437 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68438 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68439 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68440 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68441 := Call(__e, tmp68440, V4931)

			tmp68442 := Call(__e, tmp68439, tmp68441)

			tmp68443 := Call(__e, tmp68438, tmp68442)

			tmp68444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68445 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4factorise_1cond)

			tmp68446 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68448 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68449 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68450 := Call(__e, tmp68449, V4931)

			tmp68451 := Call(__e, tmp68448, tmp68450)

			tmp68452 := Call(__e, tmp68447, tmp68451)

			tmp68453 := Call(__e, tmp68446, tmp68452)

			tmp68454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68455 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68456 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68457 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68458 := Call(__e, tmp68457, V4931)

			tmp68459 := Call(__e, tmp68456, tmp68458)

			tmp68460 := Call(__e, tmp68455, tmp68459, Nil)

			tmp68461 := Call(__e, tmp68454, symshen_4f__error, tmp68460)

			tmp68462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68464 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68465 := Call(__e, tmp68464, V4931)

			tmp68466 := Call(__e, tmp68463, tmp68465)

			tmp68467 := Call(__e, tmp68462, tmp68466)

			tmp68468 := Call(__e, tmp68445, tmp68453, tmp68461, tmp68467)

			tmp68469 := Call(__e, tmp68444, tmp68468, Nil)

			tmp68470 := Call(__e, tmp68437, tmp68443, tmp68469)

			tmp68471 := Call(__e, tmp68432, tmp68436, tmp68470)

			__e.TailApply(tmp68431, symdefun, tmp68471)
			return

		} else {
			__e.Return(V4931)
			return
		}

	}, 1)

	tmp68542 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4factorise_1defun, tmp68429)

	_ = tmp68542

	tmp68543 := MakeNative(func(__e *ControlFlow) {
		V4943 := __e.Get(1)
		_ = V4943
		V4944 := __e.Get(2)
		_ = V4944
		V4945 := __e.Get(3)
		_ = V4945
		tmp68558 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68559 := Call(__e, tmp68558, V4943)

		var ifres68552 Obj

		if True == tmp68559 {
			tmp68554 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68555 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68556 := Call(__e, tmp68555, V4943)

			tmp68557 := Call(__e, tmp68554, symcond, tmp68556)

			var ifres68553 Obj

			if True == tmp68557 {
				ifres68553 = True

			} else {
				ifres68553 = False

			}

			ifres68552 = ifres68553

		} else {
			ifres68552 = False

		}

		if True == ifres68552 {
			tmp68545 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

			tmp68546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch)

			tmp68547 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4add_1returns)

			tmp68548 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68549 := Call(__e, tmp68548, V4943)

			tmp68550 := Call(__e, tmp68547, tmp68549)

			tmp68551 := Call(__e, tmp68546, tmp68550, V4944)

			__e.TailApply(tmp68545, tmp68551, V4945)
			return

		} else {
			__e.Return(V4943)
			return
		}

	}, 3)

	tmp68560 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4factorise_1cond, tmp68543)

	_ = tmp68560

	tmp68561 := MakeNative(func(__e *ControlFlow) {
		V4947 := __e.Get(1)
		_ = V4947
		tmp68610 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68611 := Call(__e, tmp68610, Nil, V4947)

		if True == tmp68611 {
			__e.Return(Nil)
			return
		} else {
			tmp68608 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp68609 := Call(__e, tmp68608, V4947)

			var ifres68584 Obj

			if True == tmp68609 {
				tmp68604 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68605 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68606 := Call(__e, tmp68605, V4947)

				tmp68607 := Call(__e, tmp68604, tmp68606)

				var ifres68586 Obj

				if True == tmp68607 {
					tmp68598 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68599 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68600 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68601 := Call(__e, tmp68600, V4947)

					tmp68602 := Call(__e, tmp68599, tmp68601)

					tmp68603 := Call(__e, tmp68598, tmp68602)

					var ifres68588 Obj

					if True == tmp68603 {
						tmp68590 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp68591 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68592 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68593 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68594 := Call(__e, tmp68593, V4947)

						tmp68595 := Call(__e, tmp68592, tmp68594)

						tmp68596 := Call(__e, tmp68591, tmp68595)

						tmp68597 := Call(__e, tmp68590, Nil, tmp68596)

						var ifres68589 Obj

						if True == tmp68597 {
							ifres68589 = True

						} else {
							ifres68589 = False

						}

						ifres68588 = ifres68589

					} else {
						ifres68588 = False

					}

					var ifres68587 Obj

					if True == ifres68588 {
						ifres68587 = True

					} else {
						ifres68587 = False

					}

					ifres68586 = ifres68587

				} else {
					ifres68586 = False

				}

				var ifres68585 Obj

				if True == ifres68586 {
					ifres68585 = True

				} else {
					ifres68585 = False

				}

				ifres68584 = ifres68585

			} else {
				ifres68584 = False

			}

			if True == ifres68584 {
				tmp68565 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68566 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68567 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68568 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68569 := Call(__e, tmp68568, V4947)

				tmp68570 := Call(__e, tmp68567, tmp68569)

				tmp68571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68572 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68573 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68575 := Call(__e, tmp68574, V4947)

				tmp68576 := Call(__e, tmp68573, tmp68575)

				tmp68577 := Call(__e, tmp68572, sym_f_freturn, tmp68576)

				tmp68578 := Call(__e, tmp68571, tmp68577, Nil)

				tmp68579 := Call(__e, tmp68566, tmp68570, tmp68578)

				tmp68580 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4add_1returns)

				tmp68581 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68582 := Call(__e, tmp68581, V4947)

				tmp68583 := Call(__e, tmp68580, tmp68582)

				__e.TailApply(tmp68565, tmp68579, tmp68583)
				return

			} else {
				tmp68564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp68564, symshen_4x_4factorise_1defun_4add_1returns)
				return

			}

		}

	}, 1)

	tmp68612 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4add_1returns, tmp68561)

	_ = tmp68612

	tmp68613 := MakeNative(func(__e *ControlFlow) {
		tmp68614 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		__e.TailApply(tmp68614, sym_f_flabel)
		return

	}, 0)

	tmp68615 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4generate_1label, tmp68613)

	_ = tmp68615

	tmp68616 := MakeNative(func(__e *ControlFlow) {
		V4950 := __e.Get(1)
		_ = V4950
		V4951 := __e.Get(2)
		_ = V4951
		tmp68617 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

		tmp68618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

		tmp68619 := Call(__e, tmp68618, V4950, V4951, Nil)

		__e.TailApply(tmp68617, tmp68619)
		return

	}, 2)

	tmp68620 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4free_1variables, tmp68616)

	_ = tmp68620

	tmp68621 := MakeNative(func(__e *ControlFlow) {
		V4963 := __e.Get(1)
		_ = V4963
		V4964 := __e.Get(2)
		_ = V4964
		V4965 := __e.Get(3)
		_ = V4965
		tmp68748 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68749 := Call(__e, tmp68748, V4963)

		var ifres68706 Obj

		if True == tmp68749 {
			tmp68744 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68745 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68746 := Call(__e, tmp68745, V4963)

			tmp68747 := Call(__e, tmp68744, symlet, tmp68746)

			var ifres68708 Obj

			if True == tmp68747 {
				tmp68740 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68741 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68742 := Call(__e, tmp68741, V4963)

				tmp68743 := Call(__e, tmp68740, tmp68742)

				var ifres68710 Obj

				if True == tmp68743 {
					tmp68734 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68735 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68736 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68737 := Call(__e, tmp68736, V4963)

					tmp68738 := Call(__e, tmp68735, tmp68737)

					tmp68739 := Call(__e, tmp68734, tmp68738)

					var ifres68712 Obj

					if True == tmp68739 {
						tmp68726 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp68727 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68728 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68730 := Call(__e, tmp68729, V4963)

						tmp68731 := Call(__e, tmp68728, tmp68730)

						tmp68732 := Call(__e, tmp68727, tmp68731)

						tmp68733 := Call(__e, tmp68726, tmp68732)

						var ifres68714 Obj

						if True == tmp68733 {
							tmp68716 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp68717 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68718 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68719 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68721 := Call(__e, tmp68720, V4963)

							tmp68722 := Call(__e, tmp68719, tmp68721)

							tmp68723 := Call(__e, tmp68718, tmp68722)

							tmp68724 := Call(__e, tmp68717, tmp68723)

							tmp68725 := Call(__e, tmp68716, Nil, tmp68724)

							var ifres68715 Obj

							if True == tmp68725 {
								ifres68715 = True

							} else {
								ifres68715 = False

							}

							ifres68714 = ifres68715

						} else {
							ifres68714 = False

						}

						var ifres68713 Obj

						if True == ifres68714 {
							ifres68713 = True

						} else {
							ifres68713 = False

						}

						ifres68712 = ifres68713

					} else {
						ifres68712 = False

					}

					var ifres68711 Obj

					if True == ifres68712 {
						ifres68711 = True

					} else {
						ifres68711 = False

					}

					ifres68710 = ifres68711

				} else {
					ifres68710 = False

				}

				var ifres68709 Obj

				if True == ifres68710 {
					ifres68709 = True

				} else {
					ifres68709 = False

				}

				ifres68708 = ifres68709

			} else {
				ifres68708 = False

			}

			var ifres68707 Obj

			if True == ifres68708 {
				ifres68707 = True

			} else {
				ifres68707 = False

			}

			ifres68706 = ifres68707

		} else {
			ifres68706 = False

		}

		if True == ifres68706 {
			tmp68683 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

			tmp68684 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68685 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68686 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68688 := Call(__e, tmp68687, V4963)

			tmp68689 := Call(__e, tmp68686, tmp68688)

			tmp68690 := Call(__e, tmp68685, tmp68689)

			tmp68691 := Call(__e, tmp68684, tmp68690)

			tmp68692 := Call(__e, PrimNS1Value(symns2_1value), symremove)

			tmp68693 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68694 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68695 := Call(__e, tmp68694, V4963)

			tmp68696 := Call(__e, tmp68693, tmp68695)

			tmp68697 := Call(__e, tmp68692, tmp68696, V4964)

			tmp68698 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

			tmp68699 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68701 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68702 := Call(__e, tmp68701, V4963)

			tmp68703 := Call(__e, tmp68700, tmp68702)

			tmp68704 := Call(__e, tmp68699, tmp68703)

			tmp68705 := Call(__e, tmp68698, tmp68704, V4964, V4965)

			__e.TailApply(tmp68683, tmp68691, tmp68697, tmp68705)
			return

		} else {
			tmp68681 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp68682 := Call(__e, tmp68681, V4963)

			var ifres68651 Obj

			if True == tmp68682 {
				tmp68677 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp68678 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68679 := Call(__e, tmp68678, V4963)

				tmp68680 := Call(__e, tmp68677, symlambda, tmp68679)

				var ifres68653 Obj

				if True == tmp68680 {
					tmp68673 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68674 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68675 := Call(__e, tmp68674, V4963)

					tmp68676 := Call(__e, tmp68673, tmp68675)

					var ifres68655 Obj

					if True == tmp68676 {
						tmp68667 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp68668 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68669 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68670 := Call(__e, tmp68669, V4963)

						tmp68671 := Call(__e, tmp68668, tmp68670)

						tmp68672 := Call(__e, tmp68667, tmp68671)

						var ifres68657 Obj

						if True == tmp68672 {
							tmp68659 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp68660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68661 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68662 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68663 := Call(__e, tmp68662, V4963)

							tmp68664 := Call(__e, tmp68661, tmp68663)

							tmp68665 := Call(__e, tmp68660, tmp68664)

							tmp68666 := Call(__e, tmp68659, Nil, tmp68665)

							var ifres68658 Obj

							if True == tmp68666 {
								ifres68658 = True

							} else {
								ifres68658 = False

							}

							ifres68657 = ifres68658

						} else {
							ifres68657 = False

						}

						var ifres68656 Obj

						if True == ifres68657 {
							ifres68656 = True

						} else {
							ifres68656 = False

						}

						ifres68655 = ifres68656

					} else {
						ifres68655 = False

					}

					var ifres68654 Obj

					if True == ifres68655 {
						ifres68654 = True

					} else {
						ifres68654 = False

					}

					ifres68653 = ifres68654

				} else {
					ifres68653 = False

				}

				var ifres68652 Obj

				if True == ifres68653 {
					ifres68652 = True

				} else {
					ifres68652 = False

				}

				ifres68651 = ifres68652

			} else {
				ifres68651 = False

			}

			if True == ifres68651 {
				tmp68638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

				tmp68639 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68640 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68641 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68642 := Call(__e, tmp68641, V4963)

				tmp68643 := Call(__e, tmp68640, tmp68642)

				tmp68644 := Call(__e, tmp68639, tmp68643)

				tmp68645 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				tmp68646 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68648 := Call(__e, tmp68647, V4963)

				tmp68649 := Call(__e, tmp68646, tmp68648)

				tmp68650 := Call(__e, tmp68645, tmp68649, V4964)

				__e.TailApply(tmp68638, tmp68644, tmp68650, V4965)
				return

			} else {
				tmp68636 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68637 := Call(__e, tmp68636, V4963)

				if True == tmp68637 {
					tmp68629 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

					tmp68630 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68631 := Call(__e, tmp68630, V4963)

					tmp68632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables_1h)

					tmp68633 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68634 := Call(__e, tmp68633, V4963)

					tmp68635 := Call(__e, tmp68632, tmp68634, V4964, V4965)

					__e.TailApply(tmp68629, tmp68631, V4964, tmp68635)
					return

				} else {
					tmp68627 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					tmp68628 := Call(__e, tmp68627, V4963, V4964)

					if True == tmp68628 {
						tmp68626 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

						__e.TailApply(tmp68626, V4963, V4965)
						return

					} else {
						__e.Return(V4965)
						return
					}

				}

			}

		}

	}, 3)

	tmp68750 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4free_1variables_1h, tmp68621)

	_ = tmp68750

	tmp68751 := MakeNative(func(__e *ControlFlow) {
		V4968 := __e.Get(1)
		_ = V4968
		V4969 := __e.Get(2)
		_ = V4969
		tmp68864 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68865 := Call(__e, tmp68864, V4968)

		var ifres68822 Obj

		if True == tmp68865 {
			tmp68860 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68861 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68862 := Call(__e, tmp68861, V4968)

			tmp68863 := Call(__e, tmp68860, sym_f_flet_1label, tmp68862)

			var ifres68824 Obj

			if True == tmp68863 {
				tmp68856 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68857 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68858 := Call(__e, tmp68857, V4968)

				tmp68859 := Call(__e, tmp68856, tmp68858)

				var ifres68826 Obj

				if True == tmp68859 {
					tmp68850 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68851 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68853 := Call(__e, tmp68852, V4968)

					tmp68854 := Call(__e, tmp68851, tmp68853)

					tmp68855 := Call(__e, tmp68850, tmp68854)

					var ifres68828 Obj

					if True == tmp68855 {
						tmp68842 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp68843 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68846 := Call(__e, tmp68845, V4968)

						tmp68847 := Call(__e, tmp68844, tmp68846)

						tmp68848 := Call(__e, tmp68843, tmp68847)

						tmp68849 := Call(__e, tmp68842, tmp68848)

						var ifres68830 Obj

						if True == tmp68849 {
							tmp68832 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp68833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68836 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68837 := Call(__e, tmp68836, V4968)

							tmp68838 := Call(__e, tmp68835, tmp68837)

							tmp68839 := Call(__e, tmp68834, tmp68838)

							tmp68840 := Call(__e, tmp68833, tmp68839)

							tmp68841 := Call(__e, tmp68832, Nil, tmp68840)

							var ifres68831 Obj

							if True == tmp68841 {
								ifres68831 = True

							} else {
								ifres68831 = False

							}

							ifres68830 = ifres68831

						} else {
							ifres68830 = False

						}

						var ifres68829 Obj

						if True == ifres68830 {
							ifres68829 = True

						} else {
							ifres68829 = False

						}

						ifres68828 = ifres68829

					} else {
						ifres68828 = False

					}

					var ifres68827 Obj

					if True == ifres68828 {
						ifres68827 = True

					} else {
						ifres68827 = False

					}

					ifres68826 = ifres68827

				} else {
					ifres68826 = False

				}

				var ifres68825 Obj

				if True == ifres68826 {
					ifres68825 = True

				} else {
					ifres68825 = False

				}

				ifres68824 = ifres68825

			} else {
				ifres68824 = False

			}

			var ifres68823 Obj

			if True == ifres68824 {
				ifres68823 = True

			} else {
				ifres68823 = False

			}

			ifres68822 = ifres68823

		} else {
			ifres68822 = False

		}

		if True == ifres68822 {
			tmp68754 := MakeNative(func(__e *ControlFlow) {
				FreeVars := __e.Get(1)
				_ = FreeVars
				tmp68755 := MakeNative(func(__e *ControlFlow) {
					NewBody := __e.Get(1)
					_ = NewBody
					tmp68756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68758 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68759 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68760 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68761 := Call(__e, tmp68760, V4968)

					tmp68762 := Call(__e, tmp68759, tmp68761)

					tmp68763 := Call(__e, tmp68758, tmp68762, FreeVars)

					tmp68764 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68765 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68766 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68767 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68768 := Call(__e, tmp68767, V4968)

					tmp68769 := Call(__e, tmp68766, tmp68768)

					tmp68770 := Call(__e, tmp68765, tmp68769)

					tmp68771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68772 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

					tmp68773 := Call(__e, tmp68772, NewBody, V4969)

					tmp68774 := Call(__e, tmp68771, tmp68773, Nil)

					tmp68775 := Call(__e, tmp68764, tmp68770, tmp68774)

					tmp68776 := Call(__e, tmp68757, tmp68763, tmp68775)

					__e.TailApply(tmp68756, sym_f_flet_1label, tmp68776)
					return

				}, 1)

				tmp68812 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp68813 := Call(__e, tmp68812, Nil, FreeVars)

				var ifres68777 Obj

				if True == tmp68813 {
					tmp68804 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68805 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68806 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68807 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68808 := Call(__e, tmp68807, V4968)

					tmp68809 := Call(__e, tmp68806, tmp68808)

					tmp68810 := Call(__e, tmp68805, tmp68809)

					tmp68811 := Call(__e, tmp68804, tmp68810)

					ifres68777 = tmp68811

				} else {
					tmp68778 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp68779 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68780 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68781 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68782 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68783 := Call(__e, tmp68782, V4968)

					tmp68784 := Call(__e, tmp68781, tmp68783)

					tmp68785 := Call(__e, tmp68780, tmp68784, FreeVars)

					tmp68786 := Call(__e, tmp68779, sym_f_fgoto_1label, tmp68785)

					tmp68787 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68788 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68789 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68790 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68791 := Call(__e, tmp68790, V4968)

					tmp68792 := Call(__e, tmp68789, tmp68791)

					tmp68793 := Call(__e, tmp68788, tmp68792, Nil)

					tmp68794 := Call(__e, tmp68787, sym_f_fgoto_1label, tmp68793)

					tmp68795 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68799 := Call(__e, tmp68798, V4968)

					tmp68800 := Call(__e, tmp68797, tmp68799)

					tmp68801 := Call(__e, tmp68796, tmp68800)

					tmp68802 := Call(__e, tmp68795, tmp68801)

					tmp68803 := Call(__e, tmp68778, tmp68786, tmp68794, tmp68802)

					ifres68777 = tmp68803

				}

				__e.TailApply(tmp68755, ifres68777)
				return

			}, 1)

			tmp68814 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4free_1variables)

			tmp68815 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68816 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68817 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68818 := Call(__e, tmp68817, V4968)

			tmp68819 := Call(__e, tmp68816, tmp68818)

			tmp68820 := Call(__e, tmp68815, tmp68819)

			tmp68821 := Call(__e, tmp68814, tmp68820, V4969)

			__e.TailApply(tmp68754, tmp68821)
			return

		} else {
			tmp68753 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp68753, symshen_4x_4factorise_1defun_4attach_1free_1variables)
			return

		}

	}, 2)

	tmp68866 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4attach_1free_1variables, tmp68751)

	_ = tmp68866

	tmp68867 := MakeNative(func(__e *ControlFlow) {
		V4976 := __e.Get(1)
		_ = V4976
		V4977 := __e.Get(2)
		_ = V4977
		tmp69182 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp69183 := Call(__e, tmp69182, V4976)

		var ifres69118 Obj

		if True == tmp69183 {
			tmp69178 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp69179 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69180 := Call(__e, tmp69179, V4976)

			tmp69181 := Call(__e, tmp69178, sym_f_flet_1label, tmp69180)

			var ifres69120 Obj

			if True == tmp69181 {
				tmp69174 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69175 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69176 := Call(__e, tmp69175, V4976)

				tmp69177 := Call(__e, tmp69174, tmp69176)

				var ifres69122 Obj

				if True == tmp69177 {
					tmp69168 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69171 := Call(__e, tmp69170, V4976)

					tmp69172 := Call(__e, tmp69169, tmp69171)

					tmp69173 := Call(__e, tmp69168, tmp69172)

					var ifres69124 Obj

					if True == tmp69173 {
						tmp69160 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69161 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69162 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69163 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69164 := Call(__e, tmp69163, V4976)

						tmp69165 := Call(__e, tmp69162, tmp69164)

						tmp69166 := Call(__e, tmp69161, tmp69165)

						tmp69167 := Call(__e, tmp69160, tmp69166)

						var ifres69126 Obj

						if True == tmp69167 {
							tmp69150 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69151 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69152 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69153 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69154 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69155 := Call(__e, tmp69154, V4976)

							tmp69156 := Call(__e, tmp69153, tmp69155)

							tmp69157 := Call(__e, tmp69152, tmp69156)

							tmp69158 := Call(__e, tmp69151, tmp69157)

							tmp69159 := Call(__e, tmp69150, Nil, tmp69158)

							var ifres69128 Obj

							if True == tmp69159 {
								tmp69130 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

								tmp69131 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

								tmp69132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp69133 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp69134 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69135 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69136 := Call(__e, tmp69135, V4976)

								tmp69137 := Call(__e, tmp69134, tmp69136)

								tmp69138 := Call(__e, tmp69133, tmp69137, Nil)

								tmp69139 := Call(__e, tmp69132, sym_f_fgoto_1label, tmp69138)

								tmp69140 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69141 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69143 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69144 := Call(__e, tmp69143, V4976)

								tmp69145 := Call(__e, tmp69142, tmp69144)

								tmp69146 := Call(__e, tmp69141, tmp69145)

								tmp69147 := Call(__e, tmp69140, tmp69146)

								tmp69148 := Call(__e, tmp69131, tmp69139, tmp69147)

								tmp69149 := Call(__e, tmp69130, tmp69148, MakeNumber(1))

								var ifres69129 Obj

								if True == tmp69149 {
									ifres69129 = True

								} else {
									ifres69129 = False

								}

								ifres69128 = ifres69129

							} else {
								ifres69128 = False

							}

							var ifres69127 Obj

							if True == ifres69128 {
								ifres69127 = True

							} else {
								ifres69127 = False

							}

							ifres69126 = ifres69127

						} else {
							ifres69126 = False

						}

						var ifres69125 Obj

						if True == ifres69126 {
							ifres69125 = True

						} else {
							ifres69125 = False

						}

						ifres69124 = ifres69125

					} else {
						ifres69124 = False

					}

					var ifres69123 Obj

					if True == ifres69124 {
						ifres69123 = True

					} else {
						ifres69123 = False

					}

					ifres69122 = ifres69123

				} else {
					ifres69122 = False

				}

				var ifres69121 Obj

				if True == ifres69122 {
					ifres69121 = True

				} else {
					ifres69121 = False

				}

				ifres69120 = ifres69121

			} else {
				ifres69120 = False

			}

			var ifres69119 Obj

			if True == ifres69120 {
				ifres69119 = True

			} else {
				ifres69119 = False

			}

			ifres69118 = ifres69119

		} else {
			ifres69118 = False

		}

		if True == ifres69118 {
			tmp69093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4attach_1free_1variables)

			tmp69094 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69095 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69096 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69097 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69098 := Call(__e, tmp69097, V4976)

			tmp69099 := Call(__e, tmp69096, tmp69098)

			tmp69100 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

			tmp69102 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69104 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69105 := Call(__e, tmp69104, V4976)

			tmp69106 := Call(__e, tmp69103, tmp69105)

			tmp69107 := Call(__e, tmp69102, tmp69106)

			tmp69108 := Call(__e, tmp69101, tmp69107, V4977)

			tmp69109 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69110 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69111 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69112 := Call(__e, tmp69111, V4976)

			tmp69113 := Call(__e, tmp69110, tmp69112)

			tmp69114 := Call(__e, tmp69109, tmp69113)

			tmp69115 := Call(__e, tmp69100, tmp69108, tmp69114)

			tmp69116 := Call(__e, tmp69095, tmp69099, tmp69115)

			tmp69117 := Call(__e, tmp69094, sym_f_flet_1label, tmp69116)

			__e.TailApply(tmp69093, tmp69117, V4977)
			return

		} else {
			tmp69091 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69092 := Call(__e, tmp69091, V4976)

			var ifres69049 Obj

			if True == tmp69092 {
				tmp69087 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp69088 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69089 := Call(__e, tmp69088, V4976)

				tmp69090 := Call(__e, tmp69087, sym_f_flet_1label, tmp69089)

				var ifres69051 Obj

				if True == tmp69090 {
					tmp69083 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69084 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69085 := Call(__e, tmp69084, V4976)

					tmp69086 := Call(__e, tmp69083, tmp69085)

					var ifres69053 Obj

					if True == tmp69086 {
						tmp69077 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69078 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69079 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69080 := Call(__e, tmp69079, V4976)

						tmp69081 := Call(__e, tmp69078, tmp69080)

						tmp69082 := Call(__e, tmp69077, tmp69081)

						var ifres69055 Obj

						if True == tmp69082 {
							tmp69069 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp69070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69072 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69073 := Call(__e, tmp69072, V4976)

							tmp69074 := Call(__e, tmp69071, tmp69073)

							tmp69075 := Call(__e, tmp69070, tmp69074)

							tmp69076 := Call(__e, tmp69069, tmp69075)

							var ifres69057 Obj

							if True == tmp69076 {
								tmp69059 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp69060 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69061 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69062 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69063 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69064 := Call(__e, tmp69063, V4976)

								tmp69065 := Call(__e, tmp69062, tmp69064)

								tmp69066 := Call(__e, tmp69061, tmp69065)

								tmp69067 := Call(__e, tmp69060, tmp69066)

								tmp69068 := Call(__e, tmp69059, Nil, tmp69067)

								var ifres69058 Obj

								if True == tmp69068 {
									ifres69058 = True

								} else {
									ifres69058 = False

								}

								ifres69057 = ifres69058

							} else {
								ifres69057 = False

							}

							var ifres69056 Obj

							if True == ifres69057 {
								ifres69056 = True

							} else {
								ifres69056 = False

							}

							ifres69055 = ifres69056

						} else {
							ifres69055 = False

						}

						var ifres69054 Obj

						if True == ifres69055 {
							ifres69054 = True

						} else {
							ifres69054 = False

						}

						ifres69053 = ifres69054

					} else {
						ifres69053 = False

					}

					var ifres69052 Obj

					if True == ifres69053 {
						ifres69052 = True

					} else {
						ifres69052 = False

					}

					ifres69051 = ifres69052

				} else {
					ifres69051 = False

				}

				var ifres69050 Obj

				if True == ifres69051 {
					ifres69050 = True

				} else {
					ifres69050 = False

				}

				ifres69049 = ifres69050

			} else {
				ifres69049 = False

			}

			if True == ifres69049 {
				tmp69022 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

				tmp69023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

				tmp69024 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69025 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69026 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69027 := Call(__e, tmp69026, V4976)

				tmp69028 := Call(__e, tmp69025, tmp69027)

				tmp69029 := Call(__e, tmp69024, tmp69028)

				tmp69030 := Call(__e, tmp69023, tmp69029, V4977)

				tmp69031 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69032 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69033 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69035 := Call(__e, tmp69034, V4976)

				tmp69036 := Call(__e, tmp69033, tmp69035)

				tmp69037 := Call(__e, tmp69032, tmp69036, Nil)

				tmp69038 := Call(__e, tmp69031, sym_f_fgoto_1label, tmp69037)

				tmp69039 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

				tmp69040 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69041 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69043 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69044 := Call(__e, tmp69043, V4976)

				tmp69045 := Call(__e, tmp69042, tmp69044)

				tmp69046 := Call(__e, tmp69041, tmp69045)

				tmp69047 := Call(__e, tmp69040, tmp69046)

				tmp69048 := Call(__e, tmp69039, tmp69047, V4977)

				__e.TailApply(tmp69022, tmp69030, tmp69038, tmp69048)
				return

			} else {
				tmp69020 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69021 := Call(__e, tmp69020, V4976)

				var ifres68978 Obj

				if True == tmp69021 {
					tmp69016 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69017 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69018 := Call(__e, tmp69017, V4976)

					tmp69019 := Call(__e, tmp69016, symif, tmp69018)

					var ifres68980 Obj

					if True == tmp69019 {
						tmp69012 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69013 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69014 := Call(__e, tmp69013, V4976)

						tmp69015 := Call(__e, tmp69012, tmp69014)

						var ifres68982 Obj

						if True == tmp69015 {
							tmp69006 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp69007 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69008 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69009 := Call(__e, tmp69008, V4976)

							tmp69010 := Call(__e, tmp69007, tmp69009)

							tmp69011 := Call(__e, tmp69006, tmp69010)

							var ifres68984 Obj

							if True == tmp69011 {
								tmp68998 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp68999 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69000 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69001 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69002 := Call(__e, tmp69001, V4976)

								tmp69003 := Call(__e, tmp69000, tmp69002)

								tmp69004 := Call(__e, tmp68999, tmp69003)

								tmp69005 := Call(__e, tmp68998, tmp69004)

								var ifres68986 Obj

								if True == tmp69005 {
									tmp68988 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp68989 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68990 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68991 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68992 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68993 := Call(__e, tmp68992, V4976)

									tmp68994 := Call(__e, tmp68991, tmp68993)

									tmp68995 := Call(__e, tmp68990, tmp68994)

									tmp68996 := Call(__e, tmp68989, tmp68995)

									tmp68997 := Call(__e, tmp68988, Nil, tmp68996)

									var ifres68987 Obj

									if True == tmp68997 {
										ifres68987 = True

									} else {
										ifres68987 = False

									}

									ifres68986 = ifres68987

								} else {
									ifres68986 = False

								}

								var ifres68985 Obj

								if True == ifres68986 {
									ifres68985 = True

								} else {
									ifres68985 = False

								}

								ifres68984 = ifres68985

							} else {
								ifres68984 = False

							}

							var ifres68983 Obj

							if True == ifres68984 {
								ifres68983 = True

							} else {
								ifres68983 = False

							}

							ifres68982 = ifres68983

						} else {
							ifres68982 = False

						}

						var ifres68981 Obj

						if True == ifres68982 {
							ifres68981 = True

						} else {
							ifres68981 = False

						}

						ifres68980 = ifres68981

					} else {
						ifres68980 = False

					}

					var ifres68979 Obj

					if True == ifres68980 {
						ifres68979 = True

					} else {
						ifres68979 = False

					}

					ifres68978 = ifres68979

				} else {
					ifres68978 = False

				}

				if True == ifres68978 {
					tmp68949 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68950 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68951 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68952 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68953 := Call(__e, tmp68952, V4976)

					tmp68954 := Call(__e, tmp68951, tmp68953)

					tmp68955 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

					tmp68957 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68960 := Call(__e, tmp68959, V4976)

					tmp68961 := Call(__e, tmp68958, tmp68960)

					tmp68962 := Call(__e, tmp68957, tmp68961)

					tmp68963 := Call(__e, tmp68956, tmp68962, V4977)

					tmp68964 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp68965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

					tmp68966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68967 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68968 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68969 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68970 := Call(__e, tmp68969, V4976)

					tmp68971 := Call(__e, tmp68968, tmp68970)

					tmp68972 := Call(__e, tmp68967, tmp68971)

					tmp68973 := Call(__e, tmp68966, tmp68972)

					tmp68974 := Call(__e, tmp68965, tmp68973, V4977)

					tmp68975 := Call(__e, tmp68964, tmp68974, Nil)

					tmp68976 := Call(__e, tmp68955, tmp68963, tmp68975)

					tmp68977 := Call(__e, tmp68950, tmp68954, tmp68976)

					__e.TailApply(tmp68949, symif, tmp68977)
					return

				} else {
					tmp68947 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68948 := Call(__e, tmp68947, V4976)

					var ifres68905 Obj

					if True == tmp68948 {
						tmp68943 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp68944 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68945 := Call(__e, tmp68944, V4976)

						tmp68946 := Call(__e, tmp68943, symlet, tmp68945)

						var ifres68907 Obj

						if True == tmp68946 {
							tmp68939 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp68940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68941 := Call(__e, tmp68940, V4976)

							tmp68942 := Call(__e, tmp68939, tmp68941)

							var ifres68909 Obj

							if True == tmp68942 {
								tmp68933 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp68934 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp68935 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp68936 := Call(__e, tmp68935, V4976)

								tmp68937 := Call(__e, tmp68934, tmp68936)

								tmp68938 := Call(__e, tmp68933, tmp68937)

								var ifres68911 Obj

								if True == tmp68938 {
									tmp68925 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp68926 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68927 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68929 := Call(__e, tmp68928, V4976)

									tmp68930 := Call(__e, tmp68927, tmp68929)

									tmp68931 := Call(__e, tmp68926, tmp68930)

									tmp68932 := Call(__e, tmp68925, tmp68931)

									var ifres68913 Obj

									if True == tmp68932 {
										tmp68915 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp68916 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp68917 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp68918 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp68919 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp68920 := Call(__e, tmp68919, V4976)

										tmp68921 := Call(__e, tmp68918, tmp68920)

										tmp68922 := Call(__e, tmp68917, tmp68921)

										tmp68923 := Call(__e, tmp68916, tmp68922)

										tmp68924 := Call(__e, tmp68915, Nil, tmp68923)

										var ifres68914 Obj

										if True == tmp68924 {
											ifres68914 = True

										} else {
											ifres68914 = False

										}

										ifres68913 = ifres68914

									} else {
										ifres68913 = False

									}

									var ifres68912 Obj

									if True == ifres68913 {
										ifres68912 = True

									} else {
										ifres68912 = False

									}

									ifres68911 = ifres68912

								} else {
									ifres68911 = False

								}

								var ifres68910 Obj

								if True == ifres68911 {
									ifres68910 = True

								} else {
									ifres68910 = False

								}

								ifres68909 = ifres68910

							} else {
								ifres68909 = False

							}

							var ifres68908 Obj

							if True == ifres68909 {
								ifres68908 = True

							} else {
								ifres68908 = False

							}

							ifres68907 = ifres68908

						} else {
							ifres68907 = False

						}

						var ifres68906 Obj

						if True == ifres68907 {
							ifres68906 = True

						} else {
							ifres68906 = False

						}

						ifres68905 = ifres68906

					} else {
						ifres68905 = False

					}

					if True == ifres68905 {
						tmp68872 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp68873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp68874 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68875 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68876 := Call(__e, tmp68875, V4976)

						tmp68877 := Call(__e, tmp68874, tmp68876)

						tmp68878 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp68879 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68880 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68881 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68882 := Call(__e, tmp68881, V4976)

						tmp68883 := Call(__e, tmp68880, tmp68882)

						tmp68884 := Call(__e, tmp68879, tmp68883)

						tmp68885 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp68886 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4inline_1mono_1labels)

						tmp68887 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68888 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68889 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68891 := Call(__e, tmp68890, V4976)

						tmp68892 := Call(__e, tmp68889, tmp68891)

						tmp68893 := Call(__e, tmp68888, tmp68892)

						tmp68894 := Call(__e, tmp68887, tmp68893)

						tmp68895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp68896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68897 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp68898 := Call(__e, tmp68897, V4976)

						tmp68899 := Call(__e, tmp68896, tmp68898)

						tmp68900 := Call(__e, tmp68895, tmp68899, V4977)

						tmp68901 := Call(__e, tmp68886, tmp68894, tmp68900)

						tmp68902 := Call(__e, tmp68885, tmp68901, Nil)

						tmp68903 := Call(__e, tmp68878, tmp68884, tmp68902)

						tmp68904 := Call(__e, tmp68873, tmp68877, tmp68903)

						__e.TailApply(tmp68872, symlet, tmp68904)
						return

					} else {
						__e.Return(V4976)
						return
					}

				}

			}

		}

	}, 2)

	tmp69184 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4inline_1mono_1labels, tmp68867)

	_ = tmp69184

	tmp69185 := MakeNative(func(__e *ControlFlow) {
		V4984 := __e.Get(1)
		_ = V4984
		V4985 := __e.Get(2)
		_ = V4985
		tmp69386 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp69387 := Call(__e, tmp69386, Nil, V4984)

		if True == tmp69387 {
			__e.Return(V4985)
			return
		} else {
			tmp69384 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69385 := Call(__e, tmp69384, V4984)

			var ifres69352 Obj

			if True == tmp69385 {
				tmp69380 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69381 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69382 := Call(__e, tmp69381, V4984)

				tmp69383 := Call(__e, tmp69380, tmp69382)

				var ifres69354 Obj

				if True == tmp69383 {
					tmp69374 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69375 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69376 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69377 := Call(__e, tmp69376, V4984)

					tmp69378 := Call(__e, tmp69375, tmp69377)

					tmp69379 := Call(__e, tmp69374, True, tmp69378)

					var ifres69356 Obj

					if True == tmp69379 {
						tmp69368 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69369 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69370 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69371 := Call(__e, tmp69370, V4984)

						tmp69372 := Call(__e, tmp69369, tmp69371)

						tmp69373 := Call(__e, tmp69368, tmp69372)

						var ifres69358 Obj

						if True == tmp69373 {
							tmp69360 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69361 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69362 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69363 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69364 := Call(__e, tmp69363, V4984)

							tmp69365 := Call(__e, tmp69362, tmp69364)

							tmp69366 := Call(__e, tmp69361, tmp69365)

							tmp69367 := Call(__e, tmp69360, Nil, tmp69366)

							var ifres69359 Obj

							if True == tmp69367 {
								ifres69359 = True

							} else {
								ifres69359 = False

							}

							ifres69358 = ifres69359

						} else {
							ifres69358 = False

						}

						var ifres69357 Obj

						if True == ifres69358 {
							ifres69357 = True

						} else {
							ifres69357 = False

						}

						ifres69356 = ifres69357

					} else {
						ifres69356 = False

					}

					var ifres69355 Obj

					if True == ifres69356 {
						ifres69355 = True

					} else {
						ifres69355 = False

					}

					ifres69354 = ifres69355

				} else {
					ifres69354 = False

				}

				var ifres69353 Obj

				if True == ifres69354 {
					ifres69353 = True

				} else {
					ifres69353 = False

				}

				ifres69352 = ifres69353

			} else {
				ifres69352 = False

			}

			if True == ifres69352 {
				tmp69347 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69348 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69349 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69350 := Call(__e, tmp69349, V4984)

				tmp69351 := Call(__e, tmp69348, tmp69350)

				__e.TailApply(tmp69347, tmp69351)
				return

			} else {
				tmp69345 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69346 := Call(__e, tmp69345, V4984)

				var ifres69267 Obj

				if True == tmp69346 {
					tmp69341 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69342 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69343 := Call(__e, tmp69342, V4984)

					tmp69344 := Call(__e, tmp69341, tmp69343)

					var ifres69269 Obj

					if True == tmp69344 {
						tmp69335 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69336 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69337 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69338 := Call(__e, tmp69337, V4984)

						tmp69339 := Call(__e, tmp69336, tmp69338)

						tmp69340 := Call(__e, tmp69335, tmp69339)

						var ifres69271 Obj

						if True == tmp69340 {
							tmp69327 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69328 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69329 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69330 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69331 := Call(__e, tmp69330, V4984)

							tmp69332 := Call(__e, tmp69329, tmp69331)

							tmp69333 := Call(__e, tmp69328, tmp69332)

							tmp69334 := Call(__e, tmp69327, symand, tmp69333)

							var ifres69273 Obj

							if True == tmp69334 {
								tmp69319 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp69320 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69321 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69322 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69323 := Call(__e, tmp69322, V4984)

								tmp69324 := Call(__e, tmp69321, tmp69323)

								tmp69325 := Call(__e, tmp69320, tmp69324)

								tmp69326 := Call(__e, tmp69319, tmp69325)

								var ifres69275 Obj

								if True == tmp69326 {
									tmp69309 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp69310 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69311 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69312 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp69313 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp69314 := Call(__e, tmp69313, V4984)

									tmp69315 := Call(__e, tmp69312, tmp69314)

									tmp69316 := Call(__e, tmp69311, tmp69315)

									tmp69317 := Call(__e, tmp69310, tmp69316)

									tmp69318 := Call(__e, tmp69309, tmp69317)

									var ifres69277 Obj

									if True == tmp69318 {
										tmp69297 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp69298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69299 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69300 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69301 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp69302 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp69303 := Call(__e, tmp69302, V4984)

										tmp69304 := Call(__e, tmp69301, tmp69303)

										tmp69305 := Call(__e, tmp69300, tmp69304)

										tmp69306 := Call(__e, tmp69299, tmp69305)

										tmp69307 := Call(__e, tmp69298, tmp69306)

										tmp69308 := Call(__e, tmp69297, Nil, tmp69307)

										var ifres69279 Obj

										if True == tmp69308 {
											tmp69291 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp69292 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69293 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69294 := Call(__e, tmp69293, V4984)

											tmp69295 := Call(__e, tmp69292, tmp69294)

											tmp69296 := Call(__e, tmp69291, tmp69295)

											var ifres69281 Obj

											if True == tmp69296 {
												tmp69283 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp69284 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp69285 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp69286 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp69287 := Call(__e, tmp69286, V4984)

												tmp69288 := Call(__e, tmp69285, tmp69287)

												tmp69289 := Call(__e, tmp69284, tmp69288)

												tmp69290 := Call(__e, tmp69283, Nil, tmp69289)

												var ifres69282 Obj

												if True == tmp69290 {
													ifres69282 = True

												} else {
													ifres69282 = False

												}

												ifres69281 = ifres69282

											} else {
												ifres69281 = False

											}

											var ifres69280 Obj

											if True == ifres69281 {
												ifres69280 = True

											} else {
												ifres69280 = False

											}

											ifres69279 = ifres69280

										} else {
											ifres69279 = False

										}

										var ifres69278 Obj

										if True == ifres69279 {
											ifres69278 = True

										} else {
											ifres69278 = False

										}

										ifres69277 = ifres69278

									} else {
										ifres69277 = False

									}

									var ifres69276 Obj

									if True == ifres69277 {
										ifres69276 = True

									} else {
										ifres69276 = False

									}

									ifres69275 = ifres69276

								} else {
									ifres69275 = False

								}

								var ifres69274 Obj

								if True == ifres69275 {
									ifres69274 = True

								} else {
									ifres69274 = False

								}

								ifres69273 = ifres69274

							} else {
								ifres69273 = False

							}

							var ifres69272 Obj

							if True == ifres69273 {
								ifres69272 = True

							} else {
								ifres69272 = False

							}

							ifres69271 = ifres69272

						} else {
							ifres69271 = False

						}

						var ifres69270 Obj

						if True == ifres69271 {
							ifres69270 = True

						} else {
							ifres69270 = False

						}

						ifres69269 = ifres69270

					} else {
						ifres69269 = False

					}

					var ifres69268 Obj

					if True == ifres69269 {
						ifres69268 = True

					} else {
						ifres69268 = False

					}

					ifres69267 = ifres69268

				} else {
					ifres69267 = False

				}

				if True == ifres69267 {
					tmp69236 := MakeNative(func(__e *ControlFlow) {
						TrueBranch := __e.Get(1)
						_ = TrueBranch
						tmp69237 := MakeNative(func(__e *ControlFlow) {
							FalseBranch := __e.Get(1)
							_ = FalseBranch
							tmp69238 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch_1h)

							tmp69239 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69241 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69242 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69243 := Call(__e, tmp69242, V4984)

							tmp69244 := Call(__e, tmp69241, tmp69243)

							tmp69245 := Call(__e, tmp69240, tmp69244)

							tmp69246 := Call(__e, tmp69239, tmp69245)

							__e.TailApply(tmp69238, tmp69246, TrueBranch, FalseBranch, V4985)
							return

						}, 1)

						tmp69247 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

						tmp69248 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69250 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69251 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69252 := Call(__e, tmp69251, V4984)

						tmp69253 := Call(__e, tmp69250, tmp69252)

						tmp69254 := Call(__e, tmp69249, tmp69253)

						tmp69255 := Call(__e, tmp69248, tmp69254)

						tmp69256 := Call(__e, tmp69247, tmp69255, V4984)

						__e.TailApply(tmp69237, tmp69256)
						return

					}, 1)

					tmp69257 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4true_1branch)

					tmp69258 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69259 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69260 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69261 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69262 := Call(__e, tmp69261, V4984)

					tmp69263 := Call(__e, tmp69260, tmp69262)

					tmp69264 := Call(__e, tmp69259, tmp69263)

					tmp69265 := Call(__e, tmp69258, tmp69264)

					tmp69266 := Call(__e, tmp69257, tmp69265, V4984)

					__e.TailApply(tmp69236, tmp69266)
					return

				} else {
					tmp69234 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69235 := Call(__e, tmp69234, V4984)

					var ifres69210 Obj

					if True == tmp69235 {
						tmp69230 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69231 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69232 := Call(__e, tmp69231, V4984)

						tmp69233 := Call(__e, tmp69230, tmp69232)

						var ifres69212 Obj

						if True == tmp69233 {
							tmp69224 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp69225 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69226 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69227 := Call(__e, tmp69226, V4984)

							tmp69228 := Call(__e, tmp69225, tmp69227)

							tmp69229 := Call(__e, tmp69224, tmp69228)

							var ifres69214 Obj

							if True == tmp69229 {
								tmp69216 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp69217 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69218 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69219 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69220 := Call(__e, tmp69219, V4984)

								tmp69221 := Call(__e, tmp69218, tmp69220)

								tmp69222 := Call(__e, tmp69217, tmp69221)

								tmp69223 := Call(__e, tmp69216, Nil, tmp69222)

								var ifres69215 Obj

								if True == tmp69223 {
									ifres69215 = True

								} else {
									ifres69215 = False

								}

								ifres69214 = ifres69215

							} else {
								ifres69214 = False

							}

							var ifres69213 Obj

							if True == ifres69214 {
								ifres69213 = True

							} else {
								ifres69213 = False

							}

							ifres69212 = ifres69213

						} else {
							ifres69212 = False

						}

						var ifres69211 Obj

						if True == ifres69212 {
							ifres69211 = True

						} else {
							ifres69211 = False

						}

						ifres69210 = ifres69211

					} else {
						ifres69210 = False

					}

					if True == ifres69210 {
						tmp69191 := MakeNative(func(__e *ControlFlow) {
							TrueBranch := __e.Get(1)
							_ = TrueBranch
							tmp69192 := MakeNative(func(__e *ControlFlow) {
								FalseBranch := __e.Get(1)
								_ = FalseBranch
								tmp69193 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch_1h)

								tmp69194 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69195 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69196 := Call(__e, tmp69195, V4984)

								tmp69197 := Call(__e, tmp69194, tmp69196)

								__e.TailApply(tmp69193, tmp69197, TrueBranch, FalseBranch, V4985)
								return

							}, 1)

							tmp69198 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

							tmp69199 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69200 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69201 := Call(__e, tmp69200, V4984)

							tmp69202 := Call(__e, tmp69199, tmp69201)

							tmp69203 := Call(__e, tmp69198, tmp69202, V4984)

							__e.TailApply(tmp69192, tmp69203)
							return

						}, 1)

						tmp69204 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4true_1branch)

						tmp69205 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69206 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69207 := Call(__e, tmp69206, V4984)

						tmp69208 := Call(__e, tmp69205, tmp69207)

						tmp69209 := Call(__e, tmp69204, tmp69208, V4984)

						__e.TailApply(tmp69191, tmp69209)
						return

					} else {
						tmp69190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp69190, symshen_4x_4factorise_1defun_4rebranch)
						return

					}

				}

			}

		}

	}, 2)

	tmp69388 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4rebranch, tmp69185)

	_ = tmp69388

	tmp69389 := MakeNative(func(__e *ControlFlow) {
		V4990 := __e.Get(1)
		_ = V4990
		V4991 := __e.Get(2)
		_ = V4991
		V4992 := __e.Get(3)
		_ = V4992
		V4993 := __e.Get(4)
		_ = V4993
		tmp69390 := MakeNative(func(__e *ControlFlow) {
			NewElse := __e.Get(1)
			_ = NewElse
			tmp69391 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4with_1labelled_1else)

			tmp69392 := MakeNative(func(__e *ControlFlow) {
				GotoElse := __e.Get(1)
				_ = GotoElse
				tmp69393 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs)

				tmp69394 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69396 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69397 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4optimize_1selectors)

				tmp69398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch)

				tmp69399 := Call(__e, tmp69398, V4991, GotoElse)

				tmp69400 := Call(__e, tmp69397, V4990, tmp69399)

				tmp69401 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69402 := Call(__e, tmp69401, GotoElse, Nil)

				tmp69403 := Call(__e, tmp69396, tmp69400, tmp69402)

				tmp69404 := Call(__e, tmp69395, V4990, tmp69403)

				tmp69405 := Call(__e, tmp69394, symif, tmp69404)

				__e.TailApply(tmp69393, tmp69405)
				return

			}, 1)

			__e.TailApply(tmp69391, NewElse, tmp69392)
			return

		}, 1)

		tmp69406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4rebranch)

		tmp69407 := Call(__e, tmp69406, V4992, V4993)

		__e.TailApply(tmp69390, tmp69407)
		return

	}, 4)

	tmp69408 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4rebranch_1h, tmp69389)

	_ = tmp69408

	tmp69409 := MakeNative(func(__e *ControlFlow) {
		V5006 := __e.Get(1)
		_ = V5006
		V5007 := __e.Get(2)
		_ = V5007
		tmp69572 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp69573 := Call(__e, tmp69572, V5007)

		var ifres69482 Obj

		if True == tmp69573 {
			tmp69568 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69569 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69570 := Call(__e, tmp69569, V5007)

			tmp69571 := Call(__e, tmp69568, tmp69570)

			var ifres69484 Obj

			if True == tmp69571 {
				tmp69562 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69563 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69564 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69565 := Call(__e, tmp69564, V5007)

				tmp69566 := Call(__e, tmp69563, tmp69565)

				tmp69567 := Call(__e, tmp69562, tmp69566)

				var ifres69486 Obj

				if True == tmp69567 {
					tmp69554 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69555 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69556 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69557 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69558 := Call(__e, tmp69557, V5007)

					tmp69559 := Call(__e, tmp69556, tmp69558)

					tmp69560 := Call(__e, tmp69555, tmp69559)

					tmp69561 := Call(__e, tmp69554, symand, tmp69560)

					var ifres69488 Obj

					if True == tmp69561 {
						tmp69546 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69547 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69548 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69549 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69550 := Call(__e, tmp69549, V5007)

						tmp69551 := Call(__e, tmp69548, tmp69550)

						tmp69552 := Call(__e, tmp69547, tmp69551)

						tmp69553 := Call(__e, tmp69546, tmp69552)

						var ifres69490 Obj

						if True == tmp69553 {
							tmp69536 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp69537 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69538 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69539 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69540 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69541 := Call(__e, tmp69540, V5007)

							tmp69542 := Call(__e, tmp69539, tmp69541)

							tmp69543 := Call(__e, tmp69538, tmp69542)

							tmp69544 := Call(__e, tmp69537, tmp69543)

							tmp69545 := Call(__e, tmp69536, tmp69544)

							var ifres69492 Obj

							if True == tmp69545 {
								tmp69524 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp69525 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69526 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69527 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69528 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69529 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69530 := Call(__e, tmp69529, V5007)

								tmp69531 := Call(__e, tmp69528, tmp69530)

								tmp69532 := Call(__e, tmp69527, tmp69531)

								tmp69533 := Call(__e, tmp69526, tmp69532)

								tmp69534 := Call(__e, tmp69525, tmp69533)

								tmp69535 := Call(__e, tmp69524, Nil, tmp69534)

								var ifres69494 Obj

								if True == tmp69535 {
									tmp69518 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp69519 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69520 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp69521 := Call(__e, tmp69520, V5007)

									tmp69522 := Call(__e, tmp69519, tmp69521)

									tmp69523 := Call(__e, tmp69518, tmp69522)

									var ifres69496 Obj

									if True == tmp69523 {
										tmp69510 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp69511 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69512 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69513 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp69514 := Call(__e, tmp69513, V5007)

										tmp69515 := Call(__e, tmp69512, tmp69514)

										tmp69516 := Call(__e, tmp69511, tmp69515)

										tmp69517 := Call(__e, tmp69510, Nil, tmp69516)

										var ifres69498 Obj

										if True == tmp69517 {
											tmp69500 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp69501 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69502 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69503 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69504 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69505 := Call(__e, tmp69504, V5007)

											tmp69506 := Call(__e, tmp69503, tmp69505)

											tmp69507 := Call(__e, tmp69502, tmp69506)

											tmp69508 := Call(__e, tmp69501, tmp69507)

											tmp69509 := Call(__e, tmp69500, tmp69508, V5006)

											var ifres69499 Obj

											if True == tmp69509 {
												ifres69499 = True

											} else {
												ifres69499 = False

											}

											ifres69498 = ifres69499

										} else {
											ifres69498 = False

										}

										var ifres69497 Obj

										if True == ifres69498 {
											ifres69497 = True

										} else {
											ifres69497 = False

										}

										ifres69496 = ifres69497

									} else {
										ifres69496 = False

									}

									var ifres69495 Obj

									if True == ifres69496 {
										ifres69495 = True

									} else {
										ifres69495 = False

									}

									ifres69494 = ifres69495

								} else {
									ifres69494 = False

								}

								var ifres69493 Obj

								if True == ifres69494 {
									ifres69493 = True

								} else {
									ifres69493 = False

								}

								ifres69492 = ifres69493

							} else {
								ifres69492 = False

							}

							var ifres69491 Obj

							if True == ifres69492 {
								ifres69491 = True

							} else {
								ifres69491 = False

							}

							ifres69490 = ifres69491

						} else {
							ifres69490 = False

						}

						var ifres69489 Obj

						if True == ifres69490 {
							ifres69489 = True

						} else {
							ifres69489 = False

						}

						ifres69488 = ifres69489

					} else {
						ifres69488 = False

					}

					var ifres69487 Obj

					if True == ifres69488 {
						ifres69487 = True

					} else {
						ifres69487 = False

					}

					ifres69486 = ifres69487

				} else {
					ifres69486 = False

				}

				var ifres69485 Obj

				if True == ifres69486 {
					ifres69485 = True

				} else {
					ifres69485 = False

				}

				ifres69484 = ifres69485

			} else {
				ifres69484 = False

			}

			var ifres69483 Obj

			if True == ifres69484 {
				ifres69483 = True

			} else {
				ifres69483 = False

			}

			ifres69482 = ifres69483

		} else {
			ifres69482 = False

		}

		if True == ifres69482 {
			tmp69453 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69455 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69457 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69458 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69459 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69460 := Call(__e, tmp69459, V5007)

			tmp69461 := Call(__e, tmp69458, tmp69460)

			tmp69462 := Call(__e, tmp69457, tmp69461)

			tmp69463 := Call(__e, tmp69456, tmp69462)

			tmp69464 := Call(__e, tmp69455, tmp69463)

			tmp69465 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69466 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69467 := Call(__e, tmp69466, V5007)

			tmp69468 := Call(__e, tmp69465, tmp69467)

			tmp69469 := Call(__e, tmp69454, tmp69464, tmp69468)

			tmp69470 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4true_1branch)

			tmp69471 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69472 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69473 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69474 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69475 := Call(__e, tmp69474, V5007)

			tmp69476 := Call(__e, tmp69473, tmp69475)

			tmp69477 := Call(__e, tmp69472, tmp69476)

			tmp69478 := Call(__e, tmp69471, tmp69477)

			tmp69479 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69480 := Call(__e, tmp69479, V5007)

			tmp69481 := Call(__e, tmp69470, tmp69478, tmp69480)

			__e.TailApply(tmp69453, tmp69469, tmp69481)
			return

		} else {
			tmp69451 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69452 := Call(__e, tmp69451, V5007)

			var ifres69419 Obj

			if True == tmp69452 {
				tmp69447 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69448 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69449 := Call(__e, tmp69448, V5007)

				tmp69450 := Call(__e, tmp69447, tmp69449)

				var ifres69421 Obj

				if True == tmp69450 {
					tmp69441 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69442 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69443 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69444 := Call(__e, tmp69443, V5007)

					tmp69445 := Call(__e, tmp69442, tmp69444)

					tmp69446 := Call(__e, tmp69441, tmp69445)

					var ifres69423 Obj

					if True == tmp69446 {
						tmp69433 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp69434 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69436 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69437 := Call(__e, tmp69436, V5007)

						tmp69438 := Call(__e, tmp69435, tmp69437)

						tmp69439 := Call(__e, tmp69434, tmp69438)

						tmp69440 := Call(__e, tmp69433, Nil, tmp69439)

						var ifres69425 Obj

						if True == tmp69440 {
							tmp69427 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69428 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69429 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69430 := Call(__e, tmp69429, V5007)

							tmp69431 := Call(__e, tmp69428, tmp69430)

							tmp69432 := Call(__e, tmp69427, tmp69431, V5006)

							var ifres69426 Obj

							if True == tmp69432 {
								ifres69426 = True

							} else {
								ifres69426 = False

							}

							ifres69425 = ifres69426

						} else {
							ifres69425 = False

						}

						var ifres69424 Obj

						if True == ifres69425 {
							ifres69424 = True

						} else {
							ifres69424 = False

						}

						ifres69423 = ifres69424

					} else {
						ifres69423 = False

					}

					var ifres69422 Obj

					if True == ifres69423 {
						ifres69422 = True

					} else {
						ifres69422 = False

					}

					ifres69421 = ifres69422

				} else {
					ifres69421 = False

				}

				var ifres69420 Obj

				if True == ifres69421 {
					ifres69420 = True

				} else {
					ifres69420 = False

				}

				ifres69419 = ifres69420

			} else {
				ifres69419 = False

			}

			if True == ifres69419 {
				tmp69412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69413 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp69414 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69415 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69416 := Call(__e, tmp69415, V5007)

				tmp69417 := Call(__e, tmp69414, tmp69416)

				tmp69418 := Call(__e, tmp69413, True, tmp69417)

				__e.TailApply(tmp69412, tmp69418, Nil)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 2)

	tmp69574 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4true_1branch, tmp69409)

	_ = tmp69574

	tmp69575 := MakeNative(func(__e *ControlFlow) {
		V5016 := __e.Get(1)
		_ = V5016
		V5017 := __e.Get(2)
		_ = V5017
		tmp69720 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp69721 := Call(__e, tmp69720, V5017)

		var ifres69630 Obj

		if True == tmp69721 {
			tmp69716 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69717 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69718 := Call(__e, tmp69717, V5017)

			tmp69719 := Call(__e, tmp69716, tmp69718)

			var ifres69632 Obj

			if True == tmp69719 {
				tmp69710 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69711 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69712 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69713 := Call(__e, tmp69712, V5017)

				tmp69714 := Call(__e, tmp69711, tmp69713)

				tmp69715 := Call(__e, tmp69710, tmp69714)

				var ifres69634 Obj

				if True == tmp69715 {
					tmp69702 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69703 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69704 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69705 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69706 := Call(__e, tmp69705, V5017)

					tmp69707 := Call(__e, tmp69704, tmp69706)

					tmp69708 := Call(__e, tmp69703, tmp69707)

					tmp69709 := Call(__e, tmp69702, symand, tmp69708)

					var ifres69636 Obj

					if True == tmp69709 {
						tmp69694 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69695 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69696 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69697 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69698 := Call(__e, tmp69697, V5017)

						tmp69699 := Call(__e, tmp69696, tmp69698)

						tmp69700 := Call(__e, tmp69695, tmp69699)

						tmp69701 := Call(__e, tmp69694, tmp69700)

						var ifres69638 Obj

						if True == tmp69701 {
							tmp69684 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp69685 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69686 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69687 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69688 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69689 := Call(__e, tmp69688, V5017)

							tmp69690 := Call(__e, tmp69687, tmp69689)

							tmp69691 := Call(__e, tmp69686, tmp69690)

							tmp69692 := Call(__e, tmp69685, tmp69691)

							tmp69693 := Call(__e, tmp69684, tmp69692)

							var ifres69640 Obj

							if True == tmp69693 {
								tmp69672 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp69673 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69674 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69675 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69676 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69677 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69678 := Call(__e, tmp69677, V5017)

								tmp69679 := Call(__e, tmp69676, tmp69678)

								tmp69680 := Call(__e, tmp69675, tmp69679)

								tmp69681 := Call(__e, tmp69674, tmp69680)

								tmp69682 := Call(__e, tmp69673, tmp69681)

								tmp69683 := Call(__e, tmp69672, Nil, tmp69682)

								var ifres69642 Obj

								if True == tmp69683 {
									tmp69666 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp69667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69668 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp69669 := Call(__e, tmp69668, V5017)

									tmp69670 := Call(__e, tmp69667, tmp69669)

									tmp69671 := Call(__e, tmp69666, tmp69670)

									var ifres69644 Obj

									if True == tmp69671 {
										tmp69658 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp69659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69661 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp69662 := Call(__e, tmp69661, V5017)

										tmp69663 := Call(__e, tmp69660, tmp69662)

										tmp69664 := Call(__e, tmp69659, tmp69663)

										tmp69665 := Call(__e, tmp69658, Nil, tmp69664)

										var ifres69646 Obj

										if True == tmp69665 {
											tmp69648 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp69649 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69651 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69652 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69653 := Call(__e, tmp69652, V5017)

											tmp69654 := Call(__e, tmp69651, tmp69653)

											tmp69655 := Call(__e, tmp69650, tmp69654)

											tmp69656 := Call(__e, tmp69649, tmp69655)

											tmp69657 := Call(__e, tmp69648, tmp69656, V5016)

											var ifres69647 Obj

											if True == tmp69657 {
												ifres69647 = True

											} else {
												ifres69647 = False

											}

											ifres69646 = ifres69647

										} else {
											ifres69646 = False

										}

										var ifres69645 Obj

										if True == ifres69646 {
											ifres69645 = True

										} else {
											ifres69645 = False

										}

										ifres69644 = ifres69645

									} else {
										ifres69644 = False

									}

									var ifres69643 Obj

									if True == ifres69644 {
										ifres69643 = True

									} else {
										ifres69643 = False

									}

									ifres69642 = ifres69643

								} else {
									ifres69642 = False

								}

								var ifres69641 Obj

								if True == ifres69642 {
									ifres69641 = True

								} else {
									ifres69641 = False

								}

								ifres69640 = ifres69641

							} else {
								ifres69640 = False

							}

							var ifres69639 Obj

							if True == ifres69640 {
								ifres69639 = True

							} else {
								ifres69639 = False

							}

							ifres69638 = ifres69639

						} else {
							ifres69638 = False

						}

						var ifres69637 Obj

						if True == ifres69638 {
							ifres69637 = True

						} else {
							ifres69637 = False

						}

						ifres69636 = ifres69637

					} else {
						ifres69636 = False

					}

					var ifres69635 Obj

					if True == ifres69636 {
						ifres69635 = True

					} else {
						ifres69635 = False

					}

					ifres69634 = ifres69635

				} else {
					ifres69634 = False

				}

				var ifres69633 Obj

				if True == ifres69634 {
					ifres69633 = True

				} else {
					ifres69633 = False

				}

				ifres69632 = ifres69633

			} else {
				ifres69632 = False

			}

			var ifres69631 Obj

			if True == ifres69632 {
				ifres69631 = True

			} else {
				ifres69631 = False

			}

			ifres69630 = ifres69631

		} else {
			ifres69630 = False

		}

		if True == ifres69630 {
			tmp69619 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

			tmp69620 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69621 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69622 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69623 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69624 := Call(__e, tmp69623, V5017)

			tmp69625 := Call(__e, tmp69622, tmp69624)

			tmp69626 := Call(__e, tmp69621, tmp69625)

			tmp69627 := Call(__e, tmp69620, tmp69626)

			tmp69628 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69629 := Call(__e, tmp69628, V5017)

			__e.TailApply(tmp69619, tmp69627, tmp69629)
			return

		} else {
			tmp69617 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69618 := Call(__e, tmp69617, V5017)

			var ifres69585 Obj

			if True == tmp69618 {
				tmp69613 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69614 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69615 := Call(__e, tmp69614, V5017)

				tmp69616 := Call(__e, tmp69613, tmp69615)

				var ifres69587 Obj

				if True == tmp69616 {
					tmp69607 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69608 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69609 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69610 := Call(__e, tmp69609, V5017)

					tmp69611 := Call(__e, tmp69608, tmp69610)

					tmp69612 := Call(__e, tmp69607, tmp69611)

					var ifres69589 Obj

					if True == tmp69612 {
						tmp69599 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp69600 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69602 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69603 := Call(__e, tmp69602, V5017)

						tmp69604 := Call(__e, tmp69601, tmp69603)

						tmp69605 := Call(__e, tmp69600, tmp69604)

						tmp69606 := Call(__e, tmp69599, Nil, tmp69605)

						var ifres69591 Obj

						if True == tmp69606 {
							tmp69593 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69594 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69595 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69596 := Call(__e, tmp69595, V5017)

							tmp69597 := Call(__e, tmp69594, tmp69596)

							tmp69598 := Call(__e, tmp69593, tmp69597, V5016)

							var ifres69592 Obj

							if True == tmp69598 {
								ifres69592 = True

							} else {
								ifres69592 = False

							}

							ifres69591 = ifres69592

						} else {
							ifres69591 = False

						}

						var ifres69590 Obj

						if True == ifres69591 {
							ifres69590 = True

						} else {
							ifres69590 = False

						}

						ifres69589 = ifres69590

					} else {
						ifres69589 = False

					}

					var ifres69588 Obj

					if True == ifres69589 {
						ifres69588 = True

					} else {
						ifres69588 = False

					}

					ifres69587 = ifres69588

				} else {
					ifres69587 = False

				}

				var ifres69586 Obj

				if True == ifres69587 {
					ifres69586 = True

				} else {
					ifres69586 = False

				}

				ifres69585 = ifres69586

			} else {
				ifres69585 = False

			}

			if True == ifres69585 {
				tmp69578 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4false_1branch)

				tmp69579 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69580 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69581 := Call(__e, tmp69580, V5017)

				tmp69582 := Call(__e, tmp69579, tmp69581)

				tmp69583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69584 := Call(__e, tmp69583, V5017)

				__e.TailApply(tmp69578, tmp69582, tmp69584)
				return

			} else {
				__e.Return(V5017)
				return
			}

		}

	}, 2)

	tmp69722 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4false_1branch, tmp69575)

	_ = tmp69722

	tmp69723 := MakeNative(func(__e *ControlFlow) {
		V5020 := __e.Get(1)
		_ = V5020
		V5021 := __e.Get(2)
		_ = V5021
		tmp69808 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp69809 := Call(__e, tmp69808, V5020)

		var ifres69778 Obj

		if True == tmp69809 {
			tmp69804 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp69805 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69806 := Call(__e, tmp69805, V5020)

			tmp69807 := Call(__e, tmp69804, sym_f_freturn, tmp69806)

			var ifres69780 Obj

			if True == tmp69807 {
				tmp69800 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69801 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp69802 := Call(__e, tmp69801, V5020)

				tmp69803 := Call(__e, tmp69800, tmp69802)

				var ifres69782 Obj

				if True == tmp69803 {
					tmp69794 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69797 := Call(__e, tmp69796, V5020)

					tmp69798 := Call(__e, tmp69795, tmp69797)

					tmp69799 := Call(__e, tmp69794, Nil, tmp69798)

					var ifres69784 Obj

					if True == tmp69799 {
						tmp69786 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						tmp69787 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69788 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69789 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69790 := Call(__e, tmp69789, V5020)

						tmp69791 := Call(__e, tmp69788, tmp69790)

						tmp69792 := Call(__e, tmp69787, tmp69791)

						tmp69793 := Call(__e, tmp69786, tmp69792)

						var ifres69785 Obj

						if True == tmp69793 {
							ifres69785 = True

						} else {
							ifres69785 = False

						}

						ifres69784 = ifres69785

					} else {
						ifres69784 = False

					}

					var ifres69783 Obj

					if True == ifres69784 {
						ifres69783 = True

					} else {
						ifres69783 = False

					}

					ifres69782 = ifres69783

				} else {
					ifres69782 = False

				}

				var ifres69781 Obj

				if True == ifres69782 {
					ifres69781 = True

				} else {
					ifres69781 = False

				}

				ifres69780 = ifres69781

			} else {
				ifres69780 = False

			}

			var ifres69779 Obj

			if True == ifres69780 {
				ifres69779 = True

			} else {
				ifres69779 = False

			}

			ifres69778 = ifres69779

		} else {
			ifres69778 = False

		}

		if True == ifres69778 {
			__e.TailApply(V5021, V5020)
			return
		} else {
			tmp69776 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp69777 := Call(__e, tmp69776, V5020)

			var ifres69764 Obj

			if True == tmp69777 {
				tmp69772 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp69773 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp69774 := Call(__e, tmp69773, V5020)

				tmp69775 := Call(__e, tmp69772, symfail, tmp69774)

				var ifres69766 Obj

				if True == tmp69775 {
					tmp69768 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69769 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69770 := Call(__e, tmp69769, V5020)

					tmp69771 := Call(__e, tmp69768, Nil, tmp69770)

					var ifres69767 Obj

					if True == tmp69771 {
						ifres69767 = True

					} else {
						ifres69767 = False

					}

					ifres69766 = ifres69767

				} else {
					ifres69766 = False

				}

				var ifres69765 Obj

				if True == ifres69766 {
					ifres69765 = True

				} else {
					ifres69765 = False

				}

				ifres69764 = ifres69765

			} else {
				ifres69764 = False

			}

			if True == ifres69764 {
				__e.TailApply(V5021, V5020)
				return
			} else {
				tmp69762 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69763 := Call(__e, tmp69762, V5020)

				var ifres69742 Obj

				if True == tmp69763 {
					tmp69758 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp69759 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp69760 := Call(__e, tmp69759, V5020)

					tmp69761 := Call(__e, tmp69758, sym_f_fgoto_1label, tmp69760)

					var ifres69744 Obj

					if True == tmp69761 {
						tmp69754 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69755 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69756 := Call(__e, tmp69755, V5020)

						tmp69757 := Call(__e, tmp69754, tmp69756)

						var ifres69746 Obj

						if True == tmp69757 {
							tmp69748 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69749 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69750 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69751 := Call(__e, tmp69750, V5020)

							tmp69752 := Call(__e, tmp69749, tmp69751)

							tmp69753 := Call(__e, tmp69748, Nil, tmp69752)

							var ifres69747 Obj

							if True == tmp69753 {
								ifres69747 = True

							} else {
								ifres69747 = False

							}

							ifres69746 = ifres69747

						} else {
							ifres69746 = False

						}

						var ifres69745 Obj

						if True == ifres69746 {
							ifres69745 = True

						} else {
							ifres69745 = False

						}

						ifres69744 = ifres69745

					} else {
						ifres69744 = False

					}

					var ifres69743 Obj

					if True == ifres69744 {
						ifres69743 = True

					} else {
						ifres69743 = False

					}

					ifres69742 = ifres69743

				} else {
					ifres69742 = False

				}

				if True == ifres69742 {
					__e.TailApply(V5021, V5020)
					return
				} else {
					tmp69727 := MakeNative(func(__e *ControlFlow) {
						Label := __e.Get(1)
						_ = Label
						tmp69728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp69729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp69730 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp69731 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp69732 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp69733 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp69734 := Call(__e, tmp69733, Label, Nil)

						tmp69735 := Call(__e, tmp69732, sym_f_fgoto_1label, tmp69734)

						tmp69736 := Call(__e, V5021, tmp69735)

						tmp69737 := Call(__e, tmp69731, tmp69736, Nil)

						tmp69738 := Call(__e, tmp69730, V5020, tmp69737)

						tmp69739 := Call(__e, tmp69729, Label, tmp69738)

						__e.TailApply(tmp69728, sym_f_flet_1label, tmp69739)
						return

					}, 1)

					tmp69740 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4generate_1label)

					tmp69741 := Call(__e, tmp69740)

					__e.TailApply(tmp69727, tmp69741)
					return

				}

			}

		}

	}, 2)

	tmp69810 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4with_1labelled_1else, tmp69723)

	_ = tmp69810

	tmp69811 := MakeNative(func(__e *ControlFlow) {
		V5024 := __e.Get(1)
		_ = V5024
		tmp70006 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp70007 := Call(__e, tmp70006, V5024)

		var ifres69856 Obj

		if True == tmp70007 {
			tmp70002 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp70003 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp70004 := Call(__e, tmp70003, V5024)

			tmp70005 := Call(__e, tmp70002, symif, tmp70004)

			var ifres69858 Obj

			if True == tmp70005 {
				tmp69998 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp69999 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70000 := Call(__e, tmp69999, V5024)

				tmp70001 := Call(__e, tmp69998, tmp70000)

				var ifres69860 Obj

				if True == tmp70001 {
					tmp69992 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp69993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69994 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp69995 := Call(__e, tmp69994, V5024)

					tmp69996 := Call(__e, tmp69993, tmp69995)

					tmp69997 := Call(__e, tmp69992, tmp69996)

					var ifres69862 Obj

					if True == tmp69997 {
						tmp69984 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp69985 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp69986 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp69988 := Call(__e, tmp69987, V5024)

						tmp69989 := Call(__e, tmp69986, tmp69988)

						tmp69990 := Call(__e, tmp69985, tmp69989)

						tmp69991 := Call(__e, tmp69984, tmp69990)

						var ifres69864 Obj

						if True == tmp69991 {
							tmp69974 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp69975 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69976 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp69977 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69978 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp69979 := Call(__e, tmp69978, V5024)

							tmp69980 := Call(__e, tmp69977, tmp69979)

							tmp69981 := Call(__e, tmp69976, tmp69980)

							tmp69982 := Call(__e, tmp69975, tmp69981)

							tmp69983 := Call(__e, tmp69974, symif, tmp69982)

							var ifres69866 Obj

							if True == tmp69983 {
								tmp69964 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp69965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp69967 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69968 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp69969 := Call(__e, tmp69968, V5024)

								tmp69970 := Call(__e, tmp69967, tmp69969)

								tmp69971 := Call(__e, tmp69966, tmp69970)

								tmp69972 := Call(__e, tmp69965, tmp69971)

								tmp69973 := Call(__e, tmp69964, tmp69972)

								var ifres69868 Obj

								if True == tmp69973 {
									tmp69952 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp69953 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69954 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69955 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp69956 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp69958 := Call(__e, tmp69957, V5024)

									tmp69959 := Call(__e, tmp69956, tmp69958)

									tmp69960 := Call(__e, tmp69955, tmp69959)

									tmp69961 := Call(__e, tmp69954, tmp69960)

									tmp69962 := Call(__e, tmp69953, tmp69961)

									tmp69963 := Call(__e, tmp69952, tmp69962)

									var ifres69870 Obj

									if True == tmp69963 {
										tmp69938 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp69939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69941 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69942 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp69943 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69944 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp69945 := Call(__e, tmp69944, V5024)

										tmp69946 := Call(__e, tmp69943, tmp69945)

										tmp69947 := Call(__e, tmp69942, tmp69946)

										tmp69948 := Call(__e, tmp69941, tmp69947)

										tmp69949 := Call(__e, tmp69940, tmp69948)

										tmp69950 := Call(__e, tmp69939, tmp69949)

										tmp69951 := Call(__e, tmp69938, tmp69950)

										var ifres69872 Obj

										if True == tmp69951 {
											tmp69922 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp69923 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69924 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69925 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69926 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69927 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp69928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69929 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp69930 := Call(__e, tmp69929, V5024)

											tmp69931 := Call(__e, tmp69928, tmp69930)

											tmp69932 := Call(__e, tmp69927, tmp69931)

											tmp69933 := Call(__e, tmp69926, tmp69932)

											tmp69934 := Call(__e, tmp69925, tmp69933)

											tmp69935 := Call(__e, tmp69924, tmp69934)

											tmp69936 := Call(__e, tmp69923, tmp69935)

											tmp69937 := Call(__e, tmp69922, Nil, tmp69936)

											var ifres69874 Obj

											if True == tmp69937 {
												tmp69914 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp69915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp69916 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp69917 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp69918 := Call(__e, tmp69917, V5024)

												tmp69919 := Call(__e, tmp69916, tmp69918)

												tmp69920 := Call(__e, tmp69915, tmp69919)

												tmp69921 := Call(__e, tmp69914, tmp69920)

												var ifres69876 Obj

												if True == tmp69921 {
													tmp69904 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp69905 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp69906 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp69907 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp69908 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp69909 := Call(__e, tmp69908, V5024)

													tmp69910 := Call(__e, tmp69907, tmp69909)

													tmp69911 := Call(__e, tmp69906, tmp69910)

													tmp69912 := Call(__e, tmp69905, tmp69911)

													tmp69913 := Call(__e, tmp69904, Nil, tmp69912)

													var ifres69878 Obj

													if True == tmp69913 {
														tmp69880 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp69881 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp69882 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69884 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69885 := Call(__e, tmp69884, V5024)

														tmp69886 := Call(__e, tmp69883, tmp69885)

														tmp69887 := Call(__e, tmp69882, tmp69886)

														tmp69888 := Call(__e, tmp69881, tmp69887)

														tmp69889 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp69890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69891 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69892 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69893 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp69894 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69895 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp69896 := Call(__e, tmp69895, V5024)

														tmp69897 := Call(__e, tmp69894, tmp69896)

														tmp69898 := Call(__e, tmp69893, tmp69897)

														tmp69899 := Call(__e, tmp69892, tmp69898)

														tmp69900 := Call(__e, tmp69891, tmp69899)

														tmp69901 := Call(__e, tmp69890, tmp69900)

														tmp69902 := Call(__e, tmp69889, tmp69901)

														tmp69903 := Call(__e, tmp69880, tmp69888, tmp69902)

														var ifres69879 Obj

														if True == tmp69903 {
															ifres69879 = True

														} else {
															ifres69879 = False

														}

														ifres69878 = ifres69879

													} else {
														ifres69878 = False

													}

													var ifres69877 Obj

													if True == ifres69878 {
														ifres69877 = True

													} else {
														ifres69877 = False

													}

													ifres69876 = ifres69877

												} else {
													ifres69876 = False

												}

												var ifres69875 Obj

												if True == ifres69876 {
													ifres69875 = True

												} else {
													ifres69875 = False

												}

												ifres69874 = ifres69875

											} else {
												ifres69874 = False

											}

											var ifres69873 Obj

											if True == ifres69874 {
												ifres69873 = True

											} else {
												ifres69873 = False

											}

											ifres69872 = ifres69873

										} else {
											ifres69872 = False

										}

										var ifres69871 Obj

										if True == ifres69872 {
											ifres69871 = True

										} else {
											ifres69871 = False

										}

										ifres69870 = ifres69871

									} else {
										ifres69870 = False

									}

									var ifres69869 Obj

									if True == ifres69870 {
										ifres69869 = True

									} else {
										ifres69869 = False

									}

									ifres69868 = ifres69869

								} else {
									ifres69868 = False

								}

								var ifres69867 Obj

								if True == ifres69868 {
									ifres69867 = True

								} else {
									ifres69867 = False

								}

								ifres69866 = ifres69867

							} else {
								ifres69866 = False

							}

							var ifres69865 Obj

							if True == ifres69866 {
								ifres69865 = True

							} else {
								ifres69865 = False

							}

							ifres69864 = ifres69865

						} else {
							ifres69864 = False

						}

						var ifres69863 Obj

						if True == ifres69864 {
							ifres69863 = True

						} else {
							ifres69863 = False

						}

						ifres69862 = ifres69863

					} else {
						ifres69862 = False

					}

					var ifres69861 Obj

					if True == ifres69862 {
						ifres69861 = True

					} else {
						ifres69861 = False

					}

					ifres69860 = ifres69861

				} else {
					ifres69860 = False

				}

				var ifres69859 Obj

				if True == ifres69860 {
					ifres69859 = True

				} else {
					ifres69859 = False

				}

				ifres69858 = ifres69859

			} else {
				ifres69858 = False

			}

			var ifres69857 Obj

			if True == ifres69858 {
				ifres69857 = True

			} else {
				ifres69857 = False

			}

			ifres69856 = ifres69857

		} else {
			ifres69856 = False

		}

		if True == ifres69856 {
			tmp69813 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69814 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69815 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69819 := Call(__e, tmp69818, V5024)

			tmp69820 := Call(__e, tmp69817, tmp69819)

			tmp69821 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69822 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69823 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69824 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69825 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69826 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69827 := Call(__e, tmp69826, V5024)

			tmp69828 := Call(__e, tmp69825, tmp69827)

			tmp69829 := Call(__e, tmp69824, tmp69828)

			tmp69830 := Call(__e, tmp69823, tmp69829)

			tmp69831 := Call(__e, tmp69822, tmp69830)

			tmp69832 := Call(__e, tmp69821, tmp69831, Nil)

			tmp69833 := Call(__e, tmp69816, tmp69820, tmp69832)

			tmp69834 := Call(__e, tmp69815, symand, tmp69833)

			tmp69835 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp69836 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69838 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69839 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp69840 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69841 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69842 := Call(__e, tmp69841, V5024)

			tmp69843 := Call(__e, tmp69840, tmp69842)

			tmp69844 := Call(__e, tmp69839, tmp69843)

			tmp69845 := Call(__e, tmp69838, tmp69844)

			tmp69846 := Call(__e, tmp69837, tmp69845)

			tmp69847 := Call(__e, tmp69836, tmp69846)

			tmp69848 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69849 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69850 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp69851 := Call(__e, tmp69850, V5024)

			tmp69852 := Call(__e, tmp69849, tmp69851)

			tmp69853 := Call(__e, tmp69848, tmp69852)

			tmp69854 := Call(__e, tmp69835, tmp69847, tmp69853)

			tmp69855 := Call(__e, tmp69814, tmp69834, tmp69854)

			__e.TailApply(tmp69813, symif, tmp69855)
			return

		} else {
			__e.Return(V5024)
			return
		}

	}, 1)

	tmp70008 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs, tmp69811)

	_ = tmp70008

	tmp70009 := MakeNative(func(__e *ControlFlow) {
		V5027 := __e.Get(1)
		_ = V5027
		V5028 := __e.Get(2)
		_ = V5028
		tmp70010 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		tmp70011 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		tmp70012 := Call(__e, tmp70011, sym_c, V5028)

		__e.TailApply(tmp70010, V5027, tmp70012)
		return

	}, 2)

	tmp70013 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4concat_c, tmp70009)

	_ = tmp70013

	tmp70014 := MakeNative(func(__e *ControlFlow) {
		V5032 := __e.Get(1)
		_ = V5032
		tmp70051 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp70052 := Call(__e, tmp70051, V5032)

		var ifres70031 Obj

		if True == tmp70052 {
			tmp70047 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp70048 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp70049 := Call(__e, tmp70048, V5032)

			tmp70050 := Call(__e, tmp70047, tmp70049)

			var ifres70033 Obj

			if True == tmp70050 {
				tmp70041 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp70042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70043 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70044 := Call(__e, tmp70043, V5032)

				tmp70045 := Call(__e, tmp70042, tmp70044)

				tmp70046 := Call(__e, tmp70041, Nil, tmp70045)

				var ifres70035 Obj

				if True == tmp70046 {
					tmp70037 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

					tmp70038 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp70039 := Call(__e, tmp70038, V5032)

					tmp70040 := Call(__e, tmp70037, tmp70039)

					var ifres70036 Obj

					if True == tmp70040 {
						ifres70036 = True

					} else {
						ifres70036 = False

					}

					ifres70035 = ifres70036

				} else {
					ifres70035 = False

				}

				var ifres70034 Obj

				if True == ifres70035 {
					ifres70034 = True

				} else {
					ifres70034 = False

				}

				ifres70033 = ifres70034

			} else {
				ifres70033 = False

			}

			var ifres70032 Obj

			if True == ifres70033 {
				ifres70032 = True

			} else {
				ifres70032 = False

			}

			ifres70031 = ifres70032

		} else {
			ifres70031 = False

		}

		if True == ifres70031 {
			tmp70022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4concat_c)

			tmp70023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4exp_1var)

			tmp70024 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp70025 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp70026 := Call(__e, tmp70025, V5032)

			tmp70027 := Call(__e, tmp70024, tmp70026)

			tmp70028 := Call(__e, tmp70023, tmp70027)

			tmp70029 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp70030 := Call(__e, tmp70029, V5032)

			__e.TailApply(tmp70022, tmp70028, tmp70030)
			return

		} else {
			tmp70020 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp70021 := Call(__e, tmp70020, V5032)

			if True == tmp70021 {
				tmp70017 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				tmp70018 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp70019 := Call(__e, tmp70018, V5032)

				__e.TailApply(tmp70017, tmp70019)
				return

			} else {
				__e.Return(V5032)
				return
			}

		}

	}, 1)

	tmp70053 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4exp_1var, tmp70014)

	_ = tmp70053

	tmp70054 := MakeNative(func(__e *ControlFlow) {
		V5035 := __e.Get(1)
		_ = V5035
		V5036 := __e.Get(2)
		_ = V5036
		tmp70055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)

		tmp70056 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4test_1_6selectors)

		tmp70057 := Call(__e, tmp70056, V5035)

		__e.TailApply(tmp70055, tmp70057, V5036)
		return

	}, 2)

	tmp70058 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4optimize_1selectors, tmp70054)

	_ = tmp70058

	tmp70059 := MakeNative(func(__e *ControlFlow) {
		V5042 := __e.Get(1)
		_ = V5042
		tmp70204 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp70205 := Call(__e, tmp70204, V5042)

		var ifres70184 Obj

		if True == tmp70205 {
			tmp70200 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp70201 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp70202 := Call(__e, tmp70201, V5042)

			tmp70203 := Call(__e, tmp70200, symcons_2, tmp70202)

			var ifres70186 Obj

			if True == tmp70203 {
				tmp70196 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp70197 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70198 := Call(__e, tmp70197, V5042)

				tmp70199 := Call(__e, tmp70196, tmp70198)

				var ifres70188 Obj

				if True == tmp70199 {
					tmp70190 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp70191 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70192 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70193 := Call(__e, tmp70192, V5042)

					tmp70194 := Call(__e, tmp70191, tmp70193)

					tmp70195 := Call(__e, tmp70190, Nil, tmp70194)

					var ifres70189 Obj

					if True == tmp70195 {
						ifres70189 = True

					} else {
						ifres70189 = False

					}

					ifres70188 = ifres70189

				} else {
					ifres70188 = False

				}

				var ifres70187 Obj

				if True == ifres70188 {
					ifres70187 = True

				} else {
					ifres70187 = False

				}

				ifres70186 = ifres70187

			} else {
				ifres70186 = False

			}

			var ifres70185 Obj

			if True == ifres70186 {
				ifres70185 = True

			} else {
				ifres70185 = False

			}

			ifres70184 = ifres70185

		} else {
			ifres70184 = False

		}

		if True == ifres70184 {
			tmp70173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70174 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70175 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp70176 := Call(__e, tmp70175, V5042)

			tmp70177 := Call(__e, tmp70174, symhd, tmp70176)

			tmp70178 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70179 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70180 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp70181 := Call(__e, tmp70180, V5042)

			tmp70182 := Call(__e, tmp70179, symtl, tmp70181)

			tmp70183 := Call(__e, tmp70178, tmp70182, Nil)

			__e.TailApply(tmp70173, tmp70177, tmp70183)
			return

		} else {
			tmp70171 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp70172 := Call(__e, tmp70171, V5042)

			var ifres70151 Obj

			if True == tmp70172 {
				tmp70167 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp70168 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp70169 := Call(__e, tmp70168, V5042)

				tmp70170 := Call(__e, tmp70167, symtuple_2, tmp70169)

				var ifres70153 Obj

				if True == tmp70170 {
					tmp70163 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp70164 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70165 := Call(__e, tmp70164, V5042)

					tmp70166 := Call(__e, tmp70163, tmp70165)

					var ifres70155 Obj

					if True == tmp70166 {
						tmp70157 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp70158 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70159 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70160 := Call(__e, tmp70159, V5042)

						tmp70161 := Call(__e, tmp70158, tmp70160)

						tmp70162 := Call(__e, tmp70157, Nil, tmp70161)

						var ifres70156 Obj

						if True == tmp70162 {
							ifres70156 = True

						} else {
							ifres70156 = False

						}

						ifres70155 = ifres70156

					} else {
						ifres70155 = False

					}

					var ifres70154 Obj

					if True == ifres70155 {
						ifres70154 = True

					} else {
						ifres70154 = False

					}

					ifres70153 = ifres70154

				} else {
					ifres70153 = False

				}

				var ifres70152 Obj

				if True == ifres70153 {
					ifres70152 = True

				} else {
					ifres70152 = False

				}

				ifres70151 = ifres70152

			} else {
				ifres70151 = False

			}

			if True == ifres70151 {
				tmp70140 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70141 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70143 := Call(__e, tmp70142, V5042)

				tmp70144 := Call(__e, tmp70141, symfst, tmp70143)

				tmp70145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70146 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70147 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70148 := Call(__e, tmp70147, V5042)

				tmp70149 := Call(__e, tmp70146, symsnd, tmp70148)

				tmp70150 := Call(__e, tmp70145, tmp70149, Nil)

				__e.TailApply(tmp70140, tmp70144, tmp70150)
				return

			} else {
				tmp70138 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp70139 := Call(__e, tmp70138, V5042)

				var ifres70118 Obj

				if True == tmp70139 {
					tmp70134 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp70135 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp70136 := Call(__e, tmp70135, V5042)

					tmp70137 := Call(__e, tmp70134, symshen_4_7string_2, tmp70136)

					var ifres70120 Obj

					if True == tmp70137 {
						tmp70130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp70131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70132 := Call(__e, tmp70131, V5042)

						tmp70133 := Call(__e, tmp70130, tmp70132)

						var ifres70122 Obj

						if True == tmp70133 {
							tmp70124 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp70125 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70126 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70127 := Call(__e, tmp70126, V5042)

							tmp70128 := Call(__e, tmp70125, tmp70127)

							tmp70129 := Call(__e, tmp70124, Nil, tmp70128)

							var ifres70123 Obj

							if True == tmp70129 {
								ifres70123 = True

							} else {
								ifres70123 = False

							}

							ifres70122 = ifres70123

						} else {
							ifres70122 = False

						}

						var ifres70121 Obj

						if True == ifres70122 {
							ifres70121 = True

						} else {
							ifres70121 = False

						}

						ifres70120 = ifres70121

					} else {
						ifres70120 = False

					}

					var ifres70119 Obj

					if True == ifres70120 {
						ifres70119 = True

					} else {
						ifres70119 = False

					}

					ifres70118 = ifres70119

				} else {
					ifres70118 = False

				}

				if True == ifres70118 {
					tmp70107 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp70108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp70109 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70110 := Call(__e, tmp70109, V5042)

					tmp70111 := Call(__e, tmp70108, symhdstr, tmp70110)

					tmp70112 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp70113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp70114 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70115 := Call(__e, tmp70114, V5042)

					tmp70116 := Call(__e, tmp70113, symtlstr, tmp70115)

					tmp70117 := Call(__e, tmp70112, tmp70116, Nil)

					__e.TailApply(tmp70107, tmp70111, tmp70117)
					return

				} else {
					tmp70105 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp70106 := Call(__e, tmp70105, V5042)

					var ifres70085 Obj

					if True == tmp70106 {
						tmp70101 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp70102 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp70103 := Call(__e, tmp70102, V5042)

						tmp70104 := Call(__e, tmp70101, symshen_4_7vector_2, tmp70103)

						var ifres70087 Obj

						if True == tmp70104 {
							tmp70097 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp70098 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70099 := Call(__e, tmp70098, V5042)

							tmp70100 := Call(__e, tmp70097, tmp70099)

							var ifres70089 Obj

							if True == tmp70100 {
								tmp70091 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp70092 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70093 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70094 := Call(__e, tmp70093, V5042)

								tmp70095 := Call(__e, tmp70092, tmp70094)

								tmp70096 := Call(__e, tmp70091, Nil, tmp70095)

								var ifres70090 Obj

								if True == tmp70096 {
									ifres70090 = True

								} else {
									ifres70090 = False

								}

								ifres70089 = ifres70090

							} else {
								ifres70089 = False

							}

							var ifres70088 Obj

							if True == ifres70089 {
								ifres70088 = True

							} else {
								ifres70088 = False

							}

							ifres70087 = ifres70088

						} else {
							ifres70087 = False

						}

						var ifres70086 Obj

						if True == ifres70087 {
							ifres70086 = True

						} else {
							ifres70086 = False

						}

						ifres70085 = ifres70086

					} else {
						ifres70085 = False

					}

					if True == ifres70085 {
						tmp70074 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70075 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70076 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70077 := Call(__e, tmp70076, V5042)

						tmp70078 := Call(__e, tmp70075, symhdv, tmp70077)

						tmp70079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70081 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70082 := Call(__e, tmp70081, V5042)

						tmp70083 := Call(__e, tmp70080, symtlv, tmp70082)

						tmp70084 := Call(__e, tmp70079, tmp70083, Nil)

						__e.TailApply(tmp70074, tmp70078, tmp70084)
						return

					} else {
						tmp70064 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp70066 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp70067 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp70068 := Call(__e, tmp70067)

							tmp70069 := Call(__e, tmp70066, Result, tmp70068)

							if True == tmp70069 {
								__e.Return(Nil)
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp70070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4apply_1selector_1handlers)

						tmp70071 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

						tmp70072 := Call(__e, tmp70071, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

						tmp70073 := Call(__e, tmp70070, tmp70072, V5042)

						__e.TailApply(tmp70064, tmp70073)
						return

					}

				}

			}

		}

	}, 1)

	tmp70206 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4test_1_6selectors, tmp70059)

	_ = tmp70206

	tmp70207 := MakeNative(func(__e *ControlFlow) {
		V5045 := __e.Get(1)
		_ = V5045
		V5046 := __e.Get(2)
		_ = V5046
		tmp70220 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp70221 := Call(__e, tmp70220, V5045)

		if True == tmp70221 {
			tmp70213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4bind_1selector)

			tmp70214 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp70215 := Call(__e, tmp70214, V5045)

			tmp70216 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)

			tmp70217 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp70218 := Call(__e, tmp70217, V5045)

			tmp70219 := Call(__e, tmp70216, tmp70218, V5046)

			__e.TailApply(tmp70213, tmp70215, tmp70219)
			return

		} else {
			tmp70211 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp70212 := Call(__e, tmp70211, Nil, V5045)

			if True == tmp70212 {
				__e.Return(V5046)
				return
			} else {
				tmp70210 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp70210, symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)
				return

			}

		}

	}, 2)

	tmp70222 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors, tmp70207)

	_ = tmp70222

	tmp70223 := MakeNative(func(__e *ControlFlow) {
		V5053 := __e.Get(1)
		_ = V5053
		V5054 := __e.Get(2)
		_ = V5054
		tmp70237 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		tmp70238 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

		tmp70239 := Call(__e, tmp70238, V5053, V5054)

		tmp70240 := Call(__e, tmp70237, tmp70239, MakeNumber(1))

		if True == tmp70240 {
			tmp70225 := MakeNative(func(__e *ControlFlow) {
				Var := __e.Get(1)
				_ = Var
				tmp70226 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70227 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70228 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70229 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp70230 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

				tmp70231 := Call(__e, tmp70230, Var, V5053, V5054)

				tmp70232 := Call(__e, tmp70229, tmp70231, Nil)

				tmp70233 := Call(__e, tmp70228, V5053, tmp70232)

				tmp70234 := Call(__e, tmp70227, Var, tmp70233)

				__e.TailApply(tmp70226, symlet, tmp70234)
				return

			}, 1)

			tmp70235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4exp_1var)

			tmp70236 := Call(__e, tmp70235, V5053)

			__e.TailApply(tmp70225, tmp70236)
			return

		} else {
			__e.Return(V5054)
			return
		}

	}, 2)

	tmp70241 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4bind_1selector, tmp70223)

	_ = tmp70241

	tmp70242 := MakeNative(func(__e *ControlFlow) {
		V5067 := __e.Get(1)
		_ = V5067
		V5068 := __e.Get(2)
		_ = V5068
		tmp70268 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp70269 := Call(__e, tmp70268, Nil, V5067)

		if True == tmp70269 {
			tmp70267 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp70267)
			return

		} else {
			tmp70244 := MakeNative(func(__e *ControlFlow) {
				Freeze := __e.Get(1)
				_ = Freeze
				tmp70257 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp70258 := Call(__e, tmp70257, V5067)

				if True == tmp70258 {
					tmp70247 := MakeNative(func(__e *ControlFlow) {
						Result := __e.Get(1)
						_ = Result
						tmp70250 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp70251 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						tmp70252 := Call(__e, tmp70251)

						tmp70253 := Call(__e, tmp70250, Result, tmp70252)

						if True == tmp70253 {
							tmp70249 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

							__e.TailApply(tmp70249, Freeze)
							return

						} else {
							__e.Return(Result)
							return
						}

					}, 1)

					tmp70254 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp70255 := Call(__e, tmp70254, V5067)

					tmp70256 := Call(__e, tmp70255, V5068)

					__e.TailApply(tmp70247, tmp70256)
					return

				} else {
					tmp70246 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

					__e.TailApply(tmp70246, Freeze)
					return

				}

			}, 1)

			tmp70259 := MakeNative(func(__e *ControlFlow) {
				tmp70265 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp70266 := Call(__e, tmp70265, V5067)

				if True == tmp70266 {
					tmp70262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4apply_1selector_1handlers)

					tmp70263 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70264 := Call(__e, tmp70263, V5067)

					__e.TailApply(tmp70262, tmp70264, V5068)
					return

				} else {
					tmp70261 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp70261, symshen_4x_4factorise_1defun_4apply_1selector_1handlers)
					return

				}

			}, 0)

			__e.TailApply(tmp70244, tmp70259)
			return

		}

	}, 2)

	tmp70270 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4apply_1selector_1handlers, tmp70242)

	_ = tmp70270

	tmp70271 := MakeNative(func(__e *ControlFlow) {
		tmp70272 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp70273 := Call(__e, tmp70272, symshen_4x_4factorise_1defun_4_dselector_1handlers_d, Nil)

		_ = tmp70273

		tmp70274 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp70275 := Call(__e, tmp70274, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, Nil)

		_ = tmp70275

		__e.Return(symshen_4x_4factorise_1defun_4done)
		return

	}, 0)

	tmp70276 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4initialise, tmp70271)

	_ = tmp70276

	tmp70277 := MakeNative(func(__e *ControlFlow) {
		V5070 := __e.Get(1)
		_ = V5070
		tmp70293 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp70294 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp70295 := Call(__e, tmp70294, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

		tmp70296 := Call(__e, tmp70293, V5070, tmp70295)

		if True == tmp70296 {
			__e.Return(V5070)
			return
		} else {
			tmp70279 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp70280 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70281 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp70282 := Call(__e, tmp70281, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

			tmp70283 := Call(__e, tmp70280, V5070, tmp70282)

			tmp70284 := Call(__e, tmp70279, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, tmp70283)

			_ = tmp70284

			tmp70285 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp70286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70287 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

			tmp70288 := Call(__e, tmp70287, V5070)

			tmp70289 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp70290 := Call(__e, tmp70289, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

			tmp70291 := Call(__e, tmp70286, tmp70288, tmp70290)

			tmp70292 := Call(__e, tmp70285, symshen_4x_4factorise_1defun_4_dselector_1handlers_d, tmp70291)

			_ = tmp70292

			__e.Return(V5070)
			return

		}

	}, 1)

	tmp70297 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4register_1selector_1handler, tmp70277)

	_ = tmp70297

	tmp70298 := MakeNative(func(__e *ControlFlow) {
		V5073 := __e.Get(1)
		_ = V5073
		V5074 := __e.Get(2)
		_ = V5074
		tmp70299 := MakeNative(func(__e *ControlFlow) {
			tmp70300 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findpos)

			__e.TailApply(tmp70300, V5073, V5074)
			return

		}, 0)

		tmp70301 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			tmp70302 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp70303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp70304 := Call(__e, tmp70303, V5073, MakeString(" is not a selector handler\n"), symshen_4a)

			__e.TailApply(tmp70302, tmp70304)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp70299, tmp70301)
		return

	}, 2)

	tmp70305 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4findpos, tmp70298)

	_ = tmp70305

	tmp70306 := MakeNative(func(__e *ControlFlow) {
		V5076 := __e.Get(1)
		_ = V5076
		tmp70307 := MakeNative(func(__e *ControlFlow) {
			Reg := __e.Get(1)
			_ = Reg
			tmp70308 := MakeNative(func(__e *ControlFlow) {
				Pos := __e.Get(1)
				_ = Pos
				tmp70309 := MakeNative(func(__e *ControlFlow) {
					RemoveReg := __e.Get(1)
					_ = RemoveReg
					tmp70310 := MakeNative(func(__e *ControlFlow) {
						RemoveFun := __e.Get(1)
						_ = RemoveFun
						__e.Return(V5076)
						return
					}, 1)

					tmp70311 := Call(__e, PrimNS1Value(symns2_1value), symset)

					tmp70312 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1nth)

					tmp70313 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					tmp70314 := Call(__e, tmp70313, symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

					tmp70315 := Call(__e, tmp70312, Pos, tmp70314)

					tmp70316 := Call(__e, tmp70311, symshen_4x_4factorise_1defun_4_dselector_1handlers_d, tmp70315)

					__e.TailApply(tmp70310, tmp70316)
					return

				}, 1)

				tmp70317 := Call(__e, PrimNS1Value(symns2_1value), symset)

				tmp70318 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				tmp70319 := Call(__e, tmp70318, V5076, Reg)

				tmp70320 := Call(__e, tmp70317, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, tmp70319)

				__e.TailApply(tmp70309, tmp70320)
				return

			}, 1)

			tmp70321 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4factorise_1defun_4findpos)

			tmp70322 := Call(__e, tmp70321, V5076, Reg)

			__e.TailApply(tmp70308, tmp70322)
			return

		}, 1)

		tmp70323 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp70324 := Call(__e, tmp70323, symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d)

		__e.TailApply(tmp70307, tmp70324)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4unregister_1selector_1handler, tmp70306)
	return

}, 0)

var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symMessage = MakeSymbol("Message")
var symshen_4legitimate_1term_2 = MakeSymbol("shen.legitimate-term?")
var symintern = MakeSymbol("intern")
var symshen_4linearise__help = MakeSymbol("shen.linearise_help")
var symshen_4embed_1and = MakeSymbol("shen.embed-and")
var symmacroexpand = MakeSymbol("macroexpand")
var sym_dport_d = MakeSymbol("*port*")
var symdatatype = MakeSymbol("datatype")
var symshen_4_5non_1ll_1rules_6 = MakeSymbol("shen.<non-ll-rules>")
var sym_h_1 = MakeSymbol(":-")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4procedure__name = MakeSymbol("shen.procedure_name")
var sym_5_a = MakeSymbol("<=")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symsubst = MakeSymbol("subst")
var symshen_4cc__body = MakeSymbol("shen.cc_body")
var symshen_4encode_1choices = MakeSymbol("shen.encode-choices")
var symshen_4a = MakeSymbol("shen.a")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var symC = MakeSymbol("C")
var symshen_4clause__form = MakeSymbol("shen.clause_form")
var symshow_1help = MakeSymbol("show-help")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4type_1signature_1of_1step = MakeSymbol("shen.type-signature-of-step")
var symshen_4type_1signature_1of_1variable_2 = MakeSymbol("shen.type-signature-of-variable?")
var symshen_4resizeprocessvector = MakeSymbol("shen.resizeprocessvector")
var symstr = MakeSymbol("str")
var symX = MakeSymbol("X")
var symdestroy = MakeSymbol("destroy")
var symshen_4type_1signature_1of_1porters = MakeSymbol("shen.type-signature-of-porters")
var symreceive = MakeSymbol("receive")
var symunify = MakeSymbol("unify")
var symshen_4require = MakeSymbol("shen.require")
var symshen_4type_1signature_1of_1simple_1error = MakeSymbol("shen.type-signature-of-simple-error")
var symshen_4_5sig_7rules_6 = MakeSymbol("shen.<sig+rules>")
var symhd = MakeSymbol("hd")
var symD = MakeSymbol("D")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var sym_dlanguage_d = MakeSymbol("*language*")
var symshen_4type_1signature_1of_1gensym = MakeSymbol("shen.type-signature-of-gensym")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4call_1help = MakeSymbol("shen.call-help")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4read_1file_1as_1Xlist_1help = MakeSymbol("shen.read-file-as-Xlist-help")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4multiple_1set = MakeSymbol("shen.multiple-set")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4_dcatch_d = MakeSymbol("shen.*catch*")
var symshen_4_5body_d_6 = MakeSymbol("shen.<body*>")
var symshen_4type_1signature_1of_1spy = MakeSymbol("shen.type-signature-of-spy")
var symContext__1957 = MakeSymbol("Context_1957")
var symspy = MakeSymbol("spy")
var symshen_4reduce = MakeSymbol("shen.reduce")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4placeholders = MakeSymbol("shen.placeholders")
var symshen_4write_1char_1and_1inc = MakeSymbol("shen.write-char-and-inc")
var sym_a = MakeSymbol("=")
var symstring_2 = MakeSymbol("string?")
var symshen_4s_1prolog__literal = MakeSymbol("shen.s-prolog_literal")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4type_1signature_1of_1stinput = MakeSymbol("shen.type-signature-of-stinput")
var symshen_4pushnew = MakeSymbol("shen.pushnew")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symM = MakeSymbol("M")
var symFinish = MakeSymbol("Finish")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symfreeze = MakeSymbol("freeze")
var symshen_4type_1signature_1of_1destroy = MakeSymbol("shen.type-signature-of-destroy")
var symshen_4construct_1base_1search_1clause = MakeSymbol("shen.construct-base-search-clause")
var symAssumption__1957 = MakeSymbol("Assumption_1957")
var symnl = MakeSymbol("nl")
var sym_5_1vector = MakeSymbol("<-vector")
var symverified = MakeSymbol("verified")
var symshen_4type_1signature_1of_1empty_2 = MakeSymbol("shen.type-signature-of-empty?")
var symshen_4type_1signature_1of_1number_2 = MakeSymbol("shen.type-signature-of-number?")
var symshen_4type_1signature_1of_1tlstr = MakeSymbol("shen.type-signature-of-tlstr")
var sym_l = MakeSymbol(",")
var symshen_4prefix_2 = MakeSymbol("shen.prefix?")
var symbound_2 = MakeSymbol("bound?")
var symshen_4type_1signature_1of_1implementation = MakeSymbol("shen.type-signature-of-implementation")
var symin = MakeSymbol("in")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4csl_1help = MakeSymbol("shen.csl-help")
var symFreeze = MakeSymbol("Freeze")
var symshen_4repl = MakeSymbol("shen.repl")
var symP = MakeSymbol("P")
var symmapcan = MakeSymbol("mapcan")
var symshen_4type_1signature_1of_1unprofile = MakeSymbol("shen.type-signature-of-unprofile")
var symshen_4curry_1synonyms = MakeSymbol("shen.curry-synonyms")
var symshen_4insert_1predicate = MakeSymbol("shen.insert-predicate")
var symshen_4left_1round = MakeSymbol("shen.left-round")
var symshen_4terminal_2 = MakeSymbol("shen.terminal?")
var symhdv = MakeSymbol("hdv")
var symshen_4type_1signature_1of_1hdstr = MakeSymbol("shen.type-signature-of-hdstr")
var symshen_4type_1signature_1of_1nth = MakeSymbol("shen.type-signature-of-nth")
var symshen_4percent = MakeSymbol("shen.percent")
var symremove = MakeSymbol("remove")
var symnot = MakeSymbol("not")
var symor = MakeSymbol("or")
var symshen_4type_1signature_1of_1shen_4insert = MakeSymbol("shen.type-signature-of-shen.insert")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4errormaxinfs = MakeSymbol("shen.errormaxinfs")
var symoutput = MakeSymbol("output")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4err_1condition = MakeSymbol("shen.err-condition")
var symerror = MakeSymbol("error")
var symshen_4non_1empty = MakeSymbol("shen.non-empty")
var symshen_4left_1rule = MakeSymbol("shen.left-rule")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4after_1codepoint = MakeSymbol("shen.after-codepoint")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var sym_c_4 = MakeSymbol("/.")
var symundefmacro = MakeSymbol("undefmacro")
var sympr = MakeSymbol("pr")
var symshen_4type_1signature_1of_1_5_1vector = MakeSymbol("shen.type-signature-of-<-vector")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4type_1signature_1of_1tc = MakeSymbol("shen.type-signature-of-tc")
var symNPP = MakeSymbol("NPP")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4recursive__cons__form = MakeSymbol("shen.recursive_cons_form")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symshen_4type_1signature_1of_1n_1_6string = MakeSymbol("shen.type-signature-of-n->string")
var symshen_4_5predigits_6 = MakeSymbol("shen.<predigits>")
var symQv = MakeSymbol("Qv")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4place = MakeSymbol("shen.place")
var symshen_4post = MakeSymbol("shen.post")
var symunify_b = MakeSymbol("unify!")
var symstream = MakeSymbol("stream")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4double_1_6singles = MakeSymbol("shen.double->singles")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4x_4factorise_1defun_4add_1returns = MakeSymbol("shen.x.factorise-defun.add-returns")
var symshen_4compile__to__kl = MakeSymbol("shen.compile_to_kl")
var symshen_4type_1signature_1of_1enable_1type_1theory = MakeSymbol("shen.type-signature-of-enable-type-theory")
var symshen_4catch_1cut = MakeSymbol("shen.catch-cut")
var symshen_4dh_2 = MakeSymbol("shen.dh?")
var symshen_4cl = MakeSymbol("shen.cl")
var symshen_4preclude_1h = MakeSymbol("shen.preclude-h")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4complexity__head = MakeSymbol("shen.complexity_head")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symshen_4_5non_1return_6 = MakeSymbol("shen.<non-return>")
var symContext1__1957 = MakeSymbol("Context1_1957")
var symshen_4x_4features_4cond_1expand_1macro = MakeSymbol("shen.x.features.cond-expand-macro")
var symI = MakeSymbol("I")
var symtc_2 = MakeSymbol("tc?")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4x_4factorise_1defun_4attach_1free_1variables = MakeSymbol("shen.x.factorise-defun.attach-free-variables")
var symshen_4application__build = MakeSymbol("shen.application_build")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symY = MakeSymbol("Y")
var symshen_4remtype = MakeSymbol("shen.remtype")
var symshen_4insert_1deref = MakeSymbol("shen.insert-deref")
var symshen_4record_1exceptions = MakeSymbol("shen.record-exceptions")
var symshen_4_5premise_6 = MakeSymbol("shen.<premise>")
var sym_dmacros_d = MakeSymbol("*macros*")
var symR = MakeSymbol("R")
var symunprofile = MakeSymbol("unprofile")
var symshen_4type_1signature_1of_1difference = MakeSymbol("shen.type-signature-of-difference")
var symshen_4type_1signature_1of_1_6 = MakeSymbol("shen.type-signature-of->")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4_5predicate_d_6 = MakeSymbol("shen.<predicate*>")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4x_4factorise_1defun_4factorise_1defun = MakeSymbol("shen.x.factorise-defun.factorise-defun")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var syminternal = MakeSymbol("internal")
var symshen_4type_1signature_1of_1freeze = MakeSymbol("shen.type-signature-of-freeze")
var symshen_4to = MakeSymbol("shen.to")
var symshen_4x_4factorise_1defun_4factorise_1cond = MakeSymbol("shen.x.factorise-defun.factorise-cond")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symps = MakeSymbol("ps")
var symget = MakeSymbol("get")
var symfunction = MakeSymbol("function")
var sym_h = MakeSymbol(":")
var symshen_4type_1signature_1of_1it = MakeSymbol("shen.type-signature-of-it")
var symshen_4type_1signature_1of_1union = MakeSymbol("shen.type-signature-of-union")
var symshen_4packageh = MakeSymbol("shen.packageh")
var symshen_4jump__stream = MakeSymbol("shen.jump_stream")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4aritycheck = MakeSymbol("shen.aritycheck")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4atom_1type = MakeSymbol("shen.atom-type")
var symshen_4type_1signature_1of_1package_2 = MakeSymbol("shen.type-signature-of-package?")
var symshen_4of = MakeSymbol("shen.of")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symshen_4f__error = MakeSymbol("shen.f_error")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4linearise__X = MakeSymbol("shen.linearise_X")
var symshen_4x_4features_4current = MakeSymbol("shen.x.features.current")
var symshen_4intprolog = MakeSymbol("shen.intprolog")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4ue_2 = MakeSymbol("shen.ue?")
var symshen_4t_d_1action = MakeSymbol("shen.t*-action")
var symns2_1value = MakeSymbol("ns2-value")
var symshen_4extract__free__vars = MakeSymbol("shen.extract_free_vars")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4type_1signature_1of_1vector_1_6 = MakeSymbol("shen.type-signature-of-vector->")
var symContext = MakeSymbol("Context")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4type_1signature_1of_1boolean_2 = MakeSymbol("shen.type-signature-of-boolean?")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symspecialise = MakeSymbol("specialise")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4type_1signature_1of_1not = MakeSymbol("shen.type-signature-of-not")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4_dsignedfuncs_d = MakeSymbol("shen.*signedfuncs*")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var symshen_4type_1signature_1of_1occurrences = MakeSymbol("shen.type-signature-of-occurrences")
var symshen_4terminator_2 = MakeSymbol("shen.terminator?")
var symshen_4x_4factorise_1defun_4free_1variables_1h = MakeSymbol("shen.x.factorise-defun.free-variables-h")
var symport = MakeSymbol("port")
var symH = MakeSymbol("H")
var symdefmacro = MakeSymbol("defmacro")
var symProcessN = MakeSymbol("ProcessN")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symParse__Y = MakeSymbol("Parse_Y")
var sym_6_6 = MakeSymbol(">>")
var symshen_4prolog_1_6shen = MakeSymbol("shen.prolog->shen")
var symshen_4mu__reduction = MakeSymbol("shen.mu_reduction")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4trim_1whitespace = MakeSymbol("shen.trim-whitespace")
var symshen_4_5side_1conditions_6 = MakeSymbol("shen.<side-conditions>")
var symshen_4extract_1pvars = MakeSymbol("shen.extract-pvars")
var symshen_4memo = MakeSymbol("shen.memo")
var sym_i = MakeSymbol("{")
var symshen_4type_1signature_1of_1fail_1if = MakeSymbol("shen.type-signature-of-fail-if")
var symshen_4type_1signature_1of_1include = MakeSymbol("shen.type-signature-of-include")
var symshen_4type_1signature_1of_1_a_a = MakeSymbol("shen.type-signature-of-==")
var symshen_4compile__prolog__procedure = MakeSymbol("shen.compile_prolog_procedure")
var symshen_4stack = MakeSymbol("shen.stack")
var symcond = MakeSymbol("cond")
var sym_dversion_d = MakeSymbol("*version*")
var symcd = MakeSymbol("cd")
var symis = MakeSymbol("is")
var sym_dos_d = MakeSymbol("*os*")
var symshen_4type_1signature_1of_1undefmacro = MakeSymbol("shen.type-signature-of-undefmacro")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4record_1it_1h = MakeSymbol("shen.record-it-h")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4x_4factorise_1defun_4false_1branch = MakeSymbol("shen.x.factorise-defun.false-branch")
var symshen_4type_1signature_1of_1remove = MakeSymbol("shen.type-signature-of-remove")
var symshen_4initialise_1signedfunc_1lambda_1forms = MakeSymbol("shen.initialise-signedfunc-lambda-forms")
var symshen_4type_1signature_1of_1read = MakeSymbol("shen.type-signature-of-read")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4code_1point = MakeSymbol("shen.code-point")
var symshen_4reduce__help = MakeSymbol("shen.reduce_help")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4type_1signature_1of_1pr = MakeSymbol("shen.type-signature-of-pr")
var symshen_4type_1signature_1of_1snd = MakeSymbol("shen.type-signature-of-snd")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4split__cc__rule = MakeSymbol("shen.split_cc_rule")
var symshen_4reassemble = MakeSymbol("shen.reassemble")
var symshen_4type_1signature_1of_1language = MakeSymbol("shen.type-signature-of-language")
var symCase = MakeSymbol("Case")
var symshen_4_5E_6 = MakeSymbol("shen.<E>")
var symshen_4_5guard_6 = MakeSymbol("shen.<guard>")
var sym__ = MakeSymbol("_")
var symshen_4function_1macro = MakeSymbol("shen.function-macro")
var symshen_4head__abstraction = MakeSymbol("shen.head_abstraction")
var symshen_4assign_1types = MakeSymbol("shen.assign-types")
var symshen_4unix = MakeSymbol("shen.unix")
var symshen_4removetype = MakeSymbol("shen.removetype")
var symshen_4read_1file_1as_1charlist = MakeSymbol("shen.read-file-as-charlist")
var symshen_4pre = MakeSymbol("shen.pre")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4semantic_1completion_1warning = MakeSymbol("shen.semantic-completion-warning")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4newhyps = MakeSymbol("shen.newhyps")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var symshen_4typecheck_1and_1evaluate = MakeSymbol("shen.typecheck-and-evaluate")
var symshen_4_5singleunderline_6 = MakeSymbol("shen.<singleunderline>")
var symshen_4alphanum_2 = MakeSymbol("shen.alphanum?")
var symshen_4find = MakeSymbol("shen.find")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symempty_2 = MakeSymbol("empty?")
var sym_8s = MakeSymbol("@s")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4type_1signature_1of_1pos = MakeSymbol("shen.type-signature-of-pos")
var symshen_4type_1signature_1of_1unspecialise = MakeSymbol("shen.type-signature-of-unspecialise")
var symshen_4then = MakeSymbol("shen.then")
var symshen_4_5semicolon_1symbol_6 = MakeSymbol("shen.<semicolon-symbol>")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4type_1signature_1of_1bound_2 = MakeSymbol("shen.type-signature-of-bound?")
var symversion = MakeSymbol("version")
var symshen_4em__help = MakeSymbol("shen.em_help")
var symshen_4construct_1premiss_1literal = MakeSymbol("shen.construct-premiss-literal")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4package_1contents = MakeSymbol("shen.package-contents")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4syntax = MakeSymbol("shen.syntax")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4lisp_1or = MakeSymbol("shen.lisp-or")
var symshen_4_5side_1condition_6 = MakeSymbol("shen.<side-condition>")
var symshen_4update_1symbol_1table = MakeSymbol("shen.update-symbol-table")
var sym_a_a = MakeSymbol("==")
var symshen_4nest_1disjunct = MakeSymbol("shen.nest-disjunct")
var symT = MakeSymbol("T")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4type_1signature_1of_1optimise = MakeSymbol("shen.type-signature-of-optimise")
var symshen_4t_d_1defh = MakeSymbol("shen.t*-defh")
var symsymbol = MakeSymbol("symbol")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symmaxinferences = MakeSymbol("maxinferences")
var sym_3 = MakeSymbol("$")
var symshen_4default__semantics = MakeSymbol("shen.default_semantics")
var symshen_4x_4factorise_1defun_4findpos = MakeSymbol("shen.x.factorise-defun.findpos")
var symshen_4lambda_1form_1entry = MakeSymbol("shen.lambda-form-entry")
var symget_1time = MakeSymbol("get-time")
var symshen_4type_1signature_1of_1maxinferences = MakeSymbol("shen.type-signature-of-maxinferences")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var sympos = MakeSymbol("pos")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4pair = MakeSymbol("shen.pair")
var symshen_4type_1signature_1of_1fst = MakeSymbol("shen.type-signature-of-fst")
var sym_f_flet_1label = MakeSymbol("%%let-label")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4type_1signature_1of_1sum = MakeSymbol("shen.type-signature-of-sum")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var sym_j = MakeSymbol("}")
var sym_8v = MakeSymbol("@v")
var symshen_4_dprocess_1counter_d = MakeSymbol("shen.*process-counter*")
var symrun = MakeSymbol("run")
var symshen_4datatype_1error = MakeSymbol("shen.datatype-error")
var symshen_4compose = MakeSymbol("shen.compose")
var symshen_4dereferencing = MakeSymbol("shen.dereferencing")
var symshen_4_5byte_6 = MakeSymbol("shen.<byte>")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4add_1end = MakeSymbol("shen.add-end")
var symfst = MakeSymbol("fst")
var symshen_4type_1signature_1of_1integer_2 = MakeSymbol("shen.type-signature-of-integer?")
var symshen_4tab = MakeSymbol("shen.tab")
var symexplode = MakeSymbol("explode")
var symshen_4_a_a_6 = MakeSymbol("shen.==>")
var symshen_4_5term_d_6 = MakeSymbol("shen.<term*>")
var symshen_4mode_1ify = MakeSymbol("shen.mode-ify")
var symshen_4variant_2 = MakeSymbol("shen.variant?")
var symcons_2 = MakeSymbol("cons?")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symK = MakeSymbol("K")
var symshen_4type_1signature_1of_1nl = MakeSymbol("shen.type-signature-of-nl")
var symshen_4cf__help = MakeSymbol("shen.cf_help")
var symshen_4rule_1_6horn_1clause_1head = MakeSymbol("shen.rule->horn-clause-head")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4x_4factorise_1defun_4test_1_6selectors = MakeSymbol("shen.x.factorise-defun.test->selectors")
var symshen_4type_1signature_1of_1include_1all_1but = MakeSymbol("shen.type-signature-of-include-all-but")
var symshen_4_5st__input2_6 = MakeSymbol("shen.<st_input2>")
var symgensym = MakeSymbol("gensym")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4active_1cons = MakeSymbol("shen.active-cons")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4catchpoint = MakeSymbol("shen.catchpoint")
var sym_f_flabel = MakeSymbol("%%label")
var symshen_4x_4factorise_1defun_4rebranch_1h = MakeSymbol("shen.x.factorise-defun.rebranch-h")
var symshen_4_dcustom_1pattern_1compiler_d = MakeSymbol("shen.*custom-pattern-compiler*")
var sym_dargv_d = MakeSymbol("*argv*")
var symshen_4type_1signature_1of_1do = MakeSymbol("shen.type-signature-of-do")
var symshen_4type_1signature_1of_1print = MakeSymbol("shen.type-signature-of-print")
var symshen_4remove_1synonyms = MakeSymbol("shen.remove-synonyms")
var symshen_4x_4features_4initialise = MakeSymbol("shen.x.features.initialise")
var symshen_4copy_1vector = MakeSymbol("shen.copy-vector")
var symshen_4newcontinuation = MakeSymbol("shen.newcontinuation")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4control_1chars = MakeSymbol("shen.control-chars")
var symshen_4_5num_6 = MakeSymbol("shen.<num>")
var symit = MakeSymbol("it")
var symshen_4type_1signature_1of_1head = MakeSymbol("shen.type-signature-of-head")
var symshen_4type_1signature_1of_1profile_1results = MakeSymbol("shen.type-signature-of-profile-results")
var symshen_4type_1signature_1of_1_1 = MakeSymbol("shen.type-signature-of--")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4free__variable__warnings = MakeSymbol("shen.free_variable_warnings")
var symshen_4type_1signature_1of_1read_1file_1as_1bytelist = MakeSymbol("shen.type-signature-of-read-file-as-bytelist")
var symshen_4catchpoint_b = MakeSymbol("shen.catchpoint!")
var symcn = MakeSymbol("cn")
var symoccurrences = MakeSymbol("occurrences")
var symintersection = MakeSymbol("intersection")
var symshen_4type_1signature_1of_1limit = MakeSymbol("shen.type-signature-of-limit")
var symshen_4type_1signature_1of_1preclude_1all_1but = MakeSymbol("shen.type-signature-of-preclude-all-but")
var symshen_4cn_1all = MakeSymbol("shen.cn-all")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var sym_d = MakeSymbol("*")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symshen_4loop = MakeSymbol("shen.loop")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4_5doubleunderline_6 = MakeSymbol("shen.<doubleunderline>")
var symshen_4t_d_1hyps = MakeSymbol("shen.t*-hyps")
var symshen_4r = MakeSymbol("shen.r")
var symcases = MakeSymbol("cases")
var symshen_4be = MakeSymbol("shen.be")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symporters = MakeSymbol("porters")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symkill = MakeSymbol("kill")
var symshen_4type_1signature_1of_1y_1or_1n_2 = MakeSymbol("shen.type-signature-of-y-or-n?")
var symmode = MakeSymbol("mode")
var symshen_4type_1signature_1of_1stoutput = MakeSymbol("shen.type-signature-of-stoutput")
var symshen_4s_1prolog__clause = MakeSymbol("shen.s-prolog_clause")
var symshen_4doubleunderline_2 = MakeSymbol("shen.doubleunderline?")
var symshen_4reverse__help = MakeSymbol("shen.reverse_help")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4ebr = MakeSymbol("shen.ebr")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4ue = MakeSymbol("shen.ue")
var sym_f_fgoto_1label = MakeSymbol("%%goto-label")
var symtry_1catch = MakeSymbol("try-catch")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4type_1signature_1of_1explode = MakeSymbol("shen.type-signature-of-explode")
var symStart = MakeSymbol("Start")
var symshen_4cons__form = MakeSymbol("shen.cons_form")
var symshen_4insert_1prolog_1variables_1help = MakeSymbol("shen.insert-prolog-variables-help")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4type_1signature_1of_1if = MakeSymbol("shen.type-signature-of-if")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4add__test = MakeSymbol("shen.add_test")
var symshen_4failed_b = MakeSymbol("shen.failed!")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symtuple_2 = MakeSymbol("tuple?")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4_5st__input_6 = MakeSymbol("shen.<st_input>")
var symshen_4default_1rule = MakeSymbol("shen.default-rule")
var symshen_4type_1signature_1of_1 = MakeSymbol("shen.type-signature-of-")
var symshen_4construct_1search_1clauses = MakeSymbol("shen.construct-search-clauses")
var symResult = MakeSymbol("Result")
var symuntrack = MakeSymbol("untrack")
var symshen_4group__clauses = MakeSymbol("shen.group_clauses")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4toplevel__evaluate = MakeSymbol("shen.toplevel_evaluate")
var symif = MakeSymbol("if")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symlineread = MakeSymbol("lineread")
var symshen_4insert__modes = MakeSymbol("shen.insert_modes")
var symshen_4initialise_1prolog = MakeSymbol("shen.initialise-prolog")
var symshen_4kill_1code = MakeSymbol("shen.kill-code")
var symshen_4variable_1match = MakeSymbol("shen.variable-match")
var symshen_4x_4factorise_1defun_4apply_1selector_1handlers = MakeSymbol("shen.x.factorise-defun.apply-selector-handlers")
var symabsvector = MakeSymbol("absvector")
var symshen_4_5signature_1help_6 = MakeSymbol("shen.<signature-help>")
var sym_8p = MakeSymbol("@p")
var symshen_4type_1signature_1of_1read_1file = MakeSymbol("shen.type-signature-of-read-file")
var symshen_4type_1signature_1of_1reverse = MakeSymbol("shen.type-signature-of-reverse")
var symshen_4the = MakeSymbol("shen.the")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symprofile_1results = MakeSymbol("profile-results")
var symwarn = MakeSymbol("warn")
var symshen_4type_1signature_1of_1read_1byte = MakeSymbol("shen.type-signature-of-read-byte")
var symshen_4_5postdigits_6 = MakeSymbol("shen.<postdigits>")
var symshen_4print_1past_1inputs = MakeSymbol("shen.print-past-inputs")
var sym_5 = MakeSymbol("<")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symshen_4read_7 = MakeSymbol("shen.read+")
var symshen_4_5clauses_d_6 = MakeSymbol("shen.<clauses*>")
var symshen_4cc__help = MakeSymbol("shen.cc_help")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4x_4factorise_1defun_4_dselector_1handlers_d = MakeSymbol("shen.x.factorise-defun.*selector-handlers*")
var symlength = MakeSymbol("length")
var symshen_4hdhd = MakeSymbol("shen.hdhd")
var symshen_4type_1signature_1of_1or = MakeSymbol("shen.type-signature-of-or")
var symshen_4continuation = MakeSymbol("shen.continuation")
var symshen_4unbindv = MakeSymbol("shen.unbindv")
var symshen_4include_1h = MakeSymbol("shen.include-h")
var symdeclare = MakeSymbol("declare")
var symshen_4type_1signature_1of_1absvector_2 = MakeSymbol("shen.type-signature-of-absvector?")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4record_1source = MakeSymbol("shen.record-source")
var symshen_4list__variables = MakeSymbol("shen.list_variables")
var symthaw = MakeSymbol("thaw")
var symshen_4type_1signature_1of_1cd = MakeSymbol("shen.type-signature-of-cd")
var symshen_4compress_150 = MakeSymbol("shen.compress-50")
var symshen_4type_1signature_1of_1arity = MakeSymbol("shen.type-signature-of-arity")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4case_1form = MakeSymbol("shen.case-form")
var symlist = MakeSymbol("list")
var symshen_4type_1signature_1of_1sterror = MakeSymbol("shen.type-signature-of-sterror")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4_5end_d_6 = MakeSymbol("shen.<end*>")
var symshen_4pop = MakeSymbol("shen.pop")
var symshen_4prolog__constant_2 = MakeSymbol("shen.prolog_constant?")
var symshen_4exclamation = MakeSymbol("shen.exclamation")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symstinput = MakeSymbol("stinput")
var syminferences = MakeSymbol("inferences")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4result_1type = MakeSymbol("shen.result-type")
var symshen_4externally_1bound = MakeSymbol("shen.externally-bound")
var sym_f_freturn = MakeSymbol("%%return")
var symshen_4x_4factorise_1defun_4unregister_1selector_1handler = MakeSymbol("shen.x.factorise-defun.unregister-selector-handler")
var symshen_4tests = MakeSymbol("shen.tests")
var symnull = MakeSymbol("null")
var symshen_4package_1macro = MakeSymbol("shen.package-macro")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4by__hypothesis = MakeSymbol("shen.by_hypothesis")
var symshen_4patthyps = MakeSymbol("shen.patthyps")
var symU = MakeSymbol("U")
var symshen_4type_1signature_1of_1length = MakeSymbol("shen.type-signature-of-length")
var symshen_4else = MakeSymbol("shen.else")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4curry_1type_1h = MakeSymbol("shen.curry-type-h")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4read_1char_1code = MakeSymbol("shen.read-char-code")
var symshen_4succeeds_2 = MakeSymbol("shen.succeeds?")
var symshen_4ues = MakeSymbol("shen.ues")
var symshen_4x_4factorise_1defun_4inline_1mono_1labels = MakeSymbol("shen.x.factorise-defun.inline-mono-labels")
var symput = MakeSymbol("put")
var symlambda = MakeSymbol("lambda")
var symshen_4compile_1macro = MakeSymbol("shen.compile-macro")
var symtail = MakeSymbol("tail")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4update__history = MakeSymbol("shen.update_history")
var symshen_4x_4factorise_1defun_4register_1selector_1handler = MakeSymbol("shen.x.factorise-defun.register-selector-handler")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symlimit = MakeSymbol("limit")
var symshen_4_5literal_d_6 = MakeSymbol("shen.<literal*>")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4_5conclusion_6 = MakeSymbol("shen.<conclusion>")
var symshen_4_dcustom_1pattern_1reducer_d = MakeSymbol("shen.*custom-pattern-reducer*")
var symshen_4abstract__rule = MakeSymbol("shen.abstract_rule")
var symshen_4type_1signature_1of_1shen_4app = MakeSymbol("shen.type-signature-of-shen.app")
var symshen_4type_1signature_1of_1_5e_6 = MakeSymbol("shen.type-signature-of-<e>")
var symshen_4collect = MakeSymbol("shen.collect")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symwhere = MakeSymbol("where")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var sym_b = MakeSymbol("!")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var symlazy = MakeSymbol("lazy")
var symbar_b = MakeSymbol("bar!")
var symS = MakeSymbol("S")
var symshen_4lzy_a_a = MakeSymbol("shen.lzy==")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symconcat = MakeSymbol("concat")
var symshen_4type_1signature_1of_1append = MakeSymbol("shen.type-signature-of-append")
var symContinuation = MakeSymbol("Continuation")
var symshen_4mk_1pvar = MakeSymbol("shen.mk-pvar")
var symshen_4decons = MakeSymbol("shen.decons")
var sym_5_1 = MakeSymbol("<-")
var symshen_4cond_1expression = MakeSymbol("shen.cond-expression")
var symshen_4aum__to__shen = MakeSymbol("shen.aum_to_shen")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symmap = MakeSymbol("map")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4type_1signature_1of_1read_1from_1string = MakeSymbol("shen.type-signature-of-read-from-string")
var symshen_4type_1signature_1of_1tlv = MakeSymbol("shen.type-signature-of-tlv")
var symshen_4continuation__call = MakeSymbol("shen.continuation_call")
var symshen_4read_1error = MakeSymbol("shen.read-error")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symstring_1_6n = MakeSymbol("string->n")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symsnd = MakeSymbol("snd")
var symread = MakeSymbol("read")
var symreturn = MakeSymbol("return")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symA = MakeSymbol("A")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4credits = MakeSymbol("shen.credits")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4hdtl = MakeSymbol("shen.hdtl")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4type_1signature_1of_1write_1byte = MakeSymbol("shen.type-signature-of-write-byte")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symB = MakeSymbol("B")
var symsterror = MakeSymbol("sterror")
var symshen_4type_1signature_1of_1internal = MakeSymbol("shen.type-signature-of-internal")
var symshen_4type_1signature_1of_1shen_4proc_1nl = MakeSymbol("shen.type-signature-of-shen.proc-nl")
var symshen_4_5st__input1_6 = MakeSymbol("shen.<st_input1>")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symshen_4type_1signature_1of_1_5_b_6 = MakeSymbol("shen.type-signature-of-<!>")
var symshen_4mu = MakeSymbol("shen.mu")
var symshen_4double = MakeSymbol("shen.double")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4type_1signature_1of_1version = MakeSymbol("shen.type-signature-of-version")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symshen_4snd_1or_1fail = MakeSymbol("shen.snd-or-fail")
var symshen_4_5simple__pattern_6 = MakeSymbol("shen.<simple_pattern>")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var sym_k = MakeSymbol(";")
var symshen_4chwild = MakeSymbol("shen.chwild")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4_5variable_2_6 = MakeSymbol("shen.<variable?>")
var symshen_4shen_1syntax_1error = MakeSymbol("shen.shen-syntax-error")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symshen_4type_1signature_1of_1_d = MakeSymbol("shen.type-signature-of-*")
var symshen_4rule_1_6horn_1clause_1body = MakeSymbol("shen.rule->horn-clause-body")
var symbind = MakeSymbol("bind")
var symshen_4type_1signature_1of_1string_2 = MakeSymbol("shen.type-signature-of-string?")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4_dvarcounter_d = MakeSymbol("shen.*varcounter*")
var symshen_4type_1signature_1of_1systemf = MakeSymbol("shen.type-signature-of-systemf")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4call__the__continuation = MakeSymbol("shen.call_the_continuation")
var symfail = MakeSymbol("fail")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4prolog_1error = MakeSymbol("shen.prolog-error")
var symshen_4hat = MakeSymbol("shen.hat")
var symshen_4recursive__descent = MakeSymbol("shen.recursive_descent")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symG = MakeSymbol("G")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4tlhd = MakeSymbol("shen.tlhd")
var symsave = MakeSymbol("save")
var symshen_4type_1signature_1of_1_5_a = MakeSymbol("shen.type-signature-of-<=")
var symsuccess = MakeSymbol("success")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symdefun = MakeSymbol("defun")
var syminput_7 = MakeSymbol("input+")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4x_4factorise_1defun_4free_1variables = MakeSymbol("shen.x.factorise-defun.free-variables")
var symread_1file = MakeSymbol("read-file")
var symshen_4type_1signature_1of_1error_1to_1string = MakeSymbol("shen.type-signature-of-error-to-string")
var symshen_4cutpoint = MakeSymbol("shen.cutpoint")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4prh = MakeSymbol("shen.prh")
var symshen_4x_4factorise_1defun_4generate_1label = MakeSymbol("shen.x.factorise-defun.generate-label")
var sym_6 = MakeSymbol(">")
var symshen_4fail__if = MakeSymbol("shen.fail_if")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4lookup_1func = MakeSymbol("shen.lookup-func")
var symlet = MakeSymbol("let")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4single = MakeSymbol("shen.single")
var symshen_4elim_1def = MakeSymbol("shen.elim-def")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4check__stream = MakeSymbol("shen.check_stream")
var symshen_4_dprologvectors_d = MakeSymbol("shen.*prologvectors*")
var symshen_4sequent = MakeSymbol("shen.sequent")
var symshen_4mult__subst = MakeSymbol("shen.mult_subst")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symvalue = MakeSymbol("value")
var sym_1 = MakeSymbol("-")
var symfindall = MakeSymbol("findall")
var symshen_4type_1signature_1of_1close = MakeSymbol("shen.type-signature-of-close")
var symshen_4remove__modes = MakeSymbol("shen.remove_modes")
var symshen_4aum = MakeSymbol("shen.aum")
var symshen_4construct_1side_1literals = MakeSymbol("shen.construct-side-literals")
var symRecord = MakeSymbol("Record")
var symreverse = MakeSymbol("reverse")
var symunit = MakeSymbol("unit")
var symshen_4type_1signature_1of_1specialise = MakeSymbol("shen.type-signature-of-specialise")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4void = MakeSymbol("shen.void")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs = MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4free__variable__check = MakeSymbol("shen.free_variable_check")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symJ = MakeSymbol("J")
var symF = MakeSymbol("F")
var symshen_4type_1signature_1of_1track = MakeSymbol("shen.type-signature-of-track")
var symshen_4rename = MakeSymbol("shen.rename")
var symshen_4typedf_2 = MakeSymbol("shen.typedf?")
var symshen_4eval_1without_1macros = MakeSymbol("shen.eval-without-macros")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4initialise__arity__table = MakeSymbol("shen.initialise_arity_table")
var symshen_4type_1signature_1of_1port = MakeSymbol("shen.type-signature-of-port")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4type_1signature_1of_1symbol_2 = MakeSymbol("shen.type-signature-of-symbol?")
var symarity = MakeSymbol("arity")
var symtype = MakeSymbol("type")
var symmake_1string = MakeSymbol("make-string")
var symshen_4type_1signature_1of_1mapcan = MakeSymbol("shen.type-signature-of-mapcan")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4decimalise = MakeSymbol("shen.decimalise")
var symshen_4x_4factorise_1defun_4concat_c = MakeSymbol("shen.x.factorise-defun.concat/")
var symprofile = MakeSymbol("profile")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4_5anysingle_6 = MakeSymbol("shen.<anysingle>")
var symnumber_2 = MakeSymbol("number?")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4demod_1rule = MakeSymbol("shen.demod-rule")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4type_1signature_1of_1_6_a = MakeSymbol("shen.type-signature-of->=")
var symshen_4semantics = MakeSymbol("shen.semantics")
var symshen_4remove_1nth = MakeSymbol("shen.remove-nth")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symshen_4intprolog_1help_1help = MakeSymbol("shen.intprolog-help-help")
var symtl = MakeSymbol("tl")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symunput = MakeSymbol("unput")
var symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d = MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*")
var symstring = MakeSymbol("string")
var symshen_4type_1signature_1of_1shen_4prhush = MakeSymbol("shen.type-signature-of-shen.prhush")
var symshen_4type_1signature_1of_1write_1to_1file = MakeSymbol("shen.type-signature-of-write-to-file")
var symshen_4typecheck_1and_1load = MakeSymbol("shen.typecheck-and-load")
var symshen_4numbyte_2 = MakeSymbol("shen.numbyte?")
var symlanguage = MakeSymbol("language")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symshen_4type_1signature_1of_1tc_2 = MakeSymbol("shen.type-signature-of-tc?")
var symshen_4grammar__symbol_2 = MakeSymbol("shen.grammar_symbol?")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symexception = MakeSymbol("exception")
var symshen_4prbytes = MakeSymbol("shen.prbytes")
var symshen_4x_4factorise_1defun_4with_1labelled_1else = MakeSymbol("shen.x.factorise-defun.with-labelled-else")
var symunion = MakeSymbol("union")
var symtime = MakeSymbol("time")
var symshen_4type_1signature_1of_1os = MakeSymbol("shen.type-signature-of-os")
var symshen_4construct_1recursive_1search_1clause = MakeSymbol("shen.construct-recursive-search-clause")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4th_d = MakeSymbol("shen.th*")
var symshen_4_5action_6 = MakeSymbol("shen.<action>")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4_5anymulti_6 = MakeSymbol("shen.<anymulti>")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4remember = MakeSymbol("shen.remember")
var symadjoin = MakeSymbol("adjoin")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4flatten = MakeSymbol("shen.flatten")
var symshen_4abstraction__build = MakeSymbol("shen.abstraction_build")
var symabort = MakeSymbol("abort")
var symshen_4type_1signature_1of_1function = MakeSymbol("shen.type-signature-of-function")
var symshen_4type_1signature_1of_1trap_1error = MakeSymbol("shen.type-signature-of-trap-error")
var symshen_4result = MakeSymbol("shen.result")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4sys_1error = MakeSymbol("shen.sys-error")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symshen_4_dmacroreg_d = MakeSymbol("shen.*macroreg*")
var symshen_4type_1signature_1of_1fail = MakeSymbol("shen.type-signature-of-fail")
var symshen_4safe_1product = MakeSymbol("shen.safe-product")
var symshen_4call_1rest = MakeSymbol("shen.call-rest")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symtlv = MakeSymbol("tlv")
var symO = MakeSymbol("O")
var symL = MakeSymbol("L")
var sym_1_6 = MakeSymbol("->")
var symW = MakeSymbol("W")
var symshen_4newline = MakeSymbol("shen.newline")
var symshen_4x_4factorise_1defun_4optimize_1selectors = MakeSymbol("shen.x.factorise-defun.optimize-selectors")
var symelement_2 = MakeSymbol("element?")
var symhash = MakeSymbol("hash")
var symstep = MakeSymbol("step")
var symshen_4assumetype = MakeSymbol("shen.assumetype")
var symshen_4decons_1string = MakeSymbol("shen.decons-string")
var symContextOut__1957 = MakeSymbol("ContextOut_1957")
var symshen_4toplineread = MakeSymbol("shen.toplineread")
var symopen = MakeSymbol("open")
var symshen_4_5premises_6 = MakeSymbol("shen.<premises>")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4x_4factorise_1defun_4bind_1repeating_1selectors = MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symrelease = MakeSymbol("release")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4type_1signature_1of_1vector_2 = MakeSymbol("shen.type-signature-of-vector?")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4construct_1search_1clause = MakeSymbol("shen.construct-search-clause")
var symshen_4interror = MakeSymbol("shen.interror")
var symshen_4type_1signature_1of_1assoc = MakeSymbol("shen.type-signature-of-assoc")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4t_d_1patterns = MakeSymbol("shen.t*-patterns")
var symdefine = MakeSymbol("define")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4x_4factorise_1defun_4true_1branch = MakeSymbol("shen.x.factorise-defun.true-branch")
var symboolean = MakeSymbol("boolean")
var symshen_4variables = MakeSymbol("shen.variables")
var symshen_4update_1demodulation_1function = MakeSymbol("shen.update-demodulation-function")
var symshen_4list_1stream = MakeSymbol("shen.list-stream")
var sym_7 = MakeSymbol("+")
var symshen_4next_150 = MakeSymbol("shen.next-50")
var symIn__1957 = MakeSymbol("In_1957")
var symThrowcontrol = MakeSymbol("Throwcontrol")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symshen_4x_4features_4_dfeatures_d = MakeSymbol("shen.x.features.*features*")
var symshen_4type_1signature_1of_1hash = MakeSymbol("shen.type-signature-of-hash")
var symAssumptions__1957 = MakeSymbol("Assumptions_1957")
var symshen_4_dinstalling_1kl_d = MakeSymbol("shen.*installing-kl*")
var symnth = MakeSymbol("nth")
var symshen_4type_1signature_1of_1str = MakeSymbol("shen.type-signature-of-str")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4aritycheck_1name = MakeSymbol("shen.aritycheck-name")
var sympreclude = MakeSymbol("preclude")
var symshen_4type_1signature_1of_1adjoin = MakeSymbol("shen.type-signature-of-adjoin")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4yacc = MakeSymbol("shen.yacc")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symfwhen = MakeSymbol("fwhen")
var sym_e_e = MakeSymbol("&&")
var symeval = MakeSymbol("eval")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4proc_1input_7 = MakeSymbol("shen.proc-input+")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4type_1signature_1of_1profile = MakeSymbol("shen.type-signature-of-profile")
var symshen_4right_1_6left = MakeSymbol("shen.right->left")
var symshen_4construct_1search_1literals = MakeSymbol("shen.construct-search-literals")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4_dmaxcomplexity_d = MakeSymbol("shen.*maxcomplexity*")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var sym_6_a = MakeSymbol(">=")
var symshen_4type_1signature_1of_1cn = MakeSymbol("shen.type-signature-of-cn")
var symshen_4synonyms_1help = MakeSymbol("shen.synonyms-help")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4find_1past_1inputs = MakeSymbol("shen.find-past-inputs")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4type_1signature_1of_1string_1_6n = MakeSymbol("shen.type-signature-of-string->n")
var symshen_4type_1signature_1of_1_5 = MakeSymbol("shen.type-signature-of-<")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var symshen_4_5comma_1symbol_6 = MakeSymbol("shen.<comma-symbol>")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4analyse_1kill = MakeSymbol("shen.analyse-kill")
var symprotect = MakeSymbol("protect")
var symE = MakeSymbol("E")
var symshen_4_dalphabet_d = MakeSymbol("shen.*alphabet*")
var symshen_4rule_1_6horn_1clause = MakeSymbol("shen.rule->horn-clause")
var symexternal = MakeSymbol("external")
var symshen_4deref = MakeSymbol("shen.deref")
var symappend = MakeSymbol("append")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4type_1signature_1of_1kill = MakeSymbol("shen.type-signature-of-kill")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symTime = MakeSymbol("Time")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4strip_1pathname = MakeSymbol("shen.strip-pathname")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4add_1macro = MakeSymbol("shen.add-macro")
var symshen_4s_1prolog = MakeSymbol("shen.s-prolog")
var symshen_4right_1rule = MakeSymbol("shen.right-rule")
var symV = MakeSymbol("V")
var symabsvector_2 = MakeSymbol("absvector?")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4carriage_1return = MakeSymbol("shen.carriage-return")
var sym_e_e_e = MakeSymbol("&&&")
var symshen_4jump__stream_2 = MakeSymbol("shen.jump_stream?")
var symcons = MakeSymbol("cons")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4explicit__modes = MakeSymbol("shen.explicit_modes")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symParse__ = MakeSymbol("Parse_")
var symshen_4x_4factorise_1defun_4exp_1var = MakeSymbol("shen.x.factorise-defun.exp-var")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4yacc__cases = MakeSymbol("shen.yacc_cases")
var symfile = MakeSymbol("file")
var symshen_4type_1signature_1of_1untrack = MakeSymbol("shen.type-signature-of-untrack")
var symshen_4type_1signature_1of_1_c = MakeSymbol("shen.type-signature-of-/")
var symshen_4intprolog_1help = MakeSymbol("shen.intprolog-help")
var symshen_4read_1file_1as_1Xlist = MakeSymbol("shen.read-file-as-Xlist")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symvector = MakeSymbol("vector")
var symnumber = MakeSymbol("number")
var symshen_4type_1signature_1of_1release = MakeSymbol("shen.type-signature-of-release")
var symshen_4start_1new_1prolog_1process = MakeSymbol("shen.start-new-prolog-process")
var symshen_4ephemeral__variable_2 = MakeSymbol("shen.ephemeral_variable?")
var symshen_4t_d_1def = MakeSymbol("shen.t*-def")
var symshen_4aritycheck_1action = MakeSymbol("shen.aritycheck-action")
var symtlstr = MakeSymbol("tlstr")
var symQ = MakeSymbol("Q")
var sympackage_2 = MakeSymbol("package?")
var symshen_4type_1signature_1of_1map = MakeSymbol("shen.type-signature-of-map")
var symshen_4nextline = MakeSymbol("shen.nextline")
var sympackage = MakeSymbol("package")
var symshen_4compile__to__lambda_7 = MakeSymbol("shen.compile_to_lambda+")
var symshen_4_dteststack_d = MakeSymbol("shen.*teststack*")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symtc = MakeSymbol("tc")
var symshen_4copy_1vector_1stage_11 = MakeSymbol("shen.copy-vector-stage-1")
var symYaccParse = MakeSymbol("YaccParse")
var symshen_4x_4factorise_1defun_4initialise = MakeSymbol("shen.x.factorise-defun.initialise")
var sym_dsterror_d = MakeSymbol("*sterror*")
var symshen_4type_1signature_1of_1hdv = MakeSymbol("shen.type-signature-of-hdv")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4split__cc__rules = MakeSymbol("shen.split_cc_rules")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4_5clause_d_6 = MakeSymbol("shen.<clause*>")
var symshen_4complexity = MakeSymbol("shen.complexity")
var symshen_4prolog_1aritycheck = MakeSymbol("shen.prolog-aritycheck")
var symhdstr = MakeSymbol("hdstr")
var symshen_4eval_1cons = MakeSymbol("shen.eval-cons")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4type_1signature_1of_1fix = MakeSymbol("shen.type-signature-of-fix")
var symshen_4function_1abstraction_1help = MakeSymbol("shen.function-abstraction-help")
var symshen_4x_4factorise_1defun_4bind_1selector = MakeSymbol("shen.x.factorise-defun.bind-selector")
var symload = MakeSymbol("load")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symset = MakeSymbol("set")
var sym_c = MakeSymbol("/")
var symshen_4type_1signature_1of_1vector = MakeSymbol("shen.type-signature-of-vector")
var symshen_4function_1abstraction = MakeSymbol("shen.function-abstraction")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symand = MakeSymbol("and")
var symfix = MakeSymbol("fix")
var symshen_4same__predicate_2 = MakeSymbol("shen.same_predicate?")
var symshen_4type_1signature_1of_1compile = MakeSymbol("shen.type-signature-of-compile")
var symshen_4type_1signature_1of_1element_2 = MakeSymbol("shen.type-signature-of-element?")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symshen_4x_4features_4add = MakeSymbol("shen.x.features.add")
var symos = MakeSymbol("os")
var symshen_4sh_2 = MakeSymbol("shen.sh?")
var symOut__1957 = MakeSymbol("Out_1957")
var symvector_2 = MakeSymbol("vector?")
var symshen_4type_1signature_1of_1tuple_2 = MakeSymbol("shen.type-signature-of-tuple?")
var symshen_4symbol_1code_2 = MakeSymbol("shen.symbol-code?")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symout = MakeSymbol("out")
var symshen_4type_1signature_1of_1ps = MakeSymbol("shen.type-signature-of-ps")
var symshen_4locally_1bind_1prolog_1vs = MakeSymbol("shen.locally-bind-prolog-vs")
var symcall = MakeSymbol("call")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4base = MakeSymbol("shen.base")
var symshen_4aah = MakeSymbol("shen.aah")
var symshen_4_5head_d_6 = MakeSymbol("shen.<head*>")
var symshen_4findallhelp = MakeSymbol("shen.findallhelp")
var symshen_4type_1signature_1of_1string_1_6symbol = MakeSymbol("shen.type-signature-of-string->symbol")
var symshen_4findpos = MakeSymbol("shen.findpos")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4x_4factorise_1defun_4done = MakeSymbol("shen.x.factorise-defun.done")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symprint = MakeSymbol("print")
var symshen_4x_4factorise_1defun_4rebranch = MakeSymbol("shen.x.factorise-defun.rebranch")
var sym_a_b = MakeSymbol("=!")
var sym_h_a = MakeSymbol(":=")
var symshen_4type_1signature_1of_1external = MakeSymbol("shen.type-signature-of-external")
var symshen_4variable = MakeSymbol("shen.variable")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symsum = MakeSymbol("sum")
var symshen_4type_1signature_1of_1get_1time = MakeSymbol("shen.type-signature-of-get-time")
var symshen_4get_1type = MakeSymbol("shen.get-type")
var symshen_4clash_2 = MakeSymbol("shen.clash?")
var symshen_4type_1signature_1of_1load = MakeSymbol("shen.type-signature-of-load")
var symshen_4safe_1multiply = MakeSymbol("shen.safe-multiply")
var symshen_4lineread_1loop = MakeSymbol("shen.lineread-loop")
var symshen_4t_d_1defhh = MakeSymbol("shen.t*-defhh")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4type_1signature_1of_1thaw = MakeSymbol("shen.type-signature-of-thaw")
var symshen_4type_1signature_1of_1_7 = MakeSymbol("shen.type-signature-of-+")
var symshen_4trim_1gubbins = MakeSymbol("shen.trim-gubbins")
var symsystemf = MakeSymbol("systemf")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symwhen = MakeSymbol("when")
var symshen_4type_1signature_1of_1occurs_1check = MakeSymbol("shen.type-signature-of-occurs-check")
var symshen_4type_1signature_1of_1tail = MakeSymbol("shen.type-signature-of-tail")
var symshen_4bld_1prolog_1call = MakeSymbol("shen.bld-prolog-call")
var symshen_4clauses_1to_1shen = MakeSymbol("shen.clauses-to-shen")
var symshen_4f = MakeSymbol("shen.f")
var symNewStream = MakeSymbol("NewStream")
var symshen_4type_1signature_1of_1cons_2 = MakeSymbol("shen.type-signature-of-cons?")
var symshen_4variancy_1test = MakeSymbol("shen.variancy-test")
var symStream = MakeSymbol("Stream")
var symN = MakeSymbol("N")
var symoptimise = MakeSymbol("optimise")
var symcut = MakeSymbol("cut")
var symshen_4nextbyte = MakeSymbol("shen.nextbyte")
var symshen_4space = MakeSymbol("shen.space")
var symshen_4x_4features_4cond_1expand = MakeSymbol("shen.x.features.cond-expand")
var symdefcc = MakeSymbol("defcc")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4make__mu__application = MakeSymbol("shen.make_mu_application")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symdifference = MakeSymbol("difference")
var symshen_4ue_1sig = MakeSymbol("shen.ue-sig")
var symshen_4compile__to__machine__code = MakeSymbol("shen.compile_to_machine_code")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4typextable = MakeSymbol("shen.typextable")
var symshen_4insert_1lazyderef = MakeSymbol("shen.insert-lazyderef")
var symshen_4valvector = MakeSymbol("shen.valvector")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symassoc = MakeSymbol("assoc")
var symcompile = MakeSymbol("compile")
var symshen_4extract__vars = MakeSymbol("shen.extract_vars")
var symshen_4_5sig_7rest_6 = MakeSymbol("shen.<sig+rest>")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4ue_1h_2 = MakeSymbol("shen.ue-h?")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symloaded = MakeSymbol("loaded")
var symshen_4type_1signature_1of_1protect = MakeSymbol("shen.type-signature-of-protect")
var symshen_4type_1signature_1of_1_a = MakeSymbol("shen.type-signature-of-=")
var symshen_4copy_1vector_1stage_12 = MakeSymbol("shen.copy-vector-stage-2")
var symshen_4prolog_1failure = MakeSymbol("shen.prolog-failure")
var syminput = MakeSymbol("input")
var symclose = MakeSymbol("close")
var symshen_4type_1signature_1of_1and = MakeSymbol("shen.type-signature-of-and")
var symshen_4type_1signature_1of_1inferences = MakeSymbol("shen.type-signature-of-inferences")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var syminclude = MakeSymbol("include")
var symshen_4udefs_d = MakeSymbol("shen.udefs*")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4strip_1protect = MakeSymbol("shen.strip-protect")
var symZ = MakeSymbol("Z")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4resize_1vector = MakeSymbol("shen.resize-vector")
var symshen_4toplineread__loop = MakeSymbol("shen.toplineread_loop")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4insert_1runon = MakeSymbol("shen.insert-runon")
var symstoutput = MakeSymbol("stoutput")
var symshen_4measure_ereturn = MakeSymbol("shen.measure&return")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symns2_1set = MakeSymbol("ns2-set")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symdo = MakeSymbol("do")
var symimplementation = MakeSymbol("implementation")
var symtrack = MakeSymbol("track")
var symshen_4digits_1_6integers = MakeSymbol("shen.digits->integers")
var symshen_4placeholder = MakeSymbol("shen.placeholder")
var symhead = MakeSymbol("head")
var symidentical = MakeSymbol("identical")
var symshen_4type_1signature_1of_1intersection = MakeSymbol("shen.type-signature-of-intersection")
var symshen_4type_1signature_1of_1preclude = MakeSymbol("shen.type-signature-of-preclude")
var symshen_4remove_1bar = MakeSymbol("shen.remove-bar")
var symshen_4toplevel = MakeSymbol("shen.toplevel")
var symshen_4make_1key = MakeSymbol("shen.make-key")
var symshen_4type_1signature_1of_1read_1file_1as_1string = MakeSymbol("shen.type-signature-of-read-file-as-string")
var symshen_4singleunderline_2 = MakeSymbol("shen.singleunderline?")
var symshen_4rules_1_6horn_1clauses = MakeSymbol("shen.rules->horn-clauses")
var symshen_4retrieve_1from_1history_1if_1needed = MakeSymbol("shen.retrieve-from-history-if-needed")
