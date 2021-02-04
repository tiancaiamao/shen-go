package main

import . "github.com/tiancaiamao/shen-go/kl"

var BuildMain = MakeNative(func(__e *ControlFlow) {
	tmp112 := MakeNative(func(__e *ControlFlow) {
		p86 := __e.Get(1)
		_ = p86
		p87 := __e.Get(2)
		_ = p87
		tmp113 := MakeNative(func(__e *ControlFlow) {
			cc88 := __e.Get(1)
			_ = cc88
			tmp114 := MakeNative(func(__e *ControlFlow) {
				env := __e.Get(1)
				_ = env
				tmp116 := PrimEqual(Nil, p87)

				if True == tmp116 {
					__e.Return(Nil)
					return
				} else {
					__e.TailApply(cc88)
					return
				}

			}, 1)

			__e.TailApply(tmp114, p86)
			return

		}, 1)

		tmp117 := MakeNative(func(__e *ControlFlow) {
			tmp118 := MakeNative(func(__e *ControlFlow) {
				cc89 := __e.Get(1)
				_ = cc89
				tmp119 := MakeNative(func(__e *ControlFlow) {
					env := __e.Get(1)
					_ = env
					tmp120 := MakeNative(func(__e *ControlFlow) {
						x := __e.Get(1)
						_ = x
						tmp127 := PrimIsNumber(x)

						var ifres122 Obj

						if True == tmp127 {
							ifres122 = True

						} else {
							tmp126 := PrimIsString(x)

							var ifres123 Obj

							if True == tmp126 {
								ifres123 = True

							} else {
								tmp125 := Call(__e, PrimNS1Value(symboolean_2), x)

								var ifres124 Obj

								if True == tmp125 {
									ifres124 = True

								} else {
									ifres124 = False

								}

								ifres123 = ifres124

							}

							ifres122 = ifres123

						}

						if True == ifres122 {
							__e.Return(x)
							return
						} else {
							__e.TailApply(cc89)
							return
						}

					}, 1)

					__e.TailApply(tmp120, p87)
					return

				}, 1)

				__e.TailApply(tmp119, p86)
				return

			}, 1)

			tmp128 := MakeNative(func(__e *ControlFlow) {
				tmp129 := MakeNative(func(__e *ControlFlow) {
					cc90 := __e.Get(1)
					_ = cc90
					tmp130 := MakeNative(func(__e *ControlFlow) {
						env := __e.Get(1)
						_ = env
						tmp193 := PrimIsPair(p87)

						var ifres189 Obj

						if True == tmp193 {
							tmp191 := Call(__e, PrimNS1Value(symnull_2), p87)

							tmp192 := PrimNot(tmp191)

							var ifres190 Obj

							if True == tmp192 {
								ifres190 = True

							} else {
								ifres190 = False

							}

							ifres189 = ifres190

						} else {
							ifres189 = False

						}

						if True == ifres189 {
							tmp187 := PrimHead(p87)

							tmp188 := PrimEqual(symif, tmp187)

							if True == tmp188 {
								tmp185 := PrimTail(p87)

								tmp186 := PrimIsPair(tmp185)

								var ifres180 Obj

								if True == tmp186 {
									tmp182 := PrimTail(p87)

									tmp183 := Call(__e, PrimNS1Value(symnull_2), tmp182)

									tmp184 := PrimNot(tmp183)

									var ifres181 Obj

									if True == tmp184 {
										ifres181 = True

									} else {
										ifres181 = False

									}

									ifres180 = ifres181

								} else {
									ifres180 = False

								}

								if True == ifres180 {
									tmp134 := MakeNative(func(__e *ControlFlow) {
										x := __e.Get(1)
										_ = x
										tmp175 := PrimTail(p87)

										tmp176 := PrimTail(tmp175)

										tmp177 := PrimIsPair(tmp176)

										var ifres169 Obj

										if True == tmp177 {
											tmp171 := PrimTail(p87)

											tmp172 := PrimTail(tmp171)

											tmp173 := Call(__e, PrimNS1Value(symnull_2), tmp172)

											tmp174 := PrimNot(tmp173)

											var ifres170 Obj

											if True == tmp174 {
												ifres170 = True

											} else {
												ifres170 = False

											}

											ifres169 = ifres170

										} else {
											ifres169 = False

										}

										if True == ifres169 {
											tmp136 := MakeNative(func(__e *ControlFlow) {
												y := __e.Get(1)
												_ = y
												tmp162 := PrimTail(p87)

												tmp163 := PrimTail(tmp162)

												tmp164 := PrimTail(tmp163)

												tmp165 := PrimIsPair(tmp164)

												var ifres155 Obj

												if True == tmp165 {
													tmp157 := PrimTail(p87)

													tmp158 := PrimTail(tmp157)

													tmp159 := PrimTail(tmp158)

													tmp160 := Call(__e, PrimNS1Value(symnull_2), tmp159)

													tmp161 := PrimNot(tmp160)

													var ifres156 Obj

													if True == tmp161 {
														ifres156 = True

													} else {
														ifres156 = False

													}

													ifres155 = ifres156

												} else {
													ifres155 = False

												}

												if True == ifres155 {
													tmp138 := MakeNative(func(__e *ControlFlow) {
														z := __e.Get(1)
														_ = z
														tmp146 := PrimTail(p87)

														tmp147 := PrimTail(tmp146)

														tmp148 := PrimTail(tmp147)

														tmp149 := PrimTail(tmp148)

														tmp150 := Call(__e, PrimNS1Value(symnull_2), tmp149)

														if True == tmp150 {
															tmp140 := Call(__e, PrimNS1Value(symkl_1_6cora), env, x)

															tmp141 := Call(__e, PrimNS1Value(symkl_1_6cora), env, y)

															tmp142 := Call(__e, PrimNS1Value(symkl_1_6cora), env, z)

															tmp143 := PrimCons(tmp142, Nil)

															tmp144 := PrimCons(tmp141, tmp143)

															tmp145 := PrimCons(tmp140, tmp144)

															__e.Return(PrimCons(symif, tmp145))
															return

														} else {
															__e.TailApply(cc90)
															return
														}

													}, 1)

													tmp151 := PrimTail(p87)

													tmp152 := PrimTail(tmp151)

													tmp153 := PrimTail(tmp152)

													tmp154 := PrimHead(tmp153)

													__e.TailApply(tmp138, tmp154)
													return

												} else {
													__e.TailApply(cc90)
													return
												}

											}, 1)

											tmp166 := PrimTail(p87)

											tmp167 := PrimTail(tmp166)

											tmp168 := PrimHead(tmp167)

											__e.TailApply(tmp136, tmp168)
											return

										} else {
											__e.TailApply(cc90)
											return
										}

									}, 1)

									tmp178 := PrimTail(p87)

									tmp179 := PrimHead(tmp178)

									__e.TailApply(tmp134, tmp179)
									return

								} else {
									__e.TailApply(cc90)
									return
								}

							} else {
								__e.TailApply(cc90)
								return
							}

						} else {
							__e.TailApply(cc90)
							return
						}

					}, 1)

					__e.TailApply(tmp130, p86)
					return

				}, 1)

				tmp194 := MakeNative(func(__e *ControlFlow) {
					tmp195 := MakeNative(func(__e *ControlFlow) {
						cc91 := __e.Get(1)
						_ = cc91
						tmp196 := MakeNative(func(__e *ControlFlow) {
							env := __e.Get(1)
							_ = env
							tmp239 := PrimIsPair(p87)

							var ifres235 Obj

							if True == tmp239 {
								tmp237 := Call(__e, PrimNS1Value(symnull_2), p87)

								tmp238 := PrimNot(tmp237)

								var ifres236 Obj

								if True == tmp238 {
									ifres236 = True

								} else {
									ifres236 = False

								}

								ifres235 = ifres236

							} else {
								ifres235 = False

							}

							if True == ifres235 {
								tmp233 := PrimHead(p87)

								tmp234 := PrimEqual(symdo, tmp233)

								if True == tmp234 {
									tmp231 := PrimTail(p87)

									tmp232 := PrimIsPair(tmp231)

									var ifres226 Obj

									if True == tmp232 {
										tmp228 := PrimTail(p87)

										tmp229 := Call(__e, PrimNS1Value(symnull_2), tmp228)

										tmp230 := PrimNot(tmp229)

										var ifres227 Obj

										if True == tmp230 {
											ifres227 = True

										} else {
											ifres227 = False

										}

										ifres226 = ifres227

									} else {
										ifres226 = False

									}

									if True == ifres226 {
										tmp200 := MakeNative(func(__e *ControlFlow) {
											x := __e.Get(1)
											_ = x
											tmp221 := PrimTail(p87)

											tmp222 := PrimTail(tmp221)

											tmp223 := PrimIsPair(tmp222)

											var ifres215 Obj

											if True == tmp223 {
												tmp217 := PrimTail(p87)

												tmp218 := PrimTail(tmp217)

												tmp219 := Call(__e, PrimNS1Value(symnull_2), tmp218)

												tmp220 := PrimNot(tmp219)

												var ifres216 Obj

												if True == tmp220 {
													ifres216 = True

												} else {
													ifres216 = False

												}

												ifres215 = ifres216

											} else {
												ifres215 = False

											}

											if True == ifres215 {
												tmp202 := MakeNative(func(__e *ControlFlow) {
													y := __e.Get(1)
													_ = y
													tmp208 := PrimTail(p87)

													tmp209 := PrimTail(tmp208)

													tmp210 := PrimTail(tmp209)

													tmp211 := Call(__e, PrimNS1Value(symnull_2), tmp210)

													if True == tmp211 {
														tmp204 := Call(__e, PrimNS1Value(symkl_1_6cora), env, x)

														tmp205 := Call(__e, PrimNS1Value(symkl_1_6cora), env, y)

														tmp206 := PrimCons(tmp205, Nil)

														tmp207 := PrimCons(tmp204, tmp206)

														__e.Return(PrimCons(symdo, tmp207))
														return

													} else {
														__e.TailApply(cc91)
														return
													}

												}, 1)

												tmp212 := PrimTail(p87)

												tmp213 := PrimTail(tmp212)

												tmp214 := PrimHead(tmp213)

												__e.TailApply(tmp202, tmp214)
												return

											} else {
												__e.TailApply(cc91)
												return
											}

										}, 1)

										tmp224 := PrimTail(p87)

										tmp225 := PrimHead(tmp224)

										__e.TailApply(tmp200, tmp225)
										return

									} else {
										__e.TailApply(cc91)
										return
									}

								} else {
									__e.TailApply(cc91)
									return
								}

							} else {
								__e.TailApply(cc91)
								return
							}

						}, 1)

						__e.TailApply(tmp196, p86)
						return

					}, 1)

					tmp240 := MakeNative(func(__e *ControlFlow) {
						tmp241 := MakeNative(func(__e *ControlFlow) {
							cc92 := __e.Get(1)
							_ = cc92
							tmp242 := MakeNative(func(__e *ControlFlow) {
								env := __e.Get(1)
								_ = env
								tmp286 := PrimIsPair(p87)

								var ifres282 Obj

								if True == tmp286 {
									tmp284 := Call(__e, PrimNS1Value(symnull_2), p87)

									tmp285 := PrimNot(tmp284)

									var ifres283 Obj

									if True == tmp285 {
										ifres283 = True

									} else {
										ifres283 = False

									}

									ifres282 = ifres283

								} else {
									ifres282 = False

								}

								if True == ifres282 {
									tmp280 := PrimHead(p87)

									tmp281 := PrimEqual(symlambda, tmp280)

									if True == tmp281 {
										tmp278 := PrimTail(p87)

										tmp279 := PrimIsPair(tmp278)

										var ifres273 Obj

										if True == tmp279 {
											tmp275 := PrimTail(p87)

											tmp276 := Call(__e, PrimNS1Value(symnull_2), tmp275)

											tmp277 := PrimNot(tmp276)

											var ifres274 Obj

											if True == tmp277 {
												ifres274 = True

											} else {
												ifres274 = False

											}

											ifres273 = ifres274

										} else {
											ifres273 = False

										}

										if True == ifres273 {
											tmp246 := MakeNative(func(__e *ControlFlow) {
												x := __e.Get(1)
												_ = x
												tmp268 := PrimTail(p87)

												tmp269 := PrimTail(tmp268)

												tmp270 := PrimIsPair(tmp269)

												var ifres262 Obj

												if True == tmp270 {
													tmp264 := PrimTail(p87)

													tmp265 := PrimTail(tmp264)

													tmp266 := Call(__e, PrimNS1Value(symnull_2), tmp265)

													tmp267 := PrimNot(tmp266)

													var ifres263 Obj

													if True == tmp267 {
														ifres263 = True

													} else {
														ifres263 = False

													}

													ifres262 = ifres263

												} else {
													ifres262 = False

												}

												if True == ifres262 {
													tmp248 := MakeNative(func(__e *ControlFlow) {
														y := __e.Get(1)
														_ = y
														tmp255 := PrimTail(p87)

														tmp256 := PrimTail(tmp255)

														tmp257 := PrimTail(tmp256)

														tmp258 := Call(__e, PrimNS1Value(symnull_2), tmp257)

														if True == tmp258 {
															tmp250 := PrimCons(x, Nil)

															tmp251 := PrimCons(x, env)

															tmp252 := Call(__e, PrimNS1Value(symkl_1_6cora), tmp251, y)

															tmp253 := PrimCons(tmp252, Nil)

															tmp254 := PrimCons(tmp250, tmp253)

															__e.Return(PrimCons(symlambda, tmp254))
															return

														} else {
															__e.TailApply(cc92)
															return
														}

													}, 1)

													tmp259 := PrimTail(p87)

													tmp260 := PrimTail(tmp259)

													tmp261 := PrimHead(tmp260)

													__e.TailApply(tmp248, tmp261)
													return

												} else {
													__e.TailApply(cc92)
													return
												}

											}, 1)

											tmp271 := PrimTail(p87)

											tmp272 := PrimHead(tmp271)

											__e.TailApply(tmp246, tmp272)
											return

										} else {
											__e.TailApply(cc92)
											return
										}

									} else {
										__e.TailApply(cc92)
										return
									}

								} else {
									__e.TailApply(cc92)
									return
								}

							}, 1)

							__e.TailApply(tmp242, p86)
							return

						}, 1)

						tmp287 := MakeNative(func(__e *ControlFlow) {
							tmp288 := MakeNative(func(__e *ControlFlow) {
								cc93 := __e.Get(1)
								_ = cc93
								tmp289 := MakeNative(func(__e *ControlFlow) {
									env := __e.Get(1)
									_ = env
									tmp354 := PrimIsPair(p87)

									var ifres350 Obj

									if True == tmp354 {
										tmp352 := Call(__e, PrimNS1Value(symnull_2), p87)

										tmp353 := PrimNot(tmp352)

										var ifres351 Obj

										if True == tmp353 {
											ifres351 = True

										} else {
											ifres351 = False

										}

										ifres350 = ifres351

									} else {
										ifres350 = False

									}

									if True == ifres350 {
										tmp348 := PrimHead(p87)

										tmp349 := PrimEqual(symdefun, tmp348)

										if True == tmp349 {
											tmp346 := PrimTail(p87)

											tmp347 := PrimIsPair(tmp346)

											var ifres341 Obj

											if True == tmp347 {
												tmp343 := PrimTail(p87)

												tmp344 := Call(__e, PrimNS1Value(symnull_2), tmp343)

												tmp345 := PrimNot(tmp344)

												var ifres342 Obj

												if True == tmp345 {
													ifres342 = True

												} else {
													ifres342 = False

												}

												ifres341 = ifres342

											} else {
												ifres341 = False

											}

											if True == ifres341 {
												tmp293 := MakeNative(func(__e *ControlFlow) {
													f := __e.Get(1)
													_ = f
													tmp336 := PrimTail(p87)

													tmp337 := PrimTail(tmp336)

													tmp338 := PrimIsPair(tmp337)

													var ifres330 Obj

													if True == tmp338 {
														tmp332 := PrimTail(p87)

														tmp333 := PrimTail(tmp332)

														tmp334 := Call(__e, PrimNS1Value(symnull_2), tmp333)

														tmp335 := PrimNot(tmp334)

														var ifres331 Obj

														if True == tmp335 {
															ifres331 = True

														} else {
															ifres331 = False

														}

														ifres330 = ifres331

													} else {
														ifres330 = False

													}

													if True == ifres330 {
														tmp295 := MakeNative(func(__e *ControlFlow) {
															x := __e.Get(1)
															_ = x
															tmp323 := PrimTail(p87)

															tmp324 := PrimTail(tmp323)

															tmp325 := PrimTail(tmp324)

															tmp326 := PrimIsPair(tmp325)

															var ifres316 Obj

															if True == tmp326 {
																tmp318 := PrimTail(p87)

																tmp319 := PrimTail(tmp318)

																tmp320 := PrimTail(tmp319)

																tmp321 := Call(__e, PrimNS1Value(symnull_2), tmp320)

																tmp322 := PrimNot(tmp321)

																var ifres317 Obj

																if True == tmp322 {
																	ifres317 = True

																} else {
																	ifres317 = False

																}

																ifres316 = ifres317

															} else {
																ifres316 = False

															}

															if True == ifres316 {
																tmp297 := MakeNative(func(__e *ControlFlow) {
																	y := __e.Get(1)
																	_ = y
																	tmp307 := PrimTail(p87)

																	tmp308 := PrimTail(tmp307)

																	tmp309 := PrimTail(tmp308)

																	tmp310 := PrimTail(tmp309)

																	tmp311 := Call(__e, PrimNS1Value(symnull_2), tmp310)

																	if True == tmp311 {
																		tmp299 := PrimCons(f, Nil)

																		tmp300 := PrimCons(symquote, tmp299)

																		tmp301 := Call(__e, PrimNS1Value(symkl_1_6cora), x, y)

																		tmp302 := PrimCons(tmp301, Nil)

																		tmp303 := PrimCons(x, tmp302)

																		tmp304 := PrimCons(symlambda, tmp303)

																		tmp305 := PrimCons(tmp304, Nil)

																		tmp306 := PrimCons(tmp300, tmp305)

																		__e.Return(PrimCons(symns2_1set, tmp306))
																		return

																	} else {
																		__e.TailApply(cc93)
																		return
																	}

																}, 1)

																tmp312 := PrimTail(p87)

																tmp313 := PrimTail(tmp312)

																tmp314 := PrimTail(tmp313)

																tmp315 := PrimHead(tmp314)

																__e.TailApply(tmp297, tmp315)
																return

															} else {
																__e.TailApply(cc93)
																return
															}

														}, 1)

														tmp327 := PrimTail(p87)

														tmp328 := PrimTail(tmp327)

														tmp329 := PrimHead(tmp328)

														__e.TailApply(tmp295, tmp329)
														return

													} else {
														__e.TailApply(cc93)
														return
													}

												}, 1)

												tmp339 := PrimTail(p87)

												tmp340 := PrimHead(tmp339)

												__e.TailApply(tmp293, tmp340)
												return

											} else {
												__e.TailApply(cc93)
												return
											}

										} else {
											__e.TailApply(cc93)
											return
										}

									} else {
										__e.TailApply(cc93)
										return
									}

								}, 1)

								__e.TailApply(tmp289, p86)
								return

							}, 1)

							tmp355 := MakeNative(func(__e *ControlFlow) {
								tmp356 := MakeNative(func(__e *ControlFlow) {
									cc94 := __e.Get(1)
									_ = cc94
									tmp357 := MakeNative(func(__e *ControlFlow) {
										env := __e.Get(1)
										_ = env
										tmp422 := PrimIsPair(p87)

										var ifres418 Obj

										if True == tmp422 {
											tmp420 := Call(__e, PrimNS1Value(symnull_2), p87)

											tmp421 := PrimNot(tmp420)

											var ifres419 Obj

											if True == tmp421 {
												ifres419 = True

											} else {
												ifres419 = False

											}

											ifres418 = ifres419

										} else {
											ifres418 = False

										}

										if True == ifres418 {
											tmp416 := PrimHead(p87)

											tmp417 := PrimEqual(symlet, tmp416)

											if True == tmp417 {
												tmp414 := PrimTail(p87)

												tmp415 := PrimIsPair(tmp414)

												var ifres409 Obj

												if True == tmp415 {
													tmp411 := PrimTail(p87)

													tmp412 := Call(__e, PrimNS1Value(symnull_2), tmp411)

													tmp413 := PrimNot(tmp412)

													var ifres410 Obj

													if True == tmp413 {
														ifres410 = True

													} else {
														ifres410 = False

													}

													ifres409 = ifres410

												} else {
													ifres409 = False

												}

												if True == ifres409 {
													tmp361 := MakeNative(func(__e *ControlFlow) {
														x := __e.Get(1)
														_ = x
														tmp404 := PrimTail(p87)

														tmp405 := PrimTail(tmp404)

														tmp406 := PrimIsPair(tmp405)

														var ifres398 Obj

														if True == tmp406 {
															tmp400 := PrimTail(p87)

															tmp401 := PrimTail(tmp400)

															tmp402 := Call(__e, PrimNS1Value(symnull_2), tmp401)

															tmp403 := PrimNot(tmp402)

															var ifres399 Obj

															if True == tmp403 {
																ifres399 = True

															} else {
																ifres399 = False

															}

															ifres398 = ifres399

														} else {
															ifres398 = False

														}

														if True == ifres398 {
															tmp363 := MakeNative(func(__e *ControlFlow) {
																y := __e.Get(1)
																_ = y
																tmp391 := PrimTail(p87)

																tmp392 := PrimTail(tmp391)

																tmp393 := PrimTail(tmp392)

																tmp394 := PrimIsPair(tmp393)

																var ifres384 Obj

																if True == tmp394 {
																	tmp386 := PrimTail(p87)

																	tmp387 := PrimTail(tmp386)

																	tmp388 := PrimTail(tmp387)

																	tmp389 := Call(__e, PrimNS1Value(symnull_2), tmp388)

																	tmp390 := PrimNot(tmp389)

																	var ifres385 Obj

																	if True == tmp390 {
																		ifres385 = True

																	} else {
																		ifres385 = False

																	}

																	ifres384 = ifres385

																} else {
																	ifres384 = False

																}

																if True == ifres384 {
																	tmp365 := MakeNative(func(__e *ControlFlow) {
																		z := __e.Get(1)
																		_ = z
																		tmp375 := PrimTail(p87)

																		tmp376 := PrimTail(tmp375)

																		tmp377 := PrimTail(tmp376)

																		tmp378 := PrimTail(tmp377)

																		tmp379 := Call(__e, PrimNS1Value(symnull_2), tmp378)

																		if True == tmp379 {
																			tmp367 := PrimCons(x, Nil)

																			tmp368 := PrimCons(x, env)

																			tmp369 := Call(__e, PrimNS1Value(symkl_1_6cora), tmp368, z)

																			tmp370 := PrimCons(tmp369, Nil)

																			tmp371 := PrimCons(tmp367, tmp370)

																			tmp372 := PrimCons(symlambda, tmp371)

																			tmp373 := Call(__e, PrimNS1Value(symkl_1_6cora), env, y)

																			tmp374 := PrimCons(tmp373, Nil)

																			__e.Return(PrimCons(tmp372, tmp374))
																			return

																		} else {
																			__e.TailApply(cc94)
																			return
																		}

																	}, 1)

																	tmp380 := PrimTail(p87)

																	tmp381 := PrimTail(tmp380)

																	tmp382 := PrimTail(tmp381)

																	tmp383 := PrimHead(tmp382)

																	__e.TailApply(tmp365, tmp383)
																	return

																} else {
																	__e.TailApply(cc94)
																	return
																}

															}, 1)

															tmp395 := PrimTail(p87)

															tmp396 := PrimTail(tmp395)

															tmp397 := PrimHead(tmp396)

															__e.TailApply(tmp363, tmp397)
															return

														} else {
															__e.TailApply(cc94)
															return
														}

													}, 1)

													tmp407 := PrimTail(p87)

													tmp408 := PrimHead(tmp407)

													__e.TailApply(tmp361, tmp408)
													return

												} else {
													__e.TailApply(cc94)
													return
												}

											} else {
												__e.TailApply(cc94)
												return
											}

										} else {
											__e.TailApply(cc94)
											return
										}

									}, 1)

									__e.TailApply(tmp357, p86)
									return

								}, 1)

								tmp423 := MakeNative(func(__e *ControlFlow) {
									tmp424 := MakeNative(func(__e *ControlFlow) {
										cc95 := __e.Get(1)
										_ = cc95
										tmp425 := MakeNative(func(__e *ControlFlow) {
											env := __e.Get(1)
											_ = env
											tmp471 := PrimIsPair(p87)

											var ifres467 Obj

											if True == tmp471 {
												tmp469 := Call(__e, PrimNS1Value(symnull_2), p87)

												tmp470 := PrimNot(tmp469)

												var ifres468 Obj

												if True == tmp470 {
													ifres468 = True

												} else {
													ifres468 = False

												}

												ifres467 = ifres468

											} else {
												ifres467 = False

											}

											if True == ifres467 {
												tmp465 := PrimHead(p87)

												tmp466 := PrimEqual(symtrap_1error, tmp465)

												if True == tmp466 {
													tmp463 := PrimTail(p87)

													tmp464 := PrimIsPair(tmp463)

													var ifres458 Obj

													if True == tmp464 {
														tmp460 := PrimTail(p87)

														tmp461 := Call(__e, PrimNS1Value(symnull_2), tmp460)

														tmp462 := PrimNot(tmp461)

														var ifres459 Obj

														if True == tmp462 {
															ifres459 = True

														} else {
															ifres459 = False

														}

														ifres458 = ifres459

													} else {
														ifres458 = False

													}

													if True == ifres458 {
														tmp429 := MakeNative(func(__e *ControlFlow) {
															body := __e.Get(1)
															_ = body
															tmp453 := PrimTail(p87)

															tmp454 := PrimTail(tmp453)

															tmp455 := PrimIsPair(tmp454)

															var ifres447 Obj

															if True == tmp455 {
																tmp449 := PrimTail(p87)

																tmp450 := PrimTail(tmp449)

																tmp451 := Call(__e, PrimNS1Value(symnull_2), tmp450)

																tmp452 := PrimNot(tmp451)

																var ifres448 Obj

																if True == tmp452 {
																	ifres448 = True

																} else {
																	ifres448 = False

																}

																ifres447 = ifres448

															} else {
																ifres447 = False

															}

															if True == ifres447 {
																tmp431 := MakeNative(func(__e *ControlFlow) {
																	handler := __e.Get(1)
																	_ = handler
																	tmp440 := PrimTail(p87)

																	tmp441 := PrimTail(tmp440)

																	tmp442 := PrimTail(tmp441)

																	tmp443 := Call(__e, PrimNS1Value(symnull_2), tmp442)

																	if True == tmp443 {
																		tmp433 := Call(__e, PrimNS1Value(symkl_1_6cora), env, body)

																		tmp434 := PrimCons(tmp433, Nil)

																		tmp435 := PrimCons(Nil, tmp434)

																		tmp436 := PrimCons(symlambda, tmp435)

																		tmp437 := Call(__e, PrimNS1Value(symkl_1_6cora), env, handler)

																		tmp438 := PrimCons(tmp437, Nil)

																		tmp439 := PrimCons(tmp436, tmp438)

																		__e.Return(PrimCons(symtry_1catch, tmp439))
																		return

																	} else {
																		__e.TailApply(cc95)
																		return
																	}

																}, 1)

																tmp444 := PrimTail(p87)

																tmp445 := PrimTail(tmp444)

																tmp446 := PrimHead(tmp445)

																__e.TailApply(tmp431, tmp446)
																return

															} else {
																__e.TailApply(cc95)
																return
															}

														}, 1)

														tmp456 := PrimTail(p87)

														tmp457 := PrimHead(tmp456)

														__e.TailApply(tmp429, tmp457)
														return

													} else {
														__e.TailApply(cc95)
														return
													}

												} else {
													__e.TailApply(cc95)
													return
												}

											} else {
												__e.TailApply(cc95)
												return
											}

										}, 1)

										__e.TailApply(tmp425, p86)
										return

									}, 1)

									tmp472 := MakeNative(func(__e *ControlFlow) {
										tmp473 := MakeNative(func(__e *ControlFlow) {
											cc96 := __e.Get(1)
											_ = cc96
											tmp474 := MakeNative(func(__e *ControlFlow) {
												env := __e.Get(1)
												_ = env
												tmp521 := PrimIsPair(p87)

												var ifres517 Obj

												if True == tmp521 {
													tmp519 := Call(__e, PrimNS1Value(symnull_2), p87)

													tmp520 := PrimNot(tmp519)

													var ifres518 Obj

													if True == tmp520 {
														ifres518 = True

													} else {
														ifres518 = False

													}

													ifres517 = ifres518

												} else {
													ifres517 = False

												}

												if True == ifres517 {
													tmp515 := PrimHead(p87)

													tmp516 := PrimEqual(symor, tmp515)

													if True == tmp516 {
														tmp513 := PrimTail(p87)

														tmp514 := PrimIsPair(tmp513)

														var ifres508 Obj

														if True == tmp514 {
															tmp510 := PrimTail(p87)

															tmp511 := Call(__e, PrimNS1Value(symnull_2), tmp510)

															tmp512 := PrimNot(tmp511)

															var ifres509 Obj

															if True == tmp512 {
																ifres509 = True

															} else {
																ifres509 = False

															}

															ifres508 = ifres509

														} else {
															ifres508 = False

														}

														if True == ifres508 {
															tmp478 := MakeNative(func(__e *ControlFlow) {
																x := __e.Get(1)
																_ = x
																tmp503 := PrimTail(p87)

																tmp504 := PrimTail(tmp503)

																tmp505 := PrimIsPair(tmp504)

																var ifres497 Obj

																if True == tmp505 {
																	tmp499 := PrimTail(p87)

																	tmp500 := PrimTail(tmp499)

																	tmp501 := Call(__e, PrimNS1Value(symnull_2), tmp500)

																	tmp502 := PrimNot(tmp501)

																	var ifres498 Obj

																	if True == tmp502 {
																		ifres498 = True

																	} else {
																		ifres498 = False

																	}

																	ifres497 = ifres498

																} else {
																	ifres497 = False

																}

																if True == ifres497 {
																	tmp480 := MakeNative(func(__e *ControlFlow) {
																		y := __e.Get(1)
																		_ = y
																		tmp490 := PrimTail(p87)

																		tmp491 := PrimTail(tmp490)

																		tmp492 := PrimTail(tmp491)

																		tmp493 := Call(__e, PrimNS1Value(symnull_2), tmp492)

																		if True == tmp493 {
																			tmp482 := PrimCons(False, Nil)

																			tmp483 := PrimCons(True, tmp482)

																			tmp484 := PrimCons(y, tmp483)

																			tmp485 := PrimCons(symif, tmp484)

																			tmp486 := PrimCons(tmp485, Nil)

																			tmp487 := PrimCons(True, tmp486)

																			tmp488 := PrimCons(x, tmp487)

																			tmp489 := PrimCons(symif, tmp488)

																			__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp489)
																			return

																		} else {
																			__e.TailApply(cc96)
																			return
																		}

																	}, 1)

																	tmp494 := PrimTail(p87)

																	tmp495 := PrimTail(tmp494)

																	tmp496 := PrimHead(tmp495)

																	__e.TailApply(tmp480, tmp496)
																	return

																} else {
																	__e.TailApply(cc96)
																	return
																}

															}, 1)

															tmp506 := PrimTail(p87)

															tmp507 := PrimHead(tmp506)

															__e.TailApply(tmp478, tmp507)
															return

														} else {
															__e.TailApply(cc96)
															return
														}

													} else {
														__e.TailApply(cc96)
														return
													}

												} else {
													__e.TailApply(cc96)
													return
												}

											}, 1)

											__e.TailApply(tmp474, p86)
											return

										}, 1)

										tmp522 := MakeNative(func(__e *ControlFlow) {
											tmp523 := MakeNative(func(__e *ControlFlow) {
												cc97 := __e.Get(1)
												_ = cc97
												tmp524 := MakeNative(func(__e *ControlFlow) {
													env := __e.Get(1)
													_ = env
													tmp571 := PrimIsPair(p87)

													var ifres567 Obj

													if True == tmp571 {
														tmp569 := Call(__e, PrimNS1Value(symnull_2), p87)

														tmp570 := PrimNot(tmp569)

														var ifres568 Obj

														if True == tmp570 {
															ifres568 = True

														} else {
															ifres568 = False

														}

														ifres567 = ifres568

													} else {
														ifres567 = False

													}

													if True == ifres567 {
														tmp565 := PrimHead(p87)

														tmp566 := PrimEqual(symand, tmp565)

														if True == tmp566 {
															tmp563 := PrimTail(p87)

															tmp564 := PrimIsPair(tmp563)

															var ifres558 Obj

															if True == tmp564 {
																tmp560 := PrimTail(p87)

																tmp561 := Call(__e, PrimNS1Value(symnull_2), tmp560)

																tmp562 := PrimNot(tmp561)

																var ifres559 Obj

																if True == tmp562 {
																	ifres559 = True

																} else {
																	ifres559 = False

																}

																ifres558 = ifres559

															} else {
																ifres558 = False

															}

															if True == ifres558 {
																tmp528 := MakeNative(func(__e *ControlFlow) {
																	x := __e.Get(1)
																	_ = x
																	tmp553 := PrimTail(p87)

																	tmp554 := PrimTail(tmp553)

																	tmp555 := PrimIsPair(tmp554)

																	var ifres547 Obj

																	if True == tmp555 {
																		tmp549 := PrimTail(p87)

																		tmp550 := PrimTail(tmp549)

																		tmp551 := Call(__e, PrimNS1Value(symnull_2), tmp550)

																		tmp552 := PrimNot(tmp551)

																		var ifres548 Obj

																		if True == tmp552 {
																			ifres548 = True

																		} else {
																			ifres548 = False

																		}

																		ifres547 = ifres548

																	} else {
																		ifres547 = False

																	}

																	if True == ifres547 {
																		tmp530 := MakeNative(func(__e *ControlFlow) {
																			y := __e.Get(1)
																			_ = y
																			tmp540 := PrimTail(p87)

																			tmp541 := PrimTail(tmp540)

																			tmp542 := PrimTail(tmp541)

																			tmp543 := Call(__e, PrimNS1Value(symnull_2), tmp542)

																			if True == tmp543 {
																				tmp532 := PrimCons(False, Nil)

																				tmp533 := PrimCons(True, tmp532)

																				tmp534 := PrimCons(y, tmp533)

																				tmp535 := PrimCons(symif, tmp534)

																				tmp536 := PrimCons(False, Nil)

																				tmp537 := PrimCons(tmp535, tmp536)

																				tmp538 := PrimCons(x, tmp537)

																				tmp539 := PrimCons(symif, tmp538)

																				__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp539)
																				return

																			} else {
																				__e.TailApply(cc97)
																				return
																			}

																		}, 1)

																		tmp544 := PrimTail(p87)

																		tmp545 := PrimTail(tmp544)

																		tmp546 := PrimHead(tmp545)

																		__e.TailApply(tmp530, tmp546)
																		return

																	} else {
																		__e.TailApply(cc97)
																		return
																	}

																}, 1)

																tmp556 := PrimTail(p87)

																tmp557 := PrimHead(tmp556)

																__e.TailApply(tmp528, tmp557)
																return

															} else {
																__e.TailApply(cc97)
																return
															}

														} else {
															__e.TailApply(cc97)
															return
														}

													} else {
														__e.TailApply(cc97)
														return
													}

												}, 1)

												__e.TailApply(tmp524, p86)
												return

											}, 1)

											tmp572 := MakeNative(func(__e *ControlFlow) {
												tmp573 := MakeNative(func(__e *ControlFlow) {
													cc98 := __e.Get(1)
													_ = cc98
													tmp574 := MakeNative(func(__e *ControlFlow) {
														env := __e.Get(1)
														_ = env
														tmp638 := PrimIsPair(p87)

														var ifres634 Obj

														if True == tmp638 {
															tmp636 := Call(__e, PrimNS1Value(symnull_2), p87)

															tmp637 := PrimNot(tmp636)

															var ifres635 Obj

															if True == tmp637 {
																ifres635 = True

															} else {
																ifres635 = False

															}

															ifres634 = ifres635

														} else {
															ifres634 = False

														}

														if True == ifres634 {
															tmp632 := PrimHead(p87)

															tmp633 := PrimEqual(symcond, tmp632)

															if True == tmp633 {
																tmp630 := PrimTail(p87)

																tmp631 := PrimIsPair(tmp630)

																var ifres625 Obj

																if True == tmp631 {
																	tmp627 := PrimTail(p87)

																	tmp628 := Call(__e, PrimNS1Value(symnull_2), tmp627)

																	tmp629 := PrimNot(tmp628)

																	var ifres626 Obj

																	if True == tmp629 {
																		ifres626 = True

																	} else {
																		ifres626 = False

																	}

																	ifres625 = ifres626

																} else {
																	ifres625 = False

																}

																if True == ifres625 {
																	tmp622 := PrimTail(p87)

																	tmp623 := PrimHead(tmp622)

																	tmp624 := PrimIsPair(tmp623)

																	var ifres616 Obj

																	if True == tmp624 {
																		tmp618 := PrimTail(p87)

																		tmp619 := PrimHead(tmp618)

																		tmp620 := Call(__e, PrimNS1Value(symnull_2), tmp619)

																		tmp621 := PrimNot(tmp620)

																		var ifres617 Obj

																		if True == tmp621 {
																			ifres617 = True

																		} else {
																			ifres617 = False

																		}

																		ifres616 = ifres617

																	} else {
																		ifres616 = False

																	}

																	if True == ifres616 {
																		tmp579 := MakeNative(func(__e *ControlFlow) {
																			__case := __e.Get(1)
																			_ = __case
																			tmp609 := PrimTail(p87)

																			tmp610 := PrimHead(tmp609)

																			tmp611 := PrimTail(tmp610)

																			tmp612 := PrimIsPair(tmp611)

																			var ifres602 Obj

																			if True == tmp612 {
																				tmp604 := PrimTail(p87)

																				tmp605 := PrimHead(tmp604)

																				tmp606 := PrimTail(tmp605)

																				tmp607 := Call(__e, PrimNS1Value(symnull_2), tmp606)

																				tmp608 := PrimNot(tmp607)

																				var ifres603 Obj

																				if True == tmp608 {
																					ifres603 = True

																				} else {
																					ifres603 = False

																				}

																				ifres602 = ifres603

																			} else {
																				ifres602 = False

																			}

																			if True == ifres602 {
																				tmp581 := MakeNative(func(__e *ControlFlow) {
																					action := __e.Get(1)
																					_ = action
																					tmp593 := PrimTail(p87)

																					tmp594 := PrimHead(tmp593)

																					tmp595 := PrimTail(tmp594)

																					tmp596 := PrimTail(tmp595)

																					tmp597 := Call(__e, PrimNS1Value(symnull_2), tmp596)

																					if True == tmp597 {
																						tmp590 := PrimTail(p87)

																						tmp591 := PrimTail(tmp590)

																						tmp592 := Call(__e, PrimNS1Value(symnull_2), tmp591)

																						if True == tmp592 {
																							tmp584 := PrimCons(MakeString("no cond match"), Nil)

																							tmp585 := PrimCons(symsimple_1error, tmp584)

																							tmp586 := PrimCons(tmp585, Nil)

																							tmp587 := PrimCons(action, tmp586)

																							tmp588 := PrimCons(__case, tmp587)

																							tmp589 := PrimCons(symif, tmp588)

																							__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp589)
																							return

																						} else {
																							__e.TailApply(cc98)
																							return
																						}

																					} else {
																						__e.TailApply(cc98)
																						return
																					}

																				}, 1)

																				tmp598 := PrimTail(p87)

																				tmp599 := PrimHead(tmp598)

																				tmp600 := PrimTail(tmp599)

																				tmp601 := PrimHead(tmp600)

																				__e.TailApply(tmp581, tmp601)
																				return

																			} else {
																				__e.TailApply(cc98)
																				return
																			}

																		}, 1)

																		tmp613 := PrimTail(p87)

																		tmp614 := PrimHead(tmp613)

																		tmp615 := PrimHead(tmp614)

																		__e.TailApply(tmp579, tmp615)
																		return

																	} else {
																		__e.TailApply(cc98)
																		return
																	}

																} else {
																	__e.TailApply(cc98)
																	return
																}

															} else {
																__e.TailApply(cc98)
																return
															}

														} else {
															__e.TailApply(cc98)
															return
														}

													}, 1)

													__e.TailApply(tmp574, p86)
													return

												}, 1)

												tmp639 := MakeNative(func(__e *ControlFlow) {
													tmp640 := MakeNative(func(__e *ControlFlow) {
														cc99 := __e.Get(1)
														_ = cc99
														tmp641 := MakeNative(func(__e *ControlFlow) {
															env := __e.Get(1)
															_ = env
															tmp642 := MakeNative(func(__e *ControlFlow) {
																n106 := __e.Get(1)
																_ = n106
																tmp688 := PrimIsPair(n106)

																if True == tmp688 {
																	tmp686 := PrimHead(n106)

																	tmp687 := PrimEqual(symcond, tmp686)

																	if True == tmp687 {
																		tmp645 := MakeNative(func(__e *ControlFlow) {
																			n107 := __e.Get(1)
																			_ = n107
																			tmp684 := PrimIsPair(n107)

																			if True == tmp684 {
																				tmp682 := PrimHead(n107)

																				tmp683 := PrimIsPair(tmp682)

																				var ifres677 Obj

																				if True == tmp683 {
																					tmp679 := PrimHead(n107)

																					tmp680 := Call(__e, PrimNS1Value(symnull_2), tmp679)

																					tmp681 := PrimNot(tmp680)

																					var ifres678 Obj

																					if True == tmp681 {
																						ifres678 = True

																					} else {
																						ifres678 = False

																					}

																					ifres677 = ifres678

																				} else {
																					ifres677 = False

																				}

																				if True == ifres677 {
																					tmp648 := MakeNative(func(__e *ControlFlow) {
																						__case := __e.Get(1)
																						_ = __case
																						tmp672 := PrimHead(n107)

																						tmp673 := PrimTail(tmp672)

																						tmp674 := PrimIsPair(tmp673)

																						var ifres666 Obj

																						if True == tmp674 {
																							tmp668 := PrimHead(n107)

																							tmp669 := PrimTail(tmp668)

																							tmp670 := Call(__e, PrimNS1Value(symnull_2), tmp669)

																							tmp671 := PrimNot(tmp670)

																							var ifres667 Obj

																							if True == tmp671 {
																								ifres667 = True

																							} else {
																								ifres667 = False

																							}

																							ifres666 = ifres667

																						} else {
																							ifres666 = False

																						}

																						if True == ifres666 {
																							tmp650 := MakeNative(func(__e *ControlFlow) {
																								action := __e.Get(1)
																								_ = action
																								tmp659 := PrimHead(n107)

																								tmp660 := PrimTail(tmp659)

																								tmp661 := PrimTail(tmp660)

																								tmp662 := Call(__e, PrimNS1Value(symnull_2), tmp661)

																								if True == tmp662 {
																									tmp652 := MakeNative(func(__e *ControlFlow) {
																										more := __e.Get(1)
																										_ = more
																										tmp653 := PrimCons(symcond, more)

																										tmp654 := PrimCons(tmp653, Nil)

																										tmp655 := PrimCons(action, tmp654)

																										tmp656 := PrimCons(__case, tmp655)

																										tmp657 := PrimCons(symif, tmp656)

																										__e.TailApply(PrimNS1Value(symkl_1_6cora), env, tmp657)
																										return

																									}, 1)

																									tmp658 := PrimTail(n107)

																									__e.TailApply(tmp652, tmp658)
																									return

																								} else {
																									__e.TailApply(cc99)
																									return
																								}

																							}, 1)

																							tmp663 := PrimHead(n107)

																							tmp664 := PrimTail(tmp663)

																							tmp665 := PrimHead(tmp664)

																							__e.TailApply(tmp650, tmp665)
																							return

																						} else {
																							__e.TailApply(cc99)
																							return
																						}

																					}, 1)

																					tmp675 := PrimHead(n107)

																					tmp676 := PrimHead(tmp675)

																					__e.TailApply(tmp648, tmp676)
																					return

																				} else {
																					__e.TailApply(cc99)
																					return
																				}

																			} else {
																				__e.TailApply(cc99)
																				return
																			}

																		}, 1)

																		tmp685 := PrimTail(n106)

																		__e.TailApply(tmp645, tmp685)
																		return

																	} else {
																		__e.TailApply(cc99)
																		return
																	}

																} else {
																	__e.TailApply(cc99)
																	return
																}

															}, 1)

															__e.TailApply(tmp642, p87)
															return

														}, 1)

														__e.TailApply(tmp641, p86)
														return

													}, 1)

													tmp689 := MakeNative(func(__e *ControlFlow) {
														tmp690 := MakeNative(func(__e *ControlFlow) {
															cc100 := __e.Get(1)
															_ = cc100
															tmp691 := MakeNative(func(__e *ControlFlow) {
																env := __e.Get(1)
																_ = env
																tmp718 := PrimIsPair(p87)

																var ifres714 Obj

																if True == tmp718 {
																	tmp716 := Call(__e, PrimNS1Value(symnull_2), p87)

																	tmp717 := PrimNot(tmp716)

																	var ifres715 Obj

																	if True == tmp717 {
																		ifres715 = True

																	} else {
																		ifres715 = False

																	}

																	ifres714 = ifres715

																} else {
																	ifres714 = False

																}

																if True == ifres714 {
																	tmp712 := PrimHead(p87)

																	tmp713 := PrimEqual(symfreeze, tmp712)

																	if True == tmp713 {
																		tmp710 := PrimTail(p87)

																		tmp711 := PrimIsPair(tmp710)

																		var ifres705 Obj

																		if True == tmp711 {
																			tmp707 := PrimTail(p87)

																			tmp708 := Call(__e, PrimNS1Value(symnull_2), tmp707)

																			tmp709 := PrimNot(tmp708)

																			var ifres706 Obj

																			if True == tmp709 {
																				ifres706 = True

																			} else {
																				ifres706 = False

																			}

																			ifres705 = ifres706

																		} else {
																			ifres705 = False

																		}

																		if True == ifres705 {
																			tmp695 := MakeNative(func(__e *ControlFlow) {
																				body := __e.Get(1)
																				_ = body
																				tmp700 := PrimTail(p87)

																				tmp701 := PrimTail(tmp700)

																				tmp702 := Call(__e, PrimNS1Value(symnull_2), tmp701)

																				if True == tmp702 {
																					tmp697 := Call(__e, PrimNS1Value(symkl_1_6cora), env, body)

																					tmp698 := PrimCons(tmp697, Nil)

																					tmp699 := PrimCons(Nil, tmp698)

																					__e.Return(PrimCons(symlambda, tmp699))
																					return

																				} else {
																					__e.TailApply(cc100)
																					return
																				}

																			}, 1)

																			tmp703 := PrimTail(p87)

																			tmp704 := PrimHead(tmp703)

																			__e.TailApply(tmp695, tmp704)
																			return

																		} else {
																			__e.TailApply(cc100)
																			return
																		}

																	} else {
																		__e.TailApply(cc100)
																		return
																	}

																} else {
																	__e.TailApply(cc100)
																	return
																}

															}, 1)

															__e.TailApply(tmp691, p86)
															return

														}, 1)

														tmp719 := MakeNative(func(__e *ControlFlow) {
															tmp720 := MakeNative(func(__e *ControlFlow) {
																cc101 := __e.Get(1)
																_ = cc101
																tmp721 := MakeNative(func(__e *ControlFlow) {
																	env := __e.Get(1)
																	_ = env
																	tmp722 := MakeNative(func(__e *ControlFlow) {
																		n105 := __e.Get(1)
																		_ = n105
																		tmp733 := PrimIsPair(n105)

																		if True == tmp733 {
																			tmp724 := MakeNative(func(__e *ControlFlow) {
																				f := __e.Get(1)
																				_ = f
																				tmp725 := MakeNative(func(__e *ControlFlow) {
																					x := __e.Get(1)
																					_ = x
																					tmp730 := PrimIsSymbol(f)

																					if True == tmp730 {
																						tmp727 := Call(__e, PrimNS1Value(symparse_1app), env, f)

																						tmp728 := Call(__e, PrimNS1Value(symkl_1_6cora), env)

																						tmp729 := Call(__e, PrimNS1Value(symmap), tmp728, x)

																						__e.Return(PrimCons(tmp727, tmp729))
																						return

																					} else {
																						__e.TailApply(cc101)
																						return
																					}

																				}, 1)

																				tmp731 := PrimTail(n105)

																				__e.TailApply(tmp725, tmp731)
																				return

																			}, 1)

																			tmp732 := PrimHead(n105)

																			__e.TailApply(tmp724, tmp732)
																			return

																		} else {
																			__e.TailApply(cc101)
																			return
																		}

																	}, 1)

																	__e.TailApply(tmp722, p87)
																	return

																}, 1)

																__e.TailApply(tmp721, p86)
																return

															}, 1)

															tmp734 := MakeNative(func(__e *ControlFlow) {
																tmp735 := MakeNative(func(__e *ControlFlow) {
																	cc102 := __e.Get(1)
																	_ = cc102
																	tmp736 := MakeNative(func(__e *ControlFlow) {
																		env := __e.Get(1)
																		_ = env
																		tmp737 := MakeNative(func(__e *ControlFlow) {
																			n104 := __e.Get(1)
																			_ = n104
																			tmp745 := PrimIsPair(n104)

																			if True == tmp745 {
																				tmp739 := MakeNative(func(__e *ControlFlow) {
																					f := __e.Get(1)
																					_ = f
																					tmp740 := MakeNative(func(__e *ControlFlow) {
																						x := __e.Get(1)
																						_ = x
																						tmp741 := Call(__e, PrimNS1Value(symkl_1_6cora), env)

																						tmp742 := PrimCons(f, x)

																						__e.TailApply(PrimNS1Value(symmap), tmp741, tmp742)
																						return

																					}, 1)

																					tmp743 := PrimTail(n104)

																					__e.TailApply(tmp740, tmp743)
																					return

																				}, 1)

																				tmp744 := PrimHead(n104)

																				__e.TailApply(tmp739, tmp744)
																				return

																			} else {
																				__e.TailApply(cc102)
																				return
																			}

																		}, 1)

																		__e.TailApply(tmp737, p87)
																		return

																	}, 1)

																	__e.TailApply(tmp736, p86)
																	return

																}, 1)

																tmp746 := MakeNative(func(__e *ControlFlow) {
																	tmp747 := MakeNative(func(__e *ControlFlow) {
																		cc103 := __e.Get(1)
																		_ = cc103
																		tmp748 := MakeNative(func(__e *ControlFlow) {
																			env := __e.Get(1)
																			_ = env
																			tmp749 := MakeNative(func(__e *ControlFlow) {
																				x := __e.Get(1)
																				_ = x
																				tmp752 := Call(__e, PrimNS1Value(symelem_2), x, env)

																				if True == tmp752 {
																					__e.Return(x)
																					return
																				} else {
																					tmp751 := PrimCons(x, Nil)

																					__e.Return(PrimCons(symquote, tmp751))
																					return

																				}

																			}, 1)

																			__e.TailApply(tmp749, p87)
																			return

																		}, 1)

																		__e.TailApply(tmp748, p86)
																		return

																	}, 1)

																	tmp753 := MakeNative(func(__e *ControlFlow) {
																		__e.TailApply(PrimNS1Value(symerror), MakeString("no match-help found!"))
																		return
																	}, 0)

																	__e.TailApply(tmp747, tmp753)
																	return

																}, 0)

																__e.TailApply(tmp735, tmp746)
																return

															}, 0)

															__e.TailApply(tmp720, tmp734)
															return

														}, 0)

														__e.TailApply(tmp690, tmp719)
														return

													}, 0)

													__e.TailApply(tmp640, tmp689)
													return

												}, 0)

												__e.TailApply(tmp573, tmp639)
												return

											}, 0)

											__e.TailApply(tmp523, tmp572)
											return

										}, 0)

										__e.TailApply(tmp473, tmp522)
										return

									}, 0)

									__e.TailApply(tmp424, tmp472)
									return

								}, 0)

								__e.TailApply(tmp356, tmp423)
								return

							}, 0)

							__e.TailApply(tmp288, tmp355)
							return

						}, 0)

						__e.TailApply(tmp241, tmp287)
						return

					}, 0)

					__e.TailApply(tmp195, tmp240)
					return

				}, 0)

				__e.TailApply(tmp129, tmp194)
				return

			}, 0)

			__e.TailApply(tmp118, tmp128)
			return

		}, 0)

		__e.TailApply(tmp113, tmp117)
		return

	}, 2)

	tmp754 := PrimNS1Set(symkl_1_6cora, tmp112)

	_ = tmp754

	tmp755 := MakeNative(func(__e *ControlFlow) {
		p108 := __e.Get(1)
		_ = p108
		p109 := __e.Get(2)
		_ = p109
		tmp756 := MakeNative(func(__e *ControlFlow) {
			cc110 := __e.Get(1)
			_ = cc110
			tmp757 := MakeNative(func(__e *ControlFlow) {
				env := __e.Get(1)
				_ = env
				tmp758 := MakeNative(func(__e *ControlFlow) {
					f := __e.Get(1)
					_ = f
					tmp760 := Call(__e, PrimNS1Value(symelem_2), f, env)

					if True == tmp760 {
						__e.Return(f)
						return
					} else {
						__e.TailApply(cc110)
						return
					}

				}, 1)

				__e.TailApply(tmp758, p109)
				return

			}, 1)

			__e.TailApply(tmp757, p108)
			return

		}, 1)

		tmp761 := MakeNative(func(__e *ControlFlow) {
			tmp762 := MakeNative(func(__e *ControlFlow) {
				cc111 := __e.Get(1)
				_ = cc111
				tmp763 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					tmp764 := MakeNative(func(__e *ControlFlow) {
						f := __e.Get(1)
						_ = f
						tmp765 := PrimCons(f, Nil)

						tmp766 := PrimCons(symquote, tmp765)

						tmp767 := PrimCons(tmp766, Nil)

						__e.Return(PrimCons(symns2_1value, tmp767))
						return

					}, 1)

					__e.TailApply(tmp764, p109)
					return

				}, 1)

				__e.TailApply(tmp763, p108)
				return

			}, 1)

			tmp768 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS1Value(symerror), MakeString("no match-help found!"))
				return
			}, 0)

			__e.TailApply(tmp762, tmp768)
			return

		}, 0)

		__e.TailApply(tmp756, tmp761)
		return

	}, 2)

	tmp769 := PrimNS1Set(symparse_1app, tmp755)

	_ = tmp769

	tmp770 := MakeNative(func(__e *ControlFlow) {
		fin := __e.Get(1)
		_ = fin
		fout := __e.Get(2)
		_ = fout
		tmp771 := MakeNative(func(__e *ControlFlow) {
			exprs := __e.Get(1)
			_ = exprs
			tmp772 := MakeNative(func(__e *ControlFlow) {
				res := __e.Get(1)
				_ = res
				__e.TailApply(PrimNS1Value(symwrite_1sexp_1to_1file), fout, res)
				return
			}, 1)

			tmp773 := Call(__e, PrimNS1Value(symkl_1_6cora), Nil)

			tmp774 := Call(__e, PrimNS1Value(symmap), tmp773, exprs)

			tmp775 := PrimCons(symbegin, tmp774)

			__e.TailApply(tmp772, tmp775)
			return

		}, 1)

		tmp776 := Call(__e, PrimNS1Value(symread_1file_1as_1sexp), fin, False)

		__e.TailApply(tmp771, tmp776)
		return

	}, 2)

	__e.Return(PrimNS1Set(symcompile_1file, tmp770))
	return

}, 0)

var symbegin = MakeSymbol("begin")
var symkl_1_6cora = MakeSymbol("kl->cora")
var symquote = MakeSymbol("quote")
var symcompile_1file = MakeSymbol("compile-file")
var symparse_1app = MakeSymbol("parse-app")
var symns2_1value = MakeSymbol("ns2-value")
var symread_1file_1as_1sexp = MakeSymbol("read-file-as-sexp")
var symwrite_1sexp_1to_1file = MakeSymbol("write-sexp-to-file")
var symnull_2 = MakeSymbol("null?")
var symelem_2 = MakeSymbol("elem?")
