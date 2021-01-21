package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionFeaturesMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

	tmp70325 := MakeNative(func(__e *ControlFlow) {
		V4885 := __e.Get(1)
		_ = V4885
		tmp70842 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp70843 := Call(__e, tmp70842, V4885)

		var ifres70830 Obj

		if True == tmp70843 {
			tmp70838 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp70839 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp70840 := Call(__e, tmp70839, V4885)

			tmp70841 := Call(__e, tmp70838, symshen_4x_4features_4cond_1expand, tmp70840)

			var ifres70832 Obj

			if True == tmp70841 {
				tmp70834 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp70835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70836 := Call(__e, tmp70835, V4885)

				tmp70837 := Call(__e, tmp70834, Nil, tmp70836)

				var ifres70833 Obj

				if True == tmp70837 {
					ifres70833 = True

				} else {
					ifres70833 = False

				}

				ifres70832 = ifres70833

			} else {
				ifres70832 = False

			}

			var ifres70831 Obj

			if True == ifres70832 {
				ifres70831 = True

			} else {
				ifres70831 = False

			}

			ifres70830 = ifres70831

		} else {
			ifres70830 = False

		}

		if True == ifres70830 {
			tmp70829 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp70829, MakeString("Unfulfilled shen.x.features.cond-expand clause."))
			return

		} else {
			tmp70827 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp70828 := Call(__e, tmp70827, V4885)

			var ifres70789 Obj

			if True == tmp70828 {
				tmp70823 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp70824 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp70825 := Call(__e, tmp70824, V4885)

				tmp70826 := Call(__e, tmp70823, symshen_4x_4features_4cond_1expand, tmp70825)

				var ifres70791 Obj

				if True == tmp70826 {
					tmp70819 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp70820 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70821 := Call(__e, tmp70820, V4885)

					tmp70822 := Call(__e, tmp70819, tmp70821)

					var ifres70793 Obj

					if True == tmp70822 {
						tmp70813 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp70814 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp70815 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70816 := Call(__e, tmp70815, V4885)

						tmp70817 := Call(__e, tmp70814, tmp70816)

						tmp70818 := Call(__e, tmp70813, True, tmp70817)

						var ifres70795 Obj

						if True == tmp70818 {
							tmp70807 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp70808 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70809 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70810 := Call(__e, tmp70809, V4885)

							tmp70811 := Call(__e, tmp70808, tmp70810)

							tmp70812 := Call(__e, tmp70807, tmp70811)

							var ifres70797 Obj

							if True == tmp70812 {
								tmp70799 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp70800 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70801 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70802 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70803 := Call(__e, tmp70802, V4885)

								tmp70804 := Call(__e, tmp70801, tmp70803)

								tmp70805 := Call(__e, tmp70800, tmp70804)

								tmp70806 := Call(__e, tmp70799, Nil, tmp70805)

								var ifres70798 Obj

								if True == tmp70806 {
									ifres70798 = True

								} else {
									ifres70798 = False

								}

								ifres70797 = ifres70798

							} else {
								ifres70797 = False

							}

							var ifres70796 Obj

							if True == ifres70797 {
								ifres70796 = True

							} else {
								ifres70796 = False

							}

							ifres70795 = ifres70796

						} else {
							ifres70795 = False

						}

						var ifres70794 Obj

						if True == ifres70795 {
							ifres70794 = True

						} else {
							ifres70794 = False

						}

						ifres70793 = ifres70794

					} else {
						ifres70793 = False

					}

					var ifres70792 Obj

					if True == ifres70793 {
						ifres70792 = True

					} else {
						ifres70792 = False

					}

					ifres70791 = ifres70792

				} else {
					ifres70791 = False

				}

				var ifres70790 Obj

				if True == ifres70791 {
					ifres70790 = True

				} else {
					ifres70790 = False

				}

				ifres70789 = ifres70790

			} else {
				ifres70789 = False

			}

			if True == ifres70789 {
				tmp70784 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp70785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70786 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70787 := Call(__e, tmp70786, V4885)

				tmp70788 := Call(__e, tmp70785, tmp70787)

				__e.TailApply(tmp70784, tmp70788)
				return

			} else {
				tmp70782 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp70783 := Call(__e, tmp70782, V4885)

				var ifres70734 Obj

				if True == tmp70783 {
					tmp70778 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp70779 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp70780 := Call(__e, tmp70779, V4885)

					tmp70781 := Call(__e, tmp70778, symshen_4x_4features_4cond_1expand, tmp70780)

					var ifres70736 Obj

					if True == tmp70781 {
						tmp70774 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp70775 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70776 := Call(__e, tmp70775, V4885)

						tmp70777 := Call(__e, tmp70774, tmp70776)

						var ifres70738 Obj

						if True == tmp70777 {
							tmp70768 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp70769 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp70770 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70771 := Call(__e, tmp70770, V4885)

							tmp70772 := Call(__e, tmp70769, tmp70771)

							tmp70773 := Call(__e, tmp70768, tmp70772)

							var ifres70740 Obj

							if True == tmp70773 {
								tmp70760 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp70761 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70762 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70764 := Call(__e, tmp70763, V4885)

								tmp70765 := Call(__e, tmp70762, tmp70764)

								tmp70766 := Call(__e, tmp70761, tmp70765)

								tmp70767 := Call(__e, tmp70760, symand, tmp70766)

								var ifres70742 Obj

								if True == tmp70767 {
									tmp70752 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp70753 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70754 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70755 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70756 := Call(__e, tmp70755, V4885)

									tmp70757 := Call(__e, tmp70754, tmp70756)

									tmp70758 := Call(__e, tmp70753, tmp70757)

									tmp70759 := Call(__e, tmp70752, Nil, tmp70758)

									var ifres70744 Obj

									if True == tmp70759 {
										tmp70746 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp70747 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70748 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70749 := Call(__e, tmp70748, V4885)

										tmp70750 := Call(__e, tmp70747, tmp70749)

										tmp70751 := Call(__e, tmp70746, tmp70750)

										var ifres70745 Obj

										if True == tmp70751 {
											ifres70745 = True

										} else {
											ifres70745 = False

										}

										ifres70744 = ifres70745

									} else {
										ifres70744 = False

									}

									var ifres70743 Obj

									if True == ifres70744 {
										ifres70743 = True

									} else {
										ifres70743 = False

									}

									ifres70742 = ifres70743

								} else {
									ifres70742 = False

								}

								var ifres70741 Obj

								if True == ifres70742 {
									ifres70741 = True

								} else {
									ifres70741 = False

								}

								ifres70740 = ifres70741

							} else {
								ifres70740 = False

							}

							var ifres70739 Obj

							if True == ifres70740 {
								ifres70739 = True

							} else {
								ifres70739 = False

							}

							ifres70738 = ifres70739

						} else {
							ifres70738 = False

						}

						var ifres70737 Obj

						if True == ifres70738 {
							ifres70737 = True

						} else {
							ifres70737 = False

						}

						ifres70736 = ifres70737

					} else {
						ifres70736 = False

					}

					var ifres70735 Obj

					if True == ifres70736 {
						ifres70735 = True

					} else {
						ifres70735 = False

					}

					ifres70734 = ifres70735

				} else {
					ifres70734 = False

				}

				if True == ifres70734 {
					tmp70729 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp70730 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70731 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp70732 := Call(__e, tmp70731, V4885)

					tmp70733 := Call(__e, tmp70730, tmp70732)

					__e.TailApply(tmp70729, tmp70733)
					return

				} else {
					tmp70727 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp70728 := Call(__e, tmp70727, V4885)

					var ifres70679 Obj

					if True == tmp70728 {
						tmp70723 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp70724 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp70725 := Call(__e, tmp70724, V4885)

						tmp70726 := Call(__e, tmp70723, symshen_4x_4features_4cond_1expand, tmp70725)

						var ifres70681 Obj

						if True == tmp70726 {
							tmp70719 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp70720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70721 := Call(__e, tmp70720, V4885)

							tmp70722 := Call(__e, tmp70719, tmp70721)

							var ifres70683 Obj

							if True == tmp70722 {
								tmp70713 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp70714 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70715 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70716 := Call(__e, tmp70715, V4885)

								tmp70717 := Call(__e, tmp70714, tmp70716)

								tmp70718 := Call(__e, tmp70713, tmp70717)

								var ifres70685 Obj

								if True == tmp70718 {
									tmp70705 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp70706 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70707 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70708 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70709 := Call(__e, tmp70708, V4885)

									tmp70710 := Call(__e, tmp70707, tmp70709)

									tmp70711 := Call(__e, tmp70706, tmp70710)

									tmp70712 := Call(__e, tmp70705, symand, tmp70711)

									var ifres70687 Obj

									if True == tmp70712 {
										tmp70697 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp70698 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70699 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70701 := Call(__e, tmp70700, V4885)

										tmp70702 := Call(__e, tmp70699, tmp70701)

										tmp70703 := Call(__e, tmp70698, tmp70702)

										tmp70704 := Call(__e, tmp70697, tmp70703)

										var ifres70689 Obj

										if True == tmp70704 {
											tmp70691 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp70692 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70693 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70694 := Call(__e, tmp70693, V4885)

											tmp70695 := Call(__e, tmp70692, tmp70694)

											tmp70696 := Call(__e, tmp70691, tmp70695)

											var ifres70690 Obj

											if True == tmp70696 {
												ifres70690 = True

											} else {
												ifres70690 = False

											}

											ifres70689 = ifres70690

										} else {
											ifres70689 = False

										}

										var ifres70688 Obj

										if True == ifres70689 {
											ifres70688 = True

										} else {
											ifres70688 = False

										}

										ifres70687 = ifres70688

									} else {
										ifres70687 = False

									}

									var ifres70686 Obj

									if True == ifres70687 {
										ifres70686 = True

									} else {
										ifres70686 = False

									}

									ifres70685 = ifres70686

								} else {
									ifres70685 = False

								}

								var ifres70684 Obj

								if True == ifres70685 {
									ifres70684 = True

								} else {
									ifres70684 = False

								}

								ifres70683 = ifres70684

							} else {
								ifres70683 = False

							}

							var ifres70682 Obj

							if True == ifres70683 {
								ifres70682 = True

							} else {
								ifres70682 = False

							}

							ifres70681 = ifres70682

						} else {
							ifres70681 = False

						}

						var ifres70680 Obj

						if True == ifres70681 {
							ifres70680 = True

						} else {
							ifres70680 = False

						}

						ifres70679 = ifres70680

					} else {
						ifres70679 = False

					}

					if True == ifres70679 {
						tmp70642 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70643 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70644 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp70645 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70646 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp70647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70648 := Call(__e, tmp70647, V4885)

						tmp70649 := Call(__e, tmp70646, tmp70648)

						tmp70650 := Call(__e, tmp70645, tmp70649)

						tmp70651 := Call(__e, tmp70644, tmp70650)

						tmp70652 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70653 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70654 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70655 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp70656 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70657 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70658 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp70659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70660 := Call(__e, tmp70659, V4885)

						tmp70661 := Call(__e, tmp70658, tmp70660)

						tmp70662 := Call(__e, tmp70657, tmp70661)

						tmp70663 := Call(__e, tmp70656, tmp70662)

						tmp70664 := Call(__e, tmp70655, symand, tmp70663)

						tmp70665 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70666 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70667 := Call(__e, tmp70666, V4885)

						tmp70668 := Call(__e, tmp70665, tmp70667)

						tmp70669 := Call(__e, tmp70654, tmp70664, tmp70668)

						tmp70670 := Call(__e, tmp70653, symshen_4x_4features_4cond_1expand, tmp70669)

						tmp70671 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70672 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70673 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp70674 := Call(__e, tmp70673, V4885)

						tmp70675 := Call(__e, tmp70672, tmp70674)

						tmp70676 := Call(__e, tmp70671, tmp70675)

						tmp70677 := Call(__e, tmp70652, tmp70670, tmp70676)

						tmp70678 := Call(__e, tmp70643, tmp70651, tmp70677)

						__e.TailApply(tmp70642, symshen_4x_4features_4cond_1expand, tmp70678)
						return

					} else {
						tmp70640 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp70641 := Call(__e, tmp70640, V4885)

						var ifres70592 Obj

						if True == tmp70641 {
							tmp70636 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp70637 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp70638 := Call(__e, tmp70637, V4885)

							tmp70639 := Call(__e, tmp70636, symshen_4x_4features_4cond_1expand, tmp70638)

							var ifres70594 Obj

							if True == tmp70639 {
								tmp70632 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp70633 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70634 := Call(__e, tmp70633, V4885)

								tmp70635 := Call(__e, tmp70632, tmp70634)

								var ifres70596 Obj

								if True == tmp70635 {
									tmp70626 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp70627 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70628 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70629 := Call(__e, tmp70628, V4885)

									tmp70630 := Call(__e, tmp70627, tmp70629)

									tmp70631 := Call(__e, tmp70626, tmp70630)

									var ifres70598 Obj

									if True == tmp70631 {
										tmp70618 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp70619 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70620 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70621 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70622 := Call(__e, tmp70621, V4885)

										tmp70623 := Call(__e, tmp70620, tmp70622)

										tmp70624 := Call(__e, tmp70619, tmp70623)

										tmp70625 := Call(__e, tmp70618, symor, tmp70624)

										var ifres70600 Obj

										if True == tmp70625 {
											tmp70610 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp70611 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70612 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp70613 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70614 := Call(__e, tmp70613, V4885)

											tmp70615 := Call(__e, tmp70612, tmp70614)

											tmp70616 := Call(__e, tmp70611, tmp70615)

											tmp70617 := Call(__e, tmp70610, Nil, tmp70616)

											var ifres70602 Obj

											if True == tmp70617 {
												tmp70604 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp70605 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70606 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70607 := Call(__e, tmp70606, V4885)

												tmp70608 := Call(__e, tmp70605, tmp70607)

												tmp70609 := Call(__e, tmp70604, tmp70608)

												var ifres70603 Obj

												if True == tmp70609 {
													ifres70603 = True

												} else {
													ifres70603 = False

												}

												ifres70602 = ifres70603

											} else {
												ifres70602 = False

											}

											var ifres70601 Obj

											if True == ifres70602 {
												ifres70601 = True

											} else {
												ifres70601 = False

											}

											ifres70600 = ifres70601

										} else {
											ifres70600 = False

										}

										var ifres70599 Obj

										if True == ifres70600 {
											ifres70599 = True

										} else {
											ifres70599 = False

										}

										ifres70598 = ifres70599

									} else {
										ifres70598 = False

									}

									var ifres70597 Obj

									if True == ifres70598 {
										ifres70597 = True

									} else {
										ifres70597 = False

									}

									ifres70596 = ifres70597

								} else {
									ifres70596 = False

								}

								var ifres70595 Obj

								if True == ifres70596 {
									ifres70595 = True

								} else {
									ifres70595 = False

								}

								ifres70594 = ifres70595

							} else {
								ifres70594 = False

							}

							var ifres70593 Obj

							if True == ifres70594 {
								ifres70593 = True

							} else {
								ifres70593 = False

							}

							ifres70592 = ifres70593

						} else {
							ifres70592 = False

						}

						if True == ifres70592 {
							tmp70585 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp70586 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70587 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70588 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp70589 := Call(__e, tmp70588, V4885)

							tmp70590 := Call(__e, tmp70587, tmp70589)

							tmp70591 := Call(__e, tmp70586, tmp70590)

							__e.TailApply(tmp70585, symshen_4x_4features_4cond_1expand, tmp70591)
							return

						} else {
							tmp70583 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp70584 := Call(__e, tmp70583, V4885)

							var ifres70535 Obj

							if True == tmp70584 {
								tmp70579 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp70580 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70581 := Call(__e, tmp70580, V4885)

								tmp70582 := Call(__e, tmp70579, symshen_4x_4features_4cond_1expand, tmp70581)

								var ifres70537 Obj

								if True == tmp70582 {
									tmp70575 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp70576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70577 := Call(__e, tmp70576, V4885)

									tmp70578 := Call(__e, tmp70575, tmp70577)

									var ifres70539 Obj

									if True == tmp70578 {
										tmp70569 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp70570 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70571 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70572 := Call(__e, tmp70571, V4885)

										tmp70573 := Call(__e, tmp70570, tmp70572)

										tmp70574 := Call(__e, tmp70569, tmp70573)

										var ifres70541 Obj

										if True == tmp70574 {
											tmp70561 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp70562 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp70563 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp70564 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70565 := Call(__e, tmp70564, V4885)

											tmp70566 := Call(__e, tmp70563, tmp70565)

											tmp70567 := Call(__e, tmp70562, tmp70566)

											tmp70568 := Call(__e, tmp70561, symor, tmp70567)

											var ifres70543 Obj

											if True == tmp70568 {
												tmp70553 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp70554 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70555 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp70556 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70557 := Call(__e, tmp70556, V4885)

												tmp70558 := Call(__e, tmp70555, tmp70557)

												tmp70559 := Call(__e, tmp70554, tmp70558)

												tmp70560 := Call(__e, tmp70553, tmp70559)

												var ifres70545 Obj

												if True == tmp70560 {
													tmp70547 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp70548 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70549 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70550 := Call(__e, tmp70549, V4885)

													tmp70551 := Call(__e, tmp70548, tmp70550)

													tmp70552 := Call(__e, tmp70547, tmp70551)

													var ifres70546 Obj

													if True == tmp70552 {
														ifres70546 = True

													} else {
														ifres70546 = False

													}

													ifres70545 = ifres70546

												} else {
													ifres70545 = False

												}

												var ifres70544 Obj

												if True == ifres70545 {
													ifres70544 = True

												} else {
													ifres70544 = False

												}

												ifres70543 = ifres70544

											} else {
												ifres70543 = False

											}

											var ifres70542 Obj

											if True == ifres70543 {
												ifres70542 = True

											} else {
												ifres70542 = False

											}

											ifres70541 = ifres70542

										} else {
											ifres70541 = False

										}

										var ifres70540 Obj

										if True == ifres70541 {
											ifres70540 = True

										} else {
											ifres70540 = False

										}

										ifres70539 = ifres70540

									} else {
										ifres70539 = False

									}

									var ifres70538 Obj

									if True == ifres70539 {
										ifres70538 = True

									} else {
										ifres70538 = False

									}

									ifres70537 = ifres70538

								} else {
									ifres70537 = False

								}

								var ifres70536 Obj

								if True == ifres70537 {
									ifres70536 = True

								} else {
									ifres70536 = False

								}

								ifres70535 = ifres70536

							} else {
								ifres70535 = False

							}

							if True == ifres70535 {
								tmp70494 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70495 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70496 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70497 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70498 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70499 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70500 := Call(__e, tmp70499, V4885)

								tmp70501 := Call(__e, tmp70498, tmp70500)

								tmp70502 := Call(__e, tmp70497, tmp70501)

								tmp70503 := Call(__e, tmp70496, tmp70502)

								tmp70504 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70505 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70506 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70507 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70508 := Call(__e, tmp70507, V4885)

								tmp70509 := Call(__e, tmp70506, tmp70508)

								tmp70510 := Call(__e, tmp70505, tmp70509)

								tmp70511 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70512 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70513 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70514 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70515 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp70516 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70517 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70518 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp70519 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70520 := Call(__e, tmp70519, V4885)

								tmp70521 := Call(__e, tmp70518, tmp70520)

								tmp70522 := Call(__e, tmp70517, tmp70521)

								tmp70523 := Call(__e, tmp70516, tmp70522)

								tmp70524 := Call(__e, tmp70515, symor, tmp70523)

								tmp70525 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70526 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp70527 := Call(__e, tmp70526, V4885)

								tmp70528 := Call(__e, tmp70525, tmp70527)

								tmp70529 := Call(__e, tmp70514, tmp70524, tmp70528)

								tmp70530 := Call(__e, tmp70513, symshen_4x_4features_4cond_1expand, tmp70529)

								tmp70531 := Call(__e, tmp70512, tmp70530, Nil)

								tmp70532 := Call(__e, tmp70511, True, tmp70531)

								tmp70533 := Call(__e, tmp70504, tmp70510, tmp70532)

								tmp70534 := Call(__e, tmp70495, tmp70503, tmp70533)

								__e.TailApply(tmp70494, symshen_4x_4features_4cond_1expand, tmp70534)
								return

							} else {
								tmp70492 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp70493 := Call(__e, tmp70492, V4885)

								var ifres70432 Obj

								if True == tmp70493 {
									tmp70488 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp70489 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70490 := Call(__e, tmp70489, V4885)

									tmp70491 := Call(__e, tmp70488, symshen_4x_4features_4cond_1expand, tmp70490)

									var ifres70434 Obj

									if True == tmp70491 {
										tmp70484 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp70485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70486 := Call(__e, tmp70485, V4885)

										tmp70487 := Call(__e, tmp70484, tmp70486)

										var ifres70436 Obj

										if True == tmp70487 {
											tmp70478 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp70479 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp70480 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70481 := Call(__e, tmp70480, V4885)

											tmp70482 := Call(__e, tmp70479, tmp70481)

											tmp70483 := Call(__e, tmp70478, tmp70482)

											var ifres70438 Obj

											if True == tmp70483 {
												tmp70470 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp70471 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp70472 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp70473 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70474 := Call(__e, tmp70473, V4885)

												tmp70475 := Call(__e, tmp70472, tmp70474)

												tmp70476 := Call(__e, tmp70471, tmp70475)

												tmp70477 := Call(__e, tmp70470, symnot, tmp70476)

												var ifres70440 Obj

												if True == tmp70477 {
													tmp70462 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp70463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70464 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp70465 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70466 := Call(__e, tmp70465, V4885)

													tmp70467 := Call(__e, tmp70464, tmp70466)

													tmp70468 := Call(__e, tmp70463, tmp70467)

													tmp70469 := Call(__e, tmp70462, tmp70468)

													var ifres70442 Obj

													if True == tmp70469 {
														tmp70452 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp70453 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp70454 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp70455 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp70456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp70457 := Call(__e, tmp70456, V4885)

														tmp70458 := Call(__e, tmp70455, tmp70457)

														tmp70459 := Call(__e, tmp70454, tmp70458)

														tmp70460 := Call(__e, tmp70453, tmp70459)

														tmp70461 := Call(__e, tmp70452, Nil, tmp70460)

														var ifres70444 Obj

														if True == tmp70461 {
															tmp70446 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp70447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp70448 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp70449 := Call(__e, tmp70448, V4885)

															tmp70450 := Call(__e, tmp70447, tmp70449)

															tmp70451 := Call(__e, tmp70446, tmp70450)

															var ifres70445 Obj

															if True == tmp70451 {
																ifres70445 = True

															} else {
																ifres70445 = False

															}

															ifres70444 = ifres70445

														} else {
															ifres70444 = False

														}

														var ifres70443 Obj

														if True == ifres70444 {
															ifres70443 = True

														} else {
															ifres70443 = False

														}

														ifres70442 = ifres70443

													} else {
														ifres70442 = False

													}

													var ifres70441 Obj

													if True == ifres70442 {
														ifres70441 = True

													} else {
														ifres70441 = False

													}

													ifres70440 = ifres70441

												} else {
													ifres70440 = False

												}

												var ifres70439 Obj

												if True == ifres70440 {
													ifres70439 = True

												} else {
													ifres70439 = False

												}

												ifres70438 = ifres70439

											} else {
												ifres70438 = False

											}

											var ifres70437 Obj

											if True == ifres70438 {
												ifres70437 = True

											} else {
												ifres70437 = False

											}

											ifres70436 = ifres70437

										} else {
											ifres70436 = False

										}

										var ifres70435 Obj

										if True == ifres70436 {
											ifres70435 = True

										} else {
											ifres70435 = False

										}

										ifres70434 = ifres70435

									} else {
										ifres70434 = False

									}

									var ifres70433 Obj

									if True == ifres70434 {
										ifres70433 = True

									} else {
										ifres70433 = False

									}

									ifres70432 = ifres70433

								} else {
									ifres70432 = False

								}

								if True == ifres70432 {
									tmp70401 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp70402 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp70403 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70404 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70405 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70406 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70407 := Call(__e, tmp70406, V4885)

									tmp70408 := Call(__e, tmp70405, tmp70407)

									tmp70409 := Call(__e, tmp70404, tmp70408)

									tmp70410 := Call(__e, tmp70403, tmp70409)

									tmp70411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp70412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp70413 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70414 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70415 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70416 := Call(__e, tmp70415, V4885)

									tmp70417 := Call(__e, tmp70414, tmp70416)

									tmp70418 := Call(__e, tmp70413, tmp70417)

									tmp70419 := Call(__e, tmp70412, symshen_4x_4features_4cond_1expand, tmp70418)

									tmp70420 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp70421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp70422 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp70423 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70424 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp70425 := Call(__e, tmp70424, V4885)

									tmp70426 := Call(__e, tmp70423, tmp70425)

									tmp70427 := Call(__e, tmp70422, tmp70426)

									tmp70428 := Call(__e, tmp70421, tmp70427, Nil)

									tmp70429 := Call(__e, tmp70420, True, tmp70428)

									tmp70430 := Call(__e, tmp70411, tmp70419, tmp70429)

									tmp70431 := Call(__e, tmp70402, tmp70410, tmp70430)

									__e.TailApply(tmp70401, symshen_4x_4features_4cond_1expand, tmp70431)
									return

								} else {
									tmp70399 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp70400 := Call(__e, tmp70399, V4885)

									var ifres70369 Obj

									if True == tmp70400 {
										tmp70395 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp70396 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70397 := Call(__e, tmp70396, V4885)

										tmp70398 := Call(__e, tmp70395, symshen_4x_4features_4cond_1expand, tmp70397)

										var ifres70371 Obj

										if True == tmp70398 {
											tmp70391 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp70392 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70393 := Call(__e, tmp70392, V4885)

											tmp70394 := Call(__e, tmp70391, tmp70393)

											var ifres70373 Obj

											if True == tmp70394 {
												tmp70385 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp70386 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70387 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70388 := Call(__e, tmp70387, V4885)

												tmp70389 := Call(__e, tmp70386, tmp70388)

												tmp70390 := Call(__e, tmp70385, tmp70389)

												var ifres70375 Obj

												if True == tmp70390 {
													tmp70377 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

													tmp70378 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp70379 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70380 := Call(__e, tmp70379, V4885)

													tmp70381 := Call(__e, tmp70378, tmp70380)

													tmp70382 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

													tmp70383 := Call(__e, tmp70382, symshen_4x_4features_4_dfeatures_d)

													tmp70384 := Call(__e, tmp70377, tmp70381, tmp70383)

													var ifres70376 Obj

													if True == tmp70384 {
														ifres70376 = True

													} else {
														ifres70376 = False

													}

													ifres70375 = ifres70376

												} else {
													ifres70375 = False

												}

												var ifres70374 Obj

												if True == ifres70375 {
													ifres70374 = True

												} else {
													ifres70374 = False

												}

												ifres70373 = ifres70374

											} else {
												ifres70373 = False

											}

											var ifres70372 Obj

											if True == ifres70373 {
												ifres70372 = True

											} else {
												ifres70372 = False

											}

											ifres70371 = ifres70372

										} else {
											ifres70371 = False

										}

										var ifres70370 Obj

										if True == ifres70371 {
											ifres70370 = True

										} else {
											ifres70370 = False

										}

										ifres70369 = ifres70370

									} else {
										ifres70369 = False

									}

									if True == ifres70369 {
										tmp70364 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70365 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70366 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp70367 := Call(__e, tmp70366, V4885)

										tmp70368 := Call(__e, tmp70365, tmp70367)

										__e.TailApply(tmp70364, tmp70368)
										return

									} else {
										tmp70362 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp70363 := Call(__e, tmp70362, V4885)

										var ifres70342 Obj

										if True == tmp70363 {
											tmp70358 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp70359 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp70360 := Call(__e, tmp70359, V4885)

											tmp70361 := Call(__e, tmp70358, symshen_4x_4features_4cond_1expand, tmp70360)

											var ifres70344 Obj

											if True == tmp70361 {
												tmp70354 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp70355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70356 := Call(__e, tmp70355, V4885)

												tmp70357 := Call(__e, tmp70354, tmp70356)

												var ifres70346 Obj

												if True == tmp70357 {
													tmp70348 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp70349 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70350 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp70351 := Call(__e, tmp70350, V4885)

													tmp70352 := Call(__e, tmp70349, tmp70351)

													tmp70353 := Call(__e, tmp70348, tmp70352)

													var ifres70347 Obj

													if True == tmp70353 {
														ifres70347 = True

													} else {
														ifres70347 = False

													}

													ifres70346 = ifres70347

												} else {
													ifres70346 = False

												}

												var ifres70345 Obj

												if True == ifres70346 {
													ifres70345 = True

												} else {
													ifres70345 = False

												}

												ifres70344 = ifres70345

											} else {
												ifres70344 = False

											}

											var ifres70343 Obj

											if True == ifres70344 {
												ifres70343 = True

											} else {
												ifres70343 = False

											}

											ifres70342 = ifres70343

										} else {
											ifres70342 = False

										}

										if True == ifres70342 {
											tmp70335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp70336 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70337 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp70339 := Call(__e, tmp70338, V4885)

											tmp70340 := Call(__e, tmp70337, tmp70339)

											tmp70341 := Call(__e, tmp70336, tmp70340)

											__e.TailApply(tmp70335, symshen_4x_4features_4cond_1expand, tmp70341)
											return

										} else {
											__e.Return(V4885)
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

	}, 1)

	tmp70844 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4cond_1expand_1macro, tmp70325)

	_ = tmp70844

	tmp70845 := MakeNative(func(__e *ControlFlow) {
		tmp70846 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp70846, symshen_4x_4features_4_dfeatures_d)
		return

	}, 0)

	tmp70847 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4current, tmp70845)

	_ = tmp70847

	tmp70848 := MakeNative(func(__e *ControlFlow) {
		V4887 := __e.Get(1)
		_ = V4887
		tmp70849 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			tmp70850 := MakeNative(func(__e *ControlFlow) {
				Old := __e.Get(1)
				_ = Old
				tmp70851 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					__e.Return(Old)
					return
				}, 1)

				tmp70852 := Call(__e, PrimNS1Value(symns2_1value), symset)

				tmp70853 := Call(__e, tmp70852, symshen_4x_4features_4_dfeatures_d, V4887)

				__e.TailApply(tmp70851, tmp70853)
				return

			}, 1)

			tmp70854 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4features_4current)

			tmp70855 := Call(__e, tmp70854)

			__e.TailApply(tmp70850, tmp70855)
			return

		}, 1)

		tmp70856 := MakeNative(func(__e *ControlFlow) {
			tmp70857 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			__e.TailApply(tmp70857, symshen_4x_4features_4_dfeatures_d)
			return

		}, 0)

		tmp70858 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp70859 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp70860 := Call(__e, tmp70859, symshen_4x_4features_4_dfeatures_d, Nil)

			_ = tmp70860

			tmp70861 := Call(__e, PrimNS1Value(symns2_1value), symshen_4set_1lambda_1form_1entry)

			tmp70862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp70863 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp70864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4features_4cond_1expand_1macro)

				__e.TailApply(tmp70864, X)
				return

			}, 1)

			tmp70865 := Call(__e, tmp70862, symshen_4x_4features_4cond_1expand_1macro, tmp70863)

			tmp70866 := Call(__e, tmp70861, tmp70865)

			_ = tmp70866

			tmp70867 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add_1macro)

			__e.TailApply(tmp70867, symshen_4x_4features_4cond_1expand_1macro)
			return

		}, 1)

		tmp70868 := Call(__e, PrimNS1Value(symtry_1catch), tmp70856, tmp70858)

		__e.TailApply(tmp70849, tmp70868)
		return

	}, 1)

	tmp70869 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4initialise, tmp70848)

	_ = tmp70869

	tmp70870 := MakeNative(func(__e *ControlFlow) {
		V4889 := __e.Get(1)
		_ = V4889
		tmp70871 := MakeNative(func(__e *ControlFlow) {
			Old := __e.Get(1)
			_ = Old
			tmp70872 := MakeNative(func(__e *ControlFlow) {
				__ := __e.Get(1)
				_ = __
				__e.Return(Old)
				return
			}, 1)

			tmp70873 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp70874 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			tmp70875 := Call(__e, tmp70874, V4889, Old)

			tmp70876 := Call(__e, tmp70873, symshen_4x_4features_4_dfeatures_d, tmp70875)

			__e.TailApply(tmp70872, tmp70876)
			return

		}, 1)

		tmp70877 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4features_4current)

		tmp70878 := Call(__e, tmp70877)

		__e.TailApply(tmp70871, tmp70878)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4features_4add, tmp70870)
	return

}, 0)
