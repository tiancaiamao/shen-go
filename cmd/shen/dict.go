package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen380 := MakeNative(func(__e Evaluator) {
			V2342 := __e.Get(1)
			_ = V2342
			gen379 := Call(__e, __e.Global(sym_5), V2342, MakeNumber(1))

			if True == gen379 {
				gen377 := Call(__e, __e.Global(symshen_4app), V2342, MakeString(""), MakeSymbol("shen.s"))

				gen378 := Call(__e, __e.Global(symcn), MakeString("invalid initial dict size: "), gen377)

				__e.TailApply(__e.Global(symsimple_1error), gen378)

				return

			} else {
				gen370 := Call(__e, __e.Global(sym_7), MakeNumber(3), V2342)

				gen371 := Call(__e, __e.Global(symabsvector), gen370)

				D := gen371
				gen372 := Call(__e, __e.Global(symaddress_1_6), D, MakeNumber(0), MakeSymbol("shen.dictionary"))

				Tag := gen372
				_ = Tag
				gen373 := Call(__e, __e.Global(symaddress_1_6), D, MakeNumber(1), V2342)

				Capacity := gen373
				_ = Capacity
				gen374 := Call(__e, __e.Global(symaddress_1_6), D, MakeNumber(2), MakeNumber(0))

				Count := gen374
				_ = Count
				gen375 := Call(__e, __e.Global(sym_7), MakeNumber(2), V2342)

				gen376 := Call(__e, __e.Global(symshen_4fillvector), D, MakeNumber(3), gen375, Nil)

				Fill := gen376
				_ = Fill
				__e.Return(D)
				return

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict"), gen380)

		gen386 := MakeNative(func(__e Evaluator) {
			V2344 := __e.Get(1)
			_ = V2344
			gen385 := Call(__e, __e.Global(symabsvector_2), V2344)

			if True == gen385 {
				gen381 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeSymbol("shen.not-dictionary"))
					return
				}, 1)
				gen382 := MakeNative(func(__e Evaluator) {
					__e.TailApply(__e.Global(sym_5_1address), V2344, MakeNumber(0))

					return
				}, 0)
				gen383 := Try(__e, gen382).Catch(gen381)
				gen384 := Call(__e, __e.Global(sym_a), gen383, MakeSymbol("shen.dictionary"))

				if True == gen384 {
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
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict?"), gen386)

		gen387 := MakeNative(func(__e Evaluator) {
			V2346 := __e.Get(1)
			_ = V2346
			__e.TailApply(__e.Global(sym_5_1address), V2346, MakeNumber(1))

			return
		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-capacity"), gen387)

		gen388 := MakeNative(func(__e Evaluator) {
			V2348 := __e.Get(1)
			_ = V2348
			__e.TailApply(__e.Global(sym_5_1address), V2348, MakeNumber(2))

			return
		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-count"), gen388)

		gen389 := MakeNative(func(__e Evaluator) {
			V2351 := __e.Get(1)
			_ = V2351
			V2352 := __e.Get(2)
			_ = V2352
			__e.TailApply(__e.Global(symaddress_1_6), V2351, MakeNumber(2), V2352)

			return
		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-count->"), gen389)

		gen391 := MakeNative(func(__e Evaluator) {
			V2355 := __e.Get(1)
			_ = V2355
			V2356 := __e.Get(2)
			_ = V2356
			gen390 := Call(__e, __e.Global(sym_7), MakeNumber(3), V2356)

			__e.TailApply(__e.Global(sym_5_1address), V2355, gen390)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<-dict-bucket"), gen391)

		gen393 := MakeNative(func(__e Evaluator) {
			V2360 := __e.Get(1)
			_ = V2360
			V2361 := __e.Get(2)
			_ = V2361
			V2362 := __e.Get(3)
			_ = V2362
			gen392 := Call(__e, __e.Global(sym_7), MakeNumber(3), V2361)

			__e.TailApply(__e.Global(symaddress_1_6), V2360, gen392, V2362)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-bucket->"), gen393)

		gen399 := MakeNative(func(__e Evaluator) {
			V2366 := __e.Get(1)
			_ = V2366
			V2367 := __e.Get(2)
			_ = V2367
			V2368 := __e.Get(3)
			_ = V2368
			gen394 := Call(__e, __e.Global(symlength), V2368)

			gen395 := Call(__e, __e.Global(symlength), V2367)

			gen396 := Call(__e, __e.Global(sym_1), gen394, gen395)

			Diff := gen396
			gen397 := Call(__e, __e.Global(symshen_4dict_1count), V2366)

			gen398 := Call(__e, __e.Global(sym_7), Diff, gen397)

			__e.TailApply(__e.Global(symshen_4dict_1count_1_6), V2366, gen398)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-update-count"), gen399)

		gen406 := MakeNative(func(__e Evaluator) {
			V2372 := __e.Get(1)
			_ = V2372
			V2373 := __e.Get(2)
			_ = V2373
			V2374 := __e.Get(3)
			_ = V2374
			gen400 := Call(__e, __e.Global(symshen_4dict_1capacity), V2372)

			gen401 := Call(__e, __e.Global(symhash), V2373, gen400)

			N := gen401
			gen402 := Call(__e, __e.Global(symshen_4_5_1dict_1bucket), V2372, N)

			Bucket := gen402
			gen403 := Call(__e, __e.Global(symshen_4assoc_1set), V2373, V2374, Bucket)

			NewBucket := gen403
			gen404 := Call(__e, __e.Global(symshen_4dict_1bucket_1_6), V2372, N, NewBucket)

			Change := gen404
			_ = Change
			gen405 := Call(__e, __e.Global(symshen_4dict_1update_1count), V2372, Bucket, NewBucket)

			Count := gen405
			_ = Count
			__e.Return(V2374)
			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict->"), gen406)

		gen414 := MakeNative(func(__e Evaluator) {
			V2377 := __e.Get(1)
			_ = V2377
			V2378 := __e.Get(2)
			_ = V2378
			gen407 := Call(__e, __e.Global(symshen_4dict_1capacity), V2377)

			gen408 := Call(__e, __e.Global(symhash), V2378, gen407)

			N := gen408
			gen409 := Call(__e, __e.Global(symshen_4_5_1dict_1bucket), V2377, N)

			Bucket := gen409
			gen410 := Call(__e, __e.Global(symassoc), V2378, Bucket)

			Result := gen410
			gen413 := Call(__e, __e.Global(symempty_2), Result)

			if True == gen413 {
				gen411 := Call(__e, __e.Global(symshen_4app), V2378, MakeString(" not found in dict\n"), MakeSymbol("shen.a"))

				gen412 := Call(__e, __e.Global(symcn), MakeString("value "), gen411)

				__e.TailApply(__e.Global(symsimple_1error), gen412)

				return

			} else {
				__e.TailApply(__e.Global(symtl), Result)

				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<-dict"), gen414)

		gen421 := MakeNative(func(__e Evaluator) {
			V2381 := __e.Get(1)
			_ = V2381
			V2382 := __e.Get(2)
			_ = V2382
			gen415 := Call(__e, __e.Global(symshen_4dict_1capacity), V2381)

			gen416 := Call(__e, __e.Global(symhash), V2382, gen415)

			N := gen416
			gen417 := Call(__e, __e.Global(symshen_4_5_1dict_1bucket), V2381, N)

			Bucket := gen417
			gen418 := Call(__e, __e.Global(symshen_4assoc_1rm), V2382, Bucket)

			NewBucket := gen418
			gen419 := Call(__e, __e.Global(symshen_4dict_1bucket_1_6), V2381, N, NewBucket)

			Change := gen419
			_ = Change
			gen420 := Call(__e, __e.Global(symshen_4dict_1update_1count), V2381, Bucket, NewBucket)

			Count := gen420
			_ = Count
			__e.Return(V2382)
			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-rm"), gen421)

		gen423 := MakeNative(func(__e Evaluator) {
			V2386 := __e.Get(1)
			_ = V2386
			V2387 := __e.Get(2)
			_ = V2387
			V2388 := __e.Get(3)
			_ = V2388
			gen422 := Call(__e, __e.Global(symshen_4dict_1capacity), V2387)

			Limit := gen422
			__e.TailApply(__e.Global(symshen_4dict_1fold_1h), V2386, V2387, V2388, MakeNumber(0), Limit)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-fold"), gen423)

		gen428 := MakeNative(func(__e Evaluator) {
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
			gen427 := Call(__e, __e.Global(sym_a), V2399, V2398)

			if True == gen427 {
				__e.Return(V2397)
				return
			} else {
				gen424 := Call(__e, __e.Global(symshen_4_5_1dict_1bucket), V2396, V2398)

				B := gen424
				gen425 := Call(__e, __e.Global(symshen_4bucket_1fold), V2395, B, V2397)

				Acc := gen425
				gen426 := Call(__e, __e.Global(sym_7), MakeNumber(1), V2398)

				__e.TailApply(__e.Global(symshen_4dict_1fold_1h), V2395, V2396, Acc, gen426, V2399)

				return

			}

		}, 5)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-fold-h"), gen428)

		gen440 := MakeNative(func(__e Evaluator) {
			V2403 := __e.Get(1)
			_ = V2403
			V2404 := __e.Get(2)
			_ = V2404
			V2405 := __e.Get(3)
			_ = V2405
			gen439 := Call(__e, __e.Global(sym_a), Nil, V2404)

			if True == gen439 {
				__e.Return(V2405)
				return
			} else {
				gen437 := Call(__e, __e.Global(symcons_2), V2404)

				var gen438 Obj
				if True == gen437 {
					gen435 := Call(__e, __e.Global(symhd), V2404)

					gen436 := Call(__e, __e.Global(symcons_2), gen435)

					if True == gen436 {
						gen438 = True
					} else {
						gen438 = False
					}

				} else {
					gen438 = False
				}
				if True == gen438 {
					gen429 := Call(__e, __e.Global(symhd), V2404)

					gen430 := Call(__e, __e.Global(symhd), gen429)

					gen431 := Call(__e, __e.Global(symhd), V2404)

					gen432 := Call(__e, __e.Global(symtl), gen431)

					gen433 := Call(__e, __e.Global(symtl), V2404)

					gen434 := Call(__e, __e.Global(symshen_4bucket_1fold), V2403, gen433, V2405)

					__e.TailApply(V2403, gen430, gen432, gen434)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.bucket-fold"))

					return
				}

			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.bucket-fold"), gen440)

		gen442 := MakeNative(func(__e Evaluator) {
			V2407 := __e.Get(1)
			_ = V2407
			gen441 := MakeNative(func(__e Evaluator) {
				K := __e.Get(1)
				_ = K
				__e.Return(MakeNative(func(__e Evaluator) {
					__ := __e.Get(1)
					_ = __
					__e.Return(MakeNative(func(__e Evaluator) {
						Acc := __e.Get(1)
						_ = Acc
						__e.TailApply(__e.Global(symcons), K, Acc)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			__e.TailApply(__e.Global(symshen_4dict_1fold), gen441, V2407, Nil)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dict-keys"), gen442)

		gen444 := MakeNative(func(__e Evaluator) {
			V2409 := __e.Get(1)
			_ = V2409
			gen443 := MakeNative(func(__e Evaluator) {
				__ := __e.Get(1)
				_ = __
				__e.Return(MakeNative(func(__e Evaluator) {
					V := __e.Get(1)
					_ = V
					__e.Return(MakeNative(func(__e Evaluator) {
						Acc := __e.Get(1)
						_ = Acc
						__e.TailApply(__e.Global(symcons), V, Acc)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			__e.TailApply(__e.Global(symshen_4dict_1fold), gen443, V2409, Nil)

			return

		}, 1)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.dict-values"), gen444)

		return

	}, 0))
}
