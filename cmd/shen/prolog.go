package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
	tmp9619 := MakeNative(func(__e *ControlFlow) {
		V2055 := __e.Get(1)
		_ = V2055
		V2056 := __e.Get(2)
		_ = V2056
		tmp9620 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimFunc(symshen_4_5defprolog_6), X)
			return
		}, 1)

		tmp9621 := PrimCons(V2055, V2056)

		__e.TailApply(PrimFunc(symcompile), tmp9620, tmp9621)
		return

	}, 2)

	tmp9622 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp9619)

	_ = tmp9622

	tmp9623 := MakeNative(func(__e *ControlFlow) {
		V2057 := __e.Get(1)
		_ = V2057
		tmp9624 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp9626 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9647 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2057)

		var ifres9627 Obj

		if True == tmp9647 {
			tmp9628 := MakeNative(func(__e *ControlFlow) {
				F := __e.Get(1)
				_ = F
				tmp9629 := MakeNative(func(__e *ControlFlow) {
					News1851 := __e.Get(1)
					_ = News1851
					tmp9630 := MakeNative(func(__e *ControlFlow) {
						Parseshen_4_5clauses_6 := __e.Get(1)
						_ = Parseshen_4_5clauses_6
						tmp9641 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5clauses_6)

						if True == tmp9641 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp9631 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5clauses_6)

							tmp9632 := MakeNative(func(__e *ControlFlow) {
								Aritycheck := __e.Get(1)
								_ = Aritycheck
								tmp9633 := MakeNative(func(__e *ControlFlow) {
									LeftLinear := __e.Get(1)
									_ = LeftLinear
									__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), F, LeftLinear)
									return
								}, 1)

								tmp9634 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									__e.TailApply(PrimFunc(symshen_4linearise_1clause), X)
									return
								}, 1)

								tmp9635 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5clauses_6)

								tmp9636 := Call(__e, PrimFunc(symmap), tmp9634, tmp9635)

								__e.TailApply(tmp9633, tmp9636)
								return

							}, 1)

							tmp9637 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5clauses_6)

							tmp9638 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), F, tmp9637)

							tmp9639 := Call(__e, tmp9632, tmp9638)

							__e.TailApply(PrimFunc(symshen_4comb), tmp9631, tmp9639)
							return

						}

					}, 1)

					tmp9642 := Call(__e, PrimFunc(symshen_4_5clauses_6), News1851)

					__e.TailApply(tmp9630, tmp9642)
					return

				}, 1)

				tmp9643 := Call(__e, PrimFunc(symshen_4tls), V2057)

				__e.TailApply(tmp9629, tmp9643)
				return

			}, 1)

			tmp9644 := Call(__e, PrimFunc(symshen_4hds), V2057)

			tmp9645 := Call(__e, tmp9628, tmp9644)

			ifres9627 = tmp9645

		} else {
			tmp9646 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres9627 = tmp9646

		}

		__e.TailApply(tmp9624, ifres9627)
		return

	}, 1)

	tmp9648 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp9623)

	_ = tmp9648

	tmp9649 := MakeNative(func(__e *ControlFlow) {
		V2058 := __e.Get(1)
		_ = V2058
		tmp9650 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp9668 {
				tmp9651 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp9653 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9654 := MakeNative(func(__e *ControlFlow) {
					Parse_5_b_6 := __e.Get(1)
					_ = Parse_5_b_6
					tmp9664 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5_b_6)

					if True == tmp9664 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp9655 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5_b_6)

						tmp9661 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)

						tmp9662 := Call(__e, PrimFunc(symempty_2), tmp9661)

						var ifres9656 Obj

						if True == tmp9662 {
							ifres9656 = Nil

						} else {
							tmp9657 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)

							tmp9658 := Call(__e, PrimFunc(symshen_4app), tmp9657, MakeString("\n ..."), symshen_4r)

							tmp9659 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp9658)

							tmp9660 := PrimSimpleError(tmp9659)

							ifres9656 = tmp9660

						}

						__e.TailApply(PrimFunc(symshen_4comb), tmp9655, ifres9656)
						return

					}

				}, 1)

				tmp9665 := Call(__e, PrimFunc(sym_5_b_6), V2058)

				tmp9666 := Call(__e, tmp9654, tmp9665)

				__e.TailApply(tmp9651, tmp9666)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9669 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5clause_6 := __e.Get(1)
			_ = Parseshen_4_5clause_6
			tmp9679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5clause_6)

			if True == tmp9679 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9670 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5clauses_6 := __e.Get(1)
					_ = Parseshen_4_5clauses_6
					tmp9676 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5clauses_6)

					if True == tmp9676 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp9671 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5clauses_6)

						tmp9672 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5clause_6)

						tmp9673 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5clauses_6)

						tmp9674 := PrimCons(tmp9672, tmp9673)

						__e.TailApply(PrimFunc(symshen_4comb), tmp9671, tmp9674)
						return

					}

				}, 1)

				tmp9677 := Call(__e, PrimFunc(symshen_4_5clauses_6), Parseshen_4_5clause_6)

				__e.TailApply(tmp9670, tmp9677)
				return

			}

		}, 1)

		tmp9680 := Call(__e, PrimFunc(symshen_4_5clause_6), V2058)

		tmp9681 := Call(__e, tmp9669, tmp9680)

		__e.TailApply(tmp9650, tmp9681)
		return

	}, 1)

	tmp9682 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp9649)

	_ = tmp9682

	tmp9683 := MakeNative(func(__e *ControlFlow) {
		V2063 := __e.Get(1)
		_ = V2063
		V2064 := __e.Get(2)
		_ = V2064
		tmp9710 := PrimIsPair(V2064)

		var ifres9706 Obj

		if True == tmp9710 {
			tmp9708 := PrimTail(V2064)

			tmp9709 := PrimEqual(Nil, tmp9708)

			var ifres9707 Obj

			if True == tmp9709 {
				ifres9707 = True

			} else {
				ifres9707 = False

			}

			ifres9706 = ifres9707

		} else {
			ifres9706 = False

		}

		if True == ifres9706 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp9704 := PrimIsPair(V2064)

			var ifres9689 Obj

			if True == tmp9704 {
				tmp9702 := PrimHead(V2064)

				tmp9703 := PrimIsPair(tmp9702)

				var ifres9691 Obj

				if True == tmp9703 {
					tmp9699 := PrimHead(V2064)

					tmp9700 := PrimTail(tmp9699)

					tmp9701 := PrimIsPair(tmp9700)

					var ifres9693 Obj

					if True == tmp9701 {
						tmp9695 := PrimHead(V2064)

						tmp9696 := PrimTail(tmp9695)

						tmp9697 := PrimTail(tmp9696)

						tmp9698 := PrimEqual(Nil, tmp9697)

						var ifres9694 Obj

						if True == tmp9698 {
							ifres9694 = True

						} else {
							ifres9694 = False

						}

						ifres9693 = ifres9694

					} else {
						ifres9693 = False

					}

					var ifres9692 Obj

					if True == ifres9693 {
						ifres9692 = True

					} else {
						ifres9692 = False

					}

					ifres9691 = ifres9692

				} else {
					ifres9691 = False

				}

				var ifres9690 Obj

				if True == ifres9691 {
					ifres9690 = True

				} else {
					ifres9690 = False

				}

				ifres9689 = ifres9690

			} else {
				ifres9689 = False

			}

			if True == ifres9689 {
				tmp9684 := PrimHead(V2064)

				tmp9685 := PrimHead(tmp9684)

				tmp9686 := Call(__e, PrimFunc(symlength), tmp9685)

				tmp9687 := PrimTail(V2064)

				__e.TailApply(PrimFunc(symshen_4pac_1h), V2063, tmp9686, tmp9687)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1arity_1check)
				return
			}

		}

	}, 2)

	tmp9711 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp9683)

	_ = tmp9711

	tmp9712 := MakeNative(func(__e *ControlFlow) {
		V2065 := __e.Get(1)
		_ = V2065
		tmp9728 := PrimIsPair(V2065)

		var ifres9719 Obj

		if True == tmp9728 {
			tmp9726 := PrimTail(V2065)

			tmp9727 := PrimIsPair(tmp9726)

			var ifres9721 Obj

			if True == tmp9727 {
				tmp9723 := PrimTail(V2065)

				tmp9724 := PrimTail(tmp9723)

				tmp9725 := PrimEqual(Nil, tmp9724)

				var ifres9722 Obj

				if True == tmp9725 {
					ifres9722 = True

				} else {
					ifres9722 = False

				}

				ifres9721 = ifres9722

			} else {
				ifres9721 = False

			}

			var ifres9720 Obj

			if True == ifres9721 {
				ifres9720 = True

			} else {
				ifres9720 = False

			}

			ifres9719 = ifres9720

		} else {
			ifres9719 = False

		}

		if True == ifres9719 {
			tmp9713 := PrimHead(V2065)

			tmp9714 := PrimTail(V2065)

			tmp9715 := PrimHead(tmp9714)

			tmp9716 := Call(__e, PrimFunc(sym_8p), tmp9713, tmp9715)

			tmp9717 := Call(__e, PrimFunc(symshen_4linearise), tmp9716)

			__e.TailApply(PrimFunc(symshen_4lch), tmp9717)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4linearise_1clause)
			return
		}

	}, 1)

	tmp9729 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp9712)

	_ = tmp9729

	tmp9730 := MakeNative(func(__e *ControlFlow) {
		V2066 := __e.Get(1)
		_ = V2066
		tmp9736 := Call(__e, PrimFunc(symtuple_2), V2066)

		if True == tmp9736 {
			tmp9731 := Call(__e, PrimFunc(symfst), V2066)

			tmp9732 := Call(__e, PrimFunc(symsnd), V2066)

			tmp9733 := Call(__e, PrimFunc(symshen_4lchh), tmp9732)

			tmp9734 := PrimCons(tmp9733, Nil)

			__e.Return(PrimCons(tmp9731, tmp9734))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lch)
			return
		}

	}, 1)

	tmp9737 := Call(__e, ns2_1set, symshen_4lch, tmp9730)

	_ = tmp9737

	tmp9738 := MakeNative(func(__e *ControlFlow) {
		V2067 := __e.Get(1)
		_ = V2067
		tmp9801 := PrimIsPair(V2067)

		var ifres9750 Obj

		if True == tmp9801 {
			tmp9799 := PrimHead(V2067)

			tmp9800 := PrimEqual(symwhere, tmp9799)

			var ifres9752 Obj

			if True == tmp9800 {
				tmp9797 := PrimTail(V2067)

				tmp9798 := PrimIsPair(tmp9797)

				var ifres9754 Obj

				if True == tmp9798 {
					tmp9794 := PrimTail(V2067)

					tmp9795 := PrimHead(tmp9794)

					tmp9796 := PrimIsPair(tmp9795)

					var ifres9756 Obj

					if True == tmp9796 {
						tmp9790 := PrimTail(V2067)

						tmp9791 := PrimHead(tmp9790)

						tmp9792 := PrimHead(tmp9791)

						tmp9793 := PrimEqual(sym_a, tmp9792)

						var ifres9758 Obj

						if True == tmp9793 {
							tmp9786 := PrimTail(V2067)

							tmp9787 := PrimHead(tmp9786)

							tmp9788 := PrimTail(tmp9787)

							tmp9789 := PrimIsPair(tmp9788)

							var ifres9760 Obj

							if True == tmp9789 {
								tmp9781 := PrimTail(V2067)

								tmp9782 := PrimHead(tmp9781)

								tmp9783 := PrimTail(tmp9782)

								tmp9784 := PrimTail(tmp9783)

								tmp9785 := PrimIsPair(tmp9784)

								var ifres9762 Obj

								if True == tmp9785 {
									tmp9775 := PrimTail(V2067)

									tmp9776 := PrimHead(tmp9775)

									tmp9777 := PrimTail(tmp9776)

									tmp9778 := PrimTail(tmp9777)

									tmp9779 := PrimTail(tmp9778)

									tmp9780 := PrimEqual(Nil, tmp9779)

									var ifres9764 Obj

									if True == tmp9780 {
										tmp9772 := PrimTail(V2067)

										tmp9773 := PrimTail(tmp9772)

										tmp9774 := PrimIsPair(tmp9773)

										var ifres9766 Obj

										if True == tmp9774 {
											tmp9768 := PrimTail(V2067)

											tmp9769 := PrimTail(tmp9768)

											tmp9770 := PrimTail(tmp9769)

											tmp9771 := PrimEqual(Nil, tmp9770)

											var ifres9767 Obj

											if True == tmp9771 {
												ifres9767 = True

											} else {
												ifres9767 = False

											}

											ifres9766 = ifres9767

										} else {
											ifres9766 = False

										}

										var ifres9765 Obj

										if True == ifres9766 {
											ifres9765 = True

										} else {
											ifres9765 = False

										}

										ifres9764 = ifres9765

									} else {
										ifres9764 = False

									}

									var ifres9763 Obj

									if True == ifres9764 {
										ifres9763 = True

									} else {
										ifres9763 = False

									}

									ifres9762 = ifres9763

								} else {
									ifres9762 = False

								}

								var ifres9761 Obj

								if True == ifres9762 {
									ifres9761 = True

								} else {
									ifres9761 = False

								}

								ifres9760 = ifres9761

							} else {
								ifres9760 = False

							}

							var ifres9759 Obj

							if True == ifres9760 {
								ifres9759 = True

							} else {
								ifres9759 = False

							}

							ifres9758 = ifres9759

						} else {
							ifres9758 = False

						}

						var ifres9757 Obj

						if True == ifres9758 {
							ifres9757 = True

						} else {
							ifres9757 = False

						}

						ifres9756 = ifres9757

					} else {
						ifres9756 = False

					}

					var ifres9755 Obj

					if True == ifres9756 {
						ifres9755 = True

					} else {
						ifres9755 = False

					}

					ifres9754 = ifres9755

				} else {
					ifres9754 = False

				}

				var ifres9753 Obj

				if True == ifres9754 {
					ifres9753 = True

				} else {
					ifres9753 = False

				}

				ifres9752 = ifres9753

			} else {
				ifres9752 = False

			}

			var ifres9751 Obj

			if True == ifres9752 {
				ifres9751 = True

			} else {
				ifres9751 = False

			}

			ifres9750 = ifres9751

		} else {
			ifres9750 = False

		}

		if True == ifres9750 {
			tmp9740 := PrimValue(symshen_4_doccurs_d)

			var ifres9739 Obj

			if True == tmp9740 {
				ifres9739 = symis_b

			} else {
				ifres9739 = symis

			}

			tmp9741 := PrimTail(V2067)

			tmp9742 := PrimHead(tmp9741)

			tmp9743 := PrimTail(tmp9742)

			tmp9744 := PrimCons(ifres9739, tmp9743)

			tmp9745 := PrimTail(V2067)

			tmp9746 := PrimTail(tmp9745)

			tmp9747 := PrimHead(tmp9746)

			tmp9748 := Call(__e, PrimFunc(symshen_4lchh), tmp9747)

			__e.Return(PrimCons(tmp9744, tmp9748))
			return

		} else {
			__e.Return(V2067)
			return
		}

	}, 1)

	tmp9802 := Call(__e, ns2_1set, symshen_4lchh, tmp9738)

	_ = tmp9802

	tmp9803 := MakeNative(func(__e *ControlFlow) {
		V2074 := __e.Get(1)
		_ = V2074
		V2075 := __e.Get(2)
		_ = V2075
		V2076 := __e.Get(3)
		_ = V2076
		tmp9819 := PrimEqual(Nil, V2076)

		if True == tmp9819 {
			__e.Return(True)
			return
		} else {
			tmp9817 := PrimIsPair(V2076)

			var ifres9813 Obj

			if True == tmp9817 {
				tmp9815 := PrimHead(V2076)

				tmp9816 := PrimIsPair(tmp9815)

				var ifres9814 Obj

				if True == tmp9816 {
					ifres9814 = True

				} else {
					ifres9814 = False

				}

				ifres9813 = ifres9814

			} else {
				ifres9813 = False

			}

			if True == ifres9813 {
				tmp9808 := PrimHead(V2076)

				tmp9809 := PrimHead(tmp9808)

				tmp9810 := Call(__e, PrimFunc(symlength), tmp9809)

				tmp9811 := PrimEqual(V2075, tmp9810)

				if True == tmp9811 {
					tmp9804 := PrimTail(V2076)

					__e.TailApply(PrimFunc(symshen_4pac_1h), V2074, V2075, tmp9804)
					return

				} else {
					tmp9805 := Call(__e, PrimFunc(symshen_4app), V2074, MakeString("\n"), symshen_4a)

					tmp9806 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp9805)

					__e.Return(PrimSimpleError(tmp9806))
					return

				}

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4pac_1h)
				return
			}

		}

	}, 3)

	tmp9820 := Call(__e, ns2_1set, symshen_4pac_1h, tmp9803)

	_ = tmp9820

	tmp9821 := MakeNative(func(__e *ControlFlow) {
		V2077 := __e.Get(1)
		_ = V2077
		tmp9822 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9824 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp9824 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9825 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5head_6 := __e.Get(1)
			_ = Parseshen_4_5head_6
			tmp9844 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5head_6)

			if True == tmp9844 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9842 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5head_6, sym_5_1_1)

				if True == tmp9842 {
					tmp9826 := MakeNative(func(__e *ControlFlow) {
						News1854 := __e.Get(1)
						_ = News1854
						tmp9827 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5body_6 := __e.Get(1)
							_ = Parseshen_4_5body_6
							tmp9838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5body_6)

							if True == tmp9838 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp9828 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5sc_6 := __e.Get(1)
									_ = Parseshen_4_5sc_6
									tmp9835 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

									if True == tmp9835 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp9829 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sc_6)

										tmp9830 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5head_6)

										tmp9831 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5body_6)

										tmp9832 := PrimCons(tmp9831, Nil)

										tmp9833 := PrimCons(tmp9830, tmp9832)

										__e.TailApply(PrimFunc(symshen_4comb), tmp9829, tmp9833)
										return

									}

								}, 1)

								tmp9836 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5body_6)

								__e.TailApply(tmp9828, tmp9836)
								return

							}

						}, 1)

						tmp9839 := Call(__e, PrimFunc(symshen_4_5body_6), News1854)

						__e.TailApply(tmp9827, tmp9839)
						return

					}, 1)

					tmp9840 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5head_6)

					__e.TailApply(tmp9826, tmp9840)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp9845 := Call(__e, PrimFunc(symshen_4_5head_6), V2077)

		tmp9846 := Call(__e, tmp9825, tmp9845)

		__e.TailApply(tmp9822, tmp9846)
		return

	}, 1)

	tmp9847 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp9821)

	_ = tmp9847

	tmp9848 := MakeNative(func(__e *ControlFlow) {
		V2078 := __e.Get(1)
		_ = V2078
		tmp9849 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9860 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp9860 {
				tmp9850 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9852 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp9852 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9853 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp9856 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp9856 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp9854 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp9854, Nil)
						return

					}

				}, 1)

				tmp9857 := Call(__e, PrimFunc(sym_5e_6), V2078)

				tmp9858 := Call(__e, tmp9853, tmp9857)

				__e.TailApply(tmp9850, tmp9858)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9861 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5hterm_6 := __e.Get(1)
			_ = Parseshen_4_5hterm_6
			tmp9871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

			if True == tmp9871 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9862 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5head_6 := __e.Get(1)
					_ = Parseshen_4_5head_6
					tmp9868 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5head_6)

					if True == tmp9868 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp9863 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5head_6)

						tmp9864 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

						tmp9865 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5head_6)

						tmp9866 := PrimCons(tmp9864, tmp9865)

						__e.TailApply(PrimFunc(symshen_4comb), tmp9863, tmp9866)
						return

					}

				}, 1)

				tmp9869 := Call(__e, PrimFunc(symshen_4_5head_6), Parseshen_4_5hterm_6)

				__e.TailApply(tmp9862, tmp9869)
				return

			}

		}, 1)

		tmp9872 := Call(__e, PrimFunc(symshen_4_5hterm_6), V2078)

		tmp9873 := Call(__e, tmp9861, tmp9872)

		__e.TailApply(tmp9849, tmp9873)
		return

	}, 1)

	tmp9874 := Call(__e, ns2_1set, symshen_4_5head_6, tmp9848)

	_ = tmp9874

	tmp9875 := MakeNative(func(__e *ControlFlow) {
		V2079 := __e.Get(1)
		_ = V2079
		tmp9876 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10047 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10047 {
				tmp9877 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10033 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp10033 {
						tmp9878 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp10000 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp10000 {
								tmp9879 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp9973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

									if True == tmp9973 {
										tmp9880 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp9946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

											if True == tmp9946 {
												tmp9881 := MakeNative(func(__e *ControlFlow) {
													Result := __e.Get(1)
													_ = Result
													tmp9915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

													if True == tmp9915 {
														tmp9882 := MakeNative(func(__e *ControlFlow) {
															Result := __e.Get(1)
															_ = Result
															tmp9884 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

															if True == tmp9884 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																__e.Return(Result)
																return
															}

														}, 1)

														tmp9913 := Call(__e, PrimFunc(symshen_4ccons_2), V2079)

														var ifres9885 Obj

														if True == tmp9913 {
															tmp9886 := MakeNative(func(__e *ControlFlow) {
																SynCons := __e.Get(1)
																_ = SynCons
																tmp9907 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, symmode)

																if True == tmp9907 {
																	tmp9887 := MakeNative(func(__e *ControlFlow) {
																		News1864 := __e.Get(1)
																		_ = News1864
																		tmp9888 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5hterm_6 := __e.Get(1)
																			_ = Parseshen_4_5hterm_6
																			tmp9903 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

																			if True == tmp9903 {
																				__e.TailApply(PrimFunc(symshen_4parse_1failure))
																				return
																			} else {
																				tmp9901 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5hterm_6, sym_1)

																				if True == tmp9901 {
																					tmp9889 := MakeNative(func(__e *ControlFlow) {
																						News1865 := __e.Get(1)
																						_ = News1865
																						tmp9890 := MakeNative(func(__e *ControlFlow) {
																							Parseshen_4_5end_6 := __e.Get(1)
																							_ = Parseshen_4_5end_6
																							tmp9897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

																							if True == tmp9897 {
																								__e.TailApply(PrimFunc(symshen_4parse_1failure))
																								return
																							} else {
																								tmp9891 := Call(__e, PrimFunc(symshen_4tlstream), V2079)

																								tmp9892 := Call(__e, PrimFunc(symshen_4in_1_6), tmp9891)

																								tmp9893 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

																								tmp9894 := PrimCons(tmp9893, Nil)

																								tmp9895 := PrimCons(symshen_4_1m, tmp9894)

																								__e.TailApply(PrimFunc(symshen_4comb), tmp9892, tmp9895)
																								return

																							}

																						}, 1)

																						tmp9898 := Call(__e, PrimFunc(symshen_4_5end_6), News1865)

																						__e.TailApply(tmp9890, tmp9898)
																						return

																					}, 1)

																					tmp9899 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5hterm_6)

																					__e.TailApply(tmp9889, tmp9899)
																					return

																				} else {
																					__e.TailApply(PrimFunc(symshen_4parse_1failure))
																					return
																				}

																			}

																		}, 1)

																		tmp9904 := Call(__e, PrimFunc(symshen_4_5hterm_6), News1864)

																		__e.TailApply(tmp9888, tmp9904)
																		return

																	}, 1)

																	tmp9905 := Call(__e, PrimFunc(symshen_4tls), SynCons)

																	__e.TailApply(tmp9887, tmp9905)
																	return

																} else {
																	__e.TailApply(PrimFunc(symshen_4parse_1failure))
																	return
																}

															}, 1)

															tmp9908 := Call(__e, PrimFunc(symshen_4hds), V2079)

															tmp9909 := Call(__e, PrimFunc(symshen_4_5_1out), V2079)

															tmp9910 := Call(__e, PrimFunc(symshen_4comb), tmp9908, tmp9909)

															tmp9911 := Call(__e, tmp9886, tmp9910)

															ifres9885 = tmp9911

														} else {
															tmp9912 := Call(__e, PrimFunc(symshen_4parse_1failure))

															ifres9885 = tmp9912

														}

														__e.TailApply(tmp9882, ifres9885)
														return

													} else {
														__e.Return(Result)
														return
													}

												}, 1)

												tmp9944 := Call(__e, PrimFunc(symshen_4ccons_2), V2079)

												var ifres9916 Obj

												if True == tmp9944 {
													tmp9917 := MakeNative(func(__e *ControlFlow) {
														SynCons := __e.Get(1)
														_ = SynCons
														tmp9938 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, symmode)

														if True == tmp9938 {
															tmp9918 := MakeNative(func(__e *ControlFlow) {
																News1862 := __e.Get(1)
																_ = News1862
																tmp9919 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5hterm_6 := __e.Get(1)
																	_ = Parseshen_4_5hterm_6
																	tmp9934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

																	if True == tmp9934 {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	} else {
																		tmp9932 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5hterm_6, sym_7)

																		if True == tmp9932 {
																			tmp9920 := MakeNative(func(__e *ControlFlow) {
																				News1863 := __e.Get(1)
																				_ = News1863
																				tmp9921 := MakeNative(func(__e *ControlFlow) {
																					Parseshen_4_5end_6 := __e.Get(1)
																					_ = Parseshen_4_5end_6
																					tmp9928 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

																					if True == tmp9928 {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					} else {
																						tmp9922 := Call(__e, PrimFunc(symshen_4tlstream), V2079)

																						tmp9923 := Call(__e, PrimFunc(symshen_4in_1_6), tmp9922)

																						tmp9924 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

																						tmp9925 := PrimCons(tmp9924, Nil)

																						tmp9926 := PrimCons(symshen_4_7m, tmp9925)

																						__e.TailApply(PrimFunc(symshen_4comb), tmp9923, tmp9926)
																						return

																					}

																				}, 1)

																				tmp9929 := Call(__e, PrimFunc(symshen_4_5end_6), News1863)

																				__e.TailApply(tmp9921, tmp9929)
																				return

																			}, 1)

																			tmp9930 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5hterm_6)

																			__e.TailApply(tmp9920, tmp9930)
																			return

																		} else {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		}

																	}

																}, 1)

																tmp9935 := Call(__e, PrimFunc(symshen_4_5hterm_6), News1862)

																__e.TailApply(tmp9919, tmp9935)
																return

															}, 1)

															tmp9936 := Call(__e, PrimFunc(symshen_4tls), SynCons)

															__e.TailApply(tmp9918, tmp9936)
															return

														} else {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														}

													}, 1)

													tmp9939 := Call(__e, PrimFunc(symshen_4hds), V2079)

													tmp9940 := Call(__e, PrimFunc(symshen_4_5_1out), V2079)

													tmp9941 := Call(__e, PrimFunc(symshen_4comb), tmp9939, tmp9940)

													tmp9942 := Call(__e, tmp9917, tmp9941)

													ifres9916 = tmp9942

												} else {
													tmp9943 := Call(__e, PrimFunc(symshen_4parse_1failure))

													ifres9916 = tmp9943

												}

												__e.TailApply(tmp9881, ifres9916)
												return

											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp9971 := Call(__e, PrimFunc(symshen_4ccons_2), V2079)

										var ifres9947 Obj

										if True == tmp9971 {
											tmp9948 := MakeNative(func(__e *ControlFlow) {
												SynCons := __e.Get(1)
												_ = SynCons
												tmp9965 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, sym_1)

												if True == tmp9965 {
													tmp9949 := MakeNative(func(__e *ControlFlow) {
														News1861 := __e.Get(1)
														_ = News1861
														tmp9950 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5hterm_6 := __e.Get(1)
															_ = Parseshen_4_5hterm_6
															tmp9961 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

															if True == tmp9961 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp9951 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5end_6 := __e.Get(1)
																	_ = Parseshen_4_5end_6
																	tmp9958 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

																	if True == tmp9958 {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	} else {
																		tmp9952 := Call(__e, PrimFunc(symshen_4tlstream), V2079)

																		tmp9953 := Call(__e, PrimFunc(symshen_4in_1_6), tmp9952)

																		tmp9954 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

																		tmp9955 := PrimCons(tmp9954, Nil)

																		tmp9956 := PrimCons(symshen_4_1m, tmp9955)

																		__e.TailApply(PrimFunc(symshen_4comb), tmp9953, tmp9956)
																		return

																	}

																}, 1)

																tmp9959 := Call(__e, PrimFunc(symshen_4_5end_6), Parseshen_4_5hterm_6)

																__e.TailApply(tmp9951, tmp9959)
																return

															}

														}, 1)

														tmp9962 := Call(__e, PrimFunc(symshen_4_5hterm_6), News1861)

														__e.TailApply(tmp9950, tmp9962)
														return

													}, 1)

													tmp9963 := Call(__e, PrimFunc(symshen_4tls), SynCons)

													__e.TailApply(tmp9949, tmp9963)
													return

												} else {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp9966 := Call(__e, PrimFunc(symshen_4hds), V2079)

											tmp9967 := Call(__e, PrimFunc(symshen_4_5_1out), V2079)

											tmp9968 := Call(__e, PrimFunc(symshen_4comb), tmp9966, tmp9967)

											tmp9969 := Call(__e, tmp9948, tmp9968)

											ifres9947 = tmp9969

										} else {
											tmp9970 := Call(__e, PrimFunc(symshen_4parse_1failure))

											ifres9947 = tmp9970

										}

										__e.TailApply(tmp9880, ifres9947)
										return

									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp9998 := Call(__e, PrimFunc(symshen_4ccons_2), V2079)

								var ifres9974 Obj

								if True == tmp9998 {
									tmp9975 := MakeNative(func(__e *ControlFlow) {
										SynCons := __e.Get(1)
										_ = SynCons
										tmp9992 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, sym_7)

										if True == tmp9992 {
											tmp9976 := MakeNative(func(__e *ControlFlow) {
												News1860 := __e.Get(1)
												_ = News1860
												tmp9977 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5hterm_6 := __e.Get(1)
													_ = Parseshen_4_5hterm_6
													tmp9988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

													if True == tmp9988 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp9978 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5end_6 := __e.Get(1)
															_ = Parseshen_4_5end_6
															tmp9985 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

															if True == tmp9985 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp9979 := Call(__e, PrimFunc(symshen_4tlstream), V2079)

																tmp9980 := Call(__e, PrimFunc(symshen_4in_1_6), tmp9979)

																tmp9981 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

																tmp9982 := PrimCons(tmp9981, Nil)

																tmp9983 := PrimCons(symshen_4_7m, tmp9982)

																__e.TailApply(PrimFunc(symshen_4comb), tmp9980, tmp9983)
																return

															}

														}, 1)

														tmp9986 := Call(__e, PrimFunc(symshen_4_5end_6), Parseshen_4_5hterm_6)

														__e.TailApply(tmp9978, tmp9986)
														return

													}

												}, 1)

												tmp9989 := Call(__e, PrimFunc(symshen_4_5hterm_6), News1860)

												__e.TailApply(tmp9977, tmp9989)
												return

											}, 1)

											tmp9990 := Call(__e, PrimFunc(symshen_4tls), SynCons)

											__e.TailApply(tmp9976, tmp9990)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp9993 := Call(__e, PrimFunc(symshen_4hds), V2079)

									tmp9994 := Call(__e, PrimFunc(symshen_4_5_1out), V2079)

									tmp9995 := Call(__e, PrimFunc(symshen_4comb), tmp9993, tmp9994)

									tmp9996 := Call(__e, tmp9975, tmp9995)

									ifres9974 = tmp9996

								} else {
									tmp9997 := Call(__e, PrimFunc(symshen_4parse_1failure))

									ifres9974 = tmp9997

								}

								__e.TailApply(tmp9879, ifres9974)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp10031 := Call(__e, PrimFunc(symshen_4ccons_2), V2079)

						var ifres10001 Obj

						if True == tmp10031 {
							tmp10002 := MakeNative(func(__e *ControlFlow) {
								SynCons := __e.Get(1)
								_ = SynCons
								tmp10025 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, symcons)

								if True == tmp10025 {
									tmp10003 := MakeNative(func(__e *ControlFlow) {
										News1859 := __e.Get(1)
										_ = News1859
										tmp10004 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5hterm1_6 := __e.Get(1)
											_ = Parseshen_4_5hterm1_6
											tmp10021 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm1_6)

											if True == tmp10021 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp10005 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5hterm2_6 := __e.Get(1)
													_ = Parseshen_4_5hterm2_6
													tmp10018 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm2_6)

													if True == tmp10018 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp10006 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5end_6 := __e.Get(1)
															_ = Parseshen_4_5end_6
															tmp10015 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

															if True == tmp10015 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp10007 := Call(__e, PrimFunc(symshen_4tlstream), V2079)

																tmp10008 := Call(__e, PrimFunc(symshen_4in_1_6), tmp10007)

																tmp10009 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm1_6)

																tmp10010 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm2_6)

																tmp10011 := PrimCons(tmp10010, Nil)

																tmp10012 := PrimCons(tmp10009, tmp10011)

																tmp10013 := PrimCons(symcons, tmp10012)

																__e.TailApply(PrimFunc(symshen_4comb), tmp10008, tmp10013)
																return

															}

														}, 1)

														tmp10016 := Call(__e, PrimFunc(symshen_4_5end_6), Parseshen_4_5hterm2_6)

														__e.TailApply(tmp10006, tmp10016)
														return

													}

												}, 1)

												tmp10019 := Call(__e, PrimFunc(symshen_4_5hterm2_6), Parseshen_4_5hterm1_6)

												__e.TailApply(tmp10005, tmp10019)
												return

											}

										}, 1)

										tmp10022 := Call(__e, PrimFunc(symshen_4_5hterm1_6), News1859)

										__e.TailApply(tmp10004, tmp10022)
										return

									}, 1)

									tmp10023 := Call(__e, PrimFunc(symshen_4tls), SynCons)

									__e.TailApply(tmp10003, tmp10023)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp10026 := Call(__e, PrimFunc(symshen_4hds), V2079)

							tmp10027 := Call(__e, PrimFunc(symshen_4_5_1out), V2079)

							tmp10028 := Call(__e, PrimFunc(symshen_4comb), tmp10026, tmp10027)

							tmp10029 := Call(__e, tmp10002, tmp10028)

							ifres10001 = tmp10029

						} else {
							tmp10030 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres10001 = tmp10030

						}

						__e.TailApply(tmp9878, ifres10001)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10045 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2079)

				var ifres10034 Obj

				if True == tmp10045 {
					tmp10035 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp10036 := MakeNative(func(__e *ControlFlow) {
							News1858 := __e.Get(1)
							_ = News1858
							tmp10039 := PrimIntern(MakeString(":"))

							tmp10040 := PrimEqual(X, tmp10039)

							if True == tmp10040 {
								tmp10037 := Call(__e, PrimFunc(symshen_4in_1_6), News1858)

								__e.TailApply(PrimFunc(symshen_4comb), tmp10037, X)
								return

							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp10041 := Call(__e, PrimFunc(symshen_4tls), V2079)

						__e.TailApply(tmp10036, tmp10041)
						return

					}, 1)

					tmp10042 := Call(__e, PrimFunc(symshen_4hds), V2079)

					tmp10043 := Call(__e, tmp10035, tmp10042)

					ifres10034 = tmp10043

				} else {
					tmp10044 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres10034 = tmp10044

				}

				__e.TailApply(tmp9877, ifres10034)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10062 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2079)

		var ifres10048 Obj

		if True == tmp10062 {
			tmp10049 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp10050 := MakeNative(func(__e *ControlFlow) {
					News1857 := __e.Get(1)
					_ = News1857
					tmp10057 := Call(__e, PrimFunc(symatom_2), X)

					var ifres10053 Obj

					if True == tmp10057 {
						tmp10055 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), X)

						tmp10056 := PrimNot(tmp10055)

						var ifres10054 Obj

						if True == tmp10056 {
							ifres10054 = True

						} else {
							ifres10054 = False

						}

						ifres10053 = ifres10054

					} else {
						ifres10053 = False

					}

					if True == ifres10053 {
						tmp10051 := Call(__e, PrimFunc(symshen_4in_1_6), News1857)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10051, X)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10058 := Call(__e, PrimFunc(symshen_4tls), V2079)

				__e.TailApply(tmp10050, tmp10058)
				return

			}, 1)

			tmp10059 := Call(__e, PrimFunc(symshen_4hds), V2079)

			tmp10060 := Call(__e, tmp10049, tmp10059)

			ifres10048 = tmp10060

		} else {
			tmp10061 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10048 = tmp10061

		}

		__e.TailApply(tmp9876, ifres10048)
		return

	}, 1)

	tmp10063 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp9875)

	_ = tmp10063

	tmp10064 := MakeNative(func(__e *ControlFlow) {
		V2080 := __e.Get(1)
		_ = V2080
		tmp10065 := PrimIntern(MakeString(";"))

		tmp10066 := PrimCons(sym_5_1_1, Nil)

		tmp10067 := PrimCons(tmp10065, tmp10066)

		__e.TailApply(PrimFunc(symelement_2), V2080, tmp10067)
		return

	}, 1)

	tmp10068 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10064)

	_ = tmp10068

	tmp10069 := MakeNative(func(__e *ControlFlow) {
		V2081 := __e.Get(1)
		_ = V2081
		tmp10082 := PrimIsSymbol(V2081)

		if True == tmp10082 {
			__e.Return(True)
			return
		} else {
			tmp10080 := PrimIsString(V2081)

			var ifres10071 Obj

			if True == tmp10080 {
				ifres10071 = True

			} else {
				tmp10079 := Call(__e, PrimFunc(symboolean_2), V2081)

				var ifres10073 Obj

				if True == tmp10079 {
					ifres10073 = True

				} else {
					tmp10078 := PrimIsNumber(V2081)

					var ifres10075 Obj

					if True == tmp10078 {
						ifres10075 = True

					} else {
						tmp10077 := Call(__e, PrimFunc(symempty_2), V2081)

						var ifres10076 Obj

						if True == tmp10077 {
							ifres10076 = True

						} else {
							ifres10076 = False

						}

						ifres10075 = ifres10076

					}

					var ifres10074 Obj

					if True == ifres10075 {
						ifres10074 = True

					} else {
						ifres10074 = False

					}

					ifres10073 = ifres10074

				}

				var ifres10072 Obj

				if True == ifres10073 {
					ifres10072 = True

				} else {
					ifres10072 = False

				}

				ifres10071 = ifres10072

			}

			if True == ifres10071 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp10083 := Call(__e, ns2_1set, symatom_2, tmp10069)

	_ = tmp10083

	tmp10084 := MakeNative(func(__e *ControlFlow) {
		V2082 := __e.Get(1)
		_ = V2082
		tmp10085 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10087 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10087 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10088 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5hterm_6 := __e.Get(1)
			_ = Parseshen_4_5hterm_6
			tmp10092 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

			if True == tmp10092 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10089 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5hterm_6)

				tmp10090 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp10089, tmp10090)
				return

			}

		}, 1)

		tmp10093 := Call(__e, PrimFunc(symshen_4_5hterm_6), V2082)

		tmp10094 := Call(__e, tmp10088, tmp10093)

		__e.TailApply(tmp10085, tmp10094)
		return

	}, 1)

	tmp10095 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10084)

	_ = tmp10095

	tmp10096 := MakeNative(func(__e *ControlFlow) {
		V2083 := __e.Get(1)
		_ = V2083
		tmp10097 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10099 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10099 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10100 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5hterm_6 := __e.Get(1)
			_ = Parseshen_4_5hterm_6
			tmp10104 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

			if True == tmp10104 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10101 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5hterm_6)

				tmp10102 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5hterm_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp10101, tmp10102)
				return

			}

		}, 1)

		tmp10105 := Call(__e, PrimFunc(symshen_4_5hterm_6), V2083)

		tmp10106 := Call(__e, tmp10100, tmp10105)

		__e.TailApply(tmp10097, tmp10106)
		return

	}, 1)

	tmp10107 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10096)

	_ = tmp10107

	tmp10108 := MakeNative(func(__e *ControlFlow) {
		V2084 := __e.Get(1)
		_ = V2084
		tmp10109 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10120 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10120 {
				tmp10110 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10112 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp10112 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10113 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp10116 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp10116 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10114 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10114, Nil)
						return

					}

				}, 1)

				tmp10117 := Call(__e, PrimFunc(sym_5e_6), V2084)

				tmp10118 := Call(__e, tmp10113, tmp10117)

				__e.TailApply(tmp10110, tmp10118)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10121 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5literal_6 := __e.Get(1)
			_ = Parseshen_4_5literal_6
			tmp10131 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5literal_6)

			if True == tmp10131 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10122 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5body_6 := __e.Get(1)
					_ = Parseshen_4_5body_6
					tmp10128 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5body_6)

					if True == tmp10128 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10123 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5body_6)

						tmp10124 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5literal_6)

						tmp10125 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5body_6)

						tmp10126 := PrimCons(tmp10124, tmp10125)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10123, tmp10126)
						return

					}

				}, 1)

				tmp10129 := Call(__e, PrimFunc(symshen_4_5body_6), Parseshen_4_5literal_6)

				__e.TailApply(tmp10122, tmp10129)
				return

			}

		}, 1)

		tmp10132 := Call(__e, PrimFunc(symshen_4_5literal_6), V2084)

		tmp10133 := Call(__e, tmp10121, tmp10132)

		__e.TailApply(tmp10109, tmp10133)
		return

	}, 1)

	tmp10134 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10108)

	_ = tmp10134

	tmp10135 := MakeNative(func(__e *ControlFlow) {
		V2085 := __e.Get(1)
		_ = V2085
		tmp10136 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10160 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10160 {
				tmp10137 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10139 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp10139 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10158 := Call(__e, PrimFunc(symshen_4ccons_2), V2085)

				var ifres10140 Obj

				if True == tmp10158 {
					tmp10141 := MakeNative(func(__e *ControlFlow) {
						SynCons := __e.Get(1)
						_ = SynCons
						tmp10142 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5bterms_6 := __e.Get(1)
							_ = Parseshen_4_5bterms_6
							tmp10151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5bterms_6)

							if True == tmp10151 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10143 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5end_6 := __e.Get(1)
									_ = Parseshen_4_5end_6
									tmp10148 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

									if True == tmp10148 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp10144 := Call(__e, PrimFunc(symshen_4tlstream), V2085)

										tmp10145 := Call(__e, PrimFunc(symshen_4in_1_6), tmp10144)

										tmp10146 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5bterms_6)

										__e.TailApply(PrimFunc(symshen_4comb), tmp10145, tmp10146)
										return

									}

								}, 1)

								tmp10149 := Call(__e, PrimFunc(symshen_4_5end_6), Parseshen_4_5bterms_6)

								__e.TailApply(tmp10143, tmp10149)
								return

							}

						}, 1)

						tmp10152 := Call(__e, PrimFunc(symshen_4_5bterms_6), SynCons)

						__e.TailApply(tmp10142, tmp10152)
						return

					}, 1)

					tmp10153 := Call(__e, PrimFunc(symshen_4hds), V2085)

					tmp10154 := Call(__e, PrimFunc(symshen_4_5_1out), V2085)

					tmp10155 := Call(__e, PrimFunc(symshen_4comb), tmp10153, tmp10154)

					tmp10156 := Call(__e, tmp10141, tmp10155)

					ifres10140 = tmp10156

				} else {
					tmp10157 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres10140 = tmp10157

				}

				__e.TailApply(tmp10137, ifres10140)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10167 := Call(__e, PrimFunc(symshen_4_ahd_2), V2085, sym_b)

		var ifres10161 Obj

		if True == tmp10167 {
			tmp10162 := MakeNative(func(__e *ControlFlow) {
				News1870 := __e.Get(1)
				_ = News1870
				tmp10163 := Call(__e, PrimFunc(symshen_4in_1_6), News1870)

				__e.TailApply(PrimFunc(symshen_4comb), tmp10163, sym_b)
				return

			}, 1)

			tmp10164 := Call(__e, PrimFunc(symshen_4tls), V2085)

			tmp10165 := Call(__e, tmp10162, tmp10164)

			ifres10161 = tmp10165

		} else {
			tmp10166 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10161 = tmp10166

		}

		__e.TailApply(tmp10136, ifres10161)
		return

	}, 1)

	tmp10168 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10135)

	_ = tmp10168

	tmp10169 := MakeNative(func(__e *ControlFlow) {
		V2086 := __e.Get(1)
		_ = V2086
		tmp10170 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10181 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10181 {
				tmp10171 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10173 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp10173 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10174 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp10177 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp10177 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10175 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10175, Nil)
						return

					}

				}, 1)

				tmp10178 := Call(__e, PrimFunc(sym_5e_6), V2086)

				tmp10179 := Call(__e, tmp10174, tmp10178)

				__e.TailApply(tmp10171, tmp10179)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10182 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5bterm_6 := __e.Get(1)
			_ = Parseshen_4_5bterm_6
			tmp10192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5bterm_6)

			if True == tmp10192 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10183 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5bterms_6 := __e.Get(1)
					_ = Parseshen_4_5bterms_6
					tmp10189 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5bterms_6)

					if True == tmp10189 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10184 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5bterms_6)

						tmp10185 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5bterm_6)

						tmp10186 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5bterms_6)

						tmp10187 := PrimCons(tmp10185, tmp10186)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10184, tmp10187)
						return

					}

				}, 1)

				tmp10190 := Call(__e, PrimFunc(symshen_4_5bterms_6), Parseshen_4_5bterm_6)

				__e.TailApply(tmp10183, tmp10190)
				return

			}

		}, 1)

		tmp10193 := Call(__e, PrimFunc(symshen_4_5bterm_6), V2086)

		tmp10194 := Call(__e, tmp10182, tmp10193)

		__e.TailApply(tmp10170, tmp10194)
		return

	}, 1)

	tmp10195 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10169)

	_ = tmp10195

	tmp10196 := MakeNative(func(__e *ControlFlow) {
		V2087 := __e.Get(1)
		_ = V2087
		tmp10197 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10235 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10235 {
				tmp10198 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10222 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp10222 {
						tmp10199 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp10201 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp10201 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp10220 := Call(__e, PrimFunc(symshen_4ccons_2), V2087)

						var ifres10202 Obj

						if True == tmp10220 {
							tmp10203 := MakeNative(func(__e *ControlFlow) {
								SynCons := __e.Get(1)
								_ = SynCons
								tmp10204 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5bterms_6 := __e.Get(1)
									_ = Parseshen_4_5bterms_6
									tmp10213 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5bterms_6)

									if True == tmp10213 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp10205 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5end_6 := __e.Get(1)
											_ = Parseshen_4_5end_6
											tmp10210 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)

											if True == tmp10210 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp10206 := Call(__e, PrimFunc(symshen_4tlstream), V2087)

												tmp10207 := Call(__e, PrimFunc(symshen_4in_1_6), tmp10206)

												tmp10208 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5bterms_6)

												__e.TailApply(PrimFunc(symshen_4comb), tmp10207, tmp10208)
												return

											}

										}, 1)

										tmp10211 := Call(__e, PrimFunc(symshen_4_5end_6), Parseshen_4_5bterms_6)

										__e.TailApply(tmp10205, tmp10211)
										return

									}

								}, 1)

								tmp10214 := Call(__e, PrimFunc(symshen_4_5bterms_6), SynCons)

								__e.TailApply(tmp10204, tmp10214)
								return

							}, 1)

							tmp10215 := Call(__e, PrimFunc(symshen_4hds), V2087)

							tmp10216 := Call(__e, PrimFunc(symshen_4_5_1out), V2087)

							tmp10217 := Call(__e, PrimFunc(symshen_4comb), tmp10215, tmp10216)

							tmp10218 := Call(__e, tmp10203, tmp10217)

							ifres10202 = tmp10218

						} else {
							tmp10219 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres10202 = tmp10219

						}

						__e.TailApply(tmp10199, ifres10202)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10233 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2087)

				var ifres10223 Obj

				if True == tmp10233 {
					tmp10224 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp10225 := MakeNative(func(__e *ControlFlow) {
							News1873 := __e.Get(1)
							_ = News1873
							tmp10228 := Call(__e, PrimFunc(symatom_2), X)

							if True == tmp10228 {
								tmp10226 := Call(__e, PrimFunc(symshen_4in_1_6), News1873)

								__e.TailApply(PrimFunc(symshen_4comb), tmp10226, X)
								return

							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp10229 := Call(__e, PrimFunc(symshen_4tls), V2087)

						__e.TailApply(tmp10225, tmp10229)
						return

					}, 1)

					tmp10230 := Call(__e, PrimFunc(symshen_4hds), V2087)

					tmp10231 := Call(__e, tmp10224, tmp10230)

					ifres10223 = tmp10231

				} else {
					tmp10232 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres10223 = tmp10232

				}

				__e.TailApply(tmp10198, ifres10223)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10236 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5wildcard_6 := __e.Get(1)
			_ = Parseshen_4_5wildcard_6
			tmp10240 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5wildcard_6)

			if True == tmp10240 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10237 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5wildcard_6)

				tmp10238 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5wildcard_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp10237, tmp10238)
				return

			}

		}, 1)

		tmp10241 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V2087)

		tmp10242 := Call(__e, tmp10236, tmp10241)

		__e.TailApply(tmp10197, tmp10242)
		return

	}, 1)

	tmp10243 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10196)

	_ = tmp10243

	tmp10244 := MakeNative(func(__e *ControlFlow) {
		V2088 := __e.Get(1)
		_ = V2088
		tmp10245 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10247 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10247 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10260 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2088)

		var ifres10248 Obj

		if True == tmp10260 {
			tmp10249 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp10250 := MakeNative(func(__e *ControlFlow) {
					News1875 := __e.Get(1)
					_ = News1875
					tmp10255 := PrimEqual(X, sym__)

					if True == tmp10255 {
						tmp10251 := Call(__e, PrimFunc(symshen_4in_1_6), News1875)

						tmp10252 := Call(__e, PrimFunc(symprotect), symY)

						tmp10253 := Call(__e, PrimFunc(symgensym), tmp10252)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10251, tmp10253)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10256 := Call(__e, PrimFunc(symshen_4tls), V2088)

				__e.TailApply(tmp10250, tmp10256)
				return

			}, 1)

			tmp10257 := Call(__e, PrimFunc(symshen_4hds), V2088)

			tmp10258 := Call(__e, tmp10249, tmp10257)

			ifres10248 = tmp10258

		} else {
			tmp10259 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10248 = tmp10259

		}

		__e.TailApply(tmp10245, ifres10248)
		return

	}, 1)

	tmp10261 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10244)

	_ = tmp10261

	tmp10262 := MakeNative(func(__e *ControlFlow) {
		V2089 := __e.Get(1)
		_ = V2089
		tmp10263 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp10265 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10276 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2089)

		var ifres10266 Obj

		if True == tmp10276 {
			tmp10267 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp10268 := MakeNative(func(__e *ControlFlow) {
					News1877 := __e.Get(1)
					_ = News1877
					tmp10271 := Call(__e, PrimFunc(symshen_4semicolon_2), X)

					if True == tmp10271 {
						tmp10269 := Call(__e, PrimFunc(symshen_4in_1_6), News1877)

						__e.TailApply(PrimFunc(symshen_4comb), tmp10269, X)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10272 := Call(__e, PrimFunc(symshen_4tls), V2089)

				__e.TailApply(tmp10268, tmp10272)
				return

			}, 1)

			tmp10273 := Call(__e, PrimFunc(symshen_4hds), V2089)

			tmp10274 := Call(__e, tmp10267, tmp10273)

			ifres10266 = tmp10274

		} else {
			tmp10275 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10266 = tmp10275

		}

		__e.TailApply(tmp10263, ifres10266)
		return

	}, 1)

	tmp10277 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10262)

	_ = tmp10277

	tmp10278 := MakeNative(func(__e *ControlFlow) {
		V2090 := __e.Get(1)
		_ = V2090
		tmp10279 := PrimIntern(MakeString(";"))

		__e.Return(PrimEqual(V2090, tmp10279))
		return

	}, 1)

	tmp10280 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp10278)

	_ = tmp10280

	tmp10281 := MakeNative(func(__e *ControlFlow) {
		V2091 := __e.Get(1)
		_ = V2091
		V2092 := __e.Get(2)
		_ = V2092
		tmp10282 := MakeNative(func(__e *ControlFlow) {
			Bindings := __e.Get(1)
			_ = Bindings
			tmp10283 := MakeNative(func(__e *ControlFlow) {
				Lock := __e.Get(1)
				_ = Lock
				tmp10284 := MakeNative(func(__e *ControlFlow) {
					Key := __e.Get(1)
					_ = Key
					tmp10285 := MakeNative(func(__e *ControlFlow) {
						Continuation := __e.Get(1)
						_ = Continuation
						tmp10286 := MakeNative(func(__e *ControlFlow) {
							Parameters := __e.Get(1)
							_ = Parameters
							tmp10287 := MakeNative(func(__e *ControlFlow) {
								HasCut_2 := __e.Get(1)
								_ = HasCut_2
								tmp10288 := MakeNative(func(__e *ControlFlow) {
									FBody := __e.Get(1)
									_ = FBody
									tmp10289 := MakeNative(func(__e *ControlFlow) {
										CutFBody := __e.Get(1)
										_ = CutFBody
										tmp10290 := MakeNative(func(__e *ControlFlow) {
											Shen := __e.Get(1)
											_ = Shen
											__e.Return(Shen)
											return
										}, 1)

										tmp10291 := PrimCons(sym_1_6, Nil)

										tmp10292 := PrimCons(Continuation, tmp10291)

										tmp10293 := PrimCons(Key, tmp10292)

										tmp10294 := PrimCons(Lock, tmp10293)

										tmp10295 := PrimCons(Bindings, tmp10294)

										tmp10296 := PrimCons(CutFBody, Nil)

										tmp10297 := Call(__e, PrimFunc(symappend), tmp10295, tmp10296)

										tmp10298 := Call(__e, PrimFunc(symappend), Parameters, tmp10297)

										tmp10299 := PrimCons(V2091, tmp10298)

										tmp10300 := PrimCons(symdefine, tmp10299)

										__e.TailApply(tmp10290, tmp10300)
										return

									}, 1)

									var ifres10301 Obj

									if True == HasCut_2 {
										tmp10302 := PrimCons(MakeNumber(1), Nil)

										tmp10303 := PrimCons(Key, tmp10302)

										tmp10304 := PrimCons(sym_7, tmp10303)

										tmp10305 := PrimCons(FBody, Nil)

										tmp10306 := PrimCons(tmp10304, tmp10305)

										tmp10307 := PrimCons(Key, tmp10306)

										tmp10308 := PrimCons(symlet, tmp10307)

										ifres10301 = tmp10308

									} else {
										ifres10301 = FBody

									}

									__e.TailApply(tmp10289, ifres10301)
									return

								}, 1)

								tmp10309 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V2092, Parameters, Bindings, Lock, Key, Continuation, HasCut_2)

								__e.TailApply(tmp10288, tmp10309)
								return

							}, 1)

							tmp10310 := Call(__e, PrimFunc(symshen_4hascut_2), V2092)

							__e.TailApply(tmp10287, tmp10310)
							return

						}, 1)

						tmp10311 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V2092)

						__e.TailApply(tmp10286, tmp10311)
						return

					}, 1)

					tmp10312 := Call(__e, PrimFunc(symprotect), symC)

					tmp10313 := Call(__e, PrimFunc(symgensym), tmp10312)

					__e.TailApply(tmp10285, tmp10313)
					return

				}, 1)

				tmp10314 := Call(__e, PrimFunc(symprotect), symK)

				tmp10315 := Call(__e, PrimFunc(symgensym), tmp10314)

				__e.TailApply(tmp10284, tmp10315)
				return

			}, 1)

			tmp10316 := Call(__e, PrimFunc(symprotect), symL)

			tmp10317 := Call(__e, PrimFunc(symgensym), tmp10316)

			__e.TailApply(tmp10283, tmp10317)
			return

		}, 1)

		tmp10318 := Call(__e, PrimFunc(symprotect), symB)

		tmp10319 := Call(__e, PrimFunc(symgensym), tmp10318)

		__e.TailApply(tmp10282, tmp10319)
		return

	}, 2)

	tmp10320 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10281)

	_ = tmp10320

	tmp10321 := MakeNative(func(__e *ControlFlow) {
		V2095 := __e.Get(1)
		_ = V2095
		tmp10331 := PrimEqual(sym_b, V2095)

		if True == tmp10331 {
			__e.Return(True)
			return
		} else {
			tmp10329 := PrimIsPair(V2095)

			if True == tmp10329 {
				tmp10326 := PrimHead(V2095)

				tmp10327 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10326)

				if True == tmp10327 {
					__e.Return(True)
					return
				} else {
					tmp10323 := PrimTail(V2095)

					tmp10324 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10323)

					if True == tmp10324 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp10332 := Call(__e, ns2_1set, symshen_4hascut_2, tmp10321)

	_ = tmp10332

	tmp10333 := MakeNative(func(__e *ControlFlow) {
		V2100 := __e.Get(1)
		_ = V2100
		tmp10342 := PrimIsPair(V2100)

		var ifres10338 Obj

		if True == tmp10342 {
			tmp10340 := PrimHead(V2100)

			tmp10341 := PrimIsPair(tmp10340)

			var ifres10339 Obj

			if True == tmp10341 {
				ifres10339 = True

			} else {
				ifres10339 = False

			}

			ifres10338 = ifres10339

		} else {
			ifres10338 = False

		}

		if True == ifres10338 {
			tmp10334 := PrimHead(V2100)

			tmp10335 := PrimHead(tmp10334)

			tmp10336 := Call(__e, PrimFunc(symlength), tmp10335)

			__e.TailApply(PrimFunc(symshen_4parameters), tmp10336)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1parameters)
			return
		}

	}, 1)

	tmp10343 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp10333)

	_ = tmp10343

	tmp10344 := MakeNative(func(__e *ControlFlow) {
		V2121 := __e.Get(1)
		_ = V2121
		V2122 := __e.Get(2)
		_ = V2122
		V2123 := __e.Get(3)
		_ = V2123
		V2124 := __e.Get(4)
		_ = V2124
		V2125 := __e.Get(5)
		_ = V2125
		V2126 := __e.Get(6)
		_ = V2126
		V2127 := __e.Get(7)
		_ = V2127
		tmp10438 := PrimEqual(Nil, V2121)

		var ifres10435 Obj

		if True == tmp10438 {
			tmp10437 := PrimEqual(True, V2127)

			var ifres10436 Obj

			if True == tmp10437 {
				ifres10436 = True

			} else {
				ifres10436 = False

			}

			ifres10435 = ifres10436

		} else {
			ifres10435 = False

		}

		if True == ifres10435 {
			tmp10345 := PrimCons(V2125, Nil)

			tmp10346 := PrimCons(V2124, tmp10345)

			__e.Return(PrimCons(symshen_4unlock, tmp10346))
			return

		} else {
			tmp10433 := PrimIsPair(V2121)

			var ifres10411 Obj

			if True == tmp10433 {
				tmp10431 := PrimHead(V2121)

				tmp10432 := PrimIsPair(tmp10431)

				var ifres10413 Obj

				if True == tmp10432 {
					tmp10428 := PrimHead(V2121)

					tmp10429 := PrimTail(tmp10428)

					tmp10430 := PrimIsPair(tmp10429)

					var ifres10415 Obj

					if True == tmp10430 {
						tmp10424 := PrimHead(V2121)

						tmp10425 := PrimTail(tmp10424)

						tmp10426 := PrimTail(tmp10425)

						tmp10427 := PrimEqual(Nil, tmp10426)

						var ifres10417 Obj

						if True == tmp10427 {
							tmp10422 := PrimTail(V2121)

							tmp10423 := PrimEqual(Nil, tmp10422)

							var ifres10419 Obj

							if True == tmp10423 {
								tmp10421 := PrimEqual(False, V2127)

								var ifres10420 Obj

								if True == tmp10421 {
									ifres10420 = True

								} else {
									ifres10420 = False

								}

								ifres10419 = ifres10420

							} else {
								ifres10419 = False

							}

							var ifres10418 Obj

							if True == ifres10419 {
								ifres10418 = True

							} else {
								ifres10418 = False

							}

							ifres10417 = ifres10418

						} else {
							ifres10417 = False

						}

						var ifres10416 Obj

						if True == ifres10417 {
							ifres10416 = True

						} else {
							ifres10416 = False

						}

						ifres10415 = ifres10416

					} else {
						ifres10415 = False

					}

					var ifres10414 Obj

					if True == ifres10415 {
						ifres10414 = True

					} else {
						ifres10414 = False

					}

					ifres10413 = ifres10414

				} else {
					ifres10413 = False

				}

				var ifres10412 Obj

				if True == ifres10413 {
					ifres10412 = True

				} else {
					ifres10412 = False

				}

				ifres10411 = ifres10412

			} else {
				ifres10411 = False

			}

			if True == ifres10411 {
				tmp10347 := MakeNative(func(__e *ControlFlow) {
					Continue := __e.Get(1)
					_ = Continue
					tmp10348 := PrimCons(V2124, Nil)

					tmp10349 := PrimCons(symshen_4unlocked_2, tmp10348)

					tmp10350 := PrimHead(V2121)

					tmp10351 := PrimHead(tmp10350)

					tmp10352 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10351, V2122, V2123, Continue)

					tmp10353 := PrimCons(False, Nil)

					tmp10354 := PrimCons(tmp10352, tmp10353)

					tmp10355 := PrimCons(tmp10349, tmp10354)

					__e.Return(PrimCons(symif, tmp10355))
					return

				}, 1)

				tmp10356 := PrimHead(V2121)

				tmp10357 := PrimHead(tmp10356)

				tmp10358 := PrimHead(V2121)

				tmp10359 := PrimTail(tmp10358)

				tmp10360 := PrimHead(tmp10359)

				tmp10361 := Call(__e, PrimFunc(symshen_4continue), tmp10357, tmp10360, V2123, V2124, V2125, V2126)

				__e.TailApply(tmp10347, tmp10361)
				return

			} else {
				tmp10409 := PrimIsPair(V2121)

				var ifres10394 Obj

				if True == tmp10409 {
					tmp10407 := PrimHead(V2121)

					tmp10408 := PrimIsPair(tmp10407)

					var ifres10396 Obj

					if True == tmp10408 {
						tmp10404 := PrimHead(V2121)

						tmp10405 := PrimTail(tmp10404)

						tmp10406 := PrimIsPair(tmp10405)

						var ifres10398 Obj

						if True == tmp10406 {
							tmp10400 := PrimHead(V2121)

							tmp10401 := PrimTail(tmp10400)

							tmp10402 := PrimTail(tmp10401)

							tmp10403 := PrimEqual(Nil, tmp10402)

							var ifres10399 Obj

							if True == tmp10403 {
								ifres10399 = True

							} else {
								ifres10399 = False

							}

							ifres10398 = ifres10399

						} else {
							ifres10398 = False

						}

						var ifres10397 Obj

						if True == ifres10398 {
							ifres10397 = True

						} else {
							ifres10397 = False

						}

						ifres10396 = ifres10397

					} else {
						ifres10396 = False

					}

					var ifres10395 Obj

					if True == ifres10396 {
						ifres10395 = True

					} else {
						ifres10395 = False

					}

					ifres10394 = ifres10395

				} else {
					ifres10394 = False

				}

				if True == ifres10394 {
					tmp10362 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp10363 := MakeNative(func(__e *ControlFlow) {
							Continue := __e.Get(1)
							_ = Continue
							tmp10364 := PrimCons(V2124, Nil)

							tmp10365 := PrimCons(symshen_4unlocked_2, tmp10364)

							tmp10366 := PrimHead(V2121)

							tmp10367 := PrimHead(tmp10366)

							tmp10368 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10367, V2122, V2123, Continue)

							tmp10369 := PrimCons(False, Nil)

							tmp10370 := PrimCons(tmp10368, tmp10369)

							tmp10371 := PrimCons(tmp10365, tmp10370)

							tmp10372 := PrimCons(symif, tmp10371)

							tmp10373 := PrimCons(False, Nil)

							tmp10374 := PrimCons(Case, tmp10373)

							tmp10375 := PrimCons(sym_a, tmp10374)

							tmp10376 := PrimTail(V2121)

							tmp10377 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp10376, V2122, V2123, V2124, V2125, V2126, V2127)

							tmp10378 := PrimCons(Case, Nil)

							tmp10379 := PrimCons(tmp10377, tmp10378)

							tmp10380 := PrimCons(tmp10375, tmp10379)

							tmp10381 := PrimCons(symif, tmp10380)

							tmp10382 := PrimCons(tmp10381, Nil)

							tmp10383 := PrimCons(tmp10372, tmp10382)

							tmp10384 := PrimCons(Case, tmp10383)

							__e.Return(PrimCons(symlet, tmp10384))
							return

						}, 1)

						tmp10385 := PrimHead(V2121)

						tmp10386 := PrimHead(tmp10385)

						tmp10387 := PrimHead(V2121)

						tmp10388 := PrimTail(tmp10387)

						tmp10389 := PrimHead(tmp10388)

						tmp10390 := Call(__e, PrimFunc(symshen_4continue), tmp10386, tmp10389, V2123, V2124, V2125, V2126)

						__e.TailApply(tmp10363, tmp10390)
						return

					}, 1)

					tmp10391 := Call(__e, PrimFunc(symgensym), symC)

					tmp10392 := Call(__e, PrimFunc(symprotect), tmp10391)

					__e.TailApply(tmp10362, tmp10392)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
					return
				}

			}

		}

	}, 7)

	tmp10439 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp10344)

	_ = tmp10439

	tmp10440 := MakeNative(func(__e *ControlFlow) {
		V2128 := __e.Get(1)
		_ = V2128
		V2129 := __e.Get(2)
		_ = V2129
		tmp10445 := Call(__e, PrimFunc(symshen_4locked_2), V2128)

		var ifres10442 Obj

		if True == tmp10445 {
			tmp10444 := Call(__e, PrimFunc(symshen_4fits_2), V2129, V2128)

			var ifres10443 Obj

			if True == tmp10444 {
				ifres10443 = True

			} else {
				ifres10443 = False

			}

			ifres10442 = ifres10443

		} else {
			ifres10442 = False

		}

		if True == ifres10442 {
			__e.TailApply(PrimFunc(symshen_4openlock), V2128)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp10446 := Call(__e, ns2_1set, symshen_4unlock, tmp10440)

	_ = tmp10446

	tmp10447 := MakeNative(func(__e *ControlFlow) {
		V2130 := __e.Get(1)
		_ = V2130
		tmp10448 := Call(__e, PrimFunc(symshen_4unlocked_2), V2130)

		__e.Return(PrimNot(tmp10448))
		return

	}, 1)

	tmp10449 := Call(__e, ns2_1set, symshen_4locked_2, tmp10447)

	_ = tmp10449

	tmp10450 := MakeNative(func(__e *ControlFlow) {
		V2131 := __e.Get(1)
		_ = V2131
		__e.Return(PrimVectorGet(V2131, MakeNumber(1)))
		return
	}, 1)

	tmp10451 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp10450)

	_ = tmp10451

	tmp10452 := MakeNative(func(__e *ControlFlow) {
		V2132 := __e.Get(1)
		_ = V2132
		tmp10453 := PrimVectorSet(V2132, MakeNumber(1), True)

		_ = tmp10453

		__e.Return(False)
		return

	}, 1)

	tmp10454 := Call(__e, ns2_1set, symshen_4openlock, tmp10452)

	_ = tmp10454

	tmp10455 := MakeNative(func(__e *ControlFlow) {
		V2133 := __e.Get(1)
		_ = V2133
		V2134 := __e.Get(2)
		_ = V2134
		tmp10456 := PrimVectorGet(V2134, MakeNumber(2))

		__e.Return(PrimEqual(V2133, tmp10456))
		return

	}, 2)

	tmp10457 := Call(__e, ns2_1set, symshen_4fits_2, tmp10455)

	_ = tmp10457

	tmp10458 := MakeNative(func(__e *ControlFlow) {
		V2137 := __e.Get(1)
		_ = V2137
		V2138 := __e.Get(2)
		_ = V2138
		V2139 := __e.Get(3)
		_ = V2139
		V2140 := __e.Get(4)
		_ = V2140
		tmp10459 := MakeNative(func(__e *ControlFlow) {
			Compute := __e.Get(1)
			_ = Compute
			tmp10464 := PrimEqual(Compute, False)

			var ifres10461 Obj

			if True == tmp10464 {
				tmp10463 := Call(__e, PrimFunc(symshen_4unlocked_2), V2138)

				var ifres10462 Obj

				if True == tmp10463 {
					ifres10462 = True

				} else {
					ifres10462 = False

				}

				ifres10461 = ifres10462

			} else {
				ifres10461 = False

			}

			if True == ifres10461 {
				__e.TailApply(PrimFunc(symshen_4lock), V2139, V2138)
				return
			} else {
				__e.Return(Compute)
				return
			}

		}, 1)

		tmp10465 := Call(__e, PrimFunc(symthaw), V2140)

		__e.TailApply(tmp10459, tmp10465)
		return

	}, 4)

	tmp10466 := Call(__e, ns2_1set, symshen_4cut, tmp10458)

	_ = tmp10466

	tmp10467 := MakeNative(func(__e *ControlFlow) {
		V2141 := __e.Get(1)
		_ = V2141
		V2142 := __e.Get(2)
		_ = V2142
		tmp10468 := MakeNative(func(__e *ControlFlow) {
			SetLock := __e.Get(1)
			_ = SetLock
			tmp10469 := MakeNative(func(__e *ControlFlow) {
				SetKey := __e.Get(1)
				_ = SetKey
				__e.Return(False)
				return
			}, 1)

			tmp10470 := PrimVectorSet(V2142, MakeNumber(2), V2141)

			__e.TailApply(tmp10469, tmp10470)
			return

		}, 1)

		tmp10471 := PrimVectorSet(V2142, MakeNumber(1), False)

		__e.TailApply(tmp10468, tmp10471)
		return

	}, 2)

	tmp10472 := Call(__e, ns2_1set, symshen_4lock, tmp10467)

	_ = tmp10472

	tmp10473 := MakeNative(func(__e *ControlFlow) {
		V2143 := __e.Get(1)
		_ = V2143
		V2144 := __e.Get(2)
		_ = V2144
		V2145 := __e.Get(3)
		_ = V2145
		V2146 := __e.Get(4)
		_ = V2146
		V2147 := __e.Get(5)
		_ = V2147
		V2148 := __e.Get(6)
		_ = V2148
		tmp10474 := MakeNative(func(__e *ControlFlow) {
			HVs := __e.Get(1)
			_ = HVs
			tmp10475 := MakeNative(func(__e *ControlFlow) {
				BVs := __e.Get(1)
				_ = BVs
				tmp10476 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					tmp10477 := MakeNative(func(__e *ControlFlow) {
						ContinuationCode := __e.Get(1)
						_ = ContinuationCode
						__e.TailApply(PrimFunc(symshen_4stpart), Free, ContinuationCode, V2145)
						return
					}, 1)

					tmp10478 := PrimCons(symshen_4incinfs, Nil)

					tmp10479 := Call(__e, PrimFunc(symshen_4compile_1body), V2144, V2145, V2146, V2147, V2148)

					tmp10480 := PrimCons(tmp10479, Nil)

					tmp10481 := PrimCons(tmp10478, tmp10480)

					tmp10482 := PrimCons(symdo, tmp10481)

					__e.TailApply(tmp10477, tmp10482)
					return

				}, 1)

				tmp10483 := Call(__e, PrimFunc(symdifference), BVs, HVs)

				__e.TailApply(tmp10476, tmp10483)
				return

			}, 1)

			tmp10484 := Call(__e, PrimFunc(symshen_4extract_1vars), V2144)

			__e.TailApply(tmp10475, tmp10484)
			return

		}, 1)

		tmp10485 := Call(__e, PrimFunc(symshen_4extract_1vars), V2143)

		__e.TailApply(tmp10474, tmp10485)
		return

	}, 6)

	tmp10486 := Call(__e, ns2_1set, symshen_4continue, tmp10473)

	_ = tmp10486

	tmp10487 := MakeNative(func(__e *ControlFlow) {
		V2165 := __e.Get(1)
		_ = V2165
		V2166 := __e.Get(2)
		_ = V2166
		V2167 := __e.Get(3)
		_ = V2167
		V2168 := __e.Get(4)
		_ = V2168
		V2169 := __e.Get(5)
		_ = V2169
		tmp10522 := PrimEqual(Nil, V2165)

		if True == tmp10522 {
			tmp10488 := PrimCons(V2169, Nil)

			__e.Return(PrimCons(symthaw, tmp10488))
			return

		} else {
			tmp10520 := PrimIsPair(V2165)

			var ifres10516 Obj

			if True == tmp10520 {
				tmp10518 := PrimHead(V2165)

				tmp10519 := PrimEqual(sym_b, tmp10518)

				var ifres10517 Obj

				if True == tmp10519 {
					ifres10517 = True

				} else {
					ifres10517 = False

				}

				ifres10516 = ifres10517

			} else {
				ifres10516 = False

			}

			if True == ifres10516 {
				tmp10489 := PrimCons(symshen_4cut, Nil)

				tmp10490 := PrimTail(V2165)

				tmp10491 := PrimCons(tmp10489, tmp10490)

				__e.TailApply(PrimFunc(symshen_4compile_1body), tmp10491, V2166, V2167, V2168, V2169)
				return

			} else {
				tmp10514 := PrimIsPair(V2165)

				var ifres10510 Obj

				if True == tmp10514 {
					tmp10512 := PrimTail(V2165)

					tmp10513 := PrimEqual(Nil, tmp10512)

					var ifres10511 Obj

					if True == tmp10513 {
						ifres10511 = True

					} else {
						ifres10511 = False

					}

					ifres10510 = ifres10511

				} else {
					ifres10510 = False

				}

				if True == ifres10510 {
					tmp10492 := PrimHead(V2165)

					tmp10493 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp10492, V2166)

					tmp10494 := PrimCons(V2169, Nil)

					tmp10495 := PrimCons(V2168, tmp10494)

					tmp10496 := PrimCons(V2167, tmp10495)

					tmp10497 := PrimCons(V2166, tmp10496)

					__e.TailApply(PrimFunc(symappend), tmp10493, tmp10497)
					return

				} else {
					tmp10508 := PrimIsPair(V2165)

					if True == tmp10508 {
						tmp10498 := MakeNative(func(__e *ControlFlow) {
							P_d := __e.Get(1)
							_ = P_d
							tmp10499 := PrimTail(V2165)

							tmp10500 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp10499, V2166, V2167, V2168, V2169)

							tmp10501 := PrimCons(tmp10500, Nil)

							tmp10502 := PrimCons(V2168, tmp10501)

							tmp10503 := PrimCons(V2167, tmp10502)

							tmp10504 := PrimCons(V2166, tmp10503)

							__e.TailApply(PrimFunc(symappend), P_d, tmp10504)
							return

						}, 1)

						tmp10505 := PrimHead(V2165)

						tmp10506 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp10505, V2166)

						__e.TailApply(tmp10498, tmp10506)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
						return
					}

				}

			}

		}

	}, 5)

	tmp10523 := Call(__e, ns2_1set, symshen_4compile_1body, tmp10487)

	_ = tmp10523

	tmp10524 := MakeNative(func(__e *ControlFlow) {
		V2186 := __e.Get(1)
		_ = V2186
		V2187 := __e.Get(2)
		_ = V2187
		V2188 := __e.Get(3)
		_ = V2188
		V2189 := __e.Get(4)
		_ = V2189
		V2190 := __e.Get(5)
		_ = V2190
		tmp10548 := PrimEqual(Nil, V2186)

		if True == tmp10548 {
			__e.Return(V2190)
			return
		} else {
			tmp10546 := PrimIsPair(V2186)

			var ifres10542 Obj

			if True == tmp10546 {
				tmp10544 := PrimHead(V2186)

				tmp10545 := PrimEqual(sym_b, tmp10544)

				var ifres10543 Obj

				if True == tmp10545 {
					ifres10543 = True

				} else {
					ifres10543 = False

				}

				ifres10542 = ifres10543

			} else {
				ifres10542 = False

			}

			if True == ifres10542 {
				tmp10525 := PrimCons(symshen_4cut, Nil)

				tmp10526 := PrimTail(V2186)

				tmp10527 := PrimCons(tmp10525, tmp10526)

				__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp10527, V2187, V2188, V2189, V2190)
				return

			} else {
				tmp10540 := PrimIsPair(V2186)

				if True == tmp10540 {
					tmp10528 := MakeNative(func(__e *ControlFlow) {
						P_d := __e.Get(1)
						_ = P_d
						tmp10529 := PrimTail(V2186)

						tmp10530 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp10529, V2187, V2188, V2189, V2190)

						tmp10531 := PrimCons(tmp10530, Nil)

						tmp10532 := PrimCons(V2189, tmp10531)

						tmp10533 := PrimCons(V2188, tmp10532)

						tmp10534 := PrimCons(V2187, tmp10533)

						tmp10535 := Call(__e, PrimFunc(symappend), P_d, tmp10534)

						tmp10536 := PrimCons(tmp10535, Nil)

						__e.Return(PrimCons(symfreeze, tmp10536))
						return

					}, 1)

					tmp10537 := PrimHead(V2186)

					tmp10538 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp10537, V2187)

					__e.TailApply(tmp10528, tmp10538)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
					return
				}

			}

		}

	}, 5)

	tmp10549 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp10524)

	_ = tmp10549

	tmp10550 := MakeNative(func(__e *ControlFlow) {
		V2195 := __e.Get(1)
		_ = V2195
		V2196 := __e.Get(2)
		_ = V2196
		tmp10565 := PrimIsPair(V2195)

		var ifres10561 Obj

		if True == tmp10565 {
			tmp10563 := PrimHead(V2195)

			tmp10564 := PrimEqual(symfork, tmp10563)

			var ifres10562 Obj

			if True == tmp10564 {
				ifres10562 = True

			} else {
				ifres10562 = False

			}

			ifres10561 = ifres10562

		} else {
			ifres10561 = False

		}

		if True == ifres10561 {
			tmp10551 := PrimTail(V2195)

			tmp10552 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp10551, V2196)

			tmp10553 := PrimCons(tmp10552, Nil)

			__e.Return(PrimCons(symfork, tmp10553))
			return

		} else {
			tmp10559 := PrimIsPair(V2195)

			if True == tmp10559 {
				tmp10554 := PrimHead(V2195)

				tmp10555 := MakeNative(func(__e *ControlFlow) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(PrimFunc(symshen_4function_1calls), Y, V2196)
					return
				}, 1)

				tmp10556 := PrimTail(V2195)

				tmp10557 := Call(__e, PrimFunc(symmap), tmp10555, tmp10556)

				__e.Return(PrimCons(tmp10554, tmp10557))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
				return
			}

		}

	}, 2)

	tmp10566 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp10550)

	_ = tmp10566

	tmp10567 := MakeNative(func(__e *ControlFlow) {
		V2203 := __e.Get(1)
		_ = V2203
		V2204 := __e.Get(2)
		_ = V2204
		tmp10577 := PrimEqual(Nil, V2203)

		if True == tmp10577 {
			__e.Return(Nil)
			return
		} else {
			tmp10575 := PrimIsPair(V2203)

			if True == tmp10575 {
				tmp10568 := PrimHead(V2203)

				tmp10569 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp10568, V2204)

				tmp10570 := PrimTail(V2203)

				tmp10571 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp10570, V2204)

				tmp10572 := PrimCons(tmp10571, Nil)

				tmp10573 := PrimCons(tmp10569, tmp10572)

				__e.Return(PrimCons(symcons, tmp10573))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
				return
			}

		}

	}, 2)

	tmp10578 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp10567)

	_ = tmp10578

	tmp10579 := MakeNative(func(__e *ControlFlow) {
		V2207 := __e.Get(1)
		_ = V2207
		V2208 := __e.Get(2)
		_ = V2208
		tmp10611 := PrimIsPair(V2207)

		var ifres10592 Obj

		if True == tmp10611 {
			tmp10609 := PrimHead(V2207)

			tmp10610 := PrimEqual(symcons, tmp10609)

			var ifres10594 Obj

			if True == tmp10610 {
				tmp10607 := PrimTail(V2207)

				tmp10608 := PrimIsPair(tmp10607)

				var ifres10596 Obj

				if True == tmp10608 {
					tmp10604 := PrimTail(V2207)

					tmp10605 := PrimTail(tmp10604)

					tmp10606 := PrimIsPair(tmp10605)

					var ifres10598 Obj

					if True == tmp10606 {
						tmp10600 := PrimTail(V2207)

						tmp10601 := PrimTail(tmp10600)

						tmp10602 := PrimTail(tmp10601)

						tmp10603 := PrimEqual(Nil, tmp10602)

						var ifres10599 Obj

						if True == tmp10603 {
							ifres10599 = True

						} else {
							ifres10599 = False

						}

						ifres10598 = ifres10599

					} else {
						ifres10598 = False

					}

					var ifres10597 Obj

					if True == ifres10598 {
						ifres10597 = True

					} else {
						ifres10597 = False

					}

					ifres10596 = ifres10597

				} else {
					ifres10596 = False

				}

				var ifres10595 Obj

				if True == ifres10596 {
					ifres10595 = True

				} else {
					ifres10595 = False

				}

				ifres10594 = ifres10595

			} else {
				ifres10594 = False

			}

			var ifres10593 Obj

			if True == ifres10594 {
				ifres10593 = True

			} else {
				ifres10593 = False

			}

			ifres10592 = ifres10593

		} else {
			ifres10592 = False

		}

		if True == ifres10592 {
			tmp10580 := PrimTail(V2207)

			tmp10581 := PrimHead(tmp10580)

			tmp10582 := Call(__e, PrimFunc(symshen_4function_1calls), tmp10581, V2208)

			tmp10583 := PrimTail(V2207)

			tmp10584 := PrimTail(tmp10583)

			tmp10585 := PrimHead(tmp10584)

			tmp10586 := Call(__e, PrimFunc(symshen_4function_1calls), tmp10585, V2208)

			tmp10587 := PrimCons(tmp10586, Nil)

			tmp10588 := PrimCons(tmp10582, tmp10587)

			__e.Return(PrimCons(symcons, tmp10588))
			return

		} else {
			tmp10590 := PrimIsPair(V2207)

			if True == tmp10590 {
				__e.TailApply(PrimFunc(symshen_4deref_1terms), V2207, V2208)
				return
			} else {
				__e.Return(V2207)
				return
			}

		}

	}, 2)

	tmp10612 := Call(__e, ns2_1set, symshen_4function_1calls, tmp10579)

	_ = tmp10612

	tmp10613 := MakeNative(func(__e *ControlFlow) {
		V2213 := __e.Get(1)
		_ = V2213
		V2214 := __e.Get(2)
		_ = V2214
		tmp10671 := PrimIsPair(V2213)

		var ifres10658 Obj

		if True == tmp10671 {
			tmp10669 := PrimHead(V2213)

			tmp10670 := PrimEqual(MakeNumber(0), tmp10669)

			var ifres10660 Obj

			if True == tmp10670 {
				tmp10667 := PrimTail(V2213)

				tmp10668 := PrimIsPair(tmp10667)

				var ifres10662 Obj

				if True == tmp10668 {
					tmp10664 := PrimTail(V2213)

					tmp10665 := PrimTail(tmp10664)

					tmp10666 := PrimEqual(Nil, tmp10665)

					var ifres10663 Obj

					if True == tmp10666 {
						ifres10663 = True

					} else {
						ifres10663 = False

					}

					ifres10662 = ifres10663

				} else {
					ifres10662 = False

				}

				var ifres10661 Obj

				if True == ifres10662 {
					ifres10661 = True

				} else {
					ifres10661 = False

				}

				ifres10660 = ifres10661

			} else {
				ifres10660 = False

			}

			var ifres10659 Obj

			if True == ifres10660 {
				ifres10659 = True

			} else {
				ifres10659 = False

			}

			ifres10658 = ifres10659

		} else {
			ifres10658 = False

		}

		if True == ifres10658 {
			tmp10620 := PrimTail(V2213)

			tmp10621 := PrimHead(tmp10620)

			tmp10622 := PrimIsVariable(tmp10621)

			if True == tmp10622 {
				tmp10614 := PrimTail(V2213)

				__e.Return(PrimHead(tmp10614))
				return

			} else {
				tmp10615 := PrimTail(V2213)

				tmp10616 := PrimHead(tmp10615)

				tmp10617 := Call(__e, PrimFunc(symshen_4app), tmp10616, MakeString("\n"), symshen_4s)

				tmp10618 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp10617)

				__e.Return(PrimSimpleError(tmp10618))
				return

			}

		} else {
			tmp10656 := PrimIsPair(V2213)

			var ifres10643 Obj

			if True == tmp10656 {
				tmp10654 := PrimHead(V2213)

				tmp10655 := PrimEqual(MakeNumber(1), tmp10654)

				var ifres10645 Obj

				if True == tmp10655 {
					tmp10652 := PrimTail(V2213)

					tmp10653 := PrimIsPair(tmp10652)

					var ifres10647 Obj

					if True == tmp10653 {
						tmp10649 := PrimTail(V2213)

						tmp10650 := PrimTail(tmp10649)

						tmp10651 := PrimEqual(Nil, tmp10650)

						var ifres10648 Obj

						if True == tmp10651 {
							ifres10648 = True

						} else {
							ifres10648 = False

						}

						ifres10647 = ifres10648

					} else {
						ifres10647 = False

					}

					var ifres10646 Obj

					if True == ifres10647 {
						ifres10646 = True

					} else {
						ifres10646 = False

					}

					ifres10645 = ifres10646

				} else {
					ifres10645 = False

				}

				var ifres10644 Obj

				if True == ifres10645 {
					ifres10644 = True

				} else {
					ifres10644 = False

				}

				ifres10643 = ifres10644

			} else {
				ifres10643 = False

			}

			if True == ifres10643 {
				tmp10632 := PrimTail(V2213)

				tmp10633 := PrimHead(tmp10632)

				tmp10634 := PrimIsVariable(tmp10633)

				if True == tmp10634 {
					tmp10623 := PrimTail(V2213)

					tmp10624 := PrimHead(tmp10623)

					tmp10625 := PrimCons(V2214, Nil)

					tmp10626 := PrimCons(tmp10624, tmp10625)

					__e.Return(PrimCons(symshen_4lazyderef, tmp10626))
					return

				} else {
					tmp10627 := PrimTail(V2213)

					tmp10628 := PrimHead(tmp10627)

					tmp10629 := Call(__e, PrimFunc(symshen_4app), tmp10628, MakeString("\n"), symshen_4s)

					tmp10630 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp10629)

					__e.Return(PrimSimpleError(tmp10630))
					return

				}

			} else {
				tmp10641 := PrimIsVariable(V2213)

				if True == tmp10641 {
					tmp10635 := PrimCons(V2214, Nil)

					tmp10636 := PrimCons(V2213, tmp10635)

					__e.Return(PrimCons(symshen_4deref, tmp10636))
					return

				} else {
					tmp10639 := PrimIsPair(V2213)

					if True == tmp10639 {
						tmp10637 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(PrimFunc(symshen_4deref_1terms), Z, V2214)
							return
						}, 1)

						__e.TailApply(PrimFunc(symmap), tmp10637, V2213)
						return

					} else {
						__e.Return(V2213)
						return
					}

				}

			}

		}

	}, 2)

	tmp10672 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp10613)

	_ = tmp10672

	tmp10673 := MakeNative(func(__e *ControlFlow) {
		V2232 := __e.Get(1)
		_ = V2232
		V2233 := __e.Get(2)
		_ = V2233
		V2234 := __e.Get(3)
		_ = V2234
		V2235 := __e.Get(4)
		_ = V2235
		V2236 := __e.Get(5)
		_ = V2236
		tmp10849 := PrimEqual(Nil, V2233)

		var ifres10846 Obj

		if True == tmp10849 {
			tmp10848 := PrimEqual(Nil, V2234)

			var ifres10847 Obj

			if True == tmp10848 {
				ifres10847 = True

			} else {
				ifres10847 = False

			}

			ifres10846 = ifres10847

		} else {
			ifres10846 = False

		}

		if True == ifres10846 {
			__e.Return(V2236)
			return
		} else {
			tmp10844 := PrimIsPair(V2233)

			var ifres10824 Obj

			if True == tmp10844 {
				tmp10842 := PrimHead(V2233)

				tmp10843 := PrimIsPair(tmp10842)

				var ifres10826 Obj

				if True == tmp10843 {
					tmp10839 := PrimHead(V2233)

					tmp10840 := PrimHead(tmp10839)

					tmp10841 := PrimEqual(symshen_4_7m, tmp10840)

					var ifres10828 Obj

					if True == tmp10841 {
						tmp10836 := PrimHead(V2233)

						tmp10837 := PrimTail(tmp10836)

						tmp10838 := PrimIsPair(tmp10837)

						var ifres10830 Obj

						if True == tmp10838 {
							tmp10832 := PrimHead(V2233)

							tmp10833 := PrimTail(tmp10832)

							tmp10834 := PrimTail(tmp10833)

							tmp10835 := PrimEqual(Nil, tmp10834)

							var ifres10831 Obj

							if True == tmp10835 {
								ifres10831 = True

							} else {
								ifres10831 = False

							}

							ifres10830 = ifres10831

						} else {
							ifres10830 = False

						}

						var ifres10829 Obj

						if True == ifres10830 {
							ifres10829 = True

						} else {
							ifres10829 = False

						}

						ifres10828 = ifres10829

					} else {
						ifres10828 = False

					}

					var ifres10827 Obj

					if True == ifres10828 {
						ifres10827 = True

					} else {
						ifres10827 = False

					}

					ifres10826 = ifres10827

				} else {
					ifres10826 = False

				}

				var ifres10825 Obj

				if True == ifres10826 {
					ifres10825 = True

				} else {
					ifres10825 = False

				}

				ifres10824 = ifres10825

			} else {
				ifres10824 = False

			}

			if True == ifres10824 {
				tmp10674 := PrimHead(V2233)

				tmp10675 := PrimTail(tmp10674)

				tmp10676 := PrimHead(tmp10675)

				tmp10677 := PrimTail(V2233)

				tmp10678 := PrimCons(V2232, tmp10677)

				tmp10679 := PrimCons(tmp10676, tmp10678)

				tmp10680 := PrimCons(symshen_4_7m, tmp10679)

				__e.TailApply(PrimFunc(symshen_4compile_1head), V2232, tmp10680, V2234, V2235, V2236)
				return

			} else {
				tmp10822 := PrimIsPair(V2233)

				var ifres10802 Obj

				if True == tmp10822 {
					tmp10820 := PrimHead(V2233)

					tmp10821 := PrimIsPair(tmp10820)

					var ifres10804 Obj

					if True == tmp10821 {
						tmp10817 := PrimHead(V2233)

						tmp10818 := PrimHead(tmp10817)

						tmp10819 := PrimEqual(symshen_4_1m, tmp10818)

						var ifres10806 Obj

						if True == tmp10819 {
							tmp10814 := PrimHead(V2233)

							tmp10815 := PrimTail(tmp10814)

							tmp10816 := PrimIsPair(tmp10815)

							var ifres10808 Obj

							if True == tmp10816 {
								tmp10810 := PrimHead(V2233)

								tmp10811 := PrimTail(tmp10810)

								tmp10812 := PrimTail(tmp10811)

								tmp10813 := PrimEqual(Nil, tmp10812)

								var ifres10809 Obj

								if True == tmp10813 {
									ifres10809 = True

								} else {
									ifres10809 = False

								}

								ifres10808 = ifres10809

							} else {
								ifres10808 = False

							}

							var ifres10807 Obj

							if True == ifres10808 {
								ifres10807 = True

							} else {
								ifres10807 = False

							}

							ifres10806 = ifres10807

						} else {
							ifres10806 = False

						}

						var ifres10805 Obj

						if True == ifres10806 {
							ifres10805 = True

						} else {
							ifres10805 = False

						}

						ifres10804 = ifres10805

					} else {
						ifres10804 = False

					}

					var ifres10803 Obj

					if True == ifres10804 {
						ifres10803 = True

					} else {
						ifres10803 = False

					}

					ifres10802 = ifres10803

				} else {
					ifres10802 = False

				}

				if True == ifres10802 {
					tmp10681 := PrimHead(V2233)

					tmp10682 := PrimTail(tmp10681)

					tmp10683 := PrimHead(tmp10682)

					tmp10684 := PrimTail(V2233)

					tmp10685 := PrimCons(V2232, tmp10684)

					tmp10686 := PrimCons(tmp10683, tmp10685)

					tmp10687 := PrimCons(symshen_4_1m, tmp10686)

					__e.TailApply(PrimFunc(symshen_4compile_1head), V2232, tmp10687, V2234, V2235, V2236)
					return

				} else {
					tmp10800 := PrimIsPair(V2233)

					var ifres10796 Obj

					if True == tmp10800 {
						tmp10798 := PrimHead(V2233)

						tmp10799 := PrimEqual(symshen_4_1m, tmp10798)

						var ifres10797 Obj

						if True == tmp10799 {
							ifres10797 = True

						} else {
							ifres10797 = False

						}

						ifres10796 = ifres10797

					} else {
						ifres10796 = False

					}

					if True == ifres10796 {
						tmp10688 := PrimTail(V2233)

						__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp10688, V2234, V2235, V2236)
						return

					} else {
						tmp10794 := PrimIsPair(V2233)

						var ifres10790 Obj

						if True == tmp10794 {
							tmp10792 := PrimHead(V2233)

							tmp10793 := PrimEqual(symshen_4_7m, tmp10792)

							var ifres10791 Obj

							if True == tmp10793 {
								ifres10791 = True

							} else {
								ifres10791 = False

							}

							ifres10790 = ifres10791

						} else {
							ifres10790 = False

						}

						if True == ifres10790 {
							tmp10689 := PrimTail(V2233)

							__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10689, V2234, V2235, V2236)
							return

						} else {
							tmp10788 := PrimIsPair(V2233)

							var ifres10781 Obj

							if True == tmp10788 {
								tmp10787 := PrimIsPair(V2234)

								var ifres10783 Obj

								if True == tmp10787 {
									tmp10785 := PrimHead(V2233)

									tmp10786 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp10785)

									var ifres10784 Obj

									if True == tmp10786 {
										ifres10784 = True

									} else {
										ifres10784 = False

									}

									ifres10783 = ifres10784

								} else {
									ifres10783 = False

								}

								var ifres10782 Obj

								if True == ifres10783 {
									ifres10782 = True

								} else {
									ifres10782 = False

								}

								ifres10781 = ifres10782

							} else {
								ifres10781 = False

							}

							if True == ifres10781 {
								tmp10690 := PrimTail(V2233)

								tmp10691 := PrimTail(V2234)

								__e.TailApply(PrimFunc(symshen_4compile_1head), V2232, tmp10690, tmp10691, V2235, V2236)
								return

							} else {
								tmp10779 := PrimIsPair(V2233)

								var ifres10775 Obj

								if True == tmp10779 {
									tmp10777 := PrimHead(V2233)

									tmp10778 := PrimIsVariable(tmp10777)

									var ifres10776 Obj

									if True == tmp10778 {
										ifres10776 = True

									} else {
										ifres10776 = False

									}

									ifres10775 = ifres10776

								} else {
									ifres10775 = False

								}

								if True == ifres10775 {
									__e.TailApply(PrimFunc(symshen_4variable_1case), V2232, V2233, V2234, V2235, V2236)
									return
								} else {
									tmp10773 := PrimEqual(symshen_4_1m, V2232)

									var ifres10766 Obj

									if True == tmp10773 {
										tmp10772 := PrimIsPair(V2233)

										var ifres10768 Obj

										if True == tmp10772 {
											tmp10770 := PrimHead(V2233)

											tmp10771 := Call(__e, PrimFunc(symatom_2), tmp10770)

											var ifres10769 Obj

											if True == tmp10771 {
												ifres10769 = True

											} else {
												ifres10769 = False

											}

											ifres10768 = ifres10769

										} else {
											ifres10768 = False

										}

										var ifres10767 Obj

										if True == ifres10768 {
											ifres10767 = True

										} else {
											ifres10767 = False

										}

										ifres10766 = ifres10767

									} else {
										ifres10766 = False

									}

									if True == ifres10766 {
										__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V2233, V2234, V2235, V2236)
										return
									} else {
										tmp10764 := PrimEqual(symshen_4_1m, V2232)

										var ifres10734 Obj

										if True == tmp10764 {
											tmp10763 := PrimIsPair(V2233)

											var ifres10736 Obj

											if True == tmp10763 {
												tmp10761 := PrimHead(V2233)

												tmp10762 := PrimIsPair(tmp10761)

												var ifres10738 Obj

												if True == tmp10762 {
													tmp10758 := PrimHead(V2233)

													tmp10759 := PrimHead(tmp10758)

													tmp10760 := PrimEqual(symcons, tmp10759)

													var ifres10740 Obj

													if True == tmp10760 {
														tmp10755 := PrimHead(V2233)

														tmp10756 := PrimTail(tmp10755)

														tmp10757 := PrimIsPair(tmp10756)

														var ifres10742 Obj

														if True == tmp10757 {
															tmp10751 := PrimHead(V2233)

															tmp10752 := PrimTail(tmp10751)

															tmp10753 := PrimTail(tmp10752)

															tmp10754 := PrimIsPair(tmp10753)

															var ifres10744 Obj

															if True == tmp10754 {
																tmp10746 := PrimHead(V2233)

																tmp10747 := PrimTail(tmp10746)

																tmp10748 := PrimTail(tmp10747)

																tmp10749 := PrimTail(tmp10748)

																tmp10750 := PrimEqual(Nil, tmp10749)

																var ifres10745 Obj

																if True == tmp10750 {
																	ifres10745 = True

																} else {
																	ifres10745 = False

																}

																ifres10744 = ifres10745

															} else {
																ifres10744 = False

															}

															var ifres10743 Obj

															if True == ifres10744 {
																ifres10743 = True

															} else {
																ifres10743 = False

															}

															ifres10742 = ifres10743

														} else {
															ifres10742 = False

														}

														var ifres10741 Obj

														if True == ifres10742 {
															ifres10741 = True

														} else {
															ifres10741 = False

														}

														ifres10740 = ifres10741

													} else {
														ifres10740 = False

													}

													var ifres10739 Obj

													if True == ifres10740 {
														ifres10739 = True

													} else {
														ifres10739 = False

													}

													ifres10738 = ifres10739

												} else {
													ifres10738 = False

												}

												var ifres10737 Obj

												if True == ifres10738 {
													ifres10737 = True

												} else {
													ifres10737 = False

												}

												ifres10736 = ifres10737

											} else {
												ifres10736 = False

											}

											var ifres10735 Obj

											if True == ifres10736 {
												ifres10735 = True

											} else {
												ifres10735 = False

											}

											ifres10734 = ifres10735

										} else {
											ifres10734 = False

										}

										if True == ifres10734 {
											__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V2233, V2234, V2235, V2236)
											return
										} else {
											tmp10732 := PrimEqual(symshen_4_7m, V2232)

											var ifres10725 Obj

											if True == tmp10732 {
												tmp10731 := PrimIsPair(V2233)

												var ifres10727 Obj

												if True == tmp10731 {
													tmp10729 := PrimHead(V2233)

													tmp10730 := Call(__e, PrimFunc(symatom_2), tmp10729)

													var ifres10728 Obj

													if True == tmp10730 {
														ifres10728 = True

													} else {
														ifres10728 = False

													}

													ifres10727 = ifres10728

												} else {
													ifres10727 = False

												}

												var ifres10726 Obj

												if True == ifres10727 {
													ifres10726 = True

												} else {
													ifres10726 = False

												}

												ifres10725 = ifres10726

											} else {
												ifres10725 = False

											}

											if True == ifres10725 {
												__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V2233, V2234, V2235, V2236)
												return
											} else {
												tmp10723 := PrimEqual(symshen_4_7m, V2232)

												var ifres10693 Obj

												if True == tmp10723 {
													tmp10722 := PrimIsPair(V2233)

													var ifres10695 Obj

													if True == tmp10722 {
														tmp10720 := PrimHead(V2233)

														tmp10721 := PrimIsPair(tmp10720)

														var ifres10697 Obj

														if True == tmp10721 {
															tmp10717 := PrimHead(V2233)

															tmp10718 := PrimHead(tmp10717)

															tmp10719 := PrimEqual(symcons, tmp10718)

															var ifres10699 Obj

															if True == tmp10719 {
																tmp10714 := PrimHead(V2233)

																tmp10715 := PrimTail(tmp10714)

																tmp10716 := PrimIsPair(tmp10715)

																var ifres10701 Obj

																if True == tmp10716 {
																	tmp10710 := PrimHead(V2233)

																	tmp10711 := PrimTail(tmp10710)

																	tmp10712 := PrimTail(tmp10711)

																	tmp10713 := PrimIsPair(tmp10712)

																	var ifres10703 Obj

																	if True == tmp10713 {
																		tmp10705 := PrimHead(V2233)

																		tmp10706 := PrimTail(tmp10705)

																		tmp10707 := PrimTail(tmp10706)

																		tmp10708 := PrimTail(tmp10707)

																		tmp10709 := PrimEqual(Nil, tmp10708)

																		var ifres10704 Obj

																		if True == tmp10709 {
																			ifres10704 = True

																		} else {
																			ifres10704 = False

																		}

																		ifres10703 = ifres10704

																	} else {
																		ifres10703 = False

																	}

																	var ifres10702 Obj

																	if True == ifres10703 {
																		ifres10702 = True

																	} else {
																		ifres10702 = False

																	}

																	ifres10701 = ifres10702

																} else {
																	ifres10701 = False

																}

																var ifres10700 Obj

																if True == ifres10701 {
																	ifres10700 = True

																} else {
																	ifres10700 = False

																}

																ifres10699 = ifres10700

															} else {
																ifres10699 = False

															}

															var ifres10698 Obj

															if True == ifres10699 {
																ifres10698 = True

															} else {
																ifres10698 = False

															}

															ifres10697 = ifres10698

														} else {
															ifres10697 = False

														}

														var ifres10696 Obj

														if True == ifres10697 {
															ifres10696 = True

														} else {
															ifres10696 = False

														}

														ifres10695 = ifres10696

													} else {
														ifres10695 = False

													}

													var ifres10694 Obj

													if True == ifres10695 {
														ifres10694 = True

													} else {
														ifres10694 = False

													}

													ifres10693 = ifres10694

												} else {
													ifres10693 = False

												}

												if True == ifres10693 {
													__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V2233, V2234, V2235, V2236)
													return
												} else {
													__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-head")))
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

		}

	}, 5)

	tmp10850 := Call(__e, ns2_1set, symshen_4compile_1head, tmp10673)

	_ = tmp10850

	tmp10851 := MakeNative(func(__e *ControlFlow) {
		V2247 := __e.Get(1)
		_ = V2247
		V2248 := __e.Get(2)
		_ = V2248
		V2249 := __e.Get(3)
		_ = V2249
		V2250 := __e.Get(4)
		_ = V2250
		V2251 := __e.Get(5)
		_ = V2251
		tmp10872 := PrimIsPair(V2248)

		var ifres10869 Obj

		if True == tmp10872 {
			tmp10871 := PrimIsPair(V2249)

			var ifres10870 Obj

			if True == tmp10871 {
				ifres10870 = True

			} else {
				ifres10870 = False

			}

			ifres10869 = ifres10870

		} else {
			ifres10869 = False

		}

		if True == ifres10869 {
			tmp10866 := PrimHead(V2249)

			tmp10867 := PrimIsVariable(tmp10866)

			if True == tmp10867 {
				tmp10852 := PrimTail(V2248)

				tmp10853 := PrimTail(V2249)

				tmp10854 := PrimHead(V2249)

				tmp10855 := PrimHead(V2248)

				tmp10856 := Call(__e, PrimFunc(symsubst), tmp10854, tmp10855, V2251)

				__e.TailApply(PrimFunc(symshen_4compile_1head), V2247, tmp10852, tmp10853, V2250, tmp10856)
				return

			} else {
				tmp10857 := PrimHead(V2248)

				tmp10858 := PrimHead(V2249)

				tmp10859 := PrimTail(V2248)

				tmp10860 := PrimTail(V2249)

				tmp10861 := Call(__e, PrimFunc(symshen_4compile_1head), V2247, tmp10859, tmp10860, V2250, V2251)

				tmp10862 := PrimCons(tmp10861, Nil)

				tmp10863 := PrimCons(tmp10858, tmp10862)

				tmp10864 := PrimCons(tmp10857, tmp10863)

				__e.Return(PrimCons(symlet, tmp10864))
				return

			}

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
			return
		}

	}, 5)

	tmp10873 := Call(__e, ns2_1set, symshen_4variable_1case, tmp10851)

	_ = tmp10873

	tmp10874 := MakeNative(func(__e *ControlFlow) {
		V2260 := __e.Get(1)
		_ = V2260
		V2261 := __e.Get(2)
		_ = V2261
		V2262 := __e.Get(3)
		_ = V2262
		V2263 := __e.Get(4)
		_ = V2263
		tmp10900 := PrimIsPair(V2260)

		var ifres10897 Obj

		if True == tmp10900 {
			tmp10899 := PrimIsPair(V2261)

			var ifres10898 Obj

			if True == tmp10899 {
				ifres10898 = True

			} else {
				ifres10898 = False

			}

			ifres10897 = ifres10898

		} else {
			ifres10897 = False

		}

		if True == ifres10897 {
			tmp10875 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp10876 := PrimHead(V2261)

				tmp10877 := PrimCons(V2262, Nil)

				tmp10878 := PrimCons(tmp10876, tmp10877)

				tmp10879 := PrimCons(symshen_4lazyderef, tmp10878)

				tmp10880 := PrimHead(V2260)

				tmp10881 := PrimCons(tmp10880, Nil)

				tmp10882 := PrimCons(Tm, tmp10881)

				tmp10883 := PrimCons(sym_a, tmp10882)

				tmp10884 := PrimTail(V2260)

				tmp10885 := PrimTail(V2261)

				tmp10886 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp10884, tmp10885, V2262, V2263)

				tmp10887 := PrimCons(False, Nil)

				tmp10888 := PrimCons(tmp10886, tmp10887)

				tmp10889 := PrimCons(tmp10883, tmp10888)

				tmp10890 := PrimCons(symif, tmp10889)

				tmp10891 := PrimCons(tmp10890, Nil)

				tmp10892 := PrimCons(tmp10879, tmp10891)

				tmp10893 := PrimCons(Tm, tmp10892)

				__e.Return(PrimCons(symlet, tmp10893))
				return

			}, 1)

			tmp10894 := Call(__e, PrimFunc(symprotect), symTm)

			tmp10895 := Call(__e, PrimFunc(symgensym), tmp10894)

			__e.TailApply(tmp10875, tmp10895)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
			return
		}

	}, 4)

	tmp10901 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp10874)

	_ = tmp10901

	tmp10902 := MakeNative(func(__e *ControlFlow) {
		V2272 := __e.Get(1)
		_ = V2272
		V2273 := __e.Get(2)
		_ = V2273
		V2274 := __e.Get(3)
		_ = V2274
		V2275 := __e.Get(4)
		_ = V2275
		tmp10968 := PrimIsPair(V2272)

		var ifres10938 Obj

		if True == tmp10968 {
			tmp10966 := PrimHead(V2272)

			tmp10967 := PrimIsPair(tmp10966)

			var ifres10940 Obj

			if True == tmp10967 {
				tmp10963 := PrimHead(V2272)

				tmp10964 := PrimHead(tmp10963)

				tmp10965 := PrimEqual(symcons, tmp10964)

				var ifres10942 Obj

				if True == tmp10965 {
					tmp10960 := PrimHead(V2272)

					tmp10961 := PrimTail(tmp10960)

					tmp10962 := PrimIsPair(tmp10961)

					var ifres10944 Obj

					if True == tmp10962 {
						tmp10956 := PrimHead(V2272)

						tmp10957 := PrimTail(tmp10956)

						tmp10958 := PrimTail(tmp10957)

						tmp10959 := PrimIsPair(tmp10958)

						var ifres10946 Obj

						if True == tmp10959 {
							tmp10951 := PrimHead(V2272)

							tmp10952 := PrimTail(tmp10951)

							tmp10953 := PrimTail(tmp10952)

							tmp10954 := PrimTail(tmp10953)

							tmp10955 := PrimEqual(Nil, tmp10954)

							var ifres10948 Obj

							if True == tmp10955 {
								tmp10950 := PrimIsPair(V2273)

								var ifres10949 Obj

								if True == tmp10950 {
									ifres10949 = True

								} else {
									ifres10949 = False

								}

								ifres10948 = ifres10949

							} else {
								ifres10948 = False

							}

							var ifres10947 Obj

							if True == ifres10948 {
								ifres10947 = True

							} else {
								ifres10947 = False

							}

							ifres10946 = ifres10947

						} else {
							ifres10946 = False

						}

						var ifres10945 Obj

						if True == ifres10946 {
							ifres10945 = True

						} else {
							ifres10945 = False

						}

						ifres10944 = ifres10945

					} else {
						ifres10944 = False

					}

					var ifres10943 Obj

					if True == ifres10944 {
						ifres10943 = True

					} else {
						ifres10943 = False

					}

					ifres10942 = ifres10943

				} else {
					ifres10942 = False

				}

				var ifres10941 Obj

				if True == ifres10942 {
					ifres10941 = True

				} else {
					ifres10941 = False

				}

				ifres10940 = ifres10941

			} else {
				ifres10940 = False

			}

			var ifres10939 Obj

			if True == ifres10940 {
				ifres10939 = True

			} else {
				ifres10939 = False

			}

			ifres10938 = ifres10939

		} else {
			ifres10938 = False

		}

		if True == ifres10938 {
			tmp10903 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp10904 := PrimHead(V2273)

				tmp10905 := PrimCons(V2274, Nil)

				tmp10906 := PrimCons(tmp10904, tmp10905)

				tmp10907 := PrimCons(symshen_4lazyderef, tmp10906)

				tmp10908 := PrimCons(Tm, Nil)

				tmp10909 := PrimCons(symcons_2, tmp10908)

				tmp10910 := PrimHead(V2272)

				tmp10911 := PrimTail(tmp10910)

				tmp10912 := PrimHead(tmp10911)

				tmp10913 := PrimHead(V2272)

				tmp10914 := PrimTail(tmp10913)

				tmp10915 := PrimTail(tmp10914)

				tmp10916 := PrimHead(tmp10915)

				tmp10917 := PrimTail(V2272)

				tmp10918 := PrimCons(tmp10916, tmp10917)

				tmp10919 := PrimCons(tmp10912, tmp10918)

				tmp10920 := PrimCons(Tm, Nil)

				tmp10921 := PrimCons(symhd, tmp10920)

				tmp10922 := PrimCons(Tm, Nil)

				tmp10923 := PrimCons(symtl, tmp10922)

				tmp10924 := PrimTail(V2273)

				tmp10925 := PrimCons(tmp10923, tmp10924)

				tmp10926 := PrimCons(tmp10921, tmp10925)

				tmp10927 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp10919, tmp10926, V2274, V2275)

				tmp10928 := PrimCons(False, Nil)

				tmp10929 := PrimCons(tmp10927, tmp10928)

				tmp10930 := PrimCons(tmp10909, tmp10929)

				tmp10931 := PrimCons(symif, tmp10930)

				tmp10932 := PrimCons(tmp10931, Nil)

				tmp10933 := PrimCons(tmp10907, tmp10932)

				tmp10934 := PrimCons(Tm, tmp10933)

				__e.Return(PrimCons(symlet, tmp10934))
				return

			}, 1)

			tmp10935 := Call(__e, PrimFunc(symprotect), symTm)

			tmp10936 := Call(__e, PrimFunc(symgensym), tmp10935)

			__e.TailApply(tmp10903, tmp10936)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
			return
		}

	}, 4)

	tmp10969 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp10902)

	_ = tmp10969

	tmp10970 := MakeNative(func(__e *ControlFlow) {
		V2284 := __e.Get(1)
		_ = V2284
		V2285 := __e.Get(2)
		_ = V2285
		V2286 := __e.Get(3)
		_ = V2286
		V2287 := __e.Get(4)
		_ = V2287
		tmp11018 := PrimIsPair(V2284)

		var ifres11015 Obj

		if True == tmp11018 {
			tmp11017 := PrimIsPair(V2285)

			var ifres11016 Obj

			if True == tmp11017 {
				ifres11016 = True

			} else {
				ifres11016 = False

			}

			ifres11015 = ifres11016

		} else {
			ifres11015 = False

		}

		if True == ifres11015 {
			tmp10971 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp10972 := MakeNative(func(__e *ControlFlow) {
					GoTo := __e.Get(1)
					_ = GoTo
					tmp10973 := PrimHead(V2285)

					tmp10974 := PrimCons(V2286, Nil)

					tmp10975 := PrimCons(tmp10973, tmp10974)

					tmp10976 := PrimCons(symshen_4lazyderef, tmp10975)

					tmp10977 := PrimTail(V2284)

					tmp10978 := PrimTail(V2285)

					tmp10979 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10977, tmp10978, V2286, V2287)

					tmp10980 := PrimCons(tmp10979, Nil)

					tmp10981 := PrimCons(symfreeze, tmp10980)

					tmp10982 := PrimHead(V2284)

					tmp10983 := PrimCons(tmp10982, Nil)

					tmp10984 := PrimCons(Tm, tmp10983)

					tmp10985 := PrimCons(sym_a, tmp10984)

					tmp10986 := PrimCons(GoTo, Nil)

					tmp10987 := PrimCons(symthaw, tmp10986)

					tmp10988 := PrimCons(Tm, Nil)

					tmp10989 := PrimCons(symshen_4pvar_2, tmp10988)

					tmp10990 := PrimHead(V2284)

					tmp10991 := Call(__e, PrimFunc(symshen_4demode), tmp10990)

					tmp10992 := PrimCons(GoTo, Nil)

					tmp10993 := PrimCons(V2286, tmp10992)

					tmp10994 := PrimCons(tmp10991, tmp10993)

					tmp10995 := PrimCons(Tm, tmp10994)

					tmp10996 := PrimCons(symshen_4bind_b, tmp10995)

					tmp10997 := PrimCons(False, Nil)

					tmp10998 := PrimCons(tmp10996, tmp10997)

					tmp10999 := PrimCons(tmp10989, tmp10998)

					tmp11000 := PrimCons(symif, tmp10999)

					tmp11001 := PrimCons(tmp11000, Nil)

					tmp11002 := PrimCons(tmp10987, tmp11001)

					tmp11003 := PrimCons(tmp10985, tmp11002)

					tmp11004 := PrimCons(symif, tmp11003)

					tmp11005 := PrimCons(tmp11004, Nil)

					tmp11006 := PrimCons(tmp10981, tmp11005)

					tmp11007 := PrimCons(GoTo, tmp11006)

					tmp11008 := PrimCons(tmp10976, tmp11007)

					tmp11009 := PrimCons(Tm, tmp11008)

					__e.Return(PrimCons(symlet, tmp11009))
					return

				}, 1)

				tmp11010 := Call(__e, PrimFunc(symprotect), symGoTo)

				tmp11011 := Call(__e, PrimFunc(symgensym), tmp11010)

				__e.TailApply(tmp10972, tmp11011)
				return

			}, 1)

			tmp11012 := Call(__e, PrimFunc(symprotect), symTm)

			tmp11013 := Call(__e, PrimFunc(symgensym), tmp11012)

			__e.TailApply(tmp10971, tmp11013)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
			return
		}

	}, 4)

	tmp11019 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp10970)

	_ = tmp11019

	tmp11020 := MakeNative(func(__e *ControlFlow) {
		V2296 := __e.Get(1)
		_ = V2296
		V2297 := __e.Get(2)
		_ = V2297
		V2298 := __e.Get(3)
		_ = V2298
		V2299 := __e.Get(4)
		_ = V2299
		tmp11118 := PrimIsPair(V2296)

		var ifres11088 Obj

		if True == tmp11118 {
			tmp11116 := PrimHead(V2296)

			tmp11117 := PrimIsPair(tmp11116)

			var ifres11090 Obj

			if True == tmp11117 {
				tmp11113 := PrimHead(V2296)

				tmp11114 := PrimHead(tmp11113)

				tmp11115 := PrimEqual(symcons, tmp11114)

				var ifres11092 Obj

				if True == tmp11115 {
					tmp11110 := PrimHead(V2296)

					tmp11111 := PrimTail(tmp11110)

					tmp11112 := PrimIsPair(tmp11111)

					var ifres11094 Obj

					if True == tmp11112 {
						tmp11106 := PrimHead(V2296)

						tmp11107 := PrimTail(tmp11106)

						tmp11108 := PrimTail(tmp11107)

						tmp11109 := PrimIsPair(tmp11108)

						var ifres11096 Obj

						if True == tmp11109 {
							tmp11101 := PrimHead(V2296)

							tmp11102 := PrimTail(tmp11101)

							tmp11103 := PrimTail(tmp11102)

							tmp11104 := PrimTail(tmp11103)

							tmp11105 := PrimEqual(Nil, tmp11104)

							var ifres11098 Obj

							if True == tmp11105 {
								tmp11100 := PrimIsPair(V2297)

								var ifres11099 Obj

								if True == tmp11100 {
									ifres11099 = True

								} else {
									ifres11099 = False

								}

								ifres11098 = ifres11099

							} else {
								ifres11098 = False

							}

							var ifres11097 Obj

							if True == ifres11098 {
								ifres11097 = True

							} else {
								ifres11097 = False

							}

							ifres11096 = ifres11097

						} else {
							ifres11096 = False

						}

						var ifres11095 Obj

						if True == ifres11096 {
							ifres11095 = True

						} else {
							ifres11095 = False

						}

						ifres11094 = ifres11095

					} else {
						ifres11094 = False

					}

					var ifres11093 Obj

					if True == ifres11094 {
						ifres11093 = True

					} else {
						ifres11093 = False

					}

					ifres11092 = ifres11093

				} else {
					ifres11092 = False

				}

				var ifres11091 Obj

				if True == ifres11092 {
					ifres11091 = True

				} else {
					ifres11091 = False

				}

				ifres11090 = ifres11091

			} else {
				ifres11090 = False

			}

			var ifres11089 Obj

			if True == ifres11090 {
				ifres11089 = True

			} else {
				ifres11089 = False

			}

			ifres11088 = ifres11089

		} else {
			ifres11088 = False

		}

		if True == ifres11088 {
			tmp11021 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp11022 := MakeNative(func(__e *ControlFlow) {
					GoTo := __e.Get(1)
					_ = GoTo
					tmp11023 := MakeNative(func(__e *ControlFlow) {
						Vars := __e.Get(1)
						_ = Vars
						tmp11024 := MakeNative(func(__e *ControlFlow) {
							Tame := __e.Get(1)
							_ = Tame
							tmp11025 := MakeNative(func(__e *ControlFlow) {
								TVars := __e.Get(1)
								_ = TVars
								tmp11026 := PrimHead(V2297)

								tmp11027 := PrimCons(V2298, Nil)

								tmp11028 := PrimCons(tmp11026, tmp11027)

								tmp11029 := PrimCons(symshen_4lazyderef, tmp11028)

								tmp11030 := PrimTail(V2296)

								tmp11031 := PrimTail(V2297)

								tmp11032 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11030, tmp11031, V2298, V2299)

								tmp11033 := Call(__e, PrimFunc(symshen_4goto), Vars, tmp11032)

								tmp11034 := PrimCons(Tm, Nil)

								tmp11035 := PrimCons(symcons_2, tmp11034)

								tmp11036 := PrimHead(V2296)

								tmp11037 := PrimTail(tmp11036)

								tmp11038 := PrimCons(Tm, Nil)

								tmp11039 := PrimCons(symhd, tmp11038)

								tmp11040 := PrimCons(Tm, Nil)

								tmp11041 := PrimCons(symtl, tmp11040)

								tmp11042 := PrimCons(tmp11041, Nil)

								tmp11043 := PrimCons(tmp11039, tmp11042)

								tmp11044 := Call(__e, PrimFunc(symshen_4invoke), GoTo, Vars)

								tmp11045 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11037, tmp11043, V2298, tmp11044)

								tmp11046 := PrimCons(Tm, Nil)

								tmp11047 := PrimCons(symshen_4pvar_2, tmp11046)

								tmp11048 := Call(__e, PrimFunc(symshen_4demode), Tame)

								tmp11049 := Call(__e, PrimFunc(symshen_4invoke), GoTo, Vars)

								tmp11050 := PrimCons(tmp11049, Nil)

								tmp11051 := PrimCons(symfreeze, tmp11050)

								tmp11052 := PrimCons(tmp11051, Nil)

								tmp11053 := PrimCons(V2298, tmp11052)

								tmp11054 := PrimCons(tmp11048, tmp11053)

								tmp11055 := PrimCons(Tm, tmp11054)

								tmp11056 := PrimCons(symshen_4bind_b, tmp11055)

								tmp11057 := Call(__e, PrimFunc(symshen_4stpart), TVars, tmp11056, V2298)

								tmp11058 := PrimCons(False, Nil)

								tmp11059 := PrimCons(tmp11057, tmp11058)

								tmp11060 := PrimCons(tmp11047, tmp11059)

								tmp11061 := PrimCons(symif, tmp11060)

								tmp11062 := PrimCons(tmp11061, Nil)

								tmp11063 := PrimCons(tmp11045, tmp11062)

								tmp11064 := PrimCons(tmp11035, tmp11063)

								tmp11065 := PrimCons(symif, tmp11064)

								tmp11066 := PrimCons(tmp11065, Nil)

								tmp11067 := PrimCons(tmp11033, tmp11066)

								tmp11068 := PrimCons(GoTo, tmp11067)

								tmp11069 := PrimCons(tmp11029, tmp11068)

								tmp11070 := PrimCons(Tm, tmp11069)

								__e.Return(PrimCons(symlet, tmp11070))
								return

							}, 1)

							tmp11071 := Call(__e, PrimFunc(symshen_4extract_1vars), Tame)

							__e.TailApply(tmp11025, tmp11071)
							return

						}, 1)

						tmp11072 := PrimHead(V2296)

						tmp11073 := Call(__e, PrimFunc(symshen_4tame), tmp11072)

						__e.TailApply(tmp11024, tmp11073)
						return

					}, 1)

					tmp11074 := PrimHead(V2296)

					tmp11075 := PrimTail(tmp11074)

					tmp11076 := PrimHead(tmp11075)

					tmp11077 := PrimHead(V2296)

					tmp11078 := PrimTail(tmp11077)

					tmp11079 := PrimTail(tmp11078)

					tmp11080 := PrimHead(tmp11079)

					tmp11081 := PrimCons(tmp11076, tmp11080)

					tmp11082 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11081)

					__e.TailApply(tmp11023, tmp11082)
					return

				}, 1)

				tmp11083 := Call(__e, PrimFunc(symprotect), symGoTo)

				tmp11084 := Call(__e, PrimFunc(symgensym), tmp11083)

				__e.TailApply(tmp11022, tmp11084)
				return

			}, 1)

			tmp11085 := Call(__e, PrimFunc(symprotect), symTm)

			tmp11086 := Call(__e, PrimFunc(symgensym), tmp11085)

			__e.TailApply(tmp11021, tmp11086)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
			return
		}

	}, 4)

	tmp11119 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11020)

	_ = tmp11119

	tmp11120 := MakeNative(func(__e *ControlFlow) {
		V2300 := __e.Get(1)
		_ = V2300
		tmp11157 := PrimIsPair(V2300)

		var ifres11144 Obj

		if True == tmp11157 {
			tmp11155 := PrimHead(V2300)

			tmp11156 := PrimEqual(symshen_4_7m, tmp11155)

			var ifres11146 Obj

			if True == tmp11156 {
				tmp11153 := PrimTail(V2300)

				tmp11154 := PrimIsPair(tmp11153)

				var ifres11148 Obj

				if True == tmp11154 {
					tmp11150 := PrimTail(V2300)

					tmp11151 := PrimTail(tmp11150)

					tmp11152 := PrimEqual(Nil, tmp11151)

					var ifres11149 Obj

					if True == tmp11152 {
						ifres11149 = True

					} else {
						ifres11149 = False

					}

					ifres11148 = ifres11149

				} else {
					ifres11148 = False

				}

				var ifres11147 Obj

				if True == ifres11148 {
					ifres11147 = True

				} else {
					ifres11147 = False

				}

				ifres11146 = ifres11147

			} else {
				ifres11146 = False

			}

			var ifres11145 Obj

			if True == ifres11146 {
				ifres11145 = True

			} else {
				ifres11145 = False

			}

			ifres11144 = ifres11145

		} else {
			ifres11144 = False

		}

		if True == ifres11144 {
			tmp11121 := PrimTail(V2300)

			tmp11122 := PrimHead(tmp11121)

			__e.TailApply(PrimFunc(symshen_4demode), tmp11122)
			return

		} else {
			tmp11142 := PrimIsPair(V2300)

			var ifres11129 Obj

			if True == tmp11142 {
				tmp11140 := PrimHead(V2300)

				tmp11141 := PrimEqual(symshen_4_1m, tmp11140)

				var ifres11131 Obj

				if True == tmp11141 {
					tmp11138 := PrimTail(V2300)

					tmp11139 := PrimIsPair(tmp11138)

					var ifres11133 Obj

					if True == tmp11139 {
						tmp11135 := PrimTail(V2300)

						tmp11136 := PrimTail(tmp11135)

						tmp11137 := PrimEqual(Nil, tmp11136)

						var ifres11134 Obj

						if True == tmp11137 {
							ifres11134 = True

						} else {
							ifres11134 = False

						}

						ifres11133 = ifres11134

					} else {
						ifres11133 = False

					}

					var ifres11132 Obj

					if True == ifres11133 {
						ifres11132 = True

					} else {
						ifres11132 = False

					}

					ifres11131 = ifres11132

				} else {
					ifres11131 = False

				}

				var ifres11130 Obj

				if True == ifres11131 {
					ifres11130 = True

				} else {
					ifres11130 = False

				}

				ifres11129 = ifres11130

			} else {
				ifres11129 = False

			}

			if True == ifres11129 {
				tmp11123 := PrimTail(V2300)

				tmp11124 := PrimHead(tmp11123)

				__e.TailApply(PrimFunc(symshen_4demode), tmp11124)
				return

			} else {
				tmp11127 := PrimIsPair(V2300)

				if True == tmp11127 {
					tmp11125 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(PrimFunc(symshen_4demode), Z)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp11125, V2300)
					return

				} else {
					__e.Return(V2300)
					return
				}

			}

		}

	}, 1)

	tmp11158 := Call(__e, ns2_1set, symshen_4demode, tmp11120)

	_ = tmp11158

	tmp11159 := MakeNative(func(__e *ControlFlow) {
		V2301 := __e.Get(1)
		_ = V2301
		tmp11165 := Call(__e, PrimFunc(symshen_4wildcard_2), V2301)

		if True == tmp11165 {
			tmp11160 := Call(__e, PrimFunc(symprotect), symY)

			__e.TailApply(PrimFunc(symgensym), tmp11160)
			return

		} else {
			tmp11163 := PrimIsPair(V2301)

			if True == tmp11163 {
				tmp11161 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimFunc(symshen_4tame), Z)
					return
				}, 1)

				__e.TailApply(PrimFunc(symmap), tmp11161, V2301)
				return

			} else {
				__e.Return(V2301)
				return
			}

		}

	}, 1)

	tmp11166 := Call(__e, ns2_1set, symshen_4tame, tmp11159)

	_ = tmp11166

	tmp11167 := MakeNative(func(__e *ControlFlow) {
		V2302 := __e.Get(1)
		_ = V2302
		V2303 := __e.Get(2)
		_ = V2303
		tmp11170 := PrimEqual(Nil, V2302)

		if True == tmp11170 {
			tmp11168 := PrimCons(V2303, Nil)

			__e.Return(PrimCons(symfreeze, tmp11168))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4goto_1h), V2302, V2303)
			return
		}

	}, 2)

	tmp11171 := Call(__e, ns2_1set, symshen_4goto, tmp11167)

	_ = tmp11171

	tmp11172 := MakeNative(func(__e *ControlFlow) {
		V2304 := __e.Get(1)
		_ = V2304
		V2305 := __e.Get(2)
		_ = V2305
		tmp11181 := PrimEqual(Nil, V2304)

		if True == tmp11181 {
			__e.Return(V2305)
			return
		} else {
			tmp11179 := PrimIsPair(V2304)

			if True == tmp11179 {
				tmp11173 := PrimHead(V2304)

				tmp11174 := PrimTail(V2304)

				tmp11175 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11174, V2305)

				tmp11176 := PrimCons(tmp11175, Nil)

				tmp11177 := PrimCons(tmp11173, tmp11176)

				__e.Return(PrimCons(symlambda, tmp11177))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4goto_1h)
				return
			}

		}

	}, 2)

	tmp11182 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11172)

	_ = tmp11182

	tmp11183 := MakeNative(func(__e *ControlFlow) {
		V2306 := __e.Get(1)
		_ = V2306
		V2307 := __e.Get(2)
		_ = V2307
		tmp11186 := PrimEqual(Nil, V2307)

		if True == tmp11186 {
			tmp11184 := PrimCons(V2306, Nil)

			__e.Return(PrimCons(symthaw, tmp11184))
			return

		} else {
			__e.Return(PrimCons(V2306, V2307))
			return
		}

	}, 2)

	tmp11187 := Call(__e, ns2_1set, symshen_4invoke, tmp11183)

	_ = tmp11187

	tmp11188 := MakeNative(func(__e *ControlFlow) {
		V2308 := __e.Get(1)
		_ = V2308
		__e.Return(PrimEqual(V2308, sym__))
		return
	}, 1)

	tmp11189 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11188)

	_ = tmp11189

	tmp11190 := MakeNative(func(__e *ControlFlow) {
		V2309 := __e.Get(1)
		_ = V2309
		tmp11191 := MakeNative(func(__e *ControlFlow) {
			tmp11196 := PrimIsVector(V2309)

			if True == tmp11196 {
				tmp11193 := PrimVectorGet(V2309, MakeNumber(0))

				tmp11194 := PrimEqual(tmp11193, symshen_4pvar)

				if True == tmp11194 {
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

		}, 0)

		tmp11197 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(try_1catch, tmp11191, tmp11197)
		return

	}, 1)

	tmp11198 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11190)

	_ = tmp11198

	tmp11199 := MakeNative(func(__e *ControlFlow) {
		V2310 := __e.Get(1)
		_ = V2310
		V2311 := __e.Get(2)
		_ = V2311
		tmp11206 := Call(__e, PrimFunc(symshen_4pvar_2), V2310)

		if True == tmp11206 {
			tmp11200 := MakeNative(func(__e *ControlFlow) {
				Value := __e.Get(1)
				_ = Value
				tmp11202 := PrimEqual(Value, symshen_4_1null_1)

				if True == tmp11202 {
					__e.Return(V2310)
					return
				} else {
					__e.TailApply(PrimFunc(symshen_4lazyderef), Value, V2311)
					return
				}

			}, 1)

			tmp11203 := PrimVectorGet(V2310, MakeNumber(1))

			tmp11204 := PrimVectorGet(V2311, tmp11203)

			__e.TailApply(tmp11200, tmp11204)
			return

		} else {
			__e.Return(V2310)
			return
		}

	}, 2)

	tmp11207 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11199)

	_ = tmp11207

	tmp11208 := MakeNative(func(__e *ControlFlow) {
		V2312 := __e.Get(1)
		_ = V2312
		V2313 := __e.Get(2)
		_ = V2313
		tmp11221 := PrimIsPair(V2312)

		if True == tmp11221 {
			tmp11209 := PrimHead(V2312)

			tmp11210 := Call(__e, PrimFunc(symshen_4deref), tmp11209, V2313)

			tmp11211 := PrimTail(V2312)

			tmp11212 := Call(__e, PrimFunc(symshen_4deref), tmp11211, V2313)

			__e.Return(PrimCons(tmp11210, tmp11212))
			return

		} else {
			tmp11219 := Call(__e, PrimFunc(symshen_4pvar_2), V2312)

			if True == tmp11219 {
				tmp11213 := MakeNative(func(__e *ControlFlow) {
					Value := __e.Get(1)
					_ = Value
					tmp11215 := PrimEqual(Value, symshen_4_1null_1)

					if True == tmp11215 {
						__e.Return(V2312)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4deref), Value, V2313)
						return
					}

				}, 1)

				tmp11216 := PrimVectorGet(V2312, MakeNumber(1))

				tmp11217 := PrimVectorGet(V2313, tmp11216)

				__e.TailApply(tmp11213, tmp11217)
				return

			} else {
				__e.Return(V2312)
				return
			}

		}

	}, 2)

	tmp11222 := Call(__e, ns2_1set, symshen_4deref, tmp11208)

	_ = tmp11222

	tmp11223 := MakeNative(func(__e *ControlFlow) {
		V2314 := __e.Get(1)
		_ = V2314
		V2315 := __e.Get(2)
		_ = V2315
		V2316 := __e.Get(3)
		_ = V2316
		V2317 := __e.Get(4)
		_ = V2317
		tmp11224 := MakeNative(func(__e *ControlFlow) {
			Bind := __e.Get(1)
			_ = Bind
			tmp11225 := MakeNative(func(__e *ControlFlow) {
				Compute := __e.Get(1)
				_ = Compute
				tmp11227 := PrimEqual(Compute, False)

				if True == tmp11227 {
					__e.TailApply(PrimFunc(symshen_4unwind), V2314, V2316, Compute)
					return
				} else {
					__e.Return(Compute)
					return
				}

			}, 1)

			tmp11228 := Call(__e, PrimFunc(symthaw), V2317)

			__e.TailApply(tmp11225, tmp11228)
			return

		}, 1)

		tmp11229 := Call(__e, PrimFunc(symshen_4bindv), V2314, V2315, V2316)

		__e.TailApply(tmp11224, tmp11229)
		return

	}, 4)

	tmp11230 := Call(__e, ns2_1set, symshen_4bind_b, tmp11223)

	_ = tmp11230

	tmp11231 := MakeNative(func(__e *ControlFlow) {
		V2318 := __e.Get(1)
		_ = V2318
		V2319 := __e.Get(2)
		_ = V2319
		V2320 := __e.Get(3)
		_ = V2320
		tmp11232 := PrimVectorGet(V2318, MakeNumber(1))

		__e.Return(PrimVectorSet(V2320, tmp11232, V2319))
		return

	}, 3)

	tmp11233 := Call(__e, ns2_1set, symshen_4bindv, tmp11231)

	_ = tmp11233

	tmp11234 := MakeNative(func(__e *ControlFlow) {
		V2321 := __e.Get(1)
		_ = V2321
		V2322 := __e.Get(2)
		_ = V2322
		V2323 := __e.Get(3)
		_ = V2323
		tmp11235 := PrimVectorGet(V2321, MakeNumber(1))

		tmp11236 := PrimVectorSet(V2322, tmp11235, symshen_4_1null_1)

		_ = tmp11236

		__e.Return(V2323)
		return

	}, 3)

	tmp11237 := Call(__e, ns2_1set, symshen_4unwind, tmp11234)

	_ = tmp11237

	tmp11238 := MakeNative(func(__e *ControlFlow) {
		V2332 := __e.Get(1)
		_ = V2332
		V2333 := __e.Get(2)
		_ = V2333
		V2334 := __e.Get(3)
		_ = V2334
		tmp11253 := PrimEqual(Nil, V2332)

		if True == tmp11253 {
			__e.Return(V2333)
			return
		} else {
			tmp11251 := PrimIsPair(V2332)

			if True == tmp11251 {
				tmp11239 := PrimHead(V2332)

				tmp11240 := PrimCons(V2334, Nil)

				tmp11241 := PrimCons(symshen_4newpv, tmp11240)

				tmp11242 := PrimTail(V2332)

				tmp11243 := Call(__e, PrimFunc(symshen_4stpart), tmp11242, V2333, V2334)

				tmp11244 := PrimCons(tmp11243, Nil)

				tmp11245 := PrimCons(V2334, tmp11244)

				tmp11246 := PrimCons(symshen_4gc, tmp11245)

				tmp11247 := PrimCons(tmp11246, Nil)

				tmp11248 := PrimCons(tmp11241, tmp11247)

				tmp11249 := PrimCons(tmp11239, tmp11248)

				__e.Return(PrimCons(symlet, tmp11249))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
				return
			}

		}

	}, 3)

	tmp11254 := Call(__e, ns2_1set, symshen_4stpart, tmp11238)

	_ = tmp11254

	tmp11255 := MakeNative(func(__e *ControlFlow) {
		V2335 := __e.Get(1)
		_ = V2335
		V2336 := __e.Get(2)
		_ = V2336
		tmp11260 := PrimEqual(V2336, False)

		if True == tmp11260 {
			tmp11256 := MakeNative(func(__e *ControlFlow) {
				N := __e.Get(1)
				_ = N
				tmp11257 := Call(__e, PrimFunc(symshen_4decrement_1ticket), N, V2335)

				_ = tmp11257

				__e.Return(V2336)
				return

			}, 1)

			tmp11258 := Call(__e, PrimFunc(symshen_4ticket_1number), V2335)

			__e.TailApply(tmp11256, tmp11258)
			return

		} else {
			__e.Return(V2336)
			return
		}

	}, 2)

	tmp11261 := Call(__e, ns2_1set, symshen_4gc, tmp11255)

	_ = tmp11261

	tmp11262 := MakeNative(func(__e *ControlFlow) {
		V2337 := __e.Get(1)
		_ = V2337
		V2338 := __e.Get(2)
		_ = V2338
		tmp11263 := PrimNumberSubtract(V2337, MakeNumber(1))

		__e.Return(PrimVectorSet(V2338, MakeNumber(1), tmp11263))
		return

	}, 2)

	tmp11264 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp11262)

	_ = tmp11264

	tmp11265 := MakeNative(func(__e *ControlFlow) {
		V2339 := __e.Get(1)
		_ = V2339
		tmp11266 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp11267 := MakeNative(func(__e *ControlFlow) {
				NewBindings := __e.Get(1)
				_ = NewBindings
				tmp11268 := MakeNative(func(__e *ControlFlow) {
					NextTicket := __e.Get(1)
					_ = NextTicket
					__e.Return(NewBindings)
					return
				}, 1)

				tmp11269 := Call(__e, PrimFunc(symshen_4nextticket), V2339, N)

				__e.TailApply(tmp11268, tmp11269)
				return

			}, 1)

			tmp11270 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), N)

			__e.TailApply(tmp11267, tmp11270)
			return

		}, 1)

		tmp11271 := Call(__e, PrimFunc(symshen_4ticket_1number), V2339)

		__e.TailApply(tmp11266, tmp11271)
		return

	}, 1)

	tmp11272 := Call(__e, ns2_1set, symshen_4newpv, tmp11265)

	_ = tmp11272

	tmp11273 := MakeNative(func(__e *ControlFlow) {
		V2340 := __e.Get(1)
		_ = V2340
		__e.Return(PrimVectorGet(V2340, MakeNumber(1)))
		return
	}, 1)

	tmp11274 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp11273)

	_ = tmp11274

	tmp11275 := MakeNative(func(__e *ControlFlow) {
		V2341 := __e.Get(1)
		_ = V2341
		V2342 := __e.Get(2)
		_ = V2342
		tmp11276 := MakeNative(func(__e *ControlFlow) {
			NewVector := __e.Get(1)
			_ = NewVector
			tmp11277 := PrimNumberAdd(V2342, MakeNumber(1))

			__e.Return(PrimVectorSet(NewVector, MakeNumber(1), tmp11277))
			return

		}, 1)

		tmp11278 := PrimVectorSet(V2341, V2342, symshen_4_1null_1)

		__e.TailApply(tmp11276, tmp11278)
		return

	}, 2)

	tmp11279 := Call(__e, ns2_1set, symshen_4nextticket, tmp11275)

	_ = tmp11279

	tmp11280 := MakeNative(func(__e *ControlFlow) {
		V2343 := __e.Get(1)
		_ = V2343
		tmp11281 := PrimAbsvector(MakeNumber(2))

		tmp11282 := PrimVectorSet(tmp11281, MakeNumber(0), symshen_4pvar)

		__e.Return(PrimVectorSet(tmp11282, MakeNumber(1), V2343))
		return

	}, 1)

	tmp11283 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11280)

	_ = tmp11283

	tmp11284 := MakeNative(func(__e *ControlFlow) {
		V2344 := __e.Get(1)
		_ = V2344
		tmp11285 := PrimVectorGet(V2344, MakeNumber(1))

		tmp11286 := Call(__e, PrimFunc(symshen_4app), tmp11285, MakeString(""), symshen_4a)

		__e.Return(PrimStringConcat(MakeString("Var"), tmp11286))
		return

	}, 1)

	tmp11287 := Call(__e, ns2_1set, symshen_4pvar, tmp11284)

	_ = tmp11287

	tmp11288 := MakeNative(func(__e *ControlFlow) {
		tmp11289 := PrimValue(symshen_4_dinfs_d)

		tmp11290 := PrimNumberAdd(MakeNumber(1), tmp11289)

		__e.Return(PrimSet(symshen_4_dinfs_d, tmp11290))
		return

	}, 0)

	tmp11291 := Call(__e, ns2_1set, symshen_4incinfs, tmp11288)

	_ = tmp11291

	tmp11292 := MakeNative(func(__e *ControlFlow) {
		V2345 := __e.Get(1)
		_ = V2345
		tmp11299 := PrimIsInteger(V2345)

		var ifres11296 Obj

		if True == tmp11299 {
			tmp11298 := PrimGreatThan(V2345, MakeNumber(0))

			var ifres11297 Obj

			if True == tmp11298 {
				ifres11297 = True

			} else {
				ifres11297 = False

			}

			ifres11296 = ifres11297

		} else {
			ifres11296 = False

		}

		if True == ifres11296 {
			__e.Return(PrimSet(symshen_4_dsize_1prolog_1vector_d, V2345))
			return
		} else {
			tmp11293 := Call(__e, PrimFunc(symshen_4app), V2345, MakeString(""), symshen_4a)

			tmp11294 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp11293)

			__e.Return(PrimSimpleError(tmp11294))
			return

		}

	}, 1)

	tmp11300 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11292)

	_ = tmp11300

	tmp11301 := MakeNative(func(__e *ControlFlow) {
		V2357 := __e.Get(1)
		_ = V2357
		V2358 := __e.Get(2)
		_ = V2358
		V2359 := __e.Get(3)
		_ = V2359
		V2360 := __e.Get(4)
		_ = V2360
		tmp11331 := PrimEqual(V2357, V2358)

		if True == tmp11331 {
			__e.TailApply(PrimFunc(symthaw), V2360)
			return
		} else {
			tmp11329 := Call(__e, PrimFunc(symshen_4pvar_2), V2357)

			var ifres11324 Obj

			if True == tmp11329 {
				tmp11326 := Call(__e, PrimFunc(symshen_4deref), V2358, V2359)

				tmp11327 := Call(__e, PrimFunc(symshen_4occurs_2), V2357, tmp11326)

				tmp11328 := PrimNot(tmp11327)

				var ifres11325 Obj

				if True == tmp11328 {
					ifres11325 = True

				} else {
					ifres11325 = False

				}

				ifres11324 = ifres11325

			} else {
				ifres11324 = False

			}

			if True == ifres11324 {
				__e.TailApply(PrimFunc(symshen_4bind_b), V2357, V2358, V2359, V2360)
				return
			} else {
				tmp11322 := Call(__e, PrimFunc(symshen_4pvar_2), V2358)

				var ifres11317 Obj

				if True == tmp11322 {
					tmp11319 := Call(__e, PrimFunc(symshen_4deref), V2357, V2359)

					tmp11320 := Call(__e, PrimFunc(symshen_4occurs_2), V2358, tmp11319)

					tmp11321 := PrimNot(tmp11320)

					var ifres11318 Obj

					if True == tmp11321 {
						ifres11318 = True

					} else {
						ifres11318 = False

					}

					ifres11317 = ifres11318

				} else {
					ifres11317 = False

				}

				if True == ifres11317 {
					__e.TailApply(PrimFunc(symshen_4bind_b), V2358, V2357, V2359, V2360)
					return
				} else {
					tmp11315 := PrimIsPair(V2357)

					var ifres11312 Obj

					if True == tmp11315 {
						tmp11314 := PrimIsPair(V2358)

						var ifres11313 Obj

						if True == tmp11314 {
							ifres11313 = True

						} else {
							ifres11313 = False

						}

						ifres11312 = ifres11313

					} else {
						ifres11312 = False

					}

					if True == ifres11312 {
						tmp11302 := PrimHead(V2357)

						tmp11303 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11302, V2359)

						tmp11304 := PrimHead(V2358)

						tmp11305 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11304, V2359)

						tmp11306 := MakeNative(func(__e *ControlFlow) {
							tmp11307 := PrimTail(V2357)

							tmp11308 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11307, V2359)

							tmp11309 := PrimTail(V2358)

							tmp11310 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11309, V2359)

							__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11308, tmp11310, V2359, V2360)
							return

						}, 0)

						__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11303, tmp11305, V2359, tmp11306)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp11332 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp11301)

	_ = tmp11332

	tmp11333 := MakeNative(func(__e *ControlFlow) {
		V2372 := __e.Get(1)
		_ = V2372
		V2373 := __e.Get(2)
		_ = V2373
		V2374 := __e.Get(3)
		_ = V2374
		V2375 := __e.Get(4)
		_ = V2375
		tmp11353 := PrimEqual(V2372, V2373)

		if True == tmp11353 {
			__e.TailApply(PrimFunc(symthaw), V2375)
			return
		} else {
			tmp11351 := Call(__e, PrimFunc(symshen_4pvar_2), V2372)

			if True == tmp11351 {
				__e.TailApply(PrimFunc(symshen_4bind_b), V2372, V2373, V2374, V2375)
				return
			} else {
				tmp11349 := Call(__e, PrimFunc(symshen_4pvar_2), V2373)

				if True == tmp11349 {
					__e.TailApply(PrimFunc(symshen_4bind_b), V2373, V2372, V2374, V2375)
					return
				} else {
					tmp11347 := PrimIsPair(V2372)

					var ifres11344 Obj

					if True == tmp11347 {
						tmp11346 := PrimIsPair(V2373)

						var ifres11345 Obj

						if True == tmp11346 {
							ifres11345 = True

						} else {
							ifres11345 = False

						}

						ifres11344 = ifres11345

					} else {
						ifres11344 = False

					}

					if True == ifres11344 {
						tmp11334 := PrimHead(V2372)

						tmp11335 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11334, V2374)

						tmp11336 := PrimHead(V2373)

						tmp11337 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11336, V2374)

						tmp11338 := MakeNative(func(__e *ControlFlow) {
							tmp11339 := PrimTail(V2372)

							tmp11340 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11339, V2374)

							tmp11341 := PrimTail(V2373)

							tmp11342 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11341, V2374)

							__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11340, tmp11342, V2374, V2375)
							return

						}, 0)

						__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11335, tmp11337, V2374, tmp11338)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp11354 := Call(__e, ns2_1set, symshen_4lzy_a, tmp11333)

	_ = tmp11354

	tmp11355 := MakeNative(func(__e *ControlFlow) {
		V2381 := __e.Get(1)
		_ = V2381
		V2382 := __e.Get(2)
		_ = V2382
		tmp11365 := PrimEqual(V2381, V2382)

		if True == tmp11365 {
			__e.Return(True)
			return
		} else {
			tmp11363 := PrimIsPair(V2382)

			if True == tmp11363 {
				tmp11360 := PrimHead(V2382)

				tmp11361 := Call(__e, PrimFunc(symshen_4occurs_2), V2381, tmp11360)

				if True == tmp11361 {
					__e.Return(True)
					return
				} else {
					tmp11357 := PrimTail(V2382)

					tmp11358 := Call(__e, PrimFunc(symshen_4occurs_2), V2381, tmp11357)

					if True == tmp11358 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp11366 := Call(__e, ns2_1set, symshen_4occurs_2, tmp11355)

	_ = tmp11366

	tmp11367 := MakeNative(func(__e *ControlFlow) {
		V2383 := __e.Get(1)
		_ = V2383
		V2384 := __e.Get(2)
		_ = V2384
		V2385 := __e.Get(3)
		_ = V2385
		V2386 := __e.Get(4)
		_ = V2386
		V2387 := __e.Get(5)
		_ = V2387
		tmp11368 := Call(__e, V2383, V2384)

		tmp11369 := Call(__e, tmp11368, V2385)

		tmp11370 := Call(__e, tmp11369, V2386)

		__e.TailApply(tmp11370, V2387)
		return

	}, 5)

	tmp11371 := Call(__e, ns2_1set, symcall, tmp11367)

	_ = tmp11371

	tmp11372 := MakeNative(func(__e *ControlFlow) {
		V2394 := __e.Get(1)
		_ = V2394
		V2395 := __e.Get(2)
		_ = V2395
		V2396 := __e.Get(3)
		_ = V2396
		V2397 := __e.Get(4)
		_ = V2397
		V2398 := __e.Get(5)
		_ = V2398
		__e.TailApply(PrimFunc(symshen_4deref), V2394, V2395)
		return
	}, 5)

	tmp11373 := Call(__e, ns2_1set, symreturn, tmp11372)

	_ = tmp11373

	tmp11374 := MakeNative(func(__e *ControlFlow) {
		V2405 := __e.Get(1)
		_ = V2405
		V2406 := __e.Get(2)
		_ = V2406
		V2407 := __e.Get(3)
		_ = V2407
		V2408 := __e.Get(4)
		_ = V2408
		V2409 := __e.Get(5)
		_ = V2409
		if True == V2405 {
			__e.TailApply(PrimFunc(symthaw), V2409)
			return
		} else {
			__e.Return(False)
			return
		}
	}, 5)

	tmp11376 := Call(__e, ns2_1set, symwhen, tmp11374)

	_ = tmp11376

	tmp11377 := MakeNative(func(__e *ControlFlow) {
		V2410 := __e.Get(1)
		_ = V2410
		V2411 := __e.Get(2)
		_ = V2411
		V2412 := __e.Get(3)
		_ = V2412
		V2413 := __e.Get(4)
		_ = V2413
		V2414 := __e.Get(5)
		_ = V2414
		V2415 := __e.Get(6)
		_ = V2415
		tmp11378 := Call(__e, PrimFunc(symshen_4lazyderef), V2410, V2412)

		tmp11379 := Call(__e, PrimFunc(symshen_4lazyderef), V2411, V2412)

		__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11378, tmp11379, V2412, V2415)
		return

	}, 6)

	tmp11380 := Call(__e, ns2_1set, symis, tmp11377)

	_ = tmp11380

	tmp11381 := MakeNative(func(__e *ControlFlow) {
		V2416 := __e.Get(1)
		_ = V2416
		V2417 := __e.Get(2)
		_ = V2417
		V2418 := __e.Get(3)
		_ = V2418
		V2419 := __e.Get(4)
		_ = V2419
		V2420 := __e.Get(5)
		_ = V2420
		V2421 := __e.Get(6)
		_ = V2421
		tmp11382 := Call(__e, PrimFunc(symshen_4lazyderef), V2416, V2418)

		tmp11383 := Call(__e, PrimFunc(symshen_4lazyderef), V2417, V2418)

		__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11382, tmp11383, V2418, V2421)
		return

	}, 6)

	tmp11384 := Call(__e, ns2_1set, symis_b, tmp11381)

	_ = tmp11384

	tmp11385 := MakeNative(func(__e *ControlFlow) {
		V2426 := __e.Get(1)
		_ = V2426
		V2427 := __e.Get(2)
		_ = V2427
		V2428 := __e.Get(3)
		_ = V2428
		V2429 := __e.Get(4)
		_ = V2429
		V2430 := __e.Get(5)
		_ = V2430
		V2431 := __e.Get(6)
		_ = V2431
		__e.TailApply(PrimFunc(symshen_4bind_b), V2426, V2427, V2428, V2431)
		return
	}, 6)

	tmp11386 := Call(__e, ns2_1set, symbind, tmp11385)

	_ = tmp11386

	tmp11387 := MakeNative(func(__e *ControlFlow) {
		V2432 := __e.Get(1)
		_ = V2432
		V2433 := __e.Get(2)
		_ = V2433
		V2434 := __e.Get(3)
		_ = V2434
		V2435 := __e.Get(4)
		_ = V2435
		V2436 := __e.Get(5)
		_ = V2436
		tmp11389 := Call(__e, PrimFunc(symshen_4lazyderef), V2432, V2433)

		tmp11390 := Call(__e, PrimFunc(symshen_4pvar_2), tmp11389)

		if True == tmp11390 {
			__e.TailApply(PrimFunc(symthaw), V2436)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 5)

	tmp11391 := Call(__e, ns2_1set, symvar_2, tmp11387)

	_ = tmp11391

	tmp11392 := MakeNative(func(__e *ControlFlow) {
		V2439 := __e.Get(1)
		_ = V2439
		__e.Return(MakeString("|prolog vector|"))
		return
	}, 1)

	tmp11393 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp11392)

	_ = tmp11393

	tmp11394 := MakeNative(func(__e *ControlFlow) {
		V2458 := __e.Get(1)
		_ = V2458
		V2459 := __e.Get(2)
		_ = V2459
		V2460 := __e.Get(3)
		_ = V2460
		V2461 := __e.Get(4)
		_ = V2461
		V2462 := __e.Get(5)
		_ = V2462
		tmp11407 := PrimEqual(Nil, V2458)

		if True == tmp11407 {
			__e.Return(False)
			return
		} else {
			tmp11405 := PrimIsPair(V2458)

			if True == tmp11405 {
				tmp11395 := MakeNative(func(__e *ControlFlow) {
					Case := __e.Get(1)
					_ = Case
					tmp11398 := PrimEqual(Case, False)

					if True == tmp11398 {
						tmp11396 := PrimTail(V2458)

						__e.TailApply(PrimFunc(symfork), tmp11396, V2459, V2460, V2461, V2462)
						return

					} else {
						__e.Return(Case)
						return
					}

				}, 1)

				tmp11399 := PrimHead(V2458)

				tmp11400 := Call(__e, tmp11399, V2459)

				tmp11401 := Call(__e, tmp11400, V2460)

				tmp11402 := Call(__e, tmp11401, V2461)

				tmp11403 := Call(__e, tmp11402, V2462)

				__e.TailApply(tmp11395, tmp11403)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
				return
			}

		}

	}, 5)

	tmp11408 := Call(__e, ns2_1set, symfork, tmp11394)

	_ = tmp11408

	tmp11409 := MakeNative(func(__e *ControlFlow) {
		V2463 := __e.Get(1)
		_ = V2463
		V2464 := __e.Get(2)
		_ = V2464
		V2465 := __e.Get(3)
		_ = V2465
		V2466 := __e.Get(4)
		_ = V2466
		V2467 := __e.Get(5)
		_ = V2467
		V2468 := __e.Get(6)
		_ = V2468
		V2469 := __e.Get(7)
		_ = V2469
		tmp11416 := Call(__e, PrimFunc(symshen_4unlocked_2), V2467)

		if True == tmp11416 {
			tmp11410 := MakeNative(func(__e *ControlFlow) {
				Store := __e.Get(1)
				_ = Store
				tmp11411 := Call(__e, PrimFunc(symshen_4incinfs))

				_ = tmp11411

				tmp11412 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimFunc(symshen_4findall_1h), V2463, V2464, V2465, Store, V2466, V2467, V2468, V2469)
					return
				}, 0)

				tmp11413 := Call(__e, PrimFunc(symis), Store, Nil, V2466, V2467, V2468, tmp11412)

				__e.TailApply(PrimFunc(symshen_4gc), V2466, tmp11413)
				return

			}, 1)

			tmp11414 := Call(__e, PrimFunc(symshen_4newpv), V2466)

			__e.TailApply(tmp11410, tmp11414)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 7)

	tmp11417 := Call(__e, ns2_1set, symfindall, tmp11409)

	_ = tmp11417

	tmp11418 := MakeNative(func(__e *ControlFlow) {
		V2470 := __e.Get(1)
		_ = V2470
		V2471 := __e.Get(2)
		_ = V2471
		V2472 := __e.Get(3)
		_ = V2472
		V2473 := __e.Get(4)
		_ = V2473
		V2474 := __e.Get(5)
		_ = V2474
		V2475 := __e.Get(6)
		_ = V2475
		V2476 := __e.Get(7)
		_ = V2476
		V2477 := __e.Get(8)
		_ = V2477
		tmp11419 := MakeNative(func(__e *ControlFlow) {
			C1894 := __e.Get(1)
			_ = C1894
			tmp11424 := PrimEqual(C1894, False)

			if True == tmp11424 {
				tmp11422 := Call(__e, PrimFunc(symshen_4unlocked_2), V2475)

				if True == tmp11422 {
					tmp11420 := Call(__e, PrimFunc(symshen_4incinfs))

					_ = tmp11420

					__e.TailApply(PrimFunc(symis_b), V2472, V2473, V2474, V2475, V2476, V2477)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C1894)
				return
			}

		}, 1)

		tmp11429 := Call(__e, PrimFunc(symshen_4unlocked_2), V2475)

		var ifres11425 Obj

		if True == tmp11429 {
			tmp11426 := Call(__e, PrimFunc(symshen_4incinfs))

			_ = tmp11426

			tmp11427 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimFunc(symshen_4overbind), V2470, V2473, V2474, V2475, V2476, V2477)
				return
			}, 0)

			tmp11428 := Call(__e, PrimFunc(symcall), V2471, V2474, V2475, V2476, tmp11427)

			ifres11425 = tmp11428

		} else {
			ifres11425 = False

		}

		__e.TailApply(tmp11419, ifres11425)
		return

	}, 8)

	tmp11430 := Call(__e, ns2_1set, symshen_4findall_1h, tmp11418)

	_ = tmp11430

	tmp11431 := MakeNative(func(__e *ControlFlow) {
		V2484 := __e.Get(1)
		_ = V2484
		V2485 := __e.Get(2)
		_ = V2485
		V2486 := __e.Get(3)
		_ = V2486
		V2487 := __e.Get(4)
		_ = V2487
		V2488 := __e.Get(5)
		_ = V2488
		V2489 := __e.Get(6)
		_ = V2489
		tmp11432 := Call(__e, PrimFunc(symshen_4deref), V2484, V2486)

		tmp11433 := Call(__e, PrimFunc(symshen_4lazyderef), V2485, V2486)

		tmp11434 := PrimCons(tmp11432, tmp11433)

		tmp11435 := Call(__e, PrimFunc(symshen_4bindv), V2485, tmp11434, V2486)

		_ = tmp11435

		__e.Return(False)
		return

	}, 6)

	tmp11436 := Call(__e, ns2_1set, symshen_4overbind, tmp11431)

	_ = tmp11436

	tmp11437 := MakeNative(func(__e *ControlFlow) {
		V2492 := __e.Get(1)
		_ = V2492
		tmp11441 := PrimEqual(sym_7, V2492)

		if True == tmp11441 {
			__e.Return(PrimSet(symshen_4_doccurs_d, True))
			return
		} else {
			tmp11439 := PrimEqual(sym_1, V2492)

			if True == tmp11439 {
				__e.Return(PrimSet(symshen_4_doccurs_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	__e.TailApply(ns2_1set, symoccurs_1check, tmp11437)
	return

}, 0)
