package main

import . "github.com/tiancaiamao/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen880 := MakeNative(func(__e *ControlFlow) {
		V2342 := __e.Get(1)
		_ = V2342
		gen878 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

		gen879 := Call(__e, gen878, V2342, MakeNumber(1))

		if True == gen879 {
			gen873 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen874 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen875 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen876 := Call(__e, gen875, V2342, MakeString(""), symshen_4s)

			gen877 := Call(__e, gen874, MakeString("invalid initial dict size: "), gen876)

			__e.TailApply(gen873, gen877)

			return

		} else {
			if True == True {
				gen868 := MakeNative(func(__e *ControlFlow) {
					D := __e.Get(1)
					_ = D
					gen865 := MakeNative(func(__e *ControlFlow) {
						Tag := __e.Get(1)
						_ = Tag
						gen862 := MakeNative(func(__e *ControlFlow) {
							Capacity := __e.Get(1)
							_ = Capacity
							gen859 := MakeNative(func(__e *ControlFlow) {
								Count := __e.Get(1)
								_ = Count
								gen854 := MakeNative(func(__e *ControlFlow) {
									Fill := __e.Get(1)
									_ = Fill
									__e.Return(D)
									return
								}, 1)

								gen855 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

								gen856 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

								gen857 := Call(__e, gen856, MakeNumber(2), V2342)

								gen858 := Call(__e, gen855, D, MakeNumber(3), gen857, Nil)

								__e.TailApply(gen854, gen858)

								return

							}, 1)

							gen860 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

							gen861 := Call(__e, gen860, D, MakeNumber(2), MakeNumber(0))

							__e.TailApply(gen859, gen861)

							return

						}, 1)

						gen863 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

						gen864 := Call(__e, gen863, D, MakeNumber(1), V2342)

						__e.TailApply(gen862, gen864)

						return

					}, 1)

					gen866 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

					gen867 := Call(__e, gen866, D, MakeNumber(0), symshen_4dictionary)

					__e.TailApply(gen865, gen867)

					return

				}, 1)

				gen869 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

				gen870 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen871 := Call(__e, gen870, MakeNumber(3), V2342)

				gen872 := Call(__e, gen869, gen871)

				__e.TailApply(gen868, gen872)

				return

			} else {
				gen853 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen853, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict, gen880)

	gen889 := MakeNative(func(__e *ControlFlow) {
		V2344 := __e.Get(1)
		_ = V2344
		gen887 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		gen888 := Call(__e, gen887, V2344)

		if True == gen888 {
			gen881 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen883 := MakeNative(func(__e *ControlFlow) {
				gen882 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(gen882, V2344, MakeNumber(0))

				return

			}, 0)

			gen884 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1dictionary)
				return
			}, 1)

			gen885 := Call(__e, PrimNS1Value(symtry_1catch), gen883, gen884)

			gen886 := Call(__e, gen881, gen885, symshen_4dictionary)

			if True == gen886 {
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

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_2, gen889)

	gen891 := MakeNative(func(__e *ControlFlow) {
		V2346 := __e.Get(1)
		_ = V2346
		gen890 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(gen890, V2346, MakeNumber(1))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1capacity, gen891)

	gen893 := MakeNative(func(__e *ControlFlow) {
		V2348 := __e.Get(1)
		_ = V2348
		gen892 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(gen892, V2348, MakeNumber(2))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1count, gen893)

	gen895 := MakeNative(func(__e *ControlFlow) {
		V2351 := __e.Get(1)
		_ = V2351
		V2352 := __e.Get(2)
		_ = V2352
		gen894 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		__e.TailApply(gen894, V2351, MakeNumber(2), V2352)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1count_1_6, gen895)

	gen899 := MakeNative(func(__e *ControlFlow) {
		V2355 := __e.Get(1)
		_ = V2355
		V2356 := __e.Get(2)
		_ = V2356
		gen896 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		gen897 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		gen898 := Call(__e, gen897, MakeNumber(3), V2356)

		__e.TailApply(gen896, V2355, gen898)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5_1dict_1bucket, gen899)

	gen903 := MakeNative(func(__e *ControlFlow) {
		V2360 := __e.Get(1)
		_ = V2360
		V2361 := __e.Get(2)
		_ = V2361
		V2362 := __e.Get(3)
		_ = V2362
		gen900 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		gen901 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		gen902 := Call(__e, gen901, MakeNumber(3), V2361)

		__e.TailApply(gen900, V2360, gen902, V2362)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1bucket_1_6, gen903)

	gen916 := MakeNative(func(__e *ControlFlow) {
		V2366 := __e.Get(1)
		_ = V2366
		V2367 := __e.Get(2)
		_ = V2367
		V2368 := __e.Get(3)
		_ = V2368
		gen909 := MakeNative(func(__e *ControlFlow) {
			Diff := __e.Get(1)
			_ = Diff
			gen904 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1count_1_6)

			gen905 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			gen906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1count)

			gen907 := Call(__e, gen906, V2366)

			gen908 := Call(__e, gen905, Diff, gen907)

			__e.TailApply(gen904, V2366, gen908)

			return

		}, 1)

		gen910 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

		gen911 := Call(__e, PrimNS1Value(symns2_1value), symlength)

		gen912 := Call(__e, gen911, V2368)

		gen913 := Call(__e, PrimNS1Value(symns2_1value), symlength)

		gen914 := Call(__e, gen913, V2367)

		gen915 := Call(__e, gen910, gen912, gen914)

		__e.TailApply(gen909, gen915)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1update_1count, gen916)

	gen934 := MakeNative(func(__e *ControlFlow) {
		V2372 := __e.Get(1)
		_ = V2372
		V2373 := __e.Get(2)
		_ = V2373
		V2374 := __e.Get(3)
		_ = V2374
		gen929 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			gen926 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				gen923 := MakeNative(func(__e *ControlFlow) {
					NewBucket := __e.Get(1)
					_ = NewBucket
					gen920 := MakeNative(func(__e *ControlFlow) {
						Change := __e.Get(1)
						_ = Change
						gen917 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							__e.Return(V2374)
							return
						}, 1)

						gen918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1update_1count)

						gen919 := Call(__e, gen918, V2372, Bucket, NewBucket)

						__e.TailApply(gen917, gen919)

						return

					}, 1)

					gen921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1bucket_1_6)

					gen922 := Call(__e, gen921, V2372, N, NewBucket)

					__e.TailApply(gen920, gen922)

					return

				}, 1)

				gen924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1set)

				gen925 := Call(__e, gen924, V2373, V2374, Bucket)

				__e.TailApply(gen923, gen925)

				return

			}, 1)

			gen927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			gen928 := Call(__e, gen927, V2372, N)

			__e.TailApply(gen926, gen928)

			return

		}, 1)

		gen930 := Call(__e, PrimNS1Value(symns2_1value), symhash)

		gen931 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		gen932 := Call(__e, gen931, V2372)

		gen933 := Call(__e, gen930, V2373, gen932)

		__e.TailApply(gen929, gen933)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1_6, gen934)

	gen954 := MakeNative(func(__e *ControlFlow) {
		V2377 := __e.Get(1)
		_ = V2377
		V2378 := __e.Get(2)
		_ = V2378
		gen949 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			gen946 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				gen943 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					gen941 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

					gen942 := Call(__e, gen941, Result)

					if True == gen942 {
						gen936 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						gen937 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						gen938 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						gen939 := Call(__e, gen938, V2378, MakeString(" not found in dict\n"), symshen_4a)

						gen940 := Call(__e, gen937, MakeString("value "), gen939)

						__e.TailApply(gen936, gen940)

						return

					} else {
						gen935 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						__e.TailApply(gen935, Result)

						return

					}

				}, 1)

				gen944 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

				gen945 := Call(__e, gen944, V2378, Bucket)

				__e.TailApply(gen943, gen945)

				return

			}, 1)

			gen947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			gen948 := Call(__e, gen947, V2377, N)

			__e.TailApply(gen946, gen948)

			return

		}, 1)

		gen950 := Call(__e, PrimNS1Value(symns2_1value), symhash)

		gen951 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		gen952 := Call(__e, gen951, V2377)

		gen953 := Call(__e, gen950, V2378, gen952)

		__e.TailApply(gen949, gen953)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5_1dict, gen954)

	gen972 := MakeNative(func(__e *ControlFlow) {
		V2381 := __e.Get(1)
		_ = V2381
		V2382 := __e.Get(2)
		_ = V2382
		gen967 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			gen964 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				gen961 := MakeNative(func(__e *ControlFlow) {
					NewBucket := __e.Get(1)
					_ = NewBucket
					gen958 := MakeNative(func(__e *ControlFlow) {
						Change := __e.Get(1)
						_ = Change
						gen955 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							__e.Return(V2382)
							return
						}, 1)

						gen956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1update_1count)

						gen957 := Call(__e, gen956, V2381, Bucket, NewBucket)

						__e.TailApply(gen955, gen957)

						return

					}, 1)

					gen959 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1bucket_1_6)

					gen960 := Call(__e, gen959, V2381, N, NewBucket)

					__e.TailApply(gen958, gen960)

					return

				}, 1)

				gen962 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1rm)

				gen963 := Call(__e, gen962, V2382, Bucket)

				__e.TailApply(gen961, gen963)

				return

			}, 1)

			gen965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			gen966 := Call(__e, gen965, V2381, N)

			__e.TailApply(gen964, gen966)

			return

		}, 1)

		gen968 := Call(__e, PrimNS1Value(symns2_1value), symhash)

		gen969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		gen970 := Call(__e, gen969, V2381)

		gen971 := Call(__e, gen968, V2382, gen970)

		__e.TailApply(gen967, gen971)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1rm, gen972)

	gen977 := MakeNative(func(__e *ControlFlow) {
		V2386 := __e.Get(1)
		_ = V2386
		V2387 := __e.Get(2)
		_ = V2387
		V2388 := __e.Get(3)
		_ = V2388
		gen974 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			gen973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold_1h)

			__e.TailApply(gen973, V2386, V2387, V2388, MakeNumber(0), Limit)

			return

		}, 1)

		gen975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		gen976 := Call(__e, gen975, V2387)

		__e.TailApply(gen974, gen976)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1fold, gen977)

	gen990 := MakeNative(func(__e *ControlFlow) {
		V2395 := __e.Get(1)
		_ = V2395
		V2396 := __e.Get(2)
		_ = V2396
		V2397 := __e.Get(3)
		_ = V2397
		V2398 := __e.Get(4)
		_ = V2398
		V2399 := __e.Get(5)
		_ = V2399
		gen988 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen989 := Call(__e, gen988, V2399, V2398)

		if True == gen989 {
			__e.Return(V2397)
			return
		} else {
			if True == True {
				gen985 := MakeNative(func(__e *ControlFlow) {
					B := __e.Get(1)
					_ = B
					gen982 := MakeNative(func(__e *ControlFlow) {
						Acc := __e.Get(1)
						_ = Acc
						gen979 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold_1h)

						gen980 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

						gen981 := Call(__e, gen980, MakeNumber(1), V2398)

						__e.TailApply(gen979, V2395, V2396, Acc, gen981, V2399)

						return

					}, 1)

					gen983 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bucket_1fold)

					gen984 := Call(__e, gen983, V2395, B, V2397)

					__e.TailApply(gen982, gen984)

					return

				}, 1)

				gen986 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

				gen987 := Call(__e, gen986, V2396, V2398)

				__e.TailApply(gen985, gen987)

				return

			} else {
				gen978 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen978, MakeString("no cond match"))

				return

			}
		}

	}, 5)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1fold_1h, gen990)

	gen1014 := MakeNative(func(__e *ControlFlow) {
		V2403 := __e.Get(1)
		_ = V2403
		V2404 := __e.Get(2)
		_ = V2404
		V2405 := __e.Get(3)
		_ = V2405
		gen1012 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen1013 := Call(__e, gen1012, Nil, V2404)

		if True == gen1013 {
			__e.Return(V2405)
			return
		} else {
			gen1009 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen1010 := Call(__e, gen1009, V2404)

			var gen1011 Obj

			if True == gen1010 {
				gen1005 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen1006 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen1007 := Call(__e, gen1006, V2404)

				gen1008 := Call(__e, gen1005, gen1007)

				if True == gen1008 {
					gen1011 = True
				} else {
					gen1011 = False
				}

			} else {
				gen1011 = False
			}

			if True == gen1011 {
				gen993 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen994 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen995 := Call(__e, gen994, V2404)

				gen996 := Call(__e, gen993, gen995)

				gen997 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen998 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen999 := Call(__e, gen998, V2404)

				gen1000 := Call(__e, gen997, gen999)

				gen1001 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bucket_1fold)

				gen1002 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen1003 := Call(__e, gen1002, V2404)

				gen1004 := Call(__e, gen1001, V2403, gen1003, V2405)

				__e.TailApply(V2403, gen996, gen1000, gen1004)

				return

			} else {
				if True == True {
					gen992 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen992, symshen_4bucket_1fold)

					return

				} else {
					gen991 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen991, MakeString("no cond match"))

					return

				}
			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4bucket_1fold, gen1014)

	gen1018 := MakeNative(func(__e *ControlFlow) {
		V2407 := __e.Get(1)
		_ = V2407
		gen1015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold)

		gen1017 := MakeNative(func(__e *ControlFlow) {
			K := __e.Get(1)
			_ = K
			__e.Return(MakeNative(func(__e *ControlFlow) {
				__ := __e.Get(1)
				_ = __
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					gen1016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					__e.TailApply(gen1016, K, Acc)

					return

				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(gen1015, gen1017, V2407, Nil)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1keys, gen1018)

	gen1022 := MakeNative(func(__e *ControlFlow) {
		V2409 := __e.Get(1)
		_ = V2409
		gen1019 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold)

		gen1021 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V := __e.Get(1)
				_ = V
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					gen1020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					__e.TailApply(gen1020, V, Acc)

					return

				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(gen1019, gen1021, V2409, Nil)

		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4dict_1values, gen1022)

	return

}, 0)
