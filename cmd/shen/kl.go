package main

import . "github.com/tiancaiamao/shen-go/klambda"

var KLMain = MakeNative(func(__e *ControlFlow) {
	tmp152 := MakeNative(func(__e *ControlFlow) {
		p123 := __e.Get(1)
		_ = p123
		p124 := __e.Get(2)
		_ = p124
		tmp153 := MakeNative(func(__e *ControlFlow) {
			val125 := __e.Get(1)
			_ = val125
			tmp154 := MakeNative(func(__e *ControlFlow) {
				cc126 := __e.Get(1)
				_ = cc126
				tmp178 := PrimIsPair(val125)

				var ifres174 Obj

				if True == tmp178 {
					tmp176 := Call(__e, PrimNS1Value(symnull_2), val125)

					tmp177 := PrimNot(tmp176)

					var ifres175 Obj

					if True == tmp177 {
						ifres175 = True

					} else {
						ifres175 = False

					}

					ifres174 = ifres175

				} else {
					ifres174 = False

				}

				if True == ifres174 {
					tmp156 := MakeNative(func(__e *ControlFlow) {
						env := __e.Get(1)
						_ = env
						tmp171 := PrimTail(val125)

						tmp172 := PrimIsPair(tmp171)

						var ifres166 Obj

						if True == tmp172 {
							tmp168 := PrimTail(val125)

							tmp169 := Call(__e, PrimNS1Value(symnull_2), tmp168)

							tmp170 := PrimNot(tmp169)

							var ifres167 Obj

							if True == tmp170 {
								ifres167 = True

							} else {
								ifres167 = False

							}

							ifres166 = ifres167

						} else {
							ifres166 = False

						}

						if True == ifres166 {
							tmp163 := PrimTail(val125)

							tmp164 := PrimHead(tmp163)

							tmp165 := PrimEqual(Nil, tmp164)

							if True == tmp165 {
								tmp160 := PrimTail(val125)

								tmp161 := PrimTail(tmp160)

								tmp162 := Call(__e, PrimNS1Value(symnull_2), tmp161)

								if True == tmp162 {
									__e.Return(Nil)
									return
								} else {
									__e.TailApply(cc126)
									return
								}

							} else {
								__e.TailApply(cc126)
								return
							}

						} else {
							__e.TailApply(cc126)
							return
						}

					}, 1)

					tmp173 := PrimHead(val125)

					__e.TailApply(tmp156, tmp173)
					return

				} else {
					__e.TailApply(cc126)
					return
				}

			}, 1)

			tmp179 := MakeNative(func(__e *ControlFlow) {
				tmp180 := MakeNative(func(__e *ControlFlow) {
					cc127 := __e.Get(1)
					_ = cc127
					tmp210 := PrimIsPair(val125)

					var ifres206 Obj

					if True == tmp210 {
						tmp208 := Call(__e, PrimNS1Value(symnull_2), val125)

						tmp209 := PrimNot(tmp208)

						var ifres207 Obj

						if True == tmp209 {
							ifres207 = True

						} else {
							ifres207 = False

						}

						ifres206 = ifres207

					} else {
						ifres206 = False

					}

					if True == ifres206 {
						tmp182 := MakeNative(func(__e *ControlFlow) {
							env := __e.Get(1)
							_ = env
							tmp203 := PrimTail(val125)

							tmp204 := PrimIsPair(tmp203)

							var ifres198 Obj

							if True == tmp204 {
								tmp200 := PrimTail(val125)

								tmp201 := Call(__e, PrimNS1Value(symnull_2), tmp200)

								tmp202 := PrimNot(tmp201)

								var ifres199 Obj

								if True == tmp202 {
									ifres199 = True

								} else {
									ifres199 = False

								}

								ifres198 = ifres199

							} else {
								ifres198 = False

							}

							if True == ifres198 {
								tmp184 := MakeNative(func(__e *ControlFlow) {
									x := __e.Get(1)
									_ = x
									tmp193 := PrimTail(val125)

									tmp194 := PrimTail(tmp193)

									tmp195 := Call(__e, PrimNS1Value(symnull_2), tmp194)

									if True == tmp195 {
										tmp192 := PrimIsNumber(x)

										var ifres187 Obj

										if True == tmp192 {
											ifres187 = True

										} else {
											tmp191 := PrimIsString(x)

											var ifres188 Obj

											if True == tmp191 {
												ifres188 = True

											} else {
												tmp190 := Call(__e, PrimNS1Value(symboolean_2), x)

												var ifres189 Obj

												if True == tmp190 {
													ifres189 = True

												} else {
													ifres189 = False

												}

												ifres188 = ifres189

											}

											ifres187 = ifres188

										}

										if True == ifres187 {
											__e.Return(x)
											return
										} else {
											__e.TailApply(cc127)
											return
										}

									} else {
										__e.TailApply(cc127)
										return
									}

								}, 1)

								tmp196 := PrimTail(val125)

								tmp197 := PrimHead(tmp196)

								__e.TailApply(tmp184, tmp197)
								return

							} else {
								__e.TailApply(cc127)
								return
							}

						}, 1)

						tmp205 := PrimHead(val125)

						__e.TailApply(tmp182, tmp205)
						return

					} else {
						__e.TailApply(cc127)
						return
					}

				}, 1)

				tmp211 := MakeNative(func(__e *ControlFlow) {
					tmp212 := MakeNative(func(__e *ControlFlow) {
						cc128 := __e.Get(1)
						_ = cc128
						tmp321 := PrimIsPair(val125)

						var ifres317 Obj

						if True == tmp321 {
							tmp319 := Call(__e, PrimNS1Value(symnull_2), val125)

							tmp320 := PrimNot(tmp319)

							var ifres318 Obj

							if True == tmp320 {
								ifres318 = True

							} else {
								ifres318 = False

							}

							ifres317 = ifres318

						} else {
							ifres317 = False

						}

						if True == ifres317 {
							tmp214 := MakeNative(func(__e *ControlFlow) {
								env := __e.Get(1)
								_ = env
								tmp314 := PrimTail(val125)

								tmp315 := PrimIsPair(tmp314)

								var ifres309 Obj

								if True == tmp315 {
									tmp311 := PrimTail(val125)

									tmp312 := Call(__e, PrimNS1Value(symnull_2), tmp311)

									tmp313 := PrimNot(tmp312)

									var ifres310 Obj

									if True == tmp313 {
										ifres310 = True

									} else {
										ifres310 = False

									}

									ifres309 = ifres310

								} else {
									ifres309 = False

								}

								if True == ifres309 {
									tmp306 := PrimTail(val125)

									tmp307 := PrimHead(tmp306)

									tmp308 := PrimIsPair(tmp307)

									var ifres300 Obj

									if True == tmp308 {
										tmp302 := PrimTail(val125)

										tmp303 := PrimHead(tmp302)

										tmp304 := Call(__e, PrimNS1Value(symnull_2), tmp303)

										tmp305 := PrimNot(tmp304)

										var ifres301 Obj

										if True == tmp305 {
											ifres301 = True

										} else {
											ifres301 = False

										}

										ifres300 = ifres301

									} else {
										ifres300 = False

									}

									if True == ifres300 {
										tmp296 := PrimTail(val125)

										tmp297 := PrimHead(tmp296)

										tmp298 := PrimHead(tmp297)

										tmp299 := PrimEqual(symif, tmp298)

										if True == tmp299 {
											tmp292 := PrimTail(val125)

											tmp293 := PrimHead(tmp292)

											tmp294 := PrimTail(tmp293)

											tmp295 := PrimIsPair(tmp294)

											var ifres285 Obj

											if True == tmp295 {
												tmp287 := PrimTail(val125)

												tmp288 := PrimHead(tmp287)

												tmp289 := PrimTail(tmp288)

												tmp290 := Call(__e, PrimNS1Value(symnull_2), tmp289)

												tmp291 := PrimNot(tmp290)

												var ifres286 Obj

												if True == tmp291 {
													ifres286 = True

												} else {
													ifres286 = False

												}

												ifres285 = ifres286

											} else {
												ifres285 = False

											}

											if True == ifres285 {
												tmp219 := MakeNative(func(__e *ControlFlow) {
													x := __e.Get(1)
													_ = x
													tmp276 := PrimTail(val125)

													tmp277 := PrimHead(tmp276)

													tmp278 := PrimTail(tmp277)

													tmp279 := PrimTail(tmp278)

													tmp280 := PrimIsPair(tmp279)

													var ifres268 Obj

													if True == tmp280 {
														tmp270 := PrimTail(val125)

														tmp271 := PrimHead(tmp270)

														tmp272 := PrimTail(tmp271)

														tmp273 := PrimTail(tmp272)

														tmp274 := Call(__e, PrimNS1Value(symnull_2), tmp273)

														tmp275 := PrimNot(tmp274)

														var ifres269 Obj

														if True == tmp275 {
															ifres269 = True

														} else {
															ifres269 = False

														}

														ifres268 = ifres269

													} else {
														ifres268 = False

													}

													if True == ifres268 {
														tmp221 := MakeNative(func(__e *ControlFlow) {
															y := __e.Get(1)
															_ = y
															tmp257 := PrimTail(val125)

															tmp258 := PrimHead(tmp257)

															tmp259 := PrimTail(tmp258)

															tmp260 := PrimTail(tmp259)

															tmp261 := PrimTail(tmp260)

															tmp262 := PrimIsPair(tmp261)

															var ifres248 Obj

															if True == tmp262 {
																tmp250 := PrimTail(val125)

																tmp251 := PrimHead(tmp250)

																tmp252 := PrimTail(tmp251)

																tmp253 := PrimTail(tmp252)

																tmp254 := PrimTail(tmp253)

																tmp255 := Call(__e, PrimNS1Value(symnull_2), tmp254)

																tmp256 := PrimNot(tmp255)

																var ifres249 Obj

																if True == tmp256 {
																	ifres249 = True

																} else {
																	ifres249 = False

																}

																ifres248 = ifres249

															} else {
																ifres248 = False

															}

															if True == ifres248 {
																tmp223 := MakeNative(func(__e *ControlFlow) {
																	z := __e.Get(1)
																	_ = z
																	tmp235 := PrimTail(val125)

																	tmp236 := PrimHead(tmp235)

																	tmp237 := PrimTail(tmp236)

																	tmp238 := PrimTail(tmp237)

																	tmp239 := PrimTail(tmp238)

																	tmp240 := PrimTail(tmp239)

																	tmp241 := Call(__e, PrimNS1Value(symnull_2), tmp240)

																	if True == tmp241 {
																		tmp232 := PrimTail(val125)

																		tmp233 := PrimTail(tmp232)

																		tmp234 := Call(__e, PrimNS1Value(symnull_2), tmp233)

																		if True == tmp234 {
																			tmp226 := Call(__e, PrimNS1Value(symkl_1_6cora), env, x)

																			tmp227 := Call(__e, PrimNS1Value(symkl_1_6cora), env, y)

																			tmp228 := Call(__e, PrimNS1Value(symkl_1_6cora), env, z)

																			tmp229 := PrimCons(tmp228, Nil)

																			tmp230 := PrimCons(tmp227, tmp229)

																			tmp231 := PrimCons(tmp226, tmp230)

																			__e.Return(PrimCons(symif, tmp231))
																			return

																		} else {
																			__e.TailApply(cc128)
																			return
																		}

																	} else {
																		__e.TailApply(cc128)
																		return
																	}

																}, 1)

																tmp242 := PrimTail(val125)

																tmp243 := PrimHead(tmp242)

																tmp244 := PrimTail(tmp243)

																tmp245 := PrimTail(tmp244)

																tmp246 := PrimTail(tmp245)

																tmp247 := PrimHead(tmp246)

																__e.TailApply(tmp223, tmp247)
																return

															} else {
																__e.TailApply(cc128)
																return
															}

														}, 1)

														tmp263 := PrimTail(val125)

														tmp264 := PrimHead(tmp263)

														tmp265 := PrimTail(tmp264)

														tmp266 := PrimTail(tmp265)

														tmp267 := PrimHead(tmp266)

														__e.TailApply(tmp221, tmp267)
														return

													} else {
														__e.TailApply(cc128)
														return
													}

												}, 1)

												tmp281 := PrimTail(val125)

												tmp282 := PrimHead(tmp281)

												tmp283 := PrimTail(tmp282)

												tmp284 := PrimHead(tmp283)

												__e.TailApply(tmp219, tmp284)
												return

											} else {
												__e.TailApply(cc128)
												return
											}

										} else {
											__e.TailApply(cc128)
											return
										}

									} else {
										__e.TailApply(cc128)
										return
									}

								} else {
									__e.TailApply(cc128)
									return
								}

							}, 1)

							tmp316 := PrimHead(val125)

							__e.TailApply(tmp214, tmp316)
							return

						} else {
							__e.TailApply(cc128)
							return
						}

					}, 1)

					tmp322 := MakeNative(func(__e *ControlFlow) {
						tmp323 := MakeNative(func(__e *ControlFlow) {
							cc129 := __e.Get(1)
							_ = cc129
							tmp406 := PrimIsPair(val125)

							var ifres402 Obj

							if True == tmp406 {
								tmp404 := Call(__e, PrimNS1Value(symnull_2), val125)

								tmp405 := PrimNot(tmp404)

								var ifres403 Obj

								if True == tmp405 {
									ifres403 = True

								} else {
									ifres403 = False

								}

								ifres402 = ifres403

							} else {
								ifres402 = False

							}

							if True == ifres402 {
								tmp325 := MakeNative(func(__e *ControlFlow) {
									env := __e.Get(1)
									_ = env
									tmp399 := PrimTail(val125)

									tmp400 := PrimIsPair(tmp399)

									var ifres394 Obj

									if True == tmp400 {
										tmp396 := PrimTail(val125)

										tmp397 := Call(__e, PrimNS1Value(symnull_2), tmp396)

										tmp398 := PrimNot(tmp397)

										var ifres395 Obj

										if True == tmp398 {
											ifres395 = True

										} else {
											ifres395 = False

										}

										ifres394 = ifres395

									} else {
										ifres394 = False

									}

									if True == ifres394 {
										tmp391 := PrimTail(val125)

										tmp392 := PrimHead(tmp391)

										tmp393 := PrimIsPair(tmp392)

										var ifres385 Obj

										if True == tmp393 {
											tmp387 := PrimTail(val125)

											tmp388 := PrimHead(tmp387)

											tmp389 := Call(__e, PrimNS1Value(symnull_2), tmp388)

											tmp390 := PrimNot(tmp389)

											var ifres386 Obj

											if True == tmp390 {
												ifres386 = True

											} else {
												ifres386 = False

											}

											ifres385 = ifres386

										} else {
											ifres385 = False

										}

										if True == ifres385 {
											tmp381 := PrimTail(val125)

											tmp382 := PrimHead(tmp381)

											tmp383 := PrimHead(tmp382)

											tmp384 := PrimEqual(symdo, tmp383)

											if True == tmp384 {
												tmp377 := PrimTail(val125)

												tmp378 := PrimHead(tmp377)

												tmp379 := PrimTail(tmp378)

												tmp380 := PrimIsPair(tmp379)

												var ifres370 Obj

												if True == tmp380 {
													tmp372 := PrimTail(val125)

													tmp373 := PrimHead(tmp372)

													tmp374 := PrimTail(tmp373)

													tmp375 := Call(__e, PrimNS1Value(symnull_2), tmp374)

													tmp376 := PrimNot(tmp375)

													var ifres371 Obj

													if True == tmp376 {
														ifres371 = True

													} else {
														ifres371 = False

													}

													ifres370 = ifres371

												} else {
													ifres370 = False

												}

												if True == ifres370 {
													tmp330 := MakeNative(func(__e *ControlFlow) {
														x := __e.Get(1)
														_ = x
														tmp361 := PrimTail(val125)

														tmp362 := PrimHead(tmp361)

														tmp363 := PrimTail(tmp362)

														tmp364 := PrimTail(tmp363)

														tmp365 := PrimIsPair(tmp364)

														var ifres353 Obj

														if True == tmp365 {
															tmp355 := PrimTail(val125)

															tmp356 := PrimHead(tmp355)

															tmp357 := PrimTail(tmp356)

															tmp358 := PrimTail(tmp357)

															tmp359 := Call(__e, PrimNS1Value(symnull_2), tmp358)

															tmp360 := PrimNot(tmp359)

															var ifres354 Obj

															if True == tmp360 {
																ifres354 = True

															} else {
																ifres354 = False

															}

															ifres353 = ifres354

														} else {
															ifres353 = False

														}

														if True == ifres353 {
															tmp332 := MakeNative(func(__e *ControlFlow) {
																y := __e.Get(1)
																_ = y
																tmp342 := PrimTail(val125)

																tmp343 := PrimHead(tmp342)

																tmp344 := PrimTail(tmp343)

																tmp345 := PrimTail(tmp344)

																tmp346 := PrimTail(tmp345)

																tmp347 := Call(__e, PrimNS1Value(symnull_2), tmp346)

																if True == tmp347 {
																	tmp339 := PrimTail(val125)

																	tmp340 := PrimTail(tmp339)

																	tmp341 := Call(__e, PrimNS1Value(symnull_2), tmp340)

																	if True == tmp341 {
																		tmp335 := Call(__e, PrimNS1Value(symkl_1_6cora), env, x)

																		tmp336 := Call(__e, PrimNS1Value(symkl_1_6cora), env, y)

																		tmp337 := PrimCons(tmp336, Nil)

																		tmp338 := PrimCons(tmp335, tmp337)

																		__e.Return(PrimCons(symdo, tmp338))
																		return

																	} else {
																		__e.TailApply(cc129)
																		return
																	}

																} else {
																	__e.TailApply(cc129)
																	return
																}

															}, 1)

															tmp348 := PrimTail(val125)

															tmp349 := PrimHead(tmp348)

															tmp350 := PrimTail(tmp349)

															tmp351 := PrimTail(tmp350)

															tmp352 := PrimHead(tmp351)

															__e.TailApply(tmp332, tmp352)
															return

														} else {
															__e.TailApply(cc129)
															return
														}

													}, 1)

													tmp366 := PrimTail(val125)

													tmp367 := PrimHead(tmp366)

													tmp368 := PrimTail(tmp367)

													tmp369 := PrimHead(tmp368)

													__e.TailApply(tmp330, tmp369)
													return

												} else {
													__e.TailApply(cc129)
													return
												}

											} else {
												__e.TailApply(cc129)
												return
											}

										} else {
											__e.TailApply(cc129)
											return
										}

									} else {
										__e.TailApply(cc129)
										return
									}

								}, 1)

								tmp401 := PrimHead(val125)

								__e.TailApply(tmp325, tmp401)
								return

							} else {
								__e.TailApply(cc129)
								return
							}

						}, 1)

						tmp407 := MakeNative(func(__e *ControlFlow) {
							tmp408 := MakeNative(func(__e *ControlFlow) {
								cc130 := __e.Get(1)
								_ = cc130
								tmp492 := PrimIsPair(val125)

								var ifres488 Obj

								if True == tmp492 {
									tmp490 := Call(__e, PrimNS1Value(symnull_2), val125)

									tmp491 := PrimNot(tmp490)

									var ifres489 Obj

									if True == tmp491 {
										ifres489 = True

									} else {
										ifres489 = False

									}

									ifres488 = ifres489

								} else {
									ifres488 = False

								}

								if True == ifres488 {
									tmp410 := MakeNative(func(__e *ControlFlow) {
										env := __e.Get(1)
										_ = env
										tmp485 := PrimTail(val125)

										tmp486 := PrimIsPair(tmp485)

										var ifres480 Obj

										if True == tmp486 {
											tmp482 := PrimTail(val125)

											tmp483 := Call(__e, PrimNS1Value(symnull_2), tmp482)

											tmp484 := PrimNot(tmp483)

											var ifres481 Obj

											if True == tmp484 {
												ifres481 = True

											} else {
												ifres481 = False

											}

											ifres480 = ifres481

										} else {
											ifres480 = False

										}

										if True == ifres480 {
											tmp477 := PrimTail(val125)

											tmp478 := PrimHead(tmp477)

											tmp479 := PrimIsPair(tmp478)

											var ifres471 Obj

											if True == tmp479 {
												tmp473 := PrimTail(val125)

												tmp474 := PrimHead(tmp473)

												tmp475 := Call(__e, PrimNS1Value(symnull_2), tmp474)

												tmp476 := PrimNot(tmp475)

												var ifres472 Obj

												if True == tmp476 {
													ifres472 = True

												} else {
													ifres472 = False

												}

												ifres471 = ifres472

											} else {
												ifres471 = False

											}

											if True == ifres471 {
												tmp467 := PrimTail(val125)

												tmp468 := PrimHead(tmp467)

												tmp469 := PrimHead(tmp468)

												tmp470 := PrimEqual(symlambda, tmp469)

												if True == tmp470 {
													tmp463 := PrimTail(val125)

													tmp464 := PrimHead(tmp463)

													tmp465 := PrimTail(tmp464)

													tmp466 := PrimIsPair(tmp465)

													var ifres456 Obj

													if True == tmp466 {
														tmp458 := PrimTail(val125)

														tmp459 := PrimHead(tmp458)

														tmp460 := PrimTail(tmp459)

														tmp461 := Call(__e, PrimNS1Value(symnull_2), tmp460)

														tmp462 := PrimNot(tmp461)

														var ifres457 Obj

														if True == tmp462 {
															ifres457 = True

														} else {
															ifres457 = False

														}

														ifres456 = ifres457

													} else {
														ifres456 = False

													}

													if True == ifres456 {
														tmp415 := MakeNative(func(__e *ControlFlow) {
															x := __e.Get(1)
															_ = x
															tmp447 := PrimTail(val125)

															tmp448 := PrimHead(tmp447)

															tmp449 := PrimTail(tmp448)

															tmp450 := PrimTail(tmp449)

															tmp451 := PrimIsPair(tmp450)

															var ifres439 Obj

															if True == tmp451 {
																tmp441 := PrimTail(val125)

																tmp442 := PrimHead(tmp441)

																tmp443 := PrimTail(tmp442)

																tmp444 := PrimTail(tmp443)

																tmp445 := Call(__e, PrimNS1Value(symnull_2), tmp444)

																tmp446 := PrimNot(tmp445)

																var ifres440 Obj

																if True == tmp446 {
																	ifres440 = True

																} else {
																	ifres440 = False

																}

																ifres439 = ifres440

															} else {
																ifres439 = False

															}

															if True == ifres439 {
																tmp417 := MakeNative(func(__e *ControlFlow) {
																	y := __e.Get(1)
																	_ = y
																	tmp428 := PrimTail(val125)

																	tmp429 := PrimHead(tmp428)

																	tmp430 := PrimTail(tmp429)

																	tmp431 := PrimTail(tmp430)

																	tmp432 := PrimTail(tmp431)

																	tmp433 := Call(__e, PrimNS1Value(symnull_2), tmp432)

																	if True == tmp433 {
																		tmp425 := PrimTail(val125)

																		tmp426 := PrimTail(tmp425)

																		tmp427 := Call(__e, PrimNS1Value(symnull_2), tmp426)

																		if True == tmp427 {
																			tmp420 := PrimCons(x, Nil)

																			tmp421 := PrimCons(x, env)

																			tmp422 := Call(__e, PrimNS1Value(symkl_1_6cora), tmp421, y)

																			tmp423 := PrimCons(tmp422, Nil)

																			tmp424 := PrimCons(tmp420, tmp423)

																			__e.Return(PrimCons(symlambda, tmp424))
																			return

																		} else {
																			__e.TailApply(cc130)
																			return
																		}

																	} else {
																		__e.TailApply(cc130)
																		return
																	}

																}, 1)

																tmp434 := PrimTail(val125)

																tmp435 := PrimHead(tmp434)

																tmp436 := PrimTail(tmp435)

																tmp437 := PrimTail(tmp436)

																tmp438 := PrimHead(tmp437)

																__e.TailApply(tmp417, tmp438)
																return

															} else {
																__e.TailApply(cc130)
																return
															}

														}, 1)

														tmp452 := PrimTail(val125)

														tmp453 := PrimHead(tmp452)

														tmp454 := PrimTail(tmp453)

														tmp455 := PrimHead(tmp454)

														__e.TailApply(tmp415, tmp455)
														return

													} else {
														__e.TailApply(cc130)
														return
													}

												} else {
													__e.TailApply(cc130)
													return
												}

											} else {
												__e.TailApply(cc130)
												return
											}

										} else {
											__e.TailApply(cc130)
											return
										}

									}, 1)

									tmp487 := PrimHead(val125)

									__e.TailApply(tmp410, tmp487)
									return

								} else {
									__e.TailApply(cc130)
									return
								}

							}, 1)

							tmp493 := MakeNative(func(__e *ControlFlow) {
								tmp494 := MakeNative(func(__e *ControlFlow) {
									cc131 := __e.Get(1)
									_ = cc131
									tmp573 := PrimIsPair(val125)

									var ifres569 Obj

									if True == tmp573 {
										tmp571 := Call(__e, PrimNS1Value(symnull_2), val125)

										tmp572 := PrimNot(tmp571)

										var ifres570 Obj

										if True == tmp572 {
											ifres570 = True

										} else {
											ifres570 = False

										}

										ifres569 = ifres570

									} else {
										ifres569 = False

									}

									if True == ifres569 {
										tmp496 := MakeNative(func(__e *ControlFlow) {
											env := __e.Get(1)
											_ = env
											tmp566 := PrimTail(val125)

											tmp567 := PrimIsPair(tmp566)

											var ifres561 Obj

											if True == tmp567 {
												tmp563 := PrimTail(val125)

												tmp564 := Call(__e, PrimNS1Value(symnull_2), tmp563)

												tmp565 := PrimNot(tmp564)

												var ifres562 Obj

												if True == tmp565 {
													ifres562 = True

												} else {
													ifres562 = False

												}

												ifres561 = ifres562

											} else {
												ifres561 = False

											}

											if True == ifres561 {
												tmp558 := PrimTail(val125)

												tmp559 := PrimHead(tmp558)

												tmp560 := PrimIsPair(tmp559)

												var ifres552 Obj

												if True == tmp560 {
													tmp554 := PrimTail(val125)

													tmp555 := PrimHead(tmp554)

													tmp556 := Call(__e, PrimNS1Value(symnull_2), tmp555)

													tmp557 := PrimNot(tmp556)

													var ifres553 Obj

													if True == tmp557 {
														ifres553 = True

													} else {
														ifres553 = False

													}

													ifres552 = ifres553

												} else {
													ifres552 = False

												}

												if True == ifres552 {
													tmp548 := PrimTail(val125)

													tmp549 := PrimHead(tmp548)

													tmp550 := PrimHead(tmp549)

													tmp551 := PrimEqual(symtype, tmp550)

													if True == tmp551 {
														tmp544 := PrimTail(val125)

														tmp545 := PrimHead(tmp544)

														tmp546 := PrimTail(tmp545)

														tmp547 := PrimIsPair(tmp546)

														var ifres537 Obj

														if True == tmp547 {
															tmp539 := PrimTail(val125)

															tmp540 := PrimHead(tmp539)

															tmp541 := PrimTail(tmp540)

															tmp542 := Call(__e, PrimNS1Value(symnull_2), tmp541)

															tmp543 := PrimNot(tmp542)

															var ifres538 Obj

															if True == tmp543 {
																ifres538 = True

															} else {
																ifres538 = False

															}

															ifres537 = ifres538

														} else {
															ifres537 = False

														}

														if True == ifres537 {
															tmp501 := MakeNative(func(__e *ControlFlow) {
																exp := __e.Get(1)
																_ = exp
																tmp528 := PrimTail(val125)

																tmp529 := PrimHead(tmp528)

																tmp530 := PrimTail(tmp529)

																tmp531 := PrimTail(tmp530)

																tmp532 := PrimIsPair(tmp531)

																var ifres520 Obj

																if True == tmp532 {
																	tmp522 := PrimTail(val125)

																	tmp523 := PrimHead(tmp522)

																	tmp524 := PrimTail(tmp523)

																	tmp525 := PrimTail(tmp524)

																	tmp526 := Call(__e, PrimNS1Value(symnull_2), tmp525)

																	tmp527 := PrimNot(tmp526)

																	var ifres521 Obj

																	if True == tmp527 {
																		ifres521 = True

																	} else {
																		ifres521 = False

																	}

																	ifres520 = ifres521

																} else {
																	ifres520 = False

																}

																if True == ifres520 {
																	tmp503 := MakeNative(func(__e *ControlFlow) {
																		__ := __e.Get(1)
																		_ = __
																		tmp509 := PrimTail(val125)

																		tmp510 := PrimHead(tmp509)

																		tmp511 := PrimTail(tmp510)

																		tmp512 := PrimTail(tmp511)

																		tmp513 := PrimTail(tmp512)

																		tmp514 := Call(__e, PrimNS1Value(symnull_2), tmp513)

																		if True == tmp514 {
																			tmp506 := PrimTail(val125)

																			tmp507 := PrimTail(tmp506)

																			tmp508 := Call(__e, PrimNS1Value(symnull_2), tmp507)

																			if True == tmp508 {
																				__e.TailApply(PrimNS1Value(symkl_1_6cora), env, exp)
																				return
																			} else {
																				__e.TailApply(cc131)
																				return
																			}

																		} else {
																			__e.TailApply(cc131)
																			return
																		}

																	}, 1)

																	tmp515 := PrimTail(val125)

																	tmp516 := PrimHead(tmp515)

																	tmp517 := PrimTail(tmp516)

																	tmp518 := PrimTail(tmp517)

																	tmp519 := PrimHead(tmp518)

																	__e.TailApply(tmp503, tmp519)
																	return

																} else {
																	__e.TailApply(cc131)
																	return
																}

															}, 1)

															tmp533 := PrimTail(val125)

															tmp534 := PrimHead(tmp533)

															tmp535 := PrimTail(tmp534)

															tmp536 := PrimHead(tmp535)

															__e.TailApply(tmp501, tmp536)
															return

														} else {
															__e.TailApply(cc131)
															return
														}

													} else {
														__e.TailApply(cc131)
														return
													}

												} else {
													__e.TailApply(cc131)
													return
												}

											} else {
												__e.TailApply(cc131)
												return
											}

										}, 1)

										tmp568 := PrimHead(val125)

										__e.TailApply(tmp496, tmp568)
										return

									} else {
										__e.TailApply(cc131)
										return
									}

								}, 1)

								tmp574 := MakeNative(func(__e *ControlFlow) {
									tmp575 := MakeNative(func(__e *ControlFlow) {
										cc132 := __e.Get(1)
										_ = cc132
										tmp690 := PrimIsPair(val125)

										var ifres686 Obj

										if True == tmp690 {
											tmp688 := Call(__e, PrimNS1Value(symnull_2), val125)

											tmp689 := PrimNot(tmp688)

											var ifres687 Obj

											if True == tmp689 {
												ifres687 = True

											} else {
												ifres687 = False

											}

											ifres686 = ifres687

										} else {
											ifres686 = False

										}

										if True == ifres686 {
											tmp577 := MakeNative(func(__e *ControlFlow) {
												env := __e.Get(1)
												_ = env
												tmp683 := PrimTail(val125)

												tmp684 := PrimIsPair(tmp683)

												var ifres678 Obj

												if True == tmp684 {
													tmp680 := PrimTail(val125)

													tmp681 := Call(__e, PrimNS1Value(symnull_2), tmp680)

													tmp682 := PrimNot(tmp681)

													var ifres679 Obj

													if True == tmp682 {
														ifres679 = True

													} else {
														ifres679 = False

													}

													ifres678 = ifres679

												} else {
													ifres678 = False

												}

												if True == ifres678 {
													tmp675 := PrimTail(val125)

													tmp676 := PrimHead(tmp675)

													tmp677 := PrimIsPair(tmp676)

													var ifres669 Obj

													if True == tmp677 {
														tmp671 := PrimTail(val125)

														tmp672 := PrimHead(tmp671)

														tmp673 := Call(__e, PrimNS1Value(symnull_2), tmp672)

														tmp674 := PrimNot(tmp673)

														var ifres670 Obj

														if True == tmp674 {
															ifres670 = True

														} else {
															ifres670 = False

														}

														ifres669 = ifres670

													} else {
														ifres669 = False

													}

													if True == ifres669 {
														tmp665 := PrimTail(val125)

														tmp666 := PrimHead(tmp665)

														tmp667 := PrimHead(tmp666)

														tmp668 := PrimEqual(symdefun, tmp667)

														if True == tmp668 {
															tmp661 := PrimTail(val125)

															tmp662 := PrimHead(tmp661)

															tmp663 := PrimTail(tmp662)

															tmp664 := PrimIsPair(tmp663)

															var ifres654 Obj

															if True == tmp664 {
																tmp656 := PrimTail(val125)

																tmp657 := PrimHead(tmp656)

																tmp658 := PrimTail(tmp657)

																tmp659 := Call(__e, PrimNS1Value(symnull_2), tmp658)

																tmp660 := PrimNot(tmp659)

																var ifres655 Obj

																if True == tmp660 {
																	ifres655 = True

																} else {
																	ifres655 = False

																}

																ifres654 = ifres655

															} else {
																ifres654 = False

															}

															if True == ifres654 {
																tmp582 := MakeNative(func(__e *ControlFlow) {
																	f := __e.Get(1)
																	_ = f
																	tmp645 := PrimTail(val125)

																	tmp646 := PrimHead(tmp645)

																	tmp647 := PrimTail(tmp646)

																	tmp648 := PrimTail(tmp647)

																	tmp649 := PrimIsPair(tmp648)

																	var ifres637 Obj

																	if True == tmp649 {
																		tmp639 := PrimTail(val125)

																		tmp640 := PrimHead(tmp639)

																		tmp641 := PrimTail(tmp640)

																		tmp642 := PrimTail(tmp641)

																		tmp643 := Call(__e, PrimNS1Value(symnull_2), tmp642)

																		tmp644 := PrimNot(tmp643)

																		var ifres638 Obj

																		if True == tmp644 {
																			ifres638 = True

																		} else {
																			ifres638 = False

																		}

																		ifres637 = ifres638

																	} else {
																		ifres637 = False

																	}

																	if True == ifres637 {
																		tmp584 := MakeNative(func(__e *ControlFlow) {
																			x := __e.Get(1)
																			_ = x
																			tmp626 := PrimTail(val125)

																			tmp627 := PrimHead(tmp626)

																			tmp628 := PrimTail(tmp627)

																			tmp629 := PrimTail(tmp628)

																			tmp630 := PrimTail(tmp629)

																			tmp631 := PrimIsPair(tmp630)

																			var ifres617 Obj

																			if True == tmp631 {
																				tmp619 := PrimTail(val125)

																				tmp620 := PrimHead(tmp619)

																				tmp621 := PrimTail(tmp620)

																				tmp622 := PrimTail(tmp621)

																				tmp623 := PrimTail(tmp622)

																				tmp624 := Call(__e, PrimNS1Value(symnull_2), tmp623)

																				tmp625 := PrimNot(tmp624)

																				var ifres618 Obj

																				if True == tmp625 {
																					ifres618 = True

																				} else {
																					ifres618 = False

																				}

																				ifres617 = ifres618

																			} else {
																				ifres617 = False

																			}

																			if True == ifres617 {
																				tmp586 := MakeNative(func(__e *ControlFlow) {
																					y := __e.Get(1)
																					_ = y
																					tmp604 := PrimTail(val125)

																					tmp605 := PrimHead(tmp604)

																					tmp606 := PrimTail(tmp605)

																					tmp607 := PrimTail(tmp606)

																					tmp608 := PrimTail(tmp607)

																					tmp609 := PrimTail(tmp608)

																					tmp610 := Call(__e, PrimNS1Value(symnull_2), tmp609)

																					if True == tmp610 {
																						tmp601 := PrimTail(val125)

																						tmp602 := PrimTail(tmp601)

																						tmp603 := Call(__e, PrimNS1Value(symnull_2), tmp602)

																						if True == tmp603 {
																							tmp589 := PrimCons(symdef, Nil)

																							tmp590 := PrimCons(symquote, tmp589)

																							tmp591 := PrimCons(tmp590, Nil)

																							tmp592 := PrimCons(symfn, tmp591)

																							tmp593 := PrimCons(f, Nil)

																							tmp594 := PrimCons(symquote, tmp593)

																							tmp595 := Call(__e, PrimNS1Value(symkl_1_6cora), x, y)

																							tmp596 := PrimCons(tmp595, Nil)

																							tmp597 := PrimCons(x, tmp596)

																							tmp598 := PrimCons(symlambda, tmp597)

																							tmp599 := PrimCons(tmp598, Nil)

																							tmp600 := PrimCons(tmp594, tmp599)

																							__e.Return(PrimCons(tmp592, tmp600))
																							return

																						} else {
																							__e.TailApply(cc132)
																							return
																						}

																					} else {
																						__e.TailApply(cc132)
																						return
																					}

																				}, 1)

																				tmp611 := PrimTail(val125)

																				tmp612 := PrimHead(tmp611)

																				tmp613 := PrimTail(tmp612)

																				tmp614 := PrimTail(tmp613)

																				tmp615 := PrimTail(tmp614)

																				tmp616 := PrimHead(tmp615)

																				__e.TailApply(tmp586, tmp616)
																				return

																			} else {
																				__e.TailApply(cc132)
																				return
																			}

																		}, 1)

																		tmp632 := PrimTail(val125)

																		tmp633 := PrimHead(tmp632)

																		tmp634 := PrimTail(tmp633)

																		tmp635 := PrimTail(tmp634)

																		tmp636 := PrimHead(tmp635)

																		__e.TailApply(tmp584, tmp636)
																		return

																	} else {
																		__e.TailApply(cc132)
																		return
																	}

																}, 1)

																tmp650 := PrimTail(val125)

																tmp651 := PrimHead(tmp650)

																tmp652 := PrimTail(tmp651)

																tmp653 := PrimHead(tmp652)

																__e.TailApply(tmp582, tmp653)
																return

															} else {
																__e.TailApply(cc132)
																return
															}

														} else {
															__e.TailApply(cc132)
															return
														}

													} else {
														__e.TailApply(cc132)
														return
													}

												} else {
													__e.TailApply(cc132)
													return
												}

											}, 1)

											tmp685 := PrimHead(val125)

											__e.TailApply(tmp577, tmp685)
											return

										} else {
											__e.TailApply(cc132)
											return
										}

									}, 1)

									tmp691 := MakeNative(func(__e *ControlFlow) {
										tmp692 := MakeNative(func(__e *ControlFlow) {
											cc133 := __e.Get(1)
											_ = cc133
											tmp803 := PrimIsPair(val125)

											var ifres799 Obj

											if True == tmp803 {
												tmp801 := Call(__e, PrimNS1Value(symnull_2), val125)

												tmp802 := PrimNot(tmp801)

												var ifres800 Obj

												if True == tmp802 {
													ifres800 = True

												} else {
													ifres800 = False

												}

												ifres799 = ifres800

											} else {
												ifres799 = False

											}

											if True == ifres799 {
												tmp694 := MakeNative(func(__e *ControlFlow) {
													env := __e.Get(1)
													_ = env
													tmp796 := PrimTail(val125)

													tmp797 := PrimIsPair(tmp796)

													var ifres791 Obj

													if True == tmp797 {
														tmp793 := PrimTail(val125)

														tmp794 := Call(__e, PrimNS1Value(symnull_2), tmp793)

														tmp795 := PrimNot(tmp794)

														var ifres792 Obj

														if True == tmp795 {
															ifres792 = True

														} else {
															ifres792 = False

														}

														ifres791 = ifres792

													} else {
														ifres791 = False

													}

													if True == ifres791 {
														tmp788 := PrimTail(val125)

														tmp789 := PrimHead(tmp788)

														tmp790 := PrimIsPair(tmp789)

														var ifres782 Obj

														if True == tmp790 {
															tmp784 := PrimTail(val125)

															tmp785 := PrimHead(tmp784)

															tmp786 := Call(__e, PrimNS1Value(symnull_2), tmp785)

															tmp787 := PrimNot(tmp786)

															var ifres783 Obj

															if True == tmp787 {
																ifres783 = True

															} else {
																ifres783 = False

															}

															ifres782 = ifres783

														} else {
															ifres782 = False

														}

														if True == ifres782 {
															tmp778 := PrimTail(val125)

															tmp779 := PrimHead(tmp778)

															tmp780 := PrimHead(tmp779)

															tmp781 := PrimEqual(symlet, tmp780)

															if True == tmp781 {
																tmp774 := PrimTail(val125)

																tmp775 := PrimHead(tmp774)

																tmp776 := PrimTail(tmp775)

																tmp777 := PrimIsPair(tmp776)

																var ifres767 Obj

																if True == tmp777 {
																	tmp769 := PrimTail(val125)

																	tmp770 := PrimHead(tmp769)

																	tmp771 := PrimTail(tmp770)

																	tmp772 := Call(__e, PrimNS1Value(symnull_2), tmp771)

																	tmp773 := PrimNot(tmp772)

																	var ifres768 Obj

																	if True == tmp773 {
																		ifres768 = True

																	} else {
																		ifres768 = False

																	}

																	ifres767 = ifres768

																} else {
																	ifres767 = False

																}

																if True == ifres767 {
																	tmp699 := MakeNative(func(__e *ControlFlow) {
																		x := __e.Get(1)
																		_ = x
																		tmp758 := PrimTail(val125)

																		tmp759 := PrimHead(tmp758)

																		tmp760 := PrimTail(tmp759)

																		tmp761 := PrimTail(tmp760)

																		tmp762 := PrimIsPair(tmp761)

																		var ifres750 Obj

																		if True == tmp762 {
																			tmp752 := PrimTail(val125)

																			tmp753 := PrimHead(tmp752)

																			tmp754 := PrimTail(tmp753)

																			tmp755 := PrimTail(tmp754)

																			tmp756 := Call(__e, PrimNS1Value(symnull_2), tmp755)

																			tmp757 := PrimNot(tmp756)

																			var ifres751 Obj

																			if True == tmp757 {
																				ifres751 = True

																			} else {
																				ifres751 = False

																			}

																			ifres750 = ifres751

																		} else {
																			ifres750 = False

																		}

																		if True == ifres750 {
																			tmp701 := MakeNative(func(__e *ControlFlow) {
																				y := __e.Get(1)
																				_ = y
																				tmp739 := PrimTail(val125)

																				tmp740 := PrimHead(tmp739)

																				tmp741 := PrimTail(tmp740)

																				tmp742 := PrimTail(tmp741)

																				tmp743 := PrimTail(tmp742)

																				tmp744 := PrimIsPair(tmp743)

																				var ifres730 Obj

																				if True == tmp744 {
																					tmp732 := PrimTail(val125)

																					tmp733 := PrimHead(tmp732)

																					tmp734 := PrimTail(tmp733)

																					tmp735 := PrimTail(tmp734)

																					tmp736 := PrimTail(tmp735)

																					tmp737 := Call(__e, PrimNS1Value(symnull_2), tmp736)

																					tmp738 := PrimNot(tmp737)

																					var ifres731 Obj

																					if True == tmp738 {
																						ifres731 = True

																					} else {
																						ifres731 = False

																					}

																					ifres730 = ifres731

																				} else {
																					ifres730 = False

																				}

																				if True == ifres730 {
																					tmp703 := MakeNative(func(__e *ControlFlow) {
																						z := __e.Get(1)
																						_ = z
																						tmp717 := PrimTail(val125)

																						tmp718 := PrimHead(tmp717)

																						tmp719 := PrimTail(tmp718)

																						tmp720 := PrimTail(tmp719)

																						tmp721 := PrimTail(tmp720)

																						tmp722 := PrimTail(tmp721)

																						tmp723 := Call(__e, PrimNS1Value(symnull_2), tmp722)

																						if True == tmp723 {
																							tmp714 := PrimTail(val125)

																							tmp715 := PrimTail(tmp714)

																							tmp716 := Call(__e, PrimNS1Value(symnull_2), tmp715)

																							if True == tmp716 {
																								tmp706 := PrimCons(x, Nil)

																								tmp707 := PrimCons(x, env)

																								tmp708 := Call(__e, PrimNS1Value(symkl_1_6cora), tmp707, z)

																								tmp709 := PrimCons(tmp708, Nil)

																								tmp710 := PrimCons(tmp706, tmp709)

																								tmp711 := PrimCons(symlambda, tmp710)

																								tmp712 := Call(__e, PrimNS1Value(symkl_1_6cora), env, y)

																								tmp713 := PrimCons(tmp712, Nil)

																								__e.Return(PrimCons(tmp711, tmp713))
																								return

																							} else {
																								__e.TailApply(cc133)
																								return
																							}

																						} else {
																							__e.TailApply(cc133)
																							return
																						}

																					}, 1)

																					tmp724 := PrimTail(val125)

																					tmp725 := PrimHead(tmp724)

																					tmp726 := PrimTail(tmp725)

																					tmp727 := PrimTail(tmp726)

																					tmp728 := PrimTail(tmp727)

																					tmp729 := PrimHead(tmp728)

																					__e.TailApply(tmp703, tmp729)
																					return

																				} else {
																					__e.TailApply(cc133)
																					return
																				}

																			}, 1)

																			tmp745 := PrimTail(val125)

																			tmp746 := PrimHead(tmp745)

																			tmp747 := PrimTail(tmp746)

																			tmp748 := PrimTail(tmp747)

																			tmp749 := PrimHead(tmp748)

																			__e.TailApply(tmp701, tmp749)
																			return

																		} else {
																			__e.TailApply(cc133)
																			return
																		}

																	}, 1)

																	tmp763 := PrimTail(val125)

																	tmp764 := PrimHead(tmp763)

																	tmp765 := PrimTail(tmp764)

																	tmp766 := PrimHead(tmp765)

																	__e.TailApply(tmp699, tmp766)
																	return

																} else {
																	__e.TailApply(cc133)
																	return
																}

															} else {
																__e.TailApply(cc133)
																return
															}

														} else {
															__e.TailApply(cc133)
															return
														}

													} else {
														__e.TailApply(cc133)
														return
													}

												}, 1)

												tmp798 := PrimHead(val125)

												__e.TailApply(tmp694, tmp798)
												return

											} else {
												__e.TailApply(cc133)
												return
											}

										}, 1)

										tmp804 := MakeNative(func(__e *ControlFlow) {
											tmp805 := MakeNative(func(__e *ControlFlow) {
												cc134 := __e.Get(1)
												_ = cc134
												tmp891 := PrimIsPair(val125)

												var ifres887 Obj

												if True == tmp891 {
													tmp889 := Call(__e, PrimNS1Value(symnull_2), val125)

													tmp890 := PrimNot(tmp889)

													var ifres888 Obj

													if True == tmp890 {
														ifres888 = True

													} else {
														ifres888 = False

													}

													ifres887 = ifres888

												} else {
													ifres887 = False

												}

												if True == ifres887 {
													tmp807 := MakeNative(func(__e *ControlFlow) {
														env := __e.Get(1)
														_ = env
														tmp884 := PrimTail(val125)

														tmp885 := PrimIsPair(tmp884)

														var ifres879 Obj

														if True == tmp885 {
															tmp881 := PrimTail(val125)

															tmp882 := Call(__e, PrimNS1Value(symnull_2), tmp881)

															tmp883 := PrimNot(tmp882)

															var ifres880 Obj

															if True == tmp883 {
																ifres880 = True

															} else {
																ifres880 = False

															}

															ifres879 = ifres880

														} else {
															ifres879 = False

														}

														if True == ifres879 {
															tmp876 := PrimTail(val125)

															tmp877 := PrimHead(tmp876)

															tmp878 := PrimIsPair(tmp877)

															var ifres870 Obj

															if True == tmp878 {
																tmp872 := PrimTail(val125)

																tmp873 := PrimHead(tmp872)

																tmp874 := Call(__e, PrimNS1Value(symnull_2), tmp873)

																tmp875 := PrimNot(tmp874)

																var ifres871 Obj

																if True == tmp875 {
																	ifres871 = True

																} else {
																	ifres871 = False

																}

																ifres870 = ifres871

															} else {
																ifres870 = False

															}

															if True == ifres870 {
																tmp866 := PrimTail(val125)

																tmp867 := PrimHead(tmp866)

																tmp868 := PrimHead(tmp867)

																tmp869 := PrimEqual(symtrap_1error, tmp868)

																if True == tmp869 {
																	tmp862 := PrimTail(val125)

																	tmp863 := PrimHead(tmp862)

																	tmp864 := PrimTail(tmp863)

																	tmp865 := PrimIsPair(tmp864)

																	var ifres855 Obj

																	if True == tmp865 {
																		tmp857 := PrimTail(val125)

																		tmp858 := PrimHead(tmp857)

																		tmp859 := PrimTail(tmp858)

																		tmp860 := Call(__e, PrimNS1Value(symnull_2), tmp859)

																		tmp861 := PrimNot(tmp860)

																		var ifres856 Obj

																		if True == tmp861 {
																			ifres856 = True

																		} else {
																			ifres856 = False

																		}

																		ifres855 = ifres856

																	} else {
																		ifres855 = False

																	}

																	if True == ifres855 {
																		tmp812 := MakeNative(func(__e *ControlFlow) {
																			body := __e.Get(1)
																			_ = body
																			tmp846 := PrimTail(val125)

																			tmp847 := PrimHead(tmp846)

																			tmp848 := PrimTail(tmp847)

																			tmp849 := PrimTail(tmp848)

																			tmp850 := PrimIsPair(tmp849)

																			var ifres838 Obj

																			if True == tmp850 {
																				tmp840 := PrimTail(val125)

																				tmp841 := PrimHead(tmp840)

																				tmp842 := PrimTail(tmp841)

																				tmp843 := PrimTail(tmp842)

																				tmp844 := Call(__e, PrimNS1Value(symnull_2), tmp843)

																				tmp845 := PrimNot(tmp844)

																				var ifres839 Obj

																				if True == tmp845 {
																					ifres839 = True

																				} else {
																					ifres839 = False

																				}

																				ifres838 = ifres839

																			} else {
																				ifres838 = False

																			}

																			if True == ifres838 {
																				tmp814 := MakeNative(func(__e *ControlFlow) {
																					handler := __e.Get(1)
																					_ = handler
																					tmp827 := PrimTail(val125)

																					tmp828 := PrimHead(tmp827)

																					tmp829 := PrimTail(tmp828)

																					tmp830 := PrimTail(tmp829)

																					tmp831 := PrimTail(tmp830)

																					tmp832 := Call(__e, PrimNS1Value(symnull_2), tmp831)

																					if True == tmp832 {
																						tmp824 := PrimTail(val125)

																						tmp825 := PrimTail(tmp824)

																						tmp826 := Call(__e, PrimNS1Value(symnull_2), tmp825)

																						if True == tmp826 {
																							tmp817 := Call(__e, PrimNS1Value(symkl_1_6cora), env, body)

																							tmp818 := PrimCons(tmp817, Nil)

																							tmp819 := PrimCons(Nil, tmp818)

																							tmp820 := PrimCons(symlambda, tmp819)

																							tmp821 := Call(__e, PrimNS1Value(symkl_1_6cora), env, handler)

																							tmp822 := PrimCons(tmp821, Nil)

																							tmp823 := PrimCons(tmp820, tmp822)

																							__e.Return(PrimCons(symtry_1catch, tmp823))
																							return

																						} else {
																							__e.TailApply(cc134)
																							return
																						}

																					} else {
																						__e.TailApply(cc134)
																						return
																					}

																				}, 1)

																				tmp833 := PrimTail(val125)

																				tmp834 := PrimHead(tmp833)

																				tmp835 := PrimTail(tmp834)

																				tmp836 := PrimTail(tmp835)

																				tmp837 := PrimHead(tmp836)

																				__e.TailApply(tmp814, tmp837)
																				return

																			} else {
																				__e.TailApply(cc134)
																				return
																			}

																		}, 1)

																		tmp851 := PrimTail(val125)

																		tmp852 := PrimHead(tmp851)

																		tmp853 := PrimTail(tmp852)

																		tmp854 := PrimHead(tmp853)

																		__e.TailApply(tmp812, tmp854)
																		return

																	} else {
																		__e.TailApply(cc134)
																		return
																	}

																} else {
																	__e.TailApply(cc134)
																	return
																}

															} else {
																__e.TailApply(cc134)
																return
															}

														} else {
															__e.TailApply(cc134)
															return
														}

													}, 1)

													tmp886 := PrimHead(val125)

													__e.TailApply(tmp807, tmp886)
													return

												} else {
													__e.TailApply(cc134)
													return
												}

											}, 1)

											tmp892 := MakeNative(func(__e *ControlFlow) {
												tmp893 := MakeNative(func(__e *ControlFlow) {
													cc135 := __e.Get(1)
													_ = cc135
													tmp980 := PrimIsPair(val125)

													var ifres976 Obj

													if True == tmp980 {
														tmp978 := Call(__e, PrimNS1Value(symnull_2), val125)

														tmp979 := PrimNot(tmp978)

														var ifres977 Obj

														if True == tmp979 {
															ifres977 = True

														} else {
															ifres977 = False

														}

														ifres976 = ifres977

													} else {
														ifres976 = False

													}

													if True == ifres976 {
														tmp895 := MakeNative(func(__e *ControlFlow) {
															env := __e.Get(1)
															_ = env
															tmp973 := PrimTail(val125)

															tmp974 := PrimIsPair(tmp973)

															var ifres968 Obj

															if True == tmp974 {
																tmp970 := PrimTail(val125)

																tmp971 := Call(__e, PrimNS1Value(symnull_2), tmp970)

																tmp972 := PrimNot(tmp971)

																var ifres969 Obj

																if True == tmp972 {
																	ifres969 = True

																} else {
																	ifres969 = False

																}

																ifres968 = ifres969

															} else {
																ifres968 = False

															}

															if True == ifres968 {
																tmp965 := PrimTail(val125)

																tmp966 := PrimHead(tmp965)

																tmp967 := PrimIsPair(tmp966)

																var ifres959 Obj

																if True == tmp967 {
																	tmp961 := PrimTail(val125)

																	tmp962 := PrimHead(tmp961)

																	tmp963 := Call(__e, PrimNS1Value(symnull_2), tmp962)

																	tmp964 := PrimNot(tmp963)

																	var ifres960 Obj

																	if True == tmp964 {
																		ifres960 = True

																	} else {
																		ifres960 = False

																	}

																	ifres959 = ifres960

																} else {
																	ifres959 = False

																}

																if True == ifres959 {
																	tmp955 := PrimTail(val125)

																	tmp956 := PrimHead(tmp955)

																	tmp957 := PrimHead(tmp956)

																	tmp958 := PrimEqual(symor, tmp957)

																	if True == tmp958 {
																		tmp951 := PrimTail(val125)

																		tmp952 := PrimHead(tmp951)

																		tmp953 := PrimTail(tmp952)

																		tmp954 := PrimIsPair(tmp953)

																		var ifres944 Obj

																		if True == tmp954 {
																			tmp946 := PrimTail(val125)

																			tmp947 := PrimHead(tmp946)

																			tmp948 := PrimTail(tmp947)

																			tmp949 := Call(__e, PrimNS1Value(symnull_2), tmp948)

																			tmp950 := PrimNot(tmp949)

																			var ifres945 Obj

																			if True == tmp950 {
																				ifres945 = True

																			} else {
																				ifres945 = False

																			}

																			ifres944 = ifres945

																		} else {
																			ifres944 = False

																		}

																		if True == ifres944 {
																			tmp900 := MakeNative(func(__e *ControlFlow) {
																				x := __e.Get(1)
																				_ = x
																				tmp935 := PrimTail(val125)

																				tmp936 := PrimHead(tmp935)

																				tmp937 := PrimTail(tmp936)

																				tmp938 := PrimTail(tmp937)

																				tmp939 := PrimIsPair(tmp938)

																				var ifres927 Obj

																				if True == tmp939 {
																					tmp929 := PrimTail(val125)

																					tmp930 := PrimHead(tmp929)

																					tmp931 := PrimTail(tmp930)

																					tmp932 := PrimTail(tmp931)

																					tmp933 := Call(__e, PrimNS1Value(symnull_2), tmp932)

																					tmp934 := PrimNot(tmp933)

																					var ifres928 Obj

																					if True == tmp934 {
																						ifres928 = True

																					} else {
																						ifres928 = False

																					}

																					ifres927 = ifres928

																				} else {
																					ifres927 = False

																				}

																				if True == ifres927 {
																					tmp902 := MakeNative(func(__e *ControlFlow) {
																						y := __e.Get(1)
																						_ = y
																						tmp916 := PrimTail(val125)

																						tmp917 := PrimHead(tmp916)

																						tmp918 := PrimTail(tmp917)

																						tmp919 := PrimTail(tmp918)

																						tmp920 := PrimTail(tmp919)

																						tmp921 := Call(__e, PrimNS1Value(symnull_2), tmp920)

																						if True == tmp921 {
																							tmp913 := PrimTail(val125)

																							tmp914 := PrimTail(tmp913)

																							tmp915 := Call(__e, PrimNS1Value(symnull_2), tmp914)

																							if True == tmp915 {
																								tmp905 := PrimCons(False, Nil)

																								tmp906 := PrimCons(True, tmp905)

																								tmp907 := PrimCons(y, tmp906)

																								tmp908 := PrimCons(symif, tmp907)

																								tmp909 := PrimCons(tmp908, Nil)

																								tmp910 := PrimCons(True, tmp909)

																								tmp911 := PrimCons(x, tmp910)

																								tmp912 := PrimCons(symif, tmp911)

																								__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp912)
																								return

																							} else {
																								__e.TailApply(cc135)
																								return
																							}

																						} else {
																							__e.TailApply(cc135)
																							return
																						}

																					}, 1)

																					tmp922 := PrimTail(val125)

																					tmp923 := PrimHead(tmp922)

																					tmp924 := PrimTail(tmp923)

																					tmp925 := PrimTail(tmp924)

																					tmp926 := PrimHead(tmp925)

																					__e.TailApply(tmp902, tmp926)
																					return

																				} else {
																					__e.TailApply(cc135)
																					return
																				}

																			}, 1)

																			tmp940 := PrimTail(val125)

																			tmp941 := PrimHead(tmp940)

																			tmp942 := PrimTail(tmp941)

																			tmp943 := PrimHead(tmp942)

																			__e.TailApply(tmp900, tmp943)
																			return

																		} else {
																			__e.TailApply(cc135)
																			return
																		}

																	} else {
																		__e.TailApply(cc135)
																		return
																	}

																} else {
																	__e.TailApply(cc135)
																	return
																}

															} else {
																__e.TailApply(cc135)
																return
															}

														}, 1)

														tmp975 := PrimHead(val125)

														__e.TailApply(tmp895, tmp975)
														return

													} else {
														__e.TailApply(cc135)
														return
													}

												}, 1)

												tmp981 := MakeNative(func(__e *ControlFlow) {
													tmp982 := MakeNative(func(__e *ControlFlow) {
														cc136 := __e.Get(1)
														_ = cc136
														tmp1069 := PrimIsPair(val125)

														var ifres1065 Obj

														if True == tmp1069 {
															tmp1067 := Call(__e, PrimNS1Value(symnull_2), val125)

															tmp1068 := PrimNot(tmp1067)

															var ifres1066 Obj

															if True == tmp1068 {
																ifres1066 = True

															} else {
																ifres1066 = False

															}

															ifres1065 = ifres1066

														} else {
															ifres1065 = False

														}

														if True == ifres1065 {
															tmp984 := MakeNative(func(__e *ControlFlow) {
																env := __e.Get(1)
																_ = env
																tmp1062 := PrimTail(val125)

																tmp1063 := PrimIsPair(tmp1062)

																var ifres1057 Obj

																if True == tmp1063 {
																	tmp1059 := PrimTail(val125)

																	tmp1060 := Call(__e, PrimNS1Value(symnull_2), tmp1059)

																	tmp1061 := PrimNot(tmp1060)

																	var ifres1058 Obj

																	if True == tmp1061 {
																		ifres1058 = True

																	} else {
																		ifres1058 = False

																	}

																	ifres1057 = ifres1058

																} else {
																	ifres1057 = False

																}

																if True == ifres1057 {
																	tmp1054 := PrimTail(val125)

																	tmp1055 := PrimHead(tmp1054)

																	tmp1056 := PrimIsPair(tmp1055)

																	var ifres1048 Obj

																	if True == tmp1056 {
																		tmp1050 := PrimTail(val125)

																		tmp1051 := PrimHead(tmp1050)

																		tmp1052 := Call(__e, PrimNS1Value(symnull_2), tmp1051)

																		tmp1053 := PrimNot(tmp1052)

																		var ifres1049 Obj

																		if True == tmp1053 {
																			ifres1049 = True

																		} else {
																			ifres1049 = False

																		}

																		ifres1048 = ifres1049

																	} else {
																		ifres1048 = False

																	}

																	if True == ifres1048 {
																		tmp1044 := PrimTail(val125)

																		tmp1045 := PrimHead(tmp1044)

																		tmp1046 := PrimHead(tmp1045)

																		tmp1047 := PrimEqual(symand, tmp1046)

																		if True == tmp1047 {
																			tmp1040 := PrimTail(val125)

																			tmp1041 := PrimHead(tmp1040)

																			tmp1042 := PrimTail(tmp1041)

																			tmp1043 := PrimIsPair(tmp1042)

																			var ifres1033 Obj

																			if True == tmp1043 {
																				tmp1035 := PrimTail(val125)

																				tmp1036 := PrimHead(tmp1035)

																				tmp1037 := PrimTail(tmp1036)

																				tmp1038 := Call(__e, PrimNS1Value(symnull_2), tmp1037)

																				tmp1039 := PrimNot(tmp1038)

																				var ifres1034 Obj

																				if True == tmp1039 {
																					ifres1034 = True

																				} else {
																					ifres1034 = False

																				}

																				ifres1033 = ifres1034

																			} else {
																				ifres1033 = False

																			}

																			if True == ifres1033 {
																				tmp989 := MakeNative(func(__e *ControlFlow) {
																					x := __e.Get(1)
																					_ = x
																					tmp1024 := PrimTail(val125)

																					tmp1025 := PrimHead(tmp1024)

																					tmp1026 := PrimTail(tmp1025)

																					tmp1027 := PrimTail(tmp1026)

																					tmp1028 := PrimIsPair(tmp1027)

																					var ifres1016 Obj

																					if True == tmp1028 {
																						tmp1018 := PrimTail(val125)

																						tmp1019 := PrimHead(tmp1018)

																						tmp1020 := PrimTail(tmp1019)

																						tmp1021 := PrimTail(tmp1020)

																						tmp1022 := Call(__e, PrimNS1Value(symnull_2), tmp1021)

																						tmp1023 := PrimNot(tmp1022)

																						var ifres1017 Obj

																						if True == tmp1023 {
																							ifres1017 = True

																						} else {
																							ifres1017 = False

																						}

																						ifres1016 = ifres1017

																					} else {
																						ifres1016 = False

																					}

																					if True == ifres1016 {
																						tmp991 := MakeNative(func(__e *ControlFlow) {
																							y := __e.Get(1)
																							_ = y
																							tmp1005 := PrimTail(val125)

																							tmp1006 := PrimHead(tmp1005)

																							tmp1007 := PrimTail(tmp1006)

																							tmp1008 := PrimTail(tmp1007)

																							tmp1009 := PrimTail(tmp1008)

																							tmp1010 := Call(__e, PrimNS1Value(symnull_2), tmp1009)

																							if True == tmp1010 {
																								tmp1002 := PrimTail(val125)

																								tmp1003 := PrimTail(tmp1002)

																								tmp1004 := Call(__e, PrimNS1Value(symnull_2), tmp1003)

																								if True == tmp1004 {
																									tmp994 := PrimCons(False, Nil)

																									tmp995 := PrimCons(True, tmp994)

																									tmp996 := PrimCons(y, tmp995)

																									tmp997 := PrimCons(symif, tmp996)

																									tmp998 := PrimCons(False, Nil)

																									tmp999 := PrimCons(tmp997, tmp998)

																									tmp1000 := PrimCons(x, tmp999)

																									tmp1001 := PrimCons(symif, tmp1000)

																									__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp1001)
																									return

																								} else {
																									__e.TailApply(cc136)
																									return
																								}

																							} else {
																								__e.TailApply(cc136)
																								return
																							}

																						}, 1)

																						tmp1011 := PrimTail(val125)

																						tmp1012 := PrimHead(tmp1011)

																						tmp1013 := PrimTail(tmp1012)

																						tmp1014 := PrimTail(tmp1013)

																						tmp1015 := PrimHead(tmp1014)

																						__e.TailApply(tmp991, tmp1015)
																						return

																					} else {
																						__e.TailApply(cc136)
																						return
																					}

																				}, 1)

																				tmp1029 := PrimTail(val125)

																				tmp1030 := PrimHead(tmp1029)

																				tmp1031 := PrimTail(tmp1030)

																				tmp1032 := PrimHead(tmp1031)

																				__e.TailApply(tmp989, tmp1032)
																				return

																			} else {
																				__e.TailApply(cc136)
																				return
																			}

																		} else {
																			__e.TailApply(cc136)
																			return
																		}

																	} else {
																		__e.TailApply(cc136)
																		return
																	}

																} else {
																	__e.TailApply(cc136)
																	return
																}

															}, 1)

															tmp1064 := PrimHead(val125)

															__e.TailApply(tmp984, tmp1064)
															return

														} else {
															__e.TailApply(cc136)
															return
														}

													}, 1)

													tmp1070 := MakeNative(func(__e *ControlFlow) {
														tmp1071 := MakeNative(func(__e *ControlFlow) {
															cc137 := __e.Get(1)
															_ = cc137
															tmp1181 := PrimIsPair(val125)

															var ifres1177 Obj

															if True == tmp1181 {
																tmp1179 := Call(__e, PrimNS1Value(symnull_2), val125)

																tmp1180 := PrimNot(tmp1179)

																var ifres1178 Obj

																if True == tmp1180 {
																	ifres1178 = True

																} else {
																	ifres1178 = False

																}

																ifres1177 = ifres1178

															} else {
																ifres1177 = False

															}

															if True == ifres1177 {
																tmp1073 := MakeNative(func(__e *ControlFlow) {
																	env := __e.Get(1)
																	_ = env
																	tmp1174 := PrimTail(val125)

																	tmp1175 := PrimIsPair(tmp1174)

																	var ifres1169 Obj

																	if True == tmp1175 {
																		tmp1171 := PrimTail(val125)

																		tmp1172 := Call(__e, PrimNS1Value(symnull_2), tmp1171)

																		tmp1173 := PrimNot(tmp1172)

																		var ifres1170 Obj

																		if True == tmp1173 {
																			ifres1170 = True

																		} else {
																			ifres1170 = False

																		}

																		ifres1169 = ifres1170

																	} else {
																		ifres1169 = False

																	}

																	if True == ifres1169 {
																		tmp1166 := PrimTail(val125)

																		tmp1167 := PrimHead(tmp1166)

																		tmp1168 := PrimIsPair(tmp1167)

																		var ifres1160 Obj

																		if True == tmp1168 {
																			tmp1162 := PrimTail(val125)

																			tmp1163 := PrimHead(tmp1162)

																			tmp1164 := Call(__e, PrimNS1Value(symnull_2), tmp1163)

																			tmp1165 := PrimNot(tmp1164)

																			var ifres1161 Obj

																			if True == tmp1165 {
																				ifres1161 = True

																			} else {
																				ifres1161 = False

																			}

																			ifres1160 = ifres1161

																		} else {
																			ifres1160 = False

																		}

																		if True == ifres1160 {
																			tmp1156 := PrimTail(val125)

																			tmp1157 := PrimHead(tmp1156)

																			tmp1158 := PrimHead(tmp1157)

																			tmp1159 := PrimEqual(symcond, tmp1158)

																			if True == tmp1159 {
																				tmp1152 := PrimTail(val125)

																				tmp1153 := PrimHead(tmp1152)

																				tmp1154 := PrimTail(tmp1153)

																				tmp1155 := PrimIsPair(tmp1154)

																				var ifres1145 Obj

																				if True == tmp1155 {
																					tmp1147 := PrimTail(val125)

																					tmp1148 := PrimHead(tmp1147)

																					tmp1149 := PrimTail(tmp1148)

																					tmp1150 := Call(__e, PrimNS1Value(symnull_2), tmp1149)

																					tmp1151 := PrimNot(tmp1150)

																					var ifres1146 Obj

																					if True == tmp1151 {
																						ifres1146 = True

																					} else {
																						ifres1146 = False

																					}

																					ifres1145 = ifres1146

																				} else {
																					ifres1145 = False

																				}

																				if True == ifres1145 {
																					tmp1140 := PrimTail(val125)

																					tmp1141 := PrimHead(tmp1140)

																					tmp1142 := PrimTail(tmp1141)

																					tmp1143 := PrimHead(tmp1142)

																					tmp1144 := PrimIsPair(tmp1143)

																					var ifres1132 Obj

																					if True == tmp1144 {
																						tmp1134 := PrimTail(val125)

																						tmp1135 := PrimHead(tmp1134)

																						tmp1136 := PrimTail(tmp1135)

																						tmp1137 := PrimHead(tmp1136)

																						tmp1138 := Call(__e, PrimNS1Value(symnull_2), tmp1137)

																						tmp1139 := PrimNot(tmp1138)

																						var ifres1133 Obj

																						if True == tmp1139 {
																							ifres1133 = True

																						} else {
																							ifres1133 = False

																						}

																						ifres1132 = ifres1133

																					} else {
																						ifres1132 = False

																					}

																					if True == ifres1132 {
																						tmp1079 := MakeNative(func(__e *ControlFlow) {
																							__case := __e.Get(1)
																							_ = __case
																							tmp1121 := PrimTail(val125)

																							tmp1122 := PrimHead(tmp1121)

																							tmp1123 := PrimTail(tmp1122)

																							tmp1124 := PrimHead(tmp1123)

																							tmp1125 := PrimTail(tmp1124)

																							tmp1126 := PrimIsPair(tmp1125)

																							var ifres1112 Obj

																							if True == tmp1126 {
																								tmp1114 := PrimTail(val125)

																								tmp1115 := PrimHead(tmp1114)

																								tmp1116 := PrimTail(tmp1115)

																								tmp1117 := PrimHead(tmp1116)

																								tmp1118 := PrimTail(tmp1117)

																								tmp1119 := Call(__e, PrimNS1Value(symnull_2), tmp1118)

																								tmp1120 := PrimNot(tmp1119)

																								var ifres1113 Obj

																								if True == tmp1120 {
																									ifres1113 = True

																								} else {
																									ifres1113 = False

																								}

																								ifres1112 = ifres1113

																							} else {
																								ifres1112 = False

																							}

																							if True == ifres1112 {
																								tmp1081 := MakeNative(func(__e *ControlFlow) {
																									action := __e.Get(1)
																									_ = action
																									tmp1099 := PrimTail(val125)

																									tmp1100 := PrimHead(tmp1099)

																									tmp1101 := PrimTail(tmp1100)

																									tmp1102 := PrimHead(tmp1101)

																									tmp1103 := PrimTail(tmp1102)

																									tmp1104 := PrimTail(tmp1103)

																									tmp1105 := Call(__e, PrimNS1Value(symnull_2), tmp1104)

																									if True == tmp1105 {
																										tmp1094 := PrimTail(val125)

																										tmp1095 := PrimHead(tmp1094)

																										tmp1096 := PrimTail(tmp1095)

																										tmp1097 := PrimTail(tmp1096)

																										tmp1098 := Call(__e, PrimNS1Value(symnull_2), tmp1097)

																										if True == tmp1098 {
																											tmp1091 := PrimTail(val125)

																											tmp1092 := PrimTail(tmp1091)

																											tmp1093 := Call(__e, PrimNS1Value(symnull_2), tmp1092)

																											if True == tmp1093 {
																												tmp1085 := PrimCons(MakeString("no cond match"), Nil)

																												tmp1086 := PrimCons(symsimple_1error, tmp1085)

																												tmp1087 := PrimCons(tmp1086, Nil)

																												tmp1088 := PrimCons(action, tmp1087)

																												tmp1089 := PrimCons(__case, tmp1088)

																												tmp1090 := PrimCons(symif, tmp1089)

																												__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp1090)
																												return

																											} else {
																												__e.TailApply(cc137)
																												return
																											}

																										} else {
																											__e.TailApply(cc137)
																											return
																										}

																									} else {
																										__e.TailApply(cc137)
																										return
																									}

																								}, 1)

																								tmp1106 := PrimTail(val125)

																								tmp1107 := PrimHead(tmp1106)

																								tmp1108 := PrimTail(tmp1107)

																								tmp1109 := PrimHead(tmp1108)

																								tmp1110 := PrimTail(tmp1109)

																								tmp1111 := PrimHead(tmp1110)

																								__e.TailApply(tmp1081, tmp1111)
																								return

																							} else {
																								__e.TailApply(cc137)
																								return
																							}

																						}, 1)

																						tmp1127 := PrimTail(val125)

																						tmp1128 := PrimHead(tmp1127)

																						tmp1129 := PrimTail(tmp1128)

																						tmp1130 := PrimHead(tmp1129)

																						tmp1131 := PrimHead(tmp1130)

																						__e.TailApply(tmp1079, tmp1131)
																						return

																					} else {
																						__e.TailApply(cc137)
																						return
																					}

																				} else {
																					__e.TailApply(cc137)
																					return
																				}

																			} else {
																				__e.TailApply(cc137)
																				return
																			}

																		} else {
																			__e.TailApply(cc137)
																			return
																		}

																	} else {
																		__e.TailApply(cc137)
																		return
																	}

																}, 1)

																tmp1176 := PrimHead(val125)

																__e.TailApply(tmp1073, tmp1176)
																return

															} else {
																__e.TailApply(cc137)
																return
															}

														}, 1)

														tmp1182 := MakeNative(func(__e *ControlFlow) {
															tmp1183 := MakeNative(func(__e *ControlFlow) {
																cc138 := __e.Get(1)
																_ = cc138
																tmp1252 := PrimIsPair(val125)

																var ifres1248 Obj

																if True == tmp1252 {
																	tmp1250 := Call(__e, PrimNS1Value(symnull_2), val125)

																	tmp1251 := PrimNot(tmp1250)

																	var ifres1249 Obj

																	if True == tmp1251 {
																		ifres1249 = True

																	} else {
																		ifres1249 = False

																	}

																	ifres1248 = ifres1249

																} else {
																	ifres1248 = False

																}

																if True == ifres1248 {
																	tmp1185 := MakeNative(func(__e *ControlFlow) {
																		env := __e.Get(1)
																		_ = env
																		tmp1245 := PrimTail(val125)

																		tmp1246 := PrimIsPair(tmp1245)

																		var ifres1240 Obj

																		if True == tmp1246 {
																			tmp1242 := PrimTail(val125)

																			tmp1243 := Call(__e, PrimNS1Value(symnull_2), tmp1242)

																			tmp1244 := PrimNot(tmp1243)

																			var ifres1241 Obj

																			if True == tmp1244 {
																				ifres1241 = True

																			} else {
																				ifres1241 = False

																			}

																			ifres1240 = ifres1241

																		} else {
																			ifres1240 = False

																		}

																		if True == ifres1240 {
																			tmp1187 := MakeNative(func(__e *ControlFlow) {
																				n145 := __e.Get(1)
																				_ = n145
																				tmp1237 := PrimIsPair(n145)

																				if True == tmp1237 {
																					tmp1235 := PrimHead(n145)

																					tmp1236 := PrimEqual(symcond, tmp1235)

																					if True == tmp1236 {
																						tmp1190 := MakeNative(func(__e *ControlFlow) {
																							n146 := __e.Get(1)
																							_ = n146
																							tmp1233 := PrimIsPair(n146)

																							if True == tmp1233 {
																								tmp1231 := PrimHead(n146)

																								tmp1232 := PrimIsPair(tmp1231)

																								var ifres1226 Obj

																								if True == tmp1232 {
																									tmp1228 := PrimHead(n146)

																									tmp1229 := Call(__e, PrimNS1Value(symnull_2), tmp1228)

																									tmp1230 := PrimNot(tmp1229)

																									var ifres1227 Obj

																									if True == tmp1230 {
																										ifres1227 = True

																									} else {
																										ifres1227 = False

																									}

																									ifres1226 = ifres1227

																								} else {
																									ifres1226 = False

																								}

																								if True == ifres1226 {
																									tmp1193 := MakeNative(func(__e *ControlFlow) {
																										__case := __e.Get(1)
																										_ = __case
																										tmp1221 := PrimHead(n146)

																										tmp1222 := PrimTail(tmp1221)

																										tmp1223 := PrimIsPair(tmp1222)

																										var ifres1215 Obj

																										if True == tmp1223 {
																											tmp1217 := PrimHead(n146)

																											tmp1218 := PrimTail(tmp1217)

																											tmp1219 := Call(__e, PrimNS1Value(symnull_2), tmp1218)

																											tmp1220 := PrimNot(tmp1219)

																											var ifres1216 Obj

																											if True == tmp1220 {
																												ifres1216 = True

																											} else {
																												ifres1216 = False

																											}

																											ifres1215 = ifres1216

																										} else {
																											ifres1215 = False

																										}

																										if True == ifres1215 {
																											tmp1195 := MakeNative(func(__e *ControlFlow) {
																												action := __e.Get(1)
																												_ = action
																												tmp1208 := PrimHead(n146)

																												tmp1209 := PrimTail(tmp1208)

																												tmp1210 := PrimTail(tmp1209)

																												tmp1211 := Call(__e, PrimNS1Value(symnull_2), tmp1210)

																												if True == tmp1211 {
																													tmp1197 := MakeNative(func(__e *ControlFlow) {
																														more := __e.Get(1)
																														_ = more
																														tmp1204 := PrimTail(val125)

																														tmp1205 := PrimTail(tmp1204)

																														tmp1206 := Call(__e, PrimNS1Value(symnull_2), tmp1205)

																														if True == tmp1206 {
																															tmp1199 := PrimCons(symcond, more)

																															tmp1200 := PrimCons(tmp1199, Nil)

																															tmp1201 := PrimCons(action, tmp1200)

																															tmp1202 := PrimCons(__case, tmp1201)

																															tmp1203 := PrimCons(symif, tmp1202)

																															__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp1203)
																															return

																														} else {
																															__e.TailApply(cc138)
																															return
																														}

																													}, 1)

																													tmp1207 := PrimTail(n146)

																													__e.TailApply(tmp1197, tmp1207)
																													return

																												} else {
																													__e.TailApply(cc138)
																													return
																												}

																											}, 1)

																											tmp1212 := PrimHead(n146)

																											tmp1213 := PrimTail(tmp1212)

																											tmp1214 := PrimHead(tmp1213)

																											__e.TailApply(tmp1195, tmp1214)
																											return

																										} else {
																											__e.TailApply(cc138)
																											return
																										}

																									}, 1)

																									tmp1224 := PrimHead(n146)

																									tmp1225 := PrimHead(tmp1224)

																									__e.TailApply(tmp1193, tmp1225)
																									return

																								} else {
																									__e.TailApply(cc138)
																									return
																								}

																							} else {
																								__e.TailApply(cc138)
																								return
																							}

																						}, 1)

																						tmp1234 := PrimTail(n145)

																						__e.TailApply(tmp1190, tmp1234)
																						return

																					} else {
																						__e.TailApply(cc138)
																						return
																					}

																				} else {
																					__e.TailApply(cc138)
																					return
																				}

																			}, 1)

																			tmp1238 := PrimTail(val125)

																			tmp1239 := PrimHead(tmp1238)

																			__e.TailApply(tmp1187, tmp1239)
																			return

																		} else {
																			__e.TailApply(cc138)
																			return
																		}

																	}, 1)

																	tmp1247 := PrimHead(val125)

																	__e.TailApply(tmp1185, tmp1247)
																	return

																} else {
																	__e.TailApply(cc138)
																	return
																}

															}, 1)

															tmp1253 := MakeNative(func(__e *ControlFlow) {
																tmp1254 := MakeNative(func(__e *ControlFlow) {
																	cc139 := __e.Get(1)
																	_ = cc139
																	tmp1315 := PrimIsPair(val125)

																	var ifres1311 Obj

																	if True == tmp1315 {
																		tmp1313 := Call(__e, PrimNS1Value(symnull_2), val125)

																		tmp1314 := PrimNot(tmp1313)

																		var ifres1312 Obj

																		if True == tmp1314 {
																			ifres1312 = True

																		} else {
																			ifres1312 = False

																		}

																		ifres1311 = ifres1312

																	} else {
																		ifres1311 = False

																	}

																	if True == ifres1311 {
																		tmp1256 := MakeNative(func(__e *ControlFlow) {
																			env := __e.Get(1)
																			_ = env
																			tmp1308 := PrimTail(val125)

																			tmp1309 := PrimIsPair(tmp1308)

																			var ifres1303 Obj

																			if True == tmp1309 {
																				tmp1305 := PrimTail(val125)

																				tmp1306 := Call(__e, PrimNS1Value(symnull_2), tmp1305)

																				tmp1307 := PrimNot(tmp1306)

																				var ifres1304 Obj

																				if True == tmp1307 {
																					ifres1304 = True

																				} else {
																					ifres1304 = False

																				}

																				ifres1303 = ifres1304

																			} else {
																				ifres1303 = False

																			}

																			if True == ifres1303 {
																				tmp1300 := PrimTail(val125)

																				tmp1301 := PrimHead(tmp1300)

																				tmp1302 := PrimIsPair(tmp1301)

																				var ifres1294 Obj

																				if True == tmp1302 {
																					tmp1296 := PrimTail(val125)

																					tmp1297 := PrimHead(tmp1296)

																					tmp1298 := Call(__e, PrimNS1Value(symnull_2), tmp1297)

																					tmp1299 := PrimNot(tmp1298)

																					var ifres1295 Obj

																					if True == tmp1299 {
																						ifres1295 = True

																					} else {
																						ifres1295 = False

																					}

																					ifres1294 = ifres1295

																				} else {
																					ifres1294 = False

																				}

																				if True == ifres1294 {
																					tmp1290 := PrimTail(val125)

																					tmp1291 := PrimHead(tmp1290)

																					tmp1292 := PrimHead(tmp1291)

																					tmp1293 := PrimEqual(symfreeze, tmp1292)

																					if True == tmp1293 {
																						tmp1286 := PrimTail(val125)

																						tmp1287 := PrimHead(tmp1286)

																						tmp1288 := PrimTail(tmp1287)

																						tmp1289 := PrimIsPair(tmp1288)

																						var ifres1279 Obj

																						if True == tmp1289 {
																							tmp1281 := PrimTail(val125)

																							tmp1282 := PrimHead(tmp1281)

																							tmp1283 := PrimTail(tmp1282)

																							tmp1284 := Call(__e, PrimNS1Value(symnull_2), tmp1283)

																							tmp1285 := PrimNot(tmp1284)

																							var ifres1280 Obj

																							if True == tmp1285 {
																								ifres1280 = True

																							} else {
																								ifres1280 = False

																							}

																							ifres1279 = ifres1280

																						} else {
																							ifres1279 = False

																						}

																						if True == ifres1279 {
																							tmp1261 := MakeNative(func(__e *ControlFlow) {
																								body := __e.Get(1)
																								_ = body
																								tmp1270 := PrimTail(val125)

																								tmp1271 := PrimHead(tmp1270)

																								tmp1272 := PrimTail(tmp1271)

																								tmp1273 := PrimTail(tmp1272)

																								tmp1274 := Call(__e, PrimNS1Value(symnull_2), tmp1273)

																								if True == tmp1274 {
																									tmp1267 := PrimTail(val125)

																									tmp1268 := PrimTail(tmp1267)

																									tmp1269 := Call(__e, PrimNS1Value(symnull_2), tmp1268)

																									if True == tmp1269 {
																										tmp1264 := Call(__e, PrimNS1Value(symkl_1_6cora), env, body)

																										tmp1265 := PrimCons(tmp1264, Nil)

																										tmp1266 := PrimCons(Nil, tmp1265)

																										__e.Return(PrimCons(symlambda, tmp1266))
																										return

																									} else {
																										__e.TailApply(cc139)
																										return
																									}

																								} else {
																									__e.TailApply(cc139)
																									return
																								}

																							}, 1)

																							tmp1275 := PrimTail(val125)

																							tmp1276 := PrimHead(tmp1275)

																							tmp1277 := PrimTail(tmp1276)

																							tmp1278 := PrimHead(tmp1277)

																							__e.TailApply(tmp1261, tmp1278)
																							return

																						} else {
																							__e.TailApply(cc139)
																							return
																						}

																					} else {
																						__e.TailApply(cc139)
																						return
																					}

																				} else {
																					__e.TailApply(cc139)
																					return
																				}

																			} else {
																				__e.TailApply(cc139)
																				return
																			}

																		}, 1)

																		tmp1310 := PrimHead(val125)

																		__e.TailApply(tmp1256, tmp1310)
																		return

																	} else {
																		__e.TailApply(cc139)
																		return
																	}

																}, 1)

																tmp1316 := MakeNative(func(__e *ControlFlow) {
																	tmp1317 := MakeNative(func(__e *ControlFlow) {
																		cc140 := __e.Get(1)
																		_ = cc140
																		tmp1351 := PrimIsPair(val125)

																		var ifres1347 Obj

																		if True == tmp1351 {
																			tmp1349 := Call(__e, PrimNS1Value(symnull_2), val125)

																			tmp1350 := PrimNot(tmp1349)

																			var ifres1348 Obj

																			if True == tmp1350 {
																				ifres1348 = True

																			} else {
																				ifres1348 = False

																			}

																			ifres1347 = ifres1348

																		} else {
																			ifres1347 = False

																		}

																		if True == ifres1347 {
																			tmp1319 := MakeNative(func(__e *ControlFlow) {
																				env := __e.Get(1)
																				_ = env
																				tmp1344 := PrimTail(val125)

																				tmp1345 := PrimIsPair(tmp1344)

																				var ifres1339 Obj

																				if True == tmp1345 {
																					tmp1341 := PrimTail(val125)

																					tmp1342 := Call(__e, PrimNS1Value(symnull_2), tmp1341)

																					tmp1343 := PrimNot(tmp1342)

																					var ifres1340 Obj

																					if True == tmp1343 {
																						ifres1340 = True

																					} else {
																						ifres1340 = False

																					}

																					ifres1339 = ifres1340

																				} else {
																					ifres1339 = False

																				}

																				if True == ifres1339 {
																					tmp1321 := MakeNative(func(__e *ControlFlow) {
																						n144 := __e.Get(1)
																						_ = n144
																						tmp1336 := PrimIsPair(n144)

																						if True == tmp1336 {
																							tmp1323 := MakeNative(func(__e *ControlFlow) {
																								f := __e.Get(1)
																								_ = f
																								tmp1324 := MakeNative(func(__e *ControlFlow) {
																									x := __e.Get(1)
																									_ = x
																									tmp1331 := PrimTail(val125)

																									tmp1332 := PrimTail(tmp1331)

																									tmp1333 := Call(__e, PrimNS1Value(symnull_2), tmp1332)

																									if True == tmp1333 {
																										tmp1330 := PrimIsSymbol(f)

																										if True == tmp1330 {
																											tmp1327 := Call(__e, PrimNS1Value(symparse_1app), env, f)

																											tmp1328 := Call(__e, PrimNS1Value(symkl_1_6cora), env)

																											tmp1329 := Call(__e, PrimNS1Value(symmap), tmp1328, x)

																											__e.Return(PrimCons(tmp1327, tmp1329))
																											return

																										} else {
																											__e.TailApply(cc140)
																											return
																										}

																									} else {
																										__e.TailApply(cc140)
																										return
																									}

																								}, 1)

																								tmp1334 := PrimTail(n144)

																								__e.TailApply(tmp1324, tmp1334)
																								return

																							}, 1)

																							tmp1335 := PrimHead(n144)

																							__e.TailApply(tmp1323, tmp1335)
																							return

																						} else {
																							__e.TailApply(cc140)
																							return
																						}

																					}, 1)

																					tmp1337 := PrimTail(val125)

																					tmp1338 := PrimHead(tmp1337)

																					__e.TailApply(tmp1321, tmp1338)
																					return

																				} else {
																					__e.TailApply(cc140)
																					return
																				}

																			}, 1)

																			tmp1346 := PrimHead(val125)

																			__e.TailApply(tmp1319, tmp1346)
																			return

																		} else {
																			__e.TailApply(cc140)
																			return
																		}

																	}, 1)

																	tmp1352 := MakeNative(func(__e *ControlFlow) {
																		tmp1353 := MakeNative(func(__e *ControlFlow) {
																			cc141 := __e.Get(1)
																			_ = cc141
																			tmp1384 := PrimIsPair(val125)

																			var ifres1380 Obj

																			if True == tmp1384 {
																				tmp1382 := Call(__e, PrimNS1Value(symnull_2), val125)

																				tmp1383 := PrimNot(tmp1382)

																				var ifres1381 Obj

																				if True == tmp1383 {
																					ifres1381 = True

																				} else {
																					ifres1381 = False

																				}

																				ifres1380 = ifres1381

																			} else {
																				ifres1380 = False

																			}

																			if True == ifres1380 {
																				tmp1355 := MakeNative(func(__e *ControlFlow) {
																					env := __e.Get(1)
																					_ = env
																					tmp1377 := PrimTail(val125)

																					tmp1378 := PrimIsPair(tmp1377)

																					var ifres1372 Obj

																					if True == tmp1378 {
																						tmp1374 := PrimTail(val125)

																						tmp1375 := Call(__e, PrimNS1Value(symnull_2), tmp1374)

																						tmp1376 := PrimNot(tmp1375)

																						var ifres1373 Obj

																						if True == tmp1376 {
																							ifres1373 = True

																						} else {
																							ifres1373 = False

																						}

																						ifres1372 = ifres1373

																					} else {
																						ifres1372 = False

																					}

																					if True == ifres1372 {
																						tmp1357 := MakeNative(func(__e *ControlFlow) {
																							n143 := __e.Get(1)
																							_ = n143
																							tmp1369 := PrimIsPair(n143)

																							if True == tmp1369 {
																								tmp1359 := MakeNative(func(__e *ControlFlow) {
																									f := __e.Get(1)
																									_ = f
																									tmp1360 := MakeNative(func(__e *ControlFlow) {
																										x := __e.Get(1)
																										_ = x
																										tmp1364 := PrimTail(val125)

																										tmp1365 := PrimTail(tmp1364)

																										tmp1366 := Call(__e, PrimNS1Value(symnull_2), tmp1365)

																										if True == tmp1366 {
																											tmp1362 := Call(__e, PrimNS1Value(symkl_1_6cora), env)

																											tmp1363 := PrimCons(f, x)

																											__e.TailApply(PrimNS1Value(symmap), tmp1362, tmp1363)
																											return

																										} else {
																											__e.TailApply(cc141)
																											return
																										}

																									}, 1)

																									tmp1367 := PrimTail(n143)

																									__e.TailApply(tmp1360, tmp1367)
																									return

																								}, 1)

																								tmp1368 := PrimHead(n143)

																								__e.TailApply(tmp1359, tmp1368)
																								return

																							} else {
																								__e.TailApply(cc141)
																								return
																							}

																						}, 1)

																						tmp1370 := PrimTail(val125)

																						tmp1371 := PrimHead(tmp1370)

																						__e.TailApply(tmp1357, tmp1371)
																						return

																					} else {
																						__e.TailApply(cc141)
																						return
																					}

																				}, 1)

																				tmp1379 := PrimHead(val125)

																				__e.TailApply(tmp1355, tmp1379)
																				return

																			} else {
																				__e.TailApply(cc141)
																				return
																			}

																		}, 1)

																		tmp1385 := MakeNative(func(__e *ControlFlow) {
																			tmp1386 := MakeNative(func(__e *ControlFlow) {
																				cc142 := __e.Get(1)
																				_ = cc142
																				tmp1412 := PrimIsPair(val125)

																				var ifres1408 Obj

																				if True == tmp1412 {
																					tmp1410 := Call(__e, PrimNS1Value(symnull_2), val125)

																					tmp1411 := PrimNot(tmp1410)

																					var ifres1409 Obj

																					if True == tmp1411 {
																						ifres1409 = True

																					} else {
																						ifres1409 = False

																					}

																					ifres1408 = ifres1409

																				} else {
																					ifres1408 = False

																				}

																				if True == ifres1408 {
																					tmp1388 := MakeNative(func(__e *ControlFlow) {
																						env := __e.Get(1)
																						_ = env
																						tmp1405 := PrimTail(val125)

																						tmp1406 := PrimIsPair(tmp1405)

																						var ifres1400 Obj

																						if True == tmp1406 {
																							tmp1402 := PrimTail(val125)

																							tmp1403 := Call(__e, PrimNS1Value(symnull_2), tmp1402)

																							tmp1404 := PrimNot(tmp1403)

																							var ifres1401 Obj

																							if True == tmp1404 {
																								ifres1401 = True

																							} else {
																								ifres1401 = False

																							}

																							ifres1400 = ifres1401

																						} else {
																							ifres1400 = False

																						}

																						if True == ifres1400 {
																							tmp1390 := MakeNative(func(__e *ControlFlow) {
																								x := __e.Get(1)
																								_ = x
																								tmp1395 := PrimTail(val125)

																								tmp1396 := PrimTail(tmp1395)

																								tmp1397 := Call(__e, PrimNS1Value(symnull_2), tmp1396)

																								if True == tmp1397 {
																									tmp1394 := Call(__e, PrimNS1Value(symelem_2), x, env)

																									if True == tmp1394 {
																										__e.Return(x)
																										return
																									} else {
																										tmp1393 := PrimCons(x, Nil)

																										__e.Return(PrimCons(symquote, tmp1393))
																										return

																									}

																								} else {
																									__e.TailApply(cc142)
																									return
																								}

																							}, 1)

																							tmp1398 := PrimTail(val125)

																							tmp1399 := PrimHead(tmp1398)

																							__e.TailApply(tmp1390, tmp1399)
																							return

																						} else {
																							__e.TailApply(cc142)
																							return
																						}

																					}, 1)

																					tmp1407 := PrimHead(val125)

																					__e.TailApply(tmp1388, tmp1407)
																					return

																				} else {
																					__e.TailApply(cc142)
																					return
																				}

																			}, 1)

																			tmp1413 := MakeNative(func(__e *ControlFlow) {
																				__e.TailApply(PrimNS1Value(symerror), MakeString("no match-help found!"))
																				return
																			}, 0)

																			__e.TailApply(tmp1386, tmp1413)
																			return

																		}, 0)

																		__e.TailApply(tmp1353, tmp1385)
																		return

																	}, 0)

																	__e.TailApply(tmp1317, tmp1352)
																	return

																}, 0)

																__e.TailApply(tmp1254, tmp1316)
																return

															}, 0)

															__e.TailApply(tmp1183, tmp1253)
															return

														}, 0)

														__e.TailApply(tmp1071, tmp1182)
														return

													}, 0)

													__e.TailApply(tmp982, tmp1070)
													return

												}, 0)

												__e.TailApply(tmp893, tmp981)
												return

											}, 0)

											__e.TailApply(tmp805, tmp892)
											return

										}, 0)

										__e.TailApply(tmp692, tmp804)
										return

									}, 0)

									__e.TailApply(tmp575, tmp691)
									return

								}, 0)

								__e.TailApply(tmp494, tmp574)
								return

							}, 0)

							__e.TailApply(tmp408, tmp493)
							return

						}, 0)

						__e.TailApply(tmp323, tmp407)
						return

					}, 0)

					__e.TailApply(tmp212, tmp322)
					return

				}, 0)

				__e.TailApply(tmp180, tmp211)
				return

			}, 0)

			__e.TailApply(tmp154, tmp179)
			return

		}, 1)

		tmp1414 := PrimCons(p124, Nil)

		tmp1415 := PrimCons(p123, tmp1414)

		__e.TailApply(tmp153, tmp1415)
		return

	}, 2)

	tmp1416 := PrimNS1Set(symkl_1_6cora, tmp152)

	_ = tmp1416

	tmp1417 := MakeNative(func(__e *ControlFlow) {
		p147 := __e.Get(1)
		_ = p147
		p148 := __e.Get(2)
		_ = p148
		tmp1418 := MakeNative(func(__e *ControlFlow) {
			val149 := __e.Get(1)
			_ = val149
			tmp1419 := MakeNative(func(__e *ControlFlow) {
				cc150 := __e.Get(1)
				_ = cc150
				tmp1444 := PrimIsPair(val149)

				var ifres1440 Obj

				if True == tmp1444 {
					tmp1442 := Call(__e, PrimNS1Value(symnull_2), val149)

					tmp1443 := PrimNot(tmp1442)

					var ifres1441 Obj

					if True == tmp1443 {
						ifres1441 = True

					} else {
						ifres1441 = False

					}

					ifres1440 = ifres1441

				} else {
					ifres1440 = False

				}

				if True == ifres1440 {
					tmp1421 := MakeNative(func(__e *ControlFlow) {
						env := __e.Get(1)
						_ = env
						tmp1437 := PrimTail(val149)

						tmp1438 := PrimIsPair(tmp1437)

						var ifres1432 Obj

						if True == tmp1438 {
							tmp1434 := PrimTail(val149)

							tmp1435 := Call(__e, PrimNS1Value(symnull_2), tmp1434)

							tmp1436 := PrimNot(tmp1435)

							var ifres1433 Obj

							if True == tmp1436 {
								ifres1433 = True

							} else {
								ifres1433 = False

							}

							ifres1432 = ifres1433

						} else {
							ifres1432 = False

						}

						if True == ifres1432 {
							tmp1423 := MakeNative(func(__e *ControlFlow) {
								f := __e.Get(1)
								_ = f
								tmp1427 := PrimTail(val149)

								tmp1428 := PrimTail(tmp1427)

								tmp1429 := Call(__e, PrimNS1Value(symnull_2), tmp1428)

								if True == tmp1429 {
									tmp1426 := Call(__e, PrimNS1Value(symelem_2), f, env)

									if True == tmp1426 {
										__e.Return(f)
										return
									} else {
										__e.TailApply(cc150)
										return
									}

								} else {
									__e.TailApply(cc150)
									return
								}

							}, 1)

							tmp1430 := PrimTail(val149)

							tmp1431 := PrimHead(tmp1430)

							__e.TailApply(tmp1423, tmp1431)
							return

						} else {
							__e.TailApply(cc150)
							return
						}

					}, 1)

					tmp1439 := PrimHead(val149)

					__e.TailApply(tmp1421, tmp1439)
					return

				} else {
					__e.TailApply(cc150)
					return
				}

			}, 1)

			tmp1445 := MakeNative(func(__e *ControlFlow) {
				tmp1446 := MakeNative(func(__e *ControlFlow) {
					cc151 := __e.Get(1)
					_ = cc151
					tmp1472 := PrimIsPair(val149)

					var ifres1468 Obj

					if True == tmp1472 {
						tmp1470 := Call(__e, PrimNS1Value(symnull_2), val149)

						tmp1471 := PrimNot(tmp1470)

						var ifres1469 Obj

						if True == tmp1471 {
							ifres1469 = True

						} else {
							ifres1469 = False

						}

						ifres1468 = ifres1469

					} else {
						ifres1468 = False

					}

					if True == ifres1468 {
						tmp1448 := MakeNative(func(__e *ControlFlow) {
							__ := __e.Get(1)
							_ = __
							tmp1465 := PrimTail(val149)

							tmp1466 := PrimIsPair(tmp1465)

							var ifres1460 Obj

							if True == tmp1466 {
								tmp1462 := PrimTail(val149)

								tmp1463 := Call(__e, PrimNS1Value(symnull_2), tmp1462)

								tmp1464 := PrimNot(tmp1463)

								var ifres1461 Obj

								if True == tmp1464 {
									ifres1461 = True

								} else {
									ifres1461 = False

								}

								ifres1460 = ifres1461

							} else {
								ifres1460 = False

							}

							if True == ifres1460 {
								tmp1450 := MakeNative(func(__e *ControlFlow) {
									f := __e.Get(1)
									_ = f
									tmp1455 := PrimTail(val149)

									tmp1456 := PrimTail(tmp1455)

									tmp1457 := Call(__e, PrimNS1Value(symnull_2), tmp1456)

									if True == tmp1457 {
										tmp1452 := PrimCons(f, Nil)

										tmp1453 := PrimCons(symquote, tmp1452)

										tmp1454 := PrimCons(tmp1453, Nil)

										__e.Return(PrimCons(symfn, tmp1454))
										return

									} else {
										__e.TailApply(cc151)
										return
									}

								}, 1)

								tmp1458 := PrimTail(val149)

								tmp1459 := PrimHead(tmp1458)

								__e.TailApply(tmp1450, tmp1459)
								return

							} else {
								__e.TailApply(cc151)
								return
							}

						}, 1)

						tmp1467 := PrimHead(val149)

						__e.TailApply(tmp1448, tmp1467)
						return

					} else {
						__e.TailApply(cc151)
						return
					}

				}, 1)

				tmp1473 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimNS1Value(symerror), MakeString("no match-help found!"))
					return
				}, 0)

				__e.TailApply(tmp1446, tmp1473)
				return

			}, 0)

			__e.TailApply(tmp1419, tmp1445)
			return

		}, 1)

		tmp1474 := PrimCons(p148, Nil)

		tmp1475 := PrimCons(p147, tmp1474)

		__e.TailApply(tmp1418, tmp1475)
		return

	}, 2)

	tmp1476 := PrimNS1Set(symparse_1app, tmp1417)

	_ = tmp1476

	tmp1477 := MakeNative(func(__e *ControlFlow) {
		fin := __e.Get(1)
		_ = fin
		fout := __e.Get(2)
		_ = fout
		tmp1478 := MakeNative(func(__e *ControlFlow) {
			exprs := __e.Get(1)
			_ = exprs
			tmp1479 := MakeNative(func(__e *ControlFlow) {
				res := __e.Get(1)
				_ = res
				__e.TailApply(PrimNS1Value(symwrite_1sexp_1to_1file), fout, res)
				return
			}, 1)

			tmp1480 := Call(__e, PrimNS1Value(symkl_1_6cora), Nil)

			tmp1481 := Call(__e, PrimNS1Value(symmap), tmp1480, exprs)

			tmp1482 := PrimCons(symbegin, tmp1481)

			__e.TailApply(tmp1479, tmp1482)
			return

		}, 1)

		tmp1483 := Call(__e, PrimNS1Value(symread_1file_1as_1sexp), fin, False)

		__e.TailApply(tmp1478, tmp1483)
		return

	}, 2)

	tmp1484 := PrimNS1Set(symcompile_1file, tmp1477)

	_ = tmp1484

	tmp1485 := PrimNS3Set(sym_dlanguage_d, MakeString("Go"))

	_ = tmp1485

	tmp1486 := PrimNS3Set(sym_dimplementation_d, MakeString("AOT+interpreter"))

	_ = tmp1486

	tmp1487 := PrimNS3Set(sym_dporters_d, MakeString("Arthur Mao"))

	_ = tmp1487

	tmp1488 := PrimNS3Set(sym_dport_d, MakeString("1.0.0"))

	_ = tmp1488

	tmp1489 := Call(__e, PrimNS2Value(symdef), symcons, PrimNS1Value(symcons))

	_ = tmp1489

	tmp1490 := Call(__e, PrimNS2Value(symdef), symcons_2, PrimNS1Value(symcons_2))

	_ = tmp1490

	tmp1491 := Call(__e, PrimNS2Value(symdef), symhd, PrimNS1Value(symcar))

	_ = tmp1491

	tmp1492 := Call(__e, PrimNS2Value(symdef), symtl, PrimNS1Value(symcdr))

	_ = tmp1492

	tmp1493 := Call(__e, PrimNS2Value(symdef), symintern, PrimNS1Value(symintern))

	_ = tmp1493

	tmp1494 := Call(__e, PrimNS2Value(symdef), sym_a, PrimNS1Value(sym_a))

	_ = tmp1494

	tmp1495 := Call(__e, PrimNS2Value(symdef), sym_7, PrimNS1Value(sym_7))

	_ = tmp1495

	tmp1496 := Call(__e, PrimNS2Value(symdef), sym_1, PrimNS1Value(sym_1))

	_ = tmp1496

	tmp1497 := Call(__e, PrimNS2Value(symdef), sym_d, PrimNS1Value(sym_d))

	_ = tmp1497

	tmp1498 := Call(__e, PrimNS2Value(symdef), sym_c, PrimNS1Value(sym_c))

	_ = tmp1498

	tmp1499 := Call(__e, PrimNS2Value(symdef), sym_5_a, PrimNS1Value(sym_5_a))

	_ = tmp1499

	tmp1500 := Call(__e, PrimNS2Value(symdef), sym_6_a, PrimNS1Value(sym_6_a))

	_ = tmp1500

	tmp1501 := Call(__e, PrimNS2Value(symdef), sym_5, PrimNS1Value(sym_5))

	_ = tmp1501

	tmp1502 := Call(__e, PrimNS2Value(symdef), sym_6, PrimNS1Value(sym_6))

	_ = tmp1502

	tmp1503 := Call(__e, PrimNS2Value(symdef), symsymbol_2, PrimNS1Value(symsymbol_2))

	_ = tmp1503

	tmp1504 := Call(__e, PrimNS2Value(symdef), symnot, PrimNS1Value(symnot))

	_ = tmp1504

	tmp1505 := Call(__e, PrimNS2Value(symdef), symnumber_2, PrimNS1Value(symnumber_2))

	_ = tmp1505

	tmp1506 := Call(__e, PrimNS2Value(symdef), symstring_2, PrimNS1Value(symstring_2))

	_ = tmp1506

	tmp1507 := Call(__e, PrimNS2Value(symdef), symabsvector, PrimNS1Value(symvector))

	_ = tmp1507

	tmp1508 := Call(__e, PrimNS2Value(symdef), symabsvector_2, PrimNS1Value(symvector_2))

	_ = tmp1508

	tmp1509 := Call(__e, PrimNS2Value(symdef), sym_5_1address, PrimNS1Value(symvector_1ref))

	_ = tmp1509

	__e.TailApply(PrimNS2Value(symdef), symaddress_1_6, PrimNS1Value(symvector_1set_b))
	return

}, 0)

