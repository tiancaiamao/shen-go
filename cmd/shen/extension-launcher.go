package main

import . "github.com/tiancaiamao/shen-go/cora"

var ExtensionLauncherMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

	tmp4608 := MakeNative(func(__e *ControlFlow) {
		V4891 := __e.Get(1)
		_ = V4891
		tmp4609 := MakeNative(func(__e *ControlFlow) {
			Contents := __e.Get(1)
			_ = Contents
			tmp4610 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4eval_1without_1macros), X)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp4610, Contents)
			return

		}, 1)

		tmp4611 := Call(__e, PrimNS2Value(symread_1file), V4891)

		__e.TailApply(tmp4609, tmp4611)
		return

	}, 1)

	tmp4612 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4quiet_1load, tmp4608)

	_ = tmp4612

	tmp4613 := MakeNative(func(__e *ControlFlow) {
		tmp4614 := Call(__e, PrimNS2Value(symversion))

		tmp4615 := Call(__e, PrimNS2Value(symlanguage))

		tmp4616 := Call(__e, PrimNS2Value(symport))

		tmp4617 := PrimCons(tmp4616, Nil)

		tmp4618 := PrimCons(tmp4615, tmp4617)

		tmp4619 := Call(__e, PrimNS2Value(symimplementation))

		tmp4620 := Call(__e, PrimNS2Value(symrelease))

		tmp4621 := PrimCons(tmp4620, Nil)

		tmp4622 := PrimCons(tmp4619, tmp4621)

		tmp4623 := PrimCons(tmp4622, Nil)

		tmp4624 := PrimCons(symimplementation, tmp4623)

		tmp4625 := PrimCons(tmp4618, tmp4624)

		tmp4626 := PrimCons(symport, tmp4625)

		tmp4627 := Call(__e, PrimNS2Value(symshen_4app), tmp4626, MakeString("\n"), symshen_4r)

		tmp4628 := PrimStringConcat(MakeString(" "), tmp4627)

		__e.TailApply(PrimNS2Value(symshen_4app), tmp4614, tmp4628, symshen_4a)
		return

	}, 0)

	tmp4629 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4version_1string, tmp4613)

	_ = tmp4629

	tmp4630 := MakeNative(func(__e *ControlFlow) {
		V4893 := __e.Get(1)
		_ = V4893
		tmp4631 := Call(__e, PrimNS2Value(symshen_4app), V4893, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)

		__e.Return(PrimStringConcat(MakeString("Usage: "), tmp4631))
		return

	}, 1)

	tmp4632 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4help_1text, tmp4630)

	_ = tmp4632

	tmp4633 := MakeNative(func(__e *ControlFlow) {
		V4895 := __e.Get(1)
		_ = V4895
		tmp4640 := PrimEqual(Nil, V4895)

		if True == tmp4640 {
			__e.Return(PrimCons(symsuccess, Nil))
			return
		} else {
			tmp4639 := PrimIsPair(V4895)

			if True == tmp4639 {
				tmp4636 := PrimHead(V4895)

				tmp4637 := Call(__e, PrimNS2Value(symthaw), tmp4636)

				_ = tmp4637

				tmp4638 := PrimTail(V4895)

				__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4execute_1all), tmp4638)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4launcher_4execute_1all)
				return
			}

		}

	}, 1)

	tmp4641 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4execute_1all, tmp4633)

	_ = tmp4641

	tmp4642 := MakeNative(func(__e *ControlFlow) {
		V4897 := __e.Get(1)
		_ = V4897
		tmp4643 := Call(__e, PrimNS2Value(symread_1from_1string), V4897)

		tmp4644 := Call(__e, PrimNS2Value(symhead), tmp4643)

		__e.TailApply(PrimNS2Value(symeval), tmp4644)
		return

	}, 1)

	tmp4645 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1string, tmp4642)

	_ = tmp4645

	tmp4646 := MakeNative(func(__e *ControlFlow) {
		V4903 := __e.Get(1)
		_ = V4903
		tmp4656 := PrimEqual(MakeString("-e"), V4903)

		if True == tmp4656 {
			__e.Return(MakeString("--eval"))
			return
		} else {
			tmp4655 := PrimEqual(MakeString("-l"), V4903)

			if True == tmp4655 {
				__e.Return(MakeString("--load"))
				return
			} else {
				tmp4654 := PrimEqual(MakeString("-q"), V4903)

				if True == tmp4654 {
					__e.Return(MakeString("--quiet"))
					return
				} else {
					tmp4653 := PrimEqual(MakeString("-s"), V4903)

					if True == tmp4653 {
						__e.Return(MakeString("--set"))
						return
					} else {
						tmp4652 := PrimEqual(MakeString("-r"), V4903)

						if True == tmp4652 {
							__e.Return(MakeString("--repl"))
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

	tmp4657 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1flag_1map, tmp4646)

	_ = tmp4657

	tmp4658 := MakeNative(func(__e *ControlFlow) {
		V4914 := __e.Get(1)
		_ = V4914
		V4915 := __e.Get(2)
		_ = V4915
		tmp4762 := PrimEqual(Nil, V4914)

		if True == tmp4762 {
			tmp4761 := Call(__e, PrimNS2Value(symreverse), V4915)

			__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4execute_1all), tmp4761)
			return

		} else {
			tmp4760 := PrimIsPair(V4914)

			var ifres4752 Obj

			if True == tmp4760 {
				tmp4758 := PrimHead(V4914)

				tmp4759 := PrimEqual(MakeString("--eval"), tmp4758)

				var ifres4754 Obj

				if True == tmp4759 {
					tmp4756 := PrimTail(V4914)

					tmp4757 := PrimIsPair(tmp4756)

					var ifres4755 Obj

					if True == tmp4757 {
						ifres4755 = True

					} else {
						ifres4755 = False

					}

					ifres4754 = ifres4755

				} else {
					ifres4754 = False

				}

				var ifres4753 Obj

				if True == ifres4754 {
					ifres4753 = True

				} else {
					ifres4753 = False

				}

				ifres4752 = ifres4753

			} else {
				ifres4752 = False

			}

			if True == ifres4752 {
				tmp4743 := PrimTail(V4914)

				tmp4744 := PrimTail(tmp4743)

				tmp4745 := MakeNative(func(__e *ControlFlow) {
					tmp4746 := PrimTail(V4914)

					tmp4747 := PrimHead(tmp4746)

					tmp4748 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4eval_1string), tmp4747)

					tmp4749 := Call(__e, PrimNS2Value(symshen_4app), tmp4748, MakeString("\n"), symshen_4a)

					tmp4750 := Call(__e, PrimNS2Value(symstoutput))

					__e.TailApply(PrimNS2Value(symshen_4prhush), tmp4749, tmp4750)
					return

				}, 0)

				tmp4751 := PrimCons(tmp4745, V4915)

				__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), tmp4744, tmp4751)
				return

			} else {
				tmp4742 := PrimIsPair(V4914)

				var ifres4734 Obj

				if True == tmp4742 {
					tmp4740 := PrimHead(V4914)

					tmp4741 := PrimEqual(MakeString("--load"), tmp4740)

					var ifres4736 Obj

					if True == tmp4741 {
						tmp4738 := PrimTail(V4914)

						tmp4739 := PrimIsPair(tmp4738)

						var ifres4737 Obj

						if True == tmp4739 {
							ifres4737 = True

						} else {
							ifres4737 = False

						}

						ifres4736 = ifres4737

					} else {
						ifres4736 = False

					}

					var ifres4735 Obj

					if True == ifres4736 {
						ifres4735 = True

					} else {
						ifres4735 = False

					}

					ifres4734 = ifres4735

				} else {
					ifres4734 = False

				}

				if True == ifres4734 {
					tmp4728 := PrimTail(V4914)

					tmp4729 := PrimTail(tmp4728)

					tmp4730 := MakeNative(func(__e *ControlFlow) {
						tmp4731 := PrimTail(V4914)

						tmp4732 := PrimHead(tmp4731)

						__e.TailApply(PrimNS2Value(symload), tmp4732)
						return

					}, 0)

					tmp4733 := PrimCons(tmp4730, V4915)

					__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), tmp4729, tmp4733)
					return

				} else {
					tmp4727 := PrimIsPair(V4914)

					var ifres4723 Obj

					if True == tmp4727 {
						tmp4725 := PrimHead(V4914)

						tmp4726 := PrimEqual(MakeString("--quiet"), tmp4725)

						var ifres4724 Obj

						if True == tmp4726 {
							ifres4724 = True

						} else {
							ifres4724 = False

						}

						ifres4723 = ifres4724

					} else {
						ifres4723 = False

					}

					if True == ifres4723 {
						tmp4720 := PrimTail(V4914)

						tmp4721 := MakeNative(func(__e *ControlFlow) {
							__e.Return(PrimNS3Set(sym_dhush_d, True))
							return
						}, 0)

						tmp4722 := PrimCons(tmp4721, V4915)

						__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), tmp4720, tmp4722)
						return

					} else {
						tmp4719 := PrimIsPair(V4914)

						var ifres4706 Obj

						if True == tmp4719 {
							tmp4717 := PrimHead(V4914)

							tmp4718 := PrimEqual(MakeString("--set"), tmp4717)

							var ifres4708 Obj

							if True == tmp4718 {
								tmp4715 := PrimTail(V4914)

								tmp4716 := PrimIsPair(tmp4715)

								var ifres4710 Obj

								if True == tmp4716 {
									tmp4712 := PrimTail(V4914)

									tmp4713 := PrimTail(tmp4712)

									tmp4714 := PrimIsPair(tmp4713)

									var ifres4711 Obj

									if True == tmp4714 {
										ifres4711 = True

									} else {
										ifres4711 = False

									}

									ifres4710 = ifres4711

								} else {
									ifres4710 = False

								}

								var ifres4709 Obj

								if True == ifres4710 {
									ifres4709 = True

								} else {
									ifres4709 = False

								}

								ifres4708 = ifres4709

							} else {
								ifres4708 = False

							}

							var ifres4707 Obj

							if True == ifres4708 {
								ifres4707 = True

							} else {
								ifres4707 = False

							}

							ifres4706 = ifres4707

						} else {
							ifres4706 = False

						}

						if True == ifres4706 {
							tmp4694 := PrimTail(V4914)

							tmp4695 := PrimTail(tmp4694)

							tmp4696 := PrimTail(tmp4695)

							tmp4697 := MakeNative(func(__e *ControlFlow) {
								tmp4698 := PrimTail(V4914)

								tmp4699 := PrimHead(tmp4698)

								tmp4700 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4eval_1string), tmp4699)

								tmp4701 := PrimTail(V4914)

								tmp4702 := PrimTail(tmp4701)

								tmp4703 := PrimHead(tmp4702)

								tmp4704 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4eval_1string), tmp4703)

								__e.Return(PrimNS3Set(tmp4700, tmp4704))
								return

							}, 0)

							tmp4705 := PrimCons(tmp4697, V4915)

							__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), tmp4696, tmp4705)
							return

						} else {
							tmp4693 := PrimIsPair(V4914)

							var ifres4689 Obj

							if True == tmp4693 {
								tmp4691 := PrimHead(V4914)

								tmp4692 := PrimEqual(MakeString("--repl"), tmp4691)

								var ifres4690 Obj

								if True == tmp4692 {
									ifres4690 = True

								} else {
									ifres4690 = False

								}

								ifres4689 = ifres4690

							} else {
								ifres4689 = False

							}

							if True == ifres4689 {
								tmp4687 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), Nil, V4915)

								_ = tmp4687

								tmp4688 := PrimTail(V4914)

								__e.Return(PrimCons(symlaunch_1repl, tmp4688))
								return

							} else {
								tmp4665 := MakeNative(func(__e *ControlFlow) {
									Freeze := __e.Get(1)
									_ = Freeze
									tmp4679 := PrimIsPair(V4914)

									if True == tmp4679 {
										tmp4667 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp4669 := Call(__e, PrimNS2Value(symfail))

											tmp4670 := PrimEqual(Result, tmp4669)

											if True == tmp4670 {
												__e.TailApply(PrimNS2Value(symthaw), Freeze)
												return
											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp4671 := MakeNative(func(__e *ControlFlow) {
											Long := __e.Get(1)
											_ = Long
											tmp4675 := PrimEqual(False, Long)

											if True == tmp4675 {
												__e.TailApply(PrimNS2Value(symfail))
												return
											} else {
												tmp4673 := PrimTail(V4914)

												tmp4674 := PrimCons(Long, tmp4673)

												__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), tmp4674, V4915)
												return

											}

										}, 1)

										tmp4676 := PrimHead(V4914)

										tmp4677 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4eval_1flag_1map), tmp4676)

										tmp4678 := Call(__e, tmp4671, tmp4677)

										__e.TailApply(tmp4667, tmp4678)
										return

									} else {
										__e.TailApply(PrimNS2Value(symthaw), Freeze)
										return
									}

								}, 1)

								tmp4680 := MakeNative(func(__e *ControlFlow) {
									tmp4686 := PrimIsPair(V4914)

									if True == tmp4686 {
										tmp4682 := PrimHead(V4914)

										tmp4683 := Call(__e, PrimNS2Value(symshen_4app), tmp4682, MakeString(""), symshen_4a)

										tmp4684 := PrimStringConcat(MakeString("Invalid eval argument: "), tmp4683)

										tmp4685 := PrimCons(tmp4684, Nil)

										__e.Return(PrimCons(symerror, tmp4685))
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4launcher_4eval_1command_1h)
										return
									}

								}, 0)

								__e.TailApply(tmp4665, tmp4680)
								return

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp4763 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1command_1h, tmp4658)

	_ = tmp4763

	tmp4764 := MakeNative(func(__e *ControlFlow) {
		V4917 := __e.Get(1)
		_ = V4917
		__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command_1h), V4917, Nil)
		return
	}, 1)

	tmp4765 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1command, tmp4764)

	_ = tmp4765

	tmp4766 := MakeNative(func(__e *ControlFlow) {
		V4920 := __e.Get(1)
		_ = V4920
		V4921 := __e.Get(2)
		_ = V4921
		tmp4767 := PrimCons(V4920, V4921)

		tmp4768 := PrimNS3Set(sym_dargv_d, tmp4767)

		_ = tmp4768

		tmp4769 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4quiet_1load), V4920)

		_ = tmp4769

		__e.Return(PrimCons(symsuccess, Nil))
		return

	}, 2)

	tmp4770 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4script_1command, tmp4766)

	_ = tmp4770

	tmp4771 := MakeNative(func(__e *ControlFlow) {
		V4923 := __e.Get(1)
		_ = V4923
		tmp4858 := PrimIsPair(V4923)

		var ifres4854 Obj

		if True == tmp4858 {
			tmp4856 := PrimTail(V4923)

			tmp4857 := PrimEqual(Nil, tmp4856)

			var ifres4855 Obj

			if True == tmp4857 {
				ifres4855 = True

			} else {
				ifres4855 = False

			}

			ifres4854 = ifres4855

		} else {
			ifres4854 = False

		}

		if True == ifres4854 {
			__e.Return(PrimCons(symlaunch_1repl, Nil))
			return
		} else {
			tmp4853 := PrimIsPair(V4923)

			var ifres4844 Obj

			if True == tmp4853 {
				tmp4851 := PrimTail(V4923)

				tmp4852 := PrimIsPair(tmp4851)

				var ifres4846 Obj

				if True == tmp4852 {
					tmp4848 := PrimTail(V4923)

					tmp4849 := PrimHead(tmp4848)

					tmp4850 := PrimEqual(MakeString("--help"), tmp4849)

					var ifres4847 Obj

					if True == tmp4850 {
						ifres4847 = True

					} else {
						ifres4847 = False

					}

					ifres4846 = ifres4847

				} else {
					ifres4846 = False

				}

				var ifres4845 Obj

				if True == ifres4846 {
					ifres4845 = True

				} else {
					ifres4845 = False

				}

				ifres4844 = ifres4845

			} else {
				ifres4844 = False

			}

			if True == ifres4844 {
				tmp4841 := PrimHead(V4923)

				tmp4842 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4help_1text), tmp4841)

				tmp4843 := PrimCons(tmp4842, Nil)

				__e.Return(PrimCons(symshow_1help, tmp4843))
				return

			} else {
				tmp4840 := PrimIsPair(V4923)

				var ifres4831 Obj

				if True == tmp4840 {
					tmp4838 := PrimTail(V4923)

					tmp4839 := PrimIsPair(tmp4838)

					var ifres4833 Obj

					if True == tmp4839 {
						tmp4835 := PrimTail(V4923)

						tmp4836 := PrimHead(tmp4835)

						tmp4837 := PrimEqual(MakeString("--version"), tmp4836)

						var ifres4834 Obj

						if True == tmp4837 {
							ifres4834 = True

						} else {
							ifres4834 = False

						}

						ifres4833 = ifres4834

					} else {
						ifres4833 = False

					}

					var ifres4832 Obj

					if True == ifres4833 {
						ifres4832 = True

					} else {
						ifres4832 = False

					}

					ifres4831 = ifres4832

				} else {
					ifres4831 = False

				}

				if True == ifres4831 {
					tmp4829 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4version_1string))

					tmp4830 := PrimCons(tmp4829, Nil)

					__e.Return(PrimCons(symsuccess, tmp4830))
					return

				} else {
					tmp4828 := PrimIsPair(V4923)

					var ifres4819 Obj

					if True == tmp4828 {
						tmp4826 := PrimTail(V4923)

						tmp4827 := PrimIsPair(tmp4826)

						var ifres4821 Obj

						if True == tmp4827 {
							tmp4823 := PrimTail(V4923)

							tmp4824 := PrimHead(tmp4823)

							tmp4825 := PrimEqual(MakeString("repl"), tmp4824)

							var ifres4822 Obj

							if True == tmp4825 {
								ifres4822 = True

							} else {
								ifres4822 = False

							}

							ifres4821 = ifres4822

						} else {
							ifres4821 = False

						}

						var ifres4820 Obj

						if True == ifres4821 {
							ifres4820 = True

						} else {
							ifres4820 = False

						}

						ifres4819 = ifres4820

					} else {
						ifres4819 = False

					}

					if True == ifres4819 {
						tmp4817 := PrimTail(V4923)

						tmp4818 := PrimTail(tmp4817)

						__e.Return(PrimCons(symlaunch_1repl, tmp4818))
						return

					} else {
						tmp4816 := PrimIsPair(V4923)

						var ifres4802 Obj

						if True == tmp4816 {
							tmp4814 := PrimTail(V4923)

							tmp4815 := PrimIsPair(tmp4814)

							var ifres4804 Obj

							if True == tmp4815 {
								tmp4811 := PrimTail(V4923)

								tmp4812 := PrimHead(tmp4811)

								tmp4813 := PrimEqual(MakeString("script"), tmp4812)

								var ifres4806 Obj

								if True == tmp4813 {
									tmp4808 := PrimTail(V4923)

									tmp4809 := PrimTail(tmp4808)

									tmp4810 := PrimIsPair(tmp4809)

									var ifres4807 Obj

									if True == tmp4810 {
										ifres4807 = True

									} else {
										ifres4807 = False

									}

									ifres4806 = ifres4807

								} else {
									ifres4806 = False

								}

								var ifres4805 Obj

								if True == ifres4806 {
									ifres4805 = True

								} else {
									ifres4805 = False

								}

								ifres4804 = ifres4805

							} else {
								ifres4804 = False

							}

							var ifres4803 Obj

							if True == ifres4804 {
								ifres4803 = True

							} else {
								ifres4803 = False

							}

							ifres4802 = ifres4803

						} else {
							ifres4802 = False

						}

						if True == ifres4802 {
							tmp4796 := PrimTail(V4923)

							tmp4797 := PrimTail(tmp4796)

							tmp4798 := PrimHead(tmp4797)

							tmp4799 := PrimTail(V4923)

							tmp4800 := PrimTail(tmp4799)

							tmp4801 := PrimTail(tmp4800)

							__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4script_1command), tmp4798, tmp4801)
							return

						} else {
							tmp4795 := PrimIsPair(V4923)

							var ifres4786 Obj

							if True == tmp4795 {
								tmp4793 := PrimTail(V4923)

								tmp4794 := PrimIsPair(tmp4793)

								var ifres4788 Obj

								if True == tmp4794 {
									tmp4790 := PrimTail(V4923)

									tmp4791 := PrimHead(tmp4790)

									tmp4792 := PrimEqual(MakeString("eval"), tmp4791)

									var ifres4789 Obj

									if True == tmp4792 {
										ifres4789 = True

									} else {
										ifres4789 = False

									}

									ifres4788 = ifres4789

								} else {
									ifres4788 = False

								}

								var ifres4787 Obj

								if True == ifres4788 {
									ifres4787 = True

								} else {
									ifres4787 = False

								}

								ifres4786 = ifres4787

							} else {
								ifres4786 = False

							}

							if True == ifres4786 {
								tmp4784 := PrimTail(V4923)

								tmp4785 := PrimTail(tmp4784)

								__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4eval_1command), tmp4785)
								return

							} else {
								tmp4783 := PrimIsPair(V4923)

								var ifres4779 Obj

								if True == tmp4783 {
									tmp4781 := PrimTail(V4923)

									tmp4782 := PrimIsPair(tmp4781)

									var ifres4780 Obj

									if True == tmp4782 {
										ifres4780 = True

									} else {
										ifres4780 = False

									}

									ifres4779 = ifres4780

								} else {
									ifres4779 = False

								}

								if True == ifres4779 {
									__e.Return(PrimCons(symunknown_1arguments, V4923))
									return
								} else {
									__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4launcher_4launch_1shen)
									return
								}

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp4859 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4launch_1shen, tmp4771)

	_ = tmp4859

	tmp4860 := MakeNative(func(__e *ControlFlow) {
		V4927 := __e.Get(1)
		_ = V4927
		tmp4959 := PrimIsPair(V4927)

		var ifres4951 Obj

		if True == tmp4959 {
			tmp4957 := PrimHead(V4927)

			tmp4958 := PrimEqual(symsuccess, tmp4957)

			var ifres4953 Obj

			if True == tmp4958 {
				tmp4955 := PrimTail(V4927)

				tmp4956 := PrimEqual(Nil, tmp4955)

				var ifres4954 Obj

				if True == tmp4956 {
					ifres4954 = True

				} else {
					ifres4954 = False

				}

				ifres4953 = ifres4954

			} else {
				ifres4953 = False

			}

			var ifres4952 Obj

			if True == ifres4953 {
				ifres4952 = True

			} else {
				ifres4952 = False

			}

			ifres4951 = ifres4952

		} else {
			ifres4951 = False

		}

		if True == ifres4951 {
			__e.Return(symshen_4x_4launcher_4done)
			return
		} else {
			tmp4950 := PrimIsPair(V4927)

			var ifres4937 Obj

			if True == tmp4950 {
				tmp4948 := PrimHead(V4927)

				tmp4949 := PrimEqual(symsuccess, tmp4948)

				var ifres4939 Obj

				if True == tmp4949 {
					tmp4946 := PrimTail(V4927)

					tmp4947 := PrimIsPair(tmp4946)

					var ifres4941 Obj

					if True == tmp4947 {
						tmp4943 := PrimTail(V4927)

						tmp4944 := PrimTail(tmp4943)

						tmp4945 := PrimEqual(Nil, tmp4944)

						var ifres4942 Obj

						if True == tmp4945 {
							ifres4942 = True

						} else {
							ifres4942 = False

						}

						ifres4941 = ifres4942

					} else {
						ifres4941 = False

					}

					var ifres4940 Obj

					if True == ifres4941 {
						ifres4940 = True

					} else {
						ifres4940 = False

					}

					ifres4939 = ifres4940

				} else {
					ifres4939 = False

				}

				var ifres4938 Obj

				if True == ifres4939 {
					ifres4938 = True

				} else {
					ifres4938 = False

				}

				ifres4937 = ifres4938

			} else {
				ifres4937 = False

			}

			if True == ifres4937 {
				tmp4933 := PrimTail(V4927)

				tmp4934 := PrimHead(tmp4933)

				tmp4935 := Call(__e, PrimNS2Value(symshen_4app), tmp4934, MakeString("\n"), symshen_4a)

				tmp4936 := Call(__e, PrimNS2Value(symstoutput))

				__e.TailApply(PrimNS2Value(symshen_4prhush), tmp4935, tmp4936)
				return

			} else {
				tmp4932 := PrimIsPair(V4927)

				var ifres4919 Obj

				if True == tmp4932 {
					tmp4930 := PrimHead(V4927)

					tmp4931 := PrimEqual(symerror, tmp4930)

					var ifres4921 Obj

					if True == tmp4931 {
						tmp4928 := PrimTail(V4927)

						tmp4929 := PrimIsPair(tmp4928)

						var ifres4923 Obj

						if True == tmp4929 {
							tmp4925 := PrimTail(V4927)

							tmp4926 := PrimTail(tmp4925)

							tmp4927 := PrimEqual(Nil, tmp4926)

							var ifres4924 Obj

							if True == tmp4927 {
								ifres4924 = True

							} else {
								ifres4924 = False

							}

							ifres4923 = ifres4924

						} else {
							ifres4923 = False

						}

						var ifres4922 Obj

						if True == ifres4923 {
							ifres4922 = True

						} else {
							ifres4922 = False

						}

						ifres4921 = ifres4922

					} else {
						ifres4921 = False

					}

					var ifres4920 Obj

					if True == ifres4921 {
						ifres4920 = True

					} else {
						ifres4920 = False

					}

					ifres4919 = ifres4920

				} else {
					ifres4919 = False

				}

				if True == ifres4919 {
					tmp4914 := PrimTail(V4927)

					tmp4915 := PrimHead(tmp4914)

					tmp4916 := Call(__e, PrimNS2Value(symshen_4app), tmp4915, MakeString("\n"), symshen_4a)

					tmp4917 := PrimStringConcat(MakeString("ERROR: "), tmp4916)

					tmp4918 := Call(__e, PrimNS2Value(symstoutput))

					__e.TailApply(PrimNS2Value(symshen_4prhush), tmp4917, tmp4918)
					return

				} else {
					tmp4913 := PrimIsPair(V4927)

					var ifres4909 Obj

					if True == tmp4913 {
						tmp4911 := PrimHead(V4927)

						tmp4912 := PrimEqual(symlaunch_1repl, tmp4911)

						var ifres4910 Obj

						if True == tmp4912 {
							ifres4910 = True

						} else {
							ifres4910 = False

						}

						ifres4909 = ifres4910

					} else {
						ifres4909 = False

					}

					if True == ifres4909 {
						__e.TailApply(PrimNS2Value(symshen_4repl))
						return
					} else {
						tmp4908 := PrimIsPair(V4927)

						var ifres4895 Obj

						if True == tmp4908 {
							tmp4906 := PrimHead(V4927)

							tmp4907 := PrimEqual(symshow_1help, tmp4906)

							var ifres4897 Obj

							if True == tmp4907 {
								tmp4904 := PrimTail(V4927)

								tmp4905 := PrimIsPair(tmp4904)

								var ifres4899 Obj

								if True == tmp4905 {
									tmp4901 := PrimTail(V4927)

									tmp4902 := PrimTail(tmp4901)

									tmp4903 := PrimEqual(Nil, tmp4902)

									var ifres4900 Obj

									if True == tmp4903 {
										ifres4900 = True

									} else {
										ifres4900 = False

									}

									ifres4899 = ifres4900

								} else {
									ifres4899 = False

								}

								var ifres4898 Obj

								if True == ifres4899 {
									ifres4898 = True

								} else {
									ifres4898 = False

								}

								ifres4897 = ifres4898

							} else {
								ifres4897 = False

							}

							var ifres4896 Obj

							if True == ifres4897 {
								ifres4896 = True

							} else {
								ifres4896 = False

							}

							ifres4895 = ifres4896

						} else {
							ifres4895 = False

						}

						if True == ifres4895 {
							tmp4891 := PrimTail(V4927)

							tmp4892 := PrimHead(tmp4891)

							tmp4893 := Call(__e, PrimNS2Value(symshen_4app), tmp4892, MakeString("\n"), symshen_4a)

							tmp4894 := Call(__e, PrimNS2Value(symstoutput))

							__e.TailApply(PrimNS2Value(symshen_4prhush), tmp4893, tmp4894)
							return

						} else {
							tmp4890 := PrimIsPair(V4927)

							var ifres4877 Obj

							if True == tmp4890 {
								tmp4888 := PrimHead(V4927)

								tmp4889 := PrimEqual(symunknown_1arguments, tmp4888)

								var ifres4879 Obj

								if True == tmp4889 {
									tmp4886 := PrimTail(V4927)

									tmp4887 := PrimIsPair(tmp4886)

									var ifres4881 Obj

									if True == tmp4887 {
										tmp4883 := PrimTail(V4927)

										tmp4884 := PrimTail(tmp4883)

										tmp4885 := PrimIsPair(tmp4884)

										var ifres4882 Obj

										if True == tmp4885 {
											ifres4882 = True

										} else {
											ifres4882 = False

										}

										ifres4881 = ifres4882

									} else {
										ifres4881 = False

									}

									var ifres4880 Obj

									if True == ifres4881 {
										ifres4880 = True

									} else {
										ifres4880 = False

									}

									ifres4879 = ifres4880

								} else {
									ifres4879 = False

								}

								var ifres4878 Obj

								if True == ifres4879 {
									ifres4878 = True

								} else {
									ifres4878 = False

								}

								ifres4877 = ifres4878

							} else {
								ifres4877 = False

							}

							if True == ifres4877 {
								tmp4867 := PrimTail(V4927)

								tmp4868 := PrimTail(tmp4867)

								tmp4869 := PrimHead(tmp4868)

								tmp4870 := PrimTail(V4927)

								tmp4871 := PrimHead(tmp4870)

								tmp4872 := Call(__e, PrimNS2Value(symshen_4app), tmp4871, MakeString(" --help' for more information.\n"), symshen_4a)

								tmp4873 := PrimStringConcat(MakeString("\nTry `"), tmp4872)

								tmp4874 := Call(__e, PrimNS2Value(symshen_4app), tmp4869, tmp4873, symshen_4a)

								tmp4875 := PrimStringConcat(MakeString("ERROR: Invalid argument: "), tmp4874)

								tmp4876 := Call(__e, PrimNS2Value(symstoutput))

								__e.TailApply(PrimNS2Value(symshen_4prhush), tmp4875, tmp4876)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4launcher_4default_1handle_1result)
								return
							}

						}

					}

				}

			}

		}

	}, 1)

	tmp4960 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4default_1handle_1result, tmp4860)

	_ = tmp4960

	tmp4961 := MakeNative(func(__e *ControlFlow) {
		V4929 := __e.Get(1)
		_ = V4929
		tmp4962 := Call(__e, PrimNS2Value(symshen_4x_4launcher_4launch_1shen), V4929)

		__e.TailApply(PrimNS2Value(symshen_4x_4launcher_4default_1handle_1result), tmp4962)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4launcher_4main, tmp4961)
	return

}, 0)
