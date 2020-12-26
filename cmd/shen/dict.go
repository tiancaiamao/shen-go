package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen2285 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2342 := __args[0]
			_ = V2342
			gen2284 := Call(__e, ShenFunc(sym_5), V2342, MakeNumber(1))

			if True == gen2284 {
				gen2282 := Call(__e, ShenFunc(symshen_4app), V2342, MakeString(""), MakeSymbol("shen.s"))

				gen2283 := Call(__e, ShenFunc(symcn), MakeString("invalid initial dict size: "), gen2282)

				__e.TailApply(ShenFunc(symsimple_1error), gen2283)

				return

			} else {
				gen2275 := Call(__e, ShenFunc(sym_7), MakeNumber(3), V2342)

				gen2276 := Call(__e, ShenFunc(symabsvector), gen2275)

				D := gen2276
				gen2277 := Call(__e, ShenFunc(symaddress_1_6), D, MakeNumber(0), MakeSymbol("shen.dictionary"))

				Tag := gen2277
				_ = Tag
				gen2278 := Call(__e, ShenFunc(symaddress_1_6), D, MakeNumber(1), V2342)

				Capacity := gen2278
				_ = Capacity
				gen2279 := Call(__e, ShenFunc(symaddress_1_6), D, MakeNumber(2), MakeNumber(0))

				Count := gen2279
				_ = Count
				gen2280 := Call(__e, ShenFunc(sym_7), MakeNumber(2), V2342)

				gen2281 := Call(__e, ShenFunc(symshen_4fillvector), D, MakeNumber(3), gen2280, Nil)

				Fill := gen2281
				_ = Fill
				__e.Return(D)
				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict"), gen2285)

		gen2291 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2344 := __args[0]
			_ = V2344
			gen2290 := Call(__e, ShenFunc(symabsvector_2), V2344)

			if True == gen2290 {
				gen2286 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeSymbol("shen.not-dictionary"))
					return
				}, 1)
				gen2287 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(sym_5_1address), V2344, MakeNumber(0))

					return
				}, 0)
				gen2288 := Try(__e, gen2287).Catch(gen2286)
				gen2289 := Call(__e, ShenFunc(sym_a), gen2288, MakeSymbol("shen.dictionary"))

				if True == gen2289 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict?"), gen2291)

		gen2292 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2346 := __args[0]
			_ = V2346
			__e.TailApply(ShenFunc(sym_5_1address), V2346, MakeNumber(1))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-capacity"), gen2292)

		gen2293 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2348 := __args[0]
			_ = V2348
			__e.TailApply(ShenFunc(sym_5_1address), V2348, MakeNumber(2))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-count"), gen2293)

		gen2294 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2351 := __args[0]
			_ = V2351
			V2352 := __args[1]
			_ = V2352
			__e.TailApply(ShenFunc(symaddress_1_6), V2351, MakeNumber(2), V2352)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-count->"), gen2294)

		gen2296 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2355 := __args[0]
			_ = V2355
			V2356 := __args[1]
			_ = V2356
			gen2295 := Call(__e, ShenFunc(sym_7), MakeNumber(3), V2356)

			__e.TailApply(ShenFunc(sym_5_1address), V2355, gen2295)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<-dict-bucket"), gen2296)

		gen2298 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2360 := __args[0]
			_ = V2360
			V2361 := __args[1]
			_ = V2361
			V2362 := __args[2]
			_ = V2362
			gen2297 := Call(__e, ShenFunc(sym_7), MakeNumber(3), V2361)

			__e.TailApply(ShenFunc(symaddress_1_6), V2360, gen2297, V2362)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-bucket->"), gen2298)

		gen2304 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2366 := __args[0]
			_ = V2366
			V2367 := __args[1]
			_ = V2367
			V2368 := __args[2]
			_ = V2368
			gen2299 := Call(__e, ShenFunc(symlength), V2368)

			gen2300 := Call(__e, ShenFunc(symlength), V2367)

			gen2301 := Call(__e, ShenFunc(sym_1), gen2299, gen2300)

			Diff := gen2301
			gen2302 := Call(__e, ShenFunc(symshen_4dict_1count), V2366)

			gen2303 := Call(__e, ShenFunc(sym_7), Diff, gen2302)

			__e.TailApply(ShenFunc(symshen_4dict_1count_1_6), V2366, gen2303)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-update-count"), gen2304)

		gen2311 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2372 := __args[0]
			_ = V2372
			V2373 := __args[1]
			_ = V2373
			V2374 := __args[2]
			_ = V2374
			gen2305 := Call(__e, ShenFunc(symshen_4dict_1capacity), V2372)

			gen2306 := Call(__e, ShenFunc(symhash), V2373, gen2305)

			N := gen2306
			gen2307 := Call(__e, ShenFunc(symshen_4_5_1dict_1bucket), V2372, N)

			Bucket := gen2307
			gen2308 := Call(__e, ShenFunc(symshen_4assoc_1set), V2373, V2374, Bucket)

			NewBucket := gen2308
			gen2309 := Call(__e, ShenFunc(symshen_4dict_1bucket_1_6), V2372, N, NewBucket)

			Change := gen2309
			_ = Change
			gen2310 := Call(__e, ShenFunc(symshen_4dict_1update_1count), V2372, Bucket, NewBucket)

			Count := gen2310
			_ = Count
			__e.Return(V2374)
			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict->"), gen2311)

		gen2319 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2377 := __args[0]
			_ = V2377
			V2378 := __args[1]
			_ = V2378
			gen2312 := Call(__e, ShenFunc(symshen_4dict_1capacity), V2377)

			gen2313 := Call(__e, ShenFunc(symhash), V2378, gen2312)

			N := gen2313
			gen2314 := Call(__e, ShenFunc(symshen_4_5_1dict_1bucket), V2377, N)

			Bucket := gen2314
			gen2315 := Call(__e, ShenFunc(symassoc), V2378, Bucket)

			Result := gen2315
			gen2318 := Call(__e, ShenFunc(symempty_2), Result)

			if True == gen2318 {
				gen2316 := Call(__e, ShenFunc(symshen_4app), V2378, MakeString(" not found in dict\n"), MakeSymbol("shen.a"))

				gen2317 := Call(__e, ShenFunc(symcn), MakeString("value "), gen2316)

				__e.TailApply(ShenFunc(symsimple_1error), gen2317)

				return

			} else {
				__e.TailApply(ShenFunc(symtl), Result)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<-dict"), gen2319)

		gen2326 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2381 := __args[0]
			_ = V2381
			V2382 := __args[1]
			_ = V2382
			gen2320 := Call(__e, ShenFunc(symshen_4dict_1capacity), V2381)

			gen2321 := Call(__e, ShenFunc(symhash), V2382, gen2320)

			N := gen2321
			gen2322 := Call(__e, ShenFunc(symshen_4_5_1dict_1bucket), V2381, N)

			Bucket := gen2322
			_ = Bucket
			gen2323 := Call(__e, ShenFunc(symshen_4assoc_1rm), V2382, Bucket)

			NewBucket := gen2323
			_ = NewBucket
			gen2324 := Call(__e, ShenFunc(symshen_4dict_1bucket_1_6), V2381, N, NewBucket)

			Change := gen2324
			_ = Change
			gen2325 := Call(__e, ShenFunc(symshen_4dict_1update_1count), V2381, Bucket, NewBucket)

			Count := gen2325
			_ = Count
			__e.Return(V2382)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-rm"), gen2326)

		gen2328 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2386 := __args[0]
			_ = V2386
			V2387 := __args[1]
			_ = V2387
			V2388 := __args[2]
			_ = V2388
			gen2327 := Call(__e, ShenFunc(symshen_4dict_1capacity), V2387)

			Limit := gen2327
			__e.TailApply(ShenFunc(symshen_4dict_1fold_1h), V2386, V2387, V2388, MakeNumber(0), Limit)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-fold"), gen2328)

		gen2333 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2395 := __args[0]
			_ = V2395
			V2396 := __args[1]
			_ = V2396
			V2397 := __args[2]
			_ = V2397
			V2398 := __args[3]
			_ = V2398
			V2399 := __args[4]
			_ = V2399
			gen2332 := Call(__e, ShenFunc(sym_a), V2399, V2398)

			if True == gen2332 {
				__e.Return(V2397)
				return
			} else {
				gen2329 := Call(__e, ShenFunc(symshen_4_5_1dict_1bucket), V2396, V2398)

				B := gen2329
				gen2330 := Call(__e, ShenFunc(symshen_4bucket_1fold), V2395, B, V2397)

				Acc := gen2330
				gen2331 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V2398)

				__e.TailApply(ShenFunc(symshen_4dict_1fold_1h), V2395, V2396, Acc, gen2331, V2399)

				return

			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-fold-h"), gen2333)

		gen2345 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2403 := __args[0]
			_ = V2403
			V2404 := __args[1]
			_ = V2404
			V2405 := __args[2]
			_ = V2405
			gen2344 := Call(__e, ShenFunc(sym_a), Nil, V2404)

			if True == gen2344 {
				__e.Return(V2405)
				return
			} else {
				gen2342 := Call(__e, ShenFunc(symcons_2), V2404)

				var gen2343 Obj
				if True == gen2342 {
					gen2340 := Call(__e, ShenFunc(symhd), V2404)

					gen2341 := Call(__e, ShenFunc(symcons_2), gen2340)

					if True == gen2341 {
						gen2343 = True
					} else {
						gen2343 = False
					}

				} else {
					gen2343 = False
				}
				if True == gen2343 {
					gen2334 := Call(__e, ShenFunc(symhd), V2404)

					gen2335 := Call(__e, ShenFunc(symhd), gen2334)

					gen2336 := Call(__e, ShenFunc(symhd), V2404)

					gen2337 := Call(__e, ShenFunc(symtl), gen2336)

					gen2338 := Call(__e, ShenFunc(symtl), V2404)

					gen2339 := Call(__e, ShenFunc(symshen_4bucket_1fold), V2403, gen2338, V2405)

					__e.TailApply(V2403, gen2335, gen2337, gen2339)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.bucket-fold"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.bucket-fold"), gen2345)

		gen2347 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2407 := __args[0]
			_ = V2407
			gen2346 := MakeNative(func(__e Evaluator, __args ...Obj) {
				K := __args[0]
				_ = K
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					__ := __args[0]
					_ = __
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						Acc := __args[0]
						_ = Acc
						__e.TailApply(ShenFunc(symcons), K, Acc)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			__e.TailApply(ShenFunc(symshen_4dict_1fold), gen2346, V2407, Nil)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dict-keys"), gen2347)

		gen2349 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2409 := __args[0]
			_ = V2409
			gen2348 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__ := __args[0]
				_ = __
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V := __args[0]
					_ = V
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						Acc := __args[0]
						_ = Acc
						__e.TailApply(ShenFunc(symcons), V, Acc)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			__e.TailApply(ShenFunc(symshen_4dict_1fold), gen2348, V2409, Nil)

			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.dict-values"), gen2349)

		return

	}, 0))
}
