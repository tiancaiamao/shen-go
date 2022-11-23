package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
	tmp7795 := MakeNative(func(__e *ControlFlow) {
		tmp7796 := Call(__e, PrimFunc(symshen_4credits))

		_ = tmp7796

		__e.TailApply(PrimFunc(symshen_4loop))
		return

	}, 0)

	tmp7797 := Call(__e, ns2_1set, symshen_4shen, tmp7795)

	_ = tmp7797

	tmp7798 := MakeNative(func(__e *ControlFlow) {
		tmp7799 := Call(__e, PrimFunc(symshen_4initialise__environment))

		_ = tmp7799

		tmp7800 := Call(__e, PrimFunc(symshen_4prompt))

		_ = tmp7800

		tmp7801 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
			return
		}, 0)

		tmp7802 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp7803 := PrimErrorToString(E)

			tmp7804 := Call(__e, PrimFunc(symstoutput))

			tmp7805 := Call(__e, PrimFunc(sympr), tmp7803, tmp7804)

			_ = tmp7805

			__e.TailApply(PrimFunc(symnl), MakeNumber(0))
			return

		}, 1)

		tmp7806 := Call(__e, try_1catch, tmp7801, tmp7802)

		_ = tmp7806

		__e.TailApply(PrimFunc(symshen_4loop))
		return

	}, 0)

	tmp7807 := Call(__e, ns2_1set, symshen_4loop, tmp7798)

	_ = tmp7807

	tmp7808 := MakeNative(func(__e *ControlFlow) {
		tmp7809 := Call(__e, PrimFunc(symstoutput))

		tmp7810 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2021, Mark Tarver\n"), tmp7809)

		_ = tmp7810

		tmp7811 := PrimValue(sym_dversion_d)

		tmp7812 := PrimValue(sym_dlanguage_d)

		tmp7813 := PrimValue(sym_dimplementation_d)

		tmp7814 := PrimValue(sym_drelease_d)

		tmp7815 := Call(__e, PrimFunc(symshen_4app), tmp7814, MakeString("\n"), symshen_4a)

		tmp7816 := PrimStringConcat(MakeString(" "), tmp7815)

		tmp7817 := Call(__e, PrimFunc(symshen_4app), tmp7813, tmp7816, symshen_4a)

		tmp7818 := PrimStringConcat(MakeString(", platform: "), tmp7817)

		tmp7819 := Call(__e, PrimFunc(symshen_4app), tmp7812, tmp7818, symshen_4a)

		tmp7820 := PrimStringConcat(MakeString(", language: "), tmp7819)

		tmp7821 := Call(__e, PrimFunc(symshen_4app), tmp7811, tmp7820, symshen_4a)

		tmp7822 := PrimStringConcat(MakeString("version: S"), tmp7821)

		tmp7823 := Call(__e, PrimFunc(symstoutput))

		tmp7824 := Call(__e, PrimFunc(sympr), tmp7822, tmp7823)

		_ = tmp7824

		tmp7825 := PrimValue(sym_dport_d)

		tmp7826 := PrimValue(sym_dporters_d)

		tmp7827 := Call(__e, PrimFunc(symshen_4app), tmp7826, MakeString("\n\n"), symshen_4a)

		tmp7828 := PrimStringConcat(MakeString(", ported by "), tmp7827)

		tmp7829 := Call(__e, PrimFunc(symshen_4app), tmp7825, tmp7828, symshen_4a)

		tmp7830 := PrimStringConcat(MakeString("port "), tmp7829)

		tmp7831 := Call(__e, PrimFunc(symstoutput))

		__e.TailApply(PrimFunc(sympr), tmp7830, tmp7831)
		return

	}, 0)

	tmp7832 := Call(__e, ns2_1set, symshen_4credits, tmp7808)

	_ = tmp7832

	tmp7833 := MakeNative(func(__e *ControlFlow) {
		tmp7834 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

		_ = tmp7834

		__e.Return(PrimSet(symshen_4_dinfs_d, MakeNumber(0)))
		return

	}, 0)

	tmp7835 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp7833)

	_ = tmp7835

	tmp7836 := MakeNative(func(__e *ControlFlow) {
		tmp7848 := PrimValue(symshen_4_dtc_d)

		if True == tmp7848 {
			tmp7837 := PrimValue(symshen_4_dhistory_d)

			tmp7838 := Call(__e, PrimFunc(symlength), tmp7837)

			tmp7839 := Call(__e, PrimFunc(symshen_4app), tmp7838, MakeString("+) "), symshen_4a)

			tmp7840 := PrimStringConcat(MakeString("\n("), tmp7839)

			tmp7841 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp7840, tmp7841)
			return

		} else {
			tmp7842 := PrimValue(symshen_4_dhistory_d)

			tmp7843 := Call(__e, PrimFunc(symlength), tmp7842)

			tmp7844 := Call(__e, PrimFunc(symshen_4app), tmp7843, MakeString("-) "), symshen_4a)

			tmp7845 := PrimStringConcat(MakeString("\n("), tmp7844)

			tmp7846 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp7845, tmp7846)
			return

		}

	}, 0)

	tmp7849 := Call(__e, ns2_1set, symshen_4prompt, tmp7836)

	_ = tmp7849

	tmp7850 := MakeNative(func(__e *ControlFlow) {
		tmp7851 := MakeNative(func(__e *ControlFlow) {
			Package := __e.Get(1)
			_ = Package
			tmp7852 := MakeNative(func(__e *ControlFlow) {
				Lineread := __e.Get(1)
				_ = Lineread
				tmp7853 := MakeNative(func(__e *ControlFlow) {
					History := __e.Get(1)
					_ = History
					tmp7854 := PrimValue(symshen_4_dtc_d)

					__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), Lineread, History, tmp7854)
					return

				}, 1)

				tmp7855 := Call(__e, PrimFunc(symshen_4update_1history))

				__e.TailApply(tmp7853, tmp7855)
				return

			}, 1)

			tmp7856 := Call(__e, PrimFunc(symstinput))

			tmp7857 := Call(__e, PrimFunc(symlineread), tmp7856)

			tmp7858 := Call(__e, PrimFunc(symshen_4package_1user_1input), Package, tmp7857)

			__e.TailApply(tmp7852, tmp7858)
			return

		}, 1)

		tmp7859 := PrimValue(symshen_4_dpackage_d)

		__e.TailApply(tmp7851, tmp7859)
		return

	}, 0)

	tmp7860 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp7850)

	_ = tmp7860

	tmp7861 := MakeNative(func(__e *ControlFlow) {
		V4467 := __e.Get(1)
		_ = V4467
		V4468 := __e.Get(2)
		_ = V4468
		tmp7868 := PrimEqual(symnull, V4467)

		if True == tmp7868 {
			__e.Return(V4468)
			return
		} else {
			tmp7862 := MakeNative(func(__e *ControlFlow) {
				Str := __e.Get(1)
				_ = Str
				tmp7863 := MakeNative(func(__e *ControlFlow) {
					External := __e.Get(1)
					_ = External
					tmp7864 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimFunc(symshen_4pui_1h), Str, External, X)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp7864, V4468)
					return

				}, 1)

				tmp7865 := Call(__e, PrimFunc(symexternal), V4467)

				__e.TailApply(tmp7863, tmp7865)
				return

			}, 1)

			tmp7866 := PrimStr(V4467)

			__e.TailApply(tmp7862, tmp7866)
			return

		}

	}, 2)

	tmp7869 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp7861)

	_ = tmp7869

	tmp7870 := MakeNative(func(__e *ControlFlow) {
		V4473 := __e.Get(1)
		_ = V4473
		V4474 := __e.Get(2)
		_ = V4474
		V4475 := __e.Get(3)
		_ = V4475
		tmp7911 := PrimIsPair(V4475)

		var ifres7898 Obj

		if True == tmp7911 {
			tmp7909 := PrimHead(V4475)

			tmp7910 := PrimEqual(symfn, tmp7909)

			var ifres7900 Obj

			if True == tmp7910 {
				tmp7907 := PrimTail(V4475)

				tmp7908 := PrimIsPair(tmp7907)

				var ifres7902 Obj

				if True == tmp7908 {
					tmp7904 := PrimTail(V4475)

					tmp7905 := PrimTail(tmp7904)

					tmp7906 := PrimEqual(Nil, tmp7905)

					var ifres7903 Obj

					if True == tmp7906 {
						ifres7903 = True

					} else {
						ifres7903 = False

					}

					ifres7902 = ifres7903

				} else {
					ifres7902 = False

				}

				var ifres7901 Obj

				if True == ifres7902 {
					ifres7901 = True

				} else {
					ifres7901 = False

				}

				ifres7900 = ifres7901

			} else {
				ifres7900 = False

			}

			var ifres7899 Obj

			if True == ifres7900 {
				ifres7899 = True

			} else {
				ifres7899 = False

			}

			ifres7898 = ifres7899

		} else {
			ifres7898 = False

		}

		if True == ifres7898 {
			tmp7876 := PrimTail(V4475)

			tmp7877 := PrimHead(tmp7876)

			tmp7878 := Call(__e, PrimFunc(symshen_4internal_2), tmp7877, V4473, V4474)

			if True == tmp7878 {
				tmp7871 := PrimTail(V4475)

				tmp7872 := PrimHead(tmp7871)

				tmp7873 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V4473, tmp7872)

				tmp7874 := PrimCons(tmp7873, Nil)

				__e.Return(PrimCons(symfn, tmp7874))
				return

			} else {
				__e.Return(V4475)
				return
			}

		} else {
			tmp7896 := PrimIsPair(V4475)

			if True == tmp7896 {
				tmp7893 := PrimHead(V4475)

				tmp7894 := Call(__e, PrimFunc(symshen_4internal_2), tmp7893, V4473, V4474)

				if True == tmp7894 {
					tmp7879 := PrimHead(V4475)

					tmp7880 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V4473, tmp7879)

					tmp7881 := MakeNative(func(__e *ControlFlow) {
						Y := __e.Get(1)
						_ = Y
						__e.TailApply(PrimFunc(symshen_4pui_1h), V4473, V4474, Y)
						return
					}, 1)

					tmp7882 := PrimTail(V4475)

					tmp7883 := Call(__e, PrimFunc(symmap), tmp7881, tmp7882)

					__e.Return(PrimCons(tmp7880, tmp7883))
					return

				} else {
					tmp7890 := PrimHead(V4475)

					tmp7891 := PrimIsPair(tmp7890)

					if True == tmp7891 {
						tmp7884 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							__e.TailApply(PrimFunc(symshen_4pui_1h), V4473, V4474, Y)
							return
						}, 1)

						__e.TailApply(PrimFunc(symmap), tmp7884, V4475)
						return

					} else {
						tmp7885 := PrimHead(V4475)

						tmp7886 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							__e.TailApply(PrimFunc(symshen_4pui_1h), V4473, V4474, Y)
							return
						}, 1)

						tmp7887 := PrimTail(V4475)

						tmp7888 := Call(__e, PrimFunc(symmap), tmp7886, tmp7887)

						__e.Return(PrimCons(tmp7885, tmp7888))
						return

					}

				}

			} else {
				__e.Return(V4475)
				return
			}

		}

	}, 3)

	tmp7912 := Call(__e, ns2_1set, symshen_4pui_1h, tmp7870)

	_ = tmp7912

	tmp7913 := MakeNative(func(__e *ControlFlow) {
		tmp7914 := Call(__e, PrimFunc(symit))

		tmp7915 := PrimValue(symshen_4_dhistory_d)

		tmp7916 := PrimCons(tmp7914, tmp7915)

		__e.Return(PrimSet(symshen_4_dhistory_d, tmp7916))
		return

	}, 0)

	tmp7917 := Call(__e, ns2_1set, symshen_4update_1history, tmp7913)

	_ = tmp7917

	tmp7918 := MakeNative(func(__e *ControlFlow) {
		V4486 := __e.Get(1)
		_ = V4486
		V4487 := __e.Get(2)
		_ = V4487
		V4488 := __e.Get(3)
		_ = V4488
		tmp8055 := PrimIsPair(V4486)

		var ifres8040 Obj

		if True == tmp8055 {
			tmp8053 := PrimTail(V4486)

			tmp8054 := PrimEqual(Nil, tmp8053)

			var ifres8042 Obj

			if True == tmp8054 {
				tmp8052 := PrimIsPair(V4487)

				var ifres8044 Obj

				if True == tmp8052 {
					tmp8050 := PrimHead(V4487)

					tmp8051 := PrimEqual(MakeString("!!"), tmp8050)

					var ifres8046 Obj

					if True == tmp8051 {
						tmp8048 := PrimTail(V4487)

						tmp8049 := PrimIsPair(tmp8048)

						var ifres8047 Obj

						if True == tmp8049 {
							ifres8047 = True

						} else {
							ifres8047 = False

						}

						ifres8046 = ifres8047

					} else {
						ifres8046 = False

					}

					var ifres8045 Obj

					if True == ifres8046 {
						ifres8045 = True

					} else {
						ifres8045 = False

					}

					ifres8044 = ifres8045

				} else {
					ifres8044 = False

				}

				var ifres8043 Obj

				if True == ifres8044 {
					ifres8043 = True

				} else {
					ifres8043 = False

				}

				ifres8042 = ifres8043

			} else {
				ifres8042 = False

			}

			var ifres8041 Obj

			if True == ifres8042 {
				ifres8041 = True

			} else {
				ifres8041 = False

			}

			ifres8040 = ifres8041

		} else {
			ifres8040 = False

		}

		if True == ifres8040 {
			tmp7919 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp7920 := MakeNative(func(__e *ControlFlow) {
					NewHistory := __e.Get(1)
					_ = NewHistory
					tmp7921 := MakeNative(func(__e *ControlFlow) {
						Print := __e.Get(1)
						_ = Print
						__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), Y, NewHistory, V4488)
						return
					}, 1)

					tmp7922 := PrimTail(V4487)

					tmp7923 := PrimHead(tmp7922)

					tmp7924 := Call(__e, PrimFunc(symshen_4app), tmp7923, MakeString("\n"), symshen_4a)

					tmp7925 := Call(__e, PrimFunc(symstoutput))

					tmp7926 := Call(__e, PrimFunc(sympr), tmp7924, tmp7925)

					__e.TailApply(tmp7921, tmp7926)
					return

				}, 1)

				tmp7927 := PrimTail(V4487)

				tmp7928 := PrimHead(tmp7927)

				tmp7929 := PrimTail(V4487)

				tmp7930 := PrimCons(tmp7928, tmp7929)

				tmp7931 := PrimSet(symshen_4_dhistory_d, tmp7930)

				__e.TailApply(tmp7920, tmp7931)
				return

			}, 1)

			tmp7932 := PrimTail(V4487)

			tmp7933 := PrimHead(tmp7932)

			tmp7934 := Call(__e, PrimFunc(symread_1from_1string), tmp7933)

			__e.TailApply(tmp7919, tmp7934)
			return

		} else {
			tmp8038 := PrimIsPair(V4486)

			var ifres8022 Obj

			if True == tmp8038 {
				tmp8036 := PrimTail(V4486)

				tmp8037 := PrimEqual(Nil, tmp8036)

				var ifres8024 Obj

				if True == tmp8037 {
					tmp8035 := PrimIsPair(V4487)

					var ifres8026 Obj

					if True == tmp8035 {
						tmp8033 := PrimHead(V4487)

						tmp8034 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8033)

						var ifres8028 Obj

						if True == tmp8034 {
							tmp8030 := PrimHead(V4487)

							tmp8031 := Call(__e, PrimFunc(symhdstr), tmp8030)

							tmp8032 := PrimEqual(MakeString("%"), tmp8031)

							var ifres8029 Obj

							if True == tmp8032 {
								ifres8029 = True

							} else {
								ifres8029 = False

							}

							ifres8028 = ifres8029

						} else {
							ifres8028 = False

						}

						var ifres8027 Obj

						if True == ifres8028 {
							ifres8027 = True

						} else {
							ifres8027 = False

						}

						ifres8026 = ifres8027

					} else {
						ifres8026 = False

					}

					var ifres8025 Obj

					if True == ifres8026 {
						ifres8025 = True

					} else {
						ifres8025 = False

					}

					ifres8024 = ifres8025

				} else {
					ifres8024 = False

				}

				var ifres8023 Obj

				if True == ifres8024 {
					ifres8023 = True

				} else {
					ifres8023 = False

				}

				ifres8022 = ifres8023

			} else {
				ifres8022 = False

			}

			if True == ifres8022 {
				tmp7935 := MakeNative(func(__e *ControlFlow) {
					Read := __e.Get(1)
					_ = Read
					tmp7936 := MakeNative(func(__e *ControlFlow) {
						Peek := __e.Get(1)
						_ = Peek
						tmp7937 := MakeNative(func(__e *ControlFlow) {
							NewHistory := __e.Get(1)
							_ = NewHistory
							__e.TailApply(PrimFunc(symabort))
							return
						}, 1)

						tmp7938 := PrimTail(V4487)

						tmp7939 := PrimSet(symshen_4_dhistory_d, tmp7938)

						__e.TailApply(tmp7937, tmp7939)
						return

					}, 1)

					tmp7940 := PrimHead(V4487)

					tmp7941 := PrimTailString(tmp7940)

					tmp7942 := PrimTail(V4487)

					tmp7943 := Call(__e, PrimFunc(symshen_4peek_1history), Read, tmp7941, tmp7942)

					__e.TailApply(tmp7936, tmp7943)
					return

				}, 1)

				tmp7944 := PrimHead(V4487)

				tmp7945 := PrimTailString(tmp7944)

				tmp7946 := Call(__e, PrimFunc(symread_1from_1string), tmp7945)

				tmp7947 := PrimHead(tmp7946)

				__e.TailApply(tmp7935, tmp7947)
				return

			} else {
				tmp8020 := PrimIsPair(V4486)

				var ifres8004 Obj

				if True == tmp8020 {
					tmp8018 := PrimTail(V4486)

					tmp8019 := PrimEqual(Nil, tmp8018)

					var ifres8006 Obj

					if True == tmp8019 {
						tmp8017 := PrimIsPair(V4487)

						var ifres8008 Obj

						if True == tmp8017 {
							tmp8015 := PrimHead(V4487)

							tmp8016 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8015)

							var ifres8010 Obj

							if True == tmp8016 {
								tmp8012 := PrimHead(V4487)

								tmp8013 := Call(__e, PrimFunc(symhdstr), tmp8012)

								tmp8014 := PrimEqual(MakeString("!"), tmp8013)

								var ifres8011 Obj

								if True == tmp8014 {
									ifres8011 = True

								} else {
									ifres8011 = False

								}

								ifres8010 = ifres8011

							} else {
								ifres8010 = False

							}

							var ifres8009 Obj

							if True == ifres8010 {
								ifres8009 = True

							} else {
								ifres8009 = False

							}

							ifres8008 = ifres8009

						} else {
							ifres8008 = False

						}

						var ifres8007 Obj

						if True == ifres8008 {
							ifres8007 = True

						} else {
							ifres8007 = False

						}

						ifres8006 = ifres8007

					} else {
						ifres8006 = False

					}

					var ifres8005 Obj

					if True == ifres8006 {
						ifres8005 = True

					} else {
						ifres8005 = False

					}

					ifres8004 = ifres8005

				} else {
					ifres8004 = False

				}

				if True == ifres8004 {
					tmp7948 := MakeNative(func(__e *ControlFlow) {
						Read := __e.Get(1)
						_ = Read
						tmp7949 := MakeNative(func(__e *ControlFlow) {
							Match := __e.Get(1)
							_ = Match
							tmp7950 := MakeNative(func(__e *ControlFlow) {
								Print := __e.Get(1)
								_ = Print
								tmp7951 := MakeNative(func(__e *ControlFlow) {
									Y := __e.Get(1)
									_ = Y
									tmp7952 := MakeNative(func(__e *ControlFlow) {
										NewHistory := __e.Get(1)
										_ = NewHistory
										__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), Y, NewHistory, V4488)
										return
									}, 1)

									tmp7953 := PrimTail(V4487)

									tmp7954 := PrimCons(Match, tmp7953)

									tmp7955 := PrimSet(symshen_4_dhistory_d, tmp7954)

									__e.TailApply(tmp7952, tmp7955)
									return

								}, 1)

								tmp7956 := Call(__e, PrimFunc(symread_1from_1string), Match)

								__e.TailApply(tmp7951, tmp7956)
								return

							}, 1)

							tmp7957 := Call(__e, PrimFunc(symshen_4app), Match, MakeString("\n"), symshen_4a)

							tmp7958 := Call(__e, PrimFunc(symstoutput))

							tmp7959 := Call(__e, PrimFunc(sympr), tmp7957, tmp7958)

							__e.TailApply(tmp7950, tmp7959)
							return

						}, 1)

						tmp7960 := PrimHead(V4487)

						tmp7961 := PrimTailString(tmp7960)

						tmp7962 := PrimTail(V4487)

						tmp7963 := Call(__e, PrimFunc(symshen_4use_1history), Read, tmp7961, tmp7962)

						__e.TailApply(tmp7949, tmp7963)
						return

					}, 1)

					tmp7964 := PrimHead(V4487)

					tmp7965 := PrimTailString(tmp7964)

					tmp7966 := Call(__e, PrimFunc(symread_1from_1string), tmp7965)

					tmp7967 := PrimHead(tmp7966)

					__e.TailApply(tmp7948, tmp7967)
					return

				} else {
					tmp8002 := PrimIsPair(V4486)

					var ifres7986 Obj

					if True == tmp8002 {
						tmp8000 := PrimTail(V4486)

						tmp8001 := PrimEqual(Nil, tmp8000)

						var ifres7988 Obj

						if True == tmp8001 {
							tmp7999 := PrimIsPair(V4487)

							var ifres7990 Obj

							if True == tmp7999 {
								tmp7997 := PrimHead(V4487)

								tmp7998 := Call(__e, PrimFunc(symshen_4_7string_2), tmp7997)

								var ifres7992 Obj

								if True == tmp7998 {
									tmp7994 := PrimHead(V4487)

									tmp7995 := Call(__e, PrimFunc(symhdstr), tmp7994)

									tmp7996 := PrimEqual(MakeString("%"), tmp7995)

									var ifres7993 Obj

									if True == tmp7996 {
										ifres7993 = True

									} else {
										ifres7993 = False

									}

									ifres7992 = ifres7993

								} else {
									ifres7992 = False

								}

								var ifres7991 Obj

								if True == ifres7992 {
									ifres7991 = True

								} else {
									ifres7991 = False

								}

								ifres7990 = ifres7991

							} else {
								ifres7990 = False

							}

							var ifres7989 Obj

							if True == ifres7990 {
								ifres7989 = True

							} else {
								ifres7989 = False

							}

							ifres7988 = ifres7989

						} else {
							ifres7988 = False

						}

						var ifres7987 Obj

						if True == ifres7988 {
							ifres7987 = True

						} else {
							ifres7987 = False

						}

						ifres7986 = ifres7987

					} else {
						ifres7986 = False

					}

					if True == ifres7986 {
						tmp7968 := MakeNative(func(__e *ControlFlow) {
							Read := __e.Get(1)
							_ = Read
							tmp7969 := MakeNative(func(__e *ControlFlow) {
								Peek := __e.Get(1)
								_ = Peek
								tmp7970 := MakeNative(func(__e *ControlFlow) {
									NewHistory := __e.Get(1)
									_ = NewHistory
									__e.TailApply(PrimFunc(symabort))
									return
								}, 1)

								tmp7971 := PrimTail(V4487)

								tmp7972 := PrimSet(symshen_4_dhistory_d, tmp7971)

								__e.TailApply(tmp7970, tmp7972)
								return

							}, 1)

							tmp7973 := PrimHead(V4487)

							tmp7974 := PrimTailString(tmp7973)

							tmp7975 := PrimTail(V4487)

							tmp7976 := Call(__e, PrimFunc(symshen_4peek_1history), Read, tmp7974, tmp7975)

							__e.TailApply(tmp7969, tmp7976)
							return

						}, 1)

						tmp7977 := PrimHead(V4487)

						tmp7978 := PrimTailString(tmp7977)

						tmp7979 := Call(__e, PrimFunc(symread_1from_1string), tmp7978)

						tmp7980 := PrimHead(tmp7979)

						__e.TailApply(tmp7968, tmp7980)
						return

					} else {
						tmp7984 := PrimEqual(True, V4488)

						if True == tmp7984 {
							__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V4486)
							return
						} else {
							tmp7982 := PrimEqual(False, V4488)

							if True == tmp7982 {
								__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V4486)
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

	tmp8056 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp7918)

	_ = tmp8056

	tmp8057 := MakeNative(func(__e *ControlFlow) {
		V4489 := __e.Get(1)
		_ = V4489
		V4490 := __e.Get(2)
		_ = V4490
		V4491 := __e.Get(3)
		_ = V4491
		tmp8063 := PrimIsInteger(V4489)

		if True == tmp8063 {
			tmp8058 := PrimNumberAdd(MakeNumber(1), V4489)

			tmp8059 := Call(__e, PrimFunc(symreverse), V4491)

			__e.TailApply(PrimFunc(symnth), tmp8058, tmp8059)
			return

		} else {
			tmp8061 := PrimIsSymbol(V4489)

			if True == tmp8061 {
				__e.TailApply(PrimFunc(symshen_4string_1match), V4490, V4491)
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
				return
			}

		}

	}, 3)

	tmp8064 := Call(__e, ns2_1set, symshen_4use_1history, tmp8057)

	_ = tmp8064

	tmp8065 := MakeNative(func(__e *ControlFlow) {
		V4492 := __e.Get(1)
		_ = V4492
		V4493 := __e.Get(2)
		_ = V4493
		V4494 := __e.Get(3)
		_ = V4494
		tmp8079 := PrimIsInteger(V4492)

		if True == tmp8079 {
			tmp8066 := PrimNumberAdd(MakeNumber(1), V4492)

			tmp8067 := Call(__e, PrimFunc(symreverse), V4494)

			tmp8068 := Call(__e, PrimFunc(symnth), tmp8066, tmp8067)

			tmp8069 := Call(__e, PrimFunc(symshen_4app), tmp8068, MakeString(""), symshen_4a)

			tmp8070 := PrimStringConcat(MakeString("\n"), tmp8069)

			tmp8071 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp8070, tmp8071)
			return

		} else {
			tmp8077 := PrimEqual(V4493, MakeString(""))

			var ifres8074 Obj

			if True == tmp8077 {
				ifres8074 = True

			} else {
				tmp8076 := PrimIsSymbol(V4492)

				var ifres8075 Obj

				if True == tmp8076 {
					ifres8075 = True

				} else {
					ifres8075 = False

				}

				ifres8074 = ifres8075

			}

			if True == ifres8074 {
				tmp8072 := Call(__e, PrimFunc(symreverse), V4494)

				__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V4493, tmp8072)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
				return
			}

		}

	}, 3)

	tmp8080 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8065)

	_ = tmp8080

	tmp8081 := MakeNative(func(__e *ControlFlow) {
		V4504 := __e.Get(1)
		_ = V4504
		V4505 := __e.Get(2)
		_ = V4505
		tmp8092 := PrimEqual(Nil, V4505)

		if True == tmp8092 {
			__e.Return(PrimSimpleError(MakeString("\ninput not found")))
			return
		} else {
			tmp8090 := PrimIsPair(V4505)

			var ifres8086 Obj

			if True == tmp8090 {
				tmp8088 := PrimHead(V4505)

				tmp8089 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V4504, tmp8088)

				var ifres8087 Obj

				if True == tmp8089 {
					ifres8087 = True

				} else {
					ifres8087 = False

				}

				ifres8086 = ifres8087

			} else {
				ifres8086 = False

			}

			if True == ifres8086 {
				__e.Return(PrimHead(V4505))
				return
			} else {
				tmp8084 := PrimIsPair(V4505)

				if True == tmp8084 {
					tmp8082 := PrimTail(V4505)

					__e.TailApply(PrimFunc(symshen_4string_1match), V4504, tmp8082)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
					return
				}

			}

		}

	}, 2)

	tmp8093 := Call(__e, ns2_1set, symshen_4string_1match, tmp8081)

	_ = tmp8093

	tmp8094 := MakeNative(func(__e *ControlFlow) {
		V4513 := __e.Get(1)
		_ = V4513
		V4514 := __e.Get(2)
		_ = V4514
		tmp8131 := PrimEqual(MakeString(""), V4513)

		if True == tmp8131 {
			__e.Return(True)
			return
		} else {
			tmp8129 := Call(__e, PrimFunc(symshen_4_7string_2), V4513)

			var ifres8124 Obj

			if True == tmp8129 {
				tmp8126 := Call(__e, PrimFunc(symhdstr), V4513)

				tmp8127 := PrimStringToNumber(tmp8126)

				tmp8128 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8127)

				var ifres8125 Obj

				if True == tmp8128 {
					ifres8125 = True

				} else {
					ifres8125 = False

				}

				ifres8124 = ifres8125

			} else {
				ifres8124 = False

			}

			if True == ifres8124 {
				tmp8095 := PrimTailString(V4513)

				__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8095, V4514)
				return

			} else {
				tmp8122 := Call(__e, PrimFunc(symshen_4_7string_2), V4514)

				var ifres8117 Obj

				if True == tmp8122 {
					tmp8119 := Call(__e, PrimFunc(symhdstr), V4514)

					tmp8120 := PrimStringToNumber(tmp8119)

					tmp8121 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8120)

					var ifres8118 Obj

					if True == tmp8121 {
						ifres8118 = True

					} else {
						ifres8118 = False

					}

					ifres8117 = ifres8118

				} else {
					ifres8117 = False

				}

				if True == ifres8117 {
					tmp8096 := PrimTailString(V4514)

					__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V4513, tmp8096)
					return

				} else {
					tmp8115 := Call(__e, PrimFunc(symshen_4_7string_2), V4514)

					var ifres8111 Obj

					if True == tmp8115 {
						tmp8113 := Call(__e, PrimFunc(symhdstr), V4514)

						tmp8114 := PrimEqual(MakeString("("), tmp8113)

						var ifres8112 Obj

						if True == tmp8114 {
							ifres8112 = True

						} else {
							ifres8112 = False

						}

						ifres8111 = ifres8112

					} else {
						ifres8111 = False

					}

					if True == ifres8111 {
						tmp8097 := PrimTailString(V4514)

						__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V4513, tmp8097)
						return

					} else {
						tmp8109 := Call(__e, PrimFunc(symshen_4_7string_2), V4513)

						var ifres8101 Obj

						if True == tmp8109 {
							tmp8108 := Call(__e, PrimFunc(symshen_4_7string_2), V4514)

							var ifres8103 Obj

							if True == tmp8108 {
								tmp8105 := Call(__e, PrimFunc(symhdstr), V4513)

								tmp8106 := Call(__e, PrimFunc(symhdstr), V4514)

								tmp8107 := PrimEqual(tmp8105, tmp8106)

								var ifres8104 Obj

								if True == tmp8107 {
									ifres8104 = True

								} else {
									ifres8104 = False

								}

								ifres8103 = ifres8104

							} else {
								ifres8103 = False

							}

							var ifres8102 Obj

							if True == ifres8103 {
								ifres8102 = True

							} else {
								ifres8102 = False

							}

							ifres8101 = ifres8102

						} else {
							ifres8101 = False

						}

						if True == ifres8101 {
							tmp8098 := PrimTailString(V4513)

							tmp8099 := PrimTailString(V4514)

							__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8098, tmp8099)
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

	tmp8132 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8094)

	_ = tmp8132

	tmp8133 := MakeNative(func(__e *ControlFlow) {
		V4525 := __e.Get(1)
		_ = V4525
		V4526 := __e.Get(2)
		_ = V4526
		V4527 := __e.Get(3)
		_ = V4527
		tmp8148 := PrimEqual(Nil, V4527)

		if True == tmp8148 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp8146 := PrimIsPair(V4527)

			if True == tmp8146 {
				tmp8141 := PrimHead(V4527)

				tmp8142 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V4526, tmp8141)

				var ifres8134 Obj

				if True == tmp8142 {
					tmp8135 := PrimHead(V4527)

					tmp8136 := Call(__e, PrimFunc(symshen_4app), tmp8135, MakeString("\n"), symshen_4a)

					tmp8137 := PrimStringConcat(MakeString(". "), tmp8136)

					tmp8138 := Call(__e, PrimFunc(symshen_4app), V4525, tmp8137, symshen_4a)

					tmp8139 := Call(__e, PrimFunc(symstoutput))

					tmp8140 := Call(__e, PrimFunc(sympr), tmp8138, tmp8139)

					ifres8134 = tmp8140

				} else {
					ifres8134 = symshen_4skip

				}

				_ = ifres8134

				tmp8143 := PrimNumberAdd(V4525, MakeNumber(1))

				tmp8144 := PrimTail(V4527)

				__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8143, V4526, tmp8144)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
				return
			}

		}

	}, 3)

	__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8133)
	return

}, 0)
