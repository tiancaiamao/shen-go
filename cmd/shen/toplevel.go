package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
	tmp6619 := MakeNative(func(__e *ControlFlow) {
		tmp6620 := Call(__e, PrimFunc(symshen_4credits))

		_ = tmp6620

		__e.TailApply(PrimFunc(symshen_4loop))
		return

	}, 0)

	tmp6621 := Call(__e, ns2_1set, symshen_4repl, tmp6619)

	_ = tmp6621

	tmp6622 := MakeNative(func(__e *ControlFlow) {
		tmp6623 := Call(__e, PrimFunc(symshen_4initialise__environment))

		_ = tmp6623

		tmp6624 := Call(__e, PrimFunc(symshen_4prompt))

		_ = tmp6624

		tmp6625 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
			return
		}, 0)

		tmp6626 := MakeNative(func(__e *ControlFlow) {
			Z5643 := __e.Get(1)
			_ = Z5643
			__e.TailApply(PrimFunc(symshen_4toplevel_1display_1exception), Z5643)
			return
		}, 1)

		tmp6627 := Call(__e, try_1catch, tmp6625, tmp6626)

		_ = tmp6627

		__e.TailApply(PrimFunc(symshen_4loop))
		return

	}, 0)

	tmp6628 := Call(__e, ns2_1set, symshen_4loop, tmp6622)

	_ = tmp6628

	tmp6629 := MakeNative(func(__e *ControlFlow) {
		V5644 := __e.Get(1)
		_ = V5644
		tmp6630 := PrimErrorToString(V5644)

		tmp6631 := Call(__e, PrimFunc(symstoutput))

		tmp6632 := Call(__e, PrimFunc(sympr), tmp6630, tmp6631)

		_ = tmp6632

		__e.TailApply(PrimFunc(symnl), MakeNumber(0))
		return

	}, 1)

	tmp6633 := Call(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp6629)

	_ = tmp6633

	tmp6634 := MakeNative(func(__e *ControlFlow) {
		tmp6635 := Call(__e, PrimFunc(symstoutput))

		tmp6636 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp6635)

		_ = tmp6636

		tmp6637 := PrimValue(sym_dversion_d)

		tmp6638 := PrimValue(sym_dlanguage_d)

		tmp6639 := PrimValue(sym_dimplementation_d)

		tmp6640 := PrimValue(sym_drelease_d)

		tmp6641 := Call(__e, PrimFunc(symshen_4app), tmp6640, MakeString("\n"), symshen_4a)

		tmp6642 := PrimStringConcat(MakeString(" "), tmp6641)

		tmp6643 := Call(__e, PrimFunc(symshen_4app), tmp6639, tmp6642, symshen_4a)

		tmp6644 := PrimStringConcat(MakeString(", platform: "), tmp6643)

		tmp6645 := Call(__e, PrimFunc(symshen_4app), tmp6638, tmp6644, symshen_4a)

		tmp6646 := PrimStringConcat(MakeString(", language: "), tmp6645)

		tmp6647 := Call(__e, PrimFunc(symshen_4app), tmp6637, tmp6646, symshen_4a)

		tmp6648 := PrimStringConcat(MakeString("version: S"), tmp6647)

		tmp6649 := Call(__e, PrimFunc(symstoutput))

		tmp6650 := Call(__e, PrimFunc(sympr), tmp6648, tmp6649)

		_ = tmp6650

		tmp6651 := PrimValue(sym_dport_d)

		tmp6652 := PrimValue(sym_dporters_d)

		tmp6653 := Call(__e, PrimFunc(symshen_4app), tmp6652, MakeString("\n\n"), symshen_4a)

		tmp6654 := PrimStringConcat(MakeString(", ported by "), tmp6653)

		tmp6655 := Call(__e, PrimFunc(symshen_4app), tmp6651, tmp6654, symshen_4a)

		tmp6656 := PrimStringConcat(MakeString("port "), tmp6655)

		tmp6657 := Call(__e, PrimFunc(symstoutput))

		__e.TailApply(PrimFunc(sympr), tmp6656, tmp6657)
		return

	}, 0)

	tmp6658 := Call(__e, ns2_1set, symshen_4credits, tmp6634)

	_ = tmp6658

	tmp6659 := MakeNative(func(__e *ControlFlow) {
		tmp6660 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

		_ = tmp6660

		__e.Return(PrimSet(symshen_4_dinfs_d, MakeNumber(0)))
		return

	}, 0)

	tmp6661 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp6659)

	_ = tmp6661

	tmp6662 := MakeNative(func(__e *ControlFlow) {
		tmp6674 := PrimValue(symshen_4_dtc_d)

		if True == tmp6674 {
			tmp6663 := PrimValue(symshen_4_dhistory_d)

			tmp6664 := Call(__e, PrimFunc(symlength), tmp6663)

			tmp6665 := Call(__e, PrimFunc(symshen_4app), tmp6664, MakeString("+) "), symshen_4a)

			tmp6666 := PrimStringConcat(MakeString("\n("), tmp6665)

			tmp6667 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp6666, tmp6667)
			return

		} else {
			tmp6668 := PrimValue(symshen_4_dhistory_d)

			tmp6669 := Call(__e, PrimFunc(symlength), tmp6668)

			tmp6670 := Call(__e, PrimFunc(symshen_4app), tmp6669, MakeString("-) "), symshen_4a)

			tmp6671 := PrimStringConcat(MakeString("\n("), tmp6670)

			tmp6672 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp6671, tmp6672)
			return

		}

	}, 0)

	tmp6675 := Call(__e, ns2_1set, symshen_4prompt, tmp6662)

	_ = tmp6675

	tmp6676 := MakeNative(func(__e *ControlFlow) {
		tmp6677 := MakeNative(func(__e *ControlFlow) {
			W5645 := __e.Get(1)
			_ = W5645
			tmp6678 := MakeNative(func(__e *ControlFlow) {
				W5646 := __e.Get(1)
				_ = W5646
				tmp6679 := MakeNative(func(__e *ControlFlow) {
					W5647 := __e.Get(1)
					_ = W5647
					tmp6680 := PrimValue(symshen_4_dtc_d)

					__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5646, W5647, tmp6680)
					return

				}, 1)

				tmp6681 := Call(__e, PrimFunc(symshen_4update_1history))

				__e.TailApply(tmp6679, tmp6681)
				return

			}, 1)

			tmp6682 := Call(__e, PrimFunc(symstinput))

			tmp6683 := Call(__e, PrimFunc(symlineread), tmp6682)

			tmp6684 := Call(__e, PrimFunc(symshen_4package_1user_1input), W5645, tmp6683)

			__e.TailApply(tmp6678, tmp6684)
			return

		}, 1)

		tmp6685 := PrimValue(symshen_4_dpackage_d)

		__e.TailApply(tmp6677, tmp6685)
		return

	}, 0)

	tmp6686 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp6676)

	_ = tmp6686

	tmp6687 := MakeNative(func(__e *ControlFlow) {
		V5648 := __e.Get(1)
		_ = V5648
		V5649 := __e.Get(2)
		_ = V5649
		tmp6694 := PrimEqual(symnull, V5648)

		if True == tmp6694 {
			__e.Return(V5649)
			return
		} else {
			tmp6688 := MakeNative(func(__e *ControlFlow) {
				W5650 := __e.Get(1)
				_ = W5650
				tmp6689 := MakeNative(func(__e *ControlFlow) {
					W5651 := __e.Get(1)
					_ = W5651
					tmp6690 := MakeNative(func(__e *ControlFlow) {
						Z5652 := __e.Get(1)
						_ = Z5652
						__e.TailApply(PrimFunc(symshen_4pui_1h), W5650, W5651, Z5652)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp6690, V5649)
					return

				}, 1)

				tmp6691 := Call(__e, PrimFunc(symexternal), V5648)

				__e.TailApply(tmp6689, tmp6691)
				return

			}, 1)

			tmp6692 := PrimStr(V5648)

			__e.TailApply(tmp6688, tmp6692)
			return

		}

	}, 2)

	tmp6695 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp6687)

	_ = tmp6695

	tmp6696 := MakeNative(func(__e *ControlFlow) {
		V5657 := __e.Get(1)
		_ = V5657
		V5658 := __e.Get(2)
		_ = V5658
		V5659 := __e.Get(3)
		_ = V5659
		tmp6737 := PrimIsPair(V5659)

		var ifres6724 Obj

		if True == tmp6737 {
			tmp6735 := PrimHead(V5659)

			tmp6736 := PrimEqual(symfn, tmp6735)

			var ifres6726 Obj

			if True == tmp6736 {
				tmp6733 := PrimTail(V5659)

				tmp6734 := PrimIsPair(tmp6733)

				var ifres6728 Obj

				if True == tmp6734 {
					tmp6730 := PrimTail(V5659)

					tmp6731 := PrimTail(tmp6730)

					tmp6732 := PrimEqual(Nil, tmp6731)

					var ifres6729 Obj

					if True == tmp6732 {
						ifres6729 = True

					} else {
						ifres6729 = False

					}

					ifres6728 = ifres6729

				} else {
					ifres6728 = False

				}

				var ifres6727 Obj

				if True == ifres6728 {
					ifres6727 = True

				} else {
					ifres6727 = False

				}

				ifres6726 = ifres6727

			} else {
				ifres6726 = False

			}

			var ifres6725 Obj

			if True == ifres6726 {
				ifres6725 = True

			} else {
				ifres6725 = False

			}

			ifres6724 = ifres6725

		} else {
			ifres6724 = False

		}

		if True == ifres6724 {
			tmp6702 := PrimTail(V5659)

			tmp6703 := PrimHead(tmp6702)

			tmp6704 := Call(__e, PrimFunc(symshen_4internal_2), tmp6703, V5657, V5658)

			if True == tmp6704 {
				tmp6697 := PrimTail(V5659)

				tmp6698 := PrimHead(tmp6697)

				tmp6699 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5657, tmp6698)

				tmp6700 := PrimCons(tmp6699, Nil)

				__e.Return(PrimCons(symfn, tmp6700))
				return

			} else {
				__e.Return(V5659)
				return
			}

		} else {
			tmp6722 := PrimIsPair(V5659)

			if True == tmp6722 {
				tmp6719 := PrimHead(V5659)

				tmp6720 := Call(__e, PrimFunc(symshen_4internal_2), tmp6719, V5657, V5658)

				if True == tmp6720 {
					tmp6705 := PrimHead(V5659)

					tmp6706 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5657, tmp6705)

					tmp6707 := MakeNative(func(__e *ControlFlow) {
						Z5660 := __e.Get(1)
						_ = Z5660
						__e.TailApply(PrimFunc(symshen_4pui_1h), V5657, V5658, Z5660)
						return
					}, 1)

					tmp6708 := PrimTail(V5659)

					tmp6709 := Call(__e, PrimFunc(symmap), tmp6707, tmp6708)

					__e.Return(PrimCons(tmp6706, tmp6709))
					return

				} else {
					tmp6716 := PrimHead(V5659)

					tmp6717 := PrimIsPair(tmp6716)

					if True == tmp6717 {
						tmp6710 := MakeNative(func(__e *ControlFlow) {
							Z5661 := __e.Get(1)
							_ = Z5661
							__e.TailApply(PrimFunc(symshen_4pui_1h), V5657, V5658, Z5661)
							return
						}, 1)

						__e.TailApply(PrimFunc(symmap), tmp6710, V5659)
						return

					} else {
						tmp6711 := PrimHead(V5659)

						tmp6712 := MakeNative(func(__e *ControlFlow) {
							Z5662 := __e.Get(1)
							_ = Z5662
							__e.TailApply(PrimFunc(symshen_4pui_1h), V5657, V5658, Z5662)
							return
						}, 1)

						tmp6713 := PrimTail(V5659)

						tmp6714 := Call(__e, PrimFunc(symmap), tmp6712, tmp6713)

						__e.Return(PrimCons(tmp6711, tmp6714))
						return

					}

				}

			} else {
				__e.Return(V5659)
				return
			}

		}

	}, 3)

	tmp6738 := Call(__e, ns2_1set, symshen_4pui_1h, tmp6696)

	_ = tmp6738

	tmp6739 := MakeNative(func(__e *ControlFlow) {
		tmp6740 := Call(__e, PrimFunc(symit))

		tmp6741 := Call(__e, PrimFunc(symshen_4trim_1it), tmp6740)

		tmp6742 := PrimValue(symshen_4_dhistory_d)

		tmp6743 := PrimCons(tmp6741, tmp6742)

		__e.Return(PrimSet(symshen_4_dhistory_d, tmp6743))
		return

	}, 0)

	tmp6744 := Call(__e, ns2_1set, symshen_4update_1history, tmp6739)

	_ = tmp6744

	tmp6745 := MakeNative(func(__e *ControlFlow) {
		V5663 := __e.Get(1)
		_ = V5663
		tmp6753 := Call(__e, PrimFunc(symshen_4_7string_2), V5663)

		var ifres6748 Obj

		if True == tmp6753 {
			tmp6750 := Call(__e, PrimFunc(symhdstr), V5663)

			tmp6751 := PrimStringToNumber(tmp6750)

			tmp6752 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp6751)

			var ifres6749 Obj

			if True == tmp6752 {
				ifres6749 = True

			} else {
				ifres6749 = False

			}

			ifres6748 = ifres6749

		} else {
			ifres6748 = False

		}

		if True == ifres6748 {
			tmp6746 := PrimTailString(V5663)

			__e.TailApply(PrimFunc(symshen_4trim_1it), tmp6746)
			return

		} else {
			__e.Return(V5663)
			return
		}

	}, 1)

	tmp6754 := Call(__e, ns2_1set, symshen_4trim_1it, tmp6745)

	_ = tmp6754

	tmp6755 := MakeNative(func(__e *ControlFlow) {
		V5682 := __e.Get(1)
		_ = V5682
		V5683 := __e.Get(2)
		_ = V5683
		V5684 := __e.Get(3)
		_ = V5684
		tmp6885 := PrimIsPair(V5682)

		var ifres6854 Obj

		if True == tmp6885 {
			tmp6883 := PrimTail(V5682)

			tmp6884 := PrimEqual(Nil, tmp6883)

			var ifres6856 Obj

			if True == tmp6884 {
				tmp6882 := PrimIsPair(V5683)

				var ifres6858 Obj

				if True == tmp6882 {
					tmp6880 := PrimHead(V5683)

					tmp6881 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6880)

					var ifres6860 Obj

					if True == tmp6881 {
						tmp6877 := PrimHead(V5683)

						tmp6878 := Call(__e, PrimFunc(symhdstr), tmp6877)

						tmp6879 := PrimEqual(MakeString("!"), tmp6878)

						var ifres6862 Obj

						if True == tmp6879 {
							tmp6874 := PrimHead(V5683)

							tmp6875 := PrimTailString(tmp6874)

							tmp6876 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6875)

							var ifres6864 Obj

							if True == tmp6876 {
								tmp6870 := PrimHead(V5683)

								tmp6871 := PrimTailString(tmp6870)

								tmp6872 := Call(__e, PrimFunc(symhdstr), tmp6871)

								tmp6873 := PrimEqual(MakeString("!"), tmp6872)

								var ifres6866 Obj

								if True == tmp6873 {
									tmp6868 := PrimTail(V5683)

									tmp6869 := PrimIsPair(tmp6868)

									var ifres6867 Obj

									if True == tmp6869 {
										ifres6867 = True

									} else {
										ifres6867 = False

									}

									ifres6866 = ifres6867

								} else {
									ifres6866 = False

								}

								var ifres6865 Obj

								if True == ifres6866 {
									ifres6865 = True

								} else {
									ifres6865 = False

								}

								ifres6864 = ifres6865

							} else {
								ifres6864 = False

							}

							var ifres6863 Obj

							if True == ifres6864 {
								ifres6863 = True

							} else {
								ifres6863 = False

							}

							ifres6862 = ifres6863

						} else {
							ifres6862 = False

						}

						var ifres6861 Obj

						if True == ifres6862 {
							ifres6861 = True

						} else {
							ifres6861 = False

						}

						ifres6860 = ifres6861

					} else {
						ifres6860 = False

					}

					var ifres6859 Obj

					if True == ifres6860 {
						ifres6859 = True

					} else {
						ifres6859 = False

					}

					ifres6858 = ifres6859

				} else {
					ifres6858 = False

				}

				var ifres6857 Obj

				if True == ifres6858 {
					ifres6857 = True

				} else {
					ifres6857 = False

				}

				ifres6856 = ifres6857

			} else {
				ifres6856 = False

			}

			var ifres6855 Obj

			if True == ifres6856 {
				ifres6855 = True

			} else {
				ifres6855 = False

			}

			ifres6854 = ifres6855

		} else {
			ifres6854 = False

		}

		if True == ifres6854 {
			tmp6756 := MakeNative(func(__e *ControlFlow) {
				W5685 := __e.Get(1)
				_ = W5685
				tmp6757 := MakeNative(func(__e *ControlFlow) {
					W5686 := __e.Get(1)
					_ = W5686
					tmp6758 := MakeNative(func(__e *ControlFlow) {
						W5687 := __e.Get(1)
						_ = W5687
						__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5685, W5686, V5684)
						return
					}, 1)

					tmp6759 := PrimTail(V5683)

					tmp6760 := PrimHead(tmp6759)

					tmp6761 := Call(__e, PrimFunc(symshen_4app), tmp6760, MakeString("\n"), symshen_4a)

					tmp6762 := Call(__e, PrimFunc(symstoutput))

					tmp6763 := Call(__e, PrimFunc(sympr), tmp6761, tmp6762)

					__e.TailApply(tmp6758, tmp6763)
					return

				}, 1)

				tmp6764 := PrimTail(V5683)

				tmp6765 := PrimHead(tmp6764)

				tmp6766 := PrimTail(V5683)

				tmp6767 := PrimCons(tmp6765, tmp6766)

				tmp6768 := PrimSet(symshen_4_dhistory_d, tmp6767)

				__e.TailApply(tmp6757, tmp6768)
				return

			}, 1)

			tmp6769 := PrimTail(V5683)

			tmp6770 := PrimHead(tmp6769)

			tmp6771 := Call(__e, PrimFunc(symread_1from_1string), tmp6770)

			__e.TailApply(tmp6756, tmp6771)
			return

		} else {
			tmp6852 := PrimIsPair(V5682)

			var ifres6836 Obj

			if True == tmp6852 {
				tmp6850 := PrimTail(V5682)

				tmp6851 := PrimEqual(Nil, tmp6850)

				var ifres6838 Obj

				if True == tmp6851 {
					tmp6849 := PrimIsPair(V5683)

					var ifres6840 Obj

					if True == tmp6849 {
						tmp6847 := PrimHead(V5683)

						tmp6848 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6847)

						var ifres6842 Obj

						if True == tmp6848 {
							tmp6844 := PrimHead(V5683)

							tmp6845 := Call(__e, PrimFunc(symhdstr), tmp6844)

							tmp6846 := PrimEqual(MakeString("!"), tmp6845)

							var ifres6843 Obj

							if True == tmp6846 {
								ifres6843 = True

							} else {
								ifres6843 = False

							}

							ifres6842 = ifres6843

						} else {
							ifres6842 = False

						}

						var ifres6841 Obj

						if True == ifres6842 {
							ifres6841 = True

						} else {
							ifres6841 = False

						}

						ifres6840 = ifres6841

					} else {
						ifres6840 = False

					}

					var ifres6839 Obj

					if True == ifres6840 {
						ifres6839 = True

					} else {
						ifres6839 = False

					}

					ifres6838 = ifres6839

				} else {
					ifres6838 = False

				}

				var ifres6837 Obj

				if True == ifres6838 {
					ifres6837 = True

				} else {
					ifres6837 = False

				}

				ifres6836 = ifres6837

			} else {
				ifres6836 = False

			}

			if True == ifres6836 {
				tmp6772 := MakeNative(func(__e *ControlFlow) {
					W5688 := __e.Get(1)
					_ = W5688
					tmp6773 := MakeNative(func(__e *ControlFlow) {
						W5689 := __e.Get(1)
						_ = W5689
						tmp6774 := MakeNative(func(__e *ControlFlow) {
							W5690 := __e.Get(1)
							_ = W5690
							tmp6775 := MakeNative(func(__e *ControlFlow) {
								W5691 := __e.Get(1)
								_ = W5691
								tmp6776 := MakeNative(func(__e *ControlFlow) {
									W5692 := __e.Get(1)
									_ = W5692
									__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5691, W5692, V5684)
									return
								}, 1)

								tmp6777 := PrimTail(V5683)

								tmp6778 := PrimCons(W5689, tmp6777)

								tmp6779 := PrimSet(symshen_4_dhistory_d, tmp6778)

								__e.TailApply(tmp6776, tmp6779)
								return

							}, 1)

							tmp6780 := Call(__e, PrimFunc(symread_1from_1string), W5689)

							__e.TailApply(tmp6775, tmp6780)
							return

						}, 1)

						tmp6781 := Call(__e, PrimFunc(symshen_4app), W5689, MakeString("\n"), symshen_4a)

						tmp6782 := Call(__e, PrimFunc(symstoutput))

						tmp6783 := Call(__e, PrimFunc(sympr), tmp6781, tmp6782)

						__e.TailApply(tmp6774, tmp6783)
						return

					}, 1)

					tmp6784 := PrimHead(V5683)

					tmp6785 := PrimTailString(tmp6784)

					tmp6786 := PrimTail(V5683)

					tmp6787 := Call(__e, PrimFunc(symshen_4use_1history), W5688, tmp6785, tmp6786)

					__e.TailApply(tmp6773, tmp6787)
					return

				}, 1)

				tmp6793 := PrimHead(V5683)

				tmp6794 := PrimTailString(tmp6793)

				tmp6795 := PrimEqual(tmp6794, MakeString(""))

				var ifres6788 Obj

				if True == tmp6795 {
					ifres6788 = Nil

				} else {
					tmp6789 := PrimHead(V5683)

					tmp6790 := PrimTailString(tmp6789)

					tmp6791 := Call(__e, PrimFunc(symread_1from_1string), tmp6790)

					tmp6792 := PrimHead(tmp6791)

					ifres6788 = tmp6792

				}

				__e.TailApply(tmp6772, ifres6788)
				return

			} else {
				tmp6834 := PrimIsPair(V5682)

				var ifres6818 Obj

				if True == tmp6834 {
					tmp6832 := PrimTail(V5682)

					tmp6833 := PrimEqual(Nil, tmp6832)

					var ifres6820 Obj

					if True == tmp6833 {
						tmp6831 := PrimIsPair(V5683)

						var ifres6822 Obj

						if True == tmp6831 {
							tmp6829 := PrimHead(V5683)

							tmp6830 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6829)

							var ifres6824 Obj

							if True == tmp6830 {
								tmp6826 := PrimHead(V5683)

								tmp6827 := Call(__e, PrimFunc(symhdstr), tmp6826)

								tmp6828 := PrimEqual(MakeString("%"), tmp6827)

								var ifres6825 Obj

								if True == tmp6828 {
									ifres6825 = True

								} else {
									ifres6825 = False

								}

								ifres6824 = ifres6825

							} else {
								ifres6824 = False

							}

							var ifres6823 Obj

							if True == ifres6824 {
								ifres6823 = True

							} else {
								ifres6823 = False

							}

							ifres6822 = ifres6823

						} else {
							ifres6822 = False

						}

						var ifres6821 Obj

						if True == ifres6822 {
							ifres6821 = True

						} else {
							ifres6821 = False

						}

						ifres6820 = ifres6821

					} else {
						ifres6820 = False

					}

					var ifres6819 Obj

					if True == ifres6820 {
						ifres6819 = True

					} else {
						ifres6819 = False

					}

					ifres6818 = ifres6819

				} else {
					ifres6818 = False

				}

				if True == ifres6818 {
					tmp6796 := MakeNative(func(__e *ControlFlow) {
						W5693 := __e.Get(1)
						_ = W5693
						tmp6797 := MakeNative(func(__e *ControlFlow) {
							W5694 := __e.Get(1)
							_ = W5694
							tmp6798 := MakeNative(func(__e *ControlFlow) {
								W5695 := __e.Get(1)
								_ = W5695
								__e.TailApply(PrimFunc(symabort))
								return
							}, 1)

							tmp6799 := PrimTail(V5683)

							tmp6800 := PrimSet(symshen_4_dhistory_d, tmp6799)

							__e.TailApply(tmp6798, tmp6800)
							return

						}, 1)

						tmp6801 := PrimHead(V5683)

						tmp6802 := PrimTailString(tmp6801)

						tmp6803 := PrimTail(V5683)

						tmp6804 := Call(__e, PrimFunc(symshen_4peek_1history), W5693, tmp6802, tmp6803)

						__e.TailApply(tmp6797, tmp6804)
						return

					}, 1)

					tmp6810 := PrimHead(V5683)

					tmp6811 := PrimTailString(tmp6810)

					tmp6812 := PrimEqual(tmp6811, MakeString(""))

					var ifres6805 Obj

					if True == tmp6812 {
						ifres6805 = Nil

					} else {
						tmp6806 := PrimHead(V5683)

						tmp6807 := PrimTailString(tmp6806)

						tmp6808 := Call(__e, PrimFunc(symread_1from_1string), tmp6807)

						tmp6809 := PrimHead(tmp6808)

						ifres6805 = tmp6809

					}

					__e.TailApply(tmp6796, ifres6805)
					return

				} else {
					tmp6816 := PrimEqual(True, V5684)

					if True == tmp6816 {
						__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5682)
						return
					} else {
						tmp6814 := PrimEqual(False, V5684)

						if True == tmp6814 {
							__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5682)
							return
						} else {
							__e.Return(PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread")))
							return
						}

					}

				}

			}

		}

	}, 3)

	tmp6886 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp6755)

	_ = tmp6886

	tmp6887 := MakeNative(func(__e *ControlFlow) {
		V5696 := __e.Get(1)
		_ = V5696
		V5697 := __e.Get(2)
		_ = V5697
		V5698 := __e.Get(3)
		_ = V5698
		tmp6893 := PrimIsInteger(V5696)

		if True == tmp6893 {
			tmp6888 := PrimNumberAdd(MakeNumber(1), V5696)

			tmp6889 := Call(__e, PrimFunc(symreverse), V5698)

			__e.TailApply(PrimFunc(symnth), tmp6888, tmp6889)
			return

		} else {
			tmp6891 := PrimIsSymbol(V5696)

			if True == tmp6891 {
				__e.TailApply(PrimFunc(symshen_4string_1match), V5697, V5698)
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
				return
			}

		}

	}, 3)

	tmp6894 := Call(__e, ns2_1set, symshen_4use_1history, tmp6887)

	_ = tmp6894

	tmp6895 := MakeNative(func(__e *ControlFlow) {
		V5699 := __e.Get(1)
		_ = V5699
		V5700 := __e.Get(2)
		_ = V5700
		V5701 := __e.Get(3)
		_ = V5701
		tmp6909 := PrimIsInteger(V5699)

		if True == tmp6909 {
			tmp6896 := PrimNumberAdd(MakeNumber(1), V5699)

			tmp6897 := Call(__e, PrimFunc(symreverse), V5701)

			tmp6898 := Call(__e, PrimFunc(symnth), tmp6896, tmp6897)

			tmp6899 := Call(__e, PrimFunc(symshen_4app), tmp6898, MakeString(""), symshen_4a)

			tmp6900 := PrimStringConcat(MakeString("\n"), tmp6899)

			tmp6901 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp6900, tmp6901)
			return

		} else {
			tmp6907 := PrimEqual(V5700, MakeString(""))

			var ifres6904 Obj

			if True == tmp6907 {
				ifres6904 = True

			} else {
				tmp6906 := PrimIsSymbol(V5699)

				var ifres6905 Obj

				if True == tmp6906 {
					ifres6905 = True

				} else {
					ifres6905 = False

				}

				ifres6904 = ifres6905

			}

			if True == ifres6904 {
				tmp6902 := Call(__e, PrimFunc(symreverse), V5701)

				__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5700, tmp6902)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
				return
			}

		}

	}, 3)

	tmp6910 := Call(__e, ns2_1set, symshen_4peek_1history, tmp6895)

	_ = tmp6910

	tmp6911 := MakeNative(func(__e *ControlFlow) {
		V5711 := __e.Get(1)
		_ = V5711
		V5712 := __e.Get(2)
		_ = V5712
		tmp6922 := PrimEqual(Nil, V5712)

		if True == tmp6922 {
			__e.Return(PrimSimpleError(MakeString("\ninput not found")))
			return
		} else {
			tmp6920 := PrimIsPair(V5712)

			var ifres6916 Obj

			if True == tmp6920 {
				tmp6918 := PrimHead(V5712)

				tmp6919 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5711, tmp6918)

				var ifres6917 Obj

				if True == tmp6919 {
					ifres6917 = True

				} else {
					ifres6917 = False

				}

				ifres6916 = ifres6917

			} else {
				ifres6916 = False

			}

			if True == ifres6916 {
				__e.Return(PrimHead(V5712))
				return
			} else {
				tmp6914 := PrimIsPair(V5712)

				if True == tmp6914 {
					tmp6912 := PrimTail(V5712)

					__e.TailApply(PrimFunc(symshen_4string_1match), V5711, tmp6912)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
					return
				}

			}

		}

	}, 2)

	tmp6923 := Call(__e, ns2_1set, symshen_4string_1match, tmp6911)

	_ = tmp6923

	tmp6924 := MakeNative(func(__e *ControlFlow) {
		V5720 := __e.Get(1)
		_ = V5720
		V5721 := __e.Get(2)
		_ = V5721
		tmp6961 := PrimEqual(MakeString(""), V5720)

		if True == tmp6961 {
			__e.Return(True)
			return
		} else {
			tmp6959 := Call(__e, PrimFunc(symshen_4_7string_2), V5720)

			var ifres6954 Obj

			if True == tmp6959 {
				tmp6956 := Call(__e, PrimFunc(symhdstr), V5720)

				tmp6957 := PrimStringToNumber(tmp6956)

				tmp6958 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp6957)

				var ifres6955 Obj

				if True == tmp6958 {
					ifres6955 = True

				} else {
					ifres6955 = False

				}

				ifres6954 = ifres6955

			} else {
				ifres6954 = False

			}

			if True == ifres6954 {
				tmp6925 := PrimTailString(V5720)

				__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp6925, V5721)
				return

			} else {
				tmp6952 := Call(__e, PrimFunc(symshen_4_7string_2), V5721)

				var ifres6947 Obj

				if True == tmp6952 {
					tmp6949 := Call(__e, PrimFunc(symhdstr), V5721)

					tmp6950 := PrimStringToNumber(tmp6949)

					tmp6951 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp6950)

					var ifres6948 Obj

					if True == tmp6951 {
						ifres6948 = True

					} else {
						ifres6948 = False

					}

					ifres6947 = ifres6948

				} else {
					ifres6947 = False

				}

				if True == ifres6947 {
					tmp6926 := PrimTailString(V5721)

					__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5720, tmp6926)
					return

				} else {
					tmp6945 := Call(__e, PrimFunc(symshen_4_7string_2), V5721)

					var ifres6941 Obj

					if True == tmp6945 {
						tmp6943 := Call(__e, PrimFunc(symhdstr), V5721)

						tmp6944 := PrimEqual(MakeString("("), tmp6943)

						var ifres6942 Obj

						if True == tmp6944 {
							ifres6942 = True

						} else {
							ifres6942 = False

						}

						ifres6941 = ifres6942

					} else {
						ifres6941 = False

					}

					if True == ifres6941 {
						tmp6927 := PrimTailString(V5721)

						__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5720, tmp6927)
						return

					} else {
						tmp6939 := Call(__e, PrimFunc(symshen_4_7string_2), V5720)

						var ifres6931 Obj

						if True == tmp6939 {
							tmp6938 := Call(__e, PrimFunc(symshen_4_7string_2), V5721)

							var ifres6933 Obj

							if True == tmp6938 {
								tmp6935 := Call(__e, PrimFunc(symhdstr), V5720)

								tmp6936 := Call(__e, PrimFunc(symhdstr), V5721)

								tmp6937 := PrimEqual(tmp6935, tmp6936)

								var ifres6934 Obj

								if True == tmp6937 {
									ifres6934 = True

								} else {
									ifres6934 = False

								}

								ifres6933 = ifres6934

							} else {
								ifres6933 = False

							}

							var ifres6932 Obj

							if True == ifres6933 {
								ifres6932 = True

							} else {
								ifres6932 = False

							}

							ifres6931 = ifres6932

						} else {
							ifres6931 = False

						}

						if True == ifres6931 {
							tmp6928 := PrimTailString(V5720)

							tmp6929 := PrimTailString(V5721)

							__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp6928, tmp6929)
							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}

	}, 2)

	tmp6962 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp6924)

	_ = tmp6962

	tmp6963 := MakeNative(func(__e *ControlFlow) {
		V5732 := __e.Get(1)
		_ = V5732
		V5733 := __e.Get(2)
		_ = V5733
		V5734 := __e.Get(3)
		_ = V5734
		tmp6978 := PrimEqual(Nil, V5734)

		if True == tmp6978 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp6976 := PrimIsPair(V5734)

			if True == tmp6976 {
				tmp6971 := PrimHead(V5734)

				tmp6972 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5733, tmp6971)

				var ifres6964 Obj

				if True == tmp6972 {
					tmp6965 := PrimHead(V5734)

					tmp6966 := Call(__e, PrimFunc(symshen_4app), tmp6965, MakeString("\n"), symshen_4a)

					tmp6967 := PrimStringConcat(MakeString(". "), tmp6966)

					tmp6968 := Call(__e, PrimFunc(symshen_4app), V5732, tmp6967, symshen_4a)

					tmp6969 := Call(__e, PrimFunc(symstoutput))

					tmp6970 := Call(__e, PrimFunc(sympr), tmp6968, tmp6969)

					ifres6964 = tmp6970

				} else {
					ifres6964 = symshen_4skip

				}

				_ = ifres6964

				tmp6973 := PrimNumberAdd(V5732, MakeNumber(1))

				tmp6974 := PrimTail(V5734)

				__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp6973, V5733, tmp6974)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
				return
			}

		}

	}, 3)

	__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp6963)
	return

}, 0)
