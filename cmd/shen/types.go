package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen28900 := MakeNative(func(__e *ControlFlow) {
		V3210 := __e.Get(1)
		_ = V3210
		V3211 := __e.Get(2)
		_ = V3211
		gen28891 := MakeNative(func(__e *ControlFlow) {
			Record := __e.Get(1)
			_ = Record
			gen28886 := MakeNative(func(__e *ControlFlow) {
				Variancy := __e.Get(1)
				_ = Variancy
				gen28881 := MakeNative(func(__e *ControlFlow) {
					Type := __e.Get(1)
					_ = Type
					gen28878 := MakeNative(func(__e *ControlFlow) {
						F_d := __e.Get(1)
						_ = F_d
						gen28875 := MakeNative(func(__e *ControlFlow) {
							Parameters := __e.Get(1)
							_ = Parameters
							gen28856 := MakeNative(func(__e *ControlFlow) {
								Clause := __e.Get(1)
								_ = Clause
								gen28853 := MakeNative(func(__e *ControlFlow) {
									AUM__instruction := __e.Get(1)
									_ = AUM__instruction
									gen28850 := MakeNative(func(__e *ControlFlow) {
										Code := __e.Get(1)
										_ = Code
										gen28833 := MakeNative(func(__e *ControlFlow) {
											ShenDef := __e.Get(1)
											_ = ShenDef
											gen28830 := MakeNative(func(__e *ControlFlow) {
												Eval := __e.Get(1)
												_ = Eval
												__e.Return(V3210)
												return
											}, 1)

											gen28831 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

											gen28832 := Call(__e, gen28831, ShenDef)

											__e.TailApply(gen28830, gen28832)

											return

										}, 1)

										gen28834 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen28835 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen28836 := Call(__e, PrimNS1Value(symns2_1value), symappend)

										gen28837 := Call(__e, PrimNS1Value(symns2_1value), symappend)

										gen28838 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen28839 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen28840 := Call(__e, gen28839, symContinuation, Nil)

										gen28841 := Call(__e, gen28838, symProcessN, gen28840)

										gen28842 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen28843 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen28844 := Call(__e, gen28843, Code, Nil)

										gen28845 := Call(__e, gen28842, sym_1_6, gen28844)

										gen28846 := Call(__e, gen28837, gen28841, gen28845)

										gen28847 := Call(__e, gen28836, Parameters, gen28846)

										gen28848 := Call(__e, gen28835, F_d, gen28847)

										gen28849 := Call(__e, gen28834, symdefine, gen28848)

										__e.TailApply(gen28833, gen28849)

										return

									}, 1)

									gen28851 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum__to__shen)

									gen28852 := Call(__e, gen28851, AUM__instruction)

									__e.TailApply(gen28850, gen28852)

									return

								}, 1)

								gen28854 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aum)

								gen28855 := Call(__e, gen28854, Clause, Parameters)

								__e.TailApply(gen28853, gen28855)

								return

							}, 1)

							gen28857 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28858 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28859 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28860 := Call(__e, gen28859, symX, Nil)

							gen28861 := Call(__e, gen28858, F_d, gen28860)

							gen28862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28863 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28864 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28865 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28866 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28867 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen28868 := Call(__e, gen28867, Type, Nil)

							gen28869 := Call(__e, gen28866, symX, gen28868)

							gen28870 := Call(__e, gen28865, symunify_b, gen28869)

							gen28871 := Call(__e, gen28864, gen28870, Nil)

							gen28872 := Call(__e, gen28863, gen28871, Nil)

							gen28873 := Call(__e, gen28862, sym_h_1, gen28872)

							gen28874 := Call(__e, gen28857, gen28861, gen28873)

							__e.TailApply(gen28856, gen28874)

							return

						}, 1)

						gen28876 := Call(__e, PrimNS1Value(symns2_1value), symshen_4parameters)

						gen28877 := Call(__e, gen28876, MakeNumber(1))

						__e.TailApply(gen28875, gen28877)

						return

					}, 1)

					gen28879 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					gen28880 := Call(__e, gen28879, symshen_4type_1signature_1of_1, V3210)

					__e.TailApply(gen28878, gen28880)

					return

				}, 1)

				gen28882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

				gen28883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				gen28884 := Call(__e, gen28883, V3211)

				gen28885 := Call(__e, gen28882, gen28884)

				__e.TailApply(gen28881, gen28885)

				return

			}, 1)

			gen28888 := MakeNative(func(__e *ControlFlow) {
				gen28887 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variancy_1test)

				__e.TailApply(gen28887, V3210, V3211)

				return

			}, 0)

			gen28889 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4skip)
				return
			}, 1)

			gen28890 := Call(__e, PrimNS1Value(symtry_1catch), gen28888, gen28889)

			__e.TailApply(gen28886, gen28890)

			return

		}, 1)

		gen28892 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen28893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen28894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen28895 := Call(__e, gen28894, V3210, V3211)

		gen28896 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen28897 := Call(__e, gen28896, symshen_4_dsignedfuncs_d)

		gen28898 := Call(__e, gen28893, gen28895, gen28897)

		gen28899 := Call(__e, gen28892, symshen_4_dsignedfuncs_d, gen28898)

		__e.TailApply(gen28891, gen28899)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symdeclare, gen28900)

	gen28909 := MakeNative(func(__e *ControlFlow) {
		V3213 := __e.Get(1)
		_ = V3213
		gen28904 := MakeNative(func(__e *ControlFlow) {
			Demod := __e.Get(1)
			_ = Demod
			gen28902 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen28903 := Call(__e, gen28902, Demod, V3213)

			if True == gen28903 {
				__e.Return(V3213)
				return
			} else {
				gen28901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				__e.TailApply(gen28901, Demod)

				return

			}

		}, 1)

		gen28905 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

		gen28906 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen28907 := Call(__e, gen28906, symshen_4_ddemodulation_1function_d)

		gen28908 := Call(__e, gen28905, gen28907, V3213)

		__e.TailApply(gen28904, gen28908)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4demodulate, gen28909)

	gen28926 := MakeNative(func(__e *ControlFlow) {
		V3216 := __e.Get(1)
		_ = V3216
		V3217 := __e.Get(2)
		_ = V3217
		gen28923 := MakeNative(func(__e *ControlFlow) {
			TypeF := __e.Get(1)
			_ = TypeF
			gen28910 := MakeNative(func(__e *ControlFlow) {
				Check := __e.Get(1)
				_ = Check
				__e.Return(symshen_4skip)
				return
			}, 1)

			gen28920 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen28921 := Call(__e, gen28920, symsymbol, TypeF)

			var gen28922 Obj

			if True == gen28921 {
				gen28922 = symshen_4skip
			} else {
				gen28918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

				gen28919 := Call(__e, gen28918, TypeF, V3217)

				if True == gen28919 {
					gen28922 = symshen_4skip
				} else {
					gen28911 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					gen28912 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen28913 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen28914 := Call(__e, gen28913, V3216, MakeString(" may create errors\n"), symshen_4a)

					gen28915 := Call(__e, gen28912, MakeString("warning: changing the type of "), gen28914)

					gen28916 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					gen28917 := Call(__e, gen28916)

					gen28922 = Call(__e, gen28911, gen28915, gen28917)

				}

			}

			__e.TailApply(gen28910, gen28922)

			return

		}, 1)

		gen28924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck)

		gen28925 := Call(__e, gen28924, V3216, symB)

		__e.TailApply(gen28923, gen28925)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4variancy_1test, gen28926)

	gen29003 := MakeNative(func(__e *ControlFlow) {
		V3230 := __e.Get(1)
		_ = V3230
		V3231 := __e.Get(2)
		_ = V3231
		gen29001 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen29002 := Call(__e, gen29001, V3231, V3230)

		if True == gen29002 {
			__e.Return(True)
			return
		} else {
			gen28998 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen28999 := Call(__e, gen28998, V3230)

			var gen29000 Obj

			if True == gen28999 {
				gen28995 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen28996 := Call(__e, gen28995, V3231)

				var gen28997 Obj

				if True == gen28996 {
					gen28989 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen28990 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen28991 := Call(__e, gen28990, V3231)

					gen28992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen28993 := Call(__e, gen28992, V3230)

					gen28994 := Call(__e, gen28989, gen28991, gen28993)

					if True == gen28994 {
						gen28997 = True
					} else {
						gen28997 = False
					}

				} else {
					gen28997 = False
				}

				if True == gen28997 {
					gen29000 = True
				} else {
					gen29000 = False
				}

			} else {
				gen29000 = False
			}

			if True == gen29000 {
				gen28984 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

				gen28985 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen28986 := Call(__e, gen28985, V3230)

				gen28987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen28988 := Call(__e, gen28987, V3231)

				__e.TailApply(gen28984, gen28986, gen28988)

				return

			} else {
				gen28981 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen28982 := Call(__e, gen28981, V3230)

				var gen28983 Obj

				if True == gen28982 {
					gen28978 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen28979 := Call(__e, gen28978, V3231)

					var gen28980 Obj

					if True == gen28979 {
						gen28973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

						gen28974 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen28975 := Call(__e, gen28974, V3230)

						gen28976 := Call(__e, gen28973, gen28975)

						var gen28977 Obj

						if True == gen28976 {
							gen28969 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

							gen28970 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen28971 := Call(__e, gen28970, V3231)

							gen28972 := Call(__e, gen28969, gen28971)

							if True == gen28972 {
								gen28977 = True
							} else {
								gen28977 = False
							}

						} else {
							gen28977 = False
						}

						if True == gen28977 {
							gen28980 = True
						} else {
							gen28980 = False
						}

					} else {
						gen28980 = False
					}

					if True == gen28980 {
						gen28983 = True
					} else {
						gen28983 = False
					}

				} else {
					gen28983 = False
				}

				if True == gen28983 {
					gen28956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

					gen28957 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					gen28958 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen28959 := Call(__e, gen28958, V3230)

					gen28960 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen28961 := Call(__e, gen28960, V3230)

					gen28962 := Call(__e, gen28957, symshen_4a, gen28959, gen28961)

					gen28963 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					gen28964 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen28965 := Call(__e, gen28964, V3231)

					gen28966 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen28967 := Call(__e, gen28966, V3231)

					gen28968 := Call(__e, gen28963, symshen_4a, gen28965, gen28967)

					__e.TailApply(gen28956, gen28962, gen28968)

					return

				} else {
					gen28953 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen28954 := Call(__e, gen28953, V3230)

					var gen28955 Obj

					if True == gen28954 {
						gen28948 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen28949 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen28950 := Call(__e, gen28949, V3230)

						gen28951 := Call(__e, gen28948, gen28950)

						var gen28952 Obj

						if True == gen28951 {
							gen28945 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen28946 := Call(__e, gen28945, V3231)

							var gen28947 Obj

							if True == gen28946 {
								gen28941 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen28942 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen28943 := Call(__e, gen28942, V3231)

								gen28944 := Call(__e, gen28941, gen28943)

								if True == gen28944 {
									gen28947 = True
								} else {
									gen28947 = False
								}

							} else {
								gen28947 = False
							}

							if True == gen28947 {
								gen28952 = True
							} else {
								gen28952 = False
							}

						} else {
							gen28952 = False
						}

						if True == gen28952 {
							gen28955 = True
						} else {
							gen28955 = False
						}

					} else {
						gen28955 = False
					}

					if True == gen28955 {
						gen28928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variant_2)

						gen28929 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						gen28930 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen28931 := Call(__e, gen28930, V3230)

						gen28932 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen28933 := Call(__e, gen28932, V3230)

						gen28934 := Call(__e, gen28929, gen28931, gen28933)

						gen28935 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						gen28936 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen28937 := Call(__e, gen28936, V3231)

						gen28938 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen28939 := Call(__e, gen28938, V3231)

						gen28940 := Call(__e, gen28935, gen28937, gen28939)

						__e.TailApply(gen28928, gen28934, gen28940)

						return

					} else {
						if True == True {
							__e.Return(False)
							return
						} else {
							gen28927 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen28927, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4variant_2, gen29003)

	gen29015 := MakeNative(func(__e *ControlFlow) {
		V3236 := __e.Get(1)
		_ = V3236
		V3237 := __e.Get(2)
		_ = V3237
		V3238 := __e.Get(3)
		_ = V3238
		gen29012 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29004)

			gen29005 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29007 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29009 := Call(__e, gen29008, symboolean, Nil)

			gen29010 := Call(__e, gen29007, sym_1_1_6, gen29009)

			gen29011 := Call(__e, gen29006, A, gen29010)

			__e.TailApply(gen29005, V3236, gen29011, V3237, V3238)

			return

		}, 1)

		gen29013 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29014 := Call(__e, gen29013, V3237)

		__e.TailApply(gen29012, gen29014)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1absvector_2, gen29015)

	gen29041 := MakeNative(func(__e *ControlFlow) {
		V3246 := __e.Get(1)
		_ = V3246
		V3247 := __e.Get(2)
		_ = V3247
		V3248 := __e.Get(3)
		_ = V3248
		gen29038 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29016 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29016)

			gen29017 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29021 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29022 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29023 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29024 := Call(__e, gen29023, A, Nil)

			gen29025 := Call(__e, gen29022, symlist, gen29024)

			gen29026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29027 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29028 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29029 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29030 := Call(__e, gen29029, A, Nil)

			gen29031 := Call(__e, gen29028, symlist, gen29030)

			gen29032 := Call(__e, gen29027, gen29031, Nil)

			gen29033 := Call(__e, gen29026, sym_1_1_6, gen29032)

			gen29034 := Call(__e, gen29021, gen29025, gen29033)

			gen29035 := Call(__e, gen29020, gen29034, Nil)

			gen29036 := Call(__e, gen29019, sym_1_1_6, gen29035)

			gen29037 := Call(__e, gen29018, A, gen29036)

			__e.TailApply(gen29017, V3246, gen29037, V3247, V3248)

			return

		}, 1)

		gen29039 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29040 := Call(__e, gen29039, V3247)

		__e.TailApply(gen29038, gen29040)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1adjoin, gen29041)

	gen29056 := MakeNative(func(__e *ControlFlow) {
		V3256 := __e.Get(1)
		_ = V3256
		V3257 := __e.Get(2)
		_ = V3257
		V3258 := __e.Get(3)
		_ = V3258
		gen29042 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29042)

		gen29043 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29047 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29048 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29049 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29050 := Call(__e, gen29049, symboolean, Nil)

		gen29051 := Call(__e, gen29048, sym_1_1_6, gen29050)

		gen29052 := Call(__e, gen29047, symboolean, gen29051)

		gen29053 := Call(__e, gen29046, gen29052, Nil)

		gen29054 := Call(__e, gen29045, sym_1_1_6, gen29053)

		gen29055 := Call(__e, gen29044, symboolean, gen29054)

		__e.TailApply(gen29043, V3256, gen29055, V3257, V3258)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1and, gen29056)

	gen29080 := MakeNative(func(__e *ControlFlow) {
		V3266 := __e.Get(1)
		_ = V3266
		V3267 := __e.Get(2)
		_ = V3267
		V3268 := __e.Get(3)
		_ = V3268
		gen29077 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29057)

			gen29058 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29059 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29060 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29061 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29063 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29064 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29065 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29066 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29067 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29068 := Call(__e, gen29067, symstring, Nil)

			gen29069 := Call(__e, gen29066, sym_1_1_6, gen29068)

			gen29070 := Call(__e, gen29065, symsymbol, gen29069)

			gen29071 := Call(__e, gen29064, gen29070, Nil)

			gen29072 := Call(__e, gen29063, sym_1_1_6, gen29071)

			gen29073 := Call(__e, gen29062, symstring, gen29072)

			gen29074 := Call(__e, gen29061, gen29073, Nil)

			gen29075 := Call(__e, gen29060, sym_1_1_6, gen29074)

			gen29076 := Call(__e, gen29059, A, gen29075)

			__e.TailApply(gen29058, V3266, gen29076, V3267, V3268)

			return

		}, 1)

		gen29078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29079 := Call(__e, gen29078, V3267)

		__e.TailApply(gen29077, gen29079)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4app, gen29080)

	gen29110 := MakeNative(func(__e *ControlFlow) {
		V3276 := __e.Get(1)
		_ = V3276
		V3277 := __e.Get(2)
		_ = V3277
		V3278 := __e.Get(3)
		_ = V3278
		gen29107 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29081)

			gen29082 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29086 := Call(__e, gen29085, A, Nil)

			gen29087 := Call(__e, gen29084, symlist, gen29086)

			gen29088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29091 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29092 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29093 := Call(__e, gen29092, A, Nil)

			gen29094 := Call(__e, gen29091, symlist, gen29093)

			gen29095 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29096 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29097 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29098 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29099 := Call(__e, gen29098, A, Nil)

			gen29100 := Call(__e, gen29097, symlist, gen29099)

			gen29101 := Call(__e, gen29096, gen29100, Nil)

			gen29102 := Call(__e, gen29095, sym_1_1_6, gen29101)

			gen29103 := Call(__e, gen29090, gen29094, gen29102)

			gen29104 := Call(__e, gen29089, gen29103, Nil)

			gen29105 := Call(__e, gen29088, sym_1_1_6, gen29104)

			gen29106 := Call(__e, gen29083, gen29087, gen29105)

			__e.TailApply(gen29082, V3276, gen29106, V3277, V3278)

			return

		}, 1)

		gen29108 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29109 := Call(__e, gen29108, V3277)

		__e.TailApply(gen29107, gen29109)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1append, gen29110)

	gen29122 := MakeNative(func(__e *ControlFlow) {
		V3286 := __e.Get(1)
		_ = V3286
		V3287 := __e.Get(2)
		_ = V3287
		V3288 := __e.Get(3)
		_ = V3288
		gen29119 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29111)

			gen29112 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29115 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29116 := Call(__e, gen29115, symnumber, Nil)

			gen29117 := Call(__e, gen29114, sym_1_1_6, gen29116)

			gen29118 := Call(__e, gen29113, A, gen29117)

			__e.TailApply(gen29112, V3286, gen29118, V3287, V3288)

			return

		}, 1)

		gen29120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29121 := Call(__e, gen29120, V3287)

		__e.TailApply(gen29119, gen29121)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1arity, gen29122)

	gen29152 := MakeNative(func(__e *ControlFlow) {
		V3296 := __e.Get(1)
		_ = V3296
		V3297 := __e.Get(2)
		_ = V3297
		V3298 := __e.Get(3)
		_ = V3298
		gen29149 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29123)

			gen29124 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29125 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29126 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29127 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29133 := Call(__e, gen29132, A, Nil)

			gen29134 := Call(__e, gen29131, symlist, gen29133)

			gen29135 := Call(__e, gen29130, gen29134, Nil)

			gen29136 := Call(__e, gen29129, symlist, gen29135)

			gen29137 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29138 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29139 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29140 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29141 := Call(__e, gen29140, A, Nil)

			gen29142 := Call(__e, gen29139, symlist, gen29141)

			gen29143 := Call(__e, gen29138, gen29142, Nil)

			gen29144 := Call(__e, gen29137, sym_1_1_6, gen29143)

			gen29145 := Call(__e, gen29128, gen29136, gen29144)

			gen29146 := Call(__e, gen29127, gen29145, Nil)

			gen29147 := Call(__e, gen29126, sym_1_1_6, gen29146)

			gen29148 := Call(__e, gen29125, A, gen29147)

			__e.TailApply(gen29124, V3296, gen29148, V3297, V3298)

			return

		}, 1)

		gen29150 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29151 := Call(__e, gen29150, V3297)

		__e.TailApply(gen29149, gen29151)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1assoc, gen29152)

	gen29164 := MakeNative(func(__e *ControlFlow) {
		V3306 := __e.Get(1)
		_ = V3306
		V3307 := __e.Get(2)
		_ = V3307
		V3308 := __e.Get(3)
		_ = V3308
		gen29161 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29153 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29153)

			gen29154 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29157 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29158 := Call(__e, gen29157, symboolean, Nil)

			gen29159 := Call(__e, gen29156, sym_1_1_6, gen29158)

			gen29160 := Call(__e, gen29155, A, gen29159)

			__e.TailApply(gen29154, V3306, gen29160, V3307, V3308)

			return

		}, 1)

		gen29162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29163 := Call(__e, gen29162, V3307)

		__e.TailApply(gen29161, gen29163)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1boolean_2, gen29164)

	gen29173 := MakeNative(func(__e *ControlFlow) {
		V3316 := __e.Get(1)
		_ = V3316
		V3317 := __e.Get(2)
		_ = V3317
		V3318 := __e.Get(3)
		_ = V3318
		gen29165 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29165)

		gen29166 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29167 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29168 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29169 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29170 := Call(__e, gen29169, symboolean, Nil)

		gen29171 := Call(__e, gen29168, sym_1_1_6, gen29170)

		gen29172 := Call(__e, gen29167, symsymbol, gen29171)

		__e.TailApply(gen29166, V3316, gen29172, V3317, V3318)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1bound_2, gen29173)

	gen29182 := MakeNative(func(__e *ControlFlow) {
		V3326 := __e.Get(1)
		_ = V3326
		V3327 := __e.Get(2)
		_ = V3327
		V3328 := __e.Get(3)
		_ = V3328
		gen29174 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29174)

		gen29175 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29177 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29178 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29179 := Call(__e, gen29178, symstring, Nil)

		gen29180 := Call(__e, gen29177, sym_1_1_6, gen29179)

		gen29181 := Call(__e, gen29176, symstring, gen29180)

		__e.TailApply(gen29175, V3326, gen29181, V3327, V3328)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cd, gen29182)

	gen29205 := MakeNative(func(__e *ControlFlow) {
		V3336 := __e.Get(1)
		_ = V3336
		V3337 := __e.Get(2)
		_ = V3337
		V3338 := __e.Get(3)
		_ = V3338
		gen29202 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29199 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29183)

				gen29184 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29185 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29186 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29187 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29188 := Call(__e, gen29187, A, Nil)

				gen29189 := Call(__e, gen29186, symstream, gen29188)

				gen29190 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29191 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29192 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29193 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29194 := Call(__e, gen29193, B, Nil)

				gen29195 := Call(__e, gen29192, symlist, gen29194)

				gen29196 := Call(__e, gen29191, gen29195, Nil)

				gen29197 := Call(__e, gen29190, sym_1_1_6, gen29196)

				gen29198 := Call(__e, gen29185, gen29189, gen29197)

				__e.TailApply(gen29184, V3336, gen29198, V3337, V3338)

				return

			}, 1)

			gen29200 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29201 := Call(__e, gen29200, V3337)

			__e.TailApply(gen29199, gen29201)

			return

		}, 1)

		gen29203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29204 := Call(__e, gen29203, V3337)

		__e.TailApply(gen29202, gen29204)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1close, gen29205)

	gen29220 := MakeNative(func(__e *ControlFlow) {
		V3346 := __e.Get(1)
		_ = V3346
		V3347 := __e.Get(2)
		_ = V3347
		V3348 := __e.Get(3)
		_ = V3348
		gen29206 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29206)

		gen29207 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29208 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29209 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29210 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29211 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29213 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29214 := Call(__e, gen29213, symstring, Nil)

		gen29215 := Call(__e, gen29212, sym_1_1_6, gen29214)

		gen29216 := Call(__e, gen29211, symstring, gen29215)

		gen29217 := Call(__e, gen29210, gen29216, Nil)

		gen29218 := Call(__e, gen29209, sym_1_1_6, gen29217)

		gen29219 := Call(__e, gen29208, symstring, gen29218)

		__e.TailApply(gen29207, V3346, gen29219, V3347, V3348)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cn, gen29220)

	gen29259 := MakeNative(func(__e *ControlFlow) {
		V3356 := __e.Get(1)
		_ = V3356
		V3357 := __e.Get(2)
		_ = V3357
		V3358 := __e.Get(3)
		_ = V3358
		gen29256 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29253 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29221 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29221)

				gen29222 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29223 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29224 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29225 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29226 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29227 := Call(__e, gen29226, B, Nil)

				gen29228 := Call(__e, gen29225, symshen_4_a_a_6, gen29227)

				gen29229 := Call(__e, gen29224, A, gen29228)

				gen29230 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29231 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29232 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29233 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29234 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29235 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29236 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29237 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29238 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29239 := Call(__e, gen29238, B, Nil)

				gen29240 := Call(__e, gen29237, sym_1_1_6, gen29239)

				gen29241 := Call(__e, gen29236, A, gen29240)

				gen29242 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29243 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29244 := Call(__e, gen29243, B, Nil)

				gen29245 := Call(__e, gen29242, sym_1_1_6, gen29244)

				gen29246 := Call(__e, gen29235, gen29241, gen29245)

				gen29247 := Call(__e, gen29234, gen29246, Nil)

				gen29248 := Call(__e, gen29233, sym_1_1_6, gen29247)

				gen29249 := Call(__e, gen29232, A, gen29248)

				gen29250 := Call(__e, gen29231, gen29249, Nil)

				gen29251 := Call(__e, gen29230, sym_1_1_6, gen29250)

				gen29252 := Call(__e, gen29223, gen29229, gen29251)

				__e.TailApply(gen29222, V3356, gen29252, V3357, V3358)

				return

			}, 1)

			gen29254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29255 := Call(__e, gen29254, V3357)

			__e.TailApply(gen29253, gen29255)

			return

		}, 1)

		gen29257 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29258 := Call(__e, gen29257, V3357)

		__e.TailApply(gen29256, gen29258)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1compile, gen29259)

	gen29271 := MakeNative(func(__e *ControlFlow) {
		V3366 := __e.Get(1)
		_ = V3366
		V3367 := __e.Get(2)
		_ = V3367
		V3368 := __e.Get(3)
		_ = V3368
		gen29268 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29260 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29260)

			gen29261 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29265 := Call(__e, gen29264, symboolean, Nil)

			gen29266 := Call(__e, gen29263, sym_1_1_6, gen29265)

			gen29267 := Call(__e, gen29262, A, gen29266)

			__e.TailApply(gen29261, V3366, gen29267, V3367, V3368)

			return

		}, 1)

		gen29269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29270 := Call(__e, gen29269, V3367)

		__e.TailApply(gen29268, gen29270)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cons_2, gen29271)

	gen29292 := MakeNative(func(__e *ControlFlow) {
		V3376 := __e.Get(1)
		_ = V3376
		V3377 := __e.Get(2)
		_ = V3377
		V3378 := __e.Get(3)
		_ = V3378
		gen29289 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29286 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29272 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29272)

				gen29273 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29274 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29275 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29276 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29278 := Call(__e, gen29277, B, Nil)

				gen29279 := Call(__e, gen29276, sym_1_1_6, gen29278)

				gen29280 := Call(__e, gen29275, A, gen29279)

				gen29281 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29282 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29283 := Call(__e, gen29282, symsymbol, Nil)

				gen29284 := Call(__e, gen29281, sym_1_1_6, gen29283)

				gen29285 := Call(__e, gen29274, gen29280, gen29284)

				__e.TailApply(gen29273, V3376, gen29285, V3377, V3378)

				return

			}, 1)

			gen29287 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29288 := Call(__e, gen29287, V3377)

			__e.TailApply(gen29286, gen29288)

			return

		}, 1)

		gen29290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29291 := Call(__e, gen29290, V3377)

		__e.TailApply(gen29289, gen29291)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1destroy, gen29292)

	gen29322 := MakeNative(func(__e *ControlFlow) {
		V3386 := __e.Get(1)
		_ = V3386
		V3387 := __e.Get(2)
		_ = V3387
		V3388 := __e.Get(3)
		_ = V3388
		gen29319 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29293 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29293)

			gen29294 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29295 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29296 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29297 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29298 := Call(__e, gen29297, A, Nil)

			gen29299 := Call(__e, gen29296, symlist, gen29298)

			gen29300 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29301 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29302 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29305 := Call(__e, gen29304, A, Nil)

			gen29306 := Call(__e, gen29303, symlist, gen29305)

			gen29307 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29308 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29309 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29310 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29311 := Call(__e, gen29310, A, Nil)

			gen29312 := Call(__e, gen29309, symlist, gen29311)

			gen29313 := Call(__e, gen29308, gen29312, Nil)

			gen29314 := Call(__e, gen29307, sym_1_1_6, gen29313)

			gen29315 := Call(__e, gen29302, gen29306, gen29314)

			gen29316 := Call(__e, gen29301, gen29315, Nil)

			gen29317 := Call(__e, gen29300, sym_1_1_6, gen29316)

			gen29318 := Call(__e, gen29295, gen29299, gen29317)

			__e.TailApply(gen29294, V3386, gen29318, V3387, V3388)

			return

		}, 1)

		gen29320 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29321 := Call(__e, gen29320, V3387)

		__e.TailApply(gen29319, gen29321)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1difference, gen29322)

	gen29343 := MakeNative(func(__e *ControlFlow) {
		V3396 := __e.Get(1)
		_ = V3396
		V3397 := __e.Get(2)
		_ = V3397
		V3398 := __e.Get(3)
		_ = V3398
		gen29340 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29337 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29323 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29323)

				gen29324 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29327 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29328 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29329 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29330 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29331 := Call(__e, gen29330, B, Nil)

				gen29332 := Call(__e, gen29329, sym_1_1_6, gen29331)

				gen29333 := Call(__e, gen29328, B, gen29332)

				gen29334 := Call(__e, gen29327, gen29333, Nil)

				gen29335 := Call(__e, gen29326, sym_1_1_6, gen29334)

				gen29336 := Call(__e, gen29325, A, gen29335)

				__e.TailApply(gen29324, V3396, gen29336, V3397, V3398)

				return

			}, 1)

			gen29338 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29339 := Call(__e, gen29338, V3397)

			__e.TailApply(gen29337, gen29339)

			return

		}, 1)

		gen29341 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29342 := Call(__e, gen29341, V3397)

		__e.TailApply(gen29340, gen29342)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1do, gen29343)

	gen29366 := MakeNative(func(__e *ControlFlow) {
		V3406 := __e.Get(1)
		_ = V3406
		V3407 := __e.Get(2)
		_ = V3407
		V3408 := __e.Get(3)
		_ = V3408
		gen29363 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29360 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29344 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29344)

				gen29345 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29349 := Call(__e, gen29348, A, Nil)

				gen29350 := Call(__e, gen29347, symlist, gen29349)

				gen29351 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29355 := Call(__e, gen29354, B, Nil)

				gen29356 := Call(__e, gen29353, symlist, gen29355)

				gen29357 := Call(__e, gen29352, gen29356, Nil)

				gen29358 := Call(__e, gen29351, symshen_4_a_a_6, gen29357)

				gen29359 := Call(__e, gen29346, gen29350, gen29358)

				__e.TailApply(gen29345, V3406, gen29359, V3407, V3408)

				return

			}, 1)

			gen29361 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29362 := Call(__e, gen29361, V3407)

			__e.TailApply(gen29360, gen29362)

			return

		}, 1)

		gen29364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29365 := Call(__e, gen29364, V3407)

		__e.TailApply(gen29363, gen29365)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5e_6, gen29366)

	gen29386 := MakeNative(func(__e *ControlFlow) {
		V3416 := __e.Get(1)
		_ = V3416
		V3417 := __e.Get(2)
		_ = V3417
		V3418 := __e.Get(3)
		_ = V3418
		gen29383 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29367 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29367)

			gen29368 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29370 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29371 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29372 := Call(__e, gen29371, A, Nil)

			gen29373 := Call(__e, gen29370, symlist, gen29372)

			gen29374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29376 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29377 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29378 := Call(__e, gen29377, A, Nil)

			gen29379 := Call(__e, gen29376, symlist, gen29378)

			gen29380 := Call(__e, gen29375, gen29379, Nil)

			gen29381 := Call(__e, gen29374, symshen_4_a_a_6, gen29380)

			gen29382 := Call(__e, gen29369, gen29373, gen29381)

			__e.TailApply(gen29368, V3416, gen29382, V3417, V3418)

			return

		}, 1)

		gen29384 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29385 := Call(__e, gen29384, V3417)

		__e.TailApply(gen29383, gen29385)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_b_6, gen29386)

	gen29408 := MakeNative(func(__e *ControlFlow) {
		V3426 := __e.Get(1)
		_ = V3426
		V3427 := __e.Get(2)
		_ = V3427
		V3428 := __e.Get(3)
		_ = V3428
		gen29405 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29387)

			gen29388 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29389 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29391 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29392 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29393 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29394 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29395 := Call(__e, gen29394, A, Nil)

			gen29396 := Call(__e, gen29393, symlist, gen29395)

			gen29397 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29398 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29399 := Call(__e, gen29398, symboolean, Nil)

			gen29400 := Call(__e, gen29397, sym_1_1_6, gen29399)

			gen29401 := Call(__e, gen29392, gen29396, gen29400)

			gen29402 := Call(__e, gen29391, gen29401, Nil)

			gen29403 := Call(__e, gen29390, sym_1_1_6, gen29402)

			gen29404 := Call(__e, gen29389, A, gen29403)

			__e.TailApply(gen29388, V3426, gen29404, V3427, V3428)

			return

		}, 1)

		gen29406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29407 := Call(__e, gen29406, V3427)

		__e.TailApply(gen29405, gen29407)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1element_2, gen29408)

	gen29420 := MakeNative(func(__e *ControlFlow) {
		V3436 := __e.Get(1)
		_ = V3436
		V3437 := __e.Get(2)
		_ = V3437
		V3438 := __e.Get(3)
		_ = V3438
		gen29417 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29409)

			gen29410 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29413 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29414 := Call(__e, gen29413, symboolean, Nil)

			gen29415 := Call(__e, gen29412, sym_1_1_6, gen29414)

			gen29416 := Call(__e, gen29411, A, gen29415)

			__e.TailApply(gen29410, V3436, gen29416, V3437, V3438)

			return

		}, 1)

		gen29418 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29419 := Call(__e, gen29418, V3437)

		__e.TailApply(gen29417, gen29419)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1empty_2, gen29420)

	gen29429 := MakeNative(func(__e *ControlFlow) {
		V3446 := __e.Get(1)
		_ = V3446
		V3447 := __e.Get(2)
		_ = V3447
		V3448 := __e.Get(3)
		_ = V3448
		gen29421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29421)

		gen29422 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29423 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29424 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29425 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29426 := Call(__e, gen29425, symboolean, Nil)

		gen29427 := Call(__e, gen29424, sym_1_1_6, gen29426)

		gen29428 := Call(__e, gen29423, symsymbol, gen29427)

		__e.TailApply(gen29422, V3446, gen29428, V3447, V3448)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1enable_1type_1theory, gen29429)

	gen29442 := MakeNative(func(__e *ControlFlow) {
		V3456 := __e.Get(1)
		_ = V3456
		V3457 := __e.Get(2)
		_ = V3457
		V3458 := __e.Get(3)
		_ = V3458
		gen29430 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29430)

		gen29431 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29432 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29434 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29435 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29436 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29437 := Call(__e, gen29436, symsymbol, Nil)

		gen29438 := Call(__e, gen29435, symlist, gen29437)

		gen29439 := Call(__e, gen29434, gen29438, Nil)

		gen29440 := Call(__e, gen29433, sym_1_1_6, gen29439)

		gen29441 := Call(__e, gen29432, symsymbol, gen29440)

		__e.TailApply(gen29431, V3456, gen29441, V3457, V3458)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1external, gen29442)

	gen29451 := MakeNative(func(__e *ControlFlow) {
		V3466 := __e.Get(1)
		_ = V3466
		V3467 := __e.Get(2)
		_ = V3467
		V3468 := __e.Get(3)
		_ = V3468
		gen29443 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29443)

		gen29444 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29448 := Call(__e, gen29447, symstring, Nil)

		gen29449 := Call(__e, gen29446, sym_1_1_6, gen29448)

		gen29450 := Call(__e, gen29445, symexception, gen29449)

		__e.TailApply(gen29444, V3466, gen29450, V3467, V3468)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1error_1to_1string, gen29451)

	gen29467 := MakeNative(func(__e *ControlFlow) {
		V3476 := __e.Get(1)
		_ = V3476
		V3477 := __e.Get(2)
		_ = V3477
		V3478 := __e.Get(3)
		_ = V3478
		gen29464 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29452)

			gen29453 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29455 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29456 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29457 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29458 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29459 := Call(__e, gen29458, symstring, Nil)

			gen29460 := Call(__e, gen29457, symlist, gen29459)

			gen29461 := Call(__e, gen29456, gen29460, Nil)

			gen29462 := Call(__e, gen29455, sym_1_1_6, gen29461)

			gen29463 := Call(__e, gen29454, A, gen29462)

			__e.TailApply(gen29453, V3476, gen29463, V3477, V3478)

			return

		}, 1)

		gen29465 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29466 := Call(__e, gen29465, V3477)

		__e.TailApply(gen29464, gen29466)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1explode, gen29467)

	gen29474 := MakeNative(func(__e *ControlFlow) {
		V3486 := __e.Get(1)
		_ = V3486
		V3487 := __e.Get(2)
		_ = V3487
		V3488 := __e.Get(3)
		_ = V3488
		gen29468 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29468)

		gen29469 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29470 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29471 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29472 := Call(__e, gen29471, symsymbol, Nil)

		gen29473 := Call(__e, gen29470, sym_1_1_6, gen29472)

		__e.TailApply(gen29469, V3486, gen29473, V3487, V3488)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fail, gen29474)

	gen29495 := MakeNative(func(__e *ControlFlow) {
		V3496 := __e.Get(1)
		_ = V3496
		V3497 := __e.Get(2)
		_ = V3497
		V3498 := __e.Get(3)
		_ = V3498
		gen29475 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29475)

		gen29476 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29477 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29478 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29479 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29480 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29481 := Call(__e, gen29480, symboolean, Nil)

		gen29482 := Call(__e, gen29479, sym_1_1_6, gen29481)

		gen29483 := Call(__e, gen29478, symsymbol, gen29482)

		gen29484 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29485 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29486 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29487 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29488 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29489 := Call(__e, gen29488, symsymbol, Nil)

		gen29490 := Call(__e, gen29487, sym_1_1_6, gen29489)

		gen29491 := Call(__e, gen29486, symsymbol, gen29490)

		gen29492 := Call(__e, gen29485, gen29491, Nil)

		gen29493 := Call(__e, gen29484, sym_1_1_6, gen29492)

		gen29494 := Call(__e, gen29477, gen29483, gen29493)

		__e.TailApply(gen29476, V3496, gen29494, V3497, V3498)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fail_1if, gen29495)

	gen29519 := MakeNative(func(__e *ControlFlow) {
		V3506 := __e.Get(1)
		_ = V3506
		V3507 := __e.Get(2)
		_ = V3507
		V3508 := __e.Get(3)
		_ = V3508
		gen29516 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29496)

			gen29497 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29498 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29499 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29500 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29501 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29502 := Call(__e, gen29501, A, Nil)

			gen29503 := Call(__e, gen29500, sym_1_1_6, gen29502)

			gen29504 := Call(__e, gen29499, A, gen29503)

			gen29505 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29506 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29507 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29508 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29509 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29510 := Call(__e, gen29509, A, Nil)

			gen29511 := Call(__e, gen29508, sym_1_1_6, gen29510)

			gen29512 := Call(__e, gen29507, A, gen29511)

			gen29513 := Call(__e, gen29506, gen29512, Nil)

			gen29514 := Call(__e, gen29505, sym_1_1_6, gen29513)

			gen29515 := Call(__e, gen29498, gen29504, gen29514)

			__e.TailApply(gen29497, V3506, gen29515, V3507, V3508)

			return

		}, 1)

		gen29517 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29518 := Call(__e, gen29517, V3507)

		__e.TailApply(gen29516, gen29518)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fix, gen29519)

	gen29535 := MakeNative(func(__e *ControlFlow) {
		V3516 := __e.Get(1)
		_ = V3516
		V3517 := __e.Get(2)
		_ = V3517
		V3518 := __e.Get(3)
		_ = V3518
		gen29532 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29520 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29520)

			gen29521 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29522 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29523 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29524 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29525 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29526 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29527 := Call(__e, gen29526, A, Nil)

			gen29528 := Call(__e, gen29525, symlazy, gen29527)

			gen29529 := Call(__e, gen29524, gen29528, Nil)

			gen29530 := Call(__e, gen29523, sym_1_1_6, gen29529)

			gen29531 := Call(__e, gen29522, A, gen29530)

			__e.TailApply(gen29521, V3516, gen29531, V3517, V3518)

			return

		}, 1)

		gen29533 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29534 := Call(__e, gen29533, V3517)

		__e.TailApply(gen29532, gen29534)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1freeze, gen29535)

	gen29556 := MakeNative(func(__e *ControlFlow) {
		V3526 := __e.Get(1)
		_ = V3526
		V3527 := __e.Get(2)
		_ = V3527
		V3528 := __e.Get(3)
		_ = V3528
		gen29553 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			gen29550 := MakeNative(func(__e *ControlFlow) {
				A := __e.Get(1)
				_ = A
				gen29536 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29536)

				gen29537 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29538 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29539 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29541 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29542 := Call(__e, gen29541, B, Nil)

				gen29543 := Call(__e, gen29540, sym_d, gen29542)

				gen29544 := Call(__e, gen29539, A, gen29543)

				gen29545 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29546 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29547 := Call(__e, gen29546, A, Nil)

				gen29548 := Call(__e, gen29545, sym_1_1_6, gen29547)

				gen29549 := Call(__e, gen29538, gen29544, gen29548)

				__e.TailApply(gen29537, V3526, gen29549, V3527, V3528)

				return

			}, 1)

			gen29551 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29552 := Call(__e, gen29551, V3527)

			__e.TailApply(gen29550, gen29552)

			return

		}, 1)

		gen29554 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29555 := Call(__e, gen29554, V3527)

		__e.TailApply(gen29553, gen29555)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fst, gen29556)

	gen29583 := MakeNative(func(__e *ControlFlow) {
		V3536 := __e.Get(1)
		_ = V3536
		V3537 := __e.Get(2)
		_ = V3537
		V3538 := __e.Get(3)
		_ = V3538
		gen29580 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29577 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29557)

				gen29558 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29559 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29561 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29562 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29563 := Call(__e, gen29562, B, Nil)

				gen29564 := Call(__e, gen29561, sym_1_1_6, gen29563)

				gen29565 := Call(__e, gen29560, A, gen29564)

				gen29566 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29567 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29569 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29570 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29571 := Call(__e, gen29570, B, Nil)

				gen29572 := Call(__e, gen29569, sym_1_1_6, gen29571)

				gen29573 := Call(__e, gen29568, A, gen29572)

				gen29574 := Call(__e, gen29567, gen29573, Nil)

				gen29575 := Call(__e, gen29566, sym_1_1_6, gen29574)

				gen29576 := Call(__e, gen29559, gen29565, gen29575)

				__e.TailApply(gen29558, V3536, gen29576, V3537, V3538)

				return

			}, 1)

			gen29578 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29579 := Call(__e, gen29578, V3537)

			__e.TailApply(gen29577, gen29579)

			return

		}, 1)

		gen29581 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29582 := Call(__e, gen29581, V3537)

		__e.TailApply(gen29580, gen29582)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1function, gen29583)

	gen29592 := MakeNative(func(__e *ControlFlow) {
		V3546 := __e.Get(1)
		_ = V3546
		V3547 := __e.Get(2)
		_ = V3547
		V3548 := __e.Get(3)
		_ = V3548
		gen29584 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29584)

		gen29585 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29586 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29587 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29588 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29589 := Call(__e, gen29588, symsymbol, Nil)

		gen29590 := Call(__e, gen29587, sym_1_1_6, gen29589)

		gen29591 := Call(__e, gen29586, symsymbol, gen29590)

		__e.TailApply(gen29585, V3546, gen29591, V3547, V3548)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1gensym, gen29592)

	gen29614 := MakeNative(func(__e *ControlFlow) {
		V3556 := __e.Get(1)
		_ = V3556
		V3557 := __e.Get(2)
		_ = V3557
		V3558 := __e.Get(3)
		_ = V3558
		gen29611 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29593)

			gen29594 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29595 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29596 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29598 := Call(__e, gen29597, A, Nil)

			gen29599 := Call(__e, gen29596, symvector, gen29598)

			gen29600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29601 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29602 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29603 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29604 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29605 := Call(__e, gen29604, A, Nil)

			gen29606 := Call(__e, gen29603, sym_1_1_6, gen29605)

			gen29607 := Call(__e, gen29602, symnumber, gen29606)

			gen29608 := Call(__e, gen29601, gen29607, Nil)

			gen29609 := Call(__e, gen29600, sym_1_1_6, gen29608)

			gen29610 := Call(__e, gen29595, gen29599, gen29609)

			__e.TailApply(gen29594, V3556, gen29610, V3557, V3558)

			return

		}, 1)

		gen29612 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29613 := Call(__e, gen29612, V3557)

		__e.TailApply(gen29611, gen29613)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_1vector, gen29614)

	gen29646 := MakeNative(func(__e *ControlFlow) {
		V3566 := __e.Get(1)
		_ = V3566
		V3567 := __e.Get(2)
		_ = V3567
		V3568 := __e.Get(3)
		_ = V3568
		gen29643 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29615 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29615)

			gen29616 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29617 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29618 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29619 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29620 := Call(__e, gen29619, A, Nil)

			gen29621 := Call(__e, gen29618, symvector, gen29620)

			gen29622 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29625 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29626 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29629 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29630 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29631 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29632 := Call(__e, gen29631, A, Nil)

			gen29633 := Call(__e, gen29630, symvector, gen29632)

			gen29634 := Call(__e, gen29629, gen29633, Nil)

			gen29635 := Call(__e, gen29628, sym_1_1_6, gen29634)

			gen29636 := Call(__e, gen29627, A, gen29635)

			gen29637 := Call(__e, gen29626, gen29636, Nil)

			gen29638 := Call(__e, gen29625, sym_1_1_6, gen29637)

			gen29639 := Call(__e, gen29624, symnumber, gen29638)

			gen29640 := Call(__e, gen29623, gen29639, Nil)

			gen29641 := Call(__e, gen29622, sym_1_1_6, gen29640)

			gen29642 := Call(__e, gen29617, gen29621, gen29641)

			__e.TailApply(gen29616, V3566, gen29642, V3567, V3568)

			return

		}, 1)

		gen29644 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29645 := Call(__e, gen29644, V3567)

		__e.TailApply(gen29643, gen29645)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector_1_6, gen29646)

	gen29662 := MakeNative(func(__e *ControlFlow) {
		V3576 := __e.Get(1)
		_ = V3576
		V3577 := __e.Get(2)
		_ = V3577
		V3578 := __e.Get(3)
		_ = V3578
		gen29659 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29647 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29647)

			gen29648 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29649 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29650 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29651 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29652 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29653 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29654 := Call(__e, gen29653, A, Nil)

			gen29655 := Call(__e, gen29652, symvector, gen29654)

			gen29656 := Call(__e, gen29651, gen29655, Nil)

			gen29657 := Call(__e, gen29650, sym_1_1_6, gen29656)

			gen29658 := Call(__e, gen29649, symnumber, gen29657)

			__e.TailApply(gen29648, V3576, gen29658, V3577, V3578)

			return

		}, 1)

		gen29660 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29661 := Call(__e, gen29660, V3577)

		__e.TailApply(gen29659, gen29661)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector, gen29662)

	gen29671 := MakeNative(func(__e *ControlFlow) {
		V3586 := __e.Get(1)
		_ = V3586
		V3587 := __e.Get(2)
		_ = V3587
		V3588 := __e.Get(3)
		_ = V3588
		gen29663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29663)

		gen29664 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29668 := Call(__e, gen29667, symnumber, Nil)

		gen29669 := Call(__e, gen29666, sym_1_1_6, gen29668)

		gen29670 := Call(__e, gen29665, symsymbol, gen29669)

		__e.TailApply(gen29664, V3586, gen29670, V3587, V3588)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1get_1time, gen29671)

	gen29689 := MakeNative(func(__e *ControlFlow) {
		V3596 := __e.Get(1)
		_ = V3596
		V3597 := __e.Get(2)
		_ = V3597
		V3598 := __e.Get(3)
		_ = V3598
		gen29686 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29672 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29672)

			gen29673 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29674 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29675 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29679 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29680 := Call(__e, gen29679, symnumber, Nil)

			gen29681 := Call(__e, gen29678, sym_1_1_6, gen29680)

			gen29682 := Call(__e, gen29677, symnumber, gen29681)

			gen29683 := Call(__e, gen29676, gen29682, Nil)

			gen29684 := Call(__e, gen29675, sym_1_1_6, gen29683)

			gen29685 := Call(__e, gen29674, A, gen29684)

			__e.TailApply(gen29673, V3596, gen29685, V3597, V3598)

			return

		}, 1)

		gen29687 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29688 := Call(__e, gen29687, V3597)

		__e.TailApply(gen29686, gen29688)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hash, gen29689)

	gen29705 := MakeNative(func(__e *ControlFlow) {
		V3606 := __e.Get(1)
		_ = V3606
		V3607 := __e.Get(2)
		_ = V3607
		V3608 := __e.Get(3)
		_ = V3608
		gen29702 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29690 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29690)

			gen29691 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29693 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29694 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29695 := Call(__e, gen29694, A, Nil)

			gen29696 := Call(__e, gen29693, symlist, gen29695)

			gen29697 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29698 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29699 := Call(__e, gen29698, A, Nil)

			gen29700 := Call(__e, gen29697, sym_1_1_6, gen29699)

			gen29701 := Call(__e, gen29692, gen29696, gen29700)

			__e.TailApply(gen29691, V3606, gen29701, V3607, V3608)

			return

		}, 1)

		gen29703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29704 := Call(__e, gen29703, V3607)

		__e.TailApply(gen29702, gen29704)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1head, gen29705)

	gen29721 := MakeNative(func(__e *ControlFlow) {
		V3616 := __e.Get(1)
		_ = V3616
		V3617 := __e.Get(2)
		_ = V3617
		V3618 := __e.Get(3)
		_ = V3618
		gen29718 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29706 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29706)

			gen29707 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29709 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29710 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29711 := Call(__e, gen29710, A, Nil)

			gen29712 := Call(__e, gen29709, symvector, gen29711)

			gen29713 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29714 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29715 := Call(__e, gen29714, A, Nil)

			gen29716 := Call(__e, gen29713, sym_1_1_6, gen29715)

			gen29717 := Call(__e, gen29708, gen29712, gen29716)

			__e.TailApply(gen29707, V3616, gen29717, V3617, V3618)

			return

		}, 1)

		gen29719 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29720 := Call(__e, gen29719, V3617)

		__e.TailApply(gen29718, gen29720)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hdv, gen29721)

	gen29730 := MakeNative(func(__e *ControlFlow) {
		V3626 := __e.Get(1)
		_ = V3626
		V3627 := __e.Get(2)
		_ = V3627
		V3628 := __e.Get(3)
		_ = V3628
		gen29722 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29722)

		gen29723 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29724 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29725 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29726 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29727 := Call(__e, gen29726, symstring, Nil)

		gen29728 := Call(__e, gen29725, sym_1_1_6, gen29727)

		gen29729 := Call(__e, gen29724, symstring, gen29728)

		__e.TailApply(gen29723, V3626, gen29729, V3627, V3628)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hdstr, gen29730)

	gen29754 := MakeNative(func(__e *ControlFlow) {
		V3636 := __e.Get(1)
		_ = V3636
		V3637 := __e.Get(2)
		_ = V3637
		V3638 := __e.Get(3)
		_ = V3638
		gen29751 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29731 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29731)

			gen29732 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29733 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29734 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29736 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29737 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29739 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29742 := Call(__e, gen29741, A, Nil)

			gen29743 := Call(__e, gen29740, sym_1_1_6, gen29742)

			gen29744 := Call(__e, gen29739, A, gen29743)

			gen29745 := Call(__e, gen29738, gen29744, Nil)

			gen29746 := Call(__e, gen29737, sym_1_1_6, gen29745)

			gen29747 := Call(__e, gen29736, A, gen29746)

			gen29748 := Call(__e, gen29735, gen29747, Nil)

			gen29749 := Call(__e, gen29734, sym_1_1_6, gen29748)

			gen29750 := Call(__e, gen29733, symboolean, gen29749)

			__e.TailApply(gen29732, V3636, gen29750, V3637, V3638)

			return

		}, 1)

		gen29752 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29753 := Call(__e, gen29752, V3637)

		__e.TailApply(gen29751, gen29753)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1if, gen29754)

	gen29761 := MakeNative(func(__e *ControlFlow) {
		V3646 := __e.Get(1)
		_ = V3646
		V3647 := __e.Get(2)
		_ = V3647
		V3648 := __e.Get(3)
		_ = V3648
		gen29755 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29755)

		gen29756 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29758 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29759 := Call(__e, gen29758, symstring, Nil)

		gen29760 := Call(__e, gen29757, sym_1_1_6, gen29759)

		__e.TailApply(gen29756, V3646, gen29760, V3647, V3648)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1it, gen29761)

	gen29768 := MakeNative(func(__e *ControlFlow) {
		V3656 := __e.Get(1)
		_ = V3656
		V3657 := __e.Get(2)
		_ = V3657
		V3658 := __e.Get(3)
		_ = V3658
		gen29762 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29762)

		gen29763 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29764 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29765 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29766 := Call(__e, gen29765, symstring, Nil)

		gen29767 := Call(__e, gen29764, sym_1_1_6, gen29766)

		__e.TailApply(gen29763, V3656, gen29767, V3657, V3658)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1implementation, gen29768)

	gen29785 := MakeNative(func(__e *ControlFlow) {
		V3666 := __e.Get(1)
		_ = V3666
		V3667 := __e.Get(2)
		_ = V3667
		V3668 := __e.Get(3)
		_ = V3668
		gen29769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29769)

		gen29770 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29772 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29773 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29774 := Call(__e, gen29773, symsymbol, Nil)

		gen29775 := Call(__e, gen29772, symlist, gen29774)

		gen29776 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29777 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29778 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29779 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29780 := Call(__e, gen29779, symsymbol, Nil)

		gen29781 := Call(__e, gen29778, symlist, gen29780)

		gen29782 := Call(__e, gen29777, gen29781, Nil)

		gen29783 := Call(__e, gen29776, sym_1_1_6, gen29782)

		gen29784 := Call(__e, gen29771, gen29775, gen29783)

		__e.TailApply(gen29770, V3666, gen29784, V3667, V3668)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1include, gen29785)

	gen29802 := MakeNative(func(__e *ControlFlow) {
		V3676 := __e.Get(1)
		_ = V3676
		V3677 := __e.Get(2)
		_ = V3677
		V3678 := __e.Get(3)
		_ = V3678
		gen29786 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29786)

		gen29787 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29788 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29789 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29790 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29791 := Call(__e, gen29790, symsymbol, Nil)

		gen29792 := Call(__e, gen29789, symlist, gen29791)

		gen29793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29794 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29795 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29796 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29797 := Call(__e, gen29796, symsymbol, Nil)

		gen29798 := Call(__e, gen29795, symlist, gen29797)

		gen29799 := Call(__e, gen29794, gen29798, Nil)

		gen29800 := Call(__e, gen29793, sym_1_1_6, gen29799)

		gen29801 := Call(__e, gen29788, gen29792, gen29800)

		__e.TailApply(gen29787, V3676, gen29801, V3677, V3678)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1include_1all_1but, gen29802)

	gen29809 := MakeNative(func(__e *ControlFlow) {
		V3686 := __e.Get(1)
		_ = V3686
		V3687 := __e.Get(2)
		_ = V3687
		V3688 := __e.Get(3)
		_ = V3688
		gen29803 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29803)

		gen29804 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29807 := Call(__e, gen29806, symnumber, Nil)

		gen29808 := Call(__e, gen29805, sym_1_1_6, gen29807)

		__e.TailApply(gen29804, V3686, gen29808, V3687, V3688)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1inferences, gen29809)

	gen29827 := MakeNative(func(__e *ControlFlow) {
		V3696 := __e.Get(1)
		_ = V3696
		V3697 := __e.Get(2)
		_ = V3697
		V3698 := __e.Get(3)
		_ = V3698
		gen29824 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29810 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29810)

			gen29811 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29812 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29813 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29814 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29815 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29817 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29818 := Call(__e, gen29817, symstring, Nil)

			gen29819 := Call(__e, gen29816, sym_1_1_6, gen29818)

			gen29820 := Call(__e, gen29815, symstring, gen29819)

			gen29821 := Call(__e, gen29814, gen29820, Nil)

			gen29822 := Call(__e, gen29813, sym_1_1_6, gen29821)

			gen29823 := Call(__e, gen29812, A, gen29822)

			__e.TailApply(gen29811, V3696, gen29823, V3697, V3698)

			return

		}, 1)

		gen29825 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29826 := Call(__e, gen29825, V3697)

		__e.TailApply(gen29824, gen29826)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4insert, gen29827)

	gen29839 := MakeNative(func(__e *ControlFlow) {
		V3706 := __e.Get(1)
		_ = V3706
		V3707 := __e.Get(2)
		_ = V3707
		V3708 := __e.Get(3)
		_ = V3708
		gen29836 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29828)

			gen29829 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29830 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29831 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29832 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29833 := Call(__e, gen29832, symboolean, Nil)

			gen29834 := Call(__e, gen29831, sym_1_1_6, gen29833)

			gen29835 := Call(__e, gen29830, A, gen29834)

			__e.TailApply(gen29829, V3706, gen29835, V3707, V3708)

			return

		}, 1)

		gen29837 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29838 := Call(__e, gen29837, V3707)

		__e.TailApply(gen29836, gen29838)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1integer_2, gen29839)

	gen29852 := MakeNative(func(__e *ControlFlow) {
		V3716 := __e.Get(1)
		_ = V3716
		V3717 := __e.Get(2)
		_ = V3717
		V3718 := __e.Get(3)
		_ = V3718
		gen29840 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29840)

		gen29841 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29842 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29843 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29844 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29845 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29846 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29847 := Call(__e, gen29846, symsymbol, Nil)

		gen29848 := Call(__e, gen29845, symlist, gen29847)

		gen29849 := Call(__e, gen29844, gen29848, Nil)

		gen29850 := Call(__e, gen29843, sym_1_1_6, gen29849)

		gen29851 := Call(__e, gen29842, symsymbol, gen29850)

		__e.TailApply(gen29841, V3716, gen29851, V3717, V3718)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1internal, gen29852)

	gen29882 := MakeNative(func(__e *ControlFlow) {
		V3726 := __e.Get(1)
		_ = V3726
		V3727 := __e.Get(2)
		_ = V3727
		V3728 := __e.Get(3)
		_ = V3728
		gen29879 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29853 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29853)

			gen29854 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29855 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29856 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29857 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29858 := Call(__e, gen29857, A, Nil)

			gen29859 := Call(__e, gen29856, symlist, gen29858)

			gen29860 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29861 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29863 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29864 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29865 := Call(__e, gen29864, A, Nil)

			gen29866 := Call(__e, gen29863, symlist, gen29865)

			gen29867 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29868 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29869 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29870 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29871 := Call(__e, gen29870, A, Nil)

			gen29872 := Call(__e, gen29869, symlist, gen29871)

			gen29873 := Call(__e, gen29868, gen29872, Nil)

			gen29874 := Call(__e, gen29867, sym_1_1_6, gen29873)

			gen29875 := Call(__e, gen29862, gen29866, gen29874)

			gen29876 := Call(__e, gen29861, gen29875, Nil)

			gen29877 := Call(__e, gen29860, sym_1_1_6, gen29876)

			gen29878 := Call(__e, gen29855, gen29859, gen29877)

			__e.TailApply(gen29854, V3726, gen29878, V3727, V3728)

			return

		}, 1)

		gen29880 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29881 := Call(__e, gen29880, V3727)

		__e.TailApply(gen29879, gen29881)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1intersection, gen29882)

	gen29892 := MakeNative(func(__e *ControlFlow) {
		V3736 := __e.Get(1)
		_ = V3736
		V3737 := __e.Get(2)
		_ = V3737
		V3738 := __e.Get(3)
		_ = V3738
		gen29889 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29883)

			gen29884 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29885 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29886 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29887 := Call(__e, gen29886, A, Nil)

			gen29888 := Call(__e, gen29885, sym_1_1_6, gen29887)

			__e.TailApply(gen29884, V3736, gen29888, V3737, V3738)

			return

		}, 1)

		gen29890 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29891 := Call(__e, gen29890, V3737)

		__e.TailApply(gen29889, gen29891)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1kill, gen29892)

	gen29899 := MakeNative(func(__e *ControlFlow) {
		V3746 := __e.Get(1)
		_ = V3746
		V3747 := __e.Get(2)
		_ = V3747
		V3748 := __e.Get(3)
		_ = V3748
		gen29893 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29893)

		gen29894 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29896 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29897 := Call(__e, gen29896, symstring, Nil)

		gen29898 := Call(__e, gen29895, sym_1_1_6, gen29897)

		__e.TailApply(gen29894, V3746, gen29898, V3747, V3748)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1language, gen29899)

	gen29915 := MakeNative(func(__e *ControlFlow) {
		V3756 := __e.Get(1)
		_ = V3756
		V3757 := __e.Get(2)
		_ = V3757
		V3758 := __e.Get(3)
		_ = V3758
		gen29912 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29900 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29900)

			gen29901 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29902 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29904 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29905 := Call(__e, gen29904, A, Nil)

			gen29906 := Call(__e, gen29903, symlist, gen29905)

			gen29907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29909 := Call(__e, gen29908, symnumber, Nil)

			gen29910 := Call(__e, gen29907, sym_1_1_6, gen29909)

			gen29911 := Call(__e, gen29902, gen29906, gen29910)

			__e.TailApply(gen29901, V3756, gen29911, V3757, V3758)

			return

		}, 1)

		gen29913 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29914 := Call(__e, gen29913, V3757)

		__e.TailApply(gen29912, gen29914)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1length, gen29915)

	gen29931 := MakeNative(func(__e *ControlFlow) {
		V3766 := __e.Get(1)
		_ = V3766
		V3767 := __e.Get(2)
		_ = V3767
		V3768 := __e.Get(3)
		_ = V3768
		gen29928 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29916 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen29916)

			gen29917 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen29918 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29919 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29920 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29921 := Call(__e, gen29920, A, Nil)

			gen29922 := Call(__e, gen29919, symvector, gen29921)

			gen29923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen29925 := Call(__e, gen29924, symnumber, Nil)

			gen29926 := Call(__e, gen29923, sym_1_1_6, gen29925)

			gen29927 := Call(__e, gen29918, gen29922, gen29926)

			__e.TailApply(gen29917, V3766, gen29927, V3767, V3768)

			return

		}, 1)

		gen29929 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29930 := Call(__e, gen29929, V3767)

		__e.TailApply(gen29928, gen29930)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1limit, gen29931)

	gen29940 := MakeNative(func(__e *ControlFlow) {
		V3776 := __e.Get(1)
		_ = V3776
		V3777 := __e.Get(2)
		_ = V3777
		V3778 := __e.Get(3)
		_ = V3778
		gen29932 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen29932)

		gen29933 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen29934 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen29937 := Call(__e, gen29936, symsymbol, Nil)

		gen29938 := Call(__e, gen29935, sym_1_1_6, gen29937)

		gen29939 := Call(__e, gen29934, symstring, gen29938)

		__e.TailApply(gen29933, V3776, gen29939, V3777, V3778)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1load, gen29940)

	gen29975 := MakeNative(func(__e *ControlFlow) {
		V3786 := __e.Get(1)
		_ = V3786
		V3787 := __e.Get(2)
		_ = V3787
		V3788 := __e.Get(3)
		_ = V3788
		gen29972 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen29969 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29941 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29941)

				gen29942 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29943 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29944 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29945 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29946 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29947 := Call(__e, gen29946, B, Nil)

				gen29948 := Call(__e, gen29945, sym_1_1_6, gen29947)

				gen29949 := Call(__e, gen29944, A, gen29948)

				gen29950 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29951 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29952 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29953 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29954 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29955 := Call(__e, gen29954, A, Nil)

				gen29956 := Call(__e, gen29953, symlist, gen29955)

				gen29957 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29958 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29959 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29960 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29961 := Call(__e, gen29960, B, Nil)

				gen29962 := Call(__e, gen29959, symlist, gen29961)

				gen29963 := Call(__e, gen29958, gen29962, Nil)

				gen29964 := Call(__e, gen29957, sym_1_1_6, gen29963)

				gen29965 := Call(__e, gen29952, gen29956, gen29964)

				gen29966 := Call(__e, gen29951, gen29965, Nil)

				gen29967 := Call(__e, gen29950, sym_1_1_6, gen29966)

				gen29968 := Call(__e, gen29943, gen29949, gen29967)

				__e.TailApply(gen29942, V3786, gen29968, V3787, V3788)

				return

			}, 1)

			gen29970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen29971 := Call(__e, gen29970, V3787)

			__e.TailApply(gen29969, gen29971)

			return

		}, 1)

		gen29973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen29974 := Call(__e, gen29973, V3787)

		__e.TailApply(gen29972, gen29974)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1map, gen29975)

	gen30014 := MakeNative(func(__e *ControlFlow) {
		V3796 := __e.Get(1)
		_ = V3796
		V3797 := __e.Get(2)
		_ = V3797
		V3798 := __e.Get(3)
		_ = V3798
		gen30011 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30008 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen29976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen29976)

				gen29977 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen29978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29979 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29980 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29981 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29982 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29984 := Call(__e, gen29983, B, Nil)

				gen29985 := Call(__e, gen29982, symlist, gen29984)

				gen29986 := Call(__e, gen29981, gen29985, Nil)

				gen29987 := Call(__e, gen29980, sym_1_1_6, gen29986)

				gen29988 := Call(__e, gen29979, A, gen29987)

				gen29989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29991 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29992 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29993 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29994 := Call(__e, gen29993, A, Nil)

				gen29995 := Call(__e, gen29992, symlist, gen29994)

				gen29996 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29997 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29998 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen29999 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30000 := Call(__e, gen29999, B, Nil)

				gen30001 := Call(__e, gen29998, symlist, gen30000)

				gen30002 := Call(__e, gen29997, gen30001, Nil)

				gen30003 := Call(__e, gen29996, sym_1_1_6, gen30002)

				gen30004 := Call(__e, gen29991, gen29995, gen30003)

				gen30005 := Call(__e, gen29990, gen30004, Nil)

				gen30006 := Call(__e, gen29989, sym_1_1_6, gen30005)

				gen30007 := Call(__e, gen29978, gen29988, gen30006)

				__e.TailApply(gen29977, V3796, gen30007, V3797, V3798)

				return

			}, 1)

			gen30009 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen30010 := Call(__e, gen30009, V3797)

			__e.TailApply(gen30008, gen30010)

			return

		}, 1)

		gen30012 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30013 := Call(__e, gen30012, V3797)

		__e.TailApply(gen30011, gen30013)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1mapcan, gen30014)

	gen30023 := MakeNative(func(__e *ControlFlow) {
		V3806 := __e.Get(1)
		_ = V3806
		V3807 := __e.Get(2)
		_ = V3807
		V3808 := __e.Get(3)
		_ = V3808
		gen30015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30015)

		gen30016 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30020 := Call(__e, gen30019, symnumber, Nil)

		gen30021 := Call(__e, gen30018, sym_1_1_6, gen30020)

		gen30022 := Call(__e, gen30017, symnumber, gen30021)

		__e.TailApply(gen30016, V3806, gen30022, V3807, V3808)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1maxinferences, gen30023)

	gen30032 := MakeNative(func(__e *ControlFlow) {
		V3816 := __e.Get(1)
		_ = V3816
		V3817 := __e.Get(2)
		_ = V3817
		V3818 := __e.Get(3)
		_ = V3818
		gen30024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30024)

		gen30025 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30027 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30028 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30029 := Call(__e, gen30028, symstring, Nil)

		gen30030 := Call(__e, gen30027, sym_1_1_6, gen30029)

		gen30031 := Call(__e, gen30026, symnumber, gen30030)

		__e.TailApply(gen30025, V3816, gen30031, V3817, V3818)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1n_1_6string, gen30032)

	gen30041 := MakeNative(func(__e *ControlFlow) {
		V3826 := __e.Get(1)
		_ = V3826
		V3827 := __e.Get(2)
		_ = V3827
		V3828 := __e.Get(3)
		_ = V3828
		gen30033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30033)

		gen30034 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30035 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30036 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30037 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30038 := Call(__e, gen30037, symnumber, Nil)

		gen30039 := Call(__e, gen30036, sym_1_1_6, gen30038)

		gen30040 := Call(__e, gen30035, symnumber, gen30039)

		__e.TailApply(gen30034, V3826, gen30040, V3827, V3828)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1nl, gen30041)

	gen30050 := MakeNative(func(__e *ControlFlow) {
		V3836 := __e.Get(1)
		_ = V3836
		V3837 := __e.Get(2)
		_ = V3837
		V3838 := __e.Get(3)
		_ = V3838
		gen30042 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30042)

		gen30043 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30047 := Call(__e, gen30046, symboolean, Nil)

		gen30048 := Call(__e, gen30045, sym_1_1_6, gen30047)

		gen30049 := Call(__e, gen30044, symboolean, gen30048)

		__e.TailApply(gen30043, V3836, gen30049, V3837, V3838)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1not, gen30050)

	gen30072 := MakeNative(func(__e *ControlFlow) {
		V3846 := __e.Get(1)
		_ = V3846
		V3847 := __e.Get(2)
		_ = V3847
		V3848 := __e.Get(3)
		_ = V3848
		gen30069 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30051 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30051)

			gen30052 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30054 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30055 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30057 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30059 := Call(__e, gen30058, A, Nil)

			gen30060 := Call(__e, gen30057, symlist, gen30059)

			gen30061 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30063 := Call(__e, gen30062, A, Nil)

			gen30064 := Call(__e, gen30061, sym_1_1_6, gen30063)

			gen30065 := Call(__e, gen30056, gen30060, gen30064)

			gen30066 := Call(__e, gen30055, gen30065, Nil)

			gen30067 := Call(__e, gen30054, sym_1_1_6, gen30066)

			gen30068 := Call(__e, gen30053, symnumber, gen30067)

			__e.TailApply(gen30052, V3846, gen30068, V3847, V3848)

			return

		}, 1)

		gen30070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30071 := Call(__e, gen30070, V3847)

		__e.TailApply(gen30069, gen30071)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1nth, gen30072)

	gen30084 := MakeNative(func(__e *ControlFlow) {
		V3856 := __e.Get(1)
		_ = V3856
		V3857 := __e.Get(2)
		_ = V3857
		V3858 := __e.Get(3)
		_ = V3858
		gen30081 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30073)

			gen30074 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30075 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30076 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30077 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30078 := Call(__e, gen30077, symboolean, Nil)

			gen30079 := Call(__e, gen30076, sym_1_1_6, gen30078)

			gen30080 := Call(__e, gen30075, A, gen30079)

			__e.TailApply(gen30074, V3856, gen30080, V3857, V3858)

			return

		}, 1)

		gen30082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30083 := Call(__e, gen30082, V3857)

		__e.TailApply(gen30081, gen30083)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1number_2, gen30084)

	gen30105 := MakeNative(func(__e *ControlFlow) {
		V3866 := __e.Get(1)
		_ = V3866
		V3867 := __e.Get(2)
		_ = V3867
		V3868 := __e.Get(3)
		_ = V3868
		gen30102 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30099 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen30085 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen30085)

				gen30086 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen30087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30091 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30092 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30093 := Call(__e, gen30092, symnumber, Nil)

				gen30094 := Call(__e, gen30091, sym_1_1_6, gen30093)

				gen30095 := Call(__e, gen30090, B, gen30094)

				gen30096 := Call(__e, gen30089, gen30095, Nil)

				gen30097 := Call(__e, gen30088, sym_1_1_6, gen30096)

				gen30098 := Call(__e, gen30087, A, gen30097)

				__e.TailApply(gen30086, V3866, gen30098, V3867, V3868)

				return

			}, 1)

			gen30100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen30101 := Call(__e, gen30100, V3867)

			__e.TailApply(gen30099, gen30101)

			return

		}, 1)

		gen30103 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30104 := Call(__e, gen30103, V3867)

		__e.TailApply(gen30102, gen30104)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1occurrences, gen30105)

	gen30114 := MakeNative(func(__e *ControlFlow) {
		V3876 := __e.Get(1)
		_ = V3876
		V3877 := __e.Get(2)
		_ = V3877
		V3878 := __e.Get(3)
		_ = V3878
		gen30106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30106)

		gen30107 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30109 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30110 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30111 := Call(__e, gen30110, symboolean, Nil)

		gen30112 := Call(__e, gen30109, sym_1_1_6, gen30111)

		gen30113 := Call(__e, gen30108, symsymbol, gen30112)

		__e.TailApply(gen30107, V3876, gen30113, V3877, V3878)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1occurs_1check, gen30114)

	gen30123 := MakeNative(func(__e *ControlFlow) {
		V3886 := __e.Get(1)
		_ = V3886
		V3887 := __e.Get(2)
		_ = V3887
		V3888 := __e.Get(3)
		_ = V3888
		gen30115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30115)

		gen30116 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30117 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30118 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30120 := Call(__e, gen30119, symboolean, Nil)

		gen30121 := Call(__e, gen30118, sym_1_1_6, gen30120)

		gen30122 := Call(__e, gen30117, symsymbol, gen30121)

		__e.TailApply(gen30116, V3886, gen30122, V3887, V3888)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1optimise, gen30123)

	gen30138 := MakeNative(func(__e *ControlFlow) {
		V3896 := __e.Get(1)
		_ = V3896
		V3897 := __e.Get(2)
		_ = V3897
		V3898 := __e.Get(3)
		_ = V3898
		gen30124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30124)

		gen30125 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30126 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30127 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30132 := Call(__e, gen30131, symboolean, Nil)

		gen30133 := Call(__e, gen30130, sym_1_1_6, gen30132)

		gen30134 := Call(__e, gen30129, symboolean, gen30133)

		gen30135 := Call(__e, gen30128, gen30134, Nil)

		gen30136 := Call(__e, gen30127, sym_1_1_6, gen30135)

		gen30137 := Call(__e, gen30126, symboolean, gen30136)

		__e.TailApply(gen30125, V3896, gen30137, V3897, V3898)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1or, gen30138)

	gen30145 := MakeNative(func(__e *ControlFlow) {
		V3906 := __e.Get(1)
		_ = V3906
		V3907 := __e.Get(2)
		_ = V3907
		V3908 := __e.Get(3)
		_ = V3908
		gen30139 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30139)

		gen30140 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30141 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30142 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30143 := Call(__e, gen30142, symstring, Nil)

		gen30144 := Call(__e, gen30141, sym_1_1_6, gen30143)

		__e.TailApply(gen30140, V3906, gen30144, V3907, V3908)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1os, gen30145)

	gen30154 := MakeNative(func(__e *ControlFlow) {
		V3916 := __e.Get(1)
		_ = V3916
		V3917 := __e.Get(2)
		_ = V3917
		V3918 := __e.Get(3)
		_ = V3918
		gen30146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30146)

		gen30147 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30148 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30150 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30151 := Call(__e, gen30150, symboolean, Nil)

		gen30152 := Call(__e, gen30149, sym_1_1_6, gen30151)

		gen30153 := Call(__e, gen30148, symsymbol, gen30152)

		__e.TailApply(gen30147, V3916, gen30153, V3917, V3918)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1package_2, gen30154)

	gen30161 := MakeNative(func(__e *ControlFlow) {
		V3926 := __e.Get(1)
		_ = V3926
		V3927 := __e.Get(2)
		_ = V3927
		V3928 := __e.Get(3)
		_ = V3928
		gen30155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30155)

		gen30156 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30157 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30158 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30159 := Call(__e, gen30158, symstring, Nil)

		gen30160 := Call(__e, gen30157, sym_1_1_6, gen30159)

		__e.TailApply(gen30156, V3926, gen30160, V3927, V3928)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1port, gen30161)

	gen30168 := MakeNative(func(__e *ControlFlow) {
		V3936 := __e.Get(1)
		_ = V3936
		V3937 := __e.Get(2)
		_ = V3937
		V3938 := __e.Get(3)
		_ = V3938
		gen30162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30162)

		gen30163 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30166 := Call(__e, gen30165, symstring, Nil)

		gen30167 := Call(__e, gen30164, sym_1_1_6, gen30166)

		__e.TailApply(gen30163, V3936, gen30167, V3937, V3938)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1porters, gen30168)

	gen30183 := MakeNative(func(__e *ControlFlow) {
		V3946 := __e.Get(1)
		_ = V3946
		V3947 := __e.Get(2)
		_ = V3947
		V3948 := __e.Get(3)
		_ = V3948
		gen30169 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30169)

		gen30170 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30171 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30174 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30175 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30177 := Call(__e, gen30176, symstring, Nil)

		gen30178 := Call(__e, gen30175, sym_1_1_6, gen30177)

		gen30179 := Call(__e, gen30174, symnumber, gen30178)

		gen30180 := Call(__e, gen30173, gen30179, Nil)

		gen30181 := Call(__e, gen30172, sym_1_1_6, gen30180)

		gen30182 := Call(__e, gen30171, symstring, gen30181)

		__e.TailApply(gen30170, V3946, gen30182, V3947, V3948)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1pos, gen30183)

	gen30202 := MakeNative(func(__e *ControlFlow) {
		V3956 := __e.Get(1)
		_ = V3956
		V3957 := __e.Get(2)
		_ = V3957
		V3958 := __e.Get(3)
		_ = V3958
		gen30184 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30184)

		gen30185 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30186 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30187 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30188 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30189 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30190 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30191 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30192 := Call(__e, gen30191, symout, Nil)

		gen30193 := Call(__e, gen30190, symstream, gen30192)

		gen30194 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30195 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30196 := Call(__e, gen30195, symstring, Nil)

		gen30197 := Call(__e, gen30194, sym_1_1_6, gen30196)

		gen30198 := Call(__e, gen30189, gen30193, gen30197)

		gen30199 := Call(__e, gen30188, gen30198, Nil)

		gen30200 := Call(__e, gen30187, sym_1_1_6, gen30199)

		gen30201 := Call(__e, gen30186, symstring, gen30200)

		__e.TailApply(gen30185, V3956, gen30201, V3957, V3958)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1pr, gen30202)

	gen30214 := MakeNative(func(__e *ControlFlow) {
		V3966 := __e.Get(1)
		_ = V3966
		V3967 := __e.Get(2)
		_ = V3967
		V3968 := __e.Get(3)
		_ = V3968
		gen30211 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30203)

			gen30204 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30205 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30206 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30207 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30208 := Call(__e, gen30207, A, Nil)

			gen30209 := Call(__e, gen30206, sym_1_1_6, gen30208)

			gen30210 := Call(__e, gen30205, A, gen30209)

			__e.TailApply(gen30204, V3966, gen30210, V3967, V3968)

			return

		}, 1)

		gen30212 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30213 := Call(__e, gen30212, V3967)

		__e.TailApply(gen30211, gen30213)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1print, gen30214)

	gen30241 := MakeNative(func(__e *ControlFlow) {
		V3976 := __e.Get(1)
		_ = V3976
		V3977 := __e.Get(2)
		_ = V3977
		V3978 := __e.Get(3)
		_ = V3978
		gen30238 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30235 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen30215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen30215)

				gen30216 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen30217 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30218 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30219 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30220 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30221 := Call(__e, gen30220, B, Nil)

				gen30222 := Call(__e, gen30219, sym_1_1_6, gen30221)

				gen30223 := Call(__e, gen30218, A, gen30222)

				gen30224 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30225 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30226 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30227 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30228 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30229 := Call(__e, gen30228, B, Nil)

				gen30230 := Call(__e, gen30227, sym_1_1_6, gen30229)

				gen30231 := Call(__e, gen30226, A, gen30230)

				gen30232 := Call(__e, gen30225, gen30231, Nil)

				gen30233 := Call(__e, gen30224, sym_1_1_6, gen30232)

				gen30234 := Call(__e, gen30217, gen30223, gen30233)

				__e.TailApply(gen30216, V3976, gen30234, V3977, V3978)

				return

			}, 1)

			gen30236 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen30237 := Call(__e, gen30236, V3977)

			__e.TailApply(gen30235, gen30237)

			return

		}, 1)

		gen30239 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30240 := Call(__e, gen30239, V3977)

		__e.TailApply(gen30238, gen30240)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1profile, gen30241)

	gen30258 := MakeNative(func(__e *ControlFlow) {
		V3986 := __e.Get(1)
		_ = V3986
		V3987 := __e.Get(2)
		_ = V3987
		V3988 := __e.Get(3)
		_ = V3988
		gen30242 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30242)

		gen30243 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30244 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30245 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30246 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30247 := Call(__e, gen30246, symsymbol, Nil)

		gen30248 := Call(__e, gen30245, symlist, gen30247)

		gen30249 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30250 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30251 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30252 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30253 := Call(__e, gen30252, symsymbol, Nil)

		gen30254 := Call(__e, gen30251, symlist, gen30253)

		gen30255 := Call(__e, gen30250, gen30254, Nil)

		gen30256 := Call(__e, gen30249, sym_1_1_6, gen30255)

		gen30257 := Call(__e, gen30244, gen30248, gen30256)

		__e.TailApply(gen30243, V3986, gen30257, V3987, V3988)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1preclude, gen30258)

	gen30267 := MakeNative(func(__e *ControlFlow) {
		V3996 := __e.Get(1)
		_ = V3996
		V3997 := __e.Get(2)
		_ = V3997
		V3998 := __e.Get(3)
		_ = V3998
		gen30259 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30259)

		gen30260 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30261 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30264 := Call(__e, gen30263, symstring, Nil)

		gen30265 := Call(__e, gen30262, sym_1_1_6, gen30264)

		gen30266 := Call(__e, gen30261, symstring, gen30265)

		__e.TailApply(gen30260, V3996, gen30266, V3997, V3998)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4proc_1nl, gen30267)

	gen30300 := MakeNative(func(__e *ControlFlow) {
		V4006 := __e.Get(1)
		_ = V4006
		V4007 := __e.Get(2)
		_ = V4007
		V4008 := __e.Get(3)
		_ = V4008
		gen30297 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30294 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen30268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen30268)

				gen30269 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen30270 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30271 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30272 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30273 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30274 := Call(__e, gen30273, B, Nil)

				gen30275 := Call(__e, gen30272, sym_1_1_6, gen30274)

				gen30276 := Call(__e, gen30271, A, gen30275)

				gen30277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30279 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30280 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30281 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30282 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30283 := Call(__e, gen30282, B, Nil)

				gen30284 := Call(__e, gen30281, sym_1_1_6, gen30283)

				gen30285 := Call(__e, gen30280, A, gen30284)

				gen30286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30287 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30288 := Call(__e, gen30287, symnumber, Nil)

				gen30289 := Call(__e, gen30286, sym_d, gen30288)

				gen30290 := Call(__e, gen30279, gen30285, gen30289)

				gen30291 := Call(__e, gen30278, gen30290, Nil)

				gen30292 := Call(__e, gen30277, sym_1_1_6, gen30291)

				gen30293 := Call(__e, gen30270, gen30276, gen30292)

				__e.TailApply(gen30269, V4006, gen30293, V4007, V4008)

				return

			}, 1)

			gen30295 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen30296 := Call(__e, gen30295, V4007)

			__e.TailApply(gen30294, gen30296)

			return

		}, 1)

		gen30298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30299 := Call(__e, gen30298, V4007)

		__e.TailApply(gen30297, gen30299)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1profile_1results, gen30300)

	gen30309 := MakeNative(func(__e *ControlFlow) {
		V4016 := __e.Get(1)
		_ = V4016
		V4017 := __e.Get(2)
		_ = V4017
		V4018 := __e.Get(3)
		_ = V4018
		gen30301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30301)

		gen30302 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30305 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30306 := Call(__e, gen30305, symsymbol, Nil)

		gen30307 := Call(__e, gen30304, sym_1_1_6, gen30306)

		gen30308 := Call(__e, gen30303, symsymbol, gen30307)

		__e.TailApply(gen30302, V4016, gen30308, V4017, V4018)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1protect, gen30309)

	gen30326 := MakeNative(func(__e *ControlFlow) {
		V4026 := __e.Get(1)
		_ = V4026
		V4027 := __e.Get(2)
		_ = V4027
		V4028 := __e.Get(3)
		_ = V4028
		gen30310 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30310)

		gen30311 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30312 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30314 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30315 := Call(__e, gen30314, symsymbol, Nil)

		gen30316 := Call(__e, gen30313, symlist, gen30315)

		gen30317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30321 := Call(__e, gen30320, symsymbol, Nil)

		gen30322 := Call(__e, gen30319, symlist, gen30321)

		gen30323 := Call(__e, gen30318, gen30322, Nil)

		gen30324 := Call(__e, gen30317, sym_1_1_6, gen30323)

		gen30325 := Call(__e, gen30312, gen30316, gen30324)

		__e.TailApply(gen30311, V4026, gen30325, V4027, V4028)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1preclude_1all_1but, gen30326)

	gen30345 := MakeNative(func(__e *ControlFlow) {
		V4036 := __e.Get(1)
		_ = V4036
		V4037 := __e.Get(2)
		_ = V4037
		V4038 := __e.Get(3)
		_ = V4038
		gen30327 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30327)

		gen30328 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30329 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30330 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30331 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30333 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30335 := Call(__e, gen30334, symout, Nil)

		gen30336 := Call(__e, gen30333, symstream, gen30335)

		gen30337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30339 := Call(__e, gen30338, symstring, Nil)

		gen30340 := Call(__e, gen30337, sym_1_1_6, gen30339)

		gen30341 := Call(__e, gen30332, gen30336, gen30340)

		gen30342 := Call(__e, gen30331, gen30341, Nil)

		gen30343 := Call(__e, gen30330, sym_1_1_6, gen30342)

		gen30344 := Call(__e, gen30329, symstring, gen30343)

		__e.TailApply(gen30328, V4036, gen30344, V4037, V4038)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4prhush, gen30345)

	gen30358 := MakeNative(func(__e *ControlFlow) {
		V4046 := __e.Get(1)
		_ = V4046
		V4047 := __e.Get(2)
		_ = V4047
		V4048 := __e.Get(3)
		_ = V4048
		gen30346 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30346)

		gen30347 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30349 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30350 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30351 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30353 := Call(__e, gen30352, symunit, Nil)

		gen30354 := Call(__e, gen30351, symlist, gen30353)

		gen30355 := Call(__e, gen30350, gen30354, Nil)

		gen30356 := Call(__e, gen30349, sym_1_1_6, gen30355)

		gen30357 := Call(__e, gen30348, symsymbol, gen30356)

		__e.TailApply(gen30347, V4046, gen30357, V4047, V4048)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1ps, gen30358)

	gen30371 := MakeNative(func(__e *ControlFlow) {
		V4056 := __e.Get(1)
		_ = V4056
		V4057 := __e.Get(2)
		_ = V4057
		V4058 := __e.Get(3)
		_ = V4058
		gen30359 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30359)

		gen30360 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30361 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30362 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30363 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30364 := Call(__e, gen30363, symin, Nil)

		gen30365 := Call(__e, gen30362, symstream, gen30364)

		gen30366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30368 := Call(__e, gen30367, symunit, Nil)

		gen30369 := Call(__e, gen30366, sym_1_1_6, gen30368)

		gen30370 := Call(__e, gen30361, gen30365, gen30369)

		__e.TailApply(gen30360, V4056, gen30370, V4057, V4058)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read, gen30371)

	gen30384 := MakeNative(func(__e *ControlFlow) {
		V4066 := __e.Get(1)
		_ = V4066
		V4067 := __e.Get(2)
		_ = V4067
		V4068 := __e.Get(3)
		_ = V4068
		gen30372 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30372)

		gen30373 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30376 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30377 := Call(__e, gen30376, symin, Nil)

		gen30378 := Call(__e, gen30375, symstream, gen30377)

		gen30379 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30380 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30381 := Call(__e, gen30380, symnumber, Nil)

		gen30382 := Call(__e, gen30379, sym_1_1_6, gen30381)

		gen30383 := Call(__e, gen30374, gen30378, gen30382)

		__e.TailApply(gen30373, V4066, gen30383, V4067, V4068)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1byte, gen30384)

	gen30397 := MakeNative(func(__e *ControlFlow) {
		V4076 := __e.Get(1)
		_ = V4076
		V4077 := __e.Get(2)
		_ = V4077
		V4078 := __e.Get(3)
		_ = V4078
		gen30385 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30385)

		gen30386 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30387 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30388 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30389 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30391 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30392 := Call(__e, gen30391, symnumber, Nil)

		gen30393 := Call(__e, gen30390, symlist, gen30392)

		gen30394 := Call(__e, gen30389, gen30393, Nil)

		gen30395 := Call(__e, gen30388, sym_1_1_6, gen30394)

		gen30396 := Call(__e, gen30387, symstring, gen30395)

		__e.TailApply(gen30386, V4076, gen30396, V4077, V4078)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file_1as_1bytelist, gen30397)

	gen30406 := MakeNative(func(__e *ControlFlow) {
		V4086 := __e.Get(1)
		_ = V4086
		V4087 := __e.Get(2)
		_ = V4087
		V4088 := __e.Get(3)
		_ = V4088
		gen30398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30398)

		gen30399 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30400 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30401 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30402 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30403 := Call(__e, gen30402, symstring, Nil)

		gen30404 := Call(__e, gen30401, sym_1_1_6, gen30403)

		gen30405 := Call(__e, gen30400, symstring, gen30404)

		__e.TailApply(gen30399, V4086, gen30405, V4087, V4088)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file_1as_1string, gen30406)

	gen30419 := MakeNative(func(__e *ControlFlow) {
		V4096 := __e.Get(1)
		_ = V4096
		V4097 := __e.Get(2)
		_ = V4097
		V4098 := __e.Get(3)
		_ = V4098
		gen30407 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30407)

		gen30408 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30409 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30410 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30413 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30414 := Call(__e, gen30413, symunit, Nil)

		gen30415 := Call(__e, gen30412, symlist, gen30414)

		gen30416 := Call(__e, gen30411, gen30415, Nil)

		gen30417 := Call(__e, gen30410, sym_1_1_6, gen30416)

		gen30418 := Call(__e, gen30409, symstring, gen30417)

		__e.TailApply(gen30408, V4096, gen30418, V4097, V4098)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file, gen30419)

	gen30432 := MakeNative(func(__e *ControlFlow) {
		V4106 := __e.Get(1)
		_ = V4106
		V4107 := __e.Get(2)
		_ = V4107
		V4108 := __e.Get(3)
		_ = V4108
		gen30420 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30420)

		gen30421 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30422 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30423 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30424 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30425 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30426 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30427 := Call(__e, gen30426, symunit, Nil)

		gen30428 := Call(__e, gen30425, symlist, gen30427)

		gen30429 := Call(__e, gen30424, gen30428, Nil)

		gen30430 := Call(__e, gen30423, sym_1_1_6, gen30429)

		gen30431 := Call(__e, gen30422, symstring, gen30430)

		__e.TailApply(gen30421, V4106, gen30431, V4107, V4108)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1from_1string, gen30432)

	gen30439 := MakeNative(func(__e *ControlFlow) {
		V4116 := __e.Get(1)
		_ = V4116
		V4117 := __e.Get(2)
		_ = V4117
		V4118 := __e.Get(3)
		_ = V4118
		gen30433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30433)

		gen30434 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30435 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30436 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30437 := Call(__e, gen30436, symstring, Nil)

		gen30438 := Call(__e, gen30435, sym_1_1_6, gen30437)

		__e.TailApply(gen30434, V4116, gen30438, V4117, V4118)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1release, gen30439)

	gen30465 := MakeNative(func(__e *ControlFlow) {
		V4126 := __e.Get(1)
		_ = V4126
		V4127 := __e.Get(2)
		_ = V4127
		V4128 := __e.Get(3)
		_ = V4128
		gen30462 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30440 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30440)

			gen30441 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30442 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30443 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30448 := Call(__e, gen30447, A, Nil)

			gen30449 := Call(__e, gen30446, symlist, gen30448)

			gen30450 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30451 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30452 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30453 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30454 := Call(__e, gen30453, A, Nil)

			gen30455 := Call(__e, gen30452, symlist, gen30454)

			gen30456 := Call(__e, gen30451, gen30455, Nil)

			gen30457 := Call(__e, gen30450, sym_1_1_6, gen30456)

			gen30458 := Call(__e, gen30445, gen30449, gen30457)

			gen30459 := Call(__e, gen30444, gen30458, Nil)

			gen30460 := Call(__e, gen30443, sym_1_1_6, gen30459)

			gen30461 := Call(__e, gen30442, A, gen30460)

			__e.TailApply(gen30441, V4126, gen30461, V4127, V4128)

			return

		}, 1)

		gen30463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30464 := Call(__e, gen30463, V4127)

		__e.TailApply(gen30462, gen30464)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1remove, gen30465)

	gen30485 := MakeNative(func(__e *ControlFlow) {
		V4136 := __e.Get(1)
		_ = V4136
		V4137 := __e.Get(2)
		_ = V4137
		V4138 := __e.Get(3)
		_ = V4138
		gen30482 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30466)

			gen30467 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30468 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30469 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30470 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30471 := Call(__e, gen30470, A, Nil)

			gen30472 := Call(__e, gen30469, symlist, gen30471)

			gen30473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30475 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30476 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30477 := Call(__e, gen30476, A, Nil)

			gen30478 := Call(__e, gen30475, symlist, gen30477)

			gen30479 := Call(__e, gen30474, gen30478, Nil)

			gen30480 := Call(__e, gen30473, sym_1_1_6, gen30479)

			gen30481 := Call(__e, gen30468, gen30472, gen30480)

			__e.TailApply(gen30467, V4136, gen30481, V4137, V4138)

			return

		}, 1)

		gen30483 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30484 := Call(__e, gen30483, V4137)

		__e.TailApply(gen30482, gen30484)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1reverse, gen30485)

	gen30497 := MakeNative(func(__e *ControlFlow) {
		V4146 := __e.Get(1)
		_ = V4146
		V4147 := __e.Get(2)
		_ = V4147
		V4148 := __e.Get(3)
		_ = V4148
		gen30494 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30486 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30486)

			gen30487 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30488 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30489 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30490 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30491 := Call(__e, gen30490, A, Nil)

			gen30492 := Call(__e, gen30489, sym_1_1_6, gen30491)

			gen30493 := Call(__e, gen30488, symstring, gen30492)

			__e.TailApply(gen30487, V4146, gen30493, V4147, V4148)

			return

		}, 1)

		gen30495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30496 := Call(__e, gen30495, V4147)

		__e.TailApply(gen30494, gen30496)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1simple_1error, gen30497)

	gen30518 := MakeNative(func(__e *ControlFlow) {
		V4156 := __e.Get(1)
		_ = V4156
		V4157 := __e.Get(2)
		_ = V4157
		V4158 := __e.Get(3)
		_ = V4158
		gen30515 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30512 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen30498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen30498)

				gen30499 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen30500 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30501 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30502 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30503 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30504 := Call(__e, gen30503, B, Nil)

				gen30505 := Call(__e, gen30502, sym_d, gen30504)

				gen30506 := Call(__e, gen30501, A, gen30505)

				gen30507 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30508 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30509 := Call(__e, gen30508, B, Nil)

				gen30510 := Call(__e, gen30507, sym_1_1_6, gen30509)

				gen30511 := Call(__e, gen30500, gen30506, gen30510)

				__e.TailApply(gen30499, V4156, gen30511, V4157, V4158)

				return

			}, 1)

			gen30513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen30514 := Call(__e, gen30513, V4157)

			__e.TailApply(gen30512, gen30514)

			return

		}, 1)

		gen30516 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30517 := Call(__e, gen30516, V4157)

		__e.TailApply(gen30515, gen30517)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1snd, gen30518)

	gen30527 := MakeNative(func(__e *ControlFlow) {
		V4166 := __e.Get(1)
		_ = V4166
		V4167 := __e.Get(2)
		_ = V4167
		V4168 := __e.Get(3)
		_ = V4168
		gen30519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30519)

		gen30520 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30521 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30522 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30523 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30524 := Call(__e, gen30523, symsymbol, Nil)

		gen30525 := Call(__e, gen30522, sym_1_1_6, gen30524)

		gen30526 := Call(__e, gen30521, symsymbol, gen30525)

		__e.TailApply(gen30520, V4166, gen30526, V4167, V4168)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1specialise, gen30527)

	gen30536 := MakeNative(func(__e *ControlFlow) {
		V4176 := __e.Get(1)
		_ = V4176
		V4177 := __e.Get(2)
		_ = V4177
		V4178 := __e.Get(3)
		_ = V4178
		gen30528 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30528)

		gen30529 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30530 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30531 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30532 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30533 := Call(__e, gen30532, symboolean, Nil)

		gen30534 := Call(__e, gen30531, sym_1_1_6, gen30533)

		gen30535 := Call(__e, gen30530, symsymbol, gen30534)

		__e.TailApply(gen30529, V4176, gen30535, V4177, V4178)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1spy, gen30536)

	gen30545 := MakeNative(func(__e *ControlFlow) {
		V4186 := __e.Get(1)
		_ = V4186
		V4187 := __e.Get(2)
		_ = V4187
		V4188 := __e.Get(3)
		_ = V4188
		gen30537 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30537)

		gen30538 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30539 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30541 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30542 := Call(__e, gen30541, symboolean, Nil)

		gen30543 := Call(__e, gen30540, sym_1_1_6, gen30542)

		gen30544 := Call(__e, gen30539, symsymbol, gen30543)

		__e.TailApply(gen30538, V4186, gen30544, V4187, V4188)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1step, gen30545)

	gen30556 := MakeNative(func(__e *ControlFlow) {
		V4196 := __e.Get(1)
		_ = V4196
		V4197 := __e.Get(2)
		_ = V4197
		V4198 := __e.Get(3)
		_ = V4198
		gen30546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30546)

		gen30547 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30548 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30549 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30550 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30551 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30552 := Call(__e, gen30551, symin, Nil)

		gen30553 := Call(__e, gen30550, symstream, gen30552)

		gen30554 := Call(__e, gen30549, gen30553, Nil)

		gen30555 := Call(__e, gen30548, sym_1_1_6, gen30554)

		__e.TailApply(gen30547, V4196, gen30555, V4197, V4198)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1stinput, gen30556)

	gen30567 := MakeNative(func(__e *ControlFlow) {
		V4206 := __e.Get(1)
		_ = V4206
		V4207 := __e.Get(2)
		_ = V4207
		V4208 := __e.Get(3)
		_ = V4208
		gen30557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30557)

		gen30558 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30559 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30561 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30562 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30563 := Call(__e, gen30562, symout, Nil)

		gen30564 := Call(__e, gen30561, symstream, gen30563)

		gen30565 := Call(__e, gen30560, gen30564, Nil)

		gen30566 := Call(__e, gen30559, sym_1_1_6, gen30565)

		__e.TailApply(gen30558, V4206, gen30566, V4207, V4208)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1sterror, gen30567)

	gen30578 := MakeNative(func(__e *ControlFlow) {
		V4216 := __e.Get(1)
		_ = V4216
		V4217 := __e.Get(2)
		_ = V4217
		V4218 := __e.Get(3)
		_ = V4218
		gen30568 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30568)

		gen30569 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30570 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30572 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30573 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30574 := Call(__e, gen30573, symout, Nil)

		gen30575 := Call(__e, gen30572, symstream, gen30574)

		gen30576 := Call(__e, gen30571, gen30575, Nil)

		gen30577 := Call(__e, gen30570, sym_1_1_6, gen30576)

		__e.TailApply(gen30569, V4216, gen30577, V4217, V4218)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1stoutput, gen30578)

	gen30590 := MakeNative(func(__e *ControlFlow) {
		V4226 := __e.Get(1)
		_ = V4226
		V4227 := __e.Get(2)
		_ = V4227
		V4228 := __e.Get(3)
		_ = V4228
		gen30587 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30579 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30579)

			gen30580 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30583 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30584 := Call(__e, gen30583, symboolean, Nil)

			gen30585 := Call(__e, gen30582, sym_1_1_6, gen30584)

			gen30586 := Call(__e, gen30581, A, gen30585)

			__e.TailApply(gen30580, V4226, gen30586, V4227, V4228)

			return

		}, 1)

		gen30588 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30589 := Call(__e, gen30588, V4227)

		__e.TailApply(gen30587, gen30589)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_2, gen30590)

	gen30602 := MakeNative(func(__e *ControlFlow) {
		V4236 := __e.Get(1)
		_ = V4236
		V4237 := __e.Get(2)
		_ = V4237
		V4238 := __e.Get(3)
		_ = V4238
		gen30599 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30591 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30591)

			gen30592 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30593 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30594 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30595 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30596 := Call(__e, gen30595, symstring, Nil)

			gen30597 := Call(__e, gen30594, sym_1_1_6, gen30596)

			gen30598 := Call(__e, gen30593, A, gen30597)

			__e.TailApply(gen30592, V4236, gen30598, V4237, V4238)

			return

		}, 1)

		gen30600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30601 := Call(__e, gen30600, V4237)

		__e.TailApply(gen30599, gen30601)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1str, gen30602)

	gen30611 := MakeNative(func(__e *ControlFlow) {
		V4246 := __e.Get(1)
		_ = V4246
		V4247 := __e.Get(2)
		_ = V4247
		V4248 := __e.Get(3)
		_ = V4248
		gen30603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30603)

		gen30604 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30605 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30606 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30607 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30608 := Call(__e, gen30607, symnumber, Nil)

		gen30609 := Call(__e, gen30606, sym_1_1_6, gen30608)

		gen30610 := Call(__e, gen30605, symstring, gen30609)

		__e.TailApply(gen30604, V4246, gen30610, V4247, V4248)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_1_6n, gen30611)

	gen30620 := MakeNative(func(__e *ControlFlow) {
		V4256 := __e.Get(1)
		_ = V4256
		V4257 := __e.Get(2)
		_ = V4257
		V4258 := __e.Get(3)
		_ = V4258
		gen30612 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30612)

		gen30613 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30614 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30615 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30616 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30617 := Call(__e, gen30616, symsymbol, Nil)

		gen30618 := Call(__e, gen30615, sym_1_1_6, gen30617)

		gen30619 := Call(__e, gen30614, symstring, gen30618)

		__e.TailApply(gen30613, V4256, gen30619, V4257, V4258)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_1_6symbol, gen30620)

	gen30633 := MakeNative(func(__e *ControlFlow) {
		V4266 := __e.Get(1)
		_ = V4266
		V4267 := __e.Get(2)
		_ = V4267
		V4268 := __e.Get(3)
		_ = V4268
		gen30621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30621)

		gen30622 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30625 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30626 := Call(__e, gen30625, symnumber, Nil)

		gen30627 := Call(__e, gen30624, symlist, gen30626)

		gen30628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30629 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30630 := Call(__e, gen30629, symnumber, Nil)

		gen30631 := Call(__e, gen30628, sym_1_1_6, gen30630)

		gen30632 := Call(__e, gen30623, gen30627, gen30631)

		__e.TailApply(gen30622, V4266, gen30632, V4267, V4268)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1sum, gen30633)

	gen30645 := MakeNative(func(__e *ControlFlow) {
		V4276 := __e.Get(1)
		_ = V4276
		V4277 := __e.Get(2)
		_ = V4277
		V4278 := __e.Get(3)
		_ = V4278
		gen30642 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30634 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30634)

			gen30635 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30636 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30637 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30638 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30639 := Call(__e, gen30638, symboolean, Nil)

			gen30640 := Call(__e, gen30637, sym_1_1_6, gen30639)

			gen30641 := Call(__e, gen30636, A, gen30640)

			__e.TailApply(gen30635, V4276, gen30641, V4277, V4278)

			return

		}, 1)

		gen30643 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30644 := Call(__e, gen30643, V4277)

		__e.TailApply(gen30642, gen30644)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1symbol_2, gen30645)

	gen30654 := MakeNative(func(__e *ControlFlow) {
		V4286 := __e.Get(1)
		_ = V4286
		V4287 := __e.Get(2)
		_ = V4287
		V4288 := __e.Get(3)
		_ = V4288
		gen30646 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30646)

		gen30647 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30648 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30649 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30650 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30651 := Call(__e, gen30650, symsymbol, Nil)

		gen30652 := Call(__e, gen30649, sym_1_1_6, gen30651)

		gen30653 := Call(__e, gen30648, symsymbol, gen30652)

		__e.TailApply(gen30647, V4286, gen30653, V4287, V4288)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1systemf, gen30654)

	gen30674 := MakeNative(func(__e *ControlFlow) {
		V4296 := __e.Get(1)
		_ = V4296
		V4297 := __e.Get(2)
		_ = V4297
		V4298 := __e.Get(3)
		_ = V4298
		gen30671 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30655 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30655)

			gen30656 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30657 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30658 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30659 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30660 := Call(__e, gen30659, A, Nil)

			gen30661 := Call(__e, gen30658, symlist, gen30660)

			gen30662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30663 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30664 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30666 := Call(__e, gen30665, A, Nil)

			gen30667 := Call(__e, gen30664, symlist, gen30666)

			gen30668 := Call(__e, gen30663, gen30667, Nil)

			gen30669 := Call(__e, gen30662, sym_1_1_6, gen30668)

			gen30670 := Call(__e, gen30657, gen30661, gen30669)

			__e.TailApply(gen30656, V4296, gen30670, V4297, V4298)

			return

		}, 1)

		gen30672 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30673 := Call(__e, gen30672, V4297)

		__e.TailApply(gen30671, gen30673)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tail, gen30674)

	gen30683 := MakeNative(func(__e *ControlFlow) {
		V4306 := __e.Get(1)
		_ = V4306
		V4307 := __e.Get(2)
		_ = V4307
		V4308 := __e.Get(3)
		_ = V4308
		gen30675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30675)

		gen30676 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30679 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30680 := Call(__e, gen30679, symstring, Nil)

		gen30681 := Call(__e, gen30678, sym_1_1_6, gen30680)

		gen30682 := Call(__e, gen30677, symstring, gen30681)

		__e.TailApply(gen30676, V4306, gen30682, V4307, V4308)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tlstr, gen30683)

	gen30703 := MakeNative(func(__e *ControlFlow) {
		V4316 := __e.Get(1)
		_ = V4316
		V4317 := __e.Get(2)
		_ = V4317
		V4318 := __e.Get(3)
		_ = V4318
		gen30700 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30684 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30684)

			gen30685 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30686 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30688 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30689 := Call(__e, gen30688, A, Nil)

			gen30690 := Call(__e, gen30687, symvector, gen30689)

			gen30691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30693 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30694 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30695 := Call(__e, gen30694, A, Nil)

			gen30696 := Call(__e, gen30693, symvector, gen30695)

			gen30697 := Call(__e, gen30692, gen30696, Nil)

			gen30698 := Call(__e, gen30691, sym_1_1_6, gen30697)

			gen30699 := Call(__e, gen30686, gen30690, gen30698)

			__e.TailApply(gen30685, V4316, gen30699, V4317, V4318)

			return

		}, 1)

		gen30701 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30702 := Call(__e, gen30701, V4317)

		__e.TailApply(gen30700, gen30702)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tlv, gen30703)

	gen30712 := MakeNative(func(__e *ControlFlow) {
		V4326 := __e.Get(1)
		_ = V4326
		V4327 := __e.Get(2)
		_ = V4327
		V4328 := __e.Get(3)
		_ = V4328
		gen30704 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30704)

		gen30705 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30707 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30709 := Call(__e, gen30708, symboolean, Nil)

		gen30710 := Call(__e, gen30707, sym_1_1_6, gen30709)

		gen30711 := Call(__e, gen30706, symsymbol, gen30710)

		__e.TailApply(gen30705, V4326, gen30711, V4327, V4328)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tc, gen30712)

	gen30719 := MakeNative(func(__e *ControlFlow) {
		V4336 := __e.Get(1)
		_ = V4336
		V4337 := __e.Get(2)
		_ = V4337
		V4338 := __e.Get(3)
		_ = V4338
		gen30713 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30713)

		gen30714 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30715 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30716 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30717 := Call(__e, gen30716, symboolean, Nil)

		gen30718 := Call(__e, gen30715, sym_1_1_6, gen30717)

		__e.TailApply(gen30714, V4336, gen30718, V4337, V4338)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tc_2, gen30719)

	gen30735 := MakeNative(func(__e *ControlFlow) {
		V4346 := __e.Get(1)
		_ = V4346
		V4347 := __e.Get(2)
		_ = V4347
		V4348 := __e.Get(3)
		_ = V4348
		gen30732 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30720)

			gen30721 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30722 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30723 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30724 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30725 := Call(__e, gen30724, A, Nil)

			gen30726 := Call(__e, gen30723, symlazy, gen30725)

			gen30727 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30729 := Call(__e, gen30728, A, Nil)

			gen30730 := Call(__e, gen30727, sym_1_1_6, gen30729)

			gen30731 := Call(__e, gen30722, gen30726, gen30730)

			__e.TailApply(gen30721, V4346, gen30731, V4347, V4348)

			return

		}, 1)

		gen30733 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30734 := Call(__e, gen30733, V4347)

		__e.TailApply(gen30732, gen30734)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1thaw, gen30735)

	gen30744 := MakeNative(func(__e *ControlFlow) {
		V4356 := __e.Get(1)
		_ = V4356
		V4357 := __e.Get(2)
		_ = V4357
		V4358 := __e.Get(3)
		_ = V4358
		gen30736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30736)

		gen30737 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30739 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30741 := Call(__e, gen30740, symsymbol, Nil)

		gen30742 := Call(__e, gen30739, sym_1_1_6, gen30741)

		gen30743 := Call(__e, gen30738, symsymbol, gen30742)

		__e.TailApply(gen30737, V4356, gen30743, V4357, V4358)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1track, gen30744)

	gen30768 := MakeNative(func(__e *ControlFlow) {
		V4366 := __e.Get(1)
		_ = V4366
		V4367 := __e.Get(2)
		_ = V4367
		V4368 := __e.Get(3)
		_ = V4368
		gen30765 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30745)

			gen30746 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30747 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30748 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30749 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30750 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30751 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30752 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30753 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30754 := Call(__e, gen30753, A, Nil)

			gen30755 := Call(__e, gen30752, sym_1_1_6, gen30754)

			gen30756 := Call(__e, gen30751, symexception, gen30755)

			gen30757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30758 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30759 := Call(__e, gen30758, A, Nil)

			gen30760 := Call(__e, gen30757, sym_1_1_6, gen30759)

			gen30761 := Call(__e, gen30750, gen30756, gen30760)

			gen30762 := Call(__e, gen30749, gen30761, Nil)

			gen30763 := Call(__e, gen30748, sym_1_1_6, gen30762)

			gen30764 := Call(__e, gen30747, A, gen30763)

			__e.TailApply(gen30746, V4366, gen30764, V4367, V4368)

			return

		}, 1)

		gen30766 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30767 := Call(__e, gen30766, V4367)

		__e.TailApply(gen30765, gen30767)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1trap_1error, gen30768)

	gen30780 := MakeNative(func(__e *ControlFlow) {
		V4376 := __e.Get(1)
		_ = V4376
		V4377 := __e.Get(2)
		_ = V4377
		V4378 := __e.Get(3)
		_ = V4378
		gen30777 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30769)

			gen30770 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30772 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30773 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30774 := Call(__e, gen30773, symboolean, Nil)

			gen30775 := Call(__e, gen30772, sym_1_1_6, gen30774)

			gen30776 := Call(__e, gen30771, A, gen30775)

			__e.TailApply(gen30770, V4376, gen30776, V4377, V4378)

			return

		}, 1)

		gen30778 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30779 := Call(__e, gen30778, V4377)

		__e.TailApply(gen30777, gen30779)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tuple_2, gen30780)

	gen30789 := MakeNative(func(__e *ControlFlow) {
		V4386 := __e.Get(1)
		_ = V4386
		V4387 := __e.Get(2)
		_ = V4387
		V4388 := __e.Get(3)
		_ = V4388
		gen30781 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30781)

		gen30782 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30783 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30784 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30785 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30786 := Call(__e, gen30785, symsymbol, Nil)

		gen30787 := Call(__e, gen30784, sym_1_1_6, gen30786)

		gen30788 := Call(__e, gen30783, symsymbol, gen30787)

		__e.TailApply(gen30782, V4386, gen30788, V4387, V4388)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1undefmacro, gen30789)

	gen30819 := MakeNative(func(__e *ControlFlow) {
		V4396 := __e.Get(1)
		_ = V4396
		V4397 := __e.Get(2)
		_ = V4397
		V4398 := __e.Get(3)
		_ = V4398
		gen30816 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30790)

			gen30791 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30792 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30794 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30795 := Call(__e, gen30794, A, Nil)

			gen30796 := Call(__e, gen30793, symlist, gen30795)

			gen30797 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30798 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30799 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30800 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30801 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30802 := Call(__e, gen30801, A, Nil)

			gen30803 := Call(__e, gen30800, symlist, gen30802)

			gen30804 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30807 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30808 := Call(__e, gen30807, A, Nil)

			gen30809 := Call(__e, gen30806, symlist, gen30808)

			gen30810 := Call(__e, gen30805, gen30809, Nil)

			gen30811 := Call(__e, gen30804, sym_1_1_6, gen30810)

			gen30812 := Call(__e, gen30799, gen30803, gen30811)

			gen30813 := Call(__e, gen30798, gen30812, Nil)

			gen30814 := Call(__e, gen30797, sym_1_1_6, gen30813)

			gen30815 := Call(__e, gen30792, gen30796, gen30814)

			__e.TailApply(gen30791, V4396, gen30815, V4397, V4398)

			return

		}, 1)

		gen30817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30818 := Call(__e, gen30817, V4397)

		__e.TailApply(gen30816, gen30818)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1union, gen30819)

	gen30846 := MakeNative(func(__e *ControlFlow) {
		V4406 := __e.Get(1)
		_ = V4406
		V4407 := __e.Get(2)
		_ = V4407
		V4408 := __e.Get(3)
		_ = V4408
		gen30843 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30840 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen30820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen30820)

				gen30821 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen30822 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30823 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30824 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30825 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30826 := Call(__e, gen30825, B, Nil)

				gen30827 := Call(__e, gen30824, sym_1_1_6, gen30826)

				gen30828 := Call(__e, gen30823, A, gen30827)

				gen30829 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30830 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30831 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30832 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30833 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen30834 := Call(__e, gen30833, B, Nil)

				gen30835 := Call(__e, gen30832, sym_1_1_6, gen30834)

				gen30836 := Call(__e, gen30831, A, gen30835)

				gen30837 := Call(__e, gen30830, gen30836, Nil)

				gen30838 := Call(__e, gen30829, sym_1_1_6, gen30837)

				gen30839 := Call(__e, gen30822, gen30828, gen30838)

				__e.TailApply(gen30821, V4406, gen30839, V4407, V4408)

				return

			}, 1)

			gen30841 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen30842 := Call(__e, gen30841, V4407)

			__e.TailApply(gen30840, gen30842)

			return

		}, 1)

		gen30844 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30845 := Call(__e, gen30844, V4407)

		__e.TailApply(gen30843, gen30845)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1unprofile, gen30846)

	gen30855 := MakeNative(func(__e *ControlFlow) {
		V4416 := __e.Get(1)
		_ = V4416
		V4417 := __e.Get(2)
		_ = V4417
		V4418 := __e.Get(3)
		_ = V4418
		gen30847 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30847)

		gen30848 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30849 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30850 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30851 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30852 := Call(__e, gen30851, symsymbol, Nil)

		gen30853 := Call(__e, gen30850, sym_1_1_6, gen30852)

		gen30854 := Call(__e, gen30849, symsymbol, gen30853)

		__e.TailApply(gen30848, V4416, gen30854, V4417, V4418)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1untrack, gen30855)

	gen30864 := MakeNative(func(__e *ControlFlow) {
		V4426 := __e.Get(1)
		_ = V4426
		V4427 := __e.Get(2)
		_ = V4427
		V4428 := __e.Get(3)
		_ = V4428
		gen30856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30856)

		gen30857 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30858 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30859 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30860 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30861 := Call(__e, gen30860, symsymbol, Nil)

		gen30862 := Call(__e, gen30859, sym_1_1_6, gen30861)

		gen30863 := Call(__e, gen30858, symsymbol, gen30862)

		__e.TailApply(gen30857, V4426, gen30863, V4427, V4428)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1unspecialise, gen30864)

	gen30876 := MakeNative(func(__e *ControlFlow) {
		V4436 := __e.Get(1)
		_ = V4436
		V4437 := __e.Get(2)
		_ = V4437
		V4438 := __e.Get(3)
		_ = V4438
		gen30873 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30865 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30865)

			gen30866 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30867 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30868 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30869 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30870 := Call(__e, gen30869, symboolean, Nil)

			gen30871 := Call(__e, gen30868, sym_1_1_6, gen30870)

			gen30872 := Call(__e, gen30867, A, gen30871)

			__e.TailApply(gen30866, V4436, gen30872, V4437, V4438)

			return

		}, 1)

		gen30874 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30875 := Call(__e, gen30874, V4437)

		__e.TailApply(gen30873, gen30875)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1variable_2, gen30876)

	gen30888 := MakeNative(func(__e *ControlFlow) {
		V4446 := __e.Get(1)
		_ = V4446
		V4447 := __e.Get(2)
		_ = V4447
		V4448 := __e.Get(3)
		_ = V4448
		gen30885 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30877 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30877)

			gen30878 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30879 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30880 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30881 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30882 := Call(__e, gen30881, symboolean, Nil)

			gen30883 := Call(__e, gen30880, sym_1_1_6, gen30882)

			gen30884 := Call(__e, gen30879, A, gen30883)

			__e.TailApply(gen30878, V4446, gen30884, V4447, V4448)

			return

		}, 1)

		gen30886 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30887 := Call(__e, gen30886, V4447)

		__e.TailApply(gen30885, gen30887)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector_2, gen30888)

	gen30895 := MakeNative(func(__e *ControlFlow) {
		V4456 := __e.Get(1)
		_ = V4456
		V4457 := __e.Get(2)
		_ = V4457
		V4458 := __e.Get(3)
		_ = V4458
		gen30889 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30889)

		gen30890 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30891 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30892 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30893 := Call(__e, gen30892, symstring, Nil)

		gen30894 := Call(__e, gen30891, sym_1_1_6, gen30893)

		__e.TailApply(gen30890, V4456, gen30894, V4457, V4458)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1version, gen30895)

	gen30913 := MakeNative(func(__e *ControlFlow) {
		V4466 := __e.Get(1)
		_ = V4466
		V4467 := __e.Get(2)
		_ = V4467
		V4468 := __e.Get(3)
		_ = V4468
		gen30910 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen30896 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen30896)

			gen30897 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen30898 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30899 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30900 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30901 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30902 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen30904 := Call(__e, gen30903, A, Nil)

			gen30905 := Call(__e, gen30902, sym_1_1_6, gen30904)

			gen30906 := Call(__e, gen30901, A, gen30905)

			gen30907 := Call(__e, gen30900, gen30906, Nil)

			gen30908 := Call(__e, gen30899, sym_1_1_6, gen30907)

			gen30909 := Call(__e, gen30898, symstring, gen30908)

			__e.TailApply(gen30897, V4466, gen30909, V4467, V4468)

			return

		}, 1)

		gen30911 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen30912 := Call(__e, gen30911, V4467)

		__e.TailApply(gen30910, gen30912)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1write_1to_1file, gen30913)

	gen30932 := MakeNative(func(__e *ControlFlow) {
		V4476 := __e.Get(1)
		_ = V4476
		V4477 := __e.Get(2)
		_ = V4477
		V4478 := __e.Get(3)
		_ = V4478
		gen30914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30914)

		gen30915 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30916 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30917 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30918 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30919 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30920 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30921 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30922 := Call(__e, gen30921, symout, Nil)

		gen30923 := Call(__e, gen30920, symstream, gen30922)

		gen30924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30926 := Call(__e, gen30925, symnumber, Nil)

		gen30927 := Call(__e, gen30924, sym_1_1_6, gen30926)

		gen30928 := Call(__e, gen30919, gen30923, gen30927)

		gen30929 := Call(__e, gen30918, gen30928, Nil)

		gen30930 := Call(__e, gen30917, sym_1_1_6, gen30929)

		gen30931 := Call(__e, gen30916, symnumber, gen30930)

		__e.TailApply(gen30915, V4476, gen30931, V4477, V4478)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1write_1byte, gen30932)

	gen30941 := MakeNative(func(__e *ControlFlow) {
		V4486 := __e.Get(1)
		_ = V4486
		V4487 := __e.Get(2)
		_ = V4487
		V4488 := __e.Get(3)
		_ = V4488
		gen30933 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30933)

		gen30934 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30938 := Call(__e, gen30937, symboolean, Nil)

		gen30939 := Call(__e, gen30936, sym_1_1_6, gen30938)

		gen30940 := Call(__e, gen30935, symstring, gen30939)

		__e.TailApply(gen30934, V4486, gen30940, V4487, V4488)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1y_1or_1n_2, gen30941)

	gen30956 := MakeNative(func(__e *ControlFlow) {
		V4496 := __e.Get(1)
		_ = V4496
		V4497 := __e.Get(2)
		_ = V4497
		V4498 := __e.Get(3)
		_ = V4498
		gen30942 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30942)

		gen30943 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30944 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30945 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30946 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30947 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30948 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30949 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30950 := Call(__e, gen30949, symboolean, Nil)

		gen30951 := Call(__e, gen30948, sym_1_1_6, gen30950)

		gen30952 := Call(__e, gen30947, symnumber, gen30951)

		gen30953 := Call(__e, gen30946, gen30952, Nil)

		gen30954 := Call(__e, gen30945, sym_1_1_6, gen30953)

		gen30955 := Call(__e, gen30944, symnumber, gen30954)

		__e.TailApply(gen30943, V4496, gen30955, V4497, V4498)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_6, gen30956)

	gen30971 := MakeNative(func(__e *ControlFlow) {
		V4506 := __e.Get(1)
		_ = V4506
		V4507 := __e.Get(2)
		_ = V4507
		V4508 := __e.Get(3)
		_ = V4508
		gen30957 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30957)

		gen30958 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30959 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30960 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30961 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30962 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30963 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30964 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30965 := Call(__e, gen30964, symboolean, Nil)

		gen30966 := Call(__e, gen30963, sym_1_1_6, gen30965)

		gen30967 := Call(__e, gen30962, symnumber, gen30966)

		gen30968 := Call(__e, gen30961, gen30967, Nil)

		gen30969 := Call(__e, gen30960, sym_1_1_6, gen30968)

		gen30970 := Call(__e, gen30959, symnumber, gen30969)

		__e.TailApply(gen30958, V4506, gen30970, V4507, V4508)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5, gen30971)

	gen30986 := MakeNative(func(__e *ControlFlow) {
		V4516 := __e.Get(1)
		_ = V4516
		V4517 := __e.Get(2)
		_ = V4517
		V4518 := __e.Get(3)
		_ = V4518
		gen30972 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30972)

		gen30973 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30974 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30975 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30976 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30977 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30979 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30980 := Call(__e, gen30979, symboolean, Nil)

		gen30981 := Call(__e, gen30978, sym_1_1_6, gen30980)

		gen30982 := Call(__e, gen30977, symnumber, gen30981)

		gen30983 := Call(__e, gen30976, gen30982, Nil)

		gen30984 := Call(__e, gen30975, sym_1_1_6, gen30983)

		gen30985 := Call(__e, gen30974, symnumber, gen30984)

		__e.TailApply(gen30973, V4516, gen30985, V4517, V4518)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_6_a, gen30986)

	gen31001 := MakeNative(func(__e *ControlFlow) {
		V4526 := __e.Get(1)
		_ = V4526
		V4527 := __e.Get(2)
		_ = V4527
		V4528 := __e.Get(3)
		_ = V4528
		gen30987 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen30987)

		gen30988 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen30989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30991 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30992 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30993 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30994 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen30995 := Call(__e, gen30994, symboolean, Nil)

		gen30996 := Call(__e, gen30993, sym_1_1_6, gen30995)

		gen30997 := Call(__e, gen30992, symnumber, gen30996)

		gen30998 := Call(__e, gen30991, gen30997, Nil)

		gen30999 := Call(__e, gen30990, sym_1_1_6, gen30998)

		gen31000 := Call(__e, gen30989, symnumber, gen30999)

		__e.TailApply(gen30988, V4526, gen31000, V4527, V4528)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_a, gen31001)

	gen31019 := MakeNative(func(__e *ControlFlow) {
		V4536 := __e.Get(1)
		_ = V4536
		V4537 := __e.Get(2)
		_ = V4537
		V4538 := __e.Get(3)
		_ = V4538
		gen31016 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen31002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			Call(__e, gen31002)

			gen31003 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			gen31004 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen31005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen31006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen31007 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen31008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen31009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen31010 := Call(__e, gen31009, symboolean, Nil)

			gen31011 := Call(__e, gen31008, sym_1_1_6, gen31010)

			gen31012 := Call(__e, gen31007, A, gen31011)

			gen31013 := Call(__e, gen31006, gen31012, Nil)

			gen31014 := Call(__e, gen31005, sym_1_1_6, gen31013)

			gen31015 := Call(__e, gen31004, A, gen31014)

			__e.TailApply(gen31003, V4536, gen31015, V4537, V4538)

			return

		}, 1)

		gen31017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen31018 := Call(__e, gen31017, V4537)

		__e.TailApply(gen31016, gen31018)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_a, gen31019)

	gen31034 := MakeNative(func(__e *ControlFlow) {
		V4546 := __e.Get(1)
		_ = V4546
		V4547 := __e.Get(2)
		_ = V4547
		V4548 := __e.Get(3)
		_ = V4548
		gen31020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen31020)

		gen31021 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen31022 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31023 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31024 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31025 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31027 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31028 := Call(__e, gen31027, symnumber, Nil)

		gen31029 := Call(__e, gen31026, sym_1_1_6, gen31028)

		gen31030 := Call(__e, gen31025, symnumber, gen31029)

		gen31031 := Call(__e, gen31024, gen31030, Nil)

		gen31032 := Call(__e, gen31023, sym_1_1_6, gen31031)

		gen31033 := Call(__e, gen31022, symnumber, gen31032)

		__e.TailApply(gen31021, V4546, gen31033, V4547, V4548)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_7, gen31034)

	gen31049 := MakeNative(func(__e *ControlFlow) {
		V4556 := __e.Get(1)
		_ = V4556
		V4557 := __e.Get(2)
		_ = V4557
		V4558 := __e.Get(3)
		_ = V4558
		gen31035 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen31035)

		gen31036 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen31037 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31038 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31039 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31040 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31041 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31042 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31043 := Call(__e, gen31042, symnumber, Nil)

		gen31044 := Call(__e, gen31041, sym_1_1_6, gen31043)

		gen31045 := Call(__e, gen31040, symnumber, gen31044)

		gen31046 := Call(__e, gen31039, gen31045, Nil)

		gen31047 := Call(__e, gen31038, sym_1_1_6, gen31046)

		gen31048 := Call(__e, gen31037, symnumber, gen31047)

		__e.TailApply(gen31036, V4556, gen31048, V4557, V4558)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_c, gen31049)

	gen31064 := MakeNative(func(__e *ControlFlow) {
		V4566 := __e.Get(1)
		_ = V4566
		V4567 := __e.Get(2)
		_ = V4567
		V4568 := __e.Get(3)
		_ = V4568
		gen31050 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen31050)

		gen31051 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen31052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31054 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31055 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31057 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31058 := Call(__e, gen31057, symnumber, Nil)

		gen31059 := Call(__e, gen31056, sym_1_1_6, gen31058)

		gen31060 := Call(__e, gen31055, symnumber, gen31059)

		gen31061 := Call(__e, gen31054, gen31060, Nil)

		gen31062 := Call(__e, gen31053, sym_1_1_6, gen31061)

		gen31063 := Call(__e, gen31052, symnumber, gen31062)

		__e.TailApply(gen31051, V4566, gen31063, V4567, V4568)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_1, gen31064)

	gen31079 := MakeNative(func(__e *ControlFlow) {
		V4576 := __e.Get(1)
		_ = V4576
		V4577 := __e.Get(2)
		_ = V4577
		V4578 := __e.Get(3)
		_ = V4578
		gen31065 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		Call(__e, gen31065)

		gen31066 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

		gen31067 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31068 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31069 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31070 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31071 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31072 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen31073 := Call(__e, gen31072, symnumber, Nil)

		gen31074 := Call(__e, gen31071, sym_1_1_6, gen31073)

		gen31075 := Call(__e, gen31070, symnumber, gen31074)

		gen31076 := Call(__e, gen31069, gen31075, Nil)

		gen31077 := Call(__e, gen31068, sym_1_1_6, gen31076)

		gen31078 := Call(__e, gen31067, symnumber, gen31077)

		__e.TailApply(gen31066, V4576, gen31078, V4577, V4578)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_d, gen31079)

	gen31100 := MakeNative(func(__e *ControlFlow) {
		V4586 := __e.Get(1)
		_ = V4586
		V4587 := __e.Get(2)
		_ = V4587
		V4588 := __e.Get(3)
		_ = V4588
		gen31097 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			gen31094 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				gen31080 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				Call(__e, gen31080)

				gen31081 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				gen31082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen31083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen31084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen31085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen31086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen31087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen31088 := Call(__e, gen31087, symboolean, Nil)

				gen31089 := Call(__e, gen31086, sym_1_1_6, gen31088)

				gen31090 := Call(__e, gen31085, B, gen31089)

				gen31091 := Call(__e, gen31084, gen31090, Nil)

				gen31092 := Call(__e, gen31083, sym_1_1_6, gen31091)

				gen31093 := Call(__e, gen31082, A, gen31092)

				__e.TailApply(gen31081, V4586, gen31093, V4587, V4588)

				return

			}, 1)

			gen31095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			gen31096 := Call(__e, gen31095, V4587)

			__e.TailApply(gen31094, gen31096)

			return

		}, 1)

		gen31098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		gen31099 := Call(__e, gen31098, V4587)

		__e.TailApply(gen31097, gen31099)

		return

	}, 3)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_a_a, gen31100)

	return

}, 0)
