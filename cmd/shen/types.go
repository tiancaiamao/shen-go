package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp101916 := MakeNative(func(__e *ControlFlow) {
		V3210 := __e.Get(1)
		_ = V3210
		V3211 := __e.Get(2)
		_ = V3211
		tmp101917 := MakeNative(func(__e *ControlFlow) {
			Record := __e.Get(1)
			_ = Record
			tmp101918 := MakeNative(func(__e *ControlFlow) {
				Variancy := __e.Get(1)
				_ = Variancy
				tmp101919 := MakeNative(func(__e *ControlFlow) {
					Type := __e.Get(1)
					_ = Type
					tmp101920 := MakeNative(func(__e *ControlFlow) {
						F_d := __e.Get(1)
						_ = F_d
						tmp101921 := MakeNative(func(__e *ControlFlow) {
							Parameters := __e.Get(1)
							_ = Parameters
							tmp101922 := MakeNative(func(__e *ControlFlow) {
								Clause := __e.Get(1)
								_ = Clause
								tmp101923 := MakeNative(func(__e *ControlFlow) {
									AUM__instruction := __e.Get(1)
									_ = AUM__instruction
									tmp101924 := MakeNative(func(__e *ControlFlow) {
										Code := __e.Get(1)
										_ = Code
										tmp101925 := MakeNative(func(__e *ControlFlow) {
											ShenDef := __e.Get(1)
											_ = ShenDef
											tmp101926 := MakeNative(func(__e *ControlFlow) {
												Eval := __e.Get(1)
												_ = Eval
												__e.Return(V3210)
												return
											}, 1)

											tmp101927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

											tmp101928 := Call(__e, tmp101927, ShenDef)

											__e.TailApply(tmp101926, tmp101928)
											return

										}, 1)

										tmp101929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp101930 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp101931 := Call(__e, PrimNS1Value(symns2_1value), symappend)

										tmp101932 := Call(__e, PrimNS1Value(symns2_1value), symappend)

										tmp101933 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp101934 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp101935 := Call(__e, tmp101934, symContinuation, Nil)

										tmp101936 := Call(__e, tmp101933, symProcessN, tmp101935)

										tmp101937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp101938 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp101939 := Call(__e, tmp101938, Code, Nil)

										tmp101940 := Call(__e, tmp101937, sym_1_6, tmp101939)

										tmp101941 := Call(__e, tmp101932, tmp101936, tmp101940)

										tmp101942 := Call(__e, tmp101931, Parameters, tmp101941)

										tmp101943 := Call(__e, tmp101930, F_d, tmp101942)

										tmp101944 := Call(__e, tmp101929, symdefine, tmp101943)

										__e.TailApply(tmp101925, tmp101944)
										return

									}, 1)

									tmp101945 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

									tmp101946 := Call(__e, tmp101945, AUM__instruction)

									__e.TailApply(tmp101924, tmp101946)
									return

								}, 1)

								tmp101947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum)

								tmp101948 := Call(__e, tmp101947, Clause, Parameters)

								__e.TailApply(tmp101923, tmp101948)
								return

							}, 1)

							tmp101949 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101950 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101951 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101952 := Call(__e, tmp101951, symX, Nil)

							tmp101953 := Call(__e, tmp101950, F_d, tmp101952)

							tmp101954 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101955 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101956 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101957 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101958 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101959 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp101960 := Call(__e, tmp101959, Type, Nil)

							tmp101961 := Call(__e, tmp101958, symX, tmp101960)

							tmp101962 := Call(__e, tmp101957, symunify_b, tmp101961)

							tmp101963 := Call(__e, tmp101956, tmp101962, Nil)

							tmp101964 := Call(__e, tmp101955, tmp101963, Nil)

							tmp101965 := Call(__e, tmp101954, sym_h_1, tmp101964)

							tmp101966 := Call(__e, tmp101949, tmp101953, tmp101965)

							__e.TailApply(tmp101922, tmp101966)
							return

						}, 1)

						tmp101967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4parameters)

						tmp101968 := Call(__e, tmp101967, MakeNumber(1))

						__e.TailApply(tmp101921, tmp101968)
						return

					}, 1)

					tmp101969 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					tmp101970 := Call(__e, tmp101969, symshen_4type_1signature_1of_1, V3210)

					__e.TailApply(tmp101920, tmp101970)
					return

				}, 1)

				tmp101971 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

				tmp101972 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				tmp101973 := Call(__e, tmp101972, V3211)

				tmp101974 := Call(__e, tmp101971, tmp101973)

				__e.TailApply(tmp101919, tmp101974)
				return

			}, 1)

			tmp101975 := MakeNative(func(__e *ControlFlow) {
				tmp101976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variancy_1test)

				__e.TailApply(tmp101976, V3210, V3211)
				return

			}, 0)

			tmp101977 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4skip)
				return
			}, 1)

			tmp101978 := Call(__e, PrimNS1Value(symtry_1catch), tmp101975, tmp101977)

			__e.TailApply(tmp101918, tmp101978)
			return

		}, 1)

		tmp101979 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp101980 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp101981 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp101982 := Call(__e, tmp101981, V3210, V3211)

		tmp101983 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp101984 := Call(__e, tmp101983, symshen_4_dsignedfuncs_d)

		tmp101985 := Call(__e, tmp101980, tmp101982, tmp101984)

		tmp101986 := Call(__e, tmp101979, symshen_4_dsignedfuncs_d, tmp101985)

		__e.TailApply(tmp101917, tmp101986)
		return

	}, 2)

	tmp101987 := Call(__e, PrimNS1Value(symns2_1set), symdeclare, tmp101916)

	_ = tmp101987

	tmp101988 := MakeNative(func(__e *ControlFlow) {
		V3213 := __e.Get(1)
		_ = V3213
		tmp101989 := MakeNative(func(__e *ControlFlow) {
			Demod := __e.Get(1)
			_ = Demod
			tmp101992 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101993 := Call(__e, tmp101992, Demod, V3213)

			if True == tmp101993 {
				__e.Return(V3213)
				return
			} else {
				tmp101991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				__e.TailApply(tmp101991, Demod)
				return

			}

		}, 1)

		tmp101994 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

		tmp101995 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp101996 := Call(__e, tmp101995, symshen_4_ddemodulation_1function_d)

		tmp101997 := Call(__e, tmp101994, tmp101996, V3213)

		__e.TailApply(tmp101989, tmp101997)
		return

	}, 1)

	tmp101998 := Call(__e, PrimNS1Value(symns2_1set), symshen_4demodulate, tmp101988)

	_ = tmp101998

	tmp101999 := MakeNative(func(__e *ControlFlow) {
		V3216 := __e.Get(1)
		_ = V3216
		V3217 := __e.Get(2)
		_ = V3217
		tmp102000 := MakeNative(func(__e *ControlFlow) {
			TypeF := __e.Get(1)
			_ = TypeF
			tmp102001 := MakeNative(func(__e *ControlFlow) {
				Check := __e.Get(1)
				_ = Check
				__e.Return(symshen_4skip)
				return
			}, 1)

			tmp102014 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp102015 := Call(__e, tmp102014, symsymbol, TypeF)

			var ifres102002 Obj

			if True == tmp102015 {
				ifres102002 = symshen_4skip

			} else {
				tmp102012 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

				tmp102013 := Call(__e, tmp102012, TypeF, V3217)

				var ifres102003 Obj

				if True == tmp102013 {
					ifres102003 = symshen_4skip

				} else {
					tmp102004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					tmp102005 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp102006 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp102007 := Call(__e, tmp102006, V3216, MakeString(" may create errors\n"), symshen_4a)

					tmp102008 := Call(__e, tmp102005, MakeString("warning: changing the type of "), tmp102007)

					tmp102009 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					tmp102010 := Call(__e, tmp102009)

					tmp102011 := Call(__e, tmp102004, tmp102008, tmp102010)

					ifres102003 = tmp102011

				}

				ifres102002 = ifres102003

			}

			__e.TailApply(tmp102001, ifres102002)
			return

		}, 1)

		tmp102016 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck)

		tmp102017 := Call(__e, tmp102016, V3216, symB)

		__e.TailApply(tmp102000, tmp102017)
		return

	}, 2)

	tmp102018 := Call(__e, PrimNS1Value(symns2_1set), symshen_4variancy_1test, tmp101999)

	_ = tmp102018

	tmp102019 := MakeNative(func(__e *ControlFlow) {
		V3230 := __e.Get(1)
		_ = V3230
		V3231 := __e.Get(2)
		_ = V3231
		tmp102105 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp102106 := Call(__e, tmp102105, V3231, V3230)

		if True == tmp102106 {
			__e.Return(True)
			return
		} else {
			tmp102103 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp102104 := Call(__e, tmp102103, V3230)

			var ifres102091 Obj

			if True == tmp102104 {
				tmp102101 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp102102 := Call(__e, tmp102101, V3231)

				var ifres102093 Obj

				if True == tmp102102 {
					tmp102095 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp102096 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp102097 := Call(__e, tmp102096, V3231)

					tmp102098 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp102099 := Call(__e, tmp102098, V3230)

					tmp102100 := Call(__e, tmp102095, tmp102097, tmp102099)

					var ifres102094 Obj

					if True == tmp102100 {
						ifres102094 = True

					} else {
						ifres102094 = False

					}

					ifres102093 = ifres102094

				} else {
					ifres102093 = False

				}

				var ifres102092 Obj

				if True == ifres102093 {
					ifres102092 = True

				} else {
					ifres102092 = False

				}

				ifres102091 = ifres102092

			} else {
				ifres102091 = False

			}

			if True == ifres102091 {
				tmp102086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

				tmp102087 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp102088 := Call(__e, tmp102087, V3230)

				tmp102089 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp102090 := Call(__e, tmp102089, V3231)

				__e.TailApply(tmp102086, tmp102088, tmp102090)
				return

			} else {
				tmp102084 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp102085 := Call(__e, tmp102084, V3230)

				var ifres102068 Obj

				if True == tmp102085 {
					tmp102082 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp102083 := Call(__e, tmp102082, V3231)

					var ifres102070 Obj

					if True == tmp102083 {
						tmp102078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

						tmp102079 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp102080 := Call(__e, tmp102079, V3230)

						tmp102081 := Call(__e, tmp102078, tmp102080)

						var ifres102072 Obj

						if True == tmp102081 {
							tmp102074 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

							tmp102075 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp102076 := Call(__e, tmp102075, V3231)

							tmp102077 := Call(__e, tmp102074, tmp102076)

							var ifres102073 Obj

							if True == tmp102077 {
								ifres102073 = True

							} else {
								ifres102073 = False

							}

							ifres102072 = ifres102073

						} else {
							ifres102072 = False

						}

						var ifres102071 Obj

						if True == ifres102072 {
							ifres102071 = True

						} else {
							ifres102071 = False

						}

						ifres102070 = ifres102071

					} else {
						ifres102070 = False

					}

					var ifres102069 Obj

					if True == ifres102070 {
						ifres102069 = True

					} else {
						ifres102069 = False

					}

					ifres102068 = ifres102069

				} else {
					ifres102068 = False

				}

				if True == ifres102068 {
					tmp102055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

					tmp102056 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp102057 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp102058 := Call(__e, tmp102057, V3230)

					tmp102059 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp102060 := Call(__e, tmp102059, V3230)

					tmp102061 := Call(__e, tmp102056, symshen_4a, tmp102058, tmp102060)

					tmp102062 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp102063 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp102064 := Call(__e, tmp102063, V3231)

					tmp102065 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp102066 := Call(__e, tmp102065, V3231)

					tmp102067 := Call(__e, tmp102062, symshen_4a, tmp102064, tmp102066)

					__e.TailApply(tmp102055, tmp102061, tmp102067)
					return

				} else {
					tmp102053 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp102054 := Call(__e, tmp102053, V3230)

					var ifres102037 Obj

					if True == tmp102054 {
						tmp102049 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp102050 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp102051 := Call(__e, tmp102050, V3230)

						tmp102052 := Call(__e, tmp102049, tmp102051)

						var ifres102039 Obj

						if True == tmp102052 {
							tmp102047 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp102048 := Call(__e, tmp102047, V3231)

							var ifres102041 Obj

							if True == tmp102048 {
								tmp102043 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp102044 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp102045 := Call(__e, tmp102044, V3231)

								tmp102046 := Call(__e, tmp102043, tmp102045)

								var ifres102042 Obj

								if True == tmp102046 {
									ifres102042 = True

								} else {
									ifres102042 = False

								}

								ifres102041 = ifres102042

							} else {
								ifres102041 = False

							}

							var ifres102040 Obj

							if True == ifres102041 {
								ifres102040 = True

							} else {
								ifres102040 = False

							}

							ifres102039 = ifres102040

						} else {
							ifres102039 = False

						}

						var ifres102038 Obj

						if True == ifres102039 {
							ifres102038 = True

						} else {
							ifres102038 = False

						}

						ifres102037 = ifres102038

					} else {
						ifres102037 = False

					}

					if True == ifres102037 {
						tmp102024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

						tmp102025 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						tmp102026 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp102027 := Call(__e, tmp102026, V3230)

						tmp102028 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp102029 := Call(__e, tmp102028, V3230)

						tmp102030 := Call(__e, tmp102025, tmp102027, tmp102029)

						tmp102031 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						tmp102032 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp102033 := Call(__e, tmp102032, V3231)

						tmp102034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp102035 := Call(__e, tmp102034, V3231)

						tmp102036 := Call(__e, tmp102031, tmp102033, tmp102035)

						__e.TailApply(tmp102024, tmp102030, tmp102036)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 2)

	tmp102107 := Call(__e, PrimNS1Value(symns2_1set), symshen_4variant_2, tmp102019)

	_ = tmp102107

	tmp102108 := MakeNative(func(__e *ControlFlow) {
		V3236 := __e.Get(1)
		_ = V3236
		V3237 := __e.Get(2)
		_ = V3237
		V3238 := __e.Get(3)
		_ = V3238
		tmp102109 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102110 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102111 := Call(__e, tmp102110)

			_ = tmp102111

			tmp102112 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102115 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102116 := Call(__e, tmp102115, symboolean, Nil)

			tmp102117 := Call(__e, tmp102114, sym_1_1_6, tmp102116)

			tmp102118 := Call(__e, tmp102113, A, tmp102117)

			__e.TailApply(tmp102112, V3236, tmp102118, V3237, V3238)
			return

		}, 1)

		tmp102119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102120 := Call(__e, tmp102119, V3237)

		__e.TailApply(tmp102109, tmp102120)
		return

	}, 3)

	tmp102121 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1absvector_2, tmp102108)

	_ = tmp102121

	tmp102122 := MakeNative(func(__e *ControlFlow) {
		V3246 := __e.Get(1)
		_ = V3246
		V3247 := __e.Get(2)
		_ = V3247
		V3248 := __e.Get(3)
		_ = V3248
		tmp102123 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102125 := Call(__e, tmp102124)

			_ = tmp102125

			tmp102126 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102127 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102133 := Call(__e, tmp102132, A, Nil)

			tmp102134 := Call(__e, tmp102131, symlist, tmp102133)

			tmp102135 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102136 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102137 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102138 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102139 := Call(__e, tmp102138, A, Nil)

			tmp102140 := Call(__e, tmp102137, symlist, tmp102139)

			tmp102141 := Call(__e, tmp102136, tmp102140, Nil)

			tmp102142 := Call(__e, tmp102135, sym_1_1_6, tmp102141)

			tmp102143 := Call(__e, tmp102130, tmp102134, tmp102142)

			tmp102144 := Call(__e, tmp102129, tmp102143, Nil)

			tmp102145 := Call(__e, tmp102128, sym_1_1_6, tmp102144)

			tmp102146 := Call(__e, tmp102127, A, tmp102145)

			__e.TailApply(tmp102126, V3246, tmp102146, V3247, V3248)
			return

		}, 1)

		tmp102147 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102148 := Call(__e, tmp102147, V3247)

		__e.TailApply(tmp102123, tmp102148)
		return

	}, 3)

	tmp102149 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1adjoin, tmp102122)

	_ = tmp102149

	tmp102150 := MakeNative(func(__e *ControlFlow) {
		V3256 := __e.Get(1)
		_ = V3256
		V3257 := __e.Get(2)
		_ = V3257
		V3258 := __e.Get(3)
		_ = V3258
		tmp102151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102152 := Call(__e, tmp102151)

		_ = tmp102152

		tmp102153 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102154 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102157 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102158 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102159 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102160 := Call(__e, tmp102159, symboolean, Nil)

		tmp102161 := Call(__e, tmp102158, sym_1_1_6, tmp102160)

		tmp102162 := Call(__e, tmp102157, symboolean, tmp102161)

		tmp102163 := Call(__e, tmp102156, tmp102162, Nil)

		tmp102164 := Call(__e, tmp102155, sym_1_1_6, tmp102163)

		tmp102165 := Call(__e, tmp102154, symboolean, tmp102164)

		__e.TailApply(tmp102153, V3256, tmp102165, V3257, V3258)
		return

	}, 3)

	tmp102166 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1and, tmp102150)

	_ = tmp102166

	tmp102167 := MakeNative(func(__e *ControlFlow) {
		V3266 := __e.Get(1)
		_ = V3266
		V3267 := __e.Get(2)
		_ = V3267
		V3268 := __e.Get(3)
		_ = V3268
		tmp102168 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102169 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102170 := Call(__e, tmp102169)

			_ = tmp102170

			tmp102171 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102174 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102175 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102177 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102178 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102179 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102180 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102181 := Call(__e, tmp102180, symstring, Nil)

			tmp102182 := Call(__e, tmp102179, sym_1_1_6, tmp102181)

			tmp102183 := Call(__e, tmp102178, symsymbol, tmp102182)

			tmp102184 := Call(__e, tmp102177, tmp102183, Nil)

			tmp102185 := Call(__e, tmp102176, sym_1_1_6, tmp102184)

			tmp102186 := Call(__e, tmp102175, symstring, tmp102185)

			tmp102187 := Call(__e, tmp102174, tmp102186, Nil)

			tmp102188 := Call(__e, tmp102173, sym_1_1_6, tmp102187)

			tmp102189 := Call(__e, tmp102172, A, tmp102188)

			__e.TailApply(tmp102171, V3266, tmp102189, V3267, V3268)
			return

		}, 1)

		tmp102190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102191 := Call(__e, tmp102190, V3267)

		__e.TailApply(tmp102168, tmp102191)
		return

	}, 3)

	tmp102192 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4app, tmp102167)

	_ = tmp102192

	tmp102193 := MakeNative(func(__e *ControlFlow) {
		V3276 := __e.Get(1)
		_ = V3276
		V3277 := __e.Get(2)
		_ = V3277
		V3278 := __e.Get(3)
		_ = V3278
		tmp102194 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102195 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102196 := Call(__e, tmp102195)

			_ = tmp102196

			tmp102197 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102198 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102199 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102200 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102201 := Call(__e, tmp102200, A, Nil)

			tmp102202 := Call(__e, tmp102199, symlist, tmp102201)

			tmp102203 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102204 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102205 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102206 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102207 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102208 := Call(__e, tmp102207, A, Nil)

			tmp102209 := Call(__e, tmp102206, symlist, tmp102208)

			tmp102210 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102211 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102213 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102214 := Call(__e, tmp102213, A, Nil)

			tmp102215 := Call(__e, tmp102212, symlist, tmp102214)

			tmp102216 := Call(__e, tmp102211, tmp102215, Nil)

			tmp102217 := Call(__e, tmp102210, sym_1_1_6, tmp102216)

			tmp102218 := Call(__e, tmp102205, tmp102209, tmp102217)

			tmp102219 := Call(__e, tmp102204, tmp102218, Nil)

			tmp102220 := Call(__e, tmp102203, sym_1_1_6, tmp102219)

			tmp102221 := Call(__e, tmp102198, tmp102202, tmp102220)

			__e.TailApply(tmp102197, V3276, tmp102221, V3277, V3278)
			return

		}, 1)

		tmp102222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102223 := Call(__e, tmp102222, V3277)

		__e.TailApply(tmp102194, tmp102223)
		return

	}, 3)

	tmp102224 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1append, tmp102193)

	_ = tmp102224

	tmp102225 := MakeNative(func(__e *ControlFlow) {
		V3286 := __e.Get(1)
		_ = V3286
		V3287 := __e.Get(2)
		_ = V3287
		V3288 := __e.Get(3)
		_ = V3288
		tmp102226 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102227 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102228 := Call(__e, tmp102227)

			_ = tmp102228

			tmp102229 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102230 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102231 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102232 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102233 := Call(__e, tmp102232, symnumber, Nil)

			tmp102234 := Call(__e, tmp102231, sym_1_1_6, tmp102233)

			tmp102235 := Call(__e, tmp102230, A, tmp102234)

			__e.TailApply(tmp102229, V3286, tmp102235, V3287, V3288)
			return

		}, 1)

		tmp102236 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102237 := Call(__e, tmp102236, V3287)

		__e.TailApply(tmp102226, tmp102237)
		return

	}, 3)

	tmp102238 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1arity, tmp102225)

	_ = tmp102238

	tmp102239 := MakeNative(func(__e *ControlFlow) {
		V3296 := __e.Get(1)
		_ = V3296
		V3297 := __e.Get(2)
		_ = V3297
		V3298 := __e.Get(3)
		_ = V3298
		tmp102240 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102241 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102242 := Call(__e, tmp102241)

			_ = tmp102242

			tmp102243 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102244 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102245 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102246 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102247 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102248 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102249 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102250 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102251 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102252 := Call(__e, tmp102251, A, Nil)

			tmp102253 := Call(__e, tmp102250, symlist, tmp102252)

			tmp102254 := Call(__e, tmp102249, tmp102253, Nil)

			tmp102255 := Call(__e, tmp102248, symlist, tmp102254)

			tmp102256 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102257 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102258 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102260 := Call(__e, tmp102259, A, Nil)

			tmp102261 := Call(__e, tmp102258, symlist, tmp102260)

			tmp102262 := Call(__e, tmp102257, tmp102261, Nil)

			tmp102263 := Call(__e, tmp102256, sym_1_1_6, tmp102262)

			tmp102264 := Call(__e, tmp102247, tmp102255, tmp102263)

			tmp102265 := Call(__e, tmp102246, tmp102264, Nil)

			tmp102266 := Call(__e, tmp102245, sym_1_1_6, tmp102265)

			tmp102267 := Call(__e, tmp102244, A, tmp102266)

			__e.TailApply(tmp102243, V3296, tmp102267, V3297, V3298)
			return

		}, 1)

		tmp102268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102269 := Call(__e, tmp102268, V3297)

		__e.TailApply(tmp102240, tmp102269)
		return

	}, 3)

	tmp102270 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1assoc, tmp102239)

	_ = tmp102270

	tmp102271 := MakeNative(func(__e *ControlFlow) {
		V3306 := __e.Get(1)
		_ = V3306
		V3307 := __e.Get(2)
		_ = V3307
		V3308 := __e.Get(3)
		_ = V3308
		tmp102272 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102273 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102274 := Call(__e, tmp102273)

			_ = tmp102274

			tmp102275 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102276 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102279 := Call(__e, tmp102278, symboolean, Nil)

			tmp102280 := Call(__e, tmp102277, sym_1_1_6, tmp102279)

			tmp102281 := Call(__e, tmp102276, A, tmp102280)

			__e.TailApply(tmp102275, V3306, tmp102281, V3307, V3308)
			return

		}, 1)

		tmp102282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102283 := Call(__e, tmp102282, V3307)

		__e.TailApply(tmp102272, tmp102283)
		return

	}, 3)

	tmp102284 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1boolean_2, tmp102271)

	_ = tmp102284

	tmp102285 := MakeNative(func(__e *ControlFlow) {
		V3316 := __e.Get(1)
		_ = V3316
		V3317 := __e.Get(2)
		_ = V3317
		V3318 := __e.Get(3)
		_ = V3318
		tmp102286 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102287 := Call(__e, tmp102286)

		_ = tmp102287

		tmp102288 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102289 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102290 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102291 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102292 := Call(__e, tmp102291, symboolean, Nil)

		tmp102293 := Call(__e, tmp102290, sym_1_1_6, tmp102292)

		tmp102294 := Call(__e, tmp102289, symsymbol, tmp102293)

		__e.TailApply(tmp102288, V3316, tmp102294, V3317, V3318)
		return

	}, 3)

	tmp102295 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1bound_2, tmp102285)

	_ = tmp102295

	tmp102296 := MakeNative(func(__e *ControlFlow) {
		V3326 := __e.Get(1)
		_ = V3326
		V3327 := __e.Get(2)
		_ = V3327
		V3328 := __e.Get(3)
		_ = V3328
		tmp102297 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102298 := Call(__e, tmp102297)

		_ = tmp102298

		tmp102299 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102300 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102301 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102302 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102303 := Call(__e, tmp102302, symstring, Nil)

		tmp102304 := Call(__e, tmp102301, sym_1_1_6, tmp102303)

		tmp102305 := Call(__e, tmp102300, symstring, tmp102304)

		__e.TailApply(tmp102299, V3326, tmp102305, V3327, V3328)
		return

	}, 3)

	tmp102306 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cd, tmp102296)

	_ = tmp102306

	tmp102307 := MakeNative(func(__e *ControlFlow) {
		V3336 := __e.Get(1)
		_ = V3336
		V3337 := __e.Get(2)
		_ = V3337
		V3338 := __e.Get(3)
		_ = V3338
		tmp102308 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102309 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp102310 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102311 := Call(__e, tmp102310)

				_ = tmp102311

				tmp102312 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102314 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102316 := Call(__e, tmp102315, A, Nil)

				tmp102317 := Call(__e, tmp102314, symstream, tmp102316)

				tmp102318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102322 := Call(__e, tmp102321, B, Nil)

				tmp102323 := Call(__e, tmp102320, symlist, tmp102322)

				tmp102324 := Call(__e, tmp102319, tmp102323, Nil)

				tmp102325 := Call(__e, tmp102318, sym_1_1_6, tmp102324)

				tmp102326 := Call(__e, tmp102313, tmp102317, tmp102325)

				__e.TailApply(tmp102312, V3336, tmp102326, V3337, V3338)
				return

			}, 1)

			tmp102327 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102328 := Call(__e, tmp102327, V3337)

			__e.TailApply(tmp102309, tmp102328)
			return

		}, 1)

		tmp102329 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102330 := Call(__e, tmp102329, V3337)

		__e.TailApply(tmp102308, tmp102330)
		return

	}, 3)

	tmp102331 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1close, tmp102307)

	_ = tmp102331

	tmp102332 := MakeNative(func(__e *ControlFlow) {
		V3346 := __e.Get(1)
		_ = V3346
		V3347 := __e.Get(2)
		_ = V3347
		V3348 := __e.Get(3)
		_ = V3348
		tmp102333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102334 := Call(__e, tmp102333)

		_ = tmp102334

		tmp102335 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102339 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102340 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102342 := Call(__e, tmp102341, symstring, Nil)

		tmp102343 := Call(__e, tmp102340, sym_1_1_6, tmp102342)

		tmp102344 := Call(__e, tmp102339, symstring, tmp102343)

		tmp102345 := Call(__e, tmp102338, tmp102344, Nil)

		tmp102346 := Call(__e, tmp102337, sym_1_1_6, tmp102345)

		tmp102347 := Call(__e, tmp102336, symstring, tmp102346)

		__e.TailApply(tmp102335, V3346, tmp102347, V3347, V3348)
		return

	}, 3)

	tmp102348 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cn, tmp102332)

	_ = tmp102348

	tmp102349 := MakeNative(func(__e *ControlFlow) {
		V3356 := __e.Get(1)
		_ = V3356
		V3357 := __e.Get(2)
		_ = V3357
		V3358 := __e.Get(3)
		_ = V3358
		tmp102350 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102351 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp102352 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102353 := Call(__e, tmp102352)

				_ = tmp102353

				tmp102354 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102359 := Call(__e, tmp102358, B, Nil)

				tmp102360 := Call(__e, tmp102357, symshen_4_a_a_6, tmp102359)

				tmp102361 := Call(__e, tmp102356, A, tmp102360)

				tmp102362 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102363 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102364 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102365 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102368 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102370 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102371 := Call(__e, tmp102370, B, Nil)

				tmp102372 := Call(__e, tmp102369, sym_1_1_6, tmp102371)

				tmp102373 := Call(__e, tmp102368, A, tmp102372)

				tmp102374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102376 := Call(__e, tmp102375, B, Nil)

				tmp102377 := Call(__e, tmp102374, sym_1_1_6, tmp102376)

				tmp102378 := Call(__e, tmp102367, tmp102373, tmp102377)

				tmp102379 := Call(__e, tmp102366, tmp102378, Nil)

				tmp102380 := Call(__e, tmp102365, sym_1_1_6, tmp102379)

				tmp102381 := Call(__e, tmp102364, A, tmp102380)

				tmp102382 := Call(__e, tmp102363, tmp102381, Nil)

				tmp102383 := Call(__e, tmp102362, sym_1_1_6, tmp102382)

				tmp102384 := Call(__e, tmp102355, tmp102361, tmp102383)

				__e.TailApply(tmp102354, V3356, tmp102384, V3357, V3358)
				return

			}, 1)

			tmp102385 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102386 := Call(__e, tmp102385, V3357)

			__e.TailApply(tmp102351, tmp102386)
			return

		}, 1)

		tmp102387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102388 := Call(__e, tmp102387, V3357)

		__e.TailApply(tmp102350, tmp102388)
		return

	}, 3)

	tmp102389 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1compile, tmp102349)

	_ = tmp102389

	tmp102390 := MakeNative(func(__e *ControlFlow) {
		V3366 := __e.Get(1)
		_ = V3366
		V3367 := __e.Get(2)
		_ = V3367
		V3368 := __e.Get(3)
		_ = V3368
		tmp102391 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102393 := Call(__e, tmp102392)

			_ = tmp102393

			tmp102394 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102396 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102397 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102398 := Call(__e, tmp102397, symboolean, Nil)

			tmp102399 := Call(__e, tmp102396, sym_1_1_6, tmp102398)

			tmp102400 := Call(__e, tmp102395, A, tmp102399)

			__e.TailApply(tmp102394, V3366, tmp102400, V3367, V3368)
			return

		}, 1)

		tmp102401 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102402 := Call(__e, tmp102401, V3367)

		__e.TailApply(tmp102391, tmp102402)
		return

	}, 3)

	tmp102403 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cons_2, tmp102390)

	_ = tmp102403

	tmp102404 := MakeNative(func(__e *ControlFlow) {
		V3376 := __e.Get(1)
		_ = V3376
		V3377 := __e.Get(2)
		_ = V3377
		V3378 := __e.Get(3)
		_ = V3378
		tmp102405 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102406 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp102407 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102408 := Call(__e, tmp102407)

				_ = tmp102408

				tmp102409 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102410 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102413 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102414 := Call(__e, tmp102413, B, Nil)

				tmp102415 := Call(__e, tmp102412, sym_1_1_6, tmp102414)

				tmp102416 := Call(__e, tmp102411, A, tmp102415)

				tmp102417 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102418 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102419 := Call(__e, tmp102418, symsymbol, Nil)

				tmp102420 := Call(__e, tmp102417, sym_1_1_6, tmp102419)

				tmp102421 := Call(__e, tmp102410, tmp102416, tmp102420)

				__e.TailApply(tmp102409, V3376, tmp102421, V3377, V3378)
				return

			}, 1)

			tmp102422 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102423 := Call(__e, tmp102422, V3377)

			__e.TailApply(tmp102406, tmp102423)
			return

		}, 1)

		tmp102424 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102425 := Call(__e, tmp102424, V3377)

		__e.TailApply(tmp102405, tmp102425)
		return

	}, 3)

	tmp102426 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1destroy, tmp102404)

	_ = tmp102426

	tmp102427 := MakeNative(func(__e *ControlFlow) {
		V3386 := __e.Get(1)
		_ = V3386
		V3387 := __e.Get(2)
		_ = V3387
		V3388 := __e.Get(3)
		_ = V3388
		tmp102428 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102429 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102430 := Call(__e, tmp102429)

			_ = tmp102430

			tmp102431 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102432 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102434 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102435 := Call(__e, tmp102434, A, Nil)

			tmp102436 := Call(__e, tmp102433, symlist, tmp102435)

			tmp102437 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102438 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102439 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102440 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102441 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102442 := Call(__e, tmp102441, A, Nil)

			tmp102443 := Call(__e, tmp102440, symlist, tmp102442)

			tmp102444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102448 := Call(__e, tmp102447, A, Nil)

			tmp102449 := Call(__e, tmp102446, symlist, tmp102448)

			tmp102450 := Call(__e, tmp102445, tmp102449, Nil)

			tmp102451 := Call(__e, tmp102444, sym_1_1_6, tmp102450)

			tmp102452 := Call(__e, tmp102439, tmp102443, tmp102451)

			tmp102453 := Call(__e, tmp102438, tmp102452, Nil)

			tmp102454 := Call(__e, tmp102437, sym_1_1_6, tmp102453)

			tmp102455 := Call(__e, tmp102432, tmp102436, tmp102454)

			__e.TailApply(tmp102431, V3386, tmp102455, V3387, V3388)
			return

		}, 1)

		tmp102456 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102457 := Call(__e, tmp102456, V3387)

		__e.TailApply(tmp102428, tmp102457)
		return

	}, 3)

	tmp102458 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1difference, tmp102427)

	_ = tmp102458

	tmp102459 := MakeNative(func(__e *ControlFlow) {
		V3396 := __e.Get(1)
		_ = V3396
		V3397 := __e.Get(2)
		_ = V3397
		V3398 := __e.Get(3)
		_ = V3398
		tmp102460 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102461 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp102462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102463 := Call(__e, tmp102462)

				_ = tmp102463

				tmp102464 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102465 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102466 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102467 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102468 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102469 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102470 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102471 := Call(__e, tmp102470, B, Nil)

				tmp102472 := Call(__e, tmp102469, sym_1_1_6, tmp102471)

				tmp102473 := Call(__e, tmp102468, B, tmp102472)

				tmp102474 := Call(__e, tmp102467, tmp102473, Nil)

				tmp102475 := Call(__e, tmp102466, sym_1_1_6, tmp102474)

				tmp102476 := Call(__e, tmp102465, A, tmp102475)

				__e.TailApply(tmp102464, V3396, tmp102476, V3397, V3398)
				return

			}, 1)

			tmp102477 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102478 := Call(__e, tmp102477, V3397)

			__e.TailApply(tmp102461, tmp102478)
			return

		}, 1)

		tmp102479 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102480 := Call(__e, tmp102479, V3397)

		__e.TailApply(tmp102460, tmp102480)
		return

	}, 3)

	tmp102481 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1do, tmp102459)

	_ = tmp102481

	tmp102482 := MakeNative(func(__e *ControlFlow) {
		V3406 := __e.Get(1)
		_ = V3406
		V3407 := __e.Get(2)
		_ = V3407
		V3408 := __e.Get(3)
		_ = V3408
		tmp102483 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102484 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp102485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102486 := Call(__e, tmp102485)

				_ = tmp102486

				tmp102487 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102488 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102489 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102490 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102491 := Call(__e, tmp102490, A, Nil)

				tmp102492 := Call(__e, tmp102489, symlist, tmp102491)

				tmp102493 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102494 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102495 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102496 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102497 := Call(__e, tmp102496, B, Nil)

				tmp102498 := Call(__e, tmp102495, symlist, tmp102497)

				tmp102499 := Call(__e, tmp102494, tmp102498, Nil)

				tmp102500 := Call(__e, tmp102493, symshen_4_a_a_6, tmp102499)

				tmp102501 := Call(__e, tmp102488, tmp102492, tmp102500)

				__e.TailApply(tmp102487, V3406, tmp102501, V3407, V3408)
				return

			}, 1)

			tmp102502 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102503 := Call(__e, tmp102502, V3407)

			__e.TailApply(tmp102484, tmp102503)
			return

		}, 1)

		tmp102504 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102505 := Call(__e, tmp102504, V3407)

		__e.TailApply(tmp102483, tmp102505)
		return

	}, 3)

	tmp102506 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5e_6, tmp102482)

	_ = tmp102506

	tmp102507 := MakeNative(func(__e *ControlFlow) {
		V3416 := __e.Get(1)
		_ = V3416
		V3417 := __e.Get(2)
		_ = V3417
		V3418 := __e.Get(3)
		_ = V3418
		tmp102508 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102509 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102510 := Call(__e, tmp102509)

			_ = tmp102510

			tmp102511 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102512 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102513 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102514 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102515 := Call(__e, tmp102514, A, Nil)

			tmp102516 := Call(__e, tmp102513, symlist, tmp102515)

			tmp102517 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102518 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102519 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102520 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102521 := Call(__e, tmp102520, A, Nil)

			tmp102522 := Call(__e, tmp102519, symlist, tmp102521)

			tmp102523 := Call(__e, tmp102518, tmp102522, Nil)

			tmp102524 := Call(__e, tmp102517, symshen_4_a_a_6, tmp102523)

			tmp102525 := Call(__e, tmp102512, tmp102516, tmp102524)

			__e.TailApply(tmp102511, V3416, tmp102525, V3417, V3418)
			return

		}, 1)

		tmp102526 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102527 := Call(__e, tmp102526, V3417)

		__e.TailApply(tmp102508, tmp102527)
		return

	}, 3)

	tmp102528 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_b_6, tmp102507)

	_ = tmp102528

	tmp102529 := MakeNative(func(__e *ControlFlow) {
		V3426 := __e.Get(1)
		_ = V3426
		V3427 := __e.Get(2)
		_ = V3427
		V3428 := __e.Get(3)
		_ = V3428
		tmp102530 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102531 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102532 := Call(__e, tmp102531)

			_ = tmp102532

			tmp102533 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102534 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102535 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102536 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102537 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102538 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102539 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102540 := Call(__e, tmp102539, A, Nil)

			tmp102541 := Call(__e, tmp102538, symlist, tmp102540)

			tmp102542 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102543 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102544 := Call(__e, tmp102543, symboolean, Nil)

			tmp102545 := Call(__e, tmp102542, sym_1_1_6, tmp102544)

			tmp102546 := Call(__e, tmp102537, tmp102541, tmp102545)

			tmp102547 := Call(__e, tmp102536, tmp102546, Nil)

			tmp102548 := Call(__e, tmp102535, sym_1_1_6, tmp102547)

			tmp102549 := Call(__e, tmp102534, A, tmp102548)

			__e.TailApply(tmp102533, V3426, tmp102549, V3427, V3428)
			return

		}, 1)

		tmp102550 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102551 := Call(__e, tmp102550, V3427)

		__e.TailApply(tmp102530, tmp102551)
		return

	}, 3)

	tmp102552 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1element_2, tmp102529)

	_ = tmp102552

	tmp102553 := MakeNative(func(__e *ControlFlow) {
		V3436 := __e.Get(1)
		_ = V3436
		V3437 := __e.Get(2)
		_ = V3437
		V3438 := __e.Get(3)
		_ = V3438
		tmp102554 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102555 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102556 := Call(__e, tmp102555)

			_ = tmp102556

			tmp102557 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102558 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102559 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102561 := Call(__e, tmp102560, symboolean, Nil)

			tmp102562 := Call(__e, tmp102559, sym_1_1_6, tmp102561)

			tmp102563 := Call(__e, tmp102558, A, tmp102562)

			__e.TailApply(tmp102557, V3436, tmp102563, V3437, V3438)
			return

		}, 1)

		tmp102564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102565 := Call(__e, tmp102564, V3437)

		__e.TailApply(tmp102554, tmp102565)
		return

	}, 3)

	tmp102566 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1empty_2, tmp102553)

	_ = tmp102566

	tmp102567 := MakeNative(func(__e *ControlFlow) {
		V3446 := __e.Get(1)
		_ = V3446
		V3447 := __e.Get(2)
		_ = V3447
		V3448 := __e.Get(3)
		_ = V3448
		tmp102568 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102569 := Call(__e, tmp102568)

		_ = tmp102569

		tmp102570 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102572 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102573 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102574 := Call(__e, tmp102573, symboolean, Nil)

		tmp102575 := Call(__e, tmp102572, sym_1_1_6, tmp102574)

		tmp102576 := Call(__e, tmp102571, symsymbol, tmp102575)

		__e.TailApply(tmp102570, V3446, tmp102576, V3447, V3448)
		return

	}, 3)

	tmp102577 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1enable_1type_1theory, tmp102567)

	_ = tmp102577

	tmp102578 := MakeNative(func(__e *ControlFlow) {
		V3456 := __e.Get(1)
		_ = V3456
		V3457 := __e.Get(2)
		_ = V3457
		V3458 := __e.Get(3)
		_ = V3458
		tmp102579 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102580 := Call(__e, tmp102579)

		_ = tmp102580

		tmp102581 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102583 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102584 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102585 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102586 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102587 := Call(__e, tmp102586, symsymbol, Nil)

		tmp102588 := Call(__e, tmp102585, symlist, tmp102587)

		tmp102589 := Call(__e, tmp102584, tmp102588, Nil)

		tmp102590 := Call(__e, tmp102583, sym_1_1_6, tmp102589)

		tmp102591 := Call(__e, tmp102582, symsymbol, tmp102590)

		__e.TailApply(tmp102581, V3456, tmp102591, V3457, V3458)
		return

	}, 3)

	tmp102592 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1external, tmp102578)

	_ = tmp102592

	tmp102593 := MakeNative(func(__e *ControlFlow) {
		V3466 := __e.Get(1)
		_ = V3466
		V3467 := __e.Get(2)
		_ = V3467
		V3468 := __e.Get(3)
		_ = V3468
		tmp102594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102595 := Call(__e, tmp102594)

		_ = tmp102595

		tmp102596 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102600 := Call(__e, tmp102599, symstring, Nil)

		tmp102601 := Call(__e, tmp102598, sym_1_1_6, tmp102600)

		tmp102602 := Call(__e, tmp102597, symexception, tmp102601)

		__e.TailApply(tmp102596, V3466, tmp102602, V3467, V3468)
		return

	}, 3)

	tmp102603 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1error_1to_1string, tmp102593)

	_ = tmp102603

	tmp102604 := MakeNative(func(__e *ControlFlow) {
		V3476 := __e.Get(1)
		_ = V3476
		V3477 := __e.Get(2)
		_ = V3477
		V3478 := __e.Get(3)
		_ = V3478
		tmp102605 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102607 := Call(__e, tmp102606)

			_ = tmp102607

			tmp102608 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102609 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102610 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102611 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102612 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102613 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102614 := Call(__e, tmp102613, symstring, Nil)

			tmp102615 := Call(__e, tmp102612, symlist, tmp102614)

			tmp102616 := Call(__e, tmp102611, tmp102615, Nil)

			tmp102617 := Call(__e, tmp102610, sym_1_1_6, tmp102616)

			tmp102618 := Call(__e, tmp102609, A, tmp102617)

			__e.TailApply(tmp102608, V3476, tmp102618, V3477, V3478)
			return

		}, 1)

		tmp102619 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102620 := Call(__e, tmp102619, V3477)

		__e.TailApply(tmp102605, tmp102620)
		return

	}, 3)

	tmp102621 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1explode, tmp102604)

	_ = tmp102621

	tmp102622 := MakeNative(func(__e *ControlFlow) {
		V3486 := __e.Get(1)
		_ = V3486
		V3487 := __e.Get(2)
		_ = V3487
		V3488 := __e.Get(3)
		_ = V3488
		tmp102623 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102624 := Call(__e, tmp102623)

		_ = tmp102624

		tmp102625 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102626 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102628 := Call(__e, tmp102627, symsymbol, Nil)

		tmp102629 := Call(__e, tmp102626, sym_1_1_6, tmp102628)

		__e.TailApply(tmp102625, V3486, tmp102629, V3487, V3488)
		return

	}, 3)

	tmp102630 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fail, tmp102622)

	_ = tmp102630

	tmp102631 := MakeNative(func(__e *ControlFlow) {
		V3496 := __e.Get(1)
		_ = V3496
		V3497 := __e.Get(2)
		_ = V3497
		V3498 := __e.Get(3)
		_ = V3498
		tmp102632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102633 := Call(__e, tmp102632)

		_ = tmp102633

		tmp102634 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102635 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102636 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102637 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102638 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102639 := Call(__e, tmp102638, symboolean, Nil)

		tmp102640 := Call(__e, tmp102637, sym_1_1_6, tmp102639)

		tmp102641 := Call(__e, tmp102636, symsymbol, tmp102640)

		tmp102642 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102643 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102644 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102645 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102646 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102647 := Call(__e, tmp102646, symsymbol, Nil)

		tmp102648 := Call(__e, tmp102645, sym_1_1_6, tmp102647)

		tmp102649 := Call(__e, tmp102644, symsymbol, tmp102648)

		tmp102650 := Call(__e, tmp102643, tmp102649, Nil)

		tmp102651 := Call(__e, tmp102642, sym_1_1_6, tmp102650)

		tmp102652 := Call(__e, tmp102635, tmp102641, tmp102651)

		__e.TailApply(tmp102634, V3496, tmp102652, V3497, V3498)
		return

	}, 3)

	tmp102653 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fail_1if, tmp102631)

	_ = tmp102653

	tmp102654 := MakeNative(func(__e *ControlFlow) {
		V3506 := __e.Get(1)
		_ = V3506
		V3507 := __e.Get(2)
		_ = V3507
		V3508 := __e.Get(3)
		_ = V3508
		tmp102655 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102656 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102657 := Call(__e, tmp102656)

			_ = tmp102657

			tmp102658 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102659 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102660 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102663 := Call(__e, tmp102662, A, Nil)

			tmp102664 := Call(__e, tmp102661, sym_1_1_6, tmp102663)

			tmp102665 := Call(__e, tmp102660, A, tmp102664)

			tmp102666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102668 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102669 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102670 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102671 := Call(__e, tmp102670, A, Nil)

			tmp102672 := Call(__e, tmp102669, sym_1_1_6, tmp102671)

			tmp102673 := Call(__e, tmp102668, A, tmp102672)

			tmp102674 := Call(__e, tmp102667, tmp102673, Nil)

			tmp102675 := Call(__e, tmp102666, sym_1_1_6, tmp102674)

			tmp102676 := Call(__e, tmp102659, tmp102665, tmp102675)

			__e.TailApply(tmp102658, V3506, tmp102676, V3507, V3508)
			return

		}, 1)

		tmp102677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102678 := Call(__e, tmp102677, V3507)

		__e.TailApply(tmp102655, tmp102678)
		return

	}, 3)

	tmp102679 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fix, tmp102654)

	_ = tmp102679

	tmp102680 := MakeNative(func(__e *ControlFlow) {
		V3516 := __e.Get(1)
		_ = V3516
		V3517 := __e.Get(2)
		_ = V3517
		V3518 := __e.Get(3)
		_ = V3518
		tmp102681 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102682 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102683 := Call(__e, tmp102682)

			_ = tmp102683

			tmp102684 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102685 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102686 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102688 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102689 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102690 := Call(__e, tmp102689, A, Nil)

			tmp102691 := Call(__e, tmp102688, symlazy, tmp102690)

			tmp102692 := Call(__e, tmp102687, tmp102691, Nil)

			tmp102693 := Call(__e, tmp102686, sym_1_1_6, tmp102692)

			tmp102694 := Call(__e, tmp102685, A, tmp102693)

			__e.TailApply(tmp102684, V3516, tmp102694, V3517, V3518)
			return

		}, 1)

		tmp102695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102696 := Call(__e, tmp102695, V3517)

		__e.TailApply(tmp102681, tmp102696)
		return

	}, 3)

	tmp102697 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1freeze, tmp102680)

	_ = tmp102697

	tmp102698 := MakeNative(func(__e *ControlFlow) {
		V3526 := __e.Get(1)
		_ = V3526
		V3527 := __e.Get(2)
		_ = V3527
		V3528 := __e.Get(3)
		_ = V3528
		tmp102699 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			tmp102700 := MakeNative(func(__e *ControlFlow) {
				A := __e.Get(1)
				_ = A
				tmp102701 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102702 := Call(__e, tmp102701)

				_ = tmp102702

				tmp102703 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102704 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102705 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102707 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102708 := Call(__e, tmp102707, B, Nil)

				tmp102709 := Call(__e, tmp102706, sym_d, tmp102708)

				tmp102710 := Call(__e, tmp102705, A, tmp102709)

				tmp102711 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102712 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102713 := Call(__e, tmp102712, A, Nil)

				tmp102714 := Call(__e, tmp102711, sym_1_1_6, tmp102713)

				tmp102715 := Call(__e, tmp102704, tmp102710, tmp102714)

				__e.TailApply(tmp102703, V3526, tmp102715, V3527, V3528)
				return

			}, 1)

			tmp102716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102717 := Call(__e, tmp102716, V3527)

			__e.TailApply(tmp102700, tmp102717)
			return

		}, 1)

		tmp102718 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102719 := Call(__e, tmp102718, V3527)

		__e.TailApply(tmp102699, tmp102719)
		return

	}, 3)

	tmp102720 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fst, tmp102698)

	_ = tmp102720

	tmp102721 := MakeNative(func(__e *ControlFlow) {
		V3536 := __e.Get(1)
		_ = V3536
		V3537 := __e.Get(2)
		_ = V3537
		V3538 := __e.Get(3)
		_ = V3538
		tmp102722 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102723 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp102724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp102725 := Call(__e, tmp102724)

				_ = tmp102725

				tmp102726 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp102727 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102730 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102731 := Call(__e, tmp102730, B, Nil)

				tmp102732 := Call(__e, tmp102729, sym_1_1_6, tmp102731)

				tmp102733 := Call(__e, tmp102728, A, tmp102732)

				tmp102734 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102736 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102737 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp102739 := Call(__e, tmp102738, B, Nil)

				tmp102740 := Call(__e, tmp102737, sym_1_1_6, tmp102739)

				tmp102741 := Call(__e, tmp102736, A, tmp102740)

				tmp102742 := Call(__e, tmp102735, tmp102741, Nil)

				tmp102743 := Call(__e, tmp102734, sym_1_1_6, tmp102742)

				tmp102744 := Call(__e, tmp102727, tmp102733, tmp102743)

				__e.TailApply(tmp102726, V3536, tmp102744, V3537, V3538)
				return

			}, 1)

			tmp102745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp102746 := Call(__e, tmp102745, V3537)

			__e.TailApply(tmp102723, tmp102746)
			return

		}, 1)

		tmp102747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102748 := Call(__e, tmp102747, V3537)

		__e.TailApply(tmp102722, tmp102748)
		return

	}, 3)

	tmp102749 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1function, tmp102721)

	_ = tmp102749

	tmp102750 := MakeNative(func(__e *ControlFlow) {
		V3546 := __e.Get(1)
		_ = V3546
		V3547 := __e.Get(2)
		_ = V3547
		V3548 := __e.Get(3)
		_ = V3548
		tmp102751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102752 := Call(__e, tmp102751)

		_ = tmp102752

		tmp102753 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102754 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102755 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102757 := Call(__e, tmp102756, symsymbol, Nil)

		tmp102758 := Call(__e, tmp102755, sym_1_1_6, tmp102757)

		tmp102759 := Call(__e, tmp102754, symsymbol, tmp102758)

		__e.TailApply(tmp102753, V3546, tmp102759, V3547, V3548)
		return

	}, 3)

	tmp102760 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1gensym, tmp102750)

	_ = tmp102760

	tmp102761 := MakeNative(func(__e *ControlFlow) {
		V3556 := __e.Get(1)
		_ = V3556
		V3557 := __e.Get(2)
		_ = V3557
		V3558 := __e.Get(3)
		_ = V3558
		tmp102762 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102763 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102764 := Call(__e, tmp102763)

			_ = tmp102764

			tmp102765 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102766 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102767 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102768 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102769 := Call(__e, tmp102768, A, Nil)

			tmp102770 := Call(__e, tmp102767, symvector, tmp102769)

			tmp102771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102772 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102773 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102774 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102775 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102776 := Call(__e, tmp102775, A, Nil)

			tmp102777 := Call(__e, tmp102774, sym_1_1_6, tmp102776)

			tmp102778 := Call(__e, tmp102773, symnumber, tmp102777)

			tmp102779 := Call(__e, tmp102772, tmp102778, Nil)

			tmp102780 := Call(__e, tmp102771, sym_1_1_6, tmp102779)

			tmp102781 := Call(__e, tmp102766, tmp102770, tmp102780)

			__e.TailApply(tmp102765, V3556, tmp102781, V3557, V3558)
			return

		}, 1)

		tmp102782 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102783 := Call(__e, tmp102782, V3557)

		__e.TailApply(tmp102762, tmp102783)
		return

	}, 3)

	tmp102784 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_1vector, tmp102761)

	_ = tmp102784

	tmp102785 := MakeNative(func(__e *ControlFlow) {
		V3566 := __e.Get(1)
		_ = V3566
		V3567 := __e.Get(2)
		_ = V3567
		V3568 := __e.Get(3)
		_ = V3568
		tmp102786 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102788 := Call(__e, tmp102787)

			_ = tmp102788

			tmp102789 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102790 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102791 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102792 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102793 := Call(__e, tmp102792, A, Nil)

			tmp102794 := Call(__e, tmp102791, symvector, tmp102793)

			tmp102795 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102796 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102797 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102798 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102799 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102800 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102801 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102802 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102803 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102804 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102805 := Call(__e, tmp102804, A, Nil)

			tmp102806 := Call(__e, tmp102803, symvector, tmp102805)

			tmp102807 := Call(__e, tmp102802, tmp102806, Nil)

			tmp102808 := Call(__e, tmp102801, sym_1_1_6, tmp102807)

			tmp102809 := Call(__e, tmp102800, A, tmp102808)

			tmp102810 := Call(__e, tmp102799, tmp102809, Nil)

			tmp102811 := Call(__e, tmp102798, sym_1_1_6, tmp102810)

			tmp102812 := Call(__e, tmp102797, symnumber, tmp102811)

			tmp102813 := Call(__e, tmp102796, tmp102812, Nil)

			tmp102814 := Call(__e, tmp102795, sym_1_1_6, tmp102813)

			tmp102815 := Call(__e, tmp102790, tmp102794, tmp102814)

			__e.TailApply(tmp102789, V3566, tmp102815, V3567, V3568)
			return

		}, 1)

		tmp102816 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102817 := Call(__e, tmp102816, V3567)

		__e.TailApply(tmp102786, tmp102817)
		return

	}, 3)

	tmp102818 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector_1_6, tmp102785)

	_ = tmp102818

	tmp102819 := MakeNative(func(__e *ControlFlow) {
		V3576 := __e.Get(1)
		_ = V3576
		V3577 := __e.Get(2)
		_ = V3577
		V3578 := __e.Get(3)
		_ = V3578
		tmp102820 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102821 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102822 := Call(__e, tmp102821)

			_ = tmp102822

			tmp102823 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102824 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102825 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102826 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102827 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102828 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102829 := Call(__e, tmp102828, A, Nil)

			tmp102830 := Call(__e, tmp102827, symvector, tmp102829)

			tmp102831 := Call(__e, tmp102826, tmp102830, Nil)

			tmp102832 := Call(__e, tmp102825, sym_1_1_6, tmp102831)

			tmp102833 := Call(__e, tmp102824, symnumber, tmp102832)

			__e.TailApply(tmp102823, V3576, tmp102833, V3577, V3578)
			return

		}, 1)

		tmp102834 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102835 := Call(__e, tmp102834, V3577)

		__e.TailApply(tmp102820, tmp102835)
		return

	}, 3)

	tmp102836 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector, tmp102819)

	_ = tmp102836

	tmp102837 := MakeNative(func(__e *ControlFlow) {
		V3586 := __e.Get(1)
		_ = V3586
		V3587 := __e.Get(2)
		_ = V3587
		V3588 := __e.Get(3)
		_ = V3588
		tmp102838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102839 := Call(__e, tmp102838)

		_ = tmp102839

		tmp102840 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102841 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102842 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102843 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102844 := Call(__e, tmp102843, symnumber, Nil)

		tmp102845 := Call(__e, tmp102842, sym_1_1_6, tmp102844)

		tmp102846 := Call(__e, tmp102841, symsymbol, tmp102845)

		__e.TailApply(tmp102840, V3586, tmp102846, V3587, V3588)
		return

	}, 3)

	tmp102847 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1get_1time, tmp102837)

	_ = tmp102847

	tmp102848 := MakeNative(func(__e *ControlFlow) {
		V3596 := __e.Get(1)
		_ = V3596
		V3597 := __e.Get(2)
		_ = V3597
		V3598 := __e.Get(3)
		_ = V3598
		tmp102849 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102850 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102851 := Call(__e, tmp102850)

			_ = tmp102851

			tmp102852 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102853 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102854 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102855 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102856 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102857 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102858 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102859 := Call(__e, tmp102858, symnumber, Nil)

			tmp102860 := Call(__e, tmp102857, sym_1_1_6, tmp102859)

			tmp102861 := Call(__e, tmp102856, symnumber, tmp102860)

			tmp102862 := Call(__e, tmp102855, tmp102861, Nil)

			tmp102863 := Call(__e, tmp102854, sym_1_1_6, tmp102862)

			tmp102864 := Call(__e, tmp102853, A, tmp102863)

			__e.TailApply(tmp102852, V3596, tmp102864, V3597, V3598)
			return

		}, 1)

		tmp102865 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102866 := Call(__e, tmp102865, V3597)

		__e.TailApply(tmp102849, tmp102866)
		return

	}, 3)

	tmp102867 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hash, tmp102848)

	_ = tmp102867

	tmp102868 := MakeNative(func(__e *ControlFlow) {
		V3606 := __e.Get(1)
		_ = V3606
		V3607 := __e.Get(2)
		_ = V3607
		V3608 := __e.Get(3)
		_ = V3608
		tmp102869 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102870 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102871 := Call(__e, tmp102870)

			_ = tmp102871

			tmp102872 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102874 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102876 := Call(__e, tmp102875, A, Nil)

			tmp102877 := Call(__e, tmp102874, symlist, tmp102876)

			tmp102878 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102879 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102880 := Call(__e, tmp102879, A, Nil)

			tmp102881 := Call(__e, tmp102878, sym_1_1_6, tmp102880)

			tmp102882 := Call(__e, tmp102873, tmp102877, tmp102881)

			__e.TailApply(tmp102872, V3606, tmp102882, V3607, V3608)
			return

		}, 1)

		tmp102883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102884 := Call(__e, tmp102883, V3607)

		__e.TailApply(tmp102869, tmp102884)
		return

	}, 3)

	tmp102885 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1head, tmp102868)

	_ = tmp102885

	tmp102886 := MakeNative(func(__e *ControlFlow) {
		V3616 := __e.Get(1)
		_ = V3616
		V3617 := __e.Get(2)
		_ = V3617
		V3618 := __e.Get(3)
		_ = V3618
		tmp102887 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102888 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102889 := Call(__e, tmp102888)

			_ = tmp102889

			tmp102890 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102891 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102892 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102894 := Call(__e, tmp102893, A, Nil)

			tmp102895 := Call(__e, tmp102892, symvector, tmp102894)

			tmp102896 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102897 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102898 := Call(__e, tmp102897, A, Nil)

			tmp102899 := Call(__e, tmp102896, sym_1_1_6, tmp102898)

			tmp102900 := Call(__e, tmp102891, tmp102895, tmp102899)

			__e.TailApply(tmp102890, V3616, tmp102900, V3617, V3618)
			return

		}, 1)

		tmp102901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102902 := Call(__e, tmp102901, V3617)

		__e.TailApply(tmp102887, tmp102902)
		return

	}, 3)

	tmp102903 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hdv, tmp102886)

	_ = tmp102903

	tmp102904 := MakeNative(func(__e *ControlFlow) {
		V3626 := __e.Get(1)
		_ = V3626
		V3627 := __e.Get(2)
		_ = V3627
		V3628 := __e.Get(3)
		_ = V3628
		tmp102905 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102906 := Call(__e, tmp102905)

		_ = tmp102906

		tmp102907 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102909 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102910 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102911 := Call(__e, tmp102910, symstring, Nil)

		tmp102912 := Call(__e, tmp102909, sym_1_1_6, tmp102911)

		tmp102913 := Call(__e, tmp102908, symstring, tmp102912)

		__e.TailApply(tmp102907, V3626, tmp102913, V3627, V3628)
		return

	}, 3)

	tmp102914 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hdstr, tmp102904)

	_ = tmp102914

	tmp102915 := MakeNative(func(__e *ControlFlow) {
		V3636 := __e.Get(1)
		_ = V3636
		V3637 := __e.Get(2)
		_ = V3637
		V3638 := __e.Get(3)
		_ = V3638
		tmp102916 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp102917 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp102918 := Call(__e, tmp102917)

			_ = tmp102918

			tmp102919 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp102920 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102921 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102922 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102926 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102927 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp102929 := Call(__e, tmp102928, A, Nil)

			tmp102930 := Call(__e, tmp102927, sym_1_1_6, tmp102929)

			tmp102931 := Call(__e, tmp102926, A, tmp102930)

			tmp102932 := Call(__e, tmp102925, tmp102931, Nil)

			tmp102933 := Call(__e, tmp102924, sym_1_1_6, tmp102932)

			tmp102934 := Call(__e, tmp102923, A, tmp102933)

			tmp102935 := Call(__e, tmp102922, tmp102934, Nil)

			tmp102936 := Call(__e, tmp102921, sym_1_1_6, tmp102935)

			tmp102937 := Call(__e, tmp102920, symboolean, tmp102936)

			__e.TailApply(tmp102919, V3636, tmp102937, V3637, V3638)
			return

		}, 1)

		tmp102938 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp102939 := Call(__e, tmp102938, V3637)

		__e.TailApply(tmp102916, tmp102939)
		return

	}, 3)

	tmp102940 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1if, tmp102915)

	_ = tmp102940

	tmp102941 := MakeNative(func(__e *ControlFlow) {
		V3646 := __e.Get(1)
		_ = V3646
		V3647 := __e.Get(2)
		_ = V3647
		V3648 := __e.Get(3)
		_ = V3648
		tmp102942 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102943 := Call(__e, tmp102942)

		_ = tmp102943

		tmp102944 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102945 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102946 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102947 := Call(__e, tmp102946, symstring, Nil)

		tmp102948 := Call(__e, tmp102945, sym_1_1_6, tmp102947)

		__e.TailApply(tmp102944, V3646, tmp102948, V3647, V3648)
		return

	}, 3)

	tmp102949 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1it, tmp102941)

	_ = tmp102949

	tmp102950 := MakeNative(func(__e *ControlFlow) {
		V3656 := __e.Get(1)
		_ = V3656
		V3657 := __e.Get(2)
		_ = V3657
		V3658 := __e.Get(3)
		_ = V3658
		tmp102951 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102952 := Call(__e, tmp102951)

		_ = tmp102952

		tmp102953 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102954 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102955 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102956 := Call(__e, tmp102955, symstring, Nil)

		tmp102957 := Call(__e, tmp102954, sym_1_1_6, tmp102956)

		__e.TailApply(tmp102953, V3656, tmp102957, V3657, V3658)
		return

	}, 3)

	tmp102958 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1implementation, tmp102950)

	_ = tmp102958

	tmp102959 := MakeNative(func(__e *ControlFlow) {
		V3666 := __e.Get(1)
		_ = V3666
		V3667 := __e.Get(2)
		_ = V3667
		V3668 := __e.Get(3)
		_ = V3668
		tmp102960 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102961 := Call(__e, tmp102960)

		_ = tmp102961

		tmp102962 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102963 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102964 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102965 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102966 := Call(__e, tmp102965, symsymbol, Nil)

		tmp102967 := Call(__e, tmp102964, symlist, tmp102966)

		tmp102968 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102969 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102970 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102971 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102972 := Call(__e, tmp102971, symsymbol, Nil)

		tmp102973 := Call(__e, tmp102970, symlist, tmp102972)

		tmp102974 := Call(__e, tmp102969, tmp102973, Nil)

		tmp102975 := Call(__e, tmp102968, sym_1_1_6, tmp102974)

		tmp102976 := Call(__e, tmp102963, tmp102967, tmp102975)

		__e.TailApply(tmp102962, V3666, tmp102976, V3667, V3668)
		return

	}, 3)

	tmp102977 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1include, tmp102959)

	_ = tmp102977

	tmp102978 := MakeNative(func(__e *ControlFlow) {
		V3676 := __e.Get(1)
		_ = V3676
		V3677 := __e.Get(2)
		_ = V3677
		V3678 := __e.Get(3)
		_ = V3678
		tmp102979 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102980 := Call(__e, tmp102979)

		_ = tmp102980

		tmp102981 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp102982 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102985 := Call(__e, tmp102984, symsymbol, Nil)

		tmp102986 := Call(__e, tmp102983, symlist, tmp102985)

		tmp102987 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102988 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp102991 := Call(__e, tmp102990, symsymbol, Nil)

		tmp102992 := Call(__e, tmp102989, symlist, tmp102991)

		tmp102993 := Call(__e, tmp102988, tmp102992, Nil)

		tmp102994 := Call(__e, tmp102987, sym_1_1_6, tmp102993)

		tmp102995 := Call(__e, tmp102982, tmp102986, tmp102994)

		__e.TailApply(tmp102981, V3676, tmp102995, V3677, V3678)
		return

	}, 3)

	tmp102996 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1include_1all_1but, tmp102978)

	_ = tmp102996

	tmp102997 := MakeNative(func(__e *ControlFlow) {
		V3686 := __e.Get(1)
		_ = V3686
		V3687 := __e.Get(2)
		_ = V3687
		V3688 := __e.Get(3)
		_ = V3688
		tmp102998 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp102999 := Call(__e, tmp102998)

		_ = tmp102999

		tmp103000 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103001 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103003 := Call(__e, tmp103002, symnumber, Nil)

		tmp103004 := Call(__e, tmp103001, sym_1_1_6, tmp103003)

		__e.TailApply(tmp103000, V3686, tmp103004, V3687, V3688)
		return

	}, 3)

	tmp103005 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1inferences, tmp102997)

	_ = tmp103005

	tmp103006 := MakeNative(func(__e *ControlFlow) {
		V3696 := __e.Get(1)
		_ = V3696
		V3697 := __e.Get(2)
		_ = V3697
		V3698 := __e.Get(3)
		_ = V3698
		tmp103007 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103008 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103009 := Call(__e, tmp103008)

			_ = tmp103009

			tmp103010 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103015 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103017 := Call(__e, tmp103016, symstring, Nil)

			tmp103018 := Call(__e, tmp103015, sym_1_1_6, tmp103017)

			tmp103019 := Call(__e, tmp103014, symstring, tmp103018)

			tmp103020 := Call(__e, tmp103013, tmp103019, Nil)

			tmp103021 := Call(__e, tmp103012, sym_1_1_6, tmp103020)

			tmp103022 := Call(__e, tmp103011, A, tmp103021)

			__e.TailApply(tmp103010, V3696, tmp103022, V3697, V3698)
			return

		}, 1)

		tmp103023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103024 := Call(__e, tmp103023, V3697)

		__e.TailApply(tmp103007, tmp103024)
		return

	}, 3)

	tmp103025 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4insert, tmp103006)

	_ = tmp103025

	tmp103026 := MakeNative(func(__e *ControlFlow) {
		V3706 := __e.Get(1)
		_ = V3706
		V3707 := __e.Get(2)
		_ = V3707
		V3708 := __e.Get(3)
		_ = V3708
		tmp103027 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103028 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103029 := Call(__e, tmp103028)

			_ = tmp103029

			tmp103030 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103031 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103032 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103033 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103034 := Call(__e, tmp103033, symboolean, Nil)

			tmp103035 := Call(__e, tmp103032, sym_1_1_6, tmp103034)

			tmp103036 := Call(__e, tmp103031, A, tmp103035)

			__e.TailApply(tmp103030, V3706, tmp103036, V3707, V3708)
			return

		}, 1)

		tmp103037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103038 := Call(__e, tmp103037, V3707)

		__e.TailApply(tmp103027, tmp103038)
		return

	}, 3)

	tmp103039 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1integer_2, tmp103026)

	_ = tmp103039

	tmp103040 := MakeNative(func(__e *ControlFlow) {
		V3716 := __e.Get(1)
		_ = V3716
		V3717 := __e.Get(2)
		_ = V3717
		V3718 := __e.Get(3)
		_ = V3718
		tmp103041 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103042 := Call(__e, tmp103041)

		_ = tmp103042

		tmp103043 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103047 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103048 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103049 := Call(__e, tmp103048, symsymbol, Nil)

		tmp103050 := Call(__e, tmp103047, symlist, tmp103049)

		tmp103051 := Call(__e, tmp103046, tmp103050, Nil)

		tmp103052 := Call(__e, tmp103045, sym_1_1_6, tmp103051)

		tmp103053 := Call(__e, tmp103044, symsymbol, tmp103052)

		__e.TailApply(tmp103043, V3716, tmp103053, V3717, V3718)
		return

	}, 3)

	tmp103054 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1internal, tmp103040)

	_ = tmp103054

	tmp103055 := MakeNative(func(__e *ControlFlow) {
		V3726 := __e.Get(1)
		_ = V3726
		V3727 := __e.Get(2)
		_ = V3727
		V3728 := __e.Get(3)
		_ = V3728
		tmp103056 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103058 := Call(__e, tmp103057)

			_ = tmp103058

			tmp103059 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103060 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103061 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103063 := Call(__e, tmp103062, A, Nil)

			tmp103064 := Call(__e, tmp103061, symlist, tmp103063)

			tmp103065 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103066 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103067 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103068 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103069 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103070 := Call(__e, tmp103069, A, Nil)

			tmp103071 := Call(__e, tmp103068, symlist, tmp103070)

			tmp103072 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103073 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103074 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103075 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103076 := Call(__e, tmp103075, A, Nil)

			tmp103077 := Call(__e, tmp103074, symlist, tmp103076)

			tmp103078 := Call(__e, tmp103073, tmp103077, Nil)

			tmp103079 := Call(__e, tmp103072, sym_1_1_6, tmp103078)

			tmp103080 := Call(__e, tmp103067, tmp103071, tmp103079)

			tmp103081 := Call(__e, tmp103066, tmp103080, Nil)

			tmp103082 := Call(__e, tmp103065, sym_1_1_6, tmp103081)

			tmp103083 := Call(__e, tmp103060, tmp103064, tmp103082)

			__e.TailApply(tmp103059, V3726, tmp103083, V3727, V3728)
			return

		}, 1)

		tmp103084 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103085 := Call(__e, tmp103084, V3727)

		__e.TailApply(tmp103056, tmp103085)
		return

	}, 3)

	tmp103086 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1intersection, tmp103055)

	_ = tmp103086

	tmp103087 := MakeNative(func(__e *ControlFlow) {
		V3736 := __e.Get(1)
		_ = V3736
		V3737 := __e.Get(2)
		_ = V3737
		V3738 := __e.Get(3)
		_ = V3738
		tmp103088 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103089 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103090 := Call(__e, tmp103089)

			_ = tmp103090

			tmp103091 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103092 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103093 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103094 := Call(__e, tmp103093, A, Nil)

			tmp103095 := Call(__e, tmp103092, sym_1_1_6, tmp103094)

			__e.TailApply(tmp103091, V3736, tmp103095, V3737, V3738)
			return

		}, 1)

		tmp103096 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103097 := Call(__e, tmp103096, V3737)

		__e.TailApply(tmp103088, tmp103097)
		return

	}, 3)

	tmp103098 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1kill, tmp103087)

	_ = tmp103098

	tmp103099 := MakeNative(func(__e *ControlFlow) {
		V3746 := __e.Get(1)
		_ = V3746
		V3747 := __e.Get(2)
		_ = V3747
		V3748 := __e.Get(3)
		_ = V3748
		tmp103100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103101 := Call(__e, tmp103100)

		_ = tmp103101

		tmp103102 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103103 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103105 := Call(__e, tmp103104, symstring, Nil)

		tmp103106 := Call(__e, tmp103103, sym_1_1_6, tmp103105)

		__e.TailApply(tmp103102, V3746, tmp103106, V3747, V3748)
		return

	}, 3)

	tmp103107 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1language, tmp103099)

	_ = tmp103107

	tmp103108 := MakeNative(func(__e *ControlFlow) {
		V3756 := __e.Get(1)
		_ = V3756
		V3757 := __e.Get(2)
		_ = V3757
		V3758 := __e.Get(3)
		_ = V3758
		tmp103109 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103110 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103111 := Call(__e, tmp103110)

			_ = tmp103111

			tmp103112 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103115 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103116 := Call(__e, tmp103115, A, Nil)

			tmp103117 := Call(__e, tmp103114, symlist, tmp103116)

			tmp103118 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103120 := Call(__e, tmp103119, symnumber, Nil)

			tmp103121 := Call(__e, tmp103118, sym_1_1_6, tmp103120)

			tmp103122 := Call(__e, tmp103113, tmp103117, tmp103121)

			__e.TailApply(tmp103112, V3756, tmp103122, V3757, V3758)
			return

		}, 1)

		tmp103123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103124 := Call(__e, tmp103123, V3757)

		__e.TailApply(tmp103109, tmp103124)
		return

	}, 3)

	tmp103125 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1length, tmp103108)

	_ = tmp103125

	tmp103126 := MakeNative(func(__e *ControlFlow) {
		V3766 := __e.Get(1)
		_ = V3766
		V3767 := __e.Get(2)
		_ = V3767
		V3768 := __e.Get(3)
		_ = V3768
		tmp103127 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103128 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103129 := Call(__e, tmp103128)

			_ = tmp103129

			tmp103130 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103133 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103134 := Call(__e, tmp103133, A, Nil)

			tmp103135 := Call(__e, tmp103132, symvector, tmp103134)

			tmp103136 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103137 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103138 := Call(__e, tmp103137, symnumber, Nil)

			tmp103139 := Call(__e, tmp103136, sym_1_1_6, tmp103138)

			tmp103140 := Call(__e, tmp103131, tmp103135, tmp103139)

			__e.TailApply(tmp103130, V3766, tmp103140, V3767, V3768)
			return

		}, 1)

		tmp103141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103142 := Call(__e, tmp103141, V3767)

		__e.TailApply(tmp103127, tmp103142)
		return

	}, 3)

	tmp103143 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1limit, tmp103126)

	_ = tmp103143

	tmp103144 := MakeNative(func(__e *ControlFlow) {
		V3776 := __e.Get(1)
		_ = V3776
		V3777 := __e.Get(2)
		_ = V3777
		V3778 := __e.Get(3)
		_ = V3778
		tmp103145 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103146 := Call(__e, tmp103145)

		_ = tmp103146

		tmp103147 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103148 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103150 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103151 := Call(__e, tmp103150, symsymbol, Nil)

		tmp103152 := Call(__e, tmp103149, sym_1_1_6, tmp103151)

		tmp103153 := Call(__e, tmp103148, symstring, tmp103152)

		__e.TailApply(tmp103147, V3776, tmp103153, V3777, V3778)
		return

	}, 3)

	tmp103154 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1load, tmp103144)

	_ = tmp103154

	tmp103155 := MakeNative(func(__e *ControlFlow) {
		V3786 := __e.Get(1)
		_ = V3786
		V3787 := __e.Get(2)
		_ = V3787
		V3788 := __e.Get(3)
		_ = V3788
		tmp103156 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103157 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp103158 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp103159 := Call(__e, tmp103158)

				_ = tmp103159

				tmp103160 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp103161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103163 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103165 := Call(__e, tmp103164, B, Nil)

				tmp103166 := Call(__e, tmp103163, sym_1_1_6, tmp103165)

				tmp103167 := Call(__e, tmp103162, A, tmp103166)

				tmp103168 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103169 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103170 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103171 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103173 := Call(__e, tmp103172, A, Nil)

				tmp103174 := Call(__e, tmp103171, symlist, tmp103173)

				tmp103175 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103177 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103178 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103179 := Call(__e, tmp103178, B, Nil)

				tmp103180 := Call(__e, tmp103177, symlist, tmp103179)

				tmp103181 := Call(__e, tmp103176, tmp103180, Nil)

				tmp103182 := Call(__e, tmp103175, sym_1_1_6, tmp103181)

				tmp103183 := Call(__e, tmp103170, tmp103174, tmp103182)

				tmp103184 := Call(__e, tmp103169, tmp103183, Nil)

				tmp103185 := Call(__e, tmp103168, sym_1_1_6, tmp103184)

				tmp103186 := Call(__e, tmp103161, tmp103167, tmp103185)

				__e.TailApply(tmp103160, V3786, tmp103186, V3787, V3788)
				return

			}, 1)

			tmp103187 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp103188 := Call(__e, tmp103187, V3787)

			__e.TailApply(tmp103157, tmp103188)
			return

		}, 1)

		tmp103189 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103190 := Call(__e, tmp103189, V3787)

		__e.TailApply(tmp103156, tmp103190)
		return

	}, 3)

	tmp103191 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1map, tmp103155)

	_ = tmp103191

	tmp103192 := MakeNative(func(__e *ControlFlow) {
		V3796 := __e.Get(1)
		_ = V3796
		V3797 := __e.Get(2)
		_ = V3797
		V3798 := __e.Get(3)
		_ = V3798
		tmp103193 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103194 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp103195 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp103196 := Call(__e, tmp103195)

				_ = tmp103196

				tmp103197 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp103198 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103199 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103200 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103201 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103202 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103203 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103204 := Call(__e, tmp103203, B, Nil)

				tmp103205 := Call(__e, tmp103202, symlist, tmp103204)

				tmp103206 := Call(__e, tmp103201, tmp103205, Nil)

				tmp103207 := Call(__e, tmp103200, sym_1_1_6, tmp103206)

				tmp103208 := Call(__e, tmp103199, A, tmp103207)

				tmp103209 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103210 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103211 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103213 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103214 := Call(__e, tmp103213, A, Nil)

				tmp103215 := Call(__e, tmp103212, symlist, tmp103214)

				tmp103216 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103217 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103218 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103219 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103220 := Call(__e, tmp103219, B, Nil)

				tmp103221 := Call(__e, tmp103218, symlist, tmp103220)

				tmp103222 := Call(__e, tmp103217, tmp103221, Nil)

				tmp103223 := Call(__e, tmp103216, sym_1_1_6, tmp103222)

				tmp103224 := Call(__e, tmp103211, tmp103215, tmp103223)

				tmp103225 := Call(__e, tmp103210, tmp103224, Nil)

				tmp103226 := Call(__e, tmp103209, sym_1_1_6, tmp103225)

				tmp103227 := Call(__e, tmp103198, tmp103208, tmp103226)

				__e.TailApply(tmp103197, V3796, tmp103227, V3797, V3798)
				return

			}, 1)

			tmp103228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp103229 := Call(__e, tmp103228, V3797)

			__e.TailApply(tmp103194, tmp103229)
			return

		}, 1)

		tmp103230 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103231 := Call(__e, tmp103230, V3797)

		__e.TailApply(tmp103193, tmp103231)
		return

	}, 3)

	tmp103232 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1mapcan, tmp103192)

	_ = tmp103232

	tmp103233 := MakeNative(func(__e *ControlFlow) {
		V3806 := __e.Get(1)
		_ = V3806
		V3807 := __e.Get(2)
		_ = V3807
		V3808 := __e.Get(3)
		_ = V3808
		tmp103234 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103235 := Call(__e, tmp103234)

		_ = tmp103235

		tmp103236 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103237 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103238 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103239 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103240 := Call(__e, tmp103239, symnumber, Nil)

		tmp103241 := Call(__e, tmp103238, sym_1_1_6, tmp103240)

		tmp103242 := Call(__e, tmp103237, symnumber, tmp103241)

		__e.TailApply(tmp103236, V3806, tmp103242, V3807, V3808)
		return

	}, 3)

	tmp103243 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1maxinferences, tmp103233)

	_ = tmp103243

	tmp103244 := MakeNative(func(__e *ControlFlow) {
		V3816 := __e.Get(1)
		_ = V3816
		V3817 := __e.Get(2)
		_ = V3817
		V3818 := __e.Get(3)
		_ = V3818
		tmp103245 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103246 := Call(__e, tmp103245)

		_ = tmp103246

		tmp103247 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103248 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103249 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103250 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103251 := Call(__e, tmp103250, symstring, Nil)

		tmp103252 := Call(__e, tmp103249, sym_1_1_6, tmp103251)

		tmp103253 := Call(__e, tmp103248, symnumber, tmp103252)

		__e.TailApply(tmp103247, V3816, tmp103253, V3817, V3818)
		return

	}, 3)

	tmp103254 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1n_1_6string, tmp103244)

	_ = tmp103254

	tmp103255 := MakeNative(func(__e *ControlFlow) {
		V3826 := __e.Get(1)
		_ = V3826
		V3827 := __e.Get(2)
		_ = V3827
		V3828 := __e.Get(3)
		_ = V3828
		tmp103256 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103257 := Call(__e, tmp103256)

		_ = tmp103257

		tmp103258 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103261 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103262 := Call(__e, tmp103261, symnumber, Nil)

		tmp103263 := Call(__e, tmp103260, sym_1_1_6, tmp103262)

		tmp103264 := Call(__e, tmp103259, symnumber, tmp103263)

		__e.TailApply(tmp103258, V3826, tmp103264, V3827, V3828)
		return

	}, 3)

	tmp103265 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1nl, tmp103255)

	_ = tmp103265

	tmp103266 := MakeNative(func(__e *ControlFlow) {
		V3836 := __e.Get(1)
		_ = V3836
		V3837 := __e.Get(2)
		_ = V3837
		V3838 := __e.Get(3)
		_ = V3838
		tmp103267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103268 := Call(__e, tmp103267)

		_ = tmp103268

		tmp103269 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103270 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103271 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103272 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103273 := Call(__e, tmp103272, symboolean, Nil)

		tmp103274 := Call(__e, tmp103271, sym_1_1_6, tmp103273)

		tmp103275 := Call(__e, tmp103270, symboolean, tmp103274)

		__e.TailApply(tmp103269, V3836, tmp103275, V3837, V3838)
		return

	}, 3)

	tmp103276 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1not, tmp103266)

	_ = tmp103276

	tmp103277 := MakeNative(func(__e *ControlFlow) {
		V3846 := __e.Get(1)
		_ = V3846
		V3847 := __e.Get(2)
		_ = V3847
		V3848 := __e.Get(3)
		_ = V3848
		tmp103278 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103279 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103280 := Call(__e, tmp103279)

			_ = tmp103280

			tmp103281 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103282 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103283 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103284 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103287 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103288 := Call(__e, tmp103287, A, Nil)

			tmp103289 := Call(__e, tmp103286, symlist, tmp103288)

			tmp103290 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103291 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103292 := Call(__e, tmp103291, A, Nil)

			tmp103293 := Call(__e, tmp103290, sym_1_1_6, tmp103292)

			tmp103294 := Call(__e, tmp103285, tmp103289, tmp103293)

			tmp103295 := Call(__e, tmp103284, tmp103294, Nil)

			tmp103296 := Call(__e, tmp103283, sym_1_1_6, tmp103295)

			tmp103297 := Call(__e, tmp103282, symnumber, tmp103296)

			__e.TailApply(tmp103281, V3846, tmp103297, V3847, V3848)
			return

		}, 1)

		tmp103298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103299 := Call(__e, tmp103298, V3847)

		__e.TailApply(tmp103278, tmp103299)
		return

	}, 3)

	tmp103300 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1nth, tmp103277)

	_ = tmp103300

	tmp103301 := MakeNative(func(__e *ControlFlow) {
		V3856 := __e.Get(1)
		_ = V3856
		V3857 := __e.Get(2)
		_ = V3857
		V3858 := __e.Get(3)
		_ = V3858
		tmp103302 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103304 := Call(__e, tmp103303)

			_ = tmp103304

			tmp103305 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103306 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103307 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103308 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103309 := Call(__e, tmp103308, symboolean, Nil)

			tmp103310 := Call(__e, tmp103307, sym_1_1_6, tmp103309)

			tmp103311 := Call(__e, tmp103306, A, tmp103310)

			__e.TailApply(tmp103305, V3856, tmp103311, V3857, V3858)
			return

		}, 1)

		tmp103312 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103313 := Call(__e, tmp103312, V3857)

		__e.TailApply(tmp103302, tmp103313)
		return

	}, 3)

	tmp103314 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1number_2, tmp103301)

	_ = tmp103314

	tmp103315 := MakeNative(func(__e *ControlFlow) {
		V3866 := __e.Get(1)
		_ = V3866
		V3867 := __e.Get(2)
		_ = V3867
		V3868 := __e.Get(3)
		_ = V3868
		tmp103316 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103317 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp103318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp103319 := Call(__e, tmp103318)

				_ = tmp103319

				tmp103320 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp103321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103327 := Call(__e, tmp103326, symnumber, Nil)

				tmp103328 := Call(__e, tmp103325, sym_1_1_6, tmp103327)

				tmp103329 := Call(__e, tmp103324, B, tmp103328)

				tmp103330 := Call(__e, tmp103323, tmp103329, Nil)

				tmp103331 := Call(__e, tmp103322, sym_1_1_6, tmp103330)

				tmp103332 := Call(__e, tmp103321, A, tmp103331)

				__e.TailApply(tmp103320, V3866, tmp103332, V3867, V3868)
				return

			}, 1)

			tmp103333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp103334 := Call(__e, tmp103333, V3867)

			__e.TailApply(tmp103317, tmp103334)
			return

		}, 1)

		tmp103335 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103336 := Call(__e, tmp103335, V3867)

		__e.TailApply(tmp103316, tmp103336)
		return

	}, 3)

	tmp103337 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1occurrences, tmp103315)

	_ = tmp103337

	tmp103338 := MakeNative(func(__e *ControlFlow) {
		V3876 := __e.Get(1)
		_ = V3876
		V3877 := __e.Get(2)
		_ = V3877
		V3878 := __e.Get(3)
		_ = V3878
		tmp103339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103340 := Call(__e, tmp103339)

		_ = tmp103340

		tmp103341 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103344 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103345 := Call(__e, tmp103344, symboolean, Nil)

		tmp103346 := Call(__e, tmp103343, sym_1_1_6, tmp103345)

		tmp103347 := Call(__e, tmp103342, symsymbol, tmp103346)

		__e.TailApply(tmp103341, V3876, tmp103347, V3877, V3878)
		return

	}, 3)

	tmp103348 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1occurs_1check, tmp103338)

	_ = tmp103348

	tmp103349 := MakeNative(func(__e *ControlFlow) {
		V3886 := __e.Get(1)
		_ = V3886
		V3887 := __e.Get(2)
		_ = V3887
		V3888 := __e.Get(3)
		_ = V3888
		tmp103350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103351 := Call(__e, tmp103350)

		_ = tmp103351

		tmp103352 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103356 := Call(__e, tmp103355, symboolean, Nil)

		tmp103357 := Call(__e, tmp103354, sym_1_1_6, tmp103356)

		tmp103358 := Call(__e, tmp103353, symsymbol, tmp103357)

		__e.TailApply(tmp103352, V3886, tmp103358, V3887, V3888)
		return

	}, 3)

	tmp103359 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1optimise, tmp103349)

	_ = tmp103359

	tmp103360 := MakeNative(func(__e *ControlFlow) {
		V3896 := __e.Get(1)
		_ = V3896
		V3897 := __e.Get(2)
		_ = V3897
		V3898 := __e.Get(3)
		_ = V3898
		tmp103361 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103362 := Call(__e, tmp103361)

		_ = tmp103362

		tmp103363 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103364 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103365 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103368 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103370 := Call(__e, tmp103369, symboolean, Nil)

		tmp103371 := Call(__e, tmp103368, sym_1_1_6, tmp103370)

		tmp103372 := Call(__e, tmp103367, symboolean, tmp103371)

		tmp103373 := Call(__e, tmp103366, tmp103372, Nil)

		tmp103374 := Call(__e, tmp103365, sym_1_1_6, tmp103373)

		tmp103375 := Call(__e, tmp103364, symboolean, tmp103374)

		__e.TailApply(tmp103363, V3896, tmp103375, V3897, V3898)
		return

	}, 3)

	tmp103376 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1or, tmp103360)

	_ = tmp103376

	tmp103377 := MakeNative(func(__e *ControlFlow) {
		V3906 := __e.Get(1)
		_ = V3906
		V3907 := __e.Get(2)
		_ = V3907
		V3908 := __e.Get(3)
		_ = V3908
		tmp103378 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103379 := Call(__e, tmp103378)

		_ = tmp103379

		tmp103380 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103381 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103382 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103383 := Call(__e, tmp103382, symstring, Nil)

		tmp103384 := Call(__e, tmp103381, sym_1_1_6, tmp103383)

		__e.TailApply(tmp103380, V3906, tmp103384, V3907, V3908)
		return

	}, 3)

	tmp103385 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1os, tmp103377)

	_ = tmp103385

	tmp103386 := MakeNative(func(__e *ControlFlow) {
		V3916 := __e.Get(1)
		_ = V3916
		V3917 := __e.Get(2)
		_ = V3917
		V3918 := __e.Get(3)
		_ = V3918
		tmp103387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103388 := Call(__e, tmp103387)

		_ = tmp103388

		tmp103389 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103391 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103392 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103393 := Call(__e, tmp103392, symboolean, Nil)

		tmp103394 := Call(__e, tmp103391, sym_1_1_6, tmp103393)

		tmp103395 := Call(__e, tmp103390, symsymbol, tmp103394)

		__e.TailApply(tmp103389, V3916, tmp103395, V3917, V3918)
		return

	}, 3)

	tmp103396 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1package_2, tmp103386)

	_ = tmp103396

	tmp103397 := MakeNative(func(__e *ControlFlow) {
		V3926 := __e.Get(1)
		_ = V3926
		V3927 := __e.Get(2)
		_ = V3927
		V3928 := __e.Get(3)
		_ = V3928
		tmp103398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103399 := Call(__e, tmp103398)

		_ = tmp103399

		tmp103400 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103401 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103402 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103403 := Call(__e, tmp103402, symstring, Nil)

		tmp103404 := Call(__e, tmp103401, sym_1_1_6, tmp103403)

		__e.TailApply(tmp103400, V3926, tmp103404, V3927, V3928)
		return

	}, 3)

	tmp103405 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1port, tmp103397)

	_ = tmp103405

	tmp103406 := MakeNative(func(__e *ControlFlow) {
		V3936 := __e.Get(1)
		_ = V3936
		V3937 := __e.Get(2)
		_ = V3937
		V3938 := __e.Get(3)
		_ = V3938
		tmp103407 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103408 := Call(__e, tmp103407)

		_ = tmp103408

		tmp103409 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103410 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103412 := Call(__e, tmp103411, symstring, Nil)

		tmp103413 := Call(__e, tmp103410, sym_1_1_6, tmp103412)

		__e.TailApply(tmp103409, V3936, tmp103413, V3937, V3938)
		return

	}, 3)

	tmp103414 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1porters, tmp103406)

	_ = tmp103414

	tmp103415 := MakeNative(func(__e *ControlFlow) {
		V3946 := __e.Get(1)
		_ = V3946
		V3947 := __e.Get(2)
		_ = V3947
		V3948 := __e.Get(3)
		_ = V3948
		tmp103416 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103417 := Call(__e, tmp103416)

		_ = tmp103417

		tmp103418 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103419 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103420 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103422 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103423 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103424 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103425 := Call(__e, tmp103424, symstring, Nil)

		tmp103426 := Call(__e, tmp103423, sym_1_1_6, tmp103425)

		tmp103427 := Call(__e, tmp103422, symnumber, tmp103426)

		tmp103428 := Call(__e, tmp103421, tmp103427, Nil)

		tmp103429 := Call(__e, tmp103420, sym_1_1_6, tmp103428)

		tmp103430 := Call(__e, tmp103419, symstring, tmp103429)

		__e.TailApply(tmp103418, V3946, tmp103430, V3947, V3948)
		return

	}, 3)

	tmp103431 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1pos, tmp103415)

	_ = tmp103431

	tmp103432 := MakeNative(func(__e *ControlFlow) {
		V3956 := __e.Get(1)
		_ = V3956
		V3957 := __e.Get(2)
		_ = V3957
		V3958 := __e.Get(3)
		_ = V3958
		tmp103433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103434 := Call(__e, tmp103433)

		_ = tmp103434

		tmp103435 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103436 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103437 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103438 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103439 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103440 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103441 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103442 := Call(__e, tmp103441, symout, Nil)

		tmp103443 := Call(__e, tmp103440, symstream, tmp103442)

		tmp103444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103446 := Call(__e, tmp103445, symstring, Nil)

		tmp103447 := Call(__e, tmp103444, sym_1_1_6, tmp103446)

		tmp103448 := Call(__e, tmp103439, tmp103443, tmp103447)

		tmp103449 := Call(__e, tmp103438, tmp103448, Nil)

		tmp103450 := Call(__e, tmp103437, sym_1_1_6, tmp103449)

		tmp103451 := Call(__e, tmp103436, symstring, tmp103450)

		__e.TailApply(tmp103435, V3956, tmp103451, V3957, V3958)
		return

	}, 3)

	tmp103452 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1pr, tmp103432)

	_ = tmp103452

	tmp103453 := MakeNative(func(__e *ControlFlow) {
		V3966 := __e.Get(1)
		_ = V3966
		V3967 := __e.Get(2)
		_ = V3967
		V3968 := __e.Get(3)
		_ = V3968
		tmp103454 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103455 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103456 := Call(__e, tmp103455)

			_ = tmp103456

			tmp103457 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103458 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103459 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103460 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103461 := Call(__e, tmp103460, A, Nil)

			tmp103462 := Call(__e, tmp103459, sym_1_1_6, tmp103461)

			tmp103463 := Call(__e, tmp103458, A, tmp103462)

			__e.TailApply(tmp103457, V3966, tmp103463, V3967, V3968)
			return

		}, 1)

		tmp103464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103465 := Call(__e, tmp103464, V3967)

		__e.TailApply(tmp103454, tmp103465)
		return

	}, 3)

	tmp103466 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1print, tmp103453)

	_ = tmp103466

	tmp103467 := MakeNative(func(__e *ControlFlow) {
		V3976 := __e.Get(1)
		_ = V3976
		V3977 := __e.Get(2)
		_ = V3977
		V3978 := __e.Get(3)
		_ = V3978
		tmp103468 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103469 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp103470 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp103471 := Call(__e, tmp103470)

				_ = tmp103471

				tmp103472 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp103473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103475 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103476 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103477 := Call(__e, tmp103476, B, Nil)

				tmp103478 := Call(__e, tmp103475, sym_1_1_6, tmp103477)

				tmp103479 := Call(__e, tmp103474, A, tmp103478)

				tmp103480 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103481 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103482 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103483 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103484 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103485 := Call(__e, tmp103484, B, Nil)

				tmp103486 := Call(__e, tmp103483, sym_1_1_6, tmp103485)

				tmp103487 := Call(__e, tmp103482, A, tmp103486)

				tmp103488 := Call(__e, tmp103481, tmp103487, Nil)

				tmp103489 := Call(__e, tmp103480, sym_1_1_6, tmp103488)

				tmp103490 := Call(__e, tmp103473, tmp103479, tmp103489)

				__e.TailApply(tmp103472, V3976, tmp103490, V3977, V3978)
				return

			}, 1)

			tmp103491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp103492 := Call(__e, tmp103491, V3977)

			__e.TailApply(tmp103469, tmp103492)
			return

		}, 1)

		tmp103493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103494 := Call(__e, tmp103493, V3977)

		__e.TailApply(tmp103468, tmp103494)
		return

	}, 3)

	tmp103495 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1profile, tmp103467)

	_ = tmp103495

	tmp103496 := MakeNative(func(__e *ControlFlow) {
		V3986 := __e.Get(1)
		_ = V3986
		V3987 := __e.Get(2)
		_ = V3987
		V3988 := __e.Get(3)
		_ = V3988
		tmp103497 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103498 := Call(__e, tmp103497)

		_ = tmp103498

		tmp103499 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103500 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103501 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103502 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103503 := Call(__e, tmp103502, symsymbol, Nil)

		tmp103504 := Call(__e, tmp103501, symlist, tmp103503)

		tmp103505 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103506 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103507 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103508 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103509 := Call(__e, tmp103508, symsymbol, Nil)

		tmp103510 := Call(__e, tmp103507, symlist, tmp103509)

		tmp103511 := Call(__e, tmp103506, tmp103510, Nil)

		tmp103512 := Call(__e, tmp103505, sym_1_1_6, tmp103511)

		tmp103513 := Call(__e, tmp103500, tmp103504, tmp103512)

		__e.TailApply(tmp103499, V3986, tmp103513, V3987, V3988)
		return

	}, 3)

	tmp103514 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1preclude, tmp103496)

	_ = tmp103514

	tmp103515 := MakeNative(func(__e *ControlFlow) {
		V3996 := __e.Get(1)
		_ = V3996
		V3997 := __e.Get(2)
		_ = V3997
		V3998 := __e.Get(3)
		_ = V3998
		tmp103516 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103517 := Call(__e, tmp103516)

		_ = tmp103517

		tmp103518 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103519 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103520 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103521 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103522 := Call(__e, tmp103521, symstring, Nil)

		tmp103523 := Call(__e, tmp103520, sym_1_1_6, tmp103522)

		tmp103524 := Call(__e, tmp103519, symstring, tmp103523)

		__e.TailApply(tmp103518, V3996, tmp103524, V3997, V3998)
		return

	}, 3)

	tmp103525 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4proc_1nl, tmp103515)

	_ = tmp103525

	tmp103526 := MakeNative(func(__e *ControlFlow) {
		V4006 := __e.Get(1)
		_ = V4006
		V4007 := __e.Get(2)
		_ = V4007
		V4008 := __e.Get(3)
		_ = V4008
		tmp103527 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103528 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp103529 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp103530 := Call(__e, tmp103529)

				_ = tmp103530

				tmp103531 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp103532 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103533 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103534 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103535 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103536 := Call(__e, tmp103535, B, Nil)

				tmp103537 := Call(__e, tmp103534, sym_1_1_6, tmp103536)

				tmp103538 := Call(__e, tmp103533, A, tmp103537)

				tmp103539 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103541 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103542 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103543 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103544 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103545 := Call(__e, tmp103544, B, Nil)

				tmp103546 := Call(__e, tmp103543, sym_1_1_6, tmp103545)

				tmp103547 := Call(__e, tmp103542, A, tmp103546)

				tmp103548 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103549 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103550 := Call(__e, tmp103549, symnumber, Nil)

				tmp103551 := Call(__e, tmp103548, sym_d, tmp103550)

				tmp103552 := Call(__e, tmp103541, tmp103547, tmp103551)

				tmp103553 := Call(__e, tmp103540, tmp103552, Nil)

				tmp103554 := Call(__e, tmp103539, sym_1_1_6, tmp103553)

				tmp103555 := Call(__e, tmp103532, tmp103538, tmp103554)

				__e.TailApply(tmp103531, V4006, tmp103555, V4007, V4008)
				return

			}, 1)

			tmp103556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp103557 := Call(__e, tmp103556, V4007)

			__e.TailApply(tmp103528, tmp103557)
			return

		}, 1)

		tmp103558 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103559 := Call(__e, tmp103558, V4007)

		__e.TailApply(tmp103527, tmp103559)
		return

	}, 3)

	tmp103560 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1profile_1results, tmp103526)

	_ = tmp103560

	tmp103561 := MakeNative(func(__e *ControlFlow) {
		V4016 := __e.Get(1)
		_ = V4016
		V4017 := __e.Get(2)
		_ = V4017
		V4018 := __e.Get(3)
		_ = V4018
		tmp103562 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103563 := Call(__e, tmp103562)

		_ = tmp103563

		tmp103564 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103565 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103566 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103567 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103568 := Call(__e, tmp103567, symsymbol, Nil)

		tmp103569 := Call(__e, tmp103566, sym_1_1_6, tmp103568)

		tmp103570 := Call(__e, tmp103565, symsymbol, tmp103569)

		__e.TailApply(tmp103564, V4016, tmp103570, V4017, V4018)
		return

	}, 3)

	tmp103571 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1protect, tmp103561)

	_ = tmp103571

	tmp103572 := MakeNative(func(__e *ControlFlow) {
		V4026 := __e.Get(1)
		_ = V4026
		V4027 := __e.Get(2)
		_ = V4027
		V4028 := __e.Get(3)
		_ = V4028
		tmp103573 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103574 := Call(__e, tmp103573)

		_ = tmp103574

		tmp103575 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103576 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103577 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103578 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103579 := Call(__e, tmp103578, symsymbol, Nil)

		tmp103580 := Call(__e, tmp103577, symlist, tmp103579)

		tmp103581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103583 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103584 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103585 := Call(__e, tmp103584, symsymbol, Nil)

		tmp103586 := Call(__e, tmp103583, symlist, tmp103585)

		tmp103587 := Call(__e, tmp103582, tmp103586, Nil)

		tmp103588 := Call(__e, tmp103581, sym_1_1_6, tmp103587)

		tmp103589 := Call(__e, tmp103576, tmp103580, tmp103588)

		__e.TailApply(tmp103575, V4026, tmp103589, V4027, V4028)
		return

	}, 3)

	tmp103590 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1preclude_1all_1but, tmp103572)

	_ = tmp103590

	tmp103591 := MakeNative(func(__e *ControlFlow) {
		V4036 := __e.Get(1)
		_ = V4036
		V4037 := __e.Get(2)
		_ = V4037
		V4038 := __e.Get(3)
		_ = V4038
		tmp103592 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103593 := Call(__e, tmp103592)

		_ = tmp103593

		tmp103594 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103595 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103596 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103601 := Call(__e, tmp103600, symout, Nil)

		tmp103602 := Call(__e, tmp103599, symstream, tmp103601)

		tmp103603 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103604 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103605 := Call(__e, tmp103604, symstring, Nil)

		tmp103606 := Call(__e, tmp103603, sym_1_1_6, tmp103605)

		tmp103607 := Call(__e, tmp103598, tmp103602, tmp103606)

		tmp103608 := Call(__e, tmp103597, tmp103607, Nil)

		tmp103609 := Call(__e, tmp103596, sym_1_1_6, tmp103608)

		tmp103610 := Call(__e, tmp103595, symstring, tmp103609)

		__e.TailApply(tmp103594, V4036, tmp103610, V4037, V4038)
		return

	}, 3)

	tmp103611 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4prhush, tmp103591)

	_ = tmp103611

	tmp103612 := MakeNative(func(__e *ControlFlow) {
		V4046 := __e.Get(1)
		_ = V4046
		V4047 := __e.Get(2)
		_ = V4047
		V4048 := __e.Get(3)
		_ = V4048
		tmp103613 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103614 := Call(__e, tmp103613)

		_ = tmp103614

		tmp103615 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103616 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103617 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103618 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103619 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103620 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103621 := Call(__e, tmp103620, symunit, Nil)

		tmp103622 := Call(__e, tmp103619, symlist, tmp103621)

		tmp103623 := Call(__e, tmp103618, tmp103622, Nil)

		tmp103624 := Call(__e, tmp103617, sym_1_1_6, tmp103623)

		tmp103625 := Call(__e, tmp103616, symsymbol, tmp103624)

		__e.TailApply(tmp103615, V4046, tmp103625, V4047, V4048)
		return

	}, 3)

	tmp103626 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1ps, tmp103612)

	_ = tmp103626

	tmp103627 := MakeNative(func(__e *ControlFlow) {
		V4056 := __e.Get(1)
		_ = V4056
		V4057 := __e.Get(2)
		_ = V4057
		V4058 := __e.Get(3)
		_ = V4058
		tmp103628 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103629 := Call(__e, tmp103628)

		_ = tmp103629

		tmp103630 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103631 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103632 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103633 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103634 := Call(__e, tmp103633, symin, Nil)

		tmp103635 := Call(__e, tmp103632, symstream, tmp103634)

		tmp103636 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103637 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103638 := Call(__e, tmp103637, symunit, Nil)

		tmp103639 := Call(__e, tmp103636, sym_1_1_6, tmp103638)

		tmp103640 := Call(__e, tmp103631, tmp103635, tmp103639)

		__e.TailApply(tmp103630, V4056, tmp103640, V4057, V4058)
		return

	}, 3)

	tmp103641 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read, tmp103627)

	_ = tmp103641

	tmp103642 := MakeNative(func(__e *ControlFlow) {
		V4066 := __e.Get(1)
		_ = V4066
		V4067 := __e.Get(2)
		_ = V4067
		V4068 := __e.Get(3)
		_ = V4068
		tmp103643 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103644 := Call(__e, tmp103643)

		_ = tmp103644

		tmp103645 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103646 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103647 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103648 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103649 := Call(__e, tmp103648, symin, Nil)

		tmp103650 := Call(__e, tmp103647, symstream, tmp103649)

		tmp103651 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103652 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103653 := Call(__e, tmp103652, symnumber, Nil)

		tmp103654 := Call(__e, tmp103651, sym_1_1_6, tmp103653)

		tmp103655 := Call(__e, tmp103646, tmp103650, tmp103654)

		__e.TailApply(tmp103645, V4066, tmp103655, V4067, V4068)
		return

	}, 3)

	tmp103656 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1byte, tmp103642)

	_ = tmp103656

	tmp103657 := MakeNative(func(__e *ControlFlow) {
		V4076 := __e.Get(1)
		_ = V4076
		V4077 := __e.Get(2)
		_ = V4077
		V4078 := __e.Get(3)
		_ = V4078
		tmp103658 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103659 := Call(__e, tmp103658)

		_ = tmp103659

		tmp103660 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103663 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103664 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103666 := Call(__e, tmp103665, symnumber, Nil)

		tmp103667 := Call(__e, tmp103664, symlist, tmp103666)

		tmp103668 := Call(__e, tmp103663, tmp103667, Nil)

		tmp103669 := Call(__e, tmp103662, sym_1_1_6, tmp103668)

		tmp103670 := Call(__e, tmp103661, symstring, tmp103669)

		__e.TailApply(tmp103660, V4076, tmp103670, V4077, V4078)
		return

	}, 3)

	tmp103671 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file_1as_1bytelist, tmp103657)

	_ = tmp103671

	tmp103672 := MakeNative(func(__e *ControlFlow) {
		V4086 := __e.Get(1)
		_ = V4086
		V4087 := __e.Get(2)
		_ = V4087
		V4088 := __e.Get(3)
		_ = V4088
		tmp103673 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103674 := Call(__e, tmp103673)

		_ = tmp103674

		tmp103675 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103679 := Call(__e, tmp103678, symstring, Nil)

		tmp103680 := Call(__e, tmp103677, sym_1_1_6, tmp103679)

		tmp103681 := Call(__e, tmp103676, symstring, tmp103680)

		__e.TailApply(tmp103675, V4086, tmp103681, V4087, V4088)
		return

	}, 3)

	tmp103682 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file_1as_1string, tmp103672)

	_ = tmp103682

	tmp103683 := MakeNative(func(__e *ControlFlow) {
		V4096 := __e.Get(1)
		_ = V4096
		V4097 := __e.Get(2)
		_ = V4097
		V4098 := __e.Get(3)
		_ = V4098
		tmp103684 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103685 := Call(__e, tmp103684)

		_ = tmp103685

		tmp103686 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103688 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103689 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103690 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103692 := Call(__e, tmp103691, symunit, Nil)

		tmp103693 := Call(__e, tmp103690, symlist, tmp103692)

		tmp103694 := Call(__e, tmp103689, tmp103693, Nil)

		tmp103695 := Call(__e, tmp103688, sym_1_1_6, tmp103694)

		tmp103696 := Call(__e, tmp103687, symstring, tmp103695)

		__e.TailApply(tmp103686, V4096, tmp103696, V4097, V4098)
		return

	}, 3)

	tmp103697 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file, tmp103683)

	_ = tmp103697

	tmp103698 := MakeNative(func(__e *ControlFlow) {
		V4106 := __e.Get(1)
		_ = V4106
		V4107 := __e.Get(2)
		_ = V4107
		V4108 := __e.Get(3)
		_ = V4108
		tmp103699 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103700 := Call(__e, tmp103699)

		_ = tmp103700

		tmp103701 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103702 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103703 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103704 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103705 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103707 := Call(__e, tmp103706, symunit, Nil)

		tmp103708 := Call(__e, tmp103705, symlist, tmp103707)

		tmp103709 := Call(__e, tmp103704, tmp103708, Nil)

		tmp103710 := Call(__e, tmp103703, sym_1_1_6, tmp103709)

		tmp103711 := Call(__e, tmp103702, symstring, tmp103710)

		__e.TailApply(tmp103701, V4106, tmp103711, V4107, V4108)
		return

	}, 3)

	tmp103712 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1from_1string, tmp103698)

	_ = tmp103712

	tmp103713 := MakeNative(func(__e *ControlFlow) {
		V4116 := __e.Get(1)
		_ = V4116
		V4117 := __e.Get(2)
		_ = V4117
		V4118 := __e.Get(3)
		_ = V4118
		tmp103714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103715 := Call(__e, tmp103714)

		_ = tmp103715

		tmp103716 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103717 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103718 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103719 := Call(__e, tmp103718, symstring, Nil)

		tmp103720 := Call(__e, tmp103717, sym_1_1_6, tmp103719)

		__e.TailApply(tmp103716, V4116, tmp103720, V4117, V4118)
		return

	}, 3)

	tmp103721 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1release, tmp103713)

	_ = tmp103721

	tmp103722 := MakeNative(func(__e *ControlFlow) {
		V4126 := __e.Get(1)
		_ = V4126
		V4127 := __e.Get(2)
		_ = V4127
		V4128 := __e.Get(3)
		_ = V4128
		tmp103723 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103725 := Call(__e, tmp103724)

			_ = tmp103725

			tmp103726 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103727 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103730 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103731 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103732 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103733 := Call(__e, tmp103732, A, Nil)

			tmp103734 := Call(__e, tmp103731, symlist, tmp103733)

			tmp103735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103736 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103737 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103739 := Call(__e, tmp103738, A, Nil)

			tmp103740 := Call(__e, tmp103737, symlist, tmp103739)

			tmp103741 := Call(__e, tmp103736, tmp103740, Nil)

			tmp103742 := Call(__e, tmp103735, sym_1_1_6, tmp103741)

			tmp103743 := Call(__e, tmp103730, tmp103734, tmp103742)

			tmp103744 := Call(__e, tmp103729, tmp103743, Nil)

			tmp103745 := Call(__e, tmp103728, sym_1_1_6, tmp103744)

			tmp103746 := Call(__e, tmp103727, A, tmp103745)

			__e.TailApply(tmp103726, V4126, tmp103746, V4127, V4128)
			return

		}, 1)

		tmp103747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103748 := Call(__e, tmp103747, V4127)

		__e.TailApply(tmp103723, tmp103748)
		return

	}, 3)

	tmp103749 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1remove, tmp103722)

	_ = tmp103749

	tmp103750 := MakeNative(func(__e *ControlFlow) {
		V4136 := __e.Get(1)
		_ = V4136
		V4137 := __e.Get(2)
		_ = V4137
		V4138 := __e.Get(3)
		_ = V4138
		tmp103751 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103752 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103753 := Call(__e, tmp103752)

			_ = tmp103753

			tmp103754 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103755 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103758 := Call(__e, tmp103757, A, Nil)

			tmp103759 := Call(__e, tmp103756, symlist, tmp103758)

			tmp103760 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103761 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103762 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103763 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103764 := Call(__e, tmp103763, A, Nil)

			tmp103765 := Call(__e, tmp103762, symlist, tmp103764)

			tmp103766 := Call(__e, tmp103761, tmp103765, Nil)

			tmp103767 := Call(__e, tmp103760, sym_1_1_6, tmp103766)

			tmp103768 := Call(__e, tmp103755, tmp103759, tmp103767)

			__e.TailApply(tmp103754, V4136, tmp103768, V4137, V4138)
			return

		}, 1)

		tmp103769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103770 := Call(__e, tmp103769, V4137)

		__e.TailApply(tmp103751, tmp103770)
		return

	}, 3)

	tmp103771 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1reverse, tmp103750)

	_ = tmp103771

	tmp103772 := MakeNative(func(__e *ControlFlow) {
		V4146 := __e.Get(1)
		_ = V4146
		V4147 := __e.Get(2)
		_ = V4147
		V4148 := __e.Get(3)
		_ = V4148
		tmp103773 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103774 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103775 := Call(__e, tmp103774)

			_ = tmp103775

			tmp103776 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103777 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103778 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103779 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103780 := Call(__e, tmp103779, A, Nil)

			tmp103781 := Call(__e, tmp103778, sym_1_1_6, tmp103780)

			tmp103782 := Call(__e, tmp103777, symstring, tmp103781)

			__e.TailApply(tmp103776, V4146, tmp103782, V4147, V4148)
			return

		}, 1)

		tmp103783 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103784 := Call(__e, tmp103783, V4147)

		__e.TailApply(tmp103773, tmp103784)
		return

	}, 3)

	tmp103785 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1simple_1error, tmp103772)

	_ = tmp103785

	tmp103786 := MakeNative(func(__e *ControlFlow) {
		V4156 := __e.Get(1)
		_ = V4156
		V4157 := __e.Get(2)
		_ = V4157
		V4158 := __e.Get(3)
		_ = V4158
		tmp103787 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103788 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp103789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp103790 := Call(__e, tmp103789)

				_ = tmp103790

				tmp103791 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp103792 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103794 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103795 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103796 := Call(__e, tmp103795, B, Nil)

				tmp103797 := Call(__e, tmp103794, sym_d, tmp103796)

				tmp103798 := Call(__e, tmp103793, A, tmp103797)

				tmp103799 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103800 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp103801 := Call(__e, tmp103800, B, Nil)

				tmp103802 := Call(__e, tmp103799, sym_1_1_6, tmp103801)

				tmp103803 := Call(__e, tmp103792, tmp103798, tmp103802)

				__e.TailApply(tmp103791, V4156, tmp103803, V4157, V4158)
				return

			}, 1)

			tmp103804 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp103805 := Call(__e, tmp103804, V4157)

			__e.TailApply(tmp103788, tmp103805)
			return

		}, 1)

		tmp103806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103807 := Call(__e, tmp103806, V4157)

		__e.TailApply(tmp103787, tmp103807)
		return

	}, 3)

	tmp103808 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1snd, tmp103786)

	_ = tmp103808

	tmp103809 := MakeNative(func(__e *ControlFlow) {
		V4166 := __e.Get(1)
		_ = V4166
		V4167 := __e.Get(2)
		_ = V4167
		V4168 := __e.Get(3)
		_ = V4168
		tmp103810 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103811 := Call(__e, tmp103810)

		_ = tmp103811

		tmp103812 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103813 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103814 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103815 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103816 := Call(__e, tmp103815, symsymbol, Nil)

		tmp103817 := Call(__e, tmp103814, sym_1_1_6, tmp103816)

		tmp103818 := Call(__e, tmp103813, symsymbol, tmp103817)

		__e.TailApply(tmp103812, V4166, tmp103818, V4167, V4168)
		return

	}, 3)

	tmp103819 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1specialise, tmp103809)

	_ = tmp103819

	tmp103820 := MakeNative(func(__e *ControlFlow) {
		V4176 := __e.Get(1)
		_ = V4176
		V4177 := __e.Get(2)
		_ = V4177
		V4178 := __e.Get(3)
		_ = V4178
		tmp103821 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103822 := Call(__e, tmp103821)

		_ = tmp103822

		tmp103823 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103824 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103825 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103826 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103827 := Call(__e, tmp103826, symboolean, Nil)

		tmp103828 := Call(__e, tmp103825, sym_1_1_6, tmp103827)

		tmp103829 := Call(__e, tmp103824, symsymbol, tmp103828)

		__e.TailApply(tmp103823, V4176, tmp103829, V4177, V4178)
		return

	}, 3)

	tmp103830 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1spy, tmp103820)

	_ = tmp103830

	tmp103831 := MakeNative(func(__e *ControlFlow) {
		V4186 := __e.Get(1)
		_ = V4186
		V4187 := __e.Get(2)
		_ = V4187
		V4188 := __e.Get(3)
		_ = V4188
		tmp103832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103833 := Call(__e, tmp103832)

		_ = tmp103833

		tmp103834 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103835 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103837 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103838 := Call(__e, tmp103837, symboolean, Nil)

		tmp103839 := Call(__e, tmp103836, sym_1_1_6, tmp103838)

		tmp103840 := Call(__e, tmp103835, symsymbol, tmp103839)

		__e.TailApply(tmp103834, V4186, tmp103840, V4187, V4188)
		return

	}, 3)

	tmp103841 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1step, tmp103831)

	_ = tmp103841

	tmp103842 := MakeNative(func(__e *ControlFlow) {
		V4196 := __e.Get(1)
		_ = V4196
		V4197 := __e.Get(2)
		_ = V4197
		V4198 := __e.Get(3)
		_ = V4198
		tmp103843 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103844 := Call(__e, tmp103843)

		_ = tmp103844

		tmp103845 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103846 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103847 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103848 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103849 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103850 := Call(__e, tmp103849, symin, Nil)

		tmp103851 := Call(__e, tmp103848, symstream, tmp103850)

		tmp103852 := Call(__e, tmp103847, tmp103851, Nil)

		tmp103853 := Call(__e, tmp103846, sym_1_1_6, tmp103852)

		__e.TailApply(tmp103845, V4196, tmp103853, V4197, V4198)
		return

	}, 3)

	tmp103854 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1stinput, tmp103842)

	_ = tmp103854

	tmp103855 := MakeNative(func(__e *ControlFlow) {
		V4206 := __e.Get(1)
		_ = V4206
		V4207 := __e.Get(2)
		_ = V4207
		V4208 := __e.Get(3)
		_ = V4208
		tmp103856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103857 := Call(__e, tmp103856)

		_ = tmp103857

		tmp103858 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103859 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103860 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103861 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103863 := Call(__e, tmp103862, symout, Nil)

		tmp103864 := Call(__e, tmp103861, symstream, tmp103863)

		tmp103865 := Call(__e, tmp103860, tmp103864, Nil)

		tmp103866 := Call(__e, tmp103859, sym_1_1_6, tmp103865)

		__e.TailApply(tmp103858, V4206, tmp103866, V4207, V4208)
		return

	}, 3)

	tmp103867 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1sterror, tmp103855)

	_ = tmp103867

	tmp103868 := MakeNative(func(__e *ControlFlow) {
		V4216 := __e.Get(1)
		_ = V4216
		V4217 := __e.Get(2)
		_ = V4217
		V4218 := __e.Get(3)
		_ = V4218
		tmp103869 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103870 := Call(__e, tmp103869)

		_ = tmp103870

		tmp103871 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103872 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103874 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103876 := Call(__e, tmp103875, symout, Nil)

		tmp103877 := Call(__e, tmp103874, symstream, tmp103876)

		tmp103878 := Call(__e, tmp103873, tmp103877, Nil)

		tmp103879 := Call(__e, tmp103872, sym_1_1_6, tmp103878)

		__e.TailApply(tmp103871, V4216, tmp103879, V4217, V4218)
		return

	}, 3)

	tmp103880 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1stoutput, tmp103868)

	_ = tmp103880

	tmp103881 := MakeNative(func(__e *ControlFlow) {
		V4226 := __e.Get(1)
		_ = V4226
		V4227 := __e.Get(2)
		_ = V4227
		V4228 := __e.Get(3)
		_ = V4228
		tmp103882 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103884 := Call(__e, tmp103883)

			_ = tmp103884

			tmp103885 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103886 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103887 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103888 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103889 := Call(__e, tmp103888, symboolean, Nil)

			tmp103890 := Call(__e, tmp103887, sym_1_1_6, tmp103889)

			tmp103891 := Call(__e, tmp103886, A, tmp103890)

			__e.TailApply(tmp103885, V4226, tmp103891, V4227, V4228)
			return

		}, 1)

		tmp103892 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103893 := Call(__e, tmp103892, V4227)

		__e.TailApply(tmp103882, tmp103893)
		return

	}, 3)

	tmp103894 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_2, tmp103881)

	_ = tmp103894

	tmp103895 := MakeNative(func(__e *ControlFlow) {
		V4236 := __e.Get(1)
		_ = V4236
		V4237 := __e.Get(2)
		_ = V4237
		V4238 := __e.Get(3)
		_ = V4238
		tmp103896 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103897 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103898 := Call(__e, tmp103897)

			_ = tmp103898

			tmp103899 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103900 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103901 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103902 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103903 := Call(__e, tmp103902, symstring, Nil)

			tmp103904 := Call(__e, tmp103901, sym_1_1_6, tmp103903)

			tmp103905 := Call(__e, tmp103900, A, tmp103904)

			__e.TailApply(tmp103899, V4236, tmp103905, V4237, V4238)
			return

		}, 1)

		tmp103906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103907 := Call(__e, tmp103906, V4237)

		__e.TailApply(tmp103896, tmp103907)
		return

	}, 3)

	tmp103908 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1str, tmp103895)

	_ = tmp103908

	tmp103909 := MakeNative(func(__e *ControlFlow) {
		V4246 := __e.Get(1)
		_ = V4246
		V4247 := __e.Get(2)
		_ = V4247
		V4248 := __e.Get(3)
		_ = V4248
		tmp103910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103911 := Call(__e, tmp103910)

		_ = tmp103911

		tmp103912 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103913 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103914 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103915 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103916 := Call(__e, tmp103915, symnumber, Nil)

		tmp103917 := Call(__e, tmp103914, sym_1_1_6, tmp103916)

		tmp103918 := Call(__e, tmp103913, symstring, tmp103917)

		__e.TailApply(tmp103912, V4246, tmp103918, V4247, V4248)
		return

	}, 3)

	tmp103919 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_1_6n, tmp103909)

	_ = tmp103919

	tmp103920 := MakeNative(func(__e *ControlFlow) {
		V4256 := __e.Get(1)
		_ = V4256
		V4257 := __e.Get(2)
		_ = V4257
		V4258 := __e.Get(3)
		_ = V4258
		tmp103921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103922 := Call(__e, tmp103921)

		_ = tmp103922

		tmp103923 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103926 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103927 := Call(__e, tmp103926, symsymbol, Nil)

		tmp103928 := Call(__e, tmp103925, sym_1_1_6, tmp103927)

		tmp103929 := Call(__e, tmp103924, symstring, tmp103928)

		__e.TailApply(tmp103923, V4256, tmp103929, V4257, V4258)
		return

	}, 3)

	tmp103930 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_1_6symbol, tmp103920)

	_ = tmp103930

	tmp103931 := MakeNative(func(__e *ControlFlow) {
		V4266 := __e.Get(1)
		_ = V4266
		V4267 := __e.Get(2)
		_ = V4267
		V4268 := __e.Get(3)
		_ = V4268
		tmp103932 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103933 := Call(__e, tmp103932)

		_ = tmp103933

		tmp103934 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103938 := Call(__e, tmp103937, symnumber, Nil)

		tmp103939 := Call(__e, tmp103936, symlist, tmp103938)

		tmp103940 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103941 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103942 := Call(__e, tmp103941, symnumber, Nil)

		tmp103943 := Call(__e, tmp103940, sym_1_1_6, tmp103942)

		tmp103944 := Call(__e, tmp103935, tmp103939, tmp103943)

		__e.TailApply(tmp103934, V4266, tmp103944, V4267, V4268)
		return

	}, 3)

	tmp103945 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1sum, tmp103931)

	_ = tmp103945

	tmp103946 := MakeNative(func(__e *ControlFlow) {
		V4276 := __e.Get(1)
		_ = V4276
		V4277 := __e.Get(2)
		_ = V4277
		V4278 := __e.Get(3)
		_ = V4278
		tmp103947 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103948 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103949 := Call(__e, tmp103948)

			_ = tmp103949

			tmp103950 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103951 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103952 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103953 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103954 := Call(__e, tmp103953, symboolean, Nil)

			tmp103955 := Call(__e, tmp103952, sym_1_1_6, tmp103954)

			tmp103956 := Call(__e, tmp103951, A, tmp103955)

			__e.TailApply(tmp103950, V4276, tmp103956, V4277, V4278)
			return

		}, 1)

		tmp103957 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103958 := Call(__e, tmp103957, V4277)

		__e.TailApply(tmp103947, tmp103958)
		return

	}, 3)

	tmp103959 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1symbol_2, tmp103946)

	_ = tmp103959

	tmp103960 := MakeNative(func(__e *ControlFlow) {
		V4286 := __e.Get(1)
		_ = V4286
		V4287 := __e.Get(2)
		_ = V4287
		V4288 := __e.Get(3)
		_ = V4288
		tmp103961 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103962 := Call(__e, tmp103961)

		_ = tmp103962

		tmp103963 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103964 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103965 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103966 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103967 := Call(__e, tmp103966, symsymbol, Nil)

		tmp103968 := Call(__e, tmp103965, sym_1_1_6, tmp103967)

		tmp103969 := Call(__e, tmp103964, symsymbol, tmp103968)

		__e.TailApply(tmp103963, V4286, tmp103969, V4287, V4288)
		return

	}, 3)

	tmp103970 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1systemf, tmp103960)

	_ = tmp103970

	tmp103971 := MakeNative(func(__e *ControlFlow) {
		V4296 := __e.Get(1)
		_ = V4296
		V4297 := __e.Get(2)
		_ = V4297
		V4298 := __e.Get(3)
		_ = V4298
		tmp103972 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp103973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp103974 := Call(__e, tmp103973)

			_ = tmp103974

			tmp103975 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp103976 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103977 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103979 := Call(__e, tmp103978, A, Nil)

			tmp103980 := Call(__e, tmp103977, symlist, tmp103979)

			tmp103981 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103982 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp103985 := Call(__e, tmp103984, A, Nil)

			tmp103986 := Call(__e, tmp103983, symlist, tmp103985)

			tmp103987 := Call(__e, tmp103982, tmp103986, Nil)

			tmp103988 := Call(__e, tmp103981, sym_1_1_6, tmp103987)

			tmp103989 := Call(__e, tmp103976, tmp103980, tmp103988)

			__e.TailApply(tmp103975, V4296, tmp103989, V4297, V4298)
			return

		}, 1)

		tmp103990 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp103991 := Call(__e, tmp103990, V4297)

		__e.TailApply(tmp103972, tmp103991)
		return

	}, 3)

	tmp103992 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tail, tmp103971)

	_ = tmp103992

	tmp103993 := MakeNative(func(__e *ControlFlow) {
		V4306 := __e.Get(1)
		_ = V4306
		V4307 := __e.Get(2)
		_ = V4307
		V4308 := __e.Get(3)
		_ = V4308
		tmp103994 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp103995 := Call(__e, tmp103994)

		_ = tmp103995

		tmp103996 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp103997 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103998 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp103999 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104000 := Call(__e, tmp103999, symstring, Nil)

		tmp104001 := Call(__e, tmp103998, sym_1_1_6, tmp104000)

		tmp104002 := Call(__e, tmp103997, symstring, tmp104001)

		__e.TailApply(tmp103996, V4306, tmp104002, V4307, V4308)
		return

	}, 3)

	tmp104003 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tlstr, tmp103993)

	_ = tmp104003

	tmp104004 := MakeNative(func(__e *ControlFlow) {
		V4316 := __e.Get(1)
		_ = V4316
		V4317 := __e.Get(2)
		_ = V4317
		V4318 := __e.Get(3)
		_ = V4318
		tmp104005 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104006 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104007 := Call(__e, tmp104006)

			_ = tmp104007

			tmp104008 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104010 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104012 := Call(__e, tmp104011, A, Nil)

			tmp104013 := Call(__e, tmp104010, symvector, tmp104012)

			tmp104014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104015 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104018 := Call(__e, tmp104017, A, Nil)

			tmp104019 := Call(__e, tmp104016, symvector, tmp104018)

			tmp104020 := Call(__e, tmp104015, tmp104019, Nil)

			tmp104021 := Call(__e, tmp104014, sym_1_1_6, tmp104020)

			tmp104022 := Call(__e, tmp104009, tmp104013, tmp104021)

			__e.TailApply(tmp104008, V4316, tmp104022, V4317, V4318)
			return

		}, 1)

		tmp104023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104024 := Call(__e, tmp104023, V4317)

		__e.TailApply(tmp104005, tmp104024)
		return

	}, 3)

	tmp104025 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tlv, tmp104004)

	_ = tmp104025

	tmp104026 := MakeNative(func(__e *ControlFlow) {
		V4326 := __e.Get(1)
		_ = V4326
		V4327 := __e.Get(2)
		_ = V4327
		V4328 := __e.Get(3)
		_ = V4328
		tmp104027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104028 := Call(__e, tmp104027)

		_ = tmp104028

		tmp104029 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104030 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104031 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104032 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104033 := Call(__e, tmp104032, symboolean, Nil)

		tmp104034 := Call(__e, tmp104031, sym_1_1_6, tmp104033)

		tmp104035 := Call(__e, tmp104030, symsymbol, tmp104034)

		__e.TailApply(tmp104029, V4326, tmp104035, V4327, V4328)
		return

	}, 3)

	tmp104036 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tc, tmp104026)

	_ = tmp104036

	tmp104037 := MakeNative(func(__e *ControlFlow) {
		V4336 := __e.Get(1)
		_ = V4336
		V4337 := __e.Get(2)
		_ = V4337
		V4338 := __e.Get(3)
		_ = V4338
		tmp104038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104039 := Call(__e, tmp104038)

		_ = tmp104039

		tmp104040 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104041 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104042 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104043 := Call(__e, tmp104042, symboolean, Nil)

		tmp104044 := Call(__e, tmp104041, sym_1_1_6, tmp104043)

		__e.TailApply(tmp104040, V4336, tmp104044, V4337, V4338)
		return

	}, 3)

	tmp104045 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tc_2, tmp104037)

	_ = tmp104045

	tmp104046 := MakeNative(func(__e *ControlFlow) {
		V4346 := __e.Get(1)
		_ = V4346
		V4347 := __e.Get(2)
		_ = V4347
		V4348 := __e.Get(3)
		_ = V4348
		tmp104047 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104048 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104049 := Call(__e, tmp104048)

			_ = tmp104049

			tmp104050 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104051 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104054 := Call(__e, tmp104053, A, Nil)

			tmp104055 := Call(__e, tmp104052, symlazy, tmp104054)

			tmp104056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104057 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104058 := Call(__e, tmp104057, A, Nil)

			tmp104059 := Call(__e, tmp104056, sym_1_1_6, tmp104058)

			tmp104060 := Call(__e, tmp104051, tmp104055, tmp104059)

			__e.TailApply(tmp104050, V4346, tmp104060, V4347, V4348)
			return

		}, 1)

		tmp104061 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104062 := Call(__e, tmp104061, V4347)

		__e.TailApply(tmp104047, tmp104062)
		return

	}, 3)

	tmp104063 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1thaw, tmp104046)

	_ = tmp104063

	tmp104064 := MakeNative(func(__e *ControlFlow) {
		V4356 := __e.Get(1)
		_ = V4356
		V4357 := __e.Get(2)
		_ = V4357
		V4358 := __e.Get(3)
		_ = V4358
		tmp104065 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104066 := Call(__e, tmp104065)

		_ = tmp104066

		tmp104067 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104068 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104069 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104070 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104071 := Call(__e, tmp104070, symsymbol, Nil)

		tmp104072 := Call(__e, tmp104069, sym_1_1_6, tmp104071)

		tmp104073 := Call(__e, tmp104068, symsymbol, tmp104072)

		__e.TailApply(tmp104067, V4356, tmp104073, V4357, V4358)
		return

	}, 3)

	tmp104074 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1track, tmp104064)

	_ = tmp104074

	tmp104075 := MakeNative(func(__e *ControlFlow) {
		V4366 := __e.Get(1)
		_ = V4366
		V4367 := __e.Get(2)
		_ = V4367
		V4368 := __e.Get(3)
		_ = V4368
		tmp104076 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104077 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104078 := Call(__e, tmp104077)

			_ = tmp104078

			tmp104079 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104081 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104087 := Call(__e, tmp104086, A, Nil)

			tmp104088 := Call(__e, tmp104085, sym_1_1_6, tmp104087)

			tmp104089 := Call(__e, tmp104084, symexception, tmp104088)

			tmp104090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104091 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104092 := Call(__e, tmp104091, A, Nil)

			tmp104093 := Call(__e, tmp104090, sym_1_1_6, tmp104092)

			tmp104094 := Call(__e, tmp104083, tmp104089, tmp104093)

			tmp104095 := Call(__e, tmp104082, tmp104094, Nil)

			tmp104096 := Call(__e, tmp104081, sym_1_1_6, tmp104095)

			tmp104097 := Call(__e, tmp104080, A, tmp104096)

			__e.TailApply(tmp104079, V4366, tmp104097, V4367, V4368)
			return

		}, 1)

		tmp104098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104099 := Call(__e, tmp104098, V4367)

		__e.TailApply(tmp104076, tmp104099)
		return

	}, 3)

	tmp104100 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1trap_1error, tmp104075)

	_ = tmp104100

	tmp104101 := MakeNative(func(__e *ControlFlow) {
		V4376 := __e.Get(1)
		_ = V4376
		V4377 := __e.Get(2)
		_ = V4377
		V4378 := __e.Get(3)
		_ = V4378
		tmp104102 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104103 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104104 := Call(__e, tmp104103)

			_ = tmp104104

			tmp104105 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104106 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104107 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104109 := Call(__e, tmp104108, symboolean, Nil)

			tmp104110 := Call(__e, tmp104107, sym_1_1_6, tmp104109)

			tmp104111 := Call(__e, tmp104106, A, tmp104110)

			__e.TailApply(tmp104105, V4376, tmp104111, V4377, V4378)
			return

		}, 1)

		tmp104112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104113 := Call(__e, tmp104112, V4377)

		__e.TailApply(tmp104102, tmp104113)
		return

	}, 3)

	tmp104114 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tuple_2, tmp104101)

	_ = tmp104114

	tmp104115 := MakeNative(func(__e *ControlFlow) {
		V4386 := __e.Get(1)
		_ = V4386
		V4387 := __e.Get(2)
		_ = V4387
		V4388 := __e.Get(3)
		_ = V4388
		tmp104116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104117 := Call(__e, tmp104116)

		_ = tmp104117

		tmp104118 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104120 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104122 := Call(__e, tmp104121, symsymbol, Nil)

		tmp104123 := Call(__e, tmp104120, sym_1_1_6, tmp104122)

		tmp104124 := Call(__e, tmp104119, symsymbol, tmp104123)

		__e.TailApply(tmp104118, V4386, tmp104124, V4387, V4388)
		return

	}, 3)

	tmp104125 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1undefmacro, tmp104115)

	_ = tmp104125

	tmp104126 := MakeNative(func(__e *ControlFlow) {
		V4396 := __e.Get(1)
		_ = V4396
		V4397 := __e.Get(2)
		_ = V4397
		V4398 := __e.Get(3)
		_ = V4398
		tmp104127 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104128 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104129 := Call(__e, tmp104128)

			_ = tmp104129

			tmp104130 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104133 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104134 := Call(__e, tmp104133, A, Nil)

			tmp104135 := Call(__e, tmp104132, symlist, tmp104134)

			tmp104136 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104137 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104138 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104139 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104140 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104141 := Call(__e, tmp104140, A, Nil)

			tmp104142 := Call(__e, tmp104139, symlist, tmp104141)

			tmp104143 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104144 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104146 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104147 := Call(__e, tmp104146, A, Nil)

			tmp104148 := Call(__e, tmp104145, symlist, tmp104147)

			tmp104149 := Call(__e, tmp104144, tmp104148, Nil)

			tmp104150 := Call(__e, tmp104143, sym_1_1_6, tmp104149)

			tmp104151 := Call(__e, tmp104138, tmp104142, tmp104150)

			tmp104152 := Call(__e, tmp104137, tmp104151, Nil)

			tmp104153 := Call(__e, tmp104136, sym_1_1_6, tmp104152)

			tmp104154 := Call(__e, tmp104131, tmp104135, tmp104153)

			__e.TailApply(tmp104130, V4396, tmp104154, V4397, V4398)
			return

		}, 1)

		tmp104155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104156 := Call(__e, tmp104155, V4397)

		__e.TailApply(tmp104127, tmp104156)
		return

	}, 3)

	tmp104157 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1union, tmp104126)

	_ = tmp104157

	tmp104158 := MakeNative(func(__e *ControlFlow) {
		V4406 := __e.Get(1)
		_ = V4406
		V4407 := __e.Get(2)
		_ = V4407
		V4408 := __e.Get(3)
		_ = V4408
		tmp104159 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104160 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp104161 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp104162 := Call(__e, tmp104161)

				_ = tmp104162

				tmp104163 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp104164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104167 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104168 := Call(__e, tmp104167, B, Nil)

				tmp104169 := Call(__e, tmp104166, sym_1_1_6, tmp104168)

				tmp104170 := Call(__e, tmp104165, A, tmp104169)

				tmp104171 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104174 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104175 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104176 := Call(__e, tmp104175, B, Nil)

				tmp104177 := Call(__e, tmp104174, sym_1_1_6, tmp104176)

				tmp104178 := Call(__e, tmp104173, A, tmp104177)

				tmp104179 := Call(__e, tmp104172, tmp104178, Nil)

				tmp104180 := Call(__e, tmp104171, sym_1_1_6, tmp104179)

				tmp104181 := Call(__e, tmp104164, tmp104170, tmp104180)

				__e.TailApply(tmp104163, V4406, tmp104181, V4407, V4408)
				return

			}, 1)

			tmp104182 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp104183 := Call(__e, tmp104182, V4407)

			__e.TailApply(tmp104160, tmp104183)
			return

		}, 1)

		tmp104184 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104185 := Call(__e, tmp104184, V4407)

		__e.TailApply(tmp104159, tmp104185)
		return

	}, 3)

	tmp104186 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1unprofile, tmp104158)

	_ = tmp104186

	tmp104187 := MakeNative(func(__e *ControlFlow) {
		V4416 := __e.Get(1)
		_ = V4416
		V4417 := __e.Get(2)
		_ = V4417
		V4418 := __e.Get(3)
		_ = V4418
		tmp104188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104189 := Call(__e, tmp104188)

		_ = tmp104189

		tmp104190 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104191 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104192 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104193 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104194 := Call(__e, tmp104193, symsymbol, Nil)

		tmp104195 := Call(__e, tmp104192, sym_1_1_6, tmp104194)

		tmp104196 := Call(__e, tmp104191, symsymbol, tmp104195)

		__e.TailApply(tmp104190, V4416, tmp104196, V4417, V4418)
		return

	}, 3)

	tmp104197 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1untrack, tmp104187)

	_ = tmp104197

	tmp104198 := MakeNative(func(__e *ControlFlow) {
		V4426 := __e.Get(1)
		_ = V4426
		V4427 := __e.Get(2)
		_ = V4427
		V4428 := __e.Get(3)
		_ = V4428
		tmp104199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104200 := Call(__e, tmp104199)

		_ = tmp104200

		tmp104201 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104202 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104203 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104204 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104205 := Call(__e, tmp104204, symsymbol, Nil)

		tmp104206 := Call(__e, tmp104203, sym_1_1_6, tmp104205)

		tmp104207 := Call(__e, tmp104202, symsymbol, tmp104206)

		__e.TailApply(tmp104201, V4426, tmp104207, V4427, V4428)
		return

	}, 3)

	tmp104208 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1unspecialise, tmp104198)

	_ = tmp104208

	tmp104209 := MakeNative(func(__e *ControlFlow) {
		V4436 := __e.Get(1)
		_ = V4436
		V4437 := __e.Get(2)
		_ = V4437
		V4438 := __e.Get(3)
		_ = V4438
		tmp104210 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104211 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104212 := Call(__e, tmp104211)

			_ = tmp104212

			tmp104213 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104214 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104215 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104216 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104217 := Call(__e, tmp104216, symboolean, Nil)

			tmp104218 := Call(__e, tmp104215, sym_1_1_6, tmp104217)

			tmp104219 := Call(__e, tmp104214, A, tmp104218)

			__e.TailApply(tmp104213, V4436, tmp104219, V4437, V4438)
			return

		}, 1)

		tmp104220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104221 := Call(__e, tmp104220, V4437)

		__e.TailApply(tmp104210, tmp104221)
		return

	}, 3)

	tmp104222 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1variable_2, tmp104209)

	_ = tmp104222

	tmp104223 := MakeNative(func(__e *ControlFlow) {
		V4446 := __e.Get(1)
		_ = V4446
		V4447 := __e.Get(2)
		_ = V4447
		V4448 := __e.Get(3)
		_ = V4448
		tmp104224 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104225 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104226 := Call(__e, tmp104225)

			_ = tmp104226

			tmp104227 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104228 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104229 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104230 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104231 := Call(__e, tmp104230, symboolean, Nil)

			tmp104232 := Call(__e, tmp104229, sym_1_1_6, tmp104231)

			tmp104233 := Call(__e, tmp104228, A, tmp104232)

			__e.TailApply(tmp104227, V4446, tmp104233, V4447, V4448)
			return

		}, 1)

		tmp104234 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104235 := Call(__e, tmp104234, V4447)

		__e.TailApply(tmp104224, tmp104235)
		return

	}, 3)

	tmp104236 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector_2, tmp104223)

	_ = tmp104236

	tmp104237 := MakeNative(func(__e *ControlFlow) {
		V4456 := __e.Get(1)
		_ = V4456
		V4457 := __e.Get(2)
		_ = V4457
		V4458 := __e.Get(3)
		_ = V4458
		tmp104238 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104239 := Call(__e, tmp104238)

		_ = tmp104239

		tmp104240 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104241 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104242 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104243 := Call(__e, tmp104242, symstring, Nil)

		tmp104244 := Call(__e, tmp104241, sym_1_1_6, tmp104243)

		__e.TailApply(tmp104240, V4456, tmp104244, V4457, V4458)
		return

	}, 3)

	tmp104245 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1version, tmp104237)

	_ = tmp104245

	tmp104246 := MakeNative(func(__e *ControlFlow) {
		V4466 := __e.Get(1)
		_ = V4466
		V4467 := __e.Get(2)
		_ = V4467
		V4468 := __e.Get(3)
		_ = V4468
		tmp104247 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104248 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104249 := Call(__e, tmp104248)

			_ = tmp104249

			tmp104250 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104251 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104252 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104253 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104254 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104255 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104256 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104257 := Call(__e, tmp104256, A, Nil)

			tmp104258 := Call(__e, tmp104255, sym_1_1_6, tmp104257)

			tmp104259 := Call(__e, tmp104254, A, tmp104258)

			tmp104260 := Call(__e, tmp104253, tmp104259, Nil)

			tmp104261 := Call(__e, tmp104252, sym_1_1_6, tmp104260)

			tmp104262 := Call(__e, tmp104251, symstring, tmp104261)

			__e.TailApply(tmp104250, V4466, tmp104262, V4467, V4468)
			return

		}, 1)

		tmp104263 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104264 := Call(__e, tmp104263, V4467)

		__e.TailApply(tmp104247, tmp104264)
		return

	}, 3)

	tmp104265 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1write_1to_1file, tmp104246)

	_ = tmp104265

	tmp104266 := MakeNative(func(__e *ControlFlow) {
		V4476 := __e.Get(1)
		_ = V4476
		V4477 := __e.Get(2)
		_ = V4477
		V4478 := __e.Get(3)
		_ = V4478
		tmp104267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104268 := Call(__e, tmp104267)

		_ = tmp104268

		tmp104269 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104270 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104271 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104272 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104273 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104274 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104275 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104276 := Call(__e, tmp104275, symout, Nil)

		tmp104277 := Call(__e, tmp104274, symstream, tmp104276)

		tmp104278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104279 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104280 := Call(__e, tmp104279, symnumber, Nil)

		tmp104281 := Call(__e, tmp104278, sym_1_1_6, tmp104280)

		tmp104282 := Call(__e, tmp104273, tmp104277, tmp104281)

		tmp104283 := Call(__e, tmp104272, tmp104282, Nil)

		tmp104284 := Call(__e, tmp104271, sym_1_1_6, tmp104283)

		tmp104285 := Call(__e, tmp104270, symnumber, tmp104284)

		__e.TailApply(tmp104269, V4476, tmp104285, V4477, V4478)
		return

	}, 3)

	tmp104286 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1write_1byte, tmp104266)

	_ = tmp104286

	tmp104287 := MakeNative(func(__e *ControlFlow) {
		V4486 := __e.Get(1)
		_ = V4486
		V4487 := __e.Get(2)
		_ = V4487
		V4488 := __e.Get(3)
		_ = V4488
		tmp104288 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104289 := Call(__e, tmp104288)

		_ = tmp104289

		tmp104290 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104291 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104292 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104293 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104294 := Call(__e, tmp104293, symboolean, Nil)

		tmp104295 := Call(__e, tmp104292, sym_1_1_6, tmp104294)

		tmp104296 := Call(__e, tmp104291, symstring, tmp104295)

		__e.TailApply(tmp104290, V4486, tmp104296, V4487, V4488)
		return

	}, 3)

	tmp104297 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1y_1or_1n_2, tmp104287)

	_ = tmp104297

	tmp104298 := MakeNative(func(__e *ControlFlow) {
		V4496 := __e.Get(1)
		_ = V4496
		V4497 := __e.Get(2)
		_ = V4497
		V4498 := __e.Get(3)
		_ = V4498
		tmp104299 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104300 := Call(__e, tmp104299)

		_ = tmp104300

		tmp104301 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104302 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104305 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104306 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104307 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104308 := Call(__e, tmp104307, symboolean, Nil)

		tmp104309 := Call(__e, tmp104306, sym_1_1_6, tmp104308)

		tmp104310 := Call(__e, tmp104305, symnumber, tmp104309)

		tmp104311 := Call(__e, tmp104304, tmp104310, Nil)

		tmp104312 := Call(__e, tmp104303, sym_1_1_6, tmp104311)

		tmp104313 := Call(__e, tmp104302, symnumber, tmp104312)

		__e.TailApply(tmp104301, V4496, tmp104313, V4497, V4498)
		return

	}, 3)

	tmp104314 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_6, tmp104298)

	_ = tmp104314

	tmp104315 := MakeNative(func(__e *ControlFlow) {
		V4506 := __e.Get(1)
		_ = V4506
		V4507 := __e.Get(2)
		_ = V4507
		V4508 := __e.Get(3)
		_ = V4508
		tmp104316 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104317 := Call(__e, tmp104316)

		_ = tmp104317

		tmp104318 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104325 := Call(__e, tmp104324, symboolean, Nil)

		tmp104326 := Call(__e, tmp104323, sym_1_1_6, tmp104325)

		tmp104327 := Call(__e, tmp104322, symnumber, tmp104326)

		tmp104328 := Call(__e, tmp104321, tmp104327, Nil)

		tmp104329 := Call(__e, tmp104320, sym_1_1_6, tmp104328)

		tmp104330 := Call(__e, tmp104319, symnumber, tmp104329)

		__e.TailApply(tmp104318, V4506, tmp104330, V4507, V4508)
		return

	}, 3)

	tmp104331 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5, tmp104315)

	_ = tmp104331

	tmp104332 := MakeNative(func(__e *ControlFlow) {
		V4516 := __e.Get(1)
		_ = V4516
		V4517 := __e.Get(2)
		_ = V4517
		V4518 := __e.Get(3)
		_ = V4518
		tmp104333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104334 := Call(__e, tmp104333)

		_ = tmp104334

		tmp104335 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104339 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104340 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104342 := Call(__e, tmp104341, symboolean, Nil)

		tmp104343 := Call(__e, tmp104340, sym_1_1_6, tmp104342)

		tmp104344 := Call(__e, tmp104339, symnumber, tmp104343)

		tmp104345 := Call(__e, tmp104338, tmp104344, Nil)

		tmp104346 := Call(__e, tmp104337, sym_1_1_6, tmp104345)

		tmp104347 := Call(__e, tmp104336, symnumber, tmp104346)

		__e.TailApply(tmp104335, V4516, tmp104347, V4517, V4518)
		return

	}, 3)

	tmp104348 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_6_a, tmp104332)

	_ = tmp104348

	tmp104349 := MakeNative(func(__e *ControlFlow) {
		V4526 := __e.Get(1)
		_ = V4526
		V4527 := __e.Get(2)
		_ = V4527
		V4528 := __e.Get(3)
		_ = V4528
		tmp104350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104351 := Call(__e, tmp104350)

		_ = tmp104351

		tmp104352 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104359 := Call(__e, tmp104358, symboolean, Nil)

		tmp104360 := Call(__e, tmp104357, sym_1_1_6, tmp104359)

		tmp104361 := Call(__e, tmp104356, symnumber, tmp104360)

		tmp104362 := Call(__e, tmp104355, tmp104361, Nil)

		tmp104363 := Call(__e, tmp104354, sym_1_1_6, tmp104362)

		tmp104364 := Call(__e, tmp104353, symnumber, tmp104363)

		__e.TailApply(tmp104352, V4526, tmp104364, V4527, V4528)
		return

	}, 3)

	tmp104365 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_a, tmp104349)

	_ = tmp104365

	tmp104366 := MakeNative(func(__e *ControlFlow) {
		V4536 := __e.Get(1)
		_ = V4536
		V4537 := __e.Get(2)
		_ = V4537
		V4538 := __e.Get(3)
		_ = V4538
		tmp104367 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp104369 := Call(__e, tmp104368)

			_ = tmp104369

			tmp104370 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp104371 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104372 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104373 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104376 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104377 := Call(__e, tmp104376, symboolean, Nil)

			tmp104378 := Call(__e, tmp104375, sym_1_1_6, tmp104377)

			tmp104379 := Call(__e, tmp104374, A, tmp104378)

			tmp104380 := Call(__e, tmp104373, tmp104379, Nil)

			tmp104381 := Call(__e, tmp104372, sym_1_1_6, tmp104380)

			tmp104382 := Call(__e, tmp104371, A, tmp104381)

			__e.TailApply(tmp104370, V4536, tmp104382, V4537, V4538)
			return

		}, 1)

		tmp104383 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104384 := Call(__e, tmp104383, V4537)

		__e.TailApply(tmp104367, tmp104384)
		return

	}, 3)

	tmp104385 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_a, tmp104366)

	_ = tmp104385

	tmp104386 := MakeNative(func(__e *ControlFlow) {
		V4546 := __e.Get(1)
		_ = V4546
		V4547 := __e.Get(2)
		_ = V4547
		V4548 := __e.Get(3)
		_ = V4548
		tmp104387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104388 := Call(__e, tmp104387)

		_ = tmp104388

		tmp104389 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104391 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104392 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104393 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104394 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104396 := Call(__e, tmp104395, symnumber, Nil)

		tmp104397 := Call(__e, tmp104394, sym_1_1_6, tmp104396)

		tmp104398 := Call(__e, tmp104393, symnumber, tmp104397)

		tmp104399 := Call(__e, tmp104392, tmp104398, Nil)

		tmp104400 := Call(__e, tmp104391, sym_1_1_6, tmp104399)

		tmp104401 := Call(__e, tmp104390, symnumber, tmp104400)

		__e.TailApply(tmp104389, V4546, tmp104401, V4547, V4548)
		return

	}, 3)

	tmp104402 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_7, tmp104386)

	_ = tmp104402

	tmp104403 := MakeNative(func(__e *ControlFlow) {
		V4556 := __e.Get(1)
		_ = V4556
		V4557 := __e.Get(2)
		_ = V4557
		V4558 := __e.Get(3)
		_ = V4558
		tmp104404 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104405 := Call(__e, tmp104404)

		_ = tmp104405

		tmp104406 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104407 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104408 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104409 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104410 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104413 := Call(__e, tmp104412, symnumber, Nil)

		tmp104414 := Call(__e, tmp104411, sym_1_1_6, tmp104413)

		tmp104415 := Call(__e, tmp104410, symnumber, tmp104414)

		tmp104416 := Call(__e, tmp104409, tmp104415, Nil)

		tmp104417 := Call(__e, tmp104408, sym_1_1_6, tmp104416)

		tmp104418 := Call(__e, tmp104407, symnumber, tmp104417)

		__e.TailApply(tmp104406, V4556, tmp104418, V4557, V4558)
		return

	}, 3)

	tmp104419 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_c, tmp104403)

	_ = tmp104419

	tmp104420 := MakeNative(func(__e *ControlFlow) {
		V4566 := __e.Get(1)
		_ = V4566
		V4567 := __e.Get(2)
		_ = V4567
		V4568 := __e.Get(3)
		_ = V4568
		tmp104421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104422 := Call(__e, tmp104421)

		_ = tmp104422

		tmp104423 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104424 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104425 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104426 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104427 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104428 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104429 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104430 := Call(__e, tmp104429, symnumber, Nil)

		tmp104431 := Call(__e, tmp104428, sym_1_1_6, tmp104430)

		tmp104432 := Call(__e, tmp104427, symnumber, tmp104431)

		tmp104433 := Call(__e, tmp104426, tmp104432, Nil)

		tmp104434 := Call(__e, tmp104425, sym_1_1_6, tmp104433)

		tmp104435 := Call(__e, tmp104424, symnumber, tmp104434)

		__e.TailApply(tmp104423, V4566, tmp104435, V4567, V4568)
		return

	}, 3)

	tmp104436 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_1, tmp104420)

	_ = tmp104436

	tmp104437 := MakeNative(func(__e *ControlFlow) {
		V4576 := __e.Get(1)
		_ = V4576
		V4577 := __e.Get(2)
		_ = V4577
		V4578 := __e.Get(3)
		_ = V4578
		tmp104438 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp104439 := Call(__e, tmp104438)

		_ = tmp104439

		tmp104440 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		tmp104441 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104442 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104443 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp104447 := Call(__e, tmp104446, symnumber, Nil)

		tmp104448 := Call(__e, tmp104445, sym_1_1_6, tmp104447)

		tmp104449 := Call(__e, tmp104444, symnumber, tmp104448)

		tmp104450 := Call(__e, tmp104443, tmp104449, Nil)

		tmp104451 := Call(__e, tmp104442, sym_1_1_6, tmp104450)

		tmp104452 := Call(__e, tmp104441, symnumber, tmp104451)

		__e.TailApply(tmp104440, V4576, tmp104452, V4577, V4578)
		return

	}, 3)

	tmp104453 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_d, tmp104437)

	_ = tmp104453

	tmp104454 := MakeNative(func(__e *ControlFlow) {
		V4586 := __e.Get(1)
		_ = V4586
		V4587 := __e.Get(2)
		_ = V4587
		V4588 := __e.Get(3)
		_ = V4588
		tmp104455 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp104456 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp104457 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp104458 := Call(__e, tmp104457)

				_ = tmp104458

				tmp104459 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				tmp104460 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104461 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104462 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104463 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104464 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104465 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104466 := Call(__e, tmp104465, symboolean, Nil)

				tmp104467 := Call(__e, tmp104464, sym_1_1_6, tmp104466)

				tmp104468 := Call(__e, tmp104463, B, tmp104467)

				tmp104469 := Call(__e, tmp104462, tmp104468, Nil)

				tmp104470 := Call(__e, tmp104461, sym_1_1_6, tmp104469)

				tmp104471 := Call(__e, tmp104460, A, tmp104470)

				__e.TailApply(tmp104459, V4586, tmp104471, V4587, V4588)
				return

			}, 1)

			tmp104472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp104473 := Call(__e, tmp104472, V4587)

			__e.TailApply(tmp104456, tmp104473)
			return

		}, 1)

		tmp104474 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp104475 := Call(__e, tmp104474, V4587)

		__e.TailApply(tmp104455, tmp104475)
		return

	}, 3)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_a_a, tmp104454)
	return

}, 0)
