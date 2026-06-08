package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
	tmp8301 := MakeNative(func(__e *ControlFlow) {
		V1553 := __e.Get(1)
		_ = V1553
		__e.TailApply(PrimFunc(symshen_4assert_d), V1553, symshen_4top)
		return
	}, 1)

	tmp8302 := Call(__e, ns2_1set, symasserta, tmp8301)

	_ = tmp8302

	tmp8303 := MakeNative(func(__e *ControlFlow) {
		V1554 := __e.Get(1)
		_ = V1554
		__e.TailApply(PrimFunc(symshen_4assert_d), V1554, symshen_4bottom)
		return
	}, 1)

	tmp8304 := Call(__e, ns2_1set, symassertz, tmp8303)

	_ = tmp8304

	tmp8305 := MakeNative(func(__e *ControlFlow) {
		V1555 := __e.Get(1)
		_ = V1555
		V1556 := __e.Get(2)
		_ = V1556
		tmp8339 := PrimIsPair(V1555)

		var ifres8330 Obj

		if True == tmp8339 {
			tmp8337 := PrimTail(V1555)

			tmp8338 := PrimIsPair(tmp8337)

			var ifres8332 Obj

			if True == tmp8338 {
				tmp8334 := PrimTail(V1555)

				tmp8335 := PrimHead(tmp8334)

				tmp8336 := PrimEqual(sym_5_1_1, tmp8335)

				var ifres8333 Obj

				if True == tmp8336 {
					ifres8333 = True

				} else {
					ifres8333 = False

				}

				ifres8332 = ifres8333

			} else {
				ifres8332 = False

			}

			var ifres8331 Obj

			if True == ifres8332 {
				ifres8331 = True

			} else {
				ifres8331 = False

			}

			ifres8330 = ifres8331

		} else {
			ifres8330 = False

		}

		if True == ifres8330 {
			tmp8306 := MakeNative(func(__e *ControlFlow) {
				W1557 := __e.Get(1)
				_ = W1557
				tmp8307 := MakeNative(func(__e *ControlFlow) {
					W1558 := __e.Get(1)
					_ = W1558
					tmp8308 := MakeNative(func(__e *ControlFlow) {
						W1559 := __e.Get(1)
						_ = W1559
						tmp8309 := MakeNative(func(__e *ControlFlow) {
							W1560 := __e.Get(1)
							_ = W1560
							tmp8310 := MakeNative(func(__e *ControlFlow) {
								W1561 := __e.Get(1)
								_ = W1561
								tmp8311 := MakeNative(func(__e *ControlFlow) {
									W1562 := __e.Get(1)
									_ = W1562
									tmp8312 := MakeNative(func(__e *ControlFlow) {
										W1563 := __e.Get(1)
										_ = W1563
										__e.Return(W1557)
										return
									}, 1)

									tmp8313 := PrimTail(V1555)

									tmp8314 := PrimTail(tmp8313)

									tmp8315 := Call(__e, PrimFunc(symshen_4insert_1info), W1557, W1558, tmp8314, V1555, V1556)

									__e.TailApply(tmp8312, tmp8315)
									return

								}, 1)

								tmp8321 := PrimEqual(W1561, MakeNumber(-1))

								var ifres8316 Obj

								if True == tmp8321 {
									tmp8317 := Call(__e, PrimFunc(symshen_4create_1skeleton), W1557, W1560)

									tmp8318 := Call(__e, PrimFunc(symeval), tmp8317)

									_ = tmp8318

									tmp8319 := PrimValue(sym_dproperty_1vector_d)

									tmp8320 := Call(__e, PrimFunc(symput), W1557, symshen_4dynamic, Nil, tmp8319)

									ifres8316 = tmp8320

								} else {
									ifres8316 = symshen_4skip

								}

								__e.TailApply(tmp8311, ifres8316)
								return

							}, 1)

							tmp8322 := Call(__e, PrimFunc(symarity), W1557)

							__e.TailApply(tmp8310, tmp8322)
							return

						}, 1)

						tmp8323 := Call(__e, PrimFunc(symshen_4parameters), W1559)

						__e.TailApply(tmp8309, tmp8323)
						return

					}, 1)

					tmp8324 := Call(__e, PrimFunc(symlength), W1558)

					__e.TailApply(tmp8308, tmp8324)
					return

				}, 1)

				tmp8325 := PrimHead(V1555)

				tmp8326 := Call(__e, PrimFunc(symshen_4terms), tmp8325)

				__e.TailApply(tmp8307, tmp8326)
				return

			}, 1)

			tmp8327 := PrimHead(V1555)

			tmp8328 := Call(__e, PrimFunc(symshen_4predicate), tmp8327)

			__e.TailApply(tmp8306, tmp8328)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assert_d)
			return
		}

	}, 2)

	tmp8340 := Call(__e, ns2_1set, symshen_4assert_d, tmp8305)

	_ = tmp8340

	tmp8341 := MakeNative(func(__e *ControlFlow) {
		V1566 := __e.Get(1)
		_ = V1566
		tmp8343 := PrimIsPair(V1566)

		if True == tmp8343 {
			__e.Return(PrimHead(V1566))
			return
		} else {
			__e.Return(V1566)
			return
		}

	}, 1)

	tmp8344 := Call(__e, ns2_1set, symshen_4predicate, tmp8341)

	_ = tmp8344

	tmp8345 := MakeNative(func(__e *ControlFlow) {
		V1571 := __e.Get(1)
		_ = V1571
		tmp8347 := PrimIsPair(V1571)

		if True == tmp8347 {
			__e.Return(PrimTail(V1571))
			return
		} else {
			__e.Return(Nil)
			return
		}

	}, 1)

	tmp8348 := Call(__e, ns2_1set, symshen_4terms, tmp8345)

	_ = tmp8348

	tmp8349 := MakeNative(func(__e *ControlFlow) {
		V1572 := __e.Get(1)
		_ = V1572
		V1573 := __e.Get(2)
		_ = V1573
		tmp8350 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1572, V1573)

		tmp8351 := PrimCons(V1572, tmp8350)

		__e.Return(PrimCons(symdefprolog, tmp8351))
		return

	}, 2)

	tmp8352 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp8349)

	_ = tmp8352

	tmp8353 := MakeNative(func(__e *ControlFlow) {
		V1574 := __e.Get(1)
		_ = V1574
		V1575 := __e.Get(2)
		_ = V1575
		tmp8354 := Call(__e, PrimFunc(symshen_4cons_1form), V1575)

		tmp8355 := PrimCons(symshen_4dynamic, Nil)

		tmp8356 := PrimCons(V1574, tmp8355)

		tmp8357 := PrimCons(symget, tmp8356)

		tmp8358 := PrimCons(tmp8357, Nil)

		tmp8359 := PrimCons(tmp8354, tmp8358)

		tmp8360 := PrimCons(symshen_4call_1dynamic, tmp8359)

		tmp8361 := PrimIntern(MakeString(";"))

		tmp8362 := PrimCons(tmp8361, Nil)

		tmp8363 := PrimCons(tmp8360, tmp8362)

		tmp8364 := PrimCons(sym_5_1_1, tmp8363)

		__e.TailApply(PrimFunc(symappend), V1575, tmp8364)
		return

	}, 2)

	tmp8365 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp8353)

	_ = tmp8365

	tmp8366 := MakeNative(func(__e *ControlFlow) {
		V1576 := __e.Get(1)
		_ = V1576
		V1577 := __e.Get(2)
		_ = V1577
		V1578 := __e.Get(3)
		_ = V1578
		V1579 := __e.Get(4)
		_ = V1579
		V1580 := __e.Get(5)
		_ = V1580
		tmp8367 := MakeNative(func(__e *ControlFlow) {
			W1581 := __e.Get(1)
			_ = W1581
			tmp8368 := MakeNative(func(__e *ControlFlow) {
				W1582 := __e.Get(1)
				_ = W1582
				tmp8369 := MakeNative(func(__e *ControlFlow) {
					W1583 := __e.Get(1)
					_ = W1583
					tmp8370 := MakeNative(func(__e *ControlFlow) {
						W1584 := __e.Get(1)
						_ = W1584
						tmp8371 := MakeNative(func(__e *ControlFlow) {
							W1585 := __e.Get(1)
							_ = W1585
							tmp8372 := PrimValue(sym_dproperty_1vector_d)

							__e.TailApply(PrimFunc(symput), V1576, symshen_4dynamic, W1585, tmp8372)
							return

						}, 1)

						tmp8377 := PrimEqual(V1580, symshen_4top)

						var ifres8373 Obj

						if True == tmp8377 {
							tmp8374 := PrimCons(W1583, W1584)

							ifres8373 = tmp8374

						} else {
							tmp8375 := PrimCons(W1583, Nil)

							tmp8376 := Call(__e, PrimFunc(symappend), W1584, tmp8375)

							ifres8373 = tmp8376

						}

						__e.TailApply(tmp8371, ifres8373)
						return

					}, 1)

					tmp8378 := PrimValue(sym_dproperty_1vector_d)

					tmp8379 := Call(__e, PrimFunc(symget), V1576, symshen_4dynamic, tmp8378)

					__e.TailApply(tmp8370, tmp8379)
					return

				}, 1)

				tmp8380 := Call(__e, PrimFunc(symfn), W1581)

				tmp8381 := PrimCons(W1581, V1579)

				tmp8382 := PrimCons(tmp8380, tmp8381)

				__e.TailApply(tmp8369, tmp8382)
				return

			}, 1)

			tmp8383 := PrimCons(W1581, Nil)

			tmp8384 := PrimCons(symdefprolog, tmp8383)

			tmp8385 := PrimCons(sym_5_1_1, V1578)

			tmp8386 := Call(__e, PrimFunc(symappend), V1577, tmp8385)

			tmp8387 := Call(__e, PrimFunc(symappend), tmp8384, tmp8386)

			tmp8388 := Call(__e, PrimFunc(symeval), tmp8387)

			__e.TailApply(tmp8368, tmp8388)
			return

		}, 1)

		tmp8389 := Call(__e, PrimFunc(symgensym), symshen_4g)

		__e.TailApply(tmp8367, tmp8389)
		return

	}, 5)

	tmp8390 := Call(__e, ns2_1set, symshen_4insert_1info, tmp8366)

	_ = tmp8390

	tmp8391 := MakeNative(func(__e *ControlFlow) {
		tmp8392 := MakeNative(func(__e *ControlFlow) {
			W1586 := __e.Get(1)
			_ = W1586
			tmp8393 := MakeNative(func(__e *ControlFlow) {
				W1587 := __e.Get(1)
				_ = W1587
				__e.Return(W1587)
				return
			}, 1)

			tmp8399 := Call(__e, PrimFunc(symempty_2), W1586)

			var ifres8394 Obj

			if True == tmp8399 {
				tmp8395 := Call(__e, PrimFunc(symgensym), symshen_4g)

				ifres8394 = tmp8395

			} else {
				tmp8396 := PrimTail(W1586)

				tmp8397 := PrimSet(symshen_4_dnames_d, tmp8396)

				_ = tmp8397

				tmp8398 := PrimHead(W1586)

				ifres8394 = tmp8398

			}

			__e.TailApply(tmp8393, ifres8394)
			return

		}, 1)

		tmp8400 := PrimValue(symshen_4_dnames_d)

		__e.TailApply(tmp8392, tmp8400)
		return

	}, 0)

	tmp8401 := Call(__e, ns2_1set, symshen_4newname, tmp8391)

	_ = tmp8401

	tmp8402 := MakeNative(func(__e *ControlFlow) {
		V1588 := __e.Get(1)
		_ = V1588
		V1589 := __e.Get(2)
		_ = V1589
		V1590 := __e.Get(3)
		_ = V1590
		V1591 := __e.Get(4)
		_ = V1591
		V1592 := __e.Get(5)
		_ = V1592
		V1593 := __e.Get(6)
		_ = V1593
		tmp8403 := MakeNative(func(__e *ControlFlow) {
			W1594 := __e.Get(1)
			_ = W1594
			tmp8414 := PrimEqual(W1594, False)

			if True == tmp8414 {
				tmp8412 := Call(__e, PrimFunc(symshen_4unlocked_2), V1591)

				if True == tmp8412 {
					tmp8404 := MakeNative(func(__e *ControlFlow) {
						W1598 := __e.Get(1)
						_ = W1598
						tmp8409 := PrimIsPair(W1598)

						if True == tmp8409 {
							tmp8405 := MakeNative(func(__e *ControlFlow) {
								W1599 := __e.Get(1)
								_ = W1599
								tmp8406 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp8406

								__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1588, W1599, V1590, V1591, V1592, V1593)
								return

							}, 1)

							tmp8407 := PrimTail(W1598)

							__e.TailApply(tmp8405, tmp8407)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp8410 := Call(__e, PrimFunc(symshen_4lazyderef), V1589, V1590)

					__e.TailApply(tmp8404, tmp8410)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W1594)
				return
			}

		}, 1)

		tmp8429 := Call(__e, PrimFunc(symshen_4unlocked_2), V1591)

		var ifres8415 Obj

		if True == tmp8429 {
			tmp8416 := MakeNative(func(__e *ControlFlow) {
				W1595 := __e.Get(1)
				_ = W1595
				tmp8426 := PrimIsPair(W1595)

				if True == tmp8426 {
					tmp8417 := MakeNative(func(__e *ControlFlow) {
						W1596 := __e.Get(1)
						_ = W1596
						tmp8422 := PrimIsPair(W1596)

						if True == tmp8422 {
							tmp8418 := MakeNative(func(__e *ControlFlow) {
								W1597 := __e.Get(1)
								_ = W1597
								tmp8419 := Call(__e, PrimFunc(symshen_4incinfs))

								_ = tmp8419

								__e.TailApply(PrimFunc(symshen_4callrec), W1597, V1588, V1590, V1591, V1592, V1593)
								return

							}, 1)

							tmp8420 := PrimHead(W1596)

							__e.TailApply(tmp8418, tmp8420)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp8423 := PrimHead(W1595)

					tmp8424 := Call(__e, PrimFunc(symshen_4lazyderef), tmp8423, V1590)

					__e.TailApply(tmp8417, tmp8424)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp8427 := Call(__e, PrimFunc(symshen_4lazyderef), V1589, V1590)

			tmp8428 := Call(__e, tmp8416, tmp8427)

			ifres8415 = tmp8428

		} else {
			ifres8415 = False

		}

		__e.TailApply(tmp8403, ifres8415)
		return

	}, 6)

	tmp8430 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp8402)

	_ = tmp8430

	tmp8431 := MakeNative(func(__e *ControlFlow) {
		V1600 := __e.Get(1)
		_ = V1600
		V1601 := __e.Get(2)
		_ = V1601
		V1602 := __e.Get(3)
		_ = V1602
		V1603 := __e.Get(4)
		_ = V1603
		V1604 := __e.Get(5)
		_ = V1604
		V1605 := __e.Get(6)
		_ = V1605
		tmp8441 := PrimEqual(Nil, V1601)

		if True == tmp8441 {
			tmp8432 := Call(__e, V1600, V1602)

			tmp8433 := Call(__e, tmp8432, V1603)

			tmp8434 := Call(__e, tmp8433, V1604)

			__e.TailApply(tmp8434, V1605)
			return

		} else {
			tmp8439 := PrimIsPair(V1601)

			if True == tmp8439 {
				tmp8435 := PrimHead(V1601)

				tmp8436 := Call(__e, V1600, tmp8435)

				tmp8437 := PrimTail(V1601)

				__e.TailApply(PrimFunc(symshen_4callrec), tmp8436, tmp8437, V1602, V1603, V1604, V1605)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4callrec)
				return
			}

		}

	}, 6)

	tmp8442 := Call(__e, ns2_1set, symshen_4callrec, tmp8431)

	_ = tmp8442

	tmp8443 := MakeNative(func(__e *ControlFlow) {
		V1606 := __e.Get(1)
		_ = V1606
		tmp8462 := PrimIsPair(V1606)

		var ifres8453 Obj

		if True == tmp8462 {
			tmp8460 := PrimTail(V1606)

			tmp8461 := PrimIsPair(tmp8460)

			var ifres8455 Obj

			if True == tmp8461 {
				tmp8457 := PrimTail(V1606)

				tmp8458 := PrimHead(tmp8457)

				tmp8459 := PrimEqual(sym_5_1_1, tmp8458)

				var ifres8456 Obj

				if True == tmp8459 {
					ifres8456 = True

				} else {
					ifres8456 = False

				}

				ifres8455 = ifres8456

			} else {
				ifres8455 = False

			}

			var ifres8454 Obj

			if True == ifres8455 {
				ifres8454 = True

			} else {
				ifres8454 = False

			}

			ifres8453 = ifres8454

		} else {
			ifres8453 = False

		}

		if True == ifres8453 {
			tmp8444 := MakeNative(func(__e *ControlFlow) {
				W1607 := __e.Get(1)
				_ = W1607
				tmp8445 := MakeNative(func(__e *ControlFlow) {
					W1608 := __e.Get(1)
					_ = W1608
					tmp8446 := Call(__e, PrimFunc(symshen_4retract_1clause), V1606, W1608)

					tmp8447 := PrimValue(sym_dproperty_1vector_d)

					__e.TailApply(PrimFunc(symput), W1607, symshen_4dynamic, tmp8446, tmp8447)
					return

				}, 1)

				tmp8448 := PrimValue(sym_dproperty_1vector_d)

				tmp8449 := Call(__e, PrimFunc(symget), W1607, symshen_4dynamic, tmp8448)

				__e.TailApply(tmp8445, tmp8449)
				return

			}, 1)

			tmp8450 := PrimHead(V1606)

			tmp8451 := Call(__e, PrimFunc(symshen_4predicate), tmp8450)

			__e.TailApply(tmp8444, tmp8451)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symretract)
			return
		}

	}, 1)

	tmp8463 := Call(__e, ns2_1set, symretract, tmp8443)

	_ = tmp8463

	tmp8464 := MakeNative(func(__e *ControlFlow) {
		V1614 := __e.Get(1)
		_ = V1614
		V1615 := __e.Get(2)
		_ = V1615
		tmp8494 := PrimEqual(Nil, V1615)

		if True == tmp8494 {
			__e.Return(Nil)
			return
		} else {
			tmp8492 := PrimIsPair(V1615)

			var ifres8477 Obj

			if True == tmp8492 {
				tmp8490 := PrimHead(V1615)

				tmp8491 := PrimIsPair(tmp8490)

				var ifres8479 Obj

				if True == tmp8491 {
					tmp8487 := PrimHead(V1615)

					tmp8488 := PrimTail(tmp8487)

					tmp8489 := PrimIsPair(tmp8488)

					var ifres8481 Obj

					if True == tmp8489 {
						tmp8483 := PrimHead(V1615)

						tmp8484 := PrimTail(tmp8483)

						tmp8485 := PrimTail(tmp8484)

						tmp8486 := PrimEqual(V1614, tmp8485)

						var ifres8482 Obj

						if True == tmp8486 {
							ifres8482 = True

						} else {
							ifres8482 = False

						}

						ifres8481 = ifres8482

					} else {
						ifres8481 = False

					}

					var ifres8480 Obj

					if True == ifres8481 {
						ifres8480 = True

					} else {
						ifres8480 = False

					}

					ifres8479 = ifres8480

				} else {
					ifres8479 = False

				}

				var ifres8478 Obj

				if True == ifres8479 {
					ifres8478 = True

				} else {
					ifres8478 = False

				}

				ifres8477 = ifres8478

			} else {
				ifres8477 = False

			}

			if True == ifres8477 {
				tmp8465 := PrimHead(V1615)

				tmp8466 := PrimTail(tmp8465)

				tmp8467 := PrimHead(tmp8466)

				tmp8468 := PrimValue(symshen_4_dnames_d)

				tmp8469 := PrimCons(tmp8467, tmp8468)

				tmp8470 := PrimSet(symshen_4_dnames_d, tmp8469)

				_ = tmp8470

				__e.Return(PrimTail(V1615))
				return

			} else {
				tmp8475 := PrimIsPair(V1615)

				if True == tmp8475 {
					tmp8471 := PrimHead(V1615)

					tmp8472 := PrimTail(V1615)

					tmp8473 := Call(__e, PrimFunc(symshen_4retract_1clause), V1614, tmp8472)

					__e.Return(PrimCons(tmp8471, tmp8473))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4retract_1clause)
					return
				}

			}

		}

	}, 2)

	tmp8495 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp8464)

	_ = tmp8495

	tmp8496 := MakeNative(func(__e *ControlFlow) {
		V1616 := __e.Get(1)
		_ = V1616
		V1617 := __e.Get(2)
		_ = V1617
		tmp8497 := MakeNative(func(__e *ControlFlow) {
			Z1618 := __e.Get(1)
			_ = Z1618
			__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1618)
			return
		}, 1)

		tmp8498 := PrimCons(V1616, V1617)

		__e.TailApply(PrimFunc(symcompile), tmp8497, tmp8498)
		return

	}, 2)

	tmp8499 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp8496)

	_ = tmp8499

	tmp8500 := MakeNative(func(__e *ControlFlow) {
		V1619 := __e.Get(1)
		_ = V1619
		tmp8501 := MakeNative(func(__e *ControlFlow) {
			W1620 := __e.Get(1)
			_ = W1620
			tmp8503 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1620)

			if True == tmp8503 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W1620)
				return
			}

		}, 1)

		tmp8525 := PrimIsPair(V1619)

		var ifres8504 Obj

		if True == tmp8525 {
			tmp8505 := MakeNative(func(__e *ControlFlow) {
				W1621 := __e.Get(1)
				_ = W1621
				tmp8506 := MakeNative(func(__e *ControlFlow) {
					W1622 := __e.Get(1)
					_ = W1622
					tmp8507 := MakeNative(func(__e *ControlFlow) {
						W1623 := __e.Get(1)
						_ = W1623
						tmp8519 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1623)

						if True == tmp8519 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp8508 := MakeNative(func(__e *ControlFlow) {
								W1624 := __e.Get(1)
								_ = W1624
								tmp8509 := MakeNative(func(__e *ControlFlow) {
									W1625 := __e.Get(1)
									_ = W1625
									tmp8510 := MakeNative(func(__e *ControlFlow) {
										W1626 := __e.Get(1)
										_ = W1626
										tmp8511 := MakeNative(func(__e *ControlFlow) {
											W1627 := __e.Get(1)
											_ = W1627
											__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), W1621, W1627)
											return
										}, 1)

										tmp8512 := MakeNative(func(__e *ControlFlow) {
											Z1628 := __e.Get(1)
											_ = Z1628
											__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1628)
											return
										}, 1)

										tmp8513 := Call(__e, PrimFunc(symmap), tmp8512, W1624)

										__e.TailApply(tmp8511, tmp8513)
										return

									}, 1)

									tmp8514 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), W1621, W1624)

									tmp8515 := Call(__e, tmp8510, tmp8514)

									__e.TailApply(PrimFunc(symshen_4comb), W1625, tmp8515)
									return

								}, 1)

								tmp8516 := Call(__e, PrimFunc(symshen_4in_1_6), W1623)

								__e.TailApply(tmp8509, tmp8516)
								return

							}, 1)

							tmp8517 := Call(__e, PrimFunc(symshen_4_5_1out), W1623)

							__e.TailApply(tmp8508, tmp8517)
							return

						}

					}, 1)

					tmp8520 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1622)

					__e.TailApply(tmp8507, tmp8520)
					return

				}, 1)

				tmp8521 := Call(__e, PrimFunc(symtail), V1619)

				__e.TailApply(tmp8506, tmp8521)
				return

			}, 1)

			tmp8522 := Call(__e, PrimFunc(symhead), V1619)

			tmp8523 := Call(__e, tmp8505, tmp8522)

			ifres8504 = tmp8523

		} else {
			tmp8524 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres8504 = tmp8524

		}

		__e.TailApply(tmp8501, ifres8504)
		return

	}, 1)

	tmp8526 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp8500)

	_ = tmp8526

	tmp8527 := MakeNative(func(__e *ControlFlow) {
		V1631 := __e.Get(1)
		_ = V1631
		V1632 := __e.Get(2)
		_ = V1632
		tmp8571 := PrimIsPair(V1632)

		var ifres8552 Obj

		if True == tmp8571 {
			tmp8569 := PrimHead(V1632)

			tmp8570 := PrimIsPair(tmp8569)

			var ifres8554 Obj

			if True == tmp8570 {
				tmp8566 := PrimHead(V1632)

				tmp8567 := PrimTail(tmp8566)

				tmp8568 := PrimIsPair(tmp8567)

				var ifres8556 Obj

				if True == tmp8568 {
					tmp8562 := PrimHead(V1632)

					tmp8563 := PrimTail(tmp8562)

					tmp8564 := PrimTail(tmp8563)

					tmp8565 := PrimEqual(Nil, tmp8564)

					var ifres8558 Obj

					if True == tmp8565 {
						tmp8560 := PrimTail(V1632)

						tmp8561 := PrimEqual(Nil, tmp8560)

						var ifres8559 Obj

						if True == tmp8561 {
							ifres8559 = True

						} else {
							ifres8559 = False

						}

						ifres8558 = ifres8559

					} else {
						ifres8558 = False

					}

					var ifres8557 Obj

					if True == ifres8558 {
						ifres8557 = True

					} else {
						ifres8557 = False

					}

					ifres8556 = ifres8557

				} else {
					ifres8556 = False

				}

				var ifres8555 Obj

				if True == ifres8556 {
					ifres8555 = True

				} else {
					ifres8555 = False

				}

				ifres8554 = ifres8555

			} else {
				ifres8554 = False

			}

			var ifres8553 Obj

			if True == ifres8554 {
				ifres8553 = True

			} else {
				ifres8553 = False

			}

			ifres8552 = ifres8553

		} else {
			ifres8552 = False

		}

		if True == ifres8552 {
			tmp8528 := PrimHead(V1632)

			tmp8529 := PrimHead(tmp8528)

			__e.TailApply(PrimFunc(symlength), tmp8529)
			return

		} else {
			tmp8550 := PrimIsPair(V1632)

			var ifres8535 Obj

			if True == tmp8550 {
				tmp8548 := PrimHead(V1632)

				tmp8549 := PrimIsPair(tmp8548)

				var ifres8537 Obj

				if True == tmp8549 {
					tmp8545 := PrimHead(V1632)

					tmp8546 := PrimTail(tmp8545)

					tmp8547 := PrimIsPair(tmp8546)

					var ifres8539 Obj

					if True == tmp8547 {
						tmp8541 := PrimHead(V1632)

						tmp8542 := PrimTail(tmp8541)

						tmp8543 := PrimTail(tmp8542)

						tmp8544 := PrimEqual(Nil, tmp8543)

						var ifres8540 Obj

						if True == tmp8544 {
							ifres8540 = True

						} else {
							ifres8540 = False

						}

						ifres8539 = ifres8540

					} else {
						ifres8539 = False

					}

					var ifres8538 Obj

					if True == ifres8539 {
						ifres8538 = True

					} else {
						ifres8538 = False

					}

					ifres8537 = ifres8538

				} else {
					ifres8537 = False

				}

				var ifres8536 Obj

				if True == ifres8537 {
					ifres8536 = True

				} else {
					ifres8536 = False

				}

				ifres8535 = ifres8536

			} else {
				ifres8535 = False

			}

			if True == ifres8535 {
				tmp8530 := PrimHead(V1632)

				tmp8531 := PrimHead(tmp8530)

				tmp8532 := Call(__e, PrimFunc(symlength), tmp8531)

				tmp8533 := PrimTail(V1632)

				__e.TailApply(PrimFunc(symshen_4pac_1h), V1631, tmp8532, tmp8533)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1arity_1check)
				return
			}

		}

	}, 2)

	tmp8572 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp8527)

	_ = tmp8572

	tmp8573 := MakeNative(func(__e *ControlFlow) {
		V1637 := __e.Get(1)
		_ = V1637
		V1638 := __e.Get(2)
		_ = V1638
		V1639 := __e.Get(3)
		_ = V1639
		tmp8589 := PrimEqual(Nil, V1639)

		if True == tmp8589 {
			__e.Return(V1638)
			return
		} else {
			tmp8587 := PrimIsPair(V1639)

			var ifres8583 Obj

			if True == tmp8587 {
				tmp8585 := PrimHead(V1639)

				tmp8586 := PrimIsPair(tmp8585)

				var ifres8584 Obj

				if True == tmp8586 {
					ifres8584 = True

				} else {
					ifres8584 = False

				}

				ifres8583 = ifres8584

			} else {
				ifres8583 = False

			}

			if True == ifres8583 {
				tmp8578 := PrimHead(V1639)

				tmp8579 := PrimHead(tmp8578)

				tmp8580 := Call(__e, PrimFunc(symlength), tmp8579)

				tmp8581 := PrimEqual(V1638, tmp8580)

				if True == tmp8581 {
					tmp8574 := PrimTail(V1639)

					__e.TailApply(PrimFunc(symshen_4pac_1h), V1637, V1638, tmp8574)
					return

				} else {
					tmp8575 := Call(__e, PrimFunc(symshen_4app), V1637, MakeString("\n"), symshen_4a)

					tmp8576 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp8575)

					__e.Return(PrimSimpleError(tmp8576))
					return

				}

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4pac_1h)
				return
			}

		}

	}, 3)

	tmp8590 := Call(__e, ns2_1set, symshen_4pac_1h, tmp8573)

	_ = tmp8590

	tmp8591 := MakeNative(func(__e *ControlFlow) {
		V1640 := __e.Get(1)
		_ = V1640
		tmp8592 := MakeNative(func(__e *ControlFlow) {
			W1641 := __e.Get(1)
			_ = W1641
			tmp8611 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1641)

			if True == tmp8611 {
				tmp8593 := MakeNative(func(__e *ControlFlow) {
					W1648 := __e.Get(1)
					_ = W1648
					tmp8595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1648)

					if True == tmp8595 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W1648)
						return
					}

				}, 1)

				tmp8596 := MakeNative(func(__e *ControlFlow) {
					W1649 := __e.Get(1)
					_ = W1649
					tmp8607 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1649)

					if True == tmp8607 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp8597 := MakeNative(func(__e *ControlFlow) {
							W1650 := __e.Get(1)
							_ = W1650
							tmp8598 := MakeNative(func(__e *ControlFlow) {
								W1651 := __e.Get(1)
								_ = W1651
								tmp8603 := Call(__e, PrimFunc(symempty_2), W1650)

								var ifres8599 Obj

								if True == tmp8603 {
									ifres8599 = Nil

								} else {
									tmp8600 := Call(__e, PrimFunc(symshen_4app), W1650, MakeString("\n ..."), symshen_4r)

									tmp8601 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp8600)

									tmp8602 := PrimSimpleError(tmp8601)

									ifres8599 = tmp8602

								}

								__e.TailApply(PrimFunc(symshen_4comb), W1651, ifres8599)
								return

							}, 1)

							tmp8604 := Call(__e, PrimFunc(symshen_4in_1_6), W1649)

							__e.TailApply(tmp8598, tmp8604)
							return

						}, 1)

						tmp8605 := Call(__e, PrimFunc(symshen_4_5_1out), W1649)

						__e.TailApply(tmp8597, tmp8605)
						return

					}

				}, 1)

				tmp8608 := Call(__e, PrimFunc(sym_5_b_6), V1640)

				tmp8609 := Call(__e, tmp8596, tmp8608)

				__e.TailApply(tmp8593, tmp8609)
				return

			} else {
				__e.Return(W1641)
				return
			}

		}, 1)

		tmp8612 := MakeNative(func(__e *ControlFlow) {
			W1642 := __e.Get(1)
			_ = W1642
			tmp8627 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1642)

			if True == tmp8627 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp8613 := MakeNative(func(__e *ControlFlow) {
					W1643 := __e.Get(1)
					_ = W1643
					tmp8614 := MakeNative(func(__e *ControlFlow) {
						W1644 := __e.Get(1)
						_ = W1644
						tmp8615 := MakeNative(func(__e *ControlFlow) {
							W1645 := __e.Get(1)
							_ = W1645
							tmp8622 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1645)

							if True == tmp8622 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp8616 := MakeNative(func(__e *ControlFlow) {
									W1646 := __e.Get(1)
									_ = W1646
									tmp8617 := MakeNative(func(__e *ControlFlow) {
										W1647 := __e.Get(1)
										_ = W1647
										tmp8618 := PrimCons(W1643, W1646)

										__e.TailApply(PrimFunc(symshen_4comb), W1647, tmp8618)
										return

									}, 1)

									tmp8619 := Call(__e, PrimFunc(symshen_4in_1_6), W1645)

									__e.TailApply(tmp8617, tmp8619)
									return

								}, 1)

								tmp8620 := Call(__e, PrimFunc(symshen_4_5_1out), W1645)

								__e.TailApply(tmp8616, tmp8620)
								return

							}

						}, 1)

						tmp8623 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1644)

						__e.TailApply(tmp8615, tmp8623)
						return

					}, 1)

					tmp8624 := Call(__e, PrimFunc(symshen_4in_1_6), W1642)

					__e.TailApply(tmp8614, tmp8624)
					return

				}, 1)

				tmp8625 := Call(__e, PrimFunc(symshen_4_5_1out), W1642)

				__e.TailApply(tmp8613, tmp8625)
				return

			}

		}, 1)

		tmp8628 := Call(__e, PrimFunc(symshen_4_5clause_6), V1640)

		tmp8629 := Call(__e, tmp8612, tmp8628)

		__e.TailApply(tmp8592, tmp8629)
		return

	}, 1)

	tmp8630 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp8591)

	_ = tmp8630

	tmp8631 := MakeNative(func(__e *ControlFlow) {
		V1652 := __e.Get(1)
		_ = V1652
		tmp8647 := PrimIsPair(V1652)

		var ifres8638 Obj

		if True == tmp8647 {
			tmp8645 := PrimTail(V1652)

			tmp8646 := PrimIsPair(tmp8645)

			var ifres8640 Obj

			if True == tmp8646 {
				tmp8642 := PrimTail(V1652)

				tmp8643 := PrimTail(tmp8642)

				tmp8644 := PrimEqual(Nil, tmp8643)

				var ifres8641 Obj

				if True == tmp8644 {
					ifres8641 = True

				} else {
					ifres8641 = False

				}

				ifres8640 = ifres8641

			} else {
				ifres8640 = False

			}

			var ifres8639 Obj

			if True == ifres8640 {
				ifres8639 = True

			} else {
				ifres8639 = False

			}

			ifres8638 = ifres8639

		} else {
			ifres8638 = False

		}

		if True == ifres8638 {
			tmp8632 := PrimHead(V1652)

			tmp8633 := PrimTail(V1652)

			tmp8634 := PrimHead(tmp8633)

			tmp8635 := Call(__e, PrimFunc(sym_8p), tmp8632, tmp8634)

			tmp8636 := Call(__e, PrimFunc(symshen_4linearise), tmp8635)

			__e.TailApply(PrimFunc(symshen_4lch), tmp8636)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4linearise_1clause)
			return
		}

	}, 1)

	tmp8648 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp8631)

	_ = tmp8648

	tmp8649 := MakeNative(func(__e *ControlFlow) {
		V1653 := __e.Get(1)
		_ = V1653
		tmp8655 := Call(__e, PrimFunc(symtuple_2), V1653)

		if True == tmp8655 {
			tmp8650 := Call(__e, PrimFunc(symfst), V1653)

			tmp8651 := Call(__e, PrimFunc(symsnd), V1653)

			tmp8652 := Call(__e, PrimFunc(symshen_4lchh), tmp8651)

			tmp8653 := PrimCons(tmp8652, Nil)

			__e.Return(PrimCons(tmp8650, tmp8653))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lch)
			return
		}

	}, 1)

	tmp8656 := Call(__e, ns2_1set, symshen_4lch, tmp8649)

	_ = tmp8656

	tmp8657 := MakeNative(func(__e *ControlFlow) {
		V1654 := __e.Get(1)
		_ = V1654
		tmp8720 := PrimIsPair(V1654)

		var ifres8669 Obj

		if True == tmp8720 {
			tmp8718 := PrimHead(V1654)

			tmp8719 := PrimEqual(symwhere, tmp8718)

			var ifres8671 Obj

			if True == tmp8719 {
				tmp8716 := PrimTail(V1654)

				tmp8717 := PrimIsPair(tmp8716)

				var ifres8673 Obj

				if True == tmp8717 {
					tmp8713 := PrimTail(V1654)

					tmp8714 := PrimHead(tmp8713)

					tmp8715 := PrimIsPair(tmp8714)

					var ifres8675 Obj

					if True == tmp8715 {
						tmp8709 := PrimTail(V1654)

						tmp8710 := PrimHead(tmp8709)

						tmp8711 := PrimHead(tmp8710)

						tmp8712 := PrimEqual(sym_a, tmp8711)

						var ifres8677 Obj

						if True == tmp8712 {
							tmp8705 := PrimTail(V1654)

							tmp8706 := PrimHead(tmp8705)

							tmp8707 := PrimTail(tmp8706)

							tmp8708 := PrimIsPair(tmp8707)

							var ifres8679 Obj

							if True == tmp8708 {
								tmp8700 := PrimTail(V1654)

								tmp8701 := PrimHead(tmp8700)

								tmp8702 := PrimTail(tmp8701)

								tmp8703 := PrimTail(tmp8702)

								tmp8704 := PrimIsPair(tmp8703)

								var ifres8681 Obj

								if True == tmp8704 {
									tmp8694 := PrimTail(V1654)

									tmp8695 := PrimHead(tmp8694)

									tmp8696 := PrimTail(tmp8695)

									tmp8697 := PrimTail(tmp8696)

									tmp8698 := PrimTail(tmp8697)

									tmp8699 := PrimEqual(Nil, tmp8698)

									var ifres8683 Obj

									if True == tmp8699 {
										tmp8691 := PrimTail(V1654)

										tmp8692 := PrimTail(tmp8691)

										tmp8693 := PrimIsPair(tmp8692)

										var ifres8685 Obj

										if True == tmp8693 {
											tmp8687 := PrimTail(V1654)

											tmp8688 := PrimTail(tmp8687)

											tmp8689 := PrimTail(tmp8688)

											tmp8690 := PrimEqual(Nil, tmp8689)

											var ifres8686 Obj

											if True == tmp8690 {
												ifres8686 = True

											} else {
												ifres8686 = False

											}

											ifres8685 = ifres8686

										} else {
											ifres8685 = False

										}

										var ifres8684 Obj

										if True == ifres8685 {
											ifres8684 = True

										} else {
											ifres8684 = False

										}

										ifres8683 = ifres8684

									} else {
										ifres8683 = False

									}

									var ifres8682 Obj

									if True == ifres8683 {
										ifres8682 = True

									} else {
										ifres8682 = False

									}

									ifres8681 = ifres8682

								} else {
									ifres8681 = False

								}

								var ifres8680 Obj

								if True == ifres8681 {
									ifres8680 = True

								} else {
									ifres8680 = False

								}

								ifres8679 = ifres8680

							} else {
								ifres8679 = False

							}

							var ifres8678 Obj

							if True == ifres8679 {
								ifres8678 = True

							} else {
								ifres8678 = False

							}

							ifres8677 = ifres8678

						} else {
							ifres8677 = False

						}

						var ifres8676 Obj

						if True == ifres8677 {
							ifres8676 = True

						} else {
							ifres8676 = False

						}

						ifres8675 = ifres8676

					} else {
						ifres8675 = False

					}

					var ifres8674 Obj

					if True == ifres8675 {
						ifres8674 = True

					} else {
						ifres8674 = False

					}

					ifres8673 = ifres8674

				} else {
					ifres8673 = False

				}

				var ifres8672 Obj

				if True == ifres8673 {
					ifres8672 = True

				} else {
					ifres8672 = False

				}

				ifres8671 = ifres8672

			} else {
				ifres8671 = False

			}

			var ifres8670 Obj

			if True == ifres8671 {
				ifres8670 = True

			} else {
				ifres8670 = False

			}

			ifres8669 = ifres8670

		} else {
			ifres8669 = False

		}

		if True == ifres8669 {
			tmp8659 := PrimValue(symshen_4_doccurs_d)

			var ifres8658 Obj

			if True == tmp8659 {
				ifres8658 = symis_b

			} else {
				ifres8658 = symis

			}

			tmp8660 := PrimTail(V1654)

			tmp8661 := PrimHead(tmp8660)

			tmp8662 := PrimTail(tmp8661)

			tmp8663 := PrimCons(ifres8658, tmp8662)

			tmp8664 := PrimTail(V1654)

			tmp8665 := PrimTail(tmp8664)

			tmp8666 := PrimHead(tmp8665)

			tmp8667 := Call(__e, PrimFunc(symshen_4lchh), tmp8666)

			__e.Return(PrimCons(tmp8663, tmp8667))
			return

		} else {
			__e.Return(V1654)
			return
		}

	}, 1)

	tmp8721 := Call(__e, ns2_1set, symshen_4lchh, tmp8657)

	_ = tmp8721

	tmp8722 := MakeNative(func(__e *ControlFlow) {
		V1655 := __e.Get(1)
		_ = V1655
		tmp8723 := MakeNative(func(__e *ControlFlow) {
			W1656 := __e.Get(1)
			_ = W1656
			tmp8725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1656)

			if True == tmp8725 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W1656)
				return
			}

		}, 1)

		tmp8726 := MakeNative(func(__e *ControlFlow) {
			W1657 := __e.Get(1)
			_ = W1657
			tmp8752 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1657)

			if True == tmp8752 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp8727 := MakeNative(func(__e *ControlFlow) {
					W1658 := __e.Get(1)
					_ = W1658
					tmp8728 := MakeNative(func(__e *ControlFlow) {
						W1659 := __e.Get(1)
						_ = W1659
						tmp8748 := Call(__e, PrimFunc(symshen_4hds_a_2), W1659, sym_5_1_1)

						if True == tmp8748 {
							tmp8729 := MakeNative(func(__e *ControlFlow) {
								W1660 := __e.Get(1)
								_ = W1660
								tmp8730 := MakeNative(func(__e *ControlFlow) {
									W1661 := __e.Get(1)
									_ = W1661
									tmp8744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1661)

									if True == tmp8744 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp8731 := MakeNative(func(__e *ControlFlow) {
											W1662 := __e.Get(1)
											_ = W1662
											tmp8732 := MakeNative(func(__e *ControlFlow) {
												W1663 := __e.Get(1)
												_ = W1663
												tmp8733 := MakeNative(func(__e *ControlFlow) {
													W1664 := __e.Get(1)
													_ = W1664
													tmp8739 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1664)

													if True == tmp8739 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp8734 := MakeNative(func(__e *ControlFlow) {
															W1665 := __e.Get(1)
															_ = W1665
															tmp8735 := PrimCons(W1662, Nil)

															tmp8736 := PrimCons(W1658, tmp8735)

															__e.TailApply(PrimFunc(symshen_4comb), W1665, tmp8736)
															return

														}, 1)

														tmp8737 := Call(__e, PrimFunc(symshen_4in_1_6), W1664)

														__e.TailApply(tmp8734, tmp8737)
														return

													}

												}, 1)

												tmp8740 := Call(__e, PrimFunc(symshen_4_5sc_6), W1663)

												__e.TailApply(tmp8733, tmp8740)
												return

											}, 1)

											tmp8741 := Call(__e, PrimFunc(symshen_4in_1_6), W1661)

											__e.TailApply(tmp8732, tmp8741)
											return

										}, 1)

										tmp8742 := Call(__e, PrimFunc(symshen_4_5_1out), W1661)

										__e.TailApply(tmp8731, tmp8742)
										return

									}

								}, 1)

								tmp8745 := Call(__e, PrimFunc(symshen_4_5body_6), W1660)

								__e.TailApply(tmp8730, tmp8745)
								return

							}, 1)

							tmp8746 := Call(__e, PrimFunc(symtail), W1659)

							__e.TailApply(tmp8729, tmp8746)
							return

						} else {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp8749 := Call(__e, PrimFunc(symshen_4in_1_6), W1657)

					__e.TailApply(tmp8728, tmp8749)
					return

				}, 1)

				tmp8750 := Call(__e, PrimFunc(symshen_4_5_1out), W1657)

				__e.TailApply(tmp8727, tmp8750)
				return

			}

		}, 1)

		tmp8753 := Call(__e, PrimFunc(symshen_4_5head_6), V1655)

		tmp8754 := Call(__e, tmp8726, tmp8753)

		__e.TailApply(tmp8723, tmp8754)
		return

	}, 1)

	tmp8755 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp8722)

	_ = tmp8755

	tmp8756 := MakeNative(func(__e *ControlFlow) {
		V1666 := __e.Get(1)
		_ = V1666
		tmp8757 := MakeNative(func(__e *ControlFlow) {
			W1667 := __e.Get(1)
			_ = W1667
			tmp8769 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1667)

			if True == tmp8769 {
				tmp8758 := MakeNative(func(__e *ControlFlow) {
					W1674 := __e.Get(1)
					_ = W1674
					tmp8760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1674)

					if True == tmp8760 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W1674)
						return
					}

				}, 1)

				tmp8761 := MakeNative(func(__e *ControlFlow) {
					W1675 := __e.Get(1)
					_ = W1675
					tmp8765 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1675)

					if True == tmp8765 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp8762 := MakeNative(func(__e *ControlFlow) {
							W1676 := __e.Get(1)
							_ = W1676
							__e.TailApply(PrimFunc(symshen_4comb), W1676, Nil)
							return
						}, 1)

						tmp8763 := Call(__e, PrimFunc(symshen_4in_1_6), W1675)

						__e.TailApply(tmp8762, tmp8763)
						return

					}

				}, 1)

				tmp8766 := Call(__e, PrimFunc(sym_5e_6), V1666)

				tmp8767 := Call(__e, tmp8761, tmp8766)

				__e.TailApply(tmp8758, tmp8767)
				return

			} else {
				__e.Return(W1667)
				return
			}

		}, 1)

		tmp8770 := MakeNative(func(__e *ControlFlow) {
			W1668 := __e.Get(1)
			_ = W1668
			tmp8785 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1668)

			if True == tmp8785 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp8771 := MakeNative(func(__e *ControlFlow) {
					W1669 := __e.Get(1)
					_ = W1669
					tmp8772 := MakeNative(func(__e *ControlFlow) {
						W1670 := __e.Get(1)
						_ = W1670
						tmp8773 := MakeNative(func(__e *ControlFlow) {
							W1671 := __e.Get(1)
							_ = W1671
							tmp8780 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1671)

							if True == tmp8780 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp8774 := MakeNative(func(__e *ControlFlow) {
									W1672 := __e.Get(1)
									_ = W1672
									tmp8775 := MakeNative(func(__e *ControlFlow) {
										W1673 := __e.Get(1)
										_ = W1673
										tmp8776 := PrimCons(W1669, W1672)

										__e.TailApply(PrimFunc(symshen_4comb), W1673, tmp8776)
										return

									}, 1)

									tmp8777 := Call(__e, PrimFunc(symshen_4in_1_6), W1671)

									__e.TailApply(tmp8775, tmp8777)
									return

								}, 1)

								tmp8778 := Call(__e, PrimFunc(symshen_4_5_1out), W1671)

								__e.TailApply(tmp8774, tmp8778)
								return

							}

						}, 1)

						tmp8781 := Call(__e, PrimFunc(symshen_4_5head_6), W1670)

						__e.TailApply(tmp8773, tmp8781)
						return

					}, 1)

					tmp8782 := Call(__e, PrimFunc(symshen_4in_1_6), W1668)

					__e.TailApply(tmp8772, tmp8782)
					return

				}, 1)

				tmp8783 := Call(__e, PrimFunc(symshen_4_5_1out), W1668)

				__e.TailApply(tmp8771, tmp8783)
				return

			}

		}, 1)

		tmp8786 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1666)

		tmp8787 := Call(__e, tmp8770, tmp8786)

		__e.TailApply(tmp8757, tmp8787)
		return

	}, 1)

	tmp8788 := Call(__e, ns2_1set, symshen_4_5head_6, tmp8756)

	_ = tmp8788

	tmp8789 := MakeNative(func(__e *ControlFlow) {
		V1677 := __e.Get(1)
		_ = V1677
		tmp8790 := MakeNative(func(__e *ControlFlow) {
			W1678 := __e.Get(1)
			_ = W1678
			tmp8978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1678)

			if True == tmp8978 {
				tmp8791 := MakeNative(func(__e *ControlFlow) {
					W1681 := __e.Get(1)
					_ = W1681
					tmp8965 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1681)

					if True == tmp8965 {
						tmp8792 := MakeNative(func(__e *ControlFlow) {
							W1684 := __e.Get(1)
							_ = W1684
							tmp8926 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1684)

							if True == tmp8926 {
								tmp8793 := MakeNative(func(__e *ControlFlow) {
									W1696 := __e.Get(1)
									_ = W1696
									tmp8896 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1696)

									if True == tmp8896 {
										tmp8794 := MakeNative(func(__e *ControlFlow) {
											W1705 := __e.Get(1)
											_ = W1705
											tmp8866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1705)

											if True == tmp8866 {
												tmp8795 := MakeNative(func(__e *ControlFlow) {
													W1714 := __e.Get(1)
													_ = W1714
													tmp8832 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1714)

													if True == tmp8832 {
														tmp8796 := MakeNative(func(__e *ControlFlow) {
															W1724 := __e.Get(1)
															_ = W1724
															tmp8798 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1724)

															if True == tmp8798 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																__e.Return(W1724)
																return
															}

														}, 1)

														tmp8830 := Call(__e, PrimFunc(symshen_4ccons_2), V1677)

														var ifres8799 Obj

														if True == tmp8830 {
															tmp8800 := MakeNative(func(__e *ControlFlow) {
																W1725 := __e.Get(1)
																_ = W1725
																tmp8801 := MakeNative(func(__e *ControlFlow) {
																	W1726 := __e.Get(1)
																	_ = W1726
																	tmp8825 := Call(__e, PrimFunc(symshen_4hds_a_2), W1725, symmode)

																	if True == tmp8825 {
																		tmp8802 := MakeNative(func(__e *ControlFlow) {
																			W1727 := __e.Get(1)
																			_ = W1727
																			tmp8803 := MakeNative(func(__e *ControlFlow) {
																				W1728 := __e.Get(1)
																				_ = W1728
																				tmp8821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1728)

																				if True == tmp8821 {
																					__e.TailApply(PrimFunc(symshen_4parse_1failure))
																					return
																				} else {
																					tmp8804 := MakeNative(func(__e *ControlFlow) {
																						W1729 := __e.Get(1)
																						_ = W1729
																						tmp8805 := MakeNative(func(__e *ControlFlow) {
																							W1730 := __e.Get(1)
																							_ = W1730
																							tmp8817 := Call(__e, PrimFunc(symshen_4hds_a_2), W1730, sym_1)

																							if True == tmp8817 {
																								tmp8806 := MakeNative(func(__e *ControlFlow) {
																									W1731 := __e.Get(1)
																									_ = W1731
																									tmp8807 := MakeNative(func(__e *ControlFlow) {
																										W1732 := __e.Get(1)
																										_ = W1732
																										tmp8813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1732)

																										if True == tmp8813 {
																											__e.TailApply(PrimFunc(symshen_4parse_1failure))
																											return
																										} else {
																											tmp8808 := MakeNative(func(__e *ControlFlow) {
																												W1733 := __e.Get(1)
																												_ = W1733
																												tmp8809 := PrimCons(W1729, Nil)

																												tmp8810 := PrimCons(symshen_4_1m, tmp8809)

																												__e.TailApply(PrimFunc(symshen_4comb), W1726, tmp8810)
																												return

																											}, 1)

																											tmp8811 := Call(__e, PrimFunc(symshen_4in_1_6), W1732)

																											__e.TailApply(tmp8808, tmp8811)
																											return

																										}

																									}, 1)

																									tmp8814 := Call(__e, PrimFunc(sym_5end_6), W1731)

																									__e.TailApply(tmp8807, tmp8814)
																									return

																								}, 1)

																								tmp8815 := Call(__e, PrimFunc(symtail), W1730)

																								__e.TailApply(tmp8806, tmp8815)
																								return

																							} else {
																								__e.TailApply(PrimFunc(symshen_4parse_1failure))
																								return
																							}

																						}, 1)

																						tmp8818 := Call(__e, PrimFunc(symshen_4in_1_6), W1728)

																						__e.TailApply(tmp8805, tmp8818)
																						return

																					}, 1)

																					tmp8819 := Call(__e, PrimFunc(symshen_4_5_1out), W1728)

																					__e.TailApply(tmp8804, tmp8819)
																					return

																				}

																			}, 1)

																			tmp8822 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1727)

																			__e.TailApply(tmp8803, tmp8822)
																			return

																		}, 1)

																		tmp8823 := Call(__e, PrimFunc(symtail), W1725)

																		__e.TailApply(tmp8802, tmp8823)
																		return

																	} else {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	}

																}, 1)

																tmp8826 := Call(__e, PrimFunc(symtail), V1677)

																__e.TailApply(tmp8801, tmp8826)
																return

															}, 1)

															tmp8827 := Call(__e, PrimFunc(symhead), V1677)

															tmp8828 := Call(__e, tmp8800, tmp8827)

															ifres8799 = tmp8828

														} else {
															tmp8829 := Call(__e, PrimFunc(symshen_4parse_1failure))

															ifres8799 = tmp8829

														}

														__e.TailApply(tmp8796, ifres8799)
														return

													} else {
														__e.Return(W1714)
														return
													}

												}, 1)

												tmp8864 := Call(__e, PrimFunc(symshen_4ccons_2), V1677)

												var ifres8833 Obj

												if True == tmp8864 {
													tmp8834 := MakeNative(func(__e *ControlFlow) {
														W1715 := __e.Get(1)
														_ = W1715
														tmp8835 := MakeNative(func(__e *ControlFlow) {
															W1716 := __e.Get(1)
															_ = W1716
															tmp8859 := Call(__e, PrimFunc(symshen_4hds_a_2), W1715, symmode)

															if True == tmp8859 {
																tmp8836 := MakeNative(func(__e *ControlFlow) {
																	W1717 := __e.Get(1)
																	_ = W1717
																	tmp8837 := MakeNative(func(__e *ControlFlow) {
																		W1718 := __e.Get(1)
																		_ = W1718
																		tmp8855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1718)

																		if True == tmp8855 {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		} else {
																			tmp8838 := MakeNative(func(__e *ControlFlow) {
																				W1719 := __e.Get(1)
																				_ = W1719
																				tmp8839 := MakeNative(func(__e *ControlFlow) {
																					W1720 := __e.Get(1)
																					_ = W1720
																					tmp8851 := Call(__e, PrimFunc(symshen_4hds_a_2), W1720, sym_7)

																					if True == tmp8851 {
																						tmp8840 := MakeNative(func(__e *ControlFlow) {
																							W1721 := __e.Get(1)
																							_ = W1721
																							tmp8841 := MakeNative(func(__e *ControlFlow) {
																								W1722 := __e.Get(1)
																								_ = W1722
																								tmp8847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1722)

																								if True == tmp8847 {
																									__e.TailApply(PrimFunc(symshen_4parse_1failure))
																									return
																								} else {
																									tmp8842 := MakeNative(func(__e *ControlFlow) {
																										W1723 := __e.Get(1)
																										_ = W1723
																										tmp8843 := PrimCons(W1719, Nil)

																										tmp8844 := PrimCons(symshen_4_7m, tmp8843)

																										__e.TailApply(PrimFunc(symshen_4comb), W1716, tmp8844)
																										return

																									}, 1)

																									tmp8845 := Call(__e, PrimFunc(symshen_4in_1_6), W1722)

																									__e.TailApply(tmp8842, tmp8845)
																									return

																								}

																							}, 1)

																							tmp8848 := Call(__e, PrimFunc(sym_5end_6), W1721)

																							__e.TailApply(tmp8841, tmp8848)
																							return

																						}, 1)

																						tmp8849 := Call(__e, PrimFunc(symtail), W1720)

																						__e.TailApply(tmp8840, tmp8849)
																						return

																					} else {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					}

																				}, 1)

																				tmp8852 := Call(__e, PrimFunc(symshen_4in_1_6), W1718)

																				__e.TailApply(tmp8839, tmp8852)
																				return

																			}, 1)

																			tmp8853 := Call(__e, PrimFunc(symshen_4_5_1out), W1718)

																			__e.TailApply(tmp8838, tmp8853)
																			return

																		}

																	}, 1)

																	tmp8856 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1717)

																	__e.TailApply(tmp8837, tmp8856)
																	return

																}, 1)

																tmp8857 := Call(__e, PrimFunc(symtail), W1715)

																__e.TailApply(tmp8836, tmp8857)
																return

															} else {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															}

														}, 1)

														tmp8860 := Call(__e, PrimFunc(symtail), V1677)

														__e.TailApply(tmp8835, tmp8860)
														return

													}, 1)

													tmp8861 := Call(__e, PrimFunc(symhead), V1677)

													tmp8862 := Call(__e, tmp8834, tmp8861)

													ifres8833 = tmp8862

												} else {
													tmp8863 := Call(__e, PrimFunc(symshen_4parse_1failure))

													ifres8833 = tmp8863

												}

												__e.TailApply(tmp8795, ifres8833)
												return

											} else {
												__e.Return(W1705)
												return
											}

										}, 1)

										tmp8894 := Call(__e, PrimFunc(symshen_4ccons_2), V1677)

										var ifres8867 Obj

										if True == tmp8894 {
											tmp8868 := MakeNative(func(__e *ControlFlow) {
												W1706 := __e.Get(1)
												_ = W1706
												tmp8869 := MakeNative(func(__e *ControlFlow) {
													W1707 := __e.Get(1)
													_ = W1707
													tmp8889 := Call(__e, PrimFunc(symshen_4hds_a_2), W1706, sym_1)

													if True == tmp8889 {
														tmp8870 := MakeNative(func(__e *ControlFlow) {
															W1708 := __e.Get(1)
															_ = W1708
															tmp8871 := MakeNative(func(__e *ControlFlow) {
																W1709 := __e.Get(1)
																_ = W1709
																tmp8885 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1709)

																if True == tmp8885 {
																	__e.TailApply(PrimFunc(symshen_4parse_1failure))
																	return
																} else {
																	tmp8872 := MakeNative(func(__e *ControlFlow) {
																		W1710 := __e.Get(1)
																		_ = W1710
																		tmp8873 := MakeNative(func(__e *ControlFlow) {
																			W1711 := __e.Get(1)
																			_ = W1711
																			tmp8874 := MakeNative(func(__e *ControlFlow) {
																				W1712 := __e.Get(1)
																				_ = W1712
																				tmp8880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1712)

																				if True == tmp8880 {
																					__e.TailApply(PrimFunc(symshen_4parse_1failure))
																					return
																				} else {
																					tmp8875 := MakeNative(func(__e *ControlFlow) {
																						W1713 := __e.Get(1)
																						_ = W1713
																						tmp8876 := PrimCons(W1710, Nil)

																						tmp8877 := PrimCons(symshen_4_1m, tmp8876)

																						__e.TailApply(PrimFunc(symshen_4comb), W1707, tmp8877)
																						return

																					}, 1)

																					tmp8878 := Call(__e, PrimFunc(symshen_4in_1_6), W1712)

																					__e.TailApply(tmp8875, tmp8878)
																					return

																				}

																			}, 1)

																			tmp8881 := Call(__e, PrimFunc(sym_5end_6), W1711)

																			__e.TailApply(tmp8874, tmp8881)
																			return

																		}, 1)

																		tmp8882 := Call(__e, PrimFunc(symshen_4in_1_6), W1709)

																		__e.TailApply(tmp8873, tmp8882)
																		return

																	}, 1)

																	tmp8883 := Call(__e, PrimFunc(symshen_4_5_1out), W1709)

																	__e.TailApply(tmp8872, tmp8883)
																	return

																}

															}, 1)

															tmp8886 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1708)

															__e.TailApply(tmp8871, tmp8886)
															return

														}, 1)

														tmp8887 := Call(__e, PrimFunc(symtail), W1706)

														__e.TailApply(tmp8870, tmp8887)
														return

													} else {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp8890 := Call(__e, PrimFunc(symtail), V1677)

												__e.TailApply(tmp8869, tmp8890)
												return

											}, 1)

											tmp8891 := Call(__e, PrimFunc(symhead), V1677)

											tmp8892 := Call(__e, tmp8868, tmp8891)

											ifres8867 = tmp8892

										} else {
											tmp8893 := Call(__e, PrimFunc(symshen_4parse_1failure))

											ifres8867 = tmp8893

										}

										__e.TailApply(tmp8794, ifres8867)
										return

									} else {
										__e.Return(W1696)
										return
									}

								}, 1)

								tmp8924 := Call(__e, PrimFunc(symshen_4ccons_2), V1677)

								var ifres8897 Obj

								if True == tmp8924 {
									tmp8898 := MakeNative(func(__e *ControlFlow) {
										W1697 := __e.Get(1)
										_ = W1697
										tmp8899 := MakeNative(func(__e *ControlFlow) {
											W1698 := __e.Get(1)
											_ = W1698
											tmp8919 := Call(__e, PrimFunc(symshen_4hds_a_2), W1697, sym_7)

											if True == tmp8919 {
												tmp8900 := MakeNative(func(__e *ControlFlow) {
													W1699 := __e.Get(1)
													_ = W1699
													tmp8901 := MakeNative(func(__e *ControlFlow) {
														W1700 := __e.Get(1)
														_ = W1700
														tmp8915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1700)

														if True == tmp8915 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp8902 := MakeNative(func(__e *ControlFlow) {
																W1701 := __e.Get(1)
																_ = W1701
																tmp8903 := MakeNative(func(__e *ControlFlow) {
																	W1702 := __e.Get(1)
																	_ = W1702
																	tmp8904 := MakeNative(func(__e *ControlFlow) {
																		W1703 := __e.Get(1)
																		_ = W1703
																		tmp8910 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1703)

																		if True == tmp8910 {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		} else {
																			tmp8905 := MakeNative(func(__e *ControlFlow) {
																				W1704 := __e.Get(1)
																				_ = W1704
																				tmp8906 := PrimCons(W1701, Nil)

																				tmp8907 := PrimCons(symshen_4_7m, tmp8906)

																				__e.TailApply(PrimFunc(symshen_4comb), W1698, tmp8907)
																				return

																			}, 1)

																			tmp8908 := Call(__e, PrimFunc(symshen_4in_1_6), W1703)

																			__e.TailApply(tmp8905, tmp8908)
																			return

																		}

																	}, 1)

																	tmp8911 := Call(__e, PrimFunc(sym_5end_6), W1702)

																	__e.TailApply(tmp8904, tmp8911)
																	return

																}, 1)

																tmp8912 := Call(__e, PrimFunc(symshen_4in_1_6), W1700)

																__e.TailApply(tmp8903, tmp8912)
																return

															}, 1)

															tmp8913 := Call(__e, PrimFunc(symshen_4_5_1out), W1700)

															__e.TailApply(tmp8902, tmp8913)
															return

														}

													}, 1)

													tmp8916 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1699)

													__e.TailApply(tmp8901, tmp8916)
													return

												}, 1)

												tmp8917 := Call(__e, PrimFunc(symtail), W1697)

												__e.TailApply(tmp8900, tmp8917)
												return

											} else {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp8920 := Call(__e, PrimFunc(symtail), V1677)

										__e.TailApply(tmp8899, tmp8920)
										return

									}, 1)

									tmp8921 := Call(__e, PrimFunc(symhead), V1677)

									tmp8922 := Call(__e, tmp8898, tmp8921)

									ifres8897 = tmp8922

								} else {
									tmp8923 := Call(__e, PrimFunc(symshen_4parse_1failure))

									ifres8897 = tmp8923

								}

								__e.TailApply(tmp8793, ifres8897)
								return

							} else {
								__e.Return(W1684)
								return
							}

						}, 1)

						tmp8963 := Call(__e, PrimFunc(symshen_4ccons_2), V1677)

						var ifres8927 Obj

						if True == tmp8963 {
							tmp8928 := MakeNative(func(__e *ControlFlow) {
								W1685 := __e.Get(1)
								_ = W1685
								tmp8929 := MakeNative(func(__e *ControlFlow) {
									W1686 := __e.Get(1)
									_ = W1686
									tmp8958 := Call(__e, PrimFunc(symshen_4hds_a_2), W1685, symcons)

									if True == tmp8958 {
										tmp8930 := MakeNative(func(__e *ControlFlow) {
											W1687 := __e.Get(1)
											_ = W1687
											tmp8931 := MakeNative(func(__e *ControlFlow) {
												W1688 := __e.Get(1)
												_ = W1688
												tmp8954 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1688)

												if True == tmp8954 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp8932 := MakeNative(func(__e *ControlFlow) {
														W1689 := __e.Get(1)
														_ = W1689
														tmp8933 := MakeNative(func(__e *ControlFlow) {
															W1690 := __e.Get(1)
															_ = W1690
															tmp8934 := MakeNative(func(__e *ControlFlow) {
																W1691 := __e.Get(1)
																_ = W1691
																tmp8949 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1691)

																if True == tmp8949 {
																	__e.TailApply(PrimFunc(symshen_4parse_1failure))
																	return
																} else {
																	tmp8935 := MakeNative(func(__e *ControlFlow) {
																		W1692 := __e.Get(1)
																		_ = W1692
																		tmp8936 := MakeNative(func(__e *ControlFlow) {
																			W1693 := __e.Get(1)
																			_ = W1693
																			tmp8937 := MakeNative(func(__e *ControlFlow) {
																				W1694 := __e.Get(1)
																				_ = W1694
																				tmp8944 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1694)

																				if True == tmp8944 {
																					__e.TailApply(PrimFunc(symshen_4parse_1failure))
																					return
																				} else {
																					tmp8938 := MakeNative(func(__e *ControlFlow) {
																						W1695 := __e.Get(1)
																						_ = W1695
																						tmp8939 := PrimCons(W1692, Nil)

																						tmp8940 := PrimCons(W1689, tmp8939)

																						tmp8941 := PrimCons(symcons, tmp8940)

																						__e.TailApply(PrimFunc(symshen_4comb), W1686, tmp8941)
																						return

																					}, 1)

																					tmp8942 := Call(__e, PrimFunc(symshen_4in_1_6), W1694)

																					__e.TailApply(tmp8938, tmp8942)
																					return

																				}

																			}, 1)

																			tmp8945 := Call(__e, PrimFunc(sym_5end_6), W1693)

																			__e.TailApply(tmp8937, tmp8945)
																			return

																		}, 1)

																		tmp8946 := Call(__e, PrimFunc(symshen_4in_1_6), W1691)

																		__e.TailApply(tmp8936, tmp8946)
																		return

																	}, 1)

																	tmp8947 := Call(__e, PrimFunc(symshen_4_5_1out), W1691)

																	__e.TailApply(tmp8935, tmp8947)
																	return

																}

															}, 1)

															tmp8950 := Call(__e, PrimFunc(symshen_4_5hterm2_6), W1690)

															__e.TailApply(tmp8934, tmp8950)
															return

														}, 1)

														tmp8951 := Call(__e, PrimFunc(symshen_4in_1_6), W1688)

														__e.TailApply(tmp8933, tmp8951)
														return

													}, 1)

													tmp8952 := Call(__e, PrimFunc(symshen_4_5_1out), W1688)

													__e.TailApply(tmp8932, tmp8952)
													return

												}

											}, 1)

											tmp8955 := Call(__e, PrimFunc(symshen_4_5hterm1_6), W1687)

											__e.TailApply(tmp8931, tmp8955)
											return

										}, 1)

										tmp8956 := Call(__e, PrimFunc(symtail), W1685)

										__e.TailApply(tmp8930, tmp8956)
										return

									} else {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp8959 := Call(__e, PrimFunc(symtail), V1677)

								__e.TailApply(tmp8929, tmp8959)
								return

							}, 1)

							tmp8960 := Call(__e, PrimFunc(symhead), V1677)

							tmp8961 := Call(__e, tmp8928, tmp8960)

							ifres8927 = tmp8961

						} else {
							tmp8962 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres8927 = tmp8962

						}

						__e.TailApply(tmp8792, ifres8927)
						return

					} else {
						__e.Return(W1681)
						return
					}

				}, 1)

				tmp8976 := PrimIsPair(V1677)

				var ifres8966 Obj

				if True == tmp8976 {
					tmp8967 := MakeNative(func(__e *ControlFlow) {
						W1682 := __e.Get(1)
						_ = W1682
						tmp8968 := MakeNative(func(__e *ControlFlow) {
							W1683 := __e.Get(1)
							_ = W1683
							tmp8970 := PrimIntern(MakeString(":"))

							tmp8971 := PrimEqual(W1682, tmp8970)

							if True == tmp8971 {
								__e.TailApply(PrimFunc(symshen_4comb), W1683, W1682)
								return
							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp8972 := Call(__e, PrimFunc(symtail), V1677)

						__e.TailApply(tmp8968, tmp8972)
						return

					}, 1)

					tmp8973 := Call(__e, PrimFunc(symhead), V1677)

					tmp8974 := Call(__e, tmp8967, tmp8973)

					ifres8966 = tmp8974

				} else {
					tmp8975 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres8966 = tmp8975

				}

				__e.TailApply(tmp8791, ifres8966)
				return

			} else {
				__e.Return(W1678)
				return
			}

		}, 1)

		tmp8992 := PrimIsPair(V1677)

		var ifres8979 Obj

		if True == tmp8992 {
			tmp8980 := MakeNative(func(__e *ControlFlow) {
				W1679 := __e.Get(1)
				_ = W1679
				tmp8981 := MakeNative(func(__e *ControlFlow) {
					W1680 := __e.Get(1)
					_ = W1680
					tmp8987 := Call(__e, PrimFunc(symatom_2), W1679)

					var ifres8983 Obj

					if True == tmp8987 {
						tmp8985 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), W1679)

						tmp8986 := PrimNot(tmp8985)

						var ifres8984 Obj

						if True == tmp8986 {
							ifres8984 = True

						} else {
							ifres8984 = False

						}

						ifres8983 = ifres8984

					} else {
						ifres8983 = False

					}

					if True == ifres8983 {
						__e.TailApply(PrimFunc(symshen_4comb), W1680, W1679)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp8988 := Call(__e, PrimFunc(symtail), V1677)

				__e.TailApply(tmp8981, tmp8988)
				return

			}, 1)

			tmp8989 := Call(__e, PrimFunc(symhead), V1677)

			tmp8990 := Call(__e, tmp8980, tmp8989)

			ifres8979 = tmp8990

		} else {
			tmp8991 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres8979 = tmp8991

		}

		__e.TailApply(tmp8790, ifres8979)
		return

	}, 1)

	tmp8993 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp8789)

	_ = tmp8993

	tmp8994 := MakeNative(func(__e *ControlFlow) {
		V1734 := __e.Get(1)
		_ = V1734
		tmp8995 := PrimIntern(MakeString(";"))

		tmp8996 := PrimCons(sym_5_1_1, Nil)

		tmp8997 := PrimCons(tmp8995, tmp8996)

		__e.TailApply(PrimFunc(symelement_2), V1734, tmp8997)
		return

	}, 1)

	tmp8998 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp8994)

	_ = tmp8998

	tmp8999 := MakeNative(func(__e *ControlFlow) {
		V1735 := __e.Get(1)
		_ = V1735
		tmp9012 := PrimIsSymbol(V1735)

		if True == tmp9012 {
			__e.Return(True)
			return
		} else {
			tmp9010 := PrimIsString(V1735)

			var ifres9001 Obj

			if True == tmp9010 {
				ifres9001 = True

			} else {
				tmp9009 := Call(__e, PrimFunc(symboolean_2), V1735)

				var ifres9003 Obj

				if True == tmp9009 {
					ifres9003 = True

				} else {
					tmp9008 := PrimIsNumber(V1735)

					var ifres9005 Obj

					if True == tmp9008 {
						ifres9005 = True

					} else {
						tmp9007 := Call(__e, PrimFunc(symempty_2), V1735)

						var ifres9006 Obj

						if True == tmp9007 {
							ifres9006 = True

						} else {
							ifres9006 = False

						}

						ifres9005 = ifres9006

					}

					var ifres9004 Obj

					if True == ifres9005 {
						ifres9004 = True

					} else {
						ifres9004 = False

					}

					ifres9003 = ifres9004

				}

				var ifres9002 Obj

				if True == ifres9003 {
					ifres9002 = True

				} else {
					ifres9002 = False

				}

				ifres9001 = ifres9002

			}

			if True == ifres9001 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp9013 := Call(__e, ns2_1set, symatom_2, tmp8999)

	_ = tmp9013

	tmp9014 := MakeNative(func(__e *ControlFlow) {
		V1736 := __e.Get(1)
		_ = V1736
		tmp9015 := MakeNative(func(__e *ControlFlow) {
			W1737 := __e.Get(1)
			_ = W1737
			tmp9017 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1737)

			if True == tmp9017 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W1737)
				return
			}

		}, 1)

		tmp9018 := MakeNative(func(__e *ControlFlow) {
			W1738 := __e.Get(1)
			_ = W1738
			tmp9024 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1738)

			if True == tmp9024 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9019 := MakeNative(func(__e *ControlFlow) {
					W1739 := __e.Get(1)
					_ = W1739
					tmp9020 := MakeNative(func(__e *ControlFlow) {
						W1740 := __e.Get(1)
						_ = W1740
						__e.TailApply(PrimFunc(symshen_4comb), W1740, W1739)
						return
					}, 1)

					tmp9021 := Call(__e, PrimFunc(symshen_4in_1_6), W1738)

					__e.TailApply(tmp9020, tmp9021)
					return

				}, 1)

				tmp9022 := Call(__e, PrimFunc(symshen_4_5_1out), W1738)

				__e.TailApply(tmp9019, tmp9022)
				return

			}

		}, 1)

		tmp9025 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1736)

		tmp9026 := Call(__e, tmp9018, tmp9025)

		__e.TailApply(tmp9015, tmp9026)
		return

	}, 1)

	tmp9027 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp9014)

	_ = tmp9027

	tmp9028 := MakeNative(func(__e *ControlFlow) {
		V1741 := __e.Get(1)
		_ = V1741
		tmp9029 := MakeNative(func(__e *ControlFlow) {
			W1742 := __e.Get(1)
			_ = W1742
			tmp9031 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1742)

			if True == tmp9031 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W1742)
				return
			}

		}, 1)

		tmp9032 := MakeNative(func(__e *ControlFlow) {
			W1743 := __e.Get(1)
			_ = W1743
			tmp9038 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1743)

			if True == tmp9038 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9033 := MakeNative(func(__e *ControlFlow) {
					W1744 := __e.Get(1)
					_ = W1744
					tmp9034 := MakeNative(func(__e *ControlFlow) {
						W1745 := __e.Get(1)
						_ = W1745
						__e.TailApply(PrimFunc(symshen_4comb), W1745, W1744)
						return
					}, 1)

					tmp9035 := Call(__e, PrimFunc(symshen_4in_1_6), W1743)

					__e.TailApply(tmp9034, tmp9035)
					return

				}, 1)

				tmp9036 := Call(__e, PrimFunc(symshen_4_5_1out), W1743)

				__e.TailApply(tmp9033, tmp9036)
				return

			}

		}, 1)

		tmp9039 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1741)

		tmp9040 := Call(__e, tmp9032, tmp9039)

		__e.TailApply(tmp9029, tmp9040)
		return

	}, 1)

	tmp9041 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp9028)

	_ = tmp9041

	tmp9042 := MakeNative(func(__e *ControlFlow) {
		V1746 := __e.Get(1)
		_ = V1746
		tmp9043 := MakeNative(func(__e *ControlFlow) {
			W1747 := __e.Get(1)
			_ = W1747
			tmp9055 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1747)

			if True == tmp9055 {
				tmp9044 := MakeNative(func(__e *ControlFlow) {
					W1754 := __e.Get(1)
					_ = W1754
					tmp9046 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1754)

					if True == tmp9046 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W1754)
						return
					}

				}, 1)

				tmp9047 := MakeNative(func(__e *ControlFlow) {
					W1755 := __e.Get(1)
					_ = W1755
					tmp9051 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1755)

					if True == tmp9051 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp9048 := MakeNative(func(__e *ControlFlow) {
							W1756 := __e.Get(1)
							_ = W1756
							__e.TailApply(PrimFunc(symshen_4comb), W1756, Nil)
							return
						}, 1)

						tmp9049 := Call(__e, PrimFunc(symshen_4in_1_6), W1755)

						__e.TailApply(tmp9048, tmp9049)
						return

					}

				}, 1)

				tmp9052 := Call(__e, PrimFunc(sym_5e_6), V1746)

				tmp9053 := Call(__e, tmp9047, tmp9052)

				__e.TailApply(tmp9044, tmp9053)
				return

			} else {
				__e.Return(W1747)
				return
			}

		}, 1)

		tmp9056 := MakeNative(func(__e *ControlFlow) {
			W1748 := __e.Get(1)
			_ = W1748
			tmp9071 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1748)

			if True == tmp9071 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9057 := MakeNative(func(__e *ControlFlow) {
					W1749 := __e.Get(1)
					_ = W1749
					tmp9058 := MakeNative(func(__e *ControlFlow) {
						W1750 := __e.Get(1)
						_ = W1750
						tmp9059 := MakeNative(func(__e *ControlFlow) {
							W1751 := __e.Get(1)
							_ = W1751
							tmp9066 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1751)

							if True == tmp9066 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp9060 := MakeNative(func(__e *ControlFlow) {
									W1752 := __e.Get(1)
									_ = W1752
									tmp9061 := MakeNative(func(__e *ControlFlow) {
										W1753 := __e.Get(1)
										_ = W1753
										tmp9062 := PrimCons(W1749, W1752)

										__e.TailApply(PrimFunc(symshen_4comb), W1753, tmp9062)
										return

									}, 1)

									tmp9063 := Call(__e, PrimFunc(symshen_4in_1_6), W1751)

									__e.TailApply(tmp9061, tmp9063)
									return

								}, 1)

								tmp9064 := Call(__e, PrimFunc(symshen_4_5_1out), W1751)

								__e.TailApply(tmp9060, tmp9064)
								return

							}

						}, 1)

						tmp9067 := Call(__e, PrimFunc(symshen_4_5body_6), W1750)

						__e.TailApply(tmp9059, tmp9067)
						return

					}, 1)

					tmp9068 := Call(__e, PrimFunc(symshen_4in_1_6), W1748)

					__e.TailApply(tmp9058, tmp9068)
					return

				}, 1)

				tmp9069 := Call(__e, PrimFunc(symshen_4_5_1out), W1748)

				__e.TailApply(tmp9057, tmp9069)
				return

			}

		}, 1)

		tmp9072 := Call(__e, PrimFunc(symshen_4_5literal_6), V1746)

		tmp9073 := Call(__e, tmp9056, tmp9072)

		__e.TailApply(tmp9043, tmp9073)
		return

	}, 1)

	tmp9074 := Call(__e, ns2_1set, symshen_4_5body_6, tmp9042)

	_ = tmp9074

	tmp9075 := MakeNative(func(__e *ControlFlow) {
		V1757 := __e.Get(1)
		_ = V1757
		tmp9076 := MakeNative(func(__e *ControlFlow) {
			W1758 := __e.Get(1)
			_ = W1758
			tmp9103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1758)

			if True == tmp9103 {
				tmp9077 := MakeNative(func(__e *ControlFlow) {
					W1760 := __e.Get(1)
					_ = W1760
					tmp9079 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1760)

					if True == tmp9079 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W1760)
						return
					}

				}, 1)

				tmp9101 := Call(__e, PrimFunc(symshen_4ccons_2), V1757)

				var ifres9080 Obj

				if True == tmp9101 {
					tmp9081 := MakeNative(func(__e *ControlFlow) {
						W1761 := __e.Get(1)
						_ = W1761
						tmp9082 := MakeNative(func(__e *ControlFlow) {
							W1762 := __e.Get(1)
							_ = W1762
							tmp9083 := MakeNative(func(__e *ControlFlow) {
								W1763 := __e.Get(1)
								_ = W1763
								tmp9095 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1763)

								if True == tmp9095 {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								} else {
									tmp9084 := MakeNative(func(__e *ControlFlow) {
										W1764 := __e.Get(1)
										_ = W1764
										tmp9085 := MakeNative(func(__e *ControlFlow) {
											W1765 := __e.Get(1)
											_ = W1765
											tmp9086 := MakeNative(func(__e *ControlFlow) {
												W1766 := __e.Get(1)
												_ = W1766
												tmp9090 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1766)

												if True == tmp9090 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp9087 := MakeNative(func(__e *ControlFlow) {
														W1767 := __e.Get(1)
														_ = W1767
														__e.TailApply(PrimFunc(symshen_4comb), W1762, W1764)
														return
													}, 1)

													tmp9088 := Call(__e, PrimFunc(symshen_4in_1_6), W1766)

													__e.TailApply(tmp9087, tmp9088)
													return

												}

											}, 1)

											tmp9091 := Call(__e, PrimFunc(sym_5end_6), W1765)

											__e.TailApply(tmp9086, tmp9091)
											return

										}, 1)

										tmp9092 := Call(__e, PrimFunc(symshen_4in_1_6), W1763)

										__e.TailApply(tmp9085, tmp9092)
										return

									}, 1)

									tmp9093 := Call(__e, PrimFunc(symshen_4_5_1out), W1763)

									__e.TailApply(tmp9084, tmp9093)
									return

								}

							}, 1)

							tmp9096 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1761)

							__e.TailApply(tmp9083, tmp9096)
							return

						}, 1)

						tmp9097 := Call(__e, PrimFunc(symtail), V1757)

						__e.TailApply(tmp9082, tmp9097)
						return

					}, 1)

					tmp9098 := Call(__e, PrimFunc(symhead), V1757)

					tmp9099 := Call(__e, tmp9081, tmp9098)

					ifres9080 = tmp9099

				} else {
					tmp9100 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres9080 = tmp9100

				}

				__e.TailApply(tmp9077, ifres9080)
				return

			} else {
				__e.Return(W1758)
				return
			}

		}, 1)

		tmp9109 := Call(__e, PrimFunc(symshen_4hds_a_2), V1757, sym_b)

		var ifres9104 Obj

		if True == tmp9109 {
			tmp9105 := MakeNative(func(__e *ControlFlow) {
				W1759 := __e.Get(1)
				_ = W1759
				__e.TailApply(PrimFunc(symshen_4comb), W1759, sym_b)
				return
			}, 1)

			tmp9106 := Call(__e, PrimFunc(symtail), V1757)

			tmp9107 := Call(__e, tmp9105, tmp9106)

			ifres9104 = tmp9107

		} else {
			tmp9108 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres9104 = tmp9108

		}

		__e.TailApply(tmp9076, ifres9104)
		return

	}, 1)

	tmp9110 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp9075)

	_ = tmp9110

	tmp9111 := MakeNative(func(__e *ControlFlow) {
		V1768 := __e.Get(1)
		_ = V1768
		tmp9112 := MakeNative(func(__e *ControlFlow) {
			W1769 := __e.Get(1)
			_ = W1769
			tmp9124 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1769)

			if True == tmp9124 {
				tmp9113 := MakeNative(func(__e *ControlFlow) {
					W1776 := __e.Get(1)
					_ = W1776
					tmp9115 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1776)

					if True == tmp9115 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W1776)
						return
					}

				}, 1)

				tmp9116 := MakeNative(func(__e *ControlFlow) {
					W1777 := __e.Get(1)
					_ = W1777
					tmp9120 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1777)

					if True == tmp9120 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp9117 := MakeNative(func(__e *ControlFlow) {
							W1778 := __e.Get(1)
							_ = W1778
							__e.TailApply(PrimFunc(symshen_4comb), W1778, Nil)
							return
						}, 1)

						tmp9118 := Call(__e, PrimFunc(symshen_4in_1_6), W1777)

						__e.TailApply(tmp9117, tmp9118)
						return

					}

				}, 1)

				tmp9121 := Call(__e, PrimFunc(sym_5e_6), V1768)

				tmp9122 := Call(__e, tmp9116, tmp9121)

				__e.TailApply(tmp9113, tmp9122)
				return

			} else {
				__e.Return(W1769)
				return
			}

		}, 1)

		tmp9125 := MakeNative(func(__e *ControlFlow) {
			W1770 := __e.Get(1)
			_ = W1770
			tmp9140 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1770)

			if True == tmp9140 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9126 := MakeNative(func(__e *ControlFlow) {
					W1771 := __e.Get(1)
					_ = W1771
					tmp9127 := MakeNative(func(__e *ControlFlow) {
						W1772 := __e.Get(1)
						_ = W1772
						tmp9128 := MakeNative(func(__e *ControlFlow) {
							W1773 := __e.Get(1)
							_ = W1773
							tmp9135 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1773)

							if True == tmp9135 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp9129 := MakeNative(func(__e *ControlFlow) {
									W1774 := __e.Get(1)
									_ = W1774
									tmp9130 := MakeNative(func(__e *ControlFlow) {
										W1775 := __e.Get(1)
										_ = W1775
										tmp9131 := PrimCons(W1771, W1774)

										__e.TailApply(PrimFunc(symshen_4comb), W1775, tmp9131)
										return

									}, 1)

									tmp9132 := Call(__e, PrimFunc(symshen_4in_1_6), W1773)

									__e.TailApply(tmp9130, tmp9132)
									return

								}, 1)

								tmp9133 := Call(__e, PrimFunc(symshen_4_5_1out), W1773)

								__e.TailApply(tmp9129, tmp9133)
								return

							}

						}, 1)

						tmp9136 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1772)

						__e.TailApply(tmp9128, tmp9136)
						return

					}, 1)

					tmp9137 := Call(__e, PrimFunc(symshen_4in_1_6), W1770)

					__e.TailApply(tmp9127, tmp9137)
					return

				}, 1)

				tmp9138 := Call(__e, PrimFunc(symshen_4_5_1out), W1770)

				__e.TailApply(tmp9126, tmp9138)
				return

			}

		}, 1)

		tmp9141 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1768)

		tmp9142 := Call(__e, tmp9125, tmp9141)

		__e.TailApply(tmp9112, tmp9142)
		return

	}, 1)

	tmp9143 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp9111)

	_ = tmp9143

	tmp9144 := MakeNative(func(__e *ControlFlow) {
		V1779 := __e.Get(1)
		_ = V1779
		tmp9145 := MakeNative(func(__e *ControlFlow) {
			W1780 := __e.Get(1)
			_ = W1780
			tmp9185 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1780)

			if True == tmp9185 {
				tmp9146 := MakeNative(func(__e *ControlFlow) {
					W1784 := __e.Get(1)
					_ = W1784
					tmp9173 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1784)

					if True == tmp9173 {
						tmp9147 := MakeNative(func(__e *ControlFlow) {
							W1787 := __e.Get(1)
							_ = W1787
							tmp9149 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1787)

							if True == tmp9149 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(W1787)
								return
							}

						}, 1)

						tmp9171 := Call(__e, PrimFunc(symshen_4ccons_2), V1779)

						var ifres9150 Obj

						if True == tmp9171 {
							tmp9151 := MakeNative(func(__e *ControlFlow) {
								W1788 := __e.Get(1)
								_ = W1788
								tmp9152 := MakeNative(func(__e *ControlFlow) {
									W1789 := __e.Get(1)
									_ = W1789
									tmp9153 := MakeNative(func(__e *ControlFlow) {
										W1790 := __e.Get(1)
										_ = W1790
										tmp9165 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1790)

										if True == tmp9165 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp9154 := MakeNative(func(__e *ControlFlow) {
												W1791 := __e.Get(1)
												_ = W1791
												tmp9155 := MakeNative(func(__e *ControlFlow) {
													W1792 := __e.Get(1)
													_ = W1792
													tmp9156 := MakeNative(func(__e *ControlFlow) {
														W1793 := __e.Get(1)
														_ = W1793
														tmp9160 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1793)

														if True == tmp9160 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp9157 := MakeNative(func(__e *ControlFlow) {
																W1794 := __e.Get(1)
																_ = W1794
																__e.TailApply(PrimFunc(symshen_4comb), W1789, W1791)
																return
															}, 1)

															tmp9158 := Call(__e, PrimFunc(symshen_4in_1_6), W1793)

															__e.TailApply(tmp9157, tmp9158)
															return

														}

													}, 1)

													tmp9161 := Call(__e, PrimFunc(sym_5end_6), W1792)

													__e.TailApply(tmp9156, tmp9161)
													return

												}, 1)

												tmp9162 := Call(__e, PrimFunc(symshen_4in_1_6), W1790)

												__e.TailApply(tmp9155, tmp9162)
												return

											}, 1)

											tmp9163 := Call(__e, PrimFunc(symshen_4_5_1out), W1790)

											__e.TailApply(tmp9154, tmp9163)
											return

										}

									}, 1)

									tmp9166 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1788)

									__e.TailApply(tmp9153, tmp9166)
									return

								}, 1)

								tmp9167 := Call(__e, PrimFunc(symtail), V1779)

								__e.TailApply(tmp9152, tmp9167)
								return

							}, 1)

							tmp9168 := Call(__e, PrimFunc(symhead), V1779)

							tmp9169 := Call(__e, tmp9151, tmp9168)

							ifres9150 = tmp9169

						} else {
							tmp9170 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres9150 = tmp9170

						}

						__e.TailApply(tmp9147, ifres9150)
						return

					} else {
						__e.Return(W1784)
						return
					}

				}, 1)

				tmp9183 := PrimIsPair(V1779)

				var ifres9174 Obj

				if True == tmp9183 {
					tmp9175 := MakeNative(func(__e *ControlFlow) {
						W1785 := __e.Get(1)
						_ = W1785
						tmp9176 := MakeNative(func(__e *ControlFlow) {
							W1786 := __e.Get(1)
							_ = W1786
							tmp9178 := Call(__e, PrimFunc(symatom_2), W1785)

							if True == tmp9178 {
								__e.TailApply(PrimFunc(symshen_4comb), W1786, W1785)
								return
							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp9179 := Call(__e, PrimFunc(symtail), V1779)

						__e.TailApply(tmp9176, tmp9179)
						return

					}, 1)

					tmp9180 := Call(__e, PrimFunc(symhead), V1779)

					tmp9181 := Call(__e, tmp9175, tmp9180)

					ifres9174 = tmp9181

				} else {
					tmp9182 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres9174 = tmp9182

				}

				__e.TailApply(tmp9146, ifres9174)
				return

			} else {
				__e.Return(W1780)
				return
			}

		}, 1)

		tmp9186 := MakeNative(func(__e *ControlFlow) {
			W1781 := __e.Get(1)
			_ = W1781
			tmp9192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1781)

			if True == tmp9192 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp9187 := MakeNative(func(__e *ControlFlow) {
					W1782 := __e.Get(1)
					_ = W1782
					tmp9188 := MakeNative(func(__e *ControlFlow) {
						W1783 := __e.Get(1)
						_ = W1783
						__e.TailApply(PrimFunc(symshen_4comb), W1783, W1782)
						return
					}, 1)

					tmp9189 := Call(__e, PrimFunc(symshen_4in_1_6), W1781)

					__e.TailApply(tmp9188, tmp9189)
					return

				}, 1)

				tmp9190 := Call(__e, PrimFunc(symshen_4_5_1out), W1781)

				__e.TailApply(tmp9187, tmp9190)
				return

			}

		}, 1)

		tmp9193 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1779)

		tmp9194 := Call(__e, tmp9186, tmp9193)

		__e.TailApply(tmp9145, tmp9194)
		return

	}, 1)

	tmp9195 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp9144)

	_ = tmp9195

	tmp9196 := MakeNative(func(__e *ControlFlow) {
		V1795 := __e.Get(1)
		_ = V1795
		tmp9197 := MakeNative(func(__e *ControlFlow) {
			W1796 := __e.Get(1)
			_ = W1796
			tmp9199 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1796)

			if True == tmp9199 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W1796)
				return
			}

		}, 1)

		tmp9210 := PrimIsPair(V1795)

		var ifres9200 Obj

		if True == tmp9210 {
			tmp9201 := MakeNative(func(__e *ControlFlow) {
				W1797 := __e.Get(1)
				_ = W1797
				tmp9202 := MakeNative(func(__e *ControlFlow) {
					W1798 := __e.Get(1)
					_ = W1798
					tmp9205 := PrimEqual(W1797, sym__)

					if True == tmp9205 {
						tmp9203 := Call(__e, PrimFunc(symgensym), symY)

						__e.TailApply(PrimFunc(symshen_4comb), W1798, tmp9203)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp9206 := Call(__e, PrimFunc(symtail), V1795)

				__e.TailApply(tmp9202, tmp9206)
				return

			}, 1)

			tmp9207 := Call(__e, PrimFunc(symhead), V1795)

			tmp9208 := Call(__e, tmp9201, tmp9207)

			ifres9200 = tmp9208

		} else {
			tmp9209 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres9200 = tmp9209

		}

		__e.TailApply(tmp9197, ifres9200)
		return

	}, 1)

	tmp9211 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp9196)

	_ = tmp9211

	tmp9212 := MakeNative(func(__e *ControlFlow) {
		V1799 := __e.Get(1)
		_ = V1799
		tmp9213 := MakeNative(func(__e *ControlFlow) {
			W1800 := __e.Get(1)
			_ = W1800
			tmp9215 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1800)

			if True == tmp9215 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W1800)
				return
			}

		}, 1)

		tmp9225 := PrimIsPair(V1799)

		var ifres9216 Obj

		if True == tmp9225 {
			tmp9217 := MakeNative(func(__e *ControlFlow) {
				W1801 := __e.Get(1)
				_ = W1801
				tmp9218 := MakeNative(func(__e *ControlFlow) {
					W1802 := __e.Get(1)
					_ = W1802
					tmp9220 := Call(__e, PrimFunc(symshen_4semicolon_2), W1801)

					if True == tmp9220 {
						__e.TailApply(PrimFunc(symshen_4comb), W1802, W1801)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp9221 := Call(__e, PrimFunc(symtail), V1799)

				__e.TailApply(tmp9218, tmp9221)
				return

			}, 1)

			tmp9222 := Call(__e, PrimFunc(symhead), V1799)

			tmp9223 := Call(__e, tmp9217, tmp9222)

			ifres9216 = tmp9223

		} else {
			tmp9224 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres9216 = tmp9224

		}

		__e.TailApply(tmp9213, ifres9216)
		return

	}, 1)

	tmp9226 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp9212)

	_ = tmp9226

	tmp9227 := MakeNative(func(__e *ControlFlow) {
		V1803 := __e.Get(1)
		_ = V1803
		V1804 := __e.Get(2)
		_ = V1804
		tmp9228 := MakeNative(func(__e *ControlFlow) {
			W1805 := __e.Get(1)
			_ = W1805
			tmp9229 := MakeNative(func(__e *ControlFlow) {
				W1806 := __e.Get(1)
				_ = W1806
				tmp9230 := MakeNative(func(__e *ControlFlow) {
					W1807 := __e.Get(1)
					_ = W1807
					tmp9231 := MakeNative(func(__e *ControlFlow) {
						W1808 := __e.Get(1)
						_ = W1808
						tmp9232 := MakeNative(func(__e *ControlFlow) {
							W1809 := __e.Get(1)
							_ = W1809
							tmp9233 := MakeNative(func(__e *ControlFlow) {
								W1810 := __e.Get(1)
								_ = W1810
								tmp9234 := MakeNative(func(__e *ControlFlow) {
									W1811 := __e.Get(1)
									_ = W1811
									tmp9235 := MakeNative(func(__e *ControlFlow) {
										W1812 := __e.Get(1)
										_ = W1812
										tmp9236 := MakeNative(func(__e *ControlFlow) {
											W1813 := __e.Get(1)
											_ = W1813
											__e.Return(W1813)
											return
										}, 1)

										tmp9237 := PrimCons(sym_1_6, Nil)

										tmp9238 := PrimCons(W1808, tmp9237)

										tmp9239 := PrimCons(W1807, tmp9238)

										tmp9240 := PrimCons(W1806, tmp9239)

										tmp9241 := PrimCons(W1805, tmp9240)

										tmp9242 := PrimCons(W1812, Nil)

										tmp9243 := Call(__e, PrimFunc(symappend), tmp9241, tmp9242)

										tmp9244 := Call(__e, PrimFunc(symappend), W1809, tmp9243)

										tmp9245 := PrimCons(V1803, tmp9244)

										tmp9246 := PrimCons(symdefine, tmp9245)

										__e.TailApply(tmp9236, tmp9246)
										return

									}, 1)

									var ifres9247 Obj

									if True == W1810 {
										tmp9248 := PrimCons(MakeNumber(1), Nil)

										tmp9249 := PrimCons(W1807, tmp9248)

										tmp9250 := PrimCons(sym_7, tmp9249)

										tmp9251 := PrimCons(W1811, Nil)

										tmp9252 := PrimCons(tmp9250, tmp9251)

										tmp9253 := PrimCons(W1807, tmp9252)

										tmp9254 := PrimCons(symlet, tmp9253)

										ifres9247 = tmp9254

									} else {
										ifres9247 = W1811

									}

									__e.TailApply(tmp9235, ifres9247)
									return

								}, 1)

								tmp9255 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1804, W1809, W1805, W1806, W1807, W1808, W1810)

								__e.TailApply(tmp9234, tmp9255)
								return

							}, 1)

							tmp9256 := Call(__e, PrimFunc(symshen_4hascut_2), V1804)

							__e.TailApply(tmp9233, tmp9256)
							return

						}, 1)

						tmp9257 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1804)

						__e.TailApply(tmp9232, tmp9257)
						return

					}, 1)

					tmp9258 := Call(__e, PrimFunc(symgensym), symC)

					__e.TailApply(tmp9231, tmp9258)
					return

				}, 1)

				tmp9259 := Call(__e, PrimFunc(symgensym), symK)

				__e.TailApply(tmp9230, tmp9259)
				return

			}, 1)

			tmp9260 := Call(__e, PrimFunc(symgensym), symL)

			__e.TailApply(tmp9229, tmp9260)
			return

		}, 1)

		tmp9261 := Call(__e, PrimFunc(symgensym), symB)

		__e.TailApply(tmp9228, tmp9261)
		return

	}, 2)

	tmp9262 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp9227)

	_ = tmp9262

	tmp9263 := MakeNative(func(__e *ControlFlow) {
		V1816 := __e.Get(1)
		_ = V1816
		tmp9273 := PrimEqual(sym_b, V1816)

		if True == tmp9273 {
			__e.Return(True)
			return
		} else {
			tmp9271 := PrimIsPair(V1816)

			if True == tmp9271 {
				tmp9268 := PrimHead(V1816)

				tmp9269 := Call(__e, PrimFunc(symshen_4hascut_2), tmp9268)

				if True == tmp9269 {
					__e.Return(True)
					return
				} else {
					tmp9265 := PrimTail(V1816)

					tmp9266 := Call(__e, PrimFunc(symshen_4hascut_2), tmp9265)

					if True == tmp9266 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp9274 := Call(__e, ns2_1set, symshen_4hascut_2, tmp9263)

	_ = tmp9274

	tmp9275 := MakeNative(func(__e *ControlFlow) {
		V1821 := __e.Get(1)
		_ = V1821
		tmp9284 := PrimIsPair(V1821)

		var ifres9280 Obj

		if True == tmp9284 {
			tmp9282 := PrimHead(V1821)

			tmp9283 := PrimIsPair(tmp9282)

			var ifres9281 Obj

			if True == tmp9283 {
				ifres9281 = True

			} else {
				ifres9281 = False

			}

			ifres9280 = ifres9281

		} else {
			ifres9280 = False

		}

		if True == ifres9280 {
			tmp9276 := PrimHead(V1821)

			tmp9277 := PrimHead(tmp9276)

			tmp9278 := Call(__e, PrimFunc(symlength), tmp9277)

			__e.TailApply(PrimFunc(symshen_4parameters), tmp9278)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1parameters)
			return
		}

	}, 1)

	tmp9285 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp9275)

	_ = tmp9285

	tmp9286 := MakeNative(func(__e *ControlFlow) {
		V1842 := __e.Get(1)
		_ = V1842
		V1843 := __e.Get(2)
		_ = V1843
		V1844 := __e.Get(3)
		_ = V1844
		V1845 := __e.Get(4)
		_ = V1845
		V1846 := __e.Get(5)
		_ = V1846
		V1847 := __e.Get(6)
		_ = V1847
		V1848 := __e.Get(7)
		_ = V1848
		tmp9379 := PrimEqual(Nil, V1842)

		var ifres9376 Obj

		if True == tmp9379 {
			tmp9378 := PrimEqual(True, V1848)

			var ifres9377 Obj

			if True == tmp9378 {
				ifres9377 = True

			} else {
				ifres9377 = False

			}

			ifres9376 = ifres9377

		} else {
			ifres9376 = False

		}

		if True == ifres9376 {
			tmp9287 := PrimCons(V1846, Nil)

			tmp9288 := PrimCons(V1845, tmp9287)

			__e.Return(PrimCons(symshen_4unlock, tmp9288))
			return

		} else {
			tmp9374 := PrimIsPair(V1842)

			var ifres9352 Obj

			if True == tmp9374 {
				tmp9372 := PrimHead(V1842)

				tmp9373 := PrimIsPair(tmp9372)

				var ifres9354 Obj

				if True == tmp9373 {
					tmp9369 := PrimHead(V1842)

					tmp9370 := PrimTail(tmp9369)

					tmp9371 := PrimIsPair(tmp9370)

					var ifres9356 Obj

					if True == tmp9371 {
						tmp9365 := PrimHead(V1842)

						tmp9366 := PrimTail(tmp9365)

						tmp9367 := PrimTail(tmp9366)

						tmp9368 := PrimEqual(Nil, tmp9367)

						var ifres9358 Obj

						if True == tmp9368 {
							tmp9363 := PrimTail(V1842)

							tmp9364 := PrimEqual(Nil, tmp9363)

							var ifres9360 Obj

							if True == tmp9364 {
								tmp9362 := PrimEqual(False, V1848)

								var ifres9361 Obj

								if True == tmp9362 {
									ifres9361 = True

								} else {
									ifres9361 = False

								}

								ifres9360 = ifres9361

							} else {
								ifres9360 = False

							}

							var ifres9359 Obj

							if True == ifres9360 {
								ifres9359 = True

							} else {
								ifres9359 = False

							}

							ifres9358 = ifres9359

						} else {
							ifres9358 = False

						}

						var ifres9357 Obj

						if True == ifres9358 {
							ifres9357 = True

						} else {
							ifres9357 = False

						}

						ifres9356 = ifres9357

					} else {
						ifres9356 = False

					}

					var ifres9355 Obj

					if True == ifres9356 {
						ifres9355 = True

					} else {
						ifres9355 = False

					}

					ifres9354 = ifres9355

				} else {
					ifres9354 = False

				}

				var ifres9353 Obj

				if True == ifres9354 {
					ifres9353 = True

				} else {
					ifres9353 = False

				}

				ifres9352 = ifres9353

			} else {
				ifres9352 = False

			}

			if True == ifres9352 {
				tmp9289 := MakeNative(func(__e *ControlFlow) {
					W1849 := __e.Get(1)
					_ = W1849
					tmp9290 := PrimCons(V1845, Nil)

					tmp9291 := PrimCons(symshen_4unlocked_2, tmp9290)

					tmp9292 := PrimHead(V1842)

					tmp9293 := PrimHead(tmp9292)

					tmp9294 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9293, V1843, V1844, W1849)

					tmp9295 := PrimCons(False, Nil)

					tmp9296 := PrimCons(tmp9294, tmp9295)

					tmp9297 := PrimCons(tmp9291, tmp9296)

					__e.Return(PrimCons(symif, tmp9297))
					return

				}, 1)

				tmp9298 := PrimHead(V1842)

				tmp9299 := PrimHead(tmp9298)

				tmp9300 := PrimHead(V1842)

				tmp9301 := PrimTail(tmp9300)

				tmp9302 := PrimHead(tmp9301)

				tmp9303 := Call(__e, PrimFunc(symshen_4continue), tmp9299, tmp9302, V1844, V1845, V1846, V1847)

				__e.TailApply(tmp9289, tmp9303)
				return

			} else {
				tmp9350 := PrimIsPair(V1842)

				var ifres9335 Obj

				if True == tmp9350 {
					tmp9348 := PrimHead(V1842)

					tmp9349 := PrimIsPair(tmp9348)

					var ifres9337 Obj

					if True == tmp9349 {
						tmp9345 := PrimHead(V1842)

						tmp9346 := PrimTail(tmp9345)

						tmp9347 := PrimIsPair(tmp9346)

						var ifres9339 Obj

						if True == tmp9347 {
							tmp9341 := PrimHead(V1842)

							tmp9342 := PrimTail(tmp9341)

							tmp9343 := PrimTail(tmp9342)

							tmp9344 := PrimEqual(Nil, tmp9343)

							var ifres9340 Obj

							if True == tmp9344 {
								ifres9340 = True

							} else {
								ifres9340 = False

							}

							ifres9339 = ifres9340

						} else {
							ifres9339 = False

						}

						var ifres9338 Obj

						if True == ifres9339 {
							ifres9338 = True

						} else {
							ifres9338 = False

						}

						ifres9337 = ifres9338

					} else {
						ifres9337 = False

					}

					var ifres9336 Obj

					if True == ifres9337 {
						ifres9336 = True

					} else {
						ifres9336 = False

					}

					ifres9335 = ifres9336

				} else {
					ifres9335 = False

				}

				if True == ifres9335 {
					tmp9304 := MakeNative(func(__e *ControlFlow) {
						W1850 := __e.Get(1)
						_ = W1850
						tmp9305 := MakeNative(func(__e *ControlFlow) {
							W1851 := __e.Get(1)
							_ = W1851
							tmp9306 := PrimCons(V1845, Nil)

							tmp9307 := PrimCons(symshen_4unlocked_2, tmp9306)

							tmp9308 := PrimHead(V1842)

							tmp9309 := PrimHead(tmp9308)

							tmp9310 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9309, V1843, V1844, W1851)

							tmp9311 := PrimCons(False, Nil)

							tmp9312 := PrimCons(tmp9310, tmp9311)

							tmp9313 := PrimCons(tmp9307, tmp9312)

							tmp9314 := PrimCons(symif, tmp9313)

							tmp9315 := PrimCons(False, Nil)

							tmp9316 := PrimCons(W1850, tmp9315)

							tmp9317 := PrimCons(sym_a, tmp9316)

							tmp9318 := PrimTail(V1842)

							tmp9319 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp9318, V1843, V1844, V1845, V1846, V1847, V1848)

							tmp9320 := PrimCons(W1850, Nil)

							tmp9321 := PrimCons(tmp9319, tmp9320)

							tmp9322 := PrimCons(tmp9317, tmp9321)

							tmp9323 := PrimCons(symif, tmp9322)

							tmp9324 := PrimCons(tmp9323, Nil)

							tmp9325 := PrimCons(tmp9314, tmp9324)

							tmp9326 := PrimCons(W1850, tmp9325)

							__e.Return(PrimCons(symlet, tmp9326))
							return

						}, 1)

						tmp9327 := PrimHead(V1842)

						tmp9328 := PrimHead(tmp9327)

						tmp9329 := PrimHead(V1842)

						tmp9330 := PrimTail(tmp9329)

						tmp9331 := PrimHead(tmp9330)

						tmp9332 := Call(__e, PrimFunc(symshen_4continue), tmp9328, tmp9331, V1844, V1845, V1846, V1847)

						__e.TailApply(tmp9305, tmp9332)
						return

					}, 1)

					tmp9333 := Call(__e, PrimFunc(symgensym), symC)

					__e.TailApply(tmp9304, tmp9333)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
					return
				}

			}

		}

	}, 7)

	tmp9380 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp9286)

	_ = tmp9380

	tmp9381 := MakeNative(func(__e *ControlFlow) {
		V1852 := __e.Get(1)
		_ = V1852
		V1853 := __e.Get(2)
		_ = V1853
		tmp9386 := Call(__e, PrimFunc(symshen_4locked_2), V1852)

		var ifres9383 Obj

		if True == tmp9386 {
			tmp9385 := Call(__e, PrimFunc(symshen_4fits_2), V1853, V1852)

			var ifres9384 Obj

			if True == tmp9385 {
				ifres9384 = True

			} else {
				ifres9384 = False

			}

			ifres9383 = ifres9384

		} else {
			ifres9383 = False

		}

		if True == ifres9383 {
			__e.TailApply(PrimFunc(symshen_4openlock), V1852)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp9387 := Call(__e, ns2_1set, symshen_4unlock, tmp9381)

	_ = tmp9387

	tmp9388 := MakeNative(func(__e *ControlFlow) {
		V1854 := __e.Get(1)
		_ = V1854
		tmp9389 := Call(__e, PrimFunc(symshen_4unlocked_2), V1854)

		__e.Return(PrimNot(tmp9389))
		return

	}, 1)

	tmp9390 := Call(__e, ns2_1set, symshen_4locked_2, tmp9388)

	_ = tmp9390

	tmp9391 := MakeNative(func(__e *ControlFlow) {
		V1855 := __e.Get(1)
		_ = V1855
		__e.Return(PrimVectorGet(V1855, MakeNumber(1)))
		return
	}, 1)

	tmp9392 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp9391)

	_ = tmp9392

	tmp9393 := MakeNative(func(__e *ControlFlow) {
		V1856 := __e.Get(1)
		_ = V1856
		tmp9394 := PrimVectorSet(V1856, MakeNumber(1), True)

		_ = tmp9394

		__e.Return(False)
		return

	}, 1)

	tmp9395 := Call(__e, ns2_1set, symshen_4openlock, tmp9393)

	_ = tmp9395

	tmp9396 := MakeNative(func(__e *ControlFlow) {
		V1857 := __e.Get(1)
		_ = V1857
		V1858 := __e.Get(2)
		_ = V1858
		tmp9397 := PrimVectorGet(V1858, MakeNumber(2))

		__e.Return(PrimEqual(V1857, tmp9397))
		return

	}, 2)

	tmp9398 := Call(__e, ns2_1set, symshen_4fits_2, tmp9396)

	_ = tmp9398

	tmp9399 := MakeNative(func(__e *ControlFlow) {
		V1861 := __e.Get(1)
		_ = V1861
		V1862 := __e.Get(2)
		_ = V1862
		V1863 := __e.Get(3)
		_ = V1863
		V1864 := __e.Get(4)
		_ = V1864
		tmp9400 := MakeNative(func(__e *ControlFlow) {
			W1865 := __e.Get(1)
			_ = W1865
			tmp9405 := PrimEqual(W1865, False)

			var ifres9402 Obj

			if True == tmp9405 {
				tmp9404 := Call(__e, PrimFunc(symshen_4unlocked_2), V1862)

				var ifres9403 Obj

				if True == tmp9404 {
					ifres9403 = True

				} else {
					ifres9403 = False

				}

				ifres9402 = ifres9403

			} else {
				ifres9402 = False

			}

			if True == ifres9402 {
				__e.TailApply(PrimFunc(symshen_4lock), V1863, V1862)
				return
			} else {
				__e.Return(W1865)
				return
			}

		}, 1)

		tmp9406 := Call(__e, PrimFunc(symthaw), V1864)

		__e.TailApply(tmp9400, tmp9406)
		return

	}, 4)

	tmp9407 := Call(__e, ns2_1set, symshen_4cut, tmp9399)

	_ = tmp9407

	tmp9408 := MakeNative(func(__e *ControlFlow) {
		V1866 := __e.Get(1)
		_ = V1866
		V1867 := __e.Get(2)
		_ = V1867
		tmp9409 := MakeNative(func(__e *ControlFlow) {
			W1868 := __e.Get(1)
			_ = W1868
			tmp9410 := MakeNative(func(__e *ControlFlow) {
				W1869 := __e.Get(1)
				_ = W1869
				__e.Return(False)
				return
			}, 1)

			tmp9411 := PrimVectorSet(V1867, MakeNumber(2), V1866)

			__e.TailApply(tmp9410, tmp9411)
			return

		}, 1)

		tmp9412 := PrimVectorSet(V1867, MakeNumber(1), False)

		__e.TailApply(tmp9409, tmp9412)
		return

	}, 2)

	tmp9413 := Call(__e, ns2_1set, symshen_4lock, tmp9408)

	_ = tmp9413

	tmp9414 := MakeNative(func(__e *ControlFlow) {
		V1870 := __e.Get(1)
		_ = V1870
		V1871 := __e.Get(2)
		_ = V1871
		V1872 := __e.Get(3)
		_ = V1872
		V1873 := __e.Get(4)
		_ = V1873
		V1874 := __e.Get(5)
		_ = V1874
		V1875 := __e.Get(6)
		_ = V1875
		tmp9415 := MakeNative(func(__e *ControlFlow) {
			W1876 := __e.Get(1)
			_ = W1876
			tmp9416 := MakeNative(func(__e *ControlFlow) {
				W1877 := __e.Get(1)
				_ = W1877
				tmp9417 := MakeNative(func(__e *ControlFlow) {
					W1878 := __e.Get(1)
					_ = W1878
					tmp9418 := MakeNative(func(__e *ControlFlow) {
						W1879 := __e.Get(1)
						_ = W1879
						__e.TailApply(PrimFunc(symshen_4stpart), W1878, W1879, V1872)
						return
					}, 1)

					tmp9419 := PrimCons(symshen_4incinfs, Nil)

					tmp9420 := Call(__e, PrimFunc(symshen_4compile_1body), V1871, V1872, V1873, V1874, V1875)

					tmp9421 := PrimCons(tmp9420, Nil)

					tmp9422 := PrimCons(tmp9419, tmp9421)

					tmp9423 := PrimCons(symdo, tmp9422)

					__e.TailApply(tmp9418, tmp9423)
					return

				}, 1)

				tmp9424 := Call(__e, PrimFunc(symdifference), W1877, W1876)

				__e.TailApply(tmp9417, tmp9424)
				return

			}, 1)

			tmp9425 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1871)

			__e.TailApply(tmp9416, tmp9425)
			return

		}, 1)

		tmp9426 := Call(__e, PrimFunc(symshen_4extract_1vars), V1870)

		__e.TailApply(tmp9415, tmp9426)
		return

	}, 6)

	tmp9427 := Call(__e, ns2_1set, symshen_4continue, tmp9414)

	_ = tmp9427

	tmp9428 := MakeNative(func(__e *ControlFlow) {
		V1882 := __e.Get(1)
		_ = V1882
		tmp9463 := PrimIsPair(V1882)

		var ifres9444 Obj

		if True == tmp9463 {
			tmp9461 := PrimHead(V1882)

			tmp9462 := PrimEqual(symlambda, tmp9461)

			var ifres9446 Obj

			if True == tmp9462 {
				tmp9459 := PrimTail(V1882)

				tmp9460 := PrimIsPair(tmp9459)

				var ifres9448 Obj

				if True == tmp9460 {
					tmp9456 := PrimTail(V1882)

					tmp9457 := PrimTail(tmp9456)

					tmp9458 := PrimIsPair(tmp9457)

					var ifres9450 Obj

					if True == tmp9458 {
						tmp9452 := PrimTail(V1882)

						tmp9453 := PrimTail(tmp9452)

						tmp9454 := PrimTail(tmp9453)

						tmp9455 := PrimEqual(Nil, tmp9454)

						var ifres9451 Obj

						if True == tmp9455 {
							ifres9451 = True

						} else {
							ifres9451 = False

						}

						ifres9450 = ifres9451

					} else {
						ifres9450 = False

					}

					var ifres9449 Obj

					if True == ifres9450 {
						ifres9449 = True

					} else {
						ifres9449 = False

					}

					ifres9448 = ifres9449

				} else {
					ifres9448 = False

				}

				var ifres9447 Obj

				if True == ifres9448 {
					ifres9447 = True

				} else {
					ifres9447 = False

				}

				ifres9446 = ifres9447

			} else {
				ifres9446 = False

			}

			var ifres9445 Obj

			if True == ifres9446 {
				ifres9445 = True

			} else {
				ifres9445 = False

			}

			ifres9444 = ifres9445

		} else {
			ifres9444 = False

		}

		if True == ifres9444 {
			tmp9429 := PrimTail(V1882)

			tmp9430 := PrimHead(tmp9429)

			tmp9431 := PrimTail(V1882)

			tmp9432 := PrimTail(tmp9431)

			tmp9433 := PrimHead(tmp9432)

			tmp9434 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp9433)

			__e.TailApply(PrimFunc(symremove), tmp9430, tmp9434)
			return

		} else {
			tmp9442 := PrimIsPair(V1882)

			if True == tmp9442 {
				tmp9435 := PrimHead(V1882)

				tmp9436 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp9435)

				tmp9437 := PrimTail(V1882)

				tmp9438 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp9437)

				__e.TailApply(PrimFunc(symunion), tmp9436, tmp9438)
				return

			} else {
				tmp9440 := PrimIsVariable(V1882)

				if True == tmp9440 {
					__e.Return(PrimCons(V1882, Nil))
					return
				} else {
					__e.Return(Nil)
					return
				}

			}

		}

	}, 1)

	tmp9464 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp9428)

	_ = tmp9464

	tmp9465 := MakeNative(func(__e *ControlFlow) {
		V1899 := __e.Get(1)
		_ = V1899
		V1900 := __e.Get(2)
		_ = V1900
		V1901 := __e.Get(3)
		_ = V1901
		V1902 := __e.Get(4)
		_ = V1902
		V1903 := __e.Get(5)
		_ = V1903
		tmp9500 := PrimEqual(Nil, V1899)

		if True == tmp9500 {
			tmp9466 := PrimCons(V1903, Nil)

			__e.Return(PrimCons(symthaw, tmp9466))
			return

		} else {
			tmp9498 := PrimIsPair(V1899)

			var ifres9494 Obj

			if True == tmp9498 {
				tmp9496 := PrimHead(V1899)

				tmp9497 := PrimEqual(sym_b, tmp9496)

				var ifres9495 Obj

				if True == tmp9497 {
					ifres9495 = True

				} else {
					ifres9495 = False

				}

				ifres9494 = ifres9495

			} else {
				ifres9494 = False

			}

			if True == ifres9494 {
				tmp9467 := PrimCons(symshen_4cut, Nil)

				tmp9468 := PrimTail(V1899)

				tmp9469 := PrimCons(tmp9467, tmp9468)

				__e.TailApply(PrimFunc(symshen_4compile_1body), tmp9469, V1900, V1901, V1902, V1903)
				return

			} else {
				tmp9492 := PrimIsPair(V1899)

				var ifres9488 Obj

				if True == tmp9492 {
					tmp9490 := PrimTail(V1899)

					tmp9491 := PrimEqual(Nil, tmp9490)

					var ifres9489 Obj

					if True == tmp9491 {
						ifres9489 = True

					} else {
						ifres9489 = False

					}

					ifres9488 = ifres9489

				} else {
					ifres9488 = False

				}

				if True == ifres9488 {
					tmp9470 := PrimHead(V1899)

					tmp9471 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9470, V1900)

					tmp9472 := PrimCons(V1903, Nil)

					tmp9473 := PrimCons(V1902, tmp9472)

					tmp9474 := PrimCons(V1901, tmp9473)

					tmp9475 := PrimCons(V1900, tmp9474)

					__e.TailApply(PrimFunc(symappend), tmp9471, tmp9475)
					return

				} else {
					tmp9486 := PrimIsPair(V1899)

					if True == tmp9486 {
						tmp9476 := MakeNative(func(__e *ControlFlow) {
							W1904 := __e.Get(1)
							_ = W1904
							tmp9477 := PrimTail(V1899)

							tmp9478 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp9477, V1900, V1901, V1902, V1903)

							tmp9479 := PrimCons(tmp9478, Nil)

							tmp9480 := PrimCons(V1902, tmp9479)

							tmp9481 := PrimCons(V1901, tmp9480)

							tmp9482 := PrimCons(V1900, tmp9481)

							__e.TailApply(PrimFunc(symappend), W1904, tmp9482)
							return

						}, 1)

						tmp9483 := PrimHead(V1899)

						tmp9484 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9483, V1900)

						__e.TailApply(tmp9476, tmp9484)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
						return
					}

				}

			}

		}

	}, 5)

	tmp9501 := Call(__e, ns2_1set, symshen_4compile_1body, tmp9465)

	_ = tmp9501

	tmp9502 := MakeNative(func(__e *ControlFlow) {
		V1921 := __e.Get(1)
		_ = V1921
		V1922 := __e.Get(2)
		_ = V1922
		V1923 := __e.Get(3)
		_ = V1923
		V1924 := __e.Get(4)
		_ = V1924
		V1925 := __e.Get(5)
		_ = V1925
		tmp9526 := PrimEqual(Nil, V1921)

		if True == tmp9526 {
			__e.Return(V1925)
			return
		} else {
			tmp9524 := PrimIsPair(V1921)

			var ifres9520 Obj

			if True == tmp9524 {
				tmp9522 := PrimHead(V1921)

				tmp9523 := PrimEqual(sym_b, tmp9522)

				var ifres9521 Obj

				if True == tmp9523 {
					ifres9521 = True

				} else {
					ifres9521 = False

				}

				ifres9520 = ifres9521

			} else {
				ifres9520 = False

			}

			if True == ifres9520 {
				tmp9503 := PrimCons(symshen_4cut, Nil)

				tmp9504 := PrimTail(V1921)

				tmp9505 := PrimCons(tmp9503, tmp9504)

				__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp9505, V1922, V1923, V1924, V1925)
				return

			} else {
				tmp9518 := PrimIsPair(V1921)

				if True == tmp9518 {
					tmp9506 := MakeNative(func(__e *ControlFlow) {
						W1926 := __e.Get(1)
						_ = W1926
						tmp9507 := PrimTail(V1921)

						tmp9508 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp9507, V1922, V1923, V1924, V1925)

						tmp9509 := PrimCons(tmp9508, Nil)

						tmp9510 := PrimCons(V1924, tmp9509)

						tmp9511 := PrimCons(V1923, tmp9510)

						tmp9512 := PrimCons(V1922, tmp9511)

						tmp9513 := Call(__e, PrimFunc(symappend), W1926, tmp9512)

						tmp9514 := PrimCons(tmp9513, Nil)

						__e.Return(PrimCons(symfreeze, tmp9514))
						return

					}, 1)

					tmp9515 := PrimHead(V1921)

					tmp9516 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9515, V1922)

					__e.TailApply(tmp9506, tmp9516)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
					return
				}

			}

		}

	}, 5)

	tmp9527 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp9502)

	_ = tmp9527

	tmp9528 := MakeNative(func(__e *ControlFlow) {
		V1931 := __e.Get(1)
		_ = V1931
		V1932 := __e.Get(2)
		_ = V1932
		tmp9543 := PrimIsPair(V1931)

		var ifres9539 Obj

		if True == tmp9543 {
			tmp9541 := PrimHead(V1931)

			tmp9542 := PrimEqual(symfork, tmp9541)

			var ifres9540 Obj

			if True == tmp9542 {
				ifres9540 = True

			} else {
				ifres9540 = False

			}

			ifres9539 = ifres9540

		} else {
			ifres9539 = False

		}

		if True == ifres9539 {
			tmp9529 := PrimTail(V1931)

			tmp9530 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp9529, V1932)

			tmp9531 := PrimCons(tmp9530, Nil)

			__e.Return(PrimCons(symfork, tmp9531))
			return

		} else {
			tmp9537 := PrimIsPair(V1931)

			if True == tmp9537 {
				tmp9532 := PrimHead(V1931)

				tmp9533 := MakeNative(func(__e *ControlFlow) {
					Z1933 := __e.Get(1)
					_ = Z1933
					__e.TailApply(PrimFunc(symshen_4function_1calls), Z1933, V1932)
					return
				}, 1)

				tmp9534 := PrimTail(V1931)

				tmp9535 := Call(__e, PrimFunc(symmap), tmp9533, tmp9534)

				__e.Return(PrimCons(tmp9532, tmp9535))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
				return
			}

		}

	}, 2)

	tmp9544 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp9528)

	_ = tmp9544

	tmp9545 := MakeNative(func(__e *ControlFlow) {
		V1940 := __e.Get(1)
		_ = V1940
		V1941 := __e.Get(2)
		_ = V1941
		tmp9555 := PrimEqual(Nil, V1940)

		if True == tmp9555 {
			__e.Return(Nil)
			return
		} else {
			tmp9553 := PrimIsPair(V1940)

			if True == tmp9553 {
				tmp9546 := PrimHead(V1940)

				tmp9547 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9546, V1941)

				tmp9548 := PrimTail(V1940)

				tmp9549 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp9548, V1941)

				tmp9550 := PrimCons(tmp9549, Nil)

				tmp9551 := PrimCons(tmp9547, tmp9550)

				__e.Return(PrimCons(symcons, tmp9551))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
				return
			}

		}

	}, 2)

	tmp9556 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp9545)

	_ = tmp9556

	tmp9557 := MakeNative(func(__e *ControlFlow) {
		V1944 := __e.Get(1)
		_ = V1944
		V1945 := __e.Get(2)
		_ = V1945
		tmp9589 := PrimIsPair(V1944)

		var ifres9570 Obj

		if True == tmp9589 {
			tmp9587 := PrimHead(V1944)

			tmp9588 := PrimEqual(symcons, tmp9587)

			var ifres9572 Obj

			if True == tmp9588 {
				tmp9585 := PrimTail(V1944)

				tmp9586 := PrimIsPair(tmp9585)

				var ifres9574 Obj

				if True == tmp9586 {
					tmp9582 := PrimTail(V1944)

					tmp9583 := PrimTail(tmp9582)

					tmp9584 := PrimIsPair(tmp9583)

					var ifres9576 Obj

					if True == tmp9584 {
						tmp9578 := PrimTail(V1944)

						tmp9579 := PrimTail(tmp9578)

						tmp9580 := PrimTail(tmp9579)

						tmp9581 := PrimEqual(Nil, tmp9580)

						var ifres9577 Obj

						if True == tmp9581 {
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
			tmp9558 := PrimTail(V1944)

			tmp9559 := PrimHead(tmp9558)

			tmp9560 := Call(__e, PrimFunc(symshen_4function_1calls), tmp9559, V1945)

			tmp9561 := PrimTail(V1944)

			tmp9562 := PrimTail(tmp9561)

			tmp9563 := PrimHead(tmp9562)

			tmp9564 := Call(__e, PrimFunc(symshen_4function_1calls), tmp9563, V1945)

			tmp9565 := PrimCons(tmp9564, Nil)

			tmp9566 := PrimCons(tmp9560, tmp9565)

			__e.Return(PrimCons(symcons, tmp9566))
			return

		} else {
			tmp9568 := PrimIsPair(V1944)

			if True == tmp9568 {
				__e.TailApply(PrimFunc(symshen_4deref_1terms), V1944, V1945, Nil)
				return
			} else {
				__e.Return(V1944)
				return
			}

		}

	}, 2)

	tmp9590 := Call(__e, ns2_1set, symshen_4function_1calls, tmp9557)

	_ = tmp9590

	tmp9591 := MakeNative(func(__e *ControlFlow) {
		V1954 := __e.Get(1)
		_ = V1954
		V1955 := __e.Get(2)
		_ = V1955
		V1956 := __e.Get(3)
		_ = V1956
		tmp9685 := PrimIsPair(V1954)

		var ifres9672 Obj

		if True == tmp9685 {
			tmp9683 := PrimHead(V1954)

			tmp9684 := PrimEqual(MakeNumber(0), tmp9683)

			var ifres9674 Obj

			if True == tmp9684 {
				tmp9681 := PrimTail(V1954)

				tmp9682 := PrimIsPair(tmp9681)

				var ifres9676 Obj

				if True == tmp9682 {
					tmp9678 := PrimTail(V1954)

					tmp9679 := PrimTail(tmp9678)

					tmp9680 := PrimEqual(Nil, tmp9679)

					var ifres9677 Obj

					if True == tmp9680 {
						ifres9677 = True

					} else {
						ifres9677 = False

					}

					ifres9676 = ifres9677

				} else {
					ifres9676 = False

				}

				var ifres9675 Obj

				if True == ifres9676 {
					ifres9675 = True

				} else {
					ifres9675 = False

				}

				ifres9674 = ifres9675

			} else {
				ifres9674 = False

			}

			var ifres9673 Obj

			if True == ifres9674 {
				ifres9673 = True

			} else {
				ifres9673 = False

			}

			ifres9672 = ifres9673

		} else {
			ifres9672 = False

		}

		if True == ifres9672 {
			tmp9598 := PrimTail(V1954)

			tmp9599 := PrimHead(tmp9598)

			tmp9600 := PrimIsVariable(tmp9599)

			if True == tmp9600 {
				tmp9592 := PrimTail(V1954)

				__e.Return(PrimHead(tmp9592))
				return

			} else {
				tmp9593 := PrimTail(V1954)

				tmp9594 := PrimHead(tmp9593)

				tmp9595 := Call(__e, PrimFunc(symshen_4app), tmp9594, MakeString("\n"), symshen_4s)

				tmp9596 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp9595)

				__e.Return(PrimSimpleError(tmp9596))
				return

			}

		} else {
			tmp9670 := PrimIsPair(V1954)

			var ifres9657 Obj

			if True == tmp9670 {
				tmp9668 := PrimHead(V1954)

				tmp9669 := PrimEqual(MakeNumber(1), tmp9668)

				var ifres9659 Obj

				if True == tmp9669 {
					tmp9666 := PrimTail(V1954)

					tmp9667 := PrimIsPair(tmp9666)

					var ifres9661 Obj

					if True == tmp9667 {
						tmp9663 := PrimTail(V1954)

						tmp9664 := PrimTail(tmp9663)

						tmp9665 := PrimEqual(Nil, tmp9664)

						var ifres9662 Obj

						if True == tmp9665 {
							ifres9662 = True

						} else {
							ifres9662 = False

						}

						ifres9661 = ifres9662

					} else {
						ifres9661 = False

					}

					var ifres9660 Obj

					if True == ifres9661 {
						ifres9660 = True

					} else {
						ifres9660 = False

					}

					ifres9659 = ifres9660

				} else {
					ifres9659 = False

				}

				var ifres9658 Obj

				if True == ifres9659 {
					ifres9658 = True

				} else {
					ifres9658 = False

				}

				ifres9657 = ifres9658

			} else {
				ifres9657 = False

			}

			if True == ifres9657 {
				tmp9610 := PrimTail(V1954)

				tmp9611 := PrimHead(tmp9610)

				tmp9612 := PrimIsVariable(tmp9611)

				if True == tmp9612 {
					tmp9601 := PrimTail(V1954)

					tmp9602 := PrimHead(tmp9601)

					tmp9603 := PrimCons(V1955, Nil)

					tmp9604 := PrimCons(tmp9602, tmp9603)

					__e.Return(PrimCons(symshen_4lazyderef, tmp9604))
					return

				} else {
					tmp9605 := PrimTail(V1954)

					tmp9606 := PrimHead(tmp9605)

					tmp9607 := Call(__e, PrimFunc(symshen_4app), tmp9606, MakeString("\n"), symshen_4s)

					tmp9608 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp9607)

					__e.Return(PrimSimpleError(tmp9608))
					return

				}

			} else {
				tmp9654 := Call(__e, PrimFunc(symelement_2), V1954, V1956)

				tmp9655 := PrimNot(tmp9654)

				var ifres9651 Obj

				if True == tmp9655 {
					tmp9653 := PrimIsVariable(V1954)

					var ifres9652 Obj

					if True == tmp9653 {
						ifres9652 = True

					} else {
						ifres9652 = False

					}

					ifres9651 = ifres9652

				} else {
					ifres9651 = False

				}

				if True == ifres9651 {
					tmp9613 := PrimCons(V1955, Nil)

					tmp9614 := PrimCons(V1954, tmp9613)

					__e.Return(PrimCons(symshen_4deref, tmp9614))
					return

				} else {
					tmp9649 := PrimIsPair(V1954)

					var ifres9630 Obj

					if True == tmp9649 {
						tmp9647 := PrimHead(V1954)

						tmp9648 := PrimEqual(symlambda, tmp9647)

						var ifres9632 Obj

						if True == tmp9648 {
							tmp9645 := PrimTail(V1954)

							tmp9646 := PrimIsPair(tmp9645)

							var ifres9634 Obj

							if True == tmp9646 {
								tmp9642 := PrimTail(V1954)

								tmp9643 := PrimTail(tmp9642)

								tmp9644 := PrimIsPair(tmp9643)

								var ifres9636 Obj

								if True == tmp9644 {
									tmp9638 := PrimTail(V1954)

									tmp9639 := PrimTail(tmp9638)

									tmp9640 := PrimTail(tmp9639)

									tmp9641 := PrimEqual(Nil, tmp9640)

									var ifres9637 Obj

									if True == tmp9641 {
										ifres9637 = True

									} else {
										ifres9637 = False

									}

									ifres9636 = ifres9637

								} else {
									ifres9636 = False

								}

								var ifres9635 Obj

								if True == ifres9636 {
									ifres9635 = True

								} else {
									ifres9635 = False

								}

								ifres9634 = ifres9635

							} else {
								ifres9634 = False

							}

							var ifres9633 Obj

							if True == ifres9634 {
								ifres9633 = True

							} else {
								ifres9633 = False

							}

							ifres9632 = ifres9633

						} else {
							ifres9632 = False

						}

						var ifres9631 Obj

						if True == ifres9632 {
							ifres9631 = True

						} else {
							ifres9631 = False

						}

						ifres9630 = ifres9631

					} else {
						ifres9630 = False

					}

					if True == ifres9630 {
						tmp9615 := PrimTail(V1954)

						tmp9616 := PrimHead(tmp9615)

						tmp9617 := PrimTail(V1954)

						tmp9618 := PrimTail(tmp9617)

						tmp9619 := PrimHead(tmp9618)

						tmp9620 := PrimTail(V1954)

						tmp9621 := PrimHead(tmp9620)

						tmp9622 := PrimCons(tmp9621, V1956)

						tmp9623 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp9619, V1955, tmp9622)

						tmp9624 := PrimCons(tmp9623, Nil)

						tmp9625 := PrimCons(tmp9616, tmp9624)

						__e.Return(PrimCons(symlambda, tmp9625))
						return

					} else {
						tmp9628 := PrimIsPair(V1954)

						if True == tmp9628 {
							tmp9626 := MakeNative(func(__e *ControlFlow) {
								Z1957 := __e.Get(1)
								_ = Z1957
								__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1957, V1955, V1956)
								return
							}, 1)

							__e.TailApply(PrimFunc(symmap), tmp9626, V1954)
							return

						} else {
							__e.Return(V1954)
							return
						}

					}

				}

			}

		}

	}, 3)

	tmp9686 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp9591)

	_ = tmp9686

	tmp9687 := MakeNative(func(__e *ControlFlow) {
		V1975 := __e.Get(1)
		_ = V1975
		V1976 := __e.Get(2)
		_ = V1976
		V1977 := __e.Get(3)
		_ = V1977
		V1978 := __e.Get(4)
		_ = V1978
		V1979 := __e.Get(5)
		_ = V1979
		tmp9863 := PrimEqual(Nil, V1976)

		var ifres9860 Obj

		if True == tmp9863 {
			tmp9862 := PrimEqual(Nil, V1977)

			var ifres9861 Obj

			if True == tmp9862 {
				ifres9861 = True

			} else {
				ifres9861 = False

			}

			ifres9860 = ifres9861

		} else {
			ifres9860 = False

		}

		if True == ifres9860 {
			__e.Return(V1979)
			return
		} else {
			tmp9858 := PrimIsPair(V1976)

			var ifres9838 Obj

			if True == tmp9858 {
				tmp9856 := PrimHead(V1976)

				tmp9857 := PrimIsPair(tmp9856)

				var ifres9840 Obj

				if True == tmp9857 {
					tmp9853 := PrimHead(V1976)

					tmp9854 := PrimHead(tmp9853)

					tmp9855 := PrimEqual(symshen_4_7m, tmp9854)

					var ifres9842 Obj

					if True == tmp9855 {
						tmp9850 := PrimHead(V1976)

						tmp9851 := PrimTail(tmp9850)

						tmp9852 := PrimIsPair(tmp9851)

						var ifres9844 Obj

						if True == tmp9852 {
							tmp9846 := PrimHead(V1976)

							tmp9847 := PrimTail(tmp9846)

							tmp9848 := PrimTail(tmp9847)

							tmp9849 := PrimEqual(Nil, tmp9848)

							var ifres9845 Obj

							if True == tmp9849 {
								ifres9845 = True

							} else {
								ifres9845 = False

							}

							ifres9844 = ifres9845

						} else {
							ifres9844 = False

						}

						var ifres9843 Obj

						if True == ifres9844 {
							ifres9843 = True

						} else {
							ifres9843 = False

						}

						ifres9842 = ifres9843

					} else {
						ifres9842 = False

					}

					var ifres9841 Obj

					if True == ifres9842 {
						ifres9841 = True

					} else {
						ifres9841 = False

					}

					ifres9840 = ifres9841

				} else {
					ifres9840 = False

				}

				var ifres9839 Obj

				if True == ifres9840 {
					ifres9839 = True

				} else {
					ifres9839 = False

				}

				ifres9838 = ifres9839

			} else {
				ifres9838 = False

			}

			if True == ifres9838 {
				tmp9688 := PrimHead(V1976)

				tmp9689 := PrimTail(tmp9688)

				tmp9690 := PrimHead(tmp9689)

				tmp9691 := PrimTail(V1976)

				tmp9692 := PrimCons(V1975, tmp9691)

				tmp9693 := PrimCons(tmp9690, tmp9692)

				tmp9694 := PrimCons(symshen_4_7m, tmp9693)

				__e.TailApply(PrimFunc(symshen_4compile_1head), V1975, tmp9694, V1977, V1978, V1979)
				return

			} else {
				tmp9836 := PrimIsPair(V1976)

				var ifres9816 Obj

				if True == tmp9836 {
					tmp9834 := PrimHead(V1976)

					tmp9835 := PrimIsPair(tmp9834)

					var ifres9818 Obj

					if True == tmp9835 {
						tmp9831 := PrimHead(V1976)

						tmp9832 := PrimHead(tmp9831)

						tmp9833 := PrimEqual(symshen_4_1m, tmp9832)

						var ifres9820 Obj

						if True == tmp9833 {
							tmp9828 := PrimHead(V1976)

							tmp9829 := PrimTail(tmp9828)

							tmp9830 := PrimIsPair(tmp9829)

							var ifres9822 Obj

							if True == tmp9830 {
								tmp9824 := PrimHead(V1976)

								tmp9825 := PrimTail(tmp9824)

								tmp9826 := PrimTail(tmp9825)

								tmp9827 := PrimEqual(Nil, tmp9826)

								var ifres9823 Obj

								if True == tmp9827 {
									ifres9823 = True

								} else {
									ifres9823 = False

								}

								ifres9822 = ifres9823

							} else {
								ifres9822 = False

							}

							var ifres9821 Obj

							if True == ifres9822 {
								ifres9821 = True

							} else {
								ifres9821 = False

							}

							ifres9820 = ifres9821

						} else {
							ifres9820 = False

						}

						var ifres9819 Obj

						if True == ifres9820 {
							ifres9819 = True

						} else {
							ifres9819 = False

						}

						ifres9818 = ifres9819

					} else {
						ifres9818 = False

					}

					var ifres9817 Obj

					if True == ifres9818 {
						ifres9817 = True

					} else {
						ifres9817 = False

					}

					ifres9816 = ifres9817

				} else {
					ifres9816 = False

				}

				if True == ifres9816 {
					tmp9695 := PrimHead(V1976)

					tmp9696 := PrimTail(tmp9695)

					tmp9697 := PrimHead(tmp9696)

					tmp9698 := PrimTail(V1976)

					tmp9699 := PrimCons(V1975, tmp9698)

					tmp9700 := PrimCons(tmp9697, tmp9699)

					tmp9701 := PrimCons(symshen_4_1m, tmp9700)

					__e.TailApply(PrimFunc(symshen_4compile_1head), V1975, tmp9701, V1977, V1978, V1979)
					return

				} else {
					tmp9814 := PrimIsPair(V1976)

					var ifres9810 Obj

					if True == tmp9814 {
						tmp9812 := PrimHead(V1976)

						tmp9813 := PrimEqual(symshen_4_1m, tmp9812)

						var ifres9811 Obj

						if True == tmp9813 {
							ifres9811 = True

						} else {
							ifres9811 = False

						}

						ifres9810 = ifres9811

					} else {
						ifres9810 = False

					}

					if True == ifres9810 {
						tmp9702 := PrimTail(V1976)

						__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp9702, V1977, V1978, V1979)
						return

					} else {
						tmp9808 := PrimIsPair(V1976)

						var ifres9804 Obj

						if True == tmp9808 {
							tmp9806 := PrimHead(V1976)

							tmp9807 := PrimEqual(symshen_4_7m, tmp9806)

							var ifres9805 Obj

							if True == tmp9807 {
								ifres9805 = True

							} else {
								ifres9805 = False

							}

							ifres9804 = ifres9805

						} else {
							ifres9804 = False

						}

						if True == ifres9804 {
							tmp9703 := PrimTail(V1976)

							__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9703, V1977, V1978, V1979)
							return

						} else {
							tmp9802 := PrimIsPair(V1976)

							var ifres9795 Obj

							if True == tmp9802 {
								tmp9801 := PrimIsPair(V1977)

								var ifres9797 Obj

								if True == tmp9801 {
									tmp9799 := PrimHead(V1976)

									tmp9800 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp9799)

									var ifres9798 Obj

									if True == tmp9800 {
										ifres9798 = True

									} else {
										ifres9798 = False

									}

									ifres9797 = ifres9798

								} else {
									ifres9797 = False

								}

								var ifres9796 Obj

								if True == ifres9797 {
									ifres9796 = True

								} else {
									ifres9796 = False

								}

								ifres9795 = ifres9796

							} else {
								ifres9795 = False

							}

							if True == ifres9795 {
								tmp9704 := PrimTail(V1976)

								tmp9705 := PrimTail(V1977)

								__e.TailApply(PrimFunc(symshen_4compile_1head), V1975, tmp9704, tmp9705, V1978, V1979)
								return

							} else {
								tmp9793 := PrimIsPair(V1976)

								var ifres9789 Obj

								if True == tmp9793 {
									tmp9791 := PrimHead(V1976)

									tmp9792 := PrimIsVariable(tmp9791)

									var ifres9790 Obj

									if True == tmp9792 {
										ifres9790 = True

									} else {
										ifres9790 = False

									}

									ifres9789 = ifres9790

								} else {
									ifres9789 = False

								}

								if True == ifres9789 {
									__e.TailApply(PrimFunc(symshen_4variable_1case), V1975, V1976, V1977, V1978, V1979)
									return
								} else {
									tmp9787 := PrimEqual(symshen_4_1m, V1975)

									var ifres9780 Obj

									if True == tmp9787 {
										tmp9786 := PrimIsPair(V1976)

										var ifres9782 Obj

										if True == tmp9786 {
											tmp9784 := PrimHead(V1976)

											tmp9785 := Call(__e, PrimFunc(symatom_2), tmp9784)

											var ifres9783 Obj

											if True == tmp9785 {
												ifres9783 = True

											} else {
												ifres9783 = False

											}

											ifres9782 = ifres9783

										} else {
											ifres9782 = False

										}

										var ifres9781 Obj

										if True == ifres9782 {
											ifres9781 = True

										} else {
											ifres9781 = False

										}

										ifres9780 = ifres9781

									} else {
										ifres9780 = False

									}

									if True == ifres9780 {
										__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V1976, V1977, V1978, V1979)
										return
									} else {
										tmp9778 := PrimEqual(symshen_4_1m, V1975)

										var ifres9748 Obj

										if True == tmp9778 {
											tmp9777 := PrimIsPair(V1976)

											var ifres9750 Obj

											if True == tmp9777 {
												tmp9775 := PrimHead(V1976)

												tmp9776 := PrimIsPair(tmp9775)

												var ifres9752 Obj

												if True == tmp9776 {
													tmp9772 := PrimHead(V1976)

													tmp9773 := PrimHead(tmp9772)

													tmp9774 := PrimEqual(symcons, tmp9773)

													var ifres9754 Obj

													if True == tmp9774 {
														tmp9769 := PrimHead(V1976)

														tmp9770 := PrimTail(tmp9769)

														tmp9771 := PrimIsPair(tmp9770)

														var ifres9756 Obj

														if True == tmp9771 {
															tmp9765 := PrimHead(V1976)

															tmp9766 := PrimTail(tmp9765)

															tmp9767 := PrimTail(tmp9766)

															tmp9768 := PrimIsPair(tmp9767)

															var ifres9758 Obj

															if True == tmp9768 {
																tmp9760 := PrimHead(V1976)

																tmp9761 := PrimTail(tmp9760)

																tmp9762 := PrimTail(tmp9761)

																tmp9763 := PrimTail(tmp9762)

																tmp9764 := PrimEqual(Nil, tmp9763)

																var ifres9759 Obj

																if True == tmp9764 {
																	ifres9759 = True

																} else {
																	ifres9759 = False

																}

																ifres9758 = ifres9759

															} else {
																ifres9758 = False

															}

															var ifres9757 Obj

															if True == ifres9758 {
																ifres9757 = True

															} else {
																ifres9757 = False

															}

															ifres9756 = ifres9757

														} else {
															ifres9756 = False

														}

														var ifres9755 Obj

														if True == ifres9756 {
															ifres9755 = True

														} else {
															ifres9755 = False

														}

														ifres9754 = ifres9755

													} else {
														ifres9754 = False

													}

													var ifres9753 Obj

													if True == ifres9754 {
														ifres9753 = True

													} else {
														ifres9753 = False

													}

													ifres9752 = ifres9753

												} else {
													ifres9752 = False

												}

												var ifres9751 Obj

												if True == ifres9752 {
													ifres9751 = True

												} else {
													ifres9751 = False

												}

												ifres9750 = ifres9751

											} else {
												ifres9750 = False

											}

											var ifres9749 Obj

											if True == ifres9750 {
												ifres9749 = True

											} else {
												ifres9749 = False

											}

											ifres9748 = ifres9749

										} else {
											ifres9748 = False

										}

										if True == ifres9748 {
											__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V1976, V1977, V1978, V1979)
											return
										} else {
											tmp9746 := PrimEqual(symshen_4_7m, V1975)

											var ifres9739 Obj

											if True == tmp9746 {
												tmp9745 := PrimIsPair(V1976)

												var ifres9741 Obj

												if True == tmp9745 {
													tmp9743 := PrimHead(V1976)

													tmp9744 := Call(__e, PrimFunc(symatom_2), tmp9743)

													var ifres9742 Obj

													if True == tmp9744 {
														ifres9742 = True

													} else {
														ifres9742 = False

													}

													ifres9741 = ifres9742

												} else {
													ifres9741 = False

												}

												var ifres9740 Obj

												if True == ifres9741 {
													ifres9740 = True

												} else {
													ifres9740 = False

												}

												ifres9739 = ifres9740

											} else {
												ifres9739 = False

											}

											if True == ifres9739 {
												__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V1976, V1977, V1978, V1979)
												return
											} else {
												tmp9737 := PrimEqual(symshen_4_7m, V1975)

												var ifres9707 Obj

												if True == tmp9737 {
													tmp9736 := PrimIsPair(V1976)

													var ifres9709 Obj

													if True == tmp9736 {
														tmp9734 := PrimHead(V1976)

														tmp9735 := PrimIsPair(tmp9734)

														var ifres9711 Obj

														if True == tmp9735 {
															tmp9731 := PrimHead(V1976)

															tmp9732 := PrimHead(tmp9731)

															tmp9733 := PrimEqual(symcons, tmp9732)

															var ifres9713 Obj

															if True == tmp9733 {
																tmp9728 := PrimHead(V1976)

																tmp9729 := PrimTail(tmp9728)

																tmp9730 := PrimIsPair(tmp9729)

																var ifres9715 Obj

																if True == tmp9730 {
																	tmp9724 := PrimHead(V1976)

																	tmp9725 := PrimTail(tmp9724)

																	tmp9726 := PrimTail(tmp9725)

																	tmp9727 := PrimIsPair(tmp9726)

																	var ifres9717 Obj

																	if True == tmp9727 {
																		tmp9719 := PrimHead(V1976)

																		tmp9720 := PrimTail(tmp9719)

																		tmp9721 := PrimTail(tmp9720)

																		tmp9722 := PrimTail(tmp9721)

																		tmp9723 := PrimEqual(Nil, tmp9722)

																		var ifres9718 Obj

																		if True == tmp9723 {
																			ifres9718 = True

																		} else {
																			ifres9718 = False

																		}

																		ifres9717 = ifres9718

																	} else {
																		ifres9717 = False

																	}

																	var ifres9716 Obj

																	if True == ifres9717 {
																		ifres9716 = True

																	} else {
																		ifres9716 = False

																	}

																	ifres9715 = ifres9716

																} else {
																	ifres9715 = False

																}

																var ifres9714 Obj

																if True == ifres9715 {
																	ifres9714 = True

																} else {
																	ifres9714 = False

																}

																ifres9713 = ifres9714

															} else {
																ifres9713 = False

															}

															var ifres9712 Obj

															if True == ifres9713 {
																ifres9712 = True

															} else {
																ifres9712 = False

															}

															ifres9711 = ifres9712

														} else {
															ifres9711 = False

														}

														var ifres9710 Obj

														if True == ifres9711 {
															ifres9710 = True

														} else {
															ifres9710 = False

														}

														ifres9709 = ifres9710

													} else {
														ifres9709 = False

													}

													var ifres9708 Obj

													if True == ifres9709 {
														ifres9708 = True

													} else {
														ifres9708 = False

													}

													ifres9707 = ifres9708

												} else {
													ifres9707 = False

												}

												if True == ifres9707 {
													__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V1976, V1977, V1978, V1979)
													return
												} else {
													__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-head")))
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

			}

		}

	}, 5)

	tmp9864 := Call(__e, ns2_1set, symshen_4compile_1head, tmp9687)

	_ = tmp9864

	tmp9865 := MakeNative(func(__e *ControlFlow) {
		V1990 := __e.Get(1)
		_ = V1990
		V1991 := __e.Get(2)
		_ = V1991
		V1992 := __e.Get(3)
		_ = V1992
		V1993 := __e.Get(4)
		_ = V1993
		V1994 := __e.Get(5)
		_ = V1994
		tmp9886 := PrimIsPair(V1991)

		var ifres9883 Obj

		if True == tmp9886 {
			tmp9885 := PrimIsPair(V1992)

			var ifres9884 Obj

			if True == tmp9885 {
				ifres9884 = True

			} else {
				ifres9884 = False

			}

			ifres9883 = ifres9884

		} else {
			ifres9883 = False

		}

		if True == ifres9883 {
			tmp9880 := PrimHead(V1992)

			tmp9881 := PrimIsVariable(tmp9880)

			if True == tmp9881 {
				tmp9866 := PrimTail(V1991)

				tmp9867 := PrimTail(V1992)

				tmp9868 := PrimHead(V1992)

				tmp9869 := PrimHead(V1991)

				tmp9870 := Call(__e, PrimFunc(symsubst), tmp9868, tmp9869, V1994)

				__e.TailApply(PrimFunc(symshen_4compile_1head), V1990, tmp9866, tmp9867, V1993, tmp9870)
				return

			} else {
				tmp9871 := PrimHead(V1991)

				tmp9872 := PrimHead(V1992)

				tmp9873 := PrimTail(V1991)

				tmp9874 := PrimTail(V1992)

				tmp9875 := Call(__e, PrimFunc(symshen_4compile_1head), V1990, tmp9873, tmp9874, V1993, V1994)

				tmp9876 := PrimCons(tmp9875, Nil)

				tmp9877 := PrimCons(tmp9872, tmp9876)

				tmp9878 := PrimCons(tmp9871, tmp9877)

				__e.Return(PrimCons(symlet, tmp9878))
				return

			}

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
			return
		}

	}, 5)

	tmp9887 := Call(__e, ns2_1set, symshen_4variable_1case, tmp9865)

	_ = tmp9887

	tmp9888 := MakeNative(func(__e *ControlFlow) {
		V2003 := __e.Get(1)
		_ = V2003
		V2004 := __e.Get(2)
		_ = V2004
		V2005 := __e.Get(3)
		_ = V2005
		V2006 := __e.Get(4)
		_ = V2006
		tmp9913 := PrimIsPair(V2003)

		var ifres9910 Obj

		if True == tmp9913 {
			tmp9912 := PrimIsPair(V2004)

			var ifres9911 Obj

			if True == tmp9912 {
				ifres9911 = True

			} else {
				ifres9911 = False

			}

			ifres9910 = ifres9911

		} else {
			ifres9910 = False

		}

		if True == ifres9910 {
			tmp9889 := MakeNative(func(__e *ControlFlow) {
				W2007 := __e.Get(1)
				_ = W2007
				tmp9890 := PrimHead(V2004)

				tmp9891 := PrimCons(V2005, Nil)

				tmp9892 := PrimCons(tmp9890, tmp9891)

				tmp9893 := PrimCons(symshen_4lazyderef, tmp9892)

				tmp9894 := PrimHead(V2003)

				tmp9895 := PrimCons(tmp9894, Nil)

				tmp9896 := PrimCons(W2007, tmp9895)

				tmp9897 := PrimCons(sym_a, tmp9896)

				tmp9898 := PrimTail(V2003)

				tmp9899 := PrimTail(V2004)

				tmp9900 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp9898, tmp9899, V2005, V2006)

				tmp9901 := PrimCons(False, Nil)

				tmp9902 := PrimCons(tmp9900, tmp9901)

				tmp9903 := PrimCons(tmp9897, tmp9902)

				tmp9904 := PrimCons(symif, tmp9903)

				tmp9905 := PrimCons(tmp9904, Nil)

				tmp9906 := PrimCons(tmp9893, tmp9905)

				tmp9907 := PrimCons(W2007, tmp9906)

				__e.Return(PrimCons(symlet, tmp9907))
				return

			}, 1)

			tmp9908 := Call(__e, PrimFunc(symgensym), symTm)

			__e.TailApply(tmp9889, tmp9908)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
			return
		}

	}, 4)

	tmp9914 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp9888)

	_ = tmp9914

	tmp9915 := MakeNative(func(__e *ControlFlow) {
		V2016 := __e.Get(1)
		_ = V2016
		V2017 := __e.Get(2)
		_ = V2017
		V2018 := __e.Get(3)
		_ = V2018
		V2019 := __e.Get(4)
		_ = V2019
		tmp9980 := PrimIsPair(V2016)

		var ifres9950 Obj

		if True == tmp9980 {
			tmp9978 := PrimHead(V2016)

			tmp9979 := PrimIsPair(tmp9978)

			var ifres9952 Obj

			if True == tmp9979 {
				tmp9975 := PrimHead(V2016)

				tmp9976 := PrimHead(tmp9975)

				tmp9977 := PrimEqual(symcons, tmp9976)

				var ifres9954 Obj

				if True == tmp9977 {
					tmp9972 := PrimHead(V2016)

					tmp9973 := PrimTail(tmp9972)

					tmp9974 := PrimIsPair(tmp9973)

					var ifres9956 Obj

					if True == tmp9974 {
						tmp9968 := PrimHead(V2016)

						tmp9969 := PrimTail(tmp9968)

						tmp9970 := PrimTail(tmp9969)

						tmp9971 := PrimIsPair(tmp9970)

						var ifres9958 Obj

						if True == tmp9971 {
							tmp9963 := PrimHead(V2016)

							tmp9964 := PrimTail(tmp9963)

							tmp9965 := PrimTail(tmp9964)

							tmp9966 := PrimTail(tmp9965)

							tmp9967 := PrimEqual(Nil, tmp9966)

							var ifres9960 Obj

							if True == tmp9967 {
								tmp9962 := PrimIsPair(V2017)

								var ifres9961 Obj

								if True == tmp9962 {
									ifres9961 = True

								} else {
									ifres9961 = False

								}

								ifres9960 = ifres9961

							} else {
								ifres9960 = False

							}

							var ifres9959 Obj

							if True == ifres9960 {
								ifres9959 = True

							} else {
								ifres9959 = False

							}

							ifres9958 = ifres9959

						} else {
							ifres9958 = False

						}

						var ifres9957 Obj

						if True == ifres9958 {
							ifres9957 = True

						} else {
							ifres9957 = False

						}

						ifres9956 = ifres9957

					} else {
						ifres9956 = False

					}

					var ifres9955 Obj

					if True == ifres9956 {
						ifres9955 = True

					} else {
						ifres9955 = False

					}

					ifres9954 = ifres9955

				} else {
					ifres9954 = False

				}

				var ifres9953 Obj

				if True == ifres9954 {
					ifres9953 = True

				} else {
					ifres9953 = False

				}

				ifres9952 = ifres9953

			} else {
				ifres9952 = False

			}

			var ifres9951 Obj

			if True == ifres9952 {
				ifres9951 = True

			} else {
				ifres9951 = False

			}

			ifres9950 = ifres9951

		} else {
			ifres9950 = False

		}

		if True == ifres9950 {
			tmp9916 := MakeNative(func(__e *ControlFlow) {
				W2020 := __e.Get(1)
				_ = W2020
				tmp9917 := PrimHead(V2017)

				tmp9918 := PrimCons(V2018, Nil)

				tmp9919 := PrimCons(tmp9917, tmp9918)

				tmp9920 := PrimCons(symshen_4lazyderef, tmp9919)

				tmp9921 := PrimCons(W2020, Nil)

				tmp9922 := PrimCons(symcons_2, tmp9921)

				tmp9923 := PrimHead(V2016)

				tmp9924 := PrimTail(tmp9923)

				tmp9925 := PrimHead(tmp9924)

				tmp9926 := PrimHead(V2016)

				tmp9927 := PrimTail(tmp9926)

				tmp9928 := PrimTail(tmp9927)

				tmp9929 := PrimHead(tmp9928)

				tmp9930 := PrimTail(V2016)

				tmp9931 := PrimCons(tmp9929, tmp9930)

				tmp9932 := PrimCons(tmp9925, tmp9931)

				tmp9933 := PrimCons(W2020, Nil)

				tmp9934 := PrimCons(symhd, tmp9933)

				tmp9935 := PrimCons(W2020, Nil)

				tmp9936 := PrimCons(symtl, tmp9935)

				tmp9937 := PrimTail(V2017)

				tmp9938 := PrimCons(tmp9936, tmp9937)

				tmp9939 := PrimCons(tmp9934, tmp9938)

				tmp9940 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp9932, tmp9939, V2018, V2019)

				tmp9941 := PrimCons(False, Nil)

				tmp9942 := PrimCons(tmp9940, tmp9941)

				tmp9943 := PrimCons(tmp9922, tmp9942)

				tmp9944 := PrimCons(symif, tmp9943)

				tmp9945 := PrimCons(tmp9944, Nil)

				tmp9946 := PrimCons(tmp9920, tmp9945)

				tmp9947 := PrimCons(W2020, tmp9946)

				__e.Return(PrimCons(symlet, tmp9947))
				return

			}, 1)

			tmp9948 := Call(__e, PrimFunc(symgensym), symTm)

			__e.TailApply(tmp9916, tmp9948)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
			return
		}

	}, 4)

	tmp9981 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp9915)

	_ = tmp9981

	tmp9982 := MakeNative(func(__e *ControlFlow) {
		V2029 := __e.Get(1)
		_ = V2029
		V2030 := __e.Get(2)
		_ = V2030
		V2031 := __e.Get(3)
		_ = V2031
		V2032 := __e.Get(4)
		_ = V2032
		tmp10028 := PrimIsPair(V2029)

		var ifres10025 Obj

		if True == tmp10028 {
			tmp10027 := PrimIsPair(V2030)

			var ifres10026 Obj

			if True == tmp10027 {
				ifres10026 = True

			} else {
				ifres10026 = False

			}

			ifres10025 = ifres10026

		} else {
			ifres10025 = False

		}

		if True == ifres10025 {
			tmp9983 := MakeNative(func(__e *ControlFlow) {
				W2033 := __e.Get(1)
				_ = W2033
				tmp9984 := MakeNative(func(__e *ControlFlow) {
					W2034 := __e.Get(1)
					_ = W2034
					tmp9985 := PrimHead(V2030)

					tmp9986 := PrimCons(V2031, Nil)

					tmp9987 := PrimCons(tmp9985, tmp9986)

					tmp9988 := PrimCons(symshen_4lazyderef, tmp9987)

					tmp9989 := PrimTail(V2029)

					tmp9990 := PrimTail(V2030)

					tmp9991 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9989, tmp9990, V2031, V2032)

					tmp9992 := PrimCons(tmp9991, Nil)

					tmp9993 := PrimCons(symfreeze, tmp9992)

					tmp9994 := PrimHead(V2029)

					tmp9995 := PrimCons(tmp9994, Nil)

					tmp9996 := PrimCons(W2033, tmp9995)

					tmp9997 := PrimCons(sym_a, tmp9996)

					tmp9998 := PrimCons(W2034, Nil)

					tmp9999 := PrimCons(symthaw, tmp9998)

					tmp10000 := PrimCons(W2033, Nil)

					tmp10001 := PrimCons(symshen_4pvar_2, tmp10000)

					tmp10002 := PrimHead(V2029)

					tmp10003 := Call(__e, PrimFunc(symshen_4demode), tmp10002)

					tmp10004 := PrimCons(W2034, Nil)

					tmp10005 := PrimCons(V2031, tmp10004)

					tmp10006 := PrimCons(tmp10003, tmp10005)

					tmp10007 := PrimCons(W2033, tmp10006)

					tmp10008 := PrimCons(symshen_4bind_b, tmp10007)

					tmp10009 := PrimCons(False, Nil)

					tmp10010 := PrimCons(tmp10008, tmp10009)

					tmp10011 := PrimCons(tmp10001, tmp10010)

					tmp10012 := PrimCons(symif, tmp10011)

					tmp10013 := PrimCons(tmp10012, Nil)

					tmp10014 := PrimCons(tmp9999, tmp10013)

					tmp10015 := PrimCons(tmp9997, tmp10014)

					tmp10016 := PrimCons(symif, tmp10015)

					tmp10017 := PrimCons(tmp10016, Nil)

					tmp10018 := PrimCons(tmp9993, tmp10017)

					tmp10019 := PrimCons(W2034, tmp10018)

					tmp10020 := PrimCons(tmp9988, tmp10019)

					tmp10021 := PrimCons(W2033, tmp10020)

					__e.Return(PrimCons(symlet, tmp10021))
					return

				}, 1)

				tmp10022 := Call(__e, PrimFunc(symgensym), symGoTo)

				__e.TailApply(tmp9984, tmp10022)
				return

			}, 1)

			tmp10023 := Call(__e, PrimFunc(symgensym), symTm)

			__e.TailApply(tmp9983, tmp10023)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
			return
		}

	}, 4)

	tmp10029 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp9982)

	_ = tmp10029

	tmp10030 := MakeNative(func(__e *ControlFlow) {
		V2043 := __e.Get(1)
		_ = V2043
		V2044 := __e.Get(2)
		_ = V2044
		V2045 := __e.Get(3)
		_ = V2045
		V2046 := __e.Get(4)
		_ = V2046
		tmp10126 := PrimIsPair(V2043)

		var ifres10096 Obj

		if True == tmp10126 {
			tmp10124 := PrimHead(V2043)

			tmp10125 := PrimIsPair(tmp10124)

			var ifres10098 Obj

			if True == tmp10125 {
				tmp10121 := PrimHead(V2043)

				tmp10122 := PrimHead(tmp10121)

				tmp10123 := PrimEqual(symcons, tmp10122)

				var ifres10100 Obj

				if True == tmp10123 {
					tmp10118 := PrimHead(V2043)

					tmp10119 := PrimTail(tmp10118)

					tmp10120 := PrimIsPair(tmp10119)

					var ifres10102 Obj

					if True == tmp10120 {
						tmp10114 := PrimHead(V2043)

						tmp10115 := PrimTail(tmp10114)

						tmp10116 := PrimTail(tmp10115)

						tmp10117 := PrimIsPair(tmp10116)

						var ifres10104 Obj

						if True == tmp10117 {
							tmp10109 := PrimHead(V2043)

							tmp10110 := PrimTail(tmp10109)

							tmp10111 := PrimTail(tmp10110)

							tmp10112 := PrimTail(tmp10111)

							tmp10113 := PrimEqual(Nil, tmp10112)

							var ifres10106 Obj

							if True == tmp10113 {
								tmp10108 := PrimIsPair(V2044)

								var ifres10107 Obj

								if True == tmp10108 {
									ifres10107 = True

								} else {
									ifres10107 = False

								}

								ifres10106 = ifres10107

							} else {
								ifres10106 = False

							}

							var ifres10105 Obj

							if True == ifres10106 {
								ifres10105 = True

							} else {
								ifres10105 = False

							}

							ifres10104 = ifres10105

						} else {
							ifres10104 = False

						}

						var ifres10103 Obj

						if True == ifres10104 {
							ifres10103 = True

						} else {
							ifres10103 = False

						}

						ifres10102 = ifres10103

					} else {
						ifres10102 = False

					}

					var ifres10101 Obj

					if True == ifres10102 {
						ifres10101 = True

					} else {
						ifres10101 = False

					}

					ifres10100 = ifres10101

				} else {
					ifres10100 = False

				}

				var ifres10099 Obj

				if True == ifres10100 {
					ifres10099 = True

				} else {
					ifres10099 = False

				}

				ifres10098 = ifres10099

			} else {
				ifres10098 = False

			}

			var ifres10097 Obj

			if True == ifres10098 {
				ifres10097 = True

			} else {
				ifres10097 = False

			}

			ifres10096 = ifres10097

		} else {
			ifres10096 = False

		}

		if True == ifres10096 {
			tmp10031 := MakeNative(func(__e *ControlFlow) {
				W2047 := __e.Get(1)
				_ = W2047
				tmp10032 := MakeNative(func(__e *ControlFlow) {
					W2048 := __e.Get(1)
					_ = W2048
					tmp10033 := MakeNative(func(__e *ControlFlow) {
						W2049 := __e.Get(1)
						_ = W2049
						tmp10034 := MakeNative(func(__e *ControlFlow) {
							W2050 := __e.Get(1)
							_ = W2050
							tmp10035 := MakeNative(func(__e *ControlFlow) {
								W2051 := __e.Get(1)
								_ = W2051
								tmp10036 := PrimHead(V2044)

								tmp10037 := PrimCons(V2045, Nil)

								tmp10038 := PrimCons(tmp10036, tmp10037)

								tmp10039 := PrimCons(symshen_4lazyderef, tmp10038)

								tmp10040 := PrimTail(V2043)

								tmp10041 := PrimTail(V2044)

								tmp10042 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10040, tmp10041, V2045, V2046)

								tmp10043 := Call(__e, PrimFunc(symshen_4goto), W2049, tmp10042)

								tmp10044 := PrimCons(W2047, Nil)

								tmp10045 := PrimCons(symcons_2, tmp10044)

								tmp10046 := PrimHead(V2043)

								tmp10047 := PrimTail(tmp10046)

								tmp10048 := PrimCons(W2047, Nil)

								tmp10049 := PrimCons(symhd, tmp10048)

								tmp10050 := PrimCons(W2047, Nil)

								tmp10051 := PrimCons(symtl, tmp10050)

								tmp10052 := PrimCons(tmp10051, Nil)

								tmp10053 := PrimCons(tmp10049, tmp10052)

								tmp10054 := Call(__e, PrimFunc(symshen_4invoke), W2048, W2049)

								tmp10055 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10047, tmp10053, V2045, tmp10054)

								tmp10056 := PrimCons(W2047, Nil)

								tmp10057 := PrimCons(symshen_4pvar_2, tmp10056)

								tmp10058 := Call(__e, PrimFunc(symshen_4demode), W2050)

								tmp10059 := Call(__e, PrimFunc(symshen_4invoke), W2048, W2049)

								tmp10060 := PrimCons(tmp10059, Nil)

								tmp10061 := PrimCons(symfreeze, tmp10060)

								tmp10062 := PrimCons(tmp10061, Nil)

								tmp10063 := PrimCons(V2045, tmp10062)

								tmp10064 := PrimCons(tmp10058, tmp10063)

								tmp10065 := PrimCons(W2047, tmp10064)

								tmp10066 := PrimCons(symshen_4bind_b, tmp10065)

								tmp10067 := Call(__e, PrimFunc(symshen_4stpart), W2051, tmp10066, V2045)

								tmp10068 := PrimCons(False, Nil)

								tmp10069 := PrimCons(tmp10067, tmp10068)

								tmp10070 := PrimCons(tmp10057, tmp10069)

								tmp10071 := PrimCons(symif, tmp10070)

								tmp10072 := PrimCons(tmp10071, Nil)

								tmp10073 := PrimCons(tmp10055, tmp10072)

								tmp10074 := PrimCons(tmp10045, tmp10073)

								tmp10075 := PrimCons(symif, tmp10074)

								tmp10076 := PrimCons(tmp10075, Nil)

								tmp10077 := PrimCons(tmp10043, tmp10076)

								tmp10078 := PrimCons(W2048, tmp10077)

								tmp10079 := PrimCons(tmp10039, tmp10078)

								tmp10080 := PrimCons(W2047, tmp10079)

								__e.Return(PrimCons(symlet, tmp10080))
								return

							}, 1)

							tmp10081 := Call(__e, PrimFunc(symshen_4extract_1vars), W2050)

							__e.TailApply(tmp10035, tmp10081)
							return

						}, 1)

						tmp10082 := PrimHead(V2043)

						tmp10083 := Call(__e, PrimFunc(symshen_4tame), tmp10082)

						__e.TailApply(tmp10034, tmp10083)
						return

					}, 1)

					tmp10084 := PrimHead(V2043)

					tmp10085 := PrimTail(tmp10084)

					tmp10086 := PrimHead(tmp10085)

					tmp10087 := PrimHead(V2043)

					tmp10088 := PrimTail(tmp10087)

					tmp10089 := PrimTail(tmp10088)

					tmp10090 := PrimHead(tmp10089)

					tmp10091 := PrimCons(tmp10086, tmp10090)

					tmp10092 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp10091)

					__e.TailApply(tmp10033, tmp10092)
					return

				}, 1)

				tmp10093 := Call(__e, PrimFunc(symgensym), symGoTo)

				__e.TailApply(tmp10032, tmp10093)
				return

			}, 1)

			tmp10094 := Call(__e, PrimFunc(symgensym), symTm)

			__e.TailApply(tmp10031, tmp10094)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
			return
		}

	}, 4)

	tmp10127 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp10030)

	_ = tmp10127

	tmp10128 := MakeNative(func(__e *ControlFlow) {
		V2052 := __e.Get(1)
		_ = V2052
		tmp10165 := PrimIsPair(V2052)

		var ifres10152 Obj

		if True == tmp10165 {
			tmp10163 := PrimHead(V2052)

			tmp10164 := PrimEqual(symshen_4_7m, tmp10163)

			var ifres10154 Obj

			if True == tmp10164 {
				tmp10161 := PrimTail(V2052)

				tmp10162 := PrimIsPair(tmp10161)

				var ifres10156 Obj

				if True == tmp10162 {
					tmp10158 := PrimTail(V2052)

					tmp10159 := PrimTail(tmp10158)

					tmp10160 := PrimEqual(Nil, tmp10159)

					var ifres10157 Obj

					if True == tmp10160 {
						ifres10157 = True

					} else {
						ifres10157 = False

					}

					ifres10156 = ifres10157

				} else {
					ifres10156 = False

				}

				var ifres10155 Obj

				if True == ifres10156 {
					ifres10155 = True

				} else {
					ifres10155 = False

				}

				ifres10154 = ifres10155

			} else {
				ifres10154 = False

			}

			var ifres10153 Obj

			if True == ifres10154 {
				ifres10153 = True

			} else {
				ifres10153 = False

			}

			ifres10152 = ifres10153

		} else {
			ifres10152 = False

		}

		if True == ifres10152 {
			tmp10129 := PrimTail(V2052)

			tmp10130 := PrimHead(tmp10129)

			__e.TailApply(PrimFunc(symshen_4demode), tmp10130)
			return

		} else {
			tmp10150 := PrimIsPair(V2052)

			var ifres10137 Obj

			if True == tmp10150 {
				tmp10148 := PrimHead(V2052)

				tmp10149 := PrimEqual(symshen_4_1m, tmp10148)

				var ifres10139 Obj

				if True == tmp10149 {
					tmp10146 := PrimTail(V2052)

					tmp10147 := PrimIsPair(tmp10146)

					var ifres10141 Obj

					if True == tmp10147 {
						tmp10143 := PrimTail(V2052)

						tmp10144 := PrimTail(tmp10143)

						tmp10145 := PrimEqual(Nil, tmp10144)

						var ifres10142 Obj

						if True == tmp10145 {
							ifres10142 = True

						} else {
							ifres10142 = False

						}

						ifres10141 = ifres10142

					} else {
						ifres10141 = False

					}

					var ifres10140 Obj

					if True == ifres10141 {
						ifres10140 = True

					} else {
						ifres10140 = False

					}

					ifres10139 = ifres10140

				} else {
					ifres10139 = False

				}

				var ifres10138 Obj

				if True == ifres10139 {
					ifres10138 = True

				} else {
					ifres10138 = False

				}

				ifres10137 = ifres10138

			} else {
				ifres10137 = False

			}

			if True == ifres10137 {
				tmp10131 := PrimTail(V2052)

				tmp10132 := PrimHead(tmp10131)

				__e.TailApply(PrimFunc(symshen_4demode), tmp10132)
				return

			} else {
				tmp10135 := PrimIsPair(V2052)

				if True == tmp10135 {
					tmp10133 := MakeNative(func(__e *ControlFlow) {
						Z2053 := __e.Get(1)
						_ = Z2053
						__e.TailApply(PrimFunc(symshen_4demode), Z2053)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp10133, V2052)
					return

				} else {
					__e.Return(V2052)
					return
				}

			}

		}

	}, 1)

	tmp10166 := Call(__e, ns2_1set, symshen_4demode, tmp10128)

	_ = tmp10166

	tmp10167 := MakeNative(func(__e *ControlFlow) {
		V2054 := __e.Get(1)
		_ = V2054
		tmp10172 := Call(__e, PrimFunc(symshen_4wildcard_2), V2054)

		if True == tmp10172 {
			__e.TailApply(PrimFunc(symgensym), symY)
			return
		} else {
			tmp10170 := PrimIsPair(V2054)

			if True == tmp10170 {
				tmp10168 := MakeNative(func(__e *ControlFlow) {
					Z2055 := __e.Get(1)
					_ = Z2055
					__e.TailApply(PrimFunc(symshen_4tame), Z2055)
					return
				}, 1)

				__e.TailApply(PrimFunc(symmap), tmp10168, V2054)
				return

			} else {
				__e.Return(V2054)
				return
			}

		}

	}, 1)

	tmp10173 := Call(__e, ns2_1set, symshen_4tame, tmp10167)

	_ = tmp10173

	tmp10174 := MakeNative(func(__e *ControlFlow) {
		V2056 := __e.Get(1)
		_ = V2056
		V2057 := __e.Get(2)
		_ = V2057
		tmp10177 := PrimEqual(Nil, V2056)

		if True == tmp10177 {
			tmp10175 := PrimCons(V2057, Nil)

			__e.Return(PrimCons(symfreeze, tmp10175))
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4goto_1h), V2056, V2057)
			return
		}

	}, 2)

	tmp10178 := Call(__e, ns2_1set, symshen_4goto, tmp10174)

	_ = tmp10178

	tmp10179 := MakeNative(func(__e *ControlFlow) {
		V2058 := __e.Get(1)
		_ = V2058
		V2059 := __e.Get(2)
		_ = V2059
		tmp10188 := PrimEqual(Nil, V2058)

		if True == tmp10188 {
			__e.Return(V2059)
			return
		} else {
			tmp10186 := PrimIsPair(V2058)

			if True == tmp10186 {
				tmp10180 := PrimHead(V2058)

				tmp10181 := PrimTail(V2058)

				tmp10182 := Call(__e, PrimFunc(symshen_4goto_1h), tmp10181, V2059)

				tmp10183 := PrimCons(tmp10182, Nil)

				tmp10184 := PrimCons(tmp10180, tmp10183)

				__e.Return(PrimCons(symlambda, tmp10184))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4goto_1h)
				return
			}

		}

	}, 2)

	tmp10189 := Call(__e, ns2_1set, symshen_4goto_1h, tmp10179)

	_ = tmp10189

	tmp10190 := MakeNative(func(__e *ControlFlow) {
		V2060 := __e.Get(1)
		_ = V2060
		V2061 := __e.Get(2)
		_ = V2061
		tmp10193 := PrimEqual(Nil, V2061)

		if True == tmp10193 {
			tmp10191 := PrimCons(V2060, Nil)

			__e.Return(PrimCons(symthaw, tmp10191))
			return

		} else {
			__e.Return(PrimCons(V2060, V2061))
			return
		}

	}, 2)

	tmp10194 := Call(__e, ns2_1set, symshen_4invoke, tmp10190)

	_ = tmp10194

	tmp10195 := MakeNative(func(__e *ControlFlow) {
		V2062 := __e.Get(1)
		_ = V2062
		__e.Return(PrimEqual(V2062, sym__))
		return
	}, 1)

	tmp10196 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp10195)

	_ = tmp10196

	tmp10197 := MakeNative(func(__e *ControlFlow) {
		V2063 := __e.Get(1)
		_ = V2063
		tmp10204 := PrimIsVector(V2063)

		if True == tmp10204 {
			tmp10199 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V2063, MakeNumber(0)))
				return
			}, 0)

			tmp10200 := MakeNative(func(__e *ControlFlow) {
				Z2064 := __e.Get(1)
				_ = Z2064
				__e.Return(symshen_4not_1pvar)
				return
			}, 1)

			tmp10201 := Call(__e, try_1catch, tmp10199, tmp10200)

			tmp10202 := PrimEqual(tmp10201, symshen_4pvar)

			if True == tmp10202 {
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

	tmp10205 := Call(__e, ns2_1set, symshen_4pvar_2, tmp10197)

	_ = tmp10205

	tmp10206 := MakeNative(func(__e *ControlFlow) {
		V2065 := __e.Get(1)
		_ = V2065
		V2066 := __e.Get(2)
		_ = V2066
		tmp10213 := Call(__e, PrimFunc(symshen_4pvar_2), V2065)

		if True == tmp10213 {
			tmp10207 := MakeNative(func(__e *ControlFlow) {
				W2067 := __e.Get(1)
				_ = W2067
				tmp10209 := PrimEqual(W2067, symshen_4_1null_1)

				if True == tmp10209 {
					__e.Return(V2065)
					return
				} else {
					__e.TailApply(PrimFunc(symshen_4lazyderef), W2067, V2066)
					return
				}

			}, 1)

			tmp10210 := PrimVectorGet(V2065, MakeNumber(1))

			tmp10211 := PrimVectorGet(V2066, tmp10210)

			__e.TailApply(tmp10207, tmp10211)
			return

		} else {
			__e.Return(V2065)
			return
		}

	}, 2)

	tmp10214 := Call(__e, ns2_1set, symshen_4lazyderef, tmp10206)

	_ = tmp10214

	tmp10215 := MakeNative(func(__e *ControlFlow) {
		V2068 := __e.Get(1)
		_ = V2068
		V2069 := __e.Get(2)
		_ = V2069
		tmp10228 := PrimIsPair(V2068)

		if True == tmp10228 {
			tmp10216 := PrimHead(V2068)

			tmp10217 := Call(__e, PrimFunc(symshen_4deref), tmp10216, V2069)

			tmp10218 := PrimTail(V2068)

			tmp10219 := Call(__e, PrimFunc(symshen_4deref), tmp10218, V2069)

			__e.Return(PrimCons(tmp10217, tmp10219))
			return

		} else {
			tmp10226 := Call(__e, PrimFunc(symshen_4pvar_2), V2068)

			if True == tmp10226 {
				tmp10220 := MakeNative(func(__e *ControlFlow) {
					W2070 := __e.Get(1)
					_ = W2070
					tmp10222 := PrimEqual(W2070, symshen_4_1null_1)

					if True == tmp10222 {
						__e.Return(V2068)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4deref), W2070, V2069)
						return
					}

				}, 1)

				tmp10223 := PrimVectorGet(V2068, MakeNumber(1))

				tmp10224 := PrimVectorGet(V2069, tmp10223)

				__e.TailApply(tmp10220, tmp10224)
				return

			} else {
				__e.Return(V2068)
				return
			}

		}

	}, 2)

	tmp10229 := Call(__e, ns2_1set, symshen_4deref, tmp10215)

	_ = tmp10229

	tmp10230 := MakeNative(func(__e *ControlFlow) {
		V2071 := __e.Get(1)
		_ = V2071
		V2072 := __e.Get(2)
		_ = V2072
		V2073 := __e.Get(3)
		_ = V2073
		V2074 := __e.Get(4)
		_ = V2074
		tmp10231 := MakeNative(func(__e *ControlFlow) {
			W2075 := __e.Get(1)
			_ = W2075
			tmp10232 := MakeNative(func(__e *ControlFlow) {
				W2076 := __e.Get(1)
				_ = W2076
				tmp10234 := PrimEqual(W2076, False)

				if True == tmp10234 {
					__e.TailApply(PrimFunc(symshen_4unwind), V2071, V2073, W2076)
					return
				} else {
					__e.Return(W2076)
					return
				}

			}, 1)

			tmp10235 := Call(__e, PrimFunc(symthaw), V2074)

			__e.TailApply(tmp10232, tmp10235)
			return

		}, 1)

		tmp10236 := Call(__e, PrimFunc(symshen_4bindv), V2071, V2072, V2073)

		__e.TailApply(tmp10231, tmp10236)
		return

	}, 4)

	tmp10237 := Call(__e, ns2_1set, symshen_4bind_b, tmp10230)

	_ = tmp10237

	tmp10238 := MakeNative(func(__e *ControlFlow) {
		V2077 := __e.Get(1)
		_ = V2077
		V2078 := __e.Get(2)
		_ = V2078
		V2079 := __e.Get(3)
		_ = V2079
		tmp10239 := PrimVectorGet(V2077, MakeNumber(1))

		__e.Return(PrimVectorSet(V2079, tmp10239, V2078))
		return

	}, 3)

	tmp10240 := Call(__e, ns2_1set, symshen_4bindv, tmp10238)

	_ = tmp10240

	tmp10241 := MakeNative(func(__e *ControlFlow) {
		V2080 := __e.Get(1)
		_ = V2080
		V2081 := __e.Get(2)
		_ = V2081
		V2082 := __e.Get(3)
		_ = V2082
		tmp10242 := PrimVectorGet(V2080, MakeNumber(1))

		tmp10243 := PrimVectorSet(V2081, tmp10242, symshen_4_1null_1)

		_ = tmp10243

		__e.Return(V2082)
		return

	}, 3)

	tmp10244 := Call(__e, ns2_1set, symshen_4unwind, tmp10241)

	_ = tmp10244

	tmp10245 := MakeNative(func(__e *ControlFlow) {
		V2091 := __e.Get(1)
		_ = V2091
		V2092 := __e.Get(2)
		_ = V2092
		V2093 := __e.Get(3)
		_ = V2093
		tmp10260 := PrimEqual(Nil, V2091)

		if True == tmp10260 {
			__e.Return(V2092)
			return
		} else {
			tmp10258 := PrimIsPair(V2091)

			if True == tmp10258 {
				tmp10246 := PrimHead(V2091)

				tmp10247 := PrimCons(V2093, Nil)

				tmp10248 := PrimCons(symshen_4newpv, tmp10247)

				tmp10249 := PrimTail(V2091)

				tmp10250 := Call(__e, PrimFunc(symshen_4stpart), tmp10249, V2092, V2093)

				tmp10251 := PrimCons(tmp10250, Nil)

				tmp10252 := PrimCons(V2093, tmp10251)

				tmp10253 := PrimCons(symshen_4gc, tmp10252)

				tmp10254 := PrimCons(tmp10253, Nil)

				tmp10255 := PrimCons(tmp10248, tmp10254)

				tmp10256 := PrimCons(tmp10246, tmp10255)

				__e.Return(PrimCons(symlet, tmp10256))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
				return
			}

		}

	}, 3)

	tmp10261 := Call(__e, ns2_1set, symshen_4stpart, tmp10245)

	_ = tmp10261

	tmp10262 := MakeNative(func(__e *ControlFlow) {
		V2094 := __e.Get(1)
		_ = V2094
		V2095 := __e.Get(2)
		_ = V2095
		tmp10267 := PrimEqual(V2095, False)

		if True == tmp10267 {
			tmp10263 := MakeNative(func(__e *ControlFlow) {
				W2096 := __e.Get(1)
				_ = W2096
				tmp10264 := Call(__e, PrimFunc(symshen_4decrement_1ticket), W2096, V2094)

				_ = tmp10264

				__e.Return(V2095)
				return

			}, 1)

			tmp10265 := Call(__e, PrimFunc(symshen_4ticket_1number), V2094)

			__e.TailApply(tmp10263, tmp10265)
			return

		} else {
			__e.Return(V2095)
			return
		}

	}, 2)

	tmp10268 := Call(__e, ns2_1set, symshen_4gc, tmp10262)

	_ = tmp10268

	tmp10269 := MakeNative(func(__e *ControlFlow) {
		V2097 := __e.Get(1)
		_ = V2097
		V2098 := __e.Get(2)
		_ = V2098
		tmp10270 := PrimNumberSubtract(V2097, MakeNumber(1))

		__e.Return(PrimVectorSet(V2098, MakeNumber(1), tmp10270))
		return

	}, 2)

	tmp10271 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp10269)

	_ = tmp10271

	tmp10272 := MakeNative(func(__e *ControlFlow) {
		V2099 := __e.Get(1)
		_ = V2099
		tmp10273 := MakeNative(func(__e *ControlFlow) {
			W2100 := __e.Get(1)
			_ = W2100
			tmp10274 := MakeNative(func(__e *ControlFlow) {
				W2101 := __e.Get(1)
				_ = W2101
				tmp10275 := MakeNative(func(__e *ControlFlow) {
					W2102 := __e.Get(1)
					_ = W2102
					__e.Return(W2101)
					return
				}, 1)

				tmp10276 := Call(__e, PrimFunc(symshen_4nextticket), V2099, W2100)

				__e.TailApply(tmp10275, tmp10276)
				return

			}, 1)

			tmp10277 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), W2100)

			__e.TailApply(tmp10274, tmp10277)
			return

		}, 1)

		tmp10278 := Call(__e, PrimFunc(symshen_4ticket_1number), V2099)

		__e.TailApply(tmp10273, tmp10278)
		return

	}, 1)

	tmp10279 := Call(__e, ns2_1set, symshen_4newpv, tmp10272)

	_ = tmp10279

	tmp10280 := MakeNative(func(__e *ControlFlow) {
		V2103 := __e.Get(1)
		_ = V2103
		__e.Return(PrimVectorGet(V2103, MakeNumber(1)))
		return
	}, 1)

	tmp10281 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp10280)

	_ = tmp10281

	tmp10282 := MakeNative(func(__e *ControlFlow) {
		V2104 := __e.Get(1)
		_ = V2104
		V2105 := __e.Get(2)
		_ = V2105
		tmp10283 := MakeNative(func(__e *ControlFlow) {
			W2106 := __e.Get(1)
			_ = W2106
			tmp10284 := PrimNumberAdd(V2105, MakeNumber(1))

			__e.Return(PrimVectorSet(W2106, MakeNumber(1), tmp10284))
			return

		}, 1)

		tmp10285 := PrimVectorSet(V2104, V2105, symshen_4_1null_1)

		__e.TailApply(tmp10283, tmp10285)
		return

	}, 2)

	tmp10286 := Call(__e, ns2_1set, symshen_4nextticket, tmp10282)

	_ = tmp10286

	tmp10287 := MakeNative(func(__e *ControlFlow) {
		V2107 := __e.Get(1)
		_ = V2107
		tmp10288 := PrimAbsvector(MakeNumber(2))

		tmp10289 := PrimVectorSet(tmp10288, MakeNumber(0), symshen_4pvar)

		__e.Return(PrimVectorSet(tmp10289, MakeNumber(1), V2107))
		return

	}, 1)

	tmp10290 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp10287)

	_ = tmp10290

	tmp10291 := MakeNative(func(__e *ControlFlow) {
		V2108 := __e.Get(1)
		_ = V2108
		tmp10292 := PrimVectorGet(V2108, MakeNumber(1))

		tmp10293 := Call(__e, PrimFunc(symshen_4app), tmp10292, MakeString(""), symshen_4a)

		__e.Return(PrimStringConcat(MakeString("Var"), tmp10293))
		return

	}, 1)

	tmp10294 := Call(__e, ns2_1set, symshen_4pvar, tmp10291)

	_ = tmp10294

	tmp10295 := MakeNative(func(__e *ControlFlow) {
		tmp10296 := PrimValue(symshen_4_dinfs_d)

		tmp10297 := PrimNumberAdd(MakeNumber(1), tmp10296)

		__e.Return(PrimSet(symshen_4_dinfs_d, tmp10297))
		return

	}, 0)

	tmp10298 := Call(__e, ns2_1set, symshen_4incinfs, tmp10295)

	_ = tmp10298

	tmp10299 := MakeNative(func(__e *ControlFlow) {
		V2109 := __e.Get(1)
		_ = V2109
		tmp10306 := PrimIsInteger(V2109)

		var ifres10303 Obj

		if True == tmp10306 {
			tmp10305 := PrimGreatThan(V2109, MakeNumber(0))

			var ifres10304 Obj

			if True == tmp10305 {
				ifres10304 = True

			} else {
				ifres10304 = False

			}

			ifres10303 = ifres10304

		} else {
			ifres10303 = False

		}

		if True == ifres10303 {
			__e.Return(PrimSet(symshen_4_dsize_1prolog_1vector_d, V2109))
			return
		} else {
			tmp10300 := Call(__e, PrimFunc(symshen_4app), V2109, MakeString(""), symshen_4a)

			tmp10301 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp10300)

			__e.Return(PrimSimpleError(tmp10301))
			return

		}

	}, 1)

	tmp10307 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp10299)

	_ = tmp10307

	tmp10308 := MakeNative(func(__e *ControlFlow) {
		V2121 := __e.Get(1)
		_ = V2121
		V2122 := __e.Get(2)
		_ = V2122
		V2123 := __e.Get(3)
		_ = V2123
		V2124 := __e.Get(4)
		_ = V2124
		tmp10338 := PrimEqual(V2121, V2122)

		if True == tmp10338 {
			__e.TailApply(PrimFunc(symthaw), V2124)
			return
		} else {
			tmp10336 := Call(__e, PrimFunc(symshen_4pvar_2), V2121)

			var ifres10331 Obj

			if True == tmp10336 {
				tmp10333 := Call(__e, PrimFunc(symshen_4deref), V2122, V2123)

				tmp10334 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2121, tmp10333)

				tmp10335 := PrimNot(tmp10334)

				var ifres10332 Obj

				if True == tmp10335 {
					ifres10332 = True

				} else {
					ifres10332 = False

				}

				ifres10331 = ifres10332

			} else {
				ifres10331 = False

			}

			if True == ifres10331 {
				__e.TailApply(PrimFunc(symshen_4bind_b), V2121, V2122, V2123, V2124)
				return
			} else {
				tmp10329 := Call(__e, PrimFunc(symshen_4pvar_2), V2122)

				var ifres10324 Obj

				if True == tmp10329 {
					tmp10326 := Call(__e, PrimFunc(symshen_4deref), V2121, V2123)

					tmp10327 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2122, tmp10326)

					tmp10328 := PrimNot(tmp10327)

					var ifres10325 Obj

					if True == tmp10328 {
						ifres10325 = True

					} else {
						ifres10325 = False

					}

					ifres10324 = ifres10325

				} else {
					ifres10324 = False

				}

				if True == ifres10324 {
					__e.TailApply(PrimFunc(symshen_4bind_b), V2122, V2121, V2123, V2124)
					return
				} else {
					tmp10322 := PrimIsPair(V2121)

					var ifres10319 Obj

					if True == tmp10322 {
						tmp10321 := PrimIsPair(V2122)

						var ifres10320 Obj

						if True == tmp10321 {
							ifres10320 = True

						} else {
							ifres10320 = False

						}

						ifres10319 = ifres10320

					} else {
						ifres10319 = False

					}

					if True == ifres10319 {
						tmp10309 := PrimHead(V2121)

						tmp10310 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10309, V2123)

						tmp10311 := PrimHead(V2122)

						tmp10312 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10311, V2123)

						tmp10313 := MakeNative(func(__e *ControlFlow) {
							tmp10314 := PrimTail(V2121)

							tmp10315 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10314, V2123)

							tmp10316 := PrimTail(V2122)

							tmp10317 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10316, V2123)

							__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp10315, tmp10317, V2123, V2124)
							return

						}, 0)

						__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp10310, tmp10312, V2123, tmp10313)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp10339 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp10308)

	_ = tmp10339

	tmp10340 := MakeNative(func(__e *ControlFlow) {
		V2136 := __e.Get(1)
		_ = V2136
		V2137 := __e.Get(2)
		_ = V2137
		V2138 := __e.Get(3)
		_ = V2138
		V2139 := __e.Get(4)
		_ = V2139
		tmp10360 := PrimEqual(V2136, V2137)

		if True == tmp10360 {
			__e.TailApply(PrimFunc(symthaw), V2139)
			return
		} else {
			tmp10358 := Call(__e, PrimFunc(symshen_4pvar_2), V2136)

			if True == tmp10358 {
				__e.TailApply(PrimFunc(symshen_4bind_b), V2136, V2137, V2138, V2139)
				return
			} else {
				tmp10356 := Call(__e, PrimFunc(symshen_4pvar_2), V2137)

				if True == tmp10356 {
					__e.TailApply(PrimFunc(symshen_4bind_b), V2137, V2136, V2138, V2139)
					return
				} else {
					tmp10354 := PrimIsPair(V2136)

					var ifres10351 Obj

					if True == tmp10354 {
						tmp10353 := PrimIsPair(V2137)

						var ifres10352 Obj

						if True == tmp10353 {
							ifres10352 = True

						} else {
							ifres10352 = False

						}

						ifres10351 = ifres10352

					} else {
						ifres10351 = False

					}

					if True == ifres10351 {
						tmp10341 := PrimHead(V2136)

						tmp10342 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10341, V2138)

						tmp10343 := PrimHead(V2137)

						tmp10344 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10343, V2138)

						tmp10345 := MakeNative(func(__e *ControlFlow) {
							tmp10346 := PrimTail(V2136)

							tmp10347 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10346, V2138)

							tmp10348 := PrimTail(V2137)

							tmp10349 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10348, V2138)

							__e.TailApply(PrimFunc(symshen_4lzy_a), tmp10347, tmp10349, V2138, V2139)
							return

						}, 0)

						__e.TailApply(PrimFunc(symshen_4lzy_a), tmp10342, tmp10344, V2138, tmp10345)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp10361 := Call(__e, ns2_1set, symshen_4lzy_a, tmp10340)

	_ = tmp10361

	tmp10362 := MakeNative(func(__e *ControlFlow) {
		V2145 := __e.Get(1)
		_ = V2145
		V2146 := __e.Get(2)
		_ = V2146
		tmp10372 := PrimEqual(V2145, V2146)

		if True == tmp10372 {
			__e.Return(True)
			return
		} else {
			tmp10370 := PrimIsPair(V2146)

			if True == tmp10370 {
				tmp10367 := PrimHead(V2146)

				tmp10368 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2145, tmp10367)

				if True == tmp10368 {
					__e.Return(True)
					return
				} else {
					tmp10364 := PrimTail(V2146)

					tmp10365 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2145, tmp10364)

					if True == tmp10365 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp10373 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp10362)

	_ = tmp10373

	tmp10374 := MakeNative(func(__e *ControlFlow) {
		V2147 := __e.Get(1)
		_ = V2147
		V2148 := __e.Get(2)
		_ = V2148
		V2149 := __e.Get(3)
		_ = V2149
		V2150 := __e.Get(4)
		_ = V2150
		V2151 := __e.Get(5)
		_ = V2151
		tmp10375 := Call(__e, V2147, V2148)

		tmp10376 := Call(__e, tmp10375, V2149)

		tmp10377 := Call(__e, tmp10376, V2150)

		__e.TailApply(tmp10377, V2151)
		return

	}, 5)

	tmp10378 := Call(__e, ns2_1set, symcall, tmp10374)

	_ = tmp10378

	tmp10379 := MakeNative(func(__e *ControlFlow) {
		V2158 := __e.Get(1)
		_ = V2158
		V2159 := __e.Get(2)
		_ = V2159
		V2160 := __e.Get(3)
		_ = V2160
		V2161 := __e.Get(4)
		_ = V2161
		V2162 := __e.Get(5)
		_ = V2162
		__e.TailApply(PrimFunc(symshen_4deref), V2158, V2159)
		return
	}, 5)

	tmp10380 := Call(__e, ns2_1set, symreturn, tmp10379)

	_ = tmp10380

	tmp10381 := MakeNative(func(__e *ControlFlow) {
		V2169 := __e.Get(1)
		_ = V2169
		V2170 := __e.Get(2)
		_ = V2170
		V2171 := __e.Get(3)
		_ = V2171
		V2172 := __e.Get(4)
		_ = V2172
		V2173 := __e.Get(5)
		_ = V2173
		if True == V2169 {
			__e.TailApply(PrimFunc(symthaw), V2173)
			return
		} else {
			__e.Return(False)
			return
		}
	}, 5)

	tmp10383 := Call(__e, ns2_1set, symwhen, tmp10381)

	_ = tmp10383

	tmp10384 := MakeNative(func(__e *ControlFlow) {
		V2174 := __e.Get(1)
		_ = V2174
		V2175 := __e.Get(2)
		_ = V2175
		V2176 := __e.Get(3)
		_ = V2176
		V2177 := __e.Get(4)
		_ = V2177
		V2178 := __e.Get(5)
		_ = V2178
		V2179 := __e.Get(6)
		_ = V2179
		tmp10385 := Call(__e, PrimFunc(symshen_4lazyderef), V2174, V2176)

		tmp10386 := Call(__e, PrimFunc(symshen_4lazyderef), V2175, V2176)

		__e.TailApply(PrimFunc(symshen_4lzy_a), tmp10385, tmp10386, V2176, V2179)
		return

	}, 6)

	tmp10387 := Call(__e, ns2_1set, symis, tmp10384)

	_ = tmp10387

	tmp10388 := MakeNative(func(__e *ControlFlow) {
		V2180 := __e.Get(1)
		_ = V2180
		V2181 := __e.Get(2)
		_ = V2181
		V2182 := __e.Get(3)
		_ = V2182
		V2183 := __e.Get(4)
		_ = V2183
		V2184 := __e.Get(5)
		_ = V2184
		V2185 := __e.Get(6)
		_ = V2185
		tmp10389 := Call(__e, PrimFunc(symshen_4lazyderef), V2180, V2182)

		tmp10390 := Call(__e, PrimFunc(symshen_4lazyderef), V2181, V2182)

		__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp10389, tmp10390, V2182, V2185)
		return

	}, 6)

	tmp10391 := Call(__e, ns2_1set, symis_b, tmp10388)

	_ = tmp10391

	tmp10392 := MakeNative(func(__e *ControlFlow) {
		V2190 := __e.Get(1)
		_ = V2190
		V2191 := __e.Get(2)
		_ = V2191
		V2192 := __e.Get(3)
		_ = V2192
		V2193 := __e.Get(4)
		_ = V2193
		V2194 := __e.Get(5)
		_ = V2194
		V2195 := __e.Get(6)
		_ = V2195
		__e.TailApply(PrimFunc(symshen_4bind_b), V2190, V2191, V2192, V2195)
		return
	}, 6)

	tmp10393 := Call(__e, ns2_1set, symbind, tmp10392)

	_ = tmp10393

	tmp10394 := MakeNative(func(__e *ControlFlow) {
		V2196 := __e.Get(1)
		_ = V2196
		V2197 := __e.Get(2)
		_ = V2197
		V2198 := __e.Get(3)
		_ = V2198
		V2199 := __e.Get(4)
		_ = V2199
		V2200 := __e.Get(5)
		_ = V2200
		tmp10396 := Call(__e, PrimFunc(symshen_4lazyderef), V2196, V2197)

		tmp10397 := Call(__e, PrimFunc(symshen_4pvar_2), tmp10396)

		if True == tmp10397 {
			__e.TailApply(PrimFunc(symthaw), V2200)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 5)

	tmp10398 := Call(__e, ns2_1set, symvar_2, tmp10394)

	_ = tmp10398

	tmp10399 := MakeNative(func(__e *ControlFlow) {
		V2203 := __e.Get(1)
		_ = V2203
		__e.Return(MakeString("|prolog vector|"))
		return
	}, 1)

	tmp10400 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp10399)

	_ = tmp10400

	tmp10401 := MakeNative(func(__e *ControlFlow) {
		V2222 := __e.Get(1)
		_ = V2222
		V2223 := __e.Get(2)
		_ = V2223
		V2224 := __e.Get(3)
		_ = V2224
		V2225 := __e.Get(4)
		_ = V2225
		V2226 := __e.Get(5)
		_ = V2226
		tmp10414 := PrimEqual(Nil, V2222)

		if True == tmp10414 {
			__e.Return(False)
			return
		} else {
			tmp10412 := PrimIsPair(V2222)

			if True == tmp10412 {
				tmp10402 := MakeNative(func(__e *ControlFlow) {
					W2227 := __e.Get(1)
					_ = W2227
					tmp10405 := PrimEqual(W2227, False)

					if True == tmp10405 {
						tmp10403 := PrimTail(V2222)

						__e.TailApply(PrimFunc(symfork), tmp10403, V2223, V2224, V2225, V2226)
						return

					} else {
						__e.Return(W2227)
						return
					}

				}, 1)

				tmp10406 := PrimHead(V2222)

				tmp10407 := Call(__e, tmp10406, V2223)

				tmp10408 := Call(__e, tmp10407, V2224)

				tmp10409 := Call(__e, tmp10408, V2225)

				tmp10410 := Call(__e, tmp10409, V2226)

				__e.TailApply(tmp10402, tmp10410)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
				return
			}

		}

	}, 5)

	tmp10415 := Call(__e, ns2_1set, symfork, tmp10401)

	_ = tmp10415

	tmp10416 := MakeNative(func(__e *ControlFlow) {
		V2228 := __e.Get(1)
		_ = V2228
		V2229 := __e.Get(2)
		_ = V2229
		V2230 := __e.Get(3)
		_ = V2230
		V2231 := __e.Get(4)
		_ = V2231
		V2232 := __e.Get(5)
		_ = V2232
		V2233 := __e.Get(6)
		_ = V2233
		V2234 := __e.Get(7)
		_ = V2234
		tmp10423 := Call(__e, PrimFunc(symshen_4unlocked_2), V2232)

		if True == tmp10423 {
			tmp10417 := MakeNative(func(__e *ControlFlow) {
				W2235 := __e.Get(1)
				_ = W2235
				tmp10418 := Call(__e, PrimFunc(symshen_4incinfs))

				_ = tmp10418

				tmp10419 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimFunc(symshen_4findall_1h), V2228, V2229, V2230, W2235, V2231, V2232, V2233, V2234)
					return
				}, 0)

				tmp10420 := Call(__e, PrimFunc(symis), W2235, Nil, V2231, V2232, V2233, tmp10419)

				__e.TailApply(PrimFunc(symshen_4gc), V2231, tmp10420)
				return

			}, 1)

			tmp10421 := Call(__e, PrimFunc(symshen_4newpv), V2231)

			__e.TailApply(tmp10417, tmp10421)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 7)

	tmp10424 := Call(__e, ns2_1set, symfindall, tmp10416)

	_ = tmp10424

	tmp10425 := MakeNative(func(__e *ControlFlow) {
		V2236 := __e.Get(1)
		_ = V2236
		V2237 := __e.Get(2)
		_ = V2237
		V2238 := __e.Get(3)
		_ = V2238
		V2239 := __e.Get(4)
		_ = V2239
		V2240 := __e.Get(5)
		_ = V2240
		V2241 := __e.Get(6)
		_ = V2241
		V2242 := __e.Get(7)
		_ = V2242
		V2243 := __e.Get(8)
		_ = V2243
		tmp10426 := MakeNative(func(__e *ControlFlow) {
			W2244 := __e.Get(1)
			_ = W2244
			tmp10431 := PrimEqual(W2244, False)

			if True == tmp10431 {
				tmp10429 := Call(__e, PrimFunc(symshen_4unlocked_2), V2241)

				if True == tmp10429 {
					tmp10427 := Call(__e, PrimFunc(symshen_4incinfs))

					_ = tmp10427

					__e.TailApply(PrimFunc(symis_b), V2238, V2239, V2240, V2241, V2242, V2243)
					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(W2244)
				return
			}

		}, 1)

		tmp10436 := Call(__e, PrimFunc(symshen_4unlocked_2), V2241)

		var ifres10432 Obj

		if True == tmp10436 {
			tmp10433 := Call(__e, PrimFunc(symshen_4incinfs))

			_ = tmp10433

			tmp10434 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimFunc(symshen_4overbind), V2236, V2239, V2240, V2241, V2242, V2243)
				return
			}, 0)

			tmp10435 := Call(__e, PrimFunc(symcall), V2237, V2240, V2241, V2242, tmp10434)

			ifres10432 = tmp10435

		} else {
			ifres10432 = False

		}

		__e.TailApply(tmp10426, ifres10432)
		return

	}, 8)

	tmp10437 := Call(__e, ns2_1set, symshen_4findall_1h, tmp10425)

	_ = tmp10437

	tmp10438 := MakeNative(func(__e *ControlFlow) {
		V2251 := __e.Get(1)
		_ = V2251
		V2252 := __e.Get(2)
		_ = V2252
		V2253 := __e.Get(3)
		_ = V2253
		V2254 := __e.Get(4)
		_ = V2254
		V2255 := __e.Get(5)
		_ = V2255
		V2256 := __e.Get(6)
		_ = V2256
		tmp10439 := Call(__e, PrimFunc(symshen_4deref), V2251, V2253)

		tmp10440 := Call(__e, PrimFunc(symshen_4lazyderef), V2252, V2253)

		tmp10441 := PrimCons(tmp10439, tmp10440)

		tmp10442 := Call(__e, PrimFunc(symshen_4bindv), V2252, tmp10441, V2253)

		_ = tmp10442

		__e.Return(False)
		return

	}, 6)

	tmp10443 := Call(__e, ns2_1set, symshen_4overbind, tmp10438)

	_ = tmp10443

	tmp10444 := MakeNative(func(__e *ControlFlow) {
		V2259 := __e.Get(1)
		_ = V2259
		tmp10448 := PrimEqual(sym_7, V2259)

		if True == tmp10448 {
			__e.Return(PrimSet(symshen_4_doccurs_d, True))
			return
		} else {
			tmp10446 := PrimEqual(sym_1, V2259)

			if True == tmp10446 {
				__e.Return(PrimSet(symshen_4_doccurs_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	__e.TailApply(ns2_1set, symoccurs_1check, tmp10444)
	return

}, 0)
