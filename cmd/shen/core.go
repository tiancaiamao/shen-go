package main

import . "github.com/tiancaiamao/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
	tmp1492 := MakeNative(func(__e *ControlFlow) {
		V520 := __e.Get(1)
		_ = V520
		tmp1493 := MakeNative(func(__e *ControlFlow) {
			W521 := __e.Get(1)
			_ = W521
			__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), W521)
			return
		}, 1)

		tmp1494 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V520)

		__e.TailApply(tmp1493, tmp1494)
		return

	}, 1)

	tmp1495 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp1492)

	_ = tmp1495

	tmp1496 := MakeNative(func(__e *ControlFlow) {
		V522 := __e.Get(1)
		_ = V522
		tmp1549 := PrimIsPair(V522)

		var ifres1523 Obj

		if True == tmp1549 {
			tmp1547 := PrimHead(V522)

			tmp1548 := PrimEqual(symdefun, tmp1547)

			var ifres1525 Obj

			if True == tmp1548 {
				tmp1545 := PrimTail(V522)

				tmp1546 := PrimIsPair(tmp1545)

				var ifres1527 Obj

				if True == tmp1546 {
					tmp1542 := PrimTail(V522)

					tmp1543 := PrimTail(tmp1542)

					tmp1544 := PrimIsPair(tmp1543)

					var ifres1529 Obj

					if True == tmp1544 {
						tmp1538 := PrimTail(V522)

						tmp1539 := PrimTail(tmp1538)

						tmp1540 := PrimTail(tmp1539)

						tmp1541 := PrimIsPair(tmp1540)

						var ifres1531 Obj

						if True == tmp1541 {
							tmp1533 := PrimTail(V522)

							tmp1534 := PrimTail(tmp1533)

							tmp1535 := PrimTail(tmp1534)

							tmp1536 := PrimTail(tmp1535)

							tmp1537 := PrimEqual(Nil, tmp1536)

							var ifres1532 Obj

							if True == tmp1537 {
								ifres1532 = True

							} else {
								ifres1532 = False

							}

							ifres1531 = ifres1532

						} else {
							ifres1531 = False

						}

						var ifres1530 Obj

						if True == ifres1531 {
							ifres1530 = True

						} else {
							ifres1530 = False

						}

						ifres1529 = ifres1530

					} else {
						ifres1529 = False

					}

					var ifres1528 Obj

					if True == ifres1529 {
						ifres1528 = True

					} else {
						ifres1528 = False

					}

					ifres1527 = ifres1528

				} else {
					ifres1527 = False

				}

				var ifres1526 Obj

				if True == ifres1527 {
					ifres1526 = True

				} else {
					ifres1526 = False

				}

				ifres1525 = ifres1526

			} else {
				ifres1525 = False

			}

			var ifres1524 Obj

			if True == ifres1525 {
				ifres1524 = True

			} else {
				ifres1524 = False

			}

			ifres1523 = ifres1524

		} else {
			ifres1523 = False

		}

		if True == ifres1523 {
			tmp1497 := MakeNative(func(__e *ControlFlow) {
				W523 := __e.Get(1)
				_ = W523
				tmp1498 := MakeNative(func(__e *ControlFlow) {
					W524 := __e.Get(1)
					_ = W524
					tmp1499 := MakeNative(func(__e *ControlFlow) {
						W525 := __e.Get(1)
						_ = W525
						tmp1500 := MakeNative(func(__e *ControlFlow) {
							W526 := __e.Get(1)
							_ = W526
							tmp1501 := PrimTail(V522)

							tmp1502 := PrimHead(tmp1501)

							__e.TailApply(PrimFunc(symshen_4fn_1print), tmp1502)
							return

						}, 1)

						tmp1503 := Call(__e, PrimFunc(symeval_1kl), V522)

						__e.TailApply(tmp1500, tmp1503)
						return

					}, 1)

					tmp1504 := PrimTail(V522)

					tmp1505 := PrimHead(tmp1504)

					tmp1506 := Call(__e, PrimFunc(symshen_4record_1kl), tmp1505, V522)

					__e.TailApply(tmp1499, tmp1506)
					return

				}, 1)

				tmp1507 := PrimTail(V522)

				tmp1508 := PrimHead(tmp1507)

				tmp1509 := PrimTail(V522)

				tmp1510 := PrimTail(tmp1509)

				tmp1511 := PrimHead(tmp1510)

				tmp1512 := Call(__e, PrimFunc(symlength), tmp1511)

				tmp1513 := Call(__e, PrimFunc(symshen_4store_1arity), tmp1508, tmp1512)

				__e.TailApply(tmp1498, tmp1513)
				return

			}, 1)

			tmp1519 := PrimTail(V522)

			tmp1520 := PrimHead(tmp1519)

			tmp1521 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp1520)

			var ifres1514 Obj

			if True == tmp1521 {
				tmp1515 := PrimTail(V522)

				tmp1516 := PrimHead(tmp1515)

				tmp1517 := Call(__e, PrimFunc(symshen_4app), tmp1516, MakeString(" is not a legitimate function name\n"), symshen_4a)

				tmp1518 := PrimSimpleError(tmp1517)

				ifres1514 = tmp1518

			} else {
				ifres1514 = symshen_4skip

			}

			__e.TailApply(tmp1497, ifres1514)
			return

		} else {
			__e.Return(V522)
			return
		}

	}, 1)

	tmp1550 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp1496)

	_ = tmp1550

	tmp1551 := MakeNative(func(__e *ControlFlow) {
		V527 := __e.Get(1)
		_ = V527
		tmp1652 := PrimIsPair(V527)

		var ifres1644 Obj

		if True == tmp1652 {
			tmp1650 := PrimHead(V527)

			tmp1651 := PrimEqual(symdefine, tmp1650)

			var ifres1646 Obj

			if True == tmp1651 {
				tmp1648 := PrimTail(V527)

				tmp1649 := PrimIsPair(tmp1648)

				var ifres1647 Obj

				if True == tmp1649 {
					ifres1647 = True

				} else {
					ifres1647 = False

				}

				ifres1646 = ifres1647

			} else {
				ifres1646 = False

			}

			var ifres1645 Obj

			if True == ifres1646 {
				ifres1645 = True

			} else {
				ifres1645 = False

			}

			ifres1644 = ifres1645

		} else {
			ifres1644 = False

		}

		if True == ifres1644 {
			tmp1552 := PrimTail(V527)

			tmp1553 := PrimHead(tmp1552)

			tmp1554 := PrimTail(V527)

			tmp1555 := PrimTail(tmp1554)

			__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp1553, tmp1555)
			return

		} else {
			tmp1642 := PrimIsPair(V527)

			var ifres1616 Obj

			if True == tmp1642 {
				tmp1640 := PrimHead(V527)

				tmp1641 := PrimEqual(symdefun, tmp1640)

				var ifres1618 Obj

				if True == tmp1641 {
					tmp1638 := PrimTail(V527)

					tmp1639 := PrimIsPair(tmp1638)

					var ifres1620 Obj

					if True == tmp1639 {
						tmp1635 := PrimTail(V527)

						tmp1636 := PrimTail(tmp1635)

						tmp1637 := PrimIsPair(tmp1636)

						var ifres1622 Obj

						if True == tmp1637 {
							tmp1631 := PrimTail(V527)

							tmp1632 := PrimTail(tmp1631)

							tmp1633 := PrimTail(tmp1632)

							tmp1634 := PrimIsPair(tmp1633)

							var ifres1624 Obj

							if True == tmp1634 {
								tmp1626 := PrimTail(V527)

								tmp1627 := PrimTail(tmp1626)

								tmp1628 := PrimTail(tmp1627)

								tmp1629 := PrimTail(tmp1628)

								tmp1630 := PrimEqual(Nil, tmp1629)

								var ifres1625 Obj

								if True == tmp1630 {
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

						var ifres1621 Obj

						if True == ifres1622 {
							ifres1621 = True

						} else {
							ifres1621 = False

						}

						ifres1620 = ifres1621

					} else {
						ifres1620 = False

					}

					var ifres1619 Obj

					if True == ifres1620 {
						ifres1619 = True

					} else {
						ifres1619 = False

					}

					ifres1618 = ifres1619

				} else {
					ifres1618 = False

				}

				var ifres1617 Obj

				if True == ifres1618 {
					ifres1617 = True

				} else {
					ifres1617 = False

				}

				ifres1616 = ifres1617

			} else {
				ifres1616 = False

			}

			if True == ifres1616 {
				__e.Return(V527)
				return
			} else {
				tmp1614 := PrimIsPair(V527)

				var ifres1595 Obj

				if True == tmp1614 {
					tmp1612 := PrimHead(V527)

					tmp1613 := PrimEqual(symtype, tmp1612)

					var ifres1597 Obj

					if True == tmp1613 {
						tmp1610 := PrimTail(V527)

						tmp1611 := PrimIsPair(tmp1610)

						var ifres1599 Obj

						if True == tmp1611 {
							tmp1607 := PrimTail(V527)

							tmp1608 := PrimTail(tmp1607)

							tmp1609 := PrimIsPair(tmp1608)

							var ifres1601 Obj

							if True == tmp1609 {
								tmp1603 := PrimTail(V527)

								tmp1604 := PrimTail(tmp1603)

								tmp1605 := PrimTail(tmp1604)

								tmp1606 := PrimEqual(Nil, tmp1605)

								var ifres1602 Obj

								if True == tmp1606 {
									ifres1602 = True

								} else {
									ifres1602 = False

								}

								ifres1601 = ifres1602

							} else {
								ifres1601 = False

							}

							var ifres1600 Obj

							if True == ifres1601 {
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

				if True == ifres1595 {
					tmp1556 := PrimTail(V527)

					tmp1557 := PrimHead(tmp1556)

					tmp1558 := PrimTail(V527)

					tmp1559 := PrimTail(tmp1558)

					tmp1560 := PrimHead(tmp1559)

					tmp1561 := Call(__e, PrimFunc(symshen_4rcons__form), tmp1560)

					tmp1562 := PrimCons(tmp1561, Nil)

					tmp1563 := PrimCons(tmp1557, tmp1562)

					__e.Return(PrimCons(symtype, tmp1563))
					return

				} else {
					tmp1593 := PrimIsPair(V527)

					var ifres1574 Obj

					if True == tmp1593 {
						tmp1591 := PrimHead(V527)

						tmp1592 := PrimEqual(syminput_7, tmp1591)

						var ifres1576 Obj

						if True == tmp1592 {
							tmp1589 := PrimTail(V527)

							tmp1590 := PrimIsPair(tmp1589)

							var ifres1578 Obj

							if True == tmp1590 {
								tmp1586 := PrimTail(V527)

								tmp1587 := PrimTail(tmp1586)

								tmp1588 := PrimIsPair(tmp1587)

								var ifres1580 Obj

								if True == tmp1588 {
									tmp1582 := PrimTail(V527)

									tmp1583 := PrimTail(tmp1582)

									tmp1584 := PrimTail(tmp1583)

									tmp1585 := PrimEqual(Nil, tmp1584)

									var ifres1581 Obj

									if True == tmp1585 {
										ifres1581 = True

									} else {
										ifres1581 = False

									}

									ifres1580 = ifres1581

								} else {
									ifres1580 = False

								}

								var ifres1579 Obj

								if True == ifres1580 {
									ifres1579 = True

								} else {
									ifres1579 = False

								}

								ifres1578 = ifres1579

							} else {
								ifres1578 = False

							}

							var ifres1577 Obj

							if True == ifres1578 {
								ifres1577 = True

							} else {
								ifres1577 = False

							}

							ifres1576 = ifres1577

						} else {
							ifres1576 = False

						}

						var ifres1575 Obj

						if True == ifres1576 {
							ifres1575 = True

						} else {
							ifres1575 = False

						}

						ifres1574 = ifres1575

					} else {
						ifres1574 = False

					}

					if True == ifres1574 {
						tmp1564 := PrimTail(V527)

						tmp1565 := PrimHead(tmp1564)

						tmp1566 := Call(__e, PrimFunc(symshen_4rcons__form), tmp1565)

						tmp1567 := PrimTail(V527)

						tmp1568 := PrimTail(tmp1567)

						tmp1569 := PrimCons(tmp1566, tmp1568)

						__e.Return(PrimCons(syminput_7, tmp1569))
						return

					} else {
						tmp1572 := PrimIsPair(V527)

						if True == tmp1572 {
							tmp1570 := MakeNative(func(__e *ControlFlow) {
								Z528 := __e.Get(1)
								_ = Z528
								__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z528)
								return
							}, 1)

							__e.TailApply(PrimFunc(symmap), tmp1570, V527)
							return

						} else {
							__e.Return(V527)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp1653 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp1551)

	_ = tmp1653

	tmp1654 := MakeNative(func(__e *ControlFlow) {
		V529 := __e.Get(1)
		_ = V529
		V530 := __e.Get(2)
		_ = V530
		tmp1655 := MakeNative(func(__e *ControlFlow) {
			Z531 := __e.Get(1)
			_ = Z531
			__e.TailApply(PrimFunc(symshen_4_5define_6), Z531)
			return
		}, 1)

		tmp1656 := PrimCons(V529, V530)

		__e.TailApply(PrimFunc(symcompile), tmp1655, tmp1656)
		return

	}, 2)

	tmp1657 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp1654)

	_ = tmp1657

	tmp1658 := MakeNative(func(__e *ControlFlow) {
		V532 := __e.Get(1)
		_ = V532
		tmp1659 := MakeNative(func(__e *ControlFlow) {
			W533 := __e.Get(1)
			_ = W533
			tmp1682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W533)

			if True == tmp1682 {
				tmp1660 := MakeNative(func(__e *ControlFlow) {
					W544 := __e.Get(1)
					_ = W544
					tmp1662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W544)

					if True == tmp1662 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W544)
						return
					}

				}, 1)

				tmp1663 := MakeNative(func(__e *ControlFlow) {
					W545 := __e.Get(1)
					_ = W545
					tmp1678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W545)

					if True == tmp1678 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp1664 := MakeNative(func(__e *ControlFlow) {
							W546 := __e.Get(1)
							_ = W546
							tmp1665 := MakeNative(func(__e *ControlFlow) {
								W547 := __e.Get(1)
								_ = W547
								tmp1666 := MakeNative(func(__e *ControlFlow) {
									W548 := __e.Get(1)
									_ = W548
									tmp1673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W548)

									if True == tmp1673 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp1667 := MakeNative(func(__e *ControlFlow) {
											W549 := __e.Get(1)
											_ = W549
											tmp1668 := MakeNative(func(__e *ControlFlow) {
												W550 := __e.Get(1)
												_ = W550
												tmp1669 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W546, W549)

												__e.TailApply(PrimFunc(symshen_4comb), W550, tmp1669)
												return

											}, 1)

											tmp1670 := Call(__e, PrimFunc(symshen_4in_1_6), W548)

											__e.TailApply(tmp1668, tmp1670)
											return

										}, 1)

										tmp1671 := Call(__e, PrimFunc(symshen_4_5_1out), W548)

										__e.TailApply(tmp1667, tmp1671)
										return

									}

								}, 1)

								tmp1674 := Call(__e, PrimFunc(symshen_4_5rules_6), W547)

								__e.TailApply(tmp1666, tmp1674)
								return

							}, 1)

							tmp1675 := Call(__e, PrimFunc(symshen_4in_1_6), W545)

							__e.TailApply(tmp1665, tmp1675)
							return

						}, 1)

						tmp1676 := Call(__e, PrimFunc(symshen_4_5_1out), W545)

						__e.TailApply(tmp1664, tmp1676)
						return

					}

				}, 1)

				tmp1679 := Call(__e, PrimFunc(symshen_4_5name_6), V532)

				tmp1680 := Call(__e, tmp1663, tmp1679)

				__e.TailApply(tmp1660, tmp1680)
				return

			} else {
				__e.Return(W533)
				return
			}

		}, 1)

		tmp1683 := MakeNative(func(__e *ControlFlow) {
			W534 := __e.Get(1)
			_ = W534
			tmp1712 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W534)

			if True == tmp1712 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp1684 := MakeNative(func(__e *ControlFlow) {
					W535 := __e.Get(1)
					_ = W535
					tmp1685 := MakeNative(func(__e *ControlFlow) {
						W536 := __e.Get(1)
						_ = W536
						tmp1708 := Call(__e, PrimFunc(symshen_4hds_a_2), W536, sym_i)

						if True == tmp1708 {
							tmp1686 := MakeNative(func(__e *ControlFlow) {
								W537 := __e.Get(1)
								_ = W537
								tmp1687 := MakeNative(func(__e *ControlFlow) {
									W538 := __e.Get(1)
									_ = W538
									tmp1704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W538)

									if True == tmp1704 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp1688 := MakeNative(func(__e *ControlFlow) {
											W539 := __e.Get(1)
											_ = W539
											tmp1701 := Call(__e, PrimFunc(symshen_4hds_a_2), W539, sym_j)

											if True == tmp1701 {
												tmp1689 := MakeNative(func(__e *ControlFlow) {
													W540 := __e.Get(1)
													_ = W540
													tmp1690 := MakeNative(func(__e *ControlFlow) {
														W541 := __e.Get(1)
														_ = W541
														tmp1697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W541)

														if True == tmp1697 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp1691 := MakeNative(func(__e *ControlFlow) {
																W542 := __e.Get(1)
																_ = W542
																tmp1692 := MakeNative(func(__e *ControlFlow) {
																	W543 := __e.Get(1)
																	_ = W543
																	tmp1693 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W535, W542)

																	__e.TailApply(PrimFunc(symshen_4comb), W543, tmp1693)
																	return

																}, 1)

																tmp1694 := Call(__e, PrimFunc(symshen_4in_1_6), W541)

																__e.TailApply(tmp1692, tmp1694)
																return

															}, 1)

															tmp1695 := Call(__e, PrimFunc(symshen_4_5_1out), W541)

															__e.TailApply(tmp1691, tmp1695)
															return

														}

													}, 1)

													tmp1698 := Call(__e, PrimFunc(symshen_4_5rules_6), W540)

													__e.TailApply(tmp1690, tmp1698)
													return

												}, 1)

												tmp1699 := Call(__e, PrimFunc(symtail), W539)

												__e.TailApply(tmp1689, tmp1699)
												return

											} else {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp1702 := Call(__e, PrimFunc(symshen_4in_1_6), W538)

										__e.TailApply(tmp1688, tmp1702)
										return

									}

								}, 1)

								tmp1705 := Call(__e, PrimFunc(symshen_4_5signature_6), W537)

								__e.TailApply(tmp1687, tmp1705)
								return

							}, 1)

							tmp1706 := Call(__e, PrimFunc(symtail), W536)

							__e.TailApply(tmp1686, tmp1706)
							return

						} else {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp1709 := Call(__e, PrimFunc(symshen_4in_1_6), W534)

					__e.TailApply(tmp1685, tmp1709)
					return

				}, 1)

				tmp1710 := Call(__e, PrimFunc(symshen_4_5_1out), W534)

				__e.TailApply(tmp1684, tmp1710)
				return

			}

		}, 1)

		tmp1713 := Call(__e, PrimFunc(symshen_4_5name_6), V532)

		tmp1714 := Call(__e, tmp1683, tmp1713)

		__e.TailApply(tmp1659, tmp1714)
		return

	}, 1)

	tmp1715 := Call(__e, ns2_1set, symshen_4_5define_6, tmp1658)

	_ = tmp1715

	tmp1716 := MakeNative(func(__e *ControlFlow) {
		V551 := __e.Get(1)
		_ = V551
		V552 := __e.Get(2)
		_ = V552
		tmp1717 := MakeNative(func(__e *ControlFlow) {
			W553 := __e.Get(1)
			_ = W553
			tmp1718 := MakeNative(func(__e *ControlFlow) {
				W555 := __e.Get(1)
				_ = W555
				tmp1719 := MakeNative(func(__e *ControlFlow) {
					W556 := __e.Get(1)
					_ = W556
					tmp1720 := MakeNative(func(__e *ControlFlow) {
						W558 := __e.Get(1)
						_ = W558
						tmp1721 := MakeNative(func(__e *ControlFlow) {
							W559 := __e.Get(1)
							_ = W559
							__e.Return(W559)
							return
						}, 1)

						tmp1722 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V551, W558, W555)

						tmp1723 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp1722)

						__e.TailApply(tmp1721, tmp1723)
						return

					}, 1)

					tmp1724 := Call(__e, PrimFunc(symshen_4unprotect), V552)

					__e.TailApply(tmp1720, tmp1724)
					return

				}, 1)

				tmp1725 := MakeNative(func(__e *ControlFlow) {
					Z557 := __e.Get(1)
					_ = Z557
					__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V551, Z557)
					return
				}, 1)

				tmp1726 := Call(__e, PrimFunc(symmap), tmp1725, V552)

				__e.TailApply(tmp1719, tmp1726)
				return

			}, 1)

			tmp1727 := Call(__e, PrimFunc(symshen_4arity_1chk), V551, W553)

			__e.TailApply(tmp1718, tmp1727)
			return

		}, 1)

		tmp1728 := MakeNative(func(__e *ControlFlow) {
			Z554 := __e.Get(1)
			_ = Z554
			__e.TailApply(PrimFunc(symfst), Z554)
			return
		}, 1)

		tmp1729 := Call(__e, PrimFunc(symmap), tmp1728, V552)

		__e.TailApply(tmp1717, tmp1729)
		return

	}, 2)

	tmp1730 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp1716)

	_ = tmp1730

	tmp1731 := MakeNative(func(__e *ControlFlow) {
		V560 := __e.Get(1)
		_ = V560
		tmp1757 := Call(__e, PrimFunc(symtuple_2), V560)

		if True == tmp1757 {
			tmp1732 := Call(__e, PrimFunc(symfst), V560)

			tmp1733 := Call(__e, PrimFunc(symshen_4unprotect), tmp1732)

			tmp1734 := Call(__e, PrimFunc(symsnd), V560)

			tmp1735 := Call(__e, PrimFunc(symshen_4unprotect), tmp1734)

			__e.TailApply(PrimFunc(sym_8p), tmp1733, tmp1735)
			return

		} else {
			tmp1755 := PrimIsPair(V560)

			var ifres1742 Obj

			if True == tmp1755 {
				tmp1753 := PrimHead(V560)

				tmp1754 := PrimEqual(symprotect, tmp1753)

				var ifres1744 Obj

				if True == tmp1754 {
					tmp1751 := PrimTail(V560)

					tmp1752 := PrimIsPair(tmp1751)

					var ifres1746 Obj

					if True == tmp1752 {
						tmp1748 := PrimTail(V560)

						tmp1749 := PrimTail(tmp1748)

						tmp1750 := PrimEqual(Nil, tmp1749)

						var ifres1747 Obj

						if True == tmp1750 {
							ifres1747 = True

						} else {
							ifres1747 = False

						}

						ifres1746 = ifres1747

					} else {
						ifres1746 = False

					}

					var ifres1745 Obj

					if True == ifres1746 {
						ifres1745 = True

					} else {
						ifres1745 = False

					}

					ifres1744 = ifres1745

				} else {
					ifres1744 = False

				}

				var ifres1743 Obj

				if True == ifres1744 {
					ifres1743 = True

				} else {
					ifres1743 = False

				}

				ifres1742 = ifres1743

			} else {
				ifres1742 = False

			}

			if True == ifres1742 {
				tmp1736 := PrimTail(V560)

				tmp1737 := PrimHead(tmp1736)

				__e.TailApply(PrimFunc(symshen_4unprotect), tmp1737)
				return

			} else {
				tmp1740 := PrimIsPair(V560)

				if True == tmp1740 {
					tmp1738 := MakeNative(func(__e *ControlFlow) {
						Z561 := __e.Get(1)
						_ = Z561
						__e.TailApply(PrimFunc(symshen_4unprotect), Z561)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp1738, V560)
					return

				} else {
					__e.Return(V560)
					return
				}

			}

		}

	}, 1)

	tmp1758 := Call(__e, ns2_1set, symshen_4unprotect, tmp1731)

	_ = tmp1758

	tmp1759 := MakeNative(func(__e *ControlFlow) {
		V562 := __e.Get(1)
		_ = V562
		tmp1760 := MakeNative(func(__e *ControlFlow) {
			W563 := __e.Get(1)
			_ = W563
			tmp1762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W563)

			if True == tmp1762 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W563)
				return
			}

		}, 1)

		tmp1778 := PrimIsPair(V562)

		var ifres1763 Obj

		if True == tmp1778 {
			tmp1764 := MakeNative(func(__e *ControlFlow) {
				W564 := __e.Get(1)
				_ = W564
				tmp1765 := MakeNative(func(__e *ControlFlow) {
					W565 := __e.Get(1)
					_ = W565
					tmp1773 := PrimIsSymbol(W564)

					var ifres1769 Obj

					if True == tmp1773 {
						tmp1771 := PrimIsVariable(W564)

						tmp1772 := PrimNot(tmp1771)

						var ifres1770 Obj

						if True == tmp1772 {
							ifres1770 = True

						} else {
							ifres1770 = False

						}

						ifres1769 = ifres1770

					} else {
						ifres1769 = False

					}

					var ifres1766 Obj

					if True == ifres1769 {
						ifres1766 = W564

					} else {
						tmp1767 := Call(__e, PrimFunc(symshen_4app), W564, MakeString(" is not a legitimate function name.\n"), symshen_4a)

						tmp1768 := PrimSimpleError(tmp1767)

						ifres1766 = tmp1768

					}

					__e.TailApply(PrimFunc(symshen_4comb), W565, ifres1766)
					return

				}, 1)

				tmp1774 := Call(__e, PrimFunc(symtail), V562)

				__e.TailApply(tmp1765, tmp1774)
				return

			}, 1)

			tmp1775 := Call(__e, PrimFunc(symhead), V562)

			tmp1776 := Call(__e, tmp1764, tmp1775)

			ifres1763 = tmp1776

		} else {
			tmp1777 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres1763 = tmp1777

		}

		__e.TailApply(tmp1760, ifres1763)
		return

	}, 1)

	tmp1779 := Call(__e, ns2_1set, symshen_4_5name_6, tmp1759)

	_ = tmp1779

	tmp1780 := MakeNative(func(__e *ControlFlow) {
		V566 := __e.Get(1)
		_ = V566
		tmp1781 := MakeNative(func(__e *ControlFlow) {
			W567 := __e.Get(1)
			_ = W567
			tmp1793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W567)

			if True == tmp1793 {
				tmp1782 := MakeNative(func(__e *ControlFlow) {
					W573 := __e.Get(1)
					_ = W573
					tmp1784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W573)

					if True == tmp1784 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W573)
						return
					}

				}, 1)

				tmp1785 := MakeNative(func(__e *ControlFlow) {
					W574 := __e.Get(1)
					_ = W574
					tmp1789 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W574)

					if True == tmp1789 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp1786 := MakeNative(func(__e *ControlFlow) {
							W575 := __e.Get(1)
							_ = W575
							__e.TailApply(PrimFunc(symshen_4comb), W575, Nil)
							return
						}, 1)

						tmp1787 := Call(__e, PrimFunc(symshen_4in_1_6), W574)

						__e.TailApply(tmp1786, tmp1787)
						return

					}

				}, 1)

				tmp1790 := Call(__e, PrimFunc(sym_5e_6), V566)

				tmp1791 := Call(__e, tmp1785, tmp1790)

				__e.TailApply(tmp1782, tmp1791)
				return

			} else {
				__e.Return(W567)
				return
			}

		}, 1)

		tmp1815 := PrimIsPair(V566)

		var ifres1794 Obj

		if True == tmp1815 {
			tmp1795 := MakeNative(func(__e *ControlFlow) {
				W568 := __e.Get(1)
				_ = W568
				tmp1796 := MakeNative(func(__e *ControlFlow) {
					W569 := __e.Get(1)
					_ = W569
					tmp1797 := MakeNative(func(__e *ControlFlow) {
						W570 := __e.Get(1)
						_ = W570
						tmp1809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W570)

						if True == tmp1809 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp1798 := MakeNative(func(__e *ControlFlow) {
								W571 := __e.Get(1)
								_ = W571
								tmp1799 := MakeNative(func(__e *ControlFlow) {
									W572 := __e.Get(1)
									_ = W572
									tmp1802 := PrimCons(sym_j, Nil)

									tmp1803 := PrimCons(sym_i, tmp1802)

									tmp1804 := Call(__e, PrimFunc(symelement_2), W568, tmp1803)

									tmp1805 := PrimNot(tmp1804)

									if True == tmp1805 {
										tmp1800 := PrimCons(W568, W571)

										__e.TailApply(PrimFunc(symshen_4comb), W572, tmp1800)
										return

									} else {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp1806 := Call(__e, PrimFunc(symshen_4in_1_6), W570)

								__e.TailApply(tmp1799, tmp1806)
								return

							}, 1)

							tmp1807 := Call(__e, PrimFunc(symshen_4_5_1out), W570)

							__e.TailApply(tmp1798, tmp1807)
							return

						}

					}, 1)

					tmp1810 := Call(__e, PrimFunc(symshen_4_5signature_6), W569)

					__e.TailApply(tmp1797, tmp1810)
					return

				}, 1)

				tmp1811 := Call(__e, PrimFunc(symtail), V566)

				__e.TailApply(tmp1796, tmp1811)
				return

			}, 1)

			tmp1812 := Call(__e, PrimFunc(symhead), V566)

			tmp1813 := Call(__e, tmp1795, tmp1812)

			ifres1794 = tmp1813

		} else {
			tmp1814 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres1794 = tmp1814

		}

		__e.TailApply(tmp1781, ifres1794)
		return

	}, 1)

	tmp1816 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp1780)

	_ = tmp1816

	tmp1817 := MakeNative(func(__e *ControlFlow) {
		V576 := __e.Get(1)
		_ = V576
		tmp1818 := MakeNative(func(__e *ControlFlow) {
			W577 := __e.Get(1)
			_ = W577
			tmp1837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W577)

			if True == tmp1837 {
				tmp1819 := MakeNative(func(__e *ControlFlow) {
					W584 := __e.Get(1)
					_ = W584
					tmp1821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W584)

					if True == tmp1821 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W584)
						return
					}

				}, 1)

				tmp1822 := MakeNative(func(__e *ControlFlow) {
					W585 := __e.Get(1)
					_ = W585
					tmp1833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W585)

					if True == tmp1833 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp1823 := MakeNative(func(__e *ControlFlow) {
							W586 := __e.Get(1)
							_ = W586
							tmp1824 := MakeNative(func(__e *ControlFlow) {
								W587 := __e.Get(1)
								_ = W587
								tmp1829 := Call(__e, PrimFunc(symempty_2), W586)

								var ifres1825 Obj

								if True == tmp1829 {
									ifres1825 = Nil

								} else {
									tmp1826 := Call(__e, PrimFunc(symshen_4app), W586, MakeString("\n ..."), symshen_4r)

									tmp1827 := PrimStringConcat(MakeString("Shen syntax error here:\n "), tmp1826)

									tmp1828 := PrimSimpleError(tmp1827)

									ifres1825 = tmp1828

								}

								__e.TailApply(PrimFunc(symshen_4comb), W587, ifres1825)
								return

							}, 1)

							tmp1830 := Call(__e, PrimFunc(symshen_4in_1_6), W585)

							__e.TailApply(tmp1824, tmp1830)
							return

						}, 1)

						tmp1831 := Call(__e, PrimFunc(symshen_4_5_1out), W585)

						__e.TailApply(tmp1823, tmp1831)
						return

					}

				}, 1)

				tmp1834 := Call(__e, PrimFunc(sym_5_b_6), V576)

				tmp1835 := Call(__e, tmp1822, tmp1834)

				__e.TailApply(tmp1819, tmp1835)
				return

			} else {
				__e.Return(W577)
				return
			}

		}, 1)

		tmp1838 := MakeNative(func(__e *ControlFlow) {
			W578 := __e.Get(1)
			_ = W578
			tmp1854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W578)

			if True == tmp1854 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp1839 := MakeNative(func(__e *ControlFlow) {
					W579 := __e.Get(1)
					_ = W579
					tmp1840 := MakeNative(func(__e *ControlFlow) {
						W580 := __e.Get(1)
						_ = W580
						tmp1841 := MakeNative(func(__e *ControlFlow) {
							W581 := __e.Get(1)
							_ = W581
							tmp1849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W581)

							if True == tmp1849 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp1842 := MakeNative(func(__e *ControlFlow) {
									W582 := __e.Get(1)
									_ = W582
									tmp1843 := MakeNative(func(__e *ControlFlow) {
										W583 := __e.Get(1)
										_ = W583
										tmp1844 := Call(__e, PrimFunc(symshen_4linearise), W579)

										tmp1845 := PrimCons(tmp1844, W582)

										__e.TailApply(PrimFunc(symshen_4comb), W583, tmp1845)
										return

									}, 1)

									tmp1846 := Call(__e, PrimFunc(symshen_4in_1_6), W581)

									__e.TailApply(tmp1843, tmp1846)
									return

								}, 1)

								tmp1847 := Call(__e, PrimFunc(symshen_4_5_1out), W581)

								__e.TailApply(tmp1842, tmp1847)
								return

							}

						}, 1)

						tmp1850 := Call(__e, PrimFunc(symshen_4_5rules_6), W580)

						__e.TailApply(tmp1841, tmp1850)
						return

					}, 1)

					tmp1851 := Call(__e, PrimFunc(symshen_4in_1_6), W578)

					__e.TailApply(tmp1840, tmp1851)
					return

				}, 1)

				tmp1852 := Call(__e, PrimFunc(symshen_4_5_1out), W578)

				__e.TailApply(tmp1839, tmp1852)
				return

			}

		}, 1)

		tmp1855 := Call(__e, PrimFunc(symshen_4_5rule_6), V576)

		tmp1856 := Call(__e, tmp1838, tmp1855)

		__e.TailApply(tmp1818, tmp1856)
		return

	}, 1)

	tmp1857 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp1817)

	_ = tmp1857

	tmp1858 := MakeNative(func(__e *ControlFlow) {
		V590 := __e.Get(1)
		_ = V590
		tmp1863 := Call(__e, PrimFunc(symtuple_2), V590)

		if True == tmp1863 {
			tmp1859 := Call(__e, PrimFunc(symfst), V590)

			tmp1860 := Call(__e, PrimFunc(symfst), V590)

			tmp1861 := Call(__e, PrimFunc(symsnd), V590)

			__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1859, tmp1860, Nil, tmp1861)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise")))
			return
		}

	}, 1)

	tmp1864 := Call(__e, ns2_1set, symshen_4linearise, tmp1858)

	_ = tmp1864

	tmp1865 := MakeNative(func(__e *ControlFlow) {
		V603 := __e.Get(1)
		_ = V603
		V604 := __e.Get(2)
		_ = V604
		V605 := __e.Get(3)
		_ = V605
		V606 := __e.Get(4)
		_ = V606
		tmp1903 := PrimEqual(Nil, V603)

		if True == tmp1903 {
			__e.TailApply(PrimFunc(sym_8p), V604, V606)
			return
		} else {
			tmp1901 := PrimIsPair(V603)

			var ifres1897 Obj

			if True == tmp1901 {
				tmp1899 := PrimHead(V603)

				tmp1900 := PrimIsPair(tmp1899)

				var ifres1898 Obj

				if True == tmp1900 {
					ifres1898 = True

				} else {
					ifres1898 = False

				}

				ifres1897 = ifres1898

			} else {
				ifres1897 = False

			}

			if True == ifres1897 {
				tmp1866 := PrimHead(V603)

				tmp1867 := PrimTail(V603)

				tmp1868 := Call(__e, PrimFunc(symappend), tmp1866, tmp1867)

				__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1868, V604, V605, V606)
				return

			} else {
				tmp1895 := PrimIsPair(V603)

				var ifres1891 Obj

				if True == tmp1895 {
					tmp1893 := PrimHead(V603)

					tmp1894 := PrimIsVariable(tmp1893)

					var ifres1892 Obj

					if True == tmp1894 {
						ifres1892 = True

					} else {
						ifres1892 = False

					}

					ifres1891 = ifres1892

				} else {
					ifres1891 = False

				}

				if True == ifres1891 {
					tmp1885 := PrimHead(V603)

					tmp1886 := Call(__e, PrimFunc(symelement_2), tmp1885, V605)

					if True == tmp1886 {
						tmp1869 := MakeNative(func(__e *ControlFlow) {
							W607 := __e.Get(1)
							_ = W607
							tmp1870 := PrimTail(V603)

							tmp1871 := PrimHead(V603)

							tmp1872 := Call(__e, PrimFunc(symshen_4rep_1X), tmp1871, W607, V604)

							tmp1873 := PrimHead(V603)

							tmp1874 := PrimCons(tmp1873, Nil)

							tmp1875 := PrimCons(W607, tmp1874)

							tmp1876 := PrimCons(sym_a, tmp1875)

							tmp1877 := PrimCons(V606, Nil)

							tmp1878 := PrimCons(tmp1876, tmp1877)

							tmp1879 := PrimCons(symwhere, tmp1878)

							__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1870, tmp1872, V605, tmp1879)
							return

						}, 1)

						tmp1880 := Call(__e, PrimFunc(symgensym), symV)

						__e.TailApply(tmp1869, tmp1880)
						return

					} else {
						tmp1881 := PrimTail(V603)

						tmp1882 := PrimHead(V603)

						tmp1883 := PrimCons(tmp1882, V605)

						__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1881, V604, tmp1883, V606)
						return

					}

				} else {
					tmp1889 := PrimIsPair(V603)

					if True == tmp1889 {
						tmp1887 := PrimTail(V603)

						__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1887, V604, V605, V606)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise-h")))
						return
					}

				}

			}

		}

	}, 4)

	tmp1904 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp1865)

	_ = tmp1904

	tmp1905 := MakeNative(func(__e *ControlFlow) {
		V608 := __e.Get(1)
		_ = V608
		tmp1906 := MakeNative(func(__e *ControlFlow) {
			W609 := __e.Get(1)
			_ = W609
			tmp1994 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W609)

			if True == tmp1994 {
				tmp1907 := MakeNative(func(__e *ControlFlow) {
					W619 := __e.Get(1)
					_ = W619
					tmp1972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W619)

					if True == tmp1972 {
						tmp1908 := MakeNative(func(__e *ControlFlow) {
							W626 := __e.Get(1)
							_ = W626
							tmp1935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W626)

							if True == tmp1935 {
								tmp1909 := MakeNative(func(__e *ControlFlow) {
									W636 := __e.Get(1)
									_ = W636
									tmp1911 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W636)

									if True == tmp1911 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										__e.Return(W636)
										return
									}

								}, 1)

								tmp1912 := MakeNative(func(__e *ControlFlow) {
									W637 := __e.Get(1)
									_ = W637
									tmp1931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W637)

									if True == tmp1931 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp1913 := MakeNative(func(__e *ControlFlow) {
											W638 := __e.Get(1)
											_ = W638
											tmp1914 := MakeNative(func(__e *ControlFlow) {
												W639 := __e.Get(1)
												_ = W639
												tmp1927 := Call(__e, PrimFunc(symshen_4hds_a_2), W639, sym_5_1)

												if True == tmp1927 {
													tmp1915 := MakeNative(func(__e *ControlFlow) {
														W640 := __e.Get(1)
														_ = W640
														tmp1924 := PrimIsPair(W640)

														if True == tmp1924 {
															tmp1916 := MakeNative(func(__e *ControlFlow) {
																W641 := __e.Get(1)
																_ = W641
																tmp1917 := MakeNative(func(__e *ControlFlow) {
																	W642 := __e.Get(1)
																	_ = W642
																	tmp1918 := PrimCons(W641, Nil)

																	tmp1919 := PrimCons(symshen_4choicepoint_b, tmp1918)

																	tmp1920 := Call(__e, PrimFunc(sym_8p), W638, tmp1919)

																	__e.TailApply(PrimFunc(symshen_4comb), W642, tmp1920)
																	return

																}, 1)

																tmp1921 := Call(__e, PrimFunc(symtail), W640)

																__e.TailApply(tmp1917, tmp1921)
																return

															}, 1)

															tmp1922 := Call(__e, PrimFunc(symhead), W640)

															__e.TailApply(tmp1916, tmp1922)
															return

														} else {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														}

													}, 1)

													tmp1925 := Call(__e, PrimFunc(symtail), W639)

													__e.TailApply(tmp1915, tmp1925)
													return

												} else {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp1928 := Call(__e, PrimFunc(symshen_4in_1_6), W637)

											__e.TailApply(tmp1914, tmp1928)
											return

										}, 1)

										tmp1929 := Call(__e, PrimFunc(symshen_4_5_1out), W637)

										__e.TailApply(tmp1913, tmp1929)
										return

									}

								}, 1)

								tmp1932 := Call(__e, PrimFunc(symshen_4_5patterns_6), V608)

								tmp1933 := Call(__e, tmp1912, tmp1932)

								__e.TailApply(tmp1909, tmp1933)
								return

							} else {
								__e.Return(W626)
								return
							}

						}, 1)

						tmp1936 := MakeNative(func(__e *ControlFlow) {
							W627 := __e.Get(1)
							_ = W627
							tmp1968 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W627)

							if True == tmp1968 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp1937 := MakeNative(func(__e *ControlFlow) {
									W628 := __e.Get(1)
									_ = W628
									tmp1938 := MakeNative(func(__e *ControlFlow) {
										W629 := __e.Get(1)
										_ = W629
										tmp1964 := Call(__e, PrimFunc(symshen_4hds_a_2), W629, sym_5_1)

										if True == tmp1964 {
											tmp1939 := MakeNative(func(__e *ControlFlow) {
												W630 := __e.Get(1)
												_ = W630
												tmp1961 := PrimIsPair(W630)

												if True == tmp1961 {
													tmp1940 := MakeNative(func(__e *ControlFlow) {
														W631 := __e.Get(1)
														_ = W631
														tmp1941 := MakeNative(func(__e *ControlFlow) {
															W632 := __e.Get(1)
															_ = W632
															tmp1957 := Call(__e, PrimFunc(symshen_4hds_a_2), W632, symwhere)

															if True == tmp1957 {
																tmp1942 := MakeNative(func(__e *ControlFlow) {
																	W633 := __e.Get(1)
																	_ = W633
																	tmp1954 := PrimIsPair(W633)

																	if True == tmp1954 {
																		tmp1943 := MakeNative(func(__e *ControlFlow) {
																			W634 := __e.Get(1)
																			_ = W634
																			tmp1944 := MakeNative(func(__e *ControlFlow) {
																				W635 := __e.Get(1)
																				_ = W635
																				tmp1945 := PrimCons(W631, Nil)

																				tmp1946 := PrimCons(symshen_4choicepoint_b, tmp1945)

																				tmp1947 := PrimCons(tmp1946, Nil)

																				tmp1948 := PrimCons(W634, tmp1947)

																				tmp1949 := PrimCons(symwhere, tmp1948)

																				tmp1950 := Call(__e, PrimFunc(sym_8p), W628, tmp1949)

																				__e.TailApply(PrimFunc(symshen_4comb), W635, tmp1950)
																				return

																			}, 1)

																			tmp1951 := Call(__e, PrimFunc(symtail), W633)

																			__e.TailApply(tmp1944, tmp1951)
																			return

																		}, 1)

																		tmp1952 := Call(__e, PrimFunc(symhead), W633)

																		__e.TailApply(tmp1943, tmp1952)
																		return

																	} else {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	}

																}, 1)

																tmp1955 := Call(__e, PrimFunc(symtail), W632)

																__e.TailApply(tmp1942, tmp1955)
																return

															} else {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															}

														}, 1)

														tmp1958 := Call(__e, PrimFunc(symtail), W630)

														__e.TailApply(tmp1941, tmp1958)
														return

													}, 1)

													tmp1959 := Call(__e, PrimFunc(symhead), W630)

													__e.TailApply(tmp1940, tmp1959)
													return

												} else {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												}

											}, 1)

											tmp1962 := Call(__e, PrimFunc(symtail), W629)

											__e.TailApply(tmp1939, tmp1962)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp1965 := Call(__e, PrimFunc(symshen_4in_1_6), W627)

									__e.TailApply(tmp1938, tmp1965)
									return

								}, 1)

								tmp1966 := Call(__e, PrimFunc(symshen_4_5_1out), W627)

								__e.TailApply(tmp1937, tmp1966)
								return

							}

						}, 1)

						tmp1969 := Call(__e, PrimFunc(symshen_4_5patterns_6), V608)

						tmp1970 := Call(__e, tmp1936, tmp1969)

						__e.TailApply(tmp1908, tmp1970)
						return

					} else {
						__e.Return(W619)
						return
					}

				}, 1)

				tmp1973 := MakeNative(func(__e *ControlFlow) {
					W620 := __e.Get(1)
					_ = W620
					tmp1990 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W620)

					if True == tmp1990 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp1974 := MakeNative(func(__e *ControlFlow) {
							W621 := __e.Get(1)
							_ = W621
							tmp1975 := MakeNative(func(__e *ControlFlow) {
								W622 := __e.Get(1)
								_ = W622
								tmp1986 := Call(__e, PrimFunc(symshen_4hds_a_2), W622, sym_1_6)

								if True == tmp1986 {
									tmp1976 := MakeNative(func(__e *ControlFlow) {
										W623 := __e.Get(1)
										_ = W623
										tmp1983 := PrimIsPair(W623)

										if True == tmp1983 {
											tmp1977 := MakeNative(func(__e *ControlFlow) {
												W624 := __e.Get(1)
												_ = W624
												tmp1978 := MakeNative(func(__e *ControlFlow) {
													W625 := __e.Get(1)
													_ = W625
													tmp1979 := Call(__e, PrimFunc(sym_8p), W621, W624)

													__e.TailApply(PrimFunc(symshen_4comb), W625, tmp1979)
													return

												}, 1)

												tmp1980 := Call(__e, PrimFunc(symtail), W623)

												__e.TailApply(tmp1978, tmp1980)
												return

											}, 1)

											tmp1981 := Call(__e, PrimFunc(symhead), W623)

											__e.TailApply(tmp1977, tmp1981)
											return

										} else {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										}

									}, 1)

									tmp1984 := Call(__e, PrimFunc(symtail), W622)

									__e.TailApply(tmp1976, tmp1984)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp1987 := Call(__e, PrimFunc(symshen_4in_1_6), W620)

							__e.TailApply(tmp1975, tmp1987)
							return

						}, 1)

						tmp1988 := Call(__e, PrimFunc(symshen_4_5_1out), W620)

						__e.TailApply(tmp1974, tmp1988)
						return

					}

				}, 1)

				tmp1991 := Call(__e, PrimFunc(symshen_4_5patterns_6), V608)

				tmp1992 := Call(__e, tmp1973, tmp1991)

				__e.TailApply(tmp1907, tmp1992)
				return

			} else {
				__e.Return(W609)
				return
			}

		}, 1)

		tmp1995 := MakeNative(func(__e *ControlFlow) {
			W610 := __e.Get(1)
			_ = W610
			tmp2025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W610)

			if True == tmp2025 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp1996 := MakeNative(func(__e *ControlFlow) {
					W611 := __e.Get(1)
					_ = W611
					tmp1997 := MakeNative(func(__e *ControlFlow) {
						W612 := __e.Get(1)
						_ = W612
						tmp2021 := Call(__e, PrimFunc(symshen_4hds_a_2), W612, sym_1_6)

						if True == tmp2021 {
							tmp1998 := MakeNative(func(__e *ControlFlow) {
								W613 := __e.Get(1)
								_ = W613
								tmp2018 := PrimIsPair(W613)

								if True == tmp2018 {
									tmp1999 := MakeNative(func(__e *ControlFlow) {
										W614 := __e.Get(1)
										_ = W614
										tmp2000 := MakeNative(func(__e *ControlFlow) {
											W615 := __e.Get(1)
											_ = W615
											tmp2014 := Call(__e, PrimFunc(symshen_4hds_a_2), W615, symwhere)

											if True == tmp2014 {
												tmp2001 := MakeNative(func(__e *ControlFlow) {
													W616 := __e.Get(1)
													_ = W616
													tmp2011 := PrimIsPair(W616)

													if True == tmp2011 {
														tmp2002 := MakeNative(func(__e *ControlFlow) {
															W617 := __e.Get(1)
															_ = W617
															tmp2003 := MakeNative(func(__e *ControlFlow) {
																W618 := __e.Get(1)
																_ = W618
																tmp2004 := PrimCons(W614, Nil)

																tmp2005 := PrimCons(W617, tmp2004)

																tmp2006 := PrimCons(symwhere, tmp2005)

																tmp2007 := Call(__e, PrimFunc(sym_8p), W611, tmp2006)

																__e.TailApply(PrimFunc(symshen_4comb), W618, tmp2007)
																return

															}, 1)

															tmp2008 := Call(__e, PrimFunc(symtail), W616)

															__e.TailApply(tmp2003, tmp2008)
															return

														}, 1)

														tmp2009 := Call(__e, PrimFunc(symhead), W616)

														__e.TailApply(tmp2002, tmp2009)
														return

													} else {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													}

												}, 1)

												tmp2012 := Call(__e, PrimFunc(symtail), W615)

												__e.TailApply(tmp2001, tmp2012)
												return

											} else {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											}

										}, 1)

										tmp2015 := Call(__e, PrimFunc(symtail), W613)

										__e.TailApply(tmp2000, tmp2015)
										return

									}, 1)

									tmp2016 := Call(__e, PrimFunc(symhead), W613)

									__e.TailApply(tmp1999, tmp2016)
									return

								} else {
									__e.TailApply(PrimFunc(symshen_4parse_1failure))
									return
								}

							}, 1)

							tmp2019 := Call(__e, PrimFunc(symtail), W612)

							__e.TailApply(tmp1998, tmp2019)
							return

						} else {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						}

					}, 1)

					tmp2022 := Call(__e, PrimFunc(symshen_4in_1_6), W610)

					__e.TailApply(tmp1997, tmp2022)
					return

				}, 1)

				tmp2023 := Call(__e, PrimFunc(symshen_4_5_1out), W610)

				__e.TailApply(tmp1996, tmp2023)
				return

			}

		}, 1)

		tmp2026 := Call(__e, PrimFunc(symshen_4_5patterns_6), V608)

		tmp2027 := Call(__e, tmp1995, tmp2026)

		__e.TailApply(tmp1906, tmp2027)
		return

	}, 1)

	tmp2028 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp1905)

	_ = tmp2028

	tmp2029 := MakeNative(func(__e *ControlFlow) {
		V643 := __e.Get(1)
		_ = V643
		tmp2030 := MakeNative(func(__e *ControlFlow) {
			W644 := __e.Get(1)
			_ = W644
			tmp2042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W644)

			if True == tmp2042 {
				tmp2031 := MakeNative(func(__e *ControlFlow) {
					W651 := __e.Get(1)
					_ = W651
					tmp2033 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W651)

					if True == tmp2033 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W651)
						return
					}

				}, 1)

				tmp2034 := MakeNative(func(__e *ControlFlow) {
					W652 := __e.Get(1)
					_ = W652
					tmp2038 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W652)

					if True == tmp2038 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp2035 := MakeNative(func(__e *ControlFlow) {
							W653 := __e.Get(1)
							_ = W653
							__e.TailApply(PrimFunc(symshen_4comb), W653, Nil)
							return
						}, 1)

						tmp2036 := Call(__e, PrimFunc(symshen_4in_1_6), W652)

						__e.TailApply(tmp2035, tmp2036)
						return

					}

				}, 1)

				tmp2039 := Call(__e, PrimFunc(sym_5e_6), V643)

				tmp2040 := Call(__e, tmp2034, tmp2039)

				__e.TailApply(tmp2031, tmp2040)
				return

			} else {
				__e.Return(W644)
				return
			}

		}, 1)

		tmp2043 := MakeNative(func(__e *ControlFlow) {
			W645 := __e.Get(1)
			_ = W645
			tmp2058 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W645)

			if True == tmp2058 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp2044 := MakeNative(func(__e *ControlFlow) {
					W646 := __e.Get(1)
					_ = W646
					tmp2045 := MakeNative(func(__e *ControlFlow) {
						W647 := __e.Get(1)
						_ = W647
						tmp2046 := MakeNative(func(__e *ControlFlow) {
							W648 := __e.Get(1)
							_ = W648
							tmp2053 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W648)

							if True == tmp2053 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp2047 := MakeNative(func(__e *ControlFlow) {
									W649 := __e.Get(1)
									_ = W649
									tmp2048 := MakeNative(func(__e *ControlFlow) {
										W650 := __e.Get(1)
										_ = W650
										tmp2049 := PrimCons(W646, W649)

										__e.TailApply(PrimFunc(symshen_4comb), W650, tmp2049)
										return

									}, 1)

									tmp2050 := Call(__e, PrimFunc(symshen_4in_1_6), W648)

									__e.TailApply(tmp2048, tmp2050)
									return

								}, 1)

								tmp2051 := Call(__e, PrimFunc(symshen_4_5_1out), W648)

								__e.TailApply(tmp2047, tmp2051)
								return

							}

						}, 1)

						tmp2054 := Call(__e, PrimFunc(symshen_4_5patterns_6), W647)

						__e.TailApply(tmp2046, tmp2054)
						return

					}, 1)

					tmp2055 := Call(__e, PrimFunc(symshen_4in_1_6), W645)

					__e.TailApply(tmp2045, tmp2055)
					return

				}, 1)

				tmp2056 := Call(__e, PrimFunc(symshen_4_5_1out), W645)

				__e.TailApply(tmp2044, tmp2056)
				return

			}

		}, 1)

		tmp2059 := Call(__e, PrimFunc(symshen_4_5pattern_6), V643)

		tmp2060 := Call(__e, tmp2043, tmp2059)

		__e.TailApply(tmp2030, tmp2060)
		return

	}, 1)

	tmp2061 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp2029)

	_ = tmp2061

	tmp2062 := MakeNative(func(__e *ControlFlow) {
		V654 := __e.Get(1)
		_ = V654
		tmp2063 := MakeNative(func(__e *ControlFlow) {
			W655 := __e.Get(1)
			_ = W655
			tmp2118 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W655)

			if True == tmp2118 {
				tmp2064 := MakeNative(func(__e *ControlFlow) {
					W669 := __e.Get(1)
					_ = W669
					tmp2092 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W669)

					if True == tmp2092 {
						tmp2065 := MakeNative(func(__e *ControlFlow) {
							W676 := __e.Get(1)
							_ = W676
							tmp2079 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W676)

							if True == tmp2079 {
								tmp2066 := MakeNative(func(__e *ControlFlow) {
									W679 := __e.Get(1)
									_ = W679
									tmp2068 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W679)

									if True == tmp2068 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										__e.Return(W679)
										return
									}

								}, 1)

								tmp2069 := MakeNative(func(__e *ControlFlow) {
									W680 := __e.Get(1)
									_ = W680
									tmp2075 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W680)

									if True == tmp2075 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp2070 := MakeNative(func(__e *ControlFlow) {
											W681 := __e.Get(1)
											_ = W681
											tmp2071 := MakeNative(func(__e *ControlFlow) {
												W682 := __e.Get(1)
												_ = W682
												__e.TailApply(PrimFunc(symshen_4comb), W682, W681)
												return
											}, 1)

											tmp2072 := Call(__e, PrimFunc(symshen_4in_1_6), W680)

											__e.TailApply(tmp2071, tmp2072)
											return

										}, 1)

										tmp2073 := Call(__e, PrimFunc(symshen_4_5_1out), W680)

										__e.TailApply(tmp2070, tmp2073)
										return

									}

								}, 1)

								tmp2076 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V654)

								tmp2077 := Call(__e, tmp2069, tmp2076)

								__e.TailApply(tmp2066, tmp2077)
								return

							} else {
								__e.Return(W676)
								return
							}

						}, 1)

						tmp2090 := PrimIsPair(V654)

						var ifres2080 Obj

						if True == tmp2090 {
							tmp2081 := MakeNative(func(__e *ControlFlow) {
								W677 := __e.Get(1)
								_ = W677
								tmp2082 := MakeNative(func(__e *ControlFlow) {
									W678 := __e.Get(1)
									_ = W678
									tmp2085 := PrimIsPair(W677)

									if True == tmp2085 {
										tmp2083 := Call(__e, PrimFunc(symshen_4constructor_1error), W677)

										__e.TailApply(PrimFunc(symshen_4comb), W678, tmp2083)
										return

									} else {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp2086 := Call(__e, PrimFunc(symtail), V654)

								__e.TailApply(tmp2082, tmp2086)
								return

							}, 1)

							tmp2087 := Call(__e, PrimFunc(symhead), V654)

							tmp2088 := Call(__e, tmp2081, tmp2087)

							ifres2080 = tmp2088

						} else {
							tmp2089 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres2080 = tmp2089

						}

						__e.TailApply(tmp2065, ifres2080)
						return

					} else {
						__e.Return(W669)
						return
					}

				}, 1)

				tmp2116 := Call(__e, PrimFunc(symshen_4ccons_2), V654)

				var ifres2093 Obj

				if True == tmp2116 {
					tmp2094 := MakeNative(func(__e *ControlFlow) {
						W670 := __e.Get(1)
						_ = W670
						tmp2095 := MakeNative(func(__e *ControlFlow) {
							W671 := __e.Get(1)
							_ = W671
							tmp2111 := Call(__e, PrimFunc(symshen_4hds_a_2), W670, symvector)

							if True == tmp2111 {
								tmp2096 := MakeNative(func(__e *ControlFlow) {
									W672 := __e.Get(1)
									_ = W672
									tmp2108 := Call(__e, PrimFunc(symshen_4hds_a_2), W672, MakeNumber(0))

									if True == tmp2108 {
										tmp2097 := MakeNative(func(__e *ControlFlow) {
											W673 := __e.Get(1)
											_ = W673
											tmp2098 := MakeNative(func(__e *ControlFlow) {
												W674 := __e.Get(1)
												_ = W674
												tmp2104 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W674)

												if True == tmp2104 {
													__e.TailApply(PrimFunc(symshen_4parse_1failure))
													return
												} else {
													tmp2099 := MakeNative(func(__e *ControlFlow) {
														W675 := __e.Get(1)
														_ = W675
														tmp2100 := PrimCons(MakeNumber(0), Nil)

														tmp2101 := PrimCons(symvector, tmp2100)

														__e.TailApply(PrimFunc(symshen_4comb), W671, tmp2101)
														return

													}, 1)

													tmp2102 := Call(__e, PrimFunc(symshen_4in_1_6), W674)

													__e.TailApply(tmp2099, tmp2102)
													return

												}

											}, 1)

											tmp2105 := Call(__e, PrimFunc(sym_5end_6), W673)

											__e.TailApply(tmp2098, tmp2105)
											return

										}, 1)

										tmp2106 := Call(__e, PrimFunc(symtail), W672)

										__e.TailApply(tmp2097, tmp2106)
										return

									} else {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									}

								}, 1)

								tmp2109 := Call(__e, PrimFunc(symtail), W670)

								__e.TailApply(tmp2096, tmp2109)
								return

							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp2112 := Call(__e, PrimFunc(symtail), V654)

						__e.TailApply(tmp2095, tmp2112)
						return

					}, 1)

					tmp2113 := Call(__e, PrimFunc(symhead), V654)

					tmp2114 := Call(__e, tmp2094, tmp2113)

					ifres2093 = tmp2114

				} else {
					tmp2115 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres2093 = tmp2115

				}

				__e.TailApply(tmp2064, ifres2093)
				return

			} else {
				__e.Return(W655)
				return
			}

		}, 1)

		tmp2159 := Call(__e, PrimFunc(symshen_4ccons_2), V654)

		var ifres2119 Obj

		if True == tmp2159 {
			tmp2120 := MakeNative(func(__e *ControlFlow) {
				W656 := __e.Get(1)
				_ = W656
				tmp2121 := MakeNative(func(__e *ControlFlow) {
					W657 := __e.Get(1)
					_ = W657
					tmp2122 := MakeNative(func(__e *ControlFlow) {
						W658 := __e.Get(1)
						_ = W658
						tmp2153 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W658)

						if True == tmp2153 {
							__e.TailApply(PrimFunc(symshen_4parse_1failure))
							return
						} else {
							tmp2123 := MakeNative(func(__e *ControlFlow) {
								W659 := __e.Get(1)
								_ = W659
								tmp2124 := MakeNative(func(__e *ControlFlow) {
									W660 := __e.Get(1)
									_ = W660
									tmp2125 := MakeNative(func(__e *ControlFlow) {
										W661 := __e.Get(1)
										_ = W661
										tmp2148 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W661)

										if True == tmp2148 {
											__e.TailApply(PrimFunc(symshen_4parse_1failure))
											return
										} else {
											tmp2126 := MakeNative(func(__e *ControlFlow) {
												W662 := __e.Get(1)
												_ = W662
												tmp2127 := MakeNative(func(__e *ControlFlow) {
													W663 := __e.Get(1)
													_ = W663
													tmp2128 := MakeNative(func(__e *ControlFlow) {
														W664 := __e.Get(1)
														_ = W664
														tmp2143 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W664)

														if True == tmp2143 {
															__e.TailApply(PrimFunc(symshen_4parse_1failure))
															return
														} else {
															tmp2129 := MakeNative(func(__e *ControlFlow) {
																W665 := __e.Get(1)
																_ = W665
																tmp2130 := MakeNative(func(__e *ControlFlow) {
																	W666 := __e.Get(1)
																	_ = W666
																	tmp2131 := MakeNative(func(__e *ControlFlow) {
																		W667 := __e.Get(1)
																		_ = W667
																		tmp2138 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W667)

																		if True == tmp2138 {
																			__e.TailApply(PrimFunc(symshen_4parse_1failure))
																			return
																		} else {
																			tmp2132 := MakeNative(func(__e *ControlFlow) {
																				W668 := __e.Get(1)
																				_ = W668
																				tmp2133 := PrimCons(W665, Nil)

																				tmp2134 := PrimCons(W662, tmp2133)

																				tmp2135 := PrimCons(W659, tmp2134)

																				__e.TailApply(PrimFunc(symshen_4comb), W657, tmp2135)
																				return

																			}, 1)

																			tmp2136 := Call(__e, PrimFunc(symshen_4in_1_6), W667)

																			__e.TailApply(tmp2132, tmp2136)
																			return

																		}

																	}, 1)

																	tmp2139 := Call(__e, PrimFunc(sym_5end_6), W666)

																	__e.TailApply(tmp2131, tmp2139)
																	return

																}, 1)

																tmp2140 := Call(__e, PrimFunc(symshen_4in_1_6), W664)

																__e.TailApply(tmp2130, tmp2140)
																return

															}, 1)

															tmp2141 := Call(__e, PrimFunc(symshen_4_5_1out), W664)

															__e.TailApply(tmp2129, tmp2141)
															return

														}

													}, 1)

													tmp2144 := Call(__e, PrimFunc(symshen_4_5pattern2_6), W663)

													__e.TailApply(tmp2128, tmp2144)
													return

												}, 1)

												tmp2145 := Call(__e, PrimFunc(symshen_4in_1_6), W661)

												__e.TailApply(tmp2127, tmp2145)
												return

											}, 1)

											tmp2146 := Call(__e, PrimFunc(symshen_4_5_1out), W661)

											__e.TailApply(tmp2126, tmp2146)
											return

										}

									}, 1)

									tmp2149 := Call(__e, PrimFunc(symshen_4_5pattern1_6), W660)

									__e.TailApply(tmp2125, tmp2149)
									return

								}, 1)

								tmp2150 := Call(__e, PrimFunc(symshen_4in_1_6), W658)

								__e.TailApply(tmp2124, tmp2150)
								return

							}, 1)

							tmp2151 := Call(__e, PrimFunc(symshen_4_5_1out), W658)

							__e.TailApply(tmp2123, tmp2151)
							return

						}

					}, 1)

					tmp2154 := Call(__e, PrimFunc(symshen_4_5constructor_6), W656)

					__e.TailApply(tmp2122, tmp2154)
					return

				}, 1)

				tmp2155 := Call(__e, PrimFunc(symtail), V654)

				__e.TailApply(tmp2121, tmp2155)
				return

			}, 1)

			tmp2156 := Call(__e, PrimFunc(symhead), V654)

			tmp2157 := Call(__e, tmp2120, tmp2156)

			ifres2119 = tmp2157

		} else {
			tmp2158 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres2119 = tmp2158

		}

		__e.TailApply(tmp2063, ifres2119)
		return

	}, 1)

	tmp2160 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp2062)

	_ = tmp2160

	tmp2161 := MakeNative(func(__e *ControlFlow) {
		V683 := __e.Get(1)
		_ = V683
		tmp2162 := MakeNative(func(__e *ControlFlow) {
			W684 := __e.Get(1)
			_ = W684
			tmp2164 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W684)

			if True == tmp2164 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W684)
				return
			}

		}, 1)

		tmp2174 := PrimIsPair(V683)

		var ifres2165 Obj

		if True == tmp2174 {
			tmp2166 := MakeNative(func(__e *ControlFlow) {
				W685 := __e.Get(1)
				_ = W685
				tmp2167 := MakeNative(func(__e *ControlFlow) {
					W686 := __e.Get(1)
					_ = W686
					tmp2169 := Call(__e, PrimFunc(symshen_4constructor_2), W685)

					if True == tmp2169 {
						__e.TailApply(PrimFunc(symshen_4comb), W686, W685)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp2170 := Call(__e, PrimFunc(symtail), V683)

				__e.TailApply(tmp2167, tmp2170)
				return

			}, 1)

			tmp2171 := Call(__e, PrimFunc(symhead), V683)

			tmp2172 := Call(__e, tmp2166, tmp2171)

			ifres2165 = tmp2172

		} else {
			tmp2173 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres2165 = tmp2173

		}

		__e.TailApply(tmp2162, ifres2165)
		return

	}, 1)

	tmp2175 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp2161)

	_ = tmp2175

	tmp2176 := MakeNative(func(__e *ControlFlow) {
		V687 := __e.Get(1)
		_ = V687
		tmp2177 := PrimCons(sym_8v, Nil)

		tmp2178 := PrimCons(sym_8s, tmp2177)

		tmp2179 := PrimCons(sym_8p, tmp2178)

		tmp2180 := PrimCons(symcons, tmp2179)

		__e.TailApply(PrimFunc(symelement_2), V687, tmp2180)
		return

	}, 1)

	tmp2181 := Call(__e, ns2_1set, symshen_4constructor_2, tmp2176)

	_ = tmp2181

	tmp2182 := MakeNative(func(__e *ControlFlow) {
		V688 := __e.Get(1)
		_ = V688
		tmp2183 := Call(__e, PrimFunc(symshen_4app), V688, MakeString(" is not a legitimate constructor\n"), symshen_4r)

		__e.Return(PrimSimpleError(tmp2183))
		return

	}, 1)

	tmp2184 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp2182)

	_ = tmp2184

	tmp2185 := MakeNative(func(__e *ControlFlow) {
		V689 := __e.Get(1)
		_ = V689
		tmp2186 := MakeNative(func(__e *ControlFlow) {
			W690 := __e.Get(1)
			_ = W690
			tmp2204 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W690)

			if True == tmp2204 {
				tmp2187 := MakeNative(func(__e *ControlFlow) {
					W693 := __e.Get(1)
					_ = W693
					tmp2189 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W693)

					if True == tmp2189 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(W693)
						return
					}

				}, 1)

				tmp2202 := PrimIsPair(V689)

				var ifres2190 Obj

				if True == tmp2202 {
					tmp2191 := MakeNative(func(__e *ControlFlow) {
						W694 := __e.Get(1)
						_ = W694
						tmp2192 := MakeNative(func(__e *ControlFlow) {
							W695 := __e.Get(1)
							_ = W695
							tmp2194 := PrimCons(sym_5_1, Nil)

							tmp2195 := PrimCons(sym_1_6, tmp2194)

							tmp2196 := Call(__e, PrimFunc(symelement_2), W694, tmp2195)

							tmp2197 := PrimNot(tmp2196)

							if True == tmp2197 {
								__e.TailApply(PrimFunc(symshen_4comb), W695, W694)
								return
							} else {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							}

						}, 1)

						tmp2198 := Call(__e, PrimFunc(symtail), V689)

						__e.TailApply(tmp2192, tmp2198)
						return

					}, 1)

					tmp2199 := Call(__e, PrimFunc(symhead), V689)

					tmp2200 := Call(__e, tmp2191, tmp2199)

					ifres2190 = tmp2200

				} else {
					tmp2201 := Call(__e, PrimFunc(symshen_4parse_1failure))

					ifres2190 = tmp2201

				}

				__e.TailApply(tmp2187, ifres2190)
				return

			} else {
				__e.Return(W690)
				return
			}

		}, 1)

		tmp2215 := PrimIsPair(V689)

		var ifres2205 Obj

		if True == tmp2215 {
			tmp2206 := MakeNative(func(__e *ControlFlow) {
				W691 := __e.Get(1)
				_ = W691
				tmp2207 := MakeNative(func(__e *ControlFlow) {
					W692 := __e.Get(1)
					_ = W692
					tmp2210 := PrimEqual(W691, sym__)

					if True == tmp2210 {
						tmp2208 := Call(__e, PrimFunc(symgensym), symY)

						__e.TailApply(PrimFunc(symshen_4comb), W692, tmp2208)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp2211 := Call(__e, PrimFunc(symtail), V689)

				__e.TailApply(tmp2207, tmp2211)
				return

			}, 1)

			tmp2212 := Call(__e, PrimFunc(symhead), V689)

			tmp2213 := Call(__e, tmp2206, tmp2212)

			ifres2205 = tmp2213

		} else {
			tmp2214 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres2205 = tmp2214

		}

		__e.TailApply(tmp2186, ifres2205)
		return

	}, 1)

	tmp2216 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp2185)

	_ = tmp2216

	tmp2217 := MakeNative(func(__e *ControlFlow) {
		V696 := __e.Get(1)
		_ = V696
		tmp2218 := MakeNative(func(__e *ControlFlow) {
			W697 := __e.Get(1)
			_ = W697
			tmp2220 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W697)

			if True == tmp2220 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W697)
				return
			}

		}, 1)

		tmp2221 := MakeNative(func(__e *ControlFlow) {
			W698 := __e.Get(1)
			_ = W698
			tmp2227 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W698)

			if True == tmp2227 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp2222 := MakeNative(func(__e *ControlFlow) {
					W699 := __e.Get(1)
					_ = W699
					tmp2223 := MakeNative(func(__e *ControlFlow) {
						W700 := __e.Get(1)
						_ = W700
						__e.TailApply(PrimFunc(symshen_4comb), W700, W699)
						return
					}, 1)

					tmp2224 := Call(__e, PrimFunc(symshen_4in_1_6), W698)

					__e.TailApply(tmp2223, tmp2224)
					return

				}, 1)

				tmp2225 := Call(__e, PrimFunc(symshen_4_5_1out), W698)

				__e.TailApply(tmp2222, tmp2225)
				return

			}

		}, 1)

		tmp2228 := Call(__e, PrimFunc(symshen_4_5pattern_6), V696)

		tmp2229 := Call(__e, tmp2221, tmp2228)

		__e.TailApply(tmp2218, tmp2229)
		return

	}, 1)

	tmp2230 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp2217)

	_ = tmp2230

	tmp2231 := MakeNative(func(__e *ControlFlow) {
		V701 := __e.Get(1)
		_ = V701
		tmp2232 := MakeNative(func(__e *ControlFlow) {
			W702 := __e.Get(1)
			_ = W702
			tmp2234 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W702)

			if True == tmp2234 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(W702)
				return
			}

		}, 1)

		tmp2235 := MakeNative(func(__e *ControlFlow) {
			W703 := __e.Get(1)
			_ = W703
			tmp2241 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W703)

			if True == tmp2241 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp2236 := MakeNative(func(__e *ControlFlow) {
					W704 := __e.Get(1)
					_ = W704
					tmp2237 := MakeNative(func(__e *ControlFlow) {
						W705 := __e.Get(1)
						_ = W705
						__e.TailApply(PrimFunc(symshen_4comb), W705, W704)
						return
					}, 1)

					tmp2238 := Call(__e, PrimFunc(symshen_4in_1_6), W703)

					__e.TailApply(tmp2237, tmp2238)
					return

				}, 1)

				tmp2239 := Call(__e, PrimFunc(symshen_4_5_1out), W703)

				__e.TailApply(tmp2236, tmp2239)
				return

			}

		}, 1)

		tmp2242 := Call(__e, PrimFunc(symshen_4_5pattern_6), V701)

		tmp2243 := Call(__e, tmp2235, tmp2242)

		__e.TailApply(tmp2232, tmp2243)
		return

	}, 1)

	tmp2244 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp2231)

	_ = tmp2244

	tmp2245 := MakeNative(func(__e *ControlFlow) {
		V706 := __e.Get(1)
		_ = V706
		tmp2246 := MakeNative(func(__e *ControlFlow) {
			W707 := __e.Get(1)
			_ = W707
			tmp2247 := MakeNative(func(__e *ControlFlow) {
				W708 := __e.Get(1)
				_ = W708
				tmp2248 := MakeNative(func(__e *ControlFlow) {
					W709 := __e.Get(1)
					_ = W709
					__e.Return(W709)
					return
				}, 1)

				tmp2249 := PrimStr(V706)

				tmp2250 := Call(__e, PrimFunc(sym_8s), tmp2249, MakeString(")"))

				tmp2251 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2250)

				tmp2252 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2251)

				tmp2253 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2252)

				tmp2254 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp2253)

				tmp2255 := PrimVectorSet(W708, MakeNumber(1), tmp2254)

				__e.TailApply(tmp2248, tmp2255)
				return

			}, 1)

			tmp2256 := PrimVectorSet(W707, MakeNumber(0), symshen_4printF)

			__e.TailApply(tmp2247, tmp2256)
			return

		}, 1)

		tmp2257 := PrimAbsvector(MakeNumber(2))

		__e.TailApply(tmp2246, tmp2257)
		return

	}, 1)

	tmp2258 := Call(__e, ns2_1set, symshen_4fn_1print, tmp2245)

	_ = tmp2258

	tmp2259 := MakeNative(func(__e *ControlFlow) {
		V710 := __e.Get(1)
		_ = V710
		__e.Return(PrimVectorGet(V710, MakeNumber(1)))
		return
	}, 1)

	tmp2260 := Call(__e, ns2_1set, symshen_4printF, tmp2259)

	_ = tmp2260

	tmp2261 := MakeNative(func(__e *ControlFlow) {
		V715 := __e.Get(1)
		_ = V715
		V716 := __e.Get(2)
		_ = V716
		tmp2285 := PrimIsPair(V716)

		var ifres2281 Obj

		if True == tmp2285 {
			tmp2283 := PrimTail(V716)

			tmp2284 := PrimEqual(Nil, tmp2283)

			var ifres2282 Obj

			if True == tmp2284 {
				ifres2282 = True

			} else {
				ifres2282 = False

			}

			ifres2281 = ifres2282

		} else {
			ifres2281 = False

		}

		if True == ifres2281 {
			tmp2262 := PrimHead(V716)

			__e.TailApply(PrimFunc(symlength), tmp2262)
			return

		} else {
			tmp2279 := PrimIsPair(V716)

			var ifres2267 Obj

			if True == tmp2279 {
				tmp2277 := PrimTail(V716)

				tmp2278 := PrimIsPair(tmp2277)

				var ifres2269 Obj

				if True == tmp2278 {
					tmp2271 := PrimHead(V716)

					tmp2272 := Call(__e, PrimFunc(symlength), tmp2271)

					tmp2273 := PrimTail(V716)

					tmp2274 := PrimHead(tmp2273)

					tmp2275 := Call(__e, PrimFunc(symlength), tmp2274)

					tmp2276 := PrimEqual(tmp2272, tmp2275)

					var ifres2270 Obj

					if True == tmp2276 {
						ifres2270 = True

					} else {
						ifres2270 = False

					}

					ifres2269 = ifres2270

				} else {
					ifres2269 = False

				}

				var ifres2268 Obj

				if True == ifres2269 {
					ifres2268 = True

				} else {
					ifres2268 = False

				}

				ifres2267 = ifres2268

			} else {
				ifres2267 = False

			}

			if True == ifres2267 {
				tmp2263 := PrimTail(V716)

				__e.TailApply(PrimFunc(symshen_4arity_1chk), V715, tmp2263)
				return

			} else {
				tmp2264 := Call(__e, PrimFunc(symshen_4app), V715, MakeString("\n"), symshen_4a)

				tmp2265 := PrimStringConcat(MakeString("arity error in "), tmp2264)

				__e.Return(PrimSimpleError(tmp2265))
				return

			}

		}

	}, 2)

	tmp2286 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp2261)

	_ = tmp2286

	tmp2287 := MakeNative(func(__e *ControlFlow) {
		V717 := __e.Get(1)
		_ = V717
		V718 := __e.Get(2)
		_ = V718
		tmp2293 := Call(__e, PrimFunc(symtuple_2), V718)

		if True == tmp2293 {
			tmp2288 := Call(__e, PrimFunc(symfst), V718)

			tmp2289 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp2288)

			tmp2290 := Call(__e, PrimFunc(symsnd), V718)

			tmp2291 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp2289, tmp2290)

			__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V717, tmp2291)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4free_1var_1chk)
			return
		}

	}, 2)

	tmp2294 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp2287)

	_ = tmp2294

	tmp2295 := MakeNative(func(__e *ControlFlow) {
		V719 := __e.Get(1)
		_ = V719
		V720 := __e.Get(2)
		_ = V720
		tmp2307 := Call(__e, PrimFunc(symempty_2), V720)

		if True == tmp2307 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp2296 := Call(__e, PrimFunc(symshen_4app), V719, MakeString(":"), symshen_4a)

			tmp2297 := PrimStringConcat(MakeString("free variables in "), tmp2296)

			tmp2298 := Call(__e, PrimFunc(symstoutput))

			tmp2299 := Call(__e, PrimFunc(sympr), tmp2297, tmp2298)

			_ = tmp2299

			tmp2300 := MakeNative(func(__e *ControlFlow) {
				Z721 := __e.Get(1)
				_ = Z721
				tmp2301 := Call(__e, PrimFunc(symshen_4app), Z721, MakeString(""), symshen_4a)

				tmp2302 := PrimStringConcat(MakeString(" "), tmp2301)

				tmp2303 := Call(__e, PrimFunc(symstoutput))

				__e.TailApply(PrimFunc(sympr), tmp2302, tmp2303)
				return

			}, 1)

			tmp2304 := Call(__e, PrimFunc(symshen_4for_1each), tmp2300, V720)

			_ = tmp2304

			tmp2305 := Call(__e, PrimFunc(symnl), MakeNumber(1))

			_ = tmp2305

			__e.TailApply(PrimFunc(symabort))
			return

		}

	}, 2)

	tmp2308 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp2295)

	_ = tmp2308

	tmp2309 := MakeNative(func(__e *ControlFlow) {
		V724 := __e.Get(1)
		_ = V724
		tmp2317 := PrimIsVariable(V724)

		if True == tmp2317 {
			__e.Return(PrimCons(V724, Nil))
			return
		} else {
			tmp2315 := PrimIsPair(V724)

			if True == tmp2315 {
				tmp2310 := PrimHead(V724)

				tmp2311 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp2310)

				tmp2312 := PrimTail(V724)

				tmp2313 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp2312)

				__e.TailApply(PrimFunc(symunion), tmp2311, tmp2313)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp2318 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp2309)

	_ = tmp2318

	tmp2319 := MakeNative(func(__e *ControlFlow) {
		V729 := __e.Get(1)
		_ = V729
		V730 := __e.Get(2)
		_ = V730
		tmp2409 := PrimIsPair(V730)

		var ifres2396 Obj

		if True == tmp2409 {
			tmp2407 := PrimHead(V730)

			tmp2408 := PrimEqual(symprotect, tmp2407)

			var ifres2398 Obj

			if True == tmp2408 {
				tmp2405 := PrimTail(V730)

				tmp2406 := PrimIsPair(tmp2405)

				var ifres2400 Obj

				if True == tmp2406 {
					tmp2402 := PrimTail(V730)

					tmp2403 := PrimTail(tmp2402)

					tmp2404 := PrimEqual(Nil, tmp2403)

					var ifres2401 Obj

					if True == tmp2404 {
						ifres2401 = True

					} else {
						ifres2401 = False

					}

					ifres2400 = ifres2401

				} else {
					ifres2400 = False

				}

				var ifres2399 Obj

				if True == ifres2400 {
					ifres2399 = True

				} else {
					ifres2399 = False

				}

				ifres2398 = ifres2399

			} else {
				ifres2398 = False

			}

			var ifres2397 Obj

			if True == ifres2398 {
				ifres2397 = True

			} else {
				ifres2397 = False

			}

			ifres2396 = ifres2397

		} else {
			ifres2396 = False

		}

		if True == ifres2396 {
			__e.Return(Nil)
			return
		} else {
			tmp2394 := PrimIsPair(V730)

			var ifres2368 Obj

			if True == tmp2394 {
				tmp2392 := PrimHead(V730)

				tmp2393 := PrimEqual(symlet, tmp2392)

				var ifres2370 Obj

				if True == tmp2393 {
					tmp2390 := PrimTail(V730)

					tmp2391 := PrimIsPair(tmp2390)

					var ifres2372 Obj

					if True == tmp2391 {
						tmp2387 := PrimTail(V730)

						tmp2388 := PrimTail(tmp2387)

						tmp2389 := PrimIsPair(tmp2388)

						var ifres2374 Obj

						if True == tmp2389 {
							tmp2383 := PrimTail(V730)

							tmp2384 := PrimTail(tmp2383)

							tmp2385 := PrimTail(tmp2384)

							tmp2386 := PrimIsPair(tmp2385)

							var ifres2376 Obj

							if True == tmp2386 {
								tmp2378 := PrimTail(V730)

								tmp2379 := PrimTail(tmp2378)

								tmp2380 := PrimTail(tmp2379)

								tmp2381 := PrimTail(tmp2380)

								tmp2382 := PrimEqual(Nil, tmp2381)

								var ifres2377 Obj

								if True == tmp2382 {
									ifres2377 = True

								} else {
									ifres2377 = False

								}

								ifres2376 = ifres2377

							} else {
								ifres2376 = False

							}

							var ifres2375 Obj

							if True == ifres2376 {
								ifres2375 = True

							} else {
								ifres2375 = False

							}

							ifres2374 = ifres2375

						} else {
							ifres2374 = False

						}

						var ifres2373 Obj

						if True == ifres2374 {
							ifres2373 = True

						} else {
							ifres2373 = False

						}

						ifres2372 = ifres2373

					} else {
						ifres2372 = False

					}

					var ifres2371 Obj

					if True == ifres2372 {
						ifres2371 = True

					} else {
						ifres2371 = False

					}

					ifres2370 = ifres2371

				} else {
					ifres2370 = False

				}

				var ifres2369 Obj

				if True == ifres2370 {
					ifres2369 = True

				} else {
					ifres2369 = False

				}

				ifres2368 = ifres2369

			} else {
				ifres2368 = False

			}

			if True == ifres2368 {
				tmp2320 := PrimTail(V730)

				tmp2321 := PrimTail(tmp2320)

				tmp2322 := PrimHead(tmp2321)

				tmp2323 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V729, tmp2322)

				tmp2324 := PrimTail(V730)

				tmp2325 := PrimHead(tmp2324)

				tmp2326 := PrimCons(tmp2325, V729)

				tmp2327 := PrimTail(V730)

				tmp2328 := PrimTail(tmp2327)

				tmp2329 := PrimTail(tmp2328)

				tmp2330 := PrimHead(tmp2329)

				tmp2331 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp2326, tmp2330)

				__e.TailApply(PrimFunc(symunion), tmp2323, tmp2331)
				return

			} else {
				tmp2366 := PrimIsPair(V730)

				var ifres2347 Obj

				if True == tmp2366 {
					tmp2364 := PrimHead(V730)

					tmp2365 := PrimEqual(symlambda, tmp2364)

					var ifres2349 Obj

					if True == tmp2365 {
						tmp2362 := PrimTail(V730)

						tmp2363 := PrimIsPair(tmp2362)

						var ifres2351 Obj

						if True == tmp2363 {
							tmp2359 := PrimTail(V730)

							tmp2360 := PrimTail(tmp2359)

							tmp2361 := PrimIsPair(tmp2360)

							var ifres2353 Obj

							if True == tmp2361 {
								tmp2355 := PrimTail(V730)

								tmp2356 := PrimTail(tmp2355)

								tmp2357 := PrimTail(tmp2356)

								tmp2358 := PrimEqual(Nil, tmp2357)

								var ifres2354 Obj

								if True == tmp2358 {
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

						var ifres2350 Obj

						if True == ifres2351 {
							ifres2350 = True

						} else {
							ifres2350 = False

						}

						ifres2349 = ifres2350

					} else {
						ifres2349 = False

					}

					var ifres2348 Obj

					if True == ifres2349 {
						ifres2348 = True

					} else {
						ifres2348 = False

					}

					ifres2347 = ifres2348

				} else {
					ifres2347 = False

				}

				if True == ifres2347 {
					tmp2332 := PrimTail(V730)

					tmp2333 := PrimHead(tmp2332)

					tmp2334 := PrimCons(tmp2333, V729)

					tmp2335 := PrimTail(V730)

					tmp2336 := PrimTail(tmp2335)

					tmp2337 := PrimHead(tmp2336)

					__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp2334, tmp2337)
					return

				} else {
					tmp2345 := PrimIsPair(V730)

					if True == tmp2345 {
						tmp2338 := PrimHead(V730)

						tmp2339 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V729, tmp2338)

						tmp2340 := PrimTail(V730)

						tmp2341 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V729, tmp2340)

						__e.TailApply(PrimFunc(symunion), tmp2339, tmp2341)
						return

					} else {
						tmp2343 := Call(__e, PrimFunc(symshen_4free_1variable_2), V730, V729)

						if True == tmp2343 {
							__e.Return(PrimCons(V730, Nil))
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

	tmp2410 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp2319)

	_ = tmp2410

	tmp2411 := MakeNative(func(__e *ControlFlow) {
		V731 := __e.Get(1)
		_ = V731
		V732 := __e.Get(2)
		_ = V732
		tmp2416 := PrimIsVariable(V731)

		if True == tmp2416 {
			tmp2413 := Call(__e, PrimFunc(symelement_2), V731, V732)

			tmp2414 := PrimNot(tmp2413)

			if True == tmp2414 {
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

	tmp2417 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp2411)

	_ = tmp2417

	tmp2418 := MakeNative(func(__e *ControlFlow) {
		V733 := __e.Get(1)
		_ = V733
		V734 := __e.Get(2)
		_ = V734
		tmp2419 := PrimValue(symshen_4_duserdefs_d)

		tmp2420 := Call(__e, PrimFunc(symadjoin), V733, tmp2419)

		tmp2421 := PrimSet(symshen_4_duserdefs_d, tmp2420)

		_ = tmp2421

		tmp2422 := PrimValue(sym_dproperty_1vector_d)

		__e.TailApply(PrimFunc(symput), V733, symshen_4source, V734, tmp2422)
		return

	}, 2)

	tmp2423 := Call(__e, ns2_1set, symshen_4record_1kl, tmp2418)

	_ = tmp2423

	tmp2424 := MakeNative(func(__e *ControlFlow) {
		V735 := __e.Get(1)
		_ = V735
		V736 := __e.Get(2)
		_ = V736
		V737 := __e.Get(3)
		_ = V737
		tmp2425 := MakeNative(func(__e *ControlFlow) {
			W738 := __e.Get(1)
			_ = W738
			tmp2426 := MakeNative(func(__e *ControlFlow) {
				W739 := __e.Get(1)
				_ = W739
				tmp2427 := MakeNative(func(__e *ControlFlow) {
					W740 := __e.Get(1)
					_ = W740
					__e.Return(W740)
					return
				}, 1)

				tmp2428 := Call(__e, PrimFunc(symshen_4cond_1form), W739)

				tmp2429 := PrimCons(tmp2428, Nil)

				tmp2430 := PrimCons(W738, tmp2429)

				tmp2431 := PrimCons(V735, tmp2430)

				tmp2432 := PrimCons(symdefun, tmp2431)

				__e.TailApply(tmp2427, tmp2432)
				return

			}, 1)

			tmp2433 := Call(__e, PrimFunc(symshen_4kl_1body), V736, W738)

			tmp2434 := Call(__e, PrimFunc(symshen_4scan_1body), V735, tmp2433)

			__e.TailApply(tmp2426, tmp2434)
			return

		}, 1)

		tmp2435 := Call(__e, PrimFunc(symshen_4parameters), V737)

		__e.TailApply(tmp2425, tmp2435)
		return

	}, 3)

	tmp2436 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp2424)

	_ = tmp2436

	tmp2437 := MakeNative(func(__e *ControlFlow) {
		V741 := __e.Get(1)
		_ = V741
		tmp2442 := PrimEqual(MakeNumber(0), V741)

		if True == tmp2442 {
			__e.Return(Nil)
			return
		} else {
			tmp2438 := Call(__e, PrimFunc(symgensym), symV)

			tmp2439 := PrimNumberSubtract(V741, MakeNumber(1))

			tmp2440 := Call(__e, PrimFunc(symshen_4parameters), tmp2439)

			__e.Return(PrimCons(tmp2438, tmp2440))
			return

		}

	}, 1)

	tmp2443 := Call(__e, ns2_1set, symshen_4parameters, tmp2437)

	_ = tmp2443

	tmp2444 := MakeNative(func(__e *ControlFlow) {
		V744 := __e.Get(1)
		_ = V744
		tmp2468 := PrimIsPair(V744)

		var ifres2448 Obj

		if True == tmp2468 {
			tmp2466 := PrimHead(V744)

			tmp2467 := PrimIsPair(tmp2466)

			var ifres2450 Obj

			if True == tmp2467 {
				tmp2463 := PrimHead(V744)

				tmp2464 := PrimHead(tmp2463)

				tmp2465 := PrimEqual(True, tmp2464)

				var ifres2452 Obj

				if True == tmp2465 {
					tmp2460 := PrimHead(V744)

					tmp2461 := PrimTail(tmp2460)

					tmp2462 := PrimIsPair(tmp2461)

					var ifres2454 Obj

					if True == tmp2462 {
						tmp2456 := PrimHead(V744)

						tmp2457 := PrimTail(tmp2456)

						tmp2458 := PrimTail(tmp2457)

						tmp2459 := PrimEqual(Nil, tmp2458)

						var ifres2455 Obj

						if True == tmp2459 {
							ifres2455 = True

						} else {
							ifres2455 = False

						}

						ifres2454 = ifres2455

					} else {
						ifres2454 = False

					}

					var ifres2453 Obj

					if True == ifres2454 {
						ifres2453 = True

					} else {
						ifres2453 = False

					}

					ifres2452 = ifres2453

				} else {
					ifres2452 = False

				}

				var ifres2451 Obj

				if True == ifres2452 {
					ifres2451 = True

				} else {
					ifres2451 = False

				}

				ifres2450 = ifres2451

			} else {
				ifres2450 = False

			}

			var ifres2449 Obj

			if True == ifres2450 {
				ifres2449 = True

			} else {
				ifres2449 = False

			}

			ifres2448 = ifres2449

		} else {
			ifres2448 = False

		}

		if True == ifres2448 {
			tmp2445 := PrimHead(V744)

			tmp2446 := PrimTail(tmp2445)

			__e.Return(PrimHead(tmp2446))
			return

		} else {
			__e.Return(PrimCons(symcond, V744))
			return
		}

	}, 1)

	tmp2469 := Call(__e, ns2_1set, symshen_4cond_1form, tmp2444)

	_ = tmp2469

	tmp2470 := MakeNative(func(__e *ControlFlow) {
		V753 := __e.Get(1)
		_ = V753
		V754 := __e.Get(2)
		_ = V754
		tmp2514 := PrimEqual(Nil, V754)

		if True == tmp2514 {
			tmp2471 := PrimCons(V753, Nil)

			tmp2472 := PrimCons(symshen_4f_1error, tmp2471)

			tmp2473 := PrimCons(tmp2472, Nil)

			tmp2474 := PrimCons(True, tmp2473)

			__e.Return(PrimCons(tmp2474, Nil))
			return

		} else {
			tmp2512 := PrimIsPair(V754)

			var ifres2508 Obj

			if True == tmp2512 {
				tmp2510 := PrimHead(V754)

				tmp2511 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp2510)

				var ifres2509 Obj

				if True == tmp2511 {
					ifres2509 = True

				} else {
					ifres2509 = False

				}

				ifres2508 = ifres2509

			} else {
				ifres2508 = False

			}

			if True == ifres2508 {
				tmp2475 := Call(__e, PrimFunc(symgensym), symFreeze)

				tmp2476 := Call(__e, PrimFunc(symgensym), symResult)

				tmp2477 := PrimHead(V754)

				tmp2478 := PrimTail(V754)

				__e.TailApply(PrimFunc(symshen_4choicepoint), V753, tmp2475, tmp2476, tmp2477, tmp2478)
				return

			} else {
				tmp2506 := PrimIsPair(V754)

				var ifres2486 Obj

				if True == tmp2506 {
					tmp2504 := PrimHead(V754)

					tmp2505 := PrimIsPair(tmp2504)

					var ifres2488 Obj

					if True == tmp2505 {
						tmp2501 := PrimHead(V754)

						tmp2502 := PrimHead(tmp2501)

						tmp2503 := PrimEqual(True, tmp2502)

						var ifres2490 Obj

						if True == tmp2503 {
							tmp2498 := PrimHead(V754)

							tmp2499 := PrimTail(tmp2498)

							tmp2500 := PrimIsPair(tmp2499)

							var ifres2492 Obj

							if True == tmp2500 {
								tmp2494 := PrimHead(V754)

								tmp2495 := PrimTail(tmp2494)

								tmp2496 := PrimTail(tmp2495)

								tmp2497 := PrimEqual(Nil, tmp2496)

								var ifres2493 Obj

								if True == tmp2497 {
									ifres2493 = True

								} else {
									ifres2493 = False

								}

								ifres2492 = ifres2493

							} else {
								ifres2492 = False

							}

							var ifres2491 Obj

							if True == ifres2492 {
								ifres2491 = True

							} else {
								ifres2491 = False

							}

							ifres2490 = ifres2491

						} else {
							ifres2490 = False

						}

						var ifres2489 Obj

						if True == ifres2490 {
							ifres2489 = True

						} else {
							ifres2489 = False

						}

						ifres2488 = ifres2489

					} else {
						ifres2488 = False

					}

					var ifres2487 Obj

					if True == ifres2488 {
						ifres2487 = True

					} else {
						ifres2487 = False

					}

					ifres2486 = ifres2487

				} else {
					ifres2486 = False

				}

				if True == ifres2486 {
					tmp2479 := PrimHead(V754)

					__e.Return(PrimCons(tmp2479, Nil))
					return

				} else {
					tmp2484 := PrimIsPair(V754)

					if True == tmp2484 {
						tmp2480 := PrimHead(V754)

						tmp2481 := PrimTail(V754)

						tmp2482 := Call(__e, PrimFunc(symshen_4scan_1body), V753, tmp2481)

						__e.Return(PrimCons(tmp2480, tmp2482))
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.scan-body")))
						return
					}

				}

			}

		}

	}, 2)

	tmp2515 := Call(__e, ns2_1set, symshen_4scan_1body, tmp2470)

	_ = tmp2515

	tmp2516 := MakeNative(func(__e *ControlFlow) {
		V761 := __e.Get(1)
		_ = V761
		tmp2551 := PrimIsPair(V761)

		var ifres2518 Obj

		if True == tmp2551 {
			tmp2549 := PrimTail(V761)

			tmp2550 := PrimIsPair(tmp2549)

			var ifres2520 Obj

			if True == tmp2550 {
				tmp2546 := PrimTail(V761)

				tmp2547 := PrimHead(tmp2546)

				tmp2548 := PrimIsPair(tmp2547)

				var ifres2522 Obj

				if True == tmp2548 {
					tmp2542 := PrimTail(V761)

					tmp2543 := PrimHead(tmp2542)

					tmp2544 := PrimHead(tmp2543)

					tmp2545 := PrimEqual(symshen_4choicepoint_b, tmp2544)

					var ifres2524 Obj

					if True == tmp2545 {
						tmp2538 := PrimTail(V761)

						tmp2539 := PrimHead(tmp2538)

						tmp2540 := PrimTail(tmp2539)

						tmp2541 := PrimIsPair(tmp2540)

						var ifres2526 Obj

						if True == tmp2541 {
							tmp2533 := PrimTail(V761)

							tmp2534 := PrimHead(tmp2533)

							tmp2535 := PrimTail(tmp2534)

							tmp2536 := PrimTail(tmp2535)

							tmp2537 := PrimEqual(Nil, tmp2536)

							var ifres2528 Obj

							if True == tmp2537 {
								tmp2530 := PrimTail(V761)

								tmp2531 := PrimTail(tmp2530)

								tmp2532 := PrimEqual(Nil, tmp2531)

								var ifres2529 Obj

								if True == tmp2532 {
									ifres2529 = True

								} else {
									ifres2529 = False

								}

								ifres2528 = ifres2529

							} else {
								ifres2528 = False

							}

							var ifres2527 Obj

							if True == ifres2528 {
								ifres2527 = True

							} else {
								ifres2527 = False

							}

							ifres2526 = ifres2527

						} else {
							ifres2526 = False

						}

						var ifres2525 Obj

						if True == ifres2526 {
							ifres2525 = True

						} else {
							ifres2525 = False

						}

						ifres2524 = ifres2525

					} else {
						ifres2524 = False

					}

					var ifres2523 Obj

					if True == ifres2524 {
						ifres2523 = True

					} else {
						ifres2523 = False

					}

					ifres2522 = ifres2523

				} else {
					ifres2522 = False

				}

				var ifres2521 Obj

				if True == ifres2522 {
					ifres2521 = True

				} else {
					ifres2521 = False

				}

				ifres2520 = ifres2521

			} else {
				ifres2520 = False

			}

			var ifres2519 Obj

			if True == ifres2520 {
				ifres2519 = True

			} else {
				ifres2519 = False

			}

			ifres2518 = ifres2519

		} else {
			ifres2518 = False

		}

		if True == ifres2518 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp2552 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp2516)

	_ = tmp2552

	tmp2553 := MakeNative(func(__e *ControlFlow) {
		V777 := __e.Get(1)
		_ = V777
		V778 := __e.Get(2)
		_ = V778
		V779 := __e.Get(3)
		_ = V779
		V780 := __e.Get(4)
		_ = V780
		V781 := __e.Get(5)
		_ = V781
		tmp2745 := PrimIsPair(V780)

		var ifres2667 Obj

		if True == tmp2745 {
			tmp2743 := PrimTail(V780)

			tmp2744 := PrimIsPair(tmp2743)

			var ifres2669 Obj

			if True == tmp2744 {
				tmp2740 := PrimTail(V780)

				tmp2741 := PrimHead(tmp2740)

				tmp2742 := PrimIsPair(tmp2741)

				var ifres2671 Obj

				if True == tmp2742 {
					tmp2736 := PrimTail(V780)

					tmp2737 := PrimHead(tmp2736)

					tmp2738 := PrimTail(tmp2737)

					tmp2739 := PrimIsPair(tmp2738)

					var ifres2673 Obj

					if True == tmp2739 {
						tmp2731 := PrimTail(V780)

						tmp2732 := PrimHead(tmp2731)

						tmp2733 := PrimTail(tmp2732)

						tmp2734 := PrimHead(tmp2733)

						tmp2735 := PrimIsPair(tmp2734)

						var ifres2675 Obj

						if True == tmp2735 {
							tmp2725 := PrimTail(V780)

							tmp2726 := PrimHead(tmp2725)

							tmp2727 := PrimTail(tmp2726)

							tmp2728 := PrimHead(tmp2727)

							tmp2729 := PrimHead(tmp2728)

							tmp2730 := PrimEqual(symfail_1if, tmp2729)

							var ifres2677 Obj

							if True == tmp2730 {
								tmp2719 := PrimTail(V780)

								tmp2720 := PrimHead(tmp2719)

								tmp2721 := PrimTail(tmp2720)

								tmp2722 := PrimHead(tmp2721)

								tmp2723 := PrimTail(tmp2722)

								tmp2724 := PrimIsPair(tmp2723)

								var ifres2679 Obj

								if True == tmp2724 {
									tmp2712 := PrimTail(V780)

									tmp2713 := PrimHead(tmp2712)

									tmp2714 := PrimTail(tmp2713)

									tmp2715 := PrimHead(tmp2714)

									tmp2716 := PrimTail(tmp2715)

									tmp2717 := PrimTail(tmp2716)

									tmp2718 := PrimIsPair(tmp2717)

									var ifres2681 Obj

									if True == tmp2718 {
										tmp2704 := PrimTail(V780)

										tmp2705 := PrimHead(tmp2704)

										tmp2706 := PrimTail(tmp2705)

										tmp2707 := PrimHead(tmp2706)

										tmp2708 := PrimTail(tmp2707)

										tmp2709 := PrimTail(tmp2708)

										tmp2710 := PrimTail(tmp2709)

										tmp2711 := PrimEqual(Nil, tmp2710)

										var ifres2683 Obj

										if True == tmp2711 {
											tmp2699 := PrimTail(V780)

											tmp2700 := PrimHead(tmp2699)

											tmp2701 := PrimTail(tmp2700)

											tmp2702 := PrimTail(tmp2701)

											tmp2703 := PrimEqual(Nil, tmp2702)

											var ifres2685 Obj

											if True == tmp2703 {
												tmp2696 := PrimTail(V780)

												tmp2697 := PrimTail(tmp2696)

												tmp2698 := PrimEqual(Nil, tmp2697)

												var ifres2687 Obj

												if True == tmp2698 {
													tmp2689 := PrimTail(V780)

													tmp2690 := PrimHead(tmp2689)

													tmp2691 := PrimTail(tmp2690)

													tmp2692 := PrimHead(tmp2691)

													tmp2693 := PrimTail(tmp2692)

													tmp2694 := PrimHead(tmp2693)

													tmp2695 := PrimEqual(V777, tmp2694)

													var ifres2688 Obj

													if True == tmp2695 {
														ifres2688 = True

													} else {
														ifres2688 = False

													}

													ifres2687 = ifres2688

												} else {
													ifres2687 = False

												}

												var ifres2686 Obj

												if True == ifres2687 {
													ifres2686 = True

												} else {
													ifres2686 = False

												}

												ifres2685 = ifres2686

											} else {
												ifres2685 = False

											}

											var ifres2684 Obj

											if True == ifres2685 {
												ifres2684 = True

											} else {
												ifres2684 = False

											}

											ifres2683 = ifres2684

										} else {
											ifres2683 = False

										}

										var ifres2682 Obj

										if True == ifres2683 {
											ifres2682 = True

										} else {
											ifres2682 = False

										}

										ifres2681 = ifres2682

									} else {
										ifres2681 = False

									}

									var ifres2680 Obj

									if True == ifres2681 {
										ifres2680 = True

									} else {
										ifres2680 = False

									}

									ifres2679 = ifres2680

								} else {
									ifres2679 = False

								}

								var ifres2678 Obj

								if True == ifres2679 {
									ifres2678 = True

								} else {
									ifres2678 = False

								}

								ifres2677 = ifres2678

							} else {
								ifres2677 = False

							}

							var ifres2676 Obj

							if True == ifres2677 {
								ifres2676 = True

							} else {
								ifres2676 = False

							}

							ifres2675 = ifres2676

						} else {
							ifres2675 = False

						}

						var ifres2674 Obj

						if True == ifres2675 {
							ifres2674 = True

						} else {
							ifres2674 = False

						}

						ifres2673 = ifres2674

					} else {
						ifres2673 = False

					}

					var ifres2672 Obj

					if True == ifres2673 {
						ifres2672 = True

					} else {
						ifres2672 = False

					}

					ifres2671 = ifres2672

				} else {
					ifres2671 = False

				}

				var ifres2670 Obj

				if True == ifres2671 {
					ifres2670 = True

				} else {
					ifres2670 = False

				}

				ifres2669 = ifres2670

			} else {
				ifres2669 = False

			}

			var ifres2668 Obj

			if True == ifres2669 {
				ifres2668 = True

			} else {
				ifres2668 = False

			}

			ifres2667 = ifres2668

		} else {
			ifres2667 = False

		}

		if True == ifres2667 {
			tmp2554 := PrimTail(V780)

			tmp2555 := PrimHead(tmp2554)

			tmp2556 := PrimTail(tmp2555)

			tmp2557 := PrimHead(tmp2556)

			tmp2558 := PrimTail(tmp2557)

			tmp2559 := PrimHead(tmp2558)

			tmp2560 := Call(__e, PrimFunc(symshen_4scan_1body), tmp2559, V781)

			tmp2561 := PrimCons(symcond, tmp2560)

			tmp2562 := PrimCons(tmp2561, Nil)

			tmp2563 := PrimCons(symfreeze, tmp2562)

			tmp2564 := PrimHead(V780)

			tmp2565 := PrimTail(V780)

			tmp2566 := PrimHead(tmp2565)

			tmp2567 := PrimTail(tmp2566)

			tmp2568 := PrimHead(tmp2567)

			tmp2569 := PrimTail(tmp2568)

			tmp2570 := PrimTail(tmp2569)

			tmp2571 := PrimHead(tmp2570)

			tmp2572 := PrimTail(V780)

			tmp2573 := PrimHead(tmp2572)

			tmp2574 := PrimTail(tmp2573)

			tmp2575 := PrimHead(tmp2574)

			tmp2576 := PrimTail(tmp2575)

			tmp2577 := PrimHead(tmp2576)

			tmp2578 := PrimCons(V779, Nil)

			tmp2579 := PrimCons(tmp2577, tmp2578)

			tmp2580 := PrimCons(V778, Nil)

			tmp2581 := PrimCons(symthaw, tmp2580)

			tmp2582 := PrimCons(V779, Nil)

			tmp2583 := PrimCons(tmp2581, tmp2582)

			tmp2584 := PrimCons(tmp2579, tmp2583)

			tmp2585 := PrimCons(symif, tmp2584)

			tmp2586 := PrimCons(tmp2585, Nil)

			tmp2587 := PrimCons(tmp2571, tmp2586)

			tmp2588 := PrimCons(V779, tmp2587)

			tmp2589 := PrimCons(symlet, tmp2588)

			tmp2590 := PrimCons(V778, Nil)

			tmp2591 := PrimCons(symthaw, tmp2590)

			tmp2592 := PrimCons(tmp2591, Nil)

			tmp2593 := PrimCons(tmp2589, tmp2592)

			tmp2594 := PrimCons(tmp2564, tmp2593)

			tmp2595 := PrimCons(symif, tmp2594)

			tmp2596 := PrimCons(tmp2595, Nil)

			tmp2597 := PrimCons(tmp2563, tmp2596)

			tmp2598 := PrimCons(V778, tmp2597)

			tmp2599 := PrimCons(symlet, tmp2598)

			tmp2600 := PrimCons(tmp2599, Nil)

			tmp2601 := PrimCons(True, tmp2600)

			__e.Return(PrimCons(tmp2601, Nil))
			return

		} else {
			tmp2665 := PrimIsPair(V780)

			var ifres2638 Obj

			if True == tmp2665 {
				tmp2663 := PrimTail(V780)

				tmp2664 := PrimIsPair(tmp2663)

				var ifres2640 Obj

				if True == tmp2664 {
					tmp2660 := PrimTail(V780)

					tmp2661 := PrimHead(tmp2660)

					tmp2662 := PrimIsPair(tmp2661)

					var ifres2642 Obj

					if True == tmp2662 {
						tmp2656 := PrimTail(V780)

						tmp2657 := PrimHead(tmp2656)

						tmp2658 := PrimTail(tmp2657)

						tmp2659 := PrimIsPair(tmp2658)

						var ifres2644 Obj

						if True == tmp2659 {
							tmp2651 := PrimTail(V780)

							tmp2652 := PrimHead(tmp2651)

							tmp2653 := PrimTail(tmp2652)

							tmp2654 := PrimTail(tmp2653)

							tmp2655 := PrimEqual(Nil, tmp2654)

							var ifres2646 Obj

							if True == tmp2655 {
								tmp2648 := PrimTail(V780)

								tmp2649 := PrimTail(tmp2648)

								tmp2650 := PrimEqual(Nil, tmp2649)

								var ifres2647 Obj

								if True == tmp2650 {
									ifres2647 = True

								} else {
									ifres2647 = False

								}

								ifres2646 = ifres2647

							} else {
								ifres2646 = False

							}

							var ifres2645 Obj

							if True == ifres2646 {
								ifres2645 = True

							} else {
								ifres2645 = False

							}

							ifres2644 = ifres2645

						} else {
							ifres2644 = False

						}

						var ifres2643 Obj

						if True == ifres2644 {
							ifres2643 = True

						} else {
							ifres2643 = False

						}

						ifres2642 = ifres2643

					} else {
						ifres2642 = False

					}

					var ifres2641 Obj

					if True == ifres2642 {
						ifres2641 = True

					} else {
						ifres2641 = False

					}

					ifres2640 = ifres2641

				} else {
					ifres2640 = False

				}

				var ifres2639 Obj

				if True == ifres2640 {
					ifres2639 = True

				} else {
					ifres2639 = False

				}

				ifres2638 = ifres2639

			} else {
				ifres2638 = False

			}

			if True == ifres2638 {
				tmp2602 := Call(__e, PrimFunc(symshen_4scan_1body), V777, V781)

				tmp2603 := PrimCons(symcond, tmp2602)

				tmp2604 := PrimCons(tmp2603, Nil)

				tmp2605 := PrimCons(symfreeze, tmp2604)

				tmp2606 := PrimHead(V780)

				tmp2607 := PrimTail(V780)

				tmp2608 := PrimHead(tmp2607)

				tmp2609 := PrimTail(tmp2608)

				tmp2610 := PrimHead(tmp2609)

				tmp2611 := PrimCons(symfail, Nil)

				tmp2612 := PrimCons(tmp2611, Nil)

				tmp2613 := PrimCons(V779, tmp2612)

				tmp2614 := PrimCons(sym_a, tmp2613)

				tmp2615 := PrimCons(V778, Nil)

				tmp2616 := PrimCons(symthaw, tmp2615)

				tmp2617 := PrimCons(V779, Nil)

				tmp2618 := PrimCons(tmp2616, tmp2617)

				tmp2619 := PrimCons(tmp2614, tmp2618)

				tmp2620 := PrimCons(symif, tmp2619)

				tmp2621 := PrimCons(tmp2620, Nil)

				tmp2622 := PrimCons(tmp2610, tmp2621)

				tmp2623 := PrimCons(V779, tmp2622)

				tmp2624 := PrimCons(symlet, tmp2623)

				tmp2625 := PrimCons(V778, Nil)

				tmp2626 := PrimCons(symthaw, tmp2625)

				tmp2627 := PrimCons(tmp2626, Nil)

				tmp2628 := PrimCons(tmp2624, tmp2627)

				tmp2629 := PrimCons(tmp2606, tmp2628)

				tmp2630 := PrimCons(symif, tmp2629)

				tmp2631 := PrimCons(tmp2630, Nil)

				tmp2632 := PrimCons(tmp2605, tmp2631)

				tmp2633 := PrimCons(V778, tmp2632)

				tmp2634 := PrimCons(symlet, tmp2633)

				tmp2635 := PrimCons(tmp2634, Nil)

				tmp2636 := PrimCons(True, tmp2635)

				__e.Return(PrimCons(tmp2636, Nil))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.choicepoint")))
				return
			}

		}

	}, 5)

	tmp2746 := Call(__e, ns2_1set, symshen_4choicepoint, tmp2553)

	_ = tmp2746

	tmp2747 := MakeNative(func(__e *ControlFlow) {
		V783 := __e.Get(1)
		_ = V783
		V784 := __e.Get(2)
		_ = V784
		V785 := __e.Get(3)
		_ = V785
		tmp2761 := PrimEqual(V783, V785)

		if True == tmp2761 {
			__e.Return(V784)
			return
		} else {
			tmp2759 := PrimIsPair(V785)

			if True == tmp2759 {
				tmp2748 := MakeNative(func(__e *ControlFlow) {
					W786 := __e.Get(1)
					_ = W786
					tmp2754 := PrimHead(V785)

					tmp2755 := PrimEqual(W786, tmp2754)

					if True == tmp2755 {
						tmp2749 := PrimHead(V785)

						tmp2750 := PrimTail(V785)

						tmp2751 := Call(__e, PrimFunc(symshen_4rep_1X), V783, V784, tmp2750)

						__e.Return(PrimCons(tmp2749, tmp2751))
						return

					} else {
						tmp2752 := PrimTail(V785)

						__e.Return(PrimCons(W786, tmp2752))
						return

					}

				}, 1)

				tmp2756 := PrimHead(V785)

				tmp2757 := Call(__e, PrimFunc(symshen_4rep_1X), V783, V784, tmp2756)

				__e.TailApply(tmp2748, tmp2757)
				return

			} else {
				__e.Return(V785)
				return
			}

		}

	}, 3)

	tmp2762 := Call(__e, ns2_1set, symshen_4rep_1X, tmp2747)

	_ = tmp2762

	tmp2763 := MakeNative(func(__e *ControlFlow) {
		V787 := __e.Get(1)
		_ = V787
		tmp2846 := PrimIsPair(V787)

		var ifres2827 Obj

		if True == tmp2846 {
			tmp2844 := PrimHead(V787)

			tmp2845 := PrimEqual(symlambda, tmp2844)

			var ifres2829 Obj

			if True == tmp2845 {
				tmp2842 := PrimTail(V787)

				tmp2843 := PrimIsPair(tmp2842)

				var ifres2831 Obj

				if True == tmp2843 {
					tmp2839 := PrimTail(V787)

					tmp2840 := PrimTail(tmp2839)

					tmp2841 := PrimIsPair(tmp2840)

					var ifres2833 Obj

					if True == tmp2841 {
						tmp2835 := PrimTail(V787)

						tmp2836 := PrimTail(tmp2835)

						tmp2837 := PrimTail(tmp2836)

						tmp2838 := PrimEqual(Nil, tmp2837)

						var ifres2834 Obj

						if True == tmp2838 {
							ifres2834 = True

						} else {
							ifres2834 = False

						}

						ifres2833 = ifres2834

					} else {
						ifres2833 = False

					}

					var ifres2832 Obj

					if True == ifres2833 {
						ifres2832 = True

					} else {
						ifres2832 = False

					}

					ifres2831 = ifres2832

				} else {
					ifres2831 = False

				}

				var ifres2830 Obj

				if True == ifres2831 {
					ifres2830 = True

				} else {
					ifres2830 = False

				}

				ifres2829 = ifres2830

			} else {
				ifres2829 = False

			}

			var ifres2828 Obj

			if True == ifres2829 {
				ifres2828 = True

			} else {
				ifres2828 = False

			}

			ifres2827 = ifres2828

		} else {
			ifres2827 = False

		}

		if True == ifres2827 {
			tmp2764 := MakeNative(func(__e *ControlFlow) {
				W788 := __e.Get(1)
				_ = W788
				tmp2765 := MakeNative(func(__e *ControlFlow) {
					W789 := __e.Get(1)
					_ = W789
					tmp2766 := MakeNative(func(__e *ControlFlow) {
						Z790 := __e.Get(1)
						_ = Z790
						__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z790)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp2766, W789)
					return

				}, 1)

				tmp2767 := PrimTail(V787)

				tmp2768 := PrimHead(tmp2767)

				tmp2769 := PrimTail(V787)

				tmp2770 := PrimTail(tmp2769)

				tmp2771 := PrimHead(tmp2770)

				tmp2772 := Call(__e, PrimFunc(symshen_4beta), tmp2768, W788, tmp2771)

				tmp2773 := PrimCons(tmp2772, Nil)

				tmp2774 := PrimCons(W788, tmp2773)

				tmp2775 := PrimCons(symlambda, tmp2774)

				__e.TailApply(tmp2765, tmp2775)
				return

			}, 1)

			tmp2776 := Call(__e, PrimFunc(symgensym), symZ)

			__e.TailApply(tmp2764, tmp2776)
			return

		} else {
			tmp2825 := PrimIsPair(V787)

			var ifres2799 Obj

			if True == tmp2825 {
				tmp2823 := PrimHead(V787)

				tmp2824 := PrimEqual(symlet, tmp2823)

				var ifres2801 Obj

				if True == tmp2824 {
					tmp2821 := PrimTail(V787)

					tmp2822 := PrimIsPair(tmp2821)

					var ifres2803 Obj

					if True == tmp2822 {
						tmp2818 := PrimTail(V787)

						tmp2819 := PrimTail(tmp2818)

						tmp2820 := PrimIsPair(tmp2819)

						var ifres2805 Obj

						if True == tmp2820 {
							tmp2814 := PrimTail(V787)

							tmp2815 := PrimTail(tmp2814)

							tmp2816 := PrimTail(tmp2815)

							tmp2817 := PrimIsPair(tmp2816)

							var ifres2807 Obj

							if True == tmp2817 {
								tmp2809 := PrimTail(V787)

								tmp2810 := PrimTail(tmp2809)

								tmp2811 := PrimTail(tmp2810)

								tmp2812 := PrimTail(tmp2811)

								tmp2813 := PrimEqual(Nil, tmp2812)

								var ifres2808 Obj

								if True == tmp2813 {
									ifres2808 = True

								} else {
									ifres2808 = False

								}

								ifres2807 = ifres2808

							} else {
								ifres2807 = False

							}

							var ifres2806 Obj

							if True == ifres2807 {
								ifres2806 = True

							} else {
								ifres2806 = False

							}

							ifres2805 = ifres2806

						} else {
							ifres2805 = False

						}

						var ifres2804 Obj

						if True == ifres2805 {
							ifres2804 = True

						} else {
							ifres2804 = False

						}

						ifres2803 = ifres2804

					} else {
						ifres2803 = False

					}

					var ifres2802 Obj

					if True == ifres2803 {
						ifres2802 = True

					} else {
						ifres2802 = False

					}

					ifres2801 = ifres2802

				} else {
					ifres2801 = False

				}

				var ifres2800 Obj

				if True == ifres2801 {
					ifres2800 = True

				} else {
					ifres2800 = False

				}

				ifres2799 = ifres2800

			} else {
				ifres2799 = False

			}

			if True == ifres2799 {
				tmp2777 := MakeNative(func(__e *ControlFlow) {
					W791 := __e.Get(1)
					_ = W791
					tmp2778 := MakeNative(func(__e *ControlFlow) {
						W792 := __e.Get(1)
						_ = W792
						tmp2779 := MakeNative(func(__e *ControlFlow) {
							Z793 := __e.Get(1)
							_ = Z793
							__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z793)
							return
						}, 1)

						__e.TailApply(PrimFunc(symmap), tmp2779, W792)
						return

					}, 1)

					tmp2780 := PrimTail(V787)

					tmp2781 := PrimTail(tmp2780)

					tmp2782 := PrimHead(tmp2781)

					tmp2783 := PrimTail(V787)

					tmp2784 := PrimHead(tmp2783)

					tmp2785 := PrimTail(V787)

					tmp2786 := PrimTail(tmp2785)

					tmp2787 := PrimTail(tmp2786)

					tmp2788 := PrimHead(tmp2787)

					tmp2789 := Call(__e, PrimFunc(symshen_4beta), tmp2784, W791, tmp2788)

					tmp2790 := PrimCons(tmp2789, Nil)

					tmp2791 := PrimCons(tmp2782, tmp2790)

					tmp2792 := PrimCons(W791, tmp2791)

					tmp2793 := PrimCons(symlet, tmp2792)

					__e.TailApply(tmp2778, tmp2793)
					return

				}, 1)

				tmp2794 := Call(__e, PrimFunc(symgensym), symW)

				__e.TailApply(tmp2777, tmp2794)
				return

			} else {
				tmp2797 := PrimIsPair(V787)

				if True == tmp2797 {
					tmp2795 := MakeNative(func(__e *ControlFlow) {
						Z794 := __e.Get(1)
						_ = Z794
						__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z794)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp2795, V787)
					return

				} else {
					__e.Return(V787)
					return
				}

			}

		}

	}, 1)

	tmp2847 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp2763)

	_ = tmp2847

	tmp2848 := MakeNative(func(__e *ControlFlow) {
		V795 := __e.Get(1)
		_ = V795
		V796 := __e.Get(2)
		_ = V796
		tmp2849 := MakeNative(func(__e *ControlFlow) {
			Z797 := __e.Get(1)
			_ = Z797
			tmp2850 := Call(__e, PrimFunc(symfst), Z797)

			tmp2851 := Call(__e, PrimFunc(symsnd), Z797)

			tmp2852 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp2851)

			__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp2850, V796, tmp2852)
			return

		}, 1)

		__e.TailApply(PrimFunc(symmap), tmp2849, V795)
		return

	}, 2)

	tmp2853 := Call(__e, ns2_1set, symshen_4kl_1body, tmp2848)

	_ = tmp2853

	tmp2854 := MakeNative(func(__e *ControlFlow) {
		V806 := __e.Get(1)
		_ = V806
		V807 := __e.Get(2)
		_ = V807
		V808 := __e.Get(3)
		_ = V808
		V809 := __e.Get(4)
		_ = V809
		tmp2984 := PrimEqual(Nil, V807)

		var ifres2959 Obj

		if True == tmp2984 {
			tmp2983 := PrimEqual(Nil, V808)

			var ifres2961 Obj

			if True == tmp2983 {
				tmp2982 := PrimIsPair(V809)

				var ifres2963 Obj

				if True == tmp2982 {
					tmp2980 := PrimHead(V809)

					tmp2981 := PrimEqual(symwhere, tmp2980)

					var ifres2965 Obj

					if True == tmp2981 {
						tmp2978 := PrimTail(V809)

						tmp2979 := PrimIsPair(tmp2978)

						var ifres2967 Obj

						if True == tmp2979 {
							tmp2975 := PrimTail(V809)

							tmp2976 := PrimTail(tmp2975)

							tmp2977 := PrimIsPair(tmp2976)

							var ifres2969 Obj

							if True == tmp2977 {
								tmp2971 := PrimTail(V809)

								tmp2972 := PrimTail(tmp2971)

								tmp2973 := PrimTail(tmp2972)

								tmp2974 := PrimEqual(Nil, tmp2973)

								var ifres2970 Obj

								if True == tmp2974 {
									ifres2970 = True

								} else {
									ifres2970 = False

								}

								ifres2969 = ifres2970

							} else {
								ifres2969 = False

							}

							var ifres2968 Obj

							if True == ifres2969 {
								ifres2968 = True

							} else {
								ifres2968 = False

							}

							ifres2967 = ifres2968

						} else {
							ifres2967 = False

						}

						var ifres2966 Obj

						if True == ifres2967 {
							ifres2966 = True

						} else {
							ifres2966 = False

						}

						ifres2965 = ifres2966

					} else {
						ifres2965 = False

					}

					var ifres2964 Obj

					if True == ifres2965 {
						ifres2964 = True

					} else {
						ifres2964 = False

					}

					ifres2963 = ifres2964

				} else {
					ifres2963 = False

				}

				var ifres2962 Obj

				if True == ifres2963 {
					ifres2962 = True

				} else {
					ifres2962 = False

				}

				ifres2961 = ifres2962

			} else {
				ifres2961 = False

			}

			var ifres2960 Obj

			if True == ifres2961 {
				ifres2960 = True

			} else {
				ifres2960 = False

			}

			ifres2959 = ifres2960

		} else {
			ifres2959 = False

		}

		if True == ifres2959 {
			tmp2855 := PrimTail(V809)

			tmp2856 := PrimHead(tmp2855)

			tmp2857 := PrimCons(tmp2856, V806)

			tmp2858 := PrimTail(V809)

			tmp2859 := PrimTail(tmp2858)

			tmp2860 := PrimHead(tmp2859)

			__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp2857, Nil, Nil, tmp2860)
			return

		} else {
			tmp2957 := PrimEqual(Nil, V807)

			var ifres2954 Obj

			if True == tmp2957 {
				tmp2956 := PrimEqual(Nil, V808)

				var ifres2955 Obj

				if True == tmp2956 {
					ifres2955 = True

				} else {
					ifres2955 = False

				}

				ifres2954 = ifres2955

			} else {
				ifres2954 = False

			}

			if True == ifres2954 {
				tmp2861 := Call(__e, PrimFunc(symreverse), V806)

				tmp2862 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp2861)

				tmp2863 := PrimCons(V809, Nil)

				__e.Return(PrimCons(tmp2862, tmp2863))
				return

			} else {
				tmp2952 := PrimIsPair(V807)

				var ifres2945 Obj

				if True == tmp2952 {
					tmp2951 := PrimIsPair(V808)

					var ifres2947 Obj

					if True == tmp2951 {
						tmp2949 := PrimHead(V807)

						tmp2950 := PrimIsVariable(tmp2949)

						var ifres2948 Obj

						if True == tmp2950 {
							ifres2948 = True

						} else {
							ifres2948 = False

						}

						ifres2947 = ifres2948

					} else {
						ifres2947 = False

					}

					var ifres2946 Obj

					if True == ifres2947 {
						ifres2946 = True

					} else {
						ifres2946 = False

					}

					ifres2945 = ifres2946

				} else {
					ifres2945 = False

				}

				if True == ifres2945 {
					tmp2864 := PrimTail(V807)

					tmp2865 := PrimTail(V808)

					tmp2866 := PrimHead(V807)

					tmp2867 := PrimHead(V808)

					tmp2868 := Call(__e, PrimFunc(symshen_4beta), tmp2866, tmp2867, V809)

					__e.TailApply(PrimFunc(symshen_4triple_1stack), V806, tmp2864, tmp2865, tmp2868)
					return

				} else {
					tmp2943 := PrimIsPair(V807)

					var ifres2918 Obj

					if True == tmp2943 {
						tmp2941 := PrimHead(V807)

						tmp2942 := PrimIsPair(tmp2941)

						var ifres2920 Obj

						if True == tmp2942 {
							tmp2938 := PrimHead(V807)

							tmp2939 := PrimTail(tmp2938)

							tmp2940 := PrimIsPair(tmp2939)

							var ifres2922 Obj

							if True == tmp2940 {
								tmp2934 := PrimHead(V807)

								tmp2935 := PrimTail(tmp2934)

								tmp2936 := PrimTail(tmp2935)

								tmp2937 := PrimIsPair(tmp2936)

								var ifres2924 Obj

								if True == tmp2937 {
									tmp2929 := PrimHead(V807)

									tmp2930 := PrimTail(tmp2929)

									tmp2931 := PrimTail(tmp2930)

									tmp2932 := PrimTail(tmp2931)

									tmp2933 := PrimEqual(Nil, tmp2932)

									var ifres2926 Obj

									if True == tmp2933 {
										tmp2928 := PrimIsPair(V808)

										var ifres2927 Obj

										if True == tmp2928 {
											ifres2927 = True

										} else {
											ifres2927 = False

										}

										ifres2926 = ifres2927

									} else {
										ifres2926 = False

									}

									var ifres2925 Obj

									if True == ifres2926 {
										ifres2925 = True

									} else {
										ifres2925 = False

									}

									ifres2924 = ifres2925

								} else {
									ifres2924 = False

								}

								var ifres2923 Obj

								if True == ifres2924 {
									ifres2923 = True

								} else {
									ifres2923 = False

								}

								ifres2922 = ifres2923

							} else {
								ifres2922 = False

							}

							var ifres2921 Obj

							if True == ifres2922 {
								ifres2921 = True

							} else {
								ifres2921 = False

							}

							ifres2920 = ifres2921

						} else {
							ifres2920 = False

						}

						var ifres2919 Obj

						if True == ifres2920 {
							ifres2919 = True

						} else {
							ifres2919 = False

						}

						ifres2918 = ifres2919

					} else {
						ifres2918 = False

					}

					if True == ifres2918 {
						tmp2869 := PrimHead(V807)

						tmp2870 := PrimHead(tmp2869)

						tmp2871 := Call(__e, PrimFunc(symshen_4op_1test), tmp2870)

						tmp2872 := PrimHead(V808)

						tmp2873 := PrimCons(tmp2872, Nil)

						tmp2874 := PrimCons(tmp2871, tmp2873)

						tmp2875 := PrimCons(tmp2874, V806)

						tmp2876 := PrimHead(V807)

						tmp2877 := PrimTail(tmp2876)

						tmp2878 := PrimHead(tmp2877)

						tmp2879 := PrimHead(V807)

						tmp2880 := PrimTail(tmp2879)

						tmp2881 := PrimTail(tmp2880)

						tmp2882 := PrimHead(tmp2881)

						tmp2883 := PrimTail(V807)

						tmp2884 := PrimCons(tmp2882, tmp2883)

						tmp2885 := PrimCons(tmp2878, tmp2884)

						tmp2886 := PrimHead(V807)

						tmp2887 := PrimHead(tmp2886)

						tmp2888 := Call(__e, PrimFunc(symshen_4op1), tmp2887)

						tmp2889 := PrimHead(V808)

						tmp2890 := PrimCons(tmp2889, Nil)

						tmp2891 := PrimCons(tmp2888, tmp2890)

						tmp2892 := PrimHead(V807)

						tmp2893 := PrimHead(tmp2892)

						tmp2894 := Call(__e, PrimFunc(symshen_4op2), tmp2893)

						tmp2895 := PrimHead(V808)

						tmp2896 := PrimCons(tmp2895, Nil)

						tmp2897 := PrimCons(tmp2894, tmp2896)

						tmp2898 := PrimTail(V808)

						tmp2899 := PrimCons(tmp2897, tmp2898)

						tmp2900 := PrimCons(tmp2891, tmp2899)

						tmp2901 := PrimHead(V807)

						tmp2902 := PrimHead(V808)

						tmp2903 := Call(__e, PrimFunc(symshen_4beta), tmp2901, tmp2902, V809)

						__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp2875, tmp2885, tmp2900, tmp2903)
						return

					} else {
						tmp2916 := PrimIsPair(V807)

						var ifres2913 Obj

						if True == tmp2916 {
							tmp2915 := PrimIsPair(V808)

							var ifres2914 Obj

							if True == tmp2915 {
								ifres2914 = True

							} else {
								ifres2914 = False

							}

							ifres2913 = ifres2914

						} else {
							ifres2913 = False

						}

						if True == ifres2913 {
							tmp2904 := PrimHead(V807)

							tmp2905 := PrimHead(V808)

							tmp2906 := PrimCons(tmp2905, Nil)

							tmp2907 := PrimCons(tmp2904, tmp2906)

							tmp2908 := PrimCons(sym_a, tmp2907)

							tmp2909 := PrimCons(tmp2908, V806)

							tmp2910 := PrimTail(V807)

							tmp2911 := PrimTail(V808)

							__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp2909, tmp2910, tmp2911, V809)
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

	tmp2985 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp2854)

	_ = tmp2985

	tmp2986 := MakeNative(func(__e *ControlFlow) {
		V812 := __e.Get(1)
		_ = V812
		tmp3005 := PrimEqual(Nil, V812)

		if True == tmp3005 {
			__e.Return(True)
			return
		} else {
			tmp3003 := PrimIsPair(V812)

			var ifres2999 Obj

			if True == tmp3003 {
				tmp3001 := PrimTail(V812)

				tmp3002 := PrimEqual(Nil, tmp3001)

				var ifres3000 Obj

				if True == tmp3002 {
					ifres3000 = True

				} else {
					ifres3000 = False

				}

				ifres2999 = ifres3000

			} else {
				ifres2999 = False

			}

			if True == ifres2999 {
				__e.Return(PrimHead(V812))
				return
			} else {
				tmp2997 := PrimIsPair(V812)

				var ifres2993 Obj

				if True == tmp2997 {
					tmp2995 := PrimTail(V812)

					tmp2996 := PrimIsPair(tmp2995)

					var ifres2994 Obj

					if True == tmp2996 {
						ifres2994 = True

					} else {
						ifres2994 = False

					}

					ifres2993 = ifres2994

				} else {
					ifres2993 = False

				}

				if True == ifres2993 {
					tmp2987 := PrimHead(V812)

					tmp2988 := PrimTail(V812)

					tmp2989 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp2988)

					tmp2990 := PrimCons(tmp2989, Nil)

					tmp2991 := PrimCons(tmp2987, tmp2990)

					__e.Return(PrimCons(symand, tmp2991))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.rectify-test")))
					return
				}

			}

		}

	}, 1)

	tmp3006 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp2986)

	_ = tmp3006

	tmp3007 := MakeNative(func(__e *ControlFlow) {
		V822 := __e.Get(1)
		_ = V822
		V823 := __e.Get(2)
		_ = V823
		V824 := __e.Get(3)
		_ = V824
		tmp3084 := PrimEqual(V822, V824)

		if True == tmp3084 {
			__e.Return(V823)
			return
		} else {
			tmp3082 := PrimIsPair(V824)

			var ifres3058 Obj

			if True == tmp3082 {
				tmp3080 := PrimHead(V824)

				tmp3081 := PrimEqual(symlambda, tmp3080)

				var ifres3060 Obj

				if True == tmp3081 {
					tmp3078 := PrimTail(V824)

					tmp3079 := PrimIsPair(tmp3078)

					var ifres3062 Obj

					if True == tmp3079 {
						tmp3075 := PrimTail(V824)

						tmp3076 := PrimTail(tmp3075)

						tmp3077 := PrimIsPair(tmp3076)

						var ifres3064 Obj

						if True == tmp3077 {
							tmp3071 := PrimTail(V824)

							tmp3072 := PrimTail(tmp3071)

							tmp3073 := PrimTail(tmp3072)

							tmp3074 := PrimEqual(Nil, tmp3073)

							var ifres3066 Obj

							if True == tmp3074 {
								tmp3068 := PrimTail(V824)

								tmp3069 := PrimHead(tmp3068)

								tmp3070 := PrimEqual(V822, tmp3069)

								var ifres3067 Obj

								if True == tmp3070 {
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

				var ifres3059 Obj

				if True == ifres3060 {
					ifres3059 = True

				} else {
					ifres3059 = False

				}

				ifres3058 = ifres3059

			} else {
				ifres3058 = False

			}

			if True == ifres3058 {
				__e.Return(V824)
				return
			} else {
				tmp3056 := PrimIsPair(V824)

				var ifres3025 Obj

				if True == tmp3056 {
					tmp3054 := PrimHead(V824)

					tmp3055 := PrimEqual(symlet, tmp3054)

					var ifres3027 Obj

					if True == tmp3055 {
						tmp3052 := PrimTail(V824)

						tmp3053 := PrimIsPair(tmp3052)

						var ifres3029 Obj

						if True == tmp3053 {
							tmp3049 := PrimTail(V824)

							tmp3050 := PrimTail(tmp3049)

							tmp3051 := PrimIsPair(tmp3050)

							var ifres3031 Obj

							if True == tmp3051 {
								tmp3045 := PrimTail(V824)

								tmp3046 := PrimTail(tmp3045)

								tmp3047 := PrimTail(tmp3046)

								tmp3048 := PrimIsPair(tmp3047)

								var ifres3033 Obj

								if True == tmp3048 {
									tmp3040 := PrimTail(V824)

									tmp3041 := PrimTail(tmp3040)

									tmp3042 := PrimTail(tmp3041)

									tmp3043 := PrimTail(tmp3042)

									tmp3044 := PrimEqual(Nil, tmp3043)

									var ifres3035 Obj

									if True == tmp3044 {
										tmp3037 := PrimTail(V824)

										tmp3038 := PrimHead(tmp3037)

										tmp3039 := PrimEqual(V822, tmp3038)

										var ifres3036 Obj

										if True == tmp3039 {
											ifres3036 = True

										} else {
											ifres3036 = False

										}

										ifres3035 = ifres3036

									} else {
										ifres3035 = False

									}

									var ifres3034 Obj

									if True == ifres3035 {
										ifres3034 = True

									} else {
										ifres3034 = False

									}

									ifres3033 = ifres3034

								} else {
									ifres3033 = False

								}

								var ifres3032 Obj

								if True == ifres3033 {
									ifres3032 = True

								} else {
									ifres3032 = False

								}

								ifres3031 = ifres3032

							} else {
								ifres3031 = False

							}

							var ifres3030 Obj

							if True == ifres3031 {
								ifres3030 = True

							} else {
								ifres3030 = False

							}

							ifres3029 = ifres3030

						} else {
							ifres3029 = False

						}

						var ifres3028 Obj

						if True == ifres3029 {
							ifres3028 = True

						} else {
							ifres3028 = False

						}

						ifres3027 = ifres3028

					} else {
						ifres3027 = False

					}

					var ifres3026 Obj

					if True == ifres3027 {
						ifres3026 = True

					} else {
						ifres3026 = False

					}

					ifres3025 = ifres3026

				} else {
					ifres3025 = False

				}

				if True == ifres3025 {
					tmp3008 := PrimTail(V824)

					tmp3009 := PrimHead(tmp3008)

					tmp3010 := PrimTail(V824)

					tmp3011 := PrimHead(tmp3010)

					tmp3012 := PrimTail(V824)

					tmp3013 := PrimTail(tmp3012)

					tmp3014 := PrimHead(tmp3013)

					tmp3015 := Call(__e, PrimFunc(symshen_4beta), tmp3011, V823, tmp3014)

					tmp3016 := PrimTail(V824)

					tmp3017 := PrimTail(tmp3016)

					tmp3018 := PrimTail(tmp3017)

					tmp3019 := PrimCons(tmp3015, tmp3018)

					tmp3020 := PrimCons(tmp3009, tmp3019)

					__e.Return(PrimCons(symlet, tmp3020))
					return

				} else {
					tmp3023 := PrimIsPair(V824)

					if True == tmp3023 {
						tmp3021 := MakeNative(func(__e *ControlFlow) {
							Z825 := __e.Get(1)
							_ = Z825
							__e.TailApply(PrimFunc(symshen_4beta), V822, V823, Z825)
							return
						}, 1)

						__e.TailApply(PrimFunc(symmap), tmp3021, V824)
						return

					} else {
						__e.Return(V824)
						return
					}

				}

			}

		}

	}, 3)

	tmp3085 := Call(__e, ns2_1set, symshen_4beta, tmp3007)

	_ = tmp3085

	tmp3086 := MakeNative(func(__e *ControlFlow) {
		V828 := __e.Get(1)
		_ = V828
		tmp3094 := PrimEqual(symcons, V828)

		if True == tmp3094 {
			__e.Return(symhd)
			return
		} else {
			tmp3092 := PrimEqual(sym_8s, V828)

			if True == tmp3092 {
				__e.Return(symhdstr)
				return
			} else {
				tmp3090 := PrimEqual(sym_8p, V828)

				if True == tmp3090 {
					__e.Return(symfst)
					return
				} else {
					tmp3088 := PrimEqual(sym_8v, V828)

					if True == tmp3088 {
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

	tmp3095 := Call(__e, ns2_1set, symshen_4op1, tmp3086)

	_ = tmp3095

	tmp3096 := MakeNative(func(__e *ControlFlow) {
		V831 := __e.Get(1)
		_ = V831
		tmp3104 := PrimEqual(symcons, V831)

		if True == tmp3104 {
			__e.Return(symtl)
			return
		} else {
			tmp3102 := PrimEqual(sym_8s, V831)

			if True == tmp3102 {
				__e.Return(symtlstr)
				return
			} else {
				tmp3100 := PrimEqual(sym_8p, V831)

				if True == tmp3100 {
					__e.Return(symsnd)
					return
				} else {
					tmp3098 := PrimEqual(sym_8v, V831)

					if True == tmp3098 {
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

	tmp3105 := Call(__e, ns2_1set, symshen_4op2, tmp3096)

	_ = tmp3105

	tmp3106 := MakeNative(func(__e *ControlFlow) {
		V834 := __e.Get(1)
		_ = V834
		tmp3114 := PrimEqual(symcons, V834)

		if True == tmp3114 {
			__e.Return(symcons_2)
			return
		} else {
			tmp3112 := PrimEqual(sym_8s, V834)

			if True == tmp3112 {
				__e.Return(symshen_4_7string_2)
				return
			} else {
				tmp3110 := PrimEqual(sym_8p, V834)

				if True == tmp3110 {
					__e.Return(symtuple_2)
					return
				} else {
					tmp3108 := PrimEqual(sym_8v, V834)

					if True == tmp3108 {
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

	tmp3115 := Call(__e, ns2_1set, symshen_4op_1test, tmp3106)

	_ = tmp3115

	tmp3116 := MakeNative(func(__e *ControlFlow) {
		V835 := __e.Get(1)
		_ = V835
		tmp3118 := PrimEqual(MakeString(""), V835)

		if True == tmp3118 {
			__e.Return(False)
			return
		} else {
			__e.Return(PrimIsString(V835))
			return
		}

	}, 1)

	tmp3119 := Call(__e, ns2_1set, symshen_4_7string_2, tmp3116)

	_ = tmp3119

	tmp3120 := MakeNative(func(__e *ControlFlow) {
		V836 := __e.Get(1)
		_ = V836
		tmp3122 := Call(__e, PrimFunc(symvector), MakeNumber(0))

		tmp3123 := PrimEqual(V836, tmp3122)

		if True == tmp3123 {
			__e.Return(False)
			return
		} else {
			__e.TailApply(PrimFunc(symvector_2), V836)
			return
		}

	}, 1)

	tmp3124 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp3120)

	_ = tmp3124

	tmp3125 := MakeNative(func(__e *ControlFlow) {
		V839 := __e.Get(1)
		_ = V839
		tmp3129 := PrimEqual(sym_7, V839)

		if True == tmp3129 {
			__e.Return(PrimSet(symshen_4_dfactorise_2_d, True))
			return
		} else {
			tmp3127 := PrimEqual(sym_1, V839)

			if True == tmp3127 {
				__e.Return(PrimSet(symshen_4_dfactorise_2_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("factorise expects a + or a -\n")))
				return
			}

		}

	}, 1)

	tmp3130 := Call(__e, ns2_1set, symfactorise, tmp3125)

	_ = tmp3130

	tmp3131 := MakeNative(func(__e *ControlFlow) {
		V840 := __e.Get(1)
		_ = V840
		tmp3133 := PrimValue(symshen_4_dfactorise_2_d)

		if True == tmp3133 {
			__e.TailApply(PrimFunc(symshen_4factor), V840)
			return
		} else {
			__e.Return(V840)
			return
		}

	}, 1)

	tmp3134 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp3131)

	_ = tmp3134

	tmp3135 := MakeNative(func(__e *ControlFlow) {
		V841 := __e.Get(1)
		_ = V841
		tmp3192 := PrimIsPair(V841)

		var ifres3151 Obj

		if True == tmp3192 {
			tmp3190 := PrimHead(V841)

			tmp3191 := PrimEqual(symdefun, tmp3190)

			var ifres3153 Obj

			if True == tmp3191 {
				tmp3188 := PrimTail(V841)

				tmp3189 := PrimIsPair(tmp3188)

				var ifres3155 Obj

				if True == tmp3189 {
					tmp3185 := PrimTail(V841)

					tmp3186 := PrimTail(tmp3185)

					tmp3187 := PrimIsPair(tmp3186)

					var ifres3157 Obj

					if True == tmp3187 {
						tmp3181 := PrimTail(V841)

						tmp3182 := PrimTail(tmp3181)

						tmp3183 := PrimTail(tmp3182)

						tmp3184 := PrimIsPair(tmp3183)

						var ifres3159 Obj

						if True == tmp3184 {
							tmp3176 := PrimTail(V841)

							tmp3177 := PrimTail(tmp3176)

							tmp3178 := PrimTail(tmp3177)

							tmp3179 := PrimHead(tmp3178)

							tmp3180 := PrimIsPair(tmp3179)

							var ifres3161 Obj

							if True == tmp3180 {
								tmp3170 := PrimTail(V841)

								tmp3171 := PrimTail(tmp3170)

								tmp3172 := PrimTail(tmp3171)

								tmp3173 := PrimHead(tmp3172)

								tmp3174 := PrimHead(tmp3173)

								tmp3175 := PrimEqual(symcond, tmp3174)

								var ifres3163 Obj

								if True == tmp3175 {
									tmp3165 := PrimTail(V841)

									tmp3166 := PrimTail(tmp3165)

									tmp3167 := PrimTail(tmp3166)

									tmp3168 := PrimTail(tmp3167)

									tmp3169 := PrimEqual(Nil, tmp3168)

									var ifres3164 Obj

									if True == tmp3169 {
										ifres3164 = True

									} else {
										ifres3164 = False

									}

									ifres3163 = ifres3164

								} else {
									ifres3163 = False

								}

								var ifres3162 Obj

								if True == ifres3163 {
									ifres3162 = True

								} else {
									ifres3162 = False

								}

								ifres3161 = ifres3162

							} else {
								ifres3161 = False

							}

							var ifres3160 Obj

							if True == ifres3161 {
								ifres3160 = True

							} else {
								ifres3160 = False

							}

							ifres3159 = ifres3160

						} else {
							ifres3159 = False

						}

						var ifres3158 Obj

						if True == ifres3159 {
							ifres3158 = True

						} else {
							ifres3158 = False

						}

						ifres3157 = ifres3158

					} else {
						ifres3157 = False

					}

					var ifres3156 Obj

					if True == ifres3157 {
						ifres3156 = True

					} else {
						ifres3156 = False

					}

					ifres3155 = ifres3156

				} else {
					ifres3155 = False

				}

				var ifres3154 Obj

				if True == ifres3155 {
					ifres3154 = True

				} else {
					ifres3154 = False

				}

				ifres3153 = ifres3154

			} else {
				ifres3153 = False

			}

			var ifres3152 Obj

			if True == ifres3153 {
				ifres3152 = True

			} else {
				ifres3152 = False

			}

			ifres3151 = ifres3152

		} else {
			ifres3151 = False

		}

		if True == ifres3151 {
			tmp3136 := PrimTail(V841)

			tmp3137 := PrimHead(tmp3136)

			tmp3138 := PrimTail(V841)

			tmp3139 := PrimTail(tmp3138)

			tmp3140 := PrimHead(tmp3139)

			tmp3141 := PrimTail(V841)

			tmp3142 := PrimTail(tmp3141)

			tmp3143 := PrimTail(tmp3142)

			tmp3144 := PrimHead(tmp3143)

			tmp3145 := PrimTail(tmp3144)

			tmp3146 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp3145)

			tmp3147 := PrimCons(tmp3146, Nil)

			tmp3148 := PrimCons(tmp3140, tmp3147)

			tmp3149 := PrimCons(tmp3137, tmp3148)

			__e.Return(PrimCons(symdefun, tmp3149))
			return

		} else {
			__e.Return(V841)
			return
		}

	}, 1)

	tmp3193 := Call(__e, ns2_1set, symshen_4factor, tmp3135)

	_ = tmp3193

	tmp3194 := MakeNative(func(__e *ControlFlow) {
		V844 := __e.Get(1)
		_ = V844
		tmp3350 := PrimIsPair(V844)

		var ifres3330 Obj

		if True == tmp3350 {
			tmp3348 := PrimHead(V844)

			tmp3349 := PrimIsPair(tmp3348)

			var ifres3332 Obj

			if True == tmp3349 {
				tmp3345 := PrimHead(V844)

				tmp3346 := PrimHead(tmp3345)

				tmp3347 := PrimEqual(True, tmp3346)

				var ifres3334 Obj

				if True == tmp3347 {
					tmp3342 := PrimHead(V844)

					tmp3343 := PrimTail(tmp3342)

					tmp3344 := PrimIsPair(tmp3343)

					var ifres3336 Obj

					if True == tmp3344 {
						tmp3338 := PrimHead(V844)

						tmp3339 := PrimTail(tmp3338)

						tmp3340 := PrimTail(tmp3339)

						tmp3341 := PrimEqual(Nil, tmp3340)

						var ifres3337 Obj

						if True == tmp3341 {
							ifres3337 = True

						} else {
							ifres3337 = False

						}

						ifres3336 = ifres3337

					} else {
						ifres3336 = False

					}

					var ifres3335 Obj

					if True == ifres3336 {
						ifres3335 = True

					} else {
						ifres3335 = False

					}

					ifres3334 = ifres3335

				} else {
					ifres3334 = False

				}

				var ifres3333 Obj

				if True == ifres3334 {
					ifres3333 = True

				} else {
					ifres3333 = False

				}

				ifres3332 = ifres3333

			} else {
				ifres3332 = False

			}

			var ifres3331 Obj

			if True == ifres3332 {
				ifres3331 = True

			} else {
				ifres3331 = False

			}

			ifres3330 = ifres3331

		} else {
			ifres3330 = False

		}

		if True == ifres3330 {
			tmp3195 := PrimHead(V844)

			tmp3196 := PrimTail(tmp3195)

			__e.Return(PrimHead(tmp3196))
			return

		} else {
			tmp3328 := PrimIsPair(V844)

			var ifres3281 Obj

			if True == tmp3328 {
				tmp3326 := PrimHead(V844)

				tmp3327 := PrimIsPair(tmp3326)

				var ifres3283 Obj

				if True == tmp3327 {
					tmp3323 := PrimHead(V844)

					tmp3324 := PrimHead(tmp3323)

					tmp3325 := PrimIsPair(tmp3324)

					var ifres3285 Obj

					if True == tmp3325 {
						tmp3319 := PrimHead(V844)

						tmp3320 := PrimHead(tmp3319)

						tmp3321 := PrimHead(tmp3320)

						tmp3322 := PrimEqual(symand, tmp3321)

						var ifres3287 Obj

						if True == tmp3322 {
							tmp3315 := PrimHead(V844)

							tmp3316 := PrimHead(tmp3315)

							tmp3317 := PrimTail(tmp3316)

							tmp3318 := PrimIsPair(tmp3317)

							var ifres3289 Obj

							if True == tmp3318 {
								tmp3310 := PrimHead(V844)

								tmp3311 := PrimHead(tmp3310)

								tmp3312 := PrimTail(tmp3311)

								tmp3313 := PrimTail(tmp3312)

								tmp3314 := PrimIsPair(tmp3313)

								var ifres3291 Obj

								if True == tmp3314 {
									tmp3304 := PrimHead(V844)

									tmp3305 := PrimHead(tmp3304)

									tmp3306 := PrimTail(tmp3305)

									tmp3307 := PrimTail(tmp3306)

									tmp3308 := PrimTail(tmp3307)

									tmp3309 := PrimEqual(Nil, tmp3308)

									var ifres3293 Obj

									if True == tmp3309 {
										tmp3301 := PrimHead(V844)

										tmp3302 := PrimTail(tmp3301)

										tmp3303 := PrimIsPair(tmp3302)

										var ifres3295 Obj

										if True == tmp3303 {
											tmp3297 := PrimHead(V844)

											tmp3298 := PrimTail(tmp3297)

											tmp3299 := PrimTail(tmp3298)

											tmp3300 := PrimEqual(Nil, tmp3299)

											var ifres3296 Obj

											if True == tmp3300 {
												ifres3296 = True

											} else {
												ifres3296 = False

											}

											ifres3295 = ifres3296

										} else {
											ifres3295 = False

										}

										var ifres3294 Obj

										if True == ifres3295 {
											ifres3294 = True

										} else {
											ifres3294 = False

										}

										ifres3293 = ifres3294

									} else {
										ifres3293 = False

									}

									var ifres3292 Obj

									if True == ifres3293 {
										ifres3292 = True

									} else {
										ifres3292 = False

									}

									ifres3291 = ifres3292

								} else {
									ifres3291 = False

								}

								var ifres3290 Obj

								if True == ifres3291 {
									ifres3290 = True

								} else {
									ifres3290 = False

								}

								ifres3289 = ifres3290

							} else {
								ifres3289 = False

							}

							var ifres3288 Obj

							if True == ifres3289 {
								ifres3288 = True

							} else {
								ifres3288 = False

							}

							ifres3287 = ifres3288

						} else {
							ifres3287 = False

						}

						var ifres3286 Obj

						if True == ifres3287 {
							ifres3286 = True

						} else {
							ifres3286 = False

						}

						ifres3285 = ifres3286

					} else {
						ifres3285 = False

					}

					var ifres3284 Obj

					if True == ifres3285 {
						ifres3284 = True

					} else {
						ifres3284 = False

					}

					ifres3283 = ifres3284

				} else {
					ifres3283 = False

				}

				var ifres3282 Obj

				if True == ifres3283 {
					ifres3282 = True

				} else {
					ifres3282 = False

				}

				ifres3281 = ifres3282

			} else {
				ifres3281 = False

			}

			if True == ifres3281 {
				tmp3197 := MakeNative(func(__e *ControlFlow) {
					W845 := __e.Get(1)
					_ = W845
					tmp3198 := MakeNative(func(__e *ControlFlow) {
						W846 := __e.Get(1)
						_ = W846
						tmp3246 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), W846)

						if True == tmp3246 {
							tmp3199 := PrimHead(V844)

							tmp3200 := PrimHead(tmp3199)

							tmp3201 := PrimHead(V844)

							tmp3202 := PrimTail(tmp3201)

							tmp3203 := PrimHead(tmp3202)

							tmp3204 := PrimTail(V844)

							tmp3205 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp3204)

							tmp3206 := PrimCons(tmp3205, Nil)

							tmp3207 := PrimCons(tmp3203, tmp3206)

							tmp3208 := PrimCons(tmp3200, tmp3207)

							__e.Return(PrimCons(symif, tmp3208))
							return

						} else {
							tmp3209 := MakeNative(func(__e *ControlFlow) {
								W847 := __e.Get(1)
								_ = W847
								tmp3210 := MakeNative(func(__e *ControlFlow) {
									W848 := __e.Get(1)
									_ = W848
									tmp3211 := MakeNative(func(__e *ControlFlow) {
										W849 := __e.Get(1)
										_ = W849
										tmp3212 := MakeNative(func(__e *ControlFlow) {
											W850 := __e.Get(1)
											_ = W850
											tmp3213 := MakeNative(func(__e *ControlFlow) {
												W851 := __e.Get(1)
												_ = W851
												__e.TailApply(PrimFunc(symshen_4remove_1indirection), W851)
												return
											}, 1)

											tmp3214 := PrimCons(W848, Nil)

											tmp3215 := PrimCons(symfreeze, tmp3214)

											tmp3216 := PrimHead(V844)

											tmp3217 := PrimHead(tmp3216)

											tmp3218 := PrimTail(tmp3217)

											tmp3219 := PrimHead(tmp3218)

											tmp3220 := PrimHead(V844)

											tmp3221 := PrimHead(tmp3220)

											tmp3222 := PrimTail(tmp3221)

											tmp3223 := PrimHead(tmp3222)

											tmp3224 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W850)

											tmp3225 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp3223, tmp3224)

											tmp3226 := PrimCons(W849, Nil)

											tmp3227 := PrimCons(symthaw, tmp3226)

											tmp3228 := PrimCons(tmp3227, Nil)

											tmp3229 := PrimCons(tmp3225, tmp3228)

											tmp3230 := PrimCons(tmp3219, tmp3229)

											tmp3231 := PrimCons(symif, tmp3230)

											tmp3232 := PrimCons(tmp3231, Nil)

											tmp3233 := PrimCons(tmp3215, tmp3232)

											tmp3234 := PrimCons(W849, tmp3233)

											tmp3235 := PrimCons(symlet, tmp3234)

											__e.TailApply(tmp3213, tmp3235)
											return

										}, 1)

										tmp3236 := PrimCons(W849, Nil)

										tmp3237 := PrimCons(symthaw, tmp3236)

										tmp3238 := PrimCons(tmp3237, Nil)

										tmp3239 := PrimCons(True, tmp3238)

										tmp3240 := PrimCons(tmp3239, W846)

										tmp3241 := Call(__e, PrimFunc(symreverse), tmp3240)

										__e.TailApply(tmp3212, tmp3241)
										return

									}, 1)

									tmp3242 := Call(__e, PrimFunc(symgensym), symGoTo)

									__e.TailApply(tmp3211, tmp3242)
									return

								}, 1)

								tmp3243 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W847)

								__e.TailApply(tmp3210, tmp3243)
								return

							}, 1)

							tmp3244 := Call(__e, PrimFunc(symsnd), W845)

							__e.TailApply(tmp3209, tmp3244)
							return

						}

					}, 1)

					tmp3247 := Call(__e, PrimFunc(symfst), W845)

					__e.TailApply(tmp3198, tmp3247)
					return

				}, 1)

				tmp3248 := PrimHead(V844)

				tmp3249 := PrimHead(tmp3248)

				tmp3250 := PrimTail(tmp3249)

				tmp3251 := PrimHead(tmp3250)

				tmp3252 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp3251, V844, Nil)

				__e.TailApply(tmp3197, tmp3252)
				return

			} else {
				tmp3279 := PrimIsPair(V844)

				var ifres3264 Obj

				if True == tmp3279 {
					tmp3277 := PrimHead(V844)

					tmp3278 := PrimIsPair(tmp3277)

					var ifres3266 Obj

					if True == tmp3278 {
						tmp3274 := PrimHead(V844)

						tmp3275 := PrimTail(tmp3274)

						tmp3276 := PrimIsPair(tmp3275)

						var ifres3268 Obj

						if True == tmp3276 {
							tmp3270 := PrimHead(V844)

							tmp3271 := PrimTail(tmp3270)

							tmp3272 := PrimTail(tmp3271)

							tmp3273 := PrimEqual(Nil, tmp3272)

							var ifres3269 Obj

							if True == tmp3273 {
								ifres3269 = True

							} else {
								ifres3269 = False

							}

							ifres3268 = ifres3269

						} else {
							ifres3268 = False

						}

						var ifres3267 Obj

						if True == ifres3268 {
							ifres3267 = True

						} else {
							ifres3267 = False

						}

						ifres3266 = ifres3267

					} else {
						ifres3266 = False

					}

					var ifres3265 Obj

					if True == ifres3266 {
						ifres3265 = True

					} else {
						ifres3265 = False

					}

					ifres3264 = ifres3265

				} else {
					ifres3264 = False

				}

				if True == ifres3264 {
					tmp3253 := PrimHead(V844)

					tmp3254 := PrimHead(tmp3253)

					tmp3255 := PrimHead(V844)

					tmp3256 := PrimTail(tmp3255)

					tmp3257 := PrimHead(tmp3256)

					tmp3258 := PrimTail(V844)

					tmp3259 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp3258)

					tmp3260 := PrimCons(tmp3259, Nil)

					tmp3261 := PrimCons(tmp3257, tmp3260)

					tmp3262 := PrimCons(tmp3254, tmp3261)

					__e.Return(PrimCons(symif, tmp3262))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1recognisors)
					return
				}

			}

		}

	}, 1)

	tmp3351 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp3194)

	_ = tmp3351

	tmp3352 := MakeNative(func(__e *ControlFlow) {
		V856 := __e.Get(1)
		_ = V856
		tmp3358 := PrimIsPair(V856)

		var ifres3354 Obj

		if True == tmp3358 {
			tmp3356 := PrimTail(V856)

			tmp3357 := PrimEqual(Nil, tmp3356)

			var ifres3355 Obj

			if True == tmp3357 {
				ifres3355 = True

			} else {
				ifres3355 = False

			}

			ifres3354 = ifres3355

		} else {
			ifres3354 = False

		}

		if True == ifres3354 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp3359 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp3352)

	_ = tmp3359

	tmp3360 := MakeNative(func(__e *ControlFlow) {
		V857 := __e.Get(1)
		_ = V857
		tmp3475 := PrimIsPair(V857)

		var ifres3375 Obj

		if True == tmp3475 {
			tmp3473 := PrimHead(V857)

			tmp3474 := PrimEqual(symlet, tmp3473)

			var ifres3377 Obj

			if True == tmp3474 {
				tmp3471 := PrimTail(V857)

				tmp3472 := PrimIsPair(tmp3471)

				var ifres3379 Obj

				if True == tmp3472 {
					tmp3468 := PrimTail(V857)

					tmp3469 := PrimTail(tmp3468)

					tmp3470 := PrimIsPair(tmp3469)

					var ifres3381 Obj

					if True == tmp3470 {
						tmp3464 := PrimTail(V857)

						tmp3465 := PrimTail(tmp3464)

						tmp3466 := PrimHead(tmp3465)

						tmp3467 := PrimIsPair(tmp3466)

						var ifres3383 Obj

						if True == tmp3467 {
							tmp3459 := PrimTail(V857)

							tmp3460 := PrimTail(tmp3459)

							tmp3461 := PrimHead(tmp3460)

							tmp3462 := PrimHead(tmp3461)

							tmp3463 := PrimEqual(symfreeze, tmp3462)

							var ifres3385 Obj

							if True == tmp3463 {
								tmp3454 := PrimTail(V857)

								tmp3455 := PrimTail(tmp3454)

								tmp3456 := PrimHead(tmp3455)

								tmp3457 := PrimTail(tmp3456)

								tmp3458 := PrimIsPair(tmp3457)

								var ifres3387 Obj

								if True == tmp3458 {
									tmp3448 := PrimTail(V857)

									tmp3449 := PrimTail(tmp3448)

									tmp3450 := PrimHead(tmp3449)

									tmp3451 := PrimTail(tmp3450)

									tmp3452 := PrimHead(tmp3451)

									tmp3453 := PrimIsPair(tmp3452)

									var ifres3389 Obj

									if True == tmp3453 {
										tmp3441 := PrimTail(V857)

										tmp3442 := PrimTail(tmp3441)

										tmp3443 := PrimHead(tmp3442)

										tmp3444 := PrimTail(tmp3443)

										tmp3445 := PrimHead(tmp3444)

										tmp3446 := PrimHead(tmp3445)

										tmp3447 := PrimEqual(symthaw, tmp3446)

										var ifres3391 Obj

										if True == tmp3447 {
											tmp3434 := PrimTail(V857)

											tmp3435 := PrimTail(tmp3434)

											tmp3436 := PrimHead(tmp3435)

											tmp3437 := PrimTail(tmp3436)

											tmp3438 := PrimHead(tmp3437)

											tmp3439 := PrimTail(tmp3438)

											tmp3440 := PrimIsPair(tmp3439)

											var ifres3393 Obj

											if True == tmp3440 {
												tmp3426 := PrimTail(V857)

												tmp3427 := PrimTail(tmp3426)

												tmp3428 := PrimHead(tmp3427)

												tmp3429 := PrimTail(tmp3428)

												tmp3430 := PrimHead(tmp3429)

												tmp3431 := PrimTail(tmp3430)

												tmp3432 := PrimTail(tmp3431)

												tmp3433 := PrimEqual(Nil, tmp3432)

												var ifres3395 Obj

												if True == tmp3433 {
													tmp3420 := PrimTail(V857)

													tmp3421 := PrimTail(tmp3420)

													tmp3422 := PrimHead(tmp3421)

													tmp3423 := PrimTail(tmp3422)

													tmp3424 := PrimTail(tmp3423)

													tmp3425 := PrimEqual(Nil, tmp3424)

													var ifres3397 Obj

													if True == tmp3425 {
														tmp3416 := PrimTail(V857)

														tmp3417 := PrimTail(tmp3416)

														tmp3418 := PrimTail(tmp3417)

														tmp3419 := PrimIsPair(tmp3418)

														var ifres3399 Obj

														if True == tmp3419 {
															tmp3411 := PrimTail(V857)

															tmp3412 := PrimTail(tmp3411)

															tmp3413 := PrimTail(tmp3412)

															tmp3414 := PrimTail(tmp3413)

															tmp3415 := PrimEqual(Nil, tmp3414)

															var ifres3401 Obj

															if True == tmp3415 {
																tmp3403 := PrimTail(V857)

																tmp3404 := PrimTail(tmp3403)

																tmp3405 := PrimHead(tmp3404)

																tmp3406 := PrimTail(tmp3405)

																tmp3407 := PrimHead(tmp3406)

																tmp3408 := PrimTail(tmp3407)

																tmp3409 := PrimHead(tmp3408)

																tmp3410 := PrimIsSymbol(tmp3409)

																var ifres3402 Obj

																if True == tmp3410 {
																	ifres3402 = True

																} else {
																	ifres3402 = False

																}

																ifres3401 = ifres3402

															} else {
																ifres3401 = False

															}

															var ifres3400 Obj

															if True == ifres3401 {
																ifres3400 = True

															} else {
																ifres3400 = False

															}

															ifres3399 = ifres3400

														} else {
															ifres3399 = False

														}

														var ifres3398 Obj

														if True == ifres3399 {
															ifres3398 = True

														} else {
															ifres3398 = False

														}

														ifres3397 = ifres3398

													} else {
														ifres3397 = False

													}

													var ifres3396 Obj

													if True == ifres3397 {
														ifres3396 = True

													} else {
														ifres3396 = False

													}

													ifres3395 = ifres3396

												} else {
													ifres3395 = False

												}

												var ifres3394 Obj

												if True == ifres3395 {
													ifres3394 = True

												} else {
													ifres3394 = False

												}

												ifres3393 = ifres3394

											} else {
												ifres3393 = False

											}

											var ifres3392 Obj

											if True == ifres3393 {
												ifres3392 = True

											} else {
												ifres3392 = False

											}

											ifres3391 = ifres3392

										} else {
											ifres3391 = False

										}

										var ifres3390 Obj

										if True == ifres3391 {
											ifres3390 = True

										} else {
											ifres3390 = False

										}

										ifres3389 = ifres3390

									} else {
										ifres3389 = False

									}

									var ifres3388 Obj

									if True == ifres3389 {
										ifres3388 = True

									} else {
										ifres3388 = False

									}

									ifres3387 = ifres3388

								} else {
									ifres3387 = False

								}

								var ifres3386 Obj

								if True == ifres3387 {
									ifres3386 = True

								} else {
									ifres3386 = False

								}

								ifres3385 = ifres3386

							} else {
								ifres3385 = False

							}

							var ifres3384 Obj

							if True == ifres3385 {
								ifres3384 = True

							} else {
								ifres3384 = False

							}

							ifres3383 = ifres3384

						} else {
							ifres3383 = False

						}

						var ifres3382 Obj

						if True == ifres3383 {
							ifres3382 = True

						} else {
							ifres3382 = False

						}

						ifres3381 = ifres3382

					} else {
						ifres3381 = False

					}

					var ifres3380 Obj

					if True == ifres3381 {
						ifres3380 = True

					} else {
						ifres3380 = False

					}

					ifres3379 = ifres3380

				} else {
					ifres3379 = False

				}

				var ifres3378 Obj

				if True == ifres3379 {
					ifres3378 = True

				} else {
					ifres3378 = False

				}

				ifres3377 = ifres3378

			} else {
				ifres3377 = False

			}

			var ifres3376 Obj

			if True == ifres3377 {
				ifres3376 = True

			} else {
				ifres3376 = False

			}

			ifres3375 = ifres3376

		} else {
			ifres3375 = False

		}

		if True == ifres3375 {
			tmp3361 := PrimTail(V857)

			tmp3362 := PrimTail(tmp3361)

			tmp3363 := PrimHead(tmp3362)

			tmp3364 := PrimTail(tmp3363)

			tmp3365 := PrimHead(tmp3364)

			tmp3366 := PrimTail(tmp3365)

			tmp3367 := PrimHead(tmp3366)

			tmp3368 := PrimTail(V857)

			tmp3369 := PrimHead(tmp3368)

			tmp3370 := PrimTail(V857)

			tmp3371 := PrimTail(tmp3370)

			tmp3372 := PrimTail(tmp3371)

			tmp3373 := PrimHead(tmp3372)

			__e.TailApply(PrimFunc(symsubst), tmp3367, tmp3369, tmp3373)
			return

		} else {
			__e.Return(V857)
			return
		}

	}, 1)

	tmp3476 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp3360)

	_ = tmp3476

	tmp3477 := MakeNative(func(__e *ControlFlow) {
		V860 := __e.Get(1)
		_ = V860
		V861 := __e.Get(2)
		_ = V861
		V862 := __e.Get(3)
		_ = V862
		tmp3576 := PrimIsPair(V861)

		var ifres3522 Obj

		if True == tmp3576 {
			tmp3574 := PrimHead(V861)

			tmp3575 := PrimIsPair(tmp3574)

			var ifres3524 Obj

			if True == tmp3575 {
				tmp3571 := PrimHead(V861)

				tmp3572 := PrimHead(tmp3571)

				tmp3573 := PrimIsPair(tmp3572)

				var ifres3526 Obj

				if True == tmp3573 {
					tmp3567 := PrimHead(V861)

					tmp3568 := PrimHead(tmp3567)

					tmp3569 := PrimHead(tmp3568)

					tmp3570 := PrimEqual(symand, tmp3569)

					var ifres3528 Obj

					if True == tmp3570 {
						tmp3563 := PrimHead(V861)

						tmp3564 := PrimHead(tmp3563)

						tmp3565 := PrimTail(tmp3564)

						tmp3566 := PrimIsPair(tmp3565)

						var ifres3530 Obj

						if True == tmp3566 {
							tmp3558 := PrimHead(V861)

							tmp3559 := PrimHead(tmp3558)

							tmp3560 := PrimTail(tmp3559)

							tmp3561 := PrimTail(tmp3560)

							tmp3562 := PrimIsPair(tmp3561)

							var ifres3532 Obj

							if True == tmp3562 {
								tmp3552 := PrimHead(V861)

								tmp3553 := PrimHead(tmp3552)

								tmp3554 := PrimTail(tmp3553)

								tmp3555 := PrimTail(tmp3554)

								tmp3556 := PrimTail(tmp3555)

								tmp3557 := PrimEqual(Nil, tmp3556)

								var ifres3534 Obj

								if True == tmp3557 {
									tmp3549 := PrimHead(V861)

									tmp3550 := PrimTail(tmp3549)

									tmp3551 := PrimIsPair(tmp3550)

									var ifres3536 Obj

									if True == tmp3551 {
										tmp3545 := PrimHead(V861)

										tmp3546 := PrimTail(tmp3545)

										tmp3547 := PrimTail(tmp3546)

										tmp3548 := PrimEqual(Nil, tmp3547)

										var ifres3538 Obj

										if True == tmp3548 {
											tmp3540 := PrimHead(V861)

											tmp3541 := PrimHead(tmp3540)

											tmp3542 := PrimTail(tmp3541)

											tmp3543 := PrimHead(tmp3542)

											tmp3544 := PrimEqual(V860, tmp3543)

											var ifres3539 Obj

											if True == tmp3544 {
												ifres3539 = True

											} else {
												ifres3539 = False

											}

											ifres3538 = ifres3539

										} else {
											ifres3538 = False

										}

										var ifres3537 Obj

										if True == ifres3538 {
											ifres3537 = True

										} else {
											ifres3537 = False

										}

										ifres3536 = ifres3537

									} else {
										ifres3536 = False

									}

									var ifres3535 Obj

									if True == ifres3536 {
										ifres3535 = True

									} else {
										ifres3535 = False

									}

									ifres3534 = ifres3535

								} else {
									ifres3534 = False

								}

								var ifres3533 Obj

								if True == ifres3534 {
									ifres3533 = True

								} else {
									ifres3533 = False

								}

								ifres3532 = ifres3533

							} else {
								ifres3532 = False

							}

							var ifres3531 Obj

							if True == ifres3532 {
								ifres3531 = True

							} else {
								ifres3531 = False

							}

							ifres3530 = ifres3531

						} else {
							ifres3530 = False

						}

						var ifres3529 Obj

						if True == ifres3530 {
							ifres3529 = True

						} else {
							ifres3529 = False

						}

						ifres3528 = ifres3529

					} else {
						ifres3528 = False

					}

					var ifres3527 Obj

					if True == ifres3528 {
						ifres3527 = True

					} else {
						ifres3527 = False

					}

					ifres3526 = ifres3527

				} else {
					ifres3526 = False

				}

				var ifres3525 Obj

				if True == ifres3526 {
					ifres3525 = True

				} else {
					ifres3525 = False

				}

				ifres3524 = ifres3525

			} else {
				ifres3524 = False

			}

			var ifres3523 Obj

			if True == ifres3524 {
				ifres3523 = True

			} else {
				ifres3523 = False

			}

			ifres3522 = ifres3523

		} else {
			ifres3522 = False

		}

		if True == ifres3522 {
			tmp3478 := PrimHead(V861)

			tmp3479 := PrimHead(tmp3478)

			tmp3480 := PrimTail(tmp3479)

			tmp3481 := PrimHead(tmp3480)

			tmp3482 := PrimTail(V861)

			tmp3483 := PrimHead(V861)

			tmp3484 := PrimHead(tmp3483)

			tmp3485 := PrimTail(tmp3484)

			tmp3486 := PrimTail(tmp3485)

			tmp3487 := PrimHead(tmp3486)

			tmp3488 := PrimHead(V861)

			tmp3489 := PrimTail(tmp3488)

			tmp3490 := PrimCons(tmp3487, tmp3489)

			tmp3491 := PrimCons(tmp3490, V862)

			__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp3481, tmp3482, tmp3491)
			return

		} else {
			tmp3520 := PrimIsPair(V861)

			var ifres3500 Obj

			if True == tmp3520 {
				tmp3518 := PrimHead(V861)

				tmp3519 := PrimIsPair(tmp3518)

				var ifres3502 Obj

				if True == tmp3519 {
					tmp3515 := PrimHead(V861)

					tmp3516 := PrimTail(tmp3515)

					tmp3517 := PrimIsPair(tmp3516)

					var ifres3504 Obj

					if True == tmp3517 {
						tmp3511 := PrimHead(V861)

						tmp3512 := PrimTail(tmp3511)

						tmp3513 := PrimTail(tmp3512)

						tmp3514 := PrimEqual(Nil, tmp3513)

						var ifres3506 Obj

						if True == tmp3514 {
							tmp3508 := PrimHead(V861)

							tmp3509 := PrimHead(tmp3508)

							tmp3510 := PrimEqual(V860, tmp3509)

							var ifres3507 Obj

							if True == tmp3510 {
								ifres3507 = True

							} else {
								ifres3507 = False

							}

							ifres3506 = ifres3507

						} else {
							ifres3506 = False

						}

						var ifres3505 Obj

						if True == ifres3506 {
							ifres3505 = True

						} else {
							ifres3505 = False

						}

						ifres3504 = ifres3505

					} else {
						ifres3504 = False

					}

					var ifres3503 Obj

					if True == ifres3504 {
						ifres3503 = True

					} else {
						ifres3503 = False

					}

					ifres3502 = ifres3503

				} else {
					ifres3502 = False

				}

				var ifres3501 Obj

				if True == ifres3502 {
					ifres3501 = True

				} else {
					ifres3501 = False

				}

				ifres3500 = ifres3501

			} else {
				ifres3500 = False

			}

			if True == ifres3500 {
				tmp3492 := PrimHead(V861)

				tmp3493 := PrimHead(tmp3492)

				tmp3494 := PrimTail(V861)

				tmp3495 := PrimHead(V861)

				tmp3496 := PrimTail(tmp3495)

				tmp3497 := PrimCons(True, tmp3496)

				tmp3498 := PrimCons(tmp3497, V862)

				__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp3493, tmp3494, tmp3498)
				return

			} else {
				__e.TailApply(PrimFunc(sym_8p), V862, V861)
				return
			}

		}

	}, 3)

	tmp3577 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp3477)

	_ = tmp3577

	tmp3578 := MakeNative(func(__e *ControlFlow) {
		V865 := __e.Get(1)
		_ = V865
		V866 := __e.Get(2)
		_ = V866
		tmp3602 := PrimIsPair(V865)

		var ifres3593 Obj

		if True == tmp3602 {
			tmp3600 := PrimTail(V865)

			tmp3601 := PrimIsPair(tmp3600)

			var ifres3595 Obj

			if True == tmp3601 {
				tmp3597 := PrimTail(V865)

				tmp3598 := PrimTail(tmp3597)

				tmp3599 := PrimEqual(Nil, tmp3598)

				var ifres3596 Obj

				if True == tmp3599 {
					ifres3596 = True

				} else {
					ifres3596 = False

				}

				ifres3595 = ifres3596

			} else {
				ifres3595 = False

			}

			var ifres3594 Obj

			if True == ifres3595 {
				ifres3594 = True

			} else {
				ifres3594 = False

			}

			ifres3593 = ifres3594

		} else {
			ifres3593 = False

		}

		if True == ifres3593 {
			tmp3579 := MakeNative(func(__e *ControlFlow) {
				W867 := __e.Get(1)
				_ = W867
				tmp3589 := PrimEqual(symshen_4skip, W867)

				if True == tmp3589 {
					__e.Return(V866)
					return
				} else {
					tmp3580 := Call(__e, PrimFunc(symshen_4op1), W867)

					tmp3581 := PrimTail(V865)

					tmp3582 := PrimCons(tmp3580, tmp3581)

					tmp3583 := Call(__e, PrimFunc(symshen_4op2), W867)

					tmp3584 := PrimTail(V865)

					tmp3585 := PrimCons(tmp3583, tmp3584)

					tmp3586 := PrimCons(tmp3585, Nil)

					tmp3587 := PrimCons(tmp3582, tmp3586)

					__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp3587, V866)
					return

				}

			}, 1)

			tmp3590 := PrimHead(V865)

			tmp3591 := Call(__e, PrimFunc(symshen_4op), tmp3590)

			__e.TailApply(tmp3579, tmp3591)
			return

		} else {
			__e.Return(V866)
			return
		}

	}, 2)

	tmp3603 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp3578)

	_ = tmp3603

	tmp3604 := MakeNative(func(__e *ControlFlow) {
		V870 := __e.Get(1)
		_ = V870
		tmp3612 := PrimEqual(symcons_2, V870)

		if True == tmp3612 {
			__e.Return(symcons)
			return
		} else {
			tmp3610 := PrimEqual(symshen_4_7string_2, V870)

			if True == tmp3610 {
				__e.Return(sym_8s)
				return
			} else {
				tmp3608 := PrimEqual(symshen_4_7vector_2, V870)

				if True == tmp3608 {
					__e.Return(sym_8v)
					return
				} else {
					tmp3606 := PrimEqual(symtuple_2, V870)

					if True == tmp3606 {
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

	tmp3613 := Call(__e, ns2_1set, symshen_4op, tmp3604)

	_ = tmp3613

	tmp3614 := MakeNative(func(__e *ControlFlow) {
		V871 := __e.Get(1)
		_ = V871
		V872 := __e.Get(2)
		_ = V872
		tmp3633 := PrimEqual(Nil, V871)

		if True == tmp3633 {
			__e.Return(V872)
			return
		} else {
			tmp3631 := PrimIsPair(V871)

			if True == tmp3631 {
				tmp3627 := PrimHead(V871)

				tmp3628 := Call(__e, PrimFunc(symoccurrences), tmp3627, V872)

				tmp3629 := PrimGreatThan(tmp3628, MakeNumber(1))

				if True == tmp3629 {
					tmp3615 := MakeNative(func(__e *ControlFlow) {
						W873 := __e.Get(1)
						_ = W873
						tmp3616 := PrimHead(V871)

						tmp3617 := PrimTail(V871)

						tmp3618 := PrimHead(V871)

						tmp3619 := Call(__e, PrimFunc(symsubst), W873, tmp3618, V872)

						tmp3620 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp3617, tmp3619)

						tmp3621 := PrimCons(tmp3620, Nil)

						tmp3622 := PrimCons(tmp3616, tmp3621)

						tmp3623 := PrimCons(W873, tmp3622)

						__e.Return(PrimCons(symlet, tmp3623))
						return

					}, 1)

					tmp3624 := Call(__e, PrimFunc(symgensym), symSelect)

					__e.TailApply(tmp3615, tmp3624)
					return

				} else {
					tmp3625 := PrimTail(V871)

					__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp3625, V872)
					return

				}

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1selectors_1h)
				return
			}

		}

	}, 2)

	__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp3614)
	return

}, 0)
