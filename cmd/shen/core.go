package main

import . "github.com/tiancaiamao/shen-go/cora"

var CoreMain = MakeNative(func(__e *ControlFlow) {
	tmp1470 := MakeNative(func(__e *ControlFlow) {
		V1447 := __e.Get(1)
		_ = V1447
		tmp1471 := MakeNative(func(__e *ControlFlow) {
			KL := __e.Get(1)
			_ = KL
			__e.TailApply(PrimNS2Value(symshen_4record_1and_1evaluate), KL)
			return
		}, 1)

		tmp1472 := Call(__e, PrimNS2Value(symshen_4shen_1_6kl_1h), V1447)

		__e.TailApply(tmp1471, tmp1472)
		return

	}, 1)

	tmp1473 := Call(__e, PrimNS2Value(symdef), symshen_4shen_1_6kl, tmp1470)

	_ = tmp1473

	tmp1474 := MakeNative(func(__e *ControlFlow) {
		V1448 := __e.Get(1)
		_ = V1448
		tmp1527 := PrimIsPair(V1448)

		var ifres1501 Obj

		if True == tmp1527 {
			tmp1525 := PrimHead(V1448)

			tmp1526 := PrimEqual(symdefun, tmp1525)

			var ifres1503 Obj

			if True == tmp1526 {
				tmp1523 := PrimTail(V1448)

				tmp1524 := PrimIsPair(tmp1523)

				var ifres1505 Obj

				if True == tmp1524 {
					tmp1520 := PrimTail(V1448)

					tmp1521 := PrimTail(tmp1520)

					tmp1522 := PrimIsPair(tmp1521)

					var ifres1507 Obj

					if True == tmp1522 {
						tmp1516 := PrimTail(V1448)

						tmp1517 := PrimTail(tmp1516)

						tmp1518 := PrimTail(tmp1517)

						tmp1519 := PrimIsPair(tmp1518)

						var ifres1509 Obj

						if True == tmp1519 {
							tmp1511 := PrimTail(V1448)

							tmp1512 := PrimTail(tmp1511)

							tmp1513 := PrimTail(tmp1512)

							tmp1514 := PrimTail(tmp1513)

							tmp1515 := PrimEqual(Nil, tmp1514)

							var ifres1510 Obj

							if True == tmp1515 {
								ifres1510 = True

							} else {
								ifres1510 = False

							}

							ifres1509 = ifres1510

						} else {
							ifres1509 = False

						}

						var ifres1508 Obj

						if True == ifres1509 {
							ifres1508 = True

						} else {
							ifres1508 = False

						}

						ifres1507 = ifres1508

					} else {
						ifres1507 = False

					}

					var ifres1506 Obj

					if True == ifres1507 {
						ifres1506 = True

					} else {
						ifres1506 = False

					}

					ifres1505 = ifres1506

				} else {
					ifres1505 = False

				}

				var ifres1504 Obj

				if True == ifres1505 {
					ifres1504 = True

				} else {
					ifres1504 = False

				}

				ifres1503 = ifres1504

			} else {
				ifres1503 = False

			}

			var ifres1502 Obj

			if True == ifres1503 {
				ifres1502 = True

			} else {
				ifres1502 = False

			}

			ifres1501 = ifres1502

		} else {
			ifres1501 = False

		}

		if True == ifres1501 {
			tmp1476 := MakeNative(func(__e *ControlFlow) {
				SysfuncChk := __e.Get(1)
				_ = SysfuncChk
				tmp1477 := MakeNative(func(__e *ControlFlow) {
					Arity := __e.Get(1)
					_ = Arity
					tmp1478 := MakeNative(func(__e *ControlFlow) {
						Record := __e.Get(1)
						_ = Record
						tmp1479 := MakeNative(func(__e *ControlFlow) {
							Eval := __e.Get(1)
							_ = Eval
							tmp1480 := PrimTail(V1448)

							tmp1481 := PrimHead(tmp1480)

							__e.TailApply(PrimNS2Value(symshen_4fn_1print), tmp1481)
							return

						}, 1)

						tmp1482 := Call(__e, PrimNS2Value(symeval_1kl), V1448)

						__e.TailApply(tmp1479, tmp1482)
						return

					}, 1)

					tmp1483 := PrimTail(V1448)

					tmp1484 := PrimHead(tmp1483)

					tmp1485 := Call(__e, PrimNS2Value(symshen_4record_1kl), tmp1484, V1448)

					__e.TailApply(tmp1478, tmp1485)
					return

				}, 1)

				tmp1486 := PrimTail(V1448)

				tmp1487 := PrimHead(tmp1486)

				tmp1488 := PrimTail(V1448)

				tmp1489 := PrimTail(tmp1488)

				tmp1490 := PrimHead(tmp1489)

				tmp1491 := Call(__e, PrimNS2Value(symlength), tmp1490)

				tmp1492 := Call(__e, PrimNS2Value(symshen_4store_1arity), tmp1487, tmp1491)

				__e.TailApply(tmp1477, tmp1492)
				return

			}, 1)

			tmp1498 := PrimTail(V1448)

			tmp1499 := PrimHead(tmp1498)

			tmp1500 := Call(__e, PrimNS2Value(symshen_4sysfunc_2), tmp1499)

			var ifres1493 Obj

			if True == tmp1500 {
				tmp1494 := PrimTail(V1448)

				tmp1495 := PrimHead(tmp1494)

				tmp1496 := Call(__e, PrimNS2Value(symshen_4app), tmp1495, MakeString(" is not a legitimate function name\n"), symshen_4a)

				tmp1497 := PrimSimpleError(tmp1496)

				ifres1493 = tmp1497

			} else {
				ifres1493 = symshen_4skip

			}

			__e.TailApply(tmp1476, ifres1493)
			return

		} else {
			__e.Return(V1448)
			return
		}

	}, 1)

	tmp1528 := Call(__e, PrimNS2Value(symdef), symshen_4record_1and_1evaluate, tmp1474)

	_ = tmp1528

	tmp1529 := MakeNative(func(__e *ControlFlow) {
		V1449 := __e.Get(1)
		_ = V1449
		tmp1630 := PrimIsPair(V1449)

		var ifres1622 Obj

		if True == tmp1630 {
			tmp1628 := PrimHead(V1449)

			tmp1629 := PrimEqual(symdefine, tmp1628)

			var ifres1624 Obj

			if True == tmp1629 {
				tmp1626 := PrimTail(V1449)

				tmp1627 := PrimIsPair(tmp1626)

				var ifres1625 Obj

				if True == tmp1627 {
					ifres1625 = True

				} else {
					ifres1625 = False

				}

				ifres1624 = ifres1625

			} else {
				ifres1624 = False

			}

			var ifres1623 Obj

			if True == ifres1624 {
				ifres1623 = True

			} else {
				ifres1623 = False

			}

			ifres1622 = ifres1623

		} else {
			ifres1622 = False

		}

		if True == ifres1622 {
			tmp1618 := PrimTail(V1449)

			tmp1619 := PrimHead(tmp1618)

			tmp1620 := PrimTail(V1449)

			tmp1621 := PrimTail(tmp1620)

			__e.TailApply(PrimNS2Value(symshen_4shendef_1_6kldef), tmp1619, tmp1621)
			return

		} else {
			tmp1617 := PrimIsPair(V1449)

			var ifres1591 Obj

			if True == tmp1617 {
				tmp1615 := PrimHead(V1449)

				tmp1616 := PrimEqual(symdefun, tmp1615)

				var ifres1593 Obj

				if True == tmp1616 {
					tmp1613 := PrimTail(V1449)

					tmp1614 := PrimIsPair(tmp1613)

					var ifres1595 Obj

					if True == tmp1614 {
						tmp1610 := PrimTail(V1449)

						tmp1611 := PrimTail(tmp1610)

						tmp1612 := PrimIsPair(tmp1611)

						var ifres1597 Obj

						if True == tmp1612 {
							tmp1606 := PrimTail(V1449)

							tmp1607 := PrimTail(tmp1606)

							tmp1608 := PrimTail(tmp1607)

							tmp1609 := PrimIsPair(tmp1608)

							var ifres1599 Obj

							if True == tmp1609 {
								tmp1601 := PrimTail(V1449)

								tmp1602 := PrimTail(tmp1601)

								tmp1603 := PrimTail(tmp1602)

								tmp1604 := PrimTail(tmp1603)

								tmp1605 := PrimEqual(Nil, tmp1604)

								var ifres1600 Obj

								if True == tmp1605 {
									ifres1600 = True

								} else {
									ifres1600 = False

								}

								ifres1599 = ifres1600

							} else {
								ifres1599 = False

							}

							var ifres1598 Obj

							if True == ifres1599 {
								ifres1598 = True

							} else {
								ifres1598 = False

							}

							ifres1597 = ifres1598

						} else {
							ifres1597 = False

						}

						var ifres1596 Obj

						if True == ifres1597 {
							ifres1596 = True

						} else {
							ifres1596 = False

						}

						ifres1595 = ifres1596

					} else {
						ifres1595 = False

					}

					var ifres1594 Obj

					if True == ifres1595 {
						ifres1594 = True

					} else {
						ifres1594 = False

					}

					ifres1593 = ifres1594

				} else {
					ifres1593 = False

				}

				var ifres1592 Obj

				if True == ifres1593 {
					ifres1592 = True

				} else {
					ifres1592 = False

				}

				ifres1591 = ifres1592

			} else {
				ifres1591 = False

			}

			if True == ifres1591 {
				__e.Return(V1449)
				return
			} else {
				tmp1590 := PrimIsPair(V1449)

				var ifres1571 Obj

				if True == tmp1590 {
					tmp1588 := PrimHead(V1449)

					tmp1589 := PrimEqual(symtype, tmp1588)

					var ifres1573 Obj

					if True == tmp1589 {
						tmp1586 := PrimTail(V1449)

						tmp1587 := PrimIsPair(tmp1586)

						var ifres1575 Obj

						if True == tmp1587 {
							tmp1583 := PrimTail(V1449)

							tmp1584 := PrimTail(tmp1583)

							tmp1585 := PrimIsPair(tmp1584)

							var ifres1577 Obj

							if True == tmp1585 {
								tmp1579 := PrimTail(V1449)

								tmp1580 := PrimTail(tmp1579)

								tmp1581 := PrimTail(tmp1580)

								tmp1582 := PrimEqual(Nil, tmp1581)

								var ifres1578 Obj

								if True == tmp1582 {
									ifres1578 = True

								} else {
									ifres1578 = False

								}

								ifres1577 = ifres1578

							} else {
								ifres1577 = False

							}

							var ifres1576 Obj

							if True == ifres1577 {
								ifres1576 = True

							} else {
								ifres1576 = False

							}

							ifres1575 = ifres1576

						} else {
							ifres1575 = False

						}

						var ifres1574 Obj

						if True == ifres1575 {
							ifres1574 = True

						} else {
							ifres1574 = False

						}

						ifres1573 = ifres1574

					} else {
						ifres1573 = False

					}

					var ifres1572 Obj

					if True == ifres1573 {
						ifres1572 = True

					} else {
						ifres1572 = False

					}

					ifres1571 = ifres1572

				} else {
					ifres1571 = False

				}

				if True == ifres1571 {
					tmp1563 := PrimTail(V1449)

					tmp1564 := PrimHead(tmp1563)

					tmp1565 := PrimTail(V1449)

					tmp1566 := PrimTail(tmp1565)

					tmp1567 := PrimHead(tmp1566)

					tmp1568 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp1567)

					tmp1569 := PrimCons(tmp1568, Nil)

					tmp1570 := PrimCons(tmp1564, tmp1569)

					__e.Return(PrimCons(symtype, tmp1570))
					return

				} else {
					tmp1562 := PrimIsPair(V1449)

					var ifres1543 Obj

					if True == tmp1562 {
						tmp1560 := PrimHead(V1449)

						tmp1561 := PrimEqual(syminput_7, tmp1560)

						var ifres1545 Obj

						if True == tmp1561 {
							tmp1558 := PrimTail(V1449)

							tmp1559 := PrimIsPair(tmp1558)

							var ifres1547 Obj

							if True == tmp1559 {
								tmp1555 := PrimTail(V1449)

								tmp1556 := PrimTail(tmp1555)

								tmp1557 := PrimIsPair(tmp1556)

								var ifres1549 Obj

								if True == tmp1557 {
									tmp1551 := PrimTail(V1449)

									tmp1552 := PrimTail(tmp1551)

									tmp1553 := PrimTail(tmp1552)

									tmp1554 := PrimEqual(Nil, tmp1553)

									var ifres1550 Obj

									if True == tmp1554 {
										ifres1550 = True

									} else {
										ifres1550 = False

									}

									ifres1549 = ifres1550

								} else {
									ifres1549 = False

								}

								var ifres1548 Obj

								if True == ifres1549 {
									ifres1548 = True

								} else {
									ifres1548 = False

								}

								ifres1547 = ifres1548

							} else {
								ifres1547 = False

							}

							var ifres1546 Obj

							if True == ifres1547 {
								ifres1546 = True

							} else {
								ifres1546 = False

							}

							ifres1545 = ifres1546

						} else {
							ifres1545 = False

						}

						var ifres1544 Obj

						if True == ifres1545 {
							ifres1544 = True

						} else {
							ifres1544 = False

						}

						ifres1543 = ifres1544

					} else {
						ifres1543 = False

					}

					if True == ifres1543 {
						tmp1537 := PrimTail(V1449)

						tmp1538 := PrimHead(tmp1537)

						tmp1539 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp1538)

						tmp1540 := PrimTail(V1449)

						tmp1541 := PrimTail(tmp1540)

						tmp1542 := PrimCons(tmp1539, tmp1541)

						__e.Return(PrimCons(syminput_7, tmp1542))
						return

					} else {
						tmp1536 := PrimIsPair(V1449)

						if True == tmp1536 {
							tmp1535 := MakeNative(func(__e *ControlFlow) {
								Z := __e.Get(1)
								_ = Z
								__e.TailApply(PrimNS2Value(symshen_4shen_1_6kl_1h), Z)
								return
							}, 1)

							__e.TailApply(PrimNS2Value(symmap), tmp1535, V1449)
							return

						} else {
							__e.Return(V1449)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp1631 := Call(__e, PrimNS2Value(symdef), symshen_4shen_1_6kl_1h, tmp1529)

	_ = tmp1631

	tmp1632 := MakeNative(func(__e *ControlFlow) {
		V1450 := __e.Get(1)
		_ = V1450
		V1451 := __e.Get(2)
		_ = V1451
		tmp1633 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_5define_6), X)
			return
		}, 1)

		tmp1634 := PrimCons(V1450, V1451)

		__e.TailApply(PrimNS2Value(symcompile), tmp1633, tmp1634)
		return

	}, 2)

	tmp1635 := Call(__e, PrimNS2Value(symdef), symshen_4shendef_1_6kldef, tmp1632)

	_ = tmp1635

	tmp1636 := MakeNative(func(__e *ControlFlow) {
		V1452 := __e.Get(1)
		_ = V1452
		tmp1637 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1655 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp1655 {
				tmp1639 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp1641 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp1641 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp1642 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5name_6 := __e.Get(1)
					_ = Parseshen_4_5name_6
					tmp1652 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5name_6)

					if True == tmp1652 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1644 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5rules_6 := __e.Get(1)
							_ = Parseshen_4_5rules_6
							tmp1650 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rules_6)

							if True == tmp1650 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp1646 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5rules_6)

								tmp1647 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5name_6)

								tmp1648 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rules_6)

								tmp1649 := Call(__e, PrimNS2Value(symshen_4shendef_1_6kldef_1h), tmp1647, tmp1648)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp1646, tmp1649)
								return

							}

						}, 1)

						tmp1651 := Call(__e, PrimNS2Value(symshen_4_5rules_6), Parseshen_4_5name_6)

						__e.TailApply(tmp1644, tmp1651)
						return

					}

				}, 1)

				tmp1653 := Call(__e, PrimNS2Value(symshen_4_5name_6), V1452)

				tmp1654 := Call(__e, tmp1642, tmp1653)

				__e.TailApply(tmp1639, tmp1654)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp1656 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5name_6 := __e.Get(1)
			_ = Parseshen_4_5name_6
			tmp1678 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5name_6)

			if True == tmp1678 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp1677 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5name_6, sym_i)

				if True == tmp1677 {
					tmp1659 := MakeNative(func(__e *ControlFlow) {
						News1313 := __e.Get(1)
						_ = News1313
						tmp1660 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5signature_6 := __e.Get(1)
							_ = Parseshen_4_5signature_6
							tmp1674 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5signature_6)

							if True == tmp1674 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp1673 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5signature_6, sym_j)

								if True == tmp1673 {
									tmp1663 := MakeNative(func(__e *ControlFlow) {
										News1314 := __e.Get(1)
										_ = News1314
										tmp1664 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5rules_6 := __e.Get(1)
											_ = Parseshen_4_5rules_6
											tmp1670 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rules_6)

											if True == tmp1670 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp1666 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5rules_6)

												tmp1667 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5name_6)

												tmp1668 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rules_6)

												tmp1669 := Call(__e, PrimNS2Value(symshen_4shendef_1_6kldef_1h), tmp1667, tmp1668)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp1666, tmp1669)
												return

											}

										}, 1)

										tmp1671 := Call(__e, PrimNS2Value(symshen_4_5rules_6), News1314)

										__e.TailApply(tmp1664, tmp1671)
										return

									}, 1)

									tmp1672 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5signature_6)

									__e.TailApply(tmp1663, tmp1672)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}

						}, 1)

						tmp1675 := Call(__e, PrimNS2Value(symshen_4_5signature_6), News1313)

						__e.TailApply(tmp1660, tmp1675)
						return

					}, 1)

					tmp1676 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5name_6)

					__e.TailApply(tmp1659, tmp1676)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp1679 := Call(__e, PrimNS2Value(symshen_4_5name_6), V1452)

		tmp1680 := Call(__e, tmp1656, tmp1679)

		__e.TailApply(tmp1637, tmp1680)
		return

	}, 1)

	tmp1681 := Call(__e, PrimNS2Value(symdef), symshen_4_5define_6, tmp1636)

	_ = tmp1681

	tmp1682 := MakeNative(func(__e *ControlFlow) {
		V1453 := __e.Get(1)
		_ = V1453
		V1454 := __e.Get(2)
		_ = V1454
		tmp1683 := MakeNative(func(__e *ControlFlow) {
			Ps := __e.Get(1)
			_ = Ps
			tmp1684 := MakeNative(func(__e *ControlFlow) {
				Arity := __e.Get(1)
				_ = Arity
				tmp1685 := MakeNative(func(__e *ControlFlow) {
					FreeVarChk := __e.Get(1)
					_ = FreeVarChk
					tmp1686 := MakeNative(func(__e *ControlFlow) {
						KL := __e.Get(1)
						_ = KL
						__e.Return(KL)
						return
					}, 1)

					tmp1687 := Call(__e, PrimNS2Value(symshen_4compile_1to_1kl), V1453, V1454, Arity)

					tmp1688 := Call(__e, PrimNS2Value(symshen_4factorise_1code), tmp1687)

					__e.TailApply(tmp1686, tmp1688)
					return

				}, 1)

				tmp1689 := MakeNative(func(__e *ControlFlow) {
					R := __e.Get(1)
					_ = R
					__e.TailApply(PrimNS2Value(symshen_4free_1var_1chk), V1453, R)
					return
				}, 1)

				tmp1690 := Call(__e, PrimNS2Value(symmap), tmp1689, V1454)

				__e.TailApply(tmp1685, tmp1690)
				return

			}, 1)

			tmp1691 := Call(__e, PrimNS2Value(symshen_4arity_1chk), V1453, Ps)

			__e.TailApply(tmp1684, tmp1691)
			return

		}, 1)

		tmp1692 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symfst), X)
			return
		}, 1)

		tmp1693 := Call(__e, PrimNS2Value(symmap), tmp1692, V1454)

		__e.TailApply(tmp1683, tmp1693)
		return

	}, 2)

	tmp1694 := Call(__e, PrimNS2Value(symdef), symshen_4shendef_1_6kldef_1h, tmp1682)

	_ = tmp1694

	tmp1695 := MakeNative(func(__e *ControlFlow) {
		V1455 := __e.Get(1)
		_ = V1455
		tmp1696 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1698 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp1698 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp1715 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1455)

		var ifres1699 Obj

		if True == tmp1715 {
			tmp1701 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp1702 := MakeNative(func(__e *ControlFlow) {
					News1316 := __e.Get(1)
					_ = News1316
					tmp1703 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1316)

					tmp1711 := PrimIsSymbol(X)

					var ifres1707 Obj

					if True == tmp1711 {
						tmp1709 := PrimIsVariable(X)

						tmp1710 := PrimNot(tmp1709)

						var ifres1708 Obj

						if True == tmp1710 {
							ifres1708 = True

						} else {
							ifres1708 = False

						}

						ifres1707 = ifres1708

					} else {
						ifres1707 = False

					}

					var ifres1704 Obj

					if True == ifres1707 {
						ifres1704 = X

					} else {
						tmp1705 := Call(__e, PrimNS2Value(symshen_4app), X, MakeString(" is not a legitimate function name.\n"), symshen_4a)

						tmp1706 := PrimSimpleError(tmp1705)

						ifres1704 = tmp1706

					}

					__e.TailApply(PrimNS2Value(symshen_4comb), tmp1703, ifres1704)
					return

				}, 1)

				tmp1712 := Call(__e, PrimNS2Value(symshen_4tls), V1455)

				__e.TailApply(tmp1702, tmp1712)
				return

			}, 1)

			tmp1713 := Call(__e, PrimNS2Value(symshen_4hds), V1455)

			tmp1714 := Call(__e, tmp1701, tmp1713)

			ifres1699 = tmp1714

		} else {
			tmp1700 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres1699 = tmp1700

		}

		__e.TailApply(tmp1696, ifres1699)
		return

	}, 1)

	tmp1716 := Call(__e, PrimNS2Value(symdef), symshen_4_5name_6, tmp1695)

	_ = tmp1716

	tmp1717 := MakeNative(func(__e *ControlFlow) {
		V1456 := __e.Get(1)
		_ = V1456
		tmp1718 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1729 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp1729 {
				tmp1720 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp1722 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp1722 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp1723 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp1726 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp1726 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1725 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp1725, Nil)
						return

					}

				}, 1)

				tmp1727 := Call(__e, PrimNS2Value(sym_5e_6), V1456)

				tmp1728 := Call(__e, tmp1723, tmp1727)

				__e.TailApply(tmp1720, tmp1728)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp1749 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1456)

		var ifres1730 Obj

		if True == tmp1749 {
			tmp1732 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp1733 := MakeNative(func(__e *ControlFlow) {
					News1318 := __e.Get(1)
					_ = News1318
					tmp1734 := MakeNative(func(__e *ControlFlow) {
						Parseshen_4_5signature_6 := __e.Get(1)
						_ = Parseshen_4_5signature_6
						tmp1744 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5signature_6)

						if True == tmp1744 {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						} else {
							tmp1740 := PrimCons(sym_j, Nil)

							tmp1741 := PrimCons(sym_i, tmp1740)

							tmp1742 := Call(__e, PrimNS2Value(symelement_2), X, tmp1741)

							tmp1743 := PrimNot(tmp1742)

							if True == tmp1743 {
								tmp1737 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5signature_6)

								tmp1738 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5signature_6)

								tmp1739 := PrimCons(X, tmp1738)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp1737, tmp1739)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							}

						}

					}, 1)

					tmp1745 := Call(__e, PrimNS2Value(symshen_4_5signature_6), News1318)

					__e.TailApply(tmp1734, tmp1745)
					return

				}, 1)

				tmp1746 := Call(__e, PrimNS2Value(symshen_4tls), V1456)

				__e.TailApply(tmp1733, tmp1746)
				return

			}, 1)

			tmp1747 := Call(__e, PrimNS2Value(symshen_4hds), V1456)

			tmp1748 := Call(__e, tmp1732, tmp1747)

			ifres1730 = tmp1748

		} else {
			tmp1731 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres1730 = tmp1731

		}

		__e.TailApply(tmp1718, ifres1730)
		return

	}, 1)

	tmp1750 := Call(__e, PrimNS2Value(symdef), symshen_4_5signature_6, tmp1717)

	_ = tmp1750

	tmp1751 := MakeNative(func(__e *ControlFlow) {
		V1457 := __e.Get(1)
		_ = V1457
		tmp1752 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1770 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp1770 {
				tmp1754 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp1756 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp1756 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp1757 := MakeNative(func(__e *ControlFlow) {
					Parse_5_b_6 := __e.Get(1)
					_ = Parse_5_b_6
					tmp1767 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5_b_6)

					if True == tmp1767 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1759 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5_b_6)

						tmp1765 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

						tmp1766 := Call(__e, PrimNS2Value(symempty_2), tmp1765)

						var ifres1760 Obj

						if True == tmp1766 {
							ifres1760 = Nil

						} else {
							tmp1761 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parse_5_b_6)

							tmp1762 := Call(__e, PrimNS2Value(symshen_4app), tmp1761, MakeString("\n ..."), symshen_4r)

							tmp1763 := PrimStringConcat(MakeString("Shen syntax error here:\n "), tmp1762)

							tmp1764 := PrimSimpleError(tmp1763)

							ifres1760 = tmp1764

						}

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp1759, ifres1760)
						return

					}

				}, 1)

				tmp1768 := Call(__e, PrimNS2Value(sym_5_b_6), V1457)

				tmp1769 := Call(__e, tmp1757, tmp1768)

				__e.TailApply(tmp1754, tmp1769)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp1771 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5rule_6 := __e.Get(1)
			_ = Parseshen_4_5rule_6
			tmp1782 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rule_6)

			if True == tmp1782 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp1773 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5rules_6 := __e.Get(1)
					_ = Parseshen_4_5rules_6
					tmp1780 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rules_6)

					if True == tmp1780 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1775 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5rules_6)

						tmp1776 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rule_6)

						tmp1777 := Call(__e, PrimNS2Value(symshen_4linearise), tmp1776)

						tmp1778 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5rules_6)

						tmp1779 := PrimCons(tmp1777, tmp1778)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp1775, tmp1779)
						return

					}

				}, 1)

				tmp1781 := Call(__e, PrimNS2Value(symshen_4_5rules_6), Parseshen_4_5rule_6)

				__e.TailApply(tmp1773, tmp1781)
				return

			}

		}, 1)

		tmp1783 := Call(__e, PrimNS2Value(symshen_4_5rule_6), V1457)

		tmp1784 := Call(__e, tmp1771, tmp1783)

		__e.TailApply(tmp1752, tmp1784)
		return

	}, 1)

	tmp1785 := Call(__e, PrimNS2Value(symdef), symshen_4_5rules_6, tmp1751)

	_ = tmp1785

	tmp1786 := MakeNative(func(__e *ControlFlow) {
		V1460 := __e.Get(1)
		_ = V1460
		tmp1791 := Call(__e, PrimNS2Value(symtuple_2), V1460)

		if True == tmp1791 {
			tmp1788 := Call(__e, PrimNS2Value(symfst), V1460)

			tmp1789 := Call(__e, PrimNS2Value(symfst), V1460)

			tmp1790 := Call(__e, PrimNS2Value(symsnd), V1460)

			__e.TailApply(PrimNS2Value(symshen_4linearise_1h), tmp1788, tmp1789, Nil, tmp1790)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise")))
			return
		}

	}, 1)

	tmp1792 := Call(__e, PrimNS2Value(symdef), symshen_4linearise, tmp1786)

	_ = tmp1792

	tmp1793 := MakeNative(func(__e *ControlFlow) {
		V1473 := __e.Get(1)
		_ = V1473
		V1474 := __e.Get(2)
		_ = V1474
		V1475 := __e.Get(3)
		_ = V1475
		V1476 := __e.Get(4)
		_ = V1476
		tmp1832 := PrimEqual(Nil, V1473)

		if True == tmp1832 {
			__e.TailApply(PrimNS2Value(sym_8p), V1474, V1476)
			return
		} else {
			tmp1831 := PrimIsPair(V1473)

			var ifres1827 Obj

			if True == tmp1831 {
				tmp1829 := PrimHead(V1473)

				tmp1830 := PrimIsPair(tmp1829)

				var ifres1828 Obj

				if True == tmp1830 {
					ifres1828 = True

				} else {
					ifres1828 = False

				}

				ifres1827 = ifres1828

			} else {
				ifres1827 = False

			}

			if True == ifres1827 {
				tmp1824 := PrimHead(V1473)

				tmp1825 := PrimTail(V1473)

				tmp1826 := Call(__e, PrimNS2Value(symappend), tmp1824, tmp1825)

				__e.TailApply(PrimNS2Value(symshen_4linearise_1h), tmp1826, V1474, V1475, V1476)
				return

			} else {
				tmp1823 := PrimIsPair(V1473)

				var ifres1819 Obj

				if True == tmp1823 {
					tmp1821 := PrimHead(V1473)

					tmp1822 := PrimIsVariable(tmp1821)

					var ifres1820 Obj

					if True == tmp1822 {
						ifres1820 = True

					} else {
						ifres1820 = False

					}

					ifres1819 = ifres1820

				} else {
					ifres1819 = False

				}

				if True == ifres1819 {
					tmp1817 := PrimHead(V1473)

					tmp1818 := Call(__e, PrimNS2Value(symelement_2), tmp1817, V1475)

					if True == tmp1818 {
						tmp1804 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							tmp1805 := PrimTail(V1473)

							tmp1806 := PrimHead(V1473)

							tmp1807 := Call(__e, PrimNS2Value(symshen_4rep_1X), tmp1806, Z, V1474)

							tmp1808 := PrimHead(V1473)

							tmp1809 := PrimCons(tmp1808, Nil)

							tmp1810 := PrimCons(Z, tmp1809)

							tmp1811 := PrimCons(sym_a, tmp1810)

							tmp1812 := PrimCons(V1476, Nil)

							tmp1813 := PrimCons(tmp1811, tmp1812)

							tmp1814 := PrimCons(symwhere, tmp1813)

							__e.TailApply(PrimNS2Value(symshen_4linearise_1h), tmp1805, tmp1807, V1475, tmp1814)
							return

						}, 1)

						tmp1815 := Call(__e, PrimNS2Value(symprotect), symV)

						tmp1816 := Call(__e, PrimNS2Value(symgensym), tmp1815)

						__e.TailApply(tmp1804, tmp1816)
						return

					} else {
						tmp1801 := PrimTail(V1473)

						tmp1802 := PrimHead(V1473)

						tmp1803 := PrimCons(tmp1802, V1475)

						__e.TailApply(PrimNS2Value(symshen_4linearise_1h), tmp1801, V1474, tmp1803, V1476)
						return

					}

				} else {
					tmp1799 := PrimIsPair(V1473)

					if True == tmp1799 {
						tmp1798 := PrimTail(V1473)

						__e.TailApply(PrimNS2Value(symshen_4linearise_1h), tmp1798, V1474, V1475, V1476)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise-h")))
						return
					}

				}

			}

		}

	}, 4)

	tmp1833 := Call(__e, PrimNS2Value(symdef), symshen_4linearise_1h, tmp1793)

	_ = tmp1833

	tmp1834 := MakeNative(func(__e *ControlFlow) {
		V1477 := __e.Get(1)
		_ = V1477
		tmp1835 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1917 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp1917 {
				tmp1837 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp1898 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp1898 {
						tmp1839 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp1864 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp1864 {
								tmp1841 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp1843 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

									if True == tmp1843 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp1844 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5patterns_6 := __e.Get(1)
									_ = Parseshen_4_5patterns_6
									tmp1861 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

									if True == tmp1861 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp1860 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)

										if True == tmp1860 {
											tmp1847 := MakeNative(func(__e *ControlFlow) {
												News1331 := __e.Get(1)
												_ = News1331
												tmp1858 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1331)

												if True == tmp1858 {
													tmp1849 := MakeNative(func(__e *ControlFlow) {
														Action := __e.Get(1)
														_ = Action
														tmp1850 := MakeNative(func(__e *ControlFlow) {
															News1332 := __e.Get(1)
															_ = News1332
															tmp1851 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1332)

															tmp1852 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

															tmp1853 := PrimCons(Action, Nil)

															tmp1854 := PrimCons(symshen_4choicepoint_b, tmp1853)

															tmp1855 := Call(__e, PrimNS2Value(sym_8p), tmp1852, tmp1854)

															__e.TailApply(PrimNS2Value(symshen_4comb), tmp1851, tmp1855)
															return

														}, 1)

														tmp1856 := Call(__e, PrimNS2Value(symshen_4tls), News1331)

														__e.TailApply(tmp1850, tmp1856)
														return

													}, 1)

													tmp1857 := Call(__e, PrimNS2Value(symshen_4hds), News1331)

													__e.TailApply(tmp1849, tmp1857)
													return

												} else {
													__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp1859 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

											__e.TailApply(tmp1847, tmp1859)
											return

										} else {
											__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
											return
										}

									}

								}, 1)

								tmp1862 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V1477)

								tmp1863 := Call(__e, tmp1844, tmp1862)

								__e.TailApply(tmp1841, tmp1863)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp1865 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5patterns_6 := __e.Get(1)
							_ = Parseshen_4_5patterns_6
							tmp1895 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

							if True == tmp1895 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp1894 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)

								if True == tmp1894 {
									tmp1868 := MakeNative(func(__e *ControlFlow) {
										News1327 := __e.Get(1)
										_ = News1327
										tmp1892 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1327)

										if True == tmp1892 {
											tmp1870 := MakeNative(func(__e *ControlFlow) {
												Action := __e.Get(1)
												_ = Action
												tmp1871 := MakeNative(func(__e *ControlFlow) {
													News1328 := __e.Get(1)
													_ = News1328
													tmp1889 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News1328, symwhere)

													if True == tmp1889 {
														tmp1873 := MakeNative(func(__e *ControlFlow) {
															News1329 := __e.Get(1)
															_ = News1329
															tmp1887 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1329)

															if True == tmp1887 {
																tmp1875 := MakeNative(func(__e *ControlFlow) {
																	Guard := __e.Get(1)
																	_ = Guard
																	tmp1876 := MakeNative(func(__e *ControlFlow) {
																		News1330 := __e.Get(1)
																		_ = News1330
																		tmp1877 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1330)

																		tmp1878 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

																		tmp1879 := PrimCons(Action, Nil)

																		tmp1880 := PrimCons(symshen_4choicepoint_b, tmp1879)

																		tmp1881 := PrimCons(tmp1880, Nil)

																		tmp1882 := PrimCons(Guard, tmp1881)

																		tmp1883 := PrimCons(symwhere, tmp1882)

																		tmp1884 := Call(__e, PrimNS2Value(sym_8p), tmp1878, tmp1883)

																		__e.TailApply(PrimNS2Value(symshen_4comb), tmp1877, tmp1884)
																		return

																	}, 1)

																	tmp1885 := Call(__e, PrimNS2Value(symshen_4tls), News1329)

																	__e.TailApply(tmp1876, tmp1885)
																	return

																}, 1)

																tmp1886 := Call(__e, PrimNS2Value(symshen_4hds), News1329)

																__e.TailApply(tmp1875, tmp1886)
																return

															} else {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															}

														}, 1)

														tmp1888 := Call(__e, PrimNS2Value(symshen_4tls), News1328)

														__e.TailApply(tmp1873, tmp1888)
														return

													} else {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp1890 := Call(__e, PrimNS2Value(symshen_4tls), News1327)

												__e.TailApply(tmp1871, tmp1890)
												return

											}, 1)

											tmp1891 := Call(__e, PrimNS2Value(symshen_4hds), News1327)

											__e.TailApply(tmp1870, tmp1891)
											return

										} else {
											__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp1893 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

									__e.TailApply(tmp1868, tmp1893)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}

						}, 1)

						tmp1896 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V1477)

						tmp1897 := Call(__e, tmp1865, tmp1896)

						__e.TailApply(tmp1839, tmp1897)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp1899 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5patterns_6 := __e.Get(1)
					_ = Parseshen_4_5patterns_6
					tmp1914 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

					if True == tmp1914 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1913 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)

						if True == tmp1913 {
							tmp1902 := MakeNative(func(__e *ControlFlow) {
								News1325 := __e.Get(1)
								_ = News1325
								tmp1911 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1325)

								if True == tmp1911 {
									tmp1904 := MakeNative(func(__e *ControlFlow) {
										Action := __e.Get(1)
										_ = Action
										tmp1905 := MakeNative(func(__e *ControlFlow) {
											News1326 := __e.Get(1)
											_ = News1326
											tmp1906 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1326)

											tmp1907 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

											tmp1908 := Call(__e, PrimNS2Value(sym_8p), tmp1907, Action)

											__e.TailApply(PrimNS2Value(symshen_4comb), tmp1906, tmp1908)
											return

										}, 1)

										tmp1909 := Call(__e, PrimNS2Value(symshen_4tls), News1325)

										__e.TailApply(tmp1905, tmp1909)
										return

									}, 1)

									tmp1910 := Call(__e, PrimNS2Value(symshen_4hds), News1325)

									__e.TailApply(tmp1904, tmp1910)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp1912 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

							__e.TailApply(tmp1902, tmp1912)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}

				}, 1)

				tmp1915 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V1477)

				tmp1916 := Call(__e, tmp1899, tmp1915)

				__e.TailApply(tmp1837, tmp1916)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp1918 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5patterns_6 := __e.Get(1)
			_ = Parseshen_4_5patterns_6
			tmp1946 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

			if True == tmp1946 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp1945 := Call(__e, PrimNS2Value(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)

				if True == tmp1945 {
					tmp1921 := MakeNative(func(__e *ControlFlow) {
						News1321 := __e.Get(1)
						_ = News1321
						tmp1943 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1321)

						if True == tmp1943 {
							tmp1923 := MakeNative(func(__e *ControlFlow) {
								Action := __e.Get(1)
								_ = Action
								tmp1924 := MakeNative(func(__e *ControlFlow) {
									News1322 := __e.Get(1)
									_ = News1322
									tmp1940 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News1322, symwhere)

									if True == tmp1940 {
										tmp1926 := MakeNative(func(__e *ControlFlow) {
											News1323 := __e.Get(1)
											_ = News1323
											tmp1938 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), News1323)

											if True == tmp1938 {
												tmp1928 := MakeNative(func(__e *ControlFlow) {
													Guard := __e.Get(1)
													_ = Guard
													tmp1929 := MakeNative(func(__e *ControlFlow) {
														News1324 := __e.Get(1)
														_ = News1324
														tmp1930 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1324)

														tmp1931 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

														tmp1932 := PrimCons(Action, Nil)

														tmp1933 := PrimCons(Guard, tmp1932)

														tmp1934 := PrimCons(symwhere, tmp1933)

														tmp1935 := Call(__e, PrimNS2Value(sym_8p), tmp1931, tmp1934)

														__e.TailApply(PrimNS2Value(symshen_4comb), tmp1930, tmp1935)
														return

													}, 1)

													tmp1936 := Call(__e, PrimNS2Value(symshen_4tls), News1323)

													__e.TailApply(tmp1929, tmp1936)
													return

												}, 1)

												tmp1937 := Call(__e, PrimNS2Value(symshen_4hds), News1323)

												__e.TailApply(tmp1928, tmp1937)
												return

											} else {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp1939 := Call(__e, PrimNS2Value(symshen_4tls), News1322)

										__e.TailApply(tmp1926, tmp1939)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp1941 := Call(__e, PrimNS2Value(symshen_4tls), News1321)

								__e.TailApply(tmp1924, tmp1941)
								return

							}, 1)

							tmp1942 := Call(__e, PrimNS2Value(symshen_4hds), News1321)

							__e.TailApply(tmp1923, tmp1942)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp1944 := Call(__e, PrimNS2Value(symshen_4tls), Parseshen_4_5patterns_6)

					__e.TailApply(tmp1921, tmp1944)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
					return
				}

			}

		}, 1)

		tmp1947 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V1477)

		tmp1948 := Call(__e, tmp1918, tmp1947)

		__e.TailApply(tmp1835, tmp1948)
		return

	}, 1)

	tmp1949 := Call(__e, PrimNS2Value(symdef), symshen_4_5rule_6, tmp1834)

	_ = tmp1949

	tmp1950 := MakeNative(func(__e *ControlFlow) {
		V1478 := __e.Get(1)
		_ = V1478
		tmp1951 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1962 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp1962 {
				tmp1953 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp1955 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp1955 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp1956 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp1959 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp1959 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1958 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp1958, Nil)
						return

					}

				}, 1)

				tmp1960 := Call(__e, PrimNS2Value(sym_5e_6), V1478)

				tmp1961 := Call(__e, tmp1956, tmp1960)

				__e.TailApply(tmp1953, tmp1961)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp1963 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5pattern_6 := __e.Get(1)
			_ = Parseshen_4_5pattern_6
			tmp1973 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5pattern_6)

			if True == tmp1973 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp1965 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5patterns_6 := __e.Get(1)
					_ = Parseshen_4_5patterns_6
					tmp1971 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)

					if True == tmp1971 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp1967 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5patterns_6)

						tmp1968 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5pattern_6)

						tmp1969 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5patterns_6)

						tmp1970 := PrimCons(tmp1968, tmp1969)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp1967, tmp1970)
						return

					}

				}, 1)

				tmp1972 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), Parseshen_4_5pattern_6)

				__e.TailApply(tmp1965, tmp1972)
				return

			}

		}, 1)

		tmp1974 := Call(__e, PrimNS2Value(symshen_4_5pattern_6), V1478)

		tmp1975 := Call(__e, tmp1963, tmp1974)

		__e.TailApply(tmp1951, tmp1975)
		return

	}, 1)

	tmp1976 := Call(__e, PrimNS2Value(symdef), symshen_4_5patterns_6, tmp1950)

	_ = tmp1976

	tmp1977 := MakeNative(func(__e *ControlFlow) {
		V1479 := __e.Get(1)
		_ = V1479
		tmp1978 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp2032 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp2032 {
				tmp1980 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp2007 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp2007 {
						tmp1982 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp1994 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp1994 {
								tmp1984 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp1986 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

									if True == tmp1986 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp1987 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5simple_1pattern_6 := __e.Get(1)
									_ = Parseshen_4_5simple_1pattern_6
									tmp1991 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5simple_1pattern_6)

									if True == tmp1991 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp1989 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5simple_1pattern_6)

										tmp1990 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5simple_1pattern_6)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp1989, tmp1990)
										return

									}

								}, 1)

								tmp1992 := Call(__e, PrimNS2Value(symshen_4_5simple_1pattern_6), V1479)

								tmp1993 := Call(__e, tmp1987, tmp1992)

								__e.TailApply(tmp1984, tmp1993)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp2006 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1479)

						var ifres1995 Obj

						if True == tmp2006 {
							tmp1997 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp1998 := MakeNative(func(__e *ControlFlow) {
									News1337 := __e.Get(1)
									_ = News1337
									tmp2002 := PrimIsPair(X)

									if True == tmp2002 {
										tmp2000 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1337)

										tmp2001 := Call(__e, PrimNS2Value(symshen_4constructor_1error), X)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp2000, tmp2001)
										return

									} else {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp2003 := Call(__e, PrimNS2Value(symshen_4tls), V1479)

								__e.TailApply(tmp1998, tmp2003)
								return

							}, 1)

							tmp2004 := Call(__e, PrimNS2Value(symshen_4hds), V1479)

							tmp2005 := Call(__e, tmp1997, tmp2004)

							ifres1995 = tmp2005

						} else {
							tmp1996 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

							ifres1995 = tmp1996

						}

						__e.TailApply(tmp1982, ifres1995)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp2031 := Call(__e, PrimNS2Value(symshen_4ccons_2), V1479)

				var ifres2008 Obj

				if True == tmp2031 {
					tmp2010 := MakeNative(func(__e *ControlFlow) {
						SynCons := __e.Get(1)
						_ = SynCons
						tmp2026 := Call(__e, PrimNS2Value(symshen_4_ahd_2), SynCons, symvector)

						if True == tmp2026 {
							tmp2012 := MakeNative(func(__e *ControlFlow) {
								News1335 := __e.Get(1)
								_ = News1335
								tmp2024 := Call(__e, PrimNS2Value(symshen_4_ahd_2), News1335, MakeNumber(0))

								if True == tmp2024 {
									tmp2014 := MakeNative(func(__e *ControlFlow) {
										News1336 := __e.Get(1)
										_ = News1336
										tmp2015 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5end_6 := __e.Get(1)
											_ = Parseshen_4_5end_6
											tmp2021 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

											if True == tmp2021 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp2017 := Call(__e, PrimNS2Value(symshen_4tlstream), V1479)

												tmp2018 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp2017)

												tmp2019 := PrimCons(MakeNumber(0), Nil)

												tmp2020 := PrimCons(symvector, tmp2019)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp2018, tmp2020)
												return

											}

										}, 1)

										tmp2022 := Call(__e, PrimNS2Value(symshen_4_5end_6), News1336)

										__e.TailApply(tmp2015, tmp2022)
										return

									}, 1)

									tmp2023 := Call(__e, PrimNS2Value(symshen_4tls), News1335)

									__e.TailApply(tmp2014, tmp2023)
									return

								} else {
									__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp2025 := Call(__e, PrimNS2Value(symshen_4tls), SynCons)

							__e.TailApply(tmp2012, tmp2025)
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp2027 := Call(__e, PrimNS2Value(symshen_4hds), V1479)

					tmp2028 := Call(__e, PrimNS2Value(symshen_4_5_1out), V1479)

					tmp2029 := Call(__e, PrimNS2Value(symshen_4comb), tmp2027, tmp2028)

					tmp2030 := Call(__e, tmp2010, tmp2029)

					ifres2008 = tmp2030

				} else {
					tmp2009 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

					ifres2008 = tmp2009

				}

				__e.TailApply(tmp1980, ifres2008)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp2064 := Call(__e, PrimNS2Value(symshen_4ccons_2), V1479)

		var ifres2033 Obj

		if True == tmp2064 {
			tmp2035 := MakeNative(func(__e *ControlFlow) {
				SynCons := __e.Get(1)
				_ = SynCons
				tmp2036 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5constructor_6 := __e.Get(1)
					_ = Parseshen_4_5constructor_6
					tmp2058 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5constructor_6)

					if True == tmp2058 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp2038 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5pattern1_6 := __e.Get(1)
							_ = Parseshen_4_5pattern1_6
							tmp2056 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5pattern1_6)

							if True == tmp2056 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp2040 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5pattern2_6 := __e.Get(1)
									_ = Parseshen_4_5pattern2_6
									tmp2054 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5pattern2_6)

									if True == tmp2054 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp2042 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5end_6 := __e.Get(1)
											_ = Parseshen_4_5end_6
											tmp2052 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5end_6)

											if True == tmp2052 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp2044 := Call(__e, PrimNS2Value(symshen_4tlstream), V1479)

												tmp2045 := Call(__e, PrimNS2Value(symshen_4in_1_6), tmp2044)

												tmp2046 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5constructor_6)

												tmp2047 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5pattern1_6)

												tmp2048 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5pattern2_6)

												tmp2049 := PrimCons(tmp2048, Nil)

												tmp2050 := PrimCons(tmp2047, tmp2049)

												tmp2051 := PrimCons(tmp2046, tmp2050)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp2045, tmp2051)
												return

											}

										}, 1)

										tmp2053 := Call(__e, PrimNS2Value(symshen_4_5end_6), Parseshen_4_5pattern2_6)

										__e.TailApply(tmp2042, tmp2053)
										return

									}

								}, 1)

								tmp2055 := Call(__e, PrimNS2Value(symshen_4_5pattern2_6), Parseshen_4_5pattern1_6)

								__e.TailApply(tmp2040, tmp2055)
								return

							}

						}, 1)

						tmp2057 := Call(__e, PrimNS2Value(symshen_4_5pattern1_6), Parseshen_4_5constructor_6)

						__e.TailApply(tmp2038, tmp2057)
						return

					}

				}, 1)

				tmp2059 := Call(__e, PrimNS2Value(symshen_4_5constructor_6), SynCons)

				__e.TailApply(tmp2036, tmp2059)
				return

			}, 1)

			tmp2060 := Call(__e, PrimNS2Value(symshen_4hds), V1479)

			tmp2061 := Call(__e, PrimNS2Value(symshen_4_5_1out), V1479)

			tmp2062 := Call(__e, PrimNS2Value(symshen_4comb), tmp2060, tmp2061)

			tmp2063 := Call(__e, tmp2035, tmp2062)

			ifres2033 = tmp2063

		} else {
			tmp2034 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres2033 = tmp2034

		}

		__e.TailApply(tmp1978, ifres2033)
		return

	}, 1)

	tmp2065 := Call(__e, PrimNS2Value(symdef), symshen_4_5pattern_6, tmp1977)

	_ = tmp2065

	tmp2066 := MakeNative(func(__e *ControlFlow) {
		V1480 := __e.Get(1)
		_ = V1480
		tmp2067 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp2069 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp2069 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp2080 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1480)

		var ifres2070 Obj

		if True == tmp2080 {
			tmp2072 := MakeNative(func(__e *ControlFlow) {
				C := __e.Get(1)
				_ = C
				tmp2073 := MakeNative(func(__e *ControlFlow) {
					News1339 := __e.Get(1)
					_ = News1339
					tmp2076 := Call(__e, PrimNS2Value(symshen_4constructor_2), C)

					if True == tmp2076 {
						tmp2075 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1339)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp2075, C)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp2077 := Call(__e, PrimNS2Value(symshen_4tls), V1480)

				__e.TailApply(tmp2073, tmp2077)
				return

			}, 1)

			tmp2078 := Call(__e, PrimNS2Value(symshen_4hds), V1480)

			tmp2079 := Call(__e, tmp2072, tmp2078)

			ifres2070 = tmp2079

		} else {
			tmp2071 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres2070 = tmp2071

		}

		__e.TailApply(tmp2067, ifres2070)
		return

	}, 1)

	tmp2081 := Call(__e, PrimNS2Value(symdef), symshen_4_5constructor_6, tmp2066)

	_ = tmp2081

	tmp2082 := MakeNative(func(__e *ControlFlow) {
		V1481 := __e.Get(1)
		_ = V1481
		tmp2083 := PrimCons(sym_8v, Nil)

		tmp2084 := PrimCons(sym_8s, tmp2083)

		tmp2085 := PrimCons(sym_8p, tmp2084)

		tmp2086 := PrimCons(symcons, tmp2085)

		__e.TailApply(PrimNS2Value(symelement_2), V1481, tmp2086)
		return

	}, 1)

	tmp2087 := Call(__e, PrimNS2Value(symdef), symshen_4constructor_2, tmp2082)

	_ = tmp2087

	tmp2088 := MakeNative(func(__e *ControlFlow) {
		V1482 := __e.Get(1)
		_ = V1482
		tmp2089 := Call(__e, PrimNS2Value(symshen_4app), V1482, MakeString(" is not a legitimate constructor\n"), symshen_4r)

		__e.Return(PrimSimpleError(tmp2089))
		return

	}, 1)

	tmp2090 := Call(__e, PrimNS2Value(symdef), symshen_4constructor_1error, tmp2088)

	_ = tmp2090

	tmp2091 := MakeNative(func(__e *ControlFlow) {
		V1483 := __e.Get(1)
		_ = V1483
		tmp2092 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp2111 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp2111 {
				tmp2094 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp2096 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp2096 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp2110 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1483)

				var ifres2097 Obj

				if True == tmp2110 {
					tmp2099 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp2100 := MakeNative(func(__e *ControlFlow) {
							News1342 := __e.Get(1)
							_ = News1342
							tmp2103 := PrimCons(sym_5_1, Nil)

							tmp2104 := PrimCons(sym_1_6, tmp2103)

							tmp2105 := Call(__e, PrimNS2Value(symelement_2), X, tmp2104)

							tmp2106 := PrimNot(tmp2105)

							if True == tmp2106 {
								tmp2102 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1342)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp2102, X)
								return

							} else {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp2107 := Call(__e, PrimNS2Value(symshen_4tls), V1483)

						__e.TailApply(tmp2100, tmp2107)
						return

					}, 1)

					tmp2108 := Call(__e, PrimNS2Value(symshen_4hds), V1483)

					tmp2109 := Call(__e, tmp2099, tmp2108)

					ifres2097 = tmp2109

				} else {
					tmp2098 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

					ifres2097 = tmp2098

				}

				__e.TailApply(tmp2094, ifres2097)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp2124 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V1483)

		var ifres2112 Obj

		if True == tmp2124 {
			tmp2114 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp2115 := MakeNative(func(__e *ControlFlow) {
					News1341 := __e.Get(1)
					_ = News1341
					tmp2120 := PrimEqual(X, sym__)

					if True == tmp2120 {
						tmp2117 := Call(__e, PrimNS2Value(symshen_4in_1_6), News1341)

						tmp2118 := Call(__e, PrimNS2Value(symprotect), symY)

						tmp2119 := Call(__e, PrimNS2Value(symgensym), tmp2118)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp2117, tmp2119)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp2121 := Call(__e, PrimNS2Value(symshen_4tls), V1483)

				__e.TailApply(tmp2115, tmp2121)
				return

			}, 1)

			tmp2122 := Call(__e, PrimNS2Value(symshen_4hds), V1483)

			tmp2123 := Call(__e, tmp2114, tmp2122)

			ifres2112 = tmp2123

		} else {
			tmp2113 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres2112 = tmp2113

		}

		__e.TailApply(tmp2092, ifres2112)
		return

	}, 1)

	tmp2125 := Call(__e, PrimNS2Value(symdef), symshen_4_5simple_1pattern_6, tmp2091)

	_ = tmp2125

	tmp2126 := MakeNative(func(__e *ControlFlow) {
		V1484 := __e.Get(1)
		_ = V1484
		tmp2127 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp2129 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp2129 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp2130 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5pattern_6 := __e.Get(1)
			_ = Parseshen_4_5pattern_6
			tmp2134 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5pattern_6)

			if True == tmp2134 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp2132 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5pattern_6)

				tmp2133 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5pattern_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp2132, tmp2133)
				return

			}

		}, 1)

		tmp2135 := Call(__e, PrimNS2Value(symshen_4_5pattern_6), V1484)

		tmp2136 := Call(__e, tmp2130, tmp2135)

		__e.TailApply(tmp2127, tmp2136)
		return

	}, 1)

	tmp2137 := Call(__e, PrimNS2Value(symdef), symshen_4_5pattern1_6, tmp2126)

	_ = tmp2137

	tmp2138 := MakeNative(func(__e *ControlFlow) {
		V1485 := __e.Get(1)
		_ = V1485
		tmp2139 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp2141 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp2141 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp2142 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5pattern_6 := __e.Get(1)
			_ = Parseshen_4_5pattern_6
			tmp2146 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5pattern_6)

			if True == tmp2146 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp2144 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5pattern_6)

				tmp2145 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5pattern_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp2144, tmp2145)
				return

			}

		}, 1)

		tmp2147 := Call(__e, PrimNS2Value(symshen_4_5pattern_6), V1485)

		tmp2148 := Call(__e, tmp2142, tmp2147)

		__e.TailApply(tmp2139, tmp2148)
		return

	}, 1)

	tmp2149 := Call(__e, PrimNS2Value(symdef), symshen_4_5pattern2_6, tmp2138)

	_ = tmp2149

	tmp2150 := MakeNative(func(__e *ControlFlow) {
		V1486 := __e.Get(1)
		_ = V1486
		tmp2151 := MakeNative(func(__e *ControlFlow) {
			V := __e.Get(1)
			_ = V
			tmp2152 := MakeNative(func(__e *ControlFlow) {
				Print := __e.Get(1)
				_ = Print
				tmp2153 := MakeNative(func(__e *ControlFlow) {
					Named := __e.Get(1)
					_ = Named
					__e.Return(Named)
					return
				}, 1)

				tmp2154 := PrimStr(V1486)

				tmp2155 := Call(__e, PrimNS2Value(sym_8s), tmp2154, MakeString(")"))

				tmp2156 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp2155)

				tmp2157 := Call(__e, PrimNS2Value(sym_8s), MakeString("n"), tmp2156)

				tmp2158 := Call(__e, PrimNS2Value(sym_8s), MakeString("f"), tmp2157)

				tmp2159 := Call(__e, PrimNS2Value(sym_8s), MakeString("("), tmp2158)

				tmp2160 := PrimVectorSet(Print, MakeNumber(1), tmp2159)

				__e.TailApply(tmp2153, tmp2160)
				return

			}, 1)

			tmp2161 := PrimVectorSet(V, MakeNumber(0), symshen_4printF)

			__e.TailApply(tmp2152, tmp2161)
			return

		}, 1)

		tmp2162 := PrimAbsvector(MakeNumber(2))

		__e.TailApply(tmp2151, tmp2162)
		return

	}, 1)

	tmp2163 := Call(__e, PrimNS2Value(symdef), symshen_4fn_1print, tmp2150)

	_ = tmp2163

	tmp2164 := MakeNative(func(__e *ControlFlow) {
		V1487 := __e.Get(1)
		_ = V1487
		__e.Return(PrimVectorGet(V1487, MakeNumber(1)))
		return
	}, 1)

	tmp2165 := Call(__e, PrimNS2Value(symdef), symshen_4printF, tmp2164)

	_ = tmp2165

	tmp2166 := MakeNative(func(__e *ControlFlow) {
		V1492 := __e.Get(1)
		_ = V1492
		V1493 := __e.Get(2)
		_ = V1493
		tmp2190 := PrimIsPair(V1493)

		var ifres2186 Obj

		if True == tmp2190 {
			tmp2188 := PrimTail(V1493)

			tmp2189 := PrimEqual(Nil, tmp2188)

			var ifres2187 Obj

			if True == tmp2189 {
				ifres2187 = True

			} else {
				ifres2187 = False

			}

			ifres2186 = ifres2187

		} else {
			ifres2186 = False

		}

		if True == ifres2186 {
			tmp2185 := PrimHead(V1493)

			__e.TailApply(PrimNS2Value(symlength), tmp2185)
			return

		} else {
			tmp2184 := PrimIsPair(V1493)

			var ifres2172 Obj

			if True == tmp2184 {
				tmp2182 := PrimTail(V1493)

				tmp2183 := PrimIsPair(tmp2182)

				var ifres2174 Obj

				if True == tmp2183 {
					tmp2176 := PrimHead(V1493)

					tmp2177 := Call(__e, PrimNS2Value(symlength), tmp2176)

					tmp2178 := PrimTail(V1493)

					tmp2179 := PrimHead(tmp2178)

					tmp2180 := Call(__e, PrimNS2Value(symlength), tmp2179)

					tmp2181 := PrimEqual(tmp2177, tmp2180)

					var ifres2175 Obj

					if True == tmp2181 {
						ifres2175 = True

					} else {
						ifres2175 = False

					}

					ifres2174 = ifres2175

				} else {
					ifres2174 = False

				}

				var ifres2173 Obj

				if True == ifres2174 {
					ifres2173 = True

				} else {
					ifres2173 = False

				}

				ifres2172 = ifres2173

			} else {
				ifres2172 = False

			}

			if True == ifres2172 {
				tmp2171 := PrimTail(V1493)

				__e.TailApply(PrimNS2Value(symshen_4arity_1chk), V1492, tmp2171)
				return

			} else {
				tmp2169 := Call(__e, PrimNS2Value(symshen_4app), V1492, MakeString("\n"), symshen_4a)

				tmp2170 := PrimStringConcat(MakeString("arity error in "), tmp2169)

				__e.Return(PrimSimpleError(tmp2170))
				return

			}

		}

	}, 2)

	tmp2191 := Call(__e, PrimNS2Value(symdef), symshen_4arity_1chk, tmp2166)

	_ = tmp2191

	tmp2192 := MakeNative(func(__e *ControlFlow) {
		V1494 := __e.Get(1)
		_ = V1494
		V1495 := __e.Get(2)
		_ = V1495
		tmp2198 := Call(__e, PrimNS2Value(symtuple_2), V1495)

		if True == tmp2198 {
			tmp2194 := Call(__e, PrimNS2Value(symfst), V1495)

			tmp2195 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp2194)

			tmp2196 := Call(__e, PrimNS2Value(symsnd), V1495)

			tmp2197 := Call(__e, PrimNS2Value(symshen_4find_1free_1vars), tmp2195, tmp2196)

			__e.TailApply(PrimNS2Value(symshen_4free_1variable_1error_1message), V1494, tmp2197)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4free_1var_1chk)
			return
		}

	}, 2)

	tmp2199 := Call(__e, PrimNS2Value(symdef), symshen_4free_1var_1chk, tmp2192)

	_ = tmp2199

	tmp2200 := MakeNative(func(__e *ControlFlow) {
		V1496 := __e.Get(1)
		_ = V1496
		V1497 := __e.Get(2)
		_ = V1497
		tmp2212 := Call(__e, PrimNS2Value(symempty_2), V1497)

		if True == tmp2212 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp2202 := Call(__e, PrimNS2Value(symshen_4app), V1496, MakeString(":"), symshen_4a)

			tmp2203 := PrimStringConcat(MakeString("free variables in "), tmp2202)

			tmp2204 := Call(__e, PrimNS2Value(symstoutput))

			tmp2205 := Call(__e, PrimNS2Value(sympr), tmp2203, tmp2204)

			_ = tmp2205

			tmp2206 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp2207 := Call(__e, PrimNS2Value(symshen_4app), X, MakeString(""), symshen_4a)

				tmp2208 := PrimStringConcat(MakeString(" "), tmp2207)

				tmp2209 := Call(__e, PrimNS2Value(symstoutput))

				__e.TailApply(PrimNS2Value(sympr), tmp2208, tmp2209)
				return

			}, 1)

			tmp2210 := Call(__e, PrimNS2Value(symmap), tmp2206, V1497)

			_ = tmp2210

			tmp2211 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

			_ = tmp2211

			__e.TailApply(PrimNS2Value(symabort))
			return

		}

	}, 2)

	tmp2213 := Call(__e, PrimNS2Value(symdef), symshen_4free_1variable_1error_1message, tmp2200)

	_ = tmp2213

	tmp2214 := MakeNative(func(__e *ControlFlow) {
		V1500 := __e.Get(1)
		_ = V1500
		tmp2222 := PrimIsVariable(V1500)

		if True == tmp2222 {
			__e.Return(PrimCons(V1500, Nil))
			return
		} else {
			tmp2221 := PrimIsPair(V1500)

			if True == tmp2221 {
				tmp2217 := PrimHead(V1500)

				tmp2218 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp2217)

				tmp2219 := PrimTail(V1500)

				tmp2220 := Call(__e, PrimNS2Value(symshen_4extract_1vars), tmp2219)

				__e.TailApply(PrimNS2Value(symunion), tmp2218, tmp2220)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp2223 := Call(__e, PrimNS2Value(symdef), symshen_4extract_1vars, tmp2214)

	_ = tmp2223

	tmp2224 := MakeNative(func(__e *ControlFlow) {
		V1505 := __e.Get(1)
		_ = V1505
		V1506 := __e.Get(2)
		_ = V1506
		tmp2314 := PrimIsPair(V1506)

		var ifres2301 Obj

		if True == tmp2314 {
			tmp2312 := PrimHead(V1506)

			tmp2313 := PrimEqual(symprotect, tmp2312)

			var ifres2303 Obj

			if True == tmp2313 {
				tmp2310 := PrimTail(V1506)

				tmp2311 := PrimIsPair(tmp2310)

				var ifres2305 Obj

				if True == tmp2311 {
					tmp2307 := PrimTail(V1506)

					tmp2308 := PrimTail(tmp2307)

					tmp2309 := PrimEqual(Nil, tmp2308)

					var ifres2306 Obj

					if True == tmp2309 {
						ifres2306 = True

					} else {
						ifres2306 = False

					}

					ifres2305 = ifres2306

				} else {
					ifres2305 = False

				}

				var ifres2304 Obj

				if True == ifres2305 {
					ifres2304 = True

				} else {
					ifres2304 = False

				}

				ifres2303 = ifres2304

			} else {
				ifres2303 = False

			}

			var ifres2302 Obj

			if True == ifres2303 {
				ifres2302 = True

			} else {
				ifres2302 = False

			}

			ifres2301 = ifres2302

		} else {
			ifres2301 = False

		}

		if True == ifres2301 {
			__e.Return(Nil)
			return
		} else {
			tmp2300 := PrimIsPair(V1506)

			var ifres2274 Obj

			if True == tmp2300 {
				tmp2298 := PrimHead(V1506)

				tmp2299 := PrimEqual(symlet, tmp2298)

				var ifres2276 Obj

				if True == tmp2299 {
					tmp2296 := PrimTail(V1506)

					tmp2297 := PrimIsPair(tmp2296)

					var ifres2278 Obj

					if True == tmp2297 {
						tmp2293 := PrimTail(V1506)

						tmp2294 := PrimTail(tmp2293)

						tmp2295 := PrimIsPair(tmp2294)

						var ifres2280 Obj

						if True == tmp2295 {
							tmp2289 := PrimTail(V1506)

							tmp2290 := PrimTail(tmp2289)

							tmp2291 := PrimTail(tmp2290)

							tmp2292 := PrimIsPair(tmp2291)

							var ifres2282 Obj

							if True == tmp2292 {
								tmp2284 := PrimTail(V1506)

								tmp2285 := PrimTail(tmp2284)

								tmp2286 := PrimTail(tmp2285)

								tmp2287 := PrimTail(tmp2286)

								tmp2288 := PrimEqual(Nil, tmp2287)

								var ifres2283 Obj

								if True == tmp2288 {
									ifres2283 = True

								} else {
									ifres2283 = False

								}

								ifres2282 = ifres2283

							} else {
								ifres2282 = False

							}

							var ifres2281 Obj

							if True == ifres2282 {
								ifres2281 = True

							} else {
								ifres2281 = False

							}

							ifres2280 = ifres2281

						} else {
							ifres2280 = False

						}

						var ifres2279 Obj

						if True == ifres2280 {
							ifres2279 = True

						} else {
							ifres2279 = False

						}

						ifres2278 = ifres2279

					} else {
						ifres2278 = False

					}

					var ifres2277 Obj

					if True == ifres2278 {
						ifres2277 = True

					} else {
						ifres2277 = False

					}

					ifres2276 = ifres2277

				} else {
					ifres2276 = False

				}

				var ifres2275 Obj

				if True == ifres2276 {
					ifres2275 = True

				} else {
					ifres2275 = False

				}

				ifres2274 = ifres2275

			} else {
				ifres2274 = False

			}

			if True == ifres2274 {
				tmp2262 := PrimTail(V1506)

				tmp2263 := PrimTail(tmp2262)

				tmp2264 := PrimHead(tmp2263)

				tmp2265 := Call(__e, PrimNS2Value(symshen_4find_1free_1vars), V1505, tmp2264)

				tmp2266 := PrimTail(V1506)

				tmp2267 := PrimHead(tmp2266)

				tmp2268 := PrimCons(tmp2267, V1505)

				tmp2269 := PrimTail(V1506)

				tmp2270 := PrimTail(tmp2269)

				tmp2271 := PrimTail(tmp2270)

				tmp2272 := PrimHead(tmp2271)

				tmp2273 := Call(__e, PrimNS2Value(symshen_4find_1free_1vars), tmp2268, tmp2272)

				__e.TailApply(PrimNS2Value(symunion), tmp2265, tmp2273)
				return

			} else {
				tmp2261 := PrimIsPair(V1506)

				var ifres2242 Obj

				if True == tmp2261 {
					tmp2259 := PrimHead(V1506)

					tmp2260 := PrimEqual(symlambda, tmp2259)

					var ifres2244 Obj

					if True == tmp2260 {
						tmp2257 := PrimTail(V1506)

						tmp2258 := PrimIsPair(tmp2257)

						var ifres2246 Obj

						if True == tmp2258 {
							tmp2254 := PrimTail(V1506)

							tmp2255 := PrimTail(tmp2254)

							tmp2256 := PrimIsPair(tmp2255)

							var ifres2248 Obj

							if True == tmp2256 {
								tmp2250 := PrimTail(V1506)

								tmp2251 := PrimTail(tmp2250)

								tmp2252 := PrimTail(tmp2251)

								tmp2253 := PrimEqual(Nil, tmp2252)

								var ifres2249 Obj

								if True == tmp2253 {
									ifres2249 = True

								} else {
									ifres2249 = False

								}

								ifres2248 = ifres2249

							} else {
								ifres2248 = False

							}

							var ifres2247 Obj

							if True == ifres2248 {
								ifres2247 = True

							} else {
								ifres2247 = False

							}

							ifres2246 = ifres2247

						} else {
							ifres2246 = False

						}

						var ifres2245 Obj

						if True == ifres2246 {
							ifres2245 = True

						} else {
							ifres2245 = False

						}

						ifres2244 = ifres2245

					} else {
						ifres2244 = False

					}

					var ifres2243 Obj

					if True == ifres2244 {
						ifres2243 = True

					} else {
						ifres2243 = False

					}

					ifres2242 = ifres2243

				} else {
					ifres2242 = False

				}

				if True == ifres2242 {
					tmp2236 := PrimTail(V1506)

					tmp2237 := PrimHead(tmp2236)

					tmp2238 := PrimCons(tmp2237, V1505)

					tmp2239 := PrimTail(V1506)

					tmp2240 := PrimTail(tmp2239)

					tmp2241 := PrimHead(tmp2240)

					__e.TailApply(PrimNS2Value(symshen_4find_1free_1vars), tmp2238, tmp2241)
					return

				} else {
					tmp2235 := PrimIsPair(V1506)

					if True == tmp2235 {
						tmp2231 := PrimHead(V1506)

						tmp2232 := Call(__e, PrimNS2Value(symshen_4find_1free_1vars), V1505, tmp2231)

						tmp2233 := PrimTail(V1506)

						tmp2234 := Call(__e, PrimNS2Value(symshen_4find_1free_1vars), V1505, tmp2233)

						__e.TailApply(PrimNS2Value(symunion), tmp2232, tmp2234)
						return

					} else {
						tmp2230 := Call(__e, PrimNS2Value(symshen_4free_1variable_2), V1506, V1505)

						if True == tmp2230 {
							__e.Return(PrimCons(V1506, Nil))
							return
						} else {
							__e.Return(Nil)
							return
						}

					}

				}

			}

		}

	}, 2)

	tmp2315 := Call(__e, PrimNS2Value(symdef), symshen_4find_1free_1vars, tmp2224)

	_ = tmp2315

	tmp2316 := MakeNative(func(__e *ControlFlow) {
		V1507 := __e.Get(1)
		_ = V1507
		V1508 := __e.Get(2)
		_ = V1508
		tmp2321 := PrimIsVariable(V1507)

		if True == tmp2321 {
			tmp2319 := Call(__e, PrimNS2Value(symelement_2), V1507, V1508)

			tmp2320 := PrimNot(tmp2319)

			if True == tmp2320 {
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

	}, 2)

	tmp2322 := Call(__e, PrimNS2Value(symdef), symshen_4free_1variable_2, tmp2316)

	_ = tmp2322

	tmp2323 := MakeNative(func(__e *ControlFlow) {
		V1509 := __e.Get(1)
		_ = V1509
		V1510 := __e.Get(2)
		_ = V1510
		tmp2324 := PrimNS3Value(sym_dproperty_1vector_d)

		__e.TailApply(PrimNS2Value(symput), V1509, symshen_4source, V1510, tmp2324)
		return

	}, 2)

	tmp2325 := Call(__e, PrimNS2Value(symdef), symshen_4record_1kl, tmp2323)

	_ = tmp2325

	tmp2326 := MakeNative(func(__e *ControlFlow) {
		V1511 := __e.Get(1)
		_ = V1511
		V1512 := __e.Get(2)
		_ = V1512
		V1513 := __e.Get(3)
		_ = V1513
		tmp2327 := MakeNative(func(__e *ControlFlow) {
			Parameters := __e.Get(1)
			_ = Parameters
			tmp2328 := MakeNative(func(__e *ControlFlow) {
				Body := __e.Get(1)
				_ = Body
				tmp2329 := MakeNative(func(__e *ControlFlow) {
					Defun := __e.Get(1)
					_ = Defun
					__e.Return(Defun)
					return
				}, 1)

				tmp2330 := Call(__e, PrimNS2Value(symshen_4cond_1form), Body)

				tmp2331 := PrimCons(tmp2330, Nil)

				tmp2332 := PrimCons(Parameters, tmp2331)

				tmp2333 := PrimCons(V1511, tmp2332)

				tmp2334 := PrimCons(symdefun, tmp2333)

				__e.TailApply(tmp2329, tmp2334)
				return

			}, 1)

			tmp2335 := Call(__e, PrimNS2Value(symshen_4kl_1body), V1512, Parameters)

			tmp2336 := Call(__e, PrimNS2Value(symshen_4scan_1body), V1511, tmp2335)

			__e.TailApply(tmp2328, tmp2336)
			return

		}, 1)

		tmp2337 := Call(__e, PrimNS2Value(symshen_4parameters), V1513)

		__e.TailApply(tmp2327, tmp2337)
		return

	}, 3)

	tmp2338 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1to_1kl, tmp2326)

	_ = tmp2338

	tmp2339 := MakeNative(func(__e *ControlFlow) {
		V1514 := __e.Get(1)
		_ = V1514
		tmp2345 := PrimEqual(MakeNumber(0), V1514)

		if True == tmp2345 {
			__e.Return(Nil)
			return
		} else {
			tmp2341 := Call(__e, PrimNS2Value(symprotect), symV)

			tmp2342 := Call(__e, PrimNS2Value(symgensym), tmp2341)

			tmp2343 := PrimNumberSubtract(V1514, MakeNumber(1))

			tmp2344 := Call(__e, PrimNS2Value(symshen_4parameters), tmp2343)

			__e.Return(PrimCons(tmp2342, tmp2344))
			return

		}

	}, 1)

	tmp2346 := Call(__e, PrimNS2Value(symdef), symshen_4parameters, tmp2339)

	_ = tmp2346

	tmp2347 := MakeNative(func(__e *ControlFlow) {
		V1517 := __e.Get(1)
		_ = V1517
		tmp2371 := PrimIsPair(V1517)

		var ifres2351 Obj

		if True == tmp2371 {
			tmp2369 := PrimHead(V1517)

			tmp2370 := PrimIsPair(tmp2369)

			var ifres2353 Obj

			if True == tmp2370 {
				tmp2366 := PrimHead(V1517)

				tmp2367 := PrimHead(tmp2366)

				tmp2368 := PrimEqual(True, tmp2367)

				var ifres2355 Obj

				if True == tmp2368 {
					tmp2363 := PrimHead(V1517)

					tmp2364 := PrimTail(tmp2363)

					tmp2365 := PrimIsPair(tmp2364)

					var ifres2357 Obj

					if True == tmp2365 {
						tmp2359 := PrimHead(V1517)

						tmp2360 := PrimTail(tmp2359)

						tmp2361 := PrimTail(tmp2360)

						tmp2362 := PrimEqual(Nil, tmp2361)

						var ifres2358 Obj

						if True == tmp2362 {
							ifres2358 = True

						} else {
							ifres2358 = False

						}

						ifres2357 = ifres2358

					} else {
						ifres2357 = False

					}

					var ifres2356 Obj

					if True == ifres2357 {
						ifres2356 = True

					} else {
						ifres2356 = False

					}

					ifres2355 = ifres2356

				} else {
					ifres2355 = False

				}

				var ifres2354 Obj

				if True == ifres2355 {
					ifres2354 = True

				} else {
					ifres2354 = False

				}

				ifres2353 = ifres2354

			} else {
				ifres2353 = False

			}

			var ifres2352 Obj

			if True == ifres2353 {
				ifres2352 = True

			} else {
				ifres2352 = False

			}

			ifres2351 = ifres2352

		} else {
			ifres2351 = False

		}

		if True == ifres2351 {
			tmp2349 := PrimHead(V1517)

			tmp2350 := PrimTail(tmp2349)

			__e.Return(PrimHead(tmp2350))
			return

		} else {
			__e.Return(PrimCons(symcond, V1517))
			return
		}

	}, 1)

	tmp2372 := Call(__e, PrimNS2Value(symdef), symshen_4cond_1form, tmp2347)

	_ = tmp2372

	tmp2373 := MakeNative(func(__e *ControlFlow) {
		V1526 := __e.Get(1)
		_ = V1526
		V1527 := __e.Get(2)
		_ = V1527
		tmp2419 := PrimEqual(Nil, V1527)

		if True == tmp2419 {
			tmp2415 := PrimCons(V1526, Nil)

			tmp2416 := PrimCons(symshen_4f_1error, tmp2415)

			tmp2417 := PrimCons(tmp2416, Nil)

			tmp2418 := PrimCons(True, tmp2417)

			__e.Return(PrimCons(tmp2418, Nil))
			return

		} else {
			tmp2414 := PrimIsPair(V1527)

			var ifres2410 Obj

			if True == tmp2414 {
				tmp2412 := PrimHead(V1527)

				tmp2413 := Call(__e, PrimNS2Value(symshen_4choicepoint_2), tmp2412)

				var ifres2411 Obj

				if True == tmp2413 {
					ifres2411 = True

				} else {
					ifres2411 = False

				}

				ifres2410 = ifres2411

			} else {
				ifres2410 = False

			}

			if True == ifres2410 {
				tmp2404 := Call(__e, PrimNS2Value(symprotect), symFreeze)

				tmp2405 := Call(__e, PrimNS2Value(symgensym), tmp2404)

				tmp2406 := Call(__e, PrimNS2Value(symprotect), symResult)

				tmp2407 := Call(__e, PrimNS2Value(symgensym), tmp2406)

				tmp2408 := PrimHead(V1527)

				tmp2409 := PrimTail(V1527)

				__e.TailApply(PrimNS2Value(symshen_4choicepoint), V1526, tmp2405, tmp2407, tmp2408, tmp2409)
				return

			} else {
				tmp2403 := PrimIsPair(V1527)

				var ifres2383 Obj

				if True == tmp2403 {
					tmp2401 := PrimHead(V1527)

					tmp2402 := PrimIsPair(tmp2401)

					var ifres2385 Obj

					if True == tmp2402 {
						tmp2398 := PrimHead(V1527)

						tmp2399 := PrimHead(tmp2398)

						tmp2400 := PrimEqual(True, tmp2399)

						var ifres2387 Obj

						if True == tmp2400 {
							tmp2395 := PrimHead(V1527)

							tmp2396 := PrimTail(tmp2395)

							tmp2397 := PrimIsPair(tmp2396)

							var ifres2389 Obj

							if True == tmp2397 {
								tmp2391 := PrimHead(V1527)

								tmp2392 := PrimTail(tmp2391)

								tmp2393 := PrimTail(tmp2392)

								tmp2394 := PrimEqual(Nil, tmp2393)

								var ifres2390 Obj

								if True == tmp2394 {
									ifres2390 = True

								} else {
									ifres2390 = False

								}

								ifres2389 = ifres2390

							} else {
								ifres2389 = False

							}

							var ifres2388 Obj

							if True == ifres2389 {
								ifres2388 = True

							} else {
								ifres2388 = False

							}

							ifres2387 = ifres2388

						} else {
							ifres2387 = False

						}

						var ifres2386 Obj

						if True == ifres2387 {
							ifres2386 = True

						} else {
							ifres2386 = False

						}

						ifres2385 = ifres2386

					} else {
						ifres2385 = False

					}

					var ifres2384 Obj

					if True == ifres2385 {
						ifres2384 = True

					} else {
						ifres2384 = False

					}

					ifres2383 = ifres2384

				} else {
					ifres2383 = False

				}

				if True == ifres2383 {
					tmp2382 := PrimHead(V1527)

					__e.Return(PrimCons(tmp2382, Nil))
					return

				} else {
					tmp2381 := PrimIsPair(V1527)

					if True == tmp2381 {
						tmp2378 := PrimHead(V1527)

						tmp2379 := PrimTail(V1527)

						tmp2380 := Call(__e, PrimNS2Value(symshen_4scan_1body), V1526, tmp2379)

						__e.Return(PrimCons(tmp2378, tmp2380))
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.scan-body")))
						return
					}

				}

			}

		}

	}, 2)

	tmp2420 := Call(__e, PrimNS2Value(symdef), symshen_4scan_1body, tmp2373)

	_ = tmp2420

	tmp2421 := MakeNative(func(__e *ControlFlow) {
		V1534 := __e.Get(1)
		_ = V1534
		tmp2456 := PrimIsPair(V1534)

		var ifres2423 Obj

		if True == tmp2456 {
			tmp2454 := PrimTail(V1534)

			tmp2455 := PrimIsPair(tmp2454)

			var ifres2425 Obj

			if True == tmp2455 {
				tmp2451 := PrimTail(V1534)

				tmp2452 := PrimHead(tmp2451)

				tmp2453 := PrimIsPair(tmp2452)

				var ifres2427 Obj

				if True == tmp2453 {
					tmp2447 := PrimTail(V1534)

					tmp2448 := PrimHead(tmp2447)

					tmp2449 := PrimHead(tmp2448)

					tmp2450 := PrimEqual(symshen_4choicepoint_b, tmp2449)

					var ifres2429 Obj

					if True == tmp2450 {
						tmp2443 := PrimTail(V1534)

						tmp2444 := PrimHead(tmp2443)

						tmp2445 := PrimTail(tmp2444)

						tmp2446 := PrimIsPair(tmp2445)

						var ifres2431 Obj

						if True == tmp2446 {
							tmp2438 := PrimTail(V1534)

							tmp2439 := PrimHead(tmp2438)

							tmp2440 := PrimTail(tmp2439)

							tmp2441 := PrimTail(tmp2440)

							tmp2442 := PrimEqual(Nil, tmp2441)

							var ifres2433 Obj

							if True == tmp2442 {
								tmp2435 := PrimTail(V1534)

								tmp2436 := PrimTail(tmp2435)

								tmp2437 := PrimEqual(Nil, tmp2436)

								var ifres2434 Obj

								if True == tmp2437 {
									ifres2434 = True

								} else {
									ifres2434 = False

								}

								ifres2433 = ifres2434

							} else {
								ifres2433 = False

							}

							var ifres2432 Obj

							if True == ifres2433 {
								ifres2432 = True

							} else {
								ifres2432 = False

							}

							ifres2431 = ifres2432

						} else {
							ifres2431 = False

						}

						var ifres2430 Obj

						if True == ifres2431 {
							ifres2430 = True

						} else {
							ifres2430 = False

						}

						ifres2429 = ifres2430

					} else {
						ifres2429 = False

					}

					var ifres2428 Obj

					if True == ifres2429 {
						ifres2428 = True

					} else {
						ifres2428 = False

					}

					ifres2427 = ifres2428

				} else {
					ifres2427 = False

				}

				var ifres2426 Obj

				if True == ifres2427 {
					ifres2426 = True

				} else {
					ifres2426 = False

				}

				ifres2425 = ifres2426

			} else {
				ifres2425 = False

			}

			var ifres2424 Obj

			if True == ifres2425 {
				ifres2424 = True

			} else {
				ifres2424 = False

			}

			ifres2423 = ifres2424

		} else {
			ifres2423 = False

		}

		if True == ifres2423 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp2457 := Call(__e, PrimNS2Value(symdef), symshen_4choicepoint_2, tmp2421)

	_ = tmp2457

	tmp2458 := MakeNative(func(__e *ControlFlow) {
		V1550 := __e.Get(1)
		_ = V1550
		V1551 := __e.Get(2)
		_ = V1551
		V1552 := __e.Get(3)
		_ = V1552
		V1553 := __e.Get(4)
		_ = V1553
		V1554 := __e.Get(5)
		_ = V1554
		tmp2650 := PrimIsPair(V1553)

		var ifres2572 Obj

		if True == tmp2650 {
			tmp2648 := PrimTail(V1553)

			tmp2649 := PrimIsPair(tmp2648)

			var ifres2574 Obj

			if True == tmp2649 {
				tmp2645 := PrimTail(V1553)

				tmp2646 := PrimHead(tmp2645)

				tmp2647 := PrimIsPair(tmp2646)

				var ifres2576 Obj

				if True == tmp2647 {
					tmp2641 := PrimTail(V1553)

					tmp2642 := PrimHead(tmp2641)

					tmp2643 := PrimTail(tmp2642)

					tmp2644 := PrimIsPair(tmp2643)

					var ifres2578 Obj

					if True == tmp2644 {
						tmp2636 := PrimTail(V1553)

						tmp2637 := PrimHead(tmp2636)

						tmp2638 := PrimTail(tmp2637)

						tmp2639 := PrimHead(tmp2638)

						tmp2640 := PrimIsPair(tmp2639)

						var ifres2580 Obj

						if True == tmp2640 {
							tmp2630 := PrimTail(V1553)

							tmp2631 := PrimHead(tmp2630)

							tmp2632 := PrimTail(tmp2631)

							tmp2633 := PrimHead(tmp2632)

							tmp2634 := PrimHead(tmp2633)

							tmp2635 := PrimEqual(symfail_1if, tmp2634)

							var ifres2582 Obj

							if True == tmp2635 {
								tmp2624 := PrimTail(V1553)

								tmp2625 := PrimHead(tmp2624)

								tmp2626 := PrimTail(tmp2625)

								tmp2627 := PrimHead(tmp2626)

								tmp2628 := PrimTail(tmp2627)

								tmp2629 := PrimIsPair(tmp2628)

								var ifres2584 Obj

								if True == tmp2629 {
									tmp2617 := PrimTail(V1553)

									tmp2618 := PrimHead(tmp2617)

									tmp2619 := PrimTail(tmp2618)

									tmp2620 := PrimHead(tmp2619)

									tmp2621 := PrimTail(tmp2620)

									tmp2622 := PrimTail(tmp2621)

									tmp2623 := PrimIsPair(tmp2622)

									var ifres2586 Obj

									if True == tmp2623 {
										tmp2609 := PrimTail(V1553)

										tmp2610 := PrimHead(tmp2609)

										tmp2611 := PrimTail(tmp2610)

										tmp2612 := PrimHead(tmp2611)

										tmp2613 := PrimTail(tmp2612)

										tmp2614 := PrimTail(tmp2613)

										tmp2615 := PrimTail(tmp2614)

										tmp2616 := PrimEqual(Nil, tmp2615)

										var ifres2588 Obj

										if True == tmp2616 {
											tmp2604 := PrimTail(V1553)

											tmp2605 := PrimHead(tmp2604)

											tmp2606 := PrimTail(tmp2605)

											tmp2607 := PrimTail(tmp2606)

											tmp2608 := PrimEqual(Nil, tmp2607)

											var ifres2590 Obj

											if True == tmp2608 {
												tmp2601 := PrimTail(V1553)

												tmp2602 := PrimTail(tmp2601)

												tmp2603 := PrimEqual(Nil, tmp2602)

												var ifres2592 Obj

												if True == tmp2603 {
													tmp2594 := PrimTail(V1553)

													tmp2595 := PrimHead(tmp2594)

													tmp2596 := PrimTail(tmp2595)

													tmp2597 := PrimHead(tmp2596)

													tmp2598 := PrimTail(tmp2597)

													tmp2599 := PrimHead(tmp2598)

													tmp2600 := PrimEqual(V1550, tmp2599)

													var ifres2593 Obj

													if True == tmp2600 {
														ifres2593 = True

													} else {
														ifres2593 = False

													}

													ifres2592 = ifres2593

												} else {
													ifres2592 = False

												}

												var ifres2591 Obj

												if True == ifres2592 {
													ifres2591 = True

												} else {
													ifres2591 = False

												}

												ifres2590 = ifres2591

											} else {
												ifres2590 = False

											}

											var ifres2589 Obj

											if True == ifres2590 {
												ifres2589 = True

											} else {
												ifres2589 = False

											}

											ifres2588 = ifres2589

										} else {
											ifres2588 = False

										}

										var ifres2587 Obj

										if True == ifres2588 {
											ifres2587 = True

										} else {
											ifres2587 = False

										}

										ifres2586 = ifres2587

									} else {
										ifres2586 = False

									}

									var ifres2585 Obj

									if True == ifres2586 {
										ifres2585 = True

									} else {
										ifres2585 = False

									}

									ifres2584 = ifres2585

								} else {
									ifres2584 = False

								}

								var ifres2583 Obj

								if True == ifres2584 {
									ifres2583 = True

								} else {
									ifres2583 = False

								}

								ifres2582 = ifres2583

							} else {
								ifres2582 = False

							}

							var ifres2581 Obj

							if True == ifres2582 {
								ifres2581 = True

							} else {
								ifres2581 = False

							}

							ifres2580 = ifres2581

						} else {
							ifres2580 = False

						}

						var ifres2579 Obj

						if True == ifres2580 {
							ifres2579 = True

						} else {
							ifres2579 = False

						}

						ifres2578 = ifres2579

					} else {
						ifres2578 = False

					}

					var ifres2577 Obj

					if True == ifres2578 {
						ifres2577 = True

					} else {
						ifres2577 = False

					}

					ifres2576 = ifres2577

				} else {
					ifres2576 = False

				}

				var ifres2575 Obj

				if True == ifres2576 {
					ifres2575 = True

				} else {
					ifres2575 = False

				}

				ifres2574 = ifres2575

			} else {
				ifres2574 = False

			}

			var ifres2573 Obj

			if True == ifres2574 {
				ifres2573 = True

			} else {
				ifres2573 = False

			}

			ifres2572 = ifres2573

		} else {
			ifres2572 = False

		}

		if True == ifres2572 {
			tmp2524 := PrimTail(V1553)

			tmp2525 := PrimHead(tmp2524)

			tmp2526 := PrimTail(tmp2525)

			tmp2527 := PrimHead(tmp2526)

			tmp2528 := PrimTail(tmp2527)

			tmp2529 := PrimHead(tmp2528)

			tmp2530 := Call(__e, PrimNS2Value(symshen_4scan_1body), tmp2529, V1554)

			tmp2531 := PrimCons(symcond, tmp2530)

			tmp2532 := PrimCons(tmp2531, Nil)

			tmp2533 := PrimCons(symfreeze, tmp2532)

			tmp2534 := PrimHead(V1553)

			tmp2535 := PrimTail(V1553)

			tmp2536 := PrimHead(tmp2535)

			tmp2537 := PrimTail(tmp2536)

			tmp2538 := PrimHead(tmp2537)

			tmp2539 := PrimTail(tmp2538)

			tmp2540 := PrimTail(tmp2539)

			tmp2541 := PrimHead(tmp2540)

			tmp2542 := PrimTail(V1553)

			tmp2543 := PrimHead(tmp2542)

			tmp2544 := PrimTail(tmp2543)

			tmp2545 := PrimHead(tmp2544)

			tmp2546 := PrimTail(tmp2545)

			tmp2547 := PrimHead(tmp2546)

			tmp2548 := PrimCons(V1552, Nil)

			tmp2549 := PrimCons(tmp2547, tmp2548)

			tmp2550 := PrimCons(V1551, Nil)

			tmp2551 := PrimCons(symthaw, tmp2550)

			tmp2552 := PrimCons(V1552, Nil)

			tmp2553 := PrimCons(tmp2551, tmp2552)

			tmp2554 := PrimCons(tmp2549, tmp2553)

			tmp2555 := PrimCons(symif, tmp2554)

			tmp2556 := PrimCons(tmp2555, Nil)

			tmp2557 := PrimCons(tmp2541, tmp2556)

			tmp2558 := PrimCons(V1552, tmp2557)

			tmp2559 := PrimCons(symlet, tmp2558)

			tmp2560 := PrimCons(V1551, Nil)

			tmp2561 := PrimCons(symthaw, tmp2560)

			tmp2562 := PrimCons(tmp2561, Nil)

			tmp2563 := PrimCons(tmp2559, tmp2562)

			tmp2564 := PrimCons(tmp2534, tmp2563)

			tmp2565 := PrimCons(symif, tmp2564)

			tmp2566 := PrimCons(tmp2565, Nil)

			tmp2567 := PrimCons(tmp2533, tmp2566)

			tmp2568 := PrimCons(V1551, tmp2567)

			tmp2569 := PrimCons(symlet, tmp2568)

			tmp2570 := PrimCons(tmp2569, Nil)

			tmp2571 := PrimCons(True, tmp2570)

			__e.Return(PrimCons(tmp2571, Nil))
			return

		} else {
			tmp2523 := PrimIsPair(V1553)

			var ifres2496 Obj

			if True == tmp2523 {
				tmp2521 := PrimTail(V1553)

				tmp2522 := PrimIsPair(tmp2521)

				var ifres2498 Obj

				if True == tmp2522 {
					tmp2518 := PrimTail(V1553)

					tmp2519 := PrimHead(tmp2518)

					tmp2520 := PrimIsPair(tmp2519)

					var ifres2500 Obj

					if True == tmp2520 {
						tmp2514 := PrimTail(V1553)

						tmp2515 := PrimHead(tmp2514)

						tmp2516 := PrimTail(tmp2515)

						tmp2517 := PrimIsPair(tmp2516)

						var ifres2502 Obj

						if True == tmp2517 {
							tmp2509 := PrimTail(V1553)

							tmp2510 := PrimHead(tmp2509)

							tmp2511 := PrimTail(tmp2510)

							tmp2512 := PrimTail(tmp2511)

							tmp2513 := PrimEqual(Nil, tmp2512)

							var ifres2504 Obj

							if True == tmp2513 {
								tmp2506 := PrimTail(V1553)

								tmp2507 := PrimTail(tmp2506)

								tmp2508 := PrimEqual(Nil, tmp2507)

								var ifres2505 Obj

								if True == tmp2508 {
									ifres2505 = True

								} else {
									ifres2505 = False

								}

								ifres2504 = ifres2505

							} else {
								ifres2504 = False

							}

							var ifres2503 Obj

							if True == ifres2504 {
								ifres2503 = True

							} else {
								ifres2503 = False

							}

							ifres2502 = ifres2503

						} else {
							ifres2502 = False

						}

						var ifres2501 Obj

						if True == ifres2502 {
							ifres2501 = True

						} else {
							ifres2501 = False

						}

						ifres2500 = ifres2501

					} else {
						ifres2500 = False

					}

					var ifres2499 Obj

					if True == ifres2500 {
						ifres2499 = True

					} else {
						ifres2499 = False

					}

					ifres2498 = ifres2499

				} else {
					ifres2498 = False

				}

				var ifres2497 Obj

				if True == ifres2498 {
					ifres2497 = True

				} else {
					ifres2497 = False

				}

				ifres2496 = ifres2497

			} else {
				ifres2496 = False

			}

			if True == ifres2496 {
				tmp2461 := Call(__e, PrimNS2Value(symshen_4scan_1body), V1550, V1554)

				tmp2462 := PrimCons(symcond, tmp2461)

				tmp2463 := PrimCons(tmp2462, Nil)

				tmp2464 := PrimCons(symfreeze, tmp2463)

				tmp2465 := PrimHead(V1553)

				tmp2466 := PrimTail(V1553)

				tmp2467 := PrimHead(tmp2466)

				tmp2468 := PrimTail(tmp2467)

				tmp2469 := PrimHead(tmp2468)

				tmp2470 := PrimCons(symfail, Nil)

				tmp2471 := PrimCons(tmp2470, Nil)

				tmp2472 := PrimCons(V1552, tmp2471)

				tmp2473 := PrimCons(sym_a, tmp2472)

				tmp2474 := PrimCons(V1551, Nil)

				tmp2475 := PrimCons(symthaw, tmp2474)

				tmp2476 := PrimCons(V1552, Nil)

				tmp2477 := PrimCons(tmp2475, tmp2476)

				tmp2478 := PrimCons(tmp2473, tmp2477)

				tmp2479 := PrimCons(symif, tmp2478)

				tmp2480 := PrimCons(tmp2479, Nil)

				tmp2481 := PrimCons(tmp2469, tmp2480)

				tmp2482 := PrimCons(V1552, tmp2481)

				tmp2483 := PrimCons(symlet, tmp2482)

				tmp2484 := PrimCons(V1551, Nil)

				tmp2485 := PrimCons(symthaw, tmp2484)

				tmp2486 := PrimCons(tmp2485, Nil)

				tmp2487 := PrimCons(tmp2483, tmp2486)

				tmp2488 := PrimCons(tmp2465, tmp2487)

				tmp2489 := PrimCons(symif, tmp2488)

				tmp2490 := PrimCons(tmp2489, Nil)

				tmp2491 := PrimCons(tmp2464, tmp2490)

				tmp2492 := PrimCons(V1551, tmp2491)

				tmp2493 := PrimCons(symlet, tmp2492)

				tmp2494 := PrimCons(tmp2493, Nil)

				tmp2495 := PrimCons(True, tmp2494)

				__e.Return(PrimCons(tmp2495, Nil))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.choicepoint")))
				return
			}

		}

	}, 5)

	tmp2651 := Call(__e, PrimNS2Value(symdef), symshen_4choicepoint, tmp2458)

	_ = tmp2651

	tmp2652 := MakeNative(func(__e *ControlFlow) {
		V1556 := __e.Get(1)
		_ = V1556
		V1557 := __e.Get(2)
		_ = V1557
		V1558 := __e.Get(3)
		_ = V1558
		tmp2666 := PrimEqual(V1556, V1558)

		if True == tmp2666 {
			__e.Return(V1557)
			return
		} else {
			tmp2665 := PrimIsPair(V1558)

			if True == tmp2665 {
				tmp2655 := MakeNative(func(__e *ControlFlow) {
					Rep := __e.Get(1)
					_ = Rep
					tmp2661 := PrimHead(V1558)

					tmp2662 := PrimEqual(Rep, tmp2661)

					if True == tmp2662 {
						tmp2658 := PrimHead(V1558)

						tmp2659 := PrimTail(V1558)

						tmp2660 := Call(__e, PrimNS2Value(symshen_4rep_1X), V1556, V1557, tmp2659)

						__e.Return(PrimCons(tmp2658, tmp2660))
						return

					} else {
						tmp2657 := PrimTail(V1558)

						__e.Return(PrimCons(Rep, tmp2657))
						return

					}

				}, 1)

				tmp2663 := PrimHead(V1558)

				tmp2664 := Call(__e, PrimNS2Value(symshen_4rep_1X), V1556, V1557, tmp2663)

				__e.TailApply(tmp2655, tmp2664)
				return

			} else {
				__e.Return(V1558)
				return
			}

		}

	}, 3)

	tmp2667 := Call(__e, PrimNS2Value(symdef), symshen_4rep_1X, tmp2652)

	_ = tmp2667

	tmp2668 := MakeNative(func(__e *ControlFlow) {
		V1559 := __e.Get(1)
		_ = V1559
		V1560 := __e.Get(2)
		_ = V1560
		tmp2669 := MakeNative(func(__e *ControlFlow) {
			R := __e.Get(1)
			_ = R
			tmp2670 := Call(__e, PrimNS2Value(symfst), R)

			tmp2671 := Call(__e, PrimNS2Value(symsnd), R)

			__e.TailApply(PrimNS2Value(symshen_4triple_1stack), Nil, tmp2670, V1560, tmp2671)
			return

		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp2669, V1559)
		return

	}, 2)

	tmp2672 := Call(__e, PrimNS2Value(symdef), symshen_4kl_1body, tmp2668)

	_ = tmp2672

	tmp2673 := MakeNative(func(__e *ControlFlow) {
		V1569 := __e.Get(1)
		_ = V1569
		V1570 := __e.Get(2)
		_ = V1570
		V1571 := __e.Get(3)
		_ = V1571
		V1572 := __e.Get(4)
		_ = V1572
		tmp2803 := PrimEqual(Nil, V1570)

		var ifres2778 Obj

		if True == tmp2803 {
			tmp2802 := PrimEqual(Nil, V1571)

			var ifres2780 Obj

			if True == tmp2802 {
				tmp2801 := PrimIsPair(V1572)

				var ifres2782 Obj

				if True == tmp2801 {
					tmp2799 := PrimHead(V1572)

					tmp2800 := PrimEqual(symwhere, tmp2799)

					var ifres2784 Obj

					if True == tmp2800 {
						tmp2797 := PrimTail(V1572)

						tmp2798 := PrimIsPair(tmp2797)

						var ifres2786 Obj

						if True == tmp2798 {
							tmp2794 := PrimTail(V1572)

							tmp2795 := PrimTail(tmp2794)

							tmp2796 := PrimIsPair(tmp2795)

							var ifres2788 Obj

							if True == tmp2796 {
								tmp2790 := PrimTail(V1572)

								tmp2791 := PrimTail(tmp2790)

								tmp2792 := PrimTail(tmp2791)

								tmp2793 := PrimEqual(Nil, tmp2792)

								var ifres2789 Obj

								if True == tmp2793 {
									ifres2789 = True

								} else {
									ifres2789 = False

								}

								ifres2788 = ifres2789

							} else {
								ifres2788 = False

							}

							var ifres2787 Obj

							if True == ifres2788 {
								ifres2787 = True

							} else {
								ifres2787 = False

							}

							ifres2786 = ifres2787

						} else {
							ifres2786 = False

						}

						var ifres2785 Obj

						if True == ifres2786 {
							ifres2785 = True

						} else {
							ifres2785 = False

						}

						ifres2784 = ifres2785

					} else {
						ifres2784 = False

					}

					var ifres2783 Obj

					if True == ifres2784 {
						ifres2783 = True

					} else {
						ifres2783 = False

					}

					ifres2782 = ifres2783

				} else {
					ifres2782 = False

				}

				var ifres2781 Obj

				if True == ifres2782 {
					ifres2781 = True

				} else {
					ifres2781 = False

				}

				ifres2780 = ifres2781

			} else {
				ifres2780 = False

			}

			var ifres2779 Obj

			if True == ifres2780 {
				ifres2779 = True

			} else {
				ifres2779 = False

			}

			ifres2778 = ifres2779

		} else {
			ifres2778 = False

		}

		if True == ifres2778 {
			tmp2772 := PrimTail(V1572)

			tmp2773 := PrimHead(tmp2772)

			tmp2774 := PrimCons(tmp2773, V1569)

			tmp2775 := PrimTail(V1572)

			tmp2776 := PrimTail(tmp2775)

			tmp2777 := PrimHead(tmp2776)

			__e.TailApply(PrimNS2Value(symshen_4triple_1stack), tmp2774, Nil, Nil, tmp2777)
			return

		} else {
			tmp2771 := PrimEqual(Nil, V1570)

			var ifres2768 Obj

			if True == tmp2771 {
				tmp2770 := PrimEqual(Nil, V1571)

				var ifres2769 Obj

				if True == tmp2770 {
					ifres2769 = True

				} else {
					ifres2769 = False

				}

				ifres2768 = ifres2769

			} else {
				ifres2768 = False

			}

			if True == ifres2768 {
				tmp2765 := Call(__e, PrimNS2Value(symreverse), V1569)

				tmp2766 := Call(__e, PrimNS2Value(symshen_4rectify_1test), tmp2765)

				tmp2767 := PrimCons(V1572, Nil)

				__e.Return(PrimCons(tmp2766, tmp2767))
				return

			} else {
				tmp2764 := PrimIsPair(V1570)

				var ifres2757 Obj

				if True == tmp2764 {
					tmp2763 := PrimIsPair(V1571)

					var ifres2759 Obj

					if True == tmp2763 {
						tmp2761 := PrimHead(V1570)

						tmp2762 := PrimIsVariable(tmp2761)

						var ifres2760 Obj

						if True == tmp2762 {
							ifres2760 = True

						} else {
							ifres2760 = False

						}

						ifres2759 = ifres2760

					} else {
						ifres2759 = False

					}

					var ifres2758 Obj

					if True == ifres2759 {
						ifres2758 = True

					} else {
						ifres2758 = False

					}

					ifres2757 = ifres2758

				} else {
					ifres2757 = False

				}

				if True == ifres2757 {
					tmp2752 := PrimTail(V1570)

					tmp2753 := PrimTail(V1571)

					tmp2754 := PrimHead(V1570)

					tmp2755 := PrimHead(V1571)

					tmp2756 := Call(__e, PrimNS2Value(symshen_4beta), tmp2754, tmp2755, V1572)

					__e.TailApply(PrimNS2Value(symshen_4triple_1stack), V1569, tmp2752, tmp2753, tmp2756)
					return

				} else {
					tmp2751 := PrimIsPair(V1570)

					var ifres2726 Obj

					if True == tmp2751 {
						tmp2749 := PrimHead(V1570)

						tmp2750 := PrimIsPair(tmp2749)

						var ifres2728 Obj

						if True == tmp2750 {
							tmp2746 := PrimHead(V1570)

							tmp2747 := PrimTail(tmp2746)

							tmp2748 := PrimIsPair(tmp2747)

							var ifres2730 Obj

							if True == tmp2748 {
								tmp2742 := PrimHead(V1570)

								tmp2743 := PrimTail(tmp2742)

								tmp2744 := PrimTail(tmp2743)

								tmp2745 := PrimIsPair(tmp2744)

								var ifres2732 Obj

								if True == tmp2745 {
									tmp2737 := PrimHead(V1570)

									tmp2738 := PrimTail(tmp2737)

									tmp2739 := PrimTail(tmp2738)

									tmp2740 := PrimTail(tmp2739)

									tmp2741 := PrimEqual(Nil, tmp2740)

									var ifres2734 Obj

									if True == tmp2741 {
										tmp2736 := PrimIsPair(V1571)

										var ifres2735 Obj

										if True == tmp2736 {
											ifres2735 = True

										} else {
											ifres2735 = False

										}

										ifres2734 = ifres2735

									} else {
										ifres2734 = False

									}

									var ifres2733 Obj

									if True == ifres2734 {
										ifres2733 = True

									} else {
										ifres2733 = False

									}

									ifres2732 = ifres2733

								} else {
									ifres2732 = False

								}

								var ifres2731 Obj

								if True == ifres2732 {
									ifres2731 = True

								} else {
									ifres2731 = False

								}

								ifres2730 = ifres2731

							} else {
								ifres2730 = False

							}

							var ifres2729 Obj

							if True == ifres2730 {
								ifres2729 = True

							} else {
								ifres2729 = False

							}

							ifres2728 = ifres2729

						} else {
							ifres2728 = False

						}

						var ifres2727 Obj

						if True == ifres2728 {
							ifres2727 = True

						} else {
							ifres2727 = False

						}

						ifres2726 = ifres2727

					} else {
						ifres2726 = False

					}

					if True == ifres2726 {
						tmp2691 := PrimHead(V1570)

						tmp2692 := PrimHead(tmp2691)

						tmp2693 := Call(__e, PrimNS2Value(symshen_4op_1test), tmp2692)

						tmp2694 := PrimHead(V1571)

						tmp2695 := PrimCons(tmp2694, Nil)

						tmp2696 := PrimCons(tmp2693, tmp2695)

						tmp2697 := PrimCons(tmp2696, V1569)

						tmp2698 := PrimHead(V1570)

						tmp2699 := PrimTail(tmp2698)

						tmp2700 := PrimHead(tmp2699)

						tmp2701 := PrimHead(V1570)

						tmp2702 := PrimTail(tmp2701)

						tmp2703 := PrimTail(tmp2702)

						tmp2704 := PrimHead(tmp2703)

						tmp2705 := PrimTail(V1570)

						tmp2706 := PrimCons(tmp2704, tmp2705)

						tmp2707 := PrimCons(tmp2700, tmp2706)

						tmp2708 := PrimHead(V1570)

						tmp2709 := PrimHead(tmp2708)

						tmp2710 := Call(__e, PrimNS2Value(symshen_4op1), tmp2709)

						tmp2711 := PrimHead(V1571)

						tmp2712 := PrimCons(tmp2711, Nil)

						tmp2713 := PrimCons(tmp2710, tmp2712)

						tmp2714 := PrimHead(V1570)

						tmp2715 := PrimHead(tmp2714)

						tmp2716 := Call(__e, PrimNS2Value(symshen_4op2), tmp2715)

						tmp2717 := PrimHead(V1571)

						tmp2718 := PrimCons(tmp2717, Nil)

						tmp2719 := PrimCons(tmp2716, tmp2718)

						tmp2720 := PrimTail(V1571)

						tmp2721 := PrimCons(tmp2719, tmp2720)

						tmp2722 := PrimCons(tmp2713, tmp2721)

						tmp2723 := PrimHead(V1570)

						tmp2724 := PrimHead(V1571)

						tmp2725 := Call(__e, PrimNS2Value(symshen_4beta), tmp2723, tmp2724, V1572)

						__e.TailApply(PrimNS2Value(symshen_4triple_1stack), tmp2697, tmp2707, tmp2722, tmp2725)
						return

					} else {
						tmp2690 := PrimIsPair(V1570)

						var ifres2687 Obj

						if True == tmp2690 {
							tmp2689 := PrimIsPair(V1571)

							var ifres2688 Obj

							if True == tmp2689 {
								ifres2688 = True

							} else {
								ifres2688 = False

							}

							ifres2687 = ifres2688

						} else {
							ifres2687 = False

						}

						if True == ifres2687 {
							tmp2679 := PrimHead(V1570)

							tmp2680 := PrimHead(V1571)

							tmp2681 := PrimCons(tmp2680, Nil)

							tmp2682 := PrimCons(tmp2679, tmp2681)

							tmp2683 := PrimCons(sym_a, tmp2682)

							tmp2684 := PrimCons(tmp2683, V1569)

							tmp2685 := PrimTail(V1570)

							tmp2686 := PrimTail(V1571)

							__e.TailApply(PrimNS2Value(symshen_4triple_1stack), tmp2684, tmp2685, tmp2686, V1572)
							return

						} else {
							__e.Return(PrimSimpleError(MakeString("implementation error in shen.triple-stack")))
							return
						}

					}

				}

			}

		}

	}, 4)

	tmp2804 := Call(__e, PrimNS2Value(symdef), symshen_4triple_1stack, tmp2673)

	_ = tmp2804

	tmp2805 := MakeNative(func(__e *ControlFlow) {
		V1575 := __e.Get(1)
		_ = V1575
		tmp2824 := PrimEqual(Nil, V1575)

		if True == tmp2824 {
			__e.Return(True)
			return
		} else {
			tmp2823 := PrimIsPair(V1575)

			var ifres2819 Obj

			if True == tmp2823 {
				tmp2821 := PrimTail(V1575)

				tmp2822 := PrimEqual(Nil, tmp2821)

				var ifres2820 Obj

				if True == tmp2822 {
					ifres2820 = True

				} else {
					ifres2820 = False

				}

				ifres2819 = ifres2820

			} else {
				ifres2819 = False

			}

			if True == ifres2819 {
				__e.Return(PrimHead(V1575))
				return
			} else {
				tmp2818 := PrimIsPair(V1575)

				var ifres2814 Obj

				if True == tmp2818 {
					tmp2816 := PrimTail(V1575)

					tmp2817 := PrimIsPair(tmp2816)

					var ifres2815 Obj

					if True == tmp2817 {
						ifres2815 = True

					} else {
						ifres2815 = False

					}

					ifres2814 = ifres2815

				} else {
					ifres2814 = False

				}

				if True == ifres2814 {
					tmp2809 := PrimHead(V1575)

					tmp2810 := PrimTail(V1575)

					tmp2811 := Call(__e, PrimNS2Value(symshen_4rectify_1test), tmp2810)

					tmp2812 := PrimCons(tmp2811, Nil)

					tmp2813 := PrimCons(tmp2809, tmp2812)

					__e.Return(PrimCons(symand, tmp2813))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.rectify-test")))
					return
				}

			}

		}

	}, 1)

	tmp2825 := Call(__e, PrimNS2Value(symdef), symshen_4rectify_1test, tmp2805)

	_ = tmp2825

	tmp2826 := MakeNative(func(__e *ControlFlow) {
		V1585 := __e.Get(1)
		_ = V1585
		V1586 := __e.Get(2)
		_ = V1586
		V1587 := __e.Get(3)
		_ = V1587
		tmp2903 := PrimEqual(V1585, V1587)

		if True == tmp2903 {
			__e.Return(V1586)
			return
		} else {
			tmp2902 := PrimIsPair(V1587)

			var ifres2878 Obj

			if True == tmp2902 {
				tmp2900 := PrimHead(V1587)

				tmp2901 := PrimEqual(symlambda, tmp2900)

				var ifres2880 Obj

				if True == tmp2901 {
					tmp2898 := PrimTail(V1587)

					tmp2899 := PrimIsPair(tmp2898)

					var ifres2882 Obj

					if True == tmp2899 {
						tmp2895 := PrimTail(V1587)

						tmp2896 := PrimTail(tmp2895)

						tmp2897 := PrimIsPair(tmp2896)

						var ifres2884 Obj

						if True == tmp2897 {
							tmp2891 := PrimTail(V1587)

							tmp2892 := PrimTail(tmp2891)

							tmp2893 := PrimTail(tmp2892)

							tmp2894 := PrimEqual(Nil, tmp2893)

							var ifres2886 Obj

							if True == tmp2894 {
								tmp2888 := PrimTail(V1587)

								tmp2889 := PrimHead(tmp2888)

								tmp2890 := PrimEqual(V1585, tmp2889)

								var ifres2887 Obj

								if True == tmp2890 {
									ifres2887 = True

								} else {
									ifres2887 = False

								}

								ifres2886 = ifres2887

							} else {
								ifres2886 = False

							}

							var ifres2885 Obj

							if True == ifres2886 {
								ifres2885 = True

							} else {
								ifres2885 = False

							}

							ifres2884 = ifres2885

						} else {
							ifres2884 = False

						}

						var ifres2883 Obj

						if True == ifres2884 {
							ifres2883 = True

						} else {
							ifres2883 = False

						}

						ifres2882 = ifres2883

					} else {
						ifres2882 = False

					}

					var ifres2881 Obj

					if True == ifres2882 {
						ifres2881 = True

					} else {
						ifres2881 = False

					}

					ifres2880 = ifres2881

				} else {
					ifres2880 = False

				}

				var ifres2879 Obj

				if True == ifres2880 {
					ifres2879 = True

				} else {
					ifres2879 = False

				}

				ifres2878 = ifres2879

			} else {
				ifres2878 = False

			}

			if True == ifres2878 {
				__e.Return(V1587)
				return
			} else {
				tmp2877 := PrimIsPair(V1587)

				var ifres2846 Obj

				if True == tmp2877 {
					tmp2875 := PrimHead(V1587)

					tmp2876 := PrimEqual(symlet, tmp2875)

					var ifres2848 Obj

					if True == tmp2876 {
						tmp2873 := PrimTail(V1587)

						tmp2874 := PrimIsPair(tmp2873)

						var ifres2850 Obj

						if True == tmp2874 {
							tmp2870 := PrimTail(V1587)

							tmp2871 := PrimTail(tmp2870)

							tmp2872 := PrimIsPair(tmp2871)

							var ifres2852 Obj

							if True == tmp2872 {
								tmp2866 := PrimTail(V1587)

								tmp2867 := PrimTail(tmp2866)

								tmp2868 := PrimTail(tmp2867)

								tmp2869 := PrimIsPair(tmp2868)

								var ifres2854 Obj

								if True == tmp2869 {
									tmp2861 := PrimTail(V1587)

									tmp2862 := PrimTail(tmp2861)

									tmp2863 := PrimTail(tmp2862)

									tmp2864 := PrimTail(tmp2863)

									tmp2865 := PrimEqual(Nil, tmp2864)

									var ifres2856 Obj

									if True == tmp2865 {
										tmp2858 := PrimTail(V1587)

										tmp2859 := PrimHead(tmp2858)

										tmp2860 := PrimEqual(V1585, tmp2859)

										var ifres2857 Obj

										if True == tmp2860 {
											ifres2857 = True

										} else {
											ifres2857 = False

										}

										ifres2856 = ifres2857

									} else {
										ifres2856 = False

									}

									var ifres2855 Obj

									if True == ifres2856 {
										ifres2855 = True

									} else {
										ifres2855 = False

									}

									ifres2854 = ifres2855

								} else {
									ifres2854 = False

								}

								var ifres2853 Obj

								if True == ifres2854 {
									ifres2853 = True

								} else {
									ifres2853 = False

								}

								ifres2852 = ifres2853

							} else {
								ifres2852 = False

							}

							var ifres2851 Obj

							if True == ifres2852 {
								ifres2851 = True

							} else {
								ifres2851 = False

							}

							ifres2850 = ifres2851

						} else {
							ifres2850 = False

						}

						var ifres2849 Obj

						if True == ifres2850 {
							ifres2849 = True

						} else {
							ifres2849 = False

						}

						ifres2848 = ifres2849

					} else {
						ifres2848 = False

					}

					var ifres2847 Obj

					if True == ifres2848 {
						ifres2847 = True

					} else {
						ifres2847 = False

					}

					ifres2846 = ifres2847

				} else {
					ifres2846 = False

				}

				if True == ifres2846 {
					tmp2833 := PrimTail(V1587)

					tmp2834 := PrimHead(tmp2833)

					tmp2835 := PrimTail(V1587)

					tmp2836 := PrimHead(tmp2835)

					tmp2837 := PrimTail(V1587)

					tmp2838 := PrimTail(tmp2837)

					tmp2839 := PrimHead(tmp2838)

					tmp2840 := Call(__e, PrimNS2Value(symshen_4beta), tmp2836, V1586, tmp2839)

					tmp2841 := PrimTail(V1587)

					tmp2842 := PrimTail(tmp2841)

					tmp2843 := PrimTail(tmp2842)

					tmp2844 := PrimCons(tmp2840, tmp2843)

					tmp2845 := PrimCons(tmp2834, tmp2844)

					__e.Return(PrimCons(symlet, tmp2845))
					return

				} else {
					tmp2832 := PrimIsPair(V1587)

					if True == tmp2832 {
						tmp2831 := MakeNative(func(__e *ControlFlow) {
							V := __e.Get(1)
							_ = V
							__e.TailApply(PrimNS2Value(symshen_4beta), V1585, V1586, V)
							return
						}, 1)

						__e.TailApply(PrimNS2Value(symmap), tmp2831, V1587)
						return

					} else {
						__e.Return(V1587)
						return
					}

				}

			}

		}

	}, 3)

	tmp2904 := Call(__e, PrimNS2Value(symdef), symshen_4beta, tmp2826)

	_ = tmp2904

	tmp2905 := MakeNative(func(__e *ControlFlow) {
		V1590 := __e.Get(1)
		_ = V1590
		tmp2913 := PrimEqual(symcons, V1590)

		if True == tmp2913 {
			__e.Return(symhd)
			return
		} else {
			tmp2912 := PrimEqual(sym_8s, V1590)

			if True == tmp2912 {
				__e.Return(symhdstr)
				return
			} else {
				tmp2911 := PrimEqual(sym_8p, V1590)

				if True == tmp2911 {
					__e.Return(symfst)
					return
				} else {
					tmp2910 := PrimEqual(sym_8v, V1590)

					if True == tmp2910 {
						__e.Return(symhdv)
						return
					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.op1")))
						return
					}

				}

			}

		}

	}, 1)

	tmp2914 := Call(__e, PrimNS2Value(symdef), symshen_4op1, tmp2905)

	_ = tmp2914

	tmp2915 := MakeNative(func(__e *ControlFlow) {
		V1593 := __e.Get(1)
		_ = V1593
		tmp2923 := PrimEqual(symcons, V1593)

		if True == tmp2923 {
			__e.Return(symtl)
			return
		} else {
			tmp2922 := PrimEqual(sym_8s, V1593)

			if True == tmp2922 {
				__e.Return(symtlstr)
				return
			} else {
				tmp2921 := PrimEqual(sym_8p, V1593)

				if True == tmp2921 {
					__e.Return(symsnd)
					return
				} else {
					tmp2920 := PrimEqual(sym_8v, V1593)

					if True == tmp2920 {
						__e.Return(symtlv)
						return
					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.op2")))
						return
					}

				}

			}

		}

	}, 1)

	tmp2924 := Call(__e, PrimNS2Value(symdef), symshen_4op2, tmp2915)

	_ = tmp2924

	tmp2925 := MakeNative(func(__e *ControlFlow) {
		V1596 := __e.Get(1)
		_ = V1596
		tmp2933 := PrimEqual(symcons, V1596)

		if True == tmp2933 {
			__e.Return(symcons_2)
			return
		} else {
			tmp2932 := PrimEqual(sym_8s, V1596)

			if True == tmp2932 {
				__e.Return(symshen_4_7string_2)
				return
			} else {
				tmp2931 := PrimEqual(sym_8p, V1596)

				if True == tmp2931 {
					__e.Return(symtuple_2)
					return
				} else {
					tmp2930 := PrimEqual(sym_8v, V1596)

					if True == tmp2930 {
						__e.Return(symshen_4_7vector_2)
						return
					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.op-test")))
						return
					}

				}

			}

		}

	}, 1)

	tmp2934 := Call(__e, PrimNS2Value(symdef), symshen_4op_1test, tmp2925)

	_ = tmp2934

	tmp2935 := MakeNative(func(__e *ControlFlow) {
		V1597 := __e.Get(1)
		_ = V1597
		tmp2937 := PrimEqual(MakeString(""), V1597)

		if True == tmp2937 {
			__e.Return(False)
			return
		} else {
			__e.Return(PrimIsString(V1597))
			return
		}

	}, 1)

	tmp2938 := Call(__e, PrimNS2Value(symdef), symshen_4_7string_2, tmp2935)

	_ = tmp2938

	tmp2939 := MakeNative(func(__e *ControlFlow) {
		V1598 := __e.Get(1)
		_ = V1598
		tmp2941 := Call(__e, PrimNS2Value(symvector), MakeNumber(0))

		tmp2942 := PrimEqual(V1598, tmp2941)

		if True == tmp2942 {
			__e.Return(False)
			return
		} else {
			__e.TailApply(PrimNS2Value(symvector_2), V1598)
			return
		}

	}, 1)

	tmp2943 := Call(__e, PrimNS2Value(symdef), symshen_4_7vector_2, tmp2939)

	_ = tmp2943

	tmp2944 := MakeNative(func(__e *ControlFlow) {
		V1601 := __e.Get(1)
		_ = V1601
		tmp2948 := PrimEqual(sym_7, V1601)

		if True == tmp2948 {
			__e.Return(PrimNS3Set(symshen_4_dfactorise_2_d, True))
			return
		} else {
			tmp2947 := PrimEqual(sym_1, V1601)

			if True == tmp2947 {
				__e.Return(PrimNS3Set(symshen_4_dfactorise_2_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("factorise expects a + or a -\n")))
				return
			}

		}

	}, 1)

	tmp2949 := Call(__e, PrimNS2Value(symdef), symfactorise, tmp2944)

	_ = tmp2949

	tmp2950 := MakeNative(func(__e *ControlFlow) {
		V1602 := __e.Get(1)
		_ = V1602
		tmp3017 := PrimIsPair(V1602)

		var ifres2973 Obj

		if True == tmp3017 {
			tmp3015 := PrimHead(V1602)

			tmp3016 := PrimEqual(symdefun, tmp3015)

			var ifres2975 Obj

			if True == tmp3016 {
				tmp3013 := PrimTail(V1602)

				tmp3014 := PrimIsPair(tmp3013)

				var ifres2977 Obj

				if True == tmp3014 {
					tmp3010 := PrimTail(V1602)

					tmp3011 := PrimTail(tmp3010)

					tmp3012 := PrimIsPair(tmp3011)

					var ifres2979 Obj

					if True == tmp3012 {
						tmp3006 := PrimTail(V1602)

						tmp3007 := PrimTail(tmp3006)

						tmp3008 := PrimTail(tmp3007)

						tmp3009 := PrimIsPair(tmp3008)

						var ifres2981 Obj

						if True == tmp3009 {
							tmp3001 := PrimTail(V1602)

							tmp3002 := PrimTail(tmp3001)

							tmp3003 := PrimTail(tmp3002)

							tmp3004 := PrimHead(tmp3003)

							tmp3005 := PrimIsPair(tmp3004)

							var ifres2983 Obj

							if True == tmp3005 {
								tmp2995 := PrimTail(V1602)

								tmp2996 := PrimTail(tmp2995)

								tmp2997 := PrimTail(tmp2996)

								tmp2998 := PrimHead(tmp2997)

								tmp2999 := PrimHead(tmp2998)

								tmp3000 := PrimEqual(symcond, tmp2999)

								var ifres2985 Obj

								if True == tmp3000 {
									tmp2990 := PrimTail(V1602)

									tmp2991 := PrimTail(tmp2990)

									tmp2992 := PrimTail(tmp2991)

									tmp2993 := PrimTail(tmp2992)

									tmp2994 := PrimEqual(Nil, tmp2993)

									var ifres2987 Obj

									if True == tmp2994 {
										tmp2989 := PrimNS3Value(symshen_4_dfactorise_2_d)

										var ifres2988 Obj

										if True == tmp2989 {
											ifres2988 = True

										} else {
											ifres2988 = False

										}

										ifres2987 = ifres2988

									} else {
										ifres2987 = False

									}

									var ifres2986 Obj

									if True == ifres2987 {
										ifres2986 = True

									} else {
										ifres2986 = False

									}

									ifres2985 = ifres2986

								} else {
									ifres2985 = False

								}

								var ifres2984 Obj

								if True == ifres2985 {
									ifres2984 = True

								} else {
									ifres2984 = False

								}

								ifres2983 = ifres2984

							} else {
								ifres2983 = False

							}

							var ifres2982 Obj

							if True == ifres2983 {
								ifres2982 = True

							} else {
								ifres2982 = False

							}

							ifres2981 = ifres2982

						} else {
							ifres2981 = False

						}

						var ifres2980 Obj

						if True == ifres2981 {
							ifres2980 = True

						} else {
							ifres2980 = False

						}

						ifres2979 = ifres2980

					} else {
						ifres2979 = False

					}

					var ifres2978 Obj

					if True == ifres2979 {
						ifres2978 = True

					} else {
						ifres2978 = False

					}

					ifres2977 = ifres2978

				} else {
					ifres2977 = False

				}

				var ifres2976 Obj

				if True == ifres2977 {
					ifres2976 = True

				} else {
					ifres2976 = False

				}

				ifres2975 = ifres2976

			} else {
				ifres2975 = False

			}

			var ifres2974 Obj

			if True == ifres2975 {
				ifres2974 = True

			} else {
				ifres2974 = False

			}

			ifres2973 = ifres2974

		} else {
			ifres2973 = False

		}

		if True == ifres2973 {
			tmp2952 := PrimTail(V1602)

			tmp2953 := PrimHead(tmp2952)

			tmp2954 := PrimTail(V1602)

			tmp2955 := PrimTail(tmp2954)

			tmp2956 := PrimHead(tmp2955)

			tmp2957 := PrimTail(V1602)

			tmp2958 := PrimTail(tmp2957)

			tmp2959 := PrimHead(tmp2958)

			tmp2960 := PrimTail(V1602)

			tmp2961 := PrimTail(tmp2960)

			tmp2962 := PrimTail(tmp2961)

			tmp2963 := PrimHead(tmp2962)

			tmp2964 := PrimTail(tmp2963)

			tmp2965 := PrimTail(V1602)

			tmp2966 := PrimHead(tmp2965)

			tmp2967 := PrimCons(tmp2966, Nil)

			tmp2968 := PrimCons(symshen_4f_1error, tmp2967)

			tmp2969 := Call(__e, PrimNS2Value(symshen_4vertical), tmp2959, tmp2964, tmp2968)

			tmp2970 := PrimCons(tmp2969, Nil)

			tmp2971 := PrimCons(tmp2956, tmp2970)

			tmp2972 := PrimCons(tmp2953, tmp2971)

			__e.Return(PrimCons(symdefun, tmp2972))
			return

		} else {
			__e.Return(V1602)
			return
		}

	}, 1)

	tmp3018 := Call(__e, PrimNS2Value(symdef), symshen_4factorise_1code, tmp2950)

	_ = tmp3018

	tmp3019 := MakeNative(func(__e *ControlFlow) {
		V1615 := __e.Get(1)
		_ = V1615
		V1616 := __e.Get(2)
		_ = V1616
		V1617 := __e.Get(3)
		_ = V1617
		tmp3131 := PrimIsPair(V1616)

		var ifres3111 Obj

		if True == tmp3131 {
			tmp3129 := PrimHead(V1616)

			tmp3130 := PrimIsPair(tmp3129)

			var ifres3113 Obj

			if True == tmp3130 {
				tmp3126 := PrimHead(V1616)

				tmp3127 := PrimHead(tmp3126)

				tmp3128 := PrimEqual(True, tmp3127)

				var ifres3115 Obj

				if True == tmp3128 {
					tmp3123 := PrimHead(V1616)

					tmp3124 := PrimTail(tmp3123)

					tmp3125 := PrimIsPair(tmp3124)

					var ifres3117 Obj

					if True == tmp3125 {
						tmp3119 := PrimHead(V1616)

						tmp3120 := PrimTail(tmp3119)

						tmp3121 := PrimTail(tmp3120)

						tmp3122 := PrimEqual(Nil, tmp3121)

						var ifres3118 Obj

						if True == tmp3122 {
							ifres3118 = True

						} else {
							ifres3118 = False

						}

						ifres3117 = ifres3118

					} else {
						ifres3117 = False

					}

					var ifres3116 Obj

					if True == ifres3117 {
						ifres3116 = True

					} else {
						ifres3116 = False

					}

					ifres3115 = ifres3116

				} else {
					ifres3115 = False

				}

				var ifres3114 Obj

				if True == ifres3115 {
					ifres3114 = True

				} else {
					ifres3114 = False

				}

				ifres3113 = ifres3114

			} else {
				ifres3113 = False

			}

			var ifres3112 Obj

			if True == ifres3113 {
				ifres3112 = True

			} else {
				ifres3112 = False

			}

			ifres3111 = ifres3112

		} else {
			ifres3111 = False

		}

		if True == ifres3111 {
			tmp3109 := PrimHead(V1616)

			tmp3110 := PrimTail(tmp3109)

			__e.Return(PrimHead(tmp3110))
			return

		} else {
			tmp3108 := PrimEqual(Nil, V1616)

			if True == tmp3108 {
				__e.Return(V1617)
				return
			} else {
				tmp3107 := PrimIsPair(V1616)

				var ifres3060 Obj

				if True == tmp3107 {
					tmp3105 := PrimHead(V1616)

					tmp3106 := PrimIsPair(tmp3105)

					var ifres3062 Obj

					if True == tmp3106 {
						tmp3102 := PrimHead(V1616)

						tmp3103 := PrimHead(tmp3102)

						tmp3104 := PrimIsPair(tmp3103)

						var ifres3064 Obj

						if True == tmp3104 {
							tmp3098 := PrimHead(V1616)

							tmp3099 := PrimHead(tmp3098)

							tmp3100 := PrimHead(tmp3099)

							tmp3101 := PrimEqual(symand, tmp3100)

							var ifres3066 Obj

							if True == tmp3101 {
								tmp3094 := PrimHead(V1616)

								tmp3095 := PrimHead(tmp3094)

								tmp3096 := PrimTail(tmp3095)

								tmp3097 := PrimIsPair(tmp3096)

								var ifres3068 Obj

								if True == tmp3097 {
									tmp3089 := PrimHead(V1616)

									tmp3090 := PrimHead(tmp3089)

									tmp3091 := PrimTail(tmp3090)

									tmp3092 := PrimTail(tmp3091)

									tmp3093 := PrimIsPair(tmp3092)

									var ifres3070 Obj

									if True == tmp3093 {
										tmp3083 := PrimHead(V1616)

										tmp3084 := PrimHead(tmp3083)

										tmp3085 := PrimTail(tmp3084)

										tmp3086 := PrimTail(tmp3085)

										tmp3087 := PrimTail(tmp3086)

										tmp3088 := PrimEqual(Nil, tmp3087)

										var ifres3072 Obj

										if True == tmp3088 {
											tmp3080 := PrimHead(V1616)

											tmp3081 := PrimTail(tmp3080)

											tmp3082 := PrimIsPair(tmp3081)

											var ifres3074 Obj

											if True == tmp3082 {
												tmp3076 := PrimHead(V1616)

												tmp3077 := PrimTail(tmp3076)

												tmp3078 := PrimTail(tmp3077)

												tmp3079 := PrimEqual(Nil, tmp3078)

												var ifres3075 Obj

												if True == tmp3079 {
													ifres3075 = True

												} else {
													ifres3075 = False

												}

												ifres3074 = ifres3075

											} else {
												ifres3074 = False

											}

											var ifres3073 Obj

											if True == ifres3074 {
												ifres3073 = True

											} else {
												ifres3073 = False

											}

											ifres3072 = ifres3073

										} else {
											ifres3072 = False

										}

										var ifres3071 Obj

										if True == ifres3072 {
											ifres3071 = True

										} else {
											ifres3071 = False

										}

										ifres3070 = ifres3071

									} else {
										ifres3070 = False

									}

									var ifres3069 Obj

									if True == ifres3070 {
										ifres3069 = True

									} else {
										ifres3069 = False

									}

									ifres3068 = ifres3069

								} else {
									ifres3068 = False

								}

								var ifres3067 Obj

								if True == ifres3068 {
									ifres3067 = True

								} else {
									ifres3067 = False

								}

								ifres3066 = ifres3067

							} else {
								ifres3066 = False

							}

							var ifres3065 Obj

							if True == ifres3066 {
								ifres3065 = True

							} else {
								ifres3065 = False

							}

							ifres3064 = ifres3065

						} else {
							ifres3064 = False

						}

						var ifres3063 Obj

						if True == ifres3064 {
							ifres3063 = True

						} else {
							ifres3063 = False

						}

						ifres3062 = ifres3063

					} else {
						ifres3062 = False

					}

					var ifres3061 Obj

					if True == ifres3062 {
						ifres3061 = True

					} else {
						ifres3061 = False

					}

					ifres3060 = ifres3061

				} else {
					ifres3060 = False

				}

				if True == ifres3060 {
					tmp3050 := MakeNative(func(__e *ControlFlow) {
						Before_7After := __e.Get(1)
						_ = Before_7After
						tmp3051 := PrimHead(V1616)

						tmp3052 := PrimHead(tmp3051)

						tmp3053 := PrimTail(tmp3052)

						tmp3054 := PrimHead(tmp3053)

						__e.TailApply(PrimNS2Value(symshen_4branch), tmp3054, V1615, Before_7After, V1617)
						return

					}, 1)

					tmp3055 := PrimHead(V1616)

					tmp3056 := PrimHead(tmp3055)

					tmp3057 := PrimTail(tmp3056)

					tmp3058 := PrimHead(tmp3057)

					tmp3059 := Call(__e, PrimNS2Value(symshen_4split_1cases), tmp3058, V1616, Nil)

					__e.TailApply(tmp3050, tmp3059)
					return

				} else {
					tmp3049 := PrimIsPair(V1616)

					var ifres3034 Obj

					if True == tmp3049 {
						tmp3047 := PrimHead(V1616)

						tmp3048 := PrimIsPair(tmp3047)

						var ifres3036 Obj

						if True == tmp3048 {
							tmp3044 := PrimHead(V1616)

							tmp3045 := PrimTail(tmp3044)

							tmp3046 := PrimIsPair(tmp3045)

							var ifres3038 Obj

							if True == tmp3046 {
								tmp3040 := PrimHead(V1616)

								tmp3041 := PrimTail(tmp3040)

								tmp3042 := PrimTail(tmp3041)

								tmp3043 := PrimEqual(Nil, tmp3042)

								var ifres3039 Obj

								if True == tmp3043 {
									ifres3039 = True

								} else {
									ifres3039 = False

								}

								ifres3038 = ifres3039

							} else {
								ifres3038 = False

							}

							var ifres3037 Obj

							if True == ifres3038 {
								ifres3037 = True

							} else {
								ifres3037 = False

							}

							ifres3036 = ifres3037

						} else {
							ifres3036 = False

						}

						var ifres3035 Obj

						if True == ifres3036 {
							ifres3035 = True

						} else {
							ifres3035 = False

						}

						ifres3034 = ifres3035

					} else {
						ifres3034 = False

					}

					if True == ifres3034 {
						tmp3024 := PrimHead(V1616)

						tmp3025 := PrimHead(tmp3024)

						tmp3026 := PrimHead(V1616)

						tmp3027 := PrimTail(tmp3026)

						tmp3028 := PrimHead(tmp3027)

						tmp3029 := PrimTail(V1616)

						tmp3030 := Call(__e, PrimNS2Value(symshen_4vertical), V1615, tmp3029, V1617)

						tmp3031 := PrimCons(tmp3030, Nil)

						tmp3032 := PrimCons(tmp3028, tmp3031)

						tmp3033 := PrimCons(tmp3025, tmp3032)

						__e.Return(PrimCons(symif, tmp3033))
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.vertical")))
						return
					}

				}

			}

		}

	}, 3)

	tmp3132 := Call(__e, PrimNS2Value(symdef), symshen_4vertical, tmp3019)

	_ = tmp3132

	tmp3133 := MakeNative(func(__e *ControlFlow) {
		V1622 := __e.Get(1)
		_ = V1622
		V1623 := __e.Get(2)
		_ = V1623
		V1624 := __e.Get(3)
		_ = V1624
		tmp3234 := PrimIsPair(V1623)

		var ifres3180 Obj

		if True == tmp3234 {
			tmp3232 := PrimHead(V1623)

			tmp3233 := PrimIsPair(tmp3232)

			var ifres3182 Obj

			if True == tmp3233 {
				tmp3229 := PrimHead(V1623)

				tmp3230 := PrimHead(tmp3229)

				tmp3231 := PrimIsPair(tmp3230)

				var ifres3184 Obj

				if True == tmp3231 {
					tmp3225 := PrimHead(V1623)

					tmp3226 := PrimHead(tmp3225)

					tmp3227 := PrimHead(tmp3226)

					tmp3228 := PrimEqual(symand, tmp3227)

					var ifres3186 Obj

					if True == tmp3228 {
						tmp3221 := PrimHead(V1623)

						tmp3222 := PrimHead(tmp3221)

						tmp3223 := PrimTail(tmp3222)

						tmp3224 := PrimIsPair(tmp3223)

						var ifres3188 Obj

						if True == tmp3224 {
							tmp3216 := PrimHead(V1623)

							tmp3217 := PrimHead(tmp3216)

							tmp3218 := PrimTail(tmp3217)

							tmp3219 := PrimTail(tmp3218)

							tmp3220 := PrimIsPair(tmp3219)

							var ifres3190 Obj

							if True == tmp3220 {
								tmp3210 := PrimHead(V1623)

								tmp3211 := PrimHead(tmp3210)

								tmp3212 := PrimTail(tmp3211)

								tmp3213 := PrimTail(tmp3212)

								tmp3214 := PrimTail(tmp3213)

								tmp3215 := PrimEqual(Nil, tmp3214)

								var ifres3192 Obj

								if True == tmp3215 {
									tmp3207 := PrimHead(V1623)

									tmp3208 := PrimTail(tmp3207)

									tmp3209 := PrimIsPair(tmp3208)

									var ifres3194 Obj

									if True == tmp3209 {
										tmp3203 := PrimHead(V1623)

										tmp3204 := PrimTail(tmp3203)

										tmp3205 := PrimTail(tmp3204)

										tmp3206 := PrimEqual(Nil, tmp3205)

										var ifres3196 Obj

										if True == tmp3206 {
											tmp3198 := PrimHead(V1623)

											tmp3199 := PrimHead(tmp3198)

											tmp3200 := PrimTail(tmp3199)

											tmp3201 := PrimHead(tmp3200)

											tmp3202 := PrimEqual(V1622, tmp3201)

											var ifres3197 Obj

											if True == tmp3202 {
												ifres3197 = True

											} else {
												ifres3197 = False

											}

											ifres3196 = ifres3197

										} else {
											ifres3196 = False

										}

										var ifres3195 Obj

										if True == ifres3196 {
											ifres3195 = True

										} else {
											ifres3195 = False

										}

										ifres3194 = ifres3195

									} else {
										ifres3194 = False

									}

									var ifres3193 Obj

									if True == ifres3194 {
										ifres3193 = True

									} else {
										ifres3193 = False

									}

									ifres3192 = ifres3193

								} else {
									ifres3192 = False

								}

								var ifres3191 Obj

								if True == ifres3192 {
									ifres3191 = True

								} else {
									ifres3191 = False

								}

								ifres3190 = ifres3191

							} else {
								ifres3190 = False

							}

							var ifres3189 Obj

							if True == ifres3190 {
								ifres3189 = True

							} else {
								ifres3189 = False

							}

							ifres3188 = ifres3189

						} else {
							ifres3188 = False

						}

						var ifres3187 Obj

						if True == ifres3188 {
							ifres3187 = True

						} else {
							ifres3187 = False

						}

						ifres3186 = ifres3187

					} else {
						ifres3186 = False

					}

					var ifres3185 Obj

					if True == ifres3186 {
						ifres3185 = True

					} else {
						ifres3185 = False

					}

					ifres3184 = ifres3185

				} else {
					ifres3184 = False

				}

				var ifres3183 Obj

				if True == ifres3184 {
					ifres3183 = True

				} else {
					ifres3183 = False

				}

				ifres3182 = ifres3183

			} else {
				ifres3182 = False

			}

			var ifres3181 Obj

			if True == ifres3182 {
				ifres3181 = True

			} else {
				ifres3181 = False

			}

			ifres3180 = ifres3181

		} else {
			ifres3180 = False

		}

		if True == ifres3180 {
			tmp3166 := PrimHead(V1623)

			tmp3167 := PrimHead(tmp3166)

			tmp3168 := PrimTail(tmp3167)

			tmp3169 := PrimHead(tmp3168)

			tmp3170 := PrimTail(V1623)

			tmp3171 := PrimHead(V1623)

			tmp3172 := PrimHead(tmp3171)

			tmp3173 := PrimTail(tmp3172)

			tmp3174 := PrimTail(tmp3173)

			tmp3175 := PrimHead(tmp3174)

			tmp3176 := PrimHead(V1623)

			tmp3177 := PrimTail(tmp3176)

			tmp3178 := PrimCons(tmp3175, tmp3177)

			tmp3179 := PrimCons(tmp3178, V1624)

			__e.TailApply(PrimNS2Value(symshen_4split_1cases), tmp3169, tmp3170, tmp3179)
			return

		} else {
			tmp3165 := PrimIsPair(V1623)

			var ifres3145 Obj

			if True == tmp3165 {
				tmp3163 := PrimHead(V1623)

				tmp3164 := PrimIsPair(tmp3163)

				var ifres3147 Obj

				if True == tmp3164 {
					tmp3160 := PrimHead(V1623)

					tmp3161 := PrimTail(tmp3160)

					tmp3162 := PrimIsPair(tmp3161)

					var ifres3149 Obj

					if True == tmp3162 {
						tmp3156 := PrimHead(V1623)

						tmp3157 := PrimTail(tmp3156)

						tmp3158 := PrimTail(tmp3157)

						tmp3159 := PrimEqual(Nil, tmp3158)

						var ifres3151 Obj

						if True == tmp3159 {
							tmp3153 := PrimHead(V1623)

							tmp3154 := PrimHead(tmp3153)

							tmp3155 := PrimEqual(V1622, tmp3154)

							var ifres3152 Obj

							if True == tmp3155 {
								ifres3152 = True

							} else {
								ifres3152 = False

							}

							ifres3151 = ifres3152

						} else {
							ifres3151 = False

						}

						var ifres3150 Obj

						if True == ifres3151 {
							ifres3150 = True

						} else {
							ifres3150 = False

						}

						ifres3149 = ifres3150

					} else {
						ifres3149 = False

					}

					var ifres3148 Obj

					if True == ifres3149 {
						ifres3148 = True

					} else {
						ifres3148 = False

					}

					ifres3147 = ifres3148

				} else {
					ifres3147 = False

				}

				var ifres3146 Obj

				if True == ifres3147 {
					ifres3146 = True

				} else {
					ifres3146 = False

				}

				ifres3145 = ifres3146

			} else {
				ifres3145 = False

			}

			if True == ifres3145 {
				tmp3138 := PrimHead(V1623)

				tmp3139 := PrimHead(tmp3138)

				tmp3140 := PrimTail(V1623)

				tmp3141 := PrimHead(V1623)

				tmp3142 := PrimTail(tmp3141)

				tmp3143 := PrimCons(True, tmp3142)

				tmp3144 := PrimCons(tmp3143, V1624)

				__e.TailApply(PrimNS2Value(symshen_4split_1cases), tmp3139, tmp3140, tmp3144)
				return

			} else {
				tmp3136 := Call(__e, PrimNS2Value(symreverse), V1624)

				tmp3137 := PrimCons(V1623, Nil)

				__e.Return(PrimCons(tmp3136, tmp3137))
				return

			}

		}

	}, 3)

	tmp3235 := Call(__e, PrimNS2Value(symdef), symshen_4split_1cases, tmp3133)

	_ = tmp3235

	tmp3236 := MakeNative(func(__e *ControlFlow) {
		V1625 := __e.Get(1)
		_ = V1625
		V1626 := __e.Get(2)
		_ = V1626
		V1627 := __e.Get(3)
		_ = V1627
		V1628 := __e.Get(4)
		_ = V1628
		tmp3257 := PrimIsPair(V1627)

		var ifres3248 Obj

		if True == tmp3257 {
			tmp3255 := PrimTail(V1627)

			tmp3256 := PrimIsPair(tmp3255)

			var ifres3250 Obj

			if True == tmp3256 {
				tmp3252 := PrimTail(V1627)

				tmp3253 := PrimTail(tmp3252)

				tmp3254 := PrimEqual(Nil, tmp3253)

				var ifres3251 Obj

				if True == tmp3254 {
					ifres3251 = True

				} else {
					ifres3251 = False

				}

				ifres3250 = ifres3251

			} else {
				ifres3250 = False

			}

			var ifres3249 Obj

			if True == ifres3250 {
				ifres3249 = True

			} else {
				ifres3249 = False

			}

			ifres3248 = ifres3249

		} else {
			ifres3248 = False

		}

		if True == ifres3248 {
			tmp3238 := MakeNative(func(__e *ControlFlow) {
				Else := __e.Get(1)
				_ = Else
				tmp3239 := MakeNative(func(__e *ControlFlow) {
					Then := __e.Get(1)
					_ = Then
					tmp3240 := PrimCons(Else, Nil)

					tmp3241 := PrimCons(Then, tmp3240)

					tmp3242 := PrimCons(V1625, tmp3241)

					__e.Return(PrimCons(symif, tmp3242))
					return

				}, 1)

				tmp3243 := PrimHead(V1627)

				tmp3244 := Call(__e, PrimNS2Value(symshen_4then), V1625, V1626, tmp3243, Else)

				__e.TailApply(tmp3239, tmp3244)
				return

			}, 1)

			tmp3245 := PrimTail(V1627)

			tmp3246 := PrimHead(tmp3245)

			tmp3247 := Call(__e, PrimNS2Value(symshen_4else), V1626, tmp3246, V1628)

			__e.TailApply(tmp3238, tmp3247)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4branch)
			return
		}

	}, 4)

	tmp3258 := Call(__e, PrimNS2Value(symdef), symshen_4branch, tmp3236)

	_ = tmp3258

	tmp3259 := MakeNative(func(__e *ControlFlow) {
		V1629 := __e.Get(1)
		_ = V1629
		V1630 := __e.Get(2)
		_ = V1630
		V1631 := __e.Get(3)
		_ = V1631
		tmp3260 := MakeNative(func(__e *ControlFlow) {
			Else := __e.Get(1)
			_ = Else
			tmp3262 := Call(__e, PrimNS2Value(symshen_4inline_2), Else)

			if True == tmp3262 {
				__e.Return(Else)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4procedure_1call), V1629, Else)
				return
			}

		}, 1)

		tmp3263 := Call(__e, PrimNS2Value(symshen_4vertical), V1629, V1630, V1631)

		__e.TailApply(tmp3260, tmp3263)
		return

	}, 3)

	tmp3264 := Call(__e, PrimNS2Value(symdef), symshen_4else, tmp3259)

	_ = tmp3264

	tmp3265 := MakeNative(func(__e *ControlFlow) {
		V1632 := __e.Get(1)
		_ = V1632
		V1633 := __e.Get(2)
		_ = V1633
		tmp3266 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			tmp3267 := MakeNative(func(__e *ControlFlow) {
				Used := __e.Get(1)
				_ = Used
				tmp3268 := MakeNative(func(__e *ControlFlow) {
					KL := __e.Get(1)
					_ = KL
					tmp3269 := MakeNative(func(__e *ControlFlow) {
						EvalKL := __e.Get(1)
						_ = EvalKL
						tmp3270 := MakeNative(func(__e *ControlFlow) {
							Record := __e.Get(1)
							_ = Record
							__e.Return(PrimCons(F, Used))
							return
						}, 1)

						tmp3271 := Call(__e, PrimNS2Value(symshen_4record_1kl), F, KL)

						__e.TailApply(tmp3270, tmp3271)
						return

					}, 1)

					tmp3272 := Call(__e, PrimNS2Value(symeval_1kl), KL)

					__e.TailApply(tmp3269, tmp3272)
					return

				}, 1)

				tmp3273 := PrimCons(V1633, Nil)

				tmp3274 := PrimCons(Used, tmp3273)

				tmp3275 := PrimCons(F, tmp3274)

				tmp3276 := PrimCons(symdefun, tmp3275)

				__e.TailApply(tmp3268, tmp3276)
				return

			}, 1)

			tmp3277 := Call(__e, PrimNS2Value(symshen_4remove_1if_1unused), V1632, V1633)

			__e.TailApply(tmp3267, tmp3277)
			return

		}, 1)

		tmp3278 := Call(__e, PrimNS2Value(symgensym), symshen_4else)

		__e.TailApply(tmp3266, tmp3278)
		return

	}, 2)

	tmp3279 := Call(__e, PrimNS2Value(symdef), symshen_4procedure_1call, tmp3265)

	_ = tmp3279

	tmp3280 := MakeNative(func(__e *ControlFlow) {
		V1640 := __e.Get(1)
		_ = V1640
		V1641 := __e.Get(2)
		_ = V1641
		tmp3291 := PrimEqual(Nil, V1640)

		if True == tmp3291 {
			__e.Return(Nil)
			return
		} else {
			tmp3290 := PrimIsPair(V1640)

			if True == tmp3290 {
				tmp3288 := PrimHead(V1640)

				tmp3289 := Call(__e, PrimNS2Value(symshen_4occurs_2), tmp3288, V1641)

				if True == tmp3289 {
					tmp3285 := PrimHead(V1640)

					tmp3286 := PrimTail(V1640)

					tmp3287 := Call(__e, PrimNS2Value(symshen_4remove_1if_1unused), tmp3286, V1641)

					__e.Return(PrimCons(tmp3285, tmp3287))
					return

				} else {
					tmp3284 := PrimTail(V1640)

					__e.TailApply(PrimNS2Value(symshen_4remove_1if_1unused), tmp3284, V1641)
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-if-unused")))
				return
			}

		}

	}, 2)

	tmp3292 := Call(__e, PrimNS2Value(symdef), symshen_4remove_1if_1unused, tmp3280)

	_ = tmp3292

	tmp3293 := MakeNative(func(__e *ControlFlow) {
		V1642 := __e.Get(1)
		_ = V1642
		V1643 := __e.Get(2)
		_ = V1643
		V1644 := __e.Get(3)
		_ = V1644
		V1645 := __e.Get(4)
		_ = V1645
		tmp3294 := Call(__e, PrimNS2Value(symshen_4selectors), V1642, V1644)

		__e.TailApply(PrimNS2Value(symshen_4horizontal), tmp3294, V1643, V1644, V1645)
		return

	}, 4)

	tmp3295 := Call(__e, PrimNS2Value(symdef), symshen_4then, tmp3293)

	_ = tmp3295

	tmp3296 := MakeNative(func(__e *ControlFlow) {
		V1654 := __e.Get(1)
		_ = V1654
		V1655 := __e.Get(2)
		_ = V1655
		V1656 := __e.Get(3)
		_ = V1656
		V1657 := __e.Get(4)
		_ = V1657
		tmp3312 := PrimEqual(Nil, V1654)

		if True == tmp3312 {
			__e.TailApply(PrimNS2Value(symshen_4vertical), V1655, V1656, V1657)
			return
		} else {
			tmp3311 := PrimIsPair(V1654)

			if True == tmp3311 {
				tmp3299 := MakeNative(func(__e *ControlFlow) {
					V := __e.Get(1)
					_ = V
					tmp3300 := PrimHead(V1654)

					tmp3301 := PrimTail(V1654)

					tmp3302 := PrimCons(V, V1655)

					tmp3303 := PrimHead(V1654)

					tmp3304 := Call(__e, PrimNS2Value(symsubst), V, tmp3303, V1656)

					tmp3305 := Call(__e, PrimNS2Value(symshen_4horizontal), tmp3301, tmp3302, tmp3304, V1657)

					tmp3306 := PrimCons(tmp3305, Nil)

					tmp3307 := PrimCons(tmp3300, tmp3306)

					tmp3308 := PrimCons(V, tmp3307)

					__e.Return(PrimCons(symlet, tmp3308))
					return

				}, 1)

				tmp3309 := Call(__e, PrimNS2Value(symprotect), symV)

				tmp3310 := Call(__e, PrimNS2Value(symgensym), tmp3309)

				__e.TailApply(tmp3299, tmp3310)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.horizontal")))
				return
			}

		}

	}, 4)

	tmp3313 := Call(__e, PrimNS2Value(symdef), symshen_4horizontal, tmp3296)

	_ = tmp3313

	tmp3314 := MakeNative(func(__e *ControlFlow) {
		V1662 := __e.Get(1)
		_ = V1662
		V1663 := __e.Get(2)
		_ = V1663
		tmp3351 := PrimIsPair(V1662)

		var ifres3337 Obj

		if True == tmp3351 {
			tmp3349 := PrimTail(V1662)

			tmp3350 := PrimIsPair(tmp3349)

			var ifres3339 Obj

			if True == tmp3350 {
				tmp3346 := PrimTail(V1662)

				tmp3347 := PrimTail(tmp3346)

				tmp3348 := PrimEqual(Nil, tmp3347)

				var ifres3341 Obj

				if True == tmp3348 {
					tmp3343 := PrimHead(V1662)

					tmp3344 := Call(__e, PrimNS2Value(symshen_4op), tmp3343)

					tmp3345 := Call(__e, PrimNS2Value(symshen_4constructor_2), tmp3344)

					var ifres3342 Obj

					if True == tmp3345 {
						ifres3342 = True

					} else {
						ifres3342 = False

					}

					ifres3341 = ifres3342

				} else {
					ifres3341 = False

				}

				var ifres3340 Obj

				if True == ifres3341 {
					ifres3340 = True

				} else {
					ifres3340 = False

				}

				ifres3339 = ifres3340

			} else {
				ifres3339 = False

			}

			var ifres3338 Obj

			if True == ifres3339 {
				ifres3338 = True

			} else {
				ifres3338 = False

			}

			ifres3337 = ifres3338

		} else {
			ifres3337 = False

		}

		if True == ifres3337 {
			tmp3316 := MakeNative(func(__e *ControlFlow) {
				Op := __e.Get(1)
				_ = Op
				tmp3317 := MakeNative(func(__e *ControlFlow) {
					Hd := __e.Get(1)
					_ = Hd
					tmp3318 := MakeNative(func(__e *ControlFlow) {
						Tl := __e.Get(1)
						_ = Tl
						tmp3319 := MakeNative(func(__e *ControlFlow) {
							RptedHd_2 := __e.Get(1)
							_ = RptedHd_2
							tmp3320 := MakeNative(func(__e *ControlFlow) {
								RptedTl_2 := __e.Get(1)
								_ = RptedTl_2
								var ifres3325 Obj

								if True == RptedHd_2 {
									var ifres3326 Obj

									if True == RptedTl_2 {
										ifres3326 = True

									} else {
										ifres3326 = False

									}

									ifres3325 = ifres3326

								} else {
									ifres3325 = False

								}

								if True == ifres3325 {
									tmp3324 := PrimCons(Tl, Nil)

									__e.Return(PrimCons(Hd, tmp3324))
									return

								} else {
									if True == RptedHd_2 {
										__e.Return(PrimCons(Hd, Nil))
										return
									} else {
										if True == RptedTl_2 {
											__e.Return(PrimCons(Tl, Nil))
											return
										} else {
											__e.Return(Nil)
											return
										}
									}
								}

							}, 1)

							tmp3327 := Call(__e, PrimNS2Value(symshen_4rpted_2), Tl, V1663)

							__e.TailApply(tmp3320, tmp3327)
							return

						}, 1)

						tmp3328 := Call(__e, PrimNS2Value(symshen_4rpted_2), Hd, V1663)

						__e.TailApply(tmp3319, tmp3328)
						return

					}, 1)

					tmp3329 := Call(__e, PrimNS2Value(symshen_4op2), Op)

					tmp3330 := PrimTail(V1662)

					tmp3331 := PrimCons(tmp3329, tmp3330)

					__e.TailApply(tmp3318, tmp3331)
					return

				}, 1)

				tmp3332 := Call(__e, PrimNS2Value(symshen_4op1), Op)

				tmp3333 := PrimTail(V1662)

				tmp3334 := PrimCons(tmp3332, tmp3333)

				__e.TailApply(tmp3317, tmp3334)
				return

			}, 1)

			tmp3335 := PrimHead(V1662)

			tmp3336 := Call(__e, PrimNS2Value(symshen_4op), tmp3335)

			__e.TailApply(tmp3316, tmp3336)
			return

		} else {
			__e.Return(Nil)
			return
		}

	}, 2)

	tmp3352 := Call(__e, PrimNS2Value(symdef), symshen_4selectors, tmp3314)

	_ = tmp3352

	tmp3353 := MakeNative(func(__e *ControlFlow) {
		V1664 := __e.Get(1)
		_ = V1664
		V1665 := __e.Get(2)
		_ = V1665
		tmp3354 := Call(__e, PrimNS2Value(symoccurrences), V1664, V1665)

		__e.Return(PrimGreatThan(tmp3354, MakeNumber(1)))
		return

	}, 2)

	tmp3355 := Call(__e, PrimNS2Value(symdef), symshen_4rpted_2, tmp3353)

	_ = tmp3355

	tmp3356 := MakeNative(func(__e *ControlFlow) {
		V1666 := __e.Get(1)
		_ = V1666
		tmp3364 := PrimIsPair(V1666)

		if True == tmp3364 {
			tmp3362 := PrimHead(V1666)

			tmp3363 := Call(__e, PrimNS2Value(symatom_2), tmp3362)

			if True == tmp3363 {
				tmp3360 := PrimTail(V1666)

				tmp3361 := Call(__e, PrimNS2Value(symshen_4inline_2), tmp3360)

				if True == tmp3361 {
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
			__e.TailApply(PrimNS2Value(symatom_2), V1666)
			return
		}

	}, 1)

	tmp3365 := Call(__e, PrimNS2Value(symdef), symshen_4inline_2, tmp3356)

	_ = tmp3365

	tmp3366 := MakeNative(func(__e *ControlFlow) {
		V1669 := __e.Get(1)
		_ = V1669
		tmp3374 := PrimEqual(symcons_2, V1669)

		if True == tmp3374 {
			__e.Return(symcons)
			return
		} else {
			tmp3373 := PrimEqual(symshen_4_7string_2, V1669)

			if True == tmp3373 {
				__e.Return(sym_8s)
				return
			} else {
				tmp3372 := PrimEqual(symshen_4_7vector_2, V1669)

				if True == tmp3372 {
					__e.Return(sym_8v)
					return
				} else {
					tmp3371 := PrimEqual(symtuple_2, V1669)

					if True == tmp3371 {
						__e.Return(sym_8p)
						return
					} else {
						__e.Return(symshen_4skip)
						return
					}

				}

			}

		}

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), symshen_4op, tmp3366)
	return

}, 0)
