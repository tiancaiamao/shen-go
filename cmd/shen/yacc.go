package main

import . "github.com/tiancaiamao/shen-go/cora"

var YaccMain = MakeNative(func(__e *ControlFlow) {
	tmp15254 := MakeNative(func(__e *ControlFlow) {
		V1141 := __e.Get(1)
		_ = V1141
		V1142 := __e.Get(2)
		_ = V1142
		tmp15255 := MakeNative(func(__e *ControlFlow) {
			Compile := __e.Get(1)
			_ = Compile
			tmp15257 := Call(__e, PrimNS2Value(symshen_4parsed_2), Compile)

			if True == tmp15257 {
				__e.TailApply(PrimNS2Value(symshen_4objectcode), Compile)
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("parse failure\n")))
				return
			}

		}, 1)

		tmp15258 := PrimCons(symshen_4no_1action, Nil)

		tmp15259 := PrimCons(V1142, tmp15258)

		tmp15260 := Call(__e, V1141, tmp15259)

		__e.TailApply(tmp15255, tmp15260)
		return

	}, 2)

	tmp15261 := Call(__e, PrimNS2Value(symdef), symcompile, tmp15254)

	_ = tmp15261

	tmp15262 := MakeNative(func(__e *ControlFlow) {
		V1147 := __e.Get(1)
		_ = V1147
		tmp15275 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), V1147)

		if True == tmp15275 {
			__e.Return(False)
			return
		} else {
			tmp15274 := PrimIsPair(V1147)

			var ifres15270 Obj

			if True == tmp15274 {
				tmp15272 := PrimHead(V1147)

				tmp15273 := PrimIsPair(tmp15272)

				var ifres15271 Obj

				if True == tmp15273 {
					ifres15271 = True

				} else {
					ifres15271 = False

				}

				ifres15270 = ifres15271

			} else {
				ifres15270 = False

			}

			if True == ifres15270 {
				tmp15265 := PrimHead(V1147)

				tmp15266 := PrimNS3Set(symshen_4_dresidue_d, tmp15265)

				_ = tmp15266

				tmp15267 := PrimHead(V1147)

				tmp15268 := Call(__e, PrimNS2Value(symshen_4app), tmp15267, MakeString("\n ..."), symshen_4r)

				tmp15269 := PrimStringConcat(MakeString("syntax error here: "), tmp15268)

				__e.Return(PrimSimpleError(tmp15269))
				return

			} else {
				__e.Return(True)
				return
			}

		}

	}, 1)

	tmp15276 := Call(__e, PrimNS2Value(symdef), symshen_4parsed_2, tmp15262)

	_ = tmp15276

	tmp15277 := MakeNative(func(__e *ControlFlow) {
		V1148 := __e.Get(1)
		_ = V1148
		tmp15278 := Call(__e, PrimNS2Value(symfail))

		__e.Return(PrimEqual(V1148, tmp15278))
		return

	}, 1)

	tmp15279 := Call(__e, PrimNS2Value(symdef), symshen_4parse_1failure_2, tmp15277)

	_ = tmp15279

	tmp15280 := MakeNative(func(__e *ControlFlow) {
		V1151 := __e.Get(1)
		_ = V1151
		tmp15293 := PrimIsPair(V1151)

		var ifres15284 Obj

		if True == tmp15293 {
			tmp15291 := PrimTail(V1151)

			tmp15292 := PrimIsPair(tmp15291)

			var ifres15286 Obj

			if True == tmp15292 {
				tmp15288 := PrimTail(V1151)

				tmp15289 := PrimTail(tmp15288)

				tmp15290 := PrimEqual(Nil, tmp15289)

				var ifres15287 Obj

				if True == tmp15290 {
					ifres15287 = True

				} else {
					ifres15287 = False

				}

				ifres15286 = ifres15287

			} else {
				ifres15286 = False

			}

			var ifres15285 Obj

			if True == ifres15286 {
				ifres15285 = True

			} else {
				ifres15285 = False

			}

			ifres15284 = ifres15285

		} else {
			ifres15284 = False

		}

		if True == ifres15284 {
			tmp15283 := PrimTail(V1151)

			__e.Return(PrimHead(tmp15283))
			return

		} else {
			tmp15282 := Call(__e, PrimNS2Value(symshen_4app), V1151, MakeString(" is not a YACC stream\n"), symshen_4s)

			__e.Return(PrimSimpleError(tmp15282))
			return

		}

	}, 1)

	tmp15294 := Call(__e, PrimNS2Value(symdef), symshen_4objectcode, tmp15280)

	_ = tmp15294

	tmp15295 := MakeNative(func(__e *ControlFlow) {
		V1152 := __e.Get(1)
		_ = V1152
		tmp15296 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_5yacc_6), X)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symcompile), tmp15296, V1152)
		return

	}, 1)

	tmp15297 := Call(__e, PrimNS2Value(symdef), symshen_4yacc_1_6shen, tmp15295)

	_ = tmp15297

	tmp15298 := MakeNative(func(__e *ControlFlow) {
		V1153 := __e.Get(1)
		_ = V1153
		tmp15299 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15301 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15301 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15334 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1153)

		var ifres15302 Obj

		if True == tmp15334 {
			tmp15304 := MakeNative(func(__e *ControlFlow) {
				F := __e.Get(1)
				_ = F
				tmp15305 := MakeNative(func(__e *ControlFlow) {
					News1034 := __e.Get(1)
					_ = News1034
					tmp15306 := MakeNative(func(__e *ControlFlow) {
						Parseshen_4_5yaccsig_6 := __e.Get(1)
						_ = Parseshen_4_5yaccsig_6
						tmp15329 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5yaccsig_6)

						if True == tmp15329 {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						} else {
							tmp15308 := MakeNative(func(__e *ControlFlow) {
								Parseshen_4_5c_1rules_6 := __e.Get(1)
								_ = Parseshen_4_5c_1rules_6
								tmp15327 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5c_1rules_6)

								if True == tmp15327 {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								} else {
									tmp15310 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5c_1rules_6)

									tmp15311 := MakeNative(func(__e *ControlFlow) {
										Stream := __e.Get(1)
										_ = Stream
										tmp15312 := MakeNative(func(__e *ControlFlow) {
											Def := __e.Get(1)
											_ = Def
											__e.Return(Def)
											return
										}, 1)

										tmp15313 := PrimCons(F, Nil)

										tmp15314 := PrimCons(symdefine, tmp15313)

										tmp15315 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5yaccsig_6)

										tmp15316 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5yaccsig_6)

										tmp15317 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5c_1rules_6)

										tmp15318 := Call(__e, PrimNS2Value(symshen_4c_1rules_1_6shen), tmp15316, Stream, tmp15317)

										tmp15319 := PrimCons(tmp15318, Nil)

										tmp15320 := PrimCons(sym_1_6, tmp15319)

										tmp15321 := PrimCons(Stream, tmp15320)

										tmp15322 := Call(__e, PrimNS2Value(symappend), tmp15315, tmp15321)

										tmp15323 := Call(__e, PrimNS2Value(symappend), tmp15314, tmp15322)

										__e.TailApply(tmp15312, tmp15323)
										return

									}, 1)

									tmp15324 := Call(__e, PrimNS2Value(symprotect), symS)

									tmp15325 := Call(__e, PrimNS2Value(symgensym), tmp15324)

									tmp15326 := Call(__e, tmp15311, tmp15325)

									__e.TailApply(PrimNS2Value(symshen_4comb), tmp15310, tmp15326)
									return

								}

							}, 1)

							tmp15328 := Call(__e, PrimNS2Value(symshen_4_5c_1rules_6), Parseshen_4_5yaccsig_6)

							__e.TailApply(tmp15308, tmp15328)
							return

						}

					}, 1)

					tmp15330 := Call(__e, PrimNS2Value(symshen_4_5yaccsig_6), News1034)

					__e.TailApply(tmp15306, tmp15330)
					return

				}, 1)

				tmp15331 := Call(__e, PrimNS2Value(symshen_4tls), V1153)

				__e.TailApply(tmp15305, tmp15331)
				return

			}, 1)

			tmp15332 := Call(__e, PrimNS2Value(symshen_4hds), V1153)

			tmp15333 := Call(__e, tmp15304, tmp15332)

			ifres15302 = tmp15333

		} else {
			tmp15303 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres15302 = tmp15303

		}

		__e.TailApply(tmp15299, ifres15302)
		return

	}, 1)

	tmp15335 := Call(__e, PrimNS2Value(symdef), symshen_4_5yacc_6, tmp15298)

	_ = tmp15335

	tmp15336 := MakeNative(func(__e *ControlFlow) {
		V1154 := __e.Get(1)
		_ = V1154
		tmp15337 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15348 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15348 {
				tmp15339 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15341 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15341 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15342 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp15345 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp15345 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15344 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15344, Nil)
						return

					}

				}, 1)

				tmp15346 := Call(__e, PrimNS2Value(sym_5e_6), V1154)

				tmp15347 := Call(__e, tmp15342, tmp15346)

				__e.TailApply(tmp15339, tmp15347)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15419 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1154)

		var ifres15349 Obj

		if True == tmp15419 {
			tmp15351 := MakeNative(func(__e *ControlFlow) {
				LC := __e.Get(1)
				_ = LC
				tmp15352 := MakeNative(func(__e *ControlFlow) {
					News1036 := __e.Get(1)
					_ = News1036
					tmp15415 := Call(__e, PrimNS2Value(symshen_4ccons_2), News1036)

					if True == tmp15415 {
						tmp15354 := MakeNative(func(__e *ControlFlow) {
							SynCons := __e.Get(1)
							_ = SynCons
							tmp15411 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, symlist)

							if True == tmp15411 {
								tmp15356 := MakeNative(func(__e *ControlFlow) {
									News1037 := __e.Get(1)
									_ = News1037
									tmp15409 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1037)

									if True == tmp15409 {
										tmp15358 := MakeNative(func(__e *ControlFlow) {
											A := __e.Get(1)
											_ = A
											tmp15359 := MakeNative(func(__e *ControlFlow) {
												News1038 := __e.Get(1)
												_ = News1038
												tmp15360 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5end_6 := __e.Get(1)
													_ = Parseshen_4_5end_6
													tmp15405 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

													if True == tmp15405 {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													} else {
														tmp15403 := Call(__e, PrimNS2Value(symshen_4tlstream), News1036)

														tmp15404 := Call(__e, PrimNS2Value(symshen_4_ahd_2), tmp15403, sym_a_a_6)

														if True == tmp15404 {
															tmp15363 := MakeNative(func(__e *ControlFlow) {
																News1039 := __e.Get(1)
																_ = News1039
																tmp15400 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1039)

																if True == tmp15400 {
																	tmp15365 := MakeNative(func(__e *ControlFlow) {
																		B := __e.Get(1)
																		_ = B
																		tmp15366 := MakeNative(func(__e *ControlFlow) {
																			News1040 := __e.Get(1)
																			_ = News1040
																			tmp15397 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1040)

																			if True == tmp15397 {
																				tmp15368 := MakeNative(func(__e *ControlFlow) {
																					RC := __e.Get(1)
																					_ = RC
																					tmp15369 := MakeNative(func(__e *ControlFlow) {
																						News1041 := __e.Get(1)
																						_ = News1041
																						tmp15394 := PrimEqual(sym_i, LC)

																						var ifres15391 Obj

																						if True == tmp15394 {
																							tmp15393 := PrimEqual(sym_j, RC)

																							var ifres15392 Obj

																							if True == tmp15393 {
																								ifres15392 = True

																							} else {
																								ifres15392 = False

																							}

																							ifres15391 = ifres15392

																						} else {
																							ifres15391 = False

																						}

																						if True == ifres15391 {
																							tmp15371 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1041)

																							tmp15372 := MakeNative(func(__e *ControlFlow) {
																								C := __e.Get(1)
																								_ = C
																								tmp15373 := PrimCons(A, Nil)

																								tmp15374 := PrimCons(symlist, tmp15373)

																								tmp15375 := Call(__e, PrimNS2Value(symprotect), C)

																								tmp15376 := PrimCons(tmp15375, Nil)

																								tmp15377 := PrimCons(tmp15374, tmp15376)

																								tmp15378 := PrimCons(symstr, tmp15377)

																								tmp15379 := PrimCons(A, Nil)

																								tmp15380 := PrimCons(symlist, tmp15379)

																								tmp15381 := PrimCons(B, Nil)

																								tmp15382 := PrimCons(tmp15380, tmp15381)

																								tmp15383 := PrimCons(symstr, tmp15382)

																								tmp15384 := PrimCons(sym_j, Nil)

																								tmp15385 := PrimCons(tmp15383, tmp15384)

																								tmp15386 := PrimCons(sym_1_1_6, tmp15385)

																								tmp15387 := PrimCons(tmp15378, tmp15386)

																								__e.Return(PrimCons(sym_i, tmp15387))
																								return

																							}, 1)

																							tmp15388 := Call(__e, PrimNS2Value(symgensym), symC)

																							tmp15389 := Call(__e, PrimNS2Value(symprotect), tmp15388)

																							tmp15390 := Call(__e, tmp15372, tmp15389)

																							__e.TailApply(PrimNS2Value(symshen_4comb), tmp15371, tmp15390)
																							return

																						} else {
																							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																							return
																						}

																					}, 1)

																					tmp15395 := Call(__e, PrimNS2Value(symshen_4tls), News1040)

																					__e.TailApply(tmp15369, tmp15395)
																					return

																				}, 1)

																				tmp15396 := Call(__e, PrimNS2Value(symshen_4hds), News1040)

																				__e.TailApply(tmp15368, tmp15396)
																				return

																			} else {
																				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																				return
																			}

																		}, 1)

																		tmp15398 := Call(__e, PrimNS2Value(symshen_4tls), News1039)

																		__e.TailApply(tmp15366, tmp15398)
																		return

																	}, 1)

																	tmp15399 := Call(__e, PrimNS2Value(symshen_4hds), News1039)

																	__e.TailApply(tmp15365, tmp15399)
																	return

																} else {
																	__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																	return
																}

															}, 1)

															tmp15401 := Call(__e, PrimNS2Value(symshen_4tlstream), News1036)

															tmp15402 := Call(__e, PrimNS2Value(symshen_4tls), tmp15401)

															__e.TailApply(tmp15363, tmp15402)
															return

														} else {
															__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
															return
														}

													}

												}, 1)

												tmp15406 := Call(__e, PrimNS2Value(symshen_4_5end_6), News1038)

												__e.TailApply(tmp15360, tmp15406)
												return

											}, 1)

											tmp15407 := Call(__e, PrimNS2Value(symshen_4tls), News1037)

											__e.TailApply(tmp15359, tmp15407)
											return

										}, 1)

										tmp15408 := Call(__e, PrimNS2Value(symshen_4hds), News1037)

										__e.TailApply(tmp15358, tmp15408)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp15410 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

								__e.TailApply(tmp15356, tmp15410)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp15412 := Call(__e, PrimNS2Value(symshen_4hds), News1036)

						tmp15413 := Call(__e, PrimNS2Value(symshen_4_5_1out), News1036)

						tmp15414 := Call(__e, PrimNS2Value(symshen_4comb), tmp15412, tmp15413)

						__e.TailApply(tmp15354, tmp15414)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15416 := Call(__e, PrimNS2Value(symshen_4tls), V1154)

				__e.TailApply(tmp15352, tmp15416)
				return

			}, 1)

			tmp15417 := Call(__e, PrimNS2Value(symshen_4hds), V1154)

			tmp15418 := Call(__e, tmp15351, tmp15417)

			ifres15349 = tmp15418

		} else {
			tmp15350 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres15349 = tmp15350

		}

		__e.TailApply(tmp15337, ifres15349)
		return

	}, 1)

	tmp15420 := Call(__e, PrimNS2Value(symdef), symshen_4_5yaccsig_6, tmp15336)

	_ = tmp15420

	tmp15421 := MakeNative(func(__e *ControlFlow) {
		V1155 := __e.Get(1)
		_ = V1155
		tmp15422 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15440 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15440 {
				tmp15424 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15426 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15426 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15427 := MakeNative(func(__e *ControlFlow) {
					Parse_5_b_6 := __e.Get(1)
					_ = Parse_5_b_6
					tmp15437 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5_b_6)

					if True == tmp15437 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15429 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5_b_6)

						tmp15435 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

						tmp15436 := Call(__e, PrimNS2Value(symempty_2), tmp15435)

						var ifres15430 Obj

						if True == tmp15436 {
							ifres15430 = Nil

						} else {
							tmp15431 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

							tmp15432 := Call(__e, PrimNS2Value(symshen_4app), tmp15431, MakeString("\n ..."), symshen_4r)

							tmp15433 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp15432)

							tmp15434 := PrimSimpleError(tmp15433)

							ifres15430 = tmp15434

						}

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15429, ifres15430)
						return

					}

				}, 1)

				tmp15438 := Call(__e, PrimNS2Value(sym_5_b_6), V1155)

				tmp15439 := Call(__e, tmp15427, tmp15438)

				__e.TailApply(tmp15424, tmp15439)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15441 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5c_1rule_6 := __e.Get(1)
			_ = Parseshen_4_5c_1rule_6
			tmp15451 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5c_1rule_6)

			if True == tmp15451 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15443 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5c_1rules_6 := __e.Get(1)
					_ = Parseshen_4_5c_1rules_6
					tmp15449 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5c_1rules_6)

					if True == tmp15449 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15445 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5c_1rules_6)

						tmp15446 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5c_1rule_6)

						tmp15447 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5c_1rules_6)

						tmp15448 := PrimCons(tmp15446, tmp15447)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15445, tmp15448)
						return

					}

				}, 1)

				tmp15450 := Call(__e, PrimNS2Value(symshen_4_5c_1rules_6), Parseshen_4_5c_1rule_6)

				__e.TailApply(tmp15443, tmp15450)
				return

			}

		}, 1)

		tmp15452 := Call(__e, PrimNS2Value(symshen_4_5c_1rule_6), V1155)

		tmp15453 := Call(__e, tmp15441, tmp15452)

		__e.TailApply(tmp15422, tmp15453)
		return

	}, 1)

	tmp15454 := Call(__e, PrimNS2Value(symdef), symshen_4_5c_1rules_6, tmp15421)

	_ = tmp15454

	tmp15455 := MakeNative(func(__e *ControlFlow) {
		V1156 := __e.Get(1)
		_ = V1156
		tmp15456 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15476 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15476 {
				tmp15458 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15460 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15460 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15461 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5syntax_6 := __e.Get(1)
					_ = Parseshen_4_5syntax_6
					tmp15473 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5syntax_6)

					if True == tmp15473 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15463 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5sc_6 := __e.Get(1)
							_ = Parseshen_4_5sc_6
							tmp15471 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

							if True == tmp15471 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp15465 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sc_6)

								tmp15466 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5syntax_6)

								tmp15467 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5syntax_6)

								tmp15468 := Call(__e, PrimNS2Value(symshen_4autocomplete), tmp15467)

								tmp15469 := PrimCons(tmp15468, Nil)

								tmp15470 := PrimCons(tmp15466, tmp15469)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp15465, tmp15470)
								return

							}

						}, 1)

						tmp15472 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5syntax_6)

						__e.TailApply(tmp15463, tmp15472)
						return

					}

				}, 1)

				tmp15474 := Call(__e, PrimNS2Value(symshen_4_5syntax_6), V1156)

				tmp15475 := Call(__e, tmp15461, tmp15474)

				__e.TailApply(tmp15458, tmp15475)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15477 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5syntax_6 := __e.Get(1)
			_ = Parseshen_4_5syntax_6
			tmp15492 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5syntax_6)

			if True == tmp15492 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15479 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5semantics_6 := __e.Get(1)
					_ = Parseshen_4_5semantics_6
					tmp15490 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5semantics_6)

					if True == tmp15490 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15481 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5sc_6 := __e.Get(1)
							_ = Parseshen_4_5sc_6
							tmp15488 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sc_6)

							if True == tmp15488 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp15483 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sc_6)

								tmp15484 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5syntax_6)

								tmp15485 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5semantics_6)

								tmp15486 := PrimCons(tmp15485, Nil)

								tmp15487 := PrimCons(tmp15484, tmp15486)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp15483, tmp15487)
								return

							}

						}, 1)

						tmp15489 := Call(__e, PrimNS2Value(symshen_4_5sc_6), Parseshen_4_5semantics_6)

						__e.TailApply(tmp15481, tmp15489)
						return

					}

				}, 1)

				tmp15491 := Call(__e, PrimNS2Value(symshen_4_5semantics_6), Parseshen_4_5syntax_6)

				__e.TailApply(tmp15479, tmp15491)
				return

			}

		}, 1)

		tmp15493 := Call(__e, PrimNS2Value(symshen_4_5syntax_6), V1156)

		tmp15494 := Call(__e, tmp15477, tmp15493)

		__e.TailApply(tmp15456, tmp15494)
		return

	}, 1)

	tmp15495 := Call(__e, PrimNS2Value(symdef), symshen_4_5c_1rule_6, tmp15455)

	_ = tmp15495

	tmp15496 := MakeNative(func(__e *ControlFlow) {
		V1157 := __e.Get(1)
		_ = V1157
		tmp15525 := PrimIsPair(V1157)

		var ifres15517 Obj

		if True == tmp15525 {
			tmp15523 := PrimTail(V1157)

			tmp15524 := PrimEqual(Nil, tmp15523)

			var ifres15519 Obj

			if True == tmp15524 {
				tmp15521 := PrimHead(V1157)

				tmp15522 := Call(__e, PrimNS2Value(symshen_4non_1terminal_2), tmp15521)

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

			var ifres15518 Obj

			if True == ifres15519 {
				ifres15518 = True

			} else {
				ifres15518 = False

			}

			ifres15517 = ifres15518

		} else {
			ifres15517 = False

		}

		if True == ifres15517 {
			__e.Return(PrimHead(V1157))
			return
		} else {
			tmp15516 := PrimIsPair(V1157)

			var ifres15512 Obj

			if True == tmp15516 {
				tmp15514 := PrimHead(V1157)

				tmp15515 := Call(__e, PrimNS2Value(symshen_4non_1terminal_2), tmp15514)

				var ifres15513 Obj

				if True == tmp15515 {
					ifres15513 = True

				} else {
					ifres15513 = False

				}

				ifres15512 = ifres15513

			} else {
				ifres15512 = False

			}

			if True == ifres15512 {
				tmp15507 := PrimHead(V1157)

				tmp15508 := PrimTail(V1157)

				tmp15509 := Call(__e, PrimNS2Value(symshen_4autocomplete), tmp15508)

				tmp15510 := PrimCons(tmp15509, Nil)

				tmp15511 := PrimCons(tmp15507, tmp15510)

				__e.Return(PrimCons(symappend, tmp15511))
				return

			} else {
				tmp15506 := PrimIsPair(V1157)

				if True == tmp15506 {
					tmp15500 := PrimHead(V1157)

					tmp15501 := Call(__e, PrimNS2Value(symshen_4autocomplete), tmp15500)

					tmp15502 := PrimTail(V1157)

					tmp15503 := Call(__e, PrimNS2Value(symshen_4autocomplete), tmp15502)

					tmp15504 := PrimCons(tmp15503, Nil)

					tmp15505 := PrimCons(tmp15501, tmp15504)

					__e.Return(PrimCons(symcons, tmp15505))
					return

				} else {
					__e.Return(V1157)
					return
				}

			}

		}

	}, 1)

	tmp15526 := Call(__e, PrimNS2Value(symdef), symshen_4autocomplete, tmp15496)

	_ = tmp15526

	tmp15527 := MakeNative(func(__e *ControlFlow) {
		V1158 := __e.Get(1)
		_ = V1158
		tmp15534 := PrimIsSymbol(V1158)

		if True == tmp15534 {
			tmp15530 := MakeNative(func(__e *ControlFlow) {
				Explode := __e.Get(1)
				_ = Explode
				tmp15531 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4_5non_1terminal_2_6), X)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symcompile), tmp15531, Explode)
				return

			}, 1)

			tmp15532 := Call(__e, PrimNS2Value(symexplode), V1158)

			tmp15533 := Call(__e, tmp15530, tmp15532)

			if True == tmp15533 {
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

	tmp15535 := Call(__e, PrimNS2Value(symdef), symshen_4non_1terminal_2, tmp15527)

	_ = tmp15535

	tmp15536 := MakeNative(func(__e *ControlFlow) {
		V1159 := __e.Get(1)
		_ = V1159
		tmp15537 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15557 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15557 {
				tmp15539 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15550 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15550 {
						tmp15541 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp15543 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp15543 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp15544 := MakeNative(func(__e *ControlFlow) {
							Parse_5_b_6 := __e.Get(1)
							_ = Parse_5_b_6
							tmp15547 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5_b_6)

							if True == tmp15547 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp15546 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5_b_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp15546, False)
								return

							}

						}, 1)

						tmp15548 := Call(__e, PrimNS2Value(sym_5_b_6), V1159)

						tmp15549 := Call(__e, tmp15544, tmp15548)

						__e.TailApply(tmp15541, tmp15549)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15551 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5non_1terminal_1name_6 := __e.Get(1)
					_ = Parseshen_4_5non_1terminal_1name_6
					tmp15554 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5non_1terminal_1name_6)

					if True == tmp15554 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15553 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5non_1terminal_1name_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15553, True)
						return

					}

				}, 1)

				tmp15555 := Call(__e, PrimNS2Value(symshen_4_5non_1terminal_1name_6), V1159)

				tmp15556 := Call(__e, tmp15551, tmp15555)

				__e.TailApply(tmp15539, tmp15556)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15558 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5packagenames_6 := __e.Get(1)
			_ = Parseshen_4_5packagenames_6
			tmp15565 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5packagenames_6)

			if True == tmp15565 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15560 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5non_1terminal_1name_6 := __e.Get(1)
					_ = Parseshen_4_5non_1terminal_1name_6
					tmp15563 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5non_1terminal_1name_6)

					if True == tmp15563 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15562 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5non_1terminal_1name_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15562, True)
						return

					}

				}, 1)

				tmp15564 := Call(__e, PrimNS2Value(symshen_4_5non_1terminal_1name_6), Parseshen_4_5packagenames_6)

				__e.TailApply(tmp15560, tmp15564)
				return

			}

		}, 1)

		tmp15566 := Call(__e, PrimNS2Value(symshen_4_5packagenames_6), V1159)

		tmp15567 := Call(__e, tmp15558, tmp15566)

		__e.TailApply(tmp15537, tmp15567)
		return

	}, 1)

	tmp15568 := Call(__e, PrimNS2Value(symdef), symshen_4_5non_1terminal_2_6, tmp15536)

	_ = tmp15568

	tmp15569 := MakeNative(func(__e *ControlFlow) {
		V1160 := __e.Get(1)
		_ = V1160
		tmp15570 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15585 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15585 {
				tmp15572 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15574 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15574 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15575 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5packagename_6 := __e.Get(1)
					_ = Parseshen_4_5packagename_6
					tmp15582 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5packagename_6)

					if True == tmp15582 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15581 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5packagename_6, MakeString("."))

						if True == tmp15581 {
							tmp15578 := MakeNative(func(__e *ControlFlow) {
								News1047 := __e.Get(1)
								_ = News1047
								tmp15579 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1047)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp15579, symshen_4skip)
								return

							}, 1)

							tmp15580 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5packagename_6)

							__e.TailApply(tmp15578, tmp15580)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}

				}, 1)

				tmp15583 := Call(__e, PrimNS2Value(symshen_4_5packagename_6), V1160)

				tmp15584 := Call(__e, tmp15575, tmp15583)

				__e.TailApply(tmp15572, tmp15584)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15586 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5packagename_6 := __e.Get(1)
			_ = Parseshen_4_5packagename_6
			tmp15597 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5packagename_6)

			if True == tmp15597 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15596 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5packagename_6, MakeString("."))

				if True == tmp15596 {
					tmp15589 := MakeNative(func(__e *ControlFlow) {
						News1046 := __e.Get(1)
						_ = News1046
						tmp15590 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5packagenames_6 := __e.Get(1)
							_ = Parseshen_4_5packagenames_6
							tmp15593 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5packagenames_6)

							if True == tmp15593 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp15592 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5packagenames_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp15592, symshen_4skip)
								return

							}

						}, 1)

						tmp15594 := Call(__e, PrimNS2Value(symshen_4_5packagenames_6), News1046)

						__e.TailApply(tmp15590, tmp15594)
						return

					}, 1)

					tmp15595 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5packagename_6)

					__e.TailApply(tmp15589, tmp15595)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp15598 := Call(__e, PrimNS2Value(symshen_4_5packagename_6), V1160)

		tmp15599 := Call(__e, tmp15586, tmp15598)

		__e.TailApply(tmp15570, tmp15599)
		return

	}, 1)

	tmp15600 := Call(__e, PrimNS2Value(symdef), symshen_4_5packagenames_6, tmp15569)

	_ = tmp15600

	tmp15601 := MakeNative(func(__e *ControlFlow) {
		V1161 := __e.Get(1)
		_ = V1161
		tmp15602 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15613 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15613 {
				tmp15604 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15606 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15606 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15607 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp15610 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp15610 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15609 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15609, symshen_4skip)
						return

					}

				}, 1)

				tmp15611 := Call(__e, PrimNS2Value(sym_5e_6), V1161)

				tmp15612 := Call(__e, tmp15607, tmp15611)

				__e.TailApply(tmp15604, tmp15612)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15614 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5packagechar_6 := __e.Get(1)
			_ = Parseshen_4_5packagechar_6
			tmp15621 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5packagechar_6)

			if True == tmp15621 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15616 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5packagename_6 := __e.Get(1)
					_ = Parseshen_4_5packagename_6
					tmp15619 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5packagename_6)

					if True == tmp15619 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15618 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5packagename_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15618, symshen_4skip)
						return

					}

				}, 1)

				tmp15620 := Call(__e, PrimNS2Value(symshen_4_5packagename_6), Parseshen_4_5packagechar_6)

				__e.TailApply(tmp15616, tmp15620)
				return

			}

		}, 1)

		tmp15622 := Call(__e, PrimNS2Value(symshen_4_5packagechar_6), V1161)

		tmp15623 := Call(__e, tmp15614, tmp15622)

		__e.TailApply(tmp15602, tmp15623)
		return

	}, 1)

	tmp15624 := Call(__e, PrimNS2Value(symdef), symshen_4_5packagename_6, tmp15601)

	_ = tmp15624

	tmp15625 := MakeNative(func(__e *ControlFlow) {
		V1162 := __e.Get(1)
		_ = V1162
		tmp15626 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15628 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15628 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15640 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1162)

		var ifres15629 Obj

		if True == tmp15640 {
			tmp15631 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp15632 := MakeNative(func(__e *ControlFlow) {
					News1050 := __e.Get(1)
					_ = News1050
					tmp15635 := PrimEqual(X, MakeString("."))

					tmp15636 := PrimNot(tmp15635)

					if True == tmp15636 {
						tmp15634 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1050)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15634, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15637 := Call(__e, PrimNS2Value(symshen_4tls), V1162)

				__e.TailApply(tmp15632, tmp15637)
				return

			}, 1)

			tmp15638 := Call(__e, PrimNS2Value(symshen_4hds), V1162)

			tmp15639 := Call(__e, tmp15631, tmp15638)

			ifres15629 = tmp15639

		} else {
			tmp15630 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres15629 = tmp15630

		}

		__e.TailApply(tmp15626, ifres15629)
		return

	}, 1)

	tmp15641 := Call(__e, PrimNS2Value(symdef), symshen_4_5packagechar_6, tmp15625)

	_ = tmp15641

	tmp15642 := MakeNative(func(__e *ControlFlow) {
		V1163 := __e.Get(1)
		_ = V1163
		tmp15643 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15645 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15645 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15666 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V1163, MakeString("<"))

		var ifres15646 Obj

		if True == tmp15666 {
			tmp15648 := MakeNative(func(__e *ControlFlow) {
				News1052 := __e.Get(1)
				_ = News1052
				tmp15649 := MakeNative(func(__e *ControlFlow) {
					Parse_5_b_6 := __e.Get(1)
					_ = Parse_5_b_6
					tmp15662 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5_b_6)

					if True == tmp15662 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15653 := MakeNative(func(__e *ControlFlow) {
							Reverse := __e.Get(1)
							_ = Reverse
							tmp15658 := PrimIsPair(Reverse)

							if True == tmp15658 {
								tmp15656 := PrimHead(Reverse)

								tmp15657 := PrimEqual(tmp15656, MakeString(">"))

								if True == tmp15657 {
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

						tmp15659 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

						tmp15660 := Call(__e, PrimNS2Value(symreverse), tmp15659)

						tmp15661 := Call(__e, tmp15653, tmp15660)

						if True == tmp15661 {
							tmp15652 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5_b_6)

							__e.TailApply(PrimNS2Value(symshen_4comb), tmp15652, symshen_4skip)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}

				}, 1)

				tmp15663 := Call(__e, PrimNS2Value(sym_5_b_6), News1052)

				__e.TailApply(tmp15649, tmp15663)
				return

			}, 1)

			tmp15664 := Call(__e, PrimNS2Value(symshen_4tls), V1163)

			tmp15665 := Call(__e, tmp15648, tmp15664)

			ifres15646 = tmp15665

		} else {
			tmp15647 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres15646 = tmp15647

		}

		__e.TailApply(tmp15643, ifres15646)
		return

	}, 1)

	tmp15667 := Call(__e, PrimNS2Value(symdef), symshen_4_5non_1terminal_1name_6, tmp15642)

	_ = tmp15667

	tmp15668 := MakeNative(func(__e *ControlFlow) {
		V1164 := __e.Get(1)
		_ = V1164
		tmp15669 := PrimIntern(MakeString(";"))

		__e.Return(PrimEqual(V1164, tmp15669))
		return

	}, 1)

	tmp15670 := Call(__e, PrimNS2Value(symdef), symshen_4semicolon_2, tmp15668)

	_ = tmp15670

	tmp15671 := MakeNative(func(__e *ControlFlow) {
		V1165 := __e.Get(1)
		_ = V1165
		tmp15672 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15674 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15674 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15685 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1165)

		var ifres15675 Obj

		if True == tmp15685 {
			tmp15677 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp15678 := MakeNative(func(__e *ControlFlow) {
					News1054 := __e.Get(1)
					_ = News1054
					tmp15681 := Call(__e, PrimNS2Value(symshen_4colon_1equal_2), X)

					if True == tmp15681 {
						tmp15680 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1054)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15680, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15682 := Call(__e, PrimNS2Value(symshen_4tls), V1165)

				__e.TailApply(tmp15678, tmp15682)
				return

			}, 1)

			tmp15683 := Call(__e, PrimNS2Value(symshen_4hds), V1165)

			tmp15684 := Call(__e, tmp15677, tmp15683)

			ifres15675 = tmp15684

		} else {
			tmp15676 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres15675 = tmp15676

		}

		__e.TailApply(tmp15672, ifres15675)
		return

	}, 1)

	tmp15686 := Call(__e, PrimNS2Value(symdef), symshen_4_5colon_1equal_6, tmp15671)

	_ = tmp15686

	tmp15687 := MakeNative(func(__e *ControlFlow) {
		V1166 := __e.Get(1)
		_ = V1166
		tmp15688 := PrimIntern(MakeString(":="))

		__e.Return(PrimEqual(tmp15688, V1166))
		return

	}, 1)

	tmp15689 := Call(__e, PrimNS2Value(symdef), symshen_4colon_1equal_2, tmp15687)

	_ = tmp15689

	tmp15690 := MakeNative(func(__e *ControlFlow) {
		V1167 := __e.Get(1)
		_ = V1167
		tmp15691 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15704 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15704 {
				tmp15693 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15695 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15695 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15696 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5syntax_1item_6 := __e.Get(1)
					_ = Parseshen_4_5syntax_1item_6
					tmp15701 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5syntax_1item_6)

					if True == tmp15701 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15698 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5syntax_1item_6)

						tmp15699 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5syntax_1item_6)

						tmp15700 := PrimCons(tmp15699, Nil)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15698, tmp15700)
						return

					}

				}, 1)

				tmp15702 := Call(__e, PrimNS2Value(symshen_4_5syntax_1item_6), V1167)

				tmp15703 := Call(__e, tmp15696, tmp15702)

				__e.TailApply(tmp15693, tmp15703)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15705 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5syntax_1item_6 := __e.Get(1)
			_ = Parseshen_4_5syntax_1item_6
			tmp15715 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5syntax_1item_6)

			if True == tmp15715 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15707 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5syntax_6 := __e.Get(1)
					_ = Parseshen_4_5syntax_6
					tmp15713 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5syntax_6)

					if True == tmp15713 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15709 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5syntax_6)

						tmp15710 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5syntax_1item_6)

						tmp15711 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5syntax_6)

						tmp15712 := PrimCons(tmp15710, tmp15711)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15709, tmp15712)
						return

					}

				}, 1)

				tmp15714 := Call(__e, PrimNS2Value(symshen_4_5syntax_6), Parseshen_4_5syntax_1item_6)

				__e.TailApply(tmp15707, tmp15714)
				return

			}

		}, 1)

		tmp15716 := Call(__e, PrimNS2Value(symshen_4_5syntax_1item_6), V1167)

		tmp15717 := Call(__e, tmp15705, tmp15716)

		__e.TailApply(tmp15691, tmp15717)
		return

	}, 1)

	tmp15718 := Call(__e, PrimNS2Value(symdef), symshen_4_5syntax_6, tmp15690)

	_ = tmp15718

	tmp15719 := MakeNative(func(__e *ControlFlow) {
		V1168 := __e.Get(1)
		_ = V1168
		tmp15720 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15722 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15722 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15733 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1168)

		var ifres15723 Obj

		if True == tmp15733 {
			tmp15725 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp15726 := MakeNative(func(__e *ControlFlow) {
					News1057 := __e.Get(1)
					_ = News1057
					tmp15729 := Call(__e, PrimNS2Value(symshen_4syntax_1item_2), X)

					if True == tmp15729 {
						tmp15728 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1057)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp15728, X)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp15730 := Call(__e, PrimNS2Value(symshen_4tls), V1168)

				__e.TailApply(tmp15726, tmp15730)
				return

			}, 1)

			tmp15731 := Call(__e, PrimNS2Value(symshen_4hds), V1168)

			tmp15732 := Call(__e, tmp15725, tmp15731)

			ifres15723 = tmp15732

		} else {
			tmp15724 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres15723 = tmp15724

		}

		__e.TailApply(tmp15720, ifres15723)
		return

	}, 1)

	tmp15734 := Call(__e, PrimNS2Value(symdef), symshen_4_5syntax_1item_6, tmp15719)

	_ = tmp15734

	tmp15735 := MakeNative(func(__e *ControlFlow) {
		V1171 := __e.Get(1)
		_ = V1171
		tmp15771 := Call(__e, PrimNS2Value(symshen_4colon_1equal_2), V1171)

		if True == tmp15771 {
			__e.Return(False)
			return
		} else {
			tmp15770 := Call(__e, PrimNS2Value(symshen_4semicolon_2), V1171)

			if True == tmp15770 {
				__e.Return(False)
				return
			} else {
				tmp15769 := Call(__e, PrimNS2Value(symatom_2), V1171)

				if True == tmp15769 {
					__e.Return(True)
					return
				} else {
					tmp15768 := PrimIsPair(V1171)

					var ifres15749 Obj

					if True == tmp15768 {
						tmp15766 := PrimHead(V1171)

						tmp15767 := PrimEqual(symcons, tmp15766)

						var ifres15751 Obj

						if True == tmp15767 {
							tmp15764 := PrimTail(V1171)

							tmp15765 := PrimIsPair(tmp15764)

							var ifres15753 Obj

							if True == tmp15765 {
								tmp15761 := PrimTail(V1171)

								tmp15762 := PrimTail(tmp15761)

								tmp15763 := PrimIsPair(tmp15762)

								var ifres15755 Obj

								if True == tmp15763 {
									tmp15757 := PrimTail(V1171)

									tmp15758 := PrimTail(tmp15757)

									tmp15759 := PrimTail(tmp15758)

									tmp15760 := PrimEqual(Nil, tmp15759)

									var ifres15756 Obj

									if True == tmp15760 {
										ifres15756 = True

									} else {
										ifres15756 = False

									}

									ifres15755 = ifres15756

								} else {
									ifres15755 = False

								}

								var ifres15754 Obj

								if True == ifres15755 {
									ifres15754 = True

								} else {
									ifres15754 = False

								}

								ifres15753 = ifres15754

							} else {
								ifres15753 = False

							}

							var ifres15752 Obj

							if True == ifres15753 {
								ifres15752 = True

							} else {
								ifres15752 = False

							}

							ifres15751 = ifres15752

						} else {
							ifres15751 = False

						}

						var ifres15750 Obj

						if True == ifres15751 {
							ifres15750 = True

						} else {
							ifres15750 = False

						}

						ifres15749 = ifres15750

					} else {
						ifres15749 = False

					}

					if True == ifres15749 {
						tmp15746 := PrimTail(V1171)

						tmp15747 := PrimHead(tmp15746)

						tmp15748 := Call(__e, PrimNS2Value(symshen_4syntax_1item_2), tmp15747)

						if True == tmp15748 {
							tmp15742 := PrimTail(V1171)

							tmp15743 := PrimTail(tmp15742)

							tmp15744 := PrimHead(tmp15743)

							tmp15745 := Call(__e, PrimNS2Value(symshen_4syntax_1item_2), tmp15744)

							if True == tmp15745 {
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

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 1)

	tmp15772 := Call(__e, PrimNS2Value(symdef), symshen_4syntax_1item_2, tmp15735)

	_ = tmp15772

	tmp15773 := MakeNative(func(__e *ControlFlow) {
		V1172 := __e.Get(1)
		_ = V1172
		tmp15774 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp15794 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp15794 {
				tmp15776 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp15778 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp15778 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp15779 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5colon_1equal_6 := __e.Get(1)
					_ = Parseshen_4_5colon_1equal_6
					tmp15791 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5colon_1equal_6)

					if True == tmp15791 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp15790 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), Parseshen_4_5colon_1equal_6)

						if True == tmp15790 {
							tmp15782 := MakeNative(func(__e *ControlFlow) {
								Semantics := __e.Get(1)
								_ = Semantics
								tmp15783 := MakeNative(func(__e *ControlFlow) {
									News1062 := __e.Get(1)
									_ = News1062
									tmp15786 := Call(__e, PrimNS2Value(symshen_4semicolon_2), Semantics)

									tmp15787 := PrimNot(tmp15786)

									if True == tmp15787 {
										tmp15785 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1062)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp15785, Semantics)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp15788 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5colon_1equal_6)

								__e.TailApply(tmp15783, tmp15788)
								return

							}, 1)

							tmp15789 := Call(__e, PrimNS2Value(symshen_4hds), Parseshen_4_5colon_1equal_6)

							__e.TailApply(tmp15782, tmp15789)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}

				}, 1)

				tmp15792 := Call(__e, PrimNS2Value(symshen_4_5colon_1equal_6), V1172)

				tmp15793 := Call(__e, tmp15779, tmp15792)

				__e.TailApply(tmp15776, tmp15793)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp15795 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5colon_1equal_6 := __e.Get(1)
			_ = Parseshen_4_5colon_1equal_6
			tmp15820 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5colon_1equal_6)

			if True == tmp15820 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp15819 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), Parseshen_4_5colon_1equal_6)

				if True == tmp15819 {
					tmp15798 := MakeNative(func(__e *ControlFlow) {
						Semantics := __e.Get(1)
						_ = Semantics
						tmp15799 := MakeNative(func(__e *ControlFlow) {
							News1059 := __e.Get(1)
							_ = News1059
							tmp15816 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News1059, symwhere)

							if True == tmp15816 {
								tmp15801 := MakeNative(func(__e *ControlFlow) {
									News1060 := __e.Get(1)
									_ = News1060
									tmp15814 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1060)

									if True == tmp15814 {
										tmp15803 := MakeNative(func(__e *ControlFlow) {
											Guard := __e.Get(1)
											_ = Guard
											tmp15804 := MakeNative(func(__e *ControlFlow) {
												News1061 := __e.Get(1)
												_ = News1061
												tmp15810 := Call(__e, PrimNS2Value(symshen_4semicolon_2), Semantics)

												tmp15811 := PrimNot(tmp15810)

												if True == tmp15811 {
													tmp15806 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1061)

													tmp15807 := PrimCons(Semantics, Nil)

													tmp15808 := PrimCons(Guard, tmp15807)

													tmp15809 := PrimCons(symwhere, tmp15808)

													__e.TailApply(PrimNS2Value(symshen_4comb), tmp15806, tmp15809)
													return

												} else {
													__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp15812 := Call(__e, PrimNS2Value(symshen_4tls), News1060)

											__e.TailApply(tmp15804, tmp15812)
											return

										}, 1)

										tmp15813 := Call(__e, PrimNS2Value(symshen_4hds), News1060)

										__e.TailApply(tmp15803, tmp15813)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp15815 := Call(__e, PrimNS2Value(symshen_4tls), News1059)

								__e.TailApply(tmp15801, tmp15815)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp15817 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5colon_1equal_6)

						__e.TailApply(tmp15799, tmp15817)
						return

					}, 1)

					tmp15818 := Call(__e, PrimNS2Value(symshen_4hds), Parseshen_4_5colon_1equal_6)

					__e.TailApply(tmp15798, tmp15818)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp15821 := Call(__e, PrimNS2Value(symshen_4_5colon_1equal_6), V1172)

		tmp15822 := Call(__e, tmp15795, tmp15821)

		__e.TailApply(tmp15774, tmp15822)
		return

	}, 1)

	tmp15823 := Call(__e, PrimNS2Value(symdef), symshen_4_5semantics_6, tmp15773)

	_ = tmp15823

	tmp15824 := MakeNative(func(__e *ControlFlow) {
		V1181 := __e.Get(1)
		_ = V1181
		V1182 := __e.Get(2)
		_ = V1182
		V1183 := __e.Get(3)
		_ = V1183
		tmp15832 := PrimEqual(Nil, V1183)

		if True == tmp15832 {
			__e.Return(PrimCons(symshen_4parse_1failure, Nil))
			return
		} else {
			tmp15831 := PrimIsPair(V1183)

			if True == tmp15831 {
				tmp15827 := PrimHead(V1183)

				tmp15828 := Call(__e, PrimNS2Value(symshen_4c_1rule_1_6shen), V1181, tmp15827, V1182)

				tmp15829 := PrimTail(V1183)

				tmp15830 := Call(__e, PrimNS2Value(symshen_4c_1rules_1_6shen), V1181, V1182, tmp15829)

				__e.TailApply(PrimNS2Value(symshen_4combine_1c_1code), tmp15828, tmp15830)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
				return
			}

		}

	}, 3)

	tmp15833 := Call(__e, PrimNS2Value(symdef), symshen_4c_1rules_1_6shen, tmp15824)

	_ = tmp15833

	tmp15834 := MakeNative(func(__e *ControlFlow) {
		__e.TailApply(PrimNS2Value(symfail))
		return
	}, 0)

	tmp15835 := Call(__e, PrimNS2Value(symdef), symshen_4parse_1failure, tmp15834)

	_ = tmp15835

	tmp15836 := MakeNative(func(__e *ControlFlow) {
		V1184 := __e.Get(1)
		_ = V1184
		V1185 := __e.Get(2)
		_ = V1185
		tmp15837 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp15838 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp15839 := PrimCons(tmp15838, Nil)

		tmp15840 := PrimCons(symshen_4parse_1failure_2, tmp15839)

		tmp15841 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp15842 := PrimCons(tmp15841, Nil)

		tmp15843 := PrimCons(V1185, tmp15842)

		tmp15844 := PrimCons(tmp15840, tmp15843)

		tmp15845 := PrimCons(symif, tmp15844)

		tmp15846 := PrimCons(tmp15845, Nil)

		tmp15847 := PrimCons(V1184, tmp15846)

		tmp15848 := PrimCons(tmp15837, tmp15847)

		__e.Return(PrimCons(symlet, tmp15848))
		return

	}, 2)

	tmp15849 := Call(__e, PrimNS2Value(symdef), symshen_4combine_1c_1code, tmp15836)

	_ = tmp15849

	tmp15850 := MakeNative(func(__e *ControlFlow) {
		V1192 := __e.Get(1)
		_ = V1192
		V1193 := __e.Get(2)
		_ = V1193
		V1194 := __e.Get(3)
		_ = V1194
		tmp15864 := PrimIsPair(V1193)

		var ifres15855 Obj

		if True == tmp15864 {
			tmp15862 := PrimTail(V1193)

			tmp15863 := PrimIsPair(tmp15862)

			var ifres15857 Obj

			if True == tmp15863 {
				tmp15859 := PrimTail(V1193)

				tmp15860 := PrimTail(tmp15859)

				tmp15861 := PrimEqual(Nil, tmp15860)

				var ifres15858 Obj

				if True == tmp15861 {
					ifres15858 = True

				} else {
					ifres15858 = False

				}

				ifres15857 = ifres15858

			} else {
				ifres15857 = False

			}

			var ifres15856 Obj

			if True == ifres15857 {
				ifres15856 = True

			} else {
				ifres15856 = False

			}

			ifres15855 = ifres15856

		} else {
			ifres15855 = False

		}

		if True == ifres15855 {
			tmp15852 := PrimHead(V1193)

			tmp15853 := PrimTail(V1193)

			tmp15854 := PrimHead(tmp15853)

			__e.TailApply(PrimNS2Value(symshen_4yacc_1syntax), V1192, V1194, tmp15852, tmp15854)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
			return
		}

	}, 3)

	tmp15865 := Call(__e, PrimNS2Value(symdef), symshen_4c_1rule_1_6shen, tmp15850)

	_ = tmp15865

	tmp15866 := MakeNative(func(__e *ControlFlow) {
		V1203 := __e.Get(1)
		_ = V1203
		V1204 := __e.Get(2)
		_ = V1204
		V1205 := __e.Get(3)
		_ = V1205
		V1206 := __e.Get(4)
		_ = V1206
		tmp15930 := PrimEqual(Nil, V1205)

		var ifres15908 Obj

		if True == tmp15930 {
			tmp15929 := PrimIsPair(V1206)

			var ifres15910 Obj

			if True == tmp15929 {
				tmp15927 := PrimHead(V1206)

				tmp15928 := PrimEqual(symwhere, tmp15927)

				var ifres15912 Obj

				if True == tmp15928 {
					tmp15925 := PrimTail(V1206)

					tmp15926 := PrimIsPair(tmp15925)

					var ifres15914 Obj

					if True == tmp15926 {
						tmp15922 := PrimTail(V1206)

						tmp15923 := PrimTail(tmp15922)

						tmp15924 := PrimIsPair(tmp15923)

						var ifres15916 Obj

						if True == tmp15924 {
							tmp15918 := PrimTail(V1206)

							tmp15919 := PrimTail(tmp15918)

							tmp15920 := PrimTail(tmp15919)

							tmp15921 := PrimEqual(Nil, tmp15920)

							var ifres15917 Obj

							if True == tmp15921 {
								ifres15917 = True

							} else {
								ifres15917 = False

							}

							ifres15916 = ifres15917

						} else {
							ifres15916 = False

						}

						var ifres15915 Obj

						if True == ifres15916 {
							ifres15915 = True

						} else {
							ifres15915 = False

						}

						ifres15914 = ifres15915

					} else {
						ifres15914 = False

					}

					var ifres15913 Obj

					if True == ifres15914 {
						ifres15913 = True

					} else {
						ifres15913 = False

					}

					ifres15912 = ifres15913

				} else {
					ifres15912 = False

				}

				var ifres15911 Obj

				if True == ifres15912 {
					ifres15911 = True

				} else {
					ifres15911 = False

				}

				ifres15910 = ifres15911

			} else {
				ifres15910 = False

			}

			var ifres15909 Obj

			if True == ifres15910 {
				ifres15909 = True

			} else {
				ifres15909 = False

			}

			ifres15908 = ifres15909

		} else {
			ifres15908 = False

		}

		if True == ifres15908 {
			tmp15897 := PrimTail(V1206)

			tmp15898 := PrimHead(tmp15897)

			tmp15899 := Call(__e, PrimNS2Value(symshen_4process_1yacc_1semantics), tmp15898)

			tmp15900 := PrimTail(V1206)

			tmp15901 := PrimTail(tmp15900)

			tmp15902 := PrimHead(tmp15901)

			tmp15903 := Call(__e, PrimNS2Value(symshen_4yacc_1syntax), V1203, V1204, Nil, tmp15902)

			tmp15904 := PrimCons(symshen_4parse_1failure, Nil)

			tmp15905 := PrimCons(tmp15904, Nil)

			tmp15906 := PrimCons(tmp15903, tmp15905)

			tmp15907 := PrimCons(tmp15899, tmp15906)

			__e.Return(PrimCons(symif, tmp15907))
			return

		} else {
			tmp15896 := PrimEqual(Nil, V1205)

			if True == tmp15896 {
				__e.TailApply(PrimNS2Value(symshen_4yacc_1semantics), V1203, V1204, V1206)
				return
			} else {
				tmp15895 := PrimIsPair(V1205)

				if True == tmp15895 {
					tmp15893 := PrimHead(V1205)

					tmp15894 := Call(__e, PrimNS2Value(symshen_4non_1terminal_2), tmp15893)

					if True == tmp15894 {
						tmp15891 := PrimHead(V1205)

						tmp15892 := PrimTail(V1205)

						__e.TailApply(PrimNS2Value(symshen_4non_1terminalcode), V1203, V1204, tmp15891, tmp15892, V1206)
						return

					} else {
						tmp15889 := PrimHead(V1205)

						tmp15890 := PrimIsVariable(tmp15889)

						if True == tmp15890 {
							tmp15887 := PrimHead(V1205)

							tmp15888 := PrimTail(V1205)

							__e.TailApply(PrimNS2Value(symshen_4variablecode), V1203, V1204, tmp15887, tmp15888, V1206)
							return

						} else {
							tmp15885 := PrimHead(V1205)

							tmp15886 := PrimEqual(sym__, tmp15885)

							if True == tmp15886 {
								tmp15883 := PrimHead(V1205)

								tmp15884 := PrimTail(V1205)

								__e.TailApply(PrimNS2Value(symshen_4wildcardcode), V1203, V1204, tmp15883, tmp15884, V1206)
								return

							} else {
								tmp15881 := PrimHead(V1205)

								tmp15882 := Call(__e, PrimNS2Value(symatom_2), tmp15881)

								if True == tmp15882 {
									tmp15879 := PrimHead(V1205)

									tmp15880 := PrimTail(V1205)

									__e.TailApply(PrimNS2Value(symshen_4terminalcode), V1203, V1204, tmp15879, tmp15880, V1206)
									return

								} else {
									tmp15877 := PrimHead(V1205)

									tmp15878 := PrimIsPair(tmp15877)

									if True == tmp15878 {
										tmp15875 := PrimHead(V1205)

										tmp15876 := PrimTail(V1205)

										__e.TailApply(PrimNS2Value(symshen_4conscode), V1203, V1204, tmp15875, tmp15876, V1206)
										return

									} else {
										__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
										return
									}

								}

							}

						}

					}

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
					return
				}

			}

		}

	}, 4)

	tmp15931 := Call(__e, PrimNS2Value(symdef), symshen_4yacc_1syntax, tmp15866)

	_ = tmp15931

	tmp15932 := MakeNative(func(__e *ControlFlow) {
		V1207 := __e.Get(1)
		_ = V1207
		V1208 := __e.Get(2)
		_ = V1208
		V1209 := __e.Get(3)
		_ = V1209
		V1210 := __e.Get(4)
		_ = V1210
		V1211 := __e.Get(5)
		_ = V1211
		tmp15933 := MakeNative(func(__e *ControlFlow) {
			ApplyNonTerminal := __e.Get(1)
			_ = ApplyNonTerminal
			tmp15934 := PrimCons(V1208, Nil)

			tmp15935 := PrimCons(V1209, tmp15934)

			tmp15936 := PrimCons(ApplyNonTerminal, Nil)

			tmp15937 := PrimCons(symshen_4parse_1failure_2, tmp15936)

			tmp15938 := PrimCons(symshen_4parse_1failure, Nil)

			tmp15939 := Call(__e, PrimNS2Value(symshen_4yacc_1syntax), V1207, ApplyNonTerminal, V1210, V1211)

			tmp15940 := PrimCons(tmp15939, Nil)

			tmp15941 := PrimCons(tmp15938, tmp15940)

			tmp15942 := PrimCons(tmp15937, tmp15941)

			tmp15943 := PrimCons(symif, tmp15942)

			tmp15944 := PrimCons(tmp15943, Nil)

			tmp15945 := PrimCons(tmp15935, tmp15944)

			tmp15946 := PrimCons(ApplyNonTerminal, tmp15945)

			__e.Return(PrimCons(symlet, tmp15946))
			return

		}, 1)

		tmp15947 := Call(__e, PrimNS2Value(symprotect), symParse)

		tmp15948 := Call(__e, PrimNS2Value(symconcat), tmp15947, V1209)

		__e.TailApply(tmp15933, tmp15948)
		return

	}, 5)

	tmp15949 := Call(__e, PrimNS2Value(symdef), symshen_4non_1terminalcode, tmp15932)

	_ = tmp15949

	tmp15950 := MakeNative(func(__e *ControlFlow) {
		V1212 := __e.Get(1)
		_ = V1212
		V1213 := __e.Get(2)
		_ = V1213
		V1214 := __e.Get(3)
		_ = V1214
		V1215 := __e.Get(4)
		_ = V1215
		V1216 := __e.Get(5)
		_ = V1216
		tmp15951 := MakeNative(func(__e *ControlFlow) {
			NewStream := __e.Get(1)
			_ = NewStream
			tmp15952 := PrimCons(V1213, Nil)

			tmp15953 := PrimCons(symshen_4non_1empty_1stream_2, tmp15952)

			tmp15954 := PrimCons(V1213, Nil)

			tmp15955 := PrimCons(symshen_4hds, tmp15954)

			tmp15956 := PrimCons(V1213, Nil)

			tmp15957 := PrimCons(symshen_4tls, tmp15956)

			tmp15958 := Call(__e, PrimNS2Value(symshen_4yacc_1syntax), V1212, NewStream, V1215, V1216)

			tmp15959 := PrimCons(tmp15958, Nil)

			tmp15960 := PrimCons(tmp15957, tmp15959)

			tmp15961 := PrimCons(NewStream, tmp15960)

			tmp15962 := PrimCons(tmp15955, tmp15961)

			tmp15963 := PrimCons(V1214, tmp15962)

			tmp15964 := PrimCons(symlet, tmp15963)

			tmp15965 := PrimCons(symshen_4parse_1failure, Nil)

			tmp15966 := PrimCons(tmp15965, Nil)

			tmp15967 := PrimCons(tmp15964, tmp15966)

			tmp15968 := PrimCons(tmp15953, tmp15967)

			__e.Return(PrimCons(symif, tmp15968))
			return

		}, 1)

		tmp15969 := Call(__e, PrimNS2Value(symprotect), symNews)

		tmp15970 := Call(__e, PrimNS2Value(symgensym), tmp15969)

		__e.TailApply(tmp15951, tmp15970)
		return

	}, 5)

	tmp15971 := Call(__e, PrimNS2Value(symdef), symshen_4variablecode, tmp15950)

	_ = tmp15971

	tmp15972 := MakeNative(func(__e *ControlFlow) {
		V1217 := __e.Get(1)
		_ = V1217
		V1218 := __e.Get(2)
		_ = V1218
		V1219 := __e.Get(3)
		_ = V1219
		V1220 := __e.Get(4)
		_ = V1220
		V1221 := __e.Get(5)
		_ = V1221
		tmp15973 := MakeNative(func(__e *ControlFlow) {
			NewStream := __e.Get(1)
			_ = NewStream
			tmp15974 := PrimCons(V1218, Nil)

			tmp15975 := PrimCons(symshen_4non_1empty_1stream_2, tmp15974)

			tmp15976 := PrimCons(V1218, Nil)

			tmp15977 := PrimCons(symshen_4tls, tmp15976)

			tmp15978 := Call(__e, PrimNS2Value(symshen_4yacc_1syntax), V1217, NewStream, V1220, V1221)

			tmp15979 := PrimCons(tmp15978, Nil)

			tmp15980 := PrimCons(tmp15977, tmp15979)

			tmp15981 := PrimCons(NewStream, tmp15980)

			tmp15982 := PrimCons(symlet, tmp15981)

			tmp15983 := PrimCons(symshen_4parse_1failure, Nil)

			tmp15984 := PrimCons(tmp15983, Nil)

			tmp15985 := PrimCons(tmp15982, tmp15984)

			tmp15986 := PrimCons(tmp15975, tmp15985)

			__e.Return(PrimCons(symif, tmp15986))
			return

		}, 1)

		tmp15987 := Call(__e, PrimNS2Value(symprotect), symNews)

		tmp15988 := Call(__e, PrimNS2Value(symgensym), tmp15987)

		__e.TailApply(tmp15973, tmp15988)
		return

	}, 5)

	tmp15989 := Call(__e, PrimNS2Value(symdef), symshen_4wildcardcode, tmp15972)

	_ = tmp15989

	tmp15990 := MakeNative(func(__e *ControlFlow) {
		V1222 := __e.Get(1)
		_ = V1222
		V1223 := __e.Get(2)
		_ = V1223
		V1224 := __e.Get(3)
		_ = V1224
		V1225 := __e.Get(4)
		_ = V1225
		V1226 := __e.Get(5)
		_ = V1226
		tmp15991 := MakeNative(func(__e *ControlFlow) {
			NewStream := __e.Get(1)
			_ = NewStream
			tmp15992 := PrimCons(V1224, Nil)

			tmp15993 := PrimCons(V1223, tmp15992)

			tmp15994 := PrimCons(symshen_4_ahd_2, tmp15993)

			tmp15995 := PrimCons(V1223, Nil)

			tmp15996 := PrimCons(symshen_4tls, tmp15995)

			tmp15997 := Call(__e, PrimNS2Value(symshen_4yacc_1syntax), V1222, NewStream, V1225, V1226)

			tmp15998 := PrimCons(tmp15997, Nil)

			tmp15999 := PrimCons(tmp15996, tmp15998)

			tmp16000 := PrimCons(NewStream, tmp15999)

			tmp16001 := PrimCons(symlet, tmp16000)

			tmp16002 := PrimCons(symshen_4parse_1failure, Nil)

			tmp16003 := PrimCons(tmp16002, Nil)

			tmp16004 := PrimCons(tmp16001, tmp16003)

			tmp16005 := PrimCons(tmp15994, tmp16004)

			__e.Return(PrimCons(symif, tmp16005))
			return

		}, 1)

		tmp16006 := Call(__e, PrimNS2Value(symprotect), symNews)

		tmp16007 := Call(__e, PrimNS2Value(symgensym), tmp16006)

		__e.TailApply(tmp15991, tmp16007)
		return

	}, 5)

	tmp16008 := Call(__e, PrimNS2Value(symdef), symshen_4terminalcode, tmp15990)

	_ = tmp16008

	tmp16009 := MakeNative(func(__e *ControlFlow) {
		V1227 := __e.Get(1)
		_ = V1227
		V1228 := __e.Get(2)
		_ = V1228
		V1229 := __e.Get(3)
		_ = V1229
		V1230 := __e.Get(4)
		_ = V1230
		V1231 := __e.Get(5)
		_ = V1231
		tmp16010 := PrimCons(V1228, Nil)

		tmp16011 := PrimCons(symshen_4ccons_2, tmp16010)

		tmp16012 := Call(__e, PrimNS2Value(symprotect), symSynCons)

		tmp16013 := PrimCons(V1228, Nil)

		tmp16014 := PrimCons(symshen_4hds, tmp16013)

		tmp16015 := PrimCons(V1228, Nil)

		tmp16016 := PrimCons(symshen_4_5_1out, tmp16015)

		tmp16017 := PrimCons(tmp16016, Nil)

		tmp16018 := PrimCons(tmp16014, tmp16017)

		tmp16019 := PrimCons(symshen_4comb, tmp16018)

		tmp16020 := Call(__e, PrimNS2Value(symprotect), symSynCons)

		tmp16021 := Call(__e, PrimNS2Value(symshen_4decons), V1229)

		tmp16022 := PrimCons(symshen_4_5end_6, Nil)

		tmp16023 := Call(__e, PrimNS2Value(symappend), tmp16021, tmp16022)

		tmp16024 := PrimCons(V1228, Nil)

		tmp16025 := PrimCons(symshen_4tlstream, tmp16024)

		tmp16026 := PrimCons(V1231, Nil)

		tmp16027 := PrimCons(V1230, tmp16026)

		tmp16028 := PrimCons(tmp16025, tmp16027)

		tmp16029 := PrimCons(symshen_4pushsemantics, tmp16028)

		tmp16030 := Call(__e, PrimNS2Value(symshen_4yacc_1syntax), V1227, tmp16020, tmp16023, tmp16029)

		tmp16031 := PrimCons(tmp16030, Nil)

		tmp16032 := PrimCons(tmp16019, tmp16031)

		tmp16033 := PrimCons(tmp16012, tmp16032)

		tmp16034 := PrimCons(symlet, tmp16033)

		tmp16035 := PrimCons(symshen_4parse_1failure, Nil)

		tmp16036 := PrimCons(tmp16035, Nil)

		tmp16037 := PrimCons(tmp16034, tmp16036)

		tmp16038 := PrimCons(tmp16011, tmp16037)

		__e.Return(PrimCons(symif, tmp16038))
		return

	}, 5)

	tmp16039 := Call(__e, PrimNS2Value(symdef), symshen_4conscode, tmp16009)

	_ = tmp16039

	tmp16040 := MakeNative(func(__e *ControlFlow) {
		V1232 := __e.Get(1)
		_ = V1232
		tmp16067 := PrimIsPair(V1232)

		var ifres16048 Obj

		if True == tmp16067 {
			tmp16065 := PrimHead(V1232)

			tmp16066 := PrimEqual(symcons, tmp16065)

			var ifres16050 Obj

			if True == tmp16066 {
				tmp16063 := PrimTail(V1232)

				tmp16064 := PrimIsPair(tmp16063)

				var ifres16052 Obj

				if True == tmp16064 {
					tmp16060 := PrimTail(V1232)

					tmp16061 := PrimTail(tmp16060)

					tmp16062 := PrimIsPair(tmp16061)

					var ifres16054 Obj

					if True == tmp16062 {
						tmp16056 := PrimTail(V1232)

						tmp16057 := PrimTail(tmp16056)

						tmp16058 := PrimTail(tmp16057)

						tmp16059 := PrimEqual(Nil, tmp16058)

						var ifres16055 Obj

						if True == tmp16059 {
							ifres16055 = True

						} else {
							ifres16055 = False

						}

						ifres16054 = ifres16055

					} else {
						ifres16054 = False

					}

					var ifres16053 Obj

					if True == ifres16054 {
						ifres16053 = True

					} else {
						ifres16053 = False

					}

					ifres16052 = ifres16053

				} else {
					ifres16052 = False

				}

				var ifres16051 Obj

				if True == ifres16052 {
					ifres16051 = True

				} else {
					ifres16051 = False

				}

				ifres16050 = ifres16051

			} else {
				ifres16050 = False

			}

			var ifres16049 Obj

			if True == ifres16050 {
				ifres16049 = True

			} else {
				ifres16049 = False

			}

			ifres16048 = ifres16049

		} else {
			ifres16048 = False

		}

		if True == ifres16048 {
			tmp16042 := PrimTail(V1232)

			tmp16043 := PrimHead(tmp16042)

			tmp16044 := PrimTail(V1232)

			tmp16045 := PrimTail(tmp16044)

			tmp16046 := PrimHead(tmp16045)

			tmp16047 := Call(__e, PrimNS2Value(symshen_4decons), tmp16046)

			__e.Return(PrimCons(tmp16043, tmp16047))
			return

		} else {
			__e.Return(V1232)
			return
		}

	}, 1)

	tmp16068 := Call(__e, PrimNS2Value(symdef), symshen_4decons, tmp16040)

	_ = tmp16068

	tmp16069 := MakeNative(func(__e *ControlFlow) {
		V1239 := __e.Get(1)
		_ = V1239
		tmp16086 := PrimIsPair(V1239)

		var ifres16073 Obj

		if True == tmp16086 {
			tmp16084 := PrimHead(V1239)

			tmp16085 := PrimIsPair(tmp16084)

			var ifres16075 Obj

			if True == tmp16085 {
				tmp16082 := PrimTail(V1239)

				tmp16083 := PrimIsPair(tmp16082)

				var ifres16077 Obj

				if True == tmp16083 {
					tmp16079 := PrimTail(V1239)

					tmp16080 := PrimTail(tmp16079)

					tmp16081 := PrimEqual(Nil, tmp16080)

					var ifres16078 Obj

					if True == tmp16081 {
						ifres16078 = True

					} else {
						ifres16078 = False

					}

					ifres16077 = ifres16078

				} else {
					ifres16077 = False

				}

				var ifres16076 Obj

				if True == ifres16077 {
					ifres16076 = True

				} else {
					ifres16076 = False

				}

				ifres16075 = ifres16076

			} else {
				ifres16075 = False

			}

			var ifres16074 Obj

			if True == ifres16075 {
				ifres16074 = True

			} else {
				ifres16074 = False

			}

			ifres16073 = ifres16074

		} else {
			ifres16073 = False

		}

		if True == ifres16073 {
			tmp16071 := PrimHead(V1239)

			tmp16072 := PrimHead(tmp16071)

			__e.Return(PrimIsPair(tmp16072))
			return

		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp16087 := Call(__e, PrimNS2Value(symdef), symshen_4ccons_2, tmp16069)

	_ = tmp16087

	tmp16088 := MakeNative(func(__e *ControlFlow) {
		V1248 := __e.Get(1)
		_ = V1248
		tmp16094 := PrimIsPair(V1248)

		var ifres16090 Obj

		if True == tmp16094 {
			tmp16092 := PrimHead(V1248)

			tmp16093 := PrimIsPair(tmp16092)

			var ifres16091 Obj

			if True == tmp16093 {
				ifres16091 = True

			} else {
				ifres16091 = False

			}

			ifres16090 = ifres16091

		} else {
			ifres16090 = False

		}

		if True == ifres16090 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp16095 := Call(__e, PrimNS2Value(symdef), symshen_4non_1empty_1stream_2, tmp16088)

	_ = tmp16095

	tmp16096 := MakeNative(func(__e *ControlFlow) {
		V1249 := __e.Get(1)
		_ = V1249
		tmp16097 := PrimHead(V1249)

		__e.Return(PrimHead(tmp16097))
		return

	}, 1)

	tmp16098 := Call(__e, PrimNS2Value(symdef), symshen_4hds, tmp16096)

	_ = tmp16098

	tmp16099 := MakeNative(func(__e *ControlFlow) {
		V1254 := __e.Get(1)
		_ = V1254
		tmp16117 := PrimIsPair(V1254)

		var ifres16104 Obj

		if True == tmp16117 {
			tmp16115 := PrimHead(V1254)

			tmp16116 := PrimIsPair(tmp16115)

			var ifres16106 Obj

			if True == tmp16116 {
				tmp16113 := PrimTail(V1254)

				tmp16114 := PrimIsPair(tmp16113)

				var ifres16108 Obj

				if True == tmp16114 {
					tmp16110 := PrimTail(V1254)

					tmp16111 := PrimTail(tmp16110)

					tmp16112 := PrimEqual(Nil, tmp16111)

					var ifres16109 Obj

					if True == tmp16112 {
						ifres16109 = True

					} else {
						ifres16109 = False

					}

					ifres16108 = ifres16109

				} else {
					ifres16108 = False

				}

				var ifres16107 Obj

				if True == ifres16108 {
					ifres16107 = True

				} else {
					ifres16107 = False

				}

				ifres16106 = ifres16107

			} else {
				ifres16106 = False

			}

			var ifres16105 Obj

			if True == ifres16106 {
				ifres16105 = True

			} else {
				ifres16105 = False

			}

			ifres16104 = ifres16105

		} else {
			ifres16104 = False

		}

		if True == ifres16104 {
			tmp16101 := PrimHead(V1254)

			tmp16102 := PrimHead(tmp16101)

			tmp16103 := PrimTail(V1254)

			__e.Return(PrimCons(tmp16102, tmp16103))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.hdstream\n")))
			return
		}

	}, 1)

	tmp16118 := Call(__e, PrimNS2Value(symdef), symshen_4hdstream, tmp16099)

	_ = tmp16118

	tmp16119 := MakeNative(func(__e *ControlFlow) {
		V1255 := __e.Get(1)
		_ = V1255
		V1256 := __e.Get(2)
		_ = V1256
		tmp16120 := PrimCons(V1256, Nil)

		__e.Return(PrimCons(V1255, tmp16120))
		return

	}, 2)

	tmp16121 := Call(__e, PrimNS2Value(symdef), symshen_4comb, tmp16119)

	_ = tmp16121

	tmp16122 := MakeNative(func(__e *ControlFlow) {
		V1261 := __e.Get(1)
		_ = V1261
		tmp16140 := PrimIsPair(V1261)

		var ifres16127 Obj

		if True == tmp16140 {
			tmp16138 := PrimHead(V1261)

			tmp16139 := PrimIsPair(tmp16138)

			var ifres16129 Obj

			if True == tmp16139 {
				tmp16136 := PrimTail(V1261)

				tmp16137 := PrimIsPair(tmp16136)

				var ifres16131 Obj

				if True == tmp16137 {
					tmp16133 := PrimTail(V1261)

					tmp16134 := PrimTail(tmp16133)

					tmp16135 := PrimEqual(Nil, tmp16134)

					var ifres16132 Obj

					if True == tmp16135 {
						ifres16132 = True

					} else {
						ifres16132 = False

					}

					ifres16131 = ifres16132

				} else {
					ifres16131 = False

				}

				var ifres16130 Obj

				if True == ifres16131 {
					ifres16130 = True

				} else {
					ifres16130 = False

				}

				ifres16129 = ifres16130

			} else {
				ifres16129 = False

			}

			var ifres16128 Obj

			if True == ifres16129 {
				ifres16128 = True

			} else {
				ifres16128 = False

			}

			ifres16127 = ifres16128

		} else {
			ifres16127 = False

		}

		if True == ifres16127 {
			tmp16124 := PrimHead(V1261)

			tmp16125 := PrimTail(tmp16124)

			tmp16126 := PrimTail(V1261)

			__e.Return(PrimCons(tmp16125, tmp16126))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.tlstream\n")))
			return
		}

	}, 1)

	tmp16141 := Call(__e, PrimNS2Value(symdef), symshen_4tlstream, tmp16122)

	_ = tmp16141

	tmp16142 := MakeNative(func(__e *ControlFlow) {
		V1271 := __e.Get(1)
		_ = V1271
		V1272 := __e.Get(2)
		_ = V1272
		tmp16153 := PrimIsPair(V1271)

		var ifres16144 Obj

		if True == tmp16153 {
			tmp16151 := PrimHead(V1271)

			tmp16152 := PrimIsPair(tmp16151)

			var ifres16146 Obj

			if True == tmp16152 {
				tmp16148 := PrimHead(V1271)

				tmp16149 := PrimHead(tmp16148)

				tmp16150 := PrimEqual(tmp16149, V1272)

				var ifres16147 Obj

				if True == tmp16150 {
					ifres16147 = True

				} else {
					ifres16147 = False

				}

				ifres16146 = ifres16147

			} else {
				ifres16146 = False

			}

			var ifres16145 Obj

			if True == ifres16146 {
				ifres16145 = True

			} else {
				ifres16145 = False

			}

			ifres16144 = ifres16145

		} else {
			ifres16144 = False

		}

		if True == ifres16144 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp16154 := Call(__e, PrimNS2Value(symdef), symshen_4_ahd_2, tmp16142)

	_ = tmp16154

	tmp16155 := MakeNative(func(__e *ControlFlow) {
		V1277 := __e.Get(1)
		_ = V1277
		tmp16173 := PrimIsPair(V1277)

		var ifres16160 Obj

		if True == tmp16173 {
			tmp16171 := PrimHead(V1277)

			tmp16172 := PrimIsPair(tmp16171)

			var ifres16162 Obj

			if True == tmp16172 {
				tmp16169 := PrimTail(V1277)

				tmp16170 := PrimIsPair(tmp16169)

				var ifres16164 Obj

				if True == tmp16170 {
					tmp16166 := PrimTail(V1277)

					tmp16167 := PrimTail(tmp16166)

					tmp16168 := PrimEqual(Nil, tmp16167)

					var ifres16165 Obj

					if True == tmp16168 {
						ifres16165 = True

					} else {
						ifres16165 = False

					}

					ifres16164 = ifres16165

				} else {
					ifres16164 = False

				}

				var ifres16163 Obj

				if True == ifres16164 {
					ifres16163 = True

				} else {
					ifres16163 = False

				}

				ifres16162 = ifres16163

			} else {
				ifres16162 = False

			}

			var ifres16161 Obj

			if True == ifres16162 {
				ifres16161 = True

			} else {
				ifres16161 = False

			}

			ifres16160 = ifres16161

		} else {
			ifres16160 = False

		}

		if True == ifres16160 {
			tmp16157 := PrimHead(V1277)

			tmp16158 := PrimTail(tmp16157)

			tmp16159 := PrimTail(V1277)

			__e.Return(PrimCons(tmp16158, tmp16159))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.tls\n")))
			return
		}

	}, 1)

	tmp16174 := Call(__e, PrimNS2Value(symdef), symshen_4tls, tmp16155)

	_ = tmp16174

	tmp16175 := MakeNative(func(__e *ControlFlow) {
		V1280 := __e.Get(1)
		_ = V1280
		V1281 := __e.Get(2)
		_ = V1281
		V1282 := __e.Get(3)
		_ = V1282
		tmp16220 := PrimIsPair(V1282)

		var ifres16194 Obj

		if True == tmp16220 {
			tmp16218 := PrimHead(V1282)

			tmp16219 := PrimEqual(symshen_4pushsemantics, tmp16218)

			var ifres16196 Obj

			if True == tmp16219 {
				tmp16216 := PrimTail(V1282)

				tmp16217 := PrimIsPair(tmp16216)

				var ifres16198 Obj

				if True == tmp16217 {
					tmp16213 := PrimTail(V1282)

					tmp16214 := PrimTail(tmp16213)

					tmp16215 := PrimIsPair(tmp16214)

					var ifres16200 Obj

					if True == tmp16215 {
						tmp16209 := PrimTail(V1282)

						tmp16210 := PrimTail(tmp16209)

						tmp16211 := PrimTail(tmp16210)

						tmp16212 := PrimIsPair(tmp16211)

						var ifres16202 Obj

						if True == tmp16212 {
							tmp16204 := PrimTail(V1282)

							tmp16205 := PrimTail(tmp16204)

							tmp16206 := PrimTail(tmp16205)

							tmp16207 := PrimTail(tmp16206)

							tmp16208 := PrimEqual(Nil, tmp16207)

							var ifres16203 Obj

							if True == tmp16208 {
								ifres16203 = True

							} else {
								ifres16203 = False

							}

							ifres16202 = ifres16203

						} else {
							ifres16202 = False

						}

						var ifres16201 Obj

						if True == ifres16202 {
							ifres16201 = True

						} else {
							ifres16201 = False

						}

						ifres16200 = ifres16201

					} else {
						ifres16200 = False

					}

					var ifres16199 Obj

					if True == ifres16200 {
						ifres16199 = True

					} else {
						ifres16199 = False

					}

					ifres16198 = ifres16199

				} else {
					ifres16198 = False

				}

				var ifres16197 Obj

				if True == ifres16198 {
					ifres16197 = True

				} else {
					ifres16197 = False

				}

				ifres16196 = ifres16197

			} else {
				ifres16196 = False

			}

			var ifres16195 Obj

			if True == ifres16196 {
				ifres16195 = True

			} else {
				ifres16195 = False

			}

			ifres16194 = ifres16195

		} else {
			ifres16194 = False

		}

		if True == ifres16194 {
			tmp16185 := PrimTail(V1282)

			tmp16186 := PrimHead(tmp16185)

			tmp16187 := PrimTail(V1282)

			tmp16188 := PrimTail(tmp16187)

			tmp16189 := PrimHead(tmp16188)

			tmp16190 := PrimTail(V1282)

			tmp16191 := PrimTail(tmp16190)

			tmp16192 := PrimTail(tmp16191)

			tmp16193 := PrimHead(tmp16192)

			__e.TailApply(PrimNS2Value(symshen_4yacc_1syntax), V1280, tmp16186, tmp16189, tmp16193)
			return

		} else {
			tmp16177 := MakeNative(func(__e *ControlFlow) {
				Process := __e.Get(1)
				_ = Process
				tmp16178 := MakeNative(func(__e *ControlFlow) {
					Annotate := __e.Get(1)
					_ = Annotate
					tmp16179 := PrimCons(V1281, Nil)

					tmp16180 := PrimCons(symshen_4in_1_6, tmp16179)

					tmp16181 := PrimCons(Annotate, Nil)

					tmp16182 := PrimCons(tmp16180, tmp16181)

					__e.Return(PrimCons(symshen_4comb, tmp16182))
					return

				}, 1)

				tmp16183 := Call(__e, PrimNS2Value(symshen_4use_1type_1info), V1280, Process)

				__e.TailApply(tmp16178, tmp16183)
				return

			}, 1)

			tmp16184 := Call(__e, PrimNS2Value(symshen_4process_1yacc_1semantics), V1282)

			__e.TailApply(tmp16177, tmp16184)
			return

		}

	}, 3)

	tmp16221 := Call(__e, PrimNS2Value(symdef), symshen_4yacc_1semantics, tmp16175)

	_ = tmp16221

	tmp16222 := MakeNative(func(__e *ControlFlow) {
		V1286 := __e.Get(1)
		_ = V1286
		V1287 := __e.Get(2)
		_ = V1287
		tmp16442 := PrimIsPair(V1286)

		var ifres16231 Obj

		if True == tmp16442 {
			tmp16440 := PrimHead(V1286)

			tmp16441 := PrimEqual(sym_i, tmp16440)

			var ifres16233 Obj

			if True == tmp16441 {
				tmp16438 := PrimTail(V1286)

				tmp16439 := PrimIsPair(tmp16438)

				var ifres16235 Obj

				if True == tmp16439 {
					tmp16435 := PrimTail(V1286)

					tmp16436 := PrimHead(tmp16435)

					tmp16437 := PrimIsPair(tmp16436)

					var ifres16237 Obj

					if True == tmp16437 {
						tmp16431 := PrimTail(V1286)

						tmp16432 := PrimHead(tmp16431)

						tmp16433 := PrimHead(tmp16432)

						tmp16434 := PrimEqual(symstr, tmp16433)

						var ifres16239 Obj

						if True == tmp16434 {
							tmp16427 := PrimTail(V1286)

							tmp16428 := PrimHead(tmp16427)

							tmp16429 := PrimTail(tmp16428)

							tmp16430 := PrimIsPair(tmp16429)

							var ifres16241 Obj

							if True == tmp16430 {
								tmp16422 := PrimTail(V1286)

								tmp16423 := PrimHead(tmp16422)

								tmp16424 := PrimTail(tmp16423)

								tmp16425 := PrimHead(tmp16424)

								tmp16426 := PrimIsPair(tmp16425)

								var ifres16243 Obj

								if True == tmp16426 {
									tmp16416 := PrimTail(V1286)

									tmp16417 := PrimHead(tmp16416)

									tmp16418 := PrimTail(tmp16417)

									tmp16419 := PrimHead(tmp16418)

									tmp16420 := PrimHead(tmp16419)

									tmp16421 := PrimEqual(symlist, tmp16420)

									var ifres16245 Obj

									if True == tmp16421 {
										tmp16410 := PrimTail(V1286)

										tmp16411 := PrimHead(tmp16410)

										tmp16412 := PrimTail(tmp16411)

										tmp16413 := PrimHead(tmp16412)

										tmp16414 := PrimTail(tmp16413)

										tmp16415 := PrimIsPair(tmp16414)

										var ifres16247 Obj

										if True == tmp16415 {
											tmp16403 := PrimTail(V1286)

											tmp16404 := PrimHead(tmp16403)

											tmp16405 := PrimTail(tmp16404)

											tmp16406 := PrimHead(tmp16405)

											tmp16407 := PrimTail(tmp16406)

											tmp16408 := PrimTail(tmp16407)

											tmp16409 := PrimEqual(Nil, tmp16408)

											var ifres16249 Obj

											if True == tmp16409 {
												tmp16398 := PrimTail(V1286)

												tmp16399 := PrimHead(tmp16398)

												tmp16400 := PrimTail(tmp16399)

												tmp16401 := PrimTail(tmp16400)

												tmp16402 := PrimIsPair(tmp16401)

												var ifres16251 Obj

												if True == tmp16402 {
													tmp16392 := PrimTail(V1286)

													tmp16393 := PrimHead(tmp16392)

													tmp16394 := PrimTail(tmp16393)

													tmp16395 := PrimTail(tmp16394)

													tmp16396 := PrimTail(tmp16395)

													tmp16397 := PrimEqual(Nil, tmp16396)

													var ifres16253 Obj

													if True == tmp16397 {
														tmp16389 := PrimTail(V1286)

														tmp16390 := PrimTail(tmp16389)

														tmp16391 := PrimIsPair(tmp16390)

														var ifres16255 Obj

														if True == tmp16391 {
															tmp16385 := PrimTail(V1286)

															tmp16386 := PrimTail(tmp16385)

															tmp16387 := PrimHead(tmp16386)

															tmp16388 := PrimEqual(sym_1_1_6, tmp16387)

															var ifres16257 Obj

															if True == tmp16388 {
																tmp16381 := PrimTail(V1286)

																tmp16382 := PrimTail(tmp16381)

																tmp16383 := PrimTail(tmp16382)

																tmp16384 := PrimIsPair(tmp16383)

																var ifres16259 Obj

																if True == tmp16384 {
																	tmp16376 := PrimTail(V1286)

																	tmp16377 := PrimTail(tmp16376)

																	tmp16378 := PrimTail(tmp16377)

																	tmp16379 := PrimHead(tmp16378)

																	tmp16380 := PrimIsPair(tmp16379)

																	var ifres16261 Obj

																	if True == tmp16380 {
																		tmp16370 := PrimTail(V1286)

																		tmp16371 := PrimTail(tmp16370)

																		tmp16372 := PrimTail(tmp16371)

																		tmp16373 := PrimHead(tmp16372)

																		tmp16374 := PrimHead(tmp16373)

																		tmp16375 := PrimEqual(symstr, tmp16374)

																		var ifres16263 Obj

																		if True == tmp16375 {
																			tmp16364 := PrimTail(V1286)

																			tmp16365 := PrimTail(tmp16364)

																			tmp16366 := PrimTail(tmp16365)

																			tmp16367 := PrimHead(tmp16366)

																			tmp16368 := PrimTail(tmp16367)

																			tmp16369 := PrimIsPair(tmp16368)

																			var ifres16265 Obj

																			if True == tmp16369 {
																				tmp16357 := PrimTail(V1286)

																				tmp16358 := PrimTail(tmp16357)

																				tmp16359 := PrimTail(tmp16358)

																				tmp16360 := PrimHead(tmp16359)

																				tmp16361 := PrimTail(tmp16360)

																				tmp16362 := PrimHead(tmp16361)

																				tmp16363 := PrimIsPair(tmp16362)

																				var ifres16267 Obj

																				if True == tmp16363 {
																					tmp16349 := PrimTail(V1286)

																					tmp16350 := PrimTail(tmp16349)

																					tmp16351 := PrimTail(tmp16350)

																					tmp16352 := PrimHead(tmp16351)

																					tmp16353 := PrimTail(tmp16352)

																					tmp16354 := PrimHead(tmp16353)

																					tmp16355 := PrimHead(tmp16354)

																					tmp16356 := PrimEqual(symlist, tmp16355)

																					var ifres16269 Obj

																					if True == tmp16356 {
																						tmp16341 := PrimTail(V1286)

																						tmp16342 := PrimTail(tmp16341)

																						tmp16343 := PrimTail(tmp16342)

																						tmp16344 := PrimHead(tmp16343)

																						tmp16345 := PrimTail(tmp16344)

																						tmp16346 := PrimHead(tmp16345)

																						tmp16347 := PrimTail(tmp16346)

																						tmp16348 := PrimIsPair(tmp16347)

																						var ifres16271 Obj

																						if True == tmp16348 {
																							tmp16332 := PrimTail(V1286)

																							tmp16333 := PrimTail(tmp16332)

																							tmp16334 := PrimTail(tmp16333)

																							tmp16335 := PrimHead(tmp16334)

																							tmp16336 := PrimTail(tmp16335)

																							tmp16337 := PrimHead(tmp16336)

																							tmp16338 := PrimTail(tmp16337)

																							tmp16339 := PrimTail(tmp16338)

																							tmp16340 := PrimEqual(Nil, tmp16339)

																							var ifres16273 Obj

																							if True == tmp16340 {
																								tmp16325 := PrimTail(V1286)

																								tmp16326 := PrimTail(tmp16325)

																								tmp16327 := PrimTail(tmp16326)

																								tmp16328 := PrimHead(tmp16327)

																								tmp16329 := PrimTail(tmp16328)

																								tmp16330 := PrimTail(tmp16329)

																								tmp16331 := PrimIsPair(tmp16330)

																								var ifres16275 Obj

																								if True == tmp16331 {
																									tmp16317 := PrimTail(V1286)

																									tmp16318 := PrimTail(tmp16317)

																									tmp16319 := PrimTail(tmp16318)

																									tmp16320 := PrimHead(tmp16319)

																									tmp16321 := PrimTail(tmp16320)

																									tmp16322 := PrimTail(tmp16321)

																									tmp16323 := PrimTail(tmp16322)

																									tmp16324 := PrimEqual(Nil, tmp16323)

																									var ifres16277 Obj

																									if True == tmp16324 {
																										tmp16312 := PrimTail(V1286)

																										tmp16313 := PrimTail(tmp16312)

																										tmp16314 := PrimTail(tmp16313)

																										tmp16315 := PrimTail(tmp16314)

																										tmp16316 := PrimIsPair(tmp16315)

																										var ifres16279 Obj

																										if True == tmp16316 {
																											tmp16306 := PrimTail(V1286)

																											tmp16307 := PrimTail(tmp16306)

																											tmp16308 := PrimTail(tmp16307)

																											tmp16309 := PrimTail(tmp16308)

																											tmp16310 := PrimHead(tmp16309)

																											tmp16311 := PrimEqual(sym_j, tmp16310)

																											var ifres16281 Obj

																											if True == tmp16311 {
																												tmp16300 := PrimTail(V1286)

																												tmp16301 := PrimTail(tmp16300)

																												tmp16302 := PrimTail(tmp16301)

																												tmp16303 := PrimTail(tmp16302)

																												tmp16304 := PrimTail(tmp16303)

																												tmp16305 := PrimEqual(Nil, tmp16304)

																												var ifres16283 Obj

																												if True == tmp16305 {
																													tmp16285 := PrimTail(V1286)

																													tmp16286 := PrimHead(tmp16285)

																													tmp16287 := PrimTail(tmp16286)

																													tmp16288 := PrimHead(tmp16287)

																													tmp16289 := PrimTail(tmp16288)

																													tmp16290 := PrimHead(tmp16289)

																													tmp16291 := PrimTail(V1286)

																													tmp16292 := PrimTail(tmp16291)

																													tmp16293 := PrimTail(tmp16292)

																													tmp16294 := PrimHead(tmp16293)

																													tmp16295 := PrimTail(tmp16294)

																													tmp16296 := PrimHead(tmp16295)

																													tmp16297 := PrimTail(tmp16296)

																													tmp16298 := PrimHead(tmp16297)

																													tmp16299 := PrimEqual(tmp16290, tmp16298)

																													var ifres16284 Obj

																													if True == tmp16299 {
																														ifres16284 = True

																													} else {
																														ifres16284 = False

																													}

																													ifres16283 = ifres16284

																												} else {
																													ifres16283 = False

																												}

																												var ifres16282 Obj

																												if True == ifres16283 {
																													ifres16282 = True

																												} else {
																													ifres16282 = False

																												}

																												ifres16281 = ifres16282

																											} else {
																												ifres16281 = False

																											}

																											var ifres16280 Obj

																											if True == ifres16281 {
																												ifres16280 = True

																											} else {
																												ifres16280 = False

																											}

																											ifres16279 = ifres16280

																										} else {
																											ifres16279 = False

																										}

																										var ifres16278 Obj

																										if True == ifres16279 {
																											ifres16278 = True

																										} else {
																											ifres16278 = False

																										}

																										ifres16277 = ifres16278

																									} else {
																										ifres16277 = False

																									}

																									var ifres16276 Obj

																									if True == ifres16277 {
																										ifres16276 = True

																									} else {
																										ifres16276 = False

																									}

																									ifres16275 = ifres16276

																								} else {
																									ifres16275 = False

																								}

																								var ifres16274 Obj

																								if True == ifres16275 {
																									ifres16274 = True

																								} else {
																									ifres16274 = False

																								}

																								ifres16273 = ifres16274

																							} else {
																								ifres16273 = False

																							}

																							var ifres16272 Obj

																							if True == ifres16273 {
																								ifres16272 = True

																							} else {
																								ifres16272 = False

																							}

																							ifres16271 = ifres16272

																						} else {
																							ifres16271 = False

																						}

																						var ifres16270 Obj

																						if True == ifres16271 {
																							ifres16270 = True

																						} else {
																							ifres16270 = False

																						}

																						ifres16269 = ifres16270

																					} else {
																						ifres16269 = False

																					}

																					var ifres16268 Obj

																					if True == ifres16269 {
																						ifres16268 = True

																					} else {
																						ifres16268 = False

																					}

																					ifres16267 = ifres16268

																				} else {
																					ifres16267 = False

																				}

																				var ifres16266 Obj

																				if True == ifres16267 {
																					ifres16266 = True

																				} else {
																					ifres16266 = False

																				}

																				ifres16265 = ifres16266

																			} else {
																				ifres16265 = False

																			}

																			var ifres16264 Obj

																			if True == ifres16265 {
																				ifres16264 = True

																			} else {
																				ifres16264 = False

																			}

																			ifres16263 = ifres16264

																		} else {
																			ifres16263 = False

																		}

																		var ifres16262 Obj

																		if True == ifres16263 {
																			ifres16262 = True

																		} else {
																			ifres16262 = False

																		}

																		ifres16261 = ifres16262

																	} else {
																		ifres16261 = False

																	}

																	var ifres16260 Obj

																	if True == ifres16261 {
																		ifres16260 = True

																	} else {
																		ifres16260 = False

																	}

																	ifres16259 = ifres16260

																} else {
																	ifres16259 = False

																}

																var ifres16258 Obj

																if True == ifres16259 {
																	ifres16258 = True

																} else {
																	ifres16258 = False

																}

																ifres16257 = ifres16258

															} else {
																ifres16257 = False

															}

															var ifres16256 Obj

															if True == ifres16257 {
																ifres16256 = True

															} else {
																ifres16256 = False

															}

															ifres16255 = ifres16256

														} else {
															ifres16255 = False

														}

														var ifres16254 Obj

														if True == ifres16255 {
															ifres16254 = True

														} else {
															ifres16254 = False

														}

														ifres16253 = ifres16254

													} else {
														ifres16253 = False

													}

													var ifres16252 Obj

													if True == ifres16253 {
														ifres16252 = True

													} else {
														ifres16252 = False

													}

													ifres16251 = ifres16252

												} else {
													ifres16251 = False

												}

												var ifres16250 Obj

												if True == ifres16251 {
													ifres16250 = True

												} else {
													ifres16250 = False

												}

												ifres16249 = ifres16250

											} else {
												ifres16249 = False

											}

											var ifres16248 Obj

											if True == ifres16249 {
												ifres16248 = True

											} else {
												ifres16248 = False

											}

											ifres16247 = ifres16248

										} else {
											ifres16247 = False

										}

										var ifres16246 Obj

										if True == ifres16247 {
											ifres16246 = True

										} else {
											ifres16246 = False

										}

										ifres16245 = ifres16246

									} else {
										ifres16245 = False

									}

									var ifres16244 Obj

									if True == ifres16245 {
										ifres16244 = True

									} else {
										ifres16244 = False

									}

									ifres16243 = ifres16244

								} else {
									ifres16243 = False

								}

								var ifres16242 Obj

								if True == ifres16243 {
									ifres16242 = True

								} else {
									ifres16242 = False

								}

								ifres16241 = ifres16242

							} else {
								ifres16241 = False

							}

							var ifres16240 Obj

							if True == ifres16241 {
								ifres16240 = True

							} else {
								ifres16240 = False

							}

							ifres16239 = ifres16240

						} else {
							ifres16239 = False

						}

						var ifres16238 Obj

						if True == ifres16239 {
							ifres16238 = True

						} else {
							ifres16238 = False

						}

						ifres16237 = ifres16238

					} else {
						ifres16237 = False

					}

					var ifres16236 Obj

					if True == ifres16237 {
						ifres16236 = True

					} else {
						ifres16236 = False

					}

					ifres16235 = ifres16236

				} else {
					ifres16235 = False

				}

				var ifres16234 Obj

				if True == ifres16235 {
					ifres16234 = True

				} else {
					ifres16234 = False

				}

				ifres16233 = ifres16234

			} else {
				ifres16233 = False

			}

			var ifres16232 Obj

			if True == ifres16233 {
				ifres16232 = True

			} else {
				ifres16232 = False

			}

			ifres16231 = ifres16232

		} else {
			ifres16231 = False

		}

		if True == ifres16231 {
			tmp16224 := PrimTail(V1286)

			tmp16225 := PrimTail(tmp16224)

			tmp16226 := PrimTail(tmp16225)

			tmp16227 := PrimHead(tmp16226)

			tmp16228 := PrimTail(tmp16227)

			tmp16229 := PrimTail(tmp16228)

			tmp16230 := PrimCons(V1287, tmp16229)

			__e.Return(PrimCons(symtype, tmp16230))
			return

		} else {
			__e.Return(V1287)
			return
		}

	}, 2)

	tmp16443 := Call(__e, PrimNS2Value(symdef), symshen_4use_1type_1info, tmp16222)

	_ = tmp16443

	tmp16444 := MakeNative(func(__e *ControlFlow) {
		V1288 := __e.Get(1)
		_ = V1288
		tmp16452 := PrimIsPair(V1288)

		if True == tmp16452 {
			tmp16451 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimNS2Value(symshen_4process_1yacc_1semantics), Z)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp16451, V1288)
			return

		} else {
			tmp16450 := Call(__e, PrimNS2Value(symshen_4non_1terminal_2), V1288)

			if True == tmp16450 {
				tmp16447 := Call(__e, PrimNS2Value(symprotect), symParse)

				tmp16448 := Call(__e, PrimNS2Value(symconcat), tmp16447, V1288)

				tmp16449 := PrimCons(tmp16448, Nil)

				__e.Return(PrimCons(symshen_4_5_1out, tmp16449))
				return

			} else {
				__e.Return(V1288)
				return
			}

		}

	}, 1)

	tmp16453 := Call(__e, PrimNS2Value(symdef), symshen_4process_1yacc_1semantics, tmp16444)

	_ = tmp16453

	tmp16454 := MakeNative(func(__e *ControlFlow) {
		V1293 := __e.Get(1)
		_ = V1293
		tmp16466 := PrimIsPair(V1293)

		var ifres16457 Obj

		if True == tmp16466 {
			tmp16464 := PrimTail(V1293)

			tmp16465 := PrimIsPair(tmp16464)

			var ifres16459 Obj

			if True == tmp16465 {
				tmp16461 := PrimTail(V1293)

				tmp16462 := PrimTail(tmp16461)

				tmp16463 := PrimEqual(Nil, tmp16462)

				var ifres16460 Obj

				if True == tmp16463 {
					ifres16460 = True

				} else {
					ifres16460 = False

				}

				ifres16459 = ifres16460

			} else {
				ifres16459 = False

			}

			var ifres16458 Obj

			if True == ifres16459 {
				ifres16458 = True

			} else {
				ifres16458 = False

			}

			ifres16457 = ifres16458

		} else {
			ifres16457 = False

		}

		if True == ifres16457 {
			tmp16456 := PrimTail(V1293)

			__e.Return(PrimHead(tmp16456))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.<-out\n")))
			return
		}

	}, 1)

	tmp16467 := Call(__e, PrimNS2Value(symdef), symshen_4_5_1out, tmp16454)

	_ = tmp16467

	tmp16468 := MakeNative(func(__e *ControlFlow) {
		V1298 := __e.Get(1)
		_ = V1298
		tmp16479 := PrimIsPair(V1298)

		var ifres16470 Obj

		if True == tmp16479 {
			tmp16477 := PrimTail(V1298)

			tmp16478 := PrimIsPair(tmp16477)

			var ifres16472 Obj

			if True == tmp16478 {
				tmp16474 := PrimTail(V1298)

				tmp16475 := PrimTail(tmp16474)

				tmp16476 := PrimEqual(Nil, tmp16475)

				var ifres16473 Obj

				if True == tmp16476 {
					ifres16473 = True

				} else {
					ifres16473 = False

				}

				ifres16472 = ifres16473

			} else {
				ifres16472 = False

			}

			var ifres16471 Obj

			if True == ifres16472 {
				ifres16471 = True

			} else {
				ifres16471 = False

			}

			ifres16470 = ifres16471

		} else {
			ifres16470 = False

		}

		if True == ifres16470 {
			__e.Return(PrimHead(V1298))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.in->\n")))
			return
		}

	}, 1)

	tmp16480 := Call(__e, PrimNS2Value(symdef), symshen_4in_1_6, tmp16468)

	_ = tmp16480

	tmp16481 := MakeNative(func(__e *ControlFlow) {
		V1303 := __e.Get(1)
		_ = V1303
		tmp16494 := PrimIsPair(V1303)

		var ifres16485 Obj

		if True == tmp16494 {
			tmp16492 := PrimTail(V1303)

			tmp16493 := PrimIsPair(tmp16492)

			var ifres16487 Obj

			if True == tmp16493 {
				tmp16489 := PrimTail(V1303)

				tmp16490 := PrimTail(tmp16489)

				tmp16491 := PrimEqual(Nil, tmp16490)

				var ifres16488 Obj

				if True == tmp16491 {
					ifres16488 = True

				} else {
					ifres16488 = False

				}

				ifres16487 = ifres16488

			} else {
				ifres16487 = False

			}

			var ifres16486 Obj

			if True == ifres16487 {
				ifres16486 = True

			} else {
				ifres16486 = False

			}

			ifres16485 = ifres16486

		} else {
			ifres16485 = False

		}

		if True == ifres16485 {
			tmp16483 := PrimHead(V1303)

			tmp16484 := PrimCons(tmp16483, Nil)

			__e.Return(PrimCons(Nil, tmp16484))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in <!>\n")))
			return
		}

	}, 1)

	tmp16495 := Call(__e, PrimNS2Value(symdef), sym_5_b_6, tmp16481)

	_ = tmp16495

	tmp16496 := MakeNative(func(__e *ControlFlow) {
		V1308 := __e.Get(1)
		_ = V1308
		tmp16509 := PrimIsPair(V1308)

		var ifres16500 Obj

		if True == tmp16509 {
			tmp16507 := PrimTail(V1308)

			tmp16508 := PrimIsPair(tmp16507)

			var ifres16502 Obj

			if True == tmp16508 {
				tmp16504 := PrimTail(V1308)

				tmp16505 := PrimTail(tmp16504)

				tmp16506 := PrimEqual(Nil, tmp16505)

				var ifres16503 Obj

				if True == tmp16506 {
					ifres16503 = True

				} else {
					ifres16503 = False

				}

				ifres16502 = ifres16503

			} else {
				ifres16502 = False

			}

			var ifres16501 Obj

			if True == ifres16502 {
				ifres16501 = True

			} else {
				ifres16501 = False

			}

			ifres16500 = ifres16501

		} else {
			ifres16500 = False

		}

		if True == ifres16500 {
			tmp16498 := PrimHead(V1308)

			tmp16499 := PrimCons(Nil, Nil)

			__e.Return(PrimCons(tmp16498, tmp16499))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in <e>\n")))
			return
		}

	}, 1)

	tmp16510 := Call(__e, PrimNS2Value(symdef), sym_5e_6, tmp16496)

	_ = tmp16510

	tmp16511 := MakeNative(func(__e *ControlFlow) {
		V1311 := __e.Get(1)
		_ = V1311
		tmp16526 := PrimIsPair(V1311)

		var ifres16513 Obj

		if True == tmp16526 {
			tmp16524 := PrimHead(V1311)

			tmp16525 := PrimEqual(Nil, tmp16524)

			var ifres16515 Obj

			if True == tmp16525 {
				tmp16522 := PrimTail(V1311)

				tmp16523 := PrimIsPair(tmp16522)

				var ifres16517 Obj

				if True == tmp16523 {
					tmp16519 := PrimTail(V1311)

					tmp16520 := PrimTail(tmp16519)

					tmp16521 := PrimEqual(Nil, tmp16520)

					var ifres16518 Obj

					if True == tmp16521 {
						ifres16518 = True

					} else {
						ifres16518 = False

					}

					ifres16517 = ifres16518

				} else {
					ifres16517 = False

				}

				var ifres16516 Obj

				if True == ifres16517 {
					ifres16516 = True

				} else {
					ifres16516 = False

				}

				ifres16515 = ifres16516

			} else {
				ifres16515 = False

			}

			var ifres16514 Obj

			if True == ifres16515 {
				ifres16514 = True

			} else {
				ifres16514 = False

			}

			ifres16513 = ifres16514

		} else {
			ifres16513 = False

		}

		if True == ifres16513 {
			__e.Return(V1311)
			return
		} else {
			__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
			return
		}

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), symshen_4_5end_6, tmp16511)
	return

}, 0)
