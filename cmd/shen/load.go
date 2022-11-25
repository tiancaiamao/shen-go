package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
	tmp9167 := MakeNative(func(__e *ControlFlow) {
		V1704 := __e.Get(1)
		_ = V1704
		tmp9168 := MakeNative(func(__e *ControlFlow) {
			TC_2 := __e.Get(1)
			_ = TC_2
			tmp9169 := MakeNative(func(__e *ControlFlow) {
				Load := __e.Get(1)
				_ = Load
				tmp9170 := MakeNative(func(__e *ControlFlow) {
					Infs := __e.Get(1)
					_ = Infs
					__e.Return(symloaded)
					return
				}, 1)

				var ifres9171 Obj

				if True == TC_2 {
					tmp9172 := Call(__e, PrimFunc(syminferences))

					tmp9173 := Call(__e, PrimFunc(symshen_4app), tmp9172, MakeString(" inferences\n"), symshen_4a)

					tmp9174 := PrimStringConcat(MakeString("\ntypechecked in "), tmp9173)

					tmp9175 := Call(__e, PrimFunc(symstoutput))

					tmp9176 := Call(__e, PrimFunc(sympr), tmp9174, tmp9175)

					ifres9171 = tmp9176

				} else {
					ifres9171 = symshen_4skip

				}

				__e.TailApply(tmp9170, ifres9171)
				return

			}, 1)

			tmp9177 := MakeNative(func(__e *ControlFlow) {
				Start := __e.Get(1)
				_ = Start
				tmp9178 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp9179 := MakeNative(func(__e *ControlFlow) {
						Finish := __e.Get(1)
						_ = Finish
						tmp9180 := MakeNative(func(__e *ControlFlow) {
							Time := __e.Get(1)
							_ = Time
							tmp9181 := MakeNative(func(__e *ControlFlow) {
								Message := __e.Get(1)
								_ = Message
								__e.Return(Result)
								return
							}, 1)

							tmp9182 := PrimStr(Time)

							tmp9183 := PrimStringConcat(tmp9182, MakeString(" secs\n"))

							tmp9184 := PrimStringConcat(MakeString("\nrun time: "), tmp9183)

							tmp9185 := Call(__e, PrimFunc(symstoutput))

							tmp9186 := Call(__e, PrimFunc(sympr), tmp9184, tmp9185)

							__e.TailApply(tmp9181, tmp9186)
							return

						}, 1)

						tmp9187 := PrimNumberSubtract(Finish, Start)

						__e.TailApply(tmp9180, tmp9187)
						return

					}, 1)

					tmp9188 := PrimGetTime(symrun)

					__e.TailApply(tmp9179, tmp9188)
					return

				}, 1)

				tmp9189 := Call(__e, PrimFunc(symread_1file), V1704)

				tmp9190 := Call(__e, PrimFunc(symshen_4load_1help), TC_2, tmp9189)

				__e.TailApply(tmp9178, tmp9190)
				return

			}, 1)

			tmp9191 := PrimGetTime(symrun)

			tmp9192 := Call(__e, tmp9177, tmp9191)

			__e.TailApply(tmp9169, tmp9192)
			return

		}, 1)

		tmp9193 := PrimValue(symshen_4_dtc_d)

		__e.TailApply(tmp9168, tmp9193)
		return

	}, 1)

	tmp9194 := Call(__e, ns2_1set, symload, tmp9167)

	_ = tmp9194

	tmp9195 := MakeNative(func(__e *ControlFlow) {
		V1707 := __e.Get(1)
		_ = V1707
		V1708 := __e.Get(2)
		_ = V1708
		tmp9197 := PrimEqual(False, V1707)

		if True == tmp9197 {
			__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V1708)
			return
		} else {
			__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V1708)
			return
		}

	}, 2)

	tmp9198 := Call(__e, ns2_1set, symshen_4load_1help, tmp9195)

	_ = tmp9198

	tmp9199 := MakeNative(func(__e *ControlFlow) {
		V1709 := __e.Get(1)
		_ = V1709
		tmp9200 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			tmp9201 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Y)

			tmp9202 := Call(__e, PrimFunc(symeval_1kl), tmp9201)

			tmp9203 := Call(__e, PrimFunc(symshen_4app), tmp9202, MakeString("\n"), symshen_4s)

			tmp9204 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp9203, tmp9204)
			return

		}, 1)

		__e.TailApply(PrimFunc(symmap), tmp9200, V1709)
		return

	}, 1)

	tmp9205 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9199)

	_ = tmp9205

	tmp9206 := MakeNative(func(__e *ControlFlow) {
		V1710 := __e.Get(1)
		_ = V1710
		tmp9207 := MakeNative(func(__e *ControlFlow) {
			Table := __e.Get(1)
			_ = Table
			tmp9208 := MakeNative(func(__e *ControlFlow) {
				Assume := __e.Get(1)
				_ = Assume
				tmp9209 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimFunc(symshen_4work_1through), V1710)
					return
				}, 0)

				tmp9210 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.TailApply(PrimFunc(symshen_4unwind_1types), E, Table)
					return
				}, 1)

				__e.TailApply(try_1catch, tmp9209, tmp9210)
				return

			}, 1)

			tmp9211 := Call(__e, PrimFunc(symshen_4assumetypes), Table)

			__e.TailApply(tmp9208, tmp9211)
			return

		}, 1)

		tmp9212 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			__e.TailApply(PrimFunc(symshen_4typetable), Y)
			return
		}, 1)

		tmp9213 := Call(__e, PrimFunc(symmapcan), tmp9212, V1710)

		__e.TailApply(tmp9207, tmp9213)
		return

	}, 1)

	tmp9214 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9206)

	_ = tmp9214

	tmp9215 := MakeNative(func(__e *ControlFlow) {
		V1715 := __e.Get(1)
		_ = V1715
		tmp9260 := PrimIsPair(V1715)

		var ifres9241 Obj

		if True == tmp9260 {
			tmp9258 := PrimHead(V1715)

			tmp9259 := PrimEqual(symdefine, tmp9258)

			var ifres9243 Obj

			if True == tmp9259 {
				tmp9256 := PrimTail(V1715)

				tmp9257 := PrimIsPair(tmp9256)

				var ifres9245 Obj

				if True == tmp9257 {
					tmp9253 := PrimTail(V1715)

					tmp9254 := PrimTail(tmp9253)

					tmp9255 := PrimIsPair(tmp9254)

					var ifres9247 Obj

					if True == tmp9255 {
						tmp9249 := PrimTail(V1715)

						tmp9250 := PrimTail(tmp9249)

						tmp9251 := PrimHead(tmp9250)

						tmp9252 := PrimEqual(sym_i, tmp9251)

						var ifres9248 Obj

						if True == tmp9252 {
							ifres9248 = True

						} else {
							ifres9248 = False

						}

						ifres9247 = ifres9248

					} else {
						ifres9247 = False

					}

					var ifres9246 Obj

					if True == ifres9247 {
						ifres9246 = True

					} else {
						ifres9246 = False

					}

					ifres9245 = ifres9246

				} else {
					ifres9245 = False

				}

				var ifres9244 Obj

				if True == ifres9245 {
					ifres9244 = True

				} else {
					ifres9244 = False

				}

				ifres9243 = ifres9244

			} else {
				ifres9243 = False

			}

			var ifres9242 Obj

			if True == ifres9243 {
				ifres9242 = True

			} else {
				ifres9242 = False

			}

			ifres9241 = ifres9242

		} else {
			ifres9241 = False

		}

		if True == ifres9241 {
			tmp9216 := PrimTail(V1715)

			tmp9217 := PrimHead(tmp9216)

			tmp9218 := PrimTail(V1715)

			tmp9219 := PrimHead(tmp9218)

			tmp9220 := PrimTail(V1715)

			tmp9221 := PrimTail(tmp9220)

			tmp9222 := PrimTail(tmp9221)

			tmp9223 := Call(__e, PrimFunc(symshen_4type_1F), tmp9219, tmp9222)

			tmp9224 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9223)

			tmp9225 := PrimCons(tmp9224, Nil)

			__e.Return(PrimCons(tmp9217, tmp9225))
			return

		} else {
			tmp9239 := PrimIsPair(V1715)

			var ifres9231 Obj

			if True == tmp9239 {
				tmp9237 := PrimHead(V1715)

				tmp9238 := PrimEqual(symdefine, tmp9237)

				var ifres9233 Obj

				if True == tmp9238 {
					tmp9235 := PrimTail(V1715)

					tmp9236 := PrimIsPair(tmp9235)

					var ifres9234 Obj

					if True == tmp9236 {
						ifres9234 = True

					} else {
						ifres9234 = False

					}

					ifres9233 = ifres9234

				} else {
					ifres9233 = False

				}

				var ifres9232 Obj

				if True == ifres9233 {
					ifres9232 = True

				} else {
					ifres9232 = False

				}

				ifres9231 = ifres9232

			} else {
				ifres9231 = False

			}

			if True == ifres9231 {
				tmp9226 := PrimTail(V1715)

				tmp9227 := PrimHead(tmp9226)

				tmp9228 := Call(__e, PrimFunc(symshen_4app), tmp9227, MakeString("\n"), symshen_4a)

				tmp9229 := PrimStringConcat(MakeString("missing { in "), tmp9228)

				__e.Return(PrimSimpleError(tmp9229))
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp9261 := Call(__e, ns2_1set, symshen_4typetable, tmp9215)

	_ = tmp9261

	tmp9262 := MakeNative(func(__e *ControlFlow) {
		V1722 := __e.Get(1)
		_ = V1722
		V1723 := __e.Get(2)
		_ = V1723
		tmp9275 := PrimIsPair(V1723)

		var ifres9271 Obj

		if True == tmp9275 {
			tmp9273 := PrimHead(V1723)

			tmp9274 := PrimEqual(sym_j, tmp9273)

			var ifres9272 Obj

			if True == tmp9274 {
				ifres9272 = True

			} else {
				ifres9272 = False

			}

			ifres9271 = ifres9272

		} else {
			ifres9271 = False

		}

		if True == ifres9271 {
			__e.Return(Nil)
			return
		} else {
			tmp9269 := PrimIsPair(V1723)

			if True == tmp9269 {
				tmp9263 := PrimHead(V1723)

				tmp9264 := PrimTail(V1723)

				tmp9265 := Call(__e, PrimFunc(symshen_4type_1F), V1722, tmp9264)

				__e.Return(PrimCons(tmp9263, tmp9265))
				return

			} else {
				tmp9266 := Call(__e, PrimFunc(symshen_4app), V1722, MakeString("\n"), symshen_4a)

				tmp9267 := PrimStringConcat(MakeString("missing } in "), tmp9266)

				__e.Return(PrimSimpleError(tmp9267))
				return

			}

		}

	}, 2)

	tmp9276 := Call(__e, ns2_1set, symshen_4type_1F, tmp9262)

	_ = tmp9276

	tmp9277 := MakeNative(func(__e *ControlFlow) {
		V1726 := __e.Get(1)
		_ = V1726
		tmp9291 := PrimEqual(Nil, V1726)

		if True == tmp9291 {
			__e.Return(Nil)
			return
		} else {
			tmp9289 := PrimIsPair(V1726)

			var ifres9285 Obj

			if True == tmp9289 {
				tmp9287 := PrimTail(V1726)

				tmp9288 := PrimIsPair(tmp9287)

				var ifres9286 Obj

				if True == tmp9288 {
					ifres9286 = True

				} else {
					ifres9286 = False

				}

				ifres9285 = ifres9286

			} else {
				ifres9285 = False

			}

			if True == ifres9285 {
				tmp9278 := PrimHead(V1726)

				tmp9279 := PrimTail(V1726)

				tmp9280 := PrimHead(tmp9279)

				tmp9281 := Call(__e, PrimFunc(symdeclare), tmp9278, tmp9280)

				_ = tmp9281

				tmp9282 := PrimTail(V1726)

				tmp9283 := PrimTail(tmp9282)

				__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9283)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
				return
			}

		}

	}, 1)

	tmp9292 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9277)

	_ = tmp9292

	tmp9293 := MakeNative(func(__e *ControlFlow) {
		V1731 := __e.Get(1)
		_ = V1731
		V1732 := __e.Get(2)
		_ = V1732
		tmp9304 := PrimIsPair(V1732)

		var ifres9300 Obj

		if True == tmp9304 {
			tmp9302 := PrimHead(V1732)

			tmp9303 := PrimIsPair(tmp9302)

			var ifres9301 Obj

			if True == tmp9303 {
				ifres9301 = True

			} else {
				ifres9301 = False

			}

			ifres9300 = ifres9301

		} else {
			ifres9300 = False

		}

		if True == ifres9300 {
			tmp9294 := PrimHead(V1732)

			tmp9295 := PrimHead(tmp9294)

			tmp9296 := Call(__e, PrimFunc(symdestroy), tmp9295)

			_ = tmp9296

			tmp9297 := PrimTail(V1732)

			__e.TailApply(PrimFunc(symshen_4unwind_1types), V1731, tmp9297)
			return

		} else {
			tmp9298 := PrimErrorToString(V1731)

			__e.Return(PrimSimpleError(tmp9298))
			return

		}

	}, 2)

	tmp9305 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9293)

	_ = tmp9305

	tmp9306 := MakeNative(func(__e *ControlFlow) {
		V1735 := __e.Get(1)
		_ = V1735
		tmp9356 := PrimEqual(Nil, V1735)

		if True == tmp9356 {
			__e.Return(Nil)
			return
		} else {
			tmp9354 := PrimIsPair(V1735)

			var ifres9339 Obj

			if True == tmp9354 {
				tmp9352 := PrimTail(V1735)

				tmp9353 := PrimIsPair(tmp9352)

				var ifres9341 Obj

				if True == tmp9353 {
					tmp9349 := PrimTail(V1735)

					tmp9350 := PrimTail(tmp9349)

					tmp9351 := PrimIsPair(tmp9350)

					var ifres9343 Obj

					if True == tmp9351 {
						tmp9345 := PrimTail(V1735)

						tmp9346 := PrimHead(tmp9345)

						tmp9347 := PrimIntern(MakeString(":"))

						tmp9348 := PrimEqual(tmp9346, tmp9347)

						var ifres9344 Obj

						if True == tmp9348 {
							ifres9344 = True

						} else {
							ifres9344 = False

						}

						ifres9343 = ifres9344

					} else {
						ifres9343 = False

					}

					var ifres9342 Obj

					if True == ifres9343 {
						ifres9342 = True

					} else {
						ifres9342 = False

					}

					ifres9341 = ifres9342

				} else {
					ifres9341 = False

				}

				var ifres9340 Obj

				if True == ifres9341 {
					ifres9340 = True

				} else {
					ifres9340 = False

				}

				ifres9339 = ifres9340

			} else {
				ifres9339 = False

			}

			if True == ifres9339 {
				tmp9307 := MakeNative(func(__e *ControlFlow) {
					Check := __e.Get(1)
					_ = Check
					tmp9323 := PrimEqual(Check, False)

					if True == tmp9323 {
						__e.TailApply(PrimFunc(symshen_4type_1error))
						return
					} else {
						tmp9308 := MakeNative(func(__e *ControlFlow) {
							Eval := __e.Get(1)
							_ = Eval
							tmp9309 := MakeNative(func(__e *ControlFlow) {
								Message := __e.Get(1)
								_ = Message
								tmp9310 := PrimTail(V1735)

								tmp9311 := PrimTail(tmp9310)

								tmp9312 := PrimTail(tmp9311)

								__e.TailApply(PrimFunc(symshen_4work_1through), tmp9312)
								return

							}, 1)

							tmp9313 := Call(__e, PrimFunc(symshen_4pretty_1type), Check)

							tmp9314 := Call(__e, PrimFunc(symshen_4app), tmp9313, MakeString("\n"), symshen_4r)

							tmp9315 := PrimStringConcat(MakeString(" : "), tmp9314)

							tmp9316 := Call(__e, PrimFunc(symshen_4app), Eval, tmp9315, symshen_4s)

							tmp9317 := Call(__e, PrimFunc(symstoutput))

							tmp9318 := Call(__e, PrimFunc(sympr), tmp9316, tmp9317)

							__e.TailApply(tmp9309, tmp9318)
							return

						}, 1)

						tmp9319 := PrimHead(V1735)

						tmp9320 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9319)

						tmp9321 := Call(__e, PrimFunc(symeval_1kl), tmp9320)

						__e.TailApply(tmp9308, tmp9321)
						return

					}

				}, 1)

				tmp9324 := PrimHead(V1735)

				tmp9325 := PrimTail(V1735)

				tmp9326 := PrimTail(tmp9325)

				tmp9327 := PrimHead(tmp9326)

				tmp9328 := Call(__e, PrimFunc(symshen_4typecheck), tmp9324, tmp9327)

				__e.TailApply(tmp9307, tmp9328)
				return

			} else {
				tmp9337 := PrimIsPair(V1735)

				if True == tmp9337 {
					tmp9329 := PrimHead(V1735)

					tmp9330 := PrimIntern(MakeString(":"))

					tmp9331 := Call(__e, PrimFunc(symprotect), symA)

					tmp9332 := PrimTail(V1735)

					tmp9333 := PrimCons(tmp9331, tmp9332)

					tmp9334 := PrimCons(tmp9330, tmp9333)

					tmp9335 := PrimCons(tmp9329, tmp9334)

					__e.TailApply(PrimFunc(symshen_4work_1through), tmp9335)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
					return
				}

			}

		}

	}, 1)

	tmp9357 := Call(__e, ns2_1set, symshen_4work_1through, tmp9306)

	_ = tmp9357

	tmp9358 := MakeNative(func(__e *ControlFlow) {
		V1737 := __e.Get(1)
		_ = V1737
		tmp9534 := PrimIsPair(V1737)

		var ifres9371 Obj

		if True == tmp9534 {
			tmp9532 := PrimHead(V1737)

			tmp9533 := PrimIsPair(tmp9532)

			var ifres9373 Obj

			if True == tmp9533 {
				tmp9529 := PrimHead(V1737)

				tmp9530 := PrimHead(tmp9529)

				tmp9531 := PrimEqual(symstr, tmp9530)

				var ifres9375 Obj

				if True == tmp9531 {
					tmp9526 := PrimHead(V1737)

					tmp9527 := PrimTail(tmp9526)

					tmp9528 := PrimIsPair(tmp9527)

					var ifres9377 Obj

					if True == tmp9528 {
						tmp9522 := PrimHead(V1737)

						tmp9523 := PrimTail(tmp9522)

						tmp9524 := PrimHead(tmp9523)

						tmp9525 := PrimIsPair(tmp9524)

						var ifres9379 Obj

						if True == tmp9525 {
							tmp9517 := PrimHead(V1737)

							tmp9518 := PrimTail(tmp9517)

							tmp9519 := PrimHead(tmp9518)

							tmp9520 := PrimHead(tmp9519)

							tmp9521 := PrimEqual(symlist, tmp9520)

							var ifres9381 Obj

							if True == tmp9521 {
								tmp9512 := PrimHead(V1737)

								tmp9513 := PrimTail(tmp9512)

								tmp9514 := PrimHead(tmp9513)

								tmp9515 := PrimTail(tmp9514)

								tmp9516 := PrimIsPair(tmp9515)

								var ifres9383 Obj

								if True == tmp9516 {
									tmp9506 := PrimHead(V1737)

									tmp9507 := PrimTail(tmp9506)

									tmp9508 := PrimHead(tmp9507)

									tmp9509 := PrimTail(tmp9508)

									tmp9510 := PrimTail(tmp9509)

									tmp9511 := PrimEqual(Nil, tmp9510)

									var ifres9385 Obj

									if True == tmp9511 {
										tmp9502 := PrimHead(V1737)

										tmp9503 := PrimTail(tmp9502)

										tmp9504 := PrimTail(tmp9503)

										tmp9505 := PrimIsPair(tmp9504)

										var ifres9387 Obj

										if True == tmp9505 {
											tmp9497 := PrimHead(V1737)

											tmp9498 := PrimTail(tmp9497)

											tmp9499 := PrimTail(tmp9498)

											tmp9500 := PrimTail(tmp9499)

											tmp9501 := PrimEqual(Nil, tmp9500)

											var ifres9389 Obj

											if True == tmp9501 {
												tmp9495 := PrimTail(V1737)

												tmp9496 := PrimIsPair(tmp9495)

												var ifres9391 Obj

												if True == tmp9496 {
													tmp9492 := PrimTail(V1737)

													tmp9493 := PrimHead(tmp9492)

													tmp9494 := PrimEqual(sym_1_1_6, tmp9493)

													var ifres9393 Obj

													if True == tmp9494 {
														tmp9489 := PrimTail(V1737)

														tmp9490 := PrimTail(tmp9489)

														tmp9491 := PrimIsPair(tmp9490)

														var ifres9395 Obj

														if True == tmp9491 {
															tmp9485 := PrimTail(V1737)

															tmp9486 := PrimTail(tmp9485)

															tmp9487 := PrimHead(tmp9486)

															tmp9488 := PrimIsPair(tmp9487)

															var ifres9397 Obj

															if True == tmp9488 {
																tmp9480 := PrimTail(V1737)

																tmp9481 := PrimTail(tmp9480)

																tmp9482 := PrimHead(tmp9481)

																tmp9483 := PrimHead(tmp9482)

																tmp9484 := PrimEqual(symstr, tmp9483)

																var ifres9399 Obj

																if True == tmp9484 {
																	tmp9475 := PrimTail(V1737)

																	tmp9476 := PrimTail(tmp9475)

																	tmp9477 := PrimHead(tmp9476)

																	tmp9478 := PrimTail(tmp9477)

																	tmp9479 := PrimIsPair(tmp9478)

																	var ifres9401 Obj

																	if True == tmp9479 {
																		tmp9469 := PrimTail(V1737)

																		tmp9470 := PrimTail(tmp9469)

																		tmp9471 := PrimHead(tmp9470)

																		tmp9472 := PrimTail(tmp9471)

																		tmp9473 := PrimHead(tmp9472)

																		tmp9474 := PrimIsPair(tmp9473)

																		var ifres9403 Obj

																		if True == tmp9474 {
																			tmp9462 := PrimTail(V1737)

																			tmp9463 := PrimTail(tmp9462)

																			tmp9464 := PrimHead(tmp9463)

																			tmp9465 := PrimTail(tmp9464)

																			tmp9466 := PrimHead(tmp9465)

																			tmp9467 := PrimHead(tmp9466)

																			tmp9468 := PrimEqual(symlist, tmp9467)

																			var ifres9405 Obj

																			if True == tmp9468 {
																				tmp9455 := PrimTail(V1737)

																				tmp9456 := PrimTail(tmp9455)

																				tmp9457 := PrimHead(tmp9456)

																				tmp9458 := PrimTail(tmp9457)

																				tmp9459 := PrimHead(tmp9458)

																				tmp9460 := PrimTail(tmp9459)

																				tmp9461 := PrimIsPair(tmp9460)

																				var ifres9407 Obj

																				if True == tmp9461 {
																					tmp9447 := PrimTail(V1737)

																					tmp9448 := PrimTail(tmp9447)

																					tmp9449 := PrimHead(tmp9448)

																					tmp9450 := PrimTail(tmp9449)

																					tmp9451 := PrimHead(tmp9450)

																					tmp9452 := PrimTail(tmp9451)

																					tmp9453 := PrimTail(tmp9452)

																					tmp9454 := PrimEqual(Nil, tmp9453)

																					var ifres9409 Obj

																					if True == tmp9454 {
																						tmp9441 := PrimTail(V1737)

																						tmp9442 := PrimTail(tmp9441)

																						tmp9443 := PrimHead(tmp9442)

																						tmp9444 := PrimTail(tmp9443)

																						tmp9445 := PrimTail(tmp9444)

																						tmp9446 := PrimIsPair(tmp9445)

																						var ifres9411 Obj

																						if True == tmp9446 {
																							tmp9434 := PrimTail(V1737)

																							tmp9435 := PrimTail(tmp9434)

																							tmp9436 := PrimHead(tmp9435)

																							tmp9437 := PrimTail(tmp9436)

																							tmp9438 := PrimTail(tmp9437)

																							tmp9439 := PrimTail(tmp9438)

																							tmp9440 := PrimEqual(Nil, tmp9439)

																							var ifres9413 Obj

																							if True == tmp9440 {
																								tmp9430 := PrimTail(V1737)

																								tmp9431 := PrimTail(tmp9430)

																								tmp9432 := PrimTail(tmp9431)

																								tmp9433 := PrimEqual(Nil, tmp9432)

																								var ifres9415 Obj

																								if True == tmp9433 {
																									tmp9417 := PrimHead(V1737)

																									tmp9418 := PrimTail(tmp9417)

																									tmp9419 := PrimHead(tmp9418)

																									tmp9420 := PrimTail(tmp9419)

																									tmp9421 := PrimHead(tmp9420)

																									tmp9422 := PrimTail(V1737)

																									tmp9423 := PrimTail(tmp9422)

																									tmp9424 := PrimHead(tmp9423)

																									tmp9425 := PrimTail(tmp9424)

																									tmp9426 := PrimHead(tmp9425)

																									tmp9427 := PrimTail(tmp9426)

																									tmp9428 := PrimHead(tmp9427)

																									tmp9429 := PrimEqual(tmp9421, tmp9428)

																									var ifres9416 Obj

																									if True == tmp9429 {
																										ifres9416 = True

																									} else {
																										ifres9416 = False

																									}

																									ifres9415 = ifres9416

																								} else {
																									ifres9415 = False

																								}

																								var ifres9414 Obj

																								if True == ifres9415 {
																									ifres9414 = True

																								} else {
																									ifres9414 = False

																								}

																								ifres9413 = ifres9414

																							} else {
																								ifres9413 = False

																							}

																							var ifres9412 Obj

																							if True == ifres9413 {
																								ifres9412 = True

																							} else {
																								ifres9412 = False

																							}

																							ifres9411 = ifres9412

																						} else {
																							ifres9411 = False

																						}

																						var ifres9410 Obj

																						if True == ifres9411 {
																							ifres9410 = True

																						} else {
																							ifres9410 = False

																						}

																						ifres9409 = ifres9410

																					} else {
																						ifres9409 = False

																					}

																					var ifres9408 Obj

																					if True == ifres9409 {
																						ifres9408 = True

																					} else {
																						ifres9408 = False

																					}

																					ifres9407 = ifres9408

																				} else {
																					ifres9407 = False

																				}

																				var ifres9406 Obj

																				if True == ifres9407 {
																					ifres9406 = True

																				} else {
																					ifres9406 = False

																				}

																				ifres9405 = ifres9406

																			} else {
																				ifres9405 = False

																			}

																			var ifres9404 Obj

																			if True == ifres9405 {
																				ifres9404 = True

																			} else {
																				ifres9404 = False

																			}

																			ifres9403 = ifres9404

																		} else {
																			ifres9403 = False

																		}

																		var ifres9402 Obj

																		if True == ifres9403 {
																			ifres9402 = True

																		} else {
																			ifres9402 = False

																		}

																		ifres9401 = ifres9402

																	} else {
																		ifres9401 = False

																	}

																	var ifres9400 Obj

																	if True == ifres9401 {
																		ifres9400 = True

																	} else {
																		ifres9400 = False

																	}

																	ifres9399 = ifres9400

																} else {
																	ifres9399 = False

																}

																var ifres9398 Obj

																if True == ifres9399 {
																	ifres9398 = True

																} else {
																	ifres9398 = False

																}

																ifres9397 = ifres9398

															} else {
																ifres9397 = False

															}

															var ifres9396 Obj

															if True == ifres9397 {
																ifres9396 = True

															} else {
																ifres9396 = False

															}

															ifres9395 = ifres9396

														} else {
															ifres9395 = False

														}

														var ifres9394 Obj

														if True == ifres9395 {
															ifres9394 = True

														} else {
															ifres9394 = False

														}

														ifres9393 = ifres9394

													} else {
														ifres9393 = False

													}

													var ifres9392 Obj

													if True == ifres9393 {
														ifres9392 = True

													} else {
														ifres9392 = False

													}

													ifres9391 = ifres9392

												} else {
													ifres9391 = False

												}

												var ifres9390 Obj

												if True == ifres9391 {
													ifres9390 = True

												} else {
													ifres9390 = False

												}

												ifres9389 = ifres9390

											} else {
												ifres9389 = False

											}

											var ifres9388 Obj

											if True == ifres9389 {
												ifres9388 = True

											} else {
												ifres9388 = False

											}

											ifres9387 = ifres9388

										} else {
											ifres9387 = False

										}

										var ifres9386 Obj

										if True == ifres9387 {
											ifres9386 = True

										} else {
											ifres9386 = False

										}

										ifres9385 = ifres9386

									} else {
										ifres9385 = False

									}

									var ifres9384 Obj

									if True == ifres9385 {
										ifres9384 = True

									} else {
										ifres9384 = False

									}

									ifres9383 = ifres9384

								} else {
									ifres9383 = False

								}

								var ifres9382 Obj

								if True == ifres9383 {
									ifres9382 = True

								} else {
									ifres9382 = False

								}

								ifres9381 = ifres9382

							} else {
								ifres9381 = False

							}

							var ifres9380 Obj

							if True == ifres9381 {
								ifres9380 = True

							} else {
								ifres9380 = False

							}

							ifres9379 = ifres9380

						} else {
							ifres9379 = False

						}

						var ifres9378 Obj

						if True == ifres9379 {
							ifres9378 = True

						} else {
							ifres9378 = False

						}

						ifres9377 = ifres9378

					} else {
						ifres9377 = False

					}

					var ifres9376 Obj

					if True == ifres9377 {
						ifres9376 = True

					} else {
						ifres9376 = False

					}

					ifres9375 = ifres9376

				} else {
					ifres9375 = False

				}

				var ifres9374 Obj

				if True == ifres9375 {
					ifres9374 = True

				} else {
					ifres9374 = False

				}

				ifres9373 = ifres9374

			} else {
				ifres9373 = False

			}

			var ifres9372 Obj

			if True == ifres9373 {
				ifres9372 = True

			} else {
				ifres9372 = False

			}

			ifres9371 = ifres9372

		} else {
			ifres9371 = False

		}

		if True == ifres9371 {
			tmp9359 := PrimTail(V1737)

			tmp9360 := PrimTail(tmp9359)

			tmp9361 := PrimHead(tmp9360)

			tmp9362 := PrimTail(tmp9361)

			tmp9363 := PrimHead(tmp9362)

			tmp9364 := PrimTail(V1737)

			tmp9365 := PrimTail(tmp9364)

			tmp9366 := PrimHead(tmp9365)

			tmp9367 := PrimTail(tmp9366)

			tmp9368 := PrimTail(tmp9367)

			tmp9369 := PrimCons(sym_a_a_6, tmp9368)

			__e.Return(PrimCons(tmp9363, tmp9369))
			return

		} else {
			__e.Return(V1737)
			return
		}

	}, 1)

	tmp9535 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9358)

	_ = tmp9535

	tmp9536 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("type error\n")))
		return
	}, 0)

	tmp9537 := Call(__e, ns2_1set, symshen_4type_1error, tmp9536)

	_ = tmp9537

	tmp9538 := MakeNative(func(__e *ControlFlow) {
		V1738 := __e.Get(1)
		_ = V1738
		tmp9539 := MakeNative(func(__e *ControlFlow) {
			KLFile := __e.Get(1)
			_ = KLFile
			tmp9540 := MakeNative(func(__e *ControlFlow) {
				Code := __e.Get(1)
				_ = Code
				tmp9541 := MakeNative(func(__e *ControlFlow) {
					Open := __e.Get(1)
					_ = Open
					tmp9542 := MakeNative(func(__e *ControlFlow) {
						KL := __e.Get(1)
						_ = KL
						tmp9543 := MakeNative(func(__e *ControlFlow) {
							Write := __e.Get(1)
							_ = Write
							__e.Return(KLFile)
							return
						}, 1)

						tmp9544 := Call(__e, PrimFunc(symshen_4write_1kl), KL, Open)

						__e.TailApply(tmp9543, tmp9544)
						return

					}, 1)

					tmp9545 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), X)
						return
					}, 1)

					tmp9546 := Call(__e, PrimFunc(symmap), tmp9545, Code)

					__e.TailApply(tmp9542, tmp9546)
					return

				}, 1)

				tmp9547 := PrimOpenStream(KLFile, symout)

				__e.TailApply(tmp9541, tmp9547)
				return

			}, 1)

			tmp9548 := Call(__e, PrimFunc(symread_1file), V1738)

			__e.TailApply(tmp9540, tmp9548)
			return

		}, 1)

		tmp9549 := Call(__e, PrimFunc(symshen_4klfile), V1738)

		__e.TailApply(tmp9539, tmp9549)
		return

	}, 1)

	tmp9550 := Call(__e, ns2_1set, symbootstrap, tmp9538)

	_ = tmp9550

	tmp9551 := MakeNative(func(__e *ControlFlow) {
		V1741 := __e.Get(1)
		_ = V1741
		V1742 := __e.Get(2)
		_ = V1742
		tmp9565 := PrimEqual(Nil, V1741)

		if True == tmp9565 {
			__e.Return(PrimCloseStream(V1742))
			return
		} else {
			tmp9563 := PrimIsPair(V1741)

			var ifres9559 Obj

			if True == tmp9563 {
				tmp9561 := PrimHead(V1741)

				tmp9562 := PrimIsPair(tmp9561)

				var ifres9560 Obj

				if True == tmp9562 {
					ifres9560 = True

				} else {
					ifres9560 = False

				}

				ifres9559 = ifres9560

			} else {
				ifres9559 = False

			}

			if True == ifres9559 {
				tmp9552 := PrimTail(V1741)

				tmp9553 := PrimHead(V1741)

				tmp9554 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp9553, V1742)

				_ = tmp9554

				__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9552, V1742)
				return

			} else {
				tmp9557 := PrimIsPair(V1741)

				if True == tmp9557 {
					tmp9555 := PrimTail(V1741)

					__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9555, V1742)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4write_1kl)
					return
				}

			}

		}

	}, 2)

	tmp9566 := Call(__e, ns2_1set, symshen_4write_1kl, tmp9551)

	_ = tmp9566

	tmp9567 := MakeNative(func(__e *ControlFlow) {
		V1745 := __e.Get(1)
		_ = V1745
		V1746 := __e.Get(2)
		_ = V1746
		tmp9607 := PrimIsPair(V1745)

		var ifres9570 Obj

		if True == tmp9607 {
			tmp9605 := PrimHead(V1745)

			tmp9606 := PrimEqual(symdefun, tmp9605)

			var ifres9572 Obj

			if True == tmp9606 {
				tmp9603 := PrimTail(V1745)

				tmp9604 := PrimIsPair(tmp9603)

				var ifres9574 Obj

				if True == tmp9604 {
					tmp9600 := PrimTail(V1745)

					tmp9601 := PrimHead(tmp9600)

					tmp9602 := PrimEqual(symfail, tmp9601)

					var ifres9576 Obj

					if True == tmp9602 {
						tmp9597 := PrimTail(V1745)

						tmp9598 := PrimTail(tmp9597)

						tmp9599 := PrimIsPair(tmp9598)

						var ifres9578 Obj

						if True == tmp9599 {
							tmp9593 := PrimTail(V1745)

							tmp9594 := PrimTail(tmp9593)

							tmp9595 := PrimHead(tmp9594)

							tmp9596 := PrimEqual(Nil, tmp9595)

							var ifres9580 Obj

							if True == tmp9596 {
								tmp9589 := PrimTail(V1745)

								tmp9590 := PrimTail(tmp9589)

								tmp9591 := PrimTail(tmp9590)

								tmp9592 := PrimIsPair(tmp9591)

								var ifres9582 Obj

								if True == tmp9592 {
									tmp9584 := PrimTail(V1745)

									tmp9585 := PrimTail(tmp9584)

									tmp9586 := PrimTail(tmp9585)

									tmp9587 := PrimTail(tmp9586)

									tmp9588 := PrimEqual(Nil, tmp9587)

									var ifres9583 Obj

									if True == tmp9588 {
										ifres9583 = True

									} else {
										ifres9583 = False

									}

									ifres9582 = ifres9583

								} else {
									ifres9582 = False

								}

								var ifres9581 Obj

								if True == ifres9582 {
									ifres9581 = True

								} else {
									ifres9581 = False

								}

								ifres9580 = ifres9581

							} else {
								ifres9580 = False

							}

							var ifres9579 Obj

							if True == ifres9580 {
								ifres9579 = True

							} else {
								ifres9579 = False

							}

							ifres9578 = ifres9579

						} else {
							ifres9578 = False

						}

						var ifres9577 Obj

						if True == ifres9578 {
							ifres9577 = True

						} else {
							ifres9577 = False

						}

						ifres9576 = ifres9577

					} else {
						ifres9576 = False

					}

					var ifres9575 Obj

					if True == ifres9576 {
						ifres9575 = True

					} else {
						ifres9575 = False

					}

					ifres9574 = ifres9575

				} else {
					ifres9574 = False

				}

				var ifres9573 Obj

				if True == ifres9574 {
					ifres9573 = True

				} else {
					ifres9573 = False

				}

				ifres9572 = ifres9573

			} else {
				ifres9572 = False

			}

			var ifres9571 Obj

			if True == ifres9572 {
				ifres9571 = True

			} else {
				ifres9571 = False

			}

			ifres9570 = ifres9571

		} else {
			ifres9570 = False

		}

		if True == ifres9570 {
			__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V1746)
			return
		} else {
			tmp9568 := Call(__e, PrimFunc(symshen_4app), V1745, MakeString("\n\n"), symshen_4r)

			__e.TailApply(PrimFunc(sympr), tmp9568, V1746)
			return

		}

	}, 2)

	tmp9608 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp9567)

	_ = tmp9608

	tmp9609 := MakeNative(func(__e *ControlFlow) {
		V1747 := __e.Get(1)
		_ = V1747
		tmp9618 := PrimEqual(MakeString(""), V1747)

		if True == tmp9618 {
			__e.Return(MakeString(".kl"))
			return
		} else {
			tmp9616 := PrimEqual(MakeString(".shen"), V1747)

			if True == tmp9616 {
				__e.Return(MakeString(".kl"))
				return
			} else {
				tmp9614 := Call(__e, PrimFunc(symshen_4_7string_2), V1747)

				if True == tmp9614 {
					tmp9610 := Call(__e, PrimFunc(symhdstr), V1747)

					tmp9611 := PrimTailString(V1747)

					tmp9612 := Call(__e, PrimFunc(symshen_4klfile), tmp9611)

					__e.TailApply(PrimFunc(sym_8s), tmp9610, tmp9612)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4klfile)
					return
				}

			}

		}

	}, 1)

	__e.TailApply(ns2_1set, symshen_4klfile, tmp9609)
	return

}, 0)
