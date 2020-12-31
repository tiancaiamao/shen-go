package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

		gen17374 := MakeNative(func(__e Evaluator) {
			V4885 := __e.Get(1)
			_ = V4885
			gen17372 := Call(__e, __e.Global(symcons_2), V4885)

			var gen17373 Obj
			if True == gen17372 {
				gen17369 := Call(__e, __e.Global(symhd), V4885)

				gen17370 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17369)

				var gen17371 Obj
				if True == gen17370 {
					gen17367 := Call(__e, __e.Global(symtl), V4885)

					gen17368 := Call(__e, __e.Global(sym_a), Nil, gen17367)

					if True == gen17368 {
						gen17371 = True
					} else {
						gen17371 = False
					}

				} else {
					gen17371 = False
				}
				if True == gen17371 {
					gen17373 = True
				} else {
					gen17373 = False
				}

			} else {
				gen17373 = False
			}
			if True == gen17373 {
				__e.TailApply(__e.Global(symsimple_1error), MakeString("Unfulfilled shen.x.features.cond-expand clause."))

				return
			} else {
				gen17365 := Call(__e, __e.Global(symcons_2), V4885)

				var gen17366 Obj
				if True == gen17365 {
					gen17362 := Call(__e, __e.Global(symhd), V4885)

					gen17363 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17362)

					var gen17364 Obj
					if True == gen17363 {
						gen17359 := Call(__e, __e.Global(symtl), V4885)

						gen17360 := Call(__e, __e.Global(symcons_2), gen17359)

						var gen17361 Obj
						if True == gen17360 {
							gen17355 := Call(__e, __e.Global(symtl), V4885)

							gen17356 := Call(__e, __e.Global(symhd), gen17355)

							gen17357 := Call(__e, __e.Global(sym_a), True, gen17356)

							var gen17358 Obj
							if True == gen17357 {
								gen17351 := Call(__e, __e.Global(symtl), V4885)

								gen17352 := Call(__e, __e.Global(symtl), gen17351)

								gen17353 := Call(__e, __e.Global(symcons_2), gen17352)

								var gen17354 Obj
								if True == gen17353 {
									gen17347 := Call(__e, __e.Global(symtl), V4885)

									gen17348 := Call(__e, __e.Global(symtl), gen17347)

									gen17349 := Call(__e, __e.Global(symtl), gen17348)

									gen17350 := Call(__e, __e.Global(sym_a), Nil, gen17349)

									if True == gen17350 {
										gen17354 = True
									} else {
										gen17354 = False
									}

								} else {
									gen17354 = False
								}
								if True == gen17354 {
									gen17358 = True
								} else {
									gen17358 = False
								}

							} else {
								gen17358 = False
							}
							if True == gen17358 {
								gen17361 = True
							} else {
								gen17361 = False
							}

						} else {
							gen17361 = False
						}
						if True == gen17361 {
							gen17364 = True
						} else {
							gen17364 = False
						}

					} else {
						gen17364 = False
					}
					if True == gen17364 {
						gen17366 = True
					} else {
						gen17366 = False
					}

				} else {
					gen17366 = False
				}
				if True == gen17366 {
					gen17345 := Call(__e, __e.Global(symtl), V4885)

					gen17346 := Call(__e, __e.Global(symtl), gen17345)

					__e.TailApply(__e.Global(symhd), gen17346)

					return

				} else {
					gen17343 := Call(__e, __e.Global(symcons_2), V4885)

					var gen17344 Obj
					if True == gen17343 {
						gen17340 := Call(__e, __e.Global(symhd), V4885)

						gen17341 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17340)

						var gen17342 Obj
						if True == gen17341 {
							gen17337 := Call(__e, __e.Global(symtl), V4885)

							gen17338 := Call(__e, __e.Global(symcons_2), gen17337)

							var gen17339 Obj
							if True == gen17338 {
								gen17333 := Call(__e, __e.Global(symtl), V4885)

								gen17334 := Call(__e, __e.Global(symhd), gen17333)

								gen17335 := Call(__e, __e.Global(symcons_2), gen17334)

								var gen17336 Obj
								if True == gen17335 {
									gen17328 := Call(__e, __e.Global(symtl), V4885)

									gen17329 := Call(__e, __e.Global(symhd), gen17328)

									gen17330 := Call(__e, __e.Global(symhd), gen17329)

									gen17331 := Call(__e, __e.Global(sym_a), MakeSymbol("and"), gen17330)

									var gen17332 Obj
									if True == gen17331 {
										gen17323 := Call(__e, __e.Global(symtl), V4885)

										gen17324 := Call(__e, __e.Global(symhd), gen17323)

										gen17325 := Call(__e, __e.Global(symtl), gen17324)

										gen17326 := Call(__e, __e.Global(sym_a), Nil, gen17325)

										var gen17327 Obj
										if True == gen17326 {
											gen17320 := Call(__e, __e.Global(symtl), V4885)

											gen17321 := Call(__e, __e.Global(symtl), gen17320)

											gen17322 := Call(__e, __e.Global(symcons_2), gen17321)

											if True == gen17322 {
												gen17327 = True
											} else {
												gen17327 = False
											}

										} else {
											gen17327 = False
										}
										if True == gen17327 {
											gen17332 = True
										} else {
											gen17332 = False
										}

									} else {
										gen17332 = False
									}
									if True == gen17332 {
										gen17336 = True
									} else {
										gen17336 = False
									}

								} else {
									gen17336 = False
								}
								if True == gen17336 {
									gen17339 = True
								} else {
									gen17339 = False
								}

							} else {
								gen17339 = False
							}
							if True == gen17339 {
								gen17342 = True
							} else {
								gen17342 = False
							}

						} else {
							gen17342 = False
						}
						if True == gen17342 {
							gen17344 = True
						} else {
							gen17344 = False
						}

					} else {
						gen17344 = False
					}
					if True == gen17344 {
						gen17318 := Call(__e, __e.Global(symtl), V4885)

						gen17319 := Call(__e, __e.Global(symtl), gen17318)

						__e.TailApply(__e.Global(symhd), gen17319)

						return

					} else {
						gen17316 := Call(__e, __e.Global(symcons_2), V4885)

						var gen17317 Obj
						if True == gen17316 {
							gen17313 := Call(__e, __e.Global(symhd), V4885)

							gen17314 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17313)

							var gen17315 Obj
							if True == gen17314 {
								gen17310 := Call(__e, __e.Global(symtl), V4885)

								gen17311 := Call(__e, __e.Global(symcons_2), gen17310)

								var gen17312 Obj
								if True == gen17311 {
									gen17306 := Call(__e, __e.Global(symtl), V4885)

									gen17307 := Call(__e, __e.Global(symhd), gen17306)

									gen17308 := Call(__e, __e.Global(symcons_2), gen17307)

									var gen17309 Obj
									if True == gen17308 {
										gen17301 := Call(__e, __e.Global(symtl), V4885)

										gen17302 := Call(__e, __e.Global(symhd), gen17301)

										gen17303 := Call(__e, __e.Global(symhd), gen17302)

										gen17304 := Call(__e, __e.Global(sym_a), MakeSymbol("and"), gen17303)

										var gen17305 Obj
										if True == gen17304 {
											gen17296 := Call(__e, __e.Global(symtl), V4885)

											gen17297 := Call(__e, __e.Global(symhd), gen17296)

											gen17298 := Call(__e, __e.Global(symtl), gen17297)

											gen17299 := Call(__e, __e.Global(symcons_2), gen17298)

											var gen17300 Obj
											if True == gen17299 {
												gen17293 := Call(__e, __e.Global(symtl), V4885)

												gen17294 := Call(__e, __e.Global(symtl), gen17293)

												gen17295 := Call(__e, __e.Global(symcons_2), gen17294)

												if True == gen17295 {
													gen17300 = True
												} else {
													gen17300 = False
												}

											} else {
												gen17300 = False
											}
											if True == gen17300 {
												gen17305 = True
											} else {
												gen17305 = False
											}

										} else {
											gen17305 = False
										}
										if True == gen17305 {
											gen17309 = True
										} else {
											gen17309 = False
										}

									} else {
										gen17309 = False
									}
									if True == gen17309 {
										gen17312 = True
									} else {
										gen17312 = False
									}

								} else {
									gen17312 = False
								}
								if True == gen17312 {
									gen17315 = True
								} else {
									gen17315 = False
								}

							} else {
								gen17315 = False
							}
							if True == gen17315 {
								gen17317 = True
							} else {
								gen17317 = False
							}

						} else {
							gen17317 = False
						}
						if True == gen17317 {
							gen17275 := Call(__e, __e.Global(symtl), V4885)

							gen17276 := Call(__e, __e.Global(symhd), gen17275)

							gen17277 := Call(__e, __e.Global(symtl), gen17276)

							gen17278 := Call(__e, __e.Global(symhd), gen17277)

							gen17279 := Call(__e, __e.Global(symtl), V4885)

							gen17280 := Call(__e, __e.Global(symhd), gen17279)

							gen17281 := Call(__e, __e.Global(symtl), gen17280)

							gen17282 := Call(__e, __e.Global(symtl), gen17281)

							gen17283 := Call(__e, __e.Global(symcons), MakeSymbol("and"), gen17282)

							gen17284 := Call(__e, __e.Global(symtl), V4885)

							gen17285 := Call(__e, __e.Global(symtl), gen17284)

							gen17286 := Call(__e, __e.Global(symcons), gen17283, gen17285)

							gen17287 := Call(__e, __e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17286)

							gen17288 := Call(__e, __e.Global(symtl), V4885)

							gen17289 := Call(__e, __e.Global(symtl), gen17288)

							gen17290 := Call(__e, __e.Global(symtl), gen17289)

							gen17291 := Call(__e, __e.Global(symcons), gen17287, gen17290)

							gen17292 := Call(__e, __e.Global(symcons), gen17278, gen17291)

							__e.TailApply(__e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17292)

							return

						} else {
							gen17273 := Call(__e, __e.Global(symcons_2), V4885)

							var gen17274 Obj
							if True == gen17273 {
								gen17270 := Call(__e, __e.Global(symhd), V4885)

								gen17271 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17270)

								var gen17272 Obj
								if True == gen17271 {
									gen17267 := Call(__e, __e.Global(symtl), V4885)

									gen17268 := Call(__e, __e.Global(symcons_2), gen17267)

									var gen17269 Obj
									if True == gen17268 {
										gen17263 := Call(__e, __e.Global(symtl), V4885)

										gen17264 := Call(__e, __e.Global(symhd), gen17263)

										gen17265 := Call(__e, __e.Global(symcons_2), gen17264)

										var gen17266 Obj
										if True == gen17265 {
											gen17258 := Call(__e, __e.Global(symtl), V4885)

											gen17259 := Call(__e, __e.Global(symhd), gen17258)

											gen17260 := Call(__e, __e.Global(symhd), gen17259)

											gen17261 := Call(__e, __e.Global(sym_a), MakeSymbol("or"), gen17260)

											var gen17262 Obj
											if True == gen17261 {
												gen17253 := Call(__e, __e.Global(symtl), V4885)

												gen17254 := Call(__e, __e.Global(symhd), gen17253)

												gen17255 := Call(__e, __e.Global(symtl), gen17254)

												gen17256 := Call(__e, __e.Global(sym_a), Nil, gen17255)

												var gen17257 Obj
												if True == gen17256 {
													gen17250 := Call(__e, __e.Global(symtl), V4885)

													gen17251 := Call(__e, __e.Global(symtl), gen17250)

													gen17252 := Call(__e, __e.Global(symcons_2), gen17251)

													if True == gen17252 {
														gen17257 = True
													} else {
														gen17257 = False
													}

												} else {
													gen17257 = False
												}
												if True == gen17257 {
													gen17262 = True
												} else {
													gen17262 = False
												}

											} else {
												gen17262 = False
											}
											if True == gen17262 {
												gen17266 = True
											} else {
												gen17266 = False
											}

										} else {
											gen17266 = False
										}
										if True == gen17266 {
											gen17269 = True
										} else {
											gen17269 = False
										}

									} else {
										gen17269 = False
									}
									if True == gen17269 {
										gen17272 = True
									} else {
										gen17272 = False
									}

								} else {
									gen17272 = False
								}
								if True == gen17272 {
									gen17274 = True
								} else {
									gen17274 = False
								}

							} else {
								gen17274 = False
							}
							if True == gen17274 {
								gen17247 := Call(__e, __e.Global(symtl), V4885)

								gen17248 := Call(__e, __e.Global(symtl), gen17247)

								gen17249 := Call(__e, __e.Global(symtl), gen17248)

								__e.TailApply(__e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17249)

								return

							} else {
								gen17245 := Call(__e, __e.Global(symcons_2), V4885)

								var gen17246 Obj
								if True == gen17245 {
									gen17242 := Call(__e, __e.Global(symhd), V4885)

									gen17243 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17242)

									var gen17244 Obj
									if True == gen17243 {
										gen17239 := Call(__e, __e.Global(symtl), V4885)

										gen17240 := Call(__e, __e.Global(symcons_2), gen17239)

										var gen17241 Obj
										if True == gen17240 {
											gen17235 := Call(__e, __e.Global(symtl), V4885)

											gen17236 := Call(__e, __e.Global(symhd), gen17235)

											gen17237 := Call(__e, __e.Global(symcons_2), gen17236)

											var gen17238 Obj
											if True == gen17237 {
												gen17230 := Call(__e, __e.Global(symtl), V4885)

												gen17231 := Call(__e, __e.Global(symhd), gen17230)

												gen17232 := Call(__e, __e.Global(symhd), gen17231)

												gen17233 := Call(__e, __e.Global(sym_a), MakeSymbol("or"), gen17232)

												var gen17234 Obj
												if True == gen17233 {
													gen17225 := Call(__e, __e.Global(symtl), V4885)

													gen17226 := Call(__e, __e.Global(symhd), gen17225)

													gen17227 := Call(__e, __e.Global(symtl), gen17226)

													gen17228 := Call(__e, __e.Global(symcons_2), gen17227)

													var gen17229 Obj
													if True == gen17228 {
														gen17222 := Call(__e, __e.Global(symtl), V4885)

														gen17223 := Call(__e, __e.Global(symtl), gen17222)

														gen17224 := Call(__e, __e.Global(symcons_2), gen17223)

														if True == gen17224 {
															gen17229 = True
														} else {
															gen17229 = False
														}

													} else {
														gen17229 = False
													}
													if True == gen17229 {
														gen17234 = True
													} else {
														gen17234 = False
													}

												} else {
													gen17234 = False
												}
												if True == gen17234 {
													gen17238 = True
												} else {
													gen17238 = False
												}

											} else {
												gen17238 = False
											}
											if True == gen17238 {
												gen17241 = True
											} else {
												gen17241 = False
											}

										} else {
											gen17241 = False
										}
										if True == gen17241 {
											gen17244 = True
										} else {
											gen17244 = False
										}

									} else {
										gen17244 = False
									}
									if True == gen17244 {
										gen17246 = True
									} else {
										gen17246 = False
									}

								} else {
									gen17246 = False
								}
								if True == gen17246 {
									gen17202 := Call(__e, __e.Global(symtl), V4885)

									gen17203 := Call(__e, __e.Global(symhd), gen17202)

									gen17204 := Call(__e, __e.Global(symtl), gen17203)

									gen17205 := Call(__e, __e.Global(symhd), gen17204)

									gen17206 := Call(__e, __e.Global(symtl), V4885)

									gen17207 := Call(__e, __e.Global(symtl), gen17206)

									gen17208 := Call(__e, __e.Global(symhd), gen17207)

									gen17209 := Call(__e, __e.Global(symtl), V4885)

									gen17210 := Call(__e, __e.Global(symhd), gen17209)

									gen17211 := Call(__e, __e.Global(symtl), gen17210)

									gen17212 := Call(__e, __e.Global(symtl), gen17211)

									gen17213 := Call(__e, __e.Global(symcons), MakeSymbol("or"), gen17212)

									gen17214 := Call(__e, __e.Global(symtl), V4885)

									gen17215 := Call(__e, __e.Global(symtl), gen17214)

									gen17216 := Call(__e, __e.Global(symcons), gen17213, gen17215)

									gen17217 := Call(__e, __e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17216)

									gen17218 := Call(__e, __e.Global(symcons), gen17217, Nil)

									gen17219 := Call(__e, __e.Global(symcons), True, gen17218)

									gen17220 := Call(__e, __e.Global(symcons), gen17208, gen17219)

									gen17221 := Call(__e, __e.Global(symcons), gen17205, gen17220)

									__e.TailApply(__e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17221)

									return

								} else {
									gen17200 := Call(__e, __e.Global(symcons_2), V4885)

									var gen17201 Obj
									if True == gen17200 {
										gen17197 := Call(__e, __e.Global(symhd), V4885)

										gen17198 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17197)

										var gen17199 Obj
										if True == gen17198 {
											gen17194 := Call(__e, __e.Global(symtl), V4885)

											gen17195 := Call(__e, __e.Global(symcons_2), gen17194)

											var gen17196 Obj
											if True == gen17195 {
												gen17190 := Call(__e, __e.Global(symtl), V4885)

												gen17191 := Call(__e, __e.Global(symhd), gen17190)

												gen17192 := Call(__e, __e.Global(symcons_2), gen17191)

												var gen17193 Obj
												if True == gen17192 {
													gen17185 := Call(__e, __e.Global(symtl), V4885)

													gen17186 := Call(__e, __e.Global(symhd), gen17185)

													gen17187 := Call(__e, __e.Global(symhd), gen17186)

													gen17188 := Call(__e, __e.Global(sym_a), MakeSymbol("not"), gen17187)

													var gen17189 Obj
													if True == gen17188 {
														gen17180 := Call(__e, __e.Global(symtl), V4885)

														gen17181 := Call(__e, __e.Global(symhd), gen17180)

														gen17182 := Call(__e, __e.Global(symtl), gen17181)

														gen17183 := Call(__e, __e.Global(symcons_2), gen17182)

														var gen17184 Obj
														if True == gen17183 {
															gen17174 := Call(__e, __e.Global(symtl), V4885)

															gen17175 := Call(__e, __e.Global(symhd), gen17174)

															gen17176 := Call(__e, __e.Global(symtl), gen17175)

															gen17177 := Call(__e, __e.Global(symtl), gen17176)

															gen17178 := Call(__e, __e.Global(sym_a), Nil, gen17177)

															var gen17179 Obj
															if True == gen17178 {
																gen17171 := Call(__e, __e.Global(symtl), V4885)

																gen17172 := Call(__e, __e.Global(symtl), gen17171)

																gen17173 := Call(__e, __e.Global(symcons_2), gen17172)

																if True == gen17173 {
																	gen17179 = True
																} else {
																	gen17179 = False
																}

															} else {
																gen17179 = False
															}
															if True == gen17179 {
																gen17184 = True
															} else {
																gen17184 = False
															}

														} else {
															gen17184 = False
														}
														if True == gen17184 {
															gen17189 = True
														} else {
															gen17189 = False
														}

													} else {
														gen17189 = False
													}
													if True == gen17189 {
														gen17193 = True
													} else {
														gen17193 = False
													}

												} else {
													gen17193 = False
												}
												if True == gen17193 {
													gen17196 = True
												} else {
													gen17196 = False
												}

											} else {
												gen17196 = False
											}
											if True == gen17196 {
												gen17199 = True
											} else {
												gen17199 = False
											}

										} else {
											gen17199 = False
										}
										if True == gen17199 {
											gen17201 = True
										} else {
											gen17201 = False
										}

									} else {
										gen17201 = False
									}
									if True == gen17201 {
										gen17156 := Call(__e, __e.Global(symtl), V4885)

										gen17157 := Call(__e, __e.Global(symhd), gen17156)

										gen17158 := Call(__e, __e.Global(symtl), gen17157)

										gen17159 := Call(__e, __e.Global(symhd), gen17158)

										gen17160 := Call(__e, __e.Global(symtl), V4885)

										gen17161 := Call(__e, __e.Global(symtl), gen17160)

										gen17162 := Call(__e, __e.Global(symtl), gen17161)

										gen17163 := Call(__e, __e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17162)

										gen17164 := Call(__e, __e.Global(symtl), V4885)

										gen17165 := Call(__e, __e.Global(symtl), gen17164)

										gen17166 := Call(__e, __e.Global(symhd), gen17165)

										gen17167 := Call(__e, __e.Global(symcons), gen17166, Nil)

										gen17168 := Call(__e, __e.Global(symcons), True, gen17167)

										gen17169 := Call(__e, __e.Global(symcons), gen17163, gen17168)

										gen17170 := Call(__e, __e.Global(symcons), gen17159, gen17169)

										__e.TailApply(__e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17170)

										return

									} else {
										gen17154 := Call(__e, __e.Global(symcons_2), V4885)

										var gen17155 Obj
										if True == gen17154 {
											gen17151 := Call(__e, __e.Global(symhd), V4885)

											gen17152 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17151)

											var gen17153 Obj
											if True == gen17152 {
												gen17148 := Call(__e, __e.Global(symtl), V4885)

												gen17149 := Call(__e, __e.Global(symcons_2), gen17148)

												var gen17150 Obj
												if True == gen17149 {
													gen17144 := Call(__e, __e.Global(symtl), V4885)

													gen17145 := Call(__e, __e.Global(symtl), gen17144)

													gen17146 := Call(__e, __e.Global(symcons_2), gen17145)

													var gen17147 Obj
													if True == gen17146 {
														gen17140 := Call(__e, __e.Global(symtl), V4885)

														gen17141 := Call(__e, __e.Global(symhd), gen17140)

														gen17142 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.features.*features*"))

														gen17143 := Call(__e, __e.Global(symelement_2), gen17141, gen17142)

														if True == gen17143 {
															gen17147 = True
														} else {
															gen17147 = False
														}

													} else {
														gen17147 = False
													}
													if True == gen17147 {
														gen17150 = True
													} else {
														gen17150 = False
													}

												} else {
													gen17150 = False
												}
												if True == gen17150 {
													gen17153 = True
												} else {
													gen17153 = False
												}

											} else {
												gen17153 = False
											}
											if True == gen17153 {
												gen17155 = True
											} else {
												gen17155 = False
											}

										} else {
											gen17155 = False
										}
										if True == gen17155 {
											gen17138 := Call(__e, __e.Global(symtl), V4885)

											gen17139 := Call(__e, __e.Global(symtl), gen17138)

											__e.TailApply(__e.Global(symhd), gen17139)

											return

										} else {
											gen17136 := Call(__e, __e.Global(symcons_2), V4885)

											var gen17137 Obj
											if True == gen17136 {
												gen17133 := Call(__e, __e.Global(symhd), V4885)

												gen17134 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen17133)

												var gen17135 Obj
												if True == gen17134 {
													gen17130 := Call(__e, __e.Global(symtl), V4885)

													gen17131 := Call(__e, __e.Global(symcons_2), gen17130)

													var gen17132 Obj
													if True == gen17131 {
														gen17127 := Call(__e, __e.Global(symtl), V4885)

														gen17128 := Call(__e, __e.Global(symtl), gen17127)

														gen17129 := Call(__e, __e.Global(symcons_2), gen17128)

														if True == gen17129 {
															gen17132 = True
														} else {
															gen17132 = False
														}

													} else {
														gen17132 = False
													}
													if True == gen17132 {
														gen17135 = True
													} else {
														gen17135 = False
													}

												} else {
													gen17135 = False
												}
												if True == gen17135 {
													gen17137 = True
												} else {
													gen17137 = False
												}

											} else {
												gen17137 = False
											}
											if True == gen17137 {
												gen17124 := Call(__e, __e.Global(symtl), V4885)

												gen17125 := Call(__e, __e.Global(symtl), gen17124)

												gen17126 := Call(__e, __e.Global(symtl), gen17125)

												__e.TailApply(__e.Global(symcons), MakeSymbol("shen.x.features.cond-expand"), gen17126)

												return

											} else {
												__e.Return(V4885)
												return
											}

										}

									}

								}

							}

						}

					}

				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.features.cond-expand-macro"), gen17374)

		gen17375 := MakeNative(func(__e Evaluator) {
			__e.TailApply(__e.Global(symvalue), MakeSymbol("shen.x.features.*features*"))

			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.features.current"), gen17375)

		gen17383 := MakeNative(func(__e Evaluator) {
			V4887 := __e.Get(1)
			_ = V4887
			gen17378 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				Call(__e, __e.Global(symset), MakeSymbol("shen.x.features.*features*"), Nil)
				gen17376 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(__e.Global(symshen_4x_4features_4cond_1expand_1macro), X)

					return
				}, 1)
				gen17377 := Call(__e, __e.Global(symcons), MakeSymbol("shen.x.features.cond-expand-macro"), gen17376)

				Call(__e, __e.Global(symshen_4set_1lambda_1form_1entry), gen17377)

				__e.TailApply(__e.Global(symshen_4add_1macro), MakeSymbol("shen.x.features.cond-expand-macro"))

				return

			}, 1)
			gen17379 := MakeNative(func(__e Evaluator) {
				__e.TailApply(__e.Global(symvalue), MakeSymbol("shen.x.features.*features*"))

				return
			}, 0)
			gen17380 := Try(__e, gen17379).Catch(gen17378)
			_ = gen17380
			gen17381 := Call(__e, __e.Global(symshen_4x_4features_4current))

			Old := gen17381
			gen17382 := Call(__e, __e.Global(symset), MakeSymbol("shen.x.features.*features*"), V4887)

			_ = gen17382
			__e.Return(Old)
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.features.initialise"), gen17383)

		gen17387 := MakeNative(func(__e Evaluator) {
			V4889 := __e.Get(1)
			_ = V4889
			gen17384 := Call(__e, __e.Global(symshen_4x_4features_4current))

			Old := gen17384
			gen17385 := Call(__e, __e.Global(symadjoin), V4889, Old)

			gen17386 := Call(__e, __e.Global(symset), MakeSymbol("shen.x.features.*features*"), gen17385)

			_ = gen17386
			__e.Return(Old)
			return

		}, 1)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.x.features.add"), gen17387)

		return

	}, 0))
}
