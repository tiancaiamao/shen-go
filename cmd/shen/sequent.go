package main

import . "github.com/tiancaiamao/cora/cora_go"

var SequentMain = MakeNative(func(__e *ControlFlow) {
	tmp10464 := MakeNative(func(__e *ControlFlow) {
		V3119 := __e.Get(1)
		_ = V3119
		tmp10465 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10467 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10467 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10485 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V3119)

		var ifres10468 Obj

		if True == tmp10485 {
			tmp10470 := MakeNative(func(__e *ControlFlow) {
				D := __e.Get(1)
				_ = D
				tmp10471 := MakeNative(func(__e *ControlFlow) {
					News2980 := __e.Get(1)
					_ = News2980
					tmp10472 := MakeNative(func(__e *ControlFlow) {
						Parseshen_4_5datatype_1rules_6 := __e.Get(1)
						_ = Parseshen_4_5datatype_1rules_6
						tmp10480 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5datatype_1rules_6)

						if True == tmp10480 {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						} else {
							tmp10474 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5datatype_1rules_6)

							tmp10475 := MakeNative(func(__e *ControlFlow) {
								Prolog := __e.Get(1)
								_ = Prolog
								tmp10476 := Call(__e, PrimNS2Value(symfn), D)

								__e.TailApply(PrimNS2Value(symshen_4remember_1datatype), D, tmp10476)
								return

							}, 1)

							tmp10477 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5datatype_1rules_6)

							tmp10478 := Call(__e, PrimNS2Value(symshen_4rules_1_6prolog), D, tmp10477)

							tmp10479 := Call(__e, tmp10475, tmp10478)

							__e.TailApply(PrimNS2Value(symshen_4comb), tmp10474, tmp10479)
							return

						}

					}, 1)

					tmp10481 := Call(__e, PrimNS2Value(symshen_4_5datatype_1rules_6), News2980)

					__e.TailApply(tmp10472, tmp10481)
					return

				}, 1)

				tmp10482 := Call(__e, PrimNS2Value(symshen_4tls), V3119)

				__e.TailApply(tmp10471, tmp10482)
				return

			}, 1)

			tmp10483 := Call(__e, PrimNS2Value(symshen_4hds), V3119)

			tmp10484 := Call(__e, tmp10470, tmp10483)

			ifres10468 = tmp10484

		} else {
			tmp10469 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres10468 = tmp10469

		}

		__e.TailApply(tmp10465, ifres10468)
		return

	}, 1)

	tmp10486 := Call(__e, PrimNS2Value(symdef), symshen_4_5datatype_6, tmp10464)

	_ = tmp10486

	tmp10487 := MakeNative(func(__e *ControlFlow) {
		V3120 := __e.Get(1)
		_ = V3120
		tmp10488 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10506 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10506 {
				tmp10490 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10492 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10492 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10493 := MakeNative(func(__e *ControlFlow) {
					Parse_5_b_6 := __e.Get(1)
					_ = Parse_5_b_6
					tmp10503 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5_b_6)

					if True == tmp10503 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10495 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5_b_6)

						tmp10501 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

						tmp10502 := Call(__e, PrimNS2Value(symempty_2), tmp10501)

						var ifres10496 Obj

						if True == tmp10502 {
							ifres10496 = Nil

						} else {
							tmp10497 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

							tmp10498 := Call(__e, PrimNS2Value(symshen_4app), tmp10497, MakeString("\n ..."), symshen_4r)

							tmp10499 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp10498)

							tmp10500 := PrimSimpleError(tmp10499)

							ifres10496 = tmp10500

						}

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10495, ifres10496)
						return

					}

				}, 1)

				tmp10504 := Call(__e, PrimNS2Value(sym_5_b_6), V3120)

				tmp10505 := Call(__e, tmp10493, tmp10504)

				__e.TailApply(tmp10490, tmp10505)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10507 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5datatype_1rule_6 := __e.Get(1)
			_ = Parseshen_4_5datatype_1rule_6
			tmp10517 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5datatype_1rule_6)

			if True == tmp10517 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10509 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5datatype_1rules_6 := __e.Get(1)
					_ = Parseshen_4_5datatype_1rules_6
					tmp10515 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5datatype_1rules_6)

					if True == tmp10515 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10511 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5datatype_1rules_6)

						tmp10512 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5datatype_1rule_6)

						tmp10513 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5datatype_1rules_6)

						tmp10514 := Call(__e, PrimNS2Value(symappend), tmp10512, tmp10513)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10511, tmp10514)
						return

					}

				}, 1)

				tmp10516 := Call(__e, PrimNS2Value(symshen_4_5datatype_1rules_6), Parseshen_4_5datatype_1rule_6)

				__e.TailApply(tmp10509, tmp10516)
				return

			}

		}, 1)

		tmp10518 := Call(__e, PrimNS2Value(symshen_4_5datatype_1rule_6), V3120)

		tmp10519 := Call(__e, tmp10507, tmp10518)

		__e.TailApply(tmp10488, tmp10519)
		return

	}, 1)

	tmp10520 := Call(__e, PrimNS2Value(symdef), symshen_4_5datatype_1rules_6, tmp10487)

	_ = tmp10520

	tmp10521 := MakeNative(func(__e *ControlFlow) {
		V3121 := __e.Get(1)
		_ = V3121
		tmp10522 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10534 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10534 {
				tmp10524 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10526 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10526 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10527 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5double_6 := __e.Get(1)
					_ = Parseshen_4_5double_6
					tmp10531 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5double_6)

					if True == tmp10531 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10529 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5double_6)

						tmp10530 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5double_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10529, tmp10530)
						return

					}

				}, 1)

				tmp10532 := Call(__e, PrimNS2Value(symshen_4_5double_6), V3121)

				tmp10533 := Call(__e, tmp10527, tmp10532)

				__e.TailApply(tmp10524, tmp10533)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10535 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5single_6 := __e.Get(1)
			_ = Parseshen_4_5single_6
			tmp10539 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5single_6)

			if True == tmp10539 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10537 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5single_6)

				tmp10538 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5single_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp10537, tmp10538)
				return

			}

		}, 1)

		tmp10540 := Call(__e, PrimNS2Value(symshen_4_5single_6), V3121)

		tmp10541 := Call(__e, tmp10535, tmp10540)

		__e.TailApply(tmp10522, tmp10541)
		return

	}, 1)

	tmp10542 := Call(__e, PrimNS2Value(symdef), symshen_4_5datatype_1rule_6, tmp10521)

	_ = tmp10542

	tmp10543 := MakeNative(func(__e *ControlFlow) {
		V3122 := __e.Get(1)
		_ = V3122
		tmp10544 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10546 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10546 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10547 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5sides_6 := __e.Get(1)
			_ = Parseshen_4_5sides_6
			tmp10573 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sides_6)

			if True == tmp10573 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10549 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5prems_6 := __e.Get(1)
					_ = Parseshen_4_5prems_6
					tmp10571 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5prems_6)

					if True == tmp10571 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10551 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5sng_6 := __e.Get(1)
							_ = Parseshen_4_5sng_6
							tmp10569 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sng_6)

							if True == tmp10569 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10553 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5conc_6 := __e.Get(1)
									_ = Parseshen_4_5conc_6
									tmp10567 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5conc_6)

									if True == tmp10567 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp10555 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5sc_6 := __e.Get(1)
											_ = Parseshen_4_5sc_6
											tmp10565 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

											if True == tmp10565 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp10557 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sc_6)

												tmp10558 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5sides_6)

												tmp10559 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5prems_6)

												tmp10560 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5conc_6)

												tmp10561 := PrimCons(tmp10560, Nil)

												tmp10562 := PrimCons(tmp10559, tmp10561)

												tmp10563 := PrimCons(tmp10558, tmp10562)

												tmp10564 := PrimCons(tmp10563, Nil)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp10557, tmp10564)
												return

											}

										}, 1)

										tmp10566 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5conc_6)

										__e.TailApply(tmp10555, tmp10566)
										return

									}

								}, 1)

								tmp10568 := Call(__e, PrimNS2Value(symshen_4_5conc_6), Parseshen_4_5sng_6)

								__e.TailApply(tmp10553, tmp10568)
								return

							}

						}, 1)

						tmp10570 := Call(__e, PrimNS2Value(symshen_4_5sng_6), Parseshen_4_5prems_6)

						__e.TailApply(tmp10551, tmp10570)
						return

					}

				}, 1)

				tmp10572 := Call(__e, PrimNS2Value(symshen_4_5prems_6), Parseshen_4_5sides_6)

				__e.TailApply(tmp10549, tmp10572)
				return

			}

		}, 1)

		tmp10574 := Call(__e, PrimNS2Value(symshen_4_5sides_6), V3122)

		tmp10575 := Call(__e, tmp10547, tmp10574)

		__e.TailApply(tmp10544, tmp10575)
		return

	}, 1)

	tmp10576 := Call(__e, PrimNS2Value(symdef), symshen_4_5single_6, tmp10543)

	_ = tmp10576

	tmp10577 := MakeNative(func(__e *ControlFlow) {
		V3123 := __e.Get(1)
		_ = V3123
		tmp10578 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10580 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10580 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10581 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5sides_6 := __e.Get(1)
			_ = Parseshen_4_5sides_6
			tmp10606 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sides_6)

			if True == tmp10606 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10583 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5formulae_6 := __e.Get(1)
					_ = Parseshen_4_5formulae_6
					tmp10604 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formulae_6)

					if True == tmp10604 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10585 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5dbl_6 := __e.Get(1)
							_ = Parseshen_4_5dbl_6
							tmp10602 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5dbl_6)

							if True == tmp10602 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10587 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5formula_6 := __e.Get(1)
									_ = Parseshen_4_5formula_6
									tmp10600 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

									if True == tmp10600 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp10589 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5sc_6 := __e.Get(1)
											_ = Parseshen_4_5sc_6
											tmp10598 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

											if True == tmp10598 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp10591 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sc_6)

												tmp10592 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5sides_6)

												tmp10593 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formulae_6)

												tmp10594 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

												tmp10595 := PrimCons(tmp10594, Nil)

												tmp10596 := PrimCons(Nil, tmp10595)

												tmp10597 := Call(__e, PrimNS2Value(symshen_4lr_1rule), tmp10592, tmp10593, tmp10596)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp10591, tmp10597)
												return

											}

										}, 1)

										tmp10599 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5formula_6)

										__e.TailApply(tmp10589, tmp10599)
										return

									}

								}, 1)

								tmp10601 := Call(__e, PrimNS2Value(symshen_4_5formula_6), Parseshen_4_5dbl_6)

								__e.TailApply(tmp10587, tmp10601)
								return

							}

						}, 1)

						tmp10603 := Call(__e, PrimNS2Value(symshen_4_5dbl_6), Parseshen_4_5formulae_6)

						__e.TailApply(tmp10585, tmp10603)
						return

					}

				}, 1)

				tmp10605 := Call(__e, PrimNS2Value(symshen_4_5formulae_6), Parseshen_4_5sides_6)

				__e.TailApply(tmp10583, tmp10605)
				return

			}

		}, 1)

		tmp10607 := Call(__e, PrimNS2Value(symshen_4_5sides_6), V3123)

		tmp10608 := Call(__e, tmp10581, tmp10607)

		__e.TailApply(tmp10578, tmp10608)
		return

	}, 1)

	tmp10609 := Call(__e, PrimNS2Value(symdef), symshen_4_5double_6, tmp10577)

	_ = tmp10609

	tmp10610 := MakeNative(func(__e *ControlFlow) {
		V3124 := __e.Get(1)
		_ = V3124
		tmp10611 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10630 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10630 {
				tmp10613 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10615 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10615 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10616 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5formula_6 := __e.Get(1)
					_ = Parseshen_4_5formula_6
					tmp10627 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

					if True == tmp10627 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10618 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5sc_6 := __e.Get(1)
							_ = Parseshen_4_5sc_6
							tmp10625 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

							if True == tmp10625 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10620 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sc_6)

								tmp10621 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

								tmp10622 := PrimCons(tmp10621, Nil)

								tmp10623 := PrimCons(Nil, tmp10622)

								tmp10624 := PrimCons(tmp10623, Nil)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10620, tmp10624)
								return

							}

						}, 1)

						tmp10626 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5formula_6)

						__e.TailApply(tmp10618, tmp10626)
						return

					}

				}, 1)

				tmp10628 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V3124)

				tmp10629 := Call(__e, tmp10616, tmp10628)

				__e.TailApply(tmp10613, tmp10629)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10631 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5formula_6 := __e.Get(1)
			_ = Parseshen_4_5formula_6
			tmp10647 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

			if True == tmp10647 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10633 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5sc_6 := __e.Get(1)
					_ = Parseshen_4_5sc_6
					tmp10645 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

					if True == tmp10645 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10635 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5formulae_6 := __e.Get(1)
							_ = Parseshen_4_5formulae_6
							tmp10643 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formulae_6)

							if True == tmp10643 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10637 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5formulae_6)

								tmp10638 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

								tmp10639 := PrimCons(tmp10638, Nil)

								tmp10640 := PrimCons(Nil, tmp10639)

								tmp10641 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formulae_6)

								tmp10642 := PrimCons(tmp10640, tmp10641)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10637, tmp10642)
								return

							}

						}, 1)

						tmp10644 := Call(__e, PrimNS2Value(symshen_4_5formulae_6), Parseshen_4_5sc_6)

						__e.TailApply(tmp10635, tmp10644)
						return

					}

				}, 1)

				tmp10646 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5formula_6)

				__e.TailApply(tmp10633, tmp10646)
				return

			}

		}, 1)

		tmp10648 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V3124)

		tmp10649 := Call(__e, tmp10631, tmp10648)

		__e.TailApply(tmp10611, tmp10649)
		return

	}, 1)

	tmp10650 := Call(__e, PrimNS2Value(symdef), symshen_4_5formulae_6, tmp10610)

	_ = tmp10650

	tmp10651 := MakeNative(func(__e *ControlFlow) {
		V3125 := __e.Get(1)
		_ = V3125
		tmp10652 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10666 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10666 {
				tmp10654 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10656 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10656 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10657 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5formula_6 := __e.Get(1)
					_ = Parseshen_4_5formula_6
					tmp10663 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

					if True == tmp10663 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10659 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5formula_6)

						tmp10660 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

						tmp10661 := PrimCons(tmp10660, Nil)

						tmp10662 := PrimCons(Nil, tmp10661)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10659, tmp10662)
						return

					}

				}, 1)

				tmp10664 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V3125)

				tmp10665 := Call(__e, tmp10657, tmp10664)

				__e.TailApply(tmp10654, tmp10665)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10667 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5ass_6 := __e.Get(1)
			_ = Parseshen_4_5ass_6
			tmp10682 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5ass_6)

			if True == tmp10682 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10681 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5ass_6, sym_6_6)

				if True == tmp10681 {
					tmp10670 := MakeNative(func(__e *ControlFlow) {
						News2987 := __e.Get(1)
						_ = News2987
						tmp10671 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5formula_6 := __e.Get(1)
							_ = Parseshen_4_5formula_6
							tmp10678 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

							if True == tmp10678 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10673 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5formula_6)

								tmp10674 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5ass_6)

								tmp10675 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

								tmp10676 := PrimCons(tmp10675, Nil)

								tmp10677 := PrimCons(tmp10674, tmp10676)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10673, tmp10677)
								return

							}

						}, 1)

						tmp10679 := Call(__e, PrimNS2Value(symshen_4_5formula_6), News2987)

						__e.TailApply(tmp10671, tmp10679)
						return

					}, 1)

					tmp10680 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5ass_6)

					__e.TailApply(tmp10670, tmp10680)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp10683 := Call(__e, PrimNS2Value(symshen_4_5ass_6), V3125)

		tmp10684 := Call(__e, tmp10667, tmp10683)

		__e.TailApply(tmp10652, tmp10684)
		return

	}, 1)

	tmp10685 := Call(__e, PrimNS2Value(symdef), symshen_4_5conc_6, tmp10651)

	_ = tmp10685

	tmp10686 := MakeNative(func(__e *ControlFlow) {
		V3126 := __e.Get(1)
		_ = V3126
		tmp10687 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10698 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10698 {
				tmp10689 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10691 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10691 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10692 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp10695 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp10695 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10694 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10694, Nil)
						return

					}

				}, 1)

				tmp10696 := Call(__e, PrimNS2Value(sym_5e_6), V3126)

				tmp10697 := Call(__e, tmp10692, tmp10696)

				__e.TailApply(tmp10689, tmp10697)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10699 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5prem_6 := __e.Get(1)
			_ = Parseshen_4_5prem_6
			tmp10713 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5prem_6)

			if True == tmp10713 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10701 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5sc_6 := __e.Get(1)
					_ = Parseshen_4_5sc_6
					tmp10711 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

					if True == tmp10711 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10703 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5prems_6 := __e.Get(1)
							_ = Parseshen_4_5prems_6
							tmp10709 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5prems_6)

							if True == tmp10709 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10705 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5prems_6)

								tmp10706 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5prem_6)

								tmp10707 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5prems_6)

								tmp10708 := PrimCons(tmp10706, tmp10707)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10705, tmp10708)
								return

							}

						}, 1)

						tmp10710 := Call(__e, PrimNS2Value(symshen_4_5prems_6), Parseshen_4_5sc_6)

						__e.TailApply(tmp10703, tmp10710)
						return

					}

				}, 1)

				tmp10712 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5prem_6)

				__e.TailApply(tmp10701, tmp10712)
				return

			}

		}, 1)

		tmp10714 := Call(__e, PrimNS2Value(symshen_4_5prem_6), V3126)

		tmp10715 := Call(__e, tmp10699, tmp10714)

		__e.TailApply(tmp10687, tmp10715)
		return

	}, 1)

	tmp10716 := Call(__e, PrimNS2Value(symdef), symshen_4_5prems_6, tmp10686)

	_ = tmp10716

	tmp10717 := MakeNative(func(__e *ControlFlow) {
		V3127 := __e.Get(1)
		_ = V3127
		tmp10718 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10753 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10753 {
				tmp10720 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10734 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10734 {
						tmp10722 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp10724 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp10724 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp10725 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5formula_6 := __e.Get(1)
							_ = Parseshen_4_5formula_6
							tmp10731 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

							if True == tmp10731 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10727 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5formula_6)

								tmp10728 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

								tmp10729 := PrimCons(tmp10728, Nil)

								tmp10730 := PrimCons(Nil, tmp10729)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10727, tmp10730)
								return

							}

						}, 1)

						tmp10732 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V3127)

						tmp10733 := Call(__e, tmp10725, tmp10732)

						__e.TailApply(tmp10722, tmp10733)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10735 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5ass_6 := __e.Get(1)
					_ = Parseshen_4_5ass_6
					tmp10750 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5ass_6)

					if True == tmp10750 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10749 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5ass_6, sym_6_6)

						if True == tmp10749 {
							tmp10738 := MakeNative(func(__e *ControlFlow) {
								News2991 := __e.Get(1)
								_ = News2991
								tmp10739 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5formula_6 := __e.Get(1)
									_ = Parseshen_4_5formula_6
									tmp10746 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

									if True == tmp10746 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp10741 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5formula_6)

										tmp10742 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5ass_6)

										tmp10743 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

										tmp10744 := PrimCons(tmp10743, Nil)

										tmp10745 := PrimCons(tmp10742, tmp10744)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp10741, tmp10745)
										return

									}

								}, 1)

								tmp10747 := Call(__e, PrimNS2Value(symshen_4_5formula_6), News2991)

								__e.TailApply(tmp10739, tmp10747)
								return

							}, 1)

							tmp10748 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5ass_6)

							__e.TailApply(tmp10738, tmp10748)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}

				}, 1)

				tmp10751 := Call(__e, PrimNS2Value(symshen_4_5ass_6), V3127)

				tmp10752 := Call(__e, tmp10735, tmp10751)

				__e.TailApply(tmp10720, tmp10752)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10760 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V3127, sym_b)

		var ifres10754 Obj

		if True == tmp10760 {
			tmp10756 := MakeNative(func(__e *ControlFlow) {
				News2990 := __e.Get(1)
				_ = News2990
				tmp10757 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2990)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp10757, sym_b)
				return

			}, 1)

			tmp10758 := Call(__e, PrimNS2Value(symshen_4tls), V3127)

			tmp10759 := Call(__e, tmp10756, tmp10758)

			ifres10754 = tmp10759

		} else {
			tmp10755 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres10754 = tmp10755

		}

		__e.TailApply(tmp10718, ifres10754)
		return

	}, 1)

	tmp10761 := Call(__e, PrimNS2Value(symdef), symshen_4_5prem_6, tmp10717)

	_ = tmp10761

	tmp10762 := MakeNative(func(__e *ControlFlow) {
		V3128 := __e.Get(1)
		_ = V3128
		tmp10763 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10785 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10785 {
				tmp10765 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10776 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10776 {
						tmp10767 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp10769 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp10769 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp10770 := MakeNative(func(__e *ControlFlow) {
							Parse_5e_6 := __e.Get(1)
							_ = Parse_5e_6
							tmp10773 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

							if True == tmp10773 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10772 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10772, Nil)
								return

							}

						}, 1)

						tmp10774 := Call(__e, PrimNS2Value(sym_5e_6), V3128)

						tmp10775 := Call(__e, tmp10770, tmp10774)

						__e.TailApply(tmp10767, tmp10775)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10777 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5formula_6 := __e.Get(1)
					_ = Parseshen_4_5formula_6
					tmp10782 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

					if True == tmp10782 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10779 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5formula_6)

						tmp10780 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

						tmp10781 := PrimCons(tmp10780, Nil)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10779, tmp10781)
						return

					}

				}, 1)

				tmp10783 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V3128)

				tmp10784 := Call(__e, tmp10777, tmp10783)

				__e.TailApply(tmp10765, tmp10784)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10786 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5formula_6 := __e.Get(1)
			_ = Parseshen_4_5formula_6
			tmp10800 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5formula_6)

			if True == tmp10800 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10788 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5iscomma_6 := __e.Get(1)
					_ = Parseshen_4_5iscomma_6
					tmp10798 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5iscomma_6)

					if True == tmp10798 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10790 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5ass_6 := __e.Get(1)
							_ = Parseshen_4_5ass_6
							tmp10796 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5ass_6)

							if True == tmp10796 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10792 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5ass_6)

								tmp10793 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5formula_6)

								tmp10794 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5ass_6)

								tmp10795 := PrimCons(tmp10793, tmp10794)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10792, tmp10795)
								return

							}

						}, 1)

						tmp10797 := Call(__e, PrimNS2Value(symshen_4_5ass_6), Parseshen_4_5iscomma_6)

						__e.TailApply(tmp10790, tmp10797)
						return

					}

				}, 1)

				tmp10799 := Call(__e, PrimNS2Value(symshen_4_5iscomma_6), Parseshen_4_5formula_6)

				__e.TailApply(tmp10788, tmp10799)
				return

			}

		}, 1)

		tmp10801 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V3128)

		tmp10802 := Call(__e, tmp10786, tmp10801)

		__e.TailApply(tmp10763, tmp10802)
		return

	}, 1)

	tmp10803 := Call(__e, PrimNS2Value(symdef), symshen_4_5ass_6, tmp10762)

	_ = tmp10803

	tmp10804 := MakeNative(func(__e *ControlFlow) {
		V3129 := __e.Get(1)
		_ = V3129
		tmp10805 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10807 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10807 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10819 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V3129)

		var ifres10808 Obj

		if True == tmp10819 {
			tmp10810 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp10811 := MakeNative(func(__e *ControlFlow) {
					News2994 := __e.Get(1)
					_ = News2994
					tmp10814 := PrimIntern(MakeString(","))

					tmp10815 := PrimEqual(X, tmp10814)

					if True == tmp10815 {
						tmp10813 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2994)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10813, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10816 := Call(__e, PrimNS2Value(symshen_4tls), V3129)

				__e.TailApply(tmp10811, tmp10816)
				return

			}, 1)

			tmp10817 := Call(__e, PrimNS2Value(symshen_4hds), V3129)

			tmp10818 := Call(__e, tmp10810, tmp10817)

			ifres10808 = tmp10818

		} else {
			tmp10809 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres10808 = tmp10809

		}

		__e.TailApply(tmp10805, ifres10808)
		return

	}, 1)

	tmp10820 := Call(__e, PrimNS2Value(symdef), symshen_4_5iscomma_6, tmp10804)

	_ = tmp10820

	tmp10821 := MakeNative(func(__e *ControlFlow) {
		V3130 := __e.Get(1)
		_ = V3130
		tmp10822 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10834 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10834 {
				tmp10824 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10826 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10826 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10827 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5expr_6 := __e.Get(1)
					_ = Parseshen_4_5expr_6
					tmp10831 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5expr_6)

					if True == tmp10831 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10829 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5expr_6)

						tmp10830 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5expr_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10829, tmp10830)
						return

					}

				}, 1)

				tmp10832 := Call(__e, PrimNS2Value(symshen_4_5expr_6), V3130)

				tmp10833 := Call(__e, tmp10827, tmp10832)

				__e.TailApply(tmp10824, tmp10833)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10835 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5expr_6 := __e.Get(1)
			_ = Parseshen_4_5expr_6
			tmp10854 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5expr_6)

			if True == tmp10854 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10837 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5iscolon_6 := __e.Get(1)
					_ = Parseshen_4_5iscolon_6
					tmp10852 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5iscolon_6)

					if True == tmp10852 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10839 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5type_6 := __e.Get(1)
							_ = Parseshen_4_5type_6
							tmp10850 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5type_6)

							if True == tmp10850 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp10841 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5type_6)

								tmp10842 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5expr_6)

								tmp10843 := Call(__e, PrimNS2Value(symshen_4curry), tmp10842)

								tmp10844 := PrimIntern(MakeString(":"))

								tmp10845 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5type_6)

								tmp10846 := Call(__e, PrimNS2Value(symshen_4rectify_1type), tmp10845)

								tmp10847 := PrimCons(tmp10846, Nil)

								tmp10848 := PrimCons(tmp10844, tmp10847)

								tmp10849 := PrimCons(tmp10843, tmp10848)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp10841, tmp10849)
								return

							}

						}, 1)

						tmp10851 := Call(__e, PrimNS2Value(symshen_4_5type_6), Parseshen_4_5iscolon_6)

						__e.TailApply(tmp10839, tmp10851)
						return

					}

				}, 1)

				tmp10853 := Call(__e, PrimNS2Value(symshen_4_5iscolon_6), Parseshen_4_5expr_6)

				__e.TailApply(tmp10837, tmp10853)
				return

			}

		}, 1)

		tmp10855 := Call(__e, PrimNS2Value(symshen_4_5expr_6), V3130)

		tmp10856 := Call(__e, tmp10835, tmp10855)

		__e.TailApply(tmp10822, tmp10856)
		return

	}, 1)

	tmp10857 := Call(__e, PrimNS2Value(symdef), symshen_4_5formula_6, tmp10821)

	_ = tmp10857

	tmp10858 := MakeNative(func(__e *ControlFlow) {
		V3131 := __e.Get(1)
		_ = V3131
		tmp10859 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10861 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10861 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10873 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V3131)

		var ifres10862 Obj

		if True == tmp10873 {
			tmp10864 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp10865 := MakeNative(func(__e *ControlFlow) {
					News2997 := __e.Get(1)
					_ = News2997
					tmp10868 := PrimIntern(MakeString(":"))

					tmp10869 := PrimEqual(X, tmp10868)

					if True == tmp10869 {
						tmp10867 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2997)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10867, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10870 := Call(__e, PrimNS2Value(symshen_4tls), V3131)

				__e.TailApply(tmp10865, tmp10870)
				return

			}, 1)

			tmp10871 := Call(__e, PrimNS2Value(symshen_4hds), V3131)

			tmp10872 := Call(__e, tmp10864, tmp10871)

			ifres10862 = tmp10872

		} else {
			tmp10863 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres10862 = tmp10863

		}

		__e.TailApply(tmp10859, ifres10862)
		return

	}, 1)

	tmp10874 := Call(__e, PrimNS2Value(symdef), symshen_4_5iscolon_6, tmp10858)

	_ = tmp10874

	tmp10875 := MakeNative(func(__e *ControlFlow) {
		V3132 := __e.Get(1)
		_ = V3132
		tmp10876 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10887 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10887 {
				tmp10878 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10880 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10880 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10881 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp10884 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp10884 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10883 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10883, Nil)
						return

					}

				}, 1)

				tmp10885 := Call(__e, PrimNS2Value(sym_5e_6), V3132)

				tmp10886 := Call(__e, tmp10881, tmp10885)

				__e.TailApply(tmp10878, tmp10886)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10888 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5side_6 := __e.Get(1)
			_ = Parseshen_4_5side_6
			tmp10898 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5side_6)

			if True == tmp10898 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp10890 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5sides_6 := __e.Get(1)
					_ = Parseshen_4_5sides_6
					tmp10896 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sides_6)

					if True == tmp10896 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp10892 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sides_6)

						tmp10893 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5side_6)

						tmp10894 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5sides_6)

						tmp10895 := PrimCons(tmp10893, tmp10894)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp10892, tmp10895)
						return

					}

				}, 1)

				tmp10897 := Call(__e, PrimNS2Value(symshen_4_5sides_6), Parseshen_4_5side_6)

				__e.TailApply(tmp10890, tmp10897)
				return

			}

		}, 1)

		tmp10899 := Call(__e, PrimNS2Value(symshen_4_5side_6), V3132)

		tmp10900 := Call(__e, tmp10888, tmp10899)

		__e.TailApply(tmp10876, tmp10900)
		return

	}, 1)

	tmp10901 := Call(__e, PrimNS2Value(symdef), symshen_4_5sides_6, tmp10875)

	_ = tmp10901

	tmp10902 := MakeNative(func(__e *ControlFlow) {
		V3133 := __e.Get(1)
		_ = V3133
		tmp10903 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp10955 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp10955 {
				tmp10905 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp10932 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp10932 {
						tmp10907 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp10909 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp10909 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp10931 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V3133, symshen_4let_b)

						var ifres10910 Obj

						if True == tmp10931 {
							tmp10912 := MakeNative(func(__e *ControlFlow) {
								News3005 := __e.Get(1)
								_ = News3005
								tmp10928 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3005)

								if True == tmp10928 {
									tmp10914 := MakeNative(func(__e *ControlFlow) {
										X := __e.Get(1)
										_ = X
										tmp10915 := MakeNative(func(__e *ControlFlow) {
											News3006 := __e.Get(1)
											_ = News3006
											tmp10925 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3006)

											if True == tmp10925 {
												tmp10917 := MakeNative(func(__e *ControlFlow) {
													Y := __e.Get(1)
													_ = Y
													tmp10918 := MakeNative(func(__e *ControlFlow) {
														News3007 := __e.Get(1)
														_ = News3007
														tmp10919 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3007)

														tmp10920 := PrimCons(Y, Nil)

														tmp10921 := PrimCons(X, tmp10920)

														tmp10922 := PrimCons(symshen_4let_b, tmp10921)

														__e.TailApply(PrimNS2Value(symshen_4comb), tmp10919, tmp10922)
														return

													}, 1)

													tmp10923 := Call(__e, PrimNS2Value(symshen_4tls), News3006)

													__e.TailApply(tmp10918, tmp10923)
													return

												}, 1)

												tmp10924 := Call(__e, PrimNS2Value(symshen_4hds), News3006)

												__e.TailApply(tmp10917, tmp10924)
												return

											} else {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp10926 := Call(__e, PrimNS2Value(symshen_4tls), News3005)

										__e.TailApply(tmp10915, tmp10926)
										return

									}, 1)

									tmp10927 := Call(__e, PrimNS2Value(symshen_4hds), News3005)

									__e.TailApply(tmp10914, tmp10927)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp10929 := Call(__e, PrimNS2Value(symshen_4tls), V3133)

							tmp10930 := Call(__e, tmp10912, tmp10929)

							ifres10910 = tmp10930

						} else {
							tmp10911 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

							ifres10910 = tmp10911

						}

						__e.TailApply(tmp10907, ifres10910)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp10954 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V3133, symlet)

				var ifres10933 Obj

				if True == tmp10954 {
					tmp10935 := MakeNative(func(__e *ControlFlow) {
						News3002 := __e.Get(1)
						_ = News3002
						tmp10951 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3002)

						if True == tmp10951 {
							tmp10937 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp10938 := MakeNative(func(__e *ControlFlow) {
									News3003 := __e.Get(1)
									_ = News3003
									tmp10948 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3003)

									if True == tmp10948 {
										tmp10940 := MakeNative(func(__e *ControlFlow) {
											Y := __e.Get(1)
											_ = Y
											tmp10941 := MakeNative(func(__e *ControlFlow) {
												News3004 := __e.Get(1)
												_ = News3004
												tmp10942 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3004)

												tmp10943 := PrimCons(Y, Nil)

												tmp10944 := PrimCons(X, tmp10943)

												tmp10945 := PrimCons(symlet, tmp10944)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp10942, tmp10945)
												return

											}, 1)

											tmp10946 := Call(__e, PrimNS2Value(symshen_4tls), News3003)

											__e.TailApply(tmp10941, tmp10946)
											return

										}, 1)

										tmp10947 := Call(__e, PrimNS2Value(symshen_4hds), News3003)

										__e.TailApply(tmp10940, tmp10947)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp10949 := Call(__e, PrimNS2Value(symshen_4tls), News3002)

								__e.TailApply(tmp10938, tmp10949)
								return

							}, 1)

							tmp10950 := Call(__e, PrimNS2Value(symshen_4hds), News3002)

							__e.TailApply(tmp10937, tmp10950)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp10952 := Call(__e, PrimNS2Value(symshen_4tls), V3133)

					tmp10953 := Call(__e, tmp10935, tmp10952)

					ifres10933 = tmp10953

				} else {
					tmp10934 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

					ifres10933 = tmp10934

				}

				__e.TailApply(tmp10905, ifres10933)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp10970 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V3133, symif)

		var ifres10956 Obj

		if True == tmp10970 {
			tmp10958 := MakeNative(func(__e *ControlFlow) {
				News3000 := __e.Get(1)
				_ = News3000
				tmp10967 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3000)

				if True == tmp10967 {
					tmp10960 := MakeNative(func(__e *ControlFlow) {
						P := __e.Get(1)
						_ = P
						tmp10961 := MakeNative(func(__e *ControlFlow) {
							News3001 := __e.Get(1)
							_ = News3001
							tmp10962 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3001)

							tmp10963 := PrimCons(P, Nil)

							tmp10964 := PrimCons(symif, tmp10963)

							__e.TailApply(PrimNS2Value(symshen_4comb), tmp10962, tmp10964)
							return

						}, 1)

						tmp10965 := Call(__e, PrimNS2Value(symshen_4tls), News3000)

						__e.TailApply(tmp10961, tmp10965)
						return

					}, 1)

					tmp10966 := Call(__e, PrimNS2Value(symshen_4hds), News3000)

					__e.TailApply(tmp10960, tmp10966)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}, 1)

			tmp10968 := Call(__e, PrimNS2Value(symshen_4tls), V3133)

			tmp10969 := Call(__e, tmp10958, tmp10968)

			ifres10956 = tmp10969

		} else {
			tmp10957 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres10956 = tmp10957

		}

		__e.TailApply(tmp10903, ifres10956)
		return

	}, 1)

	tmp10971 := Call(__e, PrimNS2Value(symdef), symshen_4_5side_6, tmp10902)

	_ = tmp10971

	tmp10972 := MakeNative(func(__e *ControlFlow) {
		V3140 := __e.Get(1)
		_ = V3140
		V3141 := __e.Get(2)
		_ = V3141
		V3142 := __e.Get(3)
		_ = V3142
		tmp11008 := PrimIsPair(V3142)

		var ifres10995 Obj

		if True == tmp11008 {
			tmp11006 := PrimHead(V3142)

			tmp11007 := PrimEqual(Nil, tmp11006)

			var ifres10997 Obj

			if True == tmp11007 {
				tmp11004 := PrimTail(V3142)

				tmp11005 := PrimIsPair(tmp11004)

				var ifres10999 Obj

				if True == tmp11005 {
					tmp11001 := PrimTail(V3142)

					tmp11002 := PrimTail(tmp11001)

					tmp11003 := PrimEqual(Nil, tmp11002)

					var ifres11000 Obj

					if True == tmp11003 {
						ifres11000 = True

					} else {
						ifres11000 = False

					}

					ifres10999 = ifres11000

				} else {
					ifres10999 = False

				}

				var ifres10998 Obj

				if True == ifres10999 {
					ifres10998 = True

				} else {
					ifres10998 = False

				}

				ifres10997 = ifres10998

			} else {
				ifres10997 = False

			}

			var ifres10996 Obj

			if True == ifres10997 {
				ifres10996 = True

			} else {
				ifres10996 = False

			}

			ifres10995 = ifres10996

		} else {
			ifres10995 = False

		}

		if True == ifres10995 {
			tmp10974 := MakeNative(func(__e *ControlFlow) {
				P := __e.Get(1)
				_ = P
				tmp10975 := MakeNative(func(__e *ControlFlow) {
					LConc := __e.Get(1)
					_ = LConc
					tmp10976 := MakeNative(func(__e *ControlFlow) {
						LPrem := __e.Get(1)
						_ = LPrem
						tmp10977 := MakeNative(func(__e *ControlFlow) {
							Left := __e.Get(1)
							_ = Left
							tmp10978 := MakeNative(func(__e *ControlFlow) {
								Right := __e.Get(1)
								_ = Right
								tmp10979 := PrimCons(Left, Nil)

								__e.Return(PrimCons(Right, tmp10979))
								return

							}, 1)

							tmp10980 := PrimCons(V3142, Nil)

							tmp10981 := PrimCons(V3141, tmp10980)

							tmp10982 := PrimCons(V3140, tmp10981)

							__e.TailApply(tmp10978, tmp10982)
							return

						}, 1)

						tmp10983 := PrimCons(LPrem, Nil)

						tmp10984 := PrimCons(LConc, Nil)

						tmp10985 := PrimCons(tmp10983, tmp10984)

						tmp10986 := PrimCons(V3140, tmp10985)

						__e.TailApply(tmp10977, tmp10986)
						return

					}, 1)

					tmp10987 := Call(__e, PrimNS2Value(symshen_4coll_1formulae), V3141)

					tmp10988 := PrimCons(P, Nil)

					tmp10989 := PrimCons(tmp10987, tmp10988)

					__e.TailApply(tmp10976, tmp10989)
					return

				}, 1)

				tmp10990 := PrimTail(V3142)

				tmp10991 := PrimCons(P, Nil)

				tmp10992 := PrimCons(tmp10990, tmp10991)

				__e.TailApply(tmp10975, tmp10992)
				return

			}, 1)

			tmp10993 := Call(__e, PrimNS2Value(symprotect), symP)

			tmp10994 := Call(__e, PrimNS2Value(symgensym), tmp10993)

			__e.TailApply(tmp10974, tmp10994)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
			return
		}

	}, 3)

	tmp11009 := Call(__e, PrimNS2Value(symdef), symshen_4lr_1rule, tmp10972)

	_ = tmp11009

	tmp11010 := MakeNative(func(__e *ControlFlow) {
		V3145 := __e.Get(1)
		_ = V3145
		tmp11039 := PrimEqual(Nil, V3145)

		if True == tmp11039 {
			__e.Return(Nil)
			return
		} else {
			tmp11038 := PrimIsPair(V3145)

			var ifres11018 Obj

			if True == tmp11038 {
				tmp11036 := PrimHead(V3145)

				tmp11037 := PrimIsPair(tmp11036)

				var ifres11020 Obj

				if True == tmp11037 {
					tmp11033 := PrimHead(V3145)

					tmp11034 := PrimHead(tmp11033)

					tmp11035 := PrimEqual(Nil, tmp11034)

					var ifres11022 Obj

					if True == tmp11035 {
						tmp11030 := PrimHead(V3145)

						tmp11031 := PrimTail(tmp11030)

						tmp11032 := PrimIsPair(tmp11031)

						var ifres11024 Obj

						if True == tmp11032 {
							tmp11026 := PrimHead(V3145)

							tmp11027 := PrimTail(tmp11026)

							tmp11028 := PrimTail(tmp11027)

							tmp11029 := PrimEqual(Nil, tmp11028)

							var ifres11025 Obj

							if True == tmp11029 {
								ifres11025 = True

							} else {
								ifres11025 = False

							}

							ifres11024 = ifres11025

						} else {
							ifres11024 = False

						}

						var ifres11023 Obj

						if True == ifres11024 {
							ifres11023 = True

						} else {
							ifres11023 = False

						}

						ifres11022 = ifres11023

					} else {
						ifres11022 = False

					}

					var ifres11021 Obj

					if True == ifres11022 {
						ifres11021 = True

					} else {
						ifres11021 = False

					}

					ifres11020 = ifres11021

				} else {
					ifres11020 = False

				}

				var ifres11019 Obj

				if True == ifres11020 {
					ifres11019 = True

				} else {
					ifres11019 = False

				}

				ifres11018 = ifres11019

			} else {
				ifres11018 = False

			}

			if True == ifres11018 {
				tmp11013 := PrimHead(V3145)

				tmp11014 := PrimTail(tmp11013)

				tmp11015 := PrimHead(tmp11014)

				tmp11016 := PrimTail(V3145)

				tmp11017 := Call(__e, PrimNS2Value(symshen_4coll_1formulae), tmp11016)

				__e.Return(PrimCons(tmp11015, tmp11017))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
				return
			}

		}

	}, 1)

	tmp11040 := Call(__e, PrimNS2Value(symdef), symshen_4coll_1formulae, tmp11010)

	_ = tmp11040

	tmp11041 := MakeNative(func(__e *ControlFlow) {
		V3146 := __e.Get(1)
		_ = V3146
		tmp11042 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp11044 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp11044 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp11057 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V3146)

		var ifres11045 Obj

		if True == tmp11057 {
			tmp11047 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp11048 := MakeNative(func(__e *ControlFlow) {
					News3009 := __e.Get(1)
					_ = News3009
					tmp11052 := Call(__e, PrimNS2Value(symshen_4key_1in_1sequent_1calculus_2), X)

					tmp11053 := PrimNot(tmp11052)

					if True == tmp11053 {
						tmp11050 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3009)

						tmp11051 := Call(__e, PrimNS2Value(symmacroexpand), X)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp11050, tmp11051)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp11054 := Call(__e, PrimNS2Value(symshen_4tls), V3146)

				__e.TailApply(tmp11048, tmp11054)
				return

			}, 1)

			tmp11055 := Call(__e, PrimNS2Value(symshen_4hds), V3146)

			tmp11056 := Call(__e, tmp11047, tmp11055)

			ifres11045 = tmp11056

		} else {
			tmp11046 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres11045 = tmp11046

		}

		__e.TailApply(tmp11042, ifres11045)
		return

	}, 1)

	tmp11058 := Call(__e, PrimNS2Value(symdef), symshen_4_5expr_6, tmp11041)

	_ = tmp11058

	tmp11059 := MakeNative(func(__e *ControlFlow) {
		V3147 := __e.Get(1)
		_ = V3147
		tmp11066 := PrimIntern(MakeString(";"))

		tmp11067 := PrimIntern(MakeString(","))

		tmp11068 := PrimIntern(MakeString(":"))

		tmp11069 := PrimCons(sym_5_1_1, Nil)

		tmp11070 := PrimCons(tmp11068, tmp11069)

		tmp11071 := PrimCons(tmp11067, tmp11070)

		tmp11072 := PrimCons(tmp11066, tmp11071)

		tmp11073 := PrimCons(sym_6_6, tmp11072)

		tmp11074 := Call(__e, PrimNS2Value(symelement_2), V3147, tmp11073)

		if True == tmp11074 {
			__e.Return(True)
			return
		} else {
			tmp11065 := Call(__e, PrimNS2Value(symshen_4sng_2), V3147)

			var ifres11062 Obj

			if True == tmp11065 {
				ifres11062 = True

			} else {
				tmp11064 := Call(__e, PrimNS2Value(symshen_4dbl_2), V3147)

				var ifres11063 Obj

				if True == tmp11064 {
					ifres11063 = True

				} else {
					ifres11063 = False

				}

				ifres11062 = ifres11063

			}

			if True == ifres11062 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp11075 := Call(__e, PrimNS2Value(symdef), symshen_4key_1in_1sequent_1calculus_2, tmp11059)

	_ = tmp11075

	tmp11076 := MakeNative(func(__e *ControlFlow) {
		V3148 := __e.Get(1)
		_ = V3148
		tmp11077 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp11079 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp11079 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp11080 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5expr_6 := __e.Get(1)
			_ = Parseshen_4_5expr_6
			tmp11084 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5expr_6)

			if True == tmp11084 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp11082 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5expr_6)

				tmp11083 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5expr_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp11082, tmp11083)
				return

			}

		}, 1)

		tmp11085 := Call(__e, PrimNS2Value(symshen_4_5expr_6), V3148)

		tmp11086 := Call(__e, tmp11080, tmp11085)

		__e.TailApply(tmp11077, tmp11086)
		return

	}, 1)

	tmp11087 := Call(__e, PrimNS2Value(symdef), symshen_4_5type_6, tmp11076)

	_ = tmp11087

	tmp11088 := MakeNative(func(__e *ControlFlow) {
		V3149 := __e.Get(1)
		_ = V3149
		tmp11089 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp11091 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp11091 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp11102 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V3149)

		var ifres11092 Obj

		if True == tmp11102 {
			tmp11094 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp11095 := MakeNative(func(__e *ControlFlow) {
					News3012 := __e.Get(1)
					_ = News3012
					tmp11098 := Call(__e, PrimNS2Value(symshen_4dbl_2), X)

					if True == tmp11098 {
						tmp11097 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3012)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp11097, X)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp11099 := Call(__e, PrimNS2Value(symshen_4tls), V3149)

				__e.TailApply(tmp11095, tmp11099)
				return

			}, 1)

			tmp11100 := Call(__e, PrimNS2Value(symshen_4hds), V3149)

			tmp11101 := Call(__e, tmp11094, tmp11100)

			ifres11092 = tmp11101

		} else {
			tmp11093 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres11092 = tmp11093

		}

		__e.TailApply(tmp11089, ifres11092)
		return

	}, 1)

	tmp11103 := Call(__e, PrimNS2Value(symdef), symshen_4_5dbl_6, tmp11088)

	_ = tmp11103

	tmp11104 := MakeNative(func(__e *ControlFlow) {
		V3150 := __e.Get(1)
		_ = V3150
		tmp11105 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp11107 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp11107 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp11118 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V3150)

		var ifres11108 Obj

		if True == tmp11118 {
			tmp11110 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp11111 := MakeNative(func(__e *ControlFlow) {
					News3014 := __e.Get(1)
					_ = News3014
					tmp11114 := Call(__e, PrimNS2Value(symshen_4sng_2), X)

					if True == tmp11114 {
						tmp11113 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3014)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp11113, X)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp11115 := Call(__e, PrimNS2Value(symshen_4tls), V3150)

				__e.TailApply(tmp11111, tmp11115)
				return

			}, 1)

			tmp11116 := Call(__e, PrimNS2Value(symshen_4hds), V3150)

			tmp11117 := Call(__e, tmp11110, tmp11116)

			ifres11108 = tmp11117

		} else {
			tmp11109 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres11108 = tmp11109

		}

		__e.TailApply(tmp11105, ifres11108)
		return

	}, 1)

	tmp11119 := Call(__e, PrimNS2Value(symdef), symshen_4_5sng_6, tmp11104)

	_ = tmp11119

	tmp11120 := MakeNative(func(__e *ControlFlow) {
		V3151 := __e.Get(1)
		_ = V3151
		tmp11125 := PrimIsSymbol(V3151)

		if True == tmp11125 {
			tmp11123 := PrimStr(V3151)

			tmp11124 := Call(__e, PrimNS2Value(symshen_4sng_1h_2), tmp11123)

			if True == tmp11124 {
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

	tmp11126 := Call(__e, PrimNS2Value(symdef), symshen_4sng_2, tmp11120)

	_ = tmp11126

	tmp11127 := MakeNative(func(__e *ControlFlow) {
		V3154 := __e.Get(1)
		_ = V3154
		tmp11136 := PrimEqual(MakeString("___"), V3154)

		if True == tmp11136 {
			__e.Return(True)
			return
		} else {
			tmp11135 := Call(__e, PrimNS2Value(symshen_4_7string_2), V3154)

			var ifres11131 Obj

			if True == tmp11135 {
				tmp11133 := Call(__e, PrimNS2Value(symhdstr), V3154)

				tmp11134 := PrimEqual(MakeString("_"), tmp11133)

				var ifres11132 Obj

				if True == tmp11134 {
					ifres11132 = True

				} else {
					ifres11132 = False

				}

				ifres11131 = ifres11132

			} else {
				ifres11131 = False

			}

			if True == ifres11131 {
				tmp11130 := PrimTailString(V3154)

				__e.TailApply(PrimNS2Value(symshen_4sng_1h_2), tmp11130)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp11137 := Call(__e, PrimNS2Value(symdef), symshen_4sng_1h_2, tmp11127)

	_ = tmp11137

	tmp11138 := MakeNative(func(__e *ControlFlow) {
		V3155 := __e.Get(1)
		_ = V3155
		tmp11143 := PrimIsSymbol(V3155)

		if True == tmp11143 {
			tmp11141 := PrimStr(V3155)

			tmp11142 := Call(__e, PrimNS2Value(symshen_4dbl_1h_2), tmp11141)

			if True == tmp11142 {
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

	tmp11144 := Call(__e, PrimNS2Value(symdef), symshen_4dbl_2, tmp11138)

	_ = tmp11144

	tmp11145 := MakeNative(func(__e *ControlFlow) {
		V3158 := __e.Get(1)
		_ = V3158
		tmp11154 := PrimEqual(MakeString("==="), V3158)

		if True == tmp11154 {
			__e.Return(True)
			return
		} else {
			tmp11153 := Call(__e, PrimNS2Value(symshen_4_7string_2), V3158)

			var ifres11149 Obj

			if True == tmp11153 {
				tmp11151 := Call(__e, PrimNS2Value(symhdstr), V3158)

				tmp11152 := PrimEqual(MakeString("="), tmp11151)

				var ifres11150 Obj

				if True == tmp11152 {
					ifres11150 = True

				} else {
					ifres11150 = False

				}

				ifres11149 = ifres11150

			} else {
				ifres11149 = False

			}

			if True == ifres11149 {
				tmp11148 := PrimTailString(V3158)

				__e.TailApply(PrimNS2Value(symshen_4dbl_1h_2), tmp11148)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp11155 := Call(__e, PrimNS2Value(symdef), symshen_4dbl_1h_2, tmp11145)

	_ = tmp11155

	tmp11156 := MakeNative(func(__e *ControlFlow) {
		V3159 := __e.Get(1)
		_ = V3159
		V3160 := __e.Get(2)
		_ = V3160
		tmp11157 := PrimNS3Value(symshen_4_ddatatypes_d)

		tmp11158 := Call(__e, PrimNS2Value(symshen_4assoc_1_6), V3159, V3160, tmp11157)

		tmp11159 := PrimNS3Set(symshen_4_ddatatypes_d, tmp11158)

		_ = tmp11159

		tmp11160 := PrimNS3Value(symshen_4_dalldatatypes_d)

		tmp11161 := Call(__e, PrimNS2Value(symshen_4assoc_1_6), V3159, V3160, tmp11160)

		tmp11162 := PrimNS3Set(symshen_4_dalldatatypes_d, tmp11161)

		_ = tmp11162

		__e.Return(V3159)
		return

	}, 2)

	tmp11163 := Call(__e, PrimNS2Value(symdef), symshen_4remember_1datatype, tmp11156)

	_ = tmp11163

	tmp11164 := MakeNative(func(__e *ControlFlow) {
		V3161 := __e.Get(1)
		_ = V3161
		V3162 := __e.Get(2)
		_ = V3162
		tmp11165 := MakeNative(func(__e *ControlFlow) {
			Clauses := __e.Get(1)
			_ = Clauses
			tmp11166 := PrimCons(V3161, Clauses)

			tmp11167 := PrimCons(symdefprolog, tmp11166)

			__e.TailApply(PrimNS2Value(symeval), tmp11167)
			return

		}, 1)

		tmp11168 := MakeNative(func(__e *ControlFlow) {
			Rule := __e.Get(1)
			_ = Rule
			__e.TailApply(PrimNS2Value(symshen_4rule_1_6clause), Rule)
			return
		}, 1)

		tmp11169 := Call(__e, PrimNS2Value(symmapcan), tmp11168, V3162)

		__e.TailApply(tmp11165, tmp11169)
		return

	}, 2)

	tmp11170 := Call(__e, PrimNS2Value(symdef), symshen_4rules_1_6prolog, tmp11164)

	_ = tmp11170

	tmp11171 := MakeNative(func(__e *ControlFlow) {
		V3165 := __e.Get(1)
		_ = V3165
		tmp11247 := PrimIsPair(V3165)

		var ifres11211 Obj

		if True == tmp11247 {
			tmp11245 := PrimTail(V3165)

			tmp11246 := PrimIsPair(tmp11245)

			var ifres11213 Obj

			if True == tmp11246 {
				tmp11242 := PrimTail(V3165)

				tmp11243 := PrimTail(tmp11242)

				tmp11244 := PrimIsPair(tmp11243)

				var ifres11215 Obj

				if True == tmp11244 {
					tmp11238 := PrimTail(V3165)

					tmp11239 := PrimTail(tmp11238)

					tmp11240 := PrimHead(tmp11239)

					tmp11241 := PrimIsPair(tmp11240)

					var ifres11217 Obj

					if True == tmp11241 {
						tmp11233 := PrimTail(V3165)

						tmp11234 := PrimTail(tmp11233)

						tmp11235 := PrimHead(tmp11234)

						tmp11236 := PrimTail(tmp11235)

						tmp11237 := PrimIsPair(tmp11236)

						var ifres11219 Obj

						if True == tmp11237 {
							tmp11227 := PrimTail(V3165)

							tmp11228 := PrimTail(tmp11227)

							tmp11229 := PrimHead(tmp11228)

							tmp11230 := PrimTail(tmp11229)

							tmp11231 := PrimTail(tmp11230)

							tmp11232 := PrimEqual(Nil, tmp11231)

							var ifres11221 Obj

							if True == tmp11232 {
								tmp11223 := PrimTail(V3165)

								tmp11224 := PrimTail(tmp11223)

								tmp11225 := PrimTail(tmp11224)

								tmp11226 := PrimEqual(Nil, tmp11225)

								var ifres11222 Obj

								if True == tmp11226 {
									ifres11222 = True

								} else {
									ifres11222 = False

								}

								ifres11221 = ifres11222

							} else {
								ifres11221 = False

							}

							var ifres11220 Obj

							if True == ifres11221 {
								ifres11220 = True

							} else {
								ifres11220 = False

							}

							ifres11219 = ifres11220

						} else {
							ifres11219 = False

						}

						var ifres11218 Obj

						if True == ifres11219 {
							ifres11218 = True

						} else {
							ifres11218 = False

						}

						ifres11217 = ifres11218

					} else {
						ifres11217 = False

					}

					var ifres11216 Obj

					if True == ifres11217 {
						ifres11216 = True

					} else {
						ifres11216 = False

					}

					ifres11215 = ifres11216

				} else {
					ifres11215 = False

				}

				var ifres11214 Obj

				if True == ifres11215 {
					ifres11214 = True

				} else {
					ifres11214 = False

				}

				ifres11213 = ifres11214

			} else {
				ifres11213 = False

			}

			var ifres11212 Obj

			if True == ifres11213 {
				ifres11212 = True

			} else {
				ifres11212 = False

			}

			ifres11211 = ifres11212

		} else {
			ifres11211 = False

		}

		if True == ifres11211 {
			tmp11173 := MakeNative(func(__e *ControlFlow) {
				Constraints := __e.Get(1)
				_ = Constraints
				tmp11174 := MakeNative(func(__e *ControlFlow) {
					HypVs := __e.Get(1)
					_ = HypVs
					tmp11175 := MakeNative(func(__e *ControlFlow) {
						Active := __e.Get(1)
						_ = Active
						tmp11176 := MakeNative(func(__e *ControlFlow) {
							Head := __e.Get(1)
							_ = Head
							tmp11177 := MakeNative(func(__e *ControlFlow) {
								Goals := __e.Get(1)
								_ = Goals
								tmp11178 := PrimCons(sym_5_1_1, Nil)

								tmp11179 := PrimIntern(MakeString(";"))

								tmp11180 := PrimCons(tmp11179, Nil)

								tmp11181 := Call(__e, PrimNS2Value(symappend), Goals, tmp11180)

								tmp11182 := Call(__e, PrimNS2Value(symappend), tmp11178, tmp11181)

								__e.TailApply(PrimNS2Value(symappend), Head, tmp11182)
								return

							}, 1)

							tmp11183 := PrimTail(V3165)

							tmp11184 := PrimTail(tmp11183)

							tmp11185 := PrimHead(tmp11184)

							tmp11186 := PrimHead(tmp11185)

							tmp11187 := PrimHead(V3165)

							tmp11188 := PrimTail(V3165)

							tmp11189 := PrimHead(tmp11188)

							tmp11190 := Call(__e, PrimNS2Value(symshen_4goals), Constraints, tmp11186, tmp11187, tmp11189, HypVs, Active)

							__e.TailApply(tmp11177, tmp11190)
							return

						}, 1)

						tmp11191 := PrimTail(V3165)

						tmp11192 := PrimTail(tmp11191)

						tmp11193 := PrimHead(tmp11192)

						tmp11194 := PrimTail(tmp11193)

						tmp11195 := PrimHead(tmp11194)

						tmp11196 := Call(__e, PrimNS2Value(symshen_4compile_1consequent), tmp11195, HypVs)

						__e.TailApply(tmp11176, tmp11196)
						return

					}, 1)

					tmp11197 := PrimTail(V3165)

					tmp11198 := PrimTail(tmp11197)

					tmp11199 := PrimHead(tmp11198)

					tmp11200 := PrimTail(tmp11199)

					tmp11201 := PrimHead(tmp11200)

					tmp11202 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp11201)

					__e.TailApply(tmp11175, tmp11202)
					return

				}, 1)

				tmp11203 := PrimTail(V3165)

				tmp11204 := PrimTail(tmp11203)

				tmp11205 := PrimHead(tmp11204)

				tmp11206 := PrimHead(tmp11205)

				tmp11207 := Call(__e, PrimNS2Value(symlength), tmp11206)

				tmp11208 := PrimNumberAdd(MakeNumber(1), tmp11207)

				tmp11209 := Call(__e, PrimNS2Value(symshen_4nvars), tmp11208)

				__e.TailApply(tmp11174, tmp11209)
				return

			}, 1)

			tmp11210 := Call(__e, PrimNS2Value(symshen_4extract_1vars), V3165)

			__e.TailApply(tmp11173, tmp11210)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.rule->clause")))
			return
		}

	}, 1)

	tmp11248 := Call(__e, PrimNS2Value(symdef), symshen_4rule_1_6clause, tmp11171)

	_ = tmp11248

	tmp11249 := MakeNative(func(__e *ControlFlow) {
		V3172 := __e.Get(1)
		_ = V3172
		V3173 := __e.Get(2)
		_ = V3173
		tmp11254 := PrimIsPair(V3173)

		if True == tmp11254 {
			tmp11251 := Call(__e, PrimNS2Value(symshen_4optimise_1typing), V3172)

			tmp11252 := PrimHead(V3173)

			tmp11253 := PrimCons(tmp11252, Nil)

			__e.Return(PrimCons(tmp11251, tmp11253))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-consequent")))
			return
		}

	}, 2)

	tmp11255 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1consequent, tmp11249)

	_ = tmp11255

	tmp11256 := MakeNative(func(__e *ControlFlow) {
		V3174 := __e.Get(1)
		_ = V3174
		tmp11262 := PrimEqual(MakeNumber(0), V3174)

		if True == tmp11262 {
			__e.Return(Nil)
			return
		} else {
			tmp11258 := Call(__e, PrimNS2Value(symprotect), symV)

			tmp11259 := Call(__e, PrimNS2Value(symgensym), tmp11258)

			tmp11260 := PrimNumberSubtract(V3174, MakeNumber(1))

			tmp11261 := Call(__e, PrimNS2Value(symshen_4nvars), tmp11260)

			__e.Return(PrimCons(tmp11259, tmp11261))
			return

		}

	}, 1)

	tmp11263 := Call(__e, PrimNS2Value(symdef), symshen_4nvars, tmp11256)

	_ = tmp11263

	tmp11264 := MakeNative(func(__e *ControlFlow) {
		V3175 := __e.Get(1)
		_ = V3175
		tmp11300 := PrimIsPair(V3175)

		var ifres11279 Obj

		if True == tmp11300 {
			tmp11298 := PrimTail(V3175)

			tmp11299 := PrimIsPair(tmp11298)

			var ifres11281 Obj

			if True == tmp11299 {
				tmp11295 := PrimTail(V3175)

				tmp11296 := PrimTail(tmp11295)

				tmp11297 := PrimIsPair(tmp11296)

				var ifres11283 Obj

				if True == tmp11297 {
					tmp11291 := PrimTail(V3175)

					tmp11292 := PrimTail(tmp11291)

					tmp11293 := PrimTail(tmp11292)

					tmp11294 := PrimEqual(Nil, tmp11293)

					var ifres11285 Obj

					if True == tmp11294 {
						tmp11287 := PrimTail(V3175)

						tmp11288 := PrimHead(tmp11287)

						tmp11289 := PrimIntern(MakeString(":"))

						tmp11290 := PrimEqual(tmp11288, tmp11289)

						var ifres11286 Obj

						if True == tmp11290 {
							ifres11286 = True

						} else {
							ifres11286 = False

						}

						ifres11285 = ifres11286

					} else {
						ifres11285 = False

					}

					var ifres11284 Obj

					if True == ifres11285 {
						ifres11284 = True

					} else {
						ifres11284 = False

					}

					ifres11283 = ifres11284

				} else {
					ifres11283 = False

				}

				var ifres11282 Obj

				if True == ifres11283 {
					ifres11282 = True

				} else {
					ifres11282 = False

				}

				ifres11281 = ifres11282

			} else {
				ifres11281 = False

			}

			var ifres11280 Obj

			if True == ifres11281 {
				ifres11280 = True

			} else {
				ifres11280 = False

			}

			ifres11279 = ifres11280

		} else {
			ifres11279 = False

		}

		if True == ifres11279 {
			tmp11268 := PrimHead(V3175)

			tmp11269 := PrimTail(V3175)

			tmp11270 := PrimHead(tmp11269)

			tmp11271 := PrimTail(V3175)

			tmp11272 := PrimTail(tmp11271)

			tmp11273 := PrimCons(sym_7, tmp11272)

			tmp11274 := PrimCons(tmp11273, Nil)

			tmp11275 := PrimCons(tmp11270, tmp11274)

			tmp11276 := PrimCons(tmp11268, tmp11275)

			tmp11277 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), tmp11276)

			tmp11278 := PrimCons(tmp11277, Nil)

			__e.Return(PrimCons(sym_1, tmp11278))
			return

		} else {
			tmp11266 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), V3175)

			tmp11267 := PrimCons(tmp11266, Nil)

			__e.Return(PrimCons(sym_7, tmp11267))
			return

		}

	}, 1)

	tmp11301 := Call(__e, PrimNS2Value(symdef), symshen_4optimise_1typing, tmp11264)

	_ = tmp11301

	tmp11302 := MakeNative(func(__e *ControlFlow) {
		V3176 := __e.Get(1)
		_ = V3176
		tmp11392 := PrimIsPair(V3176)

		var ifres11379 Obj

		if True == tmp11392 {
			tmp11390 := PrimHead(V3176)

			tmp11391 := PrimEqual(sym_1, tmp11390)

			var ifres11381 Obj

			if True == tmp11391 {
				tmp11388 := PrimTail(V3176)

				tmp11389 := PrimIsPair(tmp11388)

				var ifres11383 Obj

				if True == tmp11389 {
					tmp11385 := PrimTail(V3176)

					tmp11386 := PrimTail(tmp11385)

					tmp11387 := PrimEqual(Nil, tmp11386)

					var ifres11384 Obj

					if True == tmp11387 {
						ifres11384 = True

					} else {
						ifres11384 = False

					}

					ifres11383 = ifres11384

				} else {
					ifres11383 = False

				}

				var ifres11382 Obj

				if True == ifres11383 {
					ifres11382 = True

				} else {
					ifres11382 = False

				}

				ifres11381 = ifres11382

			} else {
				ifres11381 = False

			}

			var ifres11380 Obj

			if True == ifres11381 {
				ifres11380 = True

			} else {
				ifres11380 = False

			}

			ifres11379 = ifres11380

		} else {
			ifres11379 = False

		}

		if True == ifres11379 {
			tmp11375 := PrimTail(V3176)

			tmp11376 := PrimHead(tmp11375)

			tmp11377 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), tmp11376)

			tmp11378 := PrimCons(tmp11377, Nil)

			__e.Return(PrimCons(sym_1, tmp11378))
			return

		} else {
			tmp11374 := PrimIsPair(V3176)

			var ifres11361 Obj

			if True == tmp11374 {
				tmp11372 := PrimHead(V3176)

				tmp11373 := PrimEqual(sym_7, tmp11372)

				var ifres11363 Obj

				if True == tmp11373 {
					tmp11370 := PrimTail(V3176)

					tmp11371 := PrimIsPair(tmp11370)

					var ifres11365 Obj

					if True == tmp11371 {
						tmp11367 := PrimTail(V3176)

						tmp11368 := PrimTail(tmp11367)

						tmp11369 := PrimEqual(Nil, tmp11368)

						var ifres11366 Obj

						if True == tmp11369 {
							ifres11366 = True

						} else {
							ifres11366 = False

						}

						ifres11365 = ifres11366

					} else {
						ifres11365 = False

					}

					var ifres11364 Obj

					if True == ifres11365 {
						ifres11364 = True

					} else {
						ifres11364 = False

					}

					ifres11363 = ifres11364

				} else {
					ifres11363 = False

				}

				var ifres11362 Obj

				if True == ifres11363 {
					ifres11362 = True

				} else {
					ifres11362 = False

				}

				ifres11361 = ifres11362

			} else {
				ifres11361 = False

			}

			if True == ifres11361 {
				tmp11357 := PrimTail(V3176)

				tmp11358 := PrimHead(tmp11357)

				tmp11359 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), tmp11358)

				tmp11360 := PrimCons(tmp11359, Nil)

				__e.Return(PrimCons(sym_7, tmp11360))
				return

			} else {
				tmp11356 := PrimIsPair(V3176)

				var ifres11337 Obj

				if True == tmp11356 {
					tmp11354 := PrimHead(V3176)

					tmp11355 := PrimEqual(symmode, tmp11354)

					var ifres11339 Obj

					if True == tmp11355 {
						tmp11352 := PrimTail(V3176)

						tmp11353 := PrimIsPair(tmp11352)

						var ifres11341 Obj

						if True == tmp11353 {
							tmp11349 := PrimTail(V3176)

							tmp11350 := PrimTail(tmp11349)

							tmp11351 := PrimIsPair(tmp11350)

							var ifres11343 Obj

							if True == tmp11351 {
								tmp11345 := PrimTail(V3176)

								tmp11346 := PrimTail(tmp11345)

								tmp11347 := PrimTail(tmp11346)

								tmp11348 := PrimEqual(Nil, tmp11347)

								var ifres11344 Obj

								if True == tmp11348 {
									ifres11344 = True

								} else {
									ifres11344 = False

								}

								ifres11343 = ifres11344

							} else {
								ifres11343 = False

							}

							var ifres11342 Obj

							if True == ifres11343 {
								ifres11342 = True

							} else {
								ifres11342 = False

							}

							ifres11341 = ifres11342

						} else {
							ifres11341 = False

						}

						var ifres11340 Obj

						if True == ifres11341 {
							ifres11340 = True

						} else {
							ifres11340 = False

						}

						ifres11339 = ifres11340

					} else {
						ifres11339 = False

					}

					var ifres11338 Obj

					if True == ifres11339 {
						ifres11338 = True

					} else {
						ifres11338 = False

					}

					ifres11337 = ifres11338

				} else {
					ifres11337 = False

				}

				if True == ifres11337 {
					tmp11330 := PrimTail(V3176)

					tmp11331 := PrimTail(tmp11330)

					tmp11332 := PrimHead(tmp11331)

					tmp11333 := PrimTail(V3176)

					tmp11334 := PrimHead(tmp11333)

					tmp11335 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), tmp11334)

					tmp11336 := PrimCons(tmp11335, Nil)

					__e.Return(PrimCons(tmp11332, tmp11336))
					return

				} else {
					tmp11329 := PrimIsPair(V3176)

					var ifres11316 Obj

					if True == tmp11329 {
						tmp11327 := PrimHead(V3176)

						tmp11328 := PrimEqual(symbar_b, tmp11327)

						var ifres11318 Obj

						if True == tmp11328 {
							tmp11325 := PrimTail(V3176)

							tmp11326 := PrimIsPair(tmp11325)

							var ifres11320 Obj

							if True == tmp11326 {
								tmp11322 := PrimTail(V3176)

								tmp11323 := PrimTail(tmp11322)

								tmp11324 := PrimEqual(Nil, tmp11323)

								var ifres11321 Obj

								if True == tmp11324 {
									ifres11321 = True

								} else {
									ifres11321 = False

								}

								ifres11320 = ifres11321

							} else {
								ifres11320 = False

							}

							var ifres11319 Obj

							if True == ifres11320 {
								ifres11319 = True

							} else {
								ifres11319 = False

							}

							ifres11318 = ifres11319

						} else {
							ifres11318 = False

						}

						var ifres11317 Obj

						if True == ifres11318 {
							ifres11317 = True

						} else {
							ifres11317 = False

						}

						ifres11316 = ifres11317

					} else {
						ifres11316 = False

					}

					if True == ifres11316 {
						tmp11315 := PrimTail(V3176)

						__e.Return(PrimHead(tmp11315))
						return

					} else {
						tmp11314 := PrimIsPair(V3176)

						if True == tmp11314 {
							tmp11308 := PrimHead(V3176)

							tmp11309 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), tmp11308)

							tmp11310 := PrimTail(V3176)

							tmp11311 := Call(__e, PrimNS2Value(symshen_4cons_1form_1with_1modes), tmp11310)

							tmp11312 := PrimCons(tmp11311, Nil)

							tmp11313 := PrimCons(tmp11309, tmp11312)

							__e.Return(PrimCons(symcons, tmp11313))
							return

						} else {
							__e.Return(V3176)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp11393 := Call(__e, PrimNS2Value(symdef), symshen_4cons_1form_1with_1modes, tmp11302)

	_ = tmp11393

	tmp11394 := MakeNative(func(__e *ControlFlow) {
		V3177 := __e.Get(1)
		_ = V3177
		V3178 := __e.Get(2)
		_ = V3178
		V3179 := __e.Get(3)
		_ = V3179
		V3180 := __e.Get(4)
		_ = V3180
		V3181 := __e.Get(5)
		_ = V3181
		V3182 := __e.Get(6)
		_ = V3182
		tmp11395 := MakeNative(func(__e *ControlFlow) {
			GoalsAs := __e.Get(1)
			_ = GoalsAs
			tmp11396 := MakeNative(func(__e *ControlFlow) {
				GoalsS := __e.Get(1)
				_ = GoalsS
				tmp11397 := MakeNative(func(__e *ControlFlow) {
					GoalsP := __e.Get(1)
					_ = GoalsP
					tmp11398 := Call(__e, PrimNS2Value(symappend), GoalsS, GoalsP)

					__e.TailApply(PrimNS2Value(symappend), GoalsAs, tmp11398)
					return

				}, 1)

				tmp11399 := Call(__e, PrimNS2Value(symshen_4compile_1premises), V3180, V3181)

				__e.TailApply(tmp11397, tmp11399)
				return

			}, 1)

			tmp11400 := Call(__e, PrimNS2Value(symshen_4compile_1side_1conditions), V3179)

			__e.TailApply(tmp11396, tmp11400)
			return

		}, 1)

		tmp11401 := Call(__e, PrimNS2Value(symshen_4compile_1assumptions), V3178, V3177, V3181, V3182)

		__e.TailApply(tmp11395, tmp11401)
		return

	}, 6)

	tmp11402 := Call(__e, PrimNS2Value(symdef), symshen_4goals, tmp11394)

	_ = tmp11402

	tmp11403 := MakeNative(func(__e *ControlFlow) {
		V3197 := __e.Get(1)
		_ = V3197
		V3198 := __e.Get(2)
		_ = V3198
		V3199 := __e.Get(3)
		_ = V3199
		V3200 := __e.Get(4)
		_ = V3200
		tmp11426 := PrimEqual(Nil, V3197)

		if True == tmp11426 {
			__e.Return(Nil)
			return
		} else {
			tmp11425 := PrimIsPair(V3197)

			var ifres11418 Obj

			if True == tmp11425 {
				tmp11424 := PrimIsPair(V3199)

				var ifres11420 Obj

				if True == tmp11424 {
					tmp11422 := PrimTail(V3199)

					tmp11423 := PrimIsPair(tmp11422)

					var ifres11421 Obj

					if True == tmp11423 {
						ifres11421 = True

					} else {
						ifres11421 = False

					}

					ifres11420 = ifres11421

				} else {
					ifres11420 = False

				}

				var ifres11419 Obj

				if True == ifres11420 {
					ifres11419 = True

				} else {
					ifres11419 = False

				}

				ifres11418 = ifres11419

			} else {
				ifres11418 = False

			}

			if True == ifres11418 {
				tmp11406 := MakeNative(func(__e *ControlFlow) {
					NewActive := __e.Get(1)
					_ = NewActive
					tmp11407 := PrimHead(V3197)

					tmp11408 := PrimHead(V3199)

					tmp11409 := PrimTail(V3199)

					tmp11410 := PrimHead(tmp11409)

					tmp11411 := Call(__e, PrimNS2Value(symshen_4compile_1assumption), tmp11407, tmp11408, tmp11410, V3198, V3200)

					tmp11412 := PrimTail(V3197)

					tmp11413 := PrimTail(V3199)

					tmp11414 := Call(__e, PrimNS2Value(symshen_4compile_1assumptions), tmp11412, V3198, tmp11413, NewActive)

					__e.Return(PrimCons(tmp11411, tmp11414))
					return

				}, 1)

				tmp11415 := PrimHead(V3197)

				tmp11416 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp11415)

				tmp11417 := Call(__e, PrimNS2Value(symappend), tmp11416, V3200)

				__e.TailApply(tmp11406, tmp11417)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-assumptions")))
				return
			}

		}

	}, 4)

	tmp11427 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1assumptions, tmp11403)

	_ = tmp11427

	tmp11428 := MakeNative(func(__e *ControlFlow) {
		V3201 := __e.Get(1)
		_ = V3201
		V3202 := __e.Get(2)
		_ = V3202
		V3203 := __e.Get(3)
		_ = V3203
		V3204 := __e.Get(4)
		_ = V3204
		V3205 := __e.Get(5)
		_ = V3205
		tmp11429 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			tmp11430 := MakeNative(func(__e *ControlFlow) {
				Compile := __e.Get(1)
				_ = Compile
				tmp11431 := PrimCons(V3203, V3204)

				tmp11432 := PrimCons(Nil, tmp11431)

				tmp11433 := PrimCons(V3202, tmp11432)

				__e.Return(PrimCons(F, tmp11433))
				return

			}, 1)

			tmp11434 := Call(__e, PrimNS2Value(symshen_4compile_1search_1procedure), F, V3201, V3202, V3203, V3204, V3205)

			__e.TailApply(tmp11430, tmp11434)
			return

		}, 1)

		tmp11435 := Call(__e, PrimNS2Value(symgensym), symshen_4search)

		__e.TailApply(tmp11429, tmp11435)
		return

	}, 5)

	tmp11436 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1assumption, tmp11428)

	_ = tmp11436

	tmp11437 := MakeNative(func(__e *ControlFlow) {
		V3206 := __e.Get(1)
		_ = V3206
		V3207 := __e.Get(2)
		_ = V3207
		V3208 := __e.Get(3)
		_ = V3208
		V3209 := __e.Get(4)
		_ = V3209
		V3210 := __e.Get(5)
		_ = V3210
		V3211 := __e.Get(6)
		_ = V3211
		tmp11438 := MakeNative(func(__e *ControlFlow) {
			Past := __e.Get(1)
			_ = Past
			tmp11439 := MakeNative(func(__e *ControlFlow) {
				Base := __e.Get(1)
				_ = Base
				tmp11440 := MakeNative(func(__e *ControlFlow) {
					Recursive := __e.Get(1)
					_ = Recursive
					tmp11441 := Call(__e, PrimNS2Value(symappend), Base, Recursive)

					tmp11442 := PrimCons(V3206, tmp11441)

					tmp11443 := PrimCons(symdefprolog, tmp11442)

					__e.TailApply(PrimNS2Value(symeval), tmp11443)
					return

				}, 1)

				tmp11444 := Call(__e, PrimNS2Value(symshen_4keep_1looking), V3206, V3208, Past, V3209, V3210)

				__e.TailApply(tmp11440, tmp11444)
				return

			}, 1)

			tmp11445 := Call(__e, PrimNS2Value(symshen_4foundit_b), V3207, V3208, Past, V3209, V3210, V3211)

			__e.TailApply(tmp11439, tmp11445)
			return

		}, 1)

		tmp11446 := Call(__e, PrimNS2Value(symprotect), symPrevious)

		tmp11447 := Call(__e, PrimNS2Value(symgensym), tmp11446)

		__e.TailApply(tmp11438, tmp11447)
		return

	}, 6)

	tmp11448 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1search_1procedure, tmp11437)

	_ = tmp11448

	tmp11449 := MakeNative(func(__e *ControlFlow) {
		V3212 := __e.Get(1)
		_ = V3212
		V3213 := __e.Get(2)
		_ = V3213
		V3214 := __e.Get(3)
		_ = V3214
		V3215 := __e.Get(4)
		_ = V3215
		V3216 := __e.Get(5)
		_ = V3216
		V3217 := __e.Get(6)
		_ = V3217
		tmp11450 := MakeNative(func(__e *ControlFlow) {
			Passive := __e.Get(1)
			_ = Passive
			tmp11451 := MakeNative(func(__e *ControlFlow) {
				Table := __e.Get(1)
				_ = Table
				tmp11452 := MakeNative(func(__e *ControlFlow) {
					Head := __e.Get(1)
					_ = Head
					tmp11453 := MakeNative(func(__e *ControlFlow) {
						Body := __e.Get(1)
						_ = Body
						tmp11454 := PrimCons(sym_5_1_1, Nil)

						tmp11455 := PrimIntern(MakeString(";"))

						tmp11456 := PrimCons(tmp11455, Nil)

						tmp11457 := Call(__e, PrimNS2Value(symappend), Body, tmp11456)

						tmp11458 := Call(__e, PrimNS2Value(symappend), tmp11454, tmp11457)

						__e.TailApply(PrimNS2Value(symappend), Head, tmp11458)
						return

					}, 1)

					tmp11459 := Call(__e, PrimNS2Value(symshen_4body_1foundit_b), V3213, V3214, V3215, Table)

					__e.TailApply(tmp11453, tmp11459)
					return

				}, 1)

				tmp11460 := Call(__e, PrimNS2Value(symshen_4head_1foundit_b), V3212, V3213, V3214, V3215, V3216, Table)

				__e.TailApply(tmp11452, tmp11460)
				return

			}, 1)

			tmp11461 := Call(__e, PrimNS2Value(symshen_4tabulate_1passive), Passive)

			__e.TailApply(tmp11451, tmp11461)
			return

		}, 1)

		tmp11462 := Call(__e, PrimNS2Value(symshen_4passive), V3212, V3217)

		__e.TailApply(tmp11450, tmp11462)
		return

	}, 6)

	tmp11463 := Call(__e, PrimNS2Value(symdef), symshen_4foundit_b, tmp11449)

	_ = tmp11463

	tmp11464 := MakeNative(func(__e *ControlFlow) {
		V3218 := __e.Get(1)
		_ = V3218
		V3219 := __e.Get(2)
		_ = V3219
		V3220 := __e.Get(3)
		_ = V3220
		V3221 := __e.Get(4)
		_ = V3221
		V3222 := __e.Get(5)
		_ = V3222
		tmp11465 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp11466 := MakeNative(func(__e *ControlFlow) {
				Head := __e.Get(1)
				_ = Head
				tmp11467 := MakeNative(func(__e *ControlFlow) {
					Body := __e.Get(1)
					_ = Body
					tmp11468 := PrimCons(sym_5_1_1, Nil)

					tmp11469 := PrimIntern(MakeString(";"))

					tmp11470 := PrimCons(tmp11469, Nil)

					tmp11471 := Call(__e, PrimNS2Value(symappend), Body, tmp11470)

					tmp11472 := Call(__e, PrimNS2Value(symappend), tmp11468, tmp11471)

					__e.TailApply(PrimNS2Value(symappend), Head, tmp11472)
					return

				}, 1)

				tmp11473 := PrimCons(V3220, Nil)

				tmp11474 := PrimCons(X, tmp11473)

				tmp11475 := PrimCons(symcons, tmp11474)

				tmp11476 := PrimCons(V3221, V3222)

				tmp11477 := PrimCons(tmp11475, tmp11476)

				tmp11478 := PrimCons(V3219, tmp11477)

				tmp11479 := PrimCons(V3218, tmp11478)

				tmp11480 := PrimCons(tmp11479, Nil)

				__e.TailApply(tmp11467, tmp11480)
				return

			}, 1)

			tmp11481 := PrimCons(V3219, Nil)

			tmp11482 := PrimCons(X, tmp11481)

			tmp11483 := PrimCons(symcons, tmp11482)

			tmp11484 := PrimCons(tmp11483, Nil)

			tmp11485 := PrimCons(sym_1, tmp11484)

			tmp11486 := PrimCons(V3221, V3222)

			tmp11487 := PrimCons(V3220, tmp11486)

			tmp11488 := PrimCons(tmp11485, tmp11487)

			__e.TailApply(tmp11466, tmp11488)
			return

		}, 1)

		tmp11489 := Call(__e, PrimNS2Value(symprotect), symV)

		tmp11490 := Call(__e, PrimNS2Value(symgensym), tmp11489)

		__e.TailApply(tmp11465, tmp11490)
		return

	}, 5)

	tmp11491 := Call(__e, PrimNS2Value(symdef), symshen_4keep_1looking, tmp11464)

	_ = tmp11491

	tmp11492 := MakeNative(func(__e *ControlFlow) {
		V3227 := __e.Get(1)
		_ = V3227
		V3228 := __e.Get(2)
		_ = V3228
		tmp11500 := PrimIsPair(V3227)

		if True == tmp11500 {
			tmp11496 := PrimHead(V3227)

			tmp11497 := Call(__e, PrimNS2Value(symshen_4passive), tmp11496, V3228)

			tmp11498 := PrimTail(V3227)

			tmp11499 := Call(__e, PrimNS2Value(symshen_4passive), tmp11498, V3228)

			__e.TailApply(PrimNS2Value(symunion), tmp11497, tmp11499)
			return

		} else {
			tmp11495 := Call(__e, PrimNS2Value(symshen_4passive_2), V3227, V3228)

			if True == tmp11495 {
				__e.Return(PrimCons(V3227, Nil))
				return
			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 2)

	tmp11501 := Call(__e, PrimNS2Value(symdef), symshen_4passive, tmp11492)

	_ = tmp11501

	tmp11502 := MakeNative(func(__e *ControlFlow) {
		V3229 := __e.Get(1)
		_ = V3229
		V3230 := __e.Get(2)
		_ = V3230
		tmp11506 := Call(__e, PrimNS2Value(symelement_2), V3229, V3230)

		tmp11507 := PrimNot(tmp11506)

		if True == tmp11507 {
			tmp11505 := PrimIsVariable(V3229)

			if True == tmp11505 {
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

	}, 2)

	tmp11508 := Call(__e, PrimNS2Value(symdef), symshen_4passive_2, tmp11502)

	_ = tmp11508

	tmp11509 := MakeNative(func(__e *ControlFlow) {
		V3231 := __e.Get(1)
		_ = V3231
		tmp11510 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp11511 := Call(__e, PrimNS2Value(symprotect), symV)

			tmp11512 := Call(__e, PrimNS2Value(symgensym), tmp11511)

			__e.Return(PrimCons(X, tmp11512))
			return

		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp11510, V3231)
		return

	}, 1)

	tmp11513 := Call(__e, PrimNS2Value(symdef), symshen_4tabulate_1passive, tmp11509)

	_ = tmp11513

	tmp11514 := MakeNative(func(__e *ControlFlow) {
		V3232 := __e.Get(1)
		_ = V3232
		V3233 := __e.Get(2)
		_ = V3233
		V3234 := __e.Get(3)
		_ = V3234
		V3235 := __e.Get(4)
		_ = V3235
		V3236 := __e.Get(5)
		_ = V3236
		V3237 := __e.Get(6)
		_ = V3237
		tmp11515 := MakeNative(func(__e *ControlFlow) {
			Optimise := __e.Get(1)
			_ = Optimise
			tmp11516 := Call(__e, PrimNS2Value(symshen_4optimise_1typing), V3232)

			tmp11517 := PrimCons(V3233, Nil)

			tmp11518 := PrimCons(tmp11516, tmp11517)

			tmp11519 := PrimCons(symcons, tmp11518)

			tmp11520 := PrimCons(tmp11519, Nil)

			tmp11521 := PrimCons(sym_1, tmp11520)

			tmp11522 := PrimCons(V3235, Optimise)

			tmp11523 := PrimCons(V3234, tmp11522)

			__e.Return(PrimCons(tmp11521, tmp11523))
			return

		}, 1)

		tmp11524 := Call(__e, PrimNS2Value(symshen_4optimise_1passive), V3236, V3237)

		__e.TailApply(tmp11515, tmp11524)
		return

	}, 6)

	tmp11525 := Call(__e, PrimNS2Value(symdef), symshen_4head_1foundit_b, tmp11514)

	_ = tmp11525

	tmp11526 := MakeNative(func(__e *ControlFlow) {
		V3238 := __e.Get(1)
		_ = V3238
		V3239 := __e.Get(2)
		_ = V3239
		tmp11527 := MakeNative(func(__e *ControlFlow) {
			C := __e.Get(1)
			_ = C
			__e.TailApply(PrimNS2Value(symshen_4optimise_1passive_1h), C, V3239)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp11527, V3238)
		return

	}, 2)

	tmp11528 := Call(__e, PrimNS2Value(symdef), symshen_4optimise_1passive, tmp11526)

	_ = tmp11528

	tmp11529 := MakeNative(func(__e *ControlFlow) {
		V3240 := __e.Get(1)
		_ = V3240
		V3241 := __e.Get(2)
		_ = V3241
		tmp11530 := MakeNative(func(__e *ControlFlow) {
			Entry := __e.Get(1)
			_ = Entry
			tmp11532 := Call(__e, PrimNS2Value(symempty_2), Entry)

			if True == tmp11532 {
				__e.Return(V3240)
				return
			} else {
				__e.Return(PrimTail(Entry))
				return
			}

		}, 1)

		tmp11533 := Call(__e, PrimNS2Value(symassoc), V3240, V3241)

		__e.TailApply(tmp11530, tmp11533)
		return

	}, 2)

	tmp11534 := Call(__e, PrimNS2Value(symdef), symshen_4optimise_1passive_1h, tmp11529)

	_ = tmp11534

	tmp11535 := MakeNative(func(__e *ControlFlow) {
		V3250 := __e.Get(1)
		_ = V3250
		V3251 := __e.Get(2)
		_ = V3251
		V3252 := __e.Get(3)
		_ = V3252
		V3253 := __e.Get(4)
		_ = V3253
		tmp11562 := PrimEqual(Nil, V3253)

		if True == tmp11562 {
			tmp11552 := PrimCons(V3251, Nil)

			tmp11553 := PrimCons(MakeNumber(1), tmp11552)

			tmp11554 := PrimCons(V3250, Nil)

			tmp11555 := PrimCons(MakeNumber(1), tmp11554)

			tmp11556 := PrimCons(tmp11555, Nil)

			tmp11557 := PrimCons(tmp11553, tmp11556)

			tmp11558 := PrimCons(symappend, tmp11557)

			tmp11559 := PrimCons(tmp11558, Nil)

			tmp11560 := PrimCons(V3252, tmp11559)

			tmp11561 := PrimCons(symbind, tmp11560)

			__e.Return(PrimCons(tmp11561, Nil))
			return

		} else {
			tmp11551 := PrimIsPair(V3253)

			var ifres11547 Obj

			if True == tmp11551 {
				tmp11549 := PrimHead(V3253)

				tmp11550 := PrimIsPair(tmp11549)

				var ifres11548 Obj

				if True == tmp11550 {
					ifres11548 = True

				} else {
					ifres11548 = False

				}

				ifres11547 = ifres11548

			} else {
				ifres11547 = False

			}

			if True == ifres11547 {
				tmp11538 := PrimHead(V3253)

				tmp11539 := PrimTail(tmp11538)

				tmp11540 := PrimHead(V3253)

				tmp11541 := PrimHead(tmp11540)

				tmp11542 := PrimCons(tmp11541, Nil)

				tmp11543 := PrimCons(tmp11539, tmp11542)

				tmp11544 := PrimCons(symbind, tmp11543)

				tmp11545 := PrimTail(V3253)

				tmp11546 := Call(__e, PrimNS2Value(symshen_4body_1foundit_b), V3250, V3251, V3252, tmp11545)

				__e.Return(PrimCons(tmp11544, tmp11546))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.body-foundit!")))
				return
			}

		}

	}, 4)

	tmp11563 := Call(__e, PrimNS2Value(symdef), symshen_4body_1foundit_b, tmp11535)

	_ = tmp11563

	tmp11564 := MakeNative(func(__e *ControlFlow) {
		V3254 := __e.Get(1)
		_ = V3254
		tmp11565 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4compile_1side_1condition), X)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp11565, V3254)
		return

	}, 1)

	tmp11566 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1side_1conditions, tmp11564)

	_ = tmp11566

	tmp11567 := MakeNative(func(__e *ControlFlow) {
		V3257 := __e.Get(1)
		_ = V3257
		tmp11627 := PrimIsPair(V3257)

		var ifres11608 Obj

		if True == tmp11627 {
			tmp11625 := PrimHead(V3257)

			tmp11626 := PrimEqual(symlet, tmp11625)

			var ifres11610 Obj

			if True == tmp11626 {
				tmp11623 := PrimTail(V3257)

				tmp11624 := PrimIsPair(tmp11623)

				var ifres11612 Obj

				if True == tmp11624 {
					tmp11620 := PrimTail(V3257)

					tmp11621 := PrimTail(tmp11620)

					tmp11622 := PrimIsPair(tmp11621)

					var ifres11614 Obj

					if True == tmp11622 {
						tmp11616 := PrimTail(V3257)

						tmp11617 := PrimTail(tmp11616)

						tmp11618 := PrimTail(tmp11617)

						tmp11619 := PrimEqual(Nil, tmp11618)

						var ifres11615 Obj

						if True == tmp11619 {
							ifres11615 = True

						} else {
							ifres11615 = False

						}

						ifres11614 = ifres11615

					} else {
						ifres11614 = False

					}

					var ifres11613 Obj

					if True == ifres11614 {
						ifres11613 = True

					} else {
						ifres11613 = False

					}

					ifres11612 = ifres11613

				} else {
					ifres11612 = False

				}

				var ifres11611 Obj

				if True == ifres11612 {
					ifres11611 = True

				} else {
					ifres11611 = False

				}

				ifres11610 = ifres11611

			} else {
				ifres11610 = False

			}

			var ifres11609 Obj

			if True == ifres11610 {
				ifres11609 = True

			} else {
				ifres11609 = False

			}

			ifres11608 = ifres11609

		} else {
			ifres11608 = False

		}

		if True == ifres11608 {
			tmp11607 := PrimTail(V3257)

			__e.Return(PrimCons(symis, tmp11607))
			return

		} else {
			tmp11606 := PrimIsPair(V3257)

			var ifres11587 Obj

			if True == tmp11606 {
				tmp11604 := PrimHead(V3257)

				tmp11605 := PrimEqual(symshen_4let_b, tmp11604)

				var ifres11589 Obj

				if True == tmp11605 {
					tmp11602 := PrimTail(V3257)

					tmp11603 := PrimIsPair(tmp11602)

					var ifres11591 Obj

					if True == tmp11603 {
						tmp11599 := PrimTail(V3257)

						tmp11600 := PrimTail(tmp11599)

						tmp11601 := PrimIsPair(tmp11600)

						var ifres11593 Obj

						if True == tmp11601 {
							tmp11595 := PrimTail(V3257)

							tmp11596 := PrimTail(tmp11595)

							tmp11597 := PrimTail(tmp11596)

							tmp11598 := PrimEqual(Nil, tmp11597)

							var ifres11594 Obj

							if True == tmp11598 {
								ifres11594 = True

							} else {
								ifres11594 = False

							}

							ifres11593 = ifres11594

						} else {
							ifres11593 = False

						}

						var ifres11592 Obj

						if True == ifres11593 {
							ifres11592 = True

						} else {
							ifres11592 = False

						}

						ifres11591 = ifres11592

					} else {
						ifres11591 = False

					}

					var ifres11590 Obj

					if True == ifres11591 {
						ifres11590 = True

					} else {
						ifres11590 = False

					}

					ifres11589 = ifres11590

				} else {
					ifres11589 = False

				}

				var ifres11588 Obj

				if True == ifres11589 {
					ifres11588 = True

				} else {
					ifres11588 = False

				}

				ifres11587 = ifres11588

			} else {
				ifres11587 = False

			}

			if True == ifres11587 {
				tmp11586 := PrimTail(V3257)

				__e.Return(PrimCons(symis_b, tmp11586))
				return

			} else {
				tmp11585 := PrimIsPair(V3257)

				var ifres11572 Obj

				if True == tmp11585 {
					tmp11583 := PrimHead(V3257)

					tmp11584 := PrimEqual(symif, tmp11583)

					var ifres11574 Obj

					if True == tmp11584 {
						tmp11581 := PrimTail(V3257)

						tmp11582 := PrimIsPair(tmp11581)

						var ifres11576 Obj

						if True == tmp11582 {
							tmp11578 := PrimTail(V3257)

							tmp11579 := PrimTail(tmp11578)

							tmp11580 := PrimEqual(Nil, tmp11579)

							var ifres11577 Obj

							if True == tmp11580 {
								ifres11577 = True

							} else {
								ifres11577 = False

							}

							ifres11576 = ifres11577

						} else {
							ifres11576 = False

						}

						var ifres11575 Obj

						if True == ifres11576 {
							ifres11575 = True

						} else {
							ifres11575 = False

						}

						ifres11574 = ifres11575

					} else {
						ifres11574 = False

					}

					var ifres11573 Obj

					if True == ifres11574 {
						ifres11573 = True

					} else {
						ifres11573 = False

					}

					ifres11572 = ifres11573

				} else {
					ifres11572 = False

				}

				if True == ifres11572 {
					tmp11571 := PrimTail(V3257)

					__e.Return(PrimCons(symwhen, tmp11571))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-side-condition")))
					return
				}

			}

		}

	}, 1)

	tmp11628 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1side_1condition, tmp11567)

	_ = tmp11628

	tmp11629 := MakeNative(func(__e *ControlFlow) {
		V3258 := __e.Get(1)
		_ = V3258
		V3259 := __e.Get(2)
		_ = V3259
		tmp11630 := MakeNative(func(__e *ControlFlow) {
			Hyp := __e.Get(1)
			_ = Hyp
			tmp11631 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4compile_1premise), X, Hyp)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp11631, V3258)
			return

		}, 1)

		tmp11632 := Call(__e, PrimNS2Value(symreverse), V3259)

		tmp11633 := PrimHead(tmp11632)

		__e.TailApply(tmp11630, tmp11633)
		return

	}, 2)

	tmp11634 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1premises, tmp11629)

	_ = tmp11634

	tmp11635 := MakeNative(func(__e *ControlFlow) {
		V3266 := __e.Get(1)
		_ = V3266
		V3267 := __e.Get(2)
		_ = V3267
		tmp11652 := PrimEqual(sym_b, V3266)

		if True == tmp11652 {
			__e.Return(sym_b)
			return
		} else {
			tmp11651 := PrimIsPair(V3266)

			var ifres11642 Obj

			if True == tmp11651 {
				tmp11649 := PrimTail(V3266)

				tmp11650 := PrimIsPair(tmp11649)

				var ifres11644 Obj

				if True == tmp11650 {
					tmp11646 := PrimTail(V3266)

					tmp11647 := PrimTail(tmp11646)

					tmp11648 := PrimEqual(Nil, tmp11647)

					var ifres11645 Obj

					if True == tmp11648 {
						ifres11645 = True

					} else {
						ifres11645 = False

					}

					ifres11644 = ifres11645

				} else {
					ifres11644 = False

				}

				var ifres11643 Obj

				if True == ifres11644 {
					ifres11643 = True

				} else {
					ifres11643 = False

				}

				ifres11642 = ifres11643

			} else {
				ifres11642 = False

			}

			if True == ifres11642 {
				tmp11638 := PrimHead(V3266)

				tmp11639 := Call(__e, PrimNS2Value(symreverse), tmp11638)

				tmp11640 := PrimTail(V3266)

				tmp11641 := PrimHead(tmp11640)

				__e.TailApply(PrimNS2Value(symshen_4compile_1premise_1h), tmp11639, tmp11641, V3267)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.premise")))
				return
			}

		}

	}, 2)

	tmp11653 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1premise, tmp11635)

	_ = tmp11653

	tmp11654 := MakeNative(func(__e *ControlFlow) {
		V3274 := __e.Get(1)
		_ = V3274
		V3275 := __e.Get(2)
		_ = V3275
		V3276 := __e.Get(3)
		_ = V3276
		tmp11667 := PrimEqual(Nil, V3274)

		if True == tmp11667 {
			tmp11664 := Call(__e, PrimNS2Value(symshen_4cons_1form_1no_1modes), V3275)

			tmp11665 := PrimCons(V3276, Nil)

			tmp11666 := PrimCons(tmp11664, tmp11665)

			__e.Return(PrimCons(symshen_4system_1S, tmp11666))
			return

		} else {
			tmp11663 := PrimIsPair(V3274)

			if True == tmp11663 {
				tmp11657 := PrimTail(V3274)

				tmp11658 := PrimHead(V3274)

				tmp11659 := Call(__e, PrimNS2Value(symshen_4cons_1form_1no_1modes), tmp11658)

				tmp11660 := PrimCons(V3276, Nil)

				tmp11661 := PrimCons(tmp11659, tmp11660)

				tmp11662 := PrimCons(symcons, tmp11661)

				__e.TailApply(PrimNS2Value(symshen_4compile_1premise_1h), tmp11657, V3275, tmp11662)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-premise-h")))
				return
			}

		}

	}, 3)

	tmp11668 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1premise_1h, tmp11654)

	_ = tmp11668

	tmp11669 := MakeNative(func(__e *ControlFlow) {
		V3277 := __e.Get(1)
		_ = V3277
		tmp11693 := PrimIsPair(V3277)

		var ifres11680 Obj

		if True == tmp11693 {
			tmp11691 := PrimHead(V3277)

			tmp11692 := PrimEqual(symbar_b, tmp11691)

			var ifres11682 Obj

			if True == tmp11692 {
				tmp11689 := PrimTail(V3277)

				tmp11690 := PrimIsPair(tmp11689)

				var ifres11684 Obj

				if True == tmp11690 {
					tmp11686 := PrimTail(V3277)

					tmp11687 := PrimTail(tmp11686)

					tmp11688 := PrimEqual(Nil, tmp11687)

					var ifres11685 Obj

					if True == tmp11688 {
						ifres11685 = True

					} else {
						ifres11685 = False

					}

					ifres11684 = ifres11685

				} else {
					ifres11684 = False

				}

				var ifres11683 Obj

				if True == ifres11684 {
					ifres11683 = True

				} else {
					ifres11683 = False

				}

				ifres11682 = ifres11683

			} else {
				ifres11682 = False

			}

			var ifres11681 Obj

			if True == ifres11682 {
				ifres11681 = True

			} else {
				ifres11681 = False

			}

			ifres11680 = ifres11681

		} else {
			ifres11680 = False

		}

		if True == ifres11680 {
			tmp11679 := PrimTail(V3277)

			__e.Return(PrimHead(tmp11679))
			return

		} else {
			tmp11678 := PrimIsPair(V3277)

			if True == tmp11678 {
				tmp11672 := PrimHead(V3277)

				tmp11673 := Call(__e, PrimNS2Value(symshen_4cons_1form_1no_1modes), tmp11672)

				tmp11674 := PrimTail(V3277)

				tmp11675 := Call(__e, PrimNS2Value(symshen_4cons_1form_1no_1modes), tmp11674)

				tmp11676 := PrimCons(tmp11675, Nil)

				tmp11677 := PrimCons(tmp11673, tmp11676)

				__e.Return(PrimCons(symcons, tmp11677))
				return

			} else {
				__e.Return(V3277)
				return
			}

		}

	}, 1)

	tmp11694 := Call(__e, PrimNS2Value(symdef), symshen_4cons_1form_1no_1modes, tmp11669)

	_ = tmp11694

	tmp11695 := MakeNative(func(__e *ControlFlow) {
		V3278 := __e.Get(1)
		_ = V3278
		tmp11696 := MakeNative(func(__e *ControlFlow) {
			InternTypes := __e.Get(1)
			_ = InternTypes
			tmp11697 := MakeNative(func(__e *ControlFlow) {
				Datatypes := __e.Get(1)
				_ = Datatypes
				tmp11698 := MakeNative(func(__e *ControlFlow) {
					Remove := __e.Get(1)
					_ = Remove
					tmp11699 := MakeNative(func(__e *ControlFlow) {
						NewDatatypes := __e.Get(1)
						_ = NewDatatypes
						__e.TailApply(PrimNS2Value(symshen_4show_1datatypes), NewDatatypes)
						return
					}, 1)

					tmp11700 := PrimNS3Set(symshen_4_ddatatypes_d, Remove)

					__e.TailApply(tmp11699, tmp11700)
					return

				}, 1)

				tmp11701 := Call(__e, PrimNS2Value(symshen_4remove_1datatypes), InternTypes, Datatypes)

				__e.TailApply(tmp11698, tmp11701)
				return

			}, 1)

			tmp11702 := PrimNS3Value(symshen_4_ddatatypes_d)

			__e.TailApply(tmp11697, tmp11702)
			return

		}, 1)

		tmp11703 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp11704 := Call(__e, PrimNS2Value(symmap), tmp11703, V3278)

		__e.TailApply(tmp11696, tmp11704)
		return

	}, 1)

	tmp11705 := Call(__e, PrimNS2Value(symdef), sympreclude, tmp11695)

	_ = tmp11705

	tmp11706 := MakeNative(func(__e *ControlFlow) {
		V3283 := __e.Get(1)
		_ = V3283
		V3284 := __e.Get(2)
		_ = V3284
		tmp11713 := PrimEqual(Nil, V3283)

		if True == tmp11713 {
			__e.Return(V3284)
			return
		} else {
			tmp11712 := PrimIsPair(V3283)

			if True == tmp11712 {
				tmp11709 := PrimTail(V3283)

				tmp11710 := PrimHead(V3283)

				tmp11711 := Call(__e, PrimNS2Value(symshen_4unassoc), tmp11710, V3284)

				__e.TailApply(PrimNS2Value(symshen_4remove_1datatypes), tmp11709, tmp11711)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
				return
			}

		}

	}, 2)

	tmp11714 := Call(__e, PrimNS2Value(symdef), symshen_4remove_1datatypes, tmp11706)

	_ = tmp11714

	tmp11715 := MakeNative(func(__e *ControlFlow) {
		V3294 := __e.Get(1)
		_ = V3294
		V3295 := __e.Get(2)
		_ = V3295
		tmp11733 := PrimEqual(Nil, V3295)

		if True == tmp11733 {
			__e.Return(Nil)
			return
		} else {
			tmp11732 := PrimIsPair(V3295)

			var ifres11723 Obj

			if True == tmp11732 {
				tmp11730 := PrimHead(V3295)

				tmp11731 := PrimIsPair(tmp11730)

				var ifres11725 Obj

				if True == tmp11731 {
					tmp11727 := PrimHead(V3295)

					tmp11728 := PrimHead(tmp11727)

					tmp11729 := PrimEqual(V3294, tmp11728)

					var ifres11726 Obj

					if True == tmp11729 {
						ifres11726 = True

					} else {
						ifres11726 = False

					}

					ifres11725 = ifres11726

				} else {
					ifres11725 = False

				}

				var ifres11724 Obj

				if True == ifres11725 {
					ifres11724 = True

				} else {
					ifres11724 = False

				}

				ifres11723 = ifres11724

			} else {
				ifres11723 = False

			}

			if True == ifres11723 {
				__e.Return(PrimTail(V3295))
				return
			} else {
				tmp11722 := PrimIsPair(V3295)

				if True == tmp11722 {
					tmp11719 := PrimHead(V3295)

					tmp11720 := PrimTail(V3295)

					tmp11721 := Call(__e, PrimNS2Value(symshen_4unassoc), V3294, tmp11720)

					__e.Return(PrimCons(tmp11719, tmp11721))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.unassoc")))
					return
				}

			}

		}

	}, 2)

	tmp11734 := Call(__e, PrimNS2Value(symdef), symshen_4unassoc, tmp11715)

	_ = tmp11734

	tmp11735 := MakeNative(func(__e *ControlFlow) {
		V3296 := __e.Get(1)
		_ = V3296
		tmp11736 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(PrimHead(X))
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp11736, V3296)
		return

	}, 1)

	tmp11737 := Call(__e, PrimNS2Value(symdef), symshen_4show_1datatypes, tmp11735)

	_ = tmp11737

	tmp11738 := MakeNative(func(__e *ControlFlow) {
		V3297 := __e.Get(1)
		_ = V3297
		tmp11739 := MakeNative(func(__e *ControlFlow) {
			InternTypes := __e.Get(1)
			_ = InternTypes
			tmp11740 := MakeNative(func(__e *ControlFlow) {
				Remember := __e.Get(1)
				_ = Remember
				tmp11741 := MakeNative(func(__e *ControlFlow) {
					Datatypes := __e.Get(1)
					_ = Datatypes
					__e.TailApply(PrimNS2Value(symshen_4show_1datatypes), Datatypes)
					return
				}, 1)

				tmp11742 := PrimNS3Value(symshen_4_ddatatypes_d)

				__e.TailApply(tmp11741, tmp11742)
				return

			}, 1)

			tmp11743 := MakeNative(func(__e *ControlFlow) {
				D := __e.Get(1)
				_ = D
				tmp11744 := Call(__e, PrimNS2Value(symfn), D)

				__e.TailApply(PrimNS2Value(symshen_4remember_1datatype), D, tmp11744)
				return

			}, 1)

			tmp11745 := Call(__e, PrimNS2Value(symmap), tmp11743, InternTypes)

			__e.TailApply(tmp11740, tmp11745)
			return

		}, 1)

		tmp11746 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp11747 := Call(__e, PrimNS2Value(symmap), tmp11746, V3297)

		__e.TailApply(tmp11739, tmp11747)
		return

	}, 1)

	tmp11748 := Call(__e, PrimNS2Value(symdef), syminclude, tmp11738)

	_ = tmp11748

	tmp11749 := MakeNative(func(__e *ControlFlow) {
		V3298 := __e.Get(1)
		_ = V3298
		tmp11750 := MakeNative(func(__e *ControlFlow) {
			Initialise := __e.Get(1)
			_ = Initialise
			tmp11751 := MakeNative(func(__e *ControlFlow) {
				InternTypes := __e.Get(1)
				_ = InternTypes
				tmp11752 := MakeNative(func(__e *ControlFlow) {
					NewDatatypes := __e.Get(1)
					_ = NewDatatypes
					tmp11753 := PrimNS3Value(symshen_4_ddatatypes_d)

					__e.TailApply(PrimNS2Value(symshen_4show_1datatypes), tmp11753)
					return

				}, 1)

				tmp11754 := MakeNative(func(__e *ControlFlow) {
					D := __e.Get(1)
					_ = D
					tmp11755 := Call(__e, PrimNS2Value(symfn), D)

					__e.TailApply(PrimNS2Value(symshen_4remember_1datatype), D, tmp11755)
					return

				}, 1)

				tmp11756 := Call(__e, PrimNS2Value(symmap), tmp11754, InternTypes)

				__e.TailApply(tmp11752, tmp11756)
				return

			}, 1)

			tmp11757 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
				return
			}, 1)

			tmp11758 := Call(__e, PrimNS2Value(symmap), tmp11757, V3298)

			__e.TailApply(tmp11751, tmp11758)
			return

		}, 1)

		tmp11759 := PrimNS3Set(symshen_4_ddatatypes_d, Nil)

		__e.TailApply(tmp11750, tmp11759)
		return

	}, 1)

	tmp11760 := Call(__e, PrimNS2Value(symdef), sympreclude_1all_1but, tmp11749)

	_ = tmp11760

	tmp11761 := MakeNative(func(__e *ControlFlow) {
		V3299 := __e.Get(1)
		_ = V3299
		tmp11762 := MakeNative(func(__e *ControlFlow) {
			InternTypes := __e.Get(1)
			_ = InternTypes
			tmp11763 := MakeNative(func(__e *ControlFlow) {
				AllDatatypes := __e.Get(1)
				_ = AllDatatypes
				tmp11764 := MakeNative(func(__e *ControlFlow) {
					Datatypes := __e.Get(1)
					_ = Datatypes
					__e.TailApply(PrimNS2Value(symshen_4show_1datatypes), Datatypes)
					return
				}, 1)

				tmp11765 := Call(__e, PrimNS2Value(symshen_4remove_1datatypes), InternTypes, AllDatatypes)

				tmp11766 := PrimNS3Set(symshen_4_ddatatypes_d, tmp11765)

				__e.TailApply(tmp11764, tmp11766)
				return

			}, 1)

			tmp11767 := PrimNS3Value(symshen_4_dalldatatypes_d)

			__e.TailApply(tmp11763, tmp11767)
			return

		}, 1)

		tmp11768 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp11769 := Call(__e, PrimNS2Value(symmap), tmp11768, V3299)

		__e.TailApply(tmp11762, tmp11769)
		return

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), syminclude_1all_1but, tmp11761)
	return

}, 0)
