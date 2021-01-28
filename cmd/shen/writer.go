package main

import . "github.com/tiancaiamao/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp23536 := MakeNative(func(__e *ControlFlow) {
		V4594 := __e.Get(1)
		_ = V4594
		V4595 := __e.Get(2)
		_ = V4595
		tmp23537 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4prh), V4594, V4595, MakeNumber(0))
			return
		}, 0)

		tmp23538 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V4594)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp23537, tmp23538)
		return

	}, 2)

	tmp23539 := Call(__e, PrimNS1Value(symns2_1set), sympr, tmp23536)

	_ = tmp23539

	tmp23540 := MakeNative(func(__e *ControlFlow) {
		V4599 := __e.Get(1)
		_ = V4599
		V4600 := __e.Get(2)
		_ = V4600
		V4601 := __e.Get(3)
		_ = V4601
		tmp23541 := Call(__e, PrimNS2Value(symshen_4write_1char_1and_1inc), V4599, V4600, V4601)

		__e.TailApply(PrimNS2Value(symshen_4prh), V4599, V4600, tmp23541)
		return

	}, 3)

	tmp23542 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prh, tmp23540)

	_ = tmp23542

	tmp23543 := MakeNative(func(__e *ControlFlow) {
		V4605 := __e.Get(1)
		_ = V4605
		V4606 := __e.Get(2)
		_ = V4606
		V4607 := __e.Get(3)
		_ = V4607
		tmp23544 := PrimPos(V4605, V4607)

		tmp23545 := PrimStringToNumber(tmp23544)

		tmp23546 := PrimWriteByte(tmp23545, V4606)

		_ = tmp23546

		__e.Return(PrimNumberAdd(V4607, MakeNumber(1)))
		return

	}, 3)

	tmp23547 := Call(__e, PrimNS1Value(symns2_1set), symshen_4write_1char_1and_1inc, tmp23543)

	_ = tmp23547

	tmp23548 := MakeNative(func(__e *ControlFlow) {
		V4609 := __e.Get(1)
		_ = V4609
		tmp23549 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			tmp23550 := MakeNative(func(__e *ControlFlow) {
				Print := __e.Get(1)
				_ = Print
				__e.Return(V4609)
				return
			}, 1)

			tmp23551 := Call(__e, PrimNS2Value(symstoutput))

			tmp23552 := Call(__e, PrimNS2Value(symshen_4prhush), String, tmp23551)

			__e.TailApply(tmp23550, tmp23552)
			return

		}, 1)

		tmp23553 := Call(__e, PrimNS2Value(symshen_4insert), V4609, MakeString("~S"))

		__e.TailApply(tmp23549, tmp23553)
		return

	}, 1)

	tmp23554 := Call(__e, PrimNS1Value(symns2_1set), symprint, tmp23548)

	_ = tmp23554

	tmp23555 := MakeNative(func(__e *ControlFlow) {
		V4612 := __e.Get(1)
		_ = V4612
		V4613 := __e.Get(2)
		_ = V4613
		tmp23557 := PrimNS3Value(sym_dhush_d)

		if True == tmp23557 {
			__e.Return(V4612)
			return
		} else {
			__e.TailApply(PrimNS2Value(sympr), V4612, V4613)
			return
		}

	}, 2)

	tmp23558 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prhush, tmp23555)

	_ = tmp23558

	tmp23559 := MakeNative(func(__e *ControlFlow) {
		V4616 := __e.Get(1)
		_ = V4616
		V4617 := __e.Get(2)
		_ = V4617
		tmp23564 := PrimIsString(V4616)

		if True == tmp23564 {
			tmp23563 := Call(__e, PrimNS2Value(symshen_4proc_1nl), V4616)

			__e.TailApply(PrimNS2Value(symshen_4mkstr_1l), tmp23563, V4617)
			return

		} else {
			tmp23561 := PrimCons(V4616, Nil)

			tmp23562 := PrimCons(symshen_4proc_1nl, tmp23561)

			__e.TailApply(PrimNS2Value(symshen_4mkstr_1r), tmp23562, V4617)
			return

		}

	}, 2)

	tmp23565 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr, tmp23559)

	_ = tmp23565

	tmp23566 := MakeNative(func(__e *ControlFlow) {
		V4620 := __e.Get(1)
		_ = V4620
		V4621 := __e.Get(2)
		_ = V4621
		tmp23573 := PrimEqual(Nil, V4621)

		if True == tmp23573 {
			__e.Return(V4620)
			return
		} else {
			tmp23572 := PrimIsPair(V4621)

			if True == tmp23572 {
				tmp23569 := PrimHead(V4621)

				tmp23570 := Call(__e, PrimNS2Value(symshen_4insert_1l), tmp23569, V4620)

				tmp23571 := PrimTail(V4621)

				__e.TailApply(PrimNS2Value(symshen_4mkstr_1l), tmp23570, tmp23571)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4mkstr_1l)
				return
			}

		}

	}, 2)

	tmp23574 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr_1l, tmp23566)

	_ = tmp23574

	tmp23575 := MakeNative(func(__e *ControlFlow) {
		V4626 := __e.Get(1)
		_ = V4626
		V4627 := __e.Get(2)
		_ = V4627
		tmp23713 := PrimEqual(MakeString(""), V4627)

		if True == tmp23713 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp23712 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4627)

			var ifres23699 Obj

			if True == tmp23712 {
				tmp23710 := PrimPos(V4627, MakeNumber(0))

				tmp23711 := PrimEqual(MakeString("~"), tmp23710)

				var ifres23701 Obj

				if True == tmp23711 {
					tmp23708 := PrimTailString(V4627)

					tmp23709 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23708)

					var ifres23703 Obj

					if True == tmp23709 {
						tmp23705 := PrimTailString(V4627)

						tmp23706 := PrimPos(tmp23705, MakeNumber(0))

						tmp23707 := PrimEqual(MakeString("A"), tmp23706)

						var ifres23704 Obj

						if True == tmp23707 {
							ifres23704 = True

						} else {
							ifres23704 = False

						}

						ifres23703 = ifres23704

					} else {
						ifres23703 = False

					}

					var ifres23702 Obj

					if True == ifres23703 {
						ifres23702 = True

					} else {
						ifres23702 = False

					}

					ifres23701 = ifres23702

				} else {
					ifres23701 = False

				}

				var ifres23700 Obj

				if True == ifres23701 {
					ifres23700 = True

				} else {
					ifres23700 = False

				}

				ifres23699 = ifres23700

			} else {
				ifres23699 = False

			}

			if True == ifres23699 {
				tmp23694 := PrimTailString(V4627)

				tmp23695 := PrimTailString(tmp23694)

				tmp23696 := PrimCons(symshen_4a, Nil)

				tmp23697 := PrimCons(tmp23695, tmp23696)

				tmp23698 := PrimCons(V4626, tmp23697)

				__e.Return(PrimCons(symshen_4app, tmp23698))
				return

			} else {
				tmp23693 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4627)

				var ifres23680 Obj

				if True == tmp23693 {
					tmp23691 := PrimPos(V4627, MakeNumber(0))

					tmp23692 := PrimEqual(MakeString("~"), tmp23691)

					var ifres23682 Obj

					if True == tmp23692 {
						tmp23689 := PrimTailString(V4627)

						tmp23690 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23689)

						var ifres23684 Obj

						if True == tmp23690 {
							tmp23686 := PrimTailString(V4627)

							tmp23687 := PrimPos(tmp23686, MakeNumber(0))

							tmp23688 := PrimEqual(MakeString("R"), tmp23687)

							var ifres23685 Obj

							if True == tmp23688 {
								ifres23685 = True

							} else {
								ifres23685 = False

							}

							ifres23684 = ifres23685

						} else {
							ifres23684 = False

						}

						var ifres23683 Obj

						if True == ifres23684 {
							ifres23683 = True

						} else {
							ifres23683 = False

						}

						ifres23682 = ifres23683

					} else {
						ifres23682 = False

					}

					var ifres23681 Obj

					if True == ifres23682 {
						ifres23681 = True

					} else {
						ifres23681 = False

					}

					ifres23680 = ifres23681

				} else {
					ifres23680 = False

				}

				if True == ifres23680 {
					tmp23675 := PrimTailString(V4627)

					tmp23676 := PrimTailString(tmp23675)

					tmp23677 := PrimCons(symshen_4r, Nil)

					tmp23678 := PrimCons(tmp23676, tmp23677)

					tmp23679 := PrimCons(V4626, tmp23678)

					__e.Return(PrimCons(symshen_4app, tmp23679))
					return

				} else {
					tmp23674 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4627)

					var ifres23661 Obj

					if True == tmp23674 {
						tmp23672 := PrimPos(V4627, MakeNumber(0))

						tmp23673 := PrimEqual(MakeString("~"), tmp23672)

						var ifres23663 Obj

						if True == tmp23673 {
							tmp23670 := PrimTailString(V4627)

							tmp23671 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23670)

							var ifres23665 Obj

							if True == tmp23671 {
								tmp23667 := PrimTailString(V4627)

								tmp23668 := PrimPos(tmp23667, MakeNumber(0))

								tmp23669 := PrimEqual(MakeString("S"), tmp23668)

								var ifres23666 Obj

								if True == tmp23669 {
									ifres23666 = True

								} else {
									ifres23666 = False

								}

								ifres23665 = ifres23666

							} else {
								ifres23665 = False

							}

							var ifres23664 Obj

							if True == ifres23665 {
								ifres23664 = True

							} else {
								ifres23664 = False

							}

							ifres23663 = ifres23664

						} else {
							ifres23663 = False

						}

						var ifres23662 Obj

						if True == ifres23663 {
							ifres23662 = True

						} else {
							ifres23662 = False

						}

						ifres23661 = ifres23662

					} else {
						ifres23661 = False

					}

					if True == ifres23661 {
						tmp23656 := PrimTailString(V4627)

						tmp23657 := PrimTailString(tmp23656)

						tmp23658 := PrimCons(symshen_4s, Nil)

						tmp23659 := PrimCons(tmp23657, tmp23658)

						tmp23660 := PrimCons(V4626, tmp23659)

						__e.Return(PrimCons(symshen_4app, tmp23660))
						return

					} else {
						tmp23655 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4627)

						if True == tmp23655 {
							tmp23649 := PrimPos(V4627, MakeNumber(0))

							tmp23650 := PrimTailString(V4627)

							tmp23651 := Call(__e, PrimNS2Value(symshen_4insert_1l), V4626, tmp23650)

							tmp23652 := PrimCons(tmp23651, Nil)

							tmp23653 := PrimCons(tmp23649, tmp23652)

							tmp23654 := PrimCons(symcn, tmp23653)

							__e.TailApply(PrimNS2Value(symshen_4factor_1cn), tmp23654)
							return

						} else {
							tmp23648 := PrimIsPair(V4627)

							var ifres23629 Obj

							if True == tmp23648 {
								tmp23646 := PrimHead(V4627)

								tmp23647 := PrimEqual(symcn, tmp23646)

								var ifres23631 Obj

								if True == tmp23647 {
									tmp23644 := PrimTail(V4627)

									tmp23645 := PrimIsPair(tmp23644)

									var ifres23633 Obj

									if True == tmp23645 {
										tmp23641 := PrimTail(V4627)

										tmp23642 := PrimTail(tmp23641)

										tmp23643 := PrimIsPair(tmp23642)

										var ifres23635 Obj

										if True == tmp23643 {
											tmp23637 := PrimTail(V4627)

											tmp23638 := PrimTail(tmp23637)

											tmp23639 := PrimTail(tmp23638)

											tmp23640 := PrimEqual(Nil, tmp23639)

											var ifres23636 Obj

											if True == tmp23640 {
												ifres23636 = True

											} else {
												ifres23636 = False

											}

											ifres23635 = ifres23636

										} else {
											ifres23635 = False

										}

										var ifres23634 Obj

										if True == ifres23635 {
											ifres23634 = True

										} else {
											ifres23634 = False

										}

										ifres23633 = ifres23634

									} else {
										ifres23633 = False

									}

									var ifres23632 Obj

									if True == ifres23633 {
										ifres23632 = True

									} else {
										ifres23632 = False

									}

									ifres23631 = ifres23632

								} else {
									ifres23631 = False

								}

								var ifres23630 Obj

								if True == ifres23631 {
									ifres23630 = True

								} else {
									ifres23630 = False

								}

								ifres23629 = ifres23630

							} else {
								ifres23629 = False

							}

							if True == ifres23629 {
								tmp23621 := PrimTail(V4627)

								tmp23622 := PrimHead(tmp23621)

								tmp23623 := PrimTail(V4627)

								tmp23624 := PrimTail(tmp23623)

								tmp23625 := PrimHead(tmp23624)

								tmp23626 := Call(__e, PrimNS2Value(symshen_4insert_1l), V4626, tmp23625)

								tmp23627 := PrimCons(tmp23626, Nil)

								tmp23628 := PrimCons(tmp23622, tmp23627)

								__e.Return(PrimCons(symcn, tmp23628))
								return

							} else {
								tmp23620 := PrimIsPair(V4627)

								var ifres23594 Obj

								if True == tmp23620 {
									tmp23618 := PrimHead(V4627)

									tmp23619 := PrimEqual(symshen_4app, tmp23618)

									var ifres23596 Obj

									if True == tmp23619 {
										tmp23616 := PrimTail(V4627)

										tmp23617 := PrimIsPair(tmp23616)

										var ifres23598 Obj

										if True == tmp23617 {
											tmp23613 := PrimTail(V4627)

											tmp23614 := PrimTail(tmp23613)

											tmp23615 := PrimIsPair(tmp23614)

											var ifres23600 Obj

											if True == tmp23615 {
												tmp23609 := PrimTail(V4627)

												tmp23610 := PrimTail(tmp23609)

												tmp23611 := PrimTail(tmp23610)

												tmp23612 := PrimIsPair(tmp23611)

												var ifres23602 Obj

												if True == tmp23612 {
													tmp23604 := PrimTail(V4627)

													tmp23605 := PrimTail(tmp23604)

													tmp23606 := PrimTail(tmp23605)

													tmp23607 := PrimTail(tmp23606)

													tmp23608 := PrimEqual(Nil, tmp23607)

													var ifres23603 Obj

													if True == tmp23608 {
														ifres23603 = True

													} else {
														ifres23603 = False

													}

													ifres23602 = ifres23603

												} else {
													ifres23602 = False

												}

												var ifres23601 Obj

												if True == ifres23602 {
													ifres23601 = True

												} else {
													ifres23601 = False

												}

												ifres23600 = ifres23601

											} else {
												ifres23600 = False

											}

											var ifres23599 Obj

											if True == ifres23600 {
												ifres23599 = True

											} else {
												ifres23599 = False

											}

											ifres23598 = ifres23599

										} else {
											ifres23598 = False

										}

										var ifres23597 Obj

										if True == ifres23598 {
											ifres23597 = True

										} else {
											ifres23597 = False

										}

										ifres23596 = ifres23597

									} else {
										ifres23596 = False

									}

									var ifres23595 Obj

									if True == ifres23596 {
										ifres23595 = True

									} else {
										ifres23595 = False

									}

									ifres23594 = ifres23595

								} else {
									ifres23594 = False

								}

								if True == ifres23594 {
									tmp23583 := PrimTail(V4627)

									tmp23584 := PrimHead(tmp23583)

									tmp23585 := PrimTail(V4627)

									tmp23586 := PrimTail(tmp23585)

									tmp23587 := PrimHead(tmp23586)

									tmp23588 := Call(__e, PrimNS2Value(symshen_4insert_1l), V4626, tmp23587)

									tmp23589 := PrimTail(V4627)

									tmp23590 := PrimTail(tmp23589)

									tmp23591 := PrimTail(tmp23590)

									tmp23592 := PrimCons(tmp23588, tmp23591)

									tmp23593 := PrimCons(tmp23584, tmp23592)

									__e.Return(PrimCons(symshen_4app, tmp23593))
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4insert_1l)
									return
								}

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp23714 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1l, tmp23575)

	_ = tmp23714

	tmp23715 := MakeNative(func(__e *ControlFlow) {
		V4629 := __e.Get(1)
		_ = V4629
		tmp23800 := PrimIsPair(V4629)

		var ifres23731 Obj

		if True == tmp23800 {
			tmp23798 := PrimHead(V4629)

			tmp23799 := PrimEqual(symcn, tmp23798)

			var ifres23733 Obj

			if True == tmp23799 {
				tmp23796 := PrimTail(V4629)

				tmp23797 := PrimIsPair(tmp23796)

				var ifres23735 Obj

				if True == tmp23797 {
					tmp23793 := PrimTail(V4629)

					tmp23794 := PrimTail(tmp23793)

					tmp23795 := PrimIsPair(tmp23794)

					var ifres23737 Obj

					if True == tmp23795 {
						tmp23789 := PrimTail(V4629)

						tmp23790 := PrimTail(tmp23789)

						tmp23791 := PrimHead(tmp23790)

						tmp23792 := PrimIsPair(tmp23791)

						var ifres23739 Obj

						if True == tmp23792 {
							tmp23784 := PrimTail(V4629)

							tmp23785 := PrimTail(tmp23784)

							tmp23786 := PrimHead(tmp23785)

							tmp23787 := PrimHead(tmp23786)

							tmp23788 := PrimEqual(symcn, tmp23787)

							var ifres23741 Obj

							if True == tmp23788 {
								tmp23779 := PrimTail(V4629)

								tmp23780 := PrimTail(tmp23779)

								tmp23781 := PrimHead(tmp23780)

								tmp23782 := PrimTail(tmp23781)

								tmp23783 := PrimIsPair(tmp23782)

								var ifres23743 Obj

								if True == tmp23783 {
									tmp23773 := PrimTail(V4629)

									tmp23774 := PrimTail(tmp23773)

									tmp23775 := PrimHead(tmp23774)

									tmp23776 := PrimTail(tmp23775)

									tmp23777 := PrimTail(tmp23776)

									tmp23778 := PrimIsPair(tmp23777)

									var ifres23745 Obj

									if True == tmp23778 {
										tmp23766 := PrimTail(V4629)

										tmp23767 := PrimTail(tmp23766)

										tmp23768 := PrimHead(tmp23767)

										tmp23769 := PrimTail(tmp23768)

										tmp23770 := PrimTail(tmp23769)

										tmp23771 := PrimTail(tmp23770)

										tmp23772 := PrimEqual(Nil, tmp23771)

										var ifres23747 Obj

										if True == tmp23772 {
											tmp23762 := PrimTail(V4629)

											tmp23763 := PrimTail(tmp23762)

											tmp23764 := PrimTail(tmp23763)

											tmp23765 := PrimEqual(Nil, tmp23764)

											var ifres23749 Obj

											if True == tmp23765 {
												tmp23759 := PrimTail(V4629)

												tmp23760 := PrimHead(tmp23759)

												tmp23761 := PrimIsString(tmp23760)

												var ifres23751 Obj

												if True == tmp23761 {
													tmp23753 := PrimTail(V4629)

													tmp23754 := PrimTail(tmp23753)

													tmp23755 := PrimHead(tmp23754)

													tmp23756 := PrimTail(tmp23755)

													tmp23757 := PrimHead(tmp23756)

													tmp23758 := PrimIsString(tmp23757)

													var ifres23752 Obj

													if True == tmp23758 {
														ifres23752 = True

													} else {
														ifres23752 = False

													}

													ifres23751 = ifres23752

												} else {
													ifres23751 = False

												}

												var ifres23750 Obj

												if True == ifres23751 {
													ifres23750 = True

												} else {
													ifres23750 = False

												}

												ifres23749 = ifres23750

											} else {
												ifres23749 = False

											}

											var ifres23748 Obj

											if True == ifres23749 {
												ifres23748 = True

											} else {
												ifres23748 = False

											}

											ifres23747 = ifres23748

										} else {
											ifres23747 = False

										}

										var ifres23746 Obj

										if True == ifres23747 {
											ifres23746 = True

										} else {
											ifres23746 = False

										}

										ifres23745 = ifres23746

									} else {
										ifres23745 = False

									}

									var ifres23744 Obj

									if True == ifres23745 {
										ifres23744 = True

									} else {
										ifres23744 = False

									}

									ifres23743 = ifres23744

								} else {
									ifres23743 = False

								}

								var ifres23742 Obj

								if True == ifres23743 {
									ifres23742 = True

								} else {
									ifres23742 = False

								}

								ifres23741 = ifres23742

							} else {
								ifres23741 = False

							}

							var ifres23740 Obj

							if True == ifres23741 {
								ifres23740 = True

							} else {
								ifres23740 = False

							}

							ifres23739 = ifres23740

						} else {
							ifres23739 = False

						}

						var ifres23738 Obj

						if True == ifres23739 {
							ifres23738 = True

						} else {
							ifres23738 = False

						}

						ifres23737 = ifres23738

					} else {
						ifres23737 = False

					}

					var ifres23736 Obj

					if True == ifres23737 {
						ifres23736 = True

					} else {
						ifres23736 = False

					}

					ifres23735 = ifres23736

				} else {
					ifres23735 = False

				}

				var ifres23734 Obj

				if True == ifres23735 {
					ifres23734 = True

				} else {
					ifres23734 = False

				}

				ifres23733 = ifres23734

			} else {
				ifres23733 = False

			}

			var ifres23732 Obj

			if True == ifres23733 {
				ifres23732 = True

			} else {
				ifres23732 = False

			}

			ifres23731 = ifres23732

		} else {
			ifres23731 = False

		}

		if True == ifres23731 {
			tmp23717 := PrimTail(V4629)

			tmp23718 := PrimHead(tmp23717)

			tmp23719 := PrimTail(V4629)

			tmp23720 := PrimTail(tmp23719)

			tmp23721 := PrimHead(tmp23720)

			tmp23722 := PrimTail(tmp23721)

			tmp23723 := PrimHead(tmp23722)

			tmp23724 := PrimStringConcat(tmp23718, tmp23723)

			tmp23725 := PrimTail(V4629)

			tmp23726 := PrimTail(tmp23725)

			tmp23727 := PrimHead(tmp23726)

			tmp23728 := PrimTail(tmp23727)

			tmp23729 := PrimTail(tmp23728)

			tmp23730 := PrimCons(tmp23724, tmp23729)

			__e.Return(PrimCons(symcn, tmp23730))
			return

		} else {
			__e.Return(V4629)
			return
		}

	}, 1)

	tmp23801 := Call(__e, PrimNS1Value(symns2_1set), symshen_4factor_1cn, tmp23715)

	_ = tmp23801

	tmp23802 := MakeNative(func(__e *ControlFlow) {
		V4631 := __e.Get(1)
		_ = V4631
		tmp23828 := PrimEqual(MakeString(""), V4631)

		if True == tmp23828 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp23827 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4631)

			var ifres23814 Obj

			if True == tmp23827 {
				tmp23825 := PrimPos(V4631, MakeNumber(0))

				tmp23826 := PrimEqual(MakeString("~"), tmp23825)

				var ifres23816 Obj

				if True == tmp23826 {
					tmp23823 := PrimTailString(V4631)

					tmp23824 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23823)

					var ifres23818 Obj

					if True == tmp23824 {
						tmp23820 := PrimTailString(V4631)

						tmp23821 := PrimPos(tmp23820, MakeNumber(0))

						tmp23822 := PrimEqual(MakeString("%"), tmp23821)

						var ifres23819 Obj

						if True == tmp23822 {
							ifres23819 = True

						} else {
							ifres23819 = False

						}

						ifres23818 = ifres23819

					} else {
						ifres23818 = False

					}

					var ifres23817 Obj

					if True == ifres23818 {
						ifres23817 = True

					} else {
						ifres23817 = False

					}

					ifres23816 = ifres23817

				} else {
					ifres23816 = False

				}

				var ifres23815 Obj

				if True == ifres23816 {
					ifres23815 = True

				} else {
					ifres23815 = False

				}

				ifres23814 = ifres23815

			} else {
				ifres23814 = False

			}

			if True == ifres23814 {
				tmp23810 := PrimNumberToString(MakeNumber(10))

				tmp23811 := PrimTailString(V4631)

				tmp23812 := PrimTailString(tmp23811)

				tmp23813 := Call(__e, PrimNS2Value(symshen_4proc_1nl), tmp23812)

				__e.Return(PrimStringConcat(tmp23810, tmp23813))
				return

			} else {
				tmp23809 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4631)

				if True == tmp23809 {
					tmp23806 := PrimPos(V4631, MakeNumber(0))

					tmp23807 := PrimTailString(V4631)

					tmp23808 := Call(__e, PrimNS2Value(symshen_4proc_1nl), tmp23807)

					__e.Return(PrimStringConcat(tmp23806, tmp23808))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4proc_1nl)
					return
				}

			}

		}

	}, 1)

	tmp23829 := Call(__e, PrimNS1Value(symns2_1set), symshen_4proc_1nl, tmp23802)

	_ = tmp23829

	tmp23830 := MakeNative(func(__e *ControlFlow) {
		V4634 := __e.Get(1)
		_ = V4634
		V4635 := __e.Get(2)
		_ = V4635
		tmp23839 := PrimEqual(Nil, V4635)

		if True == tmp23839 {
			__e.Return(V4634)
			return
		} else {
			tmp23838 := PrimIsPair(V4635)

			if True == tmp23838 {
				tmp23833 := PrimHead(V4635)

				tmp23834 := PrimCons(V4634, Nil)

				tmp23835 := PrimCons(tmp23833, tmp23834)

				tmp23836 := PrimCons(symshen_4insert, tmp23835)

				tmp23837 := PrimTail(V4635)

				__e.TailApply(PrimNS2Value(symshen_4mkstr_1r), tmp23836, tmp23837)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4mkstr_1r)
				return
			}

		}

	}, 2)

	tmp23840 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr_1r, tmp23830)

	_ = tmp23840

	tmp23841 := MakeNative(func(__e *ControlFlow) {
		V4638 := __e.Get(1)
		_ = V4638
		V4639 := __e.Get(2)
		_ = V4639
		__e.TailApply(PrimNS2Value(symshen_4insert_1h), V4638, V4639, MakeString(""))
		return
	}, 2)

	tmp23842 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert, tmp23841)

	_ = tmp23842

	tmp23843 := MakeNative(func(__e *ControlFlow) {
		V4645 := __e.Get(1)
		_ = V4645
		V4646 := __e.Get(2)
		_ = V4646
		V4647 := __e.Get(3)
		_ = V4647
		tmp23904 := PrimEqual(MakeString(""), V4646)

		if True == tmp23904 {
			__e.Return(V4647)
			return
		} else {
			tmp23903 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4646)

			var ifres23890 Obj

			if True == tmp23903 {
				tmp23901 := PrimPos(V4646, MakeNumber(0))

				tmp23902 := PrimEqual(MakeString("~"), tmp23901)

				var ifres23892 Obj

				if True == tmp23902 {
					tmp23899 := PrimTailString(V4646)

					tmp23900 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23899)

					var ifres23894 Obj

					if True == tmp23900 {
						tmp23896 := PrimTailString(V4646)

						tmp23897 := PrimPos(tmp23896, MakeNumber(0))

						tmp23898 := PrimEqual(MakeString("A"), tmp23897)

						var ifres23895 Obj

						if True == tmp23898 {
							ifres23895 = True

						} else {
							ifres23895 = False

						}

						ifres23894 = ifres23895

					} else {
						ifres23894 = False

					}

					var ifres23893 Obj

					if True == ifres23894 {
						ifres23893 = True

					} else {
						ifres23893 = False

					}

					ifres23892 = ifres23893

				} else {
					ifres23892 = False

				}

				var ifres23891 Obj

				if True == ifres23892 {
					ifres23891 = True

				} else {
					ifres23891 = False

				}

				ifres23890 = ifres23891

			} else {
				ifres23890 = False

			}

			if True == ifres23890 {
				tmp23887 := PrimTailString(V4646)

				tmp23888 := PrimTailString(tmp23887)

				tmp23889 := Call(__e, PrimNS2Value(symshen_4app), V4645, tmp23888, symshen_4a)

				__e.Return(PrimStringConcat(V4647, tmp23889))
				return

			} else {
				tmp23886 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4646)

				var ifres23873 Obj

				if True == tmp23886 {
					tmp23884 := PrimPos(V4646, MakeNumber(0))

					tmp23885 := PrimEqual(MakeString("~"), tmp23884)

					var ifres23875 Obj

					if True == tmp23885 {
						tmp23882 := PrimTailString(V4646)

						tmp23883 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23882)

						var ifres23877 Obj

						if True == tmp23883 {
							tmp23879 := PrimTailString(V4646)

							tmp23880 := PrimPos(tmp23879, MakeNumber(0))

							tmp23881 := PrimEqual(MakeString("R"), tmp23880)

							var ifres23878 Obj

							if True == tmp23881 {
								ifres23878 = True

							} else {
								ifres23878 = False

							}

							ifres23877 = ifres23878

						} else {
							ifres23877 = False

						}

						var ifres23876 Obj

						if True == ifres23877 {
							ifres23876 = True

						} else {
							ifres23876 = False

						}

						ifres23875 = ifres23876

					} else {
						ifres23875 = False

					}

					var ifres23874 Obj

					if True == ifres23875 {
						ifres23874 = True

					} else {
						ifres23874 = False

					}

					ifres23873 = ifres23874

				} else {
					ifres23873 = False

				}

				if True == ifres23873 {
					tmp23870 := PrimTailString(V4646)

					tmp23871 := PrimTailString(tmp23870)

					tmp23872 := Call(__e, PrimNS2Value(symshen_4app), V4645, tmp23871, symshen_4r)

					__e.Return(PrimStringConcat(V4647, tmp23872))
					return

				} else {
					tmp23869 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4646)

					var ifres23856 Obj

					if True == tmp23869 {
						tmp23867 := PrimPos(V4646, MakeNumber(0))

						tmp23868 := PrimEqual(MakeString("~"), tmp23867)

						var ifres23858 Obj

						if True == tmp23868 {
							tmp23865 := PrimTailString(V4646)

							tmp23866 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp23865)

							var ifres23860 Obj

							if True == tmp23866 {
								tmp23862 := PrimTailString(V4646)

								tmp23863 := PrimPos(tmp23862, MakeNumber(0))

								tmp23864 := PrimEqual(MakeString("S"), tmp23863)

								var ifres23861 Obj

								if True == tmp23864 {
									ifres23861 = True

								} else {
									ifres23861 = False

								}

								ifres23860 = ifres23861

							} else {
								ifres23860 = False

							}

							var ifres23859 Obj

							if True == ifres23860 {
								ifres23859 = True

							} else {
								ifres23859 = False

							}

							ifres23858 = ifres23859

						} else {
							ifres23858 = False

						}

						var ifres23857 Obj

						if True == ifres23858 {
							ifres23857 = True

						} else {
							ifres23857 = False

						}

						ifres23856 = ifres23857

					} else {
						ifres23856 = False

					}

					if True == ifres23856 {
						tmp23853 := PrimTailString(V4646)

						tmp23854 := PrimTailString(tmp23853)

						tmp23855 := Call(__e, PrimNS2Value(symshen_4app), V4645, tmp23854, symshen_4s)

						__e.Return(PrimStringConcat(V4647, tmp23855))
						return

					} else {
						tmp23852 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4646)

						if True == tmp23852 {
							tmp23849 := PrimTailString(V4646)

							tmp23850 := PrimPos(V4646, MakeNumber(0))

							tmp23851 := PrimStringConcat(V4647, tmp23850)

							__e.TailApply(PrimNS2Value(symshen_4insert_1h), V4645, tmp23849, tmp23851)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4insert_1h)
							return
						}

					}

				}

			}

		}

	}, 3)

	tmp23905 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1h, tmp23843)

	_ = tmp23905

	tmp23906 := MakeNative(func(__e *ControlFlow) {
		V4651 := __e.Get(1)
		_ = V4651
		V4652 := __e.Get(2)
		_ = V4652
		V4653 := __e.Get(3)
		_ = V4653
		tmp23907 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), V4651, V4653)

		__e.Return(PrimStringConcat(tmp23907, V4652))
		return

	}, 3)

	tmp23908 := Call(__e, PrimNS1Value(symns2_1set), symshen_4app, tmp23906)

	_ = tmp23908

	tmp23909 := MakeNative(func(__e *ControlFlow) {
		V4661 := __e.Get(1)
		_ = V4661
		V4662 := __e.Get(2)
		_ = V4662
		tmp23917 := Call(__e, PrimNS2Value(symfail))

		tmp23918 := PrimEqual(V4661, tmp23917)

		if True == tmp23918 {
			__e.Return(MakeString("..."))
			return
		} else {
			tmp23916 := Call(__e, PrimNS2Value(symshen_4list_2), V4661)

			if True == tmp23916 {
				__e.TailApply(PrimNS2Value(symshen_4list_1_6str), V4661, V4662)
				return
			} else {
				tmp23915 := PrimIsString(V4661)

				if True == tmp23915 {
					__e.TailApply(PrimNS2Value(symshen_4str_1_6str), V4661, V4662)
					return
				} else {
					tmp23914 := PrimIsVector(V4661)

					if True == tmp23914 {
						__e.TailApply(PrimNS2Value(symshen_4vector_1_6str), V4661, V4662)
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4atom_1_6str), V4661)
						return
					}

				}

			}

		}

	}, 2)

	tmp23919 := Call(__e, PrimNS1Value(symns2_1set), symshen_4arg_1_6str, tmp23909)

	_ = tmp23919

	tmp23920 := MakeNative(func(__e *ControlFlow) {
		V4665 := __e.Get(1)
		_ = V4665
		V4666 := __e.Get(2)
		_ = V4666
		tmp23928 := PrimEqual(symshen_4r, V4666)

		if True == tmp23928 {
			tmp23925 := Call(__e, PrimNS2Value(symshen_4maxseq))

			tmp23926 := Call(__e, PrimNS2Value(symshen_4iter_1list), V4665, symshen_4r, tmp23925)

			tmp23927 := Call(__e, PrimNS2Value(sym_8s), tmp23926, MakeString(")"))

			__e.TailApply(PrimNS2Value(sym_8s), MakeString("("), tmp23927)
			return

		} else {
			tmp23922 := Call(__e, PrimNS2Value(symshen_4maxseq))

			tmp23923 := Call(__e, PrimNS2Value(symshen_4iter_1list), V4665, V4666, tmp23922)

			tmp23924 := Call(__e, PrimNS2Value(sym_8s), tmp23923, MakeString("]"))

			__e.TailApply(PrimNS2Value(sym_8s), MakeString("["), tmp23924)
			return

		}

	}, 2)

	tmp23929 := Call(__e, PrimNS1Value(symns2_1set), symshen_4list_1_6str, tmp23920)

	_ = tmp23929

	tmp23930 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dmaximum_1print_1sequence_1size_d))
		return
	}, 0)

	tmp23931 := Call(__e, PrimNS1Value(symns2_1set), symshen_4maxseq, tmp23930)

	_ = tmp23931

	tmp23932 := MakeNative(func(__e *ControlFlow) {
		V4680 := __e.Get(1)
		_ = V4680
		V4681 := __e.Get(2)
		_ = V4681
		V4682 := __e.Get(3)
		_ = V4682
		tmp23953 := PrimEqual(Nil, V4680)

		if True == tmp23953 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp23952 := PrimEqual(MakeNumber(0), V4682)

			if True == tmp23952 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				tmp23951 := PrimIsPair(V4680)

				var ifres23947 Obj

				if True == tmp23951 {
					tmp23949 := PrimTail(V4680)

					tmp23950 := PrimEqual(Nil, tmp23949)

					var ifres23948 Obj

					if True == tmp23950 {
						ifres23948 = True

					} else {
						ifres23948 = False

					}

					ifres23947 = ifres23948

				} else {
					ifres23947 = False

				}

				if True == ifres23947 {
					tmp23946 := PrimHead(V4680)

					__e.TailApply(PrimNS2Value(symshen_4arg_1_6str), tmp23946, V4681)
					return

				} else {
					tmp23945 := PrimIsPair(V4680)

					if True == tmp23945 {
						tmp23939 := PrimHead(V4680)

						tmp23940 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), tmp23939, V4681)

						tmp23941 := PrimTail(V4680)

						tmp23942 := PrimNumberSubtract(V4682, MakeNumber(1))

						tmp23943 := Call(__e, PrimNS2Value(symshen_4iter_1list), tmp23941, V4681, tmp23942)

						tmp23944 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp23943)

						__e.TailApply(PrimNS2Value(sym_8s), tmp23940, tmp23944)
						return

					} else {
						tmp23937 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), V4680, V4681)

						tmp23938 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp23937)

						__e.TailApply(PrimNS2Value(sym_8s), MakeString("|"), tmp23938)
						return

					}

				}

			}

		}

	}, 3)

	tmp23954 := Call(__e, PrimNS1Value(symns2_1set), symshen_4iter_1list, tmp23932)

	_ = tmp23954

	tmp23955 := MakeNative(func(__e *ControlFlow) {
		V4689 := __e.Get(1)
		_ = V4689
		V4690 := __e.Get(2)
		_ = V4690
		tmp23960 := PrimEqual(symshen_4a, V4690)

		if True == tmp23960 {
			__e.Return(V4689)
			return
		} else {
			tmp23957 := PrimNumberToString(MakeNumber(34))

			tmp23958 := PrimNumberToString(MakeNumber(34))

			tmp23959 := Call(__e, PrimNS2Value(sym_8s), V4689, tmp23958)

			__e.TailApply(PrimNS2Value(sym_8s), tmp23957, tmp23959)
			return

		}

	}, 2)

	tmp23961 := Call(__e, PrimNS1Value(symns2_1set), symshen_4str_1_6str, tmp23955)

	_ = tmp23961

	tmp23962 := MakeNative(func(__e *ControlFlow) {
		V4693 := __e.Get(1)
		_ = V4693
		V4694 := __e.Get(2)
		_ = V4694
		tmp23975 := Call(__e, PrimNS2Value(symshen_4print_1vector_2), V4693)

		if True == tmp23975 {
			tmp23973 := PrimVectorGet(V4693, MakeNumber(0))

			tmp23974 := Call(__e, PrimNS2Value(symfunction), tmp23973)

			__e.TailApply(tmp23974, V4693)
			return

		} else {
			tmp23972 := Call(__e, PrimNS2Value(symvector_2), V4693)

			if True == tmp23972 {
				tmp23969 := Call(__e, PrimNS2Value(symshen_4maxseq))

				tmp23970 := Call(__e, PrimNS2Value(symshen_4iter_1vector), V4693, MakeNumber(1), V4694, tmp23969)

				tmp23971 := Call(__e, PrimNS2Value(sym_8s), tmp23970, MakeString(">"))

				__e.TailApply(PrimNS2Value(sym_8s), MakeString("<"), tmp23971)
				return

			} else {
				tmp23965 := Call(__e, PrimNS2Value(symshen_4maxseq))

				tmp23966 := Call(__e, PrimNS2Value(symshen_4iter_1vector), V4693, MakeNumber(0), V4694, tmp23965)

				tmp23967 := Call(__e, PrimNS2Value(sym_8s), tmp23966, MakeString(">>"))

				tmp23968 := Call(__e, PrimNS2Value(sym_8s), MakeString("<"), tmp23967)

				__e.TailApply(PrimNS2Value(sym_8s), MakeString("<"), tmp23968)
				return

			}

		}

	}, 2)

	tmp23976 := Call(__e, PrimNS1Value(symns2_1set), symshen_4vector_1_6str, tmp23962)

	_ = tmp23976

	tmp23977 := MakeNative(func(__e *ControlFlow) {
		V4696 := __e.Get(1)
		_ = V4696
		tmp23978 := PrimNS3Value(symshen_4_dempty_1absvector_d)

		__e.Return(PrimEqual(V4696, tmp23978))
		return

	}, 1)

	tmp23979 := Call(__e, PrimNS1Value(symns2_1set), symshen_4empty_1absvector_2, tmp23977)

	_ = tmp23979

	tmp23980 := MakeNative(func(__e *ControlFlow) {
		V4698 := __e.Get(1)
		_ = V4698
		tmp24000 := Call(__e, PrimNS2Value(symshen_4empty_1absvector_2), V4698)

		tmp24001 := PrimNot(tmp24000)

		if True == tmp24001 {
			tmp23983 := MakeNative(func(__e *ControlFlow) {
				First := __e.Get(1)
				_ = First
				tmp23997 := PrimEqual(First, symshen_4tuple)

				if True == tmp23997 {
					__e.Return(True)
					return
				} else {
					tmp23996 := PrimEqual(First, symshen_4pvar)

					var ifres23986 Obj

					if True == tmp23996 {
						ifres23986 = True

					} else {
						tmp23995 := PrimEqual(First, symshen_4dictionary)

						var ifres23988 Obj

						if True == tmp23995 {
							ifres23988 = True

						} else {
							tmp23993 := PrimIsNumber(First)

							tmp23994 := PrimNot(tmp23993)

							var ifres23990 Obj

							if True == tmp23994 {
								tmp23992 := Call(__e, PrimNS2Value(symshen_4fbound_2), First)

								var ifres23991 Obj

								if True == tmp23992 {
									ifres23991 = True

								} else {
									ifres23991 = False

								}

								ifres23990 = ifres23991

							} else {
								ifres23990 = False

							}

							var ifres23989 Obj

							if True == ifres23990 {
								ifres23989 = True

							} else {
								ifres23989 = False

							}

							ifres23988 = ifres23989

						}

						var ifres23987 Obj

						if True == ifres23988 {
							ifres23987 = True

						} else {
							ifres23987 = False

						}

						ifres23986 = ifres23987

					}

					if True == ifres23986 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}, 1)

			tmp23998 := PrimVectorGet(V4698, MakeNumber(0))

			tmp23999 := Call(__e, tmp23983, tmp23998)

			if True == tmp23999 {
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

	tmp24002 := Call(__e, PrimNS1Value(symns2_1set), symshen_4print_1vector_2, tmp23980)

	_ = tmp24002

	tmp24003 := MakeNative(func(__e *ControlFlow) {
		V4700 := __e.Get(1)
		_ = V4700
		tmp24004 := MakeNative(func(__e *ControlFlow) {
			tmp24005 := Call(__e, PrimNS2Value(symshen_4lookup_1func), V4700)

			_ = tmp24005

			__e.Return(True)
			return

		}, 0)

		tmp24006 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp24004, tmp24006)
		return

	}, 1)

	tmp24007 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fbound_2, tmp24003)

	_ = tmp24007

	tmp24008 := MakeNative(func(__e *ControlFlow) {
		V4702 := __e.Get(1)
		_ = V4702
		tmp24009 := PrimVectorGet(V4702, MakeNumber(1))

		tmp24010 := PrimVectorGet(V4702, MakeNumber(2))

		tmp24011 := Call(__e, PrimNS2Value(symshen_4app), tmp24010, MakeString(")"), symshen_4s)

		tmp24012 := PrimStringConcat(MakeString(" "), tmp24011)

		tmp24013 := Call(__e, PrimNS2Value(symshen_4app), tmp24009, tmp24012, symshen_4s)

		__e.Return(PrimStringConcat(MakeString("(@p "), tmp24013))
		return

	}, 1)

	tmp24014 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tuple, tmp24008)

	_ = tmp24014

	tmp24015 := MakeNative(func(__e *ControlFlow) {
		V4704 := __e.Get(1)
		_ = V4704
		__e.Return(MakeString("(dict ...)"))
		return
	}, 1)

	tmp24016 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dictionary, tmp24015)

	_ = tmp24016

	tmp24017 := MakeNative(func(__e *ControlFlow) {
		V4715 := __e.Get(1)
		_ = V4715
		V4716 := __e.Get(2)
		_ = V4716
		V4717 := __e.Get(3)
		_ = V4717
		V4718 := __e.Get(4)
		_ = V4718
		tmp24037 := PrimEqual(MakeNumber(0), V4718)

		if True == tmp24037 {
			__e.Return(MakeString("... etc"))
			return
		} else {
			tmp24019 := MakeNative(func(__e *ControlFlow) {
				Item := __e.Get(1)
				_ = Item
				tmp24020 := MakeNative(func(__e *ControlFlow) {
					Next := __e.Get(1)
					_ = Next
					tmp24029 := PrimEqual(Item, symshen_4out_1of_1bounds)

					if True == tmp24029 {
						__e.Return(MakeString(""))
						return
					} else {
						tmp24028 := PrimEqual(Next, symshen_4out_1of_1bounds)

						if True == tmp24028 {
							__e.TailApply(PrimNS2Value(symshen_4arg_1_6str), Item, V4717)
							return
						} else {
							tmp24023 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), Item, V4717)

							tmp24024 := PrimNumberAdd(V4716, MakeNumber(1))

							tmp24025 := PrimNumberSubtract(V4718, MakeNumber(1))

							tmp24026 := Call(__e, PrimNS2Value(symshen_4iter_1vector), V4715, tmp24024, V4717, tmp24025)

							tmp24027 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp24026)

							__e.TailApply(PrimNS2Value(sym_8s), tmp24023, tmp24027)
							return

						}

					}

				}, 1)

				tmp24030 := MakeNative(func(__e *ControlFlow) {
					tmp24031 := PrimNumberAdd(V4716, MakeNumber(1))

					__e.Return(PrimVectorGet(V4715, tmp24031))
					return

				}, 0)

				tmp24032 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4out_1of_1bounds)
					return
				}, 1)

				tmp24033 := Call(__e, PrimNS1Value(symtry_1catch), tmp24030, tmp24032)

				__e.TailApply(tmp24020, tmp24033)
				return

			}, 1)

			tmp24034 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V4715, V4716))
				return
			}, 0)

			tmp24035 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4out_1of_1bounds)
				return
			}, 1)

			tmp24036 := Call(__e, PrimNS1Value(symtry_1catch), tmp24034, tmp24035)

			__e.TailApply(tmp24019, tmp24036)
			return

		}

	}, 4)

	tmp24038 := Call(__e, PrimNS1Value(symns2_1set), symshen_4iter_1vector, tmp24017)

	_ = tmp24038

	tmp24039 := MakeNative(func(__e *ControlFlow) {
		V4720 := __e.Get(1)
		_ = V4720
		tmp24040 := MakeNative(func(__e *ControlFlow) {
			__e.Return(PrimStr(V4720))
			return
		}, 0)

		tmp24041 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.TailApply(PrimNS2Value(symshen_4funexstring))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp24040, tmp24041)
		return

	}, 1)

	tmp24042 := Call(__e, PrimNS1Value(symns2_1set), symshen_4atom_1_6str, tmp24039)

	_ = tmp24042

	tmp24043 := MakeNative(func(__e *ControlFlow) {
		tmp24044 := PrimIntern(MakeString("x"))

		tmp24045 := Call(__e, PrimNS2Value(symgensym), tmp24044)

		tmp24046 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), tmp24045, symshen_4a)

		tmp24047 := Call(__e, PrimNS2Value(sym_8s), tmp24046, MakeString("\x11"))

		tmp24048 := Call(__e, PrimNS2Value(sym_8s), MakeString("e"), tmp24047)

		tmp24049 := Call(__e, PrimNS2Value(sym_8s), MakeString("n"), tmp24048)

		tmp24050 := Call(__e, PrimNS2Value(sym_8s), MakeString("u"), tmp24049)

		tmp24051 := Call(__e, PrimNS2Value(sym_8s), MakeString("f"), tmp24050)

		__e.TailApply(PrimNS2Value(sym_8s), MakeString("\x10"), tmp24051)
		return

	}, 0)

	tmp24052 := Call(__e, PrimNS1Value(symns2_1set), symshen_4funexstring, tmp24043)

	_ = tmp24052

	tmp24053 := MakeNative(func(__e *ControlFlow) {
		V4722 := __e.Get(1)
		_ = V4722
		tmp24057 := Call(__e, PrimNS2Value(symempty_2), V4722)

		if True == tmp24057 {
			__e.Return(True)
			return
		} else {
			tmp24056 := PrimIsPair(V4722)

			if True == tmp24056 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4list_2, tmp24053)
	return

}, 0)
