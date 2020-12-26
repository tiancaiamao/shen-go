package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

		gen3247 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4891 := __args[0]
			_ = V4891
			gen3245 := Call(__e, ShenFunc(symread_1file), V4891)

			Contents := gen3245
			gen3246 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4eval_1without_1macros), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symmap), gen3246, Contents)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.quiet-load"), gen3247)

		gen3263 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen3248 := Call(__e, ShenFunc(symversion))

			gen3249 := Call(__e, ShenFunc(symlanguage))

			gen3250 := Call(__e, ShenFunc(symport))

			gen3251 := Call(__e, ShenFunc(symcons), gen3250, Nil)

			gen3252 := Call(__e, ShenFunc(symcons), gen3249, gen3251)

			gen3253 := Call(__e, ShenFunc(symimplementation))

			gen3254 := Call(__e, ShenFunc(symrelease))

			gen3255 := Call(__e, ShenFunc(symcons), gen3254, Nil)

			gen3256 := Call(__e, ShenFunc(symcons), gen3253, gen3255)

			gen3257 := Call(__e, ShenFunc(symcons), gen3256, Nil)

			gen3258 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen3257)

			gen3259 := Call(__e, ShenFunc(symcons), gen3252, gen3258)

			gen3260 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen3259)

			gen3261 := Call(__e, ShenFunc(symshen_4app), gen3260, MakeString("\n"), MakeSymbol("shen.r"))

			gen3262 := Call(__e, ShenFunc(symcn), MakeString(" "), gen3261)

			__e.TailApply(ShenFunc(symshen_4app), gen3248, gen3262, MakeSymbol("shen.a"))

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.version-string"), gen3263)

		gen3265 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4893 := __args[0]
			_ = V4893
			gen3264 := Call(__e, ShenFunc(symshen_4app), V4893, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symcn), MakeString("Usage: "), gen3264)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.help-text"), gen3265)

		gen3270 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4895 := __args[0]
			_ = V4895
			gen3269 := Call(__e, ShenFunc(sym_a), Nil, V4895)

			if True == gen3269 {
				__e.TailApply(ShenFunc(symcons), MakeSymbol("success"), Nil)

				return
			} else {
				gen3268 := Call(__e, ShenFunc(symcons_2), V4895)

				if True == gen3268 {
					gen3266 := Call(__e, ShenFunc(symhd), V4895)

					Call(__e, ShenFunc(symthaw), gen3266)

					gen3267 := Call(__e, ShenFunc(symtl), V4895)

					__e.TailApply(ShenFunc(symshen_4x_4launcher_4execute_1all), gen3267)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.launcher.execute-all"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.execute-all"), gen3270)

		gen3273 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4897 := __args[0]
			_ = V4897
			gen3271 := Call(__e, ShenFunc(symread_1from_1string), V4897)

			gen3272 := Call(__e, ShenFunc(symhead), gen3271)

			__e.TailApply(ShenFunc(symeval), gen3272)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-string"), gen3273)

		gen3279 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4903 := __args[0]
			_ = V4903
			gen3278 := Call(__e, ShenFunc(sym_a), MakeString("-e"), V4903)

			if True == gen3278 {
				__e.Return(MakeString("--eval"))
				return
			} else {
				gen3277 := Call(__e, ShenFunc(sym_a), MakeString("-l"), V4903)

				if True == gen3277 {
					__e.Return(MakeString("--load"))
					return
				} else {
					gen3276 := Call(__e, ShenFunc(sym_a), MakeString("-q"), V4903)

					if True == gen3276 {
						__e.Return(MakeString("--quiet"))
						return
					} else {
						gen3275 := Call(__e, ShenFunc(sym_a), MakeString("-s"), V4903)

						if True == gen3275 {
							__e.Return(MakeString("--set"))
							return
						} else {
							gen3274 := Call(__e, ShenFunc(sym_a), MakeString("-r"), V4903)

							if True == gen3274 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-flag-map"), gen3279)

		gen3361 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4914 := __args[0]
			_ = V4914
			V4915 := __args[1]
			_ = V4915
			gen3360 := Call(__e, ShenFunc(sym_a), Nil, V4914)

			if True == gen3360 {
				gen3359 := Call(__e, ShenFunc(symreverse), V4915)

				__e.TailApply(ShenFunc(symshen_4x_4launcher_4execute_1all), gen3359)

				return

			} else {
				gen3357 := Call(__e, ShenFunc(symcons_2), V4914)

				var gen3358 Obj
				if True == gen3357 {
					gen3354 := Call(__e, ShenFunc(symhd), V4914)

					gen3355 := Call(__e, ShenFunc(sym_a), MakeString("--eval"), gen3354)

					var gen3356 Obj
					if True == gen3355 {
						gen3352 := Call(__e, ShenFunc(symtl), V4914)

						gen3353 := Call(__e, ShenFunc(symcons_2), gen3352)

						if True == gen3353 {
							gen3356 = True
						} else {
							gen3356 = False
						}

					} else {
						gen3356 = False
					}
					if True == gen3356 {
						gen3358 = True
					} else {
						gen3358 = False
					}

				} else {
					gen3358 = False
				}
				if True == gen3358 {
					gen3343 := Call(__e, ShenFunc(symtl), V4914)

					gen3344 := Call(__e, ShenFunc(symtl), gen3343)

					gen3350 := MakeNative(func(__e Evaluator, __args ...Obj) {
						gen3345 := Call(__e, ShenFunc(symtl), V4914)

						gen3346 := Call(__e, ShenFunc(symhd), gen3345)

						gen3347 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1string), gen3346)

						gen3348 := Call(__e, ShenFunc(symshen_4app), gen3347, MakeString("\n"), MakeSymbol("shen.a"))

						gen3349 := Call(__e, ShenFunc(symstoutput))

						__e.TailApply(ShenFunc(symshen_4prhush), gen3348, gen3349)

						return

					}, 0)
					gen3351 := Call(__e, ShenFunc(symcons), gen3350, V4915)

					__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen3344, gen3351)

					return

				} else {
					gen3341 := Call(__e, ShenFunc(symcons_2), V4914)

					var gen3342 Obj
					if True == gen3341 {
						gen3338 := Call(__e, ShenFunc(symhd), V4914)

						gen3339 := Call(__e, ShenFunc(sym_a), MakeString("--load"), gen3338)

						var gen3340 Obj
						if True == gen3339 {
							gen3336 := Call(__e, ShenFunc(symtl), V4914)

							gen3337 := Call(__e, ShenFunc(symcons_2), gen3336)

							if True == gen3337 {
								gen3340 = True
							} else {
								gen3340 = False
							}

						} else {
							gen3340 = False
						}
						if True == gen3340 {
							gen3342 = True
						} else {
							gen3342 = False
						}

					} else {
						gen3342 = False
					}
					if True == gen3342 {
						gen3330 := Call(__e, ShenFunc(symtl), V4914)

						gen3331 := Call(__e, ShenFunc(symtl), gen3330)

						gen3334 := MakeNative(func(__e Evaluator, __args ...Obj) {
							gen3332 := Call(__e, ShenFunc(symtl), V4914)

							gen3333 := Call(__e, ShenFunc(symhd), gen3332)

							__e.TailApply(ShenFunc(symload), gen3333)

							return

						}, 0)
						gen3335 := Call(__e, ShenFunc(symcons), gen3334, V4915)

						__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen3331, gen3335)

						return

					} else {
						gen3328 := Call(__e, ShenFunc(symcons_2), V4914)

						var gen3329 Obj
						if True == gen3328 {
							gen3326 := Call(__e, ShenFunc(symhd), V4914)

							gen3327 := Call(__e, ShenFunc(sym_a), MakeString("--quiet"), gen3326)

							if True == gen3327 {
								gen3329 = True
							} else {
								gen3329 = False
							}

						} else {
							gen3329 = False
						}
						if True == gen3329 {
							gen3323 := Call(__e, ShenFunc(symtl), V4914)

							gen3324 := MakeNative(func(__e Evaluator, __args ...Obj) {
								__e.TailApply(ShenFunc(symset), MakeSymbol("*hush*"), True)

								return
							}, 0)
							gen3325 := Call(__e, ShenFunc(symcons), gen3324, V4915)

							__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen3323, gen3325)

							return

						} else {
							gen3321 := Call(__e, ShenFunc(symcons_2), V4914)

							var gen3322 Obj
							if True == gen3321 {
								gen3318 := Call(__e, ShenFunc(symhd), V4914)

								gen3319 := Call(__e, ShenFunc(sym_a), MakeString("--set"), gen3318)

								var gen3320 Obj
								if True == gen3319 {
									gen3315 := Call(__e, ShenFunc(symtl), V4914)

									gen3316 := Call(__e, ShenFunc(symcons_2), gen3315)

									var gen3317 Obj
									if True == gen3316 {
										gen3312 := Call(__e, ShenFunc(symtl), V4914)

										gen3313 := Call(__e, ShenFunc(symtl), gen3312)

										gen3314 := Call(__e, ShenFunc(symcons_2), gen3313)

										if True == gen3314 {
											gen3317 = True
										} else {
											gen3317 = False
										}

									} else {
										gen3317 = False
									}
									if True == gen3317 {
										gen3320 = True
									} else {
										gen3320 = False
									}

								} else {
									gen3320 = False
								}
								if True == gen3320 {
									gen3322 = True
								} else {
									gen3322 = False
								}

							} else {
								gen3322 = False
							}
							if True == gen3322 {
								gen3300 := Call(__e, ShenFunc(symtl), V4914)

								gen3301 := Call(__e, ShenFunc(symtl), gen3300)

								gen3302 := Call(__e, ShenFunc(symtl), gen3301)

								gen3310 := MakeNative(func(__e Evaluator, __args ...Obj) {
									gen3303 := Call(__e, ShenFunc(symtl), V4914)

									gen3304 := Call(__e, ShenFunc(symhd), gen3303)

									gen3305 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1string), gen3304)

									gen3306 := Call(__e, ShenFunc(symtl), V4914)

									gen3307 := Call(__e, ShenFunc(symtl), gen3306)

									gen3308 := Call(__e, ShenFunc(symhd), gen3307)

									gen3309 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1string), gen3308)

									__e.TailApply(ShenFunc(symset), gen3305, gen3309)

									return

								}, 0)
								gen3311 := Call(__e, ShenFunc(symcons), gen3310, V4915)

								__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen3302, gen3311)

								return

							} else {
								gen3298 := Call(__e, ShenFunc(symcons_2), V4914)

								var gen3299 Obj
								if True == gen3298 {
									gen3296 := Call(__e, ShenFunc(symhd), V4914)

									gen3297 := Call(__e, ShenFunc(sym_a), MakeString("--repl"), gen3296)

									if True == gen3297 {
										gen3299 = True
									} else {
										gen3299 = False
									}

								} else {
									gen3299 = False
								}
								if True == gen3299 {
									Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V4915)
									gen3295 := Call(__e, ShenFunc(symtl), V4914)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("launch-repl"), gen3295)

									return

								} else {
									gen3285 := MakeNative(func(__e Evaluator, __args ...Obj) {
										gen3284 := Call(__e, ShenFunc(symcons_2), V4914)

										if True == gen3284 {
											gen3280 := Call(__e, ShenFunc(symhd), V4914)

											gen3281 := Call(__e, ShenFunc(symshen_4app), gen3280, MakeString(""), MakeSymbol("shen.a"))

											gen3282 := Call(__e, ShenFunc(symcn), MakeString("Invalid eval argument: "), gen3281)

											gen3283 := Call(__e, ShenFunc(symcons), gen3282, Nil)

											__e.TailApply(ShenFunc(symcons), MakeSymbol("error"), gen3283)

											return

										} else {
											__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.x.launcher.eval-command-h"))

											return
										}

									}, 0)
									Freeze := gen3285
									gen3294 := Call(__e, ShenFunc(symcons_2), V4914)

									if True == gen3294 {
										gen3286 := Call(__e, ShenFunc(symhd), V4914)

										gen3287 := Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1flag_1map), gen3286)

										Long := gen3287
										gen3290 := Call(__e, ShenFunc(sym_a), False, Long)

										var gen3291 Obj
										if True == gen3290 {
											gen3291 = Call(__e, ShenFunc(symfail))

										} else {
											gen3288 := Call(__e, ShenFunc(symtl), V4914)

											gen3289 := Call(__e, ShenFunc(symcons), Long, gen3288)

											gen3291 = Call(__e, ShenFunc(symshen_4x_4launcher_4eval_1command_1h), gen3289, V4915)

										}
										Result := gen3291
										gen3292 := Call(__e, ShenFunc(symfail))

										gen3293 := Call(__e, ShenFunc(sym_a), Result, gen3292)

										if True == gen3293 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-command-h"), gen3361)

		gen3362 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4917 := __args[0]
			_ = V4917
			__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command_1h), V4917, Nil)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.eval-command"), gen3362)

		gen3364 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4920 := __args[0]
			_ = V4920
			V4921 := __args[1]
			_ = V4921
			gen3363 := Call(__e, ShenFunc(symcons), V4920, V4921)

			Call(__e, ShenFunc(symset), MakeSymbol("*argv*"), gen3363)

			Call(__e, ShenFunc(symshen_4x_4launcher_4quiet_1load), V4920)
			__e.TailApply(ShenFunc(symcons), MakeSymbol("success"), Nil)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.script-command"), gen3364)

		gen3432 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4923 := __args[0]
			_ = V4923
			gen3430 := Call(__e, ShenFunc(symcons_2), V4923)

			var gen3431 Obj
			if True == gen3430 {
				gen3428 := Call(__e, ShenFunc(symtl), V4923)

				gen3429 := Call(__e, ShenFunc(sym_a), Nil, gen3428)

				if True == gen3429 {
					gen3431 = True
				} else {
					gen3431 = False
				}

			} else {
				gen3431 = False
			}
			if True == gen3431 {
				__e.TailApply(ShenFunc(symcons), MakeSymbol("launch-repl"), Nil)

				return
			} else {
				gen3426 := Call(__e, ShenFunc(symcons_2), V4923)

				var gen3427 Obj
				if True == gen3426 {
					gen3423 := Call(__e, ShenFunc(symtl), V4923)

					gen3424 := Call(__e, ShenFunc(symcons_2), gen3423)

					var gen3425 Obj
					if True == gen3424 {
						gen3420 := Call(__e, ShenFunc(symtl), V4923)

						gen3421 := Call(__e, ShenFunc(symhd), gen3420)

						gen3422 := Call(__e, ShenFunc(sym_a), MakeString("--help"), gen3421)

						if True == gen3422 {
							gen3425 = True
						} else {
							gen3425 = False
						}

					} else {
						gen3425 = False
					}
					if True == gen3425 {
						gen3427 = True
					} else {
						gen3427 = False
					}

				} else {
					gen3427 = False
				}
				if True == gen3427 {
					gen3417 := Call(__e, ShenFunc(symhd), V4923)

					gen3418 := Call(__e, ShenFunc(symshen_4x_4launcher_4help_1text), gen3417)

					gen3419 := Call(__e, ShenFunc(symcons), gen3418, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("show-help"), gen3419)

					return

				} else {
					gen3415 := Call(__e, ShenFunc(symcons_2), V4923)

					var gen3416 Obj
					if True == gen3415 {
						gen3412 := Call(__e, ShenFunc(symtl), V4923)

						gen3413 := Call(__e, ShenFunc(symcons_2), gen3412)

						var gen3414 Obj
						if True == gen3413 {
							gen3409 := Call(__e, ShenFunc(symtl), V4923)

							gen3410 := Call(__e, ShenFunc(symhd), gen3409)

							gen3411 := Call(__e, ShenFunc(sym_a), MakeString("--version"), gen3410)

							if True == gen3411 {
								gen3414 = True
							} else {
								gen3414 = False
							}

						} else {
							gen3414 = False
						}
						if True == gen3414 {
							gen3416 = True
						} else {
							gen3416 = False
						}

					} else {
						gen3416 = False
					}
					if True == gen3416 {
						gen3407 := Call(__e, ShenFunc(symshen_4x_4launcher_4version_1string))

						gen3408 := Call(__e, ShenFunc(symcons), gen3407, Nil)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("success"), gen3408)

						return

					} else {
						gen3405 := Call(__e, ShenFunc(symcons_2), V4923)

						var gen3406 Obj
						if True == gen3405 {
							gen3402 := Call(__e, ShenFunc(symtl), V4923)

							gen3403 := Call(__e, ShenFunc(symcons_2), gen3402)

							var gen3404 Obj
							if True == gen3403 {
								gen3399 := Call(__e, ShenFunc(symtl), V4923)

								gen3400 := Call(__e, ShenFunc(symhd), gen3399)

								gen3401 := Call(__e, ShenFunc(sym_a), MakeString("repl"), gen3400)

								if True == gen3401 {
									gen3404 = True
								} else {
									gen3404 = False
								}

							} else {
								gen3404 = False
							}
							if True == gen3404 {
								gen3406 = True
							} else {
								gen3406 = False
							}

						} else {
							gen3406 = False
						}
						if True == gen3406 {
							gen3397 := Call(__e, ShenFunc(symtl), V4923)

							gen3398 := Call(__e, ShenFunc(symtl), gen3397)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("launch-repl"), gen3398)

							return

						} else {
							gen3395 := Call(__e, ShenFunc(symcons_2), V4923)

							var gen3396 Obj
							if True == gen3395 {
								gen3392 := Call(__e, ShenFunc(symtl), V4923)

								gen3393 := Call(__e, ShenFunc(symcons_2), gen3392)

								var gen3394 Obj
								if True == gen3393 {
									gen3388 := Call(__e, ShenFunc(symtl), V4923)

									gen3389 := Call(__e, ShenFunc(symhd), gen3388)

									gen3390 := Call(__e, ShenFunc(sym_a), MakeString("script"), gen3389)

									var gen3391 Obj
									if True == gen3390 {
										gen3385 := Call(__e, ShenFunc(symtl), V4923)

										gen3386 := Call(__e, ShenFunc(symtl), gen3385)

										gen3387 := Call(__e, ShenFunc(symcons_2), gen3386)

										if True == gen3387 {
											gen3391 = True
										} else {
											gen3391 = False
										}

									} else {
										gen3391 = False
									}
									if True == gen3391 {
										gen3394 = True
									} else {
										gen3394 = False
									}

								} else {
									gen3394 = False
								}
								if True == gen3394 {
									gen3396 = True
								} else {
									gen3396 = False
								}

							} else {
								gen3396 = False
							}
							if True == gen3396 {
								gen3379 := Call(__e, ShenFunc(symtl), V4923)

								gen3380 := Call(__e, ShenFunc(symtl), gen3379)

								gen3381 := Call(__e, ShenFunc(symhd), gen3380)

								gen3382 := Call(__e, ShenFunc(symtl), V4923)

								gen3383 := Call(__e, ShenFunc(symtl), gen3382)

								gen3384 := Call(__e, ShenFunc(symtl), gen3383)

								__e.TailApply(ShenFunc(symshen_4x_4launcher_4script_1command), gen3381, gen3384)

								return

							} else {
								gen3377 := Call(__e, ShenFunc(symcons_2), V4923)

								var gen3378 Obj
								if True == gen3377 {
									gen3374 := Call(__e, ShenFunc(symtl), V4923)

									gen3375 := Call(__e, ShenFunc(symcons_2), gen3374)

									var gen3376 Obj
									if True == gen3375 {
										gen3371 := Call(__e, ShenFunc(symtl), V4923)

										gen3372 := Call(__e, ShenFunc(symhd), gen3371)

										gen3373 := Call(__e, ShenFunc(sym_a), MakeString("eval"), gen3372)

										if True == gen3373 {
											gen3376 = True
										} else {
											gen3376 = False
										}

									} else {
										gen3376 = False
									}
									if True == gen3376 {
										gen3378 = True
									} else {
										gen3378 = False
									}

								} else {
									gen3378 = False
								}
								if True == gen3378 {
									gen3369 := Call(__e, ShenFunc(symtl), V4923)

									gen3370 := Call(__e, ShenFunc(symtl), gen3369)

									__e.TailApply(ShenFunc(symshen_4x_4launcher_4eval_1command), gen3370)

									return

								} else {
									gen3367 := Call(__e, ShenFunc(symcons_2), V4923)

									var gen3368 Obj
									if True == gen3367 {
										gen3365 := Call(__e, ShenFunc(symtl), V4923)

										gen3366 := Call(__e, ShenFunc(symcons_2), gen3365)

										if True == gen3366 {
											gen3368 = True
										} else {
											gen3368 = False
										}

									} else {
										gen3368 = False
									}
									if True == gen3368 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.launch-shen"), gen3432)

		gen3511 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4927 := __args[0]
			_ = V4927
			gen3509 := Call(__e, ShenFunc(symcons_2), V4927)

			var gen3510 Obj
			if True == gen3509 {
				gen3506 := Call(__e, ShenFunc(symhd), V4927)

				gen3507 := Call(__e, ShenFunc(sym_a), MakeSymbol("success"), gen3506)

				var gen3508 Obj
				if True == gen3507 {
					gen3504 := Call(__e, ShenFunc(symtl), V4927)

					gen3505 := Call(__e, ShenFunc(sym_a), Nil, gen3504)

					if True == gen3505 {
						gen3508 = True
					} else {
						gen3508 = False
					}

				} else {
					gen3508 = False
				}
				if True == gen3508 {
					gen3510 = True
				} else {
					gen3510 = False
				}

			} else {
				gen3510 = False
			}
			if True == gen3510 {
				__e.Return(MakeSymbol("shen.x.launcher.done"))
				return
			} else {
				gen3502 := Call(__e, ShenFunc(symcons_2), V4927)

				var gen3503 Obj
				if True == gen3502 {
					gen3499 := Call(__e, ShenFunc(symhd), V4927)

					gen3500 := Call(__e, ShenFunc(sym_a), MakeSymbol("success"), gen3499)

					var gen3501 Obj
					if True == gen3500 {
						gen3496 := Call(__e, ShenFunc(symtl), V4927)

						gen3497 := Call(__e, ShenFunc(symcons_2), gen3496)

						var gen3498 Obj
						if True == gen3497 {
							gen3493 := Call(__e, ShenFunc(symtl), V4927)

							gen3494 := Call(__e, ShenFunc(symtl), gen3493)

							gen3495 := Call(__e, ShenFunc(sym_a), Nil, gen3494)

							if True == gen3495 {
								gen3498 = True
							} else {
								gen3498 = False
							}

						} else {
							gen3498 = False
						}
						if True == gen3498 {
							gen3501 = True
						} else {
							gen3501 = False
						}

					} else {
						gen3501 = False
					}
					if True == gen3501 {
						gen3503 = True
					} else {
						gen3503 = False
					}

				} else {
					gen3503 = False
				}
				if True == gen3503 {
					gen3489 := Call(__e, ShenFunc(symtl), V4927)

					gen3490 := Call(__e, ShenFunc(symhd), gen3489)

					gen3491 := Call(__e, ShenFunc(symshen_4app), gen3490, MakeString("\n"), MakeSymbol("shen.a"))

					gen3492 := Call(__e, ShenFunc(symstoutput))

					__e.TailApply(ShenFunc(symshen_4prhush), gen3491, gen3492)

					return

				} else {
					gen3487 := Call(__e, ShenFunc(symcons_2), V4927)

					var gen3488 Obj
					if True == gen3487 {
						gen3484 := Call(__e, ShenFunc(symhd), V4927)

						gen3485 := Call(__e, ShenFunc(sym_a), MakeSymbol("error"), gen3484)

						var gen3486 Obj
						if True == gen3485 {
							gen3481 := Call(__e, ShenFunc(symtl), V4927)

							gen3482 := Call(__e, ShenFunc(symcons_2), gen3481)

							var gen3483 Obj
							if True == gen3482 {
								gen3478 := Call(__e, ShenFunc(symtl), V4927)

								gen3479 := Call(__e, ShenFunc(symtl), gen3478)

								gen3480 := Call(__e, ShenFunc(sym_a), Nil, gen3479)

								if True == gen3480 {
									gen3483 = True
								} else {
									gen3483 = False
								}

							} else {
								gen3483 = False
							}
							if True == gen3483 {
								gen3486 = True
							} else {
								gen3486 = False
							}

						} else {
							gen3486 = False
						}
						if True == gen3486 {
							gen3488 = True
						} else {
							gen3488 = False
						}

					} else {
						gen3488 = False
					}
					if True == gen3488 {
						gen3473 := Call(__e, ShenFunc(symtl), V4927)

						gen3474 := Call(__e, ShenFunc(symhd), gen3473)

						gen3475 := Call(__e, ShenFunc(symshen_4app), gen3474, MakeString("\n"), MakeSymbol("shen.a"))

						gen3476 := Call(__e, ShenFunc(symcn), MakeString("ERROR: "), gen3475)

						gen3477 := Call(__e, ShenFunc(symstoutput))

						__e.TailApply(ShenFunc(symshen_4prhush), gen3476, gen3477)

						return

					} else {
						gen3471 := Call(__e, ShenFunc(symcons_2), V4927)

						var gen3472 Obj
						if True == gen3471 {
							gen3469 := Call(__e, ShenFunc(symhd), V4927)

							gen3470 := Call(__e, ShenFunc(sym_a), MakeSymbol("launch-repl"), gen3469)

							if True == gen3470 {
								gen3472 = True
							} else {
								gen3472 = False
							}

						} else {
							gen3472 = False
						}
						if True == gen3472 {
							__e.TailApply(ShenFunc(symshen_4repl))

							return
						} else {
							gen3467 := Call(__e, ShenFunc(symcons_2), V4927)

							var gen3468 Obj
							if True == gen3467 {
								gen3464 := Call(__e, ShenFunc(symhd), V4927)

								gen3465 := Call(__e, ShenFunc(sym_a), MakeSymbol("show-help"), gen3464)

								var gen3466 Obj
								if True == gen3465 {
									gen3461 := Call(__e, ShenFunc(symtl), V4927)

									gen3462 := Call(__e, ShenFunc(symcons_2), gen3461)

									var gen3463 Obj
									if True == gen3462 {
										gen3458 := Call(__e, ShenFunc(symtl), V4927)

										gen3459 := Call(__e, ShenFunc(symtl), gen3458)

										gen3460 := Call(__e, ShenFunc(sym_a), Nil, gen3459)

										if True == gen3460 {
											gen3463 = True
										} else {
											gen3463 = False
										}

									} else {
										gen3463 = False
									}
									if True == gen3463 {
										gen3466 = True
									} else {
										gen3466 = False
									}

								} else {
									gen3466 = False
								}
								if True == gen3466 {
									gen3468 = True
								} else {
									gen3468 = False
								}

							} else {
								gen3468 = False
							}
							if True == gen3468 {
								gen3454 := Call(__e, ShenFunc(symtl), V4927)

								gen3455 := Call(__e, ShenFunc(symhd), gen3454)

								gen3456 := Call(__e, ShenFunc(symshen_4app), gen3455, MakeString("\n"), MakeSymbol("shen.a"))

								gen3457 := Call(__e, ShenFunc(symstoutput))

								__e.TailApply(ShenFunc(symshen_4prhush), gen3456, gen3457)

								return

							} else {
								gen3452 := Call(__e, ShenFunc(symcons_2), V4927)

								var gen3453 Obj
								if True == gen3452 {
									gen3449 := Call(__e, ShenFunc(symhd), V4927)

									gen3450 := Call(__e, ShenFunc(sym_a), MakeSymbol("unknown-arguments"), gen3449)

									var gen3451 Obj
									if True == gen3450 {
										gen3446 := Call(__e, ShenFunc(symtl), V4927)

										gen3447 := Call(__e, ShenFunc(symcons_2), gen3446)

										var gen3448 Obj
										if True == gen3447 {
											gen3443 := Call(__e, ShenFunc(symtl), V4927)

											gen3444 := Call(__e, ShenFunc(symtl), gen3443)

											gen3445 := Call(__e, ShenFunc(symcons_2), gen3444)

											if True == gen3445 {
												gen3448 = True
											} else {
												gen3448 = False
											}

										} else {
											gen3448 = False
										}
										if True == gen3448 {
											gen3451 = True
										} else {
											gen3451 = False
										}

									} else {
										gen3451 = False
									}
									if True == gen3451 {
										gen3453 = True
									} else {
										gen3453 = False
									}

								} else {
									gen3453 = False
								}
								if True == gen3453 {
									gen3433 := Call(__e, ShenFunc(symtl), V4927)

									gen3434 := Call(__e, ShenFunc(symtl), gen3433)

									gen3435 := Call(__e, ShenFunc(symhd), gen3434)

									gen3436 := Call(__e, ShenFunc(symtl), V4927)

									gen3437 := Call(__e, ShenFunc(symhd), gen3436)

									gen3438 := Call(__e, ShenFunc(symshen_4app), gen3437, MakeString(" --help' for more information.\n"), MakeSymbol("shen.a"))

									gen3439 := Call(__e, ShenFunc(symcn), MakeString("\nTry `"), gen3438)

									gen3440 := Call(__e, ShenFunc(symshen_4app), gen3435, gen3439, MakeSymbol("shen.a"))

									gen3441 := Call(__e, ShenFunc(symcn), MakeString("ERROR: Invalid argument: "), gen3440)

									gen3442 := Call(__e, ShenFunc(symstoutput))

									__e.TailApply(ShenFunc(symshen_4prhush), gen3441, gen3442)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.launcher.default-handle-result"), gen3511)

		gen3513 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4929 := __args[0]
			_ = V4929
			gen3512 := Call(__e, ShenFunc(symshen_4x_4launcher_4launch_1shen), V4929)

			__e.TailApply(ShenFunc(symshen_4x_4launcher_4default_1handle_1result), gen3512)

			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.x.launcher.main"), gen3513)

		return

	}, 0))
}
