package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionFactoriseDefunMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2012-2019 Bruno Deferrari.  All rights reserved.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

	tmp3176 := MakeNative(func(__e *ControlFlow) {
		V4931 := __e.Get(1)
		_ = V4931
		tmp3239 := PrimIsPair(V4931)

		var ifres3198 Obj

		if True == tmp3239 {
			tmp3237 := PrimHead(V4931)

			tmp3238 := PrimEqual(symdefun, tmp3237)

			var ifres3200 Obj

			if True == tmp3238 {
				tmp3235 := PrimTail(V4931)

				tmp3236 := PrimIsPair(tmp3235)

				var ifres3202 Obj

				if True == tmp3236 {
					tmp3232 := PrimTail(V4931)

					tmp3233 := PrimTail(tmp3232)

					tmp3234 := PrimIsPair(tmp3233)

					var ifres3204 Obj

					if True == tmp3234 {
						tmp3228 := PrimTail(V4931)

						tmp3229 := PrimTail(tmp3228)

						tmp3230 := PrimTail(tmp3229)

						tmp3231 := PrimIsPair(tmp3230)

						var ifres3206 Obj

						if True == tmp3231 {
							tmp3223 := PrimTail(V4931)

							tmp3224 := PrimTail(tmp3223)

							tmp3225 := PrimTail(tmp3224)

							tmp3226 := PrimHead(tmp3225)

							tmp3227 := PrimIsPair(tmp3226)

							var ifres3208 Obj

							if True == tmp3227 {
								tmp3217 := PrimTail(V4931)

								tmp3218 := PrimTail(tmp3217)

								tmp3219 := PrimTail(tmp3218)

								tmp3220 := PrimHead(tmp3219)

								tmp3221 := PrimHead(tmp3220)

								tmp3222 := PrimEqual(symcond, tmp3221)

								var ifres3210 Obj

								if True == tmp3222 {
									tmp3212 := PrimTail(V4931)

									tmp3213 := PrimTail(tmp3212)

									tmp3214 := PrimTail(tmp3213)

									tmp3215 := PrimTail(tmp3214)

									tmp3216 := PrimEqual(Nil, tmp3215)

									var ifres3211 Obj

									if True == tmp3216 {
										ifres3211 = True

									} else {
										ifres3211 = False

									}

									ifres3210 = ifres3211

								} else {
									ifres3210 = False

								}

								var ifres3209 Obj

								if True == ifres3210 {
									ifres3209 = True

								} else {
									ifres3209 = False

								}

								ifres3208 = ifres3209

							} else {
								ifres3208 = False

							}

							var ifres3207 Obj

							if True == ifres3208 {
								ifres3207 = True

							} else {
								ifres3207 = False

							}

							ifres3206 = ifres3207

						} else {
							ifres3206 = False

						}

						var ifres3205 Obj

						if True == ifres3206 {
							ifres3205 = True

						} else {
							ifres3205 = False

						}

						ifres3204 = ifres3205

					} else {
						ifres3204 = False

					}

					var ifres3203 Obj

					if True == ifres3204 {
						ifres3203 = True

					} else {
						ifres3203 = False

					}

					ifres3202 = ifres3203

				} else {
					ifres3202 = False

				}

				var ifres3201 Obj

				if True == ifres3202 {
					ifres3201 = True

				} else {
					ifres3201 = False

				}

				ifres3200 = ifres3201

			} else {
				ifres3200 = False

			}

			var ifres3199 Obj

			if True == ifres3200 {
				ifres3199 = True

			} else {
				ifres3199 = False

			}

			ifres3198 = ifres3199

		} else {
			ifres3198 = False

		}

		if True == ifres3198 {
			tmp3178 := PrimTail(V4931)

			tmp3179 := PrimHead(tmp3178)

			tmp3180 := PrimTail(V4931)

			tmp3181 := PrimTail(tmp3180)

			tmp3182 := PrimHead(tmp3181)

			tmp3183 := PrimTail(V4931)

			tmp3184 := PrimTail(tmp3183)

			tmp3185 := PrimTail(tmp3184)

			tmp3186 := PrimHead(tmp3185)

			tmp3187 := PrimTail(V4931)

			tmp3188 := PrimHead(tmp3187)

			tmp3189 := PrimCons(tmp3188, Nil)

			tmp3190 := PrimCons(symshen_4f__error, tmp3189)

			tmp3191 := PrimTail(V4931)

			tmp3192 := PrimTail(tmp3191)

			tmp3193 := PrimHead(tmp3192)

			tmp3194 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4factorise_1cond), tmp3186, tmp3190, tmp3193)

			tmp3195 := PrimCons(tmp3194, Nil)

			tmp3196 := PrimCons(tmp3182, tmp3195)

			tmp3197 := PrimCons(tmp3179, tmp3196)

			__e.Return(PrimCons(symdefun, tmp3197))
			return

		} else {
			__e.Return(V4931)
			return
		}

	}, 1)

	tmp3240 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4factorise_1defun, tmp3176)

	_ = tmp3240

	tmp3241 := MakeNative(func(__e *ControlFlow) {
		V4943 := __e.Get(1)
		_ = V4943
		V4944 := __e.Get(2)
		_ = V4944
		V4945 := __e.Get(3)
		_ = V4945
		tmp3250 := PrimIsPair(V4943)

		var ifres3246 Obj

		if True == tmp3250 {
			tmp3248 := PrimHead(V4943)

			tmp3249 := PrimEqual(symcond, tmp3248)

			var ifres3247 Obj

			if True == tmp3249 {
				ifres3247 = True

			} else {
				ifres3247 = False

			}

			ifres3246 = ifres3247

		} else {
			ifres3246 = False

		}

		if True == ifres3246 {
			tmp3243 := PrimTail(V4943)

			tmp3244 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4add_1returns), tmp3243)

			tmp3245 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4rebranch), tmp3244, V4944)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3245, V4945)
			return

		} else {
			__e.Return(V4943)
			return
		}

	}, 3)

	tmp3251 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4factorise_1cond, tmp3241)

	_ = tmp3251

	tmp3252 := MakeNative(func(__e *ControlFlow) {
		V4947 := __e.Get(1)
		_ = V4947
		tmp3280 := PrimEqual(Nil, V4947)

		if True == tmp3280 {
			__e.Return(Nil)
			return
		} else {
			tmp3279 := PrimIsPair(V4947)

			var ifres3264 Obj

			if True == tmp3279 {
				tmp3277 := PrimHead(V4947)

				tmp3278 := PrimIsPair(tmp3277)

				var ifres3266 Obj

				if True == tmp3278 {
					tmp3274 := PrimHead(V4947)

					tmp3275 := PrimTail(tmp3274)

					tmp3276 := PrimIsPair(tmp3275)

					var ifres3268 Obj

					if True == tmp3276 {
						tmp3270 := PrimHead(V4947)

						tmp3271 := PrimTail(tmp3270)

						tmp3272 := PrimTail(tmp3271)

						tmp3273 := PrimEqual(Nil, tmp3272)

						var ifres3269 Obj

						if True == tmp3273 {
							ifres3269 = True

						} else {
							ifres3269 = False

						}

						ifres3268 = ifres3269

					} else {
						ifres3268 = False

					}

					var ifres3267 Obj

					if True == ifres3268 {
						ifres3267 = True

					} else {
						ifres3267 = False

					}

					ifres3266 = ifres3267

				} else {
					ifres3266 = False

				}

				var ifres3265 Obj

				if True == ifres3266 {
					ifres3265 = True

				} else {
					ifres3265 = False

				}

				ifres3264 = ifres3265

			} else {
				ifres3264 = False

			}

			if True == ifres3264 {
				tmp3255 := PrimHead(V4947)

				tmp3256 := PrimHead(tmp3255)

				tmp3257 := PrimHead(V4947)

				tmp3258 := PrimTail(tmp3257)

				tmp3259 := PrimCons(sym_f_freturn, tmp3258)

				tmp3260 := PrimCons(tmp3259, Nil)

				tmp3261 := PrimCons(tmp3256, tmp3260)

				tmp3262 := PrimTail(V4947)

				tmp3263 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4add_1returns), tmp3262)

				__e.Return(PrimCons(tmp3261, tmp3263))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4factorise_1defun_4add_1returns)
				return
			}

		}

	}, 1)

	tmp3281 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4add_1returns, tmp3252)

	_ = tmp3281

	tmp3282 := MakeNative(func(__e *ControlFlow) {
		__e.TailApply(PrimNS2Value(symgensym), sym_f_flabel)
		return
	}, 0)

	tmp3283 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4generate_1label, tmp3282)

	_ = tmp3283

	tmp3284 := MakeNative(func(__e *ControlFlow) {
		V4950 := __e.Get(1)
		_ = V4950
		V4951 := __e.Get(2)
		_ = V4951
		tmp3285 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables_1h), V4950, V4951, Nil)

		__e.TailApply(PrimNS2Value(symreverse), tmp3285)
		return

	}, 2)

	tmp3286 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4free_1variables, tmp3284)

	_ = tmp3286

	tmp3287 := MakeNative(func(__e *ControlFlow) {
		V4963 := __e.Get(1)
		_ = V4963
		V4964 := __e.Get(2)
		_ = V4964
		V4965 := __e.Get(3)
		_ = V4965
		tmp3360 := PrimIsPair(V4963)

		var ifres3334 Obj

		if True == tmp3360 {
			tmp3358 := PrimHead(V4963)

			tmp3359 := PrimEqual(symlet, tmp3358)

			var ifres3336 Obj

			if True == tmp3359 {
				tmp3356 := PrimTail(V4963)

				tmp3357 := PrimIsPair(tmp3356)

				var ifres3338 Obj

				if True == tmp3357 {
					tmp3353 := PrimTail(V4963)

					tmp3354 := PrimTail(tmp3353)

					tmp3355 := PrimIsPair(tmp3354)

					var ifres3340 Obj

					if True == tmp3355 {
						tmp3349 := PrimTail(V4963)

						tmp3350 := PrimTail(tmp3349)

						tmp3351 := PrimTail(tmp3350)

						tmp3352 := PrimIsPair(tmp3351)

						var ifres3342 Obj

						if True == tmp3352 {
							tmp3344 := PrimTail(V4963)

							tmp3345 := PrimTail(tmp3344)

							tmp3346 := PrimTail(tmp3345)

							tmp3347 := PrimTail(tmp3346)

							tmp3348 := PrimEqual(Nil, tmp3347)

							var ifres3343 Obj

							if True == tmp3348 {
								ifres3343 = True

							} else {
								ifres3343 = False

							}

							ifres3342 = ifres3343

						} else {
							ifres3342 = False

						}

						var ifres3341 Obj

						if True == ifres3342 {
							ifres3341 = True

						} else {
							ifres3341 = False

						}

						ifres3340 = ifres3341

					} else {
						ifres3340 = False

					}

					var ifres3339 Obj

					if True == ifres3340 {
						ifres3339 = True

					} else {
						ifres3339 = False

					}

					ifres3338 = ifres3339

				} else {
					ifres3338 = False

				}

				var ifres3337 Obj

				if True == ifres3338 {
					ifres3337 = True

				} else {
					ifres3337 = False

				}

				ifres3336 = ifres3337

			} else {
				ifres3336 = False

			}

			var ifres3335 Obj

			if True == ifres3336 {
				ifres3335 = True

			} else {
				ifres3335 = False

			}

			ifres3334 = ifres3335

		} else {
			ifres3334 = False

		}

		if True == ifres3334 {
			tmp3323 := PrimTail(V4963)

			tmp3324 := PrimTail(tmp3323)

			tmp3325 := PrimTail(tmp3324)

			tmp3326 := PrimHead(tmp3325)

			tmp3327 := PrimTail(V4963)

			tmp3328 := PrimHead(tmp3327)

			tmp3329 := Call(__e, PrimNS2Value(symremove), tmp3328, V4964)

			tmp3330 := PrimTail(V4963)

			tmp3331 := PrimTail(tmp3330)

			tmp3332 := PrimHead(tmp3331)

			tmp3333 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables_1h), tmp3332, V4964, V4965)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables_1h), tmp3326, tmp3329, tmp3333)
			return

		} else {
			tmp3322 := PrimIsPair(V4963)

			var ifres3303 Obj

			if True == tmp3322 {
				tmp3320 := PrimHead(V4963)

				tmp3321 := PrimEqual(symlambda, tmp3320)

				var ifres3305 Obj

				if True == tmp3321 {
					tmp3318 := PrimTail(V4963)

					tmp3319 := PrimIsPair(tmp3318)

					var ifres3307 Obj

					if True == tmp3319 {
						tmp3315 := PrimTail(V4963)

						tmp3316 := PrimTail(tmp3315)

						tmp3317 := PrimIsPair(tmp3316)

						var ifres3309 Obj

						if True == tmp3317 {
							tmp3311 := PrimTail(V4963)

							tmp3312 := PrimTail(tmp3311)

							tmp3313 := PrimTail(tmp3312)

							tmp3314 := PrimEqual(Nil, tmp3313)

							var ifres3310 Obj

							if True == tmp3314 {
								ifres3310 = True

							} else {
								ifres3310 = False

							}

							ifres3309 = ifres3310

						} else {
							ifres3309 = False

						}

						var ifres3308 Obj

						if True == ifres3309 {
							ifres3308 = True

						} else {
							ifres3308 = False

						}

						ifres3307 = ifres3308

					} else {
						ifres3307 = False

					}

					var ifres3306 Obj

					if True == ifres3307 {
						ifres3306 = True

					} else {
						ifres3306 = False

					}

					ifres3305 = ifres3306

				} else {
					ifres3305 = False

				}

				var ifres3304 Obj

				if True == ifres3305 {
					ifres3304 = True

				} else {
					ifres3304 = False

				}

				ifres3303 = ifres3304

			} else {
				ifres3303 = False

			}

			if True == ifres3303 {
				tmp3297 := PrimTail(V4963)

				tmp3298 := PrimTail(tmp3297)

				tmp3299 := PrimHead(tmp3298)

				tmp3300 := PrimTail(V4963)

				tmp3301 := PrimHead(tmp3300)

				tmp3302 := Call(__e, PrimNS2Value(symremove), tmp3301, V4964)

				__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables_1h), tmp3299, tmp3302, V4965)
				return

			} else {
				tmp3296 := PrimIsPair(V4963)

				if True == tmp3296 {
					tmp3293 := PrimTail(V4963)

					tmp3294 := PrimHead(V4963)

					tmp3295 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables_1h), tmp3294, V4964, V4965)

					__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables_1h), tmp3293, V4964, tmp3295)
					return

				} else {
					tmp3292 := Call(__e, PrimNS2Value(symelement_2), V4963, V4964)

					if True == tmp3292 {
						__e.TailApply(PrimNS2Value(symadjoin), V4963, V4965)
						return
					} else {
						__e.Return(V4965)
						return
					}

				}

			}

		}

	}, 3)

	tmp3361 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4free_1variables_1h, tmp3287)

	_ = tmp3361

	tmp3362 := MakeNative(func(__e *ControlFlow) {
		V4968 := __e.Get(1)
		_ = V4968
		V4969 := __e.Get(2)
		_ = V4969
		tmp3425 := PrimIsPair(V4968)

		var ifres3399 Obj

		if True == tmp3425 {
			tmp3423 := PrimHead(V4968)

			tmp3424 := PrimEqual(sym_f_flet_1label, tmp3423)

			var ifres3401 Obj

			if True == tmp3424 {
				tmp3421 := PrimTail(V4968)

				tmp3422 := PrimIsPair(tmp3421)

				var ifres3403 Obj

				if True == tmp3422 {
					tmp3418 := PrimTail(V4968)

					tmp3419 := PrimTail(tmp3418)

					tmp3420 := PrimIsPair(tmp3419)

					var ifres3405 Obj

					if True == tmp3420 {
						tmp3414 := PrimTail(V4968)

						tmp3415 := PrimTail(tmp3414)

						tmp3416 := PrimTail(tmp3415)

						tmp3417 := PrimIsPair(tmp3416)

						var ifres3407 Obj

						if True == tmp3417 {
							tmp3409 := PrimTail(V4968)

							tmp3410 := PrimTail(tmp3409)

							tmp3411 := PrimTail(tmp3410)

							tmp3412 := PrimTail(tmp3411)

							tmp3413 := PrimEqual(Nil, tmp3412)

							var ifres3408 Obj

							if True == tmp3413 {
								ifres3408 = True

							} else {
								ifres3408 = False

							}

							ifres3407 = ifres3408

						} else {
							ifres3407 = False

						}

						var ifres3406 Obj

						if True == ifres3407 {
							ifres3406 = True

						} else {
							ifres3406 = False

						}

						ifres3405 = ifres3406

					} else {
						ifres3405 = False

					}

					var ifres3404 Obj

					if True == ifres3405 {
						ifres3404 = True

					} else {
						ifres3404 = False

					}

					ifres3403 = ifres3404

				} else {
					ifres3403 = False

				}

				var ifres3402 Obj

				if True == ifres3403 {
					ifres3402 = True

				} else {
					ifres3402 = False

				}

				ifres3401 = ifres3402

			} else {
				ifres3401 = False

			}

			var ifres3400 Obj

			if True == ifres3401 {
				ifres3400 = True

			} else {
				ifres3400 = False

			}

			ifres3399 = ifres3400

		} else {
			ifres3399 = False

		}

		if True == ifres3399 {
			tmp3364 := MakeNative(func(__e *ControlFlow) {
				FreeVars := __e.Get(1)
				_ = FreeVars
				tmp3365 := MakeNative(func(__e *ControlFlow) {
					NewBody := __e.Get(1)
					_ = NewBody
					tmp3366 := PrimTail(V4968)

					tmp3367 := PrimHead(tmp3366)

					tmp3368 := PrimCons(tmp3367, FreeVars)

					tmp3369 := PrimTail(V4968)

					tmp3370 := PrimTail(tmp3369)

					tmp3371 := PrimHead(tmp3370)

					tmp3372 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), NewBody, V4969)

					tmp3373 := PrimCons(tmp3372, Nil)

					tmp3374 := PrimCons(tmp3371, tmp3373)

					tmp3375 := PrimCons(tmp3368, tmp3374)

					__e.Return(PrimCons(sym_f_flet_1label, tmp3375))
					return

				}, 1)

				tmp3394 := PrimEqual(Nil, FreeVars)

				var ifres3376 Obj

				if True == tmp3394 {
					tmp3390 := PrimTail(V4968)

					tmp3391 := PrimTail(tmp3390)

					tmp3392 := PrimTail(tmp3391)

					tmp3393 := PrimHead(tmp3392)

					ifres3376 = tmp3393

				} else {
					tmp3377 := PrimTail(V4968)

					tmp3378 := PrimHead(tmp3377)

					tmp3379 := PrimCons(tmp3378, FreeVars)

					tmp3380 := PrimCons(sym_f_fgoto_1label, tmp3379)

					tmp3381 := PrimTail(V4968)

					tmp3382 := PrimHead(tmp3381)

					tmp3383 := PrimCons(tmp3382, Nil)

					tmp3384 := PrimCons(sym_f_fgoto_1label, tmp3383)

					tmp3385 := PrimTail(V4968)

					tmp3386 := PrimTail(tmp3385)

					tmp3387 := PrimTail(tmp3386)

					tmp3388 := PrimHead(tmp3387)

					tmp3389 := Call(__e, PrimNS2Value(symsubst), tmp3380, tmp3384, tmp3388)

					ifres3376 = tmp3389

				}

				__e.TailApply(tmp3365, ifres3376)
				return

			}, 1)

			tmp3395 := PrimTail(V4968)

			tmp3396 := PrimTail(tmp3395)

			tmp3397 := PrimHead(tmp3396)

			tmp3398 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4free_1variables), tmp3397, V4969)

			__e.TailApply(tmp3364, tmp3398)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4factorise_1defun_4attach_1free_1variables)
			return
		}

	}, 2)

	tmp3426 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4attach_1free_1variables, tmp3362)

	_ = tmp3426

	tmp3427 := MakeNative(func(__e *ControlFlow) {
		V4976 := __e.Get(1)
		_ = V4976
		V4977 := __e.Get(2)
		_ = V4977
		tmp3606 := PrimIsPair(V4976)

		var ifres3568 Obj

		if True == tmp3606 {
			tmp3604 := PrimHead(V4976)

			tmp3605 := PrimEqual(sym_f_flet_1label, tmp3604)

			var ifres3570 Obj

			if True == tmp3605 {
				tmp3602 := PrimTail(V4976)

				tmp3603 := PrimIsPair(tmp3602)

				var ifres3572 Obj

				if True == tmp3603 {
					tmp3599 := PrimTail(V4976)

					tmp3600 := PrimTail(tmp3599)

					tmp3601 := PrimIsPair(tmp3600)

					var ifres3574 Obj

					if True == tmp3601 {
						tmp3595 := PrimTail(V4976)

						tmp3596 := PrimTail(tmp3595)

						tmp3597 := PrimTail(tmp3596)

						tmp3598 := PrimIsPair(tmp3597)

						var ifres3576 Obj

						if True == tmp3598 {
							tmp3590 := PrimTail(V4976)

							tmp3591 := PrimTail(tmp3590)

							tmp3592 := PrimTail(tmp3591)

							tmp3593 := PrimTail(tmp3592)

							tmp3594 := PrimEqual(Nil, tmp3593)

							var ifres3578 Obj

							if True == tmp3594 {
								tmp3580 := PrimTail(V4976)

								tmp3581 := PrimHead(tmp3580)

								tmp3582 := PrimCons(tmp3581, Nil)

								tmp3583 := PrimCons(sym_f_fgoto_1label, tmp3582)

								tmp3584 := PrimTail(V4976)

								tmp3585 := PrimTail(tmp3584)

								tmp3586 := PrimTail(tmp3585)

								tmp3587 := PrimHead(tmp3586)

								tmp3588 := Call(__e, PrimNS2Value(symoccurrences), tmp3583, tmp3587)

								tmp3589 := PrimGreatThan(tmp3588, MakeNumber(1))

								var ifres3579 Obj

								if True == tmp3589 {
									ifres3579 = True

								} else {
									ifres3579 = False

								}

								ifres3578 = ifres3579

							} else {
								ifres3578 = False

							}

							var ifres3577 Obj

							if True == ifres3578 {
								ifres3577 = True

							} else {
								ifres3577 = False

							}

							ifres3576 = ifres3577

						} else {
							ifres3576 = False

						}

						var ifres3575 Obj

						if True == ifres3576 {
							ifres3575 = True

						} else {
							ifres3575 = False

						}

						ifres3574 = ifres3575

					} else {
						ifres3574 = False

					}

					var ifres3573 Obj

					if True == ifres3574 {
						ifres3573 = True

					} else {
						ifres3573 = False

					}

					ifres3572 = ifres3573

				} else {
					ifres3572 = False

				}

				var ifres3571 Obj

				if True == ifres3572 {
					ifres3571 = True

				} else {
					ifres3571 = False

				}

				ifres3570 = ifres3571

			} else {
				ifres3570 = False

			}

			var ifres3569 Obj

			if True == ifres3570 {
				ifres3569 = True

			} else {
				ifres3569 = False

			}

			ifres3568 = ifres3569

		} else {
			ifres3568 = False

		}

		if True == ifres3568 {
			tmp3556 := PrimTail(V4976)

			tmp3557 := PrimHead(tmp3556)

			tmp3558 := PrimTail(V4976)

			tmp3559 := PrimTail(tmp3558)

			tmp3560 := PrimHead(tmp3559)

			tmp3561 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3560, V4977)

			tmp3562 := PrimTail(V4976)

			tmp3563 := PrimTail(tmp3562)

			tmp3564 := PrimTail(tmp3563)

			tmp3565 := PrimCons(tmp3561, tmp3564)

			tmp3566 := PrimCons(tmp3557, tmp3565)

			tmp3567 := PrimCons(sym_f_flet_1label, tmp3566)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4attach_1free_1variables), tmp3567, V4977)
			return

		} else {
			tmp3555 := PrimIsPair(V4976)

			var ifres3529 Obj

			if True == tmp3555 {
				tmp3553 := PrimHead(V4976)

				tmp3554 := PrimEqual(sym_f_flet_1label, tmp3553)

				var ifres3531 Obj

				if True == tmp3554 {
					tmp3551 := PrimTail(V4976)

					tmp3552 := PrimIsPair(tmp3551)

					var ifres3533 Obj

					if True == tmp3552 {
						tmp3548 := PrimTail(V4976)

						tmp3549 := PrimTail(tmp3548)

						tmp3550 := PrimIsPair(tmp3549)

						var ifres3535 Obj

						if True == tmp3550 {
							tmp3544 := PrimTail(V4976)

							tmp3545 := PrimTail(tmp3544)

							tmp3546 := PrimTail(tmp3545)

							tmp3547 := PrimIsPair(tmp3546)

							var ifres3537 Obj

							if True == tmp3547 {
								tmp3539 := PrimTail(V4976)

								tmp3540 := PrimTail(tmp3539)

								tmp3541 := PrimTail(tmp3540)

								tmp3542 := PrimTail(tmp3541)

								tmp3543 := PrimEqual(Nil, tmp3542)

								var ifres3538 Obj

								if True == tmp3543 {
									ifres3538 = True

								} else {
									ifres3538 = False

								}

								ifres3537 = ifres3538

							} else {
								ifres3537 = False

							}

							var ifres3536 Obj

							if True == ifres3537 {
								ifres3536 = True

							} else {
								ifres3536 = False

							}

							ifres3535 = ifres3536

						} else {
							ifres3535 = False

						}

						var ifres3534 Obj

						if True == ifres3535 {
							ifres3534 = True

						} else {
							ifres3534 = False

						}

						ifres3533 = ifres3534

					} else {
						ifres3533 = False

					}

					var ifres3532 Obj

					if True == ifres3533 {
						ifres3532 = True

					} else {
						ifres3532 = False

					}

					ifres3531 = ifres3532

				} else {
					ifres3531 = False

				}

				var ifres3530 Obj

				if True == ifres3531 {
					ifres3530 = True

				} else {
					ifres3530 = False

				}

				ifres3529 = ifres3530

			} else {
				ifres3529 = False

			}

			if True == ifres3529 {
				tmp3516 := PrimTail(V4976)

				tmp3517 := PrimTail(tmp3516)

				tmp3518 := PrimHead(tmp3517)

				tmp3519 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3518, V4977)

				tmp3520 := PrimTail(V4976)

				tmp3521 := PrimHead(tmp3520)

				tmp3522 := PrimCons(tmp3521, Nil)

				tmp3523 := PrimCons(sym_f_fgoto_1label, tmp3522)

				tmp3524 := PrimTail(V4976)

				tmp3525 := PrimTail(tmp3524)

				tmp3526 := PrimTail(tmp3525)

				tmp3527 := PrimHead(tmp3526)

				tmp3528 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3527, V4977)

				__e.TailApply(PrimNS2Value(symsubst), tmp3519, tmp3523, tmp3528)
				return

			} else {
				tmp3515 := PrimIsPair(V4976)

				var ifres3489 Obj

				if True == tmp3515 {
					tmp3513 := PrimHead(V4976)

					tmp3514 := PrimEqual(symif, tmp3513)

					var ifres3491 Obj

					if True == tmp3514 {
						tmp3511 := PrimTail(V4976)

						tmp3512 := PrimIsPair(tmp3511)

						var ifres3493 Obj

						if True == tmp3512 {
							tmp3508 := PrimTail(V4976)

							tmp3509 := PrimTail(tmp3508)

							tmp3510 := PrimIsPair(tmp3509)

							var ifres3495 Obj

							if True == tmp3510 {
								tmp3504 := PrimTail(V4976)

								tmp3505 := PrimTail(tmp3504)

								tmp3506 := PrimTail(tmp3505)

								tmp3507 := PrimIsPair(tmp3506)

								var ifres3497 Obj

								if True == tmp3507 {
									tmp3499 := PrimTail(V4976)

									tmp3500 := PrimTail(tmp3499)

									tmp3501 := PrimTail(tmp3500)

									tmp3502 := PrimTail(tmp3501)

									tmp3503 := PrimEqual(Nil, tmp3502)

									var ifres3498 Obj

									if True == tmp3503 {
										ifres3498 = True

									} else {
										ifres3498 = False

									}

									ifres3497 = ifres3498

								} else {
									ifres3497 = False

								}

								var ifres3496 Obj

								if True == ifres3497 {
									ifres3496 = True

								} else {
									ifres3496 = False

								}

								ifres3495 = ifres3496

							} else {
								ifres3495 = False

							}

							var ifres3494 Obj

							if True == ifres3495 {
								ifres3494 = True

							} else {
								ifres3494 = False

							}

							ifres3493 = ifres3494

						} else {
							ifres3493 = False

						}

						var ifres3492 Obj

						if True == ifres3493 {
							ifres3492 = True

						} else {
							ifres3492 = False

						}

						ifres3491 = ifres3492

					} else {
						ifres3491 = False

					}

					var ifres3490 Obj

					if True == ifres3491 {
						ifres3490 = True

					} else {
						ifres3490 = False

					}

					ifres3489 = ifres3490

				} else {
					ifres3489 = False

				}

				if True == ifres3489 {
					tmp3475 := PrimTail(V4976)

					tmp3476 := PrimHead(tmp3475)

					tmp3477 := PrimTail(V4976)

					tmp3478 := PrimTail(tmp3477)

					tmp3479 := PrimHead(tmp3478)

					tmp3480 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3479, V4977)

					tmp3481 := PrimTail(V4976)

					tmp3482 := PrimTail(tmp3481)

					tmp3483 := PrimTail(tmp3482)

					tmp3484 := PrimHead(tmp3483)

					tmp3485 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3484, V4977)

					tmp3486 := PrimCons(tmp3485, Nil)

					tmp3487 := PrimCons(tmp3480, tmp3486)

					tmp3488 := PrimCons(tmp3476, tmp3487)

					__e.Return(PrimCons(symif, tmp3488))
					return

				} else {
					tmp3474 := PrimIsPair(V4976)

					var ifres3448 Obj

					if True == tmp3474 {
						tmp3472 := PrimHead(V4976)

						tmp3473 := PrimEqual(symlet, tmp3472)

						var ifres3450 Obj

						if True == tmp3473 {
							tmp3470 := PrimTail(V4976)

							tmp3471 := PrimIsPair(tmp3470)

							var ifres3452 Obj

							if True == tmp3471 {
								tmp3467 := PrimTail(V4976)

								tmp3468 := PrimTail(tmp3467)

								tmp3469 := PrimIsPair(tmp3468)

								var ifres3454 Obj

								if True == tmp3469 {
									tmp3463 := PrimTail(V4976)

									tmp3464 := PrimTail(tmp3463)

									tmp3465 := PrimTail(tmp3464)

									tmp3466 := PrimIsPair(tmp3465)

									var ifres3456 Obj

									if True == tmp3466 {
										tmp3458 := PrimTail(V4976)

										tmp3459 := PrimTail(tmp3458)

										tmp3460 := PrimTail(tmp3459)

										tmp3461 := PrimTail(tmp3460)

										tmp3462 := PrimEqual(Nil, tmp3461)

										var ifres3457 Obj

										if True == tmp3462 {
											ifres3457 = True

										} else {
											ifres3457 = False

										}

										ifres3456 = ifres3457

									} else {
										ifres3456 = False

									}

									var ifres3455 Obj

									if True == ifres3456 {
										ifres3455 = True

									} else {
										ifres3455 = False

									}

									ifres3454 = ifres3455

								} else {
									ifres3454 = False

								}

								var ifres3453 Obj

								if True == ifres3454 {
									ifres3453 = True

								} else {
									ifres3453 = False

								}

								ifres3452 = ifres3453

							} else {
								ifres3452 = False

							}

							var ifres3451 Obj

							if True == ifres3452 {
								ifres3451 = True

							} else {
								ifres3451 = False

							}

							ifres3450 = ifres3451

						} else {
							ifres3450 = False

						}

						var ifres3449 Obj

						if True == ifres3450 {
							ifres3449 = True

						} else {
							ifres3449 = False

						}

						ifres3448 = ifres3449

					} else {
						ifres3448 = False

					}

					if True == ifres3448 {
						tmp3432 := PrimTail(V4976)

						tmp3433 := PrimHead(tmp3432)

						tmp3434 := PrimTail(V4976)

						tmp3435 := PrimTail(tmp3434)

						tmp3436 := PrimHead(tmp3435)

						tmp3437 := PrimTail(V4976)

						tmp3438 := PrimTail(tmp3437)

						tmp3439 := PrimTail(tmp3438)

						tmp3440 := PrimHead(tmp3439)

						tmp3441 := PrimTail(V4976)

						tmp3442 := PrimHead(tmp3441)

						tmp3443 := PrimCons(tmp3442, V4977)

						tmp3444 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4inline_1mono_1labels), tmp3440, tmp3443)

						tmp3445 := PrimCons(tmp3444, Nil)

						tmp3446 := PrimCons(tmp3436, tmp3445)

						tmp3447 := PrimCons(tmp3433, tmp3446)

						__e.Return(PrimCons(symlet, tmp3447))
						return

					} else {
						__e.Return(V4976)
						return
					}

				}

			}

		}

	}, 2)

	tmp3607 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4inline_1mono_1labels, tmp3427)

	_ = tmp3607

	tmp3608 := MakeNative(func(__e *ControlFlow) {
		V4984 := __e.Get(1)
		_ = V4984
		V4985 := __e.Get(2)
		_ = V4985
		tmp3726 := PrimEqual(Nil, V4984)

		if True == tmp3726 {
			__e.Return(V4985)
			return
		} else {
			tmp3725 := PrimIsPair(V4984)

			var ifres3705 Obj

			if True == tmp3725 {
				tmp3723 := PrimHead(V4984)

				tmp3724 := PrimIsPair(tmp3723)

				var ifres3707 Obj

				if True == tmp3724 {
					tmp3720 := PrimHead(V4984)

					tmp3721 := PrimHead(tmp3720)

					tmp3722 := PrimEqual(True, tmp3721)

					var ifres3709 Obj

					if True == tmp3722 {
						tmp3717 := PrimHead(V4984)

						tmp3718 := PrimTail(tmp3717)

						tmp3719 := PrimIsPair(tmp3718)

						var ifres3711 Obj

						if True == tmp3719 {
							tmp3713 := PrimHead(V4984)

							tmp3714 := PrimTail(tmp3713)

							tmp3715 := PrimTail(tmp3714)

							tmp3716 := PrimEqual(Nil, tmp3715)

							var ifres3712 Obj

							if True == tmp3716 {
								ifres3712 = True

							} else {
								ifres3712 = False

							}

							ifres3711 = ifres3712

						} else {
							ifres3711 = False

						}

						var ifres3710 Obj

						if True == ifres3711 {
							ifres3710 = True

						} else {
							ifres3710 = False

						}

						ifres3709 = ifres3710

					} else {
						ifres3709 = False

					}

					var ifres3708 Obj

					if True == ifres3709 {
						ifres3708 = True

					} else {
						ifres3708 = False

					}

					ifres3707 = ifres3708

				} else {
					ifres3707 = False

				}

				var ifres3706 Obj

				if True == ifres3707 {
					ifres3706 = True

				} else {
					ifres3706 = False

				}

				ifres3705 = ifres3706

			} else {
				ifres3705 = False

			}

			if True == ifres3705 {
				tmp3703 := PrimHead(V4984)

				tmp3704 := PrimTail(tmp3703)

				__e.Return(PrimHead(tmp3704))
				return

			} else {
				tmp3702 := PrimIsPair(V4984)

				var ifres3655 Obj

				if True == tmp3702 {
					tmp3700 := PrimHead(V4984)

					tmp3701 := PrimIsPair(tmp3700)

					var ifres3657 Obj

					if True == tmp3701 {
						tmp3697 := PrimHead(V4984)

						tmp3698 := PrimHead(tmp3697)

						tmp3699 := PrimIsPair(tmp3698)

						var ifres3659 Obj

						if True == tmp3699 {
							tmp3693 := PrimHead(V4984)

							tmp3694 := PrimHead(tmp3693)

							tmp3695 := PrimHead(tmp3694)

							tmp3696 := PrimEqual(symand, tmp3695)

							var ifres3661 Obj

							if True == tmp3696 {
								tmp3689 := PrimHead(V4984)

								tmp3690 := PrimHead(tmp3689)

								tmp3691 := PrimTail(tmp3690)

								tmp3692 := PrimIsPair(tmp3691)

								var ifres3663 Obj

								if True == tmp3692 {
									tmp3684 := PrimHead(V4984)

									tmp3685 := PrimHead(tmp3684)

									tmp3686 := PrimTail(tmp3685)

									tmp3687 := PrimTail(tmp3686)

									tmp3688 := PrimIsPair(tmp3687)

									var ifres3665 Obj

									if True == tmp3688 {
										tmp3678 := PrimHead(V4984)

										tmp3679 := PrimHead(tmp3678)

										tmp3680 := PrimTail(tmp3679)

										tmp3681 := PrimTail(tmp3680)

										tmp3682 := PrimTail(tmp3681)

										tmp3683 := PrimEqual(Nil, tmp3682)

										var ifres3667 Obj

										if True == tmp3683 {
											tmp3675 := PrimHead(V4984)

											tmp3676 := PrimTail(tmp3675)

											tmp3677 := PrimIsPair(tmp3676)

											var ifres3669 Obj

											if True == tmp3677 {
												tmp3671 := PrimHead(V4984)

												tmp3672 := PrimTail(tmp3671)

												tmp3673 := PrimTail(tmp3672)

												tmp3674 := PrimEqual(Nil, tmp3673)

												var ifres3670 Obj

												if True == tmp3674 {
													ifres3670 = True

												} else {
													ifres3670 = False

												}

												ifres3669 = ifres3670

											} else {
												ifres3669 = False

											}

											var ifres3668 Obj

											if True == ifres3669 {
												ifres3668 = True

											} else {
												ifres3668 = False

											}

											ifres3667 = ifres3668

										} else {
											ifres3667 = False

										}

										var ifres3666 Obj

										if True == ifres3667 {
											ifres3666 = True

										} else {
											ifres3666 = False

										}

										ifres3665 = ifres3666

									} else {
										ifres3665 = False

									}

									var ifres3664 Obj

									if True == ifres3665 {
										ifres3664 = True

									} else {
										ifres3664 = False

									}

									ifres3663 = ifres3664

								} else {
									ifres3663 = False

								}

								var ifres3662 Obj

								if True == ifres3663 {
									ifres3662 = True

								} else {
									ifres3662 = False

								}

								ifres3661 = ifres3662

							} else {
								ifres3661 = False

							}

							var ifres3660 Obj

							if True == ifres3661 {
								ifres3660 = True

							} else {
								ifres3660 = False

							}

							ifres3659 = ifres3660

						} else {
							ifres3659 = False

						}

						var ifres3658 Obj

						if True == ifres3659 {
							ifres3658 = True

						} else {
							ifres3658 = False

						}

						ifres3657 = ifres3658

					} else {
						ifres3657 = False

					}

					var ifres3656 Obj

					if True == ifres3657 {
						ifres3656 = True

					} else {
						ifres3656 = False

					}

					ifres3655 = ifres3656

				} else {
					ifres3655 = False

				}

				if True == ifres3655 {
					tmp3639 := MakeNative(func(__e *ControlFlow) {
						TrueBranch := __e.Get(1)
						_ = TrueBranch
						tmp3640 := MakeNative(func(__e *ControlFlow) {
							FalseBranch := __e.Get(1)
							_ = FalseBranch
							tmp3641 := PrimHead(V4984)

							tmp3642 := PrimHead(tmp3641)

							tmp3643 := PrimTail(tmp3642)

							tmp3644 := PrimHead(tmp3643)

							__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4rebranch_1h), tmp3644, TrueBranch, FalseBranch, V4985)
							return

						}, 1)

						tmp3645 := PrimHead(V4984)

						tmp3646 := PrimHead(tmp3645)

						tmp3647 := PrimTail(tmp3646)

						tmp3648 := PrimHead(tmp3647)

						tmp3649 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4false_1branch), tmp3648, V4984)

						__e.TailApply(tmp3640, tmp3649)
						return

					}, 1)

					tmp3650 := PrimHead(V4984)

					tmp3651 := PrimHead(tmp3650)

					tmp3652 := PrimTail(tmp3651)

					tmp3653 := PrimHead(tmp3652)

					tmp3654 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4true_1branch), tmp3653, V4984)

					__e.TailApply(tmp3639, tmp3654)
					return

				} else {
					tmp3638 := PrimIsPair(V4984)

					var ifres3623 Obj

					if True == tmp3638 {
						tmp3636 := PrimHead(V4984)

						tmp3637 := PrimIsPair(tmp3636)

						var ifres3625 Obj

						if True == tmp3637 {
							tmp3633 := PrimHead(V4984)

							tmp3634 := PrimTail(tmp3633)

							tmp3635 := PrimIsPair(tmp3634)

							var ifres3627 Obj

							if True == tmp3635 {
								tmp3629 := PrimHead(V4984)

								tmp3630 := PrimTail(tmp3629)

								tmp3631 := PrimTail(tmp3630)

								tmp3632 := PrimEqual(Nil, tmp3631)

								var ifres3628 Obj

								if True == tmp3632 {
									ifres3628 = True

								} else {
									ifres3628 = False

								}

								ifres3627 = ifres3628

							} else {
								ifres3627 = False

							}

							var ifres3626 Obj

							if True == ifres3627 {
								ifres3626 = True

							} else {
								ifres3626 = False

							}

							ifres3625 = ifres3626

						} else {
							ifres3625 = False

						}

						var ifres3624 Obj

						if True == ifres3625 {
							ifres3624 = True

						} else {
							ifres3624 = False

						}

						ifres3623 = ifres3624

					} else {
						ifres3623 = False

					}

					if True == ifres3623 {
						tmp3613 := MakeNative(func(__e *ControlFlow) {
							TrueBranch := __e.Get(1)
							_ = TrueBranch
							tmp3614 := MakeNative(func(__e *ControlFlow) {
								FalseBranch := __e.Get(1)
								_ = FalseBranch
								tmp3615 := PrimHead(V4984)

								tmp3616 := PrimHead(tmp3615)

								__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4rebranch_1h), tmp3616, TrueBranch, FalseBranch, V4985)
								return

							}, 1)

							tmp3617 := PrimHead(V4984)

							tmp3618 := PrimHead(tmp3617)

							tmp3619 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4false_1branch), tmp3618, V4984)

							__e.TailApply(tmp3614, tmp3619)
							return

						}, 1)

						tmp3620 := PrimHead(V4984)

						tmp3621 := PrimHead(tmp3620)

						tmp3622 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4true_1branch), tmp3621, V4984)

						__e.TailApply(tmp3613, tmp3622)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4factorise_1defun_4rebranch)
						return
					}

				}

			}

		}

	}, 2)

	tmp3727 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4rebranch, tmp3608)

	_ = tmp3727

	tmp3728 := MakeNative(func(__e *ControlFlow) {
		V4990 := __e.Get(1)
		_ = V4990
		V4991 := __e.Get(2)
		_ = V4991
		V4992 := __e.Get(3)
		_ = V4992
		V4993 := __e.Get(4)
		_ = V4993
		tmp3729 := MakeNative(func(__e *ControlFlow) {
			NewElse := __e.Get(1)
			_ = NewElse
			tmp3730 := MakeNative(func(__e *ControlFlow) {
				GotoElse := __e.Get(1)
				_ = GotoElse
				tmp3731 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4rebranch), V4991, GotoElse)

				tmp3732 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4optimize_1selectors), V4990, tmp3731)

				tmp3733 := PrimCons(GotoElse, Nil)

				tmp3734 := PrimCons(tmp3732, tmp3733)

				tmp3735 := PrimCons(V4990, tmp3734)

				tmp3736 := PrimCons(symif, tmp3735)

				__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs), tmp3736)
				return

			}, 1)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4with_1labelled_1else), NewElse, tmp3730)
			return

		}, 1)

		tmp3737 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4rebranch), V4992, V4993)

		__e.TailApply(tmp3729, tmp3737)
		return

	}, 4)

	tmp3738 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4rebranch_1h, tmp3728)

	_ = tmp3738

	tmp3739 := MakeNative(func(__e *ControlFlow) {
		V5006 := __e.Get(1)
		_ = V5006
		V5007 := __e.Get(2)
		_ = V5007
		tmp3834 := PrimIsPair(V5007)

		var ifres3780 Obj

		if True == tmp3834 {
			tmp3832 := PrimHead(V5007)

			tmp3833 := PrimIsPair(tmp3832)

			var ifres3782 Obj

			if True == tmp3833 {
				tmp3829 := PrimHead(V5007)

				tmp3830 := PrimHead(tmp3829)

				tmp3831 := PrimIsPair(tmp3830)

				var ifres3784 Obj

				if True == tmp3831 {
					tmp3825 := PrimHead(V5007)

					tmp3826 := PrimHead(tmp3825)

					tmp3827 := PrimHead(tmp3826)

					tmp3828 := PrimEqual(symand, tmp3827)

					var ifres3786 Obj

					if True == tmp3828 {
						tmp3821 := PrimHead(V5007)

						tmp3822 := PrimHead(tmp3821)

						tmp3823 := PrimTail(tmp3822)

						tmp3824 := PrimIsPair(tmp3823)

						var ifres3788 Obj

						if True == tmp3824 {
							tmp3816 := PrimHead(V5007)

							tmp3817 := PrimHead(tmp3816)

							tmp3818 := PrimTail(tmp3817)

							tmp3819 := PrimTail(tmp3818)

							tmp3820 := PrimIsPair(tmp3819)

							var ifres3790 Obj

							if True == tmp3820 {
								tmp3810 := PrimHead(V5007)

								tmp3811 := PrimHead(tmp3810)

								tmp3812 := PrimTail(tmp3811)

								tmp3813 := PrimTail(tmp3812)

								tmp3814 := PrimTail(tmp3813)

								tmp3815 := PrimEqual(Nil, tmp3814)

								var ifres3792 Obj

								if True == tmp3815 {
									tmp3807 := PrimHead(V5007)

									tmp3808 := PrimTail(tmp3807)

									tmp3809 := PrimIsPair(tmp3808)

									var ifres3794 Obj

									if True == tmp3809 {
										tmp3803 := PrimHead(V5007)

										tmp3804 := PrimTail(tmp3803)

										tmp3805 := PrimTail(tmp3804)

										tmp3806 := PrimEqual(Nil, tmp3805)

										var ifres3796 Obj

										if True == tmp3806 {
											tmp3798 := PrimHead(V5007)

											tmp3799 := PrimHead(tmp3798)

											tmp3800 := PrimTail(tmp3799)

											tmp3801 := PrimHead(tmp3800)

											tmp3802 := PrimEqual(tmp3801, V5006)

											var ifres3797 Obj

											if True == tmp3802 {
												ifres3797 = True

											} else {
												ifres3797 = False

											}

											ifres3796 = ifres3797

										} else {
											ifres3796 = False

										}

										var ifres3795 Obj

										if True == ifres3796 {
											ifres3795 = True

										} else {
											ifres3795 = False

										}

										ifres3794 = ifres3795

									} else {
										ifres3794 = False

									}

									var ifres3793 Obj

									if True == ifres3794 {
										ifres3793 = True

									} else {
										ifres3793 = False

									}

									ifres3792 = ifres3793

								} else {
									ifres3792 = False

								}

								var ifres3791 Obj

								if True == ifres3792 {
									ifres3791 = True

								} else {
									ifres3791 = False

								}

								ifres3790 = ifres3791

							} else {
								ifres3790 = False

							}

							var ifres3789 Obj

							if True == ifres3790 {
								ifres3789 = True

							} else {
								ifres3789 = False

							}

							ifres3788 = ifres3789

						} else {
							ifres3788 = False

						}

						var ifres3787 Obj

						if True == ifres3788 {
							ifres3787 = True

						} else {
							ifres3787 = False

						}

						ifres3786 = ifres3787

					} else {
						ifres3786 = False

					}

					var ifres3785 Obj

					if True == ifres3786 {
						ifres3785 = True

					} else {
						ifres3785 = False

					}

					ifres3784 = ifres3785

				} else {
					ifres3784 = False

				}

				var ifres3783 Obj

				if True == ifres3784 {
					ifres3783 = True

				} else {
					ifres3783 = False

				}

				ifres3782 = ifres3783

			} else {
				ifres3782 = False

			}

			var ifres3781 Obj

			if True == ifres3782 {
				ifres3781 = True

			} else {
				ifres3781 = False

			}

			ifres3780 = ifres3781

		} else {
			ifres3780 = False

		}

		if True == ifres3780 {
			tmp3766 := PrimHead(V5007)

			tmp3767 := PrimHead(tmp3766)

			tmp3768 := PrimTail(tmp3767)

			tmp3769 := PrimTail(tmp3768)

			tmp3770 := PrimHead(tmp3769)

			tmp3771 := PrimHead(V5007)

			tmp3772 := PrimTail(tmp3771)

			tmp3773 := PrimCons(tmp3770, tmp3772)

			tmp3774 := PrimHead(V5007)

			tmp3775 := PrimHead(tmp3774)

			tmp3776 := PrimTail(tmp3775)

			tmp3777 := PrimHead(tmp3776)

			tmp3778 := PrimTail(V5007)

			tmp3779 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4true_1branch), tmp3777, tmp3778)

			__e.Return(PrimCons(tmp3773, tmp3779))
			return

		} else {
			tmp3765 := PrimIsPair(V5007)

			var ifres3745 Obj

			if True == tmp3765 {
				tmp3763 := PrimHead(V5007)

				tmp3764 := PrimIsPair(tmp3763)

				var ifres3747 Obj

				if True == tmp3764 {
					tmp3760 := PrimHead(V5007)

					tmp3761 := PrimTail(tmp3760)

					tmp3762 := PrimIsPair(tmp3761)

					var ifres3749 Obj

					if True == tmp3762 {
						tmp3756 := PrimHead(V5007)

						tmp3757 := PrimTail(tmp3756)

						tmp3758 := PrimTail(tmp3757)

						tmp3759 := PrimEqual(Nil, tmp3758)

						var ifres3751 Obj

						if True == tmp3759 {
							tmp3753 := PrimHead(V5007)

							tmp3754 := PrimHead(tmp3753)

							tmp3755 := PrimEqual(tmp3754, V5006)

							var ifres3752 Obj

							if True == tmp3755 {
								ifres3752 = True

							} else {
								ifres3752 = False

							}

							ifres3751 = ifres3752

						} else {
							ifres3751 = False

						}

						var ifres3750 Obj

						if True == ifres3751 {
							ifres3750 = True

						} else {
							ifres3750 = False

						}

						ifres3749 = ifres3750

					} else {
						ifres3749 = False

					}

					var ifres3748 Obj

					if True == ifres3749 {
						ifres3748 = True

					} else {
						ifres3748 = False

					}

					ifres3747 = ifres3748

				} else {
					ifres3747 = False

				}

				var ifres3746 Obj

				if True == ifres3747 {
					ifres3746 = True

				} else {
					ifres3746 = False

				}

				ifres3745 = ifres3746

			} else {
				ifres3745 = False

			}

			if True == ifres3745 {
				tmp3742 := PrimHead(V5007)

				tmp3743 := PrimTail(tmp3742)

				tmp3744 := PrimCons(True, tmp3743)

				__e.Return(PrimCons(tmp3744, Nil))
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 2)

	tmp3835 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4true_1branch, tmp3739)

	_ = tmp3835

	tmp3836 := MakeNative(func(__e *ControlFlow) {
		V5016 := __e.Get(1)
		_ = V5016
		V5017 := __e.Get(2)
		_ = V5017
		tmp3922 := PrimIsPair(V5017)

		var ifres3868 Obj

		if True == tmp3922 {
			tmp3920 := PrimHead(V5017)

			tmp3921 := PrimIsPair(tmp3920)

			var ifres3870 Obj

			if True == tmp3921 {
				tmp3917 := PrimHead(V5017)

				tmp3918 := PrimHead(tmp3917)

				tmp3919 := PrimIsPair(tmp3918)

				var ifres3872 Obj

				if True == tmp3919 {
					tmp3913 := PrimHead(V5017)

					tmp3914 := PrimHead(tmp3913)

					tmp3915 := PrimHead(tmp3914)

					tmp3916 := PrimEqual(symand, tmp3915)

					var ifres3874 Obj

					if True == tmp3916 {
						tmp3909 := PrimHead(V5017)

						tmp3910 := PrimHead(tmp3909)

						tmp3911 := PrimTail(tmp3910)

						tmp3912 := PrimIsPair(tmp3911)

						var ifres3876 Obj

						if True == tmp3912 {
							tmp3904 := PrimHead(V5017)

							tmp3905 := PrimHead(tmp3904)

							tmp3906 := PrimTail(tmp3905)

							tmp3907 := PrimTail(tmp3906)

							tmp3908 := PrimIsPair(tmp3907)

							var ifres3878 Obj

							if True == tmp3908 {
								tmp3898 := PrimHead(V5017)

								tmp3899 := PrimHead(tmp3898)

								tmp3900 := PrimTail(tmp3899)

								tmp3901 := PrimTail(tmp3900)

								tmp3902 := PrimTail(tmp3901)

								tmp3903 := PrimEqual(Nil, tmp3902)

								var ifres3880 Obj

								if True == tmp3903 {
									tmp3895 := PrimHead(V5017)

									tmp3896 := PrimTail(tmp3895)

									tmp3897 := PrimIsPair(tmp3896)

									var ifres3882 Obj

									if True == tmp3897 {
										tmp3891 := PrimHead(V5017)

										tmp3892 := PrimTail(tmp3891)

										tmp3893 := PrimTail(tmp3892)

										tmp3894 := PrimEqual(Nil, tmp3893)

										var ifres3884 Obj

										if True == tmp3894 {
											tmp3886 := PrimHead(V5017)

											tmp3887 := PrimHead(tmp3886)

											tmp3888 := PrimTail(tmp3887)

											tmp3889 := PrimHead(tmp3888)

											tmp3890 := PrimEqual(tmp3889, V5016)

											var ifres3885 Obj

											if True == tmp3890 {
												ifres3885 = True

											} else {
												ifres3885 = False

											}

											ifres3884 = ifres3885

										} else {
											ifres3884 = False

										}

										var ifres3883 Obj

										if True == ifres3884 {
											ifres3883 = True

										} else {
											ifres3883 = False

										}

										ifres3882 = ifres3883

									} else {
										ifres3882 = False

									}

									var ifres3881 Obj

									if True == ifres3882 {
										ifres3881 = True

									} else {
										ifres3881 = False

									}

									ifres3880 = ifres3881

								} else {
									ifres3880 = False

								}

								var ifres3879 Obj

								if True == ifres3880 {
									ifres3879 = True

								} else {
									ifres3879 = False

								}

								ifres3878 = ifres3879

							} else {
								ifres3878 = False

							}

							var ifres3877 Obj

							if True == ifres3878 {
								ifres3877 = True

							} else {
								ifres3877 = False

							}

							ifres3876 = ifres3877

						} else {
							ifres3876 = False

						}

						var ifres3875 Obj

						if True == ifres3876 {
							ifres3875 = True

						} else {
							ifres3875 = False

						}

						ifres3874 = ifres3875

					} else {
						ifres3874 = False

					}

					var ifres3873 Obj

					if True == ifres3874 {
						ifres3873 = True

					} else {
						ifres3873 = False

					}

					ifres3872 = ifres3873

				} else {
					ifres3872 = False

				}

				var ifres3871 Obj

				if True == ifres3872 {
					ifres3871 = True

				} else {
					ifres3871 = False

				}

				ifres3870 = ifres3871

			} else {
				ifres3870 = False

			}

			var ifres3869 Obj

			if True == ifres3870 {
				ifres3869 = True

			} else {
				ifres3869 = False

			}

			ifres3868 = ifres3869

		} else {
			ifres3868 = False

		}

		if True == ifres3868 {
			tmp3863 := PrimHead(V5017)

			tmp3864 := PrimHead(tmp3863)

			tmp3865 := PrimTail(tmp3864)

			tmp3866 := PrimHead(tmp3865)

			tmp3867 := PrimTail(V5017)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4false_1branch), tmp3866, tmp3867)
			return

		} else {
			tmp3862 := PrimIsPair(V5017)

			var ifres3842 Obj

			if True == tmp3862 {
				tmp3860 := PrimHead(V5017)

				tmp3861 := PrimIsPair(tmp3860)

				var ifres3844 Obj

				if True == tmp3861 {
					tmp3857 := PrimHead(V5017)

					tmp3858 := PrimTail(tmp3857)

					tmp3859 := PrimIsPair(tmp3858)

					var ifres3846 Obj

					if True == tmp3859 {
						tmp3853 := PrimHead(V5017)

						tmp3854 := PrimTail(tmp3853)

						tmp3855 := PrimTail(tmp3854)

						tmp3856 := PrimEqual(Nil, tmp3855)

						var ifres3848 Obj

						if True == tmp3856 {
							tmp3850 := PrimHead(V5017)

							tmp3851 := PrimHead(tmp3850)

							tmp3852 := PrimEqual(tmp3851, V5016)

							var ifres3849 Obj

							if True == tmp3852 {
								ifres3849 = True

							} else {
								ifres3849 = False

							}

							ifres3848 = ifres3849

						} else {
							ifres3848 = False

						}

						var ifres3847 Obj

						if True == ifres3848 {
							ifres3847 = True

						} else {
							ifres3847 = False

						}

						ifres3846 = ifres3847

					} else {
						ifres3846 = False

					}

					var ifres3845 Obj

					if True == ifres3846 {
						ifres3845 = True

					} else {
						ifres3845 = False

					}

					ifres3844 = ifres3845

				} else {
					ifres3844 = False

				}

				var ifres3843 Obj

				if True == ifres3844 {
					ifres3843 = True

				} else {
					ifres3843 = False

				}

				ifres3842 = ifres3843

			} else {
				ifres3842 = False

			}

			if True == ifres3842 {
				tmp3839 := PrimHead(V5017)

				tmp3840 := PrimHead(tmp3839)

				tmp3841 := PrimTail(V5017)

				__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4false_1branch), tmp3840, tmp3841)
				return

			} else {
				__e.Return(V5017)
				return
			}

		}

	}, 2)

	tmp3923 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4false_1branch, tmp3836)

	_ = tmp3923

	tmp3924 := MakeNative(func(__e *ControlFlow) {
		V5020 := __e.Get(1)
		_ = V5020
		V5021 := __e.Get(2)
		_ = V5021
		tmp3978 := PrimIsPair(V5020)

		var ifres3959 Obj

		if True == tmp3978 {
			tmp3976 := PrimHead(V5020)

			tmp3977 := PrimEqual(sym_f_freturn, tmp3976)

			var ifres3961 Obj

			if True == tmp3977 {
				tmp3974 := PrimTail(V5020)

				tmp3975 := PrimIsPair(tmp3974)

				var ifres3963 Obj

				if True == tmp3975 {
					tmp3971 := PrimTail(V5020)

					tmp3972 := PrimTail(tmp3971)

					tmp3973 := PrimEqual(Nil, tmp3972)

					var ifres3965 Obj

					if True == tmp3973 {
						tmp3967 := PrimTail(V5020)

						tmp3968 := PrimHead(tmp3967)

						tmp3969 := PrimIsPair(tmp3968)

						tmp3970 := PrimNot(tmp3969)

						var ifres3966 Obj

						if True == tmp3970 {
							ifres3966 = True

						} else {
							ifres3966 = False

						}

						ifres3965 = ifres3966

					} else {
						ifres3965 = False

					}

					var ifres3964 Obj

					if True == ifres3965 {
						ifres3964 = True

					} else {
						ifres3964 = False

					}

					ifres3963 = ifres3964

				} else {
					ifres3963 = False

				}

				var ifres3962 Obj

				if True == ifres3963 {
					ifres3962 = True

				} else {
					ifres3962 = False

				}

				ifres3961 = ifres3962

			} else {
				ifres3961 = False

			}

			var ifres3960 Obj

			if True == ifres3961 {
				ifres3960 = True

			} else {
				ifres3960 = False

			}

			ifres3959 = ifres3960

		} else {
			ifres3959 = False

		}

		if True == ifres3959 {
			__e.TailApply(V5021, V5020)
			return
		} else {
			tmp3958 := PrimIsPair(V5020)

			var ifres3950 Obj

			if True == tmp3958 {
				tmp3956 := PrimHead(V5020)

				tmp3957 := PrimEqual(symfail, tmp3956)

				var ifres3952 Obj

				if True == tmp3957 {
					tmp3954 := PrimTail(V5020)

					tmp3955 := PrimEqual(Nil, tmp3954)

					var ifres3953 Obj

					if True == tmp3955 {
						ifres3953 = True

					} else {
						ifres3953 = False

					}

					ifres3952 = ifres3953

				} else {
					ifres3952 = False

				}

				var ifres3951 Obj

				if True == ifres3952 {
					ifres3951 = True

				} else {
					ifres3951 = False

				}

				ifres3950 = ifres3951

			} else {
				ifres3950 = False

			}

			if True == ifres3950 {
				__e.TailApply(V5021, V5020)
				return
			} else {
				tmp3949 := PrimIsPair(V5020)

				var ifres3936 Obj

				if True == tmp3949 {
					tmp3947 := PrimHead(V5020)

					tmp3948 := PrimEqual(sym_f_fgoto_1label, tmp3947)

					var ifres3938 Obj

					if True == tmp3948 {
						tmp3945 := PrimTail(V5020)

						tmp3946 := PrimIsPair(tmp3945)

						var ifres3940 Obj

						if True == tmp3946 {
							tmp3942 := PrimTail(V5020)

							tmp3943 := PrimTail(tmp3942)

							tmp3944 := PrimEqual(Nil, tmp3943)

							var ifres3941 Obj

							if True == tmp3944 {
								ifres3941 = True

							} else {
								ifres3941 = False

							}

							ifres3940 = ifres3941

						} else {
							ifres3940 = False

						}

						var ifres3939 Obj

						if True == ifres3940 {
							ifres3939 = True

						} else {
							ifres3939 = False

						}

						ifres3938 = ifres3939

					} else {
						ifres3938 = False

					}

					var ifres3937 Obj

					if True == ifres3938 {
						ifres3937 = True

					} else {
						ifres3937 = False

					}

					ifres3936 = ifres3937

				} else {
					ifres3936 = False

				}

				if True == ifres3936 {
					__e.TailApply(V5021, V5020)
					return
				} else {
					tmp3928 := MakeNative(func(__e *ControlFlow) {
						Label := __e.Get(1)
						_ = Label
						tmp3929 := PrimCons(Label, Nil)

						tmp3930 := PrimCons(sym_f_fgoto_1label, tmp3929)

						tmp3931 := Call(__e, V5021, tmp3930)

						tmp3932 := PrimCons(tmp3931, Nil)

						tmp3933 := PrimCons(V5020, tmp3932)

						tmp3934 := PrimCons(Label, tmp3933)

						__e.Return(PrimCons(sym_f_flet_1label, tmp3934))
						return

					}, 1)

					tmp3935 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4generate_1label))

					__e.TailApply(tmp3928, tmp3935)
					return

				}

			}

		}

	}, 2)

	tmp3979 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4with_1labelled_1else, tmp3924)

	_ = tmp3979

	tmp3980 := MakeNative(func(__e *ControlFlow) {
		V5024 := __e.Get(1)
		_ = V5024
		tmp4090 := PrimIsPair(V5024)

		var ifres4003 Obj

		if True == tmp4090 {
			tmp4088 := PrimHead(V5024)

			tmp4089 := PrimEqual(symif, tmp4088)

			var ifres4005 Obj

			if True == tmp4089 {
				tmp4086 := PrimTail(V5024)

				tmp4087 := PrimIsPair(tmp4086)

				var ifres4007 Obj

				if True == tmp4087 {
					tmp4083 := PrimTail(V5024)

					tmp4084 := PrimTail(tmp4083)

					tmp4085 := PrimIsPair(tmp4084)

					var ifres4009 Obj

					if True == tmp4085 {
						tmp4079 := PrimTail(V5024)

						tmp4080 := PrimTail(tmp4079)

						tmp4081 := PrimHead(tmp4080)

						tmp4082 := PrimIsPair(tmp4081)

						var ifres4011 Obj

						if True == tmp4082 {
							tmp4074 := PrimTail(V5024)

							tmp4075 := PrimTail(tmp4074)

							tmp4076 := PrimHead(tmp4075)

							tmp4077 := PrimHead(tmp4076)

							tmp4078 := PrimEqual(symif, tmp4077)

							var ifres4013 Obj

							if True == tmp4078 {
								tmp4069 := PrimTail(V5024)

								tmp4070 := PrimTail(tmp4069)

								tmp4071 := PrimHead(tmp4070)

								tmp4072 := PrimTail(tmp4071)

								tmp4073 := PrimIsPair(tmp4072)

								var ifres4015 Obj

								if True == tmp4073 {
									tmp4063 := PrimTail(V5024)

									tmp4064 := PrimTail(tmp4063)

									tmp4065 := PrimHead(tmp4064)

									tmp4066 := PrimTail(tmp4065)

									tmp4067 := PrimTail(tmp4066)

									tmp4068 := PrimIsPair(tmp4067)

									var ifres4017 Obj

									if True == tmp4068 {
										tmp4056 := PrimTail(V5024)

										tmp4057 := PrimTail(tmp4056)

										tmp4058 := PrimHead(tmp4057)

										tmp4059 := PrimTail(tmp4058)

										tmp4060 := PrimTail(tmp4059)

										tmp4061 := PrimTail(tmp4060)

										tmp4062 := PrimIsPair(tmp4061)

										var ifres4019 Obj

										if True == tmp4062 {
											tmp4048 := PrimTail(V5024)

											tmp4049 := PrimTail(tmp4048)

											tmp4050 := PrimHead(tmp4049)

											tmp4051 := PrimTail(tmp4050)

											tmp4052 := PrimTail(tmp4051)

											tmp4053 := PrimTail(tmp4052)

											tmp4054 := PrimTail(tmp4053)

											tmp4055 := PrimEqual(Nil, tmp4054)

											var ifres4021 Obj

											if True == tmp4055 {
												tmp4044 := PrimTail(V5024)

												tmp4045 := PrimTail(tmp4044)

												tmp4046 := PrimTail(tmp4045)

												tmp4047 := PrimIsPair(tmp4046)

												var ifres4023 Obj

												if True == tmp4047 {
													tmp4039 := PrimTail(V5024)

													tmp4040 := PrimTail(tmp4039)

													tmp4041 := PrimTail(tmp4040)

													tmp4042 := PrimTail(tmp4041)

													tmp4043 := PrimEqual(Nil, tmp4042)

													var ifres4025 Obj

													if True == tmp4043 {
														tmp4027 := PrimTail(V5024)

														tmp4028 := PrimTail(tmp4027)

														tmp4029 := PrimTail(tmp4028)

														tmp4030 := PrimHead(tmp4029)

														tmp4031 := PrimTail(V5024)

														tmp4032 := PrimTail(tmp4031)

														tmp4033 := PrimHead(tmp4032)

														tmp4034 := PrimTail(tmp4033)

														tmp4035 := PrimTail(tmp4034)

														tmp4036 := PrimTail(tmp4035)

														tmp4037 := PrimHead(tmp4036)

														tmp4038 := PrimEqual(tmp4030, tmp4037)

														var ifres4026 Obj

														if True == tmp4038 {
															ifres4026 = True

														} else {
															ifres4026 = False

														}

														ifres4025 = ifres4026

													} else {
														ifres4025 = False

													}

													var ifres4024 Obj

													if True == ifres4025 {
														ifres4024 = True

													} else {
														ifres4024 = False

													}

													ifres4023 = ifres4024

												} else {
													ifres4023 = False

												}

												var ifres4022 Obj

												if True == ifres4023 {
													ifres4022 = True

												} else {
													ifres4022 = False

												}

												ifres4021 = ifres4022

											} else {
												ifres4021 = False

											}

											var ifres4020 Obj

											if True == ifres4021 {
												ifres4020 = True

											} else {
												ifres4020 = False

											}

											ifres4019 = ifres4020

										} else {
											ifres4019 = False

										}

										var ifres4018 Obj

										if True == ifres4019 {
											ifres4018 = True

										} else {
											ifres4018 = False

										}

										ifres4017 = ifres4018

									} else {
										ifres4017 = False

									}

									var ifres4016 Obj

									if True == ifres4017 {
										ifres4016 = True

									} else {
										ifres4016 = False

									}

									ifres4015 = ifres4016

								} else {
									ifres4015 = False

								}

								var ifres4014 Obj

								if True == ifres4015 {
									ifres4014 = True

								} else {
									ifres4014 = False

								}

								ifres4013 = ifres4014

							} else {
								ifres4013 = False

							}

							var ifres4012 Obj

							if True == ifres4013 {
								ifres4012 = True

							} else {
								ifres4012 = False

							}

							ifres4011 = ifres4012

						} else {
							ifres4011 = False

						}

						var ifres4010 Obj

						if True == ifres4011 {
							ifres4010 = True

						} else {
							ifres4010 = False

						}

						ifres4009 = ifres4010

					} else {
						ifres4009 = False

					}

					var ifres4008 Obj

					if True == ifres4009 {
						ifres4008 = True

					} else {
						ifres4008 = False

					}

					ifres4007 = ifres4008

				} else {
					ifres4007 = False

				}

				var ifres4006 Obj

				if True == ifres4007 {
					ifres4006 = True

				} else {
					ifres4006 = False

				}

				ifres4005 = ifres4006

			} else {
				ifres4005 = False

			}

			var ifres4004 Obj

			if True == ifres4005 {
				ifres4004 = True

			} else {
				ifres4004 = False

			}

			ifres4003 = ifres4004

		} else {
			ifres4003 = False

		}

		if True == ifres4003 {
			tmp3982 := PrimTail(V5024)

			tmp3983 := PrimHead(tmp3982)

			tmp3984 := PrimTail(V5024)

			tmp3985 := PrimTail(tmp3984)

			tmp3986 := PrimHead(tmp3985)

			tmp3987 := PrimTail(tmp3986)

			tmp3988 := PrimHead(tmp3987)

			tmp3989 := PrimCons(tmp3988, Nil)

			tmp3990 := PrimCons(tmp3983, tmp3989)

			tmp3991 := PrimCons(symand, tmp3990)

			tmp3992 := PrimTail(V5024)

			tmp3993 := PrimTail(tmp3992)

			tmp3994 := PrimHead(tmp3993)

			tmp3995 := PrimTail(tmp3994)

			tmp3996 := PrimTail(tmp3995)

			tmp3997 := PrimHead(tmp3996)

			tmp3998 := PrimTail(V5024)

			tmp3999 := PrimTail(tmp3998)

			tmp4000 := PrimTail(tmp3999)

			tmp4001 := PrimCons(tmp3997, tmp4000)

			tmp4002 := PrimCons(tmp3991, tmp4001)

			__e.Return(PrimCons(symif, tmp4002))
			return

		} else {
			__e.Return(V5024)
			return
		}

	}, 1)

	tmp4091 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs, tmp3980)

	_ = tmp4091

	tmp4092 := MakeNative(func(__e *ControlFlow) {
		V5027 := __e.Get(1)
		_ = V5027
		V5028 := __e.Get(2)
		_ = V5028
		tmp4093 := Call(__e, PrimNS2Value(symconcat), sym_c, V5028)

		__e.TailApply(PrimNS2Value(symconcat), V5027, tmp4093)
		return

	}, 2)

	tmp4094 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4concat_c, tmp4092)

	_ = tmp4094

	tmp4095 := MakeNative(func(__e *ControlFlow) {
		V5032 := __e.Get(1)
		_ = V5032
		tmp4117 := PrimIsPair(V5032)

		var ifres4104 Obj

		if True == tmp4117 {
			tmp4115 := PrimTail(V5032)

			tmp4116 := PrimIsPair(tmp4115)

			var ifres4106 Obj

			if True == tmp4116 {
				tmp4112 := PrimTail(V5032)

				tmp4113 := PrimTail(tmp4112)

				tmp4114 := PrimEqual(Nil, tmp4113)

				var ifres4108 Obj

				if True == tmp4114 {
					tmp4110 := PrimHead(V5032)

					tmp4111 := PrimIsSymbol(tmp4110)

					var ifres4109 Obj

					if True == tmp4111 {
						ifres4109 = True

					} else {
						ifres4109 = False

					}

					ifres4108 = ifres4109

				} else {
					ifres4108 = False

				}

				var ifres4107 Obj

				if True == ifres4108 {
					ifres4107 = True

				} else {
					ifres4107 = False

				}

				ifres4106 = ifres4107

			} else {
				ifres4106 = False

			}

			var ifres4105 Obj

			if True == ifres4106 {
				ifres4105 = True

			} else {
				ifres4105 = False

			}

			ifres4104 = ifres4105

		} else {
			ifres4104 = False

		}

		if True == ifres4104 {
			tmp4100 := PrimTail(V5032)

			tmp4101 := PrimHead(tmp4100)

			tmp4102 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4exp_1var), tmp4101)

			tmp4103 := PrimHead(V5032)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4concat_c), tmp4102, tmp4103)
			return

		} else {
			tmp4099 := PrimIsPair(V5032)

			if True == tmp4099 {
				tmp4098 := PrimHead(V5032)

				__e.TailApply(PrimNS2Value(symgensym), tmp4098)
				return

			} else {
				__e.Return(V5032)
				return
			}

		}

	}, 1)

	tmp4118 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4exp_1var, tmp4095)

	_ = tmp4118

	tmp4119 := MakeNative(func(__e *ControlFlow) {
		V5035 := __e.Get(1)
		_ = V5035
		V5036 := __e.Get(2)
		_ = V5036
		tmp4120 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4test_1_6selectors), V5035)

		__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4bind_1repeating_1selectors), tmp4120, V5036)
		return

	}, 2)

	tmp4121 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4optimize_1selectors, tmp4119)

	_ = tmp4121

	tmp4122 := MakeNative(func(__e *ControlFlow) {
		V5042 := __e.Get(1)
		_ = V5042
		tmp4208 := PrimIsPair(V5042)

		var ifres4195 Obj

		if True == tmp4208 {
			tmp4206 := PrimHead(V5042)

			tmp4207 := PrimEqual(symcons_2, tmp4206)

			var ifres4197 Obj

			if True == tmp4207 {
				tmp4204 := PrimTail(V5042)

				tmp4205 := PrimIsPair(tmp4204)

				var ifres4199 Obj

				if True == tmp4205 {
					tmp4201 := PrimTail(V5042)

					tmp4202 := PrimTail(tmp4201)

					tmp4203 := PrimEqual(Nil, tmp4202)

					var ifres4200 Obj

					if True == tmp4203 {
						ifres4200 = True

					} else {
						ifres4200 = False

					}

					ifres4199 = ifres4200

				} else {
					ifres4199 = False

				}

				var ifres4198 Obj

				if True == ifres4199 {
					ifres4198 = True

				} else {
					ifres4198 = False

				}

				ifres4197 = ifres4198

			} else {
				ifres4197 = False

			}

			var ifres4196 Obj

			if True == ifres4197 {
				ifres4196 = True

			} else {
				ifres4196 = False

			}

			ifres4195 = ifres4196

		} else {
			ifres4195 = False

		}

		if True == ifres4195 {
			tmp4190 := PrimTail(V5042)

			tmp4191 := PrimCons(symhd, tmp4190)

			tmp4192 := PrimTail(V5042)

			tmp4193 := PrimCons(symtl, tmp4192)

			tmp4194 := PrimCons(tmp4193, Nil)

			__e.Return(PrimCons(tmp4191, tmp4194))
			return

		} else {
			tmp4189 := PrimIsPair(V5042)

			var ifres4176 Obj

			if True == tmp4189 {
				tmp4187 := PrimHead(V5042)

				tmp4188 := PrimEqual(symtuple_2, tmp4187)

				var ifres4178 Obj

				if True == tmp4188 {
					tmp4185 := PrimTail(V5042)

					tmp4186 := PrimIsPair(tmp4185)

					var ifres4180 Obj

					if True == tmp4186 {
						tmp4182 := PrimTail(V5042)

						tmp4183 := PrimTail(tmp4182)

						tmp4184 := PrimEqual(Nil, tmp4183)

						var ifres4181 Obj

						if True == tmp4184 {
							ifres4181 = True

						} else {
							ifres4181 = False

						}

						ifres4180 = ifres4181

					} else {
						ifres4180 = False

					}

					var ifres4179 Obj

					if True == ifres4180 {
						ifres4179 = True

					} else {
						ifres4179 = False

					}

					ifres4178 = ifres4179

				} else {
					ifres4178 = False

				}

				var ifres4177 Obj

				if True == ifres4178 {
					ifres4177 = True

				} else {
					ifres4177 = False

				}

				ifres4176 = ifres4177

			} else {
				ifres4176 = False

			}

			if True == ifres4176 {
				tmp4171 := PrimTail(V5042)

				tmp4172 := PrimCons(symfst, tmp4171)

				tmp4173 := PrimTail(V5042)

				tmp4174 := PrimCons(symsnd, tmp4173)

				tmp4175 := PrimCons(tmp4174, Nil)

				__e.Return(PrimCons(tmp4172, tmp4175))
				return

			} else {
				tmp4170 := PrimIsPair(V5042)

				var ifres4157 Obj

				if True == tmp4170 {
					tmp4168 := PrimHead(V5042)

					tmp4169 := PrimEqual(symshen_4_7string_2, tmp4168)

					var ifres4159 Obj

					if True == tmp4169 {
						tmp4166 := PrimTail(V5042)

						tmp4167 := PrimIsPair(tmp4166)

						var ifres4161 Obj

						if True == tmp4167 {
							tmp4163 := PrimTail(V5042)

							tmp4164 := PrimTail(tmp4163)

							tmp4165 := PrimEqual(Nil, tmp4164)

							var ifres4162 Obj

							if True == tmp4165 {
								ifres4162 = True

							} else {
								ifres4162 = False

							}

							ifres4161 = ifres4162

						} else {
							ifres4161 = False

						}

						var ifres4160 Obj

						if True == ifres4161 {
							ifres4160 = True

						} else {
							ifres4160 = False

						}

						ifres4159 = ifres4160

					} else {
						ifres4159 = False

					}

					var ifres4158 Obj

					if True == ifres4159 {
						ifres4158 = True

					} else {
						ifres4158 = False

					}

					ifres4157 = ifres4158

				} else {
					ifres4157 = False

				}

				if True == ifres4157 {
					tmp4152 := PrimTail(V5042)

					tmp4153 := PrimCons(symhdstr, tmp4152)

					tmp4154 := PrimTail(V5042)

					tmp4155 := PrimCons(symtlstr, tmp4154)

					tmp4156 := PrimCons(tmp4155, Nil)

					__e.Return(PrimCons(tmp4153, tmp4156))
					return

				} else {
					tmp4151 := PrimIsPair(V5042)

					var ifres4138 Obj

					if True == tmp4151 {
						tmp4149 := PrimHead(V5042)

						tmp4150 := PrimEqual(symshen_4_7vector_2, tmp4149)

						var ifres4140 Obj

						if True == tmp4150 {
							tmp4147 := PrimTail(V5042)

							tmp4148 := PrimIsPair(tmp4147)

							var ifres4142 Obj

							if True == tmp4148 {
								tmp4144 := PrimTail(V5042)

								tmp4145 := PrimTail(tmp4144)

								tmp4146 := PrimEqual(Nil, tmp4145)

								var ifres4143 Obj

								if True == tmp4146 {
									ifres4143 = True

								} else {
									ifres4143 = False

								}

								ifres4142 = ifres4143

							} else {
								ifres4142 = False

							}

							var ifres4141 Obj

							if True == ifres4142 {
								ifres4141 = True

							} else {
								ifres4141 = False

							}

							ifres4140 = ifres4141

						} else {
							ifres4140 = False

						}

						var ifres4139 Obj

						if True == ifres4140 {
							ifres4139 = True

						} else {
							ifres4139 = False

						}

						ifres4138 = ifres4139

					} else {
						ifres4138 = False

					}

					if True == ifres4138 {
						tmp4133 := PrimTail(V5042)

						tmp4134 := PrimCons(symhdv, tmp4133)

						tmp4135 := PrimTail(V5042)

						tmp4136 := PrimCons(symtlv, tmp4135)

						tmp4137 := PrimCons(tmp4136, Nil)

						__e.Return(PrimCons(tmp4134, tmp4137))
						return

					} else {
						tmp4127 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp4129 := Call(__e, PrimNS2Value(symfail))

							tmp4130 := PrimEqual(Result, tmp4129)

							if True == tmp4130 {
								__e.Return(Nil)
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp4131 := PrimNS3Value(symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

						tmp4132 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4apply_1selector_1handlers), tmp4131, V5042)

						__e.TailApply(tmp4127, tmp4132)
						return

					}

				}

			}

		}

	}, 1)

	tmp4209 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4test_1_6selectors, tmp4122)

	_ = tmp4209

	tmp4210 := MakeNative(func(__e *ControlFlow) {
		V5045 := __e.Get(1)
		_ = V5045
		V5046 := __e.Get(2)
		_ = V5046
		tmp4217 := PrimIsPair(V5045)

		if True == tmp4217 {
			tmp4214 := PrimHead(V5045)

			tmp4215 := PrimTail(V5045)

			tmp4216 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4bind_1repeating_1selectors), tmp4215, V5046)

			__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4bind_1selector), tmp4214, tmp4216)
			return

		} else {
			tmp4213 := PrimEqual(Nil, V5045)

			if True == tmp4213 {
				__e.Return(V5046)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors)
				return
			}

		}

	}, 2)

	tmp4218 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4bind_1repeating_1selectors, tmp4210)

	_ = tmp4218

	tmp4219 := MakeNative(func(__e *ControlFlow) {
		V5053 := __e.Get(1)
		_ = V5053
		V5054 := __e.Get(2)
		_ = V5054
		tmp4227 := Call(__e, PrimNS2Value(symoccurrences), V5053, V5054)

		tmp4228 := PrimGreatThan(tmp4227, MakeNumber(1))

		if True == tmp4228 {
			tmp4221 := MakeNative(func(__e *ControlFlow) {
				Var := __e.Get(1)
				_ = Var
				tmp4222 := Call(__e, PrimNS2Value(symsubst), Var, V5053, V5054)

				tmp4223 := PrimCons(tmp4222, Nil)

				tmp4224 := PrimCons(V5053, tmp4223)

				tmp4225 := PrimCons(Var, tmp4224)

				__e.Return(PrimCons(symlet, tmp4225))
				return

			}, 1)

			tmp4226 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4exp_1var), V5053)

			__e.TailApply(tmp4221, tmp4226)
			return

		} else {
			__e.Return(V5054)
			return
		}

	}, 2)

	tmp4229 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4bind_1selector, tmp4219)

	_ = tmp4229

	tmp4230 := MakeNative(func(__e *ControlFlow) {
		V5067 := __e.Get(1)
		_ = V5067
		V5068 := __e.Get(2)
		_ = V5068
		tmp4245 := PrimEqual(Nil, V5067)

		if True == tmp4245 {
			__e.TailApply(PrimNS2Value(symfail))
			return
		} else {
			tmp4232 := MakeNative(func(__e *ControlFlow) {
				Freeze := __e.Get(1)
				_ = Freeze
				tmp4240 := PrimIsPair(V5067)

				if True == tmp4240 {
					tmp4234 := MakeNative(func(__e *ControlFlow) {
						Result := __e.Get(1)
						_ = Result
						tmp4236 := Call(__e, PrimNS2Value(symfail))

						tmp4237 := PrimEqual(Result, tmp4236)

						if True == tmp4237 {
							__e.TailApply(PrimNS2Value(symthaw), Freeze)
							return
						} else {
							__e.Return(Result)
							return
						}

					}, 1)

					tmp4238 := PrimHead(V5067)

					tmp4239 := Call(__e, tmp4238, V5068)

					__e.TailApply(tmp4234, tmp4239)
					return

				} else {
					__e.TailApply(PrimNS2Value(symthaw), Freeze)
					return
				}

			}, 1)

			tmp4241 := MakeNative(func(__e *ControlFlow) {
				tmp4244 := PrimIsPair(V5067)

				if True == tmp4244 {
					tmp4243 := PrimTail(V5067)

					__e.TailApply(PrimNS2Value(symshen_4x_4factorise_1defun_4apply_1selector_1handlers), tmp4243, V5068)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4x_4factorise_1defun_4apply_1selector_1handlers)
					return
				}

			}, 0)

			__e.TailApply(tmp4232, tmp4241)
			return

		}

	}, 2)

	tmp4246 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4apply_1selector_1handlers, tmp4230)

	_ = tmp4246

	tmp4247 := MakeNative(func(__e *ControlFlow) {
		tmp4248 := PrimNS3Set(symshen_4x_4factorise_1defun_4_dselector_1handlers_d, Nil)

		_ = tmp4248

		tmp4249 := PrimNS3Set(symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, Nil)

		_ = tmp4249

		__e.Return(symshen_4x_4factorise_1defun_4done)
		return

	}, 0)

	tmp4250 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4initialise, tmp4247)

	_ = tmp4250

	tmp4251 := MakeNative(func(__e *ControlFlow) {
		V5070 := __e.Get(1)
		_ = V5070
		tmp4260 := PrimNS3Value(symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

		tmp4261 := Call(__e, PrimNS2Value(symelement_2), V5070, tmp4260)

		if True == tmp4261 {
			__e.Return(V5070)
			return
		} else {
			tmp4253 := PrimNS3Value(symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

			tmp4254 := PrimCons(V5070, tmp4253)

			tmp4255 := PrimNS3Set(symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, tmp4254)

			_ = tmp4255

			tmp4256 := Call(__e, PrimNS2Value(symfunction), V5070)

			tmp4257 := PrimNS3Value(symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

			tmp4258 := PrimCons(tmp4256, tmp4257)

			tmp4259 := PrimNS3Set(symshen_4x_4factorise_1defun_4_dselector_1handlers_d, tmp4258)

			_ = tmp4259

			__e.Return(V5070)
			return

		}

	}, 1)

	tmp4262 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4register_1selector_1handler, tmp4251)

	_ = tmp4262

	tmp4263 := MakeNative(func(__e *ControlFlow) {
		V5073 := __e.Get(1)
		_ = V5073
		V5074 := __e.Get(2)
		_ = V5074
		tmp4264 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4findpos), V5073, V5074)
			return
		}, 0)

		tmp4265 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			tmp4266 := Call(__e, PrimNS2Value(symshen_4app), V5073, MakeString(" is not a selector handler\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp4266))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp4264, tmp4265)
		return

	}, 2)

	tmp4267 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4findpos, tmp4263)

	_ = tmp4267

	tmp4268 := MakeNative(func(__e *ControlFlow) {
		V5076 := __e.Get(1)
		_ = V5076
		tmp4269 := MakeNative(func(__e *ControlFlow) {
			Reg := __e.Get(1)
			_ = Reg
			tmp4270 := MakeNative(func(__e *ControlFlow) {
				Pos := __e.Get(1)
				_ = Pos
				tmp4271 := MakeNative(func(__e *ControlFlow) {
					RemoveReg := __e.Get(1)
					_ = RemoveReg
					tmp4272 := MakeNative(func(__e *ControlFlow) {
						RemoveFun := __e.Get(1)
						_ = RemoveFun
						__e.Return(V5076)
						return
					}, 1)

					tmp4273 := PrimNS3Value(symshen_4x_4factorise_1defun_4_dselector_1handlers_d)

					tmp4274 := Call(__e, PrimNS2Value(symshen_4remove_1nth), Pos, tmp4273)

					tmp4275 := PrimNS3Set(symshen_4x_4factorise_1defun_4_dselector_1handlers_d, tmp4274)

					__e.TailApply(tmp4272, tmp4275)
					return

				}, 1)

				tmp4276 := Call(__e, PrimNS2Value(symremove), V5076, Reg)

				tmp4277 := PrimNS3Set(symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d, tmp4276)

				__e.TailApply(tmp4271, tmp4277)
				return

			}, 1)

			tmp4278 := Call(__e, PrimNS2Value(symshen_4x_4factorise_1defun_4findpos), V5076, Reg)

			__e.TailApply(tmp4270, tmp4278)
			return

		}, 1)

		tmp4279 := PrimNS3Value(symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d)

		__e.TailApply(tmp4269, tmp4279)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4factorise_1defun_4unregister_1selector_1handler, tmp4268)
	return

}, 0)

