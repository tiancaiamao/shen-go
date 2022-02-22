package main

import . "github.com/tiancaiamao/cora/cora_go"

var TypesMain = MakeNative(func(__e *ControlFlow) {
	tmp16527 := MakeNative(func(__e *ControlFlow) {
		V4662 := __e.Get(1)
		_ = V4662
		V4663 := __e.Get(2)
		_ = V4663
		tmp16528 := MakeNative(func(__e *ControlFlow) {
			Rectify := __e.Get(1)
			_ = Rectify
			tmp16529 := MakeNative(func(__e *ControlFlow) {
				Variant_2 := __e.Get(1)
				_ = Variant_2
				tmp16530 := MakeNative(func(__e *ControlFlow) {
					Abstraction := __e.Get(1)
					_ = Abstraction
					tmp16531 := MakeNative(func(__e *ControlFlow) {
						UpDate := __e.Get(1)
						_ = UpDate
						__e.Return(V4662)
						return
					}, 1)

					tmp16532 := PrimNS3Value(symshen_4_dsigf_d)

					tmp16533 := Call(__e, PrimNS2Value(symshen_4assoc_1_6), V4662, Abstraction, tmp16532)

					tmp16534 := PrimNS3Set(symshen_4_dsigf_d, tmp16533)

					__e.TailApply(tmp16531, tmp16534)
					return

				}, 1)

				tmp16535 := Call(__e, PrimNS2Value(symshen_4prolog_1abstraction), V4663)

				tmp16536 := Call(__e, PrimNS2Value(symeval_1kl), tmp16535)

				__e.TailApply(tmp16530, tmp16536)
				return

			}, 1)

			tmp16537 := MakeNative(func(__e *ControlFlow) {
				V4624 := __e.Get(1)
				_ = V4624
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L4625 := __e.Get(1)
					_ = L4625
					__e.Return(MakeNative(func(__e *ControlFlow) {
						K4626 := __e.Get(1)
						_ = K4626
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C4627 := __e.Get(1)
							_ = C4627
							tmp16538 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp16538

							tmp16539 := Call(__e, PrimNS2Value(symshen_4deref), V4662, V4624)

							tmp16540 := Call(__e, PrimNS2Value(symreceive), tmp16539)

							tmp16541 := Call(__e, PrimNS2Value(symshen_4deref), Rectify, V4624)

							tmp16542 := Call(__e, PrimNS2Value(symreceive), tmp16541)

							__e.TailApply(PrimNS2Value(symshen_4variancy), tmp16540, tmp16542, V4624, L4625, K4626, C4627)
							return

						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)

			tmp16543 := Call(__e, PrimNS2Value(symshen_4reset_1prolog_1vector))

			tmp16544 := Call(__e, tmp16537, tmp16543)

			tmp16545 := Call(__e, PrimNS2Value(symvector), MakeNumber(0))

			tmp16546 := Call(__e, PrimNS2Value(sym_8v), MakeNumber(0), tmp16545)

			tmp16547 := Call(__e, PrimNS2Value(sym_8v), True, tmp16546)

			tmp16548 := Call(__e, tmp16544, tmp16547)

			tmp16549 := Call(__e, tmp16548, MakeNumber(0))

			tmp16550 := MakeNative(func(__e *ControlFlow) {
				__e.Return(True)
				return
			}, 0)

			tmp16551 := Call(__e, tmp16549, tmp16550)

			__e.TailApply(tmp16529, tmp16551)
			return

		}, 1)

		tmp16552 := Call(__e, PrimNS2Value(symshen_4rectify_1type), V4663)

		__e.TailApply(tmp16528, tmp16552)
		return

	}, 2)

	tmp16553 := Call(__e, PrimNS2Value(symdef), symdeclare, tmp16527)

	_ = tmp16553

	tmp16554 := MakeNative(func(__e *ControlFlow) {
		V4664 := __e.Get(1)
		_ = V4664
		V4665 := __e.Get(2)
		_ = V4665
		V4666 := __e.Get(3)
		_ = V4666
		V4667 := __e.Get(4)
		_ = V4667
		V4668 := __e.Get(5)
		_ = V4668
		V4669 := __e.Get(6)
		_ = V4669
		tmp16561 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4667)

		if True == tmp16561 {
			tmp16556 := MakeNative(func(__e *ControlFlow) {
				A := __e.Get(1)
				_ = A
				tmp16557 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp16557

				tmp16558 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimNS2Value(symshen_4variants_2), V4664, A, V4665, V4666, V4667, V4668, V4669)
					return
				}, 0)

				tmp16559 := Call(__e, PrimNS2Value(symshen_4system_1S_1h), V4664, A, Nil, V4666, V4667, V4668, tmp16558)

				__e.TailApply(PrimNS2Value(symshen_4gc), V4666, tmp16559)
				return

			}, 1)

			tmp16560 := Call(__e, PrimNS2Value(symshen_4newpv), V4666)

			__e.TailApply(tmp16556, tmp16560)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 6)

	tmp16562 := Call(__e, PrimNS2Value(symdef), symshen_4variancy, tmp16554)

	_ = tmp16562

	tmp16563 := MakeNative(func(__e *ControlFlow) {
		V4670 := __e.Get(1)
		_ = V4670
		V4671 := __e.Get(2)
		_ = V4671
		V4672 := __e.Get(3)
		_ = V4672
		V4673 := __e.Get(4)
		_ = V4673
		V4674 := __e.Get(5)
		_ = V4674
		V4675 := __e.Get(6)
		_ = V4675
		V4676 := __e.Get(7)
		_ = V4676
		tmp16564 := MakeNative(func(__e *ControlFlow) {
			K4637 := __e.Get(1)
			_ = K4637
			tmp16565 := MakeNative(func(__e *ControlFlow) {
				C4642 := __e.Get(1)
				_ = C4642
				tmp16589 := PrimEqual(C4642, False)

				if True == tmp16589 {
					tmp16567 := MakeNative(func(__e *ControlFlow) {
						C4645 := __e.Get(1)
						_ = C4645
						tmp16584 := PrimEqual(C4645, False)

						if True == tmp16584 {
							tmp16569 := MakeNative(func(__e *ControlFlow) {
								C4646 := __e.Get(1)
								_ = C4646
								tmp16571 := PrimEqual(C4646, False)

								if True == tmp16571 {
									__e.TailApply(PrimNS2Value(symshen_4unlock), V4674, K4637)
									return
								} else {
									__e.Return(C4646)
									return
								}

							}, 1)

							tmp16583 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4674)

							var ifres16572 Obj

							if True == tmp16583 {
								tmp16573 := MakeNative(func(__e *ControlFlow) {
									Warning := __e.Get(1)
									_ = Warning
									tmp16574 := Call(__e, PrimNS2Value(symshen_4incinfs))

									_ = tmp16574

									tmp16575 := Call(__e, PrimNS2Value(symshen_4deref), V4670, V4673)

									tmp16576 := Call(__e, PrimNS2Value(symshen_4app), tmp16575, MakeString(" may create errors\n"), symshen_4a)

									tmp16577 := PrimStringConcat(MakeString("warning: changing the type of "), tmp16576)

									tmp16578 := Call(__e, PrimNS2Value(symstoutput))

									tmp16579 := Call(__e, PrimNS2Value(sympr), tmp16577, tmp16578)

									tmp16580 := Call(__e, PrimNS2Value(symis), Warning, tmp16579, V4673, V4674, K4637, V4676)

									__e.TailApply(PrimNS2Value(symshen_4gc), V4673, tmp16580)
									return

								}, 1)

								tmp16581 := Call(__e, PrimNS2Value(symshen_4newpv), V4673)

								tmp16582 := Call(__e, tmp16573, tmp16581)

								ifres16572 = tmp16582

							} else {
								ifres16572 = False

							}

							__e.TailApply(tmp16569, ifres16572)
							return

						} else {
							__e.Return(C4645)
							return
						}

					}, 1)

					tmp16588 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4674)

					var ifres16585 Obj

					if True == tmp16588 {
						tmp16586 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp16586

						tmp16587 := Call(__e, PrimNS2Value(symis_b), V4671, V4672, V4673, V4674, K4637, V4676)

						ifres16585 = tmp16587

					} else {
						ifres16585 = False

					}

					__e.TailApply(tmp16567, ifres16585)
					return

				} else {
					__e.Return(C4642)
					return
				}

			}, 1)

			tmp16601 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4674)

			var ifres16590 Obj

			if True == tmp16601 {
				tmp16591 := MakeNative(func(__e *ControlFlow) {
					Tm4643 := __e.Get(1)
					_ = Tm4643
					tmp16592 := MakeNative(func(__e *ControlFlow) {
						GoTo4644 := __e.Get(1)
						_ = GoTo4644
						tmp16596 := PrimEqual(Tm4643, symsymbol)

						if True == tmp16596 {
							__e.TailApply(PrimNS2Value(symthaw), GoTo4644)
							return
						} else {
							tmp16595 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm4643)

							if True == tmp16595 {
								__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm4643, symsymbol, V4673, GoTo4644)
								return
							} else {
								__e.Return(False)
								return
							}

						}

					}, 1)

					tmp16597 := MakeNative(func(__e *ControlFlow) {
						tmp16598 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp16598

						__e.TailApply(PrimNS2Value(symshen_4cut), V4673, V4674, K4637, V4676)
						return

					}, 0)

					__e.TailApply(tmp16592, tmp16597)
					return

				}, 1)

				tmp16599 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4671, V4673)

				tmp16600 := Call(__e, tmp16591, tmp16599)

				ifres16590 = tmp16600

			} else {
				ifres16590 = False

			}

			__e.TailApply(tmp16565, ifres16590)
			return

		}, 1)

		tmp16602 := PrimNumberAdd(V4675, MakeNumber(1))

		__e.TailApply(tmp16564, tmp16602)
		return

	}, 7)

	tmp16603 := Call(__e, PrimNS2Value(symdef), symshen_4variants_2, tmp16563)

	_ = tmp16603

	tmp16604 := MakeNative(func(__e *ControlFlow) {
		V4677 := __e.Get(1)
		_ = V4677
		tmp16605 := MakeNative(func(__e *ControlFlow) {
			Bindings := __e.Get(1)
			_ = Bindings
			tmp16606 := MakeNative(func(__e *ControlFlow) {
				Lock := __e.Get(1)
				_ = Lock
				tmp16607 := MakeNative(func(__e *ControlFlow) {
					Key := __e.Get(1)
					_ = Key
					tmp16608 := MakeNative(func(__e *ControlFlow) {
						Continuation := __e.Get(1)
						_ = Continuation
						tmp16609 := MakeNative(func(__e *ControlFlow) {
							V := __e.Get(1)
							_ = V
							tmp16610 := MakeNative(func(__e *ControlFlow) {
								Vs := __e.Get(1)
								_ = Vs
								tmp16611 := Call(__e, PrimNS2Value(symshen_4rcons__form), V4677)

								tmp16612 := PrimCons(Continuation, Nil)

								tmp16613 := PrimCons(Key, tmp16612)

								tmp16614 := PrimCons(Lock, tmp16613)

								tmp16615 := PrimCons(Bindings, tmp16614)

								tmp16616 := PrimCons(tmp16611, tmp16615)

								tmp16617 := PrimCons(V, tmp16616)

								tmp16618 := PrimCons(symis_b, tmp16617)

								tmp16619 := Call(__e, PrimNS2Value(symshen_4stpart), Vs, tmp16618, Bindings)

								tmp16620 := PrimCons(tmp16619, Nil)

								tmp16621 := PrimCons(Continuation, tmp16620)

								tmp16622 := PrimCons(symlambda, tmp16621)

								tmp16623 := PrimCons(tmp16622, Nil)

								tmp16624 := PrimCons(Key, tmp16623)

								tmp16625 := PrimCons(symlambda, tmp16624)

								tmp16626 := PrimCons(tmp16625, Nil)

								tmp16627 := PrimCons(Lock, tmp16626)

								tmp16628 := PrimCons(symlambda, tmp16627)

								tmp16629 := PrimCons(tmp16628, Nil)

								tmp16630 := PrimCons(Bindings, tmp16629)

								tmp16631 := PrimCons(symlambda, tmp16630)

								tmp16632 := PrimCons(tmp16631, Nil)

								tmp16633 := PrimCons(V, tmp16632)

								__e.Return(PrimCons(symlambda, tmp16633))
								return

							}, 1)

							tmp16634 := Call(__e, PrimNS2Value(symshen_4extract_1vars), V4677)

							__e.TailApply(tmp16610, tmp16634)
							return

						}, 1)

						tmp16635 := Call(__e, PrimNS2Value(symprotect), symV)

						tmp16636 := Call(__e, PrimNS2Value(symgensym), tmp16635)

						__e.TailApply(tmp16609, tmp16636)
						return

					}, 1)

					tmp16637 := Call(__e, PrimNS2Value(symprotect), symC)

					tmp16638 := Call(__e, PrimNS2Value(symgensym), tmp16637)

					__e.TailApply(tmp16608, tmp16638)
					return

				}, 1)

				tmp16639 := Call(__e, PrimNS2Value(symprotect), symKey)

				tmp16640 := Call(__e, PrimNS2Value(symgensym), tmp16639)

				__e.TailApply(tmp16607, tmp16640)
				return

			}, 1)

			tmp16641 := Call(__e, PrimNS2Value(symprotect), symL)

			tmp16642 := Call(__e, PrimNS2Value(symgensym), tmp16641)

			__e.TailApply(tmp16606, tmp16642)
			return

		}, 1)

		tmp16643 := Call(__e, PrimNS2Value(symprotect), symB)

		tmp16644 := Call(__e, PrimNS2Value(symgensym), tmp16643)

		__e.TailApply(tmp16605, tmp16644)
		return

	}, 1)

	tmp16645 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1abstraction, tmp16604)

	_ = tmp16645

	tmp16646 := MakeNative(func(__e *ControlFlow) {
		V4678 := __e.Get(1)
		_ = V4678
		__e.Return(V4678)
		return
	}, 1)

	tmp16647 := Call(__e, PrimNS2Value(symdef), symshen_4demod, tmp16646)

	_ = tmp16647

	tmp16648 := PrimCons(symA, Nil)

	tmp16649 := PrimCons(sym_1_1_6, tmp16648)

	tmp16650 := Call(__e, PrimNS2Value(symdeclare), symabort, tmp16649)

	_ = tmp16650

	tmp16651 := PrimCons(symboolean, Nil)

	tmp16652 := PrimCons(sym_1_1_6, tmp16651)

	tmp16653 := PrimCons(symA, tmp16652)

	tmp16654 := Call(__e, PrimNS2Value(symdeclare), symabsvector_2, tmp16653)

	_ = tmp16654

	tmp16655 := PrimCons(symA, Nil)

	tmp16656 := PrimCons(symlist, tmp16655)

	tmp16657 := PrimCons(symA, Nil)

	tmp16658 := PrimCons(symlist, tmp16657)

	tmp16659 := PrimCons(tmp16658, Nil)

	tmp16660 := PrimCons(sym_1_1_6, tmp16659)

	tmp16661 := PrimCons(tmp16656, tmp16660)

	tmp16662 := PrimCons(tmp16661, Nil)

	tmp16663 := PrimCons(sym_1_1_6, tmp16662)

	tmp16664 := PrimCons(symA, tmp16663)

	tmp16665 := Call(__e, PrimNS2Value(symdeclare), symadjoin, tmp16664)

	_ = tmp16665

	tmp16666 := PrimCons(symboolean, Nil)

	tmp16667 := PrimCons(sym_1_1_6, tmp16666)

	tmp16668 := PrimCons(symboolean, tmp16667)

	tmp16669 := PrimCons(tmp16668, Nil)

	tmp16670 := PrimCons(sym_1_1_6, tmp16669)

	tmp16671 := PrimCons(symboolean, tmp16670)

	tmp16672 := Call(__e, PrimNS2Value(symdeclare), symand, tmp16671)

	_ = tmp16672

	tmp16673 := PrimCons(symstring, Nil)

	tmp16674 := PrimCons(sym_1_1_6, tmp16673)

	tmp16675 := PrimCons(symsymbol, tmp16674)

	tmp16676 := PrimCons(tmp16675, Nil)

	tmp16677 := PrimCons(sym_1_1_6, tmp16676)

	tmp16678 := PrimCons(symstring, tmp16677)

	tmp16679 := PrimCons(tmp16678, Nil)

	tmp16680 := PrimCons(sym_1_1_6, tmp16679)

	tmp16681 := PrimCons(symA, tmp16680)

	tmp16682 := Call(__e, PrimNS2Value(symdeclare), symshen_4app, tmp16681)

	_ = tmp16682

	tmp16683 := PrimCons(symA, Nil)

	tmp16684 := PrimCons(symlist, tmp16683)

	tmp16685 := PrimCons(symA, Nil)

	tmp16686 := PrimCons(symlist, tmp16685)

	tmp16687 := PrimCons(symA, Nil)

	tmp16688 := PrimCons(symlist, tmp16687)

	tmp16689 := PrimCons(tmp16688, Nil)

	tmp16690 := PrimCons(sym_1_1_6, tmp16689)

	tmp16691 := PrimCons(tmp16686, tmp16690)

	tmp16692 := PrimCons(tmp16691, Nil)

	tmp16693 := PrimCons(sym_1_1_6, tmp16692)

	tmp16694 := PrimCons(tmp16684, tmp16693)

	tmp16695 := Call(__e, PrimNS2Value(symdeclare), symappend, tmp16694)

	_ = tmp16695

	tmp16696 := PrimCons(symnumber, Nil)

	tmp16697 := PrimCons(sym_1_1_6, tmp16696)

	tmp16698 := PrimCons(symA, tmp16697)

	tmp16699 := Call(__e, PrimNS2Value(symdeclare), symarity, tmp16698)

	_ = tmp16699

	tmp16700 := PrimCons(symA, Nil)

	tmp16701 := PrimCons(symlist, tmp16700)

	tmp16702 := PrimCons(tmp16701, Nil)

	tmp16703 := PrimCons(symlist, tmp16702)

	tmp16704 := PrimCons(symA, Nil)

	tmp16705 := PrimCons(symlist, tmp16704)

	tmp16706 := PrimCons(tmp16705, Nil)

	tmp16707 := PrimCons(sym_1_1_6, tmp16706)

	tmp16708 := PrimCons(tmp16703, tmp16707)

	tmp16709 := PrimCons(tmp16708, Nil)

	tmp16710 := PrimCons(sym_1_1_6, tmp16709)

	tmp16711 := PrimCons(symA, tmp16710)

	tmp16712 := Call(__e, PrimNS2Value(symdeclare), symassoc, tmp16711)

	_ = tmp16712

	tmp16713 := PrimCons(symboolean, Nil)

	tmp16714 := PrimCons(sym_1_1_6, tmp16713)

	tmp16715 := PrimCons(symA, tmp16714)

	tmp16716 := Call(__e, PrimNS2Value(symdeclare), symatom_2, tmp16715)

	_ = tmp16716

	tmp16717 := PrimCons(symboolean, Nil)

	tmp16718 := PrimCons(sym_1_1_6, tmp16717)

	tmp16719 := PrimCons(symA, tmp16718)

	tmp16720 := Call(__e, PrimNS2Value(symdeclare), symboolean_2, tmp16719)

	_ = tmp16720

	tmp16721 := PrimCons(symstring, Nil)

	tmp16722 := PrimCons(sym_1_1_6, tmp16721)

	tmp16723 := PrimCons(symstring, tmp16722)

	tmp16724 := Call(__e, PrimNS2Value(symdeclare), symbootstrap, tmp16723)

	_ = tmp16724

	tmp16725 := PrimCons(symboolean, Nil)

	tmp16726 := PrimCons(sym_1_1_6, tmp16725)

	tmp16727 := PrimCons(symsymbol, tmp16726)

	tmp16728 := Call(__e, PrimNS2Value(symdeclare), symbound_2, tmp16727)

	_ = tmp16728

	tmp16729 := PrimCons(symstring, Nil)

	tmp16730 := PrimCons(sym_1_1_6, tmp16729)

	tmp16731 := PrimCons(symstring, tmp16730)

	tmp16732 := Call(__e, PrimNS2Value(symdeclare), symcd, tmp16731)

	_ = tmp16732

	tmp16733 := PrimCons(symA, Nil)

	tmp16734 := PrimCons(symstream, tmp16733)

	tmp16735 := PrimCons(symB, Nil)

	tmp16736 := PrimCons(symlist, tmp16735)

	tmp16737 := PrimCons(tmp16736, Nil)

	tmp16738 := PrimCons(sym_1_1_6, tmp16737)

	tmp16739 := PrimCons(tmp16734, tmp16738)

	tmp16740 := Call(__e, PrimNS2Value(symdeclare), symclose, tmp16739)

	_ = tmp16740

	tmp16741 := PrimCons(symstring, Nil)

	tmp16742 := PrimCons(sym_1_1_6, tmp16741)

	tmp16743 := PrimCons(symstring, tmp16742)

	tmp16744 := PrimCons(tmp16743, Nil)

	tmp16745 := PrimCons(sym_1_1_6, tmp16744)

	tmp16746 := PrimCons(symstring, tmp16745)

	tmp16747 := Call(__e, PrimNS2Value(symdeclare), symcn, tmp16746)

	_ = tmp16747

	tmp16748 := PrimCons(symA, Nil)

	tmp16749 := PrimCons(symlist, tmp16748)

	tmp16750 := PrimCons(symB, Nil)

	tmp16751 := PrimCons(tmp16749, tmp16750)

	tmp16752 := PrimCons(symstr, tmp16751)

	tmp16753 := PrimCons(symA, Nil)

	tmp16754 := PrimCons(symlist, tmp16753)

	tmp16755 := PrimCons(symC, Nil)

	tmp16756 := PrimCons(tmp16754, tmp16755)

	tmp16757 := PrimCons(symstr, tmp16756)

	tmp16758 := PrimCons(tmp16757, Nil)

	tmp16759 := PrimCons(sym_1_1_6, tmp16758)

	tmp16760 := PrimCons(tmp16752, tmp16759)

	tmp16761 := PrimCons(symA, Nil)

	tmp16762 := PrimCons(symlist, tmp16761)

	tmp16763 := PrimCons(symC, Nil)

	tmp16764 := PrimCons(sym_1_1_6, tmp16763)

	tmp16765 := PrimCons(tmp16762, tmp16764)

	tmp16766 := PrimCons(tmp16765, Nil)

	tmp16767 := PrimCons(sym_1_1_6, tmp16766)

	tmp16768 := PrimCons(tmp16760, tmp16767)

	tmp16769 := Call(__e, PrimNS2Value(symdeclare), symcompile, tmp16768)

	_ = tmp16769

	tmp16770 := PrimCons(symboolean, Nil)

	tmp16771 := PrimCons(sym_1_1_6, tmp16770)

	tmp16772 := PrimCons(symA, tmp16771)

	tmp16773 := Call(__e, PrimNS2Value(symdeclare), symcons_2, tmp16772)

	_ = tmp16773

	tmp16774 := PrimCons(symB, Nil)

	tmp16775 := PrimCons(sym_1_1_6, tmp16774)

	tmp16776 := PrimCons(symA, tmp16775)

	tmp16777 := PrimCons(symsymbol, Nil)

	tmp16778 := PrimCons(sym_1_1_6, tmp16777)

	tmp16779 := PrimCons(tmp16776, tmp16778)

	tmp16780 := Call(__e, PrimNS2Value(symdeclare), symdestroy, tmp16779)

	_ = tmp16780

	tmp16781 := PrimCons(symA, Nil)

	tmp16782 := PrimCons(symlist, tmp16781)

	tmp16783 := PrimCons(symA, Nil)

	tmp16784 := PrimCons(symlist, tmp16783)

	tmp16785 := PrimCons(symA, Nil)

	tmp16786 := PrimCons(symlist, tmp16785)

	tmp16787 := PrimCons(tmp16786, Nil)

	tmp16788 := PrimCons(sym_1_1_6, tmp16787)

	tmp16789 := PrimCons(tmp16784, tmp16788)

	tmp16790 := PrimCons(tmp16789, Nil)

	tmp16791 := PrimCons(sym_1_1_6, tmp16790)

	tmp16792 := PrimCons(tmp16782, tmp16791)

	tmp16793 := Call(__e, PrimNS2Value(symdeclare), symdifference, tmp16792)

	_ = tmp16793

	tmp16794 := PrimCons(symB, Nil)

	tmp16795 := PrimCons(sym_1_1_6, tmp16794)

	tmp16796 := PrimCons(symB, tmp16795)

	tmp16797 := PrimCons(tmp16796, Nil)

	tmp16798 := PrimCons(sym_1_1_6, tmp16797)

	tmp16799 := PrimCons(symA, tmp16798)

	tmp16800 := Call(__e, PrimNS2Value(symdeclare), symdo, tmp16799)

	_ = tmp16800

	tmp16801 := PrimCons(symA, Nil)

	tmp16802 := PrimCons(symlist, tmp16801)

	tmp16803 := PrimCons(symB, Nil)

	tmp16804 := PrimCons(tmp16802, tmp16803)

	tmp16805 := PrimCons(symstr, tmp16804)

	tmp16806 := PrimCons(symA, Nil)

	tmp16807 := PrimCons(symlist, tmp16806)

	tmp16808 := PrimCons(symC, Nil)

	tmp16809 := PrimCons(symlist, tmp16808)

	tmp16810 := PrimCons(tmp16809, Nil)

	tmp16811 := PrimCons(tmp16807, tmp16810)

	tmp16812 := PrimCons(symstr, tmp16811)

	tmp16813 := PrimCons(tmp16812, Nil)

	tmp16814 := PrimCons(sym_1_1_6, tmp16813)

	tmp16815 := PrimCons(tmp16805, tmp16814)

	tmp16816 := Call(__e, PrimNS2Value(symdeclare), sym_5e_6, tmp16815)

	_ = tmp16816

	tmp16817 := PrimCons(symA, Nil)

	tmp16818 := PrimCons(symlist, tmp16817)

	tmp16819 := PrimCons(symB, Nil)

	tmp16820 := PrimCons(tmp16818, tmp16819)

	tmp16821 := PrimCons(symstr, tmp16820)

	tmp16822 := PrimCons(symA, Nil)

	tmp16823 := PrimCons(symlist, tmp16822)

	tmp16824 := PrimCons(symA, Nil)

	tmp16825 := PrimCons(symlist, tmp16824)

	tmp16826 := PrimCons(tmp16825, Nil)

	tmp16827 := PrimCons(tmp16823, tmp16826)

	tmp16828 := PrimCons(symstr, tmp16827)

	tmp16829 := PrimCons(tmp16828, Nil)

	tmp16830 := PrimCons(sym_1_1_6, tmp16829)

	tmp16831 := PrimCons(tmp16821, tmp16830)

	tmp16832 := Call(__e, PrimNS2Value(symdeclare), sym_5_b_6, tmp16831)

	_ = tmp16832

	tmp16833 := PrimCons(symA, Nil)

	tmp16834 := PrimCons(symlist, tmp16833)

	tmp16835 := PrimCons(symB, Nil)

	tmp16836 := PrimCons(tmp16834, tmp16835)

	tmp16837 := PrimCons(symstr, tmp16836)

	tmp16838 := PrimCons(symA, Nil)

	tmp16839 := PrimCons(symlist, tmp16838)

	tmp16840 := PrimCons(symB, Nil)

	tmp16841 := PrimCons(tmp16839, tmp16840)

	tmp16842 := PrimCons(symstr, tmp16841)

	tmp16843 := PrimCons(tmp16842, Nil)

	tmp16844 := PrimCons(sym_1_1_6, tmp16843)

	tmp16845 := PrimCons(tmp16837, tmp16844)

	tmp16846 := Call(__e, PrimNS2Value(symdeclare), symshen_4_5end_6, tmp16845)

	_ = tmp16846

	tmp16847 := PrimCons(symA, Nil)

	tmp16848 := PrimCons(symlist, tmp16847)

	tmp16849 := PrimCons(symB, Nil)

	tmp16850 := PrimCons(tmp16848, tmp16849)

	tmp16851 := PrimCons(symstr, tmp16850)

	tmp16852 := PrimCons(symboolean, Nil)

	tmp16853 := PrimCons(sym_1_1_6, tmp16852)

	tmp16854 := PrimCons(symA, tmp16853)

	tmp16855 := PrimCons(tmp16854, Nil)

	tmp16856 := PrimCons(sym_1_1_6, tmp16855)

	tmp16857 := PrimCons(tmp16851, tmp16856)

	tmp16858 := Call(__e, PrimNS2Value(symdeclare), symshen_4_ahd_2, tmp16857)

	_ = tmp16858

	tmp16859 := PrimCons(symA, Nil)

	tmp16860 := PrimCons(symlist, tmp16859)

	tmp16861 := PrimCons(symB, Nil)

	tmp16862 := PrimCons(tmp16860, tmp16861)

	tmp16863 := PrimCons(symstr, tmp16862)

	tmp16864 := PrimCons(symA, Nil)

	tmp16865 := PrimCons(sym_1_1_6, tmp16864)

	tmp16866 := PrimCons(tmp16863, tmp16865)

	tmp16867 := Call(__e, PrimNS2Value(symdeclare), symshen_4hds, tmp16866)

	_ = tmp16867

	tmp16868 := PrimCons(symA, Nil)

	tmp16869 := PrimCons(symlist, tmp16868)

	tmp16870 := PrimCons(symB, Nil)

	tmp16871 := PrimCons(tmp16869, tmp16870)

	tmp16872 := PrimCons(symstr, tmp16871)

	tmp16873 := PrimCons(symA, Nil)

	tmp16874 := PrimCons(symlist, tmp16873)

	tmp16875 := PrimCons(symB, Nil)

	tmp16876 := PrimCons(tmp16874, tmp16875)

	tmp16877 := PrimCons(symstr, tmp16876)

	tmp16878 := PrimCons(tmp16877, Nil)

	tmp16879 := PrimCons(sym_1_1_6, tmp16878)

	tmp16880 := PrimCons(tmp16872, tmp16879)

	tmp16881 := Call(__e, PrimNS2Value(symdeclare), symshen_4tls, tmp16880)

	_ = tmp16881

	tmp16882 := PrimCons(symA, Nil)

	tmp16883 := PrimCons(symlist, tmp16882)

	tmp16884 := PrimCons(symB, Nil)

	tmp16885 := PrimCons(tmp16883, tmp16884)

	tmp16886 := PrimCons(symstr, tmp16885)

	tmp16887 := PrimCons(symboolean, Nil)

	tmp16888 := PrimCons(sym_1_1_6, tmp16887)

	tmp16889 := PrimCons(tmp16886, tmp16888)

	tmp16890 := Call(__e, PrimNS2Value(symdeclare), symshen_4parse_1failure_2, tmp16889)

	_ = tmp16890

	tmp16891 := PrimCons(symA, Nil)

	tmp16892 := PrimCons(symlist, tmp16891)

	tmp16893 := PrimCons(symB, Nil)

	tmp16894 := PrimCons(tmp16892, tmp16893)

	tmp16895 := PrimCons(symstr, tmp16894)

	tmp16896 := PrimCons(tmp16895, Nil)

	tmp16897 := PrimCons(sym_1_1_6, tmp16896)

	tmp16898 := Call(__e, PrimNS2Value(symdeclare), symshen_4parse_1failure, tmp16897)

	_ = tmp16898

	tmp16899 := PrimCons(symA, Nil)

	tmp16900 := PrimCons(symlist, tmp16899)

	tmp16901 := PrimCons(symB, Nil)

	tmp16902 := PrimCons(tmp16900, tmp16901)

	tmp16903 := PrimCons(symstr, tmp16902)

	tmp16904 := PrimCons(symB, Nil)

	tmp16905 := PrimCons(sym_1_1_6, tmp16904)

	tmp16906 := PrimCons(tmp16903, tmp16905)

	tmp16907 := Call(__e, PrimNS2Value(symdeclare), symshen_4_5_1out, tmp16906)

	_ = tmp16907

	tmp16908 := PrimCons(symA, Nil)

	tmp16909 := PrimCons(symlist, tmp16908)

	tmp16910 := PrimCons(symB, Nil)

	tmp16911 := PrimCons(tmp16909, tmp16910)

	tmp16912 := PrimCons(symstr, tmp16911)

	tmp16913 := PrimCons(symA, Nil)

	tmp16914 := PrimCons(symlist, tmp16913)

	tmp16915 := PrimCons(tmp16914, Nil)

	tmp16916 := PrimCons(sym_1_1_6, tmp16915)

	tmp16917 := PrimCons(tmp16912, tmp16916)

	tmp16918 := Call(__e, PrimNS2Value(symdeclare), symshen_4in_1_6, tmp16917)

	_ = tmp16918

	tmp16919 := PrimCons(symA, Nil)

	tmp16920 := PrimCons(symlist, tmp16919)

	tmp16921 := PrimCons(symB, Nil)

	tmp16922 := PrimCons(tmp16920, tmp16921)

	tmp16923 := PrimCons(symstr, tmp16922)

	tmp16924 := PrimCons(symboolean, Nil)

	tmp16925 := PrimCons(sym_1_1_6, tmp16924)

	tmp16926 := PrimCons(tmp16923, tmp16925)

	tmp16927 := Call(__e, PrimNS2Value(symdeclare), symshen_4non_1empty_1stream_2, tmp16926)

	_ = tmp16927

	tmp16928 := PrimCons(symA, Nil)

	tmp16929 := PrimCons(symlist, tmp16928)

	tmp16930 := PrimCons(symA, Nil)

	tmp16931 := PrimCons(symlist, tmp16930)

	tmp16932 := PrimCons(symB, Nil)

	tmp16933 := PrimCons(tmp16931, tmp16932)

	tmp16934 := PrimCons(symstr, tmp16933)

	tmp16935 := PrimCons(tmp16934, Nil)

	tmp16936 := PrimCons(sym_1_1_6, tmp16935)

	tmp16937 := PrimCons(symB, tmp16936)

	tmp16938 := PrimCons(tmp16937, Nil)

	tmp16939 := PrimCons(sym_1_1_6, tmp16938)

	tmp16940 := PrimCons(tmp16929, tmp16939)

	tmp16941 := Call(__e, PrimNS2Value(symdeclare), symshen_4comb, tmp16940)

	_ = tmp16941

	tmp16942 := PrimCons(symB, Nil)

	tmp16943 := PrimCons(symA, tmp16942)

	tmp16944 := PrimCons(symstr, tmp16943)

	tmp16945 := PrimCons(symD, Nil)

	tmp16946 := PrimCons(symC, tmp16945)

	tmp16947 := PrimCons(symstr, tmp16946)

	tmp16948 := PrimCons(symD, Nil)

	tmp16949 := PrimCons(symC, tmp16948)

	tmp16950 := PrimCons(symstr, tmp16949)

	tmp16951 := PrimCons(tmp16950, Nil)

	tmp16952 := PrimCons(symA, tmp16951)

	tmp16953 := PrimCons(symstr, tmp16952)

	tmp16954 := PrimCons(tmp16953, Nil)

	tmp16955 := PrimCons(sym_1_1_6, tmp16954)

	tmp16956 := PrimCons(tmp16947, tmp16955)

	tmp16957 := PrimCons(tmp16956, Nil)

	tmp16958 := PrimCons(sym_1_1_6, tmp16957)

	tmp16959 := PrimCons(tmp16944, tmp16958)

	tmp16960 := Call(__e, PrimNS2Value(symdeclare), symshen_4headstream, tmp16959)

	_ = tmp16960

	tmp16961 := PrimCons(symA, Nil)

	tmp16962 := PrimCons(symlist, tmp16961)

	tmp16963 := PrimCons(symB, Nil)

	tmp16964 := PrimCons(tmp16962, tmp16963)

	tmp16965 := PrimCons(symstr, tmp16964)

	tmp16966 := PrimCons(symA, Nil)

	tmp16967 := PrimCons(symlist, tmp16966)

	tmp16968 := PrimCons(symB, Nil)

	tmp16969 := PrimCons(tmp16967, tmp16968)

	tmp16970 := PrimCons(symstr, tmp16969)

	tmp16971 := PrimCons(tmp16970, Nil)

	tmp16972 := PrimCons(sym_1_1_6, tmp16971)

	tmp16973 := PrimCons(tmp16965, tmp16972)

	tmp16974 := Call(__e, PrimNS2Value(symdeclare), symshen_4tlstream, tmp16973)

	_ = tmp16974

	tmp16975 := PrimCons(symA, Nil)

	tmp16976 := PrimCons(symlist, tmp16975)

	tmp16977 := PrimCons(symboolean, Nil)

	tmp16978 := PrimCons(sym_1_1_6, tmp16977)

	tmp16979 := PrimCons(tmp16976, tmp16978)

	tmp16980 := PrimCons(tmp16979, Nil)

	tmp16981 := PrimCons(sym_1_1_6, tmp16980)

	tmp16982 := PrimCons(symA, tmp16981)

	tmp16983 := Call(__e, PrimNS2Value(symdeclare), symelement_2, tmp16982)

	_ = tmp16983

	tmp16984 := PrimCons(symboolean, Nil)

	tmp16985 := PrimCons(sym_1_1_6, tmp16984)

	tmp16986 := PrimCons(symA, tmp16985)

	tmp16987 := Call(__e, PrimNS2Value(symdeclare), symempty_2, tmp16986)

	_ = tmp16987

	tmp16988 := PrimCons(symboolean, Nil)

	tmp16989 := PrimCons(sym_1_1_6, tmp16988)

	tmp16990 := PrimCons(symsymbol, tmp16989)

	tmp16991 := Call(__e, PrimNS2Value(symdeclare), symenable_1type_1theory, tmp16990)

	_ = tmp16991

	tmp16992 := PrimCons(symsymbol, Nil)

	tmp16993 := PrimCons(symlist, tmp16992)

	tmp16994 := PrimCons(tmp16993, Nil)

	tmp16995 := PrimCons(sym_1_1_6, tmp16994)

	tmp16996 := PrimCons(symsymbol, tmp16995)

	tmp16997 := Call(__e, PrimNS2Value(symdeclare), symexternal, tmp16996)

	_ = tmp16997

	tmp16998 := PrimCons(symstring, Nil)

	tmp16999 := PrimCons(sym_1_1_6, tmp16998)

	tmp17000 := PrimCons(symexception, tmp16999)

	tmp17001 := Call(__e, PrimNS2Value(symdeclare), symerror_1to_1string, tmp17000)

	_ = tmp17001

	tmp17002 := PrimCons(symstring, Nil)

	tmp17003 := PrimCons(symlist, tmp17002)

	tmp17004 := PrimCons(tmp17003, Nil)

	tmp17005 := PrimCons(sym_1_1_6, tmp17004)

	tmp17006 := PrimCons(symA, tmp17005)

	tmp17007 := Call(__e, PrimNS2Value(symdeclare), symexplode, tmp17006)

	_ = tmp17007

	tmp17008 := PrimCons(symsymbol, Nil)

	tmp17009 := PrimCons(sym_1_1_6, tmp17008)

	tmp17010 := PrimCons(symsymbol, tmp17009)

	tmp17011 := Call(__e, PrimNS2Value(symdeclare), symfactorise, tmp17010)

	_ = tmp17011

	tmp17012 := PrimCons(symsymbol, Nil)

	tmp17013 := PrimCons(sym_1_1_6, tmp17012)

	tmp17014 := Call(__e, PrimNS2Value(symdeclare), symfail, tmp17013)

	_ = tmp17014

	tmp17015 := PrimCons(symA, Nil)

	tmp17016 := PrimCons(sym_1_1_6, tmp17015)

	tmp17017 := PrimCons(symA, tmp17016)

	tmp17018 := PrimCons(symA, Nil)

	tmp17019 := PrimCons(sym_1_1_6, tmp17018)

	tmp17020 := PrimCons(symA, tmp17019)

	tmp17021 := PrimCons(tmp17020, Nil)

	tmp17022 := PrimCons(sym_1_1_6, tmp17021)

	tmp17023 := PrimCons(tmp17017, tmp17022)

	tmp17024 := Call(__e, PrimNS2Value(symdeclare), symfix, tmp17023)

	_ = tmp17024

	tmp17025 := PrimCons(symA, Nil)

	tmp17026 := PrimCons(symlazy, tmp17025)

	tmp17027 := PrimCons(tmp17026, Nil)

	tmp17028 := PrimCons(sym_1_1_6, tmp17027)

	tmp17029 := PrimCons(symA, tmp17028)

	tmp17030 := Call(__e, PrimNS2Value(symdeclare), symfreeze, tmp17029)

	_ = tmp17030

	tmp17031 := PrimCons(symB, Nil)

	tmp17032 := PrimCons(sym_d, tmp17031)

	tmp17033 := PrimCons(symA, tmp17032)

	tmp17034 := PrimCons(symA, Nil)

	tmp17035 := PrimCons(sym_1_1_6, tmp17034)

	tmp17036 := PrimCons(tmp17033, tmp17035)

	tmp17037 := Call(__e, PrimNS2Value(symdeclare), symfst, tmp17036)

	_ = tmp17037

	tmp17038 := PrimCons(symsymbol, Nil)

	tmp17039 := PrimCons(sym_1_1_6, tmp17038)

	tmp17040 := PrimCons(symsymbol, tmp17039)

	tmp17041 := Call(__e, PrimNS2Value(symdeclare), symgensym, tmp17040)

	_ = tmp17041

	tmp17042 := PrimCons(symA, Nil)

	tmp17043 := PrimCons(symvector, tmp17042)

	tmp17044 := PrimCons(symA, Nil)

	tmp17045 := PrimCons(sym_1_1_6, tmp17044)

	tmp17046 := PrimCons(symnumber, tmp17045)

	tmp17047 := PrimCons(tmp17046, Nil)

	tmp17048 := PrimCons(sym_1_1_6, tmp17047)

	tmp17049 := PrimCons(tmp17043, tmp17048)

	tmp17050 := Call(__e, PrimNS2Value(symdeclare), sym_5_1vector, tmp17049)

	_ = tmp17050

	tmp17051 := PrimCons(symA, Nil)

	tmp17052 := PrimCons(symvector, tmp17051)

	tmp17053 := PrimCons(symA, Nil)

	tmp17054 := PrimCons(symvector, tmp17053)

	tmp17055 := PrimCons(tmp17054, Nil)

	tmp17056 := PrimCons(sym_1_1_6, tmp17055)

	tmp17057 := PrimCons(symA, tmp17056)

	tmp17058 := PrimCons(tmp17057, Nil)

	tmp17059 := PrimCons(sym_1_1_6, tmp17058)

	tmp17060 := PrimCons(symnumber, tmp17059)

	tmp17061 := PrimCons(tmp17060, Nil)

	tmp17062 := PrimCons(sym_1_1_6, tmp17061)

	tmp17063 := PrimCons(tmp17052, tmp17062)

	tmp17064 := Call(__e, PrimNS2Value(symdeclare), symvector_1_6, tmp17063)

	_ = tmp17064

	tmp17065 := PrimCons(symA, Nil)

	tmp17066 := PrimCons(symvector, tmp17065)

	tmp17067 := PrimCons(tmp17066, Nil)

	tmp17068 := PrimCons(sym_1_1_6, tmp17067)

	tmp17069 := PrimCons(symnumber, tmp17068)

	tmp17070 := Call(__e, PrimNS2Value(symdeclare), symvector, tmp17069)

	_ = tmp17070

	tmp17071 := PrimCons(symnumber, Nil)

	tmp17072 := PrimCons(sym_1_1_6, tmp17071)

	tmp17073 := PrimCons(symsymbol, tmp17072)

	tmp17074 := Call(__e, PrimNS2Value(symdeclare), symget_1time, tmp17073)

	_ = tmp17074

	tmp17075 := PrimCons(symnumber, Nil)

	tmp17076 := PrimCons(sym_1_1_6, tmp17075)

	tmp17077 := PrimCons(symnumber, tmp17076)

	tmp17078 := PrimCons(tmp17077, Nil)

	tmp17079 := PrimCons(sym_1_1_6, tmp17078)

	tmp17080 := PrimCons(symA, tmp17079)

	tmp17081 := Call(__e, PrimNS2Value(symdeclare), symhash, tmp17080)

	_ = tmp17081

	tmp17082 := PrimCons(symA, Nil)

	tmp17083 := PrimCons(symlist, tmp17082)

	tmp17084 := PrimCons(symA, Nil)

	tmp17085 := PrimCons(sym_1_1_6, tmp17084)

	tmp17086 := PrimCons(tmp17083, tmp17085)

	tmp17087 := Call(__e, PrimNS2Value(symdeclare), symhead, tmp17086)

	_ = tmp17087

	tmp17088 := PrimCons(symA, Nil)

	tmp17089 := PrimCons(symvector, tmp17088)

	tmp17090 := PrimCons(symA, Nil)

	tmp17091 := PrimCons(sym_1_1_6, tmp17090)

	tmp17092 := PrimCons(tmp17089, tmp17091)

	tmp17093 := Call(__e, PrimNS2Value(symdeclare), symhdv, tmp17092)

	_ = tmp17093

	tmp17094 := PrimCons(symstring, Nil)

	tmp17095 := PrimCons(sym_1_1_6, tmp17094)

	tmp17096 := PrimCons(symstring, tmp17095)

	tmp17097 := Call(__e, PrimNS2Value(symdeclare), symhdstr, tmp17096)

	_ = tmp17097

	tmp17098 := PrimCons(symA, Nil)

	tmp17099 := PrimCons(sym_1_1_6, tmp17098)

	tmp17100 := PrimCons(symA, tmp17099)

	tmp17101 := PrimCons(tmp17100, Nil)

	tmp17102 := PrimCons(sym_1_1_6, tmp17101)

	tmp17103 := PrimCons(symA, tmp17102)

	tmp17104 := PrimCons(tmp17103, Nil)

	tmp17105 := PrimCons(sym_1_1_6, tmp17104)

	tmp17106 := PrimCons(symboolean, tmp17105)

	tmp17107 := Call(__e, PrimNS2Value(symdeclare), symif, tmp17106)

	_ = tmp17107

	tmp17108 := PrimCons(symsymbol, Nil)

	tmp17109 := PrimCons(sym_1_1_6, tmp17108)

	tmp17110 := PrimCons(symsymbol, tmp17109)

	tmp17111 := Call(__e, PrimNS2Value(symdeclare), symin_1package, tmp17110)

	_ = tmp17111

	tmp17112 := PrimCons(symstring, Nil)

	tmp17113 := PrimCons(sym_1_1_6, tmp17112)

	tmp17114 := Call(__e, PrimNS2Value(symdeclare), symit, tmp17113)

	_ = tmp17114

	tmp17115 := PrimCons(symstring, Nil)

	tmp17116 := PrimCons(sym_1_1_6, tmp17115)

	tmp17117 := Call(__e, PrimNS2Value(symdeclare), symimplementation, tmp17116)

	_ = tmp17117

	tmp17118 := PrimCons(symsymbol, Nil)

	tmp17119 := PrimCons(symlist, tmp17118)

	tmp17120 := PrimCons(symsymbol, Nil)

	tmp17121 := PrimCons(symlist, tmp17120)

	tmp17122 := PrimCons(tmp17121, Nil)

	tmp17123 := PrimCons(sym_1_1_6, tmp17122)

	tmp17124 := PrimCons(tmp17119, tmp17123)

	tmp17125 := Call(__e, PrimNS2Value(symdeclare), syminclude, tmp17124)

	_ = tmp17125

	tmp17126 := PrimCons(symsymbol, Nil)

	tmp17127 := PrimCons(symlist, tmp17126)

	tmp17128 := PrimCons(symsymbol, Nil)

	tmp17129 := PrimCons(symlist, tmp17128)

	tmp17130 := PrimCons(tmp17129, Nil)

	tmp17131 := PrimCons(sym_1_1_6, tmp17130)

	tmp17132 := PrimCons(tmp17127, tmp17131)

	tmp17133 := Call(__e, PrimNS2Value(symdeclare), syminclude_1all_1but, tmp17132)

	_ = tmp17133

	tmp17134 := PrimCons(symnumber, Nil)

	tmp17135 := PrimCons(sym_1_1_6, tmp17134)

	tmp17136 := Call(__e, PrimNS2Value(symdeclare), syminferences, tmp17135)

	_ = tmp17136

	tmp17137 := PrimCons(symstring, Nil)

	tmp17138 := PrimCons(sym_1_1_6, tmp17137)

	tmp17139 := PrimCons(symstring, tmp17138)

	tmp17140 := PrimCons(tmp17139, Nil)

	tmp17141 := PrimCons(sym_1_1_6, tmp17140)

	tmp17142 := PrimCons(symA, tmp17141)

	tmp17143 := Call(__e, PrimNS2Value(symdeclare), symshen_4insert, tmp17142)

	_ = tmp17143

	tmp17144 := PrimCons(symboolean, Nil)

	tmp17145 := PrimCons(sym_1_1_6, tmp17144)

	tmp17146 := PrimCons(symA, tmp17145)

	tmp17147 := Call(__e, PrimNS2Value(symdeclare), syminteger_2, tmp17146)

	_ = tmp17147

	tmp17148 := PrimCons(symsymbol, Nil)

	tmp17149 := PrimCons(symlist, tmp17148)

	tmp17150 := PrimCons(tmp17149, Nil)

	tmp17151 := PrimCons(sym_1_1_6, tmp17150)

	tmp17152 := PrimCons(symsymbol, tmp17151)

	tmp17153 := Call(__e, PrimNS2Value(symdeclare), syminternal, tmp17152)

	_ = tmp17153

	tmp17154 := PrimCons(symA, Nil)

	tmp17155 := PrimCons(symlist, tmp17154)

	tmp17156 := PrimCons(symA, Nil)

	tmp17157 := PrimCons(symlist, tmp17156)

	tmp17158 := PrimCons(symA, Nil)

	tmp17159 := PrimCons(symlist, tmp17158)

	tmp17160 := PrimCons(tmp17159, Nil)

	tmp17161 := PrimCons(sym_1_1_6, tmp17160)

	tmp17162 := PrimCons(tmp17157, tmp17161)

	tmp17163 := PrimCons(tmp17162, Nil)

	tmp17164 := PrimCons(sym_1_1_6, tmp17163)

	tmp17165 := PrimCons(tmp17155, tmp17164)

	tmp17166 := Call(__e, PrimNS2Value(symdeclare), symintersection, tmp17165)

	_ = tmp17166

	tmp17167 := PrimCons(symstring, Nil)

	tmp17168 := PrimCons(sym_1_1_6, tmp17167)

	tmp17169 := Call(__e, PrimNS2Value(symdeclare), symlanguage, tmp17168)

	_ = tmp17169

	tmp17170 := PrimCons(symA, Nil)

	tmp17171 := PrimCons(symlist, tmp17170)

	tmp17172 := PrimCons(symnumber, Nil)

	tmp17173 := PrimCons(sym_1_1_6, tmp17172)

	tmp17174 := PrimCons(tmp17171, tmp17173)

	tmp17175 := Call(__e, PrimNS2Value(symdeclare), symlength, tmp17174)

	_ = tmp17175

	tmp17176 := PrimCons(symA, Nil)

	tmp17177 := PrimCons(symvector, tmp17176)

	tmp17178 := PrimCons(symnumber, Nil)

	tmp17179 := PrimCons(sym_1_1_6, tmp17178)

	tmp17180 := PrimCons(tmp17177, tmp17179)

	tmp17181 := Call(__e, PrimNS2Value(symdeclare), symlimit, tmp17180)

	_ = tmp17181

	tmp17182 := PrimCons(symin, Nil)

	tmp17183 := PrimCons(symstream, tmp17182)

	tmp17184 := PrimCons(symunit, Nil)

	tmp17185 := PrimCons(symlist, tmp17184)

	tmp17186 := PrimCons(tmp17185, Nil)

	tmp17187 := PrimCons(sym_1_1_6, tmp17186)

	tmp17188 := PrimCons(tmp17183, tmp17187)

	tmp17189 := Call(__e, PrimNS2Value(symdeclare), symlineread, tmp17188)

	_ = tmp17189

	tmp17190 := PrimCons(symsymbol, Nil)

	tmp17191 := PrimCons(sym_1_1_6, tmp17190)

	tmp17192 := PrimCons(symstring, tmp17191)

	tmp17193 := Call(__e, PrimNS2Value(symdeclare), symload, tmp17192)

	_ = tmp17193

	tmp17194 := PrimCons(symB, Nil)

	tmp17195 := PrimCons(sym_1_1_6, tmp17194)

	tmp17196 := PrimCons(symA, tmp17195)

	tmp17197 := PrimCons(symA, Nil)

	tmp17198 := PrimCons(symlist, tmp17197)

	tmp17199 := PrimCons(symB, Nil)

	tmp17200 := PrimCons(symlist, tmp17199)

	tmp17201 := PrimCons(tmp17200, Nil)

	tmp17202 := PrimCons(sym_1_1_6, tmp17201)

	tmp17203 := PrimCons(tmp17198, tmp17202)

	tmp17204 := PrimCons(tmp17203, Nil)

	tmp17205 := PrimCons(sym_1_1_6, tmp17204)

	tmp17206 := PrimCons(tmp17196, tmp17205)

	tmp17207 := Call(__e, PrimNS2Value(symdeclare), symmap, tmp17206)

	_ = tmp17207

	tmp17208 := PrimCons(symB, Nil)

	tmp17209 := PrimCons(symlist, tmp17208)

	tmp17210 := PrimCons(tmp17209, Nil)

	tmp17211 := PrimCons(sym_1_1_6, tmp17210)

	tmp17212 := PrimCons(symA, tmp17211)

	tmp17213 := PrimCons(symA, Nil)

	tmp17214 := PrimCons(symlist, tmp17213)

	tmp17215 := PrimCons(symB, Nil)

	tmp17216 := PrimCons(symlist, tmp17215)

	tmp17217 := PrimCons(tmp17216, Nil)

	tmp17218 := PrimCons(sym_1_1_6, tmp17217)

	tmp17219 := PrimCons(tmp17214, tmp17218)

	tmp17220 := PrimCons(tmp17219, Nil)

	tmp17221 := PrimCons(sym_1_1_6, tmp17220)

	tmp17222 := PrimCons(tmp17212, tmp17221)

	tmp17223 := Call(__e, PrimNS2Value(symdeclare), symmapcan, tmp17222)

	_ = tmp17223

	tmp17224 := PrimCons(symnumber, Nil)

	tmp17225 := PrimCons(sym_1_1_6, tmp17224)

	tmp17226 := PrimCons(symnumber, tmp17225)

	tmp17227 := Call(__e, PrimNS2Value(symdeclare), symmaxinferences, tmp17226)

	_ = tmp17227

	tmp17228 := PrimCons(symstring, Nil)

	tmp17229 := PrimCons(sym_1_1_6, tmp17228)

	tmp17230 := PrimCons(symnumber, tmp17229)

	tmp17231 := Call(__e, PrimNS2Value(symdeclare), symn_1_6string, tmp17230)

	_ = tmp17231

	tmp17232 := PrimCons(symnumber, Nil)

	tmp17233 := PrimCons(sym_1_1_6, tmp17232)

	tmp17234 := PrimCons(symnumber, tmp17233)

	tmp17235 := Call(__e, PrimNS2Value(symdeclare), symnl, tmp17234)

	_ = tmp17235

	tmp17236 := PrimCons(symboolean, Nil)

	tmp17237 := PrimCons(sym_1_1_6, tmp17236)

	tmp17238 := PrimCons(symboolean, tmp17237)

	tmp17239 := Call(__e, PrimNS2Value(symdeclare), symnot, tmp17238)

	_ = tmp17239

	tmp17240 := PrimCons(symA, Nil)

	tmp17241 := PrimCons(symlist, tmp17240)

	tmp17242 := PrimCons(symA, Nil)

	tmp17243 := PrimCons(sym_1_1_6, tmp17242)

	tmp17244 := PrimCons(tmp17241, tmp17243)

	tmp17245 := PrimCons(tmp17244, Nil)

	tmp17246 := PrimCons(sym_1_1_6, tmp17245)

	tmp17247 := PrimCons(symnumber, tmp17246)

	tmp17248 := Call(__e, PrimNS2Value(symdeclare), symnth, tmp17247)

	_ = tmp17248

	tmp17249 := PrimCons(symboolean, Nil)

	tmp17250 := PrimCons(sym_1_1_6, tmp17249)

	tmp17251 := PrimCons(symA, tmp17250)

	tmp17252 := Call(__e, PrimNS2Value(symdeclare), symnumber_2, tmp17251)

	_ = tmp17252

	tmp17253 := PrimCons(symnumber, Nil)

	tmp17254 := PrimCons(sym_1_1_6, tmp17253)

	tmp17255 := PrimCons(symB, tmp17254)

	tmp17256 := PrimCons(tmp17255, Nil)

	tmp17257 := PrimCons(sym_1_1_6, tmp17256)

	tmp17258 := PrimCons(symA, tmp17257)

	tmp17259 := Call(__e, PrimNS2Value(symdeclare), symoccurrences, tmp17258)

	_ = tmp17259

	tmp17260 := PrimCons(symboolean, Nil)

	tmp17261 := PrimCons(sym_1_1_6, tmp17260)

	tmp17262 := PrimCons(symsymbol, tmp17261)

	tmp17263 := Call(__e, PrimNS2Value(symdeclare), symoccurs_1check, tmp17262)

	_ = tmp17263

	tmp17264 := PrimCons(symboolean, Nil)

	tmp17265 := PrimCons(sym_1_1_6, tmp17264)

	tmp17266 := PrimCons(symsymbol, tmp17265)

	tmp17267 := Call(__e, PrimNS2Value(symdeclare), symoptimise, tmp17266)

	_ = tmp17267

	tmp17268 := PrimCons(symboolean, Nil)

	tmp17269 := PrimCons(sym_1_1_6, tmp17268)

	tmp17270 := PrimCons(symboolean, tmp17269)

	tmp17271 := PrimCons(tmp17270, Nil)

	tmp17272 := PrimCons(sym_1_1_6, tmp17271)

	tmp17273 := PrimCons(symboolean, tmp17272)

	tmp17274 := Call(__e, PrimNS2Value(symdeclare), symor, tmp17273)

	_ = tmp17274

	tmp17275 := PrimCons(symstring, Nil)

	tmp17276 := PrimCons(sym_1_1_6, tmp17275)

	tmp17277 := Call(__e, PrimNS2Value(symdeclare), symos, tmp17276)

	_ = tmp17277

	tmp17278 := PrimCons(symboolean, Nil)

	tmp17279 := PrimCons(sym_1_1_6, tmp17278)

	tmp17280 := PrimCons(symsymbol, tmp17279)

	tmp17281 := Call(__e, PrimNS2Value(symdeclare), sympackage_2, tmp17280)

	_ = tmp17281

	tmp17282 := PrimCons(symstring, Nil)

	tmp17283 := PrimCons(sym_1_1_6, tmp17282)

	tmp17284 := Call(__e, PrimNS2Value(symdeclare), symport, tmp17283)

	_ = tmp17284

	tmp17285 := PrimCons(symstring, Nil)

	tmp17286 := PrimCons(sym_1_1_6, tmp17285)

	tmp17287 := Call(__e, PrimNS2Value(symdeclare), symporters, tmp17286)

	_ = tmp17287

	tmp17288 := PrimCons(symstring, Nil)

	tmp17289 := PrimCons(sym_1_1_6, tmp17288)

	tmp17290 := PrimCons(symnumber, tmp17289)

	tmp17291 := PrimCons(tmp17290, Nil)

	tmp17292 := PrimCons(sym_1_1_6, tmp17291)

	tmp17293 := PrimCons(symstring, tmp17292)

	tmp17294 := Call(__e, PrimNS2Value(symdeclare), sympos, tmp17293)

	_ = tmp17294

	tmp17295 := PrimCons(symout, Nil)

	tmp17296 := PrimCons(symstream, tmp17295)

	tmp17297 := PrimCons(symstring, Nil)

	tmp17298 := PrimCons(sym_1_1_6, tmp17297)

	tmp17299 := PrimCons(tmp17296, tmp17298)

	tmp17300 := PrimCons(tmp17299, Nil)

	tmp17301 := PrimCons(sym_1_1_6, tmp17300)

	tmp17302 := PrimCons(symstring, tmp17301)

	tmp17303 := Call(__e, PrimNS2Value(symdeclare), sympr, tmp17302)

	_ = tmp17303

	tmp17304 := PrimCons(symA, Nil)

	tmp17305 := PrimCons(sym_1_1_6, tmp17304)

	tmp17306 := PrimCons(symA, tmp17305)

	tmp17307 := Call(__e, PrimNS2Value(symdeclare), symprint, tmp17306)

	_ = tmp17307

	tmp17308 := PrimCons(symsymbol, Nil)

	tmp17309 := PrimCons(sym_1_1_6, tmp17308)

	tmp17310 := PrimCons(symsymbol, tmp17309)

	tmp17311 := Call(__e, PrimNS2Value(symdeclare), symprofile, tmp17310)

	_ = tmp17311

	tmp17312 := PrimCons(symsymbol, Nil)

	tmp17313 := PrimCons(symlist, tmp17312)

	tmp17314 := PrimCons(symsymbol, Nil)

	tmp17315 := PrimCons(symlist, tmp17314)

	tmp17316 := PrimCons(tmp17315, Nil)

	tmp17317 := PrimCons(sym_1_1_6, tmp17316)

	tmp17318 := PrimCons(tmp17313, tmp17317)

	tmp17319 := Call(__e, PrimNS2Value(symdeclare), sympreclude, tmp17318)

	_ = tmp17319

	tmp17320 := PrimCons(symstring, Nil)

	tmp17321 := PrimCons(sym_1_1_6, tmp17320)

	tmp17322 := PrimCons(symstring, tmp17321)

	tmp17323 := Call(__e, PrimNS2Value(symdeclare), symshen_4proc_1nl, tmp17322)

	_ = tmp17323

	tmp17324 := PrimCons(symnumber, Nil)

	tmp17325 := PrimCons(sym_d, tmp17324)

	tmp17326 := PrimCons(symsymbol, tmp17325)

	tmp17327 := PrimCons(tmp17326, Nil)

	tmp17328 := PrimCons(sym_1_1_6, tmp17327)

	tmp17329 := PrimCons(symsymbol, tmp17328)

	tmp17330 := Call(__e, PrimNS2Value(symdeclare), symprofile_1results, tmp17329)

	_ = tmp17330

	tmp17331 := PrimCons(symA, Nil)

	tmp17332 := PrimCons(sym_1_1_6, tmp17331)

	tmp17333 := PrimCons(symA, tmp17332)

	tmp17334 := Call(__e, PrimNS2Value(symdeclare), symprotect, tmp17333)

	_ = tmp17334

	tmp17335 := PrimCons(symsymbol, Nil)

	tmp17336 := PrimCons(symlist, tmp17335)

	tmp17337 := PrimCons(symsymbol, Nil)

	tmp17338 := PrimCons(symlist, tmp17337)

	tmp17339 := PrimCons(tmp17338, Nil)

	tmp17340 := PrimCons(sym_1_1_6, tmp17339)

	tmp17341 := PrimCons(tmp17336, tmp17340)

	tmp17342 := Call(__e, PrimNS2Value(symdeclare), sympreclude_1all_1but, tmp17341)

	_ = tmp17342

	tmp17343 := PrimCons(symout, Nil)

	tmp17344 := PrimCons(symstream, tmp17343)

	tmp17345 := PrimCons(symstring, Nil)

	tmp17346 := PrimCons(sym_1_1_6, tmp17345)

	tmp17347 := PrimCons(tmp17344, tmp17346)

	tmp17348 := PrimCons(tmp17347, Nil)

	tmp17349 := PrimCons(sym_1_1_6, tmp17348)

	tmp17350 := PrimCons(symstring, tmp17349)

	tmp17351 := Call(__e, PrimNS2Value(symdeclare), symshen_4prhush, tmp17350)

	_ = tmp17351

	tmp17352 := PrimCons(symnumber, Nil)

	tmp17353 := PrimCons(sym_1_1_6, tmp17352)

	tmp17354 := PrimCons(symnumber, tmp17353)

	tmp17355 := Call(__e, PrimNS2Value(symdeclare), symprolog_1memory, tmp17354)

	_ = tmp17355

	tmp17356 := PrimCons(symunit, Nil)

	tmp17357 := PrimCons(symlist, tmp17356)

	tmp17358 := PrimCons(tmp17357, Nil)

	tmp17359 := PrimCons(sym_1_1_6, tmp17358)

	tmp17360 := PrimCons(symsymbol, tmp17359)

	tmp17361 := Call(__e, PrimNS2Value(symdeclare), symps, tmp17360)

	_ = tmp17361

	tmp17362 := PrimCons(symin, Nil)

	tmp17363 := PrimCons(symstream, tmp17362)

	tmp17364 := PrimCons(symunit, Nil)

	tmp17365 := PrimCons(sym_1_1_6, tmp17364)

	tmp17366 := PrimCons(tmp17363, tmp17365)

	tmp17367 := Call(__e, PrimNS2Value(symdeclare), symread, tmp17366)

	_ = tmp17367

	tmp17368 := PrimCons(symin, Nil)

	tmp17369 := PrimCons(symstream, tmp17368)

	tmp17370 := PrimCons(symnumber, Nil)

	tmp17371 := PrimCons(sym_1_1_6, tmp17370)

	tmp17372 := PrimCons(tmp17369, tmp17371)

	tmp17373 := Call(__e, PrimNS2Value(symdeclare), symread_1byte, tmp17372)

	_ = tmp17373

	tmp17374 := PrimCons(symnumber, Nil)

	tmp17375 := PrimCons(symlist, tmp17374)

	tmp17376 := PrimCons(tmp17375, Nil)

	tmp17377 := PrimCons(sym_1_1_6, tmp17376)

	tmp17378 := PrimCons(symstring, tmp17377)

	tmp17379 := Call(__e, PrimNS2Value(symdeclare), symread_1file_1as_1bytelist, tmp17378)

	_ = tmp17379

	tmp17380 := PrimCons(symstring, Nil)

	tmp17381 := PrimCons(sym_1_1_6, tmp17380)

	tmp17382 := PrimCons(symstring, tmp17381)

	tmp17383 := Call(__e, PrimNS2Value(symdeclare), symread_1file_1as_1string, tmp17382)

	_ = tmp17383

	tmp17384 := PrimCons(symunit, Nil)

	tmp17385 := PrimCons(symlist, tmp17384)

	tmp17386 := PrimCons(tmp17385, Nil)

	tmp17387 := PrimCons(sym_1_1_6, tmp17386)

	tmp17388 := PrimCons(symstring, tmp17387)

	tmp17389 := Call(__e, PrimNS2Value(symdeclare), symread_1file, tmp17388)

	_ = tmp17389

	tmp17390 := PrimCons(symunit, Nil)

	tmp17391 := PrimCons(symlist, tmp17390)

	tmp17392 := PrimCons(tmp17391, Nil)

	tmp17393 := PrimCons(sym_1_1_6, tmp17392)

	tmp17394 := PrimCons(symstring, tmp17393)

	tmp17395 := Call(__e, PrimNS2Value(symdeclare), symread_1from_1string, tmp17394)

	_ = tmp17395

	tmp17396 := PrimCons(symunit, Nil)

	tmp17397 := PrimCons(symlist, tmp17396)

	tmp17398 := PrimCons(tmp17397, Nil)

	tmp17399 := PrimCons(sym_1_1_6, tmp17398)

	tmp17400 := PrimCons(symstring, tmp17399)

	tmp17401 := Call(__e, PrimNS2Value(symdeclare), symread_1from_1string_1unprocessed, tmp17400)

	_ = tmp17401

	tmp17402 := PrimCons(symstring, Nil)

	tmp17403 := PrimCons(sym_1_1_6, tmp17402)

	tmp17404 := Call(__e, PrimNS2Value(symdeclare), symrelease, tmp17403)

	_ = tmp17404

	tmp17405 := PrimCons(symA, Nil)

	tmp17406 := PrimCons(symlist, tmp17405)

	tmp17407 := PrimCons(symA, Nil)

	tmp17408 := PrimCons(symlist, tmp17407)

	tmp17409 := PrimCons(tmp17408, Nil)

	tmp17410 := PrimCons(sym_1_1_6, tmp17409)

	tmp17411 := PrimCons(tmp17406, tmp17410)

	tmp17412 := PrimCons(tmp17411, Nil)

	tmp17413 := PrimCons(sym_1_1_6, tmp17412)

	tmp17414 := PrimCons(symA, tmp17413)

	tmp17415 := Call(__e, PrimNS2Value(symdeclare), symremove, tmp17414)

	_ = tmp17415

	tmp17416 := PrimCons(symA, Nil)

	tmp17417 := PrimCons(symlist, tmp17416)

	tmp17418 := PrimCons(symA, Nil)

	tmp17419 := PrimCons(symlist, tmp17418)

	tmp17420 := PrimCons(tmp17419, Nil)

	tmp17421 := PrimCons(sym_1_1_6, tmp17420)

	tmp17422 := PrimCons(tmp17417, tmp17421)

	tmp17423 := Call(__e, PrimNS2Value(symdeclare), symreverse, tmp17422)

	_ = tmp17423

	tmp17424 := PrimCons(symA, Nil)

	tmp17425 := PrimCons(sym_1_1_6, tmp17424)

	tmp17426 := PrimCons(symstring, tmp17425)

	tmp17427 := Call(__e, PrimNS2Value(symdeclare), symsimple_1error, tmp17426)

	_ = tmp17427

	tmp17428 := PrimCons(symB, Nil)

	tmp17429 := PrimCons(sym_d, tmp17428)

	tmp17430 := PrimCons(symA, tmp17429)

	tmp17431 := PrimCons(symB, Nil)

	tmp17432 := PrimCons(sym_1_1_6, tmp17431)

	tmp17433 := PrimCons(tmp17430, tmp17432)

	tmp17434 := Call(__e, PrimNS2Value(symdeclare), symsnd, tmp17433)

	_ = tmp17434

	tmp17435 := PrimCons(symsymbol, Nil)

	tmp17436 := PrimCons(sym_1_1_6, tmp17435)

	tmp17437 := PrimCons(symnumber, tmp17436)

	tmp17438 := PrimCons(tmp17437, Nil)

	tmp17439 := PrimCons(sym_1_1_6, tmp17438)

	tmp17440 := PrimCons(symsymbol, tmp17439)

	tmp17441 := Call(__e, PrimNS2Value(symdeclare), symspecialise, tmp17440)

	_ = tmp17441

	tmp17442 := PrimCons(symboolean, Nil)

	tmp17443 := PrimCons(sym_1_1_6, tmp17442)

	tmp17444 := PrimCons(symsymbol, tmp17443)

	tmp17445 := Call(__e, PrimNS2Value(symdeclare), symspy, tmp17444)

	_ = tmp17445

	tmp17446 := PrimCons(symboolean, Nil)

	tmp17447 := PrimCons(sym_1_1_6, tmp17446)

	tmp17448 := PrimCons(symsymbol, tmp17447)

	tmp17449 := Call(__e, PrimNS2Value(symdeclare), symstep, tmp17448)

	_ = tmp17449

	tmp17450 := PrimCons(symin, Nil)

	tmp17451 := PrimCons(symstream, tmp17450)

	tmp17452 := PrimCons(tmp17451, Nil)

	tmp17453 := PrimCons(sym_1_1_6, tmp17452)

	tmp17454 := Call(__e, PrimNS2Value(symdeclare), symstinput, tmp17453)

	_ = tmp17454

	tmp17455 := PrimCons(symout, Nil)

	tmp17456 := PrimCons(symstream, tmp17455)

	tmp17457 := PrimCons(tmp17456, Nil)

	tmp17458 := PrimCons(sym_1_1_6, tmp17457)

	tmp17459 := Call(__e, PrimNS2Value(symdeclare), symstoutput, tmp17458)

	_ = tmp17459

	tmp17460 := PrimCons(symboolean, Nil)

	tmp17461 := PrimCons(sym_1_1_6, tmp17460)

	tmp17462 := PrimCons(symA, tmp17461)

	tmp17463 := Call(__e, PrimNS2Value(symdeclare), symstring_2, tmp17462)

	_ = tmp17463

	tmp17464 := PrimCons(symstring, Nil)

	tmp17465 := PrimCons(sym_1_1_6, tmp17464)

	tmp17466 := PrimCons(symA, tmp17465)

	tmp17467 := Call(__e, PrimNS2Value(symdeclare), symstr, tmp17466)

	_ = tmp17467

	tmp17468 := PrimCons(symnumber, Nil)

	tmp17469 := PrimCons(sym_1_1_6, tmp17468)

	tmp17470 := PrimCons(symstring, tmp17469)

	tmp17471 := Call(__e, PrimNS2Value(symdeclare), symstring_1_6n, tmp17470)

	_ = tmp17471

	tmp17472 := PrimCons(symsymbol, Nil)

	tmp17473 := PrimCons(sym_1_1_6, tmp17472)

	tmp17474 := PrimCons(symstring, tmp17473)

	tmp17475 := Call(__e, PrimNS2Value(symdeclare), symstring_1_6symbol, tmp17474)

	_ = tmp17475

	tmp17476 := PrimCons(symnumber, Nil)

	tmp17477 := PrimCons(symlist, tmp17476)

	tmp17478 := PrimCons(symnumber, Nil)

	tmp17479 := PrimCons(sym_1_1_6, tmp17478)

	tmp17480 := PrimCons(tmp17477, tmp17479)

	tmp17481 := Call(__e, PrimNS2Value(symdeclare), symsum, tmp17480)

	_ = tmp17481

	tmp17482 := PrimCons(symboolean, Nil)

	tmp17483 := PrimCons(sym_1_1_6, tmp17482)

	tmp17484 := PrimCons(symA, tmp17483)

	tmp17485 := Call(__e, PrimNS2Value(symdeclare), symsymbol_2, tmp17484)

	_ = tmp17485

	tmp17486 := PrimCons(symsymbol, Nil)

	tmp17487 := PrimCons(sym_1_1_6, tmp17486)

	tmp17488 := PrimCons(symsymbol, tmp17487)

	tmp17489 := Call(__e, PrimNS2Value(symdeclare), symsystemf, tmp17488)

	_ = tmp17489

	tmp17490 := PrimCons(symA, Nil)

	tmp17491 := PrimCons(symlist, tmp17490)

	tmp17492 := PrimCons(symA, Nil)

	tmp17493 := PrimCons(symlist, tmp17492)

	tmp17494 := PrimCons(tmp17493, Nil)

	tmp17495 := PrimCons(sym_1_1_6, tmp17494)

	tmp17496 := PrimCons(tmp17491, tmp17495)

	tmp17497 := Call(__e, PrimNS2Value(symdeclare), symtail, tmp17496)

	_ = tmp17497

	tmp17498 := PrimCons(symstring, Nil)

	tmp17499 := PrimCons(sym_1_1_6, tmp17498)

	tmp17500 := PrimCons(symstring, tmp17499)

	tmp17501 := Call(__e, PrimNS2Value(symdeclare), symtlstr, tmp17500)

	_ = tmp17501

	tmp17502 := PrimCons(symA, Nil)

	tmp17503 := PrimCons(symvector, tmp17502)

	tmp17504 := PrimCons(symA, Nil)

	tmp17505 := PrimCons(symvector, tmp17504)

	tmp17506 := PrimCons(tmp17505, Nil)

	tmp17507 := PrimCons(sym_1_1_6, tmp17506)

	tmp17508 := PrimCons(tmp17503, tmp17507)

	tmp17509 := Call(__e, PrimNS2Value(symdeclare), symtlv, tmp17508)

	_ = tmp17509

	tmp17510 := PrimCons(symboolean, Nil)

	tmp17511 := PrimCons(sym_1_1_6, tmp17510)

	tmp17512 := PrimCons(symsymbol, tmp17511)

	tmp17513 := Call(__e, PrimNS2Value(symdeclare), symtc, tmp17512)

	_ = tmp17513

	tmp17514 := PrimCons(symboolean, Nil)

	tmp17515 := PrimCons(sym_1_1_6, tmp17514)

	tmp17516 := Call(__e, PrimNS2Value(symdeclare), symtc_2, tmp17515)

	_ = tmp17516

	tmp17517 := PrimCons(symA, Nil)

	tmp17518 := PrimCons(symlazy, tmp17517)

	tmp17519 := PrimCons(symA, Nil)

	tmp17520 := PrimCons(sym_1_1_6, tmp17519)

	tmp17521 := PrimCons(tmp17518, tmp17520)

	tmp17522 := Call(__e, PrimNS2Value(symdeclare), symthaw, tmp17521)

	_ = tmp17522

	tmp17523 := PrimCons(symsymbol, Nil)

	tmp17524 := PrimCons(sym_1_1_6, tmp17523)

	tmp17525 := PrimCons(symsymbol, tmp17524)

	tmp17526 := Call(__e, PrimNS2Value(symdeclare), symtrack, tmp17525)

	_ = tmp17526

	tmp17527 := PrimCons(symA, Nil)

	tmp17528 := PrimCons(sym_1_1_6, tmp17527)

	tmp17529 := PrimCons(symexception, tmp17528)

	tmp17530 := PrimCons(symA, Nil)

	tmp17531 := PrimCons(sym_1_1_6, tmp17530)

	tmp17532 := PrimCons(tmp17529, tmp17531)

	tmp17533 := PrimCons(tmp17532, Nil)

	tmp17534 := PrimCons(sym_1_1_6, tmp17533)

	tmp17535 := PrimCons(symA, tmp17534)

	tmp17536 := Call(__e, PrimNS2Value(symdeclare), symtrap_1error, tmp17535)

	_ = tmp17536

	tmp17537 := PrimCons(symboolean, Nil)

	tmp17538 := PrimCons(sym_1_1_6, tmp17537)

	tmp17539 := PrimCons(symA, tmp17538)

	tmp17540 := Call(__e, PrimNS2Value(symdeclare), symtuple_2, tmp17539)

	_ = tmp17540

	tmp17541 := PrimCons(symsymbol, Nil)

	tmp17542 := PrimCons(sym_1_1_6, tmp17541)

	tmp17543 := PrimCons(symsymbol, tmp17542)

	tmp17544 := Call(__e, PrimNS2Value(symdeclare), symundefmacro, tmp17543)

	_ = tmp17544

	tmp17545 := PrimCons(symA, Nil)

	tmp17546 := PrimCons(symlist, tmp17545)

	tmp17547 := PrimCons(symA, Nil)

	tmp17548 := PrimCons(symlist, tmp17547)

	tmp17549 := PrimCons(symA, Nil)

	tmp17550 := PrimCons(symlist, tmp17549)

	tmp17551 := PrimCons(tmp17550, Nil)

	tmp17552 := PrimCons(sym_1_1_6, tmp17551)

	tmp17553 := PrimCons(tmp17548, tmp17552)

	tmp17554 := PrimCons(tmp17553, Nil)

	tmp17555 := PrimCons(sym_1_1_6, tmp17554)

	tmp17556 := PrimCons(tmp17546, tmp17555)

	tmp17557 := Call(__e, PrimNS2Value(symdeclare), symunion, tmp17556)

	_ = tmp17557

	tmp17558 := PrimCons(symsymbol, Nil)

	tmp17559 := PrimCons(sym_1_1_6, tmp17558)

	tmp17560 := PrimCons(symsymbol, tmp17559)

	tmp17561 := Call(__e, PrimNS2Value(symdeclare), symunprofile, tmp17560)

	_ = tmp17561

	tmp17562 := PrimCons(symsymbol, Nil)

	tmp17563 := PrimCons(sym_1_1_6, tmp17562)

	tmp17564 := PrimCons(symsymbol, tmp17563)

	tmp17565 := Call(__e, PrimNS2Value(symdeclare), symuntrack, tmp17564)

	_ = tmp17565

	tmp17566 := PrimCons(symboolean, Nil)

	tmp17567 := PrimCons(sym_1_1_6, tmp17566)

	tmp17568 := PrimCons(symA, tmp17567)

	tmp17569 := Call(__e, PrimNS2Value(symdeclare), symvariable_2, tmp17568)

	_ = tmp17569

	tmp17570 := PrimCons(symboolean, Nil)

	tmp17571 := PrimCons(sym_1_1_6, tmp17570)

	tmp17572 := PrimCons(symA, tmp17571)

	tmp17573 := Call(__e, PrimNS2Value(symdeclare), symvector_2, tmp17572)

	_ = tmp17573

	tmp17574 := PrimCons(symstring, Nil)

	tmp17575 := PrimCons(sym_1_1_6, tmp17574)

	tmp17576 := Call(__e, PrimNS2Value(symdeclare), symversion, tmp17575)

	_ = tmp17576

	tmp17577 := PrimCons(symA, Nil)

	tmp17578 := PrimCons(sym_1_1_6, tmp17577)

	tmp17579 := PrimCons(symA, tmp17578)

	tmp17580 := PrimCons(tmp17579, Nil)

	tmp17581 := PrimCons(sym_1_1_6, tmp17580)

	tmp17582 := PrimCons(symstring, tmp17581)

	tmp17583 := Call(__e, PrimNS2Value(symdeclare), symwrite_1to_1file, tmp17582)

	_ = tmp17583

	tmp17584 := PrimCons(symout, Nil)

	tmp17585 := PrimCons(symstream, tmp17584)

	tmp17586 := PrimCons(symnumber, Nil)

	tmp17587 := PrimCons(sym_1_1_6, tmp17586)

	tmp17588 := PrimCons(tmp17585, tmp17587)

	tmp17589 := PrimCons(tmp17588, Nil)

	tmp17590 := PrimCons(sym_1_1_6, tmp17589)

	tmp17591 := PrimCons(symnumber, tmp17590)

	tmp17592 := Call(__e, PrimNS2Value(symdeclare), symwrite_1byte, tmp17591)

	_ = tmp17592

	tmp17593 := PrimCons(symboolean, Nil)

	tmp17594 := PrimCons(sym_1_1_6, tmp17593)

	tmp17595 := PrimCons(symstring, tmp17594)

	tmp17596 := Call(__e, PrimNS2Value(symdeclare), symy_1or_1n_2, tmp17595)

	_ = tmp17596

	tmp17597 := PrimCons(symboolean, Nil)

	tmp17598 := PrimCons(sym_1_1_6, tmp17597)

	tmp17599 := PrimCons(symnumber, tmp17598)

	tmp17600 := PrimCons(tmp17599, Nil)

	tmp17601 := PrimCons(sym_1_1_6, tmp17600)

	tmp17602 := PrimCons(symnumber, tmp17601)

	tmp17603 := Call(__e, PrimNS2Value(symdeclare), sym_6, tmp17602)

	_ = tmp17603

	tmp17604 := PrimCons(symboolean, Nil)

	tmp17605 := PrimCons(sym_1_1_6, tmp17604)

	tmp17606 := PrimCons(symnumber, tmp17605)

	tmp17607 := PrimCons(tmp17606, Nil)

	tmp17608 := PrimCons(sym_1_1_6, tmp17607)

	tmp17609 := PrimCons(symnumber, tmp17608)

	tmp17610 := Call(__e, PrimNS2Value(symdeclare), sym_5, tmp17609)

	_ = tmp17610

	tmp17611 := PrimCons(symboolean, Nil)

	tmp17612 := PrimCons(sym_1_1_6, tmp17611)

	tmp17613 := PrimCons(symnumber, tmp17612)

	tmp17614 := PrimCons(tmp17613, Nil)

	tmp17615 := PrimCons(sym_1_1_6, tmp17614)

	tmp17616 := PrimCons(symnumber, tmp17615)

	tmp17617 := Call(__e, PrimNS2Value(symdeclare), sym_6_a, tmp17616)

	_ = tmp17617

	tmp17618 := PrimCons(symboolean, Nil)

	tmp17619 := PrimCons(sym_1_1_6, tmp17618)

	tmp17620 := PrimCons(symnumber, tmp17619)

	tmp17621 := PrimCons(tmp17620, Nil)

	tmp17622 := PrimCons(sym_1_1_6, tmp17621)

	tmp17623 := PrimCons(symnumber, tmp17622)

	tmp17624 := Call(__e, PrimNS2Value(symdeclare), sym_5_a, tmp17623)

	_ = tmp17624

	tmp17625 := PrimCons(symboolean, Nil)

	tmp17626 := PrimCons(sym_1_1_6, tmp17625)

	tmp17627 := PrimCons(symA, tmp17626)

	tmp17628 := PrimCons(tmp17627, Nil)

	tmp17629 := PrimCons(sym_1_1_6, tmp17628)

	tmp17630 := PrimCons(symA, tmp17629)

	tmp17631 := Call(__e, PrimNS2Value(symdeclare), sym_a, tmp17630)

	_ = tmp17631

	tmp17632 := PrimCons(symnumber, Nil)

	tmp17633 := PrimCons(sym_1_1_6, tmp17632)

	tmp17634 := PrimCons(symnumber, tmp17633)

	tmp17635 := PrimCons(tmp17634, Nil)

	tmp17636 := PrimCons(sym_1_1_6, tmp17635)

	tmp17637 := PrimCons(symnumber, tmp17636)

	tmp17638 := Call(__e, PrimNS2Value(symdeclare), sym_7, tmp17637)

	_ = tmp17638

	tmp17639 := PrimCons(symnumber, Nil)

	tmp17640 := PrimCons(sym_1_1_6, tmp17639)

	tmp17641 := PrimCons(symnumber, tmp17640)

	tmp17642 := PrimCons(tmp17641, Nil)

	tmp17643 := PrimCons(sym_1_1_6, tmp17642)

	tmp17644 := PrimCons(symnumber, tmp17643)

	tmp17645 := Call(__e, PrimNS2Value(symdeclare), sym_c, tmp17644)

	_ = tmp17645

	tmp17646 := PrimCons(symnumber, Nil)

	tmp17647 := PrimCons(sym_1_1_6, tmp17646)

	tmp17648 := PrimCons(symnumber, tmp17647)

	tmp17649 := PrimCons(tmp17648, Nil)

	tmp17650 := PrimCons(sym_1_1_6, tmp17649)

	tmp17651 := PrimCons(symnumber, tmp17650)

	tmp17652 := Call(__e, PrimNS2Value(symdeclare), sym_1, tmp17651)

	_ = tmp17652

	tmp17653 := PrimCons(symnumber, Nil)

	tmp17654 := PrimCons(sym_1_1_6, tmp17653)

	tmp17655 := PrimCons(symnumber, tmp17654)

	tmp17656 := PrimCons(tmp17655, Nil)

	tmp17657 := PrimCons(sym_1_1_6, tmp17656)

	tmp17658 := PrimCons(symnumber, tmp17657)

	tmp17659 := Call(__e, PrimNS2Value(symdeclare), sym_d, tmp17658)

	_ = tmp17659

	tmp17660 := PrimCons(symboolean, Nil)

	tmp17661 := PrimCons(sym_1_1_6, tmp17660)

	tmp17662 := PrimCons(symB, tmp17661)

	tmp17663 := PrimCons(tmp17662, Nil)

	tmp17664 := PrimCons(sym_1_1_6, tmp17663)

	tmp17665 := PrimCons(symA, tmp17664)

	__e.TailApply(PrimNS2Value(symdeclare), sym_a_a, tmp17665)
	return

}, 0)
