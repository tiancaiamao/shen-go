package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
	tmp3634 := MakeNative(func(__e *ControlFlow) {
		V2504 := __e.Get(1)
		_ = V2504
		tmp3635 := MakeNative(func(__e *ControlFlow) {
			W2505 := __e.Get(1)
			_ = W2505
			tmp3636 := MakeNative(func(__e *ControlFlow) {
				W2506 := __e.Get(1)
				_ = W2506
				tmp3637 := MakeNative(func(__e *ControlFlow) {
					W2509 := __e.Get(1)
					_ = W2509
					__e.Return(W2509)
					return
				}, 1)

				tmp3638 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2506)

				__e.TailApply(tmp3637, tmp3638)
				return

			}, 1)

			tmp3639 := MakeNative(func(__e *ControlFlow) {
				tmp3640 := MakeNative(func(__e *ControlFlow) {
					Z2507 := __e.Get(1)
					_ = Z2507
					__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2507)
					return
				}, 1)

				__e.TailApply(PrimFunc(symcompile), tmp3640, W2505)
				return

			}, 0)

			tmp3641 := MakeNative(func(__e *ControlFlow) {
				Z2508 := __e.Get(1)
				_ = Z2508
				tmp3642 := PrimValue(symshen_4_dresidue_d)

				__e.TailApply(PrimFunc(symshen_4reader_1error), tmp3642)
				return

			}, 1)

			tmp3643 := Call(__e, try_1catch, tmp3639, tmp3641)

			__e.TailApply(tmp3636, tmp3643)
			return

		}, 1)

		tmp3644 := PrimReadFileAsByteList(V2504)

		__e.TailApply(tmp3635, tmp3644)
		return

	}, 1)

	tmp3645 := Call(__e, ns2_1set, symread_1file, tmp3634)

	_ = tmp3645

	tmp3646 := MakeNative(func(__e *ControlFlow) {
		V2510 := __e.Get(1)
		_ = V2510
		tmp3647 := PrimValue(sym_dmaximum_1print_1sequence_1size_d)

		tmp3648 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp3647, MakeNumber(0), V2510)

		tmp3649 := PrimStringConcat(MakeString("reader error near here: "), tmp3648)

		tmp3650 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp3649)

		__e.Return(PrimSimpleError(tmp3650))
		return

	}, 1)

	tmp3651 := Call(__e, ns2_1set, symshen_4reader_1error, tmp3646)

	_ = tmp3651

	tmp3652 := MakeNative(func(__e *ControlFlow) {
		V2518 := __e.Get(1)
		_ = V2518
		V2519 := __e.Get(2)
		_ = V2519
		V2520 := __e.Get(3)
		_ = V2520
		tmp3663 := PrimEqual(Nil, V2520)

		if True == tmp3663 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp3661 := PrimEqual(V2518, V2519)

			if True == tmp3661 {
				__e.Return(MakeString(""))
				return
			} else {
				tmp3659 := PrimIsPair(V2520)

				if True == tmp3659 {
					tmp3653 := PrimHead(V2520)

					tmp3654 := PrimNumberToString(tmp3653)

					tmp3655 := PrimNumberAdd(V2519, MakeNumber(1))

					tmp3656 := PrimTail(V2520)

					tmp3657 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2518, tmp3655, tmp3656)

					__e.Return(PrimStringConcat(tmp3654, tmp3657))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4reader_1error_1message)
					return
				}

			}

		}

	}, 3)

	tmp3664 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp3652)

	_ = tmp3664

	tmp3665 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dit_d))
		return
	}, 0)

	tmp3666 := Call(__e, ns2_1set, symit, tmp3665)

	_ = tmp3666

	tmp3667 := MakeNative(func(__e *ControlFlow) {
		V2521 := __e.Get(1)
		_ = V2521
		tmp3668 := MakeNative(func(__e *ControlFlow) {
			W2522 := __e.Get(1)
			_ = W2522
			tmp3669 := MakeNative(func(__e *ControlFlow) {
				W2523 := __e.Get(1)
				_ = W2523
				tmp3670 := MakeNative(func(__e *ControlFlow) {
					W2524 := __e.Get(1)
					_ = W2524
					tmp3671 := MakeNative(func(__e *ControlFlow) {
						W2525 := __e.Get(1)
						_ = W2525
						__e.TailApply(PrimFunc(symreverse), W2524)
						return
					}, 1)

					tmp3672 := PrimCloseStream(W2522)

					__e.TailApply(tmp3671, tmp3672)
					return

				}, 1)

				tmp3673 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), W2522, W2523, Nil)

				__e.TailApply(tmp3670, tmp3673)
				return

			}, 1)

			tmp3674 := PrimReadByte(W2522)

			__e.TailApply(tmp3669, tmp3674)
			return

		}, 1)

		tmp3675 := PrimOpenStream(V2521, symin)

		__e.TailApply(tmp3668, tmp3675)
		return

	}, 1)

	tmp3676 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp3667)

	_ = tmp3676

	tmp3677 := MakeNative(func(__e *ControlFlow) {
		V2526 := __e.Get(1)
		_ = V2526
		V2527 := __e.Get(2)
		_ = V2527
		V2528 := __e.Get(3)
		_ = V2528
		tmp3681 := PrimEqual(MakeNumber(-1), V2527)

		if True == tmp3681 {
			__e.Return(V2528)
			return
		} else {
			tmp3678 := PrimReadByte(V2526)

			tmp3679 := PrimCons(V2527, V2528)

			__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2526, tmp3678, tmp3679)
			return

		}

	}, 3)

	tmp3682 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp3677)

	_ = tmp3682

	tmp3683 := MakeNative(func(__e *ControlFlow) {
		V2529 := __e.Get(1)
		_ = V2529
		tmp3684 := MakeNative(func(__e *ControlFlow) {
			W2530 := __e.Get(1)
			_ = W2530
			tmp3685 := PrimReadByte(W2530)

			__e.TailApply(PrimFunc(symshen_4rfas_1h), W2530, tmp3685, MakeString(""))
			return

		}, 1)

		tmp3686 := PrimOpenStream(V2529, symin)

		__e.TailApply(tmp3684, tmp3686)
		return

	}, 1)

	tmp3687 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp3683)

	_ = tmp3687

	tmp3688 := MakeNative(func(__e *ControlFlow) {
		V2531 := __e.Get(1)
		_ = V2531
		V2532 := __e.Get(2)
		_ = V2532
		V2533 := __e.Get(3)
		_ = V2533
		tmp3694 := PrimEqual(MakeNumber(-1), V2532)

		if True == tmp3694 {
			tmp3689 := PrimCloseStream(V2531)

			_ = tmp3689

			__e.Return(V2533)
			return

		} else {
			tmp3690 := PrimReadByte(V2531)

			tmp3691 := PrimNumberToString(V2532)

			tmp3692 := PrimStringConcat(V2533, tmp3691)

			__e.TailApply(PrimFunc(symshen_4rfas_1h), V2531, tmp3690, tmp3692)
			return

		}

	}, 3)

	tmp3695 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp3688)

	_ = tmp3695

	tmp3696 := MakeNative(func(__e *ControlFlow) {
		V2534 := __e.Get(1)
		_ = V2534
		tmp3697 := Call(__e, PrimFunc(symread), V2534)

		__e.TailApply(PrimFunc(symeval_1kl), tmp3697)
		return

	}, 1)

	tmp3698 := Call(__e, ns2_1set, syminput, tmp3696)

	_ = tmp3698

	tmp3699 := MakeNative(func(__e *ControlFlow) {
		V2535 := __e.Get(1)
		_ = V2535
		V2536 := __e.Get(2)
		_ = V2536
		tmp3700 := MakeNative(func(__e *ControlFlow) {
			W2537 := __e.Get(1)
			_ = W2537
			tmp3701 := MakeNative(func(__e *ControlFlow) {
				W2538 := __e.Get(1)
				_ = W2538
				tmp3707 := Call(__e, PrimFunc(symshen_4rectify_1type), V2535)

				tmp3708 := Call(__e, PrimFunc(symshen_4typecheck), W2538, tmp3707)

				tmp3709 := PrimEqual(False, tmp3708)

				if True == tmp3709 {
					tmp3702 := Call(__e, PrimFunc(symshen_4app), V2535, MakeString("\n"), symshen_4r)

					tmp3703 := PrimStringConcat(MakeString(" is not of type "), tmp3702)

					tmp3704 := Call(__e, PrimFunc(symshen_4app), W2538, tmp3703, symshen_4r)

					tmp3705 := PrimStringConcat(MakeString("type error: "), tmp3704)

					__e.Return(PrimSimpleError(tmp3705))
					return

				} else {
					__e.TailApply(PrimFunc(symeval_1kl), W2538)
					return
				}

			}, 1)

			tmp3710 := Call(__e, PrimFunc(symread), V2536)

			__e.TailApply(tmp3701, tmp3710)
			return

		}, 1)

		tmp3711 := Call(__e, PrimFunc(symshen_4monotype), V2535)

		__e.TailApply(tmp3700, tmp3711)
		return

	}, 2)

	tmp3712 := Call(__e, ns2_1set, syminput_7, tmp3699)

	_ = tmp3712

	tmp3713 := MakeNative(func(__e *ControlFlow) {
		V2539 := __e.Get(1)
		_ = V2539
		tmp3720 := PrimIsPair(V2539)

		if True == tmp3720 {
			tmp3714 := MakeNative(func(__e *ControlFlow) {
				Z2540 := __e.Get(1)
				_ = Z2540
				__e.TailApply(PrimFunc(symshen_4monotype), Z2540)
				return
			}, 1)

			__e.TailApply(PrimFunc(symmap), tmp3714, V2539)
			return

		} else {
			tmp3718 := PrimIsVariable(V2539)

			if True == tmp3718 {
				tmp3715 := Call(__e, PrimFunc(symshen_4app), V2539, MakeString("\n"), symshen_4a)

				tmp3716 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp3715)

				__e.Return(PrimSimpleError(tmp3716))
				return

			} else {
				__e.Return(V2539)
				return
			}

		}

	}, 1)

	tmp3721 := Call(__e, ns2_1set, symshen_4monotype, tmp3713)

	_ = tmp3721

	tmp3722 := MakeNative(func(__e *ControlFlow) {
		V2541 := __e.Get(1)
		_ = V2541
		tmp3723 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2541)

		tmp3724 := MakeNative(func(__e *ControlFlow) {
			Z2542 := __e.Get(1)
			_ = Z2542
			__e.TailApply(PrimFunc(symshen_4return_2), Z2542)
			return
		}, 1)

		__e.TailApply(PrimFunc(symshen_4read_1loop), V2541, tmp3723, Nil, tmp3724)
		return

	}, 1)

	tmp3725 := Call(__e, ns2_1set, symlineread, tmp3722)

	_ = tmp3725

	tmp3726 := MakeNative(func(__e *ControlFlow) {
		V2543 := __e.Get(1)
		_ = V2543
		tmp3727 := MakeNative(func(__e *ControlFlow) {
			W2544 := __e.Get(1)
			_ = W2544
			tmp3728 := MakeNative(func(__e *ControlFlow) {
				W2545 := __e.Get(1)
				_ = W2545
				tmp3729 := MakeNative(func(__e *ControlFlow) {
					W2547 := __e.Get(1)
					_ = W2547
					__e.Return(W2547)
					return
				}, 1)

				tmp3730 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2545)

				__e.TailApply(tmp3729, tmp3730)
				return

			}, 1)

			tmp3731 := MakeNative(func(__e *ControlFlow) {
				Z2546 := __e.Get(1)
				_ = Z2546
				__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2546)
				return
			}, 1)

			tmp3732 := Call(__e, PrimFunc(symcompile), tmp3731, W2544)

			__e.TailApply(tmp3728, tmp3732)
			return

		}, 1)

		tmp3733 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2543)

		__e.TailApply(tmp3727, tmp3733)
		return

	}, 1)

	tmp3734 := Call(__e, ns2_1set, symread_1from_1string, tmp3726)

	_ = tmp3734

	tmp3735 := MakeNative(func(__e *ControlFlow) {
		V2548 := __e.Get(1)
		_ = V2548
		tmp3736 := MakeNative(func(__e *ControlFlow) {
			W2549 := __e.Get(1)
			_ = W2549
			tmp3737 := MakeNative(func(__e *ControlFlow) {
				W2550 := __e.Get(1)
				_ = W2550
				__e.Return(W2550)
				return
			}, 1)

			tmp3738 := MakeNative(func(__e *ControlFlow) {
				Z2551 := __e.Get(1)
				_ = Z2551
				__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2551)
				return
			}, 1)

			tmp3739 := Call(__e, PrimFunc(symcompile), tmp3738, W2549)

			__e.TailApply(tmp3737, tmp3739)
			return

		}, 1)

		tmp3740 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2548)

		__e.TailApply(tmp3736, tmp3740)
		return

	}, 1)

	tmp3741 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp3735)

	_ = tmp3741

	tmp3742 := MakeNative(func(__e *ControlFlow) {
		V2552 := __e.Get(1)
		_ = V2552
		tmp3750 := PrimEqual(MakeString(""), V2552)

		if True == tmp3750 {
			__e.Return(Nil)
			return
		} else {
			tmp3748 := Call(__e, PrimFunc(symshen_4_7string_2), V2552)

			if True == tmp3748 {
				tmp3743 := Call(__e, PrimFunc(symhdstr), V2552)

				tmp3744 := PrimStringToNumber(tmp3743)

				tmp3745 := PrimTailString(V2552)

				tmp3746 := Call(__e, PrimFunc(symshen_4str_1_6bytes), tmp3745)

				__e.Return(PrimCons(tmp3744, tmp3746))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4str_1_6bytes)
				return
			}

		}

	}, 1)

	tmp3751 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp3742)

	_ = tmp3751

	tmp3752 := MakeNative(func(__e *ControlFlow) {
		V2553 := __e.Get(1)
		_ = V2553
		tmp3753 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2553)

		tmp3754 := MakeNative(func(__e *ControlFlow) {
			Z2554 := __e.Get(1)
			_ = Z2554
			__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2554)
			return
		}, 1)

		tmp3755 := Call(__e, PrimFunc(symshen_4read_1loop), V2553, tmp3753, Nil, tmp3754)

		__e.Return(PrimHead(tmp3755))
		return

	}, 1)

	tmp3756 := Call(__e, ns2_1set, symread, tmp3752)

	_ = tmp3756

	tmp3757 := MakeNative(func(__e *ControlFlow) {
		V2555 := __e.Get(1)
		_ = V2555
		tmp3760 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2555)

		if True == tmp3760 {
			tmp3758 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2555)

			__e.Return(PrimStringToNumber(tmp3758))
			return

		} else {
			__e.Return(PrimReadByte(V2555))
			return
		}

	}, 1)

	tmp3761 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp3757)

	_ = tmp3761

	tmp3762 := MakeNative(func(__e *ControlFlow) {
		V2560 := __e.Get(1)
		_ = V2560
		V2561 := __e.Get(2)
		_ = V2561
		V2562 := __e.Get(3)
		_ = V2562
		V2563 := __e.Get(4)
		_ = V2563
		tmp3785 := PrimEqual(MakeNumber(94), V2561)

		if True == tmp3785 {
			__e.Return(PrimSimpleError(MakeString("read aborted")))
			return
		} else {
			tmp3783 := PrimEqual(MakeNumber(-1), V2561)

			if True == tmp3783 {
				tmp3765 := Call(__e, PrimFunc(symempty_2), V2562)

				if True == tmp3765 {
					__e.Return(PrimSimpleError(MakeString("error: empty stream")))
					return
				} else {
					tmp3763 := MakeNative(func(__e *ControlFlow) {
						Z2564 := __e.Get(1)
						_ = Z2564
						__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2564)
						return
					}, 1)

					__e.TailApply(PrimFunc(symcompile), tmp3763, V2562)
					return

				}

			} else {
				tmp3781 := PrimEqual(MakeNumber(0), V2561)

				if True == tmp3781 {
					tmp3766 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2560)

					__e.TailApply(PrimFunc(symshen_4read_1loop), V2560, tmp3766, V2562, V2563)
					return

				} else {
					tmp3779 := Call(__e, V2563, V2561)

					if True == tmp3779 {
						tmp3767 := MakeNative(func(__e *ControlFlow) {
							W2565 := __e.Get(1)
							_ = W2565
							tmp3773 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2565)

							if True == tmp3773 {
								tmp3768 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2560)

								tmp3769 := PrimCons(V2561, Nil)

								tmp3770 := Call(__e, PrimFunc(symappend), V2562, tmp3769)

								__e.TailApply(PrimFunc(symshen_4read_1loop), V2560, tmp3768, tmp3770, V2563)
								return

							} else {
								tmp3771 := Call(__e, PrimFunc(symshen_4record_1it), V2562)

								_ = tmp3771

								__e.Return(W2565)
								return

							}

						}, 1)

						tmp3774 := Call(__e, PrimFunc(symshen_4try_1parse), V2562)

						__e.TailApply(tmp3767, tmp3774)
						return

					} else {
						tmp3775 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2560)

						tmp3776 := PrimCons(V2561, Nil)

						tmp3777 := Call(__e, PrimFunc(symappend), V2562, tmp3776)

						__e.TailApply(PrimFunc(symshen_4read_1loop), V2560, tmp3775, tmp3777, V2563)
						return

					}

				}

			}

		}

	}, 4)

	tmp3786 := Call(__e, ns2_1set, symshen_4read_1loop, tmp3762)

	_ = tmp3786

	tmp3787 := MakeNative(func(__e *ControlFlow) {
		V2566 := __e.Get(1)
		_ = V2566
		tmp3788 := MakeNative(func(__e *ControlFlow) {
			W2567 := __e.Get(1)
			_ = W2567
			tmp3790 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2567)

			if True == tmp3790 {
				__e.Return(symshen_4i_1failed_b)
				return
			} else {
				__e.TailApply(PrimFunc(symshen_4process_1sexprs), W2567)
				return
			}

		}, 1)

		tmp3791 := MakeNative(func(__e *ControlFlow) {
			tmp3792 := MakeNative(func(__e *ControlFlow) {
				Z2568 := __e.Get(1)
				_ = Z2568
				__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2568)
				return
			}, 1)

			__e.TailApply(PrimFunc(symcompile), tmp3792, V2566)
			return

		}, 0)

		tmp3793 := MakeNative(func(__e *ControlFlow) {
			Z2569 := __e.Get(1)
			_ = Z2569
			__e.Return(symshen_4i_1failed_b)
			return
		}, 1)

		tmp3794 := Call(__e, try_1catch, tmp3791, tmp3793)

		__e.TailApply(tmp3788, tmp3794)
		return

	}, 1)

	tmp3795 := Call(__e, ns2_1set, symshen_4try_1parse, tmp3787)

	_ = tmp3795

	tmp3796 := MakeNative(func(__e *ControlFlow) {
		V2572 := __e.Get(1)
		_ = V2572
		tmp3800 := PrimEqual(symshen_4i_1failed_b, V2572)

		if True == tmp3800 {
			__e.Return(True)
			return
		} else {
			tmp3798 := PrimEqual(Nil, V2572)

			if True == tmp3798 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp3801 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp3796)

	_ = tmp3801

	tmp3802 := MakeNative(func(__e *ControlFlow) {
		V2573 := __e.Get(1)
		_ = V2573
		tmp3803 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2573)

		__e.Return(PrimSet(symshen_4_dit_d, tmp3803))
		return

	}, 1)

	tmp3804 := Call(__e, ns2_1set, symshen_4record_1it, tmp3802)

	_ = tmp3804

	tmp3805 := MakeNative(func(__e *ControlFlow) {
		V2574 := __e.Get(1)
		_ = V2574
		tmp3813 := PrimEqual(Nil, V2574)

		if True == tmp3813 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp3811 := PrimIsPair(V2574)

			if True == tmp3811 {
				tmp3806 := PrimHead(V2574)

				tmp3807 := PrimNumberToString(tmp3806)

				tmp3808 := PrimTail(V2574)

				tmp3809 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp3808)

				__e.Return(PrimStringConcat(tmp3807, tmp3809))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bytes_1_6string)
				return
			}

		}

	}, 1)

	tmp3814 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp3805)

	_ = tmp3814

	tmp3815 := MakeNative(func(__e *ControlFlow) {
		V2575 := __e.Get(1)
		_ = V2575
		tmp3816 := MakeNative(func(__e *ControlFlow) {
			W2576 := __e.Get(1)
			_ = W2576
			tmp3817 := MakeNative(func(__e *ControlFlow) {
				W2577 := __e.Get(1)
				_ = W2577
				tmp3818 := MakeNative(func(__e *ControlFlow) {
					W2578 := __e.Get(1)
					_ = W2578
					tmp3819 := MakeNative(func(__e *ControlFlow) {
						Z2579 := __e.Get(1)
						_ = Z2579
						__e.TailApply(PrimFunc(symshen_4process_1applications), Z2579, W2578)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp3819, W2576)
					return

				}, 1)

				tmp3820 := Call(__e, PrimFunc(symshen_4find_1types), W2576)

				__e.TailApply(tmp3818, tmp3820)
				return

			}, 1)

			tmp3821 := Call(__e, PrimFunc(symshen_4find_1arities), W2576)

			__e.TailApply(tmp3817, tmp3821)
			return

		}, 1)

		tmp3822 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2575)

		__e.TailApply(tmp3816, tmp3822)
		return

	}, 1)

	tmp3823 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp3815)

	_ = tmp3823

	tmp3824 := MakeNative(func(__e *ControlFlow) {
		V2580 := __e.Get(1)
		_ = V2580
		tmp3846 := PrimIsPair(V2580)

		var ifres3837 Obj

		if True == tmp3846 {
			tmp3844 := PrimTail(V2580)

			tmp3845 := PrimIsPair(tmp3844)

			var ifres3839 Obj

			if True == tmp3845 {
				tmp3841 := PrimHead(V2580)

				tmp3842 := PrimIntern(MakeString(":"))

				tmp3843 := PrimEqual(tmp3841, tmp3842)

				var ifres3840 Obj

				if True == tmp3843 {
					ifres3840 = True

				} else {
					ifres3840 = False

				}

				ifres3839 = ifres3840

			} else {
				ifres3839 = False

			}

			var ifres3838 Obj

			if True == ifres3839 {
				ifres3838 = True

			} else {
				ifres3838 = False

			}

			ifres3837 = ifres3838

		} else {
			ifres3837 = False

		}

		if True == ifres3837 {
			tmp3825 := PrimTail(V2580)

			tmp3826 := PrimHead(tmp3825)

			tmp3827 := PrimTail(V2580)

			tmp3828 := PrimTail(tmp3827)

			tmp3829 := Call(__e, PrimFunc(symshen_4find_1types), tmp3828)

			__e.Return(PrimCons(tmp3826, tmp3829))
			return

		} else {
			tmp3835 := PrimIsPair(V2580)

			if True == tmp3835 {
				tmp3830 := PrimHead(V2580)

				tmp3831 := Call(__e, PrimFunc(symshen_4find_1types), tmp3830)

				tmp3832 := PrimTail(V2580)

				tmp3833 := Call(__e, PrimFunc(symshen_4find_1types), tmp3832)

				__e.TailApply(PrimFunc(symappend), tmp3831, tmp3833)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp3847 := Call(__e, ns2_1set, symshen_4find_1types, tmp3824)

	_ = tmp3847

	tmp3848 := MakeNative(func(__e *ControlFlow) {
		V2583 := __e.Get(1)
		_ = V2583
		tmp3897 := PrimIsPair(V2583)

		var ifres3878 Obj

		if True == tmp3897 {
			tmp3895 := PrimHead(V2583)

			tmp3896 := PrimEqual(symdefine, tmp3895)

			var ifres3880 Obj

			if True == tmp3896 {
				tmp3893 := PrimTail(V2583)

				tmp3894 := PrimIsPair(tmp3893)

				var ifres3882 Obj

				if True == tmp3894 {
					tmp3890 := PrimTail(V2583)

					tmp3891 := PrimTail(tmp3890)

					tmp3892 := PrimIsPair(tmp3891)

					var ifres3884 Obj

					if True == tmp3892 {
						tmp3886 := PrimTail(V2583)

						tmp3887 := PrimTail(tmp3886)

						tmp3888 := PrimHead(tmp3887)

						tmp3889 := PrimEqual(sym_i, tmp3888)

						var ifres3885 Obj

						if True == tmp3889 {
							ifres3885 = True

						} else {
							ifres3885 = False

						}

						ifres3884 = ifres3885

					} else {
						ifres3884 = False

					}

					var ifres3883 Obj

					if True == ifres3884 {
						ifres3883 = True

					} else {
						ifres3883 = False

					}

					ifres3882 = ifres3883

				} else {
					ifres3882 = False

				}

				var ifres3881 Obj

				if True == ifres3882 {
					ifres3881 = True

				} else {
					ifres3881 = False

				}

				ifres3880 = ifres3881

			} else {
				ifres3880 = False

			}

			var ifres3879 Obj

			if True == ifres3880 {
				ifres3879 = True

			} else {
				ifres3879 = False

			}

			ifres3878 = ifres3879

		} else {
			ifres3878 = False

		}

		if True == ifres3878 {
			tmp3849 := PrimTail(V2583)

			tmp3850 := PrimHead(tmp3849)

			tmp3851 := PrimTail(V2583)

			tmp3852 := PrimHead(tmp3851)

			tmp3853 := PrimTail(V2583)

			tmp3854 := PrimTail(tmp3853)

			tmp3855 := PrimTail(tmp3854)

			tmp3856 := Call(__e, PrimFunc(symshen_4find_1arity), tmp3852, MakeNumber(1), tmp3855)

			__e.TailApply(PrimFunc(symshen_4store_1arity), tmp3850, tmp3856)
			return

		} else {
			tmp3876 := PrimIsPair(V2583)

			var ifres3868 Obj

			if True == tmp3876 {
				tmp3874 := PrimHead(V2583)

				tmp3875 := PrimEqual(symdefine, tmp3874)

				var ifres3870 Obj

				if True == tmp3875 {
					tmp3872 := PrimTail(V2583)

					tmp3873 := PrimIsPair(tmp3872)

					var ifres3871 Obj

					if True == tmp3873 {
						ifres3871 = True

					} else {
						ifres3871 = False

					}

					ifres3870 = ifres3871

				} else {
					ifres3870 = False

				}

				var ifres3869 Obj

				if True == ifres3870 {
					ifres3869 = True

				} else {
					ifres3869 = False

				}

				ifres3868 = ifres3869

			} else {
				ifres3868 = False

			}

			if True == ifres3868 {
				tmp3857 := PrimTail(V2583)

				tmp3858 := PrimHead(tmp3857)

				tmp3859 := PrimTail(V2583)

				tmp3860 := PrimHead(tmp3859)

				tmp3861 := PrimTail(V2583)

				tmp3862 := PrimTail(tmp3861)

				tmp3863 := Call(__e, PrimFunc(symshen_4find_1arity), tmp3860, MakeNumber(0), tmp3862)

				__e.TailApply(PrimFunc(symshen_4store_1arity), tmp3858, tmp3863)
				return

			} else {
				tmp3866 := PrimIsPair(V2583)

				if True == tmp3866 {
					tmp3864 := MakeNative(func(__e *ControlFlow) {
						Z2584 := __e.Get(1)
						_ = Z2584
						__e.TailApply(PrimFunc(symshen_4find_1arities), Z2584)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp3864, V2583)
					return

				} else {
					__e.Return(symshen_4skip)
					return
				}

			}

		}

	}, 1)

	tmp3898 := Call(__e, ns2_1set, symshen_4find_1arities, tmp3848)

	_ = tmp3898

	tmp3899 := MakeNative(func(__e *ControlFlow) {
		V2585 := __e.Get(1)
		_ = V2585
		V2586 := __e.Get(2)
		_ = V2586
		tmp3900 := MakeNative(func(__e *ControlFlow) {
			W2587 := __e.Get(1)
			_ = W2587
			tmp3911 := PrimEqual(W2587, MakeNumber(-1))

			if True == tmp3911 {
				__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2585, V2586)
				return
			} else {
				tmp3909 := PrimEqual(W2587, V2586)

				if True == tmp3909 {
					__e.Return(symshen_4skip)
					return
				} else {
					tmp3907 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2585)

					if True == tmp3907 {
						tmp3901 := Call(__e, PrimFunc(symshen_4app), V2585, MakeString(" is a system function\n"), symshen_4a)

						__e.Return(PrimSimpleError(tmp3901))
						return

					} else {
						tmp3902 := Call(__e, PrimFunc(symshen_4app), V2585, MakeString(" may cause errors\n"), symshen_4a)

						tmp3903 := PrimStringConcat(MakeString("changing the arity of "), tmp3902)

						tmp3904 := Call(__e, PrimFunc(symstoutput))

						tmp3905 := Call(__e, PrimFunc(sympr), tmp3903, tmp3904)

						_ = tmp3905

						__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2585, V2586)
						return

					}

				}

			}

		}, 1)

		tmp3912 := Call(__e, PrimFunc(symarity), V2585)

		__e.TailApply(tmp3900, tmp3912)
		return

	}, 2)

	tmp3913 := Call(__e, ns2_1set, symshen_4store_1arity, tmp3899)

	_ = tmp3913

	tmp3914 := MakeNative(func(__e *ControlFlow) {
		V2588 := __e.Get(1)
		_ = V2588
		V2589 := __e.Get(2)
		_ = V2589
		tmp3919 := PrimEqual(MakeNumber(0), V2589)

		if True == tmp3919 {
			tmp3915 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symput), V2588, symarity, MakeNumber(0), tmp3915)
			return

		} else {
			tmp3916 := PrimValue(sym_dproperty_1vector_d)

			tmp3917 := Call(__e, PrimFunc(symput), V2588, symarity, V2589, tmp3916)

			_ = tmp3917

			__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2588, V2589)
			return

		}

	}, 2)

	tmp3920 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp3914)

	_ = tmp3920

	tmp3921 := MakeNative(func(__e *ControlFlow) {
		V2590 := __e.Get(1)
		_ = V2590
		V2591 := __e.Get(2)
		_ = V2591
		tmp3922 := MakeNative(func(__e *ControlFlow) {
			W2592 := __e.Get(1)
			_ = W2592
			tmp3923 := MakeNative(func(__e *ControlFlow) {
				W2593 := __e.Get(1)
				_ = W2593
				__e.Return(V2591)
				return
			}, 1)

			tmp3924 := PrimCons(V2590, W2592)

			tmp3925 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3924)

			__e.TailApply(tmp3923, tmp3925)
			return

		}, 1)

		tmp3926 := PrimCons(V2590, Nil)

		tmp3927 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp3926, V2591)

		tmp3928 := Call(__e, PrimFunc(symeval_1kl), tmp3927)

		__e.TailApply(tmp3922, tmp3928)
		return

	}, 2)

	tmp3929 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp3921)

	_ = tmp3929

	tmp3930 := MakeNative(func(__e *ControlFlow) {
		V2596 := __e.Get(1)
		_ = V2596
		V2597 := __e.Get(2)
		_ = V2597
		tmp3948 := PrimEqual(MakeNumber(0), V2597)

		if True == tmp3948 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp3946 := PrimEqual(MakeNumber(1), V2597)

			if True == tmp3946 {
				tmp3931 := MakeNative(func(__e *ControlFlow) {
					W2598 := __e.Get(1)
					_ = W2598
					tmp3932 := PrimCons(W2598, Nil)

					tmp3933 := Call(__e, PrimFunc(symappend), V2596, tmp3932)

					tmp3934 := PrimCons(tmp3933, Nil)

					tmp3935 := PrimCons(W2598, tmp3934)

					__e.Return(PrimCons(symlambda, tmp3935))
					return

				}, 1)

				tmp3936 := Call(__e, PrimFunc(symgensym), symY)

				__e.TailApply(tmp3931, tmp3936)
				return

			} else {
				tmp3937 := MakeNative(func(__e *ControlFlow) {
					W2599 := __e.Get(1)
					_ = W2599
					tmp3938 := PrimCons(W2599, Nil)

					tmp3939 := Call(__e, PrimFunc(symappend), V2596, tmp3938)

					tmp3940 := PrimNumberSubtract(V2597, MakeNumber(1))

					tmp3941 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp3939, tmp3940)

					tmp3942 := PrimCons(tmp3941, Nil)

					tmp3943 := PrimCons(W2599, tmp3942)

					__e.Return(PrimCons(symlambda, tmp3943))
					return

				}, 1)

				tmp3944 := Call(__e, PrimFunc(symgensym), symY)

				__e.TailApply(tmp3937, tmp3944)
				return

			}

		}

	}, 2)

	tmp3949 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp3930)

	_ = tmp3949

	tmp3950 := MakeNative(func(__e *ControlFlow) {
		V2609 := __e.Get(1)
		_ = V2609
		V2610 := __e.Get(2)
		_ = V2610
		V2611 := __e.Get(3)
		_ = V2611
		tmp3973 := PrimEqual(Nil, V2611)

		if True == tmp3973 {
			tmp3951 := PrimCons(V2609, V2610)

			__e.Return(PrimCons(tmp3951, Nil))
			return

		} else {
			tmp3971 := PrimIsPair(V2611)

			var ifres3962 Obj

			if True == tmp3971 {
				tmp3969 := PrimHead(V2611)

				tmp3970 := PrimIsPair(tmp3969)

				var ifres3964 Obj

				if True == tmp3970 {
					tmp3966 := PrimHead(V2611)

					tmp3967 := PrimHead(tmp3966)

					tmp3968 := PrimEqual(V2609, tmp3967)

					var ifres3965 Obj

					if True == tmp3968 {
						ifres3965 = True

					} else {
						ifres3965 = False

					}

					ifres3964 = ifres3965

				} else {
					ifres3964 = False

				}

				var ifres3963 Obj

				if True == ifres3964 {
					ifres3963 = True

				} else {
					ifres3963 = False

				}

				ifres3962 = ifres3963

			} else {
				ifres3962 = False

			}

			if True == ifres3962 {
				tmp3952 := PrimHead(V2611)

				tmp3953 := PrimHead(tmp3952)

				tmp3954 := PrimCons(tmp3953, V2610)

				tmp3955 := PrimTail(V2611)

				__e.Return(PrimCons(tmp3954, tmp3955))
				return

			} else {
				tmp3960 := PrimIsPair(V2611)

				if True == tmp3960 {
					tmp3956 := PrimHead(V2611)

					tmp3957 := PrimTail(V2611)

					tmp3958 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2609, V2610, tmp3957)

					__e.Return(PrimCons(tmp3956, tmp3958))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
					return
				}

			}

		}

	}, 3)

	tmp3974 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp3950)

	_ = tmp3974

	tmp3975 := MakeNative(func(__e *ControlFlow) {
		V2626 := __e.Get(1)
		_ = V2626
		V2627 := __e.Get(2)
		_ = V2627
		V2628 := __e.Get(3)
		_ = V2628
		tmp4022 := PrimEqual(MakeNumber(0), V2627)

		var ifres4015 Obj

		if True == tmp4022 {
			tmp4021 := PrimIsPair(V2628)

			var ifres4017 Obj

			if True == tmp4021 {
				tmp4019 := PrimHead(V2628)

				tmp4020 := PrimEqual(tmp4019, sym_1_6)

				var ifres4018 Obj

				if True == tmp4020 {
					ifres4018 = True

				} else {
					ifres4018 = False

				}

				ifres4017 = ifres4018

			} else {
				ifres4017 = False

			}

			var ifres4016 Obj

			if True == ifres4017 {
				ifres4016 = True

			} else {
				ifres4016 = False

			}

			ifres4015 = ifres4016

		} else {
			ifres4015 = False

		}

		if True == ifres4015 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp4013 := PrimEqual(MakeNumber(0), V2627)

			var ifres4006 Obj

			if True == tmp4013 {
				tmp4012 := PrimIsPair(V2628)

				var ifres4008 Obj

				if True == tmp4012 {
					tmp4010 := PrimHead(V2628)

					tmp4011 := PrimEqual(tmp4010, sym_5_1)

					var ifres4009 Obj

					if True == tmp4011 {
						ifres4009 = True

					} else {
						ifres4009 = False

					}

					ifres4008 = ifres4009

				} else {
					ifres4008 = False

				}

				var ifres4007 Obj

				if True == ifres4008 {
					ifres4007 = True

				} else {
					ifres4007 = False

				}

				ifres4006 = ifres4007

			} else {
				ifres4006 = False

			}

			if True == ifres4006 {
				__e.Return(MakeNumber(0))
				return
			} else {
				tmp4004 := PrimEqual(MakeNumber(0), V2627)

				var ifres4001 Obj

				if True == tmp4004 {
					tmp4003 := PrimIsPair(V2628)

					var ifres4002 Obj

					if True == tmp4003 {
						ifres4002 = True

					} else {
						ifres4002 = False

					}

					ifres4001 = ifres4002

				} else {
					ifres4001 = False

				}

				if True == ifres4001 {
					tmp3976 := PrimTail(V2628)

					tmp3977 := Call(__e, PrimFunc(symshen_4find_1arity), V2626, MakeNumber(0), tmp3976)

					__e.Return(PrimNumberAdd(MakeNumber(1), tmp3977))
					return

				} else {
					tmp3999 := PrimEqual(MakeNumber(1), V2627)

					var ifres3992 Obj

					if True == tmp3999 {
						tmp3998 := PrimIsPair(V2628)

						var ifres3994 Obj

						if True == tmp3998 {
							tmp3996 := PrimHead(V2628)

							tmp3997 := PrimEqual(sym_j, tmp3996)

							var ifres3995 Obj

							if True == tmp3997 {
								ifres3995 = True

							} else {
								ifres3995 = False

							}

							ifres3994 = ifres3995

						} else {
							ifres3994 = False

						}

						var ifres3993 Obj

						if True == ifres3994 {
							ifres3993 = True

						} else {
							ifres3993 = False

						}

						ifres3992 = ifres3993

					} else {
						ifres3992 = False

					}

					if True == ifres3992 {
						tmp3978 := PrimTail(V2628)

						__e.TailApply(PrimFunc(symshen_4find_1arity), V2626, MakeNumber(0), tmp3978)
						return

					} else {
						tmp3990 := PrimEqual(MakeNumber(1), V2627)

						var ifres3987 Obj

						if True == tmp3990 {
							tmp3989 := PrimIsPair(V2628)

							var ifres3988 Obj

							if True == tmp3989 {
								ifres3988 = True

							} else {
								ifres3988 = False

							}

							ifres3987 = ifres3988

						} else {
							ifres3987 = False

						}

						if True == ifres3987 {
							tmp3979 := PrimTail(V2628)

							__e.TailApply(PrimFunc(symshen_4find_1arity), V2626, MakeNumber(1), tmp3979)
							return

						} else {
							tmp3985 := PrimEqual(MakeNumber(1), V2627)

							if True == tmp3985 {
								tmp3980 := Call(__e, PrimFunc(symshen_4app), V2626, MakeString(" definition: missing }\n"), symshen_4a)

								tmp3981 := PrimStringConcat(MakeString("syntax error in "), tmp3980)

								__e.Return(PrimSimpleError(tmp3981))
								return

							} else {
								tmp3982 := Call(__e, PrimFunc(symshen_4app), V2626, MakeString(" definition: missing -> or <-\n"), symshen_4a)

								tmp3983 := PrimStringConcat(MakeString("syntax error in "), tmp3982)

								__e.Return(PrimSimpleError(tmp3983))
								return

							}

						}

					}

				}

			}

		}

	}, 3)

	tmp4023 := Call(__e, ns2_1set, symshen_4find_1arity, tmp3975)

	_ = tmp4023

	tmp4024 := MakeNative(func(__e *ControlFlow) {
		V2629 := __e.Get(1)
		_ = V2629
		tmp4025 := MakeNative(func(__e *ControlFlow) {
			W2630 := __e.Get(1)
			_ = W2630
			tmp4270 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2630)

			if True == tmp4270 {
				tmp4026 := MakeNative(func(__e *ControlFlow) {
					W2641 := __e.Get(1)
					_ = W2641
					tmp4238 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2641)

					if True == tmp4238 {
						tmp4027 := MakeNative(func(__e *ControlFlow) {
							W2652 := __e.Get(1)
							_ = W2652
							tmp4220 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2652)

							if True == tmp4220 {
								tmp4028 := MakeNative(func(__e *ControlFlow) {
									W2658 := __e.Get(1)
									_ = W2658
									tmp4202 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2658)

									if True == tmp4202 {
										tmp4029 := MakeNative(func(__e *ControlFlow) {
											W2664 := __e.Get(1)
											_ = W2664
											tmp4184 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2664)

											if True == tmp4184 {
												tmp4030 := MakeNative(func(__e *ControlFlow) {
													W2670 := __e.Get(1)
													_ = W2670
													tmp4165 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2670)

													if True == tmp4165 {
														tmp4031 := MakeNative(func(__e *ControlFlow) {
															W2676 := __e.Get(1)
															_ = W2676
															tmp4140 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2676)

															if True == tmp4140 {
																tmp4032 := MakeNative(func(__e *ControlFlow) {
																	W2684 := __e.Get(1)
																	_ = W2684
																	tmp4121 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2684)

																	if True == tmp4121 {
																		tmp4033 := MakeNative(func(__e *ControlFlow) {
																			W2690 := __e.Get(1)
																			_ = W2690
																			tmp4102 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2690)

																			if True == tmp4102 {
																				tmp4034 := MakeNative(func(__e *ControlFlow) {
																					W2696 := __e.Get(1)
																					_ = W2696
																					tmp4085 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2696)

																					if True == tmp4085 {
																						tmp4035 := MakeNative(func(__e *ControlFlow) {
																							W2702 := __e.Get(1)
																							_ = W2702
																							tmp4065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2702)

																							if True == tmp4065 {
																								tmp4036 := MakeNative(func(__e *ControlFlow) {
																									W2709 := __e.Get(1)
																									_ = W2709
																									tmp4048 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2709)

																									if True == tmp4048 {
																										tmp4037 := MakeNative(func(__e *ControlFlow) {
																											W2715 := __e.Get(1)
																											_ = W2715
																											tmp4039 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2715)

																											if True == tmp4039 {
																												__e.TailApply(PrimFunc(symshen_4parse_1failure))
																												return
																											} else {
																												__e.Return(W2715)
																												return
																											}

																										}, 1)

																										tmp4040 := MakeNative(func(__e *ControlFlow) {
																											W2716 := __e.Get(1)
																											_ = W2716
																											tmp4044 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2716)

																											if True == tmp4044 {
																												__e.TailApply(PrimFunc(symshen_4parse_1failure))
																												return
																											} else {
																												tmp4041 := MakeNative(func(__e *ControlFlow) {
																													W2717 := __e.Get(1)
																													_ = W2717
																													__e.TailApply(PrimFunc(symshen_4comb), W2717, Nil)
																													return
																												}, 1)

																												tmp4042 := Call(__e, PrimFunc(symshen_4in_1_6), W2716)

																												__e.TailApply(tmp4041, tmp4042)
																												return

																											}

																										}, 1)

																										tmp4045 := Call(__e, PrimFunc(sym_5e_6), V2629)

																										tmp4046 := Call(__e, tmp4040, tmp4045)

																										__e.TailApply(tmp4037, tmp4046)
																										return

																									} else {
																										__e.Return(W2709)
																										return
																									}

																								}, 1)

																								tmp4049 := MakeNative(func(__e *ControlFlow) {
																									W2710 := __e.Get(1)
																									_ = W2710
																									tmp4061 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2710)

																									if True == tmp4061 {
																										__e.TailApply(PrimFunc(symshen_4parse_1failure))
																										return
																									} else {
																										tmp4050 := MakeNative(func(__e *ControlFlow) {
																											W2711 := __e.Get(1)
																											_ = W2711
																											tmp4051 := MakeNative(func(__e *ControlFlow) {
																												W2712 := __e.Get(1)
																												_ = W2712
																												tmp4057 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2712)

																												if True == tmp4057 {
																													__e.TailApply(PrimFunc(symshen_4parse_1failure))
																													return
																												} else {
																													tmp4052 := MakeNative(func(__e *ControlFlow) {
																														W2713 := __e.Get(1)
																														_ = W2713
																														tmp4053 := MakeNative(func(__e *ControlFlow) {
																															W2714 := __e.Get(1)
																															_ = W2714
																															__e.TailApply(PrimFunc(symshen_4comb), W2714, W2713)
																															return
																														}, 1)

																														tmp4054 := Call(__e, PrimFunc(symshen_4in_1_6), W2712)

																														__e.TailApply(tmp4053, tmp4054)
																														return

																													}, 1)

																													tmp4055 := Call(__e, PrimFunc(symshen_4_5_1out), W2712)

																													__e.TailApply(tmp4052, tmp4055)
																													return

																												}

																											}, 1)

																											tmp4058 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2711)

																											__e.TailApply(tmp4051, tmp4058)
																											return

																										}, 1)

																										tmp4059 := Call(__e, PrimFunc(symshen_4in_1_6), W2710)

																										__e.TailApply(tmp4050, tmp4059)
																										return

																									}

																								}, 1)

																								tmp4062 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2629)

																								tmp4063 := Call(__e, tmp4049, tmp4062)

																								__e.TailApply(tmp4036, tmp4063)
																								return

																							} else {
																								__e.Return(W2702)
																								return
																							}

																						}, 1)

																						tmp4066 := MakeNative(func(__e *ControlFlow) {
																							W2703 := __e.Get(1)
																							_ = W2703
																							tmp4081 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2703)

																							if True == tmp4081 {
																								__e.TailApply(PrimFunc(symshen_4parse_1failure))
																								return
																							} else {
																								tmp4067 := MakeNative(func(__e *ControlFlow) {
																									W2704 := __e.Get(1)
																									_ = W2704
																									tmp4068 := MakeNative(func(__e *ControlFlow) {
																										W2705 := __e.Get(1)
																										_ = W2705
																										tmp4069 := MakeNative(func(__e *ControlFlow) {
																											W2706 := __e.Get(1)
																											_ = W2706
																											tmp4076 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2706)

																											if True == tmp4076 {
																												__e.TailApply(PrimFunc(symshen_4parse_1failure))
																												return
																											} else {
																												tmp4070 := MakeNative(func(__e *ControlFlow) {
																													W2707 := __e.Get(1)
																													_ = W2707
																													tmp4071 := MakeNative(func(__e *ControlFlow) {
																														W2708 := __e.Get(1)
																														_ = W2708
																														tmp4072 := PrimCons(W2704, W2707)

																														__e.TailApply(PrimFunc(symshen_4comb), W2708, tmp4072)
																														return

																													}, 1)

																													tmp4073 := Call(__e, PrimFunc(symshen_4in_1_6), W2706)

																													__e.TailApply(tmp4071, tmp4073)
																													return

																												}, 1)

																												tmp4074 := Call(__e, PrimFunc(symshen_4_5_1out), W2706)

																												__e.TailApply(tmp4070, tmp4074)
																												return

																											}

																										}, 1)

																										tmp4077 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2705)

																										__e.TailApply(tmp4069, tmp4077)
																										return

																									}, 1)

																									tmp4078 := Call(__e, PrimFunc(symshen_4in_1_6), W2703)

																									__e.TailApply(tmp4068, tmp4078)
																									return

																								}, 1)

																								tmp4079 := Call(__e, PrimFunc(symshen_4_5_1out), W2703)

																								__e.TailApply(tmp4067, tmp4079)
																								return

																							}

																						}, 1)

																						tmp4082 := Call(__e, PrimFunc(symshen_4_5atom_6), V2629)

																						tmp4083 := Call(__e, tmp4066, tmp4082)

																						__e.TailApply(tmp4035, tmp4083)
																						return

																					} else {
																						__e.Return(W2696)
																						return
																					}

																				}, 1)

																				tmp4086 := MakeNative(func(__e *ControlFlow) {
																					W2697 := __e.Get(1)
																					_ = W2697
																					tmp4098 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2697)

																					if True == tmp4098 {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					} else {
																						tmp4087 := MakeNative(func(__e *ControlFlow) {
																							W2698 := __e.Get(1)
																							_ = W2698
																							tmp4088 := MakeNative(func(__e *ControlFlow) {
																								W2699 := __e.Get(1)
																								_ = W2699
																								tmp4094 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2699)

																								if True == tmp4094 {
																									__e.TailApply(PrimFunc(symshen_4parse_1failure))
																									return
																								} else {
																									tmp4089 := MakeNative(func(__e *ControlFlow) {
																										W2700 := __e.Get(1)
																										_ = W2700
																										tmp4090 := MakeNative(func(__e *ControlFlow) {
																											W2701 := __e.Get(1)
																											_ = W2701
																											__e.TailApply(PrimFunc(symshen_4comb), W2701, W2700)
																											return
																										}, 1)

																										tmp4091 := Call(__e, PrimFunc(symshen_4in_1_6), W2699)

																										__e.TailApply(tmp4090, tmp4091)
																										return

																									}, 1)

																									tmp4092 := Call(__e, PrimFunc(symshen_4_5_1out), W2699)

																									__e.TailApply(tmp4089, tmp4092)
																									return

																								}

																							}, 1)

																							tmp4095 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2698)

																							__e.TailApply(tmp4088, tmp4095)
																							return

																						}, 1)

																						tmp4096 := Call(__e, PrimFunc(symshen_4in_1_6), W2697)

																						__e.TailApply(tmp4087, tmp4096)
																						return

																					}

																				}, 1)

																				tmp4099 := Call(__e, PrimFunc(symshen_4_5comment_6), V2629)

																				tmp4100 := Call(__e, tmp4086, tmp4099)

																				__e.TailApply(tmp4034, tmp4100)
																				return

																			} else {
																				__e.Return(W2690)
																				return
																			}

																		}, 1)

																		tmp4103 := MakeNative(func(__e *ControlFlow) {
																			W2691 := __e.Get(1)
																			_ = W2691
																			tmp4117 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2691)

																			if True == tmp4117 {
																				__e.TailApply(PrimFunc(symshen_4parse_1failure))
																				return
																			} else {
																				tmp4104 := MakeNative(func(__e *ControlFlow) {
																					W2692 := __e.Get(1)
																					_ = W2692
																					tmp4105 := MakeNative(func(__e *ControlFlow) {
																						W2693 := __e.Get(1)
																						_ = W2693
																						tmp4113 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2693)

																						if True == tmp4113 {
																							__e.TailApply(PrimFunc(symshen_4parse_1failure))
																							return
																						} else {
																							tmp4106 := MakeNative(func(__e *ControlFlow) {
																								W2694 := __e.Get(1)
																								_ = W2694
																								tmp4107 := MakeNative(func(__e *ControlFlow) {
																									W2695 := __e.Get(1)
																									_ = W2695
																									tmp4108 := PrimIntern(MakeString(","))

																									tmp4109 := PrimCons(tmp4108, W2694)

																									__e.TailApply(PrimFunc(symshen_4comb), W2695, tmp4109)
																									return

																								}, 1)

																								tmp4110 := Call(__e, PrimFunc(symshen_4in_1_6), W2693)

																								__e.TailApply(tmp4107, tmp4110)
																								return

																							}, 1)

																							tmp4111 := Call(__e, PrimFunc(symshen_4_5_1out), W2693)

																							__e.TailApply(tmp4106, tmp4111)
																							return

																						}

																					}, 1)

																					tmp4114 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2692)

																					__e.TailApply(tmp4105, tmp4114)
																					return

																				}, 1)

																				tmp4115 := Call(__e, PrimFunc(symshen_4in_1_6), W2691)

																				__e.TailApply(tmp4104, tmp4115)
																				return

																			}

																		}, 1)

																		tmp4118 := Call(__e, PrimFunc(symshen_4_5comma_6), V2629)

																		tmp4119 := Call(__e, tmp4103, tmp4118)

																		__e.TailApply(tmp4033, tmp4119)
																		return

																	} else {
																		__e.Return(W2684)
																		return
																	}

																}, 1)

																tmp4122 := MakeNative(func(__e *ControlFlow) {
																	W2685 := __e.Get(1)
																	_ = W2685
																	tmp4136 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2685)

																	if True == tmp4136 {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	} else {
																		tmp4123 := MakeNative(func(__e *ControlFlow) {
																			W2686 := __e.Get(1)
																			_ = W2686
																			tmp4124 := MakeNative(func(__e *ControlFlow) {
																				W2687 := __e.Get(1)
																				_ = W2687
																				tmp4132 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2687)

																				if True == tmp4132 {
																					__e.TailApply(PrimFunc(symshen_4parse_1failure))
																					return
																				} else {
																					tmp4125 := MakeNative(func(__e *ControlFlow) {
																						W2688 := __e.Get(1)
																						_ = W2688
																						tmp4126 := MakeNative(func(__e *ControlFlow) {
																							W2689 := __e.Get(1)
																							_ = W2689
																							tmp4127 := PrimIntern(MakeString(":"))

																							tmp4128 := PrimCons(tmp4127, W2688)

																							__e.TailApply(PrimFunc(symshen_4comb), W2689, tmp4128)
																							return

																						}, 1)

																						tmp4129 := Call(__e, PrimFunc(symshen_4in_1_6), W2687)

																						__e.TailApply(tmp4126, tmp4129)
																						return

																					}, 1)

																					tmp4130 := Call(__e, PrimFunc(symshen_4_5_1out), W2687)

																					__e.TailApply(tmp4125, tmp4130)
																					return

																				}

																			}, 1)

																			tmp4133 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2686)

																			__e.TailApply(tmp4124, tmp4133)
																			return

																		}, 1)

																		tmp4134 := Call(__e, PrimFunc(symshen_4in_1_6), W2685)

																		__e.TailApply(tmp4123, tmp4134)
																		return

																	}

																}, 1)

																tmp4137 := Call(__e, PrimFunc(symshen_4_5colon_6), V2629)

																tmp4138 := Call(__e, tmp4122, tmp4137)

																__e.TailApply(tmp4032, tmp4138)
																return

															} else {
																__e.Return(W2676)
																return
															}

														}, 1)

														tmp4141 := MakeNative(func(__e *ControlFlow) {
															W2677 := __e.Get(1)
															_ = W2677
															tmp4161 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2677)

															if True == tmp4161 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp4142 := MakeNative(func(__e *ControlFlow) {
																	W2678 := __e.Get(1)
																	_ = W2678
																	tmp4143 := MakeNative(func(__e *ControlFlow) {
																		W2679 := __e.Get(1)
																		_ = W2679
																		tmp4157 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2679)

																		if True == tmp4157 {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		} else {
																			tmp4144 := MakeNative(func(__e *ControlFlow) {
																				W2680 := __e.Get(1)
																				_ = W2680
																				tmp4145 := MakeNative(func(__e *ControlFlow) {
																					W2681 := __e.Get(1)
																					_ = W2681
																					tmp4153 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2681)

																					if True == tmp4153 {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					} else {
																						tmp4146 := MakeNative(func(__e *ControlFlow) {
																							W2682 := __e.Get(1)
																							_ = W2682
																							tmp4147 := MakeNative(func(__e *ControlFlow) {
																								W2683 := __e.Get(1)
																								_ = W2683
																								tmp4148 := PrimIntern(MakeString(":="))

																								tmp4149 := PrimCons(tmp4148, W2682)

																								__e.TailApply(PrimFunc(symshen_4comb), W2683, tmp4149)
																								return

																							}, 1)

																							tmp4150 := Call(__e, PrimFunc(symshen_4in_1_6), W2681)

																							__e.TailApply(tmp4147, tmp4150)
																							return

																						}, 1)

																						tmp4151 := Call(__e, PrimFunc(symshen_4_5_1out), W2681)

																						__e.TailApply(tmp4146, tmp4151)
																						return

																					}

																				}, 1)

																				tmp4154 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2680)

																				__e.TailApply(tmp4145, tmp4154)
																				return

																			}, 1)

																			tmp4155 := Call(__e, PrimFunc(symshen_4in_1_6), W2679)

																			__e.TailApply(tmp4144, tmp4155)
																			return

																		}

																	}, 1)

																	tmp4158 := Call(__e, PrimFunc(symshen_4_5equal_6), W2678)

																	__e.TailApply(tmp4143, tmp4158)
																	return

																}, 1)

																tmp4159 := Call(__e, PrimFunc(symshen_4in_1_6), W2677)

																__e.TailApply(tmp4142, tmp4159)
																return

															}

														}, 1)

														tmp4162 := Call(__e, PrimFunc(symshen_4_5colon_6), V2629)

														tmp4163 := Call(__e, tmp4141, tmp4162)

														__e.TailApply(tmp4031, tmp4163)
														return

													} else {
														__e.Return(W2670)
														return
													}

												}, 1)

												tmp4166 := MakeNative(func(__e *ControlFlow) {
													W2671 := __e.Get(1)
													_ = W2671
													tmp4180 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2671)

													if True == tmp4180 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp4167 := MakeNative(func(__e *ControlFlow) {
															W2672 := __e.Get(1)
															_ = W2672
															tmp4168 := MakeNative(func(__e *ControlFlow) {
																W2673 := __e.Get(1)
																_ = W2673
																tmp4176 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2673)

																if True == tmp4176 {
																	__e.TailApply(PrimFunc(symshen_4parse_1failure))
																	return
																} else {
																	tmp4169 := MakeNative(func(__e *ControlFlow) {
																		W2674 := __e.Get(1)
																		_ = W2674
																		tmp4170 := MakeNative(func(__e *ControlFlow) {
																			W2675 := __e.Get(1)
																			_ = W2675
																			tmp4171 := PrimIntern(MakeString(";"))

																			tmp4172 := PrimCons(tmp4171, W2674)

																			__e.TailApply(PrimFunc(symshen_4comb), W2675, tmp4172)
																			return

																		}, 1)

																		tmp4173 := Call(__e, PrimFunc(symshen_4in_1_6), W2673)

																		__e.TailApply(tmp4170, tmp4173)
																		return

																	}, 1)

																	tmp4174 := Call(__e, PrimFunc(symshen_4_5_1out), W2673)

																	__e.TailApply(tmp4169, tmp4174)
																	return

																}

															}, 1)

															tmp4177 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2672)

															__e.TailApply(tmp4168, tmp4177)
															return

														}, 1)

														tmp4178 := Call(__e, PrimFunc(symshen_4in_1_6), W2671)

														__e.TailApply(tmp4167, tmp4178)
														return

													}

												}, 1)

												tmp4181 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2629)

												tmp4182 := Call(__e, tmp4166, tmp4181)

												__e.TailApply(tmp4030, tmp4182)
												return

											} else {
												__e.Return(W2664)
												return
											}

										}, 1)

										tmp4185 := MakeNative(func(__e *ControlFlow) {
											W2665 := __e.Get(1)
											_ = W2665
											tmp4198 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2665)

											if True == tmp4198 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp4186 := MakeNative(func(__e *ControlFlow) {
													W2666 := __e.Get(1)
													_ = W2666
													tmp4187 := MakeNative(func(__e *ControlFlow) {
														W2667 := __e.Get(1)
														_ = W2667
														tmp4194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2667)

														if True == tmp4194 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp4188 := MakeNative(func(__e *ControlFlow) {
																W2668 := __e.Get(1)
																_ = W2668
																tmp4189 := MakeNative(func(__e *ControlFlow) {
																	W2669 := __e.Get(1)
																	_ = W2669
																	tmp4190 := PrimCons(symbar_b, W2668)

																	__e.TailApply(PrimFunc(symshen_4comb), W2669, tmp4190)
																	return

																}, 1)

																tmp4191 := Call(__e, PrimFunc(symshen_4in_1_6), W2667)

																__e.TailApply(tmp4189, tmp4191)
																return

															}, 1)

															tmp4192 := Call(__e, PrimFunc(symshen_4_5_1out), W2667)

															__e.TailApply(tmp4188, tmp4192)
															return

														}

													}, 1)

													tmp4195 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2666)

													__e.TailApply(tmp4187, tmp4195)
													return

												}, 1)

												tmp4196 := Call(__e, PrimFunc(symshen_4in_1_6), W2665)

												__e.TailApply(tmp4186, tmp4196)
												return

											}

										}, 1)

										tmp4199 := Call(__e, PrimFunc(symshen_4_5bar_6), V2629)

										tmp4200 := Call(__e, tmp4185, tmp4199)

										__e.TailApply(tmp4029, tmp4200)
										return

									} else {
										__e.Return(W2658)
										return
									}

								}, 1)

								tmp4203 := MakeNative(func(__e *ControlFlow) {
									W2659 := __e.Get(1)
									_ = W2659
									tmp4216 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2659)

									if True == tmp4216 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4204 := MakeNative(func(__e *ControlFlow) {
											W2660 := __e.Get(1)
											_ = W2660
											tmp4205 := MakeNative(func(__e *ControlFlow) {
												W2661 := __e.Get(1)
												_ = W2661
												tmp4212 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2661)

												if True == tmp4212 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp4206 := MakeNative(func(__e *ControlFlow) {
														W2662 := __e.Get(1)
														_ = W2662
														tmp4207 := MakeNative(func(__e *ControlFlow) {
															W2663 := __e.Get(1)
															_ = W2663
															tmp4208 := PrimCons(sym_j, W2662)

															__e.TailApply(PrimFunc(symshen_4comb), W2663, tmp4208)
															return

														}, 1)

														tmp4209 := Call(__e, PrimFunc(symshen_4in_1_6), W2661)

														__e.TailApply(tmp4207, tmp4209)
														return

													}, 1)

													tmp4210 := Call(__e, PrimFunc(symshen_4_5_1out), W2661)

													__e.TailApply(tmp4206, tmp4210)
													return

												}

											}, 1)

											tmp4213 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2660)

											__e.TailApply(tmp4205, tmp4213)
											return

										}, 1)

										tmp4214 := Call(__e, PrimFunc(symshen_4in_1_6), W2659)

										__e.TailApply(tmp4204, tmp4214)
										return

									}

								}, 1)

								tmp4217 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2629)

								tmp4218 := Call(__e, tmp4203, tmp4217)

								__e.TailApply(tmp4028, tmp4218)
								return

							} else {
								__e.Return(W2652)
								return
							}

						}, 1)

						tmp4221 := MakeNative(func(__e *ControlFlow) {
							W2653 := __e.Get(1)
							_ = W2653
							tmp4234 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2653)

							if True == tmp4234 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4222 := MakeNative(func(__e *ControlFlow) {
									W2654 := __e.Get(1)
									_ = W2654
									tmp4223 := MakeNative(func(__e *ControlFlow) {
										W2655 := __e.Get(1)
										_ = W2655
										tmp4230 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2655)

										if True == tmp4230 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp4224 := MakeNative(func(__e *ControlFlow) {
												W2656 := __e.Get(1)
												_ = W2656
												tmp4225 := MakeNative(func(__e *ControlFlow) {
													W2657 := __e.Get(1)
													_ = W2657
													tmp4226 := PrimCons(sym_i, W2656)

													__e.TailApply(PrimFunc(symshen_4comb), W2657, tmp4226)
													return

												}, 1)

												tmp4227 := Call(__e, PrimFunc(symshen_4in_1_6), W2655)

												__e.TailApply(tmp4225, tmp4227)
												return

											}, 1)

											tmp4228 := Call(__e, PrimFunc(symshen_4_5_1out), W2655)

											__e.TailApply(tmp4224, tmp4228)
											return

										}

									}, 1)

									tmp4231 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2654)

									__e.TailApply(tmp4223, tmp4231)
									return

								}, 1)

								tmp4232 := Call(__e, PrimFunc(symshen_4in_1_6), W2653)

								__e.TailApply(tmp4222, tmp4232)
								return

							}

						}, 1)

						tmp4235 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2629)

						tmp4236 := Call(__e, tmp4221, tmp4235)

						__e.TailApply(tmp4027, tmp4236)
						return

					} else {
						__e.Return(W2641)
						return
					}

				}, 1)

				tmp4239 := MakeNative(func(__e *ControlFlow) {
					W2642 := __e.Get(1)
					_ = W2642
					tmp4266 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2642)

					if True == tmp4266 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4240 := MakeNative(func(__e *ControlFlow) {
							W2643 := __e.Get(1)
							_ = W2643
							tmp4241 := MakeNative(func(__e *ControlFlow) {
								W2644 := __e.Get(1)
								_ = W2644
								tmp4262 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2644)

								if True == tmp4262 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp4242 := MakeNative(func(__e *ControlFlow) {
										W2645 := __e.Get(1)
										_ = W2645
										tmp4243 := MakeNative(func(__e *ControlFlow) {
											W2646 := __e.Get(1)
											_ = W2646
											tmp4244 := MakeNative(func(__e *ControlFlow) {
												W2647 := __e.Get(1)
												_ = W2647
												tmp4257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2647)

												if True == tmp4257 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp4245 := MakeNative(func(__e *ControlFlow) {
														W2648 := __e.Get(1)
														_ = W2648
														tmp4246 := MakeNative(func(__e *ControlFlow) {
															W2649 := __e.Get(1)
															_ = W2649
															tmp4253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2649)

															if True == tmp4253 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp4247 := MakeNative(func(__e *ControlFlow) {
																	W2650 := __e.Get(1)
																	_ = W2650
																	tmp4248 := MakeNative(func(__e *ControlFlow) {
																		W2651 := __e.Get(1)
																		_ = W2651
																		tmp4249 := Call(__e, PrimFunc(symshen_4add_1sexpr), W2645, W2650)

																		__e.TailApply(PrimFunc(symshen_4comb), W2651, tmp4249)
																		return

																	}, 1)

																	tmp4250 := Call(__e, PrimFunc(symshen_4in_1_6), W2649)

																	__e.TailApply(tmp4248, tmp4250)
																	return

																}, 1)

																tmp4251 := Call(__e, PrimFunc(symshen_4_5_1out), W2649)

																__e.TailApply(tmp4247, tmp4251)
																return

															}

														}, 1)

														tmp4254 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2648)

														__e.TailApply(tmp4246, tmp4254)
														return

													}, 1)

													tmp4255 := Call(__e, PrimFunc(symshen_4in_1_6), W2647)

													__e.TailApply(tmp4245, tmp4255)
													return

												}

											}, 1)

											tmp4258 := Call(__e, PrimFunc(symshen_4_5rrb_6), W2646)

											__e.TailApply(tmp4244, tmp4258)
											return

										}, 1)

										tmp4259 := Call(__e, PrimFunc(symshen_4in_1_6), W2644)

										__e.TailApply(tmp4243, tmp4259)
										return

									}, 1)

									tmp4260 := Call(__e, PrimFunc(symshen_4_5_1out), W2644)

									__e.TailApply(tmp4242, tmp4260)
									return

								}

							}, 1)

							tmp4263 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2643)

							__e.TailApply(tmp4241, tmp4263)
							return

						}, 1)

						tmp4264 := Call(__e, PrimFunc(symshen_4in_1_6), W2642)

						__e.TailApply(tmp4240, tmp4264)
						return

					}

				}, 1)

				tmp4267 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2629)

				tmp4268 := Call(__e, tmp4239, tmp4267)

				__e.TailApply(tmp4026, tmp4268)
				return

			} else {
				__e.Return(W2630)
				return
			}

		}, 1)

		tmp4271 := MakeNative(func(__e *ControlFlow) {
			W2631 := __e.Get(1)
			_ = W2631
			tmp4299 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2631)

			if True == tmp4299 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4272 := MakeNative(func(__e *ControlFlow) {
					W2632 := __e.Get(1)
					_ = W2632
					tmp4273 := MakeNative(func(__e *ControlFlow) {
						W2633 := __e.Get(1)
						_ = W2633
						tmp4295 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2633)

						if True == tmp4295 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4274 := MakeNative(func(__e *ControlFlow) {
								W2634 := __e.Get(1)
								_ = W2634
								tmp4275 := MakeNative(func(__e *ControlFlow) {
									W2635 := __e.Get(1)
									_ = W2635
									tmp4276 := MakeNative(func(__e *ControlFlow) {
										W2636 := __e.Get(1)
										_ = W2636
										tmp4290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2636)

										if True == tmp4290 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp4277 := MakeNative(func(__e *ControlFlow) {
												W2637 := __e.Get(1)
												_ = W2637
												tmp4278 := MakeNative(func(__e *ControlFlow) {
													W2638 := __e.Get(1)
													_ = W2638
													tmp4286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2638)

													if True == tmp4286 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp4279 := MakeNative(func(__e *ControlFlow) {
															W2639 := __e.Get(1)
															_ = W2639
															tmp4280 := MakeNative(func(__e *ControlFlow) {
																W2640 := __e.Get(1)
																_ = W2640
																tmp4281 := Call(__e, PrimFunc(symshen_4cons_1form), W2634)

																tmp4282 := PrimCons(tmp4281, W2639)

																__e.TailApply(PrimFunc(symshen_4comb), W2640, tmp4282)
																return

															}, 1)

															tmp4283 := Call(__e, PrimFunc(symshen_4in_1_6), W2638)

															__e.TailApply(tmp4280, tmp4283)
															return

														}, 1)

														tmp4284 := Call(__e, PrimFunc(symshen_4_5_1out), W2638)

														__e.TailApply(tmp4279, tmp4284)
														return

													}

												}, 1)

												tmp4287 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2637)

												__e.TailApply(tmp4278, tmp4287)
												return

											}, 1)

											tmp4288 := Call(__e, PrimFunc(symshen_4in_1_6), W2636)

											__e.TailApply(tmp4277, tmp4288)
											return

										}

									}, 1)

									tmp4291 := Call(__e, PrimFunc(symshen_4_5rsb_6), W2635)

									__e.TailApply(tmp4276, tmp4291)
									return

								}, 1)

								tmp4292 := Call(__e, PrimFunc(symshen_4in_1_6), W2633)

								__e.TailApply(tmp4275, tmp4292)
								return

							}, 1)

							tmp4293 := Call(__e, PrimFunc(symshen_4_5_1out), W2633)

							__e.TailApply(tmp4274, tmp4293)
							return

						}

					}, 1)

					tmp4296 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2632)

					__e.TailApply(tmp4273, tmp4296)
					return

				}, 1)

				tmp4297 := Call(__e, PrimFunc(symshen_4in_1_6), W2631)

				__e.TailApply(tmp4272, tmp4297)
				return

			}

		}, 1)

		tmp4300 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2629)

		tmp4301 := Call(__e, tmp4271, tmp4300)

		__e.TailApply(tmp4025, tmp4301)
		return

	}, 1)

	tmp4302 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp4024)

	_ = tmp4302

	tmp4303 := MakeNative(func(__e *ControlFlow) {
		V2718 := __e.Get(1)
		_ = V2718
		V2719 := __e.Get(2)
		_ = V2719
		tmp4321 := PrimIsPair(V2718)

		var ifres4308 Obj

		if True == tmp4321 {
			tmp4319 := PrimHead(V2718)

			tmp4320 := PrimEqual(sym_3, tmp4319)

			var ifres4310 Obj

			if True == tmp4320 {
				tmp4317 := PrimTail(V2718)

				tmp4318 := PrimIsPair(tmp4317)

				var ifres4312 Obj

				if True == tmp4318 {
					tmp4314 := PrimTail(V2718)

					tmp4315 := PrimTail(tmp4314)

					tmp4316 := PrimEqual(Nil, tmp4315)

					var ifres4313 Obj

					if True == tmp4316 {
						ifres4313 = True

					} else {
						ifres4313 = False

					}

					ifres4312 = ifres4313

				} else {
					ifres4312 = False

				}

				var ifres4311 Obj

				if True == ifres4312 {
					ifres4311 = True

				} else {
					ifres4311 = False

				}

				ifres4310 = ifres4311

			} else {
				ifres4310 = False

			}

			var ifres4309 Obj

			if True == ifres4310 {
				ifres4309 = True

			} else {
				ifres4309 = False

			}

			ifres4308 = ifres4309

		} else {
			ifres4308 = False

		}

		if True == ifres4308 {
			tmp4304 := PrimTail(V2718)

			tmp4305 := PrimHead(tmp4304)

			tmp4306 := Call(__e, PrimFunc(symexplode), tmp4305)

			__e.TailApply(PrimFunc(symappend), tmp4306, V2719)
			return

		} else {
			__e.Return(PrimCons(V2718, V2719))
			return
		}

	}, 2)

	tmp4322 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp4303)

	_ = tmp4322

	tmp4323 := MakeNative(func(__e *ControlFlow) {
		V2720 := __e.Get(1)
		_ = V2720
		tmp4324 := MakeNative(func(__e *ControlFlow) {
			W2721 := __e.Get(1)
			_ = W2721
			tmp4326 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2721)

			if True == tmp4326 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2721)
				return
			}

		}, 1)

		tmp4332 := Call(__e, PrimFunc(symshen_4hds_a_2), V2720, MakeNumber(91))

		var ifres4327 Obj

		if True == tmp4332 {
			tmp4328 := MakeNative(func(__e *ControlFlow) {
				W2722 := __e.Get(1)
				_ = W2722
				__e.TailApply(PrimFunc(symshen_4comb), W2722, symshen_4skip)
				return
			}, 1)

			tmp4329 := Call(__e, PrimFunc(symtail), V2720)

			tmp4330 := Call(__e, tmp4328, tmp4329)

			ifres4327 = tmp4330

		} else {
			tmp4331 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4327 = tmp4331

		}

		__e.TailApply(tmp4324, ifres4327)
		return

	}, 1)

	tmp4333 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp4323)

	_ = tmp4333

	tmp4334 := MakeNative(func(__e *ControlFlow) {
		V2723 := __e.Get(1)
		_ = V2723
		tmp4335 := MakeNative(func(__e *ControlFlow) {
			W2724 := __e.Get(1)
			_ = W2724
			tmp4337 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2724)

			if True == tmp4337 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2724)
				return
			}

		}, 1)

		tmp4343 := Call(__e, PrimFunc(symshen_4hds_a_2), V2723, MakeNumber(93))

		var ifres4338 Obj

		if True == tmp4343 {
			tmp4339 := MakeNative(func(__e *ControlFlow) {
				W2725 := __e.Get(1)
				_ = W2725
				__e.TailApply(PrimFunc(symshen_4comb), W2725, symshen_4skip)
				return
			}, 1)

			tmp4340 := Call(__e, PrimFunc(symtail), V2723)

			tmp4341 := Call(__e, tmp4339, tmp4340)

			ifres4338 = tmp4341

		} else {
			tmp4342 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4338 = tmp4342

		}

		__e.TailApply(tmp4335, ifres4338)
		return

	}, 1)

	tmp4344 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp4334)

	_ = tmp4344

	tmp4345 := MakeNative(func(__e *ControlFlow) {
		V2726 := __e.Get(1)
		_ = V2726
		tmp4346 := MakeNative(func(__e *ControlFlow) {
			W2727 := __e.Get(1)
			_ = W2727
			tmp4348 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2727)

			if True == tmp4348 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2727)
				return
			}

		}, 1)

		tmp4349 := MakeNative(func(__e *ControlFlow) {
			W2728 := __e.Get(1)
			_ = W2728
			tmp4355 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2728)

			if True == tmp4355 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4350 := MakeNative(func(__e *ControlFlow) {
					W2729 := __e.Get(1)
					_ = W2729
					tmp4351 := MakeNative(func(__e *ControlFlow) {
						W2730 := __e.Get(1)
						_ = W2730
						__e.TailApply(PrimFunc(symshen_4comb), W2730, W2729)
						return
					}, 1)

					tmp4352 := Call(__e, PrimFunc(symshen_4in_1_6), W2728)

					__e.TailApply(tmp4351, tmp4352)
					return

				}, 1)

				tmp4353 := Call(__e, PrimFunc(symshen_4_5_1out), W2728)

				__e.TailApply(tmp4350, tmp4353)
				return

			}

		}, 1)

		tmp4356 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2726)

		tmp4357 := Call(__e, tmp4349, tmp4356)

		__e.TailApply(tmp4346, tmp4357)
		return

	}, 1)

	tmp4358 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp4345)

	_ = tmp4358

	tmp4359 := MakeNative(func(__e *ControlFlow) {
		V2731 := __e.Get(1)
		_ = V2731
		tmp4360 := MakeNative(func(__e *ControlFlow) {
			W2732 := __e.Get(1)
			_ = W2732
			tmp4362 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2732)

			if True == tmp4362 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2732)
				return
			}

		}, 1)

		tmp4363 := MakeNative(func(__e *ControlFlow) {
			W2733 := __e.Get(1)
			_ = W2733
			tmp4369 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2733)

			if True == tmp4369 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4364 := MakeNative(func(__e *ControlFlow) {
					W2734 := __e.Get(1)
					_ = W2734
					tmp4365 := MakeNative(func(__e *ControlFlow) {
						W2735 := __e.Get(1)
						_ = W2735
						__e.TailApply(PrimFunc(symshen_4comb), W2735, W2734)
						return
					}, 1)

					tmp4366 := Call(__e, PrimFunc(symshen_4in_1_6), W2733)

					__e.TailApply(tmp4365, tmp4366)
					return

				}, 1)

				tmp4367 := Call(__e, PrimFunc(symshen_4_5_1out), W2733)

				__e.TailApply(tmp4364, tmp4367)
				return

			}

		}, 1)

		tmp4370 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2731)

		tmp4371 := Call(__e, tmp4363, tmp4370)

		__e.TailApply(tmp4360, tmp4371)
		return

	}, 1)

	tmp4372 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp4359)

	_ = tmp4372

	tmp4373 := MakeNative(func(__e *ControlFlow) {
		V2737 := __e.Get(1)
		_ = V2737
		tmp4430 := PrimEqual(Nil, V2737)

		if True == tmp4430 {
			__e.Return(Nil)
			return
		} else {
			tmp4428 := PrimIsPair(V2737)

			var ifres4408 Obj

			if True == tmp4428 {
				tmp4426 := PrimTail(V2737)

				tmp4427 := PrimIsPair(tmp4426)

				var ifres4410 Obj

				if True == tmp4427 {
					tmp4423 := PrimTail(V2737)

					tmp4424 := PrimTail(tmp4423)

					tmp4425 := PrimIsPair(tmp4424)

					var ifres4412 Obj

					if True == tmp4425 {
						tmp4419 := PrimTail(V2737)

						tmp4420 := PrimTail(tmp4419)

						tmp4421 := PrimTail(tmp4420)

						tmp4422 := PrimEqual(Nil, tmp4421)

						var ifres4414 Obj

						if True == tmp4422 {
							tmp4416 := PrimTail(V2737)

							tmp4417 := PrimHead(tmp4416)

							tmp4418 := PrimEqual(tmp4417, symbar_b)

							var ifres4415 Obj

							if True == tmp4418 {
								ifres4415 = True

							} else {
								ifres4415 = False

							}

							ifres4414 = ifres4415

						} else {
							ifres4414 = False

						}

						var ifres4413 Obj

						if True == ifres4414 {
							ifres4413 = True

						} else {
							ifres4413 = False

						}

						ifres4412 = ifres4413

					} else {
						ifres4412 = False

					}

					var ifres4411 Obj

					if True == ifres4412 {
						ifres4411 = True

					} else {
						ifres4411 = False

					}

					ifres4410 = ifres4411

				} else {
					ifres4410 = False

				}

				var ifres4409 Obj

				if True == ifres4410 {
					ifres4409 = True

				} else {
					ifres4409 = False

				}

				ifres4408 = ifres4409

			} else {
				ifres4408 = False

			}

			if True == ifres4408 {
				tmp4374 := PrimHead(V2737)

				tmp4375 := PrimTail(V2737)

				tmp4376 := PrimTail(tmp4375)

				tmp4377 := PrimCons(tmp4374, tmp4376)

				__e.Return(PrimCons(symcons, tmp4377))
				return

			} else {
				tmp4406 := PrimIsPair(V2737)

				var ifres4386 Obj

				if True == tmp4406 {
					tmp4404 := PrimTail(V2737)

					tmp4405 := PrimIsPair(tmp4404)

					var ifres4388 Obj

					if True == tmp4405 {
						tmp4401 := PrimTail(V2737)

						tmp4402 := PrimTail(tmp4401)

						tmp4403 := PrimIsPair(tmp4402)

						var ifres4390 Obj

						if True == tmp4403 {
							tmp4397 := PrimTail(V2737)

							tmp4398 := PrimTail(tmp4397)

							tmp4399 := PrimTail(tmp4398)

							tmp4400 := PrimIsPair(tmp4399)

							var ifres4392 Obj

							if True == tmp4400 {
								tmp4394 := PrimTail(V2737)

								tmp4395 := PrimHead(tmp4394)

								tmp4396 := PrimEqual(tmp4395, symbar_b)

								var ifres4393 Obj

								if True == tmp4396 {
									ifres4393 = True

								} else {
									ifres4393 = False

								}

								ifres4392 = ifres4393

							} else {
								ifres4392 = False

							}

							var ifres4391 Obj

							if True == ifres4392 {
								ifres4391 = True

							} else {
								ifres4391 = False

							}

							ifres4390 = ifres4391

						} else {
							ifres4390 = False

						}

						var ifres4389 Obj

						if True == ifres4390 {
							ifres4389 = True

						} else {
							ifres4389 = False

						}

						ifres4388 = ifres4389

					} else {
						ifres4388 = False

					}

					var ifres4387 Obj

					if True == ifres4388 {
						ifres4387 = True

					} else {
						ifres4387 = False

					}

					ifres4386 = ifres4387

				} else {
					ifres4386 = False

				}

				if True == ifres4386 {
					__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
					return
				} else {
					tmp4384 := PrimIsPair(V2737)

					if True == tmp4384 {
						tmp4378 := PrimHead(V2737)

						tmp4379 := PrimTail(V2737)

						tmp4380 := Call(__e, PrimFunc(symshen_4cons_1form), tmp4379)

						tmp4381 := PrimCons(tmp4380, Nil)

						tmp4382 := PrimCons(tmp4378, tmp4381)

						__e.Return(PrimCons(symcons, tmp4382))
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4cons_1form)
						return
					}

				}

			}

		}

	}, 1)

	tmp4431 := Call(__e, ns2_1set, symshen_4cons_1form, tmp4373)

	_ = tmp4431

	tmp4432 := MakeNative(func(__e *ControlFlow) {
		V2738 := __e.Get(1)
		_ = V2738
		tmp4433 := MakeNative(func(__e *ControlFlow) {
			W2739 := __e.Get(1)
			_ = W2739
			tmp4435 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2739)

			if True == tmp4435 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2739)
				return
			}

		}, 1)

		tmp4441 := Call(__e, PrimFunc(symshen_4hds_a_2), V2738, MakeNumber(40))

		var ifres4436 Obj

		if True == tmp4441 {
			tmp4437 := MakeNative(func(__e *ControlFlow) {
				W2740 := __e.Get(1)
				_ = W2740
				__e.TailApply(PrimFunc(symshen_4comb), W2740, symshen_4skip)
				return
			}, 1)

			tmp4438 := Call(__e, PrimFunc(symtail), V2738)

			tmp4439 := Call(__e, tmp4437, tmp4438)

			ifres4436 = tmp4439

		} else {
			tmp4440 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4436 = tmp4440

		}

		__e.TailApply(tmp4433, ifres4436)
		return

	}, 1)

	tmp4442 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp4432)

	_ = tmp4442

	tmp4443 := MakeNative(func(__e *ControlFlow) {
		V2741 := __e.Get(1)
		_ = V2741
		tmp4444 := MakeNative(func(__e *ControlFlow) {
			W2742 := __e.Get(1)
			_ = W2742
			tmp4446 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2742)

			if True == tmp4446 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2742)
				return
			}

		}, 1)

		tmp4452 := Call(__e, PrimFunc(symshen_4hds_a_2), V2741, MakeNumber(41))

		var ifres4447 Obj

		if True == tmp4452 {
			tmp4448 := MakeNative(func(__e *ControlFlow) {
				W2743 := __e.Get(1)
				_ = W2743
				__e.TailApply(PrimFunc(symshen_4comb), W2743, symshen_4skip)
				return
			}, 1)

			tmp4449 := Call(__e, PrimFunc(symtail), V2741)

			tmp4450 := Call(__e, tmp4448, tmp4449)

			ifres4447 = tmp4450

		} else {
			tmp4451 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4447 = tmp4451

		}

		__e.TailApply(tmp4444, ifres4447)
		return

	}, 1)

	tmp4453 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp4443)

	_ = tmp4453

	tmp4454 := MakeNative(func(__e *ControlFlow) {
		V2744 := __e.Get(1)
		_ = V2744
		tmp4455 := MakeNative(func(__e *ControlFlow) {
			W2745 := __e.Get(1)
			_ = W2745
			tmp4457 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2745)

			if True == tmp4457 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2745)
				return
			}

		}, 1)

		tmp4463 := Call(__e, PrimFunc(symshen_4hds_a_2), V2744, MakeNumber(123))

		var ifres4458 Obj

		if True == tmp4463 {
			tmp4459 := MakeNative(func(__e *ControlFlow) {
				W2746 := __e.Get(1)
				_ = W2746
				__e.TailApply(PrimFunc(symshen_4comb), W2746, symshen_4skip)
				return
			}, 1)

			tmp4460 := Call(__e, PrimFunc(symtail), V2744)

			tmp4461 := Call(__e, tmp4459, tmp4460)

			ifres4458 = tmp4461

		} else {
			tmp4462 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4458 = tmp4462

		}

		__e.TailApply(tmp4455, ifres4458)
		return

	}, 1)

	tmp4464 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp4454)

	_ = tmp4464

	tmp4465 := MakeNative(func(__e *ControlFlow) {
		V2747 := __e.Get(1)
		_ = V2747
		tmp4466 := MakeNative(func(__e *ControlFlow) {
			W2748 := __e.Get(1)
			_ = W2748
			tmp4468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2748)

			if True == tmp4468 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2748)
				return
			}

		}, 1)

		tmp4474 := Call(__e, PrimFunc(symshen_4hds_a_2), V2747, MakeNumber(125))

		var ifres4469 Obj

		if True == tmp4474 {
			tmp4470 := MakeNative(func(__e *ControlFlow) {
				W2749 := __e.Get(1)
				_ = W2749
				__e.TailApply(PrimFunc(symshen_4comb), W2749, symshen_4skip)
				return
			}, 1)

			tmp4471 := Call(__e, PrimFunc(symtail), V2747)

			tmp4472 := Call(__e, tmp4470, tmp4471)

			ifres4469 = tmp4472

		} else {
			tmp4473 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4469 = tmp4473

		}

		__e.TailApply(tmp4466, ifres4469)
		return

	}, 1)

	tmp4475 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp4465)

	_ = tmp4475

	tmp4476 := MakeNative(func(__e *ControlFlow) {
		V2750 := __e.Get(1)
		_ = V2750
		tmp4477 := MakeNative(func(__e *ControlFlow) {
			W2751 := __e.Get(1)
			_ = W2751
			tmp4479 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2751)

			if True == tmp4479 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2751)
				return
			}

		}, 1)

		tmp4485 := Call(__e, PrimFunc(symshen_4hds_a_2), V2750, MakeNumber(124))

		var ifres4480 Obj

		if True == tmp4485 {
			tmp4481 := MakeNative(func(__e *ControlFlow) {
				W2752 := __e.Get(1)
				_ = W2752
				__e.TailApply(PrimFunc(symshen_4comb), W2752, symshen_4skip)
				return
			}, 1)

			tmp4482 := Call(__e, PrimFunc(symtail), V2750)

			tmp4483 := Call(__e, tmp4481, tmp4482)

			ifres4480 = tmp4483

		} else {
			tmp4484 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4480 = tmp4484

		}

		__e.TailApply(tmp4477, ifres4480)
		return

	}, 1)

	tmp4486 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp4476)

	_ = tmp4486

	tmp4487 := MakeNative(func(__e *ControlFlow) {
		V2753 := __e.Get(1)
		_ = V2753
		tmp4488 := MakeNative(func(__e *ControlFlow) {
			W2754 := __e.Get(1)
			_ = W2754
			tmp4490 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2754)

			if True == tmp4490 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2754)
				return
			}

		}, 1)

		tmp4496 := Call(__e, PrimFunc(symshen_4hds_a_2), V2753, MakeNumber(59))

		var ifres4491 Obj

		if True == tmp4496 {
			tmp4492 := MakeNative(func(__e *ControlFlow) {
				W2755 := __e.Get(1)
				_ = W2755
				__e.TailApply(PrimFunc(symshen_4comb), W2755, symshen_4skip)
				return
			}, 1)

			tmp4493 := Call(__e, PrimFunc(symtail), V2753)

			tmp4494 := Call(__e, tmp4492, tmp4493)

			ifres4491 = tmp4494

		} else {
			tmp4495 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4491 = tmp4495

		}

		__e.TailApply(tmp4488, ifres4491)
		return

	}, 1)

	tmp4497 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp4487)

	_ = tmp4497

	tmp4498 := MakeNative(func(__e *ControlFlow) {
		V2756 := __e.Get(1)
		_ = V2756
		tmp4499 := MakeNative(func(__e *ControlFlow) {
			W2757 := __e.Get(1)
			_ = W2757
			tmp4501 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2757)

			if True == tmp4501 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2757)
				return
			}

		}, 1)

		tmp4507 := Call(__e, PrimFunc(symshen_4hds_a_2), V2756, MakeNumber(58))

		var ifres4502 Obj

		if True == tmp4507 {
			tmp4503 := MakeNative(func(__e *ControlFlow) {
				W2758 := __e.Get(1)
				_ = W2758
				__e.TailApply(PrimFunc(symshen_4comb), W2758, symshen_4skip)
				return
			}, 1)

			tmp4504 := Call(__e, PrimFunc(symtail), V2756)

			tmp4505 := Call(__e, tmp4503, tmp4504)

			ifres4502 = tmp4505

		} else {
			tmp4506 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4502 = tmp4506

		}

		__e.TailApply(tmp4499, ifres4502)
		return

	}, 1)

	tmp4508 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp4498)

	_ = tmp4508

	tmp4509 := MakeNative(func(__e *ControlFlow) {
		V2759 := __e.Get(1)
		_ = V2759
		tmp4510 := MakeNative(func(__e *ControlFlow) {
			W2760 := __e.Get(1)
			_ = W2760
			tmp4512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2760)

			if True == tmp4512 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2760)
				return
			}

		}, 1)

		tmp4518 := Call(__e, PrimFunc(symshen_4hds_a_2), V2759, MakeNumber(44))

		var ifres4513 Obj

		if True == tmp4518 {
			tmp4514 := MakeNative(func(__e *ControlFlow) {
				W2761 := __e.Get(1)
				_ = W2761
				__e.TailApply(PrimFunc(symshen_4comb), W2761, symshen_4skip)
				return
			}, 1)

			tmp4515 := Call(__e, PrimFunc(symtail), V2759)

			tmp4516 := Call(__e, tmp4514, tmp4515)

			ifres4513 = tmp4516

		} else {
			tmp4517 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4513 = tmp4517

		}

		__e.TailApply(tmp4510, ifres4513)
		return

	}, 1)

	tmp4519 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp4509)

	_ = tmp4519

	tmp4520 := MakeNative(func(__e *ControlFlow) {
		V2762 := __e.Get(1)
		_ = V2762
		tmp4521 := MakeNative(func(__e *ControlFlow) {
			W2763 := __e.Get(1)
			_ = W2763
			tmp4523 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2763)

			if True == tmp4523 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2763)
				return
			}

		}, 1)

		tmp4529 := Call(__e, PrimFunc(symshen_4hds_a_2), V2762, MakeNumber(61))

		var ifres4524 Obj

		if True == tmp4529 {
			tmp4525 := MakeNative(func(__e *ControlFlow) {
				W2764 := __e.Get(1)
				_ = W2764
				__e.TailApply(PrimFunc(symshen_4comb), W2764, symshen_4skip)
				return
			}, 1)

			tmp4526 := Call(__e, PrimFunc(symtail), V2762)

			tmp4527 := Call(__e, tmp4525, tmp4526)

			ifres4524 = tmp4527

		} else {
			tmp4528 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4524 = tmp4528

		}

		__e.TailApply(tmp4521, ifres4524)
		return

	}, 1)

	tmp4530 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp4520)

	_ = tmp4530

	tmp4531 := MakeNative(func(__e *ControlFlow) {
		V2765 := __e.Get(1)
		_ = V2765
		tmp4532 := MakeNative(func(__e *ControlFlow) {
			W2766 := __e.Get(1)
			_ = W2766
			tmp4544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2766)

			if True == tmp4544 {
				tmp4533 := MakeNative(func(__e *ControlFlow) {
					W2769 := __e.Get(1)
					_ = W2769
					tmp4535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2769)

					if True == tmp4535 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2769)
						return
					}

				}, 1)

				tmp4536 := MakeNative(func(__e *ControlFlow) {
					W2770 := __e.Get(1)
					_ = W2770
					tmp4540 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2770)

					if True == tmp4540 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4537 := MakeNative(func(__e *ControlFlow) {
							W2771 := __e.Get(1)
							_ = W2771
							__e.TailApply(PrimFunc(symshen_4comb), W2771, symshen_4skip)
							return
						}, 1)

						tmp4538 := Call(__e, PrimFunc(symshen_4in_1_6), W2770)

						__e.TailApply(tmp4537, tmp4538)
						return

					}

				}, 1)

				tmp4541 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2765)

				tmp4542 := Call(__e, tmp4536, tmp4541)

				__e.TailApply(tmp4533, tmp4542)
				return

			} else {
				__e.Return(W2766)
				return
			}

		}, 1)

		tmp4545 := MakeNative(func(__e *ControlFlow) {
			W2767 := __e.Get(1)
			_ = W2767
			tmp4549 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2767)

			if True == tmp4549 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4546 := MakeNative(func(__e *ControlFlow) {
					W2768 := __e.Get(1)
					_ = W2768
					__e.TailApply(PrimFunc(symshen_4comb), W2768, symshen_4skip)
					return
				}, 1)

				tmp4547 := Call(__e, PrimFunc(symshen_4in_1_6), W2767)

				__e.TailApply(tmp4546, tmp4547)
				return

			}

		}, 1)

		tmp4550 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2765)

		tmp4551 := Call(__e, tmp4545, tmp4550)

		__e.TailApply(tmp4532, tmp4551)
		return

	}, 1)

	tmp4552 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp4531)

	_ = tmp4552

	tmp4553 := MakeNative(func(__e *ControlFlow) {
		V2772 := __e.Get(1)
		_ = V2772
		tmp4554 := MakeNative(func(__e *ControlFlow) {
			W2773 := __e.Get(1)
			_ = W2773
			tmp4556 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2773)

			if True == tmp4556 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2773)
				return
			}

		}, 1)

		tmp4557 := MakeNative(func(__e *ControlFlow) {
			W2774 := __e.Get(1)
			_ = W2774
			tmp4579 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2774)

			if True == tmp4579 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4558 := MakeNative(func(__e *ControlFlow) {
					W2775 := __e.Get(1)
					_ = W2775
					tmp4559 := MakeNative(func(__e *ControlFlow) {
						W2776 := __e.Get(1)
						_ = W2776
						tmp4575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2776)

						if True == tmp4575 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4560 := MakeNative(func(__e *ControlFlow) {
								W2777 := __e.Get(1)
								_ = W2777
								tmp4561 := MakeNative(func(__e *ControlFlow) {
									W2778 := __e.Get(1)
									_ = W2778
									tmp4571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2778)

									if True == tmp4571 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4562 := MakeNative(func(__e *ControlFlow) {
											W2779 := __e.Get(1)
											_ = W2779
											tmp4563 := MakeNative(func(__e *ControlFlow) {
												W2780 := __e.Get(1)
												_ = W2780
												tmp4567 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2780)

												if True == tmp4567 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp4564 := MakeNative(func(__e *ControlFlow) {
														W2781 := __e.Get(1)
														_ = W2781
														__e.TailApply(PrimFunc(symshen_4comb), W2781, symshen_4skip)
														return
													}, 1)

													tmp4565 := Call(__e, PrimFunc(symshen_4in_1_6), W2780)

													__e.TailApply(tmp4564, tmp4565)
													return

												}

											}, 1)

											tmp4568 := Call(__e, PrimFunc(symshen_4_5returns_6), W2779)

											__e.TailApply(tmp4563, tmp4568)
											return

										}, 1)

										tmp4569 := Call(__e, PrimFunc(symshen_4in_1_6), W2778)

										__e.TailApply(tmp4562, tmp4569)
										return

									}

								}, 1)

								tmp4572 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2777)

								__e.TailApply(tmp4561, tmp4572)
								return

							}, 1)

							tmp4573 := Call(__e, PrimFunc(symshen_4in_1_6), W2776)

							__e.TailApply(tmp4560, tmp4573)
							return

						}

					}, 1)

					tmp4576 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2775)

					__e.TailApply(tmp4559, tmp4576)
					return

				}, 1)

				tmp4577 := Call(__e, PrimFunc(symshen_4in_1_6), W2774)

				__e.TailApply(tmp4558, tmp4577)
				return

			}

		}, 1)

		tmp4580 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2772)

		tmp4581 := Call(__e, tmp4557, tmp4580)

		__e.TailApply(tmp4554, tmp4581)
		return

	}, 1)

	tmp4582 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp4553)

	_ = tmp4582

	tmp4583 := MakeNative(func(__e *ControlFlow) {
		V2782 := __e.Get(1)
		_ = V2782
		tmp4584 := MakeNative(func(__e *ControlFlow) {
			W2783 := __e.Get(1)
			_ = W2783
			tmp4586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2783)

			if True == tmp4586 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2783)
				return
			}

		}, 1)

		tmp4592 := Call(__e, PrimFunc(symshen_4hds_a_2), V2782, MakeNumber(92))

		var ifres4587 Obj

		if True == tmp4592 {
			tmp4588 := MakeNative(func(__e *ControlFlow) {
				W2784 := __e.Get(1)
				_ = W2784
				__e.TailApply(PrimFunc(symshen_4comb), W2784, symshen_4skip)
				return
			}, 1)

			tmp4589 := Call(__e, PrimFunc(symtail), V2782)

			tmp4590 := Call(__e, tmp4588, tmp4589)

			ifres4587 = tmp4590

		} else {
			tmp4591 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4587 = tmp4591

		}

		__e.TailApply(tmp4584, ifres4587)
		return

	}, 1)

	tmp4593 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp4583)

	_ = tmp4593

	tmp4594 := MakeNative(func(__e *ControlFlow) {
		V2785 := __e.Get(1)
		_ = V2785
		tmp4595 := MakeNative(func(__e *ControlFlow) {
			W2786 := __e.Get(1)
			_ = W2786
			tmp4607 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2786)

			if True == tmp4607 {
				tmp4596 := MakeNative(func(__e *ControlFlow) {
					W2791 := __e.Get(1)
					_ = W2791
					tmp4598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2791)

					if True == tmp4598 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2791)
						return
					}

				}, 1)

				tmp4599 := MakeNative(func(__e *ControlFlow) {
					W2792 := __e.Get(1)
					_ = W2792
					tmp4603 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2792)

					if True == tmp4603 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4600 := MakeNative(func(__e *ControlFlow) {
							W2793 := __e.Get(1)
							_ = W2793
							__e.TailApply(PrimFunc(symshen_4comb), W2793, symshen_4skip)
							return
						}, 1)

						tmp4601 := Call(__e, PrimFunc(symshen_4in_1_6), W2792)

						__e.TailApply(tmp4600, tmp4601)
						return

					}

				}, 1)

				tmp4604 := Call(__e, PrimFunc(sym_5e_6), V2785)

				tmp4605 := Call(__e, tmp4599, tmp4604)

				__e.TailApply(tmp4596, tmp4605)
				return

			} else {
				__e.Return(W2786)
				return
			}

		}, 1)

		tmp4608 := MakeNative(func(__e *ControlFlow) {
			W2787 := __e.Get(1)
			_ = W2787
			tmp4618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2787)

			if True == tmp4618 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4609 := MakeNative(func(__e *ControlFlow) {
					W2788 := __e.Get(1)
					_ = W2788
					tmp4610 := MakeNative(func(__e *ControlFlow) {
						W2789 := __e.Get(1)
						_ = W2789
						tmp4614 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2789)

						if True == tmp4614 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4611 := MakeNative(func(__e *ControlFlow) {
								W2790 := __e.Get(1)
								_ = W2790
								__e.TailApply(PrimFunc(symshen_4comb), W2790, symshen_4skip)
								return
							}, 1)

							tmp4612 := Call(__e, PrimFunc(symshen_4in_1_6), W2789)

							__e.TailApply(tmp4611, tmp4612)
							return

						}

					}, 1)

					tmp4615 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2788)

					__e.TailApply(tmp4610, tmp4615)
					return

				}, 1)

				tmp4616 := Call(__e, PrimFunc(symshen_4in_1_6), W2787)

				__e.TailApply(tmp4609, tmp4616)
				return

			}

		}, 1)

		tmp4619 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2785)

		tmp4620 := Call(__e, tmp4608, tmp4619)

		__e.TailApply(tmp4595, tmp4620)
		return

	}, 1)

	tmp4621 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp4594)

	_ = tmp4621

	tmp4622 := MakeNative(func(__e *ControlFlow) {
		V2794 := __e.Get(1)
		_ = V2794
		tmp4623 := MakeNative(func(__e *ControlFlow) {
			W2795 := __e.Get(1)
			_ = W2795
			tmp4625 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2795)

			if True == tmp4625 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2795)
				return
			}

		}, 1)

		tmp4636 := PrimIsPair(V2794)

		var ifres4626 Obj

		if True == tmp4636 {
			tmp4627 := MakeNative(func(__e *ControlFlow) {
				W2796 := __e.Get(1)
				_ = W2796
				tmp4628 := MakeNative(func(__e *ControlFlow) {
					W2797 := __e.Get(1)
					_ = W2797
					tmp4630 := Call(__e, PrimFunc(symshen_4return_2), W2796)

					tmp4631 := PrimNot(tmp4630)

					if True == tmp4631 {
						__e.TailApply(PrimFunc(symshen_4comb), W2797, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4632 := Call(__e, PrimFunc(symtail), V2794)

				__e.TailApply(tmp4628, tmp4632)
				return

			}, 1)

			tmp4633 := Call(__e, PrimFunc(symhead), V2794)

			tmp4634 := Call(__e, tmp4627, tmp4633)

			ifres4626 = tmp4634

		} else {
			tmp4635 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4626 = tmp4635

		}

		__e.TailApply(tmp4623, ifres4626)
		return

	}, 1)

	tmp4637 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp4622)

	_ = tmp4637

	tmp4638 := MakeNative(func(__e *ControlFlow) {
		V2798 := __e.Get(1)
		_ = V2798
		tmp4639 := MakeNative(func(__e *ControlFlow) {
			W2799 := __e.Get(1)
			_ = W2799
			tmp4651 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2799)

			if True == tmp4651 {
				tmp4640 := MakeNative(func(__e *ControlFlow) {
					W2804 := __e.Get(1)
					_ = W2804
					tmp4642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2804)

					if True == tmp4642 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2804)
						return
					}

				}, 1)

				tmp4643 := MakeNative(func(__e *ControlFlow) {
					W2805 := __e.Get(1)
					_ = W2805
					tmp4647 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2805)

					if True == tmp4647 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4644 := MakeNative(func(__e *ControlFlow) {
							W2806 := __e.Get(1)
							_ = W2806
							__e.TailApply(PrimFunc(symshen_4comb), W2806, symshen_4skip)
							return
						}, 1)

						tmp4645 := Call(__e, PrimFunc(symshen_4in_1_6), W2805)

						__e.TailApply(tmp4644, tmp4645)
						return

					}

				}, 1)

				tmp4648 := Call(__e, PrimFunc(symshen_4_5return_6), V2798)

				tmp4649 := Call(__e, tmp4643, tmp4648)

				__e.TailApply(tmp4640, tmp4649)
				return

			} else {
				__e.Return(W2799)
				return
			}

		}, 1)

		tmp4652 := MakeNative(func(__e *ControlFlow) {
			W2800 := __e.Get(1)
			_ = W2800
			tmp4662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2800)

			if True == tmp4662 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4653 := MakeNative(func(__e *ControlFlow) {
					W2801 := __e.Get(1)
					_ = W2801
					tmp4654 := MakeNative(func(__e *ControlFlow) {
						W2802 := __e.Get(1)
						_ = W2802
						tmp4658 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2802)

						if True == tmp4658 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4655 := MakeNative(func(__e *ControlFlow) {
								W2803 := __e.Get(1)
								_ = W2803
								__e.TailApply(PrimFunc(symshen_4comb), W2803, symshen_4skip)
								return
							}, 1)

							tmp4656 := Call(__e, PrimFunc(symshen_4in_1_6), W2802)

							__e.TailApply(tmp4655, tmp4656)
							return

						}

					}, 1)

					tmp4659 := Call(__e, PrimFunc(symshen_4_5returns_6), W2801)

					__e.TailApply(tmp4654, tmp4659)
					return

				}, 1)

				tmp4660 := Call(__e, PrimFunc(symshen_4in_1_6), W2800)

				__e.TailApply(tmp4653, tmp4660)
				return

			}

		}, 1)

		tmp4663 := Call(__e, PrimFunc(symshen_4_5return_6), V2798)

		tmp4664 := Call(__e, tmp4652, tmp4663)

		__e.TailApply(tmp4639, tmp4664)
		return

	}, 1)

	tmp4665 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp4638)

	_ = tmp4665

	tmp4666 := MakeNative(func(__e *ControlFlow) {
		V2807 := __e.Get(1)
		_ = V2807
		tmp4667 := MakeNative(func(__e *ControlFlow) {
			W2808 := __e.Get(1)
			_ = W2808
			tmp4669 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2808)

			if True == tmp4669 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2808)
				return
			}

		}, 1)

		tmp4679 := PrimIsPair(V2807)

		var ifres4670 Obj

		if True == tmp4679 {
			tmp4671 := MakeNative(func(__e *ControlFlow) {
				W2809 := __e.Get(1)
				_ = W2809
				tmp4672 := MakeNative(func(__e *ControlFlow) {
					W2810 := __e.Get(1)
					_ = W2810
					tmp4674 := Call(__e, PrimFunc(symshen_4return_2), W2809)

					if True == tmp4674 {
						__e.TailApply(PrimFunc(symshen_4comb), W2810, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4675 := Call(__e, PrimFunc(symtail), V2807)

				__e.TailApply(tmp4672, tmp4675)
				return

			}, 1)

			tmp4676 := Call(__e, PrimFunc(symhead), V2807)

			tmp4677 := Call(__e, tmp4671, tmp4676)

			ifres4670 = tmp4677

		} else {
			tmp4678 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4670 = tmp4678

		}

		__e.TailApply(tmp4667, ifres4670)
		return

	}, 1)

	tmp4680 := Call(__e, ns2_1set, symshen_4_5return_6, tmp4666)

	_ = tmp4680

	tmp4681 := MakeNative(func(__e *ControlFlow) {
		V2811 := __e.Get(1)
		_ = V2811
		tmp4682 := PrimCons(MakeNumber(13), Nil)

		tmp4683 := PrimCons(MakeNumber(10), tmp4682)

		tmp4684 := PrimCons(MakeNumber(9), tmp4683)

		__e.TailApply(PrimFunc(symelement_2), V2811, tmp4684)
		return

	}, 1)

	tmp4685 := Call(__e, ns2_1set, symshen_4return_2, tmp4681)

	_ = tmp4685

	tmp4686 := MakeNative(func(__e *ControlFlow) {
		V2812 := __e.Get(1)
		_ = V2812
		tmp4687 := MakeNative(func(__e *ControlFlow) {
			W2813 := __e.Get(1)
			_ = W2813
			tmp4689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2813)

			if True == tmp4689 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2813)
				return
			}

		}, 1)

		tmp4690 := MakeNative(func(__e *ControlFlow) {
			W2814 := __e.Get(1)
			_ = W2814
			tmp4706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2814)

			if True == tmp4706 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4691 := MakeNative(func(__e *ControlFlow) {
					W2815 := __e.Get(1)
					_ = W2815
					tmp4692 := MakeNative(func(__e *ControlFlow) {
						W2816 := __e.Get(1)
						_ = W2816
						tmp4702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2816)

						if True == tmp4702 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4693 := MakeNative(func(__e *ControlFlow) {
								W2817 := __e.Get(1)
								_ = W2817
								tmp4694 := MakeNative(func(__e *ControlFlow) {
									W2818 := __e.Get(1)
									_ = W2818
									tmp4698 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2818)

									if True == tmp4698 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4695 := MakeNative(func(__e *ControlFlow) {
											W2819 := __e.Get(1)
											_ = W2819
											__e.TailApply(PrimFunc(symshen_4comb), W2819, symshen_4skip)
											return
										}, 1)

										tmp4696 := Call(__e, PrimFunc(symshen_4in_1_6), W2818)

										__e.TailApply(tmp4695, tmp4696)
										return

									}

								}, 1)

								tmp4699 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2817)

								__e.TailApply(tmp4694, tmp4699)
								return

							}, 1)

							tmp4700 := Call(__e, PrimFunc(symshen_4in_1_6), W2816)

							__e.TailApply(tmp4693, tmp4700)
							return

						}

					}, 1)

					tmp4703 := Call(__e, PrimFunc(symshen_4_5times_6), W2815)

					__e.TailApply(tmp4692, tmp4703)
					return

				}, 1)

				tmp4704 := Call(__e, PrimFunc(symshen_4in_1_6), W2814)

				__e.TailApply(tmp4691, tmp4704)
				return

			}

		}, 1)

		tmp4707 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2812)

		tmp4708 := Call(__e, tmp4690, tmp4707)

		__e.TailApply(tmp4687, tmp4708)
		return

	}, 1)

	tmp4709 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp4686)

	_ = tmp4709

	tmp4710 := MakeNative(func(__e *ControlFlow) {
		V2820 := __e.Get(1)
		_ = V2820
		tmp4711 := MakeNative(func(__e *ControlFlow) {
			W2821 := __e.Get(1)
			_ = W2821
			tmp4713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2821)

			if True == tmp4713 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2821)
				return
			}

		}, 1)

		tmp4719 := Call(__e, PrimFunc(symshen_4hds_a_2), V2820, MakeNumber(42))

		var ifres4714 Obj

		if True == tmp4719 {
			tmp4715 := MakeNative(func(__e *ControlFlow) {
				W2822 := __e.Get(1)
				_ = W2822
				__e.TailApply(PrimFunc(symshen_4comb), W2822, symshen_4skip)
				return
			}, 1)

			tmp4716 := Call(__e, PrimFunc(symtail), V2820)

			tmp4717 := Call(__e, tmp4715, tmp4716)

			ifres4714 = tmp4717

		} else {
			tmp4718 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4714 = tmp4718

		}

		__e.TailApply(tmp4711, ifres4714)
		return

	}, 1)

	tmp4720 := Call(__e, ns2_1set, symshen_4_5times_6, tmp4710)

	_ = tmp4720

	tmp4721 := MakeNative(func(__e *ControlFlow) {
		V2823 := __e.Get(1)
		_ = V2823
		tmp4722 := MakeNative(func(__e *ControlFlow) {
			W2824 := __e.Get(1)
			_ = W2824
			tmp4755 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2824)

			if True == tmp4755 {
				tmp4723 := MakeNative(func(__e *ControlFlow) {
					W2829 := __e.Get(1)
					_ = W2829
					tmp4740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2829)

					if True == tmp4740 {
						tmp4724 := MakeNative(func(__e *ControlFlow) {
							W2834 := __e.Get(1)
							_ = W2834
							tmp4726 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2834)

							if True == tmp4726 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W2834)
								return
							}

						}, 1)

						tmp4738 := PrimIsPair(V2823)

						var ifres4727 Obj

						if True == tmp4738 {
							tmp4728 := MakeNative(func(__e *ControlFlow) {
								W2835 := __e.Get(1)
								_ = W2835
								tmp4729 := MakeNative(func(__e *ControlFlow) {
									W2836 := __e.Get(1)
									_ = W2836
									tmp4733 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2836)

									if True == tmp4733 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4730 := MakeNative(func(__e *ControlFlow) {
											W2837 := __e.Get(1)
											_ = W2837
											__e.TailApply(PrimFunc(symshen_4comb), W2837, symshen_4skip)
											return
										}, 1)

										tmp4731 := Call(__e, PrimFunc(symshen_4in_1_6), W2836)

										__e.TailApply(tmp4730, tmp4731)
										return

									}

								}, 1)

								tmp4734 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2835)

								__e.TailApply(tmp4729, tmp4734)
								return

							}, 1)

							tmp4735 := Call(__e, PrimFunc(symtail), V2823)

							tmp4736 := Call(__e, tmp4728, tmp4735)

							ifres4727 = tmp4736

						} else {
							tmp4737 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres4727 = tmp4737

						}

						__e.TailApply(tmp4724, ifres4727)
						return

					} else {
						__e.Return(W2829)
						return
					}

				}, 1)

				tmp4741 := MakeNative(func(__e *ControlFlow) {
					W2830 := __e.Get(1)
					_ = W2830
					tmp4751 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2830)

					if True == tmp4751 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4742 := MakeNative(func(__e *ControlFlow) {
							W2831 := __e.Get(1)
							_ = W2831
							tmp4743 := MakeNative(func(__e *ControlFlow) {
								W2832 := __e.Get(1)
								_ = W2832
								tmp4747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2832)

								if True == tmp4747 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp4744 := MakeNative(func(__e *ControlFlow) {
										W2833 := __e.Get(1)
										_ = W2833
										__e.TailApply(PrimFunc(symshen_4comb), W2833, symshen_4skip)
										return
									}, 1)

									tmp4745 := Call(__e, PrimFunc(symshen_4in_1_6), W2832)

									__e.TailApply(tmp4744, tmp4745)
									return

								}

							}, 1)

							tmp4748 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2831)

							__e.TailApply(tmp4743, tmp4748)
							return

						}, 1)

						tmp4749 := Call(__e, PrimFunc(symshen_4in_1_6), W2830)

						__e.TailApply(tmp4742, tmp4749)
						return

					}

				}, 1)

				tmp4752 := Call(__e, PrimFunc(symshen_4_5times_6), V2823)

				tmp4753 := Call(__e, tmp4741, tmp4752)

				__e.TailApply(tmp4723, tmp4753)
				return

			} else {
				__e.Return(W2824)
				return
			}

		}, 1)

		tmp4756 := MakeNative(func(__e *ControlFlow) {
			W2825 := __e.Get(1)
			_ = W2825
			tmp4766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2825)

			if True == tmp4766 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4757 := MakeNative(func(__e *ControlFlow) {
					W2826 := __e.Get(1)
					_ = W2826
					tmp4758 := MakeNative(func(__e *ControlFlow) {
						W2827 := __e.Get(1)
						_ = W2827
						tmp4762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2827)

						if True == tmp4762 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4759 := MakeNative(func(__e *ControlFlow) {
								W2828 := __e.Get(1)
								_ = W2828
								__e.TailApply(PrimFunc(symshen_4comb), W2828, symshen_4skip)
								return
							}, 1)

							tmp4760 := Call(__e, PrimFunc(symshen_4in_1_6), W2827)

							__e.TailApply(tmp4759, tmp4760)
							return

						}

					}, 1)

					tmp4763 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2826)

					__e.TailApply(tmp4758, tmp4763)
					return

				}, 1)

				tmp4764 := Call(__e, PrimFunc(symshen_4in_1_6), W2825)

				__e.TailApply(tmp4757, tmp4764)
				return

			}

		}, 1)

		tmp4767 := Call(__e, PrimFunc(symshen_4_5comment_6), V2823)

		tmp4768 := Call(__e, tmp4756, tmp4767)

		__e.TailApply(tmp4722, tmp4768)
		return

	}, 1)

	tmp4769 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp4721)

	_ = tmp4769

	tmp4770 := MakeNative(func(__e *ControlFlow) {
		V2838 := __e.Get(1)
		_ = V2838
		tmp4771 := MakeNative(func(__e *ControlFlow) {
			W2839 := __e.Get(1)
			_ = W2839
			tmp4802 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2839)

			if True == tmp4802 {
				tmp4772 := MakeNative(func(__e *ControlFlow) {
					W2843 := __e.Get(1)
					_ = W2843
					tmp4791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2843)

					if True == tmp4791 {
						tmp4773 := MakeNative(func(__e *ControlFlow) {
							W2847 := __e.Get(1)
							_ = W2847
							tmp4775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2847)

							if True == tmp4775 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W2847)
								return
							}

						}, 1)

						tmp4776 := MakeNative(func(__e *ControlFlow) {
							W2848 := __e.Get(1)
							_ = W2848
							tmp4787 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2848)

							if True == tmp4787 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4777 := MakeNative(func(__e *ControlFlow) {
									W2849 := __e.Get(1)
									_ = W2849
									tmp4778 := MakeNative(func(__e *ControlFlow) {
										W2850 := __e.Get(1)
										_ = W2850
										tmp4783 := PrimEqual(W2849, MakeString("<>"))

										var ifres4779 Obj

										if True == tmp4783 {
											tmp4780 := PrimCons(MakeNumber(0), Nil)

											tmp4781 := PrimCons(symvector, tmp4780)

											ifres4779 = tmp4781

										} else {
											tmp4782 := PrimIntern(W2849)

											ifres4779 = tmp4782

										}

										__e.TailApply(PrimFunc(symshen_4comb), W2850, ifres4779)
										return

									}, 1)

									tmp4784 := Call(__e, PrimFunc(symshen_4in_1_6), W2848)

									__e.TailApply(tmp4778, tmp4784)
									return

								}, 1)

								tmp4785 := Call(__e, PrimFunc(symshen_4_5_1out), W2848)

								__e.TailApply(tmp4777, tmp4785)
								return

							}

						}, 1)

						tmp4788 := Call(__e, PrimFunc(symshen_4_5sym_6), V2838)

						tmp4789 := Call(__e, tmp4776, tmp4788)

						__e.TailApply(tmp4773, tmp4789)
						return

					} else {
						__e.Return(W2843)
						return
					}

				}, 1)

				tmp4792 := MakeNative(func(__e *ControlFlow) {
					W2844 := __e.Get(1)
					_ = W2844
					tmp4798 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2844)

					if True == tmp4798 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4793 := MakeNative(func(__e *ControlFlow) {
							W2845 := __e.Get(1)
							_ = W2845
							tmp4794 := MakeNative(func(__e *ControlFlow) {
								W2846 := __e.Get(1)
								_ = W2846
								__e.TailApply(PrimFunc(symshen_4comb), W2846, W2845)
								return
							}, 1)

							tmp4795 := Call(__e, PrimFunc(symshen_4in_1_6), W2844)

							__e.TailApply(tmp4794, tmp4795)
							return

						}, 1)

						tmp4796 := Call(__e, PrimFunc(symshen_4_5_1out), W2844)

						__e.TailApply(tmp4793, tmp4796)
						return

					}

				}, 1)

				tmp4799 := Call(__e, PrimFunc(symshen_4_5number_6), V2838)

				tmp4800 := Call(__e, tmp4792, tmp4799)

				__e.TailApply(tmp4772, tmp4800)
				return

			} else {
				__e.Return(W2839)
				return
			}

		}, 1)

		tmp4803 := MakeNative(func(__e *ControlFlow) {
			W2840 := __e.Get(1)
			_ = W2840
			tmp4809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2840)

			if True == tmp4809 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4804 := MakeNative(func(__e *ControlFlow) {
					W2841 := __e.Get(1)
					_ = W2841
					tmp4805 := MakeNative(func(__e *ControlFlow) {
						W2842 := __e.Get(1)
						_ = W2842
						__e.TailApply(PrimFunc(symshen_4comb), W2842, W2841)
						return
					}, 1)

					tmp4806 := Call(__e, PrimFunc(symshen_4in_1_6), W2840)

					__e.TailApply(tmp4805, tmp4806)
					return

				}, 1)

				tmp4807 := Call(__e, PrimFunc(symshen_4_5_1out), W2840)

				__e.TailApply(tmp4804, tmp4807)
				return

			}

		}, 1)

		tmp4810 := Call(__e, PrimFunc(symshen_4_5str_6), V2838)

		tmp4811 := Call(__e, tmp4803, tmp4810)

		__e.TailApply(tmp4771, tmp4811)
		return

	}, 1)

	tmp4812 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp4770)

	_ = tmp4812

	tmp4813 := MakeNative(func(__e *ControlFlow) {
		V2851 := __e.Get(1)
		_ = V2851
		tmp4814 := MakeNative(func(__e *ControlFlow) {
			W2852 := __e.Get(1)
			_ = W2852
			tmp4816 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2852)

			if True == tmp4816 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2852)
				return
			}

		}, 1)

		tmp4817 := MakeNative(func(__e *ControlFlow) {
			W2853 := __e.Get(1)
			_ = W2853
			tmp4832 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2853)

			if True == tmp4832 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4818 := MakeNative(func(__e *ControlFlow) {
					W2854 := __e.Get(1)
					_ = W2854
					tmp4819 := MakeNative(func(__e *ControlFlow) {
						W2855 := __e.Get(1)
						_ = W2855
						tmp4820 := MakeNative(func(__e *ControlFlow) {
							W2856 := __e.Get(1)
							_ = W2856
							tmp4827 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2856)

							if True == tmp4827 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4821 := MakeNative(func(__e *ControlFlow) {
									W2857 := __e.Get(1)
									_ = W2857
									tmp4822 := MakeNative(func(__e *ControlFlow) {
										W2858 := __e.Get(1)
										_ = W2858
										tmp4823 := PrimStringConcat(W2854, W2857)

										__e.TailApply(PrimFunc(symshen_4comb), W2858, tmp4823)
										return

									}, 1)

									tmp4824 := Call(__e, PrimFunc(symshen_4in_1_6), W2856)

									__e.TailApply(tmp4822, tmp4824)
									return

								}, 1)

								tmp4825 := Call(__e, PrimFunc(symshen_4_5_1out), W2856)

								__e.TailApply(tmp4821, tmp4825)
								return

							}

						}, 1)

						tmp4828 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2855)

						__e.TailApply(tmp4820, tmp4828)
						return

					}, 1)

					tmp4829 := Call(__e, PrimFunc(symshen_4in_1_6), W2853)

					__e.TailApply(tmp4819, tmp4829)
					return

				}, 1)

				tmp4830 := Call(__e, PrimFunc(symshen_4_5_1out), W2853)

				__e.TailApply(tmp4818, tmp4830)
				return

			}

		}, 1)

		tmp4833 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2851)

		tmp4834 := Call(__e, tmp4817, tmp4833)

		__e.TailApply(tmp4814, tmp4834)
		return

	}, 1)

	tmp4835 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp4813)

	_ = tmp4835

	tmp4836 := MakeNative(func(__e *ControlFlow) {
		V2859 := __e.Get(1)
		_ = V2859
		tmp4837 := MakeNative(func(__e *ControlFlow) {
			W2860 := __e.Get(1)
			_ = W2860
			tmp4839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2860)

			if True == tmp4839 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2860)
				return
			}

		}, 1)

		tmp4850 := PrimIsPair(V2859)

		var ifres4840 Obj

		if True == tmp4850 {
			tmp4841 := MakeNative(func(__e *ControlFlow) {
				W2861 := __e.Get(1)
				_ = W2861
				tmp4842 := MakeNative(func(__e *ControlFlow) {
					W2862 := __e.Get(1)
					_ = W2862
					tmp4845 := Call(__e, PrimFunc(symshen_4alpha_2), W2861)

					if True == tmp4845 {
						tmp4843 := PrimNumberToString(W2861)

						__e.TailApply(PrimFunc(symshen_4comb), W2862, tmp4843)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4846 := Call(__e, PrimFunc(symtail), V2859)

				__e.TailApply(tmp4842, tmp4846)
				return

			}, 1)

			tmp4847 := Call(__e, PrimFunc(symhead), V2859)

			tmp4848 := Call(__e, tmp4841, tmp4847)

			ifres4840 = tmp4848

		} else {
			tmp4849 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4840 = tmp4849

		}

		__e.TailApply(tmp4837, ifres4840)
		return

	}, 1)

	tmp4851 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp4836)

	_ = tmp4851

	tmp4852 := MakeNative(func(__e *ControlFlow) {
		V2863 := __e.Get(1)
		_ = V2863
		tmp4859 := Call(__e, PrimFunc(symshen_4lowercase_2), V2863)

		if True == tmp4859 {
			__e.Return(True)
			return
		} else {
			tmp4857 := Call(__e, PrimFunc(symshen_4uppercase_2), V2863)

			var ifres4854 Obj

			if True == tmp4857 {
				ifres4854 = True

			} else {
				tmp4856 := Call(__e, PrimFunc(symshen_4misc_2), V2863)

				var ifres4855 Obj

				if True == tmp4856 {
					ifres4855 = True

				} else {
					ifres4855 = False

				}

				ifres4854 = ifres4855

			}

			if True == ifres4854 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp4860 := Call(__e, ns2_1set, symshen_4alpha_2, tmp4852)

	_ = tmp4860

	tmp4861 := MakeNative(func(__e *ControlFlow) {
		V2864 := __e.Get(1)
		_ = V2864
		tmp4865 := PrimGreatEqual(V2864, MakeNumber(97))

		if True == tmp4865 {
			tmp4863 := PrimLessEqual(V2864, MakeNumber(122))

			if True == tmp4863 {
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

	tmp4866 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp4861)

	_ = tmp4866

	tmp4867 := MakeNative(func(__e *ControlFlow) {
		V2865 := __e.Get(1)
		_ = V2865
		tmp4871 := PrimGreatEqual(V2865, MakeNumber(65))

		if True == tmp4871 {
			tmp4869 := PrimLessEqual(V2865, MakeNumber(90))

			if True == tmp4869 {
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

	tmp4872 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp4867)

	_ = tmp4872

	tmp4873 := MakeNative(func(__e *ControlFlow) {
		V2866 := __e.Get(1)
		_ = V2866
		tmp4874 := PrimCons(MakeNumber(96), Nil)

		tmp4875 := PrimCons(MakeNumber(35), tmp4874)

		tmp4876 := PrimCons(MakeNumber(39), tmp4875)

		tmp4877 := PrimCons(MakeNumber(37), tmp4876)

		tmp4878 := PrimCons(MakeNumber(38), tmp4877)

		tmp4879 := PrimCons(MakeNumber(60), tmp4878)

		tmp4880 := PrimCons(MakeNumber(62), tmp4879)

		tmp4881 := PrimCons(MakeNumber(46), tmp4880)

		tmp4882 := PrimCons(MakeNumber(126), tmp4881)

		tmp4883 := PrimCons(MakeNumber(64), tmp4882)

		tmp4884 := PrimCons(MakeNumber(33), tmp4883)

		tmp4885 := PrimCons(MakeNumber(36), tmp4884)

		tmp4886 := PrimCons(MakeNumber(63), tmp4885)

		tmp4887 := PrimCons(MakeNumber(95), tmp4886)

		tmp4888 := PrimCons(MakeNumber(43), tmp4887)

		tmp4889 := PrimCons(MakeNumber(47), tmp4888)

		tmp4890 := PrimCons(MakeNumber(42), tmp4889)

		tmp4891 := PrimCons(MakeNumber(45), tmp4890)

		tmp4892 := PrimCons(MakeNumber(61), tmp4891)

		__e.TailApply(PrimFunc(symelement_2), V2866, tmp4892)
		return

	}, 1)

	tmp4893 := Call(__e, ns2_1set, symshen_4misc_2, tmp4873)

	_ = tmp4893

	tmp4894 := MakeNative(func(__e *ControlFlow) {
		V2867 := __e.Get(1)
		_ = V2867
		tmp4895 := MakeNative(func(__e *ControlFlow) {
			W2868 := __e.Get(1)
			_ = W2868
			tmp4907 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2868)

			if True == tmp4907 {
				tmp4896 := MakeNative(func(__e *ControlFlow) {
					W2875 := __e.Get(1)
					_ = W2875
					tmp4898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2875)

					if True == tmp4898 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2875)
						return
					}

				}, 1)

				tmp4899 := MakeNative(func(__e *ControlFlow) {
					W2876 := __e.Get(1)
					_ = W2876
					tmp4903 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2876)

					if True == tmp4903 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4900 := MakeNative(func(__e *ControlFlow) {
							W2877 := __e.Get(1)
							_ = W2877
							__e.TailApply(PrimFunc(symshen_4comb), W2877, MakeString(""))
							return
						}, 1)

						tmp4901 := Call(__e, PrimFunc(symshen_4in_1_6), W2876)

						__e.TailApply(tmp4900, tmp4901)
						return

					}

				}, 1)

				tmp4904 := Call(__e, PrimFunc(sym_5e_6), V2867)

				tmp4905 := Call(__e, tmp4899, tmp4904)

				__e.TailApply(tmp4896, tmp4905)
				return

			} else {
				__e.Return(W2868)
				return
			}

		}, 1)

		tmp4908 := MakeNative(func(__e *ControlFlow) {
			W2869 := __e.Get(1)
			_ = W2869
			tmp4923 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2869)

			if True == tmp4923 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4909 := MakeNative(func(__e *ControlFlow) {
					W2870 := __e.Get(1)
					_ = W2870
					tmp4910 := MakeNative(func(__e *ControlFlow) {
						W2871 := __e.Get(1)
						_ = W2871
						tmp4911 := MakeNative(func(__e *ControlFlow) {
							W2872 := __e.Get(1)
							_ = W2872
							tmp4918 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2872)

							if True == tmp4918 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4912 := MakeNative(func(__e *ControlFlow) {
									W2873 := __e.Get(1)
									_ = W2873
									tmp4913 := MakeNative(func(__e *ControlFlow) {
										W2874 := __e.Get(1)
										_ = W2874
										tmp4914 := PrimStringConcat(W2870, W2873)

										__e.TailApply(PrimFunc(symshen_4comb), W2874, tmp4914)
										return

									}, 1)

									tmp4915 := Call(__e, PrimFunc(symshen_4in_1_6), W2872)

									__e.TailApply(tmp4913, tmp4915)
									return

								}, 1)

								tmp4916 := Call(__e, PrimFunc(symshen_4_5_1out), W2872)

								__e.TailApply(tmp4912, tmp4916)
								return

							}

						}, 1)

						tmp4919 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2871)

						__e.TailApply(tmp4911, tmp4919)
						return

					}, 1)

					tmp4920 := Call(__e, PrimFunc(symshen_4in_1_6), W2869)

					__e.TailApply(tmp4910, tmp4920)
					return

				}, 1)

				tmp4921 := Call(__e, PrimFunc(symshen_4_5_1out), W2869)

				__e.TailApply(tmp4909, tmp4921)
				return

			}

		}, 1)

		tmp4924 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2867)

		tmp4925 := Call(__e, tmp4908, tmp4924)

		__e.TailApply(tmp4895, tmp4925)
		return

	}, 1)

	tmp4926 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp4894)

	_ = tmp4926

	tmp4927 := MakeNative(func(__e *ControlFlow) {
		V2878 := __e.Get(1)
		_ = V2878
		tmp4928 := MakeNative(func(__e *ControlFlow) {
			W2879 := __e.Get(1)
			_ = W2879
			tmp4942 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2879)

			if True == tmp4942 {
				tmp4929 := MakeNative(func(__e *ControlFlow) {
					W2883 := __e.Get(1)
					_ = W2883
					tmp4931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2883)

					if True == tmp4931 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2883)
						return
					}

				}, 1)

				tmp4932 := MakeNative(func(__e *ControlFlow) {
					W2884 := __e.Get(1)
					_ = W2884
					tmp4938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2884)

					if True == tmp4938 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4933 := MakeNative(func(__e *ControlFlow) {
							W2885 := __e.Get(1)
							_ = W2885
							tmp4934 := MakeNative(func(__e *ControlFlow) {
								W2886 := __e.Get(1)
								_ = W2886
								__e.TailApply(PrimFunc(symshen_4comb), W2886, W2885)
								return
							}, 1)

							tmp4935 := Call(__e, PrimFunc(symshen_4in_1_6), W2884)

							__e.TailApply(tmp4934, tmp4935)
							return

						}, 1)

						tmp4936 := Call(__e, PrimFunc(symshen_4_5_1out), W2884)

						__e.TailApply(tmp4933, tmp4936)
						return

					}

				}, 1)

				tmp4939 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2878)

				tmp4940 := Call(__e, tmp4932, tmp4939)

				__e.TailApply(tmp4929, tmp4940)
				return

			} else {
				__e.Return(W2879)
				return
			}

		}, 1)

		tmp4943 := MakeNative(func(__e *ControlFlow) {
			W2880 := __e.Get(1)
			_ = W2880
			tmp4949 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2880)

			if True == tmp4949 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4944 := MakeNative(func(__e *ControlFlow) {
					W2881 := __e.Get(1)
					_ = W2881
					tmp4945 := MakeNative(func(__e *ControlFlow) {
						W2882 := __e.Get(1)
						_ = W2882
						__e.TailApply(PrimFunc(symshen_4comb), W2882, W2881)
						return
					}, 1)

					tmp4946 := Call(__e, PrimFunc(symshen_4in_1_6), W2880)

					__e.TailApply(tmp4945, tmp4946)
					return

				}, 1)

				tmp4947 := Call(__e, PrimFunc(symshen_4_5_1out), W2880)

				__e.TailApply(tmp4944, tmp4947)
				return

			}

		}, 1)

		tmp4950 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2878)

		tmp4951 := Call(__e, tmp4943, tmp4950)

		__e.TailApply(tmp4928, tmp4951)
		return

	}, 1)

	tmp4952 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp4927)

	_ = tmp4952

	tmp4953 := MakeNative(func(__e *ControlFlow) {
		V2887 := __e.Get(1)
		_ = V2887
		tmp4954 := MakeNative(func(__e *ControlFlow) {
			W2888 := __e.Get(1)
			_ = W2888
			tmp4956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2888)

			if True == tmp4956 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2888)
				return
			}

		}, 1)

		tmp4967 := PrimIsPair(V2887)

		var ifres4957 Obj

		if True == tmp4967 {
			tmp4958 := MakeNative(func(__e *ControlFlow) {
				W2889 := __e.Get(1)
				_ = W2889
				tmp4959 := MakeNative(func(__e *ControlFlow) {
					W2890 := __e.Get(1)
					_ = W2890
					tmp4962 := Call(__e, PrimFunc(symshen_4digit_2), W2889)

					if True == tmp4962 {
						tmp4960 := PrimNumberToString(W2889)

						__e.TailApply(PrimFunc(symshen_4comb), W2890, tmp4960)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4963 := Call(__e, PrimFunc(symtail), V2887)

				__e.TailApply(tmp4959, tmp4963)
				return

			}, 1)

			tmp4964 := Call(__e, PrimFunc(symhead), V2887)

			tmp4965 := Call(__e, tmp4958, tmp4964)

			ifres4957 = tmp4965

		} else {
			tmp4966 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4957 = tmp4966

		}

		__e.TailApply(tmp4954, ifres4957)
		return

	}, 1)

	tmp4968 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp4953)

	_ = tmp4968

	tmp4969 := MakeNative(func(__e *ControlFlow) {
		V2891 := __e.Get(1)
		_ = V2891
		tmp4973 := PrimGreatEqual(V2891, MakeNumber(48))

		if True == tmp4973 {
			tmp4971 := PrimLessEqual(V2891, MakeNumber(57))

			if True == tmp4971 {
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

	tmp4974 := Call(__e, ns2_1set, symshen_4digit_2, tmp4969)

	_ = tmp4974

	tmp4975 := MakeNative(func(__e *ControlFlow) {
		V2892 := __e.Get(1)
		_ = V2892
		tmp4976 := MakeNative(func(__e *ControlFlow) {
			W2893 := __e.Get(1)
			_ = W2893
			tmp4978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2893)

			if True == tmp4978 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2893)
				return
			}

		}, 1)

		tmp4979 := MakeNative(func(__e *ControlFlow) {
			W2894 := __e.Get(1)
			_ = W2894
			tmp4997 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2894)

			if True == tmp4997 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4980 := MakeNative(func(__e *ControlFlow) {
					W2895 := __e.Get(1)
					_ = W2895
					tmp4981 := MakeNative(func(__e *ControlFlow) {
						W2896 := __e.Get(1)
						_ = W2896
						tmp4993 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2896)

						if True == tmp4993 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp4982 := MakeNative(func(__e *ControlFlow) {
								W2897 := __e.Get(1)
								_ = W2897
								tmp4983 := MakeNative(func(__e *ControlFlow) {
									W2898 := __e.Get(1)
									_ = W2898
									tmp4984 := MakeNative(func(__e *ControlFlow) {
										W2899 := __e.Get(1)
										_ = W2899
										tmp4988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2899)

										if True == tmp4988 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp4985 := MakeNative(func(__e *ControlFlow) {
												W2900 := __e.Get(1)
												_ = W2900
												__e.TailApply(PrimFunc(symshen_4comb), W2900, W2897)
												return
											}, 1)

											tmp4986 := Call(__e, PrimFunc(symshen_4in_1_6), W2899)

											__e.TailApply(tmp4985, tmp4986)
											return

										}

									}, 1)

									tmp4989 := Call(__e, PrimFunc(symshen_4_5dbq_6), W2898)

									__e.TailApply(tmp4984, tmp4989)
									return

								}, 1)

								tmp4990 := Call(__e, PrimFunc(symshen_4in_1_6), W2896)

								__e.TailApply(tmp4983, tmp4990)
								return

							}, 1)

							tmp4991 := Call(__e, PrimFunc(symshen_4_5_1out), W2896)

							__e.TailApply(tmp4982, tmp4991)
							return

						}

					}, 1)

					tmp4994 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2895)

					__e.TailApply(tmp4981, tmp4994)
					return

				}, 1)

				tmp4995 := Call(__e, PrimFunc(symshen_4in_1_6), W2894)

				__e.TailApply(tmp4980, tmp4995)
				return

			}

		}, 1)

		tmp4998 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2892)

		tmp4999 := Call(__e, tmp4979, tmp4998)

		__e.TailApply(tmp4976, tmp4999)
		return

	}, 1)

	tmp5000 := Call(__e, ns2_1set, symshen_4_5str_6, tmp4975)

	_ = tmp5000

	tmp5001 := MakeNative(func(__e *ControlFlow) {
		V2901 := __e.Get(1)
		_ = V2901
		tmp5002 := MakeNative(func(__e *ControlFlow) {
			W2902 := __e.Get(1)
			_ = W2902
			tmp5004 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2902)

			if True == tmp5004 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2902)
				return
			}

		}, 1)

		tmp5010 := Call(__e, PrimFunc(symshen_4hds_a_2), V2901, MakeNumber(34))

		var ifres5005 Obj

		if True == tmp5010 {
			tmp5006 := MakeNative(func(__e *ControlFlow) {
				W2903 := __e.Get(1)
				_ = W2903
				__e.TailApply(PrimFunc(symshen_4comb), W2903, symshen_4skip)
				return
			}, 1)

			tmp5007 := Call(__e, PrimFunc(symtail), V2901)

			tmp5008 := Call(__e, tmp5006, tmp5007)

			ifres5005 = tmp5008

		} else {
			tmp5009 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5005 = tmp5009

		}

		__e.TailApply(tmp5002, ifres5005)
		return

	}, 1)

	tmp5011 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp5001)

	_ = tmp5011

	tmp5012 := MakeNative(func(__e *ControlFlow) {
		V2904 := __e.Get(1)
		_ = V2904
		tmp5013 := MakeNative(func(__e *ControlFlow) {
			W2905 := __e.Get(1)
			_ = W2905
			tmp5025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2905)

			if True == tmp5025 {
				tmp5014 := MakeNative(func(__e *ControlFlow) {
					W2912 := __e.Get(1)
					_ = W2912
					tmp5016 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2912)

					if True == tmp5016 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2912)
						return
					}

				}, 1)

				tmp5017 := MakeNative(func(__e *ControlFlow) {
					W2913 := __e.Get(1)
					_ = W2913
					tmp5021 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2913)

					if True == tmp5021 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5018 := MakeNative(func(__e *ControlFlow) {
							W2914 := __e.Get(1)
							_ = W2914
							__e.TailApply(PrimFunc(symshen_4comb), W2914, MakeString(""))
							return
						}, 1)

						tmp5019 := Call(__e, PrimFunc(symshen_4in_1_6), W2913)

						__e.TailApply(tmp5018, tmp5019)
						return

					}

				}, 1)

				tmp5022 := Call(__e, PrimFunc(sym_5e_6), V2904)

				tmp5023 := Call(__e, tmp5017, tmp5022)

				__e.TailApply(tmp5014, tmp5023)
				return

			} else {
				__e.Return(W2905)
				return
			}

		}, 1)

		tmp5026 := MakeNative(func(__e *ControlFlow) {
			W2906 := __e.Get(1)
			_ = W2906
			tmp5041 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2906)

			if True == tmp5041 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5027 := MakeNative(func(__e *ControlFlow) {
					W2907 := __e.Get(1)
					_ = W2907
					tmp5028 := MakeNative(func(__e *ControlFlow) {
						W2908 := __e.Get(1)
						_ = W2908
						tmp5029 := MakeNative(func(__e *ControlFlow) {
							W2909 := __e.Get(1)
							_ = W2909
							tmp5036 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2909)

							if True == tmp5036 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5030 := MakeNative(func(__e *ControlFlow) {
									W2910 := __e.Get(1)
									_ = W2910
									tmp5031 := MakeNative(func(__e *ControlFlow) {
										W2911 := __e.Get(1)
										_ = W2911
										tmp5032 := PrimStringConcat(W2907, W2910)

										__e.TailApply(PrimFunc(symshen_4comb), W2911, tmp5032)
										return

									}, 1)

									tmp5033 := Call(__e, PrimFunc(symshen_4in_1_6), W2909)

									__e.TailApply(tmp5031, tmp5033)
									return

								}, 1)

								tmp5034 := Call(__e, PrimFunc(symshen_4_5_1out), W2909)

								__e.TailApply(tmp5030, tmp5034)
								return

							}

						}, 1)

						tmp5037 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2908)

						__e.TailApply(tmp5029, tmp5037)
						return

					}, 1)

					tmp5038 := Call(__e, PrimFunc(symshen_4in_1_6), W2906)

					__e.TailApply(tmp5028, tmp5038)
					return

				}, 1)

				tmp5039 := Call(__e, PrimFunc(symshen_4_5_1out), W2906)

				__e.TailApply(tmp5027, tmp5039)
				return

			}

		}, 1)

		tmp5042 := Call(__e, PrimFunc(symshen_4_5strc_6), V2904)

		tmp5043 := Call(__e, tmp5026, tmp5042)

		__e.TailApply(tmp5013, tmp5043)
		return

	}, 1)

	tmp5044 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp5012)

	_ = tmp5044

	tmp5045 := MakeNative(func(__e *ControlFlow) {
		V2915 := __e.Get(1)
		_ = V2915
		tmp5046 := MakeNative(func(__e *ControlFlow) {
			W2916 := __e.Get(1)
			_ = W2916
			tmp5060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2916)

			if True == tmp5060 {
				tmp5047 := MakeNative(func(__e *ControlFlow) {
					W2920 := __e.Get(1)
					_ = W2920
					tmp5049 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2920)

					if True == tmp5049 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2920)
						return
					}

				}, 1)

				tmp5050 := MakeNative(func(__e *ControlFlow) {
					W2921 := __e.Get(1)
					_ = W2921
					tmp5056 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2921)

					if True == tmp5056 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5051 := MakeNative(func(__e *ControlFlow) {
							W2922 := __e.Get(1)
							_ = W2922
							tmp5052 := MakeNative(func(__e *ControlFlow) {
								W2923 := __e.Get(1)
								_ = W2923
								__e.TailApply(PrimFunc(symshen_4comb), W2923, W2922)
								return
							}, 1)

							tmp5053 := Call(__e, PrimFunc(symshen_4in_1_6), W2921)

							__e.TailApply(tmp5052, tmp5053)
							return

						}, 1)

						tmp5054 := Call(__e, PrimFunc(symshen_4_5_1out), W2921)

						__e.TailApply(tmp5051, tmp5054)
						return

					}

				}, 1)

				tmp5057 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2915)

				tmp5058 := Call(__e, tmp5050, tmp5057)

				__e.TailApply(tmp5047, tmp5058)
				return

			} else {
				__e.Return(W2916)
				return
			}

		}, 1)

		tmp5061 := MakeNative(func(__e *ControlFlow) {
			W2917 := __e.Get(1)
			_ = W2917
			tmp5067 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2917)

			if True == tmp5067 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5062 := MakeNative(func(__e *ControlFlow) {
					W2918 := __e.Get(1)
					_ = W2918
					tmp5063 := MakeNative(func(__e *ControlFlow) {
						W2919 := __e.Get(1)
						_ = W2919
						__e.TailApply(PrimFunc(symshen_4comb), W2919, W2918)
						return
					}, 1)

					tmp5064 := Call(__e, PrimFunc(symshen_4in_1_6), W2917)

					__e.TailApply(tmp5063, tmp5064)
					return

				}, 1)

				tmp5065 := Call(__e, PrimFunc(symshen_4_5_1out), W2917)

				__e.TailApply(tmp5062, tmp5065)
				return

			}

		}, 1)

		tmp5068 := Call(__e, PrimFunc(symshen_4_5control_6), V2915)

		tmp5069 := Call(__e, tmp5061, tmp5068)

		__e.TailApply(tmp5046, tmp5069)
		return

	}, 1)

	tmp5070 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp5045)

	_ = tmp5070

	tmp5071 := MakeNative(func(__e *ControlFlow) {
		V2924 := __e.Get(1)
		_ = V2924
		tmp5072 := MakeNative(func(__e *ControlFlow) {
			W2925 := __e.Get(1)
			_ = W2925
			tmp5074 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2925)

			if True == tmp5074 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2925)
				return
			}

		}, 1)

		tmp5075 := MakeNative(func(__e *ControlFlow) {
			W2926 := __e.Get(1)
			_ = W2926
			tmp5100 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2926)

			if True == tmp5100 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5076 := MakeNative(func(__e *ControlFlow) {
					W2927 := __e.Get(1)
					_ = W2927
					tmp5077 := MakeNative(func(__e *ControlFlow) {
						W2928 := __e.Get(1)
						_ = W2928
						tmp5096 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2928)

						if True == tmp5096 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp5078 := MakeNative(func(__e *ControlFlow) {
								W2929 := __e.Get(1)
								_ = W2929
								tmp5079 := MakeNative(func(__e *ControlFlow) {
									W2930 := __e.Get(1)
									_ = W2930
									tmp5092 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2930)

									if True == tmp5092 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5080 := MakeNative(func(__e *ControlFlow) {
											W2931 := __e.Get(1)
											_ = W2931
											tmp5081 := MakeNative(func(__e *ControlFlow) {
												W2932 := __e.Get(1)
												_ = W2932
												tmp5082 := MakeNative(func(__e *ControlFlow) {
													W2933 := __e.Get(1)
													_ = W2933
													tmp5087 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2933)

													if True == tmp5087 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp5083 := MakeNative(func(__e *ControlFlow) {
															W2934 := __e.Get(1)
															_ = W2934
															tmp5084 := PrimNumberToString(W2931)

															__e.TailApply(PrimFunc(symshen_4comb), W2934, tmp5084)
															return

														}, 1)

														tmp5085 := Call(__e, PrimFunc(symshen_4in_1_6), W2933)

														__e.TailApply(tmp5083, tmp5085)
														return

													}

												}, 1)

												tmp5088 := Call(__e, PrimFunc(symshen_4_5semicolon_6), W2932)

												__e.TailApply(tmp5082, tmp5088)
												return

											}, 1)

											tmp5089 := Call(__e, PrimFunc(symshen_4in_1_6), W2930)

											__e.TailApply(tmp5081, tmp5089)
											return

										}, 1)

										tmp5090 := Call(__e, PrimFunc(symshen_4_5_1out), W2930)

										__e.TailApply(tmp5080, tmp5090)
										return

									}

								}, 1)

								tmp5093 := Call(__e, PrimFunc(symshen_4_5integer_6), W2929)

								__e.TailApply(tmp5079, tmp5093)
								return

							}, 1)

							tmp5094 := Call(__e, PrimFunc(symshen_4in_1_6), W2928)

							__e.TailApply(tmp5078, tmp5094)
							return

						}

					}, 1)

					tmp5097 := Call(__e, PrimFunc(symshen_4_5hash_6), W2927)

					__e.TailApply(tmp5077, tmp5097)
					return

				}, 1)

				tmp5098 := Call(__e, PrimFunc(symshen_4in_1_6), W2926)

				__e.TailApply(tmp5076, tmp5098)
				return

			}

		}, 1)

		tmp5101 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2924)

		tmp5102 := Call(__e, tmp5075, tmp5101)

		__e.TailApply(tmp5072, tmp5102)
		return

	}, 1)

	tmp5103 := Call(__e, ns2_1set, symshen_4_5control_6, tmp5071)

	_ = tmp5103

	tmp5104 := MakeNative(func(__e *ControlFlow) {
		V2935 := __e.Get(1)
		_ = V2935
		tmp5105 := MakeNative(func(__e *ControlFlow) {
			W2936 := __e.Get(1)
			_ = W2936
			tmp5107 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2936)

			if True == tmp5107 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2936)
				return
			}

		}, 1)

		tmp5119 := PrimIsPair(V2935)

		var ifres5108 Obj

		if True == tmp5119 {
			tmp5109 := MakeNative(func(__e *ControlFlow) {
				W2937 := __e.Get(1)
				_ = W2937
				tmp5110 := MakeNative(func(__e *ControlFlow) {
					W2938 := __e.Get(1)
					_ = W2938
					tmp5113 := PrimEqual(W2937, MakeNumber(34))

					tmp5114 := PrimNot(tmp5113)

					if True == tmp5114 {
						tmp5111 := PrimNumberToString(W2937)

						__e.TailApply(PrimFunc(symshen_4comb), W2938, tmp5111)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5115 := Call(__e, PrimFunc(symtail), V2935)

				__e.TailApply(tmp5110, tmp5115)
				return

			}, 1)

			tmp5116 := Call(__e, PrimFunc(symhead), V2935)

			tmp5117 := Call(__e, tmp5109, tmp5116)

			ifres5108 = tmp5117

		} else {
			tmp5118 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5108 = tmp5118

		}

		__e.TailApply(tmp5105, ifres5108)
		return

	}, 1)

	tmp5120 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp5104)

	_ = tmp5120

	tmp5121 := MakeNative(func(__e *ControlFlow) {
		V2939 := __e.Get(1)
		_ = V2939
		tmp5122 := MakeNative(func(__e *ControlFlow) {
			W2940 := __e.Get(1)
			_ = W2940
			tmp5124 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2940)

			if True == tmp5124 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2940)
				return
			}

		}, 1)

		tmp5130 := Call(__e, PrimFunc(symshen_4hds_a_2), V2939, MakeNumber(99))

		var ifres5125 Obj

		if True == tmp5130 {
			tmp5126 := MakeNative(func(__e *ControlFlow) {
				W2941 := __e.Get(1)
				_ = W2941
				__e.TailApply(PrimFunc(symshen_4comb), W2941, symshen_4skip)
				return
			}, 1)

			tmp5127 := Call(__e, PrimFunc(symtail), V2939)

			tmp5128 := Call(__e, tmp5126, tmp5127)

			ifres5125 = tmp5128

		} else {
			tmp5129 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5125 = tmp5129

		}

		__e.TailApply(tmp5122, ifres5125)
		return

	}, 1)

	tmp5131 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp5121)

	_ = tmp5131

	tmp5132 := MakeNative(func(__e *ControlFlow) {
		V2942 := __e.Get(1)
		_ = V2942
		tmp5133 := MakeNative(func(__e *ControlFlow) {
			W2943 := __e.Get(1)
			_ = W2943
			tmp5135 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2943)

			if True == tmp5135 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2943)
				return
			}

		}, 1)

		tmp5141 := Call(__e, PrimFunc(symshen_4hds_a_2), V2942, MakeNumber(35))

		var ifres5136 Obj

		if True == tmp5141 {
			tmp5137 := MakeNative(func(__e *ControlFlow) {
				W2944 := __e.Get(1)
				_ = W2944
				__e.TailApply(PrimFunc(symshen_4comb), W2944, symshen_4skip)
				return
			}, 1)

			tmp5138 := Call(__e, PrimFunc(symtail), V2942)

			tmp5139 := Call(__e, tmp5137, tmp5138)

			ifres5136 = tmp5139

		} else {
			tmp5140 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5136 = tmp5140

		}

		__e.TailApply(tmp5133, ifres5136)
		return

	}, 1)

	tmp5142 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp5132)

	_ = tmp5142

	tmp5143 := MakeNative(func(__e *ControlFlow) {
		V2945 := __e.Get(1)
		_ = V2945
		tmp5144 := MakeNative(func(__e *ControlFlow) {
			W2946 := __e.Get(1)
			_ = W2946
			tmp5200 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2946)

			if True == tmp5200 {
				tmp5145 := MakeNative(func(__e *ControlFlow) {
					W2952 := __e.Get(1)
					_ = W2952
					tmp5183 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2952)

					if True == tmp5183 {
						tmp5146 := MakeNative(func(__e *ControlFlow) {
							W2958 := __e.Get(1)
							_ = W2958
							tmp5172 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2958)

							if True == tmp5172 {
								tmp5147 := MakeNative(func(__e *ControlFlow) {
									W2962 := __e.Get(1)
									_ = W2962
									tmp5161 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2962)

									if True == tmp5161 {
										tmp5148 := MakeNative(func(__e *ControlFlow) {
											W2966 := __e.Get(1)
											_ = W2966
											tmp5150 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2966)

											if True == tmp5150 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												__e.Return(W2966)
												return
											}

										}, 1)

										tmp5151 := MakeNative(func(__e *ControlFlow) {
											W2967 := __e.Get(1)
											_ = W2967
											tmp5157 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2967)

											if True == tmp5157 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp5152 := MakeNative(func(__e *ControlFlow) {
													W2968 := __e.Get(1)
													_ = W2968
													tmp5153 := MakeNative(func(__e *ControlFlow) {
														W2969 := __e.Get(1)
														_ = W2969
														__e.TailApply(PrimFunc(symshen_4comb), W2969, W2968)
														return
													}, 1)

													tmp5154 := Call(__e, PrimFunc(symshen_4in_1_6), W2967)

													__e.TailApply(tmp5153, tmp5154)
													return

												}, 1)

												tmp5155 := Call(__e, PrimFunc(symshen_4_5_1out), W2967)

												__e.TailApply(tmp5152, tmp5155)
												return

											}

										}, 1)

										tmp5158 := Call(__e, PrimFunc(symshen_4_5integer_6), V2945)

										tmp5159 := Call(__e, tmp5151, tmp5158)

										__e.TailApply(tmp5148, tmp5159)
										return

									} else {
										__e.Return(W2962)
										return
									}

								}, 1)

								tmp5162 := MakeNative(func(__e *ControlFlow) {
									W2963 := __e.Get(1)
									_ = W2963
									tmp5168 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2963)

									if True == tmp5168 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5163 := MakeNative(func(__e *ControlFlow) {
											W2964 := __e.Get(1)
											_ = W2964
											tmp5164 := MakeNative(func(__e *ControlFlow) {
												W2965 := __e.Get(1)
												_ = W2965
												__e.TailApply(PrimFunc(symshen_4comb), W2965, W2964)
												return
											}, 1)

											tmp5165 := Call(__e, PrimFunc(symshen_4in_1_6), W2963)

											__e.TailApply(tmp5164, tmp5165)
											return

										}, 1)

										tmp5166 := Call(__e, PrimFunc(symshen_4_5_1out), W2963)

										__e.TailApply(tmp5163, tmp5166)
										return

									}

								}, 1)

								tmp5169 := Call(__e, PrimFunc(symshen_4_5float_6), V2945)

								tmp5170 := Call(__e, tmp5162, tmp5169)

								__e.TailApply(tmp5147, tmp5170)
								return

							} else {
								__e.Return(W2958)
								return
							}

						}, 1)

						tmp5173 := MakeNative(func(__e *ControlFlow) {
							W2959 := __e.Get(1)
							_ = W2959
							tmp5179 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2959)

							if True == tmp5179 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5174 := MakeNative(func(__e *ControlFlow) {
									W2960 := __e.Get(1)
									_ = W2960
									tmp5175 := MakeNative(func(__e *ControlFlow) {
										W2961 := __e.Get(1)
										_ = W2961
										__e.TailApply(PrimFunc(symshen_4comb), W2961, W2960)
										return
									}, 1)

									tmp5176 := Call(__e, PrimFunc(symshen_4in_1_6), W2959)

									__e.TailApply(tmp5175, tmp5176)
									return

								}, 1)

								tmp5177 := Call(__e, PrimFunc(symshen_4_5_1out), W2959)

								__e.TailApply(tmp5174, tmp5177)
								return

							}

						}, 1)

						tmp5180 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2945)

						tmp5181 := Call(__e, tmp5173, tmp5180)

						__e.TailApply(tmp5146, tmp5181)
						return

					} else {
						__e.Return(W2952)
						return
					}

				}, 1)

				tmp5184 := MakeNative(func(__e *ControlFlow) {
					W2953 := __e.Get(1)
					_ = W2953
					tmp5196 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2953)

					if True == tmp5196 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5185 := MakeNative(func(__e *ControlFlow) {
							W2954 := __e.Get(1)
							_ = W2954
							tmp5186 := MakeNative(func(__e *ControlFlow) {
								W2955 := __e.Get(1)
								_ = W2955
								tmp5192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2955)

								if True == tmp5192 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp5187 := MakeNative(func(__e *ControlFlow) {
										W2956 := __e.Get(1)
										_ = W2956
										tmp5188 := MakeNative(func(__e *ControlFlow) {
											W2957 := __e.Get(1)
											_ = W2957
											__e.TailApply(PrimFunc(symshen_4comb), W2957, W2956)
											return
										}, 1)

										tmp5189 := Call(__e, PrimFunc(symshen_4in_1_6), W2955)

										__e.TailApply(tmp5188, tmp5189)
										return

									}, 1)

									tmp5190 := Call(__e, PrimFunc(symshen_4_5_1out), W2955)

									__e.TailApply(tmp5187, tmp5190)
									return

								}

							}, 1)

							tmp5193 := Call(__e, PrimFunc(symshen_4_5number_6), W2954)

							__e.TailApply(tmp5186, tmp5193)
							return

						}, 1)

						tmp5194 := Call(__e, PrimFunc(symshen_4in_1_6), W2953)

						__e.TailApply(tmp5185, tmp5194)
						return

					}

				}, 1)

				tmp5197 := Call(__e, PrimFunc(symshen_4_5plus_6), V2945)

				tmp5198 := Call(__e, tmp5184, tmp5197)

				__e.TailApply(tmp5145, tmp5198)
				return

			} else {
				__e.Return(W2946)
				return
			}

		}, 1)

		tmp5201 := MakeNative(func(__e *ControlFlow) {
			W2947 := __e.Get(1)
			_ = W2947
			tmp5214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2947)

			if True == tmp5214 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5202 := MakeNative(func(__e *ControlFlow) {
					W2948 := __e.Get(1)
					_ = W2948
					tmp5203 := MakeNative(func(__e *ControlFlow) {
						W2949 := __e.Get(1)
						_ = W2949
						tmp5210 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2949)

						if True == tmp5210 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp5204 := MakeNative(func(__e *ControlFlow) {
								W2950 := __e.Get(1)
								_ = W2950
								tmp5205 := MakeNative(func(__e *ControlFlow) {
									W2951 := __e.Get(1)
									_ = W2951
									tmp5206 := PrimNumberSubtract(MakeNumber(0), W2950)

									__e.TailApply(PrimFunc(symshen_4comb), W2951, tmp5206)
									return

								}, 1)

								tmp5207 := Call(__e, PrimFunc(symshen_4in_1_6), W2949)

								__e.TailApply(tmp5205, tmp5207)
								return

							}, 1)

							tmp5208 := Call(__e, PrimFunc(symshen_4_5_1out), W2949)

							__e.TailApply(tmp5204, tmp5208)
							return

						}

					}, 1)

					tmp5211 := Call(__e, PrimFunc(symshen_4_5number_6), W2948)

					__e.TailApply(tmp5203, tmp5211)
					return

				}, 1)

				tmp5212 := Call(__e, PrimFunc(symshen_4in_1_6), W2947)

				__e.TailApply(tmp5202, tmp5212)
				return

			}

		}, 1)

		tmp5215 := Call(__e, PrimFunc(symshen_4_5minus_6), V2945)

		tmp5216 := Call(__e, tmp5201, tmp5215)

		__e.TailApply(tmp5144, tmp5216)
		return

	}, 1)

	tmp5217 := Call(__e, ns2_1set, symshen_4_5number_6, tmp5143)

	_ = tmp5217

	tmp5218 := MakeNative(func(__e *ControlFlow) {
		V2970 := __e.Get(1)
		_ = V2970
		tmp5219 := MakeNative(func(__e *ControlFlow) {
			W2971 := __e.Get(1)
			_ = W2971
			tmp5221 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2971)

			if True == tmp5221 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2971)
				return
			}

		}, 1)

		tmp5227 := Call(__e, PrimFunc(symshen_4hds_a_2), V2970, MakeNumber(45))

		var ifres5222 Obj

		if True == tmp5227 {
			tmp5223 := MakeNative(func(__e *ControlFlow) {
				W2972 := __e.Get(1)
				_ = W2972
				__e.TailApply(PrimFunc(symshen_4comb), W2972, symshen_4skip)
				return
			}, 1)

			tmp5224 := Call(__e, PrimFunc(symtail), V2970)

			tmp5225 := Call(__e, tmp5223, tmp5224)

			ifres5222 = tmp5225

		} else {
			tmp5226 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5222 = tmp5226

		}

		__e.TailApply(tmp5219, ifres5222)
		return

	}, 1)

	tmp5228 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp5218)

	_ = tmp5228

	tmp5229 := MakeNative(func(__e *ControlFlow) {
		V2973 := __e.Get(1)
		_ = V2973
		tmp5230 := MakeNative(func(__e *ControlFlow) {
			W2974 := __e.Get(1)
			_ = W2974
			tmp5232 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2974)

			if True == tmp5232 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2974)
				return
			}

		}, 1)

		tmp5238 := Call(__e, PrimFunc(symshen_4hds_a_2), V2973, MakeNumber(43))

		var ifres5233 Obj

		if True == tmp5238 {
			tmp5234 := MakeNative(func(__e *ControlFlow) {
				W2975 := __e.Get(1)
				_ = W2975
				__e.TailApply(PrimFunc(symshen_4comb), W2975, symshen_4skip)
				return
			}, 1)

			tmp5235 := Call(__e, PrimFunc(symtail), V2973)

			tmp5236 := Call(__e, tmp5234, tmp5235)

			ifres5233 = tmp5236

		} else {
			tmp5237 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5233 = tmp5237

		}

		__e.TailApply(tmp5230, ifres5233)
		return

	}, 1)

	tmp5239 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp5229)

	_ = tmp5239

	tmp5240 := MakeNative(func(__e *ControlFlow) {
		V2976 := __e.Get(1)
		_ = V2976
		tmp5241 := MakeNative(func(__e *ControlFlow) {
			W2977 := __e.Get(1)
			_ = W2977
			tmp5243 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2977)

			if True == tmp5243 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2977)
				return
			}

		}, 1)

		tmp5244 := MakeNative(func(__e *ControlFlow) {
			W2978 := __e.Get(1)
			_ = W2978
			tmp5251 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2978)

			if True == tmp5251 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5245 := MakeNative(func(__e *ControlFlow) {
					W2979 := __e.Get(1)
					_ = W2979
					tmp5246 := MakeNative(func(__e *ControlFlow) {
						W2980 := __e.Get(1)
						_ = W2980
						tmp5247 := Call(__e, PrimFunc(symshen_4compute_1integer), W2979)

						__e.TailApply(PrimFunc(symshen_4comb), W2980, tmp5247)
						return

					}, 1)

					tmp5248 := Call(__e, PrimFunc(symshen_4in_1_6), W2978)

					__e.TailApply(tmp5246, tmp5248)
					return

				}, 1)

				tmp5249 := Call(__e, PrimFunc(symshen_4_5_1out), W2978)

				__e.TailApply(tmp5245, tmp5249)
				return

			}

		}, 1)

		tmp5252 := Call(__e, PrimFunc(symshen_4_5digits_6), V2976)

		tmp5253 := Call(__e, tmp5244, tmp5252)

		__e.TailApply(tmp5241, tmp5253)
		return

	}, 1)

	tmp5254 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp5240)

	_ = tmp5254

	tmp5255 := MakeNative(func(__e *ControlFlow) {
		V2981 := __e.Get(1)
		_ = V2981
		tmp5256 := MakeNative(func(__e *ControlFlow) {
			W2982 := __e.Get(1)
			_ = W2982
			tmp5271 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2982)

			if True == tmp5271 {
				tmp5257 := MakeNative(func(__e *ControlFlow) {
					W2989 := __e.Get(1)
					_ = W2989
					tmp5259 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2989)

					if True == tmp5259 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W2989)
						return
					}

				}, 1)

				tmp5260 := MakeNative(func(__e *ControlFlow) {
					W2990 := __e.Get(1)
					_ = W2990
					tmp5267 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2990)

					if True == tmp5267 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5261 := MakeNative(func(__e *ControlFlow) {
							W2991 := __e.Get(1)
							_ = W2991
							tmp5262 := MakeNative(func(__e *ControlFlow) {
								W2992 := __e.Get(1)
								_ = W2992
								tmp5263 := PrimCons(W2991, Nil)

								__e.TailApply(PrimFunc(symshen_4comb), W2992, tmp5263)
								return

							}, 1)

							tmp5264 := Call(__e, PrimFunc(symshen_4in_1_6), W2990)

							__e.TailApply(tmp5262, tmp5264)
							return

						}, 1)

						tmp5265 := Call(__e, PrimFunc(symshen_4_5_1out), W2990)

						__e.TailApply(tmp5261, tmp5265)
						return

					}

				}, 1)

				tmp5268 := Call(__e, PrimFunc(symshen_4_5digit_6), V2981)

				tmp5269 := Call(__e, tmp5260, tmp5268)

				__e.TailApply(tmp5257, tmp5269)
				return

			} else {
				__e.Return(W2982)
				return
			}

		}, 1)

		tmp5272 := MakeNative(func(__e *ControlFlow) {
			W2983 := __e.Get(1)
			_ = W2983
			tmp5287 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2983)

			if True == tmp5287 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5273 := MakeNative(func(__e *ControlFlow) {
					W2984 := __e.Get(1)
					_ = W2984
					tmp5274 := MakeNative(func(__e *ControlFlow) {
						W2985 := __e.Get(1)
						_ = W2985
						tmp5275 := MakeNative(func(__e *ControlFlow) {
							W2986 := __e.Get(1)
							_ = W2986
							tmp5282 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2986)

							if True == tmp5282 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5276 := MakeNative(func(__e *ControlFlow) {
									W2987 := __e.Get(1)
									_ = W2987
									tmp5277 := MakeNative(func(__e *ControlFlow) {
										W2988 := __e.Get(1)
										_ = W2988
										tmp5278 := PrimCons(W2984, W2987)

										__e.TailApply(PrimFunc(symshen_4comb), W2988, tmp5278)
										return

									}, 1)

									tmp5279 := Call(__e, PrimFunc(symshen_4in_1_6), W2986)

									__e.TailApply(tmp5277, tmp5279)
									return

								}, 1)

								tmp5280 := Call(__e, PrimFunc(symshen_4_5_1out), W2986)

								__e.TailApply(tmp5276, tmp5280)
								return

							}

						}, 1)

						tmp5283 := Call(__e, PrimFunc(symshen_4_5digits_6), W2985)

						__e.TailApply(tmp5275, tmp5283)
						return

					}, 1)

					tmp5284 := Call(__e, PrimFunc(symshen_4in_1_6), W2983)

					__e.TailApply(tmp5274, tmp5284)
					return

				}, 1)

				tmp5285 := Call(__e, PrimFunc(symshen_4_5_1out), W2983)

				__e.TailApply(tmp5273, tmp5285)
				return

			}

		}, 1)

		tmp5288 := Call(__e, PrimFunc(symshen_4_5digit_6), V2981)

		tmp5289 := Call(__e, tmp5272, tmp5288)

		__e.TailApply(tmp5256, tmp5289)
		return

	}, 1)

	tmp5290 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp5255)

	_ = tmp5290

	tmp5291 := MakeNative(func(__e *ControlFlow) {
		V2993 := __e.Get(1)
		_ = V2993
		tmp5292 := MakeNative(func(__e *ControlFlow) {
			W2994 := __e.Get(1)
			_ = W2994
			tmp5294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2994)

			if True == tmp5294 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W2994)
				return
			}

		}, 1)

		tmp5305 := PrimIsPair(V2993)

		var ifres5295 Obj

		if True == tmp5305 {
			tmp5296 := MakeNative(func(__e *ControlFlow) {
				W2995 := __e.Get(1)
				_ = W2995
				tmp5297 := MakeNative(func(__e *ControlFlow) {
					W2996 := __e.Get(1)
					_ = W2996
					tmp5300 := Call(__e, PrimFunc(symshen_4digit_2), W2995)

					if True == tmp5300 {
						tmp5298 := Call(__e, PrimFunc(symshen_4byte_1_6digit), W2995)

						__e.TailApply(PrimFunc(symshen_4comb), W2996, tmp5298)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5301 := Call(__e, PrimFunc(symtail), V2993)

				__e.TailApply(tmp5297, tmp5301)
				return

			}, 1)

			tmp5302 := Call(__e, PrimFunc(symhead), V2993)

			tmp5303 := Call(__e, tmp5296, tmp5302)

			ifres5295 = tmp5303

		} else {
			tmp5304 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5295 = tmp5304

		}

		__e.TailApply(tmp5292, ifres5295)
		return

	}, 1)

	tmp5306 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp5291)

	_ = tmp5306

	tmp5307 := MakeNative(func(__e *ControlFlow) {
		V2997 := __e.Get(1)
		_ = V2997
		__e.Return(PrimNumberSubtract(V2997, MakeNumber(48)))
		return
	}, 1)

	tmp5308 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp5307)

	_ = tmp5308

	tmp5309 := MakeNative(func(__e *ControlFlow) {
		V2998 := __e.Get(1)
		_ = V2998
		tmp5310 := Call(__e, PrimFunc(symreverse), V2998)

		__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp5310, MakeNumber(0))
		return

	}, 1)

	tmp5311 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp5309)

	_ = tmp5311

	tmp5312 := MakeNative(func(__e *ControlFlow) {
		V3001 := __e.Get(1)
		_ = V3001
		V3002 := __e.Get(2)
		_ = V3002
		tmp5322 := PrimEqual(Nil, V3001)

		if True == tmp5322 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp5320 := PrimIsPair(V3001)

			if True == tmp5320 {
				tmp5313 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V3002)

				tmp5314 := PrimHead(V3001)

				tmp5315 := PrimNumberMultiply(tmp5313, tmp5314)

				tmp5316 := PrimTail(V3001)

				tmp5317 := PrimNumberAdd(V3002, MakeNumber(1))

				tmp5318 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp5316, tmp5317)

				__e.Return(PrimNumberAdd(tmp5315, tmp5318))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1integer_1h)
				return
			}

		}

	}, 2)

	tmp5323 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp5312)

	_ = tmp5323

	tmp5324 := MakeNative(func(__e *ControlFlow) {
		V3005 := __e.Get(1)
		_ = V3005
		V3006 := __e.Get(2)
		_ = V3006
		tmp5332 := PrimEqual(MakeNumber(0), V3006)

		if True == tmp5332 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp5330 := PrimGreatThan(V3006, MakeNumber(0))

			if True == tmp5330 {
				tmp5325 := PrimNumberSubtract(V3006, MakeNumber(1))

				tmp5326 := Call(__e, PrimFunc(symshen_4expt), V3005, tmp5325)

				__e.Return(PrimNumberMultiply(V3005, tmp5326))
				return

			} else {
				tmp5327 := PrimNumberAdd(V3006, MakeNumber(1))

				tmp5328 := Call(__e, PrimFunc(symshen_4expt), V3005, tmp5327)

				__e.Return(PrimNumberDivide(tmp5328, V3005))
				return

			}

		}

	}, 2)

	tmp5333 := Call(__e, ns2_1set, symshen_4expt, tmp5324)

	_ = tmp5333

	tmp5334 := MakeNative(func(__e *ControlFlow) {
		V3007 := __e.Get(1)
		_ = V3007
		tmp5335 := MakeNative(func(__e *ControlFlow) {
			W3008 := __e.Get(1)
			_ = W3008
			tmp5355 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3008)

			if True == tmp5355 {
				tmp5336 := MakeNative(func(__e *ControlFlow) {
					W3017 := __e.Get(1)
					_ = W3017
					tmp5338 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3017)

					if True == tmp5338 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3017)
						return
					}

				}, 1)

				tmp5339 := MakeNative(func(__e *ControlFlow) {
					W3018 := __e.Get(1)
					_ = W3018
					tmp5351 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3018)

					if True == tmp5351 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5340 := MakeNative(func(__e *ControlFlow) {
							W3019 := __e.Get(1)
							_ = W3019
							tmp5341 := MakeNative(func(__e *ControlFlow) {
								W3020 := __e.Get(1)
								_ = W3020
								tmp5347 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3020)

								if True == tmp5347 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp5342 := MakeNative(func(__e *ControlFlow) {
										W3021 := __e.Get(1)
										_ = W3021
										tmp5343 := MakeNative(func(__e *ControlFlow) {
											W3022 := __e.Get(1)
											_ = W3022
											__e.TailApply(PrimFunc(symshen_4comb), W3022, W3021)
											return
										}, 1)

										tmp5344 := Call(__e, PrimFunc(symshen_4in_1_6), W3020)

										__e.TailApply(tmp5343, tmp5344)
										return

									}, 1)

									tmp5345 := Call(__e, PrimFunc(symshen_4_5_1out), W3020)

									__e.TailApply(tmp5342, tmp5345)
									return

								}

							}, 1)

							tmp5348 := Call(__e, PrimFunc(symshen_4_5fraction_6), W3019)

							__e.TailApply(tmp5341, tmp5348)
							return

						}, 1)

						tmp5349 := Call(__e, PrimFunc(symshen_4in_1_6), W3018)

						__e.TailApply(tmp5340, tmp5349)
						return

					}

				}, 1)

				tmp5352 := Call(__e, PrimFunc(symshen_4_5stop_6), V3007)

				tmp5353 := Call(__e, tmp5339, tmp5352)

				__e.TailApply(tmp5336, tmp5353)
				return

			} else {
				__e.Return(W3008)
				return
			}

		}, 1)

		tmp5356 := MakeNative(func(__e *ControlFlow) {
			W3009 := __e.Get(1)
			_ = W3009
			tmp5377 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3009)

			if True == tmp5377 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5357 := MakeNative(func(__e *ControlFlow) {
					W3010 := __e.Get(1)
					_ = W3010
					tmp5358 := MakeNative(func(__e *ControlFlow) {
						W3011 := __e.Get(1)
						_ = W3011
						tmp5359 := MakeNative(func(__e *ControlFlow) {
							W3012 := __e.Get(1)
							_ = W3012
							tmp5372 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3012)

							if True == tmp5372 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5360 := MakeNative(func(__e *ControlFlow) {
									W3013 := __e.Get(1)
									_ = W3013
									tmp5361 := MakeNative(func(__e *ControlFlow) {
										W3014 := __e.Get(1)
										_ = W3014
										tmp5368 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3014)

										if True == tmp5368 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp5362 := MakeNative(func(__e *ControlFlow) {
												W3015 := __e.Get(1)
												_ = W3015
												tmp5363 := MakeNative(func(__e *ControlFlow) {
													W3016 := __e.Get(1)
													_ = W3016
													tmp5364 := PrimNumberAdd(W3010, W3015)

													__e.TailApply(PrimFunc(symshen_4comb), W3016, tmp5364)
													return

												}, 1)

												tmp5365 := Call(__e, PrimFunc(symshen_4in_1_6), W3014)

												__e.TailApply(tmp5363, tmp5365)
												return

											}, 1)

											tmp5366 := Call(__e, PrimFunc(symshen_4_5_1out), W3014)

											__e.TailApply(tmp5362, tmp5366)
											return

										}

									}, 1)

									tmp5369 := Call(__e, PrimFunc(symshen_4_5fraction_6), W3013)

									__e.TailApply(tmp5361, tmp5369)
									return

								}, 1)

								tmp5370 := Call(__e, PrimFunc(symshen_4in_1_6), W3012)

								__e.TailApply(tmp5360, tmp5370)
								return

							}

						}, 1)

						tmp5373 := Call(__e, PrimFunc(symshen_4_5stop_6), W3011)

						__e.TailApply(tmp5359, tmp5373)
						return

					}, 1)

					tmp5374 := Call(__e, PrimFunc(symshen_4in_1_6), W3009)

					__e.TailApply(tmp5358, tmp5374)
					return

				}, 1)

				tmp5375 := Call(__e, PrimFunc(symshen_4_5_1out), W3009)

				__e.TailApply(tmp5357, tmp5375)
				return

			}

		}, 1)

		tmp5378 := Call(__e, PrimFunc(symshen_4_5integer_6), V3007)

		tmp5379 := Call(__e, tmp5356, tmp5378)

		__e.TailApply(tmp5335, tmp5379)
		return

	}, 1)

	tmp5380 := Call(__e, ns2_1set, symshen_4_5float_6, tmp5334)

	_ = tmp5380

	tmp5381 := MakeNative(func(__e *ControlFlow) {
		V3023 := __e.Get(1)
		_ = V3023
		tmp5382 := MakeNative(func(__e *ControlFlow) {
			W3024 := __e.Get(1)
			_ = W3024
			tmp5384 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3024)

			if True == tmp5384 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3024)
				return
			}

		}, 1)

		tmp5390 := Call(__e, PrimFunc(symshen_4hds_a_2), V3023, MakeNumber(46))

		var ifres5385 Obj

		if True == tmp5390 {
			tmp5386 := MakeNative(func(__e *ControlFlow) {
				W3025 := __e.Get(1)
				_ = W3025
				__e.TailApply(PrimFunc(symshen_4comb), W3025, symshen_4skip)
				return
			}, 1)

			tmp5387 := Call(__e, PrimFunc(symtail), V3023)

			tmp5388 := Call(__e, tmp5386, tmp5387)

			ifres5385 = tmp5388

		} else {
			tmp5389 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5385 = tmp5389

		}

		__e.TailApply(tmp5382, ifres5385)
		return

	}, 1)

	tmp5391 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp5381)

	_ = tmp5391

	tmp5392 := MakeNative(func(__e *ControlFlow) {
		V3026 := __e.Get(1)
		_ = V3026
		tmp5393 := MakeNative(func(__e *ControlFlow) {
			W3027 := __e.Get(1)
			_ = W3027
			tmp5395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3027)

			if True == tmp5395 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3027)
				return
			}

		}, 1)

		tmp5396 := MakeNative(func(__e *ControlFlow) {
			W3028 := __e.Get(1)
			_ = W3028
			tmp5403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3028)

			if True == tmp5403 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5397 := MakeNative(func(__e *ControlFlow) {
					W3029 := __e.Get(1)
					_ = W3029
					tmp5398 := MakeNative(func(__e *ControlFlow) {
						W3030 := __e.Get(1)
						_ = W3030
						tmp5399 := Call(__e, PrimFunc(symshen_4compute_1fraction), W3029)

						__e.TailApply(PrimFunc(symshen_4comb), W3030, tmp5399)
						return

					}, 1)

					tmp5400 := Call(__e, PrimFunc(symshen_4in_1_6), W3028)

					__e.TailApply(tmp5398, tmp5400)
					return

				}, 1)

				tmp5401 := Call(__e, PrimFunc(symshen_4_5_1out), W3028)

				__e.TailApply(tmp5397, tmp5401)
				return

			}

		}, 1)

		tmp5404 := Call(__e, PrimFunc(symshen_4_5digits_6), V3026)

		tmp5405 := Call(__e, tmp5396, tmp5404)

		__e.TailApply(tmp5393, tmp5405)
		return

	}, 1)

	tmp5406 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp5392)

	_ = tmp5406

	tmp5407 := MakeNative(func(__e *ControlFlow) {
		V3031 := __e.Get(1)
		_ = V3031
		__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V3031, MakeNumber(-1))
		return
	}, 1)

	tmp5408 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp5407)

	_ = tmp5408

	tmp5409 := MakeNative(func(__e *ControlFlow) {
		V3034 := __e.Get(1)
		_ = V3034
		V3035 := __e.Get(2)
		_ = V3035
		tmp5419 := PrimEqual(Nil, V3034)

		if True == tmp5419 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp5417 := PrimIsPair(V3034)

			if True == tmp5417 {
				tmp5410 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V3035)

				tmp5411 := PrimHead(V3034)

				tmp5412 := PrimNumberMultiply(tmp5410, tmp5411)

				tmp5413 := PrimTail(V3034)

				tmp5414 := PrimNumberSubtract(V3035, MakeNumber(1))

				tmp5415 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp5413, tmp5414)

				__e.Return(PrimNumberAdd(tmp5412, tmp5415))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1fraction_1h)
				return
			}

		}

	}, 2)

	tmp5420 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp5409)

	_ = tmp5420

	tmp5421 := MakeNative(func(__e *ControlFlow) {
		V3036 := __e.Get(1)
		_ = V3036
		tmp5422 := MakeNative(func(__e *ControlFlow) {
			W3037 := __e.Get(1)
			_ = W3037
			tmp5451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3037)

			if True == tmp5451 {
				tmp5423 := MakeNative(func(__e *ControlFlow) {
					W3046 := __e.Get(1)
					_ = W3046
					tmp5425 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3046)

					if True == tmp5425 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3046)
						return
					}

				}, 1)

				tmp5426 := MakeNative(func(__e *ControlFlow) {
					W3047 := __e.Get(1)
					_ = W3047
					tmp5447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3047)

					if True == tmp5447 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5427 := MakeNative(func(__e *ControlFlow) {
							W3048 := __e.Get(1)
							_ = W3048
							tmp5428 := MakeNative(func(__e *ControlFlow) {
								W3049 := __e.Get(1)
								_ = W3049
								tmp5429 := MakeNative(func(__e *ControlFlow) {
									W3050 := __e.Get(1)
									_ = W3050
									tmp5442 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3050)

									if True == tmp5442 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5430 := MakeNative(func(__e *ControlFlow) {
											W3051 := __e.Get(1)
											_ = W3051
											tmp5431 := MakeNative(func(__e *ControlFlow) {
												W3052 := __e.Get(1)
												_ = W3052
												tmp5438 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3052)

												if True == tmp5438 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp5432 := MakeNative(func(__e *ControlFlow) {
														W3053 := __e.Get(1)
														_ = W3053
														tmp5433 := MakeNative(func(__e *ControlFlow) {
															W3054 := __e.Get(1)
															_ = W3054
															tmp5434 := Call(__e, PrimFunc(symshen_4compute_1E), W3048, W3053)

															__e.TailApply(PrimFunc(symshen_4comb), W3054, tmp5434)
															return

														}, 1)

														tmp5435 := Call(__e, PrimFunc(symshen_4in_1_6), W3052)

														__e.TailApply(tmp5433, tmp5435)
														return

													}, 1)

													tmp5436 := Call(__e, PrimFunc(symshen_4_5_1out), W3052)

													__e.TailApply(tmp5432, tmp5436)
													return

												}

											}, 1)

											tmp5439 := Call(__e, PrimFunc(symshen_4_5log10_6), W3051)

											__e.TailApply(tmp5431, tmp5439)
											return

										}, 1)

										tmp5440 := Call(__e, PrimFunc(symshen_4in_1_6), W3050)

										__e.TailApply(tmp5430, tmp5440)
										return

									}

								}, 1)

								tmp5443 := Call(__e, PrimFunc(symshen_4_5lowE_6), W3049)

								__e.TailApply(tmp5429, tmp5443)
								return

							}, 1)

							tmp5444 := Call(__e, PrimFunc(symshen_4in_1_6), W3047)

							__e.TailApply(tmp5428, tmp5444)
							return

						}, 1)

						tmp5445 := Call(__e, PrimFunc(symshen_4_5_1out), W3047)

						__e.TailApply(tmp5427, tmp5445)
						return

					}

				}, 1)

				tmp5448 := Call(__e, PrimFunc(symshen_4_5integer_6), V3036)

				tmp5449 := Call(__e, tmp5426, tmp5448)

				__e.TailApply(tmp5423, tmp5449)
				return

			} else {
				__e.Return(W3037)
				return
			}

		}, 1)

		tmp5452 := MakeNative(func(__e *ControlFlow) {
			W3038 := __e.Get(1)
			_ = W3038
			tmp5473 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3038)

			if True == tmp5473 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5453 := MakeNative(func(__e *ControlFlow) {
					W3039 := __e.Get(1)
					_ = W3039
					tmp5454 := MakeNative(func(__e *ControlFlow) {
						W3040 := __e.Get(1)
						_ = W3040
						tmp5455 := MakeNative(func(__e *ControlFlow) {
							W3041 := __e.Get(1)
							_ = W3041
							tmp5468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3041)

							if True == tmp5468 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5456 := MakeNative(func(__e *ControlFlow) {
									W3042 := __e.Get(1)
									_ = W3042
									tmp5457 := MakeNative(func(__e *ControlFlow) {
										W3043 := __e.Get(1)
										_ = W3043
										tmp5464 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3043)

										if True == tmp5464 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp5458 := MakeNative(func(__e *ControlFlow) {
												W3044 := __e.Get(1)
												_ = W3044
												tmp5459 := MakeNative(func(__e *ControlFlow) {
													W3045 := __e.Get(1)
													_ = W3045
													tmp5460 := Call(__e, PrimFunc(symshen_4compute_1E), W3039, W3044)

													__e.TailApply(PrimFunc(symshen_4comb), W3045, tmp5460)
													return

												}, 1)

												tmp5461 := Call(__e, PrimFunc(symshen_4in_1_6), W3043)

												__e.TailApply(tmp5459, tmp5461)
												return

											}, 1)

											tmp5462 := Call(__e, PrimFunc(symshen_4_5_1out), W3043)

											__e.TailApply(tmp5458, tmp5462)
											return

										}

									}, 1)

									tmp5465 := Call(__e, PrimFunc(symshen_4_5log10_6), W3042)

									__e.TailApply(tmp5457, tmp5465)
									return

								}, 1)

								tmp5466 := Call(__e, PrimFunc(symshen_4in_1_6), W3041)

								__e.TailApply(tmp5456, tmp5466)
								return

							}

						}, 1)

						tmp5469 := Call(__e, PrimFunc(symshen_4_5lowE_6), W3040)

						__e.TailApply(tmp5455, tmp5469)
						return

					}, 1)

					tmp5470 := Call(__e, PrimFunc(symshen_4in_1_6), W3038)

					__e.TailApply(tmp5454, tmp5470)
					return

				}, 1)

				tmp5471 := Call(__e, PrimFunc(symshen_4_5_1out), W3038)

				__e.TailApply(tmp5453, tmp5471)
				return

			}

		}, 1)

		tmp5474 := Call(__e, PrimFunc(symshen_4_5float_6), V3036)

		tmp5475 := Call(__e, tmp5452, tmp5474)

		__e.TailApply(tmp5422, tmp5475)
		return

	}, 1)

	tmp5476 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp5421)

	_ = tmp5476

	tmp5477 := MakeNative(func(__e *ControlFlow) {
		V3055 := __e.Get(1)
		_ = V3055
		tmp5478 := MakeNative(func(__e *ControlFlow) {
			W3056 := __e.Get(1)
			_ = W3056
			tmp5511 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3056)

			if True == tmp5511 {
				tmp5479 := MakeNative(func(__e *ControlFlow) {
					W3062 := __e.Get(1)
					_ = W3062
					tmp5493 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3062)

					if True == tmp5493 {
						tmp5480 := MakeNative(func(__e *ControlFlow) {
							W3068 := __e.Get(1)
							_ = W3068
							tmp5482 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3068)

							if True == tmp5482 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W3068)
								return
							}

						}, 1)

						tmp5483 := MakeNative(func(__e *ControlFlow) {
							W3069 := __e.Get(1)
							_ = W3069
							tmp5489 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3069)

							if True == tmp5489 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5484 := MakeNative(func(__e *ControlFlow) {
									W3070 := __e.Get(1)
									_ = W3070
									tmp5485 := MakeNative(func(__e *ControlFlow) {
										W3071 := __e.Get(1)
										_ = W3071
										__e.TailApply(PrimFunc(symshen_4comb), W3071, W3070)
										return
									}, 1)

									tmp5486 := Call(__e, PrimFunc(symshen_4in_1_6), W3069)

									__e.TailApply(tmp5485, tmp5486)
									return

								}, 1)

								tmp5487 := Call(__e, PrimFunc(symshen_4_5_1out), W3069)

								__e.TailApply(tmp5484, tmp5487)
								return

							}

						}, 1)

						tmp5490 := Call(__e, PrimFunc(symshen_4_5integer_6), V3055)

						tmp5491 := Call(__e, tmp5483, tmp5490)

						__e.TailApply(tmp5480, tmp5491)
						return

					} else {
						__e.Return(W3062)
						return
					}

				}, 1)

				tmp5494 := MakeNative(func(__e *ControlFlow) {
					W3063 := __e.Get(1)
					_ = W3063
					tmp5507 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3063)

					if True == tmp5507 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5495 := MakeNative(func(__e *ControlFlow) {
							W3064 := __e.Get(1)
							_ = W3064
							tmp5496 := MakeNative(func(__e *ControlFlow) {
								W3065 := __e.Get(1)
								_ = W3065
								tmp5503 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3065)

								if True == tmp5503 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp5497 := MakeNative(func(__e *ControlFlow) {
										W3066 := __e.Get(1)
										_ = W3066
										tmp5498 := MakeNative(func(__e *ControlFlow) {
											W3067 := __e.Get(1)
											_ = W3067
											tmp5499 := PrimNumberSubtract(MakeNumber(0), W3066)

											__e.TailApply(PrimFunc(symshen_4comb), W3067, tmp5499)
											return

										}, 1)

										tmp5500 := Call(__e, PrimFunc(symshen_4in_1_6), W3065)

										__e.TailApply(tmp5498, tmp5500)
										return

									}, 1)

									tmp5501 := Call(__e, PrimFunc(symshen_4_5_1out), W3065)

									__e.TailApply(tmp5497, tmp5501)
									return

								}

							}, 1)

							tmp5504 := Call(__e, PrimFunc(symshen_4_5log10_6), W3064)

							__e.TailApply(tmp5496, tmp5504)
							return

						}, 1)

						tmp5505 := Call(__e, PrimFunc(symshen_4in_1_6), W3063)

						__e.TailApply(tmp5495, tmp5505)
						return

					}

				}, 1)

				tmp5508 := Call(__e, PrimFunc(symshen_4_5minus_6), V3055)

				tmp5509 := Call(__e, tmp5494, tmp5508)

				__e.TailApply(tmp5479, tmp5509)
				return

			} else {
				__e.Return(W3056)
				return
			}

		}, 1)

		tmp5512 := MakeNative(func(__e *ControlFlow) {
			W3057 := __e.Get(1)
			_ = W3057
			tmp5524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3057)

			if True == tmp5524 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5513 := MakeNative(func(__e *ControlFlow) {
					W3058 := __e.Get(1)
					_ = W3058
					tmp5514 := MakeNative(func(__e *ControlFlow) {
						W3059 := __e.Get(1)
						_ = W3059
						tmp5520 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3059)

						if True == tmp5520 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp5515 := MakeNative(func(__e *ControlFlow) {
								W3060 := __e.Get(1)
								_ = W3060
								tmp5516 := MakeNative(func(__e *ControlFlow) {
									W3061 := __e.Get(1)
									_ = W3061
									__e.TailApply(PrimFunc(symshen_4comb), W3061, W3060)
									return
								}, 1)

								tmp5517 := Call(__e, PrimFunc(symshen_4in_1_6), W3059)

								__e.TailApply(tmp5516, tmp5517)
								return

							}, 1)

							tmp5518 := Call(__e, PrimFunc(symshen_4_5_1out), W3059)

							__e.TailApply(tmp5515, tmp5518)
							return

						}

					}, 1)

					tmp5521 := Call(__e, PrimFunc(symshen_4_5log10_6), W3058)

					__e.TailApply(tmp5514, tmp5521)
					return

				}, 1)

				tmp5522 := Call(__e, PrimFunc(symshen_4in_1_6), W3057)

				__e.TailApply(tmp5513, tmp5522)
				return

			}

		}, 1)

		tmp5525 := Call(__e, PrimFunc(symshen_4_5plus_6), V3055)

		tmp5526 := Call(__e, tmp5512, tmp5525)

		__e.TailApply(tmp5478, tmp5526)
		return

	}, 1)

	tmp5527 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp5477)

	_ = tmp5527

	tmp5528 := MakeNative(func(__e *ControlFlow) {
		V3072 := __e.Get(1)
		_ = V3072
		tmp5529 := MakeNative(func(__e *ControlFlow) {
			W3073 := __e.Get(1)
			_ = W3073
			tmp5531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3073)

			if True == tmp5531 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3073)
				return
			}

		}, 1)

		tmp5537 := Call(__e, PrimFunc(symshen_4hds_a_2), V3072, MakeNumber(101))

		var ifres5532 Obj

		if True == tmp5537 {
			tmp5533 := MakeNative(func(__e *ControlFlow) {
				W3074 := __e.Get(1)
				_ = W3074
				__e.TailApply(PrimFunc(symshen_4comb), W3074, symshen_4skip)
				return
			}, 1)

			tmp5534 := Call(__e, PrimFunc(symtail), V3072)

			tmp5535 := Call(__e, tmp5533, tmp5534)

			ifres5532 = tmp5535

		} else {
			tmp5536 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5532 = tmp5536

		}

		__e.TailApply(tmp5529, ifres5532)
		return

	}, 1)

	tmp5538 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp5528)

	_ = tmp5538

	tmp5539 := MakeNative(func(__e *ControlFlow) {
		V3075 := __e.Get(1)
		_ = V3075
		V3076 := __e.Get(2)
		_ = V3076
		tmp5540 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V3076)

		__e.Return(PrimNumberMultiply(V3075, tmp5540))
		return

	}, 2)

	tmp5541 := Call(__e, ns2_1set, symshen_4compute_1E, tmp5539)

	_ = tmp5541

	tmp5542 := MakeNative(func(__e *ControlFlow) {
		V3077 := __e.Get(1)
		_ = V3077
		tmp5543 := MakeNative(func(__e *ControlFlow) {
			W3078 := __e.Get(1)
			_ = W3078
			tmp5555 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3078)

			if True == tmp5555 {
				tmp5544 := MakeNative(func(__e *ControlFlow) {
					W3083 := __e.Get(1)
					_ = W3083
					tmp5546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3083)

					if True == tmp5546 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3083)
						return
					}

				}, 1)

				tmp5547 := MakeNative(func(__e *ControlFlow) {
					W3084 := __e.Get(1)
					_ = W3084
					tmp5551 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3084)

					if True == tmp5551 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5548 := MakeNative(func(__e *ControlFlow) {
							W3085 := __e.Get(1)
							_ = W3085
							__e.TailApply(PrimFunc(symshen_4comb), W3085, symshen_4skip)
							return
						}, 1)

						tmp5549 := Call(__e, PrimFunc(symshen_4in_1_6), W3084)

						__e.TailApply(tmp5548, tmp5549)
						return

					}

				}, 1)

				tmp5552 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3077)

				tmp5553 := Call(__e, tmp5547, tmp5552)

				__e.TailApply(tmp5544, tmp5553)
				return

			} else {
				__e.Return(W3078)
				return
			}

		}, 1)

		tmp5556 := MakeNative(func(__e *ControlFlow) {
			W3079 := __e.Get(1)
			_ = W3079
			tmp5566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3079)

			if True == tmp5566 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5557 := MakeNative(func(__e *ControlFlow) {
					W3080 := __e.Get(1)
					_ = W3080
					tmp5558 := MakeNative(func(__e *ControlFlow) {
						W3081 := __e.Get(1)
						_ = W3081
						tmp5562 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3081)

						if True == tmp5562 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp5559 := MakeNative(func(__e *ControlFlow) {
								W3082 := __e.Get(1)
								_ = W3082
								__e.TailApply(PrimFunc(symshen_4comb), W3082, symshen_4skip)
								return
							}, 1)

							tmp5560 := Call(__e, PrimFunc(symshen_4in_1_6), W3081)

							__e.TailApply(tmp5559, tmp5560)
							return

						}

					}, 1)

					tmp5563 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), W3080)

					__e.TailApply(tmp5558, tmp5563)
					return

				}, 1)

				tmp5564 := Call(__e, PrimFunc(symshen_4in_1_6), W3079)

				__e.TailApply(tmp5557, tmp5564)
				return

			}

		}, 1)

		tmp5567 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3077)

		tmp5568 := Call(__e, tmp5556, tmp5567)

		__e.TailApply(tmp5543, tmp5568)
		return

	}, 1)

	tmp5569 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp5542)

	_ = tmp5569

	tmp5570 := MakeNative(func(__e *ControlFlow) {
		V3086 := __e.Get(1)
		_ = V3086
		tmp5571 := MakeNative(func(__e *ControlFlow) {
			W3087 := __e.Get(1)
			_ = W3087
			tmp5573 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3087)

			if True == tmp5573 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3087)
				return
			}

		}, 1)

		tmp5583 := PrimIsPair(V3086)

		var ifres5574 Obj

		if True == tmp5583 {
			tmp5575 := MakeNative(func(__e *ControlFlow) {
				W3088 := __e.Get(1)
				_ = W3088
				tmp5576 := MakeNative(func(__e *ControlFlow) {
					W3089 := __e.Get(1)
					_ = W3089
					tmp5578 := Call(__e, PrimFunc(symshen_4whitespace_2), W3088)

					if True == tmp5578 {
						__e.TailApply(PrimFunc(symshen_4comb), W3089, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5579 := Call(__e, PrimFunc(symtail), V3086)

				__e.TailApply(tmp5576, tmp5579)
				return

			}, 1)

			tmp5580 := Call(__e, PrimFunc(symhead), V3086)

			tmp5581 := Call(__e, tmp5575, tmp5580)

			ifres5574 = tmp5581

		} else {
			tmp5582 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5574 = tmp5582

		}

		__e.TailApply(tmp5571, ifres5574)
		return

	}, 1)

	tmp5584 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp5570)

	_ = tmp5584

	tmp5585 := MakeNative(func(__e *ControlFlow) {
		V3092 := __e.Get(1)
		_ = V3092
		tmp5593 := PrimEqual(MakeNumber(32), V3092)

		if True == tmp5593 {
			__e.Return(True)
			return
		} else {
			tmp5591 := PrimEqual(MakeNumber(13), V3092)

			if True == tmp5591 {
				__e.Return(True)
				return
			} else {
				tmp5589 := PrimEqual(MakeNumber(10), V3092)

				if True == tmp5589 {
					__e.Return(True)
					return
				} else {
					tmp5587 := PrimEqual(MakeNumber(9), V3092)

					if True == tmp5587 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 1)

	tmp5594 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp5585)

	_ = tmp5594

	tmp5595 := MakeNative(func(__e *ControlFlow) {
		V3093 := __e.Get(1)
		_ = V3093
		tmp5618 := PrimEqual(Nil, V3093)

		if True == tmp5618 {
			__e.Return(Nil)
			return
		} else {
			tmp5616 := PrimIsPair(V3093)

			var ifres5612 Obj

			if True == tmp5616 {
				tmp5614 := PrimHead(V3093)

				tmp5615 := Call(__e, PrimFunc(symshen_4packaged_2), tmp5614)

				var ifres5613 Obj

				if True == tmp5615 {
					ifres5613 = True

				} else {
					ifres5613 = False

				}

				ifres5612 = ifres5613

			} else {
				ifres5612 = False

			}

			if True == ifres5612 {
				tmp5596 := PrimHead(V3093)

				tmp5597 := Call(__e, PrimFunc(symshen_4unpackage), tmp5596)

				tmp5598 := PrimTail(V3093)

				tmp5599 := Call(__e, PrimFunc(symappend), tmp5597, tmp5598)

				__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp5599)
				return

			} else {
				tmp5610 := PrimIsPair(V3093)

				if True == tmp5610 {
					tmp5600 := MakeNative(func(__e *ControlFlow) {
						W3094 := __e.Get(1)
						_ = W3094
						tmp5606 := Call(__e, PrimFunc(symshen_4packaged_2), W3094)

						if True == tmp5606 {
							tmp5601 := PrimTail(V3093)

							tmp5602 := PrimCons(W3094, tmp5601)

							__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp5602)
							return

						} else {
							tmp5603 := PrimTail(V3093)

							tmp5604 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp5603)

							__e.Return(PrimCons(W3094, tmp5604))
							return

						}

					}, 1)

					tmp5607 := PrimHead(V3093)

					tmp5608 := Call(__e, PrimFunc(symmacroexpand), tmp5607)

					__e.TailApply(tmp5600, tmp5608)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage_emacroexpand)
					return
				}

			}

		}

	}, 1)

	tmp5619 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp5595)

	_ = tmp5619

	tmp5620 := MakeNative(func(__e *ControlFlow) {
		V3097 := __e.Get(1)
		_ = V3097
		tmp5635 := PrimIsPair(V3097)

		var ifres5622 Obj

		if True == tmp5635 {
			tmp5633 := PrimHead(V3097)

			tmp5634 := PrimEqual(sympackage, tmp5633)

			var ifres5624 Obj

			if True == tmp5634 {
				tmp5631 := PrimTail(V3097)

				tmp5632 := PrimIsPair(tmp5631)

				var ifres5626 Obj

				if True == tmp5632 {
					tmp5628 := PrimTail(V3097)

					tmp5629 := PrimTail(tmp5628)

					tmp5630 := PrimIsPair(tmp5629)

					var ifres5627 Obj

					if True == tmp5630 {
						ifres5627 = True

					} else {
						ifres5627 = False

					}

					ifres5626 = ifres5627

				} else {
					ifres5626 = False

				}

				var ifres5625 Obj

				if True == ifres5626 {
					ifres5625 = True

				} else {
					ifres5625 = False

				}

				ifres5624 = ifres5625

			} else {
				ifres5624 = False

			}

			var ifres5623 Obj

			if True == ifres5624 {
				ifres5623 = True

			} else {
				ifres5623 = False

			}

			ifres5622 = ifres5623

		} else {
			ifres5622 = False

		}

		if True == ifres5622 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp5636 := Call(__e, ns2_1set, symshen_4packaged_2, tmp5620)

	_ = tmp5636

	tmp5637 := MakeNative(func(__e *ControlFlow) {
		V3100 := __e.Get(1)
		_ = V3100
		tmp5698 := PrimIsPair(V3100)

		var ifres5680 Obj

		if True == tmp5698 {
			tmp5696 := PrimHead(V3100)

			tmp5697 := PrimEqual(sympackage, tmp5696)

			var ifres5682 Obj

			if True == tmp5697 {
				tmp5694 := PrimTail(V3100)

				tmp5695 := PrimIsPair(tmp5694)

				var ifres5684 Obj

				if True == tmp5695 {
					tmp5691 := PrimTail(V3100)

					tmp5692 := PrimHead(tmp5691)

					tmp5693 := PrimEqual(symnull, tmp5692)

					var ifres5686 Obj

					if True == tmp5693 {
						tmp5688 := PrimTail(V3100)

						tmp5689 := PrimTail(tmp5688)

						tmp5690 := PrimIsPair(tmp5689)

						var ifres5687 Obj

						if True == tmp5690 {
							ifres5687 = True

						} else {
							ifres5687 = False

						}

						ifres5686 = ifres5687

					} else {
						ifres5686 = False

					}

					var ifres5685 Obj

					if True == ifres5686 {
						ifres5685 = True

					} else {
						ifres5685 = False

					}

					ifres5684 = ifres5685

				} else {
					ifres5684 = False

				}

				var ifres5683 Obj

				if True == ifres5684 {
					ifres5683 = True

				} else {
					ifres5683 = False

				}

				ifres5682 = ifres5683

			} else {
				ifres5682 = False

			}

			var ifres5681 Obj

			if True == ifres5682 {
				ifres5681 = True

			} else {
				ifres5681 = False

			}

			ifres5680 = ifres5681

		} else {
			ifres5680 = False

		}

		if True == ifres5680 {
			tmp5638 := PrimTail(V3100)

			tmp5639 := PrimTail(tmp5638)

			__e.Return(PrimTail(tmp5639))
			return

		} else {
			tmp5678 := PrimIsPair(V3100)

			var ifres5665 Obj

			if True == tmp5678 {
				tmp5676 := PrimHead(V3100)

				tmp5677 := PrimEqual(sympackage, tmp5676)

				var ifres5667 Obj

				if True == tmp5677 {
					tmp5674 := PrimTail(V3100)

					tmp5675 := PrimIsPair(tmp5674)

					var ifres5669 Obj

					if True == tmp5675 {
						tmp5671 := PrimTail(V3100)

						tmp5672 := PrimTail(tmp5671)

						tmp5673 := PrimIsPair(tmp5672)

						var ifres5670 Obj

						if True == tmp5673 {
							ifres5670 = True

						} else {
							ifres5670 = False

						}

						ifres5669 = ifres5670

					} else {
						ifres5669 = False

					}

					var ifres5668 Obj

					if True == ifres5669 {
						ifres5668 = True

					} else {
						ifres5668 = False

					}

					ifres5667 = ifres5668

				} else {
					ifres5667 = False

				}

				var ifres5666 Obj

				if True == ifres5667 {
					ifres5666 = True

				} else {
					ifres5666 = False

				}

				ifres5665 = ifres5666

			} else {
				ifres5665 = False

			}

			if True == ifres5665 {
				tmp5640 := MakeNative(func(__e *ControlFlow) {
					W3101 := __e.Get(1)
					_ = W3101
					tmp5641 := MakeNative(func(__e *ControlFlow) {
						W3102 := __e.Get(1)
						_ = W3102
						tmp5642 := MakeNative(func(__e *ControlFlow) {
							W3103 := __e.Get(1)
							_ = W3103
							tmp5643 := MakeNative(func(__e *ControlFlow) {
								W3104 := __e.Get(1)
								_ = W3104
								__e.Return(W3102)
								return
							}, 1)

							tmp5644 := PrimTail(V3100)

							tmp5645 := PrimHead(tmp5644)

							tmp5646 := PrimTail(V3100)

							tmp5647 := PrimTail(tmp5646)

							tmp5648 := PrimTail(tmp5647)

							tmp5649 := Call(__e, PrimFunc(symshen_4record_1internal), tmp5645, W3101, tmp5648)

							__e.TailApply(tmp5643, tmp5649)
							return

						}, 1)

						tmp5650 := PrimTail(V3100)

						tmp5651 := PrimHead(tmp5650)

						tmp5652 := Call(__e, PrimFunc(symshen_4record_1external), tmp5651, W3101)

						__e.TailApply(tmp5642, tmp5652)
						return

					}, 1)

					tmp5653 := PrimTail(V3100)

					tmp5654 := PrimHead(tmp5653)

					tmp5655 := PrimStr(tmp5654)

					tmp5656 := PrimTail(V3100)

					tmp5657 := PrimTail(tmp5656)

					tmp5658 := PrimTail(tmp5657)

					tmp5659 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp5655, W3101, tmp5658)

					__e.TailApply(tmp5641, tmp5659)
					return

				}, 1)

				tmp5660 := PrimTail(V3100)

				tmp5661 := PrimTail(tmp5660)

				tmp5662 := PrimHead(tmp5661)

				tmp5663 := Call(__e, PrimFunc(symeval), tmp5662)

				__e.TailApply(tmp5640, tmp5663)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage)
				return
			}

		}

	}, 1)

	tmp5699 := Call(__e, ns2_1set, symshen_4unpackage, tmp5637)

	_ = tmp5699

	tmp5700 := MakeNative(func(__e *ControlFlow) {
		V3105 := __e.Get(1)
		_ = V3105
		V3106 := __e.Get(2)
		_ = V3106
		V3107 := __e.Get(3)
		_ = V3107
		tmp5701 := MakeNative(func(__e *ControlFlow) {
			W3108 := __e.Get(1)
			_ = W3108
			tmp5702 := MakeNative(func(__e *ControlFlow) {
				W3110 := __e.Get(1)
				_ = W3110
				tmp5703 := Call(__e, PrimFunc(symunion), W3110, W3108)

				tmp5704 := PrimValue(sym_dproperty_1vector_d)

				__e.TailApply(PrimFunc(symput), V3105, symshen_4internal_1symbols, tmp5703, tmp5704)
				return

			}, 1)

			tmp5705 := PrimStr(V3105)

			tmp5706 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp5705, V3106, V3107)

			__e.TailApply(tmp5702, tmp5706)
			return

		}, 1)

		tmp5707 := MakeNative(func(__e *ControlFlow) {
			tmp5708 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symget), V3105, symshen_4internal_1symbols, tmp5708)
			return

		}, 0)

		tmp5709 := MakeNative(func(__e *ControlFlow) {
			Z3109 := __e.Get(1)
			_ = Z3109
			__e.Return(Nil)
			return
		}, 1)

		tmp5710 := Call(__e, try_1catch, tmp5707, tmp5709)

		__e.TailApply(tmp5701, tmp5710)
		return

	}, 3)

	tmp5711 := Call(__e, ns2_1set, symshen_4record_1internal, tmp5700)

	_ = tmp5711

	tmp5712 := MakeNative(func(__e *ControlFlow) {
		V3117 := __e.Get(1)
		_ = V3117
		V3118 := __e.Get(2)
		_ = V3118
		V3119 := __e.Get(3)
		_ = V3119
		tmp5721 := PrimIsPair(V3119)

		if True == tmp5721 {
			tmp5713 := PrimHead(V3119)

			tmp5714 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3117, V3118, tmp5713)

			tmp5715 := PrimTail(V3119)

			tmp5716 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3117, V3118, tmp5715)

			__e.TailApply(PrimFunc(symunion), tmp5714, tmp5716)
			return

		} else {
			tmp5719 := Call(__e, PrimFunc(symshen_4internal_2), V3119, V3117, V3118)

			if True == tmp5719 {
				tmp5717 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V3117, V3119)

				__e.Return(PrimCons(tmp5717, Nil))
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 3)

	tmp5722 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp5712)

	_ = tmp5722

	tmp5723 := MakeNative(func(__e *ControlFlow) {
		V3120 := __e.Get(1)
		_ = V3120
		V3121 := __e.Get(2)
		_ = V3121
		tmp5724 := MakeNative(func(__e *ControlFlow) {
			W3122 := __e.Get(1)
			_ = W3122
			tmp5725 := Call(__e, PrimFunc(symunion), V3121, W3122)

			tmp5726 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symput), V3120, symshen_4external_1symbols, tmp5725, tmp5726)
			return

		}, 1)

		tmp5727 := MakeNative(func(__e *ControlFlow) {
			tmp5728 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symget), V3120, symshen_4external_1symbols, tmp5728)
			return

		}, 0)

		tmp5729 := MakeNative(func(__e *ControlFlow) {
			Z3123 := __e.Get(1)
			_ = Z3123
			__e.Return(Nil)
			return
		}, 1)

		tmp5730 := Call(__e, try_1catch, tmp5727, tmp5729)

		__e.TailApply(tmp5724, tmp5730)
		return

	}, 2)

	tmp5731 := Call(__e, ns2_1set, symshen_4record_1external, tmp5723)

	_ = tmp5731

	tmp5732 := MakeNative(func(__e *ControlFlow) {
		V3128 := __e.Get(1)
		_ = V3128
		V3129 := __e.Get(2)
		_ = V3129
		V3130 := __e.Get(3)
		_ = V3130
		tmp5737 := PrimIsPair(V3130)

		if True == tmp5737 {
			tmp5733 := MakeNative(func(__e *ControlFlow) {
				Z3131 := __e.Get(1)
				_ = Z3131
				__e.TailApply(PrimFunc(symshen_4package_1symbols), V3128, V3129, Z3131)
				return
			}, 1)

			__e.TailApply(PrimFunc(symmap), tmp5733, V3130)
			return

		} else {
			tmp5735 := Call(__e, PrimFunc(symshen_4internal_2), V3130, V3128, V3129)

			if True == tmp5735 {
				__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V3128, V3130)
				return
			} else {
				__e.Return(V3130)
				return
			}

		}

	}, 3)

	tmp5738 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp5732)

	_ = tmp5738

	tmp5739 := MakeNative(func(__e *ControlFlow) {
		V3132 := __e.Get(1)
		_ = V3132
		V3133 := __e.Get(2)
		_ = V3133
		tmp5740 := PrimStr(V3133)

		tmp5741 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp5740)

		tmp5742 := Call(__e, PrimFunc(sym_8s), V3132, tmp5741)

		__e.Return(PrimIntern(tmp5742))
		return

	}, 2)

	tmp5743 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp5739)

	_ = tmp5743

	tmp5744 := MakeNative(func(__e *ControlFlow) {
		V3134 := __e.Get(1)
		_ = V3134
		V3135 := __e.Get(2)
		_ = V3135
		V3136 := __e.Get(3)
		_ = V3136
		tmp5774 := Call(__e, PrimFunc(symelement_2), V3134, V3136)

		tmp5775 := PrimNot(tmp5774)

		if True == tmp5775 {
			tmp5771 := Call(__e, PrimFunc(symshen_4sng_2), V3134)

			tmp5772 := PrimNot(tmp5771)

			var ifres5746 Obj

			if True == tmp5772 {
				tmp5769 := Call(__e, PrimFunc(symshen_4dbl_2), V3134)

				tmp5770 := PrimNot(tmp5769)

				var ifres5748 Obj

				if True == tmp5770 {
					tmp5768 := PrimIsSymbol(V3134)

					var ifres5750 Obj

					if True == tmp5768 {
						tmp5766 := Call(__e, PrimFunc(symshen_4sysfunc_2), V3134)

						tmp5767 := PrimNot(tmp5766)

						var ifres5752 Obj

						if True == tmp5767 {
							tmp5764 := PrimIsVariable(V3134)

							tmp5765 := PrimNot(tmp5764)

							var ifres5754 Obj

							if True == tmp5765 {
								tmp5761 := PrimStr(V3134)

								tmp5762 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp5761)

								tmp5763 := PrimNot(tmp5762)

								var ifres5756 Obj

								if True == tmp5763 {
									tmp5758 := PrimStr(V3134)

									tmp5759 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V3135, tmp5758)

									tmp5760 := PrimNot(tmp5759)

									var ifres5757 Obj

									if True == tmp5760 {
										ifres5757 = True

									} else {
										ifres5757 = False

									}

									ifres5756 = ifres5757

								} else {
									ifres5756 = False

								}

								var ifres5755 Obj

								if True == ifres5756 {
									ifres5755 = True

								} else {
									ifres5755 = False

								}

								ifres5754 = ifres5755

							} else {
								ifres5754 = False

							}

							var ifres5753 Obj

							if True == ifres5754 {
								ifres5753 = True

							} else {
								ifres5753 = False

							}

							ifres5752 = ifres5753

						} else {
							ifres5752 = False

						}

						var ifres5751 Obj

						if True == ifres5752 {
							ifres5751 = True

						} else {
							ifres5751 = False

						}

						ifres5750 = ifres5751

					} else {
						ifres5750 = False

					}

					var ifres5749 Obj

					if True == ifres5750 {
						ifres5749 = True

					} else {
						ifres5749 = False

					}

					ifres5748 = ifres5749

				} else {
					ifres5748 = False

				}

				var ifres5747 Obj

				if True == ifres5748 {
					ifres5747 = True

				} else {
					ifres5747 = False

				}

				ifres5746 = ifres5747

			} else {
				ifres5746 = False

			}

			if True == ifres5746 {
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

	}, 3)

	tmp5776 := Call(__e, ns2_1set, symshen_4internal_2, tmp5744)

	_ = tmp5776

	tmp5777 := MakeNative(func(__e *ControlFlow) {
		V3141 := __e.Get(1)
		_ = V3141
		tmp5831 := Call(__e, PrimFunc(symshen_4_7string_2), V3141)

		var ifres5779 Obj

		if True == tmp5831 {
			tmp5829 := Call(__e, PrimFunc(symhdstr), V3141)

			tmp5830 := PrimEqual(MakeString("s"), tmp5829)

			var ifres5781 Obj

			if True == tmp5830 {
				tmp5827 := PrimTailString(V3141)

				tmp5828 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5827)

				var ifres5783 Obj

				if True == tmp5828 {
					tmp5824 := PrimTailString(V3141)

					tmp5825 := Call(__e, PrimFunc(symhdstr), tmp5824)

					tmp5826 := PrimEqual(MakeString("h"), tmp5825)

					var ifres5785 Obj

					if True == tmp5826 {
						tmp5821 := PrimTailString(V3141)

						tmp5822 := PrimTailString(tmp5821)

						tmp5823 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5822)

						var ifres5787 Obj

						if True == tmp5823 {
							tmp5817 := PrimTailString(V3141)

							tmp5818 := PrimTailString(tmp5817)

							tmp5819 := Call(__e, PrimFunc(symhdstr), tmp5818)

							tmp5820 := PrimEqual(MakeString("e"), tmp5819)

							var ifres5789 Obj

							if True == tmp5820 {
								tmp5813 := PrimTailString(V3141)

								tmp5814 := PrimTailString(tmp5813)

								tmp5815 := PrimTailString(tmp5814)

								tmp5816 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5815)

								var ifres5791 Obj

								if True == tmp5816 {
									tmp5808 := PrimTailString(V3141)

									tmp5809 := PrimTailString(tmp5808)

									tmp5810 := PrimTailString(tmp5809)

									tmp5811 := Call(__e, PrimFunc(symhdstr), tmp5810)

									tmp5812 := PrimEqual(MakeString("n"), tmp5811)

									var ifres5793 Obj

									if True == tmp5812 {
										tmp5803 := PrimTailString(V3141)

										tmp5804 := PrimTailString(tmp5803)

										tmp5805 := PrimTailString(tmp5804)

										tmp5806 := PrimTailString(tmp5805)

										tmp5807 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5806)

										var ifres5795 Obj

										if True == tmp5807 {
											tmp5797 := PrimTailString(V3141)

											tmp5798 := PrimTailString(tmp5797)

											tmp5799 := PrimTailString(tmp5798)

											tmp5800 := PrimTailString(tmp5799)

											tmp5801 := Call(__e, PrimFunc(symhdstr), tmp5800)

											tmp5802 := PrimEqual(MakeString("."), tmp5801)

											var ifres5796 Obj

											if True == tmp5802 {
												ifres5796 = True

											} else {
												ifres5796 = False

											}

											ifres5795 = ifres5796

										} else {
											ifres5795 = False

										}

										var ifres5794 Obj

										if True == ifres5795 {
											ifres5794 = True

										} else {
											ifres5794 = False

										}

										ifres5793 = ifres5794

									} else {
										ifres5793 = False

									}

									var ifres5792 Obj

									if True == ifres5793 {
										ifres5792 = True

									} else {
										ifres5792 = False

									}

									ifres5791 = ifres5792

								} else {
									ifres5791 = False

								}

								var ifres5790 Obj

								if True == ifres5791 {
									ifres5790 = True

								} else {
									ifres5790 = False

								}

								ifres5789 = ifres5790

							} else {
								ifres5789 = False

							}

							var ifres5788 Obj

							if True == ifres5789 {
								ifres5788 = True

							} else {
								ifres5788 = False

							}

							ifres5787 = ifres5788

						} else {
							ifres5787 = False

						}

						var ifres5786 Obj

						if True == ifres5787 {
							ifres5786 = True

						} else {
							ifres5786 = False

						}

						ifres5785 = ifres5786

					} else {
						ifres5785 = False

					}

					var ifres5784 Obj

					if True == ifres5785 {
						ifres5784 = True

					} else {
						ifres5784 = False

					}

					ifres5783 = ifres5784

				} else {
					ifres5783 = False

				}

				var ifres5782 Obj

				if True == ifres5783 {
					ifres5782 = True

				} else {
					ifres5782 = False

				}

				ifres5781 = ifres5782

			} else {
				ifres5781 = False

			}

			var ifres5780 Obj

			if True == ifres5781 {
				ifres5780 = True

			} else {
				ifres5780 = False

			}

			ifres5779 = ifres5780

		} else {
			ifres5779 = False

		}

		if True == ifres5779 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp5832 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp5777)

	_ = tmp5832

	tmp5833 := MakeNative(func(__e *ControlFlow) {
		V3142 := __e.Get(1)
		_ = V3142
		tmp5834 := PrimValue(sym_dproperty_1vector_d)

		tmp5835 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp5834)

		__e.TailApply(PrimFunc(symelement_2), V3142, tmp5835)
		return

	}, 1)

	tmp5836 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp5833)

	_ = tmp5836

	tmp5837 := MakeNative(func(__e *ControlFlow) {
		V3150 := __e.Get(1)
		_ = V3150
		V3151 := __e.Get(2)
		_ = V3151
		tmp5858 := PrimEqual(MakeString(""), V3150)

		var ifres5851 Obj

		if True == tmp5858 {
			tmp5857 := Call(__e, PrimFunc(symshen_4_7string_2), V3151)

			var ifres5853 Obj

			if True == tmp5857 {
				tmp5855 := Call(__e, PrimFunc(symhdstr), V3151)

				tmp5856 := PrimEqual(MakeString("."), tmp5855)

				var ifres5854 Obj

				if True == tmp5856 {
					ifres5854 = True

				} else {
					ifres5854 = False

				}

				ifres5853 = ifres5854

			} else {
				ifres5853 = False

			}

			var ifres5852 Obj

			if True == ifres5853 {
				ifres5852 = True

			} else {
				ifres5852 = False

			}

			ifres5851 = ifres5852

		} else {
			ifres5851 = False

		}

		if True == ifres5851 {
			__e.Return(True)
			return
		} else {
			tmp5849 := Call(__e, PrimFunc(symshen_4_7string_2), V3150)

			var ifres5841 Obj

			if True == tmp5849 {
				tmp5848 := Call(__e, PrimFunc(symshen_4_7string_2), V3151)

				var ifres5843 Obj

				if True == tmp5848 {
					tmp5845 := Call(__e, PrimFunc(symhdstr), V3150)

					tmp5846 := Call(__e, PrimFunc(symhdstr), V3151)

					tmp5847 := PrimEqual(tmp5845, tmp5846)

					var ifres5844 Obj

					if True == tmp5847 {
						ifres5844 = True

					} else {
						ifres5844 = False

					}

					ifres5843 = ifres5844

				} else {
					ifres5843 = False

				}

				var ifres5842 Obj

				if True == ifres5843 {
					ifres5842 = True

				} else {
					ifres5842 = False

				}

				ifres5841 = ifres5842

			} else {
				ifres5841 = False

			}

			if True == ifres5841 {
				tmp5838 := PrimTailString(V3150)

				tmp5839 := PrimTailString(V3151)

				__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp5838, tmp5839)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp5859 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp5837)

	_ = tmp5859

	tmp5860 := MakeNative(func(__e *ControlFlow) {
		V3154 := __e.Get(1)
		_ = V3154
		V3155 := __e.Get(2)
		_ = V3155
		tmp5881 := Call(__e, PrimFunc(symelement_2), V3154, V3155)

		if True == tmp5881 {
			__e.Return(V3154)
			return
		} else {
			tmp5879 := PrimIsPair(V3154)

			var ifres5875 Obj

			if True == tmp5879 {
				tmp5877 := PrimHead(V3154)

				tmp5878 := PrimEqual(symcond, tmp5877)

				var ifres5876 Obj

				if True == tmp5878 {
					ifres5876 = True

				} else {
					ifres5876 = False

				}

				ifres5875 = ifres5876

			} else {
				ifres5875 = False

			}

			if True == ifres5875 {
				tmp5861 := PrimTail(V3154)

				tmp5862 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp5861, V3155)

				__e.Return(PrimCons(symcond, tmp5862))
				return

			} else {
				tmp5873 := PrimIsPair(V3154)

				var ifres5869 Obj

				if True == tmp5873 {
					tmp5871 := PrimHead(V3154)

					tmp5872 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp5871)

					var ifres5870 Obj

					if True == tmp5872 {
						ifres5870 = True

					} else {
						ifres5870 = False

					}

					ifres5869 = ifres5870

				} else {
					ifres5869 = False

				}

				if True == ifres5869 {
					tmp5863 := PrimHead(V3154)

					__e.TailApply(PrimFunc(symshen_4special_1case), tmp5863, V3154, V3155)
					return

				} else {
					tmp5867 := PrimIsPair(V3154)

					if True == tmp5867 {
						tmp5864 := MakeNative(func(__e *ControlFlow) {
							Z3156 := __e.Get(1)
							_ = Z3156
							__e.TailApply(PrimFunc(symshen_4process_1applications), Z3156, V3155)
							return
						}, 1)

						tmp5865 := Call(__e, PrimFunc(symmap), tmp5864, V3154)

						__e.TailApply(PrimFunc(symshen_4process_1application), tmp5865, V3155)
						return

					} else {
						__e.Return(V3154)
						return
					}

				}

			}

		}

	}, 2)

	tmp5882 := Call(__e, ns2_1set, symshen_4process_1applications, tmp5860)

	_ = tmp5882

	tmp5883 := MakeNative(func(__e *ControlFlow) {
		V3159 := __e.Get(1)
		_ = V3159
		tmp5893 := PrimEqual(symdefine, V3159)

		if True == tmp5893 {
			__e.Return(True)
			return
		} else {
			tmp5891 := PrimEqual(symdefun, V3159)

			if True == tmp5891 {
				__e.Return(True)
				return
			} else {
				tmp5889 := PrimEqual(symsynonyms, V3159)

				if True == tmp5889 {
					__e.Return(True)
					return
				} else {
					tmp5887 := Call(__e, PrimFunc(symshen_4special_2), V3159)

					if True == tmp5887 {
						__e.Return(True)
						return
					} else {
						tmp5885 := Call(__e, PrimFunc(symshen_4extraspecial_2), V3159)

						if True == tmp5885 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp5894 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp5883)

	_ = tmp5894

	tmp5895 := MakeNative(func(__e *ControlFlow) {
		V3164 := __e.Get(1)
		_ = V3164
		V3165 := __e.Get(2)
		_ = V3165
		V3166 := __e.Get(3)
		_ = V3166
		tmp6137 := PrimEqual(symlambda, V3164)

		var ifres6115 Obj

		if True == tmp6137 {
			tmp6136 := PrimIsPair(V3165)

			var ifres6117 Obj

			if True == tmp6136 {
				tmp6134 := PrimHead(V3165)

				tmp6135 := PrimEqual(symlambda, tmp6134)

				var ifres6119 Obj

				if True == tmp6135 {
					tmp6132 := PrimTail(V3165)

					tmp6133 := PrimIsPair(tmp6132)

					var ifres6121 Obj

					if True == tmp6133 {
						tmp6129 := PrimTail(V3165)

						tmp6130 := PrimTail(tmp6129)

						tmp6131 := PrimIsPair(tmp6130)

						var ifres6123 Obj

						if True == tmp6131 {
							tmp6125 := PrimTail(V3165)

							tmp6126 := PrimTail(tmp6125)

							tmp6127 := PrimTail(tmp6126)

							tmp6128 := PrimEqual(Nil, tmp6127)

							var ifres6124 Obj

							if True == tmp6128 {
								ifres6124 = True

							} else {
								ifres6124 = False

							}

							ifres6123 = ifres6124

						} else {
							ifres6123 = False

						}

						var ifres6122 Obj

						if True == ifres6123 {
							ifres6122 = True

						} else {
							ifres6122 = False

						}

						ifres6121 = ifres6122

					} else {
						ifres6121 = False

					}

					var ifres6120 Obj

					if True == ifres6121 {
						ifres6120 = True

					} else {
						ifres6120 = False

					}

					ifres6119 = ifres6120

				} else {
					ifres6119 = False

				}

				var ifres6118 Obj

				if True == ifres6119 {
					ifres6118 = True

				} else {
					ifres6118 = False

				}

				ifres6117 = ifres6118

			} else {
				ifres6117 = False

			}

			var ifres6116 Obj

			if True == ifres6117 {
				ifres6116 = True

			} else {
				ifres6116 = False

			}

			ifres6115 = ifres6116

		} else {
			ifres6115 = False

		}

		if True == ifres6115 {
			tmp5896 := PrimTail(V3165)

			tmp5897 := PrimHead(tmp5896)

			tmp5898 := PrimTail(V3165)

			tmp5899 := PrimTail(tmp5898)

			tmp5900 := PrimHead(tmp5899)

			tmp5901 := Call(__e, PrimFunc(symshen_4process_1applications), tmp5900, V3166)

			tmp5902 := PrimCons(tmp5901, Nil)

			tmp5903 := PrimCons(tmp5897, tmp5902)

			__e.Return(PrimCons(symlambda, tmp5903))
			return

		} else {
			tmp6113 := PrimEqual(symlet, V3164)

			var ifres6084 Obj

			if True == tmp6113 {
				tmp6112 := PrimIsPair(V3165)

				var ifres6086 Obj

				if True == tmp6112 {
					tmp6110 := PrimHead(V3165)

					tmp6111 := PrimEqual(symlet, tmp6110)

					var ifres6088 Obj

					if True == tmp6111 {
						tmp6108 := PrimTail(V3165)

						tmp6109 := PrimIsPair(tmp6108)

						var ifres6090 Obj

						if True == tmp6109 {
							tmp6105 := PrimTail(V3165)

							tmp6106 := PrimTail(tmp6105)

							tmp6107 := PrimIsPair(tmp6106)

							var ifres6092 Obj

							if True == tmp6107 {
								tmp6101 := PrimTail(V3165)

								tmp6102 := PrimTail(tmp6101)

								tmp6103 := PrimTail(tmp6102)

								tmp6104 := PrimIsPair(tmp6103)

								var ifres6094 Obj

								if True == tmp6104 {
									tmp6096 := PrimTail(V3165)

									tmp6097 := PrimTail(tmp6096)

									tmp6098 := PrimTail(tmp6097)

									tmp6099 := PrimTail(tmp6098)

									tmp6100 := PrimEqual(Nil, tmp6099)

									var ifres6095 Obj

									if True == tmp6100 {
										ifres6095 = True

									} else {
										ifres6095 = False

									}

									ifres6094 = ifres6095

								} else {
									ifres6094 = False

								}

								var ifres6093 Obj

								if True == ifres6094 {
									ifres6093 = True

								} else {
									ifres6093 = False

								}

								ifres6092 = ifres6093

							} else {
								ifres6092 = False

							}

							var ifres6091 Obj

							if True == ifres6092 {
								ifres6091 = True

							} else {
								ifres6091 = False

							}

							ifres6090 = ifres6091

						} else {
							ifres6090 = False

						}

						var ifres6089 Obj

						if True == ifres6090 {
							ifres6089 = True

						} else {
							ifres6089 = False

						}

						ifres6088 = ifres6089

					} else {
						ifres6088 = False

					}

					var ifres6087 Obj

					if True == ifres6088 {
						ifres6087 = True

					} else {
						ifres6087 = False

					}

					ifres6086 = ifres6087

				} else {
					ifres6086 = False

				}

				var ifres6085 Obj

				if True == ifres6086 {
					ifres6085 = True

				} else {
					ifres6085 = False

				}

				ifres6084 = ifres6085

			} else {
				ifres6084 = False

			}

			if True == ifres6084 {
				tmp5904 := PrimTail(V3165)

				tmp5905 := PrimHead(tmp5904)

				tmp5906 := PrimTail(V3165)

				tmp5907 := PrimTail(tmp5906)

				tmp5908 := PrimHead(tmp5907)

				tmp5909 := Call(__e, PrimFunc(symshen_4process_1applications), tmp5908, V3166)

				tmp5910 := PrimTail(V3165)

				tmp5911 := PrimTail(tmp5910)

				tmp5912 := PrimTail(tmp5911)

				tmp5913 := PrimHead(tmp5912)

				tmp5914 := Call(__e, PrimFunc(symshen_4process_1applications), tmp5913, V3166)

				tmp5915 := PrimCons(tmp5914, Nil)

				tmp5916 := PrimCons(tmp5909, tmp5915)

				tmp5917 := PrimCons(tmp5905, tmp5916)

				__e.Return(PrimCons(symlet, tmp5917))
				return

			} else {
				tmp6082 := PrimEqual(symdefun, V3164)

				var ifres6053 Obj

				if True == tmp6082 {
					tmp6081 := PrimIsPair(V3165)

					var ifres6055 Obj

					if True == tmp6081 {
						tmp6079 := PrimHead(V3165)

						tmp6080 := PrimEqual(symdefun, tmp6079)

						var ifres6057 Obj

						if True == tmp6080 {
							tmp6077 := PrimTail(V3165)

							tmp6078 := PrimIsPair(tmp6077)

							var ifres6059 Obj

							if True == tmp6078 {
								tmp6074 := PrimTail(V3165)

								tmp6075 := PrimTail(tmp6074)

								tmp6076 := PrimIsPair(tmp6075)

								var ifres6061 Obj

								if True == tmp6076 {
									tmp6070 := PrimTail(V3165)

									tmp6071 := PrimTail(tmp6070)

									tmp6072 := PrimTail(tmp6071)

									tmp6073 := PrimIsPair(tmp6072)

									var ifres6063 Obj

									if True == tmp6073 {
										tmp6065 := PrimTail(V3165)

										tmp6066 := PrimTail(tmp6065)

										tmp6067 := PrimTail(tmp6066)

										tmp6068 := PrimTail(tmp6067)

										tmp6069 := PrimEqual(Nil, tmp6068)

										var ifres6064 Obj

										if True == tmp6069 {
											ifres6064 = True

										} else {
											ifres6064 = False

										}

										ifres6063 = ifres6064

									} else {
										ifres6063 = False

									}

									var ifres6062 Obj

									if True == ifres6063 {
										ifres6062 = True

									} else {
										ifres6062 = False

									}

									ifres6061 = ifres6062

								} else {
									ifres6061 = False

								}

								var ifres6060 Obj

								if True == ifres6061 {
									ifres6060 = True

								} else {
									ifres6060 = False

								}

								ifres6059 = ifres6060

							} else {
								ifres6059 = False

							}

							var ifres6058 Obj

							if True == ifres6059 {
								ifres6058 = True

							} else {
								ifres6058 = False

							}

							ifres6057 = ifres6058

						} else {
							ifres6057 = False

						}

						var ifres6056 Obj

						if True == ifres6057 {
							ifres6056 = True

						} else {
							ifres6056 = False

						}

						ifres6055 = ifres6056

					} else {
						ifres6055 = False

					}

					var ifres6054 Obj

					if True == ifres6055 {
						ifres6054 = True

					} else {
						ifres6054 = False

					}

					ifres6053 = ifres6054

				} else {
					ifres6053 = False

				}

				if True == ifres6053 {
					__e.Return(V3165)
					return
				} else {
					tmp6051 := PrimEqual(symdefine, V3164)

					var ifres6029 Obj

					if True == tmp6051 {
						tmp6050 := PrimIsPair(V3165)

						var ifres6031 Obj

						if True == tmp6050 {
							tmp6048 := PrimHead(V3165)

							tmp6049 := PrimEqual(symdefine, tmp6048)

							var ifres6033 Obj

							if True == tmp6049 {
								tmp6046 := PrimTail(V3165)

								tmp6047 := PrimIsPair(tmp6046)

								var ifres6035 Obj

								if True == tmp6047 {
									tmp6043 := PrimTail(V3165)

									tmp6044 := PrimTail(tmp6043)

									tmp6045 := PrimIsPair(tmp6044)

									var ifres6037 Obj

									if True == tmp6045 {
										tmp6039 := PrimTail(V3165)

										tmp6040 := PrimTail(tmp6039)

										tmp6041 := PrimHead(tmp6040)

										tmp6042 := PrimEqual(sym_i, tmp6041)

										var ifres6038 Obj

										if True == tmp6042 {
											ifres6038 = True

										} else {
											ifres6038 = False

										}

										ifres6037 = ifres6038

									} else {
										ifres6037 = False

									}

									var ifres6036 Obj

									if True == ifres6037 {
										ifres6036 = True

									} else {
										ifres6036 = False

									}

									ifres6035 = ifres6036

								} else {
									ifres6035 = False

								}

								var ifres6034 Obj

								if True == ifres6035 {
									ifres6034 = True

								} else {
									ifres6034 = False

								}

								ifres6033 = ifres6034

							} else {
								ifres6033 = False

							}

							var ifres6032 Obj

							if True == ifres6033 {
								ifres6032 = True

							} else {
								ifres6032 = False

							}

							ifres6031 = ifres6032

						} else {
							ifres6031 = False

						}

						var ifres6030 Obj

						if True == ifres6031 {
							ifres6030 = True

						} else {
							ifres6030 = False

						}

						ifres6029 = ifres6030

					} else {
						ifres6029 = False

					}

					if True == ifres6029 {
						tmp5918 := PrimTail(V3165)

						tmp5919 := PrimHead(tmp5918)

						tmp5920 := PrimTail(V3165)

						tmp5921 := PrimHead(tmp5920)

						tmp5922 := PrimTail(V3165)

						tmp5923 := PrimTail(tmp5922)

						tmp5924 := PrimTail(tmp5923)

						tmp5925 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp5921, tmp5924, V3166)

						tmp5926 := PrimCons(sym_i, tmp5925)

						tmp5927 := PrimCons(tmp5919, tmp5926)

						__e.Return(PrimCons(symdefine, tmp5927))
						return

					} else {
						tmp6027 := PrimEqual(symdefine, V3164)

						var ifres6016 Obj

						if True == tmp6027 {
							tmp6026 := PrimIsPair(V3165)

							var ifres6018 Obj

							if True == tmp6026 {
								tmp6024 := PrimHead(V3165)

								tmp6025 := PrimEqual(symdefine, tmp6024)

								var ifres6020 Obj

								if True == tmp6025 {
									tmp6022 := PrimTail(V3165)

									tmp6023 := PrimIsPair(tmp6022)

									var ifres6021 Obj

									if True == tmp6023 {
										ifres6021 = True

									} else {
										ifres6021 = False

									}

									ifres6020 = ifres6021

								} else {
									ifres6020 = False

								}

								var ifres6019 Obj

								if True == ifres6020 {
									ifres6019 = True

								} else {
									ifres6019 = False

								}

								ifres6018 = ifres6019

							} else {
								ifres6018 = False

							}

							var ifres6017 Obj

							if True == ifres6018 {
								ifres6017 = True

							} else {
								ifres6017 = False

							}

							ifres6016 = ifres6017

						} else {
							ifres6016 = False

						}

						if True == ifres6016 {
							tmp5928 := PrimTail(V3165)

							tmp5929 := PrimHead(tmp5928)

							tmp5930 := MakeNative(func(__e *ControlFlow) {
								Z3167 := __e.Get(1)
								_ = Z3167
								__e.TailApply(PrimFunc(symshen_4process_1applications), Z3167, V3166)
								return
							}, 1)

							tmp5931 := PrimTail(V3165)

							tmp5932 := PrimTail(tmp5931)

							tmp5933 := Call(__e, PrimFunc(symmap), tmp5930, tmp5932)

							tmp5934 := PrimCons(tmp5929, tmp5933)

							__e.Return(PrimCons(symdefine, tmp5934))
							return

						} else {
							tmp6014 := PrimEqual(symsynonyms, V3164)

							if True == tmp6014 {
								__e.Return(PrimCons(symsynonyms, V3165))
								return
							} else {
								tmp6012 := PrimEqual(symtype, V3164)

								var ifres5990 Obj

								if True == tmp6012 {
									tmp6011 := PrimIsPair(V3165)

									var ifres5992 Obj

									if True == tmp6011 {
										tmp6009 := PrimHead(V3165)

										tmp6010 := PrimEqual(symtype, tmp6009)

										var ifres5994 Obj

										if True == tmp6010 {
											tmp6007 := PrimTail(V3165)

											tmp6008 := PrimIsPair(tmp6007)

											var ifres5996 Obj

											if True == tmp6008 {
												tmp6004 := PrimTail(V3165)

												tmp6005 := PrimTail(tmp6004)

												tmp6006 := PrimIsPair(tmp6005)

												var ifres5998 Obj

												if True == tmp6006 {
													tmp6000 := PrimTail(V3165)

													tmp6001 := PrimTail(tmp6000)

													tmp6002 := PrimTail(tmp6001)

													tmp6003 := PrimEqual(Nil, tmp6002)

													var ifres5999 Obj

													if True == tmp6003 {
														ifres5999 = True

													} else {
														ifres5999 = False

													}

													ifres5998 = ifres5999

												} else {
													ifres5998 = False

												}

												var ifres5997 Obj

												if True == ifres5998 {
													ifres5997 = True

												} else {
													ifres5997 = False

												}

												ifres5996 = ifres5997

											} else {
												ifres5996 = False

											}

											var ifres5995 Obj

											if True == ifres5996 {
												ifres5995 = True

											} else {
												ifres5995 = False

											}

											ifres5994 = ifres5995

										} else {
											ifres5994 = False

										}

										var ifres5993 Obj

										if True == ifres5994 {
											ifres5993 = True

										} else {
											ifres5993 = False

										}

										ifres5992 = ifres5993

									} else {
										ifres5992 = False

									}

									var ifres5991 Obj

									if True == ifres5992 {
										ifres5991 = True

									} else {
										ifres5991 = False

									}

									ifres5990 = ifres5991

								} else {
									ifres5990 = False

								}

								if True == ifres5990 {
									tmp5935 := PrimTail(V3165)

									tmp5936 := PrimHead(tmp5935)

									tmp5937 := Call(__e, PrimFunc(symshen_4process_1applications), tmp5936, V3166)

									tmp5938 := PrimTail(V3165)

									tmp5939 := PrimTail(tmp5938)

									tmp5940 := PrimCons(tmp5937, tmp5939)

									__e.Return(PrimCons(symtype, tmp5940))
									return

								} else {
									tmp5988 := PrimEqual(syminput_7, V3164)

									var ifres5966 Obj

									if True == tmp5988 {
										tmp5987 := PrimIsPair(V3165)

										var ifres5968 Obj

										if True == tmp5987 {
											tmp5985 := PrimHead(V3165)

											tmp5986 := PrimEqual(syminput_7, tmp5985)

											var ifres5970 Obj

											if True == tmp5986 {
												tmp5983 := PrimTail(V3165)

												tmp5984 := PrimIsPair(tmp5983)

												var ifres5972 Obj

												if True == tmp5984 {
													tmp5980 := PrimTail(V3165)

													tmp5981 := PrimTail(tmp5980)

													tmp5982 := PrimIsPair(tmp5981)

													var ifres5974 Obj

													if True == tmp5982 {
														tmp5976 := PrimTail(V3165)

														tmp5977 := PrimTail(tmp5976)

														tmp5978 := PrimTail(tmp5977)

														tmp5979 := PrimEqual(Nil, tmp5978)

														var ifres5975 Obj

														if True == tmp5979 {
															ifres5975 = True

														} else {
															ifres5975 = False

														}

														ifres5974 = ifres5975

													} else {
														ifres5974 = False

													}

													var ifres5973 Obj

													if True == ifres5974 {
														ifres5973 = True

													} else {
														ifres5973 = False

													}

													ifres5972 = ifres5973

												} else {
													ifres5972 = False

												}

												var ifres5971 Obj

												if True == ifres5972 {
													ifres5971 = True

												} else {
													ifres5971 = False

												}

												ifres5970 = ifres5971

											} else {
												ifres5970 = False

											}

											var ifres5969 Obj

											if True == ifres5970 {
												ifres5969 = True

											} else {
												ifres5969 = False

											}

											ifres5968 = ifres5969

										} else {
											ifres5968 = False

										}

										var ifres5967 Obj

										if True == ifres5968 {
											ifres5967 = True

										} else {
											ifres5967 = False

										}

										ifres5966 = ifres5967

									} else {
										ifres5966 = False

									}

									if True == ifres5966 {
										tmp5941 := PrimTail(V3165)

										tmp5942 := PrimHead(tmp5941)

										tmp5943 := PrimTail(V3165)

										tmp5944 := PrimTail(tmp5943)

										tmp5945 := PrimHead(tmp5944)

										tmp5946 := Call(__e, PrimFunc(symshen_4process_1applications), tmp5945, V3166)

										tmp5947 := PrimCons(tmp5946, Nil)

										tmp5948 := PrimCons(tmp5942, tmp5947)

										__e.Return(PrimCons(syminput_7, tmp5948))
										return

									} else {
										tmp5964 := PrimIsPair(V3165)

										var ifres5960 Obj

										if True == tmp5964 {
											tmp5962 := PrimHead(V3165)

											tmp5963 := Call(__e, PrimFunc(symshen_4special_2), tmp5962)

											var ifres5961 Obj

											if True == tmp5963 {
												ifres5961 = True

											} else {
												ifres5961 = False

											}

											ifres5960 = ifres5961

										} else {
											ifres5960 = False

										}

										if True == ifres5960 {
											tmp5949 := PrimHead(V3165)

											tmp5950 := MakeNative(func(__e *ControlFlow) {
												Z3168 := __e.Get(1)
												_ = Z3168
												__e.TailApply(PrimFunc(symshen_4process_1applications), Z3168, V3166)
												return
											}, 1)

											tmp5951 := PrimTail(V3165)

											tmp5952 := Call(__e, PrimFunc(symmap), tmp5950, tmp5951)

											__e.Return(PrimCons(tmp5949, tmp5952))
											return

										} else {
											tmp5958 := PrimIsPair(V3165)

											var ifres5954 Obj

											if True == tmp5958 {
												tmp5956 := PrimHead(V3165)

												tmp5957 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp5956)

												var ifres5955 Obj

												if True == tmp5957 {
													ifres5955 = True

												} else {
													ifres5955 = False

												}

												ifres5954 = ifres5955

											} else {
												ifres5954 = False

											}

											if True == ifres5954 {
												__e.Return(V3165)
												return
											} else {
												__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4special_1case)
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

	}, 3)

	tmp6138 := Call(__e, ns2_1set, symshen_4special_1case, tmp5895)

	_ = tmp6138

	tmp6139 := MakeNative(func(__e *ControlFlow) {
		V3171 := __e.Get(1)
		_ = V3171
		V3172 := __e.Get(2)
		_ = V3172
		tmp6169 := PrimEqual(Nil, V3171)

		if True == tmp6169 {
			__e.Return(Nil)
			return
		} else {
			tmp6167 := PrimIsPair(V3171)

			var ifres6152 Obj

			if True == tmp6167 {
				tmp6165 := PrimHead(V3171)

				tmp6166 := PrimIsPair(tmp6165)

				var ifres6154 Obj

				if True == tmp6166 {
					tmp6162 := PrimHead(V3171)

					tmp6163 := PrimTail(tmp6162)

					tmp6164 := PrimIsPair(tmp6163)

					var ifres6156 Obj

					if True == tmp6164 {
						tmp6158 := PrimHead(V3171)

						tmp6159 := PrimTail(tmp6158)

						tmp6160 := PrimTail(tmp6159)

						tmp6161 := PrimEqual(Nil, tmp6160)

						var ifres6157 Obj

						if True == tmp6161 {
							ifres6157 = True

						} else {
							ifres6157 = False

						}

						ifres6156 = ifres6157

					} else {
						ifres6156 = False

					}

					var ifres6155 Obj

					if True == ifres6156 {
						ifres6155 = True

					} else {
						ifres6155 = False

					}

					ifres6154 = ifres6155

				} else {
					ifres6154 = False

				}

				var ifres6153 Obj

				if True == ifres6154 {
					ifres6153 = True

				} else {
					ifres6153 = False

				}

				ifres6152 = ifres6153

			} else {
				ifres6152 = False

			}

			if True == ifres6152 {
				tmp6140 := PrimHead(V3171)

				tmp6141 := PrimHead(tmp6140)

				tmp6142 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6141, V3172)

				tmp6143 := PrimHead(V3171)

				tmp6144 := PrimTail(tmp6143)

				tmp6145 := PrimHead(tmp6144)

				tmp6146 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6145, V3172)

				tmp6147 := PrimCons(tmp6146, Nil)

				tmp6148 := PrimCons(tmp6142, tmp6147)

				tmp6149 := PrimTail(V3171)

				tmp6150 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp6149, V3172)

				__e.Return(PrimCons(tmp6148, tmp6150))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1cond_1clauses)
				return
			}

		}

	}, 2)

	tmp6170 := Call(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp6139)

	_ = tmp6170

	tmp6171 := MakeNative(func(__e *ControlFlow) {
		V3175 := __e.Get(1)
		_ = V3175
		V3176 := __e.Get(2)
		_ = V3176
		V3177 := __e.Get(3)
		_ = V3177
		tmp6187 := PrimIsPair(V3176)

		var ifres6183 Obj

		if True == tmp6187 {
			tmp6185 := PrimHead(V3176)

			tmp6186 := PrimEqual(sym_j, tmp6185)

			var ifres6184 Obj

			if True == tmp6186 {
				ifres6184 = True

			} else {
				ifres6184 = False

			}

			ifres6183 = ifres6184

		} else {
			ifres6183 = False

		}

		if True == ifres6183 {
			tmp6172 := MakeNative(func(__e *ControlFlow) {
				Z3178 := __e.Get(1)
				_ = Z3178
				__e.TailApply(PrimFunc(symshen_4process_1applications), Z3178, V3177)
				return
			}, 1)

			tmp6173 := PrimTail(V3176)

			tmp6174 := Call(__e, PrimFunc(symmap), tmp6172, tmp6173)

			__e.Return(PrimCons(sym_j, tmp6174))
			return

		} else {
			tmp6181 := PrimIsPair(V3176)

			if True == tmp6181 {
				tmp6175 := PrimHead(V3176)

				tmp6176 := PrimTail(V3176)

				tmp6177 := Call(__e, PrimFunc(symshen_4process_1after_1type), V3175, tmp6176, V3177)

				__e.Return(PrimCons(tmp6175, tmp6177))
				return

			} else {
				tmp6178 := Call(__e, PrimFunc(symshen_4app), V3175, MakeString("\n"), symshen_4a)

				tmp6179 := PrimStringConcat(MakeString("missing } in "), tmp6178)

				__e.Return(PrimSimpleError(tmp6179))
				return

			}

		}

	}, 3)

	tmp6188 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp6171)

	_ = tmp6188

	tmp6189 := MakeNative(func(__e *ControlFlow) {
		V3179 := __e.Get(1)
		_ = V3179
		V3180 := __e.Get(2)
		_ = V3180
		tmp6234 := PrimIsPair(V3179)

		if True == tmp6234 {
			tmp6190 := MakeNative(func(__e *ControlFlow) {
				W3181 := __e.Get(1)
				_ = W3181
				tmp6191 := MakeNative(func(__e *ControlFlow) {
					W3182 := __e.Get(1)
					_ = W3182
					tmp6228 := Call(__e, PrimFunc(symelement_2), V3179, V3180)

					if True == tmp6228 {
						__e.Return(V3179)
						return
					} else {
						tmp6225 := PrimHead(V3179)

						tmp6226 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp6225)

						if True == tmp6226 {
							__e.Return(V3179)
							return
						} else {
							tmp6223 := Call(__e, PrimFunc(symshen_4foreign_2), V3179)

							if True == tmp6223 {
								__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V3179)
								return
							} else {
								tmp6221 := Call(__e, PrimFunc(symshen_4fn_1call_2), V3179)

								if True == tmp6221 {
									__e.TailApply(PrimFunc(symshen_4fn_1call), V3179)
									return
								} else {
									tmp6219 := Call(__e, PrimFunc(symshen_4zero_1place_2), V3179)

									if True == tmp6219 {
										__e.Return(V3179)
										return
									} else {
										tmp6216 := PrimHead(V3179)

										tmp6217 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp6216, W3181)

										if True == tmp6217 {
											tmp6192 := PrimHead(V3179)

											tmp6193 := PrimCons(tmp6192, Nil)

											tmp6194 := PrimCons(symfn, tmp6193)

											tmp6195 := PrimTail(V3179)

											tmp6196 := PrimCons(tmp6194, tmp6195)

											__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6196)
											return

										} else {
											tmp6213 := PrimHead(V3179)

											tmp6214 := PrimIsVariable(tmp6213)

											if True == tmp6214 {
												__e.TailApply(PrimFunc(symshen_4simple_1curry), V3179)
												return
											} else {
												tmp6210 := PrimHead(V3179)

												tmp6211 := Call(__e, PrimFunc(symshen_4application_2), tmp6210)

												if True == tmp6211 {
													__e.TailApply(PrimFunc(symshen_4simple_1curry), V3179)
													return
												} else {
													tmp6207 := PrimHead(V3179)

													tmp6208 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp6207, W3181, W3182)

													if True == tmp6208 {
														tmp6197 := PrimNumberSubtract(W3181, W3182)

														__e.TailApply(PrimFunc(symshen_4lambda_1function), V3179, tmp6197)
														return

													} else {
														tmp6204 := PrimHead(V3179)

														tmp6205 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp6204, W3181, W3182)

														if True == tmp6205 {
															tmp6198 := PrimHead(V3179)

															tmp6199 := PrimCons(tmp6198, Nil)

															tmp6200 := PrimCons(symfn, tmp6199)

															tmp6201 := PrimTail(V3179)

															tmp6202 := PrimCons(tmp6200, tmp6201)

															__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6202)
															return

														} else {
															__e.Return(V3179)
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

				}, 1)

				tmp6229 := PrimTail(V3179)

				tmp6230 := Call(__e, PrimFunc(symlength), tmp6229)

				__e.TailApply(tmp6191, tmp6230)
				return

			}, 1)

			tmp6231 := PrimHead(V3179)

			tmp6232 := Call(__e, PrimFunc(symarity), tmp6231)

			__e.TailApply(tmp6190, tmp6232)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1application)
			return
		}

	}, 2)

	tmp6235 := Call(__e, ns2_1set, symshen_4process_1application, tmp6189)

	_ = tmp6235

	tmp6236 := MakeNative(func(__e *ControlFlow) {
		V3183 := __e.Get(1)
		_ = V3183
		tmp6262 := PrimIsPair(V3183)

		var ifres6242 Obj

		if True == tmp6262 {
			tmp6260 := PrimHead(V3183)

			tmp6261 := PrimIsPair(tmp6260)

			var ifres6244 Obj

			if True == tmp6261 {
				tmp6257 := PrimHead(V3183)

				tmp6258 := PrimHead(tmp6257)

				tmp6259 := PrimEqual(symforeign, tmp6258)

				var ifres6246 Obj

				if True == tmp6259 {
					tmp6254 := PrimHead(V3183)

					tmp6255 := PrimTail(tmp6254)

					tmp6256 := PrimIsPair(tmp6255)

					var ifres6248 Obj

					if True == tmp6256 {
						tmp6250 := PrimHead(V3183)

						tmp6251 := PrimTail(tmp6250)

						tmp6252 := PrimTail(tmp6251)

						tmp6253 := PrimEqual(Nil, tmp6252)

						var ifres6249 Obj

						if True == tmp6253 {
							ifres6249 = True

						} else {
							ifres6249 = False

						}

						ifres6248 = ifres6249

					} else {
						ifres6248 = False

					}

					var ifres6247 Obj

					if True == ifres6248 {
						ifres6247 = True

					} else {
						ifres6247 = False

					}

					ifres6246 = ifres6247

				} else {
					ifres6246 = False

				}

				var ifres6245 Obj

				if True == ifres6246 {
					ifres6245 = True

				} else {
					ifres6245 = False

				}

				ifres6244 = ifres6245

			} else {
				ifres6244 = False

			}

			var ifres6243 Obj

			if True == ifres6244 {
				ifres6243 = True

			} else {
				ifres6243 = False

			}

			ifres6242 = ifres6243

		} else {
			ifres6242 = False

		}

		if True == ifres6242 {
			tmp6237 := PrimHead(V3183)

			tmp6238 := PrimTail(tmp6237)

			tmp6239 := PrimHead(tmp6238)

			tmp6240 := PrimTail(V3183)

			__e.Return(PrimCons(tmp6239, tmp6240))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpack_1foreign)
			return
		}

	}, 1)

	tmp6263 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp6236)

	_ = tmp6263

	tmp6264 := MakeNative(func(__e *ControlFlow) {
		V3186 := __e.Get(1)
		_ = V3186
		tmp6286 := PrimIsPair(V3186)

		var ifres6266 Obj

		if True == tmp6286 {
			tmp6284 := PrimHead(V3186)

			tmp6285 := PrimIsPair(tmp6284)

			var ifres6268 Obj

			if True == tmp6285 {
				tmp6281 := PrimHead(V3186)

				tmp6282 := PrimHead(tmp6281)

				tmp6283 := PrimEqual(symforeign, tmp6282)

				var ifres6270 Obj

				if True == tmp6283 {
					tmp6278 := PrimHead(V3186)

					tmp6279 := PrimTail(tmp6278)

					tmp6280 := PrimIsPair(tmp6279)

					var ifres6272 Obj

					if True == tmp6280 {
						tmp6274 := PrimHead(V3186)

						tmp6275 := PrimTail(tmp6274)

						tmp6276 := PrimTail(tmp6275)

						tmp6277 := PrimEqual(Nil, tmp6276)

						var ifres6273 Obj

						if True == tmp6277 {
							ifres6273 = True

						} else {
							ifres6273 = False

						}

						ifres6272 = ifres6273

					} else {
						ifres6272 = False

					}

					var ifres6271 Obj

					if True == ifres6272 {
						ifres6271 = True

					} else {
						ifres6271 = False

					}

					ifres6270 = ifres6271

				} else {
					ifres6270 = False

				}

				var ifres6269 Obj

				if True == ifres6270 {
					ifres6269 = True

				} else {
					ifres6269 = False

				}

				ifres6268 = ifres6269

			} else {
				ifres6268 = False

			}

			var ifres6267 Obj

			if True == ifres6268 {
				ifres6267 = True

			} else {
				ifres6267 = False

			}

			ifres6266 = ifres6267

		} else {
			ifres6266 = False

		}

		if True == ifres6266 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp6287 := Call(__e, ns2_1set, symshen_4foreign_2, tmp6264)

	_ = tmp6287

	tmp6288 := MakeNative(func(__e *ControlFlow) {
		V3189 := __e.Get(1)
		_ = V3189
		tmp6294 := PrimIsPair(V3189)

		var ifres6290 Obj

		if True == tmp6294 {
			tmp6292 := PrimTail(V3189)

			tmp6293 := PrimEqual(Nil, tmp6292)

			var ifres6291 Obj

			if True == tmp6293 {
				ifres6291 = True

			} else {
				ifres6291 = False

			}

			ifres6290 = ifres6291

		} else {
			ifres6290 = False

		}

		if True == ifres6290 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp6295 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp6288)

	_ = tmp6295

	tmp6296 := MakeNative(func(__e *ControlFlow) {
		V3190 := __e.Get(1)
		_ = V3190
		tmp6301 := PrimIsSymbol(V3190)

		if True == tmp6301 {
			tmp6298 := PrimStr(V3190)

			tmp6299 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6298)

			if True == tmp6299 {
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

	tmp6302 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp6296)

	_ = tmp6302

	tmp6303 := MakeNative(func(__e *ControlFlow) {
		V3195 := __e.Get(1)
		_ = V3195
		tmp6333 := PrimIsPair(V3195)

		var ifres6320 Obj

		if True == tmp6333 {
			tmp6331 := PrimHead(V3195)

			tmp6332 := PrimEqual(symprotect, tmp6331)

			var ifres6322 Obj

			if True == tmp6332 {
				tmp6329 := PrimTail(V3195)

				tmp6330 := PrimIsPair(tmp6329)

				var ifres6324 Obj

				if True == tmp6330 {
					tmp6326 := PrimTail(V3195)

					tmp6327 := PrimTail(tmp6326)

					tmp6328 := PrimEqual(Nil, tmp6327)

					var ifres6325 Obj

					if True == tmp6328 {
						ifres6325 = True

					} else {
						ifres6325 = False

					}

					ifres6324 = ifres6325

				} else {
					ifres6324 = False

				}

				var ifres6323 Obj

				if True == ifres6324 {
					ifres6323 = True

				} else {
					ifres6323 = False

				}

				ifres6322 = ifres6323

			} else {
				ifres6322 = False

			}

			var ifres6321 Obj

			if True == ifres6322 {
				ifres6321 = True

			} else {
				ifres6321 = False

			}

			ifres6320 = ifres6321

		} else {
			ifres6320 = False

		}

		if True == ifres6320 {
			__e.Return(False)
			return
		} else {
			tmp6318 := PrimIsPair(V3195)

			var ifres6305 Obj

			if True == tmp6318 {
				tmp6316 := PrimHead(V3195)

				tmp6317 := PrimEqual(symforeign, tmp6316)

				var ifres6307 Obj

				if True == tmp6317 {
					tmp6314 := PrimTail(V3195)

					tmp6315 := PrimIsPair(tmp6314)

					var ifres6309 Obj

					if True == tmp6315 {
						tmp6311 := PrimTail(V3195)

						tmp6312 := PrimTail(tmp6311)

						tmp6313 := PrimEqual(Nil, tmp6312)

						var ifres6310 Obj

						if True == tmp6313 {
							ifres6310 = True

						} else {
							ifres6310 = False

						}

						ifres6309 = ifres6310

					} else {
						ifres6309 = False

					}

					var ifres6308 Obj

					if True == ifres6309 {
						ifres6308 = True

					} else {
						ifres6308 = False

					}

					ifres6307 = ifres6308

				} else {
					ifres6307 = False

				}

				var ifres6306 Obj

				if True == ifres6307 {
					ifres6306 = True

				} else {
					ifres6306 = False

				}

				ifres6305 = ifres6306

			} else {
				ifres6305 = False

			}

			if True == ifres6305 {
				__e.Return(False)
				return
			} else {
				__e.Return(PrimIsPair(V3195))
				return
			}

		}

	}, 1)

	tmp6334 := Call(__e, ns2_1set, symshen_4application_2, tmp6303)

	_ = tmp6334

	tmp6335 := MakeNative(func(__e *ControlFlow) {
		V3200 := __e.Get(1)
		_ = V3200
		V3201 := __e.Get(2)
		_ = V3201
		tmp6343 := PrimEqual(MakeNumber(-1), V3201)

		if True == tmp6343 {
			tmp6341 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V3200)

			if True == tmp6341 {
				tmp6337 := Call(__e, PrimFunc(symexternal), symshen)

				tmp6338 := Call(__e, PrimFunc(symelement_2), V3200, tmp6337)

				tmp6339 := PrimNot(tmp6338)

				if True == tmp6339 {
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
			__e.Return(False)
			return
		}

	}, 2)

	tmp6344 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp6335)

	_ = tmp6344

	tmp6345 := MakeNative(func(__e *ControlFlow) {
		V3202 := __e.Get(1)
		_ = V3202
		tmp6350 := PrimIsSymbol(V3202)

		if True == tmp6350 {
			tmp6347 := PrimIsVariable(V3202)

			tmp6348 := PrimNot(tmp6347)

			if True == tmp6348 {
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

	tmp6351 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp6345)

	_ = tmp6351

	tmp6352 := MakeNative(func(__e *ControlFlow) {
		V3203 := __e.Get(1)
		_ = V3203
		tmp6382 := PrimIsPair(V3203)

		var ifres6373 Obj

		if True == tmp6382 {
			tmp6380 := PrimTail(V3203)

			tmp6381 := PrimIsPair(tmp6380)

			var ifres6375 Obj

			if True == tmp6381 {
				tmp6377 := PrimTail(V3203)

				tmp6378 := PrimTail(tmp6377)

				tmp6379 := PrimEqual(Nil, tmp6378)

				var ifres6376 Obj

				if True == tmp6379 {
					ifres6376 = True

				} else {
					ifres6376 = False

				}

				ifres6375 = ifres6376

			} else {
				ifres6375 = False

			}

			var ifres6374 Obj

			if True == ifres6375 {
				ifres6374 = True

			} else {
				ifres6374 = False

			}

			ifres6373 = ifres6374

		} else {
			ifres6373 = False

		}

		if True == ifres6373 {
			__e.Return(V3203)
			return
		} else {
			tmp6371 := PrimIsPair(V3203)

			var ifres6362 Obj

			if True == tmp6371 {
				tmp6369 := PrimTail(V3203)

				tmp6370 := PrimIsPair(tmp6369)

				var ifres6364 Obj

				if True == tmp6370 {
					tmp6366 := PrimTail(V3203)

					tmp6367 := PrimTail(tmp6366)

					tmp6368 := PrimIsPair(tmp6367)

					var ifres6365 Obj

					if True == tmp6368 {
						ifres6365 = True

					} else {
						ifres6365 = False

					}

					ifres6364 = ifres6365

				} else {
					ifres6364 = False

				}

				var ifres6363 Obj

				if True == ifres6364 {
					ifres6363 = True

				} else {
					ifres6363 = False

				}

				ifres6362 = ifres6363

			} else {
				ifres6362 = False

			}

			if True == ifres6362 {
				tmp6353 := PrimHead(V3203)

				tmp6354 := PrimTail(V3203)

				tmp6355 := PrimHead(tmp6354)

				tmp6356 := PrimCons(tmp6355, Nil)

				tmp6357 := PrimCons(tmp6353, tmp6356)

				tmp6358 := PrimTail(V3203)

				tmp6359 := PrimTail(tmp6358)

				tmp6360 := PrimCons(tmp6357, tmp6359)

				__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6360)
				return

			} else {
				__e.Return(V3203)
				return
			}

		}

	}, 1)

	tmp6383 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp6352)

	_ = tmp6383

	tmp6384 := MakeNative(func(__e *ControlFlow) {
		V3204 := __e.Get(1)
		_ = V3204
		__e.TailApply(PrimFunc(symfn), V3204)
		return
	}, 1)

	tmp6385 := Call(__e, ns2_1set, symfunction, tmp6384)

	_ = tmp6385

	tmp6386 := MakeNative(func(__e *ControlFlow) {
		V3205 := __e.Get(1)
		_ = V3205
		tmp6393 := Call(__e, PrimFunc(symarity), V3205)

		tmp6394 := PrimEqual(tmp6393, MakeNumber(0))

		if True == tmp6394 {
			__e.TailApply(V3205)
			return
		} else {
			tmp6387 := MakeNative(func(__e *ControlFlow) {
				tmp6388 := PrimValue(sym_dproperty_1vector_d)

				__e.TailApply(PrimFunc(symget), V3205, symshen_4lambda_1form, tmp6388)
				return

			}, 0)

			tmp6389 := MakeNative(func(__e *ControlFlow) {
				Z3206 := __e.Get(1)
				_ = Z3206
				tmp6390 := Call(__e, PrimFunc(symshen_4app), V3205, MakeString(" is undefined\n"), symshen_4a)

				tmp6391 := PrimStringConcat(MakeString("fn: "), tmp6390)

				__e.Return(PrimSimpleError(tmp6391))
				return

			}, 1)

			__e.TailApply(try_1catch, tmp6387, tmp6389)
			return

		}

	}, 1)

	tmp6395 := Call(__e, ns2_1set, symfn, tmp6386)

	_ = tmp6395

	tmp6396 := MakeNative(func(__e *ControlFlow) {
		V3209 := __e.Get(1)
		_ = V3209
		tmp6426 := PrimIsPair(V3209)

		var ifres6413 Obj

		if True == tmp6426 {
			tmp6424 := PrimHead(V3209)

			tmp6425 := PrimEqual(symfn, tmp6424)

			var ifres6415 Obj

			if True == tmp6425 {
				tmp6422 := PrimTail(V3209)

				tmp6423 := PrimIsPair(tmp6422)

				var ifres6417 Obj

				if True == tmp6423 {
					tmp6419 := PrimTail(V3209)

					tmp6420 := PrimTail(tmp6419)

					tmp6421 := PrimEqual(Nil, tmp6420)

					var ifres6418 Obj

					if True == tmp6421 {
						ifres6418 = True

					} else {
						ifres6418 = False

					}

					ifres6417 = ifres6418

				} else {
					ifres6417 = False

				}

				var ifres6416 Obj

				if True == ifres6417 {
					ifres6416 = True

				} else {
					ifres6416 = False

				}

				ifres6415 = ifres6416

			} else {
				ifres6415 = False

			}

			var ifres6414 Obj

			if True == ifres6415 {
				ifres6414 = True

			} else {
				ifres6414 = False

			}

			ifres6413 = ifres6414

		} else {
			ifres6413 = False

		}

		if True == ifres6413 {
			__e.Return(True)
			return
		} else {
			tmp6411 := PrimIsPair(V3209)

			var ifres6398 Obj

			if True == tmp6411 {
				tmp6409 := PrimHead(V3209)

				tmp6410 := PrimEqual(symfunction, tmp6409)

				var ifres6400 Obj

				if True == tmp6410 {
					tmp6407 := PrimTail(V3209)

					tmp6408 := PrimIsPair(tmp6407)

					var ifres6402 Obj

					if True == tmp6408 {
						tmp6404 := PrimTail(V3209)

						tmp6405 := PrimTail(tmp6404)

						tmp6406 := PrimEqual(Nil, tmp6405)

						var ifres6403 Obj

						if True == tmp6406 {
							ifres6403 = True

						} else {
							ifres6403 = False

						}

						ifres6402 = ifres6403

					} else {
						ifres6402 = False

					}

					var ifres6401 Obj

					if True == ifres6402 {
						ifres6401 = True

					} else {
						ifres6401 = False

					}

					ifres6400 = ifres6401

				} else {
					ifres6400 = False

				}

				var ifres6399 Obj

				if True == ifres6400 {
					ifres6399 = True

				} else {
					ifres6399 = False

				}

				ifres6398 = ifres6399

			} else {
				ifres6398 = False

			}

			if True == ifres6398 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp6427 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp6396)

	_ = tmp6427

	tmp6428 := MakeNative(func(__e *ControlFlow) {
		V3210 := __e.Get(1)
		_ = V3210
		tmp6469 := PrimIsPair(V3210)

		var ifres6456 Obj

		if True == tmp6469 {
			tmp6467 := PrimHead(V3210)

			tmp6468 := PrimEqual(symfunction, tmp6467)

			var ifres6458 Obj

			if True == tmp6468 {
				tmp6465 := PrimTail(V3210)

				tmp6466 := PrimIsPair(tmp6465)

				var ifres6460 Obj

				if True == tmp6466 {
					tmp6462 := PrimTail(V3210)

					tmp6463 := PrimTail(tmp6462)

					tmp6464 := PrimEqual(Nil, tmp6463)

					var ifres6461 Obj

					if True == tmp6464 {
						ifres6461 = True

					} else {
						ifres6461 = False

					}

					ifres6460 = ifres6461

				} else {
					ifres6460 = False

				}

				var ifres6459 Obj

				if True == ifres6460 {
					ifres6459 = True

				} else {
					ifres6459 = False

				}

				ifres6458 = ifres6459

			} else {
				ifres6458 = False

			}

			var ifres6457 Obj

			if True == ifres6458 {
				ifres6457 = True

			} else {
				ifres6457 = False

			}

			ifres6456 = ifres6457

		} else {
			ifres6456 = False

		}

		if True == ifres6456 {
			tmp6429 := PrimTail(V3210)

			tmp6430 := PrimCons(symfn, tmp6429)

			__e.TailApply(PrimFunc(symshen_4fn_1call), tmp6430)
			return

		} else {
			tmp6454 := PrimIsPair(V3210)

			var ifres6441 Obj

			if True == tmp6454 {
				tmp6452 := PrimHead(V3210)

				tmp6453 := PrimEqual(symfn, tmp6452)

				var ifres6443 Obj

				if True == tmp6453 {
					tmp6450 := PrimTail(V3210)

					tmp6451 := PrimIsPair(tmp6450)

					var ifres6445 Obj

					if True == tmp6451 {
						tmp6447 := PrimTail(V3210)

						tmp6448 := PrimTail(tmp6447)

						tmp6449 := PrimEqual(Nil, tmp6448)

						var ifres6446 Obj

						if True == tmp6449 {
							ifres6446 = True

						} else {
							ifres6446 = False

						}

						ifres6445 = ifres6446

					} else {
						ifres6445 = False

					}

					var ifres6444 Obj

					if True == ifres6445 {
						ifres6444 = True

					} else {
						ifres6444 = False

					}

					ifres6443 = ifres6444

				} else {
					ifres6443 = False

				}

				var ifres6442 Obj

				if True == ifres6443 {
					ifres6442 = True

				} else {
					ifres6442 = False

				}

				ifres6441 = ifres6442

			} else {
				ifres6441 = False

			}

			if True == ifres6441 {
				tmp6431 := MakeNative(func(__e *ControlFlow) {
					W3211 := __e.Get(1)
					_ = W3211
					tmp6436 := PrimEqual(W3211, MakeNumber(-1))

					if True == tmp6436 {
						__e.Return(V3210)
						return
					} else {
						tmp6434 := PrimEqual(W3211, MakeNumber(0))

						if True == tmp6434 {
							__e.Return(PrimTail(V3210))
							return
						} else {
							tmp6432 := PrimTail(V3210)

							__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp6432, W3211)
							return

						}

					}

				}, 1)

				tmp6437 := PrimTail(V3210)

				tmp6438 := PrimHead(tmp6437)

				tmp6439 := Call(__e, PrimFunc(symarity), tmp6438)

				__e.TailApply(tmp6431, tmp6439)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4fn_1call)
				return
			}

		}

	}, 1)

	tmp6470 := Call(__e, ns2_1set, symshen_4fn_1call, tmp6428)

	_ = tmp6470

	tmp6471 := MakeNative(func(__e *ControlFlow) {
		V3212 := __e.Get(1)
		_ = V3212
		V3213 := __e.Get(2)
		_ = V3213
		V3214 := __e.Get(3)
		_ = V3214
		tmp6472 := MakeNative(func(__e *ControlFlow) {
			W3215 := __e.Get(1)
			_ = W3215
			tmp6473 := MakeNative(func(__e *ControlFlow) {
				W3216 := __e.Get(1)
				_ = W3216
				__e.Return(W3215)
				return
			}, 1)

			var ifres6479 Obj

			if True == W3215 {
				tmp6487 := Call(__e, PrimFunc(symshen_4loading_2))

				var ifres6481 Obj

				if True == tmp6487 {
					tmp6483 := PrimCons(sym_1, Nil)

					tmp6484 := PrimCons(sym_7, tmp6483)

					tmp6485 := Call(__e, PrimFunc(symelement_2), V3212, tmp6484)

					tmp6486 := PrimNot(tmp6485)

					var ifres6482 Obj

					if True == tmp6486 {
						ifres6482 = True

					} else {
						ifres6482 = False

					}

					ifres6481 = ifres6482

				} else {
					ifres6481 = False

				}

				var ifres6480 Obj

				if True == ifres6481 {
					ifres6480 = True

				} else {
					ifres6480 = False

				}

				ifres6479 = ifres6480

			} else {
				ifres6479 = False

			}

			var ifres6474 Obj

			if True == ifres6479 {
				tmp6475 := Call(__e, PrimFunc(symshen_4app), V3212, MakeString("\n"), symshen_4a)

				tmp6476 := PrimStringConcat(MakeString("partial application of "), tmp6475)

				tmp6477 := Call(__e, PrimFunc(symstoutput))

				tmp6478 := Call(__e, PrimFunc(sympr), tmp6476, tmp6477)

				ifres6474 = tmp6478

			} else {
				ifres6474 = symshen_4skip

			}

			__e.TailApply(tmp6473, ifres6474)
			return

		}, 1)

		tmp6488 := PrimGreatThan(V3213, V3214)

		__e.TailApply(tmp6472, tmp6488)
		return

	}, 3)

	tmp6489 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp6471)

	_ = tmp6489

	tmp6490 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dloading_2_d))
		return
	}, 0)

	tmp6491 := Call(__e, ns2_1set, symshen_4loading_2, tmp6490)

	_ = tmp6491

	tmp6492 := MakeNative(func(__e *ControlFlow) {
		V3221 := __e.Get(1)
		_ = V3221
		V3222 := __e.Get(2)
		_ = V3222
		V3223 := __e.Get(3)
		_ = V3223
		tmp6510 := PrimEqual(MakeNumber(-1), V3222)

		if True == tmp6510 {
			__e.Return(False)
			return
		} else {
			tmp6493 := MakeNative(func(__e *ControlFlow) {
				W3224 := __e.Get(1)
				_ = W3224
				tmp6494 := MakeNative(func(__e *ControlFlow) {
					W3225 := __e.Get(1)
					_ = W3225
					__e.Return(W3224)
					return
				}, 1)

				var ifres6505 Obj

				if True == W3224 {
					tmp6507 := Call(__e, PrimFunc(symshen_4loading_2))

					var ifres6506 Obj

					if True == tmp6507 {
						ifres6506 = True

					} else {
						ifres6506 = False

					}

					ifres6505 = ifres6506

				} else {
					ifres6505 = False

				}

				var ifres6495 Obj

				if True == ifres6505 {
					tmp6497 := PrimEqual(V3223, MakeNumber(1))

					var ifres6496 Obj

					if True == tmp6497 {
						ifres6496 = MakeString("")

					} else {
						ifres6496 = MakeString("s")

					}

					tmp6498 := Call(__e, PrimFunc(symshen_4app), ifres6496, MakeString("\n"), symshen_4a)

					tmp6499 := PrimStringConcat(MakeString(" argument"), tmp6498)

					tmp6500 := Call(__e, PrimFunc(symshen_4app), V3223, tmp6499, symshen_4a)

					tmp6501 := PrimStringConcat(MakeString(" might not like "), tmp6500)

					tmp6502 := Call(__e, PrimFunc(symshen_4app), V3221, tmp6501, symshen_4a)

					tmp6503 := Call(__e, PrimFunc(symstoutput))

					tmp6504 := Call(__e, PrimFunc(sympr), tmp6502, tmp6503)

					ifres6495 = tmp6504

				} else {
					ifres6495 = symshen_4skip

				}

				__e.TailApply(tmp6494, ifres6495)
				return

			}, 1)

			tmp6508 := PrimLessThan(V3222, V3223)

			__e.TailApply(tmp6493, tmp6508)
			return

		}

	}, 3)

	__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp6492)
	return

}, 0)
