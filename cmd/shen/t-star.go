package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
	tmp12229 := MakeNative(func(__e *ControlFlow) {
		V4829 := __e.Get(1)
		_ = V4829
		V4830 := __e.Get(2)
		_ = V4830
		tmp12230 := MakeNative(func(__e *ControlFlow) {
			W4831 := __e.Get(1)
			_ = W4831
			tmp12231 := MakeNative(func(__e *ControlFlow) {
				W4832 := __e.Get(1)
				_ = W4832
				tmp12232 := MakeNative(func(__e *ControlFlow) {
					W4833 := __e.Get(1)
					_ = W4833
					tmp12233 := MakeNative(func(__e *ControlFlow) {
						Z4834 := __e.Get(1)
						_ = Z4834
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Z4835 := __e.Get(1)
							_ = Z4835
							__e.Return(MakeNative(func(__e *ControlFlow) {
								Z4836 := __e.Get(1)
								_ = Z4836
								__e.Return(MakeNative(func(__e *ControlFlow) {
									Z4837 := __e.Get(1)
									_ = Z4837
									tmp12234 := MakeNative(func(__e *ControlFlow) {
										W4838 := __e.Get(1)
										_ = W4838
										tmp12235 := Call(__e, PrimFunc(symshen_4incinfs))

										_ = tmp12235

										tmp12236 := Call(__e, PrimFunc(symshen_4deref), W4831, Z4834)

										tmp12237 := Call(__e, PrimFunc(symreceive), tmp12236)

										tmp12238 := Call(__e, PrimFunc(symshen_4deref), W4832, Z4834)

										tmp12239 := Call(__e, PrimFunc(symreceive), tmp12238)

										tmp12240 := MakeNative(func(__e *ControlFlow) {
											tmp12241 := Call(__e, PrimFunc(symshen_4deref), W4833, Z4834)

											tmp12242 := Call(__e, PrimFunc(symreceive), tmp12241)

											tmp12243 := MakeNative(func(__e *ControlFlow) {
												__e.TailApply(PrimFunc(symreturn), W4838, Z4834, Z4835, Z4836, Z4837)
												return
											}, 0)

											__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp12242, W4838, Z4834, Z4835, Z4836, tmp12243)
											return

										}, 0)

										tmp12244 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp12237, tmp12239, W4838, Z4834, Z4835, Z4836, tmp12240)

										__e.TailApply(PrimFunc(symshen_4gc), Z4834, tmp12244)
										return

									}, 1)

									tmp12245 := Call(__e, PrimFunc(symshen_4newpv), Z4834)

									__e.TailApply(tmp12234, tmp12245)
									return

								}, 1))
								return
							}, 1))
							return
						}, 1))
						return
					}, 1)

					tmp12246 := Call(__e, PrimFunc(symshen_4prolog_1vector))

					tmp12247 := Call(__e, tmp12233, tmp12246)

					tmp12248 := Call(__e, PrimFunc(symvector), MakeNumber(0))

					tmp12249 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp12248)

					tmp12250 := Call(__e, PrimFunc(sym_8v), True, tmp12249)

					tmp12251 := Call(__e, tmp12247, tmp12250)

					tmp12252 := Call(__e, tmp12251, MakeNumber(0))

					tmp12253 := MakeNative(func(__e *ControlFlow) {
						__e.Return(True)
						return
					}, 0)

					__e.TailApply(tmp12252, tmp12253)
					return

				}, 1)

				tmp12254 := Call(__e, PrimFunc(symshen_4curry), V4829)

				__e.TailApply(tmp12232, tmp12254)
				return

			}, 1)

			tmp12255 := Call(__e, PrimFunc(symshen_4rectify_1type), V4830)

			__e.TailApply(tmp12231, tmp12255)
			return

		}, 1)

		tmp12256 := Call(__e, PrimFunc(symshen_4extract_1vars), V4830)

		__e.TailApply(tmp12230, tmp12256)
		return

	}, 2)

	tmp12257 := Call(__e, ns2_1set, symshen_4typecheck, tmp12229)

	_ = tmp12257

	tmp12258 := MakeNative(func(__e *ControlFlow) {
		V4839 := __e.Get(1)
		_ = V4839
		V4840 := __e.Get(2)
		_ = V4840
		V4841 := __e.Get(3)
		_ = V4841
		V4842 := __e.Get(4)
		_ = V4842
		V4843 := __e.Get(5)
		_ = V4843
		V4844 := __e.Get(6)
		_ = V4844
		V4845 := __e.Get(7)
		_ = V4845
		tmp12259 := MakeNative(func(__e *ControlFlow) {
			W4846 := __e.Get(1)
			_ = W4846
			tmp12277 := PrimEqual(W4846, False)

			if True == tmp12277 {
				tmp12275 := Call(__e, PrimFunc(symshen_4unlocked_2), V4843)

				if True == tmp12275 {
					tmp12260 := MakeNative(func(__e *ControlFlow) {
						W4848 := __e.Get(1)
						_ = W4848
						tmp12272 := PrimIsPair(W4848)

						if True == tmp12272 {
							tmp12261 := MakeNative(func(__e *ControlFlow) {
								W4849 := __e.Get(1)
								_ = W4849
								tmp12262 := MakeNative(func(__e *ControlFlow) {
									W4850 := __e.Get(1)
									_ = W4850
									tmp12263 := MakeNative(func(__e *ControlFlow) {
										W4851 := __e.Get(1)
										_ = W4851
										tmp12264 := Call(__e, PrimFunc(symshen_4incinfs))

										_ = tmp12264

										tmp12265 := Call(__e, PrimFunc(symshen_4deref), W4851, V4842)

										tmp12266 := Call(__e, PrimFunc(symsubst), tmp12265, W4849, V4840)

										tmp12267 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W4850, tmp12266, V4841, V4842, V4843, V4844, V4845)

										__e.TailApply(PrimFunc(symshen_4gc), V4842, tmp12267)
										return

									}, 1)

									tmp12268 := Call(__e, PrimFunc(symshen_4newpv), V4842)

									__e.TailApply(tmp12263, tmp12268)
									return

								}, 1)

								tmp12269 := PrimTail(W4848)

								__e.TailApply(tmp12262, tmp12269)
								return

							}, 1)

							tmp12270 := PrimHead(W4848)

							__e.TailApply(tmp12261, tmp12270)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp12273 := Call(__e, PrimFunc(symshen_4lazyderef), V4839, V4842)

					__e.TailApply(tmp12260, tmp12273)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W4846)
				return
			}

		}, 1)

		tmp12285 := Call(__e, PrimFunc(symshen_4unlocked_2), V4843)

		var ifres12278 Obj

		if True == tmp12285 {
			tmp12279 := MakeNative(func(__e *ControlFlow) {
				W4847 := __e.Get(1)
				_ = W4847
				tmp12282 := PrimEqual(W4847, Nil)

				if True == tmp12282 {
					tmp12280 := Call(__e, PrimFunc(symshen_4incinfs))

					_ = tmp12280

					__e.TailApply(PrimFunc(symis_b), V4840, V4841, V4842, V4843, V4844, V4845)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp12283 := Call(__e, PrimFunc(symshen_4lazyderef), V4839, V4842)

			tmp12284 := Call(__e, tmp12279, tmp12283)

			ifres12278 = tmp12284

		} else {
			ifres12278 = False

		}

		__e.TailApply(tmp12259, ifres12278)
		return

	}, 7)

	tmp12286 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp12258)

	_ = tmp12286

	tmp12287 := MakeNative(func(__e *ControlFlow) {
		V4852 := __e.Get(1)
		_ = V4852
		V4853 := __e.Get(2)
		_ = V4853
		V4854 := __e.Get(3)
		_ = V4854
		V4855 := __e.Get(4)
		_ = V4855
		V4856 := __e.Get(5)
		_ = V4856
		V4857 := __e.Get(6)
		_ = V4857
		tmp12288 := MakeNative(func(__e *ControlFlow) {
			W4858 := __e.Get(1)
			_ = W4858
			tmp12289 := MakeNative(func(__e *ControlFlow) {
				W4859 := __e.Get(1)
				_ = W4859
				tmp12302 := PrimEqual(W4859, False)

				if True == tmp12302 {
					tmp12290 := MakeNative(func(__e *ControlFlow) {
						W4865 := __e.Get(1)
						_ = W4865
						tmp12292 := PrimEqual(W4865, False)

						if True == tmp12292 {
							__e.TailApply(PrimFunc(symshen_4unlock), V4855, W4858)
							return
						} else {
							__e.Return(W4865)
							return
						}

					}, 1)

					tmp12300 := Call(__e, PrimFunc(symshen_4unlocked_2), V4855)

					var ifres12293 Obj

					if True == tmp12300 {
						tmp12294 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp12294

						tmp12295 := PrimIntern(MakeString(":"))

						tmp12296 := PrimCons(V4853, Nil)

						tmp12297 := PrimCons(tmp12295, tmp12296)

						tmp12298 := PrimCons(V4852, tmp12297)

						tmp12299 := Call(__e, PrimFunc(symshen_4system_1S), tmp12298, Nil, V4854, V4855, W4858, V4857)

						ifres12293 = tmp12299

					} else {
						ifres12293 = False

					}

					__e.TailApply(tmp12290, ifres12293)
					return

				} else {
					__e.Return(W4859)
					return
				}

			}, 1)

			tmp12331 := Call(__e, PrimFunc(symshen_4unlocked_2), V4855)

			var ifres12303 Obj

			if True == tmp12331 {
				tmp12304 := MakeNative(func(__e *ControlFlow) {
					W4860 := __e.Get(1)
					_ = W4860
					tmp12328 := PrimIsPair(W4860)

					if True == tmp12328 {
						tmp12305 := MakeNative(func(__e *ControlFlow) {
							W4861 := __e.Get(1)
							_ = W4861
							tmp12324 := PrimEqual(W4861, symdefine)

							if True == tmp12324 {
								tmp12306 := MakeNative(func(__e *ControlFlow) {
									W4862 := __e.Get(1)
									_ = W4862
									tmp12320 := PrimIsPair(W4862)

									if True == tmp12320 {
										tmp12307 := MakeNative(func(__e *ControlFlow) {
											W4863 := __e.Get(1)
											_ = W4863
											tmp12308 := MakeNative(func(__e *ControlFlow) {
												W4864 := __e.Get(1)
												_ = W4864
												tmp12309 := Call(__e, PrimFunc(symshen_4incinfs))

												_ = tmp12309

												tmp12310 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))

												tmp12311 := MakeNative(func(__e *ControlFlow) {
													tmp12312 := MakeNative(func(__e *ControlFlow) {
														tmp12313 := PrimValue(symshen_4_dspy_d)

														tmp12314 := MakeNative(func(__e *ControlFlow) {
															tmp12315 := PrimCons(W4863, W4864)

															tmp12316 := PrimCons(symdefine, tmp12315)

															__e.TailApply(PrimFunc(symshen_4t_d), tmp12316, V4853, V4854, V4855, W4858, V4857)
															return

														}, 0)

														__e.TailApply(PrimFunc(symshen_4signal_1def), tmp12313, W4863, V4854, V4855, W4858, tmp12314)
														return

													}, 0)

													__e.TailApply(PrimFunc(symshen_4cut), V4854, V4855, W4858, tmp12312)
													return

												}, 0)

												__e.TailApply(PrimFunc(symwhen), tmp12310, V4854, V4855, W4858, tmp12311)
												return

											}, 1)

											tmp12317 := PrimTail(W4862)

											__e.TailApply(tmp12308, tmp12317)
											return

										}, 1)

										tmp12318 := PrimHead(W4862)

										__e.TailApply(tmp12307, tmp12318)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp12321 := PrimTail(W4860)

								tmp12322 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12321, V4854)

								__e.TailApply(tmp12306, tmp12322)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp12325 := PrimHead(W4860)

						tmp12326 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12325, V4854)

						__e.TailApply(tmp12305, tmp12326)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp12329 := Call(__e, PrimFunc(symshen_4lazyderef), V4852, V4854)

				tmp12330 := Call(__e, tmp12304, tmp12329)

				ifres12303 = tmp12330

			} else {
				ifres12303 = False

			}

			__e.TailApply(tmp12289, ifres12303)
			return

		}, 1)

		tmp12332 := PrimNumberAdd(V4856, MakeNumber(1))

		__e.TailApply(tmp12288, tmp12332)
		return

	}, 6)

	tmp12333 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp12287)

	_ = tmp12333

	tmp12334 := MakeNative(func(__e *ControlFlow) {
		V4866 := __e.Get(1)
		_ = V4866
		V4867 := __e.Get(2)
		_ = V4867
		V4868 := __e.Get(3)
		_ = V4868
		V4869 := __e.Get(4)
		_ = V4869
		V4870 := __e.Get(5)
		_ = V4870
		V4871 := __e.Get(6)
		_ = V4871
		tmp12335 := MakeNative(func(__e *ControlFlow) {
			W4872 := __e.Get(1)
			_ = W4872
			tmp12352 := PrimEqual(W4872, False)

			if True == tmp12352 {
				tmp12350 := Call(__e, PrimFunc(symshen_4unlocked_2), V4869)

				if True == tmp12350 {
					tmp12336 := MakeNative(func(__e *ControlFlow) {
						W4874 := __e.Get(1)
						_ = W4874
						tmp12347 := PrimEqual(W4874, True)

						if True == tmp12347 {
							tmp12337 := MakeNative(func(__e *ControlFlow) {
								W4875 := __e.Get(1)
								_ = W4875
								tmp12338 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp12338

								tmp12339 := Call(__e, PrimFunc(symshen_4deref), V4867, V4868)

								tmp12340 := Call(__e, PrimFunc(symshen_4app), tmp12339, MakeString(")\n"), symshen_4a)

								tmp12341 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp12340)

								tmp12342 := Call(__e, PrimFunc(symstoutput))

								tmp12343 := Call(__e, PrimFunc(sympr), tmp12341, tmp12342)

								tmp12344 := Call(__e, PrimFunc(symis), W4875, tmp12343, V4868, V4869, V4870, V4871)

								__e.TailApply(PrimFunc(symshen_4gc), V4868, tmp12344)
								return

							}, 1)

							tmp12345 := Call(__e, PrimFunc(symshen_4newpv), V4868)

							__e.TailApply(tmp12337, tmp12345)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp12348 := Call(__e, PrimFunc(symshen_4lazyderef), V4866, V4868)

					__e.TailApply(tmp12336, tmp12348)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W4872)
				return
			}

		}, 1)

		tmp12360 := Call(__e, PrimFunc(symshen_4unlocked_2), V4869)

		var ifres12353 Obj

		if True == tmp12360 {
			tmp12354 := MakeNative(func(__e *ControlFlow) {
				W4873 := __e.Get(1)
				_ = W4873
				tmp12357 := PrimEqual(W4873, False)

				if True == tmp12357 {
					tmp12355 := Call(__e, PrimFunc(symshen_4incinfs))

					_ = tmp12355

					__e.TailApply(PrimFunc(symthaw), V4871)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp12358 := Call(__e, PrimFunc(symshen_4lazyderef), V4866, V4868)

			tmp12359 := Call(__e, tmp12354, tmp12358)

			ifres12353 = tmp12359

		} else {
			ifres12353 = False

		}

		__e.TailApply(tmp12335, ifres12353)
		return

	}, 6)

	tmp12361 := Call(__e, ns2_1set, symshen_4signal_1def, tmp12334)

	_ = tmp12361

	tmp12362 := MakeNative(func(__e *ControlFlow) {
		V4876 := __e.Get(1)
		_ = V4876
		tmp12363 := Call(__e, PrimFunc(symshen_4curry_1type), V4876)

		__e.TailApply(PrimFunc(symshen_4demodulate), tmp12363)
		return

	}, 1)

	tmp12364 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp12362)

	_ = tmp12364

	tmp12365 := MakeNative(func(__e *ControlFlow) {
		V4877 := __e.Get(1)
		_ = V4877
		tmp12366 := MakeNative(func(__e *ControlFlow) {
			tmp12367 := MakeNative(func(__e *ControlFlow) {
				W4878 := __e.Get(1)
				_ = W4878
				tmp12369 := PrimEqual(W4878, V4877)

				if True == tmp12369 {
					__e.Return(V4877)
					return
				} else {
					__e.TailApply(PrimFunc(symshen_4demodulate), W4878)
					return
				}

			}, 1)

			tmp12370 := MakeNative(func(__e *ControlFlow) {
				Z4879 := __e.Get(1)
				_ = Z4879
				__e.TailApply(PrimFunc(symshen_4demod), Z4879)
				return
			}, 1)

			tmp12371 := Call(__e, PrimFunc(symshen_4walk), tmp12370, V4877)

			__e.TailApply(tmp12367, tmp12371)
			return

		}, 0)

		tmp12372 := MakeNative(func(__e *ControlFlow) {
			Z4880 := __e.Get(1)
			_ = Z4880
			__e.Return(V4877)
			return
		}, 1)

		__e.TailApply(try_1catch, tmp12366, tmp12372)
		return

	}, 1)

	tmp12373 := Call(__e, ns2_1set, symshen_4demodulate, tmp12365)

	_ = tmp12373

	tmp12374 := MakeNative(func(__e *ControlFlow) {
		V4881 := __e.Get(1)
		_ = V4881
		tmp12498 := PrimIsPair(V4881)

		var ifres12471 Obj

		if True == tmp12498 {
			tmp12496 := PrimTail(V4881)

			tmp12497 := PrimIsPair(tmp12496)

			var ifres12473 Obj

			if True == tmp12497 {
				tmp12493 := PrimTail(V4881)

				tmp12494 := PrimHead(tmp12493)

				tmp12495 := PrimEqual(sym_1_1_6, tmp12494)

				var ifres12475 Obj

				if True == tmp12495 {
					tmp12490 := PrimTail(V4881)

					tmp12491 := PrimTail(tmp12490)

					tmp12492 := PrimIsPair(tmp12491)

					var ifres12477 Obj

					if True == tmp12492 {
						tmp12486 := PrimTail(V4881)

						tmp12487 := PrimTail(tmp12486)

						tmp12488 := PrimTail(tmp12487)

						tmp12489 := PrimIsPair(tmp12488)

						var ifres12479 Obj

						if True == tmp12489 {
							tmp12481 := PrimTail(V4881)

							tmp12482 := PrimTail(tmp12481)

							tmp12483 := PrimTail(tmp12482)

							tmp12484 := PrimHead(tmp12483)

							tmp12485 := PrimEqual(sym_1_1_6, tmp12484)

							var ifres12480 Obj

							if True == tmp12485 {
								ifres12480 = True

							} else {
								ifres12480 = False

							}

							ifres12479 = ifres12480

						} else {
							ifres12479 = False

						}

						var ifres12478 Obj

						if True == ifres12479 {
							ifres12478 = True

						} else {
							ifres12478 = False

						}

						ifres12477 = ifres12478

					} else {
						ifres12477 = False

					}

					var ifres12476 Obj

					if True == ifres12477 {
						ifres12476 = True

					} else {
						ifres12476 = False

					}

					ifres12475 = ifres12476

				} else {
					ifres12475 = False

				}

				var ifres12474 Obj

				if True == ifres12475 {
					ifres12474 = True

				} else {
					ifres12474 = False

				}

				ifres12473 = ifres12474

			} else {
				ifres12473 = False

			}

			var ifres12472 Obj

			if True == ifres12473 {
				ifres12472 = True

			} else {
				ifres12472 = False

			}

			ifres12471 = ifres12472

		} else {
			ifres12471 = False

		}

		if True == ifres12471 {
			tmp12375 := PrimHead(V4881)

			tmp12376 := PrimTail(V4881)

			tmp12377 := PrimTail(tmp12376)

			tmp12378 := PrimCons(tmp12377, Nil)

			tmp12379 := PrimCons(sym_1_1_6, tmp12378)

			tmp12380 := PrimCons(tmp12375, tmp12379)

			__e.TailApply(PrimFunc(symshen_4curry_1type), tmp12380)
			return

		} else {
			tmp12469 := PrimIsPair(V4881)

			var ifres12429 Obj

			if True == tmp12469 {
				tmp12467 := PrimHead(V4881)

				tmp12468 := PrimIsPair(tmp12467)

				var ifres12431 Obj

				if True == tmp12468 {
					tmp12464 := PrimHead(V4881)

					tmp12465 := PrimHead(tmp12464)

					tmp12466 := PrimEqual(symlist, tmp12465)

					var ifres12433 Obj

					if True == tmp12466 {
						tmp12461 := PrimHead(V4881)

						tmp12462 := PrimTail(tmp12461)

						tmp12463 := PrimIsPair(tmp12462)

						var ifres12435 Obj

						if True == tmp12463 {
							tmp12457 := PrimHead(V4881)

							tmp12458 := PrimTail(tmp12457)

							tmp12459 := PrimTail(tmp12458)

							tmp12460 := PrimEqual(Nil, tmp12459)

							var ifres12437 Obj

							if True == tmp12460 {
								tmp12455 := PrimTail(V4881)

								tmp12456 := PrimIsPair(tmp12455)

								var ifres12439 Obj

								if True == tmp12456 {
									tmp12452 := PrimTail(V4881)

									tmp12453 := PrimHead(tmp12452)

									tmp12454 := PrimEqual(sym_a_a_6, tmp12453)

									var ifres12441 Obj

									if True == tmp12454 {
										tmp12449 := PrimTail(V4881)

										tmp12450 := PrimTail(tmp12449)

										tmp12451 := PrimIsPair(tmp12450)

										var ifres12443 Obj

										if True == tmp12451 {
											tmp12445 := PrimTail(V4881)

											tmp12446 := PrimTail(tmp12445)

											tmp12447 := PrimTail(tmp12446)

											tmp12448 := PrimEqual(Nil, tmp12447)

											var ifres12444 Obj

											if True == tmp12448 {
												ifres12444 = True

											} else {
												ifres12444 = False

											}

											ifres12443 = ifres12444

										} else {
											ifres12443 = False

										}

										var ifres12442 Obj

										if True == ifres12443 {
											ifres12442 = True

										} else {
											ifres12442 = False

										}

										ifres12441 = ifres12442

									} else {
										ifres12441 = False

									}

									var ifres12440 Obj

									if True == ifres12441 {
										ifres12440 = True

									} else {
										ifres12440 = False

									}

									ifres12439 = ifres12440

								} else {
									ifres12439 = False

								}

								var ifres12438 Obj

								if True == ifres12439 {
									ifres12438 = True

								} else {
									ifres12438 = False

								}

								ifres12437 = ifres12438

							} else {
								ifres12437 = False

							}

							var ifres12436 Obj

							if True == ifres12437 {
								ifres12436 = True

							} else {
								ifres12436 = False

							}

							ifres12435 = ifres12436

						} else {
							ifres12435 = False

						}

						var ifres12434 Obj

						if True == ifres12435 {
							ifres12434 = True

						} else {
							ifres12434 = False

						}

						ifres12433 = ifres12434

					} else {
						ifres12433 = False

					}

					var ifres12432 Obj

					if True == ifres12433 {
						ifres12432 = True

					} else {
						ifres12432 = False

					}

					ifres12431 = ifres12432

				} else {
					ifres12431 = False

				}

				var ifres12430 Obj

				if True == ifres12431 {
					ifres12430 = True

				} else {
					ifres12430 = False

				}

				ifres12429 = ifres12430

			} else {
				ifres12429 = False

			}

			if True == ifres12429 {
				tmp12381 := PrimHead(V4881)

				tmp12382 := PrimHead(V4881)

				tmp12383 := PrimTail(V4881)

				tmp12384 := PrimTail(tmp12383)

				tmp12385 := PrimCons(tmp12382, tmp12384)

				tmp12386 := PrimCons(symstr, tmp12385)

				tmp12387 := PrimCons(tmp12386, Nil)

				tmp12388 := PrimCons(sym_1_1_6, tmp12387)

				tmp12389 := PrimCons(tmp12381, tmp12388)

				__e.TailApply(PrimFunc(symshen_4curry_1type), tmp12389)
				return

			} else {
				tmp12427 := PrimIsPair(V4881)

				var ifres12400 Obj

				if True == tmp12427 {
					tmp12425 := PrimTail(V4881)

					tmp12426 := PrimIsPair(tmp12425)

					var ifres12402 Obj

					if True == tmp12426 {
						tmp12422 := PrimTail(V4881)

						tmp12423 := PrimHead(tmp12422)

						tmp12424 := PrimEqual(sym_d, tmp12423)

						var ifres12404 Obj

						if True == tmp12424 {
							tmp12419 := PrimTail(V4881)

							tmp12420 := PrimTail(tmp12419)

							tmp12421 := PrimIsPair(tmp12420)

							var ifres12406 Obj

							if True == tmp12421 {
								tmp12415 := PrimTail(V4881)

								tmp12416 := PrimTail(tmp12415)

								tmp12417 := PrimTail(tmp12416)

								tmp12418 := PrimIsPair(tmp12417)

								var ifres12408 Obj

								if True == tmp12418 {
									tmp12410 := PrimTail(V4881)

									tmp12411 := PrimTail(tmp12410)

									tmp12412 := PrimTail(tmp12411)

									tmp12413 := PrimHead(tmp12412)

									tmp12414 := PrimEqual(sym_d, tmp12413)

									var ifres12409 Obj

									if True == tmp12414 {
										ifres12409 = True

									} else {
										ifres12409 = False

									}

									ifres12408 = ifres12409

								} else {
									ifres12408 = False

								}

								var ifres12407 Obj

								if True == ifres12408 {
									ifres12407 = True

								} else {
									ifres12407 = False

								}

								ifres12406 = ifres12407

							} else {
								ifres12406 = False

							}

							var ifres12405 Obj

							if True == ifres12406 {
								ifres12405 = True

							} else {
								ifres12405 = False

							}

							ifres12404 = ifres12405

						} else {
							ifres12404 = False

						}

						var ifres12403 Obj

						if True == ifres12404 {
							ifres12403 = True

						} else {
							ifres12403 = False

						}

						ifres12402 = ifres12403

					} else {
						ifres12402 = False

					}

					var ifres12401 Obj

					if True == ifres12402 {
						ifres12401 = True

					} else {
						ifres12401 = False

					}

					ifres12400 = ifres12401

				} else {
					ifres12400 = False

				}

				if True == ifres12400 {
					tmp12390 := PrimHead(V4881)

					tmp12391 := PrimTail(V4881)

					tmp12392 := PrimTail(tmp12391)

					tmp12393 := PrimCons(tmp12392, Nil)

					tmp12394 := PrimCons(sym_d, tmp12393)

					tmp12395 := PrimCons(tmp12390, tmp12394)

					__e.TailApply(PrimFunc(symshen_4curry_1type), tmp12395)
					return

				} else {
					tmp12398 := PrimIsPair(V4881)

					if True == tmp12398 {
						tmp12396 := MakeNative(func(__e *ControlFlow) {
							Z4882 := __e.Get(1)
							_ = Z4882
							__e.TailApply(PrimFunc(symshen_4curry_1type), Z4882)
							return
						}, 1)

						__e.TailApply(PrimFunc(symmap), tmp12396, V4881)
						return

					} else {
						__e.Return(V4881)
						return
					}

				}

			}

		}

	}, 1)

	tmp12499 := Call(__e, ns2_1set, symshen_4curry_1type, tmp12374)

	_ = tmp12499

	tmp12500 := MakeNative(func(__e *ControlFlow) {
		V4883 := __e.Get(1)
		_ = V4883
		tmp12618 := PrimIsPair(V4883)

		var ifres12610 Obj

		if True == tmp12618 {
			tmp12616 := PrimHead(V4883)

			tmp12617 := PrimEqual(symdefine, tmp12616)

			var ifres12612 Obj

			if True == tmp12617 {
				tmp12614 := PrimTail(V4883)

				tmp12615 := PrimIsPair(tmp12614)

				var ifres12613 Obj

				if True == tmp12615 {
					ifres12613 = True

				} else {
					ifres12613 = False

				}

				ifres12612 = ifres12613

			} else {
				ifres12612 = False

			}

			var ifres12611 Obj

			if True == ifres12612 {
				ifres12611 = True

			} else {
				ifres12611 = False

			}

			ifres12610 = ifres12611

		} else {
			ifres12610 = False

		}

		if True == ifres12610 {
			__e.Return(V4883)
			return
		} else {
			tmp12608 := PrimIsPair(V4883)

			var ifres12589 Obj

			if True == tmp12608 {
				tmp12606 := PrimHead(V4883)

				tmp12607 := PrimEqual(symtype, tmp12606)

				var ifres12591 Obj

				if True == tmp12607 {
					tmp12604 := PrimTail(V4883)

					tmp12605 := PrimIsPair(tmp12604)

					var ifres12593 Obj

					if True == tmp12605 {
						tmp12601 := PrimTail(V4883)

						tmp12602 := PrimTail(tmp12601)

						tmp12603 := PrimIsPair(tmp12602)

						var ifres12595 Obj

						if True == tmp12603 {
							tmp12597 := PrimTail(V4883)

							tmp12598 := PrimTail(tmp12597)

							tmp12599 := PrimTail(tmp12598)

							tmp12600 := PrimEqual(Nil, tmp12599)

							var ifres12596 Obj

							if True == tmp12600 {
								ifres12596 = True

							} else {
								ifres12596 = False

							}

							ifres12595 = ifres12596

						} else {
							ifres12595 = False

						}

						var ifres12594 Obj

						if True == ifres12595 {
							ifres12594 = True

						} else {
							ifres12594 = False

						}

						ifres12593 = ifres12594

					} else {
						ifres12593 = False

					}

					var ifres12592 Obj

					if True == ifres12593 {
						ifres12592 = True

					} else {
						ifres12592 = False

					}

					ifres12591 = ifres12592

				} else {
					ifres12591 = False

				}

				var ifres12590 Obj

				if True == ifres12591 {
					ifres12590 = True

				} else {
					ifres12590 = False

				}

				ifres12589 = ifres12590

			} else {
				ifres12589 = False

			}

			if True == ifres12589 {
				tmp12501 := PrimTail(V4883)

				tmp12502 := PrimHead(tmp12501)

				tmp12503 := Call(__e, PrimFunc(symshen_4curry), tmp12502)

				tmp12504 := PrimTail(V4883)

				tmp12505 := PrimTail(tmp12504)

				tmp12506 := PrimCons(tmp12503, tmp12505)

				__e.Return(PrimCons(symtype, tmp12506))
				return

			} else {
				tmp12587 := PrimIsPair(V4883)

				var ifres12568 Obj

				if True == tmp12587 {
					tmp12585 := PrimHead(V4883)

					tmp12586 := PrimEqual(syminput_7, tmp12585)

					var ifres12570 Obj

					if True == tmp12586 {
						tmp12583 := PrimTail(V4883)

						tmp12584 := PrimIsPair(tmp12583)

						var ifres12572 Obj

						if True == tmp12584 {
							tmp12580 := PrimTail(V4883)

							tmp12581 := PrimTail(tmp12580)

							tmp12582 := PrimIsPair(tmp12581)

							var ifres12574 Obj

							if True == tmp12582 {
								tmp12576 := PrimTail(V4883)

								tmp12577 := PrimTail(tmp12576)

								tmp12578 := PrimTail(tmp12577)

								tmp12579 := PrimEqual(Nil, tmp12578)

								var ifres12575 Obj

								if True == tmp12579 {
									ifres12575 = True

								} else {
									ifres12575 = False

								}

								ifres12574 = ifres12575

							} else {
								ifres12574 = False

							}

							var ifres12573 Obj

							if True == ifres12574 {
								ifres12573 = True

							} else {
								ifres12573 = False

							}

							ifres12572 = ifres12573

						} else {
							ifres12572 = False

						}

						var ifres12571 Obj

						if True == ifres12572 {
							ifres12571 = True

						} else {
							ifres12571 = False

						}

						ifres12570 = ifres12571

					} else {
						ifres12570 = False

					}

					var ifres12569 Obj

					if True == ifres12570 {
						ifres12569 = True

					} else {
						ifres12569 = False

					}

					ifres12568 = ifres12569

				} else {
					ifres12568 = False

				}

				if True == ifres12568 {
					tmp12507 := PrimTail(V4883)

					tmp12508 := PrimHead(tmp12507)

					tmp12509 := PrimTail(V4883)

					tmp12510 := PrimTail(tmp12509)

					tmp12511 := PrimHead(tmp12510)

					tmp12512 := Call(__e, PrimFunc(symshen_4curry), tmp12511)

					tmp12513 := PrimCons(tmp12512, Nil)

					tmp12514 := PrimCons(tmp12508, tmp12513)

					__e.Return(PrimCons(syminput_7, tmp12514))
					return

				} else {
					tmp12566 := PrimIsPair(V4883)

					var ifres12562 Obj

					if True == tmp12566 {
						tmp12564 := PrimHead(V4883)

						tmp12565 := Call(__e, PrimFunc(symshen_4special_2), tmp12564)

						var ifres12563 Obj

						if True == tmp12565 {
							ifres12563 = True

						} else {
							ifres12563 = False

						}

						ifres12562 = ifres12563

					} else {
						ifres12562 = False

					}

					if True == ifres12562 {
						tmp12515 := PrimHead(V4883)

						tmp12516 := MakeNative(func(__e *ControlFlow) {
							Z4884 := __e.Get(1)
							_ = Z4884
							__e.TailApply(PrimFunc(symshen_4curry), Z4884)
							return
						}, 1)

						tmp12517 := PrimTail(V4883)

						tmp12518 := Call(__e, PrimFunc(symmap), tmp12516, tmp12517)

						__e.Return(PrimCons(tmp12515, tmp12518))
						return

					} else {
						tmp12560 := PrimIsPair(V4883)

						var ifres12556 Obj

						if True == tmp12560 {
							tmp12558 := PrimHead(V4883)

							tmp12559 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp12558)

							var ifres12557 Obj

							if True == tmp12559 {
								ifres12557 = True

							} else {
								ifres12557 = False

							}

							ifres12556 = ifres12557

						} else {
							ifres12556 = False

						}

						if True == ifres12556 {
							__e.Return(V4883)
							return
						} else {
							tmp12554 := PrimIsPair(V4883)

							var ifres12545 Obj

							if True == tmp12554 {
								tmp12552 := PrimTail(V4883)

								tmp12553 := PrimIsPair(tmp12552)

								var ifres12547 Obj

								if True == tmp12553 {
									tmp12549 := PrimTail(V4883)

									tmp12550 := PrimTail(tmp12549)

									tmp12551 := PrimIsPair(tmp12550)

									var ifres12548 Obj

									if True == tmp12551 {
										ifres12548 = True

									} else {
										ifres12548 = False

									}

									ifres12547 = ifres12548

								} else {
									ifres12547 = False

								}

								var ifres12546 Obj

								if True == ifres12547 {
									ifres12546 = True

								} else {
									ifres12546 = False

								}

								ifres12545 = ifres12546

							} else {
								ifres12545 = False

							}

							if True == ifres12545 {
								tmp12519 := PrimHead(V4883)

								tmp12520 := PrimTail(V4883)

								tmp12521 := PrimHead(tmp12520)

								tmp12522 := PrimCons(tmp12521, Nil)

								tmp12523 := PrimCons(tmp12519, tmp12522)

								tmp12524 := PrimTail(V4883)

								tmp12525 := PrimTail(tmp12524)

								tmp12526 := PrimCons(tmp12523, tmp12525)

								__e.TailApply(PrimFunc(symshen_4curry), tmp12526)
								return

							} else {
								tmp12543 := PrimIsPair(V4883)

								var ifres12534 Obj

								if True == tmp12543 {
									tmp12541 := PrimTail(V4883)

									tmp12542 := PrimIsPair(tmp12541)

									var ifres12536 Obj

									if True == tmp12542 {
										tmp12538 := PrimTail(V4883)

										tmp12539 := PrimTail(tmp12538)

										tmp12540 := PrimEqual(Nil, tmp12539)

										var ifres12537 Obj

										if True == tmp12540 {
											ifres12537 = True

										} else {
											ifres12537 = False

										}

										ifres12536 = ifres12537

									} else {
										ifres12536 = False

									}

									var ifres12535 Obj

									if True == ifres12536 {
										ifres12535 = True

									} else {
										ifres12535 = False

									}

									ifres12534 = ifres12535

								} else {
									ifres12534 = False

								}

								if True == ifres12534 {
									tmp12527 := PrimHead(V4883)

									tmp12528 := Call(__e, PrimFunc(symshen_4curry), tmp12527)

									tmp12529 := PrimTail(V4883)

									tmp12530 := PrimHead(tmp12529)

									tmp12531 := Call(__e, PrimFunc(symshen_4curry), tmp12530)

									tmp12532 := PrimCons(tmp12531, Nil)

									__e.Return(PrimCons(tmp12528, tmp12532))
									return

								} else {
									__e.Return(V4883)
									return
								}

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp12619 := Call(__e, ns2_1set, symshen_4curry, tmp12500)

	_ = tmp12619

	tmp12620 := MakeNative(func(__e *ControlFlow) {
		V4885 := __e.Get(1)
		_ = V4885
		tmp12621 := PrimValue(symshen_4_dspecial_d)

		__e.TailApply(PrimFunc(symelement_2), V4885, tmp12621)
		return

	}, 1)

	tmp12622 := Call(__e, ns2_1set, symshen_4special_2, tmp12620)

	_ = tmp12622

	tmp12623 := MakeNative(func(__e *ControlFlow) {
		V4886 := __e.Get(1)
		_ = V4886
		tmp12624 := PrimValue(symshen_4_dextraspecial_d)

		__e.TailApply(PrimFunc(symelement_2), V4886, tmp12624)
		return

	}, 1)

	tmp12625 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp12623)

	_ = tmp12625

	tmp12626 := MakeNative(func(__e *ControlFlow) {
		V4887 := __e.Get(1)
		_ = V4887
		V4888 := __e.Get(2)
		_ = V4888
		V4889 := __e.Get(3)
		_ = V4889
		V4890 := __e.Get(4)
		_ = V4890
		V4891 := __e.Get(5)
		_ = V4891
		V4892 := __e.Get(6)
		_ = V4892
		tmp12627 := MakeNative(func(__e *ControlFlow) {
			W4893 := __e.Get(1)
			_ = W4893
			tmp12628 := MakeNative(func(__e *ControlFlow) {
				W4894 := __e.Get(1)
				_ = W4894
				tmp12686 := PrimEqual(W4894, False)

				if True == tmp12686 {
					tmp12629 := MakeNative(func(__e *ControlFlow) {
						W4895 := __e.Get(1)
						_ = W4895
						tmp12648 := PrimEqual(W4895, False)

						if True == tmp12648 {
							tmp12630 := MakeNative(func(__e *ControlFlow) {
								W4903 := __e.Get(1)
								_ = W4903
								tmp12640 := PrimEqual(W4903, False)

								if True == tmp12640 {
									tmp12631 := MakeNative(func(__e *ControlFlow) {
										W4904 := __e.Get(1)
										_ = W4904
										tmp12633 := PrimEqual(W4904, False)

										if True == tmp12633 {
											__e.TailApply(PrimFunc(symshen_4unlock), V4890, W4893)
											return
										} else {
											__e.Return(W4904)
											return
										}

									}, 1)

									tmp12638 := Call(__e, PrimFunc(symshen_4unlocked_2), V4890)

									var ifres12634 Obj

									if True == tmp12638 {
										tmp12635 := Call(__e, PrimFunc(symshen_4incinfs))

										_ = tmp12635

										tmp12636 := PrimValue(symshen_4_ddatatypes_d)

										tmp12637 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4887, V4888, tmp12636, V4889, V4890, W4893, V4892)

										ifres12634 = tmp12637

									} else {
										ifres12634 = False

									}

									__e.TailApply(tmp12631, ifres12634)
									return

								} else {
									__e.Return(W4903)
									return
								}

							}, 1)

							tmp12646 := Call(__e, PrimFunc(symshen_4unlocked_2), V4890)

							var ifres12641 Obj

							if True == tmp12646 {
								tmp12642 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp12642

								tmp12643 := PrimValue(symshen_4_dspy_d)

								tmp12644 := MakeNative(func(__e *ControlFlow) {
									__e.TailApply(PrimFunc(symshen_4show), V4887, V4888, V4889, V4890, W4893, V4892)
									return
								}, 0)

								tmp12645 := Call(__e, PrimFunc(symwhen), tmp12643, V4889, V4890, W4893, tmp12644)

								ifres12641 = tmp12645

							} else {
								ifres12641 = False

							}

							__e.TailApply(tmp12630, ifres12641)
							return

						} else {
							__e.Return(W4895)
							return
						}

					}, 1)

					tmp12684 := Call(__e, PrimFunc(symshen_4unlocked_2), V4890)

					var ifres12649 Obj

					if True == tmp12684 {
						tmp12650 := MakeNative(func(__e *ControlFlow) {
							W4896 := __e.Get(1)
							_ = W4896
							tmp12681 := PrimIsPair(W4896)

							if True == tmp12681 {
								tmp12651 := MakeNative(func(__e *ControlFlow) {
									W4897 := __e.Get(1)
									_ = W4897
									tmp12652 := MakeNative(func(__e *ControlFlow) {
										W4898 := __e.Get(1)
										_ = W4898
										tmp12676 := PrimIsPair(W4898)

										if True == tmp12676 {
											tmp12653 := MakeNative(func(__e *ControlFlow) {
												W4899 := __e.Get(1)
												_ = W4899
												tmp12654 := MakeNative(func(__e *ControlFlow) {
													W4900 := __e.Get(1)
													_ = W4900
													tmp12671 := PrimIsPair(W4900)

													if True == tmp12671 {
														tmp12655 := MakeNative(func(__e *ControlFlow) {
															W4901 := __e.Get(1)
															_ = W4901
															tmp12656 := MakeNative(func(__e *ControlFlow) {
																W4902 := __e.Get(1)
																_ = W4902
																tmp12666 := PrimEqual(W4902, Nil)

																if True == tmp12666 {
																	tmp12657 := Call(__e, PrimFunc(symshen_4incinfs))

																	_ = tmp12657

																	tmp12658 := Call(__e, PrimFunc(symshen_4deref), W4899, V4889)

																	tmp12659 := PrimIntern(MakeString(":"))

																	tmp12660 := PrimEqual(tmp12658, tmp12659)

																	tmp12661 := MakeNative(func(__e *ControlFlow) {
																		tmp12662 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))

																		tmp12663 := MakeNative(func(__e *ControlFlow) {
																			tmp12664 := MakeNative(func(__e *ControlFlow) {
																				__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4897, W4901, V4888, V4889, V4890, W4893, V4892)
																				return
																			}, 0)

																			__e.TailApply(PrimFunc(symshen_4cut), V4889, V4890, W4893, tmp12664)
																			return

																		}, 0)

																		__e.TailApply(PrimFunc(symwhen), tmp12662, V4889, V4890, W4893, tmp12663)
																		return

																	}, 0)

																	__e.TailApply(PrimFunc(symwhen), tmp12660, V4889, V4890, W4893, tmp12661)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp12667 := PrimTail(W4900)

															tmp12668 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12667, V4889)

															__e.TailApply(tmp12656, tmp12668)
															return

														}, 1)

														tmp12669 := PrimHead(W4900)

														__e.TailApply(tmp12655, tmp12669)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp12672 := PrimTail(W4898)

												tmp12673 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12672, V4889)

												__e.TailApply(tmp12654, tmp12673)
												return

											}, 1)

											tmp12674 := PrimHead(W4898)

											__e.TailApply(tmp12653, tmp12674)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp12677 := PrimTail(W4896)

									tmp12678 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12677, V4889)

									__e.TailApply(tmp12652, tmp12678)
									return

								}, 1)

								tmp12679 := PrimHead(W4896)

								__e.TailApply(tmp12651, tmp12679)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp12682 := Call(__e, PrimFunc(symshen_4lazyderef), V4887, V4889)

						tmp12683 := Call(__e, tmp12650, tmp12682)

						ifres12649 = tmp12683

					} else {
						ifres12649 = False

					}

					__e.TailApply(tmp12629, ifres12649)
					return

				} else {
					__e.Return(W4894)
					return
				}

			}, 1)

			tmp12691 := Call(__e, PrimFunc(symshen_4unlocked_2), V4890)

			var ifres12687 Obj

			if True == tmp12691 {
				tmp12688 := Call(__e, PrimFunc(symshen_4incinfs))

				_ = tmp12688

				tmp12689 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))

				tmp12690 := Call(__e, PrimFunc(symwhen), tmp12689, V4889, V4890, W4893, V4892)

				ifres12687 = tmp12690

			} else {
				ifres12687 = False

			}

			__e.TailApply(tmp12628, ifres12687)
			return

		}, 1)

		tmp12692 := PrimNumberAdd(V4891, MakeNumber(1))

		__e.TailApply(tmp12627, tmp12692)
		return

	}, 6)

	tmp12693 := Call(__e, ns2_1set, symshen_4system_1S, tmp12626)

	_ = tmp12693

	tmp12694 := MakeNative(func(__e *ControlFlow) {
		V4911 := __e.Get(1)
		_ = V4911
		V4912 := __e.Get(2)
		_ = V4912
		V4913 := __e.Get(3)
		_ = V4913
		V4914 := __e.Get(4)
		_ = V4914
		V4915 := __e.Get(5)
		_ = V4915
		V4916 := __e.Get(6)
		_ = V4916
		tmp12695 := Call(__e, PrimFunc(symshen_4line))

		_ = tmp12695

		tmp12696 := Call(__e, PrimFunc(symshen_4deref), V4911, V4913)

		tmp12697 := Call(__e, PrimFunc(symshen_4show_1p), tmp12696)

		_ = tmp12697

		tmp12698 := Call(__e, PrimFunc(symnl), MakeNumber(2))

		_ = tmp12698

		tmp12699 := Call(__e, PrimFunc(symshen_4deref), V4912, V4913)

		tmp12700 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp12699, MakeNumber(1))

		_ = tmp12700

		tmp12701 := Call(__e, PrimFunc(symshen_4pause_1for_1user))

		_ = tmp12701

		__e.Return(False)
		return

	}, 6)

	tmp12702 := Call(__e, ns2_1set, symshen_4show, tmp12694)

	_ = tmp12702

	tmp12703 := MakeNative(func(__e *ControlFlow) {
		tmp12704 := MakeNative(func(__e *ControlFlow) {
			W4917 := __e.Get(1)
			_ = W4917
			tmp12706 := PrimEqual(MakeNumber(1), W4917)

			var ifres12705 Obj

			if True == tmp12706 {
				ifres12705 = MakeString("")

			} else {
				ifres12705 = MakeString("s")

			}

			tmp12707 := Call(__e, PrimFunc(symshen_4app), ifres12705, MakeString(" \n?- "), symshen_4a)

			tmp12708 := PrimStringConcat(MakeString(" inference"), tmp12707)

			tmp12709 := Call(__e, PrimFunc(symshen_4app), W4917, tmp12708, symshen_4a)

			tmp12710 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp12709)

			tmp12711 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp12710, tmp12711)
			return

		}, 1)

		tmp12712 := Call(__e, PrimFunc(syminferences))

		__e.TailApply(tmp12704, tmp12712)
		return

	}, 0)

	tmp12713 := Call(__e, ns2_1set, symshen_4line, tmp12703)

	_ = tmp12713

	tmp12714 := MakeNative(func(__e *ControlFlow) {
		V4918 := __e.Get(1)
		_ = V4918
		tmp12746 := PrimIsPair(V4918)

		var ifres12725 Obj

		if True == tmp12746 {
			tmp12744 := PrimTail(V4918)

			tmp12745 := PrimIsPair(tmp12744)

			var ifres12727 Obj

			if True == tmp12745 {
				tmp12741 := PrimTail(V4918)

				tmp12742 := PrimTail(tmp12741)

				tmp12743 := PrimIsPair(tmp12742)

				var ifres12729 Obj

				if True == tmp12743 {
					tmp12737 := PrimTail(V4918)

					tmp12738 := PrimTail(tmp12737)

					tmp12739 := PrimTail(tmp12738)

					tmp12740 := PrimEqual(Nil, tmp12739)

					var ifres12731 Obj

					if True == tmp12740 {
						tmp12733 := PrimTail(V4918)

						tmp12734 := PrimHead(tmp12733)

						tmp12735 := PrimIntern(MakeString(":"))

						tmp12736 := PrimEqual(tmp12734, tmp12735)

						var ifres12732 Obj

						if True == tmp12736 {
							ifres12732 = True

						} else {
							ifres12732 = False

						}

						ifres12731 = ifres12732

					} else {
						ifres12731 = False

					}

					var ifres12730 Obj

					if True == ifres12731 {
						ifres12730 = True

					} else {
						ifres12730 = False

					}

					ifres12729 = ifres12730

				} else {
					ifres12729 = False

				}

				var ifres12728 Obj

				if True == ifres12729 {
					ifres12728 = True

				} else {
					ifres12728 = False

				}

				ifres12727 = ifres12728

			} else {
				ifres12727 = False

			}

			var ifres12726 Obj

			if True == ifres12727 {
				ifres12726 = True

			} else {
				ifres12726 = False

			}

			ifres12725 = ifres12726

		} else {
			ifres12725 = False

		}

		if True == ifres12725 {
			tmp12715 := PrimHead(V4918)

			tmp12716 := Call(__e, PrimFunc(symshen_4prterm), tmp12715)

			_ = tmp12716

			tmp12717 := Call(__e, PrimFunc(symstoutput))

			tmp12718 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp12717)

			_ = tmp12718

			tmp12719 := PrimTail(V4918)

			tmp12720 := PrimTail(tmp12719)

			tmp12721 := PrimHead(tmp12720)

			tmp12722 := Call(__e, PrimFunc(symshen_4app), tmp12721, MakeString(""), symshen_4r)

			tmp12723 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp12722, tmp12723)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4prterm), V4918)
			return
		}

	}, 1)

	tmp12747 := Call(__e, ns2_1set, symshen_4show_1p, tmp12714)

	_ = tmp12747

	tmp12748 := MakeNative(func(__e *ControlFlow) {
		V4919 := __e.Get(1)
		_ = V4919
		tmp12791 := PrimIsPair(V4919)

		var ifres12772 Obj

		if True == tmp12791 {
			tmp12789 := PrimHead(V4919)

			tmp12790 := PrimEqual(symcons, tmp12789)

			var ifres12774 Obj

			if True == tmp12790 {
				tmp12787 := PrimTail(V4919)

				tmp12788 := PrimIsPair(tmp12787)

				var ifres12776 Obj

				if True == tmp12788 {
					tmp12784 := PrimTail(V4919)

					tmp12785 := PrimTail(tmp12784)

					tmp12786 := PrimIsPair(tmp12785)

					var ifres12778 Obj

					if True == tmp12786 {
						tmp12780 := PrimTail(V4919)

						tmp12781 := PrimTail(tmp12780)

						tmp12782 := PrimTail(tmp12781)

						tmp12783 := PrimEqual(Nil, tmp12782)

						var ifres12779 Obj

						if True == tmp12783 {
							ifres12779 = True

						} else {
							ifres12779 = False

						}

						ifres12778 = ifres12779

					} else {
						ifres12778 = False

					}

					var ifres12777 Obj

					if True == ifres12778 {
						ifres12777 = True

					} else {
						ifres12777 = False

					}

					ifres12776 = ifres12777

				} else {
					ifres12776 = False

				}

				var ifres12775 Obj

				if True == ifres12776 {
					ifres12775 = True

				} else {
					ifres12775 = False

				}

				ifres12774 = ifres12775

			} else {
				ifres12774 = False

			}

			var ifres12773 Obj

			if True == ifres12774 {
				ifres12773 = True

			} else {
				ifres12773 = False

			}

			ifres12772 = ifres12773

		} else {
			ifres12772 = False

		}

		if True == ifres12772 {
			tmp12749 := Call(__e, PrimFunc(symstoutput))

			tmp12750 := Call(__e, PrimFunc(sympr), MakeString("["), tmp12749)

			_ = tmp12750

			tmp12751 := PrimTail(V4919)

			tmp12752 := PrimHead(tmp12751)

			tmp12753 := Call(__e, PrimFunc(symshen_4prterm), tmp12752)

			_ = tmp12753

			tmp12754 := PrimTail(V4919)

			tmp12755 := PrimTail(tmp12754)

			tmp12756 := PrimHead(tmp12755)

			tmp12757 := Call(__e, PrimFunc(symshen_4prtl), tmp12756)

			_ = tmp12757

			tmp12758 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp12758)
			return

		} else {
			tmp12770 := PrimIsPair(V4919)

			if True == tmp12770 {
				tmp12759 := Call(__e, PrimFunc(symstoutput))

				tmp12760 := Call(__e, PrimFunc(sympr), MakeString("("), tmp12759)

				_ = tmp12760

				tmp12761 := PrimHead(V4919)

				tmp12762 := Call(__e, PrimFunc(symshen_4prterm), tmp12761)

				_ = tmp12762

				tmp12763 := MakeNative(func(__e *ControlFlow) {
					Z4920 := __e.Get(1)
					_ = Z4920
					tmp12764 := Call(__e, PrimFunc(symstoutput))

					tmp12765 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp12764)

					_ = tmp12765

					__e.TailApply(PrimFunc(symshen_4prterm), Z4920)
					return

				}, 1)

				tmp12766 := PrimTail(V4919)

				tmp12767 := Call(__e, PrimFunc(symmap), tmp12763, tmp12766)

				_ = tmp12767

				tmp12768 := Call(__e, PrimFunc(symstoutput))

				__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp12768)
				return

			} else {
				__e.TailApply(PrimFunc(symprint), V4919)
				return
			}

		}

	}, 1)

	tmp12792 := Call(__e, ns2_1set, symshen_4prterm, tmp12748)

	_ = tmp12792

	tmp12793 := MakeNative(func(__e *ControlFlow) {
		V4921 := __e.Get(1)
		_ = V4921
		tmp12826 := PrimEqual(Nil, V4921)

		if True == tmp12826 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp12824 := PrimIsPair(V4921)

			var ifres12805 Obj

			if True == tmp12824 {
				tmp12822 := PrimHead(V4921)

				tmp12823 := PrimEqual(symcons, tmp12822)

				var ifres12807 Obj

				if True == tmp12823 {
					tmp12820 := PrimTail(V4921)

					tmp12821 := PrimIsPair(tmp12820)

					var ifres12809 Obj

					if True == tmp12821 {
						tmp12817 := PrimTail(V4921)

						tmp12818 := PrimTail(tmp12817)

						tmp12819 := PrimIsPair(tmp12818)

						var ifres12811 Obj

						if True == tmp12819 {
							tmp12813 := PrimTail(V4921)

							tmp12814 := PrimTail(tmp12813)

							tmp12815 := PrimTail(tmp12814)

							tmp12816 := PrimEqual(Nil, tmp12815)

							var ifres12812 Obj

							if True == tmp12816 {
								ifres12812 = True

							} else {
								ifres12812 = False

							}

							ifres12811 = ifres12812

						} else {
							ifres12811 = False

						}

						var ifres12810 Obj

						if True == ifres12811 {
							ifres12810 = True

						} else {
							ifres12810 = False

						}

						ifres12809 = ifres12810

					} else {
						ifres12809 = False

					}

					var ifres12808 Obj

					if True == ifres12809 {
						ifres12808 = True

					} else {
						ifres12808 = False

					}

					ifres12807 = ifres12808

				} else {
					ifres12807 = False

				}

				var ifres12806 Obj

				if True == ifres12807 {
					ifres12806 = True

				} else {
					ifres12806 = False

				}

				ifres12805 = ifres12806

			} else {
				ifres12805 = False

			}

			if True == ifres12805 {
				tmp12794 := Call(__e, PrimFunc(symstoutput))

				tmp12795 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp12794)

				_ = tmp12795

				tmp12796 := PrimTail(V4921)

				tmp12797 := PrimHead(tmp12796)

				tmp12798 := Call(__e, PrimFunc(symshen_4prterm), tmp12797)

				_ = tmp12798

				tmp12799 := PrimTail(V4921)

				tmp12800 := PrimTail(tmp12799)

				tmp12801 := PrimHead(tmp12800)

				__e.TailApply(PrimFunc(symshen_4prtl), tmp12801)
				return

			} else {
				tmp12802 := Call(__e, PrimFunc(symstoutput))

				tmp12803 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp12802)

				_ = tmp12803

				__e.TailApply(PrimFunc(symshen_4prterm), V4921)
				return

			}

		}

	}, 1)

	tmp12827 := Call(__e, ns2_1set, symshen_4prtl, tmp12793)

	_ = tmp12827

	tmp12828 := MakeNative(func(__e *ControlFlow) {
		V4928 := __e.Get(1)
		_ = V4928
		V4929 := __e.Get(2)
		_ = V4929
		tmp12841 := PrimEqual(Nil, V4928)

		if True == tmp12841 {
			tmp12829 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp12829)
			return

		} else {
			tmp12839 := PrimIsPair(V4928)

			if True == tmp12839 {
				tmp12830 := Call(__e, PrimFunc(symshen_4app), V4929, MakeString(". "), symshen_4a)

				tmp12831 := Call(__e, PrimFunc(symstoutput))

				tmp12832 := Call(__e, PrimFunc(sympr), tmp12830, tmp12831)

				_ = tmp12832

				tmp12833 := PrimHead(V4928)

				tmp12834 := Call(__e, PrimFunc(symshen_4show_1p), tmp12833)

				_ = tmp12834

				tmp12835 := Call(__e, PrimFunc(symnl), MakeNumber(1))

				_ = tmp12835

				tmp12836 := PrimTail(V4928)

				tmp12837 := PrimNumberAdd(V4929, MakeNumber(1))

				__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp12836, tmp12837)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
				return
			}

		}

	}, 2)

	tmp12842 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp12828)

	_ = tmp12842

	tmp12843 := MakeNative(func(__e *ControlFlow) {
		tmp12844 := MakeNative(func(__e *ControlFlow) {
			W4930 := __e.Get(1)
			_ = W4930
			tmp12846 := PrimEqual(W4930, MakeNumber(94))

			if True == tmp12846 {
				__e.Return(PrimSimpleError(MakeString("input aborted\n")))
				return
			} else {
				__e.TailApply(PrimFunc(symnl), MakeNumber(1))
				return
			}

		}, 1)

		tmp12847 := Call(__e, PrimFunc(symstinput))

		tmp12848 := PrimReadByte(tmp12847)

		__e.TailApply(tmp12844, tmp12848)
		return

	}, 0)

	tmp12849 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp12843)

	_ = tmp12849

	tmp12850 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
		return
	}, 0)

	tmp12851 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp12850)

	_ = tmp12851

	tmp12852 := MakeNative(func(__e *ControlFlow) {
		tmp12854 := Call(__e, PrimFunc(syminferences))

		tmp12855 := PrimValue(symshen_4_dmaxinferences_d)

		tmp12856 := PrimGreatThan(tmp12854, tmp12855)

		if True == tmp12856 {
			__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
			return
		} else {
			__e.Return(False)
			return
		}

	}, 0)

	tmp12857 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp12852)

	_ = tmp12857

	tmp12858 := MakeNative(func(__e *ControlFlow) {
		V4931 := __e.Get(1)
		_ = V4931
		V4932 := __e.Get(2)
		_ = V4932
		V4933 := __e.Get(3)
		_ = V4933
		V4934 := __e.Get(4)
		_ = V4934
		V4935 := __e.Get(5)
		_ = V4935
		V4936 := __e.Get(6)
		_ = V4936
		V4937 := __e.Get(7)
		_ = V4937
		tmp12859 := MakeNative(func(__e *ControlFlow) {
			W4938 := __e.Get(1)
			_ = W4938
			tmp12860 := MakeNative(func(__e *ControlFlow) {
				W4939 := __e.Get(1)
				_ = W4939
				tmp13774 := PrimEqual(W4939, False)

				if True == tmp13774 {
					tmp12861 := MakeNative(func(__e *ControlFlow) {
						W4940 := __e.Get(1)
						_ = W4940
						tmp13764 := PrimEqual(W4940, False)

						if True == tmp13764 {
							tmp12862 := MakeNative(func(__e *ControlFlow) {
								W4941 := __e.Get(1)
								_ = W4941
								tmp13758 := PrimEqual(W4941, False)

								if True == tmp13758 {
									tmp12863 := MakeNative(func(__e *ControlFlow) {
										W4942 := __e.Get(1)
										_ = W4942
										tmp13739 := PrimEqual(W4942, False)

										if True == tmp13739 {
											tmp12864 := MakeNative(func(__e *ControlFlow) {
												W4946 := __e.Get(1)
												_ = W4946
												tmp13706 := PrimEqual(W4946, False)

												if True == tmp13706 {
													tmp12865 := MakeNative(func(__e *ControlFlow) {
														W4952 := __e.Get(1)
														_ = W4952
														tmp13679 := PrimEqual(W4952, False)

														if True == tmp13679 {
															tmp12866 := MakeNative(func(__e *ControlFlow) {
																W4958 := __e.Get(1)
																_ = W4958
																tmp13644 := PrimEqual(W4958, False)

																if True == tmp13644 {
																	tmp12867 := MakeNative(func(__e *ControlFlow) {
																		W4965 := __e.Get(1)
																		_ = W4965
																		tmp13613 := PrimEqual(W4965, False)

																		if True == tmp13613 {
																			tmp12868 := MakeNative(func(__e *ControlFlow) {
																				W4972 := __e.Get(1)
																				_ = W4972
																				tmp13528 := PrimEqual(W4972, False)

																				if True == tmp13528 {
																					tmp12869 := MakeNative(func(__e *ControlFlow) {
																						W4993 := __e.Get(1)
																						_ = W4993
																						tmp13422 := PrimEqual(W4993, False)

																						if True == tmp13422 {
																							tmp12870 := MakeNative(func(__e *ControlFlow) {
																								W5021 := __e.Get(1)
																								_ = W5021
																								tmp13337 := PrimEqual(W5021, False)

																								if True == tmp13337 {
																									tmp12871 := MakeNative(func(__e *ControlFlow) {
																										W5042 := __e.Get(1)
																										_ = W5042
																										tmp13294 := PrimEqual(W5042, False)

																										if True == tmp13294 {
																											tmp12872 := MakeNative(func(__e *ControlFlow) {
																												W5052 := __e.Get(1)
																												_ = W5052
																												tmp13170 := PrimEqual(W5052, False)

																												if True == tmp13170 {
																													tmp12873 := MakeNative(func(__e *ControlFlow) {
																														W5082 := __e.Get(1)
																														_ = W5082
																														tmp13106 := PrimEqual(W5082, False)

																														if True == tmp13106 {
																															tmp12874 := MakeNative(func(__e *ControlFlow) {
																																W5095 := __e.Get(1)
																																_ = W5095
																																tmp13018 := PrimEqual(W5095, False)

																																if True == tmp13018 {
																																	tmp12875 := MakeNative(func(__e *ControlFlow) {
																																		W5116 := __e.Get(1)
																																		_ = W5116
																																		tmp12980 := PrimEqual(W5116, False)

																																		if True == tmp12980 {
																																			tmp12876 := MakeNative(func(__e *ControlFlow) {
																																				W5124 := __e.Get(1)
																																				_ = W5124
																																				tmp12941 := PrimEqual(W5124, False)

																																				if True == tmp12941 {
																																					tmp12877 := MakeNative(func(__e *ControlFlow) {
																																						W5132 := __e.Get(1)
																																						_ = W5132
																																						tmp12903 := PrimEqual(W5132, False)

																																						if True == tmp12903 {
																																							tmp12878 := MakeNative(func(__e *ControlFlow) {
																																								W5140 := __e.Get(1)
																																								_ = W5140
																																								tmp12892 := PrimEqual(W5140, False)

																																								if True == tmp12892 {
																																									tmp12879 := MakeNative(func(__e *ControlFlow) {
																																										W5142 := __e.Get(1)
																																										_ = W5142
																																										tmp12881 := PrimEqual(W5142, False)

																																										if True == tmp12881 {
																																											__e.TailApply(PrimFunc(symshen_4unlock), V4935, W4938)
																																											return
																																										} else {
																																											__e.Return(W5142)
																																											return
																																										}

																																									}, 1)

																																									tmp12890 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																																									var ifres12882 Obj

																																									if True == tmp12890 {
																																										tmp12883 := Call(__e, PrimFunc(symshen_4incinfs))

																																										_ = tmp12883

																																										tmp12884 := PrimIntern(MakeString(":"))

																																										tmp12885 := PrimCons(V4932, Nil)

																																										tmp12886 := PrimCons(tmp12884, tmp12885)

																																										tmp12887 := PrimCons(V4931, tmp12886)

																																										tmp12888 := PrimValue(symshen_4_ddatatypes_d)

																																										tmp12889 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp12887, V4933, tmp12888, V4934, V4935, W4938, V4937)

																																										ifres12882 = tmp12889

																																									} else {
																																										ifres12882 = False

																																									}

																																									__e.TailApply(tmp12879, ifres12882)
																																									return

																																								} else {
																																									__e.Return(W5140)
																																									return
																																								}

																																							}, 1)

																																							tmp12901 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																																							var ifres12893 Obj

																																							if True == tmp12901 {
																																								tmp12894 := MakeNative(func(__e *ControlFlow) {
																																									W5141 := __e.Get(1)
																																									_ = W5141
																																									tmp12895 := Call(__e, PrimFunc(symshen_4incinfs))

																																									_ = tmp12895

																																									tmp12896 := MakeNative(func(__e *ControlFlow) {
																																										tmp12897 := MakeNative(func(__e *ControlFlow) {
																																											__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4931, V4932, W5141, V4934, V4935, W4938, V4937)
																																											return
																																										}, 0)

																																										__e.TailApply(PrimFunc(symshen_4cut), V4934, V4935, W4938, tmp12897)
																																										return

																																									}, 0)

																																									tmp12898 := Call(__e, PrimFunc(symshen_4l_1rules), V4933, W5141, False, V4934, V4935, W4938, tmp12896)

																																									__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp12898)
																																									return

																																								}, 1)

																																								tmp12899 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																								tmp12900 := Call(__e, tmp12894, tmp12899)

																																								ifres12893 = tmp12900

																																							} else {
																																								ifres12893 = False

																																							}

																																							__e.TailApply(tmp12878, ifres12893)
																																							return

																																						} else {
																																							__e.Return(W5132)
																																							return
																																						}

																																					}, 1)

																																					tmp12939 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																																					var ifres12904 Obj

																																					if True == tmp12939 {
																																						tmp12905 := MakeNative(func(__e *ControlFlow) {
																																							W5133 := __e.Get(1)
																																							_ = W5133
																																							tmp12936 := PrimIsPair(W5133)

																																							if True == tmp12936 {
																																								tmp12906 := MakeNative(func(__e *ControlFlow) {
																																									W5134 := __e.Get(1)
																																									_ = W5134
																																									tmp12932 := PrimEqual(W5134, symset)

																																									if True == tmp12932 {
																																										tmp12907 := MakeNative(func(__e *ControlFlow) {
																																											W5135 := __e.Get(1)
																																											_ = W5135
																																											tmp12928 := PrimIsPair(W5135)

																																											if True == tmp12928 {
																																												tmp12908 := MakeNative(func(__e *ControlFlow) {
																																													W5136 := __e.Get(1)
																																													_ = W5136
																																													tmp12909 := MakeNative(func(__e *ControlFlow) {
																																														W5137 := __e.Get(1)
																																														_ = W5137
																																														tmp12923 := PrimIsPair(W5137)

																																														if True == tmp12923 {
																																															tmp12910 := MakeNative(func(__e *ControlFlow) {
																																																W5138 := __e.Get(1)
																																																_ = W5138
																																																tmp12911 := MakeNative(func(__e *ControlFlow) {
																																																	W5139 := __e.Get(1)
																																																	_ = W5139
																																																	tmp12918 := PrimEqual(W5139, Nil)

																																																	if True == tmp12918 {
																																																		tmp12912 := Call(__e, PrimFunc(symshen_4incinfs))

																																																		_ = tmp12912

																																																		tmp12913 := MakeNative(func(__e *ControlFlow) {
																																																			tmp12914 := PrimCons(W5136, Nil)

																																																			tmp12915 := PrimCons(symvalue, tmp12914)

																																																			tmp12916 := MakeNative(func(__e *ControlFlow) {
																																																				__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5138, V4932, V4933, V4934, V4935, W4938, V4937)
																																																				return
																																																			}, 0)

																																																			__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp12915, V4932, V4933, V4934, V4935, W4938, tmp12916)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5136, symsymbol, V4933, V4934, V4935, W4938, tmp12913)
																																																		return

																																																	} else {
																																																		__e.Return(False)
																																																		return
																																																	}

																																																}, 1)

																																																tmp12919 := PrimTail(W5137)

																																																tmp12920 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12919, V4934)

																																																__e.TailApply(tmp12911, tmp12920)
																																																return

																																															}, 1)

																																															tmp12921 := PrimHead(W5137)

																																															__e.TailApply(tmp12910, tmp12921)
																																															return

																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}, 1)

																																													tmp12924 := PrimTail(W5135)

																																													tmp12925 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12924, V4934)

																																													__e.TailApply(tmp12909, tmp12925)
																																													return

																																												}, 1)

																																												tmp12926 := PrimHead(W5135)

																																												__e.TailApply(tmp12908, tmp12926)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}, 1)

																																										tmp12929 := PrimTail(W5133)

																																										tmp12930 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12929, V4934)

																																										__e.TailApply(tmp12907, tmp12930)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp12933 := PrimHead(W5133)

																																								tmp12934 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12933, V4934)

																																								__e.TailApply(tmp12906, tmp12934)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp12937 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																																						tmp12938 := Call(__e, tmp12905, tmp12937)

																																						ifres12904 = tmp12938

																																					} else {
																																						ifres12904 = False

																																					}

																																					__e.TailApply(tmp12877, ifres12904)
																																					return

																																				} else {
																																					__e.Return(W5124)
																																					return
																																				}

																																			}, 1)

																																			tmp12978 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																																			var ifres12942 Obj

																																			if True == tmp12978 {
																																				tmp12943 := MakeNative(func(__e *ControlFlow) {
																																					W5125 := __e.Get(1)
																																					_ = W5125
																																					tmp12975 := PrimIsPair(W5125)

																																					if True == tmp12975 {
																																						tmp12944 := MakeNative(func(__e *ControlFlow) {
																																							W5126 := __e.Get(1)
																																							_ = W5126
																																							tmp12971 := PrimEqual(W5126, syminput_7)

																																							if True == tmp12971 {
																																								tmp12945 := MakeNative(func(__e *ControlFlow) {
																																									W5127 := __e.Get(1)
																																									_ = W5127
																																									tmp12967 := PrimIsPair(W5127)

																																									if True == tmp12967 {
																																										tmp12946 := MakeNative(func(__e *ControlFlow) {
																																											W5128 := __e.Get(1)
																																											_ = W5128
																																											tmp12947 := MakeNative(func(__e *ControlFlow) {
																																												W5129 := __e.Get(1)
																																												_ = W5129
																																												tmp12962 := PrimIsPair(W5129)

																																												if True == tmp12962 {
																																													tmp12948 := MakeNative(func(__e *ControlFlow) {
																																														W5130 := __e.Get(1)
																																														_ = W5130
																																														tmp12949 := MakeNative(func(__e *ControlFlow) {
																																															W5131 := __e.Get(1)
																																															_ = W5131
																																															tmp12957 := PrimEqual(W5131, Nil)

																																															if True == tmp12957 {
																																																tmp12950 := Call(__e, PrimFunc(symshen_4incinfs))

																																																_ = tmp12950

																																																tmp12951 := Call(__e, PrimFunc(symshen_4deref), W5128, V4934)

																																																tmp12952 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp12951)

																																																tmp12953 := MakeNative(func(__e *ControlFlow) {
																																																	tmp12954 := PrimCons(symin, Nil)

																																																	tmp12955 := PrimCons(symstream, tmp12954)

																																																	__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5130, tmp12955, V4933, V4934, V4935, W4938, V4937)
																																																	return

																																																}, 0)

																																																__e.TailApply(PrimFunc(symis_b), V4932, tmp12952, V4934, V4935, W4938, tmp12953)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}, 1)

																																														tmp12958 := PrimTail(W5129)

																																														tmp12959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12958, V4934)

																																														__e.TailApply(tmp12949, tmp12959)
																																														return

																																													}, 1)

																																													tmp12960 := PrimHead(W5129)

																																													__e.TailApply(tmp12948, tmp12960)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp12963 := PrimTail(W5127)

																																											tmp12964 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12963, V4934)

																																											__e.TailApply(tmp12947, tmp12964)
																																											return

																																										}, 1)

																																										tmp12965 := PrimHead(W5127)

																																										__e.TailApply(tmp12946, tmp12965)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp12968 := PrimTail(W5125)

																																								tmp12969 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12968, V4934)

																																								__e.TailApply(tmp12945, tmp12969)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp12972 := PrimHead(W5125)

																																						tmp12973 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12972, V4934)

																																						__e.TailApply(tmp12944, tmp12973)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp12976 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																																				tmp12977 := Call(__e, tmp12943, tmp12976)

																																				ifres12942 = tmp12977

																																			} else {
																																				ifres12942 = False

																																			}

																																			__e.TailApply(tmp12876, ifres12942)
																																			return

																																		} else {
																																			__e.Return(W5116)
																																			return
																																		}

																																	}, 1)

																																	tmp13016 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																																	var ifres12981 Obj

																																	if True == tmp13016 {
																																		tmp12982 := MakeNative(func(__e *ControlFlow) {
																																			W5117 := __e.Get(1)
																																			_ = W5117
																																			tmp13013 := PrimIsPair(W5117)

																																			if True == tmp13013 {
																																				tmp12983 := MakeNative(func(__e *ControlFlow) {
																																					W5118 := __e.Get(1)
																																					_ = W5118
																																					tmp13009 := PrimEqual(W5118, symtype)

																																					if True == tmp13009 {
																																						tmp12984 := MakeNative(func(__e *ControlFlow) {
																																							W5119 := __e.Get(1)
																																							_ = W5119
																																							tmp13005 := PrimIsPair(W5119)

																																							if True == tmp13005 {
																																								tmp12985 := MakeNative(func(__e *ControlFlow) {
																																									W5120 := __e.Get(1)
																																									_ = W5120
																																									tmp12986 := MakeNative(func(__e *ControlFlow) {
																																										W5121 := __e.Get(1)
																																										_ = W5121
																																										tmp13000 := PrimIsPair(W5121)

																																										if True == tmp13000 {
																																											tmp12987 := MakeNative(func(__e *ControlFlow) {
																																												W5122 := __e.Get(1)
																																												_ = W5122
																																												tmp12988 := MakeNative(func(__e *ControlFlow) {
																																													W5123 := __e.Get(1)
																																													_ = W5123
																																													tmp12995 := PrimEqual(W5123, Nil)

																																													if True == tmp12995 {
																																														tmp12989 := Call(__e, PrimFunc(symshen_4incinfs))

																																														_ = tmp12989

																																														tmp12990 := MakeNative(func(__e *ControlFlow) {
																																															tmp12991 := Call(__e, PrimFunc(symshen_4deref), W5122, V4934)

																																															tmp12992 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp12991)

																																															tmp12993 := MakeNative(func(__e *ControlFlow) {
																																																__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5120, V4932, V4933, V4934, V4935, W4938, V4937)
																																																return
																																															}, 0)

																																															__e.TailApply(PrimFunc(symis_b), tmp12992, V4932, V4934, V4935, W4938, tmp12993)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimFunc(symshen_4cut), V4934, V4935, W4938, tmp12990)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}, 1)

																																												tmp12996 := PrimTail(W5121)

																																												tmp12997 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12996, V4934)

																																												__e.TailApply(tmp12988, tmp12997)
																																												return

																																											}, 1)

																																											tmp12998 := PrimHead(W5121)

																																											__e.TailApply(tmp12987, tmp12998)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp13001 := PrimTail(W5119)

																																									tmp13002 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13001, V4934)

																																									__e.TailApply(tmp12986, tmp13002)
																																									return

																																								}, 1)

																																								tmp13003 := PrimHead(W5119)

																																								__e.TailApply(tmp12985, tmp13003)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp13006 := PrimTail(W5117)

																																						tmp13007 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13006, V4934)

																																						__e.TailApply(tmp12984, tmp13007)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp13010 := PrimHead(W5117)

																																				tmp13011 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13010, V4934)

																																				__e.TailApply(tmp12983, tmp13011)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp13014 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																																		tmp13015 := Call(__e, tmp12982, tmp13014)

																																		ifres12981 = tmp13015

																																	} else {
																																		ifres12981 = False

																																	}

																																	__e.TailApply(tmp12875, ifres12981)
																																	return

																																} else {
																																	__e.Return(W5095)
																																	return
																																}

																															}, 1)

																															tmp13104 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																															var ifres13019 Obj

																															if True == tmp13104 {
																																tmp13020 := MakeNative(func(__e *ControlFlow) {
																																	W5096 := __e.Get(1)
																																	_ = W5096
																																	tmp13101 := PrimIsPair(W5096)

																																	if True == tmp13101 {
																																		tmp13021 := MakeNative(func(__e *ControlFlow) {
																																			W5097 := __e.Get(1)
																																			_ = W5097
																																			tmp13097 := PrimEqual(W5097, symopen)

																																			if True == tmp13097 {
																																				tmp13022 := MakeNative(func(__e *ControlFlow) {
																																					W5098 := __e.Get(1)
																																					_ = W5098
																																					tmp13093 := PrimIsPair(W5098)

																																					if True == tmp13093 {
																																						tmp13023 := MakeNative(func(__e *ControlFlow) {
																																							W5099 := __e.Get(1)
																																							_ = W5099
																																							tmp13024 := MakeNative(func(__e *ControlFlow) {
																																								W5100 := __e.Get(1)
																																								_ = W5100
																																								tmp13088 := PrimIsPair(W5100)

																																								if True == tmp13088 {
																																									tmp13025 := MakeNative(func(__e *ControlFlow) {
																																										W5101 := __e.Get(1)
																																										_ = W5101
																																										tmp13026 := MakeNative(func(__e *ControlFlow) {
																																											W5102 := __e.Get(1)
																																											_ = W5102
																																											tmp13083 := PrimEqual(W5102, Nil)

																																											if True == tmp13083 {
																																												tmp13027 := MakeNative(func(__e *ControlFlow) {
																																													W5103 := __e.Get(1)
																																													_ = W5103
																																													tmp13028 := MakeNative(func(__e *ControlFlow) {
																																														W5104 := __e.Get(1)
																																														_ = W5104
																																														tmp13072 := PrimIsPair(W5103)

																																														if True == tmp13072 {
																																															tmp13029 := MakeNative(func(__e *ControlFlow) {
																																																W5106 := __e.Get(1)
																																																_ = W5106
																																																tmp13030 := MakeNative(func(__e *ControlFlow) {
																																																	W5107 := __e.Get(1)
																																																	_ = W5107
																																																	tmp13034 := PrimEqual(W5106, symstream)

																																																	if True == tmp13034 {
																																																		__e.TailApply(PrimFunc(symthaw), W5107)
																																																		return
																																																	} else {
																																																		tmp13032 := Call(__e, PrimFunc(symshen_4pvar_2), W5106)

																																																		if True == tmp13032 {
																																																			__e.TailApply(PrimFunc(symshen_4bind_b), W5106, symstream, V4934, W5107)
																																																			return
																																																		} else {
																																																			__e.Return(False)
																																																			return
																																																		}

																																																	}

																																																}, 1)

																																																tmp13035 := MakeNative(func(__e *ControlFlow) {
																																																	tmp13036 := MakeNative(func(__e *ControlFlow) {
																																																		W5108 := __e.Get(1)
																																																		_ = W5108
																																																		tmp13037 := MakeNative(func(__e *ControlFlow) {
																																																			W5109 := __e.Get(1)
																																																			_ = W5109
																																																			tmp13057 := PrimIsPair(W5108)

																																																			if True == tmp13057 {
																																																				tmp13038 := MakeNative(func(__e *ControlFlow) {
																																																					W5111 := __e.Get(1)
																																																					_ = W5111
																																																					tmp13039 := MakeNative(func(__e *ControlFlow) {
																																																						W5112 := __e.Get(1)
																																																						_ = W5112
																																																						tmp13040 := MakeNative(func(__e *ControlFlow) {
																																																							W5113 := __e.Get(1)
																																																							_ = W5113
																																																							tmp13044 := PrimEqual(W5112, Nil)

																																																							if True == tmp13044 {
																																																								__e.TailApply(PrimFunc(symthaw), W5113)
																																																								return
																																																							} else {
																																																								tmp13042 := Call(__e, PrimFunc(symshen_4pvar_2), W5112)

																																																								if True == tmp13042 {
																																																									__e.TailApply(PrimFunc(symshen_4bind_b), W5112, Nil, V4934, W5113)
																																																									return
																																																								} else {
																																																									__e.Return(False)
																																																									return
																																																								}

																																																							}

																																																						}, 1)

																																																						tmp13045 := MakeNative(func(__e *ControlFlow) {
																																																							__e.TailApply(W5109, W5111)
																																																							return
																																																						}, 0)

																																																						__e.TailApply(tmp13040, tmp13045)
																																																						return

																																																					}, 1)

																																																					tmp13046 := PrimTail(W5108)

																																																					tmp13047 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13046, V4934)

																																																					__e.TailApply(tmp13039, tmp13047)
																																																					return

																																																				}, 1)

																																																				tmp13048 := PrimHead(W5108)

																																																				__e.TailApply(tmp13038, tmp13048)
																																																				return

																																																			} else {
																																																				tmp13055 := Call(__e, PrimFunc(symshen_4pvar_2), W5108)

																																																				if True == tmp13055 {
																																																					tmp13049 := MakeNative(func(__e *ControlFlow) {
																																																						W5114 := __e.Get(1)
																																																						_ = W5114
																																																						tmp13050 := PrimCons(W5114, Nil)

																																																						tmp13051 := MakeNative(func(__e *ControlFlow) {
																																																							__e.TailApply(W5109, W5114)
																																																							return
																																																						}, 0)

																																																						tmp13052 := Call(__e, PrimFunc(symshen_4bind_b), W5108, tmp13050, V4934, tmp13051)

																																																						__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13052)
																																																						return

																																																					}, 1)

																																																					tmp13053 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																																					__e.TailApply(tmp13049, tmp13053)
																																																					return

																																																				} else {
																																																					__e.Return(False)
																																																					return
																																																				}

																																																			}

																																																		}, 1)

																																																		tmp13058 := MakeNative(func(__e *ControlFlow) {
																																																			Z5110 := __e.Get(1)
																																																			_ = Z5110
																																																			__e.TailApply(W5104, Z5110)
																																																			return
																																																		}, 1)

																																																		__e.TailApply(tmp13037, tmp13058)
																																																		return

																																																	}, 1)

																																																	tmp13059 := PrimTail(W5103)

																																																	tmp13060 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13059, V4934)

																																																	__e.TailApply(tmp13036, tmp13060)
																																																	return

																																																}, 0)

																																																__e.TailApply(tmp13030, tmp13035)
																																																return

																																															}, 1)

																																															tmp13061 := PrimHead(W5103)

																																															tmp13062 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13061, V4934)

																																															__e.TailApply(tmp13029, tmp13062)
																																															return

																																														} else {
																																															tmp13070 := Call(__e, PrimFunc(symshen_4pvar_2), W5103)

																																															if True == tmp13070 {
																																																tmp13063 := MakeNative(func(__e *ControlFlow) {
																																																	W5115 := __e.Get(1)
																																																	_ = W5115
																																																	tmp13064 := PrimCons(W5115, Nil)

																																																	tmp13065 := PrimCons(symstream, tmp13064)

																																																	tmp13066 := MakeNative(func(__e *ControlFlow) {
																																																		__e.TailApply(W5104, W5115)
																																																		return
																																																	}, 0)

																																																	tmp13067 := Call(__e, PrimFunc(symshen_4bind_b), W5103, tmp13065, V4934, tmp13066)

																																																	__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13067)
																																																	return

																																																}, 1)

																																																tmp13068 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																																__e.TailApply(tmp13063, tmp13068)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}

																																													}, 1)

																																													tmp13073 := MakeNative(func(__e *ControlFlow) {
																																														Z5105 := __e.Get(1)
																																														_ = Z5105
																																														tmp13074 := Call(__e, PrimFunc(symshen_4incinfs))

																																														_ = tmp13074

																																														tmp13075 := MakeNative(func(__e *ControlFlow) {
																																															tmp13076 := Call(__e, PrimFunc(symshen_4lazyderef), Z5105, V4934)

																																															tmp13077 := PrimCons(symout, Nil)

																																															tmp13078 := PrimCons(symin, tmp13077)

																																															tmp13079 := Call(__e, PrimFunc(symelement_2), tmp13076, tmp13078)

																																															tmp13080 := MakeNative(func(__e *ControlFlow) {
																																																__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5099, symstring, V4933, V4934, V4935, W4938, V4937)
																																																return
																																															}, 0)

																																															__e.TailApply(PrimFunc(symwhen), tmp13079, V4934, V4935, W4938, tmp13080)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimFunc(symis_b), W5101, Z5105, V4934, V4935, W4938, tmp13075)
																																														return

																																													}, 1)

																																													__e.TailApply(tmp13028, tmp13073)
																																													return

																																												}, 1)

																																												tmp13081 := Call(__e, PrimFunc(symshen_4lazyderef), V4932, V4934)

																																												__e.TailApply(tmp13027, tmp13081)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}, 1)

																																										tmp13084 := PrimTail(W5100)

																																										tmp13085 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13084, V4934)

																																										__e.TailApply(tmp13026, tmp13085)
																																										return

																																									}, 1)

																																									tmp13086 := PrimHead(W5100)

																																									__e.TailApply(tmp13025, tmp13086)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp13089 := PrimTail(W5098)

																																							tmp13090 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13089, V4934)

																																							__e.TailApply(tmp13024, tmp13090)
																																							return

																																						}, 1)

																																						tmp13091 := PrimHead(W5098)

																																						__e.TailApply(tmp13023, tmp13091)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp13094 := PrimTail(W5096)

																																				tmp13095 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13094, V4934)

																																				__e.TailApply(tmp13022, tmp13095)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp13098 := PrimHead(W5096)

																																		tmp13099 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13098, V4934)

																																		__e.TailApply(tmp13021, tmp13099)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13102 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																																tmp13103 := Call(__e, tmp13020, tmp13102)

																																ifres13019 = tmp13103

																															} else {
																																ifres13019 = False

																															}

																															__e.TailApply(tmp12874, ifres13019)
																															return

																														} else {
																															__e.Return(W5082)
																															return
																														}

																													}, 1)

																													tmp13168 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																													var ifres13107 Obj

																													if True == tmp13168 {
																														tmp13108 := MakeNative(func(__e *ControlFlow) {
																															W5083 := __e.Get(1)
																															_ = W5083
																															tmp13165 := PrimIsPair(W5083)

																															if True == tmp13165 {
																																tmp13109 := MakeNative(func(__e *ControlFlow) {
																																	W5084 := __e.Get(1)
																																	_ = W5084
																																	tmp13161 := PrimEqual(W5084, symlet)

																																	if True == tmp13161 {
																																		tmp13110 := MakeNative(func(__e *ControlFlow) {
																																			W5085 := __e.Get(1)
																																			_ = W5085
																																			tmp13157 := PrimIsPair(W5085)

																																			if True == tmp13157 {
																																				tmp13111 := MakeNative(func(__e *ControlFlow) {
																																					W5086 := __e.Get(1)
																																					_ = W5086
																																					tmp13112 := MakeNative(func(__e *ControlFlow) {
																																						W5087 := __e.Get(1)
																																						_ = W5087
																																						tmp13152 := PrimIsPair(W5087)

																																						if True == tmp13152 {
																																							tmp13113 := MakeNative(func(__e *ControlFlow) {
																																								W5088 := __e.Get(1)
																																								_ = W5088
																																								tmp13114 := MakeNative(func(__e *ControlFlow) {
																																									W5089 := __e.Get(1)
																																									_ = W5089
																																									tmp13147 := PrimIsPair(W5089)

																																									if True == tmp13147 {
																																										tmp13115 := MakeNative(func(__e *ControlFlow) {
																																											W5090 := __e.Get(1)
																																											_ = W5090
																																											tmp13116 := MakeNative(func(__e *ControlFlow) {
																																												W5091 := __e.Get(1)
																																												_ = W5091
																																												tmp13142 := PrimEqual(W5091, Nil)

																																												if True == tmp13142 {
																																													tmp13117 := MakeNative(func(__e *ControlFlow) {
																																														W5092 := __e.Get(1)
																																														_ = W5092
																																														tmp13118 := MakeNative(func(__e *ControlFlow) {
																																															W5093 := __e.Get(1)
																																															_ = W5093
																																															tmp13119 := MakeNative(func(__e *ControlFlow) {
																																																W5094 := __e.Get(1)
																																																_ = W5094
																																																tmp13120 := Call(__e, PrimFunc(symshen_4incinfs))

																																																_ = tmp13120

																																																tmp13121 := MakeNative(func(__e *ControlFlow) {
																																																	tmp13122 := Call(__e, PrimFunc(symshen_4lazyderef), W5086, V4934)

																																																	tmp13123 := Call(__e, PrimFunc(symshen_4freshterm), tmp13122)

																																																	tmp13124 := MakeNative(func(__e *ControlFlow) {
																																																		tmp13125 := Call(__e, PrimFunc(symshen_4lazyderef), W5086, V4934)

																																																		tmp13126 := Call(__e, PrimFunc(symshen_4lazyderef), W5093, V4934)

																																																		tmp13127 := Call(__e, PrimFunc(symshen_4lazyderef), W5090, V4934)

																																																		tmp13128 := Call(__e, PrimFunc(symshen_4beta), tmp13125, tmp13126, tmp13127)

																																																		tmp13129 := MakeNative(func(__e *ControlFlow) {
																																																			tmp13130 := PrimIntern(MakeString(":"))

																																																			tmp13131 := PrimCons(W5094, Nil)

																																																			tmp13132 := PrimCons(tmp13130, tmp13131)

																																																			tmp13133 := PrimCons(W5093, tmp13132)

																																																			tmp13134 := PrimCons(tmp13133, V4933)

																																																			__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5092, V4932, tmp13134, V4934, V4935, W4938, V4937)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimFunc(symbind), W5092, tmp13128, V4934, V4935, W4938, tmp13129)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimFunc(symbind), W5093, tmp13123, V4934, V4935, W4938, tmp13124)
																																																	return

																																																}, 0)

																																																tmp13135 := Call(__e, PrimFunc(symshen_4system_1S_1h), W5088, W5094, V4933, V4934, V4935, W4938, tmp13121)

																																																__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13135)
																																																return

																																															}, 1)

																																															tmp13136 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																															tmp13137 := Call(__e, tmp13119, tmp13136)

																																															__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13137)
																																															return

																																														}, 1)

																																														tmp13138 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																														tmp13139 := Call(__e, tmp13118, tmp13138)

																																														__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13139)
																																														return

																																													}, 1)

																																													tmp13140 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																													__e.TailApply(tmp13117, tmp13140)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp13143 := PrimTail(W5089)

																																											tmp13144 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13143, V4934)

																																											__e.TailApply(tmp13116, tmp13144)
																																											return

																																										}, 1)

																																										tmp13145 := PrimHead(W5089)

																																										__e.TailApply(tmp13115, tmp13145)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp13148 := PrimTail(W5087)

																																								tmp13149 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13148, V4934)

																																								__e.TailApply(tmp13114, tmp13149)
																																								return

																																							}, 1)

																																							tmp13150 := PrimHead(W5087)

																																							__e.TailApply(tmp13113, tmp13150)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp13153 := PrimTail(W5085)

																																					tmp13154 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13153, V4934)

																																					__e.TailApply(tmp13112, tmp13154)
																																					return

																																				}, 1)

																																				tmp13155 := PrimHead(W5085)

																																				__e.TailApply(tmp13111, tmp13155)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp13158 := PrimTail(W5083)

																																		tmp13159 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13158, V4934)

																																		__e.TailApply(tmp13110, tmp13159)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13162 := PrimHead(W5083)

																																tmp13163 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13162, V4934)

																																__e.TailApply(tmp13109, tmp13163)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13166 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																														tmp13167 := Call(__e, tmp13108, tmp13166)

																														ifres13107 = tmp13167

																													} else {
																														ifres13107 = False

																													}

																													__e.TailApply(tmp12873, ifres13107)
																													return

																												} else {
																													__e.Return(W5052)
																													return
																												}

																											}, 1)

																											tmp13292 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																											var ifres13171 Obj

																											if True == tmp13292 {
																												tmp13172 := MakeNative(func(__e *ControlFlow) {
																													W5053 := __e.Get(1)
																													_ = W5053
																													tmp13289 := PrimIsPair(W5053)

																													if True == tmp13289 {
																														tmp13173 := MakeNative(func(__e *ControlFlow) {
																															W5054 := __e.Get(1)
																															_ = W5054
																															tmp13285 := PrimEqual(W5054, symlambda)

																															if True == tmp13285 {
																																tmp13174 := MakeNative(func(__e *ControlFlow) {
																																	W5055 := __e.Get(1)
																																	_ = W5055
																																	tmp13281 := PrimIsPair(W5055)

																																	if True == tmp13281 {
																																		tmp13175 := MakeNative(func(__e *ControlFlow) {
																																			W5056 := __e.Get(1)
																																			_ = W5056
																																			tmp13176 := MakeNative(func(__e *ControlFlow) {
																																				W5057 := __e.Get(1)
																																				_ = W5057
																																				tmp13276 := PrimIsPair(W5057)

																																				if True == tmp13276 {
																																					tmp13177 := MakeNative(func(__e *ControlFlow) {
																																						W5058 := __e.Get(1)
																																						_ = W5058
																																						tmp13178 := MakeNative(func(__e *ControlFlow) {
																																							W5059 := __e.Get(1)
																																							_ = W5059
																																							tmp13271 := PrimEqual(W5059, Nil)

																																							if True == tmp13271 {
																																								tmp13179 := MakeNative(func(__e *ControlFlow) {
																																									W5060 := __e.Get(1)
																																									_ = W5060
																																									tmp13180 := MakeNative(func(__e *ControlFlow) {
																																										W5061 := __e.Get(1)
																																										_ = W5061
																																										tmp13247 := PrimIsPair(W5060)

																																										if True == tmp13247 {
																																											tmp13181 := MakeNative(func(__e *ControlFlow) {
																																												W5066 := __e.Get(1)
																																												_ = W5066
																																												tmp13182 := MakeNative(func(__e *ControlFlow) {
																																													W5067 := __e.Get(1)
																																													_ = W5067
																																													tmp13183 := MakeNative(func(__e *ControlFlow) {
																																														W5068 := __e.Get(1)
																																														_ = W5068
																																														tmp13227 := PrimIsPair(W5067)

																																														if True == tmp13227 {
																																															tmp13184 := MakeNative(func(__e *ControlFlow) {
																																																W5070 := __e.Get(1)
																																																_ = W5070
																																																tmp13185 := MakeNative(func(__e *ControlFlow) {
																																																	W5071 := __e.Get(1)
																																																	_ = W5071
																																																	tmp13189 := PrimEqual(W5070, sym_1_1_6)

																																																	if True == tmp13189 {
																																																		__e.TailApply(PrimFunc(symthaw), W5071)
																																																		return
																																																	} else {
																																																		tmp13187 := Call(__e, PrimFunc(symshen_4pvar_2), W5070)

																																																		if True == tmp13187 {
																																																			__e.TailApply(PrimFunc(symshen_4bind_b), W5070, sym_1_1_6, V4934, W5071)
																																																			return
																																																		} else {
																																																			__e.Return(False)
																																																			return
																																																		}

																																																	}

																																																}, 1)

																																																tmp13190 := MakeNative(func(__e *ControlFlow) {
																																																	tmp13191 := MakeNative(func(__e *ControlFlow) {
																																																		W5072 := __e.Get(1)
																																																		_ = W5072
																																																		tmp13192 := MakeNative(func(__e *ControlFlow) {
																																																			W5073 := __e.Get(1)
																																																			_ = W5073
																																																			tmp13212 := PrimIsPair(W5072)

																																																			if True == tmp13212 {
																																																				tmp13193 := MakeNative(func(__e *ControlFlow) {
																																																					W5075 := __e.Get(1)
																																																					_ = W5075
																																																					tmp13194 := MakeNative(func(__e *ControlFlow) {
																																																						W5076 := __e.Get(1)
																																																						_ = W5076
																																																						tmp13195 := MakeNative(func(__e *ControlFlow) {
																																																							W5077 := __e.Get(1)
																																																							_ = W5077
																																																							tmp13199 := PrimEqual(W5076, Nil)

																																																							if True == tmp13199 {
																																																								__e.TailApply(PrimFunc(symthaw), W5077)
																																																								return
																																																							} else {
																																																								tmp13197 := Call(__e, PrimFunc(symshen_4pvar_2), W5076)

																																																								if True == tmp13197 {
																																																									__e.TailApply(PrimFunc(symshen_4bind_b), W5076, Nil, V4934, W5077)
																																																									return
																																																								} else {
																																																									__e.Return(False)
																																																									return
																																																								}

																																																							}

																																																						}, 1)

																																																						tmp13200 := MakeNative(func(__e *ControlFlow) {
																																																							__e.TailApply(W5073, W5075)
																																																							return
																																																						}, 0)

																																																						__e.TailApply(tmp13195, tmp13200)
																																																						return

																																																					}, 1)

																																																					tmp13201 := PrimTail(W5072)

																																																					tmp13202 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13201, V4934)

																																																					__e.TailApply(tmp13194, tmp13202)
																																																					return

																																																				}, 1)

																																																				tmp13203 := PrimHead(W5072)

																																																				__e.TailApply(tmp13193, tmp13203)
																																																				return

																																																			} else {
																																																				tmp13210 := Call(__e, PrimFunc(symshen_4pvar_2), W5072)

																																																				if True == tmp13210 {
																																																					tmp13204 := MakeNative(func(__e *ControlFlow) {
																																																						W5078 := __e.Get(1)
																																																						_ = W5078
																																																						tmp13205 := PrimCons(W5078, Nil)

																																																						tmp13206 := MakeNative(func(__e *ControlFlow) {
																																																							__e.TailApply(W5073, W5078)
																																																							return
																																																						}, 0)

																																																						tmp13207 := Call(__e, PrimFunc(symshen_4bind_b), W5072, tmp13205, V4934, tmp13206)

																																																						__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13207)
																																																						return

																																																					}, 1)

																																																					tmp13208 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																																					__e.TailApply(tmp13204, tmp13208)
																																																					return

																																																				} else {
																																																					__e.Return(False)
																																																					return
																																																				}

																																																			}

																																																		}, 1)

																																																		tmp13213 := MakeNative(func(__e *ControlFlow) {
																																																			Z5074 := __e.Get(1)
																																																			_ = Z5074
																																																			__e.TailApply(W5068, Z5074)
																																																			return
																																																		}, 1)

																																																		__e.TailApply(tmp13192, tmp13213)
																																																		return

																																																	}, 1)

																																																	tmp13214 := PrimTail(W5067)

																																																	tmp13215 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13214, V4934)

																																																	__e.TailApply(tmp13191, tmp13215)
																																																	return

																																																}, 0)

																																																__e.TailApply(tmp13185, tmp13190)
																																																return

																																															}, 1)

																																															tmp13216 := PrimHead(W5067)

																																															tmp13217 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13216, V4934)

																																															__e.TailApply(tmp13184, tmp13217)
																																															return

																																														} else {
																																															tmp13225 := Call(__e, PrimFunc(symshen_4pvar_2), W5067)

																																															if True == tmp13225 {
																																																tmp13218 := MakeNative(func(__e *ControlFlow) {
																																																	W5079 := __e.Get(1)
																																																	_ = W5079
																																																	tmp13219 := PrimCons(W5079, Nil)

																																																	tmp13220 := PrimCons(sym_1_1_6, tmp13219)

																																																	tmp13221 := MakeNative(func(__e *ControlFlow) {
																																																		__e.TailApply(W5068, W5079)
																																																		return
																																																	}, 0)

																																																	tmp13222 := Call(__e, PrimFunc(symshen_4bind_b), W5067, tmp13220, V4934, tmp13221)

																																																	__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13222)
																																																	return

																																																}, 1)

																																																tmp13223 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																																__e.TailApply(tmp13218, tmp13223)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}

																																													}, 1)

																																													tmp13228 := MakeNative(func(__e *ControlFlow) {
																																														Z5069 := __e.Get(1)
																																														_ = Z5069
																																														tmp13229 := Call(__e, W5061, W5066)

																																														__e.TailApply(tmp13229, Z5069)
																																														return

																																													}, 1)

																																													__e.TailApply(tmp13183, tmp13228)
																																													return

																																												}, 1)

																																												tmp13230 := PrimTail(W5060)

																																												tmp13231 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13230, V4934)

																																												__e.TailApply(tmp13182, tmp13231)
																																												return

																																											}, 1)

																																											tmp13232 := PrimHead(W5060)

																																											__e.TailApply(tmp13181, tmp13232)
																																											return

																																										} else {
																																											tmp13245 := Call(__e, PrimFunc(symshen_4pvar_2), W5060)

																																											if True == tmp13245 {
																																												tmp13233 := MakeNative(func(__e *ControlFlow) {
																																													W5080 := __e.Get(1)
																																													_ = W5080
																																													tmp13234 := MakeNative(func(__e *ControlFlow) {
																																														W5081 := __e.Get(1)
																																														_ = W5081
																																														tmp13235 := PrimCons(W5081, Nil)

																																														tmp13236 := PrimCons(sym_1_1_6, tmp13235)

																																														tmp13237 := PrimCons(W5080, tmp13236)

																																														tmp13238 := MakeNative(func(__e *ControlFlow) {
																																															tmp13239 := Call(__e, W5061, W5080)

																																															__e.TailApply(tmp13239, W5081)
																																															return

																																														}, 0)

																																														tmp13240 := Call(__e, PrimFunc(symshen_4bind_b), W5060, tmp13237, V4934, tmp13238)

																																														__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13240)
																																														return

																																													}, 1)

																																													tmp13241 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																													tmp13242 := Call(__e, tmp13234, tmp13241)

																																													__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13242)
																																													return

																																												}, 1)

																																												tmp13243 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																												__e.TailApply(tmp13233, tmp13243)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}

																																									}, 1)

																																									tmp13248 := MakeNative(func(__e *ControlFlow) {
																																										Z5062 := __e.Get(1)
																																										_ = Z5062
																																										__e.Return(MakeNative(func(__e *ControlFlow) {
																																											Z5063 := __e.Get(1)
																																											_ = Z5063
																																											tmp13249 := MakeNative(func(__e *ControlFlow) {
																																												W5064 := __e.Get(1)
																																												_ = W5064
																																												tmp13250 := MakeNative(func(__e *ControlFlow) {
																																													W5065 := __e.Get(1)
																																													_ = W5065
																																													tmp13251 := Call(__e, PrimFunc(symshen_4incinfs))

																																													_ = tmp13251

																																													tmp13252 := Call(__e, PrimFunc(symshen_4lazyderef), W5056, V4934)

																																													tmp13253 := Call(__e, PrimFunc(symshen_4freshterm), tmp13252)

																																													tmp13254 := MakeNative(func(__e *ControlFlow) {
																																														tmp13255 := Call(__e, PrimFunc(symshen_4lazyderef), W5056, V4934)

																																														tmp13256 := Call(__e, PrimFunc(symshen_4deref), W5065, V4934)

																																														tmp13257 := Call(__e, PrimFunc(symshen_4deref), W5058, V4934)

																																														tmp13258 := Call(__e, PrimFunc(symshen_4beta), tmp13255, tmp13256, tmp13257)

																																														tmp13259 := MakeNative(func(__e *ControlFlow) {
																																															tmp13260 := PrimIntern(MakeString(":"))

																																															tmp13261 := PrimCons(Z5062, Nil)

																																															tmp13262 := PrimCons(tmp13260, tmp13261)

																																															tmp13263 := PrimCons(W5065, tmp13262)

																																															tmp13264 := PrimCons(tmp13263, V4933)

																																															__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5064, Z5063, tmp13264, V4934, V4935, W4938, V4937)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimFunc(symbind), W5064, tmp13258, V4934, V4935, W4938, tmp13259)
																																														return

																																													}, 0)

																																													tmp13265 := Call(__e, PrimFunc(symbind), W5065, tmp13253, V4934, V4935, W4938, tmp13254)

																																													__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13265)
																																													return

																																												}, 1)

																																												tmp13266 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																												tmp13267 := Call(__e, tmp13250, tmp13266)

																																												__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13267)
																																												return

																																											}, 1)

																																											tmp13268 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																											__e.TailApply(tmp13249, tmp13268)
																																											return

																																										}, 1))
																																										return
																																									}, 1)

																																									__e.TailApply(tmp13180, tmp13248)
																																									return

																																								}, 1)

																																								tmp13269 := Call(__e, PrimFunc(symshen_4lazyderef), V4932, V4934)

																																								__e.TailApply(tmp13179, tmp13269)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp13272 := PrimTail(W5057)

																																						tmp13273 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13272, V4934)

																																						__e.TailApply(tmp13178, tmp13273)
																																						return

																																					}, 1)

																																					tmp13274 := PrimHead(W5057)

																																					__e.TailApply(tmp13177, tmp13274)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp13277 := PrimTail(W5055)

																																			tmp13278 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13277, V4934)

																																			__e.TailApply(tmp13176, tmp13278)
																																			return

																																		}, 1)

																																		tmp13279 := PrimHead(W5055)

																																		__e.TailApply(tmp13175, tmp13279)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13282 := PrimTail(W5053)

																																tmp13283 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13282, V4934)

																																__e.TailApply(tmp13174, tmp13283)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13286 := PrimHead(W5053)

																														tmp13287 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13286, V4934)

																														__e.TailApply(tmp13173, tmp13287)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13290 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																												tmp13291 := Call(__e, tmp13172, tmp13290)

																												ifres13171 = tmp13291

																											} else {
																												ifres13171 = False

																											}

																											__e.TailApply(tmp12872, ifres13171)
																											return

																										} else {
																											__e.Return(W5042)
																											return
																										}

																									}, 1)

																									tmp13335 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																									var ifres13295 Obj

																									if True == tmp13335 {
																										tmp13296 := MakeNative(func(__e *ControlFlow) {
																											W5043 := __e.Get(1)
																											_ = W5043
																											tmp13332 := PrimIsPair(W5043)

																											if True == tmp13332 {
																												tmp13297 := MakeNative(func(__e *ControlFlow) {
																													W5044 := __e.Get(1)
																													_ = W5044
																													tmp13328 := PrimEqual(W5044, sym_8s)

																													if True == tmp13328 {
																														tmp13298 := MakeNative(func(__e *ControlFlow) {
																															W5045 := __e.Get(1)
																															_ = W5045
																															tmp13324 := PrimIsPair(W5045)

																															if True == tmp13324 {
																																tmp13299 := MakeNative(func(__e *ControlFlow) {
																																	W5046 := __e.Get(1)
																																	_ = W5046
																																	tmp13300 := MakeNative(func(__e *ControlFlow) {
																																		W5047 := __e.Get(1)
																																		_ = W5047
																																		tmp13319 := PrimIsPair(W5047)

																																		if True == tmp13319 {
																																			tmp13301 := MakeNative(func(__e *ControlFlow) {
																																				W5048 := __e.Get(1)
																																				_ = W5048
																																				tmp13302 := MakeNative(func(__e *ControlFlow) {
																																					W5049 := __e.Get(1)
																																					_ = W5049
																																					tmp13314 := PrimEqual(W5049, Nil)

																																					if True == tmp13314 {
																																						tmp13303 := MakeNative(func(__e *ControlFlow) {
																																							W5050 := __e.Get(1)
																																							_ = W5050
																																							tmp13304 := MakeNative(func(__e *ControlFlow) {
																																								W5051 := __e.Get(1)
																																								_ = W5051
																																								tmp13308 := PrimEqual(W5050, symstring)

																																								if True == tmp13308 {
																																									__e.TailApply(PrimFunc(symthaw), W5051)
																																									return
																																								} else {
																																									tmp13306 := Call(__e, PrimFunc(symshen_4pvar_2), W5050)

																																									if True == tmp13306 {
																																										__e.TailApply(PrimFunc(symshen_4bind_b), W5050, symstring, V4934, W5051)
																																										return
																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp13309 := MakeNative(func(__e *ControlFlow) {
																																								tmp13310 := Call(__e, PrimFunc(symshen_4incinfs))

																																								_ = tmp13310

																																								tmp13311 := MakeNative(func(__e *ControlFlow) {
																																									__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5048, symstring, V4933, V4934, V4935, W4938, V4937)
																																									return
																																								}, 0)

																																								__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5046, symstring, V4933, V4934, V4935, W4938, tmp13311)
																																								return

																																							}, 0)

																																							__e.TailApply(tmp13304, tmp13309)
																																							return

																																						}, 1)

																																						tmp13312 := Call(__e, PrimFunc(symshen_4lazyderef), V4932, V4934)

																																						__e.TailApply(tmp13303, tmp13312)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp13315 := PrimTail(W5047)

																																				tmp13316 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13315, V4934)

																																				__e.TailApply(tmp13302, tmp13316)
																																				return

																																			}, 1)

																																			tmp13317 := PrimHead(W5047)

																																			__e.TailApply(tmp13301, tmp13317)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp13320 := PrimTail(W5045)

																																	tmp13321 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13320, V4934)

																																	__e.TailApply(tmp13300, tmp13321)
																																	return

																																}, 1)

																																tmp13322 := PrimHead(W5045)

																																__e.TailApply(tmp13299, tmp13322)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13325 := PrimTail(W5043)

																														tmp13326 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13325, V4934)

																														__e.TailApply(tmp13298, tmp13326)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13329 := PrimHead(W5043)

																												tmp13330 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13329, V4934)

																												__e.TailApply(tmp13297, tmp13330)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp13333 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																										tmp13334 := Call(__e, tmp13296, tmp13333)

																										ifres13295 = tmp13334

																									} else {
																										ifres13295 = False

																									}

																									__e.TailApply(tmp12871, ifres13295)
																									return

																								} else {
																									__e.Return(W5021)
																									return
																								}

																							}, 1)

																							tmp13420 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																							var ifres13338 Obj

																							if True == tmp13420 {
																								tmp13339 := MakeNative(func(__e *ControlFlow) {
																									W5022 := __e.Get(1)
																									_ = W5022
																									tmp13417 := PrimIsPair(W5022)

																									if True == tmp13417 {
																										tmp13340 := MakeNative(func(__e *ControlFlow) {
																											W5023 := __e.Get(1)
																											_ = W5023
																											tmp13413 := PrimEqual(W5023, sym_8v)

																											if True == tmp13413 {
																												tmp13341 := MakeNative(func(__e *ControlFlow) {
																													W5024 := __e.Get(1)
																													_ = W5024
																													tmp13409 := PrimIsPair(W5024)

																													if True == tmp13409 {
																														tmp13342 := MakeNative(func(__e *ControlFlow) {
																															W5025 := __e.Get(1)
																															_ = W5025
																															tmp13343 := MakeNative(func(__e *ControlFlow) {
																																W5026 := __e.Get(1)
																																_ = W5026
																																tmp13404 := PrimIsPair(W5026)

																																if True == tmp13404 {
																																	tmp13344 := MakeNative(func(__e *ControlFlow) {
																																		W5027 := __e.Get(1)
																																		_ = W5027
																																		tmp13345 := MakeNative(func(__e *ControlFlow) {
																																			W5028 := __e.Get(1)
																																			_ = W5028
																																			tmp13399 := PrimEqual(W5028, Nil)

																																			if True == tmp13399 {
																																				tmp13346 := MakeNative(func(__e *ControlFlow) {
																																					W5029 := __e.Get(1)
																																					_ = W5029
																																					tmp13347 := MakeNative(func(__e *ControlFlow) {
																																						W5030 := __e.Get(1)
																																						_ = W5030
																																						tmp13391 := PrimIsPair(W5029)

																																						if True == tmp13391 {
																																							tmp13348 := MakeNative(func(__e *ControlFlow) {
																																								W5032 := __e.Get(1)
																																								_ = W5032
																																								tmp13349 := MakeNative(func(__e *ControlFlow) {
																																									W5033 := __e.Get(1)
																																									_ = W5033
																																									tmp13353 := PrimEqual(W5032, symvector)

																																									if True == tmp13353 {
																																										__e.TailApply(PrimFunc(symthaw), W5033)
																																										return
																																									} else {
																																										tmp13351 := Call(__e, PrimFunc(symshen_4pvar_2), W5032)

																																										if True == tmp13351 {
																																											__e.TailApply(PrimFunc(symshen_4bind_b), W5032, symvector, V4934, W5033)
																																											return
																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp13354 := MakeNative(func(__e *ControlFlow) {
																																									tmp13355 := MakeNative(func(__e *ControlFlow) {
																																										W5034 := __e.Get(1)
																																										_ = W5034
																																										tmp13356 := MakeNative(func(__e *ControlFlow) {
																																											W5035 := __e.Get(1)
																																											_ = W5035
																																											tmp13376 := PrimIsPair(W5034)

																																											if True == tmp13376 {
																																												tmp13357 := MakeNative(func(__e *ControlFlow) {
																																													W5037 := __e.Get(1)
																																													_ = W5037
																																													tmp13358 := MakeNative(func(__e *ControlFlow) {
																																														W5038 := __e.Get(1)
																																														_ = W5038
																																														tmp13359 := MakeNative(func(__e *ControlFlow) {
																																															W5039 := __e.Get(1)
																																															_ = W5039
																																															tmp13363 := PrimEqual(W5038, Nil)

																																															if True == tmp13363 {
																																																__e.TailApply(PrimFunc(symthaw), W5039)
																																																return
																																															} else {
																																																tmp13361 := Call(__e, PrimFunc(symshen_4pvar_2), W5038)

																																																if True == tmp13361 {
																																																	__e.TailApply(PrimFunc(symshen_4bind_b), W5038, Nil, V4934, W5039)
																																																	return
																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp13364 := MakeNative(func(__e *ControlFlow) {
																																															__e.TailApply(W5035, W5037)
																																															return
																																														}, 0)

																																														__e.TailApply(tmp13359, tmp13364)
																																														return

																																													}, 1)

																																													tmp13365 := PrimTail(W5034)

																																													tmp13366 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13365, V4934)

																																													__e.TailApply(tmp13358, tmp13366)
																																													return

																																												}, 1)

																																												tmp13367 := PrimHead(W5034)

																																												__e.TailApply(tmp13357, tmp13367)
																																												return

																																											} else {
																																												tmp13374 := Call(__e, PrimFunc(symshen_4pvar_2), W5034)

																																												if True == tmp13374 {
																																													tmp13368 := MakeNative(func(__e *ControlFlow) {
																																														W5040 := __e.Get(1)
																																														_ = W5040
																																														tmp13369 := PrimCons(W5040, Nil)

																																														tmp13370 := MakeNative(func(__e *ControlFlow) {
																																															__e.TailApply(W5035, W5040)
																																															return
																																														}, 0)

																																														tmp13371 := Call(__e, PrimFunc(symshen_4bind_b), W5034, tmp13369, V4934, tmp13370)

																																														__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13371)
																																														return

																																													}, 1)

																																													tmp13372 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																													__e.TailApply(tmp13368, tmp13372)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}

																																										}, 1)

																																										tmp13377 := MakeNative(func(__e *ControlFlow) {
																																											Z5036 := __e.Get(1)
																																											_ = Z5036
																																											__e.TailApply(W5030, Z5036)
																																											return
																																										}, 1)

																																										__e.TailApply(tmp13356, tmp13377)
																																										return

																																									}, 1)

																																									tmp13378 := PrimTail(W5029)

																																									tmp13379 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13378, V4934)

																																									__e.TailApply(tmp13355, tmp13379)
																																									return

																																								}, 0)

																																								__e.TailApply(tmp13349, tmp13354)
																																								return

																																							}, 1)

																																							tmp13380 := PrimHead(W5029)

																																							tmp13381 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13380, V4934)

																																							__e.TailApply(tmp13348, tmp13381)
																																							return

																																						} else {
																																							tmp13389 := Call(__e, PrimFunc(symshen_4pvar_2), W5029)

																																							if True == tmp13389 {
																																								tmp13382 := MakeNative(func(__e *ControlFlow) {
																																									W5041 := __e.Get(1)
																																									_ = W5041
																																									tmp13383 := PrimCons(W5041, Nil)

																																									tmp13384 := PrimCons(symvector, tmp13383)

																																									tmp13385 := MakeNative(func(__e *ControlFlow) {
																																										__e.TailApply(W5030, W5041)
																																										return
																																									}, 0)

																																									tmp13386 := Call(__e, PrimFunc(symshen_4bind_b), W5029, tmp13384, V4934, tmp13385)

																																									__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13386)
																																									return

																																								}, 1)

																																								tmp13387 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																								__e.TailApply(tmp13382, tmp13387)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp13392 := MakeNative(func(__e *ControlFlow) {
																																						Z5031 := __e.Get(1)
																																						_ = Z5031
																																						tmp13393 := Call(__e, PrimFunc(symshen_4incinfs))

																																						_ = tmp13393

																																						tmp13394 := MakeNative(func(__e *ControlFlow) {
																																							tmp13395 := PrimCons(Z5031, Nil)

																																							tmp13396 := PrimCons(symvector, tmp13395)

																																							__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5027, tmp13396, V4933, V4934, V4935, W4938, V4937)
																																							return

																																						}, 0)

																																						__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5025, Z5031, V4933, V4934, V4935, W4938, tmp13394)
																																						return

																																					}, 1)

																																					__e.TailApply(tmp13347, tmp13392)
																																					return

																																				}, 1)

																																				tmp13397 := Call(__e, PrimFunc(symshen_4lazyderef), V4932, V4934)

																																				__e.TailApply(tmp13346, tmp13397)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp13400 := PrimTail(W5026)

																																		tmp13401 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13400, V4934)

																																		__e.TailApply(tmp13345, tmp13401)
																																		return

																																	}, 1)

																																	tmp13402 := PrimHead(W5026)

																																	__e.TailApply(tmp13344, tmp13402)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp13405 := PrimTail(W5024)

																															tmp13406 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13405, V4934)

																															__e.TailApply(tmp13343, tmp13406)
																															return

																														}, 1)

																														tmp13407 := PrimHead(W5024)

																														__e.TailApply(tmp13342, tmp13407)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp13410 := PrimTail(W5022)

																												tmp13411 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13410, V4934)

																												__e.TailApply(tmp13341, tmp13411)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp13414 := PrimHead(W5022)

																										tmp13415 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13414, V4934)

																										__e.TailApply(tmp13340, tmp13415)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13418 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																								tmp13419 := Call(__e, tmp13339, tmp13418)

																								ifres13338 = tmp13419

																							} else {
																								ifres13338 = False

																							}

																							__e.TailApply(tmp12870, ifres13338)
																							return

																						} else {
																							__e.Return(W4993)
																							return
																						}

																					}, 1)

																					tmp13526 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																					var ifres13423 Obj

																					if True == tmp13526 {
																						tmp13424 := MakeNative(func(__e *ControlFlow) {
																							W4994 := __e.Get(1)
																							_ = W4994
																							tmp13523 := PrimIsPair(W4994)

																							if True == tmp13523 {
																								tmp13425 := MakeNative(func(__e *ControlFlow) {
																									W4995 := __e.Get(1)
																									_ = W4995
																									tmp13519 := PrimEqual(W4995, sym_8p)

																									if True == tmp13519 {
																										tmp13426 := MakeNative(func(__e *ControlFlow) {
																											W4996 := __e.Get(1)
																											_ = W4996
																											tmp13515 := PrimIsPair(W4996)

																											if True == tmp13515 {
																												tmp13427 := MakeNative(func(__e *ControlFlow) {
																													W4997 := __e.Get(1)
																													_ = W4997
																													tmp13428 := MakeNative(func(__e *ControlFlow) {
																														W4998 := __e.Get(1)
																														_ = W4998
																														tmp13510 := PrimIsPair(W4998)

																														if True == tmp13510 {
																															tmp13429 := MakeNative(func(__e *ControlFlow) {
																																W4999 := __e.Get(1)
																																_ = W4999
																																tmp13430 := MakeNative(func(__e *ControlFlow) {
																																	W5000 := __e.Get(1)
																																	_ = W5000
																																	tmp13505 := PrimEqual(W5000, Nil)

																																	if True == tmp13505 {
																																		tmp13431 := MakeNative(func(__e *ControlFlow) {
																																			W5001 := __e.Get(1)
																																			_ = W5001
																																			tmp13432 := MakeNative(func(__e *ControlFlow) {
																																				W5002 := __e.Get(1)
																																				_ = W5002
																																				tmp13499 := PrimIsPair(W5001)

																																				if True == tmp13499 {
																																					tmp13433 := MakeNative(func(__e *ControlFlow) {
																																						W5005 := __e.Get(1)
																																						_ = W5005
																																						tmp13434 := MakeNative(func(__e *ControlFlow) {
																																							W5006 := __e.Get(1)
																																							_ = W5006
																																							tmp13435 := MakeNative(func(__e *ControlFlow) {
																																								W5007 := __e.Get(1)
																																								_ = W5007
																																								tmp13479 := PrimIsPair(W5006)

																																								if True == tmp13479 {
																																									tmp13436 := MakeNative(func(__e *ControlFlow) {
																																										W5009 := __e.Get(1)
																																										_ = W5009
																																										tmp13437 := MakeNative(func(__e *ControlFlow) {
																																											W5010 := __e.Get(1)
																																											_ = W5010
																																											tmp13441 := PrimEqual(W5009, sym_d)

																																											if True == tmp13441 {
																																												__e.TailApply(PrimFunc(symthaw), W5010)
																																												return
																																											} else {
																																												tmp13439 := Call(__e, PrimFunc(symshen_4pvar_2), W5009)

																																												if True == tmp13439 {
																																													__e.TailApply(PrimFunc(symshen_4bind_b), W5009, sym_d, V4934, W5010)
																																													return
																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}

																																										}, 1)

																																										tmp13442 := MakeNative(func(__e *ControlFlow) {
																																											tmp13443 := MakeNative(func(__e *ControlFlow) {
																																												W5011 := __e.Get(1)
																																												_ = W5011
																																												tmp13444 := MakeNative(func(__e *ControlFlow) {
																																													W5012 := __e.Get(1)
																																													_ = W5012
																																													tmp13464 := PrimIsPair(W5011)

																																													if True == tmp13464 {
																																														tmp13445 := MakeNative(func(__e *ControlFlow) {
																																															W5014 := __e.Get(1)
																																															_ = W5014
																																															tmp13446 := MakeNative(func(__e *ControlFlow) {
																																																W5015 := __e.Get(1)
																																																_ = W5015
																																																tmp13447 := MakeNative(func(__e *ControlFlow) {
																																																	W5016 := __e.Get(1)
																																																	_ = W5016
																																																	tmp13451 := PrimEqual(W5015, Nil)

																																																	if True == tmp13451 {
																																																		__e.TailApply(PrimFunc(symthaw), W5016)
																																																		return
																																																	} else {
																																																		tmp13449 := Call(__e, PrimFunc(symshen_4pvar_2), W5015)

																																																		if True == tmp13449 {
																																																			__e.TailApply(PrimFunc(symshen_4bind_b), W5015, Nil, V4934, W5016)
																																																			return
																																																		} else {
																																																			__e.Return(False)
																																																			return
																																																		}

																																																	}

																																																}, 1)

																																																tmp13452 := MakeNative(func(__e *ControlFlow) {
																																																	__e.TailApply(W5012, W5014)
																																																	return
																																																}, 0)

																																																__e.TailApply(tmp13447, tmp13452)
																																																return

																																															}, 1)

																																															tmp13453 := PrimTail(W5011)

																																															tmp13454 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13453, V4934)

																																															__e.TailApply(tmp13446, tmp13454)
																																															return

																																														}, 1)

																																														tmp13455 := PrimHead(W5011)

																																														__e.TailApply(tmp13445, tmp13455)
																																														return

																																													} else {
																																														tmp13462 := Call(__e, PrimFunc(symshen_4pvar_2), W5011)

																																														if True == tmp13462 {
																																															tmp13456 := MakeNative(func(__e *ControlFlow) {
																																																W5017 := __e.Get(1)
																																																_ = W5017
																																																tmp13457 := PrimCons(W5017, Nil)

																																																tmp13458 := MakeNative(func(__e *ControlFlow) {
																																																	__e.TailApply(W5012, W5017)
																																																	return
																																																}, 0)

																																																tmp13459 := Call(__e, PrimFunc(symshen_4bind_b), W5011, tmp13457, V4934, tmp13458)

																																																__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13459)
																																																return

																																															}, 1)

																																															tmp13460 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																															__e.TailApply(tmp13456, tmp13460)
																																															return

																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}

																																												}, 1)

																																												tmp13465 := MakeNative(func(__e *ControlFlow) {
																																													Z5013 := __e.Get(1)
																																													_ = Z5013
																																													__e.TailApply(W5007, Z5013)
																																													return
																																												}, 1)

																																												__e.TailApply(tmp13444, tmp13465)
																																												return

																																											}, 1)

																																											tmp13466 := PrimTail(W5006)

																																											tmp13467 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13466, V4934)

																																											__e.TailApply(tmp13443, tmp13467)
																																											return

																																										}, 0)

																																										__e.TailApply(tmp13437, tmp13442)
																																										return

																																									}, 1)

																																									tmp13468 := PrimHead(W5006)

																																									tmp13469 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13468, V4934)

																																									__e.TailApply(tmp13436, tmp13469)
																																									return

																																								} else {
																																									tmp13477 := Call(__e, PrimFunc(symshen_4pvar_2), W5006)

																																									if True == tmp13477 {
																																										tmp13470 := MakeNative(func(__e *ControlFlow) {
																																											W5018 := __e.Get(1)
																																											_ = W5018
																																											tmp13471 := PrimCons(W5018, Nil)

																																											tmp13472 := PrimCons(sym_d, tmp13471)

																																											tmp13473 := MakeNative(func(__e *ControlFlow) {
																																												__e.TailApply(W5007, W5018)
																																												return
																																											}, 0)

																																											tmp13474 := Call(__e, PrimFunc(symshen_4bind_b), W5006, tmp13472, V4934, tmp13473)

																																											__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13474)
																																											return

																																										}, 1)

																																										tmp13475 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																										__e.TailApply(tmp13470, tmp13475)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp13480 := MakeNative(func(__e *ControlFlow) {
																																								Z5008 := __e.Get(1)
																																								_ = Z5008
																																								tmp13481 := Call(__e, W5002, W5005)

																																								__e.TailApply(tmp13481, Z5008)
																																								return

																																							}, 1)

																																							__e.TailApply(tmp13435, tmp13480)
																																							return

																																						}, 1)

																																						tmp13482 := PrimTail(W5001)

																																						tmp13483 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13482, V4934)

																																						__e.TailApply(tmp13434, tmp13483)
																																						return

																																					}, 1)

																																					tmp13484 := PrimHead(W5001)

																																					__e.TailApply(tmp13433, tmp13484)
																																					return

																																				} else {
																																					tmp13497 := Call(__e, PrimFunc(symshen_4pvar_2), W5001)

																																					if True == tmp13497 {
																																						tmp13485 := MakeNative(func(__e *ControlFlow) {
																																							W5019 := __e.Get(1)
																																							_ = W5019
																																							tmp13486 := MakeNative(func(__e *ControlFlow) {
																																								W5020 := __e.Get(1)
																																								_ = W5020
																																								tmp13487 := PrimCons(W5020, Nil)

																																								tmp13488 := PrimCons(sym_d, tmp13487)

																																								tmp13489 := PrimCons(W5019, tmp13488)

																																								tmp13490 := MakeNative(func(__e *ControlFlow) {
																																									tmp13491 := Call(__e, W5002, W5019)

																																									__e.TailApply(tmp13491, W5020)
																																									return

																																								}, 0)

																																								tmp13492 := Call(__e, PrimFunc(symshen_4bind_b), W5001, tmp13489, V4934, tmp13490)

																																								__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13492)
																																								return

																																							}, 1)

																																							tmp13493 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																							tmp13494 := Call(__e, tmp13486, tmp13493)

																																							__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13494)
																																							return

																																						}, 1)

																																						tmp13495 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																						__e.TailApply(tmp13485, tmp13495)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp13500 := MakeNative(func(__e *ControlFlow) {
																																				Z5003 := __e.Get(1)
																																				_ = Z5003
																																				__e.Return(MakeNative(func(__e *ControlFlow) {
																																					Z5004 := __e.Get(1)
																																					_ = Z5004
																																					tmp13501 := Call(__e, PrimFunc(symshen_4incinfs))

																																					_ = tmp13501

																																					tmp13502 := MakeNative(func(__e *ControlFlow) {
																																						__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4999, Z5004, V4933, V4934, V4935, W4938, V4937)
																																						return
																																					}, 0)

																																					__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4997, Z5003, V4933, V4934, V4935, W4938, tmp13502)
																																					return

																																				}, 1))
																																				return
																																			}, 1)

																																			__e.TailApply(tmp13432, tmp13500)
																																			return

																																		}, 1)

																																		tmp13503 := Call(__e, PrimFunc(symshen_4lazyderef), V4932, V4934)

																																		__e.TailApply(tmp13431, tmp13503)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp13506 := PrimTail(W4998)

																																tmp13507 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13506, V4934)

																																__e.TailApply(tmp13430, tmp13507)
																																return

																															}, 1)

																															tmp13508 := PrimHead(W4998)

																															__e.TailApply(tmp13429, tmp13508)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp13511 := PrimTail(W4996)

																													tmp13512 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13511, V4934)

																													__e.TailApply(tmp13428, tmp13512)
																													return

																												}, 1)

																												tmp13513 := PrimHead(W4996)

																												__e.TailApply(tmp13427, tmp13513)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp13516 := PrimTail(W4994)

																										tmp13517 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13516, V4934)

																										__e.TailApply(tmp13426, tmp13517)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13520 := PrimHead(W4994)

																								tmp13521 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13520, V4934)

																								__e.TailApply(tmp13425, tmp13521)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13524 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																						tmp13525 := Call(__e, tmp13424, tmp13524)

																						ifres13423 = tmp13525

																					} else {
																						ifres13423 = False

																					}

																					__e.TailApply(tmp12869, ifres13423)
																					return

																				} else {
																					__e.Return(W4972)
																					return
																				}

																			}, 1)

																			tmp13611 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																			var ifres13529 Obj

																			if True == tmp13611 {
																				tmp13530 := MakeNative(func(__e *ControlFlow) {
																					W4973 := __e.Get(1)
																					_ = W4973
																					tmp13608 := PrimIsPair(W4973)

																					if True == tmp13608 {
																						tmp13531 := MakeNative(func(__e *ControlFlow) {
																							W4974 := __e.Get(1)
																							_ = W4974
																							tmp13604 := PrimEqual(W4974, symcons)

																							if True == tmp13604 {
																								tmp13532 := MakeNative(func(__e *ControlFlow) {
																									W4975 := __e.Get(1)
																									_ = W4975
																									tmp13600 := PrimIsPair(W4975)

																									if True == tmp13600 {
																										tmp13533 := MakeNative(func(__e *ControlFlow) {
																											W4976 := __e.Get(1)
																											_ = W4976
																											tmp13534 := MakeNative(func(__e *ControlFlow) {
																												W4977 := __e.Get(1)
																												_ = W4977
																												tmp13595 := PrimIsPair(W4977)

																												if True == tmp13595 {
																													tmp13535 := MakeNative(func(__e *ControlFlow) {
																														W4978 := __e.Get(1)
																														_ = W4978
																														tmp13536 := MakeNative(func(__e *ControlFlow) {
																															W4979 := __e.Get(1)
																															_ = W4979
																															tmp13590 := PrimEqual(W4979, Nil)

																															if True == tmp13590 {
																																tmp13537 := MakeNative(func(__e *ControlFlow) {
																																	W4980 := __e.Get(1)
																																	_ = W4980
																																	tmp13538 := MakeNative(func(__e *ControlFlow) {
																																		W4981 := __e.Get(1)
																																		_ = W4981
																																		tmp13582 := PrimIsPair(W4980)

																																		if True == tmp13582 {
																																			tmp13539 := MakeNative(func(__e *ControlFlow) {
																																				W4983 := __e.Get(1)
																																				_ = W4983
																																				tmp13540 := MakeNative(func(__e *ControlFlow) {
																																					W4984 := __e.Get(1)
																																					_ = W4984
																																					tmp13544 := PrimEqual(W4983, symlist)

																																					if True == tmp13544 {
																																						__e.TailApply(PrimFunc(symthaw), W4984)
																																						return
																																					} else {
																																						tmp13542 := Call(__e, PrimFunc(symshen_4pvar_2), W4983)

																																						if True == tmp13542 {
																																							__e.TailApply(PrimFunc(symshen_4bind_b), W4983, symlist, V4934, W4984)
																																							return
																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp13545 := MakeNative(func(__e *ControlFlow) {
																																					tmp13546 := MakeNative(func(__e *ControlFlow) {
																																						W4985 := __e.Get(1)
																																						_ = W4985
																																						tmp13547 := MakeNative(func(__e *ControlFlow) {
																																							W4986 := __e.Get(1)
																																							_ = W4986
																																							tmp13567 := PrimIsPair(W4985)

																																							if True == tmp13567 {
																																								tmp13548 := MakeNative(func(__e *ControlFlow) {
																																									W4988 := __e.Get(1)
																																									_ = W4988
																																									tmp13549 := MakeNative(func(__e *ControlFlow) {
																																										W4989 := __e.Get(1)
																																										_ = W4989
																																										tmp13550 := MakeNative(func(__e *ControlFlow) {
																																											W4990 := __e.Get(1)
																																											_ = W4990
																																											tmp13554 := PrimEqual(W4989, Nil)

																																											if True == tmp13554 {
																																												__e.TailApply(PrimFunc(symthaw), W4990)
																																												return
																																											} else {
																																												tmp13552 := Call(__e, PrimFunc(symshen_4pvar_2), W4989)

																																												if True == tmp13552 {
																																													__e.TailApply(PrimFunc(symshen_4bind_b), W4989, Nil, V4934, W4990)
																																													return
																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}

																																										}, 1)

																																										tmp13555 := MakeNative(func(__e *ControlFlow) {
																																											__e.TailApply(W4986, W4988)
																																											return
																																										}, 0)

																																										__e.TailApply(tmp13550, tmp13555)
																																										return

																																									}, 1)

																																									tmp13556 := PrimTail(W4985)

																																									tmp13557 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13556, V4934)

																																									__e.TailApply(tmp13549, tmp13557)
																																									return

																																								}, 1)

																																								tmp13558 := PrimHead(W4985)

																																								__e.TailApply(tmp13548, tmp13558)
																																								return

																																							} else {
																																								tmp13565 := Call(__e, PrimFunc(symshen_4pvar_2), W4985)

																																								if True == tmp13565 {
																																									tmp13559 := MakeNative(func(__e *ControlFlow) {
																																										W4991 := __e.Get(1)
																																										_ = W4991
																																										tmp13560 := PrimCons(W4991, Nil)

																																										tmp13561 := MakeNative(func(__e *ControlFlow) {
																																											__e.TailApply(W4986, W4991)
																																											return
																																										}, 0)

																																										tmp13562 := Call(__e, PrimFunc(symshen_4bind_b), W4985, tmp13560, V4934, tmp13561)

																																										__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13562)
																																										return

																																									}, 1)

																																									tmp13563 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																									__e.TailApply(tmp13559, tmp13563)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp13568 := MakeNative(func(__e *ControlFlow) {
																																							Z4987 := __e.Get(1)
																																							_ = Z4987
																																							__e.TailApply(W4981, Z4987)
																																							return
																																						}, 1)

																																						__e.TailApply(tmp13547, tmp13568)
																																						return

																																					}, 1)

																																					tmp13569 := PrimTail(W4980)

																																					tmp13570 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13569, V4934)

																																					__e.TailApply(tmp13546, tmp13570)
																																					return

																																				}, 0)

																																				__e.TailApply(tmp13540, tmp13545)
																																				return

																																			}, 1)

																																			tmp13571 := PrimHead(W4980)

																																			tmp13572 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13571, V4934)

																																			__e.TailApply(tmp13539, tmp13572)
																																			return

																																		} else {
																																			tmp13580 := Call(__e, PrimFunc(symshen_4pvar_2), W4980)

																																			if True == tmp13580 {
																																				tmp13573 := MakeNative(func(__e *ControlFlow) {
																																					W4992 := __e.Get(1)
																																					_ = W4992
																																					tmp13574 := PrimCons(W4992, Nil)

																																					tmp13575 := PrimCons(symlist, tmp13574)

																																					tmp13576 := MakeNative(func(__e *ControlFlow) {
																																						__e.TailApply(W4981, W4992)
																																						return
																																					}, 0)

																																					tmp13577 := Call(__e, PrimFunc(symshen_4bind_b), W4980, tmp13575, V4934, tmp13576)

																																					__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13577)
																																					return

																																				}, 1)

																																				tmp13578 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																																				__e.TailApply(tmp13573, tmp13578)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp13583 := MakeNative(func(__e *ControlFlow) {
																																		Z4982 := __e.Get(1)
																																		_ = Z4982
																																		tmp13584 := Call(__e, PrimFunc(symshen_4incinfs))

																																		_ = tmp13584

																																		tmp13585 := MakeNative(func(__e *ControlFlow) {
																																			tmp13586 := PrimCons(Z4982, Nil)

																																			tmp13587 := PrimCons(symlist, tmp13586)

																																			__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4978, tmp13587, V4933, V4934, V4935, W4938, V4937)
																																			return

																																		}, 0)

																																		__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4976, Z4982, V4933, V4934, V4935, W4938, tmp13585)
																																		return

																																	}, 1)

																																	__e.TailApply(tmp13538, tmp13583)
																																	return

																																}, 1)

																																tmp13588 := Call(__e, PrimFunc(symshen_4lazyderef), V4932, V4934)

																																__e.TailApply(tmp13537, tmp13588)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp13591 := PrimTail(W4977)

																														tmp13592 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13591, V4934)

																														__e.TailApply(tmp13536, tmp13592)
																														return

																													}, 1)

																													tmp13593 := PrimHead(W4977)

																													__e.TailApply(tmp13535, tmp13593)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp13596 := PrimTail(W4975)

																											tmp13597 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13596, V4934)

																											__e.TailApply(tmp13534, tmp13597)
																											return

																										}, 1)

																										tmp13598 := PrimHead(W4975)

																										__e.TailApply(tmp13533, tmp13598)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13601 := PrimTail(W4973)

																								tmp13602 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13601, V4934)

																								__e.TailApply(tmp13532, tmp13602)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13605 := PrimHead(W4973)

																						tmp13606 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13605, V4934)

																						__e.TailApply(tmp13531, tmp13606)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp13609 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																				tmp13610 := Call(__e, tmp13530, tmp13609)

																				ifres13529 = tmp13610

																			} else {
																				ifres13529 = False

																			}

																			__e.TailApply(tmp12868, ifres13529)
																			return

																		} else {
																			__e.Return(W4965)
																			return
																		}

																	}, 1)

																	tmp13642 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

																	var ifres13614 Obj

																	if True == tmp13642 {
																		tmp13615 := MakeNative(func(__e *ControlFlow) {
																			W4966 := __e.Get(1)
																			_ = W4966
																			tmp13639 := PrimIsPair(W4966)

																			if True == tmp13639 {
																				tmp13616 := MakeNative(func(__e *ControlFlow) {
																					W4967 := __e.Get(1)
																					_ = W4967
																					tmp13617 := MakeNative(func(__e *ControlFlow) {
																						W4968 := __e.Get(1)
																						_ = W4968
																						tmp13634 := PrimIsPair(W4968)

																						if True == tmp13634 {
																							tmp13618 := MakeNative(func(__e *ControlFlow) {
																								W4969 := __e.Get(1)
																								_ = W4969
																								tmp13619 := MakeNative(func(__e *ControlFlow) {
																									W4970 := __e.Get(1)
																									_ = W4970
																									tmp13629 := PrimEqual(W4970, Nil)

																									if True == tmp13629 {
																										tmp13620 := MakeNative(func(__e *ControlFlow) {
																											W4971 := __e.Get(1)
																											_ = W4971
																											tmp13621 := Call(__e, PrimFunc(symshen_4incinfs))

																											_ = tmp13621

																											tmp13622 := PrimCons(V4932, Nil)

																											tmp13623 := PrimCons(sym_1_1_6, tmp13622)

																											tmp13624 := PrimCons(W4971, tmp13623)

																											tmp13625 := MakeNative(func(__e *ControlFlow) {
																												__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4969, W4971, V4933, V4934, V4935, W4938, V4937)
																												return
																											}, 0)

																											tmp13626 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4967, tmp13624, V4933, V4934, V4935, W4938, tmp13625)

																											__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13626)
																											return

																										}, 1)

																										tmp13627 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																										__e.TailApply(tmp13620, tmp13627)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp13630 := PrimTail(W4968)

																								tmp13631 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13630, V4934)

																								__e.TailApply(tmp13619, tmp13631)
																								return

																							}, 1)

																							tmp13632 := PrimHead(W4968)

																							__e.TailApply(tmp13618, tmp13632)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp13635 := PrimTail(W4966)

																					tmp13636 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13635, V4934)

																					__e.TailApply(tmp13617, tmp13636)
																					return

																				}, 1)

																				tmp13637 := PrimHead(W4966)

																				__e.TailApply(tmp13616, tmp13637)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp13640 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																		tmp13641 := Call(__e, tmp13615, tmp13640)

																		ifres13614 = tmp13641

																	} else {
																		ifres13614 = False

																	}

																	__e.TailApply(tmp12867, ifres13614)
																	return

																} else {
																	__e.Return(W4958)
																	return
																}

															}, 1)

															tmp13677 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

															var ifres13645 Obj

															if True == tmp13677 {
																tmp13646 := MakeNative(func(__e *ControlFlow) {
																	W4959 := __e.Get(1)
																	_ = W4959
																	tmp13674 := PrimIsPair(W4959)

																	if True == tmp13674 {
																		tmp13647 := MakeNative(func(__e *ControlFlow) {
																			W4960 := __e.Get(1)
																			_ = W4960
																			tmp13648 := MakeNative(func(__e *ControlFlow) {
																				W4961 := __e.Get(1)
																				_ = W4961
																				tmp13669 := PrimIsPair(W4961)

																				if True == tmp13669 {
																					tmp13649 := MakeNative(func(__e *ControlFlow) {
																						W4962 := __e.Get(1)
																						_ = W4962
																						tmp13650 := MakeNative(func(__e *ControlFlow) {
																							W4963 := __e.Get(1)
																							_ = W4963
																							tmp13664 := PrimEqual(W4963, Nil)

																							if True == tmp13664 {
																								tmp13651 := MakeNative(func(__e *ControlFlow) {
																									W4964 := __e.Get(1)
																									_ = W4964
																									tmp13652 := Call(__e, PrimFunc(symshen_4incinfs))

																									_ = tmp13652

																									tmp13653 := Call(__e, PrimFunc(symshen_4lazyderef), W4960, V4934)

																									tmp13654 := PrimIsPair(tmp13653)

																									tmp13655 := PrimNot(tmp13654)

																									tmp13656 := MakeNative(func(__e *ControlFlow) {
																										tmp13657 := PrimCons(V4932, Nil)

																										tmp13658 := PrimCons(sym_1_1_6, tmp13657)

																										tmp13659 := PrimCons(W4964, tmp13658)

																										tmp13660 := MakeNative(func(__e *ControlFlow) {
																											__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4962, W4964, V4933, V4934, V4935, W4938, V4937)
																											return
																										}, 0)

																										__e.TailApply(PrimFunc(symshen_4lookupsig), W4960, tmp13659, V4934, V4935, W4938, tmp13660)
																										return

																									}, 0)

																									tmp13661 := Call(__e, PrimFunc(symwhen), tmp13655, V4934, V4935, W4938, tmp13656)

																									__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp13661)
																									return

																								}, 1)

																								tmp13662 := Call(__e, PrimFunc(symshen_4newpv), V4934)

																								__e.TailApply(tmp13651, tmp13662)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp13665 := PrimTail(W4961)

																						tmp13666 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13665, V4934)

																						__e.TailApply(tmp13650, tmp13666)
																						return

																					}, 1)

																					tmp13667 := PrimHead(W4961)

																					__e.TailApply(tmp13649, tmp13667)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp13670 := PrimTail(W4959)

																			tmp13671 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13670, V4934)

																			__e.TailApply(tmp13648, tmp13671)
																			return

																		}, 1)

																		tmp13672 := PrimHead(W4959)

																		__e.TailApply(tmp13647, tmp13672)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp13675 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

																tmp13676 := Call(__e, tmp13646, tmp13675)

																ifres13645 = tmp13676

															} else {
																ifres13645 = False

															}

															__e.TailApply(tmp12866, ifres13645)
															return

														} else {
															__e.Return(W4952)
															return
														}

													}, 1)

													tmp13704 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

													var ifres13680 Obj

													if True == tmp13704 {
														tmp13681 := MakeNative(func(__e *ControlFlow) {
															W4953 := __e.Get(1)
															_ = W4953
															tmp13701 := PrimIsPair(W4953)

															if True == tmp13701 {
																tmp13682 := MakeNative(func(__e *ControlFlow) {
																	W4954 := __e.Get(1)
																	_ = W4954
																	tmp13697 := PrimEqual(W4954, symfn)

																	if True == tmp13697 {
																		tmp13683 := MakeNative(func(__e *ControlFlow) {
																			W4955 := __e.Get(1)
																			_ = W4955
																			tmp13693 := PrimIsPair(W4955)

																			if True == tmp13693 {
																				tmp13684 := MakeNative(func(__e *ControlFlow) {
																					W4956 := __e.Get(1)
																					_ = W4956
																					tmp13685 := MakeNative(func(__e *ControlFlow) {
																						W4957 := __e.Get(1)
																						_ = W4957
																						tmp13688 := PrimEqual(W4957, Nil)

																						if True == tmp13688 {
																							tmp13686 := Call(__e, PrimFunc(symshen_4incinfs))

																							_ = tmp13686

																							__e.TailApply(PrimFunc(symshen_4lookupsig), W4956, V4932, V4934, V4935, W4938, V4937)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp13689 := PrimTail(W4955)

																					tmp13690 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13689, V4934)

																					__e.TailApply(tmp13685, tmp13690)
																					return

																				}, 1)

																				tmp13691 := PrimHead(W4955)

																				__e.TailApply(tmp13684, tmp13691)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp13694 := PrimTail(W4953)

																		tmp13695 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13694, V4934)

																		__e.TailApply(tmp13683, tmp13695)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp13698 := PrimHead(W4953)

																tmp13699 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13698, V4934)

																__e.TailApply(tmp13682, tmp13699)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp13702 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

														tmp13703 := Call(__e, tmp13681, tmp13702)

														ifres13680 = tmp13703

													} else {
														ifres13680 = False

													}

													__e.TailApply(tmp12865, ifres13680)
													return

												} else {
													__e.Return(W4946)
													return
												}

											}, 1)

											tmp13737 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

											var ifres13707 Obj

											if True == tmp13737 {
												tmp13708 := MakeNative(func(__e *ControlFlow) {
													W4947 := __e.Get(1)
													_ = W4947
													tmp13734 := PrimIsPair(W4947)

													if True == tmp13734 {
														tmp13709 := MakeNative(func(__e *ControlFlow) {
															W4948 := __e.Get(1)
															_ = W4948
															tmp13730 := PrimEqual(W4948, symfn)

															if True == tmp13730 {
																tmp13710 := MakeNative(func(__e *ControlFlow) {
																	W4949 := __e.Get(1)
																	_ = W4949
																	tmp13726 := PrimIsPair(W4949)

																	if True == tmp13726 {
																		tmp13711 := MakeNative(func(__e *ControlFlow) {
																			W4950 := __e.Get(1)
																			_ = W4950
																			tmp13712 := MakeNative(func(__e *ControlFlow) {
																				W4951 := __e.Get(1)
																				_ = W4951
																				tmp13721 := PrimEqual(W4951, Nil)

																				if True == tmp13721 {
																					tmp13713 := Call(__e, PrimFunc(symshen_4incinfs))

																					_ = tmp13713

																					tmp13714 := Call(__e, PrimFunc(symshen_4deref), W4950, V4934)

																					tmp13715 := Call(__e, PrimFunc(symarity), tmp13714)

																					tmp13716 := PrimEqual(tmp13715, MakeNumber(0))

																					tmp13717 := MakeNative(func(__e *ControlFlow) {
																						tmp13718 := MakeNative(func(__e *ControlFlow) {
																							tmp13719 := PrimCons(W4950, Nil)

																							__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp13719, V4932, V4933, V4934, V4935, W4938, V4937)
																							return

																						}, 0)

																						__e.TailApply(PrimFunc(symshen_4cut), V4934, V4935, W4938, tmp13718)
																						return

																					}, 0)

																					__e.TailApply(PrimFunc(symwhen), tmp13716, V4934, V4935, W4938, tmp13717)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp13722 := PrimTail(W4949)

																			tmp13723 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13722, V4934)

																			__e.TailApply(tmp13712, tmp13723)
																			return

																		}, 1)

																		tmp13724 := PrimHead(W4949)

																		__e.TailApply(tmp13711, tmp13724)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp13727 := PrimTail(W4947)

																tmp13728 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13727, V4934)

																__e.TailApply(tmp13710, tmp13728)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp13731 := PrimHead(W4947)

														tmp13732 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13731, V4934)

														__e.TailApply(tmp13709, tmp13732)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp13735 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

												tmp13736 := Call(__e, tmp13708, tmp13735)

												ifres13707 = tmp13736

											} else {
												ifres13707 = False

											}

											__e.TailApply(tmp12864, ifres13707)
											return

										} else {
											__e.Return(W4942)
											return
										}

									}, 1)

									tmp13756 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

									var ifres13740 Obj

									if True == tmp13756 {
										tmp13741 := MakeNative(func(__e *ControlFlow) {
											W4943 := __e.Get(1)
											_ = W4943
											tmp13753 := PrimIsPair(W4943)

											if True == tmp13753 {
												tmp13742 := MakeNative(func(__e *ControlFlow) {
													W4944 := __e.Get(1)
													_ = W4944
													tmp13743 := MakeNative(func(__e *ControlFlow) {
														W4945 := __e.Get(1)
														_ = W4945
														tmp13748 := PrimEqual(W4945, Nil)

														if True == tmp13748 {
															tmp13744 := Call(__e, PrimFunc(symshen_4incinfs))

															_ = tmp13744

															tmp13745 := PrimCons(V4932, Nil)

															tmp13746 := PrimCons(sym_1_1_6, tmp13745)

															__e.TailApply(PrimFunc(symshen_4lookupsig), W4944, tmp13746, V4934, V4935, W4938, V4937)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp13749 := PrimTail(W4943)

													tmp13750 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13749, V4934)

													__e.TailApply(tmp13743, tmp13750)
													return

												}, 1)

												tmp13751 := PrimHead(W4943)

												__e.TailApply(tmp13742, tmp13751)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp13754 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

										tmp13755 := Call(__e, tmp13741, tmp13754)

										ifres13740 = tmp13755

									} else {
										ifres13740 = False

									}

									__e.TailApply(tmp12863, ifres13740)
									return

								} else {
									__e.Return(W4941)
									return
								}

							}, 1)

							tmp13762 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

							var ifres13759 Obj

							if True == tmp13762 {
								tmp13760 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp13760

								tmp13761 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4931, V4932, V4933, V4934, V4935, W4938, V4937)

								ifres13759 = tmp13761

							} else {
								ifres13759 = False

							}

							__e.TailApply(tmp12862, ifres13759)
							return

						} else {
							__e.Return(W4940)
							return
						}

					}, 1)

					tmp13772 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

					var ifres13765 Obj

					if True == tmp13772 {
						tmp13766 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp13766

						tmp13767 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)

						tmp13768 := PrimIsPair(tmp13767)

						tmp13769 := PrimNot(tmp13768)

						tmp13770 := MakeNative(func(__e *ControlFlow) {
							__e.TailApply(PrimFunc(symshen_4primitive), V4931, V4932, V4934, V4935, W4938, V4937)
							return
						}, 0)

						tmp13771 := Call(__e, PrimFunc(symwhen), tmp13769, V4934, V4935, W4938, tmp13770)

						ifres13765 = tmp13771

					} else {
						ifres13765 = False

					}

					__e.TailApply(tmp12861, ifres13765)
					return

				} else {
					__e.Return(W4939)
					return
				}

			}, 1)

			tmp13784 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)

			var ifres13775 Obj

			if True == tmp13784 {
				tmp13776 := Call(__e, PrimFunc(symshen_4incinfs))

				_ = tmp13776

				tmp13777 := PrimValue(symshen_4_dspy_d)

				tmp13778 := MakeNative(func(__e *ControlFlow) {
					tmp13779 := PrimIntern(MakeString(":"))

					tmp13780 := PrimCons(V4932, Nil)

					tmp13781 := PrimCons(tmp13779, tmp13780)

					tmp13782 := PrimCons(V4931, tmp13781)

					__e.TailApply(PrimFunc(symshen_4show), tmp13782, V4933, V4934, V4935, W4938, V4937)
					return

				}, 0)

				tmp13783 := Call(__e, PrimFunc(symwhen), tmp13777, V4934, V4935, W4938, tmp13778)

				ifres13775 = tmp13783

			} else {
				ifres13775 = False

			}

			__e.TailApply(tmp12860, ifres13775)
			return

		}, 1)

		tmp13785 := PrimNumberAdd(V4936, MakeNumber(1))

		__e.TailApply(tmp12859, tmp13785)
		return

	}, 7)

	tmp13786 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp12858)

	_ = tmp13786

	tmp13787 := MakeNative(func(__e *ControlFlow) {
		V5143 := __e.Get(1)
		_ = V5143
		V5144 := __e.Get(2)
		_ = V5144
		V5145 := __e.Get(3)
		_ = V5145
		V5146 := __e.Get(4)
		_ = V5146
		V5147 := __e.Get(5)
		_ = V5147
		V5148 := __e.Get(6)
		_ = V5148
		tmp13788 := MakeNative(func(__e *ControlFlow) {
			W5149 := __e.Get(1)
			_ = W5149
			tmp13896 := PrimEqual(W5149, False)

			if True == tmp13896 {
				tmp13789 := MakeNative(func(__e *ControlFlow) {
					W5152 := __e.Get(1)
					_ = W5152
					tmp13880 := PrimEqual(W5152, False)

					if True == tmp13880 {
						tmp13790 := MakeNative(func(__e *ControlFlow) {
							W5155 := __e.Get(1)
							_ = W5155
							tmp13864 := PrimEqual(W5155, False)

							if True == tmp13864 {
								tmp13791 := MakeNative(func(__e *ControlFlow) {
									W5158 := __e.Get(1)
									_ = W5158
									tmp13848 := PrimEqual(W5158, False)

									if True == tmp13848 {
										tmp13846 := Call(__e, PrimFunc(symshen_4unlocked_2), V5146)

										if True == tmp13846 {
											tmp13792 := MakeNative(func(__e *ControlFlow) {
												W5161 := __e.Get(1)
												_ = W5161
												tmp13843 := PrimEqual(W5161, Nil)

												if True == tmp13843 {
													tmp13793 := MakeNative(func(__e *ControlFlow) {
														W5162 := __e.Get(1)
														_ = W5162
														tmp13794 := MakeNative(func(__e *ControlFlow) {
															W5163 := __e.Get(1)
															_ = W5163
															tmp13838 := PrimIsPair(W5162)

															if True == tmp13838 {
																tmp13795 := MakeNative(func(__e *ControlFlow) {
																	W5165 := __e.Get(1)
																	_ = W5165
																	tmp13796 := MakeNative(func(__e *ControlFlow) {
																		W5166 := __e.Get(1)
																		_ = W5166
																		tmp13800 := PrimEqual(W5165, symlist)

																		if True == tmp13800 {
																			__e.TailApply(PrimFunc(symthaw), W5166)
																			return
																		} else {
																			tmp13798 := Call(__e, PrimFunc(symshen_4pvar_2), W5165)

																			if True == tmp13798 {
																				__e.TailApply(PrimFunc(symshen_4bind_b), W5165, symlist, V5145, W5166)
																				return
																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp13801 := MakeNative(func(__e *ControlFlow) {
																		tmp13802 := MakeNative(func(__e *ControlFlow) {
																			W5167 := __e.Get(1)
																			_ = W5167
																			tmp13803 := MakeNative(func(__e *ControlFlow) {
																				W5168 := __e.Get(1)
																				_ = W5168
																				tmp13823 := PrimIsPair(W5167)

																				if True == tmp13823 {
																					tmp13804 := MakeNative(func(__e *ControlFlow) {
																						W5170 := __e.Get(1)
																						_ = W5170
																						tmp13805 := MakeNative(func(__e *ControlFlow) {
																							W5171 := __e.Get(1)
																							_ = W5171
																							tmp13806 := MakeNative(func(__e *ControlFlow) {
																								W5172 := __e.Get(1)
																								_ = W5172
																								tmp13810 := PrimEqual(W5171, Nil)

																								if True == tmp13810 {
																									__e.TailApply(PrimFunc(symthaw), W5172)
																									return
																								} else {
																									tmp13808 := Call(__e, PrimFunc(symshen_4pvar_2), W5171)

																									if True == tmp13808 {
																										__e.TailApply(PrimFunc(symshen_4bind_b), W5171, Nil, V5145, W5172)
																										return
																									} else {
																										__e.Return(False)
																										return
																									}

																								}

																							}, 1)

																							tmp13811 := MakeNative(func(__e *ControlFlow) {
																								__e.TailApply(W5168, W5170)
																								return
																							}, 0)

																							__e.TailApply(tmp13806, tmp13811)
																							return

																						}, 1)

																						tmp13812 := PrimTail(W5167)

																						tmp13813 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13812, V5145)

																						__e.TailApply(tmp13805, tmp13813)
																						return

																					}, 1)

																					tmp13814 := PrimHead(W5167)

																					__e.TailApply(tmp13804, tmp13814)
																					return

																				} else {
																					tmp13821 := Call(__e, PrimFunc(symshen_4pvar_2), W5167)

																					if True == tmp13821 {
																						tmp13815 := MakeNative(func(__e *ControlFlow) {
																							W5173 := __e.Get(1)
																							_ = W5173
																							tmp13816 := PrimCons(W5173, Nil)

																							tmp13817 := MakeNative(func(__e *ControlFlow) {
																								__e.TailApply(W5168, W5173)
																								return
																							}, 0)

																							tmp13818 := Call(__e, PrimFunc(symshen_4bind_b), W5167, tmp13816, V5145, tmp13817)

																							__e.TailApply(PrimFunc(symshen_4gc), V5145, tmp13818)
																							return

																						}, 1)

																						tmp13819 := Call(__e, PrimFunc(symshen_4newpv), V5145)

																						__e.TailApply(tmp13815, tmp13819)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}

																			}, 1)

																			tmp13824 := MakeNative(func(__e *ControlFlow) {
																				Z5169 := __e.Get(1)
																				_ = Z5169
																				__e.TailApply(W5163, Z5169)
																				return
																			}, 1)

																			__e.TailApply(tmp13803, tmp13824)
																			return

																		}, 1)

																		tmp13825 := PrimTail(W5162)

																		tmp13826 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13825, V5145)

																		__e.TailApply(tmp13802, tmp13826)
																		return

																	}, 0)

																	__e.TailApply(tmp13796, tmp13801)
																	return

																}, 1)

																tmp13827 := PrimHead(W5162)

																tmp13828 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13827, V5145)

																__e.TailApply(tmp13795, tmp13828)
																return

															} else {
																tmp13836 := Call(__e, PrimFunc(symshen_4pvar_2), W5162)

																if True == tmp13836 {
																	tmp13829 := MakeNative(func(__e *ControlFlow) {
																		W5174 := __e.Get(1)
																		_ = W5174
																		tmp13830 := PrimCons(W5174, Nil)

																		tmp13831 := PrimCons(symlist, tmp13830)

																		tmp13832 := MakeNative(func(__e *ControlFlow) {
																			__e.TailApply(W5163, W5174)
																			return
																		}, 0)

																		tmp13833 := Call(__e, PrimFunc(symshen_4bind_b), W5162, tmp13831, V5145, tmp13832)

																		__e.TailApply(PrimFunc(symshen_4gc), V5145, tmp13833)
																		return

																	}, 1)

																	tmp13834 := Call(__e, PrimFunc(symshen_4newpv), V5145)

																	__e.TailApply(tmp13829, tmp13834)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}

														}, 1)

														tmp13839 := MakeNative(func(__e *ControlFlow) {
															Z5164 := __e.Get(1)
															_ = Z5164
															tmp13840 := Call(__e, PrimFunc(symshen_4incinfs))

															_ = tmp13840

															__e.TailApply(PrimFunc(symthaw), V5148)
															return

														}, 1)

														__e.TailApply(tmp13794, tmp13839)
														return

													}, 1)

													tmp13841 := Call(__e, PrimFunc(symshen_4lazyderef), V5144, V5145)

													__e.TailApply(tmp13793, tmp13841)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp13844 := Call(__e, PrimFunc(symshen_4lazyderef), V5143, V5145)

											__e.TailApply(tmp13792, tmp13844)
											return

										} else {
											__e.Return(False)
											return
										}

									} else {
										__e.Return(W5158)
										return
									}

								}, 1)

								tmp13862 := Call(__e, PrimFunc(symshen_4unlocked_2), V5146)

								var ifres13849 Obj

								if True == tmp13862 {
									tmp13850 := MakeNative(func(__e *ControlFlow) {
										W5159 := __e.Get(1)
										_ = W5159
										tmp13851 := MakeNative(func(__e *ControlFlow) {
											W5160 := __e.Get(1)
											_ = W5160
											tmp13855 := PrimEqual(W5159, symsymbol)

											if True == tmp13855 {
												__e.TailApply(PrimFunc(symthaw), W5160)
												return
											} else {
												tmp13853 := Call(__e, PrimFunc(symshen_4pvar_2), W5159)

												if True == tmp13853 {
													__e.TailApply(PrimFunc(symshen_4bind_b), W5159, symsymbol, V5145, W5160)
													return
												} else {
													__e.Return(False)
													return
												}

											}

										}, 1)

										tmp13856 := MakeNative(func(__e *ControlFlow) {
											tmp13857 := Call(__e, PrimFunc(symshen_4incinfs))

											_ = tmp13857

											tmp13858 := Call(__e, PrimFunc(symshen_4lazyderef), V5143, V5145)

											tmp13859 := PrimIsSymbol(tmp13858)

											__e.TailApply(PrimFunc(symwhen), tmp13859, V5145, V5146, V5147, V5148)
											return

										}, 0)

										__e.TailApply(tmp13851, tmp13856)
										return

									}, 1)

									tmp13860 := Call(__e, PrimFunc(symshen_4lazyderef), V5144, V5145)

									tmp13861 := Call(__e, tmp13850, tmp13860)

									ifres13849 = tmp13861

								} else {
									ifres13849 = False

								}

								__e.TailApply(tmp13791, ifres13849)
								return

							} else {
								__e.Return(W5155)
								return
							}

						}, 1)

						tmp13878 := Call(__e, PrimFunc(symshen_4unlocked_2), V5146)

						var ifres13865 Obj

						if True == tmp13878 {
							tmp13866 := MakeNative(func(__e *ControlFlow) {
								W5156 := __e.Get(1)
								_ = W5156
								tmp13867 := MakeNative(func(__e *ControlFlow) {
									W5157 := __e.Get(1)
									_ = W5157
									tmp13871 := PrimEqual(W5156, symstring)

									if True == tmp13871 {
										__e.TailApply(PrimFunc(symthaw), W5157)
										return
									} else {
										tmp13869 := Call(__e, PrimFunc(symshen_4pvar_2), W5156)

										if True == tmp13869 {
											__e.TailApply(PrimFunc(symshen_4bind_b), W5156, symstring, V5145, W5157)
											return
										} else {
											__e.Return(False)
											return
										}

									}

								}, 1)

								tmp13872 := MakeNative(func(__e *ControlFlow) {
									tmp13873 := Call(__e, PrimFunc(symshen_4incinfs))

									_ = tmp13873

									tmp13874 := Call(__e, PrimFunc(symshen_4lazyderef), V5143, V5145)

									tmp13875 := PrimIsString(tmp13874)

									__e.TailApply(PrimFunc(symwhen), tmp13875, V5145, V5146, V5147, V5148)
									return

								}, 0)

								__e.TailApply(tmp13867, tmp13872)
								return

							}, 1)

							tmp13876 := Call(__e, PrimFunc(symshen_4lazyderef), V5144, V5145)

							tmp13877 := Call(__e, tmp13866, tmp13876)

							ifres13865 = tmp13877

						} else {
							ifres13865 = False

						}

						__e.TailApply(tmp13790, ifres13865)
						return

					} else {
						__e.Return(W5152)
						return
					}

				}, 1)

				tmp13894 := Call(__e, PrimFunc(symshen_4unlocked_2), V5146)

				var ifres13881 Obj

				if True == tmp13894 {
					tmp13882 := MakeNative(func(__e *ControlFlow) {
						W5153 := __e.Get(1)
						_ = W5153
						tmp13883 := MakeNative(func(__e *ControlFlow) {
							W5154 := __e.Get(1)
							_ = W5154
							tmp13887 := PrimEqual(W5153, symboolean)

							if True == tmp13887 {
								__e.TailApply(PrimFunc(symthaw), W5154)
								return
							} else {
								tmp13885 := Call(__e, PrimFunc(symshen_4pvar_2), W5153)

								if True == tmp13885 {
									__e.TailApply(PrimFunc(symshen_4bind_b), W5153, symboolean, V5145, W5154)
									return
								} else {
									__e.Return(False)
									return
								}

							}

						}, 1)

						tmp13888 := MakeNative(func(__e *ControlFlow) {
							tmp13889 := Call(__e, PrimFunc(symshen_4incinfs))

							_ = tmp13889

							tmp13890 := Call(__e, PrimFunc(symshen_4lazyderef), V5143, V5145)

							tmp13891 := Call(__e, PrimFunc(symboolean_2), tmp13890)

							__e.TailApply(PrimFunc(symwhen), tmp13891, V5145, V5146, V5147, V5148)
							return

						}, 0)

						__e.TailApply(tmp13883, tmp13888)
						return

					}, 1)

					tmp13892 := Call(__e, PrimFunc(symshen_4lazyderef), V5144, V5145)

					tmp13893 := Call(__e, tmp13882, tmp13892)

					ifres13881 = tmp13893

				} else {
					ifres13881 = False

				}

				__e.TailApply(tmp13789, ifres13881)
				return

			} else {
				__e.Return(W5149)
				return
			}

		}, 1)

		tmp13910 := Call(__e, PrimFunc(symshen_4unlocked_2), V5146)

		var ifres13897 Obj

		if True == tmp13910 {
			tmp13898 := MakeNative(func(__e *ControlFlow) {
				W5150 := __e.Get(1)
				_ = W5150
				tmp13899 := MakeNative(func(__e *ControlFlow) {
					W5151 := __e.Get(1)
					_ = W5151
					tmp13903 := PrimEqual(W5150, symnumber)

					if True == tmp13903 {
						__e.TailApply(PrimFunc(symthaw), W5151)
						return
					} else {
						tmp13901 := Call(__e, PrimFunc(symshen_4pvar_2), W5150)

						if True == tmp13901 {
							__e.TailApply(PrimFunc(symshen_4bind_b), W5150, symnumber, V5145, W5151)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				tmp13904 := MakeNative(func(__e *ControlFlow) {
					tmp13905 := Call(__e, PrimFunc(symshen_4incinfs))

					_ = tmp13905

					tmp13906 := Call(__e, PrimFunc(symshen_4lazyderef), V5143, V5145)

					tmp13907 := PrimIsNumber(tmp13906)

					__e.TailApply(PrimFunc(symwhen), tmp13907, V5145, V5146, V5147, V5148)
					return

				}, 0)

				__e.TailApply(tmp13899, tmp13904)
				return

			}, 1)

			tmp13908 := Call(__e, PrimFunc(symshen_4lazyderef), V5144, V5145)

			tmp13909 := Call(__e, tmp13898, tmp13908)

			ifres13897 = tmp13909

		} else {
			ifres13897 = False

		}

		__e.TailApply(tmp13788, ifres13897)
		return

	}, 6)

	tmp13911 := Call(__e, ns2_1set, symshen_4primitive, tmp13787)

	_ = tmp13911

	tmp13912 := MakeNative(func(__e *ControlFlow) {
		V5175 := __e.Get(1)
		_ = V5175
		V5176 := __e.Get(2)
		_ = V5176
		V5177 := __e.Get(3)
		_ = V5177
		V5178 := __e.Get(4)
		_ = V5178
		V5179 := __e.Get(5)
		_ = V5179
		V5180 := __e.Get(6)
		_ = V5180
		V5181 := __e.Get(7)
		_ = V5181
		tmp13913 := MakeNative(func(__e *ControlFlow) {
			W5182 := __e.Get(1)
			_ = W5182
			tmp13924 := PrimEqual(W5182, False)

			if True == tmp13924 {
				tmp13922 := Call(__e, PrimFunc(symshen_4unlocked_2), V5179)

				if True == tmp13922 {
					tmp13914 := MakeNative(func(__e *ControlFlow) {
						W5191 := __e.Get(1)
						_ = W5191
						tmp13919 := PrimIsPair(W5191)

						if True == tmp13919 {
							tmp13915 := MakeNative(func(__e *ControlFlow) {
								W5192 := __e.Get(1)
								_ = W5192
								tmp13916 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp13916

								__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V5175, V5176, W5192, V5178, V5179, V5180, V5181)
								return

							}, 1)

							tmp13917 := PrimTail(W5191)

							__e.TailApply(tmp13915, tmp13917)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp13920 := Call(__e, PrimFunc(symshen_4lazyderef), V5177, V5178)

					__e.TailApply(tmp13914, tmp13920)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W5182)
				return
			}

		}, 1)

		tmp13966 := Call(__e, PrimFunc(symshen_4unlocked_2), V5179)

		var ifres13925 Obj

		if True == tmp13966 {
			tmp13926 := MakeNative(func(__e *ControlFlow) {
				W5183 := __e.Get(1)
				_ = W5183
				tmp13963 := PrimIsPair(W5183)

				if True == tmp13963 {
					tmp13927 := MakeNative(func(__e *ControlFlow) {
						W5184 := __e.Get(1)
						_ = W5184
						tmp13959 := PrimIsPair(W5184)

						if True == tmp13959 {
							tmp13928 := MakeNative(func(__e *ControlFlow) {
								W5185 := __e.Get(1)
								_ = W5185
								tmp13929 := MakeNative(func(__e *ControlFlow) {
									W5186 := __e.Get(1)
									_ = W5186
									tmp13954 := PrimIsPair(W5186)

									if True == tmp13954 {
										tmp13930 := MakeNative(func(__e *ControlFlow) {
											W5187 := __e.Get(1)
											_ = W5187
											tmp13931 := MakeNative(func(__e *ControlFlow) {
												W5188 := __e.Get(1)
												_ = W5188
												tmp13949 := PrimIsPair(W5188)

												if True == tmp13949 {
													tmp13932 := MakeNative(func(__e *ControlFlow) {
														W5189 := __e.Get(1)
														_ = W5189
														tmp13933 := MakeNative(func(__e *ControlFlow) {
															W5190 := __e.Get(1)
															_ = W5190
															tmp13944 := PrimEqual(W5190, Nil)

															if True == tmp13944 {
																tmp13934 := Call(__e, PrimFunc(symshen_4incinfs))

																_ = tmp13934

																tmp13935 := Call(__e, PrimFunc(symshen_4deref), W5187, V5178)

																tmp13936 := PrimIntern(MakeString(":"))

																tmp13937 := PrimEqual(tmp13935, tmp13936)

																tmp13938 := MakeNative(func(__e *ControlFlow) {
																	tmp13939 := Call(__e, PrimFunc(symshen_4deref), V5175, V5178)

																	tmp13940 := Call(__e, PrimFunc(symshen_4deref), W5185, V5178)

																	tmp13941 := PrimEqual(tmp13939, tmp13940)

																	tmp13942 := MakeNative(func(__e *ControlFlow) {
																		__e.TailApply(PrimFunc(symis_b), V5176, W5189, V5178, V5179, V5180, V5181)
																		return
																	}, 0)

																	__e.TailApply(PrimFunc(symwhen), tmp13941, V5178, V5179, V5180, tmp13942)
																	return

																}, 0)

																__e.TailApply(PrimFunc(symwhen), tmp13937, V5178, V5179, V5180, tmp13938)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp13945 := PrimTail(W5188)

														tmp13946 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13945, V5178)

														__e.TailApply(tmp13933, tmp13946)
														return

													}, 1)

													tmp13947 := PrimHead(W5188)

													__e.TailApply(tmp13932, tmp13947)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp13950 := PrimTail(W5186)

											tmp13951 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13950, V5178)

											__e.TailApply(tmp13931, tmp13951)
											return

										}, 1)

										tmp13952 := PrimHead(W5186)

										__e.TailApply(tmp13930, tmp13952)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp13955 := PrimTail(W5184)

								tmp13956 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13955, V5178)

								__e.TailApply(tmp13929, tmp13956)
								return

							}, 1)

							tmp13957 := PrimHead(W5184)

							__e.TailApply(tmp13928, tmp13957)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp13960 := PrimHead(W5183)

					tmp13961 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13960, V5178)

					__e.TailApply(tmp13927, tmp13961)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp13964 := Call(__e, PrimFunc(symshen_4lazyderef), V5177, V5178)

			tmp13965 := Call(__e, tmp13926, tmp13964)

			ifres13925 = tmp13965

		} else {
			ifres13925 = False

		}

		__e.TailApply(tmp13913, ifres13925)
		return

	}, 7)

	tmp13967 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp13912)

	_ = tmp13967

	tmp13968 := MakeNative(func(__e *ControlFlow) {
		V5193 := __e.Get(1)
		_ = V5193
		V5194 := __e.Get(2)
		_ = V5194
		V5195 := __e.Get(3)
		_ = V5195
		V5196 := __e.Get(4)
		_ = V5196
		V5197 := __e.Get(5)
		_ = V5197
		V5198 := __e.Get(6)
		_ = V5198
		tmp13973 := Call(__e, PrimFunc(symshen_4unlocked_2), V5196)

		if True == tmp13973 {
			tmp13969 := Call(__e, PrimFunc(symshen_4incinfs))

			_ = tmp13969

			tmp13970 := PrimValue(symshen_4_dsigf_d)

			tmp13971 := Call(__e, PrimFunc(symassoc), V5193, tmp13970)

			__e.TailApply(PrimFunc(symshen_4sigf), tmp13971, V5194, V5195, V5196, V5197, V5198)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 6)

	tmp13974 := Call(__e, ns2_1set, symshen_4lookupsig, tmp13968)

	_ = tmp13974

	tmp13975 := MakeNative(func(__e *ControlFlow) {
		V5213 := __e.Get(1)
		_ = V5213
		V5214 := __e.Get(2)
		_ = V5214
		V5215 := __e.Get(3)
		_ = V5215
		V5216 := __e.Get(4)
		_ = V5216
		V5217 := __e.Get(5)
		_ = V5217
		V5218 := __e.Get(6)
		_ = V5218
		tmp13982 := PrimIsPair(V5213)

		if True == tmp13982 {
			tmp13976 := PrimTail(V5213)

			tmp13977 := Call(__e, tmp13976, V5214)

			tmp13978 := Call(__e, tmp13977, V5215)

			tmp13979 := Call(__e, tmp13978, V5216)

			tmp13980 := Call(__e, tmp13979, V5217)

			__e.TailApply(tmp13980, V5218)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 6)

	tmp13983 := Call(__e, ns2_1set, symshen_4sigf, tmp13975)

	_ = tmp13983

	tmp13984 := MakeNative(func(__e *ControlFlow) {
		V5219 := __e.Get(1)
		_ = V5219
		tmp13985 := MakeNative(func(__e *ControlFlow) {
			W5220 := __e.Get(1)
			_ = W5220
			tmp13986 := MakeNative(func(__e *ControlFlow) {
				W5221 := __e.Get(1)
				_ = W5221
				tmp13987 := MakeNative(func(__e *ControlFlow) {
					W5222 := __e.Get(1)
					_ = W5222
					tmp13988 := MakeNative(func(__e *ControlFlow) {
						W5223 := __e.Get(1)
						_ = W5223
						__e.Return(W5223)
						return
					}, 1)

					tmp13989 := PrimValue(symshen_4_dgensym_d)

					tmp13990 := PrimNumberAdd(MakeNumber(1), tmp13989)

					tmp13991 := PrimSet(symshen_4_dgensym_d, tmp13990)

					tmp13992 := PrimVectorSet(W5222, MakeNumber(2), tmp13991)

					__e.TailApply(tmp13988, tmp13992)
					return

				}, 1)

				tmp13993 := PrimVectorSet(W5221, MakeNumber(1), V5219)

				__e.TailApply(tmp13987, tmp13993)
				return

			}, 1)

			tmp13994 := PrimVectorSet(W5220, MakeNumber(0), symshen_4print_1freshterm)

			__e.TailApply(tmp13986, tmp13994)
			return

		}, 1)

		tmp13995 := PrimAbsvector(MakeNumber(3))

		__e.TailApply(tmp13985, tmp13995)
		return

	}, 1)

	tmp13996 := Call(__e, ns2_1set, symshen_4freshterm, tmp13984)

	_ = tmp13996

	tmp13997 := MakeNative(func(__e *ControlFlow) {
		V5224 := __e.Get(1)
		_ = V5224
		tmp13998 := PrimVectorGet(V5224, MakeNumber(1))

		tmp13999 := PrimStr(tmp13998)

		__e.Return(PrimStringConcat(MakeString("&&"), tmp13999))
		return

	}, 1)

	tmp14000 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp13997)

	_ = tmp14000

	tmp14001 := MakeNative(func(__e *ControlFlow) {
		V5225 := __e.Get(1)
		_ = V5225
		V5226 := __e.Get(2)
		_ = V5226
		V5227 := __e.Get(3)
		_ = V5227
		V5228 := __e.Get(4)
		_ = V5228
		V5229 := __e.Get(5)
		_ = V5229
		V5230 := __e.Get(6)
		_ = V5230
		V5231 := __e.Get(7)
		_ = V5231
		tmp14002 := MakeNative(func(__e *ControlFlow) {
			W5232 := __e.Get(1)
			_ = W5232
			tmp14013 := PrimEqual(W5232, False)

			if True == tmp14013 {
				tmp14011 := Call(__e, PrimFunc(symshen_4unlocked_2), V5229)

				if True == tmp14011 {
					tmp14003 := MakeNative(func(__e *ControlFlow) {
						W5236 := __e.Get(1)
						_ = W5236
						tmp14008 := PrimIsPair(W5236)

						if True == tmp14008 {
							tmp14004 := MakeNative(func(__e *ControlFlow) {
								W5237 := __e.Get(1)
								_ = W5237
								tmp14005 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp14005

								__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V5225, V5226, W5237, V5228, V5229, V5230, V5231)
								return

							}, 1)

							tmp14006 := PrimTail(W5236)

							__e.TailApply(tmp14004, tmp14006)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp14009 := Call(__e, PrimFunc(symshen_4lazyderef), V5227, V5228)

					__e.TailApply(tmp14003, tmp14009)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W5232)
				return
			}

		}, 1)

		tmp14033 := Call(__e, PrimFunc(symshen_4unlocked_2), V5229)

		var ifres14014 Obj

		if True == tmp14033 {
			tmp14015 := MakeNative(func(__e *ControlFlow) {
				W5233 := __e.Get(1)
				_ = W5233
				tmp14030 := PrimIsPair(W5233)

				if True == tmp14030 {
					tmp14016 := MakeNative(func(__e *ControlFlow) {
						W5234 := __e.Get(1)
						_ = W5234
						tmp14026 := PrimIsPair(W5234)

						if True == tmp14026 {
							tmp14017 := MakeNative(func(__e *ControlFlow) {
								W5235 := __e.Get(1)
								_ = W5235
								tmp14018 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp14018

								tmp14019 := Call(__e, PrimFunc(symshen_4deref), W5235, V5228)

								tmp14020 := Call(__e, PrimFunc(symshen_4deref), V5225, V5228)

								tmp14021 := Call(__e, tmp14019, tmp14020)

								tmp14022 := Call(__e, PrimFunc(symshen_4deref), V5226, V5228)

								tmp14023 := Call(__e, tmp14021, tmp14022)

								__e.TailApply(PrimFunc(symcall), tmp14023, V5228, V5229, V5230, V5231)
								return

							}, 1)

							tmp14024 := PrimTail(W5234)

							__e.TailApply(tmp14017, tmp14024)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp14027 := PrimHead(W5233)

					tmp14028 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14027, V5228)

					__e.TailApply(tmp14016, tmp14028)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp14031 := Call(__e, PrimFunc(symshen_4lazyderef), V5227, V5228)

			tmp14032 := Call(__e, tmp14015, tmp14031)

			ifres14014 = tmp14032

		} else {
			ifres14014 = False

		}

		__e.TailApply(tmp14002, ifres14014)
		return

	}, 7)

	tmp14034 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp14001)

	_ = tmp14034

	tmp14035 := MakeNative(func(__e *ControlFlow) {
		V5238 := __e.Get(1)
		_ = V5238
		V5239 := __e.Get(2)
		_ = V5239
		V5240 := __e.Get(3)
		_ = V5240
		V5241 := __e.Get(4)
		_ = V5241
		V5242 := __e.Get(5)
		_ = V5242
		V5243 := __e.Get(6)
		_ = V5243
		V5244 := __e.Get(7)
		_ = V5244
		tmp14036 := MakeNative(func(__e *ControlFlow) {
			W5245 := __e.Get(1)
			_ = W5245
			tmp14037 := MakeNative(func(__e *ControlFlow) {
				W5246 := __e.Get(1)
				_ = W5246
				tmp14467 := PrimEqual(W5246, False)

				if True == tmp14467 {
					tmp14038 := MakeNative(func(__e *ControlFlow) {
						W5249 := __e.Get(1)
						_ = W5249
						tmp14367 := PrimEqual(W5249, False)

						if True == tmp14367 {
							tmp14039 := MakeNative(func(__e *ControlFlow) {
								W5269 := __e.Get(1)
								_ = W5269
								tmp14262 := PrimEqual(W5269, False)

								if True == tmp14262 {
									tmp14040 := MakeNative(func(__e *ControlFlow) {
										W5291 := __e.Get(1)
										_ = W5291
										tmp14181 := PrimEqual(W5291, False)

										if True == tmp14181 {
											tmp14041 := MakeNative(func(__e *ControlFlow) {
												W5307 := __e.Get(1)
												_ = W5307
												tmp14081 := PrimEqual(W5307, False)

												if True == tmp14081 {
													tmp14042 := MakeNative(func(__e *ControlFlow) {
														W5327 := __e.Get(1)
														_ = W5327
														tmp14044 := PrimEqual(W5327, False)

														if True == tmp14044 {
															__e.TailApply(PrimFunc(symshen_4unlock), V5242, W5245)
															return
														} else {
															__e.Return(W5327)
															return
														}

													}, 1)

													tmp14079 := Call(__e, PrimFunc(symshen_4unlocked_2), V5242)

													var ifres14045 Obj

													if True == tmp14079 {
														tmp14046 := MakeNative(func(__e *ControlFlow) {
															W5328 := __e.Get(1)
															_ = W5328
															tmp14076 := PrimIsPair(W5328)

															if True == tmp14076 {
																tmp14047 := MakeNative(func(__e *ControlFlow) {
																	W5329 := __e.Get(1)
																	_ = W5329
																	tmp14048 := MakeNative(func(__e *ControlFlow) {
																		W5330 := __e.Get(1)
																		_ = W5330
																		tmp14049 := MakeNative(func(__e *ControlFlow) {
																			W5331 := __e.Get(1)
																			_ = W5331
																			tmp14050 := MakeNative(func(__e *ControlFlow) {
																				W5332 := __e.Get(1)
																				_ = W5332
																				tmp14068 := PrimIsPair(W5331)

																				if True == tmp14068 {
																					tmp14051 := MakeNative(func(__e *ControlFlow) {
																						W5335 := __e.Get(1)
																						_ = W5335
																						tmp14052 := MakeNative(func(__e *ControlFlow) {
																							W5336 := __e.Get(1)
																							_ = W5336
																							tmp14053 := Call(__e, W5332, W5335)

																							__e.TailApply(tmp14053, W5336)
																							return

																						}, 1)

																						tmp14054 := PrimTail(W5331)

																						__e.TailApply(tmp14052, tmp14054)
																						return

																					}, 1)

																					tmp14055 := PrimHead(W5331)

																					__e.TailApply(tmp14051, tmp14055)
																					return

																				} else {
																					tmp14066 := Call(__e, PrimFunc(symshen_4pvar_2), W5331)

																					if True == tmp14066 {
																						tmp14056 := MakeNative(func(__e *ControlFlow) {
																							W5337 := __e.Get(1)
																							_ = W5337
																							tmp14057 := MakeNative(func(__e *ControlFlow) {
																								W5338 := __e.Get(1)
																								_ = W5338
																								tmp14058 := PrimCons(W5337, W5338)

																								tmp14059 := MakeNative(func(__e *ControlFlow) {
																									tmp14060 := Call(__e, W5332, W5337)

																									__e.TailApply(tmp14060, W5338)
																									return

																								}, 0)

																								tmp14061 := Call(__e, PrimFunc(symshen_4bind_b), W5331, tmp14058, V5241, tmp14059)

																								__e.TailApply(PrimFunc(symshen_4gc), V5241, tmp14061)
																								return

																							}, 1)

																							tmp14062 := Call(__e, PrimFunc(symshen_4newpv), V5241)

																							tmp14063 := Call(__e, tmp14057, tmp14062)

																							__e.TailApply(PrimFunc(symshen_4gc), V5241, tmp14063)
																							return

																						}, 1)

																						tmp14064 := Call(__e, PrimFunc(symshen_4newpv), V5241)

																						__e.TailApply(tmp14056, tmp14064)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}

																			}, 1)

																			tmp14069 := MakeNative(func(__e *ControlFlow) {
																				Z5333 := __e.Get(1)
																				_ = Z5333
																				__e.Return(MakeNative(func(__e *ControlFlow) {
																					Z5334 := __e.Get(1)
																					_ = Z5334
																					tmp14070 := Call(__e, PrimFunc(symshen_4incinfs))

																					_ = tmp14070

																					tmp14071 := MakeNative(func(__e *ControlFlow) {
																						__e.TailApply(PrimFunc(symshen_4l_1rules), W5330, Z5334, V5240, V5241, V5242, W5245, V5244)
																						return
																					}, 0)

																					__e.TailApply(PrimFunc(symbind), Z5333, W5329, V5241, V5242, W5245, tmp14071)
																					return

																				}, 1))
																				return
																			}, 1)

																			__e.TailApply(tmp14050, tmp14069)
																			return

																		}, 1)

																		tmp14072 := Call(__e, PrimFunc(symshen_4lazyderef), V5239, V5241)

																		__e.TailApply(tmp14049, tmp14072)
																		return

																	}, 1)

																	tmp14073 := PrimTail(W5328)

																	__e.TailApply(tmp14048, tmp14073)
																	return

																}, 1)

																tmp14074 := PrimHead(W5328)

																__e.TailApply(tmp14047, tmp14074)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14077 := Call(__e, PrimFunc(symshen_4lazyderef), V5238, V5241)

														tmp14078 := Call(__e, tmp14046, tmp14077)

														ifres14045 = tmp14078

													} else {
														ifres14045 = False

													}

													__e.TailApply(tmp14042, ifres14045)
													return

												} else {
													__e.Return(W5307)
													return
												}

											}, 1)

											tmp14179 := Call(__e, PrimFunc(symshen_4unlocked_2), V5242)

											var ifres14082 Obj

											if True == tmp14179 {
												tmp14083 := MakeNative(func(__e *ControlFlow) {
													W5308 := __e.Get(1)
													_ = W5308
													tmp14176 := PrimIsPair(W5308)

													if True == tmp14176 {
														tmp14084 := MakeNative(func(__e *ControlFlow) {
															W5309 := __e.Get(1)
															_ = W5309
															tmp14172 := PrimIsPair(W5309)

															if True == tmp14172 {
																tmp14085 := MakeNative(func(__e *ControlFlow) {
																	W5310 := __e.Get(1)
																	_ = W5310
																	tmp14168 := PrimIsPair(W5310)

																	if True == tmp14168 {
																		tmp14086 := MakeNative(func(__e *ControlFlow) {
																			W5311 := __e.Get(1)
																			_ = W5311
																			tmp14164 := PrimEqual(W5311, sym_8v)

																			if True == tmp14164 {
																				tmp14087 := MakeNative(func(__e *ControlFlow) {
																					W5312 := __e.Get(1)
																					_ = W5312
																					tmp14160 := PrimIsPair(W5312)

																					if True == tmp14160 {
																						tmp14088 := MakeNative(func(__e *ControlFlow) {
																							W5313 := __e.Get(1)
																							_ = W5313
																							tmp14089 := MakeNative(func(__e *ControlFlow) {
																								W5314 := __e.Get(1)
																								_ = W5314
																								tmp14155 := PrimIsPair(W5314)

																								if True == tmp14155 {
																									tmp14090 := MakeNative(func(__e *ControlFlow) {
																										W5315 := __e.Get(1)
																										_ = W5315
																										tmp14091 := MakeNative(func(__e *ControlFlow) {
																											W5316 := __e.Get(1)
																											_ = W5316
																											tmp14150 := PrimEqual(W5316, Nil)

																											if True == tmp14150 {
																												tmp14092 := MakeNative(func(__e *ControlFlow) {
																													W5317 := __e.Get(1)
																													_ = W5317
																													tmp14146 := PrimIsPair(W5317)

																													if True == tmp14146 {
																														tmp14093 := MakeNative(func(__e *ControlFlow) {
																															W5318 := __e.Get(1)
																															_ = W5318
																															tmp14094 := MakeNative(func(__e *ControlFlow) {
																																W5319 := __e.Get(1)
																																_ = W5319
																																tmp14141 := PrimIsPair(W5319)

																																if True == tmp14141 {
																																	tmp14095 := MakeNative(func(__e *ControlFlow) {
																																		W5320 := __e.Get(1)
																																		_ = W5320
																																		tmp14137 := PrimIsPair(W5320)

																																		if True == tmp14137 {
																																			tmp14096 := MakeNative(func(__e *ControlFlow) {
																																				W5321 := __e.Get(1)
																																				_ = W5321
																																				tmp14133 := PrimEqual(W5321, symvector)

																																				if True == tmp14133 {
																																					tmp14097 := MakeNative(func(__e *ControlFlow) {
																																						W5322 := __e.Get(1)
																																						_ = W5322
																																						tmp14129 := PrimIsPair(W5322)

																																						if True == tmp14129 {
																																							tmp14098 := MakeNative(func(__e *ControlFlow) {
																																								W5323 := __e.Get(1)
																																								_ = W5323
																																								tmp14099 := MakeNative(func(__e *ControlFlow) {
																																									W5324 := __e.Get(1)
																																									_ = W5324
																																									tmp14124 := PrimEqual(W5324, Nil)

																																									if True == tmp14124 {
																																										tmp14100 := MakeNative(func(__e *ControlFlow) {
																																											W5325 := __e.Get(1)
																																											_ = W5325
																																											tmp14120 := PrimEqual(W5325, Nil)

																																											if True == tmp14120 {
																																												tmp14101 := MakeNative(func(__e *ControlFlow) {
																																													W5326 := __e.Get(1)
																																													_ = W5326
																																													tmp14102 := Call(__e, PrimFunc(symshen_4incinfs))

																																													_ = tmp14102

																																													tmp14103 := Call(__e, PrimFunc(symshen_4deref), W5318, V5241)

																																													tmp14104 := PrimIntern(MakeString(":"))

																																													tmp14105 := PrimEqual(tmp14103, tmp14104)

																																													tmp14106 := MakeNative(func(__e *ControlFlow) {
																																														tmp14107 := MakeNative(func(__e *ControlFlow) {
																																															tmp14108 := PrimCons(W5323, Nil)

																																															tmp14109 := PrimCons(W5318, tmp14108)

																																															tmp14110 := PrimCons(W5313, tmp14109)

																																															tmp14111 := PrimCons(W5323, Nil)

																																															tmp14112 := PrimCons(symvector, tmp14111)

																																															tmp14113 := PrimCons(tmp14112, Nil)

																																															tmp14114 := PrimCons(W5318, tmp14113)

																																															tmp14115 := PrimCons(W5315, tmp14114)

																																															tmp14116 := PrimCons(tmp14115, W5326)

																																															tmp14117 := PrimCons(tmp14110, tmp14116)

																																															__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14117, V5239, True, V5241, V5242, W5245, V5244)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimFunc(symshen_4cut), V5241, V5242, W5245, tmp14107)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimFunc(symwhen), tmp14105, V5241, V5242, W5245, tmp14106)
																																													return

																																												}, 1)

																																												tmp14118 := PrimTail(W5308)

																																												__e.TailApply(tmp14101, tmp14118)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}, 1)

																																										tmp14121 := PrimTail(W5319)

																																										tmp14122 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14121, V5241)

																																										__e.TailApply(tmp14100, tmp14122)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp14125 := PrimTail(W5322)

																																								tmp14126 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14125, V5241)

																																								__e.TailApply(tmp14099, tmp14126)
																																								return

																																							}, 1)

																																							tmp14127 := PrimHead(W5322)

																																							__e.TailApply(tmp14098, tmp14127)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp14130 := PrimTail(W5320)

																																					tmp14131 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14130, V5241)

																																					__e.TailApply(tmp14097, tmp14131)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp14134 := PrimHead(W5320)

																																			tmp14135 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14134, V5241)

																																			__e.TailApply(tmp14096, tmp14135)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp14138 := PrimHead(W5319)

																																	tmp14139 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14138, V5241)

																																	__e.TailApply(tmp14095, tmp14139)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp14142 := PrimTail(W5317)

																															tmp14143 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14142, V5241)

																															__e.TailApply(tmp14094, tmp14143)
																															return

																														}, 1)

																														tmp14144 := PrimHead(W5317)

																														__e.TailApply(tmp14093, tmp14144)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp14147 := PrimTail(W5309)

																												tmp14148 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14147, V5241)

																												__e.TailApply(tmp14092, tmp14148)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp14151 := PrimTail(W5314)

																										tmp14152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14151, V5241)

																										__e.TailApply(tmp14091, tmp14152)
																										return

																									}, 1)

																									tmp14153 := PrimHead(W5314)

																									__e.TailApply(tmp14090, tmp14153)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp14156 := PrimTail(W5312)

																							tmp14157 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14156, V5241)

																							__e.TailApply(tmp14089, tmp14157)
																							return

																						}, 1)

																						tmp14158 := PrimHead(W5312)

																						__e.TailApply(tmp14088, tmp14158)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp14161 := PrimTail(W5310)

																				tmp14162 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14161, V5241)

																				__e.TailApply(tmp14087, tmp14162)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp14165 := PrimHead(W5310)

																		tmp14166 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14165, V5241)

																		__e.TailApply(tmp14086, tmp14166)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp14169 := PrimHead(W5309)

																tmp14170 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14169, V5241)

																__e.TailApply(tmp14085, tmp14170)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14173 := PrimHead(W5308)

														tmp14174 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14173, V5241)

														__e.TailApply(tmp14084, tmp14174)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14177 := Call(__e, PrimFunc(symshen_4lazyderef), V5238, V5241)

												tmp14178 := Call(__e, tmp14083, tmp14177)

												ifres14082 = tmp14178

											} else {
												ifres14082 = False

											}

											__e.TailApply(tmp14041, ifres14082)
											return

										} else {
											__e.Return(W5291)
											return
										}

									}, 1)

									tmp14260 := Call(__e, PrimFunc(symshen_4unlocked_2), V5242)

									var ifres14182 Obj

									if True == tmp14260 {
										tmp14183 := MakeNative(func(__e *ControlFlow) {
											W5292 := __e.Get(1)
											_ = W5292
											tmp14257 := PrimIsPair(W5292)

											if True == tmp14257 {
												tmp14184 := MakeNative(func(__e *ControlFlow) {
													W5293 := __e.Get(1)
													_ = W5293
													tmp14253 := PrimIsPair(W5293)

													if True == tmp14253 {
														tmp14185 := MakeNative(func(__e *ControlFlow) {
															W5294 := __e.Get(1)
															_ = W5294
															tmp14249 := PrimIsPair(W5294)

															if True == tmp14249 {
																tmp14186 := MakeNative(func(__e *ControlFlow) {
																	W5295 := __e.Get(1)
																	_ = W5295
																	tmp14245 := PrimEqual(W5295, sym_8s)

																	if True == tmp14245 {
																		tmp14187 := MakeNative(func(__e *ControlFlow) {
																			W5296 := __e.Get(1)
																			_ = W5296
																			tmp14241 := PrimIsPair(W5296)

																			if True == tmp14241 {
																				tmp14188 := MakeNative(func(__e *ControlFlow) {
																					W5297 := __e.Get(1)
																					_ = W5297
																					tmp14189 := MakeNative(func(__e *ControlFlow) {
																						W5298 := __e.Get(1)
																						_ = W5298
																						tmp14236 := PrimIsPair(W5298)

																						if True == tmp14236 {
																							tmp14190 := MakeNative(func(__e *ControlFlow) {
																								W5299 := __e.Get(1)
																								_ = W5299
																								tmp14191 := MakeNative(func(__e *ControlFlow) {
																									W5300 := __e.Get(1)
																									_ = W5300
																									tmp14231 := PrimEqual(W5300, Nil)

																									if True == tmp14231 {
																										tmp14192 := MakeNative(func(__e *ControlFlow) {
																											W5301 := __e.Get(1)
																											_ = W5301
																											tmp14227 := PrimIsPair(W5301)

																											if True == tmp14227 {
																												tmp14193 := MakeNative(func(__e *ControlFlow) {
																													W5302 := __e.Get(1)
																													_ = W5302
																													tmp14194 := MakeNative(func(__e *ControlFlow) {
																														W5303 := __e.Get(1)
																														_ = W5303
																														tmp14222 := PrimIsPair(W5303)

																														if True == tmp14222 {
																															tmp14195 := MakeNative(func(__e *ControlFlow) {
																																W5304 := __e.Get(1)
																																_ = W5304
																																tmp14218 := PrimEqual(W5304, symstring)

																																if True == tmp14218 {
																																	tmp14196 := MakeNative(func(__e *ControlFlow) {
																																		W5305 := __e.Get(1)
																																		_ = W5305
																																		tmp14214 := PrimEqual(W5305, Nil)

																																		if True == tmp14214 {
																																			tmp14197 := MakeNative(func(__e *ControlFlow) {
																																				W5306 := __e.Get(1)
																																				_ = W5306
																																				tmp14198 := Call(__e, PrimFunc(symshen_4incinfs))

																																				_ = tmp14198

																																				tmp14199 := Call(__e, PrimFunc(symshen_4deref), W5302, V5241)

																																				tmp14200 := PrimIntern(MakeString(":"))

																																				tmp14201 := PrimEqual(tmp14199, tmp14200)

																																				tmp14202 := MakeNative(func(__e *ControlFlow) {
																																					tmp14203 := MakeNative(func(__e *ControlFlow) {
																																						tmp14204 := PrimCons(symstring, Nil)

																																						tmp14205 := PrimCons(W5302, tmp14204)

																																						tmp14206 := PrimCons(W5297, tmp14205)

																																						tmp14207 := PrimCons(symstring, Nil)

																																						tmp14208 := PrimCons(W5302, tmp14207)

																																						tmp14209 := PrimCons(W5299, tmp14208)

																																						tmp14210 := PrimCons(tmp14209, W5306)

																																						tmp14211 := PrimCons(tmp14206, tmp14210)

																																						__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14211, V5239, True, V5241, V5242, W5245, V5244)
																																						return

																																					}, 0)

																																					__e.TailApply(PrimFunc(symshen_4cut), V5241, V5242, W5245, tmp14203)
																																					return

																																				}, 0)

																																				__e.TailApply(PrimFunc(symwhen), tmp14201, V5241, V5242, W5245, tmp14202)
																																				return

																																			}, 1)

																																			tmp14212 := PrimTail(W5292)

																																			__e.TailApply(tmp14197, tmp14212)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp14215 := PrimTail(W5303)

																																	tmp14216 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14215, V5241)

																																	__e.TailApply(tmp14196, tmp14216)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp14219 := PrimHead(W5303)

																															tmp14220 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14219, V5241)

																															__e.TailApply(tmp14195, tmp14220)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp14223 := PrimTail(W5301)

																													tmp14224 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14223, V5241)

																													__e.TailApply(tmp14194, tmp14224)
																													return

																												}, 1)

																												tmp14225 := PrimHead(W5301)

																												__e.TailApply(tmp14193, tmp14225)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp14228 := PrimTail(W5293)

																										tmp14229 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14228, V5241)

																										__e.TailApply(tmp14192, tmp14229)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp14232 := PrimTail(W5298)

																								tmp14233 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14232, V5241)

																								__e.TailApply(tmp14191, tmp14233)
																								return

																							}, 1)

																							tmp14234 := PrimHead(W5298)

																							__e.TailApply(tmp14190, tmp14234)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp14237 := PrimTail(W5296)

																					tmp14238 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14237, V5241)

																					__e.TailApply(tmp14189, tmp14238)
																					return

																				}, 1)

																				tmp14239 := PrimHead(W5296)

																				__e.TailApply(tmp14188, tmp14239)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp14242 := PrimTail(W5294)

																		tmp14243 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14242, V5241)

																		__e.TailApply(tmp14187, tmp14243)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp14246 := PrimHead(W5294)

																tmp14247 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14246, V5241)

																__e.TailApply(tmp14186, tmp14247)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14250 := PrimHead(W5293)

														tmp14251 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14250, V5241)

														__e.TailApply(tmp14185, tmp14251)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14254 := PrimHead(W5292)

												tmp14255 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14254, V5241)

												__e.TailApply(tmp14184, tmp14255)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14258 := Call(__e, PrimFunc(symshen_4lazyderef), V5238, V5241)

										tmp14259 := Call(__e, tmp14183, tmp14258)

										ifres14182 = tmp14259

									} else {
										ifres14182 = False

									}

									__e.TailApply(tmp14040, ifres14182)
									return

								} else {
									__e.Return(W5269)
									return
								}

							}, 1)

							tmp14365 := Call(__e, PrimFunc(symshen_4unlocked_2), V5242)

							var ifres14263 Obj

							if True == tmp14365 {
								tmp14264 := MakeNative(func(__e *ControlFlow) {
									W5270 := __e.Get(1)
									_ = W5270
									tmp14362 := PrimIsPair(W5270)

									if True == tmp14362 {
										tmp14265 := MakeNative(func(__e *ControlFlow) {
											W5271 := __e.Get(1)
											_ = W5271
											tmp14358 := PrimIsPair(W5271)

											if True == tmp14358 {
												tmp14266 := MakeNative(func(__e *ControlFlow) {
													W5272 := __e.Get(1)
													_ = W5272
													tmp14354 := PrimIsPair(W5272)

													if True == tmp14354 {
														tmp14267 := MakeNative(func(__e *ControlFlow) {
															W5273 := __e.Get(1)
															_ = W5273
															tmp14350 := PrimEqual(W5273, sym_8p)

															if True == tmp14350 {
																tmp14268 := MakeNative(func(__e *ControlFlow) {
																	W5274 := __e.Get(1)
																	_ = W5274
																	tmp14346 := PrimIsPair(W5274)

																	if True == tmp14346 {
																		tmp14269 := MakeNative(func(__e *ControlFlow) {
																			W5275 := __e.Get(1)
																			_ = W5275
																			tmp14270 := MakeNative(func(__e *ControlFlow) {
																				W5276 := __e.Get(1)
																				_ = W5276
																				tmp14341 := PrimIsPair(W5276)

																				if True == tmp14341 {
																					tmp14271 := MakeNative(func(__e *ControlFlow) {
																						W5277 := __e.Get(1)
																						_ = W5277
																						tmp14272 := MakeNative(func(__e *ControlFlow) {
																							W5278 := __e.Get(1)
																							_ = W5278
																							tmp14336 := PrimEqual(W5278, Nil)

																							if True == tmp14336 {
																								tmp14273 := MakeNative(func(__e *ControlFlow) {
																									W5279 := __e.Get(1)
																									_ = W5279
																									tmp14332 := PrimIsPair(W5279)

																									if True == tmp14332 {
																										tmp14274 := MakeNative(func(__e *ControlFlow) {
																											W5280 := __e.Get(1)
																											_ = W5280
																											tmp14275 := MakeNative(func(__e *ControlFlow) {
																												W5281 := __e.Get(1)
																												_ = W5281
																												tmp14327 := PrimIsPair(W5281)

																												if True == tmp14327 {
																													tmp14276 := MakeNative(func(__e *ControlFlow) {
																														W5282 := __e.Get(1)
																														_ = W5282
																														tmp14323 := PrimIsPair(W5282)

																														if True == tmp14323 {
																															tmp14277 := MakeNative(func(__e *ControlFlow) {
																																W5283 := __e.Get(1)
																																_ = W5283
																																tmp14278 := MakeNative(func(__e *ControlFlow) {
																																	W5284 := __e.Get(1)
																																	_ = W5284
																																	tmp14318 := PrimIsPair(W5284)

																																	if True == tmp14318 {
																																		tmp14279 := MakeNative(func(__e *ControlFlow) {
																																			W5285 := __e.Get(1)
																																			_ = W5285
																																			tmp14314 := PrimEqual(W5285, sym_d)

																																			if True == tmp14314 {
																																				tmp14280 := MakeNative(func(__e *ControlFlow) {
																																					W5286 := __e.Get(1)
																																					_ = W5286
																																					tmp14310 := PrimIsPair(W5286)

																																					if True == tmp14310 {
																																						tmp14281 := MakeNative(func(__e *ControlFlow) {
																																							W5287 := __e.Get(1)
																																							_ = W5287
																																							tmp14282 := MakeNative(func(__e *ControlFlow) {
																																								W5288 := __e.Get(1)
																																								_ = W5288
																																								tmp14305 := PrimEqual(W5288, Nil)

																																								if True == tmp14305 {
																																									tmp14283 := MakeNative(func(__e *ControlFlow) {
																																										W5289 := __e.Get(1)
																																										_ = W5289
																																										tmp14301 := PrimEqual(W5289, Nil)

																																										if True == tmp14301 {
																																											tmp14284 := MakeNative(func(__e *ControlFlow) {
																																												W5290 := __e.Get(1)
																																												_ = W5290
																																												tmp14285 := Call(__e, PrimFunc(symshen_4incinfs))

																																												_ = tmp14285

																																												tmp14286 := Call(__e, PrimFunc(symshen_4deref), W5280, V5241)

																																												tmp14287 := PrimIntern(MakeString(":"))

																																												tmp14288 := PrimEqual(tmp14286, tmp14287)

																																												tmp14289 := MakeNative(func(__e *ControlFlow) {
																																													tmp14290 := MakeNative(func(__e *ControlFlow) {
																																														tmp14291 := PrimCons(W5283, Nil)

																																														tmp14292 := PrimCons(W5280, tmp14291)

																																														tmp14293 := PrimCons(W5275, tmp14292)

																																														tmp14294 := PrimCons(W5287, Nil)

																																														tmp14295 := PrimCons(W5280, tmp14294)

																																														tmp14296 := PrimCons(W5277, tmp14295)

																																														tmp14297 := PrimCons(tmp14296, W5290)

																																														tmp14298 := PrimCons(tmp14293, tmp14297)

																																														__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14298, V5239, True, V5241, V5242, W5245, V5244)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimFunc(symshen_4cut), V5241, V5242, W5245, tmp14290)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimFunc(symwhen), tmp14288, V5241, V5242, W5245, tmp14289)
																																												return

																																											}, 1)

																																											tmp14299 := PrimTail(W5270)

																																											__e.TailApply(tmp14284, tmp14299)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp14302 := PrimTail(W5281)

																																									tmp14303 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14302, V5241)

																																									__e.TailApply(tmp14283, tmp14303)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp14306 := PrimTail(W5286)

																																							tmp14307 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14306, V5241)

																																							__e.TailApply(tmp14282, tmp14307)
																																							return

																																						}, 1)

																																						tmp14308 := PrimHead(W5286)

																																						__e.TailApply(tmp14281, tmp14308)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp14311 := PrimTail(W5284)

																																				tmp14312 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14311, V5241)

																																				__e.TailApply(tmp14280, tmp14312)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp14315 := PrimHead(W5284)

																																		tmp14316 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14315, V5241)

																																		__e.TailApply(tmp14279, tmp14316)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp14319 := PrimTail(W5282)

																																tmp14320 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14319, V5241)

																																__e.TailApply(tmp14278, tmp14320)
																																return

																															}, 1)

																															tmp14321 := PrimHead(W5282)

																															__e.TailApply(tmp14277, tmp14321)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp14324 := PrimHead(W5281)

																													tmp14325 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14324, V5241)

																													__e.TailApply(tmp14276, tmp14325)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp14328 := PrimTail(W5279)

																											tmp14329 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14328, V5241)

																											__e.TailApply(tmp14275, tmp14329)
																											return

																										}, 1)

																										tmp14330 := PrimHead(W5279)

																										__e.TailApply(tmp14274, tmp14330)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp14333 := PrimTail(W5271)

																								tmp14334 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14333, V5241)

																								__e.TailApply(tmp14273, tmp14334)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp14337 := PrimTail(W5276)

																						tmp14338 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14337, V5241)

																						__e.TailApply(tmp14272, tmp14338)
																						return

																					}, 1)

																					tmp14339 := PrimHead(W5276)

																					__e.TailApply(tmp14271, tmp14339)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp14342 := PrimTail(W5274)

																			tmp14343 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14342, V5241)

																			__e.TailApply(tmp14270, tmp14343)
																			return

																		}, 1)

																		tmp14344 := PrimHead(W5274)

																		__e.TailApply(tmp14269, tmp14344)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp14347 := PrimTail(W5272)

																tmp14348 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14347, V5241)

																__e.TailApply(tmp14268, tmp14348)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14351 := PrimHead(W5272)

														tmp14352 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14351, V5241)

														__e.TailApply(tmp14267, tmp14352)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14355 := PrimHead(W5271)

												tmp14356 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14355, V5241)

												__e.TailApply(tmp14266, tmp14356)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14359 := PrimHead(W5270)

										tmp14360 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14359, V5241)

										__e.TailApply(tmp14265, tmp14360)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14363 := Call(__e, PrimFunc(symshen_4lazyderef), V5238, V5241)

								tmp14364 := Call(__e, tmp14264, tmp14363)

								ifres14263 = tmp14364

							} else {
								ifres14263 = False

							}

							__e.TailApply(tmp14039, ifres14263)
							return

						} else {
							__e.Return(W5249)
							return
						}

					}, 1)

					tmp14465 := Call(__e, PrimFunc(symshen_4unlocked_2), V5242)

					var ifres14368 Obj

					if True == tmp14465 {
						tmp14369 := MakeNative(func(__e *ControlFlow) {
							W5250 := __e.Get(1)
							_ = W5250
							tmp14462 := PrimIsPair(W5250)

							if True == tmp14462 {
								tmp14370 := MakeNative(func(__e *ControlFlow) {
									W5251 := __e.Get(1)
									_ = W5251
									tmp14458 := PrimIsPair(W5251)

									if True == tmp14458 {
										tmp14371 := MakeNative(func(__e *ControlFlow) {
											W5252 := __e.Get(1)
											_ = W5252
											tmp14454 := PrimIsPair(W5252)

											if True == tmp14454 {
												tmp14372 := MakeNative(func(__e *ControlFlow) {
													W5253 := __e.Get(1)
													_ = W5253
													tmp14450 := PrimEqual(W5253, symcons)

													if True == tmp14450 {
														tmp14373 := MakeNative(func(__e *ControlFlow) {
															W5254 := __e.Get(1)
															_ = W5254
															tmp14446 := PrimIsPair(W5254)

															if True == tmp14446 {
																tmp14374 := MakeNative(func(__e *ControlFlow) {
																	W5255 := __e.Get(1)
																	_ = W5255
																	tmp14375 := MakeNative(func(__e *ControlFlow) {
																		W5256 := __e.Get(1)
																		_ = W5256
																		tmp14441 := PrimIsPair(W5256)

																		if True == tmp14441 {
																			tmp14376 := MakeNative(func(__e *ControlFlow) {
																				W5257 := __e.Get(1)
																				_ = W5257
																				tmp14377 := MakeNative(func(__e *ControlFlow) {
																					W5258 := __e.Get(1)
																					_ = W5258
																					tmp14436 := PrimEqual(W5258, Nil)

																					if True == tmp14436 {
																						tmp14378 := MakeNative(func(__e *ControlFlow) {
																							W5259 := __e.Get(1)
																							_ = W5259
																							tmp14432 := PrimIsPair(W5259)

																							if True == tmp14432 {
																								tmp14379 := MakeNative(func(__e *ControlFlow) {
																									W5260 := __e.Get(1)
																									_ = W5260
																									tmp14380 := MakeNative(func(__e *ControlFlow) {
																										W5261 := __e.Get(1)
																										_ = W5261
																										tmp14427 := PrimIsPair(W5261)

																										if True == tmp14427 {
																											tmp14381 := MakeNative(func(__e *ControlFlow) {
																												W5262 := __e.Get(1)
																												_ = W5262
																												tmp14423 := PrimIsPair(W5262)

																												if True == tmp14423 {
																													tmp14382 := MakeNative(func(__e *ControlFlow) {
																														W5263 := __e.Get(1)
																														_ = W5263
																														tmp14419 := PrimEqual(W5263, symlist)

																														if True == tmp14419 {
																															tmp14383 := MakeNative(func(__e *ControlFlow) {
																																W5264 := __e.Get(1)
																																_ = W5264
																																tmp14415 := PrimIsPair(W5264)

																																if True == tmp14415 {
																																	tmp14384 := MakeNative(func(__e *ControlFlow) {
																																		W5265 := __e.Get(1)
																																		_ = W5265
																																		tmp14385 := MakeNative(func(__e *ControlFlow) {
																																			W5266 := __e.Get(1)
																																			_ = W5266
																																			tmp14410 := PrimEqual(W5266, Nil)

																																			if True == tmp14410 {
																																				tmp14386 := MakeNative(func(__e *ControlFlow) {
																																					W5267 := __e.Get(1)
																																					_ = W5267
																																					tmp14406 := PrimEqual(W5267, Nil)

																																					if True == tmp14406 {
																																						tmp14387 := MakeNative(func(__e *ControlFlow) {
																																							W5268 := __e.Get(1)
																																							_ = W5268
																																							tmp14388 := Call(__e, PrimFunc(symshen_4incinfs))

																																							_ = tmp14388

																																							tmp14389 := Call(__e, PrimFunc(symshen_4deref), W5260, V5241)

																																							tmp14390 := PrimIntern(MakeString(":"))

																																							tmp14391 := PrimEqual(tmp14389, tmp14390)

																																							tmp14392 := MakeNative(func(__e *ControlFlow) {
																																								tmp14393 := MakeNative(func(__e *ControlFlow) {
																																									tmp14394 := PrimCons(W5265, Nil)

																																									tmp14395 := PrimCons(W5260, tmp14394)

																																									tmp14396 := PrimCons(W5255, tmp14395)

																																									tmp14397 := PrimCons(W5265, Nil)

																																									tmp14398 := PrimCons(symlist, tmp14397)

																																									tmp14399 := PrimCons(tmp14398, Nil)

																																									tmp14400 := PrimCons(W5260, tmp14399)

																																									tmp14401 := PrimCons(W5257, tmp14400)

																																									tmp14402 := PrimCons(tmp14401, W5268)

																																									tmp14403 := PrimCons(tmp14396, tmp14402)

																																									__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14403, V5239, True, V5241, V5242, W5245, V5244)
																																									return

																																								}, 0)

																																								__e.TailApply(PrimFunc(symshen_4cut), V5241, V5242, W5245, tmp14393)
																																								return

																																							}, 0)

																																							__e.TailApply(PrimFunc(symwhen), tmp14391, V5241, V5242, W5245, tmp14392)
																																							return

																																						}, 1)

																																						tmp14404 := PrimTail(W5250)

																																						__e.TailApply(tmp14387, tmp14404)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp14407 := PrimTail(W5261)

																																				tmp14408 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14407, V5241)

																																				__e.TailApply(tmp14386, tmp14408)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp14411 := PrimTail(W5264)

																																		tmp14412 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14411, V5241)

																																		__e.TailApply(tmp14385, tmp14412)
																																		return

																																	}, 1)

																																	tmp14413 := PrimHead(W5264)

																																	__e.TailApply(tmp14384, tmp14413)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp14416 := PrimTail(W5262)

																															tmp14417 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14416, V5241)

																															__e.TailApply(tmp14383, tmp14417)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp14420 := PrimHead(W5262)

																													tmp14421 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14420, V5241)

																													__e.TailApply(tmp14382, tmp14421)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp14424 := PrimHead(W5261)

																											tmp14425 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14424, V5241)

																											__e.TailApply(tmp14381, tmp14425)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp14428 := PrimTail(W5259)

																									tmp14429 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14428, V5241)

																									__e.TailApply(tmp14380, tmp14429)
																									return

																								}, 1)

																								tmp14430 := PrimHead(W5259)

																								__e.TailApply(tmp14379, tmp14430)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp14433 := PrimTail(W5251)

																						tmp14434 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14433, V5241)

																						__e.TailApply(tmp14378, tmp14434)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp14437 := PrimTail(W5256)

																				tmp14438 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14437, V5241)

																				__e.TailApply(tmp14377, tmp14438)
																				return

																			}, 1)

																			tmp14439 := PrimHead(W5256)

																			__e.TailApply(tmp14376, tmp14439)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp14442 := PrimTail(W5254)

																	tmp14443 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14442, V5241)

																	__e.TailApply(tmp14375, tmp14443)
																	return

																}, 1)

																tmp14444 := PrimHead(W5254)

																__e.TailApply(tmp14374, tmp14444)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp14447 := PrimTail(W5252)

														tmp14448 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14447, V5241)

														__e.TailApply(tmp14373, tmp14448)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp14451 := PrimHead(W5252)

												tmp14452 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14451, V5241)

												__e.TailApply(tmp14372, tmp14452)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp14455 := PrimHead(W5251)

										tmp14456 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14455, V5241)

										__e.TailApply(tmp14371, tmp14456)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14459 := PrimHead(W5250)

								tmp14460 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14459, V5241)

								__e.TailApply(tmp14370, tmp14460)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14463 := Call(__e, PrimFunc(symshen_4lazyderef), V5238, V5241)

						tmp14464 := Call(__e, tmp14369, tmp14463)

						ifres14368 = tmp14464

					} else {
						ifres14368 = False

					}

					__e.TailApply(tmp14038, ifres14368)
					return

				} else {
					__e.Return(W5246)
					return
				}

			}, 1)

			tmp14480 := Call(__e, PrimFunc(symshen_4unlocked_2), V5242)

			var ifres14468 Obj

			if True == tmp14480 {
				tmp14469 := MakeNative(func(__e *ControlFlow) {
					W5247 := __e.Get(1)
					_ = W5247
					tmp14477 := PrimEqual(W5247, Nil)

					if True == tmp14477 {
						tmp14470 := MakeNative(func(__e *ControlFlow) {
							W5248 := __e.Get(1)
							_ = W5248
							tmp14474 := PrimEqual(W5248, True)

							if True == tmp14474 {
								tmp14471 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp14471

								tmp14472 := MakeNative(func(__e *ControlFlow) {
									__e.TailApply(PrimFunc(symbind), V5239, Nil, V5241, V5242, W5245, V5244)
									return
								}, 0)

								__e.TailApply(PrimFunc(symshen_4cut), V5241, V5242, W5245, tmp14472)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14475 := Call(__e, PrimFunc(symshen_4lazyderef), V5240, V5241)

						__e.TailApply(tmp14470, tmp14475)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14478 := Call(__e, PrimFunc(symshen_4lazyderef), V5238, V5241)

				tmp14479 := Call(__e, tmp14469, tmp14478)

				ifres14468 = tmp14479

			} else {
				ifres14468 = False

			}

			__e.TailApply(tmp14037, ifres14468)
			return

		}, 1)

		tmp14481 := PrimNumberAdd(V5243, MakeNumber(1))

		__e.TailApply(tmp14036, tmp14481)
		return

	}, 7)

	tmp14482 := Call(__e, ns2_1set, symshen_4l_1rules, tmp14035)

	_ = tmp14482

	tmp14483 := MakeNative(func(__e *ControlFlow) {
		V5339 := __e.Get(1)
		_ = V5339
		V5340 := __e.Get(2)
		_ = V5340
		V5341 := __e.Get(3)
		_ = V5341
		V5342 := __e.Get(4)
		_ = V5342
		V5343 := __e.Get(5)
		_ = V5343
		V5344 := __e.Get(6)
		_ = V5344
		tmp14484 := MakeNative(func(__e *ControlFlow) {
			W5345 := __e.Get(1)
			_ = W5345
			tmp14485 := MakeNative(func(__e *ControlFlow) {
				W5346 := __e.Get(1)
				_ = W5346
				tmp14487 := PrimEqual(W5346, False)

				if True == tmp14487 {
					__e.TailApply(PrimFunc(symshen_4unlock), V5342, W5345)
					return
				} else {
					__e.Return(W5346)
					return
				}

			}, 1)

			tmp14535 := Call(__e, PrimFunc(symshen_4unlocked_2), V5342)

			var ifres14488 Obj

			if True == tmp14535 {
				tmp14489 := MakeNative(func(__e *ControlFlow) {
					W5347 := __e.Get(1)
					_ = W5347
					tmp14532 := PrimIsPair(W5347)

					if True == tmp14532 {
						tmp14490 := MakeNative(func(__e *ControlFlow) {
							W5348 := __e.Get(1)
							_ = W5348
							tmp14528 := PrimEqual(W5348, symdefine)

							if True == tmp14528 {
								tmp14491 := MakeNative(func(__e *ControlFlow) {
									W5349 := __e.Get(1)
									_ = W5349
									tmp14524 := PrimIsPair(W5349)

									if True == tmp14524 {
										tmp14492 := MakeNative(func(__e *ControlFlow) {
											W5350 := __e.Get(1)
											_ = W5350
											tmp14493 := MakeNative(func(__e *ControlFlow) {
												W5351 := __e.Get(1)
												_ = W5351
												tmp14494 := MakeNative(func(__e *ControlFlow) {
													W5352 := __e.Get(1)
													_ = W5352
													tmp14495 := MakeNative(func(__e *ControlFlow) {
														W5353 := __e.Get(1)
														_ = W5353
														tmp14496 := MakeNative(func(__e *ControlFlow) {
															W5354 := __e.Get(1)
															_ = W5354
															tmp14497 := MakeNative(func(__e *ControlFlow) {
																W5355 := __e.Get(1)
																_ = W5355
																tmp14498 := Call(__e, PrimFunc(symshen_4incinfs))

																_ = tmp14498

																tmp14499 := MakeNative(func(__e *ControlFlow) {
																	tmp14500 := PrimCons(W5350, W5351)

																	tmp14501 := Call(__e, PrimFunc(symshen_4sigxrules), tmp14500)

																	tmp14502 := MakeNative(func(__e *ControlFlow) {
																		tmp14503 := Call(__e, PrimFunc(symshen_4lazyderef), W5352, V5341)

																		tmp14504 := Call(__e, PrimFunc(symfst), tmp14503)

																		tmp14505 := MakeNative(func(__e *ControlFlow) {
																			tmp14506 := Call(__e, PrimFunc(symshen_4lazyderef), W5352, V5341)

																			tmp14507 := Call(__e, PrimFunc(symsnd), tmp14506)

																			tmp14508 := MakeNative(func(__e *ControlFlow) {
																				tmp14509 := Call(__e, PrimFunc(symshen_4deref), W5355, V5341)

																				tmp14510 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp14509)

																				tmp14511 := MakeNative(func(__e *ControlFlow) {
																					tmp14512 := MakeNative(func(__e *ControlFlow) {
																						__e.TailApply(PrimFunc(symis), W5355, V5340, V5341, V5342, W5345, V5344)
																						return
																					}, 0)

																					__e.TailApply(PrimFunc(symshen_4t_d_1rules), W5350, W5353, W5354, MakeNumber(1), V5341, V5342, W5345, tmp14512)
																					return

																				}, 0)

																				__e.TailApply(PrimFunc(symbind), W5354, tmp14510, V5341, V5342, W5345, tmp14511)
																				return

																			}, 0)

																			__e.TailApply(PrimFunc(symbind), W5353, tmp14507, V5341, V5342, W5345, tmp14508)
																			return

																		}, 0)

																		__e.TailApply(PrimFunc(symbind), W5355, tmp14504, V5341, V5342, W5345, tmp14505)
																		return

																	}, 0)

																	__e.TailApply(PrimFunc(symbind), W5352, tmp14501, V5341, V5342, W5345, tmp14502)
																	return

																}, 0)

																tmp14513 := Call(__e, PrimFunc(symshen_4cut), V5341, V5342, W5345, tmp14499)

																__e.TailApply(PrimFunc(symshen_4gc), V5341, tmp14513)
																return

															}, 1)

															tmp14514 := Call(__e, PrimFunc(symshen_4newpv), V5341)

															tmp14515 := Call(__e, tmp14497, tmp14514)

															__e.TailApply(PrimFunc(symshen_4gc), V5341, tmp14515)
															return

														}, 1)

														tmp14516 := Call(__e, PrimFunc(symshen_4newpv), V5341)

														tmp14517 := Call(__e, tmp14496, tmp14516)

														__e.TailApply(PrimFunc(symshen_4gc), V5341, tmp14517)
														return

													}, 1)

													tmp14518 := Call(__e, PrimFunc(symshen_4newpv), V5341)

													tmp14519 := Call(__e, tmp14495, tmp14518)

													__e.TailApply(PrimFunc(symshen_4gc), V5341, tmp14519)
													return

												}, 1)

												tmp14520 := Call(__e, PrimFunc(symshen_4newpv), V5341)

												__e.TailApply(tmp14494, tmp14520)
												return

											}, 1)

											tmp14521 := PrimTail(W5349)

											__e.TailApply(tmp14493, tmp14521)
											return

										}, 1)

										tmp14522 := PrimHead(W5349)

										__e.TailApply(tmp14492, tmp14522)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp14525 := PrimTail(W5347)

								tmp14526 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14525, V5341)

								__e.TailApply(tmp14491, tmp14526)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14529 := PrimHead(W5347)

						tmp14530 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14529, V5341)

						__e.TailApply(tmp14490, tmp14530)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14533 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5341)

				tmp14534 := Call(__e, tmp14489, tmp14533)

				ifres14488 = tmp14534

			} else {
				ifres14488 = False

			}

			__e.TailApply(tmp14485, ifres14488)
			return

		}, 1)

		tmp14536 := PrimNumberAdd(V5343, MakeNumber(1))

		__e.TailApply(tmp14484, tmp14536)
		return

	}, 6)

	tmp14537 := Call(__e, ns2_1set, symshen_4t_d, tmp14483)

	_ = tmp14537

	tmp14538 := MakeNative(func(__e *ControlFlow) {
		V5356 := __e.Get(1)
		_ = V5356
		tmp14539 := MakeNative(func(__e *ControlFlow) {
			Z5357 := __e.Get(1)
			_ = Z5357
			__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5357)
			return
		}, 1)

		__e.TailApply(PrimFunc(symcompile), tmp14539, V5356)
		return

	}, 1)

	tmp14540 := Call(__e, ns2_1set, symshen_4sigxrules, tmp14538)

	_ = tmp14540

	tmp14541 := MakeNative(func(__e *ControlFlow) {
		V5358 := __e.Get(1)
		_ = V5358
		tmp14542 := MakeNative(func(__e *ControlFlow) {
			W5359 := __e.Get(1)
			_ = W5359
			tmp14544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5359)

			if True == tmp14544 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W5359)
				return
			}

		}, 1)

		tmp14577 := PrimIsPair(V5358)

		var ifres14545 Obj

		if True == tmp14577 {
			tmp14546 := MakeNative(func(__e *ControlFlow) {
				W5360 := __e.Get(1)
				_ = W5360
				tmp14573 := Call(__e, PrimFunc(symshen_4hds_a_2), W5360, sym_i)

				if True == tmp14573 {
					tmp14547 := MakeNative(func(__e *ControlFlow) {
						W5361 := __e.Get(1)
						_ = W5361
						tmp14548 := MakeNative(func(__e *ControlFlow) {
							W5362 := __e.Get(1)
							_ = W5362
							tmp14569 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5362)

							if True == tmp14569 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp14549 := MakeNative(func(__e *ControlFlow) {
									W5363 := __e.Get(1)
									_ = W5363
									tmp14550 := MakeNative(func(__e *ControlFlow) {
										W5364 := __e.Get(1)
										_ = W5364
										tmp14565 := Call(__e, PrimFunc(symshen_4hds_a_2), W5364, sym_j)

										if True == tmp14565 {
											tmp14551 := MakeNative(func(__e *ControlFlow) {
												W5365 := __e.Get(1)
												_ = W5365
												tmp14552 := MakeNative(func(__e *ControlFlow) {
													W5366 := __e.Get(1)
													_ = W5366
													tmp14561 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5366)

													if True == tmp14561 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp14553 := MakeNative(func(__e *ControlFlow) {
															W5367 := __e.Get(1)
															_ = W5367
															tmp14554 := MakeNative(func(__e *ControlFlow) {
																W5368 := __e.Get(1)
																_ = W5368
																tmp14555 := MakeNative(func(__e *ControlFlow) {
																	W5369 := __e.Get(1)
																	_ = W5369
																	__e.TailApply(PrimFunc(sym_8p), W5369, W5367)
																	return
																}, 1)

																tmp14556 := Call(__e, PrimFunc(symshen_4rectify_1type), W5363)

																tmp14557 := Call(__e, tmp14555, tmp14556)

																__e.TailApply(PrimFunc(symshen_4comb), W5368, tmp14557)
																return

															}, 1)

															tmp14558 := Call(__e, PrimFunc(symshen_4in_1_6), W5366)

															__e.TailApply(tmp14554, tmp14558)
															return

														}, 1)

														tmp14559 := Call(__e, PrimFunc(symshen_4_5_1out), W5366)

														__e.TailApply(tmp14553, tmp14559)
														return

													}

												}, 1)

												tmp14562 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5365)

												__e.TailApply(tmp14552, tmp14562)
												return

											}, 1)

											tmp14563 := Call(__e, PrimFunc(symtail), W5364)

											__e.TailApply(tmp14551, tmp14563)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp14566 := Call(__e, PrimFunc(symshen_4in_1_6), W5362)

									__e.TailApply(tmp14550, tmp14566)
									return

								}, 1)

								tmp14567 := Call(__e, PrimFunc(symshen_4_5_1out), W5362)

								__e.TailApply(tmp14549, tmp14567)
								return

							}

						}, 1)

						tmp14570 := Call(__e, PrimFunc(symshen_4_5signature_6), W5361)

						__e.TailApply(tmp14548, tmp14570)
						return

					}, 1)

					tmp14571 := Call(__e, PrimFunc(symtail), W5360)

					__e.TailApply(tmp14547, tmp14571)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4parse_1failure))
					return
				}

			}, 1)

			tmp14574 := Call(__e, PrimFunc(symtail), V5358)

			tmp14575 := Call(__e, tmp14546, tmp14574)

			ifres14545 = tmp14575

		} else {
			tmp14576 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres14545 = tmp14576

		}

		__e.TailApply(tmp14542, ifres14545)
		return

	}, 1)

	tmp14578 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp14541)

	_ = tmp14578

	tmp14579 := MakeNative(func(__e *ControlFlow) {
		V5370 := __e.Get(1)
		_ = V5370
		tmp14580 := MakeNative(func(__e *ControlFlow) {
			W5371 := __e.Get(1)
			_ = W5371
			tmp14581 := MakeNative(func(__e *ControlFlow) {
				W5372 := __e.Get(1)
				_ = W5372
				__e.TailApply(PrimFunc(symshen_4freshen_1type), W5372, V5370)
				return
			}, 1)

			tmp14582 := MakeNative(func(__e *ControlFlow) {
				Z5373 := __e.Get(1)
				_ = Z5373
				tmp14583 := Call(__e, PrimFunc(symconcat), sym_e, Z5373)

				tmp14584 := Call(__e, PrimFunc(symshen_4freshterm), tmp14583)

				__e.Return(PrimCons(Z5373, tmp14584))
				return

			}, 1)

			tmp14585 := Call(__e, PrimFunc(symmap), tmp14582, W5371)

			__e.TailApply(tmp14581, tmp14585)
			return

		}, 1)

		tmp14586 := Call(__e, PrimFunc(symshen_4extract_1vars), V5370)

		__e.TailApply(tmp14580, tmp14586)
		return

	}, 1)

	tmp14587 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp14579)

	_ = tmp14587

	tmp14588 := MakeNative(func(__e *ControlFlow) {
		V5374 := __e.Get(1)
		_ = V5374
		V5375 := __e.Get(2)
		_ = V5375
		tmp14602 := PrimEqual(Nil, V5374)

		if True == tmp14602 {
			__e.Return(V5375)
			return
		} else {
			tmp14600 := PrimIsPair(V5374)

			var ifres14596 Obj

			if True == tmp14600 {
				tmp14598 := PrimHead(V5374)

				tmp14599 := PrimIsPair(tmp14598)

				var ifres14597 Obj

				if True == tmp14599 {
					ifres14597 = True

				} else {
					ifres14597 = False

				}

				ifres14596 = ifres14597

			} else {
				ifres14596 = False

			}

			if True == ifres14596 {
				tmp14589 := PrimTail(V5374)

				tmp14590 := PrimHead(V5374)

				tmp14591 := PrimTail(tmp14590)

				tmp14592 := PrimHead(V5374)

				tmp14593 := PrimHead(tmp14592)

				tmp14594 := Call(__e, PrimFunc(symsubst), tmp14591, tmp14593, V5375)

				__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp14589, tmp14594)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
				return
			}

		}

	}, 2)

	tmp14603 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp14588)

	_ = tmp14603

	tmp14604 := MakeNative(func(__e *ControlFlow) {
		V5376 := __e.Get(1)
		_ = V5376
		tmp14605 := MakeNative(func(__e *ControlFlow) {
			W5377 := __e.Get(1)
			_ = W5377
			tmp14620 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5377)

			if True == tmp14620 {
				tmp14606 := MakeNative(func(__e *ControlFlow) {
					W5384 := __e.Get(1)
					_ = W5384
					tmp14608 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5384)

					if True == tmp14608 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W5384)
						return
					}

				}, 1)

				tmp14609 := MakeNative(func(__e *ControlFlow) {
					W5385 := __e.Get(1)
					_ = W5385
					tmp14616 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5385)

					if True == tmp14616 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp14610 := MakeNative(func(__e *ControlFlow) {
							W5386 := __e.Get(1)
							_ = W5386
							tmp14611 := MakeNative(func(__e *ControlFlow) {
								W5387 := __e.Get(1)
								_ = W5387
								tmp14612 := PrimCons(W5386, Nil)

								__e.TailApply(PrimFunc(symshen_4comb), W5387, tmp14612)
								return

							}, 1)

							tmp14613 := Call(__e, PrimFunc(symshen_4in_1_6), W5385)

							__e.TailApply(tmp14611, tmp14613)
							return

						}, 1)

						tmp14614 := Call(__e, PrimFunc(symshen_4_5_1out), W5385)

						__e.TailApply(tmp14610, tmp14614)
						return

					}

				}, 1)

				tmp14617 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5376)

				tmp14618 := Call(__e, tmp14609, tmp14617)

				__e.TailApply(tmp14606, tmp14618)
				return

			} else {
				__e.Return(W5377)
				return
			}

		}, 1)

		tmp14621 := MakeNative(func(__e *ControlFlow) {
			W5378 := __e.Get(1)
			_ = W5378
			tmp14636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5378)

			if True == tmp14636 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp14622 := MakeNative(func(__e *ControlFlow) {
					W5379 := __e.Get(1)
					_ = W5379
					tmp14623 := MakeNative(func(__e *ControlFlow) {
						W5380 := __e.Get(1)
						_ = W5380
						tmp14624 := MakeNative(func(__e *ControlFlow) {
							W5381 := __e.Get(1)
							_ = W5381
							tmp14631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5381)

							if True == tmp14631 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp14625 := MakeNative(func(__e *ControlFlow) {
									W5382 := __e.Get(1)
									_ = W5382
									tmp14626 := MakeNative(func(__e *ControlFlow) {
										W5383 := __e.Get(1)
										_ = W5383
										tmp14627 := PrimCons(W5379, W5382)

										__e.TailApply(PrimFunc(symshen_4comb), W5383, tmp14627)
										return

									}, 1)

									tmp14628 := Call(__e, PrimFunc(symshen_4in_1_6), W5381)

									__e.TailApply(tmp14626, tmp14628)
									return

								}, 1)

								tmp14629 := Call(__e, PrimFunc(symshen_4_5_1out), W5381)

								__e.TailApply(tmp14625, tmp14629)
								return

							}

						}, 1)

						tmp14632 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5380)

						__e.TailApply(tmp14624, tmp14632)
						return

					}, 1)

					tmp14633 := Call(__e, PrimFunc(symshen_4in_1_6), W5378)

					__e.TailApply(tmp14623, tmp14633)
					return

				}, 1)

				tmp14634 := Call(__e, PrimFunc(symshen_4_5_1out), W5378)

				__e.TailApply(tmp14622, tmp14634)
				return

			}

		}, 1)

		tmp14637 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5376)

		tmp14638 := Call(__e, tmp14621, tmp14637)

		__e.TailApply(tmp14605, tmp14638)
		return

	}, 1)

	tmp14639 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp14604)

	_ = tmp14639

	tmp14640 := MakeNative(func(__e *ControlFlow) {
		V5388 := __e.Get(1)
		_ = V5388
		tmp14641 := MakeNative(func(__e *ControlFlow) {
			W5389 := __e.Get(1)
			_ = W5389
			tmp14727 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5389)

			if True == tmp14727 {
				tmp14642 := MakeNative(func(__e *ControlFlow) {
					W5399 := __e.Get(1)
					_ = W5399
					tmp14691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5399)

					if True == tmp14691 {
						tmp14643 := MakeNative(func(__e *ControlFlow) {
							W5409 := __e.Get(1)
							_ = W5409
							tmp14668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5409)

							if True == tmp14668 {
								tmp14644 := MakeNative(func(__e *ControlFlow) {
									W5416 := __e.Get(1)
									_ = W5416
									tmp14646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5416)

									if True == tmp14646 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										__e.Return(W5416)
										return
									}

								}, 1)

								tmp14647 := MakeNative(func(__e *ControlFlow) {
									W5417 := __e.Get(1)
									_ = W5417
									tmp14664 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5417)

									if True == tmp14664 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp14648 := MakeNative(func(__e *ControlFlow) {
											W5418 := __e.Get(1)
											_ = W5418
											tmp14649 := MakeNative(func(__e *ControlFlow) {
												W5419 := __e.Get(1)
												_ = W5419
												tmp14660 := Call(__e, PrimFunc(symshen_4hds_a_2), W5419, sym_1_6)

												if True == tmp14660 {
													tmp14650 := MakeNative(func(__e *ControlFlow) {
														W5420 := __e.Get(1)
														_ = W5420
														tmp14657 := PrimIsPair(W5420)

														if True == tmp14657 {
															tmp14651 := MakeNative(func(__e *ControlFlow) {
																W5421 := __e.Get(1)
																_ = W5421
																tmp14652 := MakeNative(func(__e *ControlFlow) {
																	W5422 := __e.Get(1)
																	_ = W5422
																	tmp14653 := Call(__e, PrimFunc(sym_8p), W5418, W5421)

																	__e.TailApply(PrimFunc(symshen_4comb), W5422, tmp14653)
																	return

																}, 1)

																tmp14654 := Call(__e, PrimFunc(symtail), W5420)

																__e.TailApply(tmp14652, tmp14654)
																return

															}, 1)

															tmp14655 := Call(__e, PrimFunc(symhead), W5420)

															__e.TailApply(tmp14651, tmp14655)
															return

														} else {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														}

													}, 1)

													tmp14658 := Call(__e, PrimFunc(symtail), W5419)

													__e.TailApply(tmp14650, tmp14658)
													return

												} else {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp14661 := Call(__e, PrimFunc(symshen_4in_1_6), W5417)

											__e.TailApply(tmp14649, tmp14661)
											return

										}, 1)

										tmp14662 := Call(__e, PrimFunc(symshen_4_5_1out), W5417)

										__e.TailApply(tmp14648, tmp14662)
										return

									}

								}, 1)

								tmp14665 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5388)

								tmp14666 := Call(__e, tmp14647, tmp14665)

								__e.TailApply(tmp14644, tmp14666)
								return

							} else {
								__e.Return(W5409)
								return
							}

						}, 1)

						tmp14669 := MakeNative(func(__e *ControlFlow) {
							W5410 := __e.Get(1)
							_ = W5410
							tmp14687 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5410)

							if True == tmp14687 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp14670 := MakeNative(func(__e *ControlFlow) {
									W5411 := __e.Get(1)
									_ = W5411
									tmp14671 := MakeNative(func(__e *ControlFlow) {
										W5412 := __e.Get(1)
										_ = W5412
										tmp14683 := Call(__e, PrimFunc(symshen_4hds_a_2), W5412, sym_5_1)

										if True == tmp14683 {
											tmp14672 := MakeNative(func(__e *ControlFlow) {
												W5413 := __e.Get(1)
												_ = W5413
												tmp14680 := PrimIsPair(W5413)

												if True == tmp14680 {
													tmp14673 := MakeNative(func(__e *ControlFlow) {
														W5414 := __e.Get(1)
														_ = W5414
														tmp14674 := MakeNative(func(__e *ControlFlow) {
															W5415 := __e.Get(1)
															_ = W5415
															tmp14675 := Call(__e, PrimFunc(symshen_4correct), W5414)

															tmp14676 := Call(__e, PrimFunc(sym_8p), W5411, tmp14675)

															__e.TailApply(PrimFunc(symshen_4comb), W5415, tmp14676)
															return

														}, 1)

														tmp14677 := Call(__e, PrimFunc(symtail), W5413)

														__e.TailApply(tmp14674, tmp14677)
														return

													}, 1)

													tmp14678 := Call(__e, PrimFunc(symhead), W5413)

													__e.TailApply(tmp14673, tmp14678)
													return

												} else {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp14681 := Call(__e, PrimFunc(symtail), W5412)

											__e.TailApply(tmp14672, tmp14681)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp14684 := Call(__e, PrimFunc(symshen_4in_1_6), W5410)

									__e.TailApply(tmp14671, tmp14684)
									return

								}, 1)

								tmp14685 := Call(__e, PrimFunc(symshen_4_5_1out), W5410)

								__e.TailApply(tmp14670, tmp14685)
								return

							}

						}, 1)

						tmp14688 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5388)

						tmp14689 := Call(__e, tmp14669, tmp14688)

						__e.TailApply(tmp14643, tmp14689)
						return

					} else {
						__e.Return(W5399)
						return
					}

				}, 1)

				tmp14692 := MakeNative(func(__e *ControlFlow) {
					W5400 := __e.Get(1)
					_ = W5400
					tmp14723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5400)

					if True == tmp14723 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp14693 := MakeNative(func(__e *ControlFlow) {
							W5401 := __e.Get(1)
							_ = W5401
							tmp14694 := MakeNative(func(__e *ControlFlow) {
								W5402 := __e.Get(1)
								_ = W5402
								tmp14719 := Call(__e, PrimFunc(symshen_4hds_a_2), W5402, sym_5_1)

								if True == tmp14719 {
									tmp14695 := MakeNative(func(__e *ControlFlow) {
										W5403 := __e.Get(1)
										_ = W5403
										tmp14716 := PrimIsPair(W5403)

										if True == tmp14716 {
											tmp14696 := MakeNative(func(__e *ControlFlow) {
												W5404 := __e.Get(1)
												_ = W5404
												tmp14697 := MakeNative(func(__e *ControlFlow) {
													W5405 := __e.Get(1)
													_ = W5405
													tmp14712 := Call(__e, PrimFunc(symshen_4hds_a_2), W5405, symwhere)

													if True == tmp14712 {
														tmp14698 := MakeNative(func(__e *ControlFlow) {
															W5406 := __e.Get(1)
															_ = W5406
															tmp14709 := PrimIsPair(W5406)

															if True == tmp14709 {
																tmp14699 := MakeNative(func(__e *ControlFlow) {
																	W5407 := __e.Get(1)
																	_ = W5407
																	tmp14700 := MakeNative(func(__e *ControlFlow) {
																		W5408 := __e.Get(1)
																		_ = W5408
																		tmp14701 := PrimCons(W5404, Nil)

																		tmp14702 := PrimCons(W5407, tmp14701)

																		tmp14703 := PrimCons(symwhere, tmp14702)

																		tmp14704 := Call(__e, PrimFunc(symshen_4correct), tmp14703)

																		tmp14705 := Call(__e, PrimFunc(sym_8p), W5401, tmp14704)

																		__e.TailApply(PrimFunc(symshen_4comb), W5408, tmp14705)
																		return

																	}, 1)

																	tmp14706 := Call(__e, PrimFunc(symtail), W5406)

																	__e.TailApply(tmp14700, tmp14706)
																	return

																}, 1)

																tmp14707 := Call(__e, PrimFunc(symhead), W5406)

																__e.TailApply(tmp14699, tmp14707)
																return

															} else {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															}

														}, 1)

														tmp14710 := Call(__e, PrimFunc(symtail), W5405)

														__e.TailApply(tmp14698, tmp14710)
														return

													} else {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp14713 := Call(__e, PrimFunc(symtail), W5403)

												__e.TailApply(tmp14697, tmp14713)
												return

											}, 1)

											tmp14714 := Call(__e, PrimFunc(symhead), W5403)

											__e.TailApply(tmp14696, tmp14714)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp14717 := Call(__e, PrimFunc(symtail), W5402)

									__e.TailApply(tmp14695, tmp14717)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp14720 := Call(__e, PrimFunc(symshen_4in_1_6), W5400)

							__e.TailApply(tmp14694, tmp14720)
							return

						}, 1)

						tmp14721 := Call(__e, PrimFunc(symshen_4_5_1out), W5400)

						__e.TailApply(tmp14693, tmp14721)
						return

					}

				}, 1)

				tmp14724 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5388)

				tmp14725 := Call(__e, tmp14692, tmp14724)

				__e.TailApply(tmp14642, tmp14725)
				return

			} else {
				__e.Return(W5389)
				return
			}

		}, 1)

		tmp14728 := MakeNative(func(__e *ControlFlow) {
			W5390 := __e.Get(1)
			_ = W5390
			tmp14758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5390)

			if True == tmp14758 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp14729 := MakeNative(func(__e *ControlFlow) {
					W5391 := __e.Get(1)
					_ = W5391
					tmp14730 := MakeNative(func(__e *ControlFlow) {
						W5392 := __e.Get(1)
						_ = W5392
						tmp14754 := Call(__e, PrimFunc(symshen_4hds_a_2), W5392, sym_1_6)

						if True == tmp14754 {
							tmp14731 := MakeNative(func(__e *ControlFlow) {
								W5393 := __e.Get(1)
								_ = W5393
								tmp14751 := PrimIsPair(W5393)

								if True == tmp14751 {
									tmp14732 := MakeNative(func(__e *ControlFlow) {
										W5394 := __e.Get(1)
										_ = W5394
										tmp14733 := MakeNative(func(__e *ControlFlow) {
											W5395 := __e.Get(1)
											_ = W5395
											tmp14747 := Call(__e, PrimFunc(symshen_4hds_a_2), W5395, symwhere)

											if True == tmp14747 {
												tmp14734 := MakeNative(func(__e *ControlFlow) {
													W5396 := __e.Get(1)
													_ = W5396
													tmp14744 := PrimIsPair(W5396)

													if True == tmp14744 {
														tmp14735 := MakeNative(func(__e *ControlFlow) {
															W5397 := __e.Get(1)
															_ = W5397
															tmp14736 := MakeNative(func(__e *ControlFlow) {
																W5398 := __e.Get(1)
																_ = W5398
																tmp14737 := PrimCons(W5394, Nil)

																tmp14738 := PrimCons(W5397, tmp14737)

																tmp14739 := PrimCons(symwhere, tmp14738)

																tmp14740 := Call(__e, PrimFunc(sym_8p), W5391, tmp14739)

																__e.TailApply(PrimFunc(symshen_4comb), W5398, tmp14740)
																return

															}, 1)

															tmp14741 := Call(__e, PrimFunc(symtail), W5396)

															__e.TailApply(tmp14736, tmp14741)
															return

														}, 1)

														tmp14742 := Call(__e, PrimFunc(symhead), W5396)

														__e.TailApply(tmp14735, tmp14742)
														return

													} else {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp14745 := Call(__e, PrimFunc(symtail), W5395)

												__e.TailApply(tmp14734, tmp14745)
												return

											} else {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp14748 := Call(__e, PrimFunc(symtail), W5393)

										__e.TailApply(tmp14733, tmp14748)
										return

									}, 1)

									tmp14749 := Call(__e, PrimFunc(symhead), W5393)

									__e.TailApply(tmp14732, tmp14749)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp14752 := Call(__e, PrimFunc(symtail), W5392)

							__e.TailApply(tmp14731, tmp14752)
							return

						} else {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp14755 := Call(__e, PrimFunc(symshen_4in_1_6), W5390)

					__e.TailApply(tmp14730, tmp14755)
					return

				}, 1)

				tmp14756 := Call(__e, PrimFunc(symshen_4_5_1out), W5390)

				__e.TailApply(tmp14729, tmp14756)
				return

			}

		}, 1)

		tmp14759 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5388)

		tmp14760 := Call(__e, tmp14728, tmp14759)

		__e.TailApply(tmp14641, tmp14760)
		return

	}, 1)

	tmp14761 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp14640)

	_ = tmp14761

	tmp14762 := MakeNative(func(__e *ControlFlow) {
		V5423 := __e.Get(1)
		_ = V5423
		tmp14910 := PrimIsPair(V5423)

		var ifres14854 Obj

		if True == tmp14910 {
			tmp14908 := PrimHead(V5423)

			tmp14909 := PrimEqual(symwhere, tmp14908)

			var ifres14856 Obj

			if True == tmp14909 {
				tmp14906 := PrimTail(V5423)

				tmp14907 := PrimIsPair(tmp14906)

				var ifres14858 Obj

				if True == tmp14907 {
					tmp14903 := PrimTail(V5423)

					tmp14904 := PrimTail(tmp14903)

					tmp14905 := PrimIsPair(tmp14904)

					var ifres14860 Obj

					if True == tmp14905 {
						tmp14899 := PrimTail(V5423)

						tmp14900 := PrimTail(tmp14899)

						tmp14901 := PrimHead(tmp14900)

						tmp14902 := PrimIsPair(tmp14901)

						var ifres14862 Obj

						if True == tmp14902 {
							tmp14894 := PrimTail(V5423)

							tmp14895 := PrimTail(tmp14894)

							tmp14896 := PrimHead(tmp14895)

							tmp14897 := PrimHead(tmp14896)

							tmp14898 := PrimEqual(symfail_1if, tmp14897)

							var ifres14864 Obj

							if True == tmp14898 {
								tmp14889 := PrimTail(V5423)

								tmp14890 := PrimTail(tmp14889)

								tmp14891 := PrimHead(tmp14890)

								tmp14892 := PrimTail(tmp14891)

								tmp14893 := PrimIsPair(tmp14892)

								var ifres14866 Obj

								if True == tmp14893 {
									tmp14883 := PrimTail(V5423)

									tmp14884 := PrimTail(tmp14883)

									tmp14885 := PrimHead(tmp14884)

									tmp14886 := PrimTail(tmp14885)

									tmp14887 := PrimTail(tmp14886)

									tmp14888 := PrimIsPair(tmp14887)

									var ifres14868 Obj

									if True == tmp14888 {
										tmp14876 := PrimTail(V5423)

										tmp14877 := PrimTail(tmp14876)

										tmp14878 := PrimHead(tmp14877)

										tmp14879 := PrimTail(tmp14878)

										tmp14880 := PrimTail(tmp14879)

										tmp14881 := PrimTail(tmp14880)

										tmp14882 := PrimEqual(Nil, tmp14881)

										var ifres14870 Obj

										if True == tmp14882 {
											tmp14872 := PrimTail(V5423)

											tmp14873 := PrimTail(tmp14872)

											tmp14874 := PrimTail(tmp14873)

											tmp14875 := PrimEqual(Nil, tmp14874)

											var ifres14871 Obj

											if True == tmp14875 {
												ifres14871 = True

											} else {
												ifres14871 = False

											}

											ifres14870 = ifres14871

										} else {
											ifres14870 = False

										}

										var ifres14869 Obj

										if True == ifres14870 {
											ifres14869 = True

										} else {
											ifres14869 = False

										}

										ifres14868 = ifres14869

									} else {
										ifres14868 = False

									}

									var ifres14867 Obj

									if True == ifres14868 {
										ifres14867 = True

									} else {
										ifres14867 = False

									}

									ifres14866 = ifres14867

								} else {
									ifres14866 = False

								}

								var ifres14865 Obj

								if True == ifres14866 {
									ifres14865 = True

								} else {
									ifres14865 = False

								}

								ifres14864 = ifres14865

							} else {
								ifres14864 = False

							}

							var ifres14863 Obj

							if True == ifres14864 {
								ifres14863 = True

							} else {
								ifres14863 = False

							}

							ifres14862 = ifres14863

						} else {
							ifres14862 = False

						}

						var ifres14861 Obj

						if True == ifres14862 {
							ifres14861 = True

						} else {
							ifres14861 = False

						}

						ifres14860 = ifres14861

					} else {
						ifres14860 = False

					}

					var ifres14859 Obj

					if True == ifres14860 {
						ifres14859 = True

					} else {
						ifres14859 = False

					}

					ifres14858 = ifres14859

				} else {
					ifres14858 = False

				}

				var ifres14857 Obj

				if True == ifres14858 {
					ifres14857 = True

				} else {
					ifres14857 = False

				}

				ifres14856 = ifres14857

			} else {
				ifres14856 = False

			}

			var ifres14855 Obj

			if True == ifres14856 {
				ifres14855 = True

			} else {
				ifres14855 = False

			}

			ifres14854 = ifres14855

		} else {
			ifres14854 = False

		}

		if True == ifres14854 {
			tmp14763 := PrimTail(V5423)

			tmp14764 := PrimHead(tmp14763)

			tmp14765 := PrimTail(V5423)

			tmp14766 := PrimTail(tmp14765)

			tmp14767 := PrimHead(tmp14766)

			tmp14768 := PrimTail(tmp14767)

			tmp14769 := PrimCons(tmp14768, Nil)

			tmp14770 := PrimCons(symnot, tmp14769)

			tmp14771 := PrimCons(tmp14770, Nil)

			tmp14772 := PrimCons(tmp14764, tmp14771)

			tmp14773 := PrimCons(symand, tmp14772)

			tmp14774 := PrimTail(V5423)

			tmp14775 := PrimTail(tmp14774)

			tmp14776 := PrimHead(tmp14775)

			tmp14777 := PrimTail(tmp14776)

			tmp14778 := PrimTail(tmp14777)

			tmp14779 := PrimCons(tmp14773, tmp14778)

			__e.Return(PrimCons(symwhere, tmp14779))
			return

		} else {
			tmp14852 := PrimIsPair(V5423)

			var ifres14833 Obj

			if True == tmp14852 {
				tmp14850 := PrimHead(V5423)

				tmp14851 := PrimEqual(symwhere, tmp14850)

				var ifres14835 Obj

				if True == tmp14851 {
					tmp14848 := PrimTail(V5423)

					tmp14849 := PrimIsPair(tmp14848)

					var ifres14837 Obj

					if True == tmp14849 {
						tmp14845 := PrimTail(V5423)

						tmp14846 := PrimTail(tmp14845)

						tmp14847 := PrimIsPair(tmp14846)

						var ifres14839 Obj

						if True == tmp14847 {
							tmp14841 := PrimTail(V5423)

							tmp14842 := PrimTail(tmp14841)

							tmp14843 := PrimTail(tmp14842)

							tmp14844 := PrimEqual(Nil, tmp14843)

							var ifres14840 Obj

							if True == tmp14844 {
								ifres14840 = True

							} else {
								ifres14840 = False

							}

							ifres14839 = ifres14840

						} else {
							ifres14839 = False

						}

						var ifres14838 Obj

						if True == ifres14839 {
							ifres14838 = True

						} else {
							ifres14838 = False

						}

						ifres14837 = ifres14838

					} else {
						ifres14837 = False

					}

					var ifres14836 Obj

					if True == ifres14837 {
						ifres14836 = True

					} else {
						ifres14836 = False

					}

					ifres14835 = ifres14836

				} else {
					ifres14835 = False

				}

				var ifres14834 Obj

				if True == ifres14835 {
					ifres14834 = True

				} else {
					ifres14834 = False

				}

				ifres14833 = ifres14834

			} else {
				ifres14833 = False

			}

			if True == ifres14833 {
				tmp14780 := PrimTail(V5423)

				tmp14781 := PrimHead(tmp14780)

				tmp14782 := PrimTail(V5423)

				tmp14783 := PrimTail(tmp14782)

				tmp14784 := PrimHead(tmp14783)

				tmp14785 := PrimCons(symfail, Nil)

				tmp14786 := PrimCons(tmp14785, Nil)

				tmp14787 := PrimCons(tmp14784, tmp14786)

				tmp14788 := PrimCons(sym_a, tmp14787)

				tmp14789 := PrimCons(tmp14788, Nil)

				tmp14790 := PrimCons(symnot, tmp14789)

				tmp14791 := PrimCons(tmp14790, Nil)

				tmp14792 := PrimCons(tmp14781, tmp14791)

				tmp14793 := PrimCons(symand, tmp14792)

				tmp14794 := PrimTail(V5423)

				tmp14795 := PrimTail(tmp14794)

				tmp14796 := PrimCons(tmp14793, tmp14795)

				__e.Return(PrimCons(symwhere, tmp14796))
				return

			} else {
				tmp14831 := PrimIsPair(V5423)

				var ifres14812 Obj

				if True == tmp14831 {
					tmp14829 := PrimHead(V5423)

					tmp14830 := PrimEqual(symfail_1if, tmp14829)

					var ifres14814 Obj

					if True == tmp14830 {
						tmp14827 := PrimTail(V5423)

						tmp14828 := PrimIsPair(tmp14827)

						var ifres14816 Obj

						if True == tmp14828 {
							tmp14824 := PrimTail(V5423)

							tmp14825 := PrimTail(tmp14824)

							tmp14826 := PrimIsPair(tmp14825)

							var ifres14818 Obj

							if True == tmp14826 {
								tmp14820 := PrimTail(V5423)

								tmp14821 := PrimTail(tmp14820)

								tmp14822 := PrimTail(tmp14821)

								tmp14823 := PrimEqual(Nil, tmp14822)

								var ifres14819 Obj

								if True == tmp14823 {
									ifres14819 = True

								} else {
									ifres14819 = False

								}

								ifres14818 = ifres14819

							} else {
								ifres14818 = False

							}

							var ifres14817 Obj

							if True == ifres14818 {
								ifres14817 = True

							} else {
								ifres14817 = False

							}

							ifres14816 = ifres14817

						} else {
							ifres14816 = False

						}

						var ifres14815 Obj

						if True == ifres14816 {
							ifres14815 = True

						} else {
							ifres14815 = False

						}

						ifres14814 = ifres14815

					} else {
						ifres14814 = False

					}

					var ifres14813 Obj

					if True == ifres14814 {
						ifres14813 = True

					} else {
						ifres14813 = False

					}

					ifres14812 = ifres14813

				} else {
					ifres14812 = False

				}

				if True == ifres14812 {
					tmp14797 := PrimTail(V5423)

					tmp14798 := PrimCons(tmp14797, Nil)

					tmp14799 := PrimCons(symnot, tmp14798)

					tmp14800 := PrimTail(V5423)

					tmp14801 := PrimTail(tmp14800)

					tmp14802 := PrimCons(tmp14799, tmp14801)

					__e.Return(PrimCons(symwhere, tmp14802))
					return

				} else {
					tmp14803 := PrimCons(symfail, Nil)

					tmp14804 := PrimCons(tmp14803, Nil)

					tmp14805 := PrimCons(V5423, tmp14804)

					tmp14806 := PrimCons(sym_a, tmp14805)

					tmp14807 := PrimCons(tmp14806, Nil)

					tmp14808 := PrimCons(symnot, tmp14807)

					tmp14809 := PrimCons(V5423, Nil)

					tmp14810 := PrimCons(tmp14808, tmp14809)

					__e.Return(PrimCons(symwhere, tmp14810))
					return

				}

			}

		}

	}, 1)

	tmp14911 := Call(__e, ns2_1set, symshen_4correct, tmp14762)

	_ = tmp14911

	tmp14912 := MakeNative(func(__e *ControlFlow) {
		V5424 := __e.Get(1)
		_ = V5424
		V5425 := __e.Get(2)
		_ = V5425
		V5426 := __e.Get(3)
		_ = V5426
		V5427 := __e.Get(4)
		_ = V5427
		V5428 := __e.Get(5)
		_ = V5428
		V5429 := __e.Get(6)
		_ = V5429
		V5430 := __e.Get(7)
		_ = V5430
		V5431 := __e.Get(8)
		_ = V5431
		tmp14913 := MakeNative(func(__e *ControlFlow) {
			W5432 := __e.Get(1)
			_ = W5432
			tmp14914 := MakeNative(func(__e *ControlFlow) {
				W5433 := __e.Get(1)
				_ = W5433
				tmp14944 := PrimEqual(W5433, False)

				if True == tmp14944 {
					tmp14915 := MakeNative(func(__e *ControlFlow) {
						W5435 := __e.Get(1)
						_ = W5435
						tmp14917 := PrimEqual(W5435, False)

						if True == tmp14917 {
							__e.TailApply(PrimFunc(symshen_4unlock), V5429, W5432)
							return
						} else {
							__e.Return(W5435)
							return
						}

					}, 1)

					tmp14942 := Call(__e, PrimFunc(symshen_4unlocked_2), V5429)

					var ifres14918 Obj

					if True == tmp14942 {
						tmp14919 := MakeNative(func(__e *ControlFlow) {
							W5436 := __e.Get(1)
							_ = W5436
							tmp14939 := PrimIsPair(W5436)

							if True == tmp14939 {
								tmp14920 := MakeNative(func(__e *ControlFlow) {
									W5437 := __e.Get(1)
									_ = W5437
									tmp14921 := MakeNative(func(__e *ControlFlow) {
										W5438 := __e.Get(1)
										_ = W5438
										tmp14922 := MakeNative(func(__e *ControlFlow) {
											W5439 := __e.Get(1)
											_ = W5439
											tmp14923 := Call(__e, PrimFunc(symshen_4incinfs))

											_ = tmp14923

											tmp14924 := Call(__e, PrimFunc(symshen_4deref), W5437, V5428)

											tmp14925 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp14924)

											tmp14926 := MakeNative(func(__e *ControlFlow) {
												tmp14927 := Call(__e, PrimFunc(symshen_4lazyderef), W5439, V5428)

												tmp14928 := Call(__e, PrimFunc(symfst), tmp14927)

												tmp14929 := Call(__e, PrimFunc(symshen_4lazyderef), W5439, V5428)

												tmp14930 := Call(__e, PrimFunc(symsnd), tmp14929)

												tmp14931 := MakeNative(func(__e *ControlFlow) {
													tmp14932 := MakeNative(func(__e *ControlFlow) {
														tmp14933 := PrimNumberAdd(V5427, MakeNumber(1))

														__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5424, W5438, V5426, tmp14933, V5428, V5429, W5432, V5431)
														return

													}, 0)

													__e.TailApply(PrimFunc(symshen_4cut), V5428, V5429, W5432, tmp14932)
													return

												}, 0)

												__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5424, V5427, tmp14928, tmp14930, V5426, V5428, V5429, W5432, tmp14931)
												return

											}, 0)

											tmp14934 := Call(__e, PrimFunc(symbind), W5439, tmp14925, V5428, V5429, W5432, tmp14926)

											__e.TailApply(PrimFunc(symshen_4gc), V5428, tmp14934)
											return

										}, 1)

										tmp14935 := Call(__e, PrimFunc(symshen_4newpv), V5428)

										__e.TailApply(tmp14922, tmp14935)
										return

									}, 1)

									tmp14936 := PrimTail(W5436)

									__e.TailApply(tmp14921, tmp14936)
									return

								}, 1)

								tmp14937 := PrimHead(W5436)

								__e.TailApply(tmp14920, tmp14937)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp14940 := Call(__e, PrimFunc(symshen_4lazyderef), V5425, V5428)

						tmp14941 := Call(__e, tmp14919, tmp14940)

						ifres14918 = tmp14941

					} else {
						ifres14918 = False

					}

					__e.TailApply(tmp14915, ifres14918)
					return

				} else {
					__e.Return(W5433)
					return
				}

			}, 1)

			tmp14952 := Call(__e, PrimFunc(symshen_4unlocked_2), V5429)

			var ifres14945 Obj

			if True == tmp14952 {
				tmp14946 := MakeNative(func(__e *ControlFlow) {
					W5434 := __e.Get(1)
					_ = W5434
					tmp14949 := PrimEqual(W5434, Nil)

					if True == tmp14949 {
						tmp14947 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp14947

						__e.TailApply(PrimFunc(symthaw), V5431)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp14950 := Call(__e, PrimFunc(symshen_4lazyderef), V5425, V5428)

				tmp14951 := Call(__e, tmp14946, tmp14950)

				ifres14945 = tmp14951

			} else {
				ifres14945 = False

			}

			__e.TailApply(tmp14914, ifres14945)
			return

		}, 1)

		tmp14953 := PrimNumberAdd(V5430, MakeNumber(1))

		__e.TailApply(tmp14913, tmp14953)
		return

	}, 8)

	tmp14954 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp14912)

	_ = tmp14954

	tmp14955 := MakeNative(func(__e *ControlFlow) {
		V5440 := __e.Get(1)
		_ = V5440
		tmp14968 := Call(__e, PrimFunc(symtuple_2), V5440)

		if True == tmp14968 {
			tmp14956 := MakeNative(func(__e *ControlFlow) {
				W5441 := __e.Get(1)
				_ = W5441
				tmp14957 := MakeNative(func(__e *ControlFlow) {
					W5442 := __e.Get(1)
					_ = W5442
					tmp14958 := Call(__e, PrimFunc(symfst), V5440)

					tmp14959 := Call(__e, PrimFunc(symshen_4freshen), W5442, tmp14958)

					tmp14960 := Call(__e, PrimFunc(symsnd), V5440)

					tmp14961 := Call(__e, PrimFunc(symshen_4freshen), W5442, tmp14960)

					__e.TailApply(PrimFunc(sym_8p), tmp14959, tmp14961)
					return

				}, 1)

				tmp14962 := MakeNative(func(__e *ControlFlow) {
					Z5443 := __e.Get(1)
					_ = Z5443
					tmp14963 := Call(__e, PrimFunc(symshen_4freshterm), Z5443)

					__e.Return(PrimCons(Z5443, tmp14963))
					return

				}, 1)

				tmp14964 := Call(__e, PrimFunc(symmap), tmp14962, W5441)

				__e.TailApply(tmp14957, tmp14964)
				return

			}, 1)

			tmp14965 := Call(__e, PrimFunc(symfst), V5440)

			tmp14966 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp14965)

			__e.TailApply(tmp14956, tmp14966)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
			return
		}

	}, 1)

	tmp14969 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp14955)

	_ = tmp14969

	tmp14970 := MakeNative(func(__e *ControlFlow) {
		V5444 := __e.Get(1)
		_ = V5444
		V5445 := __e.Get(2)
		_ = V5445
		tmp14984 := PrimEqual(Nil, V5444)

		if True == tmp14984 {
			__e.Return(V5445)
			return
		} else {
			tmp14982 := PrimIsPair(V5444)

			var ifres14978 Obj

			if True == tmp14982 {
				tmp14980 := PrimHead(V5444)

				tmp14981 := PrimIsPair(tmp14980)

				var ifres14979 Obj

				if True == tmp14981 {
					ifres14979 = True

				} else {
					ifres14979 = False

				}

				ifres14978 = ifres14979

			} else {
				ifres14978 = False

			}

			if True == ifres14978 {
				tmp14971 := PrimTail(V5444)

				tmp14972 := PrimHead(V5444)

				tmp14973 := PrimHead(tmp14972)

				tmp14974 := PrimHead(V5444)

				tmp14975 := PrimTail(tmp14974)

				tmp14976 := Call(__e, PrimFunc(symshen_4beta), tmp14973, tmp14975, V5445)

				__e.TailApply(PrimFunc(symshen_4freshen), tmp14971, tmp14976)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
				return
			}

		}

	}, 2)

	tmp14985 := Call(__e, ns2_1set, symshen_4freshen, tmp14970)

	_ = tmp14985

	tmp14986 := MakeNative(func(__e *ControlFlow) {
		V5446 := __e.Get(1)
		_ = V5446
		V5447 := __e.Get(2)
		_ = V5447
		V5448 := __e.Get(3)
		_ = V5448
		V5449 := __e.Get(4)
		_ = V5449
		V5450 := __e.Get(5)
		_ = V5450
		V5451 := __e.Get(6)
		_ = V5451
		V5452 := __e.Get(7)
		_ = V5452
		V5453 := __e.Get(8)
		_ = V5453
		V5454 := __e.Get(9)
		_ = V5454
		tmp14987 := MakeNative(func(__e *ControlFlow) {
			W5455 := __e.Get(1)
			_ = W5455
			tmp15000 := PrimEqual(W5455, False)

			if True == tmp15000 {
				tmp14998 := Call(__e, PrimFunc(symshen_4unlocked_2), V5452)

				if True == tmp14998 {
					tmp14988 := MakeNative(func(__e *ControlFlow) {
						W5456 := __e.Get(1)
						_ = W5456
						tmp14989 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp14989

						tmp14990 := Call(__e, PrimFunc(symshen_4app), V5446, MakeString("\n"), symshen_4a)

						tmp14991 := PrimStringConcat(MakeString(" of "), tmp14990)

						tmp14992 := Call(__e, PrimFunc(symshen_4app), V5447, tmp14991, symshen_4a)

						tmp14993 := PrimStringConcat(MakeString("type error in rule "), tmp14992)

						tmp14994 := PrimSimpleError(tmp14993)

						tmp14995 := Call(__e, PrimFunc(symbind), W5456, tmp14994, V5451, V5452, V5453, V5454)

						__e.TailApply(PrimFunc(symshen_4gc), V5451, tmp14995)
						return

					}, 1)

					tmp14996 := Call(__e, PrimFunc(symshen_4newpv), V5451)

					__e.TailApply(tmp14988, tmp14996)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W5455)
				return
			}

		}, 1)

		tmp15004 := Call(__e, PrimFunc(symshen_4unlocked_2), V5452)

		var ifres15001 Obj

		if True == tmp15004 {
			tmp15002 := Call(__e, PrimFunc(symshen_4incinfs))

			_ = tmp15002

			tmp15003 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5448, V5449, V5450, V5451, V5452, V5453, V5454)

			ifres15001 = tmp15003

		} else {
			ifres15001 = False

		}

		__e.TailApply(tmp14987, ifres15001)
		return

	}, 9)

	tmp15005 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp14986)

	_ = tmp15005

	tmp15006 := MakeNative(func(__e *ControlFlow) {
		V5457 := __e.Get(1)
		_ = V5457
		V5458 := __e.Get(2)
		_ = V5458
		V5459 := __e.Get(3)
		_ = V5459
		V5460 := __e.Get(4)
		_ = V5460
		V5461 := __e.Get(5)
		_ = V5461
		V5462 := __e.Get(6)
		_ = V5462
		V5463 := __e.Get(7)
		_ = V5463
		tmp15007 := MakeNative(func(__e *ControlFlow) {
			W5464 := __e.Get(1)
			_ = W5464
			tmp15008 := MakeNative(func(__e *ControlFlow) {
				W5465 := __e.Get(1)
				_ = W5465
				tmp15031 := PrimEqual(W5465, False)

				if True == tmp15031 {
					tmp15009 := MakeNative(func(__e *ControlFlow) {
						W5472 := __e.Get(1)
						_ = W5472
						tmp15011 := PrimEqual(W5472, False)

						if True == tmp15011 {
							__e.TailApply(PrimFunc(symshen_4unlock), V5461, W5464)
							return
						} else {
							__e.Return(W5472)
							return
						}

					}, 1)

					tmp15029 := Call(__e, PrimFunc(symshen_4unlocked_2), V5461)

					var ifres15012 Obj

					if True == tmp15029 {
						tmp15013 := MakeNative(func(__e *ControlFlow) {
							W5473 := __e.Get(1)
							_ = W5473
							tmp15014 := MakeNative(func(__e *ControlFlow) {
								W5474 := __e.Get(1)
								_ = W5474
								tmp15015 := MakeNative(func(__e *ControlFlow) {
									W5475 := __e.Get(1)
									_ = W5475
									tmp15016 := Call(__e, PrimFunc(symshen_4incinfs))

									_ = tmp15016

									tmp15017 := Call(__e, PrimFunc(symshen_4freshterms), V5457)

									tmp15018 := MakeNative(func(__e *ControlFlow) {
										tmp15019 := MakeNative(func(__e *ControlFlow) {
											tmp15020 := MakeNative(func(__e *ControlFlow) {
												tmp15021 := MakeNative(func(__e *ControlFlow) {
													__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5458, W5474, W5475, V5460, V5461, W5464, V5463)
													return
												}, 0)

												__e.TailApply(PrimFunc(symshen_4myassume), V5457, V5459, W5475, V5460, V5461, W5464, tmp15021)
												return

											}, 0)

											__e.TailApply(PrimFunc(symshen_4cut), V5460, V5461, W5464, tmp15020)
											return

										}, 0)

										__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5457, V5459, W5473, W5474, V5460, V5461, W5464, tmp15019)
										return

									}, 0)

									tmp15022 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp15017, W5473, V5460, V5461, W5464, tmp15018)

									__e.TailApply(PrimFunc(symshen_4gc), V5460, tmp15022)
									return

								}, 1)

								tmp15023 := Call(__e, PrimFunc(symshen_4newpv), V5460)

								tmp15024 := Call(__e, tmp15015, tmp15023)

								__e.TailApply(PrimFunc(symshen_4gc), V5460, tmp15024)
								return

							}, 1)

							tmp15025 := Call(__e, PrimFunc(symshen_4newpv), V5460)

							tmp15026 := Call(__e, tmp15014, tmp15025)

							__e.TailApply(PrimFunc(symshen_4gc), V5460, tmp15026)
							return

						}, 1)

						tmp15027 := Call(__e, PrimFunc(symshen_4newpv), V5460)

						tmp15028 := Call(__e, tmp15013, tmp15027)

						ifres15012 = tmp15028

					} else {
						ifres15012 = False

					}

					__e.TailApply(tmp15009, ifres15012)
					return

				} else {
					__e.Return(W5465)
					return
				}

			}, 1)

			tmp15061 := Call(__e, PrimFunc(symshen_4unlocked_2), V5461)

			var ifres15032 Obj

			if True == tmp15061 {
				tmp15033 := MakeNative(func(__e *ControlFlow) {
					W5466 := __e.Get(1)
					_ = W5466
					tmp15058 := PrimEqual(W5466, Nil)

					if True == tmp15058 {
						tmp15034 := MakeNative(func(__e *ControlFlow) {
							W5467 := __e.Get(1)
							_ = W5467
							tmp15055 := PrimIsPair(W5467)

							if True == tmp15055 {
								tmp15035 := MakeNative(func(__e *ControlFlow) {
									W5468 := __e.Get(1)
									_ = W5468
									tmp15051 := PrimEqual(W5468, sym_1_1_6)

									if True == tmp15051 {
										tmp15036 := MakeNative(func(__e *ControlFlow) {
											W5469 := __e.Get(1)
											_ = W5469
											tmp15047 := PrimIsPair(W5469)

											if True == tmp15047 {
												tmp15037 := MakeNative(func(__e *ControlFlow) {
													W5470 := __e.Get(1)
													_ = W5470
													tmp15038 := MakeNative(func(__e *ControlFlow) {
														W5471 := __e.Get(1)
														_ = W5471
														tmp15042 := PrimEqual(W5471, Nil)

														if True == tmp15042 {
															tmp15039 := Call(__e, PrimFunc(symshen_4incinfs))

															_ = tmp15039

															tmp15040 := MakeNative(func(__e *ControlFlow) {
																__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5458, W5470, Nil, V5460, V5461, W5464, V5463)
																return
															}, 0)

															__e.TailApply(PrimFunc(symshen_4cut), V5460, V5461, W5464, tmp15040)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp15043 := PrimTail(W5469)

													tmp15044 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15043, V5460)

													__e.TailApply(tmp15038, tmp15044)
													return

												}, 1)

												tmp15045 := PrimHead(W5469)

												__e.TailApply(tmp15037, tmp15045)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp15048 := PrimTail(W5467)

										tmp15049 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15048, V5460)

										__e.TailApply(tmp15036, tmp15049)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp15052 := PrimHead(W5467)

								tmp15053 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15052, V5460)

								__e.TailApply(tmp15035, tmp15053)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp15056 := Call(__e, PrimFunc(symshen_4lazyderef), V5459, V5460)

						__e.TailApply(tmp15034, tmp15056)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp15059 := Call(__e, PrimFunc(symshen_4lazyderef), V5457, V5460)

				tmp15060 := Call(__e, tmp15033, tmp15059)

				ifres15032 = tmp15060

			} else {
				ifres15032 = False

			}

			__e.TailApply(tmp15008, ifres15032)
			return

		}, 1)

		tmp15062 := PrimNumberAdd(V5462, MakeNumber(1))

		__e.TailApply(tmp15007, tmp15062)
		return

	}, 7)

	tmp15063 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp15006)

	_ = tmp15063

	tmp15064 := MakeNative(func(__e *ControlFlow) {
		V5476 := __e.Get(1)
		_ = V5476
		V5477 := __e.Get(2)
		_ = V5477
		V5478 := __e.Get(3)
		_ = V5478
		V5479 := __e.Get(4)
		_ = V5479
		V5480 := __e.Get(5)
		_ = V5480
		V5481 := __e.Get(6)
		_ = V5481
		V5482 := __e.Get(7)
		_ = V5482
		tmp15065 := MakeNative(func(__e *ControlFlow) {
			W5483 := __e.Get(1)
			_ = W5483
			tmp15218 := PrimEqual(W5483, False)

			if True == tmp15218 {
				tmp15216 := Call(__e, PrimFunc(symshen_4unlocked_2), V5480)

				if True == tmp15216 {
					tmp15066 := MakeNative(func(__e *ControlFlow) {
						W5487 := __e.Get(1)
						_ = W5487
						tmp15213 := PrimIsPair(W5487)

						if True == tmp15213 {
							tmp15067 := MakeNative(func(__e *ControlFlow) {
								W5488 := __e.Get(1)
								_ = W5488
								tmp15068 := MakeNative(func(__e *ControlFlow) {
									W5489 := __e.Get(1)
									_ = W5489
									tmp15069 := MakeNative(func(__e *ControlFlow) {
										W5490 := __e.Get(1)
										_ = W5490
										tmp15208 := PrimIsPair(W5490)

										if True == tmp15208 {
											tmp15070 := MakeNative(func(__e *ControlFlow) {
												W5491 := __e.Get(1)
												_ = W5491
												tmp15071 := MakeNative(func(__e *ControlFlow) {
													W5492 := __e.Get(1)
													_ = W5492
													tmp15203 := PrimIsPair(W5492)

													if True == tmp15203 {
														tmp15072 := MakeNative(func(__e *ControlFlow) {
															W5493 := __e.Get(1)
															_ = W5493
															tmp15199 := PrimEqual(W5493, sym_1_1_6)

															if True == tmp15199 {
																tmp15073 := MakeNative(func(__e *ControlFlow) {
																	W5494 := __e.Get(1)
																	_ = W5494
																	tmp15195 := PrimIsPair(W5494)

																	if True == tmp15195 {
																		tmp15074 := MakeNative(func(__e *ControlFlow) {
																			W5495 := __e.Get(1)
																			_ = W5495
																			tmp15075 := MakeNative(func(__e *ControlFlow) {
																				W5496 := __e.Get(1)
																				_ = W5496
																				tmp15190 := PrimEqual(W5496, Nil)

																				if True == tmp15190 {
																					tmp15076 := MakeNative(func(__e *ControlFlow) {
																						W5497 := __e.Get(1)
																						_ = W5497
																						tmp15077 := MakeNative(func(__e *ControlFlow) {
																							W5498 := __e.Get(1)
																							_ = W5498
																							tmp15181 := PrimIsPair(W5497)

																							if True == tmp15181 {
																								tmp15078 := MakeNative(func(__e *ControlFlow) {
																									W5503 := __e.Get(1)
																									_ = W5503
																									tmp15079 := MakeNative(func(__e *ControlFlow) {
																										W5504 := __e.Get(1)
																										_ = W5504
																										tmp15149 := PrimIsPair(W5503)

																										if True == tmp15149 {
																											tmp15080 := MakeNative(func(__e *ControlFlow) {
																												W5509 := __e.Get(1)
																												_ = W5509
																												tmp15081 := MakeNative(func(__e *ControlFlow) {
																													W5510 := __e.Get(1)
																													_ = W5510
																													tmp15082 := MakeNative(func(__e *ControlFlow) {
																														W5511 := __e.Get(1)
																														_ = W5511
																														tmp15124 := PrimIsPair(W5510)

																														if True == tmp15124 {
																															tmp15083 := MakeNative(func(__e *ControlFlow) {
																																W5514 := __e.Get(1)
																																_ = W5514
																																tmp15084 := MakeNative(func(__e *ControlFlow) {
																																	W5515 := __e.Get(1)
																																	_ = W5515
																																	tmp15085 := MakeNative(func(__e *ControlFlow) {
																																		W5516 := __e.Get(1)
																																		_ = W5516
																																		tmp15105 := PrimIsPair(W5515)

																																		if True == tmp15105 {
																																			tmp15086 := MakeNative(func(__e *ControlFlow) {
																																				W5518 := __e.Get(1)
																																				_ = W5518
																																				tmp15087 := MakeNative(func(__e *ControlFlow) {
																																					W5519 := __e.Get(1)
																																					_ = W5519
																																					tmp15088 := MakeNative(func(__e *ControlFlow) {
																																						W5520 := __e.Get(1)
																																						_ = W5520
																																						tmp15092 := PrimEqual(W5519, Nil)

																																						if True == tmp15092 {
																																							__e.TailApply(PrimFunc(symthaw), W5520)
																																							return
																																						} else {
																																							tmp15090 := Call(__e, PrimFunc(symshen_4pvar_2), W5519)

																																							if True == tmp15090 {
																																								__e.TailApply(PrimFunc(symshen_4bind_b), W5519, Nil, V5479, W5520)
																																								return
																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp15093 := MakeNative(func(__e *ControlFlow) {
																																						__e.TailApply(W5516, W5518)
																																						return
																																					}, 0)

																																					__e.TailApply(tmp15088, tmp15093)
																																					return

																																				}, 1)

																																				tmp15094 := PrimTail(W5515)

																																				tmp15095 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15094, V5479)

																																				__e.TailApply(tmp15087, tmp15095)
																																				return

																																			}, 1)

																																			tmp15096 := PrimHead(W5515)

																																			__e.TailApply(tmp15086, tmp15096)
																																			return

																																		} else {
																																			tmp15103 := Call(__e, PrimFunc(symshen_4pvar_2), W5515)

																																			if True == tmp15103 {
																																				tmp15097 := MakeNative(func(__e *ControlFlow) {
																																					W5521 := __e.Get(1)
																																					_ = W5521
																																					tmp15098 := PrimCons(W5521, Nil)

																																					tmp15099 := MakeNative(func(__e *ControlFlow) {
																																						__e.TailApply(W5516, W5521)
																																						return
																																					}, 0)

																																					tmp15100 := Call(__e, PrimFunc(symshen_4bind_b), W5515, tmp15098, V5479, tmp15099)

																																					__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15100)
																																					return

																																				}, 1)

																																				tmp15101 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																																				__e.TailApply(tmp15097, tmp15101)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp15106 := MakeNative(func(__e *ControlFlow) {
																																		Z5517 := __e.Get(1)
																																		_ = Z5517
																																		tmp15107 := Call(__e, W5511, W5514)

																																		__e.TailApply(tmp15107, Z5517)
																																		return

																																	}, 1)

																																	__e.TailApply(tmp15085, tmp15106)
																																	return

																																}, 1)

																																tmp15108 := PrimTail(W5510)

																																tmp15109 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15108, V5479)

																																__e.TailApply(tmp15084, tmp15109)
																																return

																															}, 1)

																															tmp15110 := PrimHead(W5510)

																															__e.TailApply(tmp15083, tmp15110)
																															return

																														} else {
																															tmp15122 := Call(__e, PrimFunc(symshen_4pvar_2), W5510)

																															if True == tmp15122 {
																																tmp15111 := MakeNative(func(__e *ControlFlow) {
																																	W5522 := __e.Get(1)
																																	_ = W5522
																																	tmp15112 := MakeNative(func(__e *ControlFlow) {
																																		W5523 := __e.Get(1)
																																		_ = W5523
																																		tmp15113 := PrimCons(W5523, Nil)

																																		tmp15114 := PrimCons(W5522, tmp15113)

																																		tmp15115 := MakeNative(func(__e *ControlFlow) {
																																			tmp15116 := Call(__e, W5511, W5522)

																																			__e.TailApply(tmp15116, W5523)
																																			return

																																		}, 0)

																																		tmp15117 := Call(__e, PrimFunc(symshen_4bind_b), W5510, tmp15114, V5479, tmp15115)

																																		__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15117)
																																		return

																																	}, 1)

																																	tmp15118 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																																	tmp15119 := Call(__e, tmp15112, tmp15118)

																																	__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15119)
																																	return

																																}, 1)

																																tmp15120 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																																__e.TailApply(tmp15111, tmp15120)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp15125 := MakeNative(func(__e *ControlFlow) {
																														Z5512 := __e.Get(1)
																														_ = Z5512
																														__e.Return(MakeNative(func(__e *ControlFlow) {
																															Z5513 := __e.Get(1)
																															_ = Z5513
																															tmp15126 := Call(__e, W5504, W5509)

																															tmp15127 := Call(__e, tmp15126, Z5512)

																															__e.TailApply(tmp15127, Z5513)
																															return

																														}, 1))
																														return
																													}, 1)

																													__e.TailApply(tmp15082, tmp15125)
																													return

																												}, 1)

																												tmp15128 := PrimTail(W5503)

																												tmp15129 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15128, V5479)

																												__e.TailApply(tmp15081, tmp15129)
																												return

																											}, 1)

																											tmp15130 := PrimHead(W5503)

																											__e.TailApply(tmp15080, tmp15130)
																											return

																										} else {
																											tmp15147 := Call(__e, PrimFunc(symshen_4pvar_2), W5503)

																											if True == tmp15147 {
																												tmp15131 := MakeNative(func(__e *ControlFlow) {
																													W5524 := __e.Get(1)
																													_ = W5524
																													tmp15132 := MakeNative(func(__e *ControlFlow) {
																														W5525 := __e.Get(1)
																														_ = W5525
																														tmp15133 := MakeNative(func(__e *ControlFlow) {
																															W5526 := __e.Get(1)
																															_ = W5526
																															tmp15134 := PrimCons(W5526, Nil)

																															tmp15135 := PrimCons(W5525, tmp15134)

																															tmp15136 := PrimCons(W5524, tmp15135)

																															tmp15137 := MakeNative(func(__e *ControlFlow) {
																																tmp15138 := Call(__e, W5504, W5524)

																																tmp15139 := Call(__e, tmp15138, W5525)

																																__e.TailApply(tmp15139, W5526)
																																return

																															}, 0)

																															tmp15140 := Call(__e, PrimFunc(symshen_4bind_b), W5503, tmp15136, V5479, tmp15137)

																															__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15140)
																															return

																														}, 1)

																														tmp15141 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																														tmp15142 := Call(__e, tmp15133, tmp15141)

																														__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15142)
																														return

																													}, 1)

																													tmp15143 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																													tmp15144 := Call(__e, tmp15132, tmp15143)

																													__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15144)
																													return

																												}, 1)

																												tmp15145 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																												__e.TailApply(tmp15131, tmp15145)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}

																									}, 1)

																									tmp15150 := MakeNative(func(__e *ControlFlow) {
																										Z5505 := __e.Get(1)
																										_ = Z5505
																										__e.Return(MakeNative(func(__e *ControlFlow) {
																											Z5506 := __e.Get(1)
																											_ = Z5506
																											__e.Return(MakeNative(func(__e *ControlFlow) {
																												Z5507 := __e.Get(1)
																												_ = Z5507
																												tmp15151 := MakeNative(func(__e *ControlFlow) {
																													W5508 := __e.Get(1)
																													_ = W5508
																													tmp15152 := Call(__e, W5498, Z5505)

																													tmp15153 := Call(__e, tmp15152, Z5506)

																													tmp15154 := Call(__e, tmp15153, Z5507)

																													__e.TailApply(tmp15154, W5508)
																													return

																												}, 1)

																												tmp15155 := PrimTail(W5497)

																												__e.TailApply(tmp15151, tmp15155)
																												return

																											}, 1))
																											return
																										}, 1))
																										return
																									}, 1)

																									__e.TailApply(tmp15079, tmp15150)
																									return

																								}, 1)

																								tmp15156 := PrimHead(W5497)

																								tmp15157 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15156, V5479)

																								__e.TailApply(tmp15078, tmp15157)
																								return

																							} else {
																								tmp15179 := Call(__e, PrimFunc(symshen_4pvar_2), W5497)

																								if True == tmp15179 {
																									tmp15158 := MakeNative(func(__e *ControlFlow) {
																										W5527 := __e.Get(1)
																										_ = W5527
																										tmp15159 := MakeNative(func(__e *ControlFlow) {
																											W5528 := __e.Get(1)
																											_ = W5528
																											tmp15160 := MakeNative(func(__e *ControlFlow) {
																												W5529 := __e.Get(1)
																												_ = W5529
																												tmp15161 := MakeNative(func(__e *ControlFlow) {
																													W5530 := __e.Get(1)
																													_ = W5530
																													tmp15162 := PrimCons(W5529, Nil)

																													tmp15163 := PrimCons(W5528, tmp15162)

																													tmp15164 := PrimCons(W5527, tmp15163)

																													tmp15165 := PrimCons(tmp15164, W5530)

																													tmp15166 := MakeNative(func(__e *ControlFlow) {
																														tmp15167 := Call(__e, W5498, W5527)

																														tmp15168 := Call(__e, tmp15167, W5528)

																														tmp15169 := Call(__e, tmp15168, W5529)

																														__e.TailApply(tmp15169, W5530)
																														return

																													}, 0)

																													tmp15170 := Call(__e, PrimFunc(symshen_4bind_b), W5497, tmp15165, V5479, tmp15166)

																													__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15170)
																													return

																												}, 1)

																												tmp15171 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																												tmp15172 := Call(__e, tmp15161, tmp15171)

																												__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15172)
																												return

																											}, 1)

																											tmp15173 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																											tmp15174 := Call(__e, tmp15160, tmp15173)

																											__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15174)
																											return

																										}, 1)

																										tmp15175 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																										tmp15176 := Call(__e, tmp15159, tmp15175)

																										__e.TailApply(PrimFunc(symshen_4gc), V5479, tmp15176)
																										return

																									}, 1)

																									tmp15177 := Call(__e, PrimFunc(symshen_4newpv), V5479)

																									__e.TailApply(tmp15158, tmp15177)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}

																						}, 1)

																						tmp15182 := MakeNative(func(__e *ControlFlow) {
																							Z5499 := __e.Get(1)
																							_ = Z5499
																							__e.Return(MakeNative(func(__e *ControlFlow) {
																								Z5500 := __e.Get(1)
																								_ = Z5500
																								__e.Return(MakeNative(func(__e *ControlFlow) {
																									Z5501 := __e.Get(1)
																									_ = Z5501
																									__e.Return(MakeNative(func(__e *ControlFlow) {
																										Z5502 := __e.Get(1)
																										_ = Z5502
																										tmp15183 := Call(__e, PrimFunc(symshen_4incinfs))

																										_ = tmp15183

																										tmp15184 := MakeNative(func(__e *ControlFlow) {
																											tmp15185 := MakeNative(func(__e *ControlFlow) {
																												tmp15186 := PrimIntern(MakeString(":"))

																												tmp15187 := MakeNative(func(__e *ControlFlow) {
																													__e.TailApply(PrimFunc(symshen_4myassume), W5489, W5495, Z5502, V5479, V5480, V5481, V5482)
																													return
																												}, 0)

																												__e.TailApply(PrimFunc(symbind), Z5500, tmp15186, V5479, V5480, V5481, tmp15187)
																												return

																											}, 0)

																											__e.TailApply(PrimFunc(symis_b), W5488, Z5499, V5479, V5480, V5481, tmp15185)
																											return

																										}, 0)

																										__e.TailApply(PrimFunc(symis_b), W5491, Z5501, V5479, V5480, V5481, tmp15184)
																										return

																									}, 1))
																									return
																								}, 1))
																								return
																							}, 1))
																							return
																						}, 1)

																						__e.TailApply(tmp15077, tmp15182)
																						return

																					}, 1)

																					tmp15188 := Call(__e, PrimFunc(symshen_4lazyderef), V5478, V5479)

																					__e.TailApply(tmp15076, tmp15188)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp15191 := PrimTail(W5494)

																			tmp15192 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15191, V5479)

																			__e.TailApply(tmp15075, tmp15192)
																			return

																		}, 1)

																		tmp15193 := PrimHead(W5494)

																		__e.TailApply(tmp15074, tmp15193)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp15196 := PrimTail(W5492)

																tmp15197 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15196, V5479)

																__e.TailApply(tmp15073, tmp15197)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp15200 := PrimHead(W5492)

														tmp15201 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15200, V5479)

														__e.TailApply(tmp15072, tmp15201)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp15204 := PrimTail(W5490)

												tmp15205 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15204, V5479)

												__e.TailApply(tmp15071, tmp15205)
												return

											}, 1)

											tmp15206 := PrimHead(W5490)

											__e.TailApply(tmp15070, tmp15206)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp15209 := Call(__e, PrimFunc(symshen_4lazyderef), V5477, V5479)

									__e.TailApply(tmp15069, tmp15209)
									return

								}, 1)

								tmp15210 := PrimTail(W5487)

								__e.TailApply(tmp15068, tmp15210)
								return

							}, 1)

							tmp15211 := PrimHead(W5487)

							__e.TailApply(tmp15067, tmp15211)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp15214 := Call(__e, PrimFunc(symshen_4lazyderef), V5476, V5479)

					__e.TailApply(tmp15066, tmp15214)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W5483)
				return
			}

		}, 1)

		tmp15234 := Call(__e, PrimFunc(symshen_4unlocked_2), V5480)

		var ifres15219 Obj

		if True == tmp15234 {
			tmp15220 := MakeNative(func(__e *ControlFlow) {
				W5484 := __e.Get(1)
				_ = W5484
				tmp15231 := PrimEqual(W5484, Nil)

				if True == tmp15231 {
					tmp15221 := MakeNative(func(__e *ControlFlow) {
						W5485 := __e.Get(1)
						_ = W5485
						tmp15222 := MakeNative(func(__e *ControlFlow) {
							W5486 := __e.Get(1)
							_ = W5486
							tmp15226 := PrimEqual(W5485, Nil)

							if True == tmp15226 {
								__e.TailApply(PrimFunc(symthaw), W5486)
								return
							} else {
								tmp15224 := Call(__e, PrimFunc(symshen_4pvar_2), W5485)

								if True == tmp15224 {
									__e.TailApply(PrimFunc(symshen_4bind_b), W5485, Nil, V5479, W5486)
									return
								} else {
									__e.Return(False)
									return
								}

							}

						}, 1)

						tmp15227 := MakeNative(func(__e *ControlFlow) {
							tmp15228 := Call(__e, PrimFunc(symshen_4incinfs))

							_ = tmp15228

							__e.TailApply(PrimFunc(symthaw), V5482)
							return

						}, 0)

						__e.TailApply(tmp15222, tmp15227)
						return

					}, 1)

					tmp15229 := Call(__e, PrimFunc(symshen_4lazyderef), V5478, V5479)

					__e.TailApply(tmp15221, tmp15229)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp15232 := Call(__e, PrimFunc(symshen_4lazyderef), V5476, V5479)

			tmp15233 := Call(__e, tmp15220, tmp15232)

			ifres15219 = tmp15233

		} else {
			ifres15219 = False

		}

		__e.TailApply(tmp15065, ifres15219)
		return

	}, 7)

	tmp15235 := Call(__e, ns2_1set, symshen_4myassume, tmp15064)

	_ = tmp15235

	tmp15236 := MakeNative(func(__e *ControlFlow) {
		V5533 := __e.Get(1)
		_ = V5533
		tmp15259 := PrimEqual(Nil, V5533)

		if True == tmp15259 {
			__e.Return(Nil)
			return
		} else {
			tmp15257 := PrimIsPair(V5533)

			var ifres15253 Obj

			if True == tmp15257 {
				tmp15255 := PrimHead(V5533)

				tmp15256 := PrimIsPair(tmp15255)

				var ifres15254 Obj

				if True == tmp15256 {
					ifres15254 = True

				} else {
					ifres15254 = False

				}

				ifres15253 = ifres15254

			} else {
				ifres15253 = False

			}

			if True == ifres15253 {
				tmp15237 := PrimHead(V5533)

				tmp15238 := PrimTail(V5533)

				tmp15239 := Call(__e, PrimFunc(symappend), tmp15237, tmp15238)

				__e.TailApply(PrimFunc(symshen_4freshterms), tmp15239)
				return

			} else {
				tmp15251 := PrimIsPair(V5533)

				var ifres15247 Obj

				if True == tmp15251 {
					tmp15249 := PrimHead(V5533)

					tmp15250 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp15249)

					var ifres15248 Obj

					if True == tmp15250 {
						ifres15248 = True

					} else {
						ifres15248 = False

					}

					ifres15247 = ifres15248

				} else {
					ifres15247 = False

				}

				if True == ifres15247 {
					tmp15240 := PrimHead(V5533)

					tmp15241 := PrimTail(V5533)

					tmp15242 := Call(__e, PrimFunc(symshen_4freshterms), tmp15241)

					__e.TailApply(PrimFunc(symadjoin), tmp15240, tmp15242)
					return

				} else {
					tmp15245 := PrimIsPair(V5533)

					if True == tmp15245 {
						tmp15243 := PrimTail(V5533)

						__e.TailApply(PrimFunc(symshen_4freshterms), tmp15243)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshterms)
						return
					}

				}

			}

		}

	}, 1)

	tmp15260 := Call(__e, ns2_1set, symshen_4freshterms, tmp15236)

	_ = tmp15260

	tmp15261 := MakeNative(func(__e *ControlFlow) {
		V5534 := __e.Get(1)
		_ = V5534
		V5535 := __e.Get(2)
		_ = V5535
		V5536 := __e.Get(3)
		_ = V5536
		V5537 := __e.Get(4)
		_ = V5537
		V5538 := __e.Get(5)
		_ = V5538
		V5539 := __e.Get(6)
		_ = V5539
		tmp15262 := MakeNative(func(__e *ControlFlow) {
			W5540 := __e.Get(1)
			_ = W5540
			tmp15386 := PrimEqual(W5540, False)

			if True == tmp15386 {
				tmp15384 := Call(__e, PrimFunc(symshen_4unlocked_2), V5537)

				if True == tmp15384 {
					tmp15263 := MakeNative(func(__e *ControlFlow) {
						W5544 := __e.Get(1)
						_ = W5544
						tmp15381 := PrimIsPair(W5544)

						if True == tmp15381 {
							tmp15264 := MakeNative(func(__e *ControlFlow) {
								W5545 := __e.Get(1)
								_ = W5545
								tmp15265 := MakeNative(func(__e *ControlFlow) {
									W5546 := __e.Get(1)
									_ = W5546
									tmp15266 := MakeNative(func(__e *ControlFlow) {
										W5547 := __e.Get(1)
										_ = W5547
										tmp15267 := MakeNative(func(__e *ControlFlow) {
											W5548 := __e.Get(1)
											_ = W5548
											tmp15371 := PrimIsPair(W5547)

											if True == tmp15371 {
												tmp15268 := MakeNative(func(__e *ControlFlow) {
													W5553 := __e.Get(1)
													_ = W5553
													tmp15269 := MakeNative(func(__e *ControlFlow) {
														W5554 := __e.Get(1)
														_ = W5554
														tmp15339 := PrimIsPair(W5553)

														if True == tmp15339 {
															tmp15270 := MakeNative(func(__e *ControlFlow) {
																W5559 := __e.Get(1)
																_ = W5559
																tmp15271 := MakeNative(func(__e *ControlFlow) {
																	W5560 := __e.Get(1)
																	_ = W5560
																	tmp15272 := MakeNative(func(__e *ControlFlow) {
																		W5561 := __e.Get(1)
																		_ = W5561
																		tmp15314 := PrimIsPair(W5560)

																		if True == tmp15314 {
																			tmp15273 := MakeNative(func(__e *ControlFlow) {
																				W5564 := __e.Get(1)
																				_ = W5564
																				tmp15274 := MakeNative(func(__e *ControlFlow) {
																					W5565 := __e.Get(1)
																					_ = W5565
																					tmp15275 := MakeNative(func(__e *ControlFlow) {
																						W5566 := __e.Get(1)
																						_ = W5566
																						tmp15295 := PrimIsPair(W5565)

																						if True == tmp15295 {
																							tmp15276 := MakeNative(func(__e *ControlFlow) {
																								W5568 := __e.Get(1)
																								_ = W5568
																								tmp15277 := MakeNative(func(__e *ControlFlow) {
																									W5569 := __e.Get(1)
																									_ = W5569
																									tmp15278 := MakeNative(func(__e *ControlFlow) {
																										W5570 := __e.Get(1)
																										_ = W5570
																										tmp15282 := PrimEqual(W5569, Nil)

																										if True == tmp15282 {
																											__e.TailApply(PrimFunc(symthaw), W5570)
																											return
																										} else {
																											tmp15280 := Call(__e, PrimFunc(symshen_4pvar_2), W5569)

																											if True == tmp15280 {
																												__e.TailApply(PrimFunc(symshen_4bind_b), W5569, Nil, V5536, W5570)
																												return
																											} else {
																												__e.Return(False)
																												return
																											}

																										}

																									}, 1)

																									tmp15283 := MakeNative(func(__e *ControlFlow) {
																										__e.TailApply(W5566, W5568)
																										return
																									}, 0)

																									__e.TailApply(tmp15278, tmp15283)
																									return

																								}, 1)

																								tmp15284 := PrimTail(W5565)

																								tmp15285 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15284, V5536)

																								__e.TailApply(tmp15277, tmp15285)
																								return

																							}, 1)

																							tmp15286 := PrimHead(W5565)

																							__e.TailApply(tmp15276, tmp15286)
																							return

																						} else {
																							tmp15293 := Call(__e, PrimFunc(symshen_4pvar_2), W5565)

																							if True == tmp15293 {
																								tmp15287 := MakeNative(func(__e *ControlFlow) {
																									W5571 := __e.Get(1)
																									_ = W5571
																									tmp15288 := PrimCons(W5571, Nil)

																									tmp15289 := MakeNative(func(__e *ControlFlow) {
																										__e.TailApply(W5566, W5571)
																										return
																									}, 0)

																									tmp15290 := Call(__e, PrimFunc(symshen_4bind_b), W5565, tmp15288, V5536, tmp15289)

																									__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15290)
																									return

																								}, 1)

																								tmp15291 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																								__e.TailApply(tmp15287, tmp15291)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}

																					}, 1)

																					tmp15296 := MakeNative(func(__e *ControlFlow) {
																						Z5567 := __e.Get(1)
																						_ = Z5567
																						tmp15297 := Call(__e, W5561, W5564)

																						__e.TailApply(tmp15297, Z5567)
																						return

																					}, 1)

																					__e.TailApply(tmp15275, tmp15296)
																					return

																				}, 1)

																				tmp15298 := PrimTail(W5560)

																				tmp15299 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15298, V5536)

																				__e.TailApply(tmp15274, tmp15299)
																				return

																			}, 1)

																			tmp15300 := PrimHead(W5560)

																			__e.TailApply(tmp15273, tmp15300)
																			return

																		} else {
																			tmp15312 := Call(__e, PrimFunc(symshen_4pvar_2), W5560)

																			if True == tmp15312 {
																				tmp15301 := MakeNative(func(__e *ControlFlow) {
																					W5572 := __e.Get(1)
																					_ = W5572
																					tmp15302 := MakeNative(func(__e *ControlFlow) {
																						W5573 := __e.Get(1)
																						_ = W5573
																						tmp15303 := PrimCons(W5573, Nil)

																						tmp15304 := PrimCons(W5572, tmp15303)

																						tmp15305 := MakeNative(func(__e *ControlFlow) {
																							tmp15306 := Call(__e, W5561, W5572)

																							__e.TailApply(tmp15306, W5573)
																							return

																						}, 0)

																						tmp15307 := Call(__e, PrimFunc(symshen_4bind_b), W5560, tmp15304, V5536, tmp15305)

																						__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15307)
																						return

																					}, 1)

																					tmp15308 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																					tmp15309 := Call(__e, tmp15302, tmp15308)

																					__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15309)
																					return

																				}, 1)

																				tmp15310 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																				__e.TailApply(tmp15301, tmp15310)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp15315 := MakeNative(func(__e *ControlFlow) {
																		Z5562 := __e.Get(1)
																		_ = Z5562
																		__e.Return(MakeNative(func(__e *ControlFlow) {
																			Z5563 := __e.Get(1)
																			_ = Z5563
																			tmp15316 := Call(__e, W5554, W5559)

																			tmp15317 := Call(__e, tmp15316, Z5562)

																			__e.TailApply(tmp15317, Z5563)
																			return

																		}, 1))
																		return
																	}, 1)

																	__e.TailApply(tmp15272, tmp15315)
																	return

																}, 1)

																tmp15318 := PrimTail(W5553)

																tmp15319 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15318, V5536)

																__e.TailApply(tmp15271, tmp15319)
																return

															}, 1)

															tmp15320 := PrimHead(W5553)

															__e.TailApply(tmp15270, tmp15320)
															return

														} else {
															tmp15337 := Call(__e, PrimFunc(symshen_4pvar_2), W5553)

															if True == tmp15337 {
																tmp15321 := MakeNative(func(__e *ControlFlow) {
																	W5574 := __e.Get(1)
																	_ = W5574
																	tmp15322 := MakeNative(func(__e *ControlFlow) {
																		W5575 := __e.Get(1)
																		_ = W5575
																		tmp15323 := MakeNative(func(__e *ControlFlow) {
																			W5576 := __e.Get(1)
																			_ = W5576
																			tmp15324 := PrimCons(W5576, Nil)

																			tmp15325 := PrimCons(W5575, tmp15324)

																			tmp15326 := PrimCons(W5574, tmp15325)

																			tmp15327 := MakeNative(func(__e *ControlFlow) {
																				tmp15328 := Call(__e, W5554, W5574)

																				tmp15329 := Call(__e, tmp15328, W5575)

																				__e.TailApply(tmp15329, W5576)
																				return

																			}, 0)

																			tmp15330 := Call(__e, PrimFunc(symshen_4bind_b), W5553, tmp15326, V5536, tmp15327)

																			__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15330)
																			return

																		}, 1)

																		tmp15331 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																		tmp15332 := Call(__e, tmp15323, tmp15331)

																		__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15332)
																		return

																	}, 1)

																	tmp15333 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																	tmp15334 := Call(__e, tmp15322, tmp15333)

																	__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15334)
																	return

																}, 1)

																tmp15335 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																__e.TailApply(tmp15321, tmp15335)
																return

															} else {
																__e.Return(False)
																return
															}

														}

													}, 1)

													tmp15340 := MakeNative(func(__e *ControlFlow) {
														Z5555 := __e.Get(1)
														_ = Z5555
														__e.Return(MakeNative(func(__e *ControlFlow) {
															Z5556 := __e.Get(1)
															_ = Z5556
															__e.Return(MakeNative(func(__e *ControlFlow) {
																Z5557 := __e.Get(1)
																_ = Z5557
																tmp15341 := MakeNative(func(__e *ControlFlow) {
																	W5558 := __e.Get(1)
																	_ = W5558
																	tmp15342 := Call(__e, W5548, Z5555)

																	tmp15343 := Call(__e, tmp15342, Z5556)

																	tmp15344 := Call(__e, tmp15343, Z5557)

																	__e.TailApply(tmp15344, W5558)
																	return

																}, 1)

																tmp15345 := PrimTail(W5547)

																__e.TailApply(tmp15341, tmp15345)
																return

															}, 1))
															return
														}, 1))
														return
													}, 1)

													__e.TailApply(tmp15269, tmp15340)
													return

												}, 1)

												tmp15346 := PrimHead(W5547)

												tmp15347 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15346, V5536)

												__e.TailApply(tmp15268, tmp15347)
												return

											} else {
												tmp15369 := Call(__e, PrimFunc(symshen_4pvar_2), W5547)

												if True == tmp15369 {
													tmp15348 := MakeNative(func(__e *ControlFlow) {
														W5577 := __e.Get(1)
														_ = W5577
														tmp15349 := MakeNative(func(__e *ControlFlow) {
															W5578 := __e.Get(1)
															_ = W5578
															tmp15350 := MakeNative(func(__e *ControlFlow) {
																W5579 := __e.Get(1)
																_ = W5579
																tmp15351 := MakeNative(func(__e *ControlFlow) {
																	W5580 := __e.Get(1)
																	_ = W5580
																	tmp15352 := PrimCons(W5579, Nil)

																	tmp15353 := PrimCons(W5578, tmp15352)

																	tmp15354 := PrimCons(W5577, tmp15353)

																	tmp15355 := PrimCons(tmp15354, W5580)

																	tmp15356 := MakeNative(func(__e *ControlFlow) {
																		tmp15357 := Call(__e, W5548, W5577)

																		tmp15358 := Call(__e, tmp15357, W5578)

																		tmp15359 := Call(__e, tmp15358, W5579)

																		__e.TailApply(tmp15359, W5580)
																		return

																	}, 0)

																	tmp15360 := Call(__e, PrimFunc(symshen_4bind_b), W5547, tmp15355, V5536, tmp15356)

																	__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15360)
																	return

																}, 1)

																tmp15361 := Call(__e, PrimFunc(symshen_4newpv), V5536)

																tmp15362 := Call(__e, tmp15351, tmp15361)

																__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15362)
																return

															}, 1)

															tmp15363 := Call(__e, PrimFunc(symshen_4newpv), V5536)

															tmp15364 := Call(__e, tmp15350, tmp15363)

															__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15364)
															return

														}, 1)

														tmp15365 := Call(__e, PrimFunc(symshen_4newpv), V5536)

														tmp15366 := Call(__e, tmp15349, tmp15365)

														__e.TailApply(PrimFunc(symshen_4gc), V5536, tmp15366)
														return

													}, 1)

													tmp15367 := Call(__e, PrimFunc(symshen_4newpv), V5536)

													__e.TailApply(tmp15348, tmp15367)
													return

												} else {
													__e.Return(False)
													return
												}

											}

										}, 1)

										tmp15372 := MakeNative(func(__e *ControlFlow) {
											Z5549 := __e.Get(1)
											_ = Z5549
											__e.Return(MakeNative(func(__e *ControlFlow) {
												Z5550 := __e.Get(1)
												_ = Z5550
												__e.Return(MakeNative(func(__e *ControlFlow) {
													Z5551 := __e.Get(1)
													_ = Z5551
													__e.Return(MakeNative(func(__e *ControlFlow) {
														Z5552 := __e.Get(1)
														_ = Z5552
														tmp15373 := Call(__e, PrimFunc(symshen_4incinfs))

														_ = tmp15373

														tmp15374 := MakeNative(func(__e *ControlFlow) {
															tmp15375 := PrimIntern(MakeString(":"))

															tmp15376 := MakeNative(func(__e *ControlFlow) {
																__e.TailApply(PrimFunc(symshen_4p_1hyps), W5546, Z5552, V5536, V5537, V5538, V5539)
																return
															}, 0)

															__e.TailApply(PrimFunc(symbind), Z5550, tmp15375, V5536, V5537, V5538, tmp15376)
															return

														}, 0)

														__e.TailApply(PrimFunc(symbind), Z5549, W5545, V5536, V5537, V5538, tmp15374)
														return

													}, 1))
													return
												}, 1))
												return
											}, 1))
											return
										}, 1)

										__e.TailApply(tmp15267, tmp15372)
										return

									}, 1)

									tmp15377 := Call(__e, PrimFunc(symshen_4lazyderef), V5535, V5536)

									__e.TailApply(tmp15266, tmp15377)
									return

								}, 1)

								tmp15378 := PrimTail(W5544)

								__e.TailApply(tmp15265, tmp15378)
								return

							}, 1)

							tmp15379 := PrimHead(W5544)

							__e.TailApply(tmp15264, tmp15379)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp15382 := Call(__e, PrimFunc(symshen_4lazyderef), V5534, V5536)

					__e.TailApply(tmp15263, tmp15382)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W5540)
				return
			}

		}, 1)

		tmp15402 := Call(__e, PrimFunc(symshen_4unlocked_2), V5537)

		var ifres15387 Obj

		if True == tmp15402 {
			tmp15388 := MakeNative(func(__e *ControlFlow) {
				W5541 := __e.Get(1)
				_ = W5541
				tmp15399 := PrimEqual(W5541, Nil)

				if True == tmp15399 {
					tmp15389 := MakeNative(func(__e *ControlFlow) {
						W5542 := __e.Get(1)
						_ = W5542
						tmp15390 := MakeNative(func(__e *ControlFlow) {
							W5543 := __e.Get(1)
							_ = W5543
							tmp15394 := PrimEqual(W5542, Nil)

							if True == tmp15394 {
								__e.TailApply(PrimFunc(symthaw), W5543)
								return
							} else {
								tmp15392 := Call(__e, PrimFunc(symshen_4pvar_2), W5542)

								if True == tmp15392 {
									__e.TailApply(PrimFunc(symshen_4bind_b), W5542, Nil, V5536, W5543)
									return
								} else {
									__e.Return(False)
									return
								}

							}

						}, 1)

						tmp15395 := MakeNative(func(__e *ControlFlow) {
							tmp15396 := Call(__e, PrimFunc(symshen_4incinfs))

							_ = tmp15396

							__e.TailApply(PrimFunc(symthaw), V5539)
							return

						}, 0)

						__e.TailApply(tmp15390, tmp15395)
						return

					}, 1)

					tmp15397 := Call(__e, PrimFunc(symshen_4lazyderef), V5535, V5536)

					__e.TailApply(tmp15389, tmp15397)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp15400 := Call(__e, PrimFunc(symshen_4lazyderef), V5534, V5536)

			tmp15401 := Call(__e, tmp15388, tmp15400)

			ifres15387 = tmp15401

		} else {
			ifres15387 = False

		}

		__e.TailApply(tmp15262, ifres15387)
		return

	}, 6)

	tmp15403 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp15261)

	_ = tmp15403

	tmp15404 := MakeNative(func(__e *ControlFlow) {
		V5581 := __e.Get(1)
		_ = V5581
		V5582 := __e.Get(2)
		_ = V5582
		V5583 := __e.Get(3)
		_ = V5583
		V5584 := __e.Get(4)
		_ = V5584
		V5585 := __e.Get(5)
		_ = V5585
		V5586 := __e.Get(6)
		_ = V5586
		V5587 := __e.Get(7)
		_ = V5587
		tmp15405 := MakeNative(func(__e *ControlFlow) {
			W5588 := __e.Get(1)
			_ = W5588
			tmp15406 := MakeNative(func(__e *ControlFlow) {
				W5589 := __e.Get(1)
				_ = W5589
				tmp15416 := PrimEqual(W5589, False)

				if True == tmp15416 {
					tmp15407 := MakeNative(func(__e *ControlFlow) {
						W5598 := __e.Get(1)
						_ = W5598
						tmp15409 := PrimEqual(W5598, False)

						if True == tmp15409 {
							__e.TailApply(PrimFunc(symshen_4unlock), V5585, W5588)
							return
						} else {
							__e.Return(W5598)
							return
						}

					}, 1)

					tmp15414 := Call(__e, PrimFunc(symshen_4unlocked_2), V5585)

					var ifres15410 Obj

					if True == tmp15414 {
						tmp15411 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp15411

						tmp15412 := Call(__e, PrimFunc(symshen_4curry), V5581)

						tmp15413 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp15412, V5582, V5583, V5584, V5585, W5588, V5587)

						ifres15410 = tmp15413

					} else {
						ifres15410 = False

					}

					__e.TailApply(tmp15407, ifres15410)
					return

				} else {
					__e.Return(W5589)
					return
				}

			}, 1)

			tmp15461 := Call(__e, PrimFunc(symshen_4unlocked_2), V5585)

			var ifres15417 Obj

			if True == tmp15461 {
				tmp15418 := MakeNative(func(__e *ControlFlow) {
					W5590 := __e.Get(1)
					_ = W5590
					tmp15458 := PrimIsPair(W5590)

					if True == tmp15458 {
						tmp15419 := MakeNative(func(__e *ControlFlow) {
							W5591 := __e.Get(1)
							_ = W5591
							tmp15454 := PrimEqual(W5591, symwhere)

							if True == tmp15454 {
								tmp15420 := MakeNative(func(__e *ControlFlow) {
									W5592 := __e.Get(1)
									_ = W5592
									tmp15450 := PrimIsPair(W5592)

									if True == tmp15450 {
										tmp15421 := MakeNative(func(__e *ControlFlow) {
											W5593 := __e.Get(1)
											_ = W5593
											tmp15422 := MakeNative(func(__e *ControlFlow) {
												W5594 := __e.Get(1)
												_ = W5594
												tmp15445 := PrimIsPair(W5594)

												if True == tmp15445 {
													tmp15423 := MakeNative(func(__e *ControlFlow) {
														W5595 := __e.Get(1)
														_ = W5595
														tmp15424 := MakeNative(func(__e *ControlFlow) {
															W5596 := __e.Get(1)
															_ = W5596
															tmp15440 := PrimEqual(W5596, Nil)

															if True == tmp15440 {
																tmp15425 := MakeNative(func(__e *ControlFlow) {
																	W5597 := __e.Get(1)
																	_ = W5597
																	tmp15426 := Call(__e, PrimFunc(symshen_4incinfs))

																	_ = tmp15426

																	tmp15427 := MakeNative(func(__e *ControlFlow) {
																		tmp15428 := Call(__e, PrimFunc(symshen_4curry), W5593)

																		tmp15429 := MakeNative(func(__e *ControlFlow) {
																			tmp15430 := MakeNative(func(__e *ControlFlow) {
																				tmp15431 := MakeNative(func(__e *ControlFlow) {
																					tmp15432 := PrimIntern(MakeString(":"))

																					tmp15433 := PrimCons(symverified, Nil)

																					tmp15434 := PrimCons(tmp15432, tmp15433)

																					tmp15435 := PrimCons(W5597, tmp15434)

																					tmp15436 := PrimCons(tmp15435, V5583)

																					__e.TailApply(PrimFunc(symshen_4t_d_1correct), W5595, V5582, tmp15436, V5584, V5585, W5588, V5587)
																					return

																				}, 0)

																				__e.TailApply(PrimFunc(symshen_4cut), V5584, V5585, W5588, tmp15431)
																				return

																			}, 0)

																			__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5597, symboolean, V5583, V5584, V5585, W5588, tmp15430)
																			return

																		}, 0)

																		__e.TailApply(PrimFunc(symbind), W5597, tmp15428, V5584, V5585, W5588, tmp15429)
																		return

																	}, 0)

																	tmp15437 := Call(__e, PrimFunc(symshen_4cut), V5584, V5585, W5588, tmp15427)

																	__e.TailApply(PrimFunc(symshen_4gc), V5584, tmp15437)
																	return

																}, 1)

																tmp15438 := Call(__e, PrimFunc(symshen_4newpv), V5584)

																__e.TailApply(tmp15425, tmp15438)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp15441 := PrimTail(W5594)

														tmp15442 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15441, V5584)

														__e.TailApply(tmp15424, tmp15442)
														return

													}, 1)

													tmp15443 := PrimHead(W5594)

													__e.TailApply(tmp15423, tmp15443)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp15446 := PrimTail(W5592)

											tmp15447 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15446, V5584)

											__e.TailApply(tmp15422, tmp15447)
											return

										}, 1)

										tmp15448 := PrimHead(W5592)

										__e.TailApply(tmp15421, tmp15448)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp15451 := PrimTail(W5590)

								tmp15452 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15451, V5584)

								__e.TailApply(tmp15420, tmp15452)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp15455 := PrimHead(W5590)

						tmp15456 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15455, V5584)

						__e.TailApply(tmp15419, tmp15456)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp15459 := Call(__e, PrimFunc(symshen_4lazyderef), V5581, V5584)

				tmp15460 := Call(__e, tmp15418, tmp15459)

				ifres15417 = tmp15460

			} else {
				ifres15417 = False

			}

			__e.TailApply(tmp15406, ifres15417)
			return

		}, 1)

		tmp15462 := PrimNumberAdd(V5586, MakeNumber(1))

		__e.TailApply(tmp15405, tmp15462)
		return

	}, 7)

	tmp15463 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp15404)

	_ = tmp15463

	tmp15464 := MakeNative(func(__e *ControlFlow) {
		V5599 := __e.Get(1)
		_ = V5599
		V5600 := __e.Get(2)
		_ = V5600
		V5601 := __e.Get(3)
		_ = V5601
		V5602 := __e.Get(4)
		_ = V5602
		V5603 := __e.Get(5)
		_ = V5603
		V5604 := __e.Get(6)
		_ = V5604
		V5605 := __e.Get(7)
		_ = V5605
		V5606 := __e.Get(8)
		_ = V5606
		tmp15465 := MakeNative(func(__e *ControlFlow) {
			W5607 := __e.Get(1)
			_ = W5607
			tmp15507 := PrimEqual(W5607, False)

			if True == tmp15507 {
				tmp15505 := Call(__e, PrimFunc(symshen_4unlocked_2), V5604)

				if True == tmp15505 {
					tmp15466 := MakeNative(func(__e *ControlFlow) {
						W5609 := __e.Get(1)
						_ = W5609
						tmp15502 := PrimIsPair(W5609)

						if True == tmp15502 {
							tmp15467 := MakeNative(func(__e *ControlFlow) {
								W5610 := __e.Get(1)
								_ = W5610
								tmp15468 := MakeNative(func(__e *ControlFlow) {
									W5611 := __e.Get(1)
									_ = W5611
									tmp15469 := MakeNative(func(__e *ControlFlow) {
										W5612 := __e.Get(1)
										_ = W5612
										tmp15497 := PrimIsPair(W5612)

										if True == tmp15497 {
											tmp15470 := MakeNative(func(__e *ControlFlow) {
												W5613 := __e.Get(1)
												_ = W5613
												tmp15471 := MakeNative(func(__e *ControlFlow) {
													W5614 := __e.Get(1)
													_ = W5614
													tmp15492 := PrimIsPair(W5614)

													if True == tmp15492 {
														tmp15472 := MakeNative(func(__e *ControlFlow) {
															W5615 := __e.Get(1)
															_ = W5615
															tmp15488 := PrimEqual(W5615, sym_1_1_6)

															if True == tmp15488 {
																tmp15473 := MakeNative(func(__e *ControlFlow) {
																	W5616 := __e.Get(1)
																	_ = W5616
																	tmp15484 := PrimIsPair(W5616)

																	if True == tmp15484 {
																		tmp15474 := MakeNative(func(__e *ControlFlow) {
																			W5617 := __e.Get(1)
																			_ = W5617
																			tmp15475 := MakeNative(func(__e *ControlFlow) {
																				W5618 := __e.Get(1)
																				_ = W5618
																				tmp15479 := PrimEqual(W5618, Nil)

																				if True == tmp15479 {
																					tmp15476 := Call(__e, PrimFunc(symshen_4incinfs))

																					_ = tmp15476

																					tmp15477 := MakeNative(func(__e *ControlFlow) {
																						__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W5611, W5617, V5601, V5602, V5603, V5604, V5605, V5606)
																						return
																					}, 0)

																					__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5610, W5613, V5601, V5603, V5604, V5605, tmp15477)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp15480 := PrimTail(W5616)

																			tmp15481 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15480, V5603)

																			__e.TailApply(tmp15475, tmp15481)
																			return

																		}, 1)

																		tmp15482 := PrimHead(W5616)

																		__e.TailApply(tmp15474, tmp15482)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp15485 := PrimTail(W5614)

																tmp15486 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15485, V5603)

																__e.TailApply(tmp15473, tmp15486)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp15489 := PrimHead(W5614)

														tmp15490 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15489, V5603)

														__e.TailApply(tmp15472, tmp15490)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp15493 := PrimTail(W5612)

												tmp15494 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15493, V5603)

												__e.TailApply(tmp15471, tmp15494)
												return

											}, 1)

											tmp15495 := PrimHead(W5612)

											__e.TailApply(tmp15470, tmp15495)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp15498 := Call(__e, PrimFunc(symshen_4lazyderef), V5600, V5603)

									__e.TailApply(tmp15469, tmp15498)
									return

								}, 1)

								tmp15499 := PrimTail(W5609)

								__e.TailApply(tmp15468, tmp15499)
								return

							}, 1)

							tmp15500 := PrimHead(W5609)

							__e.TailApply(tmp15467, tmp15500)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp15503 := Call(__e, PrimFunc(symshen_4lazyderef), V5599, V5603)

					__e.TailApply(tmp15466, tmp15503)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W5607)
				return
			}

		}, 1)

		tmp15515 := Call(__e, PrimFunc(symshen_4unlocked_2), V5604)

		var ifres15508 Obj

		if True == tmp15515 {
			tmp15509 := MakeNative(func(__e *ControlFlow) {
				W5608 := __e.Get(1)
				_ = W5608
				tmp15512 := PrimEqual(W5608, Nil)

				if True == tmp15512 {
					tmp15510 := Call(__e, PrimFunc(symshen_4incinfs))

					_ = tmp15510

					__e.TailApply(PrimFunc(symis_b), V5600, V5602, V5603, V5604, V5605, V5606)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp15513 := Call(__e, PrimFunc(symshen_4lazyderef), V5599, V5603)

			tmp15514 := Call(__e, tmp15509, tmp15513)

			ifres15508 = tmp15514

		} else {
			ifres15508 = False

		}

		__e.TailApply(tmp15465, ifres15508)
		return

	}, 8)

	tmp15516 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp15464)

	_ = tmp15516

	tmp15517 := MakeNative(func(__e *ControlFlow) {
		V5619 := __e.Get(1)
		_ = V5619
		tmp15526 := PrimIsVector(V5619)

		if True == tmp15526 {
			tmp15523 := PrimIsString(V5619)

			tmp15524 := PrimNot(tmp15523)

			var ifres15519 Obj

			if True == tmp15524 {
				tmp15521 := PrimVectorGet(V5619, MakeNumber(0))

				tmp15522 := PrimEqual(tmp15521, symshen_4print_1freshterm)

				var ifres15520 Obj

				if True == tmp15522 {
					ifres15520 = True

				} else {
					ifres15520 = False

				}

				ifres15519 = ifres15520

			} else {
				ifres15519 = False

			}

			if True == ifres15519 {
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

	__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp15517)
	return

}, 0)
