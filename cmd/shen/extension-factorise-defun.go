package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2012-2019 Bruno Deferrari.  All rights reserved.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

		gen2405 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4931 := __args[0]
			_ = V4931
			gen2403 := Call(__e, ShenFunc(symcons_2), V4931)

			var gen2404 Obj
			if True == gen2403 {
				gen2400 := Call(__e, ShenFunc(symhd), V4931)

				gen2401 := Call(__e, ShenFunc(sym_a), MakeSymbol("defun"), gen2400)

				var gen2402 Obj
				if True == gen2401 {
					gen2397 := Call(__e, ShenFunc(symtl), V4931)

					gen2398 := Call(__e, ShenFunc(symcons_2), gen2397)

					var gen2399 Obj
					if True == gen2398 {
						gen2393 := Call(__e, ShenFunc(symtl), V4931)

						gen2394 := Call(__e, ShenFunc(symtl), gen2393)

						gen2395 := Call(__e, ShenFunc(symcons_2), gen2394)

						var gen2396 Obj
						if True == gen2395 {
							gen2388 := Call(__e, ShenFunc(symtl), V4931)

							gen2389 := Call(__e, ShenFunc(symtl), gen2388)

							gen2390 := Call(__e, ShenFunc(symtl), gen2389)

							gen2391 := Call(__e, ShenFunc(symcons_2), gen2390)

							var gen2392 Obj
							if True == gen2391 {
								gen2382 := Call(__e, ShenFunc(symtl), V4931)

								gen2383 := Call(__e, ShenFunc(symtl), gen2382)

								gen2384 := Call(__e, ShenFunc(symtl), gen2383)

								gen2385 := Call(__e, ShenFunc(symhd), gen2384)

								gen2386 := Call(__e, ShenFunc(symcons_2), gen2385)

								var gen2387 Obj
								if True == gen2386 {
									gen2375 := Call(__e, ShenFunc(symtl), V4931)

									gen2376 := Call(__e, ShenFunc(symtl), gen2375)

									gen2377 := Call(__e, ShenFunc(symtl), gen2376)

									gen2378 := Call(__e, ShenFunc(symhd), gen2377)

									gen2379 := Call(__e, ShenFunc(symhd), gen2378)

									gen2380 := Call(__e, ShenFunc(sym_a), MakeSymbol("cond"), gen2379)

									var gen2381 Obj
									if True == gen2380 {
										gen2370 := Call(__e, ShenFunc(symtl), V4931)

										gen2371 := Call(__e, ShenFunc(symtl), gen2370)

										gen2372 := Call(__e, ShenFunc(symtl), gen2371)

										gen2373 := Call(__e, ShenFunc(symtl), gen2372)

										gen2374 := Call(__e, ShenFunc(sym_a), Nil, gen2373)

										if True == gen2374 {
											gen2381 = True
										} else {
											gen2381 = False
										}

									} else {
										gen2381 = False
									}
									if True == gen2381 {
										gen2387 = True
									} else {
										gen2387 = False
									}

								} else {
									gen2387 = False
								}
								if True == gen2387 {
									gen2392 = True
								} else {
									gen2392 = False
								}

							} else {
								gen2392 = False
							}
							if True == gen2392 {
								gen2396 = True
							} else {
								gen2396 = False
							}

						} else {
							gen2396 = False
						}
						if True == gen2396 {
							gen2399 = True
						} else {
							gen2399 = False
						}

					} else {
						gen2399 = False
					}
					if True == gen2399 {
						gen2402 = True
					} else {
						gen2402 = False
					}

				} else {
					gen2402 = False
				}
				if True == gen2402 {
					gen2404 = True
				} else {
					gen2404 = False
				}

			} else {
				gen2404 = False
			}
			if True == gen2404 {
				gen2350 := Call(__e, ShenFunc(symtl), V4931)

				gen2351 := Call(__e, ShenFunc(symhd), gen2350)

				gen2352 := Call(__e, ShenFunc(symtl), V4931)

				gen2353 := Call(__e, ShenFunc(symtl), gen2352)

				gen2354 := Call(__e, ShenFunc(symhd), gen2353)

				gen2355 := Call(__e, ShenFunc(symtl), V4931)

				gen2356 := Call(__e, ShenFunc(symtl), gen2355)

				gen2357 := Call(__e, ShenFunc(symtl), gen2356)

				gen2358 := Call(__e, ShenFunc(symhd), gen2357)

				gen2359 := Call(__e, ShenFunc(symtl), V4931)

				gen2360 := Call(__e, ShenFunc(symhd), gen2359)

				gen2361 := Call(__e, ShenFunc(symcons), gen2360, Nil)

				gen2362 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.f_error"), gen2361)

				gen2363 := Call(__e, ShenFunc(symtl), V4931)

				gen2364 := Call(__e, ShenFunc(symtl), gen2363)

				gen2365 := Call(__e, ShenFunc(symhd), gen2364)

				gen2366 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4factorise_1cond), gen2358, gen2362, gen2365)

				gen2367 := Call(__e, ShenFunc(symcons), gen2366, Nil)

				gen2368 := Call(__e, ShenFunc(symcons), gen2354, gen2367)

				gen2369 := Call(__e, ShenFunc(symcons), gen2351, gen2368)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("defun"), gen2369)

				return

			} else {
				__e.Return(V4931)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.factorise-defun"), gen2405)

		gen2413 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4943 := __args[0]
			_ = V4943
			V4944 := __args[1]
			_ = V4944
			V4945 := __args[2]
			_ = V4945
			gen2411 := Call(__e, ShenFunc(symcons_2), V4943)

			var gen2412 Obj
			if True == gen2411 {
				gen2409 := Call(__e, ShenFunc(symhd), V4943)

				gen2410 := Call(__e, ShenFunc(sym_a), MakeSymbol("cond"), gen2409)

				if True == gen2410 {
					gen2412 = True
				} else {
					gen2412 = False
				}

			} else {
				gen2412 = False
			}
			if True == gen2412 {
				gen2406 := Call(__e, ShenFunc(symtl), V4943)

				gen2407 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4add_1returns), gen2406)

				gen2408 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4rebranch), gen2407, V4944)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2408, V4945)

				return

			} else {
				__e.Return(V4943)
				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.factorise-cond"), gen2413)

		gen2437 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4947 := __args[0]
			_ = V4947
			gen2436 := Call(__e, ShenFunc(sym_a), Nil, V4947)

			if True == gen2436 {
				__e.Return(Nil)
				return
			} else {
				gen2434 := Call(__e, ShenFunc(symcons_2), V4947)

				var gen2435 Obj
				if True == gen2434 {
					gen2431 := Call(__e, ShenFunc(symhd), V4947)

					gen2432 := Call(__e, ShenFunc(symcons_2), gen2431)

					var gen2433 Obj
					if True == gen2432 {
						gen2427 := Call(__e, ShenFunc(symhd), V4947)

						gen2428 := Call(__e, ShenFunc(symtl), gen2427)

						gen2429 := Call(__e, ShenFunc(symcons_2), gen2428)

						var gen2430 Obj
						if True == gen2429 {
							gen2423 := Call(__e, ShenFunc(symhd), V4947)

							gen2424 := Call(__e, ShenFunc(symtl), gen2423)

							gen2425 := Call(__e, ShenFunc(symtl), gen2424)

							gen2426 := Call(__e, ShenFunc(sym_a), Nil, gen2425)

							if True == gen2426 {
								gen2430 = True
							} else {
								gen2430 = False
							}

						} else {
							gen2430 = False
						}
						if True == gen2430 {
							gen2433 = True
						} else {
							gen2433 = False
						}

					} else {
						gen2433 = False
					}
					if True == gen2433 {
						gen2435 = True
					} else {
						gen2435 = False
					}

				} else {
					gen2435 = False
				}
				if True == gen2435 {
					gen2414 := Call(__e, ShenFunc(symhd), V4947)

					gen2415 := Call(__e, ShenFunc(symhd), gen2414)

					gen2416 := Call(__e, ShenFunc(symhd), V4947)

					gen2417 := Call(__e, ShenFunc(symtl), gen2416)

					gen2418 := Call(__e, ShenFunc(symcons), MakeSymbol("%%return"), gen2417)

					gen2419 := Call(__e, ShenFunc(symcons), gen2418, Nil)

					gen2420 := Call(__e, ShenFunc(symcons), gen2415, gen2419)

					gen2421 := Call(__e, ShenFunc(symtl), V4947)

					gen2422 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4add_1returns), gen2421)

					__e.TailApply(ShenFunc(symcons), gen2420, gen2422)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.add-returns"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.add-returns"), gen2437)

		gen2438 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symgensym), MakeSymbol("%%label"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.generate-label"), gen2438)

		gen2440 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4950 := __args[0]
			_ = V4950
			V4951 := __args[1]
			_ = V4951
			gen2439 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4free_1variables_1h), V4950, V4951, Nil)

			__e.TailApply(ShenFunc(symreverse), gen2439)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.free-variables"), gen2440)

		gen2501 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4963 := __args[0]
			_ = V4963
			V4964 := __args[1]
			_ = V4964
			V4965 := __args[2]
			_ = V4965
			gen2499 := Call(__e, ShenFunc(symcons_2), V4963)

			var gen2500 Obj
			if True == gen2499 {
				gen2496 := Call(__e, ShenFunc(symhd), V4963)

				gen2497 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen2496)

				var gen2498 Obj
				if True == gen2497 {
					gen2493 := Call(__e, ShenFunc(symtl), V4963)

					gen2494 := Call(__e, ShenFunc(symcons_2), gen2493)

					var gen2495 Obj
					if True == gen2494 {
						gen2489 := Call(__e, ShenFunc(symtl), V4963)

						gen2490 := Call(__e, ShenFunc(symtl), gen2489)

						gen2491 := Call(__e, ShenFunc(symcons_2), gen2490)

						var gen2492 Obj
						if True == gen2491 {
							gen2484 := Call(__e, ShenFunc(symtl), V4963)

							gen2485 := Call(__e, ShenFunc(symtl), gen2484)

							gen2486 := Call(__e, ShenFunc(symtl), gen2485)

							gen2487 := Call(__e, ShenFunc(symcons_2), gen2486)

							var gen2488 Obj
							if True == gen2487 {
								gen2479 := Call(__e, ShenFunc(symtl), V4963)

								gen2480 := Call(__e, ShenFunc(symtl), gen2479)

								gen2481 := Call(__e, ShenFunc(symtl), gen2480)

								gen2482 := Call(__e, ShenFunc(symtl), gen2481)

								gen2483 := Call(__e, ShenFunc(sym_a), Nil, gen2482)

								if True == gen2483 {
									gen2488 = True
								} else {
									gen2488 = False
								}

							} else {
								gen2488 = False
							}
							if True == gen2488 {
								gen2492 = True
							} else {
								gen2492 = False
							}

						} else {
							gen2492 = False
						}
						if True == gen2492 {
							gen2495 = True
						} else {
							gen2495 = False
						}

					} else {
						gen2495 = False
					}
					if True == gen2495 {
						gen2498 = True
					} else {
						gen2498 = False
					}

				} else {
					gen2498 = False
				}
				if True == gen2498 {
					gen2500 = True
				} else {
					gen2500 = False
				}

			} else {
				gen2500 = False
			}
			if True == gen2500 {
				gen2468 := Call(__e, ShenFunc(symtl), V4963)

				gen2469 := Call(__e, ShenFunc(symtl), gen2468)

				gen2470 := Call(__e, ShenFunc(symtl), gen2469)

				gen2471 := Call(__e, ShenFunc(symhd), gen2470)

				gen2472 := Call(__e, ShenFunc(symtl), V4963)

				gen2473 := Call(__e, ShenFunc(symhd), gen2472)

				gen2474 := Call(__e, ShenFunc(symremove), gen2473, V4964)

				gen2475 := Call(__e, ShenFunc(symtl), V4963)

				gen2476 := Call(__e, ShenFunc(symtl), gen2475)

				gen2477 := Call(__e, ShenFunc(symhd), gen2476)

				gen2478 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4free_1variables_1h), gen2477, V4964, V4965)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4free_1variables_1h), gen2471, gen2474, gen2478)

				return

			} else {
				gen2466 := Call(__e, ShenFunc(symcons_2), V4963)

				var gen2467 Obj
				if True == gen2466 {
					gen2463 := Call(__e, ShenFunc(symhd), V4963)

					gen2464 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen2463)

					var gen2465 Obj
					if True == gen2464 {
						gen2460 := Call(__e, ShenFunc(symtl), V4963)

						gen2461 := Call(__e, ShenFunc(symcons_2), gen2460)

						var gen2462 Obj
						if True == gen2461 {
							gen2456 := Call(__e, ShenFunc(symtl), V4963)

							gen2457 := Call(__e, ShenFunc(symtl), gen2456)

							gen2458 := Call(__e, ShenFunc(symcons_2), gen2457)

							var gen2459 Obj
							if True == gen2458 {
								gen2452 := Call(__e, ShenFunc(symtl), V4963)

								gen2453 := Call(__e, ShenFunc(symtl), gen2452)

								gen2454 := Call(__e, ShenFunc(symtl), gen2453)

								gen2455 := Call(__e, ShenFunc(sym_a), Nil, gen2454)

								if True == gen2455 {
									gen2459 = True
								} else {
									gen2459 = False
								}

							} else {
								gen2459 = False
							}
							if True == gen2459 {
								gen2462 = True
							} else {
								gen2462 = False
							}

						} else {
							gen2462 = False
						}
						if True == gen2462 {
							gen2465 = True
						} else {
							gen2465 = False
						}

					} else {
						gen2465 = False
					}
					if True == gen2465 {
						gen2467 = True
					} else {
						gen2467 = False
					}

				} else {
					gen2467 = False
				}
				if True == gen2467 {
					gen2446 := Call(__e, ShenFunc(symtl), V4963)

					gen2447 := Call(__e, ShenFunc(symtl), gen2446)

					gen2448 := Call(__e, ShenFunc(symhd), gen2447)

					gen2449 := Call(__e, ShenFunc(symtl), V4963)

					gen2450 := Call(__e, ShenFunc(symhd), gen2449)

					gen2451 := Call(__e, ShenFunc(symremove), gen2450, V4964)

					__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4free_1variables_1h), gen2448, gen2451, V4965)

					return

				} else {
					gen2445 := Call(__e, ShenFunc(symcons_2), V4963)

					if True == gen2445 {
						gen2442 := Call(__e, ShenFunc(symtl), V4963)

						gen2443 := Call(__e, ShenFunc(symhd), V4963)

						gen2444 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4free_1variables_1h), gen2443, V4964, V4965)

						__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4free_1variables_1h), gen2442, V4964, gen2444)

						return

					} else {
						gen2441 := Call(__e, ShenFunc(symelement_2), V4963, V4964)

						if True == gen2441 {
							__e.TailApply(ShenFunc(symadjoin), V4963, V4965)

							return
						} else {
							__e.Return(V4965)
							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.free-variables-h"), gen2501)

		gen2555 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4968 := __args[0]
			_ = V4968
			V4969 := __args[1]
			_ = V4969
			gen2553 := Call(__e, ShenFunc(symcons_2), V4968)

			var gen2554 Obj
			if True == gen2553 {
				gen2550 := Call(__e, ShenFunc(symhd), V4968)

				gen2551 := Call(__e, ShenFunc(sym_a), MakeSymbol("%%let-label"), gen2550)

				var gen2552 Obj
				if True == gen2551 {
					gen2547 := Call(__e, ShenFunc(symtl), V4968)

					gen2548 := Call(__e, ShenFunc(symcons_2), gen2547)

					var gen2549 Obj
					if True == gen2548 {
						gen2543 := Call(__e, ShenFunc(symtl), V4968)

						gen2544 := Call(__e, ShenFunc(symtl), gen2543)

						gen2545 := Call(__e, ShenFunc(symcons_2), gen2544)

						var gen2546 Obj
						if True == gen2545 {
							gen2538 := Call(__e, ShenFunc(symtl), V4968)

							gen2539 := Call(__e, ShenFunc(symtl), gen2538)

							gen2540 := Call(__e, ShenFunc(symtl), gen2539)

							gen2541 := Call(__e, ShenFunc(symcons_2), gen2540)

							var gen2542 Obj
							if True == gen2541 {
								gen2533 := Call(__e, ShenFunc(symtl), V4968)

								gen2534 := Call(__e, ShenFunc(symtl), gen2533)

								gen2535 := Call(__e, ShenFunc(symtl), gen2534)

								gen2536 := Call(__e, ShenFunc(symtl), gen2535)

								gen2537 := Call(__e, ShenFunc(sym_a), Nil, gen2536)

								if True == gen2537 {
									gen2542 = True
								} else {
									gen2542 = False
								}

							} else {
								gen2542 = False
							}
							if True == gen2542 {
								gen2546 = True
							} else {
								gen2546 = False
							}

						} else {
							gen2546 = False
						}
						if True == gen2546 {
							gen2549 = True
						} else {
							gen2549 = False
						}

					} else {
						gen2549 = False
					}
					if True == gen2549 {
						gen2552 = True
					} else {
						gen2552 = False
					}

				} else {
					gen2552 = False
				}
				if True == gen2552 {
					gen2554 = True
				} else {
					gen2554 = False
				}

			} else {
				gen2554 = False
			}
			if True == gen2554 {
				gen2502 := Call(__e, ShenFunc(symtl), V4968)

				gen2503 := Call(__e, ShenFunc(symtl), gen2502)

				gen2504 := Call(__e, ShenFunc(symhd), gen2503)

				gen2505 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4free_1variables), gen2504, V4969)

				FreeVars := gen2505
				gen2521 := Call(__e, ShenFunc(sym_a), Nil, FreeVars)

				var gen2522 Obj
				if True == gen2521 {
					gen2518 := Call(__e, ShenFunc(symtl), V4968)

					gen2519 := Call(__e, ShenFunc(symtl), gen2518)

					gen2520 := Call(__e, ShenFunc(symtl), gen2519)

					gen2522 = Call(__e, ShenFunc(symhd), gen2520)

				} else {
					gen2506 := Call(__e, ShenFunc(symtl), V4968)

					gen2507 := Call(__e, ShenFunc(symhd), gen2506)

					gen2508 := Call(__e, ShenFunc(symcons), gen2507, FreeVars)

					gen2509 := Call(__e, ShenFunc(symcons), MakeSymbol("%%goto-label"), gen2508)

					gen2510 := Call(__e, ShenFunc(symtl), V4968)

					gen2511 := Call(__e, ShenFunc(symhd), gen2510)

					gen2512 := Call(__e, ShenFunc(symcons), gen2511, Nil)

					gen2513 := Call(__e, ShenFunc(symcons), MakeSymbol("%%goto-label"), gen2512)

					gen2514 := Call(__e, ShenFunc(symtl), V4968)

					gen2515 := Call(__e, ShenFunc(symtl), gen2514)

					gen2516 := Call(__e, ShenFunc(symtl), gen2515)

					gen2517 := Call(__e, ShenFunc(symhd), gen2516)

					gen2522 = Call(__e, ShenFunc(symsubst), gen2509, gen2513, gen2517)

				}
				NewBody := gen2522
				gen2523 := Call(__e, ShenFunc(symtl), V4968)

				gen2524 := Call(__e, ShenFunc(symhd), gen2523)

				gen2525 := Call(__e, ShenFunc(symcons), gen2524, FreeVars)

				gen2526 := Call(__e, ShenFunc(symtl), V4968)

				gen2527 := Call(__e, ShenFunc(symtl), gen2526)

				gen2528 := Call(__e, ShenFunc(symhd), gen2527)

				gen2529 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), NewBody, V4969)

				gen2530 := Call(__e, ShenFunc(symcons), gen2529, Nil)

				gen2531 := Call(__e, ShenFunc(symcons), gen2528, gen2530)

				gen2532 := Call(__e, ShenFunc(symcons), gen2525, gen2531)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("%%let-label"), gen2532)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.attach-free-variables"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.attach-free-variables"), gen2555)

		gen2710 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4976 := __args[0]
			_ = V4976
			V4977 := __args[1]
			_ = V4977
			gen2708 := Call(__e, ShenFunc(symcons_2), V4976)

			var gen2709 Obj
			if True == gen2708 {
				gen2705 := Call(__e, ShenFunc(symhd), V4976)

				gen2706 := Call(__e, ShenFunc(sym_a), MakeSymbol("%%let-label"), gen2705)

				var gen2707 Obj
				if True == gen2706 {
					gen2702 := Call(__e, ShenFunc(symtl), V4976)

					gen2703 := Call(__e, ShenFunc(symcons_2), gen2702)

					var gen2704 Obj
					if True == gen2703 {
						gen2698 := Call(__e, ShenFunc(symtl), V4976)

						gen2699 := Call(__e, ShenFunc(symtl), gen2698)

						gen2700 := Call(__e, ShenFunc(symcons_2), gen2699)

						var gen2701 Obj
						if True == gen2700 {
							gen2693 := Call(__e, ShenFunc(symtl), V4976)

							gen2694 := Call(__e, ShenFunc(symtl), gen2693)

							gen2695 := Call(__e, ShenFunc(symtl), gen2694)

							gen2696 := Call(__e, ShenFunc(symcons_2), gen2695)

							var gen2697 Obj
							if True == gen2696 {
								gen2687 := Call(__e, ShenFunc(symtl), V4976)

								gen2688 := Call(__e, ShenFunc(symtl), gen2687)

								gen2689 := Call(__e, ShenFunc(symtl), gen2688)

								gen2690 := Call(__e, ShenFunc(symtl), gen2689)

								gen2691 := Call(__e, ShenFunc(sym_a), Nil, gen2690)

								var gen2692 Obj
								if True == gen2691 {
									gen2677 := Call(__e, ShenFunc(symtl), V4976)

									gen2678 := Call(__e, ShenFunc(symhd), gen2677)

									gen2679 := Call(__e, ShenFunc(symcons), gen2678, Nil)

									gen2680 := Call(__e, ShenFunc(symcons), MakeSymbol("%%goto-label"), gen2679)

									gen2681 := Call(__e, ShenFunc(symtl), V4976)

									gen2682 := Call(__e, ShenFunc(symtl), gen2681)

									gen2683 := Call(__e, ShenFunc(symtl), gen2682)

									gen2684 := Call(__e, ShenFunc(symhd), gen2683)

									gen2685 := Call(__e, ShenFunc(symoccurrences), gen2680, gen2684)

									gen2686 := Call(__e, ShenFunc(sym_6), gen2685, MakeNumber(1))

									if True == gen2686 {
										gen2692 = True
									} else {
										gen2692 = False
									}

								} else {
									gen2692 = False
								}
								if True == gen2692 {
									gen2697 = True
								} else {
									gen2697 = False
								}

							} else {
								gen2697 = False
							}
							if True == gen2697 {
								gen2701 = True
							} else {
								gen2701 = False
							}

						} else {
							gen2701 = False
						}
						if True == gen2701 {
							gen2704 = True
						} else {
							gen2704 = False
						}

					} else {
						gen2704 = False
					}
					if True == gen2704 {
						gen2707 = True
					} else {
						gen2707 = False
					}

				} else {
					gen2707 = False
				}
				if True == gen2707 {
					gen2709 = True
				} else {
					gen2709 = False
				}

			} else {
				gen2709 = False
			}
			if True == gen2709 {
				gen2665 := Call(__e, ShenFunc(symtl), V4976)

				gen2666 := Call(__e, ShenFunc(symhd), gen2665)

				gen2667 := Call(__e, ShenFunc(symtl), V4976)

				gen2668 := Call(__e, ShenFunc(symtl), gen2667)

				gen2669 := Call(__e, ShenFunc(symhd), gen2668)

				gen2670 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2669, V4977)

				gen2671 := Call(__e, ShenFunc(symtl), V4976)

				gen2672 := Call(__e, ShenFunc(symtl), gen2671)

				gen2673 := Call(__e, ShenFunc(symtl), gen2672)

				gen2674 := Call(__e, ShenFunc(symcons), gen2670, gen2673)

				gen2675 := Call(__e, ShenFunc(symcons), gen2666, gen2674)

				gen2676 := Call(__e, ShenFunc(symcons), MakeSymbol("%%let-label"), gen2675)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4attach_1free_1variables), gen2676, V4977)

				return

			} else {
				gen2663 := Call(__e, ShenFunc(symcons_2), V4976)

				var gen2664 Obj
				if True == gen2663 {
					gen2660 := Call(__e, ShenFunc(symhd), V4976)

					gen2661 := Call(__e, ShenFunc(sym_a), MakeSymbol("%%let-label"), gen2660)

					var gen2662 Obj
					if True == gen2661 {
						gen2657 := Call(__e, ShenFunc(symtl), V4976)

						gen2658 := Call(__e, ShenFunc(symcons_2), gen2657)

						var gen2659 Obj
						if True == gen2658 {
							gen2653 := Call(__e, ShenFunc(symtl), V4976)

							gen2654 := Call(__e, ShenFunc(symtl), gen2653)

							gen2655 := Call(__e, ShenFunc(symcons_2), gen2654)

							var gen2656 Obj
							if True == gen2655 {
								gen2648 := Call(__e, ShenFunc(symtl), V4976)

								gen2649 := Call(__e, ShenFunc(symtl), gen2648)

								gen2650 := Call(__e, ShenFunc(symtl), gen2649)

								gen2651 := Call(__e, ShenFunc(symcons_2), gen2650)

								var gen2652 Obj
								if True == gen2651 {
									gen2643 := Call(__e, ShenFunc(symtl), V4976)

									gen2644 := Call(__e, ShenFunc(symtl), gen2643)

									gen2645 := Call(__e, ShenFunc(symtl), gen2644)

									gen2646 := Call(__e, ShenFunc(symtl), gen2645)

									gen2647 := Call(__e, ShenFunc(sym_a), Nil, gen2646)

									if True == gen2647 {
										gen2652 = True
									} else {
										gen2652 = False
									}

								} else {
									gen2652 = False
								}
								if True == gen2652 {
									gen2656 = True
								} else {
									gen2656 = False
								}

							} else {
								gen2656 = False
							}
							if True == gen2656 {
								gen2659 = True
							} else {
								gen2659 = False
							}

						} else {
							gen2659 = False
						}
						if True == gen2659 {
							gen2662 = True
						} else {
							gen2662 = False
						}

					} else {
						gen2662 = False
					}
					if True == gen2662 {
						gen2664 = True
					} else {
						gen2664 = False
					}

				} else {
					gen2664 = False
				}
				if True == gen2664 {
					gen2630 := Call(__e, ShenFunc(symtl), V4976)

					gen2631 := Call(__e, ShenFunc(symtl), gen2630)

					gen2632 := Call(__e, ShenFunc(symhd), gen2631)

					gen2633 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2632, V4977)

					gen2634 := Call(__e, ShenFunc(symtl), V4976)

					gen2635 := Call(__e, ShenFunc(symhd), gen2634)

					gen2636 := Call(__e, ShenFunc(symcons), gen2635, Nil)

					gen2637 := Call(__e, ShenFunc(symcons), MakeSymbol("%%goto-label"), gen2636)

					gen2638 := Call(__e, ShenFunc(symtl), V4976)

					gen2639 := Call(__e, ShenFunc(symtl), gen2638)

					gen2640 := Call(__e, ShenFunc(symtl), gen2639)

					gen2641 := Call(__e, ShenFunc(symhd), gen2640)

					gen2642 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2641, V4977)

					__e.TailApply(ShenFunc(symsubst), gen2633, gen2637, gen2642)

					return

				} else {
					gen2628 := Call(__e, ShenFunc(symcons_2), V4976)

					var gen2629 Obj
					if True == gen2628 {
						gen2625 := Call(__e, ShenFunc(symhd), V4976)

						gen2626 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen2625)

						var gen2627 Obj
						if True == gen2626 {
							gen2622 := Call(__e, ShenFunc(symtl), V4976)

							gen2623 := Call(__e, ShenFunc(symcons_2), gen2622)

							var gen2624 Obj
							if True == gen2623 {
								gen2618 := Call(__e, ShenFunc(symtl), V4976)

								gen2619 := Call(__e, ShenFunc(symtl), gen2618)

								gen2620 := Call(__e, ShenFunc(symcons_2), gen2619)

								var gen2621 Obj
								if True == gen2620 {
									gen2613 := Call(__e, ShenFunc(symtl), V4976)

									gen2614 := Call(__e, ShenFunc(symtl), gen2613)

									gen2615 := Call(__e, ShenFunc(symtl), gen2614)

									gen2616 := Call(__e, ShenFunc(symcons_2), gen2615)

									var gen2617 Obj
									if True == gen2616 {
										gen2608 := Call(__e, ShenFunc(symtl), V4976)

										gen2609 := Call(__e, ShenFunc(symtl), gen2608)

										gen2610 := Call(__e, ShenFunc(symtl), gen2609)

										gen2611 := Call(__e, ShenFunc(symtl), gen2610)

										gen2612 := Call(__e, ShenFunc(sym_a), Nil, gen2611)

										if True == gen2612 {
											gen2617 = True
										} else {
											gen2617 = False
										}

									} else {
										gen2617 = False
									}
									if True == gen2617 {
										gen2621 = True
									} else {
										gen2621 = False
									}

								} else {
									gen2621 = False
								}
								if True == gen2621 {
									gen2624 = True
								} else {
									gen2624 = False
								}

							} else {
								gen2624 = False
							}
							if True == gen2624 {
								gen2627 = True
							} else {
								gen2627 = False
							}

						} else {
							gen2627 = False
						}
						if True == gen2627 {
							gen2629 = True
						} else {
							gen2629 = False
						}

					} else {
						gen2629 = False
					}
					if True == gen2629 {
						gen2594 := Call(__e, ShenFunc(symtl), V4976)

						gen2595 := Call(__e, ShenFunc(symhd), gen2594)

						gen2596 := Call(__e, ShenFunc(symtl), V4976)

						gen2597 := Call(__e, ShenFunc(symtl), gen2596)

						gen2598 := Call(__e, ShenFunc(symhd), gen2597)

						gen2599 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2598, V4977)

						gen2600 := Call(__e, ShenFunc(symtl), V4976)

						gen2601 := Call(__e, ShenFunc(symtl), gen2600)

						gen2602 := Call(__e, ShenFunc(symtl), gen2601)

						gen2603 := Call(__e, ShenFunc(symhd), gen2602)

						gen2604 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2603, V4977)

						gen2605 := Call(__e, ShenFunc(symcons), gen2604, Nil)

						gen2606 := Call(__e, ShenFunc(symcons), gen2599, gen2605)

						gen2607 := Call(__e, ShenFunc(symcons), gen2595, gen2606)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen2607)

						return

					} else {
						gen2592 := Call(__e, ShenFunc(symcons_2), V4976)

						var gen2593 Obj
						if True == gen2592 {
							gen2589 := Call(__e, ShenFunc(symhd), V4976)

							gen2590 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen2589)

							var gen2591 Obj
							if True == gen2590 {
								gen2586 := Call(__e, ShenFunc(symtl), V4976)

								gen2587 := Call(__e, ShenFunc(symcons_2), gen2586)

								var gen2588 Obj
								if True == gen2587 {
									gen2582 := Call(__e, ShenFunc(symtl), V4976)

									gen2583 := Call(__e, ShenFunc(symtl), gen2582)

									gen2584 := Call(__e, ShenFunc(symcons_2), gen2583)

									var gen2585 Obj
									if True == gen2584 {
										gen2577 := Call(__e, ShenFunc(symtl), V4976)

										gen2578 := Call(__e, ShenFunc(symtl), gen2577)

										gen2579 := Call(__e, ShenFunc(symtl), gen2578)

										gen2580 := Call(__e, ShenFunc(symcons_2), gen2579)

										var gen2581 Obj
										if True == gen2580 {
											gen2572 := Call(__e, ShenFunc(symtl), V4976)

											gen2573 := Call(__e, ShenFunc(symtl), gen2572)

											gen2574 := Call(__e, ShenFunc(symtl), gen2573)

											gen2575 := Call(__e, ShenFunc(symtl), gen2574)

											gen2576 := Call(__e, ShenFunc(sym_a), Nil, gen2575)

											if True == gen2576 {
												gen2581 = True
											} else {
												gen2581 = False
											}

										} else {
											gen2581 = False
										}
										if True == gen2581 {
											gen2585 = True
										} else {
											gen2585 = False
										}

									} else {
										gen2585 = False
									}
									if True == gen2585 {
										gen2588 = True
									} else {
										gen2588 = False
									}

								} else {
									gen2588 = False
								}
								if True == gen2588 {
									gen2591 = True
								} else {
									gen2591 = False
								}

							} else {
								gen2591 = False
							}
							if True == gen2591 {
								gen2593 = True
							} else {
								gen2593 = False
							}

						} else {
							gen2593 = False
						}
						if True == gen2593 {
							gen2556 := Call(__e, ShenFunc(symtl), V4976)

							gen2557 := Call(__e, ShenFunc(symhd), gen2556)

							gen2558 := Call(__e, ShenFunc(symtl), V4976)

							gen2559 := Call(__e, ShenFunc(symtl), gen2558)

							gen2560 := Call(__e, ShenFunc(symhd), gen2559)

							gen2561 := Call(__e, ShenFunc(symtl), V4976)

							gen2562 := Call(__e, ShenFunc(symtl), gen2561)

							gen2563 := Call(__e, ShenFunc(symtl), gen2562)

							gen2564 := Call(__e, ShenFunc(symhd), gen2563)

							gen2565 := Call(__e, ShenFunc(symtl), V4976)

							gen2566 := Call(__e, ShenFunc(symhd), gen2565)

							gen2567 := Call(__e, ShenFunc(symcons), gen2566, V4977)

							gen2568 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen2564, gen2567)

							gen2569 := Call(__e, ShenFunc(symcons), gen2568, Nil)

							gen2570 := Call(__e, ShenFunc(symcons), gen2560, gen2569)

							gen2571 := Call(__e, ShenFunc(symcons), gen2557, gen2570)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen2571)

							return

						} else {
							__e.Return(V4976)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.inline-mono-labels"), gen2710)

		gen2806 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4984 := __args[0]
			_ = V4984
			V4985 := __args[1]
			_ = V4985
			gen2805 := Call(__e, ShenFunc(sym_a), Nil, V4984)

			if True == gen2805 {
				__e.Return(V4985)
				return
			} else {
				gen2803 := Call(__e, ShenFunc(symcons_2), V4984)

				var gen2804 Obj
				if True == gen2803 {
					gen2800 := Call(__e, ShenFunc(symhd), V4984)

					gen2801 := Call(__e, ShenFunc(symcons_2), gen2800)

					var gen2802 Obj
					if True == gen2801 {
						gen2796 := Call(__e, ShenFunc(symhd), V4984)

						gen2797 := Call(__e, ShenFunc(symhd), gen2796)

						gen2798 := Call(__e, ShenFunc(sym_a), True, gen2797)

						var gen2799 Obj
						if True == gen2798 {
							gen2792 := Call(__e, ShenFunc(symhd), V4984)

							gen2793 := Call(__e, ShenFunc(symtl), gen2792)

							gen2794 := Call(__e, ShenFunc(symcons_2), gen2793)

							var gen2795 Obj
							if True == gen2794 {
								gen2788 := Call(__e, ShenFunc(symhd), V4984)

								gen2789 := Call(__e, ShenFunc(symtl), gen2788)

								gen2790 := Call(__e, ShenFunc(symtl), gen2789)

								gen2791 := Call(__e, ShenFunc(sym_a), Nil, gen2790)

								if True == gen2791 {
									gen2795 = True
								} else {
									gen2795 = False
								}

							} else {
								gen2795 = False
							}
							if True == gen2795 {
								gen2799 = True
							} else {
								gen2799 = False
							}

						} else {
							gen2799 = False
						}
						if True == gen2799 {
							gen2802 = True
						} else {
							gen2802 = False
						}

					} else {
						gen2802 = False
					}
					if True == gen2802 {
						gen2804 = True
					} else {
						gen2804 = False
					}

				} else {
					gen2804 = False
				}
				if True == gen2804 {
					gen2786 := Call(__e, ShenFunc(symhd), V4984)

					gen2787 := Call(__e, ShenFunc(symtl), gen2786)

					__e.TailApply(ShenFunc(symhd), gen2787)

					return

				} else {
					gen2784 := Call(__e, ShenFunc(symcons_2), V4984)

					var gen2785 Obj
					if True == gen2784 {
						gen2781 := Call(__e, ShenFunc(symhd), V4984)

						gen2782 := Call(__e, ShenFunc(symcons_2), gen2781)

						var gen2783 Obj
						if True == gen2782 {
							gen2777 := Call(__e, ShenFunc(symhd), V4984)

							gen2778 := Call(__e, ShenFunc(symhd), gen2777)

							gen2779 := Call(__e, ShenFunc(symcons_2), gen2778)

							var gen2780 Obj
							if True == gen2779 {
								gen2772 := Call(__e, ShenFunc(symhd), V4984)

								gen2773 := Call(__e, ShenFunc(symhd), gen2772)

								gen2774 := Call(__e, ShenFunc(symhd), gen2773)

								gen2775 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen2774)

								var gen2776 Obj
								if True == gen2775 {
									gen2767 := Call(__e, ShenFunc(symhd), V4984)

									gen2768 := Call(__e, ShenFunc(symhd), gen2767)

									gen2769 := Call(__e, ShenFunc(symtl), gen2768)

									gen2770 := Call(__e, ShenFunc(symcons_2), gen2769)

									var gen2771 Obj
									if True == gen2770 {
										gen2761 := Call(__e, ShenFunc(symhd), V4984)

										gen2762 := Call(__e, ShenFunc(symhd), gen2761)

										gen2763 := Call(__e, ShenFunc(symtl), gen2762)

										gen2764 := Call(__e, ShenFunc(symtl), gen2763)

										gen2765 := Call(__e, ShenFunc(symcons_2), gen2764)

										var gen2766 Obj
										if True == gen2765 {
											gen2754 := Call(__e, ShenFunc(symhd), V4984)

											gen2755 := Call(__e, ShenFunc(symhd), gen2754)

											gen2756 := Call(__e, ShenFunc(symtl), gen2755)

											gen2757 := Call(__e, ShenFunc(symtl), gen2756)

											gen2758 := Call(__e, ShenFunc(symtl), gen2757)

											gen2759 := Call(__e, ShenFunc(sym_a), Nil, gen2758)

											var gen2760 Obj
											if True == gen2759 {
												gen2750 := Call(__e, ShenFunc(symhd), V4984)

												gen2751 := Call(__e, ShenFunc(symtl), gen2750)

												gen2752 := Call(__e, ShenFunc(symcons_2), gen2751)

												var gen2753 Obj
												if True == gen2752 {
													gen2746 := Call(__e, ShenFunc(symhd), V4984)

													gen2747 := Call(__e, ShenFunc(symtl), gen2746)

													gen2748 := Call(__e, ShenFunc(symtl), gen2747)

													gen2749 := Call(__e, ShenFunc(sym_a), Nil, gen2748)

													if True == gen2749 {
														gen2753 = True
													} else {
														gen2753 = False
													}

												} else {
													gen2753 = False
												}
												if True == gen2753 {
													gen2760 = True
												} else {
													gen2760 = False
												}

											} else {
												gen2760 = False
											}
											if True == gen2760 {
												gen2766 = True
											} else {
												gen2766 = False
											}

										} else {
											gen2766 = False
										}
										if True == gen2766 {
											gen2771 = True
										} else {
											gen2771 = False
										}

									} else {
										gen2771 = False
									}
									if True == gen2771 {
										gen2776 = True
									} else {
										gen2776 = False
									}

								} else {
									gen2776 = False
								}
								if True == gen2776 {
									gen2780 = True
								} else {
									gen2780 = False
								}

							} else {
								gen2780 = False
							}
							if True == gen2780 {
								gen2783 = True
							} else {
								gen2783 = False
							}

						} else {
							gen2783 = False
						}
						if True == gen2783 {
							gen2785 = True
						} else {
							gen2785 = False
						}

					} else {
						gen2785 = False
					}
					if True == gen2785 {
						gen2732 := Call(__e, ShenFunc(symhd), V4984)

						gen2733 := Call(__e, ShenFunc(symhd), gen2732)

						gen2734 := Call(__e, ShenFunc(symtl), gen2733)

						gen2735 := Call(__e, ShenFunc(symhd), gen2734)

						gen2736 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4true_1branch), gen2735, V4984)

						TrueBranch := gen2736
						gen2737 := Call(__e, ShenFunc(symhd), V4984)

						gen2738 := Call(__e, ShenFunc(symhd), gen2737)

						gen2739 := Call(__e, ShenFunc(symtl), gen2738)

						gen2740 := Call(__e, ShenFunc(symhd), gen2739)

						gen2741 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4false_1branch), gen2740, V4984)

						FalseBranch := gen2741
						gen2742 := Call(__e, ShenFunc(symhd), V4984)

						gen2743 := Call(__e, ShenFunc(symhd), gen2742)

						gen2744 := Call(__e, ShenFunc(symtl), gen2743)

						gen2745 := Call(__e, ShenFunc(symhd), gen2744)

						__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4rebranch_1h), gen2745, TrueBranch, FalseBranch, V4985)

						return

					} else {
						gen2730 := Call(__e, ShenFunc(symcons_2), V4984)

						var gen2731 Obj
						if True == gen2730 {
							gen2727 := Call(__e, ShenFunc(symhd), V4984)

							gen2728 := Call(__e, ShenFunc(symcons_2), gen2727)

							var gen2729 Obj
							if True == gen2728 {
								gen2723 := Call(__e, ShenFunc(symhd), V4984)

								gen2724 := Call(__e, ShenFunc(symtl), gen2723)

								gen2725 := Call(__e, ShenFunc(symcons_2), gen2724)

								var gen2726 Obj
								if True == gen2725 {
									gen2719 := Call(__e, ShenFunc(symhd), V4984)

									gen2720 := Call(__e, ShenFunc(symtl), gen2719)

									gen2721 := Call(__e, ShenFunc(symtl), gen2720)

									gen2722 := Call(__e, ShenFunc(sym_a), Nil, gen2721)

									if True == gen2722 {
										gen2726 = True
									} else {
										gen2726 = False
									}

								} else {
									gen2726 = False
								}
								if True == gen2726 {
									gen2729 = True
								} else {
									gen2729 = False
								}

							} else {
								gen2729 = False
							}
							if True == gen2729 {
								gen2731 = True
							} else {
								gen2731 = False
							}

						} else {
							gen2731 = False
						}
						if True == gen2731 {
							gen2711 := Call(__e, ShenFunc(symhd), V4984)

							gen2712 := Call(__e, ShenFunc(symhd), gen2711)

							gen2713 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4true_1branch), gen2712, V4984)

							TrueBranch := gen2713
							gen2714 := Call(__e, ShenFunc(symhd), V4984)

							gen2715 := Call(__e, ShenFunc(symhd), gen2714)

							gen2716 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4false_1branch), gen2715, V4984)

							FalseBranch := gen2716
							gen2717 := Call(__e, ShenFunc(symhd), V4984)

							gen2718 := Call(__e, ShenFunc(symhd), gen2717)

							__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4rebranch_1h), gen2718, TrueBranch, FalseBranch, V4985)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.rebranch"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.rebranch"), gen2806)

		gen2815 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4990 := __args[0]
			_ = V4990
			V4991 := __args[1]
			_ = V4991
			V4992 := __args[2]
			_ = V4992
			V4993 := __args[3]
			_ = V4993
			gen2807 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4rebranch), V4992, V4993)

			NewElse := gen2807
			gen2814 := MakeNative(func(__e Evaluator, __args ...Obj) {
				GotoElse := __args[0]
				_ = GotoElse
				gen2808 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4rebranch), V4991, GotoElse)

				gen2809 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4optimize_1selectors), V4990, gen2808)

				gen2810 := Call(__e, ShenFunc(symcons), GotoElse, Nil)

				gen2811 := Call(__e, ShenFunc(symcons), gen2809, gen2810)

				gen2812 := Call(__e, ShenFunc(symcons), V4990, gen2811)

				gen2813 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen2812)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs), gen2813)

				return

			}, 1)
			__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4with_1labelled_1else), NewElse, gen2814)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.rebranch-h"), gen2815)

		gen2896 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5006 := __args[0]
			_ = V5006
			V5007 := __args[1]
			_ = V5007
			gen2894 := Call(__e, ShenFunc(symcons_2), V5007)

			var gen2895 Obj
			if True == gen2894 {
				gen2891 := Call(__e, ShenFunc(symhd), V5007)

				gen2892 := Call(__e, ShenFunc(symcons_2), gen2891)

				var gen2893 Obj
				if True == gen2892 {
					gen2887 := Call(__e, ShenFunc(symhd), V5007)

					gen2888 := Call(__e, ShenFunc(symhd), gen2887)

					gen2889 := Call(__e, ShenFunc(symcons_2), gen2888)

					var gen2890 Obj
					if True == gen2889 {
						gen2882 := Call(__e, ShenFunc(symhd), V5007)

						gen2883 := Call(__e, ShenFunc(symhd), gen2882)

						gen2884 := Call(__e, ShenFunc(symhd), gen2883)

						gen2885 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen2884)

						var gen2886 Obj
						if True == gen2885 {
							gen2877 := Call(__e, ShenFunc(symhd), V5007)

							gen2878 := Call(__e, ShenFunc(symhd), gen2877)

							gen2879 := Call(__e, ShenFunc(symtl), gen2878)

							gen2880 := Call(__e, ShenFunc(symcons_2), gen2879)

							var gen2881 Obj
							if True == gen2880 {
								gen2871 := Call(__e, ShenFunc(symhd), V5007)

								gen2872 := Call(__e, ShenFunc(symhd), gen2871)

								gen2873 := Call(__e, ShenFunc(symtl), gen2872)

								gen2874 := Call(__e, ShenFunc(symtl), gen2873)

								gen2875 := Call(__e, ShenFunc(symcons_2), gen2874)

								var gen2876 Obj
								if True == gen2875 {
									gen2864 := Call(__e, ShenFunc(symhd), V5007)

									gen2865 := Call(__e, ShenFunc(symhd), gen2864)

									gen2866 := Call(__e, ShenFunc(symtl), gen2865)

									gen2867 := Call(__e, ShenFunc(symtl), gen2866)

									gen2868 := Call(__e, ShenFunc(symtl), gen2867)

									gen2869 := Call(__e, ShenFunc(sym_a), Nil, gen2868)

									var gen2870 Obj
									if True == gen2869 {
										gen2860 := Call(__e, ShenFunc(symhd), V5007)

										gen2861 := Call(__e, ShenFunc(symtl), gen2860)

										gen2862 := Call(__e, ShenFunc(symcons_2), gen2861)

										var gen2863 Obj
										if True == gen2862 {
											gen2855 := Call(__e, ShenFunc(symhd), V5007)

											gen2856 := Call(__e, ShenFunc(symtl), gen2855)

											gen2857 := Call(__e, ShenFunc(symtl), gen2856)

											gen2858 := Call(__e, ShenFunc(sym_a), Nil, gen2857)

											var gen2859 Obj
											if True == gen2858 {
												gen2850 := Call(__e, ShenFunc(symhd), V5007)

												gen2851 := Call(__e, ShenFunc(symhd), gen2850)

												gen2852 := Call(__e, ShenFunc(symtl), gen2851)

												gen2853 := Call(__e, ShenFunc(symhd), gen2852)

												gen2854 := Call(__e, ShenFunc(sym_a), gen2853, V5006)

												if True == gen2854 {
													gen2859 = True
												} else {
													gen2859 = False
												}

											} else {
												gen2859 = False
											}
											if True == gen2859 {
												gen2863 = True
											} else {
												gen2863 = False
											}

										} else {
											gen2863 = False
										}
										if True == gen2863 {
											gen2870 = True
										} else {
											gen2870 = False
										}

									} else {
										gen2870 = False
									}
									if True == gen2870 {
										gen2876 = True
									} else {
										gen2876 = False
									}

								} else {
									gen2876 = False
								}
								if True == gen2876 {
									gen2881 = True
								} else {
									gen2881 = False
								}

							} else {
								gen2881 = False
							}
							if True == gen2881 {
								gen2886 = True
							} else {
								gen2886 = False
							}

						} else {
							gen2886 = False
						}
						if True == gen2886 {
							gen2890 = True
						} else {
							gen2890 = False
						}

					} else {
						gen2890 = False
					}
					if True == gen2890 {
						gen2893 = True
					} else {
						gen2893 = False
					}

				} else {
					gen2893 = False
				}
				if True == gen2893 {
					gen2895 = True
				} else {
					gen2895 = False
				}

			} else {
				gen2895 = False
			}
			if True == gen2895 {
				gen2836 := Call(__e, ShenFunc(symhd), V5007)

				gen2837 := Call(__e, ShenFunc(symhd), gen2836)

				gen2838 := Call(__e, ShenFunc(symtl), gen2837)

				gen2839 := Call(__e, ShenFunc(symtl), gen2838)

				gen2840 := Call(__e, ShenFunc(symhd), gen2839)

				gen2841 := Call(__e, ShenFunc(symhd), V5007)

				gen2842 := Call(__e, ShenFunc(symtl), gen2841)

				gen2843 := Call(__e, ShenFunc(symcons), gen2840, gen2842)

				gen2844 := Call(__e, ShenFunc(symhd), V5007)

				gen2845 := Call(__e, ShenFunc(symhd), gen2844)

				gen2846 := Call(__e, ShenFunc(symtl), gen2845)

				gen2847 := Call(__e, ShenFunc(symhd), gen2846)

				gen2848 := Call(__e, ShenFunc(symtl), V5007)

				gen2849 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4true_1branch), gen2847, gen2848)

				__e.TailApply(ShenFunc(symcons), gen2843, gen2849)

				return

			} else {
				gen2834 := Call(__e, ShenFunc(symcons_2), V5007)

				var gen2835 Obj
				if True == gen2834 {
					gen2831 := Call(__e, ShenFunc(symhd), V5007)

					gen2832 := Call(__e, ShenFunc(symcons_2), gen2831)

					var gen2833 Obj
					if True == gen2832 {
						gen2827 := Call(__e, ShenFunc(symhd), V5007)

						gen2828 := Call(__e, ShenFunc(symtl), gen2827)

						gen2829 := Call(__e, ShenFunc(symcons_2), gen2828)

						var gen2830 Obj
						if True == gen2829 {
							gen2822 := Call(__e, ShenFunc(symhd), V5007)

							gen2823 := Call(__e, ShenFunc(symtl), gen2822)

							gen2824 := Call(__e, ShenFunc(symtl), gen2823)

							gen2825 := Call(__e, ShenFunc(sym_a), Nil, gen2824)

							var gen2826 Obj
							if True == gen2825 {
								gen2819 := Call(__e, ShenFunc(symhd), V5007)

								gen2820 := Call(__e, ShenFunc(symhd), gen2819)

								gen2821 := Call(__e, ShenFunc(sym_a), gen2820, V5006)

								if True == gen2821 {
									gen2826 = True
								} else {
									gen2826 = False
								}

							} else {
								gen2826 = False
							}
							if True == gen2826 {
								gen2830 = True
							} else {
								gen2830 = False
							}

						} else {
							gen2830 = False
						}
						if True == gen2830 {
							gen2833 = True
						} else {
							gen2833 = False
						}

					} else {
						gen2833 = False
					}
					if True == gen2833 {
						gen2835 = True
					} else {
						gen2835 = False
					}

				} else {
					gen2835 = False
				}
				if True == gen2835 {
					gen2816 := Call(__e, ShenFunc(symhd), V5007)

					gen2817 := Call(__e, ShenFunc(symtl), gen2816)

					gen2818 := Call(__e, ShenFunc(symcons), True, gen2817)

					__e.TailApply(ShenFunc(symcons), gen2818, Nil)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.true-branch"), gen2896)

		gen2968 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5016 := __args[0]
			_ = V5016
			V5017 := __args[1]
			_ = V5017
			gen2966 := Call(__e, ShenFunc(symcons_2), V5017)

			var gen2967 Obj
			if True == gen2966 {
				gen2963 := Call(__e, ShenFunc(symhd), V5017)

				gen2964 := Call(__e, ShenFunc(symcons_2), gen2963)

				var gen2965 Obj
				if True == gen2964 {
					gen2959 := Call(__e, ShenFunc(symhd), V5017)

					gen2960 := Call(__e, ShenFunc(symhd), gen2959)

					gen2961 := Call(__e, ShenFunc(symcons_2), gen2960)

					var gen2962 Obj
					if True == gen2961 {
						gen2954 := Call(__e, ShenFunc(symhd), V5017)

						gen2955 := Call(__e, ShenFunc(symhd), gen2954)

						gen2956 := Call(__e, ShenFunc(symhd), gen2955)

						gen2957 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen2956)

						var gen2958 Obj
						if True == gen2957 {
							gen2949 := Call(__e, ShenFunc(symhd), V5017)

							gen2950 := Call(__e, ShenFunc(symhd), gen2949)

							gen2951 := Call(__e, ShenFunc(symtl), gen2950)

							gen2952 := Call(__e, ShenFunc(symcons_2), gen2951)

							var gen2953 Obj
							if True == gen2952 {
								gen2943 := Call(__e, ShenFunc(symhd), V5017)

								gen2944 := Call(__e, ShenFunc(symhd), gen2943)

								gen2945 := Call(__e, ShenFunc(symtl), gen2944)

								gen2946 := Call(__e, ShenFunc(symtl), gen2945)

								gen2947 := Call(__e, ShenFunc(symcons_2), gen2946)

								var gen2948 Obj
								if True == gen2947 {
									gen2936 := Call(__e, ShenFunc(symhd), V5017)

									gen2937 := Call(__e, ShenFunc(symhd), gen2936)

									gen2938 := Call(__e, ShenFunc(symtl), gen2937)

									gen2939 := Call(__e, ShenFunc(symtl), gen2938)

									gen2940 := Call(__e, ShenFunc(symtl), gen2939)

									gen2941 := Call(__e, ShenFunc(sym_a), Nil, gen2940)

									var gen2942 Obj
									if True == gen2941 {
										gen2932 := Call(__e, ShenFunc(symhd), V5017)

										gen2933 := Call(__e, ShenFunc(symtl), gen2932)

										gen2934 := Call(__e, ShenFunc(symcons_2), gen2933)

										var gen2935 Obj
										if True == gen2934 {
											gen2927 := Call(__e, ShenFunc(symhd), V5017)

											gen2928 := Call(__e, ShenFunc(symtl), gen2927)

											gen2929 := Call(__e, ShenFunc(symtl), gen2928)

											gen2930 := Call(__e, ShenFunc(sym_a), Nil, gen2929)

											var gen2931 Obj
											if True == gen2930 {
												gen2922 := Call(__e, ShenFunc(symhd), V5017)

												gen2923 := Call(__e, ShenFunc(symhd), gen2922)

												gen2924 := Call(__e, ShenFunc(symtl), gen2923)

												gen2925 := Call(__e, ShenFunc(symhd), gen2924)

												gen2926 := Call(__e, ShenFunc(sym_a), gen2925, V5016)

												if True == gen2926 {
													gen2931 = True
												} else {
													gen2931 = False
												}

											} else {
												gen2931 = False
											}
											if True == gen2931 {
												gen2935 = True
											} else {
												gen2935 = False
											}

										} else {
											gen2935 = False
										}
										if True == gen2935 {
											gen2942 = True
										} else {
											gen2942 = False
										}

									} else {
										gen2942 = False
									}
									if True == gen2942 {
										gen2948 = True
									} else {
										gen2948 = False
									}

								} else {
									gen2948 = False
								}
								if True == gen2948 {
									gen2953 = True
								} else {
									gen2953 = False
								}

							} else {
								gen2953 = False
							}
							if True == gen2953 {
								gen2958 = True
							} else {
								gen2958 = False
							}

						} else {
							gen2958 = False
						}
						if True == gen2958 {
							gen2962 = True
						} else {
							gen2962 = False
						}

					} else {
						gen2962 = False
					}
					if True == gen2962 {
						gen2965 = True
					} else {
						gen2965 = False
					}

				} else {
					gen2965 = False
				}
				if True == gen2965 {
					gen2967 = True
				} else {
					gen2967 = False
				}

			} else {
				gen2967 = False
			}
			if True == gen2967 {
				gen2917 := Call(__e, ShenFunc(symhd), V5017)

				gen2918 := Call(__e, ShenFunc(symhd), gen2917)

				gen2919 := Call(__e, ShenFunc(symtl), gen2918)

				gen2920 := Call(__e, ShenFunc(symhd), gen2919)

				gen2921 := Call(__e, ShenFunc(symtl), V5017)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4false_1branch), gen2920, gen2921)

				return

			} else {
				gen2915 := Call(__e, ShenFunc(symcons_2), V5017)

				var gen2916 Obj
				if True == gen2915 {
					gen2912 := Call(__e, ShenFunc(symhd), V5017)

					gen2913 := Call(__e, ShenFunc(symcons_2), gen2912)

					var gen2914 Obj
					if True == gen2913 {
						gen2908 := Call(__e, ShenFunc(symhd), V5017)

						gen2909 := Call(__e, ShenFunc(symtl), gen2908)

						gen2910 := Call(__e, ShenFunc(symcons_2), gen2909)

						var gen2911 Obj
						if True == gen2910 {
							gen2903 := Call(__e, ShenFunc(symhd), V5017)

							gen2904 := Call(__e, ShenFunc(symtl), gen2903)

							gen2905 := Call(__e, ShenFunc(symtl), gen2904)

							gen2906 := Call(__e, ShenFunc(sym_a), Nil, gen2905)

							var gen2907 Obj
							if True == gen2906 {
								gen2900 := Call(__e, ShenFunc(symhd), V5017)

								gen2901 := Call(__e, ShenFunc(symhd), gen2900)

								gen2902 := Call(__e, ShenFunc(sym_a), gen2901, V5016)

								if True == gen2902 {
									gen2907 = True
								} else {
									gen2907 = False
								}

							} else {
								gen2907 = False
							}
							if True == gen2907 {
								gen2911 = True
							} else {
								gen2911 = False
							}

						} else {
							gen2911 = False
						}
						if True == gen2911 {
							gen2914 = True
						} else {
							gen2914 = False
						}

					} else {
						gen2914 = False
					}
					if True == gen2914 {
						gen2916 = True
					} else {
						gen2916 = False
					}

				} else {
					gen2916 = False
				}
				if True == gen2916 {
					gen2897 := Call(__e, ShenFunc(symhd), V5017)

					gen2898 := Call(__e, ShenFunc(symhd), gen2897)

					gen2899 := Call(__e, ShenFunc(symtl), V5017)

					__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4false_1branch), gen2898, gen2899)

					return

				} else {
					__e.Return(V5017)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.false-branch"), gen2968)

		gen3010 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5020 := __args[0]
			_ = V5020
			V5021 := __args[1]
			_ = V5021
			gen3008 := Call(__e, ShenFunc(symcons_2), V5020)

			var gen3009 Obj
			if True == gen3008 {
				gen3005 := Call(__e, ShenFunc(symhd), V5020)

				gen3006 := Call(__e, ShenFunc(sym_a), MakeSymbol("%%return"), gen3005)

				var gen3007 Obj
				if True == gen3006 {
					gen3002 := Call(__e, ShenFunc(symtl), V5020)

					gen3003 := Call(__e, ShenFunc(symcons_2), gen3002)

					var gen3004 Obj
					if True == gen3003 {
						gen2998 := Call(__e, ShenFunc(symtl), V5020)

						gen2999 := Call(__e, ShenFunc(symtl), gen2998)

						gen3000 := Call(__e, ShenFunc(sym_a), Nil, gen2999)

						var gen3001 Obj
						if True == gen3000 {
							gen2994 := Call(__e, ShenFunc(symtl), V5020)

							gen2995 := Call(__e, ShenFunc(symhd), gen2994)

							gen2996 := Call(__e, ShenFunc(symcons_2), gen2995)

							gen2997 := Call(__e, ShenFunc(symnot), gen2996)

							if True == gen2997 {
								gen3001 = True
							} else {
								gen3001 = False
							}

						} else {
							gen3001 = False
						}
						if True == gen3001 {
							gen3004 = True
						} else {
							gen3004 = False
						}

					} else {
						gen3004 = False
					}
					if True == gen3004 {
						gen3007 = True
					} else {
						gen3007 = False
					}

				} else {
					gen3007 = False
				}
				if True == gen3007 {
					gen3009 = True
				} else {
					gen3009 = False
				}

			} else {
				gen3009 = False
			}
			if True == gen3009 {
				__e.TailApply(V5021, V5020)

				return
			} else {
				gen2992 := Call(__e, ShenFunc(symcons_2), V5020)

				var gen2993 Obj
				if True == gen2992 {
					gen2989 := Call(__e, ShenFunc(symhd), V5020)

					gen2990 := Call(__e, ShenFunc(sym_a), MakeSymbol("fail"), gen2989)

					var gen2991 Obj
					if True == gen2990 {
						gen2987 := Call(__e, ShenFunc(symtl), V5020)

						gen2988 := Call(__e, ShenFunc(sym_a), Nil, gen2987)

						if True == gen2988 {
							gen2991 = True
						} else {
							gen2991 = False
						}

					} else {
						gen2991 = False
					}
					if True == gen2991 {
						gen2993 = True
					} else {
						gen2993 = False
					}

				} else {
					gen2993 = False
				}
				if True == gen2993 {
					__e.TailApply(V5021, V5020)

					return
				} else {
					gen2985 := Call(__e, ShenFunc(symcons_2), V5020)

					var gen2986 Obj
					if True == gen2985 {
						gen2982 := Call(__e, ShenFunc(symhd), V5020)

						gen2983 := Call(__e, ShenFunc(sym_a), MakeSymbol("%%goto-label"), gen2982)

						var gen2984 Obj
						if True == gen2983 {
							gen2979 := Call(__e, ShenFunc(symtl), V5020)

							gen2980 := Call(__e, ShenFunc(symcons_2), gen2979)

							var gen2981 Obj
							if True == gen2980 {
								gen2976 := Call(__e, ShenFunc(symtl), V5020)

								gen2977 := Call(__e, ShenFunc(symtl), gen2976)

								gen2978 := Call(__e, ShenFunc(sym_a), Nil, gen2977)

								if True == gen2978 {
									gen2981 = True
								} else {
									gen2981 = False
								}

							} else {
								gen2981 = False
							}
							if True == gen2981 {
								gen2984 = True
							} else {
								gen2984 = False
							}

						} else {
							gen2984 = False
						}
						if True == gen2984 {
							gen2986 = True
						} else {
							gen2986 = False
						}

					} else {
						gen2986 = False
					}
					if True == gen2986 {
						__e.TailApply(V5021, V5020)

						return
					} else {
						gen2969 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4generate_1label))

						Label := gen2969
						gen2970 := Call(__e, ShenFunc(symcons), Label, Nil)

						gen2971 := Call(__e, ShenFunc(symcons), MakeSymbol("%%goto-label"), gen2970)

						gen2972 := Call(__e, V5021, gen2971)

						gen2973 := Call(__e, ShenFunc(symcons), gen2972, Nil)

						gen2974 := Call(__e, ShenFunc(symcons), V5020, gen2973)

						gen2975 := Call(__e, ShenFunc(symcons), Label, gen2974)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("%%let-label"), gen2975)

						return

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.with-labelled-else"), gen3010)

		gen3108 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5024 := __args[0]
			_ = V5024
			gen3106 := Call(__e, ShenFunc(symcons_2), V5024)

			var gen3107 Obj
			if True == gen3106 {
				gen3103 := Call(__e, ShenFunc(symhd), V5024)

				gen3104 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen3103)

				var gen3105 Obj
				if True == gen3104 {
					gen3100 := Call(__e, ShenFunc(symtl), V5024)

					gen3101 := Call(__e, ShenFunc(symcons_2), gen3100)

					var gen3102 Obj
					if True == gen3101 {
						gen3096 := Call(__e, ShenFunc(symtl), V5024)

						gen3097 := Call(__e, ShenFunc(symtl), gen3096)

						gen3098 := Call(__e, ShenFunc(symcons_2), gen3097)

						var gen3099 Obj
						if True == gen3098 {
							gen3091 := Call(__e, ShenFunc(symtl), V5024)

							gen3092 := Call(__e, ShenFunc(symtl), gen3091)

							gen3093 := Call(__e, ShenFunc(symhd), gen3092)

							gen3094 := Call(__e, ShenFunc(symcons_2), gen3093)

							var gen3095 Obj
							if True == gen3094 {
								gen3085 := Call(__e, ShenFunc(symtl), V5024)

								gen3086 := Call(__e, ShenFunc(symtl), gen3085)

								gen3087 := Call(__e, ShenFunc(symhd), gen3086)

								gen3088 := Call(__e, ShenFunc(symhd), gen3087)

								gen3089 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen3088)

								var gen3090 Obj
								if True == gen3089 {
									gen3079 := Call(__e, ShenFunc(symtl), V5024)

									gen3080 := Call(__e, ShenFunc(symtl), gen3079)

									gen3081 := Call(__e, ShenFunc(symhd), gen3080)

									gen3082 := Call(__e, ShenFunc(symtl), gen3081)

									gen3083 := Call(__e, ShenFunc(symcons_2), gen3082)

									var gen3084 Obj
									if True == gen3083 {
										gen3072 := Call(__e, ShenFunc(symtl), V5024)

										gen3073 := Call(__e, ShenFunc(symtl), gen3072)

										gen3074 := Call(__e, ShenFunc(symhd), gen3073)

										gen3075 := Call(__e, ShenFunc(symtl), gen3074)

										gen3076 := Call(__e, ShenFunc(symtl), gen3075)

										gen3077 := Call(__e, ShenFunc(symcons_2), gen3076)

										var gen3078 Obj
										if True == gen3077 {
											gen3064 := Call(__e, ShenFunc(symtl), V5024)

											gen3065 := Call(__e, ShenFunc(symtl), gen3064)

											gen3066 := Call(__e, ShenFunc(symhd), gen3065)

											gen3067 := Call(__e, ShenFunc(symtl), gen3066)

											gen3068 := Call(__e, ShenFunc(symtl), gen3067)

											gen3069 := Call(__e, ShenFunc(symtl), gen3068)

											gen3070 := Call(__e, ShenFunc(symcons_2), gen3069)

											var gen3071 Obj
											if True == gen3070 {
												gen3055 := Call(__e, ShenFunc(symtl), V5024)

												gen3056 := Call(__e, ShenFunc(symtl), gen3055)

												gen3057 := Call(__e, ShenFunc(symhd), gen3056)

												gen3058 := Call(__e, ShenFunc(symtl), gen3057)

												gen3059 := Call(__e, ShenFunc(symtl), gen3058)

												gen3060 := Call(__e, ShenFunc(symtl), gen3059)

												gen3061 := Call(__e, ShenFunc(symtl), gen3060)

												gen3062 := Call(__e, ShenFunc(sym_a), Nil, gen3061)

												var gen3063 Obj
												if True == gen3062 {
													gen3050 := Call(__e, ShenFunc(symtl), V5024)

													gen3051 := Call(__e, ShenFunc(symtl), gen3050)

													gen3052 := Call(__e, ShenFunc(symtl), gen3051)

													gen3053 := Call(__e, ShenFunc(symcons_2), gen3052)

													var gen3054 Obj
													if True == gen3053 {
														gen3044 := Call(__e, ShenFunc(symtl), V5024)

														gen3045 := Call(__e, ShenFunc(symtl), gen3044)

														gen3046 := Call(__e, ShenFunc(symtl), gen3045)

														gen3047 := Call(__e, ShenFunc(symtl), gen3046)

														gen3048 := Call(__e, ShenFunc(sym_a), Nil, gen3047)

														var gen3049 Obj
														if True == gen3048 {
															gen3032 := Call(__e, ShenFunc(symtl), V5024)

															gen3033 := Call(__e, ShenFunc(symtl), gen3032)

															gen3034 := Call(__e, ShenFunc(symtl), gen3033)

															gen3035 := Call(__e, ShenFunc(symhd), gen3034)

															gen3036 := Call(__e, ShenFunc(symtl), V5024)

															gen3037 := Call(__e, ShenFunc(symtl), gen3036)

															gen3038 := Call(__e, ShenFunc(symhd), gen3037)

															gen3039 := Call(__e, ShenFunc(symtl), gen3038)

															gen3040 := Call(__e, ShenFunc(symtl), gen3039)

															gen3041 := Call(__e, ShenFunc(symtl), gen3040)

															gen3042 := Call(__e, ShenFunc(symhd), gen3041)

															gen3043 := Call(__e, ShenFunc(sym_a), gen3035, gen3042)

															if True == gen3043 {
																gen3049 = True
															} else {
																gen3049 = False
															}

														} else {
															gen3049 = False
														}
														if True == gen3049 {
															gen3054 = True
														} else {
															gen3054 = False
														}

													} else {
														gen3054 = False
													}
													if True == gen3054 {
														gen3063 = True
													} else {
														gen3063 = False
													}

												} else {
													gen3063 = False
												}
												if True == gen3063 {
													gen3071 = True
												} else {
													gen3071 = False
												}

											} else {
												gen3071 = False
											}
											if True == gen3071 {
												gen3078 = True
											} else {
												gen3078 = False
											}

										} else {
											gen3078 = False
										}
										if True == gen3078 {
											gen3084 = True
										} else {
											gen3084 = False
										}

									} else {
										gen3084 = False
									}
									if True == gen3084 {
										gen3090 = True
									} else {
										gen3090 = False
									}

								} else {
									gen3090 = False
								}
								if True == gen3090 {
									gen3095 = True
								} else {
									gen3095 = False
								}

							} else {
								gen3095 = False
							}
							if True == gen3095 {
								gen3099 = True
							} else {
								gen3099 = False
							}

						} else {
							gen3099 = False
						}
						if True == gen3099 {
							gen3102 = True
						} else {
							gen3102 = False
						}

					} else {
						gen3102 = False
					}
					if True == gen3102 {
						gen3105 = True
					} else {
						gen3105 = False
					}

				} else {
					gen3105 = False
				}
				if True == gen3105 {
					gen3107 = True
				} else {
					gen3107 = False
				}

			} else {
				gen3107 = False
			}
			if True == gen3107 {
				gen3011 := Call(__e, ShenFunc(symtl), V5024)

				gen3012 := Call(__e, ShenFunc(symhd), gen3011)

				gen3013 := Call(__e, ShenFunc(symtl), V5024)

				gen3014 := Call(__e, ShenFunc(symtl), gen3013)

				gen3015 := Call(__e, ShenFunc(symhd), gen3014)

				gen3016 := Call(__e, ShenFunc(symtl), gen3015)

				gen3017 := Call(__e, ShenFunc(symhd), gen3016)

				gen3018 := Call(__e, ShenFunc(symcons), gen3017, Nil)

				gen3019 := Call(__e, ShenFunc(symcons), gen3012, gen3018)

				gen3020 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen3019)

				gen3021 := Call(__e, ShenFunc(symtl), V5024)

				gen3022 := Call(__e, ShenFunc(symtl), gen3021)

				gen3023 := Call(__e, ShenFunc(symhd), gen3022)

				gen3024 := Call(__e, ShenFunc(symtl), gen3023)

				gen3025 := Call(__e, ShenFunc(symtl), gen3024)

				gen3026 := Call(__e, ShenFunc(symhd), gen3025)

				gen3027 := Call(__e, ShenFunc(symtl), V5024)

				gen3028 := Call(__e, ShenFunc(symtl), gen3027)

				gen3029 := Call(__e, ShenFunc(symtl), gen3028)

				gen3030 := Call(__e, ShenFunc(symcons), gen3026, gen3029)

				gen3031 := Call(__e, ShenFunc(symcons), gen3020, gen3030)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen3031)

				return

			} else {
				__e.Return(V5024)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs"), gen3108)

		gen3110 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5027 := __args[0]
			_ = V5027
			V5028 := __args[1]
			_ = V5028
			gen3109 := Call(__e, ShenFunc(symconcat), MakeSymbol("/"), V5028)

			__e.TailApply(ShenFunc(symconcat), V5027, gen3109)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.concat/"), gen3110)

		gen3128 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5032 := __args[0]
			_ = V5032
			gen3126 := Call(__e, ShenFunc(symcons_2), V5032)

			var gen3127 Obj
			if True == gen3126 {
				gen3123 := Call(__e, ShenFunc(symtl), V5032)

				gen3124 := Call(__e, ShenFunc(symcons_2), gen3123)

				var gen3125 Obj
				if True == gen3124 {
					gen3119 := Call(__e, ShenFunc(symtl), V5032)

					gen3120 := Call(__e, ShenFunc(symtl), gen3119)

					gen3121 := Call(__e, ShenFunc(sym_a), Nil, gen3120)

					var gen3122 Obj
					if True == gen3121 {
						gen3117 := Call(__e, ShenFunc(symhd), V5032)

						gen3118 := Call(__e, ShenFunc(symsymbol_2), gen3117)

						if True == gen3118 {
							gen3122 = True
						} else {
							gen3122 = False
						}

					} else {
						gen3122 = False
					}
					if True == gen3122 {
						gen3125 = True
					} else {
						gen3125 = False
					}

				} else {
					gen3125 = False
				}
				if True == gen3125 {
					gen3127 = True
				} else {
					gen3127 = False
				}

			} else {
				gen3127 = False
			}
			if True == gen3127 {
				gen3113 := Call(__e, ShenFunc(symtl), V5032)

				gen3114 := Call(__e, ShenFunc(symhd), gen3113)

				gen3115 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4exp_1var), gen3114)

				gen3116 := Call(__e, ShenFunc(symhd), V5032)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4concat_c), gen3115, gen3116)

				return

			} else {
				gen3112 := Call(__e, ShenFunc(symcons_2), V5032)

				if True == gen3112 {
					gen3111 := Call(__e, ShenFunc(symhd), V5032)

					__e.TailApply(ShenFunc(symgensym), gen3111)

					return

				} else {
					__e.Return(V5032)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.exp-var"), gen3128)

		gen3130 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5035 := __args[0]
			_ = V5035
			V5036 := __args[1]
			_ = V5036
			gen3129 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4test_1_6selectors), V5035)

			__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4bind_1repeating_1selectors), gen3129, V5036)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.optimize-selectors"), gen3130)

		gen3199 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5042 := __args[0]
			_ = V5042
			gen3197 := Call(__e, ShenFunc(symcons_2), V5042)

			var gen3198 Obj
			if True == gen3197 {
				gen3194 := Call(__e, ShenFunc(symhd), V5042)

				gen3195 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons?"), gen3194)

				var gen3196 Obj
				if True == gen3195 {
					gen3191 := Call(__e, ShenFunc(symtl), V5042)

					gen3192 := Call(__e, ShenFunc(symcons_2), gen3191)

					var gen3193 Obj
					if True == gen3192 {
						gen3188 := Call(__e, ShenFunc(symtl), V5042)

						gen3189 := Call(__e, ShenFunc(symtl), gen3188)

						gen3190 := Call(__e, ShenFunc(sym_a), Nil, gen3189)

						if True == gen3190 {
							gen3193 = True
						} else {
							gen3193 = False
						}

					} else {
						gen3193 = False
					}
					if True == gen3193 {
						gen3196 = True
					} else {
						gen3196 = False
					}

				} else {
					gen3196 = False
				}
				if True == gen3196 {
					gen3198 = True
				} else {
					gen3198 = False
				}

			} else {
				gen3198 = False
			}
			if True == gen3198 {
				gen3183 := Call(__e, ShenFunc(symtl), V5042)

				gen3184 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen3183)

				gen3185 := Call(__e, ShenFunc(symtl), V5042)

				gen3186 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen3185)

				gen3187 := Call(__e, ShenFunc(symcons), gen3186, Nil)

				__e.TailApply(ShenFunc(symcons), gen3184, gen3187)

				return

			} else {
				gen3181 := Call(__e, ShenFunc(symcons_2), V5042)

				var gen3182 Obj
				if True == gen3181 {
					gen3178 := Call(__e, ShenFunc(symhd), V5042)

					gen3179 := Call(__e, ShenFunc(sym_a), MakeSymbol("tuple?"), gen3178)

					var gen3180 Obj
					if True == gen3179 {
						gen3175 := Call(__e, ShenFunc(symtl), V5042)

						gen3176 := Call(__e, ShenFunc(symcons_2), gen3175)

						var gen3177 Obj
						if True == gen3176 {
							gen3172 := Call(__e, ShenFunc(symtl), V5042)

							gen3173 := Call(__e, ShenFunc(symtl), gen3172)

							gen3174 := Call(__e, ShenFunc(sym_a), Nil, gen3173)

							if True == gen3174 {
								gen3177 = True
							} else {
								gen3177 = False
							}

						} else {
							gen3177 = False
						}
						if True == gen3177 {
							gen3180 = True
						} else {
							gen3180 = False
						}

					} else {
						gen3180 = False
					}
					if True == gen3180 {
						gen3182 = True
					} else {
						gen3182 = False
					}

				} else {
					gen3182 = False
				}
				if True == gen3182 {
					gen3167 := Call(__e, ShenFunc(symtl), V5042)

					gen3168 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen3167)

					gen3169 := Call(__e, ShenFunc(symtl), V5042)

					gen3170 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen3169)

					gen3171 := Call(__e, ShenFunc(symcons), gen3170, Nil)

					__e.TailApply(ShenFunc(symcons), gen3168, gen3171)

					return

				} else {
					gen3165 := Call(__e, ShenFunc(symcons_2), V5042)

					var gen3166 Obj
					if True == gen3165 {
						gen3162 := Call(__e, ShenFunc(symhd), V5042)

						gen3163 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.+string?"), gen3162)

						var gen3164 Obj
						if True == gen3163 {
							gen3159 := Call(__e, ShenFunc(symtl), V5042)

							gen3160 := Call(__e, ShenFunc(symcons_2), gen3159)

							var gen3161 Obj
							if True == gen3160 {
								gen3156 := Call(__e, ShenFunc(symtl), V5042)

								gen3157 := Call(__e, ShenFunc(symtl), gen3156)

								gen3158 := Call(__e, ShenFunc(sym_a), Nil, gen3157)

								if True == gen3158 {
									gen3161 = True
								} else {
									gen3161 = False
								}

							} else {
								gen3161 = False
							}
							if True == gen3161 {
								gen3164 = True
							} else {
								gen3164 = False
							}

						} else {
							gen3164 = False
						}
						if True == gen3164 {
							gen3166 = True
						} else {
							gen3166 = False
						}

					} else {
						gen3166 = False
					}
					if True == gen3166 {
						gen3151 := Call(__e, ShenFunc(symtl), V5042)

						gen3152 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen3151)

						gen3153 := Call(__e, ShenFunc(symtl), V5042)

						gen3154 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen3153)

						gen3155 := Call(__e, ShenFunc(symcons), gen3154, Nil)

						__e.TailApply(ShenFunc(symcons), gen3152, gen3155)

						return

					} else {
						gen3149 := Call(__e, ShenFunc(symcons_2), V5042)

						var gen3150 Obj
						if True == gen3149 {
							gen3146 := Call(__e, ShenFunc(symhd), V5042)

							gen3147 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.+vector?"), gen3146)

							var gen3148 Obj
							if True == gen3147 {
								gen3143 := Call(__e, ShenFunc(symtl), V5042)

								gen3144 := Call(__e, ShenFunc(symcons_2), gen3143)

								var gen3145 Obj
								if True == gen3144 {
									gen3140 := Call(__e, ShenFunc(symtl), V5042)

									gen3141 := Call(__e, ShenFunc(symtl), gen3140)

									gen3142 := Call(__e, ShenFunc(sym_a), Nil, gen3141)

									if True == gen3142 {
										gen3145 = True
									} else {
										gen3145 = False
									}

								} else {
									gen3145 = False
								}
								if True == gen3145 {
									gen3148 = True
								} else {
									gen3148 = False
								}

							} else {
								gen3148 = False
							}
							if True == gen3148 {
								gen3150 = True
							} else {
								gen3150 = False
							}

						} else {
							gen3150 = False
						}
						if True == gen3150 {
							gen3135 := Call(__e, ShenFunc(symtl), V5042)

							gen3136 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen3135)

							gen3137 := Call(__e, ShenFunc(symtl), V5042)

							gen3138 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen3137)

							gen3139 := Call(__e, ShenFunc(symcons), gen3138, Nil)

							__e.TailApply(ShenFunc(symcons), gen3136, gen3139)

							return

						} else {
							gen3131 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

							gen3132 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4apply_1selector_1handlers), gen3131, V5042)

							Result := gen3132
							gen3133 := Call(__e, ShenFunc(symfail))

							gen3134 := Call(__e, ShenFunc(sym_a), Result, gen3133)

							if True == gen3134 {
								__e.Return(Nil)
								return
							} else {
								__e.Return(Result)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.test->selectors"), gen3199)

		gen3205 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5045 := __args[0]
			_ = V5045
			V5046 := __args[1]
			_ = V5046
			gen3204 := Call(__e, ShenFunc(symcons_2), V5045)

			if True == gen3204 {
				gen3201 := Call(__e, ShenFunc(symhd), V5045)

				gen3202 := Call(__e, ShenFunc(symtl), V5045)

				gen3203 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4bind_1repeating_1selectors), gen3202, V5046)

				__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4bind_1selector), gen3201, gen3203)

				return

			} else {
				gen3200 := Call(__e, ShenFunc(sym_a), Nil, V5045)

				if True == gen3200 {
					__e.Return(V5046)
					return
				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors"), gen3205)

		gen3213 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5053 := __args[0]
			_ = V5053
			V5054 := __args[1]
			_ = V5054
			gen3211 := Call(__e, ShenFunc(symoccurrences), V5053, V5054)

			gen3212 := Call(__e, ShenFunc(sym_6), gen3211, MakeNumber(1))

			if True == gen3212 {
				gen3206 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4exp_1var), V5053)

				Var := gen3206
				gen3207 := Call(__e, ShenFunc(symsubst), Var, V5053, V5054)

				gen3208 := Call(__e, ShenFunc(symcons), gen3207, Nil)

				gen3209 := Call(__e, ShenFunc(symcons), V5053, gen3208)

				gen3210 := Call(__e, ShenFunc(symcons), Var, gen3209)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen3210)

				return

			} else {
				__e.Return(V5054)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.bind-selector"), gen3213)

		gen3223 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5067 := __args[0]
			_ = V5067
			V5068 := __args[1]
			_ = V5068
			gen3222 := Call(__e, ShenFunc(sym_a), Nil, V5067)

			if True == gen3222 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				gen3216 := MakeNative(func(__e Evaluator, __args ...Obj) {
					gen3215 := Call(__e, ShenFunc(symcons_2), V5067)

					if True == gen3215 {
						gen3214 := Call(__e, ShenFunc(symtl), V5067)

						__e.TailApply(ShenFunc(symshen_4x_4factorise_1defun_4apply_1selector_1handlers), gen3214, V5068)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.apply-selector-handlers"))

						return
					}

				}, 0)
				Freeze := gen3216
				gen3221 := Call(__e, ShenFunc(symcons_2), V5067)

				if True == gen3221 {
					gen3217 := Call(__e, ShenFunc(symhd), V5067)

					f30 := gen3217
					gen3218 := Call(__e, f30, V5068)

					Result := gen3218
					gen3219 := Call(__e, ShenFunc(symfail))

					gen3220 := Call(__e, ShenFunc(sym_a), Result, gen3219)

					if True == gen3220 {
						__e.TailApply(ShenFunc(symthaw), Freeze)

						return
					} else {
						__e.Return(Result)
						return
					}

				} else {
					__e.TailApply(ShenFunc(symthaw), Freeze)

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.apply-selector-handlers"), gen3223)

		gen3224 := MakeNative(func(__e Evaluator, __args ...Obj) {
			Call(__e, ShenFunc(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"), Nil)
			__e.Return(MakeSymbol("shen.x.factorise-defun.done"))
			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.initialise"), gen3224)

		gen3232 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5070 := __args[0]
			_ = V5070
			gen3230 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

			gen3231 := Call(__e, ShenFunc(symelement_2), V5070, gen3230)

			if True == gen3231 {
				__e.Return(V5070)
				return
			} else {
				gen3225 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

				gen3226 := Call(__e, ShenFunc(symcons), V5070, gen3225)

				Call(__e, ShenFunc(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"), gen3226)

				gen3227 := Call(__e, ShenFunc(symfunction), V5070)

				gen3228 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

				gen3229 := Call(__e, ShenFunc(symcons), gen3227, gen3228)

				Call(__e, ShenFunc(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"), gen3229)

				__e.Return(V5070)
				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.register-selector-handler"), gen3232)

		gen3236 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5073 := __args[0]
			_ = V5073
			V5074 := __args[1]
			_ = V5074
			gen3234 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__ := __args[0]
				_ = __
				gen3233 := Call(__e, ShenFunc(symshen_4app), V5073, MakeString(" is not a selector handler\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen3233)

				return

			}, 1)
			gen3235 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4findpos), V5073, V5074)

				return
			}, 0)
			__e.Return(Try(__e, gen3235).Catch(gen3234))
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.findpos"), gen3236)

		gen3244 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V5076 := __args[0]
			_ = V5076
			gen3237 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"))

			Reg := gen3237
			gen3238 := Call(__e, ShenFunc(symshen_4x_4factorise_1defun_4findpos), V5076, Reg)

			Pos := gen3238
			gen3239 := Call(__e, ShenFunc(symremove), V5076, Reg)

			gen3240 := Call(__e, ShenFunc(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"), gen3239)

			RemoveReg := gen3240
			_ = RemoveReg
			gen3241 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

			gen3242 := Call(__e, ShenFunc(symshen_4remove_1nth), Pos, gen3241)

			gen3243 := Call(__e, ShenFunc(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"), gen3242)

			RemoveFun := gen3243
			_ = RemoveFun
			__e.Return(V5076)
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.x.factorise-defun.unregister-selector-handler"), gen3244)

		return

	}, 0))
}

var symport = MakeSymbol("port")
var symshen_4case_1form = MakeSymbol("shen.case-form")
var symshen_4type_1signature_1of_1internal = MakeSymbol("shen.type-signature-of-internal")
var symshen_4type_1signature_1of_1port = MakeSymbol("shen.type-signature-of-port")
var symshen_4x_4factorise_1defun_4exp_1var = MakeSymbol("shen.x.factorise-defun.exp-var")
var symshen_4_5conclusion_6 = MakeSymbol("shen.<conclusion>")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symbind = MakeSymbol("bind")
var symtc = MakeSymbol("tc")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4placeholders = MakeSymbol("shen.placeholders")
var symshen_4type_1signature_1of_1gensym = MakeSymbol("shen.type-signature-of-gensym")
var symshen_4typecheck_1and_1evaluate = MakeSymbol("shen.typecheck-and-evaluate")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4em__help = MakeSymbol("shen.em_help")
var symshen_4ephemeral__variable_2 = MakeSymbol("shen.ephemeral_variable?")
var symread_1file = MakeSymbol("read-file")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4flatten = MakeSymbol("shen.flatten")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4type_1signature_1of_1_a = MakeSymbol("shen.type-signature-of-=")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4aah = MakeSymbol("shen.aah")
var symshen_4double_1_6singles = MakeSymbol("shen.double->singles")
var sym_c = MakeSymbol("/")
var symshen_4type_1signature_1of_1_5 = MakeSymbol("shen.type-signature-of-<")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4strip_1protect = MakeSymbol("shen.strip-protect")
var symconcat = MakeSymbol("concat")
var symshen_4compress_150 = MakeSymbol("shen.compress-50")
var symshen_4newhyps = MakeSymbol("shen.newhyps")
var symshen_4type_1signature_1of_1cn = MakeSymbol("shen.type-signature-of-cn")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4newcontinuation = MakeSymbol("shen.newcontinuation")
var symshen_4patthyps = MakeSymbol("shen.patthyps")
var symshen_4_5anymulti_6 = MakeSymbol("shen.<anymulti>")
var symshen_4type_1signature_1of_1map = MakeSymbol("shen.type-signature-of-map")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4abstract__rule = MakeSymbol("shen.abstract_rule")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4type_1signature_1of_1profile = MakeSymbol("shen.type-signature-of-profile")
var symshen_4externally_1bound = MakeSymbol("shen.externally-bound")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4yacc__cases = MakeSymbol("shen.yacc_cases")
var symshen_4compile__to__machine__code = MakeSymbol("shen.compile_to_machine_code")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symshen_4_5side_1condition_6 = MakeSymbol("shen.<side-condition>")
var symcut = MakeSymbol("cut")
var symshen_4datatype_1error = MakeSymbol("shen.datatype-error")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4packageh = MakeSymbol("shen.packageh")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4percent = MakeSymbol("shen.percent")
var symstr = MakeSymbol("str")
var symshen_4embed_1and = MakeSymbol("shen.embed-and")
var symshen_4insert__modes = MakeSymbol("shen.insert_modes")
var symshen_4mu__reduction = MakeSymbol("shen.mu_reduction")
var symshen_4_5comma_1symbol_6 = MakeSymbol("shen.<comma-symbol>")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4terminator_2 = MakeSymbol("shen.terminator?")
var symshen_4type_1signature_1of_1tc = MakeSymbol("shen.type-signature-of-tc")
var symspy = MakeSymbol("spy")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4compile__prolog__procedure = MakeSymbol("shen.compile_prolog_procedure")
var symshen_4type_1signature_1of_1shen_4prhush = MakeSymbol("shen.type-signature-of-shen.prhush")
var symspecialise = MakeSymbol("specialise")
var symshen_4find_1past_1inputs = MakeSymbol("shen.find-past-inputs")
var symshen_4reverse__help = MakeSymbol("shen.reverse_help")
var symshen_4type_1signature_1of_1freeze = MakeSymbol("shen.type-signature-of-freeze")
var symopen = MakeSymbol("open")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4_5side_1conditions_6 = MakeSymbol("shen.<side-conditions>")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var syminferences = MakeSymbol("inferences")
var symshen_4function_1macro = MakeSymbol("shen.function-macro")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4ebr = MakeSymbol("shen.ebr")
var symshen_4type_1signature_1of_1if = MakeSymbol("shen.type-signature-of-if")
var symshen_4type_1signature_1of_1write_1byte = MakeSymbol("shen.type-signature-of-write-byte")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4cons__form = MakeSymbol("shen.cons_form")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4symbol_1code_2 = MakeSymbol("shen.symbol-code?")
var symshen_4make__mu__application = MakeSymbol("shen.make_mu_application")
var symset = MakeSymbol("set")
var symshen_4clause__form = MakeSymbol("shen.clause_form")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4start_1new_1prolog_1process = MakeSymbol("shen.start-new-prolog-process")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs = MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4lineread_1loop = MakeSymbol("shen.lineread-loop")
var symshen_4newline = MakeSymbol("shen.newline")
var symshen_4call__the__continuation = MakeSymbol("shen.call_the_continuation")
var symshen_4right_1_6left = MakeSymbol("shen.right->left")
var symshen_4_5semicolon_1symbol_6 = MakeSymbol("shen.<semicolon-symbol>")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4type_1signature_1of_1fail_1if = MakeSymbol("shen.type-signature-of-fail-if")
var symshen_4find = MakeSymbol("shen.find")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var sym_6 = MakeSymbol(">")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4x_4factorise_1defun_4generate_1label = MakeSymbol("shen.x.factorise-defun.generate-label")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4_5num_6 = MakeSymbol("shen.<num>")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4continuation__call = MakeSymbol("shen.continuation_call")
var symshen_4type_1signature_1of_1pos = MakeSymbol("shen.type-signature-of-pos")
var symshen_4x_4factorise_1defun_4free_1variables = MakeSymbol("shen.x.factorise-defun.free-variables")
var symshen_4type_1signature_1of_1append = MakeSymbol("shen.type-signature-of-append")
var symshen_4trim_1gubbins = MakeSymbol("shen.trim-gubbins")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4eval_1cons = MakeSymbol("shen.eval-cons")
var symuntrack = MakeSymbol("untrack")
var symshen_4type_1signature_1of_1and = MakeSymbol("shen.type-signature-of-and")
var symshen_4_5singleunderline_6 = MakeSymbol("shen.<singleunderline>")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symtl = MakeSymbol("tl")
var symshen_4print_1past_1inputs = MakeSymbol("shen.print-past-inputs")
var symshen_4prh = MakeSymbol("shen.prh")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4type_1signature_1of_1close = MakeSymbol("shen.type-signature-of-close")
var symhd = MakeSymbol("hd")
var symabsvector_2 = MakeSymbol("absvector?")
var symvector = MakeSymbol("vector")
var symshen_4_5st__input1_6 = MakeSymbol("shen.<st_input1>")
var symshen_4_5anysingle_6 = MakeSymbol("shen.<anysingle>")
var symshen_4rule_1_6horn_1clause_1body = MakeSymbol("shen.rule->horn-clause-body")
var symshen_4type_1signature_1of_1shen_4insert = MakeSymbol("shen.type-signature-of-shen.insert")
var symtuple_2 = MakeSymbol("tuple?")
var symshen_4read_1char_1code = MakeSymbol("shen.read-char-code")
var symsum = MakeSymbol("sum")
var symshen_4tlhd = MakeSymbol("shen.tlhd")
var symshen_4type_1signature_1of_1remove = MakeSymbol("shen.type-signature-of-remove")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4copy_1vector = MakeSymbol("shen.copy-vector")
var symshen_4procedure__name = MakeSymbol("shen.procedure_name")
var symshen_4type_1signature_1of_1bound_2 = MakeSymbol("shen.type-signature-of-bound?")
var symshen_4complexity = MakeSymbol("shen.complexity")
var symshen_4mk_1pvar = MakeSymbol("shen.mk-pvar")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symshen_4type_1signature_1of_1reverse = MakeSymbol("shen.type-signature-of-reverse")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_5doubleunderline_6 = MakeSymbol("shen.<doubleunderline>")
var symshen_4semantic_1completion_1warning = MakeSymbol("shen.semantic-completion-warning")
var symfindall = MakeSymbol("findall")
var symshen_4type_1signature_1of_1destroy = MakeSymbol("shen.type-signature-of-destroy")
var symshen_4type_1signature_1of_1track = MakeSymbol("shen.type-signature-of-track")
var symshen_4compile__to__kl = MakeSymbol("shen.compile_to_kl")
var symshen_4t_d_1hyps = MakeSymbol("shen.t*-hyps")
var symshen_4variant_2 = MakeSymbol("shen.variant?")
var symshen_4kill_1code = MakeSymbol("shen.kill-code")
var symshen_4type_1signature_1of_1vector_2 = MakeSymbol("shen.type-signature-of-vector?")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4pre = MakeSymbol("shen.pre")
var symshen_4remember = MakeSymbol("shen.remember")
var symvalue = MakeSymbol("value")
var symshen_4next_150 = MakeSymbol("shen.next-50")
var symshen_4update_1demodulation_1function = MakeSymbol("shen.update-demodulation-function")
var symshen_4type_1signature_1of_1do = MakeSymbol("shen.type-signature-of-do")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4multiple_1set = MakeSymbol("shen.multiple-set")
var symshen_4read_1file_1as_1Xlist = MakeSymbol("shen.read-file-as-Xlist")
var symshen_4type_1signature_1of_1_a_a = MakeSymbol("shen.type-signature-of-==")
var symtlstr = MakeSymbol("tlstr")
var symshen_4ues = MakeSymbol("shen.ues")
var sym_5_1address = MakeSymbol("<-address")
var symget = MakeSymbol("get")
var symshen_4terminal_2 = MakeSymbol("shen.terminal?")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4copy_1vector_1stage_12 = MakeSymbol("shen.copy-vector-stage-2")
var symshen_4ue_2 = MakeSymbol("shen.ue?")
var symtype = MakeSymbol("type")
var symadjoin = MakeSymbol("adjoin")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4type_1signature_1of_1empty_2 = MakeSymbol("shen.type-signature-of-empty?")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4type_1signature_1of_1function = MakeSymbol("shen.type-signature-of-function")
var symshen_4type_1signature_1of_1nth = MakeSymbol("shen.type-signature-of-nth")
var sympreclude = MakeSymbol("preclude")
var symmap = MakeSymbol("map")
var symmapcan = MakeSymbol("mapcan")
var symshen_4construct_1search_1clauses = MakeSymbol("shen.construct-search-clauses")
var symdeclare = MakeSymbol("declare")
var symshen_4hdtl = MakeSymbol("shen.hdtl")
var symshen_4cond_1expression = MakeSymbol("shen.cond-expression")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symshen_4encode_1choices = MakeSymbol("shen.encode-choices")
var symshen_4prolog__constant_2 = MakeSymbol("shen.prolog_constant?")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4type_1signature_1of_1hdv = MakeSymbol("shen.type-signature-of-hdv")
var symshen_4safe_1multiply = MakeSymbol("shen.safe-multiply")
var symshen_4remove_1synonyms = MakeSymbol("shen.remove-synonyms")
var symtail = MakeSymbol("tail")
var symshen_4alphanum_2 = MakeSymbol("shen.alphanum?")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4th_d = MakeSymbol("shen.th*")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4aum__to__shen = MakeSymbol("shen.aum_to_shen")
var symshen_4memo = MakeSymbol("shen.memo")
var symshen_4type_1signature_1of_1porters = MakeSymbol("shen.type-signature-of-porters")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4initialise_1prolog = MakeSymbol("shen.initialise-prolog")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4read_1error = MakeSymbol("shen.read-error")
var symhdv = MakeSymbol("hdv")
var symshen_4type_1signature_1of_1length = MakeSymbol("shen.type-signature-of-length")
var symshen_4call_1help = MakeSymbol("shen.call-help")
var symshen_4add_1end = MakeSymbol("shen.add-end")
var symshen_4variancy_1test = MakeSymbol("shen.variancy-test")
var syminternal = MakeSymbol("internal")
var symcompile = MakeSymbol("compile")
var symlimit = MakeSymbol("limit")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4s_1prolog__clause = MakeSymbol("shen.s-prolog_clause")
var symshen_4include_1h = MakeSymbol("shen.include-h")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4cn_1all = MakeSymbol("shen.cn-all")
var symshen_4cf__help = MakeSymbol("shen.cf_help")
var symshen_4type_1signature_1of_1include_1all_1but = MakeSymbol("shen.type-signature-of-include-all-but")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4_5body_d_6 = MakeSymbol("shen.<body*>")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4findallhelp = MakeSymbol("shen.findallhelp")
var symshen_4type_1signature_1of_1inferences = MakeSymbol("shen.type-signature-of-inferences")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4update__history = MakeSymbol("shen.update_history")
var symshen_4shen_1syntax_1error = MakeSymbol("shen.shen-syntax-error")
var symshen_4mode_1ify = MakeSymbol("shen.mode-ify")
var symshen_4linearise__X = MakeSymbol("shen.linearise_X")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4decons_1string = MakeSymbol("shen.decons-string")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4type_1signature_1of_1vector_1_6 = MakeSymbol("shen.type-signature-of-vector->")
var symexternal = MakeSymbol("external")
var symshen_4type_1signature_1of_1_d = MakeSymbol("shen.type-signature-of-*")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symintern = MakeSymbol("intern")
var symarity = MakeSymbol("arity")
var symshen_4abstraction__build = MakeSymbol("shen.abstraction_build")
var symshen_4_5predicate_d_6 = MakeSymbol("shen.<predicate*>")
var symshen_4f__error = MakeSymbol("shen.f_error")
var symshen_4prbytes = MakeSymbol("shen.prbytes")
var symshen_4pushnew = MakeSymbol("shen.pushnew")
var symshen_4x_4factorise_1defun_4bind_1selector = MakeSymbol("shen.x.factorise-defun.bind-selector")
var symshen_4rule_1_6horn_1clause = MakeSymbol("shen.rule->horn-clause")
var symshen_4type_1signature_1of_1systemf = MakeSymbol("shen.type-signature-of-systemf")
var symshen_4type_1signature_1of_1_1 = MakeSymbol("shen.type-signature-of--")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4package_1macro = MakeSymbol("shen.package-macro")
var symshen_4mod = MakeSymbol("shen.mod")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4ue_1sig = MakeSymbol("shen.ue-sig")
var symshen_4update_1symbol_1table = MakeSymbol("shen.update-symbol-table")
var symshen_4jump__stream_2 = MakeSymbol("shen.jump_stream?")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4_5predigits_6 = MakeSymbol("shen.<predigits>")
var symshen_4type_1signature_1of_1read_1file_1as_1bytelist = MakeSymbol("shen.type-signature-of-read-file-as-bytelist")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4errormaxinfs = MakeSymbol("shen.errormaxinfs")
var symshen_4result_1type = MakeSymbol("shen.result-type")
var symidentical = MakeSymbol("identical")
var symunput = MakeSymbol("unput")
var symshen_4semantics = MakeSymbol("shen.semantics")
var symshen_4type_1signature_1of_1arity = MakeSymbol("shen.type-signature-of-arity")
var symshen_4type_1signature_1of_1error_1to_1string = MakeSymbol("shen.type-signature-of-error-to-string")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4prolog_1error = MakeSymbol("shen.prolog-error")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4head__abstraction = MakeSymbol("shen.head_abstraction")
var symshen_4valvector = MakeSymbol("shen.valvector")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symunprofile = MakeSymbol("unprofile")
var symcn = MakeSymbol("cn")
var symshen_4group__clauses = MakeSymbol("shen.group_clauses")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4_5postdigits_6 = MakeSymbol("shen.<postdigits>")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symshen_4list_1stream = MakeSymbol("shen.list-stream")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4active_1cons = MakeSymbol("shen.active-cons")
var symunify_b = MakeSymbol("unify!")
var symshen_4type_1signature_1of_1optimise = MakeSymbol("shen.type-signature-of-optimise")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symshen_4assumetype = MakeSymbol("shen.assumetype")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symshen_4doubleunderline_2 = MakeSymbol("shen.doubleunderline?")
var symshen_4type_1signature_1of_1specialise = MakeSymbol("shen.type-signature-of-specialise")
var symshen_4typextable = MakeSymbol("shen.typextable")
var symshen_4remove_1bar = MakeSymbol("shen.remove-bar")
var symshen_4udefs_d = MakeSymbol("shen.udefs*")
var symshen_4type_1signature_1of_1_5_1vector = MakeSymbol("shen.type-signature-of-<-vector")
var symshen_4parameters = MakeSymbol("shen.parameters")
var sympos = MakeSymbol("pos")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symshen_4type_1signature_1of_1_7 = MakeSymbol("shen.type-signature-of-+")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symabort = MakeSymbol("abort")
var symshen_4code_1point = MakeSymbol("shen.code-point")
var symshen_4type_1signature_1of_1enable_1type_1theory = MakeSymbol("shen.type-signature-of-enable-type-theory")
var symtrack = MakeSymbol("track")
var symshen_4rules_1_6horn_1clauses = MakeSymbol("shen.rules->horn-clauses")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4eval_1without_1macros = MakeSymbol("shen.eval-without-macros")
var symshen_4type_1signature_1of_1unspecialise = MakeSymbol("shen.type-signature-of-unspecialise")
var symshen_4x_4factorise_1defun_4free_1variables_1h = MakeSymbol("shen.x.factorise-defun.free-variables-h")
var symshen_4x_4factorise_1defun_4test_1_6selectors = MakeSymbol("shen.x.factorise-defun.test->selectors")
var symshen_4prolog_1aritycheck = MakeSymbol("shen.prolog-aritycheck")
var symshen_4ue = MakeSymbol("shen.ue")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symshen_4type_1signature_1of_1vector = MakeSymbol("shen.type-signature-of-vector")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4type_1signature_1of_1cons_2 = MakeSymbol("shen.type-signature-of-cons?")
var symshen_4type_1signature_1of_1head = MakeSymbol("shen.type-signature-of-head")
var symshen_4compile__to__lambda_7 = MakeSymbol("shen.compile_to_lambda+")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4aritycheck_1name = MakeSymbol("shen.aritycheck-name")
var symshen_4typedf_2 = MakeSymbol("shen.typedf?")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symelement_2 = MakeSymbol("element?")
var symshen_4unbindv = MakeSymbol("shen.unbindv")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4split__cc__rule = MakeSymbol("shen.split_cc_rule")
var symshen_4variable_1match = MakeSymbol("shen.variable-match")
var symshen_4resize_1vector = MakeSymbol("shen.resize-vector")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4curry_1type_1h = MakeSymbol("shen.curry-type-h")
var symshen_4split__cc__rules = MakeSymbol("shen.split_cc_rules")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4type_1signature_1of_1stinput = MakeSymbol("shen.type-signature-of-stinput")
var symoccurs_1check = MakeSymbol("occurs-check")
var sym_5_1vector = MakeSymbol("<-vector")
var symshen_4x_4factorise_1defun_4concat_c = MakeSymbol("shen.x.factorise-defun.concat/")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4curry_1synonyms = MakeSymbol("shen.curry-synonyms")
var symshen_4type_1signature_1of_1sterror = MakeSymbol("shen.type-signature-of-sterror")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4_5term_d_6 = MakeSymbol("shen.<term*>")
var symprofile = MakeSymbol("profile")
var symsnd = MakeSymbol("snd")
var symshen_4package_1contents = MakeSymbol("shen.package-contents")
var symshen_4collect = MakeSymbol("shen.collect")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4type_1signature_1of_1nl = MakeSymbol("shen.type-signature-of-nl")
var symshen_4type_1signature_1of_1os = MakeSymbol("shen.type-signature-of-os")
var symshen_4reduce__help = MakeSymbol("shen.reduce_help")
var symshen_4x_4factorise_1defun_4rebranch_1h = MakeSymbol("shen.x.factorise-defun.rebranch-h")
var symshen_4initialise__arity__table = MakeSymbol("shen.initialise_arity_table")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4tab = MakeSymbol("shen.tab")
var symshen_4mult__subst = MakeSymbol("shen.mult_subst")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4type_1signature_1of_1adjoin = MakeSymbol("shen.type-signature-of-adjoin")
var symshen_4type_1signature_1of_1pr = MakeSymbol("shen.type-signature-of-pr")
var symshen_4dh_2 = MakeSymbol("shen.dh?")
var symshen_4construct_1premiss_1literal = MakeSymbol("shen.construct-premiss-literal")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4left_1round = MakeSymbol("shen.left-round")
var symshen_4type_1signature_1of_1read_1byte = MakeSymbol("shen.type-signature-of-read-byte")
var symshen_4type_1signature_1of_1stoutput = MakeSymbol("shen.type-signature-of-stoutput")
var symshen_4type_1signature_1of_1cd = MakeSymbol("shen.type-signature-of-cd")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4construct_1search_1clause = MakeSymbol("shen.construct-search-clause")
var symshen_4complexity__head = MakeSymbol("shen.complexity_head")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symcons = MakeSymbol("cons")
var symshen_4elim_1def = MakeSymbol("shen.elim-def")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4type_1signature_1of_1mapcan = MakeSymbol("shen.type-signature-of-mapcan")
var symshen_4record_1exceptions = MakeSymbol("shen.record-exceptions")
var symshen_4type_1signature_1of_1y_1or_1n_2 = MakeSymbol("shen.type-signature-of-y-or-n?")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symshen_4type_1signature_1of_1load = MakeSymbol("shen.type-signature-of-load")
var symfix = MakeSymbol("fix")
var symfst = MakeSymbol("fst")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4t_d_1defh = MakeSymbol("shen.t*-defh")
var symshen_4type_1signature_1of_1fix = MakeSymbol("shen.type-signature-of-fix")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symlength = MakeSymbol("length")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4aritycheck_1action = MakeSymbol("shen.aritycheck-action")
var symshen_4default_1rule = MakeSymbol("shen.default-rule")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symnth = MakeSymbol("nth")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4reduce = MakeSymbol("shen.reduce")
var symshen_4recursive__descent = MakeSymbol("shen.recursive_descent")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4initialise_1signedfunc_1lambda_1forms = MakeSymbol("shen.initialise-signedfunc-lambda-forms")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4_5end_d_6 = MakeSymbol("shen.<end*>")
var sympackage_2 = MakeSymbol("package?")
var symunion = MakeSymbol("union")
var symsubst = MakeSymbol("subst")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var symshen_4aritycheck = MakeSymbol("shen.aritycheck")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4t_d_1action = MakeSymbol("shen.t*-action")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4check__stream = MakeSymbol("shen.check_stream")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symshen_4type_1signature_1of_1trap_1error = MakeSymbol("shen.type-signature-of-trap-error")
var symreverse = MakeSymbol("reverse")
var symshen_4pair = MakeSymbol("shen.pair")
var symshen_4free__variable__warnings = MakeSymbol("shen.free_variable_warnings")
var symthaw = MakeSymbol("thaw")
var symremove = MakeSymbol("remove")
var symshen_4prolog_1failure = MakeSymbol("shen.prolog-failure")
var symshen_4type_1signature_1of_1package_2 = MakeSymbol("shen.type-signature-of-package?")
var symshen_4loop = MakeSymbol("shen.loop")
var symassoc = MakeSymbol("assoc")
var symnumber_2 = MakeSymbol("number?")
var symshen_4x_4factorise_1defun_4findpos = MakeSymbol("shen.x.factorise-defun.findpos")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4digits_1_6integers = MakeSymbol("shen.digits->integers")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4type_1signature_1of_1shen_4proc_1nl = MakeSymbol("shen.type-signature-of-shen.proc-nl")
var symshen_4type_1signature_1of_1tc_2 = MakeSymbol("shen.type-signature-of-tc?")
var symshen_4locally_1bind_1prolog_1vs = MakeSymbol("shen.locally-bind-prolog-vs")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4type_1signature_1of_1shen_4app = MakeSymbol("shen.type-signature-of-shen.app")
var symversion = MakeSymbol("version")
var symshen_4toplineread__loop = MakeSymbol("shen.toplineread_loop")
var symshen_4toplevel__evaluate = MakeSymbol("shen.toplevel_evaluate")
var symshen_4control_1chars = MakeSymbol("shen.control-chars")
var symshen_4make_1key = MakeSymbol("shen.make-key")
var symabsvector = MakeSymbol("absvector")
var symstinput = MakeSymbol("stinput")
var symnot = MakeSymbol("not")
var symshen_4_5signature_1help_6 = MakeSymbol("shen.<signature-help>")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4type_1signature_1of_1integer_2 = MakeSymbol("shen.type-signature-of-integer?")
var symnl = MakeSymbol("nl")
var symgensym = MakeSymbol("gensym")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symshen_4type_1signature_1of_1maxinferences = MakeSymbol("shen.type-signature-of-maxinferences")
var symshen_4type_1signature_1of_1_6_a = MakeSymbol("shen.type-signature-of->=")
var symshen_4free__variable__check = MakeSymbol("shen.free_variable_check")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4type_1signature_1of_1fail = MakeSymbol("shen.type-signature-of-fail")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4type_1signature_1of_1read_1file_1as_1string = MakeSymbol("shen.type-signature-of-read-file-as-string")
var sym_7 = MakeSymbol("+")
var symshen_4insert_1predicate = MakeSymbol("shen.insert-predicate")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4type_1signature_1of_1write_1to_1file = MakeSymbol("shen.type-signature-of-write-to-file")
var symshen_4toplevel = MakeSymbol("shen.toplevel")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4type_1signature_1of_1read_1from_1string = MakeSymbol("shen.type-signature-of-read-from-string")
var symshen_4prolog_1_6shen = MakeSymbol("shen.prolog->shen")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4x_4factorise_1defun_4inline_1mono_1labels = MakeSymbol("shen.x.factorise-defun.inline-mono-labels")
var symshen_4extract_1pvars = MakeSymbol("shen.extract-pvars")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4singleunderline_2 = MakeSymbol("shen.singleunderline?")
var symempty_2 = MakeSymbol("empty?")
var symget_1time = MakeSymbol("get-time")
var symshen_4type_1signature_1of_1boolean_2 = MakeSymbol("shen.type-signature-of-boolean?")
var symprint = MakeSymbol("print")
var symimplementation = MakeSymbol("implementation")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symcd = MakeSymbol("cd")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4type_1signature_1of_1ps = MakeSymbol("shen.type-signature-of-ps")
var symshen_4syntax = MakeSymbol("shen.syntax")
var symshen_4jump__stream = MakeSymbol("shen.jump_stream")
var symshen_4legitimate_1term_2 = MakeSymbol("shen.legitimate-term?")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4type_1signature_1of_1snd = MakeSymbol("shen.type-signature-of-snd")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4lookup_1func = MakeSymbol("shen.lookup-func")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symfunction = MakeSymbol("function")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var symshen_4sh_2 = MakeSymbol("shen.sh?")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4err_1condition = MakeSymbol("shen.err-condition")
var symdifference = MakeSymbol("difference")
var sym_8p = MakeSymbol("@p")
var symshen_4type_1signature_1of_1number_2 = MakeSymbol("shen.type-signature-of-number?")
var symlanguage = MakeSymbol("language")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4type_1signature_1of_1spy = MakeSymbol("shen.type-signature-of-spy")
var symshen_4type_1signature_1of_1untrack = MakeSymbol("shen.type-signature-of-untrack")
var symstep = MakeSymbol("step")
var symlineread = MakeSymbol("lineread")
var symstoutput = MakeSymbol("stoutput")
var symshen_4preclude_1h = MakeSymbol("shen.preclude-h")
var symshen_4insert_1deref = MakeSymbol("shen.insert-deref")
var symshen_4t_d = MakeSymbol("shen.t*")
var symsystemf = MakeSymbol("systemf")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4_5sig_7rules_6 = MakeSymbol("shen.<sig+rules>")
var symshen_4type_1signature_1of_1version = MakeSymbol("shen.type-signature-of-version")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4insert_1prolog_1variables_1help = MakeSymbol("shen.insert-prolog-variables-help")
var symshen_4remove_1nth = MakeSymbol("shen.remove-nth")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4type_1signature_1of_1occurs_1check = MakeSymbol("shen.type-signature-of-occurs-check")
var symshen_4type_1signature_1of_1tlv = MakeSymbol("shen.type-signature-of-tlv")
var symread = MakeSymbol("read")
var symshen_4type_1signature_1of_1intersection = MakeSymbol("shen.type-signature-of-intersection")
var syminclude = MakeSymbol("include")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4type_1signature_1of_1element_2 = MakeSymbol("shen.type-signature-of-element?")
var symshen_4lzy_a_a = MakeSymbol("shen.lzy==")
var symshen_4x_4factorise_1defun_4apply_1selector_1handlers = MakeSymbol("shen.x.factorise-defun.apply-selector-handlers")
var symshen_4read_1file_1as_1Xlist_1help = MakeSymbol("shen.read-file-as-Xlist-help")
var symshen_4remove__modes = MakeSymbol("shen.remove_modes")
var symshen_4x_4factorise_1defun_4add_1returns = MakeSymbol("shen.x.factorise-defun.add-returns")
var symshen_4x_4factorise_1defun_4optimize_1selectors = MakeSymbol("shen.x.factorise-defun.optimize-selectors")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4s_1prolog = MakeSymbol("shen.s-prolog")
var symshen_4default__semantics = MakeSymbol("shen.default_semantics")
var symdestroy = MakeSymbol("destroy")
var symshen_4x_4factorise_1defun_4bind_1repeating_1selectors = MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors")
var symshen_4add_1macro = MakeSymbol("shen.add-macro")
var symshen_4_5premises_6 = MakeSymbol("shen.<premises>")
var symcall = MakeSymbol("call")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4type_1signature_1of_1n_1_6string = MakeSymbol("shen.type-signature-of-n->string")
var symshen_4type_1signature_1of_1step = MakeSymbol("shen.type-signature-of-step")
var symdefun = MakeSymbol("defun")
var symcons_2 = MakeSymbol("cons?")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4type_1signature_1of_1read = MakeSymbol("shen.type-signature-of-read")
var symshen_4copy_1vector_1stage_11 = MakeSymbol("shen.copy-vector-stage-1")
var symbound_2 = MakeSymbol("bound?")
var symshen_4type_1signature_1of_1absvector_2 = MakeSymbol("shen.type-signature-of-absvector?")
var symunify = MakeSymbol("unify")
var symshen_4type_1signature_1of_1print = MakeSymbol("shen.type-signature-of-print")
var symshen_4type_1signature_1of_1protect = MakeSymbol("shen.type-signature-of-protect")
var symshen_4_5simple__pattern_6 = MakeSymbol("shen.<simple_pattern>")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4_5head_d_6 = MakeSymbol("shen.<head*>")
var symshen_4construct_1base_1search_1clause = MakeSymbol("shen.construct-base-search-clause")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4explicit__modes = MakeSymbol("shen.explicit_modes")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4prefix_2 = MakeSymbol("shen.prefix?")
var symshen_4get_1type = MakeSymbol("shen.get-type")
var symshen_4atom_1type = MakeSymbol("shen.atom-type")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4synonyms_1help = MakeSymbol("shen.synonyms-help")
var symshen_4safe_1product = MakeSymbol("shen.safe-product")
var symshen_4right_1rule = MakeSymbol("shen.right-rule")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4assign_1types = MakeSymbol("shen.assign-types")
var symshen_4x_4factorise_1defun_4with_1labelled_1else = MakeSymbol("shen.x.factorise-defun.with-labelled-else")
var symshen_4_5guard_6 = MakeSymbol("shen.<guard>")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4placeholder = MakeSymbol("shen.placeholder")
var symshen_4t_d_1defhh = MakeSymbol("shen.t*-defhh")
var symexplode = MakeSymbol("explode")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4compose = MakeSymbol("shen.compose")
var symshen_4type_1signature_1of_1_5_a = MakeSymbol("shen.type-signature-of-<=")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4aum = MakeSymbol("shen.aum")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var sym_1 = MakeSymbol("-")
var symshen_4base = MakeSymbol("shen.base")
var symshen_4type_1signature_1of_1symbol_2 = MakeSymbol("shen.type-signature-of-symbol?")
var symshen_4intprolog_1help_1help = MakeSymbol("shen.intprolog-help-help")
var symshen_4application__build = MakeSymbol("shen.application_build")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4grammar__symbol_2 = MakeSymbol("shen.grammar_symbol?")
var sympr = MakeSymbol("pr")
var symshen_4rule_1_6horn_1clause_1head = MakeSymbol("shen.rule->horn-clause-head")
var symeval = MakeSymbol("eval")
var symshen_4x_4factorise_1defun_4rebranch = MakeSymbol("shen.x.factorise-defun.rebranch")
var symshen_4hat = MakeSymbol("shen.hat")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4x_4features_4current = MakeSymbol("shen.x.features.current")
var symshen_4_5sig_7rest_6 = MakeSymbol("shen.<sig+rest>")
var symshen_4type_1signature_1of_1it = MakeSymbol("shen.type-signature-of-it")
var symshen_4type_1signature_1of_1_c = MakeSymbol("shen.type-signature-of-/")
var symshen_4strip_1pathname = MakeSymbol("shen.strip-pathname")
var syminput_7 = MakeSymbol("input+")
var symshen_4linearise__help = MakeSymbol("shen.linearise_help")
var symclose = MakeSymbol("close")
var symshen_4post = MakeSymbol("shen.post")
var symreturn = MakeSymbol("return")
var symshen_4type_1signature_1of_1occurrences = MakeSymbol("shen.type-signature-of-occurrences")
var symshen_4read_1file_1as_1charlist = MakeSymbol("shen.read-file-as-charlist")
var symshen_4type_1signature_1of_1sum = MakeSymbol("shen.type-signature-of-sum")
var symshen_4same__predicate_2 = MakeSymbol("shen.same_predicate?")
var symshen_4t_d_1patterns = MakeSymbol("shen.t*-patterns")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symfwhen = MakeSymbol("fwhen")
var symshen_4type_1signature_1of_1tuple_2 = MakeSymbol("shen.type-signature-of-tuple?")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4x_4factorise_1defun_4true_1branch = MakeSymbol("shen.x.factorise-defun.true-branch")
var symshen_4_5st__input2_6 = MakeSymbol("shen.<st_input2>")
var symshen_4type_1signature_1of_1_6 = MakeSymbol("shen.type-signature-of->")
var symvector_2 = MakeSymbol("vector?")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4type_1signature_1of_1difference = MakeSymbol("shen.type-signature-of-difference")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4type_1signature_1of_1tail = MakeSymbol("shen.type-signature-of-tail")
var symshen_4type_1signature_1of_1hdstr = MakeSymbol("shen.type-signature-of-hdstr")
var symshen_4clauses_1to_1shen = MakeSymbol("shen.clauses-to-shen")
var sym_8v = MakeSymbol("@v")
var symshen_4left_1rule = MakeSymbol("shen.left-rule")
var symshen_4cc__body = MakeSymbol("shen.cc_body")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4type_1signature_1of_1or = MakeSymbol("shen.type-signature-of-or")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4demod_1rule = MakeSymbol("shen.demod-rule")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var symshen_4type_1signature_1of_1get_1time = MakeSymbol("shen.type-signature-of-get-time")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4type_1signature_1of_1_5_b_6 = MakeSymbol("shen.type-signature-of-<!>")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4_5non_1return_6 = MakeSymbol("shen.<non-return>")
var symshen_4insert_1lazyderef = MakeSymbol("shen.insert-lazyderef")
var symrelease = MakeSymbol("release")
var symshen_4_5literal_d_6 = MakeSymbol("shen.<literal*>")
var symshen_4type_1signature_1of_1compile = MakeSymbol("shen.type-signature-of-compile")
var symshen_4type_1signature_1of_1read_1file = MakeSymbol("shen.type-signature-of-read-file")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4write_1char_1and_1inc = MakeSymbol("shen.write-char-and-inc")
var symshen_4catchpoint = MakeSymbol("shen.catchpoint")
var symshen_4proc_1input_7 = MakeSymbol("shen.proc-input+")
var sym_6_a = MakeSymbol(">=")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var sym_a_a = MakeSymbol("==")
var symshen_4expt = MakeSymbol("shen.expt")
var sym_5 = MakeSymbol("<")
var symshen_4clash_2 = MakeSymbol("shen.clash?")
var symintersection = MakeSymbol("intersection")
var symshen_4type_1signature_1of_1external = MakeSymbol("shen.type-signature-of-external")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4type_1signature_1of_1assoc = MakeSymbol("shen.type-signature-of-assoc")
var symshen_4type_1signature_1of_1string_1_6symbol = MakeSymbol("shen.type-signature-of-string->symbol")
var symshen_4intprolog_1help = MakeSymbol("shen.intprolog-help")
var symshen_4compile_1macro = MakeSymbol("shen.compile-macro")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4trim_1whitespace = MakeSymbol("shen.trim-whitespace")
var symshen_4chwild = MakeSymbol("shen.chwild")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4bld_1prolog_1call = MakeSymbol("shen.bld-prolog-call")
var symshen_4t_d_1def = MakeSymbol("shen.t*-def")
var symshen_4ue_1h_2 = MakeSymbol("shen.ue-h?")
var symshen_4toplineread = MakeSymbol("shen.toplineread")
var symshen_4type_1signature_1of_1union = MakeSymbol("shen.type-signature-of-union")
var symshen_4exclamation = MakeSymbol("shen.exclamation")
var symshen_4construct_1recursive_1search_1clause = MakeSymbol("shen.construct-recursive-search-clause")
var symshen_4_5clauses_d_6 = MakeSymbol("shen.<clauses*>")
var symshen_4catch_1cut = MakeSymbol("shen.catch-cut")
var symshen_4x_4features_4cond_1expand_1macro = MakeSymbol("shen.x.features.cond-expand-macro")
var sym_8s = MakeSymbol("@s")
var symshen_4yacc = MakeSymbol("shen.yacc")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4insert_1runon = MakeSymbol("shen.insert-runon")
var symshen_4_5st__input_6 = MakeSymbol("shen.<st_input>")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4type_1signature_1of_1limit = MakeSymbol("shen.type-signature-of-limit")
var symshen_4by__hypothesis = MakeSymbol("shen.by_hypothesis")
var symhash = MakeSymbol("hash")
var symfail = MakeSymbol("fail")
var symshen_4nest_1disjunct = MakeSymbol("shen.nest-disjunct")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4type_1signature_1of_1not = MakeSymbol("shen.type-signature-of-not")
var symshen_4type_1signature_1of_1str = MakeSymbol("shen.type-signature-of-str")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4extract__vars = MakeSymbol("shen.extract_vars")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4call_1rest = MakeSymbol("shen.call-rest")
var symshen_4type_1signature_1of_1preclude_1all_1but = MakeSymbol("shen.type-signature-of-preclude-all-but")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4remtype = MakeSymbol("shen.remtype")
var symhdstr = MakeSymbol("hdstr")
var symshen_4x_4factorise_1defun_4attach_1free_1variables = MakeSymbol("shen.x.factorise-defun.attach-free-variables")
var symshen_4_5non_1ll_1rules_6 = MakeSymbol("shen.<non-ll-rules>")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4type_1signature_1of_1simple_1error = MakeSymbol("shen.type-signature-of-simple-error")
var symput = MakeSymbol("put")
var symshen_4type_1signature_1of_1fst = MakeSymbol("shen.type-signature-of-fst")
var symshen_4type_1signature_1of_1language = MakeSymbol("shen.type-signature-of-language")
var symshen_4retrieve_1from_1history_1if_1needed = MakeSymbol("shen.retrieve-from-history-if-needed")
var symshen_4sequent = MakeSymbol("shen.sequent")
var symshen_4_5premise_6 = MakeSymbol("shen.<premise>")
var symshen_4type_1signature_1of_1explode = MakeSymbol("shen.type-signature-of-explode")
var symshen_4_5action_6 = MakeSymbol("shen.<action>")
var symshen_4_5E_6 = MakeSymbol("shen.<E>")
var symshen_4type_1signature_1of_1thaw = MakeSymbol("shen.type-signature-of-thaw")
var symshen_4hdhd = MakeSymbol("shen.hdhd")
var symshen_4lisp_1or = MakeSymbol("shen.lisp-or")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4type_1signature_1of_1hash = MakeSymbol("shen.type-signature-of-hash")
var symshen_4construct_1side_1literals = MakeSymbol("shen.construct-side-literals")
var symshen_4record_1it_1h = MakeSymbol("shen.record-it-h")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4type_1signature_1of_1unprofile = MakeSymbol("shen.type-signature-of-unprofile")
var sym_a = MakeSymbol("=")
var symshen_4carriage_1return = MakeSymbol("shen.carriage-return")
var sym_d = MakeSymbol("*")
var symshen_4type_1signature_1of_1profile_1results = MakeSymbol("shen.type-signature-of-profile-results")
var symshen_4space = MakeSymbol("shen.space")
var symshen_4extract__free__vars = MakeSymbol("shen.extract_free_vars")
var symshen_4cutpoint = MakeSymbol("shen.cutpoint")
var symshen_4type_1signature_1of_1_5e_6 = MakeSymbol("shen.type-signature-of-<e>")
var symshen_4type_1signature_1of_1include = MakeSymbol("shen.type-signature-of-include")
var symshen_4construct_1search_1literals = MakeSymbol("shen.construct-search-literals")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4after_1codepoint = MakeSymbol("shen.after-codepoint")
var symshen_4type_1signature_1of_1kill = MakeSymbol("shen.type-signature-of-kill")
var symshen_4removetype = MakeSymbol("shen.removetype")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4type_1signature_1of_1string_2 = MakeSymbol("shen.type-signature-of-string?")
var symshen_4x_4factorise_1defun_4false_1branch = MakeSymbol("shen.x.factorise-defun.false-branch")
var symshen_4_5variable_2_6 = MakeSymbol("shen.<variable?>")
var symshen_4csl_1help = MakeSymbol("shen.csl-help")
var symps = MakeSymbol("ps")
var symshen_4record_1source = MakeSymbol("shen.record-source")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4type_1signature_1of_1variable_2 = MakeSymbol("shen.type-signature-of-variable?")
var symload = MakeSymbol("load")
var symshen_4s_1prolog__literal = MakeSymbol("shen.s-prolog_literal")
var symshen_4type_1signature_1of_1tlstr = MakeSymbol("shen.type-signature-of-tlstr")
var symprotect = MakeSymbol("protect")
var symshen_4type_1signature_1of_1implementation = MakeSymbol("shen.type-signature-of-implementation")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4typecheck_1and_1load = MakeSymbol("shen.typecheck-and-load")
var symappend = MakeSymbol("append")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4cc__help = MakeSymbol("shen.cc_help")
var symshen_4list__variables = MakeSymbol("shen.list_variables")
var symmacroexpand = MakeSymbol("macroexpand")
var symshen_4add__test = MakeSymbol("shen.add_test")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4type_1signature_1of_1preclude = MakeSymbol("shen.type-signature-of-preclude")
var sym_5_a = MakeSymbol("<=")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4x_4factorise_1defun_4factorise_1cond = MakeSymbol("shen.x.factorise-defun.factorise-cond")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4numbyte_2 = MakeSymbol("shen.numbyte?")
var symshen_4function_1abstraction_1help = MakeSymbol("shen.function-abstraction-help")
var symoptimise = MakeSymbol("optimise")
var symshen_4type_1signature_1of_1release = MakeSymbol("shen.type-signature-of-release")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symstring_2 = MakeSymbol("string?")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4function_1abstraction = MakeSymbol("shen.function-abstraction")
var syminput = MakeSymbol("input")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4findpos = MakeSymbol("shen.findpos")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4decimalise = MakeSymbol("shen.decimalise")
var symshen_4resizeprocessvector = MakeSymbol("shen.resizeprocessvector")
var symshen_4_5clause_d_6 = MakeSymbol("shen.<clause*>")
var symshen_4recursive__cons__form = MakeSymbol("shen.recursive_cons_form")
var symshen_4type_1signature_1of_1string_1_6n = MakeSymbol("shen.type-signature-of-string->n")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symhead = MakeSymbol("head")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symshen_4type_1signature_1of_1undefmacro = MakeSymbol("shen.type-signature-of-undefmacro")
