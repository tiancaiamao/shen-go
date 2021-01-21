package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp85897 := MakeNative(func(__e *ControlFlow) {
		V1430 := __e.Get(1)
		_ = V1430
		tmp85898 := Call(__e, PrimNS1Value(symns2_1value), symread_1byte)

		__e.TailApply(tmp85898, V1430)
		return

	}, 1)

	tmp85899 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1char_1code, tmp85897)

	_ = tmp85899

	tmp85900 := MakeNative(func(__e *ControlFlow) {
		V1432 := __e.Get(1)
		_ = V1432
		tmp85901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist)

		tmp85902 := MakeNative(func(__e *ControlFlow) {
			S := __e.Get(1)
			_ = S
			tmp85903 := Call(__e, PrimNS1Value(symns2_1value), symread_1byte)

			__e.TailApply(tmp85903, S)
			return

		}, 1)

		__e.TailApply(tmp85901, V1432, tmp85902)
		return

	}, 1)

	tmp85904 := Call(__e, PrimNS1Value(symns2_1set), symread_1file_1as_1bytelist, tmp85900)

	_ = tmp85904

	tmp85905 := MakeNative(func(__e *ControlFlow) {
		V1434 := __e.Get(1)
		_ = V1434
		tmp85906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist)

		tmp85907 := MakeNative(func(__e *ControlFlow) {
			S := __e.Get(1)
			_ = S
			tmp85908 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

			__e.TailApply(tmp85908, S)
			return

		}, 1)

		__e.TailApply(tmp85906, V1434, tmp85907)
		return

	}, 1)

	tmp85909 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1charlist, tmp85905)

	_ = tmp85909

	tmp85910 := MakeNative(func(__e *ControlFlow) {
		V1437 := __e.Get(1)
		_ = V1437
		V1438 := __e.Get(2)
		_ = V1438
		tmp85911 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp85912 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp85913 := MakeNative(func(__e *ControlFlow) {
					Xs := __e.Get(1)
					_ = Xs
					tmp85914 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						tmp85915 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						__e.TailApply(tmp85915, Xs)
						return

					}, 1)

					tmp85916 := Call(__e, PrimNS1Value(symns2_1value), symclose)

					tmp85917 := Call(__e, tmp85916, Stream)

					__e.TailApply(tmp85914, tmp85917)
					return

				}, 1)

				tmp85918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist_1help)

				tmp85919 := Call(__e, tmp85918, Stream, V1438, X, Nil)

				__e.TailApply(tmp85913, tmp85919)
				return

			}, 1)

			tmp85920 := Call(__e, V1438, Stream)

			__e.TailApply(tmp85912, tmp85920)
			return

		}, 1)

		tmp85921 := Call(__e, PrimNS1Value(symns2_1value), symopen)

		tmp85922 := Call(__e, tmp85921, V1437, symin)

		__e.TailApply(tmp85911, tmp85922)
		return

	}, 2)

	tmp85923 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1Xlist, tmp85910)

	_ = tmp85923

	tmp85924 := MakeNative(func(__e *ControlFlow) {
		V1443 := __e.Get(1)
		_ = V1443
		V1444 := __e.Get(2)
		_ = V1444
		V1445 := __e.Get(3)
		_ = V1445
		V1446 := __e.Get(4)
		_ = V1446
		tmp85930 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85931 := Call(__e, tmp85930, MakeNumber(-1), V1445)

		if True == tmp85931 {
			__e.Return(V1446)
			return
		} else {
			tmp85926 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist_1help)

			tmp85927 := Call(__e, V1444, V1443)

			tmp85928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp85929 := Call(__e, tmp85928, V1445, V1446)

			__e.TailApply(tmp85926, V1443, V1444, tmp85927, tmp85929)
			return

		}

	}, 4)

	tmp85932 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1Xlist_1help, tmp85924)

	_ = tmp85932

	tmp85933 := MakeNative(func(__e *ControlFlow) {
		V1448 := __e.Get(1)
		_ = V1448
		tmp85934 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp85935 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rfas_1h)

			tmp85936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

			tmp85937 := Call(__e, tmp85936, Stream)

			__e.TailApply(tmp85935, Stream, tmp85937, MakeString(""))
			return

		}, 1)

		tmp85938 := Call(__e, PrimNS1Value(symns2_1value), symopen)

		tmp85939 := Call(__e, tmp85938, V1448, symin)

		__e.TailApply(tmp85934, tmp85939)
		return

	}, 1)

	tmp85940 := Call(__e, PrimNS1Value(symns2_1set), symread_1file_1as_1string, tmp85933)

	_ = tmp85940

	tmp85941 := MakeNative(func(__e *ControlFlow) {
		V1452 := __e.Get(1)
		_ = V1452
		V1453 := __e.Get(2)
		_ = V1453
		V1454 := __e.Get(3)
		_ = V1454
		tmp85952 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp85953 := Call(__e, tmp85952, MakeNumber(-1), V1453)

		if True == tmp85953 {
			tmp85950 := Call(__e, PrimNS1Value(symns2_1value), symclose)

			tmp85951 := Call(__e, tmp85950, V1452)

			_ = tmp85951

			__e.Return(V1454)
			return

		} else {
			tmp85943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rfas_1h)

			tmp85944 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

			tmp85945 := Call(__e, tmp85944, V1452)

			tmp85946 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp85947 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			tmp85948 := Call(__e, tmp85947, V1453)

			tmp85949 := Call(__e, tmp85946, V1454, tmp85948)

			__e.TailApply(tmp85943, V1452, tmp85945, tmp85949)
			return

		}

	}, 3)

	tmp85954 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rfas_1h, tmp85941)

	_ = tmp85954

	tmp85955 := MakeNative(func(__e *ControlFlow) {
		V1456 := __e.Get(1)
		_ = V1456
		tmp85956 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

		tmp85957 := Call(__e, PrimNS1Value(symns2_1value), symread)

		tmp85958 := Call(__e, tmp85957, V1456)

		__e.TailApply(tmp85956, tmp85958)
		return

	}, 1)

	tmp85959 := Call(__e, PrimNS1Value(symns2_1set), syminput, tmp85955)

	_ = tmp85959

	tmp85960 := MakeNative(func(__e *ControlFlow) {
		V1459 := __e.Get(1)
		_ = V1459
		V1460 := __e.Get(2)
		_ = V1460
		tmp85961 := MakeNative(func(__e *ControlFlow) {
			Mono_2 := __e.Get(1)
			_ = Mono_2
			tmp85962 := MakeNative(func(__e *ControlFlow) {
				Input := __e.Get(1)
				_ = Input
				tmp85974 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp85975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck)

				tmp85976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				tmp85977 := Call(__e, tmp85976, V1459)

				tmp85978 := Call(__e, tmp85975, Input, tmp85977)

				tmp85979 := Call(__e, tmp85974, False, tmp85978)

				if True == tmp85979 {
					tmp85965 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					tmp85966 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp85967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp85968 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp85969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp85970 := Call(__e, tmp85969, V1459, MakeString("\n"), symshen_4r)

					tmp85971 := Call(__e, tmp85968, MakeString(" is not of type "), tmp85970)

					tmp85972 := Call(__e, tmp85967, Input, tmp85971, symshen_4r)

					tmp85973 := Call(__e, tmp85966, MakeString("type error: "), tmp85972)

					__e.TailApply(tmp85965, tmp85973)
					return

				} else {
					tmp85964 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

					__e.TailApply(tmp85964, Input)
					return

				}

			}, 1)

			tmp85980 := Call(__e, PrimNS1Value(symns2_1value), symread)

			tmp85981 := Call(__e, tmp85980, V1460)

			__e.TailApply(tmp85962, tmp85981)
			return

		}, 1)

		tmp85982 := Call(__e, PrimNS1Value(symns2_1value), symshen_4monotype)

		tmp85983 := Call(__e, tmp85982, V1459)

		__e.TailApply(tmp85961, tmp85983)
		return

	}, 2)

	tmp85984 := Call(__e, PrimNS1Value(symns2_1set), syminput_7, tmp85960)

	_ = tmp85984

	tmp85985 := MakeNative(func(__e *ControlFlow) {
		V1462 := __e.Get(1)
		_ = V1462
		tmp85998 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp85999 := Call(__e, tmp85998, V1462)

		if True == tmp85999 {
			tmp85995 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp85996 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				tmp85997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4monotype)

				__e.TailApply(tmp85997, Z)
				return

			}, 1)

			__e.TailApply(tmp85995, tmp85996, V1462)
			return

		} else {
			tmp85993 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			tmp85994 := Call(__e, tmp85993, V1462)

			if True == tmp85994 {
				tmp85988 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp85989 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp85990 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp85991 := Call(__e, tmp85990, V1462, MakeString("\n"), symshen_4a)

				tmp85992 := Call(__e, tmp85989, MakeString("input+ expects a monotype: not "), tmp85991)

				__e.TailApply(tmp85988, tmp85992)
				return

			} else {
				__e.Return(V1462)
				return
			}

		}

	}, 1)

	tmp86000 := Call(__e, PrimNS1Value(symns2_1set), symshen_4monotype, tmp85985)

	_ = tmp86000

	tmp86001 := MakeNative(func(__e *ControlFlow) {
		V1464 := __e.Get(1)
		_ = V1464
		tmp86002 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86003 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1loop)

		tmp86004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

		tmp86005 := Call(__e, tmp86004, V1464)

		tmp86006 := Call(__e, tmp86003, V1464, tmp86005, Nil)

		__e.TailApply(tmp86002, tmp86006)
		return

	}, 1)

	tmp86007 := Call(__e, PrimNS1Value(symns2_1set), symread, tmp86001)

	_ = tmp86007

	tmp86008 := MakeNative(func(__e *ControlFlow) {
		tmp86009 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp86009, symshen_4_dit_d)
		return

	}, 0)

	tmp86010 := Call(__e, PrimNS1Value(symns2_1set), symit, tmp86008)

	_ = tmp86010

	tmp86011 := MakeNative(func(__e *ControlFlow) {
		V1472 := __e.Get(1)
		_ = V1472
		V1473 := __e.Get(2)
		_ = V1473
		V1474 := __e.Get(3)
		_ = V1474
		tmp86059 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp86060 := Call(__e, tmp86059, MakeNumber(94), V1473)

		if True == tmp86060 {
			tmp86058 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp86058, MakeString("read aborted"))
			return

		} else {
			tmp86056 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86057 := Call(__e, tmp86056, MakeNumber(-1), V1473)

			if True == tmp86057 {
				tmp86054 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				tmp86055 := Call(__e, tmp86054, V1474)

				if True == tmp86055 {
					tmp86053 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(tmp86053, MakeString("error: empty stream"))
					return

				} else {
					tmp86049 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

					tmp86050 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp86051 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

						__e.TailApply(tmp86051, X)
						return

					}, 1)

					tmp86052 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(E)
						return
					}, 1)

					__e.TailApply(tmp86049, tmp86050, V1474, tmp86052)
					return

				}

			} else {
				tmp86046 := Call(__e, PrimNS1Value(symns2_1value), symshen_4terminator_2)

				tmp86047 := Call(__e, tmp86046, V1473)

				if True == tmp86047 {
					tmp86022 := MakeNative(func(__e *ControlFlow) {
						AllChars := __e.Get(1)
						_ = AllChars
						tmp86023 := MakeNative(func(__e *ControlFlow) {
							It := __e.Get(1)
							_ = It
							tmp86024 := MakeNative(func(__e *ControlFlow) {
								Read := __e.Get(1)
								_ = Read
								tmp86033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp86034 := Call(__e, tmp86033, Read, symshen_4nextbyte)

								var ifres86029 Obj

								if True == tmp86034 {
									ifres86029 = True

								} else {
									tmp86031 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

									tmp86032 := Call(__e, tmp86031, Read)

									var ifres86030 Obj

									if True == tmp86032 {
										ifres86030 = True

									} else {
										ifres86030 = False

									}

									ifres86029 = ifres86030

								}

								if True == ifres86029 {
									tmp86026 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1loop)

									tmp86027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

									tmp86028 := Call(__e, tmp86027, V1472)

									__e.TailApply(tmp86026, V1472, tmp86028, AllChars)
									return

								} else {
									__e.Return(Read)
									return
								}

							}, 1)

							tmp86035 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

							tmp86036 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp86037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

								__e.TailApply(tmp86037, X)
								return

							}, 1)

							tmp86038 := MakeNative(func(__e *ControlFlow) {
								E := __e.Get(1)
								_ = E
								__e.Return(symshen_4nextbyte)
								return
							}, 1)

							tmp86039 := Call(__e, tmp86035, tmp86036, AllChars, tmp86038)

							__e.TailApply(tmp86024, tmp86039)
							return

						}, 1)

						tmp86040 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it)

						tmp86041 := Call(__e, tmp86040, AllChars)

						__e.TailApply(tmp86023, tmp86041)
						return

					}, 1)

					tmp86042 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp86043 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp86044 := Call(__e, tmp86043, V1473, Nil)

					tmp86045 := Call(__e, tmp86042, V1474, tmp86044)

					__e.TailApply(tmp86022, tmp86045)
					return

				} else {
					tmp86015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1loop)

					tmp86016 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

					tmp86017 := Call(__e, tmp86016, V1472)

					tmp86018 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp86019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp86020 := Call(__e, tmp86019, V1473, Nil)

					tmp86021 := Call(__e, tmp86018, V1474, tmp86020)

					__e.TailApply(tmp86015, V1472, tmp86017, tmp86021)
					return

				}

			}

		}

	}, 3)

	tmp86061 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1loop, tmp86011)

	_ = tmp86061

	tmp86062 := MakeNative(func(__e *ControlFlow) {
		V1476 := __e.Get(1)
		_ = V1476
		tmp86063 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp86064 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86065 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86066 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86067 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86068 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86069 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86070 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp86071 := Call(__e, tmp86070, MakeNumber(93), Nil)

		tmp86072 := Call(__e, tmp86069, MakeNumber(41), tmp86071)

		tmp86073 := Call(__e, tmp86068, MakeNumber(34), tmp86072)

		tmp86074 := Call(__e, tmp86067, MakeNumber(32), tmp86073)

		tmp86075 := Call(__e, tmp86066, MakeNumber(13), tmp86074)

		tmp86076 := Call(__e, tmp86065, MakeNumber(10), tmp86075)

		tmp86077 := Call(__e, tmp86064, MakeNumber(9), tmp86076)

		__e.TailApply(tmp86063, V1476, tmp86077)
		return

	}, 1)

	tmp86078 := Call(__e, PrimNS1Value(symns2_1set), symshen_4terminator_2, tmp86062)

	_ = tmp86078

	tmp86079 := MakeNative(func(__e *ControlFlow) {
		V1478 := __e.Get(1)
		_ = V1478
		tmp86080 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lineread_1loop)

		tmp86081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

		tmp86082 := Call(__e, tmp86081, V1478)

		__e.TailApply(tmp86080, tmp86082, Nil, V1478)
		return

	}, 1)

	tmp86083 := Call(__e, PrimNS1Value(symns2_1set), symlineread, tmp86079)

	_ = tmp86083

	tmp86084 := MakeNative(func(__e *ControlFlow) {
		V1483 := __e.Get(1)
		_ = V1483
		V1484 := __e.Get(2)
		_ = V1484
		V1485 := __e.Get(3)
		_ = V1485
		tmp86141 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp86142 := Call(__e, tmp86141, MakeNumber(-1), V1483)

		if True == tmp86142 {
			tmp86139 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

			tmp86140 := Call(__e, tmp86139, V1484)

			if True == tmp86140 {
				tmp86138 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp86138, MakeString("empty stream"))
				return

			} else {
				tmp86134 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

				tmp86135 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp86136 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

					__e.TailApply(tmp86136, X)
					return

				}, 1)

				tmp86137 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(E)
					return
				}, 1)

				__e.TailApply(tmp86134, tmp86135, V1484, tmp86137)
				return

			}

		} else {
			tmp86129 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86130 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hat)

			tmp86131 := Call(__e, tmp86130)

			tmp86132 := Call(__e, tmp86129, V1483, tmp86131)

			if True == tmp86132 {
				tmp86128 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp86128, MakeString("line read aborted"))
				return

			} else {
				tmp86118 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp86119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp86120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

				tmp86121 := Call(__e, tmp86120)

				tmp86122 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp86123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4carriage_1return)

				tmp86124 := Call(__e, tmp86123)

				tmp86125 := Call(__e, tmp86122, tmp86124, Nil)

				tmp86126 := Call(__e, tmp86119, tmp86121, tmp86125)

				tmp86127 := Call(__e, tmp86118, V1483, tmp86126)

				if True == tmp86127 {
					tmp86095 := MakeNative(func(__e *ControlFlow) {
						Line := __e.Get(1)
						_ = Line
						tmp86096 := MakeNative(func(__e *ControlFlow) {
							It := __e.Get(1)
							_ = It
							tmp86109 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp86110 := Call(__e, tmp86109, Line, symshen_4nextline)

							var ifres86105 Obj

							if True == tmp86110 {
								ifres86105 = True

							} else {
								tmp86107 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

								tmp86108 := Call(__e, tmp86107, Line)

								var ifres86106 Obj

								if True == tmp86108 {
									ifres86106 = True

								} else {
									ifres86106 = False

								}

								ifres86105 = ifres86106

							}

							if True == ifres86105 {
								tmp86098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lineread_1loop)

								tmp86099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

								tmp86100 := Call(__e, tmp86099, V1485)

								tmp86101 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								tmp86102 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp86103 := Call(__e, tmp86102, V1483, Nil)

								tmp86104 := Call(__e, tmp86101, V1484, tmp86103)

								__e.TailApply(tmp86098, tmp86100, tmp86104, V1485)
								return

							} else {
								__e.Return(Line)
								return
							}

						}, 1)

						tmp86111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it)

						tmp86112 := Call(__e, tmp86111, V1484)

						__e.TailApply(tmp86096, tmp86112)
						return

					}, 1)

					tmp86113 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

					tmp86114 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp86115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

						__e.TailApply(tmp86115, X)
						return

					}, 1)

					tmp86116 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(symshen_4nextline)
						return
					}, 1)

					tmp86117 := Call(__e, tmp86113, tmp86114, V1484, tmp86116)

					__e.TailApply(tmp86095, tmp86117)
					return

				} else {
					tmp86088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lineread_1loop)

					tmp86089 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

					tmp86090 := Call(__e, tmp86089, V1485)

					tmp86091 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp86092 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp86093 := Call(__e, tmp86092, V1483, Nil)

					tmp86094 := Call(__e, tmp86091, V1484, tmp86093)

					__e.TailApply(tmp86088, tmp86090, tmp86094, V1485)
					return

				}

			}

		}

	}, 3)

	tmp86143 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lineread_1loop, tmp86084)

	_ = tmp86143

	tmp86144 := MakeNative(func(__e *ControlFlow) {
		V1487 := __e.Get(1)
		_ = V1487
		tmp86145 := MakeNative(func(__e *ControlFlow) {
			TrimLeft := __e.Get(1)
			_ = TrimLeft
			tmp86146 := MakeNative(func(__e *ControlFlow) {
				TrimRight := __e.Get(1)
				_ = TrimRight
				tmp86147 := MakeNative(func(__e *ControlFlow) {
					Trimmed := __e.Get(1)
					_ = Trimmed
					tmp86148 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it_1h)

					__e.TailApply(tmp86148, Trimmed)
					return

				}, 1)

				tmp86149 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

				tmp86150 := Call(__e, tmp86149, TrimRight)

				__e.TailApply(tmp86147, tmp86150)
				return

			}, 1)

			tmp86151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1whitespace)

			tmp86152 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			tmp86153 := Call(__e, tmp86152, TrimLeft)

			tmp86154 := Call(__e, tmp86151, tmp86153)

			__e.TailApply(tmp86146, tmp86154)
			return

		}, 1)

		tmp86155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1whitespace)

		tmp86156 := Call(__e, tmp86155, V1487)

		__e.TailApply(tmp86145, tmp86156)
		return

	}, 1)

	tmp86157 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1it, tmp86144)

	_ = tmp86157

	tmp86158 := MakeNative(func(__e *ControlFlow) {
		V1489 := __e.Get(1)
		_ = V1489
		tmp86177 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86178 := Call(__e, tmp86177, V1489)

		var ifres86163 Obj

		if True == tmp86178 {
			tmp86165 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

			tmp86166 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp86167 := Call(__e, tmp86166, V1489)

			tmp86168 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp86169 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp86170 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp86171 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp86172 := Call(__e, tmp86171, MakeNumber(32), Nil)

			tmp86173 := Call(__e, tmp86170, MakeNumber(13), tmp86172)

			tmp86174 := Call(__e, tmp86169, MakeNumber(10), tmp86173)

			tmp86175 := Call(__e, tmp86168, MakeNumber(9), tmp86174)

			tmp86176 := Call(__e, tmp86165, tmp86167, tmp86175)

			var ifres86164 Obj

			if True == tmp86176 {
				ifres86164 = True

			} else {
				ifres86164 = False

			}

			ifres86163 = ifres86164

		} else {
			ifres86163 = False

		}

		if True == ifres86163 {
			tmp86160 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1whitespace)

			tmp86161 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp86162 := Call(__e, tmp86161, V1489)

			__e.TailApply(tmp86160, tmp86162)
			return

		} else {
			__e.Return(V1489)
			return
		}

	}, 1)

	tmp86179 := Call(__e, PrimNS1Value(symns2_1set), symshen_4trim_1whitespace, tmp86158)

	_ = tmp86179

	tmp86180 := MakeNative(func(__e *ControlFlow) {
		V1491 := __e.Get(1)
		_ = V1491
		tmp86181 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp86182 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cn_1all)

		tmp86183 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp86184 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp86185 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			__e.TailApply(tmp86185, X)
			return

		}, 1)

		tmp86186 := Call(__e, tmp86183, tmp86184, V1491)

		tmp86187 := Call(__e, tmp86182, tmp86186)

		tmp86188 := Call(__e, tmp86181, symshen_4_dit_d, tmp86187)

		_ = tmp86188

		__e.Return(V1491)
		return

	}, 1)

	tmp86189 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1it_1h, tmp86180)

	_ = tmp86189

	tmp86190 := MakeNative(func(__e *ControlFlow) {
		V1493 := __e.Get(1)
		_ = V1493
		tmp86203 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp86204 := Call(__e, tmp86203, Nil, V1493)

		if True == tmp86204 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp86201 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp86202 := Call(__e, tmp86201, V1493)

			if True == tmp86202 {
				tmp86194 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp86195 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86196 := Call(__e, tmp86195, V1493)

				tmp86197 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cn_1all)

				tmp86198 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp86199 := Call(__e, tmp86198, V1493)

				tmp86200 := Call(__e, tmp86197, tmp86199)

				__e.TailApply(tmp86194, tmp86196, tmp86200)
				return

			} else {
				tmp86193 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp86193, symshen_4cn_1all)
				return

			}

		}

	}, 1)

	tmp86205 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cn_1all, tmp86190)

	_ = tmp86205

	tmp86206 := MakeNative(func(__e *ControlFlow) {
		V1495 := __e.Get(1)
		_ = V1495
		tmp86207 := MakeNative(func(__e *ControlFlow) {
			Charlist := __e.Get(1)
			_ = Charlist
			tmp86208 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			tmp86209 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp86210 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

				__e.TailApply(tmp86210, X)
				return

			}, 1)

			tmp86211 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp86212 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1error)

				__e.TailApply(tmp86212, X)
				return

			}, 1)

			__e.TailApply(tmp86208, tmp86209, Charlist, tmp86211)
			return

		}, 1)

		tmp86213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1charlist)

		tmp86214 := Call(__e, tmp86213, V1495)

		__e.TailApply(tmp86207, tmp86214)
		return

	}, 1)

	tmp86215 := Call(__e, PrimNS1Value(symns2_1set), symread_1file, tmp86206)

	_ = tmp86215

	tmp86216 := MakeNative(func(__e *ControlFlow) {
		V1497 := __e.Get(1)
		_ = V1497
		tmp86217 := MakeNative(func(__e *ControlFlow) {
			Ns := __e.Get(1)
			_ = Ns
			tmp86218 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			tmp86219 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp86220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

				__e.TailApply(tmp86220, X)
				return

			}, 1)

			tmp86221 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp86222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1error)

				__e.TailApply(tmp86222, X)
				return

			}, 1)

			__e.TailApply(tmp86218, tmp86219, Ns, tmp86221)
			return

		}, 1)

		tmp86223 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp86224 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp86225 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

			__e.TailApply(tmp86225, X)
			return

		}, 1)

		tmp86226 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

		tmp86227 := Call(__e, tmp86226, V1497)

		tmp86228 := Call(__e, tmp86223, tmp86224, tmp86227)

		__e.TailApply(tmp86217, tmp86228)
		return

	}, 1)

	tmp86229 := Call(__e, PrimNS1Value(symns2_1set), symread_1from_1string, tmp86216)

	_ = tmp86229

	tmp86230 := MakeNative(func(__e *ControlFlow) {
		V1505 := __e.Get(1)
		_ = V1505
		tmp86262 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86263 := Call(__e, tmp86262, V1505)

		var ifres86242 Obj

		if True == tmp86263 {
			tmp86258 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp86259 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp86260 := Call(__e, tmp86259, V1505)

			tmp86261 := Call(__e, tmp86258, tmp86260)

			var ifres86244 Obj

			if True == tmp86261 {
				tmp86254 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp86255 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp86256 := Call(__e, tmp86255, V1505)

				tmp86257 := Call(__e, tmp86254, tmp86256)

				var ifres86246 Obj

				if True == tmp86257 {
					tmp86248 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp86249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp86250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp86251 := Call(__e, tmp86250, V1505)

					tmp86252 := Call(__e, tmp86249, tmp86251)

					tmp86253 := Call(__e, tmp86248, Nil, tmp86252)

					var ifres86247 Obj

					if True == tmp86253 {
						ifres86247 = True

					} else {
						ifres86247 = False

					}

					ifres86246 = ifres86247

				} else {
					ifres86246 = False

				}

				var ifres86245 Obj

				if True == ifres86246 {
					ifres86245 = True

				} else {
					ifres86245 = False

				}

				ifres86244 = ifres86245

			} else {
				ifres86244 = False

			}

			var ifres86243 Obj

			if True == ifres86244 {
				ifres86243 = True

			} else {
				ifres86243 = False

			}

			ifres86242 = ifres86243

		} else {
			ifres86242 = False

		}

		if True == ifres86242 {
			tmp86233 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp86234 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp86235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp86236 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compress_150)

			tmp86237 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp86238 := Call(__e, tmp86237, V1505)

			tmp86239 := Call(__e, tmp86236, MakeNumber(50), tmp86238)

			tmp86240 := Call(__e, tmp86235, tmp86239, MakeString("\n"), symshen_4a)

			tmp86241 := Call(__e, tmp86234, MakeString("read error here:\n\n "), tmp86240)

			__e.TailApply(tmp86233, tmp86241)
			return

		} else {
			tmp86232 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp86232, MakeString("read error\n"))
			return

		}

	}, 1)

	tmp86264 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1error, tmp86230)

	_ = tmp86264

	tmp86265 := MakeNative(func(__e *ControlFlow) {
		V1512 := __e.Get(1)
		_ = V1512
		V1513 := __e.Get(2)
		_ = V1513
		tmp86285 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp86286 := Call(__e, tmp86285, Nil, V1513)

		if True == tmp86286 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp86283 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86284 := Call(__e, tmp86283, MakeNumber(0), V1512)

			if True == tmp86284 {
				__e.Return(MakeString(""))
				return
			} else {
				tmp86281 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp86282 := Call(__e, tmp86281, V1513)

				if True == tmp86282 {
					tmp86270 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp86271 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					tmp86272 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp86273 := Call(__e, tmp86272, V1513)

					tmp86274 := Call(__e, tmp86271, tmp86273)

					tmp86275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compress_150)

					tmp86276 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					tmp86277 := Call(__e, tmp86276, V1512, MakeNumber(1))

					tmp86278 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp86279 := Call(__e, tmp86278, V1513)

					tmp86280 := Call(__e, tmp86275, tmp86277, tmp86279)

					__e.TailApply(tmp86270, tmp86274, tmp86280)
					return

				} else {
					tmp86269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp86269, symshen_4compress_150)
					return

				}

			}

		}

	}, 2)

	tmp86287 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compress_150, tmp86265)

	_ = tmp86287

	tmp86288 := MakeNative(func(__e *ControlFlow) {
		V1515 := __e.Get(1)
		_ = V1515
		tmp86289 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp86787 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86788 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp86789 := Call(__e, tmp86788)

			tmp86790 := Call(__e, tmp86787, YaccParse, tmp86789)

			if True == tmp86790 {
				tmp86291 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp86727 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp86728 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp86729 := Call(__e, tmp86728)

					tmp86730 := Call(__e, tmp86727, YaccParse, tmp86729)

					if True == tmp86730 {
						tmp86293 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp86693 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp86694 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp86695 := Call(__e, tmp86694)

							tmp86696 := Call(__e, tmp86693, YaccParse, tmp86695)

							if True == tmp86696 {
								tmp86295 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									tmp86659 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp86660 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp86661 := Call(__e, tmp86660)

									tmp86662 := Call(__e, tmp86659, YaccParse, tmp86661)

									if True == tmp86662 {
										tmp86297 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											tmp86625 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp86626 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp86627 := Call(__e, tmp86626)

											tmp86628 := Call(__e, tmp86625, YaccParse, tmp86627)

											if True == tmp86628 {
												tmp86299 := MakeNative(func(__e *ControlFlow) {
													YaccParse := __e.Get(1)
													_ = YaccParse
													tmp86591 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp86592 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp86593 := Call(__e, tmp86592)

													tmp86594 := Call(__e, tmp86591, YaccParse, tmp86593)

													if True == tmp86594 {
														tmp86301 := MakeNative(func(__e *ControlFlow) {
															YaccParse := __e.Get(1)
															_ = YaccParse
															tmp86546 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp86547 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp86548 := Call(__e, tmp86547)

															tmp86549 := Call(__e, tmp86546, YaccParse, tmp86548)

															if True == tmp86549 {
																tmp86303 := MakeNative(func(__e *ControlFlow) {
																	YaccParse := __e.Get(1)
																	_ = YaccParse
																	tmp86501 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp86502 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																	tmp86503 := Call(__e, tmp86502)

																	tmp86504 := Call(__e, tmp86501, YaccParse, tmp86503)

																	if True == tmp86504 {
																		tmp86305 := MakeNative(func(__e *ControlFlow) {
																			YaccParse := __e.Get(1)
																			_ = YaccParse
																			tmp86467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp86468 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			tmp86469 := Call(__e, tmp86468)

																			tmp86470 := Call(__e, tmp86467, YaccParse, tmp86469)

																			if True == tmp86470 {
																				tmp86307 := MakeNative(func(__e *ControlFlow) {
																					YaccParse := __e.Get(1)
																					_ = YaccParse
																					tmp86431 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp86432 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					tmp86433 := Call(__e, tmp86432)

																					tmp86434 := Call(__e, tmp86431, YaccParse, tmp86433)

																					if True == tmp86434 {
																						tmp86309 := MakeNative(func(__e *ControlFlow) {
																							YaccParse := __e.Get(1)
																							_ = YaccParse
																							tmp86399 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							tmp86400 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																							tmp86401 := Call(__e, tmp86400)

																							tmp86402 := Call(__e, tmp86399, YaccParse, tmp86401)

																							if True == tmp86402 {
																								tmp86311 := MakeNative(func(__e *ControlFlow) {
																									YaccParse := __e.Get(1)
																									_ = YaccParse
																									tmp86361 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									tmp86362 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																									tmp86363 := Call(__e, tmp86362)

																									tmp86364 := Call(__e, tmp86361, YaccParse, tmp86363)

																									if True == tmp86364 {
																										tmp86313 := MakeNative(func(__e *ControlFlow) {
																											YaccParse := __e.Get(1)
																											_ = YaccParse
																											tmp86329 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											tmp86330 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																											tmp86331 := Call(__e, tmp86330)

																											tmp86332 := Call(__e, tmp86329, YaccParse, tmp86331)

																											if True == tmp86332 {
																												tmp86315 := MakeNative(func(__e *ControlFlow) {
																													Parse___5e_6 := __e.Get(1)
																													_ = Parse___5e_6
																													tmp86321 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																													tmp86322 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																													tmp86323 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																													tmp86324 := Call(__e, tmp86323)

																													tmp86325 := Call(__e, tmp86322, tmp86324, Parse___5e_6)

																													tmp86326 := Call(__e, tmp86321, tmp86325)

																													if True == tmp86326 {
																														tmp86318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																														tmp86319 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														tmp86320 := Call(__e, tmp86319, Parse___5e_6)

																														__e.TailApply(tmp86318, tmp86320, Nil)
																														return

																													} else {
																														tmp86317 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																														__e.TailApply(tmp86317)
																														return

																													}

																												}, 1)

																												tmp86327 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

																												tmp86328 := Call(__e, tmp86327, V1515)

																												__e.TailApply(tmp86315, tmp86328)
																												return

																											} else {
																												__e.Return(YaccParse)
																												return
																											}

																										}, 1)

																										tmp86333 := MakeNative(func(__e *ControlFlow) {
																											Parse__shen_4_5whitespaces_6 := __e.Get(1)
																											_ = Parse__shen_4_5whitespaces_6
																											tmp86352 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																											tmp86353 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											tmp86354 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																											tmp86355 := Call(__e, tmp86354)

																											tmp86356 := Call(__e, tmp86353, tmp86355, Parse__shen_4_5whitespaces_6)

																											tmp86357 := Call(__e, tmp86352, tmp86356)

																											if True == tmp86357 {
																												tmp86336 := MakeNative(func(__e *ControlFlow) {
																													Parse__shen_4_5st__input_6 := __e.Get(1)
																													_ = Parse__shen_4_5st__input_6
																													tmp86344 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																													tmp86345 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																													tmp86346 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																													tmp86347 := Call(__e, tmp86346)

																													tmp86348 := Call(__e, tmp86345, tmp86347, Parse__shen_4_5st__input_6)

																													tmp86349 := Call(__e, tmp86344, tmp86348)

																													if True == tmp86349 {
																														tmp86339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																														tmp86340 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														tmp86341 := Call(__e, tmp86340, Parse__shen_4_5st__input_6)

																														tmp86342 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																														tmp86343 := Call(__e, tmp86342, Parse__shen_4_5st__input_6)

																														__e.TailApply(tmp86339, tmp86341, tmp86343)
																														return

																													} else {
																														tmp86338 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																														__e.TailApply(tmp86338)
																														return

																													}

																												}, 1)

																												tmp86350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																												tmp86351 := Call(__e, tmp86350, Parse__shen_4_5whitespaces_6)

																												__e.TailApply(tmp86336, tmp86351)
																												return

																											} else {
																												tmp86335 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																												__e.TailApply(tmp86335)
																												return

																											}

																										}, 1)

																										tmp86358 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespaces_6)

																										tmp86359 := Call(__e, tmp86358, V1515)

																										tmp86360 := Call(__e, tmp86333, tmp86359)

																										__e.TailApply(tmp86313, tmp86360)
																										return

																									} else {
																										__e.Return(YaccParse)
																										return
																									}

																								}, 1)

																								tmp86365 := MakeNative(func(__e *ControlFlow) {
																									Parse__shen_4_5atom_6 := __e.Get(1)
																									_ = Parse__shen_4_5atom_6
																									tmp86390 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																									tmp86391 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									tmp86392 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																									tmp86393 := Call(__e, tmp86392)

																									tmp86394 := Call(__e, tmp86391, tmp86393, Parse__shen_4_5atom_6)

																									tmp86395 := Call(__e, tmp86390, tmp86394)

																									if True == tmp86395 {
																										tmp86368 := MakeNative(func(__e *ControlFlow) {
																											Parse__shen_4_5st__input_6 := __e.Get(1)
																											_ = Parse__shen_4_5st__input_6
																											tmp86382 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																											tmp86383 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											tmp86384 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																											tmp86385 := Call(__e, tmp86384)

																											tmp86386 := Call(__e, tmp86383, tmp86385, Parse__shen_4_5st__input_6)

																											tmp86387 := Call(__e, tmp86382, tmp86386)

																											if True == tmp86387 {
																												tmp86371 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																												tmp86372 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																												tmp86373 := Call(__e, tmp86372, Parse__shen_4_5st__input_6)

																												tmp86374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																												tmp86375 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

																												tmp86376 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																												tmp86377 := Call(__e, tmp86376, Parse__shen_4_5atom_6)

																												tmp86378 := Call(__e, tmp86375, tmp86377)

																												tmp86379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																												tmp86380 := Call(__e, tmp86379, Parse__shen_4_5st__input_6)

																												tmp86381 := Call(__e, tmp86374, tmp86378, tmp86380)

																												__e.TailApply(tmp86371, tmp86373, tmp86381)
																												return

																											} else {
																												tmp86370 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																												__e.TailApply(tmp86370)
																												return

																											}

																										}, 1)

																										tmp86388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																										tmp86389 := Call(__e, tmp86388, Parse__shen_4_5atom_6)

																										__e.TailApply(tmp86368, tmp86389)
																										return

																									} else {
																										tmp86367 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																										__e.TailApply(tmp86367)
																										return

																									}

																								}, 1)

																								tmp86396 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5atom_6)

																								tmp86397 := Call(__e, tmp86396, V1515)

																								tmp86398 := Call(__e, tmp86365, tmp86397)

																								__e.TailApply(tmp86311, tmp86398)
																								return

																							} else {
																								__e.Return(YaccParse)
																								return
																							}

																						}, 1)

																						tmp86403 := MakeNative(func(__e *ControlFlow) {
																							Parse__shen_4_5comment_6 := __e.Get(1)
																							_ = Parse__shen_4_5comment_6
																							tmp86422 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																							tmp86423 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							tmp86424 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																							tmp86425 := Call(__e, tmp86424)

																							tmp86426 := Call(__e, tmp86423, tmp86425, Parse__shen_4_5comment_6)

																							tmp86427 := Call(__e, tmp86422, tmp86426)

																							if True == tmp86427 {
																								tmp86406 := MakeNative(func(__e *ControlFlow) {
																									Parse__shen_4_5st__input_6 := __e.Get(1)
																									_ = Parse__shen_4_5st__input_6
																									tmp86414 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																									tmp86415 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									tmp86416 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																									tmp86417 := Call(__e, tmp86416)

																									tmp86418 := Call(__e, tmp86415, tmp86417, Parse__shen_4_5st__input_6)

																									tmp86419 := Call(__e, tmp86414, tmp86418)

																									if True == tmp86419 {
																										tmp86409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																										tmp86410 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																										tmp86411 := Call(__e, tmp86410, Parse__shen_4_5st__input_6)

																										tmp86412 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																										tmp86413 := Call(__e, tmp86412, Parse__shen_4_5st__input_6)

																										__e.TailApply(tmp86409, tmp86411, tmp86413)
																										return

																									} else {
																										tmp86408 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																										__e.TailApply(tmp86408)
																										return

																									}

																								}, 1)

																								tmp86420 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																								tmp86421 := Call(__e, tmp86420, Parse__shen_4_5comment_6)

																								__e.TailApply(tmp86406, tmp86421)
																								return

																							} else {
																								tmp86405 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																								__e.TailApply(tmp86405)
																								return

																							}

																						}, 1)

																						tmp86428 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comment_6)

																						tmp86429 := Call(__e, tmp86428, V1515)

																						tmp86430 := Call(__e, tmp86403, tmp86429)

																						__e.TailApply(tmp86309, tmp86430)
																						return

																					} else {
																						__e.Return(YaccParse)
																						return
																					}

																				}, 1)

																				tmp86435 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5comma_6 := __e.Get(1)
																					_ = Parse__shen_4_5comma_6
																					tmp86458 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																					tmp86459 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp86460 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					tmp86461 := Call(__e, tmp86460)

																					tmp86462 := Call(__e, tmp86459, tmp86461, Parse__shen_4_5comma_6)

																					tmp86463 := Call(__e, tmp86458, tmp86462)

																					if True == tmp86463 {
																						tmp86438 := MakeNative(func(__e *ControlFlow) {
																							Parse__shen_4_5st__input_6 := __e.Get(1)
																							_ = Parse__shen_4_5st__input_6
																							tmp86450 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																							tmp86451 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							tmp86452 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																							tmp86453 := Call(__e, tmp86452)

																							tmp86454 := Call(__e, tmp86451, tmp86453, Parse__shen_4_5st__input_6)

																							tmp86455 := Call(__e, tmp86450, tmp86454)

																							if True == tmp86455 {
																								tmp86441 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																								tmp86442 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																								tmp86443 := Call(__e, tmp86442, Parse__shen_4_5st__input_6)

																								tmp86444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																								tmp86445 := Call(__e, PrimNS1Value(symns2_1value), symintern)

																								tmp86446 := Call(__e, tmp86445, MakeString(","))

																								tmp86447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																								tmp86448 := Call(__e, tmp86447, Parse__shen_4_5st__input_6)

																								tmp86449 := Call(__e, tmp86444, tmp86446, tmp86448)

																								__e.TailApply(tmp86441, tmp86443, tmp86449)
																								return

																							} else {
																								tmp86440 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																								__e.TailApply(tmp86440)
																								return

																							}

																						}, 1)

																						tmp86456 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																						tmp86457 := Call(__e, tmp86456, Parse__shen_4_5comma_6)

																						__e.TailApply(tmp86438, tmp86457)
																						return

																					} else {
																						tmp86437 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																						__e.TailApply(tmp86437)
																						return

																					}

																				}, 1)

																				tmp86464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comma_6)

																				tmp86465 := Call(__e, tmp86464, V1515)

																				tmp86466 := Call(__e, tmp86435, tmp86465)

																				__e.TailApply(tmp86307, tmp86466)
																				return

																			} else {
																				__e.Return(YaccParse)
																				return
																			}

																		}, 1)

																		tmp86471 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5colon_6 := __e.Get(1)
																			_ = Parse__shen_4_5colon_6
																			tmp86492 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																			tmp86493 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp86494 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			tmp86495 := Call(__e, tmp86494)

																			tmp86496 := Call(__e, tmp86493, tmp86495, Parse__shen_4_5colon_6)

																			tmp86497 := Call(__e, tmp86492, tmp86496)

																			if True == tmp86497 {
																				tmp86474 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5st__input_6 := __e.Get(1)
																					_ = Parse__shen_4_5st__input_6
																					tmp86484 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																					tmp86485 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp86486 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					tmp86487 := Call(__e, tmp86486)

																					tmp86488 := Call(__e, tmp86485, tmp86487, Parse__shen_4_5st__input_6)

																					tmp86489 := Call(__e, tmp86484, tmp86488)

																					if True == tmp86489 {
																						tmp86477 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																						tmp86478 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																						tmp86479 := Call(__e, tmp86478, Parse__shen_4_5st__input_6)

																						tmp86480 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						tmp86481 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																						tmp86482 := Call(__e, tmp86481, Parse__shen_4_5st__input_6)

																						tmp86483 := Call(__e, tmp86480, sym_h, tmp86482)

																						__e.TailApply(tmp86477, tmp86479, tmp86483)
																						return

																					} else {
																						tmp86476 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																						__e.TailApply(tmp86476)
																						return

																					}

																				}, 1)

																				tmp86490 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																				tmp86491 := Call(__e, tmp86490, Parse__shen_4_5colon_6)

																				__e.TailApply(tmp86474, tmp86491)
																				return

																			} else {
																				tmp86473 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																				__e.TailApply(tmp86473)
																				return

																			}

																		}, 1)

																		tmp86498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5colon_6)

																		tmp86499 := Call(__e, tmp86498, V1515)

																		tmp86500 := Call(__e, tmp86471, tmp86499)

																		__e.TailApply(tmp86305, tmp86500)
																		return

																	} else {
																		__e.Return(YaccParse)
																		return
																	}

																}, 1)

																tmp86505 := MakeNative(func(__e *ControlFlow) {
																	Parse__shen_4_5colon_6 := __e.Get(1)
																	_ = Parse__shen_4_5colon_6
																	tmp86537 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																	tmp86538 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp86539 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																	tmp86540 := Call(__e, tmp86539)

																	tmp86541 := Call(__e, tmp86538, tmp86540, Parse__shen_4_5colon_6)

																	tmp86542 := Call(__e, tmp86537, tmp86541)

																	if True == tmp86542 {
																		tmp86508 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5minus_6 := __e.Get(1)
																			_ = Parse__shen_4_5minus_6
																			tmp86529 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																			tmp86530 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp86531 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			tmp86532 := Call(__e, tmp86531)

																			tmp86533 := Call(__e, tmp86530, tmp86532, Parse__shen_4_5minus_6)

																			tmp86534 := Call(__e, tmp86529, tmp86533)

																			if True == tmp86534 {
																				tmp86511 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5st__input_6 := __e.Get(1)
																					_ = Parse__shen_4_5st__input_6
																					tmp86521 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																					tmp86522 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp86523 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					tmp86524 := Call(__e, tmp86523)

																					tmp86525 := Call(__e, tmp86522, tmp86524, Parse__shen_4_5st__input_6)

																					tmp86526 := Call(__e, tmp86521, tmp86525)

																					if True == tmp86526 {
																						tmp86514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																						tmp86515 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																						tmp86516 := Call(__e, tmp86515, Parse__shen_4_5st__input_6)

																						tmp86517 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						tmp86518 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																						tmp86519 := Call(__e, tmp86518, Parse__shen_4_5st__input_6)

																						tmp86520 := Call(__e, tmp86517, sym_h_1, tmp86519)

																						__e.TailApply(tmp86514, tmp86516, tmp86520)
																						return

																					} else {
																						tmp86513 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																						__e.TailApply(tmp86513)
																						return

																					}

																				}, 1)

																				tmp86527 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																				tmp86528 := Call(__e, tmp86527, Parse__shen_4_5minus_6)

																				__e.TailApply(tmp86511, tmp86528)
																				return

																			} else {
																				tmp86510 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																				__e.TailApply(tmp86510)
																				return

																			}

																		}, 1)

																		tmp86535 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5minus_6)

																		tmp86536 := Call(__e, tmp86535, Parse__shen_4_5colon_6)

																		__e.TailApply(tmp86508, tmp86536)
																		return

																	} else {
																		tmp86507 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																		__e.TailApply(tmp86507)
																		return

																	}

																}, 1)

																tmp86543 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5colon_6)

																tmp86544 := Call(__e, tmp86543, V1515)

																tmp86545 := Call(__e, tmp86505, tmp86544)

																__e.TailApply(tmp86303, tmp86545)
																return

															} else {
																__e.Return(YaccParse)
																return
															}

														}, 1)

														tmp86550 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5colon_6 := __e.Get(1)
															_ = Parse__shen_4_5colon_6
															tmp86582 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp86583 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp86584 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp86585 := Call(__e, tmp86584)

															tmp86586 := Call(__e, tmp86583, tmp86585, Parse__shen_4_5colon_6)

															tmp86587 := Call(__e, tmp86582, tmp86586)

															if True == tmp86587 {
																tmp86553 := MakeNative(func(__e *ControlFlow) {
																	Parse__shen_4_5equal_6 := __e.Get(1)
																	_ = Parse__shen_4_5equal_6
																	tmp86574 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																	tmp86575 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp86576 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																	tmp86577 := Call(__e, tmp86576)

																	tmp86578 := Call(__e, tmp86575, tmp86577, Parse__shen_4_5equal_6)

																	tmp86579 := Call(__e, tmp86574, tmp86578)

																	if True == tmp86579 {
																		tmp86556 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5st__input_6 := __e.Get(1)
																			_ = Parse__shen_4_5st__input_6
																			tmp86566 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																			tmp86567 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp86568 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			tmp86569 := Call(__e, tmp86568)

																			tmp86570 := Call(__e, tmp86567, tmp86569, Parse__shen_4_5st__input_6)

																			tmp86571 := Call(__e, tmp86566, tmp86570)

																			if True == tmp86571 {
																				tmp86559 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																				tmp86560 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																				tmp86561 := Call(__e, tmp86560, Parse__shen_4_5st__input_6)

																				tmp86562 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																				tmp86563 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																				tmp86564 := Call(__e, tmp86563, Parse__shen_4_5st__input_6)

																				tmp86565 := Call(__e, tmp86562, sym_h_a, tmp86564)

																				__e.TailApply(tmp86559, tmp86561, tmp86565)
																				return

																			} else {
																				tmp86558 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																				__e.TailApply(tmp86558)
																				return

																			}

																		}, 1)

																		tmp86572 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																		tmp86573 := Call(__e, tmp86572, Parse__shen_4_5equal_6)

																		__e.TailApply(tmp86556, tmp86573)
																		return

																	} else {
																		tmp86555 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																		__e.TailApply(tmp86555)
																		return

																	}

																}, 1)

																tmp86580 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5equal_6)

																tmp86581 := Call(__e, tmp86580, Parse__shen_4_5colon_6)

																__e.TailApply(tmp86553, tmp86581)
																return

															} else {
																tmp86552 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(tmp86552)
																return

															}

														}, 1)

														tmp86588 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5colon_6)

														tmp86589 := Call(__e, tmp86588, V1515)

														tmp86590 := Call(__e, tmp86550, tmp86589)

														__e.TailApply(tmp86301, tmp86590)
														return

													} else {
														__e.Return(YaccParse)
														return
													}

												}, 1)

												tmp86595 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5semicolon_6 := __e.Get(1)
													_ = Parse__shen_4_5semicolon_6
													tmp86616 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp86617 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp86618 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp86619 := Call(__e, tmp86618)

													tmp86620 := Call(__e, tmp86617, tmp86619, Parse__shen_4_5semicolon_6)

													tmp86621 := Call(__e, tmp86616, tmp86620)

													if True == tmp86621 {
														tmp86598 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5st__input_6 := __e.Get(1)
															_ = Parse__shen_4_5st__input_6
															tmp86608 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp86609 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp86610 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp86611 := Call(__e, tmp86610)

															tmp86612 := Call(__e, tmp86609, tmp86611, Parse__shen_4_5st__input_6)

															tmp86613 := Call(__e, tmp86608, tmp86612)

															if True == tmp86613 {
																tmp86601 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																tmp86602 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp86603 := Call(__e, tmp86602, Parse__shen_4_5st__input_6)

																tmp86604 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp86605 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp86606 := Call(__e, tmp86605, Parse__shen_4_5st__input_6)

																tmp86607 := Call(__e, tmp86604, sym_k, tmp86606)

																__e.TailApply(tmp86601, tmp86603, tmp86607)
																return

															} else {
																tmp86600 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(tmp86600)
																return

															}

														}, 1)

														tmp86614 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

														tmp86615 := Call(__e, tmp86614, Parse__shen_4_5semicolon_6)

														__e.TailApply(tmp86598, tmp86615)
														return

													} else {
														tmp86597 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp86597)
														return

													}

												}, 1)

												tmp86622 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_6)

												tmp86623 := Call(__e, tmp86622, V1515)

												tmp86624 := Call(__e, tmp86595, tmp86623)

												__e.TailApply(tmp86299, tmp86624)
												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										tmp86629 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5bar_6 := __e.Get(1)
											_ = Parse__shen_4_5bar_6
											tmp86650 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp86651 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp86652 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp86653 := Call(__e, tmp86652)

											tmp86654 := Call(__e, tmp86651, tmp86653, Parse__shen_4_5bar_6)

											tmp86655 := Call(__e, tmp86650, tmp86654)

											if True == tmp86655 {
												tmp86632 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5st__input_6 := __e.Get(1)
													_ = Parse__shen_4_5st__input_6
													tmp86642 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp86643 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp86644 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp86645 := Call(__e, tmp86644)

													tmp86646 := Call(__e, tmp86643, tmp86645, Parse__shen_4_5st__input_6)

													tmp86647 := Call(__e, tmp86642, tmp86646)

													if True == tmp86647 {
														tmp86635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														tmp86636 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp86637 := Call(__e, tmp86636, Parse__shen_4_5st__input_6)

														tmp86638 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp86639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp86640 := Call(__e, tmp86639, Parse__shen_4_5st__input_6)

														tmp86641 := Call(__e, tmp86638, symbar_b, tmp86640)

														__e.TailApply(tmp86635, tmp86637, tmp86641)
														return

													} else {
														tmp86634 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp86634)
														return

													}

												}, 1)

												tmp86648 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

												tmp86649 := Call(__e, tmp86648, Parse__shen_4_5bar_6)

												__e.TailApply(tmp86632, tmp86649)
												return

											} else {
												tmp86631 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp86631)
												return

											}

										}, 1)

										tmp86656 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5bar_6)

										tmp86657 := Call(__e, tmp86656, V1515)

										tmp86658 := Call(__e, tmp86629, tmp86657)

										__e.TailApply(tmp86297, tmp86658)
										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								tmp86663 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5rcurly_6 := __e.Get(1)
									_ = Parse__shen_4_5rcurly_6
									tmp86684 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp86685 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp86686 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp86687 := Call(__e, tmp86686)

									tmp86688 := Call(__e, tmp86685, tmp86687, Parse__shen_4_5rcurly_6)

									tmp86689 := Call(__e, tmp86684, tmp86688)

									if True == tmp86689 {
										tmp86666 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5st__input_6 := __e.Get(1)
											_ = Parse__shen_4_5st__input_6
											tmp86676 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp86677 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp86678 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp86679 := Call(__e, tmp86678)

											tmp86680 := Call(__e, tmp86677, tmp86679, Parse__shen_4_5st__input_6)

											tmp86681 := Call(__e, tmp86676, tmp86680)

											if True == tmp86681 {
												tmp86669 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp86670 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp86671 := Call(__e, tmp86670, Parse__shen_4_5st__input_6)

												tmp86672 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp86673 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp86674 := Call(__e, tmp86673, Parse__shen_4_5st__input_6)

												tmp86675 := Call(__e, tmp86672, sym_j, tmp86674)

												__e.TailApply(tmp86669, tmp86671, tmp86675)
												return

											} else {
												tmp86668 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp86668)
												return

											}

										}, 1)

										tmp86682 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

										tmp86683 := Call(__e, tmp86682, Parse__shen_4_5rcurly_6)

										__e.TailApply(tmp86666, tmp86683)
										return

									} else {
										tmp86665 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp86665)
										return

									}

								}, 1)

								tmp86690 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rcurly_6)

								tmp86691 := Call(__e, tmp86690, V1515)

								tmp86692 := Call(__e, tmp86663, tmp86691)

								__e.TailApply(tmp86295, tmp86692)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp86697 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5lcurly_6 := __e.Get(1)
							_ = Parse__shen_4_5lcurly_6
							tmp86718 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp86719 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp86720 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp86721 := Call(__e, tmp86720)

							tmp86722 := Call(__e, tmp86719, tmp86721, Parse__shen_4_5lcurly_6)

							tmp86723 := Call(__e, tmp86718, tmp86722)

							if True == tmp86723 {
								tmp86700 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5st__input_6 := __e.Get(1)
									_ = Parse__shen_4_5st__input_6
									tmp86710 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp86711 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp86712 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp86713 := Call(__e, tmp86712)

									tmp86714 := Call(__e, tmp86711, tmp86713, Parse__shen_4_5st__input_6)

									tmp86715 := Call(__e, tmp86710, tmp86714)

									if True == tmp86715 {
										tmp86703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp86704 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp86705 := Call(__e, tmp86704, Parse__shen_4_5st__input_6)

										tmp86706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp86707 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp86708 := Call(__e, tmp86707, Parse__shen_4_5st__input_6)

										tmp86709 := Call(__e, tmp86706, sym_i, tmp86708)

										__e.TailApply(tmp86703, tmp86705, tmp86709)
										return

									} else {
										tmp86702 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp86702)
										return

									}

								}, 1)

								tmp86716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

								tmp86717 := Call(__e, tmp86716, Parse__shen_4_5lcurly_6)

								__e.TailApply(tmp86700, tmp86717)
								return

							} else {
								tmp86699 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp86699)
								return

							}

						}, 1)

						tmp86724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5lcurly_6)

						tmp86725 := Call(__e, tmp86724, V1515)

						tmp86726 := Call(__e, tmp86697, tmp86725)

						__e.TailApply(tmp86293, tmp86726)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp86731 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5lrb_6 := __e.Get(1)
					_ = Parse__shen_4_5lrb_6
					tmp86778 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp86779 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp86780 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp86781 := Call(__e, tmp86780)

					tmp86782 := Call(__e, tmp86779, tmp86781, Parse__shen_4_5lrb_6)

					tmp86783 := Call(__e, tmp86778, tmp86782)

					if True == tmp86783 {
						tmp86734 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5st__input1_6 := __e.Get(1)
							_ = Parse__shen_4_5st__input1_6
							tmp86770 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp86771 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp86772 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp86773 := Call(__e, tmp86772)

							tmp86774 := Call(__e, tmp86771, tmp86773, Parse__shen_4_5st__input1_6)

							tmp86775 := Call(__e, tmp86770, tmp86774)

							if True == tmp86775 {
								tmp86737 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5rrb_6 := __e.Get(1)
									_ = Parse__shen_4_5rrb_6
									tmp86762 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp86763 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp86764 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp86765 := Call(__e, tmp86764)

									tmp86766 := Call(__e, tmp86763, tmp86765, Parse__shen_4_5rrb_6)

									tmp86767 := Call(__e, tmp86762, tmp86766)

									if True == tmp86767 {
										tmp86740 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5st__input2_6 := __e.Get(1)
											_ = Parse__shen_4_5st__input2_6
											tmp86754 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp86755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp86756 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp86757 := Call(__e, tmp86756)

											tmp86758 := Call(__e, tmp86755, tmp86757, Parse__shen_4_5st__input2_6)

											tmp86759 := Call(__e, tmp86754, tmp86758)

											if True == tmp86759 {
												tmp86743 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp86744 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp86745 := Call(__e, tmp86744, Parse__shen_4_5st__input2_6)

												tmp86746 := Call(__e, PrimNS1Value(symns2_1value), symshen_4package_1macro)

												tmp86747 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

												tmp86748 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp86749 := Call(__e, tmp86748, Parse__shen_4_5st__input1_6)

												tmp86750 := Call(__e, tmp86747, tmp86749)

												tmp86751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp86752 := Call(__e, tmp86751, Parse__shen_4_5st__input2_6)

												tmp86753 := Call(__e, tmp86746, tmp86750, tmp86752)

												__e.TailApply(tmp86743, tmp86745, tmp86753)
												return

											} else {
												tmp86742 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp86742)
												return

											}

										}, 1)

										tmp86760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input2_6)

										tmp86761 := Call(__e, tmp86760, Parse__shen_4_5rrb_6)

										__e.TailApply(tmp86740, tmp86761)
										return

									} else {
										tmp86739 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp86739)
										return

									}

								}, 1)

								tmp86768 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rrb_6)

								tmp86769 := Call(__e, tmp86768, Parse__shen_4_5st__input1_6)

								__e.TailApply(tmp86737, tmp86769)
								return

							} else {
								tmp86736 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp86736)
								return

							}

						}, 1)

						tmp86776 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input1_6)

						tmp86777 := Call(__e, tmp86776, Parse__shen_4_5lrb_6)

						__e.TailApply(tmp86734, tmp86777)
						return

					} else {
						tmp86733 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp86733)
						return

					}

				}, 1)

				tmp86784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5lrb_6)

				tmp86785 := Call(__e, tmp86784, V1515)

				tmp86786 := Call(__e, tmp86731, tmp86785)

				__e.TailApply(tmp86291, tmp86786)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp86791 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5lsb_6 := __e.Get(1)
			_ = Parse__shen_4_5lsb_6
			tmp86840 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp86841 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86842 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp86843 := Call(__e, tmp86842)

			tmp86844 := Call(__e, tmp86841, tmp86843, Parse__shen_4_5lsb_6)

			tmp86845 := Call(__e, tmp86840, tmp86844)

			if True == tmp86845 {
				tmp86794 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5st__input1_6 := __e.Get(1)
					_ = Parse__shen_4_5st__input1_6
					tmp86832 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp86833 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp86834 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp86835 := Call(__e, tmp86834)

					tmp86836 := Call(__e, tmp86833, tmp86835, Parse__shen_4_5st__input1_6)

					tmp86837 := Call(__e, tmp86832, tmp86836)

					if True == tmp86837 {
						tmp86797 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rsb_6 := __e.Get(1)
							_ = Parse__shen_4_5rsb_6
							tmp86824 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp86825 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp86826 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp86827 := Call(__e, tmp86826)

							tmp86828 := Call(__e, tmp86825, tmp86827, Parse__shen_4_5rsb_6)

							tmp86829 := Call(__e, tmp86824, tmp86828)

							if True == tmp86829 {
								tmp86800 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5st__input2_6 := __e.Get(1)
									_ = Parse__shen_4_5st__input2_6
									tmp86816 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp86817 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp86818 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp86819 := Call(__e, tmp86818)

									tmp86820 := Call(__e, tmp86817, tmp86819, Parse__shen_4_5st__input2_6)

									tmp86821 := Call(__e, tmp86816, tmp86820)

									if True == tmp86821 {
										tmp86803 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp86804 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp86805 := Call(__e, tmp86804, Parse__shen_4_5st__input2_6)

										tmp86806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp86807 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

										tmp86808 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cons__form)

										tmp86809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp86810 := Call(__e, tmp86809, Parse__shen_4_5st__input1_6)

										tmp86811 := Call(__e, tmp86808, tmp86810)

										tmp86812 := Call(__e, tmp86807, tmp86811)

										tmp86813 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp86814 := Call(__e, tmp86813, Parse__shen_4_5st__input2_6)

										tmp86815 := Call(__e, tmp86806, tmp86812, tmp86814)

										__e.TailApply(tmp86803, tmp86805, tmp86815)
										return

									} else {
										tmp86802 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp86802)
										return

									}

								}, 1)

								tmp86822 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input2_6)

								tmp86823 := Call(__e, tmp86822, Parse__shen_4_5rsb_6)

								__e.TailApply(tmp86800, tmp86823)
								return

							} else {
								tmp86799 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp86799)
								return

							}

						}, 1)

						tmp86830 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rsb_6)

						tmp86831 := Call(__e, tmp86830, Parse__shen_4_5st__input1_6)

						__e.TailApply(tmp86797, tmp86831)
						return

					} else {
						tmp86796 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp86796)
						return

					}

				}, 1)

				tmp86838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input1_6)

				tmp86839 := Call(__e, tmp86838, Parse__shen_4_5lsb_6)

				__e.TailApply(tmp86794, tmp86839)
				return

			} else {
				tmp86793 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp86793)
				return

			}

		}, 1)

		tmp86846 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5lsb_6)

		tmp86847 := Call(__e, tmp86846, V1515)

		tmp86848 := Call(__e, tmp86791, tmp86847)

		__e.TailApply(tmp86289, tmp86848)
		return

	}, 1)

	tmp86849 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input_6, tmp86288)

	_ = tmp86849

	tmp86850 := MakeNative(func(__e *ControlFlow) {
		V1518 := __e.Get(1)
		_ = V1518
		tmp86869 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86870 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86871 := Call(__e, tmp86870, V1518)

		tmp86872 := Call(__e, tmp86869, tmp86871)

		var ifres86863 Obj

		if True == tmp86872 {
			tmp86865 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp86867 := Call(__e, tmp86866, V1518)

			tmp86868 := Call(__e, tmp86865, MakeNumber(91), tmp86867)

			var ifres86864 Obj

			if True == tmp86868 {
				ifres86864 = True

			} else {
				ifres86864 = False

			}

			ifres86863 = ifres86864

		} else {
			ifres86863 = False

		}

		if True == ifres86863 {
			tmp86853 := MakeNative(func(__e *ControlFlow) {
				NewStream1516 := __e.Get(1)
				_ = NewStream1516
				tmp86854 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86855 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86856 := Call(__e, tmp86855, NewStream1516)

				__e.TailApply(tmp86854, tmp86856, symshen_4skip)
				return

			}, 1)

			tmp86857 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp86858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp86859 := Call(__e, tmp86858, V1518)

			tmp86860 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp86861 := Call(__e, tmp86860, V1518)

			tmp86862 := Call(__e, tmp86857, tmp86859, tmp86861)

			__e.TailApply(tmp86853, tmp86862)
			return

		} else {
			tmp86852 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86852)
			return

		}

	}, 1)

	tmp86873 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lsb_6, tmp86850)

	_ = tmp86873

	tmp86874 := MakeNative(func(__e *ControlFlow) {
		V1521 := __e.Get(1)
		_ = V1521
		tmp86893 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86894 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86895 := Call(__e, tmp86894, V1521)

		tmp86896 := Call(__e, tmp86893, tmp86895)

		var ifres86887 Obj

		if True == tmp86896 {
			tmp86889 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86890 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp86891 := Call(__e, tmp86890, V1521)

			tmp86892 := Call(__e, tmp86889, MakeNumber(93), tmp86891)

			var ifres86888 Obj

			if True == tmp86892 {
				ifres86888 = True

			} else {
				ifres86888 = False

			}

			ifres86887 = ifres86888

		} else {
			ifres86887 = False

		}

		if True == ifres86887 {
			tmp86877 := MakeNative(func(__e *ControlFlow) {
				NewStream1519 := __e.Get(1)
				_ = NewStream1519
				tmp86878 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86879 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86880 := Call(__e, tmp86879, NewStream1519)

				__e.TailApply(tmp86878, tmp86880, symshen_4skip)
				return

			}, 1)

			tmp86881 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp86882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp86883 := Call(__e, tmp86882, V1521)

			tmp86884 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp86885 := Call(__e, tmp86884, V1521)

			tmp86886 := Call(__e, tmp86881, tmp86883, tmp86885)

			__e.TailApply(tmp86877, tmp86886)
			return

		} else {
			tmp86876 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86876)
			return

		}

	}, 1)

	tmp86897 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rsb_6, tmp86874)

	_ = tmp86897

	tmp86898 := MakeNative(func(__e *ControlFlow) {
		V1524 := __e.Get(1)
		_ = V1524
		tmp86917 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86918 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86919 := Call(__e, tmp86918, V1524)

		tmp86920 := Call(__e, tmp86917, tmp86919)

		var ifres86911 Obj

		if True == tmp86920 {
			tmp86913 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp86915 := Call(__e, tmp86914, V1524)

			tmp86916 := Call(__e, tmp86913, MakeNumber(123), tmp86915)

			var ifres86912 Obj

			if True == tmp86916 {
				ifres86912 = True

			} else {
				ifres86912 = False

			}

			ifres86911 = ifres86912

		} else {
			ifres86911 = False

		}

		if True == ifres86911 {
			tmp86901 := MakeNative(func(__e *ControlFlow) {
				NewStream1522 := __e.Get(1)
				_ = NewStream1522
				tmp86902 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86903 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86904 := Call(__e, tmp86903, NewStream1522)

				__e.TailApply(tmp86902, tmp86904, symshen_4skip)
				return

			}, 1)

			tmp86905 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp86906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp86907 := Call(__e, tmp86906, V1524)

			tmp86908 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp86909 := Call(__e, tmp86908, V1524)

			tmp86910 := Call(__e, tmp86905, tmp86907, tmp86909)

			__e.TailApply(tmp86901, tmp86910)
			return

		} else {
			tmp86900 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86900)
			return

		}

	}, 1)

	tmp86921 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lcurly_6, tmp86898)

	_ = tmp86921

	tmp86922 := MakeNative(func(__e *ControlFlow) {
		V1527 := __e.Get(1)
		_ = V1527
		tmp86941 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86942 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86943 := Call(__e, tmp86942, V1527)

		tmp86944 := Call(__e, tmp86941, tmp86943)

		var ifres86935 Obj

		if True == tmp86944 {
			tmp86937 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86938 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp86939 := Call(__e, tmp86938, V1527)

			tmp86940 := Call(__e, tmp86937, MakeNumber(125), tmp86939)

			var ifres86936 Obj

			if True == tmp86940 {
				ifres86936 = True

			} else {
				ifres86936 = False

			}

			ifres86935 = ifres86936

		} else {
			ifres86935 = False

		}

		if True == ifres86935 {
			tmp86925 := MakeNative(func(__e *ControlFlow) {
				NewStream1525 := __e.Get(1)
				_ = NewStream1525
				tmp86926 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86927 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86928 := Call(__e, tmp86927, NewStream1525)

				__e.TailApply(tmp86926, tmp86928, symshen_4skip)
				return

			}, 1)

			tmp86929 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp86930 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp86931 := Call(__e, tmp86930, V1527)

			tmp86932 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp86933 := Call(__e, tmp86932, V1527)

			tmp86934 := Call(__e, tmp86929, tmp86931, tmp86933)

			__e.TailApply(tmp86925, tmp86934)
			return

		} else {
			tmp86924 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86924)
			return

		}

	}, 1)

	tmp86945 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rcurly_6, tmp86922)

	_ = tmp86945

	tmp86946 := MakeNative(func(__e *ControlFlow) {
		V1530 := __e.Get(1)
		_ = V1530
		tmp86965 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86967 := Call(__e, tmp86966, V1530)

		tmp86968 := Call(__e, tmp86965, tmp86967)

		var ifres86959 Obj

		if True == tmp86968 {
			tmp86961 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86962 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp86963 := Call(__e, tmp86962, V1530)

			tmp86964 := Call(__e, tmp86961, MakeNumber(124), tmp86963)

			var ifres86960 Obj

			if True == tmp86964 {
				ifres86960 = True

			} else {
				ifres86960 = False

			}

			ifres86959 = ifres86960

		} else {
			ifres86959 = False

		}

		if True == ifres86959 {
			tmp86949 := MakeNative(func(__e *ControlFlow) {
				NewStream1528 := __e.Get(1)
				_ = NewStream1528
				tmp86950 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86951 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86952 := Call(__e, tmp86951, NewStream1528)

				__e.TailApply(tmp86950, tmp86952, symshen_4skip)
				return

			}, 1)

			tmp86953 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp86954 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp86955 := Call(__e, tmp86954, V1530)

			tmp86956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp86957 := Call(__e, tmp86956, V1530)

			tmp86958 := Call(__e, tmp86953, tmp86955, tmp86957)

			__e.TailApply(tmp86949, tmp86958)
			return

		} else {
			tmp86948 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86948)
			return

		}

	}, 1)

	tmp86969 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5bar_6, tmp86946)

	_ = tmp86969

	tmp86970 := MakeNative(func(__e *ControlFlow) {
		V1533 := __e.Get(1)
		_ = V1533
		tmp86989 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp86990 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp86991 := Call(__e, tmp86990, V1533)

		tmp86992 := Call(__e, tmp86989, tmp86991)

		var ifres86983 Obj

		if True == tmp86992 {
			tmp86985 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp86986 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp86987 := Call(__e, tmp86986, V1533)

			tmp86988 := Call(__e, tmp86985, MakeNumber(59), tmp86987)

			var ifres86984 Obj

			if True == tmp86988 {
				ifres86984 = True

			} else {
				ifres86984 = False

			}

			ifres86983 = ifres86984

		} else {
			ifres86983 = False

		}

		if True == ifres86983 {
			tmp86973 := MakeNative(func(__e *ControlFlow) {
				NewStream1531 := __e.Get(1)
				_ = NewStream1531
				tmp86974 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86975 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp86976 := Call(__e, tmp86975, NewStream1531)

				__e.TailApply(tmp86974, tmp86976, symshen_4skip)
				return

			}, 1)

			tmp86977 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp86978 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp86979 := Call(__e, tmp86978, V1533)

			tmp86980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp86981 := Call(__e, tmp86980, V1533)

			tmp86982 := Call(__e, tmp86977, tmp86979, tmp86981)

			__e.TailApply(tmp86973, tmp86982)
			return

		} else {
			tmp86972 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86972)
			return

		}

	}, 1)

	tmp86993 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5semicolon_6, tmp86970)

	_ = tmp86993

	tmp86994 := MakeNative(func(__e *ControlFlow) {
		V1536 := __e.Get(1)
		_ = V1536
		tmp87013 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87014 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87015 := Call(__e, tmp87014, V1536)

		tmp87016 := Call(__e, tmp87013, tmp87015)

		var ifres87007 Obj

		if True == tmp87016 {
			tmp87009 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87010 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87011 := Call(__e, tmp87010, V1536)

			tmp87012 := Call(__e, tmp87009, MakeNumber(58), tmp87011)

			var ifres87008 Obj

			if True == tmp87012 {
				ifres87008 = True

			} else {
				ifres87008 = False

			}

			ifres87007 = ifres87008

		} else {
			ifres87007 = False

		}

		if True == ifres87007 {
			tmp86997 := MakeNative(func(__e *ControlFlow) {
				NewStream1534 := __e.Get(1)
				_ = NewStream1534
				tmp86998 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp86999 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87000 := Call(__e, tmp86999, NewStream1534)

				__e.TailApply(tmp86998, tmp87000, symshen_4skip)
				return

			}, 1)

			tmp87001 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp87002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp87003 := Call(__e, tmp87002, V1536)

			tmp87004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp87005 := Call(__e, tmp87004, V1536)

			tmp87006 := Call(__e, tmp87001, tmp87003, tmp87005)

			__e.TailApply(tmp86997, tmp87006)
			return

		} else {
			tmp86996 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp86996)
			return

		}

	}, 1)

	tmp87017 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5colon_6, tmp86994)

	_ = tmp87017

	tmp87018 := MakeNative(func(__e *ControlFlow) {
		V1539 := __e.Get(1)
		_ = V1539
		tmp87037 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87038 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87039 := Call(__e, tmp87038, V1539)

		tmp87040 := Call(__e, tmp87037, tmp87039)

		var ifres87031 Obj

		if True == tmp87040 {
			tmp87033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87034 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87035 := Call(__e, tmp87034, V1539)

			tmp87036 := Call(__e, tmp87033, MakeNumber(44), tmp87035)

			var ifres87032 Obj

			if True == tmp87036 {
				ifres87032 = True

			} else {
				ifres87032 = False

			}

			ifres87031 = ifres87032

		} else {
			ifres87031 = False

		}

		if True == ifres87031 {
			tmp87021 := MakeNative(func(__e *ControlFlow) {
				NewStream1537 := __e.Get(1)
				_ = NewStream1537
				tmp87022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87023 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87024 := Call(__e, tmp87023, NewStream1537)

				__e.TailApply(tmp87022, tmp87024, symshen_4skip)
				return

			}, 1)

			tmp87025 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp87026 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp87027 := Call(__e, tmp87026, V1539)

			tmp87028 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp87029 := Call(__e, tmp87028, V1539)

			tmp87030 := Call(__e, tmp87025, tmp87027, tmp87029)

			__e.TailApply(tmp87021, tmp87030)
			return

		} else {
			tmp87020 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87020)
			return

		}

	}, 1)

	tmp87041 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comma_6, tmp87018)

	_ = tmp87041

	tmp87042 := MakeNative(func(__e *ControlFlow) {
		V1542 := __e.Get(1)
		_ = V1542
		tmp87061 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87062 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87063 := Call(__e, tmp87062, V1542)

		tmp87064 := Call(__e, tmp87061, tmp87063)

		var ifres87055 Obj

		if True == tmp87064 {
			tmp87057 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87058 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87059 := Call(__e, tmp87058, V1542)

			tmp87060 := Call(__e, tmp87057, MakeNumber(61), tmp87059)

			var ifres87056 Obj

			if True == tmp87060 {
				ifres87056 = True

			} else {
				ifres87056 = False

			}

			ifres87055 = ifres87056

		} else {
			ifres87055 = False

		}

		if True == ifres87055 {
			tmp87045 := MakeNative(func(__e *ControlFlow) {
				NewStream1540 := __e.Get(1)
				_ = NewStream1540
				tmp87046 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87047 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87048 := Call(__e, tmp87047, NewStream1540)

				__e.TailApply(tmp87046, tmp87048, symshen_4skip)
				return

			}, 1)

			tmp87049 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp87050 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp87051 := Call(__e, tmp87050, V1542)

			tmp87052 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp87053 := Call(__e, tmp87052, V1542)

			tmp87054 := Call(__e, tmp87049, tmp87051, tmp87053)

			__e.TailApply(tmp87045, tmp87054)
			return

		} else {
			tmp87044 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87044)
			return

		}

	}, 1)

	tmp87065 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5equal_6, tmp87042)

	_ = tmp87065

	tmp87066 := MakeNative(func(__e *ControlFlow) {
		V1545 := __e.Get(1)
		_ = V1545
		tmp87085 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87086 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87087 := Call(__e, tmp87086, V1545)

		tmp87088 := Call(__e, tmp87085, tmp87087)

		var ifres87079 Obj

		if True == tmp87088 {
			tmp87081 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87083 := Call(__e, tmp87082, V1545)

			tmp87084 := Call(__e, tmp87081, MakeNumber(45), tmp87083)

			var ifres87080 Obj

			if True == tmp87084 {
				ifres87080 = True

			} else {
				ifres87080 = False

			}

			ifres87079 = ifres87080

		} else {
			ifres87079 = False

		}

		if True == ifres87079 {
			tmp87069 := MakeNative(func(__e *ControlFlow) {
				NewStream1543 := __e.Get(1)
				_ = NewStream1543
				tmp87070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87071 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87072 := Call(__e, tmp87071, NewStream1543)

				__e.TailApply(tmp87070, tmp87072, symshen_4skip)
				return

			}, 1)

			tmp87073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp87074 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp87075 := Call(__e, tmp87074, V1545)

			tmp87076 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp87077 := Call(__e, tmp87076, V1545)

			tmp87078 := Call(__e, tmp87073, tmp87075, tmp87077)

			__e.TailApply(tmp87069, tmp87078)
			return

		} else {
			tmp87068 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87068)
			return

		}

	}, 1)

	tmp87089 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5minus_6, tmp87066)

	_ = tmp87089

	tmp87090 := MakeNative(func(__e *ControlFlow) {
		V1548 := __e.Get(1)
		_ = V1548
		tmp87109 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87110 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87111 := Call(__e, tmp87110, V1548)

		tmp87112 := Call(__e, tmp87109, tmp87111)

		var ifres87103 Obj

		if True == tmp87112 {
			tmp87105 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87107 := Call(__e, tmp87106, V1548)

			tmp87108 := Call(__e, tmp87105, MakeNumber(40), tmp87107)

			var ifres87104 Obj

			if True == tmp87108 {
				ifres87104 = True

			} else {
				ifres87104 = False

			}

			ifres87103 = ifres87104

		} else {
			ifres87103 = False

		}

		if True == ifres87103 {
			tmp87093 := MakeNative(func(__e *ControlFlow) {
				NewStream1546 := __e.Get(1)
				_ = NewStream1546
				tmp87094 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87095 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87096 := Call(__e, tmp87095, NewStream1546)

				__e.TailApply(tmp87094, tmp87096, symshen_4skip)
				return

			}, 1)

			tmp87097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp87098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp87099 := Call(__e, tmp87098, V1548)

			tmp87100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp87101 := Call(__e, tmp87100, V1548)

			tmp87102 := Call(__e, tmp87097, tmp87099, tmp87101)

			__e.TailApply(tmp87093, tmp87102)
			return

		} else {
			tmp87092 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87092)
			return

		}

	}, 1)

	tmp87113 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lrb_6, tmp87090)

	_ = tmp87113

	tmp87114 := MakeNative(func(__e *ControlFlow) {
		V1551 := __e.Get(1)
		_ = V1551
		tmp87133 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87134 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87135 := Call(__e, tmp87134, V1551)

		tmp87136 := Call(__e, tmp87133, tmp87135)

		var ifres87127 Obj

		if True == tmp87136 {
			tmp87129 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87130 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87131 := Call(__e, tmp87130, V1551)

			tmp87132 := Call(__e, tmp87129, MakeNumber(41), tmp87131)

			var ifres87128 Obj

			if True == tmp87132 {
				ifres87128 = True

			} else {
				ifres87128 = False

			}

			ifres87127 = ifres87128

		} else {
			ifres87127 = False

		}

		if True == ifres87127 {
			tmp87117 := MakeNative(func(__e *ControlFlow) {
				NewStream1549 := __e.Get(1)
				_ = NewStream1549
				tmp87118 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87119 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87120 := Call(__e, tmp87119, NewStream1549)

				__e.TailApply(tmp87118, tmp87120, symshen_4skip)
				return

			}, 1)

			tmp87121 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp87122 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp87123 := Call(__e, tmp87122, V1551)

			tmp87124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp87125 := Call(__e, tmp87124, V1551)

			tmp87126 := Call(__e, tmp87121, tmp87123, tmp87125)

			__e.TailApply(tmp87117, tmp87126)
			return

		} else {
			tmp87116 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87116)
			return

		}

	}, 1)

	tmp87137 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rrb_6, tmp87114)

	_ = tmp87137

	tmp87138 := MakeNative(func(__e *ControlFlow) {
		V1553 := __e.Get(1)
		_ = V1553
		tmp87139 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp87191 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87192 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87193 := Call(__e, tmp87192)

			tmp87194 := Call(__e, tmp87191, YaccParse, tmp87193)

			if True == tmp87194 {
				tmp87141 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp87170 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87171 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87172 := Call(__e, tmp87171)

					tmp87173 := Call(__e, tmp87170, YaccParse, tmp87172)

					if True == tmp87173 {
						tmp87143 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5sym_6 := __e.Get(1)
							_ = Parse__shen_4_5sym_6
							tmp87162 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp87163 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp87164 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp87165 := Call(__e, tmp87164)

							tmp87166 := Call(__e, tmp87163, tmp87165, Parse__shen_4_5sym_6)

							tmp87167 := Call(__e, tmp87162, tmp87166)

							if True == tmp87167 {
								tmp87146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp87147 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp87148 := Call(__e, tmp87147, Parse__shen_4_5sym_6)

								tmp87158 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp87159 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp87160 := Call(__e, tmp87159, Parse__shen_4_5sym_6)

								tmp87161 := Call(__e, tmp87158, tmp87160, MakeString("<>"))

								var ifres87149 Obj

								if True == tmp87161 {
									tmp87154 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp87155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp87156 := Call(__e, tmp87155, MakeNumber(0), Nil)

									tmp87157 := Call(__e, tmp87154, symvector, tmp87156)

									ifres87149 = tmp87157

								} else {
									tmp87150 := Call(__e, PrimNS1Value(symns2_1value), symintern)

									tmp87151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp87152 := Call(__e, tmp87151, Parse__shen_4_5sym_6)

									tmp87153 := Call(__e, tmp87150, tmp87152)

									ifres87149 = tmp87153

								}

								__e.TailApply(tmp87146, tmp87148, ifres87149)
								return

							} else {
								tmp87145 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp87145)
								return

							}

						}, 1)

						tmp87168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5sym_6)

						tmp87169 := Call(__e, tmp87168, V1553)

						__e.TailApply(tmp87143, tmp87169)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp87174 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5number_6 := __e.Get(1)
					_ = Parse__shen_4_5number_6
					tmp87182 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87183 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87184 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87185 := Call(__e, tmp87184)

					tmp87186 := Call(__e, tmp87183, tmp87185, Parse__shen_4_5number_6)

					tmp87187 := Call(__e, tmp87182, tmp87186)

					if True == tmp87187 {
						tmp87177 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87178 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87179 := Call(__e, tmp87178, Parse__shen_4_5number_6)

						tmp87180 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87181 := Call(__e, tmp87180, Parse__shen_4_5number_6)

						__e.TailApply(tmp87177, tmp87179, tmp87181)
						return

					} else {
						tmp87176 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87176)
						return

					}

				}, 1)

				tmp87188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5number_6)

				tmp87189 := Call(__e, tmp87188, V1553)

				tmp87190 := Call(__e, tmp87174, tmp87189)

				__e.TailApply(tmp87141, tmp87190)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp87195 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5str_6 := __e.Get(1)
			_ = Parse__shen_4_5str_6
			tmp87205 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp87206 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87207 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87208 := Call(__e, tmp87207)

			tmp87209 := Call(__e, tmp87206, tmp87208, Parse__shen_4_5str_6)

			tmp87210 := Call(__e, tmp87205, tmp87209)

			if True == tmp87210 {
				tmp87198 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87199 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87200 := Call(__e, tmp87199, Parse__shen_4_5str_6)

				tmp87201 := Call(__e, PrimNS1Value(symns2_1value), symshen_4control_1chars)

				tmp87202 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp87203 := Call(__e, tmp87202, Parse__shen_4_5str_6)

				tmp87204 := Call(__e, tmp87201, tmp87203)

				__e.TailApply(tmp87198, tmp87200, tmp87204)
				return

			} else {
				tmp87197 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp87197)
				return

			}

		}, 1)

		tmp87211 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5str_6)

		tmp87212 := Call(__e, tmp87211, V1553)

		tmp87213 := Call(__e, tmp87195, tmp87212)

		__e.TailApply(tmp87139, tmp87213)
		return

	}, 1)

	tmp87214 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5atom_6, tmp87138)

	_ = tmp87214

	tmp87215 := MakeNative(func(__e *ControlFlow) {
		V1555 := __e.Get(1)
		_ = V1555
		tmp87272 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp87273 := Call(__e, tmp87272, Nil, V1555)

		if True == tmp87273 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp87270 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp87271 := Call(__e, tmp87270, V1555)

			var ifres87250 Obj

			if True == tmp87271 {
				tmp87266 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp87267 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87268 := Call(__e, tmp87267, V1555)

				tmp87269 := Call(__e, tmp87266, MakeString("c"), tmp87268)

				var ifres87252 Obj

				if True == tmp87269 {
					tmp87262 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp87263 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp87264 := Call(__e, tmp87263, V1555)

					tmp87265 := Call(__e, tmp87262, tmp87264)

					var ifres87254 Obj

					if True == tmp87265 {
						tmp87256 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp87257 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp87259 := Call(__e, tmp87258, V1555)

						tmp87260 := Call(__e, tmp87257, tmp87259)

						tmp87261 := Call(__e, tmp87256, MakeString("#"), tmp87260)

						var ifres87255 Obj

						if True == tmp87261 {
							ifres87255 = True

						} else {
							ifres87255 = False

						}

						ifres87254 = ifres87255

					} else {
						ifres87254 = False

					}

					var ifres87253 Obj

					if True == ifres87254 {
						ifres87253 = True

					} else {
						ifres87253 = False

					}

					ifres87252 = ifres87253

				} else {
					ifres87252 = False

				}

				var ifres87251 Obj

				if True == ifres87252 {
					ifres87251 = True

				} else {
					ifres87251 = False

				}

				ifres87250 = ifres87251

			} else {
				ifres87250 = False

			}

			if True == ifres87250 {
				tmp87229 := MakeNative(func(__e *ControlFlow) {
					CodePoint := __e.Get(1)
					_ = CodePoint
					tmp87230 := MakeNative(func(__e *ControlFlow) {
						AfterCodePoint := __e.Get(1)
						_ = AfterCodePoint
						tmp87231 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp87232 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

						tmp87233 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decimalise)

						tmp87234 := Call(__e, tmp87233, CodePoint)

						tmp87235 := Call(__e, tmp87232, tmp87234)

						tmp87236 := Call(__e, PrimNS1Value(symns2_1value), symshen_4control_1chars)

						tmp87237 := Call(__e, tmp87236, AfterCodePoint)

						__e.TailApply(tmp87231, tmp87235, tmp87237)
						return

					}, 1)

					tmp87238 := Call(__e, PrimNS1Value(symns2_1value), symshen_4after_1codepoint)

					tmp87239 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp87240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp87241 := Call(__e, tmp87240, V1555)

					tmp87242 := Call(__e, tmp87239, tmp87241)

					tmp87243 := Call(__e, tmp87238, tmp87242)

					__e.TailApply(tmp87230, tmp87243)
					return

				}, 1)

				tmp87244 := Call(__e, PrimNS1Value(symns2_1value), symshen_4code_1point)

				tmp87245 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp87246 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp87247 := Call(__e, tmp87246, V1555)

				tmp87248 := Call(__e, tmp87245, tmp87247)

				tmp87249 := Call(__e, tmp87244, tmp87248)

				__e.TailApply(tmp87229, tmp87249)
				return

			} else {
				tmp87227 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp87228 := Call(__e, tmp87227, V1555)

				if True == tmp87228 {
					tmp87220 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

					tmp87221 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp87222 := Call(__e, tmp87221, V1555)

					tmp87223 := Call(__e, PrimNS1Value(symns2_1value), symshen_4control_1chars)

					tmp87224 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp87225 := Call(__e, tmp87224, V1555)

					tmp87226 := Call(__e, tmp87223, tmp87225)

					__e.TailApply(tmp87220, tmp87222, tmp87226)
					return

				} else {
					tmp87219 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp87219, symshen_4control_1chars)
					return

				}

			}

		}

	}, 1)

	tmp87274 := Call(__e, PrimNS1Value(symns2_1set), symshen_4control_1chars, tmp87215)

	_ = tmp87274

	tmp87275 := MakeNative(func(__e *ControlFlow) {
		V1559 := __e.Get(1)
		_ = V1559
		tmp87326 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87327 := Call(__e, tmp87326, V1559)

		var ifres87320 Obj

		if True == tmp87327 {
			tmp87322 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87323 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp87324 := Call(__e, tmp87323, V1559)

			tmp87325 := Call(__e, tmp87322, MakeString(";"), tmp87324)

			var ifres87321 Obj

			if True == tmp87325 {
				ifres87321 = True

			} else {
				ifres87321 = False

			}

			ifres87320 = ifres87321

		} else {
			ifres87320 = False

		}

		if True == ifres87320 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp87318 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp87319 := Call(__e, tmp87318, V1559)

			var ifres87290 Obj

			if True == tmp87319 {
				tmp87292 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp87293 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87294 := Call(__e, tmp87293, V1559)

				tmp87295 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87296 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87297 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87298 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87299 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87300 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87301 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87302 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87305 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87306 := Call(__e, tmp87305, MakeString("0"), Nil)

				tmp87307 := Call(__e, tmp87304, MakeString("9"), tmp87306)

				tmp87308 := Call(__e, tmp87303, MakeString("8"), tmp87307)

				tmp87309 := Call(__e, tmp87302, MakeString("7"), tmp87308)

				tmp87310 := Call(__e, tmp87301, MakeString("6"), tmp87309)

				tmp87311 := Call(__e, tmp87300, MakeString("5"), tmp87310)

				tmp87312 := Call(__e, tmp87299, MakeString("4"), tmp87311)

				tmp87313 := Call(__e, tmp87298, MakeString("3"), tmp87312)

				tmp87314 := Call(__e, tmp87297, MakeString("2"), tmp87313)

				tmp87315 := Call(__e, tmp87296, MakeString("1"), tmp87314)

				tmp87316 := Call(__e, tmp87295, MakeString("0"), tmp87315)

				tmp87317 := Call(__e, tmp87292, tmp87294, tmp87316)

				var ifres87291 Obj

				if True == tmp87317 {
					ifres87291 = True

				} else {
					ifres87291 = False

				}

				ifres87290 = ifres87291

			} else {
				ifres87290 = False

			}

			if True == ifres87290 {
				tmp87283 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87284 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87285 := Call(__e, tmp87284, V1559)

				tmp87286 := Call(__e, PrimNS1Value(symns2_1value), symshen_4code_1point)

				tmp87287 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp87288 := Call(__e, tmp87287, V1559)

				tmp87289 := Call(__e, tmp87286, tmp87288)

				__e.TailApply(tmp87283, tmp87285, tmp87289)
				return

			} else {
				tmp87278 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp87279 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp87280 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp87281 := Call(__e, tmp87280, V1559, MakeString("\n"), symshen_4a)

				tmp87282 := Call(__e, tmp87279, MakeString("code point parse error "), tmp87281)

				__e.TailApply(tmp87278, tmp87282)
				return

			}

		}

	}, 1)

	tmp87328 := Call(__e, PrimNS1Value(symns2_1set), symshen_4code_1point, tmp87275)

	_ = tmp87328

	tmp87329 := MakeNative(func(__e *ControlFlow) {
		V1565 := __e.Get(1)
		_ = V1565
		tmp87348 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp87349 := Call(__e, tmp87348, Nil, V1565)

		if True == tmp87349 {
			__e.Return(Nil)
			return
		} else {
			tmp87346 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp87347 := Call(__e, tmp87346, V1565)

			var ifres87340 Obj

			if True == tmp87347 {
				tmp87342 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp87343 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87344 := Call(__e, tmp87343, V1565)

				tmp87345 := Call(__e, tmp87342, MakeString(";"), tmp87344)

				var ifres87341 Obj

				if True == tmp87345 {
					ifres87341 = True

				} else {
					ifres87341 = False

				}

				ifres87340 = ifres87341

			} else {
				ifres87340 = False

			}

			if True == ifres87340 {
				tmp87339 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				__e.TailApply(tmp87339, V1565)
				return

			} else {
				tmp87337 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp87338 := Call(__e, tmp87337, V1565)

				if True == tmp87338 {
					tmp87334 := Call(__e, PrimNS1Value(symns2_1value), symshen_4after_1codepoint)

					tmp87335 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp87336 := Call(__e, tmp87335, V1565)

					__e.TailApply(tmp87334, tmp87336)
					return

				} else {
					tmp87333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp87333, symshen_4after_1codepoint)
					return

				}

			}

		}

	}, 1)

	tmp87350 := Call(__e, PrimNS1Value(symns2_1set), symshen_4after_1codepoint, tmp87329)

	_ = tmp87350

	tmp87351 := MakeNative(func(__e *ControlFlow) {
		V1567 := __e.Get(1)
		_ = V1567
		tmp87352 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

		tmp87353 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

		tmp87354 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

		tmp87355 := Call(__e, tmp87354, V1567)

		tmp87356 := Call(__e, tmp87353, tmp87355)

		__e.TailApply(tmp87352, tmp87356, MakeNumber(0))
		return

	}, 1)

	tmp87357 := Call(__e, PrimNS1Value(symns2_1set), symshen_4decimalise, tmp87351)

	_ = tmp87357

	tmp87358 := MakeNative(func(__e *ControlFlow) {
		V1573 := __e.Get(1)
		_ = V1573
		tmp87497 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87498 := Call(__e, tmp87497, V1573)

		var ifres87491 Obj

		if True == tmp87498 {
			tmp87493 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87494 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp87495 := Call(__e, tmp87494, V1573)

			tmp87496 := Call(__e, tmp87493, MakeString("0"), tmp87495)

			var ifres87492 Obj

			if True == tmp87496 {
				ifres87492 = True

			} else {
				ifres87492 = False

			}

			ifres87491 = ifres87492

		} else {
			ifres87491 = False

		}

		if True == ifres87491 {
			tmp87486 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp87487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

			tmp87488 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp87489 := Call(__e, tmp87488, V1573)

			tmp87490 := Call(__e, tmp87487, tmp87489)

			__e.TailApply(tmp87486, MakeNumber(0), tmp87490)
			return

		} else {
			tmp87484 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp87485 := Call(__e, tmp87484, V1573)

			var ifres87478 Obj

			if True == tmp87485 {
				tmp87480 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp87481 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87482 := Call(__e, tmp87481, V1573)

				tmp87483 := Call(__e, tmp87480, MakeString("1"), tmp87482)

				var ifres87479 Obj

				if True == tmp87483 {
					ifres87479 = True

				} else {
					ifres87479 = False

				}

				ifres87478 = ifres87479

			} else {
				ifres87478 = False

			}

			if True == ifres87478 {
				tmp87473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp87474 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

				tmp87475 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp87476 := Call(__e, tmp87475, V1573)

				tmp87477 := Call(__e, tmp87474, tmp87476)

				__e.TailApply(tmp87473, MakeNumber(1), tmp87477)
				return

			} else {
				tmp87471 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp87472 := Call(__e, tmp87471, V1573)

				var ifres87465 Obj

				if True == tmp87472 {
					tmp87467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87468 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp87469 := Call(__e, tmp87468, V1573)

					tmp87470 := Call(__e, tmp87467, MakeString("2"), tmp87469)

					var ifres87466 Obj

					if True == tmp87470 {
						ifres87466 = True

					} else {
						ifres87466 = False

					}

					ifres87465 = ifres87466

				} else {
					ifres87465 = False

				}

				if True == ifres87465 {
					tmp87460 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp87461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

					tmp87462 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp87463 := Call(__e, tmp87462, V1573)

					tmp87464 := Call(__e, tmp87461, tmp87463)

					__e.TailApply(tmp87460, MakeNumber(2), tmp87464)
					return

				} else {
					tmp87458 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp87459 := Call(__e, tmp87458, V1573)

					var ifres87452 Obj

					if True == tmp87459 {
						tmp87454 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp87455 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87456 := Call(__e, tmp87455, V1573)

						tmp87457 := Call(__e, tmp87454, MakeString("3"), tmp87456)

						var ifres87453 Obj

						if True == tmp87457 {
							ifres87453 = True

						} else {
							ifres87453 = False

						}

						ifres87452 = ifres87453

					} else {
						ifres87452 = False

					}

					if True == ifres87452 {
						tmp87447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp87448 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

						tmp87449 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp87450 := Call(__e, tmp87449, V1573)

						tmp87451 := Call(__e, tmp87448, tmp87450)

						__e.TailApply(tmp87447, MakeNumber(3), tmp87451)
						return

					} else {
						tmp87445 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp87446 := Call(__e, tmp87445, V1573)

						var ifres87439 Obj

						if True == tmp87446 {
							tmp87441 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp87442 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp87443 := Call(__e, tmp87442, V1573)

							tmp87444 := Call(__e, tmp87441, MakeString("4"), tmp87443)

							var ifres87440 Obj

							if True == tmp87444 {
								ifres87440 = True

							} else {
								ifres87440 = False

							}

							ifres87439 = ifres87440

						} else {
							ifres87439 = False

						}

						if True == ifres87439 {
							tmp87434 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp87435 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

							tmp87436 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp87437 := Call(__e, tmp87436, V1573)

							tmp87438 := Call(__e, tmp87435, tmp87437)

							__e.TailApply(tmp87434, MakeNumber(4), tmp87438)
							return

						} else {
							tmp87432 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp87433 := Call(__e, tmp87432, V1573)

							var ifres87426 Obj

							if True == tmp87433 {
								tmp87428 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp87429 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp87430 := Call(__e, tmp87429, V1573)

								tmp87431 := Call(__e, tmp87428, MakeString("5"), tmp87430)

								var ifres87427 Obj

								if True == tmp87431 {
									ifres87427 = True

								} else {
									ifres87427 = False

								}

								ifres87426 = ifres87427

							} else {
								ifres87426 = False

							}

							if True == ifres87426 {
								tmp87421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp87422 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

								tmp87423 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp87424 := Call(__e, tmp87423, V1573)

								tmp87425 := Call(__e, tmp87422, tmp87424)

								__e.TailApply(tmp87421, MakeNumber(5), tmp87425)
								return

							} else {
								tmp87419 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp87420 := Call(__e, tmp87419, V1573)

								var ifres87413 Obj

								if True == tmp87420 {
									tmp87415 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp87416 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp87417 := Call(__e, tmp87416, V1573)

									tmp87418 := Call(__e, tmp87415, MakeString("6"), tmp87417)

									var ifres87414 Obj

									if True == tmp87418 {
										ifres87414 = True

									} else {
										ifres87414 = False

									}

									ifres87413 = ifres87414

								} else {
									ifres87413 = False

								}

								if True == ifres87413 {
									tmp87408 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp87409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

									tmp87410 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp87411 := Call(__e, tmp87410, V1573)

									tmp87412 := Call(__e, tmp87409, tmp87411)

									__e.TailApply(tmp87408, MakeNumber(6), tmp87412)
									return

								} else {
									tmp87406 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp87407 := Call(__e, tmp87406, V1573)

									var ifres87400 Obj

									if True == tmp87407 {
										tmp87402 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp87403 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp87404 := Call(__e, tmp87403, V1573)

										tmp87405 := Call(__e, tmp87402, MakeString("7"), tmp87404)

										var ifres87401 Obj

										if True == tmp87405 {
											ifres87401 = True

										} else {
											ifres87401 = False

										}

										ifres87400 = ifres87401

									} else {
										ifres87400 = False

									}

									if True == ifres87400 {
										tmp87395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp87396 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

										tmp87397 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp87398 := Call(__e, tmp87397, V1573)

										tmp87399 := Call(__e, tmp87396, tmp87398)

										__e.TailApply(tmp87395, MakeNumber(7), tmp87399)
										return

									} else {
										tmp87393 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp87394 := Call(__e, tmp87393, V1573)

										var ifres87387 Obj

										if True == tmp87394 {
											tmp87389 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp87390 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp87391 := Call(__e, tmp87390, V1573)

											tmp87392 := Call(__e, tmp87389, MakeString("8"), tmp87391)

											var ifres87388 Obj

											if True == tmp87392 {
												ifres87388 = True

											} else {
												ifres87388 = False

											}

											ifres87387 = ifres87388

										} else {
											ifres87387 = False

										}

										if True == ifres87387 {
											tmp87382 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp87383 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

											tmp87384 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp87385 := Call(__e, tmp87384, V1573)

											tmp87386 := Call(__e, tmp87383, tmp87385)

											__e.TailApply(tmp87382, MakeNumber(8), tmp87386)
											return

										} else {
											tmp87380 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp87381 := Call(__e, tmp87380, V1573)

											var ifres87374 Obj

											if True == tmp87381 {
												tmp87376 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp87377 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp87378 := Call(__e, tmp87377, V1573)

												tmp87379 := Call(__e, tmp87376, MakeString("9"), tmp87378)

												var ifres87375 Obj

												if True == tmp87379 {
													ifres87375 = True

												} else {
													ifres87375 = False

												}

												ifres87374 = ifres87375

											} else {
												ifres87374 = False

											}

											if True == ifres87374 {
												tmp87369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp87370 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

												tmp87371 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp87372 := Call(__e, tmp87371, V1573)

												tmp87373 := Call(__e, tmp87370, tmp87372)

												__e.TailApply(tmp87369, MakeNumber(9), tmp87373)
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

	tmp87499 := Call(__e, PrimNS1Value(symns2_1set), symshen_4digits_1_6integers, tmp87358)

	_ = tmp87499

	tmp87500 := MakeNative(func(__e *ControlFlow) {
		V1575 := __e.Get(1)
		_ = V1575
		tmp87501 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alpha_6 := __e.Get(1)
			_ = Parse__shen_4_5alpha_6
			tmp87524 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp87525 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87526 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87527 := Call(__e, tmp87526)

			tmp87528 := Call(__e, tmp87525, tmp87527, Parse__shen_4_5alpha_6)

			tmp87529 := Call(__e, tmp87524, tmp87528)

			if True == tmp87529 {
				tmp87504 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5alphanums_6 := __e.Get(1)
					_ = Parse__shen_4_5alphanums_6
					tmp87516 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87517 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87518 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87519 := Call(__e, tmp87518)

					tmp87520 := Call(__e, tmp87517, tmp87519, Parse__shen_4_5alphanums_6)

					tmp87521 := Call(__e, tmp87516, tmp87520)

					if True == tmp87521 {
						tmp87507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87508 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87509 := Call(__e, tmp87508, Parse__shen_4_5alphanums_6)

						tmp87510 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp87511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87512 := Call(__e, tmp87511, Parse__shen_4_5alpha_6)

						tmp87513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87514 := Call(__e, tmp87513, Parse__shen_4_5alphanums_6)

						tmp87515 := Call(__e, tmp87510, tmp87512, tmp87514)

						__e.TailApply(tmp87507, tmp87509, tmp87515)
						return

					} else {
						tmp87506 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87506)
						return

					}

				}, 1)

				tmp87522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alphanums_6)

				tmp87523 := Call(__e, tmp87522, Parse__shen_4_5alpha_6)

				__e.TailApply(tmp87504, tmp87523)
				return

			} else {
				tmp87503 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp87503)
				return

			}

		}, 1)

		tmp87530 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alpha_6)

		tmp87531 := Call(__e, tmp87530, V1575)

		__e.TailApply(tmp87501, tmp87531)
		return

	}, 1)

	tmp87532 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sym_6, tmp87500)

	_ = tmp87532

	tmp87533 := MakeNative(func(__e *ControlFlow) {
		V1577 := __e.Get(1)
		_ = V1577
		tmp87534 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp87550 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87551 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87552 := Call(__e, tmp87551)

			tmp87553 := Call(__e, tmp87550, YaccParse, tmp87552)

			if True == tmp87553 {
				tmp87536 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp87542 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87543 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87544 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87545 := Call(__e, tmp87544)

					tmp87546 := Call(__e, tmp87543, tmp87545, Parse___5e_6)

					tmp87547 := Call(__e, tmp87542, tmp87546)

					if True == tmp87547 {
						tmp87539 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87540 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87541 := Call(__e, tmp87540, Parse___5e_6)

						__e.TailApply(tmp87539, tmp87541, MakeString(""))
						return

					} else {
						tmp87538 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87538)
						return

					}

				}, 1)

				tmp87548 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp87549 := Call(__e, tmp87548, V1577)

				__e.TailApply(tmp87536, tmp87549)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp87554 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alphanum_6 := __e.Get(1)
			_ = Parse__shen_4_5alphanum_6
			tmp87577 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp87578 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87579 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87580 := Call(__e, tmp87579)

			tmp87581 := Call(__e, tmp87578, tmp87580, Parse__shen_4_5alphanum_6)

			tmp87582 := Call(__e, tmp87577, tmp87581)

			if True == tmp87582 {
				tmp87557 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5alphanums_6 := __e.Get(1)
					_ = Parse__shen_4_5alphanums_6
					tmp87569 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87570 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87571 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87572 := Call(__e, tmp87571)

					tmp87573 := Call(__e, tmp87570, tmp87572, Parse__shen_4_5alphanums_6)

					tmp87574 := Call(__e, tmp87569, tmp87573)

					if True == tmp87574 {
						tmp87560 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87561 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87562 := Call(__e, tmp87561, Parse__shen_4_5alphanums_6)

						tmp87563 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp87564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87565 := Call(__e, tmp87564, Parse__shen_4_5alphanum_6)

						tmp87566 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87567 := Call(__e, tmp87566, Parse__shen_4_5alphanums_6)

						tmp87568 := Call(__e, tmp87563, tmp87565, tmp87567)

						__e.TailApply(tmp87560, tmp87562, tmp87568)
						return

					} else {
						tmp87559 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87559)
						return

					}

				}, 1)

				tmp87575 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alphanums_6)

				tmp87576 := Call(__e, tmp87575, Parse__shen_4_5alphanum_6)

				__e.TailApply(tmp87557, tmp87576)
				return

			} else {
				tmp87556 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp87556)
				return

			}

		}, 1)

		tmp87583 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alphanum_6)

		tmp87584 := Call(__e, tmp87583, V1577)

		tmp87585 := Call(__e, tmp87554, tmp87584)

		__e.TailApply(tmp87534, tmp87585)
		return

	}, 1)

	tmp87586 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alphanums_6, tmp87533)

	_ = tmp87586

	tmp87587 := MakeNative(func(__e *ControlFlow) {
		V1579 := __e.Get(1)
		_ = V1579
		tmp87588 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp87606 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87607 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87608 := Call(__e, tmp87607)

			tmp87609 := Call(__e, tmp87606, YaccParse, tmp87608)

			if True == tmp87609 {
				tmp87590 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5num_6 := __e.Get(1)
					_ = Parse__shen_4_5num_6
					tmp87598 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87599 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87600 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87601 := Call(__e, tmp87600)

					tmp87602 := Call(__e, tmp87599, tmp87601, Parse__shen_4_5num_6)

					tmp87603 := Call(__e, tmp87598, tmp87602)

					if True == tmp87603 {
						tmp87593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87594 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87595 := Call(__e, tmp87594, Parse__shen_4_5num_6)

						tmp87596 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87597 := Call(__e, tmp87596, Parse__shen_4_5num_6)

						__e.TailApply(tmp87593, tmp87595, tmp87597)
						return

					} else {
						tmp87592 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87592)
						return

					}

				}, 1)

				tmp87604 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5num_6)

				tmp87605 := Call(__e, tmp87604, V1579)

				__e.TailApply(tmp87590, tmp87605)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp87610 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alpha_6 := __e.Get(1)
			_ = Parse__shen_4_5alpha_6
			tmp87618 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp87619 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87620 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87621 := Call(__e, tmp87620)

			tmp87622 := Call(__e, tmp87619, tmp87621, Parse__shen_4_5alpha_6)

			tmp87623 := Call(__e, tmp87618, tmp87622)

			if True == tmp87623 {
				tmp87613 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87614 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87615 := Call(__e, tmp87614, Parse__shen_4_5alpha_6)

				tmp87616 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp87617 := Call(__e, tmp87616, Parse__shen_4_5alpha_6)

				__e.TailApply(tmp87613, tmp87615, tmp87617)
				return

			} else {
				tmp87612 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp87612)
				return

			}

		}, 1)

		tmp87624 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alpha_6)

		tmp87625 := Call(__e, tmp87624, V1579)

		tmp87626 := Call(__e, tmp87610, tmp87625)

		__e.TailApply(tmp87588, tmp87626)
		return

	}, 1)

	tmp87627 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alphanum_6, tmp87587)

	_ = tmp87627

	tmp87628 := MakeNative(func(__e *ControlFlow) {
		V1581 := __e.Get(1)
		_ = V1581
		tmp87649 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87650 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87651 := Call(__e, tmp87650, V1581)

		tmp87652 := Call(__e, tmp87649, tmp87651)

		if True == tmp87652 {
			tmp87631 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp87645 := Call(__e, PrimNS1Value(symns2_1value), symshen_4numbyte_2)

				tmp87646 := Call(__e, tmp87645, Parse__Char)

				if True == tmp87646 {
					tmp87634 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87635 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp87636 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87637 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp87638 := Call(__e, tmp87637, V1581)

					tmp87639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp87640 := Call(__e, tmp87639, V1581)

					tmp87641 := Call(__e, tmp87636, tmp87638, tmp87640)

					tmp87642 := Call(__e, tmp87635, tmp87641)

					tmp87643 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					tmp87644 := Call(__e, tmp87643, Parse__Char)

					__e.TailApply(tmp87634, tmp87642, tmp87644)
					return

				} else {
					tmp87633 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp87633)
					return

				}

			}, 1)

			tmp87647 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87648 := Call(__e, tmp87647, V1581)

			__e.TailApply(tmp87631, tmp87648)
			return

		} else {
			tmp87630 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87630)
			return

		}

	}, 1)

	tmp87653 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5num_6, tmp87628)

	_ = tmp87653

	tmp87654 := MakeNative(func(__e *ControlFlow) {
		V1587 := __e.Get(1)
		_ = V1587
		tmp87683 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp87684 := Call(__e, tmp87683, MakeNumber(48), V1587)

		if True == tmp87684 {
			__e.Return(True)
			return
		} else {
			tmp87681 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87682 := Call(__e, tmp87681, MakeNumber(49), V1587)

			if True == tmp87682 {
				__e.Return(True)
				return
			} else {
				tmp87679 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp87680 := Call(__e, tmp87679, MakeNumber(50), V1587)

				if True == tmp87680 {
					__e.Return(True)
					return
				} else {
					tmp87677 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87678 := Call(__e, tmp87677, MakeNumber(51), V1587)

					if True == tmp87678 {
						__e.Return(True)
						return
					} else {
						tmp87675 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp87676 := Call(__e, tmp87675, MakeNumber(52), V1587)

						if True == tmp87676 {
							__e.Return(True)
							return
						} else {
							tmp87673 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp87674 := Call(__e, tmp87673, MakeNumber(53), V1587)

							if True == tmp87674 {
								__e.Return(True)
								return
							} else {
								tmp87671 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp87672 := Call(__e, tmp87671, MakeNumber(54), V1587)

								if True == tmp87672 {
									__e.Return(True)
									return
								} else {
									tmp87669 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp87670 := Call(__e, tmp87669, MakeNumber(55), V1587)

									if True == tmp87670 {
										__e.Return(True)
										return
									} else {
										tmp87667 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp87668 := Call(__e, tmp87667, MakeNumber(56), V1587)

										if True == tmp87668 {
											__e.Return(True)
											return
										} else {
											tmp87665 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp87666 := Call(__e, tmp87665, MakeNumber(57), V1587)

											if True == tmp87666 {
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

	tmp87685 := Call(__e, PrimNS1Value(symns2_1set), symshen_4numbyte_2, tmp87654)

	_ = tmp87685

	tmp87686 := MakeNative(func(__e *ControlFlow) {
		V1589 := __e.Get(1)
		_ = V1589
		tmp87707 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87708 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87709 := Call(__e, tmp87708, V1589)

		tmp87710 := Call(__e, tmp87707, tmp87709)

		if True == tmp87710 {
			tmp87689 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp87703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4symbol_1code_2)

				tmp87704 := Call(__e, tmp87703, Parse__Char)

				if True == tmp87704 {
					tmp87692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87693 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp87694 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp87696 := Call(__e, tmp87695, V1589)

					tmp87697 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp87698 := Call(__e, tmp87697, V1589)

					tmp87699 := Call(__e, tmp87694, tmp87696, tmp87698)

					tmp87700 := Call(__e, tmp87693, tmp87699)

					tmp87701 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					tmp87702 := Call(__e, tmp87701, Parse__Char)

					__e.TailApply(tmp87692, tmp87700, tmp87702)
					return

				} else {
					tmp87691 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp87691)
					return

				}

			}, 1)

			tmp87705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87706 := Call(__e, tmp87705, V1589)

			__e.TailApply(tmp87689, tmp87706)
			return

		} else {
			tmp87688 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87688)
			return

		}

	}, 1)

	tmp87711 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alpha_6, tmp87686)

	_ = tmp87711

	tmp87712 := MakeNative(func(__e *ControlFlow) {
		V1591 := __e.Get(1)
		_ = V1591
		tmp87755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp87756 := Call(__e, tmp87755, V1591, MakeNumber(126))

		if True == tmp87756 {
			__e.Return(True)
			return
		} else {
			tmp87753 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			tmp87754 := Call(__e, tmp87753, V1591, MakeNumber(94))

			var ifres87749 Obj

			if True == tmp87754 {
				tmp87751 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

				tmp87752 := Call(__e, tmp87751, V1591, MakeNumber(123))

				var ifres87750 Obj

				if True == tmp87752 {
					ifres87750 = True

				} else {
					ifres87750 = False

				}

				ifres87749 = ifres87750

			} else {
				ifres87749 = False

			}

			var ifres87715 Obj

			if True == ifres87749 {
				ifres87715 = True

			} else {
				tmp87747 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

				tmp87748 := Call(__e, tmp87747, V1591, MakeNumber(59))

				var ifres87743 Obj

				if True == tmp87748 {
					tmp87745 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

					tmp87746 := Call(__e, tmp87745, V1591, MakeNumber(91))

					var ifres87744 Obj

					if True == tmp87746 {
						ifres87744 = True

					} else {
						ifres87744 = False

					}

					ifres87743 = ifres87744

				} else {
					ifres87743 = False

				}

				var ifres87717 Obj

				if True == ifres87743 {
					ifres87717 = True

				} else {
					tmp87741 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					tmp87742 := Call(__e, tmp87741, V1591, MakeNumber(41))

					var ifres87731 Obj

					if True == tmp87742 {
						tmp87739 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

						tmp87740 := Call(__e, tmp87739, V1591, MakeNumber(58))

						var ifres87733 Obj

						if True == tmp87740 {
							tmp87735 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp87736 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp87737 := Call(__e, tmp87736, V1591, MakeNumber(44))

							tmp87738 := Call(__e, tmp87735, tmp87737)

							var ifres87734 Obj

							if True == tmp87738 {
								ifres87734 = True

							} else {
								ifres87734 = False

							}

							ifres87733 = ifres87734

						} else {
							ifres87733 = False

						}

						var ifres87732 Obj

						if True == ifres87733 {
							ifres87732 = True

						} else {
							ifres87732 = False

						}

						ifres87731 = ifres87732

					} else {
						ifres87731 = False

					}

					var ifres87719 Obj

					if True == ifres87731 {
						ifres87719 = True

					} else {
						tmp87729 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

						tmp87730 := Call(__e, tmp87729, V1591, MakeNumber(34))

						var ifres87725 Obj

						if True == tmp87730 {
							tmp87727 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

							tmp87728 := Call(__e, tmp87727, V1591, MakeNumber(40))

							var ifres87726 Obj

							if True == tmp87728 {
								ifres87726 = True

							} else {
								ifres87726 = False

							}

							ifres87725 = ifres87726

						} else {
							ifres87725 = False

						}

						var ifres87721 Obj

						if True == ifres87725 {
							ifres87721 = True

						} else {
							tmp87723 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp87724 := Call(__e, tmp87723, V1591, MakeNumber(33))

							var ifres87722 Obj

							if True == tmp87724 {
								ifres87722 = True

							} else {
								ifres87722 = False

							}

							ifres87721 = ifres87722

						}

						var ifres87720 Obj

						if True == ifres87721 {
							ifres87720 = True

						} else {
							ifres87720 = False

						}

						ifres87719 = ifres87720

					}

					var ifres87718 Obj

					if True == ifres87719 {
						ifres87718 = True

					} else {
						ifres87718 = False

					}

					ifres87717 = ifres87718

				}

				var ifres87716 Obj

				if True == ifres87717 {
					ifres87716 = True

				} else {
					ifres87716 = False

				}

				ifres87715 = ifres87716

			}

			if True == ifres87715 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp87757 := Call(__e, PrimNS1Value(symns2_1set), symshen_4symbol_1code_2, tmp87712)

	_ = tmp87757

	tmp87758 := MakeNative(func(__e *ControlFlow) {
		V1593 := __e.Get(1)
		_ = V1593
		tmp87759 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5dbq_6 := __e.Get(1)
			_ = Parse__shen_4_5dbq_6
			tmp87789 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp87790 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87791 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87792 := Call(__e, tmp87791)

			tmp87793 := Call(__e, tmp87790, tmp87792, Parse__shen_4_5dbq_6)

			tmp87794 := Call(__e, tmp87789, tmp87793)

			if True == tmp87794 {
				tmp87762 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5strcontents_6 := __e.Get(1)
					_ = Parse__shen_4_5strcontents_6
					tmp87781 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87782 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87783 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87784 := Call(__e, tmp87783)

					tmp87785 := Call(__e, tmp87782, tmp87784, Parse__shen_4_5strcontents_6)

					tmp87786 := Call(__e, tmp87781, tmp87785)

					if True == tmp87786 {
						tmp87765 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5dbq_6 := __e.Get(1)
							_ = Parse__shen_4_5dbq_6
							tmp87773 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp87774 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp87775 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp87776 := Call(__e, tmp87775)

							tmp87777 := Call(__e, tmp87774, tmp87776, Parse__shen_4_5dbq_6)

							tmp87778 := Call(__e, tmp87773, tmp87777)

							if True == tmp87778 {
								tmp87768 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp87769 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp87770 := Call(__e, tmp87769, Parse__shen_4_5dbq_6)

								tmp87771 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp87772 := Call(__e, tmp87771, Parse__shen_4_5strcontents_6)

								__e.TailApply(tmp87768, tmp87770, tmp87772)
								return

							} else {
								tmp87767 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp87767)
								return

							}

						}, 1)

						tmp87779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5dbq_6)

						tmp87780 := Call(__e, tmp87779, Parse__shen_4_5strcontents_6)

						__e.TailApply(tmp87765, tmp87780)
						return

					} else {
						tmp87764 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87764)
						return

					}

				}, 1)

				tmp87787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5strcontents_6)

				tmp87788 := Call(__e, tmp87787, Parse__shen_4_5dbq_6)

				__e.TailApply(tmp87762, tmp87788)
				return

			} else {
				tmp87761 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp87761)
				return

			}

		}, 1)

		tmp87795 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5dbq_6)

		tmp87796 := Call(__e, tmp87795, V1593)

		__e.TailApply(tmp87759, tmp87796)
		return

	}, 1)

	tmp87797 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5str_6, tmp87758)

	_ = tmp87797

	tmp87798 := MakeNative(func(__e *ControlFlow) {
		V1595 := __e.Get(1)
		_ = V1595
		tmp87817 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87818 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87819 := Call(__e, tmp87818, V1595)

		tmp87820 := Call(__e, tmp87817, tmp87819)

		if True == tmp87820 {
			tmp87801 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp87813 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp87814 := Call(__e, tmp87813, Parse__Char, MakeNumber(34))

				if True == tmp87814 {
					tmp87804 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87805 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp87806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87807 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp87808 := Call(__e, tmp87807, V1595)

					tmp87809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp87810 := Call(__e, tmp87809, V1595)

					tmp87811 := Call(__e, tmp87806, tmp87808, tmp87810)

					tmp87812 := Call(__e, tmp87805, tmp87811)

					__e.TailApply(tmp87804, tmp87812, Parse__Char)
					return

				} else {
					tmp87803 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp87803)
					return

				}

			}, 1)

			tmp87815 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87816 := Call(__e, tmp87815, V1595)

			__e.TailApply(tmp87801, tmp87816)
			return

		} else {
			tmp87800 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87800)
			return

		}

	}, 1)

	tmp87821 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5dbq_6, tmp87798)

	_ = tmp87821

	tmp87822 := MakeNative(func(__e *ControlFlow) {
		V1597 := __e.Get(1)
		_ = V1597
		tmp87823 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp87839 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87840 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87841 := Call(__e, tmp87840)

			tmp87842 := Call(__e, tmp87839, YaccParse, tmp87841)

			if True == tmp87842 {
				tmp87825 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp87831 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87832 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87833 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87834 := Call(__e, tmp87833)

					tmp87835 := Call(__e, tmp87832, tmp87834, Parse___5e_6)

					tmp87836 := Call(__e, tmp87831, tmp87835)

					if True == tmp87836 {
						tmp87828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87829 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87830 := Call(__e, tmp87829, Parse___5e_6)

						__e.TailApply(tmp87828, tmp87830, Nil)
						return

					} else {
						tmp87827 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87827)
						return

					}

				}, 1)

				tmp87837 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp87838 := Call(__e, tmp87837, V1597)

				__e.TailApply(tmp87825, tmp87838)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp87843 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5strc_6 := __e.Get(1)
			_ = Parse__shen_4_5strc_6
			tmp87866 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp87867 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp87868 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp87869 := Call(__e, tmp87868)

			tmp87870 := Call(__e, tmp87867, tmp87869, Parse__shen_4_5strc_6)

			tmp87871 := Call(__e, tmp87866, tmp87870)

			if True == tmp87871 {
				tmp87846 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5strcontents_6 := __e.Get(1)
					_ = Parse__shen_4_5strcontents_6
					tmp87858 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp87859 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp87860 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp87861 := Call(__e, tmp87860)

					tmp87862 := Call(__e, tmp87859, tmp87861, Parse__shen_4_5strcontents_6)

					tmp87863 := Call(__e, tmp87858, tmp87862)

					if True == tmp87863 {
						tmp87849 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp87850 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp87851 := Call(__e, tmp87850, Parse__shen_4_5strcontents_6)

						tmp87852 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp87853 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87854 := Call(__e, tmp87853, Parse__shen_4_5strc_6)

						tmp87855 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp87856 := Call(__e, tmp87855, Parse__shen_4_5strcontents_6)

						tmp87857 := Call(__e, tmp87852, tmp87854, tmp87856)

						__e.TailApply(tmp87849, tmp87851, tmp87857)
						return

					} else {
						tmp87848 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp87848)
						return

					}

				}, 1)

				tmp87864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5strcontents_6)

				tmp87865 := Call(__e, tmp87864, Parse__shen_4_5strc_6)

				__e.TailApply(tmp87846, tmp87865)
				return

			} else {
				tmp87845 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp87845)
				return

			}

		}, 1)

		tmp87872 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5strc_6)

		tmp87873 := Call(__e, tmp87872, V1597)

		tmp87874 := Call(__e, tmp87843, tmp87873)

		__e.TailApply(tmp87823, tmp87874)
		return

	}, 1)

	tmp87875 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5strcontents_6, tmp87822)

	_ = tmp87875

	tmp87876 := MakeNative(func(__e *ControlFlow) {
		V1599 := __e.Get(1)
		_ = V1599
		tmp87893 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87894 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87895 := Call(__e, tmp87894, V1599)

		tmp87896 := Call(__e, tmp87893, tmp87895)

		if True == tmp87896 {
			tmp87879 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp87880 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87881 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp87882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp87883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp87884 := Call(__e, tmp87883, V1599)

				tmp87885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp87886 := Call(__e, tmp87885, V1599)

				tmp87887 := Call(__e, tmp87882, tmp87884, tmp87886)

				tmp87888 := Call(__e, tmp87881, tmp87887)

				tmp87889 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				tmp87890 := Call(__e, tmp87889, Parse__Char)

				__e.TailApply(tmp87880, tmp87888, tmp87890)
				return

			}, 1)

			tmp87891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87892 := Call(__e, tmp87891, V1599)

			__e.TailApply(tmp87879, tmp87892)
			return

		} else {
			tmp87878 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87878)
			return

		}

	}, 1)

	tmp87897 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5byte_6, tmp87876)

	_ = tmp87897

	tmp87898 := MakeNative(func(__e *ControlFlow) {
		V1601 := __e.Get(1)
		_ = V1601
		tmp87921 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp87922 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp87923 := Call(__e, tmp87922, V1601)

		tmp87924 := Call(__e, tmp87921, tmp87923)

		if True == tmp87924 {
			tmp87901 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp87915 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp87916 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp87917 := Call(__e, tmp87916, Parse__Char, MakeNumber(34))

				tmp87918 := Call(__e, tmp87915, tmp87917)

				if True == tmp87918 {
					tmp87904 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87905 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp87906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp87907 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp87908 := Call(__e, tmp87907, V1601)

					tmp87909 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp87910 := Call(__e, tmp87909, V1601)

					tmp87911 := Call(__e, tmp87906, tmp87908, tmp87910)

					tmp87912 := Call(__e, tmp87905, tmp87911)

					tmp87913 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					tmp87914 := Call(__e, tmp87913, Parse__Char)

					__e.TailApply(tmp87904, tmp87912, tmp87914)
					return

				} else {
					tmp87903 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp87903)
					return

				}

			}, 1)

			tmp87919 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp87920 := Call(__e, tmp87919, V1601)

			__e.TailApply(tmp87901, tmp87920)
			return

		} else {
			tmp87900 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp87900)
			return

		}

	}, 1)

	tmp87925 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5strc_6, tmp87898)

	_ = tmp87925

	tmp87926 := MakeNative(func(__e *ControlFlow) {
		V1603 := __e.Get(1)
		_ = V1603
		tmp87927 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88176 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88177 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88178 := Call(__e, tmp88177)

			tmp88179 := Call(__e, tmp88176, YaccParse, tmp88178)

			if True == tmp88179 {
				tmp87929 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp88144 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88145 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88146 := Call(__e, tmp88145)

					tmp88147 := Call(__e, tmp88144, YaccParse, tmp88146)

					if True == tmp88147 {
						tmp87931 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp88063 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88064 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp88065 := Call(__e, tmp88064)

							tmp88066 := Call(__e, tmp88063, YaccParse, tmp88065)

							if True == tmp88066 {
								tmp87933 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									tmp88010 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp88011 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp88012 := Call(__e, tmp88011)

									tmp88013 := Call(__e, tmp88010, YaccParse, tmp88012)

									if True == tmp88013 {
										tmp87935 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											tmp87957 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp87958 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp87959 := Call(__e, tmp87958)

											tmp87960 := Call(__e, tmp87957, YaccParse, tmp87959)

											if True == tmp87960 {
												tmp87937 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5digits_6 := __e.Get(1)
													_ = Parse__shen_4_5digits_6
													tmp87949 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp87950 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp87951 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp87952 := Call(__e, tmp87951)

													tmp87953 := Call(__e, tmp87950, tmp87952, Parse__shen_4_5digits_6)

													tmp87954 := Call(__e, tmp87949, tmp87953)

													if True == tmp87954 {
														tmp87940 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														tmp87941 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp87942 := Call(__e, tmp87941, Parse__shen_4_5digits_6)

														tmp87943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

														tmp87944 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

														tmp87945 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp87946 := Call(__e, tmp87945, Parse__shen_4_5digits_6)

														tmp87947 := Call(__e, tmp87944, tmp87946)

														tmp87948 := Call(__e, tmp87943, tmp87947, MakeNumber(0))

														__e.TailApply(tmp87940, tmp87942, tmp87948)
														return

													} else {
														tmp87939 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp87939)
														return

													}

												}, 1)

												tmp87955 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

												tmp87956 := Call(__e, tmp87955, V1603)

												__e.TailApply(tmp87937, tmp87956)
												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										tmp87961 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5predigits_6 := __e.Get(1)
											_ = Parse__shen_4_5predigits_6
											tmp88001 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp88002 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp88003 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp88004 := Call(__e, tmp88003)

											tmp88005 := Call(__e, tmp88002, tmp88004, Parse__shen_4_5predigits_6)

											tmp88006 := Call(__e, tmp88001, tmp88005)

											if True == tmp88006 {
												tmp87964 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5stop_6 := __e.Get(1)
													_ = Parse__shen_4_5stop_6
													tmp87993 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp87994 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp87995 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp87996 := Call(__e, tmp87995)

													tmp87997 := Call(__e, tmp87994, tmp87996, Parse__shen_4_5stop_6)

													tmp87998 := Call(__e, tmp87993, tmp87997)

													if True == tmp87998 {
														tmp87967 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5postdigits_6 := __e.Get(1)
															_ = Parse__shen_4_5postdigits_6
															tmp87985 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp87986 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp87987 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp87988 := Call(__e, tmp87987)

															tmp87989 := Call(__e, tmp87986, tmp87988, Parse__shen_4_5postdigits_6)

															tmp87990 := Call(__e, tmp87985, tmp87989)

															if True == tmp87990 {
																tmp87970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																tmp87971 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp87972 := Call(__e, tmp87971, Parse__shen_4_5postdigits_6)

																tmp87973 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

																tmp87974 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

																tmp87975 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

																tmp87976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp87977 := Call(__e, tmp87976, Parse__shen_4_5predigits_6)

																tmp87978 := Call(__e, tmp87975, tmp87977)

																tmp87979 := Call(__e, tmp87974, tmp87978, MakeNumber(0))

																tmp87980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4post)

																tmp87981 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp87982 := Call(__e, tmp87981, Parse__shen_4_5postdigits_6)

																tmp87983 := Call(__e, tmp87980, tmp87982, MakeNumber(1))

																tmp87984 := Call(__e, tmp87973, tmp87979, tmp87983)

																__e.TailApply(tmp87970, tmp87972, tmp87984)
																return

															} else {
																tmp87969 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(tmp87969)
																return

															}

														}, 1)

														tmp87991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5postdigits_6)

														tmp87992 := Call(__e, tmp87991, Parse__shen_4_5stop_6)

														__e.TailApply(tmp87967, tmp87992)
														return

													} else {
														tmp87966 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp87966)
														return

													}

												}, 1)

												tmp87999 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5stop_6)

												tmp88000 := Call(__e, tmp87999, Parse__shen_4_5predigits_6)

												__e.TailApply(tmp87964, tmp88000)
												return

											} else {
												tmp87963 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp87963)
												return

											}

										}, 1)

										tmp88007 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5predigits_6)

										tmp88008 := Call(__e, tmp88007, V1603)

										tmp88009 := Call(__e, tmp87961, tmp88008)

										__e.TailApply(tmp87935, tmp88009)
										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								tmp88014 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5digits_6 := __e.Get(1)
									_ = Parse__shen_4_5digits_6
									tmp88054 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp88055 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp88056 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp88057 := Call(__e, tmp88056)

									tmp88058 := Call(__e, tmp88055, tmp88057, Parse__shen_4_5digits_6)

									tmp88059 := Call(__e, tmp88054, tmp88058)

									if True == tmp88059 {
										tmp88017 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5E_6 := __e.Get(1)
											_ = Parse__shen_4_5E_6
											tmp88046 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp88047 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp88048 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp88049 := Call(__e, tmp88048)

											tmp88050 := Call(__e, tmp88047, tmp88049, Parse__shen_4_5E_6)

											tmp88051 := Call(__e, tmp88046, tmp88050)

											if True == tmp88051 {
												tmp88020 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5log10_6 := __e.Get(1)
													_ = Parse__shen_4_5log10_6
													tmp88038 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp88039 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp88040 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp88041 := Call(__e, tmp88040)

													tmp88042 := Call(__e, tmp88039, tmp88041, Parse__shen_4_5log10_6)

													tmp88043 := Call(__e, tmp88038, tmp88042)

													if True == tmp88043 {
														tmp88023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														tmp88024 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp88025 := Call(__e, tmp88024, Parse__shen_4_5log10_6)

														tmp88026 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

														tmp88027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

														tmp88028 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp88029 := Call(__e, tmp88028, Parse__shen_4_5log10_6)

														tmp88030 := Call(__e, tmp88027, MakeNumber(10), tmp88029)

														tmp88031 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

														tmp88032 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

														tmp88033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp88034 := Call(__e, tmp88033, Parse__shen_4_5digits_6)

														tmp88035 := Call(__e, tmp88032, tmp88034)

														tmp88036 := Call(__e, tmp88031, tmp88035, MakeNumber(0))

														tmp88037 := Call(__e, tmp88026, tmp88030, tmp88036)

														__e.TailApply(tmp88023, tmp88025, tmp88037)
														return

													} else {
														tmp88022 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp88022)
														return

													}

												}, 1)

												tmp88044 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5log10_6)

												tmp88045 := Call(__e, tmp88044, Parse__shen_4_5E_6)

												__e.TailApply(tmp88020, tmp88045)
												return

											} else {
												tmp88019 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp88019)
												return

											}

										}, 1)

										tmp88052 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5E_6)

										tmp88053 := Call(__e, tmp88052, Parse__shen_4_5digits_6)

										__e.TailApply(tmp88017, tmp88053)
										return

									} else {
										tmp88016 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp88016)
										return

									}

								}, 1)

								tmp88060 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

								tmp88061 := Call(__e, tmp88060, V1603)

								tmp88062 := Call(__e, tmp88014, tmp88061)

								__e.TailApply(tmp87933, tmp88062)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp88067 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5predigits_6 := __e.Get(1)
							_ = Parse__shen_4_5predigits_6
							tmp88135 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp88136 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88137 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp88138 := Call(__e, tmp88137)

							tmp88139 := Call(__e, tmp88136, tmp88138, Parse__shen_4_5predigits_6)

							tmp88140 := Call(__e, tmp88135, tmp88139)

							if True == tmp88140 {
								tmp88070 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5stop_6 := __e.Get(1)
									_ = Parse__shen_4_5stop_6
									tmp88127 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp88128 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp88129 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp88130 := Call(__e, tmp88129)

									tmp88131 := Call(__e, tmp88128, tmp88130, Parse__shen_4_5stop_6)

									tmp88132 := Call(__e, tmp88127, tmp88131)

									if True == tmp88132 {
										tmp88073 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5postdigits_6 := __e.Get(1)
											_ = Parse__shen_4_5postdigits_6
											tmp88119 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp88120 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp88121 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp88122 := Call(__e, tmp88121)

											tmp88123 := Call(__e, tmp88120, tmp88122, Parse__shen_4_5postdigits_6)

											tmp88124 := Call(__e, tmp88119, tmp88123)

											if True == tmp88124 {
												tmp88076 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5E_6 := __e.Get(1)
													_ = Parse__shen_4_5E_6
													tmp88111 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp88112 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp88113 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp88114 := Call(__e, tmp88113)

													tmp88115 := Call(__e, tmp88112, tmp88114, Parse__shen_4_5E_6)

													tmp88116 := Call(__e, tmp88111, tmp88115)

													if True == tmp88116 {
														tmp88079 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5log10_6 := __e.Get(1)
															_ = Parse__shen_4_5log10_6
															tmp88103 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp88104 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp88105 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp88106 := Call(__e, tmp88105)

															tmp88107 := Call(__e, tmp88104, tmp88106, Parse__shen_4_5log10_6)

															tmp88108 := Call(__e, tmp88103, tmp88107)

															if True == tmp88108 {
																tmp88082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																tmp88083 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp88084 := Call(__e, tmp88083, Parse__shen_4_5log10_6)

																tmp88085 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

																tmp88086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

																tmp88087 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp88088 := Call(__e, tmp88087, Parse__shen_4_5log10_6)

																tmp88089 := Call(__e, tmp88086, MakeNumber(10), tmp88088)

																tmp88090 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

																tmp88091 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

																tmp88092 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

																tmp88093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp88094 := Call(__e, tmp88093, Parse__shen_4_5predigits_6)

																tmp88095 := Call(__e, tmp88092, tmp88094)

																tmp88096 := Call(__e, tmp88091, tmp88095, MakeNumber(0))

																tmp88097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4post)

																tmp88098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp88099 := Call(__e, tmp88098, Parse__shen_4_5postdigits_6)

																tmp88100 := Call(__e, tmp88097, tmp88099, MakeNumber(1))

																tmp88101 := Call(__e, tmp88090, tmp88096, tmp88100)

																tmp88102 := Call(__e, tmp88085, tmp88089, tmp88101)

																__e.TailApply(tmp88082, tmp88084, tmp88102)
																return

															} else {
																tmp88081 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(tmp88081)
																return

															}

														}, 1)

														tmp88109 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5log10_6)

														tmp88110 := Call(__e, tmp88109, Parse__shen_4_5E_6)

														__e.TailApply(tmp88079, tmp88110)
														return

													} else {
														tmp88078 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp88078)
														return

													}

												}, 1)

												tmp88117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5E_6)

												tmp88118 := Call(__e, tmp88117, Parse__shen_4_5postdigits_6)

												__e.TailApply(tmp88076, tmp88118)
												return

											} else {
												tmp88075 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp88075)
												return

											}

										}, 1)

										tmp88125 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5postdigits_6)

										tmp88126 := Call(__e, tmp88125, Parse__shen_4_5stop_6)

										__e.TailApply(tmp88073, tmp88126)
										return

									} else {
										tmp88072 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp88072)
										return

									}

								}, 1)

								tmp88133 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5stop_6)

								tmp88134 := Call(__e, tmp88133, Parse__shen_4_5predigits_6)

								__e.TailApply(tmp88070, tmp88134)
								return

							} else {
								tmp88069 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp88069)
								return

							}

						}, 1)

						tmp88141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5predigits_6)

						tmp88142 := Call(__e, tmp88141, V1603)

						tmp88143 := Call(__e, tmp88067, tmp88142)

						__e.TailApply(tmp87931, tmp88143)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp88148 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5plus_6 := __e.Get(1)
					_ = Parse__shen_4_5plus_6
					tmp88167 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88168 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88169 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88170 := Call(__e, tmp88169)

					tmp88171 := Call(__e, tmp88168, tmp88170, Parse__shen_4_5plus_6)

					tmp88172 := Call(__e, tmp88167, tmp88171)

					if True == tmp88172 {
						tmp88151 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5number_6 := __e.Get(1)
							_ = Parse__shen_4_5number_6
							tmp88159 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp88160 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88161 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp88162 := Call(__e, tmp88161)

							tmp88163 := Call(__e, tmp88160, tmp88162, Parse__shen_4_5number_6)

							tmp88164 := Call(__e, tmp88159, tmp88163)

							if True == tmp88164 {
								tmp88154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp88155 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp88156 := Call(__e, tmp88155, Parse__shen_4_5number_6)

								tmp88157 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp88158 := Call(__e, tmp88157, Parse__shen_4_5number_6)

								__e.TailApply(tmp88154, tmp88156, tmp88158)
								return

							} else {
								tmp88153 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp88153)
								return

							}

						}, 1)

						tmp88165 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5number_6)

						tmp88166 := Call(__e, tmp88165, Parse__shen_4_5plus_6)

						__e.TailApply(tmp88151, tmp88166)
						return

					} else {
						tmp88150 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88150)
						return

					}

				}, 1)

				tmp88173 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5plus_6)

				tmp88174 := Call(__e, tmp88173, V1603)

				tmp88175 := Call(__e, tmp88148, tmp88174)

				__e.TailApply(tmp87929, tmp88175)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88180 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5minus_6 := __e.Get(1)
			_ = Parse__shen_4_5minus_6
			tmp88201 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88202 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88203 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88204 := Call(__e, tmp88203)

			tmp88205 := Call(__e, tmp88202, tmp88204, Parse__shen_4_5minus_6)

			tmp88206 := Call(__e, tmp88201, tmp88205)

			if True == tmp88206 {
				tmp88183 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5number_6 := __e.Get(1)
					_ = Parse__shen_4_5number_6
					tmp88193 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88194 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88195 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88196 := Call(__e, tmp88195)

					tmp88197 := Call(__e, tmp88194, tmp88196, Parse__shen_4_5number_6)

					tmp88198 := Call(__e, tmp88193, tmp88197)

					if True == tmp88198 {
						tmp88186 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88187 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88188 := Call(__e, tmp88187, Parse__shen_4_5number_6)

						tmp88189 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						tmp88190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp88191 := Call(__e, tmp88190, Parse__shen_4_5number_6)

						tmp88192 := Call(__e, tmp88189, MakeNumber(0), tmp88191)

						__e.TailApply(tmp88186, tmp88188, tmp88192)
						return

					} else {
						tmp88185 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88185)
						return

					}

				}, 1)

				tmp88199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5number_6)

				tmp88200 := Call(__e, tmp88199, Parse__shen_4_5minus_6)

				__e.TailApply(tmp88183, tmp88200)
				return

			} else {
				tmp88182 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88182)
				return

			}

		}, 1)

		tmp88207 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5minus_6)

		tmp88208 := Call(__e, tmp88207, V1603)

		tmp88209 := Call(__e, tmp88180, tmp88208)

		__e.TailApply(tmp87927, tmp88209)
		return

	}, 1)

	tmp88210 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5number_6, tmp87926)

	_ = tmp88210

	tmp88211 := MakeNative(func(__e *ControlFlow) {
		V1606 := __e.Get(1)
		_ = V1606
		tmp88230 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88231 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88232 := Call(__e, tmp88231, V1606)

		tmp88233 := Call(__e, tmp88230, tmp88232)

		var ifres88224 Obj

		if True == tmp88233 {
			tmp88226 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88227 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88228 := Call(__e, tmp88227, V1606)

			tmp88229 := Call(__e, tmp88226, MakeNumber(101), tmp88228)

			var ifres88225 Obj

			if True == tmp88229 {
				ifres88225 = True

			} else {
				ifres88225 = False

			}

			ifres88224 = ifres88225

		} else {
			ifres88224 = False

		}

		if True == ifres88224 {
			tmp88214 := MakeNative(func(__e *ControlFlow) {
				NewStream1604 := __e.Get(1)
				_ = NewStream1604
				tmp88215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88216 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88217 := Call(__e, tmp88216, NewStream1604)

				__e.TailApply(tmp88215, tmp88217, symshen_4skip)
				return

			}, 1)

			tmp88218 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp88219 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp88220 := Call(__e, tmp88219, V1606)

			tmp88221 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp88222 := Call(__e, tmp88221, V1606)

			tmp88223 := Call(__e, tmp88218, tmp88220, tmp88222)

			__e.TailApply(tmp88214, tmp88223)
			return

		} else {
			tmp88213 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88213)
			return

		}

	}, 1)

	tmp88234 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5E_6, tmp88211)

	_ = tmp88234

	tmp88235 := MakeNative(func(__e *ControlFlow) {
		V1608 := __e.Get(1)
		_ = V1608
		tmp88236 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88258 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88259 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88260 := Call(__e, tmp88259)

			tmp88261 := Call(__e, tmp88258, YaccParse, tmp88260)

			if True == tmp88261 {
				tmp88238 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					tmp88250 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88251 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88252 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88253 := Call(__e, tmp88252)

					tmp88254 := Call(__e, tmp88251, tmp88253, Parse__shen_4_5digits_6)

					tmp88255 := Call(__e, tmp88250, tmp88254)

					if True == tmp88255 {
						tmp88241 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88242 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88243 := Call(__e, tmp88242, Parse__shen_4_5digits_6)

						tmp88244 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

						tmp88245 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						tmp88246 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp88247 := Call(__e, tmp88246, Parse__shen_4_5digits_6)

						tmp88248 := Call(__e, tmp88245, tmp88247)

						tmp88249 := Call(__e, tmp88244, tmp88248, MakeNumber(0))

						__e.TailApply(tmp88241, tmp88243, tmp88249)
						return

					} else {
						tmp88240 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88240)
						return

					}

				}, 1)

				tmp88256 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

				tmp88257 := Call(__e, tmp88256, V1608)

				__e.TailApply(tmp88238, tmp88257)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88262 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5minus_6 := __e.Get(1)
			_ = Parse__shen_4_5minus_6
			tmp88287 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88288 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88289 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88290 := Call(__e, tmp88289)

			tmp88291 := Call(__e, tmp88288, tmp88290, Parse__shen_4_5minus_6)

			tmp88292 := Call(__e, tmp88287, tmp88291)

			if True == tmp88292 {
				tmp88265 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					tmp88279 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88280 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88281 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88282 := Call(__e, tmp88281)

					tmp88283 := Call(__e, tmp88280, tmp88282, Parse__shen_4_5digits_6)

					tmp88284 := Call(__e, tmp88279, tmp88283)

					if True == tmp88284 {
						tmp88268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88269 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88270 := Call(__e, tmp88269, Parse__shen_4_5digits_6)

						tmp88271 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						tmp88272 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

						tmp88273 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						tmp88274 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp88275 := Call(__e, tmp88274, Parse__shen_4_5digits_6)

						tmp88276 := Call(__e, tmp88273, tmp88275)

						tmp88277 := Call(__e, tmp88272, tmp88276, MakeNumber(0))

						tmp88278 := Call(__e, tmp88271, MakeNumber(0), tmp88277)

						__e.TailApply(tmp88268, tmp88270, tmp88278)
						return

					} else {
						tmp88267 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88267)
						return

					}

				}, 1)

				tmp88285 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

				tmp88286 := Call(__e, tmp88285, Parse__shen_4_5minus_6)

				__e.TailApply(tmp88265, tmp88286)
				return

			} else {
				tmp88264 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88264)
				return

			}

		}, 1)

		tmp88293 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5minus_6)

		tmp88294 := Call(__e, tmp88293, V1608)

		tmp88295 := Call(__e, tmp88262, tmp88294)

		__e.TailApply(tmp88236, tmp88295)
		return

	}, 1)

	tmp88296 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5log10_6, tmp88235)

	_ = tmp88296

	tmp88297 := MakeNative(func(__e *ControlFlow) {
		V1610 := __e.Get(1)
		_ = V1610
		tmp88316 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88317 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88318 := Call(__e, tmp88317, V1610)

		tmp88319 := Call(__e, tmp88316, tmp88318)

		if True == tmp88319 {
			tmp88300 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp88312 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp88313 := Call(__e, tmp88312, Parse__Char, MakeNumber(43))

				if True == tmp88313 {
					tmp88303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88304 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp88305 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp88307 := Call(__e, tmp88306, V1610)

					tmp88308 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp88309 := Call(__e, tmp88308, V1610)

					tmp88310 := Call(__e, tmp88305, tmp88307, tmp88309)

					tmp88311 := Call(__e, tmp88304, tmp88310)

					__e.TailApply(tmp88303, tmp88311, Parse__Char)
					return

				} else {
					tmp88302 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp88302)
					return

				}

			}, 1)

			tmp88314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88315 := Call(__e, tmp88314, V1610)

			__e.TailApply(tmp88300, tmp88315)
			return

		} else {
			tmp88299 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88299)
			return

		}

	}, 1)

	tmp88320 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5plus_6, tmp88297)

	_ = tmp88320

	tmp88321 := MakeNative(func(__e *ControlFlow) {
		V1612 := __e.Get(1)
		_ = V1612
		tmp88340 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88341 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88342 := Call(__e, tmp88341, V1612)

		tmp88343 := Call(__e, tmp88340, tmp88342)

		if True == tmp88343 {
			tmp88324 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				tmp88336 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp88337 := Call(__e, tmp88336, Parse__Char, MakeNumber(46))

				if True == tmp88337 {
					tmp88327 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88328 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp88329 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88330 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp88331 := Call(__e, tmp88330, V1612)

					tmp88332 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp88333 := Call(__e, tmp88332, V1612)

					tmp88334 := Call(__e, tmp88329, tmp88331, tmp88333)

					tmp88335 := Call(__e, tmp88328, tmp88334)

					__e.TailApply(tmp88327, tmp88335, Parse__Char)
					return

				} else {
					tmp88326 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp88326)
					return

				}

			}, 1)

			tmp88338 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88339 := Call(__e, tmp88338, V1612)

			__e.TailApply(tmp88324, tmp88339)
			return

		} else {
			tmp88323 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88323)
			return

		}

	}, 1)

	tmp88344 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5stop_6, tmp88321)

	_ = tmp88344

	tmp88345 := MakeNative(func(__e *ControlFlow) {
		V1614 := __e.Get(1)
		_ = V1614
		tmp88346 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88362 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88363 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88364 := Call(__e, tmp88363)

			tmp88365 := Call(__e, tmp88362, YaccParse, tmp88364)

			if True == tmp88365 {
				tmp88348 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp88354 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88355 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88356 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88357 := Call(__e, tmp88356)

					tmp88358 := Call(__e, tmp88355, tmp88357, Parse___5e_6)

					tmp88359 := Call(__e, tmp88354, tmp88358)

					if True == tmp88359 {
						tmp88351 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88352 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88353 := Call(__e, tmp88352, Parse___5e_6)

						__e.TailApply(tmp88351, tmp88353, Nil)
						return

					} else {
						tmp88350 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88350)
						return

					}

				}, 1)

				tmp88360 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp88361 := Call(__e, tmp88360, V1614)

				__e.TailApply(tmp88348, tmp88361)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88366 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digits_6 := __e.Get(1)
			_ = Parse__shen_4_5digits_6
			tmp88374 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88375 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88376 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88377 := Call(__e, tmp88376)

			tmp88378 := Call(__e, tmp88375, tmp88377, Parse__shen_4_5digits_6)

			tmp88379 := Call(__e, tmp88374, tmp88378)

			if True == tmp88379 {
				tmp88369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88370 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88371 := Call(__e, tmp88370, Parse__shen_4_5digits_6)

				tmp88372 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp88373 := Call(__e, tmp88372, Parse__shen_4_5digits_6)

				__e.TailApply(tmp88369, tmp88371, tmp88373)
				return

			} else {
				tmp88368 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88368)
				return

			}

		}, 1)

		tmp88380 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

		tmp88381 := Call(__e, tmp88380, V1614)

		tmp88382 := Call(__e, tmp88366, tmp88381)

		__e.TailApply(tmp88346, tmp88382)
		return

	}, 1)

	tmp88383 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5predigits_6, tmp88345)

	_ = tmp88383

	tmp88384 := MakeNative(func(__e *ControlFlow) {
		V1616 := __e.Get(1)
		_ = V1616
		tmp88385 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digits_6 := __e.Get(1)
			_ = Parse__shen_4_5digits_6
			tmp88393 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88394 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88395 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88396 := Call(__e, tmp88395)

			tmp88397 := Call(__e, tmp88394, tmp88396, Parse__shen_4_5digits_6)

			tmp88398 := Call(__e, tmp88393, tmp88397)

			if True == tmp88398 {
				tmp88388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88389 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88390 := Call(__e, tmp88389, Parse__shen_4_5digits_6)

				tmp88391 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp88392 := Call(__e, tmp88391, Parse__shen_4_5digits_6)

				__e.TailApply(tmp88388, tmp88390, tmp88392)
				return

			} else {
				tmp88387 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88387)
				return

			}

		}, 1)

		tmp88399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

		tmp88400 := Call(__e, tmp88399, V1616)

		__e.TailApply(tmp88385, tmp88400)
		return

	}, 1)

	tmp88401 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5postdigits_6, tmp88384)

	_ = tmp88401

	tmp88402 := MakeNative(func(__e *ControlFlow) {
		V1618 := __e.Get(1)
		_ = V1618
		tmp88403 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88423 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88424 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88425 := Call(__e, tmp88424)

			tmp88426 := Call(__e, tmp88423, YaccParse, tmp88425)

			if True == tmp88426 {
				tmp88405 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digit_6 := __e.Get(1)
					_ = Parse__shen_4_5digit_6
					tmp88415 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88416 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88417 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88418 := Call(__e, tmp88417)

					tmp88419 := Call(__e, tmp88416, tmp88418, Parse__shen_4_5digit_6)

					tmp88420 := Call(__e, tmp88415, tmp88419)

					if True == tmp88420 {
						tmp88408 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88409 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88410 := Call(__e, tmp88409, Parse__shen_4_5digit_6)

						tmp88411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp88412 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp88413 := Call(__e, tmp88412, Parse__shen_4_5digit_6)

						tmp88414 := Call(__e, tmp88411, tmp88413, Nil)

						__e.TailApply(tmp88408, tmp88410, tmp88414)
						return

					} else {
						tmp88407 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88407)
						return

					}

				}, 1)

				tmp88421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digit_6)

				tmp88422 := Call(__e, tmp88421, V1618)

				__e.TailApply(tmp88405, tmp88422)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88427 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digit_6 := __e.Get(1)
			_ = Parse__shen_4_5digit_6
			tmp88450 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88451 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88452 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88453 := Call(__e, tmp88452)

			tmp88454 := Call(__e, tmp88451, tmp88453, Parse__shen_4_5digit_6)

			tmp88455 := Call(__e, tmp88450, tmp88454)

			if True == tmp88455 {
				tmp88430 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					tmp88442 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88443 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88444 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88445 := Call(__e, tmp88444)

					tmp88446 := Call(__e, tmp88443, tmp88445, Parse__shen_4_5digits_6)

					tmp88447 := Call(__e, tmp88442, tmp88446)

					if True == tmp88447 {
						tmp88433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88434 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88435 := Call(__e, tmp88434, Parse__shen_4_5digits_6)

						tmp88436 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp88437 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp88438 := Call(__e, tmp88437, Parse__shen_4_5digit_6)

						tmp88439 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp88440 := Call(__e, tmp88439, Parse__shen_4_5digits_6)

						tmp88441 := Call(__e, tmp88436, tmp88438, tmp88440)

						__e.TailApply(tmp88433, tmp88435, tmp88441)
						return

					} else {
						tmp88432 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88432)
						return

					}

				}, 1)

				tmp88448 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

				tmp88449 := Call(__e, tmp88448, Parse__shen_4_5digit_6)

				__e.TailApply(tmp88430, tmp88449)
				return

			} else {
				tmp88429 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88429)
				return

			}

		}, 1)

		tmp88456 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digit_6)

		tmp88457 := Call(__e, tmp88456, V1618)

		tmp88458 := Call(__e, tmp88427, tmp88457)

		__e.TailApply(tmp88403, tmp88458)
		return

	}, 1)

	tmp88459 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5digits_6, tmp88402)

	_ = tmp88459

	tmp88460 := MakeNative(func(__e *ControlFlow) {
		V1620 := __e.Get(1)
		_ = V1620
		tmp88481 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88482 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88483 := Call(__e, tmp88482, V1620)

		tmp88484 := Call(__e, tmp88481, tmp88483)

		if True == tmp88484 {
			tmp88463 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp88477 := Call(__e, PrimNS1Value(symns2_1value), symshen_4numbyte_2)

				tmp88478 := Call(__e, tmp88477, Parse__X)

				if True == tmp88478 {
					tmp88466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88467 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp88468 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88469 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp88470 := Call(__e, tmp88469, V1620)

					tmp88471 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp88472 := Call(__e, tmp88471, V1620)

					tmp88473 := Call(__e, tmp88468, tmp88470, tmp88472)

					tmp88474 := Call(__e, tmp88467, tmp88473)

					tmp88475 := Call(__e, PrimNS1Value(symns2_1value), symshen_4byte_1_6digit)

					tmp88476 := Call(__e, tmp88475, Parse__X)

					__e.TailApply(tmp88466, tmp88474, tmp88476)
					return

				} else {
					tmp88465 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp88465)
					return

				}

			}, 1)

			tmp88479 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88480 := Call(__e, tmp88479, V1620)

			__e.TailApply(tmp88463, tmp88480)
			return

		} else {
			tmp88462 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88462)
			return

		}

	}, 1)

	tmp88485 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5digit_6, tmp88460)

	_ = tmp88485

	tmp88486 := MakeNative(func(__e *ControlFlow) {
		V1622 := __e.Get(1)
		_ = V1622
		tmp88516 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp88517 := Call(__e, tmp88516, MakeNumber(48), V1622)

		if True == tmp88517 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp88514 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88515 := Call(__e, tmp88514, MakeNumber(49), V1622)

			if True == tmp88515 {
				__e.Return(MakeNumber(1))
				return
			} else {
				tmp88512 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp88513 := Call(__e, tmp88512, MakeNumber(50), V1622)

				if True == tmp88513 {
					__e.Return(MakeNumber(2))
					return
				} else {
					tmp88510 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88511 := Call(__e, tmp88510, MakeNumber(51), V1622)

					if True == tmp88511 {
						__e.Return(MakeNumber(3))
						return
					} else {
						tmp88508 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp88509 := Call(__e, tmp88508, MakeNumber(52), V1622)

						if True == tmp88509 {
							__e.Return(MakeNumber(4))
							return
						} else {
							tmp88506 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88507 := Call(__e, tmp88506, MakeNumber(53), V1622)

							if True == tmp88507 {
								__e.Return(MakeNumber(5))
								return
							} else {
								tmp88504 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp88505 := Call(__e, tmp88504, MakeNumber(54), V1622)

								if True == tmp88505 {
									__e.Return(MakeNumber(6))
									return
								} else {
									tmp88502 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp88503 := Call(__e, tmp88502, MakeNumber(55), V1622)

									if True == tmp88503 {
										__e.Return(MakeNumber(7))
										return
									} else {
										tmp88500 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp88501 := Call(__e, tmp88500, MakeNumber(56), V1622)

										if True == tmp88501 {
											__e.Return(MakeNumber(8))
											return
										} else {
											tmp88498 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp88499 := Call(__e, tmp88498, MakeNumber(57), V1622)

											if True == tmp88499 {
												__e.Return(MakeNumber(9))
												return
											} else {
												tmp88497 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

												__e.TailApply(tmp88497, symshen_4byte_1_6digit)
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

	tmp88518 := Call(__e, PrimNS1Value(symns2_1set), symshen_4byte_1_6digit, tmp88486)

	_ = tmp88518

	tmp88519 := MakeNative(func(__e *ControlFlow) {
		V1627 := __e.Get(1)
		_ = V1627
		V1628 := __e.Get(2)
		_ = V1628
		tmp88538 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp88539 := Call(__e, tmp88538, Nil, V1627)

		if True == tmp88539 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp88536 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp88537 := Call(__e, tmp88536, V1627)

			if True == tmp88537 {
				tmp88523 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp88524 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				tmp88525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				tmp88526 := Call(__e, tmp88525, MakeNumber(10), V1628)

				tmp88527 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88528 := Call(__e, tmp88527, V1627)

				tmp88529 := Call(__e, tmp88524, tmp88526, tmp88528)

				tmp88530 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

				tmp88531 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp88532 := Call(__e, tmp88531, V1627)

				tmp88533 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp88534 := Call(__e, tmp88533, V1628, MakeNumber(1))

				tmp88535 := Call(__e, tmp88530, tmp88532, tmp88534)

				__e.TailApply(tmp88523, tmp88529, tmp88535)
				return

			} else {
				tmp88522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp88522, symshen_4pre)
				return

			}

		}

	}, 2)

	tmp88540 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pre, tmp88519)

	_ = tmp88540

	tmp88541 := MakeNative(func(__e *ControlFlow) {
		V1633 := __e.Get(1)
		_ = V1633
		V1634 := __e.Get(2)
		_ = V1634
		tmp88562 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp88563 := Call(__e, tmp88562, Nil, V1633)

		if True == tmp88563 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp88560 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp88561 := Call(__e, tmp88560, V1633)

			if True == tmp88561 {
				tmp88545 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp88546 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				tmp88547 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				tmp88548 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp88549 := Call(__e, tmp88548, MakeNumber(0), V1634)

				tmp88550 := Call(__e, tmp88547, MakeNumber(10), tmp88549)

				tmp88551 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88552 := Call(__e, tmp88551, V1633)

				tmp88553 := Call(__e, tmp88546, tmp88550, tmp88552)

				tmp88554 := Call(__e, PrimNS1Value(symns2_1value), symshen_4post)

				tmp88555 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp88556 := Call(__e, tmp88555, V1633)

				tmp88557 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp88558 := Call(__e, tmp88557, V1634, MakeNumber(1))

				tmp88559 := Call(__e, tmp88554, tmp88556, tmp88558)

				__e.TailApply(tmp88545, tmp88553, tmp88559)
				return

			} else {
				tmp88544 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp88544, symshen_4post)
				return

			}

		}

	}, 2)

	tmp88564 := Call(__e, PrimNS1Value(symns2_1set), symshen_4post, tmp88541)

	_ = tmp88564

	tmp88565 := MakeNative(func(__e *ControlFlow) {
		V1639 := __e.Get(1)
		_ = V1639
		V1640 := __e.Get(2)
		_ = V1640
		tmp88582 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp88583 := Call(__e, tmp88582, MakeNumber(0), V1640)

		if True == tmp88583 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp88580 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			tmp88581 := Call(__e, tmp88580, V1640, MakeNumber(0))

			if True == tmp88581 {
				tmp88575 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				tmp88576 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				tmp88577 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp88578 := Call(__e, tmp88577, V1640, MakeNumber(1))

				tmp88579 := Call(__e, tmp88576, V1639, tmp88578)

				__e.TailApply(tmp88575, V1639, tmp88579)
				return

			} else {
				tmp88568 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				tmp88569 := Call(__e, PrimNS1Value(symns2_1value), sym_c)

				tmp88570 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				tmp88571 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp88572 := Call(__e, tmp88571, V1640, MakeNumber(1))

				tmp88573 := Call(__e, tmp88570, V1639, tmp88572)

				tmp88574 := Call(__e, tmp88569, tmp88573, V1639)

				__e.TailApply(tmp88568, MakeNumber(1), tmp88574)
				return

			}

		}

	}, 2)

	tmp88584 := Call(__e, PrimNS1Value(symns2_1set), symshen_4expt, tmp88565)

	_ = tmp88584

	tmp88585 := MakeNative(func(__e *ControlFlow) {
		V1642 := __e.Get(1)
		_ = V1642
		tmp88586 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5st__input_6 := __e.Get(1)
			_ = Parse__shen_4_5st__input_6
			tmp88594 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88595 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88596 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88597 := Call(__e, tmp88596)

			tmp88598 := Call(__e, tmp88595, tmp88597, Parse__shen_4_5st__input_6)

			tmp88599 := Call(__e, tmp88594, tmp88598)

			if True == tmp88599 {
				tmp88589 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88590 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88591 := Call(__e, tmp88590, Parse__shen_4_5st__input_6)

				tmp88592 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp88593 := Call(__e, tmp88592, Parse__shen_4_5st__input_6)

				__e.TailApply(tmp88589, tmp88591, tmp88593)
				return

			} else {
				tmp88588 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88588)
				return

			}

		}, 1)

		tmp88600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

		tmp88601 := Call(__e, tmp88600, V1642)

		__e.TailApply(tmp88586, tmp88601)
		return

	}, 1)

	tmp88602 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input1_6, tmp88585)

	_ = tmp88602

	tmp88603 := MakeNative(func(__e *ControlFlow) {
		V1644 := __e.Get(1)
		_ = V1644
		tmp88604 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5st__input_6 := __e.Get(1)
			_ = Parse__shen_4_5st__input_6
			tmp88612 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88613 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88614 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88615 := Call(__e, tmp88614)

			tmp88616 := Call(__e, tmp88613, tmp88615, Parse__shen_4_5st__input_6)

			tmp88617 := Call(__e, tmp88612, tmp88616)

			if True == tmp88617 {
				tmp88607 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88608 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88609 := Call(__e, tmp88608, Parse__shen_4_5st__input_6)

				tmp88610 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp88611 := Call(__e, tmp88610, Parse__shen_4_5st__input_6)

				__e.TailApply(tmp88607, tmp88609, tmp88611)
				return

			} else {
				tmp88606 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88606)
				return

			}

		}, 1)

		tmp88618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

		tmp88619 := Call(__e, tmp88618, V1644)

		__e.TailApply(tmp88604, tmp88619)
		return

	}, 1)

	tmp88620 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input2_6, tmp88603)

	_ = tmp88620

	tmp88621 := MakeNative(func(__e *ControlFlow) {
		V1646 := __e.Get(1)
		_ = V1646
		tmp88622 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88638 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88639 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88640 := Call(__e, tmp88639)

			tmp88641 := Call(__e, tmp88638, YaccParse, tmp88640)

			if True == tmp88641 {
				tmp88624 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5multiline_6 := __e.Get(1)
					_ = Parse__shen_4_5multiline_6
					tmp88630 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88631 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88632 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88633 := Call(__e, tmp88632)

					tmp88634 := Call(__e, tmp88631, tmp88633, Parse__shen_4_5multiline_6)

					tmp88635 := Call(__e, tmp88630, tmp88634)

					if True == tmp88635 {
						tmp88627 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88628 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88629 := Call(__e, tmp88628, Parse__shen_4_5multiline_6)

						__e.TailApply(tmp88627, tmp88629, symshen_4skip)
						return

					} else {
						tmp88626 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88626)
						return

					}

				}, 1)

				tmp88636 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5multiline_6)

				tmp88637 := Call(__e, tmp88636, V1646)

				__e.TailApply(tmp88624, tmp88637)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88642 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5singleline_6 := __e.Get(1)
			_ = Parse__shen_4_5singleline_6
			tmp88648 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88649 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88650 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88651 := Call(__e, tmp88650)

			tmp88652 := Call(__e, tmp88649, tmp88651, Parse__shen_4_5singleline_6)

			tmp88653 := Call(__e, tmp88648, tmp88652)

			if True == tmp88653 {
				tmp88645 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88646 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88647 := Call(__e, tmp88646, Parse__shen_4_5singleline_6)

				__e.TailApply(tmp88645, tmp88647, symshen_4skip)
				return

			} else {
				tmp88644 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88644)
				return

			}

		}, 1)

		tmp88654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5singleline_6)

		tmp88655 := Call(__e, tmp88654, V1646)

		tmp88656 := Call(__e, tmp88642, tmp88655)

		__e.TailApply(tmp88622, tmp88656)
		return

	}, 1)

	tmp88657 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comment_6, tmp88621)

	_ = tmp88657

	tmp88658 := MakeNative(func(__e *ControlFlow) {
		V1648 := __e.Get(1)
		_ = V1648
		tmp88659 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5backslash_6 := __e.Get(1)
			_ = Parse__shen_4_5backslash_6
			tmp88698 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88699 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88700 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88701 := Call(__e, tmp88700)

			tmp88702 := Call(__e, tmp88699, tmp88701, Parse__shen_4_5backslash_6)

			tmp88703 := Call(__e, tmp88698, tmp88702)

			if True == tmp88703 {
				tmp88662 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5backslash_6 := __e.Get(1)
					_ = Parse__shen_4_5backslash_6
					tmp88690 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88691 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88692 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88693 := Call(__e, tmp88692)

					tmp88694 := Call(__e, tmp88691, tmp88693, Parse__shen_4_5backslash_6)

					tmp88695 := Call(__e, tmp88690, tmp88694)

					if True == tmp88695 {
						tmp88665 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5anysingle_6 := __e.Get(1)
							_ = Parse__shen_4_5anysingle_6
							tmp88682 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp88683 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88684 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp88685 := Call(__e, tmp88684)

							tmp88686 := Call(__e, tmp88683, tmp88685, Parse__shen_4_5anysingle_6)

							tmp88687 := Call(__e, tmp88682, tmp88686)

							if True == tmp88687 {
								tmp88668 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5return_6 := __e.Get(1)
									_ = Parse__shen_4_5return_6
									tmp88674 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp88675 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp88676 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp88677 := Call(__e, tmp88676)

									tmp88678 := Call(__e, tmp88675, tmp88677, Parse__shen_4_5return_6)

									tmp88679 := Call(__e, tmp88674, tmp88678)

									if True == tmp88679 {
										tmp88671 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp88672 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp88673 := Call(__e, tmp88672, Parse__shen_4_5return_6)

										__e.TailApply(tmp88671, tmp88673, symshen_4skip)
										return

									} else {
										tmp88670 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp88670)
										return

									}

								}, 1)

								tmp88680 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5return_6)

								tmp88681 := Call(__e, tmp88680, Parse__shen_4_5anysingle_6)

								__e.TailApply(tmp88668, tmp88681)
								return

							} else {
								tmp88667 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp88667)
								return

							}

						}, 1)

						tmp88688 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anysingle_6)

						tmp88689 := Call(__e, tmp88688, Parse__shen_4_5backslash_6)

						__e.TailApply(tmp88665, tmp88689)
						return

					} else {
						tmp88664 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88664)
						return

					}

				}, 1)

				tmp88696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

				tmp88697 := Call(__e, tmp88696, Parse__shen_4_5backslash_6)

				__e.TailApply(tmp88662, tmp88697)
				return

			} else {
				tmp88661 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88661)
				return

			}

		}, 1)

		tmp88704 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

		tmp88705 := Call(__e, tmp88704, V1648)

		__e.TailApply(tmp88659, tmp88705)
		return

	}, 1)

	tmp88706 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5singleline_6, tmp88658)

	_ = tmp88706

	tmp88707 := MakeNative(func(__e *ControlFlow) {
		V1651 := __e.Get(1)
		_ = V1651
		tmp88726 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88727 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88728 := Call(__e, tmp88727, V1651)

		tmp88729 := Call(__e, tmp88726, tmp88728)

		var ifres88720 Obj

		if True == tmp88729 {
			tmp88722 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88723 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88724 := Call(__e, tmp88723, V1651)

			tmp88725 := Call(__e, tmp88722, MakeNumber(92), tmp88724)

			var ifres88721 Obj

			if True == tmp88725 {
				ifres88721 = True

			} else {
				ifres88721 = False

			}

			ifres88720 = ifres88721

		} else {
			ifres88720 = False

		}

		if True == ifres88720 {
			tmp88710 := MakeNative(func(__e *ControlFlow) {
				NewStream1649 := __e.Get(1)
				_ = NewStream1649
				tmp88711 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88712 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88713 := Call(__e, tmp88712, NewStream1649)

				__e.TailApply(tmp88711, tmp88713, symshen_4skip)
				return

			}, 1)

			tmp88714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp88715 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp88716 := Call(__e, tmp88715, V1651)

			tmp88717 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp88718 := Call(__e, tmp88717, V1651)

			tmp88719 := Call(__e, tmp88714, tmp88716, tmp88718)

			__e.TailApply(tmp88710, tmp88719)
			return

		} else {
			tmp88709 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88709)
			return

		}

	}, 1)

	tmp88730 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5backslash_6, tmp88707)

	_ = tmp88730

	tmp88731 := MakeNative(func(__e *ControlFlow) {
		V1653 := __e.Get(1)
		_ = V1653
		tmp88732 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88748 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88749 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88750 := Call(__e, tmp88749)

			tmp88751 := Call(__e, tmp88748, YaccParse, tmp88750)

			if True == tmp88751 {
				tmp88734 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp88740 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88741 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88742 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88743 := Call(__e, tmp88742)

					tmp88744 := Call(__e, tmp88741, tmp88743, Parse___5e_6)

					tmp88745 := Call(__e, tmp88740, tmp88744)

					if True == tmp88745 {
						tmp88737 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88738 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88739 := Call(__e, tmp88738, Parse___5e_6)

						__e.TailApply(tmp88737, tmp88739, symshen_4skip)
						return

					} else {
						tmp88736 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88736)
						return

					}

				}, 1)

				tmp88746 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp88747 := Call(__e, tmp88746, V1653)

				__e.TailApply(tmp88734, tmp88747)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88752 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5non_1return_6 := __e.Get(1)
			_ = Parse__shen_4_5non_1return_6
			tmp88769 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88770 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88771 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88772 := Call(__e, tmp88771)

			tmp88773 := Call(__e, tmp88770, tmp88772, Parse__shen_4_5non_1return_6)

			tmp88774 := Call(__e, tmp88769, tmp88773)

			if True == tmp88774 {
				tmp88755 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5anysingle_6 := __e.Get(1)
					_ = Parse__shen_4_5anysingle_6
					tmp88761 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88762 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88763 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88764 := Call(__e, tmp88763)

					tmp88765 := Call(__e, tmp88762, tmp88764, Parse__shen_4_5anysingle_6)

					tmp88766 := Call(__e, tmp88761, tmp88765)

					if True == tmp88766 {
						tmp88758 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88759 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88760 := Call(__e, tmp88759, Parse__shen_4_5anysingle_6)

						__e.TailApply(tmp88758, tmp88760, symshen_4skip)
						return

					} else {
						tmp88757 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88757)
						return

					}

				}, 1)

				tmp88767 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anysingle_6)

				tmp88768 := Call(__e, tmp88767, Parse__shen_4_5non_1return_6)

				__e.TailApply(tmp88755, tmp88768)
				return

			} else {
				tmp88754 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88754)
				return

			}

		}, 1)

		tmp88775 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5non_1return_6)

		tmp88776 := Call(__e, tmp88775, V1653)

		tmp88777 := Call(__e, tmp88752, tmp88776)

		__e.TailApply(tmp88732, tmp88777)
		return

	}, 1)

	tmp88778 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5anysingle_6, tmp88731)

	_ = tmp88778

	tmp88779 := MakeNative(func(__e *ControlFlow) {
		V1655 := __e.Get(1)
		_ = V1655
		tmp88804 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88805 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88806 := Call(__e, tmp88805, V1655)

		tmp88807 := Call(__e, tmp88804, tmp88806)

		if True == tmp88807 {
			tmp88782 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp88794 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp88795 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp88796 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp88797 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp88798 := Call(__e, tmp88797, MakeNumber(13), Nil)

				tmp88799 := Call(__e, tmp88796, MakeNumber(10), tmp88798)

				tmp88800 := Call(__e, tmp88795, Parse__X, tmp88799)

				tmp88801 := Call(__e, tmp88794, tmp88800)

				if True == tmp88801 {
					tmp88785 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88786 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp88787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88788 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp88789 := Call(__e, tmp88788, V1655)

					tmp88790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp88791 := Call(__e, tmp88790, V1655)

					tmp88792 := Call(__e, tmp88787, tmp88789, tmp88791)

					tmp88793 := Call(__e, tmp88786, tmp88792)

					__e.TailApply(tmp88785, tmp88793, symshen_4skip)
					return

				} else {
					tmp88784 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp88784)
					return

				}

			}, 1)

			tmp88802 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88803 := Call(__e, tmp88802, V1655)

			__e.TailApply(tmp88782, tmp88803)
			return

		} else {
			tmp88781 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88781)
			return

		}

	}, 1)

	tmp88808 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5non_1return_6, tmp88779)

	_ = tmp88808

	tmp88809 := MakeNative(func(__e *ControlFlow) {
		V1657 := __e.Get(1)
		_ = V1657
		tmp88832 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88833 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88834 := Call(__e, tmp88833, V1657)

		tmp88835 := Call(__e, tmp88832, tmp88834)

		if True == tmp88835 {
			tmp88812 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp88824 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp88825 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp88826 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp88827 := Call(__e, tmp88826, MakeNumber(13), Nil)

				tmp88828 := Call(__e, tmp88825, MakeNumber(10), tmp88827)

				tmp88829 := Call(__e, tmp88824, Parse__X, tmp88828)

				if True == tmp88829 {
					tmp88815 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88816 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp88817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp88818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp88819 := Call(__e, tmp88818, V1657)

					tmp88820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp88821 := Call(__e, tmp88820, V1657)

					tmp88822 := Call(__e, tmp88817, tmp88819, tmp88821)

					tmp88823 := Call(__e, tmp88816, tmp88822)

					__e.TailApply(tmp88815, tmp88823, symshen_4skip)
					return

				} else {
					tmp88814 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp88814)
					return

				}

			}, 1)

			tmp88830 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88831 := Call(__e, tmp88830, V1657)

			__e.TailApply(tmp88812, tmp88831)
			return

		} else {
			tmp88811 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88811)
			return

		}

	}, 1)

	tmp88836 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5return_6, tmp88809)

	_ = tmp88836

	tmp88837 := MakeNative(func(__e *ControlFlow) {
		V1659 := __e.Get(1)
		_ = V1659
		tmp88838 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5backslash_6 := __e.Get(1)
			_ = Parse__shen_4_5backslash_6
			tmp88866 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88867 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88868 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88869 := Call(__e, tmp88868)

			tmp88870 := Call(__e, tmp88867, tmp88869, Parse__shen_4_5backslash_6)

			tmp88871 := Call(__e, tmp88866, tmp88870)

			if True == tmp88871 {
				tmp88841 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5times_6 := __e.Get(1)
					_ = Parse__shen_4_5times_6
					tmp88858 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88859 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88860 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88861 := Call(__e, tmp88860)

					tmp88862 := Call(__e, tmp88859, tmp88861, Parse__shen_4_5times_6)

					tmp88863 := Call(__e, tmp88858, tmp88862)

					if True == tmp88863 {
						tmp88844 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5anymulti_6 := __e.Get(1)
							_ = Parse__shen_4_5anymulti_6
							tmp88850 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp88851 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88852 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp88853 := Call(__e, tmp88852)

							tmp88854 := Call(__e, tmp88851, tmp88853, Parse__shen_4_5anymulti_6)

							tmp88855 := Call(__e, tmp88850, tmp88854)

							if True == tmp88855 {
								tmp88847 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp88848 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp88849 := Call(__e, tmp88848, Parse__shen_4_5anymulti_6)

								__e.TailApply(tmp88847, tmp88849, symshen_4skip)
								return

							} else {
								tmp88846 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp88846)
								return

							}

						}, 1)

						tmp88856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anymulti_6)

						tmp88857 := Call(__e, tmp88856, Parse__shen_4_5times_6)

						__e.TailApply(tmp88844, tmp88857)
						return

					} else {
						tmp88843 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88843)
						return

					}

				}, 1)

				tmp88864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5times_6)

				tmp88865 := Call(__e, tmp88864, Parse__shen_4_5backslash_6)

				__e.TailApply(tmp88841, tmp88865)
				return

			} else {
				tmp88840 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88840)
				return

			}

		}, 1)

		tmp88872 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

		tmp88873 := Call(__e, tmp88872, V1659)

		__e.TailApply(tmp88838, tmp88873)
		return

	}, 1)

	tmp88874 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5multiline_6, tmp88837)

	_ = tmp88874

	tmp88875 := MakeNative(func(__e *ControlFlow) {
		V1662 := __e.Get(1)
		_ = V1662
		tmp88894 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp88895 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp88896 := Call(__e, tmp88895, V1662)

		tmp88897 := Call(__e, tmp88894, tmp88896)

		var ifres88888 Obj

		if True == tmp88897 {
			tmp88890 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp88892 := Call(__e, tmp88891, V1662)

			tmp88893 := Call(__e, tmp88890, MakeNumber(42), tmp88892)

			var ifres88889 Obj

			if True == tmp88893 {
				ifres88889 = True

			} else {
				ifres88889 = False

			}

			ifres88888 = ifres88889

		} else {
			ifres88888 = False

		}

		if True == ifres88888 {
			tmp88878 := MakeNative(func(__e *ControlFlow) {
				NewStream1660 := __e.Get(1)
				_ = NewStream1660
				tmp88879 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp88880 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp88881 := Call(__e, tmp88880, NewStream1660)

				__e.TailApply(tmp88879, tmp88881, symshen_4skip)
				return

			}, 1)

			tmp88882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp88883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp88884 := Call(__e, tmp88883, V1662)

			tmp88885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp88886 := Call(__e, tmp88885, V1662)

			tmp88887 := Call(__e, tmp88882, tmp88884, tmp88886)

			__e.TailApply(tmp88878, tmp88887)
			return

		} else {
			tmp88877 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp88877)
			return

		}

	}, 1)

	tmp88898 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5times_6, tmp88875)

	_ = tmp88898

	tmp88899 := MakeNative(func(__e *ControlFlow) {
		V1664 := __e.Get(1)
		_ = V1664
		tmp88900 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp88963 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88964 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88965 := Call(__e, tmp88964)

			tmp88966 := Call(__e, tmp88963, YaccParse, tmp88965)

			if True == tmp88966 {
				tmp88902 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp88933 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88934 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88935 := Call(__e, tmp88934)

					tmp88936 := Call(__e, tmp88933, YaccParse, tmp88935)

					if True == tmp88936 {
						tmp88929 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp88930 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88931 := Call(__e, tmp88930, V1664)

						tmp88932 := Call(__e, tmp88929, tmp88931)

						if True == tmp88932 {
							tmp88906 := MakeNative(func(__e *ControlFlow) {
								Parse__X := __e.Get(1)
								_ = Parse__X
								tmp88907 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5anymulti_6 := __e.Get(1)
									_ = Parse__shen_4_5anymulti_6
									tmp88913 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp88914 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp88915 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp88916 := Call(__e, tmp88915)

									tmp88917 := Call(__e, tmp88914, tmp88916, Parse__shen_4_5anymulti_6)

									tmp88918 := Call(__e, tmp88913, tmp88917)

									if True == tmp88918 {
										tmp88910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp88911 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp88912 := Call(__e, tmp88911, Parse__shen_4_5anymulti_6)

										__e.TailApply(tmp88910, tmp88912, symshen_4skip)
										return

									} else {
										tmp88909 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp88909)
										return

									}

								}, 1)

								tmp88919 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anymulti_6)

								tmp88920 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp88921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

								tmp88922 := Call(__e, tmp88921, V1664)

								tmp88923 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp88924 := Call(__e, tmp88923, V1664)

								tmp88925 := Call(__e, tmp88920, tmp88922, tmp88924)

								tmp88926 := Call(__e, tmp88919, tmp88925)

								__e.TailApply(tmp88907, tmp88926)
								return

							}, 1)

							tmp88927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							tmp88928 := Call(__e, tmp88927, V1664)

							__e.TailApply(tmp88906, tmp88928)
							return

						} else {
							tmp88905 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp88905)
							return

						}

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp88937 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5times_6 := __e.Get(1)
					_ = Parse__shen_4_5times_6
					tmp88954 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88955 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88956 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88957 := Call(__e, tmp88956)

					tmp88958 := Call(__e, tmp88955, tmp88957, Parse__shen_4_5times_6)

					tmp88959 := Call(__e, tmp88954, tmp88958)

					if True == tmp88959 {
						tmp88940 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5backslash_6 := __e.Get(1)
							_ = Parse__shen_4_5backslash_6
							tmp88946 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp88947 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp88948 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp88949 := Call(__e, tmp88948)

							tmp88950 := Call(__e, tmp88947, tmp88949, Parse__shen_4_5backslash_6)

							tmp88951 := Call(__e, tmp88946, tmp88950)

							if True == tmp88951 {
								tmp88943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp88944 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp88945 := Call(__e, tmp88944, Parse__shen_4_5backslash_6)

								__e.TailApply(tmp88943, tmp88945, symshen_4skip)
								return

							} else {
								tmp88942 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp88942)
								return

							}

						}, 1)

						tmp88952 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

						tmp88953 := Call(__e, tmp88952, Parse__shen_4_5times_6)

						__e.TailApply(tmp88940, tmp88953)
						return

					} else {
						tmp88939 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88939)
						return

					}

				}, 1)

				tmp88960 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5times_6)

				tmp88961 := Call(__e, tmp88960, V1664)

				tmp88962 := Call(__e, tmp88937, tmp88961)

				__e.TailApply(tmp88902, tmp88962)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp88967 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5comment_6 := __e.Get(1)
			_ = Parse__shen_4_5comment_6
			tmp88984 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp88985 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp88986 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp88987 := Call(__e, tmp88986)

			tmp88988 := Call(__e, tmp88985, tmp88987, Parse__shen_4_5comment_6)

			tmp88989 := Call(__e, tmp88984, tmp88988)

			if True == tmp88989 {
				tmp88970 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5anymulti_6 := __e.Get(1)
					_ = Parse__shen_4_5anymulti_6
					tmp88976 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp88977 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp88978 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp88979 := Call(__e, tmp88978)

					tmp88980 := Call(__e, tmp88977, tmp88979, Parse__shen_4_5anymulti_6)

					tmp88981 := Call(__e, tmp88976, tmp88980)

					if True == tmp88981 {
						tmp88973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp88974 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp88975 := Call(__e, tmp88974, Parse__shen_4_5anymulti_6)

						__e.TailApply(tmp88973, tmp88975, symshen_4skip)
						return

					} else {
						tmp88972 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88972)
						return

					}

				}, 1)

				tmp88982 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anymulti_6)

				tmp88983 := Call(__e, tmp88982, Parse__shen_4_5comment_6)

				__e.TailApply(tmp88970, tmp88983)
				return

			} else {
				tmp88969 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp88969)
				return

			}

		}, 1)

		tmp88990 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comment_6)

		tmp88991 := Call(__e, tmp88990, V1664)

		tmp88992 := Call(__e, tmp88967, tmp88991)

		__e.TailApply(tmp88900, tmp88992)
		return

	}, 1)

	tmp88993 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5anymulti_6, tmp88899)

	_ = tmp88993

	tmp88994 := MakeNative(func(__e *ControlFlow) {
		V1666 := __e.Get(1)
		_ = V1666
		tmp88995 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp89011 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89012 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89013 := Call(__e, tmp89012)

			tmp89014 := Call(__e, tmp89011, YaccParse, tmp89013)

			if True == tmp89014 {
				tmp88997 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5whitespace_6 := __e.Get(1)
					_ = Parse__shen_4_5whitespace_6
					tmp89003 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89004 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89005 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89006 := Call(__e, tmp89005)

					tmp89007 := Call(__e, tmp89004, tmp89006, Parse__shen_4_5whitespace_6)

					tmp89008 := Call(__e, tmp89003, tmp89007)

					if True == tmp89008 {
						tmp89000 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89001 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89002 := Call(__e, tmp89001, Parse__shen_4_5whitespace_6)

						__e.TailApply(tmp89000, tmp89002, symshen_4skip)
						return

					} else {
						tmp88999 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp88999)
						return

					}

				}, 1)

				tmp89009 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespace_6)

				tmp89010 := Call(__e, tmp89009, V1666)

				__e.TailApply(tmp88997, tmp89010)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp89015 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5whitespace_6 := __e.Get(1)
			_ = Parse__shen_4_5whitespace_6
			tmp89032 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp89033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89034 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89035 := Call(__e, tmp89034)

			tmp89036 := Call(__e, tmp89033, tmp89035, Parse__shen_4_5whitespace_6)

			tmp89037 := Call(__e, tmp89032, tmp89036)

			if True == tmp89037 {
				tmp89018 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5whitespaces_6 := __e.Get(1)
					_ = Parse__shen_4_5whitespaces_6
					tmp89024 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89025 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89026 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89027 := Call(__e, tmp89026)

					tmp89028 := Call(__e, tmp89025, tmp89027, Parse__shen_4_5whitespaces_6)

					tmp89029 := Call(__e, tmp89024, tmp89028)

					if True == tmp89029 {
						tmp89021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89022 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89023 := Call(__e, tmp89022, Parse__shen_4_5whitespaces_6)

						__e.TailApply(tmp89021, tmp89023, symshen_4skip)
						return

					} else {
						tmp89020 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89020)
						return

					}

				}, 1)

				tmp89030 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespaces_6)

				tmp89031 := Call(__e, tmp89030, Parse__shen_4_5whitespace_6)

				__e.TailApply(tmp89018, tmp89031)
				return

			} else {
				tmp89017 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp89017)
				return

			}

		}, 1)

		tmp89038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespace_6)

		tmp89039 := Call(__e, tmp89038, V1666)

		tmp89040 := Call(__e, tmp89015, tmp89039)

		__e.TailApply(tmp88995, tmp89040)
		return

	}, 1)

	tmp89041 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5whitespaces_6, tmp88994)

	_ = tmp89041

	tmp89042 := MakeNative(func(__e *ControlFlow) {
		V1668 := __e.Get(1)
		_ = V1668
		tmp89075 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89076 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp89077 := Call(__e, tmp89076, V1668)

		tmp89078 := Call(__e, tmp89075, tmp89077)

		if True == tmp89078 {
			tmp89045 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp89057 := MakeNative(func(__e *ControlFlow) {
					Parse__Case := __e.Get(1)
					_ = Parse__Case
					tmp89070 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89071 := Call(__e, tmp89070, Parse__Case, MakeNumber(32))

					if True == tmp89071 {
						__e.Return(True)
						return
					} else {
						tmp89068 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp89069 := Call(__e, tmp89068, Parse__Case, MakeNumber(13))

						var ifres89060 Obj

						if True == tmp89069 {
							ifres89060 = True

						} else {
							tmp89066 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp89067 := Call(__e, tmp89066, Parse__Case, MakeNumber(10))

							var ifres89062 Obj

							if True == tmp89067 {
								ifres89062 = True

							} else {
								tmp89064 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp89065 := Call(__e, tmp89064, Parse__Case, MakeNumber(9))

								var ifres89063 Obj

								if True == tmp89065 {
									ifres89063 = True

								} else {
									ifres89063 = False

								}

								ifres89062 = ifres89063

							}

							var ifres89061 Obj

							if True == ifres89062 {
								ifres89061 = True

							} else {
								ifres89061 = False

							}

							ifres89060 = ifres89061

						}

						if True == ifres89060 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				tmp89072 := Call(__e, tmp89057, Parse__X)

				if True == tmp89072 {
					tmp89048 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89049 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89050 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89051 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp89052 := Call(__e, tmp89051, V1668)

					tmp89053 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp89054 := Call(__e, tmp89053, V1668)

					tmp89055 := Call(__e, tmp89050, tmp89052, tmp89054)

					tmp89056 := Call(__e, tmp89049, tmp89055)

					__e.TailApply(tmp89048, tmp89056, symshen_4skip)
					return

				} else {
					tmp89047 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp89047)
					return

				}

			}, 1)

			tmp89073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp89074 := Call(__e, tmp89073, V1668)

			__e.TailApply(tmp89045, tmp89074)
			return

		} else {
			tmp89044 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp89044)
			return

		}

	}, 1)

	tmp89079 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5whitespace_6, tmp89042)

	_ = tmp89079

	tmp89080 := MakeNative(func(__e *ControlFlow) {
		V1670 := __e.Get(1)
		_ = V1670
		tmp89141 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp89142 := Call(__e, tmp89141, Nil, V1670)

		if True == tmp89142 {
			__e.Return(Nil)
			return
		} else {
			tmp89139 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp89140 := Call(__e, tmp89139, V1670)

			var ifres89107 Obj

			if True == tmp89140 {
				tmp89135 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp89136 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89137 := Call(__e, tmp89136, V1670)

				tmp89138 := Call(__e, tmp89135, tmp89137)

				var ifres89109 Obj

				if True == tmp89138 {
					tmp89129 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp89130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89132 := Call(__e, tmp89131, V1670)

					tmp89133 := Call(__e, tmp89130, tmp89132)

					tmp89134 := Call(__e, tmp89129, tmp89133)

					var ifres89111 Obj

					if True == tmp89134 {
						tmp89121 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp89122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89124 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89125 := Call(__e, tmp89124, V1670)

						tmp89126 := Call(__e, tmp89123, tmp89125)

						tmp89127 := Call(__e, tmp89122, tmp89126)

						tmp89128 := Call(__e, tmp89121, Nil, tmp89127)

						var ifres89113 Obj

						if True == tmp89128 {
							tmp89115 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp89116 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp89117 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp89118 := Call(__e, tmp89117, V1670)

							tmp89119 := Call(__e, tmp89116, tmp89118)

							tmp89120 := Call(__e, tmp89115, tmp89119, symbar_b)

							var ifres89114 Obj

							if True == tmp89120 {
								ifres89114 = True

							} else {
								ifres89114 = False

							}

							ifres89113 = ifres89114

						} else {
							ifres89113 = False

						}

						var ifres89112 Obj

						if True == ifres89113 {
							ifres89112 = True

						} else {
							ifres89112 = False

						}

						ifres89111 = ifres89112

					} else {
						ifres89111 = False

					}

					var ifres89110 Obj

					if True == ifres89111 {
						ifres89110 = True

					} else {
						ifres89110 = False

					}

					ifres89109 = ifres89110

				} else {
					ifres89109 = False

				}

				var ifres89108 Obj

				if True == ifres89109 {
					ifres89108 = True

				} else {
					ifres89108 = False

				}

				ifres89107 = ifres89108

			} else {
				ifres89107 = False

			}

			if True == ifres89107 {
				tmp89098 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp89099 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp89100 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp89101 := Call(__e, tmp89100, V1670)

				tmp89102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89104 := Call(__e, tmp89103, V1670)

				tmp89105 := Call(__e, tmp89102, tmp89104)

				tmp89106 := Call(__e, tmp89099, tmp89101, tmp89105)

				__e.TailApply(tmp89098, symcons, tmp89106)
				return

			} else {
				tmp89096 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp89097 := Call(__e, tmp89096, V1670)

				if True == tmp89097 {
					tmp89085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp89086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp89087 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89088 := Call(__e, tmp89087, V1670)

					tmp89089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp89090 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cons__form)

					tmp89091 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89092 := Call(__e, tmp89091, V1670)

					tmp89093 := Call(__e, tmp89090, tmp89092)

					tmp89094 := Call(__e, tmp89089, tmp89093, Nil)

					tmp89095 := Call(__e, tmp89086, tmp89088, tmp89094)

					__e.TailApply(tmp89085, symcons, tmp89095)
					return

				} else {
					tmp89084 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp89084, symshen_4cons__form)
					return

				}

			}

		}

	}, 1)

	tmp89143 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cons__form, tmp89080)

	_ = tmp89143

	tmp89144 := MakeNative(func(__e *ControlFlow) {
		V1675 := __e.Get(1)
		_ = V1675
		V1676 := __e.Get(2)
		_ = V1676
		tmp89284 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89285 := Call(__e, tmp89284, V1675)

		var ifres89264 Obj

		if True == tmp89285 {
			tmp89280 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89281 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp89282 := Call(__e, tmp89281, V1675)

			tmp89283 := Call(__e, tmp89280, sym_3, tmp89282)

			var ifres89266 Obj

			if True == tmp89283 {
				tmp89276 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp89277 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89278 := Call(__e, tmp89277, V1675)

				tmp89279 := Call(__e, tmp89276, tmp89278)

				var ifres89268 Obj

				if True == tmp89279 {
					tmp89270 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89271 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89272 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89273 := Call(__e, tmp89272, V1675)

					tmp89274 := Call(__e, tmp89271, tmp89273)

					tmp89275 := Call(__e, tmp89270, Nil, tmp89274)

					var ifres89269 Obj

					if True == tmp89275 {
						ifres89269 = True

					} else {
						ifres89269 = False

					}

					ifres89268 = ifres89269

				} else {
					ifres89268 = False

				}

				var ifres89267 Obj

				if True == ifres89268 {
					ifres89267 = True

				} else {
					ifres89267 = False

				}

				ifres89266 = ifres89267

			} else {
				ifres89266 = False

			}

			var ifres89265 Obj

			if True == ifres89266 {
				ifres89265 = True

			} else {
				ifres89265 = False

			}

			ifres89264 = ifres89265

		} else {
			ifres89264 = False

		}

		if True == ifres89264 {
			tmp89257 := Call(__e, PrimNS1Value(symns2_1value), symappend)

			tmp89258 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

			tmp89259 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp89260 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp89261 := Call(__e, tmp89260, V1675)

			tmp89262 := Call(__e, tmp89259, tmp89261)

			tmp89263 := Call(__e, tmp89258, tmp89262)

			__e.TailApply(tmp89257, tmp89263, V1676)
			return

		} else {
			tmp89255 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp89256 := Call(__e, tmp89255, V1675)

			var ifres89227 Obj

			if True == tmp89256 {
				tmp89251 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp89252 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp89253 := Call(__e, tmp89252, V1675)

				tmp89254 := Call(__e, tmp89251, sympackage, tmp89253)

				var ifres89229 Obj

				if True == tmp89254 {
					tmp89247 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp89248 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89249 := Call(__e, tmp89248, V1675)

					tmp89250 := Call(__e, tmp89247, tmp89249)

					var ifres89231 Obj

					if True == tmp89250 {
						tmp89241 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp89242 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89243 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89244 := Call(__e, tmp89243, V1675)

						tmp89245 := Call(__e, tmp89242, tmp89244)

						tmp89246 := Call(__e, tmp89241, symnull, tmp89245)

						var ifres89233 Obj

						if True == tmp89246 {
							tmp89235 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp89236 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp89237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp89238 := Call(__e, tmp89237, V1675)

							tmp89239 := Call(__e, tmp89236, tmp89238)

							tmp89240 := Call(__e, tmp89235, tmp89239)

							var ifres89234 Obj

							if True == tmp89240 {
								ifres89234 = True

							} else {
								ifres89234 = False

							}

							ifres89233 = ifres89234

						} else {
							ifres89233 = False

						}

						var ifres89232 Obj

						if True == ifres89233 {
							ifres89232 = True

						} else {
							ifres89232 = False

						}

						ifres89231 = ifres89232

					} else {
						ifres89231 = False

					}

					var ifres89230 Obj

					if True == ifres89231 {
						ifres89230 = True

					} else {
						ifres89230 = False

					}

					ifres89229 = ifres89230

				} else {
					ifres89229 = False

				}

				var ifres89228 Obj

				if True == ifres89229 {
					ifres89228 = True

				} else {
					ifres89228 = False

				}

				ifres89227 = ifres89228

			} else {
				ifres89227 = False

			}

			if True == ifres89227 {
				tmp89220 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				tmp89221 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89222 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89223 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89224 := Call(__e, tmp89223, V1675)

				tmp89225 := Call(__e, tmp89222, tmp89224)

				tmp89226 := Call(__e, tmp89221, tmp89225)

				__e.TailApply(tmp89220, tmp89226, V1676)
				return

			} else {
				tmp89218 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp89219 := Call(__e, tmp89218, V1675)

				var ifres89198 Obj

				if True == tmp89219 {
					tmp89214 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89215 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89216 := Call(__e, tmp89215, V1675)

					tmp89217 := Call(__e, tmp89214, sympackage, tmp89216)

					var ifres89200 Obj

					if True == tmp89217 {
						tmp89210 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp89211 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89212 := Call(__e, tmp89211, V1675)

						tmp89213 := Call(__e, tmp89210, tmp89212)

						var ifres89202 Obj

						if True == tmp89213 {
							tmp89204 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp89205 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp89206 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp89207 := Call(__e, tmp89206, V1675)

							tmp89208 := Call(__e, tmp89205, tmp89207)

							tmp89209 := Call(__e, tmp89204, tmp89208)

							var ifres89203 Obj

							if True == tmp89209 {
								ifres89203 = True

							} else {
								ifres89203 = False

							}

							ifres89202 = ifres89203

						} else {
							ifres89202 = False

						}

						var ifres89201 Obj

						if True == ifres89202 {
							ifres89201 = True

						} else {
							ifres89201 = False

						}

						ifres89200 = ifres89201

					} else {
						ifres89200 = False

					}

					var ifres89199 Obj

					if True == ifres89200 {
						ifres89199 = True

					} else {
						ifres89199 = False

					}

					ifres89198 = ifres89199

				} else {
					ifres89198 = False

				}

				if True == ifres89198 {
					tmp89149 := MakeNative(func(__e *ControlFlow) {
						ListofExceptions := __e.Get(1)
						_ = ListofExceptions
						tmp89150 := MakeNative(func(__e *ControlFlow) {
							External := __e.Get(1)
							_ = External
							tmp89151 := MakeNative(func(__e *ControlFlow) {
								PackageNameDot := __e.Get(1)
								_ = PackageNameDot
								tmp89152 := MakeNative(func(__e *ControlFlow) {
									ExpPackageNameDot := __e.Get(1)
									_ = ExpPackageNameDot
									tmp89153 := MakeNative(func(__e *ControlFlow) {
										Packaged := __e.Get(1)
										_ = Packaged
										tmp89154 := MakeNative(func(__e *ControlFlow) {
											Internal := __e.Get(1)
											_ = Internal
											tmp89155 := Call(__e, PrimNS1Value(symns2_1value), symappend)

											__e.TailApply(tmp89155, Packaged, V1676)
											return

										}, 1)

										tmp89156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1internal)

										tmp89157 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp89158 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp89159 := Call(__e, tmp89158, V1675)

										tmp89160 := Call(__e, tmp89157, tmp89159)

										tmp89161 := Call(__e, PrimNS1Value(symns2_1value), symshen_4internal_1symbols)

										tmp89162 := Call(__e, tmp89161, ExpPackageNameDot, Packaged)

										tmp89163 := Call(__e, tmp89156, tmp89160, tmp89162)

										__e.TailApply(tmp89154, tmp89163)
										return

									}, 1)

									tmp89164 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

									tmp89165 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp89166 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp89167 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp89168 := Call(__e, tmp89167, V1675)

									tmp89169 := Call(__e, tmp89166, tmp89168)

									tmp89170 := Call(__e, tmp89165, tmp89169)

									tmp89171 := Call(__e, tmp89164, PackageNameDot, ListofExceptions, tmp89170, ExpPackageNameDot)

									__e.TailApply(tmp89153, tmp89171)
									return

								}, 1)

								tmp89172 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

								tmp89173 := Call(__e, tmp89172, PackageNameDot)

								__e.TailApply(tmp89152, tmp89173)
								return

							}, 1)

							tmp89174 := Call(__e, PrimNS1Value(symns2_1value), symintern)

							tmp89175 := Call(__e, PrimNS1Value(symns2_1value), symcn)

							tmp89176 := Call(__e, PrimNS1Value(symns2_1value), symstr)

							tmp89177 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp89178 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp89179 := Call(__e, tmp89178, V1675)

							tmp89180 := Call(__e, tmp89177, tmp89179)

							tmp89181 := Call(__e, tmp89176, tmp89180)

							tmp89182 := Call(__e, tmp89175, tmp89181, MakeString("."))

							tmp89183 := Call(__e, tmp89174, tmp89182)

							__e.TailApply(tmp89151, tmp89183)
							return

						}, 1)

						tmp89184 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1exceptions)

						tmp89185 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89187 := Call(__e, tmp89186, V1675)

						tmp89188 := Call(__e, tmp89185, tmp89187)

						tmp89189 := Call(__e, tmp89184, ListofExceptions, tmp89188)

						__e.TailApply(tmp89150, tmp89189)
						return

					}, 1)

					tmp89190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

					tmp89191 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89192 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89193 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89194 := Call(__e, tmp89193, V1675)

					tmp89195 := Call(__e, tmp89192, tmp89194)

					tmp89196 := Call(__e, tmp89191, tmp89195)

					tmp89197 := Call(__e, tmp89190, tmp89196)

					__e.TailApply(tmp89149, tmp89197)
					return

				} else {
					tmp89148 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					__e.TailApply(tmp89148, V1675, V1676)
					return

				}

			}

		}

	}, 2)

	tmp89286 := Call(__e, PrimNS1Value(symns2_1set), symshen_4package_1macro, tmp89144)

	_ = tmp89286

	tmp89287 := MakeNative(func(__e *ControlFlow) {
		V1679 := __e.Get(1)
		_ = V1679
		V1680 := __e.Get(2)
		_ = V1680
		tmp89288 := MakeNative(func(__e *ControlFlow) {
			CurrExceptions := __e.Get(1)
			_ = CurrExceptions
			tmp89289 := MakeNative(func(__e *ControlFlow) {
				AllExceptions := __e.Get(1)
				_ = AllExceptions
				tmp89290 := Call(__e, PrimNS1Value(symns2_1value), symput)

				tmp89291 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp89292 := Call(__e, tmp89291, sym_dproperty_1vector_d)

				__e.TailApply(tmp89290, V1680, symshen_4external_1symbols, AllExceptions, tmp89292)
				return

			}, 1)

			tmp89293 := Call(__e, PrimNS1Value(symns2_1value), symunion)

			tmp89294 := Call(__e, tmp89293, V1679, CurrExceptions)

			__e.TailApply(tmp89289, tmp89294)
			return

		}, 1)

		tmp89295 := MakeNative(func(__e *ControlFlow) {
			tmp89296 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp89297 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp89298 := Call(__e, tmp89297, sym_dproperty_1vector_d)

			__e.TailApply(tmp89296, V1680, symshen_4external_1symbols, tmp89298)
			return

		}, 0)

		tmp89299 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp89300 := Call(__e, PrimNS1Value(symtry_1catch), tmp89295, tmp89299)

		__e.TailApply(tmp89288, tmp89300)
		return

	}, 2)

	tmp89301 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1exceptions, tmp89287)

	_ = tmp89301

	tmp89302 := MakeNative(func(__e *ControlFlow) {
		V1683 := __e.Get(1)
		_ = V1683
		V1684 := __e.Get(2)
		_ = V1684
		tmp89303 := Call(__e, PrimNS1Value(symns2_1value), symput)

		tmp89304 := Call(__e, PrimNS1Value(symns2_1value), symunion)

		tmp89305 := MakeNative(func(__e *ControlFlow) {
			tmp89306 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp89307 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp89308 := Call(__e, tmp89307, sym_dproperty_1vector_d)

			__e.TailApply(tmp89306, V1683, symshen_4internal_1symbols, tmp89308)
			return

		}, 0)

		tmp89309 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp89310 := Call(__e, PrimNS1Value(symtry_1catch), tmp89305, tmp89309)

		tmp89311 := Call(__e, tmp89304, V1684, tmp89310)

		tmp89312 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp89313 := Call(__e, tmp89312, sym_dproperty_1vector_d)

		__e.TailApply(tmp89303, V1683, symshen_4internal_1symbols, tmp89311, tmp89313)
		return

	}, 2)

	tmp89314 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1internal, tmp89302)

	_ = tmp89314

	tmp89315 := MakeNative(func(__e *ControlFlow) {
		V1695 := __e.Get(1)
		_ = V1695
		V1696 := __e.Get(2)
		_ = V1696
		tmp89336 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		tmp89337 := Call(__e, tmp89336, V1696)

		var ifres89330 Obj

		if True == tmp89337 {
			tmp89332 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

			tmp89333 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

			tmp89334 := Call(__e, tmp89333, V1696)

			tmp89335 := Call(__e, tmp89332, V1695, tmp89334)

			var ifres89331 Obj

			if True == tmp89335 {
				ifres89331 = True

			} else {
				ifres89331 = False

			}

			ifres89330 = ifres89331

		} else {
			ifres89330 = False

		}

		if True == ifres89330 {
			tmp89329 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp89329, V1696, Nil)
			return

		} else {
			tmp89327 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp89328 := Call(__e, tmp89327, V1696)

			if True == tmp89328 {
				tmp89318 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				tmp89319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4internal_1symbols)

				tmp89320 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp89321 := Call(__e, tmp89320, V1696)

				tmp89322 := Call(__e, tmp89319, V1695, tmp89321)

				tmp89323 := Call(__e, PrimNS1Value(symns2_1value), symshen_4internal_1symbols)

				tmp89324 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89325 := Call(__e, tmp89324, V1696)

				tmp89326 := Call(__e, tmp89323, V1695, tmp89325)

				__e.TailApply(tmp89318, tmp89322, tmp89326)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 2)

	tmp89338 := Call(__e, PrimNS1Value(symns2_1set), symshen_4internal_1symbols, tmp89315)

	_ = tmp89338

	tmp89339 := MakeNative(func(__e *ControlFlow) {
		V1713 := __e.Get(1)
		_ = V1713
		V1714 := __e.Get(2)
		_ = V1714
		V1715 := __e.Get(3)
		_ = V1715
		V1716 := __e.Get(4)
		_ = V1716
		tmp89399 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89400 := Call(__e, tmp89399, V1715)

		if True == tmp89400 {
			tmp89390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp89391 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

			tmp89392 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp89393 := Call(__e, tmp89392, V1715)

			tmp89394 := Call(__e, tmp89391, V1713, V1714, tmp89393, V1716)

			tmp89395 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

			tmp89396 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp89397 := Call(__e, tmp89396, V1715)

			tmp89398 := Call(__e, tmp89395, V1713, V1714, tmp89397, V1716)

			__e.TailApply(tmp89390, tmp89394, tmp89398)
			return

		} else {
			tmp89388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sysfunc_2)

			tmp89389 := Call(__e, tmp89388, V1715)

			var ifres89372 Obj

			if True == tmp89389 {
				ifres89372 = True

			} else {
				tmp89386 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				tmp89387 := Call(__e, tmp89386, V1715)

				var ifres89374 Obj

				if True == tmp89387 {
					ifres89374 = True

				} else {
					tmp89384 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					tmp89385 := Call(__e, tmp89384, V1715, V1714)

					var ifres89376 Obj

					if True == tmp89385 {
						ifres89376 = True

					} else {
						tmp89382 := Call(__e, PrimNS1Value(symns2_1value), symshen_4doubleunderline_2)

						tmp89383 := Call(__e, tmp89382, V1715)

						var ifres89378 Obj

						if True == tmp89383 {
							ifres89378 = True

						} else {
							tmp89380 := Call(__e, PrimNS1Value(symns2_1value), symshen_4singleunderline_2)

							tmp89381 := Call(__e, tmp89380, V1715)

							var ifres89379 Obj

							if True == tmp89381 {
								ifres89379 = True

							} else {
								ifres89379 = False

							}

							ifres89378 = ifres89379

						}

						var ifres89377 Obj

						if True == ifres89378 {
							ifres89377 = True

						} else {
							ifres89377 = False

						}

						ifres89376 = ifres89377

					}

					var ifres89375 Obj

					if True == ifres89376 {
						ifres89375 = True

					} else {
						ifres89375 = False

					}

					ifres89374 = ifres89375

				}

				var ifres89373 Obj

				if True == ifres89374 {
					ifres89373 = True

				} else {
					ifres89373 = False

				}

				ifres89372 = ifres89373

			}

			if True == ifres89372 {
				__e.Return(V1715)
				return
			} else {
				tmp89370 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

				tmp89371 := Call(__e, tmp89370, V1715)

				var ifres89344 Obj

				if True == tmp89371 {
					tmp89346 := MakeNative(func(__e *ControlFlow) {
						ExplodeX := __e.Get(1)
						_ = ExplodeX
						tmp89353 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						tmp89354 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

						tmp89355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89359 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89360 := Call(__e, tmp89359, MakeString("."), Nil)

						tmp89361 := Call(__e, tmp89358, MakeString("n"), tmp89360)

						tmp89362 := Call(__e, tmp89357, MakeString("e"), tmp89361)

						tmp89363 := Call(__e, tmp89356, MakeString("h"), tmp89362)

						tmp89364 := Call(__e, tmp89355, MakeString("s"), tmp89363)

						tmp89365 := Call(__e, tmp89354, tmp89364, ExplodeX)

						tmp89366 := Call(__e, tmp89353, tmp89365)

						if True == tmp89366 {
							tmp89349 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp89350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

							tmp89351 := Call(__e, tmp89350, V1716, ExplodeX)

							tmp89352 := Call(__e, tmp89349, tmp89351)

							if True == tmp89352 {
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

					tmp89367 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

					tmp89368 := Call(__e, tmp89367, V1715)

					tmp89369 := Call(__e, tmp89346, tmp89368)

					var ifres89345 Obj

					if True == tmp89369 {
						ifres89345 = True

					} else {
						ifres89345 = False

					}

					ifres89344 = ifres89345

				} else {
					ifres89344 = False

				}

				if True == ifres89344 {
					tmp89343 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					__e.TailApply(tmp89343, V1713, V1715)
					return

				} else {
					__e.Return(V1715)
					return
				}

			}

		}

	}, 4)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4packageh, tmp89339)
	return

}, 0)
