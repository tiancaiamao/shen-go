package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

		gen17390 := MakeNative(func(__e Evaluator) {
			V4891 := __e.Get(1)
			_ = V4891
			gen17388 := Call(__e, ShenFunc(symread_1file), V4891)

			Contents := gen17388
			gen17389 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4eval_1without_1macros), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symmap), gen17389, Contents)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.quiet-load"), gen17390)

		gen17406 := MakeNative(func(__e Evaluator) {
			gen17391 := Call(__e, ShenFunc(symversion))

			gen17392 := Call(__e, ShenFunc(symlanguage))

			gen17393 := Call(__e, ShenFunc(symport))

			gen17394 := Call(__e, ShenFunc(symcons), gen17393, Nil)

			gen17395 := Call(__e, ShenFunc(symcons), gen17392, gen17394)

			gen17396 := Call(__e, ShenFunc(symimplementation))

			gen17397 := Call(__e, ShenFunc(symrelease))

			gen17398 := Call(__e, ShenFunc(symcons), gen17397, Nil)

			gen17399 := Call(__e, ShenFunc(symcons), gen17396, gen17398)

			gen17400 := Call(__e, ShenFunc(symcons), gen17399, Nil)

			gen17401 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen17400)

			gen17402 := Call(__e, ShenFunc(symcons), gen17395, gen17401)

			gen17403 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen17402)

			gen17404 := Call(__e, ShenFunc(symshen_4app), gen17403, MakeString("\n"), MakeSymbol("shen.r"))

			gen17405 := Call(__e, ShenFunc(symcn), MakeString(" "), gen17404)

			__e.TailApply(ShenFunc(symshen_4app), gen17391, gen17405, MakeSymbol("shen.a"))

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.version-string"), gen17406)

		gen17408 := MakeNative(func(__e Evaluator) {
			V4893 := __e.Get(1)
			_ = V4893
			gen17407 := Call(__e, ShenFunc(symshen_4app), V4893, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symcn), MakeString("Usage: "), gen17407)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.help-text"), gen17408)

		gen17413 := MakeNative(func(__e Evaluator) {
			V4895 := __e.Get(1)
			_ = V4895
			gen17412 := Call(__e, ShenFunc(sym_a), Nil, V4895)

			if True == gen17412 {
				__e.TailApply(ShenFunc(symcons), MakeSymbol("success"), Nil)

				return
			} else {
				gen17411 := Call(__e, ShenFunc(symcons_2), V4895)

				if True == gen17411 {
					gen17409 := Call(__e, ShenFunc(symhd), V4895)

					Call(__e, ShenFunc(symthaw), gen17409)

					gen17410 := Call(__e, ShenFunc(symtl), V4895)

					__e.TailApply(ShenFunc(symshen_4x_4launcher_4execute_1all), gen17410)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.launcher.execute-all"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.execute-all"), gen17413)

		gen17416 := MakeNative(func(__e Evaluator) {
			V4897 := __e.Get(1)
			_ = V4897
			gen17414 := Call(__e, ShenFunc(symread_1from_1string), V4897)

			gen17415 := Call(__e, ShenFunc(symhead), gen17414)

			__e.TailApply(ShenFunc(symeval), gen17415)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-string"), gen17416)

		gen17422 := MakeNative(func(__e Evaluator) {
			V4903 := __e.Get(1)
			_ = V4903
			gen17421 := Call(__e, ShenFunc(sym_a), MakeString("-e"), V4903)

			if True == gen17421 {
				__e.Return(MakeString("--eval"))
				return
			} else {
				gen17420 := Call(__e, ShenFunc(sym_a), MakeString("-l"), V4903)

				if True == gen17420 {
					__e.Return(MakeString("--load"))
					return
				} else {
					gen17419 := Call(__e, ShenFunc(sym_a), MakeString("-q"), V4903)

					if True == gen17419 {
						__e.Return(MakeString("--quiet"))
						return
					} else {
						gen17418 := Call(__e, ShenFunc(sym_a), MakeString("-s"), V4903)

						if True == gen17418 {
							__e.Return(MakeString("--set"))
							return
						} else {
							gen17417 := Call(__e, ShenFunc(sym_a), MakeString("-r"), V4903)

							if True == gen17417 {
								__e.Return(MakeString("--repl"))
								return
							} else {
								__e.Return(False)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-flag-map"), gen17422)

		gen17504 := MakeNative(func(__e Evaluator) {
			V4914 := __e.Get(1)
			_ = V4914
			V4915 := __e.Get(2)
			_ = V4915
			gen17503 := Call(__e, ShenFunc(sym_a), Nil, V4914)

			if True == gen17503 {
				gen17502 := Call(__e, ShenFunc(symreverse), V4915)

				__e.TailApply(ShenFunc(symshen_4x_4launcher_4execute_1all), gen17502)

				return

			} else {
				gen17500 := Call(__e, ShenFunc(symcons_2), V4914)

				var gen17501 Obj
				if True == gen17500 {
					gen17497 := Call(__e, ShenFunc(symhd), V4914)

					gen17498 := Call(__e, ShenFunc(sym_a), MakeString("--eval"), gen17497)

					var gen17499 Obj
					if True == gen17498 {
						gen17495 := Call(__e, ShenFunc(symtl), V4914)

						gen17496 := Call(__e, ShenFunc(symcons_2), gen17495)

						if True == gen17496 {
							gen17499 = True
						} else {
							gen17499 = False
						}

					} else {
						gen17499 = False
					}
					if True == gen17499 {
						gen17501 = True
					} else {
						gen17501 = False
					}

				} else {
					gen17501 = False
				}
				if True == gen17501 {
					gen17486 := Call(__e, ShenFunc(symtl), V4914)

					gen17487 := Call(__e, ShenFunc(symtl), gen17486)

					gen17493 := MakeNative(func(__e Evaluator) {
						gen17488 := Call(__e, ShenFunc(symtl), V4914)

						gen17489 := Call(__e, ShenFunc(symhd), gen17488)

						gen17490 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1string), gen17489)

						gen17491 := Call(__e, ShenFunc(symshen_4app), gen17490, MakeString("\n"), MakeSymbol("shen.a"))

						gen17492 := Call(__e, ShenFunc(symstoutput))

						__e.TailApply(ShenFunc(symshen_4prhush), gen17491, gen17492)

						return

					}, 0)
					gen17494 := Call(__e, ShenFunc(symcons), gen17493, V4915)

					__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen17487, gen17494)

					return

				} else {
					gen17484 := Call(__e, ShenFunc(symcons_2), V4914)

					var gen17485 Obj
					if True == gen17484 {
						gen17481 := Call(__e, ShenFunc(symhd), V4914)

						gen17482 := Call(__e, ShenFunc(sym_a), MakeString("--load"), gen17481)

						var gen17483 Obj
						if True == gen17482 {
							gen17479 := Call(__e, ShenFunc(symtl), V4914)

							gen17480 := Call(__e, ShenFunc(symcons_2), gen17479)

							if True == gen17480 {
								gen17483 = True
							} else {
								gen17483 = False
							}

						} else {
							gen17483 = False
						}
						if True == gen17483 {
							gen17485 = True
						} else {
							gen17485 = False
						}

					} else {
						gen17485 = False
					}
					if True == gen17485 {
						gen17473 := Call(__e, ShenFunc(symtl), V4914)

						gen17474 := Call(__e, ShenFunc(symtl), gen17473)

						gen17477 := MakeNative(func(__e Evaluator) {
							gen17475 := Call(__e, ShenFunc(symtl), V4914)

							gen17476 := Call(__e, ShenFunc(symhd), gen17475)

							__e.TailApply(ShenFunc(symload), gen17476)

							return

						}, 0)
						gen17478 := Call(__e, ShenFunc(symcons), gen17477, V4915)

						__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen17474, gen17478)

						return

					} else {
						gen17471 := Call(__e, ShenFunc(symcons_2), V4914)

						var gen17472 Obj
						if True == gen17471 {
							gen17469 := Call(__e, ShenFunc(symhd), V4914)

							gen17470 := Call(__e, ShenFunc(sym_a), MakeString("--quiet"), gen17469)

							if True == gen17470 {
								gen17472 = True
							} else {
								gen17472 = False
							}

						} else {
							gen17472 = False
						}
						if True == gen17472 {
							gen17466 := Call(__e, ShenFunc(symtl), V4914)

							gen17467 := MakeNative(func(__e Evaluator) {
								__e.TailApply(ShenFunc(symset), MakeSymbol("*hush*"), True)

								return
							}, 0)
							gen17468 := Call(__e, ShenFunc(symcons), gen17467, V4915)

							__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen17466, gen17468)

							return

						} else {
							gen17464 := Call(__e, ShenFunc(symcons_2), V4914)

							var gen17465 Obj
							if True == gen17464 {
								gen17461 := Call(__e, ShenFunc(symhd), V4914)

								gen17462 := Call(__e, ShenFunc(sym_a), MakeString("--set"), gen17461)

								var gen17463 Obj
								if True == gen17462 {
									gen17458 := Call(__e, ShenFunc(symtl), V4914)

									gen17459 := Call(__e, ShenFunc(symcons_2), gen17458)

									var gen17460 Obj
									if True == gen17459 {
										gen17455 := Call(__e, ShenFunc(symtl), V4914)

										gen17456 := Call(__e, ShenFunc(symtl), gen17455)

										gen17457 := Call(__e, ShenFunc(symcons_2), gen17456)

										if True == gen17457 {
											gen17460 = True
										} else {
											gen17460 = False
										}

									} else {
										gen17460 = False
									}
									if True == gen17460 {
										gen17463 = True
									} else {
										gen17463 = False
									}

								} else {
									gen17463 = False
								}
								if True == gen17463 {
									gen17465 = True
								} else {
									gen17465 = False
								}

							} else {
								gen17465 = False
							}
							if True == gen17465 {
								gen17443 := Call(__e, ShenFunc(symtl), V4914)

								gen17444 := Call(__e, ShenFunc(symtl), gen17443)

								gen17445 := Call(__e, ShenFunc(symtl), gen17444)

								gen17453 := MakeNative(func(__e Evaluator) {
									gen17446 := Call(__e, ShenFunc(symtl), V4914)

									gen17447 := Call(__e, ShenFunc(symhd), gen17446)

									gen17448 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1string), gen17447)

									gen17449 := Call(__e, ShenFunc(symtl), V4914)

									gen17450 := Call(__e, ShenFunc(symtl), gen17449)

									gen17451 := Call(__e, ShenFunc(symhd), gen17450)

									gen17452 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1string), gen17451)

									__e.TailApply(ShenFunc(symset), gen17448, gen17452)

									return

								}, 0)
								gen17454 := Call(__e, ShenFunc(symcons), gen17453, V4915)

								__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen17445, gen17454)

								return

							} else {
								gen17441 := Call(__e, ShenFunc(symcons_2), V4914)

								var gen17442 Obj
								if True == gen17441 {
									gen17439 := Call(__e, ShenFunc(symhd), V4914)

									gen17440 := Call(__e, ShenFunc(sym_a), MakeString("--repl"), gen17439)

									if True == gen17440 {
										gen17442 = True
									} else {
										gen17442 = False
									}

								} else {
									gen17442 = False
								}
								if True == gen17442 {
									Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V4915)
									gen17438 := Call(__e, ShenFunc(symtl), V4914)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("launch-repl"), gen17438)

									return

								} else {
									gen17428 := MakeNative(func(__e Evaluator) {
										gen17427 := Call(__e, ShenFunc(symcons_2), V4914)

										if True == gen17427 {
											gen17423 := Call(__e, ShenFunc(symhd), V4914)

											gen17424 := Call(__e, ShenFunc(symshen_4app), gen17423, MakeString(""), MakeSymbol("shen.a"))

											gen17425 := Call(__e, ShenFunc(symcn), MakeString("Invalid eval argument: "), gen17424)

											gen17426 := Call(__e, ShenFunc(symcons), gen17425, Nil)

											__e.TailApply(ShenFunc(symcons), MakeSymbol("error"), gen17426)

											return

										} else {
											__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.launcher.eval-command-h"))

											return
										}

									}, 0)
									Freeze := gen17428
									gen17437 := Call(__e, ShenFunc(symcons_2), V4914)

									if True == gen17437 {
										gen17429 := Call(__e, ShenFunc(symhd), V4914)

										gen17430 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1flag_1map), gen17429)

										Long := gen17430
										gen17433 := Call(__e, ShenFunc(sym_a), False, Long)

										var gen17434 Obj
										if True == gen17433 {
											gen17434 = Call(__e, ShenFunc(symfail))

										} else {
											gen17431 := Call(__e, ShenFunc(symtl), V4914)

											gen17432 := Call(__e, ShenFunc(symcons), Long, gen17431)

											gen17434 = Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen17432, V4915)

										}
										Result := gen17434
										gen17435 := Call(__e, ShenFunc(symfail))

										gen17436 := Call(__e, ShenFunc(sym_a), Result, gen17435)

										if True == gen17436 {
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

							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-command-h"), gen17504)

		gen17505 := MakeNative(func(__e Evaluator) {
			V4917 := __e.Get(1)
			_ = V4917
			__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), V4917, Nil)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-command"), gen17505)

		gen17507 := MakeNative(func(__e Evaluator) {
			V4920 := __e.Get(1)
			_ = V4920
			V4921 := __e.Get(2)
			_ = V4921
			gen17506 := Call(__e, ShenFunc(symcons), V4920, V4921)

			Call(__e, ShenFunc(symset), MakeSymbol("*argv*"), gen17506)

			Call(__e, ShenFunc(symshen_4x_4launcher_4quiet_1load), V4920)
			__e.TailApply(ShenFunc(symcons), MakeSymbol("success"), Nil)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.script-command"), gen17507)

		gen17575 := MakeNative(func(__e Evaluator) {
			V4923 := __e.Get(1)
			_ = V4923
			gen17573 := Call(__e, ShenFunc(symcons_2), V4923)

			var gen17574 Obj
			if True == gen17573 {
				gen17571 := Call(__e, ShenFunc(symtl), V4923)

				gen17572 := Call(__e, ShenFunc(sym_a), Nil, gen17571)

				if True == gen17572 {
					gen17574 = True
				} else {
					gen17574 = False
				}

			} else {
				gen17574 = False
			}
			if True == gen17574 {
				__e.TailApply(ShenFunc(symcons), MakeSymbol("launch-repl"), Nil)

				return
			} else {
				gen17569 := Call(__e, ShenFunc(symcons_2), V4923)

				var gen17570 Obj
				if True == gen17569 {
					gen17566 := Call(__e, ShenFunc(symtl), V4923)

					gen17567 := Call(__e, ShenFunc(symcons_2), gen17566)

					var gen17568 Obj
					if True == gen17567 {
						gen17563 := Call(__e, ShenFunc(symtl), V4923)

						gen17564 := Call(__e, ShenFunc(symhd), gen17563)

						gen17565 := Call(__e, ShenFunc(sym_a), MakeString("--help"), gen17564)

						if True == gen17565 {
							gen17568 = True
						} else {
							gen17568 = False
						}

					} else {
						gen17568 = False
					}
					if True == gen17568 {
						gen17570 = True
					} else {
						gen17570 = False
					}

				} else {
					gen17570 = False
				}
				if True == gen17570 {
					gen17560 := Call(__e, ShenFunc(symhd), V4923)

					gen17561 := Call(__e, ShenFunc(symshen_4x_4launcher_4help_1text), gen17560)

					gen17562 := Call(__e, ShenFunc(symcons), gen17561, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("show-help"), gen17562)

					return

				} else {
					gen17558 := Call(__e, ShenFunc(symcons_2), V4923)

					var gen17559 Obj
					if True == gen17558 {
						gen17555 := Call(__e, ShenFunc(symtl), V4923)

						gen17556 := Call(__e, ShenFunc(symcons_2), gen17555)

						var gen17557 Obj
						if True == gen17556 {
							gen17552 := Call(__e, ShenFunc(symtl), V4923)

							gen17553 := Call(__e, ShenFunc(symhd), gen17552)

							gen17554 := Call(__e, ShenFunc(sym_a), MakeString("--version"), gen17553)

							if True == gen17554 {
								gen17557 = True
							} else {
								gen17557 = False
							}

						} else {
							gen17557 = False
						}
						if True == gen17557 {
							gen17559 = True
						} else {
							gen17559 = False
						}

					} else {
						gen17559 = False
					}
					if True == gen17559 {
						gen17550 := Call(__e, ShenFunc(symshen_4x_4launcher_4version_1string))

						gen17551 := Call(__e, ShenFunc(symcons), gen17550, Nil)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("success"), gen17551)

						return

					} else {
						gen17548 := Call(__e, ShenFunc(symcons_2), V4923)

						var gen17549 Obj
						if True == gen17548 {
							gen17545 := Call(__e, ShenFunc(symtl), V4923)

							gen17546 := Call(__e, ShenFunc(symcons_2), gen17545)

							var gen17547 Obj
							if True == gen17546 {
								gen17542 := Call(__e, ShenFunc(symtl), V4923)

								gen17543 := Call(__e, ShenFunc(symhd), gen17542)

								gen17544 := Call(__e, ShenFunc(sym_a), MakeString("repl"), gen17543)

								if True == gen17544 {
									gen17547 = True
								} else {
									gen17547 = False
								}

							} else {
								gen17547 = False
							}
							if True == gen17547 {
								gen17549 = True
							} else {
								gen17549 = False
							}

						} else {
							gen17549 = False
						}
						if True == gen17549 {
							gen17540 := Call(__e, ShenFunc(symtl), V4923)

							gen17541 := Call(__e, ShenFunc(symtl), gen17540)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("launch-repl"), gen17541)

							return

						} else {
							gen17538 := Call(__e, ShenFunc(symcons_2), V4923)

							var gen17539 Obj
							if True == gen17538 {
								gen17535 := Call(__e, ShenFunc(symtl), V4923)

								gen17536 := Call(__e, ShenFunc(symcons_2), gen17535)

								var gen17537 Obj
								if True == gen17536 {
									gen17531 := Call(__e, ShenFunc(symtl), V4923)

									gen17532 := Call(__e, ShenFunc(symhd), gen17531)

									gen17533 := Call(__e, ShenFunc(sym_a), MakeString("script"), gen17532)

									var gen17534 Obj
									if True == gen17533 {
										gen17528 := Call(__e, ShenFunc(symtl), V4923)

										gen17529 := Call(__e, ShenFunc(symtl), gen17528)

										gen17530 := Call(__e, ShenFunc(symcons_2), gen17529)

										if True == gen17530 {
											gen17534 = True
										} else {
											gen17534 = False
										}

									} else {
										gen17534 = False
									}
									if True == gen17534 {
										gen17537 = True
									} else {
										gen17537 = False
									}

								} else {
									gen17537 = False
								}
								if True == gen17537 {
									gen17539 = True
								} else {
									gen17539 = False
								}

							} else {
								gen17539 = False
							}
							if True == gen17539 {
								gen17522 := Call(__e, ShenFunc(symtl), V4923)

								gen17523 := Call(__e, ShenFunc(symtl), gen17522)

								gen17524 := Call(__e, ShenFunc(symhd), gen17523)

								gen17525 := Call(__e, ShenFunc(symtl), V4923)

								gen17526 := Call(__e, ShenFunc(symtl), gen17525)

								gen17527 := Call(__e, ShenFunc(symtl), gen17526)

								__e.TailApply(ShenFunc(symshen_4x_4launcher_4script_1command), gen17524, gen17527)

								return

							} else {
								gen17520 := Call(__e, ShenFunc(symcons_2), V4923)

								var gen17521 Obj
								if True == gen17520 {
									gen17517 := Call(__e, ShenFunc(symtl), V4923)

									gen17518 := Call(__e, ShenFunc(symcons_2), gen17517)

									var gen17519 Obj
									if True == gen17518 {
										gen17514 := Call(__e, ShenFunc(symtl), V4923)

										gen17515 := Call(__e, ShenFunc(symhd), gen17514)

										gen17516 := Call(__e, ShenFunc(sym_a), MakeString("eval"), gen17515)

										if True == gen17516 {
											gen17519 = True
										} else {
											gen17519 = False
										}

									} else {
										gen17519 = False
									}
									if True == gen17519 {
										gen17521 = True
									} else {
										gen17521 = False
									}

								} else {
									gen17521 = False
								}
								if True == gen17521 {
									gen17512 := Call(__e, ShenFunc(symtl), V4923)

									gen17513 := Call(__e, ShenFunc(symtl), gen17512)

									__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command), gen17513)

									return

								} else {
									gen17510 := Call(__e, ShenFunc(symcons_2), V4923)

									var gen17511 Obj
									if True == gen17510 {
										gen17508 := Call(__e, ShenFunc(symtl), V4923)

										gen17509 := Call(__e, ShenFunc(symcons_2), gen17508)

										if True == gen17509 {
											gen17511 = True
										} else {
											gen17511 = False
										}

									} else {
										gen17511 = False
									}
									if True == gen17511 {
										__e.TailApply(ShenFunc(symcons), MakeSymbol("unknown-arguments"), V4923)

										return
									} else {
										__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.launcher.launch-shen"))

										return
									}

								}

							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.launch-shen"), gen17575)

		gen17654 := MakeNative(func(__e Evaluator) {
			V4927 := __e.Get(1)
			_ = V4927
			gen17652 := Call(__e, ShenFunc(symcons_2), V4927)

			var gen17653 Obj
			if True == gen17652 {
				gen17649 := Call(__e, ShenFunc(symhd), V4927)

				gen17650 := Call(__e, ShenFunc(sym_a), MakeSymbol("success"), gen17649)

				var gen17651 Obj
				if True == gen17650 {
					gen17647 := Call(__e, ShenFunc(symtl), V4927)

					gen17648 := Call(__e, ShenFunc(sym_a), Nil, gen17647)

					if True == gen17648 {
						gen17651 = True
					} else {
						gen17651 = False
					}

				} else {
					gen17651 = False
				}
				if True == gen17651 {
					gen17653 = True
				} else {
					gen17653 = False
				}

			} else {
				gen17653 = False
			}
			if True == gen17653 {
				__e.Return(MakeSymbol("shen.x.launcher.done"))
				return
			} else {
				gen17645 := Call(__e, ShenFunc(symcons_2), V4927)

				var gen17646 Obj
				if True == gen17645 {
					gen17642 := Call(__e, ShenFunc(symhd), V4927)

					gen17643 := Call(__e, ShenFunc(sym_a), MakeSymbol("success"), gen17642)

					var gen17644 Obj
					if True == gen17643 {
						gen17639 := Call(__e, ShenFunc(symtl), V4927)

						gen17640 := Call(__e, ShenFunc(symcons_2), gen17639)

						var gen17641 Obj
						if True == gen17640 {
							gen17636 := Call(__e, ShenFunc(symtl), V4927)

							gen17637 := Call(__e, ShenFunc(symtl), gen17636)

							gen17638 := Call(__e, ShenFunc(sym_a), Nil, gen17637)

							if True == gen17638 {
								gen17641 = True
							} else {
								gen17641 = False
							}

						} else {
							gen17641 = False
						}
						if True == gen17641 {
							gen17644 = True
						} else {
							gen17644 = False
						}

					} else {
						gen17644 = False
					}
					if True == gen17644 {
						gen17646 = True
					} else {
						gen17646 = False
					}

				} else {
					gen17646 = False
				}
				if True == gen17646 {
					gen17632 := Call(__e, ShenFunc(symtl), V4927)

					gen17633 := Call(__e, ShenFunc(symhd), gen17632)

					gen17634 := Call(__e, ShenFunc(symshen_4app), gen17633, MakeString("\n"), MakeSymbol("shen.a"))

					gen17635 := Call(__e, ShenFunc(symstoutput))

					__e.TailApply(ShenFunc(symshen_4prhush), gen17634, gen17635)

					return

				} else {
					gen17630 := Call(__e, ShenFunc(symcons_2), V4927)

					var gen17631 Obj
					if True == gen17630 {
						gen17627 := Call(__e, ShenFunc(symhd), V4927)

						gen17628 := Call(__e, ShenFunc(sym_a), MakeSymbol("error"), gen17627)

						var gen17629 Obj
						if True == gen17628 {
							gen17624 := Call(__e, ShenFunc(symtl), V4927)

							gen17625 := Call(__e, ShenFunc(symcons_2), gen17624)

							var gen17626 Obj
							if True == gen17625 {
								gen17621 := Call(__e, ShenFunc(symtl), V4927)

								gen17622 := Call(__e, ShenFunc(symtl), gen17621)

								gen17623 := Call(__e, ShenFunc(sym_a), Nil, gen17622)

								if True == gen17623 {
									gen17626 = True
								} else {
									gen17626 = False
								}

							} else {
								gen17626 = False
							}
							if True == gen17626 {
								gen17629 = True
							} else {
								gen17629 = False
							}

						} else {
							gen17629 = False
						}
						if True == gen17629 {
							gen17631 = True
						} else {
							gen17631 = False
						}

					} else {
						gen17631 = False
					}
					if True == gen17631 {
						gen17616 := Call(__e, ShenFunc(symtl), V4927)

						gen17617 := Call(__e, ShenFunc(symhd), gen17616)

						gen17618 := Call(__e, ShenFunc(symshen_4app), gen17617, MakeString("\n"), MakeSymbol("shen.a"))

						gen17619 := Call(__e, ShenFunc(symcn), MakeString("ERROR: "), gen17618)

						gen17620 := Call(__e, ShenFunc(symstoutput))

						__e.TailApply(ShenFunc(symshen_4prhush), gen17619, gen17620)

						return

					} else {
						gen17614 := Call(__e, ShenFunc(symcons_2), V4927)

						var gen17615 Obj
						if True == gen17614 {
							gen17612 := Call(__e, ShenFunc(symhd), V4927)

							gen17613 := Call(__e, ShenFunc(sym_a), MakeSymbol("launch-repl"), gen17612)

							if True == gen17613 {
								gen17615 = True
							} else {
								gen17615 = False
							}

						} else {
							gen17615 = False
						}
						if True == gen17615 {
							__e.TailApply(ShenFunc(symshen_4repl))

							return
						} else {
							gen17610 := Call(__e, ShenFunc(symcons_2), V4927)

							var gen17611 Obj
							if True == gen17610 {
								gen17607 := Call(__e, ShenFunc(symhd), V4927)

								gen17608 := Call(__e, ShenFunc(sym_a), MakeSymbol("show-help"), gen17607)

								var gen17609 Obj
								if True == gen17608 {
									gen17604 := Call(__e, ShenFunc(symtl), V4927)

									gen17605 := Call(__e, ShenFunc(symcons_2), gen17604)

									var gen17606 Obj
									if True == gen17605 {
										gen17601 := Call(__e, ShenFunc(symtl), V4927)

										gen17602 := Call(__e, ShenFunc(symtl), gen17601)

										gen17603 := Call(__e, ShenFunc(sym_a), Nil, gen17602)

										if True == gen17603 {
											gen17606 = True
										} else {
											gen17606 = False
										}

									} else {
										gen17606 = False
									}
									if True == gen17606 {
										gen17609 = True
									} else {
										gen17609 = False
									}

								} else {
									gen17609 = False
								}
								if True == gen17609 {
									gen17611 = True
								} else {
									gen17611 = False
								}

							} else {
								gen17611 = False
							}
							if True == gen17611 {
								gen17597 := Call(__e, ShenFunc(symtl), V4927)

								gen17598 := Call(__e, ShenFunc(symhd), gen17597)

								gen17599 := Call(__e, ShenFunc(symshen_4app), gen17598, MakeString("\n"), MakeSymbol("shen.a"))

								gen17600 := Call(__e, ShenFunc(symstoutput))

								__e.TailApply(ShenFunc(symshen_4prhush), gen17599, gen17600)

								return

							} else {
								gen17595 := Call(__e, ShenFunc(symcons_2), V4927)

								var gen17596 Obj
								if True == gen17595 {
									gen17592 := Call(__e, ShenFunc(symhd), V4927)

									gen17593 := Call(__e, ShenFunc(sym_a), MakeSymbol("unknown-arguments"), gen17592)

									var gen17594 Obj
									if True == gen17593 {
										gen17589 := Call(__e, ShenFunc(symtl), V4927)

										gen17590 := Call(__e, ShenFunc(symcons_2), gen17589)

										var gen17591 Obj
										if True == gen17590 {
											gen17586 := Call(__e, ShenFunc(symtl), V4927)

											gen17587 := Call(__e, ShenFunc(symtl), gen17586)

											gen17588 := Call(__e, ShenFunc(symcons_2), gen17587)

											if True == gen17588 {
												gen17591 = True
											} else {
												gen17591 = False
											}

										} else {
											gen17591 = False
										}
										if True == gen17591 {
											gen17594 = True
										} else {
											gen17594 = False
										}

									} else {
										gen17594 = False
									}
									if True == gen17594 {
										gen17596 = True
									} else {
										gen17596 = False
									}

								} else {
									gen17596 = False
								}
								if True == gen17596 {
									gen17576 := Call(__e, ShenFunc(symtl), V4927)

									gen17577 := Call(__e, ShenFunc(symtl), gen17576)

									gen17578 := Call(__e, ShenFunc(symhd), gen17577)

									gen17579 := Call(__e, ShenFunc(symtl), V4927)

									gen17580 := Call(__e, ShenFunc(symhd), gen17579)

									gen17581 := Call(__e, ShenFunc(symshen_4app), gen17580, MakeString(" --help' for more information.\n"), MakeSymbol("shen.a"))

									gen17582 := Call(__e, ShenFunc(symcn), MakeString("\nTry `"), gen17581)

									gen17583 := Call(__e, ShenFunc(symshen_4app), gen17578, gen17582, MakeSymbol("shen.a"))

									gen17584 := Call(__e, ShenFunc(symcn), MakeString("ERROR: Invalid argument: "), gen17583)

									gen17585 := Call(__e, ShenFunc(symstoutput))

									__e.TailApply(ShenFunc(symshen_4prhush), gen17584, gen17585)

									return

								} else {
									__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.launcher.default-handle-result"))

									return
								}

							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.default-handle-result"), gen17654)

		gen17656 := MakeNative(func(__e Evaluator) {
			V4929 := __e.Get(1)
			_ = V4929
			gen17655 := Call(__e, ShenFunc(symshen_4x_4launcher_4launch_1shen), V4929)

			__e.TailApply(ShenFunc(symshen_4x_4launcher_4default_1handle_1result), gen17655)

			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.x.launcher.main"), gen17656)

		return

	}, 0))
}
