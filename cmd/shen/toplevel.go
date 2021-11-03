package main

import . "github.com/tiancaiamao/shen-go/cora"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
	tmp6817 := MakeNative(func(__e *ControlFlow) {
		tmp6818 := Call(__e, PrimNS2Value(symshen_4credits))

		_ = tmp6818

		__e.TailApply(PrimNS2Value(symshen_4loop))
		return

	}, 0)

	tmp6819 := Call(__e, PrimNS2Value(symdef), symshen_4shen, tmp6817)

	_ = tmp6819

	tmp6820 := MakeNative(func(__e *ControlFlow) {
		tmp6821 := Call(__e, PrimNS2Value(symshen_4initialise__environment))

		_ = tmp6821

		tmp6822 := Call(__e, PrimNS2Value(symshen_4prompt))

		_ = tmp6822

		tmp6823 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4read_1evaluate_1print))
			return
		}, 0)

		tmp6824 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp6825 := PrimErrorToString(E)

			tmp6826 := Call(__e, PrimNS2Value(symstoutput))

			tmp6827 := Call(__e, PrimNS2Value(sympr), tmp6825, tmp6826)

			_ = tmp6827

			__e.TailApply(PrimNS2Value(symnl), MakeNumber(0))
			return

		}, 1)

		tmp6828 := Call(__e, PrimNS1Value(symtry_1catch), tmp6823, tmp6824)

		_ = tmp6828

		__e.TailApply(PrimNS2Value(symshen_4loop))
		return

	}, 0)

	tmp6829 := Call(__e, PrimNS2Value(symdef), symshen_4loop, tmp6820)

	_ = tmp6829

	tmp6830 := MakeNative(func(__e *ControlFlow) {
		tmp6831 := Call(__e, PrimNS2Value(symstoutput))

		tmp6832 := Call(__e, PrimNS2Value(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2021, Mark Tarver\n"), tmp6831)

		_ = tmp6832

		tmp6833 := PrimNS3Value(sym_dversion_d)

		tmp6834 := PrimNS3Value(sym_dlanguage_d)

		tmp6835 := PrimNS3Value(sym_dimplementation_d)

		tmp6836 := PrimNS3Value(sym_drelease_d)

		tmp6837 := Call(__e, PrimNS2Value(symshen_4app), tmp6836, MakeString("\n"), symshen_4a)

		tmp6838 := PrimStringConcat(MakeString(" "), tmp6837)

		tmp6839 := Call(__e, PrimNS2Value(symshen_4app), tmp6835, tmp6838, symshen_4a)

		tmp6840 := PrimStringConcat(MakeString(", platform: "), tmp6839)

		tmp6841 := Call(__e, PrimNS2Value(symshen_4app), tmp6834, tmp6840, symshen_4a)

		tmp6842 := PrimStringConcat(MakeString(", language: "), tmp6841)

		tmp6843 := Call(__e, PrimNS2Value(symshen_4app), tmp6833, tmp6842, symshen_4a)

		tmp6844 := PrimStringConcat(MakeString("version: S"), tmp6843)

		tmp6845 := Call(__e, PrimNS2Value(symstoutput))

		tmp6846 := Call(__e, PrimNS2Value(sympr), tmp6844, tmp6845)

		_ = tmp6846

		tmp6847 := PrimNS3Value(sym_dport_d)

		tmp6848 := PrimNS3Value(sym_dporters_d)

		tmp6849 := Call(__e, PrimNS2Value(symshen_4app), tmp6848, MakeString("\n\n"), symshen_4a)

		tmp6850 := PrimStringConcat(MakeString(", ported by "), tmp6849)

		tmp6851 := Call(__e, PrimNS2Value(symshen_4app), tmp6847, tmp6850, symshen_4a)

		tmp6852 := PrimStringConcat(MakeString("port "), tmp6851)

		tmp6853 := Call(__e, PrimNS2Value(symstoutput))

		__e.TailApply(PrimNS2Value(sympr), tmp6852, tmp6853)
		return

	}, 0)

	tmp6854 := Call(__e, PrimNS2Value(symdef), symshen_4credits, tmp6830)

	_ = tmp6854

	tmp6855 := MakeNative(func(__e *ControlFlow) {
		tmp6856 := PrimNS3Set(symshen_4_dcall_d, MakeNumber(0))

		_ = tmp6856

		__e.Return(PrimNS3Set(symshen_4_dinfs_d, MakeNumber(0)))
		return

	}, 0)

	tmp6857 := Call(__e, PrimNS2Value(symdef), symshen_4initialise__environment, tmp6855)

	_ = tmp6857

	tmp6858 := MakeNative(func(__e *ControlFlow) {
		tmp6870 := PrimNS3Value(symshen_4_dtc_d)

		if True == tmp6870 {
			tmp6865 := PrimNS3Value(symshen_4_dhistory_d)

			tmp6866 := Call(__e, PrimNS2Value(symlength), tmp6865)

			tmp6867 := Call(__e, PrimNS2Value(symshen_4app), tmp6866, MakeString("+) "), symshen_4a)

			tmp6868 := PrimStringConcat(MakeString("\n("), tmp6867)

			tmp6869 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp6868, tmp6869)
			return

		} else {
			tmp6860 := PrimNS3Value(symshen_4_dhistory_d)

			tmp6861 := Call(__e, PrimNS2Value(symlength), tmp6860)

			tmp6862 := Call(__e, PrimNS2Value(symshen_4app), tmp6861, MakeString("-) "), symshen_4a)

			tmp6863 := PrimStringConcat(MakeString("\n("), tmp6862)

			tmp6864 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp6863, tmp6864)
			return

		}

	}, 0)

	tmp6871 := Call(__e, PrimNS2Value(symdef), symshen_4prompt, tmp6858)

	_ = tmp6871

	tmp6872 := MakeNative(func(__e *ControlFlow) {
		tmp6873 := MakeNative(func(__e *ControlFlow) {
			Package := __e.Get(1)
			_ = Package
			tmp6874 := MakeNative(func(__e *ControlFlow) {
				Lineread := __e.Get(1)
				_ = Lineread
				tmp6875 := MakeNative(func(__e *ControlFlow) {
					History := __e.Get(1)
					_ = History
					tmp6876 := PrimNS3Value(symshen_4_dtc_d)

					__e.TailApply(PrimNS2Value(symshen_4evaluate_1lineread), Lineread, History, tmp6876)
					return

				}, 1)

				tmp6877 := Call(__e, PrimNS2Value(symshen_4update_1history))

				__e.TailApply(tmp6875, tmp6877)
				return

			}, 1)

			tmp6878 := Call(__e, PrimNS2Value(symstinput))

			tmp6879 := Call(__e, PrimNS2Value(symlineread), tmp6878)

			tmp6880 := Call(__e, PrimNS2Value(symshen_4package_1user_1input), Package, tmp6879)

			__e.TailApply(tmp6874, tmp6880)
			return

		}, 1)

		tmp6881 := PrimNS3Value(symshen_4_dpackage_d)

		__e.TailApply(tmp6873, tmp6881)
		return

	}, 0)

	tmp6882 := Call(__e, PrimNS2Value(symdef), symshen_4read_1evaluate_1print, tmp6872)

	_ = tmp6882

	tmp6883 := MakeNative(func(__e *ControlFlow) {
		V4467 := __e.Get(1)
		_ = V4467
		V4468 := __e.Get(2)
		_ = V4468
		tmp6890 := PrimEqual(symnull, V4467)

		if True == tmp6890 {
			__e.Return(V4468)
			return
		} else {
			tmp6885 := MakeNative(func(__e *ControlFlow) {
				Str := __e.Get(1)
				_ = Str
				tmp6886 := MakeNative(func(__e *ControlFlow) {
					External := __e.Get(1)
					_ = External
					tmp6887 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4pui_1h), Str, External, X)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symmap), tmp6887, V4468)
					return

				}, 1)

				tmp6888 := Call(__e, PrimNS2Value(symexternal), V4467)

				__e.TailApply(tmp6886, tmp6888)
				return

			}, 1)

			tmp6889 := PrimStr(V4467)

			__e.TailApply(tmp6885, tmp6889)
			return

		}

	}, 2)

	tmp6891 := Call(__e, PrimNS2Value(symdef), symshen_4package_1user_1input, tmp6883)

	_ = tmp6891

	tmp6892 := MakeNative(func(__e *ControlFlow) {
		V4473 := __e.Get(1)
		_ = V4473
		V4474 := __e.Get(2)
		_ = V4474
		V4475 := __e.Get(3)
		_ = V4475
		tmp6933 := PrimIsPair(V4475)

		var ifres6920 Obj

		if True == tmp6933 {
			tmp6931 := PrimHead(V4475)

			tmp6932 := PrimEqual(symfn, tmp6931)

			var ifres6922 Obj

			if True == tmp6932 {
				tmp6929 := PrimTail(V4475)

				tmp6930 := PrimIsPair(tmp6929)

				var ifres6924 Obj

				if True == tmp6930 {
					tmp6926 := PrimTail(V4475)

					tmp6927 := PrimTail(tmp6926)

					tmp6928 := PrimEqual(Nil, tmp6927)

					var ifres6925 Obj

					if True == tmp6928 {
						ifres6925 = True

					} else {
						ifres6925 = False

					}

					ifres6924 = ifres6925

				} else {
					ifres6924 = False

				}

				var ifres6923 Obj

				if True == ifres6924 {
					ifres6923 = True

				} else {
					ifres6923 = False

				}

				ifres6922 = ifres6923

			} else {
				ifres6922 = False

			}

			var ifres6921 Obj

			if True == ifres6922 {
				ifres6921 = True

			} else {
				ifres6921 = False

			}

			ifres6920 = ifres6921

		} else {
			ifres6920 = False

		}

		if True == ifres6920 {
			tmp6917 := PrimTail(V4475)

			tmp6918 := PrimHead(tmp6917)

			tmp6919 := Call(__e, PrimNS2Value(symshen_4internal_2), tmp6918, V4473, V4474)

			if True == tmp6919 {
				tmp6913 := PrimTail(V4475)

				tmp6914 := PrimHead(tmp6913)

				tmp6915 := Call(__e, PrimNS2Value(symshen_4intern_1in_1package), V4473, tmp6914)

				tmp6916 := PrimCons(tmp6915, Nil)

				__e.Return(PrimCons(symfn, tmp6916))
				return

			} else {
				__e.Return(V4475)
				return
			}

		} else {
			tmp6911 := PrimIsPair(V4475)

			if True == tmp6911 {
				tmp6909 := PrimHead(V4475)

				tmp6910 := Call(__e, PrimNS2Value(symshen_4internal_2), tmp6909, V4473, V4474)

				if True == tmp6910 {
					tmp6904 := PrimHead(V4475)

					tmp6905 := Call(__e, PrimNS2Value(symshen_4intern_1in_1package), V4473, tmp6904)

					tmp6906 := MakeNative(func(__e *ControlFlow) {
						Y := __e.Get(1)
						_ = Y
						__e.TailApply(PrimNS2Value(symshen_4pui_1h), V4473, V4474, Y)
						return
					}, 1)

					tmp6907 := PrimTail(V4475)

					tmp6908 := Call(__e, PrimNS2Value(symmap), tmp6906, tmp6907)

					__e.Return(PrimCons(tmp6905, tmp6908))
					return

				} else {
					tmp6902 := PrimHead(V4475)

					tmp6903 := PrimIsPair(tmp6902)

					if True == tmp6903 {
						tmp6901 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							__e.TailApply(PrimNS2Value(symshen_4pui_1h), V4473, V4474, Y)
							return
						}, 1)

						__e.TailApply(PrimNS2Value(symmap), tmp6901, V4475)
						return

					} else {
						tmp6897 := PrimHead(V4475)

						tmp6898 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							__e.TailApply(PrimNS2Value(symshen_4pui_1h), V4473, V4474, Y)
							return
						}, 1)

						tmp6899 := PrimTail(V4475)

						tmp6900 := Call(__e, PrimNS2Value(symmap), tmp6898, tmp6899)

						__e.Return(PrimCons(tmp6897, tmp6900))
						return

					}

				}

			} else {
				__e.Return(V4475)
				return
			}

		}

	}, 3)

	tmp6934 := Call(__e, PrimNS2Value(symdef), symshen_4pui_1h, tmp6892)

	_ = tmp6934

	tmp6935 := MakeNative(func(__e *ControlFlow) {
		tmp6936 := Call(__e, PrimNS2Value(symit))

		tmp6937 := PrimNS3Value(symshen_4_dhistory_d)

		tmp6938 := PrimCons(tmp6936, tmp6937)

		__e.Return(PrimNS3Set(symshen_4_dhistory_d, tmp6938))
		return

	}, 0)

	tmp6939 := Call(__e, PrimNS2Value(symdef), symshen_4update_1history, tmp6935)

	_ = tmp6939

	tmp6940 := MakeNative(func(__e *ControlFlow) {
		V4486 := __e.Get(1)
		_ = V4486
		V4487 := __e.Get(2)
		_ = V4487
		V4488 := __e.Get(3)
		_ = V4488
		tmp7077 := PrimIsPair(V4486)

		var ifres7062 Obj

		if True == tmp7077 {
			tmp7075 := PrimTail(V4486)

			tmp7076 := PrimEqual(Nil, tmp7075)

			var ifres7064 Obj

			if True == tmp7076 {
				tmp7074 := PrimIsPair(V4487)

				var ifres7066 Obj

				if True == tmp7074 {
					tmp7072 := PrimHead(V4487)

					tmp7073 := PrimEqual(MakeString("!!"), tmp7072)

					var ifres7068 Obj

					if True == tmp7073 {
						tmp7070 := PrimTail(V4487)

						tmp7071 := PrimIsPair(tmp7070)

						var ifres7069 Obj

						if True == tmp7071 {
							ifres7069 = True

						} else {
							ifres7069 = False

						}

						ifres7068 = ifres7069

					} else {
						ifres7068 = False

					}

					var ifres7067 Obj

					if True == ifres7068 {
						ifres7067 = True

					} else {
						ifres7067 = False

					}

					ifres7066 = ifres7067

				} else {
					ifres7066 = False

				}

				var ifres7065 Obj

				if True == ifres7066 {
					ifres7065 = True

				} else {
					ifres7065 = False

				}

				ifres7064 = ifres7065

			} else {
				ifres7064 = False

			}

			var ifres7063 Obj

			if True == ifres7064 {
				ifres7063 = True

			} else {
				ifres7063 = False

			}

			ifres7062 = ifres7063

		} else {
			ifres7062 = False

		}

		if True == ifres7062 {
			tmp7046 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp7047 := MakeNative(func(__e *ControlFlow) {
					NewHistory := __e.Get(1)
					_ = NewHistory
					tmp7048 := MakeNative(func(__e *ControlFlow) {
						Print := __e.Get(1)
						_ = Print
						__e.TailApply(PrimNS2Value(symshen_4evaluate_1lineread), Y, NewHistory, V4488)
						return
					}, 1)

					tmp7049 := PrimTail(V4487)

					tmp7050 := PrimHead(tmp7049)

					tmp7051 := Call(__e, PrimNS2Value(symshen_4app), tmp7050, MakeString("\n"), symshen_4a)

					tmp7052 := Call(__e, PrimNS2Value(symstoutput))

					tmp7053 := Call(__e, PrimNS2Value(sympr), tmp7051, tmp7052)

					__e.TailApply(tmp7048, tmp7053)
					return

				}, 1)

				tmp7054 := PrimTail(V4487)

				tmp7055 := PrimHead(tmp7054)

				tmp7056 := PrimTail(V4487)

				tmp7057 := PrimCons(tmp7055, tmp7056)

				tmp7058 := PrimNS3Set(symshen_4_dhistory_d, tmp7057)

				__e.TailApply(tmp7047, tmp7058)
				return

			}, 1)

			tmp7059 := PrimTail(V4487)

			tmp7060 := PrimHead(tmp7059)

			tmp7061 := Call(__e, PrimNS2Value(symread_1from_1string), tmp7060)

			__e.TailApply(tmp7046, tmp7061)
			return

		} else {
			tmp7045 := PrimIsPair(V4486)

			var ifres7029 Obj

			if True == tmp7045 {
				tmp7043 := PrimTail(V4486)

				tmp7044 := PrimEqual(Nil, tmp7043)

				var ifres7031 Obj

				if True == tmp7044 {
					tmp7042 := PrimIsPair(V4487)

					var ifres7033 Obj

					if True == tmp7042 {
						tmp7040 := PrimHead(V4487)

						tmp7041 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp7040)

						var ifres7035 Obj

						if True == tmp7041 {
							tmp7037 := PrimHead(V4487)

							tmp7038 := Call(__e, PrimNS2Value(symhdstr), tmp7037)

							tmp7039 := PrimEqual(MakeString("%"), tmp7038)

							var ifres7036 Obj

							if True == tmp7039 {
								ifres7036 = True

							} else {
								ifres7036 = False

							}

							ifres7035 = ifres7036

						} else {
							ifres7035 = False

						}

						var ifres7034 Obj

						if True == ifres7035 {
							ifres7034 = True

						} else {
							ifres7034 = False

						}

						ifres7033 = ifres7034

					} else {
						ifres7033 = False

					}

					var ifres7032 Obj

					if True == ifres7033 {
						ifres7032 = True

					} else {
						ifres7032 = False

					}

					ifres7031 = ifres7032

				} else {
					ifres7031 = False

				}

				var ifres7030 Obj

				if True == ifres7031 {
					ifres7030 = True

				} else {
					ifres7030 = False

				}

				ifres7029 = ifres7030

			} else {
				ifres7029 = False

			}

			if True == ifres7029 {
				tmp7016 := MakeNative(func(__e *ControlFlow) {
					Read := __e.Get(1)
					_ = Read
					tmp7017 := MakeNative(func(__e *ControlFlow) {
						Peek := __e.Get(1)
						_ = Peek
						tmp7018 := MakeNative(func(__e *ControlFlow) {
							NewHistory := __e.Get(1)
							_ = NewHistory
							__e.TailApply(PrimNS2Value(symabort))
							return
						}, 1)

						tmp7019 := PrimTail(V4487)

						tmp7020 := PrimNS3Set(symshen_4_dhistory_d, tmp7019)

						__e.TailApply(tmp7018, tmp7020)
						return

					}, 1)

					tmp7021 := PrimHead(V4487)

					tmp7022 := PrimTailString(tmp7021)

					tmp7023 := PrimTail(V4487)

					tmp7024 := Call(__e, PrimNS2Value(symshen_4peek_1history), Read, tmp7022, tmp7023)

					__e.TailApply(tmp7017, tmp7024)
					return

				}, 1)

				tmp7025 := PrimHead(V4487)

				tmp7026 := PrimTailString(tmp7025)

				tmp7027 := Call(__e, PrimNS2Value(symread_1from_1string), tmp7026)

				tmp7028 := PrimHead(tmp7027)

				__e.TailApply(tmp7016, tmp7028)
				return

			} else {
				tmp7015 := PrimIsPair(V4486)

				var ifres6999 Obj

				if True == tmp7015 {
					tmp7013 := PrimTail(V4486)

					tmp7014 := PrimEqual(Nil, tmp7013)

					var ifres7001 Obj

					if True == tmp7014 {
						tmp7012 := PrimIsPair(V4487)

						var ifres7003 Obj

						if True == tmp7012 {
							tmp7010 := PrimHead(V4487)

							tmp7011 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp7010)

							var ifres7005 Obj

							if True == tmp7011 {
								tmp7007 := PrimHead(V4487)

								tmp7008 := Call(__e, PrimNS2Value(symhdstr), tmp7007)

								tmp7009 := PrimEqual(MakeString("!"), tmp7008)

								var ifres7006 Obj

								if True == tmp7009 {
									ifres7006 = True

								} else {
									ifres7006 = False

								}

								ifres7005 = ifres7006

							} else {
								ifres7005 = False

							}

							var ifres7004 Obj

							if True == ifres7005 {
								ifres7004 = True

							} else {
								ifres7004 = False

							}

							ifres7003 = ifres7004

						} else {
							ifres7003 = False

						}

						var ifres7002 Obj

						if True == ifres7003 {
							ifres7002 = True

						} else {
							ifres7002 = False

						}

						ifres7001 = ifres7002

					} else {
						ifres7001 = False

					}

					var ifres7000 Obj

					if True == ifres7001 {
						ifres7000 = True

					} else {
						ifres7000 = False

					}

					ifres6999 = ifres7000

				} else {
					ifres6999 = False

				}

				if True == ifres6999 {
					tmp6979 := MakeNative(func(__e *ControlFlow) {
						Read := __e.Get(1)
						_ = Read
						tmp6980 := MakeNative(func(__e *ControlFlow) {
							Match := __e.Get(1)
							_ = Match
							tmp6981 := MakeNative(func(__e *ControlFlow) {
								Print := __e.Get(1)
								_ = Print
								tmp6982 := MakeNative(func(__e *ControlFlow) {
									Y := __e.Get(1)
									_ = Y
									tmp6983 := MakeNative(func(__e *ControlFlow) {
										NewHistory := __e.Get(1)
										_ = NewHistory
										__e.TailApply(PrimNS2Value(symshen_4evaluate_1lineread), Y, NewHistory, V4488)
										return
									}, 1)

									tmp6984 := PrimTail(V4487)

									tmp6985 := PrimCons(Match, tmp6984)

									tmp6986 := PrimNS3Set(symshen_4_dhistory_d, tmp6985)

									__e.TailApply(tmp6983, tmp6986)
									return

								}, 1)

								tmp6987 := Call(__e, PrimNS2Value(symread_1from_1string), Match)

								__e.TailApply(tmp6982, tmp6987)
								return

							}, 1)

							tmp6988 := Call(__e, PrimNS2Value(symshen_4app), Match, MakeString("\n"), symshen_4a)

							tmp6989 := Call(__e, PrimNS2Value(symstoutput))

							tmp6990 := Call(__e, PrimNS2Value(sympr), tmp6988, tmp6989)

							__e.TailApply(tmp6981, tmp6990)
							return

						}, 1)

						tmp6991 := PrimHead(V4487)

						tmp6992 := PrimTailString(tmp6991)

						tmp6993 := PrimTail(V4487)

						tmp6994 := Call(__e, PrimNS2Value(symshen_4use_1history), Read, tmp6992, tmp6993)

						__e.TailApply(tmp6980, tmp6994)
						return

					}, 1)

					tmp6995 := PrimHead(V4487)

					tmp6996 := PrimTailString(tmp6995)

					tmp6997 := Call(__e, PrimNS2Value(symread_1from_1string), tmp6996)

					tmp6998 := PrimHead(tmp6997)

					__e.TailApply(tmp6979, tmp6998)
					return

				} else {
					tmp6978 := PrimIsPair(V4486)

					var ifres6962 Obj

					if True == tmp6978 {
						tmp6976 := PrimTail(V4486)

						tmp6977 := PrimEqual(Nil, tmp6976)

						var ifres6964 Obj

						if True == tmp6977 {
							tmp6975 := PrimIsPair(V4487)

							var ifres6966 Obj

							if True == tmp6975 {
								tmp6973 := PrimHead(V4487)

								tmp6974 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp6973)

								var ifres6968 Obj

								if True == tmp6974 {
									tmp6970 := PrimHead(V4487)

									tmp6971 := Call(__e, PrimNS2Value(symhdstr), tmp6970)

									tmp6972 := PrimEqual(MakeString("%"), tmp6971)

									var ifres6969 Obj

									if True == tmp6972 {
										ifres6969 = True

									} else {
										ifres6969 = False

									}

									ifres6968 = ifres6969

								} else {
									ifres6968 = False

								}

								var ifres6967 Obj

								if True == ifres6968 {
									ifres6967 = True

								} else {
									ifres6967 = False

								}

								ifres6966 = ifres6967

							} else {
								ifres6966 = False

							}

							var ifres6965 Obj

							if True == ifres6966 {
								ifres6965 = True

							} else {
								ifres6965 = False

							}

							ifres6964 = ifres6965

						} else {
							ifres6964 = False

						}

						var ifres6963 Obj

						if True == ifres6964 {
							ifres6963 = True

						} else {
							ifres6963 = False

						}

						ifres6962 = ifres6963

					} else {
						ifres6962 = False

					}

					if True == ifres6962 {
						tmp6949 := MakeNative(func(__e *ControlFlow) {
							Read := __e.Get(1)
							_ = Read
							tmp6950 := MakeNative(func(__e *ControlFlow) {
								Peek := __e.Get(1)
								_ = Peek
								tmp6951 := MakeNative(func(__e *ControlFlow) {
									NewHistory := __e.Get(1)
									_ = NewHistory
									__e.TailApply(PrimNS2Value(symabort))
									return
								}, 1)

								tmp6952 := PrimTail(V4487)

								tmp6953 := PrimNS3Set(symshen_4_dhistory_d, tmp6952)

								__e.TailApply(tmp6951, tmp6953)
								return

							}, 1)

							tmp6954 := PrimHead(V4487)

							tmp6955 := PrimTailString(tmp6954)

							tmp6956 := PrimTail(V4487)

							tmp6957 := Call(__e, PrimNS2Value(symshen_4peek_1history), Read, tmp6955, tmp6956)

							__e.TailApply(tmp6950, tmp6957)
							return

						}, 1)

						tmp6958 := PrimHead(V4487)

						tmp6959 := PrimTailString(tmp6958)

						tmp6960 := Call(__e, PrimNS2Value(symread_1from_1string), tmp6959)

						tmp6961 := PrimHead(tmp6960)

						__e.TailApply(tmp6949, tmp6961)
						return

					} else {
						tmp6948 := PrimEqual(True, V4488)

						if True == tmp6948 {
							__e.TailApply(PrimNS2Value(symshen_4check_1eval_1and_1print), V4486)
							return
						} else {
							tmp6947 := PrimEqual(False, V4488)

							if True == tmp6947 {
								__e.TailApply(PrimNS2Value(symshen_4eval_1and_1print), V4486)
								return
							} else {
								__e.Return(PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread")))
								return
							}

						}

					}

				}

			}

		}

	}, 3)

	tmp7078 := Call(__e, PrimNS2Value(symdef), symshen_4evaluate_1lineread, tmp6940)

	_ = tmp7078

	tmp7079 := MakeNative(func(__e *ControlFlow) {
		V4489 := __e.Get(1)
		_ = V4489
		V4490 := __e.Get(2)
		_ = V4490
		V4491 := __e.Get(3)
		_ = V4491
		tmp7085 := PrimIsInteger(V4489)

		if True == tmp7085 {
			tmp7083 := PrimNumberAdd(MakeNumber(1), V4489)

			tmp7084 := Call(__e, PrimNS2Value(symreverse), V4491)

			__e.TailApply(PrimNS2Value(symnth), tmp7083, tmp7084)
			return

		} else {
			tmp7082 := PrimIsSymbol(V4489)

			if True == tmp7082 {
				__e.TailApply(PrimNS2Value(symshen_4string_1match), V4490, V4491)
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
				return
			}

		}

	}, 3)

	tmp7086 := Call(__e, PrimNS2Value(symdef), symshen_4use_1history, tmp7079)

	_ = tmp7086

	tmp7087 := MakeNative(func(__e *ControlFlow) {
		V4492 := __e.Get(1)
		_ = V4492
		V4493 := __e.Get(2)
		_ = V4493
		V4494 := __e.Get(3)
		_ = V4494
		tmp7101 := PrimIsInteger(V4492)

		if True == tmp7101 {
			tmp7095 := PrimNumberAdd(MakeNumber(1), V4492)

			tmp7096 := Call(__e, PrimNS2Value(symreverse), V4494)

			tmp7097 := Call(__e, PrimNS2Value(symnth), tmp7095, tmp7096)

			tmp7098 := Call(__e, PrimNS2Value(symshen_4app), tmp7097, MakeString(""), symshen_4a)

			tmp7099 := PrimStringConcat(MakeString("\n"), tmp7098)

			tmp7100 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp7099, tmp7100)
			return

		} else {
			tmp7094 := PrimEqual(V4493, MakeString(""))

			var ifres7091 Obj

			if True == tmp7094 {
				ifres7091 = True

			} else {
				tmp7093 := PrimIsSymbol(V4492)

				var ifres7092 Obj

				if True == tmp7093 {
					ifres7092 = True

				} else {
					ifres7092 = False

				}

				ifres7091 = ifres7092

			}

			if True == ifres7091 {
				tmp7090 := Call(__e, PrimNS2Value(symreverse), V4494)

				__e.TailApply(PrimNS2Value(symshen_4recursive_1string_1match), MakeNumber(0), V4493, tmp7090)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
				return
			}

		}

	}, 3)

	tmp7102 := Call(__e, PrimNS2Value(symdef), symshen_4peek_1history, tmp7087)

	_ = tmp7102

	tmp7103 := MakeNative(func(__e *ControlFlow) {
		V4504 := __e.Get(1)
		_ = V4504
		V4505 := __e.Get(2)
		_ = V4505
		tmp7114 := PrimEqual(Nil, V4505)

		if True == tmp7114 {
			__e.Return(PrimSimpleError(MakeString("\ninput not found")))
			return
		} else {
			tmp7113 := PrimIsPair(V4505)

			var ifres7109 Obj

			if True == tmp7113 {
				tmp7111 := PrimHead(V4505)

				tmp7112 := Call(__e, PrimNS2Value(symshen_4string_1prefix_2), V4504, tmp7111)

				var ifres7110 Obj

				if True == tmp7112 {
					ifres7110 = True

				} else {
					ifres7110 = False

				}

				ifres7109 = ifres7110

			} else {
				ifres7109 = False

			}

			if True == ifres7109 {
				__e.Return(PrimHead(V4505))
				return
			} else {
				tmp7108 := PrimIsPair(V4505)

				if True == tmp7108 {
					tmp7107 := PrimTail(V4505)

					__e.TailApply(PrimNS2Value(symshen_4string_1match), V4504, tmp7107)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
					return
				}

			}

		}

	}, 2)

	tmp7115 := Call(__e, PrimNS2Value(symdef), symshen_4string_1match, tmp7103)

	_ = tmp7115

	tmp7116 := MakeNative(func(__e *ControlFlow) {
		V4513 := __e.Get(1)
		_ = V4513
		V4514 := __e.Get(2)
		_ = V4514
		tmp7153 := PrimEqual(MakeString(""), V4513)

		if True == tmp7153 {
			__e.Return(True)
			return
		} else {
			tmp7152 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4513)

			var ifres7147 Obj

			if True == tmp7152 {
				tmp7149 := Call(__e, PrimNS2Value(symhdstr), V4513)

				tmp7150 := PrimStringToNumber(tmp7149)

				tmp7151 := Call(__e, PrimNS2Value(symshen_4whitespace_2), tmp7150)

				var ifres7148 Obj

				if True == tmp7151 {
					ifres7148 = True

				} else {
					ifres7148 = False

				}

				ifres7147 = ifres7148

			} else {
				ifres7147 = False

			}

			if True == ifres7147 {
				tmp7146 := PrimTailString(V4513)

				__e.TailApply(PrimNS2Value(symshen_4string_1prefix_2), tmp7146, V4514)
				return

			} else {
				tmp7145 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4514)

				var ifres7140 Obj

				if True == tmp7145 {
					tmp7142 := Call(__e, PrimNS2Value(symhdstr), V4514)

					tmp7143 := PrimStringToNumber(tmp7142)

					tmp7144 := Call(__e, PrimNS2Value(symshen_4whitespace_2), tmp7143)

					var ifres7141 Obj

					if True == tmp7144 {
						ifres7141 = True

					} else {
						ifres7141 = False

					}

					ifres7140 = ifres7141

				} else {
					ifres7140 = False

				}

				if True == ifres7140 {
					tmp7139 := PrimTailString(V4514)

					__e.TailApply(PrimNS2Value(symshen_4string_1prefix_2), V4513, tmp7139)
					return

				} else {
					tmp7138 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4514)

					var ifres7134 Obj

					if True == tmp7138 {
						tmp7136 := Call(__e, PrimNS2Value(symhdstr), V4514)

						tmp7137 := PrimEqual(MakeString("("), tmp7136)

						var ifres7135 Obj

						if True == tmp7137 {
							ifres7135 = True

						} else {
							ifres7135 = False

						}

						ifres7134 = ifres7135

					} else {
						ifres7134 = False

					}

					if True == ifres7134 {
						tmp7133 := PrimTailString(V4514)

						__e.TailApply(PrimNS2Value(symshen_4string_1prefix_2), V4513, tmp7133)
						return

					} else {
						tmp7132 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4513)

						var ifres7124 Obj

						if True == tmp7132 {
							tmp7131 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4514)

							var ifres7126 Obj

							if True == tmp7131 {
								tmp7128 := Call(__e, PrimNS2Value(symhdstr), V4513)

								tmp7129 := Call(__e, PrimNS2Value(symhdstr), V4514)

								tmp7130 := PrimEqual(tmp7128, tmp7129)

								var ifres7127 Obj

								if True == tmp7130 {
									ifres7127 = True

								} else {
									ifres7127 = False

								}

								ifres7126 = ifres7127

							} else {
								ifres7126 = False

							}

							var ifres7125 Obj

							if True == ifres7126 {
								ifres7125 = True

							} else {
								ifres7125 = False

							}

							ifres7124 = ifres7125

						} else {
							ifres7124 = False

						}

						if True == ifres7124 {
							tmp7122 := PrimTailString(V4513)

							tmp7123 := PrimTailString(V4514)

							__e.TailApply(PrimNS2Value(symshen_4string_1prefix_2), tmp7122, tmp7123)
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

	tmp7154 := Call(__e, PrimNS2Value(symdef), symshen_4string_1prefix_2, tmp7116)

	_ = tmp7154

	tmp7155 := MakeNative(func(__e *ControlFlow) {
		V4525 := __e.Get(1)
		_ = V4525
		V4526 := __e.Get(2)
		_ = V4526
		V4527 := __e.Get(3)
		_ = V4527
		tmp7170 := PrimEqual(Nil, V4527)

		if True == tmp7170 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp7169 := PrimIsPair(V4527)

			if True == tmp7169 {
				tmp7165 := PrimHead(V4527)

				tmp7166 := Call(__e, PrimNS2Value(symshen_4string_1prefix_2), V4526, tmp7165)

				var ifres7158 Obj

				if True == tmp7166 {
					tmp7159 := PrimHead(V4527)

					tmp7160 := Call(__e, PrimNS2Value(symshen_4app), tmp7159, MakeString("\n"), symshen_4a)

					tmp7161 := PrimStringConcat(MakeString(". "), tmp7160)

					tmp7162 := Call(__e, PrimNS2Value(symshen_4app), V4525, tmp7161, symshen_4a)

					tmp7163 := Call(__e, PrimNS2Value(symstoutput))

					tmp7164 := Call(__e, PrimNS2Value(sympr), tmp7162, tmp7163)

					ifres7158 = tmp7164

				} else {
					ifres7158 = symshen_4skip

				}

				_ = ifres7158

				tmp7167 := PrimNumberAdd(V4525, MakeNumber(1))

				tmp7168 := PrimTail(V4527)

				__e.TailApply(PrimNS2Value(symshen_4recursive_1string_1match), tmp7167, V4526, tmp7168)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
				return
			}

		}

	}, 3)

	__e.TailApply(PrimNS2Value(symdef), symshen_4recursive_1string_1match, tmp7155)
	return

}, 0)
