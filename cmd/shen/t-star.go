package main

import . "github.com/tiancaiamao/shen-go/klambda"

var TStarMain = MakeNative(func(__e *ControlFlow) {
	tmp12217 := MakeNative(func(__e *ControlFlow) {
		V4258 := __e.Get(1)
		_ = V4258
		V4259 := __e.Get(2)
		_ = V4259
		tmp12218 := MakeNative(func(__e *ControlFlow) {
			Vs := __e.Get(1)
			_ = Vs
			tmp12219 := MakeNative(func(__e *ControlFlow) {
				A_d := __e.Get(1)
				_ = A_d
				tmp12220 := MakeNative(func(__e *ControlFlow) {
					Curried := __e.Get(1)
					_ = Curried
					tmp12221 := MakeNative(func(__e *ControlFlow) {
						V3660 := __e.Get(1)
						_ = V3660
						__e.Return(MakeNative(func(__e *ControlFlow) {
							L3661 := __e.Get(1)
							_ = L3661
							__e.Return(MakeNative(func(__e *ControlFlow) {
								K3662 := __e.Get(1)
								_ = K3662
								__e.Return(MakeNative(func(__e *ControlFlow) {
									C3663 := __e.Get(1)
									_ = C3663
									tmp12222 := MakeNative(func(__e *ControlFlow) {
										Out := __e.Get(1)
										_ = Out
										tmp12223 := Call(__e, PrimNS2Value(symshen_4incinfs))

										_ = tmp12223

										tmp12224 := Call(__e, PrimNS2Value(symshen_4deref), Vs, V3660)

										tmp12225 := Call(__e, PrimNS2Value(symreceive), tmp12224)

										tmp12226 := Call(__e, PrimNS2Value(symshen_4deref), A_d, V3660)

										tmp12227 := Call(__e, PrimNS2Value(symreceive), tmp12226)

										tmp12228 := MakeNative(func(__e *ControlFlow) {
											tmp12229 := Call(__e, PrimNS2Value(symshen_4deref), Curried, V3660)

											tmp12230 := Call(__e, PrimNS2Value(symreceive), tmp12229)

											tmp12231 := MakeNative(func(__e *ControlFlow) {
												__e.TailApply(PrimNS2Value(symreturn), Out, V3660, L3661, K3662, C3663)
												return
											}, 0)

											__e.TailApply(PrimNS2Value(symshen_4toplevel_1forms), tmp12230, Out, V3660, L3661, K3662, tmp12231)
											return

										}, 0)

										tmp12232 := Call(__e, PrimNS2Value(symshen_4insert_1prolog_1variables), tmp12225, tmp12227, Out, V3660, L3661, K3662, tmp12228)

										__e.TailApply(PrimNS2Value(symshen_4gc), V3660, tmp12232)
										return

									}, 1)

									tmp12233 := Call(__e, PrimNS2Value(symshen_4newpv), V3660)

									__e.TailApply(tmp12222, tmp12233)
									return

								}, 1))
								return
							}, 1))
							return
						}, 1))
						return
					}, 1)

					tmp12234 := Call(__e, PrimNS2Value(symshen_4reset_1prolog_1vector))

					tmp12235 := Call(__e, tmp12221, tmp12234)

					tmp12236 := Call(__e, PrimNS2Value(symvector), MakeNumber(0))

					tmp12237 := Call(__e, PrimNS2Value(sym_8v), MakeNumber(0), tmp12236)

					tmp12238 := Call(__e, PrimNS2Value(sym_8v), True, tmp12237)

					tmp12239 := Call(__e, tmp12235, tmp12238)

					tmp12240 := Call(__e, tmp12239, MakeNumber(0))

					tmp12241 := MakeNative(func(__e *ControlFlow) {
						__e.Return(True)
						return
					}, 0)

					__e.TailApply(tmp12240, tmp12241)
					return

				}, 1)

				tmp12242 := Call(__e, PrimNS2Value(symshen_4curry), V4258)

				__e.TailApply(tmp12220, tmp12242)
				return

			}, 1)

			tmp12243 := Call(__e, PrimNS2Value(symshen_4rectify_1type), V4259)

			__e.TailApply(tmp12219, tmp12243)
			return

		}, 1)

		tmp12244 := Call(__e, PrimNS2Value(symshen_4extract_1vars), V4259)

		__e.TailApply(tmp12218, tmp12244)
		return

	}, 2)

	tmp12245 := Call(__e, PrimNS2Value(symdef), symshen_4typecheck, tmp12217)

	_ = tmp12245

	tmp12246 := MakeNative(func(__e *ControlFlow) {
		V4260 := __e.Get(1)
		_ = V4260
		V4261 := __e.Get(2)
		_ = V4261
		V4262 := __e.Get(3)
		_ = V4262
		V4263 := __e.Get(4)
		_ = V4263
		V4264 := __e.Get(5)
		_ = V4264
		V4265 := __e.Get(6)
		_ = V4265
		V4266 := __e.Get(7)
		_ = V4266
		tmp12247 := MakeNative(func(__e *ControlFlow) {
			C3672 := __e.Get(1)
			_ = C3672
			tmp12265 := PrimEqual(C3672, False)

			if True == tmp12265 {
				tmp12264 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4264)

				if True == tmp12264 {
					tmp12250 := MakeNative(func(__e *ControlFlow) {
						Tm3674 := __e.Get(1)
						_ = Tm3674
						tmp12262 := PrimIsPair(Tm3674)

						if True == tmp12262 {
							tmp12252 := MakeNative(func(__e *ControlFlow) {
								V := __e.Get(1)
								_ = V
								tmp12253 := MakeNative(func(__e *ControlFlow) {
									Vs := __e.Get(1)
									_ = Vs
									tmp12254 := MakeNative(func(__e *ControlFlow) {
										X := __e.Get(1)
										_ = X
										tmp12255 := Call(__e, PrimNS2Value(symshen_4incinfs))

										_ = tmp12255

										tmp12256 := Call(__e, PrimNS2Value(symshen_4deref), X, V4263)

										tmp12257 := Call(__e, PrimNS2Value(symsubst), tmp12256, V, V4261)

										tmp12258 := Call(__e, PrimNS2Value(symshen_4insert_1prolog_1variables), Vs, tmp12257, V4262, V4263, V4264, V4265, V4266)

										__e.TailApply(PrimNS2Value(symshen_4gc), V4263, tmp12258)
										return

									}, 1)

									tmp12259 := Call(__e, PrimNS2Value(symshen_4newpv), V4263)

									__e.TailApply(tmp12254, tmp12259)
									return

								}, 1)

								tmp12260 := PrimTail(Tm3674)

								__e.TailApply(tmp12253, tmp12260)
								return

							}, 1)

							tmp12261 := PrimHead(Tm3674)

							__e.TailApply(tmp12252, tmp12261)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp12263 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4260, V4263)

					__e.TailApply(tmp12250, tmp12263)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C3672)
				return
			}

		}, 1)

		tmp12273 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4264)

		var ifres12266 Obj

		if True == tmp12273 {
			tmp12267 := MakeNative(func(__e *ControlFlow) {
				Tm3673 := __e.Get(1)
				_ = Tm3673
				tmp12270 := PrimEqual(Tm3673, Nil)

				if True == tmp12270 {
					tmp12269 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp12269

					__e.TailApply(PrimNS2Value(symis_b), V4261, V4262, V4263, V4264, V4265, V4266)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp12271 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4260, V4263)

			tmp12272 := Call(__e, tmp12267, tmp12271)

			ifres12266 = tmp12272

		} else {
			ifres12266 = False

		}

		__e.TailApply(tmp12247, ifres12266)
		return

	}, 7)

	tmp12274 := Call(__e, PrimNS2Value(symdef), symshen_4insert_1prolog_1variables, tmp12246)

	_ = tmp12274

	tmp12275 := MakeNative(func(__e *ControlFlow) {
		V4267 := __e.Get(1)
		_ = V4267
		V4268 := __e.Get(2)
		_ = V4268
		V4269 := __e.Get(3)
		_ = V4269
		V4270 := __e.Get(4)
		_ = V4270
		V4271 := __e.Get(5)
		_ = V4271
		V4272 := __e.Get(6)
		_ = V4272
		tmp12276 := MakeNative(func(__e *ControlFlow) {
			K3677 := __e.Get(1)
			_ = K3677
			tmp12277 := MakeNative(func(__e *ControlFlow) {
				C3681 := __e.Get(1)
				_ = C3681
				tmp12290 := PrimEqual(C3681, False)

				if True == tmp12290 {
					tmp12279 := MakeNative(func(__e *ControlFlow) {
						C3685 := __e.Get(1)
						_ = C3685
						tmp12281 := PrimEqual(C3685, False)

						if True == tmp12281 {
							__e.TailApply(PrimNS2Value(symshen_4unlock), V4270, K3677)
							return
						} else {
							__e.Return(C3685)
							return
						}

					}, 1)

					tmp12289 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4270)

					var ifres12282 Obj

					if True == tmp12289 {
						tmp12283 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp12283

						tmp12284 := PrimIntern(MakeString(":"))

						tmp12285 := PrimCons(V4268, Nil)

						tmp12286 := PrimCons(tmp12284, tmp12285)

						tmp12287 := PrimCons(V4267, tmp12286)

						tmp12288 := Call(__e, PrimNS2Value(symshen_4system_1S), tmp12287, Nil, V4269, V4270, K3677, V4272)

						ifres12282 = tmp12288

					} else {
						ifres12282 = False

					}

					__e.TailApply(tmp12279, ifres12282)
					return

				} else {
					__e.Return(C3681)
					return
				}

			}, 1)

			tmp12317 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4270)

			var ifres12291 Obj

			if True == tmp12317 {
				tmp12292 := MakeNative(func(__e *ControlFlow) {
					Tm3682 := __e.Get(1)
					_ = Tm3682
					tmp12314 := PrimIsPair(Tm3682)

					if True == tmp12314 {
						tmp12294 := MakeNative(func(__e *ControlFlow) {
							Tm3683 := __e.Get(1)
							_ = Tm3683
							tmp12311 := PrimEqual(Tm3683, symdefine)

							if True == tmp12311 {
								tmp12296 := MakeNative(func(__e *ControlFlow) {
									Tm3684 := __e.Get(1)
									_ = Tm3684
									tmp12308 := PrimIsPair(Tm3684)

									if True == tmp12308 {
										tmp12298 := MakeNative(func(__e *ControlFlow) {
											F := __e.Get(1)
											_ = F
											tmp12299 := MakeNative(func(__e *ControlFlow) {
												X := __e.Get(1)
												_ = X
												tmp12300 := Call(__e, PrimNS2Value(symshen_4incinfs))

												_ = tmp12300

												tmp12301 := MakeNative(func(__e *ControlFlow) {
													tmp12302 := PrimNS3Value(symshen_4_dspy_d)

													tmp12303 := MakeNative(func(__e *ControlFlow) {
														tmp12304 := PrimCons(F, X)

														tmp12305 := PrimCons(symdefine, tmp12304)

														__e.TailApply(PrimNS2Value(symshen_4t_d), tmp12305, V4268, V4269, V4270, K3677, V4272)
														return

													}, 0)

													__e.TailApply(PrimNS2Value(symshen_4signal_1def), tmp12302, F, V4269, V4270, K3677, tmp12303)
													return

												}, 0)

												__e.TailApply(PrimNS2Value(symshen_4cut), V4269, V4270, K3677, tmp12301)
												return

											}, 1)

											tmp12306 := PrimTail(Tm3684)

											__e.TailApply(tmp12299, tmp12306)
											return

										}, 1)

										tmp12307 := PrimHead(Tm3684)

										__e.TailApply(tmp12298, tmp12307)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp12309 := PrimTail(Tm3682)

								tmp12310 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12309, V4269)

								__e.TailApply(tmp12296, tmp12310)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp12312 := PrimHead(Tm3682)

						tmp12313 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12312, V4269)

						__e.TailApply(tmp12294, tmp12313)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp12315 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4267, V4269)

				tmp12316 := Call(__e, tmp12292, tmp12315)

				ifres12291 = tmp12316

			} else {
				ifres12291 = False

			}

			__e.TailApply(tmp12277, ifres12291)
			return

		}, 1)

		tmp12318 := PrimNumberAdd(V4271, MakeNumber(1))

		__e.TailApply(tmp12276, tmp12318)
		return

	}, 6)

	tmp12319 := Call(__e, PrimNS2Value(symdef), symshen_4toplevel_1forms, tmp12275)

	_ = tmp12319

	tmp12320 := MakeNative(func(__e *ControlFlow) {
		V4273 := __e.Get(1)
		_ = V4273
		V4274 := __e.Get(2)
		_ = V4274
		V4275 := __e.Get(3)
		_ = V4275
		V4276 := __e.Get(4)
		_ = V4276
		V4277 := __e.Get(5)
		_ = V4277
		V4278 := __e.Get(6)
		_ = V4278
		tmp12321 := MakeNative(func(__e *ControlFlow) {
			C3692 := __e.Get(1)
			_ = C3692
			tmp12338 := PrimEqual(C3692, False)

			if True == tmp12338 {
				tmp12337 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4276)

				if True == tmp12337 {
					tmp12324 := MakeNative(func(__e *ControlFlow) {
						Tm3694 := __e.Get(1)
						_ = Tm3694
						tmp12335 := PrimEqual(Tm3694, True)

						if True == tmp12335 {
							tmp12326 := MakeNative(func(__e *ControlFlow) {
								ShowF := __e.Get(1)
								_ = ShowF
								tmp12327 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp12327

								tmp12328 := Call(__e, PrimNS2Value(symshen_4deref), V4274, V4275)

								tmp12329 := Call(__e, PrimNS2Value(symshen_4app), tmp12328, MakeString(")\n"), symshen_4a)

								tmp12330 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp12329)

								tmp12331 := Call(__e, PrimNS2Value(symstoutput))

								tmp12332 := Call(__e, PrimNS2Value(sympr), tmp12330, tmp12331)

								tmp12333 := Call(__e, PrimNS2Value(symis), ShowF, tmp12332, V4275, V4276, V4277, V4278)

								__e.TailApply(PrimNS2Value(symshen_4gc), V4275, tmp12333)
								return

							}, 1)

							tmp12334 := Call(__e, PrimNS2Value(symshen_4newpv), V4275)

							__e.TailApply(tmp12326, tmp12334)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp12336 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4273, V4275)

					__e.TailApply(tmp12324, tmp12336)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C3692)
				return
			}

		}, 1)

		tmp12346 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4276)

		var ifres12339 Obj

		if True == tmp12346 {
			tmp12340 := MakeNative(func(__e *ControlFlow) {
				Tm3693 := __e.Get(1)
				_ = Tm3693
				tmp12343 := PrimEqual(Tm3693, False)

				if True == tmp12343 {
					tmp12342 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp12342

					__e.TailApply(PrimNS2Value(symthaw), V4278)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp12344 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4273, V4275)

			tmp12345 := Call(__e, tmp12340, tmp12344)

			ifres12339 = tmp12345

		} else {
			ifres12339 = False

		}

		__e.TailApply(tmp12321, ifres12339)
		return

	}, 6)

	tmp12347 := Call(__e, PrimNS2Value(symdef), symshen_4signal_1def, tmp12320)

	_ = tmp12347

	tmp12348 := MakeNative(func(__e *ControlFlow) {
		V4279 := __e.Get(1)
		_ = V4279
		tmp12349 := Call(__e, PrimNS2Value(symshen_4curry_1type), V4279)

		__e.TailApply(PrimNS2Value(symshen_4demodulate), tmp12349)
		return

	}, 1)

	tmp12350 := Call(__e, PrimNS2Value(symdef), symshen_4rectify_1type, tmp12348)

	_ = tmp12350

	tmp12351 := MakeNative(func(__e *ControlFlow) {
		V4280 := __e.Get(1)
		_ = V4280
		tmp12352 := MakeNative(func(__e *ControlFlow) {
			tmp12353 := MakeNative(func(__e *ControlFlow) {
				Demod := __e.Get(1)
				_ = Demod
				tmp12355 := PrimEqual(Demod, V4280)

				if True == tmp12355 {
					__e.Return(V4280)
					return
				} else {
					__e.TailApply(PrimNS2Value(symshen_4demodulate), Demod)
					return
				}

			}, 1)

			tmp12356 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimNS2Value(symshen_4demod), Y)
				return
			}, 1)

			tmp12357 := Call(__e, PrimNS2Value(symshen_4walk), tmp12356, V4280)

			__e.TailApply(tmp12353, tmp12357)
			return

		}, 0)

		tmp12358 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V4280)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp12352, tmp12358)
		return

	}, 1)

	tmp12359 := Call(__e, PrimNS2Value(symdef), symshen_4demodulate, tmp12351)

	_ = tmp12359

	tmp12360 := MakeNative(func(__e *ControlFlow) {
		V4281 := __e.Get(1)
		_ = V4281
		tmp12473 := PrimIsPair(V4281)

		var ifres12446 Obj

		if True == tmp12473 {
			tmp12471 := PrimTail(V4281)

			tmp12472 := PrimIsPair(tmp12471)

			var ifres12448 Obj

			if True == tmp12472 {
				tmp12468 := PrimTail(V4281)

				tmp12469 := PrimHead(tmp12468)

				tmp12470 := PrimEqual(sym_1_1_6, tmp12469)

				var ifres12450 Obj

				if True == tmp12470 {
					tmp12465 := PrimTail(V4281)

					tmp12466 := PrimTail(tmp12465)

					tmp12467 := PrimIsPair(tmp12466)

					var ifres12452 Obj

					if True == tmp12467 {
						tmp12461 := PrimTail(V4281)

						tmp12462 := PrimTail(tmp12461)

						tmp12463 := PrimTail(tmp12462)

						tmp12464 := PrimIsPair(tmp12463)

						var ifres12454 Obj

						if True == tmp12464 {
							tmp12456 := PrimTail(V4281)

							tmp12457 := PrimTail(tmp12456)

							tmp12458 := PrimTail(tmp12457)

							tmp12459 := PrimHead(tmp12458)

							tmp12460 := PrimEqual(sym_1_1_6, tmp12459)

							var ifres12455 Obj

							if True == tmp12460 {
								ifres12455 = True

							} else {
								ifres12455 = False

							}

							ifres12454 = ifres12455

						} else {
							ifres12454 = False

						}

						var ifres12453 Obj

						if True == ifres12454 {
							ifres12453 = True

						} else {
							ifres12453 = False

						}

						ifres12452 = ifres12453

					} else {
						ifres12452 = False

					}

					var ifres12451 Obj

					if True == ifres12452 {
						ifres12451 = True

					} else {
						ifres12451 = False

					}

					ifres12450 = ifres12451

				} else {
					ifres12450 = False

				}

				var ifres12449 Obj

				if True == ifres12450 {
					ifres12449 = True

				} else {
					ifres12449 = False

				}

				ifres12448 = ifres12449

			} else {
				ifres12448 = False

			}

			var ifres12447 Obj

			if True == ifres12448 {
				ifres12447 = True

			} else {
				ifres12447 = False

			}

			ifres12446 = ifres12447

		} else {
			ifres12446 = False

		}

		if True == ifres12446 {
			tmp12440 := PrimHead(V4281)

			tmp12441 := PrimTail(V4281)

			tmp12442 := PrimTail(tmp12441)

			tmp12443 := PrimCons(tmp12442, Nil)

			tmp12444 := PrimCons(sym_1_1_6, tmp12443)

			tmp12445 := PrimCons(tmp12440, tmp12444)

			__e.TailApply(PrimNS2Value(symshen_4curry_1type), tmp12445)
			return

		} else {
			tmp12439 := PrimIsPair(V4281)

			var ifres12419 Obj

			if True == tmp12439 {
				tmp12437 := PrimTail(V4281)

				tmp12438 := PrimIsPair(tmp12437)

				var ifres12421 Obj

				if True == tmp12438 {
					tmp12434 := PrimTail(V4281)

					tmp12435 := PrimHead(tmp12434)

					tmp12436 := PrimEqual(sym_a_a_6, tmp12435)

					var ifres12423 Obj

					if True == tmp12436 {
						tmp12431 := PrimTail(V4281)

						tmp12432 := PrimTail(tmp12431)

						tmp12433 := PrimIsPair(tmp12432)

						var ifres12425 Obj

						if True == tmp12433 {
							tmp12427 := PrimTail(V4281)

							tmp12428 := PrimTail(tmp12427)

							tmp12429 := PrimTail(tmp12428)

							tmp12430 := PrimEqual(Nil, tmp12429)

							var ifres12426 Obj

							if True == tmp12430 {
								ifres12426 = True

							} else {
								ifres12426 = False

							}

							ifres12425 = ifres12426

						} else {
							ifres12425 = False

						}

						var ifres12424 Obj

						if True == ifres12425 {
							ifres12424 = True

						} else {
							ifres12424 = False

						}

						ifres12423 = ifres12424

					} else {
						ifres12423 = False

					}

					var ifres12422 Obj

					if True == ifres12423 {
						ifres12422 = True

					} else {
						ifres12422 = False

					}

					ifres12421 = ifres12422

				} else {
					ifres12421 = False

				}

				var ifres12420 Obj

				if True == ifres12421 {
					ifres12420 = True

				} else {
					ifres12420 = False

				}

				ifres12419 = ifres12420

			} else {
				ifres12419 = False

			}

			if True == ifres12419 {
				tmp12401 := PrimHead(V4281)

				tmp12402 := Call(__e, PrimNS2Value(symprotect), symA)

				tmp12403 := PrimCons(tmp12402, Nil)

				tmp12404 := PrimCons(sym_d, tmp12403)

				tmp12405 := PrimCons(tmp12401, tmp12404)

				tmp12406 := PrimCons(symboolean, Nil)

				tmp12407 := PrimCons(symvector, tmp12406)

				tmp12408 := PrimHead(V4281)

				tmp12409 := PrimTail(V4281)

				tmp12410 := PrimTail(tmp12409)

				tmp12411 := PrimCons(sym_d, tmp12410)

				tmp12412 := PrimCons(tmp12408, tmp12411)

				tmp12413 := PrimCons(tmp12412, Nil)

				tmp12414 := PrimCons(sym_1_1_6, tmp12413)

				tmp12415 := PrimCons(tmp12407, tmp12414)

				tmp12416 := PrimCons(tmp12415, Nil)

				tmp12417 := PrimCons(sym_1_1_6, tmp12416)

				tmp12418 := PrimCons(tmp12405, tmp12417)

				__e.TailApply(PrimNS2Value(symshen_4curry_1type), tmp12418)
				return

			} else {
				tmp12400 := PrimIsPair(V4281)

				var ifres12373 Obj

				if True == tmp12400 {
					tmp12398 := PrimTail(V4281)

					tmp12399 := PrimIsPair(tmp12398)

					var ifres12375 Obj

					if True == tmp12399 {
						tmp12395 := PrimTail(V4281)

						tmp12396 := PrimHead(tmp12395)

						tmp12397 := PrimEqual(sym_d, tmp12396)

						var ifres12377 Obj

						if True == tmp12397 {
							tmp12392 := PrimTail(V4281)

							tmp12393 := PrimTail(tmp12392)

							tmp12394 := PrimIsPair(tmp12393)

							var ifres12379 Obj

							if True == tmp12394 {
								tmp12388 := PrimTail(V4281)

								tmp12389 := PrimTail(tmp12388)

								tmp12390 := PrimTail(tmp12389)

								tmp12391 := PrimIsPair(tmp12390)

								var ifres12381 Obj

								if True == tmp12391 {
									tmp12383 := PrimTail(V4281)

									tmp12384 := PrimTail(tmp12383)

									tmp12385 := PrimTail(tmp12384)

									tmp12386 := PrimHead(tmp12385)

									tmp12387 := PrimEqual(sym_d, tmp12386)

									var ifres12382 Obj

									if True == tmp12387 {
										ifres12382 = True

									} else {
										ifres12382 = False

									}

									ifres12381 = ifres12382

								} else {
									ifres12381 = False

								}

								var ifres12380 Obj

								if True == ifres12381 {
									ifres12380 = True

								} else {
									ifres12380 = False

								}

								ifres12379 = ifres12380

							} else {
								ifres12379 = False

							}

							var ifres12378 Obj

							if True == ifres12379 {
								ifres12378 = True

							} else {
								ifres12378 = False

							}

							ifres12377 = ifres12378

						} else {
							ifres12377 = False

						}

						var ifres12376 Obj

						if True == ifres12377 {
							ifres12376 = True

						} else {
							ifres12376 = False

						}

						ifres12375 = ifres12376

					} else {
						ifres12375 = False

					}

					var ifres12374 Obj

					if True == ifres12375 {
						ifres12374 = True

					} else {
						ifres12374 = False

					}

					ifres12373 = ifres12374

				} else {
					ifres12373 = False

				}

				if True == ifres12373 {
					tmp12367 := PrimHead(V4281)

					tmp12368 := PrimTail(V4281)

					tmp12369 := PrimTail(tmp12368)

					tmp12370 := PrimCons(tmp12369, Nil)

					tmp12371 := PrimCons(sym_d, tmp12370)

					tmp12372 := PrimCons(tmp12367, tmp12371)

					__e.TailApply(PrimNS2Value(symshen_4curry_1type), tmp12372)
					return

				} else {
					tmp12366 := PrimIsPair(V4281)

					if True == tmp12366 {
						tmp12365 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(PrimNS2Value(symshen_4curry_1type), Z)
							return
						}, 1)

						__e.TailApply(PrimNS2Value(symmap), tmp12365, V4281)
						return

					} else {
						__e.Return(V4281)
						return
					}

				}

			}

		}

	}, 1)

	tmp12474 := Call(__e, PrimNS2Value(symdef), symshen_4curry_1type, tmp12360)

	_ = tmp12474

	tmp12475 := MakeNative(func(__e *ControlFlow) {
		V4282 := __e.Get(1)
		_ = V4282
		tmp12593 := PrimIsPair(V4282)

		var ifres12585 Obj

		if True == tmp12593 {
			tmp12591 := PrimHead(V4282)

			tmp12592 := PrimEqual(symdefine, tmp12591)

			var ifres12587 Obj

			if True == tmp12592 {
				tmp12589 := PrimTail(V4282)

				tmp12590 := PrimIsPair(tmp12589)

				var ifres12588 Obj

				if True == tmp12590 {
					ifres12588 = True

				} else {
					ifres12588 = False

				}

				ifres12587 = ifres12588

			} else {
				ifres12587 = False

			}

			var ifres12586 Obj

			if True == ifres12587 {
				ifres12586 = True

			} else {
				ifres12586 = False

			}

			ifres12585 = ifres12586

		} else {
			ifres12585 = False

		}

		if True == ifres12585 {
			__e.Return(V4282)
			return
		} else {
			tmp12584 := PrimIsPair(V4282)

			var ifres12565 Obj

			if True == tmp12584 {
				tmp12582 := PrimHead(V4282)

				tmp12583 := PrimEqual(symtype, tmp12582)

				var ifres12567 Obj

				if True == tmp12583 {
					tmp12580 := PrimTail(V4282)

					tmp12581 := PrimIsPair(tmp12580)

					var ifres12569 Obj

					if True == tmp12581 {
						tmp12577 := PrimTail(V4282)

						tmp12578 := PrimTail(tmp12577)

						tmp12579 := PrimIsPair(tmp12578)

						var ifres12571 Obj

						if True == tmp12579 {
							tmp12573 := PrimTail(V4282)

							tmp12574 := PrimTail(tmp12573)

							tmp12575 := PrimTail(tmp12574)

							tmp12576 := PrimEqual(Nil, tmp12575)

							var ifres12572 Obj

							if True == tmp12576 {
								ifres12572 = True

							} else {
								ifres12572 = False

							}

							ifres12571 = ifres12572

						} else {
							ifres12571 = False

						}

						var ifres12570 Obj

						if True == ifres12571 {
							ifres12570 = True

						} else {
							ifres12570 = False

						}

						ifres12569 = ifres12570

					} else {
						ifres12569 = False

					}

					var ifres12568 Obj

					if True == ifres12569 {
						ifres12568 = True

					} else {
						ifres12568 = False

					}

					ifres12567 = ifres12568

				} else {
					ifres12567 = False

				}

				var ifres12566 Obj

				if True == ifres12567 {
					ifres12566 = True

				} else {
					ifres12566 = False

				}

				ifres12565 = ifres12566

			} else {
				ifres12565 = False

			}

			if True == ifres12565 {
				tmp12559 := PrimTail(V4282)

				tmp12560 := PrimHead(tmp12559)

				tmp12561 := Call(__e, PrimNS2Value(symshen_4curry), tmp12560)

				tmp12562 := PrimTail(V4282)

				tmp12563 := PrimTail(tmp12562)

				tmp12564 := PrimCons(tmp12561, tmp12563)

				__e.Return(PrimCons(symtype, tmp12564))
				return

			} else {
				tmp12558 := PrimIsPair(V4282)

				var ifres12539 Obj

				if True == tmp12558 {
					tmp12556 := PrimHead(V4282)

					tmp12557 := PrimEqual(syminput_7, tmp12556)

					var ifres12541 Obj

					if True == tmp12557 {
						tmp12554 := PrimTail(V4282)

						tmp12555 := PrimIsPair(tmp12554)

						var ifres12543 Obj

						if True == tmp12555 {
							tmp12551 := PrimTail(V4282)

							tmp12552 := PrimTail(tmp12551)

							tmp12553 := PrimIsPair(tmp12552)

							var ifres12545 Obj

							if True == tmp12553 {
								tmp12547 := PrimTail(V4282)

								tmp12548 := PrimTail(tmp12547)

								tmp12549 := PrimTail(tmp12548)

								tmp12550 := PrimEqual(Nil, tmp12549)

								var ifres12546 Obj

								if True == tmp12550 {
									ifres12546 = True

								} else {
									ifres12546 = False

								}

								ifres12545 = ifres12546

							} else {
								ifres12545 = False

							}

							var ifres12544 Obj

							if True == ifres12545 {
								ifres12544 = True

							} else {
								ifres12544 = False

							}

							ifres12543 = ifres12544

						} else {
							ifres12543 = False

						}

						var ifres12542 Obj

						if True == ifres12543 {
							ifres12542 = True

						} else {
							ifres12542 = False

						}

						ifres12541 = ifres12542

					} else {
						ifres12541 = False

					}

					var ifres12540 Obj

					if True == ifres12541 {
						ifres12540 = True

					} else {
						ifres12540 = False

					}

					ifres12539 = ifres12540

				} else {
					ifres12539 = False

				}

				if True == ifres12539 {
					tmp12531 := PrimTail(V4282)

					tmp12532 := PrimHead(tmp12531)

					tmp12533 := PrimTail(V4282)

					tmp12534 := PrimTail(tmp12533)

					tmp12535 := PrimHead(tmp12534)

					tmp12536 := Call(__e, PrimNS2Value(symshen_4curry), tmp12535)

					tmp12537 := PrimCons(tmp12536, Nil)

					tmp12538 := PrimCons(tmp12532, tmp12537)

					__e.Return(PrimCons(syminput_7, tmp12538))
					return

				} else {
					tmp12530 := PrimIsPair(V4282)

					var ifres12526 Obj

					if True == tmp12530 {
						tmp12528 := PrimHead(V4282)

						tmp12529 := Call(__e, PrimNS2Value(symshen_4special_2), tmp12528)

						var ifres12527 Obj

						if True == tmp12529 {
							ifres12527 = True

						} else {
							ifres12527 = False

						}

						ifres12526 = ifres12527

					} else {
						ifres12526 = False

					}

					if True == ifres12526 {
						tmp12522 := PrimHead(V4282)

						tmp12523 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							__e.TailApply(PrimNS2Value(symshen_4curry), Y)
							return
						}, 1)

						tmp12524 := PrimTail(V4282)

						tmp12525 := Call(__e, PrimNS2Value(symmap), tmp12523, tmp12524)

						__e.Return(PrimCons(tmp12522, tmp12525))
						return

					} else {
						tmp12521 := PrimIsPair(V4282)

						var ifres12517 Obj

						if True == tmp12521 {
							tmp12519 := PrimHead(V4282)

							tmp12520 := Call(__e, PrimNS2Value(symshen_4extraspecial_2), tmp12519)

							var ifres12518 Obj

							if True == tmp12520 {
								ifres12518 = True

							} else {
								ifres12518 = False

							}

							ifres12517 = ifres12518

						} else {
							ifres12517 = False

						}

						if True == ifres12517 {
							__e.Return(V4282)
							return
						} else {
							tmp12516 := PrimIsPair(V4282)

							var ifres12507 Obj

							if True == tmp12516 {
								tmp12514 := PrimTail(V4282)

								tmp12515 := PrimIsPair(tmp12514)

								var ifres12509 Obj

								if True == tmp12515 {
									tmp12511 := PrimTail(V4282)

									tmp12512 := PrimTail(tmp12511)

									tmp12513 := PrimIsPair(tmp12512)

									var ifres12510 Obj

									if True == tmp12513 {
										ifres12510 = True

									} else {
										ifres12510 = False

									}

									ifres12509 = ifres12510

								} else {
									ifres12509 = False

								}

								var ifres12508 Obj

								if True == ifres12509 {
									ifres12508 = True

								} else {
									ifres12508 = False

								}

								ifres12507 = ifres12508

							} else {
								ifres12507 = False

							}

							if True == ifres12507 {
								tmp12499 := PrimHead(V4282)

								tmp12500 := PrimTail(V4282)

								tmp12501 := PrimHead(tmp12500)

								tmp12502 := PrimCons(tmp12501, Nil)

								tmp12503 := PrimCons(tmp12499, tmp12502)

								tmp12504 := PrimTail(V4282)

								tmp12505 := PrimTail(tmp12504)

								tmp12506 := PrimCons(tmp12503, tmp12505)

								__e.TailApply(PrimNS2Value(symshen_4curry), tmp12506)
								return

							} else {
								tmp12498 := PrimIsPair(V4282)

								var ifres12489 Obj

								if True == tmp12498 {
									tmp12496 := PrimTail(V4282)

									tmp12497 := PrimIsPair(tmp12496)

									var ifres12491 Obj

									if True == tmp12497 {
										tmp12493 := PrimTail(V4282)

										tmp12494 := PrimTail(tmp12493)

										tmp12495 := PrimEqual(Nil, tmp12494)

										var ifres12492 Obj

										if True == tmp12495 {
											ifres12492 = True

										} else {
											ifres12492 = False

										}

										ifres12491 = ifres12492

									} else {
										ifres12491 = False

									}

									var ifres12490 Obj

									if True == ifres12491 {
										ifres12490 = True

									} else {
										ifres12490 = False

									}

									ifres12489 = ifres12490

								} else {
									ifres12489 = False

								}

								if True == ifres12489 {
									tmp12483 := PrimHead(V4282)

									tmp12484 := Call(__e, PrimNS2Value(symshen_4curry), tmp12483)

									tmp12485 := PrimTail(V4282)

									tmp12486 := PrimHead(tmp12485)

									tmp12487 := Call(__e, PrimNS2Value(symshen_4curry), tmp12486)

									tmp12488 := PrimCons(tmp12487, Nil)

									__e.Return(PrimCons(tmp12484, tmp12488))
									return

								} else {
									__e.Return(V4282)
									return
								}

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp12594 := Call(__e, PrimNS2Value(symdef), symshen_4curry, tmp12475)

	_ = tmp12594

	tmp12595 := MakeNative(func(__e *ControlFlow) {
		V4283 := __e.Get(1)
		_ = V4283
		tmp12596 := PrimNS3Value(symshen_4_dspecial_d)

		__e.TailApply(PrimNS2Value(symelement_2), V4283, tmp12596)
		return

	}, 1)

	tmp12597 := Call(__e, PrimNS2Value(symdef), symshen_4special_2, tmp12595)

	_ = tmp12597

	tmp12598 := MakeNative(func(__e *ControlFlow) {
		V4284 := __e.Get(1)
		_ = V4284
		tmp12599 := PrimNS3Value(symshen_4_dextraspecial_d)

		__e.TailApply(PrimNS2Value(symelement_2), V4284, tmp12599)
		return

	}, 1)

	tmp12600 := Call(__e, PrimNS2Value(symdef), symshen_4extraspecial_2, tmp12598)

	_ = tmp12600

	tmp12601 := MakeNative(func(__e *ControlFlow) {
		V4285 := __e.Get(1)
		_ = V4285
		V4286 := __e.Get(2)
		_ = V4286
		V4287 := __e.Get(3)
		_ = V4287
		V4288 := __e.Get(4)
		_ = V4288
		V4289 := __e.Get(5)
		_ = V4289
		V4290 := __e.Get(6)
		_ = V4290
		tmp12602 := MakeNative(func(__e *ControlFlow) {
			K3697 := __e.Get(1)
			_ = K3697
			tmp12603 := MakeNative(func(__e *ControlFlow) {
				C3701 := __e.Get(1)
				_ = C3701
				tmp12661 := PrimEqual(C3701, False)

				if True == tmp12661 {
					tmp12605 := MakeNative(func(__e *ControlFlow) {
						C3702 := __e.Get(1)
						_ = C3702
						tmp12624 := PrimEqual(C3702, False)

						if True == tmp12624 {
							tmp12607 := MakeNative(func(__e *ControlFlow) {
								C3707 := __e.Get(1)
								_ = C3707
								tmp12617 := PrimEqual(C3707, False)

								if True == tmp12617 {
									tmp12609 := MakeNative(func(__e *ControlFlow) {
										C3708 := __e.Get(1)
										_ = C3708
										tmp12611 := PrimEqual(C3708, False)

										if True == tmp12611 {
											__e.TailApply(PrimNS2Value(symshen_4unlock), V4288, K3697)
											return
										} else {
											__e.Return(C3708)
											return
										}

									}, 1)

									tmp12616 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4288)

									var ifres12612 Obj

									if True == tmp12616 {
										tmp12613 := Call(__e, PrimNS2Value(symshen_4incinfs))

										_ = tmp12613

										tmp12614 := PrimNS3Value(symshen_4_ddatatypes_d)

										tmp12615 := Call(__e, PrimNS2Value(symshen_4search_1user_1datatypes), V4285, V4286, tmp12614, V4287, V4288, K3697, V4290)

										ifres12612 = tmp12615

									} else {
										ifres12612 = False

									}

									__e.TailApply(tmp12609, ifres12612)
									return

								} else {
									__e.Return(C3707)
									return
								}

							}, 1)

							tmp12623 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4288)

							var ifres12618 Obj

							if True == tmp12623 {
								tmp12619 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp12619

								tmp12620 := PrimNS3Value(symshen_4_dspy_d)

								tmp12621 := MakeNative(func(__e *ControlFlow) {
									__e.TailApply(PrimNS2Value(symshen_4show), V4285, V4286, V4287, V4288, K3697, V4290)
									return
								}, 0)

								tmp12622 := Call(__e, PrimNS2Value(symwhen), tmp12620, V4287, V4288, K3697, tmp12621)

								ifres12618 = tmp12622

							} else {
								ifres12618 = False

							}

							__e.TailApply(tmp12607, ifres12618)
							return

						} else {
							__e.Return(C3702)
							return
						}

					}, 1)

					tmp12660 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4288)

					var ifres12625 Obj

					if True == tmp12660 {
						tmp12626 := MakeNative(func(__e *ControlFlow) {
							Tm3703 := __e.Get(1)
							_ = Tm3703
							tmp12657 := PrimIsPair(Tm3703)

							if True == tmp12657 {
								tmp12628 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									tmp12629 := MakeNative(func(__e *ControlFlow) {
										Tm3704 := __e.Get(1)
										_ = Tm3704
										tmp12653 := PrimIsPair(Tm3704)

										if True == tmp12653 {
											tmp12631 := MakeNative(func(__e *ControlFlow) {
												Colon := __e.Get(1)
												_ = Colon
												tmp12632 := MakeNative(func(__e *ControlFlow) {
													Tm3705 := __e.Get(1)
													_ = Tm3705
													tmp12649 := PrimIsPair(Tm3705)

													if True == tmp12649 {
														tmp12634 := MakeNative(func(__e *ControlFlow) {
															A := __e.Get(1)
															_ = A
															tmp12635 := MakeNative(func(__e *ControlFlow) {
																Tm3706 := __e.Get(1)
																_ = Tm3706
																tmp12645 := PrimEqual(Tm3706, Nil)

																if True == tmp12645 {
																	tmp12637 := Call(__e, PrimNS2Value(symshen_4incinfs))

																	_ = tmp12637

																	tmp12638 := Call(__e, PrimNS2Value(symshen_4deref), Colon, V4287)

																	tmp12639 := PrimIntern(MakeString(":"))

																	tmp12640 := PrimEqual(tmp12638, tmp12639)

																	tmp12641 := MakeNative(func(__e *ControlFlow) {
																		tmp12642 := Call(__e, PrimNS2Value(symshen_4type_1theory_1enabled_2))

																		tmp12643 := MakeNative(func(__e *ControlFlow) {
																			tmp12644 := MakeNative(func(__e *ControlFlow) {
																				__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, A, V4286, V4287, V4288, K3697, V4290)
																				return
																			}, 0)

																			__e.TailApply(PrimNS2Value(symshen_4cut), V4287, V4288, K3697, tmp12644)
																			return

																		}, 0)

																		__e.TailApply(PrimNS2Value(symwhen), tmp12642, V4287, V4288, K3697, tmp12643)
																		return

																	}, 0)

																	__e.TailApply(PrimNS2Value(symwhen), tmp12640, V4287, V4288, K3697, tmp12641)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp12646 := PrimTail(Tm3705)

															tmp12647 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12646, V4287)

															__e.TailApply(tmp12635, tmp12647)
															return

														}, 1)

														tmp12648 := PrimHead(Tm3705)

														__e.TailApply(tmp12634, tmp12648)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp12650 := PrimTail(Tm3704)

												tmp12651 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12650, V4287)

												__e.TailApply(tmp12632, tmp12651)
												return

											}, 1)

											tmp12652 := PrimHead(Tm3704)

											__e.TailApply(tmp12631, tmp12652)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp12654 := PrimTail(Tm3703)

									tmp12655 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12654, V4287)

									__e.TailApply(tmp12629, tmp12655)
									return

								}, 1)

								tmp12656 := PrimHead(Tm3703)

								__e.TailApply(tmp12628, tmp12656)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp12658 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4285, V4287)

						tmp12659 := Call(__e, tmp12626, tmp12658)

						ifres12625 = tmp12659

					} else {
						ifres12625 = False

					}

					__e.TailApply(tmp12605, ifres12625)
					return

				} else {
					__e.Return(C3701)
					return
				}

			}, 1)

			tmp12666 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4288)

			var ifres12662 Obj

			if True == tmp12666 {
				tmp12663 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp12663

				tmp12664 := Call(__e, PrimNS2Value(symshen_4maxinfexceeded_2))

				tmp12665 := Call(__e, PrimNS2Value(symwhen), tmp12664, V4287, V4288, K3697, V4290)

				ifres12662 = tmp12665

			} else {
				ifres12662 = False

			}

			__e.TailApply(tmp12603, ifres12662)
			return

		}, 1)

		tmp12667 := PrimNumberAdd(V4289, MakeNumber(1))

		__e.TailApply(tmp12602, tmp12667)
		return

	}, 6)

	tmp12668 := Call(__e, PrimNS2Value(symdef), symshen_4system_1S, tmp12601)

	_ = tmp12668

	tmp12669 := MakeNative(func(__e *ControlFlow) {
		V4297 := __e.Get(1)
		_ = V4297
		V4298 := __e.Get(2)
		_ = V4298
		V4299 := __e.Get(3)
		_ = V4299
		V4300 := __e.Get(4)
		_ = V4300
		V4301 := __e.Get(5)
		_ = V4301
		V4302 := __e.Get(6)
		_ = V4302
		tmp12670 := Call(__e, PrimNS2Value(symshen_4line))

		_ = tmp12670

		tmp12671 := Call(__e, PrimNS2Value(symshen_4deref), V4297, V4299)

		tmp12672 := Call(__e, PrimNS2Value(symshen_4show_1p), tmp12671)

		_ = tmp12672

		tmp12673 := Call(__e, PrimNS2Value(symnl), MakeNumber(2))

		_ = tmp12673

		tmp12674 := Call(__e, PrimNS2Value(symshen_4deref), V4298, V4299)

		tmp12675 := Call(__e, PrimNS2Value(symshen_4show_1assumptions), tmp12674, MakeNumber(1))

		_ = tmp12675

		tmp12676 := Call(__e, PrimNS2Value(symshen_4pause_1for_1user))

		_ = tmp12676

		__e.Return(False)
		return

	}, 6)

	tmp12677 := Call(__e, PrimNS2Value(symdef), symshen_4show, tmp12669)

	_ = tmp12677

	tmp12678 := MakeNative(func(__e *ControlFlow) {
		tmp12679 := MakeNative(func(__e *ControlFlow) {
			Infs := __e.Get(1)
			_ = Infs
			tmp12681 := PrimEqual(MakeNumber(1), Infs)

			var ifres12680 Obj

			if True == tmp12681 {
				ifres12680 = MakeString("")

			} else {
				ifres12680 = MakeString("s")

			}

			tmp12682 := Call(__e, PrimNS2Value(symshen_4app), ifres12680, MakeString(" \n?- "), symshen_4a)

			tmp12683 := PrimStringConcat(MakeString(" inference"), tmp12682)

			tmp12684 := Call(__e, PrimNS2Value(symshen_4app), Infs, tmp12683, symshen_4a)

			tmp12685 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp12684)

			tmp12686 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp12685, tmp12686)
			return

		}, 1)

		tmp12687 := Call(__e, PrimNS2Value(syminferences))

		__e.TailApply(tmp12679, tmp12687)
		return

	}, 0)

	tmp12688 := Call(__e, PrimNS2Value(symdef), symshen_4line, tmp12678)

	_ = tmp12688

	tmp12689 := MakeNative(func(__e *ControlFlow) {
		V4303 := __e.Get(1)
		_ = V4303
		tmp12721 := PrimIsPair(V4303)

		var ifres12700 Obj

		if True == tmp12721 {
			tmp12719 := PrimTail(V4303)

			tmp12720 := PrimIsPair(tmp12719)

			var ifres12702 Obj

			if True == tmp12720 {
				tmp12716 := PrimTail(V4303)

				tmp12717 := PrimTail(tmp12716)

				tmp12718 := PrimIsPair(tmp12717)

				var ifres12704 Obj

				if True == tmp12718 {
					tmp12712 := PrimTail(V4303)

					tmp12713 := PrimTail(tmp12712)

					tmp12714 := PrimTail(tmp12713)

					tmp12715 := PrimEqual(Nil, tmp12714)

					var ifres12706 Obj

					if True == tmp12715 {
						tmp12708 := PrimTail(V4303)

						tmp12709 := PrimHead(tmp12708)

						tmp12710 := PrimIntern(MakeString(":"))

						tmp12711 := PrimEqual(tmp12709, tmp12710)

						var ifres12707 Obj

						if True == tmp12711 {
							ifres12707 = True

						} else {
							ifres12707 = False

						}

						ifres12706 = ifres12707

					} else {
						ifres12706 = False

					}

					var ifres12705 Obj

					if True == ifres12706 {
						ifres12705 = True

					} else {
						ifres12705 = False

					}

					ifres12704 = ifres12705

				} else {
					ifres12704 = False

				}

				var ifres12703 Obj

				if True == ifres12704 {
					ifres12703 = True

				} else {
					ifres12703 = False

				}

				ifres12702 = ifres12703

			} else {
				ifres12702 = False

			}

			var ifres12701 Obj

			if True == ifres12702 {
				ifres12701 = True

			} else {
				ifres12701 = False

			}

			ifres12700 = ifres12701

		} else {
			ifres12700 = False

		}

		if True == ifres12700 {
			tmp12691 := PrimHead(V4303)

			tmp12692 := Call(__e, PrimNS2Value(symshen_4prterm), tmp12691)

			_ = tmp12692

			tmp12693 := Call(__e, PrimNS2Value(symstoutput))

			tmp12694 := Call(__e, PrimNS2Value(sympr), MakeString(" : "), tmp12693)

			_ = tmp12694

			tmp12695 := PrimTail(V4303)

			tmp12696 := PrimTail(tmp12695)

			tmp12697 := PrimHead(tmp12696)

			tmp12698 := Call(__e, PrimNS2Value(symshen_4app), tmp12697, MakeString(""), symshen_4r)

			tmp12699 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp12698, tmp12699)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4prterm), V4303)
			return
		}

	}, 1)

	tmp12722 := Call(__e, PrimNS2Value(symdef), symshen_4show_1p, tmp12689)

	_ = tmp12722

	tmp12723 := MakeNative(func(__e *ControlFlow) {
		V4304 := __e.Get(1)
		_ = V4304
		tmp12766 := PrimIsPair(V4304)

		var ifres12747 Obj

		if True == tmp12766 {
			tmp12764 := PrimHead(V4304)

			tmp12765 := PrimEqual(symcons, tmp12764)

			var ifres12749 Obj

			if True == tmp12765 {
				tmp12762 := PrimTail(V4304)

				tmp12763 := PrimIsPair(tmp12762)

				var ifres12751 Obj

				if True == tmp12763 {
					tmp12759 := PrimTail(V4304)

					tmp12760 := PrimTail(tmp12759)

					tmp12761 := PrimIsPair(tmp12760)

					var ifres12753 Obj

					if True == tmp12761 {
						tmp12755 := PrimTail(V4304)

						tmp12756 := PrimTail(tmp12755)

						tmp12757 := PrimTail(tmp12756)

						tmp12758 := PrimEqual(Nil, tmp12757)

						var ifres12754 Obj

						if True == tmp12758 {
							ifres12754 = True

						} else {
							ifres12754 = False

						}

						ifres12753 = ifres12754

					} else {
						ifres12753 = False

					}

					var ifres12752 Obj

					if True == ifres12753 {
						ifres12752 = True

					} else {
						ifres12752 = False

					}

					ifres12751 = ifres12752

				} else {
					ifres12751 = False

				}

				var ifres12750 Obj

				if True == ifres12751 {
					ifres12750 = True

				} else {
					ifres12750 = False

				}

				ifres12749 = ifres12750

			} else {
				ifres12749 = False

			}

			var ifres12748 Obj

			if True == ifres12749 {
				ifres12748 = True

			} else {
				ifres12748 = False

			}

			ifres12747 = ifres12748

		} else {
			ifres12747 = False

		}

		if True == ifres12747 {
			tmp12737 := Call(__e, PrimNS2Value(symstoutput))

			tmp12738 := Call(__e, PrimNS2Value(sympr), MakeString("["), tmp12737)

			_ = tmp12738

			tmp12739 := PrimTail(V4304)

			tmp12740 := PrimHead(tmp12739)

			tmp12741 := Call(__e, PrimNS2Value(symshen_4prterm), tmp12740)

			_ = tmp12741

			tmp12742 := PrimTail(V4304)

			tmp12743 := PrimTail(tmp12742)

			tmp12744 := PrimHead(tmp12743)

			tmp12745 := Call(__e, PrimNS2Value(symshen_4prtl), tmp12744)

			_ = tmp12745

			tmp12746 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), MakeString("]"), tmp12746)
			return

		} else {
			tmp12736 := PrimIsPair(V4304)

			if True == tmp12736 {
				tmp12726 := Call(__e, PrimNS2Value(symstoutput))

				tmp12727 := Call(__e, PrimNS2Value(sympr), MakeString("("), tmp12726)

				_ = tmp12727

				tmp12728 := PrimHead(V4304)

				tmp12729 := Call(__e, PrimNS2Value(symshen_4prterm), tmp12728)

				_ = tmp12729

				tmp12730 := MakeNative(func(__e *ControlFlow) {
					Y := __e.Get(1)
					_ = Y
					tmp12731 := Call(__e, PrimNS2Value(symstoutput))

					tmp12732 := Call(__e, PrimNS2Value(sympr), MakeString(" "), tmp12731)

					_ = tmp12732

					__e.TailApply(PrimNS2Value(symshen_4prterm), Y)
					return

				}, 1)

				tmp12733 := PrimTail(V4304)

				tmp12734 := Call(__e, PrimNS2Value(symmap), tmp12730, tmp12733)

				_ = tmp12734

				tmp12735 := Call(__e, PrimNS2Value(symstoutput))

				__e.TailApply(PrimNS2Value(sympr), MakeString(")"), tmp12735)
				return

			} else {
				__e.TailApply(PrimNS2Value(symprint), V4304)
				return
			}

		}

	}, 1)

	tmp12767 := Call(__e, PrimNS2Value(symdef), symshen_4prterm, tmp12723)

	_ = tmp12767

	tmp12768 := MakeNative(func(__e *ControlFlow) {
		V4305 := __e.Get(1)
		_ = V4305
		tmp12801 := PrimEqual(Nil, V4305)

		if True == tmp12801 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp12800 := PrimIsPair(V4305)

			var ifres12781 Obj

			if True == tmp12800 {
				tmp12798 := PrimHead(V4305)

				tmp12799 := PrimEqual(symcons, tmp12798)

				var ifres12783 Obj

				if True == tmp12799 {
					tmp12796 := PrimTail(V4305)

					tmp12797 := PrimIsPair(tmp12796)

					var ifres12785 Obj

					if True == tmp12797 {
						tmp12793 := PrimTail(V4305)

						tmp12794 := PrimTail(tmp12793)

						tmp12795 := PrimIsPair(tmp12794)

						var ifres12787 Obj

						if True == tmp12795 {
							tmp12789 := PrimTail(V4305)

							tmp12790 := PrimTail(tmp12789)

							tmp12791 := PrimTail(tmp12790)

							tmp12792 := PrimEqual(Nil, tmp12791)

							var ifres12788 Obj

							if True == tmp12792 {
								ifres12788 = True

							} else {
								ifres12788 = False

							}

							ifres12787 = ifres12788

						} else {
							ifres12787 = False

						}

						var ifres12786 Obj

						if True == ifres12787 {
							ifres12786 = True

						} else {
							ifres12786 = False

						}

						ifres12785 = ifres12786

					} else {
						ifres12785 = False

					}

					var ifres12784 Obj

					if True == ifres12785 {
						ifres12784 = True

					} else {
						ifres12784 = False

					}

					ifres12783 = ifres12784

				} else {
					ifres12783 = False

				}

				var ifres12782 Obj

				if True == ifres12783 {
					ifres12782 = True

				} else {
					ifres12782 = False

				}

				ifres12781 = ifres12782

			} else {
				ifres12781 = False

			}

			if True == ifres12781 {
				tmp12773 := Call(__e, PrimNS2Value(symstoutput))

				tmp12774 := Call(__e, PrimNS2Value(sympr), MakeString(" "), tmp12773)

				_ = tmp12774

				tmp12775 := PrimTail(V4305)

				tmp12776 := PrimHead(tmp12775)

				tmp12777 := Call(__e, PrimNS2Value(symshen_4prterm), tmp12776)

				_ = tmp12777

				tmp12778 := PrimTail(V4305)

				tmp12779 := PrimTail(tmp12778)

				tmp12780 := PrimHead(tmp12779)

				__e.TailApply(PrimNS2Value(symshen_4prtl), tmp12780)
				return

			} else {
				tmp12771 := Call(__e, PrimNS2Value(symstoutput))

				tmp12772 := Call(__e, PrimNS2Value(sympr), MakeString(" | "), tmp12771)

				_ = tmp12772

				__e.TailApply(PrimNS2Value(symshen_4prterm), V4305)
				return

			}

		}

	}, 1)

	tmp12802 := Call(__e, PrimNS2Value(symdef), symshen_4prtl, tmp12768)

	_ = tmp12802

	tmp12803 := MakeNative(func(__e *ControlFlow) {
		V4312 := __e.Get(1)
		_ = V4312
		V4313 := __e.Get(2)
		_ = V4313
		tmp12816 := PrimEqual(Nil, V4312)

		if True == tmp12816 {
			tmp12815 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), MakeString("\n> "), tmp12815)
			return

		} else {
			tmp12814 := PrimIsPair(V4312)

			if True == tmp12814 {
				tmp12806 := Call(__e, PrimNS2Value(symshen_4app), V4313, MakeString(". "), symshen_4a)

				tmp12807 := Call(__e, PrimNS2Value(symstoutput))

				tmp12808 := Call(__e, PrimNS2Value(sympr), tmp12806, tmp12807)

				_ = tmp12808

				tmp12809 := PrimHead(V4312)

				tmp12810 := Call(__e, PrimNS2Value(symshen_4show_1p), tmp12809)

				_ = tmp12810

				tmp12811 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

				_ = tmp12811

				tmp12812 := PrimTail(V4312)

				tmp12813 := PrimNumberAdd(V4313, MakeNumber(1))

				__e.TailApply(PrimNS2Value(symshen_4show_1assumptions), tmp12812, tmp12813)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
				return
			}

		}

	}, 2)

	tmp12817 := Call(__e, PrimNS2Value(symdef), symshen_4show_1assumptions, tmp12803)

	_ = tmp12817

	tmp12818 := MakeNative(func(__e *ControlFlow) {
		tmp12819 := MakeNative(func(__e *ControlFlow) {
			Byte := __e.Get(1)
			_ = Byte
			tmp12821 := PrimEqual(Byte, MakeNumber(94))

			if True == tmp12821 {
				__e.Return(PrimSimpleError(MakeString("input aborted\n")))
				return
			} else {
				__e.TailApply(PrimNS2Value(symnl), MakeNumber(1))
				return
			}

		}, 1)

		tmp12822 := Call(__e, PrimNS2Value(symstinput))

		tmp12823 := PrimReadByte(tmp12822)

		__e.TailApply(tmp12819, tmp12823)
		return

	}, 0)

	tmp12824 := Call(__e, PrimNS2Value(symdef), symshen_4pause_1for_1user, tmp12818)

	_ = tmp12824

	tmp12825 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dshen_1type_1theory_1enabled_2_d))
		return
	}, 0)

	tmp12826 := Call(__e, PrimNS2Value(symdef), symshen_4type_1theory_1enabled_2, tmp12825)

	_ = tmp12826

	tmp12827 := MakeNative(func(__e *ControlFlow) {
		tmp12829 := Call(__e, PrimNS2Value(syminferences))

		tmp12830 := PrimNS3Value(symshen_4_dmaxinferences_d)

		tmp12831 := PrimGreatThan(tmp12829, tmp12830)

		if True == tmp12831 {
			__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
			return
		} else {
			__e.Return(False)
			return
		}

	}, 0)

	tmp12832 := Call(__e, PrimNS2Value(symdef), symshen_4maxinfexceeded_2, tmp12827)

	_ = tmp12832

	tmp12833 := MakeNative(func(__e *ControlFlow) {
		V4314 := __e.Get(1)
		_ = V4314
		V4315 := __e.Get(2)
		_ = V4315
		V4316 := __e.Get(3)
		_ = V4316
		V4317 := __e.Get(4)
		_ = V4317
		V4318 := __e.Get(5)
		_ = V4318
		V4319 := __e.Get(6)
		_ = V4319
		V4320 := __e.Get(7)
		_ = V4320
		tmp12834 := MakeNative(func(__e *ControlFlow) {
			K3712 := __e.Get(1)
			_ = K3712
			tmp12835 := MakeNative(func(__e *ControlFlow) {
				C3717 := __e.Get(1)
				_ = C3717
				tmp13715 := PrimEqual(C3717, False)

				if True == tmp13715 {
					tmp12837 := MakeNative(func(__e *ControlFlow) {
						C3718 := __e.Get(1)
						_ = C3718
						tmp13706 := PrimEqual(C3718, False)

						if True == tmp13706 {
							tmp12839 := MakeNative(func(__e *ControlFlow) {
								C3719 := __e.Get(1)
								_ = C3719
								tmp13701 := PrimEqual(C3719, False)

								if True == tmp13701 {
									tmp12841 := MakeNative(func(__e *ControlFlow) {
										C3720 := __e.Get(1)
										_ = C3720
										tmp13683 := PrimEqual(C3720, False)

										if True == tmp13683 {
											tmp12843 := MakeNative(func(__e *ControlFlow) {
												C3723 := __e.Get(1)
												_ = C3723
												tmp13657 := PrimEqual(C3723, False)

												if True == tmp13657 {
													tmp12845 := MakeNative(func(__e *ControlFlow) {
														C3728 := __e.Get(1)
														_ = C3728
														tmp13623 := PrimEqual(C3728, False)

														if True == tmp13623 {
															tmp12847 := MakeNative(func(__e *ControlFlow) {
																C3732 := __e.Get(1)
																_ = C3732
																tmp13593 := PrimEqual(C3732, False)

																if True == tmp13593 {
																	tmp12849 := MakeNative(func(__e *ControlFlow) {
																		C3736 := __e.Get(1)
																		_ = C3736
																		tmp13509 := PrimEqual(C3736, False)

																		if True == tmp13509 {
																			tmp12851 := MakeNative(func(__e *ControlFlow) {
																				C3750 := __e.Get(1)
																				_ = C3750
																				tmp13404 := PrimEqual(C3750, False)

																				if True == tmp13404 {
																					tmp12853 := MakeNative(func(__e *ControlFlow) {
																						C3766 := __e.Get(1)
																						_ = C3766
																						tmp13320 := PrimEqual(C3766, False)

																						if True == tmp13320 {
																							tmp12855 := MakeNative(func(__e *ControlFlow) {
																								C3780 := __e.Get(1)
																								_ = C3780
																								tmp13278 := PrimEqual(C3780, False)

																								if True == tmp13278 {
																									tmp12857 := MakeNative(func(__e *ControlFlow) {
																										C3788 := __e.Get(1)
																										_ = C3788
																										tmp13155 := PrimEqual(C3788, False)

																										if True == tmp13155 {
																											tmp12859 := MakeNative(func(__e *ControlFlow) {
																												C3804 := __e.Get(1)
																												_ = C3804
																												tmp13092 := PrimEqual(C3804, False)

																												if True == tmp13092 {
																													tmp12861 := MakeNative(func(__e *ControlFlow) {
																														C3811 := __e.Get(1)
																														_ = C3811
																														tmp13005 := PrimEqual(C3811, False)

																														if True == tmp13005 {
																															tmp12863 := MakeNative(func(__e *ControlFlow) {
																																C3825 := __e.Get(1)
																																_ = C3825
																																tmp12968 := PrimEqual(C3825, False)

																																if True == tmp12968 {
																																	tmp12865 := MakeNative(func(__e *ControlFlow) {
																																		C3831 := __e.Get(1)
																																		_ = C3831
																																		tmp12930 := PrimEqual(C3831, False)

																																		if True == tmp12930 {
																																			tmp12867 := MakeNative(func(__e *ControlFlow) {
																																				C3837 := __e.Get(1)
																																				_ = C3837
																																				tmp12893 := PrimEqual(C3837, False)

																																				if True == tmp12893 {
																																					tmp12869 := MakeNative(func(__e *ControlFlow) {
																																						C3843 := __e.Get(1)
																																						_ = C3843
																																						tmp12883 := PrimEqual(C3843, False)

																																						if True == tmp12883 {
																																							tmp12871 := MakeNative(func(__e *ControlFlow) {
																																								C3844 := __e.Get(1)
																																								_ = C3844
																																								tmp12873 := PrimEqual(C3844, False)

																																								if True == tmp12873 {
																																									__e.TailApply(PrimNS2Value(symshen_4unlock), V4318, K3712)
																																									return
																																								} else {
																																									__e.Return(C3844)
																																									return
																																								}

																																							}, 1)

																																							tmp12882 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																																							var ifres12874 Obj

																																							if True == tmp12882 {
																																								tmp12875 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp12875

																																								tmp12876 := PrimIntern(MakeString(":"))

																																								tmp12877 := PrimCons(V4315, Nil)

																																								tmp12878 := PrimCons(tmp12876, tmp12877)

																																								tmp12879 := PrimCons(V4314, tmp12878)

																																								tmp12880 := PrimNS3Value(symshen_4_ddatatypes_d)

																																								tmp12881 := Call(__e, PrimNS2Value(symshen_4search_1user_1datatypes), tmp12879, V4316, tmp12880, V4317, V4318, K3712, V4320)

																																								ifres12874 = tmp12881

																																							} else {
																																								ifres12874 = False

																																							}

																																							__e.TailApply(tmp12871, ifres12874)
																																							return

																																						} else {
																																							__e.Return(C3843)
																																							return
																																						}

																																					}, 1)

																																					tmp12892 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																																					var ifres12884 Obj

																																					if True == tmp12892 {
																																						tmp12885 := MakeNative(func(__e *ControlFlow) {
																																							Normalised := __e.Get(1)
																																							_ = Normalised
																																							tmp12886 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																							_ = tmp12886

																																							tmp12887 := MakeNative(func(__e *ControlFlow) {
																																								tmp12888 := MakeNative(func(__e *ControlFlow) {
																																									__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), V4314, V4315, Normalised, V4317, V4318, K3712, V4320)
																																									return
																																								}, 0)

																																								__e.TailApply(PrimNS2Value(symshen_4cut), V4317, V4318, K3712, tmp12888)
																																								return

																																							}, 0)

																																							tmp12889 := Call(__e, PrimNS2Value(symshen_4l_1rules), V4316, Normalised, False, V4317, V4318, K3712, tmp12887)

																																							__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp12889)
																																							return

																																						}, 1)

																																						tmp12890 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																						tmp12891 := Call(__e, tmp12885, tmp12890)

																																						ifres12884 = tmp12891

																																					} else {
																																						ifres12884 = False

																																					}

																																					__e.TailApply(tmp12869, ifres12884)
																																					return

																																				} else {
																																					__e.Return(C3837)
																																					return
																																				}

																																			}, 1)

																																			tmp12929 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																																			var ifres12894 Obj

																																			if True == tmp12929 {
																																				tmp12895 := MakeNative(func(__e *ControlFlow) {
																																					Tm3838 := __e.Get(1)
																																					_ = Tm3838
																																					tmp12926 := PrimIsPair(Tm3838)

																																					if True == tmp12926 {
																																						tmp12897 := MakeNative(func(__e *ControlFlow) {
																																							Tm3839 := __e.Get(1)
																																							_ = Tm3839
																																							tmp12923 := PrimEqual(Tm3839, symset)

																																							if True == tmp12923 {
																																								tmp12899 := MakeNative(func(__e *ControlFlow) {
																																									Tm3840 := __e.Get(1)
																																									_ = Tm3840
																																									tmp12920 := PrimIsPair(Tm3840)

																																									if True == tmp12920 {
																																										tmp12901 := MakeNative(func(__e *ControlFlow) {
																																											Var := __e.Get(1)
																																											_ = Var
																																											tmp12902 := MakeNative(func(__e *ControlFlow) {
																																												Tm3841 := __e.Get(1)
																																												_ = Tm3841
																																												tmp12916 := PrimIsPair(Tm3841)

																																												if True == tmp12916 {
																																													tmp12904 := MakeNative(func(__e *ControlFlow) {
																																														Val := __e.Get(1)
																																														_ = Val
																																														tmp12905 := MakeNative(func(__e *ControlFlow) {
																																															Tm3842 := __e.Get(1)
																																															_ = Tm3842
																																															tmp12912 := PrimEqual(Tm3842, Nil)

																																															if True == tmp12912 {
																																																tmp12907 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																_ = tmp12907

																																																tmp12908 := MakeNative(func(__e *ControlFlow) {
																																																	tmp12909 := PrimCons(Var, Nil)

																																																	tmp12910 := PrimCons(symvalue, tmp12909)

																																																	tmp12911 := MakeNative(func(__e *ControlFlow) {
																																																		__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Val, V4315, V4316, V4317, V4318, K3712, V4320)
																																																		return
																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), tmp12910, V4315, V4316, V4317, V4318, K3712, tmp12911)
																																																	return

																																																}, 0)

																																																__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Var, symsymbol, V4316, V4317, V4318, K3712, tmp12908)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}, 1)

																																														tmp12913 := PrimTail(Tm3841)

																																														tmp12914 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12913, V4317)

																																														__e.TailApply(tmp12905, tmp12914)
																																														return

																																													}, 1)

																																													tmp12915 := PrimHead(Tm3841)

																																													__e.TailApply(tmp12904, tmp12915)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp12917 := PrimTail(Tm3840)

																																											tmp12918 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12917, V4317)

																																											__e.TailApply(tmp12902, tmp12918)
																																											return

																																										}, 1)

																																										tmp12919 := PrimHead(Tm3840)

																																										__e.TailApply(tmp12901, tmp12919)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp12921 := PrimTail(Tm3838)

																																								tmp12922 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12921, V4317)

																																								__e.TailApply(tmp12899, tmp12922)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp12924 := PrimHead(Tm3838)

																																						tmp12925 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12924, V4317)

																																						__e.TailApply(tmp12897, tmp12925)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp12927 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																																				tmp12928 := Call(__e, tmp12895, tmp12927)

																																				ifres12894 = tmp12928

																																			} else {
																																				ifres12894 = False

																																			}

																																			__e.TailApply(tmp12867, ifres12894)
																																			return

																																		} else {
																																			__e.Return(C3831)
																																			return
																																		}

																																	}, 1)

																																	tmp12967 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																																	var ifres12931 Obj

																																	if True == tmp12967 {
																																		tmp12932 := MakeNative(func(__e *ControlFlow) {
																																			Tm3832 := __e.Get(1)
																																			_ = Tm3832
																																			tmp12964 := PrimIsPair(Tm3832)

																																			if True == tmp12964 {
																																				tmp12934 := MakeNative(func(__e *ControlFlow) {
																																					Tm3833 := __e.Get(1)
																																					_ = Tm3833
																																					tmp12961 := PrimEqual(Tm3833, syminput_7)

																																					if True == tmp12961 {
																																						tmp12936 := MakeNative(func(__e *ControlFlow) {
																																							Tm3834 := __e.Get(1)
																																							_ = Tm3834
																																							tmp12958 := PrimIsPair(Tm3834)

																																							if True == tmp12958 {
																																								tmp12938 := MakeNative(func(__e *ControlFlow) {
																																									A := __e.Get(1)
																																									_ = A
																																									tmp12939 := MakeNative(func(__e *ControlFlow) {
																																										Tm3835 := __e.Get(1)
																																										_ = Tm3835
																																										tmp12954 := PrimIsPair(Tm3835)

																																										if True == tmp12954 {
																																											tmp12941 := MakeNative(func(__e *ControlFlow) {
																																												Stream := __e.Get(1)
																																												_ = Stream
																																												tmp12942 := MakeNative(func(__e *ControlFlow) {
																																													Tm3836 := __e.Get(1)
																																													_ = Tm3836
																																													tmp12950 := PrimEqual(Tm3836, Nil)

																																													if True == tmp12950 {
																																														tmp12944 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																														_ = tmp12944

																																														tmp12945 := Call(__e, PrimNS2Value(symshen_4deref), A, V4317)

																																														tmp12946 := Call(__e, PrimNS2Value(symshen_4rectify_1type), tmp12945)

																																														tmp12947 := MakeNative(func(__e *ControlFlow) {
																																															tmp12948 := PrimCons(symin, Nil)

																																															tmp12949 := PrimCons(symstream, tmp12948)

																																															__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Stream, tmp12949, V4316, V4317, V4318, K3712, V4320)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimNS2Value(symis_b), V4315, tmp12946, V4317, V4318, K3712, tmp12947)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}, 1)

																																												tmp12951 := PrimTail(Tm3835)

																																												tmp12952 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12951, V4317)

																																												__e.TailApply(tmp12942, tmp12952)
																																												return

																																											}, 1)

																																											tmp12953 := PrimHead(Tm3835)

																																											__e.TailApply(tmp12941, tmp12953)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp12955 := PrimTail(Tm3834)

																																									tmp12956 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12955, V4317)

																																									__e.TailApply(tmp12939, tmp12956)
																																									return

																																								}, 1)

																																								tmp12957 := PrimHead(Tm3834)

																																								__e.TailApply(tmp12938, tmp12957)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp12959 := PrimTail(Tm3832)

																																						tmp12960 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12959, V4317)

																																						__e.TailApply(tmp12936, tmp12960)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp12962 := PrimHead(Tm3832)

																																				tmp12963 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12962, V4317)

																																				__e.TailApply(tmp12934, tmp12963)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp12965 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																																		tmp12966 := Call(__e, tmp12932, tmp12965)

																																		ifres12931 = tmp12966

																																	} else {
																																		ifres12931 = False

																																	}

																																	__e.TailApply(tmp12865, ifres12931)
																																	return

																																} else {
																																	__e.Return(C3825)
																																	return
																																}

																															}, 1)

																															tmp13004 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																															var ifres12969 Obj

																															if True == tmp13004 {
																																tmp12970 := MakeNative(func(__e *ControlFlow) {
																																	Tm3826 := __e.Get(1)
																																	_ = Tm3826
																																	tmp13001 := PrimIsPair(Tm3826)

																																	if True == tmp13001 {
																																		tmp12972 := MakeNative(func(__e *ControlFlow) {
																																			Tm3827 := __e.Get(1)
																																			_ = Tm3827
																																			tmp12998 := PrimEqual(Tm3827, symtype)

																																			if True == tmp12998 {
																																				tmp12974 := MakeNative(func(__e *ControlFlow) {
																																					Tm3828 := __e.Get(1)
																																					_ = Tm3828
																																					tmp12995 := PrimIsPair(Tm3828)

																																					if True == tmp12995 {
																																						tmp12976 := MakeNative(func(__e *ControlFlow) {
																																							X := __e.Get(1)
																																							_ = X
																																							tmp12977 := MakeNative(func(__e *ControlFlow) {
																																								Tm3829 := __e.Get(1)
																																								_ = Tm3829
																																								tmp12991 := PrimIsPair(Tm3829)

																																								if True == tmp12991 {
																																									tmp12979 := MakeNative(func(__e *ControlFlow) {
																																										A := __e.Get(1)
																																										_ = A
																																										tmp12980 := MakeNative(func(__e *ControlFlow) {
																																											Tm3830 := __e.Get(1)
																																											_ = Tm3830
																																											tmp12987 := PrimEqual(Tm3830, Nil)

																																											if True == tmp12987 {
																																												tmp12982 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																												_ = tmp12982

																																												tmp12983 := MakeNative(func(__e *ControlFlow) {
																																													tmp12984 := Call(__e, PrimNS2Value(symshen_4deref), A, V4317)

																																													tmp12985 := Call(__e, PrimNS2Value(symshen_4rectify_1type), tmp12984)

																																													tmp12986 := MakeNative(func(__e *ControlFlow) {
																																														__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, V4315, V4316, V4317, V4318, K3712, V4320)
																																														return
																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symis_b), tmp12985, V4315, V4317, V4318, K3712, tmp12986)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symshen_4cut), V4317, V4318, K3712, tmp12983)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}, 1)

																																										tmp12988 := PrimTail(Tm3829)

																																										tmp12989 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12988, V4317)

																																										__e.TailApply(tmp12980, tmp12989)
																																										return

																																									}, 1)

																																									tmp12990 := PrimHead(Tm3829)

																																									__e.TailApply(tmp12979, tmp12990)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp12992 := PrimTail(Tm3828)

																																							tmp12993 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12992, V4317)

																																							__e.TailApply(tmp12977, tmp12993)
																																							return

																																						}, 1)

																																						tmp12994 := PrimHead(Tm3828)

																																						__e.TailApply(tmp12976, tmp12994)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp12996 := PrimTail(Tm3826)

																																				tmp12997 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12996, V4317)

																																				__e.TailApply(tmp12974, tmp12997)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp12999 := PrimHead(Tm3826)

																																		tmp13000 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12999, V4317)

																																		__e.TailApply(tmp12972, tmp13000)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13002 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																																tmp13003 := Call(__e, tmp12970, tmp13002)

																																ifres12969 = tmp13003

																															} else {
																																ifres12969 = False

																															}

																															__e.TailApply(tmp12863, ifres12969)
																															return

																														} else {
																															__e.Return(C3811)
																															return
																														}

																													}, 1)

																													tmp13091 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																													var ifres13006 Obj

																													if True == tmp13091 {
																														tmp13007 := MakeNative(func(__e *ControlFlow) {
																															Tm3812 := __e.Get(1)
																															_ = Tm3812
																															tmp13088 := PrimIsPair(Tm3812)

																															if True == tmp13088 {
																																tmp13009 := MakeNative(func(__e *ControlFlow) {
																																	Tm3813 := __e.Get(1)
																																	_ = Tm3813
																																	tmp13085 := PrimEqual(Tm3813, symopen)

																																	if True == tmp13085 {
																																		tmp13011 := MakeNative(func(__e *ControlFlow) {
																																			Tm3814 := __e.Get(1)
																																			_ = Tm3814
																																			tmp13082 := PrimIsPair(Tm3814)

																																			if True == tmp13082 {
																																				tmp13013 := MakeNative(func(__e *ControlFlow) {
																																					File := __e.Get(1)
																																					_ = File
																																					tmp13014 := MakeNative(func(__e *ControlFlow) {
																																						Tm3815 := __e.Get(1)
																																						_ = Tm3815
																																						tmp13078 := PrimIsPair(Tm3815)

																																						if True == tmp13078 {
																																							tmp13016 := MakeNative(func(__e *ControlFlow) {
																																								V3709 := __e.Get(1)
																																								_ = V3709
																																								tmp13017 := MakeNative(func(__e *ControlFlow) {
																																									Tm3816 := __e.Get(1)
																																									_ = Tm3816
																																									tmp13074 := PrimEqual(Tm3816, Nil)

																																									if True == tmp13074 {
																																										tmp13019 := MakeNative(func(__e *ControlFlow) {
																																											Tm3817 := __e.Get(1)
																																											_ = Tm3817
																																											tmp13020 := MakeNative(func(__e *ControlFlow) {
																																												GoTo3818 := __e.Get(1)
																																												_ = GoTo3818
																																												tmp13064 := PrimIsPair(Tm3817)

																																												if True == tmp13064 {
																																													tmp13030 := MakeNative(func(__e *ControlFlow) {
																																														Tm3819 := __e.Get(1)
																																														_ = Tm3819
																																														tmp13031 := MakeNative(func(__e *ControlFlow) {
																																															GoTo3820 := __e.Get(1)
																																															_ = GoTo3820
																																															tmp13035 := PrimEqual(Tm3819, symstream)

																																															if True == tmp13035 {
																																																__e.TailApply(PrimNS2Value(symthaw), GoTo3820)
																																																return
																																															} else {
																																																tmp13034 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3819)

																																																if True == tmp13034 {
																																																	__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3819, symstream, V4317, GoTo3820)
																																																	return
																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp13036 := MakeNative(func(__e *ControlFlow) {
																																															tmp13037 := MakeNative(func(__e *ControlFlow) {
																																																Tm3821 := __e.Get(1)
																																																_ = Tm3821
																																																tmp13038 := MakeNative(func(__e *ControlFlow) {
																																																	GoTo3822 := __e.Get(1)
																																																	_ = GoTo3822
																																																	tmp13058 := PrimIsPair(Tm3821)

																																																	if True == tmp13058 {
																																																		tmp13047 := MakeNative(func(__e *ControlFlow) {
																																																			D := __e.Get(1)
																																																			_ = D
																																																			tmp13048 := MakeNative(func(__e *ControlFlow) {
																																																				Tm3823 := __e.Get(1)
																																																				_ = Tm3823
																																																				tmp13049 := MakeNative(func(__e *ControlFlow) {
																																																					GoTo3824 := __e.Get(1)
																																																					_ = GoTo3824
																																																					tmp13053 := PrimEqual(Tm3823, Nil)

																																																					if True == tmp13053 {
																																																						__e.TailApply(PrimNS2Value(symthaw), GoTo3824)
																																																						return
																																																					} else {
																																																						tmp13052 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3823)

																																																						if True == tmp13052 {
																																																							__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3823, Nil, V4317, GoTo3824)
																																																							return
																																																						} else {
																																																							__e.Return(False)
																																																							return
																																																						}

																																																					}

																																																				}, 1)

																																																				tmp13054 := MakeNative(func(__e *ControlFlow) {
																																																					__e.TailApply(GoTo3822, D)
																																																					return
																																																				}, 0)

																																																				__e.TailApply(tmp13049, tmp13054)
																																																				return

																																																			}, 1)

																																																			tmp13055 := PrimTail(Tm3821)

																																																			tmp13056 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13055, V4317)

																																																			__e.TailApply(tmp13048, tmp13056)
																																																			return

																																																		}, 1)

																																																		tmp13057 := PrimHead(Tm3821)

																																																		__e.TailApply(tmp13047, tmp13057)
																																																		return

																																																	} else {
																																																		tmp13046 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3821)

																																																		if True == tmp13046 {
																																																			tmp13041 := MakeNative(func(__e *ControlFlow) {
																																																				D := __e.Get(1)
																																																				_ = D
																																																				tmp13042 := PrimCons(D, Nil)

																																																				tmp13043 := MakeNative(func(__e *ControlFlow) {
																																																					__e.TailApply(GoTo3822, D)
																																																					return
																																																				}, 0)

																																																				tmp13044 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3821, tmp13042, V4317, tmp13043)

																																																				__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13044)
																																																				return

																																																			}, 1)

																																																			tmp13045 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																																			__e.TailApply(tmp13041, tmp13045)
																																																			return

																																																		} else {
																																																			__e.Return(False)
																																																			return
																																																		}

																																																	}

																																																}, 1)

																																																tmp13059 := MakeNative(func(__e *ControlFlow) {
																																																	D := __e.Get(1)
																																																	_ = D
																																																	__e.TailApply(GoTo3818, D)
																																																	return
																																																}, 1)

																																																__e.TailApply(tmp13038, tmp13059)
																																																return

																																															}, 1)

																																															tmp13060 := PrimTail(Tm3817)

																																															tmp13061 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13060, V4317)

																																															__e.TailApply(tmp13037, tmp13061)
																																															return

																																														}, 0)

																																														__e.TailApply(tmp13031, tmp13036)
																																														return

																																													}, 1)

																																													tmp13062 := PrimHead(Tm3817)

																																													tmp13063 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13062, V4317)

																																													__e.TailApply(tmp13030, tmp13063)
																																													return

																																												} else {
																																													tmp13029 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3817)

																																													if True == tmp13029 {
																																														tmp13023 := MakeNative(func(__e *ControlFlow) {
																																															D := __e.Get(1)
																																															_ = D
																																															tmp13024 := PrimCons(D, Nil)

																																															tmp13025 := PrimCons(symstream, tmp13024)

																																															tmp13026 := MakeNative(func(__e *ControlFlow) {
																																																__e.TailApply(GoTo3818, D)
																																																return
																																															}, 0)

																																															tmp13027 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3817, tmp13025, V4317, tmp13026)

																																															__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13027)
																																															return

																																														}, 1)

																																														tmp13028 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																														__e.TailApply(tmp13023, tmp13028)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp13065 := MakeNative(func(__e *ControlFlow) {
																																												D := __e.Get(1)
																																												_ = D
																																												tmp13066 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																												_ = tmp13066

																																												tmp13067 := MakeNative(func(__e *ControlFlow) {
																																													tmp13068 := Call(__e, PrimNS2Value(symshen_4lazyderef), D, V4317)

																																													tmp13069 := PrimCons(symout, Nil)

																																													tmp13070 := PrimCons(symin, tmp13069)

																																													tmp13071 := Call(__e, PrimNS2Value(symelement_2), tmp13068, tmp13070)

																																													tmp13072 := MakeNative(func(__e *ControlFlow) {
																																														__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), File, symstring, V4316, V4317, V4318, K3712, V4320)
																																														return
																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symwhen), tmp13071, V4317, V4318, K3712, tmp13072)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symis_b), V3709, D, V4317, V4318, K3712, tmp13067)
																																												return

																																											}, 1)

																																											__e.TailApply(tmp13020, tmp13065)
																																											return

																																										}, 1)

																																										tmp13073 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4315, V4317)

																																										__e.TailApply(tmp13019, tmp13073)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp13075 := PrimTail(Tm3815)

																																								tmp13076 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13075, V4317)

																																								__e.TailApply(tmp13017, tmp13076)
																																								return

																																							}, 1)

																																							tmp13077 := PrimHead(Tm3815)

																																							__e.TailApply(tmp13016, tmp13077)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp13079 := PrimTail(Tm3814)

																																					tmp13080 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13079, V4317)

																																					__e.TailApply(tmp13014, tmp13080)
																																					return

																																				}, 1)

																																				tmp13081 := PrimHead(Tm3814)

																																				__e.TailApply(tmp13013, tmp13081)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp13083 := PrimTail(Tm3812)

																																		tmp13084 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13083, V4317)

																																		__e.TailApply(tmp13011, tmp13084)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13086 := PrimHead(Tm3812)

																																tmp13087 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13086, V4317)

																																__e.TailApply(tmp13009, tmp13087)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13089 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																														tmp13090 := Call(__e, tmp13007, tmp13089)

																														ifres13006 = tmp13090

																													} else {
																														ifres13006 = False

																													}

																													__e.TailApply(tmp12861, ifres13006)
																													return

																												} else {
																													__e.Return(C3804)
																													return
																												}

																											}, 1)

																											tmp13154 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																											var ifres13093 Obj

																											if True == tmp13154 {
																												tmp13094 := MakeNative(func(__e *ControlFlow) {
																													Tm3805 := __e.Get(1)
																													_ = Tm3805
																													tmp13151 := PrimIsPair(Tm3805)

																													if True == tmp13151 {
																														tmp13096 := MakeNative(func(__e *ControlFlow) {
																															Tm3806 := __e.Get(1)
																															_ = Tm3806
																															tmp13148 := PrimEqual(Tm3806, symlet)

																															if True == tmp13148 {
																																tmp13098 := MakeNative(func(__e *ControlFlow) {
																																	Tm3807 := __e.Get(1)
																																	_ = Tm3807
																																	tmp13145 := PrimIsPair(Tm3807)

																																	if True == tmp13145 {
																																		tmp13100 := MakeNative(func(__e *ControlFlow) {
																																			X := __e.Get(1)
																																			_ = X
																																			tmp13101 := MakeNative(func(__e *ControlFlow) {
																																				Tm3808 := __e.Get(1)
																																				_ = Tm3808
																																				tmp13141 := PrimIsPair(Tm3808)

																																				if True == tmp13141 {
																																					tmp13103 := MakeNative(func(__e *ControlFlow) {
																																						Y := __e.Get(1)
																																						_ = Y
																																						tmp13104 := MakeNative(func(__e *ControlFlow) {
																																							Tm3809 := __e.Get(1)
																																							_ = Tm3809
																																							tmp13137 := PrimIsPair(Tm3809)

																																							if True == tmp13137 {
																																								tmp13106 := MakeNative(func(__e *ControlFlow) {
																																									Z := __e.Get(1)
																																									_ = Z
																																									tmp13107 := MakeNative(func(__e *ControlFlow) {
																																										Tm3810 := __e.Get(1)
																																										_ = Tm3810
																																										tmp13133 := PrimEqual(Tm3810, Nil)

																																										if True == tmp13133 {
																																											tmp13109 := MakeNative(func(__e *ControlFlow) {
																																												W := __e.Get(1)
																																												_ = W
																																												tmp13110 := MakeNative(func(__e *ControlFlow) {
																																													New := __e.Get(1)
																																													_ = New
																																													tmp13111 := MakeNative(func(__e *ControlFlow) {
																																														B := __e.Get(1)
																																														_ = B
																																														tmp13112 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																														_ = tmp13112

																																														tmp13113 := MakeNative(func(__e *ControlFlow) {
																																															tmp13114 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V4317)

																																															tmp13115 := Call(__e, PrimNS2Value(symshen_4freshterm), tmp13114)

																																															tmp13116 := MakeNative(func(__e *ControlFlow) {
																																																tmp13117 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V4317)

																																																tmp13118 := Call(__e, PrimNS2Value(symshen_4lazyderef), New, V4317)

																																																tmp13119 := Call(__e, PrimNS2Value(symshen_4lazyderef), Z, V4317)

																																																tmp13120 := Call(__e, PrimNS2Value(symshen_4beta), tmp13117, tmp13118, tmp13119)

																																																tmp13121 := MakeNative(func(__e *ControlFlow) {
																																																	tmp13122 := PrimIntern(MakeString(":"))

																																																	tmp13123 := PrimCons(B, Nil)

																																																	tmp13124 := PrimCons(tmp13122, tmp13123)

																																																	tmp13125 := PrimCons(New, tmp13124)

																																																	tmp13126 := PrimCons(tmp13125, V4316)

																																																	__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), W, V4315, tmp13126, V4317, V4318, K3712, V4320)
																																																	return

																																																}, 0)

																																																__e.TailApply(PrimNS2Value(symbind), W, tmp13120, V4317, V4318, K3712, tmp13121)
																																																return

																																															}, 0)

																																															__e.TailApply(PrimNS2Value(symbind), New, tmp13115, V4317, V4318, K3712, tmp13116)
																																															return

																																														}, 0)

																																														tmp13127 := Call(__e, PrimNS2Value(symshen_4system_1S_1h), Y, B, V4316, V4317, V4318, K3712, tmp13113)

																																														__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13127)
																																														return

																																													}, 1)

																																													tmp13128 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																													tmp13129 := Call(__e, tmp13111, tmp13128)

																																													__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13129)
																																													return

																																												}, 1)

																																												tmp13130 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																												tmp13131 := Call(__e, tmp13110, tmp13130)

																																												__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13131)
																																												return

																																											}, 1)

																																											tmp13132 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																											__e.TailApply(tmp13109, tmp13132)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp13134 := PrimTail(Tm3809)

																																									tmp13135 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13134, V4317)

																																									__e.TailApply(tmp13107, tmp13135)
																																									return

																																								}, 1)

																																								tmp13136 := PrimHead(Tm3809)

																																								__e.TailApply(tmp13106, tmp13136)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp13138 := PrimTail(Tm3808)

																																						tmp13139 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13138, V4317)

																																						__e.TailApply(tmp13104, tmp13139)
																																						return

																																					}, 1)

																																					tmp13140 := PrimHead(Tm3808)

																																					__e.TailApply(tmp13103, tmp13140)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp13142 := PrimTail(Tm3807)

																																			tmp13143 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13142, V4317)

																																			__e.TailApply(tmp13101, tmp13143)
																																			return

																																		}, 1)

																																		tmp13144 := PrimHead(Tm3807)

																																		__e.TailApply(tmp13100, tmp13144)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13146 := PrimTail(Tm3805)

																																tmp13147 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13146, V4317)

																																__e.TailApply(tmp13098, tmp13147)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13149 := PrimHead(Tm3805)

																														tmp13150 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13149, V4317)

																														__e.TailApply(tmp13096, tmp13150)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13152 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																												tmp13153 := Call(__e, tmp13094, tmp13152)

																												ifres13093 = tmp13153

																											} else {
																												ifres13093 = False

																											}

																											__e.TailApply(tmp12859, ifres13093)
																											return

																										} else {
																											__e.Return(C3788)
																											return
																										}

																									}, 1)

																									tmp13277 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																									var ifres13156 Obj

																									if True == tmp13277 {
																										tmp13157 := MakeNative(func(__e *ControlFlow) {
																											Tm3789 := __e.Get(1)
																											_ = Tm3789
																											tmp13274 := PrimIsPair(Tm3789)

																											if True == tmp13274 {
																												tmp13159 := MakeNative(func(__e *ControlFlow) {
																													Tm3790 := __e.Get(1)
																													_ = Tm3790
																													tmp13271 := PrimEqual(Tm3790, symlambda)

																													if True == tmp13271 {
																														tmp13161 := MakeNative(func(__e *ControlFlow) {
																															Tm3791 := __e.Get(1)
																															_ = Tm3791
																															tmp13268 := PrimIsPair(Tm3791)

																															if True == tmp13268 {
																																tmp13163 := MakeNative(func(__e *ControlFlow) {
																																	X := __e.Get(1)
																																	_ = X
																																	tmp13164 := MakeNative(func(__e *ControlFlow) {
																																		Tm3792 := __e.Get(1)
																																		_ = Tm3792
																																		tmp13264 := PrimIsPair(Tm3792)

																																		if True == tmp13264 {
																																			tmp13166 := MakeNative(func(__e *ControlFlow) {
																																				Y := __e.Get(1)
																																				_ = Y
																																				tmp13167 := MakeNative(func(__e *ControlFlow) {
																																					Tm3793 := __e.Get(1)
																																					_ = Tm3793
																																					tmp13260 := PrimEqual(Tm3793, Nil)

																																					if True == tmp13260 {
																																						tmp13169 := MakeNative(func(__e *ControlFlow) {
																																							Tm3794 := __e.Get(1)
																																							_ = Tm3794
																																							tmp13170 := MakeNative(func(__e *ControlFlow) {
																																								GoTo3795 := __e.Get(1)
																																								_ = GoTo3795
																																								tmp13237 := PrimIsPair(Tm3794)

																																								if True == tmp13237 {
																																									tmp13185 := MakeNative(func(__e *ControlFlow) {
																																										A := __e.Get(1)
																																										_ = A
																																										tmp13186 := MakeNative(func(__e *ControlFlow) {
																																											Tm3796 := __e.Get(1)
																																											_ = Tm3796
																																											tmp13187 := MakeNative(func(__e *ControlFlow) {
																																												GoTo3797 := __e.Get(1)
																																												_ = GoTo3797
																																												tmp13231 := PrimIsPair(Tm3796)

																																												if True == tmp13231 {
																																													tmp13197 := MakeNative(func(__e *ControlFlow) {
																																														Tm3798 := __e.Get(1)
																																														_ = Tm3798
																																														tmp13198 := MakeNative(func(__e *ControlFlow) {
																																															GoTo3799 := __e.Get(1)
																																															_ = GoTo3799
																																															tmp13202 := PrimEqual(Tm3798, sym_1_1_6)

																																															if True == tmp13202 {
																																																__e.TailApply(PrimNS2Value(symthaw), GoTo3799)
																																																return
																																															} else {
																																																tmp13201 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3798)

																																																if True == tmp13201 {
																																																	__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3798, sym_1_1_6, V4317, GoTo3799)
																																																	return
																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp13203 := MakeNative(func(__e *ControlFlow) {
																																															tmp13204 := MakeNative(func(__e *ControlFlow) {
																																																Tm3800 := __e.Get(1)
																																																_ = Tm3800
																																																tmp13205 := MakeNative(func(__e *ControlFlow) {
																																																	GoTo3801 := __e.Get(1)
																																																	_ = GoTo3801
																																																	tmp13225 := PrimIsPair(Tm3800)

																																																	if True == tmp13225 {
																																																		tmp13214 := MakeNative(func(__e *ControlFlow) {
																																																			B := __e.Get(1)
																																																			_ = B
																																																			tmp13215 := MakeNative(func(__e *ControlFlow) {
																																																				Tm3802 := __e.Get(1)
																																																				_ = Tm3802
																																																				tmp13216 := MakeNative(func(__e *ControlFlow) {
																																																					GoTo3803 := __e.Get(1)
																																																					_ = GoTo3803
																																																					tmp13220 := PrimEqual(Tm3802, Nil)

																																																					if True == tmp13220 {
																																																						__e.TailApply(PrimNS2Value(symthaw), GoTo3803)
																																																						return
																																																					} else {
																																																						tmp13219 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3802)

																																																						if True == tmp13219 {
																																																							__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3802, Nil, V4317, GoTo3803)
																																																							return
																																																						} else {
																																																							__e.Return(False)
																																																							return
																																																						}

																																																					}

																																																				}, 1)

																																																				tmp13221 := MakeNative(func(__e *ControlFlow) {
																																																					__e.TailApply(GoTo3801, B)
																																																					return
																																																				}, 0)

																																																				__e.TailApply(tmp13216, tmp13221)
																																																				return

																																																			}, 1)

																																																			tmp13222 := PrimTail(Tm3800)

																																																			tmp13223 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13222, V4317)

																																																			__e.TailApply(tmp13215, tmp13223)
																																																			return

																																																		}, 1)

																																																		tmp13224 := PrimHead(Tm3800)

																																																		__e.TailApply(tmp13214, tmp13224)
																																																		return

																																																	} else {
																																																		tmp13213 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3800)

																																																		if True == tmp13213 {
																																																			tmp13208 := MakeNative(func(__e *ControlFlow) {
																																																				B := __e.Get(1)
																																																				_ = B
																																																				tmp13209 := PrimCons(B, Nil)

																																																				tmp13210 := MakeNative(func(__e *ControlFlow) {
																																																					__e.TailApply(GoTo3801, B)
																																																					return
																																																				}, 0)

																																																				tmp13211 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3800, tmp13209, V4317, tmp13210)

																																																				__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13211)
																																																				return

																																																			}, 1)

																																																			tmp13212 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																																			__e.TailApply(tmp13208, tmp13212)
																																																			return

																																																		} else {
																																																			__e.Return(False)
																																																			return
																																																		}

																																																	}

																																																}, 1)

																																																tmp13226 := MakeNative(func(__e *ControlFlow) {
																																																	B := __e.Get(1)
																																																	_ = B
																																																	__e.TailApply(GoTo3797, B)
																																																	return
																																																}, 1)

																																																__e.TailApply(tmp13205, tmp13226)
																																																return

																																															}, 1)

																																															tmp13227 := PrimTail(Tm3796)

																																															tmp13228 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13227, V4317)

																																															__e.TailApply(tmp13204, tmp13228)
																																															return

																																														}, 0)

																																														__e.TailApply(tmp13198, tmp13203)
																																														return

																																													}, 1)

																																													tmp13229 := PrimHead(Tm3796)

																																													tmp13230 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13229, V4317)

																																													__e.TailApply(tmp13197, tmp13230)
																																													return

																																												} else {
																																													tmp13196 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3796)

																																													if True == tmp13196 {
																																														tmp13190 := MakeNative(func(__e *ControlFlow) {
																																															B := __e.Get(1)
																																															_ = B
																																															tmp13191 := PrimCons(B, Nil)

																																															tmp13192 := PrimCons(sym_1_1_6, tmp13191)

																																															tmp13193 := MakeNative(func(__e *ControlFlow) {
																																																__e.TailApply(GoTo3797, B)
																																																return
																																															}, 0)

																																															tmp13194 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3796, tmp13192, V4317, tmp13193)

																																															__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13194)
																																															return

																																														}, 1)

																																														tmp13195 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																														__e.TailApply(tmp13190, tmp13195)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp13232 := MakeNative(func(__e *ControlFlow) {
																																												B := __e.Get(1)
																																												_ = B
																																												tmp13233 := Call(__e, GoTo3795, A)

																																												__e.TailApply(tmp13233, B)
																																												return

																																											}, 1)

																																											__e.TailApply(tmp13187, tmp13232)
																																											return

																																										}, 1)

																																										tmp13234 := PrimTail(Tm3794)

																																										tmp13235 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13234, V4317)

																																										__e.TailApply(tmp13186, tmp13235)
																																										return

																																									}, 1)

																																									tmp13236 := PrimHead(Tm3794)

																																									__e.TailApply(tmp13185, tmp13236)
																																									return

																																								} else {
																																									tmp13184 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3794)

																																									if True == tmp13184 {
																																										tmp13173 := MakeNative(func(__e *ControlFlow) {
																																											A := __e.Get(1)
																																											_ = A
																																											tmp13174 := MakeNative(func(__e *ControlFlow) {
																																												B := __e.Get(1)
																																												_ = B
																																												tmp13175 := PrimCons(B, Nil)

																																												tmp13176 := PrimCons(sym_1_1_6, tmp13175)

																																												tmp13177 := PrimCons(A, tmp13176)

																																												tmp13178 := MakeNative(func(__e *ControlFlow) {
																																													tmp13179 := Call(__e, GoTo3795, A)

																																													__e.TailApply(tmp13179, B)
																																													return

																																												}, 0)

																																												tmp13180 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3794, tmp13177, V4317, tmp13178)

																																												__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13180)
																																												return

																																											}, 1)

																																											tmp13181 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																											tmp13182 := Call(__e, tmp13174, tmp13181)

																																											__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13182)
																																											return

																																										}, 1)

																																										tmp13183 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																										__e.TailApply(tmp13173, tmp13183)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp13238 := MakeNative(func(__e *ControlFlow) {
																																								A := __e.Get(1)
																																								_ = A
																																								__e.Return(MakeNative(func(__e *ControlFlow) {
																																									B := __e.Get(1)
																																									_ = B
																																									tmp13239 := MakeNative(func(__e *ControlFlow) {
																																										Z := __e.Get(1)
																																										_ = Z
																																										tmp13240 := MakeNative(func(__e *ControlFlow) {
																																											New := __e.Get(1)
																																											_ = New
																																											tmp13241 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp13241

																																											tmp13242 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V4317)

																																											tmp13243 := Call(__e, PrimNS2Value(symshen_4freshterm), tmp13242)

																																											tmp13244 := MakeNative(func(__e *ControlFlow) {
																																												tmp13245 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V4317)

																																												tmp13246 := Call(__e, PrimNS2Value(symshen_4deref), New, V4317)

																																												tmp13247 := Call(__e, PrimNS2Value(symshen_4deref), Y, V4317)

																																												tmp13248 := Call(__e, PrimNS2Value(symshen_4beta), tmp13245, tmp13246, tmp13247)

																																												tmp13249 := MakeNative(func(__e *ControlFlow) {
																																													tmp13250 := PrimIntern(MakeString(":"))

																																													tmp13251 := PrimCons(A, Nil)

																																													tmp13252 := PrimCons(tmp13250, tmp13251)

																																													tmp13253 := PrimCons(New, tmp13252)

																																													tmp13254 := PrimCons(tmp13253, V4316)

																																													__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Z, B, tmp13254, V4317, V4318, K3712, V4320)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symbind), Z, tmp13248, V4317, V4318, K3712, tmp13249)
																																												return

																																											}, 0)

																																											tmp13255 := Call(__e, PrimNS2Value(symbind), New, tmp13243, V4317, V4318, K3712, tmp13244)

																																											__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13255)
																																											return

																																										}, 1)

																																										tmp13256 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																										tmp13257 := Call(__e, tmp13240, tmp13256)

																																										__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13257)
																																										return

																																									}, 1)

																																									tmp13258 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																									__e.TailApply(tmp13239, tmp13258)
																																									return

																																								}, 1))
																																								return
																																							}, 1)

																																							__e.TailApply(tmp13170, tmp13238)
																																							return

																																						}, 1)

																																						tmp13259 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4315, V4317)

																																						__e.TailApply(tmp13169, tmp13259)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp13261 := PrimTail(Tm3792)

																																				tmp13262 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13261, V4317)

																																				__e.TailApply(tmp13167, tmp13262)
																																				return

																																			}, 1)

																																			tmp13263 := PrimHead(Tm3792)

																																			__e.TailApply(tmp13166, tmp13263)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp13265 := PrimTail(Tm3791)

																																	tmp13266 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13265, V4317)

																																	__e.TailApply(tmp13164, tmp13266)
																																	return

																																}, 1)

																																tmp13267 := PrimHead(Tm3791)

																																__e.TailApply(tmp13163, tmp13267)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13269 := PrimTail(Tm3789)

																														tmp13270 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13269, V4317)

																														__e.TailApply(tmp13161, tmp13270)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13272 := PrimHead(Tm3789)

																												tmp13273 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13272, V4317)

																												__e.TailApply(tmp13159, tmp13273)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp13275 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																										tmp13276 := Call(__e, tmp13157, tmp13275)

																										ifres13156 = tmp13276

																									} else {
																										ifres13156 = False

																									}

																									__e.TailApply(tmp12857, ifres13156)
																									return

																								} else {
																									__e.Return(C3780)
																									return
																								}

																							}, 1)

																							tmp13319 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																							var ifres13279 Obj

																							if True == tmp13319 {
																								tmp13280 := MakeNative(func(__e *ControlFlow) {
																									Tm3781 := __e.Get(1)
																									_ = Tm3781
																									tmp13316 := PrimIsPair(Tm3781)

																									if True == tmp13316 {
																										tmp13282 := MakeNative(func(__e *ControlFlow) {
																											Tm3782 := __e.Get(1)
																											_ = Tm3782
																											tmp13313 := PrimEqual(Tm3782, sym_8s)

																											if True == tmp13313 {
																												tmp13284 := MakeNative(func(__e *ControlFlow) {
																													Tm3783 := __e.Get(1)
																													_ = Tm3783
																													tmp13310 := PrimIsPair(Tm3783)

																													if True == tmp13310 {
																														tmp13286 := MakeNative(func(__e *ControlFlow) {
																															X := __e.Get(1)
																															_ = X
																															tmp13287 := MakeNative(func(__e *ControlFlow) {
																																Tm3784 := __e.Get(1)
																																_ = Tm3784
																																tmp13306 := PrimIsPair(Tm3784)

																																if True == tmp13306 {
																																	tmp13289 := MakeNative(func(__e *ControlFlow) {
																																		Y := __e.Get(1)
																																		_ = Y
																																		tmp13290 := MakeNative(func(__e *ControlFlow) {
																																			Tm3785 := __e.Get(1)
																																			_ = Tm3785
																																			tmp13302 := PrimEqual(Tm3785, Nil)

																																			if True == tmp13302 {
																																				tmp13292 := MakeNative(func(__e *ControlFlow) {
																																					Tm3786 := __e.Get(1)
																																					_ = Tm3786
																																					tmp13293 := MakeNative(func(__e *ControlFlow) {
																																						GoTo3787 := __e.Get(1)
																																						_ = GoTo3787
																																						tmp13297 := PrimEqual(Tm3786, symstring)

																																						if True == tmp13297 {
																																							__e.TailApply(PrimNS2Value(symthaw), GoTo3787)
																																							return
																																						} else {
																																							tmp13296 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3786)

																																							if True == tmp13296 {
																																								__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3786, symstring, V4317, GoTo3787)
																																								return
																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp13298 := MakeNative(func(__e *ControlFlow) {
																																						tmp13299 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																						_ = tmp13299

																																						tmp13300 := MakeNative(func(__e *ControlFlow) {
																																							__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Y, symstring, V4316, V4317, V4318, K3712, V4320)
																																							return
																																						}, 0)

																																						__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, symstring, V4316, V4317, V4318, K3712, tmp13300)
																																						return

																																					}, 0)

																																					__e.TailApply(tmp13293, tmp13298)
																																					return

																																				}, 1)

																																				tmp13301 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4315, V4317)

																																				__e.TailApply(tmp13292, tmp13301)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp13303 := PrimTail(Tm3784)

																																		tmp13304 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13303, V4317)

																																		__e.TailApply(tmp13290, tmp13304)
																																		return

																																	}, 1)

																																	tmp13305 := PrimHead(Tm3784)

																																	__e.TailApply(tmp13289, tmp13305)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp13307 := PrimTail(Tm3783)

																															tmp13308 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13307, V4317)

																															__e.TailApply(tmp13287, tmp13308)
																															return

																														}, 1)

																														tmp13309 := PrimHead(Tm3783)

																														__e.TailApply(tmp13286, tmp13309)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13311 := PrimTail(Tm3781)

																												tmp13312 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13311, V4317)

																												__e.TailApply(tmp13284, tmp13312)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp13314 := PrimHead(Tm3781)

																										tmp13315 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13314, V4317)

																										__e.TailApply(tmp13282, tmp13315)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13317 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																								tmp13318 := Call(__e, tmp13280, tmp13317)

																								ifres13279 = tmp13318

																							} else {
																								ifres13279 = False

																							}

																							__e.TailApply(tmp12855, ifres13279)
																							return

																						} else {
																							__e.Return(C3766)
																							return
																						}

																					}, 1)

																					tmp13403 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																					var ifres13321 Obj

																					if True == tmp13403 {
																						tmp13322 := MakeNative(func(__e *ControlFlow) {
																							Tm3767 := __e.Get(1)
																							_ = Tm3767
																							tmp13400 := PrimIsPair(Tm3767)

																							if True == tmp13400 {
																								tmp13324 := MakeNative(func(__e *ControlFlow) {
																									Tm3768 := __e.Get(1)
																									_ = Tm3768
																									tmp13397 := PrimEqual(Tm3768, sym_8v)

																									if True == tmp13397 {
																										tmp13326 := MakeNative(func(__e *ControlFlow) {
																											Tm3769 := __e.Get(1)
																											_ = Tm3769
																											tmp13394 := PrimIsPair(Tm3769)

																											if True == tmp13394 {
																												tmp13328 := MakeNative(func(__e *ControlFlow) {
																													X := __e.Get(1)
																													_ = X
																													tmp13329 := MakeNative(func(__e *ControlFlow) {
																														Tm3770 := __e.Get(1)
																														_ = Tm3770
																														tmp13390 := PrimIsPair(Tm3770)

																														if True == tmp13390 {
																															tmp13331 := MakeNative(func(__e *ControlFlow) {
																																Y := __e.Get(1)
																																_ = Y
																																tmp13332 := MakeNative(func(__e *ControlFlow) {
																																	Tm3771 := __e.Get(1)
																																	_ = Tm3771
																																	tmp13386 := PrimEqual(Tm3771, Nil)

																																	if True == tmp13386 {
																																		tmp13334 := MakeNative(func(__e *ControlFlow) {
																																			Tm3772 := __e.Get(1)
																																			_ = Tm3772
																																			tmp13335 := MakeNative(func(__e *ControlFlow) {
																																				GoTo3773 := __e.Get(1)
																																				_ = GoTo3773
																																				tmp13379 := PrimIsPair(Tm3772)

																																				if True == tmp13379 {
																																					tmp13345 := MakeNative(func(__e *ControlFlow) {
																																						Tm3774 := __e.Get(1)
																																						_ = Tm3774
																																						tmp13346 := MakeNative(func(__e *ControlFlow) {
																																							GoTo3775 := __e.Get(1)
																																							_ = GoTo3775
																																							tmp13350 := PrimEqual(Tm3774, symvector)

																																							if True == tmp13350 {
																																								__e.TailApply(PrimNS2Value(symthaw), GoTo3775)
																																								return
																																							} else {
																																								tmp13349 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3774)

																																								if True == tmp13349 {
																																									__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3774, symvector, V4317, GoTo3775)
																																									return
																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp13351 := MakeNative(func(__e *ControlFlow) {
																																							tmp13352 := MakeNative(func(__e *ControlFlow) {
																																								Tm3776 := __e.Get(1)
																																								_ = Tm3776
																																								tmp13353 := MakeNative(func(__e *ControlFlow) {
																																									GoTo3777 := __e.Get(1)
																																									_ = GoTo3777
																																									tmp13373 := PrimIsPair(Tm3776)

																																									if True == tmp13373 {
																																										tmp13362 := MakeNative(func(__e *ControlFlow) {
																																											A := __e.Get(1)
																																											_ = A
																																											tmp13363 := MakeNative(func(__e *ControlFlow) {
																																												Tm3778 := __e.Get(1)
																																												_ = Tm3778
																																												tmp13364 := MakeNative(func(__e *ControlFlow) {
																																													GoTo3779 := __e.Get(1)
																																													_ = GoTo3779
																																													tmp13368 := PrimEqual(Tm3778, Nil)

																																													if True == tmp13368 {
																																														__e.TailApply(PrimNS2Value(symthaw), GoTo3779)
																																														return
																																													} else {
																																														tmp13367 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3778)

																																														if True == tmp13367 {
																																															__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3778, Nil, V4317, GoTo3779)
																																															return
																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}

																																												}, 1)

																																												tmp13369 := MakeNative(func(__e *ControlFlow) {
																																													__e.TailApply(GoTo3777, A)
																																													return
																																												}, 0)

																																												__e.TailApply(tmp13364, tmp13369)
																																												return

																																											}, 1)

																																											tmp13370 := PrimTail(Tm3776)

																																											tmp13371 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13370, V4317)

																																											__e.TailApply(tmp13363, tmp13371)
																																											return

																																										}, 1)

																																										tmp13372 := PrimHead(Tm3776)

																																										__e.TailApply(tmp13362, tmp13372)
																																										return

																																									} else {
																																										tmp13361 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3776)

																																										if True == tmp13361 {
																																											tmp13356 := MakeNative(func(__e *ControlFlow) {
																																												A := __e.Get(1)
																																												_ = A
																																												tmp13357 := PrimCons(A, Nil)

																																												tmp13358 := MakeNative(func(__e *ControlFlow) {
																																													__e.TailApply(GoTo3777, A)
																																													return
																																												}, 0)

																																												tmp13359 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3776, tmp13357, V4317, tmp13358)

																																												__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13359)
																																												return

																																											}, 1)

																																											tmp13360 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																											__e.TailApply(tmp13356, tmp13360)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp13374 := MakeNative(func(__e *ControlFlow) {
																																									A := __e.Get(1)
																																									_ = A
																																									__e.TailApply(GoTo3773, A)
																																									return
																																								}, 1)

																																								__e.TailApply(tmp13353, tmp13374)
																																								return

																																							}, 1)

																																							tmp13375 := PrimTail(Tm3772)

																																							tmp13376 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13375, V4317)

																																							__e.TailApply(tmp13352, tmp13376)
																																							return

																																						}, 0)

																																						__e.TailApply(tmp13346, tmp13351)
																																						return

																																					}, 1)

																																					tmp13377 := PrimHead(Tm3772)

																																					tmp13378 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13377, V4317)

																																					__e.TailApply(tmp13345, tmp13378)
																																					return

																																				} else {
																																					tmp13344 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3772)

																																					if True == tmp13344 {
																																						tmp13338 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp13339 := PrimCons(A, Nil)

																																							tmp13340 := PrimCons(symvector, tmp13339)

																																							tmp13341 := MakeNative(func(__e *ControlFlow) {
																																								__e.TailApply(GoTo3773, A)
																																								return
																																							}, 0)

																																							tmp13342 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3772, tmp13340, V4317, tmp13341)

																																							__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13342)
																																							return

																																						}, 1)

																																						tmp13343 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																						__e.TailApply(tmp13338, tmp13343)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp13380 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp13381 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp13381

																																				tmp13382 := MakeNative(func(__e *ControlFlow) {
																																					tmp13383 := PrimCons(A, Nil)

																																					tmp13384 := PrimCons(symvector, tmp13383)

																																					__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Y, tmp13384, V4316, V4317, V4318, K3712, V4320)
																																					return

																																				}, 0)

																																				__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, A, V4316, V4317, V4318, K3712, tmp13382)
																																				return

																																			}, 1)

																																			__e.TailApply(tmp13335, tmp13380)
																																			return

																																		}, 1)

																																		tmp13385 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4315, V4317)

																																		__e.TailApply(tmp13334, tmp13385)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13387 := PrimTail(Tm3770)

																																tmp13388 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13387, V4317)

																																__e.TailApply(tmp13332, tmp13388)
																																return

																															}, 1)

																															tmp13389 := PrimHead(Tm3770)

																															__e.TailApply(tmp13331, tmp13389)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp13391 := PrimTail(Tm3769)

																													tmp13392 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13391, V4317)

																													__e.TailApply(tmp13329, tmp13392)
																													return

																												}, 1)

																												tmp13393 := PrimHead(Tm3769)

																												__e.TailApply(tmp13328, tmp13393)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp13395 := PrimTail(Tm3767)

																										tmp13396 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13395, V4317)

																										__e.TailApply(tmp13326, tmp13396)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13398 := PrimHead(Tm3767)

																								tmp13399 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13398, V4317)

																								__e.TailApply(tmp13324, tmp13399)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13401 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																						tmp13402 := Call(__e, tmp13322, tmp13401)

																						ifres13321 = tmp13402

																					} else {
																						ifres13321 = False

																					}

																					__e.TailApply(tmp12853, ifres13321)
																					return

																				} else {
																					__e.Return(C3750)
																					return
																				}

																			}, 1)

																			tmp13508 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																			var ifres13405 Obj

																			if True == tmp13508 {
																				tmp13406 := MakeNative(func(__e *ControlFlow) {
																					Tm3751 := __e.Get(1)
																					_ = Tm3751
																					tmp13505 := PrimIsPair(Tm3751)

																					if True == tmp13505 {
																						tmp13408 := MakeNative(func(__e *ControlFlow) {
																							Tm3752 := __e.Get(1)
																							_ = Tm3752
																							tmp13502 := PrimEqual(Tm3752, sym_8p)

																							if True == tmp13502 {
																								tmp13410 := MakeNative(func(__e *ControlFlow) {
																									Tm3753 := __e.Get(1)
																									_ = Tm3753
																									tmp13499 := PrimIsPair(Tm3753)

																									if True == tmp13499 {
																										tmp13412 := MakeNative(func(__e *ControlFlow) {
																											X := __e.Get(1)
																											_ = X
																											tmp13413 := MakeNative(func(__e *ControlFlow) {
																												Tm3754 := __e.Get(1)
																												_ = Tm3754
																												tmp13495 := PrimIsPair(Tm3754)

																												if True == tmp13495 {
																													tmp13415 := MakeNative(func(__e *ControlFlow) {
																														Y := __e.Get(1)
																														_ = Y
																														tmp13416 := MakeNative(func(__e *ControlFlow) {
																															Tm3755 := __e.Get(1)
																															_ = Tm3755
																															tmp13491 := PrimEqual(Tm3755, Nil)

																															if True == tmp13491 {
																																tmp13418 := MakeNative(func(__e *ControlFlow) {
																																	Tm3756 := __e.Get(1)
																																	_ = Tm3756
																																	tmp13419 := MakeNative(func(__e *ControlFlow) {
																																		GoTo3757 := __e.Get(1)
																																		_ = GoTo3757
																																		tmp13486 := PrimIsPair(Tm3756)

																																		if True == tmp13486 {
																																			tmp13434 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp13435 := MakeNative(func(__e *ControlFlow) {
																																					Tm3758 := __e.Get(1)
																																					_ = Tm3758
																																					tmp13436 := MakeNative(func(__e *ControlFlow) {
																																						GoTo3759 := __e.Get(1)
																																						_ = GoTo3759
																																						tmp13480 := PrimIsPair(Tm3758)

																																						if True == tmp13480 {
																																							tmp13446 := MakeNative(func(__e *ControlFlow) {
																																								Tm3760 := __e.Get(1)
																																								_ = Tm3760
																																								tmp13447 := MakeNative(func(__e *ControlFlow) {
																																									GoTo3761 := __e.Get(1)
																																									_ = GoTo3761
																																									tmp13451 := PrimEqual(Tm3760, sym_d)

																																									if True == tmp13451 {
																																										__e.TailApply(PrimNS2Value(symthaw), GoTo3761)
																																										return
																																									} else {
																																										tmp13450 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3760)

																																										if True == tmp13450 {
																																											__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3760, sym_d, V4317, GoTo3761)
																																											return
																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp13452 := MakeNative(func(__e *ControlFlow) {
																																									tmp13453 := MakeNative(func(__e *ControlFlow) {
																																										Tm3762 := __e.Get(1)
																																										_ = Tm3762
																																										tmp13454 := MakeNative(func(__e *ControlFlow) {
																																											GoTo3763 := __e.Get(1)
																																											_ = GoTo3763
																																											tmp13474 := PrimIsPair(Tm3762)

																																											if True == tmp13474 {
																																												tmp13463 := MakeNative(func(__e *ControlFlow) {
																																													B := __e.Get(1)
																																													_ = B
																																													tmp13464 := MakeNative(func(__e *ControlFlow) {
																																														Tm3764 := __e.Get(1)
																																														_ = Tm3764
																																														tmp13465 := MakeNative(func(__e *ControlFlow) {
																																															GoTo3765 := __e.Get(1)
																																															_ = GoTo3765
																																															tmp13469 := PrimEqual(Tm3764, Nil)

																																															if True == tmp13469 {
																																																__e.TailApply(PrimNS2Value(symthaw), GoTo3765)
																																																return
																																															} else {
																																																tmp13468 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3764)

																																																if True == tmp13468 {
																																																	__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3764, Nil, V4317, GoTo3765)
																																																	return
																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp13470 := MakeNative(func(__e *ControlFlow) {
																																															__e.TailApply(GoTo3763, B)
																																															return
																																														}, 0)

																																														__e.TailApply(tmp13465, tmp13470)
																																														return

																																													}, 1)

																																													tmp13471 := PrimTail(Tm3762)

																																													tmp13472 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13471, V4317)

																																													__e.TailApply(tmp13464, tmp13472)
																																													return

																																												}, 1)

																																												tmp13473 := PrimHead(Tm3762)

																																												__e.TailApply(tmp13463, tmp13473)
																																												return

																																											} else {
																																												tmp13462 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3762)

																																												if True == tmp13462 {
																																													tmp13457 := MakeNative(func(__e *ControlFlow) {
																																														B := __e.Get(1)
																																														_ = B
																																														tmp13458 := PrimCons(B, Nil)

																																														tmp13459 := MakeNative(func(__e *ControlFlow) {
																																															__e.TailApply(GoTo3763, B)
																																															return
																																														}, 0)

																																														tmp13460 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3762, tmp13458, V4317, tmp13459)

																																														__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13460)
																																														return

																																													}, 1)

																																													tmp13461 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																													__e.TailApply(tmp13457, tmp13461)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}

																																										}, 1)

																																										tmp13475 := MakeNative(func(__e *ControlFlow) {
																																											B := __e.Get(1)
																																											_ = B
																																											__e.TailApply(GoTo3759, B)
																																											return
																																										}, 1)

																																										__e.TailApply(tmp13454, tmp13475)
																																										return

																																									}, 1)

																																									tmp13476 := PrimTail(Tm3758)

																																									tmp13477 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13476, V4317)

																																									__e.TailApply(tmp13453, tmp13477)
																																									return

																																								}, 0)

																																								__e.TailApply(tmp13447, tmp13452)
																																								return

																																							}, 1)

																																							tmp13478 := PrimHead(Tm3758)

																																							tmp13479 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13478, V4317)

																																							__e.TailApply(tmp13446, tmp13479)
																																							return

																																						} else {
																																							tmp13445 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3758)

																																							if True == tmp13445 {
																																								tmp13439 := MakeNative(func(__e *ControlFlow) {
																																									B := __e.Get(1)
																																									_ = B
																																									tmp13440 := PrimCons(B, Nil)

																																									tmp13441 := PrimCons(sym_d, tmp13440)

																																									tmp13442 := MakeNative(func(__e *ControlFlow) {
																																										__e.TailApply(GoTo3759, B)
																																										return
																																									}, 0)

																																									tmp13443 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3758, tmp13441, V4317, tmp13442)

																																									__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13443)
																																									return

																																								}, 1)

																																								tmp13444 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																								__e.TailApply(tmp13439, tmp13444)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp13481 := MakeNative(func(__e *ControlFlow) {
																																						B := __e.Get(1)
																																						_ = B
																																						tmp13482 := Call(__e, GoTo3757, A)

																																						__e.TailApply(tmp13482, B)
																																						return

																																					}, 1)

																																					__e.TailApply(tmp13436, tmp13481)
																																					return

																																				}, 1)

																																				tmp13483 := PrimTail(Tm3756)

																																				tmp13484 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13483, V4317)

																																				__e.TailApply(tmp13435, tmp13484)
																																				return

																																			}, 1)

																																			tmp13485 := PrimHead(Tm3756)

																																			__e.TailApply(tmp13434, tmp13485)
																																			return

																																		} else {
																																			tmp13433 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3756)

																																			if True == tmp13433 {
																																				tmp13422 := MakeNative(func(__e *ControlFlow) {
																																					A := __e.Get(1)
																																					_ = A
																																					tmp13423 := MakeNative(func(__e *ControlFlow) {
																																						B := __e.Get(1)
																																						_ = B
																																						tmp13424 := PrimCons(B, Nil)

																																						tmp13425 := PrimCons(sym_d, tmp13424)

																																						tmp13426 := PrimCons(A, tmp13425)

																																						tmp13427 := MakeNative(func(__e *ControlFlow) {
																																							tmp13428 := Call(__e, GoTo3757, A)

																																							__e.TailApply(tmp13428, B)
																																							return

																																						}, 0)

																																						tmp13429 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3756, tmp13426, V4317, tmp13427)

																																						__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13429)
																																						return

																																					}, 1)

																																					tmp13430 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																					tmp13431 := Call(__e, tmp13423, tmp13430)

																																					__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13431)
																																					return

																																				}, 1)

																																				tmp13432 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																				__e.TailApply(tmp13422, tmp13432)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp13487 := MakeNative(func(__e *ControlFlow) {
																																		A := __e.Get(1)
																																		_ = A
																																		__e.Return(MakeNative(func(__e *ControlFlow) {
																																			B := __e.Get(1)
																																			_ = B
																																			tmp13488 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp13488

																																			tmp13489 := MakeNative(func(__e *ControlFlow) {
																																				__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Y, B, V4316, V4317, V4318, K3712, V4320)
																																				return
																																			}, 0)

																																			__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, A, V4316, V4317, V4318, K3712, tmp13489)
																																			return

																																		}, 1))
																																		return
																																	}, 1)

																																	__e.TailApply(tmp13419, tmp13487)
																																	return

																																}, 1)

																																tmp13490 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4315, V4317)

																																__e.TailApply(tmp13418, tmp13490)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13492 := PrimTail(Tm3754)

																														tmp13493 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13492, V4317)

																														__e.TailApply(tmp13416, tmp13493)
																														return

																													}, 1)

																													tmp13494 := PrimHead(Tm3754)

																													__e.TailApply(tmp13415, tmp13494)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp13496 := PrimTail(Tm3753)

																											tmp13497 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13496, V4317)

																											__e.TailApply(tmp13413, tmp13497)
																											return

																										}, 1)

																										tmp13498 := PrimHead(Tm3753)

																										__e.TailApply(tmp13412, tmp13498)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13500 := PrimTail(Tm3751)

																								tmp13501 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13500, V4317)

																								__e.TailApply(tmp13410, tmp13501)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13503 := PrimHead(Tm3751)

																						tmp13504 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13503, V4317)

																						__e.TailApply(tmp13408, tmp13504)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp13506 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																				tmp13507 := Call(__e, tmp13406, tmp13506)

																				ifres13405 = tmp13507

																			} else {
																				ifres13405 = False

																			}

																			__e.TailApply(tmp12851, ifres13405)
																			return

																		} else {
																			__e.Return(C3736)
																			return
																		}

																	}, 1)

																	tmp13592 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

																	var ifres13510 Obj

																	if True == tmp13592 {
																		tmp13511 := MakeNative(func(__e *ControlFlow) {
																			Tm3737 := __e.Get(1)
																			_ = Tm3737
																			tmp13589 := PrimIsPair(Tm3737)

																			if True == tmp13589 {
																				tmp13513 := MakeNative(func(__e *ControlFlow) {
																					Tm3738 := __e.Get(1)
																					_ = Tm3738
																					tmp13586 := PrimEqual(Tm3738, symcons)

																					if True == tmp13586 {
																						tmp13515 := MakeNative(func(__e *ControlFlow) {
																							Tm3739 := __e.Get(1)
																							_ = Tm3739
																							tmp13583 := PrimIsPair(Tm3739)

																							if True == tmp13583 {
																								tmp13517 := MakeNative(func(__e *ControlFlow) {
																									X := __e.Get(1)
																									_ = X
																									tmp13518 := MakeNative(func(__e *ControlFlow) {
																										Tm3740 := __e.Get(1)
																										_ = Tm3740
																										tmp13579 := PrimIsPair(Tm3740)

																										if True == tmp13579 {
																											tmp13520 := MakeNative(func(__e *ControlFlow) {
																												Y := __e.Get(1)
																												_ = Y
																												tmp13521 := MakeNative(func(__e *ControlFlow) {
																													Tm3741 := __e.Get(1)
																													_ = Tm3741
																													tmp13575 := PrimEqual(Tm3741, Nil)

																													if True == tmp13575 {
																														tmp13523 := MakeNative(func(__e *ControlFlow) {
																															Tm3742 := __e.Get(1)
																															_ = Tm3742
																															tmp13524 := MakeNative(func(__e *ControlFlow) {
																																GoTo3743 := __e.Get(1)
																																_ = GoTo3743
																																tmp13568 := PrimIsPair(Tm3742)

																																if True == tmp13568 {
																																	tmp13534 := MakeNative(func(__e *ControlFlow) {
																																		Tm3744 := __e.Get(1)
																																		_ = Tm3744
																																		tmp13535 := MakeNative(func(__e *ControlFlow) {
																																			GoTo3745 := __e.Get(1)
																																			_ = GoTo3745
																																			tmp13539 := PrimEqual(Tm3744, symlist)

																																			if True == tmp13539 {
																																				__e.TailApply(PrimNS2Value(symthaw), GoTo3745)
																																				return
																																			} else {
																																				tmp13538 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3744)

																																				if True == tmp13538 {
																																					__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3744, symlist, V4317, GoTo3745)
																																					return
																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp13540 := MakeNative(func(__e *ControlFlow) {
																																			tmp13541 := MakeNative(func(__e *ControlFlow) {
																																				Tm3746 := __e.Get(1)
																																				_ = Tm3746
																																				tmp13542 := MakeNative(func(__e *ControlFlow) {
																																					GoTo3747 := __e.Get(1)
																																					_ = GoTo3747
																																					tmp13562 := PrimIsPair(Tm3746)

																																					if True == tmp13562 {
																																						tmp13551 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp13552 := MakeNative(func(__e *ControlFlow) {
																																								Tm3748 := __e.Get(1)
																																								_ = Tm3748
																																								tmp13553 := MakeNative(func(__e *ControlFlow) {
																																									GoTo3749 := __e.Get(1)
																																									_ = GoTo3749
																																									tmp13557 := PrimEqual(Tm3748, Nil)

																																									if True == tmp13557 {
																																										__e.TailApply(PrimNS2Value(symthaw), GoTo3749)
																																										return
																																									} else {
																																										tmp13556 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3748)

																																										if True == tmp13556 {
																																											__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3748, Nil, V4317, GoTo3749)
																																											return
																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp13558 := MakeNative(func(__e *ControlFlow) {
																																									__e.TailApply(GoTo3747, A)
																																									return
																																								}, 0)

																																								__e.TailApply(tmp13553, tmp13558)
																																								return

																																							}, 1)

																																							tmp13559 := PrimTail(Tm3746)

																																							tmp13560 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13559, V4317)

																																							__e.TailApply(tmp13552, tmp13560)
																																							return

																																						}, 1)

																																						tmp13561 := PrimHead(Tm3746)

																																						__e.TailApply(tmp13551, tmp13561)
																																						return

																																					} else {
																																						tmp13550 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3746)

																																						if True == tmp13550 {
																																							tmp13545 := MakeNative(func(__e *ControlFlow) {
																																								A := __e.Get(1)
																																								_ = A
																																								tmp13546 := PrimCons(A, Nil)

																																								tmp13547 := MakeNative(func(__e *ControlFlow) {
																																									__e.TailApply(GoTo3747, A)
																																									return
																																								}, 0)

																																								tmp13548 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3746, tmp13546, V4317, tmp13547)

																																								__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13548)
																																								return

																																							}, 1)

																																							tmp13549 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																							__e.TailApply(tmp13545, tmp13549)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp13563 := MakeNative(func(__e *ControlFlow) {
																																					A := __e.Get(1)
																																					_ = A
																																					__e.TailApply(GoTo3743, A)
																																					return
																																				}, 1)

																																				__e.TailApply(tmp13542, tmp13563)
																																				return

																																			}, 1)

																																			tmp13564 := PrimTail(Tm3742)

																																			tmp13565 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13564, V4317)

																																			__e.TailApply(tmp13541, tmp13565)
																																			return

																																		}, 0)

																																		__e.TailApply(tmp13535, tmp13540)
																																		return

																																	}, 1)

																																	tmp13566 := PrimHead(Tm3742)

																																	tmp13567 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13566, V4317)

																																	__e.TailApply(tmp13534, tmp13567)
																																	return

																																} else {
																																	tmp13533 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3742)

																																	if True == tmp13533 {
																																		tmp13527 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp13528 := PrimCons(A, Nil)

																																			tmp13529 := PrimCons(symlist, tmp13528)

																																			tmp13530 := MakeNative(func(__e *ControlFlow) {
																																				__e.TailApply(GoTo3743, A)
																																				return
																																			}, 0)

																																			tmp13531 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3742, tmp13529, V4317, tmp13530)

																																			__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13531)
																																			return

																																		}, 1)

																																		tmp13532 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																																		__e.TailApply(tmp13527, tmp13532)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp13569 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp13570 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																_ = tmp13570

																																tmp13571 := MakeNative(func(__e *ControlFlow) {
																																	tmp13572 := PrimCons(A, Nil)

																																	tmp13573 := PrimCons(symlist, tmp13572)

																																	__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), Y, tmp13573, V4316, V4317, V4318, K3712, V4320)
																																	return

																																}, 0)

																																__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, A, V4316, V4317, V4318, K3712, tmp13571)
																																return

																															}, 1)

																															__e.TailApply(tmp13524, tmp13569)
																															return

																														}, 1)

																														tmp13574 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4315, V4317)

																														__e.TailApply(tmp13523, tmp13574)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13576 := PrimTail(Tm3740)

																												tmp13577 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13576, V4317)

																												__e.TailApply(tmp13521, tmp13577)
																												return

																											}, 1)

																											tmp13578 := PrimHead(Tm3740)

																											__e.TailApply(tmp13520, tmp13578)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp13580 := PrimTail(Tm3739)

																									tmp13581 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13580, V4317)

																									__e.TailApply(tmp13518, tmp13581)
																									return

																								}, 1)

																								tmp13582 := PrimHead(Tm3739)

																								__e.TailApply(tmp13517, tmp13582)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13584 := PrimTail(Tm3737)

																						tmp13585 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13584, V4317)

																						__e.TailApply(tmp13515, tmp13585)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp13587 := PrimHead(Tm3737)

																				tmp13588 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13587, V4317)

																				__e.TailApply(tmp13513, tmp13588)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp13590 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																		tmp13591 := Call(__e, tmp13511, tmp13590)

																		ifres13510 = tmp13591

																	} else {
																		ifres13510 = False

																	}

																	__e.TailApply(tmp12849, ifres13510)
																	return

																} else {
																	__e.Return(C3732)
																	return
																}

															}, 1)

															tmp13622 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

															var ifres13594 Obj

															if True == tmp13622 {
																tmp13595 := MakeNative(func(__e *ControlFlow) {
																	Tm3733 := __e.Get(1)
																	_ = Tm3733
																	tmp13619 := PrimIsPair(Tm3733)

																	if True == tmp13619 {
																		tmp13597 := MakeNative(func(__e *ControlFlow) {
																			F := __e.Get(1)
																			_ = F
																			tmp13598 := MakeNative(func(__e *ControlFlow) {
																				Tm3734 := __e.Get(1)
																				_ = Tm3734
																				tmp13615 := PrimIsPair(Tm3734)

																				if True == tmp13615 {
																					tmp13600 := MakeNative(func(__e *ControlFlow) {
																						X := __e.Get(1)
																						_ = X
																						tmp13601 := MakeNative(func(__e *ControlFlow) {
																							Tm3735 := __e.Get(1)
																							_ = Tm3735
																							tmp13611 := PrimEqual(Tm3735, Nil)

																							if True == tmp13611 {
																								tmp13603 := MakeNative(func(__e *ControlFlow) {
																									B := __e.Get(1)
																									_ = B
																									tmp13604 := Call(__e, PrimNS2Value(symshen_4incinfs))

																									_ = tmp13604

																									tmp13605 := PrimCons(V4315, Nil)

																									tmp13606 := PrimCons(sym_1_1_6, tmp13605)

																									tmp13607 := PrimCons(B, tmp13606)

																									tmp13608 := MakeNative(func(__e *ControlFlow) {
																										__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, B, V4316, V4317, V4318, K3712, V4320)
																										return
																									}, 0)

																									tmp13609 := Call(__e, PrimNS2Value(symshen_4system_1S_1h), F, tmp13607, V4316, V4317, V4318, K3712, tmp13608)

																									__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13609)
																									return

																								}, 1)

																								tmp13610 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																								__e.TailApply(tmp13603, tmp13610)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13612 := PrimTail(Tm3734)

																						tmp13613 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13612, V4317)

																						__e.TailApply(tmp13601, tmp13613)
																						return

																					}, 1)

																					tmp13614 := PrimHead(Tm3734)

																					__e.TailApply(tmp13600, tmp13614)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp13616 := PrimTail(Tm3733)

																			tmp13617 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13616, V4317)

																			__e.TailApply(tmp13598, tmp13617)
																			return

																		}, 1)

																		tmp13618 := PrimHead(Tm3733)

																		__e.TailApply(tmp13597, tmp13618)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp13620 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

																tmp13621 := Call(__e, tmp13595, tmp13620)

																ifres13594 = tmp13621

															} else {
																ifres13594 = False

															}

															__e.TailApply(tmp12847, ifres13594)
															return

														} else {
															__e.Return(C3728)
															return
														}

													}, 1)

													tmp13656 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

													var ifres13624 Obj

													if True == tmp13656 {
														tmp13625 := MakeNative(func(__e *ControlFlow) {
															Tm3729 := __e.Get(1)
															_ = Tm3729
															tmp13653 := PrimIsPair(Tm3729)

															if True == tmp13653 {
																tmp13627 := MakeNative(func(__e *ControlFlow) {
																	F := __e.Get(1)
																	_ = F
																	tmp13628 := MakeNative(func(__e *ControlFlow) {
																		Tm3730 := __e.Get(1)
																		_ = Tm3730
																		tmp13649 := PrimIsPair(Tm3730)

																		if True == tmp13649 {
																			tmp13630 := MakeNative(func(__e *ControlFlow) {
																				X := __e.Get(1)
																				_ = X
																				tmp13631 := MakeNative(func(__e *ControlFlow) {
																					Tm3731 := __e.Get(1)
																					_ = Tm3731
																					tmp13645 := PrimEqual(Tm3731, Nil)

																					if True == tmp13645 {
																						tmp13633 := MakeNative(func(__e *ControlFlow) {
																							B := __e.Get(1)
																							_ = B
																							tmp13634 := Call(__e, PrimNS2Value(symshen_4incinfs))

																							_ = tmp13634

																							tmp13635 := Call(__e, PrimNS2Value(symshen_4lazyderef), F, V4317)

																							tmp13636 := PrimIsPair(tmp13635)

																							tmp13637 := PrimNot(tmp13636)

																							tmp13638 := MakeNative(func(__e *ControlFlow) {
																								tmp13639 := PrimCons(V4315, Nil)

																								tmp13640 := PrimCons(sym_1_1_6, tmp13639)

																								tmp13641 := PrimCons(B, tmp13640)

																								tmp13642 := MakeNative(func(__e *ControlFlow) {
																									__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), X, B, V4316, V4317, V4318, K3712, V4320)
																									return
																								}, 0)

																								__e.TailApply(PrimNS2Value(symshen_4lookupsig), F, tmp13641, V4317, V4318, K3712, tmp13642)
																								return

																							}, 0)

																							tmp13643 := Call(__e, PrimNS2Value(symwhen), tmp13637, V4317, V4318, K3712, tmp13638)

																							__e.TailApply(PrimNS2Value(symshen_4gc), V4317, tmp13643)
																							return

																						}, 1)

																						tmp13644 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

																						__e.TailApply(tmp13633, tmp13644)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp13646 := PrimTail(Tm3730)

																				tmp13647 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13646, V4317)

																				__e.TailApply(tmp13631, tmp13647)
																				return

																			}, 1)

																			tmp13648 := PrimHead(Tm3730)

																			__e.TailApply(tmp13630, tmp13648)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp13650 := PrimTail(Tm3729)

																	tmp13651 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13650, V4317)

																	__e.TailApply(tmp13628, tmp13651)
																	return

																}, 1)

																tmp13652 := PrimHead(Tm3729)

																__e.TailApply(tmp13627, tmp13652)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp13654 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

														tmp13655 := Call(__e, tmp13625, tmp13654)

														ifres13624 = tmp13655

													} else {
														ifres13624 = False

													}

													__e.TailApply(tmp12845, ifres13624)
													return

												} else {
													__e.Return(C3723)
													return
												}

											}, 1)

											tmp13682 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

											var ifres13658 Obj

											if True == tmp13682 {
												tmp13659 := MakeNative(func(__e *ControlFlow) {
													Tm3724 := __e.Get(1)
													_ = Tm3724
													tmp13679 := PrimIsPair(Tm3724)

													if True == tmp13679 {
														tmp13661 := MakeNative(func(__e *ControlFlow) {
															Tm3725 := __e.Get(1)
															_ = Tm3725
															tmp13676 := PrimEqual(Tm3725, symfn)

															if True == tmp13676 {
																tmp13663 := MakeNative(func(__e *ControlFlow) {
																	Tm3726 := __e.Get(1)
																	_ = Tm3726
																	tmp13673 := PrimIsPair(Tm3726)

																	if True == tmp13673 {
																		tmp13665 := MakeNative(func(__e *ControlFlow) {
																			F := __e.Get(1)
																			_ = F
																			tmp13666 := MakeNative(func(__e *ControlFlow) {
																				Tm3727 := __e.Get(1)
																				_ = Tm3727
																				tmp13669 := PrimEqual(Tm3727, Nil)

																				if True == tmp13669 {
																					tmp13668 := Call(__e, PrimNS2Value(symshen_4incinfs))

																					_ = tmp13668

																					__e.TailApply(PrimNS2Value(symshen_4lookupsig), F, V4315, V4317, V4318, K3712, V4320)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp13670 := PrimTail(Tm3726)

																			tmp13671 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13670, V4317)

																			__e.TailApply(tmp13666, tmp13671)
																			return

																		}, 1)

																		tmp13672 := PrimHead(Tm3726)

																		__e.TailApply(tmp13665, tmp13672)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp13674 := PrimTail(Tm3724)

																tmp13675 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13674, V4317)

																__e.TailApply(tmp13663, tmp13675)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp13677 := PrimHead(Tm3724)

														tmp13678 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13677, V4317)

														__e.TailApply(tmp13661, tmp13678)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp13680 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

												tmp13681 := Call(__e, tmp13659, tmp13680)

												ifres13658 = tmp13681

											} else {
												ifres13658 = False

											}

											__e.TailApply(tmp12843, ifres13658)
											return

										} else {
											__e.Return(C3720)
											return
										}

									}, 1)

									tmp13700 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

									var ifres13684 Obj

									if True == tmp13700 {
										tmp13685 := MakeNative(func(__e *ControlFlow) {
											Tm3721 := __e.Get(1)
											_ = Tm3721
											tmp13697 := PrimIsPair(Tm3721)

											if True == tmp13697 {
												tmp13687 := MakeNative(func(__e *ControlFlow) {
													F := __e.Get(1)
													_ = F
													tmp13688 := MakeNative(func(__e *ControlFlow) {
														Tm3722 := __e.Get(1)
														_ = Tm3722
														tmp13693 := PrimEqual(Tm3722, Nil)

														if True == tmp13693 {
															tmp13690 := Call(__e, PrimNS2Value(symshen_4incinfs))

															_ = tmp13690

															tmp13691 := PrimCons(V4315, Nil)

															tmp13692 := PrimCons(sym_1_1_6, tmp13691)

															__e.TailApply(PrimNS2Value(symshen_4lookupsig), F, tmp13692, V4317, V4318, K3712, V4320)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp13694 := PrimTail(Tm3721)

													tmp13695 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13694, V4317)

													__e.TailApply(tmp13688, tmp13695)
													return

												}, 1)

												tmp13696 := PrimHead(Tm3721)

												__e.TailApply(tmp13687, tmp13696)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp13698 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

										tmp13699 := Call(__e, tmp13685, tmp13698)

										ifres13684 = tmp13699

									} else {
										ifres13684 = False

									}

									__e.TailApply(tmp12841, ifres13684)
									return

								} else {
									__e.Return(C3719)
									return
								}

							}, 1)

							tmp13705 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

							var ifres13702 Obj

							if True == tmp13705 {
								tmp13703 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp13703

								tmp13704 := Call(__e, PrimNS2Value(symshen_4by_1hypothesis), V4314, V4315, V4316, V4317, V4318, K3712, V4320)

								ifres13702 = tmp13704

							} else {
								ifres13702 = False

							}

							__e.TailApply(tmp12839, ifres13702)
							return

						} else {
							__e.Return(C3718)
							return
						}

					}, 1)

					tmp13714 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

					var ifres13707 Obj

					if True == tmp13714 {
						tmp13708 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp13708

						tmp13709 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4314, V4317)

						tmp13710 := PrimIsPair(tmp13709)

						tmp13711 := PrimNot(tmp13710)

						tmp13712 := MakeNative(func(__e *ControlFlow) {
							__e.TailApply(PrimNS2Value(symshen_4primitive), V4314, V4315, V4317, V4318, K3712, V4320)
							return
						}, 0)

						tmp13713 := Call(__e, PrimNS2Value(symwhen), tmp13711, V4317, V4318, K3712, tmp13712)

						ifres13707 = tmp13713

					} else {
						ifres13707 = False

					}

					__e.TailApply(tmp12837, ifres13707)
					return

				} else {
					__e.Return(C3717)
					return
				}

			}, 1)

			tmp13725 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4318)

			var ifres13716 Obj

			if True == tmp13725 {
				tmp13717 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp13717

				tmp13718 := PrimNS3Value(symshen_4_dspy_d)

				tmp13719 := MakeNative(func(__e *ControlFlow) {
					tmp13720 := PrimIntern(MakeString(":"))

					tmp13721 := PrimCons(V4315, Nil)

					tmp13722 := PrimCons(tmp13720, tmp13721)

					tmp13723 := PrimCons(V4314, tmp13722)

					__e.TailApply(PrimNS2Value(symshen_4show), tmp13723, V4316, V4317, V4318, K3712, V4320)
					return

				}, 0)

				tmp13724 := Call(__e, PrimNS2Value(symwhen), tmp13718, V4317, V4318, K3712, tmp13719)

				ifres13716 = tmp13724

			} else {
				ifres13716 = False

			}

			__e.TailApply(tmp12835, ifres13716)
			return

		}, 1)

		tmp13726 := PrimNumberAdd(V4319, MakeNumber(1))

		__e.TailApply(tmp12834, tmp13726)
		return

	}, 7)

	tmp13727 := Call(__e, PrimNS2Value(symdef), symshen_4system_1S_1h, tmp12833)

	_ = tmp13727

	tmp13728 := MakeNative(func(__e *ControlFlow) {
		V4321 := __e.Get(1)
		_ = V4321
		V4322 := __e.Get(2)
		_ = V4322
		V4323 := __e.Get(3)
		_ = V4323
		V4324 := __e.Get(4)
		_ = V4324
		V4325 := __e.Get(5)
		_ = V4325
		V4326 := __e.Get(6)
		_ = V4326
		tmp13729 := MakeNative(func(__e *ControlFlow) {
			C3851 := __e.Get(1)
			_ = C3851
			tmp13837 := PrimEqual(C3851, False)

			if True == tmp13837 {
				tmp13731 := MakeNative(func(__e *ControlFlow) {
					C3854 := __e.Get(1)
					_ = C3854
					tmp13822 := PrimEqual(C3854, False)

					if True == tmp13822 {
						tmp13733 := MakeNative(func(__e *ControlFlow) {
							C3857 := __e.Get(1)
							_ = C3857
							tmp13807 := PrimEqual(C3857, False)

							if True == tmp13807 {
								tmp13735 := MakeNative(func(__e *ControlFlow) {
									C3860 := __e.Get(1)
									_ = C3860
									tmp13792 := PrimEqual(C3860, False)

									if True == tmp13792 {
										tmp13791 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4324)

										if True == tmp13791 {
											tmp13738 := MakeNative(func(__e *ControlFlow) {
												Tm3863 := __e.Get(1)
												_ = Tm3863
												tmp13789 := PrimEqual(Tm3863, Nil)

												if True == tmp13789 {
													tmp13740 := MakeNative(func(__e *ControlFlow) {
														Tm3864 := __e.Get(1)
														_ = Tm3864
														tmp13741 := MakeNative(func(__e *ControlFlow) {
															GoTo3865 := __e.Get(1)
															_ = GoTo3865
															tmp13785 := PrimIsPair(Tm3864)

															if True == tmp13785 {
																tmp13751 := MakeNative(func(__e *ControlFlow) {
																	Tm3866 := __e.Get(1)
																	_ = Tm3866
																	tmp13752 := MakeNative(func(__e *ControlFlow) {
																		GoTo3867 := __e.Get(1)
																		_ = GoTo3867
																		tmp13756 := PrimEqual(Tm3866, symlist)

																		if True == tmp13756 {
																			__e.TailApply(PrimNS2Value(symthaw), GoTo3867)
																			return
																		} else {
																			tmp13755 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3866)

																			if True == tmp13755 {
																				__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3866, symlist, V4323, GoTo3867)
																				return
																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp13757 := MakeNative(func(__e *ControlFlow) {
																		tmp13758 := MakeNative(func(__e *ControlFlow) {
																			Tm3868 := __e.Get(1)
																			_ = Tm3868
																			tmp13759 := MakeNative(func(__e *ControlFlow) {
																				GoTo3869 := __e.Get(1)
																				_ = GoTo3869
																				tmp13779 := PrimIsPair(Tm3868)

																				if True == tmp13779 {
																					tmp13768 := MakeNative(func(__e *ControlFlow) {
																						A := __e.Get(1)
																						_ = A
																						tmp13769 := MakeNative(func(__e *ControlFlow) {
																							Tm3870 := __e.Get(1)
																							_ = Tm3870
																							tmp13770 := MakeNative(func(__e *ControlFlow) {
																								GoTo3871 := __e.Get(1)
																								_ = GoTo3871
																								tmp13774 := PrimEqual(Tm3870, Nil)

																								if True == tmp13774 {
																									__e.TailApply(PrimNS2Value(symthaw), GoTo3871)
																									return
																								} else {
																									tmp13773 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3870)

																									if True == tmp13773 {
																										__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3870, Nil, V4323, GoTo3871)
																										return
																									} else {
																										__e.Return(False)
																										return
																									}

																								}

																							}, 1)

																							tmp13775 := MakeNative(func(__e *ControlFlow) {
																								__e.TailApply(GoTo3869, A)
																								return
																							}, 0)

																							__e.TailApply(tmp13770, tmp13775)
																							return

																						}, 1)

																						tmp13776 := PrimTail(Tm3868)

																						tmp13777 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13776, V4323)

																						__e.TailApply(tmp13769, tmp13777)
																						return

																					}, 1)

																					tmp13778 := PrimHead(Tm3868)

																					__e.TailApply(tmp13768, tmp13778)
																					return

																				} else {
																					tmp13767 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3868)

																					if True == tmp13767 {
																						tmp13762 := MakeNative(func(__e *ControlFlow) {
																							A := __e.Get(1)
																							_ = A
																							tmp13763 := PrimCons(A, Nil)

																							tmp13764 := MakeNative(func(__e *ControlFlow) {
																								__e.TailApply(GoTo3869, A)
																								return
																							}, 0)

																							tmp13765 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3868, tmp13763, V4323, tmp13764)

																							__e.TailApply(PrimNS2Value(symshen_4gc), V4323, tmp13765)
																							return

																						}, 1)

																						tmp13766 := Call(__e, PrimNS2Value(symshen_4newpv), V4323)

																						__e.TailApply(tmp13762, tmp13766)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}

																			}, 1)

																			tmp13780 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				__e.TailApply(GoTo3865, A)
																				return
																			}, 1)

																			__e.TailApply(tmp13759, tmp13780)
																			return

																		}, 1)

																		tmp13781 := PrimTail(Tm3864)

																		tmp13782 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13781, V4323)

																		__e.TailApply(tmp13758, tmp13782)
																		return

																	}, 0)

																	__e.TailApply(tmp13752, tmp13757)
																	return

																}, 1)

																tmp13783 := PrimHead(Tm3864)

																tmp13784 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13783, V4323)

																__e.TailApply(tmp13751, tmp13784)
																return

															} else {
																tmp13750 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3864)

																if True == tmp13750 {
																	tmp13744 := MakeNative(func(__e *ControlFlow) {
																		A := __e.Get(1)
																		_ = A
																		tmp13745 := PrimCons(A, Nil)

																		tmp13746 := PrimCons(symlist, tmp13745)

																		tmp13747 := MakeNative(func(__e *ControlFlow) {
																			__e.TailApply(GoTo3865, A)
																			return
																		}, 0)

																		tmp13748 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3864, tmp13746, V4323, tmp13747)

																		__e.TailApply(PrimNS2Value(symshen_4gc), V4323, tmp13748)
																		return

																	}, 1)

																	tmp13749 := Call(__e, PrimNS2Value(symshen_4newpv), V4323)

																	__e.TailApply(tmp13744, tmp13749)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}

														}, 1)

														tmp13786 := MakeNative(func(__e *ControlFlow) {
															A := __e.Get(1)
															_ = A
															tmp13787 := Call(__e, PrimNS2Value(symshen_4incinfs))

															_ = tmp13787

															__e.TailApply(PrimNS2Value(symthaw), V4326)
															return

														}, 1)

														__e.TailApply(tmp13741, tmp13786)
														return

													}, 1)

													tmp13788 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4322, V4323)

													__e.TailApply(tmp13740, tmp13788)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp13790 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4321, V4323)

											__e.TailApply(tmp13738, tmp13790)
											return

										} else {
											__e.Return(False)
											return
										}

									} else {
										__e.Return(C3860)
										return
									}

								}, 1)

								tmp13806 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4324)

								var ifres13793 Obj

								if True == tmp13806 {
									tmp13794 := MakeNative(func(__e *ControlFlow) {
										Tm3861 := __e.Get(1)
										_ = Tm3861
										tmp13795 := MakeNative(func(__e *ControlFlow) {
											GoTo3862 := __e.Get(1)
											_ = GoTo3862
											tmp13799 := PrimEqual(Tm3861, symsymbol)

											if True == tmp13799 {
												__e.TailApply(PrimNS2Value(symthaw), GoTo3862)
												return
											} else {
												tmp13798 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3861)

												if True == tmp13798 {
													__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3861, symsymbol, V4323, GoTo3862)
													return
												} else {
													__e.Return(False)
													return
												}

											}

										}, 1)

										tmp13800 := MakeNative(func(__e *ControlFlow) {
											tmp13801 := Call(__e, PrimNS2Value(symshen_4incinfs))

											_ = tmp13801

											tmp13802 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4321, V4323)

											tmp13803 := PrimIsSymbol(tmp13802)

											__e.TailApply(PrimNS2Value(symwhen), tmp13803, V4323, V4324, V4325, V4326)
											return

										}, 0)

										__e.TailApply(tmp13795, tmp13800)
										return

									}, 1)

									tmp13804 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4322, V4323)

									tmp13805 := Call(__e, tmp13794, tmp13804)

									ifres13793 = tmp13805

								} else {
									ifres13793 = False

								}

								__e.TailApply(tmp13735, ifres13793)
								return

							} else {
								__e.Return(C3857)
								return
							}

						}, 1)

						tmp13821 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4324)

						var ifres13808 Obj

						if True == tmp13821 {
							tmp13809 := MakeNative(func(__e *ControlFlow) {
								Tm3858 := __e.Get(1)
								_ = Tm3858
								tmp13810 := MakeNative(func(__e *ControlFlow) {
									GoTo3859 := __e.Get(1)
									_ = GoTo3859
									tmp13814 := PrimEqual(Tm3858, symstring)

									if True == tmp13814 {
										__e.TailApply(PrimNS2Value(symthaw), GoTo3859)
										return
									} else {
										tmp13813 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3858)

										if True == tmp13813 {
											__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3858, symstring, V4323, GoTo3859)
											return
										} else {
											__e.Return(False)
											return
										}

									}

								}, 1)

								tmp13815 := MakeNative(func(__e *ControlFlow) {
									tmp13816 := Call(__e, PrimNS2Value(symshen_4incinfs))

									_ = tmp13816

									tmp13817 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4321, V4323)

									tmp13818 := PrimIsString(tmp13817)

									__e.TailApply(PrimNS2Value(symwhen), tmp13818, V4323, V4324, V4325, V4326)
									return

								}, 0)

								__e.TailApply(tmp13810, tmp13815)
								return

							}, 1)

							tmp13819 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4322, V4323)

							tmp13820 := Call(__e, tmp13809, tmp13819)

							ifres13808 = tmp13820

						} else {
							ifres13808 = False

						}

						__e.TailApply(tmp13733, ifres13808)
						return

					} else {
						__e.Return(C3854)
						return
					}

				}, 1)

				tmp13836 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4324)

				var ifres13823 Obj

				if True == tmp13836 {
					tmp13824 := MakeNative(func(__e *ControlFlow) {
						Tm3855 := __e.Get(1)
						_ = Tm3855
						tmp13825 := MakeNative(func(__e *ControlFlow) {
							GoTo3856 := __e.Get(1)
							_ = GoTo3856
							tmp13829 := PrimEqual(Tm3855, symboolean)

							if True == tmp13829 {
								__e.TailApply(PrimNS2Value(symthaw), GoTo3856)
								return
							} else {
								tmp13828 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3855)

								if True == tmp13828 {
									__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3855, symboolean, V4323, GoTo3856)
									return
								} else {
									__e.Return(False)
									return
								}

							}

						}, 1)

						tmp13830 := MakeNative(func(__e *ControlFlow) {
							tmp13831 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp13831

							tmp13832 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4321, V4323)

							tmp13833 := Call(__e, PrimNS2Value(symboolean_2), tmp13832)

							__e.TailApply(PrimNS2Value(symwhen), tmp13833, V4323, V4324, V4325, V4326)
							return

						}, 0)

						__e.TailApply(tmp13825, tmp13830)
						return

					}, 1)

					tmp13834 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4322, V4323)

					tmp13835 := Call(__e, tmp13824, tmp13834)

					ifres13823 = tmp13835

				} else {
					ifres13823 = False

				}

				__e.TailApply(tmp13731, ifres13823)
				return

			} else {
				__e.Return(C3851)
				return
			}

		}, 1)

		tmp13851 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4324)

		var ifres13838 Obj

		if True == tmp13851 {
			tmp13839 := MakeNative(func(__e *ControlFlow) {
				Tm3852 := __e.Get(1)
				_ = Tm3852
				tmp13840 := MakeNative(func(__e *ControlFlow) {
					GoTo3853 := __e.Get(1)
					_ = GoTo3853
					tmp13844 := PrimEqual(Tm3852, symnumber)

					if True == tmp13844 {
						__e.TailApply(PrimNS2Value(symthaw), GoTo3853)
						return
					} else {
						tmp13843 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3852)

						if True == tmp13843 {
							__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm3852, symnumber, V4323, GoTo3853)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				tmp13845 := MakeNative(func(__e *ControlFlow) {
					tmp13846 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp13846

					tmp13847 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4321, V4323)

					tmp13848 := PrimIsNumber(tmp13847)

					__e.TailApply(PrimNS2Value(symwhen), tmp13848, V4323, V4324, V4325, V4326)
					return

				}, 0)

				__e.TailApply(tmp13840, tmp13845)
				return

			}, 1)

			tmp13849 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4322, V4323)

			tmp13850 := Call(__e, tmp13839, tmp13849)

			ifres13838 = tmp13850

		} else {
			ifres13838 = False

		}

		__e.TailApply(tmp13729, ifres13838)
		return

	}, 6)

	tmp13852 := Call(__e, PrimNS2Value(symdef), symshen_4primitive, tmp13728)

	_ = tmp13852

	tmp13853 := MakeNative(func(__e *ControlFlow) {
		V4327 := __e.Get(1)
		_ = V4327
		V4328 := __e.Get(2)
		_ = V4328
		V4329 := __e.Get(3)
		_ = V4329
		V4330 := __e.Get(4)
		_ = V4330
		V4331 := __e.Get(5)
		_ = V4331
		V4332 := __e.Get(6)
		_ = V4332
		V4333 := __e.Get(7)
		_ = V4333
		tmp13854 := MakeNative(func(__e *ControlFlow) {
			C3879 := __e.Get(1)
			_ = C3879
			tmp13865 := PrimEqual(C3879, False)

			if True == tmp13865 {
				tmp13864 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4331)

				if True == tmp13864 {
					tmp13857 := MakeNative(func(__e *ControlFlow) {
						Tm3885 := __e.Get(1)
						_ = Tm3885
						tmp13862 := PrimIsPair(Tm3885)

						if True == tmp13862 {
							tmp13859 := MakeNative(func(__e *ControlFlow) {
								Hyp := __e.Get(1)
								_ = Hyp
								tmp13860 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp13860

								__e.TailApply(PrimNS2Value(symshen_4by_1hypothesis), V4327, V4328, Hyp, V4330, V4331, V4332, V4333)
								return

							}, 1)

							tmp13861 := PrimTail(Tm3885)

							__e.TailApply(tmp13859, tmp13861)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp13863 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4329, V4330)

					__e.TailApply(tmp13857, tmp13863)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C3879)
				return
			}

		}, 1)

		tmp13907 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4331)

		var ifres13866 Obj

		if True == tmp13907 {
			tmp13867 := MakeNative(func(__e *ControlFlow) {
				Tm3880 := __e.Get(1)
				_ = Tm3880
				tmp13904 := PrimIsPair(Tm3880)

				if True == tmp13904 {
					tmp13869 := MakeNative(func(__e *ControlFlow) {
						Tm3881 := __e.Get(1)
						_ = Tm3881
						tmp13901 := PrimIsPair(Tm3881)

						if True == tmp13901 {
							tmp13871 := MakeNative(func(__e *ControlFlow) {
								Y := __e.Get(1)
								_ = Y
								tmp13872 := MakeNative(func(__e *ControlFlow) {
									Tm3882 := __e.Get(1)
									_ = Tm3882
									tmp13897 := PrimIsPair(Tm3882)

									if True == tmp13897 {
										tmp13874 := MakeNative(func(__e *ControlFlow) {
											Colon := __e.Get(1)
											_ = Colon
											tmp13875 := MakeNative(func(__e *ControlFlow) {
												Tm3883 := __e.Get(1)
												_ = Tm3883
												tmp13893 := PrimIsPair(Tm3883)

												if True == tmp13893 {
													tmp13877 := MakeNative(func(__e *ControlFlow) {
														B := __e.Get(1)
														_ = B
														tmp13878 := MakeNative(func(__e *ControlFlow) {
															Tm3884 := __e.Get(1)
															_ = Tm3884
															tmp13889 := PrimEqual(Tm3884, Nil)

															if True == tmp13889 {
																tmp13880 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp13880

																tmp13881 := Call(__e, PrimNS2Value(symshen_4deref), Colon, V4330)

																tmp13882 := PrimIntern(MakeString(":"))

																tmp13883 := PrimEqual(tmp13881, tmp13882)

																tmp13884 := MakeNative(func(__e *ControlFlow) {
																	tmp13885 := Call(__e, PrimNS2Value(symshen_4deref), V4327, V4330)

																	tmp13886 := Call(__e, PrimNS2Value(symshen_4deref), Y, V4330)

																	tmp13887 := PrimEqual(tmp13885, tmp13886)

																	tmp13888 := MakeNative(func(__e *ControlFlow) {
																		__e.TailApply(PrimNS2Value(symis_b), V4328, B, V4330, V4331, V4332, V4333)
																		return
																	}, 0)

																	__e.TailApply(PrimNS2Value(symwhen), tmp13887, V4330, V4331, V4332, tmp13888)
																	return

																}, 0)

																__e.TailApply(PrimNS2Value(symwhen), tmp13883, V4330, V4331, V4332, tmp13884)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp13890 := PrimTail(Tm3883)

														tmp13891 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13890, V4330)

														__e.TailApply(tmp13878, tmp13891)
														return

													}, 1)

													tmp13892 := PrimHead(Tm3883)

													__e.TailApply(tmp13877, tmp13892)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp13894 := PrimTail(Tm3882)

											tmp13895 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13894, V4330)

											__e.TailApply(tmp13875, tmp13895)
											return

										}, 1)

										tmp13896 := PrimHead(Tm3882)

										__e.TailApply(tmp13874, tmp13896)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp13898 := PrimTail(Tm3881)

								tmp13899 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13898, V4330)

								__e.TailApply(tmp13872, tmp13899)
								return

							}, 1)

							tmp13900 := PrimHead(Tm3881)

							__e.TailApply(tmp13871, tmp13900)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp13902 := PrimHead(Tm3880)

					tmp13903 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13902, V4330)

					__e.TailApply(tmp13869, tmp13903)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp13905 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4329, V4330)

			tmp13906 := Call(__e, tmp13867, tmp13905)

			ifres13866 = tmp13906

		} else {
			ifres13866 = False

		}

		__e.TailApply(tmp13854, ifres13866)
		return

	}, 7)

	tmp13908 := Call(__e, PrimNS2Value(symdef), symshen_4by_1hypothesis, tmp13853)

	_ = tmp13908

	tmp13909 := MakeNative(func(__e *ControlFlow) {
		V4334 := __e.Get(1)
		_ = V4334
		V4335 := __e.Get(2)
		_ = V4335
		V4336 := __e.Get(3)
		_ = V4336
		V4337 := __e.Get(4)
		_ = V4337
		V4338 := __e.Get(5)
		_ = V4338
		V4339 := __e.Get(6)
		_ = V4339
		tmp13914 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4337)

		if True == tmp13914 {
			tmp13911 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp13911

			tmp13912 := PrimNS3Value(symshen_4_dsigf_d)

			tmp13913 := Call(__e, PrimNS2Value(symassoc), V4334, tmp13912)

			__e.TailApply(PrimNS2Value(symshen_4sigf), tmp13913, V4335, V4336, V4337, V4338, V4339)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 6)

	tmp13915 := Call(__e, PrimNS2Value(symdef), symshen_4lookupsig, tmp13909)

	_ = tmp13915

	tmp13916 := MakeNative(func(__e *ControlFlow) {
		V4354 := __e.Get(1)
		_ = V4354
		V4355 := __e.Get(2)
		_ = V4355
		V4356 := __e.Get(3)
		_ = V4356
		V4357 := __e.Get(4)
		_ = V4357
		V4358 := __e.Get(5)
		_ = V4358
		V4359 := __e.Get(6)
		_ = V4359
		tmp13923 := PrimIsPair(V4354)

		if True == tmp13923 {
			tmp13918 := PrimTail(V4354)

			tmp13919 := Call(__e, tmp13918, V4355)

			tmp13920 := Call(__e, tmp13919, V4356)

			tmp13921 := Call(__e, tmp13920, V4357)

			tmp13922 := Call(__e, tmp13921, V4358)

			__e.TailApply(tmp13922, V4359)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 6)

	tmp13924 := Call(__e, PrimNS2Value(symdef), symshen_4sigf, tmp13916)

	_ = tmp13924

	tmp13925 := MakeNative(func(__e *ControlFlow) {
		V4360 := __e.Get(1)
		_ = V4360
		tmp13926 := MakeNative(func(__e *ControlFlow) {
			V := __e.Get(1)
			_ = V
			tmp13927 := MakeNative(func(__e *ControlFlow) {
				V0 := __e.Get(1)
				_ = V0
				tmp13928 := MakeNative(func(__e *ControlFlow) {
					V1 := __e.Get(1)
					_ = V1
					tmp13929 := MakeNative(func(__e *ControlFlow) {
						V2 := __e.Get(1)
						_ = V2
						__e.Return(V2)
						return
					}, 1)

					tmp13930 := PrimNS3Value(symshen_4_dgensym_d)

					tmp13931 := PrimNumberAdd(MakeNumber(1), tmp13930)

					tmp13932 := PrimNS3Set(symshen_4_dgensym_d, tmp13931)

					tmp13933 := PrimVectorSet(V1, MakeNumber(2), tmp13932)

					__e.TailApply(tmp13929, tmp13933)
					return

				}, 1)

				tmp13934 := PrimVectorSet(V0, MakeNumber(1), V4360)

				__e.TailApply(tmp13928, tmp13934)
				return

			}, 1)

			tmp13935 := PrimVectorSet(V, MakeNumber(0), symshen_4print_1freshterm)

			__e.TailApply(tmp13927, tmp13935)
			return

		}, 1)

		tmp13936 := PrimAbsvector(MakeNumber(3))

		__e.TailApply(tmp13926, tmp13936)
		return

	}, 1)

	tmp13937 := Call(__e, PrimNS2Value(symdef), symshen_4freshterm, tmp13925)

	_ = tmp13937

	tmp13938 := MakeNative(func(__e *ControlFlow) {
		V4361 := __e.Get(1)
		_ = V4361
		tmp13939 := PrimVectorGet(V4361, MakeNumber(1))

		tmp13940 := PrimStr(tmp13939)

		__e.Return(PrimStringConcat(MakeString("&&"), tmp13940))
		return

	}, 1)

	tmp13941 := Call(__e, PrimNS2Value(symdef), symshen_4print_1freshterm, tmp13938)

	_ = tmp13941

	tmp13942 := MakeNative(func(__e *ControlFlow) {
		V4362 := __e.Get(1)
		_ = V4362
		V4363 := __e.Get(2)
		_ = V4363
		V4364 := __e.Get(3)
		_ = V4364
		V4365 := __e.Get(4)
		_ = V4365
		V4366 := __e.Get(5)
		_ = V4366
		V4367 := __e.Get(6)
		_ = V4367
		V4368 := __e.Get(7)
		_ = V4368
		tmp13943 := MakeNative(func(__e *ControlFlow) {
			C3899 := __e.Get(1)
			_ = C3899
			tmp13954 := PrimEqual(C3899, False)

			if True == tmp13954 {
				tmp13953 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4366)

				if True == tmp13953 {
					tmp13946 := MakeNative(func(__e *ControlFlow) {
						Tm3902 := __e.Get(1)
						_ = Tm3902
						tmp13951 := PrimIsPair(Tm3902)

						if True == tmp13951 {
							tmp13948 := MakeNative(func(__e *ControlFlow) {
								Ds := __e.Get(1)
								_ = Ds
								tmp13949 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp13949

								__e.TailApply(PrimNS2Value(symshen_4search_1user_1datatypes), V4362, V4363, Ds, V4365, V4366, V4367, V4368)
								return

							}, 1)

							tmp13950 := PrimTail(Tm3902)

							__e.TailApply(tmp13948, tmp13950)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp13952 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4364, V4365)

					__e.TailApply(tmp13946, tmp13952)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C3899)
				return
			}

		}, 1)

		tmp13974 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4366)

		var ifres13955 Obj

		if True == tmp13974 {
			tmp13956 := MakeNative(func(__e *ControlFlow) {
				Tm3900 := __e.Get(1)
				_ = Tm3900
				tmp13971 := PrimIsPair(Tm3900)

				if True == tmp13971 {
					tmp13958 := MakeNative(func(__e *ControlFlow) {
						Tm3901 := __e.Get(1)
						_ = Tm3901
						tmp13968 := PrimIsPair(Tm3901)

						if True == tmp13968 {
							tmp13960 := MakeNative(func(__e *ControlFlow) {
								Fn := __e.Get(1)
								_ = Fn
								tmp13961 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp13961

								tmp13962 := Call(__e, PrimNS2Value(symshen_4deref), Fn, V4365)

								tmp13963 := Call(__e, PrimNS2Value(symshen_4deref), V4362, V4365)

								tmp13964 := Call(__e, tmp13962, tmp13963)

								tmp13965 := Call(__e, PrimNS2Value(symshen_4deref), V4363, V4365)

								tmp13966 := Call(__e, tmp13964, tmp13965)

								__e.TailApply(PrimNS2Value(symcall), tmp13966, V4365, V4366, V4367, V4368)
								return

							}, 1)

							tmp13967 := PrimTail(Tm3901)

							__e.TailApply(tmp13960, tmp13967)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp13969 := PrimHead(Tm3900)

					tmp13970 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp13969, V4365)

					__e.TailApply(tmp13958, tmp13970)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp13972 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4364, V4365)

			tmp13973 := Call(__e, tmp13956, tmp13972)

			ifres13955 = tmp13973

		} else {
			ifres13955 = False

		}

		__e.TailApply(tmp13943, ifres13955)
		return

	}, 7)

	tmp13975 := Call(__e, PrimNS2Value(symdef), symshen_4search_1user_1datatypes, tmp13942)

	_ = tmp13975

	tmp13976 := MakeNative(func(__e *ControlFlow) {
		V4369 := __e.Get(1)
		_ = V4369
		V4370 := __e.Get(2)
		_ = V4370
		V4371 := __e.Get(3)
		_ = V4371
		V4372 := __e.Get(4)
		_ = V4372
		V4373 := __e.Get(5)
		_ = V4373
		V4374 := __e.Get(6)
		_ = V4374
		V4375 := __e.Get(7)
		_ = V4375
		tmp13977 := MakeNative(func(__e *ControlFlow) {
			K3905 := __e.Get(1)
			_ = K3905
			tmp13978 := MakeNative(func(__e *ControlFlow) {
				C3910 := __e.Get(1)
				_ = C3910
				tmp14408 := PrimEqual(C3910, False)

				if True == tmp14408 {
					tmp13980 := MakeNative(func(__e *ControlFlow) {
						C3913 := __e.Get(1)
						_ = C3913
						tmp14309 := PrimEqual(C3913, False)

						if True == tmp14309 {
							tmp13982 := MakeNative(func(__e *ControlFlow) {
								C3928 := __e.Get(1)
								_ = C3928
								tmp14205 := PrimEqual(C3928, False)

								if True == tmp14205 {
									tmp13984 := MakeNative(func(__e *ControlFlow) {
										C3944 := __e.Get(1)
										_ = C3944
										tmp14125 := PrimEqual(C3944, False)

										if True == tmp14125 {
											tmp13986 := MakeNative(func(__e *ControlFlow) {
												C3956 := __e.Get(1)
												_ = C3956
												tmp14026 := PrimEqual(C3956, False)

												if True == tmp14026 {
													tmp13988 := MakeNative(func(__e *ControlFlow) {
														C3971 := __e.Get(1)
														_ = C3971
														tmp13990 := PrimEqual(C3971, False)

														if True == tmp13990 {
															__e.TailApply(PrimNS2Value(symshen_4unlock), V4373, K3905)
															return
														} else {
															__e.Return(C3971)
															return
														}

													}, 1)

													tmp14025 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4373)

													var ifres13991 Obj

													if True == tmp14025 {
														tmp13992 := MakeNative(func(__e *ControlFlow) {
															Tm3972 := __e.Get(1)
															_ = Tm3972
															tmp14022 := PrimIsPair(Tm3972)

															if True == tmp14022 {
																tmp13994 := MakeNative(func(__e *ControlFlow) {
																	P := __e.Get(1)
																	_ = P
																	tmp13995 := MakeNative(func(__e *ControlFlow) {
																		Hyp := __e.Get(1)
																		_ = Hyp
																		tmp13996 := MakeNative(func(__e *ControlFlow) {
																			Tm3973 := __e.Get(1)
																			_ = Tm3973
																			tmp13997 := MakeNative(func(__e *ControlFlow) {
																				GoTo3974 := __e.Get(1)
																				_ = GoTo3974
																				tmp14015 := PrimIsPair(Tm3973)

																				if True == tmp14015 {
																					tmp14010 := MakeNative(func(__e *ControlFlow) {
																						Q := __e.Get(1)
																						_ = Q
																						tmp14011 := MakeNative(func(__e *ControlFlow) {
																							Normalised := __e.Get(1)
																							_ = Normalised
																							tmp14012 := Call(__e, GoTo3974, Q)

																							__e.TailApply(tmp14012, Normalised)
																							return

																						}, 1)

																						tmp14013 := PrimTail(Tm3973)

																						__e.TailApply(tmp14011, tmp14013)
																						return

																					}, 1)

																					tmp14014 := PrimHead(Tm3973)

																					__e.TailApply(tmp14010, tmp14014)
																					return

																				} else {
																					tmp14009 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm3973)

																					if True == tmp14009 {
																						tmp14000 := MakeNative(func(__e *ControlFlow) {
																							Q := __e.Get(1)
																							_ = Q
																							tmp14001 := MakeNative(func(__e *ControlFlow) {
																								Normalised := __e.Get(1)
																								_ = Normalised
																								tmp14002 := PrimCons(Q, Normalised)

																								tmp14003 := MakeNative(func(__e *ControlFlow) {
																									tmp14004 := Call(__e, GoTo3974, Q)

																									__e.TailApply(tmp14004, Normalised)
																									return

																								}, 0)

																								tmp14005 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm3973, tmp14002, V4372, tmp14003)

																								__e.TailApply(PrimNS2Value(symshen_4gc), V4372, tmp14005)
																								return

																							}, 1)

																							tmp14006 := Call(__e, PrimNS2Value(symshen_4newpv), V4372)

																							tmp14007 := Call(__e, tmp14001, tmp14006)

																							__e.TailApply(PrimNS2Value(symshen_4gc), V4372, tmp14007)
																							return

																						}, 1)

																						tmp14008 := Call(__e, PrimNS2Value(symshen_4newpv), V4372)

																						__e.TailApply(tmp14000, tmp14008)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}

																			}, 1)

																			tmp14016 := MakeNative(func(__e *ControlFlow) {
																				Q := __e.Get(1)
																				_ = Q
																				__e.Return(MakeNative(func(__e *ControlFlow) {
																					Normalised := __e.Get(1)
																					_ = Normalised
																					tmp14017 := Call(__e, PrimNS2Value(symshen_4incinfs))

																					_ = tmp14017

																					tmp14018 := MakeNative(func(__e *ControlFlow) {
																						__e.TailApply(PrimNS2Value(symshen_4l_1rules), Hyp, Normalised, V4371, V4372, V4373, K3905, V4375)
																						return
																					}, 0)

																					__e.TailApply(PrimNS2Value(symbind), Q, P, V4372, V4373, K3905, tmp14018)
																					return

																				}, 1))
																				return
																			}, 1)

																			__e.TailApply(tmp13997, tmp14016)
																			return

																		}, 1)

																		tmp14019 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4370, V4372)

																		__e.TailApply(tmp13996, tmp14019)
																		return

																	}, 1)

																	tmp14020 := PrimTail(Tm3972)

																	__e.TailApply(tmp13995, tmp14020)
																	return

																}, 1)

																tmp14021 := PrimHead(Tm3972)

																__e.TailApply(tmp13994, tmp14021)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14023 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4369, V4372)

														tmp14024 := Call(__e, tmp13992, tmp14023)

														ifres13991 = tmp14024

													} else {
														ifres13991 = False

													}

													__e.TailApply(tmp13988, ifres13991)
													return

												} else {
													__e.Return(C3956)
													return
												}

											}, 1)

											tmp14124 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4373)

											var ifres14027 Obj

											if True == tmp14124 {
												tmp14028 := MakeNative(func(__e *ControlFlow) {
													Tm3957 := __e.Get(1)
													_ = Tm3957
													tmp14121 := PrimIsPair(Tm3957)

													if True == tmp14121 {
														tmp14030 := MakeNative(func(__e *ControlFlow) {
															Tm3958 := __e.Get(1)
															_ = Tm3958
															tmp14118 := PrimIsPair(Tm3958)

															if True == tmp14118 {
																tmp14032 := MakeNative(func(__e *ControlFlow) {
																	Tm3959 := __e.Get(1)
																	_ = Tm3959
																	tmp14115 := PrimIsPair(Tm3959)

																	if True == tmp14115 {
																		tmp14034 := MakeNative(func(__e *ControlFlow) {
																			Tm3960 := __e.Get(1)
																			_ = Tm3960
																			tmp14112 := PrimEqual(Tm3960, sym_8v)

																			if True == tmp14112 {
																				tmp14036 := MakeNative(func(__e *ControlFlow) {
																					Tm3961 := __e.Get(1)
																					_ = Tm3961
																					tmp14109 := PrimIsPair(Tm3961)

																					if True == tmp14109 {
																						tmp14038 := MakeNative(func(__e *ControlFlow) {
																							X := __e.Get(1)
																							_ = X
																							tmp14039 := MakeNative(func(__e *ControlFlow) {
																								Tm3962 := __e.Get(1)
																								_ = Tm3962
																								tmp14105 := PrimIsPair(Tm3962)

																								if True == tmp14105 {
																									tmp14041 := MakeNative(func(__e *ControlFlow) {
																										Y := __e.Get(1)
																										_ = Y
																										tmp14042 := MakeNative(func(__e *ControlFlow) {
																											Tm3963 := __e.Get(1)
																											_ = Tm3963
																											tmp14101 := PrimEqual(Tm3963, Nil)

																											if True == tmp14101 {
																												tmp14044 := MakeNative(func(__e *ControlFlow) {
																													Tm3964 := __e.Get(1)
																													_ = Tm3964
																													tmp14098 := PrimIsPair(Tm3964)

																													if True == tmp14098 {
																														tmp14046 := MakeNative(func(__e *ControlFlow) {
																															Colon := __e.Get(1)
																															_ = Colon
																															tmp14047 := MakeNative(func(__e *ControlFlow) {
																																Tm3965 := __e.Get(1)
																																_ = Tm3965
																																tmp14094 := PrimIsPair(Tm3965)

																																if True == tmp14094 {
																																	tmp14049 := MakeNative(func(__e *ControlFlow) {
																																		Tm3966 := __e.Get(1)
																																		_ = Tm3966
																																		tmp14091 := PrimIsPair(Tm3966)

																																		if True == tmp14091 {
																																			tmp14051 := MakeNative(func(__e *ControlFlow) {
																																				Tm3967 := __e.Get(1)
																																				_ = Tm3967
																																				tmp14088 := PrimEqual(Tm3967, symvector)

																																				if True == tmp14088 {
																																					tmp14053 := MakeNative(func(__e *ControlFlow) {
																																						Tm3968 := __e.Get(1)
																																						_ = Tm3968
																																						tmp14085 := PrimIsPair(Tm3968)

																																						if True == tmp14085 {
																																							tmp14055 := MakeNative(func(__e *ControlFlow) {
																																								A := __e.Get(1)
																																								_ = A
																																								tmp14056 := MakeNative(func(__e *ControlFlow) {
																																									Tm3969 := __e.Get(1)
																																									_ = Tm3969
																																									tmp14081 := PrimEqual(Tm3969, Nil)

																																									if True == tmp14081 {
																																										tmp14058 := MakeNative(func(__e *ControlFlow) {
																																											Tm3970 := __e.Get(1)
																																											_ = Tm3970
																																											tmp14078 := PrimEqual(Tm3970, Nil)

																																											if True == tmp14078 {
																																												tmp14060 := MakeNative(func(__e *ControlFlow) {
																																													Hyp := __e.Get(1)
																																													_ = Hyp
																																													tmp14061 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																													_ = tmp14061

																																													tmp14062 := Call(__e, PrimNS2Value(symshen_4deref), Colon, V4372)

																																													tmp14063 := PrimIntern(MakeString(":"))

																																													tmp14064 := PrimEqual(tmp14062, tmp14063)

																																													tmp14065 := MakeNative(func(__e *ControlFlow) {
																																														tmp14066 := MakeNative(func(__e *ControlFlow) {
																																															tmp14067 := PrimCons(A, Nil)

																																															tmp14068 := PrimCons(Colon, tmp14067)

																																															tmp14069 := PrimCons(X, tmp14068)

																																															tmp14070 := PrimCons(A, Nil)

																																															tmp14071 := PrimCons(symvector, tmp14070)

																																															tmp14072 := PrimCons(tmp14071, Nil)

																																															tmp14073 := PrimCons(Colon, tmp14072)

																																															tmp14074 := PrimCons(Y, tmp14073)

																																															tmp14075 := PrimCons(tmp14074, Hyp)

																																															tmp14076 := PrimCons(tmp14069, tmp14075)

																																															__e.TailApply(PrimNS2Value(symshen_4l_1rules), tmp14076, V4370, True, V4372, V4373, K3905, V4375)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimNS2Value(symshen_4cut), V4372, V4373, K3905, tmp14066)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symwhen), tmp14064, V4372, V4373, K3905, tmp14065)
																																													return

																																												}, 1)

																																												tmp14077 := PrimTail(Tm3957)

																																												__e.TailApply(tmp14060, tmp14077)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}, 1)

																																										tmp14079 := PrimTail(Tm3965)

																																										tmp14080 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14079, V4372)

																																										__e.TailApply(tmp14058, tmp14080)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp14082 := PrimTail(Tm3968)

																																								tmp14083 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14082, V4372)

																																								__e.TailApply(tmp14056, tmp14083)
																																								return

																																							}, 1)

																																							tmp14084 := PrimHead(Tm3968)

																																							__e.TailApply(tmp14055, tmp14084)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp14086 := PrimTail(Tm3966)

																																					tmp14087 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14086, V4372)

																																					__e.TailApply(tmp14053, tmp14087)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp14089 := PrimHead(Tm3966)

																																			tmp14090 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14089, V4372)

																																			__e.TailApply(tmp14051, tmp14090)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp14092 := PrimHead(Tm3965)

																																	tmp14093 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14092, V4372)

																																	__e.TailApply(tmp14049, tmp14093)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp14095 := PrimTail(Tm3964)

																															tmp14096 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14095, V4372)

																															__e.TailApply(tmp14047, tmp14096)
																															return

																														}, 1)

																														tmp14097 := PrimHead(Tm3964)

																														__e.TailApply(tmp14046, tmp14097)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp14099 := PrimTail(Tm3958)

																												tmp14100 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14099, V4372)

																												__e.TailApply(tmp14044, tmp14100)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp14102 := PrimTail(Tm3962)

																										tmp14103 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14102, V4372)

																										__e.TailApply(tmp14042, tmp14103)
																										return

																									}, 1)

																									tmp14104 := PrimHead(Tm3962)

																									__e.TailApply(tmp14041, tmp14104)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp14106 := PrimTail(Tm3961)

																							tmp14107 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14106, V4372)

																							__e.TailApply(tmp14039, tmp14107)
																							return

																						}, 1)

																						tmp14108 := PrimHead(Tm3961)

																						__e.TailApply(tmp14038, tmp14108)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp14110 := PrimTail(Tm3959)

																				tmp14111 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14110, V4372)

																				__e.TailApply(tmp14036, tmp14111)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp14113 := PrimHead(Tm3959)

																		tmp14114 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14113, V4372)

																		__e.TailApply(tmp14034, tmp14114)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp14116 := PrimHead(Tm3958)

																tmp14117 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14116, V4372)

																__e.TailApply(tmp14032, tmp14117)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14119 := PrimHead(Tm3957)

														tmp14120 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14119, V4372)

														__e.TailApply(tmp14030, tmp14120)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14122 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4369, V4372)

												tmp14123 := Call(__e, tmp14028, tmp14122)

												ifres14027 = tmp14123

											} else {
												ifres14027 = False

											}

											__e.TailApply(tmp13986, ifres14027)
											return

										} else {
											__e.Return(C3944)
											return
										}

									}, 1)

									tmp14204 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4373)

									var ifres14126 Obj

									if True == tmp14204 {
										tmp14127 := MakeNative(func(__e *ControlFlow) {
											Tm3945 := __e.Get(1)
											_ = Tm3945
											tmp14201 := PrimIsPair(Tm3945)

											if True == tmp14201 {
												tmp14129 := MakeNative(func(__e *ControlFlow) {
													Tm3946 := __e.Get(1)
													_ = Tm3946
													tmp14198 := PrimIsPair(Tm3946)

													if True == tmp14198 {
														tmp14131 := MakeNative(func(__e *ControlFlow) {
															Tm3947 := __e.Get(1)
															_ = Tm3947
															tmp14195 := PrimIsPair(Tm3947)

															if True == tmp14195 {
																tmp14133 := MakeNative(func(__e *ControlFlow) {
																	Tm3948 := __e.Get(1)
																	_ = Tm3948
																	tmp14192 := PrimEqual(Tm3948, sym_8s)

																	if True == tmp14192 {
																		tmp14135 := MakeNative(func(__e *ControlFlow) {
																			Tm3949 := __e.Get(1)
																			_ = Tm3949
																			tmp14189 := PrimIsPair(Tm3949)

																			if True == tmp14189 {
																				tmp14137 := MakeNative(func(__e *ControlFlow) {
																					X := __e.Get(1)
																					_ = X
																					tmp14138 := MakeNative(func(__e *ControlFlow) {
																						Tm3950 := __e.Get(1)
																						_ = Tm3950
																						tmp14185 := PrimIsPair(Tm3950)

																						if True == tmp14185 {
																							tmp14140 := MakeNative(func(__e *ControlFlow) {
																								Y := __e.Get(1)
																								_ = Y
																								tmp14141 := MakeNative(func(__e *ControlFlow) {
																									Tm3951 := __e.Get(1)
																									_ = Tm3951
																									tmp14181 := PrimEqual(Tm3951, Nil)

																									if True == tmp14181 {
																										tmp14143 := MakeNative(func(__e *ControlFlow) {
																											Tm3952 := __e.Get(1)
																											_ = Tm3952
																											tmp14178 := PrimIsPair(Tm3952)

																											if True == tmp14178 {
																												tmp14145 := MakeNative(func(__e *ControlFlow) {
																													Colon := __e.Get(1)
																													_ = Colon
																													tmp14146 := MakeNative(func(__e *ControlFlow) {
																														Tm3953 := __e.Get(1)
																														_ = Tm3953
																														tmp14174 := PrimIsPair(Tm3953)

																														if True == tmp14174 {
																															tmp14148 := MakeNative(func(__e *ControlFlow) {
																																Tm3954 := __e.Get(1)
																																_ = Tm3954
																																tmp14171 := PrimEqual(Tm3954, symstring)

																																if True == tmp14171 {
																																	tmp14150 := MakeNative(func(__e *ControlFlow) {
																																		Tm3955 := __e.Get(1)
																																		_ = Tm3955
																																		tmp14168 := PrimEqual(Tm3955, Nil)

																																		if True == tmp14168 {
																																			tmp14152 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp14153 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp14153

																																				tmp14154 := Call(__e, PrimNS2Value(symshen_4deref), Colon, V4372)

																																				tmp14155 := PrimIntern(MakeString(":"))

																																				tmp14156 := PrimEqual(tmp14154, tmp14155)

																																				tmp14157 := MakeNative(func(__e *ControlFlow) {
																																					tmp14158 := MakeNative(func(__e *ControlFlow) {
																																						tmp14159 := PrimCons(symstring, Nil)

																																						tmp14160 := PrimCons(Colon, tmp14159)

																																						tmp14161 := PrimCons(X, tmp14160)

																																						tmp14162 := PrimCons(symstring, Nil)

																																						tmp14163 := PrimCons(Colon, tmp14162)

																																						tmp14164 := PrimCons(Y, tmp14163)

																																						tmp14165 := PrimCons(tmp14164, Hyp)

																																						tmp14166 := PrimCons(tmp14161, tmp14165)

																																						__e.TailApply(PrimNS2Value(symshen_4l_1rules), tmp14166, V4370, True, V4372, V4373, K3905, V4375)
																																						return

																																					}, 0)

																																					__e.TailApply(PrimNS2Value(symshen_4cut), V4372, V4373, K3905, tmp14158)
																																					return

																																				}, 0)

																																				__e.TailApply(PrimNS2Value(symwhen), tmp14156, V4372, V4373, K3905, tmp14157)
																																				return

																																			}, 1)

																																			tmp14167 := PrimTail(Tm3945)

																																			__e.TailApply(tmp14152, tmp14167)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp14169 := PrimTail(Tm3953)

																																	tmp14170 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14169, V4372)

																																	__e.TailApply(tmp14150, tmp14170)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp14172 := PrimHead(Tm3953)

																															tmp14173 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14172, V4372)

																															__e.TailApply(tmp14148, tmp14173)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp14175 := PrimTail(Tm3952)

																													tmp14176 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14175, V4372)

																													__e.TailApply(tmp14146, tmp14176)
																													return

																												}, 1)

																												tmp14177 := PrimHead(Tm3952)

																												__e.TailApply(tmp14145, tmp14177)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp14179 := PrimTail(Tm3946)

																										tmp14180 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14179, V4372)

																										__e.TailApply(tmp14143, tmp14180)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp14182 := PrimTail(Tm3950)

																								tmp14183 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14182, V4372)

																								__e.TailApply(tmp14141, tmp14183)
																								return

																							}, 1)

																							tmp14184 := PrimHead(Tm3950)

																							__e.TailApply(tmp14140, tmp14184)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp14186 := PrimTail(Tm3949)

																					tmp14187 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14186, V4372)

																					__e.TailApply(tmp14138, tmp14187)
																					return

																				}, 1)

																				tmp14188 := PrimHead(Tm3949)

																				__e.TailApply(tmp14137, tmp14188)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp14190 := PrimTail(Tm3947)

																		tmp14191 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14190, V4372)

																		__e.TailApply(tmp14135, tmp14191)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp14193 := PrimHead(Tm3947)

																tmp14194 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14193, V4372)

																__e.TailApply(tmp14133, tmp14194)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14196 := PrimHead(Tm3946)

														tmp14197 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14196, V4372)

														__e.TailApply(tmp14131, tmp14197)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14199 := PrimHead(Tm3945)

												tmp14200 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14199, V4372)

												__e.TailApply(tmp14129, tmp14200)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14202 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4369, V4372)

										tmp14203 := Call(__e, tmp14127, tmp14202)

										ifres14126 = tmp14203

									} else {
										ifres14126 = False

									}

									__e.TailApply(tmp13984, ifres14126)
									return

								} else {
									__e.Return(C3928)
									return
								}

							}, 1)

							tmp14308 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4373)

							var ifres14206 Obj

							if True == tmp14308 {
								tmp14207 := MakeNative(func(__e *ControlFlow) {
									Tm3929 := __e.Get(1)
									_ = Tm3929
									tmp14305 := PrimIsPair(Tm3929)

									if True == tmp14305 {
										tmp14209 := MakeNative(func(__e *ControlFlow) {
											Tm3930 := __e.Get(1)
											_ = Tm3930
											tmp14302 := PrimIsPair(Tm3930)

											if True == tmp14302 {
												tmp14211 := MakeNative(func(__e *ControlFlow) {
													Tm3931 := __e.Get(1)
													_ = Tm3931
													tmp14299 := PrimIsPair(Tm3931)

													if True == tmp14299 {
														tmp14213 := MakeNative(func(__e *ControlFlow) {
															Tm3932 := __e.Get(1)
															_ = Tm3932
															tmp14296 := PrimEqual(Tm3932, sym_8p)

															if True == tmp14296 {
																tmp14215 := MakeNative(func(__e *ControlFlow) {
																	Tm3933 := __e.Get(1)
																	_ = Tm3933
																	tmp14293 := PrimIsPair(Tm3933)

																	if True == tmp14293 {
																		tmp14217 := MakeNative(func(__e *ControlFlow) {
																			X := __e.Get(1)
																			_ = X
																			tmp14218 := MakeNative(func(__e *ControlFlow) {
																				Tm3934 := __e.Get(1)
																				_ = Tm3934
																				tmp14289 := PrimIsPair(Tm3934)

																				if True == tmp14289 {
																					tmp14220 := MakeNative(func(__e *ControlFlow) {
																						Y := __e.Get(1)
																						_ = Y
																						tmp14221 := MakeNative(func(__e *ControlFlow) {
																							Tm3935 := __e.Get(1)
																							_ = Tm3935
																							tmp14285 := PrimEqual(Tm3935, Nil)

																							if True == tmp14285 {
																								tmp14223 := MakeNative(func(__e *ControlFlow) {
																									Tm3936 := __e.Get(1)
																									_ = Tm3936
																									tmp14282 := PrimIsPair(Tm3936)

																									if True == tmp14282 {
																										tmp14225 := MakeNative(func(__e *ControlFlow) {
																											Colon := __e.Get(1)
																											_ = Colon
																											tmp14226 := MakeNative(func(__e *ControlFlow) {
																												Tm3937 := __e.Get(1)
																												_ = Tm3937
																												tmp14278 := PrimIsPair(Tm3937)

																												if True == tmp14278 {
																													tmp14228 := MakeNative(func(__e *ControlFlow) {
																														Tm3938 := __e.Get(1)
																														_ = Tm3938
																														tmp14275 := PrimIsPair(Tm3938)

																														if True == tmp14275 {
																															tmp14230 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp14231 := MakeNative(func(__e *ControlFlow) {
																																	Tm3939 := __e.Get(1)
																																	_ = Tm3939
																																	tmp14271 := PrimIsPair(Tm3939)

																																	if True == tmp14271 {
																																		tmp14233 := MakeNative(func(__e *ControlFlow) {
																																			Tm3940 := __e.Get(1)
																																			_ = Tm3940
																																			tmp14268 := PrimEqual(Tm3940, sym_d)

																																			if True == tmp14268 {
																																				tmp14235 := MakeNative(func(__e *ControlFlow) {
																																					Tm3941 := __e.Get(1)
																																					_ = Tm3941
																																					tmp14265 := PrimIsPair(Tm3941)

																																					if True == tmp14265 {
																																						tmp14237 := MakeNative(func(__e *ControlFlow) {
																																							B := __e.Get(1)
																																							_ = B
																																							tmp14238 := MakeNative(func(__e *ControlFlow) {
																																								Tm3942 := __e.Get(1)
																																								_ = Tm3942
																																								tmp14261 := PrimEqual(Tm3942, Nil)

																																								if True == tmp14261 {
																																									tmp14240 := MakeNative(func(__e *ControlFlow) {
																																										Tm3943 := __e.Get(1)
																																										_ = Tm3943
																																										tmp14258 := PrimEqual(Tm3943, Nil)

																																										if True == tmp14258 {
																																											tmp14242 := MakeNative(func(__e *ControlFlow) {
																																												Hyp := __e.Get(1)
																																												_ = Hyp
																																												tmp14243 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																												_ = tmp14243

																																												tmp14244 := Call(__e, PrimNS2Value(symshen_4deref), Colon, V4372)

																																												tmp14245 := PrimIntern(MakeString(":"))

																																												tmp14246 := PrimEqual(tmp14244, tmp14245)

																																												tmp14247 := MakeNative(func(__e *ControlFlow) {
																																													tmp14248 := MakeNative(func(__e *ControlFlow) {
																																														tmp14249 := PrimCons(A, Nil)

																																														tmp14250 := PrimCons(Colon, tmp14249)

																																														tmp14251 := PrimCons(X, tmp14250)

																																														tmp14252 := PrimCons(B, Nil)

																																														tmp14253 := PrimCons(Colon, tmp14252)

																																														tmp14254 := PrimCons(Y, tmp14253)

																																														tmp14255 := PrimCons(tmp14254, Hyp)

																																														tmp14256 := PrimCons(tmp14251, tmp14255)

																																														__e.TailApply(PrimNS2Value(symshen_4l_1rules), tmp14256, V4370, True, V4372, V4373, K3905, V4375)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symshen_4cut), V4372, V4373, K3905, tmp14248)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symwhen), tmp14246, V4372, V4373, K3905, tmp14247)
																																												return

																																											}, 1)

																																											tmp14257 := PrimTail(Tm3929)

																																											__e.TailApply(tmp14242, tmp14257)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp14259 := PrimTail(Tm3937)

																																									tmp14260 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14259, V4372)

																																									__e.TailApply(tmp14240, tmp14260)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp14262 := PrimTail(Tm3941)

																																							tmp14263 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14262, V4372)

																																							__e.TailApply(tmp14238, tmp14263)
																																							return

																																						}, 1)

																																						tmp14264 := PrimHead(Tm3941)

																																						__e.TailApply(tmp14237, tmp14264)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp14266 := PrimTail(Tm3939)

																																				tmp14267 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14266, V4372)

																																				__e.TailApply(tmp14235, tmp14267)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp14269 := PrimHead(Tm3939)

																																		tmp14270 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14269, V4372)

																																		__e.TailApply(tmp14233, tmp14270)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp14272 := PrimTail(Tm3938)

																																tmp14273 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14272, V4372)

																																__e.TailApply(tmp14231, tmp14273)
																																return

																															}, 1)

																															tmp14274 := PrimHead(Tm3938)

																															__e.TailApply(tmp14230, tmp14274)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp14276 := PrimHead(Tm3937)

																													tmp14277 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14276, V4372)

																													__e.TailApply(tmp14228, tmp14277)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp14279 := PrimTail(Tm3936)

																											tmp14280 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14279, V4372)

																											__e.TailApply(tmp14226, tmp14280)
																											return

																										}, 1)

																										tmp14281 := PrimHead(Tm3936)

																										__e.TailApply(tmp14225, tmp14281)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp14283 := PrimTail(Tm3930)

																								tmp14284 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14283, V4372)

																								__e.TailApply(tmp14223, tmp14284)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp14286 := PrimTail(Tm3934)

																						tmp14287 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14286, V4372)

																						__e.TailApply(tmp14221, tmp14287)
																						return

																					}, 1)

																					tmp14288 := PrimHead(Tm3934)

																					__e.TailApply(tmp14220, tmp14288)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp14290 := PrimTail(Tm3933)

																			tmp14291 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14290, V4372)

																			__e.TailApply(tmp14218, tmp14291)
																			return

																		}, 1)

																		tmp14292 := PrimHead(Tm3933)

																		__e.TailApply(tmp14217, tmp14292)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp14294 := PrimTail(Tm3931)

																tmp14295 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14294, V4372)

																__e.TailApply(tmp14215, tmp14295)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14297 := PrimHead(Tm3931)

														tmp14298 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14297, V4372)

														__e.TailApply(tmp14213, tmp14298)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14300 := PrimHead(Tm3930)

												tmp14301 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14300, V4372)

												__e.TailApply(tmp14211, tmp14301)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14303 := PrimHead(Tm3929)

										tmp14304 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14303, V4372)

										__e.TailApply(tmp14209, tmp14304)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14306 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4369, V4372)

								tmp14307 := Call(__e, tmp14207, tmp14306)

								ifres14206 = tmp14307

							} else {
								ifres14206 = False

							}

							__e.TailApply(tmp13982, ifres14206)
							return

						} else {
							__e.Return(C3913)
							return
						}

					}, 1)

					tmp14407 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4373)

					var ifres14310 Obj

					if True == tmp14407 {
						tmp14311 := MakeNative(func(__e *ControlFlow) {
							Tm3914 := __e.Get(1)
							_ = Tm3914
							tmp14404 := PrimIsPair(Tm3914)

							if True == tmp14404 {
								tmp14313 := MakeNative(func(__e *ControlFlow) {
									Tm3915 := __e.Get(1)
									_ = Tm3915
									tmp14401 := PrimIsPair(Tm3915)

									if True == tmp14401 {
										tmp14315 := MakeNative(func(__e *ControlFlow) {
											Tm3916 := __e.Get(1)
											_ = Tm3916
											tmp14398 := PrimIsPair(Tm3916)

											if True == tmp14398 {
												tmp14317 := MakeNative(func(__e *ControlFlow) {
													Tm3917 := __e.Get(1)
													_ = Tm3917
													tmp14395 := PrimEqual(Tm3917, symcons)

													if True == tmp14395 {
														tmp14319 := MakeNative(func(__e *ControlFlow) {
															Tm3918 := __e.Get(1)
															_ = Tm3918
															tmp14392 := PrimIsPair(Tm3918)

															if True == tmp14392 {
																tmp14321 := MakeNative(func(__e *ControlFlow) {
																	X := __e.Get(1)
																	_ = X
																	tmp14322 := MakeNative(func(__e *ControlFlow) {
																		Tm3919 := __e.Get(1)
																		_ = Tm3919
																		tmp14388 := PrimIsPair(Tm3919)

																		if True == tmp14388 {
																			tmp14324 := MakeNative(func(__e *ControlFlow) {
																				Y := __e.Get(1)
																				_ = Y
																				tmp14325 := MakeNative(func(__e *ControlFlow) {
																					Tm3920 := __e.Get(1)
																					_ = Tm3920
																					tmp14384 := PrimEqual(Tm3920, Nil)

																					if True == tmp14384 {
																						tmp14327 := MakeNative(func(__e *ControlFlow) {
																							Tm3921 := __e.Get(1)
																							_ = Tm3921
																							tmp14381 := PrimIsPair(Tm3921)

																							if True == tmp14381 {
																								tmp14329 := MakeNative(func(__e *ControlFlow) {
																									Colon := __e.Get(1)
																									_ = Colon
																									tmp14330 := MakeNative(func(__e *ControlFlow) {
																										Tm3922 := __e.Get(1)
																										_ = Tm3922
																										tmp14377 := PrimIsPair(Tm3922)

																										if True == tmp14377 {
																											tmp14332 := MakeNative(func(__e *ControlFlow) {
																												Tm3923 := __e.Get(1)
																												_ = Tm3923
																												tmp14374 := PrimIsPair(Tm3923)

																												if True == tmp14374 {
																													tmp14334 := MakeNative(func(__e *ControlFlow) {
																														Tm3924 := __e.Get(1)
																														_ = Tm3924
																														tmp14371 := PrimEqual(Tm3924, symlist)

																														if True == tmp14371 {
																															tmp14336 := MakeNative(func(__e *ControlFlow) {
																																Tm3925 := __e.Get(1)
																																_ = Tm3925
																																tmp14368 := PrimIsPair(Tm3925)

																																if True == tmp14368 {
																																	tmp14338 := MakeNative(func(__e *ControlFlow) {
																																		A := __e.Get(1)
																																		_ = A
																																		tmp14339 := MakeNative(func(__e *ControlFlow) {
																																			Tm3926 := __e.Get(1)
																																			_ = Tm3926
																																			tmp14364 := PrimEqual(Tm3926, Nil)

																																			if True == tmp14364 {
																																				tmp14341 := MakeNative(func(__e *ControlFlow) {
																																					Tm3927 := __e.Get(1)
																																					_ = Tm3927
																																					tmp14361 := PrimEqual(Tm3927, Nil)

																																					if True == tmp14361 {
																																						tmp14343 := MakeNative(func(__e *ControlFlow) {
																																							Hyp := __e.Get(1)
																																							_ = Hyp
																																							tmp14344 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																							_ = tmp14344

																																							tmp14345 := Call(__e, PrimNS2Value(symshen_4deref), Colon, V4372)

																																							tmp14346 := PrimIntern(MakeString(":"))

																																							tmp14347 := PrimEqual(tmp14345, tmp14346)

																																							tmp14348 := MakeNative(func(__e *ControlFlow) {
																																								tmp14349 := MakeNative(func(__e *ControlFlow) {
																																									tmp14350 := PrimCons(A, Nil)

																																									tmp14351 := PrimCons(Colon, tmp14350)

																																									tmp14352 := PrimCons(X, tmp14351)

																																									tmp14353 := PrimCons(A, Nil)

																																									tmp14354 := PrimCons(symlist, tmp14353)

																																									tmp14355 := PrimCons(tmp14354, Nil)

																																									tmp14356 := PrimCons(Colon, tmp14355)

																																									tmp14357 := PrimCons(Y, tmp14356)

																																									tmp14358 := PrimCons(tmp14357, Hyp)

																																									tmp14359 := PrimCons(tmp14352, tmp14358)

																																									__e.TailApply(PrimNS2Value(symshen_4l_1rules), tmp14359, V4370, True, V4372, V4373, K3905, V4375)
																																									return

																																								}, 0)

																																								__e.TailApply(PrimNS2Value(symshen_4cut), V4372, V4373, K3905, tmp14349)
																																								return

																																							}, 0)

																																							__e.TailApply(PrimNS2Value(symwhen), tmp14347, V4372, V4373, K3905, tmp14348)
																																							return

																																						}, 1)

																																						tmp14360 := PrimTail(Tm3914)

																																						__e.TailApply(tmp14343, tmp14360)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp14362 := PrimTail(Tm3922)

																																				tmp14363 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14362, V4372)

																																				__e.TailApply(tmp14341, tmp14363)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp14365 := PrimTail(Tm3925)

																																		tmp14366 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14365, V4372)

																																		__e.TailApply(tmp14339, tmp14366)
																																		return

																																	}, 1)

																																	tmp14367 := PrimHead(Tm3925)

																																	__e.TailApply(tmp14338, tmp14367)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp14369 := PrimTail(Tm3923)

																															tmp14370 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14369, V4372)

																															__e.TailApply(tmp14336, tmp14370)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp14372 := PrimHead(Tm3923)

																													tmp14373 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14372, V4372)

																													__e.TailApply(tmp14334, tmp14373)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp14375 := PrimHead(Tm3922)

																											tmp14376 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14375, V4372)

																											__e.TailApply(tmp14332, tmp14376)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp14378 := PrimTail(Tm3921)

																									tmp14379 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14378, V4372)

																									__e.TailApply(tmp14330, tmp14379)
																									return

																								}, 1)

																								tmp14380 := PrimHead(Tm3921)

																								__e.TailApply(tmp14329, tmp14380)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp14382 := PrimTail(Tm3915)

																						tmp14383 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14382, V4372)

																						__e.TailApply(tmp14327, tmp14383)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp14385 := PrimTail(Tm3919)

																				tmp14386 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14385, V4372)

																				__e.TailApply(tmp14325, tmp14386)
																				return

																			}, 1)

																			tmp14387 := PrimHead(Tm3919)

																			__e.TailApply(tmp14324, tmp14387)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp14389 := PrimTail(Tm3918)

																	tmp14390 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14389, V4372)

																	__e.TailApply(tmp14322, tmp14390)
																	return

																}, 1)

																tmp14391 := PrimHead(Tm3918)

																__e.TailApply(tmp14321, tmp14391)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14393 := PrimTail(Tm3916)

														tmp14394 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14393, V4372)

														__e.TailApply(tmp14319, tmp14394)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14396 := PrimHead(Tm3916)

												tmp14397 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14396, V4372)

												__e.TailApply(tmp14317, tmp14397)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14399 := PrimHead(Tm3915)

										tmp14400 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14399, V4372)

										__e.TailApply(tmp14315, tmp14400)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14402 := PrimHead(Tm3914)

								tmp14403 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14402, V4372)

								__e.TailApply(tmp14313, tmp14403)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14405 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4369, V4372)

						tmp14406 := Call(__e, tmp14311, tmp14405)

						ifres14310 = tmp14406

					} else {
						ifres14310 = False

					}

					__e.TailApply(tmp13980, ifres14310)
					return

				} else {
					__e.Return(C3910)
					return
				}

			}, 1)

			tmp14421 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4373)

			var ifres14409 Obj

			if True == tmp14421 {
				tmp14410 := MakeNative(func(__e *ControlFlow) {
					Tm3911 := __e.Get(1)
					_ = Tm3911
					tmp14418 := PrimEqual(Tm3911, Nil)

					if True == tmp14418 {
						tmp14412 := MakeNative(func(__e *ControlFlow) {
							Tm3912 := __e.Get(1)
							_ = Tm3912
							tmp14416 := PrimEqual(Tm3912, True)

							if True == tmp14416 {
								tmp14414 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp14414

								tmp14415 := MakeNative(func(__e *ControlFlow) {
									__e.TailApply(PrimNS2Value(symbind), V4370, Nil, V4372, V4373, K3905, V4375)
									return
								}, 0)

								__e.TailApply(PrimNS2Value(symshen_4cut), V4372, V4373, K3905, tmp14415)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14417 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4371, V4372)

						__e.TailApply(tmp14412, tmp14417)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14419 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4369, V4372)

				tmp14420 := Call(__e, tmp14410, tmp14419)

				ifres14409 = tmp14420

			} else {
				ifres14409 = False

			}

			__e.TailApply(tmp13978, ifres14409)
			return

		}, 1)

		tmp14422 := PrimNumberAdd(V4374, MakeNumber(1))

		__e.TailApply(tmp13977, tmp14422)
		return

	}, 7)

	tmp14423 := Call(__e, PrimNS2Value(symdef), symshen_4l_1rules, tmp13976)

	_ = tmp14423

	tmp14424 := MakeNative(func(__e *ControlFlow) {
		V4376 := __e.Get(1)
		_ = V4376
		V4377 := __e.Get(2)
		_ = V4377
		V4378 := __e.Get(3)
		_ = V4378
		V4379 := __e.Get(4)
		_ = V4379
		V4380 := __e.Get(5)
		_ = V4380
		V4381 := __e.Get(6)
		_ = V4381
		tmp14425 := MakeNative(func(__e *ControlFlow) {
			K3977 := __e.Get(1)
			_ = K3977
			tmp14426 := MakeNative(func(__e *ControlFlow) {
				C3981 := __e.Get(1)
				_ = C3981
				tmp14428 := PrimEqual(C3981, False)

				if True == tmp14428 {
					__e.TailApply(PrimNS2Value(symshen_4unlock), V4379, K3977)
					return
				} else {
					__e.Return(C3981)
					return
				}

			}, 1)

			tmp14476 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4379)

			var ifres14429 Obj

			if True == tmp14476 {
				tmp14430 := MakeNative(func(__e *ControlFlow) {
					Tm3982 := __e.Get(1)
					_ = Tm3982
					tmp14473 := PrimIsPair(Tm3982)

					if True == tmp14473 {
						tmp14432 := MakeNative(func(__e *ControlFlow) {
							Tm3983 := __e.Get(1)
							_ = Tm3983
							tmp14470 := PrimEqual(Tm3983, symdefine)

							if True == tmp14470 {
								tmp14434 := MakeNative(func(__e *ControlFlow) {
									Tm3984 := __e.Get(1)
									_ = Tm3984
									tmp14467 := PrimIsPair(Tm3984)

									if True == tmp14467 {
										tmp14436 := MakeNative(func(__e *ControlFlow) {
											F := __e.Get(1)
											_ = F
											tmp14437 := MakeNative(func(__e *ControlFlow) {
												X := __e.Get(1)
												_ = X
												tmp14438 := MakeNative(func(__e *ControlFlow) {
													SigxRules := __e.Get(1)
													_ = SigxRules
													tmp14439 := MakeNative(func(__e *ControlFlow) {
														Rules := __e.Get(1)
														_ = Rules
														tmp14440 := MakeNative(func(__e *ControlFlow) {
															FreshSig := __e.Get(1)
															_ = FreshSig
															tmp14441 := MakeNative(func(__e *ControlFlow) {
																Sig := __e.Get(1)
																_ = Sig
																tmp14442 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp14442

																tmp14443 := MakeNative(func(__e *ControlFlow) {
																	tmp14444 := PrimCons(F, X)

																	tmp14445 := Call(__e, PrimNS2Value(symshen_4sigxrules), tmp14444)

																	tmp14446 := MakeNative(func(__e *ControlFlow) {
																		tmp14447 := Call(__e, PrimNS2Value(symshen_4lazyderef), SigxRules, V4378)

																		tmp14448 := Call(__e, PrimNS2Value(symfst), tmp14447)

																		tmp14449 := MakeNative(func(__e *ControlFlow) {
																			tmp14450 := Call(__e, PrimNS2Value(symshen_4lazyderef), SigxRules, V4378)

																			tmp14451 := Call(__e, PrimNS2Value(symsnd), tmp14450)

																			tmp14452 := MakeNative(func(__e *ControlFlow) {
																				tmp14453 := Call(__e, PrimNS2Value(symshen_4deref), Sig, V4378)

																				tmp14454 := Call(__e, PrimNS2Value(symshen_4freshen_1sig), tmp14453)

																				tmp14455 := MakeNative(func(__e *ControlFlow) {
																					tmp14456 := MakeNative(func(__e *ControlFlow) {
																						__e.TailApply(PrimNS2Value(symis), Sig, V4377, V4378, V4379, K3977, V4381)
																						return
																					}, 0)

																					__e.TailApply(PrimNS2Value(symshen_4t_d_1rules), F, Rules, FreshSig, MakeNumber(1), V4378, V4379, K3977, tmp14456)
																					return

																				}, 0)

																				__e.TailApply(PrimNS2Value(symbind), FreshSig, tmp14454, V4378, V4379, K3977, tmp14455)
																				return

																			}, 0)

																			__e.TailApply(PrimNS2Value(symbind), Rules, tmp14451, V4378, V4379, K3977, tmp14452)
																			return

																		}, 0)

																		__e.TailApply(PrimNS2Value(symbind), Sig, tmp14448, V4378, V4379, K3977, tmp14449)
																		return

																	}, 0)

																	__e.TailApply(PrimNS2Value(symbind), SigxRules, tmp14445, V4378, V4379, K3977, tmp14446)
																	return

																}, 0)

																tmp14457 := Call(__e, PrimNS2Value(symshen_4cut), V4378, V4379, K3977, tmp14443)

																__e.TailApply(PrimNS2Value(symshen_4gc), V4378, tmp14457)
																return

															}, 1)

															tmp14458 := Call(__e, PrimNS2Value(symshen_4newpv), V4378)

															tmp14459 := Call(__e, tmp14441, tmp14458)

															__e.TailApply(PrimNS2Value(symshen_4gc), V4378, tmp14459)
															return

														}, 1)

														tmp14460 := Call(__e, PrimNS2Value(symshen_4newpv), V4378)

														tmp14461 := Call(__e, tmp14440, tmp14460)

														__e.TailApply(PrimNS2Value(symshen_4gc), V4378, tmp14461)
														return

													}, 1)

													tmp14462 := Call(__e, PrimNS2Value(symshen_4newpv), V4378)

													tmp14463 := Call(__e, tmp14439, tmp14462)

													__e.TailApply(PrimNS2Value(symshen_4gc), V4378, tmp14463)
													return

												}, 1)

												tmp14464 := Call(__e, PrimNS2Value(symshen_4newpv), V4378)

												__e.TailApply(tmp14438, tmp14464)
												return

											}, 1)

											tmp14465 := PrimTail(Tm3984)

											__e.TailApply(tmp14437, tmp14465)
											return

										}, 1)

										tmp14466 := PrimHead(Tm3984)

										__e.TailApply(tmp14436, tmp14466)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14468 := PrimTail(Tm3982)

								tmp14469 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14468, V4378)

								__e.TailApply(tmp14434, tmp14469)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14471 := PrimHead(Tm3982)

						tmp14472 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14471, V4378)

						__e.TailApply(tmp14432, tmp14472)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14474 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4376, V4378)

				tmp14475 := Call(__e, tmp14430, tmp14474)

				ifres14429 = tmp14475

			} else {
				ifres14429 = False

			}

			__e.TailApply(tmp14426, ifres14429)
			return

		}, 1)

		tmp14477 := PrimNumberAdd(V4380, MakeNumber(1))

		__e.TailApply(tmp14425, tmp14477)
		return

	}, 6)

	tmp14478 := Call(__e, PrimNS2Value(symdef), symshen_4t_d, tmp14424)

	_ = tmp14478

	tmp14479 := MakeNative(func(__e *ControlFlow) {
		V4382 := __e.Get(1)
		_ = V4382
		tmp14480 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_5sig_drules_6), X)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symcompile), tmp14480, V4382)
		return

	}, 1)

	tmp14481 := Call(__e, PrimNS2Value(symdef), symshen_4sigxrules, tmp14479)

	_ = tmp14481

	tmp14482 := MakeNative(func(__e *ControlFlow) {
		V4383 := __e.Get(1)
		_ = V4383
		tmp14483 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp14485 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp14485 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp14515 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V4383)

		var ifres14486 Obj

		if True == tmp14515 {
			tmp14488 := MakeNative(func(__e *ControlFlow) {
				F := __e.Get(1)
				_ = F
				tmp14489 := MakeNative(func(__e *ControlFlow) {
					News3986 := __e.Get(1)
					_ = News3986
					tmp14511 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News3986, sym_i)

					if True == tmp14511 {
						tmp14491 := MakeNative(func(__e *ControlFlow) {
							News3987 := __e.Get(1)
							_ = News3987
							tmp14492 := MakeNative(func(__e *ControlFlow) {
								Parseshen_4_5signature_6 := __e.Get(1)
								_ = Parseshen_4_5signature_6
								tmp14508 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5signature_6)

								if True == tmp14508 {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								} else {
									tmp14507 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5signature_6, sym_j)

									if True == tmp14507 {
										tmp14495 := MakeNative(func(__e *ControlFlow) {
											News3988 := __e.Get(1)
											_ = News3988
											tmp14496 := MakeNative(func(__e *ControlFlow) {
												Parseshen_4_5rules_d_6 := __e.Get(1)
												_ = Parseshen_4_5rules_d_6
												tmp14504 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rules_d_6)

												if True == tmp14504 {
													__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
													return
												} else {
													tmp14498 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5rules_d_6)

													tmp14499 := MakeNative(func(__e *ControlFlow) {
														Rectified := __e.Get(1)
														_ = Rectified
														tmp14500 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rules_d_6)

														__e.TailApply(PrimNS2Value(sym_8p), Rectified, tmp14500)
														return

													}, 1)

													tmp14501 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5signature_6)

													tmp14502 := Call(__e, PrimNS2Value(symshen_4rectify_1type), tmp14501)

													tmp14503 := Call(__e, tmp14499, tmp14502)

													__e.TailApply(PrimNS2Value(symshen_4comb), tmp14498, tmp14503)
													return

												}

											}, 1)

											tmp14505 := Call(__e, PrimNS2Value(symshen_4_5rules_d_6), News3988)

											__e.TailApply(tmp14496, tmp14505)
											return

										}, 1)

										tmp14506 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5signature_6)

										__e.TailApply(tmp14495, tmp14506)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}

							}, 1)

							tmp14509 := Call(__e, PrimNS2Value(symshen_4_5signature_6), News3987)

							__e.TailApply(tmp14492, tmp14509)
							return

						}, 1)

						tmp14510 := Call(__e, PrimNS2Value(symshen_4tls), News3986)

						__e.TailApply(tmp14491, tmp14510)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp14512 := Call(__e, PrimNS2Value(symshen_4tls), V4383)

				__e.TailApply(tmp14489, tmp14512)
				return

			}, 1)

			tmp14513 := Call(__e, PrimNS2Value(symshen_4hds), V4383)

			tmp14514 := Call(__e, tmp14488, tmp14513)

			ifres14486 = tmp14514

		} else {
			tmp14487 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres14486 = tmp14487

		}

		__e.TailApply(tmp14483, ifres14486)
		return

	}, 1)

	tmp14516 := Call(__e, PrimNS2Value(symdef), symshen_4_5sig_drules_6, tmp14482)

	_ = tmp14516

	tmp14517 := MakeNative(func(__e *ControlFlow) {
		V4384 := __e.Get(1)
		_ = V4384
		tmp14518 := MakeNative(func(__e *ControlFlow) {
			Vs := __e.Get(1)
			_ = Vs
			tmp14519 := MakeNative(func(__e *ControlFlow) {
				Assoc := __e.Get(1)
				_ = Assoc
				__e.TailApply(PrimNS2Value(symshen_4freshen_1type), Assoc, V4384)
				return
			}, 1)

			tmp14520 := MakeNative(func(__e *ControlFlow) {
				V := __e.Get(1)
				_ = V
				tmp14521 := Call(__e, PrimNS2Value(symconcat), sym_e, V)

				tmp14522 := Call(__e, PrimNS2Value(symshen_4freshterm), tmp14521)

				__e.Return(PrimCons(V, tmp14522))
				return

			}, 1)

			tmp14523 := Call(__e, PrimNS2Value(symmap), tmp14520, Vs)

			__e.TailApply(tmp14519, tmp14523)
			return

		}, 1)

		tmp14524 := Call(__e, PrimNS2Value(symshen_4extract_1vars), V4384)

		__e.TailApply(tmp14518, tmp14524)
		return

	}, 1)

	tmp14525 := Call(__e, PrimNS2Value(symdef), symshen_4freshen_1sig, tmp14517)

	_ = tmp14525

	tmp14526 := MakeNative(func(__e *ControlFlow) {
		V4385 := __e.Get(1)
		_ = V4385
		V4386 := __e.Get(2)
		_ = V4386
		tmp14540 := PrimEqual(Nil, V4385)

		if True == tmp14540 {
			__e.Return(V4386)
			return
		} else {
			tmp14539 := PrimIsPair(V4385)

			var ifres14535 Obj

			if True == tmp14539 {
				tmp14537 := PrimHead(V4385)

				tmp14538 := PrimIsPair(tmp14537)

				var ifres14536 Obj

				if True == tmp14538 {
					ifres14536 = True

				} else {
					ifres14536 = False

				}

				ifres14535 = ifres14536

			} else {
				ifres14535 = False

			}

			if True == ifres14535 {
				tmp14529 := PrimTail(V4385)

				tmp14530 := PrimHead(V4385)

				tmp14531 := PrimTail(tmp14530)

				tmp14532 := PrimHead(V4385)

				tmp14533 := PrimHead(tmp14532)

				tmp14534 := Call(__e, PrimNS2Value(symsubst), tmp14531, tmp14533, V4386)

				__e.TailApply(PrimNS2Value(symshen_4freshen_1type), tmp14529, tmp14534)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4freshen_1type)
				return
			}

		}

	}, 2)

	tmp14541 := Call(__e, PrimNS2Value(symdef), symshen_4freshen_1type, tmp14526)

	_ = tmp14541

	tmp14542 := MakeNative(func(__e *ControlFlow) {
		V4387 := __e.Get(1)
		_ = V4387
		tmp14543 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp14557 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp14557 {
				tmp14545 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp14547 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp14547 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp14548 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5rule_d_6 := __e.Get(1)
					_ = Parseshen_4_5rule_d_6
					tmp14554 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rule_d_6)

					if True == tmp14554 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp14550 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5rule_d_6)

						tmp14551 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rule_d_6)

						tmp14552 := Call(__e, PrimNS2Value(symshen_4linearise), tmp14551)

						tmp14553 := PrimCons(tmp14552, Nil)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp14550, tmp14553)
						return

					}

				}, 1)

				tmp14555 := Call(__e, PrimNS2Value(symshen_4_5rule_d_6), V4387)

				tmp14556 := Call(__e, tmp14548, tmp14555)

				__e.TailApply(tmp14545, tmp14556)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp14558 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5rule_d_6 := __e.Get(1)
			_ = Parseshen_4_5rule_d_6
			tmp14569 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rule_d_6)

			if True == tmp14569 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp14560 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5rules_d_6 := __e.Get(1)
					_ = Parseshen_4_5rules_d_6
					tmp14567 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rules_d_6)

					if True == tmp14567 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp14562 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5rules_d_6)

						tmp14563 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rule_d_6)

						tmp14564 := Call(__e, PrimNS2Value(symshen_4linearise), tmp14563)

						tmp14565 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rules_d_6)

						tmp14566 := PrimCons(tmp14564, tmp14565)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp14562, tmp14566)
						return

					}

				}, 1)

				tmp14568 := Call(__e, PrimNS2Value(symshen_4_5rules_d_6), Parseshen_4_5rule_d_6)

				__e.TailApply(tmp14560, tmp14568)
				return

			}

		}, 1)

		tmp14570 := Call(__e, PrimNS2Value(symshen_4_5rule_d_6), V4387)

		tmp14571 := Call(__e, tmp14558, tmp14570)

		__e.TailApply(tmp14543, tmp14571)
		return

	}, 1)

	tmp14572 := Call(__e, PrimNS2Value(symdef), symshen_4_5rules_d_6, tmp14542)

	_ = tmp14572

	tmp14573 := MakeNative(func(__e *ControlFlow) {
		V4388 := __e.Get(1)
		_ = V4388
		tmp14574 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp14654 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp14654 {
				tmp14576 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp14621 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp14621 {
						tmp14578 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp14601 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp14601 {
								tmp14580 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp14582 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

									if True == tmp14582 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp14583 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5patterns_6 := __e.Get(1)
									_ = Parseshen_4_5patterns_6
									tmp14598 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

									if True == tmp14598 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp14597 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)

										if True == tmp14597 {
											tmp14586 := MakeNative(func(__e *ControlFlow) {
												News4001 := __e.Get(1)
												_ = News4001
												tmp14595 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News4001)

												if True == tmp14595 {
													tmp14588 := MakeNative(func(__e *ControlFlow) {
														Action := __e.Get(1)
														_ = Action
														tmp14589 := MakeNative(func(__e *ControlFlow) {
															News4002 := __e.Get(1)
															_ = News4002
															tmp14590 := Call(__e, PrimNS2Value(symshen_4in_1_6), News4002)

															tmp14591 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

															tmp14592 := Call(__e, PrimNS2Value(sym_8p), tmp14591, Action)

															__e.TailApply(PrimNS2Value(symshen_4comb), tmp14590, tmp14592)
															return

														}, 1)

														tmp14593 := Call(__e, PrimNS2Value(symshen_4tls), News4001)

														__e.TailApply(tmp14589, tmp14593)
														return

													}, 1)

													tmp14594 := Call(__e, PrimNS2Value(symshen_4hds), News4001)

													__e.TailApply(tmp14588, tmp14594)
													return

												} else {
													__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp14596 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

											__e.TailApply(tmp14586, tmp14596)
											return

										} else {
											__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
											return
										}

									}

								}, 1)

								tmp14599 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V4388)

								tmp14600 := Call(__e, tmp14583, tmp14599)

								__e.TailApply(tmp14580, tmp14600)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp14602 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5patterns_6 := __e.Get(1)
							_ = Parseshen_4_5patterns_6
							tmp14618 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

							if True == tmp14618 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp14617 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)

								if True == tmp14617 {
									tmp14605 := MakeNative(func(__e *ControlFlow) {
										News3999 := __e.Get(1)
										_ = News3999
										tmp14615 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3999)

										if True == tmp14615 {
											tmp14607 := MakeNative(func(__e *ControlFlow) {
												Action := __e.Get(1)
												_ = Action
												tmp14608 := MakeNative(func(__e *ControlFlow) {
													News4000 := __e.Get(1)
													_ = News4000
													tmp14609 := Call(__e, PrimNS2Value(symshen_4in_1_6), News4000)

													tmp14610 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

													tmp14611 := Call(__e, PrimNS2Value(symshen_4correct), Action)

													tmp14612 := Call(__e, PrimNS2Value(sym_8p), tmp14610, tmp14611)

													__e.TailApply(PrimNS2Value(symshen_4comb), tmp14609, tmp14612)
													return

												}, 1)

												tmp14613 := Call(__e, PrimNS2Value(symshen_4tls), News3999)

												__e.TailApply(tmp14608, tmp14613)
												return

											}, 1)

											tmp14614 := Call(__e, PrimNS2Value(symshen_4hds), News3999)

											__e.TailApply(tmp14607, tmp14614)
											return

										} else {
											__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp14616 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

									__e.TailApply(tmp14605, tmp14616)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}

						}, 1)

						tmp14619 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V4388)

						tmp14620 := Call(__e, tmp14602, tmp14619)

						__e.TailApply(tmp14578, tmp14620)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp14622 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5patterns_6 := __e.Get(1)
					_ = Parseshen_4_5patterns_6
					tmp14651 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

					if True == tmp14651 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp14650 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)

						if True == tmp14650 {
							tmp14625 := MakeNative(func(__e *ControlFlow) {
								News3995 := __e.Get(1)
								_ = News3995
								tmp14648 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3995)

								if True == tmp14648 {
									tmp14627 := MakeNative(func(__e *ControlFlow) {
										Action := __e.Get(1)
										_ = Action
										tmp14628 := MakeNative(func(__e *ControlFlow) {
											News3996 := __e.Get(1)
											_ = News3996
											tmp14645 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News3996, symwhere)

											if True == tmp14645 {
												tmp14630 := MakeNative(func(__e *ControlFlow) {
													News3997 := __e.Get(1)
													_ = News3997
													tmp14643 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3997)

													if True == tmp14643 {
														tmp14632 := MakeNative(func(__e *ControlFlow) {
															Guard := __e.Get(1)
															_ = Guard
															tmp14633 := MakeNative(func(__e *ControlFlow) {
																News3998 := __e.Get(1)
																_ = News3998
																tmp14634 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3998)

																tmp14635 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

																tmp14636 := PrimCons(Action, Nil)

																tmp14637 := PrimCons(Guard, tmp14636)

																tmp14638 := PrimCons(symwhere, tmp14637)

																tmp14639 := Call(__e, PrimNS2Value(symshen_4correct), tmp14638)

																tmp14640 := Call(__e, PrimNS2Value(sym_8p), tmp14635, tmp14639)

																__e.TailApply(PrimNS2Value(symshen_4comb), tmp14634, tmp14640)
																return

															}, 1)

															tmp14641 := Call(__e, PrimNS2Value(symshen_4tls), News3997)

															__e.TailApply(tmp14633, tmp14641)
															return

														}, 1)

														tmp14642 := Call(__e, PrimNS2Value(symshen_4hds), News3997)

														__e.TailApply(tmp14632, tmp14642)
														return

													} else {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp14644 := Call(__e, PrimNS2Value(symshen_4tls), News3996)

												__e.TailApply(tmp14630, tmp14644)
												return

											} else {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp14646 := Call(__e, PrimNS2Value(symshen_4tls), News3995)

										__e.TailApply(tmp14628, tmp14646)
										return

									}, 1)

									tmp14647 := Call(__e, PrimNS2Value(symshen_4hds), News3995)

									__e.TailApply(tmp14627, tmp14647)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp14649 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

							__e.TailApply(tmp14625, tmp14649)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}

				}, 1)

				tmp14652 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V4388)

				tmp14653 := Call(__e, tmp14622, tmp14652)

				__e.TailApply(tmp14576, tmp14653)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp14655 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5patterns_6 := __e.Get(1)
			_ = Parseshen_4_5patterns_6
			tmp14683 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

			if True == tmp14683 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp14682 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)

				if True == tmp14682 {
					tmp14658 := MakeNative(func(__e *ControlFlow) {
						News3991 := __e.Get(1)
						_ = News3991
						tmp14680 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3991)

						if True == tmp14680 {
							tmp14660 := MakeNative(func(__e *ControlFlow) {
								Action := __e.Get(1)
								_ = Action
								tmp14661 := MakeNative(func(__e *ControlFlow) {
									News3992 := __e.Get(1)
									_ = News3992
									tmp14677 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News3992, symwhere)

									if True == tmp14677 {
										tmp14663 := MakeNative(func(__e *ControlFlow) {
											News3993 := __e.Get(1)
											_ = News3993
											tmp14675 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News3993)

											if True == tmp14675 {
												tmp14665 := MakeNative(func(__e *ControlFlow) {
													Guard := __e.Get(1)
													_ = Guard
													tmp14666 := MakeNative(func(__e *ControlFlow) {
														News3994 := __e.Get(1)
														_ = News3994
														tmp14667 := Call(__e, PrimNS2Value(symshen_4in_1_6), News3994)

														tmp14668 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

														tmp14669 := PrimCons(Action, Nil)

														tmp14670 := PrimCons(Guard, tmp14669)

														tmp14671 := PrimCons(symwhere, tmp14670)

														tmp14672 := Call(__e, PrimNS2Value(sym_8p), tmp14668, tmp14671)

														__e.TailApply(PrimNS2Value(symshen_4comb), tmp14667, tmp14672)
														return

													}, 1)

													tmp14673 := Call(__e, PrimNS2Value(symshen_4tls), News3993)

													__e.TailApply(tmp14666, tmp14673)
													return

												}, 1)

												tmp14674 := Call(__e, PrimNS2Value(symshen_4hds), News3993)

												__e.TailApply(tmp14665, tmp14674)
												return

											} else {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp14676 := Call(__e, PrimNS2Value(symshen_4tls), News3992)

										__e.TailApply(tmp14663, tmp14676)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp14678 := Call(__e, PrimNS2Value(symshen_4tls), News3991)

								__e.TailApply(tmp14661, tmp14678)
								return

							}, 1)

							tmp14679 := Call(__e, PrimNS2Value(symshen_4hds), News3991)

							__e.TailApply(tmp14660, tmp14679)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp14681 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

					__e.TailApply(tmp14658, tmp14681)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp14684 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V4388)

		tmp14685 := Call(__e, tmp14655, tmp14684)

		__e.TailApply(tmp14574, tmp14685)
		return

	}, 1)

	tmp14686 := Call(__e, PrimNS2Value(symdef), symshen_4_5rule_d_6, tmp14573)

	_ = tmp14686

	tmp14687 := MakeNative(func(__e *ControlFlow) {
		V4389 := __e.Get(1)
		_ = V4389
		tmp14835 := PrimIsPair(V4389)

		var ifres14779 Obj

		if True == tmp14835 {
			tmp14833 := PrimHead(V4389)

			tmp14834 := PrimEqual(symwhere, tmp14833)

			var ifres14781 Obj

			if True == tmp14834 {
				tmp14831 := PrimTail(V4389)

				tmp14832 := PrimIsPair(tmp14831)

				var ifres14783 Obj

				if True == tmp14832 {
					tmp14828 := PrimTail(V4389)

					tmp14829 := PrimTail(tmp14828)

					tmp14830 := PrimIsPair(tmp14829)

					var ifres14785 Obj

					if True == tmp14830 {
						tmp14824 := PrimTail(V4389)

						tmp14825 := PrimTail(tmp14824)

						tmp14826 := PrimHead(tmp14825)

						tmp14827 := PrimIsPair(tmp14826)

						var ifres14787 Obj

						if True == tmp14827 {
							tmp14819 := PrimTail(V4389)

							tmp14820 := PrimTail(tmp14819)

							tmp14821 := PrimHead(tmp14820)

							tmp14822 := PrimHead(tmp14821)

							tmp14823 := PrimEqual(symfail_1if, tmp14822)

							var ifres14789 Obj

							if True == tmp14823 {
								tmp14814 := PrimTail(V4389)

								tmp14815 := PrimTail(tmp14814)

								tmp14816 := PrimHead(tmp14815)

								tmp14817 := PrimTail(tmp14816)

								tmp14818 := PrimIsPair(tmp14817)

								var ifres14791 Obj

								if True == tmp14818 {
									tmp14808 := PrimTail(V4389)

									tmp14809 := PrimTail(tmp14808)

									tmp14810 := PrimHead(tmp14809)

									tmp14811 := PrimTail(tmp14810)

									tmp14812 := PrimTail(tmp14811)

									tmp14813 := PrimIsPair(tmp14812)

									var ifres14793 Obj

									if True == tmp14813 {
										tmp14801 := PrimTail(V4389)

										tmp14802 := PrimTail(tmp14801)

										tmp14803 := PrimHead(tmp14802)

										tmp14804 := PrimTail(tmp14803)

										tmp14805 := PrimTail(tmp14804)

										tmp14806 := PrimTail(tmp14805)

										tmp14807 := PrimEqual(Nil, tmp14806)

										var ifres14795 Obj

										if True == tmp14807 {
											tmp14797 := PrimTail(V4389)

											tmp14798 := PrimTail(tmp14797)

											tmp14799 := PrimTail(tmp14798)

											tmp14800 := PrimEqual(Nil, tmp14799)

											var ifres14796 Obj

											if True == tmp14800 {
												ifres14796 = True

											} else {
												ifres14796 = False

											}

											ifres14795 = ifres14796

										} else {
											ifres14795 = False

										}

										var ifres14794 Obj

										if True == ifres14795 {
											ifres14794 = True

										} else {
											ifres14794 = False

										}

										ifres14793 = ifres14794

									} else {
										ifres14793 = False

									}

									var ifres14792 Obj

									if True == ifres14793 {
										ifres14792 = True

									} else {
										ifres14792 = False

									}

									ifres14791 = ifres14792

								} else {
									ifres14791 = False

								}

								var ifres14790 Obj

								if True == ifres14791 {
									ifres14790 = True

								} else {
									ifres14790 = False

								}

								ifres14789 = ifres14790

							} else {
								ifres14789 = False

							}

							var ifres14788 Obj

							if True == ifres14789 {
								ifres14788 = True

							} else {
								ifres14788 = False

							}

							ifres14787 = ifres14788

						} else {
							ifres14787 = False

						}

						var ifres14786 Obj

						if True == ifres14787 {
							ifres14786 = True

						} else {
							ifres14786 = False

						}

						ifres14785 = ifres14786

					} else {
						ifres14785 = False

					}

					var ifres14784 Obj

					if True == ifres14785 {
						ifres14784 = True

					} else {
						ifres14784 = False

					}

					ifres14783 = ifres14784

				} else {
					ifres14783 = False

				}

				var ifres14782 Obj

				if True == ifres14783 {
					ifres14782 = True

				} else {
					ifres14782 = False

				}

				ifres14781 = ifres14782

			} else {
				ifres14781 = False

			}

			var ifres14780 Obj

			if True == ifres14781 {
				ifres14780 = True

			} else {
				ifres14780 = False

			}

			ifres14779 = ifres14780

		} else {
			ifres14779 = False

		}

		if True == ifres14779 {
			tmp14762 := PrimTail(V4389)

			tmp14763 := PrimHead(tmp14762)

			tmp14764 := PrimTail(V4389)

			tmp14765 := PrimTail(tmp14764)

			tmp14766 := PrimHead(tmp14765)

			tmp14767 := PrimTail(tmp14766)

			tmp14768 := PrimCons(tmp14767, Nil)

			tmp14769 := PrimCons(symnot, tmp14768)

			tmp14770 := PrimCons(tmp14769, Nil)

			tmp14771 := PrimCons(tmp14763, tmp14770)

			tmp14772 := PrimCons(symand, tmp14771)

			tmp14773 := PrimTail(V4389)

			tmp14774 := PrimTail(tmp14773)

			tmp14775 := PrimHead(tmp14774)

			tmp14776 := PrimTail(tmp14775)

			tmp14777 := PrimTail(tmp14776)

			tmp14778 := PrimCons(tmp14772, tmp14777)

			__e.Return(PrimCons(symwhere, tmp14778))
			return

		} else {
			tmp14761 := PrimIsPair(V4389)

			var ifres14742 Obj

			if True == tmp14761 {
				tmp14759 := PrimHead(V4389)

				tmp14760 := PrimEqual(symwhere, tmp14759)

				var ifres14744 Obj

				if True == tmp14760 {
					tmp14757 := PrimTail(V4389)

					tmp14758 := PrimIsPair(tmp14757)

					var ifres14746 Obj

					if True == tmp14758 {
						tmp14754 := PrimTail(V4389)

						tmp14755 := PrimTail(tmp14754)

						tmp14756 := PrimIsPair(tmp14755)

						var ifres14748 Obj

						if True == tmp14756 {
							tmp14750 := PrimTail(V4389)

							tmp14751 := PrimTail(tmp14750)

							tmp14752 := PrimTail(tmp14751)

							tmp14753 := PrimEqual(Nil, tmp14752)

							var ifres14749 Obj

							if True == tmp14753 {
								ifres14749 = True

							} else {
								ifres14749 = False

							}

							ifres14748 = ifres14749

						} else {
							ifres14748 = False

						}

						var ifres14747 Obj

						if True == ifres14748 {
							ifres14747 = True

						} else {
							ifres14747 = False

						}

						ifres14746 = ifres14747

					} else {
						ifres14746 = False

					}

					var ifres14745 Obj

					if True == ifres14746 {
						ifres14745 = True

					} else {
						ifres14745 = False

					}

					ifres14744 = ifres14745

				} else {
					ifres14744 = False

				}

				var ifres14743 Obj

				if True == ifres14744 {
					ifres14743 = True

				} else {
					ifres14743 = False

				}

				ifres14742 = ifres14743

			} else {
				ifres14742 = False

			}

			if True == ifres14742 {
				tmp14725 := PrimTail(V4389)

				tmp14726 := PrimHead(tmp14725)

				tmp14727 := PrimTail(V4389)

				tmp14728 := PrimTail(tmp14727)

				tmp14729 := PrimHead(tmp14728)

				tmp14730 := PrimCons(symfail, Nil)

				tmp14731 := PrimCons(tmp14730, Nil)

				tmp14732 := PrimCons(tmp14729, tmp14731)

				tmp14733 := PrimCons(sym_a, tmp14732)

				tmp14734 := PrimCons(tmp14733, Nil)

				tmp14735 := PrimCons(symnot, tmp14734)

				tmp14736 := PrimCons(tmp14735, Nil)

				tmp14737 := PrimCons(tmp14726, tmp14736)

				tmp14738 := PrimCons(symand, tmp14737)

				tmp14739 := PrimTail(V4389)

				tmp14740 := PrimTail(tmp14739)

				tmp14741 := PrimCons(tmp14738, tmp14740)

				__e.Return(PrimCons(symwhere, tmp14741))
				return

			} else {
				tmp14724 := PrimIsPair(V4389)

				var ifres14705 Obj

				if True == tmp14724 {
					tmp14722 := PrimHead(V4389)

					tmp14723 := PrimEqual(symfail_1if, tmp14722)

					var ifres14707 Obj

					if True == tmp14723 {
						tmp14720 := PrimTail(V4389)

						tmp14721 := PrimIsPair(tmp14720)

						var ifres14709 Obj

						if True == tmp14721 {
							tmp14717 := PrimTail(V4389)

							tmp14718 := PrimTail(tmp14717)

							tmp14719 := PrimIsPair(tmp14718)

							var ifres14711 Obj

							if True == tmp14719 {
								tmp14713 := PrimTail(V4389)

								tmp14714 := PrimTail(tmp14713)

								tmp14715 := PrimTail(tmp14714)

								tmp14716 := PrimEqual(Nil, tmp14715)

								var ifres14712 Obj

								if True == tmp14716 {
									ifres14712 = True

								} else {
									ifres14712 = False

								}

								ifres14711 = ifres14712

							} else {
								ifres14711 = False

							}

							var ifres14710 Obj

							if True == ifres14711 {
								ifres14710 = True

							} else {
								ifres14710 = False

							}

							ifres14709 = ifres14710

						} else {
							ifres14709 = False

						}

						var ifres14708 Obj

						if True == ifres14709 {
							ifres14708 = True

						} else {
							ifres14708 = False

						}

						ifres14707 = ifres14708

					} else {
						ifres14707 = False

					}

					var ifres14706 Obj

					if True == ifres14707 {
						ifres14706 = True

					} else {
						ifres14706 = False

					}

					ifres14705 = ifres14706

				} else {
					ifres14705 = False

				}

				if True == ifres14705 {
					tmp14699 := PrimTail(V4389)

					tmp14700 := PrimCons(tmp14699, Nil)

					tmp14701 := PrimCons(symnot, tmp14700)

					tmp14702 := PrimTail(V4389)

					tmp14703 := PrimTail(tmp14702)

					tmp14704 := PrimCons(tmp14701, tmp14703)

					__e.Return(PrimCons(symwhere, tmp14704))
					return

				} else {
					tmp14691 := PrimCons(symfail, Nil)

					tmp14692 := PrimCons(tmp14691, Nil)

					tmp14693 := PrimCons(V4389, tmp14692)

					tmp14694 := PrimCons(sym_a, tmp14693)

					tmp14695 := PrimCons(tmp14694, Nil)

					tmp14696 := PrimCons(symnot, tmp14695)

					tmp14697 := PrimCons(V4389, Nil)

					tmp14698 := PrimCons(tmp14696, tmp14697)

					__e.Return(PrimCons(symwhere, tmp14698))
					return

				}

			}

		}

	}, 1)

	tmp14836 := Call(__e, PrimNS2Value(symdef), symshen_4correct, tmp14687)

	_ = tmp14836

	tmp14837 := MakeNative(func(__e *ControlFlow) {
		V4390 := __e.Get(1)
		_ = V4390
		V4391 := __e.Get(2)
		_ = V4391
		V4392 := __e.Get(3)
		_ = V4392
		V4393 := __e.Get(4)
		_ = V4393
		V4394 := __e.Get(5)
		_ = V4394
		V4395 := __e.Get(6)
		_ = V4395
		V4396 := __e.Get(7)
		_ = V4396
		V4397 := __e.Get(8)
		_ = V4397
		tmp14838 := MakeNative(func(__e *ControlFlow) {
			K4005 := __e.Get(1)
			_ = K4005
			tmp14839 := MakeNative(func(__e *ControlFlow) {
				C4011 := __e.Get(1)
				_ = C4011
				tmp14869 := PrimEqual(C4011, False)

				if True == tmp14869 {
					tmp14841 := MakeNative(func(__e *ControlFlow) {
						C4013 := __e.Get(1)
						_ = C4013
						tmp14843 := PrimEqual(C4013, False)

						if True == tmp14843 {
							__e.TailApply(PrimNS2Value(symshen_4unlock), V4395, K4005)
							return
						} else {
							__e.Return(C4013)
							return
						}

					}, 1)

					tmp14868 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4395)

					var ifres14844 Obj

					if True == tmp14868 {
						tmp14845 := MakeNative(func(__e *ControlFlow) {
							Tm4014 := __e.Get(1)
							_ = Tm4014
							tmp14865 := PrimIsPair(Tm4014)

							if True == tmp14865 {
								tmp14847 := MakeNative(func(__e *ControlFlow) {
									Rule := __e.Get(1)
									_ = Rule
									tmp14848 := MakeNative(func(__e *ControlFlow) {
										Rules := __e.Get(1)
										_ = Rules
										tmp14849 := MakeNative(func(__e *ControlFlow) {
											Fresh := __e.Get(1)
											_ = Fresh
											tmp14850 := Call(__e, PrimNS2Value(symshen_4incinfs))

											_ = tmp14850

											tmp14851 := Call(__e, PrimNS2Value(symshen_4deref), Rule, V4394)

											tmp14852 := Call(__e, PrimNS2Value(symshen_4freshen_1rule), tmp14851)

											tmp14853 := MakeNative(func(__e *ControlFlow) {
												tmp14854 := Call(__e, PrimNS2Value(symshen_4lazyderef), Fresh, V4394)

												tmp14855 := Call(__e, PrimNS2Value(symfst), tmp14854)

												tmp14856 := Call(__e, PrimNS2Value(symshen_4lazyderef), Fresh, V4394)

												tmp14857 := Call(__e, PrimNS2Value(symsnd), tmp14856)

												tmp14858 := MakeNative(func(__e *ControlFlow) {
													tmp14859 := MakeNative(func(__e *ControlFlow) {
														tmp14860 := PrimNumberAdd(V4393, MakeNumber(1))

														__e.TailApply(PrimNS2Value(symshen_4t_d_1rules), V4390, Rules, V4392, tmp14860, V4394, V4395, K4005, V4397)
														return

													}, 0)

													__e.TailApply(PrimNS2Value(symshen_4cut), V4394, V4395, K4005, tmp14859)
													return

												}, 0)

												__e.TailApply(PrimNS2Value(symshen_4t_d_1rule), V4390, V4393, tmp14855, tmp14857, V4392, V4394, V4395, K4005, tmp14858)
												return

											}, 0)

											tmp14861 := Call(__e, PrimNS2Value(symbind), Fresh, tmp14852, V4394, V4395, K4005, tmp14853)

											__e.TailApply(PrimNS2Value(symshen_4gc), V4394, tmp14861)
											return

										}, 1)

										tmp14862 := Call(__e, PrimNS2Value(symshen_4newpv), V4394)

										__e.TailApply(tmp14849, tmp14862)
										return

									}, 1)

									tmp14863 := PrimTail(Tm4014)

									__e.TailApply(tmp14848, tmp14863)
									return

								}, 1)

								tmp14864 := PrimHead(Tm4014)

								__e.TailApply(tmp14847, tmp14864)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14866 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4391, V4394)

						tmp14867 := Call(__e, tmp14845, tmp14866)

						ifres14844 = tmp14867

					} else {
						ifres14844 = False

					}

					__e.TailApply(tmp14841, ifres14844)
					return

				} else {
					__e.Return(C4011)
					return
				}

			}, 1)

			tmp14877 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4395)

			var ifres14870 Obj

			if True == tmp14877 {
				tmp14871 := MakeNative(func(__e *ControlFlow) {
					Tm4012 := __e.Get(1)
					_ = Tm4012
					tmp14874 := PrimEqual(Tm4012, Nil)

					if True == tmp14874 {
						tmp14873 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp14873

						__e.TailApply(PrimNS2Value(symthaw), V4397)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14875 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4391, V4394)

				tmp14876 := Call(__e, tmp14871, tmp14875)

				ifres14870 = tmp14876

			} else {
				ifres14870 = False

			}

			__e.TailApply(tmp14839, ifres14870)
			return

		}, 1)

		tmp14878 := PrimNumberAdd(V4396, MakeNumber(1))

		__e.TailApply(tmp14838, tmp14878)
		return

	}, 8)

	tmp14879 := Call(__e, PrimNS2Value(symdef), symshen_4t_d_1rules, tmp14837)

	_ = tmp14879

	tmp14880 := MakeNative(func(__e *ControlFlow) {
		V4398 := __e.Get(1)
		_ = V4398
		tmp14893 := Call(__e, PrimNS2Value(symtuple_2), V4398)

		if True == tmp14893 {
			tmp14882 := MakeNative(func(__e *ControlFlow) {
				Vs := __e.Get(1)
				_ = Vs
				tmp14883 := MakeNative(func(__e *ControlFlow) {
					Assoc := __e.Get(1)
					_ = Assoc
					tmp14884 := Call(__e, PrimNS2Value(symfst), V4398)

					tmp14885 := Call(__e, PrimNS2Value(symshen_4freshen), Assoc, tmp14884)

					tmp14886 := Call(__e, PrimNS2Value(symsnd), V4398)

					tmp14887 := Call(__e, PrimNS2Value(symshen_4freshen), Assoc, tmp14886)

					__e.TailApply(PrimNS2Value(sym_8p), tmp14885, tmp14887)
					return

				}, 1)

				tmp14888 := MakeNative(func(__e *ControlFlow) {
					V := __e.Get(1)
					_ = V
					tmp14889 := Call(__e, PrimNS2Value(symshen_4freshterm), V)

					__e.Return(PrimCons(V, tmp14889))
					return

				}, 1)

				tmp14890 := Call(__e, PrimNS2Value(symmap), tmp14888, Vs)

				__e.TailApply(tmp14883, tmp14890)
				return

			}, 1)

			tmp14891 := Call(__e, PrimNS2Value(symfst), V4398)

			tmp14892 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp14891)

			__e.TailApply(tmp14882, tmp14892)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4freshen_1rule)
			return
		}

	}, 1)

	tmp14894 := Call(__e, PrimNS2Value(symdef), symshen_4freshen_1rule, tmp14880)

	_ = tmp14894

	tmp14895 := MakeNative(func(__e *ControlFlow) {
		V4399 := __e.Get(1)
		_ = V4399
		V4400 := __e.Get(2)
		_ = V4400
		tmp14909 := PrimEqual(Nil, V4399)

		if True == tmp14909 {
			__e.Return(V4400)
			return
		} else {
			tmp14908 := PrimIsPair(V4399)

			var ifres14904 Obj

			if True == tmp14908 {
				tmp14906 := PrimHead(V4399)

				tmp14907 := PrimIsPair(tmp14906)

				var ifres14905 Obj

				if True == tmp14907 {
					ifres14905 = True

				} else {
					ifres14905 = False

				}

				ifres14904 = ifres14905

			} else {
				ifres14904 = False

			}

			if True == ifres14904 {
				tmp14898 := PrimTail(V4399)

				tmp14899 := PrimHead(V4399)

				tmp14900 := PrimHead(tmp14899)

				tmp14901 := PrimHead(V4399)

				tmp14902 := PrimTail(tmp14901)

				tmp14903 := Call(__e, PrimNS2Value(symshen_4beta), tmp14900, tmp14902, V4400)

				__e.TailApply(PrimNS2Value(symshen_4freshen), tmp14898, tmp14903)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4freshen)
				return
			}

		}

	}, 2)

	tmp14910 := Call(__e, PrimNS2Value(symdef), symshen_4freshen, tmp14895)

	_ = tmp14910

	tmp14911 := MakeNative(func(__e *ControlFlow) {
		V4401 := __e.Get(1)
		_ = V4401
		V4402 := __e.Get(2)
		_ = V4402
		V4403 := __e.Get(3)
		_ = V4403
		V4404 := __e.Get(4)
		_ = V4404
		V4405 := __e.Get(5)
		_ = V4405
		V4406 := __e.Get(6)
		_ = V4406
		V4407 := __e.Get(7)
		_ = V4407
		V4408 := __e.Get(8)
		_ = V4408
		V4409 := __e.Get(9)
		_ = V4409
		tmp14912 := MakeNative(func(__e *ControlFlow) {
			C4024 := __e.Get(1)
			_ = C4024
			tmp14925 := PrimEqual(C4024, False)

			if True == tmp14925 {
				tmp14924 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4407)

				if True == tmp14924 {
					tmp14915 := MakeNative(func(__e *ControlFlow) {
						Err := __e.Get(1)
						_ = Err
						tmp14916 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp14916

						tmp14917 := Call(__e, PrimNS2Value(symshen_4app), V4401, MakeString("\n"), symshen_4a)

						tmp14918 := PrimStringConcat(MakeString(" of "), tmp14917)

						tmp14919 := Call(__e, PrimNS2Value(symshen_4app), V4402, tmp14918, symshen_4a)

						tmp14920 := PrimStringConcat(MakeString("type error in rule "), tmp14919)

						tmp14921 := PrimSimpleError(tmp14920)

						tmp14922 := Call(__e, PrimNS2Value(symbind), Err, tmp14921, V4406, V4407, V4408, V4409)

						__e.TailApply(PrimNS2Value(symshen_4gc), V4406, tmp14922)
						return

					}, 1)

					tmp14923 := Call(__e, PrimNS2Value(symshen_4newpv), V4406)

					__e.TailApply(tmp14915, tmp14923)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C4024)
				return
			}

		}, 1)

		tmp14929 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4407)

		var ifres14926 Obj

		if True == tmp14929 {
			tmp14927 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp14927

			tmp14928 := Call(__e, PrimNS2Value(symshen_4t_d_1rule_1h), V4403, V4404, V4405, V4406, V4407, V4408, V4409)

			ifres14926 = tmp14928

		} else {
			ifres14926 = False

		}

		__e.TailApply(tmp14912, ifres14926)
		return

	}, 9)

	tmp14930 := Call(__e, PrimNS2Value(symdef), symshen_4t_d_1rule, tmp14911)

	_ = tmp14930

	tmp14931 := MakeNative(func(__e *ControlFlow) {
		V4410 := __e.Get(1)
		_ = V4410
		V4411 := __e.Get(2)
		_ = V4411
		V4412 := __e.Get(3)
		_ = V4412
		V4413 := __e.Get(4)
		_ = V4413
		V4414 := __e.Get(5)
		_ = V4414
		V4415 := __e.Get(6)
		_ = V4415
		V4416 := __e.Get(7)
		_ = V4416
		tmp14932 := MakeNative(func(__e *ControlFlow) {
			K4027 := __e.Get(1)
			_ = K4027
			tmp14933 := MakeNative(func(__e *ControlFlow) {
				C4032 := __e.Get(1)
				_ = C4032
				tmp14950 := PrimEqual(C4032, False)

				if True == tmp14950 {
					tmp14935 := MakeNative(func(__e *ControlFlow) {
						C4038 := __e.Get(1)
						_ = C4038
						tmp14937 := PrimEqual(C4038, False)

						if True == tmp14937 {
							__e.TailApply(PrimNS2Value(symshen_4unlock), V4414, K4027)
							return
						} else {
							__e.Return(C4038)
							return
						}

					}, 1)

					tmp14949 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4414)

					var ifres14938 Obj

					if True == tmp14949 {
						tmp14939 := MakeNative(func(__e *ControlFlow) {
							B := __e.Get(1)
							_ = B
							tmp14940 := MakeNative(func(__e *ControlFlow) {
								Hyps := __e.Get(1)
								_ = Hyps
								tmp14941 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp14941

								tmp14942 := MakeNative(func(__e *ControlFlow) {
									tmp14943 := MakeNative(func(__e *ControlFlow) {
										__e.TailApply(PrimNS2Value(symshen_4t_d_1correct), V4411, B, Hyps, V4413, V4414, K4027, V4416)
										return
									}, 0)

									__e.TailApply(PrimNS2Value(symshen_4cut), V4413, V4414, K4027, tmp14943)
									return

								}, 0)

								tmp14944 := Call(__e, PrimNS2Value(symshen_4t_d_1integrity), V4410, V4412, Hyps, B, V4413, V4414, K4027, tmp14942)

								__e.TailApply(PrimNS2Value(symshen_4gc), V4413, tmp14944)
								return

							}, 1)

							tmp14945 := Call(__e, PrimNS2Value(symshen_4newpv), V4413)

							tmp14946 := Call(__e, tmp14940, tmp14945)

							__e.TailApply(PrimNS2Value(symshen_4gc), V4413, tmp14946)
							return

						}, 1)

						tmp14947 := Call(__e, PrimNS2Value(symshen_4newpv), V4413)

						tmp14948 := Call(__e, tmp14939, tmp14947)

						ifres14938 = tmp14948

					} else {
						ifres14938 = False

					}

					__e.TailApply(tmp14935, ifres14938)
					return

				} else {
					__e.Return(C4032)
					return
				}

			}, 1)

			tmp14980 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4414)

			var ifres14951 Obj

			if True == tmp14980 {
				tmp14952 := MakeNative(func(__e *ControlFlow) {
					Tm4033 := __e.Get(1)
					_ = Tm4033
					tmp14977 := PrimEqual(Tm4033, Nil)

					if True == tmp14977 {
						tmp14954 := MakeNative(func(__e *ControlFlow) {
							Tm4034 := __e.Get(1)
							_ = Tm4034
							tmp14975 := PrimIsPair(Tm4034)

							if True == tmp14975 {
								tmp14956 := MakeNative(func(__e *ControlFlow) {
									Tm4035 := __e.Get(1)
									_ = Tm4035
									tmp14972 := PrimEqual(Tm4035, sym_1_1_6)

									if True == tmp14972 {
										tmp14958 := MakeNative(func(__e *ControlFlow) {
											Tm4036 := __e.Get(1)
											_ = Tm4036
											tmp14969 := PrimIsPair(Tm4036)

											if True == tmp14969 {
												tmp14960 := MakeNative(func(__e *ControlFlow) {
													A := __e.Get(1)
													_ = A
													tmp14961 := MakeNative(func(__e *ControlFlow) {
														Tm4037 := __e.Get(1)
														_ = Tm4037
														tmp14965 := PrimEqual(Tm4037, Nil)

														if True == tmp14965 {
															tmp14963 := Call(__e, PrimNS2Value(symshen_4incinfs))

															_ = tmp14963

															tmp14964 := MakeNative(func(__e *ControlFlow) {
																__e.TailApply(PrimNS2Value(symshen_4t_d_1correct), V4411, A, Nil, V4413, V4414, K4027, V4416)
																return
															}, 0)

															__e.TailApply(PrimNS2Value(symshen_4cut), V4413, V4414, K4027, tmp14964)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp14966 := PrimTail(Tm4036)

													tmp14967 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14966, V4413)

													__e.TailApply(tmp14961, tmp14967)
													return

												}, 1)

												tmp14968 := PrimHead(Tm4036)

												__e.TailApply(tmp14960, tmp14968)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14970 := PrimTail(Tm4034)

										tmp14971 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14970, V4413)

										__e.TailApply(tmp14958, tmp14971)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14973 := PrimHead(Tm4034)

								tmp14974 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp14973, V4413)

								__e.TailApply(tmp14956, tmp14974)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14976 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4412, V4413)

						__e.TailApply(tmp14954, tmp14976)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14978 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4410, V4413)

				tmp14979 := Call(__e, tmp14952, tmp14978)

				ifres14951 = tmp14979

			} else {
				ifres14951 = False

			}

			__e.TailApply(tmp14933, ifres14951)
			return

		}, 1)

		tmp14981 := PrimNumberAdd(V4415, MakeNumber(1))

		__e.TailApply(tmp14932, tmp14981)
		return

	}, 7)

	tmp14982 := Call(__e, PrimNS2Value(symdef), symshen_4t_d_1rule_1h, tmp14931)

	_ = tmp14982

	tmp14983 := MakeNative(func(__e *ControlFlow) {
		V4417 := __e.Get(1)
		_ = V4417
		V4418 := __e.Get(2)
		_ = V4418
		V4419 := __e.Get(3)
		_ = V4419
		V4420 := __e.Get(4)
		_ = V4420
		V4421 := __e.Get(5)
		_ = V4421
		V4422 := __e.Get(6)
		_ = V4422
		V4423 := __e.Get(7)
		_ = V4423
		tmp14984 := MakeNative(func(__e *ControlFlow) {
			K4041 := __e.Get(1)
			_ = K4041
			tmp14985 := MakeNative(func(__e *ControlFlow) {
				C4046 := __e.Get(1)
				_ = C4046
				tmp14995 := PrimEqual(C4046, False)

				if True == tmp14995 {
					tmp14987 := MakeNative(func(__e *ControlFlow) {
						C4052 := __e.Get(1)
						_ = C4052
						tmp14989 := PrimEqual(C4052, False)

						if True == tmp14989 {
							__e.TailApply(PrimNS2Value(symshen_4unlock), V4421, K4041)
							return
						} else {
							__e.Return(C4052)
							return
						}

					}, 1)

					tmp14994 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4421)

					var ifres14990 Obj

					if True == tmp14994 {
						tmp14991 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp14991

						tmp14992 := Call(__e, PrimNS2Value(symshen_4curry), V4417)

						tmp14993 := Call(__e, PrimNS2Value(symshen_4system_1S_1h), tmp14992, V4418, V4419, V4420, V4421, K4041, V4423)

						ifres14990 = tmp14993

					} else {
						ifres14990 = False

					}

					__e.TailApply(tmp14987, ifres14990)
					return

				} else {
					__e.Return(C4046)
					return
				}

			}, 1)

			tmp15040 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4421)

			var ifres14996 Obj

			if True == tmp15040 {
				tmp14997 := MakeNative(func(__e *ControlFlow) {
					Tm4047 := __e.Get(1)
					_ = Tm4047
					tmp15037 := PrimIsPair(Tm4047)

					if True == tmp15037 {
						tmp14999 := MakeNative(func(__e *ControlFlow) {
							Tm4048 := __e.Get(1)
							_ = Tm4048
							tmp15034 := PrimEqual(Tm4048, symwhere)

							if True == tmp15034 {
								tmp15001 := MakeNative(func(__e *ControlFlow) {
									Tm4049 := __e.Get(1)
									_ = Tm4049
									tmp15031 := PrimIsPair(Tm4049)

									if True == tmp15031 {
										tmp15003 := MakeNative(func(__e *ControlFlow) {
											G := __e.Get(1)
											_ = G
											tmp15004 := MakeNative(func(__e *ControlFlow) {
												Tm4050 := __e.Get(1)
												_ = Tm4050
												tmp15027 := PrimIsPair(Tm4050)

												if True == tmp15027 {
													tmp15006 := MakeNative(func(__e *ControlFlow) {
														R := __e.Get(1)
														_ = R
														tmp15007 := MakeNative(func(__e *ControlFlow) {
															Tm4051 := __e.Get(1)
															_ = Tm4051
															tmp15023 := PrimEqual(Tm4051, Nil)

															if True == tmp15023 {
																tmp15009 := MakeNative(func(__e *ControlFlow) {
																	CurryG := __e.Get(1)
																	_ = CurryG
																	tmp15010 := Call(__e, PrimNS2Value(symshen_4incinfs))

																	_ = tmp15010

																	tmp15011 := MakeNative(func(__e *ControlFlow) {
																		tmp15012 := Call(__e, PrimNS2Value(symshen_4curry), G)

																		tmp15013 := MakeNative(func(__e *ControlFlow) {
																			tmp15014 := MakeNative(func(__e *ControlFlow) {
																				tmp15015 := MakeNative(func(__e *ControlFlow) {
																					tmp15016 := PrimIntern(MakeString(":"))

																					tmp15017 := PrimCons(symverified, Nil)

																					tmp15018 := PrimCons(tmp15016, tmp15017)

																					tmp15019 := PrimCons(CurryG, tmp15018)

																					tmp15020 := PrimCons(tmp15019, V4419)

																					__e.TailApply(PrimNS2Value(symshen_4t_d_1correct), R, V4418, tmp15020, V4420, V4421, K4041, V4423)
																					return

																				}, 0)

																				__e.TailApply(PrimNS2Value(symshen_4cut), V4420, V4421, K4041, tmp15015)
																				return

																			}, 0)

																			__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), CurryG, symboolean, V4419, V4420, V4421, K4041, tmp15014)
																			return

																		}, 0)

																		__e.TailApply(PrimNS2Value(symbind), CurryG, tmp15012, V4420, V4421, K4041, tmp15013)
																		return

																	}, 0)

																	tmp15021 := Call(__e, PrimNS2Value(symshen_4cut), V4420, V4421, K4041, tmp15011)

																	__e.TailApply(PrimNS2Value(symshen_4gc), V4420, tmp15021)
																	return

																}, 1)

																tmp15022 := Call(__e, PrimNS2Value(symshen_4newpv), V4420)

																__e.TailApply(tmp15009, tmp15022)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp15024 := PrimTail(Tm4050)

														tmp15025 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15024, V4420)

														__e.TailApply(tmp15007, tmp15025)
														return

													}, 1)

													tmp15026 := PrimHead(Tm4050)

													__e.TailApply(tmp15006, tmp15026)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp15028 := PrimTail(Tm4049)

											tmp15029 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15028, V4420)

											__e.TailApply(tmp15004, tmp15029)
											return

										}, 1)

										tmp15030 := PrimHead(Tm4049)

										__e.TailApply(tmp15003, tmp15030)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp15032 := PrimTail(Tm4047)

								tmp15033 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15032, V4420)

								__e.TailApply(tmp15001, tmp15033)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp15035 := PrimHead(Tm4047)

						tmp15036 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15035, V4420)

						__e.TailApply(tmp14999, tmp15036)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp15038 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4417, V4420)

				tmp15039 := Call(__e, tmp14997, tmp15038)

				ifres14996 = tmp15039

			} else {
				ifres14996 = False

			}

			__e.TailApply(tmp14985, ifres14996)
			return

		}, 1)

		tmp15041 := PrimNumberAdd(V4422, MakeNumber(1))

		__e.TailApply(tmp14984, tmp15041)
		return

	}, 7)

	tmp15042 := Call(__e, PrimNS2Value(symdef), symshen_4t_d_1correct, tmp14983)

	_ = tmp15042

	tmp15043 := MakeNative(func(__e *ControlFlow) {
		V4424 := __e.Get(1)
		_ = V4424
		V4425 := __e.Get(2)
		_ = V4425
		V4426 := __e.Get(3)
		_ = V4426
		V4427 := __e.Get(4)
		_ = V4427
		V4428 := __e.Get(5)
		_ = V4428
		V4429 := __e.Get(6)
		_ = V4429
		V4430 := __e.Get(7)
		_ = V4430
		V4431 := __e.Get(8)
		_ = V4431
		tmp15044 := MakeNative(func(__e *ControlFlow) {
			K4056 := __e.Get(1)
			_ = K4056
			tmp15045 := MakeNative(func(__e *ControlFlow) {
				C4062 := __e.Get(1)
				_ = C4062
				tmp15124 := PrimEqual(C4062, False)

				if True == tmp15124 {
					tmp15047 := MakeNative(func(__e *ControlFlow) {
						C4066 := __e.Get(1)
						_ = C4066
						tmp15049 := PrimEqual(C4066, False)

						if True == tmp15049 {
							__e.TailApply(PrimNS2Value(symshen_4unlock), V4429, K4056)
							return
						} else {
							__e.Return(C4066)
							return
						}

					}, 1)

					tmp15123 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4429)

					var ifres15050 Obj

					if True == tmp15123 {
						tmp15051 := MakeNative(func(__e *ControlFlow) {
							Tm4067 := __e.Get(1)
							_ = Tm4067
							tmp15120 := PrimIsPair(Tm4067)

							if True == tmp15120 {
								tmp15053 := MakeNative(func(__e *ControlFlow) {
									P := __e.Get(1)
									_ = P
									tmp15054 := MakeNative(func(__e *ControlFlow) {
										Ps := __e.Get(1)
										_ = Ps
										tmp15055 := MakeNative(func(__e *ControlFlow) {
											Tm4068 := __e.Get(1)
											_ = Tm4068
											tmp15116 := PrimIsPair(Tm4068)

											if True == tmp15116 {
												tmp15057 := MakeNative(func(__e *ControlFlow) {
													A := __e.Get(1)
													_ = A
													tmp15058 := MakeNative(func(__e *ControlFlow) {
														Tm4069 := __e.Get(1)
														_ = Tm4069
														tmp15112 := PrimIsPair(Tm4069)

														if True == tmp15112 {
															tmp15060 := MakeNative(func(__e *ControlFlow) {
																Tm4070 := __e.Get(1)
																_ = Tm4070
																tmp15109 := PrimEqual(Tm4070, sym_1_1_6)

																if True == tmp15109 {
																	tmp15062 := MakeNative(func(__e *ControlFlow) {
																		Tm4071 := __e.Get(1)
																		_ = Tm4071
																		tmp15106 := PrimIsPair(Tm4071)

																		if True == tmp15106 {
																			tmp15064 := MakeNative(func(__e *ControlFlow) {
																				B := __e.Get(1)
																				_ = B
																				tmp15065 := MakeNative(func(__e *ControlFlow) {
																					Tm4072 := __e.Get(1)
																					_ = Tm4072
																					tmp15102 := PrimEqual(Tm4072, Nil)

																					if True == tmp15102 {
																						tmp15067 := MakeNative(func(__e *ControlFlow) {
																							Tm4073 := __e.Get(1)
																							_ = Tm4073
																							tmp15068 := MakeNative(func(__e *ControlFlow) {
																								GoTo4074 := __e.Get(1)
																								_ = GoTo4074
																								tmp15086 := PrimIsPair(Tm4073)

																								if True == tmp15086 {
																									tmp15081 := MakeNative(func(__e *ControlFlow) {
																										Hyp := __e.Get(1)
																										_ = Hyp
																										tmp15082 := MakeNative(func(__e *ControlFlow) {
																											Hyps := __e.Get(1)
																											_ = Hyps
																											tmp15083 := Call(__e, GoTo4074, Hyp)

																											__e.TailApply(tmp15083, Hyps)
																											return

																										}, 1)

																										tmp15084 := PrimTail(Tm4073)

																										__e.TailApply(tmp15082, tmp15084)
																										return

																									}, 1)

																									tmp15085 := PrimHead(Tm4073)

																									__e.TailApply(tmp15081, tmp15085)
																									return

																								} else {
																									tmp15080 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm4073)

																									if True == tmp15080 {
																										tmp15071 := MakeNative(func(__e *ControlFlow) {
																											Hyp := __e.Get(1)
																											_ = Hyp
																											tmp15072 := MakeNative(func(__e *ControlFlow) {
																												Hyps := __e.Get(1)
																												_ = Hyps
																												tmp15073 := PrimCons(Hyp, Hyps)

																												tmp15074 := MakeNative(func(__e *ControlFlow) {
																													tmp15075 := Call(__e, GoTo4074, Hyp)

																													__e.TailApply(tmp15075, Hyps)
																													return

																												}, 0)

																												tmp15076 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm4073, tmp15073, V4428, tmp15074)

																												__e.TailApply(PrimNS2Value(symshen_4gc), V4428, tmp15076)
																												return

																											}, 1)

																											tmp15077 := Call(__e, PrimNS2Value(symshen_4newpv), V4428)

																											tmp15078 := Call(__e, tmp15072, tmp15077)

																											__e.TailApply(PrimNS2Value(symshen_4gc), V4428, tmp15078)
																											return

																										}, 1)

																										tmp15079 := Call(__e, PrimNS2Value(symshen_4newpv), V4428)

																										__e.TailApply(tmp15071, tmp15079)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}

																							}, 1)

																							tmp15087 := MakeNative(func(__e *ControlFlow) {
																								Hyp := __e.Get(1)
																								_ = Hyp
																								__e.Return(MakeNative(func(__e *ControlFlow) {
																									Hyps := __e.Get(1)
																									_ = Hyps
																									tmp15088 := MakeNative(func(__e *ControlFlow) {
																										PHyps := __e.Get(1)
																										_ = PHyps
																										tmp15089 := Call(__e, PrimNS2Value(symshen_4incinfs))

																										_ = tmp15089

																										tmp15090 := PrimIntern(MakeString(":"))

																										tmp15091 := PrimCons(A, Nil)

																										tmp15092 := PrimCons(tmp15090, tmp15091)

																										tmp15093 := PrimCons(P, tmp15092)

																										tmp15094 := MakeNative(func(__e *ControlFlow) {
																											tmp15095 := MakeNative(func(__e *ControlFlow) {
																												tmp15096 := MakeNative(func(__e *ControlFlow) {
																													tmp15097 := MakeNative(func(__e *ControlFlow) {
																														tmp15098 := MakeNative(func(__e *ControlFlow) {
																															__e.TailApply(PrimNS2Value(symshen_4t_d_1integrity), Ps, B, Hyps, V4427, V4428, V4429, K4056, V4431)
																															return
																														}, 0)

																														__e.TailApply(PrimNS2Value(symshen_4cut), V4428, V4429, K4056, tmp15098)
																														return

																													}, 0)

																													__e.TailApply(PrimNS2Value(symshen_4system_1S_1h), P, A, PHyps, V4428, V4429, K4056, tmp15097)
																													return

																												}, 0)

																												__e.TailApply(PrimNS2Value(symshen_4cut), V4428, V4429, K4056, tmp15096)
																												return

																											}, 0)

																											__e.TailApply(PrimNS2Value(symshen_4p_1hyps), P, PHyps, V4428, V4429, K4056, tmp15095)
																											return

																										}, 0)

																										tmp15099 := Call(__e, PrimNS2Value(symbind), Hyp, tmp15093, V4428, V4429, K4056, tmp15094)

																										__e.TailApply(PrimNS2Value(symshen_4gc), V4428, tmp15099)
																										return

																									}, 1)

																									tmp15100 := Call(__e, PrimNS2Value(symshen_4newpv), V4428)

																									__e.TailApply(tmp15088, tmp15100)
																									return

																								}, 1))
																								return
																							}, 1)

																							__e.TailApply(tmp15068, tmp15087)
																							return

																						}, 1)

																						tmp15101 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4426, V4428)

																						__e.TailApply(tmp15067, tmp15101)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp15103 := PrimTail(Tm4071)

																				tmp15104 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15103, V4428)

																				__e.TailApply(tmp15065, tmp15104)
																				return

																			}, 1)

																			tmp15105 := PrimHead(Tm4071)

																			__e.TailApply(tmp15064, tmp15105)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp15107 := PrimTail(Tm4069)

																	tmp15108 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15107, V4428)

																	__e.TailApply(tmp15062, tmp15108)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp15110 := PrimHead(Tm4069)

															tmp15111 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15110, V4428)

															__e.TailApply(tmp15060, tmp15111)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp15113 := PrimTail(Tm4068)

													tmp15114 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp15113, V4428)

													__e.TailApply(tmp15058, tmp15114)
													return

												}, 1)

												tmp15115 := PrimHead(Tm4068)

												__e.TailApply(tmp15057, tmp15115)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp15117 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4425, V4428)

										__e.TailApply(tmp15055, tmp15117)
										return

									}, 1)

									tmp15118 := PrimTail(Tm4067)

									__e.TailApply(tmp15054, tmp15118)
									return

								}, 1)

								tmp15119 := PrimHead(Tm4067)

								__e.TailApply(tmp15053, tmp15119)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp15121 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4424, V4428)

						tmp15122 := Call(__e, tmp15051, tmp15121)

						ifres15050 = tmp15122

					} else {
						ifres15050 = False

					}

					__e.TailApply(tmp15047, ifres15050)
					return

				} else {
					__e.Return(C4062)
					return
				}

			}, 1)

			tmp15140 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4429)

			var ifres15125 Obj

			if True == tmp15140 {
				tmp15126 := MakeNative(func(__e *ControlFlow) {
					Tm4063 := __e.Get(1)
					_ = Tm4063
					tmp15137 := PrimEqual(Tm4063, Nil)

					if True == tmp15137 {
						tmp15128 := MakeNative(func(__e *ControlFlow) {
							Tm4064 := __e.Get(1)
							_ = Tm4064
							tmp15129 := MakeNative(func(__e *ControlFlow) {
								GoTo4065 := __e.Get(1)
								_ = GoTo4065
								tmp15133 := PrimEqual(Tm4064, Nil)

								if True == tmp15133 {
									__e.TailApply(PrimNS2Value(symthaw), GoTo4065)
									return
								} else {
									tmp15132 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm4064)

									if True == tmp15132 {
										__e.TailApply(PrimNS2Value(symshen_4bind_b), Tm4064, Nil, V4428, GoTo4065)
										return
									} else {
										__e.Return(False)
										return
									}

								}

							}, 1)

							tmp15134 := MakeNative(func(__e *ControlFlow) {
								tmp15135 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp15135

								__e.TailApply(PrimNS2Value(symis_b), V4425, V4427, V4428, V4429, K4056, V4431)
								return

							}, 0)

							__e.TailApply(tmp15129, tmp15134)
							return

						}, 1)

						tmp15136 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4426, V4428)

						__e.TailApply(tmp15128, tmp15136)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp15138 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4424, V4428)

				tmp15139 := Call(__e, tmp15126, tmp15138)

				ifres15125 = tmp15139

			} else {
				ifres15125 = False

			}

			__e.TailApply(tmp15045, ifres15125)
			return

		}, 1)

		tmp15141 := PrimNumberAdd(V4430, MakeNumber(1))

		__e.TailApply(tmp15044, tmp15141)
		return

	}, 8)

	tmp15142 := Call(__e, PrimNS2Value(symdef), symshen_4t_d_1integrity, tmp15043)

	_ = tmp15142

	tmp15143 := MakeNative(func(__e *ControlFlow) {
		V4432 := __e.Get(1)
		_ = V4432
		V4433 := __e.Get(2)
		_ = V4433
		V4434 := __e.Get(3)
		_ = V4434
		V4435 := __e.Get(4)
		_ = V4435
		V4436 := __e.Get(5)
		_ = V4436
		V4437 := __e.Get(6)
		_ = V4437
		tmp15144 := MakeNative(func(__e *ControlFlow) {
			K4077 := __e.Get(1)
			_ = K4077
			tmp15145 := MakeNative(func(__e *ControlFlow) {
				C4081 := __e.Get(1)
				_ = C4081
				tmp15178 := PrimEqual(C4081, False)

				if True == tmp15178 {
					tmp15147 := MakeNative(func(__e *ControlFlow) {
						C4082 := __e.Get(1)
						_ = C4082
						tmp15156 := PrimEqual(C4082, False)

						if True == tmp15156 {
							tmp15149 := MakeNative(func(__e *ControlFlow) {
								C4084 := __e.Get(1)
								_ = C4084
								tmp15151 := PrimEqual(C4084, False)

								if True == tmp15151 {
									__e.TailApply(PrimNS2Value(symshen_4unlock), V4435, K4077)
									return
								} else {
									__e.Return(C4084)
									return
								}

							}, 1)

							tmp15155 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4435)

							var ifres15152 Obj

							if True == tmp15155 {
								tmp15153 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp15153

								tmp15154 := Call(__e, PrimNS2Value(symbind), V4433, Nil, V4434, V4435, K4077, V4437)

								ifres15152 = tmp15154

							} else {
								ifres15152 = False

							}

							__e.TailApply(tmp15149, ifres15152)
							return

						} else {
							__e.Return(C4082)
							return
						}

					}, 1)

					tmp15177 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4435)

					var ifres15157 Obj

					if True == tmp15177 {
						tmp15158 := MakeNative(func(__e *ControlFlow) {
							Tm4083 := __e.Get(1)
							_ = Tm4083
							tmp15174 := PrimIsPair(Tm4083)

							if True == tmp15174 {
								tmp15160 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									tmp15161 := MakeNative(func(__e *ControlFlow) {
										Y := __e.Get(1)
										_ = Y
										tmp15162 := MakeNative(func(__e *ControlFlow) {
											XHyps := __e.Get(1)
											_ = XHyps
											tmp15163 := MakeNative(func(__e *ControlFlow) {
												YHyps := __e.Get(1)
												_ = YHyps
												tmp15164 := Call(__e, PrimNS2Value(symshen_4incinfs))

												_ = tmp15164

												tmp15165 := MakeNative(func(__e *ControlFlow) {
													tmp15166 := MakeNative(func(__e *ControlFlow) {
														tmp15167 := MakeNative(func(__e *ControlFlow) {
															__e.TailApply(PrimNS2Value(symshen_4join), XHyps, YHyps, V4433, V4434, V4435, K4077, V4437)
															return
														}, 0)

														__e.TailApply(PrimNS2Value(symshen_4p_1hyps), Y, YHyps, V4434, V4435, K4077, tmp15167)
														return

													}, 0)

													__e.TailApply(PrimNS2Value(symshen_4p_1hyps), X, XHyps, V4434, V4435, K4077, tmp15166)
													return

												}, 0)

												tmp15168 := Call(__e, PrimNS2Value(symshen_4cut), V4434, V4435, K4077, tmp15165)

												__e.TailApply(PrimNS2Value(symshen_4gc), V4434, tmp15168)
												return

											}, 1)

											tmp15169 := Call(__e, PrimNS2Value(symshen_4newpv), V4434)

											tmp15170 := Call(__e, tmp15163, tmp15169)

											__e.TailApply(PrimNS2Value(symshen_4gc), V4434, tmp15170)
											return

										}, 1)

										tmp15171 := Call(__e, PrimNS2Value(symshen_4newpv), V4434)

										__e.TailApply(tmp15162, tmp15171)
										return

									}, 1)

									tmp15172 := PrimTail(Tm4083)

									__e.TailApply(tmp15161, tmp15172)
									return

								}, 1)

								tmp15173 := PrimHead(Tm4083)

								__e.TailApply(tmp15160, tmp15173)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp15175 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4432, V4434)

						tmp15176 := Call(__e, tmp15158, tmp15175)

						ifres15157 = tmp15176

					} else {
						ifres15157 = False

					}

					__e.TailApply(tmp15147, ifres15157)
					return

				} else {
					__e.Return(C4081)
					return
				}

			}, 1)

			tmp15194 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4435)

			var ifres15179 Obj

			if True == tmp15194 {
				tmp15180 := MakeNative(func(__e *ControlFlow) {
					A := __e.Get(1)
					_ = A
					tmp15181 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp15181

					tmp15182 := Call(__e, PrimNS2Value(symshen_4deref), V4432, V4434)

					tmp15183 := Call(__e, PrimNS2Value(symshen_4freshterm_2), tmp15182)

					tmp15184 := MakeNative(func(__e *ControlFlow) {
						tmp15185 := MakeNative(func(__e *ControlFlow) {
							tmp15186 := PrimIntern(MakeString(":"))

							tmp15187 := PrimCons(A, Nil)

							tmp15188 := PrimCons(tmp15186, tmp15187)

							tmp15189 := PrimCons(V4432, tmp15188)

							tmp15190 := PrimCons(tmp15189, Nil)

							__e.TailApply(PrimNS2Value(symbind), V4433, tmp15190, V4434, V4435, K4077, V4437)
							return

						}, 0)

						__e.TailApply(PrimNS2Value(symshen_4cut), V4434, V4435, K4077, tmp15185)
						return

					}, 0)

					tmp15191 := Call(__e, PrimNS2Value(symwhen), tmp15183, V4434, V4435, K4077, tmp15184)

					__e.TailApply(PrimNS2Value(symshen_4gc), V4434, tmp15191)
					return

				}, 1)

				tmp15192 := Call(__e, PrimNS2Value(symshen_4newpv), V4434)

				tmp15193 := Call(__e, tmp15180, tmp15192)

				ifres15179 = tmp15193

			} else {
				ifres15179 = False

			}

			__e.TailApply(tmp15145, ifres15179)
			return

		}, 1)

		tmp15195 := PrimNumberAdd(V4436, MakeNumber(1))

		__e.TailApply(tmp15144, tmp15195)
		return

	}, 6)

	tmp15196 := Call(__e, PrimNS2Value(symdef), symshen_4p_1hyps, tmp15143)

	_ = tmp15196

	tmp15197 := MakeNative(func(__e *ControlFlow) {
		V4438 := __e.Get(1)
		_ = V4438
		tmp15206 := PrimIsVector(V4438)

		if True == tmp15206 {
			tmp15204 := PrimIsString(V4438)

			tmp15205 := PrimNot(tmp15204)

			var ifres15200 Obj

			if True == tmp15205 {
				tmp15202 := PrimVectorGet(V4438, MakeNumber(0))

				tmp15203 := PrimEqual(tmp15202, symshen_4print_1freshterm)

				var ifres15201 Obj

				if True == tmp15203 {
					ifres15201 = True

				} else {
					ifres15201 = False

				}

				ifres15200 = ifres15201

			} else {
				ifres15200 = False

			}

			if True == ifres15200 {
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

	tmp15207 := Call(__e, PrimNS2Value(symdef), symshen_4freshterm_2, tmp15197)

	_ = tmp15207

	tmp15208 := MakeNative(func(__e *ControlFlow) {
		V4439 := __e.Get(1)
		_ = V4439
		V4440 := __e.Get(2)
		_ = V4440
		V4441 := __e.Get(3)
		_ = V4441
		V4442 := __e.Get(4)
		_ = V4442
		V4443 := __e.Get(5)
		_ = V4443
		V4444 := __e.Get(6)
		_ = V4444
		V4445 := __e.Get(7)
		_ = V4445
		tmp15209 := MakeNative(func(__e *ControlFlow) {
			C4092 := __e.Get(1)
			_ = C4092
			tmp15245 := PrimEqual(C4092, False)

			if True == tmp15245 {
				tmp15244 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4443)

				if True == tmp15244 {
					tmp15212 := MakeNative(func(__e *ControlFlow) {
						Tm4094 := __e.Get(1)
						_ = Tm4094
						tmp15242 := PrimIsPair(Tm4094)

						if True == tmp15242 {
							tmp15214 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp15215 := MakeNative(func(__e *ControlFlow) {
									Y := __e.Get(1)
									_ = Y
									tmp15216 := MakeNative(func(__e *ControlFlow) {
										Tm4095 := __e.Get(1)
										_ = Tm4095
										tmp15217 := MakeNative(func(__e *ControlFlow) {
											GoTo4096 := __e.Get(1)
											_ = GoTo4096
											tmp15235 := PrimIsPair(Tm4095)

											if True == tmp15235 {
												tmp15230 := MakeNative(func(__e *ControlFlow) {
													X_d := __e.Get(1)
													_ = X_d
													tmp15231 := MakeNative(func(__e *ControlFlow) {
														Z := __e.Get(1)
														_ = Z
														tmp15232 := Call(__e, GoTo4096, X_d)

														__e.TailApply(tmp15232, Z)
														return

													}, 1)

													tmp15233 := PrimTail(Tm4095)

													__e.TailApply(tmp15231, tmp15233)
													return

												}, 1)

												tmp15234 := PrimHead(Tm4095)

												__e.TailApply(tmp15230, tmp15234)
												return

											} else {
												tmp15229 := Call(__e, PrimNS2Value(symshen_4pvar_2), Tm4095)

												if True == tmp15229 {
													tmp15220 := MakeNative(func(__e *ControlFlow) {
														X_d := __e.Get(1)
														_ = X_d
														tmp15221 := MakeNative(func(__e *ControlFlow) {
															Z := __e.Get(1)
															_ = Z
															tmp15222 := PrimCons(X_d, Z)

															tmp15223 := MakeNative(func(__e *ControlFlow) {
																tmp15224 := Call(__e, GoTo4096, X_d)

																__e.TailApply(tmp15224, Z)
																return

															}, 0)

															tmp15225 := Call(__e, PrimNS2Value(symshen_4bind_b), Tm4095, tmp15222, V4442, tmp15223)

															__e.TailApply(PrimNS2Value(symshen_4gc), V4442, tmp15225)
															return

														}, 1)

														tmp15226 := Call(__e, PrimNS2Value(symshen_4newpv), V4442)

														tmp15227 := Call(__e, tmp15221, tmp15226)

														__e.TailApply(PrimNS2Value(symshen_4gc), V4442, tmp15227)
														return

													}, 1)

													tmp15228 := Call(__e, PrimNS2Value(symshen_4newpv), V4442)

													__e.TailApply(tmp15220, tmp15228)
													return

												} else {
													__e.Return(False)
													return
												}

											}

										}, 1)

										tmp15236 := MakeNative(func(__e *ControlFlow) {
											X_d := __e.Get(1)
											_ = X_d
											__e.Return(MakeNative(func(__e *ControlFlow) {
												Z := __e.Get(1)
												_ = Z
												tmp15237 := Call(__e, PrimNS2Value(symshen_4incinfs))

												_ = tmp15237

												tmp15238 := MakeNative(func(__e *ControlFlow) {
													__e.TailApply(PrimNS2Value(symshen_4join), Y, V4440, Z, V4442, V4443, V4444, V4445)
													return
												}, 0)

												__e.TailApply(PrimNS2Value(symbind), X_d, X, V4442, V4443, V4444, tmp15238)
												return

											}, 1))
											return
										}, 1)

										__e.TailApply(tmp15217, tmp15236)
										return

									}, 1)

									tmp15239 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4441, V4442)

									__e.TailApply(tmp15216, tmp15239)
									return

								}, 1)

								tmp15240 := PrimTail(Tm4094)

								__e.TailApply(tmp15215, tmp15240)
								return

							}, 1)

							tmp15241 := PrimHead(Tm4094)

							__e.TailApply(tmp15214, tmp15241)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp15243 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4439, V4442)

					__e.TailApply(tmp15212, tmp15243)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(C4092)
				return
			}

		}, 1)

		tmp15253 := Call(__e, PrimNS2Value(symshen_4unlocked_2), V4443)

		var ifres15246 Obj

		if True == tmp15253 {
			tmp15247 := MakeNative(func(__e *ControlFlow) {
				Tm4093 := __e.Get(1)
				_ = Tm4093
				tmp15250 := PrimEqual(Tm4093, Nil)

				if True == tmp15250 {
					tmp15249 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp15249

					__e.TailApply(PrimNS2Value(symbind), V4441, V4440, V4442, V4443, V4444, V4445)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp15251 := Call(__e, PrimNS2Value(symshen_4lazyderef), V4439, V4442)

			tmp15252 := Call(__e, tmp15247, tmp15251)

			ifres15246 = tmp15252

		} else {
			ifres15246 = False

		}

		__e.TailApply(tmp15209, ifres15246)
		return

	}, 7)

	__e.TailApply(PrimNS2Value(symdef), symshen_4join, tmp15208)
	return

}, 0)
