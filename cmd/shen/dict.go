package main

import . "github.com/tiancaiamao/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp68239 := MakeNative(func(__e *ControlFlow) {
		V2342 := __e.Get(1)
		_ = V2342
		tmp68265 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

		tmp68266 := Call(__e, tmp68265, V2342, MakeNumber(1))

		if True == tmp68266 {
			tmp68260 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp68261 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp68262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp68263 := Call(__e, tmp68262, V2342, MakeString(""), symshen_4s)

			tmp68264 := Call(__e, tmp68261, MakeString("invalid initial dict size: "), tmp68263)

			__e.TailApply(tmp68260, tmp68264)
			return

		} else {
			tmp68241 := MakeNative(func(__e *ControlFlow) {
				D := __e.Get(1)
				_ = D
				tmp68242 := MakeNative(func(__e *ControlFlow) {
					Tag := __e.Get(1)
					_ = Tag
					tmp68243 := MakeNative(func(__e *ControlFlow) {
						Capacity := __e.Get(1)
						_ = Capacity
						tmp68244 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							tmp68245 := MakeNative(func(__e *ControlFlow) {
								Fill := __e.Get(1)
								_ = Fill
								__e.Return(D)
								return
							}, 1)

							tmp68246 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

							tmp68247 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

							tmp68248 := Call(__e, tmp68247, MakeNumber(2), V2342)

							tmp68249 := Call(__e, tmp68246, D, MakeNumber(3), tmp68248, Nil)

							__e.TailApply(tmp68245, tmp68249)
							return

						}, 1)

						tmp68250 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

						tmp68251 := Call(__e, tmp68250, D, MakeNumber(2), MakeNumber(0))

						__e.TailApply(tmp68244, tmp68251)
						return

					}, 1)

					tmp68252 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

					tmp68253 := Call(__e, tmp68252, D, MakeNumber(1), V2342)

					__e.TailApply(tmp68243, tmp68253)
					return

				}, 1)

				tmp68254 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

				tmp68255 := Call(__e, tmp68254, D, MakeNumber(0), symshen_4dictionary)

				__e.TailApply(tmp68242, tmp68255)
				return

			}, 1)

			tmp68256 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

			tmp68257 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp68258 := Call(__e, tmp68257, MakeNumber(3), V2342)

			tmp68259 := Call(__e, tmp68256, tmp68258)

			__e.TailApply(tmp68241, tmp68259)
			return

		}

	}, 1)

	tmp68267 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict, tmp68239)

	_ = tmp68267

	tmp68268 := MakeNative(func(__e *ControlFlow) {
		V2344 := __e.Get(1)
		_ = V2344
		tmp68277 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		tmp68278 := Call(__e, tmp68277, V2344)

		if True == tmp68278 {
			tmp68271 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68272 := MakeNative(func(__e *ControlFlow) {
				tmp68273 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(tmp68273, V2344, MakeNumber(0))
				return

			}, 0)

			tmp68274 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1dictionary)
				return
			}, 1)

			tmp68275 := Call(__e, PrimNS1Value(symtry_1catch), tmp68272, tmp68274)

			tmp68276 := Call(__e, tmp68271, tmp68275, symshen_4dictionary)

			if True == tmp68276 {
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

	tmp68279 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_2, tmp68268)

	_ = tmp68279

	tmp68280 := MakeNative(func(__e *ControlFlow) {
		V2346 := __e.Get(1)
		_ = V2346
		tmp68281 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(tmp68281, V2346, MakeNumber(1))
		return

	}, 1)

	tmp68282 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1capacity, tmp68280)

	_ = tmp68282

	tmp68283 := MakeNative(func(__e *ControlFlow) {
		V2348 := __e.Get(1)
		_ = V2348
		tmp68284 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(tmp68284, V2348, MakeNumber(2))
		return

	}, 1)

	tmp68285 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1count, tmp68283)

	_ = tmp68285

	tmp68286 := MakeNative(func(__e *ControlFlow) {
		V2351 := __e.Get(1)
		_ = V2351
		V2352 := __e.Get(2)
		_ = V2352
		tmp68287 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		__e.TailApply(tmp68287, V2351, MakeNumber(2), V2352)
		return

	}, 2)

	tmp68288 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1count_1_6, tmp68286)

	_ = tmp68288

	tmp68289 := MakeNative(func(__e *ControlFlow) {
		V2355 := __e.Get(1)
		_ = V2355
		V2356 := __e.Get(2)
		_ = V2356
		tmp68290 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp68291 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp68292 := Call(__e, tmp68291, MakeNumber(3), V2356)

		__e.TailApply(tmp68290, V2355, tmp68292)
		return

	}, 2)

	tmp68293 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5_1dict_1bucket, tmp68289)

	_ = tmp68293

	tmp68294 := MakeNative(func(__e *ControlFlow) {
		V2360 := __e.Get(1)
		_ = V2360
		V2361 := __e.Get(2)
		_ = V2361
		V2362 := __e.Get(3)
		_ = V2362
		tmp68295 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

		tmp68296 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp68297 := Call(__e, tmp68296, MakeNumber(3), V2361)

		__e.TailApply(tmp68295, V2360, tmp68297, V2362)
		return

	}, 3)

	tmp68298 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1bucket_1_6, tmp68294)

	_ = tmp68298

	tmp68299 := MakeNative(func(__e *ControlFlow) {
		V2366 := __e.Get(1)
		_ = V2366
		V2367 := __e.Get(2)
		_ = V2367
		V2368 := __e.Get(3)
		_ = V2368
		tmp68300 := MakeNative(func(__e *ControlFlow) {
			Diff := __e.Get(1)
			_ = Diff
			tmp68301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1count_1_6)

			tmp68302 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp68303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1count)

			tmp68304 := Call(__e, tmp68303, V2366)

			tmp68305 := Call(__e, tmp68302, Diff, tmp68304)

			__e.TailApply(tmp68301, V2366, tmp68305)
			return

		}, 1)

		tmp68306 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

		tmp68307 := Call(__e, PrimNS1Value(symns2_1value), symlength)

		tmp68308 := Call(__e, tmp68307, V2368)

		tmp68309 := Call(__e, PrimNS1Value(symns2_1value), symlength)

		tmp68310 := Call(__e, tmp68309, V2367)

		tmp68311 := Call(__e, tmp68306, tmp68308, tmp68310)

		__e.TailApply(tmp68300, tmp68311)
		return

	}, 3)

	tmp68312 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1update_1count, tmp68299)

	_ = tmp68312

	tmp68313 := MakeNative(func(__e *ControlFlow) {
		V2372 := __e.Get(1)
		_ = V2372
		V2373 := __e.Get(2)
		_ = V2373
		V2374 := __e.Get(3)
		_ = V2374
		tmp68314 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp68315 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				tmp68316 := MakeNative(func(__e *ControlFlow) {
					NewBucket := __e.Get(1)
					_ = NewBucket
					tmp68317 := MakeNative(func(__e *ControlFlow) {
						Change := __e.Get(1)
						_ = Change
						tmp68318 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							__e.Return(V2374)
							return
						}, 1)

						tmp68319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1update_1count)

						tmp68320 := Call(__e, tmp68319, V2372, Bucket, NewBucket)

						__e.TailApply(tmp68318, tmp68320)
						return

					}, 1)

					tmp68321 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1bucket_1_6)

					tmp68322 := Call(__e, tmp68321, V2372, N, NewBucket)

					__e.TailApply(tmp68317, tmp68322)
					return

				}, 1)

				tmp68323 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1set)

				tmp68324 := Call(__e, tmp68323, V2373, V2374, Bucket)

				__e.TailApply(tmp68316, tmp68324)
				return

			}, 1)

			tmp68325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			tmp68326 := Call(__e, tmp68325, V2372, N)

			__e.TailApply(tmp68315, tmp68326)
			return

		}, 1)

		tmp68327 := Call(__e, PrimNS1Value(symns2_1value), symhash)

		tmp68328 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		tmp68329 := Call(__e, tmp68328, V2372)

		tmp68330 := Call(__e, tmp68327, V2373, tmp68329)

		__e.TailApply(tmp68314, tmp68330)
		return

	}, 3)

	tmp68331 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1_6, tmp68313)

	_ = tmp68331

	tmp68332 := MakeNative(func(__e *ControlFlow) {
		V2377 := __e.Get(1)
		_ = V2377
		V2378 := __e.Get(2)
		_ = V2378
		tmp68333 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp68334 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				tmp68335 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp68343 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

					tmp68344 := Call(__e, tmp68343, Result)

					if True == tmp68344 {
						tmp68338 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						tmp68339 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						tmp68340 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						tmp68341 := Call(__e, tmp68340, V2378, MakeString(" not found in dict\n"), symshen_4a)

						tmp68342 := Call(__e, tmp68339, MakeString("value "), tmp68341)

						__e.TailApply(tmp68338, tmp68342)
						return

					} else {
						tmp68337 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						__e.TailApply(tmp68337, Result)
						return

					}

				}, 1)

				tmp68345 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

				tmp68346 := Call(__e, tmp68345, V2378, Bucket)

				__e.TailApply(tmp68335, tmp68346)
				return

			}, 1)

			tmp68347 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			tmp68348 := Call(__e, tmp68347, V2377, N)

			__e.TailApply(tmp68334, tmp68348)
			return

		}, 1)

		tmp68349 := Call(__e, PrimNS1Value(symns2_1value), symhash)

		tmp68350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		tmp68351 := Call(__e, tmp68350, V2377)

		tmp68352 := Call(__e, tmp68349, V2378, tmp68351)

		__e.TailApply(tmp68333, tmp68352)
		return

	}, 2)

	tmp68353 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5_1dict, tmp68332)

	_ = tmp68353

	tmp68354 := MakeNative(func(__e *ControlFlow) {
		V2381 := __e.Get(1)
		_ = V2381
		V2382 := __e.Get(2)
		_ = V2382
		tmp68355 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp68356 := MakeNative(func(__e *ControlFlow) {
				Bucket := __e.Get(1)
				_ = Bucket
				tmp68357 := MakeNative(func(__e *ControlFlow) {
					NewBucket := __e.Get(1)
					_ = NewBucket
					tmp68358 := MakeNative(func(__e *ControlFlow) {
						Change := __e.Get(1)
						_ = Change
						tmp68359 := MakeNative(func(__e *ControlFlow) {
							Count := __e.Get(1)
							_ = Count
							__e.Return(V2382)
							return
						}, 1)

						tmp68360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1update_1count)

						tmp68361 := Call(__e, tmp68360, V2381, Bucket, NewBucket)

						__e.TailApply(tmp68359, tmp68361)
						return

					}, 1)

					tmp68362 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1bucket_1_6)

					tmp68363 := Call(__e, tmp68362, V2381, N, NewBucket)

					__e.TailApply(tmp68358, tmp68363)
					return

				}, 1)

				tmp68364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1rm)

				tmp68365 := Call(__e, tmp68364, V2382, Bucket)

				__e.TailApply(tmp68357, tmp68365)
				return

			}, 1)

			tmp68366 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			tmp68367 := Call(__e, tmp68366, V2381, N)

			__e.TailApply(tmp68356, tmp68367)
			return

		}, 1)

		tmp68368 := Call(__e, PrimNS1Value(symns2_1value), symhash)

		tmp68369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		tmp68370 := Call(__e, tmp68369, V2381)

		tmp68371 := Call(__e, tmp68368, V2382, tmp68370)

		__e.TailApply(tmp68355, tmp68371)
		return

	}, 2)

	tmp68372 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1rm, tmp68354)

	_ = tmp68372

	tmp68373 := MakeNative(func(__e *ControlFlow) {
		V2386 := __e.Get(1)
		_ = V2386
		V2387 := __e.Get(2)
		_ = V2387
		V2388 := __e.Get(3)
		_ = V2388
		tmp68374 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp68375 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold_1h)

			__e.TailApply(tmp68375, V2386, V2387, V2388, MakeNumber(0), Limit)
			return

		}, 1)

		tmp68376 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1capacity)

		tmp68377 := Call(__e, tmp68376, V2387)

		__e.TailApply(tmp68374, tmp68377)
		return

	}, 3)

	tmp68378 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1fold, tmp68373)

	_ = tmp68378

	tmp68379 := MakeNative(func(__e *ControlFlow) {
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
		tmp68390 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68391 := Call(__e, tmp68390, V2399, V2398)

		if True == tmp68391 {
			__e.Return(V2397)
			return
		} else {
			tmp68381 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp68382 := MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					tmp68383 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold_1h)

					tmp68384 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					tmp68385 := Call(__e, tmp68384, MakeNumber(1), V2398)

					__e.TailApply(tmp68383, V2395, V2396, Acc, tmp68385, V2399)
					return

				}, 1)

				tmp68386 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bucket_1fold)

				tmp68387 := Call(__e, tmp68386, V2395, B, V2397)

				__e.TailApply(tmp68382, tmp68387)
				return

			}, 1)

			tmp68388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict_1bucket)

			tmp68389 := Call(__e, tmp68388, V2396, V2398)

			__e.TailApply(tmp68381, tmp68389)
			return

		}

	}, 5)

	tmp68392 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1fold_1h, tmp68379)

	_ = tmp68392

	tmp68393 := MakeNative(func(__e *ControlFlow) {
		V2403 := __e.Get(1)
		_ = V2403
		V2404 := __e.Get(2)
		_ = V2404
		V2405 := __e.Get(3)
		_ = V2405
		tmp68417 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68418 := Call(__e, tmp68417, Nil, V2404)

		if True == tmp68418 {
			__e.Return(V2405)
			return
		} else {
			tmp68415 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp68416 := Call(__e, tmp68415, V2404)

			var ifres68409 Obj

			if True == tmp68416 {
				tmp68411 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68412 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68413 := Call(__e, tmp68412, V2404)

				tmp68414 := Call(__e, tmp68411, tmp68413)

				var ifres68410 Obj

				if True == tmp68414 {
					ifres68410 = True

				} else {
					ifres68410 = False

				}

				ifres68409 = ifres68410

			} else {
				ifres68409 = False

			}

			if True == ifres68409 {
				tmp68397 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68398 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68399 := Call(__e, tmp68398, V2404)

				tmp68400 := Call(__e, tmp68397, tmp68399)

				tmp68401 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68402 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68403 := Call(__e, tmp68402, V2404)

				tmp68404 := Call(__e, tmp68401, tmp68403)

				tmp68405 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bucket_1fold)

				tmp68406 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68407 := Call(__e, tmp68406, V2404)

				tmp68408 := Call(__e, tmp68405, V2403, tmp68407, V2405)

				__e.TailApply(V2403, tmp68400, tmp68404, tmp68408)
				return

			} else {
				tmp68396 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp68396, symshen_4bucket_1fold)
				return

			}

		}

	}, 3)

	tmp68419 := Call(__e, PrimNS1Value(symns2_1set), symshen_4bucket_1fold, tmp68393)

	_ = tmp68419

	tmp68420 := MakeNative(func(__e *ControlFlow) {
		V2407 := __e.Get(1)
		_ = V2407
		tmp68421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold)

		tmp68422 := MakeNative(func(__e *ControlFlow) {
			K := __e.Get(1)
			_ = K
			__e.Return(MakeNative(func(__e *ControlFlow) {
				__ := __e.Get(1)
				_ = __
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					tmp68423 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					__e.TailApply(tmp68423, K, Acc)
					return

				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(tmp68421, tmp68422, V2407, Nil)
		return

	}, 1)

	tmp68424 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dict_1keys, tmp68420)

	_ = tmp68424

	tmp68425 := MakeNative(func(__e *ControlFlow) {
		V2409 := __e.Get(1)
		_ = V2409
		tmp68426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1fold)

		tmp68427 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V := __e.Get(1)
				_ = V
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Acc := __e.Get(1)
					_ = Acc
					tmp68428 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					__e.TailApply(tmp68428, V, Acc)
					return

				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(tmp68426, tmp68427, V2409, Nil)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4dict_1values, tmp68425)
	return

}, 0)
