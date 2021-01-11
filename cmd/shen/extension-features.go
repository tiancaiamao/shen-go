package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionFeaturesMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")
	gen36849 := MakeNative(func(__e *ControlFlow) {
		V4885 := __e.Get(1)
		_ = V4885
		gen36846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen36847 := Call(__e, gen36846, V4885)

		var gen36848 Obj

		if True == gen36847 {
			gen36841 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen36842 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen36843 := Call(__e, gen36842, V4885)

			gen36844 := Call(__e, gen36841, symshen_4x_4features_4cond_1expand, gen36843)

			var gen36845 Obj

			if True == gen36844 {
				gen36837 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen36838 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen36839 := Call(__e, gen36838, V4885)

				gen36840 := Call(__e, gen36837, Nil, gen36839)

				if True == gen36840 {
					gen36845 = True
				} else {
					gen36845 = False
				}

			} else {
				gen36845 = False
			}

			if True == gen36845 {
				gen36848 = True
			} else {
				gen36848 = False
			}

		} else {
			gen36848 = False
		}

		if True == gen36848 {
			gen36836 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(gen36836, MakeString("Unfulfilled shen.x.features.cond-expand clause."))

			return

		} else {
			gen36833 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen36834 := Call(__e, gen36833, V4885)

			var gen36835 Obj

			if True == gen36834 {
				gen36828 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen36829 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen36830 := Call(__e, gen36829, V4885)

				gen36831 := Call(__e, gen36828, symshen_4x_4features_4cond_1expand, gen36830)

				var gen36832 Obj

				if True == gen36831 {
					gen36823 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen36824 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen36825 := Call(__e, gen36824, V4885)

					gen36826 := Call(__e, gen36823, gen36825)

					var gen36827 Obj

					if True == gen36826 {
						gen36816 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen36817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen36818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36819 := Call(__e, gen36818, V4885)

						gen36820 := Call(__e, gen36817, gen36819)

						gen36821 := Call(__e, gen36816, True, gen36820)

						var gen36822 Obj

						if True == gen36821 {
							gen36809 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen36810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36811 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36812 := Call(__e, gen36811, V4885)

							gen36813 := Call(__e, gen36810, gen36812)

							gen36814 := Call(__e, gen36809, gen36813)

							var gen36815 Obj

							if True == gen36814 {
								gen36801 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen36802 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36803 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36804 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36805 := Call(__e, gen36804, V4885)

								gen36806 := Call(__e, gen36803, gen36805)

								gen36807 := Call(__e, gen36802, gen36806)

								gen36808 := Call(__e, gen36801, Nil, gen36807)

								if True == gen36808 {
									gen36815 = True
								} else {
									gen36815 = False
								}

							} else {
								gen36815 = False
							}

							if True == gen36815 {
								gen36822 = True
							} else {
								gen36822 = False
							}

						} else {
							gen36822 = False
						}

						if True == gen36822 {
							gen36827 = True
						} else {
							gen36827 = False
						}

					} else {
						gen36827 = False
					}

					if True == gen36827 {
						gen36832 = True
					} else {
						gen36832 = False
					}

				} else {
					gen36832 = False
				}

				if True == gen36832 {
					gen36835 = True
				} else {
					gen36835 = False
				}

			} else {
				gen36835 = False
			}

			if True == gen36835 {
				gen36796 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen36797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen36798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen36799 := Call(__e, gen36798, V4885)

				gen36800 := Call(__e, gen36797, gen36799)

				__e.TailApply(gen36796, gen36800)

				return

			} else {
				gen36793 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen36794 := Call(__e, gen36793, V4885)

				var gen36795 Obj

				if True == gen36794 {
					gen36788 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen36789 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen36790 := Call(__e, gen36789, V4885)

					gen36791 := Call(__e, gen36788, symshen_4x_4features_4cond_1expand, gen36790)

					var gen36792 Obj

					if True == gen36791 {
						gen36783 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen36784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36785 := Call(__e, gen36784, V4885)

						gen36786 := Call(__e, gen36783, gen36785)

						var gen36787 Obj

						if True == gen36786 {
							gen36776 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen36777 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen36778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36779 := Call(__e, gen36778, V4885)

							gen36780 := Call(__e, gen36777, gen36779)

							gen36781 := Call(__e, gen36776, gen36780)

							var gen36782 Obj

							if True == gen36781 {
								gen36767 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen36768 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36769 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36770 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36771 := Call(__e, gen36770, V4885)

								gen36772 := Call(__e, gen36769, gen36771)

								gen36773 := Call(__e, gen36768, gen36772)

								gen36774 := Call(__e, gen36767, symand, gen36773)

								var gen36775 Obj

								if True == gen36774 {
									gen36758 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen36759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36760 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36761 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36762 := Call(__e, gen36761, V4885)

									gen36763 := Call(__e, gen36760, gen36762)

									gen36764 := Call(__e, gen36759, gen36763)

									gen36765 := Call(__e, gen36758, Nil, gen36764)

									var gen36766 Obj

									if True == gen36765 {
										gen36752 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36753 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36754 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36755 := Call(__e, gen36754, V4885)

										gen36756 := Call(__e, gen36753, gen36755)

										gen36757 := Call(__e, gen36752, gen36756)

										if True == gen36757 {
											gen36766 = True
										} else {
											gen36766 = False
										}

									} else {
										gen36766 = False
									}

									if True == gen36766 {
										gen36775 = True
									} else {
										gen36775 = False
									}

								} else {
									gen36775 = False
								}

								if True == gen36775 {
									gen36782 = True
								} else {
									gen36782 = False
								}

							} else {
								gen36782 = False
							}

							if True == gen36782 {
								gen36787 = True
							} else {
								gen36787 = False
							}

						} else {
							gen36787 = False
						}

						if True == gen36787 {
							gen36792 = True
						} else {
							gen36792 = False
						}

					} else {
						gen36792 = False
					}

					if True == gen36792 {
						gen36795 = True
					} else {
						gen36795 = False
					}

				} else {
					gen36795 = False
				}

				if True == gen36795 {
					gen36747 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen36748 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen36749 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen36750 := Call(__e, gen36749, V4885)

					gen36751 := Call(__e, gen36748, gen36750)

					__e.TailApply(gen36747, gen36751)

					return

				} else {
					gen36744 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen36745 := Call(__e, gen36744, V4885)

					var gen36746 Obj

					if True == gen36745 {
						gen36739 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen36740 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen36741 := Call(__e, gen36740, V4885)

						gen36742 := Call(__e, gen36739, symshen_4x_4features_4cond_1expand, gen36741)

						var gen36743 Obj

						if True == gen36742 {
							gen36734 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen36735 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36736 := Call(__e, gen36735, V4885)

							gen36737 := Call(__e, gen36734, gen36736)

							var gen36738 Obj

							if True == gen36737 {
								gen36727 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen36728 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36730 := Call(__e, gen36729, V4885)

								gen36731 := Call(__e, gen36728, gen36730)

								gen36732 := Call(__e, gen36727, gen36731)

								var gen36733 Obj

								if True == gen36732 {
									gen36718 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen36719 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36720 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36722 := Call(__e, gen36721, V4885)

									gen36723 := Call(__e, gen36720, gen36722)

									gen36724 := Call(__e, gen36719, gen36723)

									gen36725 := Call(__e, gen36718, symand, gen36724)

									var gen36726 Obj

									if True == gen36725 {
										gen36709 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36710 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36711 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen36712 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36713 := Call(__e, gen36712, V4885)

										gen36714 := Call(__e, gen36711, gen36713)

										gen36715 := Call(__e, gen36710, gen36714)

										gen36716 := Call(__e, gen36709, gen36715)

										var gen36717 Obj

										if True == gen36716 {
											gen36703 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen36704 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36705 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36706 := Call(__e, gen36705, V4885)

											gen36707 := Call(__e, gen36704, gen36706)

											gen36708 := Call(__e, gen36703, gen36707)

											if True == gen36708 {
												gen36717 = True
											} else {
												gen36717 = False
											}

										} else {
											gen36717 = False
										}

										if True == gen36717 {
											gen36726 = True
										} else {
											gen36726 = False
										}

									} else {
										gen36726 = False
									}

									if True == gen36726 {
										gen36733 = True
									} else {
										gen36733 = False
									}

								} else {
									gen36733 = False
								}

								if True == gen36733 {
									gen36738 = True
								} else {
									gen36738 = False
								}

							} else {
								gen36738 = False
							}

							if True == gen36738 {
								gen36743 = True
							} else {
								gen36743 = False
							}

						} else {
							gen36743 = False
						}

						if True == gen36743 {
							gen36746 = True
						} else {
							gen36746 = False
						}

					} else {
						gen36746 = False
					}

					if True == gen36746 {
						gen36666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen36667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen36668 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen36669 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36670 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen36671 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36672 := Call(__e, gen36671, V4885)

						gen36673 := Call(__e, gen36670, gen36672)

						gen36674 := Call(__e, gen36669, gen36673)

						gen36675 := Call(__e, gen36668, gen36674)

						gen36676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen36677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen36678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen36679 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen36680 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36681 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36682 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen36683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36684 := Call(__e, gen36683, V4885)

						gen36685 := Call(__e, gen36682, gen36684)

						gen36686 := Call(__e, gen36681, gen36685)

						gen36687 := Call(__e, gen36680, gen36686)

						gen36688 := Call(__e, gen36679, symand, gen36687)

						gen36689 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36690 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36691 := Call(__e, gen36690, V4885)

						gen36692 := Call(__e, gen36689, gen36691)

						gen36693 := Call(__e, gen36678, gen36688, gen36692)

						gen36694 := Call(__e, gen36677, symshen_4x_4features_4cond_1expand, gen36693)

						gen36695 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36696 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36697 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen36698 := Call(__e, gen36697, V4885)

						gen36699 := Call(__e, gen36696, gen36698)

						gen36700 := Call(__e, gen36695, gen36699)

						gen36701 := Call(__e, gen36676, gen36694, gen36700)

						gen36702 := Call(__e, gen36667, gen36675, gen36701)

						__e.TailApply(gen36666, symshen_4x_4features_4cond_1expand, gen36702)

						return

					} else {
						gen36663 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen36664 := Call(__e, gen36663, V4885)

						var gen36665 Obj

						if True == gen36664 {
							gen36658 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen36659 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen36660 := Call(__e, gen36659, V4885)

							gen36661 := Call(__e, gen36658, symshen_4x_4features_4cond_1expand, gen36660)

							var gen36662 Obj

							if True == gen36661 {
								gen36653 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen36654 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36655 := Call(__e, gen36654, V4885)

								gen36656 := Call(__e, gen36653, gen36655)

								var gen36657 Obj

								if True == gen36656 {
									gen36646 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen36647 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36648 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36649 := Call(__e, gen36648, V4885)

									gen36650 := Call(__e, gen36647, gen36649)

									gen36651 := Call(__e, gen36646, gen36650)

									var gen36652 Obj

									if True == gen36651 {
										gen36637 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen36638 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen36639 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen36640 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36641 := Call(__e, gen36640, V4885)

										gen36642 := Call(__e, gen36639, gen36641)

										gen36643 := Call(__e, gen36638, gen36642)

										gen36644 := Call(__e, gen36637, symor, gen36643)

										var gen36645 Obj

										if True == gen36644 {
											gen36628 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen36629 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36630 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36631 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36632 := Call(__e, gen36631, V4885)

											gen36633 := Call(__e, gen36630, gen36632)

											gen36634 := Call(__e, gen36629, gen36633)

											gen36635 := Call(__e, gen36628, Nil, gen36634)

											var gen36636 Obj

											if True == gen36635 {
												gen36622 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												gen36623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36624 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36625 := Call(__e, gen36624, V4885)

												gen36626 := Call(__e, gen36623, gen36625)

												gen36627 := Call(__e, gen36622, gen36626)

												if True == gen36627 {
													gen36636 = True
												} else {
													gen36636 = False
												}

											} else {
												gen36636 = False
											}

											if True == gen36636 {
												gen36645 = True
											} else {
												gen36645 = False
											}

										} else {
											gen36645 = False
										}

										if True == gen36645 {
											gen36652 = True
										} else {
											gen36652 = False
										}

									} else {
										gen36652 = False
									}

									if True == gen36652 {
										gen36657 = True
									} else {
										gen36657 = False
									}

								} else {
									gen36657 = False
								}

								if True == gen36657 {
									gen36662 = True
								} else {
									gen36662 = False
								}

							} else {
								gen36662 = False
							}

							if True == gen36662 {
								gen36665 = True
							} else {
								gen36665 = False
							}

						} else {
							gen36665 = False
						}

						if True == gen36665 {
							gen36615 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen36616 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36617 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36618 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen36619 := Call(__e, gen36618, V4885)

							gen36620 := Call(__e, gen36617, gen36619)

							gen36621 := Call(__e, gen36616, gen36620)

							__e.TailApply(gen36615, symshen_4x_4features_4cond_1expand, gen36621)

							return

						} else {
							gen36612 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen36613 := Call(__e, gen36612, V4885)

							var gen36614 Obj

							if True == gen36613 {
								gen36607 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen36608 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36609 := Call(__e, gen36608, V4885)

								gen36610 := Call(__e, gen36607, symshen_4x_4features_4cond_1expand, gen36609)

								var gen36611 Obj

								if True == gen36610 {
									gen36602 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen36603 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36604 := Call(__e, gen36603, V4885)

									gen36605 := Call(__e, gen36602, gen36604)

									var gen36606 Obj

									if True == gen36605 {
										gen36595 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36596 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen36597 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36598 := Call(__e, gen36597, V4885)

										gen36599 := Call(__e, gen36596, gen36598)

										gen36600 := Call(__e, gen36595, gen36599)

										var gen36601 Obj

										if True == gen36600 {
											gen36586 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen36587 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36588 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36589 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36590 := Call(__e, gen36589, V4885)

											gen36591 := Call(__e, gen36588, gen36590)

											gen36592 := Call(__e, gen36587, gen36591)

											gen36593 := Call(__e, gen36586, symor, gen36592)

											var gen36594 Obj

											if True == gen36593 {
												gen36577 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												gen36578 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36579 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen36580 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36581 := Call(__e, gen36580, V4885)

												gen36582 := Call(__e, gen36579, gen36581)

												gen36583 := Call(__e, gen36578, gen36582)

												gen36584 := Call(__e, gen36577, gen36583)

												var gen36585 Obj

												if True == gen36584 {
													gen36571 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													gen36572 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36573 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36574 := Call(__e, gen36573, V4885)

													gen36575 := Call(__e, gen36572, gen36574)

													gen36576 := Call(__e, gen36571, gen36575)

													if True == gen36576 {
														gen36585 = True
													} else {
														gen36585 = False
													}

												} else {
													gen36585 = False
												}

												if True == gen36585 {
													gen36594 = True
												} else {
													gen36594 = False
												}

											} else {
												gen36594 = False
											}

											if True == gen36594 {
												gen36601 = True
											} else {
												gen36601 = False
											}

										} else {
											gen36601 = False
										}

										if True == gen36601 {
											gen36606 = True
										} else {
											gen36606 = False
										}

									} else {
										gen36606 = False
									}

									if True == gen36606 {
										gen36611 = True
									} else {
										gen36611 = False
									}

								} else {
									gen36611 = False
								}

								if True == gen36611 {
									gen36614 = True
								} else {
									gen36614 = False
								}

							} else {
								gen36614 = False
							}

							if True == gen36614 {
								gen36530 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36531 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36532 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36533 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36534 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36535 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36536 := Call(__e, gen36535, V4885)

								gen36537 := Call(__e, gen36534, gen36536)

								gen36538 := Call(__e, gen36533, gen36537)

								gen36539 := Call(__e, gen36532, gen36538)

								gen36540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36541 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36542 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36543 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36544 := Call(__e, gen36543, V4885)

								gen36545 := Call(__e, gen36542, gen36544)

								gen36546 := Call(__e, gen36541, gen36545)

								gen36547 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36548 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36549 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36550 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36551 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36552 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36553 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36554 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen36555 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36556 := Call(__e, gen36555, V4885)

								gen36557 := Call(__e, gen36554, gen36556)

								gen36558 := Call(__e, gen36553, gen36557)

								gen36559 := Call(__e, gen36552, gen36558)

								gen36560 := Call(__e, gen36551, symor, gen36559)

								gen36561 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36562 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36563 := Call(__e, gen36562, V4885)

								gen36564 := Call(__e, gen36561, gen36563)

								gen36565 := Call(__e, gen36550, gen36560, gen36564)

								gen36566 := Call(__e, gen36549, symshen_4x_4features_4cond_1expand, gen36565)

								gen36567 := Call(__e, gen36548, gen36566, Nil)

								gen36568 := Call(__e, gen36547, True, gen36567)

								gen36569 := Call(__e, gen36540, gen36546, gen36568)

								gen36570 := Call(__e, gen36531, gen36539, gen36569)

								__e.TailApply(gen36530, symshen_4x_4features_4cond_1expand, gen36570)

								return

							} else {
								gen36527 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen36528 := Call(__e, gen36527, V4885)

								var gen36529 Obj

								if True == gen36528 {
									gen36522 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen36523 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36524 := Call(__e, gen36523, V4885)

									gen36525 := Call(__e, gen36522, symshen_4x_4features_4cond_1expand, gen36524)

									var gen36526 Obj

									if True == gen36525 {
										gen36517 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36518 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36519 := Call(__e, gen36518, V4885)

										gen36520 := Call(__e, gen36517, gen36519)

										var gen36521 Obj

										if True == gen36520 {
											gen36510 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen36511 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36512 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36513 := Call(__e, gen36512, V4885)

											gen36514 := Call(__e, gen36511, gen36513)

											gen36515 := Call(__e, gen36510, gen36514)

											var gen36516 Obj

											if True == gen36515 {
												gen36501 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen36502 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen36503 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen36504 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36505 := Call(__e, gen36504, V4885)

												gen36506 := Call(__e, gen36503, gen36505)

												gen36507 := Call(__e, gen36502, gen36506)

												gen36508 := Call(__e, gen36501, symnot, gen36507)

												var gen36509 Obj

												if True == gen36508 {
													gen36492 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													gen36493 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36494 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													gen36495 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36496 := Call(__e, gen36495, V4885)

													gen36497 := Call(__e, gen36494, gen36496)

													gen36498 := Call(__e, gen36493, gen36497)

													gen36499 := Call(__e, gen36492, gen36498)

													var gen36500 Obj

													if True == gen36499 {
														gen36481 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														gen36482 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen36483 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen36484 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen36485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														gen36486 := Call(__e, gen36485, V4885)

														gen36487 := Call(__e, gen36484, gen36486)

														gen36488 := Call(__e, gen36483, gen36487)

														gen36489 := Call(__e, gen36482, gen36488)

														gen36490 := Call(__e, gen36481, Nil, gen36489)

														var gen36491 Obj

														if True == gen36490 {
															gen36475 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															gen36476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															gen36477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															gen36478 := Call(__e, gen36477, V4885)

															gen36479 := Call(__e, gen36476, gen36478)

															gen36480 := Call(__e, gen36475, gen36479)

															if True == gen36480 {
																gen36491 = True
															} else {
																gen36491 = False
															}

														} else {
															gen36491 = False
														}

														if True == gen36491 {
															gen36500 = True
														} else {
															gen36500 = False
														}

													} else {
														gen36500 = False
													}

													if True == gen36500 {
														gen36509 = True
													} else {
														gen36509 = False
													}

												} else {
													gen36509 = False
												}

												if True == gen36509 {
													gen36516 = True
												} else {
													gen36516 = False
												}

											} else {
												gen36516 = False
											}

											if True == gen36516 {
												gen36521 = True
											} else {
												gen36521 = False
											}

										} else {
											gen36521 = False
										}

										if True == gen36521 {
											gen36526 = True
										} else {
											gen36526 = False
										}

									} else {
										gen36526 = False
									}

									if True == gen36526 {
										gen36529 = True
									} else {
										gen36529 = False
									}

								} else {
									gen36529 = False
								}

								if True == gen36529 {
									gen36444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen36445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen36446 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36448 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36449 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36450 := Call(__e, gen36449, V4885)

									gen36451 := Call(__e, gen36448, gen36450)

									gen36452 := Call(__e, gen36447, gen36451)

									gen36453 := Call(__e, gen36446, gen36452)

									gen36454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen36455 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen36456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36457 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36458 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36459 := Call(__e, gen36458, V4885)

									gen36460 := Call(__e, gen36457, gen36459)

									gen36461 := Call(__e, gen36456, gen36460)

									gen36462 := Call(__e, gen36455, symshen_4x_4features_4cond_1expand, gen36461)

									gen36463 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen36464 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen36465 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen36466 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36467 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen36468 := Call(__e, gen36467, V4885)

									gen36469 := Call(__e, gen36466, gen36468)

									gen36470 := Call(__e, gen36465, gen36469)

									gen36471 := Call(__e, gen36464, gen36470, Nil)

									gen36472 := Call(__e, gen36463, True, gen36471)

									gen36473 := Call(__e, gen36454, gen36462, gen36472)

									gen36474 := Call(__e, gen36445, gen36453, gen36473)

									__e.TailApply(gen36444, symshen_4x_4features_4cond_1expand, gen36474)

									return

								} else {
									gen36441 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen36442 := Call(__e, gen36441, V4885)

									var gen36443 Obj

									if True == gen36442 {
										gen36436 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen36437 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen36438 := Call(__e, gen36437, V4885)

										gen36439 := Call(__e, gen36436, symshen_4x_4features_4cond_1expand, gen36438)

										var gen36440 Obj

										if True == gen36439 {
											gen36431 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen36432 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36433 := Call(__e, gen36432, V4885)

											gen36434 := Call(__e, gen36431, gen36433)

											var gen36435 Obj

											if True == gen36434 {
												gen36424 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												gen36425 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36426 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36427 := Call(__e, gen36426, V4885)

												gen36428 := Call(__e, gen36425, gen36427)

												gen36429 := Call(__e, gen36424, gen36428)

												var gen36430 Obj

												if True == gen36429 {
													gen36416 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

													gen36417 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													gen36418 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36419 := Call(__e, gen36418, V4885)

													gen36420 := Call(__e, gen36417, gen36419)

													gen36421 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

													gen36422 := Call(__e, gen36421, symshen_4x_4features_4_dfeatures_d)

													gen36423 := Call(__e, gen36416, gen36420, gen36422)

													if True == gen36423 {
														gen36430 = True
													} else {
														gen36430 = False
													}

												} else {
													gen36430 = False
												}

												if True == gen36430 {
													gen36435 = True
												} else {
													gen36435 = False
												}

											} else {
												gen36435 = False
											}

											if True == gen36435 {
												gen36440 = True
											} else {
												gen36440 = False
											}

										} else {
											gen36440 = False
										}

										if True == gen36440 {
											gen36443 = True
										} else {
											gen36443 = False
										}

									} else {
										gen36443 = False
									}

									if True == gen36443 {
										gen36411 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen36412 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36413 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen36414 := Call(__e, gen36413, V4885)

										gen36415 := Call(__e, gen36412, gen36414)

										__e.TailApply(gen36411, gen36415)

										return

									} else {
										gen36408 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36409 := Call(__e, gen36408, V4885)

										var gen36410 Obj

										if True == gen36409 {
											gen36403 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen36404 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36405 := Call(__e, gen36404, V4885)

											gen36406 := Call(__e, gen36403, symshen_4x_4features_4cond_1expand, gen36405)

											var gen36407 Obj

											if True == gen36406 {
												gen36398 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												gen36399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen36400 := Call(__e, gen36399, V4885)

												gen36401 := Call(__e, gen36398, gen36400)

												var gen36402 Obj

												if True == gen36401 {
													gen36392 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													gen36393 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36394 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36395 := Call(__e, gen36394, V4885)

													gen36396 := Call(__e, gen36393, gen36395)

													gen36397 := Call(__e, gen36392, gen36396)

													if True == gen36397 {
														gen36402 = True
													} else {
														gen36402 = False
													}

												} else {
													gen36402 = False
												}

												if True == gen36402 {
													gen36407 = True
												} else {
													gen36407 = False
												}

											} else {
												gen36407 = False
											}

											if True == gen36407 {
												gen36410 = True
											} else {
												gen36410 = False
											}

										} else {
											gen36410 = False
										}

										if True == gen36410 {
											gen36385 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											gen36386 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36387 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36388 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen36389 := Call(__e, gen36388, V4885)

											gen36390 := Call(__e, gen36387, gen36389)

											gen36391 := Call(__e, gen36386, gen36390)

											__e.TailApply(gen36385, symshen_4x_4features_4cond_1expand, gen36391)

											return

										} else {
											if True == True {
												__e.Return(V4885)
												return
											} else {
												gen36384 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

												__e.TailApply(gen36384, MakeString("no cond match"))

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

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4cond_1expand_1macro, gen36849)

	gen36851 := MakeNative(func(__e *ControlFlow) {
		gen36850 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen36850, symshen_4x_4features_4_dfeatures_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4current, gen36851)

	gen36870 := MakeNative(func(__e *ControlFlow) {
		V4887 := __e.Get(1)
		_ = V4887
		gen36858 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			gen36855 := MakeNative(func(__e *ControlFlow) {
				Old := __e.Get(1)
				_ = Old
				gen36852 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					__e.Return(Old)
					return
				}, 1)

				gen36853 := Call(__e, PrimNS1Value(symns2_1value), symset)

				gen36854 := Call(__e, gen36853, symshen_4x_4features_4_dfeatures_d, V4887)

				__e.TailApply(gen36852, gen36854)

				return

			}, 1)

			gen36856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4features_4current)

			gen36857 := Call(__e, gen36856)

			__e.TailApply(gen36855, gen36857)

			return

		}, 1)

		gen36860 := MakeNative(func(__e *ControlFlow) {
			gen36859 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			__e.TailApply(gen36859, symshen_4x_4features_4_dfeatures_d)

			return

		}, 0)

		gen36868 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen36861 := Call(__e, PrimNS1Value(symns2_1value), symset)

			Call(__e, gen36861, symshen_4x_4features_4_dfeatures_d, Nil)

			gen36862 := Call(__e, PrimNS1Value(symns2_1value), symshen_4set_1lambda_1form_1entry)

			gen36863 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen36865 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen36864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4features_4cond_1expand_1macro)

				__e.TailApply(gen36864, X)

				return

			}, 1)

			gen36866 := Call(__e, gen36863, symshen_4x_4features_4cond_1expand_1macro, gen36865)

			Call(__e, gen36862, gen36866)

			gen36867 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add_1macro)

			__e.TailApply(gen36867, symshen_4x_4features_4cond_1expand_1macro)

			return

		}, 1)

		gen36869 := Call(__e, PrimNS1Value(symtry_1catch), gen36860, gen36868)

		__e.TailApply(gen36858, gen36869)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4initialise, gen36870)

	gen36879 := MakeNative(func(__e *ControlFlow) {
		V4889 := __e.Get(1)
		_ = V4889
		gen36876 := MakeNative(func(__e *ControlFlow) {
			Old := __e.Get(1)
			_ = Old
			gen36871 := MakeNative(func(__e *ControlFlow) {
				__ := __e.Get(1)
				_ = __
				__e.Return(Old)
				return
			}, 1)

			gen36872 := Call(__e, PrimNS1Value(symns2_1value), symset)

			gen36873 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			gen36874 := Call(__e, gen36873, V4889, Old)

			gen36875 := Call(__e, gen36872, symshen_4x_4features_4_dfeatures_d, gen36874)

			__e.TailApply(gen36871, gen36875)

			return

		}, 1)

		gen36877 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4features_4current)

		gen36878 := Call(__e, gen36877)

		__e.TailApply(gen36876, gen36878)

		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4features_4add, gen36879)

	return

}, 0)
