package main

import . "github.com/tiancaiamao/shen-go/cora"

var PrologMain = MakeNative(func(__e *ControlFlow) {
	tmp8641 := MakeNative(func(__e *ControlFlow) {
		V2055 := __e.Get(1)
		_ = V2055
		V2056 := __e.Get(2)
		_ = V2056
		tmp8642 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_5defprolog_6), X)
			return
		}, 1)

		tmp8643 := PrimCons(V2055, V2056)

		__e.TailApply(PrimNS2Value(symcompile), tmp8642, tmp8643)
		return

	}, 2)

	tmp8644 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1prolog, tmp8641)

	_ = tmp8644

	tmp8645 := MakeNative(func(__e *ControlFlow) {
		V2057 := __e.Get(1)
		_ = V2057
		tmp8646 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp8648 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp8648 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp8669 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2057)

		var ifres8649 Obj

		if True == tmp8669 {
			tmp8651 := MakeNative(func(__e *ControlFlow) {
				F := __e.Get(1)
				_ = F
				tmp8652 := MakeNative(func(__e *ControlFlow) {
					News1851 := __e.Get(1)
					_ = News1851
					tmp8653 := MakeNative(func(__e *ControlFlow) {
						Parseshen_4_5clauses_6 := __e.Get(1)
						_ = Parseshen_4_5clauses_6
						tmp8664 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5clauses_6)

						if True == tmp8664 {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						} else {
							tmp8655 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5clauses_6)

							tmp8656 := MakeNative(func(__e *ControlFlow) {
								Aritycheck := __e.Get(1)
								_ = Aritycheck
								tmp8657 := MakeNative(func(__e *ControlFlow) {
									LeftLinear := __e.Get(1)
									_ = LeftLinear
									__e.TailApply(PrimNS2Value(symshen_4horn_1clause_1procedure), F, LeftLinear)
									return
								}, 1)

								tmp8658 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									__e.TailApply(PrimNS2Value(symshen_4linearise_1clause), X)
									return
								}, 1)

								tmp8659 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5clauses_6)

								tmp8660 := Call(__e, PrimNS2Value(symmap), tmp8658, tmp8659)

								__e.TailApply(tmp8657, tmp8660)
								return

							}, 1)

							tmp8661 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5clauses_6)

							tmp8662 := Call(__e, PrimNS2Value(symshen_4prolog_1arity_1check), F, tmp8661)

							tmp8663 := Call(__e, tmp8656, tmp8662)

							__e.TailApply(PrimNS2Value(symshen_4comb), tmp8655, tmp8663)
							return

						}

					}, 1)

					tmp8665 := Call(__e, PrimNS2Value(symshen_4_5clauses_6), News1851)

					__e.TailApply(tmp8653, tmp8665)
					return

				}, 1)

				tmp8666 := Call(__e, PrimNS2Value(symshen_4tls), V2057)

				__e.TailApply(tmp8652, tmp8666)
				return

			}, 1)

			tmp8667 := Call(__e, PrimNS2Value(symshen_4hds), V2057)

			tmp8668 := Call(__e, tmp8651, tmp8667)

			ifres8649 = tmp8668

		} else {
			tmp8650 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres8649 = tmp8650

		}

		__e.TailApply(tmp8646, ifres8649)
		return

	}, 1)

	tmp8670 := Call(__e, PrimNS2Value(symdef), symshen_4_5defprolog_6, tmp8645)

	_ = tmp8670

	tmp8671 := MakeNative(func(__e *ControlFlow) {
		V2058 := __e.Get(1)
		_ = V2058
		tmp8672 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp8690 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp8690 {
				tmp8674 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp8676 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp8676 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp8677 := MakeNative(func(__e *ControlFlow) {
					Parse_5_b_6 := __e.Get(1)
					_ = Parse_5_b_6
					tmp8687 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5_b_6)

					if True == tmp8687 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp8679 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5_b_6)

						tmp8685 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

						tmp8686 := Call(__e, PrimNS2Value(symempty_2), tmp8685)

						var ifres8680 Obj

						if True == tmp8686 {
							ifres8680 = Nil

						} else {
							tmp8681 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

							tmp8682 := Call(__e, PrimNS2Value(symshen_4app), tmp8681, MakeString("\n ..."), symshen_4r)

							tmp8683 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp8682)

							tmp8684 := PrimSimpleError(tmp8683)

							ifres8680 = tmp8684

						}

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp8679, ifres8680)
						return

					}

				}, 1)

				tmp8688 := Call(__e, PrimNS2Value(sym_5_b_6), V2058)

				tmp8689 := Call(__e, tmp8677, tmp8688)

				__e.TailApply(tmp8674, tmp8689)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp8691 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5clause_6 := __e.Get(1)
			_ = Parseshen_4_5clause_6
			tmp8701 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5clause_6)

			if True == tmp8701 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp8693 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5clauses_6 := __e.Get(1)
					_ = Parseshen_4_5clauses_6
					tmp8699 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5clauses_6)

					if True == tmp8699 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp8695 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5clauses_6)

						tmp8696 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5clause_6)

						tmp8697 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5clauses_6)

						tmp8698 := PrimCons(tmp8696, tmp8697)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp8695, tmp8698)
						return

					}

				}, 1)

				tmp8700 := Call(__e, PrimNS2Value(symshen_4_5clauses_6), Parseshen_4_5clause_6)

				__e.TailApply(tmp8693, tmp8700)
				return

			}

		}, 1)

		tmp8702 := Call(__e, PrimNS2Value(symshen_4_5clause_6), V2058)

		tmp8703 := Call(__e, tmp8691, tmp8702)

		__e.TailApply(tmp8672, tmp8703)
		return

	}, 1)

	tmp8704 := Call(__e, PrimNS2Value(symdef), symshen_4_5clauses_6, tmp8671)

	_ = tmp8704

	tmp8705 := MakeNative(func(__e *ControlFlow) {
		V2063 := __e.Get(1)
		_ = V2063
		V2064 := __e.Get(2)
		_ = V2064
		tmp8732 := PrimIsPair(V2064)

		var ifres8728 Obj

		if True == tmp8732 {
			tmp8730 := PrimTail(V2064)

			tmp8731 := PrimEqual(Nil, tmp8730)

			var ifres8729 Obj

			if True == tmp8731 {
				ifres8729 = True

			} else {
				ifres8729 = False

			}

			ifres8728 = ifres8729

		} else {
			ifres8728 = False

		}

		if True == ifres8728 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp8727 := PrimIsPair(V2064)

			var ifres8712 Obj

			if True == tmp8727 {
				tmp8725 := PrimHead(V2064)

				tmp8726 := PrimIsPair(tmp8725)

				var ifres8714 Obj

				if True == tmp8726 {
					tmp8722 := PrimHead(V2064)

					tmp8723 := PrimTail(tmp8722)

					tmp8724 := PrimIsPair(tmp8723)

					var ifres8716 Obj

					if True == tmp8724 {
						tmp8718 := PrimHead(V2064)

						tmp8719 := PrimTail(tmp8718)

						tmp8720 := PrimTail(tmp8719)

						tmp8721 := PrimEqual(Nil, tmp8720)

						var ifres8717 Obj

						if True == tmp8721 {
							ifres8717 = True

						} else {
							ifres8717 = False

						}

						ifres8716 = ifres8717

					} else {
						ifres8716 = False

					}

					var ifres8715 Obj

					if True == ifres8716 {
						ifres8715 = True

					} else {
						ifres8715 = False

					}

					ifres8714 = ifres8715

				} else {
					ifres8714 = False

				}

				var ifres8713 Obj

				if True == ifres8714 {
					ifres8713 = True

				} else {
					ifres8713 = False

				}

				ifres8712 = ifres8713

			} else {
				ifres8712 = False

			}

			if True == ifres8712 {
				tmp8708 := PrimHead(V2064)

				tmp8709 := PrimHead(tmp8708)

				tmp8710 := Call(__e, PrimNS2Value(symlength), tmp8709)

				tmp8711 := PrimTail(V2064)

				__e.TailApply(PrimNS2Value(symshen_4pac_1h), V2063, tmp8710, tmp8711)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4prolog_1arity_1check)
				return
			}

		}

	}, 2)

	tmp8733 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1arity_1check, tmp8705)

	_ = tmp8733

	tmp8734 := MakeNative(func(__e *ControlFlow) {
		V2065 := __e.Get(1)
		_ = V2065
		tmp8750 := PrimIsPair(V2065)

		var ifres8741 Obj

		if True == tmp8750 {
			tmp8748 := PrimTail(V2065)

			tmp8749 := PrimIsPair(tmp8748)

			var ifres8743 Obj

			if True == tmp8749 {
				tmp8745 := PrimTail(V2065)

				tmp8746 := PrimTail(tmp8745)

				tmp8747 := PrimEqual(Nil, tmp8746)

				var ifres8744 Obj

				if True == tmp8747 {
					ifres8744 = True

				} else {
					ifres8744 = False

				}

				ifres8743 = ifres8744

			} else {
				ifres8743 = False

			}

			var ifres8742 Obj

			if True == ifres8743 {
				ifres8742 = True

			} else {
				ifres8742 = False

			}

			ifres8741 = ifres8742

		} else {
			ifres8741 = False

		}

		if True == ifres8741 {
			tmp8736 := PrimHead(V2065)

			tmp8737 := PrimTail(V2065)

			tmp8738 := PrimHead(tmp8737)

			tmp8739 := Call(__e, PrimNS2Value(sym_8p), tmp8736, tmp8738)

			tmp8740 := Call(__e, PrimNS2Value(symshen_4linearise), tmp8739)

			__e.TailApply(PrimNS2Value(symshen_4lch), tmp8740)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4linearise_1clause)
			return
		}

	}, 1)

	tmp8751 := Call(__e, PrimNS2Value(symdef), symshen_4linearise_1clause, tmp8734)

	_ = tmp8751

	tmp8752 := MakeNative(func(__e *ControlFlow) {
		V2066 := __e.Get(1)
		_ = V2066
		tmp8758 := Call(__e, PrimNS2Value(symtuple_2), V2066)

		if True == tmp8758 {
			tmp8754 := Call(__e, PrimNS2Value(symfst), V2066)

			tmp8755 := Call(__e, PrimNS2Value(symsnd), V2066)

			tmp8756 := Call(__e, PrimNS2Value(symshen_4lchh), tmp8755)

			tmp8757 := PrimCons(tmp8756, Nil)

			__e.Return(PrimCons(tmp8754, tmp8757))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4lch)
			return
		}

	}, 1)

	tmp8759 := Call(__e, PrimNS2Value(symdef), symshen_4lch, tmp8752)

	_ = tmp8759

	tmp8760 := MakeNative(func(__e *ControlFlow) {
		V2067 := __e.Get(1)
		_ = V2067
		tmp8823 := PrimIsPair(V2067)

		var ifres8772 Obj

		if True == tmp8823 {
			tmp8821 := PrimHead(V2067)

			tmp8822 := PrimEqual(symwhere, tmp8821)

			var ifres8774 Obj

			if True == tmp8822 {
				tmp8819 := PrimTail(V2067)

				tmp8820 := PrimIsPair(tmp8819)

				var ifres8776 Obj

				if True == tmp8820 {
					tmp8816 := PrimTail(V2067)

					tmp8817 := PrimHead(tmp8816)

					tmp8818 := PrimIsPair(tmp8817)

					var ifres8778 Obj

					if True == tmp8818 {
						tmp8812 := PrimTail(V2067)

						tmp8813 := PrimHead(tmp8812)

						tmp8814 := PrimHead(tmp8813)

						tmp8815 := PrimEqual(sym_a, tmp8814)

						var ifres8780 Obj

						if True == tmp8815 {
							tmp8808 := PrimTail(V2067)

							tmp8809 := PrimHead(tmp8808)

							tmp8810 := PrimTail(tmp8809)

							tmp8811 := PrimIsPair(tmp8810)

							var ifres8782 Obj

							if True == tmp8811 {
								tmp8803 := PrimTail(V2067)

								tmp8804 := PrimHead(tmp8803)

								tmp8805 := PrimTail(tmp8804)

								tmp8806 := PrimTail(tmp8805)

								tmp8807 := PrimIsPair(tmp8806)

								var ifres8784 Obj

								if True == tmp8807 {
									tmp8797 := PrimTail(V2067)

									tmp8798 := PrimHead(tmp8797)

									tmp8799 := PrimTail(tmp8798)

									tmp8800 := PrimTail(tmp8799)

									tmp8801 := PrimTail(tmp8800)

									tmp8802 := PrimEqual(Nil, tmp8801)

									var ifres8786 Obj

									if True == tmp8802 {
										tmp8794 := PrimTail(V2067)

										tmp8795 := PrimTail(tmp8794)

										tmp8796 := PrimIsPair(tmp8795)

										var ifres8788 Obj

										if True == tmp8796 {
											tmp8790 := PrimTail(V2067)

											tmp8791 := PrimTail(tmp8790)

											tmp8792 := PrimTail(tmp8791)

											tmp8793 := PrimEqual(Nil, tmp8792)

											var ifres8789 Obj

											if True == tmp8793 {
												ifres8789 = True

											} else {
												ifres8789 = False

											}

											ifres8788 = ifres8789

										} else {
											ifres8788 = False

										}

										var ifres8787 Obj

										if True == ifres8788 {
											ifres8787 = True

										} else {
											ifres8787 = False

										}

										ifres8786 = ifres8787

									} else {
										ifres8786 = False

									}

									var ifres8785 Obj

									if True == ifres8786 {
										ifres8785 = True

									} else {
										ifres8785 = False

									}

									ifres8784 = ifres8785

								} else {
									ifres8784 = False

								}

								var ifres8783 Obj

								if True == ifres8784 {
									ifres8783 = True

								} else {
									ifres8783 = False

								}

								ifres8782 = ifres8783

							} else {
								ifres8782 = False

							}

							var ifres8781 Obj

							if True == ifres8782 {
								ifres8781 = True

							} else {
								ifres8781 = False

							}

							ifres8780 = ifres8781

						} else {
							ifres8780 = False

						}

						var ifres8779 Obj

						if True == ifres8780 {
							ifres8779 = True

						} else {
							ifres8779 = False

						}

						ifres8778 = ifres8779

					} else {
						ifres8778 = False

					}

					var ifres8777 Obj

					if True == ifres8778 {
						ifres8777 = True

					} else {
						ifres8777 = False

					}

					ifres8776 = ifres8777

				} else {
					ifres8776 = False

				}

				var ifres8775 Obj

				if True == ifres8776 {
					ifres8775 = True

				} else {
					ifres8775 = False

				}

				ifres8774 = ifres8775

			} else {
				ifres8774 = False

			}

			var ifres8773 Obj

			if True == ifres8774 {
				ifres8773 = True

			} else {
				ifres8773 = False

			}

			ifres8772 = ifres8773

		} else {
			ifres8772 = False

		}

		if True == ifres8772 {
			tmp8763 := PrimNS3Value(symshen_4_doccurs_d)

			var ifres8762 Obj

			if True == tmp8763 {
				ifres8762 = symis_b

			} else {
				ifres8762 = symis

			}

			tmp8764 := PrimTail(V2067)

			tmp8765 := PrimHead(tmp8764)

			tmp8766 := PrimTail(tmp8765)

			tmp8767 := PrimCons(ifres8762, tmp8766)

			tmp8768 := PrimTail(V2067)

			tmp8769 := PrimTail(tmp8768)

			tmp8770 := PrimHead(tmp8769)

			tmp8771 := Call(__e, PrimNS2Value(symshen_4lchh), tmp8770)

			__e.Return(PrimCons(tmp8767, tmp8771))
			return

		} else {
			__e.Return(V2067)
			return
		}

	}, 1)

	tmp8824 := Call(__e, PrimNS2Value(symdef), symshen_4lchh, tmp8760)

	_ = tmp8824

	tmp8825 := MakeNative(func(__e *ControlFlow) {
		V2074 := __e.Get(1)
		_ = V2074
		V2075 := __e.Get(2)
		_ = V2075
		V2076 := __e.Get(3)
		_ = V2076
		tmp8841 := PrimEqual(Nil, V2076)

		if True == tmp8841 {
			__e.Return(True)
			return
		} else {
			tmp8840 := PrimIsPair(V2076)

			var ifres8836 Obj

			if True == tmp8840 {
				tmp8838 := PrimHead(V2076)

				tmp8839 := PrimIsPair(tmp8838)

				var ifres8837 Obj

				if True == tmp8839 {
					ifres8837 = True

				} else {
					ifres8837 = False

				}

				ifres8836 = ifres8837

			} else {
				ifres8836 = False

			}

			if True == ifres8836 {
				tmp8832 := PrimHead(V2076)

				tmp8833 := PrimHead(tmp8832)

				tmp8834 := Call(__e, PrimNS2Value(symlength), tmp8833)

				tmp8835 := PrimEqual(V2075, tmp8834)

				if True == tmp8835 {
					tmp8831 := PrimTail(V2076)

					__e.TailApply(PrimNS2Value(symshen_4pac_1h), V2074, V2075, tmp8831)
					return

				} else {
					tmp8829 := Call(__e, PrimNS2Value(symshen_4app), V2074, MakeString("\n"), symshen_4a)

					tmp8830 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp8829)

					__e.Return(PrimSimpleError(tmp8830))
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4pac_1h)
				return
			}

		}

	}, 3)

	tmp8842 := Call(__e, PrimNS2Value(symdef), symshen_4pac_1h, tmp8825)

	_ = tmp8842

	tmp8843 := MakeNative(func(__e *ControlFlow) {
		V2077 := __e.Get(1)
		_ = V2077
		tmp8844 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp8846 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp8846 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp8847 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5head_6 := __e.Get(1)
			_ = Parseshen_4_5head_6
			tmp8866 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5head_6)

			if True == tmp8866 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp8865 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5head_6, sym_5_1_1)

				if True == tmp8865 {
					tmp8850 := MakeNative(func(__e *ControlFlow) {
						News1854 := __e.Get(1)
						_ = News1854
						tmp8851 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5body_6 := __e.Get(1)
							_ = Parseshen_4_5body_6
							tmp8862 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5body_6)

							if True == tmp8862 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp8853 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5sc_6 := __e.Get(1)
									_ = Parseshen_4_5sc_6
									tmp8860 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

									if True == tmp8860 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp8855 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sc_6)

										tmp8856 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5head_6)

										tmp8857 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5body_6)

										tmp8858 := PrimCons(tmp8857, Nil)

										tmp8859 := PrimCons(tmp8856, tmp8858)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp8855, tmp8859)
										return

									}

								}, 1)

								tmp8861 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5body_6)

								__e.TailApply(tmp8853, tmp8861)
								return

							}

						}, 1)

						tmp8863 := Call(__e, PrimNS2Value(symshen_4_5body_6), News1854)

						__e.TailApply(tmp8851, tmp8863)
						return

					}, 1)

					tmp8864 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5head_6)

					__e.TailApply(tmp8850, tmp8864)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp8867 := Call(__e, PrimNS2Value(symshen_4_5head_6), V2077)

		tmp8868 := Call(__e, tmp8847, tmp8867)

		__e.TailApply(tmp8844, tmp8868)
		return

	}, 1)

	tmp8869 := Call(__e, PrimNS2Value(symdef), symshen_4_5clause_6, tmp8843)

	_ = tmp8869

	tmp8870 := MakeNative(func(__e *ControlFlow) {
		V2078 := __e.Get(1)
		_ = V2078
		tmp8871 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp8882 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp8882 {
				tmp8873 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp8875 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp8875 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp8876 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp8879 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp8879 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp8878 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp8878, Nil)
						return

					}

				}, 1)

				tmp8880 := Call(__e, PrimNS2Value(sym_5e_6), V2078)

				tmp8881 := Call(__e, tmp8876, tmp8880)

				__e.TailApply(tmp8873, tmp8881)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp8883 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5hterm_6 := __e.Get(1)
			_ = Parseshen_4_5hterm_6
			tmp8893 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

			if True == tmp8893 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp8885 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5head_6 := __e.Get(1)
					_ = Parseshen_4_5head_6
					tmp8891 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5head_6)

					if True == tmp8891 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp8887 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5head_6)

						tmp8888 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

						tmp8889 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5head_6)

						tmp8890 := PrimCons(tmp8888, tmp8889)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp8887, tmp8890)
						return

					}

				}, 1)

				tmp8892 := Call(__e, PrimNS2Value(symshen_4_5head_6), Parseshen_4_5hterm_6)

				__e.TailApply(tmp8885, tmp8892)
				return

			}

		}, 1)

		tmp8894 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), V2078)

		tmp8895 := Call(__e, tmp8883, tmp8894)

		__e.TailApply(tmp8871, tmp8895)
		return

	}, 1)

	tmp8896 := Call(__e, PrimNS2Value(symdef), symshen_4_5head_6, tmp8870)

	_ = tmp8896

	tmp8897 := MakeNative(func(__e *ControlFlow) {
		V2079 := __e.Get(1)
		_ = V2079
		tmp8898 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9069 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9069 {
				tmp8900 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9056 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp9056 {
						tmp8902 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp9024 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp9024 {
								tmp8904 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp8998 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

									if True == tmp8998 {
										tmp8906 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp8972 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

											if True == tmp8972 {
												tmp8908 := MakeNative(func(__e *ControlFlow) {
													Result := __e.Get(1)
													_ = Result
													tmp8942 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

													if True == tmp8942 {
														tmp8910 := MakeNative(func(__e *ControlFlow) {
															Result := __e.Get(1)
															_ = Result
															tmp8912 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

															if True == tmp8912 {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															} else {
																__e.Return(Result)
																return
															}

														}, 1)

														tmp8941 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2079)

														var ifres8913 Obj

														if True == tmp8941 {
															tmp8915 := MakeNative(func(__e *ControlFlow) {
																SynCons := __e.Get(1)
																_ = SynCons
																tmp8936 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, symmode)

																if True == tmp8936 {
																	tmp8917 := MakeNative(func(__e *ControlFlow) {
																		News1864 := __e.Get(1)
																		_ = News1864
																		tmp8918 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5hterm_6 := __e.Get(1)
																			_ = Parseshen_4_5hterm_6
																			tmp8933 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

																			if True == tmp8933 {
																				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																				return
																			} else {
																				tmp8932 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5hterm_6, sym_1)

																				if True == tmp8932 {
																					tmp8921 := MakeNative(func(__e *ControlFlow) {
																						News1865 := __e.Get(1)
																						_ = News1865
																						tmp8922 := MakeNative(func(__e *ControlFlow) {
																							Parseshen_4_5end_6 := __e.Get(1)
																							_ = Parseshen_4_5end_6
																							tmp8929 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

																							if True == tmp8929 {
																								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																								return
																							} else {
																								tmp8924 := Call(__e, PrimNS2Value(symshen_4tlstream), V2079)

																								tmp8925 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp8924)

																								tmp8926 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

																								tmp8927 := PrimCons(tmp8926, Nil)

																								tmp8928 := PrimCons(symshen_4_1m, tmp8927)

																								__e.TailApply(PrimNS2Value(symshen_4comb), tmp8925, tmp8928)
																								return

																							}

																						}, 1)

																						tmp8930 := Call(__e, PrimNS2Value(symshen_4_5end_6), News1865)

																						__e.TailApply(tmp8922, tmp8930)
																						return

																					}, 1)

																					tmp8931 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5hterm_6)

																					__e.TailApply(tmp8921, tmp8931)
																					return

																				} else {
																					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																					return
																				}

																			}

																		}, 1)

																		tmp8934 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), News1864)

																		__e.TailApply(tmp8918, tmp8934)
																		return

																	}, 1)

																	tmp8935 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

																	__e.TailApply(tmp8917, tmp8935)
																	return

																} else {
																	__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																	return
																}

															}, 1)

															tmp8937 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

															tmp8938 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2079)

															tmp8939 := Call(__e, PrimNS2Value(symshen_4comb), tmp8937, tmp8938)

															tmp8940 := Call(__e, tmp8915, tmp8939)

															ifres8913 = tmp8940

														} else {
															tmp8914 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

															ifres8913 = tmp8914

														}

														__e.TailApply(tmp8910, ifres8913)
														return

													} else {
														__e.Return(Result)
														return
													}

												}, 1)

												tmp8971 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2079)

												var ifres8943 Obj

												if True == tmp8971 {
													tmp8945 := MakeNative(func(__e *ControlFlow) {
														SynCons := __e.Get(1)
														_ = SynCons
														tmp8966 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, symmode)

														if True == tmp8966 {
															tmp8947 := MakeNative(func(__e *ControlFlow) {
																News1862 := __e.Get(1)
																_ = News1862
																tmp8948 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5hterm_6 := __e.Get(1)
																	_ = Parseshen_4_5hterm_6
																	tmp8963 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

																	if True == tmp8963 {
																		__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																		return
																	} else {
																		tmp8962 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5hterm_6, sym_7)

																		if True == tmp8962 {
																			tmp8951 := MakeNative(func(__e *ControlFlow) {
																				News1863 := __e.Get(1)
																				_ = News1863
																				tmp8952 := MakeNative(func(__e *ControlFlow) {
																					Parseshen_4_5end_6 := __e.Get(1)
																					_ = Parseshen_4_5end_6
																					tmp8959 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

																					if True == tmp8959 {
																						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																						return
																					} else {
																						tmp8954 := Call(__e, PrimNS2Value(symshen_4tlstream), V2079)

																						tmp8955 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp8954)

																						tmp8956 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

																						tmp8957 := PrimCons(tmp8956, Nil)

																						tmp8958 := PrimCons(symshen_4_7m, tmp8957)

																						__e.TailApply(PrimNS2Value(symshen_4comb), tmp8955, tmp8958)
																						return

																					}

																				}, 1)

																				tmp8960 := Call(__e, PrimNS2Value(symshen_4_5end_6), News1863)

																				__e.TailApply(tmp8952, tmp8960)
																				return

																			}, 1)

																			tmp8961 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5hterm_6)

																			__e.TailApply(tmp8951, tmp8961)
																			return

																		} else {
																			__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																			return
																		}

																	}

																}, 1)

																tmp8964 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), News1862)

																__e.TailApply(tmp8948, tmp8964)
																return

															}, 1)

															tmp8965 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

															__e.TailApply(tmp8947, tmp8965)
															return

														} else {
															__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
															return
														}

													}, 1)

													tmp8967 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

													tmp8968 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2079)

													tmp8969 := Call(__e, PrimNS2Value(symshen_4comb), tmp8967, tmp8968)

													tmp8970 := Call(__e, tmp8945, tmp8969)

													ifres8943 = tmp8970

												} else {
													tmp8944 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

													ifres8943 = tmp8944

												}

												__e.TailApply(tmp8908, ifres8943)
												return

											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp8997 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2079)

										var ifres8973 Obj

										if True == tmp8997 {
											tmp8975 := MakeNative(func(__e *ControlFlow) {
												SynCons := __e.Get(1)
												_ = SynCons
												tmp8992 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, sym_1)

												if True == tmp8992 {
													tmp8977 := MakeNative(func(__e *ControlFlow) {
														News1861 := __e.Get(1)
														_ = News1861
														tmp8978 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5hterm_6 := __e.Get(1)
															_ = Parseshen_4_5hterm_6
															tmp8989 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

															if True == tmp8989 {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															} else {
																tmp8980 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5end_6 := __e.Get(1)
																	_ = Parseshen_4_5end_6
																	tmp8987 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

																	if True == tmp8987 {
																		__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																		return
																	} else {
																		tmp8982 := Call(__e, PrimNS2Value(symshen_4tlstream), V2079)

																		tmp8983 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp8982)

																		tmp8984 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

																		tmp8985 := PrimCons(tmp8984, Nil)

																		tmp8986 := PrimCons(symshen_4_1m, tmp8985)

																		__e.TailApply(PrimNS2Value(symshen_4comb), tmp8983, tmp8986)
																		return

																	}

																}, 1)

																tmp8988 := Call(__e, PrimNS2Value(symshen_4_5end_6), Parseshen_4_5hterm_6)

																__e.TailApply(tmp8980, tmp8988)
																return

															}

														}, 1)

														tmp8990 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), News1861)

														__e.TailApply(tmp8978, tmp8990)
														return

													}, 1)

													tmp8991 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

													__e.TailApply(tmp8977, tmp8991)
													return

												} else {
													__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp8993 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

											tmp8994 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2079)

											tmp8995 := Call(__e, PrimNS2Value(symshen_4comb), tmp8993, tmp8994)

											tmp8996 := Call(__e, tmp8975, tmp8995)

											ifres8973 = tmp8996

										} else {
											tmp8974 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

											ifres8973 = tmp8974

										}

										__e.TailApply(tmp8906, ifres8973)
										return

									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp9023 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2079)

								var ifres8999 Obj

								if True == tmp9023 {
									tmp9001 := MakeNative(func(__e *ControlFlow) {
										SynCons := __e.Get(1)
										_ = SynCons
										tmp9018 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, sym_7)

										if True == tmp9018 {
											tmp9003 := MakeNative(func(__e *ControlFlow) {
												News1860 := __e.Get(1)
												_ = News1860
												tmp9004 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5hterm_6 := __e.Get(1)
													_ = Parseshen_4_5hterm_6
													tmp9015 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

													if True == tmp9015 {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													} else {
														tmp9006 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5end_6 := __e.Get(1)
															_ = Parseshen_4_5end_6
															tmp9013 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

															if True == tmp9013 {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															} else {
																tmp9008 := Call(__e, PrimNS2Value(symshen_4tlstream), V2079)

																tmp9009 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp9008)

																tmp9010 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

																tmp9011 := PrimCons(tmp9010, Nil)

																tmp9012 := PrimCons(symshen_4_7m, tmp9011)

																__e.TailApply(PrimNS2Value(symshen_4comb), tmp9009, tmp9012)
																return

															}

														}, 1)

														tmp9014 := Call(__e, PrimNS2Value(symshen_4_5end_6), Parseshen_4_5hterm_6)

														__e.TailApply(tmp9006, tmp9014)
														return

													}

												}, 1)

												tmp9016 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), News1860)

												__e.TailApply(tmp9004, tmp9016)
												return

											}, 1)

											tmp9017 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

											__e.TailApply(tmp9003, tmp9017)
											return

										} else {
											__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp9019 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

									tmp9020 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2079)

									tmp9021 := Call(__e, PrimNS2Value(symshen_4comb), tmp9019, tmp9020)

									tmp9022 := Call(__e, tmp9001, tmp9021)

									ifres8999 = tmp9022

								} else {
									tmp9000 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

									ifres8999 = tmp9000

								}

								__e.TailApply(tmp8904, ifres8999)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp9055 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2079)

						var ifres9025 Obj

						if True == tmp9055 {
							tmp9027 := MakeNative(func(__e *ControlFlow) {
								SynCons := __e.Get(1)
								_ = SynCons
								tmp9050 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, symcons)

								if True == tmp9050 {
									tmp9029 := MakeNative(func(__e *ControlFlow) {
										News1859 := __e.Get(1)
										_ = News1859
										tmp9030 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5hterm1_6 := __e.Get(1)
											_ = Parseshen_4_5hterm1_6
											tmp9047 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm1_6)

											if True == tmp9047 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp9032 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5hterm2_6 := __e.Get(1)
													_ = Parseshen_4_5hterm2_6
													tmp9045 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm2_6)

													if True == tmp9045 {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													} else {
														tmp9034 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5end_6 := __e.Get(1)
															_ = Parseshen_4_5end_6
															tmp9043 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

															if True == tmp9043 {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															} else {
																tmp9036 := Call(__e, PrimNS2Value(symshen_4tlstream), V2079)

																tmp9037 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp9036)

																tmp9038 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm1_6)

																tmp9039 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm2_6)

																tmp9040 := PrimCons(tmp9039, Nil)

																tmp9041 := PrimCons(tmp9038, tmp9040)

																tmp9042 := PrimCons(symcons, tmp9041)

																__e.TailApply(PrimNS2Value(symshen_4comb), tmp9037, tmp9042)
																return

															}

														}, 1)

														tmp9044 := Call(__e, PrimNS2Value(symshen_4_5end_6), Parseshen_4_5hterm2_6)

														__e.TailApply(tmp9034, tmp9044)
														return

													}

												}, 1)

												tmp9046 := Call(__e, PrimNS2Value(symshen_4_5hterm2_6), Parseshen_4_5hterm1_6)

												__e.TailApply(tmp9032, tmp9046)
												return

											}

										}, 1)

										tmp9048 := Call(__e, PrimNS2Value(symshen_4_5hterm1_6), News1859)

										__e.TailApply(tmp9030, tmp9048)
										return

									}, 1)

									tmp9049 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

									__e.TailApply(tmp9029, tmp9049)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp9051 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

							tmp9052 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2079)

							tmp9053 := Call(__e, PrimNS2Value(symshen_4comb), tmp9051, tmp9052)

							tmp9054 := Call(__e, tmp9027, tmp9053)

							ifres9025 = tmp9054

						} else {
							tmp9026 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

							ifres9025 = tmp9026

						}

						__e.TailApply(tmp8902, ifres9025)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9068 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2079)

				var ifres9057 Obj

				if True == tmp9068 {
					tmp9059 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp9060 := MakeNative(func(__e *ControlFlow) {
							News1858 := __e.Get(1)
							_ = News1858
							tmp9063 := PrimIntern(MakeString(":"))

							tmp9064 := PrimEqual(X, tmp9063)

							if True == tmp9064 {
								tmp9062 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1858)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp9062, X)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp9065 := Call(__e, PrimNS2Value(symshen_4tls), V2079)

						__e.TailApply(tmp9060, tmp9065)
						return

					}, 1)

					tmp9066 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

					tmp9067 := Call(__e, tmp9059, tmp9066)

					ifres9057 = tmp9067

				} else {
					tmp9058 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

					ifres9057 = tmp9058

				}

				__e.TailApply(tmp8900, ifres9057)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9084 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2079)

		var ifres9070 Obj

		if True == tmp9084 {
			tmp9072 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp9073 := MakeNative(func(__e *ControlFlow) {
					News1857 := __e.Get(1)
					_ = News1857
					tmp9080 := Call(__e, PrimNS2Value(symatom_2), X)

					var ifres9076 Obj

					if True == tmp9080 {
						tmp9078 := Call(__e, PrimNS2Value(symshen_4prolog_1keyword_2), X)

						tmp9079 := PrimNot(tmp9078)

						var ifres9077 Obj

						if True == tmp9079 {
							ifres9077 = True

						} else {
							ifres9077 = False

						}

						ifres9076 = ifres9077

					} else {
						ifres9076 = False

					}

					if True == ifres9076 {
						tmp9075 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1857)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9075, X)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp9081 := Call(__e, PrimNS2Value(symshen_4tls), V2079)

				__e.TailApply(tmp9073, tmp9081)
				return

			}, 1)

			tmp9082 := Call(__e, PrimNS2Value(symshen_4hds), V2079)

			tmp9083 := Call(__e, tmp9072, tmp9082)

			ifres9070 = tmp9083

		} else {
			tmp9071 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres9070 = tmp9071

		}

		__e.TailApply(tmp8898, ifres9070)
		return

	}, 1)

	tmp9085 := Call(__e, PrimNS2Value(symdef), symshen_4_5hterm_6, tmp8897)

	_ = tmp9085

	tmp9086 := MakeNative(func(__e *ControlFlow) {
		V2080 := __e.Get(1)
		_ = V2080
		tmp9087 := PrimIntern(MakeString(";"))

		tmp9088 := PrimCons(sym_5_1_1, Nil)

		tmp9089 := PrimCons(tmp9087, tmp9088)

		__e.TailApply(PrimNS2Value(symelement_2), V2080, tmp9089)
		return

	}, 1)

	tmp9090 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1keyword_2, tmp9086)

	_ = tmp9090

	tmp9091 := MakeNative(func(__e *ControlFlow) {
		V2081 := __e.Get(1)
		_ = V2081
		tmp9104 := PrimIsSymbol(V2081)

		if True == tmp9104 {
			__e.Return(True)
			return
		} else {
			tmp9103 := PrimIsString(V2081)

			var ifres9094 Obj

			if True == tmp9103 {
				ifres9094 = True

			} else {
				tmp9102 := Call(__e, PrimNS2Value(symboolean_2), V2081)

				var ifres9096 Obj

				if True == tmp9102 {
					ifres9096 = True

				} else {
					tmp9101 := PrimIsNumber(V2081)

					var ifres9098 Obj

					if True == tmp9101 {
						ifres9098 = True

					} else {
						tmp9100 := Call(__e, PrimNS2Value(symempty_2), V2081)

						var ifres9099 Obj

						if True == tmp9100 {
							ifres9099 = True

						} else {
							ifres9099 = False

						}

						ifres9098 = ifres9099

					}

					var ifres9097 Obj

					if True == ifres9098 {
						ifres9097 = True

					} else {
						ifres9097 = False

					}

					ifres9096 = ifres9097

				}

				var ifres9095 Obj

				if True == ifres9096 {
					ifres9095 = True

				} else {
					ifres9095 = False

				}

				ifres9094 = ifres9095

			}

			if True == ifres9094 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp9105 := Call(__e, PrimNS2Value(symdef), symatom_2, tmp9091)

	_ = tmp9105

	tmp9106 := MakeNative(func(__e *ControlFlow) {
		V2082 := __e.Get(1)
		_ = V2082
		tmp9107 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9109 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9109 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9110 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5hterm_6 := __e.Get(1)
			_ = Parseshen_4_5hterm_6
			tmp9114 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

			if True == tmp9114 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp9112 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5hterm_6)

				tmp9113 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp9112, tmp9113)
				return

			}

		}, 1)

		tmp9115 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), V2082)

		tmp9116 := Call(__e, tmp9110, tmp9115)

		__e.TailApply(tmp9107, tmp9116)
		return

	}, 1)

	tmp9117 := Call(__e, PrimNS2Value(symdef), symshen_4_5hterm1_6, tmp9106)

	_ = tmp9117

	tmp9118 := MakeNative(func(__e *ControlFlow) {
		V2083 := __e.Get(1)
		_ = V2083
		tmp9119 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9121 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9121 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9122 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5hterm_6 := __e.Get(1)
			_ = Parseshen_4_5hterm_6
			tmp9126 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hterm_6)

			if True == tmp9126 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp9124 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5hterm_6)

				tmp9125 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5hterm_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp9124, tmp9125)
				return

			}

		}, 1)

		tmp9127 := Call(__e, PrimNS2Value(symshen_4_5hterm_6), V2083)

		tmp9128 := Call(__e, tmp9122, tmp9127)

		__e.TailApply(tmp9119, tmp9128)
		return

	}, 1)

	tmp9129 := Call(__e, PrimNS2Value(symdef), symshen_4_5hterm2_6, tmp9118)

	_ = tmp9129

	tmp9130 := MakeNative(func(__e *ControlFlow) {
		V2084 := __e.Get(1)
		_ = V2084
		tmp9131 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9142 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9142 {
				tmp9133 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9135 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp9135 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9136 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp9139 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp9139 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp9138 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9138, Nil)
						return

					}

				}, 1)

				tmp9140 := Call(__e, PrimNS2Value(sym_5e_6), V2084)

				tmp9141 := Call(__e, tmp9136, tmp9140)

				__e.TailApply(tmp9133, tmp9141)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9143 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5literal_6 := __e.Get(1)
			_ = Parseshen_4_5literal_6
			tmp9153 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5literal_6)

			if True == tmp9153 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp9145 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5body_6 := __e.Get(1)
					_ = Parseshen_4_5body_6
					tmp9151 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5body_6)

					if True == tmp9151 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp9147 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5body_6)

						tmp9148 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5literal_6)

						tmp9149 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5body_6)

						tmp9150 := PrimCons(tmp9148, tmp9149)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9147, tmp9150)
						return

					}

				}, 1)

				tmp9152 := Call(__e, PrimNS2Value(symshen_4_5body_6), Parseshen_4_5literal_6)

				__e.TailApply(tmp9145, tmp9152)
				return

			}

		}, 1)

		tmp9154 := Call(__e, PrimNS2Value(symshen_4_5literal_6), V2084)

		tmp9155 := Call(__e, tmp9143, tmp9154)

		__e.TailApply(tmp9131, tmp9155)
		return

	}, 1)

	tmp9156 := Call(__e, PrimNS2Value(symdef), symshen_4_5body_6, tmp9130)

	_ = tmp9156

	tmp9157 := MakeNative(func(__e *ControlFlow) {
		V2085 := __e.Get(1)
		_ = V2085
		tmp9158 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9182 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9182 {
				tmp9160 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9162 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp9162 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9181 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2085)

				var ifres9163 Obj

				if True == tmp9181 {
					tmp9165 := MakeNative(func(__e *ControlFlow) {
						SynCons := __e.Get(1)
						_ = SynCons
						tmp9166 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5bterms_6 := __e.Get(1)
							_ = Parseshen_4_5bterms_6
							tmp9175 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5bterms_6)

							if True == tmp9175 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp9168 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5end_6 := __e.Get(1)
									_ = Parseshen_4_5end_6
									tmp9173 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

									if True == tmp9173 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp9170 := Call(__e, PrimNS2Value(symshen_4tlstream), V2085)

										tmp9171 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp9170)

										tmp9172 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5bterms_6)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp9171, tmp9172)
										return

									}

								}, 1)

								tmp9174 := Call(__e, PrimNS2Value(symshen_4_5end_6), Parseshen_4_5bterms_6)

								__e.TailApply(tmp9168, tmp9174)
								return

							}

						}, 1)

						tmp9176 := Call(__e, PrimNS2Value(symshen_4_5bterms_6), SynCons)

						__e.TailApply(tmp9166, tmp9176)
						return

					}, 1)

					tmp9177 := Call(__e, PrimNS2Value(symshen_4hds), V2085)

					tmp9178 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2085)

					tmp9179 := Call(__e, PrimNS2Value(symshen_4comb), tmp9177, tmp9178)

					tmp9180 := Call(__e, tmp9165, tmp9179)

					ifres9163 = tmp9180

				} else {
					tmp9164 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

					ifres9163 = tmp9164

				}

				__e.TailApply(tmp9160, ifres9163)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9189 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2085, sym_b)

		var ifres9183 Obj

		if True == tmp9189 {
			tmp9185 := MakeNative(func(__e *ControlFlow) {
				News1870 := __e.Get(1)
				_ = News1870
				tmp9186 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1870)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp9186, sym_b)
				return

			}, 1)

			tmp9187 := Call(__e, PrimNS2Value(symshen_4tls), V2085)

			tmp9188 := Call(__e, tmp9185, tmp9187)

			ifres9183 = tmp9188

		} else {
			tmp9184 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres9183 = tmp9184

		}

		__e.TailApply(tmp9158, ifres9183)
		return

	}, 1)

	tmp9190 := Call(__e, PrimNS2Value(symdef), symshen_4_5literal_6, tmp9157)

	_ = tmp9190

	tmp9191 := MakeNative(func(__e *ControlFlow) {
		V2086 := __e.Get(1)
		_ = V2086
		tmp9192 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9203 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9203 {
				tmp9194 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9196 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp9196 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9197 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp9200 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp9200 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp9199 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9199, Nil)
						return

					}

				}, 1)

				tmp9201 := Call(__e, PrimNS2Value(sym_5e_6), V2086)

				tmp9202 := Call(__e, tmp9197, tmp9201)

				__e.TailApply(tmp9194, tmp9202)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9204 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5bterm_6 := __e.Get(1)
			_ = Parseshen_4_5bterm_6
			tmp9214 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5bterm_6)

			if True == tmp9214 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp9206 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5bterms_6 := __e.Get(1)
					_ = Parseshen_4_5bterms_6
					tmp9212 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5bterms_6)

					if True == tmp9212 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp9208 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5bterms_6)

						tmp9209 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5bterm_6)

						tmp9210 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5bterms_6)

						tmp9211 := PrimCons(tmp9209, tmp9210)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9208, tmp9211)
						return

					}

				}, 1)

				tmp9213 := Call(__e, PrimNS2Value(symshen_4_5bterms_6), Parseshen_4_5bterm_6)

				__e.TailApply(tmp9206, tmp9213)
				return

			}

		}, 1)

		tmp9215 := Call(__e, PrimNS2Value(symshen_4_5bterm_6), V2086)

		tmp9216 := Call(__e, tmp9204, tmp9215)

		__e.TailApply(tmp9192, tmp9216)
		return

	}, 1)

	tmp9217 := Call(__e, PrimNS2Value(symdef), symshen_4_5bterms_6, tmp9191)

	_ = tmp9217

	tmp9218 := MakeNative(func(__e *ControlFlow) {
		V2087 := __e.Get(1)
		_ = V2087
		tmp9219 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9257 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9257 {
				tmp9221 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9245 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp9245 {
						tmp9223 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp9225 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp9225 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp9244 := Call(__e, PrimNS2Value(symshen_4ccons_2), V2087)

						var ifres9226 Obj

						if True == tmp9244 {
							tmp9228 := MakeNative(func(__e *ControlFlow) {
								SynCons := __e.Get(1)
								_ = SynCons
								tmp9229 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5bterms_6 := __e.Get(1)
									_ = Parseshen_4_5bterms_6
									tmp9238 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5bterms_6)

									if True == tmp9238 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp9231 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5end_6 := __e.Get(1)
											_ = Parseshen_4_5end_6
											tmp9236 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

											if True == tmp9236 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp9233 := Call(__e, PrimNS2Value(symshen_4tlstream), V2087)

												tmp9234 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp9233)

												tmp9235 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5bterms_6)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp9234, tmp9235)
												return

											}

										}, 1)

										tmp9237 := Call(__e, PrimNS2Value(symshen_4_5end_6), Parseshen_4_5bterms_6)

										__e.TailApply(tmp9231, tmp9237)
										return

									}

								}, 1)

								tmp9239 := Call(__e, PrimNS2Value(symshen_4_5bterms_6), SynCons)

								__e.TailApply(tmp9229, tmp9239)
								return

							}, 1)

							tmp9240 := Call(__e, PrimNS2Value(symshen_4hds), V2087)

							tmp9241 := Call(__e, PrimNS2Value(symshen_4_5_1out), V2087)

							tmp9242 := Call(__e, PrimNS2Value(symshen_4comb), tmp9240, tmp9241)

							tmp9243 := Call(__e, tmp9228, tmp9242)

							ifres9226 = tmp9243

						} else {
							tmp9227 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

							ifres9226 = tmp9227

						}

						__e.TailApply(tmp9223, ifres9226)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp9256 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2087)

				var ifres9246 Obj

				if True == tmp9256 {
					tmp9248 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp9249 := MakeNative(func(__e *ControlFlow) {
							News1873 := __e.Get(1)
							_ = News1873
							tmp9252 := Call(__e, PrimNS2Value(symatom_2), X)

							if True == tmp9252 {
								tmp9251 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1873)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp9251, X)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp9253 := Call(__e, PrimNS2Value(symshen_4tls), V2087)

						__e.TailApply(tmp9249, tmp9253)
						return

					}, 1)

					tmp9254 := Call(__e, PrimNS2Value(symshen_4hds), V2087)

					tmp9255 := Call(__e, tmp9248, tmp9254)

					ifres9246 = tmp9255

				} else {
					tmp9247 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

					ifres9246 = tmp9247

				}

				__e.TailApply(tmp9221, ifres9246)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9258 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5wildcard_6 := __e.Get(1)
			_ = Parseshen_4_5wildcard_6
			tmp9262 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5wildcard_6)

			if True == tmp9262 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp9260 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5wildcard_6)

				tmp9261 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5wildcard_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp9260, tmp9261)
				return

			}

		}, 1)

		tmp9263 := Call(__e, PrimNS2Value(symshen_4_5wildcard_6), V2087)

		tmp9264 := Call(__e, tmp9258, tmp9263)

		__e.TailApply(tmp9219, tmp9264)
		return

	}, 1)

	tmp9265 := Call(__e, PrimNS2Value(symdef), symshen_4_5bterm_6, tmp9218)

	_ = tmp9265

	tmp9266 := MakeNative(func(__e *ControlFlow) {
		V2088 := __e.Get(1)
		_ = V2088
		tmp9267 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9269 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9269 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9282 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2088)

		var ifres9270 Obj

		if True == tmp9282 {
			tmp9272 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp9273 := MakeNative(func(__e *ControlFlow) {
					News1875 := __e.Get(1)
					_ = News1875
					tmp9278 := PrimEqual(X, sym__)

					if True == tmp9278 {
						tmp9275 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1875)

						tmp9276 := Call(__e, PrimNS2Value(symprotect), symY)

						tmp9277 := Call(__e, PrimNS2Value(symgensym), tmp9276)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9275, tmp9277)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp9279 := Call(__e, PrimNS2Value(symshen_4tls), V2088)

				__e.TailApply(tmp9273, tmp9279)
				return

			}, 1)

			tmp9280 := Call(__e, PrimNS2Value(symshen_4hds), V2088)

			tmp9281 := Call(__e, tmp9272, tmp9280)

			ifres9270 = tmp9281

		} else {
			tmp9271 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres9270 = tmp9271

		}

		__e.TailApply(tmp9267, ifres9270)
		return

	}, 1)

	tmp9283 := Call(__e, PrimNS2Value(symdef), symshen_4_5wildcard_6, tmp9266)

	_ = tmp9283

	tmp9284 := MakeNative(func(__e *ControlFlow) {
		V2089 := __e.Get(1)
		_ = V2089
		tmp9285 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9287 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp9287 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9298 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2089)

		var ifres9288 Obj

		if True == tmp9298 {
			tmp9290 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp9291 := MakeNative(func(__e *ControlFlow) {
					News1877 := __e.Get(1)
					_ = News1877
					tmp9294 := Call(__e, PrimNS2Value(symshen_4semicolon_2), X)

					if True == tmp9294 {
						tmp9293 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1877)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp9293, X)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp9295 := Call(__e, PrimNS2Value(symshen_4tls), V2089)

				__e.TailApply(tmp9291, tmp9295)
				return

			}, 1)

			tmp9296 := Call(__e, PrimNS2Value(symshen_4hds), V2089)

			tmp9297 := Call(__e, tmp9290, tmp9296)

			ifres9288 = tmp9297

		} else {
			tmp9289 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres9288 = tmp9289

		}

		__e.TailApply(tmp9285, ifres9288)
		return

	}, 1)

	tmp9299 := Call(__e, PrimNS2Value(symdef), symshen_4_5sc_6, tmp9284)

	_ = tmp9299

	tmp9300 := MakeNative(func(__e *ControlFlow) {
		V2090 := __e.Get(1)
		_ = V2090
		tmp9301 := PrimIntern(MakeString(";"))

		__e.Return(PrimEqual(V2090, tmp9301))
		return

	}, 1)

	tmp9302 := Call(__e, PrimNS2Value(symdef), symshen_4semicolon_2, tmp9300)

	_ = tmp9302

	tmp9303 := MakeNative(func(__e *ControlFlow) {
		V2091 := __e.Get(1)
		_ = V2091
		V2092 := __e.Get(2)
		_ = V2092
		tmp9304 := MakeNative(func(__e *ControlFlow) {
			Bindings := __e.Get(1)
			_ = Bindings
			tmp9305 := MakeNative(func(__e *ControlFlow) {
				Lock := __e.Get(1)
				_ = Lock
				tmp9306 := MakeNative(func(__e *ControlFlow) {
					Key := __e.Get(1)
					_ = Key
					tmp9307 := MakeNative(func(__e *ControlFlow) {
						Continuation := __e.Get(1)
						_ = Continuation
						tmp9308 := MakeNative(func(__e *ControlFlow) {
							Parameters := __e.Get(1)
							_ = Parameters
							tmp9309 := MakeNative(func(__e *ControlFlow) {
								HasCut_2 := __e.Get(1)
								_ = HasCut_2
								tmp9310 := MakeNative(func(__e *ControlFlow) {
									FBody := __e.Get(1)
									_ = FBody
									tmp9311 := MakeNative(func(__e *ControlFlow) {
										CutFBody := __e.Get(1)
										_ = CutFBody
										tmp9312 := MakeNative(func(__e *ControlFlow) {
											Shen := __e.Get(1)
											_ = Shen
											__e.Return(Shen)
											return
										}, 1)

										tmp9313 := PrimCons(sym_1_6, Nil)

										tmp9314 := PrimCons(Continuation, tmp9313)

										tmp9315 := PrimCons(Key, tmp9314)

										tmp9316 := PrimCons(Lock, tmp9315)

										tmp9317 := PrimCons(Bindings, tmp9316)

										tmp9318 := PrimCons(CutFBody, Nil)

										tmp9319 := Call(__e, PrimNS2Value(symappend), tmp9317, tmp9318)

										tmp9320 := Call(__e, PrimNS2Value(symappend), Parameters, tmp9319)

										tmp9321 := PrimCons(V2091, tmp9320)

										tmp9322 := PrimCons(symdefine, tmp9321)

										__e.TailApply(tmp9312, tmp9322)
										return

									}, 1)

									var ifres9323 Obj

									if True == HasCut_2 {
										tmp9324 := PrimCons(MakeNumber(1), Nil)

										tmp9325 := PrimCons(Key, tmp9324)

										tmp9326 := PrimCons(sym_7, tmp9325)

										tmp9327 := PrimCons(FBody, Nil)

										tmp9328 := PrimCons(tmp9326, tmp9327)

										tmp9329 := PrimCons(Key, tmp9328)

										tmp9330 := PrimCons(symlet, tmp9329)

										ifres9323 = tmp9330

									} else {
										ifres9323 = FBody

									}

									__e.TailApply(tmp9311, ifres9323)
									return

								}, 1)

								tmp9331 := Call(__e, PrimNS2Value(symshen_4prolog_1fbody), V2092, Parameters, Bindings, Lock, Key, Continuation, HasCut_2)

								__e.TailApply(tmp9310, tmp9331)
								return

							}, 1)

							tmp9332 := Call(__e, PrimNS2Value(symshen_4hascut_2), V2092)

							__e.TailApply(tmp9309, tmp9332)
							return

						}, 1)

						tmp9333 := Call(__e, PrimNS2Value(symshen_4prolog_1parameters), V2092)

						__e.TailApply(tmp9308, tmp9333)
						return

					}, 1)

					tmp9334 := Call(__e, PrimNS2Value(symprotect), symC)

					tmp9335 := Call(__e, PrimNS2Value(symgensym), tmp9334)

					__e.TailApply(tmp9307, tmp9335)
					return

				}, 1)

				tmp9336 := Call(__e, PrimNS2Value(symprotect), symK)

				tmp9337 := Call(__e, PrimNS2Value(symgensym), tmp9336)

				__e.TailApply(tmp9306, tmp9337)
				return

			}, 1)

			tmp9338 := Call(__e, PrimNS2Value(symprotect), symL)

			tmp9339 := Call(__e, PrimNS2Value(symgensym), tmp9338)

			__e.TailApply(tmp9305, tmp9339)
			return

		}, 1)

		tmp9340 := Call(__e, PrimNS2Value(symprotect), symB)

		tmp9341 := Call(__e, PrimNS2Value(symgensym), tmp9340)

		__e.TailApply(tmp9304, tmp9341)
		return

	}, 2)

	tmp9342 := Call(__e, PrimNS2Value(symdef), symshen_4horn_1clause_1procedure, tmp9303)

	_ = tmp9342

	tmp9343 := MakeNative(func(__e *ControlFlow) {
		V2095 := __e.Get(1)
		_ = V2095
		tmp9353 := PrimEqual(sym_b, V2095)

		if True == tmp9353 {
			__e.Return(True)
			return
		} else {
			tmp9352 := PrimIsPair(V2095)

			if True == tmp9352 {
				tmp9350 := PrimHead(V2095)

				tmp9351 := Call(__e, PrimNS2Value(symshen_4hascut_2), tmp9350)

				if True == tmp9351 {
					__e.Return(True)
					return
				} else {
					tmp9348 := PrimTail(V2095)

					tmp9349 := Call(__e, PrimNS2Value(symshen_4hascut_2), tmp9348)

					if True == tmp9349 {
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

	tmp9354 := Call(__e, PrimNS2Value(symdef), symshen_4hascut_2, tmp9343)

	_ = tmp9354

	tmp9355 := MakeNative(func(__e *ControlFlow) {
		V2100 := __e.Get(1)
		_ = V2100
		tmp9364 := PrimIsPair(V2100)

		var ifres9360 Obj

		if True == tmp9364 {
			tmp9362 := PrimHead(V2100)

			tmp9363 := PrimIsPair(tmp9362)

			var ifres9361 Obj

			if True == tmp9363 {
				ifres9361 = True

			} else {
				ifres9361 = False

			}

			ifres9360 = ifres9361

		} else {
			ifres9360 = False

		}

		if True == ifres9360 {
			tmp9357 := PrimHead(V2100)

			tmp9358 := PrimHead(tmp9357)

			tmp9359 := Call(__e, PrimNS2Value(symlength), tmp9358)

			__e.TailApply(PrimNS2Value(symshen_4parameters), tmp9359)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4prolog_1parameters)
			return
		}

	}, 1)

	tmp9365 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1parameters, tmp9355)

	_ = tmp9365

	tmp9366 := MakeNative(func(__e *ControlFlow) {
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
		tmp9460 := PrimEqual(Nil, V2121)

		var ifres9457 Obj

		if True == tmp9460 {
			tmp9459 := PrimEqual(True, V2127)

			var ifres9458 Obj

			if True == tmp9459 {
				ifres9458 = True

			} else {
				ifres9458 = False

			}

			ifres9457 = ifres9458

		} else {
			ifres9457 = False

		}

		if True == ifres9457 {
			tmp9455 := PrimCons(V2125, Nil)

			tmp9456 := PrimCons(V2124, tmp9455)

			__e.Return(PrimCons(symshen_4unlock, tmp9456))
			return

		} else {
			tmp9454 := PrimIsPair(V2121)

			var ifres9432 Obj

			if True == tmp9454 {
				tmp9452 := PrimHead(V2121)

				tmp9453 := PrimIsPair(tmp9452)

				var ifres9434 Obj

				if True == tmp9453 {
					tmp9449 := PrimHead(V2121)

					tmp9450 := PrimTail(tmp9449)

					tmp9451 := PrimIsPair(tmp9450)

					var ifres9436 Obj

					if True == tmp9451 {
						tmp9445 := PrimHead(V2121)

						tmp9446 := PrimTail(tmp9445)

						tmp9447 := PrimTail(tmp9446)

						tmp9448 := PrimEqual(Nil, tmp9447)

						var ifres9438 Obj

						if True == tmp9448 {
							tmp9443 := PrimTail(V2121)

							tmp9444 := PrimEqual(Nil, tmp9443)

							var ifres9440 Obj

							if True == tmp9444 {
								tmp9442 := PrimEqual(False, V2127)

								var ifres9441 Obj

								if True == tmp9442 {
									ifres9441 = True

								} else {
									ifres9441 = False

								}

								ifres9440 = ifres9441

							} else {
								ifres9440 = False

							}

							var ifres9439 Obj

							if True == ifres9440 {
								ifres9439 = True

							} else {
								ifres9439 = False

							}

							ifres9438 = ifres9439

						} else {
							ifres9438 = False

						}

						var ifres9437 Obj

						if True == ifres9438 {
							ifres9437 = True

						} else {
							ifres9437 = False

						}

						ifres9436 = ifres9437

					} else {
						ifres9436 = False

					}

					var ifres9435 Obj

					if True == ifres9436 {
						ifres9435 = True

					} else {
						ifres9435 = False

					}

					ifres9434 = ifres9435

				} else {
					ifres9434 = False

				}

				var ifres9433 Obj

				if True == ifres9434 {
					ifres9433 = True

				} else {
					ifres9433 = False

				}

				ifres9432 = ifres9433

			} else {
				ifres9432 = False

			}

			if True == ifres9432 {
				tmp9417 := MakeNative(func(__e *ControlFlow) {
					Continue := __e.Get(1)
					_ = Continue
					tmp9418 := PrimCons(V2124, Nil)

					tmp9419 := PrimCons(symshen_4unlocked_2, tmp9418)

					tmp9420 := PrimHead(V2121)

					tmp9421 := PrimHead(tmp9420)

					tmp9422 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_7m, tmp9421, V2122, V2123, Continue)

					tmp9423 := PrimCons(False, Nil)

					tmp9424 := PrimCons(tmp9422, tmp9423)

					tmp9425 := PrimCons(tmp9419, tmp9424)

					__e.Return(PrimCons(symif, tmp9425))
					return

				}, 1)

				tmp9426 := PrimHead(V2121)

				tmp9427 := PrimHead(tmp9426)

				tmp9428 := PrimHead(V2121)

				tmp9429 := PrimTail(tmp9428)

				tmp9430 := PrimHead(tmp9429)

				tmp9431 := Call(__e, PrimNS2Value(symshen_4continue), tmp9427, tmp9430, V2123, V2124, V2125, V2126)

				__e.TailApply(tmp9417, tmp9431)
				return

			} else {
				tmp9416 := PrimIsPair(V2121)

				var ifres9401 Obj

				if True == tmp9416 {
					tmp9414 := PrimHead(V2121)

					tmp9415 := PrimIsPair(tmp9414)

					var ifres9403 Obj

					if True == tmp9415 {
						tmp9411 := PrimHead(V2121)

						tmp9412 := PrimTail(tmp9411)

						tmp9413 := PrimIsPair(tmp9412)

						var ifres9405 Obj

						if True == tmp9413 {
							tmp9407 := PrimHead(V2121)

							tmp9408 := PrimTail(tmp9407)

							tmp9409 := PrimTail(tmp9408)

							tmp9410 := PrimEqual(Nil, tmp9409)

							var ifres9406 Obj

							if True == tmp9410 {
								ifres9406 = True

							} else {
								ifres9406 = False

							}

							ifres9405 = ifres9406

						} else {
							ifres9405 = False

						}

						var ifres9404 Obj

						if True == ifres9405 {
							ifres9404 = True

						} else {
							ifres9404 = False

						}

						ifres9403 = ifres9404

					} else {
						ifres9403 = False

					}

					var ifres9402 Obj

					if True == ifres9403 {
						ifres9402 = True

					} else {
						ifres9402 = False

					}

					ifres9401 = ifres9402

				} else {
					ifres9401 = False

				}

				if True == ifres9401 {
					tmp9370 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp9371 := MakeNative(func(__e *ControlFlow) {
							Continue := __e.Get(1)
							_ = Continue
							tmp9372 := PrimCons(V2124, Nil)

							tmp9373 := PrimCons(symshen_4unlocked_2, tmp9372)

							tmp9374 := PrimHead(V2121)

							tmp9375 := PrimHead(tmp9374)

							tmp9376 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_7m, tmp9375, V2122, V2123, Continue)

							tmp9377 := PrimCons(False, Nil)

							tmp9378 := PrimCons(tmp9376, tmp9377)

							tmp9379 := PrimCons(tmp9373, tmp9378)

							tmp9380 := PrimCons(symif, tmp9379)

							tmp9381 := PrimCons(False, Nil)

							tmp9382 := PrimCons(Case, tmp9381)

							tmp9383 := PrimCons(sym_a, tmp9382)

							tmp9384 := PrimTail(V2121)

							tmp9385 := Call(__e, PrimNS2Value(symshen_4prolog_1fbody), tmp9384, V2122, V2123, V2124, V2125, V2126, V2127)

							tmp9386 := PrimCons(Case, Nil)

							tmp9387 := PrimCons(tmp9385, tmp9386)

							tmp9388 := PrimCons(tmp9383, tmp9387)

							tmp9389 := PrimCons(symif, tmp9388)

							tmp9390 := PrimCons(tmp9389, Nil)

							tmp9391 := PrimCons(tmp9380, tmp9390)

							tmp9392 := PrimCons(Case, tmp9391)

							__e.Return(PrimCons(symlet, tmp9392))
							return

						}, 1)

						tmp9393 := PrimHead(V2121)

						tmp9394 := PrimHead(tmp9393)

						tmp9395 := PrimHead(V2121)

						tmp9396 := PrimTail(tmp9395)

						tmp9397 := PrimHead(tmp9396)

						tmp9398 := Call(__e, PrimNS2Value(symshen_4continue), tmp9394, tmp9397, V2123, V2124, V2125, V2126)

						__e.TailApply(tmp9371, tmp9398)
						return

					}, 1)

					tmp9399 := Call(__e, PrimNS2Value(symgensym), symC)

					tmp9400 := Call(__e, PrimNS2Value(symprotect), tmp9399)

					__e.TailApply(tmp9370, tmp9400)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
					return
				}

			}

		}

	}, 7)

	tmp9461 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1fbody, tmp9366)

	_ = tmp9461

	tmp9462 := MakeNative(func(__e *ControlFlow) {
		V2128 := __e.Get(1)
		_ = V2128
		V2129 := __e.Get(2)
		_ = V2129
		tmp9467 := Call(__e, PrimNS2Value(symshen_4locked_2), V2128)

		var ifres9464 Obj

		if True == tmp9467 {
			tmp9466 := Call(__e, PrimNS2Value(symshen_4fits_2), V2129, V2128)

			var ifres9465 Obj

			if True == tmp9466 {
				ifres9465 = True

			} else {
				ifres9465 = False

			}

			ifres9464 = ifres9465

		} else {
			ifres9464 = False

		}

		if True == ifres9464 {
			__e.TailApply(PrimNS2Value(symshen_4openlock), V2128)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp9468 := Call(__e, PrimNS2Value(symdef), symshen_4unlock, tmp9462)

	_ = tmp9468

	tmp9469 := MakeNative(func(__e *ControlFlow) {
		V2130 := __e.Get(1)
		_ = V2130
		tmp9470 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V2130)

		__e.Return(PrimNot(tmp9470))
		return

	}, 1)

	tmp9471 := Call(__e, PrimNS2Value(symdef), symshen_4locked_2, tmp9469)

	_ = tmp9471

	tmp9472 := MakeNative(func(__e *ControlFlow) {
		V2131 := __e.Get(1)
		_ = V2131
		__e.Return(PrimVectorGet(V2131, MakeNumber(1)))
		return
	}, 1)

	tmp9473 := Call(__e, PrimNS2Value(symdef), symshen_4unlocked_2, tmp9472)

	_ = tmp9473

	tmp9474 := MakeNative(func(__e *ControlFlow) {
		V2132 := __e.Get(1)
		_ = V2132
		tmp9475 := PrimVectorSet(V2132, MakeNumber(1), True)

		_ = tmp9475

		__e.Return(False)
		return

	}, 1)

	tmp9476 := Call(__e, PrimNS2Value(symdef), symshen_4openlock, tmp9474)

	_ = tmp9476

	tmp9477 := MakeNative(func(__e *ControlFlow) {
		V2133 := __e.Get(1)
		_ = V2133
		V2134 := __e.Get(2)
		_ = V2134
		tmp9478 := PrimVectorGet(V2134, MakeNumber(2))

		__e.Return(PrimEqual(V2133, tmp9478))
		return

	}, 2)

	tmp9479 := Call(__e, PrimNS2Value(symdef), symshen_4fits_2, tmp9477)

	_ = tmp9479

	tmp9480 := MakeNative(func(__e *ControlFlow) {
		V2137 := __e.Get(1)
		_ = V2137
		V2138 := __e.Get(2)
		_ = V2138
		V2139 := __e.Get(3)
		_ = V2139
		V2140 := __e.Get(4)
		_ = V2140
		tmp9481 := MakeNative(func(__e *ControlFlow) {
			Compute := __e.Get(1)
			_ = Compute
			tmp9486 := PrimEqual(Compute, False)

			var ifres9483 Obj

			if True == tmp9486 {
				tmp9485 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V2138)

				var ifres9484 Obj

				if True == tmp9485 {
					ifres9484 = True

				} else {
					ifres9484 = False

				}

				ifres9483 = ifres9484

			} else {
				ifres9483 = False

			}

			if True == ifres9483 {
				__e.TailApply(PrimNS2Value(symshen_4lock), V2139, V2138)
				return
			} else {
				__e.Return(Compute)
				return
			}

		}, 1)

		tmp9487 := Call(__e, PrimNS2Value(symthaw), V2140)

		__e.TailApply(tmp9481, tmp9487)
		return

	}, 4)

	tmp9488 := Call(__e, PrimNS2Value(symdef), symshen_4cut, tmp9480)

	_ = tmp9488

	tmp9489 := MakeNative(func(__e *ControlFlow) {
		V2141 := __e.Get(1)
		_ = V2141
		V2142 := __e.Get(2)
		_ = V2142
		tmp9490 := MakeNative(func(__e *ControlFlow) {
			SetLock := __e.Get(1)
			_ = SetLock
			tmp9491 := MakeNative(func(__e *ControlFlow) {
				SetKey := __e.Get(1)
				_ = SetKey
				__e.Return(False)
				return
			}, 1)

			tmp9492 := PrimVectorSet(V2142, MakeNumber(2), V2141)

			__e.TailApply(tmp9491, tmp9492)
			return

		}, 1)

		tmp9493 := PrimVectorSet(V2142, MakeNumber(1), False)

		__e.TailApply(tmp9490, tmp9493)
		return

	}, 2)

	tmp9494 := Call(__e, PrimNS2Value(symdef), symshen_4lock, tmp9489)

	_ = tmp9494

	tmp9495 := MakeNative(func(__e *ControlFlow) {
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
		tmp9496 := MakeNative(func(__e *ControlFlow) {
			HVs := __e.Get(1)
			_ = HVs
			tmp9497 := MakeNative(func(__e *ControlFlow) {
				BVs := __e.Get(1)
				_ = BVs
				tmp9498 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					tmp9499 := MakeNative(func(__e *ControlFlow) {
						ContinuationCode := __e.Get(1)
						_ = ContinuationCode
						__e.TailApply(PrimNS2Value(symshen_4stpart), Free, ContinuationCode, V2145)
						return
					}, 1)

					tmp9500 := PrimCons(symshen_4incinfs, Nil)

					tmp9501 := Call(__e, PrimNS2Value(symshen_4compile_1body), V2144, V2145, V2146, V2147, V2148)

					tmp9502 := PrimCons(tmp9501, Nil)

					tmp9503 := PrimCons(tmp9500, tmp9502)

					tmp9504 := PrimCons(symdo, tmp9503)

					__e.TailApply(tmp9499, tmp9504)
					return

				}, 1)

				tmp9505 := Call(__e, PrimNS2Value(symdifference), BVs, HVs)

				__e.TailApply(tmp9498, tmp9505)
				return

			}, 1)

			tmp9506 := Call(__e, PrimNS2Value(symshen_4extract_1vars), V2144)

			__e.TailApply(tmp9497, tmp9506)
			return

		}, 1)

		tmp9507 := Call(__e, PrimNS2Value(symshen_4extract_1vars), V2143)

		__e.TailApply(tmp9496, tmp9507)
		return

	}, 6)

	tmp9508 := Call(__e, PrimNS2Value(symdef), symshen_4continue, tmp9495)

	_ = tmp9508

	tmp9509 := MakeNative(func(__e *ControlFlow) {
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
		tmp9544 := PrimEqual(Nil, V2165)

		if True == tmp9544 {
			tmp9543 := PrimCons(V2169, Nil)

			__e.Return(PrimCons(symthaw, tmp9543))
			return

		} else {
			tmp9542 := PrimIsPair(V2165)

			var ifres9538 Obj

			if True == tmp9542 {
				tmp9540 := PrimHead(V2165)

				tmp9541 := PrimEqual(sym_b, tmp9540)

				var ifres9539 Obj

				if True == tmp9541 {
					ifres9539 = True

				} else {
					ifres9539 = False

				}

				ifres9538 = ifres9539

			} else {
				ifres9538 = False

			}

			if True == ifres9538 {
				tmp9535 := PrimCons(symshen_4cut, Nil)

				tmp9536 := PrimTail(V2165)

				tmp9537 := PrimCons(tmp9535, tmp9536)

				__e.TailApply(PrimNS2Value(symshen_4compile_1body), tmp9537, V2166, V2167, V2168, V2169)
				return

			} else {
				tmp9534 := PrimIsPair(V2165)

				var ifres9530 Obj

				if True == tmp9534 {
					tmp9532 := PrimTail(V2165)

					tmp9533 := PrimEqual(Nil, tmp9532)

					var ifres9531 Obj

					if True == tmp9533 {
						ifres9531 = True

					} else {
						ifres9531 = False

					}

					ifres9530 = ifres9531

				} else {
					ifres9530 = False

				}

				if True == ifres9530 {
					tmp9524 := PrimHead(V2165)

					tmp9525 := Call(__e, PrimNS2Value(symshen_4deref_1calls), tmp9524, V2166)

					tmp9526 := PrimCons(V2169, Nil)

					tmp9527 := PrimCons(V2168, tmp9526)

					tmp9528 := PrimCons(V2167, tmp9527)

					tmp9529 := PrimCons(V2166, tmp9528)

					__e.TailApply(PrimNS2Value(symappend), tmp9525, tmp9529)
					return

				} else {
					tmp9523 := PrimIsPair(V2165)

					if True == tmp9523 {
						tmp9514 := MakeNative(func(__e *ControlFlow) {
							P_d := __e.Get(1)
							_ = P_d
							tmp9515 := PrimTail(V2165)

							tmp9516 := Call(__e, PrimNS2Value(symshen_4freeze_1literals), tmp9515, V2166, V2167, V2168, V2169)

							tmp9517 := PrimCons(tmp9516, Nil)

							tmp9518 := PrimCons(V2168, tmp9517)

							tmp9519 := PrimCons(V2167, tmp9518)

							tmp9520 := PrimCons(V2166, tmp9519)

							__e.TailApply(PrimNS2Value(symappend), P_d, tmp9520)
							return

						}, 1)

						tmp9521 := PrimHead(V2165)

						tmp9522 := Call(__e, PrimNS2Value(symshen_4deref_1calls), tmp9521, V2166)

						__e.TailApply(tmp9514, tmp9522)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
						return
					}

				}

			}

		}

	}, 5)

	tmp9545 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1body, tmp9509)

	_ = tmp9545

	tmp9546 := MakeNative(func(__e *ControlFlow) {
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
		tmp9570 := PrimEqual(Nil, V2186)

		if True == tmp9570 {
			__e.Return(V2190)
			return
		} else {
			tmp9569 := PrimIsPair(V2186)

			var ifres9565 Obj

			if True == tmp9569 {
				tmp9567 := PrimHead(V2186)

				tmp9568 := PrimEqual(sym_b, tmp9567)

				var ifres9566 Obj

				if True == tmp9568 {
					ifres9566 = True

				} else {
					ifres9566 = False

				}

				ifres9565 = ifres9566

			} else {
				ifres9565 = False

			}

			if True == ifres9565 {
				tmp9562 := PrimCons(symshen_4cut, Nil)

				tmp9563 := PrimTail(V2186)

				tmp9564 := PrimCons(tmp9562, tmp9563)

				__e.TailApply(PrimNS2Value(symshen_4freeze_1literals), tmp9564, V2187, V2188, V2189, V2190)
				return

			} else {
				tmp9561 := PrimIsPair(V2186)

				if True == tmp9561 {
					tmp9550 := MakeNative(func(__e *ControlFlow) {
						P_d := __e.Get(1)
						_ = P_d
						tmp9551 := PrimTail(V2186)

						tmp9552 := Call(__e, PrimNS2Value(symshen_4freeze_1literals), tmp9551, V2187, V2188, V2189, V2190)

						tmp9553 := PrimCons(tmp9552, Nil)

						tmp9554 := PrimCons(V2189, tmp9553)

						tmp9555 := PrimCons(V2188, tmp9554)

						tmp9556 := PrimCons(V2187, tmp9555)

						tmp9557 := Call(__e, PrimNS2Value(symappend), P_d, tmp9556)

						tmp9558 := PrimCons(tmp9557, Nil)

						__e.Return(PrimCons(symfreeze, tmp9558))
						return

					}, 1)

					tmp9559 := PrimHead(V2186)

					tmp9560 := Call(__e, PrimNS2Value(symshen_4deref_1calls), tmp9559, V2187)

					__e.TailApply(tmp9550, tmp9560)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
					return
				}

			}

		}

	}, 5)

	tmp9571 := Call(__e, PrimNS2Value(symdef), symshen_4freeze_1literals, tmp9546)

	_ = tmp9571

	tmp9572 := MakeNative(func(__e *ControlFlow) {
		V2195 := __e.Get(1)
		_ = V2195
		V2196 := __e.Get(2)
		_ = V2196
		tmp9587 := PrimIsPair(V2195)

		var ifres9583 Obj

		if True == tmp9587 {
			tmp9585 := PrimHead(V2195)

			tmp9586 := PrimEqual(symfork, tmp9585)

			var ifres9584 Obj

			if True == tmp9586 {
				ifres9584 = True

			} else {
				ifres9584 = False

			}

			ifres9583 = ifres9584

		} else {
			ifres9583 = False

		}

		if True == ifres9583 {
			tmp9580 := PrimTail(V2195)

			tmp9581 := Call(__e, PrimNS2Value(symshen_4deref_1forked_1literals), tmp9580, V2196)

			tmp9582 := PrimCons(tmp9581, Nil)

			__e.Return(PrimCons(symfork, tmp9582))
			return

		} else {
			tmp9579 := PrimIsPair(V2195)

			if True == tmp9579 {
				tmp9575 := PrimHead(V2195)

				tmp9576 := MakeNative(func(__e *ControlFlow) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(PrimNS2Value(symshen_4function_1calls), Y, V2196)
					return
				}, 1)

				tmp9577 := PrimTail(V2195)

				tmp9578 := Call(__e, PrimNS2Value(symmap), tmp9576, tmp9577)

				__e.Return(PrimCons(tmp9575, tmp9578))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
				return
			}

		}

	}, 2)

	tmp9588 := Call(__e, PrimNS2Value(symdef), symshen_4deref_1calls, tmp9572)

	_ = tmp9588

	tmp9589 := MakeNative(func(__e *ControlFlow) {
		V2203 := __e.Get(1)
		_ = V2203
		V2204 := __e.Get(2)
		_ = V2204
		tmp9599 := PrimEqual(Nil, V2203)

		if True == tmp9599 {
			__e.Return(Nil)
			return
		} else {
			tmp9598 := PrimIsPair(V2203)

			if True == tmp9598 {
				tmp9592 := PrimHead(V2203)

				tmp9593 := Call(__e, PrimNS2Value(symshen_4deref_1calls), tmp9592, V2204)

				tmp9594 := PrimTail(V2203)

				tmp9595 := Call(__e, PrimNS2Value(symshen_4deref_1forked_1literals), tmp9594, V2204)

				tmp9596 := PrimCons(tmp9595, Nil)

				tmp9597 := PrimCons(tmp9593, tmp9596)

				__e.Return(PrimCons(symcons, tmp9597))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
				return
			}

		}

	}, 2)

	tmp9600 := Call(__e, PrimNS2Value(symdef), symshen_4deref_1forked_1literals, tmp9589)

	_ = tmp9600

	tmp9601 := MakeNative(func(__e *ControlFlow) {
		V2207 := __e.Get(1)
		_ = V2207
		V2208 := __e.Get(2)
		_ = V2208
		tmp9633 := PrimIsPair(V2207)

		var ifres9614 Obj

		if True == tmp9633 {
			tmp9631 := PrimHead(V2207)

			tmp9632 := PrimEqual(symcons, tmp9631)

			var ifres9616 Obj

			if True == tmp9632 {
				tmp9629 := PrimTail(V2207)

				tmp9630 := PrimIsPair(tmp9629)

				var ifres9618 Obj

				if True == tmp9630 {
					tmp9626 := PrimTail(V2207)

					tmp9627 := PrimTail(tmp9626)

					tmp9628 := PrimIsPair(tmp9627)

					var ifres9620 Obj

					if True == tmp9628 {
						tmp9622 := PrimTail(V2207)

						tmp9623 := PrimTail(tmp9622)

						tmp9624 := PrimTail(tmp9623)

						tmp9625 := PrimEqual(Nil, tmp9624)

						var ifres9621 Obj

						if True == tmp9625 {
							ifres9621 = True

						} else {
							ifres9621 = False

						}

						ifres9620 = ifres9621

					} else {
						ifres9620 = False

					}

					var ifres9619 Obj

					if True == ifres9620 {
						ifres9619 = True

					} else {
						ifres9619 = False

					}

					ifres9618 = ifres9619

				} else {
					ifres9618 = False

				}

				var ifres9617 Obj

				if True == ifres9618 {
					ifres9617 = True

				} else {
					ifres9617 = False

				}

				ifres9616 = ifres9617

			} else {
				ifres9616 = False

			}

			var ifres9615 Obj

			if True == ifres9616 {
				ifres9615 = True

			} else {
				ifres9615 = False

			}

			ifres9614 = ifres9615

		} else {
			ifres9614 = False

		}

		if True == ifres9614 {
			tmp9605 := PrimTail(V2207)

			tmp9606 := PrimHead(tmp9605)

			tmp9607 := Call(__e, PrimNS2Value(symshen_4function_1calls), tmp9606, V2208)

			tmp9608 := PrimTail(V2207)

			tmp9609 := PrimTail(tmp9608)

			tmp9610 := PrimHead(tmp9609)

			tmp9611 := Call(__e, PrimNS2Value(symshen_4function_1calls), tmp9610, V2208)

			tmp9612 := PrimCons(tmp9611, Nil)

			tmp9613 := PrimCons(tmp9607, tmp9612)

			__e.Return(PrimCons(symcons, tmp9613))
			return

		} else {
			tmp9604 := PrimIsPair(V2207)

			if True == tmp9604 {
				__e.TailApply(PrimNS2Value(symshen_4deref_1terms), V2207, V2208)
				return
			} else {
				__e.Return(V2207)
				return
			}

		}

	}, 2)

	tmp9634 := Call(__e, PrimNS2Value(symdef), symshen_4function_1calls, tmp9601)

	_ = tmp9634

	tmp9635 := MakeNative(func(__e *ControlFlow) {
		V2213 := __e.Get(1)
		_ = V2213
		V2214 := __e.Get(2)
		_ = V2214
		tmp9693 := PrimIsPair(V2213)

		var ifres9680 Obj

		if True == tmp9693 {
			tmp9691 := PrimHead(V2213)

			tmp9692 := PrimEqual(MakeNumber(0), tmp9691)

			var ifres9682 Obj

			if True == tmp9692 {
				tmp9689 := PrimTail(V2213)

				tmp9690 := PrimIsPair(tmp9689)

				var ifres9684 Obj

				if True == tmp9690 {
					tmp9686 := PrimTail(V2213)

					tmp9687 := PrimTail(tmp9686)

					tmp9688 := PrimEqual(Nil, tmp9687)

					var ifres9685 Obj

					if True == tmp9688 {
						ifres9685 = True

					} else {
						ifres9685 = False

					}

					ifres9684 = ifres9685

				} else {
					ifres9684 = False

				}

				var ifres9683 Obj

				if True == ifres9684 {
					ifres9683 = True

				} else {
					ifres9683 = False

				}

				ifres9682 = ifres9683

			} else {
				ifres9682 = False

			}

			var ifres9681 Obj

			if True == ifres9682 {
				ifres9681 = True

			} else {
				ifres9681 = False

			}

			ifres9680 = ifres9681

		} else {
			ifres9680 = False

		}

		if True == ifres9680 {
			tmp9677 := PrimTail(V2213)

			tmp9678 := PrimHead(tmp9677)

			tmp9679 := PrimIsVariable(tmp9678)

			if True == tmp9679 {
				tmp9676 := PrimTail(V2213)

				__e.Return(PrimHead(tmp9676))
				return

			} else {
				tmp9672 := PrimTail(V2213)

				tmp9673 := PrimHead(tmp9672)

				tmp9674 := Call(__e, PrimNS2Value(symshen_4app), tmp9673, MakeString("\n"), symshen_4s)

				tmp9675 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp9674)

				__e.Return(PrimSimpleError(tmp9675))
				return

			}

		} else {
			tmp9670 := PrimIsPair(V2213)

			var ifres9657 Obj

			if True == tmp9670 {
				tmp9668 := PrimHead(V2213)

				tmp9669 := PrimEqual(MakeNumber(1), tmp9668)

				var ifres9659 Obj

				if True == tmp9669 {
					tmp9666 := PrimTail(V2213)

					tmp9667 := PrimIsPair(tmp9666)

					var ifres9661 Obj

					if True == tmp9667 {
						tmp9663 := PrimTail(V2213)

						tmp9664 := PrimTail(tmp9663)

						tmp9665 := PrimEqual(Nil, tmp9664)

						var ifres9662 Obj

						if True == tmp9665 {
							ifres9662 = True

						} else {
							ifres9662 = False

						}

						ifres9661 = ifres9662

					} else {
						ifres9661 = False

					}

					var ifres9660 Obj

					if True == ifres9661 {
						ifres9660 = True

					} else {
						ifres9660 = False

					}

					ifres9659 = ifres9660

				} else {
					ifres9659 = False

				}

				var ifres9658 Obj

				if True == ifres9659 {
					ifres9658 = True

				} else {
					ifres9658 = False

				}

				ifres9657 = ifres9658

			} else {
				ifres9657 = False

			}

			if True == ifres9657 {
				tmp9654 := PrimTail(V2213)

				tmp9655 := PrimHead(tmp9654)

				tmp9656 := PrimIsVariable(tmp9655)

				if True == tmp9656 {
					tmp9650 := PrimTail(V2213)

					tmp9651 := PrimHead(tmp9650)

					tmp9652 := PrimCons(V2214, Nil)

					tmp9653 := PrimCons(tmp9651, tmp9652)

					__e.Return(PrimCons(symshen_4lazyderef, tmp9653))
					return

				} else {
					tmp9646 := PrimTail(V2213)

					tmp9647 := PrimHead(tmp9646)

					tmp9648 := Call(__e, PrimNS2Value(symshen_4app), tmp9647, MakeString("\n"), symshen_4s)

					tmp9649 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp9648)

					__e.Return(PrimSimpleError(tmp9649))
					return

				}

			} else {
				tmp9644 := PrimIsVariable(V2213)

				if True == tmp9644 {
					tmp9642 := PrimCons(V2214, Nil)

					tmp9643 := PrimCons(V2213, tmp9642)

					__e.Return(PrimCons(symshen_4deref, tmp9643))
					return

				} else {
					tmp9641 := PrimIsPair(V2213)

					if True == tmp9641 {
						tmp9640 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(PrimNS2Value(symshen_4deref_1terms), Z, V2214)
							return
						}, 1)

						__e.TailApply(PrimNS2Value(symmap), tmp9640, V2213)
						return

					} else {
						__e.Return(V2213)
						return
					}

				}

			}

		}

	}, 2)

	tmp9694 := Call(__e, PrimNS2Value(symdef), symshen_4deref_1terms, tmp9635)

	_ = tmp9694

	tmp9695 := MakeNative(func(__e *ControlFlow) {
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
		tmp9871 := PrimEqual(Nil, V2233)

		var ifres9868 Obj

		if True == tmp9871 {
			tmp9870 := PrimEqual(Nil, V2234)

			var ifres9869 Obj

			if True == tmp9870 {
				ifres9869 = True

			} else {
				ifres9869 = False

			}

			ifres9868 = ifres9869

		} else {
			ifres9868 = False

		}

		if True == ifres9868 {
			__e.Return(V2236)
			return
		} else {
			tmp9867 := PrimIsPair(V2233)

			var ifres9847 Obj

			if True == tmp9867 {
				tmp9865 := PrimHead(V2233)

				tmp9866 := PrimIsPair(tmp9865)

				var ifres9849 Obj

				if True == tmp9866 {
					tmp9862 := PrimHead(V2233)

					tmp9863 := PrimHead(tmp9862)

					tmp9864 := PrimEqual(symshen_4_7m, tmp9863)

					var ifres9851 Obj

					if True == tmp9864 {
						tmp9859 := PrimHead(V2233)

						tmp9860 := PrimTail(tmp9859)

						tmp9861 := PrimIsPair(tmp9860)

						var ifres9853 Obj

						if True == tmp9861 {
							tmp9855 := PrimHead(V2233)

							tmp9856 := PrimTail(tmp9855)

							tmp9857 := PrimTail(tmp9856)

							tmp9858 := PrimEqual(Nil, tmp9857)

							var ifres9854 Obj

							if True == tmp9858 {
								ifres9854 = True

							} else {
								ifres9854 = False

							}

							ifres9853 = ifres9854

						} else {
							ifres9853 = False

						}

						var ifres9852 Obj

						if True == ifres9853 {
							ifres9852 = True

						} else {
							ifres9852 = False

						}

						ifres9851 = ifres9852

					} else {
						ifres9851 = False

					}

					var ifres9850 Obj

					if True == ifres9851 {
						ifres9850 = True

					} else {
						ifres9850 = False

					}

					ifres9849 = ifres9850

				} else {
					ifres9849 = False

				}

				var ifres9848 Obj

				if True == ifres9849 {
					ifres9848 = True

				} else {
					ifres9848 = False

				}

				ifres9847 = ifres9848

			} else {
				ifres9847 = False

			}

			if True == ifres9847 {
				tmp9840 := PrimHead(V2233)

				tmp9841 := PrimTail(tmp9840)

				tmp9842 := PrimHead(tmp9841)

				tmp9843 := PrimTail(V2233)

				tmp9844 := PrimCons(V2232, tmp9843)

				tmp9845 := PrimCons(tmp9842, tmp9844)

				tmp9846 := PrimCons(symshen_4_7m, tmp9845)

				__e.TailApply(PrimNS2Value(symshen_4compile_1head), V2232, tmp9846, V2234, V2235, V2236)
				return

			} else {
				tmp9839 := PrimIsPair(V2233)

				var ifres9819 Obj

				if True == tmp9839 {
					tmp9837 := PrimHead(V2233)

					tmp9838 := PrimIsPair(tmp9837)

					var ifres9821 Obj

					if True == tmp9838 {
						tmp9834 := PrimHead(V2233)

						tmp9835 := PrimHead(tmp9834)

						tmp9836 := PrimEqual(symshen_4_1m, tmp9835)

						var ifres9823 Obj

						if True == tmp9836 {
							tmp9831 := PrimHead(V2233)

							tmp9832 := PrimTail(tmp9831)

							tmp9833 := PrimIsPair(tmp9832)

							var ifres9825 Obj

							if True == tmp9833 {
								tmp9827 := PrimHead(V2233)

								tmp9828 := PrimTail(tmp9827)

								tmp9829 := PrimTail(tmp9828)

								tmp9830 := PrimEqual(Nil, tmp9829)

								var ifres9826 Obj

								if True == tmp9830 {
									ifres9826 = True

								} else {
									ifres9826 = False

								}

								ifres9825 = ifres9826

							} else {
								ifres9825 = False

							}

							var ifres9824 Obj

							if True == ifres9825 {
								ifres9824 = True

							} else {
								ifres9824 = False

							}

							ifres9823 = ifres9824

						} else {
							ifres9823 = False

						}

						var ifres9822 Obj

						if True == ifres9823 {
							ifres9822 = True

						} else {
							ifres9822 = False

						}

						ifres9821 = ifres9822

					} else {
						ifres9821 = False

					}

					var ifres9820 Obj

					if True == ifres9821 {
						ifres9820 = True

					} else {
						ifres9820 = False

					}

					ifres9819 = ifres9820

				} else {
					ifres9819 = False

				}

				if True == ifres9819 {
					tmp9812 := PrimHead(V2233)

					tmp9813 := PrimTail(tmp9812)

					tmp9814 := PrimHead(tmp9813)

					tmp9815 := PrimTail(V2233)

					tmp9816 := PrimCons(V2232, tmp9815)

					tmp9817 := PrimCons(tmp9814, tmp9816)

					tmp9818 := PrimCons(symshen_4_1m, tmp9817)

					__e.TailApply(PrimNS2Value(symshen_4compile_1head), V2232, tmp9818, V2234, V2235, V2236)
					return

				} else {
					tmp9811 := PrimIsPair(V2233)

					var ifres9807 Obj

					if True == tmp9811 {
						tmp9809 := PrimHead(V2233)

						tmp9810 := PrimEqual(symshen_4_1m, tmp9809)

						var ifres9808 Obj

						if True == tmp9810 {
							ifres9808 = True

						} else {
							ifres9808 = False

						}

						ifres9807 = ifres9808

					} else {
						ifres9807 = False

					}

					if True == ifres9807 {
						tmp9806 := PrimTail(V2233)

						__e.TailApply(PrimNS2Value(symshen_4compile_1head), symshen_4_1m, tmp9806, V2234, V2235, V2236)
						return

					} else {
						tmp9805 := PrimIsPair(V2233)

						var ifres9801 Obj

						if True == tmp9805 {
							tmp9803 := PrimHead(V2233)

							tmp9804 := PrimEqual(symshen_4_7m, tmp9803)

							var ifres9802 Obj

							if True == tmp9804 {
								ifres9802 = True

							} else {
								ifres9802 = False

							}

							ifres9801 = ifres9802

						} else {
							ifres9801 = False

						}

						if True == ifres9801 {
							tmp9800 := PrimTail(V2233)

							__e.TailApply(PrimNS2Value(symshen_4compile_1head), symshen_4_7m, tmp9800, V2234, V2235, V2236)
							return

						} else {
							tmp9799 := PrimIsPair(V2233)

							var ifres9792 Obj

							if True == tmp9799 {
								tmp9798 := PrimIsPair(V2234)

								var ifres9794 Obj

								if True == tmp9798 {
									tmp9796 := PrimHead(V2233)

									tmp9797 := Call(__e, PrimNS2Value(symshen_4wildcard_2), tmp9796)

									var ifres9795 Obj

									if True == tmp9797 {
										ifres9795 = True

									} else {
										ifres9795 = False

									}

									ifres9794 = ifres9795

								} else {
									ifres9794 = False

								}

								var ifres9793 Obj

								if True == ifres9794 {
									ifres9793 = True

								} else {
									ifres9793 = False

								}

								ifres9792 = ifres9793

							} else {
								ifres9792 = False

							}

							if True == ifres9792 {
								tmp9790 := PrimTail(V2233)

								tmp9791 := PrimTail(V2234)

								__e.TailApply(PrimNS2Value(symshen_4compile_1head), V2232, tmp9790, tmp9791, V2235, V2236)
								return

							} else {
								tmp9789 := PrimIsPair(V2233)

								var ifres9785 Obj

								if True == tmp9789 {
									tmp9787 := PrimHead(V2233)

									tmp9788 := PrimIsVariable(tmp9787)

									var ifres9786 Obj

									if True == tmp9788 {
										ifres9786 = True

									} else {
										ifres9786 = False

									}

									ifres9785 = ifres9786

								} else {
									ifres9785 = False

								}

								if True == ifres9785 {
									__e.TailApply(PrimNS2Value(symshen_4variable_1case), V2232, V2233, V2234, V2235, V2236)
									return
								} else {
									tmp9784 := PrimEqual(symshen_4_1m, V2232)

									var ifres9777 Obj

									if True == tmp9784 {
										tmp9783 := PrimIsPair(V2233)

										var ifres9779 Obj

										if True == tmp9783 {
											tmp9781 := PrimHead(V2233)

											tmp9782 := Call(__e, PrimNS2Value(symatom_2), tmp9781)

											var ifres9780 Obj

											if True == tmp9782 {
												ifres9780 = True

											} else {
												ifres9780 = False

											}

											ifres9779 = ifres9780

										} else {
											ifres9779 = False

										}

										var ifres9778 Obj

										if True == ifres9779 {
											ifres9778 = True

										} else {
											ifres9778 = False

										}

										ifres9777 = ifres9778

									} else {
										ifres9777 = False

									}

									if True == ifres9777 {
										__e.TailApply(PrimNS2Value(symshen_4atom_1case_1minus), V2233, V2234, V2235, V2236)
										return
									} else {
										tmp9776 := PrimEqual(symshen_4_1m, V2232)

										var ifres9746 Obj

										if True == tmp9776 {
											tmp9775 := PrimIsPair(V2233)

											var ifres9748 Obj

											if True == tmp9775 {
												tmp9773 := PrimHead(V2233)

												tmp9774 := PrimIsPair(tmp9773)

												var ifres9750 Obj

												if True == tmp9774 {
													tmp9770 := PrimHead(V2233)

													tmp9771 := PrimHead(tmp9770)

													tmp9772 := PrimEqual(symcons, tmp9771)

													var ifres9752 Obj

													if True == tmp9772 {
														tmp9767 := PrimHead(V2233)

														tmp9768 := PrimTail(tmp9767)

														tmp9769 := PrimIsPair(tmp9768)

														var ifres9754 Obj

														if True == tmp9769 {
															tmp9763 := PrimHead(V2233)

															tmp9764 := PrimTail(tmp9763)

															tmp9765 := PrimTail(tmp9764)

															tmp9766 := PrimIsPair(tmp9765)

															var ifres9756 Obj

															if True == tmp9766 {
																tmp9758 := PrimHead(V2233)

																tmp9759 := PrimTail(tmp9758)

																tmp9760 := PrimTail(tmp9759)

																tmp9761 := PrimTail(tmp9760)

																tmp9762 := PrimEqual(Nil, tmp9761)

																var ifres9757 Obj

																if True == tmp9762 {
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

												var ifres9749 Obj

												if True == ifres9750 {
													ifres9749 = True

												} else {
													ifres9749 = False

												}

												ifres9748 = ifres9749

											} else {
												ifres9748 = False

											}

											var ifres9747 Obj

											if True == ifres9748 {
												ifres9747 = True

											} else {
												ifres9747 = False

											}

											ifres9746 = ifres9747

										} else {
											ifres9746 = False

										}

										if True == ifres9746 {
											__e.TailApply(PrimNS2Value(symshen_4cons_1case_1minus), V2233, V2234, V2235, V2236)
											return
										} else {
											tmp9745 := PrimEqual(symshen_4_7m, V2232)

											var ifres9738 Obj

											if True == tmp9745 {
												tmp9744 := PrimIsPair(V2233)

												var ifres9740 Obj

												if True == tmp9744 {
													tmp9742 := PrimHead(V2233)

													tmp9743 := Call(__e, PrimNS2Value(symatom_2), tmp9742)

													var ifres9741 Obj

													if True == tmp9743 {
														ifres9741 = True

													} else {
														ifres9741 = False

													}

													ifres9740 = ifres9741

												} else {
													ifres9740 = False

												}

												var ifres9739 Obj

												if True == ifres9740 {
													ifres9739 = True

												} else {
													ifres9739 = False

												}

												ifres9738 = ifres9739

											} else {
												ifres9738 = False

											}

											if True == ifres9738 {
												__e.TailApply(PrimNS2Value(symshen_4atom_1case_1plus), V2233, V2234, V2235, V2236)
												return
											} else {
												tmp9737 := PrimEqual(symshen_4_7m, V2232)

												var ifres9707 Obj

												if True == tmp9737 {
													tmp9736 := PrimIsPair(V2233)

													var ifres9709 Obj

													if True == tmp9736 {
														tmp9734 := PrimHead(V2233)

														tmp9735 := PrimIsPair(tmp9734)

														var ifres9711 Obj

														if True == tmp9735 {
															tmp9731 := PrimHead(V2233)

															tmp9732 := PrimHead(tmp9731)

															tmp9733 := PrimEqual(symcons, tmp9732)

															var ifres9713 Obj

															if True == tmp9733 {
																tmp9728 := PrimHead(V2233)

																tmp9729 := PrimTail(tmp9728)

																tmp9730 := PrimIsPair(tmp9729)

																var ifres9715 Obj

																if True == tmp9730 {
																	tmp9724 := PrimHead(V2233)

																	tmp9725 := PrimTail(tmp9724)

																	tmp9726 := PrimTail(tmp9725)

																	tmp9727 := PrimIsPair(tmp9726)

																	var ifres9717 Obj

																	if True == tmp9727 {
																		tmp9719 := PrimHead(V2233)

																		tmp9720 := PrimTail(tmp9719)

																		tmp9721 := PrimTail(tmp9720)

																		tmp9722 := PrimTail(tmp9721)

																		tmp9723 := PrimEqual(Nil, tmp9722)

																		var ifres9718 Obj

																		if True == tmp9723 {
																			ifres9718 = True

																		} else {
																			ifres9718 = False

																		}

																		ifres9717 = ifres9718

																	} else {
																		ifres9717 = False

																	}

																	var ifres9716 Obj

																	if True == ifres9717 {
																		ifres9716 = True

																	} else {
																		ifres9716 = False

																	}

																	ifres9715 = ifres9716

																} else {
																	ifres9715 = False

																}

																var ifres9714 Obj

																if True == ifres9715 {
																	ifres9714 = True

																} else {
																	ifres9714 = False

																}

																ifres9713 = ifres9714

															} else {
																ifres9713 = False

															}

															var ifres9712 Obj

															if True == ifres9713 {
																ifres9712 = True

															} else {
																ifres9712 = False

															}

															ifres9711 = ifres9712

														} else {
															ifres9711 = False

														}

														var ifres9710 Obj

														if True == ifres9711 {
															ifres9710 = True

														} else {
															ifres9710 = False

														}

														ifres9709 = ifres9710

													} else {
														ifres9709 = False

													}

													var ifres9708 Obj

													if True == ifres9709 {
														ifres9708 = True

													} else {
														ifres9708 = False

													}

													ifres9707 = ifres9708

												} else {
													ifres9707 = False

												}

												if True == ifres9707 {
													__e.TailApply(PrimNS2Value(symshen_4cons_1case_1plus), V2233, V2234, V2235, V2236)
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

	tmp9872 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1head, tmp9695)

	_ = tmp9872

	tmp9873 := MakeNative(func(__e *ControlFlow) {
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
		tmp9894 := PrimIsPair(V2248)

		var ifres9891 Obj

		if True == tmp9894 {
			tmp9893 := PrimIsPair(V2249)

			var ifres9892 Obj

			if True == tmp9893 {
				ifres9892 = True

			} else {
				ifres9892 = False

			}

			ifres9891 = ifres9892

		} else {
			ifres9891 = False

		}

		if True == ifres9891 {
			tmp9889 := PrimHead(V2249)

			tmp9890 := PrimIsVariable(tmp9889)

			if True == tmp9890 {
				tmp9884 := PrimTail(V2248)

				tmp9885 := PrimTail(V2249)

				tmp9886 := PrimHead(V2249)

				tmp9887 := PrimHead(V2248)

				tmp9888 := Call(__e, PrimNS2Value(symsubst), tmp9886, tmp9887, V2251)

				__e.TailApply(PrimNS2Value(symshen_4compile_1head), V2247, tmp9884, tmp9885, V2250, tmp9888)
				return

			} else {
				tmp9876 := PrimHead(V2248)

				tmp9877 := PrimHead(V2249)

				tmp9878 := PrimTail(V2248)

				tmp9879 := PrimTail(V2249)

				tmp9880 := Call(__e, PrimNS2Value(symshen_4compile_1head), V2247, tmp9878, tmp9879, V2250, V2251)

				tmp9881 := PrimCons(tmp9880, Nil)

				tmp9882 := PrimCons(tmp9877, tmp9881)

				tmp9883 := PrimCons(tmp9876, tmp9882)

				__e.Return(PrimCons(symlet, tmp9883))
				return

			}

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
			return
		}

	}, 5)

	tmp9895 := Call(__e, PrimNS2Value(symdef), symshen_4variable_1case, tmp9873)

	_ = tmp9895

	tmp9896 := MakeNative(func(__e *ControlFlow) {
		V2260 := __e.Get(1)
		_ = V2260
		V2261 := __e.Get(2)
		_ = V2261
		V2262 := __e.Get(3)
		_ = V2262
		V2263 := __e.Get(4)
		_ = V2263
		tmp9922 := PrimIsPair(V2260)

		var ifres9919 Obj

		if True == tmp9922 {
			tmp9921 := PrimIsPair(V2261)

			var ifres9920 Obj

			if True == tmp9921 {
				ifres9920 = True

			} else {
				ifres9920 = False

			}

			ifres9919 = ifres9920

		} else {
			ifres9919 = False

		}

		if True == ifres9919 {
			tmp9898 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp9899 := PrimHead(V2261)

				tmp9900 := PrimCons(V2262, Nil)

				tmp9901 := PrimCons(tmp9899, tmp9900)

				tmp9902 := PrimCons(symshen_4lazyderef, tmp9901)

				tmp9903 := PrimHead(V2260)

				tmp9904 := PrimCons(tmp9903, Nil)

				tmp9905 := PrimCons(Tm, tmp9904)

				tmp9906 := PrimCons(sym_a, tmp9905)

				tmp9907 := PrimTail(V2260)

				tmp9908 := PrimTail(V2261)

				tmp9909 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_1m, tmp9907, tmp9908, V2262, V2263)

				tmp9910 := PrimCons(False, Nil)

				tmp9911 := PrimCons(tmp9909, tmp9910)

				tmp9912 := PrimCons(tmp9906, tmp9911)

				tmp9913 := PrimCons(symif, tmp9912)

				tmp9914 := PrimCons(tmp9913, Nil)

				tmp9915 := PrimCons(tmp9902, tmp9914)

				tmp9916 := PrimCons(Tm, tmp9915)

				__e.Return(PrimCons(symlet, tmp9916))
				return

			}, 1)

			tmp9917 := Call(__e, PrimNS2Value(symprotect), symTm)

			tmp9918 := Call(__e, PrimNS2Value(symgensym), tmp9917)

			__e.TailApply(tmp9898, tmp9918)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
			return
		}

	}, 4)

	tmp9923 := Call(__e, PrimNS2Value(symdef), symshen_4atom_1case_1minus, tmp9896)

	_ = tmp9923

	tmp9924 := MakeNative(func(__e *ControlFlow) {
		V2272 := __e.Get(1)
		_ = V2272
		V2273 := __e.Get(2)
		_ = V2273
		V2274 := __e.Get(3)
		_ = V2274
		V2275 := __e.Get(4)
		_ = V2275
		tmp9990 := PrimIsPair(V2272)

		var ifres9960 Obj

		if True == tmp9990 {
			tmp9988 := PrimHead(V2272)

			tmp9989 := PrimIsPair(tmp9988)

			var ifres9962 Obj

			if True == tmp9989 {
				tmp9985 := PrimHead(V2272)

				tmp9986 := PrimHead(tmp9985)

				tmp9987 := PrimEqual(symcons, tmp9986)

				var ifres9964 Obj

				if True == tmp9987 {
					tmp9982 := PrimHead(V2272)

					tmp9983 := PrimTail(tmp9982)

					tmp9984 := PrimIsPair(tmp9983)

					var ifres9966 Obj

					if True == tmp9984 {
						tmp9978 := PrimHead(V2272)

						tmp9979 := PrimTail(tmp9978)

						tmp9980 := PrimTail(tmp9979)

						tmp9981 := PrimIsPair(tmp9980)

						var ifres9968 Obj

						if True == tmp9981 {
							tmp9973 := PrimHead(V2272)

							tmp9974 := PrimTail(tmp9973)

							tmp9975 := PrimTail(tmp9974)

							tmp9976 := PrimTail(tmp9975)

							tmp9977 := PrimEqual(Nil, tmp9976)

							var ifres9970 Obj

							if True == tmp9977 {
								tmp9972 := PrimIsPair(V2273)

								var ifres9971 Obj

								if True == tmp9972 {
									ifres9971 = True

								} else {
									ifres9971 = False

								}

								ifres9970 = ifres9971

							} else {
								ifres9970 = False

							}

							var ifres9969 Obj

							if True == ifres9970 {
								ifres9969 = True

							} else {
								ifres9969 = False

							}

							ifres9968 = ifres9969

						} else {
							ifres9968 = False

						}

						var ifres9967 Obj

						if True == ifres9968 {
							ifres9967 = True

						} else {
							ifres9967 = False

						}

						ifres9966 = ifres9967

					} else {
						ifres9966 = False

					}

					var ifres9965 Obj

					if True == ifres9966 {
						ifres9965 = True

					} else {
						ifres9965 = False

					}

					ifres9964 = ifres9965

				} else {
					ifres9964 = False

				}

				var ifres9963 Obj

				if True == ifres9964 {
					ifres9963 = True

				} else {
					ifres9963 = False

				}

				ifres9962 = ifres9963

			} else {
				ifres9962 = False

			}

			var ifres9961 Obj

			if True == ifres9962 {
				ifres9961 = True

			} else {
				ifres9961 = False

			}

			ifres9960 = ifres9961

		} else {
			ifres9960 = False

		}

		if True == ifres9960 {
			tmp9926 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp9927 := PrimHead(V2273)

				tmp9928 := PrimCons(V2274, Nil)

				tmp9929 := PrimCons(tmp9927, tmp9928)

				tmp9930 := PrimCons(symshen_4lazyderef, tmp9929)

				tmp9931 := PrimCons(Tm, Nil)

				tmp9932 := PrimCons(symcons_2, tmp9931)

				tmp9933 := PrimHead(V2272)

				tmp9934 := PrimTail(tmp9933)

				tmp9935 := PrimHead(tmp9934)

				tmp9936 := PrimHead(V2272)

				tmp9937 := PrimTail(tmp9936)

				tmp9938 := PrimTail(tmp9937)

				tmp9939 := PrimHead(tmp9938)

				tmp9940 := PrimTail(V2272)

				tmp9941 := PrimCons(tmp9939, tmp9940)

				tmp9942 := PrimCons(tmp9935, tmp9941)

				tmp9943 := PrimCons(Tm, Nil)

				tmp9944 := PrimCons(symhd, tmp9943)

				tmp9945 := PrimCons(Tm, Nil)

				tmp9946 := PrimCons(symtl, tmp9945)

				tmp9947 := PrimTail(V2273)

				tmp9948 := PrimCons(tmp9946, tmp9947)

				tmp9949 := PrimCons(tmp9944, tmp9948)

				tmp9950 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_1m, tmp9942, tmp9949, V2274, V2275)

				tmp9951 := PrimCons(False, Nil)

				tmp9952 := PrimCons(tmp9950, tmp9951)

				tmp9953 := PrimCons(tmp9932, tmp9952)

				tmp9954 := PrimCons(symif, tmp9953)

				tmp9955 := PrimCons(tmp9954, Nil)

				tmp9956 := PrimCons(tmp9930, tmp9955)

				tmp9957 := PrimCons(Tm, tmp9956)

				__e.Return(PrimCons(symlet, tmp9957))
				return

			}, 1)

			tmp9958 := Call(__e, PrimNS2Value(symprotect), symTm)

			tmp9959 := Call(__e, PrimNS2Value(symgensym), tmp9958)

			__e.TailApply(tmp9926, tmp9959)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
			return
		}

	}, 4)

	tmp9991 := Call(__e, PrimNS2Value(symdef), symshen_4cons_1case_1minus, tmp9924)

	_ = tmp9991

	tmp9992 := MakeNative(func(__e *ControlFlow) {
		V2284 := __e.Get(1)
		_ = V2284
		V2285 := __e.Get(2)
		_ = V2285
		V2286 := __e.Get(3)
		_ = V2286
		V2287 := __e.Get(4)
		_ = V2287
		tmp10040 := PrimIsPair(V2284)

		var ifres10037 Obj

		if True == tmp10040 {
			tmp10039 := PrimIsPair(V2285)

			var ifres10038 Obj

			if True == tmp10039 {
				ifres10038 = True

			} else {
				ifres10038 = False

			}

			ifres10037 = ifres10038

		} else {
			ifres10037 = False

		}

		if True == ifres10037 {
			tmp9994 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp9995 := MakeNative(func(__e *ControlFlow) {
					GoTo := __e.Get(1)
					_ = GoTo
					tmp9996 := PrimHead(V2285)

					tmp9997 := PrimCons(V2286, Nil)

					tmp9998 := PrimCons(tmp9996, tmp9997)

					tmp9999 := PrimCons(symshen_4lazyderef, tmp9998)

					tmp10000 := PrimTail(V2284)

					tmp10001 := PrimTail(V2285)

					tmp10002 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_7m, tmp10000, tmp10001, V2286, V2287)

					tmp10003 := PrimCons(tmp10002, Nil)

					tmp10004 := PrimCons(symfreeze, tmp10003)

					tmp10005 := PrimHead(V2284)

					tmp10006 := PrimCons(tmp10005, Nil)

					tmp10007 := PrimCons(Tm, tmp10006)

					tmp10008 := PrimCons(sym_a, tmp10007)

					tmp10009 := PrimCons(GoTo, Nil)

					tmp10010 := PrimCons(symthaw, tmp10009)

					tmp10011 := PrimCons(Tm, Nil)

					tmp10012 := PrimCons(symshen_4pvar_2, tmp10011)

					tmp10013 := PrimHead(V2284)

					tmp10014 := Call(__e, PrimNS2Value(symshen_4demode), tmp10013)

					tmp10015 := PrimCons(GoTo, Nil)

					tmp10016 := PrimCons(V2286, tmp10015)

					tmp10017 := PrimCons(tmp10014, tmp10016)

					tmp10018 := PrimCons(Tm, tmp10017)

					tmp10019 := PrimCons(symshen_4bind_b, tmp10018)

					tmp10020 := PrimCons(False, Nil)

					tmp10021 := PrimCons(tmp10019, tmp10020)

					tmp10022 := PrimCons(tmp10012, tmp10021)

					tmp10023 := PrimCons(symif, tmp10022)

					tmp10024 := PrimCons(tmp10023, Nil)

					tmp10025 := PrimCons(tmp10010, tmp10024)

					tmp10026 := PrimCons(tmp10008, tmp10025)

					tmp10027 := PrimCons(symif, tmp10026)

					tmp10028 := PrimCons(tmp10027, Nil)

					tmp10029 := PrimCons(tmp10004, tmp10028)

					tmp10030 := PrimCons(GoTo, tmp10029)

					tmp10031 := PrimCons(tmp9999, tmp10030)

					tmp10032 := PrimCons(Tm, tmp10031)

					__e.Return(PrimCons(symlet, tmp10032))
					return

				}, 1)

				tmp10033 := Call(__e, PrimNS2Value(symprotect), symGoTo)

				tmp10034 := Call(__e, PrimNS2Value(symgensym), tmp10033)

				__e.TailApply(tmp9995, tmp10034)
				return

			}, 1)

			tmp10035 := Call(__e, PrimNS2Value(symprotect), symTm)

			tmp10036 := Call(__e, PrimNS2Value(symgensym), tmp10035)

			__e.TailApply(tmp9994, tmp10036)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
			return
		}

	}, 4)

	tmp10041 := Call(__e, PrimNS2Value(symdef), symshen_4atom_1case_1plus, tmp9992)

	_ = tmp10041

	tmp10042 := MakeNative(func(__e *ControlFlow) {
		V2296 := __e.Get(1)
		_ = V2296
		V2297 := __e.Get(2)
		_ = V2297
		V2298 := __e.Get(3)
		_ = V2298
		V2299 := __e.Get(4)
		_ = V2299
		tmp10140 := PrimIsPair(V2296)

		var ifres10110 Obj

		if True == tmp10140 {
			tmp10138 := PrimHead(V2296)

			tmp10139 := PrimIsPair(tmp10138)

			var ifres10112 Obj

			if True == tmp10139 {
				tmp10135 := PrimHead(V2296)

				tmp10136 := PrimHead(tmp10135)

				tmp10137 := PrimEqual(symcons, tmp10136)

				var ifres10114 Obj

				if True == tmp10137 {
					tmp10132 := PrimHead(V2296)

					tmp10133 := PrimTail(tmp10132)

					tmp10134 := PrimIsPair(tmp10133)

					var ifres10116 Obj

					if True == tmp10134 {
						tmp10128 := PrimHead(V2296)

						tmp10129 := PrimTail(tmp10128)

						tmp10130 := PrimTail(tmp10129)

						tmp10131 := PrimIsPair(tmp10130)

						var ifres10118 Obj

						if True == tmp10131 {
							tmp10123 := PrimHead(V2296)

							tmp10124 := PrimTail(tmp10123)

							tmp10125 := PrimTail(tmp10124)

							tmp10126 := PrimTail(tmp10125)

							tmp10127 := PrimEqual(Nil, tmp10126)

							var ifres10120 Obj

							if True == tmp10127 {
								tmp10122 := PrimIsPair(V2297)

								var ifres10121 Obj

								if True == tmp10122 {
									ifres10121 = True

								} else {
									ifres10121 = False

								}

								ifres10120 = ifres10121

							} else {
								ifres10120 = False

							}

							var ifres10119 Obj

							if True == ifres10120 {
								ifres10119 = True

							} else {
								ifres10119 = False

							}

							ifres10118 = ifres10119

						} else {
							ifres10118 = False

						}

						var ifres10117 Obj

						if True == ifres10118 {
							ifres10117 = True

						} else {
							ifres10117 = False

						}

						ifres10116 = ifres10117

					} else {
						ifres10116 = False

					}

					var ifres10115 Obj

					if True == ifres10116 {
						ifres10115 = True

					} else {
						ifres10115 = False

					}

					ifres10114 = ifres10115

				} else {
					ifres10114 = False

				}

				var ifres10113 Obj

				if True == ifres10114 {
					ifres10113 = True

				} else {
					ifres10113 = False

				}

				ifres10112 = ifres10113

			} else {
				ifres10112 = False

			}

			var ifres10111 Obj

			if True == ifres10112 {
				ifres10111 = True

			} else {
				ifres10111 = False

			}

			ifres10110 = ifres10111

		} else {
			ifres10110 = False

		}

		if True == ifres10110 {
			tmp10044 := MakeNative(func(__e *ControlFlow) {
				Tm := __e.Get(1)
				_ = Tm
				tmp10045 := MakeNative(func(__e *ControlFlow) {
					GoTo := __e.Get(1)
					_ = GoTo
					tmp10046 := MakeNative(func(__e *ControlFlow) {
						Vars := __e.Get(1)
						_ = Vars
						tmp10047 := MakeNative(func(__e *ControlFlow) {
							Tame := __e.Get(1)
							_ = Tame
							tmp10048 := MakeNative(func(__e *ControlFlow) {
								TVars := __e.Get(1)
								_ = TVars
								tmp10049 := PrimHead(V2297)

								tmp10050 := PrimCons(V2298, Nil)

								tmp10051 := PrimCons(tmp10049, tmp10050)

								tmp10052 := PrimCons(symshen_4lazyderef, tmp10051)

								tmp10053 := PrimTail(V2296)

								tmp10054 := PrimTail(V2297)

								tmp10055 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_7m, tmp10053, tmp10054, V2298, V2299)

								tmp10056 := Call(__e, PrimNS2Value(symshen_4goto), Vars, tmp10055)

								tmp10057 := PrimCons(Tm, Nil)

								tmp10058 := PrimCons(symcons_2, tmp10057)

								tmp10059 := PrimHead(V2296)

								tmp10060 := PrimTail(tmp10059)

								tmp10061 := PrimCons(Tm, Nil)

								tmp10062 := PrimCons(symhd, tmp10061)

								tmp10063 := PrimCons(Tm, Nil)

								tmp10064 := PrimCons(symtl, tmp10063)

								tmp10065 := PrimCons(tmp10064, Nil)

								tmp10066 := PrimCons(tmp10062, tmp10065)

								tmp10067 := Call(__e, PrimNS2Value(symshen_4invoke), GoTo, Vars)

								tmp10068 := Call(__e, PrimNS2Value(symshen_4compile_1head), symshen_4_7m, tmp10060, tmp10066, V2298, tmp10067)

								tmp10069 := PrimCons(Tm, Nil)

								tmp10070 := PrimCons(symshen_4pvar_2, tmp10069)

								tmp10071 := Call(__e, PrimNS2Value(symshen_4demode), Tame)

								tmp10072 := Call(__e, PrimNS2Value(symshen_4invoke), GoTo, Vars)

								tmp10073 := PrimCons(tmp10072, Nil)

								tmp10074 := PrimCons(symfreeze, tmp10073)

								tmp10075 := PrimCons(tmp10074, Nil)

								tmp10076 := PrimCons(V2298, tmp10075)

								tmp10077 := PrimCons(tmp10071, tmp10076)

								tmp10078 := PrimCons(Tm, tmp10077)

								tmp10079 := PrimCons(symshen_4bind_b, tmp10078)

								tmp10080 := Call(__e, PrimNS2Value(symshen_4stpart), TVars, tmp10079, V2298)

								tmp10081 := PrimCons(False, Nil)

								tmp10082 := PrimCons(tmp10080, tmp10081)

								tmp10083 := PrimCons(tmp10070, tmp10082)

								tmp10084 := PrimCons(symif, tmp10083)

								tmp10085 := PrimCons(tmp10084, Nil)

								tmp10086 := PrimCons(tmp10068, tmp10085)

								tmp10087 := PrimCons(tmp10058, tmp10086)

								tmp10088 := PrimCons(symif, tmp10087)

								tmp10089 := PrimCons(tmp10088, Nil)

								tmp10090 := PrimCons(tmp10056, tmp10089)

								tmp10091 := PrimCons(GoTo, tmp10090)

								tmp10092 := PrimCons(tmp10052, tmp10091)

								tmp10093 := PrimCons(Tm, tmp10092)

								__e.Return(PrimCons(symlet, tmp10093))
								return

							}, 1)

							tmp10094 := Call(__e, PrimNS2Value(symshen_4extract_1vars), Tame)

							__e.TailApply(tmp10048, tmp10094)
							return

						}, 1)

						tmp10095 := PrimHead(V2296)

						tmp10096 := Call(__e, PrimNS2Value(symshen_4tame), tmp10095)

						__e.TailApply(tmp10047, tmp10096)
						return

					}, 1)

					tmp10097 := PrimHead(V2296)

					tmp10098 := PrimTail(tmp10097)

					tmp10099 := PrimHead(tmp10098)

					tmp10100 := PrimHead(V2296)

					tmp10101 := PrimTail(tmp10100)

					tmp10102 := PrimTail(tmp10101)

					tmp10103 := PrimHead(tmp10102)

					tmp10104 := PrimCons(tmp10099, tmp10103)

					tmp10105 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp10104)

					__e.TailApply(tmp10046, tmp10105)
					return

				}, 1)

				tmp10106 := Call(__e, PrimNS2Value(symprotect), symGoTo)

				tmp10107 := Call(__e, PrimNS2Value(symgensym), tmp10106)

				__e.TailApply(tmp10045, tmp10107)
				return

			}, 1)

			tmp10108 := Call(__e, PrimNS2Value(symprotect), symTm)

			tmp10109 := Call(__e, PrimNS2Value(symgensym), tmp10108)

			__e.TailApply(tmp10044, tmp10109)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
			return
		}

	}, 4)

	tmp10141 := Call(__e, PrimNS2Value(symdef), symshen_4cons_1case_1plus, tmp10042)

	_ = tmp10141

	tmp10142 := MakeNative(func(__e *ControlFlow) {
		V2300 := __e.Get(1)
		_ = V2300
		tmp10179 := PrimIsPair(V2300)

		var ifres10166 Obj

		if True == tmp10179 {
			tmp10177 := PrimHead(V2300)

			tmp10178 := PrimEqual(symshen_4_7m, tmp10177)

			var ifres10168 Obj

			if True == tmp10178 {
				tmp10175 := PrimTail(V2300)

				tmp10176 := PrimIsPair(tmp10175)

				var ifres10170 Obj

				if True == tmp10176 {
					tmp10172 := PrimTail(V2300)

					tmp10173 := PrimTail(tmp10172)

					tmp10174 := PrimEqual(Nil, tmp10173)

					var ifres10171 Obj

					if True == tmp10174 {
						ifres10171 = True

					} else {
						ifres10171 = False

					}

					ifres10170 = ifres10171

				} else {
					ifres10170 = False

				}

				var ifres10169 Obj

				if True == ifres10170 {
					ifres10169 = True

				} else {
					ifres10169 = False

				}

				ifres10168 = ifres10169

			} else {
				ifres10168 = False

			}

			var ifres10167 Obj

			if True == ifres10168 {
				ifres10167 = True

			} else {
				ifres10167 = False

			}

			ifres10166 = ifres10167

		} else {
			ifres10166 = False

		}

		if True == ifres10166 {
			tmp10164 := PrimTail(V2300)

			tmp10165 := PrimHead(tmp10164)

			__e.TailApply(PrimNS2Value(symshen_4demode), tmp10165)
			return

		} else {
			tmp10163 := PrimIsPair(V2300)

			var ifres10150 Obj

			if True == tmp10163 {
				tmp10161 := PrimHead(V2300)

				tmp10162 := PrimEqual(symshen_4_1m, tmp10161)

				var ifres10152 Obj

				if True == tmp10162 {
					tmp10159 := PrimTail(V2300)

					tmp10160 := PrimIsPair(tmp10159)

					var ifres10154 Obj

					if True == tmp10160 {
						tmp10156 := PrimTail(V2300)

						tmp10157 := PrimTail(tmp10156)

						tmp10158 := PrimEqual(Nil, tmp10157)

						var ifres10155 Obj

						if True == tmp10158 {
							ifres10155 = True

						} else {
							ifres10155 = False

						}

						ifres10154 = ifres10155

					} else {
						ifres10154 = False

					}

					var ifres10153 Obj

					if True == ifres10154 {
						ifres10153 = True

					} else {
						ifres10153 = False

					}

					ifres10152 = ifres10153

				} else {
					ifres10152 = False

				}

				var ifres10151 Obj

				if True == ifres10152 {
					ifres10151 = True

				} else {
					ifres10151 = False

				}

				ifres10150 = ifres10151

			} else {
				ifres10150 = False

			}

			if True == ifres10150 {
				tmp10148 := PrimTail(V2300)

				tmp10149 := PrimHead(tmp10148)

				__e.TailApply(PrimNS2Value(symshen_4demode), tmp10149)
				return

			} else {
				tmp10147 := PrimIsPair(V2300)

				if True == tmp10147 {
					tmp10146 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(PrimNS2Value(symshen_4demode), Z)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symmap), tmp10146, V2300)
					return

				} else {
					__e.Return(V2300)
					return
				}

			}

		}

	}, 1)

	tmp10180 := Call(__e, PrimNS2Value(symdef), symshen_4demode, tmp10142)

	_ = tmp10180

	tmp10181 := MakeNative(func(__e *ControlFlow) {
		V2301 := __e.Get(1)
		_ = V2301
		tmp10187 := Call(__e, PrimNS2Value(symshen_4wildcard_2), V2301)

		if True == tmp10187 {
			tmp10186 := Call(__e, PrimNS2Value(symprotect), symY)

			__e.TailApply(PrimNS2Value(symgensym), tmp10186)
			return

		} else {
			tmp10185 := PrimIsPair(V2301)

			if True == tmp10185 {
				tmp10184 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symshen_4tame), Z)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symmap), tmp10184, V2301)
				return

			} else {
				__e.Return(V2301)
				return
			}

		}

	}, 1)

	tmp10188 := Call(__e, PrimNS2Value(symdef), symshen_4tame, tmp10181)

	_ = tmp10188

	tmp10189 := MakeNative(func(__e *ControlFlow) {
		V2302 := __e.Get(1)
		_ = V2302
		V2303 := __e.Get(2)
		_ = V2303
		tmp10192 := PrimEqual(Nil, V2302)

		if True == tmp10192 {
			tmp10191 := PrimCons(V2303, Nil)

			__e.Return(PrimCons(symfreeze, tmp10191))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4goto_1h), V2302, V2303)
			return
		}

	}, 2)

	tmp10193 := Call(__e, PrimNS2Value(symdef), symshen_4goto, tmp10189)

	_ = tmp10193

	tmp10194 := MakeNative(func(__e *ControlFlow) {
		V2304 := __e.Get(1)
		_ = V2304
		V2305 := __e.Get(2)
		_ = V2305
		tmp10203 := PrimEqual(Nil, V2304)

		if True == tmp10203 {
			__e.Return(V2305)
			return
		} else {
			tmp10202 := PrimIsPair(V2304)

			if True == tmp10202 {
				tmp10197 := PrimHead(V2304)

				tmp10198 := PrimTail(V2304)

				tmp10199 := Call(__e, PrimNS2Value(symshen_4goto_1h), tmp10198, V2305)

				tmp10200 := PrimCons(tmp10199, Nil)

				tmp10201 := PrimCons(tmp10197, tmp10200)

				__e.Return(PrimCons(symlambda, tmp10201))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4goto_1h)
				return
			}

		}

	}, 2)

	tmp10204 := Call(__e, PrimNS2Value(symdef), symshen_4goto_1h, tmp10194)

	_ = tmp10204

	tmp10205 := MakeNative(func(__e *ControlFlow) {
		V2306 := __e.Get(1)
		_ = V2306
		V2307 := __e.Get(2)
		_ = V2307
		tmp10208 := PrimEqual(Nil, V2307)

		if True == tmp10208 {
			tmp10207 := PrimCons(V2306, Nil)

			__e.Return(PrimCons(symthaw, tmp10207))
			return

		} else {
			__e.Return(PrimCons(V2306, V2307))
			return
		}

	}, 2)

	tmp10209 := Call(__e, PrimNS2Value(symdef), symshen_4invoke, tmp10205)

	_ = tmp10209

	tmp10210 := MakeNative(func(__e *ControlFlow) {
		V2308 := __e.Get(1)
		_ = V2308
		__e.Return(PrimEqual(V2308, sym__))
		return
	}, 1)

	tmp10211 := Call(__e, PrimNS2Value(symdef), symshen_4wildcard_2, tmp10210)

	_ = tmp10211

	tmp10212 := MakeNative(func(__e *ControlFlow) {
		V2309 := __e.Get(1)
		_ = V2309
		tmp10213 := MakeNative(func(__e *ControlFlow) {
			tmp10218 := PrimIsVector(V2309)

			if True == tmp10218 {
				tmp10216 := PrimVectorGet(V2309, MakeNumber(0))

				tmp10217 := PrimEqual(tmp10216, symshen_4pvar)

				if True == tmp10217 {
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

		tmp10219 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp10213, tmp10219)
		return

	}, 1)

	tmp10220 := Call(__e, PrimNS2Value(symdef), symshen_4pvar_2, tmp10212)

	_ = tmp10220

	tmp10221 := MakeNative(func(__e *ControlFlow) {
		V2310 := __e.Get(1)
		_ = V2310
		V2311 := __e.Get(2)
		_ = V2311
		tmp10228 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2310)

		if True == tmp10228 {
			tmp10223 := MakeNative(func(__e *ControlFlow) {
				Value := __e.Get(1)
				_ = Value
				tmp10225 := PrimEqual(Value, symshen_4_1null_1)

				if True == tmp10225 {
					__e.Return(V2310)
					return
				} else {
					__e.TailApply(PrimNS2Value(symshen_4lazyderef), Value, V2311)
					return
				}

			}, 1)

			tmp10226 := PrimVectorGet(V2310, MakeNumber(1))

			tmp10227 := PrimVectorGet(V2311, tmp10226)

			__e.TailApply(tmp10223, tmp10227)
			return

		} else {
			__e.Return(V2310)
			return
		}

	}, 2)

	tmp10229 := Call(__e, PrimNS2Value(symdef), symshen_4lazyderef, tmp10221)

	_ = tmp10229

	tmp10230 := MakeNative(func(__e *ControlFlow) {
		V2312 := __e.Get(1)
		_ = V2312
		V2313 := __e.Get(2)
		_ = V2313
		tmp10243 := PrimIsPair(V2312)

		if True == tmp10243 {
			tmp10239 := PrimHead(V2312)

			tmp10240 := Call(__e, PrimNS2Value(symshen_4deref), tmp10239, V2313)

			tmp10241 := PrimTail(V2312)

			tmp10242 := Call(__e, PrimNS2Value(symshen_4deref), tmp10241, V2313)

			__e.Return(PrimCons(tmp10240, tmp10242))
			return

		} else {
			tmp10238 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2312)

			if True == tmp10238 {
				tmp10233 := MakeNative(func(__e *ControlFlow) {
					Value := __e.Get(1)
					_ = Value
					tmp10235 := PrimEqual(Value, symshen_4_1null_1)

					if True == tmp10235 {
						__e.Return(V2312)
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4deref), Value, V2313)
						return
					}

				}, 1)

				tmp10236 := PrimVectorGet(V2312, MakeNumber(1))

				tmp10237 := PrimVectorGet(V2313, tmp10236)

				__e.TailApply(tmp10233, tmp10237)
				return

			} else {
				__e.Return(V2312)
				return
			}

		}

	}, 2)

	tmp10244 := Call(__e, PrimNS2Value(symdef), symshen_4deref, tmp10230)

	_ = tmp10244

	tmp10245 := MakeNative(func(__e *ControlFlow) {
		V2314 := __e.Get(1)
		_ = V2314
		V2315 := __e.Get(2)
		_ = V2315
		V2316 := __e.Get(3)
		_ = V2316
		V2317 := __e.Get(4)
		_ = V2317
		tmp10246 := MakeNative(func(__e *ControlFlow) {
			Bind := __e.Get(1)
			_ = Bind
			tmp10247 := MakeNative(func(__e *ControlFlow) {
				Compute := __e.Get(1)
				_ = Compute
				tmp10249 := PrimEqual(Compute, False)

				if True == tmp10249 {
					__e.TailApply(PrimNS2Value(symshen_4unwind), V2314, V2316, Compute)
					return
				} else {
					__e.Return(Compute)
					return
				}

			}, 1)

			tmp10250 := Call(__e, PrimNS2Value(symthaw), V2317)

			__e.TailApply(tmp10247, tmp10250)
			return

		}, 1)

		tmp10251 := Call(__e, PrimNS2Value(symshen_4bindv), V2314, V2315, V2316)

		__e.TailApply(tmp10246, tmp10251)
		return

	}, 4)

	tmp10252 := Call(__e, PrimNS2Value(symdef), symshen_4bind_b, tmp10245)

	_ = tmp10252

	tmp10253 := MakeNative(func(__e *ControlFlow) {
		V2318 := __e.Get(1)
		_ = V2318
		V2319 := __e.Get(2)
		_ = V2319
		V2320 := __e.Get(3)
		_ = V2320
		tmp10254 := PrimVectorGet(V2318, MakeNumber(1))

		__e.Return(PrimVectorSet(V2320, tmp10254, V2319))
		return

	}, 3)

	tmp10255 := Call(__e, PrimNS2Value(symdef), symshen_4bindv, tmp10253)

	_ = tmp10255

	tmp10256 := MakeNative(func(__e *ControlFlow) {
		V2321 := __e.Get(1)
		_ = V2321
		V2322 := __e.Get(2)
		_ = V2322
		V2323 := __e.Get(3)
		_ = V2323
		tmp10257 := PrimVectorGet(V2321, MakeNumber(1))

		tmp10258 := PrimVectorSet(V2322, tmp10257, symshen_4_1null_1)

		_ = tmp10258

		__e.Return(V2323)
		return

	}, 3)

	tmp10259 := Call(__e, PrimNS2Value(symdef), symshen_4unwind, tmp10256)

	_ = tmp10259

	tmp10260 := MakeNative(func(__e *ControlFlow) {
		V2332 := __e.Get(1)
		_ = V2332
		V2333 := __e.Get(2)
		_ = V2333
		V2334 := __e.Get(3)
		_ = V2334
		tmp10275 := PrimEqual(Nil, V2332)

		if True == tmp10275 {
			__e.Return(V2333)
			return
		} else {
			tmp10274 := PrimIsPair(V2332)

			if True == tmp10274 {
				tmp10263 := PrimHead(V2332)

				tmp10264 := PrimCons(V2334, Nil)

				tmp10265 := PrimCons(symshen_4newpv, tmp10264)

				tmp10266 := PrimTail(V2332)

				tmp10267 := Call(__e, PrimNS2Value(symshen_4stpart), tmp10266, V2333, V2334)

				tmp10268 := PrimCons(tmp10267, Nil)

				tmp10269 := PrimCons(V2334, tmp10268)

				tmp10270 := PrimCons(symshen_4gc, tmp10269)

				tmp10271 := PrimCons(tmp10270, Nil)

				tmp10272 := PrimCons(tmp10265, tmp10271)

				tmp10273 := PrimCons(tmp10263, tmp10272)

				__e.Return(PrimCons(symlet, tmp10273))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
				return
			}

		}

	}, 3)

	tmp10276 := Call(__e, PrimNS2Value(symdef), symshen_4stpart, tmp10260)

	_ = tmp10276

	tmp10277 := MakeNative(func(__e *ControlFlow) {
		V2335 := __e.Get(1)
		_ = V2335
		V2336 := __e.Get(2)
		_ = V2336
		tmp10282 := PrimEqual(V2336, False)

		if True == tmp10282 {
			tmp10279 := MakeNative(func(__e *ControlFlow) {
				N := __e.Get(1)
				_ = N
				tmp10280 := Call(__e, PrimNS2Value(symshen_4decrement_1ticket), N, V2335)

				_ = tmp10280

				__e.Return(V2336)
				return

			}, 1)

			tmp10281 := Call(__e, PrimNS2Value(symshen_4ticket_1number), V2335)

			__e.TailApply(tmp10279, tmp10281)
			return

		} else {
			__e.Return(V2336)
			return
		}

	}, 2)

	tmp10283 := Call(__e, PrimNS2Value(symdef), symshen_4gc, tmp10277)

	_ = tmp10283

	tmp10284 := MakeNative(func(__e *ControlFlow) {
		V2337 := __e.Get(1)
		_ = V2337
		V2338 := __e.Get(2)
		_ = V2338
		tmp10285 := PrimNumberSubtract(V2337, MakeNumber(1))

		__e.Return(PrimVectorSet(V2338, MakeNumber(1), tmp10285))
		return

	}, 2)

	tmp10286 := Call(__e, PrimNS2Value(symdef), symshen_4decrement_1ticket, tmp10284)

	_ = tmp10286

	tmp10287 := MakeNative(func(__e *ControlFlow) {
		V2339 := __e.Get(1)
		_ = V2339
		tmp10288 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp10289 := MakeNative(func(__e *ControlFlow) {
				NewBindings := __e.Get(1)
				_ = NewBindings
				tmp10290 := MakeNative(func(__e *ControlFlow) {
					NextTicket := __e.Get(1)
					_ = NextTicket
					__e.Return(NewBindings)
					return
				}, 1)

				tmp10291 := Call(__e, PrimNS2Value(symshen_4nextticket), V2339, N)

				__e.TailApply(tmp10290, tmp10291)
				return

			}, 1)

			tmp10292 := Call(__e, PrimNS2Value(symshen_4make_1prolog_1variable), N)

			__e.TailApply(tmp10289, tmp10292)
			return

		}, 1)

		tmp10293 := Call(__e, PrimNS2Value(symshen_4ticket_1number), V2339)

		__e.TailApply(tmp10288, tmp10293)
		return

	}, 1)

	tmp10294 := Call(__e, PrimNS2Value(symdef), symshen_4newpv, tmp10287)

	_ = tmp10294

	tmp10295 := MakeNative(func(__e *ControlFlow) {
		V2340 := __e.Get(1)
		_ = V2340
		__e.Return(PrimVectorGet(V2340, MakeNumber(1)))
		return
	}, 1)

	tmp10296 := Call(__e, PrimNS2Value(symdef), symshen_4ticket_1number, tmp10295)

	_ = tmp10296

	tmp10297 := MakeNative(func(__e *ControlFlow) {
		V2341 := __e.Get(1)
		_ = V2341
		V2342 := __e.Get(2)
		_ = V2342
		tmp10298 := MakeNative(func(__e *ControlFlow) {
			NewVector := __e.Get(1)
			_ = NewVector
			tmp10299 := PrimNumberAdd(V2342, MakeNumber(1))

			__e.Return(PrimVectorSet(NewVector, MakeNumber(1), tmp10299))
			return

		}, 1)

		tmp10300 := PrimVectorSet(V2341, V2342, symshen_4_1null_1)

		__e.TailApply(tmp10298, tmp10300)
		return

	}, 2)

	tmp10301 := Call(__e, PrimNS2Value(symdef), symshen_4nextticket, tmp10297)

	_ = tmp10301

	tmp10302 := MakeNative(func(__e *ControlFlow) {
		V2343 := __e.Get(1)
		_ = V2343
		tmp10303 := PrimAbsvector(MakeNumber(2))

		tmp10304 := PrimVectorSet(tmp10303, MakeNumber(0), symshen_4pvar)

		__e.Return(PrimVectorSet(tmp10304, MakeNumber(1), V2343))
		return

	}, 1)

	tmp10305 := Call(__e, PrimNS2Value(symdef), symshen_4make_1prolog_1variable, tmp10302)

	_ = tmp10305

	tmp10306 := MakeNative(func(__e *ControlFlow) {
		V2344 := __e.Get(1)
		_ = V2344
		tmp10307 := PrimVectorGet(V2344, MakeNumber(1))

		tmp10308 := Call(__e, PrimNS2Value(symshen_4app), tmp10307, MakeString(""), symshen_4a)

		__e.Return(PrimStringConcat(MakeString("Var"), tmp10308))
		return

	}, 1)

	tmp10309 := Call(__e, PrimNS2Value(symdef), symshen_4pvar, tmp10306)

	_ = tmp10309

	tmp10310 := MakeNative(func(__e *ControlFlow) {
		tmp10311 := PrimNS3Value(symshen_4_dinfs_d)

		tmp10312 := PrimNumberAdd(MakeNumber(1), tmp10311)

		__e.Return(PrimNS3Set(symshen_4_dinfs_d, tmp10312))
		return

	}, 0)

	tmp10313 := Call(__e, PrimNS2Value(symdef), symshen_4incinfs, tmp10310)

	_ = tmp10313

	tmp10314 := MakeNative(func(__e *ControlFlow) {
		V2345 := __e.Get(1)
		_ = V2345
		tmp10321 := PrimIsInteger(V2345)

		var ifres10318 Obj

		if True == tmp10321 {
			tmp10320 := PrimGreatThan(V2345, MakeNumber(0))

			var ifres10319 Obj

			if True == tmp10320 {
				ifres10319 = True

			} else {
				ifres10319 = False

			}

			ifres10318 = ifres10319

		} else {
			ifres10318 = False

		}

		if True == ifres10318 {
			__e.Return(PrimNS3Set(symshen_4_dsize_1prolog_1vector_d, V2345))
			return
		} else {
			tmp10316 := Call(__e, PrimNS2Value(symshen_4app), V2345, MakeString(""), symshen_4a)

			tmp10317 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp10316)

			__e.Return(PrimSimpleError(tmp10317))
			return

		}

	}, 1)

	tmp10322 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1vector_1size, tmp10314)

	_ = tmp10322

	tmp10323 := MakeNative(func(__e *ControlFlow) {
		V2357 := __e.Get(1)
		_ = V2357
		V2358 := __e.Get(2)
		_ = V2358
		V2359 := __e.Get(3)
		_ = V2359
		V2360 := __e.Get(4)
		_ = V2360
		tmp10353 := PrimEqual(V2357, V2358)

		if True == tmp10353 {
			__e.TailApply(PrimNS2Value(symthaw), V2360)
			return
		} else {
			tmp10352 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2357)

			var ifres10347 Obj

			if True == tmp10352 {
				tmp10349 := Call(__e, PrimNS2Value(symshen_4deref), V2358, V2359)

				tmp10350 := Call(__e, PrimNS2Value(symshen_4occurs_2), V2357, tmp10349)

				tmp10351 := PrimNot(tmp10350)

				var ifres10348 Obj

				if True == tmp10351 {
					ifres10348 = True

				} else {
					ifres10348 = False

				}

				ifres10347 = ifres10348

			} else {
				ifres10347 = False

			}

			if True == ifres10347 {
				__e.TailApply(PrimNS2Value(symshen_4bind_b), V2357, V2358, V2359, V2360)
				return
			} else {
				tmp10346 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2358)

				var ifres10341 Obj

				if True == tmp10346 {
					tmp10343 := Call(__e, PrimNS2Value(symshen_4deref), V2357, V2359)

					tmp10344 := Call(__e, PrimNS2Value(symshen_4occurs_2), V2358, tmp10343)

					tmp10345 := PrimNot(tmp10344)

					var ifres10342 Obj

					if True == tmp10345 {
						ifres10342 = True

					} else {
						ifres10342 = False

					}

					ifres10341 = ifres10342

				} else {
					ifres10341 = False

				}

				if True == ifres10341 {
					__e.TailApply(PrimNS2Value(symshen_4bind_b), V2358, V2357, V2359, V2360)
					return
				} else {
					tmp10340 := PrimIsPair(V2357)

					var ifres10337 Obj

					if True == tmp10340 {
						tmp10339 := PrimIsPair(V2358)

						var ifres10338 Obj

						if True == tmp10339 {
							ifres10338 = True

						} else {
							ifres10338 = False

						}

						ifres10337 = ifres10338

					} else {
						ifres10337 = False

					}

					if True == ifres10337 {
						tmp10328 := PrimHead(V2357)

						tmp10329 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10328, V2359)

						tmp10330 := PrimHead(V2358)

						tmp10331 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10330, V2359)

						tmp10332 := MakeNative(func(__e *ControlFlow) {
							tmp10333 := PrimTail(V2357)

							tmp10334 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10333, V2359)

							tmp10335 := PrimTail(V2358)

							tmp10336 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10335, V2359)

							__e.TailApply(PrimNS2Value(symshen_4lzy_a_b), tmp10334, tmp10336, V2359, V2360)
							return

						}, 0)

						__e.TailApply(PrimNS2Value(symshen_4lzy_a_b), tmp10329, tmp10331, V2359, tmp10332)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp10354 := Call(__e, PrimNS2Value(symdef), symshen_4lzy_a_b, tmp10323)

	_ = tmp10354

	tmp10355 := MakeNative(func(__e *ControlFlow) {
		V2372 := __e.Get(1)
		_ = V2372
		V2373 := __e.Get(2)
		_ = V2373
		V2374 := __e.Get(3)
		_ = V2374
		V2375 := __e.Get(4)
		_ = V2375
		tmp10375 := PrimEqual(V2372, V2373)

		if True == tmp10375 {
			__e.TailApply(PrimNS2Value(symthaw), V2375)
			return
		} else {
			tmp10374 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2372)

			if True == tmp10374 {
				__e.TailApply(PrimNS2Value(symshen_4bind_b), V2372, V2373, V2374, V2375)
				return
			} else {
				tmp10373 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2373)

				if True == tmp10373 {
					__e.TailApply(PrimNS2Value(symshen_4bind_b), V2373, V2372, V2374, V2375)
					return
				} else {
					tmp10372 := PrimIsPair(V2372)

					var ifres10369 Obj

					if True == tmp10372 {
						tmp10371 := PrimIsPair(V2373)

						var ifres10370 Obj

						if True == tmp10371 {
							ifres10370 = True

						} else {
							ifres10370 = False

						}

						ifres10369 = ifres10370

					} else {
						ifres10369 = False

					}

					if True == ifres10369 {
						tmp10360 := PrimHead(V2372)

						tmp10361 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10360, V2374)

						tmp10362 := PrimHead(V2373)

						tmp10363 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10362, V2374)

						tmp10364 := MakeNative(func(__e *ControlFlow) {
							tmp10365 := PrimTail(V2372)

							tmp10366 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10365, V2374)

							tmp10367 := PrimTail(V2373)

							tmp10368 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp10367, V2374)

							__e.TailApply(PrimNS2Value(symshen_4lzy_a), tmp10366, tmp10368, V2374, V2375)
							return

						}, 0)

						__e.TailApply(PrimNS2Value(symshen_4lzy_a), tmp10361, tmp10363, V2374, tmp10364)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp10376 := Call(__e, PrimNS2Value(symdef), symshen_4lzy_a, tmp10355)

	_ = tmp10376

	tmp10377 := MakeNative(func(__e *ControlFlow) {
		V2381 := __e.Get(1)
		_ = V2381
		V2382 := __e.Get(2)
		_ = V2382
		tmp10387 := PrimEqual(V2381, V2382)

		if True == tmp10387 {
			__e.Return(True)
			return
		} else {
			tmp10386 := PrimIsPair(V2382)

			if True == tmp10386 {
				tmp10384 := PrimHead(V2382)

				tmp10385 := Call(__e, PrimNS2Value(symshen_4occurs_2), V2381, tmp10384)

				if True == tmp10385 {
					__e.Return(True)
					return
				} else {
					tmp10382 := PrimTail(V2382)

					tmp10383 := Call(__e, PrimNS2Value(symshen_4occurs_2), V2381, tmp10382)

					if True == tmp10383 {
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

	tmp10388 := Call(__e, PrimNS2Value(symdef), symshen_4occurs_2, tmp10377)

	_ = tmp10388

	tmp10389 := MakeNative(func(__e *ControlFlow) {
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
		tmp10390 := Call(__e, V2383, V2384)

		tmp10391 := Call(__e, tmp10390, V2385)

		tmp10392 := Call(__e, tmp10391, V2386)

		__e.TailApply(tmp10392, V2387)
		return

	}, 5)

	tmp10393 := Call(__e, PrimNS2Value(symdef), symcall, tmp10389)

	_ = tmp10393

	tmp10394 := MakeNative(func(__e *ControlFlow) {
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
		__e.TailApply(PrimNS2Value(symshen_4deref), V2394, V2395)
		return
	}, 5)

	tmp10395 := Call(__e, PrimNS2Value(symdef), symreturn, tmp10394)

	_ = tmp10395

	tmp10396 := MakeNative(func(__e *ControlFlow) {
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
			__e.TailApply(PrimNS2Value(symthaw), V2409)
			return
		} else {
			__e.Return(False)
			return
		}
	}, 5)

	tmp10398 := Call(__e, PrimNS2Value(symdef), symwhen, tmp10396)

	_ = tmp10398

	tmp10399 := MakeNative(func(__e *ControlFlow) {
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
		tmp10400 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2410, V2412)

		tmp10401 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2411, V2412)

		__e.TailApply(PrimNS2Value(symshen_4lzy_a), tmp10400, tmp10401, V2412, V2415)
		return

	}, 6)

	tmp10402 := Call(__e, PrimNS2Value(symdef), symis, tmp10399)

	_ = tmp10402

	tmp10403 := MakeNative(func(__e *ControlFlow) {
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
		tmp10404 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2416, V2418)

		tmp10405 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2417, V2418)

		__e.TailApply(PrimNS2Value(symshen_4lzy_a_b), tmp10404, tmp10405, V2418, V2421)
		return

	}, 6)

	tmp10406 := Call(__e, PrimNS2Value(symdef), symis_b, tmp10403)

	_ = tmp10406

	tmp10407 := MakeNative(func(__e *ControlFlow) {
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
		__e.TailApply(PrimNS2Value(symshen_4bind_b), V2426, V2427, V2428, V2431)
		return
	}, 6)

	tmp10408 := Call(__e, PrimNS2Value(symdef), symbind, tmp10407)

	_ = tmp10408

	tmp10409 := MakeNative(func(__e *ControlFlow) {
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
		tmp10411 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2432, V2433)

		tmp10412 := Call(__e, PrimNS2Value(symshen_4pvar_2), tmp10411)

		if True == tmp10412 {
			__e.TailApply(PrimNS2Value(symthaw), V2436)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 5)

	tmp10413 := Call(__e, PrimNS2Value(symdef), symvar_2, tmp10409)

	_ = tmp10413

	tmp10414 := MakeNative(func(__e *ControlFlow) {
		V2439 := __e.Get(1)
		_ = V2439
		__e.Return(MakeString("|prolog vector|"))
		return
	}, 1)

	tmp10415 := Call(__e, PrimNS2Value(symdef), symshen_4print_1prolog_1vector, tmp10414)

	_ = tmp10415

	tmp10416 := MakeNative(func(__e *ControlFlow) {
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
		tmp10429 := PrimEqual(Nil, V2458)

		if True == tmp10429 {
			__e.Return(False)
			return
		} else {
			tmp10428 := PrimIsPair(V2458)

			if True == tmp10428 {
				tmp10419 := MakeNative(func(__e *ControlFlow) {
					Case := __e.Get(1)
					_ = Case
					tmp10422 := PrimEqual(Case, False)

					if True == tmp10422 {
						tmp10421 := PrimTail(V2458)

						__e.TailApply(PrimNS2Value(symfork), tmp10421, V2459, V2460, V2461, V2462)
						return

					} else {
						__e.Return(Case)
						return
					}

				}, 1)

				tmp10423 := PrimHead(V2458)

				tmp10424 := Call(__e, tmp10423, V2459)

				tmp10425 := Call(__e, tmp10424, V2460)

				tmp10426 := Call(__e, tmp10425, V2461)

				tmp10427 := Call(__e, tmp10426, V2462)

				__e.TailApply(tmp10419, tmp10427)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
				return
			}

		}

	}, 5)

	tmp10430 := Call(__e, PrimNS2Value(symdef), symfork, tmp10416)

	_ = tmp10430

	tmp10431 := MakeNative(func(__e *ControlFlow) {
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
		tmp10438 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V2467)

		if True == tmp10438 {
			tmp10433 := MakeNative(func(__e *ControlFlow) {
				Store := __e.Get(1)
				_ = Store
				tmp10434 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp10434

				tmp10435 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimNS2Value(symshen_4findall_1h), V2463, V2464, V2465, Store, V2466, V2467, V2468, V2469)
					return
				}, 0)

				tmp10436 := Call(__e, PrimNS2Value(symis), Store, Nil, V2466, V2467, V2468, tmp10435)

				__e.TailApply(PrimNS2Value(symshen_4gc), V2466, tmp10436)
				return

			}, 1)

			tmp10437 := Call(__e, PrimNS2Value(symshen_4newpv), V2466)

			__e.TailApply(tmp10433, tmp10437)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 7)

	tmp10439 := Call(__e, PrimNS2Value(symdef), symfindall, tmp10431)

	_ = tmp10439

	tmp10440 := MakeNative(func(__e *ControlFlow) {
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
		tmp10441 := MakeNative(func(__e *ControlFlow) {
			C1894 := __e.Get(1)
			_ = C1894
			tmp10446 := PrimEqual(C1894, False)

			if True == tmp10446 {
				tmp10445 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V2475)

				if True == tmp10445 {
					tmp10444 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp10444

					__e.TailApply(PrimNS2Value(symis_b), V2472, V2473, V2474, V2475, V2476, V2477)
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

		tmp10451 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V2475)

		var ifres10447 Obj

		if True == tmp10451 {
			tmp10448 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp10448

			tmp10449 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(symshen_4overbind), V2470, V2473, V2474, V2475, V2476, V2477)
				return
			}, 0)

			tmp10450 := Call(__e, PrimNS2Value(symcall), V2471, V2474, V2475, V2476, tmp10449)

			ifres10447 = tmp10450

		} else {
			ifres10447 = False

		}

		__e.TailApply(tmp10441, ifres10447)
		return

	}, 8)

	tmp10452 := Call(__e, PrimNS2Value(symdef), symshen_4findall_1h, tmp10440)

	_ = tmp10452

	tmp10453 := MakeNative(func(__e *ControlFlow) {
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
		tmp10454 := Call(__e, PrimNS2Value(symshen_4deref), V2484, V2486)

		tmp10455 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2485, V2486)

		tmp10456 := PrimCons(tmp10454, tmp10455)

		tmp10457 := Call(__e, PrimNS2Value(symshen_4bindv), V2485, tmp10456, V2486)

		_ = tmp10457

		__e.Return(False)
		return

	}, 6)

	tmp10458 := Call(__e, PrimNS2Value(symdef), symshen_4overbind, tmp10453)

	_ = tmp10458

	tmp10459 := MakeNative(func(__e *ControlFlow) {
		V2492 := __e.Get(1)
		_ = V2492
		tmp10463 := PrimEqual(sym_7, V2492)

		if True == tmp10463 {
			__e.Return(PrimNS3Set(symshen_4_doccurs_d, True))
			return
		} else {
			tmp10462 := PrimEqual(sym_1, V2492)

			if True == tmp10462 {
				__e.Return(PrimNS3Set(symshen_4_doccurs_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), symoccurs_1check, tmp10459)
	return

}, 0)
