package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
	tmp6511 := MakeNative(func(__e *ControlFlow) {
		tmp6512 := MakeNative(func(__e *ControlFlow) {
			Z880 := __e.Get(1)
			_ = Z880
			__e.TailApply(PrimFunc(symshen_4typename), Z880)
			return
		}, 1)

		tmp6513 := PrimValue(symshen_4_dalldatatypes_d)

		__e.TailApply(PrimFunc(symmap), tmp6512, tmp6513)
		return

	}, 0)

	tmp6514 := Call(__e, ns2_1set, symdatatypes, tmp6511)

	_ = tmp6514

	tmp6515 := MakeNative(func(__e *ControlFlow) {
		tmp6516 := MakeNative(func(__e *ControlFlow) {
			Z881 := __e.Get(1)
			_ = Z881
			__e.TailApply(PrimFunc(symshen_4typename), Z881)
			return
		}, 1)

		tmp6517 := PrimValue(symshen_4_ddatatypes_d)

		__e.TailApply(PrimFunc(symmap), tmp6516, tmp6517)
		return

	}, 0)

	tmp6518 := Call(__e, ns2_1set, symshen_4included, tmp6515)

	_ = tmp6518

	tmp6519 := MakeNative(func(__e *ControlFlow) {
		V884 := __e.Get(1)
		_ = V884
		tmp6524 := PrimIsPair(V884)

		if True == tmp6524 {
			tmp6520 := PrimHead(V884)

			tmp6521 := PrimStr(tmp6520)

			tmp6522 := Call(__e, PrimFunc(symshen_4typename_1h), tmp6521)

			__e.Return(PrimIntern(tmp6522))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename)
			return
		}

	}, 1)

	tmp6525 := Call(__e, ns2_1set, symshen_4typename, tmp6519)

	_ = tmp6525

	tmp6526 := MakeNative(func(__e *ControlFlow) {
		V885 := __e.Get(1)
		_ = V885
		tmp6533 := PrimEqual(MakeString("#type"), V885)

		if True == tmp6533 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp6531 := Call(__e, PrimFunc(symshen_4_7string_2), V885)

			if True == tmp6531 {
				tmp6527 := Call(__e, PrimFunc(symhdstr), V885)

				tmp6528 := PrimTailString(V885)

				tmp6529 := Call(__e, PrimFunc(symshen_4typename_1h), tmp6528)

				__e.Return(PrimStringConcat(tmp6527, tmp6529))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename_1h)
				return
			}

		}

	}, 1)

	tmp6534 := Call(__e, ns2_1set, symshen_4typename_1h, tmp6526)

	_ = tmp6534

	tmp6535 := MakeNative(func(__e *ControlFlow) {
		V886 := __e.Get(1)
		_ = V886
		tmp6539 := PrimLessThan(V886, MakeNumber(0))

		if True == tmp6539 {
			__e.Return(PrimValue(symshen_4_dprolog_1memory_d))
			return
		} else {
			tmp6537 := PrimIsInteger(V886)

			if True == tmp6537 {
				__e.Return(PrimSet(symshen_4_dprolog_1memory_d, V886))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("prolog memory expects an integer value\n")))
				return
			}

		}

	}, 1)

	tmp6540 := Call(__e, ns2_1set, symprolog_1memory, tmp6535)

	_ = tmp6540

	tmp6541 := MakeNative(func(__e *ControlFlow) {
		V887 := __e.Get(1)
		_ = V887
		tmp6542 := MakeNative(func(__e *ControlFlow) {
			tmp6543 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symget), V887, symarity, tmp6543)
			return

		}, 0)

		tmp6544 := MakeNative(func(__e *ControlFlow) {
			Z888 := __e.Get(1)
			_ = Z888
			__e.Return(MakeNumber(-1))
			return
		}, 1)

		__e.TailApply(try_1catch, tmp6542, tmp6544)
		return

	}, 1)

	tmp6545 := Call(__e, ns2_1set, symarity, tmp6541)

	_ = tmp6545

	tmp6546 := MakeNative(func(__e *ControlFlow) {
		V891 := __e.Get(1)
		_ = V891
		tmp6562 := PrimEqual(Nil, V891)

		if True == tmp6562 {
			__e.Return(Nil)
			return
		} else {
			tmp6560 := PrimIsPair(V891)

			var ifres6556 Obj

			if True == tmp6560 {
				tmp6558 := PrimTail(V891)

				tmp6559 := PrimIsPair(tmp6558)

				var ifres6557 Obj

				if True == tmp6559 {
					ifres6557 = True

				} else {
					ifres6557 = False

				}

				ifres6556 = ifres6557

			} else {
				ifres6556 = False

			}

			if True == ifres6556 {
				tmp6547 := MakeNative(func(__e *ControlFlow) {
					W892 := __e.Get(1)
					_ = W892
					tmp6548 := PrimTail(V891)

					tmp6549 := PrimTail(tmp6548)

					__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp6549)
					return

				}, 1)

				tmp6550 := PrimHead(V891)

				tmp6551 := PrimTail(V891)

				tmp6552 := PrimHead(tmp6551)

				tmp6553 := PrimValue(sym_dproperty_1vector_d)

				tmp6554 := Call(__e, PrimFunc(symput), tmp6550, symarity, tmp6552, tmp6553)

				__e.TailApply(tmp6547, tmp6554)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
				return
			}

		}

	}, 1)

	tmp6563 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp6546)

	_ = tmp6563

	tmp6564 := MakeNative(func(__e *ControlFlow) {
		V893 := __e.Get(1)
		_ = V893
		tmp6565 := MakeNative(func(__e *ControlFlow) {
			W894 := __e.Get(1)
			_ = W894
			tmp6566 := MakeNative(func(__e *ControlFlow) {
				W895 := __e.Get(1)
				_ = W895
				__e.Return(V893)
				return
			}, 1)

			tmp6567 := Call(__e, PrimFunc(symadjoin), V893, W894)

			tmp6568 := PrimValue(sym_dproperty_1vector_d)

			tmp6569 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp6567, tmp6568)

			__e.TailApply(tmp6566, tmp6569)
			return

		}, 1)

		tmp6570 := PrimValue(sym_dproperty_1vector_d)

		tmp6571 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp6570)

		__e.TailApply(tmp6565, tmp6571)
		return

	}, 1)

	tmp6572 := Call(__e, ns2_1set, symsystemf, tmp6564)

	_ = tmp6572

	tmp6573 := MakeNative(func(__e *ControlFlow) {
		V896 := __e.Get(1)
		_ = V896
		V897 := __e.Get(2)
		_ = V897
		tmp6575 := Call(__e, PrimFunc(symelement_2), V896, V897)

		if True == tmp6575 {
			__e.Return(V897)
			return
		} else {
			__e.Return(PrimCons(V896, V897))
			return
		}

	}, 2)

	tmp6576 := Call(__e, ns2_1set, symadjoin, tmp6573)

	_ = tmp6576

	tmp6577 := MakeNative(func(__e *ControlFlow) {
		V898 := __e.Get(1)
		_ = V898
		tmp6578 := MakeNative(func(__e *ControlFlow) {
			W899 := __e.Get(1)
			_ = W899
			tmp6586 := PrimEqual(W899, MakeNumber(-1))

			var ifres6583 Obj

			if True == tmp6586 {
				ifres6583 = True

			} else {
				tmp6585 := PrimEqual(W899, MakeNumber(0))

				var ifres6584 Obj

				if True == tmp6585 {
					ifres6584 = True

				} else {
					ifres6584 = False

				}

				ifres6583 = ifres6584

			}

			if True == ifres6583 {
				__e.Return(Nil)
				return
			} else {
				tmp6579 := PrimCons(V898, Nil)

				tmp6580 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp6579, W899)

				tmp6581 := Call(__e, PrimFunc(symeval_1kl), tmp6580)

				__e.Return(PrimCons(V898, tmp6581))
				return

			}

		}, 1)

		tmp6587 := Call(__e, PrimFunc(symarity), V898)

		__e.TailApply(tmp6578, tmp6587)
		return

	}, 1)

	tmp6588 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp6577)

	_ = tmp6588

	tmp6589 := MakeNative(func(__e *ControlFlow) {
		V900 := __e.Get(1)
		_ = V900
		tmp6594 := PrimIsPair(V900)

		if True == tmp6594 {
			tmp6590 := PrimHead(V900)

			tmp6591 := PrimTail(V900)

			tmp6592 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symput), tmp6590, symshen_4lambda_1form, tmp6591, tmp6592)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
			return
		}

	}, 1)

	tmp6595 := Call(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp6589)

	_ = tmp6595

	tmp6596 := MakeNative(func(__e *ControlFlow) {
		V901 := __e.Get(1)
		_ = V901
		tmp6597 := MakeNative(func(__e *ControlFlow) {
			W902 := __e.Get(1)
			_ = W902
			tmp6598 := MakeNative(func(__e *ControlFlow) {
				Z904 := __e.Get(1)
				_ = Z904
				__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), Z904)
				return
			}, 1)

			tmp6599 := MakeNative(func(__e *ControlFlow) {
				Z905 := __e.Get(1)
				_ = Z905
				__e.TailApply(PrimFunc(symshen_4tuple), Z905)
				return
			}, 1)

			tmp6600 := PrimCons(symshen_4tuple, tmp6599)

			tmp6601 := MakeNative(func(__e *ControlFlow) {
				Z906 := __e.Get(1)
				_ = Z906
				__e.TailApply(PrimFunc(symshen_4pvar), Z906)
				return
			}, 1)

			tmp6602 := PrimCons(symshen_4pvar, tmp6601)

			tmp6603 := MakeNative(func(__e *ControlFlow) {
				Z907 := __e.Get(1)
				_ = Z907
				__e.TailApply(PrimFunc(symshen_4dictionary), Z907)
				return
			}, 1)

			tmp6604 := PrimCons(symshen_4dictionary, tmp6603)

			tmp6605 := MakeNative(func(__e *ControlFlow) {
				Z908 := __e.Get(1)
				_ = Z908
				__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z908)
				return
			}, 1)

			tmp6606 := PrimCons(symshen_4print_1prolog_1vector, tmp6605)

			tmp6607 := MakeNative(func(__e *ControlFlow) {
				Z909 := __e.Get(1)
				_ = Z909
				__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z909)
				return
			}, 1)

			tmp6608 := PrimCons(symshen_4print_1freshterm, tmp6607)

			tmp6609 := MakeNative(func(__e *ControlFlow) {
				Z910 := __e.Get(1)
				_ = Z910
				__e.TailApply(PrimFunc(symshen_4printF), Z910)
				return
			}, 1)

			tmp6610 := PrimCons(symshen_4printF, tmp6609)

			tmp6611 := PrimCons(tmp6610, W902)

			tmp6612 := PrimCons(tmp6608, tmp6611)

			tmp6613 := PrimCons(tmp6606, tmp6612)

			tmp6614 := PrimCons(tmp6604, tmp6613)

			tmp6615 := PrimCons(tmp6602, tmp6614)

			tmp6616 := PrimCons(tmp6600, tmp6615)

			__e.TailApply(PrimFunc(symshen_4for_1each), tmp6598, tmp6616)
			return

		}, 1)

		tmp6617 := MakeNative(func(__e *ControlFlow) {
			Z903 := __e.Get(1)
			_ = Z903
			__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z903)
			return
		}, 1)

		tmp6618 := Call(__e, PrimFunc(symmap), tmp6617, V901)

		__e.TailApply(tmp6597, tmp6618)
		return

	}, 1)

	__e.TailApply(ns2_1set, symshen_4build_1lambda_1table, tmp6596)
	return

}, 0)