var symshen_4eval_1cons = MakeSymbol("shen.eval-cons")
var symshen_4_5st__input2_6 = MakeSymbol("shen.<st_input2>")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symos = MakeSymbol("os")
var symlazy = MakeSymbol("lazy")
var symshen_4_5end_d_6 = MakeSymbol("shen.<end*>")
var symshen_4type_1signature_1of_1release = MakeSymbol("shen.type-signature-of-release")
var symshen_4type_1signature_1of_1specialise = MakeSymbol("shen.type-signature-of-specialise")
var symshen_4type_1signature_1of_1sum = MakeSymbol("shen.type-signature-of-sum")
var symshen_4explicit__modes = MakeSymbol("shen.explicit_modes")
var symshen_4packageh = MakeSymbol("shen.packageh")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symimplementation = MakeSymbol("implementation")
var symshen_4type_1signature_1of_1hdstr = MakeSymbol("shen.type-signature-of-hdstr")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4reverse__help = MakeSymbol("shen.reverse_help")
var symshen_4toplineread__loop = MakeSymbol("shen.toplineread_loop")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symif = MakeSymbol("if")
var symthaw = MakeSymbol("thaw")
var syminput = MakeSymbol("input")
var symnumber = MakeSymbol("number")
var symshen_4type_1signature_1of_1fail = MakeSymbol("shen.type-signature-of-fail")
var symshen_4recursive__descent = MakeSymbol("shen.recursive_descent")
var symshen_4newline = MakeSymbol("shen.newline")
var symelement_2 = MakeSymbol("element?")
var symshen_4free__variable__warnings = MakeSymbol("shen.free_variable_warnings")
var symshen_4compose = MakeSymbol("shen.compose")
var symU = MakeSymbol("U")
var symshen_4be = MakeSymbol("shen.be")
var symshen_4type_1signature_1of_1read = MakeSymbol("shen.type-signature-of-read")
var symshen_4initialise_1prolog = MakeSymbol("shen.initialise-prolog")
var symshen_4analyse_1kill = MakeSymbol("shen.analyse-kill")
var symshen_4compile__to__kl = MakeSymbol("shen.compile_to_kl")
var symcond = MakeSymbol("cond")
var symshen_4type_1signature_1of_1pr = MakeSymbol("shen.type-signature-of-pr")
var symshen_4embed_1and = MakeSymbol("shen.embed-and")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4aritycheck = MakeSymbol("shen.aritycheck")
var symshen_4type_1signature_1of_1_1 = MakeSymbol("shen.type-signature-of--")
var symshen_4type_1signature_1of_1stinput = MakeSymbol("shen.type-signature-of-stinput")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symtail = MakeSymbol("tail")
var sympr = MakeSymbol("pr")
var symexplode = MakeSymbol("explode")
var symshen_4copy_1vector_1stage_12 = MakeSymbol("shen.copy-vector-stage-2")
var symshen_4preclude_1h = MakeSymbol("shen.preclude-h")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4hdhd = MakeSymbol("shen.hdhd")
var symshen_4type_1signature_1of_1spy = MakeSymbol("shen.type-signature-of-spy")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4_5anymulti_6 = MakeSymbol("shen.<anymulti>")
var symshen_4x_4features_4current = MakeSymbol("shen.x.features.current")
var symshen_4type_1signature_1of_1snd = MakeSymbol("shen.type-signature-of-snd")
var symshen_4non_1empty = MakeSymbol("shen.non-empty")
var symshen_4read_1char_1code = MakeSymbol("shen.read-char-code")
var symshen_4check__stream = MakeSymbol("shen.check_stream")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symstring = MakeSymbol("string")
var symMessage = MakeSymbol("Message")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4print_1past_1inputs = MakeSymbol("shen.print-past-inputs")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4active_1cons = MakeSymbol("shen.active-cons")
var symloaded = MakeSymbol("loaded")
var symshen_4type_1signature_1of_1shen_4proc_1nl = MakeSymbol("shen.type-signature-of-shen.proc-nl")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4mult__subst = MakeSymbol("shen.mult_subst")
var symshen_4insert_1runon = MakeSymbol("shen.insert-runon")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4group__clauses = MakeSymbol("shen.group_clauses")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4unbindv = MakeSymbol("shen.unbindv")
var symunify_b = MakeSymbol("unify!")
var symshen_4removetype = MakeSymbol("shen.removetype")
var symStart = MakeSymbol("Start")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4ephemeral__variable_2 = MakeSymbol("shen.ephemeral_variable?")
var symshen_4construct_1search_1clause = MakeSymbol("shen.construct-search-clause")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var sym_d = MakeSymbol("*")
var symshen_4aah = MakeSymbol("shen.aah")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4type_1signature_1of_1fail_1if = MakeSymbol("shen.type-signature-of-fail-if")
var symshen_4prbytes = MakeSymbol("shen.prbytes")
var symabsvector = MakeSymbol("absvector")
var symshen_4type_1signature_1of_1explode = MakeSymbol("shen.type-signature-of-explode")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4hat = MakeSymbol("shen.hat")
var symshen_4_5side_1conditions_6 = MakeSymbol("shen.<side-conditions>")
var symshen_4split__cc__rules = MakeSymbol("shen.split_cc_rules")
var symshen_4initialise__arity__table = MakeSymbol("shen.initialise_arity_table")
var sym_i = MakeSymbol("{")
var symshen_4err_1condition = MakeSymbol("shen.err-condition")
var symshen_4sigf = MakeSymbol("shen.sigf")
var sym_6 = MakeSymbol(">")
var symshen_4_5premise_6 = MakeSymbol("shen.<premise>")
var symshen_4exclamation = MakeSymbol("shen.exclamation")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4_5side_1condition_6 = MakeSymbol("shen.<side-condition>")
var symshen_4dh_2 = MakeSymbol("shen.dh?")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var sym_h_1 = MakeSymbol(":-")
var symshen_4type_1signature_1of_1_a_a = MakeSymbol("shen.type-signature-of-==")
var symCase = MakeSymbol("Case")
var symshen_4find_1past_1inputs = MakeSymbol("shen.find-past-inputs")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4jump__stream_2 = MakeSymbol("shen.jump_stream?")
var symshen_4cond_1expression = MakeSymbol("shen.cond-expression")
var symshen_4type_1signature_1of_1get_1time = MakeSymbol("shen.type-signature-of-get-time")
var symshen_4failed_b = MakeSymbol("shen.failed!")
var symshen_4the = MakeSymbol("shen.the")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symlambda = MakeSymbol("lambda")
var symshen_4abstraction__build = MakeSymbol("shen.abstraction_build")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4type_1signature_1of_1compile = MakeSymbol("shen.type-signature-of-compile")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4_dteststack_d = MakeSymbol("shen.*teststack*")
var symversion = MakeSymbol("version")
var symL = MakeSymbol("L")
var symshen_4type_1signature_1of_1untrack = MakeSymbol("shen.type-signature-of-untrack")
var symshen_4copy_1vector = MakeSymbol("shen.copy-vector")
var symshen_4code_1point = MakeSymbol("shen.code-point")
var symspecialise = MakeSymbol("specialise")
var symshen_4_5guard_6 = MakeSymbol("shen.<guard>")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symshen_4type_1signature_1of_1_7 = MakeSymbol("shen.type-signature-of-+")
var symshen_4prefix_2 = MakeSymbol("shen.prefix?")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var sym_c = MakeSymbol("/")
var symshen_4type_1signature_1of_1shen_4prhush = MakeSymbol("shen.type-signature-of-shen.prhush")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4resize_1vector = MakeSymbol("shen.resize-vector")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var symshen_4prolog__constant_2 = MakeSymbol("shen.prolog_constant?")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4x_4factorise_1defun_4test_1_6selectors = MakeSymbol("shen.x.factorise-defun.test->selectors")
var symsuccess = MakeSymbol("success")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4x_4factorise_1defun_4free_1variables_1h = MakeSymbol("shen.x.factorise-defun.free-variables-h")
var symstr = MakeSymbol("str")
var symshen_4type_1signature_1of_1cd = MakeSymbol("shen.type-signature-of-cd")
var symshen_4type_1signature_1of_1integer_2 = MakeSymbol("shen.type-signature-of-integer?")
var symshen_4type_1signature_1of_1variable_2 = MakeSymbol("shen.type-signature-of-variable?")
var symshen_4procedure__name = MakeSymbol("shen.procedure_name")
var symshen_4fail__if = MakeSymbol("shen.fail_if")
var symeval = MakeSymbol("eval")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4extract_1pvars = MakeSymbol("shen.extract-pvars")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symshen_4call_1rest = MakeSymbol("shen.call-rest")
var symshen_4doubleunderline_2 = MakeSymbol("shen.doubleunderline?")
var symshen_4type_1signature_1of_1string_2 = MakeSymbol("shen.type-signature-of-string?")
var symport = MakeSymbol("port")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symconcat = MakeSymbol("concat")
var symtype = MakeSymbol("type")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4x_4factorise_1defun_4bind_1selector = MakeSymbol("shen.x.factorise-defun.bind-selector")
var symshen_4compile_1macro = MakeSymbol("shen.compile-macro")
var sympackage_2 = MakeSymbol("package?")
var symshen_4type_1signature_1of_1close = MakeSymbol("shen.type-signature-of-close")
var symContinuation = MakeSymbol("Continuation")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symreceive = MakeSymbol("receive")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4pair = MakeSymbol("shen.pair")
var symshen_4placeholder = MakeSymbol("shen.placeholder")
var symshen_4x_4factorise_1defun_4concat_c = MakeSymbol("shen.x.factorise-defun.concat/")
var symshen_4extract__free__vars = MakeSymbol("shen.extract_free_vars")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4to = MakeSymbol("shen.to")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symput = MakeSymbol("put")
var sym_7 = MakeSymbol("+")
var symcn = MakeSymbol("cn")
var symshen_4list_1stream = MakeSymbol("shen.list-stream")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4type_1signature_1of_1if = MakeSymbol("shen.type-signature-of-if")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4skip = MakeSymbol("shen.skip")
var symporters = MakeSymbol("porters")
var symshen_4type_1signature_1of_1assoc = MakeSymbol("shen.type-signature-of-assoc")
var symshen_4single = MakeSymbol("shen.single")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4memo = MakeSymbol("shen.memo")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4unix = MakeSymbol("shen.unix")
var symshen_4elim_1def = MakeSymbol("shen.elim-def")
var symshen_4mode_1ify = MakeSymbol("shen.mode-ify")
var symshen_4include_1h = MakeSymbol("shen.include-h")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4insert_1deref = MakeSymbol("shen.insert-deref")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4x_4factorise_1defun_4initialise = MakeSymbol("shen.x.factorise-defun.initialise")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var sym_6_a = MakeSymbol(">=")
var symshen_4type_1signature_1of_1tail = MakeSymbol("shen.type-signature-of-tail")
var symshen_4type_1signature_1of_1read_1file_1as_1bytelist = MakeSymbol("shen.type-signature-of-read-file-as-bytelist")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symshen_4cc__body = MakeSymbol("shen.cc_body")
var symshen_4prolog_1aritycheck = MakeSymbol("shen.prolog-aritycheck")
var symshen_4_5comma_1symbol_6 = MakeSymbol("shen.<comma-symbol>")
var symContext = MakeSymbol("Context")
var symunspecialise = MakeSymbol("unspecialise")
var symsubst = MakeSymbol("subst")
var symshen_4type_1signature_1of_1optimise = MakeSymbol("shen.type-signature-of-optimise")
var symshen_4type_1signature_1of_1trap_1error = MakeSymbol("shen.type-signature-of-trap-error")
var symshen_4_5non_1return_6 = MakeSymbol("shen.<non-return>")
var symshen_4_dcustom_1pattern_1compiler_d = MakeSymbol("shen.*custom-pattern-compiler*")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symuntrack = MakeSymbol("untrack")
var symshen_4abstract__rule = MakeSymbol("shen.abstract_rule")
var symshen_4type_1signature_1of_1_d = MakeSymbol("shen.type-signature-of-*")
var symadjoin = MakeSymbol("adjoin")
var symshen_4_5signature_1help_6 = MakeSymbol("shen.<signature-help>")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symshen_4read_1file_1as_1Xlist = MakeSymbol("shen.read-file-as-Xlist")
var symshen_4_5action_6 = MakeSymbol("shen.<action>")
var symload = MakeSymbol("load")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symtry_1catch = MakeSymbol("try-catch")
var symshen_4type_1signature_1of_1_5_a = MakeSymbol("shen.type-signature-of-<=")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4function_1abstraction_1help = MakeSymbol("shen.function-abstraction-help")
var symshen_4clause__form = MakeSymbol("shen.clause_form")
var symN = MakeSymbol("N")
var symshen_4type_1signature_1of_1y_1or_1n_2 = MakeSymbol("shen.type-signature-of-y-or-n?")
var symshen_4start_1new_1prolog_1process = MakeSymbol("shen.start-new-prolog-process")
var symshen_4newcontinuation = MakeSymbol("shen.newcontinuation")
var symshen_4base = MakeSymbol("shen.base")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symfreeze = MakeSymbol("freeze")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symshen_4type_1signature_1of_1do = MakeSymbol("shen.type-signature-of-do")
var symshen_4type_1signature_1of_1track = MakeSymbol("shen.type-signature-of-track")
var symshen_4cc__help = MakeSymbol("shen.cc_help")
var symshen_4valvector = MakeSymbol("shen.valvector")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var sym_8p = MakeSymbol("@p")
var symfindall = MakeSymbol("findall")
var symdifference = MakeSymbol("difference")
var symshen_4rules_1_6horn_1clauses = MakeSymbol("shen.rules->horn-clauses")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4x_4factorise_1defun_4rebranch_1h = MakeSymbol("shen.x.factorise-defun.rebranch-h")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symunit = MakeSymbol("unit")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4rule_1_6horn_1clause_1head = MakeSymbol("shen.rule->horn-clause-head")
var symZ = MakeSymbol("Z")
var symshen_4mk_1pvar = MakeSymbol("shen.mk-pvar")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4x_4factorise_1defun_4_dselector_1handlers_d = MakeSymbol("shen.x.factorise-defun.*selector-handlers*")
var symA = MakeSymbol("A")
var symshen_4of = MakeSymbol("shen.of")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symhdstr = MakeSymbol("hdstr")
var symshen_4record_1exceptions = MakeSymbol("shen.record-exceptions")
var symNewStream = MakeSymbol("NewStream")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4result_1type = MakeSymbol("shen.result-type")
var symtl = MakeSymbol("tl")
var symshen_4cf__help = MakeSymbol("shen.cf_help")
var symshen_4ue_1sig = MakeSymbol("shen.ue-sig")
var symshen_4type_1signature_1of_1package_2 = MakeSymbol("shen.type-signature-of-package?")
var symshen_4type_1signature_1of_1preclude_1all_1but = MakeSymbol("shen.type-signature-of-preclude-all-but")
var symshen_4locally_1bind_1prolog_1vs = MakeSymbol("shen.locally-bind-prolog-vs")
var symshen_4_5body_d_6 = MakeSymbol("shen.<body*>")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var sym_j = MakeSymbol("}")
var symO = MakeSymbol("O")
var symshen_4x_4factorise_1defun_4apply_1selector_1handlers = MakeSymbol("shen.x.factorise-defun.apply-selector-handlers")
var symshen_4_5anysingle_6 = MakeSymbol("shen.<anysingle>")
var symwhen = MakeSymbol("when")
var symshen_4safe_1product = MakeSymbol("shen.safe-product")
var symshen_4_5E_6 = MakeSymbol("shen.<E>")
var symshen_4_5variable_2_6 = MakeSymbol("shen.<variable?>")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var sym_h_a = MakeSymbol(":=")
var symshen_4_5doubleunderline_6 = MakeSymbol("shen.<doubleunderline>")
var symshen_4typecheck_1and_1evaluate = MakeSymbol("shen.typecheck-and-evaluate")
var symshen_4t_d_1action = MakeSymbol("shen.t*-action")
var symshen_4x_4factorise_1defun_4done = MakeSymbol("shen.x.factorise-defun.done")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4type_1signature_1of_1fix = MakeSymbol("shen.type-signature-of-fix")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4x_4features_4initialise = MakeSymbol("shen.x.features.initialise")
var symshen_4type_1signature_1of_1or = MakeSymbol("shen.type-signature-of-or")
var symshen_4aum = MakeSymbol("shen.aum")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symdo = MakeSymbol("do")
var sym_l = MakeSymbol(",")
var symshen_4x_4factorise_1defun_4unregister_1selector_1handler = MakeSymbol("shen.x.factorise-defun.unregister-selector-handler")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4type_1signature_1of_1freeze = MakeSymbol("shen.type-signature-of-freeze")
var symshen_4type_1signature_1of_1vector_2 = MakeSymbol("shen.type-signature-of-vector?")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4_dprologvectors_d = MakeSymbol("shen.*prologvectors*")
var symtrack = MakeSymbol("track")
var symcd = MakeSymbol("cd")
var symshen_4type_1signature_1of_1tc_2 = MakeSymbol("shen.type-signature-of-tc?")
var symIn__1957 = MakeSymbol("In_1957")
var symnth = MakeSymbol("nth")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var sym_a_b = MakeSymbol("=!")
var sym_1 = MakeSymbol("-")
var symshen_4update_1demodulation_1function = MakeSymbol("shen.update-demodulation-function")
var sym_dlanguage_d = MakeSymbol("*language*")
var symshen_4type_1signature_1of_1cons_2 = MakeSymbol("shen.type-signature-of-cons?")
var symContext__1957 = MakeSymbol("Context_1957")
var symshen_4errormaxinfs = MakeSymbol("shen.errormaxinfs")
var symshen_4r = MakeSymbol("shen.r")
var symunify = MakeSymbol("unify")
var symshen_4complexity__head = MakeSymbol("shen.complexity_head")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symshen_4type_1signature_1of_1tlstr = MakeSymbol("shen.type-signature-of-tlstr")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4s_1prolog__literal = MakeSymbol("shen.s-prolog_literal")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symsum = MakeSymbol("sum")
var symfix = MakeSymbol("fix")
var symshen_4x_4factorise_1defun_4_dselector_1handlers_1reg_d = MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4head__abstraction = MakeSymbol("shen.head_abstraction")
var symfunction = MakeSymbol("function")
var symout = MakeSymbol("out")
var symshen_4type_1signature_1of_1map = MakeSymbol("shen.type-signature-of-map")
var symshen_4update__history = MakeSymbol("shen.update_history")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symE = MakeSymbol("E")
var symshen_4type_1signature_1of_1unspecialise = MakeSymbol("shen.type-signature-of-unspecialise")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4type_1signature_1of_1vector = MakeSymbol("shen.type-signature-of-vector")
var symshen_4semantics = MakeSymbol("shen.semantics")
var symFinish = MakeSymbol("Finish")
var symProcessN = MakeSymbol("ProcessN")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4hdtl = MakeSymbol("shen.hdtl")
var sympreclude = MakeSymbol("preclude")
var symsystemf = MakeSymbol("systemf")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4_5predigits_6 = MakeSymbol("shen.<predigits>")
var symshen_4recursive__cons__form = MakeSymbol("shen.recursive_cons_form")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4x_4factorise_1defun_4exp_1var = MakeSymbol("shen.x.factorise-defun.exp-var")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4type_1signature_1of_1include_1all_1but = MakeSymbol("shen.type-signature-of-include-all-but")
var symshen_4type_1signature_1of_1maxinferences = MakeSymbol("shen.type-signature-of-maxinferences")
var symshen_4_5st__input1_6 = MakeSymbol("shen.<st_input1>")
var symmake_1string = MakeSymbol("make-string")
var symshen_4type_1signature_1of_1adjoin = MakeSymbol("shen.type-signature-of-adjoin")
var symshen_4type_1signature_1of_1preclude = MakeSymbol("shen.type-signature-of-preclude")
var symshen_4_5clauses_d_6 = MakeSymbol("shen.<clauses*>")
var symshen_4assign_1types = MakeSymbol("shen.assign-types")
var symM = MakeSymbol("M")
var syminclude = MakeSymbol("include")
var symshen_4left_1round = MakeSymbol("shen.left-round")
var sympackage = MakeSymbol("package")
var symshen_4compile__to__lambda_7 = MakeSymbol("shen.compile_to_lambda+")
var symshen_4read_1error = MakeSymbol("shen.read-error")
var symW = MakeSymbol("W")
var symset = MakeSymbol("set")
var symshen_4type_1signature_1of_1shen_4app = MakeSymbol("shen.type-signature-of-shen.app")
var symshen_4em__help = MakeSymbol("shen.em_help")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symtrap_1error = MakeSymbol("trap-error")
var symabort = MakeSymbol("abort")
var symcases = MakeSymbol("cases")
var symshen_4compile__prolog__procedure = MakeSymbol("shen.compile_prolog_procedure")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4find = MakeSymbol("shen.find")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4extract__vars = MakeSymbol("shen.extract_vars")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symshen_4strip_1protect = MakeSymbol("shen.strip-protect")
var symshen_4trim_1whitespace = MakeSymbol("shen.trim-whitespace")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symQ = MakeSymbol("Q")
var symshen_4remove_1synonyms = MakeSymbol("shen.remove-synonyms")
var symshen_4remtype = MakeSymbol("shen.remtype")
var symshen_4void = MakeSymbol("shen.void")
var sym_f_fgoto_1label = MakeSymbol("%%goto-label")
var symshen_4type_1signature_1of_1stoutput = MakeSymbol("shen.type-signature-of-stoutput")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4linearise__help = MakeSymbol("shen.linearise_help")
var symshen_4require = MakeSymbol("shen.require")
var symshen_4type_1signature_1of_1sterror = MakeSymbol("shen.type-signature-of-sterror")
var symshen_4insert_1lazyderef = MakeSymbol("shen.insert-lazyderef")
var symshen_4by__hypothesis = MakeSymbol("shen.by_hypothesis")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var sym_8v = MakeSymbol("@v")
var symshen_4add__test = MakeSymbol("shen.add_test")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4type_1signature_1of_1read_1file_1as_1string = MakeSymbol("shen.type-signature-of-read-file-as-string")
var symshen_4trim_1gubbins = MakeSymbol("shen.trim-gubbins")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs = MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs")
var symshen_4x_4factorise_1defun_4findpos = MakeSymbol("shen.x.factorise-defun.findpos")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4type_1signature_1of_1kill = MakeSymbol("shen.type-signature-of-kill")
var symshen_4type_1signature_1of_1os = MakeSymbol("shen.type-signature-of-os")
var symshen_4t_d_1patterns = MakeSymbol("shen.t*-patterns")
var symshen_4default__semantics = MakeSymbol("shen.default_semantics")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4remove__modes = MakeSymbol("shen.remove_modes")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4typedf_2 = MakeSymbol("shen.typedf?")
var symshen_4type_1signature_1of_1hash = MakeSymbol("shen.type-signature-of-hash")
var symshen_4_5clause_d_6 = MakeSymbol("shen.<clause*>")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symResult = MakeSymbol("Result")
var symshen_4type_1signature_1of_1_a = MakeSymbol("shen.type-signature-of-=")
var symshen_4type_1signature_1of_1remove = MakeSymbol("shen.type-signature-of-remove")
var symshen_4findallhelp = MakeSymbol("shen.findallhelp")
var symshen_4decons = MakeSymbol("shen.decons")
var symtc = MakeSymbol("tc")
var symverified = MakeSymbol("verified")
var symshen_4decimalise = MakeSymbol("shen.decimalise")
var symshen_4construct_1base_1search_1clause = MakeSymbol("shen.construct-base-search-clause")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4type_1signature_1of_1and = MakeSymbol("shen.type-signature-of-and")
var symshen_4type_1signature_1of_1not = MakeSymbol("shen.type-signature-of-not")
var symshen_4type_1signature_1of_1read_1byte = MakeSymbol("shen.type-signature-of-read-byte")
var symshen_4nest_1disjunct = MakeSymbol("shen.nest-disjunct")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symX = MakeSymbol("X")
var symshen_4_dcatch_d = MakeSymbol("shen.*catch*")
var symshen_4control_1chars = MakeSymbol("shen.control-chars")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symshen_4case_1form = MakeSymbol("shen.case-form")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4toplevel = MakeSymbol("shen.toplevel")
var symshen_4lambda_1form_1entry = MakeSymbol("shen.lambda-form-entry")
var sym_a = MakeSymbol("=")
var symrelease = MakeSymbol("release")
var symshen_4decons_1string = MakeSymbol("shen.decons-string")
var symshen_4csl_1help = MakeSymbol("shen.csl-help")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4x_4factorise_1defun_4free_1variables = MakeSymbol("shen.x.factorise-defun.free-variables")
var symprofile_1results = MakeSymbol("profile-results")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4type_1signature_1of_1print = MakeSymbol("shen.type-signature-of-print")
var symcons = MakeSymbol("cons")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4variant_2 = MakeSymbol("shen.variant?")
var symstep = MakeSymbol("step")
var symprint = MakeSymbol("print")
var symfile = MakeSymbol("file")
var symshen_4type_1signature_1of_1version = MakeSymbol("shen.type-signature-of-version")
var symshen_4externally_1bound = MakeSymbol("shen.externally-bound")
var symbar_b = MakeSymbol("bar!")
var symshen_4clash_2 = MakeSymbol("shen.clash?")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symshen_4mu__reduction = MakeSymbol("shen.mu_reduction")
var symshen_4retrieve_1from_1history_1if_1needed = MakeSymbol("shen.retrieve-from-history-if-needed")
var symshen_4package_1contents = MakeSymbol("shen.package-contents")
var sym_5_1_1 = MakeSymbol("<--")
var symQv = MakeSymbol("Qv")
var symAssumption__1957 = MakeSymbol("Assumption_1957")
var symor = MakeSymbol("or")
var symshen_4type_1signature_1of_1append = MakeSymbol("shen.type-signature-of-append")
var symThrowcontrol = MakeSymbol("Throwcontrol")
var symH = MakeSymbol("H")
var symshen_4_5sig_7rest_6 = MakeSymbol("shen.<sig+rest>")
var symmacroexpand = MakeSymbol("macroexpand")
var sym_dport_d = MakeSymbol("*port*")
var symshen_4type_1signature_1of_1step = MakeSymbol("shen.type-signature-of-step")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4_dalphabet_d = MakeSymbol("shen.*alphabet*")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4nextline = MakeSymbol("shen.nextline")
var symshen_4newhyps = MakeSymbol("shen.newhyps")
var symmap = MakeSymbol("map")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4measure_ereturn = MakeSymbol("shen.measure&return")
var symshen_4_5sig_7rules_6 = MakeSymbol("shen.<sig+rules>")
var symshen_4x_4factorise_1defun_4add_1returns = MakeSymbol("shen.x.factorise-defun.add-returns")
var sym_f_freturn = MakeSymbol("%%return")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symunprofile = MakeSymbol("unprofile")
var symshen_4aum__to__shen = MakeSymbol("shen.aum_to_shen")
var symshen_4snd_1or_1fail = MakeSymbol("shen.snd-or-fail")
var symshen_4list__variables = MakeSymbol("shen.list_variables")
var symTime = MakeSymbol("Time")
var symshen_4catch_1cut = MakeSymbol("shen.catch-cut")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4proc_1input_7 = MakeSymbol("shen.proc-input+")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4digits_1_6integers = MakeSymbol("shen.digits->integers")
var sym_dsterror_d = MakeSymbol("*sterror*")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var sym_b = MakeSymbol("!")
var symshen_4intprolog = MakeSymbol("shen.intprolog")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4atom_1type = MakeSymbol("shen.atom-type")
var symshen_4construct_1search_1clauses = MakeSymbol("shen.construct-search-clauses")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4x_4features_4cond_1expand_1macro = MakeSymbol("shen.x.features.cond-expand-macro")
var symshen_4type_1signature_1of_1load = MakeSymbol("shen.type-signature-of-load")
var symreverse = MakeSymbol("reverse")
var symexternal = MakeSymbol("external")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symshen_4remove_1bar = MakeSymbol("shen.remove-bar")
var symshen_4rule_1_6horn_1clause_1body = MakeSymbol("shen.rule->horn-clause-body")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symStream = MakeSymbol("Stream")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4type_1signature_1of_1_5_b_6 = MakeSymbol("shen.type-signature-of-<!>")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4multiple_1set = MakeSymbol("shen.multiple-set")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4type_1signature_1of_1element_2 = MakeSymbol("shen.type-signature-of-element?")
var symshen_4type_1signature_1of_1union = MakeSymbol("shen.type-signature-of-union")
var symshen_4resizeprocessvector = MakeSymbol("shen.resizeprocessvector")
var symhash = MakeSymbol("hash")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symtlstr = MakeSymbol("tlstr")
var symshen_4_5st__input_6 = MakeSymbol("shen.<st_input>")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4type_1signature_1of_1str = MakeSymbol("shen.type-signature-of-str")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symfst = MakeSymbol("fst")
var symshen_4kill_1code = MakeSymbol("shen.kill-code")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4type_1signature_1of_1function = MakeSymbol("shen.type-signature-of-function")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4initialise_1signedfunc_1lambda_1forms = MakeSymbol("shen.initialise-signedfunc-lambda-forms")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var sym_5_1 = MakeSymbol("<-")
var symshen_4type_1signature_1of_1length = MakeSymbol("shen.type-signature-of-length")
var symshen_4type_1signature_1of_1nl = MakeSymbol("shen.type-signature-of-nl")
var symshen_4lookup_1func = MakeSymbol("shen.lookup-func")
var symshen_4x_4factorise_1defun_4generate_1label = MakeSymbol("shen.x.factorise-defun.generate-label")
var symshen_4type_1signature_1of_1reverse = MakeSymbol("shen.type-signature-of-reverse")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4synonyms_1help = MakeSymbol("shen.synonyms-help")
var symshen_4x_4factorise_1defun_4with_1labelled_1else = MakeSymbol("shen.x.factorise-defun.with-labelled-else")
var symshen_4_5simple__pattern_6 = MakeSymbol("shen.<simple_pattern>")
var symshen_4s_1prolog__clause = MakeSymbol("shen.s-prolog_clause")
var symshen_4catchpoint = MakeSymbol("shen.catchpoint")
var symshen_4s = MakeSymbol("shen.s")
var symcut = MakeSymbol("cut")
var symlist = MakeSymbol("list")
var symshen_4symbol_1code_2 = MakeSymbol("shen.symbol-code?")
var symshen_4tlhd = MakeSymbol("shen.tlhd")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var sym_3 = MakeSymbol("$")
var symshen_4type_1signature_1of_1occurrences = MakeSymbol("shen.type-signature-of-occurrences")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symvector = MakeSymbol("vector")
var symshen_4reduce = MakeSymbol("shen.reduce")
var symshen_4continuation__call = MakeSymbol("shen.continuation_call")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4x_4features_4cond_1expand = MakeSymbol("shen.x.features.cond-expand")
var symclose = MakeSymbol("close")
var symshen_4curry_1synonyms = MakeSymbol("shen.curry-synonyms")
var symshen_4update_1symbol_1table = MakeSymbol("shen.update-symbol-table")
var symshen_4ue_2 = MakeSymbol("shen.ue?")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4construct_1premiss_1literal = MakeSymbol("shen.construct-premiss-literal")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symand = MakeSymbol("and")
var symshen_4_5conclusion_6 = MakeSymbol("shen.<conclusion>")
var symshen_4type_1signature_1of_1_c = MakeSymbol("shen.type-signature-of-/")
var symshen_4nextbyte = MakeSymbol("shen.nextbyte")
var symshen_4yacc__cases = MakeSymbol("shen.yacc_cases")
var symshen_4_dinstalling_1kl_d = MakeSymbol("shen.*installing-kl*")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var symshen_4type_1signature_1of_1write_1to_1file = MakeSymbol("shen.type-signature-of-write-to-file")
var symshen_4typextable = MakeSymbol("shen.typextable")
var symget_1time = MakeSymbol("get-time")
var symshen_4x_4factorise_1defun_4inline_1mono_1labels = MakeSymbol("shen.x.factorise-defun.inline-mono-labels")
var symshen_4type_1signature_1of_1external = MakeSymbol("shen.type-signature-of-external")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symtuple_2 = MakeSymbol("tuple?")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4x_4factorise_1defun_4true_1branch = MakeSymbol("shen.x.factorise-defun.true-branch")
var symhead = MakeSymbol("head")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4syntax = MakeSymbol("shen.syntax")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4collect = MakeSymbol("shen.collect")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4grammar__symbol_2 = MakeSymbol("shen.grammar_symbol?")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symdefun = MakeSymbol("defun")
var symshen_4encode_1choices = MakeSymbol("shen.encode-choices")
var symshen_4x_4factorise_1defun_4bind_1repeating_1selectors = MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors")
var symshen_4_dvarcounter_d = MakeSymbol("shen.*varcounter*")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4_a_a_6 = MakeSymbol("shen.==>")
var symshen_4prolog_1error = MakeSymbol("shen.prolog-error")
var symshen_4rename = MakeSymbol("shen.rename")
var symhdv = MakeSymbol("hdv")
var symshen_4call__the__continuation = MakeSymbol("shen.call_the_continuation")
var symshen_4_5semicolon_1symbol_6 = MakeSymbol("shen.<semicolon-symbol>")
var symContextOut__1957 = MakeSymbol("ContextOut_1957")
var symeval_1kl = MakeSymbol("eval-kl")
var symgensym = MakeSymbol("gensym")
var symshen_4tests = MakeSymbol("shen.tests")
var syminferences = MakeSymbol("inferences")
var symshen_4demod_1rule = MakeSymbol("shen.demod-rule")
var symshen_4x_4factorise_1defun_4attach_1free_1variables = MakeSymbol("shen.x.factorise-defun.attach-free-variables")
var symshen_4add_1end = MakeSymbol("shen.add-end")
var sym_f_flet_1label = MakeSymbol("%%let-label")
var sym_dargv_d = MakeSymbol("*argv*")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_5singleunderline_6 = MakeSymbol("shen.<singleunderline>")
var symAssumptions__1957 = MakeSymbol("Assumptions_1957")
var symsave = MakeSymbol("save")
var symshen_4type_1signature_1of_1unprofile = MakeSymbol("shen.type-signature-of-unprofile")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symfwhen = MakeSymbol("fwhen")
var symshen_4insert__modes = MakeSymbol("shen.insert_modes")
var sym_1_1_6 = MakeSymbol("-->")
var symC = MakeSymbol("C")
var symvector_2 = MakeSymbol("vector?")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4toplevel__evaluate = MakeSymbol("shen.toplevel_evaluate")
var sym_1_6 = MakeSymbol("->")
var symreturn = MakeSymbol("return")
var symmode = MakeSymbol("mode")
var symshen_4strip_1pathname = MakeSymbol("shen.strip-pathname")
var symshen_4ebr = MakeSymbol("shen.ebr")
var symundefmacro = MakeSymbol("undefmacro")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symvariable_2 = MakeSymbol("variable?")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symshen_4type_1signature_1of_1enable_1type_1theory = MakeSymbol("shen.type-signature-of-enable-type-theory")
var symshen_4type_1signature_1of_1protect = MakeSymbol("shen.type-signature-of-protect")
var symshen_4type_1signature_1of_1_6 = MakeSymbol("shen.type-signature-of->")
var symshen_4_5num_6 = MakeSymbol("shen.<num>")
var symV = MakeSymbol("V")
var symParse__Y = MakeSymbol("Parse_Y")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4right_1rule = MakeSymbol("shen.right-rule")
var symshen_4x_4factorise_1defun_4factorise_1cond = MakeSymbol("shen.x.factorise-defun.factorise-cond")
var symshen_4variancy_1test = MakeSymbol("shen.variancy-test")
var symB = MakeSymbol("B")
var symshen_4remove_1nth = MakeSymbol("shen.remove-nth")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symsymbol = MakeSymbol("symbol")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4x_4factorise_1defun_4register_1selector_1handler = MakeSymbol("shen.x.factorise-defun.register-selector-handler")
var symshen_4type_1signature_1of_1inferences = MakeSymbol("shen.type-signature-of-inferences")
var symshen_4package_1macro = MakeSymbol("shen.package-macro")
var symshen_4numbyte_2 = MakeSymbol("shen.numbyte?")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var symshen_4type_1signature_1of_1boolean_2 = MakeSymbol("shen.type-signature-of-boolean?")
var symshen_4insert_1prolog_1variables_1help = MakeSymbol("shen.insert-prolog-variables-help")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symshen_4type_1signature_1of_1pos = MakeSymbol("shen.type-signature-of-pos")
var symshen_4compile__to__machine__code = MakeSymbol("shen.compile_to_machine_code")
var symshen_4_dcustom_1pattern_1reducer_d = MakeSymbol("shen.*custom-pattern-reducer*")
var symshen_4x_4features_4add = MakeSymbol("shen.x.features.add")
var symprofile = MakeSymbol("profile")
var symshen_4variable = MakeSymbol("shen.variable")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symboolean_2 = MakeSymbol("boolean?")
var symlanguage = MakeSymbol("language")
var symdefcc = MakeSymbol("defcc")
var symshen_4ue_1h_2 = MakeSymbol("shen.ue-h?")
var symshen_4complexity = MakeSymbol("shen.complexity")
var symshen_4_5postdigits_6 = MakeSymbol("shen.<postdigits>")
var symshen_4toplineread = MakeSymbol("shen.toplineread")
var symremove = MakeSymbol("remove")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4type_1signature_1of_1internal = MakeSymbol("shen.type-signature-of-internal")
var symS = MakeSymbol("S")
var symshen_4then = MakeSymbol("shen.then")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4type_1signature_1of_1symbol_2 = MakeSymbol("shen.type-signature-of-symbol?")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4_dprocess_1counter_d = MakeSymbol("shen.*process-counter*")
var symdatatype = MakeSymbol("datatype")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4type_1signature_1of_1shen_4insert = MakeSymbol("shen.type-signature-of-shen.insert")
var symshen_4tab = MakeSymbol("shen.tab")
var symshen_4prolog_1failure = MakeSymbol("shen.prolog-failure")
var symarity = MakeSymbol("arity")
var symshen_4lzy_a_a = MakeSymbol("shen.lzy==")
var symshen_4percent = MakeSymbol("shen.percent")
var symK = MakeSymbol("K")
var symstring_2 = MakeSymbol("string?")
var symintersection = MakeSymbol("intersection")
var symis = MakeSymbol("is")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4record_1source = MakeSymbol("shen.record-source")
var symshen_4function_1macro = MakeSymbol("shen.function-macro")
var symshen_4default_1rule = MakeSymbol("shen.default-rule")
var symfail = MakeSymbol("fail")
var symshen_4dereferencing = MakeSymbol("shen.dereferencing")
var symshen_4_5byte_6 = MakeSymbol("shen.<byte>")
var symshen_4safe_1multiply = MakeSymbol("shen.safe-multiply")
var symshen_4yacc = MakeSymbol("shen.yacc")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var sym_5_a = MakeSymbol("<=")
var symkill = MakeSymbol("kill")
var symshen_4read_1file_1as_1charlist = MakeSymbol("shen.read-file-as-charlist")
var symshen_4_5non_1ll_1rules_6 = MakeSymbol("shen.<non-ll-rules>")
var symshen_4f__error = MakeSymbol("shen.f_error")
var symappend = MakeSymbol("append")
var symunion = MakeSymbol("union")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var syminternal = MakeSymbol("internal")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symexception = MakeSymbol("exception")
var symshen_4left_1rule = MakeSymbol("shen.left-rule")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4_dmacroreg_d = MakeSymbol("shen.*macroreg*")
var symopen = MakeSymbol("open")
var symshen_4type_1signature_1of_1vector_1_6 = MakeSymbol("shen.type-signature-of-vector->")
var symshen_4type_1signature_1of_1write_1byte = MakeSymbol("shen.type-signature-of-write-byte")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4_5head_d_6 = MakeSymbol("shen.<head*>")
var symshen_4make_1key = MakeSymbol("shen.make-key")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symunput = MakeSymbol("unput")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4type_1signature_1of_1implementation = MakeSymbol("shen.type-signature-of-implementation")
var symshen_4succeeds_2 = MakeSymbol("shen.succeeds?")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symshen_4type_1signature_1of_1undefmacro = MakeSymbol("shen.type-signature-of-undefmacro")
var symshen_4type_1signature_1of_1nth = MakeSymbol("shen.type-signature-of-nth")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4dict = MakeSymbol("shen.dict")
var symcompile = MakeSymbol("compile")
var symsterror = MakeSymbol("sterror")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var sym_k = MakeSymbol(";")
var symshen_4_5literal_d_6 = MakeSymbol("shen.<literal*>")
var symshen_4eval_1without_1macros = MakeSymbol("shen.eval-without-macros")
var symit = MakeSymbol("it")
var symshen_4type_1signature_1of_1_6_a = MakeSymbol("shen.type-signature-of->=")
var sym_8s = MakeSymbol("@s")
var symaddress_1_6 = MakeSymbol("address->")
var symcall = MakeSymbol("call")
var symshen_4shen_1syntax_1error = MakeSymbol("shen.shen-syntax-error")
var symidentical = MakeSymbol("identical")
var symshen_4cl = MakeSymbol("shen.cl")
var symshen_4type_1signature_1of_1destroy = MakeSymbol("shen.type-signature-of-destroy")
var symshen_4type_1signature_1of_1empty_2 = MakeSymbol("shen.type-signature-of-empty?")
var symshen_4type_1signature_1of_1_5 = MakeSymbol("shen.type-signature-of-<")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4clauses_1to_1shen = MakeSymbol("shen.clauses-to-shen")
var symshen_4t_d_1defhh = MakeSymbol("shen.t*-defhh")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symerror = MakeSymbol("error")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4singleunderline_2 = MakeSymbol("shen.singleunderline?")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symoutput = MakeSymbol("output")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4intprolog_1help = MakeSymbol("shen.intprolog-help")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4sequent = MakeSymbol("shen.sequent")
var symR = MakeSymbol("R")
var symD = MakeSymbol("D")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4type_1signature_1of_1simple_1error = MakeSymbol("shen.type-signature-of-simple-error")
var symshen_4cn_1all = MakeSymbol("shen.cn-all")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4terminal_2 = MakeSymbol("shen.terminal?")
var symshen_4type_1signature_1of_1limit = MakeSymbol("shen.type-signature-of-limit")
var symshen_4compress_150 = MakeSymbol("shen.compress-50")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4t_d_1defh = MakeSymbol("shen.t*-defh")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4type_1signature_1of_1difference = MakeSymbol("shen.type-signature-of-difference")
var symshen_4type_1signature_1of_1_5e_6 = MakeSymbol("shen.type-signature-of-<e>")
var symmapcan = MakeSymbol("mapcan")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symspy = MakeSymbol("spy")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4udefs_d = MakeSymbol("shen.udefs*")
var symY = MakeSymbol("Y")
var symshen_4type_1signature_1of_1port = MakeSymbol("shen.type-signature-of-port")
var symshen_4type_1signature_1of_1bound_2 = MakeSymbol("shen.type-signature-of-bound?")
var symshen_4same__predicate_2 = MakeSymbol("shen.same_predicate?")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4reassemble = MakeSymbol("shen.reassemble")
var symtime = MakeSymbol("time")
var symshen_4type_1signature_1of_1arity = MakeSymbol("shen.type-signature-of-arity")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4line = MakeSymbol("shen.line")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4type_1signature_1of_1hdv = MakeSymbol("shen.type-signature-of-hdv")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4intprolog_1help_1help = MakeSymbol("shen.intprolog-help-help")
var symshen_4double = MakeSymbol("shen.double")
var symshen_4type_1signature_1of_1gensym = MakeSymbol("shen.type-signature-of-gensym")
var symshen_4typecheck_1and_1load = MakeSymbol("shen.typecheck-and-load")
var symshen_4chwild = MakeSymbol("shen.chwild")
var symshen_4type_1signature_1of_1it = MakeSymbol("shen.type-signature-of-it")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4after_1codepoint = MakeSymbol("shen.after-codepoint")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4t_d_1def = MakeSymbol("shen.t*-def")
var symcons_2 = MakeSymbol("cons?")
var symfail_1if = MakeSymbol("fail-if")
var symwarn = MakeSymbol("warn")
var symshen_4semantic_1completion_1warning = MakeSymbol("shen.semantic-completion-warning")
var sym_f_flabel = MakeSymbol("%%label")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symshen_4type_1signature_1of_1cn = MakeSymbol("shen.type-signature-of-cn")
var symshen_4terminator_2 = MakeSymbol("shen.terminator?")
var symshen_4alphanum_2 = MakeSymbol("shen.alphanum?")
var symbind = MakeSymbol("bind")
var symshen_4type_1signature_1of_1ps = MakeSymbol("shen.type-signature-of-ps")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symContext1__1957 = MakeSymbol("Context1_1957")
var sym_5 = MakeSymbol("<")
var symshen_4type_1signature_1of_1tlv = MakeSymbol("shen.type-signature-of-tlv")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4patthyps = MakeSymbol("shen.patthyps")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4curry_1type_1h = MakeSymbol("shen.curry-type-h")
var symshen_4_dsignedfuncs_d = MakeSymbol("shen.*signedfuncs*")
var symlineread = MakeSymbol("lineread")
var symshen_4type_1signature_1of_1number_2 = MakeSymbol("shen.type-signature-of-number?")
var symshen_4type_1signature_1of_1read_1from_1string = MakeSymbol("shen.type-signature-of-read-from-string")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symnot = MakeSymbol("not")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var symshen_4type_1signature_1of_1mapcan = MakeSymbol("shen.type-signature-of-mapcan")
var sym_h = MakeSymbol(":")
var syminput_7 = MakeSymbol("input+")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4double_1_6singles = MakeSymbol("shen.double->singles")
var symsnd = MakeSymbol("snd")
var symrun = MakeSymbol("run")
var symshen_4lisp_1or = MakeSymbol("shen.lisp-or")
var symT = MakeSymbol("T")
var symshen_4type_1signature_1of_1profile = MakeSymbol("shen.type-signature-of-profile")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4flatten = MakeSymbol("shen.flatten")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symshen_4interror = MakeSymbol("shen.interror")
var symshen_4type_1signature_1of_1_5_1vector = MakeSymbol("shen.type-signature-of-<-vector")
var symshen_4type_1signature_1of_1read_1file = MakeSymbol("shen.type-signature-of-read-file")
var symget = MakeSymbol("get")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symF = MakeSymbol("F")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4pre = MakeSymbol("shen.pre")
var symOut__1957 = MakeSymbol("Out_1957")
var symshen_4findpos = MakeSymbol("shen.findpos")
var sympos = MakeSymbol("pos")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var symintern = MakeSymbol("intern")
var symshen_4x_4factorise_1defun_4false_1branch = MakeSymbol("shen.x.factorise-defun.false-branch")
var sym_c_4 = MakeSymbol("/.")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4legitimate_1term_2 = MakeSymbol("shen.legitimate-term?")
var symin = MakeSymbol("in")
var symshen_4right_1_6left = MakeSymbol("shen.right->left")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symprotect = MakeSymbol("protect")
var symnumber_2 = MakeSymbol("number?")
var symshen_4get_1type = MakeSymbol("shen.get-type")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symshen_4_5predicate_d_6 = MakeSymbol("shen.<predicate*>")
var symshen_4else = MakeSymbol("shen.else")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4t_d_1hyps = MakeSymbol("shen.t*-hyps")
var symshen_4application__build = MakeSymbol("shen.application_build")
var symlimit = MakeSymbol("limit")
var symshen_4type_1signature_1of_1porters = MakeSymbol("shen.type-signature-of-porters")
var symshen_4sys_1error = MakeSymbol("shen.sys-error")
var symG = MakeSymbol("G")
var symoptimise = MakeSymbol("optimise")
var symshen_4th_d = MakeSymbol("shen.th*")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symParse__ = MakeSymbol("Parse_")
var symshen_4_dmaxcomplexity_d = MakeSymbol("shen.*maxcomplexity*")
var symdeclare = MakeSymbol("declare")
var symnull = MakeSymbol("null")
var symNPP = MakeSymbol("NPP")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var sym_e_e_e = MakeSymbol("&&&")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4linearise__X = MakeSymbol("shen.linearise_X")
var symstinput = MakeSymbol("stinput")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var sym_6_6 = MakeSymbol(">>")
var symshen_4type_1signature_1of_1occurs_1check = MakeSymbol("shen.type-signature-of-occurs-check")
var symshen_4post = MakeSymbol("shen.post")
var symshen_4remember = MakeSymbol("shen.remember")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symFreeze = MakeSymbol("Freeze")
var symshen_4add_1macro = MakeSymbol("shen.add-macro")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symlet = MakeSymbol("let")
var symboolean = MakeSymbol("boolean")
var symtc_2 = MakeSymbol("tc?")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symP = MakeSymbol("P")
var symshen_4pop = MakeSymbol("shen.pop")
var symshen_4ues = MakeSymbol("shen.ues")
var symshen_4catchpoint_b = MakeSymbol("shen.catchpoint!")
var symshen_4construct_1search_1literals = MakeSymbol("shen.construct-search-literals")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symlength = MakeSymbol("length")
var symshen_4type_1signature_1of_1tc = MakeSymbol("shen.type-signature-of-tc")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var symshen_4read_1file_1as_1Xlist_1help = MakeSymbol("shen.read-file-as-Xlist-help")
var symshen_4_5premises_6 = MakeSymbol("shen.<premises>")
var symshen_4type_1signature_1of_1string_1_6symbol = MakeSymbol("shen.type-signature-of-string->symbol")
var symshen_4type_1signature_1of_1thaw = MakeSymbol("shen.type-signature-of-thaw")
var symshen_4insert_1predicate = MakeSymbol("shen.insert-predicate")
var sym__ = MakeSymbol("_")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4variables = MakeSymbol("shen.variables")
var symshen_4space = MakeSymbol("shen.space")
var symempty_2 = MakeSymbol("empty?")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4type_1signature_1of_1string_1_6n = MakeSymbol("shen.type-signature-of-string->n")
var symshen_4x_4features_4_dfeatures_d = MakeSymbol("shen.x.features.*features*")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symshen_4aritycheck_1action = MakeSymbol("shen.aritycheck-action")
var symshen_4rule_1_6horn_1clause = MakeSymbol("shen.rule->horn-clause")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symRecord = MakeSymbol("Record")
var symshen_4type_1signature_1of_1 = MakeSymbol("shen.type-signature-of-")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symstream = MakeSymbol("stream")
var symshen_4make__mu__application = MakeSymbol("shen.make_mu_application")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4variable_1match = MakeSymbol("shen.variable-match")
var symshen_4type_1signature_1of_1profile_1results = MakeSymbol("shen.type-signature-of-profile-results")
var symshen_4stack = MakeSymbol("shen.stack")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4type_1signature_1of_1head = MakeSymbol("shen.type-signature-of-head")
var symshen_4type_1signature_1of_1intersection = MakeSymbol("shen.type-signature-of-intersection")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4aritycheck_1name = MakeSymbol("shen.aritycheck-name")
var sym_e_e = MakeSymbol("&&")
var symshen_4construct_1recursive_1search_1clause = MakeSymbol("shen.construct-recursive-search-clause")
var symshen_4ue = MakeSymbol("shen.ue")
var symshen_4prh = MakeSymbol("shen.prh")
var symshen_4jump__stream = MakeSymbol("shen.jump_stream")
var symshen_4x_4factorise_1defun_4factorise_1defun = MakeSymbol("shen.x.factorise-defun.factorise-defun")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4type_1signature_1of_1absvector_2 = MakeSymbol("shen.type-signature-of-absvector?")
var symshen_4result = MakeSymbol("shen.result")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4x_4factorise_1defun_4optimize_1selectors = MakeSymbol("shen.x.factorise-defun.optimize-selectors")
var symvalue = MakeSymbol("value")
var symshen_4assumetype = MakeSymbol("shen.assumetype")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4s_1prolog = MakeSymbol("shen.s-prolog")
var symshen_4place = MakeSymbol("shen.place")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symread_1file = MakeSymbol("read-file")
var symshen_4type_1signature_1of_1tuple_2 = MakeSymbol("shen.type-signature-of-tuple?")
var symshen_4datatype_1error = MakeSymbol("shen.datatype-error")
var symshen_4prolog_1_6shen = MakeSymbol("shen.prolog->shen")
var symshen_4call_1help = MakeSymbol("shen.call-help")
var symshen_4construct_1side_1literals = MakeSymbol("shen.construct-side-literals")
var symassoc = MakeSymbol("assoc")
var symshow_1help = MakeSymbol("show-help")
var sym_a_a = MakeSymbol("==")
var symYaccParse = MakeSymbol("YaccParse")
var symshen_4type_1signature_1of_1n_1_6string = MakeSymbol("shen.type-signature-of-n->string")
var symshen_4x_4factorise_1defun_4rebranch = MakeSymbol("shen.x.factorise-defun.rebranch")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symps = MakeSymbol("ps")
var sym_dos_d = MakeSymbol("*os*")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4next_150 = MakeSymbol("shen.next-50")
var symtlv = MakeSymbol("tlv")
var symshen_4bld_1prolog_1call = MakeSymbol("shen.bld-prolog-call")
var symshen_4type_1signature_1of_1include = MakeSymbol("shen.type-signature-of-include")
var symshen_4record_1it_1h = MakeSymbol("shen.record-it-h")
var symshen_4carriage_1return = MakeSymbol("shen.carriage-return")
var symshen_4sh_2 = MakeSymbol("shen.sh?")
var symbound_2 = MakeSymbol("bound?")
var symshen_4type_1signature_1of_1systemf = MakeSymbol("shen.type-signature-of-systemf")
var symshen_4mu = MakeSymbol("shen.mu")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symdefine = MakeSymbol("define")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symnl = MakeSymbol("nl")
var symshen_4write_1char_1and_1inc = MakeSymbol("shen.write-char-and-inc")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4type_1signature_1of_1error_1to_1string = MakeSymbol("shen.type-signature-of-error-to-string")
var symshen_4type_1signature_1of_1fst = MakeSymbol("shen.type-signature-of-fst")
var symshen_4type_1signature_1of_1language = MakeSymbol("shen.type-signature-of-language")
var symshen_4function_1abstraction = MakeSymbol("shen.function-abstraction")
var symshen_4cutpoint = MakeSymbol("shen.cutpoint")
var symshen_4lineread_1loop = MakeSymbol("shen.lineread-loop")
var symstoutput = MakeSymbol("stoutput")
var symhd = MakeSymbol("hd")
var symI = MakeSymbol("I")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4placeholders = MakeSymbol("shen.placeholders")
var symshen_4reduce__help = MakeSymbol("shen.reduce_help")
var symread = MakeSymbol("read")
var symshen_4cons__form = MakeSymbol("shen.cons_form")
var symshen_4_5term_d_6 = MakeSymbol("shen.<term*>")
var symshen_4copy_1vector_1stage_11 = MakeSymbol("shen.copy-vector-stage-1")
var sym_5e_6 = MakeSymbol("<e>")
var symwhere = MakeSymbol("where")
var sym_5_1vector = MakeSymbol("<-vector")
var symshen_4read_7 = MakeSymbol("shen.read+")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4continuation = MakeSymbol("shen.continuation")
var symshen_4pushnew = MakeSymbol("shen.pushnew")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symns2_1set = MakeSymbol("ns2-set")
var symshen_4free__variable__check = MakeSymbol("shen.free_variable_check")
var symJ = MakeSymbol("J")
var symshen_4split__cc__rule = MakeSymbol("shen.split_cc_rule")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symdestroy = MakeSymbol("destroy")
