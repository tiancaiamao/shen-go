package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
	tmp10449 := MakeNative(func(__e *ControlFlow) {
		V3341 := __e.Get(1)
		_ = V3341
		tmp10450 := MakeNative(func(__e *ControlFlow) {
			W3342 := __e.Get(1)
			_ = W3342
			tmp10452 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3342)

			if True == tmp10452 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3342)
				return
			}

		}, 1)

		tmp10472 := PrimIsPair(V3341)

		var ifres10453 Obj

		if True == tmp10472 {
			tmp10454 := MakeNative(func(__e *ControlFlow) {
				W3343 := __e.Get(1)
				_ = W3343
				tmp10455 := MakeNative(func(__e *ControlFlow) {
					W3344 := __e.Get(1)
					_ = W3344
					tmp10456 := MakeNative(func(__e *ControlFlow) {
						W3345 := __e.Get(1)
						_ = W3345
						tmp10466 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3345)

						if True == tmp10466 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp10457 := MakeNative(func(__e *ControlFlow) {
								W3346 := __e.Get(1)
								_ = W3346
								tmp10458 := MakeNative(func(__e *ControlFlow) {
									W3347 := __e.Get(1)
									_ = W3347
									tmp10459 := MakeNative(func(__e *ControlFlow) {
										W3348 := __e.Get(1)
										_ = W3348
										tmp10460 := Call(__e, PrimFunc(symfn), W3343)

										__e.TailApply(PrimFunc(symshen_4remember_1datatype), W3343, tmp10460)
										return

									}, 1)

									tmp10461 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W3343, W3346)

									tmp10462 := Call(__e, tmp10459, tmp10461)

									__e.TailApply(PrimFunc(symshen_4comb), W3347, tmp10462)
									return

								}, 1)

								tmp10463 := Call(__e, PrimFunc(symshen_4in_1_6), W3345)

								__e.TailApply(tmp10458, tmp10463)
								return

							}, 1)

							tmp10464 := Call(__e, PrimFunc(symshen_4_5_1out), W3345)

							__e.TailApply(tmp10457, tmp10464)
							return

						}

					}, 1)

					tmp10467 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3344)

					__e.TailApply(tmp10456, tmp10467)
					return

				}, 1)

				tmp10468 := Call(__e, PrimFunc(symtail), V3341)

				__e.TailApply(tmp10455, tmp10468)
				return

			}, 1)

			tmp10469 := Call(__e, PrimFunc(symhead), V3341)

			tmp10470 := Call(__e, tmp10454, tmp10469)

			ifres10453 = tmp10470

		} else {
			tmp10471 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10453 = tmp10471

		}

		__e.TailApply(tmp10450, ifres10453)
		return

	}, 1)

	tmp10473 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp10449)

	_ = tmp10473

	tmp10474 := MakeNative(func(__e *ControlFlow) {
		V3349 := __e.Get(1)
		_ = V3349
		V3350 := __e.Get(2)
		_ = V3350
		tmp10475 := PrimValue(symshen_4_ddatatypes_d)

		tmp10476 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3349, V3350, tmp10475)

		tmp10477 := PrimSet(symshen_4_ddatatypes_d, tmp10476)

		_ = tmp10477

		tmp10478 := PrimValue(symshen_4_dalldatatypes_d)

		tmp10479 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3349, V3350, tmp10478)

		tmp10480 := PrimSet(symshen_4_dalldatatypes_d, tmp10479)

		_ = tmp10480

		__e.Return(V3349)
		return

	}, 2)

	tmp10481 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp10474)

	_ = tmp10481

	tmp10482 := MakeNative(func(__e *ControlFlow) {
		V3351 := __e.Get(1)
		_ = V3351
		tmp10483 := MakeNative(func(__e *ControlFlow) {
			W3352 := __e.Get(1)
			_ = W3352
			tmp10502 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3352)

			if True == tmp10502 {
				tmp10484 := MakeNative(func(__e *ControlFlow) {
					W3359 := __e.Get(1)
					_ = W3359
					tmp10486 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3359)

					if True == tmp10486 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3359)
						return
					}

				}, 1)

				tmp10487 := MakeNative(func(__e *ControlFlow) {
					W3360 := __e.Get(1)
					_ = W3360
					tmp10498 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3360)

					if True == tmp10498 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10488 := MakeNative(func(__e *ControlFlow) {
							W3361 := __e.Get(1)
							_ = W3361
							tmp10489 := MakeNative(func(__e *ControlFlow) {
								W3362 := __e.Get(1)
								_ = W3362
								tmp10494 := Call(__e, PrimFunc(symempty_2), W3361)

								var ifres10490 Obj

								if True == tmp10494 {
									ifres10490 = Nil

								} else {
									tmp10491 := Call(__e, PrimFunc(symshen_4app), W3361, MakeString("\n ..."), symshen_4r)

									tmp10492 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp10491)

									tmp10493 := PrimSimpleError(tmp10492)

									ifres10490 = tmp10493

								}

								__e.TailApply(PrimFunc(symshen_4comb), W3362, ifres10490)
								return

							}, 1)

							tmp10495 := Call(__e, PrimFunc(symshen_4in_1_6), W3360)

							__e.TailApply(tmp10489, tmp10495)
							return

						}, 1)

						tmp10496 := Call(__e, PrimFunc(symshen_4_5_1out), W3360)

						__e.TailApply(tmp10488, tmp10496)
						return

					}

				}, 1)

				tmp10499 := Call(__e, PrimFunc(sym_5_b_6), V3351)

				tmp10500 := Call(__e, tmp10487, tmp10499)

				__e.TailApply(tmp10484, tmp10500)
				return

			} else {
				__e.Return(W3352)
				return
			}

		}, 1)

		tmp10503 := MakeNative(func(__e *ControlFlow) {
			W3353 := __e.Get(1)
			_ = W3353
			tmp10518 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3353)

			if True == tmp10518 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10504 := MakeNative(func(__e *ControlFlow) {
					W3354 := __e.Get(1)
					_ = W3354
					tmp10505 := MakeNative(func(__e *ControlFlow) {
						W3355 := __e.Get(1)
						_ = W3355
						tmp10506 := MakeNative(func(__e *ControlFlow) {
							W3356 := __e.Get(1)
							_ = W3356
							tmp10513 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3356)

							if True == tmp10513 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10507 := MakeNative(func(__e *ControlFlow) {
									W3357 := __e.Get(1)
									_ = W3357
									tmp10508 := MakeNative(func(__e *ControlFlow) {
										W3358 := __e.Get(1)
										_ = W3358
										tmp10509 := Call(__e, PrimFunc(symappend), W3354, W3357)

										__e.TailApply(PrimFunc(symshen_4comb), W3358, tmp10509)
										return

									}, 1)

									tmp10510 := Call(__e, PrimFunc(symshen_4in_1_6), W3356)

									__e.TailApply(tmp10508, tmp10510)
									return

								}, 1)

								tmp10511 := Call(__e, PrimFunc(symshen_4_5_1out), W3356)

								__e.TailApply(tmp10507, tmp10511)
								return

							}

						}, 1)

						tmp10514 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3355)

						__e.TailApply(tmp10506, tmp10514)
						return

					}, 1)

					tmp10515 := Call(__e, PrimFunc(symshen_4in_1_6), W3353)

					__e.TailApply(tmp10505, tmp10515)
					return

				}, 1)

				tmp10516 := Call(__e, PrimFunc(symshen_4_5_1out), W3353)

				__e.TailApply(tmp10504, tmp10516)
				return

			}

		}, 1)

		tmp10519 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3351)

		tmp10520 := Call(__e, tmp10503, tmp10519)

		__e.TailApply(tmp10483, tmp10520)
		return

	}, 1)

	tmp10521 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp10482)

	_ = tmp10521

	tmp10522 := MakeNative(func(__e *ControlFlow) {
		V3363 := __e.Get(1)
		_ = V3363
		tmp10523 := MakeNative(func(__e *ControlFlow) {
			W3364 := __e.Get(1)
			_ = W3364
			tmp10537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3364)

			if True == tmp10537 {
				tmp10524 := MakeNative(func(__e *ControlFlow) {
					W3368 := __e.Get(1)
					_ = W3368
					tmp10526 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3368)

					if True == tmp10526 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3368)
						return
					}

				}, 1)

				tmp10527 := MakeNative(func(__e *ControlFlow) {
					W3369 := __e.Get(1)
					_ = W3369
					tmp10533 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3369)

					if True == tmp10533 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10528 := MakeNative(func(__e *ControlFlow) {
							W3370 := __e.Get(1)
							_ = W3370
							tmp10529 := MakeNative(func(__e *ControlFlow) {
								W3371 := __e.Get(1)
								_ = W3371
								__e.TailApply(PrimFunc(symshen_4comb), W3371, W3370)
								return
							}, 1)

							tmp10530 := Call(__e, PrimFunc(symshen_4in_1_6), W3369)

							__e.TailApply(tmp10529, tmp10530)
							return

						}, 1)

						tmp10531 := Call(__e, PrimFunc(symshen_4_5_1out), W3369)

						__e.TailApply(tmp10528, tmp10531)
						return

					}

				}, 1)

				tmp10534 := Call(__e, PrimFunc(symshen_4_5double_6), V3363)

				tmp10535 := Call(__e, tmp10527, tmp10534)

				__e.TailApply(tmp10524, tmp10535)
				return

			} else {
				__e.Return(W3364)
				return
			}

		}, 1)

		tmp10538 := MakeNative(func(__e *ControlFlow) {
			W3365 := __e.Get(1)
			_ = W3365
			tmp10544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3365)

			if True == tmp10544 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10539 := MakeNative(func(__e *ControlFlow) {
					W3366 := __e.Get(1)
					_ = W3366
					tmp10540 := MakeNative(func(__e *ControlFlow) {
						W3367 := __e.Get(1)
						_ = W3367
						__e.TailApply(PrimFunc(symshen_4comb), W3367, W3366)
						return
					}, 1)

					tmp10541 := Call(__e, PrimFunc(symshen_4in_1_6), W3365)

					__e.TailApply(tmp10540, tmp10541)
					return

				}, 1)

				tmp10542 := Call(__e, PrimFunc(symshen_4_5_1out), W3365)

				__e.TailApply(tmp10539, tmp10542)
				return

			}

		}, 1)

		tmp10545 := Call(__e, PrimFunc(symshen_4_5single_6), V3363)

		tmp10546 := Call(__e, tmp10538, tmp10545)

		__e.TailApply(tmp10523, tmp10546)
		return

	}, 1)

	tmp10547 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp10522)

	_ = tmp10547

	tmp10548 := MakeNative(func(__e *ControlFlow) {
		V3372 := __e.Get(1)
		_ = V3372
		tmp10549 := MakeNative(func(__e *ControlFlow) {
			W3373 := __e.Get(1)
			_ = W3373
			tmp10551 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3373)

			if True == tmp10551 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3373)
				return
			}

		}, 1)

		tmp10552 := MakeNative(func(__e *ControlFlow) {
			W3374 := __e.Get(1)
			_ = W3374
			tmp10590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3374)

			if True == tmp10590 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10553 := MakeNative(func(__e *ControlFlow) {
					W3375 := __e.Get(1)
					_ = W3375
					tmp10554 := MakeNative(func(__e *ControlFlow) {
						W3376 := __e.Get(1)
						_ = W3376
						tmp10555 := MakeNative(func(__e *ControlFlow) {
							W3377 := __e.Get(1)
							_ = W3377
							tmp10585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3377)

							if True == tmp10585 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10556 := MakeNative(func(__e *ControlFlow) {
									W3378 := __e.Get(1)
									_ = W3378
									tmp10557 := MakeNative(func(__e *ControlFlow) {
										W3379 := __e.Get(1)
										_ = W3379
										tmp10558 := MakeNative(func(__e *ControlFlow) {
											W3380 := __e.Get(1)
											_ = W3380
											tmp10580 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3380)

											if True == tmp10580 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp10559 := MakeNative(func(__e *ControlFlow) {
													W3381 := __e.Get(1)
													_ = W3381
													tmp10560 := MakeNative(func(__e *ControlFlow) {
														W3382 := __e.Get(1)
														_ = W3382
														tmp10576 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3382)

														if True == tmp10576 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp10561 := MakeNative(func(__e *ControlFlow) {
																W3383 := __e.Get(1)
																_ = W3383
																tmp10562 := MakeNative(func(__e *ControlFlow) {
																	W3384 := __e.Get(1)
																	_ = W3384
																	tmp10563 := MakeNative(func(__e *ControlFlow) {
																		W3385 := __e.Get(1)
																		_ = W3385
																		tmp10571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3385)

																		if True == tmp10571 {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		} else {
																			tmp10564 := MakeNative(func(__e *ControlFlow) {
																				W3386 := __e.Get(1)
																				_ = W3386
																				tmp10565 := PrimCons(W3383, Nil)

																				tmp10566 := PrimCons(W3378, tmp10565)

																				tmp10567 := PrimCons(W3375, tmp10566)

																				tmp10568 := PrimCons(tmp10567, Nil)

																				__e.TailApply(PrimFunc(symshen_4comb), W3386, tmp10568)
																				return

																			}, 1)

																			tmp10569 := Call(__e, PrimFunc(symshen_4in_1_6), W3385)

																			__e.TailApply(tmp10564, tmp10569)
																			return

																		}

																	}, 1)

																	tmp10572 := Call(__e, PrimFunc(symshen_4_5sc_6), W3384)

																	__e.TailApply(tmp10563, tmp10572)
																	return

																}, 1)

																tmp10573 := Call(__e, PrimFunc(symshen_4in_1_6), W3382)

																__e.TailApply(tmp10562, tmp10573)
																return

															}, 1)

															tmp10574 := Call(__e, PrimFunc(symshen_4_5_1out), W3382)

															__e.TailApply(tmp10561, tmp10574)
															return

														}

													}, 1)

													tmp10577 := Call(__e, PrimFunc(symshen_4_5conc_6), W3381)

													__e.TailApply(tmp10560, tmp10577)
													return

												}, 1)

												tmp10578 := Call(__e, PrimFunc(symshen_4in_1_6), W3380)

												__e.TailApply(tmp10559, tmp10578)
												return

											}

										}, 1)

										tmp10581 := Call(__e, PrimFunc(symshen_4_5sng_6), W3379)

										__e.TailApply(tmp10558, tmp10581)
										return

									}, 1)

									tmp10582 := Call(__e, PrimFunc(symshen_4in_1_6), W3377)

									__e.TailApply(tmp10557, tmp10582)
									return

								}, 1)

								tmp10583 := Call(__e, PrimFunc(symshen_4_5_1out), W3377)

								__e.TailApply(tmp10556, tmp10583)
								return

							}

						}, 1)

						tmp10586 := Call(__e, PrimFunc(symshen_4_5prems_6), W3376)

						__e.TailApply(tmp10555, tmp10586)
						return

					}, 1)

					tmp10587 := Call(__e, PrimFunc(symshen_4in_1_6), W3374)

					__e.TailApply(tmp10554, tmp10587)
					return

				}, 1)

				tmp10588 := Call(__e, PrimFunc(symshen_4_5_1out), W3374)

				__e.TailApply(tmp10553, tmp10588)
				return

			}

		}, 1)

		tmp10591 := Call(__e, PrimFunc(symshen_4_5sides_6), V3372)

		tmp10592 := Call(__e, tmp10552, tmp10591)

		__e.TailApply(tmp10549, tmp10592)
		return

	}, 1)

	tmp10593 := Call(__e, ns2_1set, symshen_4_5single_6, tmp10548)

	_ = tmp10593

	tmp10594 := MakeNative(func(__e *ControlFlow) {
		V3387 := __e.Get(1)
		_ = V3387
		tmp10595 := MakeNative(func(__e *ControlFlow) {
			W3388 := __e.Get(1)
			_ = W3388
			tmp10597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3388)

			if True == tmp10597 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3388)
				return
			}

		}, 1)

		tmp10598 := MakeNative(func(__e *ControlFlow) {
			W3389 := __e.Get(1)
			_ = W3389
			tmp10635 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3389)

			if True == tmp10635 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10599 := MakeNative(func(__e *ControlFlow) {
					W3390 := __e.Get(1)
					_ = W3390
					tmp10600 := MakeNative(func(__e *ControlFlow) {
						W3391 := __e.Get(1)
						_ = W3391
						tmp10601 := MakeNative(func(__e *ControlFlow) {
							W3392 := __e.Get(1)
							_ = W3392
							tmp10630 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3392)

							if True == tmp10630 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10602 := MakeNative(func(__e *ControlFlow) {
									W3393 := __e.Get(1)
									_ = W3393
									tmp10603 := MakeNative(func(__e *ControlFlow) {
										W3394 := __e.Get(1)
										_ = W3394
										tmp10604 := MakeNative(func(__e *ControlFlow) {
											W3395 := __e.Get(1)
											_ = W3395
											tmp10625 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3395)

											if True == tmp10625 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp10605 := MakeNative(func(__e *ControlFlow) {
													W3396 := __e.Get(1)
													_ = W3396
													tmp10606 := MakeNative(func(__e *ControlFlow) {
														W3397 := __e.Get(1)
														_ = W3397
														tmp10621 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3397)

														if True == tmp10621 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp10607 := MakeNative(func(__e *ControlFlow) {
																W3398 := __e.Get(1)
																_ = W3398
																tmp10608 := MakeNative(func(__e *ControlFlow) {
																	W3399 := __e.Get(1)
																	_ = W3399
																	tmp10609 := MakeNative(func(__e *ControlFlow) {
																		W3400 := __e.Get(1)
																		_ = W3400
																		tmp10616 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3400)

																		if True == tmp10616 {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		} else {
																			tmp10610 := MakeNative(func(__e *ControlFlow) {
																				W3401 := __e.Get(1)
																				_ = W3401
																				tmp10611 := PrimCons(W3398, Nil)

																				tmp10612 := PrimCons(Nil, tmp10611)

																				tmp10613 := Call(__e, PrimFunc(symshen_4lr_1rule), W3390, W3393, tmp10612)

																				__e.TailApply(PrimFunc(symshen_4comb), W3401, tmp10613)
																				return

																			}, 1)

																			tmp10614 := Call(__e, PrimFunc(symshen_4in_1_6), W3400)

																			__e.TailApply(tmp10610, tmp10614)
																			return

																		}

																	}, 1)

																	tmp10617 := Call(__e, PrimFunc(symshen_4_5sc_6), W3399)

																	__e.TailApply(tmp10609, tmp10617)
																	return

																}, 1)

																tmp10618 := Call(__e, PrimFunc(symshen_4in_1_6), W3397)

																__e.TailApply(tmp10608, tmp10618)
																return

															}, 1)

															tmp10619 := Call(__e, PrimFunc(symshen_4_5_1out), W3397)

															__e.TailApply(tmp10607, tmp10619)
															return

														}

													}, 1)

													tmp10622 := Call(__e, PrimFunc(symshen_4_5formula_6), W3396)

													__e.TailApply(tmp10606, tmp10622)
													return

												}, 1)

												tmp10623 := Call(__e, PrimFunc(symshen_4in_1_6), W3395)

												__e.TailApply(tmp10605, tmp10623)
												return

											}

										}, 1)

										tmp10626 := Call(__e, PrimFunc(symshen_4_5dbl_6), W3394)

										__e.TailApply(tmp10604, tmp10626)
										return

									}, 1)

									tmp10627 := Call(__e, PrimFunc(symshen_4in_1_6), W3392)

									__e.TailApply(tmp10603, tmp10627)
									return

								}, 1)

								tmp10628 := Call(__e, PrimFunc(symshen_4_5_1out), W3392)

								__e.TailApply(tmp10602, tmp10628)
								return

							}

						}, 1)

						tmp10631 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3391)

						__e.TailApply(tmp10601, tmp10631)
						return

					}, 1)

					tmp10632 := Call(__e, PrimFunc(symshen_4in_1_6), W3389)

					__e.TailApply(tmp10600, tmp10632)
					return

				}, 1)

				tmp10633 := Call(__e, PrimFunc(symshen_4_5_1out), W3389)

				__e.TailApply(tmp10599, tmp10633)
				return

			}

		}, 1)

		tmp10636 := Call(__e, PrimFunc(symshen_4_5sides_6), V3387)

		tmp10637 := Call(__e, tmp10598, tmp10636)

		__e.TailApply(tmp10595, tmp10637)
		return

	}, 1)

	tmp10638 := Call(__e, ns2_1set, symshen_4_5double_6, tmp10594)

	_ = tmp10638

	tmp10639 := MakeNative(func(__e *ControlFlow) {
		V3402 := __e.Get(1)
		_ = V3402
		tmp10640 := MakeNative(func(__e *ControlFlow) {
			W3403 := __e.Get(1)
			_ = W3403
			tmp10663 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3403)

			if True == tmp10663 {
				tmp10641 := MakeNative(func(__e *ControlFlow) {
					W3412 := __e.Get(1)
					_ = W3412
					tmp10643 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3412)

					if True == tmp10643 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3412)
						return
					}

				}, 1)

				tmp10644 := MakeNative(func(__e *ControlFlow) {
					W3413 := __e.Get(1)
					_ = W3413
					tmp10659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3413)

					if True == tmp10659 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10645 := MakeNative(func(__e *ControlFlow) {
							W3414 := __e.Get(1)
							_ = W3414
							tmp10646 := MakeNative(func(__e *ControlFlow) {
								W3415 := __e.Get(1)
								_ = W3415
								tmp10647 := MakeNative(func(__e *ControlFlow) {
									W3416 := __e.Get(1)
									_ = W3416
									tmp10654 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3416)

									if True == tmp10654 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp10648 := MakeNative(func(__e *ControlFlow) {
											W3417 := __e.Get(1)
											_ = W3417
											tmp10649 := PrimCons(W3414, Nil)

											tmp10650 := PrimCons(Nil, tmp10649)

											tmp10651 := PrimCons(tmp10650, Nil)

											__e.TailApply(PrimFunc(symshen_4comb), W3417, tmp10651)
											return

										}, 1)

										tmp10652 := Call(__e, PrimFunc(symshen_4in_1_6), W3416)

										__e.TailApply(tmp10648, tmp10652)
										return

									}

								}, 1)

								tmp10655 := Call(__e, PrimFunc(symshen_4_5sc_6), W3415)

								__e.TailApply(tmp10647, tmp10655)
								return

							}, 1)

							tmp10656 := Call(__e, PrimFunc(symshen_4in_1_6), W3413)

							__e.TailApply(tmp10646, tmp10656)
							return

						}, 1)

						tmp10657 := Call(__e, PrimFunc(symshen_4_5_1out), W3413)

						__e.TailApply(tmp10645, tmp10657)
						return

					}

				}, 1)

				tmp10660 := Call(__e, PrimFunc(symshen_4_5formula_6), V3402)

				tmp10661 := Call(__e, tmp10644, tmp10660)

				__e.TailApply(tmp10641, tmp10661)
				return

			} else {
				__e.Return(W3403)
				return
			}

		}, 1)

		tmp10664 := MakeNative(func(__e *ControlFlow) {
			W3404 := __e.Get(1)
			_ = W3404
			tmp10687 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3404)

			if True == tmp10687 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10665 := MakeNative(func(__e *ControlFlow) {
					W3405 := __e.Get(1)
					_ = W3405
					tmp10666 := MakeNative(func(__e *ControlFlow) {
						W3406 := __e.Get(1)
						_ = W3406
						tmp10667 := MakeNative(func(__e *ControlFlow) {
							W3407 := __e.Get(1)
							_ = W3407
							tmp10682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3407)

							if True == tmp10682 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10668 := MakeNative(func(__e *ControlFlow) {
									W3408 := __e.Get(1)
									_ = W3408
									tmp10669 := MakeNative(func(__e *ControlFlow) {
										W3409 := __e.Get(1)
										_ = W3409
										tmp10678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3409)

										if True == tmp10678 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp10670 := MakeNative(func(__e *ControlFlow) {
												W3410 := __e.Get(1)
												_ = W3410
												tmp10671 := MakeNative(func(__e *ControlFlow) {
													W3411 := __e.Get(1)
													_ = W3411
													tmp10672 := PrimCons(W3405, Nil)

													tmp10673 := PrimCons(Nil, tmp10672)

													tmp10674 := PrimCons(tmp10673, W3410)

													__e.TailApply(PrimFunc(symshen_4comb), W3411, tmp10674)
													return

												}, 1)

												tmp10675 := Call(__e, PrimFunc(symshen_4in_1_6), W3409)

												__e.TailApply(tmp10671, tmp10675)
												return

											}, 1)

											tmp10676 := Call(__e, PrimFunc(symshen_4_5_1out), W3409)

											__e.TailApply(tmp10670, tmp10676)
											return

										}

									}, 1)

									tmp10679 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3408)

									__e.TailApply(tmp10669, tmp10679)
									return

								}, 1)

								tmp10680 := Call(__e, PrimFunc(symshen_4in_1_6), W3407)

								__e.TailApply(tmp10668, tmp10680)
								return

							}

						}, 1)

						tmp10683 := Call(__e, PrimFunc(symshen_4_5sc_6), W3406)

						__e.TailApply(tmp10667, tmp10683)
						return

					}, 1)

					tmp10684 := Call(__e, PrimFunc(symshen_4in_1_6), W3404)

					__e.TailApply(tmp10666, tmp10684)
					return

				}, 1)

				tmp10685 := Call(__e, PrimFunc(symshen_4_5_1out), W3404)

				__e.TailApply(tmp10665, tmp10685)
				return

			}

		}, 1)

		tmp10688 := Call(__e, PrimFunc(symshen_4_5formula_6), V3402)

		tmp10689 := Call(__e, tmp10664, tmp10688)

		__e.TailApply(tmp10640, tmp10689)
		return

	}, 1)

	tmp10690 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp10639)

	_ = tmp10690

	tmp10691 := MakeNative(func(__e *ControlFlow) {
		V3418 := __e.Get(1)
		_ = V3418
		tmp10692 := MakeNative(func(__e *ControlFlow) {
			W3419 := __e.Get(1)
			_ = W3419
			tmp10708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3419)

			if True == tmp10708 {
				tmp10693 := MakeNative(func(__e *ControlFlow) {
					W3427 := __e.Get(1)
					_ = W3427
					tmp10695 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3427)

					if True == tmp10695 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3427)
						return
					}

				}, 1)

				tmp10696 := MakeNative(func(__e *ControlFlow) {
					W3428 := __e.Get(1)
					_ = W3428
					tmp10704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3428)

					if True == tmp10704 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10697 := MakeNative(func(__e *ControlFlow) {
							W3429 := __e.Get(1)
							_ = W3429
							tmp10698 := MakeNative(func(__e *ControlFlow) {
								W3430 := __e.Get(1)
								_ = W3430
								tmp10699 := PrimCons(W3429, Nil)

								tmp10700 := PrimCons(Nil, tmp10699)

								__e.TailApply(PrimFunc(symshen_4comb), W3430, tmp10700)
								return

							}, 1)

							tmp10701 := Call(__e, PrimFunc(symshen_4in_1_6), W3428)

							__e.TailApply(tmp10698, tmp10701)
							return

						}, 1)

						tmp10702 := Call(__e, PrimFunc(symshen_4_5_1out), W3428)

						__e.TailApply(tmp10697, tmp10702)
						return

					}

				}, 1)

				tmp10705 := Call(__e, PrimFunc(symshen_4_5formula_6), V3418)

				tmp10706 := Call(__e, tmp10696, tmp10705)

				__e.TailApply(tmp10693, tmp10706)
				return

			} else {
				__e.Return(W3419)
				return
			}

		}, 1)

		tmp10709 := MakeNative(func(__e *ControlFlow) {
			W3420 := __e.Get(1)
			_ = W3420
			tmp10729 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3420)

			if True == tmp10729 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10710 := MakeNative(func(__e *ControlFlow) {
					W3421 := __e.Get(1)
					_ = W3421
					tmp10711 := MakeNative(func(__e *ControlFlow) {
						W3422 := __e.Get(1)
						_ = W3422
						tmp10725 := Call(__e, PrimFunc(symshen_4hds_a_2), W3422, sym_6_6)

						if True == tmp10725 {
							tmp10712 := MakeNative(func(__e *ControlFlow) {
								W3423 := __e.Get(1)
								_ = W3423
								tmp10713 := MakeNative(func(__e *ControlFlow) {
									W3424 := __e.Get(1)
									_ = W3424
									tmp10721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3424)

									if True == tmp10721 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp10714 := MakeNative(func(__e *ControlFlow) {
											W3425 := __e.Get(1)
											_ = W3425
											tmp10715 := MakeNative(func(__e *ControlFlow) {
												W3426 := __e.Get(1)
												_ = W3426
												tmp10716 := PrimCons(W3425, Nil)

												tmp10717 := PrimCons(W3421, tmp10716)

												__e.TailApply(PrimFunc(symshen_4comb), W3426, tmp10717)
												return

											}, 1)

											tmp10718 := Call(__e, PrimFunc(symshen_4in_1_6), W3424)

											__e.TailApply(tmp10715, tmp10718)
											return

										}, 1)

										tmp10719 := Call(__e, PrimFunc(symshen_4_5_1out), W3424)

										__e.TailApply(tmp10714, tmp10719)
										return

									}

								}, 1)

								tmp10722 := Call(__e, PrimFunc(symshen_4_5formula_6), W3423)

								__e.TailApply(tmp10713, tmp10722)
								return

							}, 1)

							tmp10723 := Call(__e, PrimFunc(symtail), W3422)

							__e.TailApply(tmp10712, tmp10723)
							return

						} else {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp10726 := Call(__e, PrimFunc(symshen_4in_1_6), W3420)

					__e.TailApply(tmp10711, tmp10726)
					return

				}, 1)

				tmp10727 := Call(__e, PrimFunc(symshen_4_5_1out), W3420)

				__e.TailApply(tmp10710, tmp10727)
				return

			}

		}, 1)

		tmp10730 := Call(__e, PrimFunc(symshen_4_5ass_6), V3418)

		tmp10731 := Call(__e, tmp10709, tmp10730)

		__e.TailApply(tmp10692, tmp10731)
		return

	}, 1)

	tmp10732 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp10691)

	_ = tmp10732

	tmp10733 := MakeNative(func(__e *ControlFlow) {
		V3431 := __e.Get(1)
		_ = V3431
		tmp10734 := MakeNative(func(__e *ControlFlow) {
			W3432 := __e.Get(1)
			_ = W3432
			tmp10746 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3432)

			if True == tmp10746 {
				tmp10735 := MakeNative(func(__e *ControlFlow) {
					W3441 := __e.Get(1)
					_ = W3441
					tmp10737 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3441)

					if True == tmp10737 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3441)
						return
					}

				}, 1)

				tmp10738 := MakeNative(func(__e *ControlFlow) {
					W3442 := __e.Get(1)
					_ = W3442
					tmp10742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3442)

					if True == tmp10742 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10739 := MakeNative(func(__e *ControlFlow) {
							W3443 := __e.Get(1)
							_ = W3443
							__e.TailApply(PrimFunc(symshen_4comb), W3443, Nil)
							return
						}, 1)

						tmp10740 := Call(__e, PrimFunc(symshen_4in_1_6), W3442)

						__e.TailApply(tmp10739, tmp10740)
						return

					}

				}, 1)

				tmp10743 := Call(__e, PrimFunc(sym_5e_6), V3431)

				tmp10744 := Call(__e, tmp10738, tmp10743)

				__e.TailApply(tmp10735, tmp10744)
				return

			} else {
				__e.Return(W3432)
				return
			}

		}, 1)

		tmp10747 := MakeNative(func(__e *ControlFlow) {
			W3433 := __e.Get(1)
			_ = W3433
			tmp10768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3433)

			if True == tmp10768 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10748 := MakeNative(func(__e *ControlFlow) {
					W3434 := __e.Get(1)
					_ = W3434
					tmp10749 := MakeNative(func(__e *ControlFlow) {
						W3435 := __e.Get(1)
						_ = W3435
						tmp10750 := MakeNative(func(__e *ControlFlow) {
							W3436 := __e.Get(1)
							_ = W3436
							tmp10763 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3436)

							if True == tmp10763 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10751 := MakeNative(func(__e *ControlFlow) {
									W3437 := __e.Get(1)
									_ = W3437
									tmp10752 := MakeNative(func(__e *ControlFlow) {
										W3438 := __e.Get(1)
										_ = W3438
										tmp10759 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3438)

										if True == tmp10759 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp10753 := MakeNative(func(__e *ControlFlow) {
												W3439 := __e.Get(1)
												_ = W3439
												tmp10754 := MakeNative(func(__e *ControlFlow) {
													W3440 := __e.Get(1)
													_ = W3440
													tmp10755 := PrimCons(W3434, W3439)

													__e.TailApply(PrimFunc(symshen_4comb), W3440, tmp10755)
													return

												}, 1)

												tmp10756 := Call(__e, PrimFunc(symshen_4in_1_6), W3438)

												__e.TailApply(tmp10754, tmp10756)
												return

											}, 1)

											tmp10757 := Call(__e, PrimFunc(symshen_4_5_1out), W3438)

											__e.TailApply(tmp10753, tmp10757)
											return

										}

									}, 1)

									tmp10760 := Call(__e, PrimFunc(symshen_4_5prems_6), W3437)

									__e.TailApply(tmp10752, tmp10760)
									return

								}, 1)

								tmp10761 := Call(__e, PrimFunc(symshen_4in_1_6), W3436)

								__e.TailApply(tmp10751, tmp10761)
								return

							}

						}, 1)

						tmp10764 := Call(__e, PrimFunc(symshen_4_5sc_6), W3435)

						__e.TailApply(tmp10750, tmp10764)
						return

					}, 1)

					tmp10765 := Call(__e, PrimFunc(symshen_4in_1_6), W3433)

					__e.TailApply(tmp10749, tmp10765)
					return

				}, 1)

				tmp10766 := Call(__e, PrimFunc(symshen_4_5_1out), W3433)

				__e.TailApply(tmp10748, tmp10766)
				return

			}

		}, 1)

		tmp10769 := Call(__e, PrimFunc(symshen_4_5prem_6), V3431)

		tmp10770 := Call(__e, tmp10747, tmp10769)

		__e.TailApply(tmp10734, tmp10770)
		return

	}, 1)

	tmp10771 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp10733)

	_ = tmp10771

	tmp10772 := MakeNative(func(__e *ControlFlow) {
		V3444 := __e.Get(1)
		_ = V3444
		tmp10773 := MakeNative(func(__e *ControlFlow) {
			W3445 := __e.Get(1)
			_ = W3445
			tmp10815 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3445)

			if True == tmp10815 {
				tmp10774 := MakeNative(func(__e *ControlFlow) {
					W3447 := __e.Get(1)
					_ = W3447
					tmp10790 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3447)

					if True == tmp10790 {
						tmp10775 := MakeNative(func(__e *ControlFlow) {
							W3455 := __e.Get(1)
							_ = W3455
							tmp10777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3455)

							if True == tmp10777 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W3455)
								return
							}

						}, 1)

						tmp10778 := MakeNative(func(__e *ControlFlow) {
							W3456 := __e.Get(1)
							_ = W3456
							tmp10786 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3456)

							if True == tmp10786 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10779 := MakeNative(func(__e *ControlFlow) {
									W3457 := __e.Get(1)
									_ = W3457
									tmp10780 := MakeNative(func(__e *ControlFlow) {
										W3458 := __e.Get(1)
										_ = W3458
										tmp10781 := PrimCons(W3457, Nil)

										tmp10782 := PrimCons(Nil, tmp10781)

										__e.TailApply(PrimFunc(symshen_4comb), W3458, tmp10782)
										return

									}, 1)

									tmp10783 := Call(__e, PrimFunc(symshen_4in_1_6), W3456)

									__e.TailApply(tmp10780, tmp10783)
									return

								}, 1)

								tmp10784 := Call(__e, PrimFunc(symshen_4_5_1out), W3456)

								__e.TailApply(tmp10779, tmp10784)
								return

							}

						}, 1)

						tmp10787 := Call(__e, PrimFunc(symshen_4_5formula_6), V3444)

						tmp10788 := Call(__e, tmp10778, tmp10787)

						__e.TailApply(tmp10775, tmp10788)
						return

					} else {
						__e.Return(W3447)
						return
					}

				}, 1)

				tmp10791 := MakeNative(func(__e *ControlFlow) {
					W3448 := __e.Get(1)
					_ = W3448
					tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3448)

					if True == tmp10811 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10792 := MakeNative(func(__e *ControlFlow) {
							W3449 := __e.Get(1)
							_ = W3449
							tmp10793 := MakeNative(func(__e *ControlFlow) {
								W3450 := __e.Get(1)
								_ = W3450
								tmp10807 := Call(__e, PrimFunc(symshen_4hds_a_2), W3450, sym_6_6)

								if True == tmp10807 {
									tmp10794 := MakeNative(func(__e *ControlFlow) {
										W3451 := __e.Get(1)
										_ = W3451
										tmp10795 := MakeNative(func(__e *ControlFlow) {
											W3452 := __e.Get(1)
											_ = W3452
											tmp10803 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3452)

											if True == tmp10803 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp10796 := MakeNative(func(__e *ControlFlow) {
													W3453 := __e.Get(1)
													_ = W3453
													tmp10797 := MakeNative(func(__e *ControlFlow) {
														W3454 := __e.Get(1)
														_ = W3454
														tmp10798 := PrimCons(W3453, Nil)

														tmp10799 := PrimCons(W3449, tmp10798)

														__e.TailApply(PrimFunc(symshen_4comb), W3454, tmp10799)
														return

													}, 1)

													tmp10800 := Call(__e, PrimFunc(symshen_4in_1_6), W3452)

													__e.TailApply(tmp10797, tmp10800)
													return

												}, 1)

												tmp10801 := Call(__e, PrimFunc(symshen_4_5_1out), W3452)

												__e.TailApply(tmp10796, tmp10801)
												return

											}

										}, 1)

										tmp10804 := Call(__e, PrimFunc(symshen_4_5formula_6), W3451)

										__e.TailApply(tmp10795, tmp10804)
										return

									}, 1)

									tmp10805 := Call(__e, PrimFunc(symtail), W3450)

									__e.TailApply(tmp10794, tmp10805)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp10808 := Call(__e, PrimFunc(symshen_4in_1_6), W3448)

							__e.TailApply(tmp10793, tmp10808)
							return

						}, 1)

						tmp10809 := Call(__e, PrimFunc(symshen_4_5_1out), W3448)

						__e.TailApply(tmp10792, tmp10809)
						return

					}

				}, 1)

				tmp10812 := Call(__e, PrimFunc(symshen_4_5ass_6), V3444)

				tmp10813 := Call(__e, tmp10791, tmp10812)

				__e.TailApply(tmp10774, tmp10813)
				return

			} else {
				__e.Return(W3445)
				return
			}

		}, 1)

		tmp10821 := Call(__e, PrimFunc(symshen_4hds_a_2), V3444, sym_b)

		var ifres10816 Obj

		if True == tmp10821 {
			tmp10817 := MakeNative(func(__e *ControlFlow) {
				W3446 := __e.Get(1)
				_ = W3446
				__e.TailApply(PrimFunc(symshen_4comb), W3446, sym_b)
				return
			}, 1)

			tmp10818 := Call(__e, PrimFunc(symtail), V3444)

			tmp10819 := Call(__e, tmp10817, tmp10818)

			ifres10816 = tmp10819

		} else {
			tmp10820 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10816 = tmp10820

		}

		__e.TailApply(tmp10773, ifres10816)
		return

	}, 1)

	tmp10822 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp10772)

	_ = tmp10822

	tmp10823 := MakeNative(func(__e *ControlFlow) {
		V3459 := __e.Get(1)
		_ = V3459
		tmp10824 := MakeNative(func(__e *ControlFlow) {
			W3460 := __e.Get(1)
			_ = W3460
			tmp10849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3460)

			if True == tmp10849 {
				tmp10825 := MakeNative(func(__e *ControlFlow) {
					W3469 := __e.Get(1)
					_ = W3469
					tmp10837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3469)

					if True == tmp10837 {
						tmp10826 := MakeNative(func(__e *ControlFlow) {
							W3473 := __e.Get(1)
							_ = W3473
							tmp10828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3473)

							if True == tmp10828 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W3473)
								return
							}

						}, 1)

						tmp10829 := MakeNative(func(__e *ControlFlow) {
							W3474 := __e.Get(1)
							_ = W3474
							tmp10833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3474)

							if True == tmp10833 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10830 := MakeNative(func(__e *ControlFlow) {
									W3475 := __e.Get(1)
									_ = W3475
									__e.TailApply(PrimFunc(symshen_4comb), W3475, Nil)
									return
								}, 1)

								tmp10831 := Call(__e, PrimFunc(symshen_4in_1_6), W3474)

								__e.TailApply(tmp10830, tmp10831)
								return

							}

						}, 1)

						tmp10834 := Call(__e, PrimFunc(sym_5e_6), V3459)

						tmp10835 := Call(__e, tmp10829, tmp10834)

						__e.TailApply(tmp10826, tmp10835)
						return

					} else {
						__e.Return(W3469)
						return
					}

				}, 1)

				tmp10838 := MakeNative(func(__e *ControlFlow) {
					W3470 := __e.Get(1)
					_ = W3470
					tmp10845 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3470)

					if True == tmp10845 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10839 := MakeNative(func(__e *ControlFlow) {
							W3471 := __e.Get(1)
							_ = W3471
							tmp10840 := MakeNative(func(__e *ControlFlow) {
								W3472 := __e.Get(1)
								_ = W3472
								tmp10841 := PrimCons(W3471, Nil)

								__e.TailApply(PrimFunc(symshen_4comb), W3472, tmp10841)
								return

							}, 1)

							tmp10842 := Call(__e, PrimFunc(symshen_4in_1_6), W3470)

							__e.TailApply(tmp10840, tmp10842)
							return

						}, 1)

						tmp10843 := Call(__e, PrimFunc(symshen_4_5_1out), W3470)

						__e.TailApply(tmp10839, tmp10843)
						return

					}

				}, 1)

				tmp10846 := Call(__e, PrimFunc(symshen_4_5formula_6), V3459)

				tmp10847 := Call(__e, tmp10838, tmp10846)

				__e.TailApply(tmp10825, tmp10847)
				return

			} else {
				__e.Return(W3460)
				return
			}

		}, 1)

		tmp10850 := MakeNative(func(__e *ControlFlow) {
			W3461 := __e.Get(1)
			_ = W3461
			tmp10871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3461)

			if True == tmp10871 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10851 := MakeNative(func(__e *ControlFlow) {
					W3462 := __e.Get(1)
					_ = W3462
					tmp10852 := MakeNative(func(__e *ControlFlow) {
						W3463 := __e.Get(1)
						_ = W3463
						tmp10853 := MakeNative(func(__e *ControlFlow) {
							W3464 := __e.Get(1)
							_ = W3464
							tmp10866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3464)

							if True == tmp10866 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10854 := MakeNative(func(__e *ControlFlow) {
									W3465 := __e.Get(1)
									_ = W3465
									tmp10855 := MakeNative(func(__e *ControlFlow) {
										W3466 := __e.Get(1)
										_ = W3466
										tmp10862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3466)

										if True == tmp10862 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp10856 := MakeNative(func(__e *ControlFlow) {
												W3467 := __e.Get(1)
												_ = W3467
												tmp10857 := MakeNative(func(__e *ControlFlow) {
													W3468 := __e.Get(1)
													_ = W3468
													tmp10858 := PrimCons(W3462, W3467)

													__e.TailApply(PrimFunc(symshen_4comb), W3468, tmp10858)
													return

												}, 1)

												tmp10859 := Call(__e, PrimFunc(symshen_4in_1_6), W3466)

												__e.TailApply(tmp10857, tmp10859)
												return

											}, 1)

											tmp10860 := Call(__e, PrimFunc(symshen_4_5_1out), W3466)

											__e.TailApply(tmp10856, tmp10860)
											return

										}

									}, 1)

									tmp10863 := Call(__e, PrimFunc(symshen_4_5ass_6), W3465)

									__e.TailApply(tmp10855, tmp10863)
									return

								}, 1)

								tmp10864 := Call(__e, PrimFunc(symshen_4in_1_6), W3464)

								__e.TailApply(tmp10854, tmp10864)
								return

							}

						}, 1)

						tmp10867 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W3463)

						__e.TailApply(tmp10853, tmp10867)
						return

					}, 1)

					tmp10868 := Call(__e, PrimFunc(symshen_4in_1_6), W3461)

					__e.TailApply(tmp10852, tmp10868)
					return

				}, 1)

				tmp10869 := Call(__e, PrimFunc(symshen_4_5_1out), W3461)

				__e.TailApply(tmp10851, tmp10869)
				return

			}

		}, 1)

		tmp10872 := Call(__e, PrimFunc(symshen_4_5formula_6), V3459)

		tmp10873 := Call(__e, tmp10850, tmp10872)

		__e.TailApply(tmp10824, tmp10873)
		return

	}, 1)

	tmp10874 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp10823)

	_ = tmp10874

	tmp10875 := MakeNative(func(__e *ControlFlow) {
		V3476 := __e.Get(1)
		_ = V3476
		tmp10876 := MakeNative(func(__e *ControlFlow) {
			W3477 := __e.Get(1)
			_ = W3477
			tmp10878 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3477)

			if True == tmp10878 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3477)
				return
			}

		}, 1)

		tmp10889 := PrimIsPair(V3476)

		var ifres10879 Obj

		if True == tmp10889 {
			tmp10880 := MakeNative(func(__e *ControlFlow) {
				W3478 := __e.Get(1)
				_ = W3478
				tmp10881 := MakeNative(func(__e *ControlFlow) {
					W3479 := __e.Get(1)
					_ = W3479
					tmp10883 := PrimIntern(MakeString(","))

					tmp10884 := PrimEqual(W3478, tmp10883)

					if True == tmp10884 {
						__e.TailApply(PrimFunc(symshen_4comb), W3479, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10885 := Call(__e, PrimFunc(symtail), V3476)

				__e.TailApply(tmp10881, tmp10885)
				return

			}, 1)

			tmp10886 := Call(__e, PrimFunc(symhead), V3476)

			tmp10887 := Call(__e, tmp10880, tmp10886)

			ifres10879 = tmp10887

		} else {
			tmp10888 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10879 = tmp10888

		}

		__e.TailApply(tmp10876, ifres10879)
		return

	}, 1)

	tmp10890 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp10875)

	_ = tmp10890

	tmp10891 := MakeNative(func(__e *ControlFlow) {
		V3480 := __e.Get(1)
		_ = V3480
		tmp10892 := MakeNative(func(__e *ControlFlow) {
			W3481 := __e.Get(1)
			_ = W3481
			tmp10906 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3481)

			if True == tmp10906 {
				tmp10893 := MakeNative(func(__e *ControlFlow) {
					W3490 := __e.Get(1)
					_ = W3490
					tmp10895 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3490)

					if True == tmp10895 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3490)
						return
					}

				}, 1)

				tmp10896 := MakeNative(func(__e *ControlFlow) {
					W3491 := __e.Get(1)
					_ = W3491
					tmp10902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3491)

					if True == tmp10902 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10897 := MakeNative(func(__e *ControlFlow) {
							W3492 := __e.Get(1)
							_ = W3492
							tmp10898 := MakeNative(func(__e *ControlFlow) {
								W3493 := __e.Get(1)
								_ = W3493
								__e.TailApply(PrimFunc(symshen_4comb), W3493, W3492)
								return
							}, 1)

							tmp10899 := Call(__e, PrimFunc(symshen_4in_1_6), W3491)

							__e.TailApply(tmp10898, tmp10899)
							return

						}, 1)

						tmp10900 := Call(__e, PrimFunc(symshen_4_5_1out), W3491)

						__e.TailApply(tmp10897, tmp10900)
						return

					}

				}, 1)

				tmp10903 := Call(__e, PrimFunc(symshen_4_5expr_6), V3480)

				tmp10904 := Call(__e, tmp10896, tmp10903)

				__e.TailApply(tmp10893, tmp10904)
				return

			} else {
				__e.Return(W3481)
				return
			}

		}, 1)

		tmp10907 := MakeNative(func(__e *ControlFlow) {
			W3482 := __e.Get(1)
			_ = W3482
			tmp10933 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3482)

			if True == tmp10933 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10908 := MakeNative(func(__e *ControlFlow) {
					W3483 := __e.Get(1)
					_ = W3483
					tmp10909 := MakeNative(func(__e *ControlFlow) {
						W3484 := __e.Get(1)
						_ = W3484
						tmp10910 := MakeNative(func(__e *ControlFlow) {
							W3485 := __e.Get(1)
							_ = W3485
							tmp10928 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3485)

							if True == tmp10928 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10911 := MakeNative(func(__e *ControlFlow) {
									W3486 := __e.Get(1)
									_ = W3486
									tmp10912 := MakeNative(func(__e *ControlFlow) {
										W3487 := __e.Get(1)
										_ = W3487
										tmp10924 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3487)

										if True == tmp10924 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp10913 := MakeNative(func(__e *ControlFlow) {
												W3488 := __e.Get(1)
												_ = W3488
												tmp10914 := MakeNative(func(__e *ControlFlow) {
													W3489 := __e.Get(1)
													_ = W3489
													tmp10915 := Call(__e, PrimFunc(symshen_4curry), W3483)

													tmp10916 := PrimIntern(MakeString(":"))

													tmp10917 := Call(__e, PrimFunc(symshen_4rectify_1type), W3488)

													tmp10918 := PrimCons(tmp10917, Nil)

													tmp10919 := PrimCons(tmp10916, tmp10918)

													tmp10920 := PrimCons(tmp10915, tmp10919)

													__e.TailApply(PrimFunc(symshen_4comb), W3489, tmp10920)
													return

												}, 1)

												tmp10921 := Call(__e, PrimFunc(symshen_4in_1_6), W3487)

												__e.TailApply(tmp10914, tmp10921)
												return

											}, 1)

											tmp10922 := Call(__e, PrimFunc(symshen_4_5_1out), W3487)

											__e.TailApply(tmp10913, tmp10922)
											return

										}

									}, 1)

									tmp10925 := Call(__e, PrimFunc(symshen_4_5type_6), W3486)

									__e.TailApply(tmp10912, tmp10925)
									return

								}, 1)

								tmp10926 := Call(__e, PrimFunc(symshen_4in_1_6), W3485)

								__e.TailApply(tmp10911, tmp10926)
								return

							}

						}, 1)

						tmp10929 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W3484)

						__e.TailApply(tmp10910, tmp10929)
						return

					}, 1)

					tmp10930 := Call(__e, PrimFunc(symshen_4in_1_6), W3482)

					__e.TailApply(tmp10909, tmp10930)
					return

				}, 1)

				tmp10931 := Call(__e, PrimFunc(symshen_4_5_1out), W3482)

				__e.TailApply(tmp10908, tmp10931)
				return

			}

		}, 1)

		tmp10934 := Call(__e, PrimFunc(symshen_4_5expr_6), V3480)

		tmp10935 := Call(__e, tmp10907, tmp10934)

		__e.TailApply(tmp10892, tmp10935)
		return

	}, 1)

	tmp10936 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp10891)

	_ = tmp10936

	tmp10937 := MakeNative(func(__e *ControlFlow) {
		V3494 := __e.Get(1)
		_ = V3494
		tmp10938 := MakeNative(func(__e *ControlFlow) {
			W3495 := __e.Get(1)
			_ = W3495
			tmp10940 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3495)

			if True == tmp10940 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3495)
				return
			}

		}, 1)

		tmp10951 := PrimIsPair(V3494)

		var ifres10941 Obj

		if True == tmp10951 {
			tmp10942 := MakeNative(func(__e *ControlFlow) {
				W3496 := __e.Get(1)
				_ = W3496
				tmp10943 := MakeNative(func(__e *ControlFlow) {
					W3497 := __e.Get(1)
					_ = W3497
					tmp10945 := PrimIntern(MakeString(":"))

					tmp10946 := PrimEqual(W3496, tmp10945)

					if True == tmp10946 {
						__e.TailApply(PrimFunc(symshen_4comb), W3497, symshen_4skip)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp10947 := Call(__e, PrimFunc(symtail), V3494)

				__e.TailApply(tmp10943, tmp10947)
				return

			}, 1)

			tmp10948 := Call(__e, PrimFunc(symhead), V3494)

			tmp10949 := Call(__e, tmp10942, tmp10948)

			ifres10941 = tmp10949

		} else {
			tmp10950 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres10941 = tmp10950

		}

		__e.TailApply(tmp10938, ifres10941)
		return

	}, 1)

	tmp10952 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp10937)

	_ = tmp10952

	tmp10953 := MakeNative(func(__e *ControlFlow) {
		V3498 := __e.Get(1)
		_ = V3498
		tmp10954 := MakeNative(func(__e *ControlFlow) {
			W3499 := __e.Get(1)
			_ = W3499
			tmp10966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3499)

			if True == tmp10966 {
				tmp10955 := MakeNative(func(__e *ControlFlow) {
					W3506 := __e.Get(1)
					_ = W3506
					tmp10957 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3506)

					if True == tmp10957 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W3506)
						return
					}

				}, 1)

				tmp10958 := MakeNative(func(__e *ControlFlow) {
					W3507 := __e.Get(1)
					_ = W3507
					tmp10962 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3507)

					if True == tmp10962 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp10959 := MakeNative(func(__e *ControlFlow) {
							W3508 := __e.Get(1)
							_ = W3508
							__e.TailApply(PrimFunc(symshen_4comb), W3508, Nil)
							return
						}, 1)

						tmp10960 := Call(__e, PrimFunc(symshen_4in_1_6), W3507)

						__e.TailApply(tmp10959, tmp10960)
						return

					}

				}, 1)

				tmp10963 := Call(__e, PrimFunc(sym_5e_6), V3498)

				tmp10964 := Call(__e, tmp10958, tmp10963)

				__e.TailApply(tmp10955, tmp10964)
				return

			} else {
				__e.Return(W3499)
				return
			}

		}, 1)

		tmp10967 := MakeNative(func(__e *ControlFlow) {
			W3500 := __e.Get(1)
			_ = W3500
			tmp10982 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3500)

			if True == tmp10982 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp10968 := MakeNative(func(__e *ControlFlow) {
					W3501 := __e.Get(1)
					_ = W3501
					tmp10969 := MakeNative(func(__e *ControlFlow) {
						W3502 := __e.Get(1)
						_ = W3502
						tmp10970 := MakeNative(func(__e *ControlFlow) {
							W3503 := __e.Get(1)
							_ = W3503
							tmp10977 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3503)

							if True == tmp10977 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp10971 := MakeNative(func(__e *ControlFlow) {
									W3504 := __e.Get(1)
									_ = W3504
									tmp10972 := MakeNative(func(__e *ControlFlow) {
										W3505 := __e.Get(1)
										_ = W3505
										tmp10973 := PrimCons(W3501, W3504)

										__e.TailApply(PrimFunc(symshen_4comb), W3505, tmp10973)
										return

									}, 1)

									tmp10974 := Call(__e, PrimFunc(symshen_4in_1_6), W3503)

									__e.TailApply(tmp10972, tmp10974)
									return

								}, 1)

								tmp10975 := Call(__e, PrimFunc(symshen_4_5_1out), W3503)

								__e.TailApply(tmp10971, tmp10975)
								return

							}

						}, 1)

						tmp10978 := Call(__e, PrimFunc(symshen_4_5sides_6), W3502)

						__e.TailApply(tmp10970, tmp10978)
						return

					}, 1)

					tmp10979 := Call(__e, PrimFunc(symshen_4in_1_6), W3500)

					__e.TailApply(tmp10969, tmp10979)
					return

				}, 1)

				tmp10980 := Call(__e, PrimFunc(symshen_4_5_1out), W3500)

				__e.TailApply(tmp10968, tmp10980)
				return

			}

		}, 1)

		tmp10983 := Call(__e, PrimFunc(symshen_4_5side_6), V3498)

		tmp10984 := Call(__e, tmp10967, tmp10983)

		__e.TailApply(tmp10954, tmp10984)
		return

	}, 1)

	tmp10985 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp10953)

	_ = tmp10985

	tmp10986 := MakeNative(func(__e *ControlFlow) {
		V3509 := __e.Get(1)
		_ = V3509
		tmp10987 := MakeNative(func(__e *ControlFlow) {
			W3510 := __e.Get(1)
			_ = W3510
			tmp11032 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3510)

			if True == tmp11032 {
				tmp10988 := MakeNative(func(__e *ControlFlow) {
					W3514 := __e.Get(1)
					_ = W3514
					tmp11009 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3514)

					if True == tmp11009 {
						tmp10989 := MakeNative(func(__e *ControlFlow) {
							W3520 := __e.Get(1)
							_ = W3520
							tmp10991 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3520)

							if True == tmp10991 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W3520)
								return
							}

						}, 1)

						tmp11007 := Call(__e, PrimFunc(symshen_4hds_a_2), V3509, symctxt)

						var ifres10992 Obj

						if True == tmp11007 {
							tmp10993 := MakeNative(func(__e *ControlFlow) {
								W3521 := __e.Get(1)
								_ = W3521
								tmp11003 := PrimIsPair(W3521)

								if True == tmp11003 {
									tmp10994 := MakeNative(func(__e *ControlFlow) {
										W3522 := __e.Get(1)
										_ = W3522
										tmp10995 := MakeNative(func(__e *ControlFlow) {
											W3523 := __e.Get(1)
											_ = W3523
											tmp10999 := PrimIsVariable(W3522)

											if True == tmp10999 {
												tmp10996 := PrimCons(W3522, Nil)

												tmp10997 := PrimCons(symctxt, tmp10996)

												__e.TailApply(PrimFunc(symshen_4comb), W3523, tmp10997)
												return

											} else {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp11000 := Call(__e, PrimFunc(symtail), W3521)

										__e.TailApply(tmp10995, tmp11000)
										return

									}, 1)

									tmp11001 := Call(__e, PrimFunc(symhead), W3521)

									__e.TailApply(tmp10994, tmp11001)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp11004 := Call(__e, PrimFunc(symtail), V3509)

							tmp11005 := Call(__e, tmp10993, tmp11004)

							ifres10992 = tmp11005

						} else {
							tmp11006 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres10992 = tmp11006

						}

						__e.TailApply(tmp10989, ifres10992)
						return

					} else {
						__e.Return(W3514)
						return
					}

				}, 1)

				tmp11030 := Call(__e, PrimFunc(symshen_4hds_a_2), V3509, symlet)

				var ifres11010 Obj

				if True == tmp11030 {
					tmp11011 := MakeNative(func(__e *ControlFlow) {
						W3515 := __e.Get(1)
						_ = W3515
						tmp11026 := PrimIsPair(W3515)

						if True == tmp11026 {
							tmp11012 := MakeNative(func(__e *ControlFlow) {
								W3516 := __e.Get(1)
								_ = W3516
								tmp11013 := MakeNative(func(__e *ControlFlow) {
									W3517 := __e.Get(1)
									_ = W3517
									tmp11022 := PrimIsPair(W3517)

									if True == tmp11022 {
										tmp11014 := MakeNative(func(__e *ControlFlow) {
											W3518 := __e.Get(1)
											_ = W3518
											tmp11015 := MakeNative(func(__e *ControlFlow) {
												W3519 := __e.Get(1)
												_ = W3519
												tmp11016 := PrimCons(W3518, Nil)

												tmp11017 := PrimCons(W3516, tmp11016)

												tmp11018 := PrimCons(symlet, tmp11017)

												__e.TailApply(PrimFunc(symshen_4comb), W3519, tmp11018)
												return

											}, 1)

											tmp11019 := Call(__e, PrimFunc(symtail), W3517)

											__e.TailApply(tmp11015, tmp11019)
											return

										}, 1)

										tmp11020 := Call(__e, PrimFunc(symhead), W3517)

										__e.TailApply(tmp11014, tmp11020)
										return

									} else {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp11023 := Call(__e, PrimFunc(symtail), W3515)

								__e.TailApply(tmp11013, tmp11023)
								return

							}, 1)

							tmp11024 := Call(__e, PrimFunc(symhead), W3515)

							__e.TailApply(tmp11012, tmp11024)
							return

						} else {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp11027 := Call(__e, PrimFunc(symtail), V3509)

					tmp11028 := Call(__e, tmp11011, tmp11027)

					ifres11010 = tmp11028

				} else {
					tmp11029 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres11010 = tmp11029

				}

				__e.TailApply(tmp10988, ifres11010)
				return

			} else {
				__e.Return(W3510)
				return
			}

		}, 1)

		tmp11046 := Call(__e, PrimFunc(symshen_4hds_a_2), V3509, symif)

		var ifres11033 Obj

		if True == tmp11046 {
			tmp11034 := MakeNative(func(__e *ControlFlow) {
				W3511 := __e.Get(1)
				_ = W3511
				tmp11042 := PrimIsPair(W3511)

				if True == tmp11042 {
					tmp11035 := MakeNative(func(__e *ControlFlow) {
						W3512 := __e.Get(1)
						_ = W3512
						tmp11036 := MakeNative(func(__e *ControlFlow) {
							W3513 := __e.Get(1)
							_ = W3513
							tmp11037 := PrimCons(W3512, Nil)

							tmp11038 := PrimCons(symif, tmp11037)

							__e.TailApply(PrimFunc(symshen_4comb), W3513, tmp11038)
							return

						}, 1)

						tmp11039 := Call(__e, PrimFunc(symtail), W3511)

						__e.TailApply(tmp11036, tmp11039)
						return

					}, 1)

					tmp11040 := Call(__e, PrimFunc(symhead), W3511)

					__e.TailApply(tmp11035, tmp11040)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4parse_1failure))
					return
				}

			}, 1)

			tmp11043 := Call(__e, PrimFunc(symtail), V3509)

			tmp11044 := Call(__e, tmp11034, tmp11043)

			ifres11033 = tmp11044

		} else {
			tmp11045 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres11033 = tmp11045

		}

		__e.TailApply(tmp10987, ifres11033)
		return

	}, 1)

	tmp11047 := Call(__e, ns2_1set, symshen_4_5side_6, tmp10986)

	_ = tmp11047

	tmp11048 := MakeNative(func(__e *ControlFlow) {
		V3530 := __e.Get(1)
		_ = V3530
		V3531 := __e.Get(2)
		_ = V3531
		V3532 := __e.Get(3)
		_ = V3532
		tmp11083 := PrimIsPair(V3532)

		var ifres11070 Obj

		if True == tmp11083 {
			tmp11081 := PrimHead(V3532)

			tmp11082 := PrimEqual(Nil, tmp11081)

			var ifres11072 Obj

			if True == tmp11082 {
				tmp11079 := PrimTail(V3532)

				tmp11080 := PrimIsPair(tmp11079)

				var ifres11074 Obj

				if True == tmp11080 {
					tmp11076 := PrimTail(V3532)

					tmp11077 := PrimTail(tmp11076)

					tmp11078 := PrimEqual(Nil, tmp11077)

					var ifres11075 Obj

					if True == tmp11078 {
						ifres11075 = True

					} else {
						ifres11075 = False

					}

					ifres11074 = ifres11075

				} else {
					ifres11074 = False

				}

				var ifres11073 Obj

				if True == ifres11074 {
					ifres11073 = True

				} else {
					ifres11073 = False

				}

				ifres11072 = ifres11073

			} else {
				ifres11072 = False

			}

			var ifres11071 Obj

			if True == ifres11072 {
				ifres11071 = True

			} else {
				ifres11071 = False

			}

			ifres11070 = ifres11071

		} else {
			ifres11070 = False

		}

		if True == ifres11070 {
			tmp11049 := MakeNative(func(__e *ControlFlow) {
				W3533 := __e.Get(1)
				_ = W3533
				tmp11050 := MakeNative(func(__e *ControlFlow) {
					W3534 := __e.Get(1)
					_ = W3534
					tmp11051 := MakeNative(func(__e *ControlFlow) {
						W3535 := __e.Get(1)
						_ = W3535
						tmp11052 := MakeNative(func(__e *ControlFlow) {
							W3536 := __e.Get(1)
							_ = W3536
							tmp11053 := MakeNative(func(__e *ControlFlow) {
								W3537 := __e.Get(1)
								_ = W3537
								tmp11054 := PrimCons(W3536, Nil)

								__e.Return(PrimCons(W3537, tmp11054))
								return

							}, 1)

							tmp11055 := PrimCons(V3532, Nil)

							tmp11056 := PrimCons(V3531, tmp11055)

							tmp11057 := PrimCons(V3530, tmp11056)

							__e.TailApply(tmp11053, tmp11057)
							return

						}, 1)

						tmp11058 := PrimCons(W3535, Nil)

						tmp11059 := PrimCons(W3534, Nil)

						tmp11060 := PrimCons(tmp11058, tmp11059)

						tmp11061 := PrimCons(V3530, tmp11060)

						__e.TailApply(tmp11052, tmp11061)
						return

					}, 1)

					tmp11062 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3531)

					tmp11063 := PrimCons(W3533, Nil)

					tmp11064 := PrimCons(tmp11062, tmp11063)

					__e.TailApply(tmp11051, tmp11064)
					return

				}, 1)

				tmp11065 := PrimTail(V3532)

				tmp11066 := PrimCons(W3533, Nil)

				tmp11067 := PrimCons(tmp11065, tmp11066)

				__e.TailApply(tmp11050, tmp11067)
				return

			}, 1)

			tmp11068 := Call(__e, PrimFunc(symgensym), symP)

			__e.TailApply(tmp11049, tmp11068)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
			return
		}

	}, 3)

	tmp11084 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp11048)

	_ = tmp11084

	tmp11085 := MakeNative(func(__e *ControlFlow) {
		V3540 := __e.Get(1)
		_ = V3540
		tmp11114 := PrimEqual(Nil, V3540)

		if True == tmp11114 {
			__e.Return(Nil)
			return
		} else {
			tmp11112 := PrimIsPair(V3540)

			var ifres11092 Obj

			if True == tmp11112 {
				tmp11110 := PrimHead(V3540)

				tmp11111 := PrimIsPair(tmp11110)

				var ifres11094 Obj

				if True == tmp11111 {
					tmp11107 := PrimHead(V3540)

					tmp11108 := PrimHead(tmp11107)

					tmp11109 := PrimEqual(Nil, tmp11108)

					var ifres11096 Obj

					if True == tmp11109 {
						tmp11104 := PrimHead(V3540)

						tmp11105 := PrimTail(tmp11104)

						tmp11106 := PrimIsPair(tmp11105)

						var ifres11098 Obj

						if True == tmp11106 {
							tmp11100 := PrimHead(V3540)

							tmp11101 := PrimTail(tmp11100)

							tmp11102 := PrimTail(tmp11101)

							tmp11103 := PrimEqual(Nil, tmp11102)

							var ifres11099 Obj

							if True == tmp11103 {
								ifres11099 = True

							} else {
								ifres11099 = False

							}

							ifres11098 = ifres11099

						} else {
							ifres11098 = False

						}

						var ifres11097 Obj

						if True == ifres11098 {
							ifres11097 = True

						} else {
							ifres11097 = False

						}

						ifres11096 = ifres11097

					} else {
						ifres11096 = False

					}

					var ifres11095 Obj

					if True == ifres11096 {
						ifres11095 = True

					} else {
						ifres11095 = False

					}

					ifres11094 = ifres11095

				} else {
					ifres11094 = False

				}

				var ifres11093 Obj

				if True == ifres11094 {
					ifres11093 = True

				} else {
					ifres11093 = False

				}

				ifres11092 = ifres11093

			} else {
				ifres11092 = False

			}

			if True == ifres11092 {
				tmp11086 := PrimHead(V3540)

				tmp11087 := PrimTail(tmp11086)

				tmp11088 := PrimHead(tmp11087)

				tmp11089 := PrimTail(V3540)

				tmp11090 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp11089)

				__e.Return(PrimCons(tmp11088, tmp11090))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
				return
			}

		}

	}, 1)

	tmp11115 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp11085)

	_ = tmp11115

	tmp11116 := MakeNative(func(__e *ControlFlow) {
		V3541 := __e.Get(1)
		_ = V3541
		tmp11117 := MakeNative(func(__e *ControlFlow) {
			W3542 := __e.Get(1)
			_ = W3542
			tmp11119 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3542)

			if True == tmp11119 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3542)
				return
			}

		}, 1)

		tmp11131 := PrimIsPair(V3541)

		var ifres11120 Obj

		if True == tmp11131 {
			tmp11121 := MakeNative(func(__e *ControlFlow) {
				W3543 := __e.Get(1)
				_ = W3543
				tmp11122 := MakeNative(func(__e *ControlFlow) {
					W3544 := __e.Get(1)
					_ = W3544
					tmp11125 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W3543)

					tmp11126 := PrimNot(tmp11125)

					if True == tmp11126 {
						tmp11123 := Call(__e, PrimFunc(symmacroexpand), W3543)

						__e.TailApply(PrimFunc(symshen_4comb), W3544, tmp11123)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp11127 := Call(__e, PrimFunc(symtail), V3541)

				__e.TailApply(tmp11122, tmp11127)
				return

			}, 1)

			tmp11128 := Call(__e, PrimFunc(symhead), V3541)

			tmp11129 := Call(__e, tmp11121, tmp11128)

			ifres11120 = tmp11129

		} else {
			tmp11130 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres11120 = tmp11130

		}

		__e.TailApply(tmp11117, ifres11120)
		return

	}, 1)

	tmp11132 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp11116)

	_ = tmp11132

	tmp11133 := MakeNative(func(__e *ControlFlow) {
		V3545 := __e.Get(1)
		_ = V3545
		tmp11140 := PrimIntern(MakeString(";"))

		tmp11141 := PrimIntern(MakeString(","))

		tmp11142 := PrimIntern(MakeString(":"))

		tmp11143 := PrimCons(sym_5_1_1, Nil)

		tmp11144 := PrimCons(tmp11142, tmp11143)

		tmp11145 := PrimCons(tmp11141, tmp11144)

		tmp11146 := PrimCons(tmp11140, tmp11145)

		tmp11147 := PrimCons(sym_6_6, tmp11146)

		tmp11148 := Call(__e, PrimFunc(symelement_2), V3545, tmp11147)

		if True == tmp11148 {
			__e.Return(True)
			return
		} else {
			tmp11138 := Call(__e, PrimFunc(symshen_4sng_2), V3545)

			var ifres11135 Obj

			if True == tmp11138 {
				ifres11135 = True

			} else {
				tmp11137 := Call(__e, PrimFunc(symshen_4dbl_2), V3545)

				var ifres11136 Obj

				if True == tmp11137 {
					ifres11136 = True

				} else {
					ifres11136 = False

				}

				ifres11135 = ifres11136

			}

			if True == ifres11135 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp11149 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp11133)

	_ = tmp11149

	tmp11150 := MakeNative(func(__e *ControlFlow) {
		V3546 := __e.Get(1)
		_ = V3546
		tmp11151 := MakeNative(func(__e *ControlFlow) {
			W3547 := __e.Get(1)
			_ = W3547
			tmp11153 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3547)

			if True == tmp11153 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3547)
				return
			}

		}, 1)

		tmp11154 := MakeNative(func(__e *ControlFlow) {
			W3548 := __e.Get(1)
			_ = W3548
			tmp11160 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3548)

			if True == tmp11160 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp11155 := MakeNative(func(__e *ControlFlow) {
					W3549 := __e.Get(1)
					_ = W3549
					tmp11156 := MakeNative(func(__e *ControlFlow) {
						W3550 := __e.Get(1)
						_ = W3550
						__e.TailApply(PrimFunc(symshen_4comb), W3550, W3549)
						return
					}, 1)

					tmp11157 := Call(__e, PrimFunc(symshen_4in_1_6), W3548)

					__e.TailApply(tmp11156, tmp11157)
					return

				}, 1)

				tmp11158 := Call(__e, PrimFunc(symshen_4_5_1out), W3548)

				__e.TailApply(tmp11155, tmp11158)
				return

			}

		}, 1)

		tmp11161 := Call(__e, PrimFunc(symshen_4_5expr_6), V3546)

		tmp11162 := Call(__e, tmp11154, tmp11161)

		__e.TailApply(tmp11151, tmp11162)
		return

	}, 1)

	tmp11163 := Call(__e, ns2_1set, symshen_4_5type_6, tmp11150)

	_ = tmp11163

	tmp11164 := MakeNative(func(__e *ControlFlow) {
		V3551 := __e.Get(1)
		_ = V3551
		tmp11165 := MakeNative(func(__e *ControlFlow) {
			W3552 := __e.Get(1)
			_ = W3552
			tmp11167 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3552)

			if True == tmp11167 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3552)
				return
			}

		}, 1)

		tmp11177 := PrimIsPair(V3551)

		var ifres11168 Obj

		if True == tmp11177 {
			tmp11169 := MakeNative(func(__e *ControlFlow) {
				W3553 := __e.Get(1)
				_ = W3553
				tmp11170 := MakeNative(func(__e *ControlFlow) {
					W3554 := __e.Get(1)
					_ = W3554
					tmp11172 := Call(__e, PrimFunc(symshen_4dbl_2), W3553)

					if True == tmp11172 {
						__e.TailApply(PrimFunc(symshen_4comb), W3554, W3553)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp11173 := Call(__e, PrimFunc(symtail), V3551)

				__e.TailApply(tmp11170, tmp11173)
				return

			}, 1)

			tmp11174 := Call(__e, PrimFunc(symhead), V3551)

			tmp11175 := Call(__e, tmp11169, tmp11174)

			ifres11168 = tmp11175

		} else {
			tmp11176 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres11168 = tmp11176

		}

		__e.TailApply(tmp11165, ifres11168)
		return

	}, 1)

	tmp11178 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp11164)

	_ = tmp11178

	tmp11179 := MakeNative(func(__e *ControlFlow) {
		V3555 := __e.Get(1)
		_ = V3555
		tmp11180 := MakeNative(func(__e *ControlFlow) {
			W3556 := __e.Get(1)
			_ = W3556
			tmp11182 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3556)

			if True == tmp11182 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W3556)
				return
			}

		}, 1)

		tmp11192 := PrimIsPair(V3555)

		var ifres11183 Obj

		if True == tmp11192 {
			tmp11184 := MakeNative(func(__e *ControlFlow) {
				W3557 := __e.Get(1)
				_ = W3557
				tmp11185 := MakeNative(func(__e *ControlFlow) {
					W3558 := __e.Get(1)
					_ = W3558
					tmp11187 := Call(__e, PrimFunc(symshen_4sng_2), W3557)

					if True == tmp11187 {
						__e.TailApply(PrimFunc(symshen_4comb), W3558, W3557)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp11188 := Call(__e, PrimFunc(symtail), V3555)

				__e.TailApply(tmp11185, tmp11188)
				return

			}, 1)

			tmp11189 := Call(__e, PrimFunc(symhead), V3555)

			tmp11190 := Call(__e, tmp11184, tmp11189)

			ifres11183 = tmp11190

		} else {
			tmp11191 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres11183 = tmp11191

		}

		__e.TailApply(tmp11180, ifres11183)
		return

	}, 1)

	tmp11193 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp11179)

	_ = tmp11193

	tmp11194 := MakeNative(func(__e *ControlFlow) {
		V3559 := __e.Get(1)
		_ = V3559
		tmp11199 := PrimIsSymbol(V3559)

		if True == tmp11199 {
			tmp11196 := PrimStr(V3559)

			tmp11197 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp11196)

			if True == tmp11197 {
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

	tmp11200 := Call(__e, ns2_1set, symshen_4sng_2, tmp11194)

	_ = tmp11200

	tmp11201 := MakeNative(func(__e *ControlFlow) {
		V3562 := __e.Get(1)
		_ = V3562
		tmp11210 := PrimEqual(MakeString("___"), V3562)

		if True == tmp11210 {
			__e.Return(True)
			return
		} else {
			tmp11208 := Call(__e, PrimFunc(symshen_4_7string_2), V3562)

			var ifres11204 Obj

			if True == tmp11208 {
				tmp11206 := Call(__e, PrimFunc(symhdstr), V3562)

				tmp11207 := PrimEqual(MakeString("_"), tmp11206)

				var ifres11205 Obj

				if True == tmp11207 {
					ifres11205 = True

				} else {
					ifres11205 = False

				}

				ifres11204 = ifres11205

			} else {
				ifres11204 = False

			}

			if True == ifres11204 {
				tmp11202 := PrimTailString(V3562)

				__e.TailApply(PrimFunc(symshen_4sng_1h_2), tmp11202)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp11211 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp11201)

	_ = tmp11211

	tmp11212 := MakeNative(func(__e *ControlFlow) {
		V3563 := __e.Get(1)
		_ = V3563
		tmp11217 := PrimIsSymbol(V3563)

		if True == tmp11217 {
			tmp11214 := PrimStr(V3563)

			tmp11215 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp11214)

			if True == tmp11215 {
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

	tmp11218 := Call(__e, ns2_1set, symshen_4dbl_2, tmp11212)

	_ = tmp11218

	tmp11219 := MakeNative(func(__e *ControlFlow) {
		V3566 := __e.Get(1)
		_ = V3566
		tmp11228 := PrimEqual(MakeString("==="), V3566)

		if True == tmp11228 {
			__e.Return(True)
			return
		} else {
			tmp11226 := Call(__e, PrimFunc(symshen_4_7string_2), V3566)

			var ifres11222 Obj

			if True == tmp11226 {
				tmp11224 := Call(__e, PrimFunc(symhdstr), V3566)

				tmp11225 := PrimEqual(MakeString("="), tmp11224)

				var ifres11223 Obj

				if True == tmp11225 {
					ifres11223 = True

				} else {
					ifres11223 = False

				}

				ifres11222 = ifres11223

			} else {
				ifres11222 = False

			}

			if True == ifres11222 {
				tmp11220 := PrimTailString(V3566)

				__e.TailApply(PrimFunc(symshen_4dbl_1h_2), tmp11220)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp11229 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp11219)

	_ = tmp11229

	tmp11230 := MakeNative(func(__e *ControlFlow) {
		V3567 := __e.Get(1)
		_ = V3567
		V3568 := __e.Get(2)
		_ = V3568
		tmp11231 := MakeNative(func(__e *ControlFlow) {
			W3569 := __e.Get(1)
			_ = W3569
			tmp11232 := MakeNative(func(__e *ControlFlow) {
				W3571 := __e.Get(1)
				_ = W3571
				__e.TailApply(PrimFunc(symeval), W3571)
				return
			}, 1)

			tmp11233 := PrimCons(V3567, W3569)

			tmp11234 := PrimCons(symdefprolog, tmp11233)

			__e.TailApply(tmp11232, tmp11234)
			return

		}, 1)

		tmp11235 := MakeNative(func(__e *ControlFlow) {
			Z3570 := __e.Get(1)
			_ = Z3570
			__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3570)
			return
		}, 1)

		tmp11236 := Call(__e, PrimFunc(symmapcan), tmp11235, V3568)

		__e.TailApply(tmp11231, tmp11236)
		return

	}, 2)

	tmp11237 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp11230)

	_ = tmp11237

	tmp11238 := MakeNative(func(__e *ControlFlow) {
		V3572 := __e.Get(1)
		_ = V3572
		tmp11299 := PrimIsPair(V3572)

		var ifres11263 Obj

		if True == tmp11299 {
			tmp11297 := PrimTail(V3572)

			tmp11298 := PrimIsPair(tmp11297)

			var ifres11265 Obj

			if True == tmp11298 {
				tmp11294 := PrimTail(V3572)

				tmp11295 := PrimTail(tmp11294)

				tmp11296 := PrimIsPair(tmp11295)

				var ifres11267 Obj

				if True == tmp11296 {
					tmp11290 := PrimTail(V3572)

					tmp11291 := PrimTail(tmp11290)

					tmp11292 := PrimHead(tmp11291)

					tmp11293 := PrimIsPair(tmp11292)

					var ifres11269 Obj

					if True == tmp11293 {
						tmp11285 := PrimTail(V3572)

						tmp11286 := PrimTail(tmp11285)

						tmp11287 := PrimHead(tmp11286)

						tmp11288 := PrimTail(tmp11287)

						tmp11289 := PrimIsPair(tmp11288)

						var ifres11271 Obj

						if True == tmp11289 {
							tmp11279 := PrimTail(V3572)

							tmp11280 := PrimTail(tmp11279)

							tmp11281 := PrimHead(tmp11280)

							tmp11282 := PrimTail(tmp11281)

							tmp11283 := PrimTail(tmp11282)

							tmp11284 := PrimEqual(Nil, tmp11283)

							var ifres11273 Obj

							if True == tmp11284 {
								tmp11275 := PrimTail(V3572)

								tmp11276 := PrimTail(tmp11275)

								tmp11277 := PrimTail(tmp11276)

								tmp11278 := PrimEqual(Nil, tmp11277)

								var ifres11274 Obj

								if True == tmp11278 {
									ifres11274 = True

								} else {
									ifres11274 = False

								}

								ifres11273 = ifres11274

							} else {
								ifres11273 = False

							}

							var ifres11272 Obj

							if True == ifres11273 {
								ifres11272 = True

							} else {
								ifres11272 = False

							}

							ifres11271 = ifres11272

						} else {
							ifres11271 = False

						}

						var ifres11270 Obj

						if True == ifres11271 {
							ifres11270 = True

						} else {
							ifres11270 = False

						}

						ifres11269 = ifres11270

					} else {
						ifres11269 = False

					}

					var ifres11268 Obj

					if True == ifres11269 {
						ifres11268 = True

					} else {
						ifres11268 = False

					}

					ifres11267 = ifres11268

				} else {
					ifres11267 = False

				}

				var ifres11266 Obj

				if True == ifres11267 {
					ifres11266 = True

				} else {
					ifres11266 = False

				}

				ifres11265 = ifres11266

			} else {
				ifres11265 = False

			}

			var ifres11264 Obj

			if True == ifres11265 {
				ifres11264 = True

			} else {
				ifres11264 = False

			}

			ifres11263 = ifres11264

		} else {
			ifres11263 = False

		}

		if True == ifres11263 {
			tmp11239 := MakeNative(func(__e *ControlFlow) {
				W3573 := __e.Get(1)
				_ = W3573
				tmp11240 := PrimTail(V3572)

				tmp11241 := PrimTail(tmp11240)

				tmp11242 := PrimHead(tmp11241)

				tmp11243 := PrimTail(tmp11242)

				tmp11244 := PrimHead(tmp11243)

				tmp11245 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp11244)

				tmp11246 := PrimCons(sym_5_1_1, Nil)

				tmp11247 := PrimHead(V3572)

				tmp11248 := PrimTail(V3572)

				tmp11249 := PrimHead(tmp11248)

				tmp11250 := PrimTail(V3572)

				tmp11251 := PrimTail(tmp11250)

				tmp11252 := PrimHead(tmp11251)

				tmp11253 := PrimHead(tmp11252)

				tmp11254 := Call(__e, PrimFunc(symshen_4rule_1_6body), W3573, symAssumptions, tmp11247, tmp11249, tmp11253)

				tmp11255 := Call(__e, PrimFunc(symappend), tmp11246, tmp11254)

				__e.TailApply(PrimFunc(symappend), tmp11245, tmp11255)
				return

			}, 1)

			tmp11256 := PrimTail(V3572)

			tmp11257 := PrimTail(tmp11256)

			tmp11258 := PrimHead(tmp11257)

			tmp11259 := PrimTail(tmp11258)

			tmp11260 := PrimHead(tmp11259)

			tmp11261 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11260)

			__e.TailApply(tmp11239, tmp11261)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6clause)
			return
		}

	}, 1)

	tmp11300 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp11238)

	_ = tmp11300

	tmp11301 := MakeNative(func(__e *ControlFlow) {
		V3574 := __e.Get(1)
		_ = V3574
		tmp11302 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3574)

		tmp11303 := PrimCons(symAssumptions, Nil)

		__e.Return(PrimCons(tmp11302, tmp11303))
		return

	}, 1)

	tmp11304 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp11301)

	_ = tmp11304

	tmp11305 := MakeNative(func(__e *ControlFlow) {
		V3575 := __e.Get(1)
		_ = V3575
		tmp11306 := PrimCons(V3575, Nil)

		__e.Return(PrimCons(symshen_4_8ch, tmp11306))
		return

	}, 1)

	tmp11307 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp11305)

	_ = tmp11307

	tmp11308 := MakeNative(func(__e *ControlFlow) {
		V3576 := __e.Get(1)
		_ = V3576
		tmp11309 := PrimCons(V3576, Nil)

		__e.Return(PrimCons(symshen_4_8c, tmp11309))
		return

	}, 1)

	tmp11310 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp11308)

	_ = tmp11310

	tmp11311 := MakeNative(func(__e *ControlFlow) {
		V3577 := __e.Get(1)
		_ = V3577
		V3578 := __e.Get(2)
		_ = V3578
		V3579 := __e.Get(3)
		_ = V3579
		V3580 := __e.Get(4)
		_ = V3580
		V3581 := __e.Get(5)
		_ = V3581
		tmp11346 := PrimEqual(Nil, V3581)

		if True == tmp11346 {
			__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3577, V3578, V3579, V3580)
			return
		} else {
			tmp11344 := PrimEqual(Nil, V3580)

			var ifres11337 Obj

			if True == tmp11344 {
				tmp11343 := PrimIsPair(V3581)

				var ifres11339 Obj

				if True == tmp11343 {
					tmp11341 := PrimTail(V3581)

					tmp11342 := PrimEqual(Nil, tmp11341)

					var ifres11340 Obj

					if True == tmp11342 {
						ifres11340 = True

					} else {
						ifres11340 = False

					}

					ifres11339 = ifres11340

				} else {
					ifres11339 = False

				}

				var ifres11338 Obj

				if True == ifres11339 {
					ifres11338 = True

				} else {
					ifres11338 = False

				}

				ifres11337 = ifres11338

			} else {
				ifres11337 = False

			}

			if True == ifres11337 {
				tmp11312 := MakeNative(func(__e *ControlFlow) {
					W3582 := __e.Get(1)
					_ = W3582
					tmp11313 := MakeNative(func(__e *ControlFlow) {
						W3583 := __e.Get(1)
						_ = W3583
						tmp11314 := PrimHead(V3581)

						tmp11315 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp11314, V3578, W3583, W3582)

						tmp11316 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3577, V3578, V3579, Nil)

						__e.Return(PrimCons(tmp11315, tmp11316))
						return

					}, 1)

					tmp11317 := PrimHead(V3581)

					tmp11318 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3577, tmp11317)

					__e.TailApply(tmp11313, tmp11318)
					return

				}, 1)

				tmp11319 := PrimHead(V3581)

				tmp11320 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp11319, V3577)

				__e.TailApply(tmp11312, tmp11320)
				return

			} else {
				tmp11335 := PrimIsPair(V3581)

				if True == tmp11335 {
					tmp11321 := MakeNative(func(__e *ControlFlow) {
						W3584 := __e.Get(1)
						_ = W3584
						tmp11322 := MakeNative(func(__e *ControlFlow) {
							W3585 := __e.Get(1)
							_ = W3585
							tmp11323 := MakeNative(func(__e *ControlFlow) {
								W3586 := __e.Get(1)
								_ = W3586
								tmp11324 := PrimHead(V3581)

								tmp11325 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp11324, V3578, W3586, W3585, W3584)

								tmp11326 := Call(__e, PrimFunc(symappend), V3577, W3585)

								tmp11327 := PrimTail(V3581)

								tmp11328 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp11326, W3584, V3579, V3580, tmp11327)

								__e.Return(PrimCons(tmp11325, tmp11328))
								return

							}, 1)

							tmp11329 := PrimHead(V3581)

							tmp11330 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3577, tmp11329)

							__e.TailApply(tmp11323, tmp11330)
							return

						}, 1)

						tmp11331 := PrimHead(V3581)

						tmp11332 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp11331, V3577)

						__e.TailApply(tmp11322, tmp11332)
						return

					}, 1)

					tmp11333 := Call(__e, PrimFunc(symgensym), symNewAssumptions)

					__e.TailApply(tmp11321, tmp11333)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6body)
					return
				}

			}

		}

	}, 5)

	tmp11347 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp11311)

	_ = tmp11347

	tmp11348 := MakeNative(func(__e *ControlFlow) {
		V3587 := __e.Get(1)
		_ = V3587
		V3588 := __e.Get(2)
		_ = V3588
		V3589 := __e.Get(3)
		_ = V3589
		V3590 := __e.Get(4)
		_ = V3590
		tmp11349 := MakeNative(func(__e *ControlFlow) {
			W3591 := __e.Get(1)
			_ = W3591
			tmp11350 := MakeNative(func(__e *ControlFlow) {
				W3592 := __e.Get(1)
				_ = W3592
				tmp11351 := Call(__e, PrimFunc(symappend), V3589, V3590)

				tmp11352 := PrimCons(V3588, tmp11351)

				__e.Return(PrimCons(W3591, tmp11352))
				return

			}, 1)

			tmp11353 := Call(__e, PrimFunc(symshen_4member_1clause), W3591, V3587, V3589, V3590)

			__e.TailApply(tmp11350, tmp11353)
			return

		}, 1)

		tmp11354 := Call(__e, PrimFunc(symgensym), symshen_4member)

		__e.TailApply(tmp11349, tmp11354)
		return

	}, 4)

	tmp11355 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp11348)

	_ = tmp11355

	tmp11356 := MakeNative(func(__e *ControlFlow) {
		V3595 := __e.Get(1)
		_ = V3595
		V3596 := __e.Get(2)
		_ = V3596
		tmp11370 := PrimEqual(Nil, V3595)

		if True == tmp11370 {
			__e.Return(Nil)
			return
		} else {
			tmp11368 := PrimIsPair(V3595)

			var ifres11364 Obj

			if True == tmp11368 {
				tmp11366 := PrimHead(V3595)

				tmp11367 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp11366, V3596)

				var ifres11365 Obj

				if True == tmp11367 {
					ifres11365 = True

				} else {
					ifres11365 = False

				}

				ifres11364 = ifres11365

			} else {
				ifres11364 = False

			}

			if True == ifres11364 {
				tmp11357 := PrimHead(V3595)

				tmp11358 := PrimTail(V3595)

				tmp11359 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp11358, V3596)

				__e.Return(PrimCons(tmp11357, tmp11359))
				return

			} else {
				tmp11362 := PrimIsPair(V3595)

				if True == tmp11362 {
					tmp11360 := PrimTail(V3595)

					__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp11360, V3596)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4remove_1bystanders)
					return
				}

			}

		}

	}, 2)

	tmp11371 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp11356)

	_ = tmp11371

	tmp11372 := MakeNative(func(__e *ControlFlow) {
		V3597 := __e.Get(1)
		_ = V3597
		V3598 := __e.Get(2)
		_ = V3598
		V3599 := __e.Get(3)
		_ = V3599
		V3600 := __e.Get(4)
		_ = V3600
		tmp11373 := MakeNative(func(__e *ControlFlow) {
			W3601 := __e.Get(1)
			_ = W3601
			tmp11374 := MakeNative(func(__e *ControlFlow) {
				W3602 := __e.Get(1)
				_ = W3602
				tmp11375 := MakeNative(func(__e *ControlFlow) {
					W3603 := __e.Get(1)
					_ = W3603
					tmp11376 := MakeNative(func(__e *ControlFlow) {
						W3608 := __e.Get(1)
						_ = W3608
						__e.TailApply(PrimFunc(symeval), W3608)
						return
					}, 1)

					tmp11377 := Call(__e, PrimFunc(symappend), W3602, W3603)

					tmp11378 := PrimCons(V3597, tmp11377)

					tmp11379 := PrimCons(symdefprolog, tmp11378)

					__e.TailApply(tmp11376, tmp11379)
					return

				}, 1)

				tmp11380 := MakeNative(func(__e *ControlFlow) {
					W3604 := __e.Get(1)
					_ = W3604
					tmp11381 := MakeNative(func(__e *ControlFlow) {
						W3605 := __e.Get(1)
						_ = W3605
						tmp11382 := MakeNative(func(__e *ControlFlow) {
							W3606 := __e.Get(1)
							_ = W3606
							tmp11383 := MakeNative(func(__e *ControlFlow) {
								W3607 := __e.Get(1)
								_ = W3607
								tmp11384 := PrimCons(sym_5_1_1, Nil)

								tmp11385 := PrimIntern(MakeString(";"))

								tmp11386 := PrimCons(tmp11385, Nil)

								tmp11387 := Call(__e, PrimFunc(symappend), W3607, tmp11386)

								tmp11388 := Call(__e, PrimFunc(symappend), tmp11384, tmp11387)

								__e.TailApply(PrimFunc(symappend), W3606, tmp11388)
								return

							}, 1)

							tmp11389 := PrimCons(W3604, W3605)

							tmp11390 := PrimCons(V3597, tmp11389)

							tmp11391 := PrimCons(tmp11390, Nil)

							__e.TailApply(tmp11383, tmp11391)
							return

						}, 1)

						tmp11392 := PrimCons(W3604, Nil)

						tmp11393 := PrimCons(sym__, tmp11392)

						tmp11394 := PrimCons(symcons, tmp11393)

						tmp11395 := PrimCons(tmp11394, Nil)

						tmp11396 := PrimCons(sym_1, tmp11395)

						tmp11397 := PrimCons(tmp11396, Nil)

						tmp11398 := Call(__e, PrimFunc(symappend), tmp11397, W3605)

						__e.TailApply(tmp11382, tmp11398)
						return

					}, 1)

					tmp11399 := Call(__e, PrimFunc(symappend), V3599, V3600)

					__e.TailApply(tmp11381, tmp11399)
					return

				}, 1)

				tmp11400 := Call(__e, PrimFunc(symgensym), symHypotheses)

				tmp11401 := Call(__e, tmp11380, tmp11400)

				__e.TailApply(tmp11375, tmp11401)
				return

			}, 1)

			tmp11402 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3598)

			tmp11403 := PrimCons(sym__, Nil)

			tmp11404 := PrimCons(tmp11402, tmp11403)

			tmp11405 := PrimCons(symcons, tmp11404)

			tmp11406 := PrimCons(tmp11405, Nil)

			tmp11407 := PrimCons(sym_1, tmp11406)

			tmp11408 := PrimCons(tmp11407, Nil)

			tmp11409 := PrimCons(sym_5_1_1, Nil)

			tmp11410 := Call(__e, PrimFunc(symshen_4passive_1bind), V3600, W3601)

			tmp11411 := PrimIntern(MakeString(";"))

			tmp11412 := PrimCons(tmp11411, Nil)

			tmp11413 := Call(__e, PrimFunc(symappend), tmp11410, tmp11412)

			tmp11414 := Call(__e, PrimFunc(symappend), tmp11409, tmp11413)

			tmp11415 := Call(__e, PrimFunc(symappend), W3601, tmp11414)

			tmp11416 := Call(__e, PrimFunc(symappend), V3599, tmp11415)

			tmp11417 := Call(__e, PrimFunc(symappend), tmp11408, tmp11416)

			__e.TailApply(tmp11374, tmp11417)
			return

		}, 1)

		tmp11418 := Call(__e, PrimFunc(symlength), V3600)

		tmp11419 := Call(__e, PrimFunc(symshen_4nvars), tmp11418)

		__e.TailApply(tmp11373, tmp11419)
		return

	}, 4)

	tmp11420 := Call(__e, ns2_1set, symshen_4member_1clause, tmp11372)

	_ = tmp11420

	tmp11421 := MakeNative(func(__e *ControlFlow) {
		V3609 := __e.Get(1)
		_ = V3609
		tmp11426 := PrimEqual(MakeNumber(0), V3609)

		if True == tmp11426 {
			__e.Return(Nil)
			return
		} else {
			tmp11422 := Call(__e, PrimFunc(symgensym), symNewV)

			tmp11423 := PrimNumberSubtract(V3609, MakeNumber(1))

			tmp11424 := Call(__e, PrimFunc(symshen_4nvars), tmp11423)

			__e.Return(PrimCons(tmp11422, tmp11424))
			return

		}

	}, 1)

	tmp11427 := Call(__e, ns2_1set, symshen_4nvars, tmp11421)

	_ = tmp11427

	tmp11428 := MakeNative(func(__e *ControlFlow) {
		V3610 := __e.Get(1)
		_ = V3610
		V3611 := __e.Get(2)
		_ = V3611
		tmp11446 := PrimEqual(Nil, V3610)

		var ifres11443 Obj

		if True == tmp11446 {
			tmp11445 := PrimEqual(Nil, V3611)

			var ifres11444 Obj

			if True == tmp11445 {
				ifres11444 = True

			} else {
				ifres11444 = False

			}

			ifres11443 = ifres11444

		} else {
			ifres11443 = False

		}

		if True == ifres11443 {
			__e.Return(Nil)
			return
		} else {
			tmp11441 := PrimIsPair(V3610)

			var ifres11438 Obj

			if True == tmp11441 {
				tmp11440 := PrimIsPair(V3611)

				var ifres11439 Obj

				if True == tmp11440 {
					ifres11439 = True

				} else {
					ifres11439 = False

				}

				ifres11438 = ifres11439

			} else {
				ifres11438 = False

			}

			if True == ifres11438 {
				tmp11429 := PrimHead(V3611)

				tmp11430 := PrimHead(V3610)

				tmp11431 := PrimCons(tmp11430, Nil)

				tmp11432 := PrimCons(tmp11429, tmp11431)

				tmp11433 := PrimCons(symbind, tmp11432)

				tmp11434 := PrimTail(V3610)

				tmp11435 := PrimTail(V3611)

				tmp11436 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp11434, tmp11435)

				__e.Return(PrimCons(tmp11433, tmp11436))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4passive_1bind)
				return
			}

		}

	}, 2)

	tmp11447 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp11428)

	_ = tmp11447

	tmp11448 := MakeNative(func(__e *ControlFlow) {
		V3612 := __e.Get(1)
		_ = V3612
		V3613 := __e.Get(2)
		_ = V3613
		V3614 := __e.Get(3)
		_ = V3614
		V3615 := __e.Get(4)
		_ = V3615
		V3616 := __e.Get(5)
		_ = V3616
		tmp11449 := MakeNative(func(__e *ControlFlow) {
			W3617 := __e.Get(1)
			_ = W3617
			tmp11450 := MakeNative(func(__e *ControlFlow) {
				W3618 := __e.Get(1)
				_ = W3618
				tmp11451 := Call(__e, PrimFunc(symappend), V3614, V3615)

				tmp11452 := PrimCons(V3616, tmp11451)

				tmp11453 := PrimCons(V3613, tmp11452)

				__e.Return(PrimCons(W3617, tmp11453))
				return

			}, 1)

			tmp11454 := Call(__e, PrimFunc(symshen_4consume_1clause), W3617, V3612, V3614, V3615, V3616)

			__e.TailApply(tmp11450, tmp11454)
			return

		}, 1)

		tmp11455 := Call(__e, PrimFunc(symgensym), symshen_4consume)

		__e.TailApply(tmp11449, tmp11455)
		return

	}, 5)

	tmp11456 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp11448)

	_ = tmp11456

	tmp11457 := MakeNative(func(__e *ControlFlow) {
		V3619 := __e.Get(1)
		_ = V3619
		V3620 := __e.Get(2)
		_ = V3620
		V3621 := __e.Get(3)
		_ = V3621
		V3622 := __e.Get(4)
		_ = V3622
		V3623 := __e.Get(5)
		_ = V3623
		tmp11458 := MakeNative(func(__e *ControlFlow) {
			W3624 := __e.Get(1)
			_ = W3624
			tmp11459 := MakeNative(func(__e *ControlFlow) {
				W3625 := __e.Get(1)
				_ = W3625
				tmp11460 := MakeNative(func(__e *ControlFlow) {
					W3626 := __e.Get(1)
					_ = W3626
					tmp11461 := MakeNative(func(__e *ControlFlow) {
						W3627 := __e.Get(1)
						_ = W3627
						tmp11462 := MakeNative(func(__e *ControlFlow) {
							W3633 := __e.Get(1)
							_ = W3633
							__e.TailApply(PrimFunc(symeval), W3633)
							return
						}, 1)

						tmp11463 := Call(__e, PrimFunc(symappend), W3626, W3627)

						tmp11464 := PrimCons(V3619, tmp11463)

						tmp11465 := PrimCons(symdefprolog, tmp11464)

						__e.TailApply(tmp11462, tmp11465)
						return

					}, 1)

					tmp11466 := MakeNative(func(__e *ControlFlow) {
						W3628 := __e.Get(1)
						_ = W3628
						tmp11467 := MakeNative(func(__e *ControlFlow) {
							W3629 := __e.Get(1)
							_ = W3629
							tmp11468 := MakeNative(func(__e *ControlFlow) {
								W3630 := __e.Get(1)
								_ = W3630
								tmp11469 := MakeNative(func(__e *ControlFlow) {
									W3631 := __e.Get(1)
									_ = W3631
									tmp11470 := MakeNative(func(__e *ControlFlow) {
										W3632 := __e.Get(1)
										_ = W3632
										tmp11471 := PrimCons(sym_5_1_1, Nil)

										tmp11472 := PrimIntern(MakeString(";"))

										tmp11473 := PrimCons(tmp11472, Nil)

										tmp11474 := Call(__e, PrimFunc(symappend), W3632, tmp11473)

										tmp11475 := Call(__e, PrimFunc(symappend), tmp11471, tmp11474)

										__e.TailApply(PrimFunc(symappend), W3631, tmp11475)
										return

									}, 1)

									tmp11476 := PrimCons(W3625, Nil)

									tmp11477 := PrimCons(W3630, tmp11476)

									tmp11478 := PrimCons(symbind, tmp11477)

									tmp11479 := PrimCons(V3623, W3629)

									tmp11480 := PrimCons(W3628, tmp11479)

									tmp11481 := PrimCons(V3619, tmp11480)

									tmp11482 := PrimCons(tmp11481, Nil)

									tmp11483 := PrimCons(tmp11478, tmp11482)

									__e.TailApply(tmp11470, tmp11483)
									return

								}, 1)

								tmp11484 := PrimCons(W3628, Nil)

								tmp11485 := PrimCons(W3625, tmp11484)

								tmp11486 := PrimCons(symcons, tmp11485)

								tmp11487 := PrimCons(tmp11486, Nil)

								tmp11488 := PrimCons(sym_1, tmp11487)

								tmp11489 := PrimCons(V3623, Nil)

								tmp11490 := PrimCons(W3630, tmp11489)

								tmp11491 := PrimCons(symcons, tmp11490)

								tmp11492 := PrimCons(tmp11491, W3629)

								tmp11493 := PrimCons(tmp11488, tmp11492)

								__e.TailApply(tmp11469, tmp11493)
								return

							}, 1)

							tmp11494 := Call(__e, PrimFunc(symgensym), symAssumptions)

							__e.TailApply(tmp11468, tmp11494)
							return

						}, 1)

						tmp11495 := Call(__e, PrimFunc(symappend), V3621, V3622)

						__e.TailApply(tmp11467, tmp11495)
						return

					}, 1)

					tmp11496 := Call(__e, PrimFunc(symgensym), symHypotheses)

					tmp11497 := Call(__e, tmp11466, tmp11496)

					__e.TailApply(tmp11461, tmp11497)
					return

				}, 1)

				tmp11498 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3620)

				tmp11499 := PrimCons(W3625, Nil)

				tmp11500 := PrimCons(tmp11498, tmp11499)

				tmp11501 := PrimCons(symcons, tmp11500)

				tmp11502 := PrimCons(tmp11501, Nil)

				tmp11503 := PrimCons(sym_1, tmp11502)

				tmp11504 := PrimCons(sym_5_1_1, Nil)

				tmp11505 := Call(__e, PrimFunc(symshen_4passive_1bind), V3622, W3624)

				tmp11506 := PrimCons(W3625, Nil)

				tmp11507 := PrimCons(V3623, tmp11506)

				tmp11508 := PrimCons(symbind, tmp11507)

				tmp11509 := PrimIntern(MakeString(";"))

				tmp11510 := PrimCons(tmp11509, Nil)

				tmp11511 := PrimCons(tmp11508, tmp11510)

				tmp11512 := Call(__e, PrimFunc(symappend), tmp11505, tmp11511)

				tmp11513 := Call(__e, PrimFunc(symappend), tmp11504, tmp11512)

				tmp11514 := Call(__e, PrimFunc(symappend), W3624, tmp11513)

				tmp11515 := Call(__e, PrimFunc(symappend), V3621, tmp11514)

				tmp11516 := PrimCons(V3623, tmp11515)

				tmp11517 := PrimCons(tmp11503, tmp11516)

				__e.TailApply(tmp11460, tmp11517)
				return

			}, 1)

			tmp11518 := Call(__e, PrimFunc(symgensym), symAssumption)

			__e.TailApply(tmp11459, tmp11518)
			return

		}, 1)

		tmp11519 := Call(__e, PrimFunc(symlength), V3622)

		tmp11520 := Call(__e, PrimFunc(symshen_4nvars), tmp11519)

		__e.TailApply(tmp11458, tmp11520)
		return

	}, 5)

	tmp11521 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp11457)

	_ = tmp11521

	tmp11522 := MakeNative(func(__e *ControlFlow) {
		V3634 := __e.Get(1)
		_ = V3634
		V3635 := __e.Get(2)
		_ = V3635
		tmp11523 := Call(__e, PrimFunc(symshen_4extract_1vars), V3634)

		__e.TailApply(PrimFunc(symdifference), tmp11523, V3635)
		return

	}, 2)

	tmp11524 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp11522)

	_ = tmp11524

	tmp11525 := MakeNative(func(__e *ControlFlow) {
		V3638 := __e.Get(1)
		_ = V3638
		V3639 := __e.Get(2)
		_ = V3639
		V3640 := __e.Get(3)
		_ = V3640
		V3641 := __e.Get(4)
		_ = V3641
		V3642 := __e.Get(5)
		_ = V3642
		tmp11653 := PrimEqual(Nil, V3641)

		if True == tmp11653 {
			__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3638, V3640, V3642)
			return
		} else {
			tmp11651 := PrimIsPair(V3641)

			var ifres11631 Obj

			if True == tmp11651 {
				tmp11649 := PrimHead(V3641)

				tmp11650 := PrimIsPair(tmp11649)

				var ifres11633 Obj

				if True == tmp11650 {
					tmp11646 := PrimHead(V3641)

					tmp11647 := PrimHead(tmp11646)

					tmp11648 := PrimEqual(symif, tmp11647)

					var ifres11635 Obj

					if True == tmp11648 {
						tmp11643 := PrimHead(V3641)

						tmp11644 := PrimTail(tmp11643)

						tmp11645 := PrimIsPair(tmp11644)

						var ifres11637 Obj

						if True == tmp11645 {
							tmp11639 := PrimHead(V3641)

							tmp11640 := PrimTail(tmp11639)

							tmp11641 := PrimTail(tmp11640)

							tmp11642 := PrimEqual(Nil, tmp11641)

							var ifres11638 Obj

							if True == tmp11642 {
								ifres11638 = True

							} else {
								ifres11638 = False

							}

							ifres11637 = ifres11638

						} else {
							ifres11637 = False

						}

						var ifres11636 Obj

						if True == ifres11637 {
							ifres11636 = True

						} else {
							ifres11636 = False

						}

						ifres11635 = ifres11636

					} else {
						ifres11635 = False

					}

					var ifres11634 Obj

					if True == ifres11635 {
						ifres11634 = True

					} else {
						ifres11634 = False

					}

					ifres11633 = ifres11634

				} else {
					ifres11633 = False

				}

				var ifres11632 Obj

				if True == ifres11633 {
					ifres11632 = True

				} else {
					ifres11632 = False

				}

				ifres11631 = ifres11632

			} else {
				ifres11631 = False

			}

			if True == ifres11631 {
				tmp11526 := PrimHead(V3641)

				tmp11527 := PrimTail(tmp11526)

				tmp11528 := PrimCons(symwhen, tmp11527)

				tmp11529 := PrimTail(V3641)

				tmp11530 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3638, V3639, V3640, tmp11529, V3642)

				__e.Return(PrimCons(tmp11528, tmp11530))
				return

			} else {
				tmp11629 := PrimIsPair(V3641)

				var ifres11602 Obj

				if True == tmp11629 {
					tmp11627 := PrimHead(V3641)

					tmp11628 := PrimIsPair(tmp11627)

					var ifres11604 Obj

					if True == tmp11628 {
						tmp11624 := PrimHead(V3641)

						tmp11625 := PrimHead(tmp11624)

						tmp11626 := PrimEqual(symlet, tmp11625)

						var ifres11606 Obj

						if True == tmp11626 {
							tmp11621 := PrimHead(V3641)

							tmp11622 := PrimTail(tmp11621)

							tmp11623 := PrimIsPair(tmp11622)

							var ifres11608 Obj

							if True == tmp11623 {
								tmp11617 := PrimHead(V3641)

								tmp11618 := PrimTail(tmp11617)

								tmp11619 := PrimTail(tmp11618)

								tmp11620 := PrimIsPair(tmp11619)

								var ifres11610 Obj

								if True == tmp11620 {
									tmp11612 := PrimHead(V3641)

									tmp11613 := PrimTail(tmp11612)

									tmp11614 := PrimTail(tmp11613)

									tmp11615 := PrimTail(tmp11614)

									tmp11616 := PrimEqual(Nil, tmp11615)

									var ifres11611 Obj

									if True == tmp11616 {
										ifres11611 = True

									} else {
										ifres11611 = False

									}

									ifres11610 = ifres11611

								} else {
									ifres11610 = False

								}

								var ifres11609 Obj

								if True == ifres11610 {
									ifres11609 = True

								} else {
									ifres11609 = False

								}

								ifres11608 = ifres11609

							} else {
								ifres11608 = False

							}

							var ifres11607 Obj

							if True == ifres11608 {
								ifres11607 = True

							} else {
								ifres11607 = False

							}

							ifres11606 = ifres11607

						} else {
							ifres11606 = False

						}

						var ifres11605 Obj

						if True == ifres11606 {
							ifres11605 = True

						} else {
							ifres11605 = False

						}

						ifres11604 = ifres11605

					} else {
						ifres11604 = False

					}

					var ifres11603 Obj

					if True == ifres11604 {
						ifres11603 = True

					} else {
						ifres11603 = False

					}

					ifres11602 = ifres11603

				} else {
					ifres11602 = False

				}

				if True == ifres11602 {
					tmp11546 := PrimHead(V3641)

					tmp11547 := PrimTail(tmp11546)

					tmp11548 := PrimHead(tmp11547)

					tmp11549 := Call(__e, PrimFunc(symelement_2), tmp11548, V3639)

					if True == tmp11549 {
						tmp11531 := PrimHead(V3641)

						tmp11532 := PrimTail(tmp11531)

						tmp11533 := PrimCons(symis_b, tmp11532)

						tmp11534 := PrimTail(V3641)

						tmp11535 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3638, V3639, V3640, tmp11534, V3642)

						__e.Return(PrimCons(tmp11533, tmp11535))
						return

					} else {
						tmp11536 := PrimHead(V3641)

						tmp11537 := PrimTail(tmp11536)

						tmp11538 := PrimCons(symbind, tmp11537)

						tmp11539 := PrimHead(V3641)

						tmp11540 := PrimTail(tmp11539)

						tmp11541 := PrimHead(tmp11540)

						tmp11542 := PrimCons(tmp11541, V3639)

						tmp11543 := PrimTail(V3641)

						tmp11544 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3638, tmp11542, V3640, tmp11543, V3642)

						__e.Return(PrimCons(tmp11538, tmp11544))
						return

					}

				} else {
					tmp11600 := PrimIsPair(V3641)

					var ifres11580 Obj

					if True == tmp11600 {
						tmp11598 := PrimHead(V3641)

						tmp11599 := PrimIsPair(tmp11598)

						var ifres11582 Obj

						if True == tmp11599 {
							tmp11595 := PrimHead(V3641)

							tmp11596 := PrimHead(tmp11595)

							tmp11597 := PrimEqual(symctxt, tmp11596)

							var ifres11584 Obj

							if True == tmp11597 {
								tmp11592 := PrimHead(V3641)

								tmp11593 := PrimTail(tmp11592)

								tmp11594 := PrimIsPair(tmp11593)

								var ifres11586 Obj

								if True == tmp11594 {
									tmp11588 := PrimHead(V3641)

									tmp11589 := PrimTail(tmp11588)

									tmp11590 := PrimTail(tmp11589)

									tmp11591 := PrimEqual(Nil, tmp11590)

									var ifres11587 Obj

									if True == tmp11591 {
										ifres11587 = True

									} else {
										ifres11587 = False

									}

									ifres11586 = ifres11587

								} else {
									ifres11586 = False

								}

								var ifres11585 Obj

								if True == ifres11586 {
									ifres11585 = True

								} else {
									ifres11585 = False

								}

								ifres11584 = ifres11585

							} else {
								ifres11584 = False

							}

							var ifres11583 Obj

							if True == ifres11584 {
								ifres11583 = True

							} else {
								ifres11583 = False

							}

							ifres11582 = ifres11583

						} else {
							ifres11582 = False

						}

						var ifres11581 Obj

						if True == ifres11582 {
							ifres11581 = True

						} else {
							ifres11581 = False

						}

						ifres11580 = ifres11581

					} else {
						ifres11580 = False

					}

					if True == ifres11580 {
						tmp11575 := PrimHead(V3641)

						tmp11576 := PrimTail(tmp11575)

						tmp11577 := PrimHead(tmp11576)

						tmp11578 := Call(__e, PrimFunc(symelement_2), tmp11577, V3639)

						if True == tmp11578 {
							tmp11550 := PrimHead(V3641)

							tmp11551 := PrimTail(tmp11550)

							tmp11552 := PrimHead(tmp11551)

							tmp11553 := PrimCons(tmp11552, V3638)

							tmp11554 := PrimTail(V3641)

							__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp11553, V3639, V3640, tmp11554, V3642)
							return

						} else {
							tmp11555 := PrimHead(V3641)

							tmp11556 := PrimTail(tmp11555)

							tmp11557 := PrimHead(tmp11556)

							tmp11558 := PrimCons(V3640, Nil)

							tmp11559 := PrimCons(tmp11557, tmp11558)

							tmp11560 := PrimCons(symbind, tmp11559)

							tmp11561 := PrimHead(V3641)

							tmp11562 := PrimTail(tmp11561)

							tmp11563 := PrimHead(tmp11562)

							tmp11564 := PrimCons(tmp11563, V3638)

							tmp11565 := PrimHead(V3641)

							tmp11566 := PrimTail(tmp11565)

							tmp11567 := PrimHead(tmp11566)

							tmp11568 := PrimCons(tmp11567, V3639)

							tmp11569 := PrimHead(V3641)

							tmp11570 := PrimTail(tmp11569)

							tmp11571 := PrimHead(tmp11570)

							tmp11572 := PrimTail(V3641)

							tmp11573 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp11564, tmp11568, tmp11571, tmp11572, V3642)

							__e.Return(PrimCons(tmp11560, tmp11573))
							return

						}

					} else {
						__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4side_1conditions_1_6goals)
						return
					}

				}

			}

		}

	}, 5)

	tmp11654 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp11525)

	_ = tmp11654

	tmp11655 := MakeNative(func(__e *ControlFlow) {
		V3647 := __e.Get(1)
		_ = V3647
		V3648 := __e.Get(2)
		_ = V3648
		V3649 := __e.Get(3)
		_ = V3649
		tmp11705 := PrimEqual(Nil, V3649)

		if True == tmp11705 {
			tmp11656 := PrimIntern(MakeString(";"))

			__e.Return(PrimCons(tmp11656, Nil))
			return

		} else {
			tmp11703 := PrimIsPair(V3649)

			var ifres11699 Obj

			if True == tmp11703 {
				tmp11701 := PrimHead(V3649)

				tmp11702 := PrimEqual(sym_b, tmp11701)

				var ifres11700 Obj

				if True == tmp11702 {
					ifres11700 = True

				} else {
					ifres11700 = False

				}

				ifres11699 = ifres11700

			} else {
				ifres11699 = False

			}

			if True == ifres11699 {
				tmp11657 := PrimTail(V3649)

				tmp11658 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3647, V3648, tmp11657)

				__e.Return(PrimCons(sym_b, tmp11658))
				return

			} else {
				tmp11697 := PrimIsPair(V3649)

				var ifres11693 Obj

				if True == tmp11697 {
					tmp11695 := PrimHead(V3649)

					tmp11696 := PrimEqual(symfail, tmp11695)

					var ifres11694 Obj

					if True == tmp11696 {
						ifres11694 = True

					} else {
						ifres11694 = False

					}

					ifres11693 = ifres11694

				} else {
					ifres11693 = False

				}

				if True == ifres11693 {
					tmp11659 := PrimCons(False, Nil)

					tmp11660 := PrimCons(symwhen, tmp11659)

					tmp11661 := PrimTail(V3649)

					tmp11662 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3647, V3648, tmp11661)

					__e.Return(PrimCons(tmp11660, tmp11662))
					return

				} else {
					tmp11691 := PrimIsPair(V3649)

					var ifres11676 Obj

					if True == tmp11691 {
						tmp11689 := PrimHead(V3649)

						tmp11690 := PrimIsPair(tmp11689)

						var ifres11678 Obj

						if True == tmp11690 {
							tmp11686 := PrimHead(V3649)

							tmp11687 := PrimTail(tmp11686)

							tmp11688 := PrimIsPair(tmp11687)

							var ifres11680 Obj

							if True == tmp11688 {
								tmp11682 := PrimHead(V3649)

								tmp11683 := PrimTail(tmp11682)

								tmp11684 := PrimTail(tmp11683)

								tmp11685 := PrimEqual(Nil, tmp11684)

								var ifres11681 Obj

								if True == tmp11685 {
									ifres11681 = True

								} else {
									ifres11681 = False

								}

								ifres11680 = ifres11681

							} else {
								ifres11680 = False

							}

							var ifres11679 Obj

							if True == ifres11680 {
								ifres11679 = True

							} else {
								ifres11679 = False

							}

							ifres11678 = ifres11679

						} else {
							ifres11678 = False

						}

						var ifres11677 Obj

						if True == ifres11678 {
							ifres11677 = True

						} else {
							ifres11677 = False

						}

						ifres11676 = ifres11677

					} else {
						ifres11676 = False

					}

					if True == ifres11676 {
						tmp11663 := PrimHead(V3649)

						tmp11664 := PrimTail(tmp11663)

						tmp11665 := PrimHead(tmp11664)

						tmp11666 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp11665)

						tmp11667 := PrimHead(V3649)

						tmp11668 := PrimHead(tmp11667)

						tmp11669 := Call(__e, PrimFunc(symshen_4construct_1context), V3647, tmp11668, V3648)

						tmp11670 := PrimCons(tmp11669, Nil)

						tmp11671 := PrimCons(tmp11666, tmp11670)

						tmp11672 := PrimCons(symshen_4system_1S, tmp11671)

						tmp11673 := PrimTail(V3649)

						tmp11674 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3647, V3648, tmp11673)

						__e.Return(PrimCons(tmp11672, tmp11674))
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4premises_1_6goals)
						return
					}

				}

			}

		}

	}, 3)

	tmp11706 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp11655)

	_ = tmp11706

	tmp11707 := MakeNative(func(__e *ControlFlow) {
		V3653 := __e.Get(1)
		_ = V3653
		V3654 := __e.Get(2)
		_ = V3654
		V3655 := __e.Get(3)
		_ = V3655
		tmp11727 := PrimEqual(Nil, V3654)

		if True == tmp11727 {
			__e.Return(V3655)
			return
		} else {
			tmp11725 := PrimIsPair(V3654)

			var ifres11717 Obj

			if True == tmp11725 {
				tmp11723 := PrimTail(V3654)

				tmp11724 := PrimEqual(Nil, tmp11723)

				var ifres11719 Obj

				if True == tmp11724 {
					tmp11721 := PrimHead(V3654)

					tmp11722 := Call(__e, PrimFunc(symelement_2), tmp11721, V3653)

					var ifres11720 Obj

					if True == tmp11722 {
						ifres11720 = True

					} else {
						ifres11720 = False

					}

					ifres11719 = ifres11720

				} else {
					ifres11719 = False

				}

				var ifres11718 Obj

				if True == ifres11719 {
					ifres11718 = True

				} else {
					ifres11718 = False

				}

				ifres11717 = ifres11718

			} else {
				ifres11717 = False

			}

			if True == ifres11717 {
				__e.Return(PrimHead(V3654))
				return
			} else {
				tmp11715 := PrimIsPair(V3654)

				if True == tmp11715 {
					tmp11708 := PrimHead(V3654)

					tmp11709 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp11708)

					tmp11710 := PrimTail(V3654)

					tmp11711 := Call(__e, PrimFunc(symshen_4construct_1context), V3653, tmp11710, V3655)

					tmp11712 := PrimCons(tmp11711, Nil)

					tmp11713 := PrimCons(tmp11709, tmp11712)

					__e.Return(PrimCons(symcons, tmp11713))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4construct_1context)
					return
				}

			}

		}

	}, 3)

	tmp11728 := Call(__e, ns2_1set, symshen_4construct_1context, tmp11707)

	_ = tmp11728

	tmp11729 := MakeNative(func(__e *ControlFlow) {
		V3656 := __e.Get(1)
		_ = V3656
		tmp11730 := MakeNative(func(__e *ControlFlow) {
			W3657 := __e.Get(1)
			_ = W3657
			tmp11731 := MakeNative(func(__e *ControlFlow) {
				W3659 := __e.Get(1)
				_ = W3659
				tmp11732 := MakeNative(func(__e *ControlFlow) {
					W3660 := __e.Get(1)
					_ = W3660
					tmp11733 := MakeNative(func(__e *ControlFlow) {
						W3661 := __e.Get(1)
						_ = W3661
						__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3661)
						return
					}, 1)

					tmp11734 := PrimSet(symshen_4_ddatatypes_d, W3660)

					__e.TailApply(tmp11733, tmp11734)
					return

				}, 1)

				tmp11735 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3657, W3659)

				__e.TailApply(tmp11732, tmp11735)
				return

			}, 1)

			tmp11736 := PrimValue(symshen_4_ddatatypes_d)

			__e.TailApply(tmp11731, tmp11736)
			return

		}, 1)

		tmp11737 := MakeNative(func(__e *ControlFlow) {
			Z3658 := __e.Get(1)
			_ = Z3658
			__e.TailApply(PrimFunc(symshen_4intern_1type), Z3658)
			return
		}, 1)

		tmp11738 := Call(__e, PrimFunc(symmap), tmp11737, V3656)

		__e.TailApply(tmp11730, tmp11738)
		return

	}, 1)

	tmp11739 := Call(__e, ns2_1set, sympreclude, tmp11729)

	_ = tmp11739

	tmp11740 := MakeNative(func(__e *ControlFlow) {
		V3666 := __e.Get(1)
		_ = V3666
		V3667 := __e.Get(2)
		_ = V3667
		tmp11747 := PrimEqual(Nil, V3666)

		if True == tmp11747 {
			__e.Return(V3667)
			return
		} else {
			tmp11745 := PrimIsPair(V3666)

			if True == tmp11745 {
				tmp11741 := PrimTail(V3666)

				tmp11742 := PrimHead(V3666)

				tmp11743 := Call(__e, PrimFunc(symshen_4unassoc), tmp11742, V3667)

				__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp11741, tmp11743)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
				return
			}

		}

	}, 2)

	tmp11748 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp11740)

	_ = tmp11748

	tmp11749 := MakeNative(func(__e *ControlFlow) {
		V3668 := __e.Get(1)
		_ = V3668
		tmp11750 := MakeNative(func(__e *ControlFlow) {
			Z3669 := __e.Get(1)
			_ = Z3669
			__e.Return(PrimHead(Z3669))
			return
		}, 1)

		__e.TailApply(PrimFunc(symmap), tmp11750, V3668)
		return

	}, 1)

	tmp11751 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp11749)

	_ = tmp11751

	tmp11752 := MakeNative(func(__e *ControlFlow) {
		V3670 := __e.Get(1)
		_ = V3670
		tmp11753 := MakeNative(func(__e *ControlFlow) {
			W3671 := __e.Get(1)
			_ = W3671
			tmp11754 := MakeNative(func(__e *ControlFlow) {
				W3673 := __e.Get(1)
				_ = W3673
				tmp11755 := MakeNative(func(__e *ControlFlow) {
					W3675 := __e.Get(1)
					_ = W3675
					__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3675)
					return
				}, 1)

				tmp11756 := PrimValue(symshen_4_ddatatypes_d)

				__e.TailApply(tmp11755, tmp11756)
				return

			}, 1)

			tmp11757 := MakeNative(func(__e *ControlFlow) {
				Z3674 := __e.Get(1)
				_ = Z3674
				tmp11758 := Call(__e, PrimFunc(symfn), Z3674)

				__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3674, tmp11758)
				return

			}, 1)

			tmp11759 := Call(__e, PrimFunc(symmap), tmp11757, W3671)

			__e.TailApply(tmp11754, tmp11759)
			return

		}, 1)

		tmp11760 := MakeNative(func(__e *ControlFlow) {
			Z3672 := __e.Get(1)
			_ = Z3672
			__e.TailApply(PrimFunc(symshen_4intern_1type), Z3672)
			return
		}, 1)

		tmp11761 := Call(__e, PrimFunc(symmap), tmp11760, V3670)

		__e.TailApply(tmp11753, tmp11761)
		return

	}, 1)

	tmp11762 := Call(__e, ns2_1set, syminclude, tmp11752)

	_ = tmp11762

	tmp11763 := MakeNative(func(__e *ControlFlow) {
		V3676 := __e.Get(1)
		_ = V3676
		tmp11764 := MakeNative(func(__e *ControlFlow) {
			W3677 := __e.Get(1)
			_ = W3677
			tmp11765 := MakeNative(func(__e *ControlFlow) {
				W3678 := __e.Get(1)
				_ = W3678
				tmp11766 := MakeNative(func(__e *ControlFlow) {
					W3680 := __e.Get(1)
					_ = W3680
					tmp11767 := PrimValue(symshen_4_ddatatypes_d)

					__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp11767)
					return

				}, 1)

				tmp11768 := MakeNative(func(__e *ControlFlow) {
					Z3681 := __e.Get(1)
					_ = Z3681
					tmp11769 := Call(__e, PrimFunc(symfn), Z3681)

					__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3681, tmp11769)
					return

				}, 1)

				tmp11770 := Call(__e, PrimFunc(symmap), tmp11768, W3678)

				__e.TailApply(tmp11766, tmp11770)
				return

			}, 1)

			tmp11771 := MakeNative(func(__e *ControlFlow) {
				Z3679 := __e.Get(1)
				_ = Z3679
				__e.TailApply(PrimFunc(symshen_4intern_1type), Z3679)
				return
			}, 1)

			tmp11772 := Call(__e, PrimFunc(symmap), tmp11771, V3676)

			__e.TailApply(tmp11765, tmp11772)
			return

		}, 1)

		tmp11773 := PrimSet(symshen_4_ddatatypes_d, Nil)

		__e.TailApply(tmp11764, tmp11773)
		return

	}, 1)

	tmp11774 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp11763)

	_ = tmp11774

	tmp11775 := MakeNative(func(__e *ControlFlow) {
		V3682 := __e.Get(1)
		_ = V3682
		tmp11776 := MakeNative(func(__e *ControlFlow) {
			W3683 := __e.Get(1)
			_ = W3683
			tmp11777 := MakeNative(func(__e *ControlFlow) {
				W3685 := __e.Get(1)
				_ = W3685
				tmp11778 := MakeNative(func(__e *ControlFlow) {
					W3686 := __e.Get(1)
					_ = W3686
					__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3686)
					return
				}, 1)

				tmp11779 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3683, W3685)

				tmp11780 := PrimSet(symshen_4_ddatatypes_d, tmp11779)

				__e.TailApply(tmp11778, tmp11780)
				return

			}, 1)

			tmp11781 := PrimValue(symshen_4_dalldatatypes_d)

			__e.TailApply(tmp11777, tmp11781)
			return

		}, 1)

		tmp11782 := MakeNative(func(__e *ControlFlow) {
			Z3684 := __e.Get(1)
			_ = Z3684
			__e.TailApply(PrimFunc(symshen_4intern_1type), Z3684)
			return
		}, 1)

		tmp11783 := Call(__e, PrimFunc(symmap), tmp11782, V3682)

		__e.TailApply(tmp11776, tmp11783)
		return

	}, 1)

	__e.TailApply(ns2_1set, syminclude_1all_1but, tmp11775)
	return

}, 0)
