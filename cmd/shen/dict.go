package main

import . "github.com/tiancaiamao/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp3056 := MakeNative(func(__e *ControlFlow) {
		V2342 := __e.Get(1)
		_ = V2342
		tmp3072 := PrimLessThan(V2342, MakeNumber(1))

		if True == tmp3072 {
			tmp3070 := Call(__e, PrimNS2Value(symshen_4app), V2342, MakeString(""), symshen_4s)

			tmp3071 := PrimStringConcat(MakeString("invalid initial dict size: "), tmp3070)

			__e.Return(PrimSimpleError(tmp3071))
			return

		} else {
			tmp3058 := MakeNative(func(__e *ControlFlow) {
				D := __e.Get(1)
				_ = D
				tmp3059 := MakeNative(func(__e *ControlFlow) {
					Tag := __e.Get(1)
					_ = Tag
					tmp3060 := MakeNative(func(__e *ControlFlow) {
						Capacity := __e.Get(1)
						_ = Capacity
						tmp3061 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							tmp3062 := MakeNative(func(__e *ControlFlow) {
								Fill := __e.Get(1)
								_ = Fill
								__e.Return(D)
								return
							}, 1)

							tmp3063 := PrimNumberAdd(MakeNumber(2), V2342)

							tmp3064 := Call(__e, PrimNS2Value(symshen_4fillvector), D, MakeNumber(3), tmp3063, Nil)

							__e.TailApply(tmp3062, tmp3064)
							return

						}, 1)

						tmp3065 := PrimVectorSet(D, MakeNumber(2), MakeNumber(0))

						__e.TailApply(tmp3061, tmp3065)
						return

					}, 1)

					tmp3066 := PrimVectorSet(D, MakeNumber(1), V2342)

					__e.TailApply(tmp3060, tmp3066)
					return

				}, 1)

				tmp3067 := PrimVectorSet(D, MakeNumber(0), symshen_4dictionary)

				__e.TailApply(tmp3059, tmp3067)
				return

			}, 1)

			tmp3068 := PrimNumberAdd(MakeNumber(3), V2342)

			tmp3069 := PrimAbsvector(tmp3068)

			__e.TailApply(tmp3058, tmp3069)
			return

		}

	}, 1)

	tmp3073 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict, tmp3056)

	_ = tmp3073

	tmp3074 := MakeNative(func(__e *ControlFlow) {
		V2344 := __e.Get(1)
		_ = V2344
		tmp3081 := PrimIsVector(V2344)

		if True == tmp3081 {
			tmp3077 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V2344, MakeNumber(0)))
				return
			}, 0)

			tmp3078 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1dictionary)
				return
			}, 1)

			tmp3079 := Call(__e, PrimNS1Value(symtry_1catch), tmp3077, tmp3078)

			tmp3080 := PrimEqual(tmp3079, symshen_4dictionary)

			if True == tmp3080 {
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

	tmp3082 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_2, tmp3074)

	_ = tmp3082

	tmp3083 := MakeNative(func(__e *ControlFlow) {
		V2346 := __e.Get(1)
		_ = V2346
		__e.Return(PrimVectorGet(V2346, MakeNumber(1)))
		return
	}, 1)

	tmp3084 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1capacity, tmp3083)

	_ = tmp3084

	tmp3085 := MakeNative(func(__e *ControlFlow) {
		V2348 := __e.Get(1)
		_ = V2348
		__e.Return(PrimVectorGet(V2348, MakeNumber(2)))
		return
	}, 1)

	tmp3086 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1count, tmp3085)

	_ = tmp3086

	tmp3087 := MakeNative(func(__e *ControlFlow) {
		V2351 := __e.Get(1)
		_ = V2351
		V2352 := __e.Get(2)
		_ = V2352
		__e.Return(PrimVectorSet(V2351, MakeNumber(2), V2352))
		return
	}, 2)

	tmp3088 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1count_1_6, tmp3087)

	_ = tmp3088

	tmp3089 := MakeNative(func(__e *ControlFlow) {
		V2355 := __e.Get(1)
		_ = V2355
		V2356 := __e.Get(2)
		_ = V2356
		tmp3090 := PrimNumberAdd(MakeNumber(3), V2356)

		__e.Return(PrimVectorGet(V2355, tmp3090))
		return

	}, 2)

	tmp3091 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5_1dict_1bucket, tmp3089)

	_ = tmp3091

	tmp3092 := MakeNative(func(__e *ControlFlow) {
		V2360 := __e.Get(1)
		_ = V2360
		V2361 := __e.Get(2)
		_ = V2361
		V2362 := __e.Get(3)
		_ = V2362
		tmp3093 := PrimNumberAdd(MakeNumber(3), V2361)

		__e.Return(PrimVectorSet(V2360, tmp3093, V2362))
		return

	}, 3)

	tmp3094 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1bucket_1_6, tmp3092)

	_ = tmp3094

	tmp3095 := MakeNative(func(__e *ControlFlow) {
		V2366 := __e.Get(1)
		_ = V2366
		V2367 := __e.Get(2)
		_ = V2367
		V2368 := __e.Get(3)
		_ = V2368
		tmp3096 := MakeNative(func(__e *ControlFlow) {
			Diff := __e.Get(1)
			_ = Diff
			tmp3097 := Call(__e, PrimNS2Value(symshen_4dict_1count), V2366)

			tmp3098 := PrimNumberAdd(Diff, tmp3097)

			__e.TailApply(PrimNS2Value(symshen_4dict_1count_1_6), V2366, tmp3098)
			return

		}, 1)

		tmp3099 := Call(__e, PrimNS2Value(symlength), V2368)

		tmp3100 := Call(__e, PrimNS2Value(symlength), V2367)

		tmp3101 := PrimNumberSubtract(tmp3099, tmp3100)

		__e.TailApply(tmp3096, tmp3101)
		return

	}, 3)

	tmp3102 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1update_1count, tmp3095)

	_ = tmp3102

	tmp3103 := MakeNative(func(__e *ControlFlow) {
		V2372 := __e.Get(1)
		_ = V2372
		V2373 := __e.Get(2)
		_ = V2373
		V2374 := __e.Get(3)
		_ = V2374
		tmp3104 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp3105 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				tmp3106 := MakeNative(func(__e *ControlFlow) {
					NewBucket := __e.Get(1)
					_ = NewBucket
					tmp3107 := MakeNative(func(__e *ControlFlow) {
						Change := __e.Get(1)
						_ = Change
						tmp3108 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							__e.Return(V2374)
							return
						}, 1)

						tmp3109 := Call(__e, PrimNS2Value(symshen_4dict_1update_1count), V2372, Bucket, NewBucket)

						__e.TailApply(tmp3108, tmp3109)
						return

					}, 1)

					tmp3110 := Call(__e, PrimNS2Value(symshen_4dict_1bucket_1_6), V2372, N, NewBucket)

					__e.TailApply(tmp3107, tmp3110)
					return

				}, 1)

				tmp3111 := Call(__e, PrimNS2Value(symshen_4assoc_1set), V2373, V2374, Bucket)

				__e.TailApply(tmp3106, tmp3111)
				return

			}, 1)

			tmp3112 := Call(__e, PrimNS2Value(symshen_4_5_1dict_1bucket), V2372, N)

			__e.TailApply(tmp3105, tmp3112)
			return

		}, 1)

		tmp3113 := Call(__e, PrimNS2Value(symshen_4dict_1capacity), V2372)

		tmp3114 := Call(__e, PrimNS2Value(symhash), V2373, tmp3113)

		__e.TailApply(tmp3104, tmp3114)
		return

	}, 3)

	tmp3115 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1_6, tmp3103)

	_ = tmp3115

	tmp3116 := MakeNative(func(__e *ControlFlow) {
		V2377 := __e.Get(1)
		_ = V2377
		V2378 := __e.Get(2)
		_ = V2378
		tmp3117 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp3118 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				tmp3119 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp3123 := Call(__e, PrimNS2Value(symempty_2), Result)

					if True == tmp3123 {
						tmp3121 := Call(__e, PrimNS2Value(symshen_4app), V2378, MakeString(" not found in dict\n"), symshen_4a)

						tmp3122 := PrimStringConcat(MakeString("value "), tmp3121)

						__e.Return(PrimSimpleError(tmp3122))
						return

					} else {
						__e.Return(PrimTail(Result))
						return
					}

				}, 1)

				tmp3124 := Call(__e, PrimNS2Value(symassoc), V2378, Bucket)

				__e.TailApply(tmp3119, tmp3124)
				return

			}, 1)

			tmp3125 := Call(__e, PrimNS2Value(symshen_4_5_1dict_1bucket), V2377, N)

			__e.TailApply(tmp3118, tmp3125)
			return

		}, 1)

		tmp3126 := Call(__e, PrimNS2Value(symshen_4dict_1capacity), V2377)

		tmp3127 := Call(__e, PrimNS2Value(symhash), V2378, tmp3126)

		__e.TailApply(tmp3117, tmp3127)
		return

	}, 2)

	tmp3128 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5_1dict, tmp3116)

	_ = tmp3128

	tmp3129 := MakeNative(func(__e *ControlFlow) {
		V2381 := __e.Get(1)
		_ = V2381
		V2382 := __e.Get(2)
		_ = V2382
		tmp3130 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp3131 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				tmp3132 := MakeNative(func(__e *ControlFlow) {
					NewBucket := __e.Get(1)
					_ = NewBucket
					tmp3133 := MakeNative(func(__e *ControlFlow) {
						Change := __e.Get(1)
						_ = Change
						tmp3134 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							__e.Return(V2382)
							return
						}, 1)

						tmp3135 := Call(__e, PrimNS2Value(symshen_4dict_1update_1count), V2381, Bucket, NewBucket)

						__e.TailApply(tmp3134, tmp3135)
						return

					}, 1)

					tmp3136 := Call(__e, PrimNS2Value(symshen_4dict_1bucket_1_6), V2381, N, NewBucket)

					__e.TailApply(tmp3133, tmp3136)
					return

				}, 1)

				tmp3137 := Call(__e, PrimNS2Value(symshen_4assoc_1rm), V2382, Bucket)

				__e.TailApply(tmp3132, tmp3137)
				return

			}, 1)

			tmp3138 := Call(__e, PrimNS2Value(symshen_4_5_1dict_1bucket), V2381, N)

			__e.TailApply(tmp3131, tmp3138)
			return

		}, 1)

		tmp3139 := Call(__e, PrimNS2Value(symshen_4dict_1capacity), V2381)

		tmp3140 := Call(__e, PrimNS2Value(symhash), V2382, tmp3139)

		__e.TailApply(tmp3130, tmp3140)
		return

	}, 2)

	tmp3141 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1rm, tmp3129)

	_ = tmp3141

	tmp3142 := MakeNative(func(__e *ControlFlow) {
		V2386 := __e.Get(1)
		_ = V2386
		V2387 := __e.Get(2)
		_ = V2387
		V2388 := __e.Get(3)
		_ = V2388
		tmp3143 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			__e.TailApply(PrimNS2Value(symshen_4dict_1fold_1h), V2386, V2387, V2388, MakeNumber(0), Limit)
			return
		}, 1)

		tmp3144 := Call(__e, PrimNS2Value(symshen_4dict_1capacity), V2387)

		__e.TailApply(tmp3143, tmp3144)
		return

	}, 3)

	tmp3145 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1fold, tmp3142)

	_ = tmp3145

	tmp3146 := MakeNative(func(__e *ControlFlow) {
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
		tmp3153 := PrimEqual(V2399, V2398)

		if True == tmp3153 {
			__e.Return(V2397)
			return
		} else {
			tmp3148 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp3149 := MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					tmp3150 := PrimNumberAdd(MakeNumber(1), V2398)

					__e.TailApply(PrimNS2Value(symshen_4dict_1fold_1h), V2395, V2396, Acc, tmp3150, V2399)
					return

				}, 1)

				tmp3151 := Call(__e, PrimNS2Value(symshen_4bucket_1fold), V2395, B, V2397)

				__e.TailApply(tmp3149, tmp3151)
				return

			}, 1)

			tmp3152 := Call(__e, PrimNS2Value(symshen_4_5_1dict_1bucket), V2396, V2398)

			__e.TailApply(tmp3148, tmp3152)
			return

		}

	}, 5)

	tmp3154 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1fold_1h, tmp3146)

	_ = tmp3154

	tmp3155 := MakeNative(func(__e *ControlFlow) {
		V2403 := __e.Get(1)
		_ = V2403
		V2404 := __e.Get(2)
		_ = V2404
		V2405 := __e.Get(3)
		_ = V2405
		tmp3169 := PrimEqual(Nil, V2404)

		if True == tmp3169 {
			__e.Return(V2405)
			return
		} else {
			tmp3168 := PrimIsPair(V2404)

			var ifres3164 Obj

			if True == tmp3168 {
				tmp3166 := PrimHead(V2404)

				tmp3167 := PrimIsPair(tmp3166)

				var ifres3165 Obj

				if True == tmp3167 {
					ifres3165 = True

				} else {
					ifres3165 = False

				}

				ifres3164 = ifres3165

			} else {
				ifres3164 = False

			}

			if True == ifres3164 {
				tmp3158 := PrimHead(V2404)

				tmp3159 := PrimHead(tmp3158)

				tmp3160 := PrimHead(V2404)

				tmp3161 := PrimTail(tmp3160)

				tmp3162 := PrimTail(V2404)

				tmp3163 := Call(__e, PrimNS2Value(symshen_4bucket_1fold), V2403, tmp3162, V2405)

				__e.TailApply(V2403, tmp3159, tmp3161, tmp3163)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4bucket_1fold)
				return
			}

		}

	}, 3)

	tmp3170 := Call(__e, PrimNS1Value(symns2_1set), symshen_4bucket_1fold, tmp3155)

	_ = tmp3170

	tmp3171 := MakeNative(func(__e *ControlFlow) {
		V2407 := __e.Get(1)
		_ = V2407
		tmp3172 := MakeNative(func(__e *ControlFlow) {
			K := __e.Get(1)
			_ = K
			__e.Return(MakeNative(func(__e *ControlFlow) {
				__ := __e.Get(1)
				_ = __
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					__e.Return(PrimCons(K, Acc))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symshen_4dict_1fold), tmp3172, V2407, Nil)
		return

	}, 1)

	tmp3173 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1keys, tmp3171)

	_ = tmp3173

	tmp3174 := MakeNative(func(__e *ControlFlow) {
		V2409 := __e.Get(1)
		_ = V2409
		tmp3175 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V := __e.Get(1)
				_ = V
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					__e.Return(PrimCons(V, Acc))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symshen_4dict_1fold), tmp3175, V2409, Nil)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4dict_1values, tmp3174)
	return

}, 0)
