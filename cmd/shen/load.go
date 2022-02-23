package main

import . "github.com/tiancaiamao/shen-go/klambda"

var LoadMain = MakeNative(func(__e *ControlFlow) {
	tmp8189 := MakeNative(func(__e *ControlFlow) {
		V1704 := __e.Get(1)
		_ = V1704
		tmp8190 := MakeNative(func(__e *ControlFlow) {
			TC_2 := __e.Get(1)
			_ = TC_2
			tmp8191 := MakeNative(func(__e *ControlFlow) {
				Load := __e.Get(1)
				_ = Load
				tmp8192 := MakeNative(func(__e *ControlFlow) {
					Infs := __e.Get(1)
					_ = Infs
					__e.Return(symloaded)
					return
				}, 1)

				var ifres8193 Obj

				if True == TC_2 {
					tmp8194 := Call(__e, PrimNS2Value(syminferences))

					tmp8195 := Call(__e, PrimNS2Value(symshen_4app), tmp8194, MakeString(" inferences\n"), symshen_4a)

					tmp8196 := PrimStringConcat(MakeString("\ntypechecked in "), tmp8195)

					tmp8197 := Call(__e, PrimNS2Value(symstoutput))

					tmp8198 := Call(__e, PrimNS2Value(sympr), tmp8196, tmp8197)

					ifres8193 = tmp8198

				} else {
					ifres8193 = symshen_4skip

				}

				__e.TailApply(tmp8192, ifres8193)
				return

			}, 1)

			tmp8199 := MakeNative(func(__e *ControlFlow) {
				Start := __e.Get(1)
				_ = Start
				tmp8200 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp8201 := MakeNative(func(__e *ControlFlow) {
						Finish := __e.Get(1)
						_ = Finish
						tmp8202 := MakeNative(func(__e *ControlFlow) {
							Time := __e.Get(1)
							_ = Time
							tmp8203 := MakeNative(func(__e *ControlFlow) {
								Message := __e.Get(1)
								_ = Message
								__e.Return(Result)
								return
							}, 1)

							tmp8204 := PrimStr(Time)

							tmp8205 := PrimStringConcat(tmp8204, MakeString(" secs\n"))

							tmp8206 := PrimStringConcat(MakeString("\nrun time: "), tmp8205)

							tmp8207 := Call(__e, PrimNS2Value(symstoutput))

							tmp8208 := Call(__e, PrimNS2Value(sympr), tmp8206, tmp8207)

							__e.TailApply(tmp8203, tmp8208)
							return

						}, 1)

						tmp8209 := PrimNumberSubtract(Finish, Start)

						__e.TailApply(tmp8202, tmp8209)
						return

					}, 1)

					tmp8210 := PrimGetTime(symrun)

					__e.TailApply(tmp8201, tmp8210)
					return

				}, 1)

				tmp8211 := Call(__e, PrimNS2Value(symread_1file), V1704)

				tmp8212 := Call(__e, PrimNS2Value(symshen_4load_1help), TC_2, tmp8211)

				__e.TailApply(tmp8200, tmp8212)
				return

			}, 1)

			tmp8213 := PrimGetTime(symrun)

			tmp8214 := Call(__e, tmp8199, tmp8213)

			__e.TailApply(tmp8191, tmp8214)
			return

		}, 1)

		tmp8215 := PrimNS3Value(symshen_4_dtc_d)

		__e.TailApply(tmp8190, tmp8215)
		return

	}, 1)

	tmp8216 := Call(__e, PrimNS2Value(symdef), symload, tmp8189)

	_ = tmp8216

	tmp8217 := MakeNative(func(__e *ControlFlow) {
		V1707 := __e.Get(1)
		_ = V1707
		V1708 := __e.Get(2)
		_ = V1708
		tmp8219 := PrimEqual(False, V1707)

		if True == tmp8219 {
			__e.TailApply(PrimNS2Value(symshen_4eval_1and_1print), V1708)
			return
		} else {
			__e.TailApply(PrimNS2Value(symshen_4check_1eval_1and_1print), V1708)
			return
		}

	}, 2)

	tmp8220 := Call(__e, PrimNS2Value(symdef), symshen_4load_1help, tmp8217)

	_ = tmp8220

	tmp8221 := MakeNative(func(__e *ControlFlow) {
		V1709 := __e.Get(1)
		_ = V1709
		tmp8222 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			tmp8223 := Call(__e, PrimNS2Value(symshen_4shen_1_6kl), Y)

			tmp8224 := Call(__e, PrimNS2Value(symeval_1kl), tmp8223)

			tmp8225 := Call(__e, PrimNS2Value(symshen_4app), tmp8224, MakeString("\n"), symshen_4s)

			tmp8226 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp8225, tmp8226)
			return

		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp8222, V1709)
		return

	}, 1)

	tmp8227 := Call(__e, PrimNS2Value(symdef), symshen_4eval_1and_1print, tmp8221)

	_ = tmp8227

	tmp8228 := MakeNative(func(__e *ControlFlow) {
		V1710 := __e.Get(1)
		_ = V1710
		tmp8229 := MakeNative(func(__e *ControlFlow) {
			Table := __e.Get(1)
			_ = Table
			tmp8230 := MakeNative(func(__e *ControlFlow) {
				Assume := __e.Get(1)
				_ = Assume
				tmp8231 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimNS2Value(symshen_4work_1through), V1710)
					return
				}, 0)

				tmp8232 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.TailApply(PrimNS2Value(symshen_4unwind_1types), E, Table)
					return
				}, 1)

				__e.TailApply(PrimNS1Value(symtry_1catch), tmp8231, tmp8232)
				return

			}, 1)

			tmp8233 := Call(__e, PrimNS2Value(symshen_4assumetypes), Table)

			__e.TailApply(tmp8230, tmp8233)
			return

		}, 1)

		tmp8234 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			__e.TailApply(PrimNS2Value(symshen_4typetable), Y)
			return
		}, 1)

		tmp8235 := Call(__e, PrimNS2Value(symmapcan), tmp8234, V1710)

		__e.TailApply(tmp8229, tmp8235)
		return

	}, 1)

	tmp8236 := Call(__e, PrimNS2Value(symdef), symshen_4check_1eval_1and_1print, tmp8228)

	_ = tmp8236

	tmp8237 := MakeNative(func(__e *ControlFlow) {
		V1715 := __e.Get(1)
		_ = V1715
		tmp8282 := PrimIsPair(V1715)

		var ifres8263 Obj

		if True == tmp8282 {
			tmp8280 := PrimHead(V1715)

			tmp8281 := PrimEqual(symdefine, tmp8280)

			var ifres8265 Obj

			if True == tmp8281 {
				tmp8278 := PrimTail(V1715)

				tmp8279 := PrimIsPair(tmp8278)

				var ifres8267 Obj

				if True == tmp8279 {
					tmp8275 := PrimTail(V1715)

					tmp8276 := PrimTail(tmp8275)

					tmp8277 := PrimIsPair(tmp8276)

					var ifres8269 Obj

					if True == tmp8277 {
						tmp8271 := PrimTail(V1715)

						tmp8272 := PrimTail(tmp8271)

						tmp8273 := PrimHead(tmp8272)

						tmp8274 := PrimEqual(sym_i, tmp8273)

						var ifres8270 Obj

						if True == tmp8274 {
							ifres8270 = True

						} else {
							ifres8270 = False

						}

						ifres8269 = ifres8270

					} else {
						ifres8269 = False

					}

					var ifres8268 Obj

					if True == ifres8269 {
						ifres8268 = True

					} else {
						ifres8268 = False

					}

					ifres8267 = ifres8268

				} else {
					ifres8267 = False

				}

				var ifres8266 Obj

				if True == ifres8267 {
					ifres8266 = True

				} else {
					ifres8266 = False

				}

				ifres8265 = ifres8266

			} else {
				ifres8265 = False

			}

			var ifres8264 Obj

			if True == ifres8265 {
				ifres8264 = True

			} else {
				ifres8264 = False

			}

			ifres8263 = ifres8264

		} else {
			ifres8263 = False

		}

		if True == ifres8263 {
			tmp8253 := PrimTail(V1715)

			tmp8254 := PrimHead(tmp8253)

			tmp8255 := PrimTail(V1715)

			tmp8256 := PrimHead(tmp8255)

			tmp8257 := PrimTail(V1715)

			tmp8258 := PrimTail(tmp8257)

			tmp8259 := PrimTail(tmp8258)

			tmp8260 := Call(__e, PrimNS2Value(symshen_4type_1F), tmp8256, tmp8259)

			tmp8261 := Call(__e, PrimNS2Value(symshen_4rectify_1type), tmp8260)

			tmp8262 := PrimCons(tmp8261, Nil)

			__e.Return(PrimCons(tmp8254, tmp8262))
			return

		} else {
			tmp8252 := PrimIsPair(V1715)

			var ifres8244 Obj

			if True == tmp8252 {
				tmp8250 := PrimHead(V1715)

				tmp8251 := PrimEqual(symdefine, tmp8250)

				var ifres8246 Obj

				if True == tmp8251 {
					tmp8248 := PrimTail(V1715)

					tmp8249 := PrimIsPair(tmp8248)

					var ifres8247 Obj

					if True == tmp8249 {
						ifres8247 = True

					} else {
						ifres8247 = False

					}

					ifres8246 = ifres8247

				} else {
					ifres8246 = False

				}

				var ifres8245 Obj

				if True == ifres8246 {
					ifres8245 = True

				} else {
					ifres8245 = False

				}

				ifres8244 = ifres8245

			} else {
				ifres8244 = False

			}

			if True == ifres8244 {
				tmp8240 := PrimTail(V1715)

				tmp8241 := PrimHead(tmp8240)

				tmp8242 := Call(__e, PrimNS2Value(symshen_4app), tmp8241, MakeString("\n"), symshen_4a)

				tmp8243 := PrimStringConcat(MakeString("missing { in "), tmp8242)

				__e.Return(PrimSimpleError(tmp8243))
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp8283 := Call(__e, PrimNS2Value(symdef), symshen_4typetable, tmp8237)

	_ = tmp8283

	tmp8284 := MakeNative(func(__e *ControlFlow) {
		V1722 := __e.Get(1)
		_ = V1722
		V1723 := __e.Get(2)
		_ = V1723
		tmp8297 := PrimIsPair(V1723)

		var ifres8293 Obj

		if True == tmp8297 {
			tmp8295 := PrimHead(V1723)

			tmp8296 := PrimEqual(sym_j, tmp8295)

			var ifres8294 Obj

			if True == tmp8296 {
				ifres8294 = True

			} else {
				ifres8294 = False

			}

			ifres8293 = ifres8294

		} else {
			ifres8293 = False

		}

		if True == ifres8293 {
			__e.Return(Nil)
			return
		} else {
			tmp8292 := PrimIsPair(V1723)

			if True == tmp8292 {
				tmp8289 := PrimHead(V1723)

				tmp8290 := PrimTail(V1723)

				tmp8291 := Call(__e, PrimNS2Value(symshen_4type_1F), V1722, tmp8290)

				__e.Return(PrimCons(tmp8289, tmp8291))
				return

			} else {
				tmp8287 := Call(__e, PrimNS2Value(symshen_4app), V1722, MakeString("\n"), symshen_4a)

				tmp8288 := PrimStringConcat(MakeString("missing } in "), tmp8287)

				__e.Return(PrimSimpleError(tmp8288))
				return

			}

		}

	}, 2)

	tmp8298 := Call(__e, PrimNS2Value(symdef), symshen_4type_1F, tmp8284)

	_ = tmp8298

	tmp8299 := MakeNative(func(__e *ControlFlow) {
		V1726 := __e.Get(1)
		_ = V1726
		tmp8313 := PrimEqual(Nil, V1726)

		if True == tmp8313 {
			__e.Return(Nil)
			return
		} else {
			tmp8312 := PrimIsPair(V1726)

			var ifres8308 Obj

			if True == tmp8312 {
				tmp8310 := PrimTail(V1726)

				tmp8311 := PrimIsPair(tmp8310)

				var ifres8309 Obj

				if True == tmp8311 {
					ifres8309 = True

				} else {
					ifres8309 = False

				}

				ifres8308 = ifres8309

			} else {
				ifres8308 = False

			}

			if True == ifres8308 {
				tmp8302 := PrimHead(V1726)

				tmp8303 := PrimTail(V1726)

				tmp8304 := PrimHead(tmp8303)

				tmp8305 := Call(__e, PrimNS2Value(symdeclare), tmp8302, tmp8304)

				_ = tmp8305

				tmp8306 := PrimTail(V1726)

				tmp8307 := PrimTail(tmp8306)

				__e.TailApply(PrimNS2Value(symshen_4assumetypes), tmp8307)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
				return
			}

		}

	}, 1)

	tmp8314 := Call(__e, PrimNS2Value(symdef), symshen_4assumetypes, tmp8299)

	_ = tmp8314

	tmp8315 := MakeNative(func(__e *ControlFlow) {
		V1731 := __e.Get(1)
		_ = V1731
		V1732 := __e.Get(2)
		_ = V1732
		tmp8326 := PrimIsPair(V1732)

		var ifres8322 Obj

		if True == tmp8326 {
			tmp8324 := PrimHead(V1732)

			tmp8325 := PrimIsPair(tmp8324)

			var ifres8323 Obj

			if True == tmp8325 {
				ifres8323 = True

			} else {
				ifres8323 = False

			}

			ifres8322 = ifres8323

		} else {
			ifres8322 = False

		}

		if True == ifres8322 {
			tmp8318 := PrimHead(V1732)

			tmp8319 := PrimHead(tmp8318)

			tmp8320 := Call(__e, PrimNS2Value(symdestroy), tmp8319)

			_ = tmp8320

			tmp8321 := PrimTail(V1732)

			__e.TailApply(PrimNS2Value(symshen_4unwind_1types), V1731, tmp8321)
			return

		} else {
			tmp8317 := PrimErrorToString(V1731)

			__e.Return(PrimSimpleError(tmp8317))
			return

		}

	}, 2)

	tmp8327 := Call(__e, PrimNS2Value(symdef), symshen_4unwind_1types, tmp8315)

	_ = tmp8327

	tmp8328 := MakeNative(func(__e *ControlFlow) {
		V1735 := __e.Get(1)
		_ = V1735
		tmp8378 := PrimEqual(Nil, V1735)

		if True == tmp8378 {
			__e.Return(Nil)
			return
		} else {
			tmp8377 := PrimIsPair(V1735)

			var ifres8362 Obj

			if True == tmp8377 {
				tmp8375 := PrimTail(V1735)

				tmp8376 := PrimIsPair(tmp8375)

				var ifres8364 Obj

				if True == tmp8376 {
					tmp8372 := PrimTail(V1735)

					tmp8373 := PrimTail(tmp8372)

					tmp8374 := PrimIsPair(tmp8373)

					var ifres8366 Obj

					if True == tmp8374 {
						tmp8368 := PrimTail(V1735)

						tmp8369 := PrimHead(tmp8368)

						tmp8370 := PrimIntern(MakeString(":"))

						tmp8371 := PrimEqual(tmp8369, tmp8370)

						var ifres8367 Obj

						if True == tmp8371 {
							ifres8367 = True

						} else {
							ifres8367 = False

						}

						ifres8366 = ifres8367

					} else {
						ifres8366 = False

					}

					var ifres8365 Obj

					if True == ifres8366 {
						ifres8365 = True

					} else {
						ifres8365 = False

					}

					ifres8364 = ifres8365

				} else {
					ifres8364 = False

				}

				var ifres8363 Obj

				if True == ifres8364 {
					ifres8363 = True

				} else {
					ifres8363 = False

				}

				ifres8362 = ifres8363

			} else {
				ifres8362 = False

			}

			if True == ifres8362 {
				tmp8340 := MakeNative(func(__e *ControlFlow) {
					Check := __e.Get(1)
					_ = Check
					tmp8356 := PrimEqual(Check, False)

					if True == tmp8356 {
						__e.TailApply(PrimNS2Value(symshen_4type_1error))
						return
					} else {
						tmp8342 := MakeNative(func(__e *ControlFlow) {
							Eval := __e.Get(1)
							_ = Eval
							tmp8343 := MakeNative(func(__e *ControlFlow) {
								Message := __e.Get(1)
								_ = Message
								tmp8344 := PrimTail(V1735)

								tmp8345 := PrimTail(tmp8344)

								tmp8346 := PrimTail(tmp8345)

								__e.TailApply(PrimNS2Value(symshen_4work_1through), tmp8346)
								return

							}, 1)

							tmp8347 := Call(__e, PrimNS2Value(symshen_4pretty_1type), Check)

							tmp8348 := Call(__e, PrimNS2Value(symshen_4app), tmp8347, MakeString("\n"), symshen_4r)

							tmp8349 := PrimStringConcat(MakeString(" : "), tmp8348)

							tmp8350 := Call(__e, PrimNS2Value(symshen_4app), Eval, tmp8349, symshen_4s)

							tmp8351 := Call(__e, PrimNS2Value(symstoutput))

							tmp8352 := Call(__e, PrimNS2Value(sympr), tmp8350, tmp8351)

							__e.TailApply(tmp8343, tmp8352)
							return

						}, 1)

						tmp8353 := PrimHead(V1735)

						tmp8354 := Call(__e, PrimNS2Value(symshen_4shen_1_6kl), tmp8353)

						tmp8355 := Call(__e, PrimNS2Value(symeval_1kl), tmp8354)

						__e.TailApply(tmp8342, tmp8355)
						return

					}

				}, 1)

				tmp8357 := PrimHead(V1735)

				tmp8358 := PrimTail(V1735)

				tmp8359 := PrimTail(tmp8358)

				tmp8360 := PrimHead(tmp8359)

				tmp8361 := Call(__e, PrimNS2Value(symshen_4typecheck), tmp8357, tmp8360)

				__e.TailApply(tmp8340, tmp8361)
				return

			} else {
				tmp8339 := PrimIsPair(V1735)

				if True == tmp8339 {
					tmp8332 := PrimHead(V1735)

					tmp8333 := PrimIntern(MakeString(":"))

					tmp8334 := Call(__e, PrimNS2Value(symprotect), symA)

					tmp8335 := PrimTail(V1735)

					tmp8336 := PrimCons(tmp8334, tmp8335)

					tmp8337 := PrimCons(tmp8333, tmp8336)

					tmp8338 := PrimCons(tmp8332, tmp8337)

					__e.TailApply(PrimNS2Value(symshen_4work_1through), tmp8338)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
					return
				}

			}

		}

	}, 1)

	tmp8379 := Call(__e, PrimNS2Value(symdef), symshen_4work_1through, tmp8328)

	_ = tmp8379

	tmp8380 := MakeNative(func(__e *ControlFlow) {
		V1737 := __e.Get(1)
		_ = V1737
		tmp8556 := PrimIsPair(V1737)

		var ifres8393 Obj

		if True == tmp8556 {
			tmp8554 := PrimHead(V1737)

			tmp8555 := PrimIsPair(tmp8554)

			var ifres8395 Obj

			if True == tmp8555 {
				tmp8551 := PrimHead(V1737)

				tmp8552 := PrimHead(tmp8551)

				tmp8553 := PrimEqual(symstr, tmp8552)

				var ifres8397 Obj

				if True == tmp8553 {
					tmp8548 := PrimHead(V1737)

					tmp8549 := PrimTail(tmp8548)

					tmp8550 := PrimIsPair(tmp8549)

					var ifres8399 Obj

					if True == tmp8550 {
						tmp8544 := PrimHead(V1737)

						tmp8545 := PrimTail(tmp8544)

						tmp8546 := PrimHead(tmp8545)

						tmp8547 := PrimIsPair(tmp8546)

						var ifres8401 Obj

						if True == tmp8547 {
							tmp8539 := PrimHead(V1737)

							tmp8540 := PrimTail(tmp8539)

							tmp8541 := PrimHead(tmp8540)

							tmp8542 := PrimHead(tmp8541)

							tmp8543 := PrimEqual(symlist, tmp8542)

							var ifres8403 Obj

							if True == tmp8543 {
								tmp8534 := PrimHead(V1737)

								tmp8535 := PrimTail(tmp8534)

								tmp8536 := PrimHead(tmp8535)

								tmp8537 := PrimTail(tmp8536)

								tmp8538 := PrimIsPair(tmp8537)

								var ifres8405 Obj

								if True == tmp8538 {
									tmp8528 := PrimHead(V1737)

									tmp8529 := PrimTail(tmp8528)

									tmp8530 := PrimHead(tmp8529)

									tmp8531 := PrimTail(tmp8530)

									tmp8532 := PrimTail(tmp8531)

									tmp8533 := PrimEqual(Nil, tmp8532)

									var ifres8407 Obj

									if True == tmp8533 {
										tmp8524 := PrimHead(V1737)

										tmp8525 := PrimTail(tmp8524)

										tmp8526 := PrimTail(tmp8525)

										tmp8527 := PrimIsPair(tmp8526)

										var ifres8409 Obj

										if True == tmp8527 {
											tmp8519 := PrimHead(V1737)

											tmp8520 := PrimTail(tmp8519)

											tmp8521 := PrimTail(tmp8520)

											tmp8522 := PrimTail(tmp8521)

											tmp8523 := PrimEqual(Nil, tmp8522)

											var ifres8411 Obj

											if True == tmp8523 {
												tmp8517 := PrimTail(V1737)

												tmp8518 := PrimIsPair(tmp8517)

												var ifres8413 Obj

												if True == tmp8518 {
													tmp8514 := PrimTail(V1737)

													tmp8515 := PrimHead(tmp8514)

													tmp8516 := PrimEqual(sym_1_1_6, tmp8515)

													var ifres8415 Obj

													if True == tmp8516 {
														tmp8511 := PrimTail(V1737)

														tmp8512 := PrimTail(tmp8511)

														tmp8513 := PrimIsPair(tmp8512)

														var ifres8417 Obj

														if True == tmp8513 {
															tmp8507 := PrimTail(V1737)

															tmp8508 := PrimTail(tmp8507)

															tmp8509 := PrimHead(tmp8508)

															tmp8510 := PrimIsPair(tmp8509)

															var ifres8419 Obj

															if True == tmp8510 {
																tmp8502 := PrimTail(V1737)

																tmp8503 := PrimTail(tmp8502)

																tmp8504 := PrimHead(tmp8503)

																tmp8505 := PrimHead(tmp8504)

																tmp8506 := PrimEqual(symstr, tmp8505)

																var ifres8421 Obj

																if True == tmp8506 {
																	tmp8497 := PrimTail(V1737)

																	tmp8498 := PrimTail(tmp8497)

																	tmp8499 := PrimHead(tmp8498)

																	tmp8500 := PrimTail(tmp8499)

																	tmp8501 := PrimIsPair(tmp8500)

																	var ifres8423 Obj

																	if True == tmp8501 {
																		tmp8491 := PrimTail(V1737)

																		tmp8492 := PrimTail(tmp8491)

																		tmp8493 := PrimHead(tmp8492)

																		tmp8494 := PrimTail(tmp8493)

																		tmp8495 := PrimHead(tmp8494)

																		tmp8496 := PrimIsPair(tmp8495)

																		var ifres8425 Obj

																		if True == tmp8496 {
																			tmp8484 := PrimTail(V1737)

																			tmp8485 := PrimTail(tmp8484)

																			tmp8486 := PrimHead(tmp8485)

																			tmp8487 := PrimTail(tmp8486)

																			tmp8488 := PrimHead(tmp8487)

																			tmp8489 := PrimHead(tmp8488)

																			tmp8490 := PrimEqual(symlist, tmp8489)

																			var ifres8427 Obj

																			if True == tmp8490 {
																				tmp8477 := PrimTail(V1737)

																				tmp8478 := PrimTail(tmp8477)

																				tmp8479 := PrimHead(tmp8478)

																				tmp8480 := PrimTail(tmp8479)

																				tmp8481 := PrimHead(tmp8480)

																				tmp8482 := PrimTail(tmp8481)

																				tmp8483 := PrimIsPair(tmp8482)

																				var ifres8429 Obj

																				if True == tmp8483 {
																					tmp8469 := PrimTail(V1737)

																					tmp8470 := PrimTail(tmp8469)

																					tmp8471 := PrimHead(tmp8470)

																					tmp8472 := PrimTail(tmp8471)

																					tmp8473 := PrimHead(tmp8472)

																					tmp8474 := PrimTail(tmp8473)

																					tmp8475 := PrimTail(tmp8474)

																					tmp8476 := PrimEqual(Nil, tmp8475)

																					var ifres8431 Obj

																					if True == tmp8476 {
																						tmp8463 := PrimTail(V1737)

																						tmp8464 := PrimTail(tmp8463)

																						tmp8465 := PrimHead(tmp8464)

																						tmp8466 := PrimTail(tmp8465)

																						tmp8467 := PrimTail(tmp8466)

																						tmp8468 := PrimIsPair(tmp8467)

																						var ifres8433 Obj

																						if True == tmp8468 {
																							tmp8456 := PrimTail(V1737)

																							tmp8457 := PrimTail(tmp8456)

																							tmp8458 := PrimHead(tmp8457)

																							tmp8459 := PrimTail(tmp8458)

																							tmp8460 := PrimTail(tmp8459)

																							tmp8461 := PrimTail(tmp8460)

																							tmp8462 := PrimEqual(Nil, tmp8461)

																							var ifres8435 Obj

																							if True == tmp8462 {
																								tmp8452 := PrimTail(V1737)

																								tmp8453 := PrimTail(tmp8452)

																								tmp8454 := PrimTail(tmp8453)

																								tmp8455 := PrimEqual(Nil, tmp8454)

																								var ifres8437 Obj

																								if True == tmp8455 {
																									tmp8439 := PrimHead(V1737)

																									tmp8440 := PrimTail(tmp8439)

																									tmp8441 := PrimHead(tmp8440)

																									tmp8442 := PrimTail(tmp8441)

																									tmp8443 := PrimHead(tmp8442)

																									tmp8444 := PrimTail(V1737)

																									tmp8445 := PrimTail(tmp8444)

																									tmp8446 := PrimHead(tmp8445)

																									tmp8447 := PrimTail(tmp8446)

																									tmp8448 := PrimHead(tmp8447)

																									tmp8449 := PrimTail(tmp8448)

																									tmp8450 := PrimHead(tmp8449)

																									tmp8451 := PrimEqual(tmp8443, tmp8450)

																									var ifres8438 Obj

																									if True == tmp8451 {
																										ifres8438 = True

																									} else {
																										ifres8438 = False

																									}

																									ifres8437 = ifres8438

																								} else {
																									ifres8437 = False

																								}

																								var ifres8436 Obj

																								if True == ifres8437 {
																									ifres8436 = True

																								} else {
																									ifres8436 = False

																								}

																								ifres8435 = ifres8436

																							} else {
																								ifres8435 = False

																							}

																							var ifres8434 Obj

																							if True == ifres8435 {
																								ifres8434 = True

																							} else {
																								ifres8434 = False

																							}

																							ifres8433 = ifres8434

																						} else {
																							ifres8433 = False

																						}

																						var ifres8432 Obj

																						if True == ifres8433 {
																							ifres8432 = True

																						} else {
																							ifres8432 = False

																						}

																						ifres8431 = ifres8432

																					} else {
																						ifres8431 = False

																					}

																					var ifres8430 Obj

																					if True == ifres8431 {
																						ifres8430 = True

																					} else {
																						ifres8430 = False

																					}

																					ifres8429 = ifres8430

																				} else {
																					ifres8429 = False

																				}

																				var ifres8428 Obj

																				if True == ifres8429 {
																					ifres8428 = True

																				} else {
																					ifres8428 = False

																				}

																				ifres8427 = ifres8428

																			} else {
																				ifres8427 = False

																			}

																			var ifres8426 Obj

																			if True == ifres8427 {
																				ifres8426 = True

																			} else {
																				ifres8426 = False

																			}

																			ifres8425 = ifres8426

																		} else {
																			ifres8425 = False

																		}

																		var ifres8424 Obj

																		if True == ifres8425 {
																			ifres8424 = True

																		} else {
																			ifres8424 = False

																		}

																		ifres8423 = ifres8424

																	} else {
																		ifres8423 = False

																	}

																	var ifres8422 Obj

																	if True == ifres8423 {
																		ifres8422 = True

																	} else {
																		ifres8422 = False

																	}

																	ifres8421 = ifres8422

																} else {
																	ifres8421 = False

																}

																var ifres8420 Obj

																if True == ifres8421 {
																	ifres8420 = True

																} else {
																	ifres8420 = False

																}

																ifres8419 = ifres8420

															} else {
																ifres8419 = False

															}

															var ifres8418 Obj

															if True == ifres8419 {
																ifres8418 = True

															} else {
																ifres8418 = False

															}

															ifres8417 = ifres8418

														} else {
															ifres8417 = False

														}

														var ifres8416 Obj

														if True == ifres8417 {
															ifres8416 = True

														} else {
															ifres8416 = False

														}

														ifres8415 = ifres8416

													} else {
														ifres8415 = False

													}

													var ifres8414 Obj

													if True == ifres8415 {
														ifres8414 = True

													} else {
														ifres8414 = False

													}

													ifres8413 = ifres8414

												} else {
													ifres8413 = False

												}

												var ifres8412 Obj

												if True == ifres8413 {
													ifres8412 = True

												} else {
													ifres8412 = False

												}

												ifres8411 = ifres8412

											} else {
												ifres8411 = False

											}

											var ifres8410 Obj

											if True == ifres8411 {
												ifres8410 = True

											} else {
												ifres8410 = False

											}

											ifres8409 = ifres8410

										} else {
											ifres8409 = False

										}

										var ifres8408 Obj

										if True == ifres8409 {
											ifres8408 = True

										} else {
											ifres8408 = False

										}

										ifres8407 = ifres8408

									} else {
										ifres8407 = False

									}

									var ifres8406 Obj

									if True == ifres8407 {
										ifres8406 = True

									} else {
										ifres8406 = False

									}

									ifres8405 = ifres8406

								} else {
									ifres8405 = False

								}

								var ifres8404 Obj

								if True == ifres8405 {
									ifres8404 = True

								} else {
									ifres8404 = False

								}

								ifres8403 = ifres8404

							} else {
								ifres8403 = False

							}

							var ifres8402 Obj

							if True == ifres8403 {
								ifres8402 = True

							} else {
								ifres8402 = False

							}

							ifres8401 = ifres8402

						} else {
							ifres8401 = False

						}

						var ifres8400 Obj

						if True == ifres8401 {
							ifres8400 = True

						} else {
							ifres8400 = False

						}

						ifres8399 = ifres8400

					} else {
						ifres8399 = False

					}

					var ifres8398 Obj

					if True == ifres8399 {
						ifres8398 = True

					} else {
						ifres8398 = False

					}

					ifres8397 = ifres8398

				} else {
					ifres8397 = False

				}

				var ifres8396 Obj

				if True == ifres8397 {
					ifres8396 = True

				} else {
					ifres8396 = False

				}

				ifres8395 = ifres8396

			} else {
				ifres8395 = False

			}

			var ifres8394 Obj

			if True == ifres8395 {
				ifres8394 = True

			} else {
				ifres8394 = False

			}

			ifres8393 = ifres8394

		} else {
			ifres8393 = False

		}

		if True == ifres8393 {
			tmp8382 := PrimTail(V1737)

			tmp8383 := PrimTail(tmp8382)

			tmp8384 := PrimHead(tmp8383)

			tmp8385 := PrimTail(tmp8384)

			tmp8386 := PrimHead(tmp8385)

			tmp8387 := PrimTail(V1737)

			tmp8388 := PrimTail(tmp8387)

			tmp8389 := PrimHead(tmp8388)

			tmp8390 := PrimTail(tmp8389)

			tmp8391 := PrimTail(tmp8390)

			tmp8392 := PrimCons(sym_a_a_6, tmp8391)

			__e.Return(PrimCons(tmp8386, tmp8392))
			return

		} else {
			__e.Return(V1737)
			return
		}

	}, 1)

	tmp8557 := Call(__e, PrimNS2Value(symdef), symshen_4pretty_1type, tmp8380)

	_ = tmp8557

	tmp8558 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("type error\n")))
		return
	}, 0)

	tmp8559 := Call(__e, PrimNS2Value(symdef), symshen_4type_1error, tmp8558)

	_ = tmp8559

	tmp8560 := MakeNative(func(__e *ControlFlow) {
		V1738 := __e.Get(1)
		_ = V1738
		tmp8561 := MakeNative(func(__e *ControlFlow) {
			KLFile := __e.Get(1)
			_ = KLFile
			tmp8562 := MakeNative(func(__e *ControlFlow) {
				Code := __e.Get(1)
				_ = Code
				tmp8563 := MakeNative(func(__e *ControlFlow) {
					Open := __e.Get(1)
					_ = Open
					tmp8564 := MakeNative(func(__e *ControlFlow) {
						KL := __e.Get(1)
						_ = KL
						tmp8565 := MakeNative(func(__e *ControlFlow) {
							Write := __e.Get(1)
							_ = Write
							__e.Return(KLFile)
							return
						}, 1)

						tmp8566 := Call(__e, PrimNS2Value(symshen_4write_1kl), KL, Open)

						__e.TailApply(tmp8565, tmp8566)
						return

					}, 1)

					tmp8567 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4shen_1_6kl_1h), X)
						return
					}, 1)

					tmp8568 := Call(__e, PrimNS2Value(symmap), tmp8567, Code)

					__e.TailApply(tmp8564, tmp8568)
					return

				}, 1)

				tmp8569 := PrimOpenStream(KLFile, symout)

				__e.TailApply(tmp8563, tmp8569)
				return

			}, 1)

			tmp8570 := Call(__e, PrimNS2Value(symread_1file), V1738)

			__e.TailApply(tmp8562, tmp8570)
			return

		}, 1)

		tmp8571 := Call(__e, PrimNS2Value(symshen_4klfile), V1738)

		__e.TailApply(tmp8561, tmp8571)
		return

	}, 1)

	tmp8572 := Call(__e, PrimNS2Value(symdef), symbootstrap, tmp8560)

	_ = tmp8572

	tmp8573 := MakeNative(func(__e *ControlFlow) {
		V1741 := __e.Get(1)
		_ = V1741
		V1742 := __e.Get(2)
		_ = V1742
		tmp8587 := PrimEqual(Nil, V1741)

		if True == tmp8587 {
			__e.Return(PrimCloseStream(V1742))
			return
		} else {
			tmp8586 := PrimIsPair(V1741)

			var ifres8582 Obj

			if True == tmp8586 {
				tmp8584 := PrimHead(V1741)

				tmp8585 := PrimIsPair(tmp8584)

				var ifres8583 Obj

				if True == tmp8585 {
					ifres8583 = True

				} else {
					ifres8583 = False

				}

				ifres8582 = ifres8583

			} else {
				ifres8582 = False

			}

			if True == ifres8582 {
				tmp8579 := PrimTail(V1741)

				tmp8580 := PrimHead(V1741)

				tmp8581 := Call(__e, PrimNS2Value(symshen_4write_1kl_1h), tmp8580, V1742)

				_ = tmp8581

				__e.TailApply(PrimNS2Value(symshen_4write_1kl), tmp8579, V1742)
				return

			} else {
				tmp8578 := PrimIsPair(V1741)

				if True == tmp8578 {
					tmp8577 := PrimTail(V1741)

					__e.TailApply(PrimNS2Value(symshen_4write_1kl), tmp8577, V1742)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4write_1kl)
					return
				}

			}

		}

	}, 2)

	tmp8588 := Call(__e, PrimNS2Value(symdef), symshen_4write_1kl, tmp8573)

	_ = tmp8588

	tmp8589 := MakeNative(func(__e *ControlFlow) {
		V1745 := __e.Get(1)
		_ = V1745
		V1746 := __e.Get(2)
		_ = V1746
		tmp8629 := PrimIsPair(V1745)

		var ifres8592 Obj

		if True == tmp8629 {
			tmp8627 := PrimHead(V1745)

			tmp8628 := PrimEqual(symdefun, tmp8627)

			var ifres8594 Obj

			if True == tmp8628 {
				tmp8625 := PrimTail(V1745)

				tmp8626 := PrimIsPair(tmp8625)

				var ifres8596 Obj

				if True == tmp8626 {
					tmp8622 := PrimTail(V1745)

					tmp8623 := PrimHead(tmp8622)

					tmp8624 := PrimEqual(symfail, tmp8623)

					var ifres8598 Obj

					if True == tmp8624 {
						tmp8619 := PrimTail(V1745)

						tmp8620 := PrimTail(tmp8619)

						tmp8621 := PrimIsPair(tmp8620)

						var ifres8600 Obj

						if True == tmp8621 {
							tmp8615 := PrimTail(V1745)

							tmp8616 := PrimTail(tmp8615)

							tmp8617 := PrimHead(tmp8616)

							tmp8618 := PrimEqual(Nil, tmp8617)

							var ifres8602 Obj

							if True == tmp8618 {
								tmp8611 := PrimTail(V1745)

								tmp8612 := PrimTail(tmp8611)

								tmp8613 := PrimTail(tmp8612)

								tmp8614 := PrimIsPair(tmp8613)

								var ifres8604 Obj

								if True == tmp8614 {
									tmp8606 := PrimTail(V1745)

									tmp8607 := PrimTail(tmp8606)

									tmp8608 := PrimTail(tmp8607)

									tmp8609 := PrimTail(tmp8608)

									tmp8610 := PrimEqual(Nil, tmp8609)

									var ifres8605 Obj

									if True == tmp8610 {
										ifres8605 = True

									} else {
										ifres8605 = False

									}

									ifres8604 = ifres8605

								} else {
									ifres8604 = False

								}

								var ifres8603 Obj

								if True == ifres8604 {
									ifres8603 = True

								} else {
									ifres8603 = False

								}

								ifres8602 = ifres8603

							} else {
								ifres8602 = False

							}

							var ifres8601 Obj

							if True == ifres8602 {
								ifres8601 = True

							} else {
								ifres8601 = False

							}

							ifres8600 = ifres8601

						} else {
							ifres8600 = False

						}

						var ifres8599 Obj

						if True == ifres8600 {
							ifres8599 = True

						} else {
							ifres8599 = False

						}

						ifres8598 = ifres8599

					} else {
						ifres8598 = False

					}

					var ifres8597 Obj

					if True == ifres8598 {
						ifres8597 = True

					} else {
						ifres8597 = False

					}

					ifres8596 = ifres8597

				} else {
					ifres8596 = False

				}

				var ifres8595 Obj

				if True == ifres8596 {
					ifres8595 = True

				} else {
					ifres8595 = False

				}

				ifres8594 = ifres8595

			} else {
				ifres8594 = False

			}

			var ifres8593 Obj

			if True == ifres8594 {
				ifres8593 = True

			} else {
				ifres8593 = False

			}

			ifres8592 = ifres8593

		} else {
			ifres8592 = False

		}

		if True == ifres8592 {
			__e.TailApply(PrimNS2Value(sympr), MakeString("(defun fail () shen.fail!)"), V1746)
			return
		} else {
			tmp8591 := Call(__e, PrimNS2Value(symshen_4app), V1745, MakeString("\n\n"), symshen_4r)

			__e.TailApply(PrimNS2Value(sympr), tmp8591, V1746)
			return

		}

	}, 2)

	tmp8630 := Call(__e, PrimNS2Value(symdef), symshen_4write_1kl_1h, tmp8589)

	_ = tmp8630

	tmp8631 := MakeNative(func(__e *ControlFlow) {
		V1747 := __e.Get(1)
		_ = V1747
		tmp8640 := PrimEqual(MakeString(""), V1747)

		if True == tmp8640 {
			__e.Return(MakeString(".kl"))
			return
		} else {
			tmp8639 := PrimEqual(MakeString(".shen"), V1747)

			if True == tmp8639 {
				__e.Return(MakeString(".kl"))
				return
			} else {
				tmp8638 := Call(__e, PrimNS2Value(symshen_4_7string_2), V1747)

				if True == tmp8638 {
					tmp8635 := Call(__e, PrimNS2Value(symhdstr), V1747)

					tmp8636 := PrimTailString(V1747)

					tmp8637 := Call(__e, PrimNS2Value(symshen_4klfile), tmp8636)

					__e.TailApply(PrimNS2Value(sym_8s), tmp8635, tmp8637)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4klfile)
					return
				}

			}

		}

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), symshen_4klfile, tmp8631)
	return

}, 0)
