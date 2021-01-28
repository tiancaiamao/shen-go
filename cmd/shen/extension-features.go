package main

import . "github.com/tiancaiamao/shen-go/kl"

var ExtensionFeaturesMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

	tmp4280 := MakeNative(func(__e *ControlFlow) {
		V4885 := __e.Get(1)
		_ = V4885
		tmp4584 := PrimIsPair(V4885)

		var ifres4576 Obj

		if True == tmp4584 {
			tmp4582 := PrimHead(V4885)

			tmp4583 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4582)

			var ifres4578 Obj

			if True == tmp4583 {
				tmp4580 := PrimTail(V4885)

				tmp4581 := PrimEqual(Nil, tmp4580)

				var ifres4579 Obj

				if True == tmp4581 {
					ifres4579 = True

				} else {
					ifres4579 = False

				}

				ifres4578 = ifres4579

			} else {
				ifres4578 = False

			}

			var ifres4577 Obj

			if True == ifres4578 {
				ifres4577 = True

			} else {
				ifres4577 = False

			}

			ifres4576 = ifres4577

		} else {
			ifres4576 = False

		}

		if True == ifres4576 {
			__e.Return(PrimSimpleError(MakeString("Unfulfilled shen.x.features.cond-expand clause.")))
			return
		} else {
			tmp4575 := PrimIsPair(V4885)

			var ifres4551 Obj

			if True == tmp4575 {
				tmp4573 := PrimHead(V4885)

				tmp4574 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4573)

				var ifres4553 Obj

				if True == tmp4574 {
					tmp4571 := PrimTail(V4885)

					tmp4572 := PrimIsPair(tmp4571)

					var ifres4555 Obj

					if True == tmp4572 {
						tmp4568 := PrimTail(V4885)

						tmp4569 := PrimHead(tmp4568)

						tmp4570 := PrimEqual(True, tmp4569)

						var ifres4557 Obj

						if True == tmp4570 {
							tmp4565 := PrimTail(V4885)

							tmp4566 := PrimTail(tmp4565)

							tmp4567 := PrimIsPair(tmp4566)

							var ifres4559 Obj

							if True == tmp4567 {
								tmp4561 := PrimTail(V4885)

								tmp4562 := PrimTail(tmp4561)

								tmp4563 := PrimTail(tmp4562)

								tmp4564 := PrimEqual(Nil, tmp4563)

								var ifres4560 Obj

								if True == tmp4564 {
									ifres4560 = True

								} else {
									ifres4560 = False

								}

								ifres4559 = ifres4560

							} else {
								ifres4559 = False

							}

							var ifres4558 Obj

							if True == ifres4559 {
								ifres4558 = True

							} else {
								ifres4558 = False

							}

							ifres4557 = ifres4558

						} else {
							ifres4557 = False

						}

						var ifres4556 Obj

						if True == ifres4557 {
							ifres4556 = True

						} else {
							ifres4556 = False

						}

						ifres4555 = ifres4556

					} else {
						ifres4555 = False

					}

					var ifres4554 Obj

					if True == ifres4555 {
						ifres4554 = True

					} else {
						ifres4554 = False

					}

					ifres4553 = ifres4554

				} else {
					ifres4553 = False

				}

				var ifres4552 Obj

				if True == ifres4553 {
					ifres4552 = True

				} else {
					ifres4552 = False

				}

				ifres4551 = ifres4552

			} else {
				ifres4551 = False

			}

			if True == ifres4551 {
				tmp4549 := PrimTail(V4885)

				tmp4550 := PrimTail(tmp4549)

				__e.Return(PrimHead(tmp4550))
				return

			} else {
				tmp4548 := PrimIsPair(V4885)

				var ifres4518 Obj

				if True == tmp4548 {
					tmp4546 := PrimHead(V4885)

					tmp4547 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4546)

					var ifres4520 Obj

					if True == tmp4547 {
						tmp4544 := PrimTail(V4885)

						tmp4545 := PrimIsPair(tmp4544)

						var ifres4522 Obj

						if True == tmp4545 {
							tmp4541 := PrimTail(V4885)

							tmp4542 := PrimHead(tmp4541)

							tmp4543 := PrimIsPair(tmp4542)

							var ifres4524 Obj

							if True == tmp4543 {
								tmp4537 := PrimTail(V4885)

								tmp4538 := PrimHead(tmp4537)

								tmp4539 := PrimHead(tmp4538)

								tmp4540 := PrimEqual(symand, tmp4539)

								var ifres4526 Obj

								if True == tmp4540 {
									tmp4533 := PrimTail(V4885)

									tmp4534 := PrimHead(tmp4533)

									tmp4535 := PrimTail(tmp4534)

									tmp4536 := PrimEqual(Nil, tmp4535)

									var ifres4528 Obj

									if True == tmp4536 {
										tmp4530 := PrimTail(V4885)

										tmp4531 := PrimTail(tmp4530)

										tmp4532 := PrimIsPair(tmp4531)

										var ifres4529 Obj

										if True == tmp4532 {
											ifres4529 = True

										} else {
											ifres4529 = False

										}

										ifres4528 = ifres4529

									} else {
										ifres4528 = False

									}

									var ifres4527 Obj

									if True == ifres4528 {
										ifres4527 = True

									} else {
										ifres4527 = False

									}

									ifres4526 = ifres4527

								} else {
									ifres4526 = False

								}

								var ifres4525 Obj

								if True == ifres4526 {
									ifres4525 = True

								} else {
									ifres4525 = False

								}

								ifres4524 = ifres4525

							} else {
								ifres4524 = False

							}

							var ifres4523 Obj

							if True == ifres4524 {
								ifres4523 = True

							} else {
								ifres4523 = False

							}

							ifres4522 = ifres4523

						} else {
							ifres4522 = False

						}

						var ifres4521 Obj

						if True == ifres4522 {
							ifres4521 = True

						} else {
							ifres4521 = False

						}

						ifres4520 = ifres4521

					} else {
						ifres4520 = False

					}

					var ifres4519 Obj

					if True == ifres4520 {
						ifres4519 = True

					} else {
						ifres4519 = False

					}

					ifres4518 = ifres4519

				} else {
					ifres4518 = False

				}

				if True == ifres4518 {
					tmp4516 := PrimTail(V4885)

					tmp4517 := PrimTail(tmp4516)

					__e.Return(PrimHead(tmp4517))
					return

				} else {
					tmp4515 := PrimIsPair(V4885)

					var ifres4485 Obj

					if True == tmp4515 {
						tmp4513 := PrimHead(V4885)

						tmp4514 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4513)

						var ifres4487 Obj

						if True == tmp4514 {
							tmp4511 := PrimTail(V4885)

							tmp4512 := PrimIsPair(tmp4511)

							var ifres4489 Obj

							if True == tmp4512 {
								tmp4508 := PrimTail(V4885)

								tmp4509 := PrimHead(tmp4508)

								tmp4510 := PrimIsPair(tmp4509)

								var ifres4491 Obj

								if True == tmp4510 {
									tmp4504 := PrimTail(V4885)

									tmp4505 := PrimHead(tmp4504)

									tmp4506 := PrimHead(tmp4505)

									tmp4507 := PrimEqual(symand, tmp4506)

									var ifres4493 Obj

									if True == tmp4507 {
										tmp4500 := PrimTail(V4885)

										tmp4501 := PrimHead(tmp4500)

										tmp4502 := PrimTail(tmp4501)

										tmp4503 := PrimIsPair(tmp4502)

										var ifres4495 Obj

										if True == tmp4503 {
											tmp4497 := PrimTail(V4885)

											tmp4498 := PrimTail(tmp4497)

											tmp4499 := PrimIsPair(tmp4498)

											var ifres4496 Obj

											if True == tmp4499 {
												ifres4496 = True

											} else {
												ifres4496 = False

											}

											ifres4495 = ifres4496

										} else {
											ifres4495 = False

										}

										var ifres4494 Obj

										if True == ifres4495 {
											ifres4494 = True

										} else {
											ifres4494 = False

										}

										ifres4493 = ifres4494

									} else {
										ifres4493 = False

									}

									var ifres4492 Obj

									if True == ifres4493 {
										ifres4492 = True

									} else {
										ifres4492 = False

									}

									ifres4491 = ifres4492

								} else {
									ifres4491 = False

								}

								var ifres4490 Obj

								if True == ifres4491 {
									ifres4490 = True

								} else {
									ifres4490 = False

								}

								ifres4489 = ifres4490

							} else {
								ifres4489 = False

							}

							var ifres4488 Obj

							if True == ifres4489 {
								ifres4488 = True

							} else {
								ifres4488 = False

							}

							ifres4487 = ifres4488

						} else {
							ifres4487 = False

						}

						var ifres4486 Obj

						if True == ifres4487 {
							ifres4486 = True

						} else {
							ifres4486 = False

						}

						ifres4485 = ifres4486

					} else {
						ifres4485 = False

					}

					if True == ifres4485 {
						tmp4467 := PrimTail(V4885)

						tmp4468 := PrimHead(tmp4467)

						tmp4469 := PrimTail(tmp4468)

						tmp4470 := PrimHead(tmp4469)

						tmp4471 := PrimTail(V4885)

						tmp4472 := PrimHead(tmp4471)

						tmp4473 := PrimTail(tmp4472)

						tmp4474 := PrimTail(tmp4473)

						tmp4475 := PrimCons(symand, tmp4474)

						tmp4476 := PrimTail(V4885)

						tmp4477 := PrimTail(tmp4476)

						tmp4478 := PrimCons(tmp4475, tmp4477)

						tmp4479 := PrimCons(symshen_4x_4features_4cond_1expand, tmp4478)

						tmp4480 := PrimTail(V4885)

						tmp4481 := PrimTail(tmp4480)

						tmp4482 := PrimTail(tmp4481)

						tmp4483 := PrimCons(tmp4479, tmp4482)

						tmp4484 := PrimCons(tmp4470, tmp4483)

						__e.Return(PrimCons(symshen_4x_4features_4cond_1expand, tmp4484))
						return

					} else {
						tmp4466 := PrimIsPair(V4885)

						var ifres4436 Obj

						if True == tmp4466 {
							tmp4464 := PrimHead(V4885)

							tmp4465 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4464)

							var ifres4438 Obj

							if True == tmp4465 {
								tmp4462 := PrimTail(V4885)

								tmp4463 := PrimIsPair(tmp4462)

								var ifres4440 Obj

								if True == tmp4463 {
									tmp4459 := PrimTail(V4885)

									tmp4460 := PrimHead(tmp4459)

									tmp4461 := PrimIsPair(tmp4460)

									var ifres4442 Obj

									if True == tmp4461 {
										tmp4455 := PrimTail(V4885)

										tmp4456 := PrimHead(tmp4455)

										tmp4457 := PrimHead(tmp4456)

										tmp4458 := PrimEqual(symor, tmp4457)

										var ifres4444 Obj

										if True == tmp4458 {
											tmp4451 := PrimTail(V4885)

											tmp4452 := PrimHead(tmp4451)

											tmp4453 := PrimTail(tmp4452)

											tmp4454 := PrimEqual(Nil, tmp4453)

											var ifres4446 Obj

											if True == tmp4454 {
												tmp4448 := PrimTail(V4885)

												tmp4449 := PrimTail(tmp4448)

												tmp4450 := PrimIsPair(tmp4449)

												var ifres4447 Obj

												if True == tmp4450 {
													ifres4447 = True

												} else {
													ifres4447 = False

												}

												ifres4446 = ifres4447

											} else {
												ifres4446 = False

											}

											var ifres4445 Obj

											if True == ifres4446 {
												ifres4445 = True

											} else {
												ifres4445 = False

											}

											ifres4444 = ifres4445

										} else {
											ifres4444 = False

										}

										var ifres4443 Obj

										if True == ifres4444 {
											ifres4443 = True

										} else {
											ifres4443 = False

										}

										ifres4442 = ifres4443

									} else {
										ifres4442 = False

									}

									var ifres4441 Obj

									if True == ifres4442 {
										ifres4441 = True

									} else {
										ifres4441 = False

									}

									ifres4440 = ifres4441

								} else {
									ifres4440 = False

								}

								var ifres4439 Obj

								if True == ifres4440 {
									ifres4439 = True

								} else {
									ifres4439 = False

								}

								ifres4438 = ifres4439

							} else {
								ifres4438 = False

							}

							var ifres4437 Obj

							if True == ifres4438 {
								ifres4437 = True

							} else {
								ifres4437 = False

							}

							ifres4436 = ifres4437

						} else {
							ifres4436 = False

						}

						if True == ifres4436 {
							tmp4433 := PrimTail(V4885)

							tmp4434 := PrimTail(tmp4433)

							tmp4435 := PrimTail(tmp4434)

							__e.Return(PrimCons(symshen_4x_4features_4cond_1expand, tmp4435))
							return

						} else {
							tmp4432 := PrimIsPair(V4885)

							var ifres4402 Obj

							if True == tmp4432 {
								tmp4430 := PrimHead(V4885)

								tmp4431 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4430)

								var ifres4404 Obj

								if True == tmp4431 {
									tmp4428 := PrimTail(V4885)

									tmp4429 := PrimIsPair(tmp4428)

									var ifres4406 Obj

									if True == tmp4429 {
										tmp4425 := PrimTail(V4885)

										tmp4426 := PrimHead(tmp4425)

										tmp4427 := PrimIsPair(tmp4426)

										var ifres4408 Obj

										if True == tmp4427 {
											tmp4421 := PrimTail(V4885)

											tmp4422 := PrimHead(tmp4421)

											tmp4423 := PrimHead(tmp4422)

											tmp4424 := PrimEqual(symor, tmp4423)

											var ifres4410 Obj

											if True == tmp4424 {
												tmp4417 := PrimTail(V4885)

												tmp4418 := PrimHead(tmp4417)

												tmp4419 := PrimTail(tmp4418)

												tmp4420 := PrimIsPair(tmp4419)

												var ifres4412 Obj

												if True == tmp4420 {
													tmp4414 := PrimTail(V4885)

													tmp4415 := PrimTail(tmp4414)

													tmp4416 := PrimIsPair(tmp4415)

													var ifres4413 Obj

													if True == tmp4416 {
														ifres4413 = True

													} else {
														ifres4413 = False

													}

													ifres4412 = ifres4413

												} else {
													ifres4412 = False

												}

												var ifres4411 Obj

												if True == ifres4412 {
													ifres4411 = True

												} else {
													ifres4411 = False

												}

												ifres4410 = ifres4411

											} else {
												ifres4410 = False

											}

											var ifres4409 Obj

											if True == ifres4410 {
												ifres4409 = True

											} else {
												ifres4409 = False

											}

											ifres4408 = ifres4409

										} else {
											ifres4408 = False

										}

										var ifres4407 Obj

										if True == ifres4408 {
											ifres4407 = True

										} else {
											ifres4407 = False

										}

										ifres4406 = ifres4407

									} else {
										ifres4406 = False

									}

									var ifres4405 Obj

									if True == ifres4406 {
										ifres4405 = True

									} else {
										ifres4405 = False

									}

									ifres4404 = ifres4405

								} else {
									ifres4404 = False

								}

								var ifres4403 Obj

								if True == ifres4404 {
									ifres4403 = True

								} else {
									ifres4403 = False

								}

								ifres4402 = ifres4403

							} else {
								ifres4402 = False

							}

							if True == ifres4402 {
								tmp4382 := PrimTail(V4885)

								tmp4383 := PrimHead(tmp4382)

								tmp4384 := PrimTail(tmp4383)

								tmp4385 := PrimHead(tmp4384)

								tmp4386 := PrimTail(V4885)

								tmp4387 := PrimTail(tmp4386)

								tmp4388 := PrimHead(tmp4387)

								tmp4389 := PrimTail(V4885)

								tmp4390 := PrimHead(tmp4389)

								tmp4391 := PrimTail(tmp4390)

								tmp4392 := PrimTail(tmp4391)

								tmp4393 := PrimCons(symor, tmp4392)

								tmp4394 := PrimTail(V4885)

								tmp4395 := PrimTail(tmp4394)

								tmp4396 := PrimCons(tmp4393, tmp4395)

								tmp4397 := PrimCons(symshen_4x_4features_4cond_1expand, tmp4396)

								tmp4398 := PrimCons(tmp4397, Nil)

								tmp4399 := PrimCons(True, tmp4398)

								tmp4400 := PrimCons(tmp4388, tmp4399)

								tmp4401 := PrimCons(tmp4385, tmp4400)

								__e.Return(PrimCons(symshen_4x_4features_4cond_1expand, tmp4401))
								return

							} else {
								tmp4381 := PrimIsPair(V4885)

								var ifres4344 Obj

								if True == tmp4381 {
									tmp4379 := PrimHead(V4885)

									tmp4380 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4379)

									var ifres4346 Obj

									if True == tmp4380 {
										tmp4377 := PrimTail(V4885)

										tmp4378 := PrimIsPair(tmp4377)

										var ifres4348 Obj

										if True == tmp4378 {
											tmp4374 := PrimTail(V4885)

											tmp4375 := PrimHead(tmp4374)

											tmp4376 := PrimIsPair(tmp4375)

											var ifres4350 Obj

											if True == tmp4376 {
												tmp4370 := PrimTail(V4885)

												tmp4371 := PrimHead(tmp4370)

												tmp4372 := PrimHead(tmp4371)

												tmp4373 := PrimEqual(symnot, tmp4372)

												var ifres4352 Obj

												if True == tmp4373 {
													tmp4366 := PrimTail(V4885)

													tmp4367 := PrimHead(tmp4366)

													tmp4368 := PrimTail(tmp4367)

													tmp4369 := PrimIsPair(tmp4368)

													var ifres4354 Obj

													if True == tmp4369 {
														tmp4361 := PrimTail(V4885)

														tmp4362 := PrimHead(tmp4361)

														tmp4363 := PrimTail(tmp4362)

														tmp4364 := PrimTail(tmp4363)

														tmp4365 := PrimEqual(Nil, tmp4364)

														var ifres4356 Obj

														if True == tmp4365 {
															tmp4358 := PrimTail(V4885)

															tmp4359 := PrimTail(tmp4358)

															tmp4360 := PrimIsPair(tmp4359)

															var ifres4357 Obj

															if True == tmp4360 {
																ifres4357 = True

															} else {
																ifres4357 = False

															}

															ifres4356 = ifres4357

														} else {
															ifres4356 = False

														}

														var ifres4355 Obj

														if True == ifres4356 {
															ifres4355 = True

														} else {
															ifres4355 = False

														}

														ifres4354 = ifres4355

													} else {
														ifres4354 = False

													}

													var ifres4353 Obj

													if True == ifres4354 {
														ifres4353 = True

													} else {
														ifres4353 = False

													}

													ifres4352 = ifres4353

												} else {
													ifres4352 = False

												}

												var ifres4351 Obj

												if True == ifres4352 {
													ifres4351 = True

												} else {
													ifres4351 = False

												}

												ifres4350 = ifres4351

											} else {
												ifres4350 = False

											}

											var ifres4349 Obj

											if True == ifres4350 {
												ifres4349 = True

											} else {
												ifres4349 = False

											}

											ifres4348 = ifres4349

										} else {
											ifres4348 = False

										}

										var ifres4347 Obj

										if True == ifres4348 {
											ifres4347 = True

										} else {
											ifres4347 = False

										}

										ifres4346 = ifres4347

									} else {
										ifres4346 = False

									}

									var ifres4345 Obj

									if True == ifres4346 {
										ifres4345 = True

									} else {
										ifres4345 = False

									}

									ifres4344 = ifres4345

								} else {
									ifres4344 = False

								}

								if True == ifres4344 {
									tmp4329 := PrimTail(V4885)

									tmp4330 := PrimHead(tmp4329)

									tmp4331 := PrimTail(tmp4330)

									tmp4332 := PrimHead(tmp4331)

									tmp4333 := PrimTail(V4885)

									tmp4334 := PrimTail(tmp4333)

									tmp4335 := PrimTail(tmp4334)

									tmp4336 := PrimCons(symshen_4x_4features_4cond_1expand, tmp4335)

									tmp4337 := PrimTail(V4885)

									tmp4338 := PrimTail(tmp4337)

									tmp4339 := PrimHead(tmp4338)

									tmp4340 := PrimCons(tmp4339, Nil)

									tmp4341 := PrimCons(True, tmp4340)

									tmp4342 := PrimCons(tmp4336, tmp4341)

									tmp4343 := PrimCons(tmp4332, tmp4342)

									__e.Return(PrimCons(symshen_4x_4features_4cond_1expand, tmp4343))
									return

								} else {
									tmp4328 := PrimIsPair(V4885)

									var ifres4309 Obj

									if True == tmp4328 {
										tmp4326 := PrimHead(V4885)

										tmp4327 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4326)

										var ifres4311 Obj

										if True == tmp4327 {
											tmp4324 := PrimTail(V4885)

											tmp4325 := PrimIsPair(tmp4324)

											var ifres4313 Obj

											if True == tmp4325 {
												tmp4321 := PrimTail(V4885)

												tmp4322 := PrimTail(tmp4321)

												tmp4323 := PrimIsPair(tmp4322)

												var ifres4315 Obj

												if True == tmp4323 {
													tmp4317 := PrimTail(V4885)

													tmp4318 := PrimHead(tmp4317)

													tmp4319 := PrimNS3Value(symshen_4x_4features_4_dfeatures_d)

													tmp4320 := Call(__e, PrimNS2Value(symelement_2), tmp4318, tmp4319)

													var ifres4316 Obj

													if True == tmp4320 {
														ifres4316 = True

													} else {
														ifres4316 = False

													}

													ifres4315 = ifres4316

												} else {
													ifres4315 = False

												}

												var ifres4314 Obj

												if True == ifres4315 {
													ifres4314 = True

												} else {
													ifres4314 = False

												}

												ifres4313 = ifres4314

											} else {
												ifres4313 = False

											}

											var ifres4312 Obj

											if True == ifres4313 {
												ifres4312 = True

											} else {
												ifres4312 = False

											}

											ifres4311 = ifres4312

										} else {
											ifres4311 = False

										}

										var ifres4310 Obj

										if True == ifres4311 {
											ifres4310 = True

										} else {
											ifres4310 = False

										}

										ifres4309 = ifres4310

									} else {
										ifres4309 = False

									}

									if True == ifres4309 {
										tmp4307 := PrimTail(V4885)

										tmp4308 := PrimTail(tmp4307)

										__e.Return(PrimHead(tmp4308))
										return

									} else {
										tmp4306 := PrimIsPair(V4885)

										var ifres4293 Obj

										if True == tmp4306 {
											tmp4304 := PrimHead(V4885)

											tmp4305 := PrimEqual(symshen_4x_4features_4cond_1expand, tmp4304)

											var ifres4295 Obj

											if True == tmp4305 {
												tmp4302 := PrimTail(V4885)

												tmp4303 := PrimIsPair(tmp4302)

												var ifres4297 Obj

												if True == tmp4303 {
													tmp4299 := PrimTail(V4885)

													tmp4300 := PrimTail(tmp4299)

													tmp4301 := PrimIsPair(tmp4300)

													var ifres4298 Obj

													if True == tmp4301 {
														ifres4298 = True

													} else {
														ifres4298 = False

													}

													ifres4297 = ifres4298

												} else {
													ifres4297 = False

												}

												var ifres4296 Obj

												if True == ifres4297 {
													ifres4296 = True

												} else {
													ifres4296 = False

												}

												ifres4295 = ifres4296

											} else {
												ifres4295 = False

											}

											var ifres4294 Obj

											if True == ifres4295 {
												ifres4294 = True

											} else {
												ifres4294 = False

											}

											ifres4293 = ifres4294

										} else {
											ifres4293 = False

										}

										if True == ifres4293 {
											tmp4290 := PrimTail(V4885)

											tmp4291 := PrimTail(tmp4290)

											tmp4292 := PrimTail(tmp4291)

											__e.Return(PrimCons(symshen_4x_4features_4cond_1expand, tmp4292))
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

	tmp4585 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4cond_1expand_1macro, tmp4280)

	_ = tmp4585

	tmp4586 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4x_4features_4_dfeatures_d))
		return
	}, 0)

	tmp4587 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4current, tmp4586)

	_ = tmp4587

	tmp4588 := MakeNative(func(__e *ControlFlow) {
		V4887 := __e.Get(1)
		_ = V4887
		tmp4589 := MakeNative(func(__e *ControlFlow) {
			__ := __e.Get(1)
			_ = __
			tmp4590 := MakeNative(func(__e *ControlFlow) {
				Old := __e.Get(1)
				_ = Old
				tmp4591 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					__e.Return(Old)
					return
				}, 1)

				tmp4592 := PrimNS3Set(symshen_4x_4features_4_dfeatures_d, V4887)

				__e.TailApply(tmp4591, tmp4592)
				return

			}, 1)

			tmp4593 := Call(__e, PrimNS2Value(symshen_4x_4features_4current))

			__e.TailApply(tmp4590, tmp4593)
			return

		}, 1)

		tmp4594 := MakeNative(func(__e *ControlFlow) {
			__e.Return(PrimNS3Value(symshen_4x_4features_4_dfeatures_d))
			return
		}, 0)

		tmp4595 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp4596 := PrimNS3Set(symshen_4x_4features_4_dfeatures_d, Nil)

			_ = tmp4596

			tmp4597 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4x_4features_4cond_1expand_1macro), X)
				return
			}, 1)

			tmp4598 := PrimCons(symshen_4x_4features_4cond_1expand_1macro, tmp4597)

			tmp4599 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp4598)

			_ = tmp4599

			__e.TailApply(PrimNS2Value(symshen_4add_1macro), symshen_4x_4features_4cond_1expand_1macro)
			return

		}, 1)

		tmp4600 := Call(__e, PrimNS1Value(symtry_1catch), tmp4594, tmp4595)

		__e.TailApply(tmp4589, tmp4600)
		return

	}, 1)

	tmp4601 := Call(__e, PrimNS1Value(symns2_1set), symshen_4x_4features_4initialise, tmp4588)

	_ = tmp4601

	tmp4602 := MakeNative(func(__e *ControlFlow) {
		V4889 := __e.Get(1)
		_ = V4889
		tmp4603 := MakeNative(func(__e *ControlFlow) {
			Old := __e.Get(1)
			_ = Old
			tmp4604 := MakeNative(func(__e *ControlFlow) {
				__ := __e.Get(1)
				_ = __
				__e.Return(Old)
				return
			}, 1)

			tmp4605 := Call(__e, PrimNS2Value(symadjoin), V4889, Old)

			tmp4606 := PrimNS3Set(symshen_4x_4features_4_dfeatures_d, tmp4605)

			__e.TailApply(tmp4604, tmp4606)
			return

		}, 1)

		tmp4607 := Call(__e, PrimNS2Value(symshen_4x_4features_4current))

		__e.TailApply(tmp4603, tmp4607)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4x_4features_4add, tmp4602)
	return

}, 0)
