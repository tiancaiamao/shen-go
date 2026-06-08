package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
	tmp15527 := MakeNative(func(__e *ControlFlow) {
		V112 := __e.Get(1)
		_ = V112
		V113 := __e.Get(2)
		_ = V113
		tmp15528 := MakeNative(func(__e *ControlFlow) {
			W114 := __e.Get(1)
			_ = W114
			tmp15535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W114)

			if True == tmp15535 {
				__e.Return(PrimSimpleError(MakeString("parse failure\n")))
				return
			} else {
				tmp15533 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), W114)

				if True == tmp15533 {
					tmp15529 := Call(__e, PrimFunc(symshen_4in_1_6), W114)

					tmp15530 := PrimSet(symshen_4_dresidue_d, tmp15529)

					_ = tmp15530

					tmp15531 := PrimValue(symshen_4_dresidue_d)

					__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp15531)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4_5_1out), W114)
					return
				}

			}

		}, 1)

		tmp15536 := Call(__e, V112, V113)

		__e.TailApply(tmp15528, tmp15536)
		return

	}, 2)

	tmp15537 := Call(__e, ns2_1set, symcompile, tmp15527)

	_ = tmp15537

	tmp15538 := MakeNative(func(__e *ControlFlow) {
		V115 := __e.Get(1)
		_ = V115
		tmp15539 := PrimValue(sym_dmaximum_1print_1sequence_1size_d)

		tmp15540 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp15539, MakeNumber(0), V115)

		tmp15541 := PrimStringConcat(MakeString("syntax error here: "), tmp15540)

		tmp15542 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp15541)

		__e.Return(PrimSimpleError(tmp15542))
		return

	}, 1)

	tmp15543 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp15538)

	_ = tmp15543

	tmp15544 := MakeNative(func(__e *ControlFlow) {
		V123 := __e.Get(1)
		_ = V123
		V124 := __e.Get(2)
		_ = V124
		V125 := __e.Get(3)
		_ = V125
		tmp15555 := PrimEqual(Nil, V125)

		if True == tmp15555 {
			__e.Return(MakeString("\n"))
			return
		} else {
			tmp15553 := PrimEqual(V123, V124)

			if True == tmp15553 {
				__e.Return(MakeString("...etc \n"))
				return
			} else {
				tmp15551 := PrimIsPair(V125)

				if True == tmp15551 {
					tmp15545 := PrimHead(V125)

					tmp15546 := Call(__e, PrimFunc(symshen_4app), tmp15545, MakeString(" "), symshen_4s)

					tmp15547 := PrimNumberAdd(V124, MakeNumber(1))

					tmp15548 := PrimTail(V125)

					tmp15549 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp15547, tmp15548)

					__e.Return(PrimStringConcat(tmp15546, tmp15549))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4syntax_1error_1message)
					return
				}

			}

		}

	}, 3)

	tmp15556 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp15544)

	_ = tmp15556

	tmp15557 := MakeNative(func(__e *ControlFlow) {
		V126 := __e.Get(1)
		_ = V126
		tmp15558 := Call(__e, PrimFunc(symfail))

		__e.Return(PrimEqual(V126, tmp15558))
		return

	}, 1)

	tmp15559 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp15557)

	_ = tmp15559

	tmp15560 := MakeNative(func(__e *ControlFlow) {
		V127 := __e.Get(1)
		_ = V127
		tmp15561 := Call(__e, PrimFunc(symshen_4in_1_6), V127)

		__e.Return(PrimIsPair(tmp15561))
		return

	}, 1)

	tmp15562 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp15560)

	_ = tmp15562

	tmp15563 := MakeNative(func(__e *ControlFlow) {
		V130 := __e.Get(1)
		_ = V130
		tmp15576 := PrimIsPair(V130)

		var ifres15567 Obj

		if True == tmp15576 {
			tmp15574 := PrimTail(V130)

			tmp15575 := PrimIsPair(tmp15574)

			var ifres15569 Obj

			if True == tmp15575 {
				tmp15571 := PrimTail(V130)

				tmp15572 := PrimTail(tmp15571)

				tmp15573 := PrimEqual(Nil, tmp15572)

				var ifres15570 Obj

				if True == tmp15573 {
					ifres15570 = True

				} else {
					ifres15570 = False

				}

				ifres15569 = ifres15570

			} else {
				ifres15569 = False

			}

			var ifres15568 Obj

			if True == ifres15569 {
				ifres15568 = True

			} else {
				ifres15568 = False

			}

			ifres15567 = ifres15568

		} else {
			ifres15567 = False

		}

		if True == ifres15567 {
			tmp15564 := PrimTail(V130)

			__e.Return(PrimHead(tmp15564))
			return

		} else {
			tmp15565 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)

			__e.Return(PrimSimpleError(tmp15565))
			return

		}

	}, 1)

	tmp15577 := Call(__e, ns2_1set, symshen_4objectcode, tmp15563)

	_ = tmp15577

	tmp15578 := MakeNative(func(__e *ControlFlow) {
		V131 := __e.Get(1)
		_ = V131
		tmp15579 := MakeNative(func(__e *ControlFlow) {
			Z132 := __e.Get(1)
			_ = Z132
			__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
			return
		}, 1)

		__e.TailApply(PrimFunc(symcompile), tmp15579, V131)
		return

	}, 1)

	tmp15580 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp15578)

	_ = tmp15580

	tmp15581 := MakeNative(func(__e *ControlFlow) {
		V133 := __e.Get(1)
		_ = V133
		tmp15582 := MakeNative(func(__e *ControlFlow) {
			W134 := __e.Get(1)
			_ = W134
			tmp15584 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W134)

			if True == tmp15584 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W134)
				return
			}

		}, 1)

		tmp15620 := PrimIsPair(V133)

		var ifres15585 Obj

		if True == tmp15620 {
			tmp15586 := MakeNative(func(__e *ControlFlow) {
				W135 := __e.Get(1)
				_ = W135
				tmp15587 := MakeNative(func(__e *ControlFlow) {
					W136 := __e.Get(1)
					_ = W136
					tmp15588 := MakeNative(func(__e *ControlFlow) {
						W137 := __e.Get(1)
						_ = W137
						tmp15614 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W137)

						if True == tmp15614 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp15589 := MakeNative(func(__e *ControlFlow) {
								W138 := __e.Get(1)
								_ = W138
								tmp15590 := MakeNative(func(__e *ControlFlow) {
									W139 := __e.Get(1)
									_ = W139
									tmp15591 := MakeNative(func(__e *ControlFlow) {
										W140 := __e.Get(1)
										_ = W140
										tmp15609 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W140)

										if True == tmp15609 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp15592 := MakeNative(func(__e *ControlFlow) {
												W141 := __e.Get(1)
												_ = W141
												tmp15593 := MakeNative(func(__e *ControlFlow) {
													W142 := __e.Get(1)
													_ = W142
													tmp15594 := MakeNative(func(__e *ControlFlow) {
														W143 := __e.Get(1)
														_ = W143
														tmp15595 := MakeNative(func(__e *ControlFlow) {
															W144 := __e.Get(1)
															_ = W144
															__e.Return(W144)
															return
														}, 1)

														tmp15596 := PrimCons(W135, Nil)

														tmp15597 := PrimCons(symdefine, tmp15596)

														tmp15598 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), W138, W143, W141)

														tmp15599 := PrimCons(tmp15598, Nil)

														tmp15600 := PrimCons(sym_1_6, tmp15599)

														tmp15601 := PrimCons(W143, tmp15600)

														tmp15602 := Call(__e, PrimFunc(symappend), W138, tmp15601)

														tmp15603 := Call(__e, PrimFunc(symappend), tmp15597, tmp15602)

														__e.TailApply(tmp15595, tmp15603)
														return

													}, 1)

													tmp15604 := Call(__e, PrimFunc(symgensym), symS)

													tmp15605 := Call(__e, tmp15594, tmp15604)

													__e.TailApply(PrimFunc(symshen_4comb), W142, tmp15605)
													return

												}, 1)

												tmp15606 := Call(__e, PrimFunc(symshen_4in_1_6), W140)

												__e.TailApply(tmp15593, tmp15606)
												return

											}, 1)

											tmp15607 := Call(__e, PrimFunc(symshen_4_5_1out), W140)

											__e.TailApply(tmp15592, tmp15607)
											return

										}

									}, 1)

									tmp15610 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W139)

									__e.TailApply(tmp15591, tmp15610)
									return

								}, 1)

								tmp15611 := Call(__e, PrimFunc(symshen_4in_1_6), W137)

								__e.TailApply(tmp15590, tmp15611)
								return

							}, 1)

							tmp15612 := Call(__e, PrimFunc(symshen_4_5_1out), W137)

							__e.TailApply(tmp15589, tmp15612)
							return

						}

					}, 1)

					tmp15615 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), W136)

					__e.TailApply(tmp15588, tmp15615)
					return

				}, 1)

				tmp15616 := Call(__e, PrimFunc(symtail), V133)

				__e.TailApply(tmp15587, tmp15616)
				return

			}, 1)

			tmp15617 := Call(__e, PrimFunc(symhead), V133)

			tmp15618 := Call(__e, tmp15586, tmp15617)

			ifres15585 = tmp15618

		} else {
			tmp15619 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres15585 = tmp15619

		}

		__e.TailApply(tmp15582, ifres15585)
		return

	}, 1)

	tmp15621 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp15581)

	_ = tmp15621

	tmp15622 := MakeNative(func(__e *ControlFlow) {
		V145 := __e.Get(1)
		_ = V145
		tmp15623 := MakeNative(func(__e *ControlFlow) {
			W146 := __e.Get(1)
			_ = W146
			tmp15635 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W146)

			if True == tmp15635 {
				tmp15624 := MakeNative(func(__e *ControlFlow) {
					W161 := __e.Get(1)
					_ = W161
					tmp15626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W161)

					if True == tmp15626 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W161)
						return
					}

				}, 1)

				tmp15627 := MakeNative(func(__e *ControlFlow) {
					W162 := __e.Get(1)
					_ = W162
					tmp15631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W162)

					if True == tmp15631 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15628 := MakeNative(func(__e *ControlFlow) {
							W163 := __e.Get(1)
							_ = W163
							__e.TailApply(PrimFunc(symshen_4comb), W163, Nil)
							return
						}, 1)

						tmp15629 := Call(__e, PrimFunc(symshen_4in_1_6), W162)

						__e.TailApply(tmp15628, tmp15629)
						return

					}

				}, 1)

				tmp15632 := Call(__e, PrimFunc(sym_5e_6), V145)

				tmp15633 := Call(__e, tmp15627, tmp15632)

				__e.TailApply(tmp15624, tmp15633)
				return

			} else {
				__e.Return(W146)
				return
			}

		}, 1)

		tmp15698 := PrimIsPair(V145)

		var ifres15636 Obj

		if True == tmp15698 {
			tmp15637 := MakeNative(func(__e *ControlFlow) {
				W147 := __e.Get(1)
				_ = W147
				tmp15638 := MakeNative(func(__e *ControlFlow) {
					W148 := __e.Get(1)
					_ = W148
					tmp15693 := Call(__e, PrimFunc(symshen_4ccons_2), W148)

					if True == tmp15693 {
						tmp15639 := MakeNative(func(__e *ControlFlow) {
							W149 := __e.Get(1)
							_ = W149
							tmp15640 := MakeNative(func(__e *ControlFlow) {
								W150 := __e.Get(1)
								_ = W150
								tmp15689 := Call(__e, PrimFunc(symshen_4hds_a_2), W149, symlist)

								if True == tmp15689 {
									tmp15641 := MakeNative(func(__e *ControlFlow) {
										W151 := __e.Get(1)
										_ = W151
										tmp15686 := PrimIsPair(W151)

										if True == tmp15686 {
											tmp15642 := MakeNative(func(__e *ControlFlow) {
												W152 := __e.Get(1)
												_ = W152
												tmp15643 := MakeNative(func(__e *ControlFlow) {
													W153 := __e.Get(1)
													_ = W153
													tmp15644 := MakeNative(func(__e *ControlFlow) {
														W154 := __e.Get(1)
														_ = W154
														tmp15681 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W154)

														if True == tmp15681 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp15645 := MakeNative(func(__e *ControlFlow) {
																W155 := __e.Get(1)
																_ = W155
																tmp15678 := Call(__e, PrimFunc(symshen_4hds_a_2), W150, sym_a_a_6)

																if True == tmp15678 {
																	tmp15646 := MakeNative(func(__e *ControlFlow) {
																		W156 := __e.Get(1)
																		_ = W156
																		tmp15675 := PrimIsPair(W156)

																		if True == tmp15675 {
																			tmp15647 := MakeNative(func(__e *ControlFlow) {
																				W157 := __e.Get(1)
																				_ = W157
																				tmp15648 := MakeNative(func(__e *ControlFlow) {
																					W158 := __e.Get(1)
																					_ = W158
																					tmp15671 := PrimIsPair(W158)

																					if True == tmp15671 {
																						tmp15649 := MakeNative(func(__e *ControlFlow) {
																							W159 := __e.Get(1)
																							_ = W159
																							tmp15650 := MakeNative(func(__e *ControlFlow) {
																								W160 := __e.Get(1)
																								_ = W160
																								tmp15667 := PrimEqual(sym_i, W147)

																								var ifres15664 Obj

																								if True == tmp15667 {
																									tmp15666 := PrimEqual(sym_j, W159)

																									var ifres15665 Obj

																									if True == tmp15666 {
																										ifres15665 = True

																									} else {
																										ifres15665 = False

																									}

																									ifres15664 = ifres15665

																								} else {
																									ifres15664 = False

																								}

																								if True == ifres15664 {
																									tmp15651 := PrimCons(W152, Nil)

																									tmp15652 := PrimCons(symlist, tmp15651)

																									tmp15653 := PrimCons(W152, Nil)

																									tmp15654 := PrimCons(symlist, tmp15653)

																									tmp15655 := PrimCons(W157, Nil)

																									tmp15656 := PrimCons(tmp15654, tmp15655)

																									tmp15657 := PrimCons(symstr, tmp15656)

																									tmp15658 := PrimCons(sym_j, Nil)

																									tmp15659 := PrimCons(tmp15657, tmp15658)

																									tmp15660 := PrimCons(sym_1_1_6, tmp15659)

																									tmp15661 := PrimCons(tmp15652, tmp15660)

																									tmp15662 := PrimCons(sym_i, tmp15661)

																									__e.TailApply(PrimFunc(symshen_4comb), W160, tmp15662)
																									return

																								} else {
																									__e.TailApply(PrimFunc(symshen_4parse_1failure))
																									return
																								}

																							}, 1)

																							tmp15668 := Call(__e, PrimFunc(symtail), W158)

																							__e.TailApply(tmp15650, tmp15668)
																							return

																						}, 1)

																						tmp15669 := Call(__e, PrimFunc(symhead), W158)

																						__e.TailApply(tmp15649, tmp15669)
																						return

																					} else {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					}

																				}, 1)

																				tmp15672 := Call(__e, PrimFunc(symtail), W156)

																				__e.TailApply(tmp15648, tmp15672)
																				return

																			}, 1)

																			tmp15673 := Call(__e, PrimFunc(symhead), W156)

																			__e.TailApply(tmp15647, tmp15673)
																			return

																		} else {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		}

																	}, 1)

																	tmp15676 := Call(__e, PrimFunc(symtail), W150)

																	__e.TailApply(tmp15646, tmp15676)
																	return

																} else {
																	__e.TailApply(PrimFunc(symshen_4parse_1failure))
																	return
																}

															}, 1)

															tmp15679 := Call(__e, PrimFunc(symshen_4in_1_6), W154)

															__e.TailApply(tmp15645, tmp15679)
															return

														}

													}, 1)

													tmp15682 := Call(__e, PrimFunc(sym_5end_6), W153)

													__e.TailApply(tmp15644, tmp15682)
													return

												}, 1)

												tmp15683 := Call(__e, PrimFunc(symtail), W151)

												__e.TailApply(tmp15643, tmp15683)
												return

											}, 1)

											tmp15684 := Call(__e, PrimFunc(symhead), W151)

											__e.TailApply(tmp15642, tmp15684)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp15687 := Call(__e, PrimFunc(symtail), W149)

									__e.TailApply(tmp15641, tmp15687)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp15690 := Call(__e, PrimFunc(symtail), W148)

							__e.TailApply(tmp15640, tmp15690)
							return

						}, 1)

						tmp15691 := Call(__e, PrimFunc(symhead), W148)

						__e.TailApply(tmp15639, tmp15691)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15694 := Call(__e, PrimFunc(symtail), V145)

				__e.TailApply(tmp15638, tmp15694)
				return

			}, 1)

			tmp15695 := Call(__e, PrimFunc(symhead), V145)

			tmp15696 := Call(__e, tmp15637, tmp15695)

			ifres15636 = tmp15696

		} else {
			tmp15697 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres15636 = tmp15697

		}

		__e.TailApply(tmp15623, ifres15636)
		return

	}, 1)

	tmp15699 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp15622)

	_ = tmp15699

	tmp15700 := MakeNative(func(__e *ControlFlow) {
		V164 := __e.Get(1)
		_ = V164
		tmp15701 := MakeNative(func(__e *ControlFlow) {
			W165 := __e.Get(1)
			_ = W165
			tmp15720 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W165)

			if True == tmp15720 {
				tmp15702 := MakeNative(func(__e *ControlFlow) {
					W172 := __e.Get(1)
					_ = W172
					tmp15704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W172)

					if True == tmp15704 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W172)
						return
					}

				}, 1)

				tmp15705 := MakeNative(func(__e *ControlFlow) {
					W173 := __e.Get(1)
					_ = W173
					tmp15716 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W173)

					if True == tmp15716 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15706 := MakeNative(func(__e *ControlFlow) {
							W174 := __e.Get(1)
							_ = W174
							tmp15707 := MakeNative(func(__e *ControlFlow) {
								W175 := __e.Get(1)
								_ = W175
								tmp15712 := Call(__e, PrimFunc(symempty_2), W174)

								var ifres15708 Obj

								if True == tmp15712 {
									ifres15708 = Nil

								} else {
									tmp15709 := Call(__e, PrimFunc(symshen_4app), W174, MakeString("\n ..."), symshen_4r)

									tmp15710 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp15709)

									tmp15711 := PrimSimpleError(tmp15710)

									ifres15708 = tmp15711

								}

								__e.TailApply(PrimFunc(symshen_4comb), W175, ifres15708)
								return

							}, 1)

							tmp15713 := Call(__e, PrimFunc(symshen_4in_1_6), W173)

							__e.TailApply(tmp15707, tmp15713)
							return

						}, 1)

						tmp15714 := Call(__e, PrimFunc(symshen_4_5_1out), W173)

						__e.TailApply(tmp15706, tmp15714)
						return

					}

				}, 1)

				tmp15717 := Call(__e, PrimFunc(sym_5_b_6), V164)

				tmp15718 := Call(__e, tmp15705, tmp15717)

				__e.TailApply(tmp15702, tmp15718)
				return

			} else {
				__e.Return(W165)
				return
			}

		}, 1)

		tmp15721 := MakeNative(func(__e *ControlFlow) {
			W166 := __e.Get(1)
			_ = W166
			tmp15736 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W166)

			if True == tmp15736 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp15722 := MakeNative(func(__e *ControlFlow) {
					W167 := __e.Get(1)
					_ = W167
					tmp15723 := MakeNative(func(__e *ControlFlow) {
						W168 := __e.Get(1)
						_ = W168
						tmp15724 := MakeNative(func(__e *ControlFlow) {
							W169 := __e.Get(1)
							_ = W169
							tmp15731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W169)

							if True == tmp15731 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp15725 := MakeNative(func(__e *ControlFlow) {
									W170 := __e.Get(1)
									_ = W170
									tmp15726 := MakeNative(func(__e *ControlFlow) {
										W171 := __e.Get(1)
										_ = W171
										tmp15727 := PrimCons(W167, W170)

										__e.TailApply(PrimFunc(symshen_4comb), W171, tmp15727)
										return

									}, 1)

									tmp15728 := Call(__e, PrimFunc(symshen_4in_1_6), W169)

									__e.TailApply(tmp15726, tmp15728)
									return

								}, 1)

								tmp15729 := Call(__e, PrimFunc(symshen_4_5_1out), W169)

								__e.TailApply(tmp15725, tmp15729)
								return

							}

						}, 1)

						tmp15732 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W168)

						__e.TailApply(tmp15724, tmp15732)
						return

					}, 1)

					tmp15733 := Call(__e, PrimFunc(symshen_4in_1_6), W166)

					__e.TailApply(tmp15723, tmp15733)
					return

				}, 1)

				tmp15734 := Call(__e, PrimFunc(symshen_4_5_1out), W166)

				__e.TailApply(tmp15722, tmp15734)
				return

			}

		}, 1)

		tmp15737 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)

		tmp15738 := Call(__e, tmp15721, tmp15737)

		__e.TailApply(tmp15701, tmp15738)
		return

	}, 1)

	tmp15739 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp15700)

	_ = tmp15739

	tmp15740 := MakeNative(func(__e *ControlFlow) {
		V176 := __e.Get(1)
		_ = V176
		tmp15741 := MakeNative(func(__e *ControlFlow) {
			W177 := __e.Get(1)
			_ = W177
			tmp15764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W177)

			if True == tmp15764 {
				tmp15742 := MakeNative(func(__e *ControlFlow) {
					W186 := __e.Get(1)
					_ = W186
					tmp15744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W186)

					if True == tmp15744 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W186)
						return
					}

				}, 1)

				tmp15745 := MakeNative(func(__e *ControlFlow) {
					W187 := __e.Get(1)
					_ = W187
					tmp15760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W187)

					if True == tmp15760 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15746 := MakeNative(func(__e *ControlFlow) {
							W188 := __e.Get(1)
							_ = W188
							tmp15747 := MakeNative(func(__e *ControlFlow) {
								W189 := __e.Get(1)
								_ = W189
								tmp15748 := MakeNative(func(__e *ControlFlow) {
									W190 := __e.Get(1)
									_ = W190
									tmp15755 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W190)

									if True == tmp15755 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp15749 := MakeNative(func(__e *ControlFlow) {
											W191 := __e.Get(1)
											_ = W191
											tmp15750 := Call(__e, PrimFunc(symshen_4autocomplete), W188)

											tmp15751 := PrimCons(tmp15750, Nil)

											tmp15752 := PrimCons(W188, tmp15751)

											__e.TailApply(PrimFunc(symshen_4comb), W191, tmp15752)
											return

										}, 1)

										tmp15753 := Call(__e, PrimFunc(symshen_4in_1_6), W190)

										__e.TailApply(tmp15749, tmp15753)
										return

									}

								}, 1)

								tmp15756 := Call(__e, PrimFunc(symshen_4_5sc_6), W189)

								__e.TailApply(tmp15748, tmp15756)
								return

							}, 1)

							tmp15757 := Call(__e, PrimFunc(symshen_4in_1_6), W187)

							__e.TailApply(tmp15747, tmp15757)
							return

						}, 1)

						tmp15758 := Call(__e, PrimFunc(symshen_4_5_1out), W187)

						__e.TailApply(tmp15746, tmp15758)
						return

					}

				}, 1)

				tmp15761 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)

				tmp15762 := Call(__e, tmp15745, tmp15761)

				__e.TailApply(tmp15742, tmp15762)
				return

			} else {
				__e.Return(W177)
				return
			}

		}, 1)

		tmp15765 := MakeNative(func(__e *ControlFlow) {
			W178 := __e.Get(1)
			_ = W178
			tmp15787 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W178)

			if True == tmp15787 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp15766 := MakeNative(func(__e *ControlFlow) {
					W179 := __e.Get(1)
					_ = W179
					tmp15767 := MakeNative(func(__e *ControlFlow) {
						W180 := __e.Get(1)
						_ = W180
						tmp15768 := MakeNative(func(__e *ControlFlow) {
							W181 := __e.Get(1)
							_ = W181
							tmp15782 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W181)

							if True == tmp15782 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp15769 := MakeNative(func(__e *ControlFlow) {
									W182 := __e.Get(1)
									_ = W182
									tmp15770 := MakeNative(func(__e *ControlFlow) {
										W183 := __e.Get(1)
										_ = W183
										tmp15771 := MakeNative(func(__e *ControlFlow) {
											W184 := __e.Get(1)
											_ = W184
											tmp15777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W184)

											if True == tmp15777 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp15772 := MakeNative(func(__e *ControlFlow) {
													W185 := __e.Get(1)
													_ = W185
													tmp15773 := PrimCons(W182, Nil)

													tmp15774 := PrimCons(W179, tmp15773)

													__e.TailApply(PrimFunc(symshen_4comb), W185, tmp15774)
													return

												}, 1)

												tmp15775 := Call(__e, PrimFunc(symshen_4in_1_6), W184)

												__e.TailApply(tmp15772, tmp15775)
												return

											}

										}, 1)

										tmp15778 := Call(__e, PrimFunc(symshen_4_5sc_6), W183)

										__e.TailApply(tmp15771, tmp15778)
										return

									}, 1)

									tmp15779 := Call(__e, PrimFunc(symshen_4in_1_6), W181)

									__e.TailApply(tmp15770, tmp15779)
									return

								}, 1)

								tmp15780 := Call(__e, PrimFunc(symshen_4_5_1out), W181)

								__e.TailApply(tmp15769, tmp15780)
								return

							}

						}, 1)

						tmp15783 := Call(__e, PrimFunc(symshen_4_5semantics_6), W180)

						__e.TailApply(tmp15768, tmp15783)
						return

					}, 1)

					tmp15784 := Call(__e, PrimFunc(symshen_4in_1_6), W178)

					__e.TailApply(tmp15767, tmp15784)
					return

				}, 1)

				tmp15785 := Call(__e, PrimFunc(symshen_4_5_1out), W178)

				__e.TailApply(tmp15766, tmp15785)
				return

			}

		}, 1)

		tmp15788 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)

		tmp15789 := Call(__e, tmp15765, tmp15788)

		__e.TailApply(tmp15741, tmp15789)
		return

	}, 1)

	tmp15790 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp15740)

	_ = tmp15790

	tmp15791 := MakeNative(func(__e *ControlFlow) {
		V192 := __e.Get(1)
		_ = V192
		tmp15820 := PrimIsPair(V192)

		var ifres15812 Obj

		if True == tmp15820 {
			tmp15818 := PrimTail(V192)

			tmp15819 := PrimEqual(Nil, tmp15818)

			var ifres15814 Obj

			if True == tmp15819 {
				tmp15816 := PrimHead(V192)

				tmp15817 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp15816)

				var ifres15815 Obj

				if True == tmp15817 {
					ifres15815 = True

				} else {
					ifres15815 = False

				}

				ifres15814 = ifres15815

			} else {
				ifres15814 = False

			}

			var ifres15813 Obj

			if True == ifres15814 {
				ifres15813 = True

			} else {
				ifres15813 = False

			}

			ifres15812 = ifres15813

		} else {
			ifres15812 = False

		}

		if True == ifres15812 {
			__e.Return(PrimHead(V192))
			return
		} else {
			tmp15810 := PrimIsPair(V192)

			var ifres15806 Obj

			if True == tmp15810 {
				tmp15808 := PrimHead(V192)

				tmp15809 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp15808)

				var ifres15807 Obj

				if True == tmp15809 {
					ifres15807 = True

				} else {
					ifres15807 = False

				}

				ifres15806 = ifres15807

			} else {
				ifres15806 = False

			}

			if True == ifres15806 {
				tmp15792 := PrimHead(V192)

				tmp15793 := PrimTail(V192)

				tmp15794 := Call(__e, PrimFunc(symshen_4autocomplete), tmp15793)

				tmp15795 := PrimCons(tmp15794, Nil)

				tmp15796 := PrimCons(tmp15792, tmp15795)

				__e.Return(PrimCons(symappend, tmp15796))
				return

			} else {
				tmp15804 := PrimIsPair(V192)

				if True == tmp15804 {
					tmp15797 := PrimHead(V192)

					tmp15798 := Call(__e, PrimFunc(symshen_4autocomplete), tmp15797)

					tmp15799 := PrimTail(V192)

					tmp15800 := Call(__e, PrimFunc(symshen_4autocomplete), tmp15799)

					tmp15801 := PrimCons(tmp15800, Nil)

					tmp15802 := PrimCons(tmp15798, tmp15801)

					__e.Return(PrimCons(symcons, tmp15802))
					return

				} else {
					__e.Return(V192)
					return
				}

			}

		}

	}, 1)

	tmp15821 := Call(__e, ns2_1set, symshen_4autocomplete, tmp15791)

	_ = tmp15821

	tmp15822 := MakeNative(func(__e *ControlFlow) {
		V193 := __e.Get(1)
		_ = V193
		tmp15829 := PrimIsSymbol(V193)

		if True == tmp15829 {
			tmp15824 := MakeNative(func(__e *ControlFlow) {
				W194 := __e.Get(1)
				_ = W194
				tmp15825 := MakeNative(func(__e *ControlFlow) {
					Z195 := __e.Get(1)
					_ = Z195
					__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
					return
				}, 1)

				__e.TailApply(PrimFunc(symcompile), tmp15825, W194)
				return

			}, 1)

			tmp15826 := Call(__e, PrimFunc(symexplode), V193)

			tmp15827 := Call(__e, tmp15824, tmp15826)

			if True == tmp15827 {
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

	tmp15830 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp15822)

	_ = tmp15830

	tmp15831 := MakeNative(func(__e *ControlFlow) {
		V196 := __e.Get(1)
		_ = V196
		tmp15832 := MakeNative(func(__e *ControlFlow) {
			W197 := __e.Get(1)
			_ = W197
			tmp15854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W197)

			if True == tmp15854 {
				tmp15833 := MakeNative(func(__e *ControlFlow) {
					W202 := __e.Get(1)
					_ = W202
					tmp15845 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W202)

					if True == tmp15845 {
						tmp15834 := MakeNative(func(__e *ControlFlow) {
							W205 := __e.Get(1)
							_ = W205
							tmp15836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W205)

							if True == tmp15836 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W205)
								return
							}

						}, 1)

						tmp15837 := MakeNative(func(__e *ControlFlow) {
							W206 := __e.Get(1)
							_ = W206
							tmp15841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W206)

							if True == tmp15841 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp15838 := MakeNative(func(__e *ControlFlow) {
									W207 := __e.Get(1)
									_ = W207
									__e.TailApply(PrimFunc(symshen_4comb), W207, False)
									return
								}, 1)

								tmp15839 := Call(__e, PrimFunc(symshen_4in_1_6), W206)

								__e.TailApply(tmp15838, tmp15839)
								return

							}

						}, 1)

						tmp15842 := Call(__e, PrimFunc(sym_5_b_6), V196)

						tmp15843 := Call(__e, tmp15837, tmp15842)

						__e.TailApply(tmp15834, tmp15843)
						return

					} else {
						__e.Return(W202)
						return
					}

				}, 1)

				tmp15846 := MakeNative(func(__e *ControlFlow) {
					W203 := __e.Get(1)
					_ = W203
					tmp15850 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W203)

					if True == tmp15850 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15847 := MakeNative(func(__e *ControlFlow) {
							W204 := __e.Get(1)
							_ = W204
							__e.TailApply(PrimFunc(symshen_4comb), W204, True)
							return
						}, 1)

						tmp15848 := Call(__e, PrimFunc(symshen_4in_1_6), W203)

						__e.TailApply(tmp15847, tmp15848)
						return

					}

				}, 1)

				tmp15851 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)

				tmp15852 := Call(__e, tmp15846, tmp15851)

				__e.TailApply(tmp15833, tmp15852)
				return

			} else {
				__e.Return(W197)
				return
			}

		}, 1)

		tmp15855 := MakeNative(func(__e *ControlFlow) {
			W198 := __e.Get(1)
			_ = W198
			tmp15865 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W198)

			if True == tmp15865 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp15856 := MakeNative(func(__e *ControlFlow) {
					W199 := __e.Get(1)
					_ = W199
					tmp15857 := MakeNative(func(__e *ControlFlow) {
						W200 := __e.Get(1)
						_ = W200
						tmp15861 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W200)

						if True == tmp15861 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp15858 := MakeNative(func(__e *ControlFlow) {
								W201 := __e.Get(1)
								_ = W201
								__e.TailApply(PrimFunc(symshen_4comb), W201, True)
								return
							}, 1)

							tmp15859 := Call(__e, PrimFunc(symshen_4in_1_6), W200)

							__e.TailApply(tmp15858, tmp15859)
							return

						}

					}, 1)

					tmp15862 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), W199)

					__e.TailApply(tmp15857, tmp15862)
					return

				}, 1)

				tmp15863 := Call(__e, PrimFunc(symshen_4in_1_6), W198)

				__e.TailApply(tmp15856, tmp15863)
				return

			}

		}, 1)

		tmp15866 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)

		tmp15867 := Call(__e, tmp15855, tmp15866)

		__e.TailApply(tmp15832, tmp15867)
		return

	}, 1)

	tmp15868 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp15831)

	_ = tmp15868

	tmp15869 := MakeNative(func(__e *ControlFlow) {
		V208 := __e.Get(1)
		_ = V208
		tmp15870 := MakeNative(func(__e *ControlFlow) {
			W209 := __e.Get(1)
			_ = W209
			tmp15886 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W209)

			if True == tmp15886 {
				tmp15871 := MakeNative(func(__e *ControlFlow) {
					W215 := __e.Get(1)
					_ = W215
					tmp15873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W215)

					if True == tmp15873 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W215)
						return
					}

				}, 1)

				tmp15874 := MakeNative(func(__e *ControlFlow) {
					W216 := __e.Get(1)
					_ = W216
					tmp15882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W216)

					if True == tmp15882 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15875 := MakeNative(func(__e *ControlFlow) {
							W217 := __e.Get(1)
							_ = W217
							tmp15879 := Call(__e, PrimFunc(symshen_4hds_a_2), W217, MakeString("."))

							if True == tmp15879 {
								tmp15876 := MakeNative(func(__e *ControlFlow) {
									W218 := __e.Get(1)
									_ = W218
									__e.TailApply(PrimFunc(symshen_4comb), W218, symshen_4skip)
									return
								}, 1)

								tmp15877 := Call(__e, PrimFunc(symtail), W217)

								__e.TailApply(tmp15876, tmp15877)
								return

							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp15880 := Call(__e, PrimFunc(symshen_4in_1_6), W216)

						__e.TailApply(tmp15875, tmp15880)
						return

					}

				}, 1)

				tmp15883 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)

				tmp15884 := Call(__e, tmp15874, tmp15883)

				__e.TailApply(tmp15871, tmp15884)
				return

			} else {
				__e.Return(W209)
				return
			}

		}, 1)

		tmp15887 := MakeNative(func(__e *ControlFlow) {
			W210 := __e.Get(1)
			_ = W210
			tmp15901 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W210)

			if True == tmp15901 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp15888 := MakeNative(func(__e *ControlFlow) {
					W211 := __e.Get(1)
					_ = W211
					tmp15898 := Call(__e, PrimFunc(symshen_4hds_a_2), W211, MakeString("."))

					if True == tmp15898 {
						tmp15889 := MakeNative(func(__e *ControlFlow) {
							W212 := __e.Get(1)
							_ = W212
							tmp15890 := MakeNative(func(__e *ControlFlow) {
								W213 := __e.Get(1)
								_ = W213
								tmp15894 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W213)

								if True == tmp15894 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp15891 := MakeNative(func(__e *ControlFlow) {
										W214 := __e.Get(1)
										_ = W214
										__e.TailApply(PrimFunc(symshen_4comb), W214, symshen_4skip)
										return
									}, 1)

									tmp15892 := Call(__e, PrimFunc(symshen_4in_1_6), W213)

									__e.TailApply(tmp15891, tmp15892)
									return

								}

							}, 1)

							tmp15895 := Call(__e, PrimFunc(symshen_4_5packagenames_6), W212)

							__e.TailApply(tmp15890, tmp15895)
							return

						}, 1)

						tmp15896 := Call(__e, PrimFunc(symtail), W211)

						__e.TailApply(tmp15889, tmp15896)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15899 := Call(__e, PrimFunc(symshen_4in_1_6), W210)

				__e.TailApply(tmp15888, tmp15899)
				return

			}

		}, 1)

		tmp15902 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)

		tmp15903 := Call(__e, tmp15887, tmp15902)

		__e.TailApply(tmp15870, tmp15903)
		return

	}, 1)

	tmp15904 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp15869)

	_ = tmp15904

	tmp15905 := MakeNative(func(__e *ControlFlow) {
		V219 := __e.Get(1)
		_ = V219
		tmp15906 := MakeNative(func(__e *ControlFlow) {
			W220 := __e.Get(1)
			_ = W220
			tmp15918 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W220)

			if True == tmp15918 {
				tmp15907 := MakeNative(func(__e *ControlFlow) {
					W225 := __e.Get(1)
					_ = W225
					tmp15909 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W225)

					if True == tmp15909 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W225)
						return
					}

				}, 1)

				tmp15910 := MakeNative(func(__e *ControlFlow) {
					W226 := __e.Get(1)
					_ = W226
					tmp15914 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W226)

					if True == tmp15914 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15911 := MakeNative(func(__e *ControlFlow) {
							W227 := __e.Get(1)
							_ = W227
							__e.TailApply(PrimFunc(symshen_4comb), W227, symshen_4skip)
							return
						}, 1)

						tmp15912 := Call(__e, PrimFunc(symshen_4in_1_6), W226)

						__e.TailApply(tmp15911, tmp15912)
						return

					}

				}, 1)

				tmp15915 := Call(__e, PrimFunc(sym_5e_6), V219)

				tmp15916 := Call(__e, tmp15910, tmp15915)

				__e.TailApply(tmp15907, tmp15916)
				return

			} else {
				__e.Return(W220)
				return
			}

		}, 1)

		tmp15919 := MakeNative(func(__e *ControlFlow) {
			W221 := __e.Get(1)
			_ = W221
			tmp15929 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W221)

			if True == tmp15929 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp15920 := MakeNative(func(__e *ControlFlow) {
					W222 := __e.Get(1)
					_ = W222
					tmp15921 := MakeNative(func(__e *ControlFlow) {
						W223 := __e.Get(1)
						_ = W223
						tmp15925 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W223)

						if True == tmp15925 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp15922 := MakeNative(func(__e *ControlFlow) {
								W224 := __e.Get(1)
								_ = W224
								__e.TailApply(PrimFunc(symshen_4comb), W224, symshen_4skip)
								return
							}, 1)

							tmp15923 := Call(__e, PrimFunc(symshen_4in_1_6), W223)

							__e.TailApply(tmp15922, tmp15923)
							return

						}

					}, 1)

					tmp15926 := Call(__e, PrimFunc(symshen_4_5packagename_6), W222)

					__e.TailApply(tmp15921, tmp15926)
					return

				}, 1)

				tmp15927 := Call(__e, PrimFunc(symshen_4in_1_6), W221)

				__e.TailApply(tmp15920, tmp15927)
				return

			}

		}, 1)

		tmp15930 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)

		tmp15931 := Call(__e, tmp15919, tmp15930)

		__e.TailApply(tmp15906, tmp15931)
		return

	}, 1)

	tmp15932 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp15905)

	_ = tmp15932

	tmp15933 := MakeNative(func(__e *ControlFlow) {
		V228 := __e.Get(1)
		_ = V228
		tmp15934 := MakeNative(func(__e *ControlFlow) {
			W229 := __e.Get(1)
			_ = W229
			tmp15936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W229)

			if True == tmp15936 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W229)
				return
			}

		}, 1)

		tmp15947 := PrimIsPair(V228)

		var ifres15937 Obj

		if True == tmp15947 {
			tmp15938 := MakeNative(func(__e *ControlFlow) {
				W230 := __e.Get(1)
				_ = W230
				tmp15939 := MakeNative(func(__e *ControlFlow) {
					W231 := __e.Get(1)
					_ = W231
					tmp15941 := PrimEqual(W230, MakeString("."))

					tmp15942 := PrimNot(tmp15941)

					if True == tmp15942 {
						__e.TailApply(PrimFunc(symshen_4comb), W231, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15943 := Call(__e, PrimFunc(symtail), V228)

				__e.TailApply(tmp15939, tmp15943)
				return

			}, 1)

			tmp15944 := Call(__e, PrimFunc(symhead), V228)

			tmp15945 := Call(__e, tmp15938, tmp15944)

			ifres15937 = tmp15945

		} else {
			tmp15946 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres15937 = tmp15946

		}

		__e.TailApply(tmp15934, ifres15937)
		return

	}, 1)

	tmp15948 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp15933)

	_ = tmp15948

	tmp15949 := MakeNative(func(__e *ControlFlow) {
		V232 := __e.Get(1)
		_ = V232
		tmp15950 := MakeNative(func(__e *ControlFlow) {
			W233 := __e.Get(1)
			_ = W233
			tmp15952 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W233)

			if True == tmp15952 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W233)
				return
			}

		}, 1)

		tmp15975 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))

		var ifres15953 Obj

		if True == tmp15975 {
			tmp15954 := MakeNative(func(__e *ControlFlow) {
				W234 := __e.Get(1)
				_ = W234
				tmp15955 := MakeNative(func(__e *ControlFlow) {
					W235 := __e.Get(1)
					_ = W235
					tmp15970 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W235)

					if True == tmp15970 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp15956 := MakeNative(func(__e *ControlFlow) {
							W236 := __e.Get(1)
							_ = W236
							tmp15957 := MakeNative(func(__e *ControlFlow) {
								W237 := __e.Get(1)
								_ = W237
								tmp15959 := MakeNative(func(__e *ControlFlow) {
									W238 := __e.Get(1)
									_ = W238
									tmp15964 := PrimIsPair(W238)

									if True == tmp15964 {
										tmp15961 := PrimHead(W238)

										tmp15962 := PrimEqual(tmp15961, MakeString(">"))

										if True == tmp15962 {
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

								tmp15965 := Call(__e, PrimFunc(symreverse), W236)

								tmp15966 := Call(__e, tmp15959, tmp15965)

								if True == tmp15966 {
									__e.TailApply(PrimFunc(symshen_4comb), W237, symshen_4skip)
									return
								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp15967 := Call(__e, PrimFunc(symshen_4in_1_6), W235)

							__e.TailApply(tmp15957, tmp15967)
							return

						}, 1)

						tmp15968 := Call(__e, PrimFunc(symshen_4_5_1out), W235)

						__e.TailApply(tmp15956, tmp15968)
						return

					}

				}, 1)

				tmp15971 := Call(__e, PrimFunc(sym_5_b_6), W234)

				__e.TailApply(tmp15955, tmp15971)
				return

			}, 1)

			tmp15972 := Call(__e, PrimFunc(symtail), V232)

			tmp15973 := Call(__e, tmp15954, tmp15972)

			ifres15953 = tmp15973

		} else {
			tmp15974 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres15953 = tmp15974

		}

		__e.TailApply(tmp15950, ifres15953)
		return

	}, 1)

	tmp15976 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp15949)

	_ = tmp15976

	tmp15977 := MakeNative(func(__e *ControlFlow) {
		V239 := __e.Get(1)
		_ = V239
		tmp15978 := PrimIntern(MakeString(";"))

		__e.Return(PrimEqual(V239, tmp15978))
		return

	}, 1)

	tmp15979 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp15977)

	_ = tmp15979

	tmp15980 := MakeNative(func(__e *ControlFlow) {
		V240 := __e.Get(1)
		_ = V240
		tmp15981 := MakeNative(func(__e *ControlFlow) {
			W241 := __e.Get(1)
			_ = W241
			tmp15983 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W241)

			if True == tmp15983 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W241)
				return
			}

		}, 1)

		tmp15993 := PrimIsPair(V240)

		var ifres15984 Obj

		if True == tmp15993 {
			tmp15985 := MakeNative(func(__e *ControlFlow) {
				W242 := __e.Get(1)
				_ = W242
				tmp15986 := MakeNative(func(__e *ControlFlow) {
					W243 := __e.Get(1)
					_ = W243
					tmp15988 := Call(__e, PrimFunc(symshen_4colon_1equal_2), W242)

					if True == tmp15988 {
						__e.TailApply(PrimFunc(symshen_4comb), W243, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15989 := Call(__e, PrimFunc(symtail), V240)

				__e.TailApply(tmp15986, tmp15989)
				return

			}, 1)

			tmp15990 := Call(__e, PrimFunc(symhead), V240)

			tmp15991 := Call(__e, tmp15985, tmp15990)

			ifres15984 = tmp15991

		} else {
			tmp15992 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres15984 = tmp15992

		}

		__e.TailApply(tmp15981, ifres15984)
		return

	}, 1)

	tmp15994 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp15980)

	_ = tmp15994

	tmp15995 := MakeNative(func(__e *ControlFlow) {
		V244 := __e.Get(1)
		_ = V244
		tmp15996 := PrimIntern(MakeString(":="))

		__e.Return(PrimEqual(tmp15996, V244))
		return

	}, 1)

	tmp15997 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp15995)

	_ = tmp15997

	tmp15998 := MakeNative(func(__e *ControlFlow) {
		V245 := __e.Get(1)
		_ = V245
		tmp15999 := MakeNative(func(__e *ControlFlow) {
			W246 := __e.Get(1)
			_ = W246
			tmp16014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W246)

			if True == tmp16014 {
				tmp16000 := MakeNative(func(__e *ControlFlow) {
					W253 := __e.Get(1)
					_ = W253
					tmp16002 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W253)

					if True == tmp16002 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W253)
						return
					}

				}, 1)

				tmp16003 := MakeNative(func(__e *ControlFlow) {
					W254 := __e.Get(1)
					_ = W254
					tmp16010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W254)

					if True == tmp16010 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp16004 := MakeNative(func(__e *ControlFlow) {
							W255 := __e.Get(1)
							_ = W255
							tmp16005 := MakeNative(func(__e *ControlFlow) {
								W256 := __e.Get(1)
								_ = W256
								tmp16006 := PrimCons(W255, Nil)

								__e.TailApply(PrimFunc(symshen_4comb), W256, tmp16006)
								return

							}, 1)

							tmp16007 := Call(__e, PrimFunc(symshen_4in_1_6), W254)

							__e.TailApply(tmp16005, tmp16007)
							return

						}, 1)

						tmp16008 := Call(__e, PrimFunc(symshen_4_5_1out), W254)

						__e.TailApply(tmp16004, tmp16008)
						return

					}

				}, 1)

				tmp16011 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)

				tmp16012 := Call(__e, tmp16003, tmp16011)

				__e.TailApply(tmp16000, tmp16012)
				return

			} else {
				__e.Return(W246)
				return
			}

		}, 1)

		tmp16015 := MakeNative(func(__e *ControlFlow) {
			W247 := __e.Get(1)
			_ = W247
			tmp16030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W247)

			if True == tmp16030 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp16016 := MakeNative(func(__e *ControlFlow) {
					W248 := __e.Get(1)
					_ = W248
					tmp16017 := MakeNative(func(__e *ControlFlow) {
						W249 := __e.Get(1)
						_ = W249
						tmp16018 := MakeNative(func(__e *ControlFlow) {
							W250 := __e.Get(1)
							_ = W250
							tmp16025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W250)

							if True == tmp16025 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp16019 := MakeNative(func(__e *ControlFlow) {
									W251 := __e.Get(1)
									_ = W251
									tmp16020 := MakeNative(func(__e *ControlFlow) {
										W252 := __e.Get(1)
										_ = W252
										tmp16021 := PrimCons(W248, W251)

										__e.TailApply(PrimFunc(symshen_4comb), W252, tmp16021)
										return

									}, 1)

									tmp16022 := Call(__e, PrimFunc(symshen_4in_1_6), W250)

									__e.TailApply(tmp16020, tmp16022)
									return

								}, 1)

								tmp16023 := Call(__e, PrimFunc(symshen_4_5_1out), W250)

								__e.TailApply(tmp16019, tmp16023)
								return

							}

						}, 1)

						tmp16026 := Call(__e, PrimFunc(symshen_4_5syntax_6), W249)

						__e.TailApply(tmp16018, tmp16026)
						return

					}, 1)

					tmp16027 := Call(__e, PrimFunc(symshen_4in_1_6), W247)

					__e.TailApply(tmp16017, tmp16027)
					return

				}, 1)

				tmp16028 := Call(__e, PrimFunc(symshen_4_5_1out), W247)

				__e.TailApply(tmp16016, tmp16028)
				return

			}

		}, 1)

		tmp16031 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)

		tmp16032 := Call(__e, tmp16015, tmp16031)

		__e.TailApply(tmp15999, tmp16032)
		return

	}, 1)

	tmp16033 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp15998)

	_ = tmp16033

	tmp16034 := MakeNative(func(__e *ControlFlow) {
		V257 := __e.Get(1)
		_ = V257
		tmp16035 := MakeNative(func(__e *ControlFlow) {
			W258 := __e.Get(1)
			_ = W258
			tmp16037 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W258)

			if True == tmp16037 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W258)
				return
			}

		}, 1)

		tmp16047 := PrimIsPair(V257)

		var ifres16038 Obj

		if True == tmp16047 {
			tmp16039 := MakeNative(func(__e *ControlFlow) {
				W259 := __e.Get(1)
				_ = W259
				tmp16040 := MakeNative(func(__e *ControlFlow) {
					W260 := __e.Get(1)
					_ = W260
					tmp16042 := Call(__e, PrimFunc(symshen_4syntax_1item_2), W259)

					if True == tmp16042 {
						__e.TailApply(PrimFunc(symshen_4comb), W260, W259)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp16043 := Call(__e, PrimFunc(symtail), V257)

				__e.TailApply(tmp16040, tmp16043)
				return

			}, 1)

			tmp16044 := Call(__e, PrimFunc(symhead), V257)

			tmp16045 := Call(__e, tmp16039, tmp16044)

			ifres16038 = tmp16045

		} else {
			tmp16046 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres16038 = tmp16046

		}

		__e.TailApply(tmp16035, ifres16038)
		return

	}, 1)

	tmp16048 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp16034)

	_ = tmp16048

	tmp16049 := MakeNative(func(__e *ControlFlow) {
		V263 := __e.Get(1)
		_ = V263
		tmp16085 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)

		if True == tmp16085 {
			__e.Return(False)
			return
		} else {
			tmp16083 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)

			if True == tmp16083 {
				__e.Return(False)
				return
			} else {
				tmp16081 := Call(__e, PrimFunc(symatom_2), V263)

				if True == tmp16081 {
					__e.Return(True)
					return
				} else {
					tmp16079 := PrimIsPair(V263)

					var ifres16060 Obj

					if True == tmp16079 {
						tmp16077 := PrimHead(V263)

						tmp16078 := PrimEqual(symcons, tmp16077)

						var ifres16062 Obj

						if True == tmp16078 {
							tmp16075 := PrimTail(V263)

							tmp16076 := PrimIsPair(tmp16075)

							var ifres16064 Obj

							if True == tmp16076 {
								tmp16072 := PrimTail(V263)

								tmp16073 := PrimTail(tmp16072)

								tmp16074 := PrimIsPair(tmp16073)

								var ifres16066 Obj

								if True == tmp16074 {
									tmp16068 := PrimTail(V263)

									tmp16069 := PrimTail(tmp16068)

									tmp16070 := PrimTail(tmp16069)

									tmp16071 := PrimEqual(Nil, tmp16070)

									var ifres16067 Obj

									if True == tmp16071 {
										ifres16067 = True

									} else {
										ifres16067 = False

									}

									ifres16066 = ifres16067

								} else {
									ifres16066 = False

								}

								var ifres16065 Obj

								if True == ifres16066 {
									ifres16065 = True

								} else {
									ifres16065 = False

								}

								ifres16064 = ifres16065

							} else {
								ifres16064 = False

							}

							var ifres16063 Obj

							if True == ifres16064 {
								ifres16063 = True

							} else {
								ifres16063 = False

							}

							ifres16062 = ifres16063

						} else {
							ifres16062 = False

						}

						var ifres16061 Obj

						if True == ifres16062 {
							ifres16061 = True

						} else {
							ifres16061 = False

						}

						ifres16060 = ifres16061

					} else {
						ifres16060 = False

					}

					if True == ifres16060 {
						tmp16056 := PrimTail(V263)

						tmp16057 := PrimHead(tmp16056)

						tmp16058 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp16057)

						if True == tmp16058 {
							tmp16051 := PrimTail(V263)

							tmp16052 := PrimTail(tmp16051)

							tmp16053 := PrimHead(tmp16052)

							tmp16054 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp16053)

							if True == tmp16054 {
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

				}

			}

		}

	}, 1)

	tmp16086 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp16049)

	_ = tmp16086

	tmp16087 := MakeNative(func(__e *ControlFlow) {
		V264 := __e.Get(1)
		_ = V264
		tmp16088 := MakeNative(func(__e *ControlFlow) {
			W265 := __e.Get(1)
			_ = W265
			tmp16109 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W265)

			if True == tmp16109 {
				tmp16089 := MakeNative(func(__e *ControlFlow) {
					W273 := __e.Get(1)
					_ = W273
					tmp16091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W273)

					if True == tmp16091 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W273)
						return
					}

				}, 1)

				tmp16092 := MakeNative(func(__e *ControlFlow) {
					W274 := __e.Get(1)
					_ = W274
					tmp16105 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W274)

					if True == tmp16105 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp16093 := MakeNative(func(__e *ControlFlow) {
							W275 := __e.Get(1)
							_ = W275
							tmp16102 := PrimIsPair(W275)

							if True == tmp16102 {
								tmp16094 := MakeNative(func(__e *ControlFlow) {
									W276 := __e.Get(1)
									_ = W276
									tmp16095 := MakeNative(func(__e *ControlFlow) {
										W277 := __e.Get(1)
										_ = W277
										tmp16097 := Call(__e, PrimFunc(symshen_4semicolon_2), W276)

										tmp16098 := PrimNot(tmp16097)

										if True == tmp16098 {
											__e.TailApply(PrimFunc(symshen_4comb), W277, W276)
											return
										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp16099 := Call(__e, PrimFunc(symtail), W275)

									__e.TailApply(tmp16095, tmp16099)
									return

								}, 1)

								tmp16100 := Call(__e, PrimFunc(symhead), W275)

								__e.TailApply(tmp16094, tmp16100)
								return

							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp16103 := Call(__e, PrimFunc(symshen_4in_1_6), W274)

						__e.TailApply(tmp16093, tmp16103)
						return

					}

				}, 1)

				tmp16106 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)

				tmp16107 := Call(__e, tmp16092, tmp16106)

				__e.TailApply(tmp16089, tmp16107)
				return

			} else {
				__e.Return(W265)
				return
			}

		}, 1)

		tmp16110 := MakeNative(func(__e *ControlFlow) {
			W266 := __e.Get(1)
			_ = W266
			tmp16136 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W266)

			if True == tmp16136 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp16111 := MakeNative(func(__e *ControlFlow) {
					W267 := __e.Get(1)
					_ = W267
					tmp16133 := PrimIsPair(W267)

					if True == tmp16133 {
						tmp16112 := MakeNative(func(__e *ControlFlow) {
							W268 := __e.Get(1)
							_ = W268
							tmp16113 := MakeNative(func(__e *ControlFlow) {
								W269 := __e.Get(1)
								_ = W269
								tmp16129 := Call(__e, PrimFunc(symshen_4hds_a_2), W269, symwhere)

								if True == tmp16129 {
									tmp16114 := MakeNative(func(__e *ControlFlow) {
										W270 := __e.Get(1)
										_ = W270
										tmp16126 := PrimIsPair(W270)

										if True == tmp16126 {
											tmp16115 := MakeNative(func(__e *ControlFlow) {
												W271 := __e.Get(1)
												_ = W271
												tmp16116 := MakeNative(func(__e *ControlFlow) {
													W272 := __e.Get(1)
													_ = W272
													tmp16121 := Call(__e, PrimFunc(symshen_4semicolon_2), W268)

													tmp16122 := PrimNot(tmp16121)

													if True == tmp16122 {
														tmp16117 := PrimCons(W268, Nil)

														tmp16118 := PrimCons(W271, tmp16117)

														tmp16119 := PrimCons(symwhere, tmp16118)

														__e.TailApply(PrimFunc(symshen_4comb), W272, tmp16119)
														return

													} else {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp16123 := Call(__e, PrimFunc(symtail), W270)

												__e.TailApply(tmp16116, tmp16123)
												return

											}, 1)

											tmp16124 := Call(__e, PrimFunc(symhead), W270)

											__e.TailApply(tmp16115, tmp16124)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp16127 := Call(__e, PrimFunc(symtail), W269)

									__e.TailApply(tmp16114, tmp16127)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp16130 := Call(__e, PrimFunc(symtail), W267)

							__e.TailApply(tmp16113, tmp16130)
							return

						}, 1)

						tmp16131 := Call(__e, PrimFunc(symhead), W267)

						__e.TailApply(tmp16112, tmp16131)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp16134 := Call(__e, PrimFunc(symshen_4in_1_6), W266)

				__e.TailApply(tmp16111, tmp16134)
				return

			}

		}, 1)

		tmp16137 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)

		tmp16138 := Call(__e, tmp16110, tmp16137)

		__e.TailApply(tmp16088, tmp16138)
		return

	}, 1)

	tmp16139 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp16087)

	_ = tmp16139

	tmp16140 := MakeNative(func(__e *ControlFlow) {
		V286 := __e.Get(1)
		_ = V286
		V287 := __e.Get(2)
		_ = V287
		V288 := __e.Get(3)
		_ = V288
		tmp16148 := PrimEqual(Nil, V288)

		if True == tmp16148 {
			__e.Return(PrimCons(symshen_4parse_1failure, Nil))
			return
		} else {
			tmp16146 := PrimIsPair(V288)

			if True == tmp16146 {
				tmp16141 := PrimHead(V288)

				tmp16142 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp16141, V287)

				tmp16143 := PrimTail(V288)

				tmp16144 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp16143)

				__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp16142, tmp16144)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
				return
			}

		}

	}, 3)

	tmp16149 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp16140)

	_ = tmp16149

	tmp16150 := MakeNative(func(__e *ControlFlow) {
		__e.TailApply(PrimFunc(symfail))
		return
	}, 0)

	tmp16151 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp16150)

	_ = tmp16151

	tmp16152 := MakeNative(func(__e *ControlFlow) {
		V289 := __e.Get(1)
		_ = V289
		V290 := __e.Get(2)
		_ = V290
		tmp16153 := PrimCons(symResult, Nil)

		tmp16154 := PrimCons(symshen_4parse_1failure_2, tmp16153)

		tmp16155 := PrimCons(symResult, Nil)

		tmp16156 := PrimCons(V290, tmp16155)

		tmp16157 := PrimCons(tmp16154, tmp16156)

		tmp16158 := PrimCons(symif, tmp16157)

		tmp16159 := PrimCons(tmp16158, Nil)

		tmp16160 := PrimCons(V289, tmp16159)

		tmp16161 := PrimCons(symResult, tmp16160)

		__e.Return(PrimCons(symlet, tmp16161))
		return

	}, 2)

	tmp16162 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp16152)

	_ = tmp16162

	tmp16163 := MakeNative(func(__e *ControlFlow) {
		V297 := __e.Get(1)
		_ = V297
		V298 := __e.Get(2)
		_ = V298
		V299 := __e.Get(3)
		_ = V299
		tmp16177 := PrimIsPair(V298)

		var ifres16168 Obj

		if True == tmp16177 {
			tmp16175 := PrimTail(V298)

			tmp16176 := PrimIsPair(tmp16175)

			var ifres16170 Obj

			if True == tmp16176 {
				tmp16172 := PrimTail(V298)

				tmp16173 := PrimTail(tmp16172)

				tmp16174 := PrimEqual(Nil, tmp16173)

				var ifres16171 Obj

				if True == tmp16174 {
					ifres16171 = True

				} else {
					ifres16171 = False

				}

				ifres16170 = ifres16171

			} else {
				ifres16170 = False

			}

			var ifres16169 Obj

			if True == ifres16170 {
				ifres16169 = True

			} else {
				ifres16169 = False

			}

			ifres16168 = ifres16169

		} else {
			ifres16168 = False

		}

		if True == ifres16168 {
			tmp16164 := PrimHead(V298)

			tmp16165 := PrimTail(V298)

			tmp16166 := PrimHead(tmp16165)

			__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp16164, tmp16166)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
			return
		}

	}, 3)

	tmp16178 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp16163)

	_ = tmp16178

	tmp16179 := MakeNative(func(__e *ControlFlow) {
		V308 := __e.Get(1)
		_ = V308
		V309 := __e.Get(2)
		_ = V309
		V310 := __e.Get(3)
		_ = V310
		V311 := __e.Get(4)
		_ = V311
		tmp16243 := PrimEqual(Nil, V310)

		var ifres16221 Obj

		if True == tmp16243 {
			tmp16242 := PrimIsPair(V311)

			var ifres16223 Obj

			if True == tmp16242 {
				tmp16240 := PrimHead(V311)

				tmp16241 := PrimEqual(symwhere, tmp16240)

				var ifres16225 Obj

				if True == tmp16241 {
					tmp16238 := PrimTail(V311)

					tmp16239 := PrimIsPair(tmp16238)

					var ifres16227 Obj

					if True == tmp16239 {
						tmp16235 := PrimTail(V311)

						tmp16236 := PrimTail(tmp16235)

						tmp16237 := PrimIsPair(tmp16236)

						var ifres16229 Obj

						if True == tmp16237 {
							tmp16231 := PrimTail(V311)

							tmp16232 := PrimTail(tmp16231)

							tmp16233 := PrimTail(tmp16232)

							tmp16234 := PrimEqual(Nil, tmp16233)

							var ifres16230 Obj

							if True == tmp16234 {
								ifres16230 = True

							} else {
								ifres16230 = False

							}

							ifres16229 = ifres16230

						} else {
							ifres16229 = False

						}

						var ifres16228 Obj

						if True == ifres16229 {
							ifres16228 = True

						} else {
							ifres16228 = False

						}

						ifres16227 = ifres16228

					} else {
						ifres16227 = False

					}

					var ifres16226 Obj

					if True == ifres16227 {
						ifres16226 = True

					} else {
						ifres16226 = False

					}

					ifres16225 = ifres16226

				} else {
					ifres16225 = False

				}

				var ifres16224 Obj

				if True == ifres16225 {
					ifres16224 = True

				} else {
					ifres16224 = False

				}

				ifres16223 = ifres16224

			} else {
				ifres16223 = False

			}

			var ifres16222 Obj

			if True == ifres16223 {
				ifres16222 = True

			} else {
				ifres16222 = False

			}

			ifres16221 = ifres16222

		} else {
			ifres16221 = False

		}

		if True == ifres16221 {
			tmp16180 := PrimTail(V311)

			tmp16181 := PrimHead(tmp16180)

			tmp16182 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp16181)

			tmp16183 := PrimTail(V311)

			tmp16184 := PrimTail(tmp16183)

			tmp16185 := PrimHead(tmp16184)

			tmp16186 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp16185)

			tmp16187 := PrimCons(symshen_4parse_1failure, Nil)

			tmp16188 := PrimCons(tmp16187, Nil)

			tmp16189 := PrimCons(tmp16186, tmp16188)

			tmp16190 := PrimCons(tmp16182, tmp16189)

			__e.Return(PrimCons(symif, tmp16190))
			return

		} else {
			tmp16219 := PrimEqual(Nil, V310)

			if True == tmp16219 {
				__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
				return
			} else {
				tmp16217 := PrimIsPair(V310)

				if True == tmp16217 {
					tmp16214 := PrimHead(V310)

					tmp16215 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16214)

					if True == tmp16215 {
						tmp16191 := PrimHead(V310)

						tmp16192 := PrimTail(V310)

						__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp16191, tmp16192, V311)
						return

					} else {
						tmp16211 := PrimHead(V310)

						tmp16212 := PrimIsVariable(tmp16211)

						if True == tmp16212 {
							tmp16193 := PrimHead(V310)

							tmp16194 := PrimTail(V310)

							__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp16193, tmp16194, V311)
							return

						} else {
							tmp16208 := PrimHead(V310)

							tmp16209 := PrimEqual(sym__, tmp16208)

							if True == tmp16209 {
								tmp16195 := PrimHead(V310)

								tmp16196 := PrimTail(V310)

								__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp16195, tmp16196, V311)
								return

							} else {
								tmp16205 := PrimHead(V310)

								tmp16206 := Call(__e, PrimFunc(symatom_2), tmp16205)

								if True == tmp16206 {
									tmp16197 := PrimHead(V310)

									tmp16198 := PrimTail(V310)

									__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp16197, tmp16198, V311)
									return

								} else {
									tmp16202 := PrimHead(V310)

									tmp16203 := PrimIsPair(tmp16202)

									if True == tmp16203 {
										tmp16199 := PrimHead(V310)

										tmp16200 := PrimTail(V310)

										__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp16199, tmp16200, V311)
										return

									} else {
										__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
										return
									}

								}

							}

						}

					}

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
					return
				}

			}

		}

	}, 4)

	tmp16244 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp16179)

	_ = tmp16244

	tmp16245 := MakeNative(func(__e *ControlFlow) {
		V312 := __e.Get(1)
		_ = V312
		V313 := __e.Get(2)
		_ = V313
		V314 := __e.Get(3)
		_ = V314
		V315 := __e.Get(4)
		_ = V315
		V316 := __e.Get(5)
		_ = V316
		tmp16246 := MakeNative(func(__e *ControlFlow) {
			W317 := __e.Get(1)
			_ = W317
			tmp16247 := MakeNative(func(__e *ControlFlow) {
				W318 := __e.Get(1)
				_ = W318
				tmp16248 := MakeNative(func(__e *ControlFlow) {
					W319 := __e.Get(1)
					_ = W319
					tmp16249 := PrimCons(V313, Nil)

					tmp16250 := PrimCons(V314, tmp16249)

					tmp16251 := PrimCons(W317, Nil)

					tmp16252 := PrimCons(symshen_4parse_1failure_2, tmp16251)

					tmp16253 := PrimCons(symshen_4parse_1failure, Nil)

					tmp16254 := MakeNative(func(__e *ControlFlow) {
						W320 := __e.Get(1)
						_ = W320
						tmp16264 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)

						var ifres16261 Obj

						if True == tmp16264 {
							ifres16261 = True

						} else {
							tmp16263 := Call(__e, PrimFunc(symshen_4occurs_1check_2), W318, V316)

							var ifres16262 Obj

							if True == tmp16263 {
								ifres16262 = True

							} else {
								ifres16262 = False

							}

							ifres16261 = ifres16262

						}

						if True == ifres16261 {
							tmp16255 := PrimCons(W317, Nil)

							tmp16256 := PrimCons(symshen_4_5_1out, tmp16255)

							tmp16257 := PrimCons(W320, Nil)

							tmp16258 := PrimCons(tmp16256, tmp16257)

							tmp16259 := PrimCons(W318, tmp16258)

							__e.Return(PrimCons(symlet, tmp16259))
							return

						} else {
							__e.Return(W320)
							return
						}

					}, 1)

					tmp16265 := PrimCons(W317, Nil)

					tmp16266 := PrimCons(symshen_4in_1_6, tmp16265)

					tmp16267 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, W319, V315, V316)

					tmp16268 := PrimCons(tmp16267, Nil)

					tmp16269 := PrimCons(tmp16266, tmp16268)

					tmp16270 := PrimCons(W319, tmp16269)

					tmp16271 := PrimCons(symlet, tmp16270)

					tmp16272 := Call(__e, tmp16254, tmp16271)

					tmp16273 := PrimCons(tmp16272, Nil)

					tmp16274 := PrimCons(tmp16253, tmp16273)

					tmp16275 := PrimCons(tmp16252, tmp16274)

					tmp16276 := PrimCons(symif, tmp16275)

					tmp16277 := PrimCons(tmp16276, Nil)

					tmp16278 := PrimCons(tmp16250, tmp16277)

					tmp16279 := PrimCons(W317, tmp16278)

					__e.Return(PrimCons(symlet, tmp16279))
					return

				}, 1)

				tmp16280 := Call(__e, PrimFunc(symconcat), symRemainder, V314)

				__e.TailApply(tmp16248, tmp16280)
				return

			}, 1)

			tmp16281 := Call(__e, PrimFunc(symconcat), symAction, V314)

			__e.TailApply(tmp16247, tmp16281)
			return

		}, 1)

		tmp16282 := Call(__e, PrimFunc(symconcat), symParse, V314)

		__e.TailApply(tmp16246, tmp16282)
		return

	}, 5)

	tmp16283 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp16245)

	_ = tmp16283

	tmp16284 := MakeNative(func(__e *ControlFlow) {
		V321 := __e.Get(1)
		_ = V321
		V322 := __e.Get(2)
		_ = V322
		V323 := __e.Get(3)
		_ = V323
		V324 := __e.Get(4)
		_ = V324
		V325 := __e.Get(5)
		_ = V325
		tmp16285 := MakeNative(func(__e *ControlFlow) {
			W326 := __e.Get(1)
			_ = W326
			tmp16286 := PrimCons(V322, Nil)

			tmp16287 := PrimCons(symcons_2, tmp16286)

			tmp16288 := MakeNative(func(__e *ControlFlow) {
				W327 := __e.Get(1)
				_ = W327
				tmp16295 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)

				if True == tmp16295 {
					tmp16289 := PrimCons(V322, Nil)

					tmp16290 := PrimCons(symhead, tmp16289)

					tmp16291 := PrimCons(W327, Nil)

					tmp16292 := PrimCons(tmp16290, tmp16291)

					tmp16293 := PrimCons(V323, tmp16292)

					__e.Return(PrimCons(symlet, tmp16293))
					return

				} else {
					__e.Return(W327)
					return
				}

			}, 1)

			tmp16296 := PrimCons(V322, Nil)

			tmp16297 := PrimCons(symtail, tmp16296)

			tmp16298 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, W326, V324, V325)

			tmp16299 := PrimCons(tmp16298, Nil)

			tmp16300 := PrimCons(tmp16297, tmp16299)

			tmp16301 := PrimCons(W326, tmp16300)

			tmp16302 := PrimCons(symlet, tmp16301)

			tmp16303 := Call(__e, tmp16288, tmp16302)

			tmp16304 := PrimCons(symshen_4parse_1failure, Nil)

			tmp16305 := PrimCons(tmp16304, Nil)

			tmp16306 := PrimCons(tmp16303, tmp16305)

			tmp16307 := PrimCons(tmp16287, tmp16306)

			__e.Return(PrimCons(symif, tmp16307))
			return

		}, 1)

		tmp16308 := Call(__e, PrimFunc(symgensym), symRemainder)

		__e.TailApply(tmp16285, tmp16308)
		return

	}, 5)

	tmp16309 := Call(__e, ns2_1set, symshen_4variablecode, tmp16284)

	_ = tmp16309

	tmp16310 := MakeNative(func(__e *ControlFlow) {
		V328 := __e.Get(1)
		_ = V328
		V329 := __e.Get(2)
		_ = V329
		V330 := __e.Get(3)
		_ = V330
		V331 := __e.Get(4)
		_ = V331
		V332 := __e.Get(5)
		_ = V332
		tmp16311 := MakeNative(func(__e *ControlFlow) {
			W333 := __e.Get(1)
			_ = W333
			tmp16312 := PrimCons(V329, Nil)

			tmp16313 := PrimCons(symcons_2, tmp16312)

			tmp16314 := PrimCons(V329, Nil)

			tmp16315 := PrimCons(symtail, tmp16314)

			tmp16316 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, W333, V331, V332)

			tmp16317 := PrimCons(tmp16316, Nil)

			tmp16318 := PrimCons(tmp16315, tmp16317)

			tmp16319 := PrimCons(W333, tmp16318)

			tmp16320 := PrimCons(symlet, tmp16319)

			tmp16321 := PrimCons(symshen_4parse_1failure, Nil)

			tmp16322 := PrimCons(tmp16321, Nil)

			tmp16323 := PrimCons(tmp16320, tmp16322)

			tmp16324 := PrimCons(tmp16313, tmp16323)

			__e.Return(PrimCons(symif, tmp16324))
			return

		}, 1)

		tmp16325 := Call(__e, PrimFunc(symgensym), symRemainder)

		__e.TailApply(tmp16311, tmp16325)
		return

	}, 5)

	tmp16326 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp16310)

	_ = tmp16326

	tmp16327 := MakeNative(func(__e *ControlFlow) {
		V334 := __e.Get(1)
		_ = V334
		V335 := __e.Get(2)
		_ = V335
		V336 := __e.Get(3)
		_ = V336
		V337 := __e.Get(4)
		_ = V337
		V338 := __e.Get(5)
		_ = V338
		tmp16328 := MakeNative(func(__e *ControlFlow) {
			W339 := __e.Get(1)
			_ = W339
			tmp16329 := PrimCons(V336, Nil)

			tmp16330 := PrimCons(V335, tmp16329)

			tmp16331 := PrimCons(symshen_4hds_a_2, tmp16330)

			tmp16332 := PrimCons(V335, Nil)

			tmp16333 := PrimCons(symtail, tmp16332)

			tmp16334 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, W339, V337, V338)

			tmp16335 := PrimCons(tmp16334, Nil)

			tmp16336 := PrimCons(tmp16333, tmp16335)

			tmp16337 := PrimCons(W339, tmp16336)

			tmp16338 := PrimCons(symlet, tmp16337)

			tmp16339 := PrimCons(symshen_4parse_1failure, Nil)

			tmp16340 := PrimCons(tmp16339, Nil)

			tmp16341 := PrimCons(tmp16338, tmp16340)

			tmp16342 := PrimCons(tmp16331, tmp16341)

			__e.Return(PrimCons(symif, tmp16342))
			return

		}, 1)

		tmp16343 := Call(__e, PrimFunc(symgensym), symRemainder)

		__e.TailApply(tmp16328, tmp16343)
		return

	}, 5)

	tmp16344 := Call(__e, ns2_1set, symshen_4terminalcode, tmp16327)

	_ = tmp16344

	tmp16345 := MakeNative(func(__e *ControlFlow) {
		V347 := __e.Get(1)
		_ = V347
		V348 := __e.Get(2)
		_ = V348
		tmp16351 := PrimIsPair(V347)

		var ifres16347 Obj

		if True == tmp16351 {
			tmp16349 := PrimHead(V347)

			tmp16350 := PrimEqual(tmp16349, V348)

			var ifres16348 Obj

			if True == tmp16350 {
				ifres16348 = True

			} else {
				ifres16348 = False

			}

			ifres16347 = ifres16348

		} else {
			ifres16347 = False

		}

		if True == ifres16347 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp16352 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp16345)

	_ = tmp16352

	tmp16353 := MakeNative(func(__e *ControlFlow) {
		V349 := __e.Get(1)
		_ = V349
		V350 := __e.Get(2)
		_ = V350
		V351 := __e.Get(3)
		_ = V351
		V352 := __e.Get(4)
		_ = V352
		V353 := __e.Get(5)
		_ = V353
		tmp16354 := MakeNative(func(__e *ControlFlow) {
			W354 := __e.Get(1)
			_ = W354
			tmp16355 := MakeNative(func(__e *ControlFlow) {
				W355 := __e.Get(1)
				_ = W355
				tmp16356 := MakeNative(func(__e *ControlFlow) {
					W356 := __e.Get(1)
					_ = W356
					tmp16357 := PrimCons(V350, Nil)

					tmp16358 := PrimCons(symshen_4ccons_2, tmp16357)

					tmp16359 := PrimCons(V350, Nil)

					tmp16360 := PrimCons(symhead, tmp16359)

					tmp16361 := PrimCons(V350, Nil)

					tmp16362 := PrimCons(symtail, tmp16361)

					tmp16363 := Call(__e, PrimFunc(symshen_4decons), V351)

					tmp16364 := PrimCons(sym_5end_6, Nil)

					tmp16365 := Call(__e, PrimFunc(symappend), tmp16363, tmp16364)

					tmp16366 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W356, V352, V353)

					tmp16367 := PrimCons(tmp16366, Nil)

					tmp16368 := PrimCons(symshen_4processed, tmp16367)

					tmp16369 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W355, tmp16365, tmp16368)

					tmp16370 := PrimCons(tmp16369, Nil)

					tmp16371 := PrimCons(tmp16362, tmp16370)

					tmp16372 := PrimCons(W356, tmp16371)

					tmp16373 := PrimCons(tmp16360, tmp16372)

					tmp16374 := PrimCons(W355, tmp16373)

					tmp16375 := PrimCons(symlet, tmp16374)

					tmp16376 := PrimCons(symshen_4parse_1failure, Nil)

					tmp16377 := PrimCons(tmp16376, Nil)

					tmp16378 := PrimCons(tmp16375, tmp16377)

					tmp16379 := PrimCons(tmp16358, tmp16378)

					__e.Return(PrimCons(symif, tmp16379))
					return

				}, 1)

				tmp16380 := Call(__e, PrimFunc(symgensym), symTl)

				__e.TailApply(tmp16356, tmp16380)
				return

			}, 1)

			tmp16381 := Call(__e, PrimFunc(symgensym), symHd)

			__e.TailApply(tmp16355, tmp16381)
			return

		}, 1)

		tmp16382 := Call(__e, PrimFunc(symgensym), symRemainder)

		__e.TailApply(tmp16354, tmp16382)
		return

	}, 5)

	tmp16383 := Call(__e, ns2_1set, symshen_4conscode, tmp16353)

	_ = tmp16383

	tmp16384 := MakeNative(func(__e *ControlFlow) {
		V365 := __e.Get(1)
		_ = V365
		tmp16390 := PrimIsPair(V365)

		var ifres16386 Obj

		if True == tmp16390 {
			tmp16388 := PrimHead(V365)

			tmp16389 := PrimIsPair(tmp16388)

			var ifres16387 Obj

			if True == tmp16389 {
				ifres16387 = True

			} else {
				ifres16387 = False

			}

			ifres16386 = ifres16387

		} else {
			ifres16386 = False

		}

		if True == ifres16386 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp16391 := Call(__e, ns2_1set, symshen_4ccons_2, tmp16384)

	_ = tmp16391

	tmp16392 := MakeNative(func(__e *ControlFlow) {
		V366 := __e.Get(1)
		_ = V366
		tmp16419 := PrimIsPair(V366)

		var ifres16400 Obj

		if True == tmp16419 {
			tmp16417 := PrimHead(V366)

			tmp16418 := PrimEqual(symcons, tmp16417)

			var ifres16402 Obj

			if True == tmp16418 {
				tmp16415 := PrimTail(V366)

				tmp16416 := PrimIsPair(tmp16415)

				var ifres16404 Obj

				if True == tmp16416 {
					tmp16412 := PrimTail(V366)

					tmp16413 := PrimTail(tmp16412)

					tmp16414 := PrimIsPair(tmp16413)

					var ifres16406 Obj

					if True == tmp16414 {
						tmp16408 := PrimTail(V366)

						tmp16409 := PrimTail(tmp16408)

						tmp16410 := PrimTail(tmp16409)

						tmp16411 := PrimEqual(Nil, tmp16410)

						var ifres16407 Obj

						if True == tmp16411 {
							ifres16407 = True

						} else {
							ifres16407 = False

						}

						ifres16406 = ifres16407

					} else {
						ifres16406 = False

					}

					var ifres16405 Obj

					if True == ifres16406 {
						ifres16405 = True

					} else {
						ifres16405 = False

					}

					ifres16404 = ifres16405

				} else {
					ifres16404 = False

				}

				var ifres16403 Obj

				if True == ifres16404 {
					ifres16403 = True

				} else {
					ifres16403 = False

				}

				ifres16402 = ifres16403

			} else {
				ifres16402 = False

			}

			var ifres16401 Obj

			if True == ifres16402 {
				ifres16401 = True

			} else {
				ifres16401 = False

			}

			ifres16400 = ifres16401

		} else {
			ifres16400 = False

		}

		if True == ifres16400 {
			tmp16393 := PrimTail(V366)

			tmp16394 := PrimHead(tmp16393)

			tmp16395 := PrimTail(V366)

			tmp16396 := PrimTail(tmp16395)

			tmp16397 := PrimHead(tmp16396)

			tmp16398 := Call(__e, PrimFunc(symshen_4decons), tmp16397)

			__e.Return(PrimCons(tmp16394, tmp16398))
			return

		} else {
			__e.Return(V366)
			return
		}

	}, 1)

	tmp16420 := Call(__e, ns2_1set, symshen_4decons, tmp16392)

	_ = tmp16420

	tmp16421 := MakeNative(func(__e *ControlFlow) {
		V367 := __e.Get(1)
		_ = V367
		V368 := __e.Get(2)
		_ = V368
		tmp16422 := PrimCons(V368, Nil)

		__e.Return(PrimCons(V367, tmp16422))
		return

	}, 2)

	tmp16423 := Call(__e, ns2_1set, symshen_4comb, tmp16421)

	_ = tmp16423

	tmp16424 := MakeNative(func(__e *ControlFlow) {
		V373 := __e.Get(1)
		_ = V373
		V374 := __e.Get(2)
		_ = V374
		V375 := __e.Get(3)
		_ = V375
		tmp16446 := PrimIsPair(V375)

		var ifres16433 Obj

		if True == tmp16446 {
			tmp16444 := PrimHead(V375)

			tmp16445 := PrimEqual(symshen_4processed, tmp16444)

			var ifres16435 Obj

			if True == tmp16445 {
				tmp16442 := PrimTail(V375)

				tmp16443 := PrimIsPair(tmp16442)

				var ifres16437 Obj

				if True == tmp16443 {
					tmp16439 := PrimTail(V375)

					tmp16440 := PrimTail(tmp16439)

					tmp16441 := PrimEqual(Nil, tmp16440)

					var ifres16438 Obj

					if True == tmp16441 {
						ifres16438 = True

					} else {
						ifres16438 = False

					}

					ifres16437 = ifres16438

				} else {
					ifres16437 = False

				}

				var ifres16436 Obj

				if True == ifres16437 {
					ifres16436 = True

				} else {
					ifres16436 = False

				}

				ifres16435 = ifres16436

			} else {
				ifres16435 = False

			}

			var ifres16434 Obj

			if True == ifres16435 {
				ifres16434 = True

			} else {
				ifres16434 = False

			}

			ifres16433 = ifres16434

		} else {
			ifres16433 = False

		}

		if True == ifres16433 {
			tmp16425 := PrimTail(V375)

			__e.Return(PrimHead(tmp16425))
			return

		} else {
			tmp16426 := MakeNative(func(__e *ControlFlow) {
				W376 := __e.Get(1)
				_ = W376
				tmp16427 := MakeNative(func(__e *ControlFlow) {
					W377 := __e.Get(1)
					_ = W377
					tmp16428 := PrimCons(W377, Nil)

					tmp16429 := PrimCons(V374, tmp16428)

					__e.Return(PrimCons(symshen_4comb, tmp16429))
					return

				}, 1)

				tmp16430 := Call(__e, PrimFunc(symshen_4use_1type_1info), V373, W376)

				__e.TailApply(tmp16427, tmp16430)
				return

			}, 1)

			tmp16431 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V375)

			__e.TailApply(tmp16426, tmp16431)
			return

		}

	}, 3)

	tmp16447 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp16424)

	_ = tmp16447

	tmp16448 := MakeNative(func(__e *ControlFlow) {
		V381 := __e.Get(1)
		_ = V381
		V382 := __e.Get(2)
		_ = V382
		tmp16636 := PrimIsPair(V381)

		var ifres16457 Obj

		if True == tmp16636 {
			tmp16634 := PrimHead(V381)

			tmp16635 := PrimEqual(sym_i, tmp16634)

			var ifres16459 Obj

			if True == tmp16635 {
				tmp16632 := PrimTail(V381)

				tmp16633 := PrimIsPair(tmp16632)

				var ifres16461 Obj

				if True == tmp16633 {
					tmp16629 := PrimTail(V381)

					tmp16630 := PrimHead(tmp16629)

					tmp16631 := PrimIsPair(tmp16630)

					var ifres16463 Obj

					if True == tmp16631 {
						tmp16625 := PrimTail(V381)

						tmp16626 := PrimHead(tmp16625)

						tmp16627 := PrimHead(tmp16626)

						tmp16628 := PrimEqual(symlist, tmp16627)

						var ifres16465 Obj

						if True == tmp16628 {
							tmp16621 := PrimTail(V381)

							tmp16622 := PrimHead(tmp16621)

							tmp16623 := PrimTail(tmp16622)

							tmp16624 := PrimIsPair(tmp16623)

							var ifres16467 Obj

							if True == tmp16624 {
								tmp16616 := PrimTail(V381)

								tmp16617 := PrimHead(tmp16616)

								tmp16618 := PrimTail(tmp16617)

								tmp16619 := PrimTail(tmp16618)

								tmp16620 := PrimEqual(Nil, tmp16619)

								var ifres16469 Obj

								if True == tmp16620 {
									tmp16613 := PrimTail(V381)

									tmp16614 := PrimTail(tmp16613)

									tmp16615 := PrimIsPair(tmp16614)

									var ifres16471 Obj

									if True == tmp16615 {
										tmp16609 := PrimTail(V381)

										tmp16610 := PrimTail(tmp16609)

										tmp16611 := PrimHead(tmp16610)

										tmp16612 := PrimEqual(sym_1_1_6, tmp16611)

										var ifres16473 Obj

										if True == tmp16612 {
											tmp16605 := PrimTail(V381)

											tmp16606 := PrimTail(tmp16605)

											tmp16607 := PrimTail(tmp16606)

											tmp16608 := PrimIsPair(tmp16607)

											var ifres16475 Obj

											if True == tmp16608 {
												tmp16600 := PrimTail(V381)

												tmp16601 := PrimTail(tmp16600)

												tmp16602 := PrimTail(tmp16601)

												tmp16603 := PrimHead(tmp16602)

												tmp16604 := PrimIsPair(tmp16603)

												var ifres16477 Obj

												if True == tmp16604 {
													tmp16594 := PrimTail(V381)

													tmp16595 := PrimTail(tmp16594)

													tmp16596 := PrimTail(tmp16595)

													tmp16597 := PrimHead(tmp16596)

													tmp16598 := PrimHead(tmp16597)

													tmp16599 := PrimEqual(symstr, tmp16598)

													var ifres16479 Obj

													if True == tmp16599 {
														tmp16588 := PrimTail(V381)

														tmp16589 := PrimTail(tmp16588)

														tmp16590 := PrimTail(tmp16589)

														tmp16591 := PrimHead(tmp16590)

														tmp16592 := PrimTail(tmp16591)

														tmp16593 := PrimIsPair(tmp16592)

														var ifres16481 Obj

														if True == tmp16593 {
															tmp16581 := PrimTail(V381)

															tmp16582 := PrimTail(tmp16581)

															tmp16583 := PrimTail(tmp16582)

															tmp16584 := PrimHead(tmp16583)

															tmp16585 := PrimTail(tmp16584)

															tmp16586 := PrimHead(tmp16585)

															tmp16587 := PrimIsPair(tmp16586)

															var ifres16483 Obj

															if True == tmp16587 {
																tmp16573 := PrimTail(V381)

																tmp16574 := PrimTail(tmp16573)

																tmp16575 := PrimTail(tmp16574)

																tmp16576 := PrimHead(tmp16575)

																tmp16577 := PrimTail(tmp16576)

																tmp16578 := PrimHead(tmp16577)

																tmp16579 := PrimHead(tmp16578)

																tmp16580 := PrimEqual(symlist, tmp16579)

																var ifres16485 Obj

																if True == tmp16580 {
																	tmp16565 := PrimTail(V381)

																	tmp16566 := PrimTail(tmp16565)

																	tmp16567 := PrimTail(tmp16566)

																	tmp16568 := PrimHead(tmp16567)

																	tmp16569 := PrimTail(tmp16568)

																	tmp16570 := PrimHead(tmp16569)

																	tmp16571 := PrimTail(tmp16570)

																	tmp16572 := PrimIsPair(tmp16571)

																	var ifres16487 Obj

																	if True == tmp16572 {
																		tmp16556 := PrimTail(V381)

																		tmp16557 := PrimTail(tmp16556)

																		tmp16558 := PrimTail(tmp16557)

																		tmp16559 := PrimHead(tmp16558)

																		tmp16560 := PrimTail(tmp16559)

																		tmp16561 := PrimHead(tmp16560)

																		tmp16562 := PrimTail(tmp16561)

																		tmp16563 := PrimTail(tmp16562)

																		tmp16564 := PrimEqual(Nil, tmp16563)

																		var ifres16489 Obj

																		if True == tmp16564 {
																			tmp16549 := PrimTail(V381)

																			tmp16550 := PrimTail(tmp16549)

																			tmp16551 := PrimTail(tmp16550)

																			tmp16552 := PrimHead(tmp16551)

																			tmp16553 := PrimTail(tmp16552)

																			tmp16554 := PrimTail(tmp16553)

																			tmp16555 := PrimIsPair(tmp16554)

																			var ifres16491 Obj

																			if True == tmp16555 {
																				tmp16541 := PrimTail(V381)

																				tmp16542 := PrimTail(tmp16541)

																				tmp16543 := PrimTail(tmp16542)

																				tmp16544 := PrimHead(tmp16543)

																				tmp16545 := PrimTail(tmp16544)

																				tmp16546 := PrimTail(tmp16545)

																				tmp16547 := PrimTail(tmp16546)

																				tmp16548 := PrimEqual(Nil, tmp16547)

																				var ifres16493 Obj

																				if True == tmp16548 {
																					tmp16536 := PrimTail(V381)

																					tmp16537 := PrimTail(tmp16536)

																					tmp16538 := PrimTail(tmp16537)

																					tmp16539 := PrimTail(tmp16538)

																					tmp16540 := PrimIsPair(tmp16539)

																					var ifres16495 Obj

																					if True == tmp16540 {
																						tmp16530 := PrimTail(V381)

																						tmp16531 := PrimTail(tmp16530)

																						tmp16532 := PrimTail(tmp16531)

																						tmp16533 := PrimTail(tmp16532)

																						tmp16534 := PrimHead(tmp16533)

																						tmp16535 := PrimEqual(sym_j, tmp16534)

																						var ifres16497 Obj

																						if True == tmp16535 {
																							tmp16524 := PrimTail(V381)

																							tmp16525 := PrimTail(tmp16524)

																							tmp16526 := PrimTail(tmp16525)

																							tmp16527 := PrimTail(tmp16526)

																							tmp16528 := PrimTail(tmp16527)

																							tmp16529 := PrimEqual(Nil, tmp16528)

																							var ifres16499 Obj

																							if True == tmp16529 {
																								tmp16511 := PrimTail(V381)

																								tmp16512 := PrimHead(tmp16511)

																								tmp16513 := PrimTail(tmp16512)

																								tmp16514 := PrimHead(tmp16513)

																								tmp16515 := PrimTail(V381)

																								tmp16516 := PrimTail(tmp16515)

																								tmp16517 := PrimTail(tmp16516)

																								tmp16518 := PrimHead(tmp16517)

																								tmp16519 := PrimTail(tmp16518)

																								tmp16520 := PrimHead(tmp16519)

																								tmp16521 := PrimTail(tmp16520)

																								tmp16522 := PrimHead(tmp16521)

																								tmp16523 := PrimEqual(tmp16514, tmp16522)

																								var ifres16501 Obj

																								if True == tmp16523 {
																									tmp16503 := PrimTail(V381)

																									tmp16504 := PrimTail(tmp16503)

																									tmp16505 := PrimTail(tmp16504)

																									tmp16506 := PrimHead(tmp16505)

																									tmp16507 := PrimTail(tmp16506)

																									tmp16508 := PrimTail(tmp16507)

																									tmp16509 := PrimHead(tmp16508)

																									tmp16510 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp16509)

																									var ifres16502 Obj

																									if True == tmp16510 {
																										ifres16502 = True

																									} else {
																										ifres16502 = False

																									}

																									ifres16501 = ifres16502

																								} else {
																									ifres16501 = False

																								}

																								var ifres16500 Obj

																								if True == ifres16501 {
																									ifres16500 = True

																								} else {
																									ifres16500 = False

																								}

																								ifres16499 = ifres16500

																							} else {
																								ifres16499 = False

																							}

																							var ifres16498 Obj

																							if True == ifres16499 {
																								ifres16498 = True

																							} else {
																								ifres16498 = False

																							}

																							ifres16497 = ifres16498

																						} else {
																							ifres16497 = False

																						}

																						var ifres16496 Obj

																						if True == ifres16497 {
																							ifres16496 = True

																						} else {
																							ifres16496 = False

																						}

																						ifres16495 = ifres16496

																					} else {
																						ifres16495 = False

																					}

																					var ifres16494 Obj

																					if True == ifres16495 {
																						ifres16494 = True

																					} else {
																						ifres16494 = False

																					}

																					ifres16493 = ifres16494

																				} else {
																					ifres16493 = False

																				}

																				var ifres16492 Obj

																				if True == ifres16493 {
																					ifres16492 = True

																				} else {
																					ifres16492 = False

																				}

																				ifres16491 = ifres16492

																			} else {
																				ifres16491 = False

																			}

																			var ifres16490 Obj

																			if True == ifres16491 {
																				ifres16490 = True

																			} else {
																				ifres16490 = False

																			}

																			ifres16489 = ifres16490

																		} else {
																			ifres16489 = False

																		}

																		var ifres16488 Obj

																		if True == ifres16489 {
																			ifres16488 = True

																		} else {
																			ifres16488 = False

																		}

																		ifres16487 = ifres16488

																	} else {
																		ifres16487 = False

																	}

																	var ifres16486 Obj

																	if True == ifres16487 {
																		ifres16486 = True

																	} else {
																		ifres16486 = False

																	}

																	ifres16485 = ifres16486

																} else {
																	ifres16485 = False

																}

																var ifres16484 Obj

																if True == ifres16485 {
																	ifres16484 = True

																} else {
																	ifres16484 = False

																}

																ifres16483 = ifres16484

															} else {
																ifres16483 = False

															}

															var ifres16482 Obj

															if True == ifres16483 {
																ifres16482 = True

															} else {
																ifres16482 = False

															}

															ifres16481 = ifres16482

														} else {
															ifres16481 = False

														}

														var ifres16480 Obj

														if True == ifres16481 {
															ifres16480 = True

														} else {
															ifres16480 = False

														}

														ifres16479 = ifres16480

													} else {
														ifres16479 = False

													}

													var ifres16478 Obj

													if True == ifres16479 {
														ifres16478 = True

													} else {
														ifres16478 = False

													}

													ifres16477 = ifres16478

												} else {
													ifres16477 = False

												}

												var ifres16476 Obj

												if True == ifres16477 {
													ifres16476 = True

												} else {
													ifres16476 = False

												}

												ifres16475 = ifres16476

											} else {
												ifres16475 = False

											}

											var ifres16474 Obj

											if True == ifres16475 {
												ifres16474 = True

											} else {
												ifres16474 = False

											}

											ifres16473 = ifres16474

										} else {
											ifres16473 = False

										}

										var ifres16472 Obj

										if True == ifres16473 {
											ifres16472 = True

										} else {
											ifres16472 = False

										}

										ifres16471 = ifres16472

									} else {
										ifres16471 = False

									}

									var ifres16470 Obj

									if True == ifres16471 {
										ifres16470 = True

									} else {
										ifres16470 = False

									}

									ifres16469 = ifres16470

								} else {
									ifres16469 = False

								}

								var ifres16468 Obj

								if True == ifres16469 {
									ifres16468 = True

								} else {
									ifres16468 = False

								}

								ifres16467 = ifres16468

							} else {
								ifres16467 = False

							}

							var ifres16466 Obj

							if True == ifres16467 {
								ifres16466 = True

							} else {
								ifres16466 = False

							}

							ifres16465 = ifres16466

						} else {
							ifres16465 = False

						}

						var ifres16464 Obj

						if True == ifres16465 {
							ifres16464 = True

						} else {
							ifres16464 = False

						}

						ifres16463 = ifres16464

					} else {
						ifres16463 = False

					}

					var ifres16462 Obj

					if True == ifres16463 {
						ifres16462 = True

					} else {
						ifres16462 = False

					}

					ifres16461 = ifres16462

				} else {
					ifres16461 = False

				}

				var ifres16460 Obj

				if True == ifres16461 {
					ifres16460 = True

				} else {
					ifres16460 = False

				}

				ifres16459 = ifres16460

			} else {
				ifres16459 = False

			}

			var ifres16458 Obj

			if True == ifres16459 {
				ifres16458 = True

			} else {
				ifres16458 = False

			}

			ifres16457 = ifres16458

		} else {
			ifres16457 = False

		}

		if True == ifres16457 {
			tmp16449 := PrimTail(V381)

			tmp16450 := PrimTail(tmp16449)

			tmp16451 := PrimTail(tmp16450)

			tmp16452 := PrimHead(tmp16451)

			tmp16453 := PrimTail(tmp16452)

			tmp16454 := PrimTail(tmp16453)

			tmp16455 := PrimCons(V382, tmp16454)

			__e.Return(PrimCons(symtype, tmp16455))
			return

		} else {
			__e.Return(V382)
			return
		}

	}, 2)

	tmp16637 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp16448)

	_ = tmp16637

	tmp16638 := MakeNative(func(__e *ControlFlow) {
		V385 := __e.Get(1)
		_ = V385
		tmp16648 := PrimIsVariable(V385)

		if True == tmp16648 {
			__e.Return(False)
			return
		} else {
			tmp16646 := PrimIsPair(V385)

			if True == tmp16646 {
				tmp16643 := PrimHead(V385)

				tmp16644 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp16643)

				if True == tmp16644 {
					tmp16640 := PrimTail(V385)

					tmp16641 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp16640)

					if True == tmp16641 {
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
				__e.Return(True)
				return
			}

		}

	}, 1)

	tmp16649 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp16638)

	_ = tmp16649

	tmp16650 := MakeNative(func(__e *ControlFlow) {
		V386 := __e.Get(1)
		_ = V386
		tmp16676 := PrimIsPair(V386)

		var ifres16658 Obj

		if True == tmp16676 {
			tmp16674 := PrimHead(V386)

			tmp16675 := PrimEqual(symprotect, tmp16674)

			var ifres16660 Obj

			if True == tmp16675 {
				tmp16672 := PrimTail(V386)

				tmp16673 := PrimIsPair(tmp16672)

				var ifres16662 Obj

				if True == tmp16673 {
					tmp16669 := PrimTail(V386)

					tmp16670 := PrimTail(tmp16669)

					tmp16671 := PrimEqual(Nil, tmp16670)

					var ifres16664 Obj

					if True == tmp16671 {
						tmp16666 := PrimTail(V386)

						tmp16667 := PrimHead(tmp16666)

						tmp16668 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16667)

						var ifres16665 Obj

						if True == tmp16668 {
							ifres16665 = True

						} else {
							ifres16665 = False

						}

						ifres16664 = ifres16665

					} else {
						ifres16664 = False

					}

					var ifres16663 Obj

					if True == ifres16664 {
						ifres16663 = True

					} else {
						ifres16663 = False

					}

					ifres16662 = ifres16663

				} else {
					ifres16662 = False

				}

				var ifres16661 Obj

				if True == ifres16662 {
					ifres16661 = True

				} else {
					ifres16661 = False

				}

				ifres16660 = ifres16661

			} else {
				ifres16660 = False

			}

			var ifres16659 Obj

			if True == ifres16660 {
				ifres16659 = True

			} else {
				ifres16659 = False

			}

			ifres16658 = ifres16659

		} else {
			ifres16658 = False

		}

		if True == ifres16658 {
			tmp16651 := PrimTail(V386)

			__e.Return(PrimHead(tmp16651))
			return

		} else {
			tmp16656 := PrimIsPair(V386)

			if True == tmp16656 {
				tmp16652 := MakeNative(func(__e *ControlFlow) {
					Z387 := __e.Get(1)
					_ = Z387
					__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z387)
					return
				}, 1)

				__e.TailApply(PrimFunc(symmap), tmp16652, V386)
				return

			} else {
				tmp16654 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V386)

				if True == tmp16654 {
					__e.TailApply(PrimFunc(symconcat), symAction, V386)
					return
				} else {
					__e.Return(V386)
					return
				}

			}

		}

	}, 1)

	tmp16677 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp16650)

	_ = tmp16677

	tmp16678 := MakeNative(func(__e *ControlFlow) {
		V390 := __e.Get(1)
		_ = V390
		tmp16691 := PrimIsPair(V390)

		var ifres16682 Obj

		if True == tmp16691 {
			tmp16689 := PrimTail(V390)

			tmp16690 := PrimIsPair(tmp16689)

			var ifres16684 Obj

			if True == tmp16690 {
				tmp16686 := PrimTail(V390)

				tmp16687 := PrimTail(tmp16686)

				tmp16688 := PrimEqual(Nil, tmp16687)

				var ifres16685 Obj

				if True == tmp16688 {
					ifres16685 = True

				} else {
					ifres16685 = False

				}

				ifres16684 = ifres16685

			} else {
				ifres16684 = False

			}

			var ifres16683 Obj

			if True == ifres16684 {
				ifres16683 = True

			} else {
				ifres16683 = False

			}

			ifres16682 = ifres16683

		} else {
			ifres16682 = False

		}

		if True == ifres16682 {
			tmp16679 := PrimTail(V390)

			__e.Return(PrimHead(tmp16679))
			return

		} else {
			tmp16680 := PrimTail(V390)

			__e.Return(PrimHead(tmp16680))
			return

		}

	}, 1)

	tmp16692 := Call(__e, ns2_1set, symshen_4_5_1out, tmp16678)

	_ = tmp16692

	tmp16693 := MakeNative(func(__e *ControlFlow) {
		V391 := __e.Get(1)
		_ = V391
		__e.Return(PrimHead(V391))
		return
	}, 1)

	tmp16694 := Call(__e, ns2_1set, symshen_4in_1_6, tmp16693)

	_ = tmp16694

	tmp16695 := MakeNative(func(__e *ControlFlow) {
		V392 := __e.Get(1)
		_ = V392
		tmp16696 := PrimCons(V392, Nil)

		__e.Return(PrimCons(Nil, tmp16696))
		return

	}, 1)

	tmp16697 := Call(__e, ns2_1set, sym_5_b_6, tmp16695)

	_ = tmp16697

	tmp16698 := MakeNative(func(__e *ControlFlow) {
		V393 := __e.Get(1)
		_ = V393
		tmp16699 := PrimCons(Nil, Nil)

		__e.Return(PrimCons(V393, tmp16699))
		return

	}, 1)

	tmp16700 := Call(__e, ns2_1set, sym_5e_6, tmp16698)

	_ = tmp16700

	tmp16701 := MakeNative(func(__e *ControlFlow) {
		V396 := __e.Get(1)
		_ = V396
		tmp16704 := PrimEqual(Nil, V396)

		if True == tmp16704 {
			tmp16702 := PrimCons(Nil, Nil)

			__e.Return(PrimCons(Nil, tmp16702))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4parse_1failure))
			return
		}

	}, 1)

	__e.TailApply(ns2_1set, sym_5end_6, tmp16701)
	return

}, 0)
