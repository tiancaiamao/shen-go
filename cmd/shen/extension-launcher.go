package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionLauncherMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")
	gen36886 := MakeNative(func(__e *ControlFlow) {
		V4891 := __e.Get(1)
		_ = V4891
		gen36883 := MakeNative(func(__e *ControlFlow) {
			Contents := __e.Get(1)
			_ = Contents
			gen36880 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			gen36882 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen36881 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				__e.TailApply(gen36881, X)

				return

			}, 1)

			__e.TailApply(gen36880, gen36882, Contents)

			return

		}, 1)

		gen36884 := Call(__e, PrimNS1Value(symns2_1value), symread_1file)

		gen36885 := Call(__e, gen36884, V4891)

		__e.TailApply(gen36883, gen36885)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4quiet_1load, gen36886)

	gen36918 := MakeNative(func(__e *ControlFlow) {
		gen36887 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen36888 := Call(__e, PrimNS1Value(symns2_1value), symversion)

		gen36889 := Call(__e, gen36888)

		gen36890 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen36891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen36892 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36895 := Call(__e, PrimNS1Value(symns2_1value), symlanguage)

		gen36896 := Call(__e, gen36895)

		gen36897 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36898 := Call(__e, PrimNS1Value(symns2_1value), symport)

		gen36899 := Call(__e, gen36898)

		gen36900 := Call(__e, gen36897, gen36899, Nil)

		gen36901 := Call(__e, gen36894, gen36896, gen36900)

		gen36902 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36904 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36905 := Call(__e, PrimNS1Value(symns2_1value), symimplementation)

		gen36906 := Call(__e, gen36905)

		gen36907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen36908 := Call(__e, PrimNS1Value(symns2_1value), symrelease)

		gen36909 := Call(__e, gen36908)

		gen36910 := Call(__e, gen36907, gen36909, Nil)

		gen36911 := Call(__e, gen36904, gen36906, gen36910)

		gen36912 := Call(__e, gen36903, gen36911, Nil)

		gen36913 := Call(__e, gen36902, symimplementation, gen36912)

		gen36914 := Call(__e, gen36893, gen36901, gen36913)

		gen36915 := Call(__e, gen36892, symport, gen36914)

		gen36916 := Call(__e, gen36891, gen36915, MakeString("\n"), symshen_4r)

		gen36917 := Call(__e, gen36890, MakeString(" "), gen36916)

		__e.TailApply(gen36887, gen36889, gen36917, symshen_4a)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4version_1string, gen36918)

	gen36922 := MakeNative(func(__e *ControlFlow) {
		V4893 := __e.Get(1)
		_ = V4893
		gen36919 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen36920 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen36921 := Call(__e, gen36920, V4893, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)

		__e.TailApply(gen36919, MakeString("Usage: "), gen36921)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4help_1text, gen36922)

	gen36936 := MakeNative(func(__e *ControlFlow) {
		V4895 := __e.Get(1)
		_ = V4895
		gen36934 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen36935 := Call(__e, gen36934, Nil, V4895)

		if True == gen36935 {
			gen36933 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(gen36933, symsuccess, Nil)

			return

		} else {
			gen36931 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen36932 := Call(__e, gen36931, V4895)

			if True == gen36932 {
				gen36925 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

				gen36926 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen36927 := Call(__e, gen36926, V4895)

				Call(__e, gen36925, gen36927)

				gen36928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4execute_1all)

				gen36929 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen36930 := Call(__e, gen36929, V4895)

				__e.TailApply(gen36928, gen36930)

				return

			} else {
				if True == True {
					gen36924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen36924, symshen_4x_4launcher_4execute_1all)

					return

				} else {
					gen36923 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen36923, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4execute_1all, gen36936)

	gen36942 := MakeNative(func(__e *ControlFlow) {
		V4897 := __e.Get(1)
		_ = V4897
		gen36937 := Call(__e, PrimNS1Value(symns2_1value), symeval)

		gen36938 := Call(__e, PrimNS1Value(symns2_1value), symhead)

		gen36939 := Call(__e, PrimNS1Value(symns2_1value), symread_1from_1string)

		gen36940 := Call(__e, gen36939, V4897)

		gen36941 := Call(__e, gen36938, gen36940)

		__e.TailApply(gen36937, gen36941)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1string, gen36942)

	gen36954 := MakeNative(func(__e *ControlFlow) {
		V4903 := __e.Get(1)
		_ = V4903
		gen36952 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen36953 := Call(__e, gen36952, MakeString("-e"), V4903)

		if True == gen36953 {
			__e.Return(MakeString("--eval"))
			return
		} else {
			gen36950 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen36951 := Call(__e, gen36950, MakeString("-l"), V4903)

			if True == gen36951 {
				__e.Return(MakeString("--load"))
				return
			} else {
				gen36948 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen36949 := Call(__e, gen36948, MakeString("-q"), V4903)

				if True == gen36949 {
					__e.Return(MakeString("--quiet"))
					return
				} else {
					gen36946 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen36947 := Call(__e, gen36946, MakeString("-s"), V4903)

					if True == gen36947 {
						__e.Return(MakeString("--set"))
						return
					} else {
						gen36944 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen36945 := Call(__e, gen36944, MakeString("-r"), V4903)

						if True == gen36945 {
							__e.Return(MakeString("--repl"))
							return
						} else {
							if True == True {
								__e.Return(False)
								return
							} else {
								gen36943 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

								__e.TailApply(gen36943, MakeString("no cond match"))

								return

							}
						}

					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1flag_1map, gen36954)

	gen37124 := MakeNative(func(__e *ControlFlow) {
		V4914 := __e.Get(1)
		_ = V4914
		V4915 := __e.Get(2)
		_ = V4915
		gen37122 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen37123 := Call(__e, gen37122, Nil, V4914)

		if True == gen37123 {
			gen37119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4execute_1all)

			gen37120 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			gen37121 := Call(__e, gen37120, V4915)

			__e.TailApply(gen37119, gen37121)

			return

		} else {
			gen37116 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen37117 := Call(__e, gen37116, V4914)

			var gen37118 Obj

			if True == gen37117 {
				gen37111 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen37112 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37113 := Call(__e, gen37112, V4914)

				gen37114 := Call(__e, gen37111, MakeString("--eval"), gen37113)

				var gen37115 Obj

				if True == gen37114 {
					gen37107 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37108 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37109 := Call(__e, gen37108, V4914)

					gen37110 := Call(__e, gen37107, gen37109)

					if True == gen37110 {
						gen37115 = True
					} else {
						gen37115 = False
					}

				} else {
					gen37115 = False
				}

				if True == gen37115 {
					gen37118 = True
				} else {
					gen37118 = False
				}

			} else {
				gen37118 = False
			}

			if True == gen37118 {
				gen37088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

				gen37089 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37090 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37091 := Call(__e, gen37090, V4914)

				gen37092 := Call(__e, gen37089, gen37091)

				gen37093 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37105 := MakeNative(func(__e *ControlFlow) {
					gen37094 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					gen37095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen37096 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1string)

					gen37097 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37098 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37099 := Call(__e, gen37098, V4914)

					gen37100 := Call(__e, gen37097, gen37099)

					gen37101 := Call(__e, gen37096, gen37100)

					gen37102 := Call(__e, gen37095, gen37101, MakeString("\n"), symshen_4a)

					gen37103 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					gen37104 := Call(__e, gen37103)

					__e.TailApply(gen37094, gen37102, gen37104)

					return

				}, 0)

				gen37106 := Call(__e, gen37093, gen37105, V4915)

				__e.TailApply(gen37088, gen37092, gen37106)

				return

			} else {
				gen37085 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37086 := Call(__e, gen37085, V4914)

				var gen37087 Obj

				if True == gen37086 {
					gen37080 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen37081 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37082 := Call(__e, gen37081, V4914)

					gen37083 := Call(__e, gen37080, MakeString("--load"), gen37082)

					var gen37084 Obj

					if True == gen37083 {
						gen37076 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37077 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37078 := Call(__e, gen37077, V4914)

						gen37079 := Call(__e, gen37076, gen37078)

						if True == gen37079 {
							gen37084 = True
						} else {
							gen37084 = False
						}

					} else {
						gen37084 = False
					}

					if True == gen37084 {
						gen37087 = True
					} else {
						gen37087 = False
					}

				} else {
					gen37087 = False
				}

				if True == gen37087 {
					gen37063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

					gen37064 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37065 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37066 := Call(__e, gen37065, V4914)

					gen37067 := Call(__e, gen37064, gen37066)

					gen37068 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37074 := MakeNative(func(__e *ControlFlow) {
						gen37069 := Call(__e, PrimNS1Value(symns2_1value), symload)

						gen37070 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37072 := Call(__e, gen37071, V4914)

						gen37073 := Call(__e, gen37070, gen37072)

						__e.TailApply(gen37069, gen37073)

						return

					}, 0)

					gen37075 := Call(__e, gen37068, gen37074, V4915)

					__e.TailApply(gen37063, gen37067, gen37075)

					return

				} else {
					gen37060 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37061 := Call(__e, gen37060, V4914)

					var gen37062 Obj

					if True == gen37061 {
						gen37056 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen37057 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37058 := Call(__e, gen37057, V4914)

						gen37059 := Call(__e, gen37056, MakeString("--quiet"), gen37058)

						if True == gen37059 {
							gen37062 = True
						} else {
							gen37062 = False
						}

					} else {
						gen37062 = False
					}

					if True == gen37062 {
						gen37049 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

						gen37050 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37051 := Call(__e, gen37050, V4914)

						gen37052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37054 := MakeNative(func(__e *ControlFlow) {
							gen37053 := Call(__e, PrimNS1Value(symns2_1value), symset)

							__e.TailApply(gen37053, sym_dhush_d, True)

							return

						}, 0)

						gen37055 := Call(__e, gen37052, gen37054, V4915)

						__e.TailApply(gen37049, gen37051, gen37055)

						return

					} else {
						gen37046 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37047 := Call(__e, gen37046, V4914)

						var gen37048 Obj

						if True == gen37047 {
							gen37041 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37042 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen37043 := Call(__e, gen37042, V4914)

							gen37044 := Call(__e, gen37041, MakeString("--set"), gen37043)

							var gen37045 Obj

							if True == gen37044 {
								gen37036 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen37037 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37038 := Call(__e, gen37037, V4914)

								gen37039 := Call(__e, gen37036, gen37038)

								var gen37040 Obj

								if True == gen37039 {
									gen37030 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen37031 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37032 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37033 := Call(__e, gen37032, V4914)

									gen37034 := Call(__e, gen37031, gen37033)

									gen37035 := Call(__e, gen37030, gen37034)

									if True == gen37035 {
										gen37040 = True
									} else {
										gen37040 = False
									}

								} else {
									gen37040 = False
								}

								if True == gen37040 {
									gen37045 = True
								} else {
									gen37045 = False
								}

							} else {
								gen37045 = False
							}

							if True == gen37045 {
								gen37048 = True
							} else {
								gen37048 = False
							}

						} else {
							gen37048 = False
						}

						if True == gen37048 {
							gen37005 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

							gen37006 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37007 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37008 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37009 := Call(__e, gen37008, V4914)

							gen37010 := Call(__e, gen37007, gen37009)

							gen37011 := Call(__e, gen37006, gen37010)

							gen37012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen37028 := MakeNative(func(__e *ControlFlow) {
								gen37013 := Call(__e, PrimNS1Value(symns2_1value), symset)

								gen37014 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1string)

								gen37015 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37016 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37017 := Call(__e, gen37016, V4914)

								gen37018 := Call(__e, gen37015, gen37017)

								gen37019 := Call(__e, gen37014, gen37018)

								gen37020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1string)

								gen37021 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37022 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37023 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37024 := Call(__e, gen37023, V4914)

								gen37025 := Call(__e, gen37022, gen37024)

								gen37026 := Call(__e, gen37021, gen37025)

								gen37027 := Call(__e, gen37020, gen37026)

								__e.TailApply(gen37013, gen37019, gen37027)

								return

							}, 0)

							gen37029 := Call(__e, gen37012, gen37028, V4915)

							__e.TailApply(gen37005, gen37011, gen37029)

							return

						} else {
							gen37002 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37003 := Call(__e, gen37002, V4914)

							var gen37004 Obj

							if True == gen37003 {
								gen36998 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen36999 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37000 := Call(__e, gen36999, V4914)

								gen37001 := Call(__e, gen36998, MakeString("--repl"), gen37000)

								if True == gen37001 {
									gen37004 = True
								} else {
									gen37004 = False
								}

							} else {
								gen37004 = False
							}

							if True == gen37004 {
								gen36994 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

								Call(__e, gen36994, Nil, V4915)

								gen36995 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen36996 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen36997 := Call(__e, gen36996, V4914)

								__e.TailApply(gen36995, symlaunch_1repl, gen36997)

								return

							} else {
								if True == True {
									gen36979 := MakeNative(func(__e *ControlFlow) {
										Freeze := __e.Get(1)
										_ = Freeze
										gen36977 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36978 := Call(__e, gen36977, V4914)

										if True == gen36978 {
											gen36962 := MakeNative(func(__e *ControlFlow) {
												Result := __e.Get(1)
												_ = Result
												gen36958 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen36959 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												gen36960 := Call(__e, gen36959)

												gen36961 := Call(__e, gen36958, Result, gen36960)

												if True == gen36961 {
													gen36957 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

													__e.TailApply(gen36957, Freeze)

													return

												} else {
													__e.Return(Result)
													return
												}

											}, 1)

											gen36971 := MakeNative(func(__e *ControlFlow) {
												Long := __e.Get(1)
												_ = Long
												gen36969 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen36970 := Call(__e, gen36969, False, Long)

												if True == gen36970 {
													gen36968 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													__e.TailApply(gen36968)

													return

												} else {
													gen36963 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

													gen36964 := Call(__e, PrimNS1Value(symns2_1value), symcons)

													gen36965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen36966 := Call(__e, gen36965, V4914)

													gen36967 := Call(__e, gen36964, Long, gen36966)

													__e.TailApply(gen36963, gen36967, V4915)

													return

												}

											}, 1)

											gen36972 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1flag_1map)

											gen36973 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36974 := Call(__e, gen36973, V4914)

											gen36975 := Call(__e, gen36972, gen36974)

											gen36976 := Call(__e, gen36971, gen36975)

											__e.TailApply(gen36962, gen36976)

											return

										} else {
											gen36956 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

											__e.TailApply(gen36956, Freeze)

											return

										}

									}, 1)

									gen36993 := MakeNative(func(__e *ControlFlow) {
										gen36991 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen36992 := Call(__e, gen36991, V4914)

										if True == gen36992 {
											gen36982 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											gen36983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											gen36984 := Call(__e, PrimNS1Value(symns2_1value), symcn)

											gen36985 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

											gen36986 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen36987 := Call(__e, gen36986, V4914)

											gen36988 := Call(__e, gen36985, gen36987, MakeString(""), symshen_4a)

											gen36989 := Call(__e, gen36984, MakeString("Invalid eval argument: "), gen36988)

											gen36990 := Call(__e, gen36983, gen36989, Nil)

											__e.TailApply(gen36982, symerror, gen36990)

											return

										} else {
											if True == True {
												gen36981 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

												__e.TailApply(gen36981, symshen_4x_4launcher_4eval_1command_1h)

												return

											} else {
												gen36980 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

												__e.TailApply(gen36980, MakeString("no cond match"))

												return

											}
										}

									}, 0)

									__e.TailApply(gen36979, gen36993)

									return

								} else {
									gen36955 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

									__e.TailApply(gen36955, MakeString("no cond match"))

									return

								}
							}

						}

					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1command_1h, gen37124)

	gen37126 := MakeNative(func(__e *ControlFlow) {
		V4917 := __e.Get(1)
		_ = V4917
		gen37125 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command_1h)

		__e.TailApply(gen37125, V4917, Nil)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4eval_1command, gen37126)

	gen37132 := MakeNative(func(__e *ControlFlow) {
		V4920 := __e.Get(1)
		_ = V4920
		V4921 := __e.Get(2)
		_ = V4921
		gen37127 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen37128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen37129 := Call(__e, gen37128, V4920, V4921)

		Call(__e, gen37127, sym_dargv_d, gen37129)

		gen37130 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4quiet_1load)

		Call(__e, gen37130, V4920)

		gen37131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		__e.TailApply(gen37131, symsuccess, Nil)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4script_1command, gen37132)

	gen37263 := MakeNative(func(__e *ControlFlow) {
		V4923 := __e.Get(1)
		_ = V4923
		gen37260 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen37261 := Call(__e, gen37260, V4923)

		var gen37262 Obj

		if True == gen37261 {
			gen37256 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen37257 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen37258 := Call(__e, gen37257, V4923)

			gen37259 := Call(__e, gen37256, Nil, gen37258)

			if True == gen37259 {
				gen37262 = True
			} else {
				gen37262 = False
			}

		} else {
			gen37262 = False
		}

		if True == gen37262 {
			gen37255 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(gen37255, symlaunch_1repl, Nil)

			return

		} else {
			gen37252 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen37253 := Call(__e, gen37252, V4923)

			var gen37254 Obj

			if True == gen37253 {
				gen37247 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37248 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37249 := Call(__e, gen37248, V4923)

				gen37250 := Call(__e, gen37247, gen37249)

				var gen37251 Obj

				if True == gen37250 {
					gen37241 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen37242 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37243 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37244 := Call(__e, gen37243, V4923)

					gen37245 := Call(__e, gen37242, gen37244)

					gen37246 := Call(__e, gen37241, MakeString("--help"), gen37245)

					if True == gen37246 {
						gen37251 = True
					} else {
						gen37251 = False
					}

				} else {
					gen37251 = False
				}

				if True == gen37251 {
					gen37254 = True
				} else {
					gen37254 = False
				}

			} else {
				gen37254 = False
			}

			if True == gen37254 {
				gen37234 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37235 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen37236 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4help_1text)

				gen37237 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37238 := Call(__e, gen37237, V4923)

				gen37239 := Call(__e, gen37236, gen37238)

				gen37240 := Call(__e, gen37235, gen37239, Nil)

				__e.TailApply(gen37234, symshow_1help, gen37240)

				return

			} else {
				gen37231 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37232 := Call(__e, gen37231, V4923)

				var gen37233 Obj

				if True == gen37232 {
					gen37226 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37227 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37228 := Call(__e, gen37227, V4923)

					gen37229 := Call(__e, gen37226, gen37228)

					var gen37230 Obj

					if True == gen37229 {
						gen37220 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen37221 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37222 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37223 := Call(__e, gen37222, V4923)

						gen37224 := Call(__e, gen37221, gen37223)

						gen37225 := Call(__e, gen37220, MakeString("--version"), gen37224)

						if True == gen37225 {
							gen37230 = True
						} else {
							gen37230 = False
						}

					} else {
						gen37230 = False
					}

					if True == gen37230 {
						gen37233 = True
					} else {
						gen37233 = False
					}

				} else {
					gen37233 = False
				}

				if True == gen37233 {
					gen37215 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37216 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen37217 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4version_1string)

					gen37218 := Call(__e, gen37217)

					gen37219 := Call(__e, gen37216, gen37218, Nil)

					__e.TailApply(gen37215, symsuccess, gen37219)

					return

				} else {
					gen37212 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37213 := Call(__e, gen37212, V4923)

					var gen37214 Obj

					if True == gen37213 {
						gen37207 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37208 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37209 := Call(__e, gen37208, V4923)

						gen37210 := Call(__e, gen37207, gen37209)

						var gen37211 Obj

						if True == gen37210 {
							gen37201 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37202 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen37203 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37204 := Call(__e, gen37203, V4923)

							gen37205 := Call(__e, gen37202, gen37204)

							gen37206 := Call(__e, gen37201, MakeString("repl"), gen37205)

							if True == gen37206 {
								gen37211 = True
							} else {
								gen37211 = False
							}

						} else {
							gen37211 = False
						}

						if True == gen37211 {
							gen37214 = True
						} else {
							gen37214 = False
						}

					} else {
						gen37214 = False
					}

					if True == gen37214 {
						gen37196 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen37197 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37198 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37199 := Call(__e, gen37198, V4923)

						gen37200 := Call(__e, gen37197, gen37199)

						__e.TailApply(gen37196, symlaunch_1repl, gen37200)

						return

					} else {
						gen37193 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37194 := Call(__e, gen37193, V4923)

						var gen37195 Obj

						if True == gen37194 {
							gen37188 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37189 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37190 := Call(__e, gen37189, V4923)

							gen37191 := Call(__e, gen37188, gen37190)

							var gen37192 Obj

							if True == gen37191 {
								gen37181 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen37182 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37183 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37184 := Call(__e, gen37183, V4923)

								gen37185 := Call(__e, gen37182, gen37184)

								gen37186 := Call(__e, gen37181, MakeString("script"), gen37185)

								var gen37187 Obj

								if True == gen37186 {
									gen37175 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen37176 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37177 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37178 := Call(__e, gen37177, V4923)

									gen37179 := Call(__e, gen37176, gen37178)

									gen37180 := Call(__e, gen37175, gen37179)

									if True == gen37180 {
										gen37187 = True
									} else {
										gen37187 = False
									}

								} else {
									gen37187 = False
								}

								if True == gen37187 {
									gen37192 = True
								} else {
									gen37192 = False
								}

							} else {
								gen37192 = False
							}

							if True == gen37192 {
								gen37195 = True
							} else {
								gen37195 = False
							}

						} else {
							gen37195 = False
						}

						if True == gen37195 {
							gen37162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4script_1command)

							gen37163 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen37164 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37165 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37166 := Call(__e, gen37165, V4923)

							gen37167 := Call(__e, gen37164, gen37166)

							gen37168 := Call(__e, gen37163, gen37167)

							gen37169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37172 := Call(__e, gen37171, V4923)

							gen37173 := Call(__e, gen37170, gen37172)

							gen37174 := Call(__e, gen37169, gen37173)

							__e.TailApply(gen37162, gen37168, gen37174)

							return

						} else {
							gen37159 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37160 := Call(__e, gen37159, V4923)

							var gen37161 Obj

							if True == gen37160 {
								gen37154 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen37155 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37156 := Call(__e, gen37155, V4923)

								gen37157 := Call(__e, gen37154, gen37156)

								var gen37158 Obj

								if True == gen37157 {
									gen37148 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen37149 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen37150 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37151 := Call(__e, gen37150, V4923)

									gen37152 := Call(__e, gen37149, gen37151)

									gen37153 := Call(__e, gen37148, MakeString("eval"), gen37152)

									if True == gen37153 {
										gen37158 = True
									} else {
										gen37158 = False
									}

								} else {
									gen37158 = False
								}

								if True == gen37158 {
									gen37161 = True
								} else {
									gen37161 = False
								}

							} else {
								gen37161 = False
							}

							if True == gen37161 {
								gen37143 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4eval_1command)

								gen37144 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37145 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37146 := Call(__e, gen37145, V4923)

								gen37147 := Call(__e, gen37144, gen37146)

								__e.TailApply(gen37143, gen37147)

								return

							} else {
								gen37140 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen37141 := Call(__e, gen37140, V4923)

								var gen37142 Obj

								if True == gen37141 {
									gen37136 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen37137 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37138 := Call(__e, gen37137, V4923)

									gen37139 := Call(__e, gen37136, gen37138)

									if True == gen37139 {
										gen37142 = True
									} else {
										gen37142 = False
									}

								} else {
									gen37142 = False
								}

								if True == gen37142 {
									gen37135 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									__e.TailApply(gen37135, symunknown_1arguments, V4923)

									return

								} else {
									if True == True {
										gen37134 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

										__e.TailApply(gen37134, symshen_4x_4launcher_4launch_1shen)

										return

									} else {
										gen37133 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

										__e.TailApply(gen37133, MakeString("no cond match"))

										return

									}
								}

							}

						}

					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4launch_1shen, gen37263)

	gen37412 := MakeNative(func(__e *ControlFlow) {
		V4927 := __e.Get(1)
		_ = V4927
		gen37409 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen37410 := Call(__e, gen37409, V4927)

		var gen37411 Obj

		if True == gen37410 {
			gen37404 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen37405 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen37406 := Call(__e, gen37405, V4927)

			gen37407 := Call(__e, gen37404, symsuccess, gen37406)

			var gen37408 Obj

			if True == gen37407 {
				gen37400 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen37401 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37402 := Call(__e, gen37401, V4927)

				gen37403 := Call(__e, gen37400, Nil, gen37402)

				if True == gen37403 {
					gen37408 = True
				} else {
					gen37408 = False
				}

			} else {
				gen37408 = False
			}

			if True == gen37408 {
				gen37411 = True
			} else {
				gen37411 = False
			}

		} else {
			gen37411 = False
		}

		if True == gen37411 {
			__e.Return(symshen_4x_4launcher_4done)
			return
		} else {
			gen37397 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen37398 := Call(__e, gen37397, V4927)

			var gen37399 Obj

			if True == gen37398 {
				gen37392 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen37393 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37394 := Call(__e, gen37393, V4927)

				gen37395 := Call(__e, gen37392, symsuccess, gen37394)

				var gen37396 Obj

				if True == gen37395 {
					gen37387 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37388 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37389 := Call(__e, gen37388, V4927)

					gen37390 := Call(__e, gen37387, gen37389)

					var gen37391 Obj

					if True == gen37390 {
						gen37381 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen37382 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37383 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37384 := Call(__e, gen37383, V4927)

						gen37385 := Call(__e, gen37382, gen37384)

						gen37386 := Call(__e, gen37381, Nil, gen37385)

						if True == gen37386 {
							gen37391 = True
						} else {
							gen37391 = False
						}

					} else {
						gen37391 = False
					}

					if True == gen37391 {
						gen37396 = True
					} else {
						gen37396 = False
					}

				} else {
					gen37396 = False
				}

				if True == gen37396 {
					gen37399 = True
				} else {
					gen37399 = False
				}

			} else {
				gen37399 = False
			}

			if True == gen37399 {
				gen37372 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				gen37373 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen37374 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen37375 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen37376 := Call(__e, gen37375, V4927)

				gen37377 := Call(__e, gen37374, gen37376)

				gen37378 := Call(__e, gen37373, gen37377, MakeString("\n"), symshen_4a)

				gen37379 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				gen37380 := Call(__e, gen37379)

				__e.TailApply(gen37372, gen37378, gen37380)

				return

			} else {
				gen37369 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen37370 := Call(__e, gen37369, V4927)

				var gen37371 Obj

				if True == gen37370 {
					gen37364 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen37365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37366 := Call(__e, gen37365, V4927)

					gen37367 := Call(__e, gen37364, symerror, gen37366)

					var gen37368 Obj

					if True == gen37367 {
						gen37359 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37360 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen37361 := Call(__e, gen37360, V4927)

						gen37362 := Call(__e, gen37359, gen37361)

						var gen37363 Obj

						if True == gen37362 {
							gen37353 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37354 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37356 := Call(__e, gen37355, V4927)

							gen37357 := Call(__e, gen37354, gen37356)

							gen37358 := Call(__e, gen37353, Nil, gen37357)

							if True == gen37358 {
								gen37363 = True
							} else {
								gen37363 = False
							}

						} else {
							gen37363 = False
						}

						if True == gen37363 {
							gen37368 = True
						} else {
							gen37368 = False
						}

					} else {
						gen37368 = False
					}

					if True == gen37368 {
						gen37371 = True
					} else {
						gen37371 = False
					}

				} else {
					gen37371 = False
				}

				if True == gen37371 {
					gen37342 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					gen37343 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen37344 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen37345 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen37346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen37347 := Call(__e, gen37346, V4927)

					gen37348 := Call(__e, gen37345, gen37347)

					gen37349 := Call(__e, gen37344, gen37348, MakeString("\n"), symshen_4a)

					gen37350 := Call(__e, gen37343, MakeString("ERROR: "), gen37349)

					gen37351 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					gen37352 := Call(__e, gen37351)

					__e.TailApply(gen37342, gen37350, gen37352)

					return

				} else {
					gen37339 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen37340 := Call(__e, gen37339, V4927)

					var gen37341 Obj

					if True == gen37340 {
						gen37335 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen37336 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen37337 := Call(__e, gen37336, V4927)

						gen37338 := Call(__e, gen37335, symlaunch_1repl, gen37337)

						if True == gen37338 {
							gen37341 = True
						} else {
							gen37341 = False
						}

					} else {
						gen37341 = False
					}

					if True == gen37341 {
						gen37334 := Call(__e, PrimNS1Value(symns2_1value), symshen_4repl)

						__e.TailApply(gen37334)

						return

					} else {
						gen37331 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen37332 := Call(__e, gen37331, V4927)

						var gen37333 Obj

						if True == gen37332 {
							gen37326 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen37327 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen37328 := Call(__e, gen37327, V4927)

							gen37329 := Call(__e, gen37326, symshow_1help, gen37328)

							var gen37330 Obj

							if True == gen37329 {
								gen37321 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen37322 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37323 := Call(__e, gen37322, V4927)

								gen37324 := Call(__e, gen37321, gen37323)

								var gen37325 Obj

								if True == gen37324 {
									gen37315 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen37316 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37317 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37318 := Call(__e, gen37317, V4927)

									gen37319 := Call(__e, gen37316, gen37318)

									gen37320 := Call(__e, gen37315, Nil, gen37319)

									if True == gen37320 {
										gen37325 = True
									} else {
										gen37325 = False
									}

								} else {
									gen37325 = False
								}

								if True == gen37325 {
									gen37330 = True
								} else {
									gen37330 = False
								}

							} else {
								gen37330 = False
							}

							if True == gen37330 {
								gen37333 = True
							} else {
								gen37333 = False
							}

						} else {
							gen37333 = False
						}

						if True == gen37333 {
							gen37306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

							gen37307 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

							gen37308 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen37309 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen37310 := Call(__e, gen37309, V4927)

							gen37311 := Call(__e, gen37308, gen37310)

							gen37312 := Call(__e, gen37307, gen37311, MakeString("\n"), symshen_4a)

							gen37313 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

							gen37314 := Call(__e, gen37313)

							__e.TailApply(gen37306, gen37312, gen37314)

							return

						} else {
							gen37303 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen37304 := Call(__e, gen37303, V4927)

							var gen37305 Obj

							if True == gen37304 {
								gen37298 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen37299 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37300 := Call(__e, gen37299, V4927)

								gen37301 := Call(__e, gen37298, symunknown_1arguments, gen37300)

								var gen37302 Obj

								if True == gen37301 {
									gen37293 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen37294 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen37295 := Call(__e, gen37294, V4927)

									gen37296 := Call(__e, gen37293, gen37295)

									var gen37297 Obj

									if True == gen37296 {
										gen37287 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen37288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen37289 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen37290 := Call(__e, gen37289, V4927)

										gen37291 := Call(__e, gen37288, gen37290)

										gen37292 := Call(__e, gen37287, gen37291)

										if True == gen37292 {
											gen37297 = True
										} else {
											gen37297 = False
										}

									} else {
										gen37297 = False
									}

									if True == gen37297 {
										gen37302 = True
									} else {
										gen37302 = False
									}

								} else {
									gen37302 = False
								}

								if True == gen37302 {
									gen37305 = True
								} else {
									gen37305 = False
								}

							} else {
								gen37305 = False
							}

							if True == gen37305 {
								gen37266 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

								gen37267 := Call(__e, PrimNS1Value(symns2_1value), symcn)

								gen37268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

								gen37269 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37270 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37271 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37272 := Call(__e, gen37271, V4927)

								gen37273 := Call(__e, gen37270, gen37272)

								gen37274 := Call(__e, gen37269, gen37273)

								gen37275 := Call(__e, PrimNS1Value(symns2_1value), symcn)

								gen37276 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

								gen37277 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen37278 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen37279 := Call(__e, gen37278, V4927)

								gen37280 := Call(__e, gen37277, gen37279)

								gen37281 := Call(__e, gen37276, gen37280, MakeString(" --help' for more information.\n"), symshen_4a)

								gen37282 := Call(__e, gen37275, MakeString("\nTry `"), gen37281)

								gen37283 := Call(__e, gen37268, gen37274, gen37282, symshen_4a)

								gen37284 := Call(__e, gen37267, MakeString("ERROR: Invalid argument: "), gen37283)

								gen37285 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

								gen37286 := Call(__e, gen37285)

								__e.TailApply(gen37266, gen37284, gen37286)

								return

							} else {
								if True == True {
									gen37265 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

									__e.TailApply(gen37265, symshen_4x_4launcher_4default_1handle_1result)

									return

								} else {
									gen37264 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

									__e.TailApply(gen37264, MakeString("no cond match"))

									return

								}
							}

						}

					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4launcher_4default_1handle_1result, gen37412)

	gen37416 := MakeNative(func(__e *ControlFlow) {
		V4929 := __e.Get(1)
		_ = V4929
		gen37413 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4default_1handle_1result)

		gen37414 := Call(__e, PrimNS1Value(symns2_1value), symshen_4x_4launcher_4launch_1shen)

		gen37415 := Call(__e, gen37414, V4929)

		__e.TailApply(gen37413, gen37415)

		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4launcher_4main, gen37416)

	return

}, 0)