var symreturn = MakeSymbol("return")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var syminferences = MakeSymbol("inferences")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4compile_1premises = MakeSymbol("shen.compile-premises")
var symnull = MakeSymbol("null")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var sympackage_2 = MakeSymbol("package?")
var sym_c = MakeSymbol("/")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symnumber = MakeSymbol("number")
var symloaded = MakeSymbol("loaded")
var symlimit = MakeSymbol("limit")
var symshen_4_ahd_2 = MakeSymbol("shen.=hd?")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symlength = MakeSymbol("length")
var symremove = MakeSymbol("remove")
var symshen_4branch = MakeSymbol("shen.branch")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4search = MakeSymbol("shen.search")
var symcdr = MakeSymbol("cdr")
var symshen_4mod = MakeSymbol("shen.mod")
var symfreeze = MakeSymbol("freeze")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var sym_3 = MakeSymbol("$")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symhd = MakeSymbol("hd")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symFinish = MakeSymbol("Finish")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symarity = MakeSymbol("arity")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var symmaxinferences = MakeSymbol("maxinferences")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symwhere = MakeSymbol("where")
var symfactorise = MakeSymbol("factorise")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4hds = MakeSymbol("shen.hds")
var symundefmacro = MakeSymbol("undefmacro")
var symuntrack = MakeSymbol("untrack")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symunput = MakeSymbol("unput")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symdefine = MakeSymbol("define")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4compile_1assumptions = MakeSymbol("shen.compile-assumptions")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symin = MakeSymbol("in")
var symshen = MakeSymbol("shen")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symshen_4_dlambdatable_d = MakeSymbol("shen.*lambdatable*")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symsymbol_2 = MakeSymbol("symbol?")
var symlanguage = MakeSymbol("language")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symspy = MakeSymbol("spy")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symlambda = MakeSymbol("lambda")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var syminclude = MakeSymbol("include")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4compile_1side_1condition = MakeSymbol("shen.compile-side-condition")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symhdstr = MakeSymbol("hdstr")
var symFreeze = MakeSymbol("Freeze")
var symmake_1string = MakeSymbol("make-string")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symnot = MakeSymbol("not")
var symbound_2 = MakeSymbol("bound?")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symcompile_1file = MakeSymbol("compile-file")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4goals = MakeSymbol("shen.goals")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symvector = MakeSymbol("vector")
var symconcat = MakeSymbol("concat")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symY = MakeSymbol("Y")
var symX = MakeSymbol("X")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symV = MakeSymbol("V")
var sym_1_6 = MakeSymbol("->")
var symnth = MakeSymbol("nth")
var symstinput = MakeSymbol("stinput")
var symnl = MakeSymbol("nl")
var symabort = MakeSymbol("abort")
var symshen_4selectors = MakeSymbol("shen.selectors")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symTm = MakeSymbol("Tm")
var symshen_4nchars = MakeSymbol("shen.nchars")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4deref = MakeSymbol("shen.deref")
var symparse_1app = MakeSymbol("parse-app")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symdatatype = MakeSymbol("datatype")
var symverified = MakeSymbol("verified")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symread = MakeSymbol("read")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4shen = MakeSymbol("shen.shen")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4optimise_1passive_1h = MakeSymbol("shen.optimise-passive-h")
var symtc_2 = MakeSymbol("tc?")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4u_b_1macro = MakeSymbol("shen.u!-macro")
var symu_b = MakeSymbol("u!")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4t = MakeSymbol("shen.t")
var sym__ = MakeSymbol("_")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symmacroexpand = MakeSymbol("macroexpand")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4decons = MakeSymbol("shen.decons")
var sym_d = MakeSymbol("*")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symset = MakeSymbol("set")
var symtuple_2 = MakeSymbol("tuple?")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var symis_b = MakeSymbol("is!")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symassoc = MakeSymbol("assoc")
var symmap = MakeSymbol("map")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symport = MakeSymbol("port")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symwrite_1sexp_1to_1file = MakeSymbol("write-sexp-to-file")
var symdifference = MakeSymbol("difference")
var sym_e_e = MakeSymbol("&&")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4compile_1search_1procedure = MakeSymbol("shen.compile-search-procedure")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symvariable_2 = MakeSymbol("variable?")
var symreverse = MakeSymbol("reverse")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4compile_1consequent = MakeSymbol("shen.compile-consequent")
var sym_a_a = MakeSymbol("==")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symstring_2 = MakeSymbol("string?")
var sym_c_4 = MakeSymbol("/.")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var sym_6 = MakeSymbol(">")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symversion = MakeSymbol("version")
var symshen_4else = MakeSymbol("shen.else")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symC = MakeSymbol("C")
var symif = MakeSymbol("if")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4headstream = MakeSymbol("shen.headstream")
var sympr = MakeSymbol("pr")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4split_1cases = MakeSymbol("shen.split-cases")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var sympackage = MakeSymbol("package")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symsum = MakeSymbol("sum")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var sym_8v = MakeSymbol("@v")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symfst = MakeSymbol("fst")
var symlazy = MakeSymbol("lazy")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symexternal = MakeSymbol("external")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symKey = MakeSymbol("Key")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symshen_4no_1action = MakeSymbol("shen.no-action")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symmapcan = MakeSymbol("mapcan")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4insert = MakeSymbol("shen.insert")
var sym_a = MakeSymbol("=")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symhash = MakeSymbol("hash")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symstoutput = MakeSymbol("stoutput")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4keep_1looking = MakeSymbol("shen.keep-looking")
var symshen_4vertical = MakeSymbol("shen.vertical")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symsave = MakeSymbol("save")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symfix = MakeSymbol("fix")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symGoTo = MakeSymbol("GoTo")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symimplementation = MakeSymbol("implementation")
var symdefun = MakeSymbol("defun")
var symfindall = MakeSymbol("findall")
var symempty_2 = MakeSymbol("empty?")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symParse = MakeSymbol("Parse")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symsystemf = MakeSymbol("systemf")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symstring = MakeSymbol("string")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var sym_8p = MakeSymbol("@p")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4pushsemantics = MakeSymbol("shen.pushsemantics")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var sym_b = MakeSymbol("!")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symD = MakeSymbol("D")
var symbegin = MakeSymbol("begin")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var sym_1 = MakeSymbol("-")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4eos = MakeSymbol("shen.eos")
var symdefprolog = MakeSymbol("defprolog")
var symSynCons = MakeSymbol("SynCons")
var symshen_4compile_1assumption = MakeSymbol("shen.compile-assumption")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symtlstr = MakeSymbol("tlstr")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symunit = MakeSymbol("unit")
var symshen_4tabulate_1passive = MakeSymbol("shen.tabulate-passive")
var symshen_4remove_1pointer = MakeSymbol("shen.remove-pointer")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symlist = MakeSymbol("list")
var sym_a_a_6 = MakeSymbol("==>")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symbar_b = MakeSymbol("bar!")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var sym_5_a = MakeSymbol("<=")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symdefcc = MakeSymbol("defcc")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symgensym = MakeSymbol("gensym")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var sym_i = MakeSymbol("{")
var symshen_4non_1empty_1stream_2 = MakeSymbol("shen.non-empty-stream?")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4multiples = MakeSymbol("shen.multiples")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symmode = MakeSymbol("mode")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4remove_1if_1unused = MakeSymbol("shen.remove-if-unused")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symget = MakeSymbol("get")
var symshen_4procedure_1call = MakeSymbol("shen.procedure-call")
var sym_dlanguage_d = MakeSymbol("*language*")
var symcn = MakeSymbol("cn")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4optimise_1passive = MakeSymbol("shen.optimise-passive")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var syminternal = MakeSymbol("internal")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symMessage = MakeSymbol("Message")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symcons = MakeSymbol("cons")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var sym_5_1 = MakeSymbol("<-")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symvector_1set_b = MakeSymbol("vector-set!")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4compile_1premise = MakeSymbol("shen.compile-premise")
var symshen_4source = MakeSymbol("shen.source")
var symwhen = MakeSymbol("when")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symNews = MakeSymbol("News")
var symintersection = MakeSymbol("intersection")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4passive_2 = MakeSymbol("shen.passive?")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symout = MakeSymbol("out")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var sym_dos_d = MakeSymbol("*os*")
var symtype = MakeSymbol("type")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4rpted_2 = MakeSymbol("shen.rpted?")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symshen_4join = MakeSymbol("shen.join")
var symquote = MakeSymbol("quote")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4reset_1prolog_1vector = MakeSymbol("shen.reset-prolog-vector")
var symfunction = MakeSymbol("function")
var symvalue = MakeSymbol("value")
var symatom_2 = MakeSymbol("atom?")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var sym_drelease_d = MakeSymbol("*release*")
var symcompile = MakeSymbol("compile")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symfile = MakeSymbol("file")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symcases = MakeSymbol("cases")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4let_b = MakeSymbol("shen.let!")
var sym_6_a = MakeSymbol(">=")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symtry_1catch = MakeSymbol("try-catch")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symsubst = MakeSymbol("subst")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symos = MakeSymbol("os")
var symwarn = MakeSymbol("warn")
var symdo = MakeSymbol("do")
var symput = MakeSymbol("put")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4tlstream = MakeSymbol("shen.tlstream")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symand = MakeSymbol("and")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symResult = MakeSymbol("Result")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symsnd = MakeSymbol("snd")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symlineread = MakeSymbol("lineread")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var sympos = MakeSymbol("pos")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symprotect = MakeSymbol("protect")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var sym_7 = MakeSymbol("+")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4unix = MakeSymbol("shen.unix")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symP = MakeSymbol("P")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symRecord = MakeSymbol("Record")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symthaw = MakeSymbol("thaw")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symclose = MakeSymbol("close")
var symboolean = MakeSymbol("boolean")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symcons_2 = MakeSymbol("cons?")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symkl_1_6cora = MakeSymbol("kl->cora")
var symoutput = MakeSymbol("output")
var symshen_4parsed_2 = MakeSymbol("shen.parsed?")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var symload = MakeSymbol("load")
var sym_6_6 = MakeSymbol(">>")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4body_1foundit_b = MakeSymbol("shen.body-foundit!")
var symin_1package = MakeSymbol("in-package")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4initialise_1lambda_1tables = MakeSymbol("shen.initialise-lambda-tables")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symabsvector = MakeSymbol("absvector")
var symshen_4passive = MakeSymbol("shen.passive")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symappend = MakeSymbol("append")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4compile_1side_1conditions = MakeSymbol("shen.compile-side-conditions")
var symshen_4inline_2 = MakeSymbol("shen.inline?")
var symrun = MakeSymbol("run")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symshen_4curry = MakeSymbol("shen.curry")
var symtlv = MakeSymbol("tlv")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symshen_4tame = MakeSymbol("shen.tame")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4cons_1form_1with_1modes = MakeSymbol("shen.cons-form-with-modes")
var symhead = MakeSymbol("head")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symK = MakeSymbol("K")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symopen = MakeSymbol("open")
var symwrite_1byte = MakeSymbol("write-byte")
var symA = MakeSymbol("A")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4modh = MakeSymbol("shen.modh")
var syminput = MakeSymbol("input")
var symrelease = MakeSymbol("release")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symelem_2 = MakeSymbol("elem?")
var syminteger_2 = MakeSymbol("integer?")
var symprint = MakeSymbol("print")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symPrevious = MakeSymbol("Previous")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var sym_j = MakeSymbol("}")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symstr = MakeSymbol("str")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symvector_2 = MakeSymbol("vector?")
var symtc = MakeSymbol("tc")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symsymbol = MakeSymbol("symbol")
var symfail = MakeSymbol("fail")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4correct = MakeSymbol("shen.correct")
var symshen_4defmacro_1macro = MakeSymbol("shen.defmacro-macro")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symfresh = MakeSymbol("fresh")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symlet = MakeSymbol("let")
var symerror = MakeSymbol("error")
var symeval = MakeSymbol("eval")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symfn = MakeSymbol("fn")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4interror = MakeSymbol("shen.interror")
var symtime = MakeSymbol("time")
var symshen_4head_1foundit_b = MakeSymbol("shen.head-foundit!")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symL = MakeSymbol("L")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symstep = MakeSymbol("step")
var sym_5_1vector = MakeSymbol("<-vector")
var symunion = MakeSymbol("union")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symreceive = MakeSymbol("receive")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4print_1residue = MakeSymbol("shen.print-residue")
var symread_1file = MakeSymbol("read-file")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symcar = MakeSymbol("car")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symnumber_2 = MakeSymbol("number?")
var symprofile = MakeSymbol("profile")
var symread_1file_1as_1sexp = MakeSymbol("read-file-as-sexp")
var symshen_4foundit_b = MakeSymbol("shen.foundit!")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4tls = MakeSymbol("shen.tls")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var sym_5 = MakeSymbol("<")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symbind = MakeSymbol("bind")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var sym_dport_d = MakeSymbol("*port*")
var syminline = MakeSymbol("inline")
var symtail = MakeSymbol("tail")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symB = MakeSymbol("B")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4change_1pointer_1value = MakeSymbol("shen.change-pointer-value")
var symporters = MakeSymbol("porters")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symintern = MakeSymbol("intern")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symshen_4s = MakeSymbol("shen.s")
var symadjoin = MakeSymbol("adjoin")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symtl = MakeSymbol("tl")
var symshen_4defcc_1macro = MakeSymbol("shen.defcc-macro")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symshen_4_dprolog_1vector_d = MakeSymbol("shen.*prolog-vector*")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4cons_1form_1no_1modes = MakeSymbol("shen.cons-form-no-modes")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var sympreclude = MakeSymbol("preclude")
var symis = MakeSymbol("is")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4horizontal = MakeSymbol("shen.horizontal")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symelement_2 = MakeSymbol("element?")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symStart = MakeSymbol("Start")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4hdstream = MakeSymbol("shen.hdstream")
var symvar_2 = MakeSymbol("var?")
var symshen_4pprint = MakeSymbol("shen.pprint")
var symfork = MakeSymbol("fork")
var symTime = MakeSymbol("Time")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4op2 = MakeSymbol("shen.op2")
var symor = MakeSymbol("or")
var symdef = MakeSymbol("def")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symS = MakeSymbol("S")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var syminput_7 = MakeSymbol("input+")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symunprofile = MakeSymbol("unprofile")
var symfail_1if = MakeSymbol("fail-if")
var symdeclare = MakeSymbol("declare")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4_5end_6 = MakeSymbol("shen.<end>")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4then = MakeSymbol("shen.then")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symstream = MakeSymbol("stream")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symcd = MakeSymbol("cd")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symspecialise = MakeSymbol("specialise")
var symshen_4compile_1premise_1h = MakeSymbol("shen.compile-premise-h")
var symcall = MakeSymbol("call")
var symexception = MakeSymbol("exception")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symcond = MakeSymbol("cond")
var symshen_4beta = MakeSymbol("shen.beta")
var symtrack = MakeSymbol("track")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4optimise_1typing = MakeSymbol("shen.optimise-typing")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symps = MakeSymbol("ps")
var symget_1time = MakeSymbol("get-time")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symhdv = MakeSymbol("hdv")
var symit = MakeSymbol("it")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symdestroy = MakeSymbol("destroy")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symoptimise = MakeSymbol("optimise")
var symvector_1ref = MakeSymbol("vector-ref")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symnull_2 = MakeSymbol("null?")
var sym_8s = MakeSymbol("@s")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symexplode = MakeSymbol("explode")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var sym_e = MakeSymbol("&")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
