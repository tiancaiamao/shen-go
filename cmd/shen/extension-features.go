package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

		gen11675 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4885 := __args[0]
			_ = V4885
			gen11673 := Call(__e, ShenFunc(symcons_2), V4885)

			var gen11674 Obj
			if True == gen11673 {
				gen11670 := Call(__e, ShenFunc(symhd), V4885)

				gen11671 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11670)

				var gen11672 Obj
				if True == gen11671 {
					gen11668 := Call(__e, ShenFunc(symtl), V4885)

					gen11669 := Call(__e, ShenFunc(sym_a), Nil, gen11668)

					if True == gen11669 {
						gen11672 = True
					} else {
						gen11672 = False
					}

				} else {
					gen11672 = False
				}
				if True == gen11672 {
					gen11674 = True
				} else {
					gen11674 = False
				}

			} else {
				gen11674 = False
			}
			if True == gen11674 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("Unfulfilled shen.x.features.cond-expand clause."))

				return
			} else {
				gen11666 := Call(__e, ShenFunc(symcons_2), V4885)

				var gen11667 Obj
				if True == gen11666 {
					gen11663 := Call(__e, ShenFunc(symhd), V4885)

					gen11664 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11663)

					var gen11665 Obj
					if True == gen11664 {
						gen11660 := Call(__e, ShenFunc(symtl), V4885)

						gen11661 := Call(__e, ShenFunc(symcons_2), gen11660)

						var gen11662 Obj
						if True == gen11661 {
							gen11656 := Call(__e, ShenFunc(symtl), V4885)

							gen11657 := Call(__e, ShenFunc(symhd), gen11656)

							gen11658 := Call(__e, ShenFunc(sym_a), True, gen11657)

							var gen11659 Obj
							if True == gen11658 {
								gen11652 := Call(__e, ShenFunc(symtl), V4885)

								gen11653 := Call(__e, ShenFunc(symtl), gen11652)

								gen11654 := Call(__e, ShenFunc(symcons_2), gen11653)

								var gen11655 Obj
								if True == gen11654 {
									gen11648 := Call(__e, ShenFunc(symtl), V4885)

									gen11649 := Call(__e, ShenFunc(symtl), gen11648)

									gen11650 := Call(__e, ShenFunc(symtl), gen11649)

									gen11651 := Call(__e, ShenFunc(sym_a), Nil, gen11650)

									if True == gen11651 {
										gen11655 = True
									} else {
										gen11655 = False
									}

								} else {
									gen11655 = False
								}
								if True == gen11655 {
									gen11659 = True
								} else {
									gen11659 = False
								}

							} else {
								gen11659 = False
							}
							if True == gen11659 {
								gen11662 = True
							} else {
								gen11662 = False
							}

						} else {
							gen11662 = False
						}
						if True == gen11662 {
							gen11665 = True
						} else {
							gen11665 = False
						}

					} else {
						gen11665 = False
					}
					if True == gen11665 {
						gen11667 = True
					} else {
						gen11667 = False
					}

				} else {
					gen11667 = False
				}
				if True == gen11667 {
					gen11646 := Call(__e, ShenFunc(symtl), V4885)

					gen11647 := Call(__e, ShenFunc(symtl), gen11646)

					__e.TailApply(ShenFunc(symhd), gen11647)

					return

				} else {
					gen11644 := Call(__e, ShenFunc(symcons_2), V4885)

					var gen11645 Obj
					if True == gen11644 {
						gen11641 := Call(__e, ShenFunc(symhd), V4885)

						gen11642 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11641)

						var gen11643 Obj
						if True == gen11642 {
							gen11638 := Call(__e, ShenFunc(symtl), V4885)

							gen11639 := Call(__e, ShenFunc(symcons_2), gen11638)

							var gen11640 Obj
							if True == gen11639 {
								gen11634 := Call(__e, ShenFunc(symtl), V4885)

								gen11635 := Call(__e, ShenFunc(symhd), gen11634)

								gen11636 := Call(__e, ShenFunc(symcons_2), gen11635)

								var gen11637 Obj
								if True == gen11636 {
									gen11629 := Call(__e, ShenFunc(symtl), V4885)

									gen11630 := Call(__e, ShenFunc(symhd), gen11629)

									gen11631 := Call(__e, ShenFunc(symhd), gen11630)

									gen11632 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen11631)

									var gen11633 Obj
									if True == gen11632 {
										gen11624 := Call(__e, ShenFunc(symtl), V4885)

										gen11625 := Call(__e, ShenFunc(symhd), gen11624)

										gen11626 := Call(__e, ShenFunc(symtl), gen11625)

										gen11627 := Call(__e, ShenFunc(sym_a), Nil, gen11626)

										var gen11628 Obj
										if True == gen11627 {
											gen11621 := Call(__e, ShenFunc(symtl), V4885)

											gen11622 := Call(__e, ShenFunc(symtl), gen11621)

											gen11623 := Call(__e, ShenFunc(symcons_2), gen11622)

											if True == gen11623 {
												gen11628 = True
											} else {
												gen11628 = False
											}

										} else {
											gen11628 = False
										}
										if True == gen11628 {
											gen11633 = True
										} else {
											gen11633 = False
										}

									} else {
										gen11633 = False
									}
									if True == gen11633 {
										gen11637 = True
									} else {
										gen11637 = False
									}

								} else {
									gen11637 = False
								}
								if True == gen11637 {
									gen11640 = True
								} else {
									gen11640 = False
								}

							} else {
								gen11640 = False
							}
							if True == gen11640 {
								gen11643 = True
							} else {
								gen11643 = False
							}

						} else {
							gen11643 = False
						}
						if True == gen11643 {
							gen11645 = True
						} else {
							gen11645 = False
						}

					} else {
						gen11645 = False
					}
					if True == gen11645 {
						gen11619 := Call(__e, ShenFunc(symtl), V4885)

						gen11620 := Call(__e, ShenFunc(symtl), gen11619)

						__e.TailApply(ShenFunc(symhd), gen11620)

						return

					} else {
						gen11617 := Call(__e, ShenFunc(symcons_2), V4885)

						var gen11618 Obj
						if True == gen11617 {
							gen11614 := Call(__e, ShenFunc(symhd), V4885)

							gen11615 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11614)

							var gen11616 Obj
							if True == gen11615 {
								gen11611 := Call(__e, ShenFunc(symtl), V4885)

								gen11612 := Call(__e, ShenFunc(symcons_2), gen11611)

								var gen11613 Obj
								if True == gen11612 {
									gen11607 := Call(__e, ShenFunc(symtl), V4885)

									gen11608 := Call(__e, ShenFunc(symhd), gen11607)

									gen11609 := Call(__e, ShenFunc(symcons_2), gen11608)

									var gen11610 Obj
									if True == gen11609 {
										gen11602 := Call(__e, ShenFunc(symtl), V4885)

										gen11603 := Call(__e, ShenFunc(symhd), gen11602)

										gen11604 := Call(__e, ShenFunc(symhd), gen11603)

										gen11605 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen11604)

										var gen11606 Obj
										if True == gen11605 {
											gen11597 := Call(__e, ShenFunc(symtl), V4885)

											gen11598 := Call(__e, ShenFunc(symhd), gen11597)

											gen11599 := Call(__e, ShenFunc(symtl), gen11598)

											gen11600 := Call(__e, ShenFunc(symcons_2), gen11599)

											var gen11601 Obj
											if True == gen11600 {
												gen11594 := Call(__e, ShenFunc(symtl), V4885)

												gen11595 := Call(__e, ShenFunc(symtl), gen11594)

												gen11596 := Call(__e, ShenFunc(symcons_2), gen11595)

												if True == gen11596 {
													gen11601 = True
												} else {
													gen11601 = False
												}

											} else {
												gen11601 = False
											}
											if True == gen11601 {
												gen11606 = True
											} else {
												gen11606 = False
											}

										} else {
											gen11606 = False
										}
										if True == gen11606 {
											gen11610 = True
										} else {
											gen11610 = False
										}

									} else {
										gen11610 = False
									}
									if True == gen11610 {
										gen11613 = True
									} else {
										gen11613 = False
									}

								} else {
									gen11613 = False
								}
								if True == gen11613 {
									gen11616 = True
								} else {
									gen11616 = False
								}

							} else {
								gen11616 = False
							}
							if True == gen11616 {
								gen11618 = True
							} else {
								gen11618 = False
							}

						} else {
							gen11618 = False
						}
						if True == gen11618 {
							gen11576 := Call(__e, ShenFunc(symtl), V4885)

							gen11577 := Call(__e, ShenFunc(symhd), gen11576)

							gen11578 := Call(__e, ShenFunc(symtl), gen11577)

							gen11579 := Call(__e, ShenFunc(symhd), gen11578)

							gen11580 := Call(__e, ShenFunc(symtl), V4885)

							gen11581 := Call(__e, ShenFunc(symhd), gen11580)

							gen11582 := Call(__e, ShenFunc(symtl), gen11581)

							gen11583 := Call(__e, ShenFunc(symtl), gen11582)

							gen11584 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen11583)

							gen11585 := Call(__e, ShenFunc(symtl), V4885)

							gen11586 := Call(__e, ShenFunc(symtl), gen11585)

							gen11587 := Call(__e, ShenFunc(symcons), gen11584, gen11586)

							gen11588 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11587)

							gen11589 := Call(__e, ShenFunc(symtl), V4885)

							gen11590 := Call(__e, ShenFunc(symtl), gen11589)

							gen11591 := Call(__e, ShenFunc(symtl), gen11590)

							gen11592 := Call(__e, ShenFunc(symcons), gen11588, gen11591)

							gen11593 := Call(__e, ShenFunc(symcons), gen11579, gen11592)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11593)

							return

						} else {
							gen11574 := Call(__e, ShenFunc(symcons_2), V4885)

							var gen11575 Obj
							if True == gen11574 {
								gen11571 := Call(__e, ShenFunc(symhd), V4885)

								gen11572 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11571)

								var gen11573 Obj
								if True == gen11572 {
									gen11568 := Call(__e, ShenFunc(symtl), V4885)

									gen11569 := Call(__e, ShenFunc(symcons_2), gen11568)

									var gen11570 Obj
									if True == gen11569 {
										gen11564 := Call(__e, ShenFunc(symtl), V4885)

										gen11565 := Call(__e, ShenFunc(symhd), gen11564)

										gen11566 := Call(__e, ShenFunc(symcons_2), gen11565)

										var gen11567 Obj
										if True == gen11566 {
											gen11559 := Call(__e, ShenFunc(symtl), V4885)

											gen11560 := Call(__e, ShenFunc(symhd), gen11559)

											gen11561 := Call(__e, ShenFunc(symhd), gen11560)

											gen11562 := Call(__e, ShenFunc(sym_a), MakeSymbol("or"), gen11561)

											var gen11563 Obj
											if True == gen11562 {
												gen11554 := Call(__e, ShenFunc(symtl), V4885)

												gen11555 := Call(__e, ShenFunc(symhd), gen11554)

												gen11556 := Call(__e, ShenFunc(symtl), gen11555)

												gen11557 := Call(__e, ShenFunc(sym_a), Nil, gen11556)

												var gen11558 Obj
												if True == gen11557 {
													gen11551 := Call(__e, ShenFunc(symtl), V4885)

													gen11552 := Call(__e, ShenFunc(symtl), gen11551)

													gen11553 := Call(__e, ShenFunc(symcons_2), gen11552)

													if True == gen11553 {
														gen11558 = True
													} else {
														gen11558 = False
													}

												} else {
													gen11558 = False
												}
												if True == gen11558 {
													gen11563 = True
												} else {
													gen11563 = False
												}

											} else {
												gen11563 = False
											}
											if True == gen11563 {
												gen11567 = True
											} else {
												gen11567 = False
											}

										} else {
											gen11567 = False
										}
										if True == gen11567 {
											gen11570 = True
										} else {
											gen11570 = False
										}

									} else {
										gen11570 = False
									}
									if True == gen11570 {
										gen11573 = True
									} else {
										gen11573 = False
									}

								} else {
									gen11573 = False
								}
								if True == gen11573 {
									gen11575 = True
								} else {
									gen11575 = False
								}

							} else {
								gen11575 = False
							}
							if True == gen11575 {
								gen11548 := Call(__e, ShenFunc(symtl), V4885)

								gen11549 := Call(__e, ShenFunc(symtl), gen11548)

								gen11550 := Call(__e, ShenFunc(symtl), gen11549)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11550)

								return

							} else {
								gen11546 := Call(__e, ShenFunc(symcons_2), V4885)

								var gen11547 Obj
								if True == gen11546 {
									gen11543 := Call(__e, ShenFunc(symhd), V4885)

									gen11544 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11543)

									var gen11545 Obj
									if True == gen11544 {
										gen11540 := Call(__e, ShenFunc(symtl), V4885)

										gen11541 := Call(__e, ShenFunc(symcons_2), gen11540)

										var gen11542 Obj
										if True == gen11541 {
											gen11536 := Call(__e, ShenFunc(symtl), V4885)

											gen11537 := Call(__e, ShenFunc(symhd), gen11536)

											gen11538 := Call(__e, ShenFunc(symcons_2), gen11537)

											var gen11539 Obj
											if True == gen11538 {
												gen11531 := Call(__e, ShenFunc(symtl), V4885)

												gen11532 := Call(__e, ShenFunc(symhd), gen11531)

												gen11533 := Call(__e, ShenFunc(symhd), gen11532)

												gen11534 := Call(__e, ShenFunc(sym_a), MakeSymbol("or"), gen11533)

												var gen11535 Obj
												if True == gen11534 {
													gen11526 := Call(__e, ShenFunc(symtl), V4885)

													gen11527 := Call(__e, ShenFunc(symhd), gen11526)

													gen11528 := Call(__e, ShenFunc(symtl), gen11527)

													gen11529 := Call(__e, ShenFunc(symcons_2), gen11528)

													var gen11530 Obj
													if True == gen11529 {
														gen11523 := Call(__e, ShenFunc(symtl), V4885)

														gen11524 := Call(__e, ShenFunc(symtl), gen11523)

														gen11525 := Call(__e, ShenFunc(symcons_2), gen11524)

														if True == gen11525 {
															gen11530 = True
														} else {
															gen11530 = False
														}

													} else {
														gen11530 = False
													}
													if True == gen11530 {
														gen11535 = True
													} else {
														gen11535 = False
													}

												} else {
													gen11535 = False
												}
												if True == gen11535 {
													gen11539 = True
												} else {
													gen11539 = False
												}

											} else {
												gen11539 = False
											}
											if True == gen11539 {
												gen11542 = True
											} else {
												gen11542 = False
											}

										} else {
											gen11542 = False
										}
										if True == gen11542 {
											gen11545 = True
										} else {
											gen11545 = False
										}

									} else {
										gen11545 = False
									}
									if True == gen11545 {
										gen11547 = True
									} else {
										gen11547 = False
									}

								} else {
									gen11547 = False
								}
								if True == gen11547 {
									gen11503 := Call(__e, ShenFunc(symtl), V4885)

									gen11504 := Call(__e, ShenFunc(symhd), gen11503)

									gen11505 := Call(__e, ShenFunc(symtl), gen11504)

									gen11506 := Call(__e, ShenFunc(symhd), gen11505)

									gen11507 := Call(__e, ShenFunc(symtl), V4885)

									gen11508 := Call(__e, ShenFunc(symtl), gen11507)

									gen11509 := Call(__e, ShenFunc(symhd), gen11508)

									gen11510 := Call(__e, ShenFunc(symtl), V4885)

									gen11511 := Call(__e, ShenFunc(symhd), gen11510)

									gen11512 := Call(__e, ShenFunc(symtl), gen11511)

									gen11513 := Call(__e, ShenFunc(symtl), gen11512)

									gen11514 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen11513)

									gen11515 := Call(__e, ShenFunc(symtl), V4885)

									gen11516 := Call(__e, ShenFunc(symtl), gen11515)

									gen11517 := Call(__e, ShenFunc(symcons), gen11514, gen11516)

									gen11518 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11517)

									gen11519 := Call(__e, ShenFunc(symcons), gen11518, Nil)

									gen11520 := Call(__e, ShenFunc(symcons), True, gen11519)

									gen11521 := Call(__e, ShenFunc(symcons), gen11509, gen11520)

									gen11522 := Call(__e, ShenFunc(symcons), gen11506, gen11521)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11522)

									return

								} else {
									gen11501 := Call(__e, ShenFunc(symcons_2), V4885)

									var gen11502 Obj
									if True == gen11501 {
										gen11498 := Call(__e, ShenFunc(symhd), V4885)

										gen11499 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11498)

										var gen11500 Obj
										if True == gen11499 {
											gen11495 := Call(__e, ShenFunc(symtl), V4885)

											gen11496 := Call(__e, ShenFunc(symcons_2), gen11495)

											var gen11497 Obj
											if True == gen11496 {
												gen11491 := Call(__e, ShenFunc(symtl), V4885)

												gen11492 := Call(__e, ShenFunc(symhd), gen11491)

												gen11493 := Call(__e, ShenFunc(symcons_2), gen11492)

												var gen11494 Obj
												if True == gen11493 {
													gen11486 := Call(__e, ShenFunc(symtl), V4885)

													gen11487 := Call(__e, ShenFunc(symhd), gen11486)

													gen11488 := Call(__e, ShenFunc(symhd), gen11487)

													gen11489 := Call(__e, ShenFunc(sym_a), MakeSymbol("not"), gen11488)

													var gen11490 Obj
													if True == gen11489 {
														gen11481 := Call(__e, ShenFunc(symtl), V4885)

														gen11482 := Call(__e, ShenFunc(symhd), gen11481)

														gen11483 := Call(__e, ShenFunc(symtl), gen11482)

														gen11484 := Call(__e, ShenFunc(symcons_2), gen11483)

														var gen11485 Obj
														if True == gen11484 {
															gen11475 := Call(__e, ShenFunc(symtl), V4885)

															gen11476 := Call(__e, ShenFunc(symhd), gen11475)

															gen11477 := Call(__e, ShenFunc(symtl), gen11476)

															gen11478 := Call(__e, ShenFunc(symtl), gen11477)

															gen11479 := Call(__e, ShenFunc(sym_a), Nil, gen11478)

															var gen11480 Obj
															if True == gen11479 {
																gen11472 := Call(__e, ShenFunc(symtl), V4885)

																gen11473 := Call(__e, ShenFunc(symtl), gen11472)

																gen11474 := Call(__e, ShenFunc(symcons_2), gen11473)

																if True == gen11474 {
																	gen11480 = True
																} else {
																	gen11480 = False
																}

															} else {
																gen11480 = False
															}
															if True == gen11480 {
																gen11485 = True
															} else {
																gen11485 = False
															}

														} else {
															gen11485 = False
														}
														if True == gen11485 {
															gen11490 = True
														} else {
															gen11490 = False
														}

													} else {
														gen11490 = False
													}
													if True == gen11490 {
														gen11494 = True
													} else {
														gen11494 = False
													}

												} else {
													gen11494 = False
												}
												if True == gen11494 {
													gen11497 = True
												} else {
													gen11497 = False
												}

											} else {
												gen11497 = False
											}
											if True == gen11497 {
												gen11500 = True
											} else {
												gen11500 = False
											}

										} else {
											gen11500 = False
										}
										if True == gen11500 {
											gen11502 = True
										} else {
											gen11502 = False
										}

									} else {
										gen11502 = False
									}
									if True == gen11502 {
										gen11457 := Call(__e, ShenFunc(symtl), V4885)

										gen11458 := Call(__e, ShenFunc(symhd), gen11457)

										gen11459 := Call(__e, ShenFunc(symtl), gen11458)

										gen11460 := Call(__e, ShenFunc(symhd), gen11459)

										gen11461 := Call(__e, ShenFunc(symtl), V4885)

										gen11462 := Call(__e, ShenFunc(symtl), gen11461)

										gen11463 := Call(__e, ShenFunc(symtl), gen11462)

										gen11464 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11463)

										gen11465 := Call(__e, ShenFunc(symtl), V4885)

										gen11466 := Call(__e, ShenFunc(symtl), gen11465)

										gen11467 := Call(__e, ShenFunc(symhd), gen11466)

										gen11468 := Call(__e, ShenFunc(symcons), gen11467, Nil)

										gen11469 := Call(__e, ShenFunc(symcons), True, gen11468)

										gen11470 := Call(__e, ShenFunc(symcons), gen11464, gen11469)

										gen11471 := Call(__e, ShenFunc(symcons), gen11460, gen11470)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11471)

										return

									} else {
										gen11455 := Call(__e, ShenFunc(symcons_2), V4885)

										var gen11456 Obj
										if True == gen11455 {
											gen11452 := Call(__e, ShenFunc(symhd), V4885)

											gen11453 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11452)

											var gen11454 Obj
											if True == gen11453 {
												gen11449 := Call(__e, ShenFunc(symtl), V4885)

												gen11450 := Call(__e, ShenFunc(symcons_2), gen11449)

												var gen11451 Obj
												if True == gen11450 {
													gen11445 := Call(__e, ShenFunc(symtl), V4885)

													gen11446 := Call(__e, ShenFunc(symtl), gen11445)

													gen11447 := Call(__e, ShenFunc(symcons_2), gen11446)

													var gen11448 Obj
													if True == gen11447 {
														gen11441 := Call(__e, ShenFunc(symtl), V4885)

														gen11442 := Call(__e, ShenFunc(symhd), gen11441)

														gen11443 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.x.features.*features*"))

														gen11444 := Call(__e, ShenFunc(symelement_2), gen11442, gen11443)

														if True == gen11444 {
															gen11448 = True
														} else {
															gen11448 = False
														}

													} else {
														gen11448 = False
													}
													if True == gen11448 {
														gen11451 = True
													} else {
														gen11451 = False
													}

												} else {
													gen11451 = False
												}
												if True == gen11451 {
													gen11454 = True
												} else {
													gen11454 = False
												}

											} else {
												gen11454 = False
											}
											if True == gen11454 {
												gen11456 = True
											} else {
												gen11456 = False
											}

										} else {
											gen11456 = False
										}
										if True == gen11456 {
											gen11439 := Call(__e, ShenFunc(symtl), V4885)

											gen11440 := Call(__e, ShenFunc(symtl), gen11439)

											__e.TailApply(ShenFunc(symhd), gen11440)

											return

										} else {
											gen11437 := Call(__e, ShenFunc(symcons_2), V4885)

											var gen11438 Obj
											if True == gen11437 {
												gen11434 := Call(__e, ShenFunc(symhd), V4885)

												gen11435 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.x.features.cond-expand"), gen11434)

												var gen11436 Obj
												if True == gen11435 {
													gen11431 := Call(__e, ShenFunc(symtl), V4885)

													gen11432 := Call(__e, ShenFunc(symcons_2), gen11431)

													var gen11433 Obj
													if True == gen11432 {
														gen11428 := Call(__e, ShenFunc(symtl), V4885)

														gen11429 := Call(__e, ShenFunc(symtl), gen11428)

														gen11430 := Call(__e, ShenFunc(symcons_2), gen11429)

														if True == gen11430 {
															gen11433 = True
														} else {
															gen11433 = False
														}

													} else {
														gen11433 = False
													}
													if True == gen11433 {
														gen11436 = True
													} else {
														gen11436 = False
													}

												} else {
													gen11436 = False
												}
												if True == gen11436 {
													gen11438 = True
												} else {
													gen11438 = False
												}

											} else {
												gen11438 = False
											}
											if True == gen11438 {
												gen11425 := Call(__e, ShenFunc(symtl), V4885)

												gen11426 := Call(__e, ShenFunc(symtl), gen11425)

												gen11427 := Call(__e, ShenFunc(symtl), gen11426)

												__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand"), gen11427)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.features.cond-expand-macro"), gen11675)

		gen11676 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.x.features.*features*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.features.current"), gen11676)

		gen11684 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4887 := __args[0]
			_ = V4887
			gen11679 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				Call(__e, ShenFunc(symset), MakeSymbol("shen.x.features.*features*"), Nil)
				gen11677 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4x_4features_4cond_1expand_1macro), X)

					return
				}, 1)
				gen11678 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.x.features.cond-expand-macro"), gen11677)

				Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen11678)

				__e.TailApply(ShenFunc(symshen_4add_1macro), MakeSymbol("shen.x.features.cond-expand-macro"))

				return

			}, 1)
			gen11680 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.x.features.*features*"))

				return
			}, 0)
			gen11681 := Try(__e, gen11680).Catch(gen11679)
			_ = gen11681
			gen11682 := Call(__e, ShenFunc(symshen_4x_4features_4current))

			Old := gen11682
			_ = Old
			gen11683 := Call(__e, ShenFunc(symset), MakeSymbol("shen.x.features.*features*"), V4887)

			_ = gen11683
			__e.Return(Old)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.x.features.initialise"), gen11684)

		gen11688 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4889 := __args[0]
			_ = V4889
			gen11685 := Call(__e, ShenFunc(symshen_4x_4features_4current))

			Old := gen11685
			gen11686 := Call(__e, ShenFunc(symadjoin), V4889, Old)

			gen11687 := Call(__e, ShenFunc(symset), MakeSymbol("shen.x.features.*features*"), gen11686)

			_ = gen11687
			__e.Return(Old)
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.x.features.add"), gen11688)

		return

	}, 0))
}
