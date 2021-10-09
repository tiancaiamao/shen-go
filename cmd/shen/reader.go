package main

import . "github.com/tiancaiamao/shen-go/cora"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp12857 := MakeNative(func(__e *ControlFlow) {
		V1430 := __e.Get(1)
		_ = V1430
		__e.Return(PrimReadByte(V1430))
		return
	}, 1)

	tmp12858 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1char_1code, tmp12857)

	_ = tmp12858

	tmp12859 := MakeNative(func(__e *ControlFlow) {
		V1432 := __e.Get(1)
		_ = V1432
		tmp12860 := MakeNative(func(__e *ControlFlow) {
			S := __e.Get(1)
			_ = S
			__e.Return(PrimReadByte(S))
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symshen_4read_1file_1as_1Xlist), V1432, tmp12860)
		return

	}, 1)

	tmp12861 := Call(__e, PrimNS1Value(symns2_1set), symread_1file_1as_1bytelist, tmp12859)

	_ = tmp12861

	tmp12862 := MakeNative(func(__e *ControlFlow) {
		V1434 := __e.Get(1)
		_ = V1434
		tmp12863 := MakeNative(func(__e *ControlFlow) {
			S := __e.Get(1)
			_ = S
			__e.TailApply(PrimNS2Value(symshen_4read_1char_1code), S)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symshen_4read_1file_1as_1Xlist), V1434, tmp12863)
		return

	}, 1)

	tmp12864 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1charlist, tmp12862)

	_ = tmp12864

	tmp12865 := MakeNative(func(__e *ControlFlow) {
		V1437 := __e.Get(1)
		_ = V1437
		V1438 := __e.Get(2)
		_ = V1438
		tmp12866 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp12867 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp12868 := MakeNative(func(__e *ControlFlow) {
					Xs := __e.Get(1)
					_ = Xs
					tmp12869 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.TailApply(PrimNS2Value(symreverse), Xs)
						return
					}, 1)

					tmp12870 := PrimCloseStream(Stream)

					__e.TailApply(tmp12869, tmp12870)
					return

				}, 1)

				tmp12871 := Call(__e, PrimNS2Value(symshen_4read_1file_1as_1Xlist_1help), Stream, V1438, X, Nil)

				__e.TailApply(tmp12868, tmp12871)
				return

			}, 1)

			tmp12872 := Call(__e, V1438, Stream)

			__e.TailApply(tmp12867, tmp12872)
			return

		}, 1)

		tmp12873 := PrimOpenStream(V1437, symin)

		__e.TailApply(tmp12866, tmp12873)
		return

	}, 2)

	tmp12874 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1Xlist, tmp12865)

	_ = tmp12874

	tmp12875 := MakeNative(func(__e *ControlFlow) {
		V1443 := __e.Get(1)
		_ = V1443
		V1444 := __e.Get(2)
		_ = V1444
		V1445 := __e.Get(3)
		_ = V1445
		V1446 := __e.Get(4)
		_ = V1446
		tmp12879 := PrimEqual(MakeNumber(-1), V1445)

		if True == tmp12879 {
			__e.Return(V1446)
			return
		} else {
			tmp12877 := Call(__e, V1444, V1443)

			tmp12878 := PrimCons(V1445, V1446)

			__e.TailApply(PrimNS2Value(symshen_4read_1file_1as_1Xlist_1help), V1443, V1444, tmp12877, tmp12878)
			return

		}

	}, 4)

	tmp12880 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1Xlist_1help, tmp12875)

	_ = tmp12880

	tmp12881 := MakeNative(func(__e *ControlFlow) {
		V1448 := __e.Get(1)
		_ = V1448
		tmp12882 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp12883 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), Stream)

			__e.TailApply(PrimNS2Value(symshen_4rfas_1h), Stream, tmp12883, MakeString(""))
			return

		}, 1)

		tmp12884 := PrimOpenStream(V1448, symin)

		__e.TailApply(tmp12882, tmp12884)
		return

	}, 1)

	tmp12885 := Call(__e, PrimNS1Value(symns2_1set), symread_1file_1as_1string, tmp12881)

	_ = tmp12885

	tmp12886 := MakeNative(func(__e *ControlFlow) {
		V1452 := __e.Get(1)
		_ = V1452
		V1453 := __e.Get(2)
		_ = V1453
		V1454 := __e.Get(3)
		_ = V1454
		tmp12892 := PrimEqual(MakeNumber(-1), V1453)

		if True == tmp12892 {
			tmp12891 := PrimCloseStream(V1452)

			_ = tmp12891

			__e.Return(V1454)
			return

		} else {
			tmp12888 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1452)

			tmp12889 := PrimNumberToString(V1453)

			tmp12890 := PrimStringConcat(V1454, tmp12889)

			__e.TailApply(PrimNS2Value(symshen_4rfas_1h), V1452, tmp12888, tmp12890)
			return

		}

	}, 3)

	tmp12893 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rfas_1h, tmp12886)

	_ = tmp12893

	tmp12894 := MakeNative(func(__e *ControlFlow) {
		V1456 := __e.Get(1)
		_ = V1456
		tmp12895 := Call(__e, PrimNS2Value(symread), V1456)

		__e.TailApply(PrimNS2Value(symeval_1kl), tmp12895)
		return

	}, 1)

	tmp12896 := Call(__e, PrimNS1Value(symns2_1set), syminput, tmp12894)

	_ = tmp12896

	tmp12897 := MakeNative(func(__e *ControlFlow) {
		V1459 := __e.Get(1)
		_ = V1459
		V1460 := __e.Get(2)
		_ = V1460
		tmp12898 := MakeNative(func(__e *ControlFlow) {
			Mono_2 := __e.Get(1)
			_ = Mono_2
			tmp12899 := MakeNative(func(__e *ControlFlow) {
				Input := __e.Get(1)
				_ = Input
				tmp12905 := Call(__e, PrimNS2Value(symshen_4demodulate), V1459)

				tmp12906 := Call(__e, PrimNS2Value(symshen_4typecheck), Input, tmp12905)

				tmp12907 := PrimEqual(False, tmp12906)

				if True == tmp12907 {
					tmp12901 := Call(__e, PrimNS2Value(symshen_4app), V1459, MakeString("\n"), symshen_4r)

					tmp12902 := PrimStringConcat(MakeString(" is not of type "), tmp12901)

					tmp12903 := Call(__e, PrimNS2Value(symshen_4app), Input, tmp12902, symshen_4r)

					tmp12904 := PrimStringConcat(MakeString("type error: "), tmp12903)

					__e.Return(PrimSimpleError(tmp12904))
					return

				} else {
					__e.TailApply(PrimNS2Value(symeval_1kl), Input)
					return
				}

			}, 1)

			tmp12908 := Call(__e, PrimNS2Value(symread), V1460)

			__e.TailApply(tmp12899, tmp12908)
			return

		}, 1)

		tmp12909 := Call(__e, PrimNS2Value(symshen_4monotype), V1459)

		__e.TailApply(tmp12898, tmp12909)
		return

	}, 2)

	tmp12910 := Call(__e, PrimNS1Value(symns2_1set), syminput_7, tmp12897)

	_ = tmp12910

	tmp12911 := MakeNative(func(__e *ControlFlow) {
		V1462 := __e.Get(1)
		_ = V1462
		tmp12918 := PrimIsPair(V1462)

		if True == tmp12918 {
			tmp12917 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimNS2Value(symshen_4monotype), Z)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp12917, V1462)
			return

		} else {
			tmp12916 := PrimIsVariable(V1462)

			if True == tmp12916 {
				tmp12914 := Call(__e, PrimNS2Value(symshen_4app), V1462, MakeString("\n"), symshen_4a)

				tmp12915 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp12914)

				__e.Return(PrimSimpleError(tmp12915))
				return

			} else {
				__e.Return(V1462)
				return
			}

		}

	}, 1)

	tmp12919 := Call(__e, PrimNS1Value(symns2_1set), symshen_4monotype, tmp12911)

	_ = tmp12919

	tmp12920 := MakeNative(func(__e *ControlFlow) {
		V1464 := __e.Get(1)
		_ = V1464
		tmp12921 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1464)

		tmp12922 := Call(__e, PrimNS2Value(symshen_4read_1loop), V1464, tmp12921, Nil)

		__e.Return(PrimHead(tmp12922))
		return

	}, 1)

	tmp12923 := Call(__e, PrimNS1Value(symns2_1set), symread, tmp12920)

	_ = tmp12923

	tmp12924 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dit_d))
		return
	}, 0)

	tmp12925 := Call(__e, PrimNS1Value(symns2_1set), symit, tmp12924)

	_ = tmp12925

	tmp12926 := MakeNative(func(__e *ControlFlow) {
		V1472 := __e.Get(1)
		_ = V1472
		V1473 := __e.Get(2)
		_ = V1473
		V1474 := __e.Get(3)
		_ = V1474
		tmp12954 := PrimEqual(MakeNumber(94), V1473)

		if True == tmp12954 {
			__e.Return(PrimSimpleError(MakeString("read aborted")))
			return
		} else {
			tmp12953 := PrimEqual(MakeNumber(-1), V1473)

			if True == tmp12953 {
				tmp12952 := Call(__e, PrimNS2Value(symempty_2), V1474)

				if True == tmp12952 {
					__e.Return(PrimSimpleError(MakeString("error: empty stream")))
					return
				} else {
					tmp12950 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
						return
					}, 1)

					tmp12951 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(E)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symcompile), tmp12950, V1474, tmp12951)
					return

				}

			} else {
				tmp12948 := Call(__e, PrimNS2Value(symshen_4terminator_2), V1473)

				if True == tmp12948 {
					tmp12933 := MakeNative(func(__e *ControlFlow) {
						AllChars := __e.Get(1)
						_ = AllChars
						tmp12934 := MakeNative(func(__e *ControlFlow) {
							It := __e.Get(1)
							_ = It
							tmp12935 := MakeNative(func(__e *ControlFlow) {
								Read := __e.Get(1)
								_ = Read
								tmp12941 := PrimEqual(Read, symshen_4nextbyte)

								var ifres12938 Obj

								if True == tmp12941 {
									ifres12938 = True

								} else {
									tmp12940 := Call(__e, PrimNS2Value(symempty_2), Read)

									var ifres12939 Obj

									if True == tmp12940 {
										ifres12939 = True

									} else {
										ifres12939 = False

									}

									ifres12938 = ifres12939

								}

								if True == ifres12938 {
									tmp12937 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1472)

									__e.TailApply(PrimNS2Value(symshen_4read_1loop), V1472, tmp12937, AllChars)
									return

								} else {
									__e.Return(Read)
									return
								}

							}, 1)

							tmp12942 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
								return
							}, 1)

							tmp12943 := MakeNative(func(__e *ControlFlow) {
								E := __e.Get(1)
								_ = E
								__e.Return(symshen_4nextbyte)
								return
							}, 1)

							tmp12944 := Call(__e, PrimNS2Value(symcompile), tmp12942, AllChars, tmp12943)

							__e.TailApply(tmp12935, tmp12944)
							return

						}, 1)

						tmp12945 := Call(__e, PrimNS2Value(symshen_4record_1it), AllChars)

						__e.TailApply(tmp12934, tmp12945)
						return

					}, 1)

					tmp12946 := PrimCons(V1473, Nil)

					tmp12947 := Call(__e, PrimNS2Value(symappend), V1474, tmp12946)

					__e.TailApply(tmp12933, tmp12947)
					return

				} else {
					tmp12930 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1472)

					tmp12931 := PrimCons(V1473, Nil)

					tmp12932 := Call(__e, PrimNS2Value(symappend), V1474, tmp12931)

					__e.TailApply(PrimNS2Value(symshen_4read_1loop), V1472, tmp12930, tmp12932)
					return

				}

			}

		}

	}, 3)

	tmp12955 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1loop, tmp12926)

	_ = tmp12955

	tmp12956 := MakeNative(func(__e *ControlFlow) {
		V1476 := __e.Get(1)
		_ = V1476
		tmp12957 := PrimCons(MakeNumber(93), Nil)

		tmp12958 := PrimCons(MakeNumber(41), tmp12957)

		tmp12959 := PrimCons(MakeNumber(34), tmp12958)

		tmp12960 := PrimCons(MakeNumber(32), tmp12959)

		tmp12961 := PrimCons(MakeNumber(13), tmp12960)

		tmp12962 := PrimCons(MakeNumber(10), tmp12961)

		tmp12963 := PrimCons(MakeNumber(9), tmp12962)

		__e.TailApply(PrimNS2Value(symelement_2), V1476, tmp12963)
		return

	}, 1)

	tmp12964 := Call(__e, PrimNS1Value(symns2_1set), symshen_4terminator_2, tmp12956)

	_ = tmp12964

	tmp12965 := MakeNative(func(__e *ControlFlow) {
		V1478 := __e.Get(1)
		_ = V1478
		tmp12966 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1478)

		__e.TailApply(PrimNS2Value(symshen_4lineread_1loop), tmp12966, Nil, V1478)
		return

	}, 1)

	tmp12967 := Call(__e, PrimNS1Value(symns2_1set), symlineread, tmp12965)

	_ = tmp12967

	tmp12968 := MakeNative(func(__e *ControlFlow) {
		V1483 := __e.Get(1)
		_ = V1483
		V1484 := __e.Get(2)
		_ = V1484
		V1485 := __e.Get(3)
		_ = V1485
		tmp13000 := PrimEqual(MakeNumber(-1), V1483)

		if True == tmp13000 {
			tmp12999 := Call(__e, PrimNS2Value(symempty_2), V1484)

			if True == tmp12999 {
				__e.Return(PrimSimpleError(MakeString("empty stream")))
				return
			} else {
				tmp12997 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
					return
				}, 1)

				tmp12998 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(E)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symcompile), tmp12997, V1484, tmp12998)
				return

			}

		} else {
			tmp12994 := Call(__e, PrimNS2Value(symshen_4hat))

			tmp12995 := PrimEqual(V1483, tmp12994)

			if True == tmp12995 {
				__e.Return(PrimSimpleError(MakeString("line read aborted")))
				return
			} else {
				tmp12989 := Call(__e, PrimNS2Value(symshen_4newline))

				tmp12990 := Call(__e, PrimNS2Value(symshen_4carriage_1return))

				tmp12991 := PrimCons(tmp12990, Nil)

				tmp12992 := PrimCons(tmp12989, tmp12991)

				tmp12993 := Call(__e, PrimNS2Value(symelement_2), V1483, tmp12992)

				if True == tmp12993 {
					tmp12975 := MakeNative(func(__e *ControlFlow) {
						Line := __e.Get(1)
						_ = Line
						tmp12976 := MakeNative(func(__e *ControlFlow) {
							It := __e.Get(1)
							_ = It
							tmp12984 := PrimEqual(Line, symshen_4nextline)

							var ifres12981 Obj

							if True == tmp12984 {
								ifres12981 = True

							} else {
								tmp12983 := Call(__e, PrimNS2Value(symempty_2), Line)

								var ifres12982 Obj

								if True == tmp12983 {
									ifres12982 = True

								} else {
									ifres12982 = False

								}

								ifres12981 = ifres12982

							}

							if True == ifres12981 {
								tmp12978 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1485)

								tmp12979 := PrimCons(V1483, Nil)

								tmp12980 := Call(__e, PrimNS2Value(symappend), V1484, tmp12979)

								__e.TailApply(PrimNS2Value(symshen_4lineread_1loop), tmp12978, tmp12980, V1485)
								return

							} else {
								__e.Return(Line)
								return
							}

						}, 1)

						tmp12985 := Call(__e, PrimNS2Value(symshen_4record_1it), V1484)

						__e.TailApply(tmp12976, tmp12985)
						return

					}, 1)

					tmp12986 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
						return
					}, 1)

					tmp12987 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(symshen_4nextline)
						return
					}, 1)

					tmp12988 := Call(__e, PrimNS2Value(symcompile), tmp12986, V1484, tmp12987)

					__e.TailApply(tmp12975, tmp12988)
					return

				} else {
					tmp12972 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), V1485)

					tmp12973 := PrimCons(V1483, Nil)

					tmp12974 := Call(__e, PrimNS2Value(symappend), V1484, tmp12973)

					__e.TailApply(PrimNS2Value(symshen_4lineread_1loop), tmp12972, tmp12974, V1485)
					return

				}

			}

		}

	}, 3)

	tmp13001 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lineread_1loop, tmp12968)

	_ = tmp13001

	tmp13002 := MakeNative(func(__e *ControlFlow) {
		V1487 := __e.Get(1)
		_ = V1487
		tmp13003 := MakeNative(func(__e *ControlFlow) {
			TrimLeft := __e.Get(1)
			_ = TrimLeft
			tmp13004 := MakeNative(func(__e *ControlFlow) {
				TrimRight := __e.Get(1)
				_ = TrimRight
				tmp13005 := MakeNative(func(__e *ControlFlow) {
					Trimmed := __e.Get(1)
					_ = Trimmed
					__e.TailApply(PrimNS2Value(symshen_4record_1it_1h), Trimmed)
					return
				}, 1)

				tmp13006 := Call(__e, PrimNS2Value(symreverse), TrimRight)

				__e.TailApply(tmp13005, tmp13006)
				return

			}, 1)

			tmp13007 := Call(__e, PrimNS2Value(symreverse), TrimLeft)

			tmp13008 := Call(__e, PrimNS2Value(symshen_4trim_1whitespace), tmp13007)

			__e.TailApply(tmp13004, tmp13008)
			return

		}, 1)

		tmp13009 := Call(__e, PrimNS2Value(symshen_4trim_1whitespace), V1487)

		__e.TailApply(tmp13003, tmp13009)
		return

	}, 1)

	tmp13010 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1it, tmp13002)

	_ = tmp13010

	tmp13011 := MakeNative(func(__e *ControlFlow) {
		V1489 := __e.Get(1)
		_ = V1489
		tmp13022 := PrimIsPair(V1489)

		var ifres13014 Obj

		if True == tmp13022 {
			tmp13016 := PrimHead(V1489)

			tmp13017 := PrimCons(MakeNumber(32), Nil)

			tmp13018 := PrimCons(MakeNumber(13), tmp13017)

			tmp13019 := PrimCons(MakeNumber(10), tmp13018)

			tmp13020 := PrimCons(MakeNumber(9), tmp13019)

			tmp13021 := Call(__e, PrimNS2Value(symelement_2), tmp13016, tmp13020)

			var ifres13015 Obj

			if True == tmp13021 {
				ifres13015 = True

			} else {
				ifres13015 = False

			}

			ifres13014 = ifres13015

		} else {
			ifres13014 = False

		}

		if True == ifres13014 {
			tmp13013 := PrimTail(V1489)

			__e.TailApply(PrimNS2Value(symshen_4trim_1whitespace), tmp13013)
			return

		} else {
			__e.Return(V1489)
			return
		}

	}, 1)

	tmp13023 := Call(__e, PrimNS1Value(symns2_1set), symshen_4trim_1whitespace, tmp13011)

	_ = tmp13023

	tmp13024 := MakeNative(func(__e *ControlFlow) {
		V1491 := __e.Get(1)
		_ = V1491
		tmp13025 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(PrimNumberToString(X))
			return
		}, 1)

		tmp13026 := Call(__e, PrimNS2Value(symmap), tmp13025, V1491)

		tmp13027 := Call(__e, PrimNS2Value(symshen_4cn_1all), tmp13026)

		tmp13028 := PrimNS3Set(symshen_4_dit_d, tmp13027)

		_ = tmp13028

		__e.Return(V1491)
		return

	}, 1)

	tmp13029 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1it_1h, tmp13024)

	_ = tmp13029

	tmp13030 := MakeNative(func(__e *ControlFlow) {
		V1493 := __e.Get(1)
		_ = V1493
		tmp13037 := PrimEqual(Nil, V1493)

		if True == tmp13037 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp13036 := PrimIsPair(V1493)

			if True == tmp13036 {
				tmp13033 := PrimHead(V1493)

				tmp13034 := PrimTail(V1493)

				tmp13035 := Call(__e, PrimNS2Value(symshen_4cn_1all), tmp13034)

				__e.Return(PrimStringConcat(tmp13033, tmp13035))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4cn_1all)
				return
			}

		}

	}, 1)

	tmp13038 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cn_1all, tmp13030)

	_ = tmp13038

	tmp13039 := MakeNative(func(__e *ControlFlow) {
		V1495 := __e.Get(1)
		_ = V1495
		tmp13040 := MakeNative(func(__e *ControlFlow) {
			Charlist := __e.Get(1)
			_ = Charlist
			tmp13041 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
				return
			}, 1)

			tmp13042 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4read_1error), X)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symcompile), tmp13041, Charlist, tmp13042)
			return

		}, 1)

		tmp13043 := Call(__e, PrimNS2Value(symshen_4read_1file_1as_1charlist), V1495)

		__e.TailApply(tmp13040, tmp13043)
		return

	}, 1)

	tmp13044 := Call(__e, PrimNS1Value(symns2_1set), symread_1file, tmp13039)

	_ = tmp13044

	tmp13045 := MakeNative(func(__e *ControlFlow) {
		V1497 := __e.Get(1)
		_ = V1497
		tmp13046 := MakeNative(func(__e *ControlFlow) {
			Ns := __e.Get(1)
			_ = Ns
			tmp13047 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
				return
			}, 1)

			tmp13048 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4read_1error), X)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symcompile), tmp13047, Ns, tmp13048)
			return

		}, 1)

		tmp13049 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(PrimStringToNumber(X))
			return
		}, 1)

		tmp13050 := Call(__e, PrimNS2Value(symexplode), V1497)

		tmp13051 := Call(__e, PrimNS2Value(symmap), tmp13049, tmp13050)

		__e.TailApply(tmp13046, tmp13051)
		return

	}, 1)

	tmp13052 := Call(__e, PrimNS1Value(symns2_1set), symread_1from_1string, tmp13045)

	_ = tmp13052

	tmp13053 := MakeNative(func(__e *ControlFlow) {
		V1505 := __e.Get(1)
		_ = V1505
		tmp13072 := PrimIsPair(V1505)

		var ifres13059 Obj

		if True == tmp13072 {
			tmp13070 := PrimHead(V1505)

			tmp13071 := PrimIsPair(tmp13070)

			var ifres13061 Obj

			if True == tmp13071 {
				tmp13068 := PrimTail(V1505)

				tmp13069 := PrimIsPair(tmp13068)

				var ifres13063 Obj

				if True == tmp13069 {
					tmp13065 := PrimTail(V1505)

					tmp13066 := PrimTail(tmp13065)

					tmp13067 := PrimEqual(Nil, tmp13066)

					var ifres13064 Obj

					if True == tmp13067 {
						ifres13064 = True

					} else {
						ifres13064 = False

					}

					ifres13063 = ifres13064

				} else {
					ifres13063 = False

				}

				var ifres13062 Obj

				if True == ifres13063 {
					ifres13062 = True

				} else {
					ifres13062 = False

				}

				ifres13061 = ifres13062

			} else {
				ifres13061 = False

			}

			var ifres13060 Obj

			if True == ifres13061 {
				ifres13060 = True

			} else {
				ifres13060 = False

			}

			ifres13059 = ifres13060

		} else {
			ifres13059 = False

		}

		if True == ifres13059 {
			tmp13055 := PrimHead(V1505)

			tmp13056 := Call(__e, PrimNS2Value(symshen_4compress_150), MakeNumber(50), tmp13055)

			tmp13057 := Call(__e, PrimNS2Value(symshen_4app), tmp13056, MakeString("\n"), symshen_4a)

			tmp13058 := PrimStringConcat(MakeString("read error here:\n\n "), tmp13057)

			__e.Return(PrimSimpleError(tmp13058))
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("read error\n")))
			return
		}

	}, 1)

	tmp13073 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1error, tmp13053)

	_ = tmp13073

	tmp13074 := MakeNative(func(__e *ControlFlow) {
		V1512 := __e.Get(1)
		_ = V1512
		V1513 := __e.Get(2)
		_ = V1513
		tmp13085 := PrimEqual(Nil, V1513)

		if True == tmp13085 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp13084 := PrimEqual(MakeNumber(0), V1512)

			if True == tmp13084 {
				__e.Return(MakeString(""))
				return
			} else {
				tmp13083 := PrimIsPair(V1513)

				if True == tmp13083 {
					tmp13078 := PrimHead(V1513)

					tmp13079 := PrimNumberToString(tmp13078)

					tmp13080 := PrimNumberSubtract(V1512, MakeNumber(1))

					tmp13081 := PrimTail(V1513)

					tmp13082 := Call(__e, PrimNS2Value(symshen_4compress_150), tmp13080, tmp13081)

					__e.Return(PrimStringConcat(tmp13079, tmp13082))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4compress_150)
					return
				}

			}

		}

	}, 2)

	tmp13086 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compress_150, tmp13074)

	_ = tmp13086

	tmp13087 := MakeNative(func(__e *ControlFlow) {
		V1515 := __e.Get(1)
		_ = V1515
		tmp13088 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp13364 := Call(__e, PrimNS2Value(symfail))

			tmp13365 := PrimEqual(YaccParse, tmp13364)

			if True == tmp13365 {
				tmp13090 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp13332 := Call(__e, PrimNS2Value(symfail))

					tmp13333 := PrimEqual(YaccParse, tmp13332)

					if True == tmp13333 {
						tmp13092 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp13314 := Call(__e, PrimNS2Value(symfail))

							tmp13315 := PrimEqual(YaccParse, tmp13314)

							if True == tmp13315 {
								tmp13094 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									tmp13296 := Call(__e, PrimNS2Value(symfail))

									tmp13297 := PrimEqual(YaccParse, tmp13296)

									if True == tmp13297 {
										tmp13096 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											tmp13278 := Call(__e, PrimNS2Value(symfail))

											tmp13279 := PrimEqual(YaccParse, tmp13278)

											if True == tmp13279 {
												tmp13098 := MakeNative(func(__e *ControlFlow) {
													YaccParse := __e.Get(1)
													_ = YaccParse
													tmp13260 := Call(__e, PrimNS2Value(symfail))

													tmp13261 := PrimEqual(YaccParse, tmp13260)

													if True == tmp13261 {
														tmp13100 := MakeNative(func(__e *ControlFlow) {
															YaccParse := __e.Get(1)
															_ = YaccParse
															tmp13236 := Call(__e, PrimNS2Value(symfail))

															tmp13237 := PrimEqual(YaccParse, tmp13236)

															if True == tmp13237 {
																tmp13102 := MakeNative(func(__e *ControlFlow) {
																	YaccParse := __e.Get(1)
																	_ = YaccParse
																	tmp13212 := Call(__e, PrimNS2Value(symfail))

																	tmp13213 := PrimEqual(YaccParse, tmp13212)

																	if True == tmp13213 {
																		tmp13104 := MakeNative(func(__e *ControlFlow) {
																			YaccParse := __e.Get(1)
																			_ = YaccParse
																			tmp13194 := Call(__e, PrimNS2Value(symfail))

																			tmp13195 := PrimEqual(YaccParse, tmp13194)

																			if True == tmp13195 {
																				tmp13106 := MakeNative(func(__e *ControlFlow) {
																					YaccParse := __e.Get(1)
																					_ = YaccParse
																					tmp13175 := Call(__e, PrimNS2Value(symfail))

																					tmp13176 := PrimEqual(YaccParse, tmp13175)

																					if True == tmp13176 {
																						tmp13108 := MakeNative(func(__e *ControlFlow) {
																							YaccParse := __e.Get(1)
																							_ = YaccParse
																							tmp13158 := Call(__e, PrimNS2Value(symfail))

																							tmp13159 := PrimEqual(YaccParse, tmp13158)

																							if True == tmp13159 {
																								tmp13110 := MakeNative(func(__e *ControlFlow) {
																									YaccParse := __e.Get(1)
																									_ = YaccParse
																									tmp13138 := Call(__e, PrimNS2Value(symfail))

																									tmp13139 := PrimEqual(YaccParse, tmp13138)

																									if True == tmp13139 {
																										tmp13112 := MakeNative(func(__e *ControlFlow) {
																											YaccParse := __e.Get(1)
																											_ = YaccParse
																											tmp13121 := Call(__e, PrimNS2Value(symfail))

																											tmp13122 := PrimEqual(YaccParse, tmp13121)

																											if True == tmp13122 {
																												tmp13114 := MakeNative(func(__e *ControlFlow) {
																													Parse___5e_6 := __e.Get(1)
																													_ = Parse___5e_6
																													tmp13117 := Call(__e, PrimNS2Value(symfail))

																													tmp13118 := PrimEqual(tmp13117, Parse___5e_6)

																													tmp13119 := PrimNot(tmp13118)

																													if True == tmp13119 {
																														tmp13116 := PrimHead(Parse___5e_6)

																														__e.TailApply(PrimNS2Value(symshen_4pair), tmp13116, Nil)
																														return

																													} else {
																														__e.TailApply(PrimNS2Value(symfail))
																														return
																													}

																												}, 1)

																												tmp13120 := Call(__e, PrimNS2Value(sym_5e_6), V1515)

																												__e.TailApply(tmp13114, tmp13120)
																												return

																											} else {
																												__e.Return(YaccParse)
																												return
																											}

																										}, 1)

																										tmp13123 := MakeNative(func(__e *ControlFlow) {
																											Parse__shen_4_5whitespaces_6 := __e.Get(1)
																											_ = Parse__shen_4_5whitespaces_6
																											tmp13133 := Call(__e, PrimNS2Value(symfail))

																											tmp13134 := PrimEqual(tmp13133, Parse__shen_4_5whitespaces_6)

																											tmp13135 := PrimNot(tmp13134)

																											if True == tmp13135 {
																												tmp13125 := MakeNative(func(__e *ControlFlow) {
																													Parse__shen_4_5st__input_6 := __e.Get(1)
																													_ = Parse__shen_4_5st__input_6
																													tmp13129 := Call(__e, PrimNS2Value(symfail))

																													tmp13130 := PrimEqual(tmp13129, Parse__shen_4_5st__input_6)

																													tmp13131 := PrimNot(tmp13130)

																													if True == tmp13131 {
																														tmp13127 := PrimHead(Parse__shen_4_5st__input_6)

																														tmp13128 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																														__e.TailApply(PrimNS2Value(symshen_4pair), tmp13127, tmp13128)
																														return

																													} else {
																														__e.TailApply(PrimNS2Value(symfail))
																														return
																													}

																												}, 1)

																												tmp13132 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5whitespaces_6)

																												__e.TailApply(tmp13125, tmp13132)
																												return

																											} else {
																												__e.TailApply(PrimNS2Value(symfail))
																												return
																											}

																										}, 1)

																										tmp13136 := Call(__e, PrimNS2Value(symshen_4_5whitespaces_6), V1515)

																										tmp13137 := Call(__e, tmp13123, tmp13136)

																										__e.TailApply(tmp13112, tmp13137)
																										return

																									} else {
																										__e.Return(YaccParse)
																										return
																									}

																								}, 1)

																								tmp13140 := MakeNative(func(__e *ControlFlow) {
																									Parse__shen_4_5atom_6 := __e.Get(1)
																									_ = Parse__shen_4_5atom_6
																									tmp13153 := Call(__e, PrimNS2Value(symfail))

																									tmp13154 := PrimEqual(tmp13153, Parse__shen_4_5atom_6)

																									tmp13155 := PrimNot(tmp13154)

																									if True == tmp13155 {
																										tmp13142 := MakeNative(func(__e *ControlFlow) {
																											Parse__shen_4_5st__input_6 := __e.Get(1)
																											_ = Parse__shen_4_5st__input_6
																											tmp13149 := Call(__e, PrimNS2Value(symfail))

																											tmp13150 := PrimEqual(tmp13149, Parse__shen_4_5st__input_6)

																											tmp13151 := PrimNot(tmp13150)

																											if True == tmp13151 {
																												tmp13144 := PrimHead(Parse__shen_4_5st__input_6)

																												tmp13145 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5atom_6)

																												tmp13146 := Call(__e, PrimNS2Value(symmacroexpand), tmp13145)

																												tmp13147 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																												tmp13148 := PrimCons(tmp13146, tmp13147)

																												__e.TailApply(PrimNS2Value(symshen_4pair), tmp13144, tmp13148)
																												return

																											} else {
																												__e.TailApply(PrimNS2Value(symfail))
																												return
																											}

																										}, 1)

																										tmp13152 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5atom_6)

																										__e.TailApply(tmp13142, tmp13152)
																										return

																									} else {
																										__e.TailApply(PrimNS2Value(symfail))
																										return
																									}

																								}, 1)

																								tmp13156 := Call(__e, PrimNS2Value(symshen_4_5atom_6), V1515)

																								tmp13157 := Call(__e, tmp13140, tmp13156)

																								__e.TailApply(tmp13110, tmp13157)
																								return

																							} else {
																								__e.Return(YaccParse)
																								return
																							}

																						}, 1)

																						tmp13160 := MakeNative(func(__e *ControlFlow) {
																							Parse__shen_4_5comment_6 := __e.Get(1)
																							_ = Parse__shen_4_5comment_6
																							tmp13170 := Call(__e, PrimNS2Value(symfail))

																							tmp13171 := PrimEqual(tmp13170, Parse__shen_4_5comment_6)

																							tmp13172 := PrimNot(tmp13171)

																							if True == tmp13172 {
																								tmp13162 := MakeNative(func(__e *ControlFlow) {
																									Parse__shen_4_5st__input_6 := __e.Get(1)
																									_ = Parse__shen_4_5st__input_6
																									tmp13166 := Call(__e, PrimNS2Value(symfail))

																									tmp13167 := PrimEqual(tmp13166, Parse__shen_4_5st__input_6)

																									tmp13168 := PrimNot(tmp13167)

																									if True == tmp13168 {
																										tmp13164 := PrimHead(Parse__shen_4_5st__input_6)

																										tmp13165 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																										__e.TailApply(PrimNS2Value(symshen_4pair), tmp13164, tmp13165)
																										return

																									} else {
																										__e.TailApply(PrimNS2Value(symfail))
																										return
																									}

																								}, 1)

																								tmp13169 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5comment_6)

																								__e.TailApply(tmp13162, tmp13169)
																								return

																							} else {
																								__e.TailApply(PrimNS2Value(symfail))
																								return
																							}

																						}, 1)

																						tmp13173 := Call(__e, PrimNS2Value(symshen_4_5comment_6), V1515)

																						tmp13174 := Call(__e, tmp13160, tmp13173)

																						__e.TailApply(tmp13108, tmp13174)
																						return

																					} else {
																						__e.Return(YaccParse)
																						return
																					}

																				}, 1)

																				tmp13177 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5comma_6 := __e.Get(1)
																					_ = Parse__shen_4_5comma_6
																					tmp13189 := Call(__e, PrimNS2Value(symfail))

																					tmp13190 := PrimEqual(tmp13189, Parse__shen_4_5comma_6)

																					tmp13191 := PrimNot(tmp13190)

																					if True == tmp13191 {
																						tmp13179 := MakeNative(func(__e *ControlFlow) {
																							Parse__shen_4_5st__input_6 := __e.Get(1)
																							_ = Parse__shen_4_5st__input_6
																							tmp13185 := Call(__e, PrimNS2Value(symfail))

																							tmp13186 := PrimEqual(tmp13185, Parse__shen_4_5st__input_6)

																							tmp13187 := PrimNot(tmp13186)

																							if True == tmp13187 {
																								tmp13181 := PrimHead(Parse__shen_4_5st__input_6)

																								tmp13182 := PrimIntern(MakeString(","))

																								tmp13183 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																								tmp13184 := PrimCons(tmp13182, tmp13183)

																								__e.TailApply(PrimNS2Value(symshen_4pair), tmp13181, tmp13184)
																								return

																							} else {
																								__e.TailApply(PrimNS2Value(symfail))
																								return
																							}

																						}, 1)

																						tmp13188 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5comma_6)

																						__e.TailApply(tmp13179, tmp13188)
																						return

																					} else {
																						__e.TailApply(PrimNS2Value(symfail))
																						return
																					}

																				}, 1)

																				tmp13192 := Call(__e, PrimNS2Value(symshen_4_5comma_6), V1515)

																				tmp13193 := Call(__e, tmp13177, tmp13192)

																				__e.TailApply(tmp13106, tmp13193)
																				return

																			} else {
																				__e.Return(YaccParse)
																				return
																			}

																		}, 1)

																		tmp13196 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5colon_6 := __e.Get(1)
																			_ = Parse__shen_4_5colon_6
																			tmp13207 := Call(__e, PrimNS2Value(symfail))

																			tmp13208 := PrimEqual(tmp13207, Parse__shen_4_5colon_6)

																			tmp13209 := PrimNot(tmp13208)

																			if True == tmp13209 {
																				tmp13198 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5st__input_6 := __e.Get(1)
																					_ = Parse__shen_4_5st__input_6
																					tmp13203 := Call(__e, PrimNS2Value(symfail))

																					tmp13204 := PrimEqual(tmp13203, Parse__shen_4_5st__input_6)

																					tmp13205 := PrimNot(tmp13204)

																					if True == tmp13205 {
																						tmp13200 := PrimHead(Parse__shen_4_5st__input_6)

																						tmp13201 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																						tmp13202 := PrimCons(sym_h, tmp13201)

																						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13200, tmp13202)
																						return

																					} else {
																						__e.TailApply(PrimNS2Value(symfail))
																						return
																					}

																				}, 1)

																				tmp13206 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5colon_6)

																				__e.TailApply(tmp13198, tmp13206)
																				return

																			} else {
																				__e.TailApply(PrimNS2Value(symfail))
																				return
																			}

																		}, 1)

																		tmp13210 := Call(__e, PrimNS2Value(symshen_4_5colon_6), V1515)

																		tmp13211 := Call(__e, tmp13196, tmp13210)

																		__e.TailApply(tmp13104, tmp13211)
																		return

																	} else {
																		__e.Return(YaccParse)
																		return
																	}

																}, 1)

																tmp13214 := MakeNative(func(__e *ControlFlow) {
																	Parse__shen_4_5colon_6 := __e.Get(1)
																	_ = Parse__shen_4_5colon_6
																	tmp13231 := Call(__e, PrimNS2Value(symfail))

																	tmp13232 := PrimEqual(tmp13231, Parse__shen_4_5colon_6)

																	tmp13233 := PrimNot(tmp13232)

																	if True == tmp13233 {
																		tmp13216 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5minus_6 := __e.Get(1)
																			_ = Parse__shen_4_5minus_6
																			tmp13227 := Call(__e, PrimNS2Value(symfail))

																			tmp13228 := PrimEqual(tmp13227, Parse__shen_4_5minus_6)

																			tmp13229 := PrimNot(tmp13228)

																			if True == tmp13229 {
																				tmp13218 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5st__input_6 := __e.Get(1)
																					_ = Parse__shen_4_5st__input_6
																					tmp13223 := Call(__e, PrimNS2Value(symfail))

																					tmp13224 := PrimEqual(tmp13223, Parse__shen_4_5st__input_6)

																					tmp13225 := PrimNot(tmp13224)

																					if True == tmp13225 {
																						tmp13220 := PrimHead(Parse__shen_4_5st__input_6)

																						tmp13221 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																						tmp13222 := PrimCons(sym_h_1, tmp13221)

																						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13220, tmp13222)
																						return

																					} else {
																						__e.TailApply(PrimNS2Value(symfail))
																						return
																					}

																				}, 1)

																				tmp13226 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5minus_6)

																				__e.TailApply(tmp13218, tmp13226)
																				return

																			} else {
																				__e.TailApply(PrimNS2Value(symfail))
																				return
																			}

																		}, 1)

																		tmp13230 := Call(__e, PrimNS2Value(symshen_4_5minus_6), Parse__shen_4_5colon_6)

																		__e.TailApply(tmp13216, tmp13230)
																		return

																	} else {
																		__e.TailApply(PrimNS2Value(symfail))
																		return
																	}

																}, 1)

																tmp13234 := Call(__e, PrimNS2Value(symshen_4_5colon_6), V1515)

																tmp13235 := Call(__e, tmp13214, tmp13234)

																__e.TailApply(tmp13102, tmp13235)
																return

															} else {
																__e.Return(YaccParse)
																return
															}

														}, 1)

														tmp13238 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5colon_6 := __e.Get(1)
															_ = Parse__shen_4_5colon_6
															tmp13255 := Call(__e, PrimNS2Value(symfail))

															tmp13256 := PrimEqual(tmp13255, Parse__shen_4_5colon_6)

															tmp13257 := PrimNot(tmp13256)

															if True == tmp13257 {
																tmp13240 := MakeNative(func(__e *ControlFlow) {
																	Parse__shen_4_5equal_6 := __e.Get(1)
																	_ = Parse__shen_4_5equal_6
																	tmp13251 := Call(__e, PrimNS2Value(symfail))

																	tmp13252 := PrimEqual(tmp13251, Parse__shen_4_5equal_6)

																	tmp13253 := PrimNot(tmp13252)

																	if True == tmp13253 {
																		tmp13242 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5st__input_6 := __e.Get(1)
																			_ = Parse__shen_4_5st__input_6
																			tmp13247 := Call(__e, PrimNS2Value(symfail))

																			tmp13248 := PrimEqual(tmp13247, Parse__shen_4_5st__input_6)

																			tmp13249 := PrimNot(tmp13248)

																			if True == tmp13249 {
																				tmp13244 := PrimHead(Parse__shen_4_5st__input_6)

																				tmp13245 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																				tmp13246 := PrimCons(sym_h_a, tmp13245)

																				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13244, tmp13246)
																				return

																			} else {
																				__e.TailApply(PrimNS2Value(symfail))
																				return
																			}

																		}, 1)

																		tmp13250 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5equal_6)

																		__e.TailApply(tmp13242, tmp13250)
																		return

																	} else {
																		__e.TailApply(PrimNS2Value(symfail))
																		return
																	}

																}, 1)

																tmp13254 := Call(__e, PrimNS2Value(symshen_4_5equal_6), Parse__shen_4_5colon_6)

																__e.TailApply(tmp13240, tmp13254)
																return

															} else {
																__e.TailApply(PrimNS2Value(symfail))
																return
															}

														}, 1)

														tmp13258 := Call(__e, PrimNS2Value(symshen_4_5colon_6), V1515)

														tmp13259 := Call(__e, tmp13238, tmp13258)

														__e.TailApply(tmp13100, tmp13259)
														return

													} else {
														__e.Return(YaccParse)
														return
													}

												}, 1)

												tmp13262 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5semicolon_6 := __e.Get(1)
													_ = Parse__shen_4_5semicolon_6
													tmp13273 := Call(__e, PrimNS2Value(symfail))

													tmp13274 := PrimEqual(tmp13273, Parse__shen_4_5semicolon_6)

													tmp13275 := PrimNot(tmp13274)

													if True == tmp13275 {
														tmp13264 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5st__input_6 := __e.Get(1)
															_ = Parse__shen_4_5st__input_6
															tmp13269 := Call(__e, PrimNS2Value(symfail))

															tmp13270 := PrimEqual(tmp13269, Parse__shen_4_5st__input_6)

															tmp13271 := PrimNot(tmp13270)

															if True == tmp13271 {
																tmp13266 := PrimHead(Parse__shen_4_5st__input_6)

																tmp13267 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

																tmp13268 := PrimCons(sym_k, tmp13267)

																__e.TailApply(PrimNS2Value(symshen_4pair), tmp13266, tmp13268)
																return

															} else {
																__e.TailApply(PrimNS2Value(symfail))
																return
															}

														}, 1)

														tmp13272 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5semicolon_6)

														__e.TailApply(tmp13264, tmp13272)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp13276 := Call(__e, PrimNS2Value(symshen_4_5semicolon_6), V1515)

												tmp13277 := Call(__e, tmp13262, tmp13276)

												__e.TailApply(tmp13098, tmp13277)
												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										tmp13280 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5bar_6 := __e.Get(1)
											_ = Parse__shen_4_5bar_6
											tmp13291 := Call(__e, PrimNS2Value(symfail))

											tmp13292 := PrimEqual(tmp13291, Parse__shen_4_5bar_6)

											tmp13293 := PrimNot(tmp13292)

											if True == tmp13293 {
												tmp13282 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5st__input_6 := __e.Get(1)
													_ = Parse__shen_4_5st__input_6
													tmp13287 := Call(__e, PrimNS2Value(symfail))

													tmp13288 := PrimEqual(tmp13287, Parse__shen_4_5st__input_6)

													tmp13289 := PrimNot(tmp13288)

													if True == tmp13289 {
														tmp13284 := PrimHead(Parse__shen_4_5st__input_6)

														tmp13285 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

														tmp13286 := PrimCons(symbar_b, tmp13285)

														__e.TailApply(PrimNS2Value(symshen_4pair), tmp13284, tmp13286)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp13290 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5bar_6)

												__e.TailApply(tmp13282, tmp13290)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp13294 := Call(__e, PrimNS2Value(symshen_4_5bar_6), V1515)

										tmp13295 := Call(__e, tmp13280, tmp13294)

										__e.TailApply(tmp13096, tmp13295)
										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								tmp13298 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5rcurly_6 := __e.Get(1)
									_ = Parse__shen_4_5rcurly_6
									tmp13309 := Call(__e, PrimNS2Value(symfail))

									tmp13310 := PrimEqual(tmp13309, Parse__shen_4_5rcurly_6)

									tmp13311 := PrimNot(tmp13310)

									if True == tmp13311 {
										tmp13300 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5st__input_6 := __e.Get(1)
											_ = Parse__shen_4_5st__input_6
											tmp13305 := Call(__e, PrimNS2Value(symfail))

											tmp13306 := PrimEqual(tmp13305, Parse__shen_4_5st__input_6)

											tmp13307 := PrimNot(tmp13306)

											if True == tmp13307 {
												tmp13302 := PrimHead(Parse__shen_4_5st__input_6)

												tmp13303 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

												tmp13304 := PrimCons(sym_j, tmp13303)

												__e.TailApply(PrimNS2Value(symshen_4pair), tmp13302, tmp13304)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp13308 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5rcurly_6)

										__e.TailApply(tmp13300, tmp13308)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp13312 := Call(__e, PrimNS2Value(symshen_4_5rcurly_6), V1515)

								tmp13313 := Call(__e, tmp13298, tmp13312)

								__e.TailApply(tmp13094, tmp13313)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp13316 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5lcurly_6 := __e.Get(1)
							_ = Parse__shen_4_5lcurly_6
							tmp13327 := Call(__e, PrimNS2Value(symfail))

							tmp13328 := PrimEqual(tmp13327, Parse__shen_4_5lcurly_6)

							tmp13329 := PrimNot(tmp13328)

							if True == tmp13329 {
								tmp13318 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5st__input_6 := __e.Get(1)
									_ = Parse__shen_4_5st__input_6
									tmp13323 := Call(__e, PrimNS2Value(symfail))

									tmp13324 := PrimEqual(tmp13323, Parse__shen_4_5st__input_6)

									tmp13325 := PrimNot(tmp13324)

									if True == tmp13325 {
										tmp13320 := PrimHead(Parse__shen_4_5st__input_6)

										tmp13321 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

										tmp13322 := PrimCons(sym_i, tmp13321)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp13320, tmp13322)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp13326 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), Parse__shen_4_5lcurly_6)

								__e.TailApply(tmp13318, tmp13326)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp13330 := Call(__e, PrimNS2Value(symshen_4_5lcurly_6), V1515)

						tmp13331 := Call(__e, tmp13316, tmp13330)

						__e.TailApply(tmp13092, tmp13331)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp13334 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5lrb_6 := __e.Get(1)
					_ = Parse__shen_4_5lrb_6
					tmp13359 := Call(__e, PrimNS2Value(symfail))

					tmp13360 := PrimEqual(tmp13359, Parse__shen_4_5lrb_6)

					tmp13361 := PrimNot(tmp13360)

					if True == tmp13361 {
						tmp13336 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5st__input1_6 := __e.Get(1)
							_ = Parse__shen_4_5st__input1_6
							tmp13355 := Call(__e, PrimNS2Value(symfail))

							tmp13356 := PrimEqual(tmp13355, Parse__shen_4_5st__input1_6)

							tmp13357 := PrimNot(tmp13356)

							if True == tmp13357 {
								tmp13338 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5rrb_6 := __e.Get(1)
									_ = Parse__shen_4_5rrb_6
									tmp13351 := Call(__e, PrimNS2Value(symfail))

									tmp13352 := PrimEqual(tmp13351, Parse__shen_4_5rrb_6)

									tmp13353 := PrimNot(tmp13352)

									if True == tmp13353 {
										tmp13340 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5st__input2_6 := __e.Get(1)
											_ = Parse__shen_4_5st__input2_6
											tmp13347 := Call(__e, PrimNS2Value(symfail))

											tmp13348 := PrimEqual(tmp13347, Parse__shen_4_5st__input2_6)

											tmp13349 := PrimNot(tmp13348)

											if True == tmp13349 {
												tmp13342 := PrimHead(Parse__shen_4_5st__input2_6)

												tmp13343 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input1_6)

												tmp13344 := Call(__e, PrimNS2Value(symmacroexpand), tmp13343)

												tmp13345 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input2_6)

												tmp13346 := Call(__e, PrimNS2Value(symshen_4package_1macro), tmp13344, tmp13345)

												__e.TailApply(PrimNS2Value(symshen_4pair), tmp13342, tmp13346)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp13350 := Call(__e, PrimNS2Value(symshen_4_5st__input2_6), Parse__shen_4_5rrb_6)

										__e.TailApply(tmp13340, tmp13350)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp13354 := Call(__e, PrimNS2Value(symshen_4_5rrb_6), Parse__shen_4_5st__input1_6)

								__e.TailApply(tmp13338, tmp13354)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp13358 := Call(__e, PrimNS2Value(symshen_4_5st__input1_6), Parse__shen_4_5lrb_6)

						__e.TailApply(tmp13336, tmp13358)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13362 := Call(__e, PrimNS2Value(symshen_4_5lrb_6), V1515)

				tmp13363 := Call(__e, tmp13334, tmp13362)

				__e.TailApply(tmp13090, tmp13363)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp13366 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5lsb_6 := __e.Get(1)
			_ = Parse__shen_4_5lsb_6
			tmp13392 := Call(__e, PrimNS2Value(symfail))

			tmp13393 := PrimEqual(tmp13392, Parse__shen_4_5lsb_6)

			tmp13394 := PrimNot(tmp13393)

			if True == tmp13394 {
				tmp13368 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5st__input1_6 := __e.Get(1)
					_ = Parse__shen_4_5st__input1_6
					tmp13388 := Call(__e, PrimNS2Value(symfail))

					tmp13389 := PrimEqual(tmp13388, Parse__shen_4_5st__input1_6)

					tmp13390 := PrimNot(tmp13389)

					if True == tmp13390 {
						tmp13370 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rsb_6 := __e.Get(1)
							_ = Parse__shen_4_5rsb_6
							tmp13384 := Call(__e, PrimNS2Value(symfail))

							tmp13385 := PrimEqual(tmp13384, Parse__shen_4_5rsb_6)

							tmp13386 := PrimNot(tmp13385)

							if True == tmp13386 {
								tmp13372 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5st__input2_6 := __e.Get(1)
									_ = Parse__shen_4_5st__input2_6
									tmp13380 := Call(__e, PrimNS2Value(symfail))

									tmp13381 := PrimEqual(tmp13380, Parse__shen_4_5st__input2_6)

									tmp13382 := PrimNot(tmp13381)

									if True == tmp13382 {
										tmp13374 := PrimHead(Parse__shen_4_5st__input2_6)

										tmp13375 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input1_6)

										tmp13376 := Call(__e, PrimNS2Value(symshen_4cons__form), tmp13375)

										tmp13377 := Call(__e, PrimNS2Value(symmacroexpand), tmp13376)

										tmp13378 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input2_6)

										tmp13379 := PrimCons(tmp13377, tmp13378)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp13374, tmp13379)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp13383 := Call(__e, PrimNS2Value(symshen_4_5st__input2_6), Parse__shen_4_5rsb_6)

								__e.TailApply(tmp13372, tmp13383)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp13387 := Call(__e, PrimNS2Value(symshen_4_5rsb_6), Parse__shen_4_5st__input1_6)

						__e.TailApply(tmp13370, tmp13387)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13391 := Call(__e, PrimNS2Value(symshen_4_5st__input1_6), Parse__shen_4_5lsb_6)

				__e.TailApply(tmp13368, tmp13391)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13395 := Call(__e, PrimNS2Value(symshen_4_5lsb_6), V1515)

		tmp13396 := Call(__e, tmp13366, tmp13395)

		__e.TailApply(tmp13088, tmp13396)
		return

	}, 1)

	tmp13397 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input_6, tmp13087)

	_ = tmp13397

	tmp13398 := MakeNative(func(__e *ControlFlow) {
		V1518 := __e.Get(1)
		_ = V1518
		tmp13409 := PrimHead(V1518)

		tmp13410 := PrimIsPair(tmp13409)

		var ifres13405 Obj

		if True == tmp13410 {
			tmp13407 := Call(__e, PrimNS2Value(symshen_4hdhd), V1518)

			tmp13408 := PrimEqual(MakeNumber(91), tmp13407)

			var ifres13406 Obj

			if True == tmp13408 {
				ifres13406 = True

			} else {
				ifres13406 = False

			}

			ifres13405 = ifres13406

		} else {
			ifres13405 = False

		}

		if True == ifres13405 {
			tmp13400 := MakeNative(func(__e *ControlFlow) {
				NewStream1516 := __e.Get(1)
				_ = NewStream1516
				tmp13401 := PrimHead(NewStream1516)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13401, symshen_4skip)
				return

			}, 1)

			tmp13402 := Call(__e, PrimNS2Value(symshen_4tlhd), V1518)

			tmp13403 := Call(__e, PrimNS2Value(symshen_4hdtl), V1518)

			tmp13404 := Call(__e, PrimNS2Value(symshen_4pair), tmp13402, tmp13403)

			__e.TailApply(tmp13400, tmp13404)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13411 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lsb_6, tmp13398)

	_ = tmp13411

	tmp13412 := MakeNative(func(__e *ControlFlow) {
		V1521 := __e.Get(1)
		_ = V1521
		tmp13423 := PrimHead(V1521)

		tmp13424 := PrimIsPair(tmp13423)

		var ifres13419 Obj

		if True == tmp13424 {
			tmp13421 := Call(__e, PrimNS2Value(symshen_4hdhd), V1521)

			tmp13422 := PrimEqual(MakeNumber(93), tmp13421)

			var ifres13420 Obj

			if True == tmp13422 {
				ifres13420 = True

			} else {
				ifres13420 = False

			}

			ifres13419 = ifres13420

		} else {
			ifres13419 = False

		}

		if True == ifres13419 {
			tmp13414 := MakeNative(func(__e *ControlFlow) {
				NewStream1519 := __e.Get(1)
				_ = NewStream1519
				tmp13415 := PrimHead(NewStream1519)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13415, symshen_4skip)
				return

			}, 1)

			tmp13416 := Call(__e, PrimNS2Value(symshen_4tlhd), V1521)

			tmp13417 := Call(__e, PrimNS2Value(symshen_4hdtl), V1521)

			tmp13418 := Call(__e, PrimNS2Value(symshen_4pair), tmp13416, tmp13417)

			__e.TailApply(tmp13414, tmp13418)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13425 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rsb_6, tmp13412)

	_ = tmp13425

	tmp13426 := MakeNative(func(__e *ControlFlow) {
		V1524 := __e.Get(1)
		_ = V1524
		tmp13437 := PrimHead(V1524)

		tmp13438 := PrimIsPair(tmp13437)

		var ifres13433 Obj

		if True == tmp13438 {
			tmp13435 := Call(__e, PrimNS2Value(symshen_4hdhd), V1524)

			tmp13436 := PrimEqual(MakeNumber(123), tmp13435)

			var ifres13434 Obj

			if True == tmp13436 {
				ifres13434 = True

			} else {
				ifres13434 = False

			}

			ifres13433 = ifres13434

		} else {
			ifres13433 = False

		}

		if True == ifres13433 {
			tmp13428 := MakeNative(func(__e *ControlFlow) {
				NewStream1522 := __e.Get(1)
				_ = NewStream1522
				tmp13429 := PrimHead(NewStream1522)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13429, symshen_4skip)
				return

			}, 1)

			tmp13430 := Call(__e, PrimNS2Value(symshen_4tlhd), V1524)

			tmp13431 := Call(__e, PrimNS2Value(symshen_4hdtl), V1524)

			tmp13432 := Call(__e, PrimNS2Value(symshen_4pair), tmp13430, tmp13431)

			__e.TailApply(tmp13428, tmp13432)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13439 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lcurly_6, tmp13426)

	_ = tmp13439

	tmp13440 := MakeNative(func(__e *ControlFlow) {
		V1527 := __e.Get(1)
		_ = V1527
		tmp13451 := PrimHead(V1527)

		tmp13452 := PrimIsPair(tmp13451)

		var ifres13447 Obj

		if True == tmp13452 {
			tmp13449 := Call(__e, PrimNS2Value(symshen_4hdhd), V1527)

			tmp13450 := PrimEqual(MakeNumber(125), tmp13449)

			var ifres13448 Obj

			if True == tmp13450 {
				ifres13448 = True

			} else {
				ifres13448 = False

			}

			ifres13447 = ifres13448

		} else {
			ifres13447 = False

		}

		if True == ifres13447 {
			tmp13442 := MakeNative(func(__e *ControlFlow) {
				NewStream1525 := __e.Get(1)
				_ = NewStream1525
				tmp13443 := PrimHead(NewStream1525)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13443, symshen_4skip)
				return

			}, 1)

			tmp13444 := Call(__e, PrimNS2Value(symshen_4tlhd), V1527)

			tmp13445 := Call(__e, PrimNS2Value(symshen_4hdtl), V1527)

			tmp13446 := Call(__e, PrimNS2Value(symshen_4pair), tmp13444, tmp13445)

			__e.TailApply(tmp13442, tmp13446)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13453 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rcurly_6, tmp13440)

	_ = tmp13453

	tmp13454 := MakeNative(func(__e *ControlFlow) {
		V1530 := __e.Get(1)
		_ = V1530
		tmp13465 := PrimHead(V1530)

		tmp13466 := PrimIsPair(tmp13465)

		var ifres13461 Obj

		if True == tmp13466 {
			tmp13463 := Call(__e, PrimNS2Value(symshen_4hdhd), V1530)

			tmp13464 := PrimEqual(MakeNumber(124), tmp13463)

			var ifres13462 Obj

			if True == tmp13464 {
				ifres13462 = True

			} else {
				ifres13462 = False

			}

			ifres13461 = ifres13462

		} else {
			ifres13461 = False

		}

		if True == ifres13461 {
			tmp13456 := MakeNative(func(__e *ControlFlow) {
				NewStream1528 := __e.Get(1)
				_ = NewStream1528
				tmp13457 := PrimHead(NewStream1528)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13457, symshen_4skip)
				return

			}, 1)

			tmp13458 := Call(__e, PrimNS2Value(symshen_4tlhd), V1530)

			tmp13459 := Call(__e, PrimNS2Value(symshen_4hdtl), V1530)

			tmp13460 := Call(__e, PrimNS2Value(symshen_4pair), tmp13458, tmp13459)

			__e.TailApply(tmp13456, tmp13460)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13467 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5bar_6, tmp13454)

	_ = tmp13467

	tmp13468 := MakeNative(func(__e *ControlFlow) {
		V1533 := __e.Get(1)
		_ = V1533
		tmp13479 := PrimHead(V1533)

		tmp13480 := PrimIsPair(tmp13479)

		var ifres13475 Obj

		if True == tmp13480 {
			tmp13477 := Call(__e, PrimNS2Value(symshen_4hdhd), V1533)

			tmp13478 := PrimEqual(MakeNumber(59), tmp13477)

			var ifres13476 Obj

			if True == tmp13478 {
				ifres13476 = True

			} else {
				ifres13476 = False

			}

			ifres13475 = ifres13476

		} else {
			ifres13475 = False

		}

		if True == ifres13475 {
			tmp13470 := MakeNative(func(__e *ControlFlow) {
				NewStream1531 := __e.Get(1)
				_ = NewStream1531
				tmp13471 := PrimHead(NewStream1531)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13471, symshen_4skip)
				return

			}, 1)

			tmp13472 := Call(__e, PrimNS2Value(symshen_4tlhd), V1533)

			tmp13473 := Call(__e, PrimNS2Value(symshen_4hdtl), V1533)

			tmp13474 := Call(__e, PrimNS2Value(symshen_4pair), tmp13472, tmp13473)

			__e.TailApply(tmp13470, tmp13474)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13481 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5semicolon_6, tmp13468)

	_ = tmp13481

	tmp13482 := MakeNative(func(__e *ControlFlow) {
		V1536 := __e.Get(1)
		_ = V1536
		tmp13493 := PrimHead(V1536)

		tmp13494 := PrimIsPair(tmp13493)

		var ifres13489 Obj

		if True == tmp13494 {
			tmp13491 := Call(__e, PrimNS2Value(symshen_4hdhd), V1536)

			tmp13492 := PrimEqual(MakeNumber(58), tmp13491)

			var ifres13490 Obj

			if True == tmp13492 {
				ifres13490 = True

			} else {
				ifres13490 = False

			}

			ifres13489 = ifres13490

		} else {
			ifres13489 = False

		}

		if True == ifres13489 {
			tmp13484 := MakeNative(func(__e *ControlFlow) {
				NewStream1534 := __e.Get(1)
				_ = NewStream1534
				tmp13485 := PrimHead(NewStream1534)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13485, symshen_4skip)
				return

			}, 1)

			tmp13486 := Call(__e, PrimNS2Value(symshen_4tlhd), V1536)

			tmp13487 := Call(__e, PrimNS2Value(symshen_4hdtl), V1536)

			tmp13488 := Call(__e, PrimNS2Value(symshen_4pair), tmp13486, tmp13487)

			__e.TailApply(tmp13484, tmp13488)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13495 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5colon_6, tmp13482)

	_ = tmp13495

	tmp13496 := MakeNative(func(__e *ControlFlow) {
		V1539 := __e.Get(1)
		_ = V1539
		tmp13507 := PrimHead(V1539)

		tmp13508 := PrimIsPair(tmp13507)

		var ifres13503 Obj

		if True == tmp13508 {
			tmp13505 := Call(__e, PrimNS2Value(symshen_4hdhd), V1539)

			tmp13506 := PrimEqual(MakeNumber(44), tmp13505)

			var ifres13504 Obj

			if True == tmp13506 {
				ifres13504 = True

			} else {
				ifres13504 = False

			}

			ifres13503 = ifres13504

		} else {
			ifres13503 = False

		}

		if True == ifres13503 {
			tmp13498 := MakeNative(func(__e *ControlFlow) {
				NewStream1537 := __e.Get(1)
				_ = NewStream1537
				tmp13499 := PrimHead(NewStream1537)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13499, symshen_4skip)
				return

			}, 1)

			tmp13500 := Call(__e, PrimNS2Value(symshen_4tlhd), V1539)

			tmp13501 := Call(__e, PrimNS2Value(symshen_4hdtl), V1539)

			tmp13502 := Call(__e, PrimNS2Value(symshen_4pair), tmp13500, tmp13501)

			__e.TailApply(tmp13498, tmp13502)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13509 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comma_6, tmp13496)

	_ = tmp13509

	tmp13510 := MakeNative(func(__e *ControlFlow) {
		V1542 := __e.Get(1)
		_ = V1542
		tmp13521 := PrimHead(V1542)

		tmp13522 := PrimIsPair(tmp13521)

		var ifres13517 Obj

		if True == tmp13522 {
			tmp13519 := Call(__e, PrimNS2Value(symshen_4hdhd), V1542)

			tmp13520 := PrimEqual(MakeNumber(61), tmp13519)

			var ifres13518 Obj

			if True == tmp13520 {
				ifres13518 = True

			} else {
				ifres13518 = False

			}

			ifres13517 = ifres13518

		} else {
			ifres13517 = False

		}

		if True == ifres13517 {
			tmp13512 := MakeNative(func(__e *ControlFlow) {
				NewStream1540 := __e.Get(1)
				_ = NewStream1540
				tmp13513 := PrimHead(NewStream1540)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13513, symshen_4skip)
				return

			}, 1)

			tmp13514 := Call(__e, PrimNS2Value(symshen_4tlhd), V1542)

			tmp13515 := Call(__e, PrimNS2Value(symshen_4hdtl), V1542)

			tmp13516 := Call(__e, PrimNS2Value(symshen_4pair), tmp13514, tmp13515)

			__e.TailApply(tmp13512, tmp13516)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13523 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5equal_6, tmp13510)

	_ = tmp13523

	tmp13524 := MakeNative(func(__e *ControlFlow) {
		V1545 := __e.Get(1)
		_ = V1545
		tmp13535 := PrimHead(V1545)

		tmp13536 := PrimIsPair(tmp13535)

		var ifres13531 Obj

		if True == tmp13536 {
			tmp13533 := Call(__e, PrimNS2Value(symshen_4hdhd), V1545)

			tmp13534 := PrimEqual(MakeNumber(45), tmp13533)

			var ifres13532 Obj

			if True == tmp13534 {
				ifres13532 = True

			} else {
				ifres13532 = False

			}

			ifres13531 = ifres13532

		} else {
			ifres13531 = False

		}

		if True == ifres13531 {
			tmp13526 := MakeNative(func(__e *ControlFlow) {
				NewStream1543 := __e.Get(1)
				_ = NewStream1543
				tmp13527 := PrimHead(NewStream1543)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13527, symshen_4skip)
				return

			}, 1)

			tmp13528 := Call(__e, PrimNS2Value(symshen_4tlhd), V1545)

			tmp13529 := Call(__e, PrimNS2Value(symshen_4hdtl), V1545)

			tmp13530 := Call(__e, PrimNS2Value(symshen_4pair), tmp13528, tmp13529)

			__e.TailApply(tmp13526, tmp13530)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13537 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5minus_6, tmp13524)

	_ = tmp13537

	tmp13538 := MakeNative(func(__e *ControlFlow) {
		V1548 := __e.Get(1)
		_ = V1548
		tmp13549 := PrimHead(V1548)

		tmp13550 := PrimIsPair(tmp13549)

		var ifres13545 Obj

		if True == tmp13550 {
			tmp13547 := Call(__e, PrimNS2Value(symshen_4hdhd), V1548)

			tmp13548 := PrimEqual(MakeNumber(40), tmp13547)

			var ifres13546 Obj

			if True == tmp13548 {
				ifres13546 = True

			} else {
				ifres13546 = False

			}

			ifres13545 = ifres13546

		} else {
			ifres13545 = False

		}

		if True == ifres13545 {
			tmp13540 := MakeNative(func(__e *ControlFlow) {
				NewStream1546 := __e.Get(1)
				_ = NewStream1546
				tmp13541 := PrimHead(NewStream1546)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13541, symshen_4skip)
				return

			}, 1)

			tmp13542 := Call(__e, PrimNS2Value(symshen_4tlhd), V1548)

			tmp13543 := Call(__e, PrimNS2Value(symshen_4hdtl), V1548)

			tmp13544 := Call(__e, PrimNS2Value(symshen_4pair), tmp13542, tmp13543)

			__e.TailApply(tmp13540, tmp13544)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13551 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lrb_6, tmp13538)

	_ = tmp13551

	tmp13552 := MakeNative(func(__e *ControlFlow) {
		V1551 := __e.Get(1)
		_ = V1551
		tmp13563 := PrimHead(V1551)

		tmp13564 := PrimIsPair(tmp13563)

		var ifres13559 Obj

		if True == tmp13564 {
			tmp13561 := Call(__e, PrimNS2Value(symshen_4hdhd), V1551)

			tmp13562 := PrimEqual(MakeNumber(41), tmp13561)

			var ifres13560 Obj

			if True == tmp13562 {
				ifres13560 = True

			} else {
				ifres13560 = False

			}

			ifres13559 = ifres13560

		} else {
			ifres13559 = False

		}

		if True == ifres13559 {
			tmp13554 := MakeNative(func(__e *ControlFlow) {
				NewStream1549 := __e.Get(1)
				_ = NewStream1549
				tmp13555 := PrimHead(NewStream1549)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13555, symshen_4skip)
				return

			}, 1)

			tmp13556 := Call(__e, PrimNS2Value(symshen_4tlhd), V1551)

			tmp13557 := Call(__e, PrimNS2Value(symshen_4hdtl), V1551)

			tmp13558 := Call(__e, PrimNS2Value(symshen_4pair), tmp13556, tmp13557)

			__e.TailApply(tmp13554, tmp13558)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13565 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rrb_6, tmp13552)

	_ = tmp13565

	tmp13566 := MakeNative(func(__e *ControlFlow) {
		V1553 := __e.Get(1)
		_ = V1553
		tmp13567 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp13596 := Call(__e, PrimNS2Value(symfail))

			tmp13597 := PrimEqual(YaccParse, tmp13596)

			if True == tmp13597 {
				tmp13569 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp13585 := Call(__e, PrimNS2Value(symfail))

					tmp13586 := PrimEqual(YaccParse, tmp13585)

					if True == tmp13586 {
						tmp13571 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5sym_6 := __e.Get(1)
							_ = Parse__shen_4_5sym_6
							tmp13581 := Call(__e, PrimNS2Value(symfail))

							tmp13582 := PrimEqual(tmp13581, Parse__shen_4_5sym_6)

							tmp13583 := PrimNot(tmp13582)

							if True == tmp13583 {
								tmp13573 := PrimHead(Parse__shen_4_5sym_6)

								tmp13579 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5sym_6)

								tmp13580 := PrimEqual(tmp13579, MakeString("<>"))

								var ifres13574 Obj

								if True == tmp13580 {
									tmp13577 := PrimCons(MakeNumber(0), Nil)

									tmp13578 := PrimCons(symvector, tmp13577)

									ifres13574 = tmp13578

								} else {
									tmp13575 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5sym_6)

									tmp13576 := PrimIntern(tmp13575)

									ifres13574 = tmp13576

								}

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp13573, ifres13574)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp13584 := Call(__e, PrimNS2Value(symshen_4_5sym_6), V1553)

						__e.TailApply(tmp13571, tmp13584)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp13587 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5number_6 := __e.Get(1)
					_ = Parse__shen_4_5number_6
					tmp13591 := Call(__e, PrimNS2Value(symfail))

					tmp13592 := PrimEqual(tmp13591, Parse__shen_4_5number_6)

					tmp13593 := PrimNot(tmp13592)

					if True == tmp13593 {
						tmp13589 := PrimHead(Parse__shen_4_5number_6)

						tmp13590 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5number_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13589, tmp13590)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13594 := Call(__e, PrimNS2Value(symshen_4_5number_6), V1553)

				tmp13595 := Call(__e, tmp13587, tmp13594)

				__e.TailApply(tmp13569, tmp13595)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp13598 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5str_6 := __e.Get(1)
			_ = Parse__shen_4_5str_6
			tmp13603 := Call(__e, PrimNS2Value(symfail))

			tmp13604 := PrimEqual(tmp13603, Parse__shen_4_5str_6)

			tmp13605 := PrimNot(tmp13604)

			if True == tmp13605 {
				tmp13600 := PrimHead(Parse__shen_4_5str_6)

				tmp13601 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5str_6)

				tmp13602 := Call(__e, PrimNS2Value(symshen_4control_1chars), tmp13601)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13600, tmp13602)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13606 := Call(__e, PrimNS2Value(symshen_4_5str_6), V1553)

		tmp13607 := Call(__e, tmp13598, tmp13606)

		__e.TailApply(tmp13567, tmp13607)
		return

	}, 1)

	tmp13608 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5atom_6, tmp13566)

	_ = tmp13608

	tmp13609 := MakeNative(func(__e *ControlFlow) {
		V1555 := __e.Get(1)
		_ = V1555
		tmp13642 := PrimEqual(Nil, V1555)

		if True == tmp13642 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp13641 := PrimIsPair(V1555)

			var ifres13628 Obj

			if True == tmp13641 {
				tmp13639 := PrimHead(V1555)

				tmp13640 := PrimEqual(MakeString("c"), tmp13639)

				var ifres13630 Obj

				if True == tmp13640 {
					tmp13637 := PrimTail(V1555)

					tmp13638 := PrimIsPair(tmp13637)

					var ifres13632 Obj

					if True == tmp13638 {
						tmp13634 := PrimTail(V1555)

						tmp13635 := PrimHead(tmp13634)

						tmp13636 := PrimEqual(MakeString("#"), tmp13635)

						var ifres13633 Obj

						if True == tmp13636 {
							ifres13633 = True

						} else {
							ifres13633 = False

						}

						ifres13632 = ifres13633

					} else {
						ifres13632 = False

					}

					var ifres13631 Obj

					if True == ifres13632 {
						ifres13631 = True

					} else {
						ifres13631 = False

					}

					ifres13630 = ifres13631

				} else {
					ifres13630 = False

				}

				var ifres13629 Obj

				if True == ifres13630 {
					ifres13629 = True

				} else {
					ifres13629 = False

				}

				ifres13628 = ifres13629

			} else {
				ifres13628 = False

			}

			if True == ifres13628 {
				tmp13617 := MakeNative(func(__e *ControlFlow) {
					CodePoint := __e.Get(1)
					_ = CodePoint
					tmp13618 := MakeNative(func(__e *ControlFlow) {
						AfterCodePoint := __e.Get(1)
						_ = AfterCodePoint
						tmp13619 := Call(__e, PrimNS2Value(symshen_4decimalise), CodePoint)

						tmp13620 := PrimNumberToString(tmp13619)

						tmp13621 := Call(__e, PrimNS2Value(symshen_4control_1chars), AfterCodePoint)

						__e.TailApply(PrimNS2Value(sym_8s), tmp13620, tmp13621)
						return

					}, 1)

					tmp13622 := PrimTail(V1555)

					tmp13623 := PrimTail(tmp13622)

					tmp13624 := Call(__e, PrimNS2Value(symshen_4after_1codepoint), tmp13623)

					__e.TailApply(tmp13618, tmp13624)
					return

				}, 1)

				tmp13625 := PrimTail(V1555)

				tmp13626 := PrimTail(tmp13625)

				tmp13627 := Call(__e, PrimNS2Value(symshen_4code_1point), tmp13626)

				__e.TailApply(tmp13617, tmp13627)
				return

			} else {
				tmp13616 := PrimIsPair(V1555)

				if True == tmp13616 {
					tmp13613 := PrimHead(V1555)

					tmp13614 := PrimTail(V1555)

					tmp13615 := Call(__e, PrimNS2Value(symshen_4control_1chars), tmp13614)

					__e.TailApply(PrimNS2Value(sym_8s), tmp13613, tmp13615)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4control_1chars)
					return
				}

			}

		}

	}, 1)

	tmp13643 := Call(__e, PrimNS1Value(symns2_1set), symshen_4control_1chars, tmp13609)

	_ = tmp13643

	tmp13644 := MakeNative(func(__e *ControlFlow) {
		V1559 := __e.Get(1)
		_ = V1559
		tmp13672 := PrimIsPair(V1559)

		var ifres13668 Obj

		if True == tmp13672 {
			tmp13670 := PrimHead(V1559)

			tmp13671 := PrimEqual(MakeString(";"), tmp13670)

			var ifres13669 Obj

			if True == tmp13671 {
				ifres13669 = True

			} else {
				ifres13669 = False

			}

			ifres13668 = ifres13669

		} else {
			ifres13668 = False

		}

		if True == ifres13668 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp13667 := PrimIsPair(V1559)

			var ifres13652 Obj

			if True == tmp13667 {
				tmp13654 := PrimHead(V1559)

				tmp13655 := PrimCons(MakeString("0"), Nil)

				tmp13656 := PrimCons(MakeString("9"), tmp13655)

				tmp13657 := PrimCons(MakeString("8"), tmp13656)

				tmp13658 := PrimCons(MakeString("7"), tmp13657)

				tmp13659 := PrimCons(MakeString("6"), tmp13658)

				tmp13660 := PrimCons(MakeString("5"), tmp13659)

				tmp13661 := PrimCons(MakeString("4"), tmp13660)

				tmp13662 := PrimCons(MakeString("3"), tmp13661)

				tmp13663 := PrimCons(MakeString("2"), tmp13662)

				tmp13664 := PrimCons(MakeString("1"), tmp13663)

				tmp13665 := PrimCons(MakeString("0"), tmp13664)

				tmp13666 := Call(__e, PrimNS2Value(symelement_2), tmp13654, tmp13665)

				var ifres13653 Obj

				if True == tmp13666 {
					ifres13653 = True

				} else {
					ifres13653 = False

				}

				ifres13652 = ifres13653

			} else {
				ifres13652 = False

			}

			if True == ifres13652 {
				tmp13649 := PrimHead(V1559)

				tmp13650 := PrimTail(V1559)

				tmp13651 := Call(__e, PrimNS2Value(symshen_4code_1point), tmp13650)

				__e.Return(PrimCons(tmp13649, tmp13651))
				return

			} else {
				tmp13647 := Call(__e, PrimNS2Value(symshen_4app), V1559, MakeString("\n"), symshen_4a)

				tmp13648 := PrimStringConcat(MakeString("code point parse error "), tmp13647)

				__e.Return(PrimSimpleError(tmp13648))
				return

			}

		}

	}, 1)

	tmp13673 := Call(__e, PrimNS1Value(symns2_1set), symshen_4code_1point, tmp13644)

	_ = tmp13673

	tmp13674 := MakeNative(func(__e *ControlFlow) {
		V1565 := __e.Get(1)
		_ = V1565
		tmp13685 := PrimEqual(Nil, V1565)

		if True == tmp13685 {
			__e.Return(Nil)
			return
		} else {
			tmp13684 := PrimIsPair(V1565)

			var ifres13680 Obj

			if True == tmp13684 {
				tmp13682 := PrimHead(V1565)

				tmp13683 := PrimEqual(MakeString(";"), tmp13682)

				var ifres13681 Obj

				if True == tmp13683 {
					ifres13681 = True

				} else {
					ifres13681 = False

				}

				ifres13680 = ifres13681

			} else {
				ifres13680 = False

			}

			if True == ifres13680 {
				__e.Return(PrimTail(V1565))
				return
			} else {
				tmp13679 := PrimIsPair(V1565)

				if True == tmp13679 {
					tmp13678 := PrimTail(V1565)

					__e.TailApply(PrimNS2Value(symshen_4after_1codepoint), tmp13678)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4after_1codepoint)
					return
				}

			}

		}

	}, 1)

	tmp13686 := Call(__e, PrimNS1Value(symns2_1set), symshen_4after_1codepoint, tmp13674)

	_ = tmp13686

	tmp13687 := MakeNative(func(__e *ControlFlow) {
		V1567 := __e.Get(1)
		_ = V1567
		tmp13688 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), V1567)

		tmp13689 := Call(__e, PrimNS2Value(symreverse), tmp13688)

		__e.TailApply(PrimNS2Value(symshen_4pre), tmp13689, MakeNumber(0))
		return

	}, 1)

	tmp13690 := Call(__e, PrimNS1Value(symns2_1set), symshen_4decimalise, tmp13687)

	_ = tmp13690

	tmp13691 := MakeNative(func(__e *ControlFlow) {
		V1573 := __e.Get(1)
		_ = V1573
		tmp13771 := PrimIsPair(V1573)

		var ifres13767 Obj

		if True == tmp13771 {
			tmp13769 := PrimHead(V1573)

			tmp13770 := PrimEqual(MakeString("0"), tmp13769)

			var ifres13768 Obj

			if True == tmp13770 {
				ifres13768 = True

			} else {
				ifres13768 = False

			}

			ifres13767 = ifres13768

		} else {
			ifres13767 = False

		}

		if True == ifres13767 {
			tmp13765 := PrimTail(V1573)

			tmp13766 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13765)

			__e.Return(PrimCons(MakeNumber(0), tmp13766))
			return

		} else {
			tmp13764 := PrimIsPair(V1573)

			var ifres13760 Obj

			if True == tmp13764 {
				tmp13762 := PrimHead(V1573)

				tmp13763 := PrimEqual(MakeString("1"), tmp13762)

				var ifres13761 Obj

				if True == tmp13763 {
					ifres13761 = True

				} else {
					ifres13761 = False

				}

				ifres13760 = ifres13761

			} else {
				ifres13760 = False

			}

			if True == ifres13760 {
				tmp13758 := PrimTail(V1573)

				tmp13759 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13758)

				__e.Return(PrimCons(MakeNumber(1), tmp13759))
				return

			} else {
				tmp13757 := PrimIsPair(V1573)

				var ifres13753 Obj

				if True == tmp13757 {
					tmp13755 := PrimHead(V1573)

					tmp13756 := PrimEqual(MakeString("2"), tmp13755)

					var ifres13754 Obj

					if True == tmp13756 {
						ifres13754 = True

					} else {
						ifres13754 = False

					}

					ifres13753 = ifres13754

				} else {
					ifres13753 = False

				}

				if True == ifres13753 {
					tmp13751 := PrimTail(V1573)

					tmp13752 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13751)

					__e.Return(PrimCons(MakeNumber(2), tmp13752))
					return

				} else {
					tmp13750 := PrimIsPair(V1573)

					var ifres13746 Obj

					if True == tmp13750 {
						tmp13748 := PrimHead(V1573)

						tmp13749 := PrimEqual(MakeString("3"), tmp13748)

						var ifres13747 Obj

						if True == tmp13749 {
							ifres13747 = True

						} else {
							ifres13747 = False

						}

						ifres13746 = ifres13747

					} else {
						ifres13746 = False

					}

					if True == ifres13746 {
						tmp13744 := PrimTail(V1573)

						tmp13745 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13744)

						__e.Return(PrimCons(MakeNumber(3), tmp13745))
						return

					} else {
						tmp13743 := PrimIsPair(V1573)

						var ifres13739 Obj

						if True == tmp13743 {
							tmp13741 := PrimHead(V1573)

							tmp13742 := PrimEqual(MakeString("4"), tmp13741)

							var ifres13740 Obj

							if True == tmp13742 {
								ifres13740 = True

							} else {
								ifres13740 = False

							}

							ifres13739 = ifres13740

						} else {
							ifres13739 = False

						}

						if True == ifres13739 {
							tmp13737 := PrimTail(V1573)

							tmp13738 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13737)

							__e.Return(PrimCons(MakeNumber(4), tmp13738))
							return

						} else {
							tmp13736 := PrimIsPair(V1573)

							var ifres13732 Obj

							if True == tmp13736 {
								tmp13734 := PrimHead(V1573)

								tmp13735 := PrimEqual(MakeString("5"), tmp13734)

								var ifres13733 Obj

								if True == tmp13735 {
									ifres13733 = True

								} else {
									ifres13733 = False

								}

								ifres13732 = ifres13733

							} else {
								ifres13732 = False

							}

							if True == ifres13732 {
								tmp13730 := PrimTail(V1573)

								tmp13731 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13730)

								__e.Return(PrimCons(MakeNumber(5), tmp13731))
								return

							} else {
								tmp13729 := PrimIsPair(V1573)

								var ifres13725 Obj

								if True == tmp13729 {
									tmp13727 := PrimHead(V1573)

									tmp13728 := PrimEqual(MakeString("6"), tmp13727)

									var ifres13726 Obj

									if True == tmp13728 {
										ifres13726 = True

									} else {
										ifres13726 = False

									}

									ifres13725 = ifres13726

								} else {
									ifres13725 = False

								}

								if True == ifres13725 {
									tmp13723 := PrimTail(V1573)

									tmp13724 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13723)

									__e.Return(PrimCons(MakeNumber(6), tmp13724))
									return

								} else {
									tmp13722 := PrimIsPair(V1573)

									var ifres13718 Obj

									if True == tmp13722 {
										tmp13720 := PrimHead(V1573)

										tmp13721 := PrimEqual(MakeString("7"), tmp13720)

										var ifres13719 Obj

										if True == tmp13721 {
											ifres13719 = True

										} else {
											ifres13719 = False

										}

										ifres13718 = ifres13719

									} else {
										ifres13718 = False

									}

									if True == ifres13718 {
										tmp13716 := PrimTail(V1573)

										tmp13717 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13716)

										__e.Return(PrimCons(MakeNumber(7), tmp13717))
										return

									} else {
										tmp13715 := PrimIsPair(V1573)

										var ifres13711 Obj

										if True == tmp13715 {
											tmp13713 := PrimHead(V1573)

											tmp13714 := PrimEqual(MakeString("8"), tmp13713)

											var ifres13712 Obj

											if True == tmp13714 {
												ifres13712 = True

											} else {
												ifres13712 = False

											}

											ifres13711 = ifres13712

										} else {
											ifres13711 = False

										}

										if True == ifres13711 {
											tmp13709 := PrimTail(V1573)

											tmp13710 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13709)

											__e.Return(PrimCons(MakeNumber(8), tmp13710))
											return

										} else {
											tmp13708 := PrimIsPair(V1573)

											var ifres13704 Obj

											if True == tmp13708 {
												tmp13706 := PrimHead(V1573)

												tmp13707 := PrimEqual(MakeString("9"), tmp13706)

												var ifres13705 Obj

												if True == tmp13707 {
													ifres13705 = True

												} else {
													ifres13705 = False

												}

												ifres13704 = ifres13705

											} else {
												ifres13704 = False

											}

											if True == ifres13704 {
												tmp13702 := PrimTail(V1573)

												tmp13703 := Call(__e, PrimNS2Value(symshen_4digits_1_6integers), tmp13702)

												__e.Return(PrimCons(MakeNumber(9), tmp13703))
												return

											} else {
												__e.Return(Nil)
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

	}, 1)

	tmp13772 := Call(__e, PrimNS1Value(symns2_1set), symshen_4digits_1_6integers, tmp13691)

	_ = tmp13772

	tmp13773 := MakeNative(func(__e *ControlFlow) {
		V1575 := __e.Get(1)
		_ = V1575
		tmp13774 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alpha_6 := __e.Get(1)
			_ = Parse__shen_4_5alpha_6
			tmp13786 := Call(__e, PrimNS2Value(symfail))

			tmp13787 := PrimEqual(tmp13786, Parse__shen_4_5alpha_6)

			tmp13788 := PrimNot(tmp13787)

			if True == tmp13788 {
				tmp13776 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5alphanums_6 := __e.Get(1)
					_ = Parse__shen_4_5alphanums_6
					tmp13782 := Call(__e, PrimNS2Value(symfail))

					tmp13783 := PrimEqual(tmp13782, Parse__shen_4_5alphanums_6)

					tmp13784 := PrimNot(tmp13783)

					if True == tmp13784 {
						tmp13778 := PrimHead(Parse__shen_4_5alphanums_6)

						tmp13779 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5alpha_6)

						tmp13780 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5alphanums_6)

						tmp13781 := Call(__e, PrimNS2Value(sym_8s), tmp13779, tmp13780)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13778, tmp13781)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13785 := Call(__e, PrimNS2Value(symshen_4_5alphanums_6), Parse__shen_4_5alpha_6)

				__e.TailApply(tmp13776, tmp13785)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13789 := Call(__e, PrimNS2Value(symshen_4_5alpha_6), V1575)

		__e.TailApply(tmp13774, tmp13789)
		return

	}, 1)

	tmp13790 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sym_6, tmp13773)

	_ = tmp13790

	tmp13791 := MakeNative(func(__e *ControlFlow) {
		V1577 := __e.Get(1)
		_ = V1577
		tmp13792 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp13801 := Call(__e, PrimNS2Value(symfail))

			tmp13802 := PrimEqual(YaccParse, tmp13801)

			if True == tmp13802 {
				tmp13794 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp13797 := Call(__e, PrimNS2Value(symfail))

					tmp13798 := PrimEqual(tmp13797, Parse___5e_6)

					tmp13799 := PrimNot(tmp13798)

					if True == tmp13799 {
						tmp13796 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13796, MakeString(""))
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13800 := Call(__e, PrimNS2Value(sym_5e_6), V1577)

				__e.TailApply(tmp13794, tmp13800)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp13803 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alphanum_6 := __e.Get(1)
			_ = Parse__shen_4_5alphanum_6
			tmp13815 := Call(__e, PrimNS2Value(symfail))

			tmp13816 := PrimEqual(tmp13815, Parse__shen_4_5alphanum_6)

			tmp13817 := PrimNot(tmp13816)

			if True == tmp13817 {
				tmp13805 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5alphanums_6 := __e.Get(1)
					_ = Parse__shen_4_5alphanums_6
					tmp13811 := Call(__e, PrimNS2Value(symfail))

					tmp13812 := PrimEqual(tmp13811, Parse__shen_4_5alphanums_6)

					tmp13813 := PrimNot(tmp13812)

					if True == tmp13813 {
						tmp13807 := PrimHead(Parse__shen_4_5alphanums_6)

						tmp13808 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5alphanum_6)

						tmp13809 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5alphanums_6)

						tmp13810 := Call(__e, PrimNS2Value(sym_8s), tmp13808, tmp13809)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13807, tmp13810)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13814 := Call(__e, PrimNS2Value(symshen_4_5alphanums_6), Parse__shen_4_5alphanum_6)

				__e.TailApply(tmp13805, tmp13814)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13818 := Call(__e, PrimNS2Value(symshen_4_5alphanum_6), V1577)

		tmp13819 := Call(__e, tmp13803, tmp13818)

		__e.TailApply(tmp13792, tmp13819)
		return

	}, 1)

	tmp13820 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alphanums_6, tmp13791)

	_ = tmp13820

	tmp13821 := MakeNative(func(__e *ControlFlow) {
		V1579 := __e.Get(1)
		_ = V1579
		tmp13822 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp13832 := Call(__e, PrimNS2Value(symfail))

			tmp13833 := PrimEqual(YaccParse, tmp13832)

			if True == tmp13833 {
				tmp13824 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5num_6 := __e.Get(1)
					_ = Parse__shen_4_5num_6
					tmp13828 := Call(__e, PrimNS2Value(symfail))

					tmp13829 := PrimEqual(tmp13828, Parse__shen_4_5num_6)

					tmp13830 := PrimNot(tmp13829)

					if True == tmp13830 {
						tmp13826 := PrimHead(Parse__shen_4_5num_6)

						tmp13827 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5num_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13826, tmp13827)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13831 := Call(__e, PrimNS2Value(symshen_4_5num_6), V1579)

				__e.TailApply(tmp13824, tmp13831)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp13834 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alpha_6 := __e.Get(1)
			_ = Parse__shen_4_5alpha_6
			tmp13838 := Call(__e, PrimNS2Value(symfail))

			tmp13839 := PrimEqual(tmp13838, Parse__shen_4_5alpha_6)

			tmp13840 := PrimNot(tmp13839)

			if True == tmp13840 {
				tmp13836 := PrimHead(Parse__shen_4_5alpha_6)

				tmp13837 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5alpha_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13836, tmp13837)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13841 := Call(__e, PrimNS2Value(symshen_4_5alpha_6), V1579)

		tmp13842 := Call(__e, tmp13834, tmp13841)

		__e.TailApply(tmp13822, tmp13842)
		return

	}, 1)

	tmp13843 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alphanum_6, tmp13821)

	_ = tmp13843

	tmp13844 := MakeNative(func(__e *ControlFlow) {
		V1581 := __e.Get(1)
		_ = V1581
		tmp13855 := PrimHead(V1581)

		tmp13856 := PrimIsPair(tmp13855)

		if True == tmp13856 {
			tmp13846 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp13853 := Call(__e, PrimNS2Value(symshen_4numbyte_2), Parse__Char)

				if True == tmp13853 {
					tmp13848 := Call(__e, PrimNS2Value(symshen_4tlhd), V1581)

					tmp13849 := Call(__e, PrimNS2Value(symshen_4hdtl), V1581)

					tmp13850 := Call(__e, PrimNS2Value(symshen_4pair), tmp13848, tmp13849)

					tmp13851 := PrimHead(tmp13850)

					tmp13852 := PrimNumberToString(Parse__Char)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp13851, tmp13852)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp13854 := Call(__e, PrimNS2Value(symshen_4hdhd), V1581)

			__e.TailApply(tmp13846, tmp13854)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13857 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5num_6, tmp13844)

	_ = tmp13857

	tmp13858 := MakeNative(func(__e *ControlFlow) {
		V1587 := __e.Get(1)
		_ = V1587
		tmp13878 := PrimEqual(MakeNumber(48), V1587)

		if True == tmp13878 {
			__e.Return(True)
			return
		} else {
			tmp13877 := PrimEqual(MakeNumber(49), V1587)

			if True == tmp13877 {
				__e.Return(True)
				return
			} else {
				tmp13876 := PrimEqual(MakeNumber(50), V1587)

				if True == tmp13876 {
					__e.Return(True)
					return
				} else {
					tmp13875 := PrimEqual(MakeNumber(51), V1587)

					if True == tmp13875 {
						__e.Return(True)
						return
					} else {
						tmp13874 := PrimEqual(MakeNumber(52), V1587)

						if True == tmp13874 {
							__e.Return(True)
							return
						} else {
							tmp13873 := PrimEqual(MakeNumber(53), V1587)

							if True == tmp13873 {
								__e.Return(True)
								return
							} else {
								tmp13872 := PrimEqual(MakeNumber(54), V1587)

								if True == tmp13872 {
									__e.Return(True)
									return
								} else {
									tmp13871 := PrimEqual(MakeNumber(55), V1587)

									if True == tmp13871 {
										__e.Return(True)
										return
									} else {
										tmp13870 := PrimEqual(MakeNumber(56), V1587)

										if True == tmp13870 {
											__e.Return(True)
											return
										} else {
											tmp13869 := PrimEqual(MakeNumber(57), V1587)

											if True == tmp13869 {
												__e.Return(True)
												return
											} else {
												__e.Return(False)
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

	}, 1)

	tmp13879 := Call(__e, PrimNS1Value(symns2_1set), symshen_4numbyte_2, tmp13858)

	_ = tmp13879

	tmp13880 := MakeNative(func(__e *ControlFlow) {
		V1589 := __e.Get(1)
		_ = V1589
		tmp13891 := PrimHead(V1589)

		tmp13892 := PrimIsPair(tmp13891)

		if True == tmp13892 {
			tmp13882 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp13889 := Call(__e, PrimNS2Value(symshen_4symbol_1code_2), Parse__Char)

				if True == tmp13889 {
					tmp13884 := Call(__e, PrimNS2Value(symshen_4tlhd), V1589)

					tmp13885 := Call(__e, PrimNS2Value(symshen_4hdtl), V1589)

					tmp13886 := Call(__e, PrimNS2Value(symshen_4pair), tmp13884, tmp13885)

					tmp13887 := PrimHead(tmp13886)

					tmp13888 := PrimNumberToString(Parse__Char)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp13887, tmp13888)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp13890 := Call(__e, PrimNS2Value(symshen_4hdhd), V1589)

			__e.TailApply(tmp13882, tmp13890)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13893 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alpha_6, tmp13880)

	_ = tmp13893

	tmp13894 := MakeNative(func(__e *ControlFlow) {
		V1591 := __e.Get(1)
		_ = V1591
		tmp13926 := PrimEqual(V1591, MakeNumber(126))

		if True == tmp13926 {
			__e.Return(True)
			return
		} else {
			tmp13925 := PrimGreatThan(V1591, MakeNumber(94))

			var ifres13922 Obj

			if True == tmp13925 {
				tmp13924 := PrimLessThan(V1591, MakeNumber(123))

				var ifres13923 Obj

				if True == tmp13924 {
					ifres13923 = True

				} else {
					ifres13923 = False

				}

				ifres13922 = ifres13923

			} else {
				ifres13922 = False

			}

			var ifres13897 Obj

			if True == ifres13922 {
				ifres13897 = True

			} else {
				tmp13921 := PrimGreatThan(V1591, MakeNumber(59))

				var ifres13918 Obj

				if True == tmp13921 {
					tmp13920 := PrimLessThan(V1591, MakeNumber(91))

					var ifres13919 Obj

					if True == tmp13920 {
						ifres13919 = True

					} else {
						ifres13919 = False

					}

					ifres13918 = ifres13919

				} else {
					ifres13918 = False

				}

				var ifres13899 Obj

				if True == ifres13918 {
					ifres13899 = True

				} else {
					tmp13917 := PrimGreatThan(V1591, MakeNumber(41))

					var ifres13910 Obj

					if True == tmp13917 {
						tmp13916 := PrimLessThan(V1591, MakeNumber(58))

						var ifres13912 Obj

						if True == tmp13916 {
							tmp13914 := PrimEqual(V1591, MakeNumber(44))

							tmp13915 := PrimNot(tmp13914)

							var ifres13913 Obj

							if True == tmp13915 {
								ifres13913 = True

							} else {
								ifres13913 = False

							}

							ifres13912 = ifres13913

						} else {
							ifres13912 = False

						}

						var ifres13911 Obj

						if True == ifres13912 {
							ifres13911 = True

						} else {
							ifres13911 = False

						}

						ifres13910 = ifres13911

					} else {
						ifres13910 = False

					}

					var ifres13901 Obj

					if True == ifres13910 {
						ifres13901 = True

					} else {
						tmp13909 := PrimGreatThan(V1591, MakeNumber(34))

						var ifres13906 Obj

						if True == tmp13909 {
							tmp13908 := PrimLessThan(V1591, MakeNumber(40))

							var ifres13907 Obj

							if True == tmp13908 {
								ifres13907 = True

							} else {
								ifres13907 = False

							}

							ifres13906 = ifres13907

						} else {
							ifres13906 = False

						}

						var ifres13903 Obj

						if True == ifres13906 {
							ifres13903 = True

						} else {
							tmp13905 := PrimEqual(V1591, MakeNumber(33))

							var ifres13904 Obj

							if True == tmp13905 {
								ifres13904 = True

							} else {
								ifres13904 = False

							}

							ifres13903 = ifres13904

						}

						var ifres13902 Obj

						if True == ifres13903 {
							ifres13902 = True

						} else {
							ifres13902 = False

						}

						ifres13901 = ifres13902

					}

					var ifres13900 Obj

					if True == ifres13901 {
						ifres13900 = True

					} else {
						ifres13900 = False

					}

					ifres13899 = ifres13900

				}

				var ifres13898 Obj

				if True == ifres13899 {
					ifres13898 = True

				} else {
					ifres13898 = False

				}

				ifres13897 = ifres13898

			}

			if True == ifres13897 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp13927 := Call(__e, PrimNS1Value(symns2_1set), symshen_4symbol_1code_2, tmp13894)

	_ = tmp13927

	tmp13928 := MakeNative(func(__e *ControlFlow) {
		V1593 := __e.Get(1)
		_ = V1593
		tmp13929 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5dbq_6 := __e.Get(1)
			_ = Parse__shen_4_5dbq_6
			tmp13945 := Call(__e, PrimNS2Value(symfail))

			tmp13946 := PrimEqual(tmp13945, Parse__shen_4_5dbq_6)

			tmp13947 := PrimNot(tmp13946)

			if True == tmp13947 {
				tmp13931 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5strcontents_6 := __e.Get(1)
					_ = Parse__shen_4_5strcontents_6
					tmp13941 := Call(__e, PrimNS2Value(symfail))

					tmp13942 := PrimEqual(tmp13941, Parse__shen_4_5strcontents_6)

					tmp13943 := PrimNot(tmp13942)

					if True == tmp13943 {
						tmp13933 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5dbq_6 := __e.Get(1)
							_ = Parse__shen_4_5dbq_6
							tmp13937 := Call(__e, PrimNS2Value(symfail))

							tmp13938 := PrimEqual(tmp13937, Parse__shen_4_5dbq_6)

							tmp13939 := PrimNot(tmp13938)

							if True == tmp13939 {
								tmp13935 := PrimHead(Parse__shen_4_5dbq_6)

								tmp13936 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5strcontents_6)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp13935, tmp13936)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp13940 := Call(__e, PrimNS2Value(symshen_4_5dbq_6), Parse__shen_4_5strcontents_6)

						__e.TailApply(tmp13933, tmp13940)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13944 := Call(__e, PrimNS2Value(symshen_4_5strcontents_6), Parse__shen_4_5dbq_6)

				__e.TailApply(tmp13931, tmp13944)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13948 := Call(__e, PrimNS2Value(symshen_4_5dbq_6), V1593)

		__e.TailApply(tmp13929, tmp13948)
		return

	}, 1)

	tmp13949 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5str_6, tmp13928)

	_ = tmp13949

	tmp13950 := MakeNative(func(__e *ControlFlow) {
		V1595 := __e.Get(1)
		_ = V1595
		tmp13960 := PrimHead(V1595)

		tmp13961 := PrimIsPair(tmp13960)

		if True == tmp13961 {
			tmp13952 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp13958 := PrimEqual(Parse__Char, MakeNumber(34))

				if True == tmp13958 {
					tmp13954 := Call(__e, PrimNS2Value(symshen_4tlhd), V1595)

					tmp13955 := Call(__e, PrimNS2Value(symshen_4hdtl), V1595)

					tmp13956 := Call(__e, PrimNS2Value(symshen_4pair), tmp13954, tmp13955)

					tmp13957 := PrimHead(tmp13956)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp13957, Parse__Char)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp13959 := Call(__e, PrimNS2Value(symshen_4hdhd), V1595)

			__e.TailApply(tmp13952, tmp13959)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp13962 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5dbq_6, tmp13950)

	_ = tmp13962

	tmp13963 := MakeNative(func(__e *ControlFlow) {
		V1597 := __e.Get(1)
		_ = V1597
		tmp13964 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp13973 := Call(__e, PrimNS2Value(symfail))

			tmp13974 := PrimEqual(YaccParse, tmp13973)

			if True == tmp13974 {
				tmp13966 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp13969 := Call(__e, PrimNS2Value(symfail))

					tmp13970 := PrimEqual(tmp13969, Parse___5e_6)

					tmp13971 := PrimNot(tmp13970)

					if True == tmp13971 {
						tmp13968 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13968, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13972 := Call(__e, PrimNS2Value(sym_5e_6), V1597)

				__e.TailApply(tmp13966, tmp13972)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp13975 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5strc_6 := __e.Get(1)
			_ = Parse__shen_4_5strc_6
			tmp13987 := Call(__e, PrimNS2Value(symfail))

			tmp13988 := PrimEqual(tmp13987, Parse__shen_4_5strc_6)

			tmp13989 := PrimNot(tmp13988)

			if True == tmp13989 {
				tmp13977 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5strcontents_6 := __e.Get(1)
					_ = Parse__shen_4_5strcontents_6
					tmp13983 := Call(__e, PrimNS2Value(symfail))

					tmp13984 := PrimEqual(tmp13983, Parse__shen_4_5strcontents_6)

					tmp13985 := PrimNot(tmp13984)

					if True == tmp13985 {
						tmp13979 := PrimHead(Parse__shen_4_5strcontents_6)

						tmp13980 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5strc_6)

						tmp13981 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5strcontents_6)

						tmp13982 := PrimCons(tmp13980, tmp13981)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp13979, tmp13982)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp13986 := Call(__e, PrimNS2Value(symshen_4_5strcontents_6), Parse__shen_4_5strc_6)

				__e.TailApply(tmp13977, tmp13986)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp13990 := Call(__e, PrimNS2Value(symshen_4_5strc_6), V1597)

		tmp13991 := Call(__e, tmp13975, tmp13990)

		__e.TailApply(tmp13964, tmp13991)
		return

	}, 1)

	tmp13992 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5strcontents_6, tmp13963)

	_ = tmp13992

	tmp13993 := MakeNative(func(__e *ControlFlow) {
		V1599 := __e.Get(1)
		_ = V1599
		tmp14002 := PrimHead(V1599)

		tmp14003 := PrimIsPair(tmp14002)

		if True == tmp14003 {
			tmp13995 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp13996 := Call(__e, PrimNS2Value(symshen_4tlhd), V1599)

				tmp13997 := Call(__e, PrimNS2Value(symshen_4hdtl), V1599)

				tmp13998 := Call(__e, PrimNS2Value(symshen_4pair), tmp13996, tmp13997)

				tmp13999 := PrimHead(tmp13998)

				tmp14000 := PrimNumberToString(Parse__Char)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp13999, tmp14000)
				return

			}, 1)

			tmp14001 := Call(__e, PrimNS2Value(symshen_4hdhd), V1599)

			__e.TailApply(tmp13995, tmp14001)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14004 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5byte_6, tmp13993)

	_ = tmp14004

	tmp14005 := MakeNative(func(__e *ControlFlow) {
		V1601 := __e.Get(1)
		_ = V1601
		tmp14017 := PrimHead(V1601)

		tmp14018 := PrimIsPair(tmp14017)

		if True == tmp14018 {
			tmp14007 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp14014 := PrimEqual(Parse__Char, MakeNumber(34))

				tmp14015 := PrimNot(tmp14014)

				if True == tmp14015 {
					tmp14009 := Call(__e, PrimNS2Value(symshen_4tlhd), V1601)

					tmp14010 := Call(__e, PrimNS2Value(symshen_4hdtl), V1601)

					tmp14011 := Call(__e, PrimNS2Value(symshen_4pair), tmp14009, tmp14010)

					tmp14012 := PrimHead(tmp14011)

					tmp14013 := PrimNumberToString(Parse__Char)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14012, tmp14013)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14016 := Call(__e, PrimNS2Value(symshen_4hdhd), V1601)

			__e.TailApply(tmp14007, tmp14016)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14019 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5strc_6, tmp14005)

	_ = tmp14019

	tmp14020 := MakeNative(func(__e *ControlFlow) {
		V1603 := __e.Get(1)
		_ = V1603
		tmp14021 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14157 := Call(__e, PrimNS2Value(symfail))

			tmp14158 := PrimEqual(YaccParse, tmp14157)

			if True == tmp14158 {
				tmp14023 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp14140 := Call(__e, PrimNS2Value(symfail))

					tmp14141 := PrimEqual(YaccParse, tmp14140)

					if True == tmp14141 {
						tmp14025 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp14097 := Call(__e, PrimNS2Value(symfail))

							tmp14098 := PrimEqual(YaccParse, tmp14097)

							if True == tmp14098 {
								tmp14027 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									tmp14069 := Call(__e, PrimNS2Value(symfail))

									tmp14070 := PrimEqual(YaccParse, tmp14069)

									if True == tmp14070 {
										tmp14029 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											tmp14041 := Call(__e, PrimNS2Value(symfail))

											tmp14042 := PrimEqual(YaccParse, tmp14041)

											if True == tmp14042 {
												tmp14031 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5digits_6 := __e.Get(1)
													_ = Parse__shen_4_5digits_6
													tmp14037 := Call(__e, PrimNS2Value(symfail))

													tmp14038 := PrimEqual(tmp14037, Parse__shen_4_5digits_6)

													tmp14039 := PrimNot(tmp14038)

													if True == tmp14039 {
														tmp14033 := PrimHead(Parse__shen_4_5digits_6)

														tmp14034 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

														tmp14035 := Call(__e, PrimNS2Value(symreverse), tmp14034)

														tmp14036 := Call(__e, PrimNS2Value(symshen_4pre), tmp14035, MakeNumber(0))

														__e.TailApply(PrimNS2Value(symshen_4pair), tmp14033, tmp14036)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp14040 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V1603)

												__e.TailApply(tmp14031, tmp14040)
												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										tmp14043 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5predigits_6 := __e.Get(1)
											_ = Parse__shen_4_5predigits_6
											tmp14064 := Call(__e, PrimNS2Value(symfail))

											tmp14065 := PrimEqual(tmp14064, Parse__shen_4_5predigits_6)

											tmp14066 := PrimNot(tmp14065)

											if True == tmp14066 {
												tmp14045 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5stop_6 := __e.Get(1)
													_ = Parse__shen_4_5stop_6
													tmp14060 := Call(__e, PrimNS2Value(symfail))

													tmp14061 := PrimEqual(tmp14060, Parse__shen_4_5stop_6)

													tmp14062 := PrimNot(tmp14061)

													if True == tmp14062 {
														tmp14047 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5postdigits_6 := __e.Get(1)
															_ = Parse__shen_4_5postdigits_6
															tmp14056 := Call(__e, PrimNS2Value(symfail))

															tmp14057 := PrimEqual(tmp14056, Parse__shen_4_5postdigits_6)

															tmp14058 := PrimNot(tmp14057)

															if True == tmp14058 {
																tmp14049 := PrimHead(Parse__shen_4_5postdigits_6)

																tmp14050 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5predigits_6)

																tmp14051 := Call(__e, PrimNS2Value(symreverse), tmp14050)

																tmp14052 := Call(__e, PrimNS2Value(symshen_4pre), tmp14051, MakeNumber(0))

																tmp14053 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5postdigits_6)

																tmp14054 := Call(__e, PrimNS2Value(symshen_4post), tmp14053, MakeNumber(1))

																tmp14055 := PrimNumberAdd(tmp14052, tmp14054)

																__e.TailApply(PrimNS2Value(symshen_4pair), tmp14049, tmp14055)
																return

															} else {
																__e.TailApply(PrimNS2Value(symfail))
																return
															}

														}, 1)

														tmp14059 := Call(__e, PrimNS2Value(symshen_4_5postdigits_6), Parse__shen_4_5stop_6)

														__e.TailApply(tmp14047, tmp14059)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp14063 := Call(__e, PrimNS2Value(symshen_4_5stop_6), Parse__shen_4_5predigits_6)

												__e.TailApply(tmp14045, tmp14063)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp14067 := Call(__e, PrimNS2Value(symshen_4_5predigits_6), V1603)

										tmp14068 := Call(__e, tmp14043, tmp14067)

										__e.TailApply(tmp14029, tmp14068)
										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								tmp14071 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5digits_6 := __e.Get(1)
									_ = Parse__shen_4_5digits_6
									tmp14092 := Call(__e, PrimNS2Value(symfail))

									tmp14093 := PrimEqual(tmp14092, Parse__shen_4_5digits_6)

									tmp14094 := PrimNot(tmp14093)

									if True == tmp14094 {
										tmp14073 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5E_6 := __e.Get(1)
											_ = Parse__shen_4_5E_6
											tmp14088 := Call(__e, PrimNS2Value(symfail))

											tmp14089 := PrimEqual(tmp14088, Parse__shen_4_5E_6)

											tmp14090 := PrimNot(tmp14089)

											if True == tmp14090 {
												tmp14075 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5log10_6 := __e.Get(1)
													_ = Parse__shen_4_5log10_6
													tmp14084 := Call(__e, PrimNS2Value(symfail))

													tmp14085 := PrimEqual(tmp14084, Parse__shen_4_5log10_6)

													tmp14086 := PrimNot(tmp14085)

													if True == tmp14086 {
														tmp14077 := PrimHead(Parse__shen_4_5log10_6)

														tmp14078 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5log10_6)

														tmp14079 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), tmp14078)

														tmp14080 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

														tmp14081 := Call(__e, PrimNS2Value(symreverse), tmp14080)

														tmp14082 := Call(__e, PrimNS2Value(symshen_4pre), tmp14081, MakeNumber(0))

														tmp14083 := PrimNumberMultiply(tmp14079, tmp14082)

														__e.TailApply(PrimNS2Value(symshen_4pair), tmp14077, tmp14083)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp14087 := Call(__e, PrimNS2Value(symshen_4_5log10_6), Parse__shen_4_5E_6)

												__e.TailApply(tmp14075, tmp14087)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp14091 := Call(__e, PrimNS2Value(symshen_4_5E_6), Parse__shen_4_5digits_6)

										__e.TailApply(tmp14073, tmp14091)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp14095 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V1603)

								tmp14096 := Call(__e, tmp14071, tmp14095)

								__e.TailApply(tmp14027, tmp14096)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp14099 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5predigits_6 := __e.Get(1)
							_ = Parse__shen_4_5predigits_6
							tmp14135 := Call(__e, PrimNS2Value(symfail))

							tmp14136 := PrimEqual(tmp14135, Parse__shen_4_5predigits_6)

							tmp14137 := PrimNot(tmp14136)

							if True == tmp14137 {
								tmp14101 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5stop_6 := __e.Get(1)
									_ = Parse__shen_4_5stop_6
									tmp14131 := Call(__e, PrimNS2Value(symfail))

									tmp14132 := PrimEqual(tmp14131, Parse__shen_4_5stop_6)

									tmp14133 := PrimNot(tmp14132)

									if True == tmp14133 {
										tmp14103 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5postdigits_6 := __e.Get(1)
											_ = Parse__shen_4_5postdigits_6
											tmp14127 := Call(__e, PrimNS2Value(symfail))

											tmp14128 := PrimEqual(tmp14127, Parse__shen_4_5postdigits_6)

											tmp14129 := PrimNot(tmp14128)

											if True == tmp14129 {
												tmp14105 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5E_6 := __e.Get(1)
													_ = Parse__shen_4_5E_6
													tmp14123 := Call(__e, PrimNS2Value(symfail))

													tmp14124 := PrimEqual(tmp14123, Parse__shen_4_5E_6)

													tmp14125 := PrimNot(tmp14124)

													if True == tmp14125 {
														tmp14107 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5log10_6 := __e.Get(1)
															_ = Parse__shen_4_5log10_6
															tmp14119 := Call(__e, PrimNS2Value(symfail))

															tmp14120 := PrimEqual(tmp14119, Parse__shen_4_5log10_6)

															tmp14121 := PrimNot(tmp14120)

															if True == tmp14121 {
																tmp14109 := PrimHead(Parse__shen_4_5log10_6)

																tmp14110 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5log10_6)

																tmp14111 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), tmp14110)

																tmp14112 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5predigits_6)

																tmp14113 := Call(__e, PrimNS2Value(symreverse), tmp14112)

																tmp14114 := Call(__e, PrimNS2Value(symshen_4pre), tmp14113, MakeNumber(0))

																tmp14115 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5postdigits_6)

																tmp14116 := Call(__e, PrimNS2Value(symshen_4post), tmp14115, MakeNumber(1))

																tmp14117 := PrimNumberAdd(tmp14114, tmp14116)

																tmp14118 := PrimNumberMultiply(tmp14111, tmp14117)

																__e.TailApply(PrimNS2Value(symshen_4pair), tmp14109, tmp14118)
																return

															} else {
																__e.TailApply(PrimNS2Value(symfail))
																return
															}

														}, 1)

														tmp14122 := Call(__e, PrimNS2Value(symshen_4_5log10_6), Parse__shen_4_5E_6)

														__e.TailApply(tmp14107, tmp14122)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp14126 := Call(__e, PrimNS2Value(symshen_4_5E_6), Parse__shen_4_5postdigits_6)

												__e.TailApply(tmp14105, tmp14126)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp14130 := Call(__e, PrimNS2Value(symshen_4_5postdigits_6), Parse__shen_4_5stop_6)

										__e.TailApply(tmp14103, tmp14130)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp14134 := Call(__e, PrimNS2Value(symshen_4_5stop_6), Parse__shen_4_5predigits_6)

								__e.TailApply(tmp14101, tmp14134)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14138 := Call(__e, PrimNS2Value(symshen_4_5predigits_6), V1603)

						tmp14139 := Call(__e, tmp14099, tmp14138)

						__e.TailApply(tmp14025, tmp14139)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp14142 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5plus_6 := __e.Get(1)
					_ = Parse__shen_4_5plus_6
					tmp14152 := Call(__e, PrimNS2Value(symfail))

					tmp14153 := PrimEqual(tmp14152, Parse__shen_4_5plus_6)

					tmp14154 := PrimNot(tmp14153)

					if True == tmp14154 {
						tmp14144 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5number_6 := __e.Get(1)
							_ = Parse__shen_4_5number_6
							tmp14148 := Call(__e, PrimNS2Value(symfail))

							tmp14149 := PrimEqual(tmp14148, Parse__shen_4_5number_6)

							tmp14150 := PrimNot(tmp14149)

							if True == tmp14150 {
								tmp14146 := PrimHead(Parse__shen_4_5number_6)

								tmp14147 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5number_6)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp14146, tmp14147)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14151 := Call(__e, PrimNS2Value(symshen_4_5number_6), Parse__shen_4_5plus_6)

						__e.TailApply(tmp14144, tmp14151)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14155 := Call(__e, PrimNS2Value(symshen_4_5plus_6), V1603)

				tmp14156 := Call(__e, tmp14142, tmp14155)

				__e.TailApply(tmp14023, tmp14156)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14159 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5minus_6 := __e.Get(1)
			_ = Parse__shen_4_5minus_6
			tmp14170 := Call(__e, PrimNS2Value(symfail))

			tmp14171 := PrimEqual(tmp14170, Parse__shen_4_5minus_6)

			tmp14172 := PrimNot(tmp14171)

			if True == tmp14172 {
				tmp14161 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5number_6 := __e.Get(1)
					_ = Parse__shen_4_5number_6
					tmp14166 := Call(__e, PrimNS2Value(symfail))

					tmp14167 := PrimEqual(tmp14166, Parse__shen_4_5number_6)

					tmp14168 := PrimNot(tmp14167)

					if True == tmp14168 {
						tmp14163 := PrimHead(Parse__shen_4_5number_6)

						tmp14164 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5number_6)

						tmp14165 := PrimNumberSubtract(MakeNumber(0), tmp14164)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14163, tmp14165)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14169 := Call(__e, PrimNS2Value(symshen_4_5number_6), Parse__shen_4_5minus_6)

				__e.TailApply(tmp14161, tmp14169)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14173 := Call(__e, PrimNS2Value(symshen_4_5minus_6), V1603)

		tmp14174 := Call(__e, tmp14159, tmp14173)

		__e.TailApply(tmp14021, tmp14174)
		return

	}, 1)

	tmp14175 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5number_6, tmp14020)

	_ = tmp14175

	tmp14176 := MakeNative(func(__e *ControlFlow) {
		V1606 := __e.Get(1)
		_ = V1606
		tmp14187 := PrimHead(V1606)

		tmp14188 := PrimIsPair(tmp14187)

		var ifres14183 Obj

		if True == tmp14188 {
			tmp14185 := Call(__e, PrimNS2Value(symshen_4hdhd), V1606)

			tmp14186 := PrimEqual(MakeNumber(101), tmp14185)

			var ifres14184 Obj

			if True == tmp14186 {
				ifres14184 = True

			} else {
				ifres14184 = False

			}

			ifres14183 = ifres14184

		} else {
			ifres14183 = False

		}

		if True == ifres14183 {
			tmp14178 := MakeNative(func(__e *ControlFlow) {
				NewStream1604 := __e.Get(1)
				_ = NewStream1604
				tmp14179 := PrimHead(NewStream1604)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14179, symshen_4skip)
				return

			}, 1)

			tmp14180 := Call(__e, PrimNS2Value(symshen_4tlhd), V1606)

			tmp14181 := Call(__e, PrimNS2Value(symshen_4hdtl), V1606)

			tmp14182 := Call(__e, PrimNS2Value(symshen_4pair), tmp14180, tmp14181)

			__e.TailApply(tmp14178, tmp14182)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14189 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5E_6, tmp14176)

	_ = tmp14189

	tmp14190 := MakeNative(func(__e *ControlFlow) {
		V1608 := __e.Get(1)
		_ = V1608
		tmp14191 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14203 := Call(__e, PrimNS2Value(symfail))

			tmp14204 := PrimEqual(YaccParse, tmp14203)

			if True == tmp14204 {
				tmp14193 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					tmp14199 := Call(__e, PrimNS2Value(symfail))

					tmp14200 := PrimEqual(tmp14199, Parse__shen_4_5digits_6)

					tmp14201 := PrimNot(tmp14200)

					if True == tmp14201 {
						tmp14195 := PrimHead(Parse__shen_4_5digits_6)

						tmp14196 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

						tmp14197 := Call(__e, PrimNS2Value(symreverse), tmp14196)

						tmp14198 := Call(__e, PrimNS2Value(symshen_4pre), tmp14197, MakeNumber(0))

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14195, tmp14198)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14202 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V1608)

				__e.TailApply(tmp14193, tmp14202)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14205 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5minus_6 := __e.Get(1)
			_ = Parse__shen_4_5minus_6
			tmp14218 := Call(__e, PrimNS2Value(symfail))

			tmp14219 := PrimEqual(tmp14218, Parse__shen_4_5minus_6)

			tmp14220 := PrimNot(tmp14219)

			if True == tmp14220 {
				tmp14207 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					tmp14214 := Call(__e, PrimNS2Value(symfail))

					tmp14215 := PrimEqual(tmp14214, Parse__shen_4_5digits_6)

					tmp14216 := PrimNot(tmp14215)

					if True == tmp14216 {
						tmp14209 := PrimHead(Parse__shen_4_5digits_6)

						tmp14210 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

						tmp14211 := Call(__e, PrimNS2Value(symreverse), tmp14210)

						tmp14212 := Call(__e, PrimNS2Value(symshen_4pre), tmp14211, MakeNumber(0))

						tmp14213 := PrimNumberSubtract(MakeNumber(0), tmp14212)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14209, tmp14213)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14217 := Call(__e, PrimNS2Value(symshen_4_5digits_6), Parse__shen_4_5minus_6)

				__e.TailApply(tmp14207, tmp14217)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14221 := Call(__e, PrimNS2Value(symshen_4_5minus_6), V1608)

		tmp14222 := Call(__e, tmp14205, tmp14221)

		__e.TailApply(tmp14191, tmp14222)
		return

	}, 1)

	tmp14223 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5log10_6, tmp14190)

	_ = tmp14223

	tmp14224 := MakeNative(func(__e *ControlFlow) {
		V1610 := __e.Get(1)
		_ = V1610
		tmp14234 := PrimHead(V1610)

		tmp14235 := PrimIsPair(tmp14234)

		if True == tmp14235 {
			tmp14226 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp14232 := PrimEqual(Parse__Char, MakeNumber(43))

				if True == tmp14232 {
					tmp14228 := Call(__e, PrimNS2Value(symshen_4tlhd), V1610)

					tmp14229 := Call(__e, PrimNS2Value(symshen_4hdtl), V1610)

					tmp14230 := Call(__e, PrimNS2Value(symshen_4pair), tmp14228, tmp14229)

					tmp14231 := PrimHead(tmp14230)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14231, Parse__Char)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14233 := Call(__e, PrimNS2Value(symshen_4hdhd), V1610)

			__e.TailApply(tmp14226, tmp14233)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14236 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5plus_6, tmp14224)

	_ = tmp14236

	tmp14237 := MakeNative(func(__e *ControlFlow) {
		V1612 := __e.Get(1)
		_ = V1612
		tmp14247 := PrimHead(V1612)

		tmp14248 := PrimIsPair(tmp14247)

		if True == tmp14248 {
			tmp14239 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp14245 := PrimEqual(Parse__Char, MakeNumber(46))

				if True == tmp14245 {
					tmp14241 := Call(__e, PrimNS2Value(symshen_4tlhd), V1612)

					tmp14242 := Call(__e, PrimNS2Value(symshen_4hdtl), V1612)

					tmp14243 := Call(__e, PrimNS2Value(symshen_4pair), tmp14241, tmp14242)

					tmp14244 := PrimHead(tmp14243)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14244, Parse__Char)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14246 := Call(__e, PrimNS2Value(symshen_4hdhd), V1612)

			__e.TailApply(tmp14239, tmp14246)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14249 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5stop_6, tmp14237)

	_ = tmp14249

	tmp14250 := MakeNative(func(__e *ControlFlow) {
		V1614 := __e.Get(1)
		_ = V1614
		tmp14251 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14260 := Call(__e, PrimNS2Value(symfail))

			tmp14261 := PrimEqual(YaccParse, tmp14260)

			if True == tmp14261 {
				tmp14253 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp14256 := Call(__e, PrimNS2Value(symfail))

					tmp14257 := PrimEqual(tmp14256, Parse___5e_6)

					tmp14258 := PrimNot(tmp14257)

					if True == tmp14258 {
						tmp14255 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14255, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14259 := Call(__e, PrimNS2Value(sym_5e_6), V1614)

				__e.TailApply(tmp14253, tmp14259)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14262 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digits_6 := __e.Get(1)
			_ = Parse__shen_4_5digits_6
			tmp14266 := Call(__e, PrimNS2Value(symfail))

			tmp14267 := PrimEqual(tmp14266, Parse__shen_4_5digits_6)

			tmp14268 := PrimNot(tmp14267)

			if True == tmp14268 {
				tmp14264 := PrimHead(Parse__shen_4_5digits_6)

				tmp14265 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14264, tmp14265)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14269 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V1614)

		tmp14270 := Call(__e, tmp14262, tmp14269)

		__e.TailApply(tmp14251, tmp14270)
		return

	}, 1)

	tmp14271 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5predigits_6, tmp14250)

	_ = tmp14271

	tmp14272 := MakeNative(func(__e *ControlFlow) {
		V1616 := __e.Get(1)
		_ = V1616
		tmp14273 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digits_6 := __e.Get(1)
			_ = Parse__shen_4_5digits_6
			tmp14277 := Call(__e, PrimNS2Value(symfail))

			tmp14278 := PrimEqual(tmp14277, Parse__shen_4_5digits_6)

			tmp14279 := PrimNot(tmp14278)

			if True == tmp14279 {
				tmp14275 := PrimHead(Parse__shen_4_5digits_6)

				tmp14276 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14275, tmp14276)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14280 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V1616)

		__e.TailApply(tmp14273, tmp14280)
		return

	}, 1)

	tmp14281 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5postdigits_6, tmp14272)

	_ = tmp14281

	tmp14282 := MakeNative(func(__e *ControlFlow) {
		V1618 := __e.Get(1)
		_ = V1618
		tmp14283 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14294 := Call(__e, PrimNS2Value(symfail))

			tmp14295 := PrimEqual(YaccParse, tmp14294)

			if True == tmp14295 {
				tmp14285 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digit_6 := __e.Get(1)
					_ = Parse__shen_4_5digit_6
					tmp14290 := Call(__e, PrimNS2Value(symfail))

					tmp14291 := PrimEqual(tmp14290, Parse__shen_4_5digit_6)

					tmp14292 := PrimNot(tmp14291)

					if True == tmp14292 {
						tmp14287 := PrimHead(Parse__shen_4_5digit_6)

						tmp14288 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digit_6)

						tmp14289 := PrimCons(tmp14288, Nil)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14287, tmp14289)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14293 := Call(__e, PrimNS2Value(symshen_4_5digit_6), V1618)

				__e.TailApply(tmp14285, tmp14293)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14296 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digit_6 := __e.Get(1)
			_ = Parse__shen_4_5digit_6
			tmp14308 := Call(__e, PrimNS2Value(symfail))

			tmp14309 := PrimEqual(tmp14308, Parse__shen_4_5digit_6)

			tmp14310 := PrimNot(tmp14309)

			if True == tmp14310 {
				tmp14298 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					tmp14304 := Call(__e, PrimNS2Value(symfail))

					tmp14305 := PrimEqual(tmp14304, Parse__shen_4_5digits_6)

					tmp14306 := PrimNot(tmp14305)

					if True == tmp14306 {
						tmp14300 := PrimHead(Parse__shen_4_5digits_6)

						tmp14301 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digit_6)

						tmp14302 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5digits_6)

						tmp14303 := PrimCons(tmp14301, tmp14302)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14300, tmp14303)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14307 := Call(__e, PrimNS2Value(symshen_4_5digits_6), Parse__shen_4_5digit_6)

				__e.TailApply(tmp14298, tmp14307)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14311 := Call(__e, PrimNS2Value(symshen_4_5digit_6), V1618)

		tmp14312 := Call(__e, tmp14296, tmp14311)

		__e.TailApply(tmp14283, tmp14312)
		return

	}, 1)

	tmp14313 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5digits_6, tmp14282)

	_ = tmp14313

	tmp14314 := MakeNative(func(__e *ControlFlow) {
		V1620 := __e.Get(1)
		_ = V1620
		tmp14325 := PrimHead(V1620)

		tmp14326 := PrimIsPair(tmp14325)

		if True == tmp14326 {
			tmp14316 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp14323 := Call(__e, PrimNS2Value(symshen_4numbyte_2), Parse__X)

				if True == tmp14323 {
					tmp14318 := Call(__e, PrimNS2Value(symshen_4tlhd), V1620)

					tmp14319 := Call(__e, PrimNS2Value(symshen_4hdtl), V1620)

					tmp14320 := Call(__e, PrimNS2Value(symshen_4pair), tmp14318, tmp14319)

					tmp14321 := PrimHead(tmp14320)

					tmp14322 := Call(__e, PrimNS2Value(symshen_4byte_1_6digit), Parse__X)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14321, tmp14322)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14324 := Call(__e, PrimNS2Value(symshen_4hdhd), V1620)

			__e.TailApply(tmp14316, tmp14324)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14327 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5digit_6, tmp14314)

	_ = tmp14327

	tmp14328 := MakeNative(func(__e *ControlFlow) {
		V1622 := __e.Get(1)
		_ = V1622
		tmp14348 := PrimEqual(MakeNumber(48), V1622)

		if True == tmp14348 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp14347 := PrimEqual(MakeNumber(49), V1622)

			if True == tmp14347 {
				__e.Return(MakeNumber(1))
				return
			} else {
				tmp14346 := PrimEqual(MakeNumber(50), V1622)

				if True == tmp14346 {
					__e.Return(MakeNumber(2))
					return
				} else {
					tmp14345 := PrimEqual(MakeNumber(51), V1622)

					if True == tmp14345 {
						__e.Return(MakeNumber(3))
						return
					} else {
						tmp14344 := PrimEqual(MakeNumber(52), V1622)

						if True == tmp14344 {
							__e.Return(MakeNumber(4))
							return
						} else {
							tmp14343 := PrimEqual(MakeNumber(53), V1622)

							if True == tmp14343 {
								__e.Return(MakeNumber(5))
								return
							} else {
								tmp14342 := PrimEqual(MakeNumber(54), V1622)

								if True == tmp14342 {
									__e.Return(MakeNumber(6))
									return
								} else {
									tmp14341 := PrimEqual(MakeNumber(55), V1622)

									if True == tmp14341 {
										__e.Return(MakeNumber(7))
										return
									} else {
										tmp14340 := PrimEqual(MakeNumber(56), V1622)

										if True == tmp14340 {
											__e.Return(MakeNumber(8))
											return
										} else {
											tmp14339 := PrimEqual(MakeNumber(57), V1622)

											if True == tmp14339 {
												__e.Return(MakeNumber(9))
												return
											} else {
												__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4byte_1_6digit)
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

	}, 1)

	tmp14349 := Call(__e, PrimNS1Value(symns2_1set), symshen_4byte_1_6digit, tmp14328)

	_ = tmp14349

	tmp14350 := MakeNative(func(__e *ControlFlow) {
		V1627 := __e.Get(1)
		_ = V1627
		V1628 := __e.Get(2)
		_ = V1628
		tmp14360 := PrimEqual(Nil, V1627)

		if True == tmp14360 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp14359 := PrimIsPair(V1627)

			if True == tmp14359 {
				tmp14353 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), V1628)

				tmp14354 := PrimHead(V1627)

				tmp14355 := PrimNumberMultiply(tmp14353, tmp14354)

				tmp14356 := PrimTail(V1627)

				tmp14357 := PrimNumberAdd(V1628, MakeNumber(1))

				tmp14358 := Call(__e, PrimNS2Value(symshen_4pre), tmp14356, tmp14357)

				__e.Return(PrimNumberAdd(tmp14355, tmp14358))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4pre)
				return
			}

		}

	}, 2)

	tmp14361 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pre, tmp14350)

	_ = tmp14361

	tmp14362 := MakeNative(func(__e *ControlFlow) {
		V1633 := __e.Get(1)
		_ = V1633
		V1634 := __e.Get(2)
		_ = V1634
		tmp14373 := PrimEqual(Nil, V1633)

		if True == tmp14373 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp14372 := PrimIsPair(V1633)

			if True == tmp14372 {
				tmp14365 := PrimNumberSubtract(MakeNumber(0), V1634)

				tmp14366 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), tmp14365)

				tmp14367 := PrimHead(V1633)

				tmp14368 := PrimNumberMultiply(tmp14366, tmp14367)

				tmp14369 := PrimTail(V1633)

				tmp14370 := PrimNumberAdd(V1634, MakeNumber(1))

				tmp14371 := Call(__e, PrimNS2Value(symshen_4post), tmp14369, tmp14370)

				__e.Return(PrimNumberAdd(tmp14368, tmp14371))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4post)
				return
			}

		}

	}, 2)

	tmp14374 := Call(__e, PrimNS1Value(symns2_1set), symshen_4post, tmp14362)

	_ = tmp14374

	tmp14375 := MakeNative(func(__e *ControlFlow) {
		V1639 := __e.Get(1)
		_ = V1639
		V1640 := __e.Get(2)
		_ = V1640
		tmp14384 := PrimEqual(MakeNumber(0), V1640)

		if True == tmp14384 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp14383 := PrimGreatThan(V1640, MakeNumber(0))

			if True == tmp14383 {
				tmp14381 := PrimNumberSubtract(V1640, MakeNumber(1))

				tmp14382 := Call(__e, PrimNS2Value(symshen_4expt), V1639, tmp14381)

				__e.Return(PrimNumberMultiply(V1639, tmp14382))
				return

			} else {
				tmp14378 := PrimNumberAdd(V1640, MakeNumber(1))

				tmp14379 := Call(__e, PrimNS2Value(symshen_4expt), V1639, tmp14378)

				tmp14380 := PrimNumberDivide(tmp14379, V1639)

				__e.Return(PrimNumberMultiply(MakeNumber(1), tmp14380))
				return

			}

		}

	}, 2)

	tmp14385 := Call(__e, PrimNS1Value(symns2_1set), symshen_4expt, tmp14375)

	_ = tmp14385

	tmp14386 := MakeNative(func(__e *ControlFlow) {
		V1642 := __e.Get(1)
		_ = V1642
		tmp14387 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5st__input_6 := __e.Get(1)
			_ = Parse__shen_4_5st__input_6
			tmp14391 := Call(__e, PrimNS2Value(symfail))

			tmp14392 := PrimEqual(tmp14391, Parse__shen_4_5st__input_6)

			tmp14393 := PrimNot(tmp14392)

			if True == tmp14393 {
				tmp14389 := PrimHead(Parse__shen_4_5st__input_6)

				tmp14390 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14389, tmp14390)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14394 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), V1642)

		__e.TailApply(tmp14387, tmp14394)
		return

	}, 1)

	tmp14395 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input1_6, tmp14386)

	_ = tmp14395

	tmp14396 := MakeNative(func(__e *ControlFlow) {
		V1644 := __e.Get(1)
		_ = V1644
		tmp14397 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5st__input_6 := __e.Get(1)
			_ = Parse__shen_4_5st__input_6
			tmp14401 := Call(__e, PrimNS2Value(symfail))

			tmp14402 := PrimEqual(tmp14401, Parse__shen_4_5st__input_6)

			tmp14403 := PrimNot(tmp14402)

			if True == tmp14403 {
				tmp14399 := PrimHead(Parse__shen_4_5st__input_6)

				tmp14400 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5st__input_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14399, tmp14400)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14404 := Call(__e, PrimNS2Value(symshen_4_5st__input_6), V1644)

		__e.TailApply(tmp14397, tmp14404)
		return

	}, 1)

	tmp14405 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input2_6, tmp14396)

	_ = tmp14405

	tmp14406 := MakeNative(func(__e *ControlFlow) {
		V1646 := __e.Get(1)
		_ = V1646
		tmp14407 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14416 := Call(__e, PrimNS2Value(symfail))

			tmp14417 := PrimEqual(YaccParse, tmp14416)

			if True == tmp14417 {
				tmp14409 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5multiline_6 := __e.Get(1)
					_ = Parse__shen_4_5multiline_6
					tmp14412 := Call(__e, PrimNS2Value(symfail))

					tmp14413 := PrimEqual(tmp14412, Parse__shen_4_5multiline_6)

					tmp14414 := PrimNot(tmp14413)

					if True == tmp14414 {
						tmp14411 := PrimHead(Parse__shen_4_5multiline_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14411, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14415 := Call(__e, PrimNS2Value(symshen_4_5multiline_6), V1646)

				__e.TailApply(tmp14409, tmp14415)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14418 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5singleline_6 := __e.Get(1)
			_ = Parse__shen_4_5singleline_6
			tmp14421 := Call(__e, PrimNS2Value(symfail))

			tmp14422 := PrimEqual(tmp14421, Parse__shen_4_5singleline_6)

			tmp14423 := PrimNot(tmp14422)

			if True == tmp14423 {
				tmp14420 := PrimHead(Parse__shen_4_5singleline_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14420, symshen_4skip)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14424 := Call(__e, PrimNS2Value(symshen_4_5singleline_6), V1646)

		tmp14425 := Call(__e, tmp14418, tmp14424)

		__e.TailApply(tmp14407, tmp14425)
		return

	}, 1)

	tmp14426 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comment_6, tmp14406)

	_ = tmp14426

	tmp14427 := MakeNative(func(__e *ControlFlow) {
		V1648 := __e.Get(1)
		_ = V1648
		tmp14428 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5backslash_6 := __e.Get(1)
			_ = Parse__shen_4_5backslash_6
			tmp14449 := Call(__e, PrimNS2Value(symfail))

			tmp14450 := PrimEqual(tmp14449, Parse__shen_4_5backslash_6)

			tmp14451 := PrimNot(tmp14450)

			if True == tmp14451 {
				tmp14430 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5backslash_6 := __e.Get(1)
					_ = Parse__shen_4_5backslash_6
					tmp14445 := Call(__e, PrimNS2Value(symfail))

					tmp14446 := PrimEqual(tmp14445, Parse__shen_4_5backslash_6)

					tmp14447 := PrimNot(tmp14446)

					if True == tmp14447 {
						tmp14432 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5anysingle_6 := __e.Get(1)
							_ = Parse__shen_4_5anysingle_6
							tmp14441 := Call(__e, PrimNS2Value(symfail))

							tmp14442 := PrimEqual(tmp14441, Parse__shen_4_5anysingle_6)

							tmp14443 := PrimNot(tmp14442)

							if True == tmp14443 {
								tmp14434 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5return_6 := __e.Get(1)
									_ = Parse__shen_4_5return_6
									tmp14437 := Call(__e, PrimNS2Value(symfail))

									tmp14438 := PrimEqual(tmp14437, Parse__shen_4_5return_6)

									tmp14439 := PrimNot(tmp14438)

									if True == tmp14439 {
										tmp14436 := PrimHead(Parse__shen_4_5return_6)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp14436, symshen_4skip)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp14440 := Call(__e, PrimNS2Value(symshen_4_5return_6), Parse__shen_4_5anysingle_6)

								__e.TailApply(tmp14434, tmp14440)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14444 := Call(__e, PrimNS2Value(symshen_4_5anysingle_6), Parse__shen_4_5backslash_6)

						__e.TailApply(tmp14432, tmp14444)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14448 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), Parse__shen_4_5backslash_6)

				__e.TailApply(tmp14430, tmp14448)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14452 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), V1648)

		__e.TailApply(tmp14428, tmp14452)
		return

	}, 1)

	tmp14453 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5singleline_6, tmp14427)

	_ = tmp14453

	tmp14454 := MakeNative(func(__e *ControlFlow) {
		V1651 := __e.Get(1)
		_ = V1651
		tmp14465 := PrimHead(V1651)

		tmp14466 := PrimIsPair(tmp14465)

		var ifres14461 Obj

		if True == tmp14466 {
			tmp14463 := Call(__e, PrimNS2Value(symshen_4hdhd), V1651)

			tmp14464 := PrimEqual(MakeNumber(92), tmp14463)

			var ifres14462 Obj

			if True == tmp14464 {
				ifres14462 = True

			} else {
				ifres14462 = False

			}

			ifres14461 = ifres14462

		} else {
			ifres14461 = False

		}

		if True == ifres14461 {
			tmp14456 := MakeNative(func(__e *ControlFlow) {
				NewStream1649 := __e.Get(1)
				_ = NewStream1649
				tmp14457 := PrimHead(NewStream1649)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14457, symshen_4skip)
				return

			}, 1)

			tmp14458 := Call(__e, PrimNS2Value(symshen_4tlhd), V1651)

			tmp14459 := Call(__e, PrimNS2Value(symshen_4hdtl), V1651)

			tmp14460 := Call(__e, PrimNS2Value(symshen_4pair), tmp14458, tmp14459)

			__e.TailApply(tmp14456, tmp14460)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14467 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5backslash_6, tmp14454)

	_ = tmp14467

	tmp14468 := MakeNative(func(__e *ControlFlow) {
		V1653 := __e.Get(1)
		_ = V1653
		tmp14469 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14478 := Call(__e, PrimNS2Value(symfail))

			tmp14479 := PrimEqual(YaccParse, tmp14478)

			if True == tmp14479 {
				tmp14471 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp14474 := Call(__e, PrimNS2Value(symfail))

					tmp14475 := PrimEqual(tmp14474, Parse___5e_6)

					tmp14476 := PrimNot(tmp14475)

					if True == tmp14476 {
						tmp14473 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14473, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14477 := Call(__e, PrimNS2Value(sym_5e_6), V1653)

				__e.TailApply(tmp14471, tmp14477)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14480 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5non_1return_6 := __e.Get(1)
			_ = Parse__shen_4_5non_1return_6
			tmp14489 := Call(__e, PrimNS2Value(symfail))

			tmp14490 := PrimEqual(tmp14489, Parse__shen_4_5non_1return_6)

			tmp14491 := PrimNot(tmp14490)

			if True == tmp14491 {
				tmp14482 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5anysingle_6 := __e.Get(1)
					_ = Parse__shen_4_5anysingle_6
					tmp14485 := Call(__e, PrimNS2Value(symfail))

					tmp14486 := PrimEqual(tmp14485, Parse__shen_4_5anysingle_6)

					tmp14487 := PrimNot(tmp14486)

					if True == tmp14487 {
						tmp14484 := PrimHead(Parse__shen_4_5anysingle_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14484, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14488 := Call(__e, PrimNS2Value(symshen_4_5anysingle_6), Parse__shen_4_5non_1return_6)

				__e.TailApply(tmp14482, tmp14488)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14492 := Call(__e, PrimNS2Value(symshen_4_5non_1return_6), V1653)

		tmp14493 := Call(__e, tmp14480, tmp14492)

		__e.TailApply(tmp14469, tmp14493)
		return

	}, 1)

	tmp14494 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5anysingle_6, tmp14468)

	_ = tmp14494

	tmp14495 := MakeNative(func(__e *ControlFlow) {
		V1655 := __e.Get(1)
		_ = V1655
		tmp14508 := PrimHead(V1655)

		tmp14509 := PrimIsPair(tmp14508)

		if True == tmp14509 {
			tmp14497 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp14503 := PrimCons(MakeNumber(13), Nil)

				tmp14504 := PrimCons(MakeNumber(10), tmp14503)

				tmp14505 := Call(__e, PrimNS2Value(symelement_2), Parse__X, tmp14504)

				tmp14506 := PrimNot(tmp14505)

				if True == tmp14506 {
					tmp14499 := Call(__e, PrimNS2Value(symshen_4tlhd), V1655)

					tmp14500 := Call(__e, PrimNS2Value(symshen_4hdtl), V1655)

					tmp14501 := Call(__e, PrimNS2Value(symshen_4pair), tmp14499, tmp14500)

					tmp14502 := PrimHead(tmp14501)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14502, symshen_4skip)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14507 := Call(__e, PrimNS2Value(symshen_4hdhd), V1655)

			__e.TailApply(tmp14497, tmp14507)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14510 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5non_1return_6, tmp14495)

	_ = tmp14510

	tmp14511 := MakeNative(func(__e *ControlFlow) {
		V1657 := __e.Get(1)
		_ = V1657
		tmp14523 := PrimHead(V1657)

		tmp14524 := PrimIsPair(tmp14523)

		if True == tmp14524 {
			tmp14513 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp14519 := PrimCons(MakeNumber(13), Nil)

				tmp14520 := PrimCons(MakeNumber(10), tmp14519)

				tmp14521 := Call(__e, PrimNS2Value(symelement_2), Parse__X, tmp14520)

				if True == tmp14521 {
					tmp14515 := Call(__e, PrimNS2Value(symshen_4tlhd), V1657)

					tmp14516 := Call(__e, PrimNS2Value(symshen_4hdtl), V1657)

					tmp14517 := Call(__e, PrimNS2Value(symshen_4pair), tmp14515, tmp14516)

					tmp14518 := PrimHead(tmp14517)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14518, symshen_4skip)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14522 := Call(__e, PrimNS2Value(symshen_4hdhd), V1657)

			__e.TailApply(tmp14513, tmp14522)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14525 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5return_6, tmp14511)

	_ = tmp14525

	tmp14526 := MakeNative(func(__e *ControlFlow) {
		V1659 := __e.Get(1)
		_ = V1659
		tmp14527 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5backslash_6 := __e.Get(1)
			_ = Parse__shen_4_5backslash_6
			tmp14542 := Call(__e, PrimNS2Value(symfail))

			tmp14543 := PrimEqual(tmp14542, Parse__shen_4_5backslash_6)

			tmp14544 := PrimNot(tmp14543)

			if True == tmp14544 {
				tmp14529 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5times_6 := __e.Get(1)
					_ = Parse__shen_4_5times_6
					tmp14538 := Call(__e, PrimNS2Value(symfail))

					tmp14539 := PrimEqual(tmp14538, Parse__shen_4_5times_6)

					tmp14540 := PrimNot(tmp14539)

					if True == tmp14540 {
						tmp14531 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5anymulti_6 := __e.Get(1)
							_ = Parse__shen_4_5anymulti_6
							tmp14534 := Call(__e, PrimNS2Value(symfail))

							tmp14535 := PrimEqual(tmp14534, Parse__shen_4_5anymulti_6)

							tmp14536 := PrimNot(tmp14535)

							if True == tmp14536 {
								tmp14533 := PrimHead(Parse__shen_4_5anymulti_6)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp14533, symshen_4skip)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14537 := Call(__e, PrimNS2Value(symshen_4_5anymulti_6), Parse__shen_4_5times_6)

						__e.TailApply(tmp14531, tmp14537)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14541 := Call(__e, PrimNS2Value(symshen_4_5times_6), Parse__shen_4_5backslash_6)

				__e.TailApply(tmp14529, tmp14541)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14545 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), V1659)

		__e.TailApply(tmp14527, tmp14545)
		return

	}, 1)

	tmp14546 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5multiline_6, tmp14526)

	_ = tmp14546

	tmp14547 := MakeNative(func(__e *ControlFlow) {
		V1662 := __e.Get(1)
		_ = V1662
		tmp14558 := PrimHead(V1662)

		tmp14559 := PrimIsPair(tmp14558)

		var ifres14554 Obj

		if True == tmp14559 {
			tmp14556 := Call(__e, PrimNS2Value(symshen_4hdhd), V1662)

			tmp14557 := PrimEqual(MakeNumber(42), tmp14556)

			var ifres14555 Obj

			if True == tmp14557 {
				ifres14555 = True

			} else {
				ifres14555 = False

			}

			ifres14554 = ifres14555

		} else {
			ifres14554 = False

		}

		if True == ifres14554 {
			tmp14549 := MakeNative(func(__e *ControlFlow) {
				NewStream1660 := __e.Get(1)
				_ = NewStream1660
				tmp14550 := PrimHead(NewStream1660)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp14550, symshen_4skip)
				return

			}, 1)

			tmp14551 := Call(__e, PrimNS2Value(symshen_4tlhd), V1662)

			tmp14552 := Call(__e, PrimNS2Value(symshen_4hdtl), V1662)

			tmp14553 := Call(__e, PrimNS2Value(symshen_4pair), tmp14551, tmp14552)

			__e.TailApply(tmp14549, tmp14553)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14560 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5times_6, tmp14547)

	_ = tmp14560

	tmp14561 := MakeNative(func(__e *ControlFlow) {
		V1664 := __e.Get(1)
		_ = V1664
		tmp14562 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14597 := Call(__e, PrimNS2Value(symfail))

			tmp14598 := PrimEqual(YaccParse, tmp14597)

			if True == tmp14598 {
				tmp14564 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp14581 := Call(__e, PrimNS2Value(symfail))

					tmp14582 := PrimEqual(YaccParse, tmp14581)

					if True == tmp14582 {
						tmp14579 := PrimHead(V1664)

						tmp14580 := PrimIsPair(tmp14579)

						if True == tmp14580 {
							tmp14567 := MakeNative(func(__e *ControlFlow) {
								Parse__X := __e.Get(1)
								_ = Parse__X
								tmp14568 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5anymulti_6 := __e.Get(1)
									_ = Parse__shen_4_5anymulti_6
									tmp14571 := Call(__e, PrimNS2Value(symfail))

									tmp14572 := PrimEqual(tmp14571, Parse__shen_4_5anymulti_6)

									tmp14573 := PrimNot(tmp14572)

									if True == tmp14573 {
										tmp14570 := PrimHead(Parse__shen_4_5anymulti_6)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp14570, symshen_4skip)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp14574 := Call(__e, PrimNS2Value(symshen_4tlhd), V1664)

								tmp14575 := Call(__e, PrimNS2Value(symshen_4hdtl), V1664)

								tmp14576 := Call(__e, PrimNS2Value(symshen_4pair), tmp14574, tmp14575)

								tmp14577 := Call(__e, PrimNS2Value(symshen_4_5anymulti_6), tmp14576)

								__e.TailApply(tmp14568, tmp14577)
								return

							}, 1)

							tmp14578 := Call(__e, PrimNS2Value(symshen_4hdhd), V1664)

							__e.TailApply(tmp14567, tmp14578)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp14583 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5times_6 := __e.Get(1)
					_ = Parse__shen_4_5times_6
					tmp14592 := Call(__e, PrimNS2Value(symfail))

					tmp14593 := PrimEqual(tmp14592, Parse__shen_4_5times_6)

					tmp14594 := PrimNot(tmp14593)

					if True == tmp14594 {
						tmp14585 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5backslash_6 := __e.Get(1)
							_ = Parse__shen_4_5backslash_6
							tmp14588 := Call(__e, PrimNS2Value(symfail))

							tmp14589 := PrimEqual(tmp14588, Parse__shen_4_5backslash_6)

							tmp14590 := PrimNot(tmp14589)

							if True == tmp14590 {
								tmp14587 := PrimHead(Parse__shen_4_5backslash_6)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp14587, symshen_4skip)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14591 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), Parse__shen_4_5times_6)

						__e.TailApply(tmp14585, tmp14591)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14595 := Call(__e, PrimNS2Value(symshen_4_5times_6), V1664)

				tmp14596 := Call(__e, tmp14583, tmp14595)

				__e.TailApply(tmp14564, tmp14596)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14599 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5comment_6 := __e.Get(1)
			_ = Parse__shen_4_5comment_6
			tmp14608 := Call(__e, PrimNS2Value(symfail))

			tmp14609 := PrimEqual(tmp14608, Parse__shen_4_5comment_6)

			tmp14610 := PrimNot(tmp14609)

			if True == tmp14610 {
				tmp14601 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5anymulti_6 := __e.Get(1)
					_ = Parse__shen_4_5anymulti_6
					tmp14604 := Call(__e, PrimNS2Value(symfail))

					tmp14605 := PrimEqual(tmp14604, Parse__shen_4_5anymulti_6)

					tmp14606 := PrimNot(tmp14605)

					if True == tmp14606 {
						tmp14603 := PrimHead(Parse__shen_4_5anymulti_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14603, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14607 := Call(__e, PrimNS2Value(symshen_4_5anymulti_6), Parse__shen_4_5comment_6)

				__e.TailApply(tmp14601, tmp14607)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14611 := Call(__e, PrimNS2Value(symshen_4_5comment_6), V1664)

		tmp14612 := Call(__e, tmp14599, tmp14611)

		__e.TailApply(tmp14562, tmp14612)
		return

	}, 1)

	tmp14613 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5anymulti_6, tmp14561)

	_ = tmp14613

	tmp14614 := MakeNative(func(__e *ControlFlow) {
		V1666 := __e.Get(1)
		_ = V1666
		tmp14615 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14624 := Call(__e, PrimNS2Value(symfail))

			tmp14625 := PrimEqual(YaccParse, tmp14624)

			if True == tmp14625 {
				tmp14617 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5whitespace_6 := __e.Get(1)
					_ = Parse__shen_4_5whitespace_6
					tmp14620 := Call(__e, PrimNS2Value(symfail))

					tmp14621 := PrimEqual(tmp14620, Parse__shen_4_5whitespace_6)

					tmp14622 := PrimNot(tmp14621)

					if True == tmp14622 {
						tmp14619 := PrimHead(Parse__shen_4_5whitespace_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14619, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14623 := Call(__e, PrimNS2Value(symshen_4_5whitespace_6), V1666)

				__e.TailApply(tmp14617, tmp14623)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14626 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5whitespace_6 := __e.Get(1)
			_ = Parse__shen_4_5whitespace_6
			tmp14635 := Call(__e, PrimNS2Value(symfail))

			tmp14636 := PrimEqual(tmp14635, Parse__shen_4_5whitespace_6)

			tmp14637 := PrimNot(tmp14636)

			if True == tmp14637 {
				tmp14628 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5whitespaces_6 := __e.Get(1)
					_ = Parse__shen_4_5whitespaces_6
					tmp14631 := Call(__e, PrimNS2Value(symfail))

					tmp14632 := PrimEqual(tmp14631, Parse__shen_4_5whitespaces_6)

					tmp14633 := PrimNot(tmp14632)

					if True == tmp14633 {
						tmp14630 := PrimHead(Parse__shen_4_5whitespaces_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14630, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14634 := Call(__e, PrimNS2Value(symshen_4_5whitespaces_6), Parse__shen_4_5whitespace_6)

				__e.TailApply(tmp14628, tmp14634)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14638 := Call(__e, PrimNS2Value(symshen_4_5whitespace_6), V1666)

		tmp14639 := Call(__e, tmp14626, tmp14638)

		__e.TailApply(tmp14615, tmp14639)
		return

	}, 1)

	tmp14640 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5whitespaces_6, tmp14614)

	_ = tmp14640

	tmp14641 := MakeNative(func(__e *ControlFlow) {
		V1668 := __e.Get(1)
		_ = V1668
		tmp14662 := PrimHead(V1668)

		tmp14663 := PrimIsPair(tmp14662)

		if True == tmp14663 {
			tmp14643 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp14649 := MakeNative(func(__e *ControlFlow) {
					Parse__Case := __e.Get(1)
					_ = Parse__Case
					tmp14659 := PrimEqual(Parse__Case, MakeNumber(32))

					if True == tmp14659 {
						__e.Return(True)
						return
					} else {
						tmp14658 := PrimEqual(Parse__Case, MakeNumber(13))

						var ifres14652 Obj

						if True == tmp14658 {
							ifres14652 = True

						} else {
							tmp14657 := PrimEqual(Parse__Case, MakeNumber(10))

							var ifres14654 Obj

							if True == tmp14657 {
								ifres14654 = True

							} else {
								tmp14656 := PrimEqual(Parse__Case, MakeNumber(9))

								var ifres14655 Obj

								if True == tmp14656 {
									ifres14655 = True

								} else {
									ifres14655 = False

								}

								ifres14654 = ifres14655

							}

							var ifres14653 Obj

							if True == ifres14654 {
								ifres14653 = True

							} else {
								ifres14653 = False

							}

							ifres14652 = ifres14653

						}

						if True == ifres14652 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				tmp14660 := Call(__e, tmp14649, Parse__X)

				if True == tmp14660 {
					tmp14645 := Call(__e, PrimNS2Value(symshen_4tlhd), V1668)

					tmp14646 := Call(__e, PrimNS2Value(symshen_4hdtl), V1668)

					tmp14647 := Call(__e, PrimNS2Value(symshen_4pair), tmp14645, tmp14646)

					tmp14648 := PrimHead(tmp14647)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp14648, symshen_4skip)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp14661 := Call(__e, PrimNS2Value(symshen_4hdhd), V1668)

			__e.TailApply(tmp14643, tmp14661)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp14664 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5whitespace_6, tmp14641)

	_ = tmp14664

	tmp14665 := MakeNative(func(__e *ControlFlow) {
		V1670 := __e.Get(1)
		_ = V1670
		tmp14700 := PrimEqual(Nil, V1670)

		if True == tmp14700 {
			__e.Return(Nil)
			return
		} else {
			tmp14699 := PrimIsPair(V1670)

			var ifres14679 Obj

			if True == tmp14699 {
				tmp14697 := PrimTail(V1670)

				tmp14698 := PrimIsPair(tmp14697)

				var ifres14681 Obj

				if True == tmp14698 {
					tmp14694 := PrimTail(V1670)

					tmp14695 := PrimTail(tmp14694)

					tmp14696 := PrimIsPair(tmp14695)

					var ifres14683 Obj

					if True == tmp14696 {
						tmp14690 := PrimTail(V1670)

						tmp14691 := PrimTail(tmp14690)

						tmp14692 := PrimTail(tmp14691)

						tmp14693 := PrimEqual(Nil, tmp14692)

						var ifres14685 Obj

						if True == tmp14693 {
							tmp14687 := PrimTail(V1670)

							tmp14688 := PrimHead(tmp14687)

							tmp14689 := PrimEqual(tmp14688, symbar_b)

							var ifres14686 Obj

							if True == tmp14689 {
								ifres14686 = True

							} else {
								ifres14686 = False

							}

							ifres14685 = ifres14686

						} else {
							ifres14685 = False

						}

						var ifres14684 Obj

						if True == ifres14685 {
							ifres14684 = True

						} else {
							ifres14684 = False

						}

						ifres14683 = ifres14684

					} else {
						ifres14683 = False

					}

					var ifres14682 Obj

					if True == ifres14683 {
						ifres14682 = True

					} else {
						ifres14682 = False

					}

					ifres14681 = ifres14682

				} else {
					ifres14681 = False

				}

				var ifres14680 Obj

				if True == ifres14681 {
					ifres14680 = True

				} else {
					ifres14680 = False

				}

				ifres14679 = ifres14680

			} else {
				ifres14679 = False

			}

			if True == ifres14679 {
				tmp14675 := PrimHead(V1670)

				tmp14676 := PrimTail(V1670)

				tmp14677 := PrimTail(tmp14676)

				tmp14678 := PrimCons(tmp14675, tmp14677)

				__e.Return(PrimCons(symcons, tmp14678))
				return

			} else {
				tmp14674 := PrimIsPair(V1670)

				if True == tmp14674 {
					tmp14669 := PrimHead(V1670)

					tmp14670 := PrimTail(V1670)

					tmp14671 := Call(__e, PrimNS2Value(symshen_4cons__form), tmp14670)

					tmp14672 := PrimCons(tmp14671, Nil)

					tmp14673 := PrimCons(tmp14669, tmp14672)

					__e.Return(PrimCons(symcons, tmp14673))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4cons__form)
					return
				}

			}

		}

	}, 1)

	tmp14701 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cons__form, tmp14665)

	_ = tmp14701

	tmp14702 := MakeNative(func(__e *ControlFlow) {
		V1675 := __e.Get(1)
		_ = V1675
		V1676 := __e.Get(2)
		_ = V1676
		tmp14785 := PrimIsPair(V1675)

		var ifres14772 Obj

		if True == tmp14785 {
			tmp14783 := PrimHead(V1675)

			tmp14784 := PrimEqual(sym_3, tmp14783)

			var ifres14774 Obj

			if True == tmp14784 {
				tmp14781 := PrimTail(V1675)

				tmp14782 := PrimIsPair(tmp14781)

				var ifres14776 Obj

				if True == tmp14782 {
					tmp14778 := PrimTail(V1675)

					tmp14779 := PrimTail(tmp14778)

					tmp14780 := PrimEqual(Nil, tmp14779)

					var ifres14777 Obj

					if True == tmp14780 {
						ifres14777 = True

					} else {
						ifres14777 = False

					}

					ifres14776 = ifres14777

				} else {
					ifres14776 = False

				}

				var ifres14775 Obj

				if True == ifres14776 {
					ifres14775 = True

				} else {
					ifres14775 = False

				}

				ifres14774 = ifres14775

			} else {
				ifres14774 = False

			}

			var ifres14773 Obj

			if True == ifres14774 {
				ifres14773 = True

			} else {
				ifres14773 = False

			}

			ifres14772 = ifres14773

		} else {
			ifres14772 = False

		}

		if True == ifres14772 {
			tmp14769 := PrimTail(V1675)

			tmp14770 := PrimHead(tmp14769)

			tmp14771 := Call(__e, PrimNS2Value(symexplode), tmp14770)

			__e.TailApply(PrimNS2Value(symappend), tmp14771, V1676)
			return

		} else {
			tmp14768 := PrimIsPair(V1675)

			var ifres14750 Obj

			if True == tmp14768 {
				tmp14766 := PrimHead(V1675)

				tmp14767 := PrimEqual(sympackage, tmp14766)

				var ifres14752 Obj

				if True == tmp14767 {
					tmp14764 := PrimTail(V1675)

					tmp14765 := PrimIsPair(tmp14764)

					var ifres14754 Obj

					if True == tmp14765 {
						tmp14761 := PrimTail(V1675)

						tmp14762 := PrimHead(tmp14761)

						tmp14763 := PrimEqual(symnull, tmp14762)

						var ifres14756 Obj

						if True == tmp14763 {
							tmp14758 := PrimTail(V1675)

							tmp14759 := PrimTail(tmp14758)

							tmp14760 := PrimIsPair(tmp14759)

							var ifres14757 Obj

							if True == tmp14760 {
								ifres14757 = True

							} else {
								ifres14757 = False

							}

							ifres14756 = ifres14757

						} else {
							ifres14756 = False

						}

						var ifres14755 Obj

						if True == ifres14756 {
							ifres14755 = True

						} else {
							ifres14755 = False

						}

						ifres14754 = ifres14755

					} else {
						ifres14754 = False

					}

					var ifres14753 Obj

					if True == ifres14754 {
						ifres14753 = True

					} else {
						ifres14753 = False

					}

					ifres14752 = ifres14753

				} else {
					ifres14752 = False

				}

				var ifres14751 Obj

				if True == ifres14752 {
					ifres14751 = True

				} else {
					ifres14751 = False

				}

				ifres14750 = ifres14751

			} else {
				ifres14750 = False

			}

			if True == ifres14750 {
				tmp14747 := PrimTail(V1675)

				tmp14748 := PrimTail(tmp14747)

				tmp14749 := PrimTail(tmp14748)

				__e.TailApply(PrimNS2Value(symappend), tmp14749, V1676)
				return

			} else {
				tmp14746 := PrimIsPair(V1675)

				var ifres14733 Obj

				if True == tmp14746 {
					tmp14744 := PrimHead(V1675)

					tmp14745 := PrimEqual(sympackage, tmp14744)

					var ifres14735 Obj

					if True == tmp14745 {
						tmp14742 := PrimTail(V1675)

						tmp14743 := PrimIsPair(tmp14742)

						var ifres14737 Obj

						if True == tmp14743 {
							tmp14739 := PrimTail(V1675)

							tmp14740 := PrimTail(tmp14739)

							tmp14741 := PrimIsPair(tmp14740)

							var ifres14738 Obj

							if True == tmp14741 {
								ifres14738 = True

							} else {
								ifres14738 = False

							}

							ifres14737 = ifres14738

						} else {
							ifres14737 = False

						}

						var ifres14736 Obj

						if True == ifres14737 {
							ifres14736 = True

						} else {
							ifres14736 = False

						}

						ifres14735 = ifres14736

					} else {
						ifres14735 = False

					}

					var ifres14734 Obj

					if True == ifres14735 {
						ifres14734 = True

					} else {
						ifres14734 = False

					}

					ifres14733 = ifres14734

				} else {
					ifres14733 = False

				}

				if True == ifres14733 {
					tmp14706 := MakeNative(func(__e *ControlFlow) {
						ListofExceptions := __e.Get(1)
						_ = ListofExceptions
						tmp14707 := MakeNative(func(__e *ControlFlow) {
							External := __e.Get(1)
							_ = External
							tmp14708 := MakeNative(func(__e *ControlFlow) {
								PackageNameDot := __e.Get(1)
								_ = PackageNameDot
								tmp14709 := MakeNative(func(__e *ControlFlow) {
									ExpPackageNameDot := __e.Get(1)
									_ = ExpPackageNameDot
									tmp14710 := MakeNative(func(__e *ControlFlow) {
										Packaged := __e.Get(1)
										_ = Packaged
										tmp14711 := MakeNative(func(__e *ControlFlow) {
											Internal := __e.Get(1)
											_ = Internal
											__e.TailApply(PrimNS2Value(symappend), Packaged, V1676)
											return
										}, 1)

										tmp14712 := PrimTail(V1675)

										tmp14713 := PrimHead(tmp14712)

										tmp14714 := Call(__e, PrimNS2Value(symshen_4internal_1symbols), ExpPackageNameDot, Packaged)

										tmp14715 := Call(__e, PrimNS2Value(symshen_4record_1internal), tmp14713, tmp14714)

										__e.TailApply(tmp14711, tmp14715)
										return

									}, 1)

									tmp14716 := PrimTail(V1675)

									tmp14717 := PrimTail(tmp14716)

									tmp14718 := PrimTail(tmp14717)

									tmp14719 := Call(__e, PrimNS2Value(symshen_4packageh), PackageNameDot, ListofExceptions, tmp14718, ExpPackageNameDot)

									__e.TailApply(tmp14710, tmp14719)
									return

								}, 1)

								tmp14720 := Call(__e, PrimNS2Value(symexplode), PackageNameDot)

								__e.TailApply(tmp14709, tmp14720)
								return

							}, 1)

							tmp14721 := PrimTail(V1675)

							tmp14722 := PrimHead(tmp14721)

							tmp14723 := PrimStr(tmp14722)

							tmp14724 := PrimStringConcat(tmp14723, MakeString("."))

							tmp14725 := PrimIntern(tmp14724)

							__e.TailApply(tmp14708, tmp14725)
							return

						}, 1)

						tmp14726 := PrimTail(V1675)

						tmp14727 := PrimHead(tmp14726)

						tmp14728 := Call(__e, PrimNS2Value(symshen_4record_1exceptions), ListofExceptions, tmp14727)

						__e.TailApply(tmp14707, tmp14728)
						return

					}, 1)

					tmp14729 := PrimTail(V1675)

					tmp14730 := PrimTail(tmp14729)

					tmp14731 := PrimHead(tmp14730)

					tmp14732 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), tmp14731)

					__e.TailApply(tmp14706, tmp14732)
					return

				} else {
					__e.Return(PrimCons(V1675, V1676))
					return
				}

			}

		}

	}, 2)

	tmp14786 := Call(__e, PrimNS1Value(symns2_1set), symshen_4package_1macro, tmp14702)

	_ = tmp14786

	tmp14787 := MakeNative(func(__e *ControlFlow) {
		V1679 := __e.Get(1)
		_ = V1679
		V1680 := __e.Get(2)
		_ = V1680
		tmp14788 := MakeNative(func(__e *ControlFlow) {
			CurrExceptions := __e.Get(1)
			_ = CurrExceptions
			tmp14789 := MakeNative(func(__e *ControlFlow) {
				AllExceptions := __e.Get(1)
				_ = AllExceptions
				tmp14790 := PrimNS3Value(sym_dproperty_1vector_d)

				__e.TailApply(PrimNS2Value(symput), V1680, symshen_4external_1symbols, AllExceptions, tmp14790)
				return

			}, 1)

			tmp14791 := Call(__e, PrimNS2Value(symunion), V1679, CurrExceptions)

			__e.TailApply(tmp14789, tmp14791)
			return

		}, 1)

		tmp14792 := MakeNative(func(__e *ControlFlow) {
			tmp14793 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V1680, symshen_4external_1symbols, tmp14793)
			return

		}, 0)

		tmp14794 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp14795 := Call(__e, PrimNS1Value(symtry_1catch), tmp14792, tmp14794)

		__e.TailApply(tmp14788, tmp14795)
		return

	}, 2)

	tmp14796 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1exceptions, tmp14787)

	_ = tmp14796

	tmp14797 := MakeNative(func(__e *ControlFlow) {
		V1683 := __e.Get(1)
		_ = V1683
		V1684 := __e.Get(2)
		_ = V1684
		tmp14798 := MakeNative(func(__e *ControlFlow) {
			tmp14799 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V1683, symshen_4internal_1symbols, tmp14799)
			return

		}, 0)

		tmp14800 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp14801 := Call(__e, PrimNS1Value(symtry_1catch), tmp14798, tmp14800)

		tmp14802 := Call(__e, PrimNS2Value(symunion), V1684, tmp14801)

		tmp14803 := PrimNS3Value(sym_dproperty_1vector_d)

		__e.TailApply(PrimNS2Value(symput), V1683, symshen_4internal_1symbols, tmp14802, tmp14803)
		return

	}, 2)

	tmp14804 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1internal, tmp14797)

	_ = tmp14804

	tmp14805 := MakeNative(func(__e *ControlFlow) {
		V1695 := __e.Get(1)
		_ = V1695
		V1696 := __e.Get(2)
		_ = V1696
		tmp14817 := PrimIsSymbol(V1696)

		var ifres14813 Obj

		if True == tmp14817 {
			tmp14815 := Call(__e, PrimNS2Value(symexplode), V1696)

			tmp14816 := Call(__e, PrimNS2Value(symshen_4prefix_2), V1695, tmp14815)

			var ifres14814 Obj

			if True == tmp14816 {
				ifres14814 = True

			} else {
				ifres14814 = False

			}

			ifres14813 = ifres14814

		} else {
			ifres14813 = False

		}

		if True == ifres14813 {
			__e.Return(PrimCons(V1696, Nil))
			return
		} else {
			tmp14812 := PrimIsPair(V1696)

			if True == tmp14812 {
				tmp14808 := PrimHead(V1696)

				tmp14809 := Call(__e, PrimNS2Value(symshen_4internal_1symbols), V1695, tmp14808)

				tmp14810 := PrimTail(V1696)

				tmp14811 := Call(__e, PrimNS2Value(symshen_4internal_1symbols), V1695, tmp14810)

				__e.TailApply(PrimNS2Value(symunion), tmp14809, tmp14811)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 2)

	tmp14818 := Call(__e, PrimNS1Value(symns2_1set), symshen_4internal_1symbols, tmp14805)

	_ = tmp14818

	tmp14819 := MakeNative(func(__e *ControlFlow) {
		V1713 := __e.Get(1)
		_ = V1713
		V1714 := __e.Get(2)
		_ = V1714
		V1715 := __e.Get(3)
		_ = V1715
		V1716 := __e.Get(4)
		_ = V1716
		tmp14857 := PrimIsPair(V1715)

		if True == tmp14857 {
			tmp14853 := PrimHead(V1715)

			tmp14854 := Call(__e, PrimNS2Value(symshen_4packageh), V1713, V1714, tmp14853, V1716)

			tmp14855 := PrimTail(V1715)

			tmp14856 := Call(__e, PrimNS2Value(symshen_4packageh), V1713, V1714, tmp14855, V1716)

			__e.Return(PrimCons(tmp14854, tmp14856))
			return

		} else {
			tmp14852 := Call(__e, PrimNS2Value(symshen_4sysfunc_2), V1715)

			var ifres14840 Obj

			if True == tmp14852 {
				ifres14840 = True

			} else {
				tmp14851 := PrimIsVariable(V1715)

				var ifres14842 Obj

				if True == tmp14851 {
					ifres14842 = True

				} else {
					tmp14850 := Call(__e, PrimNS2Value(symelement_2), V1715, V1714)

					var ifres14844 Obj

					if True == tmp14850 {
						ifres14844 = True

					} else {
						tmp14849 := Call(__e, PrimNS2Value(symshen_4doubleunderline_2), V1715)

						var ifres14846 Obj

						if True == tmp14849 {
							ifres14846 = True

						} else {
							tmp14848 := Call(__e, PrimNS2Value(symshen_4singleunderline_2), V1715)

							var ifres14847 Obj

							if True == tmp14848 {
								ifres14847 = True

							} else {
								ifres14847 = False

							}

							ifres14846 = ifres14847

						}

						var ifres14845 Obj

						if True == ifres14846 {
							ifres14845 = True

						} else {
							ifres14845 = False

						}

						ifres14844 = ifres14845

					}

					var ifres14843 Obj

					if True == ifres14844 {
						ifres14843 = True

					} else {
						ifres14843 = False

					}

					ifres14842 = ifres14843

				}

				var ifres14841 Obj

				if True == ifres14842 {
					ifres14841 = True

				} else {
					ifres14841 = False

				}

				ifres14840 = ifres14841

			}

			if True == ifres14840 {
				__e.Return(V1715)
				return
			} else {
				tmp14839 := PrimIsSymbol(V1715)

				var ifres14823 Obj

				if True == tmp14839 {
					tmp14825 := MakeNative(func(__e *ControlFlow) {
						ExplodeX := __e.Get(1)
						_ = ExplodeX
						tmp14830 := PrimCons(MakeString("."), Nil)

						tmp14831 := PrimCons(MakeString("n"), tmp14830)

						tmp14832 := PrimCons(MakeString("e"), tmp14831)

						tmp14833 := PrimCons(MakeString("h"), tmp14832)

						tmp14834 := PrimCons(MakeString("s"), tmp14833)

						tmp14835 := Call(__e, PrimNS2Value(symshen_4prefix_2), tmp14834, ExplodeX)

						tmp14836 := PrimNot(tmp14835)

						if True == tmp14836 {
							tmp14828 := Call(__e, PrimNS2Value(symshen_4prefix_2), V1716, ExplodeX)

							tmp14829 := PrimNot(tmp14828)

							if True == tmp14829 {
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

					tmp14837 := Call(__e, PrimNS2Value(symexplode), V1715)

					tmp14838 := Call(__e, tmp14825, tmp14837)

					var ifres14824 Obj

					if True == tmp14838 {
						ifres14824 = True

					} else {
						ifres14824 = False

					}

					ifres14823 = ifres14824

				} else {
					ifres14823 = False

				}

				if True == ifres14823 {
					__e.TailApply(PrimNS2Value(symconcat), V1713, V1715)
					return
				} else {
					__e.Return(V1715)
					return
				}

			}

		}

	}, 4)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4packageh, tmp14819)
	return

}, 0)
