package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionLauncherMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

	tmp70879 := MakeNative(func(__e *ControlFlow) {
		V4891 := __e.Get(1)
		_ = V4891
		tmp70880 := MakeNative(func(__e *ControlFlow) {
			Contents := __e.Get(1)
			_ = Contents
			tmp70881 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp70882 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp70883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				__e.TailApply(tmp70883, X)
				return

			}, 1)

			__e.TailApply(tmp70881, tmp70882, Contents)
			return

		}, 1)

		tmp70884 := Call(__e, PrimNS1Value(symns2_1value), symread_1file)

		tmp70885 := Call(__e, tmp70884, V4891)

		__e.TailApply(tmp70880, tmp70885)
		return

	}, 1)

	tmp70886 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4quiet_1load, tmp70879)

	_ = tmp70886

	tmp70887 := MakeNative(func(__e *ControlFlow) {
		tmp70888 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp70889 := Call(__e, PrimNS1Value(symns2_1value), symversion)

		tmp70890 := Call(__e, tmp70889)

		tmp70891 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp70892 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp70893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70896 := Call(__e, PrimNS1Value(symns2_1value), symlanguage)

		tmp70897 := Call(__e, tmp70896)

		tmp70898 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70899 := Call(__e, PrimNS1Value(symns2_1value), symport)

		tmp70900 := Call(__e, tmp70899)

		tmp70901 := Call(__e, tmp70898, tmp70900, Nil)

		tmp70902 := Call(__e, tmp70895, tmp70897, tmp70901)

		tmp70903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70904 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70905 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70906 := Call(__e, PrimNS1Value(symns2_1value), symimplementation)

		tmp70907 := Call(__e, tmp70906)

		tmp70908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp70909 := Call(__e, PrimNS1Value(symns2_1value), symrelease)

		tmp70910 := Call(__e, tmp70909)

		tmp70911 := Call(__e, tmp70908, tmp70910, Nil)

		tmp70912 := Call(__e, tmp70905, tmp70907, tmp70911)

		tmp70913 := Call(__e, tmp70904, tmp70912, Nil)

		tmp70914 := Call(__e, tmp70903, symimplementation, tmp70913)

		tmp70915 := Call(__e, tmp70894, tmp70902, tmp70914)

		tmp70916 := Call(__e, tmp70893, symport, tmp70915)

		tmp70917 := Call(__e, tmp70892, tmp70916, MakeString("\n"), symshen_4r)

		tmp70918 := Call(__e, tmp70891, MakeString(" "), tmp70917)

		__e.TailApply(tmp70888, tmp70890, tmp70918, symshen_4a)
		return

	}, 0)

	tmp70919 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4version_1string, tmp70887)

	_ = tmp70919

	tmp70920 := MakeNative(func(__e *ControlFlow) {
		V4893 := __e.Get(1)
		_ = V4893
		tmp70921 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp70922 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp70923 := Call(__e, tmp70922, V4893, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)

		__e.TailApply(tmp70921, MakeString("Usage: "), tmp70923)
		return

	}, 1)

	tmp70924 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4help_1text, tmp70920)

	_ = tmp70924

	tmp70925 := MakeNative(func(__e *ControlFlow) {
		V4895 := __e.Get(1)
		_ = V4895
		tmp70939 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp70940 := Call(__e, tmp70939, Nil, V4895)

		if True == tmp70940 {
			tmp70938 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp70938, symsuccess, Nil)
			return

		} else {
			tmp70936 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp70937 := Call(__e, tmp70936, V4895)

			if True == tmp70937 {
				tmp70929 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

				tmp70930 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp70931 := Call(__e, tmp70930, V4895)

				tmp70932 := Call(__e, tmp70929, tmp70931)

				_ = tmp70932

				tmp70933 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4execute_1all)

				tmp70934 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp70935 := Call(__e, tmp70934, V4895)

				__e.TailApply(tmp70933, tmp70935)
				return

			} else {
				tmp70928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp70928, symshen_4x_4launcher_4execute_1all)
				return

			}

		}

	}, 1)

	tmp70941 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4execute_1all, tmp70925)

	_ = tmp70941

	tmp70942 := MakeNative(func(__e *ControlFlow) {
		V4897 := __e.Get(1)
		_ = V4897
		tmp70943 := Call(__e, PrimNS1Value(symns2_1value), symeval)

		tmp70944 := Call(__e, PrimNS1Value(symns2_1value), symhead)

		tmp70945 := Call(__e, PrimNS1Value(symns2_1value), symread_1from_1string)

		tmp70946 := Call(__e, tmp70945, V4897)

		tmp70947 := Call(__e, tmp70944, tmp70946)

		__e.TailApply(tmp70943, tmp70947)
		return

	}, 1)

	tmp70948 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1string, tmp70942)

	_ = tmp70948

	tmp70949 := MakeNative(func(__e *ControlFlow) {
		V4903 := __e.Get(1)
		_ = V4903
		tmp70963 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp70964 := Call(__e, tmp70963, MakeString("-e"), V4903)

		if True == tmp70964 {
			__e.Return(MakeString("--eval"))
			return
		} else {
			tmp70961 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp70962 := Call(__e, tmp70961, MakeString("-l"), V4903)

			if True == tmp70962 {
				__e.Return(MakeString("--load"))
				return
			} else {
				tmp70959 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp70960 := Call(__e, tmp70959, MakeString("-q"), V4903)

				if True == tmp70960 {
					__e.Return(MakeString("--quiet"))
					return
				} else {
					tmp70957 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp70958 := Call(__e, tmp70957, MakeString("-s"), V4903)

					if True == tmp70958 {
						__e.Return(MakeString("--set"))
						return
					} else {
						tmp70955 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp70956 := Call(__e, tmp70955, MakeString("-r"), V4903)

						if True == tmp70956 {
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

	tmp70965 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1flag_1map, tmp70949)

	_ = tmp70965

	tmp70966 := MakeNative(func(__e *ControlFlow) {
		V4914 := __e.Get(1)
		_ = V4914
		V4915 := __e.Get(2)
		_ = V4915
		tmp71152 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp71153 := Call(__e, tmp71152, Nil, V4914)

		if True == tmp71153 {
			tmp71149 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4execute_1all)

			tmp71150 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			tmp71151 := Call(__e, tmp71150, V4915)

			__e.TailApply(tmp71149, tmp71151)
			return

		} else {
			tmp71147 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp71148 := Call(__e, tmp71147, V4914)

			var ifres71135 Obj

			if True == tmp71148 {
				tmp71143 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp71144 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp71145 := Call(__e, tmp71144, V4914)

				tmp71146 := Call(__e, tmp71143, MakeString("--eval"), tmp71145)

				var ifres71137 Obj

				if True == tmp71146 {
					tmp71139 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp71140 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71141 := Call(__e, tmp71140, V4914)

					tmp71142 := Call(__e, tmp71139, tmp71141)

					var ifres71138 Obj

					if True == tmp71142 {
						ifres71138 = True

					} else {
						ifres71138 = False

					}

					ifres71137 = ifres71138

				} else {
					ifres71137 = False

				}

				var ifres71136 Obj

				if True == ifres71137 {
					ifres71136 = True

				} else {
					ifres71136 = False

				}

				ifres71135 = ifres71136

			} else {
				ifres71135 = False

			}

			if True == ifres71135 {
				tmp71116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

				tmp71117 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp71118 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp71119 := Call(__e, tmp71118, V4914)

				tmp71120 := Call(__e, tmp71117, tmp71119)

				tmp71121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp71122 := MakeNative(func(__e *ControlFlow) {
					tmp71123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					tmp71124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp71125 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1string)

					tmp71126 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp71127 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71128 := Call(__e, tmp71127, V4914)

					tmp71129 := Call(__e, tmp71126, tmp71128)

					tmp71130 := Call(__e, tmp71125, tmp71129)

					tmp71131 := Call(__e, tmp71124, tmp71130, MakeString("\n"), symshen_4a)

					tmp71132 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					tmp71133 := Call(__e, tmp71132)

					__e.TailApply(tmp71123, tmp71131, tmp71133)
					return

				}, 0)

				tmp71134 := Call(__e, tmp71121, tmp71122, V4915)

				__e.TailApply(tmp71116, tmp71120, tmp71134)
				return

			} else {
				tmp71114 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp71115 := Call(__e, tmp71114, V4914)

				var ifres71102 Obj

				if True == tmp71115 {
					tmp71110 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp71111 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp71112 := Call(__e, tmp71111, V4914)

					tmp71113 := Call(__e, tmp71110, MakeString("--load"), tmp71112)

					var ifres71104 Obj

					if True == tmp71113 {
						tmp71106 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp71107 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71108 := Call(__e, tmp71107, V4914)

						tmp71109 := Call(__e, tmp71106, tmp71108)

						var ifres71105 Obj

						if True == tmp71109 {
							ifres71105 = True

						} else {
							ifres71105 = False

						}

						ifres71104 = ifres71105

					} else {
						ifres71104 = False

					}

					var ifres71103 Obj

					if True == ifres71104 {
						ifres71103 = True

					} else {
						ifres71103 = False

					}

					ifres71102 = ifres71103

				} else {
					ifres71102 = False

				}

				if True == ifres71102 {
					tmp71089 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

					tmp71090 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71091 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71092 := Call(__e, tmp71091, V4914)

					tmp71093 := Call(__e, tmp71090, tmp71092)

					tmp71094 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp71095 := MakeNative(func(__e *ControlFlow) {
						tmp71096 := Call(__e, PrimNS1Value(symns2_1value), symload)

						tmp71097 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp71098 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71099 := Call(__e, tmp71098, V4914)

						tmp71100 := Call(__e, tmp71097, tmp71099)

						__e.TailApply(tmp71096, tmp71100)
						return

					}, 0)

					tmp71101 := Call(__e, tmp71094, tmp71095, V4915)

					__e.TailApply(tmp71089, tmp71093, tmp71101)
					return

				} else {
					tmp71087 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp71088 := Call(__e, tmp71087, V4914)

					var ifres71081 Obj

					if True == tmp71088 {
						tmp71083 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp71084 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp71085 := Call(__e, tmp71084, V4914)

						tmp71086 := Call(__e, tmp71083, MakeString("--quiet"), tmp71085)

						var ifres71082 Obj

						if True == tmp71086 {
							ifres71082 = True

						} else {
							ifres71082 = False

						}

						ifres71081 = ifres71082

					} else {
						ifres71081 = False

					}

					if True == ifres71081 {
						tmp71074 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

						tmp71075 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71076 := Call(__e, tmp71075, V4914)

						tmp71077 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp71078 := MakeNative(func(__e *ControlFlow) {
							tmp71079 := Call(__e, PrimNS1Value(symns2_1value), symset)

							__e.TailApply(tmp71079, sym_dhush_d, True)
							return

						}, 0)

						tmp71080 := Call(__e, tmp71077, tmp71078, V4915)

						__e.TailApply(tmp71074, tmp71076, tmp71080)
						return

					} else {
						tmp71072 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp71073 := Call(__e, tmp71072, V4914)

						var ifres71052 Obj

						if True == tmp71073 {
							tmp71068 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp71069 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp71070 := Call(__e, tmp71069, V4914)

							tmp71071 := Call(__e, tmp71068, MakeString("--set"), tmp71070)

							var ifres71054 Obj

							if True == tmp71071 {
								tmp71064 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp71065 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71066 := Call(__e, tmp71065, V4914)

								tmp71067 := Call(__e, tmp71064, tmp71066)

								var ifres71056 Obj

								if True == tmp71067 {
									tmp71058 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp71059 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71060 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71061 := Call(__e, tmp71060, V4914)

									tmp71062 := Call(__e, tmp71059, tmp71061)

									tmp71063 := Call(__e, tmp71058, tmp71062)

									var ifres71057 Obj

									if True == tmp71063 {
										ifres71057 = True

									} else {
										ifres71057 = False

									}

									ifres71056 = ifres71057

								} else {
									ifres71056 = False

								}

								var ifres71055 Obj

								if True == ifres71056 {
									ifres71055 = True

								} else {
									ifres71055 = False

								}

								ifres71054 = ifres71055

							} else {
								ifres71054 = False

							}

							var ifres71053 Obj

							if True == ifres71054 {
								ifres71053 = True

							} else {
								ifres71053 = False

							}

							ifres71052 = ifres71053

						} else {
							ifres71052 = False

						}

						if True == ifres71052 {
							tmp71027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

							tmp71028 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71029 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71031 := Call(__e, tmp71030, V4914)

							tmp71032 := Call(__e, tmp71029, tmp71031)

							tmp71033 := Call(__e, tmp71028, tmp71032)

							tmp71034 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp71035 := MakeNative(func(__e *ControlFlow) {
								tmp71036 := Call(__e, PrimNS1Value(symns2_1value), symset)

								tmp71037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1string)

								tmp71038 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71039 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71040 := Call(__e, tmp71039, V4914)

								tmp71041 := Call(__e, tmp71038, tmp71040)

								tmp71042 := Call(__e, tmp71037, tmp71041)

								tmp71043 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1string)

								tmp71044 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71045 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71046 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71047 := Call(__e, tmp71046, V4914)

								tmp71048 := Call(__e, tmp71045, tmp71047)

								tmp71049 := Call(__e, tmp71044, tmp71048)

								tmp71050 := Call(__e, tmp71043, tmp71049)

								__e.TailApply(tmp71036, tmp71042, tmp71050)
								return

							}, 0)

							tmp71051 := Call(__e, tmp71034, tmp71035, V4915)

							__e.TailApply(tmp71027, tmp71033, tmp71051)
							return

						} else {
							tmp71025 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp71026 := Call(__e, tmp71025, V4914)

							var ifres71019 Obj

							if True == tmp71026 {
								tmp71021 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp71022 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71023 := Call(__e, tmp71022, V4914)

								tmp71024 := Call(__e, tmp71021, MakeString("--repl"), tmp71023)

								var ifres71020 Obj

								if True == tmp71024 {
									ifres71020 = True

								} else {
									ifres71020 = False

								}

								ifres71019 = ifres71020

							} else {
								ifres71019 = False

							}

							if True == ifres71019 {
								tmp71014 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

								tmp71015 := Call(__e, tmp71014, Nil, V4915)

								_ = tmp71015

								tmp71016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp71017 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71018 := Call(__e, tmp71017, V4914)

								__e.TailApply(tmp71016, symlaunch_1repl, tmp71018)
								return

							} else {
								tmp70973 := MakeNative(func(__e *ControlFlow) {
									Freeze := __e.Get(1)
									_ = Freeze
									tmp70998 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp70999 := Call(__e, tmp70998, V4914)

									if True == tmp70999 {
										tmp70976 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp70979 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp70980 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp70981 := Call(__e, tmp70980)

											tmp70982 := Call(__e, tmp70979, Result, tmp70981)

											if True == tmp70982 {
												tmp70978 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

												__e.TailApply(tmp70978, Freeze)
												return

											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp70983 := MakeNative(func(__e *ControlFlow) {
											Long := __e.Get(1)
											_ = Long
											tmp70991 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp70992 := Call(__e, tmp70991, False, Long)

											if True == tmp70992 {
												tmp70990 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp70990)
												return

											} else {
												tmp70985 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

												tmp70986 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp70987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp70988 := Call(__e, tmp70987, V4914)

												tmp70989 := Call(__e, tmp70986, Long, tmp70988)

												__e.TailApply(tmp70985, tmp70989, V4915)
												return

											}

										}, 1)

										tmp70993 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1flag_1map)

										tmp70994 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp70995 := Call(__e, tmp70994, V4914)

										tmp70996 := Call(__e, tmp70993, tmp70995)

										tmp70997 := Call(__e, tmp70983, tmp70996)

										__e.TailApply(tmp70976, tmp70997)
										return

									} else {
										tmp70975 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

										__e.TailApply(tmp70975, Freeze)
										return

									}

								}, 1)

								tmp71000 := MakeNative(func(__e *ControlFlow) {
									tmp71012 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp71013 := Call(__e, tmp71012, V4914)

									if True == tmp71013 {
										tmp71003 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp71004 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp71005 := Call(__e, PrimNS1Value(symns2_1value), symcn)

										tmp71006 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

										tmp71007 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp71008 := Call(__e, tmp71007, V4914)

										tmp71009 := Call(__e, tmp71006, tmp71008, MakeString(""), symshen_4a)

										tmp71010 := Call(__e, tmp71005, MakeString("Invalid eval argument: "), tmp71009)

										tmp71011 := Call(__e, tmp71004, tmp71010, Nil)

										__e.TailApply(tmp71003, symerror, tmp71011)
										return

									} else {
										tmp71002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

										__e.TailApply(tmp71002, symshen_4x_4launcher_4eval_1command_1h)
										return

									}

								}, 0)

								__e.TailApply(tmp70973, tmp71000)
								return

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp71154 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1command_1h, tmp70966)

	_ = tmp71154

	tmp71155 := MakeNative(func(__e *ControlFlow) {
		V4917 := __e.Get(1)
		_ = V4917
		tmp71156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

		__e.TailApply(tmp71156, V4917, Nil)
		return

	}, 1)

	tmp71157 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1command, tmp71155)

	_ = tmp71157

	tmp71158 := MakeNative(func(__e *ControlFlow) {
		V4920 := __e.Get(1)
		_ = V4920
		V4921 := __e.Get(2)
		_ = V4921
		tmp71159 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp71160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp71161 := Call(__e, tmp71160, V4920, V4921)

		tmp71162 := Call(__e, tmp71159, sym_dargv_d, tmp71161)

		_ = tmp71162

		tmp71163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4quiet_1load)

		tmp71164 := Call(__e, tmp71163, V4920)

		_ = tmp71164

		tmp71165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		__e.TailApply(tmp71165, symsuccess, Nil)
		return

	}, 2)

	tmp71166 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4script_1command, tmp71158)

	_ = tmp71166

	tmp71167 := MakeNative(func(__e *ControlFlow) {
		V4923 := __e.Get(1)
		_ = V4923
		tmp71315 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp71316 := Call(__e, tmp71315, V4923)

		var ifres71309 Obj

		if True == tmp71316 {
			tmp71311 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp71312 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp71313 := Call(__e, tmp71312, V4923)

			tmp71314 := Call(__e, tmp71311, Nil, tmp71313)

			var ifres71310 Obj

			if True == tmp71314 {
				ifres71310 = True

			} else {
				ifres71310 = False

			}

			ifres71309 = ifres71310

		} else {
			ifres71309 = False

		}

		if True == ifres71309 {
			tmp71308 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp71308, symlaunch_1repl, Nil)
			return

		} else {
			tmp71306 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp71307 := Call(__e, tmp71306, V4923)

			var ifres71292 Obj

			if True == tmp71307 {
				tmp71302 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp71303 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp71304 := Call(__e, tmp71303, V4923)

				tmp71305 := Call(__e, tmp71302, tmp71304)

				var ifres71294 Obj

				if True == tmp71305 {
					tmp71296 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp71297 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp71298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71299 := Call(__e, tmp71298, V4923)

					tmp71300 := Call(__e, tmp71297, tmp71299)

					tmp71301 := Call(__e, tmp71296, MakeString("--help"), tmp71300)

					var ifres71295 Obj

					if True == tmp71301 {
						ifres71295 = True

					} else {
						ifres71295 = False

					}

					ifres71294 = ifres71295

				} else {
					ifres71294 = False

				}

				var ifres71293 Obj

				if True == ifres71294 {
					ifres71293 = True

				} else {
					ifres71293 = False

				}

				ifres71292 = ifres71293

			} else {
				ifres71292 = False

			}

			if True == ifres71292 {
				tmp71285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp71286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp71287 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4help_1text)

				tmp71288 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp71289 := Call(__e, tmp71288, V4923)

				tmp71290 := Call(__e, tmp71287, tmp71289)

				tmp71291 := Call(__e, tmp71286, tmp71290, Nil)

				__e.TailApply(tmp71285, symshow_1help, tmp71291)
				return

			} else {
				tmp71283 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp71284 := Call(__e, tmp71283, V4923)

				var ifres71269 Obj

				if True == tmp71284 {
					tmp71279 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp71280 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71281 := Call(__e, tmp71280, V4923)

					tmp71282 := Call(__e, tmp71279, tmp71281)

					var ifres71271 Obj

					if True == tmp71282 {
						tmp71273 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp71274 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp71275 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71276 := Call(__e, tmp71275, V4923)

						tmp71277 := Call(__e, tmp71274, tmp71276)

						tmp71278 := Call(__e, tmp71273, MakeString("--version"), tmp71277)

						var ifres71272 Obj

						if True == tmp71278 {
							ifres71272 = True

						} else {
							ifres71272 = False

						}

						ifres71271 = ifres71272

					} else {
						ifres71271 = False

					}

					var ifres71270 Obj

					if True == ifres71271 {
						ifres71270 = True

					} else {
						ifres71270 = False

					}

					ifres71269 = ifres71270

				} else {
					ifres71269 = False

				}

				if True == ifres71269 {
					tmp71264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp71265 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp71266 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4version_1string)

					tmp71267 := Call(__e, tmp71266)

					tmp71268 := Call(__e, tmp71265, tmp71267, Nil)

					__e.TailApply(tmp71264, symsuccess, tmp71268)
					return

				} else {
					tmp71262 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp71263 := Call(__e, tmp71262, V4923)

					var ifres71248 Obj

					if True == tmp71263 {
						tmp71258 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp71259 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71260 := Call(__e, tmp71259, V4923)

						tmp71261 := Call(__e, tmp71258, tmp71260)

						var ifres71250 Obj

						if True == tmp71261 {
							tmp71252 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp71253 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp71254 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71255 := Call(__e, tmp71254, V4923)

							tmp71256 := Call(__e, tmp71253, tmp71255)

							tmp71257 := Call(__e, tmp71252, MakeString("repl"), tmp71256)

							var ifres71251 Obj

							if True == tmp71257 {
								ifres71251 = True

							} else {
								ifres71251 = False

							}

							ifres71250 = ifres71251

						} else {
							ifres71250 = False

						}

						var ifres71249 Obj

						if True == ifres71250 {
							ifres71249 = True

						} else {
							ifres71249 = False

						}

						ifres71248 = ifres71249

					} else {
						ifres71248 = False

					}

					if True == ifres71248 {
						tmp71243 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp71244 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71245 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71246 := Call(__e, tmp71245, V4923)

						tmp71247 := Call(__e, tmp71244, tmp71246)

						__e.TailApply(tmp71243, symlaunch_1repl, tmp71247)
						return

					} else {
						tmp71241 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp71242 := Call(__e, tmp71241, V4923)

						var ifres71219 Obj

						if True == tmp71242 {
							tmp71237 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp71238 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71239 := Call(__e, tmp71238, V4923)

							tmp71240 := Call(__e, tmp71237, tmp71239)

							var ifres71221 Obj

							if True == tmp71240 {
								tmp71231 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp71232 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71233 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71234 := Call(__e, tmp71233, V4923)

								tmp71235 := Call(__e, tmp71232, tmp71234)

								tmp71236 := Call(__e, tmp71231, MakeString("script"), tmp71235)

								var ifres71223 Obj

								if True == tmp71236 {
									tmp71225 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp71226 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71227 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71228 := Call(__e, tmp71227, V4923)

									tmp71229 := Call(__e, tmp71226, tmp71228)

									tmp71230 := Call(__e, tmp71225, tmp71229)

									var ifres71224 Obj

									if True == tmp71230 {
										ifres71224 = True

									} else {
										ifres71224 = False

									}

									ifres71223 = ifres71224

								} else {
									ifres71223 = False

								}

								var ifres71222 Obj

								if True == ifres71223 {
									ifres71222 = True

								} else {
									ifres71222 = False

								}

								ifres71221 = ifres71222

							} else {
								ifres71221 = False

							}

							var ifres71220 Obj

							if True == ifres71221 {
								ifres71220 = True

							} else {
								ifres71220 = False

							}

							ifres71219 = ifres71220

						} else {
							ifres71219 = False

						}

						if True == ifres71219 {
							tmp71206 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4script_1command)

							tmp71207 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp71208 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71209 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71210 := Call(__e, tmp71209, V4923)

							tmp71211 := Call(__e, tmp71208, tmp71210)

							tmp71212 := Call(__e, tmp71207, tmp71211)

							tmp71213 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71214 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71215 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71216 := Call(__e, tmp71215, V4923)

							tmp71217 := Call(__e, tmp71214, tmp71216)

							tmp71218 := Call(__e, tmp71213, tmp71217)

							__e.TailApply(tmp71206, tmp71212, tmp71218)
							return

						} else {
							tmp71204 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp71205 := Call(__e, tmp71204, V4923)

							var ifres71190 Obj

							if True == tmp71205 {
								tmp71200 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp71201 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71202 := Call(__e, tmp71201, V4923)

								tmp71203 := Call(__e, tmp71200, tmp71202)

								var ifres71192 Obj

								if True == tmp71203 {
									tmp71194 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp71195 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp71196 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71197 := Call(__e, tmp71196, V4923)

									tmp71198 := Call(__e, tmp71195, tmp71197)

									tmp71199 := Call(__e, tmp71194, MakeString("eval"), tmp71198)

									var ifres71193 Obj

									if True == tmp71199 {
										ifres71193 = True

									} else {
										ifres71193 = False

									}

									ifres71192 = ifres71193

								} else {
									ifres71192 = False

								}

								var ifres71191 Obj

								if True == ifres71192 {
									ifres71191 = True

								} else {
									ifres71191 = False

								}

								ifres71190 = ifres71191

							} else {
								ifres71190 = False

							}

							if True == ifres71190 {
								tmp71185 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command)

								tmp71186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71188 := Call(__e, tmp71187, V4923)

								tmp71189 := Call(__e, tmp71186, tmp71188)

								__e.TailApply(tmp71185, tmp71189)
								return

							} else {
								tmp71183 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp71184 := Call(__e, tmp71183, V4923)

								var ifres71177 Obj

								if True == tmp71184 {
									tmp71179 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp71180 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71181 := Call(__e, tmp71180, V4923)

									tmp71182 := Call(__e, tmp71179, tmp71181)

									var ifres71178 Obj

									if True == tmp71182 {
										ifres71178 = True

									} else {
										ifres71178 = False

									}

									ifres71177 = ifres71178

								} else {
									ifres71177 = False

								}

								if True == ifres71177 {
									tmp71176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									__e.TailApply(tmp71176, symunknown_1arguments, V4923)
									return

								} else {
									tmp71175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

									__e.TailApply(tmp71175, symshen_4x_4launcher_4launch_1shen)
									return

								}

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp71317 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4launch_1shen, tmp71167)

	_ = tmp71317

	tmp71318 := MakeNative(func(__e *ControlFlow) {
		V4927 := __e.Get(1)
		_ = V4927
		tmp71485 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp71486 := Call(__e, tmp71485, V4927)

		var ifres71473 Obj

		if True == tmp71486 {
			tmp71481 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp71482 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp71483 := Call(__e, tmp71482, V4927)

			tmp71484 := Call(__e, tmp71481, symsuccess, tmp71483)

			var ifres71475 Obj

			if True == tmp71484 {
				tmp71477 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp71478 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp71479 := Call(__e, tmp71478, V4927)

				tmp71480 := Call(__e, tmp71477, Nil, tmp71479)

				var ifres71476 Obj

				if True == tmp71480 {
					ifres71476 = True

				} else {
					ifres71476 = False

				}

				ifres71475 = ifres71476

			} else {
				ifres71475 = False

			}

			var ifres71474 Obj

			if True == ifres71475 {
				ifres71474 = True

			} else {
				ifres71474 = False

			}

			ifres71473 = ifres71474

		} else {
			ifres71473 = False

		}

		if True == ifres71473 {
			__e.Return(symshen_4x_4launcher_4done)
			return
		} else {
			tmp71471 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp71472 := Call(__e, tmp71471, V4927)

			var ifres71451 Obj

			if True == tmp71472 {
				tmp71467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp71468 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp71469 := Call(__e, tmp71468, V4927)

				tmp71470 := Call(__e, tmp71467, symsuccess, tmp71469)

				var ifres71453 Obj

				if True == tmp71470 {
					tmp71463 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp71464 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71465 := Call(__e, tmp71464, V4927)

					tmp71466 := Call(__e, tmp71463, tmp71465)

					var ifres71455 Obj

					if True == tmp71466 {
						tmp71457 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp71458 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71459 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71460 := Call(__e, tmp71459, V4927)

						tmp71461 := Call(__e, tmp71458, tmp71460)

						tmp71462 := Call(__e, tmp71457, Nil, tmp71461)

						var ifres71456 Obj

						if True == tmp71462 {
							ifres71456 = True

						} else {
							ifres71456 = False

						}

						ifres71455 = ifres71456

					} else {
						ifres71455 = False

					}

					var ifres71454 Obj

					if True == ifres71455 {
						ifres71454 = True

					} else {
						ifres71454 = False

					}

					ifres71453 = ifres71454

				} else {
					ifres71453 = False

				}

				var ifres71452 Obj

				if True == ifres71453 {
					ifres71452 = True

				} else {
					ifres71452 = False

				}

				ifres71451 = ifres71452

			} else {
				ifres71451 = False

			}

			if True == ifres71451 {
				tmp71442 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp71443 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp71444 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp71445 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp71446 := Call(__e, tmp71445, V4927)

				tmp71447 := Call(__e, tmp71444, tmp71446)

				tmp71448 := Call(__e, tmp71443, tmp71447, MakeString("\n"), symshen_4a)

				tmp71449 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp71450 := Call(__e, tmp71449)

				__e.TailApply(tmp71442, tmp71448, tmp71450)
				return

			} else {
				tmp71440 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp71441 := Call(__e, tmp71440, V4927)

				var ifres71420 Obj

				if True == tmp71441 {
					tmp71436 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp71437 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp71438 := Call(__e, tmp71437, V4927)

					tmp71439 := Call(__e, tmp71436, symerror, tmp71438)

					var ifres71422 Obj

					if True == tmp71439 {
						tmp71432 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp71433 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp71434 := Call(__e, tmp71433, V4927)

						tmp71435 := Call(__e, tmp71432, tmp71434)

						var ifres71424 Obj

						if True == tmp71435 {
							tmp71426 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp71427 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71428 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71429 := Call(__e, tmp71428, V4927)

							tmp71430 := Call(__e, tmp71427, tmp71429)

							tmp71431 := Call(__e, tmp71426, Nil, tmp71430)

							var ifres71425 Obj

							if True == tmp71431 {
								ifres71425 = True

							} else {
								ifres71425 = False

							}

							ifres71424 = ifres71425

						} else {
							ifres71424 = False

						}

						var ifres71423 Obj

						if True == ifres71424 {
							ifres71423 = True

						} else {
							ifres71423 = False

						}

						ifres71422 = ifres71423

					} else {
						ifres71422 = False

					}

					var ifres71421 Obj

					if True == ifres71422 {
						ifres71421 = True

					} else {
						ifres71421 = False

					}

					ifres71420 = ifres71421

				} else {
					ifres71420 = False

				}

				if True == ifres71420 {
					tmp71409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					tmp71410 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp71411 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp71412 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp71413 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp71414 := Call(__e, tmp71413, V4927)

					tmp71415 := Call(__e, tmp71412, tmp71414)

					tmp71416 := Call(__e, tmp71411, tmp71415, MakeString("\n"), symshen_4a)

					tmp71417 := Call(__e, tmp71410, MakeString("ERROR: "), tmp71416)

					tmp71418 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					tmp71419 := Call(__e, tmp71418)

					__e.TailApply(tmp71409, tmp71417, tmp71419)
					return

				} else {
					tmp71407 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp71408 := Call(__e, tmp71407, V4927)

					var ifres71401 Obj

					if True == tmp71408 {
						tmp71403 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp71404 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp71405 := Call(__e, tmp71404, V4927)

						tmp71406 := Call(__e, tmp71403, symlaunch_1repl, tmp71405)

						var ifres71402 Obj

						if True == tmp71406 {
							ifres71402 = True

						} else {
							ifres71402 = False

						}

						ifres71401 = ifres71402

					} else {
						ifres71401 = False

					}

					if True == ifres71401 {
						tmp71400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4repl)

						__e.TailApply(tmp71400)
						return

					} else {
						tmp71398 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp71399 := Call(__e, tmp71398, V4927)

						var ifres71378 Obj

						if True == tmp71399 {
							tmp71394 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp71395 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp71396 := Call(__e, tmp71395, V4927)

							tmp71397 := Call(__e, tmp71394, symshow_1help, tmp71396)

							var ifres71380 Obj

							if True == tmp71397 {
								tmp71390 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp71391 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71392 := Call(__e, tmp71391, V4927)

								tmp71393 := Call(__e, tmp71390, tmp71392)

								var ifres71382 Obj

								if True == tmp71393 {
									tmp71384 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp71385 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71386 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71387 := Call(__e, tmp71386, V4927)

									tmp71388 := Call(__e, tmp71385, tmp71387)

									tmp71389 := Call(__e, tmp71384, Nil, tmp71388)

									var ifres71383 Obj

									if True == tmp71389 {
										ifres71383 = True

									} else {
										ifres71383 = False

									}

									ifres71382 = ifres71383

								} else {
									ifres71382 = False

								}

								var ifres71381 Obj

								if True == ifres71382 {
									ifres71381 = True

								} else {
									ifres71381 = False

								}

								ifres71380 = ifres71381

							} else {
								ifres71380 = False

							}

							var ifres71379 Obj

							if True == ifres71380 {
								ifres71379 = True

							} else {
								ifres71379 = False

							}

							ifres71378 = ifres71379

						} else {
							ifres71378 = False

						}

						if True == ifres71378 {
							tmp71369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

							tmp71370 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

							tmp71371 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp71372 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp71373 := Call(__e, tmp71372, V4927)

							tmp71374 := Call(__e, tmp71371, tmp71373)

							tmp71375 := Call(__e, tmp71370, tmp71374, MakeString("\n"), symshen_4a)

							tmp71376 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

							tmp71377 := Call(__e, tmp71376)

							__e.TailApply(tmp71369, tmp71375, tmp71377)
							return

						} else {
							tmp71367 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp71368 := Call(__e, tmp71367, V4927)

							var ifres71347 Obj

							if True == tmp71368 {
								tmp71363 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp71364 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71365 := Call(__e, tmp71364, V4927)

								tmp71366 := Call(__e, tmp71363, symunknown_1arguments, tmp71365)

								var ifres71349 Obj

								if True == tmp71366 {
									tmp71359 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp71360 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp71361 := Call(__e, tmp71360, V4927)

									tmp71362 := Call(__e, tmp71359, tmp71361)

									var ifres71351 Obj

									if True == tmp71362 {
										tmp71353 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp71354 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp71355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp71356 := Call(__e, tmp71355, V4927)

										tmp71357 := Call(__e, tmp71354, tmp71356)

										tmp71358 := Call(__e, tmp71353, tmp71357)

										var ifres71352 Obj

										if True == tmp71358 {
											ifres71352 = True

										} else {
											ifres71352 = False

										}

										ifres71351 = ifres71352

									} else {
										ifres71351 = False

									}

									var ifres71350 Obj

									if True == ifres71351 {
										ifres71350 = True

									} else {
										ifres71350 = False

									}

									ifres71349 = ifres71350

								} else {
									ifres71349 = False

								}

								var ifres71348 Obj

								if True == ifres71349 {
									ifres71348 = True

								} else {
									ifres71348 = False

								}

								ifres71347 = ifres71348

							} else {
								ifres71347 = False

							}

							if True == ifres71347 {
								tmp71326 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

								tmp71327 := Call(__e, PrimNS1Value(symns2_1value), symcn)

								tmp71328 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

								tmp71329 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71330 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71331 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71332 := Call(__e, tmp71331, V4927)

								tmp71333 := Call(__e, tmp71330, tmp71332)

								tmp71334 := Call(__e, tmp71329, tmp71333)

								tmp71335 := Call(__e, PrimNS1Value(symns2_1value), symcn)

								tmp71336 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

								tmp71337 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp71338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp71339 := Call(__e, tmp71338, V4927)

								tmp71340 := Call(__e, tmp71337, tmp71339)

								tmp71341 := Call(__e, tmp71336, tmp71340, MakeString(" --help' for more information.\n"), symshen_4a)

								tmp71342 := Call(__e, tmp71335, MakeString("\nTry `"), tmp71341)

								tmp71343 := Call(__e, tmp71328, tmp71334, tmp71342, symshen_4a)

								tmp71344 := Call(__e, tmp71327, MakeString("ERROR: Invalid argument: "), tmp71343)

								tmp71345 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

								tmp71346 := Call(__e, tmp71345)

								__e.TailApply(tmp71326, tmp71344, tmp71346)
								return

							} else {
								tmp71325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

								__e.TailApply(tmp71325, symshen_4x_4launcher_4default_1handle_1result)
								return

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp71487 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4default_1handle_1result, tmp71318)

	_ = tmp71487

	tmp71488 := MakeNative(func(__e *ControlFlow) {
		V4929 := __e.Get(1)
		_ = V4929
		tmp71489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4default_1handle_1result)

		tmp71490 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4launch_1shen)

		tmp71491 := Call(__e, tmp71490, V4929)

		__e.TailApply(tmp71489, tmp71491)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4launcher_4main, tmp71488)
	return

}, 0)
