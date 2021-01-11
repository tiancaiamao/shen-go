package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen10087 := MakeNative(func(__e *ControlFlow) {
		V1430 := __e.Get(1)
		_ = V1430
		gen10086 := Call(__e, PrimNS1Value(symns2_1value), symread_1byte)

		__e.TailApply(gen10086, V1430)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1char_1code, gen10087)

	gen10091 := MakeNative(func(__e *ControlFlow) {
		V1432 := __e.Get(1)
		_ = V1432
		gen10088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist)

		gen10090 := MakeNative(func(__e *ControlFlow) {
			S := __e.Get(1)
			_ = S
			gen10089 := Call(__e, PrimNS1Value(symns2_1value), symread_1byte)

			__e.TailApply(gen10089, S)

			return

		}, 1)

		__e.TailApply(gen10088, V1432, gen10090)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symread_1file_1as_1bytelist, gen10091)

	gen10095 := MakeNative(func(__e *ControlFlow) {
		V1434 := __e.Get(1)
		_ = V1434
		gen10092 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist)

		gen10094 := MakeNative(func(__e *ControlFlow) {
			S := __e.Get(1)
			_ = S
			gen10093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

			__e.TailApply(gen10093, S)

			return

		}, 1)

		__e.TailApply(gen10092, V1434, gen10094)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1charlist, gen10095)

	gen10108 := MakeNative(func(__e *ControlFlow) {
		V1437 := __e.Get(1)
		_ = V1437
		V1438 := __e.Get(2)
		_ = V1438
		gen10105 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			gen10103 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen10100 := MakeNative(func(__e *ControlFlow) {
					Xs := __e.Get(1)
					_ = Xs
					gen10097 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						gen10096 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						__e.TailApply(gen10096, Xs)

						return

					}, 1)

					gen10098 := Call(__e, PrimNS1Value(symns2_1value), symclose)

					gen10099 := Call(__e, gen10098, Stream)

					__e.TailApply(gen10097, gen10099)

					return

				}, 1)

				gen10101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist_1help)

				gen10102 := Call(__e, gen10101, Stream, V1438, X, Nil)

				__e.TailApply(gen10100, gen10102)

				return

			}, 1)

			gen10104 := Call(__e, V1438, Stream)

			__e.TailApply(gen10103, gen10104)

			return

		}, 1)

		gen10106 := Call(__e, PrimNS1Value(symns2_1value), symopen)

		gen10107 := Call(__e, gen10106, V1437, symin)

		__e.TailApply(gen10105, gen10107)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1Xlist, gen10108)

	gen10116 := MakeNative(func(__e *ControlFlow) {
		V1443 := __e.Get(1)
		_ = V1443
		V1444 := __e.Get(2)
		_ = V1444
		V1445 := __e.Get(3)
		_ = V1445
		V1446 := __e.Get(4)
		_ = V1446
		gen10114 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10115 := Call(__e, gen10114, MakeNumber(-1), V1445)

		if True == gen10115 {
			__e.Return(V1446)
			return
		} else {
			if True == True {
				gen10110 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1Xlist_1help)

				gen10111 := Call(__e, V1444, V1443)

				gen10112 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen10113 := Call(__e, gen10112, V1445, V1446)

				__e.TailApply(gen10110, V1443, V1444, gen10111, gen10113)

				return

			} else {
				gen10109 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10109, MakeString("no cond match"))

				return

			}
		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1file_1as_1Xlist_1help, gen10116)

	gen10123 := MakeNative(func(__e *ControlFlow) {
		V1448 := __e.Get(1)
		_ = V1448
		gen10120 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			gen10117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rfas_1h)

			gen10118 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

			gen10119 := Call(__e, gen10118, Stream)

			__e.TailApply(gen10117, Stream, gen10119, MakeString(""))

			return

		}, 1)

		gen10121 := Call(__e, PrimNS1Value(symns2_1value), symopen)

		gen10122 := Call(__e, gen10121, V1448, symin)

		__e.TailApply(gen10120, gen10122)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symread_1file_1as_1string, gen10123)

	gen10135 := MakeNative(func(__e *ControlFlow) {
		V1452 := __e.Get(1)
		_ = V1452
		V1453 := __e.Get(2)
		_ = V1453
		V1454 := __e.Get(3)
		_ = V1454
		gen10133 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10134 := Call(__e, gen10133, MakeNumber(-1), V1453)

		if True == gen10134 {
			gen10132 := Call(__e, PrimNS1Value(symns2_1value), symclose)

			Call(__e, gen10132, V1452)

			__e.Return(V1454)
			return

		} else {
			if True == True {
				gen10125 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rfas_1h)

				gen10126 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

				gen10127 := Call(__e, gen10126, V1452)

				gen10128 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen10129 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				gen10130 := Call(__e, gen10129, V1453)

				gen10131 := Call(__e, gen10128, V1454, gen10130)

				__e.TailApply(gen10125, V1452, gen10127, gen10131)

				return

			} else {
				gen10124 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10124, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4rfas_1h, gen10135)

	gen10139 := MakeNative(func(__e *ControlFlow) {
		V1456 := __e.Get(1)
		_ = V1456
		gen10136 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

		gen10137 := Call(__e, PrimNS1Value(symns2_1value), symread)

		gen10138 := Call(__e, gen10137, V1456)

		__e.TailApply(gen10136, gen10138)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), syminput, gen10139)

	gen10162 := MakeNative(func(__e *ControlFlow) {
		V1459 := __e.Get(1)
		_ = V1459
		V1460 := __e.Get(2)
		_ = V1460
		gen10159 := MakeNative(func(__e *ControlFlow) {
			Mono_2 := __e.Get(1)
			_ = Mono_2
			gen10156 := MakeNative(func(__e *ControlFlow) {
				Input := __e.Get(1)
				_ = Input
				gen10150 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen10151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck)

				gen10152 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				gen10153 := Call(__e, gen10152, V1459)

				gen10154 := Call(__e, gen10151, Input, gen10153)

				gen10155 := Call(__e, gen10150, False, gen10154)

				if True == gen10155 {
					gen10141 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					gen10142 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen10143 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen10144 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen10145 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen10146 := Call(__e, gen10145, V1459, MakeString("\n"), symshen_4r)

					gen10147 := Call(__e, gen10144, MakeString(" is not of type "), gen10146)

					gen10148 := Call(__e, gen10143, Input, gen10147, symshen_4r)

					gen10149 := Call(__e, gen10142, MakeString("type error: "), gen10148)

					__e.TailApply(gen10141, gen10149)

					return

				} else {
					gen10140 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

					__e.TailApply(gen10140, Input)

					return

				}

			}, 1)

			gen10157 := Call(__e, PrimNS1Value(symns2_1value), symread)

			gen10158 := Call(__e, gen10157, V1460)

			__e.TailApply(gen10156, gen10158)

			return

		}, 1)

		gen10160 := Call(__e, PrimNS1Value(symns2_1value), symshen_4monotype)

		gen10161 := Call(__e, gen10160, V1459)

		__e.TailApply(gen10159, gen10161)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), syminput_7, gen10162)

	gen10176 := MakeNative(func(__e *ControlFlow) {
		V1462 := __e.Get(1)
		_ = V1462
		gen10174 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10175 := Call(__e, gen10174, V1462)

		if True == gen10175 {
			gen10171 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			gen10173 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				gen10172 := Call(__e, PrimNS1Value(symns2_1value), symshen_4monotype)

				__e.TailApply(gen10172, Z)

				return

			}, 1)

			__e.TailApply(gen10171, gen10173, V1462)

			return

		} else {
			if True == True {
				gen10169 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				gen10170 := Call(__e, gen10169, V1462)

				if True == gen10170 {
					gen10164 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					gen10165 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen10166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen10167 := Call(__e, gen10166, V1462, MakeString("\n"), symshen_4a)

					gen10168 := Call(__e, gen10165, MakeString("input+ expects a monotype: not "), gen10167)

					__e.TailApply(gen10164, gen10168)

					return

				} else {
					__e.Return(V1462)
					return
				}

			} else {
				gen10163 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10163, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4monotype, gen10176)

	gen10182 := MakeNative(func(__e *ControlFlow) {
		V1464 := __e.Get(1)
		_ = V1464
		gen10177 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10178 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1loop)

		gen10179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

		gen10180 := Call(__e, gen10179, V1464)

		gen10181 := Call(__e, gen10178, V1464, gen10180, Nil)

		__e.TailApply(gen10177, gen10181)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symread, gen10182)

	gen10184 := MakeNative(func(__e *ControlFlow) {
		gen10183 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen10183, symshen_4_dit_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symit, gen10184)

	gen10229 := MakeNative(func(__e *ControlFlow) {
		V1472 := __e.Get(1)
		_ = V1472
		V1473 := __e.Get(2)
		_ = V1473
		V1474 := __e.Get(3)
		_ = V1474
		gen10227 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10228 := Call(__e, gen10227, MakeNumber(94), V1473)

		if True == gen10228 {
			gen10226 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(gen10226, MakeString("read aborted"))

			return

		} else {
			gen10224 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10225 := Call(__e, gen10224, MakeNumber(-1), V1473)

			if True == gen10225 {
				gen10222 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				gen10223 := Call(__e, gen10222, V1474)

				if True == gen10223 {
					gen10221 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen10221, MakeString("error: empty stream"))

					return

				} else {
					gen10217 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

					gen10219 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						gen10218 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

						__e.TailApply(gen10218, X)

						return

					}, 1)

					gen10220 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(E)
						return
					}, 1)

					__e.TailApply(gen10217, gen10219, V1474, gen10220)

					return

				}

			} else {
				gen10215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4terminator_2)

				gen10216 := Call(__e, gen10215, V1473)

				if True == gen10216 {
					gen10210 := MakeNative(func(__e *ControlFlow) {
						AllChars := __e.Get(1)
						_ = AllChars
						gen10207 := MakeNative(func(__e *ControlFlow) {
							It := __e.Get(1)
							_ = It
							gen10201 := MakeNative(func(__e *ControlFlow) {
								Read := __e.Get(1)
								_ = Read
								gen10198 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen10199 := Call(__e, gen10198, Read, symshen_4nextbyte)

								var gen10200 Obj

								if True == gen10199 {
									gen10200 = True
								} else {
									gen10196 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

									gen10197 := Call(__e, gen10196, Read)

									if True == gen10197 {
										gen10200 = True
									} else {
										gen10200 = False
									}

								}

								if True == gen10200 {
									gen10193 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1loop)

									gen10194 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

									gen10195 := Call(__e, gen10194, V1472)

									__e.TailApply(gen10193, V1472, gen10195, AllChars)

									return

								} else {
									__e.Return(Read)
									return
								}

							}, 1)

							gen10202 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

							gen10204 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								gen10203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

								__e.TailApply(gen10203, X)

								return

							}, 1)

							gen10205 := MakeNative(func(__e *ControlFlow) {
								E := __e.Get(1)
								_ = E
								__e.Return(symshen_4nextbyte)
								return
							}, 1)

							gen10206 := Call(__e, gen10202, gen10204, AllChars, gen10205)

							__e.TailApply(gen10201, gen10206)

							return

						}, 1)

						gen10208 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it)

						gen10209 := Call(__e, gen10208, AllChars)

						__e.TailApply(gen10207, gen10209)

						return

					}, 1)

					gen10211 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					gen10212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen10213 := Call(__e, gen10212, V1473, Nil)

					gen10214 := Call(__e, gen10211, V1474, gen10213)

					__e.TailApply(gen10210, gen10214)

					return

				} else {
					if True == True {
						gen10186 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1loop)

						gen10187 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

						gen10188 := Call(__e, gen10187, V1472)

						gen10189 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						gen10190 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen10191 := Call(__e, gen10190, V1473, Nil)

						gen10192 := Call(__e, gen10189, V1474, gen10191)

						__e.TailApply(gen10186, V1472, gen10188, gen10192)

						return

					} else {
						gen10185 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen10185, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1loop, gen10229)

	gen10245 := MakeNative(func(__e *ControlFlow) {
		V1476 := __e.Get(1)
		_ = V1476
		gen10230 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen10231 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10232 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10233 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10234 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10235 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10236 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10237 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10238 := Call(__e, gen10237, MakeNumber(93), Nil)

		gen10239 := Call(__e, gen10236, MakeNumber(41), gen10238)

		gen10240 := Call(__e, gen10235, MakeNumber(34), gen10239)

		gen10241 := Call(__e, gen10234, MakeNumber(32), gen10240)

		gen10242 := Call(__e, gen10233, MakeNumber(13), gen10241)

		gen10243 := Call(__e, gen10232, MakeNumber(10), gen10242)

		gen10244 := Call(__e, gen10231, MakeNumber(9), gen10243)

		__e.TailApply(gen10230, V1476, gen10244)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4terminator_2, gen10245)

	gen10249 := MakeNative(func(__e *ControlFlow) {
		V1478 := __e.Get(1)
		_ = V1478
		gen10246 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lineread_1loop)

		gen10247 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

		gen10248 := Call(__e, gen10247, V1478)

		__e.TailApply(gen10246, gen10248, Nil, V1478)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symlineread, gen10249)

	gen10303 := MakeNative(func(__e *ControlFlow) {
		V1483 := __e.Get(1)
		_ = V1483
		V1484 := __e.Get(2)
		_ = V1484
		V1485 := __e.Get(3)
		_ = V1485
		gen10301 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10302 := Call(__e, gen10301, MakeNumber(-1), V1483)

		if True == gen10302 {
			gen10299 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

			gen10300 := Call(__e, gen10299, V1484)

			if True == gen10300 {
				gen10298 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10298, MakeString("empty stream"))

				return

			} else {
				gen10294 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

				gen10296 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen10295 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

					__e.TailApply(gen10295, X)

					return

				}, 1)

				gen10297 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(E)
					return
				}, 1)

				__e.TailApply(gen10294, gen10296, V1484, gen10297)

				return

			}

		} else {
			gen10290 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10291 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hat)

			gen10292 := Call(__e, gen10291)

			gen10293 := Call(__e, gen10290, V1483, gen10292)

			if True == gen10293 {
				gen10289 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10289, MakeString("line read aborted"))

				return

			} else {
				gen10279 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen10280 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen10281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

				gen10282 := Call(__e, gen10281)

				gen10283 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen10284 := Call(__e, PrimNS1Value(symns2_1value), symshen_4carriage_1return)

				gen10285 := Call(__e, gen10284)

				gen10286 := Call(__e, gen10283, gen10285, Nil)

				gen10287 := Call(__e, gen10280, gen10282, gen10286)

				gen10288 := Call(__e, gen10279, V1483, gen10287)

				if True == gen10288 {
					gen10273 := MakeNative(func(__e *ControlFlow) {
						Line := __e.Get(1)
						_ = Line
						gen10270 := MakeNative(func(__e *ControlFlow) {
							It := __e.Get(1)
							_ = It
							gen10267 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen10268 := Call(__e, gen10267, Line, symshen_4nextline)

							var gen10269 Obj

							if True == gen10268 {
								gen10269 = True
							} else {
								gen10265 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

								gen10266 := Call(__e, gen10265, Line)

								if True == gen10266 {
									gen10269 = True
								} else {
									gen10269 = False
								}

							}

							if True == gen10269 {
								gen10258 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lineread_1loop)

								gen10259 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

								gen10260 := Call(__e, gen10259, V1485)

								gen10261 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								gen10262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen10263 := Call(__e, gen10262, V1483, Nil)

								gen10264 := Call(__e, gen10261, V1484, gen10263)

								__e.TailApply(gen10258, gen10260, gen10264, V1485)

								return

							} else {
								__e.Return(Line)
								return
							}

						}, 1)

						gen10271 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it)

						gen10272 := Call(__e, gen10271, V1484)

						__e.TailApply(gen10270, gen10272)

						return

					}, 1)

					gen10274 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

					gen10276 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						gen10275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

						__e.TailApply(gen10275, X)

						return

					}, 1)

					gen10277 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(symshen_4nextline)
						return
					}, 1)

					gen10278 := Call(__e, gen10274, gen10276, V1484, gen10277)

					__e.TailApply(gen10273, gen10278)

					return

				} else {
					if True == True {
						gen10251 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lineread_1loop)

						gen10252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

						gen10253 := Call(__e, gen10252, V1485)

						gen10254 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						gen10255 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen10256 := Call(__e, gen10255, V1483, Nil)

						gen10257 := Call(__e, gen10254, V1484, gen10256)

						__e.TailApply(gen10251, gen10253, gen10257, V1485)

						return

					} else {
						gen10250 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen10250, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4lineread_1loop, gen10303)

	gen10316 := MakeNative(func(__e *ControlFlow) {
		V1487 := __e.Get(1)
		_ = V1487
		gen10313 := MakeNative(func(__e *ControlFlow) {
			TrimLeft := __e.Get(1)
			_ = TrimLeft
			gen10308 := MakeNative(func(__e *ControlFlow) {
				TrimRight := __e.Get(1)
				_ = TrimRight
				gen10305 := MakeNative(func(__e *ControlFlow) {
					Trimmed := __e.Get(1)
					_ = Trimmed
					gen10304 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it_1h)

					__e.TailApply(gen10304, Trimmed)

					return

				}, 1)

				gen10306 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

				gen10307 := Call(__e, gen10306, TrimRight)

				__e.TailApply(gen10305, gen10307)

				return

			}, 1)

			gen10309 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1whitespace)

			gen10310 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			gen10311 := Call(__e, gen10310, TrimLeft)

			gen10312 := Call(__e, gen10309, gen10311)

			__e.TailApply(gen10308, gen10312)

			return

		}, 1)

		gen10314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1whitespace)

		gen10315 := Call(__e, gen10314, V1487)

		__e.TailApply(gen10313, gen10315)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1it, gen10316)

	gen10336 := MakeNative(func(__e *ControlFlow) {
		V1489 := __e.Get(1)
		_ = V1489
		gen10333 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10334 := Call(__e, gen10333, V1489)

		var gen10335 Obj

		if True == gen10334 {
			gen10321 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

			gen10322 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen10323 := Call(__e, gen10322, V1489)

			gen10324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10327 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10328 := Call(__e, gen10327, MakeNumber(32), Nil)

			gen10329 := Call(__e, gen10326, MakeNumber(13), gen10328)

			gen10330 := Call(__e, gen10325, MakeNumber(10), gen10329)

			gen10331 := Call(__e, gen10324, MakeNumber(9), gen10330)

			gen10332 := Call(__e, gen10321, gen10323, gen10331)

			if True == gen10332 {
				gen10335 = True
			} else {
				gen10335 = False
			}

		} else {
			gen10335 = False
		}

		if True == gen10335 {
			gen10318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1whitespace)

			gen10319 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen10320 := Call(__e, gen10319, V1489)

			__e.TailApply(gen10318, gen10320)

			return

		} else {
			if True == True {
				__e.Return(V1489)
				return
			} else {
				gen10317 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10317, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4trim_1whitespace, gen10336)

	gen10344 := MakeNative(func(__e *ControlFlow) {
		V1491 := __e.Get(1)
		_ = V1491
		gen10337 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen10338 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cn_1all)

		gen10339 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen10341 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen10340 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			__e.TailApply(gen10340, X)

			return

		}, 1)

		gen10342 := Call(__e, gen10339, gen10341, V1491)

		gen10343 := Call(__e, gen10338, gen10342)

		Call(__e, gen10337, symshen_4_dit_d, gen10343)

		__e.Return(V1491)
		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1it_1h, gen10344)

	gen10358 := MakeNative(func(__e *ControlFlow) {
		V1493 := __e.Get(1)
		_ = V1493
		gen10356 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10357 := Call(__e, gen10356, Nil, V1493)

		if True == gen10357 {
			__e.Return(MakeString(""))
			return
		} else {
			gen10354 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen10355 := Call(__e, gen10354, V1493)

			if True == gen10355 {
				gen10347 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen10348 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen10349 := Call(__e, gen10348, V1493)

				gen10350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cn_1all)

				gen10351 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10352 := Call(__e, gen10351, V1493)

				gen10353 := Call(__e, gen10350, gen10352)

				__e.TailApply(gen10347, gen10349, gen10353)

				return

			} else {
				if True == True {
					gen10346 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen10346, symshen_4cn_1all)

					return

				} else {
					gen10345 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen10345, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4cn_1all, gen10358)

	gen10367 := MakeNative(func(__e *ControlFlow) {
		V1495 := __e.Get(1)
		_ = V1495
		gen10364 := MakeNative(func(__e *ControlFlow) {
			Charlist := __e.Get(1)
			_ = Charlist
			gen10359 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			gen10361 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen10360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

				__e.TailApply(gen10360, X)

				return

			}, 1)

			gen10363 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen10362 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1error)

				__e.TailApply(gen10362, X)

				return

			}, 1)

			__e.TailApply(gen10359, gen10361, Charlist, gen10363)

			return

		}, 1)

		gen10365 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1file_1as_1charlist)

		gen10366 := Call(__e, gen10365, V1495)

		__e.TailApply(gen10364, gen10366)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symread_1file, gen10367)

	gen10380 := MakeNative(func(__e *ControlFlow) {
		V1497 := __e.Get(1)
		_ = V1497
		gen10373 := MakeNative(func(__e *ControlFlow) {
			Ns := __e.Get(1)
			_ = Ns
			gen10368 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			gen10370 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen10369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

				__e.TailApply(gen10369, X)

				return

			}, 1)

			gen10372 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen10371 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1error)

				__e.TailApply(gen10371, X)

				return

			}, 1)

			__e.TailApply(gen10368, gen10370, Ns, gen10372)

			return

		}, 1)

		gen10374 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen10376 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen10375 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

			__e.TailApply(gen10375, X)

			return

		}, 1)

		gen10377 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

		gen10378 := Call(__e, gen10377, V1497)

		gen10379 := Call(__e, gen10374, gen10376, gen10378)

		__e.TailApply(gen10373, gen10379)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symread_1from_1string, gen10380)

	gen10411 := MakeNative(func(__e *ControlFlow) {
		V1505 := __e.Get(1)
		_ = V1505
		gen10408 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10409 := Call(__e, gen10408, V1505)

		var gen10410 Obj

		if True == gen10409 {
			gen10403 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen10404 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen10405 := Call(__e, gen10404, V1505)

			gen10406 := Call(__e, gen10403, gen10405)

			var gen10407 Obj

			if True == gen10406 {
				gen10398 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen10399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10400 := Call(__e, gen10399, V1505)

				gen10401 := Call(__e, gen10398, gen10400)

				var gen10402 Obj

				if True == gen10401 {
					gen10392 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen10393 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen10394 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen10395 := Call(__e, gen10394, V1505)

					gen10396 := Call(__e, gen10393, gen10395)

					gen10397 := Call(__e, gen10392, Nil, gen10396)

					if True == gen10397 {
						gen10402 = True
					} else {
						gen10402 = False
					}

				} else {
					gen10402 = False
				}

				if True == gen10402 {
					gen10407 = True
				} else {
					gen10407 = False
				}

			} else {
				gen10407 = False
			}

			if True == gen10407 {
				gen10410 = True
			} else {
				gen10410 = False
			}

		} else {
			gen10410 = False
		}

		if True == gen10410 {
			gen10383 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen10384 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen10385 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen10386 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compress_150)

			gen10387 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen10388 := Call(__e, gen10387, V1505)

			gen10389 := Call(__e, gen10386, MakeNumber(50), gen10388)

			gen10390 := Call(__e, gen10385, gen10389, MakeString("\n"), symshen_4a)

			gen10391 := Call(__e, gen10384, MakeString("read error here:\n\n "), gen10390)

			__e.TailApply(gen10383, gen10391)

			return

		} else {
			if True == True {
				gen10382 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10382, MakeString("read error\n"))

				return

			} else {
				gen10381 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10381, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1error, gen10411)

	gen10431 := MakeNative(func(__e *ControlFlow) {
		V1512 := __e.Get(1)
		_ = V1512
		V1513 := __e.Get(2)
		_ = V1513
		gen10429 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10430 := Call(__e, gen10429, Nil, V1513)

		if True == gen10430 {
			__e.Return(MakeString(""))
			return
		} else {
			gen10427 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10428 := Call(__e, gen10427, MakeNumber(0), V1512)

			if True == gen10428 {
				__e.Return(MakeString(""))
				return
			} else {
				gen10425 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen10426 := Call(__e, gen10425, V1513)

				if True == gen10426 {
					gen10414 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen10415 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					gen10416 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen10417 := Call(__e, gen10416, V1513)

					gen10418 := Call(__e, gen10415, gen10417)

					gen10419 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compress_150)

					gen10420 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					gen10421 := Call(__e, gen10420, V1512, MakeNumber(1))

					gen10422 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen10423 := Call(__e, gen10422, V1513)

					gen10424 := Call(__e, gen10419, gen10421, gen10423)

					__e.TailApply(gen10414, gen10418, gen10424)

					return

				} else {
					if True == True {
						gen10413 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen10413, symshen_4compress_150)

						return

					} else {
						gen10412 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen10412, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4compress_150, gen10431)

	gen10946 := MakeNative(func(__e *ControlFlow) {
		V1515 := __e.Get(1)
		_ = V1515
		gen10891 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen10887 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10888 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen10889 := Call(__e, gen10888)

			gen10890 := Call(__e, gen10887, YaccParse, gen10889)

			if True == gen10890 {
				gen10834 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					gen10830 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen10831 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen10832 := Call(__e, gen10831)

					gen10833 := Call(__e, gen10830, YaccParse, gen10832)

					if True == gen10833 {
						gen10801 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							gen10797 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen10798 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen10799 := Call(__e, gen10798)

							gen10800 := Call(__e, gen10797, YaccParse, gen10799)

							if True == gen10800 {
								gen10768 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									gen10764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen10765 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen10766 := Call(__e, gen10765)

									gen10767 := Call(__e, gen10764, YaccParse, gen10766)

									if True == gen10767 {
										gen10735 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											gen10731 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen10732 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen10733 := Call(__e, gen10732)

											gen10734 := Call(__e, gen10731, YaccParse, gen10733)

											if True == gen10734 {
												gen10702 := MakeNative(func(__e *ControlFlow) {
													YaccParse := __e.Get(1)
													_ = YaccParse
													gen10698 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen10699 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen10700 := Call(__e, gen10699)

													gen10701 := Call(__e, gen10698, YaccParse, gen10700)

													if True == gen10701 {
														gen10659 := MakeNative(func(__e *ControlFlow) {
															YaccParse := __e.Get(1)
															_ = YaccParse
															gen10655 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															gen10656 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															gen10657 := Call(__e, gen10656)

															gen10658 := Call(__e, gen10655, YaccParse, gen10657)

															if True == gen10658 {
																gen10616 := MakeNative(func(__e *ControlFlow) {
																	YaccParse := __e.Get(1)
																	_ = YaccParse
																	gen10612 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	gen10613 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																	gen10614 := Call(__e, gen10613)

																	gen10615 := Call(__e, gen10612, YaccParse, gen10614)

																	if True == gen10615 {
																		gen10583 := MakeNative(func(__e *ControlFlow) {
																			YaccParse := __e.Get(1)
																			_ = YaccParse
																			gen10579 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			gen10580 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			gen10581 := Call(__e, gen10580)

																			gen10582 := Call(__e, gen10579, YaccParse, gen10581)

																			if True == gen10582 {
																				gen10548 := MakeNative(func(__e *ControlFlow) {
																					YaccParse := __e.Get(1)
																					_ = YaccParse
																					gen10544 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					gen10545 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					gen10546 := Call(__e, gen10545)

																					gen10547 := Call(__e, gen10544, YaccParse, gen10546)

																					if True == gen10547 {
																						gen10517 := MakeNative(func(__e *ControlFlow) {
																							YaccParse := __e.Get(1)
																							_ = YaccParse
																							gen10513 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							gen10514 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																							gen10515 := Call(__e, gen10514)

																							gen10516 := Call(__e, gen10513, YaccParse, gen10515)

																							if True == gen10516 {
																								gen10480 := MakeNative(func(__e *ControlFlow) {
																									YaccParse := __e.Get(1)
																									_ = YaccParse
																									gen10476 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									gen10477 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																									gen10478 := Call(__e, gen10477)

																									gen10479 := Call(__e, gen10476, YaccParse, gen10478)

																									if True == gen10479 {
																										gen10449 := MakeNative(func(__e *ControlFlow) {
																											YaccParse := __e.Get(1)
																											_ = YaccParse
																											gen10445 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											gen10446 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																											gen10447 := Call(__e, gen10446)

																											gen10448 := Call(__e, gen10445, YaccParse, gen10447)

																											if True == gen10448 {
																												gen10442 := MakeNative(func(__e *ControlFlow) {
																													Parse___5e_6 := __e.Get(1)
																													_ = Parse___5e_6
																													gen10436 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																													gen10437 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																													gen10438 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																													gen10439 := Call(__e, gen10438)

																													gen10440 := Call(__e, gen10437, gen10439, Parse___5e_6)

																													gen10441 := Call(__e, gen10436, gen10440)

																													if True == gen10441 {
																														gen10433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																														gen10434 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														gen10435 := Call(__e, gen10434, Parse___5e_6)

																														__e.TailApply(gen10433, gen10435, Nil)

																														return

																													} else {
																														gen10432 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																														__e.TailApply(gen10432)

																														return

																													}

																												}, 1)

																												gen10443 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

																												gen10444 := Call(__e, gen10443, V1515)

																												__e.TailApply(gen10442, gen10444)

																												return

																											} else {
																												__e.Return(YaccParse)
																												return
																											}

																										}, 1)

																										gen10472 := MakeNative(func(__e *ControlFlow) {
																											Parse__shen_4_5whitespaces_6 := __e.Get(1)
																											_ = Parse__shen_4_5whitespaces_6
																											gen10466 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																											gen10467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											gen10468 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																											gen10469 := Call(__e, gen10468)

																											gen10470 := Call(__e, gen10467, gen10469, Parse__shen_4_5whitespaces_6)

																											gen10471 := Call(__e, gen10466, gen10470)

																											if True == gen10471 {
																												gen10463 := MakeNative(func(__e *ControlFlow) {
																													Parse__shen_4_5st__input_6 := __e.Get(1)
																													_ = Parse__shen_4_5st__input_6
																													gen10457 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																													gen10458 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																													gen10459 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																													gen10460 := Call(__e, gen10459)

																													gen10461 := Call(__e, gen10458, gen10460, Parse__shen_4_5st__input_6)

																													gen10462 := Call(__e, gen10457, gen10461)

																													if True == gen10462 {
																														gen10452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																														gen10453 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														gen10454 := Call(__e, gen10453, Parse__shen_4_5st__input_6)

																														gen10455 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																														gen10456 := Call(__e, gen10455, Parse__shen_4_5st__input_6)

																														__e.TailApply(gen10452, gen10454, gen10456)

																														return

																													} else {
																														gen10451 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																														__e.TailApply(gen10451)

																														return

																													}

																												}, 1)

																												gen10464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																												gen10465 := Call(__e, gen10464, Parse__shen_4_5whitespaces_6)

																												__e.TailApply(gen10463, gen10465)

																												return

																											} else {
																												gen10450 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																												__e.TailApply(gen10450)

																												return

																											}

																										}, 1)

																										gen10473 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespaces_6)

																										gen10474 := Call(__e, gen10473, V1515)

																										gen10475 := Call(__e, gen10472, gen10474)

																										__e.TailApply(gen10449, gen10475)

																										return

																									} else {
																										__e.Return(YaccParse)
																										return
																									}

																								}, 1)

																								gen10509 := MakeNative(func(__e *ControlFlow) {
																									Parse__shen_4_5atom_6 := __e.Get(1)
																									_ = Parse__shen_4_5atom_6
																									gen10503 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																									gen10504 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									gen10505 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																									gen10506 := Call(__e, gen10505)

																									gen10507 := Call(__e, gen10504, gen10506, Parse__shen_4_5atom_6)

																									gen10508 := Call(__e, gen10503, gen10507)

																									if True == gen10508 {
																										gen10500 := MakeNative(func(__e *ControlFlow) {
																											Parse__shen_4_5st__input_6 := __e.Get(1)
																											_ = Parse__shen_4_5st__input_6
																											gen10494 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																											gen10495 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											gen10496 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																											gen10497 := Call(__e, gen10496)

																											gen10498 := Call(__e, gen10495, gen10497, Parse__shen_4_5st__input_6)

																											gen10499 := Call(__e, gen10494, gen10498)

																											if True == gen10499 {
																												gen10483 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																												gen10484 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																												gen10485 := Call(__e, gen10484, Parse__shen_4_5st__input_6)

																												gen10486 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																												gen10487 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

																												gen10488 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																												gen10489 := Call(__e, gen10488, Parse__shen_4_5atom_6)

																												gen10490 := Call(__e, gen10487, gen10489)

																												gen10491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																												gen10492 := Call(__e, gen10491, Parse__shen_4_5st__input_6)

																												gen10493 := Call(__e, gen10486, gen10490, gen10492)

																												__e.TailApply(gen10483, gen10485, gen10493)

																												return

																											} else {
																												gen10482 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																												__e.TailApply(gen10482)

																												return

																											}

																										}, 1)

																										gen10501 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																										gen10502 := Call(__e, gen10501, Parse__shen_4_5atom_6)

																										__e.TailApply(gen10500, gen10502)

																										return

																									} else {
																										gen10481 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																										__e.TailApply(gen10481)

																										return

																									}

																								}, 1)

																								gen10510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5atom_6)

																								gen10511 := Call(__e, gen10510, V1515)

																								gen10512 := Call(__e, gen10509, gen10511)

																								__e.TailApply(gen10480, gen10512)

																								return

																							} else {
																								__e.Return(YaccParse)
																								return
																							}

																						}, 1)

																						gen10540 := MakeNative(func(__e *ControlFlow) {
																							Parse__shen_4_5comment_6 := __e.Get(1)
																							_ = Parse__shen_4_5comment_6
																							gen10534 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																							gen10535 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							gen10536 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																							gen10537 := Call(__e, gen10536)

																							gen10538 := Call(__e, gen10535, gen10537, Parse__shen_4_5comment_6)

																							gen10539 := Call(__e, gen10534, gen10538)

																							if True == gen10539 {
																								gen10531 := MakeNative(func(__e *ControlFlow) {
																									Parse__shen_4_5st__input_6 := __e.Get(1)
																									_ = Parse__shen_4_5st__input_6
																									gen10525 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																									gen10526 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									gen10527 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																									gen10528 := Call(__e, gen10527)

																									gen10529 := Call(__e, gen10526, gen10528, Parse__shen_4_5st__input_6)

																									gen10530 := Call(__e, gen10525, gen10529)

																									if True == gen10530 {
																										gen10520 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																										gen10521 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																										gen10522 := Call(__e, gen10521, Parse__shen_4_5st__input_6)

																										gen10523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																										gen10524 := Call(__e, gen10523, Parse__shen_4_5st__input_6)

																										__e.TailApply(gen10520, gen10522, gen10524)

																										return

																									} else {
																										gen10519 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																										__e.TailApply(gen10519)

																										return

																									}

																								}, 1)

																								gen10532 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																								gen10533 := Call(__e, gen10532, Parse__shen_4_5comment_6)

																								__e.TailApply(gen10531, gen10533)

																								return

																							} else {
																								gen10518 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																								__e.TailApply(gen10518)

																								return

																							}

																						}, 1)

																						gen10541 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comment_6)

																						gen10542 := Call(__e, gen10541, V1515)

																						gen10543 := Call(__e, gen10540, gen10542)

																						__e.TailApply(gen10517, gen10543)

																						return

																					} else {
																						__e.Return(YaccParse)
																						return
																					}

																				}, 1)

																				gen10575 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5comma_6 := __e.Get(1)
																					_ = Parse__shen_4_5comma_6
																					gen10569 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																					gen10570 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					gen10571 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					gen10572 := Call(__e, gen10571)

																					gen10573 := Call(__e, gen10570, gen10572, Parse__shen_4_5comma_6)

																					gen10574 := Call(__e, gen10569, gen10573)

																					if True == gen10574 {
																						gen10566 := MakeNative(func(__e *ControlFlow) {
																							Parse__shen_4_5st__input_6 := __e.Get(1)
																							_ = Parse__shen_4_5st__input_6
																							gen10560 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																							gen10561 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							gen10562 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																							gen10563 := Call(__e, gen10562)

																							gen10564 := Call(__e, gen10561, gen10563, Parse__shen_4_5st__input_6)

																							gen10565 := Call(__e, gen10560, gen10564)

																							if True == gen10565 {
																								gen10551 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																								gen10552 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																								gen10553 := Call(__e, gen10552, Parse__shen_4_5st__input_6)

																								gen10554 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																								gen10555 := Call(__e, PrimNS1Value(symns2_1value), symintern)

																								gen10556 := Call(__e, gen10555, MakeString(","))

																								gen10557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																								gen10558 := Call(__e, gen10557, Parse__shen_4_5st__input_6)

																								gen10559 := Call(__e, gen10554, gen10556, gen10558)

																								__e.TailApply(gen10551, gen10553, gen10559)

																								return

																							} else {
																								gen10550 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																								__e.TailApply(gen10550)

																								return

																							}

																						}, 1)

																						gen10567 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																						gen10568 := Call(__e, gen10567, Parse__shen_4_5comma_6)

																						__e.TailApply(gen10566, gen10568)

																						return

																					} else {
																						gen10549 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																						__e.TailApply(gen10549)

																						return

																					}

																				}, 1)

																				gen10576 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comma_6)

																				gen10577 := Call(__e, gen10576, V1515)

																				gen10578 := Call(__e, gen10575, gen10577)

																				__e.TailApply(gen10548, gen10578)

																				return

																			} else {
																				__e.Return(YaccParse)
																				return
																			}

																		}, 1)

																		gen10608 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5colon_6 := __e.Get(1)
																			_ = Parse__shen_4_5colon_6
																			gen10602 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																			gen10603 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			gen10604 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			gen10605 := Call(__e, gen10604)

																			gen10606 := Call(__e, gen10603, gen10605, Parse__shen_4_5colon_6)

																			gen10607 := Call(__e, gen10602, gen10606)

																			if True == gen10607 {
																				gen10599 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5st__input_6 := __e.Get(1)
																					_ = Parse__shen_4_5st__input_6
																					gen10593 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																					gen10594 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					gen10595 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					gen10596 := Call(__e, gen10595)

																					gen10597 := Call(__e, gen10594, gen10596, Parse__shen_4_5st__input_6)

																					gen10598 := Call(__e, gen10593, gen10597)

																					if True == gen10598 {
																						gen10586 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																						gen10587 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																						gen10588 := Call(__e, gen10587, Parse__shen_4_5st__input_6)

																						gen10589 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						gen10590 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																						gen10591 := Call(__e, gen10590, Parse__shen_4_5st__input_6)

																						gen10592 := Call(__e, gen10589, sym_h, gen10591)

																						__e.TailApply(gen10586, gen10588, gen10592)

																						return

																					} else {
																						gen10585 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																						__e.TailApply(gen10585)

																						return

																					}

																				}, 1)

																				gen10600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																				gen10601 := Call(__e, gen10600, Parse__shen_4_5colon_6)

																				__e.TailApply(gen10599, gen10601)

																				return

																			} else {
																				gen10584 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																				__e.TailApply(gen10584)

																				return

																			}

																		}, 1)

																		gen10609 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5colon_6)

																		gen10610 := Call(__e, gen10609, V1515)

																		gen10611 := Call(__e, gen10608, gen10610)

																		__e.TailApply(gen10583, gen10611)

																		return

																	} else {
																		__e.Return(YaccParse)
																		return
																	}

																}, 1)

																gen10651 := MakeNative(func(__e *ControlFlow) {
																	Parse__shen_4_5colon_6 := __e.Get(1)
																	_ = Parse__shen_4_5colon_6
																	gen10645 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																	gen10646 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	gen10647 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																	gen10648 := Call(__e, gen10647)

																	gen10649 := Call(__e, gen10646, gen10648, Parse__shen_4_5colon_6)

																	gen10650 := Call(__e, gen10645, gen10649)

																	if True == gen10650 {
																		gen10642 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5minus_6 := __e.Get(1)
																			_ = Parse__shen_4_5minus_6
																			gen10636 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																			gen10637 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			gen10638 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			gen10639 := Call(__e, gen10638)

																			gen10640 := Call(__e, gen10637, gen10639, Parse__shen_4_5minus_6)

																			gen10641 := Call(__e, gen10636, gen10640)

																			if True == gen10641 {
																				gen10633 := MakeNative(func(__e *ControlFlow) {
																					Parse__shen_4_5st__input_6 := __e.Get(1)
																					_ = Parse__shen_4_5st__input_6
																					gen10627 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																					gen10628 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					gen10629 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																					gen10630 := Call(__e, gen10629)

																					gen10631 := Call(__e, gen10628, gen10630, Parse__shen_4_5st__input_6)

																					gen10632 := Call(__e, gen10627, gen10631)

																					if True == gen10632 {
																						gen10620 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																						gen10621 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																						gen10622 := Call(__e, gen10621, Parse__shen_4_5st__input_6)

																						gen10623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						gen10624 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																						gen10625 := Call(__e, gen10624, Parse__shen_4_5st__input_6)

																						gen10626 := Call(__e, gen10623, sym_h_1, gen10625)

																						__e.TailApply(gen10620, gen10622, gen10626)

																						return

																					} else {
																						gen10619 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																						__e.TailApply(gen10619)

																						return

																					}

																				}, 1)

																				gen10634 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																				gen10635 := Call(__e, gen10634, Parse__shen_4_5minus_6)

																				__e.TailApply(gen10633, gen10635)

																				return

																			} else {
																				gen10618 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																				__e.TailApply(gen10618)

																				return

																			}

																		}, 1)

																		gen10643 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5minus_6)

																		gen10644 := Call(__e, gen10643, Parse__shen_4_5colon_6)

																		__e.TailApply(gen10642, gen10644)

																		return

																	} else {
																		gen10617 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																		__e.TailApply(gen10617)

																		return

																	}

																}, 1)

																gen10652 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5colon_6)

																gen10653 := Call(__e, gen10652, V1515)

																gen10654 := Call(__e, gen10651, gen10653)

																__e.TailApply(gen10616, gen10654)

																return

															} else {
																__e.Return(YaccParse)
																return
															}

														}, 1)

														gen10694 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5colon_6 := __e.Get(1)
															_ = Parse__shen_4_5colon_6
															gen10688 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															gen10689 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															gen10690 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															gen10691 := Call(__e, gen10690)

															gen10692 := Call(__e, gen10689, gen10691, Parse__shen_4_5colon_6)

															gen10693 := Call(__e, gen10688, gen10692)

															if True == gen10693 {
																gen10685 := MakeNative(func(__e *ControlFlow) {
																	Parse__shen_4_5equal_6 := __e.Get(1)
																	_ = Parse__shen_4_5equal_6
																	gen10679 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																	gen10680 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	gen10681 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																	gen10682 := Call(__e, gen10681)

																	gen10683 := Call(__e, gen10680, gen10682, Parse__shen_4_5equal_6)

																	gen10684 := Call(__e, gen10679, gen10683)

																	if True == gen10684 {
																		gen10676 := MakeNative(func(__e *ControlFlow) {
																			Parse__shen_4_5st__input_6 := __e.Get(1)
																			_ = Parse__shen_4_5st__input_6
																			gen10670 := Call(__e, PrimNS1Value(symns2_1value), symnot)

																			gen10671 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			gen10672 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																			gen10673 := Call(__e, gen10672)

																			gen10674 := Call(__e, gen10671, gen10673, Parse__shen_4_5st__input_6)

																			gen10675 := Call(__e, gen10670, gen10674)

																			if True == gen10675 {
																				gen10663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																				gen10664 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																				gen10665 := Call(__e, gen10664, Parse__shen_4_5st__input_6)

																				gen10666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																				gen10667 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																				gen10668 := Call(__e, gen10667, Parse__shen_4_5st__input_6)

																				gen10669 := Call(__e, gen10666, sym_h_a, gen10668)

																				__e.TailApply(gen10663, gen10665, gen10669)

																				return

																			} else {
																				gen10662 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																				__e.TailApply(gen10662)

																				return

																			}

																		}, 1)

																		gen10677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

																		gen10678 := Call(__e, gen10677, Parse__shen_4_5equal_6)

																		__e.TailApply(gen10676, gen10678)

																		return

																	} else {
																		gen10661 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																		__e.TailApply(gen10661)

																		return

																	}

																}, 1)

																gen10686 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5equal_6)

																gen10687 := Call(__e, gen10686, Parse__shen_4_5colon_6)

																__e.TailApply(gen10685, gen10687)

																return

															} else {
																gen10660 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(gen10660)

																return

															}

														}, 1)

														gen10695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5colon_6)

														gen10696 := Call(__e, gen10695, V1515)

														gen10697 := Call(__e, gen10694, gen10696)

														__e.TailApply(gen10659, gen10697)

														return

													} else {
														__e.Return(YaccParse)
														return
													}

												}, 1)

												gen10727 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5semicolon_6 := __e.Get(1)
													_ = Parse__shen_4_5semicolon_6
													gen10721 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													gen10722 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen10723 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen10724 := Call(__e, gen10723)

													gen10725 := Call(__e, gen10722, gen10724, Parse__shen_4_5semicolon_6)

													gen10726 := Call(__e, gen10721, gen10725)

													if True == gen10726 {
														gen10718 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5st__input_6 := __e.Get(1)
															_ = Parse__shen_4_5st__input_6
															gen10712 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															gen10713 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															gen10714 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															gen10715 := Call(__e, gen10714)

															gen10716 := Call(__e, gen10713, gen10715, Parse__shen_4_5st__input_6)

															gen10717 := Call(__e, gen10712, gen10716)

															if True == gen10717 {
																gen10705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																gen10706 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																gen10707 := Call(__e, gen10706, Parse__shen_4_5st__input_6)

																gen10708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																gen10709 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																gen10710 := Call(__e, gen10709, Parse__shen_4_5st__input_6)

																gen10711 := Call(__e, gen10708, sym_k, gen10710)

																__e.TailApply(gen10705, gen10707, gen10711)

																return

															} else {
																gen10704 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(gen10704)

																return

															}

														}, 1)

														gen10719 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

														gen10720 := Call(__e, gen10719, Parse__shen_4_5semicolon_6)

														__e.TailApply(gen10718, gen10720)

														return

													} else {
														gen10703 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(gen10703)

														return

													}

												}, 1)

												gen10728 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_6)

												gen10729 := Call(__e, gen10728, V1515)

												gen10730 := Call(__e, gen10727, gen10729)

												__e.TailApply(gen10702, gen10730)

												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										gen10760 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5bar_6 := __e.Get(1)
											_ = Parse__shen_4_5bar_6
											gen10754 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen10755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen10756 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen10757 := Call(__e, gen10756)

											gen10758 := Call(__e, gen10755, gen10757, Parse__shen_4_5bar_6)

											gen10759 := Call(__e, gen10754, gen10758)

											if True == gen10759 {
												gen10751 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5st__input_6 := __e.Get(1)
													_ = Parse__shen_4_5st__input_6
													gen10745 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													gen10746 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen10747 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen10748 := Call(__e, gen10747)

													gen10749 := Call(__e, gen10746, gen10748, Parse__shen_4_5st__input_6)

													gen10750 := Call(__e, gen10745, gen10749)

													if True == gen10750 {
														gen10738 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														gen10739 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen10740 := Call(__e, gen10739, Parse__shen_4_5st__input_6)

														gen10741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														gen10742 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														gen10743 := Call(__e, gen10742, Parse__shen_4_5st__input_6)

														gen10744 := Call(__e, gen10741, symbar_b, gen10743)

														__e.TailApply(gen10738, gen10740, gen10744)

														return

													} else {
														gen10737 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(gen10737)

														return

													}

												}, 1)

												gen10752 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

												gen10753 := Call(__e, gen10752, Parse__shen_4_5bar_6)

												__e.TailApply(gen10751, gen10753)

												return

											} else {
												gen10736 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen10736)

												return

											}

										}, 1)

										gen10761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5bar_6)

										gen10762 := Call(__e, gen10761, V1515)

										gen10763 := Call(__e, gen10760, gen10762)

										__e.TailApply(gen10735, gen10763)

										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								gen10793 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5rcurly_6 := __e.Get(1)
									_ = Parse__shen_4_5rcurly_6
									gen10787 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen10788 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen10789 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen10790 := Call(__e, gen10789)

									gen10791 := Call(__e, gen10788, gen10790, Parse__shen_4_5rcurly_6)

									gen10792 := Call(__e, gen10787, gen10791)

									if True == gen10792 {
										gen10784 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5st__input_6 := __e.Get(1)
											_ = Parse__shen_4_5st__input_6
											gen10778 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen10779 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen10780 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen10781 := Call(__e, gen10780)

											gen10782 := Call(__e, gen10779, gen10781, Parse__shen_4_5st__input_6)

											gen10783 := Call(__e, gen10778, gen10782)

											if True == gen10783 {
												gen10771 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												gen10772 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen10773 := Call(__e, gen10772, Parse__shen_4_5st__input_6)

												gen10774 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												gen10775 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												gen10776 := Call(__e, gen10775, Parse__shen_4_5st__input_6)

												gen10777 := Call(__e, gen10774, sym_j, gen10776)

												__e.TailApply(gen10771, gen10773, gen10777)

												return

											} else {
												gen10770 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen10770)

												return

											}

										}, 1)

										gen10785 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

										gen10786 := Call(__e, gen10785, Parse__shen_4_5rcurly_6)

										__e.TailApply(gen10784, gen10786)

										return

									} else {
										gen10769 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen10769)

										return

									}

								}, 1)

								gen10794 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rcurly_6)

								gen10795 := Call(__e, gen10794, V1515)

								gen10796 := Call(__e, gen10793, gen10795)

								__e.TailApply(gen10768, gen10796)

								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						gen10826 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5lcurly_6 := __e.Get(1)
							_ = Parse__shen_4_5lcurly_6
							gen10820 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen10821 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen10822 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen10823 := Call(__e, gen10822)

							gen10824 := Call(__e, gen10821, gen10823, Parse__shen_4_5lcurly_6)

							gen10825 := Call(__e, gen10820, gen10824)

							if True == gen10825 {
								gen10817 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5st__input_6 := __e.Get(1)
									_ = Parse__shen_4_5st__input_6
									gen10811 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen10812 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen10813 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen10814 := Call(__e, gen10813)

									gen10815 := Call(__e, gen10812, gen10814, Parse__shen_4_5st__input_6)

									gen10816 := Call(__e, gen10811, gen10815)

									if True == gen10816 {
										gen10804 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen10805 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen10806 := Call(__e, gen10805, Parse__shen_4_5st__input_6)

										gen10807 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen10808 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen10809 := Call(__e, gen10808, Parse__shen_4_5st__input_6)

										gen10810 := Call(__e, gen10807, sym_i, gen10809)

										__e.TailApply(gen10804, gen10806, gen10810)

										return

									} else {
										gen10803 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen10803)

										return

									}

								}, 1)

								gen10818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

								gen10819 := Call(__e, gen10818, Parse__shen_4_5lcurly_6)

								__e.TailApply(gen10817, gen10819)

								return

							} else {
								gen10802 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen10802)

								return

							}

						}, 1)

						gen10827 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5lcurly_6)

						gen10828 := Call(__e, gen10827, V1515)

						gen10829 := Call(__e, gen10826, gen10828)

						__e.TailApply(gen10801, gen10829)

						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				gen10883 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5lrb_6 := __e.Get(1)
					_ = Parse__shen_4_5lrb_6
					gen10877 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen10878 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen10879 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen10880 := Call(__e, gen10879)

					gen10881 := Call(__e, gen10878, gen10880, Parse__shen_4_5lrb_6)

					gen10882 := Call(__e, gen10877, gen10881)

					if True == gen10882 {
						gen10874 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5st__input1_6 := __e.Get(1)
							_ = Parse__shen_4_5st__input1_6
							gen10868 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen10869 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen10870 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen10871 := Call(__e, gen10870)

							gen10872 := Call(__e, gen10869, gen10871, Parse__shen_4_5st__input1_6)

							gen10873 := Call(__e, gen10868, gen10872)

							if True == gen10873 {
								gen10865 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5rrb_6 := __e.Get(1)
									_ = Parse__shen_4_5rrb_6
									gen10859 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen10860 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen10861 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen10862 := Call(__e, gen10861)

									gen10863 := Call(__e, gen10860, gen10862, Parse__shen_4_5rrb_6)

									gen10864 := Call(__e, gen10859, gen10863)

									if True == gen10864 {
										gen10856 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5st__input2_6 := __e.Get(1)
											_ = Parse__shen_4_5st__input2_6
											gen10850 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen10851 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen10852 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen10853 := Call(__e, gen10852)

											gen10854 := Call(__e, gen10851, gen10853, Parse__shen_4_5st__input2_6)

											gen10855 := Call(__e, gen10850, gen10854)

											if True == gen10855 {
												gen10839 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												gen10840 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen10841 := Call(__e, gen10840, Parse__shen_4_5st__input2_6)

												gen10842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4package_1macro)

												gen10843 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

												gen10844 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												gen10845 := Call(__e, gen10844, Parse__shen_4_5st__input1_6)

												gen10846 := Call(__e, gen10843, gen10845)

												gen10847 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												gen10848 := Call(__e, gen10847, Parse__shen_4_5st__input2_6)

												gen10849 := Call(__e, gen10842, gen10846, gen10848)

												__e.TailApply(gen10839, gen10841, gen10849)

												return

											} else {
												gen10838 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen10838)

												return

											}

										}, 1)

										gen10857 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input2_6)

										gen10858 := Call(__e, gen10857, Parse__shen_4_5rrb_6)

										__e.TailApply(gen10856, gen10858)

										return

									} else {
										gen10837 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen10837)

										return

									}

								}, 1)

								gen10866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rrb_6)

								gen10867 := Call(__e, gen10866, Parse__shen_4_5st__input1_6)

								__e.TailApply(gen10865, gen10867)

								return

							} else {
								gen10836 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen10836)

								return

							}

						}, 1)

						gen10875 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input1_6)

						gen10876 := Call(__e, gen10875, Parse__shen_4_5lrb_6)

						__e.TailApply(gen10874, gen10876)

						return

					} else {
						gen10835 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen10835)

						return

					}

				}, 1)

				gen10884 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5lrb_6)

				gen10885 := Call(__e, gen10884, V1515)

				gen10886 := Call(__e, gen10883, gen10885)

				__e.TailApply(gen10834, gen10886)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen10942 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5lsb_6 := __e.Get(1)
			_ = Parse__shen_4_5lsb_6
			gen10936 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen10937 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10938 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen10939 := Call(__e, gen10938)

			gen10940 := Call(__e, gen10937, gen10939, Parse__shen_4_5lsb_6)

			gen10941 := Call(__e, gen10936, gen10940)

			if True == gen10941 {
				gen10933 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5st__input1_6 := __e.Get(1)
					_ = Parse__shen_4_5st__input1_6
					gen10927 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen10928 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen10929 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen10930 := Call(__e, gen10929)

					gen10931 := Call(__e, gen10928, gen10930, Parse__shen_4_5st__input1_6)

					gen10932 := Call(__e, gen10927, gen10931)

					if True == gen10932 {
						gen10924 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rsb_6 := __e.Get(1)
							_ = Parse__shen_4_5rsb_6
							gen10918 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen10919 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen10920 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen10921 := Call(__e, gen10920)

							gen10922 := Call(__e, gen10919, gen10921, Parse__shen_4_5rsb_6)

							gen10923 := Call(__e, gen10918, gen10922)

							if True == gen10923 {
								gen10915 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5st__input2_6 := __e.Get(1)
									_ = Parse__shen_4_5st__input2_6
									gen10909 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen10910 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen10911 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen10912 := Call(__e, gen10911)

									gen10913 := Call(__e, gen10910, gen10912, Parse__shen_4_5st__input2_6)

									gen10914 := Call(__e, gen10909, gen10913)

									if True == gen10914 {
										gen10896 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen10897 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen10898 := Call(__e, gen10897, Parse__shen_4_5st__input2_6)

										gen10899 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen10900 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

										gen10901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cons__form)

										gen10902 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen10903 := Call(__e, gen10902, Parse__shen_4_5st__input1_6)

										gen10904 := Call(__e, gen10901, gen10903)

										gen10905 := Call(__e, gen10900, gen10904)

										gen10906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen10907 := Call(__e, gen10906, Parse__shen_4_5st__input2_6)

										gen10908 := Call(__e, gen10899, gen10905, gen10907)

										__e.TailApply(gen10896, gen10898, gen10908)

										return

									} else {
										gen10895 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen10895)

										return

									}

								}, 1)

								gen10916 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input2_6)

								gen10917 := Call(__e, gen10916, Parse__shen_4_5rsb_6)

								__e.TailApply(gen10915, gen10917)

								return

							} else {
								gen10894 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen10894)

								return

							}

						}, 1)

						gen10925 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rsb_6)

						gen10926 := Call(__e, gen10925, Parse__shen_4_5st__input1_6)

						__e.TailApply(gen10924, gen10926)

						return

					} else {
						gen10893 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen10893)

						return

					}

				}, 1)

				gen10934 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input1_6)

				gen10935 := Call(__e, gen10934, Parse__shen_4_5lsb_6)

				__e.TailApply(gen10933, gen10935)

				return

			} else {
				gen10892 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen10892)

				return

			}

		}, 1)

		gen10943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5lsb_6)

		gen10944 := Call(__e, gen10943, V1515)

		gen10945 := Call(__e, gen10942, gen10944)

		__e.TailApply(gen10891, gen10945)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input_6, gen10946)

	gen10967 := MakeNative(func(__e *ControlFlow) {
		V1518 := __e.Get(1)
		_ = V1518
		gen10962 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10963 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10964 := Call(__e, gen10963, V1518)

		gen10965 := Call(__e, gen10962, gen10964)

		var gen10966 Obj

		if True == gen10965 {
			gen10958 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10959 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen10960 := Call(__e, gen10959, V1518)

			gen10961 := Call(__e, gen10958, MakeNumber(91), gen10960)

			if True == gen10961 {
				gen10966 = True
			} else {
				gen10966 = False
			}

		} else {
			gen10966 = False
		}

		if True == gen10966 {
			gen10951 := MakeNative(func(__e *ControlFlow) {
				NewStream1516 := __e.Get(1)
				_ = NewStream1516
				gen10948 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen10949 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen10950 := Call(__e, gen10949, NewStream1516)

				__e.TailApply(gen10948, gen10950, symshen_4skip)

				return

			}, 1)

			gen10952 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen10953 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen10954 := Call(__e, gen10953, V1518)

			gen10955 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen10956 := Call(__e, gen10955, V1518)

			gen10957 := Call(__e, gen10952, gen10954, gen10956)

			__e.TailApply(gen10951, gen10957)

			return

		} else {
			gen10947 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen10947)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lsb_6, gen10967)

	gen10988 := MakeNative(func(__e *ControlFlow) {
		V1521 := __e.Get(1)
		_ = V1521
		gen10983 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10984 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10985 := Call(__e, gen10984, V1521)

		gen10986 := Call(__e, gen10983, gen10985)

		var gen10987 Obj

		if True == gen10986 {
			gen10979 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen10980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen10981 := Call(__e, gen10980, V1521)

			gen10982 := Call(__e, gen10979, MakeNumber(93), gen10981)

			if True == gen10982 {
				gen10987 = True
			} else {
				gen10987 = False
			}

		} else {
			gen10987 = False
		}

		if True == gen10987 {
			gen10972 := MakeNative(func(__e *ControlFlow) {
				NewStream1519 := __e.Get(1)
				_ = NewStream1519
				gen10969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen10970 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen10971 := Call(__e, gen10970, NewStream1519)

				__e.TailApply(gen10969, gen10971, symshen_4skip)

				return

			}, 1)

			gen10973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen10974 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen10975 := Call(__e, gen10974, V1521)

			gen10976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen10977 := Call(__e, gen10976, V1521)

			gen10978 := Call(__e, gen10973, gen10975, gen10977)

			__e.TailApply(gen10972, gen10978)

			return

		} else {
			gen10968 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen10968)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rsb_6, gen10988)

	gen11009 := MakeNative(func(__e *ControlFlow) {
		V1524 := __e.Get(1)
		_ = V1524
		gen11004 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11005 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11006 := Call(__e, gen11005, V1524)

		gen11007 := Call(__e, gen11004, gen11006)

		var gen11008 Obj

		if True == gen11007 {
			gen11000 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11001 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11002 := Call(__e, gen11001, V1524)

			gen11003 := Call(__e, gen11000, MakeNumber(123), gen11002)

			if True == gen11003 {
				gen11008 = True
			} else {
				gen11008 = False
			}

		} else {
			gen11008 = False
		}

		if True == gen11008 {
			gen10993 := MakeNative(func(__e *ControlFlow) {
				NewStream1522 := __e.Get(1)
				_ = NewStream1522
				gen10990 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen10991 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen10992 := Call(__e, gen10991, NewStream1522)

				__e.TailApply(gen10990, gen10992, symshen_4skip)

				return

			}, 1)

			gen10994 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen10995 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen10996 := Call(__e, gen10995, V1524)

			gen10997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen10998 := Call(__e, gen10997, V1524)

			gen10999 := Call(__e, gen10994, gen10996, gen10998)

			__e.TailApply(gen10993, gen10999)

			return

		} else {
			gen10989 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen10989)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lcurly_6, gen11009)

	gen11030 := MakeNative(func(__e *ControlFlow) {
		V1527 := __e.Get(1)
		_ = V1527
		gen11025 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11026 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11027 := Call(__e, gen11026, V1527)

		gen11028 := Call(__e, gen11025, gen11027)

		var gen11029 Obj

		if True == gen11028 {
			gen11021 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11023 := Call(__e, gen11022, V1527)

			gen11024 := Call(__e, gen11021, MakeNumber(125), gen11023)

			if True == gen11024 {
				gen11029 = True
			} else {
				gen11029 = False
			}

		} else {
			gen11029 = False
		}

		if True == gen11029 {
			gen11014 := MakeNative(func(__e *ControlFlow) {
				NewStream1525 := __e.Get(1)
				_ = NewStream1525
				gen11011 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11012 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11013 := Call(__e, gen11012, NewStream1525)

				__e.TailApply(gen11011, gen11013, symshen_4skip)

				return

			}, 1)

			gen11015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11016 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11017 := Call(__e, gen11016, V1527)

			gen11018 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11019 := Call(__e, gen11018, V1527)

			gen11020 := Call(__e, gen11015, gen11017, gen11019)

			__e.TailApply(gen11014, gen11020)

			return

		} else {
			gen11010 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11010)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rcurly_6, gen11030)

	gen11051 := MakeNative(func(__e *ControlFlow) {
		V1530 := __e.Get(1)
		_ = V1530
		gen11046 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11047 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11048 := Call(__e, gen11047, V1530)

		gen11049 := Call(__e, gen11046, gen11048)

		var gen11050 Obj

		if True == gen11049 {
			gen11042 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11043 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11044 := Call(__e, gen11043, V1530)

			gen11045 := Call(__e, gen11042, MakeNumber(124), gen11044)

			if True == gen11045 {
				gen11050 = True
			} else {
				gen11050 = False
			}

		} else {
			gen11050 = False
		}

		if True == gen11050 {
			gen11035 := MakeNative(func(__e *ControlFlow) {
				NewStream1528 := __e.Get(1)
				_ = NewStream1528
				gen11032 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11033 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11034 := Call(__e, gen11033, NewStream1528)

				__e.TailApply(gen11032, gen11034, symshen_4skip)

				return

			}, 1)

			gen11036 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11038 := Call(__e, gen11037, V1530)

			gen11039 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11040 := Call(__e, gen11039, V1530)

			gen11041 := Call(__e, gen11036, gen11038, gen11040)

			__e.TailApply(gen11035, gen11041)

			return

		} else {
			gen11031 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11031)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5bar_6, gen11051)

	gen11072 := MakeNative(func(__e *ControlFlow) {
		V1533 := __e.Get(1)
		_ = V1533
		gen11067 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11068 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11069 := Call(__e, gen11068, V1533)

		gen11070 := Call(__e, gen11067, gen11069)

		var gen11071 Obj

		if True == gen11070 {
			gen11063 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11064 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11065 := Call(__e, gen11064, V1533)

			gen11066 := Call(__e, gen11063, MakeNumber(59), gen11065)

			if True == gen11066 {
				gen11071 = True
			} else {
				gen11071 = False
			}

		} else {
			gen11071 = False
		}

		if True == gen11071 {
			gen11056 := MakeNative(func(__e *ControlFlow) {
				NewStream1531 := __e.Get(1)
				_ = NewStream1531
				gen11053 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11054 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11055 := Call(__e, gen11054, NewStream1531)

				__e.TailApply(gen11053, gen11055, symshen_4skip)

				return

			}, 1)

			gen11057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11058 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11059 := Call(__e, gen11058, V1533)

			gen11060 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11061 := Call(__e, gen11060, V1533)

			gen11062 := Call(__e, gen11057, gen11059, gen11061)

			__e.TailApply(gen11056, gen11062)

			return

		} else {
			gen11052 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11052)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5semicolon_6, gen11072)

	gen11093 := MakeNative(func(__e *ControlFlow) {
		V1536 := __e.Get(1)
		_ = V1536
		gen11088 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11089 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11090 := Call(__e, gen11089, V1536)

		gen11091 := Call(__e, gen11088, gen11090)

		var gen11092 Obj

		if True == gen11091 {
			gen11084 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11085 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11086 := Call(__e, gen11085, V1536)

			gen11087 := Call(__e, gen11084, MakeNumber(58), gen11086)

			if True == gen11087 {
				gen11092 = True
			} else {
				gen11092 = False
			}

		} else {
			gen11092 = False
		}

		if True == gen11092 {
			gen11077 := MakeNative(func(__e *ControlFlow) {
				NewStream1534 := __e.Get(1)
				_ = NewStream1534
				gen11074 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11075 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11076 := Call(__e, gen11075, NewStream1534)

				__e.TailApply(gen11074, gen11076, symshen_4skip)

				return

			}, 1)

			gen11078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11079 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11080 := Call(__e, gen11079, V1536)

			gen11081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11082 := Call(__e, gen11081, V1536)

			gen11083 := Call(__e, gen11078, gen11080, gen11082)

			__e.TailApply(gen11077, gen11083)

			return

		} else {
			gen11073 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11073)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5colon_6, gen11093)

	gen11114 := MakeNative(func(__e *ControlFlow) {
		V1539 := __e.Get(1)
		_ = V1539
		gen11109 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11110 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11111 := Call(__e, gen11110, V1539)

		gen11112 := Call(__e, gen11109, gen11111)

		var gen11113 Obj

		if True == gen11112 {
			gen11105 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11107 := Call(__e, gen11106, V1539)

			gen11108 := Call(__e, gen11105, MakeNumber(44), gen11107)

			if True == gen11108 {
				gen11113 = True
			} else {
				gen11113 = False
			}

		} else {
			gen11113 = False
		}

		if True == gen11113 {
			gen11098 := MakeNative(func(__e *ControlFlow) {
				NewStream1537 := __e.Get(1)
				_ = NewStream1537
				gen11095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11096 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11097 := Call(__e, gen11096, NewStream1537)

				__e.TailApply(gen11095, gen11097, symshen_4skip)

				return

			}, 1)

			gen11099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11101 := Call(__e, gen11100, V1539)

			gen11102 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11103 := Call(__e, gen11102, V1539)

			gen11104 := Call(__e, gen11099, gen11101, gen11103)

			__e.TailApply(gen11098, gen11104)

			return

		} else {
			gen11094 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11094)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comma_6, gen11114)

	gen11135 := MakeNative(func(__e *ControlFlow) {
		V1542 := __e.Get(1)
		_ = V1542
		gen11130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11131 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11132 := Call(__e, gen11131, V1542)

		gen11133 := Call(__e, gen11130, gen11132)

		var gen11134 Obj

		if True == gen11133 {
			gen11126 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11128 := Call(__e, gen11127, V1542)

			gen11129 := Call(__e, gen11126, MakeNumber(61), gen11128)

			if True == gen11129 {
				gen11134 = True
			} else {
				gen11134 = False
			}

		} else {
			gen11134 = False
		}

		if True == gen11134 {
			gen11119 := MakeNative(func(__e *ControlFlow) {
				NewStream1540 := __e.Get(1)
				_ = NewStream1540
				gen11116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11117 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11118 := Call(__e, gen11117, NewStream1540)

				__e.TailApply(gen11116, gen11118, symshen_4skip)

				return

			}, 1)

			gen11120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11121 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11122 := Call(__e, gen11121, V1542)

			gen11123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11124 := Call(__e, gen11123, V1542)

			gen11125 := Call(__e, gen11120, gen11122, gen11124)

			__e.TailApply(gen11119, gen11125)

			return

		} else {
			gen11115 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11115)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5equal_6, gen11135)

	gen11156 := MakeNative(func(__e *ControlFlow) {
		V1545 := __e.Get(1)
		_ = V1545
		gen11151 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11152 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11153 := Call(__e, gen11152, V1545)

		gen11154 := Call(__e, gen11151, gen11153)

		var gen11155 Obj

		if True == gen11154 {
			gen11147 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11148 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11149 := Call(__e, gen11148, V1545)

			gen11150 := Call(__e, gen11147, MakeNumber(45), gen11149)

			if True == gen11150 {
				gen11155 = True
			} else {
				gen11155 = False
			}

		} else {
			gen11155 = False
		}

		if True == gen11155 {
			gen11140 := MakeNative(func(__e *ControlFlow) {
				NewStream1543 := __e.Get(1)
				_ = NewStream1543
				gen11137 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11138 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11139 := Call(__e, gen11138, NewStream1543)

				__e.TailApply(gen11137, gen11139, symshen_4skip)

				return

			}, 1)

			gen11141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11142 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11143 := Call(__e, gen11142, V1545)

			gen11144 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11145 := Call(__e, gen11144, V1545)

			gen11146 := Call(__e, gen11141, gen11143, gen11145)

			__e.TailApply(gen11140, gen11146)

			return

		} else {
			gen11136 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11136)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5minus_6, gen11156)

	gen11177 := MakeNative(func(__e *ControlFlow) {
		V1548 := __e.Get(1)
		_ = V1548
		gen11172 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11173 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11174 := Call(__e, gen11173, V1548)

		gen11175 := Call(__e, gen11172, gen11174)

		var gen11176 Obj

		if True == gen11175 {
			gen11168 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11169 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11170 := Call(__e, gen11169, V1548)

			gen11171 := Call(__e, gen11168, MakeNumber(40), gen11170)

			if True == gen11171 {
				gen11176 = True
			} else {
				gen11176 = False
			}

		} else {
			gen11176 = False
		}

		if True == gen11176 {
			gen11161 := MakeNative(func(__e *ControlFlow) {
				NewStream1546 := __e.Get(1)
				_ = NewStream1546
				gen11158 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11159 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11160 := Call(__e, gen11159, NewStream1546)

				__e.TailApply(gen11158, gen11160, symshen_4skip)

				return

			}, 1)

			gen11162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11164 := Call(__e, gen11163, V1548)

			gen11165 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11166 := Call(__e, gen11165, V1548)

			gen11167 := Call(__e, gen11162, gen11164, gen11166)

			__e.TailApply(gen11161, gen11167)

			return

		} else {
			gen11157 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11157)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5lrb_6, gen11177)

	gen11198 := MakeNative(func(__e *ControlFlow) {
		V1551 := __e.Get(1)
		_ = V1551
		gen11193 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11194 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11195 := Call(__e, gen11194, V1551)

		gen11196 := Call(__e, gen11193, gen11195)

		var gen11197 Obj

		if True == gen11196 {
			gen11189 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11191 := Call(__e, gen11190, V1551)

			gen11192 := Call(__e, gen11189, MakeNumber(41), gen11191)

			if True == gen11192 {
				gen11197 = True
			} else {
				gen11197 = False
			}

		} else {
			gen11197 = False
		}

		if True == gen11197 {
			gen11182 := MakeNative(func(__e *ControlFlow) {
				NewStream1549 := __e.Get(1)
				_ = NewStream1549
				gen11179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11180 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11181 := Call(__e, gen11180, NewStream1549)

				__e.TailApply(gen11179, gen11181, symshen_4skip)

				return

			}, 1)

			gen11183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen11184 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen11185 := Call(__e, gen11184, V1551)

			gen11186 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen11187 := Call(__e, gen11186, V1551)

			gen11188 := Call(__e, gen11183, gen11185, gen11187)

			__e.TailApply(gen11182, gen11188)

			return

		} else {
			gen11178 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11178)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rrb_6, gen11198)

	gen11267 := MakeNative(func(__e *ControlFlow) {
		V1553 := __e.Get(1)
		_ = V1553
		gen11248 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen11244 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11245 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11246 := Call(__e, gen11245)

			gen11247 := Call(__e, gen11244, YaccParse, gen11246)

			if True == gen11247 {
				gen11227 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					gen11223 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11224 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11225 := Call(__e, gen11224)

					gen11226 := Call(__e, gen11223, YaccParse, gen11225)

					if True == gen11226 {
						gen11220 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5sym_6 := __e.Get(1)
							_ = Parse__shen_4_5sym_6
							gen11214 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen11215 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen11216 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen11217 := Call(__e, gen11216)

							gen11218 := Call(__e, gen11215, gen11217, Parse__shen_4_5sym_6)

							gen11219 := Call(__e, gen11214, gen11218)

							if True == gen11219 {
								gen11200 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen11201 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen11202 := Call(__e, gen11201, Parse__shen_4_5sym_6)

								gen11209 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen11210 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen11211 := Call(__e, gen11210, Parse__shen_4_5sym_6)

								gen11212 := Call(__e, gen11209, gen11211, MakeString("<>"))

								var gen11213 Obj

								if True == gen11212 {
									gen11206 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen11207 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen11208 := Call(__e, gen11207, MakeNumber(0), Nil)

									gen11213 = Call(__e, gen11206, symvector, gen11208)

								} else {
									gen11203 := Call(__e, PrimNS1Value(symns2_1value), symintern)

									gen11204 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									gen11205 := Call(__e, gen11204, Parse__shen_4_5sym_6)

									gen11213 = Call(__e, gen11203, gen11205)

								}

								__e.TailApply(gen11200, gen11202, gen11213)

								return

							} else {
								gen11199 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen11199)

								return

							}

						}, 1)

						gen11221 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5sym_6)

						gen11222 := Call(__e, gen11221, V1553)

						__e.TailApply(gen11220, gen11222)

						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				gen11240 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5number_6 := __e.Get(1)
					_ = Parse__shen_4_5number_6
					gen11234 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11235 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11236 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11237 := Call(__e, gen11236)

					gen11238 := Call(__e, gen11235, gen11237, Parse__shen_4_5number_6)

					gen11239 := Call(__e, gen11234, gen11238)

					if True == gen11239 {
						gen11229 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11230 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11231 := Call(__e, gen11230, Parse__shen_4_5number_6)

						gen11232 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11233 := Call(__e, gen11232, Parse__shen_4_5number_6)

						__e.TailApply(gen11229, gen11231, gen11233)

						return

					} else {
						gen11228 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11228)

						return

					}

				}, 1)

				gen11241 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5number_6)

				gen11242 := Call(__e, gen11241, V1553)

				gen11243 := Call(__e, gen11240, gen11242)

				__e.TailApply(gen11227, gen11243)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen11263 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5str_6 := __e.Get(1)
			_ = Parse__shen_4_5str_6
			gen11257 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen11258 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11259 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11260 := Call(__e, gen11259)

			gen11261 := Call(__e, gen11258, gen11260, Parse__shen_4_5str_6)

			gen11262 := Call(__e, gen11257, gen11261)

			if True == gen11262 {
				gen11250 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11251 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11252 := Call(__e, gen11251, Parse__shen_4_5str_6)

				gen11253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4control_1chars)

				gen11254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen11255 := Call(__e, gen11254, Parse__shen_4_5str_6)

				gen11256 := Call(__e, gen11253, gen11255)

				__e.TailApply(gen11250, gen11252, gen11256)

				return

			} else {
				gen11249 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen11249)

				return

			}

		}, 1)

		gen11264 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5str_6)

		gen11265 := Call(__e, gen11264, V1553)

		gen11266 := Call(__e, gen11263, gen11265)

		__e.TailApply(gen11248, gen11266)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5atom_6, gen11267)

	gen11321 := MakeNative(func(__e *ControlFlow) {
		V1555 := __e.Get(1)
		_ = V1555
		gen11319 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen11320 := Call(__e, gen11319, Nil, V1555)

		if True == gen11320 {
			__e.Return(MakeString(""))
			return
		} else {
			gen11316 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen11317 := Call(__e, gen11316, V1555)

			var gen11318 Obj

			if True == gen11317 {
				gen11311 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen11312 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11313 := Call(__e, gen11312, V1555)

				gen11314 := Call(__e, gen11311, MakeString("c"), gen11313)

				var gen11315 Obj

				if True == gen11314 {
					gen11306 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen11307 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen11308 := Call(__e, gen11307, V1555)

					gen11309 := Call(__e, gen11306, gen11308)

					var gen11310 Obj

					if True == gen11309 {
						gen11300 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen11301 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11302 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen11303 := Call(__e, gen11302, V1555)

						gen11304 := Call(__e, gen11301, gen11303)

						gen11305 := Call(__e, gen11300, MakeString("#"), gen11304)

						if True == gen11305 {
							gen11310 = True
						} else {
							gen11310 = False
						}

					} else {
						gen11310 = False
					}

					if True == gen11310 {
						gen11315 = True
					} else {
						gen11315 = False
					}

				} else {
					gen11315 = False
				}

				if True == gen11315 {
					gen11318 = True
				} else {
					gen11318 = False
				}

			} else {
				gen11318 = False
			}

			if True == gen11318 {
				gen11293 := MakeNative(func(__e *ControlFlow) {
					CodePoint := __e.Get(1)
					_ = CodePoint
					gen11286 := MakeNative(func(__e *ControlFlow) {
						AfterCodePoint := __e.Get(1)
						_ = AfterCodePoint
						gen11279 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						gen11280 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

						gen11281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decimalise)

						gen11282 := Call(__e, gen11281, CodePoint)

						gen11283 := Call(__e, gen11280, gen11282)

						gen11284 := Call(__e, PrimNS1Value(symns2_1value), symshen_4control_1chars)

						gen11285 := Call(__e, gen11284, AfterCodePoint)

						__e.TailApply(gen11279, gen11283, gen11285)

						return

					}, 1)

					gen11287 := Call(__e, PrimNS1Value(symns2_1value), symshen_4after_1codepoint)

					gen11288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen11289 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen11290 := Call(__e, gen11289, V1555)

					gen11291 := Call(__e, gen11288, gen11290)

					gen11292 := Call(__e, gen11287, gen11291)

					__e.TailApply(gen11286, gen11292)

					return

				}, 1)

				gen11294 := Call(__e, PrimNS1Value(symns2_1value), symshen_4code_1point)

				gen11295 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen11296 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen11297 := Call(__e, gen11296, V1555)

				gen11298 := Call(__e, gen11295, gen11297)

				gen11299 := Call(__e, gen11294, gen11298)

				__e.TailApply(gen11293, gen11299)

				return

			} else {
				gen11277 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen11278 := Call(__e, gen11277, V1555)

				if True == gen11278 {
					gen11270 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

					gen11271 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen11272 := Call(__e, gen11271, V1555)

					gen11273 := Call(__e, PrimNS1Value(symns2_1value), symshen_4control_1chars)

					gen11274 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen11275 := Call(__e, gen11274, V1555)

					gen11276 := Call(__e, gen11273, gen11275)

					__e.TailApply(gen11270, gen11272, gen11276)

					return

				} else {
					if True == True {
						gen11269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen11269, symshen_4control_1chars)

						return

					} else {
						gen11268 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen11268, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4control_1chars, gen11321)

	gen11371 := MakeNative(func(__e *ControlFlow) {
		V1559 := __e.Get(1)
		_ = V1559
		gen11368 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11369 := Call(__e, gen11368, V1559)

		var gen11370 Obj

		if True == gen11369 {
			gen11364 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen11366 := Call(__e, gen11365, V1559)

			gen11367 := Call(__e, gen11364, MakeString(";"), gen11366)

			if True == gen11367 {
				gen11370 = True
			} else {
				gen11370 = False
			}

		} else {
			gen11370 = False
		}

		if True == gen11370 {
			__e.Return(MakeString(""))
			return
		} else {
			gen11361 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen11362 := Call(__e, gen11361, V1559)

			var gen11363 Obj

			if True == gen11362 {
				gen11335 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen11336 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11337 := Call(__e, gen11336, V1559)

				gen11338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11339 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11340 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11344 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11345 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11349 := Call(__e, gen11348, MakeString("0"), Nil)

				gen11350 := Call(__e, gen11347, MakeString("9"), gen11349)

				gen11351 := Call(__e, gen11346, MakeString("8"), gen11350)

				gen11352 := Call(__e, gen11345, MakeString("7"), gen11351)

				gen11353 := Call(__e, gen11344, MakeString("6"), gen11352)

				gen11354 := Call(__e, gen11343, MakeString("5"), gen11353)

				gen11355 := Call(__e, gen11342, MakeString("4"), gen11354)

				gen11356 := Call(__e, gen11341, MakeString("3"), gen11355)

				gen11357 := Call(__e, gen11340, MakeString("2"), gen11356)

				gen11358 := Call(__e, gen11339, MakeString("1"), gen11357)

				gen11359 := Call(__e, gen11338, MakeString("0"), gen11358)

				gen11360 := Call(__e, gen11335, gen11337, gen11359)

				if True == gen11360 {
					gen11363 = True
				} else {
					gen11363 = False
				}

			} else {
				gen11363 = False
			}

			if True == gen11363 {
				gen11328 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11329 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11330 := Call(__e, gen11329, V1559)

				gen11331 := Call(__e, PrimNS1Value(symns2_1value), symshen_4code_1point)

				gen11332 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen11333 := Call(__e, gen11332, V1559)

				gen11334 := Call(__e, gen11331, gen11333)

				__e.TailApply(gen11328, gen11330, gen11334)

				return

			} else {
				if True == True {
					gen11323 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					gen11324 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen11325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen11326 := Call(__e, gen11325, V1559, MakeString("\n"), symshen_4a)

					gen11327 := Call(__e, gen11324, MakeString("code point parse error "), gen11326)

					__e.TailApply(gen11323, gen11327)

					return

				} else {
					gen11322 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen11322, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4code_1point, gen11371)

	gen11389 := MakeNative(func(__e *ControlFlow) {
		V1565 := __e.Get(1)
		_ = V1565
		gen11387 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen11388 := Call(__e, gen11387, Nil, V1565)

		if True == gen11388 {
			__e.Return(Nil)
			return
		} else {
			gen11384 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen11385 := Call(__e, gen11384, V1565)

			var gen11386 Obj

			if True == gen11385 {
				gen11380 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen11381 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11382 := Call(__e, gen11381, V1565)

				gen11383 := Call(__e, gen11380, MakeString(";"), gen11382)

				if True == gen11383 {
					gen11386 = True
				} else {
					gen11386 = False
				}

			} else {
				gen11386 = False
			}

			if True == gen11386 {
				gen11379 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				__e.TailApply(gen11379, V1565)

				return

			} else {
				gen11377 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen11378 := Call(__e, gen11377, V1565)

				if True == gen11378 {
					gen11374 := Call(__e, PrimNS1Value(symns2_1value), symshen_4after_1codepoint)

					gen11375 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen11376 := Call(__e, gen11375, V1565)

					__e.TailApply(gen11374, gen11376)

					return

				} else {
					if True == True {
						gen11373 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen11373, symshen_4after_1codepoint)

						return

					} else {
						gen11372 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen11372, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4after_1codepoint, gen11389)

	gen11395 := MakeNative(func(__e *ControlFlow) {
		V1567 := __e.Get(1)
		_ = V1567
		gen11390 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

		gen11391 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

		gen11392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

		gen11393 := Call(__e, gen11392, V1567)

		gen11394 := Call(__e, gen11391, gen11393)

		__e.TailApply(gen11390, gen11394, MakeNumber(0))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4decimalise, gen11395)

	gen11517 := MakeNative(func(__e *ControlFlow) {
		V1573 := __e.Get(1)
		_ = V1573
		gen11514 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11515 := Call(__e, gen11514, V1573)

		var gen11516 Obj

		if True == gen11515 {
			gen11510 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11511 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen11512 := Call(__e, gen11511, V1573)

			gen11513 := Call(__e, gen11510, MakeString("0"), gen11512)

			if True == gen11513 {
				gen11516 = True
			} else {
				gen11516 = False
			}

		} else {
			gen11516 = False
		}

		if True == gen11516 {
			gen11505 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen11506 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

			gen11507 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen11508 := Call(__e, gen11507, V1573)

			gen11509 := Call(__e, gen11506, gen11508)

			__e.TailApply(gen11505, MakeNumber(0), gen11509)

			return

		} else {
			gen11502 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen11503 := Call(__e, gen11502, V1573)

			var gen11504 Obj

			if True == gen11503 {
				gen11498 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen11499 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11500 := Call(__e, gen11499, V1573)

				gen11501 := Call(__e, gen11498, MakeString("1"), gen11500)

				if True == gen11501 {
					gen11504 = True
				} else {
					gen11504 = False
				}

			} else {
				gen11504 = False
			}

			if True == gen11504 {
				gen11493 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen11494 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

				gen11495 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen11496 := Call(__e, gen11495, V1573)

				gen11497 := Call(__e, gen11494, gen11496)

				__e.TailApply(gen11493, MakeNumber(1), gen11497)

				return

			} else {
				gen11490 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen11491 := Call(__e, gen11490, V1573)

				var gen11492 Obj

				if True == gen11491 {
					gen11486 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11487 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen11488 := Call(__e, gen11487, V1573)

					gen11489 := Call(__e, gen11486, MakeString("2"), gen11488)

					if True == gen11489 {
						gen11492 = True
					} else {
						gen11492 = False
					}

				} else {
					gen11492 = False
				}

				if True == gen11492 {
					gen11481 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen11482 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

					gen11483 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen11484 := Call(__e, gen11483, V1573)

					gen11485 := Call(__e, gen11482, gen11484)

					__e.TailApply(gen11481, MakeNumber(2), gen11485)

					return

				} else {
					gen11478 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen11479 := Call(__e, gen11478, V1573)

					var gen11480 Obj

					if True == gen11479 {
						gen11474 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen11475 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11476 := Call(__e, gen11475, V1573)

						gen11477 := Call(__e, gen11474, MakeString("3"), gen11476)

						if True == gen11477 {
							gen11480 = True
						} else {
							gen11480 = False
						}

					} else {
						gen11480 = False
					}

					if True == gen11480 {
						gen11469 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen11470 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

						gen11471 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen11472 := Call(__e, gen11471, V1573)

						gen11473 := Call(__e, gen11470, gen11472)

						__e.TailApply(gen11469, MakeNumber(3), gen11473)

						return

					} else {
						gen11466 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen11467 := Call(__e, gen11466, V1573)

						var gen11468 Obj

						if True == gen11467 {
							gen11462 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen11463 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen11464 := Call(__e, gen11463, V1573)

							gen11465 := Call(__e, gen11462, MakeString("4"), gen11464)

							if True == gen11465 {
								gen11468 = True
							} else {
								gen11468 = False
							}

						} else {
							gen11468 = False
						}

						if True == gen11468 {
							gen11457 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen11458 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

							gen11459 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen11460 := Call(__e, gen11459, V1573)

							gen11461 := Call(__e, gen11458, gen11460)

							__e.TailApply(gen11457, MakeNumber(4), gen11461)

							return

						} else {
							gen11454 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen11455 := Call(__e, gen11454, V1573)

							var gen11456 Obj

							if True == gen11455 {
								gen11450 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen11451 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen11452 := Call(__e, gen11451, V1573)

								gen11453 := Call(__e, gen11450, MakeString("5"), gen11452)

								if True == gen11453 {
									gen11456 = True
								} else {
									gen11456 = False
								}

							} else {
								gen11456 = False
							}

							if True == gen11456 {
								gen11445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen11446 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

								gen11447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen11448 := Call(__e, gen11447, V1573)

								gen11449 := Call(__e, gen11446, gen11448)

								__e.TailApply(gen11445, MakeNumber(5), gen11449)

								return

							} else {
								gen11442 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen11443 := Call(__e, gen11442, V1573)

								var gen11444 Obj

								if True == gen11443 {
									gen11438 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen11439 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen11440 := Call(__e, gen11439, V1573)

									gen11441 := Call(__e, gen11438, MakeString("6"), gen11440)

									if True == gen11441 {
										gen11444 = True
									} else {
										gen11444 = False
									}

								} else {
									gen11444 = False
								}

								if True == gen11444 {
									gen11433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen11434 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

									gen11435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen11436 := Call(__e, gen11435, V1573)

									gen11437 := Call(__e, gen11434, gen11436)

									__e.TailApply(gen11433, MakeNumber(6), gen11437)

									return

								} else {
									gen11430 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen11431 := Call(__e, gen11430, V1573)

									var gen11432 Obj

									if True == gen11431 {
										gen11426 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen11427 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen11428 := Call(__e, gen11427, V1573)

										gen11429 := Call(__e, gen11426, MakeString("7"), gen11428)

										if True == gen11429 {
											gen11432 = True
										} else {
											gen11432 = False
										}

									} else {
										gen11432 = False
									}

									if True == gen11432 {
										gen11421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen11422 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

										gen11423 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen11424 := Call(__e, gen11423, V1573)

										gen11425 := Call(__e, gen11422, gen11424)

										__e.TailApply(gen11421, MakeNumber(7), gen11425)

										return

									} else {
										gen11418 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen11419 := Call(__e, gen11418, V1573)

										var gen11420 Obj

										if True == gen11419 {
											gen11414 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen11415 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen11416 := Call(__e, gen11415, V1573)

											gen11417 := Call(__e, gen11414, MakeString("8"), gen11416)

											if True == gen11417 {
												gen11420 = True
											} else {
												gen11420 = False
											}

										} else {
											gen11420 = False
										}

										if True == gen11420 {
											gen11409 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											gen11410 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

											gen11411 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen11412 := Call(__e, gen11411, V1573)

											gen11413 := Call(__e, gen11410, gen11412)

											__e.TailApply(gen11409, MakeNumber(8), gen11413)

											return

										} else {
											gen11406 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen11407 := Call(__e, gen11406, V1573)

											var gen11408 Obj

											if True == gen11407 {
												gen11402 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen11403 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen11404 := Call(__e, gen11403, V1573)

												gen11405 := Call(__e, gen11402, MakeString("9"), gen11404)

												if True == gen11405 {
													gen11408 = True
												} else {
													gen11408 = False
												}

											} else {
												gen11408 = False
											}

											if True == gen11408 {
												gen11397 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												gen11398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digits_1_6integers)

												gen11399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen11400 := Call(__e, gen11399, V1573)

												gen11401 := Call(__e, gen11398, gen11400)

												__e.TailApply(gen11397, MakeNumber(9), gen11401)

												return

											} else {
												if True == True {
													__e.Return(Nil)
													return
												} else {
													gen11396 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

													__e.TailApply(gen11396, MakeString("no cond match"))

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

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4digits_1_6integers, gen11517)

	gen11547 := MakeNative(func(__e *ControlFlow) {
		V1575 := __e.Get(1)
		_ = V1575
		gen11544 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alpha_6 := __e.Get(1)
			_ = Parse__shen_4_5alpha_6
			gen11538 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen11539 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11540 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11541 := Call(__e, gen11540)

			gen11542 := Call(__e, gen11539, gen11541, Parse__shen_4_5alpha_6)

			gen11543 := Call(__e, gen11538, gen11542)

			if True == gen11543 {
				gen11535 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5alphanums_6 := __e.Get(1)
					_ = Parse__shen_4_5alphanums_6
					gen11529 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11530 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11531 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11532 := Call(__e, gen11531)

					gen11533 := Call(__e, gen11530, gen11532, Parse__shen_4_5alphanums_6)

					gen11534 := Call(__e, gen11529, gen11533)

					if True == gen11534 {
						gen11520 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11521 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11522 := Call(__e, gen11521, Parse__shen_4_5alphanums_6)

						gen11523 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						gen11524 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11525 := Call(__e, gen11524, Parse__shen_4_5alpha_6)

						gen11526 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11527 := Call(__e, gen11526, Parse__shen_4_5alphanums_6)

						gen11528 := Call(__e, gen11523, gen11525, gen11527)

						__e.TailApply(gen11520, gen11522, gen11528)

						return

					} else {
						gen11519 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11519)

						return

					}

				}, 1)

				gen11536 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alphanums_6)

				gen11537 := Call(__e, gen11536, Parse__shen_4_5alpha_6)

				__e.TailApply(gen11535, gen11537)

				return

			} else {
				gen11518 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen11518)

				return

			}

		}, 1)

		gen11545 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alpha_6)

		gen11546 := Call(__e, gen11545, V1575)

		__e.TailApply(gen11544, gen11546)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sym_6, gen11547)

	gen11596 := MakeNative(func(__e *ControlFlow) {
		V1577 := __e.Get(1)
		_ = V1577
		gen11565 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen11561 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11562 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11563 := Call(__e, gen11562)

			gen11564 := Call(__e, gen11561, YaccParse, gen11563)

			if True == gen11564 {
				gen11558 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen11552 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11553 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11554 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11555 := Call(__e, gen11554)

					gen11556 := Call(__e, gen11553, gen11555, Parse___5e_6)

					gen11557 := Call(__e, gen11552, gen11556)

					if True == gen11557 {
						gen11549 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11550 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11551 := Call(__e, gen11550, Parse___5e_6)

						__e.TailApply(gen11549, gen11551, MakeString(""))

						return

					} else {
						gen11548 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11548)

						return

					}

				}, 1)

				gen11559 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen11560 := Call(__e, gen11559, V1577)

				__e.TailApply(gen11558, gen11560)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen11592 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alphanum_6 := __e.Get(1)
			_ = Parse__shen_4_5alphanum_6
			gen11586 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen11587 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11588 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11589 := Call(__e, gen11588)

			gen11590 := Call(__e, gen11587, gen11589, Parse__shen_4_5alphanum_6)

			gen11591 := Call(__e, gen11586, gen11590)

			if True == gen11591 {
				gen11583 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5alphanums_6 := __e.Get(1)
					_ = Parse__shen_4_5alphanums_6
					gen11577 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11578 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11579 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11580 := Call(__e, gen11579)

					gen11581 := Call(__e, gen11578, gen11580, Parse__shen_4_5alphanums_6)

					gen11582 := Call(__e, gen11577, gen11581)

					if True == gen11582 {
						gen11568 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11569 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11570 := Call(__e, gen11569, Parse__shen_4_5alphanums_6)

						gen11571 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						gen11572 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11573 := Call(__e, gen11572, Parse__shen_4_5alphanum_6)

						gen11574 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11575 := Call(__e, gen11574, Parse__shen_4_5alphanums_6)

						gen11576 := Call(__e, gen11571, gen11573, gen11575)

						__e.TailApply(gen11568, gen11570, gen11576)

						return

					} else {
						gen11567 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11567)

						return

					}

				}, 1)

				gen11584 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alphanums_6)

				gen11585 := Call(__e, gen11584, Parse__shen_4_5alphanum_6)

				__e.TailApply(gen11583, gen11585)

				return

			} else {
				gen11566 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen11566)

				return

			}

		}, 1)

		gen11593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alphanum_6)

		gen11594 := Call(__e, gen11593, V1577)

		gen11595 := Call(__e, gen11592, gen11594)

		__e.TailApply(gen11565, gen11595)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alphanums_6, gen11596)

	gen11633 := MakeNative(func(__e *ControlFlow) {
		V1579 := __e.Get(1)
		_ = V1579
		gen11616 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen11612 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11613 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11614 := Call(__e, gen11613)

			gen11615 := Call(__e, gen11612, YaccParse, gen11614)

			if True == gen11615 {
				gen11609 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5num_6 := __e.Get(1)
					_ = Parse__shen_4_5num_6
					gen11603 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11604 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11605 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11606 := Call(__e, gen11605)

					gen11607 := Call(__e, gen11604, gen11606, Parse__shen_4_5num_6)

					gen11608 := Call(__e, gen11603, gen11607)

					if True == gen11608 {
						gen11598 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11599 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11600 := Call(__e, gen11599, Parse__shen_4_5num_6)

						gen11601 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11602 := Call(__e, gen11601, Parse__shen_4_5num_6)

						__e.TailApply(gen11598, gen11600, gen11602)

						return

					} else {
						gen11597 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11597)

						return

					}

				}, 1)

				gen11610 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5num_6)

				gen11611 := Call(__e, gen11610, V1579)

				__e.TailApply(gen11609, gen11611)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen11629 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5alpha_6 := __e.Get(1)
			_ = Parse__shen_4_5alpha_6
			gen11623 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen11624 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11625 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11626 := Call(__e, gen11625)

			gen11627 := Call(__e, gen11624, gen11626, Parse__shen_4_5alpha_6)

			gen11628 := Call(__e, gen11623, gen11627)

			if True == gen11628 {
				gen11618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11619 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11620 := Call(__e, gen11619, Parse__shen_4_5alpha_6)

				gen11621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen11622 := Call(__e, gen11621, Parse__shen_4_5alpha_6)

				__e.TailApply(gen11618, gen11620, gen11622)

				return

			} else {
				gen11617 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen11617)

				return

			}

		}, 1)

		gen11630 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5alpha_6)

		gen11631 := Call(__e, gen11630, V1579)

		gen11632 := Call(__e, gen11629, gen11631)

		__e.TailApply(gen11616, gen11632)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alphanum_6, gen11633)

	gen11656 := MakeNative(func(__e *ControlFlow) {
		V1581 := __e.Get(1)
		_ = V1581
		gen11652 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11653 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11654 := Call(__e, gen11653, V1581)

		gen11655 := Call(__e, gen11652, gen11654)

		if True == gen11655 {
			gen11649 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen11647 := Call(__e, PrimNS1Value(symns2_1value), symshen_4numbyte_2)

				gen11648 := Call(__e, gen11647, Parse__Char)

				if True == gen11648 {
					gen11636 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11637 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen11638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen11640 := Call(__e, gen11639, V1581)

					gen11641 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen11642 := Call(__e, gen11641, V1581)

					gen11643 := Call(__e, gen11638, gen11640, gen11642)

					gen11644 := Call(__e, gen11637, gen11643)

					gen11645 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					gen11646 := Call(__e, gen11645, Parse__Char)

					__e.TailApply(gen11636, gen11644, gen11646)

					return

				} else {
					gen11635 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen11635)

					return

				}

			}, 1)

			gen11650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11651 := Call(__e, gen11650, V1581)

			__e.TailApply(gen11649, gen11651)

			return

		} else {
			gen11634 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11634)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5num_6, gen11656)

	gen11678 := MakeNative(func(__e *ControlFlow) {
		V1587 := __e.Get(1)
		_ = V1587
		gen11676 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen11677 := Call(__e, gen11676, MakeNumber(48), V1587)

		if True == gen11677 {
			__e.Return(True)
			return
		} else {
			gen11674 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11675 := Call(__e, gen11674, MakeNumber(49), V1587)

			if True == gen11675 {
				__e.Return(True)
				return
			} else {
				gen11672 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen11673 := Call(__e, gen11672, MakeNumber(50), V1587)

				if True == gen11673 {
					__e.Return(True)
					return
				} else {
					gen11670 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11671 := Call(__e, gen11670, MakeNumber(51), V1587)

					if True == gen11671 {
						__e.Return(True)
						return
					} else {
						gen11668 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen11669 := Call(__e, gen11668, MakeNumber(52), V1587)

						if True == gen11669 {
							__e.Return(True)
							return
						} else {
							gen11666 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen11667 := Call(__e, gen11666, MakeNumber(53), V1587)

							if True == gen11667 {
								__e.Return(True)
								return
							} else {
								gen11664 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen11665 := Call(__e, gen11664, MakeNumber(54), V1587)

								if True == gen11665 {
									__e.Return(True)
									return
								} else {
									gen11662 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen11663 := Call(__e, gen11662, MakeNumber(55), V1587)

									if True == gen11663 {
										__e.Return(True)
										return
									} else {
										gen11660 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen11661 := Call(__e, gen11660, MakeNumber(56), V1587)

										if True == gen11661 {
											__e.Return(True)
											return
										} else {
											gen11658 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen11659 := Call(__e, gen11658, MakeNumber(57), V1587)

											if True == gen11659 {
												__e.Return(True)
												return
											} else {
												if True == True {
													__e.Return(False)
													return
												} else {
													gen11657 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

													__e.TailApply(gen11657, MakeString("no cond match"))

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

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4numbyte_2, gen11678)

	gen11701 := MakeNative(func(__e *ControlFlow) {
		V1589 := __e.Get(1)
		_ = V1589
		gen11697 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11698 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11699 := Call(__e, gen11698, V1589)

		gen11700 := Call(__e, gen11697, gen11699)

		if True == gen11700 {
			gen11694 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen11692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4symbol_1code_2)

				gen11693 := Call(__e, gen11692, Parse__Char)

				if True == gen11693 {
					gen11681 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11682 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen11683 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11684 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen11685 := Call(__e, gen11684, V1589)

					gen11686 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen11687 := Call(__e, gen11686, V1589)

					gen11688 := Call(__e, gen11683, gen11685, gen11687)

					gen11689 := Call(__e, gen11682, gen11688)

					gen11690 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					gen11691 := Call(__e, gen11690, Parse__Char)

					__e.TailApply(gen11681, gen11689, gen11691)

					return

				} else {
					gen11680 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen11680)

					return

				}

			}, 1)

			gen11695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11696 := Call(__e, gen11695, V1589)

			__e.TailApply(gen11694, gen11696)

			return

		} else {
			gen11679 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11679)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5alpha_6, gen11701)

	gen11735 := MakeNative(func(__e *ControlFlow) {
		V1591 := __e.Get(1)
		_ = V1591
		gen11733 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen11734 := Call(__e, gen11733, V1591, MakeNumber(126))

		if True == gen11734 {
			__e.Return(True)
			return
		} else {
			gen11729 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			gen11730 := Call(__e, gen11729, V1591, MakeNumber(94))

			var gen11731 Obj

			if True == gen11730 {
				gen11727 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

				gen11728 := Call(__e, gen11727, V1591, MakeNumber(123))

				if True == gen11728 {
					gen11731 = True
				} else {
					gen11731 = False
				}

			} else {
				gen11731 = False
			}

			var gen11732 Obj

			if True == gen11731 {
				gen11732 = True
			} else {
				gen11723 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

				gen11724 := Call(__e, gen11723, V1591, MakeNumber(59))

				var gen11725 Obj

				if True == gen11724 {
					gen11721 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

					gen11722 := Call(__e, gen11721, V1591, MakeNumber(91))

					if True == gen11722 {
						gen11725 = True
					} else {
						gen11725 = False
					}

				} else {
					gen11725 = False
				}

				var gen11726 Obj

				if True == gen11725 {
					gen11726 = True
				} else {
					gen11717 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					gen11718 := Call(__e, gen11717, V1591, MakeNumber(41))

					var gen11719 Obj

					if True == gen11718 {
						gen11714 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

						gen11715 := Call(__e, gen11714, V1591, MakeNumber(58))

						var gen11716 Obj

						if True == gen11715 {
							gen11710 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen11711 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen11712 := Call(__e, gen11711, V1591, MakeNumber(44))

							gen11713 := Call(__e, gen11710, gen11712)

							if True == gen11713 {
								gen11716 = True
							} else {
								gen11716 = False
							}

						} else {
							gen11716 = False
						}

						if True == gen11716 {
							gen11719 = True
						} else {
							gen11719 = False
						}

					} else {
						gen11719 = False
					}

					var gen11720 Obj

					if True == gen11719 {
						gen11720 = True
					} else {
						gen11706 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

						gen11707 := Call(__e, gen11706, V1591, MakeNumber(34))

						var gen11708 Obj

						if True == gen11707 {
							gen11704 := Call(__e, PrimNS1Value(symns2_1value), sym_5)

							gen11705 := Call(__e, gen11704, V1591, MakeNumber(40))

							if True == gen11705 {
								gen11708 = True
							} else {
								gen11708 = False
							}

						} else {
							gen11708 = False
						}

						var gen11709 Obj

						if True == gen11708 {
							gen11709 = True
						} else {
							gen11702 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen11703 := Call(__e, gen11702, V1591, MakeNumber(33))

							if True == gen11703 {
								gen11709 = True
							} else {
								gen11709 = False
							}

						}

						if True == gen11709 {
							gen11720 = True
						} else {
							gen11720 = False
						}

					}

					if True == gen11720 {
						gen11726 = True
					} else {
						gen11726 = False
					}

				}

				if True == gen11726 {
					gen11732 = True
				} else {
					gen11732 = False
				}

			}

			if True == gen11732 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4symbol_1code_2, gen11735)

	gen11771 := MakeNative(func(__e *ControlFlow) {
		V1593 := __e.Get(1)
		_ = V1593
		gen11768 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5dbq_6 := __e.Get(1)
			_ = Parse__shen_4_5dbq_6
			gen11762 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen11763 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11764 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11765 := Call(__e, gen11764)

			gen11766 := Call(__e, gen11763, gen11765, Parse__shen_4_5dbq_6)

			gen11767 := Call(__e, gen11762, gen11766)

			if True == gen11767 {
				gen11759 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5strcontents_6 := __e.Get(1)
					_ = Parse__shen_4_5strcontents_6
					gen11753 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11754 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11755 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11756 := Call(__e, gen11755)

					gen11757 := Call(__e, gen11754, gen11756, Parse__shen_4_5strcontents_6)

					gen11758 := Call(__e, gen11753, gen11757)

					if True == gen11758 {
						gen11750 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5dbq_6 := __e.Get(1)
							_ = Parse__shen_4_5dbq_6
							gen11744 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen11745 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen11746 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen11747 := Call(__e, gen11746)

							gen11748 := Call(__e, gen11745, gen11747, Parse__shen_4_5dbq_6)

							gen11749 := Call(__e, gen11744, gen11748)

							if True == gen11749 {
								gen11739 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen11740 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen11741 := Call(__e, gen11740, Parse__shen_4_5dbq_6)

								gen11742 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen11743 := Call(__e, gen11742, Parse__shen_4_5strcontents_6)

								__e.TailApply(gen11739, gen11741, gen11743)

								return

							} else {
								gen11738 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen11738)

								return

							}

						}, 1)

						gen11751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5dbq_6)

						gen11752 := Call(__e, gen11751, Parse__shen_4_5strcontents_6)

						__e.TailApply(gen11750, gen11752)

						return

					} else {
						gen11737 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11737)

						return

					}

				}, 1)

				gen11760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5strcontents_6)

				gen11761 := Call(__e, gen11760, Parse__shen_4_5dbq_6)

				__e.TailApply(gen11759, gen11761)

				return

			} else {
				gen11736 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen11736)

				return

			}

		}, 1)

		gen11769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5dbq_6)

		gen11770 := Call(__e, gen11769, V1593)

		__e.TailApply(gen11768, gen11770)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5str_6, gen11771)

	gen11792 := MakeNative(func(__e *ControlFlow) {
		V1595 := __e.Get(1)
		_ = V1595
		gen11788 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11789 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11790 := Call(__e, gen11789, V1595)

		gen11791 := Call(__e, gen11788, gen11790)

		if True == gen11791 {
			gen11785 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen11783 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen11784 := Call(__e, gen11783, Parse__Char, MakeNumber(34))

				if True == gen11784 {
					gen11774 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11775 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen11776 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11777 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen11778 := Call(__e, gen11777, V1595)

					gen11779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen11780 := Call(__e, gen11779, V1595)

					gen11781 := Call(__e, gen11776, gen11778, gen11780)

					gen11782 := Call(__e, gen11775, gen11781)

					__e.TailApply(gen11774, gen11782, Parse__Char)

					return

				} else {
					gen11773 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen11773)

					return

				}

			}, 1)

			gen11786 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11787 := Call(__e, gen11786, V1595)

			__e.TailApply(gen11785, gen11787)

			return

		} else {
			gen11772 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11772)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5dbq_6, gen11792)

	gen11841 := MakeNative(func(__e *ControlFlow) {
		V1597 := __e.Get(1)
		_ = V1597
		gen11810 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen11806 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11807 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11808 := Call(__e, gen11807)

			gen11809 := Call(__e, gen11806, YaccParse, gen11808)

			if True == gen11809 {
				gen11803 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen11797 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11798 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11799 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11800 := Call(__e, gen11799)

					gen11801 := Call(__e, gen11798, gen11800, Parse___5e_6)

					gen11802 := Call(__e, gen11797, gen11801)

					if True == gen11802 {
						gen11794 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11795 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11796 := Call(__e, gen11795, Parse___5e_6)

						__e.TailApply(gen11794, gen11796, Nil)

						return

					} else {
						gen11793 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11793)

						return

					}

				}, 1)

				gen11804 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen11805 := Call(__e, gen11804, V1597)

				__e.TailApply(gen11803, gen11805)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen11837 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5strc_6 := __e.Get(1)
			_ = Parse__shen_4_5strc_6
			gen11831 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen11832 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen11833 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen11834 := Call(__e, gen11833)

			gen11835 := Call(__e, gen11832, gen11834, Parse__shen_4_5strc_6)

			gen11836 := Call(__e, gen11831, gen11835)

			if True == gen11836 {
				gen11828 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5strcontents_6 := __e.Get(1)
					_ = Parse__shen_4_5strcontents_6
					gen11822 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen11823 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen11824 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen11825 := Call(__e, gen11824)

					gen11826 := Call(__e, gen11823, gen11825, Parse__shen_4_5strcontents_6)

					gen11827 := Call(__e, gen11822, gen11826)

					if True == gen11827 {
						gen11813 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen11814 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen11815 := Call(__e, gen11814, Parse__shen_4_5strcontents_6)

						gen11816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen11817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11818 := Call(__e, gen11817, Parse__shen_4_5strc_6)

						gen11819 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen11820 := Call(__e, gen11819, Parse__shen_4_5strcontents_6)

						gen11821 := Call(__e, gen11816, gen11818, gen11820)

						__e.TailApply(gen11813, gen11815, gen11821)

						return

					} else {
						gen11812 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen11812)

						return

					}

				}, 1)

				gen11829 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5strcontents_6)

				gen11830 := Call(__e, gen11829, Parse__shen_4_5strc_6)

				__e.TailApply(gen11828, gen11830)

				return

			} else {
				gen11811 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen11811)

				return

			}

		}, 1)

		gen11838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5strc_6)

		gen11839 := Call(__e, gen11838, V1597)

		gen11840 := Call(__e, gen11837, gen11839)

		__e.TailApply(gen11810, gen11840)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5strcontents_6, gen11841)

	gen11861 := MakeNative(func(__e *ControlFlow) {
		V1599 := __e.Get(1)
		_ = V1599
		gen11857 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11858 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11859 := Call(__e, gen11858, V1599)

		gen11860 := Call(__e, gen11857, gen11859)

		if True == gen11860 {
			gen11854 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen11843 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11844 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen11845 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen11846 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				gen11847 := Call(__e, gen11846, V1599)

				gen11848 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen11849 := Call(__e, gen11848, V1599)

				gen11850 := Call(__e, gen11845, gen11847, gen11849)

				gen11851 := Call(__e, gen11844, gen11850)

				gen11852 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				gen11853 := Call(__e, gen11852, Parse__Char)

				__e.TailApply(gen11843, gen11851, gen11853)

				return

			}, 1)

			gen11855 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11856 := Call(__e, gen11855, V1599)

			__e.TailApply(gen11854, gen11856)

			return

		} else {
			gen11842 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11842)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5byte_6, gen11861)

	gen11886 := MakeNative(func(__e *ControlFlow) {
		V1601 := __e.Get(1)
		_ = V1601
		gen11882 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen11883 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen11884 := Call(__e, gen11883, V1601)

		gen11885 := Call(__e, gen11882, gen11884)

		if True == gen11885 {
			gen11879 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen11875 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				gen11876 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen11877 := Call(__e, gen11876, Parse__Char, MakeNumber(34))

				gen11878 := Call(__e, gen11875, gen11877)

				if True == gen11878 {
					gen11864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11865 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen11866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen11867 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen11868 := Call(__e, gen11867, V1601)

					gen11869 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen11870 := Call(__e, gen11869, V1601)

					gen11871 := Call(__e, gen11866, gen11868, gen11870)

					gen11872 := Call(__e, gen11865, gen11871)

					gen11873 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

					gen11874 := Call(__e, gen11873, Parse__Char)

					__e.TailApply(gen11864, gen11872, gen11874)

					return

				} else {
					gen11863 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen11863)

					return

				}

			}, 1)

			gen11880 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen11881 := Call(__e, gen11880, V1601)

			__e.TailApply(gen11879, gen11881)

			return

		} else {
			gen11862 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen11862)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5strc_6, gen11886)

	gen12149 := MakeNative(func(__e *ControlFlow) {
		V1603 := __e.Get(1)
		_ = V1603
		gen12120 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12116 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12117 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12118 := Call(__e, gen12117)

			gen12119 := Call(__e, gen12116, YaccParse, gen12118)

			if True == gen12119 {
				gen12089 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					gen12085 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12086 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12087 := Call(__e, gen12086)

					gen12088 := Call(__e, gen12085, YaccParse, gen12087)

					if True == gen12088 {
						gen12012 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							gen12008 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12009 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen12010 := Call(__e, gen12009)

							gen12011 := Call(__e, gen12008, YaccParse, gen12010)

							if True == gen12011 {
								gen11961 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									gen11957 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen11958 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen11959 := Call(__e, gen11958)

									gen11960 := Call(__e, gen11957, YaccParse, gen11959)

									if True == gen11960 {
										gen11910 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											gen11906 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen11907 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen11908 := Call(__e, gen11907)

											gen11909 := Call(__e, gen11906, YaccParse, gen11908)

											if True == gen11909 {
												gen11903 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5digits_6 := __e.Get(1)
													_ = Parse__shen_4_5digits_6
													gen11897 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													gen11898 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen11899 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen11900 := Call(__e, gen11899)

													gen11901 := Call(__e, gen11898, gen11900, Parse__shen_4_5digits_6)

													gen11902 := Call(__e, gen11897, gen11901)

													if True == gen11902 {
														gen11888 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														gen11889 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen11890 := Call(__e, gen11889, Parse__shen_4_5digits_6)

														gen11891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

														gen11892 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

														gen11893 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														gen11894 := Call(__e, gen11893, Parse__shen_4_5digits_6)

														gen11895 := Call(__e, gen11892, gen11894)

														gen11896 := Call(__e, gen11891, gen11895, MakeNumber(0))

														__e.TailApply(gen11888, gen11890, gen11896)

														return

													} else {
														gen11887 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(gen11887)

														return

													}

												}, 1)

												gen11904 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

												gen11905 := Call(__e, gen11904, V1603)

												__e.TailApply(gen11903, gen11905)

												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										gen11953 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5predigits_6 := __e.Get(1)
											_ = Parse__shen_4_5predigits_6
											gen11947 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen11948 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen11949 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen11950 := Call(__e, gen11949)

											gen11951 := Call(__e, gen11948, gen11950, Parse__shen_4_5predigits_6)

											gen11952 := Call(__e, gen11947, gen11951)

											if True == gen11952 {
												gen11944 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5stop_6 := __e.Get(1)
													_ = Parse__shen_4_5stop_6
													gen11938 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													gen11939 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen11940 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen11941 := Call(__e, gen11940)

													gen11942 := Call(__e, gen11939, gen11941, Parse__shen_4_5stop_6)

													gen11943 := Call(__e, gen11938, gen11942)

													if True == gen11943 {
														gen11935 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5postdigits_6 := __e.Get(1)
															_ = Parse__shen_4_5postdigits_6
															gen11929 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															gen11930 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															gen11931 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															gen11932 := Call(__e, gen11931)

															gen11933 := Call(__e, gen11930, gen11932, Parse__shen_4_5postdigits_6)

															gen11934 := Call(__e, gen11929, gen11933)

															if True == gen11934 {
																gen11914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																gen11915 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																gen11916 := Call(__e, gen11915, Parse__shen_4_5postdigits_6)

																gen11917 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

																gen11918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

																gen11919 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

																gen11920 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																gen11921 := Call(__e, gen11920, Parse__shen_4_5predigits_6)

																gen11922 := Call(__e, gen11919, gen11921)

																gen11923 := Call(__e, gen11918, gen11922, MakeNumber(0))

																gen11924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4post)

																gen11925 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																gen11926 := Call(__e, gen11925, Parse__shen_4_5postdigits_6)

																gen11927 := Call(__e, gen11924, gen11926, MakeNumber(1))

																gen11928 := Call(__e, gen11917, gen11923, gen11927)

																__e.TailApply(gen11914, gen11916, gen11928)

																return

															} else {
																gen11913 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(gen11913)

																return

															}

														}, 1)

														gen11936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5postdigits_6)

														gen11937 := Call(__e, gen11936, Parse__shen_4_5stop_6)

														__e.TailApply(gen11935, gen11937)

														return

													} else {
														gen11912 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(gen11912)

														return

													}

												}, 1)

												gen11945 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5stop_6)

												gen11946 := Call(__e, gen11945, Parse__shen_4_5predigits_6)

												__e.TailApply(gen11944, gen11946)

												return

											} else {
												gen11911 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen11911)

												return

											}

										}, 1)

										gen11954 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5predigits_6)

										gen11955 := Call(__e, gen11954, V1603)

										gen11956 := Call(__e, gen11953, gen11955)

										__e.TailApply(gen11910, gen11956)

										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								gen12004 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5digits_6 := __e.Get(1)
									_ = Parse__shen_4_5digits_6
									gen11998 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen11999 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen12000 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen12001 := Call(__e, gen12000)

									gen12002 := Call(__e, gen11999, gen12001, Parse__shen_4_5digits_6)

									gen12003 := Call(__e, gen11998, gen12002)

									if True == gen12003 {
										gen11995 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5E_6 := __e.Get(1)
											_ = Parse__shen_4_5E_6
											gen11989 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen11990 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen11991 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen11992 := Call(__e, gen11991)

											gen11993 := Call(__e, gen11990, gen11992, Parse__shen_4_5E_6)

											gen11994 := Call(__e, gen11989, gen11993)

											if True == gen11994 {
												gen11986 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5log10_6 := __e.Get(1)
													_ = Parse__shen_4_5log10_6
													gen11980 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													gen11981 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen11982 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen11983 := Call(__e, gen11982)

													gen11984 := Call(__e, gen11981, gen11983, Parse__shen_4_5log10_6)

													gen11985 := Call(__e, gen11980, gen11984)

													if True == gen11985 {
														gen11965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														gen11966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														gen11967 := Call(__e, gen11966, Parse__shen_4_5log10_6)

														gen11968 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

														gen11969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

														gen11970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														gen11971 := Call(__e, gen11970, Parse__shen_4_5log10_6)

														gen11972 := Call(__e, gen11969, MakeNumber(10), gen11971)

														gen11973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

														gen11974 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

														gen11975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														gen11976 := Call(__e, gen11975, Parse__shen_4_5digits_6)

														gen11977 := Call(__e, gen11974, gen11976)

														gen11978 := Call(__e, gen11973, gen11977, MakeNumber(0))

														gen11979 := Call(__e, gen11968, gen11972, gen11978)

														__e.TailApply(gen11965, gen11967, gen11979)

														return

													} else {
														gen11964 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(gen11964)

														return

													}

												}, 1)

												gen11987 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5log10_6)

												gen11988 := Call(__e, gen11987, Parse__shen_4_5E_6)

												__e.TailApply(gen11986, gen11988)

												return

											} else {
												gen11963 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen11963)

												return

											}

										}, 1)

										gen11996 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5E_6)

										gen11997 := Call(__e, gen11996, Parse__shen_4_5digits_6)

										__e.TailApply(gen11995, gen11997)

										return

									} else {
										gen11962 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen11962)

										return

									}

								}, 1)

								gen12005 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

								gen12006 := Call(__e, gen12005, V1603)

								gen12007 := Call(__e, gen12004, gen12006)

								__e.TailApply(gen11961, gen12007)

								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						gen12081 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5predigits_6 := __e.Get(1)
							_ = Parse__shen_4_5predigits_6
							gen12075 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen12076 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12077 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen12078 := Call(__e, gen12077)

							gen12079 := Call(__e, gen12076, gen12078, Parse__shen_4_5predigits_6)

							gen12080 := Call(__e, gen12075, gen12079)

							if True == gen12080 {
								gen12072 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5stop_6 := __e.Get(1)
									_ = Parse__shen_4_5stop_6
									gen12066 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen12067 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen12068 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen12069 := Call(__e, gen12068)

									gen12070 := Call(__e, gen12067, gen12069, Parse__shen_4_5stop_6)

									gen12071 := Call(__e, gen12066, gen12070)

									if True == gen12071 {
										gen12063 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5postdigits_6 := __e.Get(1)
											_ = Parse__shen_4_5postdigits_6
											gen12057 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen12058 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen12059 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen12060 := Call(__e, gen12059)

											gen12061 := Call(__e, gen12058, gen12060, Parse__shen_4_5postdigits_6)

											gen12062 := Call(__e, gen12057, gen12061)

											if True == gen12062 {
												gen12054 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5E_6 := __e.Get(1)
													_ = Parse__shen_4_5E_6
													gen12048 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													gen12049 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen12050 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													gen12051 := Call(__e, gen12050)

													gen12052 := Call(__e, gen12049, gen12051, Parse__shen_4_5E_6)

													gen12053 := Call(__e, gen12048, gen12052)

													if True == gen12053 {
														gen12045 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5log10_6 := __e.Get(1)
															_ = Parse__shen_4_5log10_6
															gen12039 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															gen12040 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															gen12041 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															gen12042 := Call(__e, gen12041)

															gen12043 := Call(__e, gen12040, gen12042, Parse__shen_4_5log10_6)

															gen12044 := Call(__e, gen12039, gen12043)

															if True == gen12044 {
																gen12018 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																gen12019 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																gen12020 := Call(__e, gen12019, Parse__shen_4_5log10_6)

																gen12021 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

																gen12022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

																gen12023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																gen12024 := Call(__e, gen12023, Parse__shen_4_5log10_6)

																gen12025 := Call(__e, gen12022, MakeNumber(10), gen12024)

																gen12026 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

																gen12027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

																gen12028 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

																gen12029 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																gen12030 := Call(__e, gen12029, Parse__shen_4_5predigits_6)

																gen12031 := Call(__e, gen12028, gen12030)

																gen12032 := Call(__e, gen12027, gen12031, MakeNumber(0))

																gen12033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4post)

																gen12034 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																gen12035 := Call(__e, gen12034, Parse__shen_4_5postdigits_6)

																gen12036 := Call(__e, gen12033, gen12035, MakeNumber(1))

																gen12037 := Call(__e, gen12026, gen12032, gen12036)

																gen12038 := Call(__e, gen12021, gen12025, gen12037)

																__e.TailApply(gen12018, gen12020, gen12038)

																return

															} else {
																gen12017 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(gen12017)

																return

															}

														}, 1)

														gen12046 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5log10_6)

														gen12047 := Call(__e, gen12046, Parse__shen_4_5E_6)

														__e.TailApply(gen12045, gen12047)

														return

													} else {
														gen12016 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(gen12016)

														return

													}

												}, 1)

												gen12055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5E_6)

												gen12056 := Call(__e, gen12055, Parse__shen_4_5postdigits_6)

												__e.TailApply(gen12054, gen12056)

												return

											} else {
												gen12015 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen12015)

												return

											}

										}, 1)

										gen12064 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5postdigits_6)

										gen12065 := Call(__e, gen12064, Parse__shen_4_5stop_6)

										__e.TailApply(gen12063, gen12065)

										return

									} else {
										gen12014 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen12014)

										return

									}

								}, 1)

								gen12073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5stop_6)

								gen12074 := Call(__e, gen12073, Parse__shen_4_5predigits_6)

								__e.TailApply(gen12072, gen12074)

								return

							} else {
								gen12013 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen12013)

								return

							}

						}, 1)

						gen12082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5predigits_6)

						gen12083 := Call(__e, gen12082, V1603)

						gen12084 := Call(__e, gen12081, gen12083)

						__e.TailApply(gen12012, gen12084)

						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				gen12112 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5plus_6 := __e.Get(1)
					_ = Parse__shen_4_5plus_6
					gen12106 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12107 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12108 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12109 := Call(__e, gen12108)

					gen12110 := Call(__e, gen12107, gen12109, Parse__shen_4_5plus_6)

					gen12111 := Call(__e, gen12106, gen12110)

					if True == gen12111 {
						gen12103 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5number_6 := __e.Get(1)
							_ = Parse__shen_4_5number_6
							gen12097 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen12098 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12099 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen12100 := Call(__e, gen12099)

							gen12101 := Call(__e, gen12098, gen12100, Parse__shen_4_5number_6)

							gen12102 := Call(__e, gen12097, gen12101)

							if True == gen12102 {
								gen12092 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen12093 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen12094 := Call(__e, gen12093, Parse__shen_4_5number_6)

								gen12095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen12096 := Call(__e, gen12095, Parse__shen_4_5number_6)

								__e.TailApply(gen12092, gen12094, gen12096)

								return

							} else {
								gen12091 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen12091)

								return

							}

						}, 1)

						gen12104 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5number_6)

						gen12105 := Call(__e, gen12104, Parse__shen_4_5plus_6)

						__e.TailApply(gen12103, gen12105)

						return

					} else {
						gen12090 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12090)

						return

					}

				}, 1)

				gen12113 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5plus_6)

				gen12114 := Call(__e, gen12113, V1603)

				gen12115 := Call(__e, gen12112, gen12114)

				__e.TailApply(gen12089, gen12115)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12145 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5minus_6 := __e.Get(1)
			_ = Parse__shen_4_5minus_6
			gen12139 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12140 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12141 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12142 := Call(__e, gen12141)

			gen12143 := Call(__e, gen12140, gen12142, Parse__shen_4_5minus_6)

			gen12144 := Call(__e, gen12139, gen12143)

			if True == gen12144 {
				gen12136 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5number_6 := __e.Get(1)
					_ = Parse__shen_4_5number_6
					gen12130 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12131 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12132 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12133 := Call(__e, gen12132)

					gen12134 := Call(__e, gen12131, gen12133, Parse__shen_4_5number_6)

					gen12135 := Call(__e, gen12130, gen12134)

					if True == gen12135 {
						gen12123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12124 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12125 := Call(__e, gen12124, Parse__shen_4_5number_6)

						gen12126 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						gen12127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen12128 := Call(__e, gen12127, Parse__shen_4_5number_6)

						gen12129 := Call(__e, gen12126, MakeNumber(0), gen12128)

						__e.TailApply(gen12123, gen12125, gen12129)

						return

					} else {
						gen12122 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12122)

						return

					}

				}, 1)

				gen12137 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5number_6)

				gen12138 := Call(__e, gen12137, Parse__shen_4_5minus_6)

				__e.TailApply(gen12136, gen12138)

				return

			} else {
				gen12121 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12121)

				return

			}

		}, 1)

		gen12146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5minus_6)

		gen12147 := Call(__e, gen12146, V1603)

		gen12148 := Call(__e, gen12145, gen12147)

		__e.TailApply(gen12120, gen12148)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5number_6, gen12149)

	gen12170 := MakeNative(func(__e *ControlFlow) {
		V1606 := __e.Get(1)
		_ = V1606
		gen12165 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12166 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12167 := Call(__e, gen12166, V1606)

		gen12168 := Call(__e, gen12165, gen12167)

		var gen12169 Obj

		if True == gen12168 {
			gen12161 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12163 := Call(__e, gen12162, V1606)

			gen12164 := Call(__e, gen12161, MakeNumber(101), gen12163)

			if True == gen12164 {
				gen12169 = True
			} else {
				gen12169 = False
			}

		} else {
			gen12169 = False
		}

		if True == gen12169 {
			gen12154 := MakeNative(func(__e *ControlFlow) {
				NewStream1604 := __e.Get(1)
				_ = NewStream1604
				gen12151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12152 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12153 := Call(__e, gen12152, NewStream1604)

				__e.TailApply(gen12151, gen12153, symshen_4skip)

				return

			}, 1)

			gen12155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen12156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen12157 := Call(__e, gen12156, V1606)

			gen12158 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen12159 := Call(__e, gen12158, V1606)

			gen12160 := Call(__e, gen12155, gen12157, gen12159)

			__e.TailApply(gen12154, gen12160)

			return

		} else {
			gen12150 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12150)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5E_6, gen12170)

	gen12227 := MakeNative(func(__e *ControlFlow) {
		V1608 := __e.Get(1)
		_ = V1608
		gen12194 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12190 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12191 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12192 := Call(__e, gen12191)

			gen12193 := Call(__e, gen12190, YaccParse, gen12192)

			if True == gen12193 {
				gen12187 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					gen12181 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12182 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12183 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12184 := Call(__e, gen12183)

					gen12185 := Call(__e, gen12182, gen12184, Parse__shen_4_5digits_6)

					gen12186 := Call(__e, gen12181, gen12185)

					if True == gen12186 {
						gen12172 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12173 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12174 := Call(__e, gen12173, Parse__shen_4_5digits_6)

						gen12175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

						gen12176 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						gen12177 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen12178 := Call(__e, gen12177, Parse__shen_4_5digits_6)

						gen12179 := Call(__e, gen12176, gen12178)

						gen12180 := Call(__e, gen12175, gen12179, MakeNumber(0))

						__e.TailApply(gen12172, gen12174, gen12180)

						return

					} else {
						gen12171 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12171)

						return

					}

				}, 1)

				gen12188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

				gen12189 := Call(__e, gen12188, V1608)

				__e.TailApply(gen12187, gen12189)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12223 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5minus_6 := __e.Get(1)
			_ = Parse__shen_4_5minus_6
			gen12217 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12218 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12219 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12220 := Call(__e, gen12219)

			gen12221 := Call(__e, gen12218, gen12220, Parse__shen_4_5minus_6)

			gen12222 := Call(__e, gen12217, gen12221)

			if True == gen12222 {
				gen12214 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					gen12208 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12209 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12210 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12211 := Call(__e, gen12210)

					gen12212 := Call(__e, gen12209, gen12211, Parse__shen_4_5digits_6)

					gen12213 := Call(__e, gen12208, gen12212)

					if True == gen12213 {
						gen12197 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12198 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12199 := Call(__e, gen12198, Parse__shen_4_5digits_6)

						gen12200 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						gen12201 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

						gen12202 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						gen12203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen12204 := Call(__e, gen12203, Parse__shen_4_5digits_6)

						gen12205 := Call(__e, gen12202, gen12204)

						gen12206 := Call(__e, gen12201, gen12205, MakeNumber(0))

						gen12207 := Call(__e, gen12200, MakeNumber(0), gen12206)

						__e.TailApply(gen12197, gen12199, gen12207)

						return

					} else {
						gen12196 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12196)

						return

					}

				}, 1)

				gen12215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

				gen12216 := Call(__e, gen12215, Parse__shen_4_5minus_6)

				__e.TailApply(gen12214, gen12216)

				return

			} else {
				gen12195 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12195)

				return

			}

		}, 1)

		gen12224 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5minus_6)

		gen12225 := Call(__e, gen12224, V1608)

		gen12226 := Call(__e, gen12223, gen12225)

		__e.TailApply(gen12194, gen12226)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5log10_6, gen12227)

	gen12248 := MakeNative(func(__e *ControlFlow) {
		V1610 := __e.Get(1)
		_ = V1610
		gen12244 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12245 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12246 := Call(__e, gen12245, V1610)

		gen12247 := Call(__e, gen12244, gen12246)

		if True == gen12247 {
			gen12241 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen12239 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen12240 := Call(__e, gen12239, Parse__Char, MakeNumber(43))

				if True == gen12240 {
					gen12230 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12231 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12232 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12233 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen12234 := Call(__e, gen12233, V1610)

					gen12235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen12236 := Call(__e, gen12235, V1610)

					gen12237 := Call(__e, gen12232, gen12234, gen12236)

					gen12238 := Call(__e, gen12231, gen12237)

					__e.TailApply(gen12230, gen12238, Parse__Char)

					return

				} else {
					gen12229 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen12229)

					return

				}

			}, 1)

			gen12242 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12243 := Call(__e, gen12242, V1610)

			__e.TailApply(gen12241, gen12243)

			return

		} else {
			gen12228 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12228)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5plus_6, gen12248)

	gen12269 := MakeNative(func(__e *ControlFlow) {
		V1612 := __e.Get(1)
		_ = V1612
		gen12265 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12266 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12267 := Call(__e, gen12266, V1612)

		gen12268 := Call(__e, gen12265, gen12267)

		if True == gen12268 {
			gen12262 := MakeNative(func(__e *ControlFlow) {
				Parse__Char := __e.Get(1)
				_ = Parse__Char
				gen12260 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen12261 := Call(__e, gen12260, Parse__Char, MakeNumber(46))

				if True == gen12261 {
					gen12251 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12252 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen12255 := Call(__e, gen12254, V1612)

					gen12256 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen12257 := Call(__e, gen12256, V1612)

					gen12258 := Call(__e, gen12253, gen12255, gen12257)

					gen12259 := Call(__e, gen12252, gen12258)

					__e.TailApply(gen12251, gen12259, Parse__Char)

					return

				} else {
					gen12250 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen12250)

					return

				}

			}, 1)

			gen12263 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12264 := Call(__e, gen12263, V1612)

			__e.TailApply(gen12262, gen12264)

			return

		} else {
			gen12249 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12249)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5stop_6, gen12269)

	gen12304 := MakeNative(func(__e *ControlFlow) {
		V1614 := __e.Get(1)
		_ = V1614
		gen12287 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12283 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12284 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12285 := Call(__e, gen12284)

			gen12286 := Call(__e, gen12283, YaccParse, gen12285)

			if True == gen12286 {
				gen12280 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen12274 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12275 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12276 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12277 := Call(__e, gen12276)

					gen12278 := Call(__e, gen12275, gen12277, Parse___5e_6)

					gen12279 := Call(__e, gen12274, gen12278)

					if True == gen12279 {
						gen12271 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12272 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12273 := Call(__e, gen12272, Parse___5e_6)

						__e.TailApply(gen12271, gen12273, Nil)

						return

					} else {
						gen12270 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12270)

						return

					}

				}, 1)

				gen12281 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen12282 := Call(__e, gen12281, V1614)

				__e.TailApply(gen12280, gen12282)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12300 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digits_6 := __e.Get(1)
			_ = Parse__shen_4_5digits_6
			gen12294 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12295 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12296 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12297 := Call(__e, gen12296)

			gen12298 := Call(__e, gen12295, gen12297, Parse__shen_4_5digits_6)

			gen12299 := Call(__e, gen12294, gen12298)

			if True == gen12299 {
				gen12289 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12290 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12291 := Call(__e, gen12290, Parse__shen_4_5digits_6)

				gen12292 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen12293 := Call(__e, gen12292, Parse__shen_4_5digits_6)

				__e.TailApply(gen12289, gen12291, gen12293)

				return

			} else {
				gen12288 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12288)

				return

			}

		}, 1)

		gen12301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

		gen12302 := Call(__e, gen12301, V1614)

		gen12303 := Call(__e, gen12300, gen12302)

		__e.TailApply(gen12287, gen12303)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5predigits_6, gen12304)

	gen12320 := MakeNative(func(__e *ControlFlow) {
		V1616 := __e.Get(1)
		_ = V1616
		gen12317 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digits_6 := __e.Get(1)
			_ = Parse__shen_4_5digits_6
			gen12311 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12312 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12313 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12314 := Call(__e, gen12313)

			gen12315 := Call(__e, gen12312, gen12314, Parse__shen_4_5digits_6)

			gen12316 := Call(__e, gen12311, gen12315)

			if True == gen12316 {
				gen12306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12307 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12308 := Call(__e, gen12307, Parse__shen_4_5digits_6)

				gen12309 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen12310 := Call(__e, gen12309, Parse__shen_4_5digits_6)

				__e.TailApply(gen12306, gen12308, gen12310)

				return

			} else {
				gen12305 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12305)

				return

			}

		}, 1)

		gen12318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

		gen12319 := Call(__e, gen12318, V1616)

		__e.TailApply(gen12317, gen12319)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5postdigits_6, gen12320)

	gen12373 := MakeNative(func(__e *ControlFlow) {
		V1618 := __e.Get(1)
		_ = V1618
		gen12342 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12338 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12339 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12340 := Call(__e, gen12339)

			gen12341 := Call(__e, gen12338, YaccParse, gen12340)

			if True == gen12341 {
				gen12335 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digit_6 := __e.Get(1)
					_ = Parse__shen_4_5digit_6
					gen12329 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12330 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12331 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12332 := Call(__e, gen12331)

					gen12333 := Call(__e, gen12330, gen12332, Parse__shen_4_5digit_6)

					gen12334 := Call(__e, gen12329, gen12333)

					if True == gen12334 {
						gen12322 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12323 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12324 := Call(__e, gen12323, Parse__shen_4_5digit_6)

						gen12325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen12326 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen12327 := Call(__e, gen12326, Parse__shen_4_5digit_6)

						gen12328 := Call(__e, gen12325, gen12327, Nil)

						__e.TailApply(gen12322, gen12324, gen12328)

						return

					} else {
						gen12321 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12321)

						return

					}

				}, 1)

				gen12336 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digit_6)

				gen12337 := Call(__e, gen12336, V1618)

				__e.TailApply(gen12335, gen12337)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12369 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5digit_6 := __e.Get(1)
			_ = Parse__shen_4_5digit_6
			gen12363 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12364 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12365 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12366 := Call(__e, gen12365)

			gen12367 := Call(__e, gen12364, gen12366, Parse__shen_4_5digit_6)

			gen12368 := Call(__e, gen12363, gen12367)

			if True == gen12368 {
				gen12360 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5digits_6 := __e.Get(1)
					_ = Parse__shen_4_5digits_6
					gen12354 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12355 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12356 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12357 := Call(__e, gen12356)

					gen12358 := Call(__e, gen12355, gen12357, Parse__shen_4_5digits_6)

					gen12359 := Call(__e, gen12354, gen12358)

					if True == gen12359 {
						gen12345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12346 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12347 := Call(__e, gen12346, Parse__shen_4_5digits_6)

						gen12348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen12349 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen12350 := Call(__e, gen12349, Parse__shen_4_5digit_6)

						gen12351 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen12352 := Call(__e, gen12351, Parse__shen_4_5digits_6)

						gen12353 := Call(__e, gen12348, gen12350, gen12352)

						__e.TailApply(gen12345, gen12347, gen12353)

						return

					} else {
						gen12344 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12344)

						return

					}

				}, 1)

				gen12361 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digits_6)

				gen12362 := Call(__e, gen12361, Parse__shen_4_5digit_6)

				__e.TailApply(gen12360, gen12362)

				return

			} else {
				gen12343 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12343)

				return

			}

		}, 1)

		gen12370 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5digit_6)

		gen12371 := Call(__e, gen12370, V1618)

		gen12372 := Call(__e, gen12369, gen12371)

		__e.TailApply(gen12342, gen12372)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5digits_6, gen12373)

	gen12396 := MakeNative(func(__e *ControlFlow) {
		V1620 := __e.Get(1)
		_ = V1620
		gen12392 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12393 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12394 := Call(__e, gen12393, V1620)

		gen12395 := Call(__e, gen12392, gen12394)

		if True == gen12395 {
			gen12389 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen12387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4numbyte_2)

				gen12388 := Call(__e, gen12387, Parse__X)

				if True == gen12388 {
					gen12376 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12377 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12378 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen12380 := Call(__e, gen12379, V1620)

					gen12381 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen12382 := Call(__e, gen12381, V1620)

					gen12383 := Call(__e, gen12378, gen12380, gen12382)

					gen12384 := Call(__e, gen12377, gen12383)

					gen12385 := Call(__e, PrimNS1Value(symns2_1value), symshen_4byte_1_6digit)

					gen12386 := Call(__e, gen12385, Parse__X)

					__e.TailApply(gen12376, gen12384, gen12386)

					return

				} else {
					gen12375 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen12375)

					return

				}

			}, 1)

			gen12390 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12391 := Call(__e, gen12390, V1620)

			__e.TailApply(gen12389, gen12391)

			return

		} else {
			gen12374 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12374)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5digit_6, gen12396)

	gen12419 := MakeNative(func(__e *ControlFlow) {
		V1622 := __e.Get(1)
		_ = V1622
		gen12417 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen12418 := Call(__e, gen12417, MakeNumber(48), V1622)

		if True == gen12418 {
			__e.Return(MakeNumber(0))
			return
		} else {
			gen12415 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12416 := Call(__e, gen12415, MakeNumber(49), V1622)

			if True == gen12416 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen12413 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen12414 := Call(__e, gen12413, MakeNumber(50), V1622)

				if True == gen12414 {
					__e.Return(MakeNumber(2))
					return
				} else {
					gen12411 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12412 := Call(__e, gen12411, MakeNumber(51), V1622)

					if True == gen12412 {
						__e.Return(MakeNumber(3))
						return
					} else {
						gen12409 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen12410 := Call(__e, gen12409, MakeNumber(52), V1622)

						if True == gen12410 {
							__e.Return(MakeNumber(4))
							return
						} else {
							gen12407 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12408 := Call(__e, gen12407, MakeNumber(53), V1622)

							if True == gen12408 {
								__e.Return(MakeNumber(5))
								return
							} else {
								gen12405 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen12406 := Call(__e, gen12405, MakeNumber(54), V1622)

								if True == gen12406 {
									__e.Return(MakeNumber(6))
									return
								} else {
									gen12403 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen12404 := Call(__e, gen12403, MakeNumber(55), V1622)

									if True == gen12404 {
										__e.Return(MakeNumber(7))
										return
									} else {
										gen12401 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen12402 := Call(__e, gen12401, MakeNumber(56), V1622)

										if True == gen12402 {
											__e.Return(MakeNumber(8))
											return
										} else {
											gen12399 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen12400 := Call(__e, gen12399, MakeNumber(57), V1622)

											if True == gen12400 {
												__e.Return(MakeNumber(9))
												return
											} else {
												if True == True {
													gen12398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

													__e.TailApply(gen12398, symshen_4byte_1_6digit)

													return

												} else {
													gen12397 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

													__e.TailApply(gen12397, MakeString("no cond match"))

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

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4byte_1_6digit, gen12419)

	gen12439 := MakeNative(func(__e *ControlFlow) {
		V1627 := __e.Get(1)
		_ = V1627
		V1628 := __e.Get(2)
		_ = V1628
		gen12437 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen12438 := Call(__e, gen12437, Nil, V1627)

		if True == gen12438 {
			__e.Return(MakeNumber(0))
			return
		} else {
			gen12435 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen12436 := Call(__e, gen12435, V1627)

			if True == gen12436 {
				gen12422 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen12423 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				gen12424 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				gen12425 := Call(__e, gen12424, MakeNumber(10), V1628)

				gen12426 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12427 := Call(__e, gen12426, V1627)

				gen12428 := Call(__e, gen12423, gen12425, gen12427)

				gen12429 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pre)

				gen12430 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen12431 := Call(__e, gen12430, V1627)

				gen12432 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen12433 := Call(__e, gen12432, V1628, MakeNumber(1))

				gen12434 := Call(__e, gen12429, gen12431, gen12433)

				__e.TailApply(gen12422, gen12428, gen12434)

				return

			} else {
				if True == True {
					gen12421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen12421, symshen_4pre)

					return

				} else {
					gen12420 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen12420, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4pre, gen12439)

	gen12461 := MakeNative(func(__e *ControlFlow) {
		V1633 := __e.Get(1)
		_ = V1633
		V1634 := __e.Get(2)
		_ = V1634
		gen12459 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen12460 := Call(__e, gen12459, Nil, V1633)

		if True == gen12460 {
			__e.Return(MakeNumber(0))
			return
		} else {
			gen12457 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen12458 := Call(__e, gen12457, V1633)

			if True == gen12458 {
				gen12442 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen12443 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				gen12444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				gen12445 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				gen12446 := Call(__e, gen12445, MakeNumber(0), V1634)

				gen12447 := Call(__e, gen12444, MakeNumber(10), gen12446)

				gen12448 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12449 := Call(__e, gen12448, V1633)

				gen12450 := Call(__e, gen12443, gen12447, gen12449)

				gen12451 := Call(__e, PrimNS1Value(symns2_1value), symshen_4post)

				gen12452 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen12453 := Call(__e, gen12452, V1633)

				gen12454 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen12455 := Call(__e, gen12454, V1634, MakeNumber(1))

				gen12456 := Call(__e, gen12451, gen12453, gen12455)

				__e.TailApply(gen12442, gen12450, gen12456)

				return

			} else {
				if True == True {
					gen12441 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen12441, symshen_4post)

					return

				} else {
					gen12440 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen12440, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4post, gen12461)

	gen12479 := MakeNative(func(__e *ControlFlow) {
		V1639 := __e.Get(1)
		_ = V1639
		V1640 := __e.Get(2)
		_ = V1640
		gen12477 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen12478 := Call(__e, gen12477, MakeNumber(0), V1640)

		if True == gen12478 {
			__e.Return(MakeNumber(1))
			return
		} else {
			gen12475 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			gen12476 := Call(__e, gen12475, V1640, MakeNumber(0))

			if True == gen12476 {
				gen12470 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				gen12471 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

				gen12472 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				gen12473 := Call(__e, gen12472, V1640, MakeNumber(1))

				gen12474 := Call(__e, gen12471, V1639, gen12473)

				__e.TailApply(gen12470, V1639, gen12474)

				return

			} else {
				if True == True {
					gen12463 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

					gen12464 := Call(__e, PrimNS1Value(symns2_1value), sym_c)

					gen12465 := Call(__e, PrimNS1Value(symns2_1value), symshen_4expt)

					gen12466 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					gen12467 := Call(__e, gen12466, V1640, MakeNumber(1))

					gen12468 := Call(__e, gen12465, V1639, gen12467)

					gen12469 := Call(__e, gen12464, gen12468, V1639)

					__e.TailApply(gen12463, MakeNumber(1), gen12469)

					return

				} else {
					gen12462 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen12462, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4expt, gen12479)

	gen12495 := MakeNative(func(__e *ControlFlow) {
		V1642 := __e.Get(1)
		_ = V1642
		gen12492 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5st__input_6 := __e.Get(1)
			_ = Parse__shen_4_5st__input_6
			gen12486 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12487 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12488 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12489 := Call(__e, gen12488)

			gen12490 := Call(__e, gen12487, gen12489, Parse__shen_4_5st__input_6)

			gen12491 := Call(__e, gen12486, gen12490)

			if True == gen12491 {
				gen12481 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12482 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12483 := Call(__e, gen12482, Parse__shen_4_5st__input_6)

				gen12484 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen12485 := Call(__e, gen12484, Parse__shen_4_5st__input_6)

				__e.TailApply(gen12481, gen12483, gen12485)

				return

			} else {
				gen12480 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12480)

				return

			}

		}, 1)

		gen12493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

		gen12494 := Call(__e, gen12493, V1642)

		__e.TailApply(gen12492, gen12494)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input1_6, gen12495)

	gen12511 := MakeNative(func(__e *ControlFlow) {
		V1644 := __e.Get(1)
		_ = V1644
		gen12508 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5st__input_6 := __e.Get(1)
			_ = Parse__shen_4_5st__input_6
			gen12502 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12503 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12504 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12505 := Call(__e, gen12504)

			gen12506 := Call(__e, gen12503, gen12505, Parse__shen_4_5st__input_6)

			gen12507 := Call(__e, gen12502, gen12506)

			if True == gen12507 {
				gen12497 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12498 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12499 := Call(__e, gen12498, Parse__shen_4_5st__input_6)

				gen12500 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen12501 := Call(__e, gen12500, Parse__shen_4_5st__input_6)

				__e.TailApply(gen12497, gen12499, gen12501)

				return

			} else {
				gen12496 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12496)

				return

			}

		}, 1)

		gen12509 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

		gen12510 := Call(__e, gen12509, V1644)

		__e.TailApply(gen12508, gen12510)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5st__input2_6, gen12511)

	gen12544 := MakeNative(func(__e *ControlFlow) {
		V1646 := __e.Get(1)
		_ = V1646
		gen12529 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12525 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12526 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12527 := Call(__e, gen12526)

			gen12528 := Call(__e, gen12525, YaccParse, gen12527)

			if True == gen12528 {
				gen12522 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5multiline_6 := __e.Get(1)
					_ = Parse__shen_4_5multiline_6
					gen12516 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12517 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12518 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12519 := Call(__e, gen12518)

					gen12520 := Call(__e, gen12517, gen12519, Parse__shen_4_5multiline_6)

					gen12521 := Call(__e, gen12516, gen12520)

					if True == gen12521 {
						gen12513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12514 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12515 := Call(__e, gen12514, Parse__shen_4_5multiline_6)

						__e.TailApply(gen12513, gen12515, symshen_4skip)

						return

					} else {
						gen12512 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12512)

						return

					}

				}, 1)

				gen12523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5multiline_6)

				gen12524 := Call(__e, gen12523, V1646)

				__e.TailApply(gen12522, gen12524)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12540 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5singleline_6 := __e.Get(1)
			_ = Parse__shen_4_5singleline_6
			gen12534 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12535 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12536 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12537 := Call(__e, gen12536)

			gen12538 := Call(__e, gen12535, gen12537, Parse__shen_4_5singleline_6)

			gen12539 := Call(__e, gen12534, gen12538)

			if True == gen12539 {
				gen12531 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12532 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12533 := Call(__e, gen12532, Parse__shen_4_5singleline_6)

				__e.TailApply(gen12531, gen12533, symshen_4skip)

				return

			} else {
				gen12530 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12530)

				return

			}

		}, 1)

		gen12541 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5singleline_6)

		gen12542 := Call(__e, gen12541, V1646)

		gen12543 := Call(__e, gen12540, gen12542)

		__e.TailApply(gen12529, gen12543)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comment_6, gen12544)

	gen12588 := MakeNative(func(__e *ControlFlow) {
		V1648 := __e.Get(1)
		_ = V1648
		gen12585 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5backslash_6 := __e.Get(1)
			_ = Parse__shen_4_5backslash_6
			gen12579 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12580 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12581 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12582 := Call(__e, gen12581)

			gen12583 := Call(__e, gen12580, gen12582, Parse__shen_4_5backslash_6)

			gen12584 := Call(__e, gen12579, gen12583)

			if True == gen12584 {
				gen12576 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5backslash_6 := __e.Get(1)
					_ = Parse__shen_4_5backslash_6
					gen12570 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12571 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12572 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12573 := Call(__e, gen12572)

					gen12574 := Call(__e, gen12571, gen12573, Parse__shen_4_5backslash_6)

					gen12575 := Call(__e, gen12570, gen12574)

					if True == gen12575 {
						gen12567 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5anysingle_6 := __e.Get(1)
							_ = Parse__shen_4_5anysingle_6
							gen12561 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen12562 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12563 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen12564 := Call(__e, gen12563)

							gen12565 := Call(__e, gen12562, gen12564, Parse__shen_4_5anysingle_6)

							gen12566 := Call(__e, gen12561, gen12565)

							if True == gen12566 {
								gen12558 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5return_6 := __e.Get(1)
									_ = Parse__shen_4_5return_6
									gen12552 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen12553 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen12554 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen12555 := Call(__e, gen12554)

									gen12556 := Call(__e, gen12553, gen12555, Parse__shen_4_5return_6)

									gen12557 := Call(__e, gen12552, gen12556)

									if True == gen12557 {
										gen12549 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen12550 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen12551 := Call(__e, gen12550, Parse__shen_4_5return_6)

										__e.TailApply(gen12549, gen12551, symshen_4skip)

										return

									} else {
										gen12548 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen12548)

										return

									}

								}, 1)

								gen12559 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5return_6)

								gen12560 := Call(__e, gen12559, Parse__shen_4_5anysingle_6)

								__e.TailApply(gen12558, gen12560)

								return

							} else {
								gen12547 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen12547)

								return

							}

						}, 1)

						gen12568 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anysingle_6)

						gen12569 := Call(__e, gen12568, Parse__shen_4_5backslash_6)

						__e.TailApply(gen12567, gen12569)

						return

					} else {
						gen12546 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12546)

						return

					}

				}, 1)

				gen12577 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

				gen12578 := Call(__e, gen12577, Parse__shen_4_5backslash_6)

				__e.TailApply(gen12576, gen12578)

				return

			} else {
				gen12545 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12545)

				return

			}

		}, 1)

		gen12586 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

		gen12587 := Call(__e, gen12586, V1648)

		__e.TailApply(gen12585, gen12587)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5singleline_6, gen12588)

	gen12609 := MakeNative(func(__e *ControlFlow) {
		V1651 := __e.Get(1)
		_ = V1651
		gen12604 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12605 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12606 := Call(__e, gen12605, V1651)

		gen12607 := Call(__e, gen12604, gen12606)

		var gen12608 Obj

		if True == gen12607 {
			gen12600 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12601 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12602 := Call(__e, gen12601, V1651)

			gen12603 := Call(__e, gen12600, MakeNumber(92), gen12602)

			if True == gen12603 {
				gen12608 = True
			} else {
				gen12608 = False
			}

		} else {
			gen12608 = False
		}

		if True == gen12608 {
			gen12593 := MakeNative(func(__e *ControlFlow) {
				NewStream1649 := __e.Get(1)
				_ = NewStream1649
				gen12590 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12591 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12592 := Call(__e, gen12591, NewStream1649)

				__e.TailApply(gen12590, gen12592, symshen_4skip)

				return

			}, 1)

			gen12594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen12595 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen12596 := Call(__e, gen12595, V1651)

			gen12597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen12598 := Call(__e, gen12597, V1651)

			gen12599 := Call(__e, gen12594, gen12596, gen12598)

			__e.TailApply(gen12593, gen12599)

			return

		} else {
			gen12589 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12589)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5backslash_6, gen12609)

	gen12652 := MakeNative(func(__e *ControlFlow) {
		V1653 := __e.Get(1)
		_ = V1653
		gen12627 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12623 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12624 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12625 := Call(__e, gen12624)

			gen12626 := Call(__e, gen12623, YaccParse, gen12625)

			if True == gen12626 {
				gen12620 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen12614 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12615 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12616 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12617 := Call(__e, gen12616)

					gen12618 := Call(__e, gen12615, gen12617, Parse___5e_6)

					gen12619 := Call(__e, gen12614, gen12618)

					if True == gen12619 {
						gen12611 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12612 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12613 := Call(__e, gen12612, Parse___5e_6)

						__e.TailApply(gen12611, gen12613, symshen_4skip)

						return

					} else {
						gen12610 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12610)

						return

					}

				}, 1)

				gen12621 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen12622 := Call(__e, gen12621, V1653)

				__e.TailApply(gen12620, gen12622)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12648 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5non_1return_6 := __e.Get(1)
			_ = Parse__shen_4_5non_1return_6
			gen12642 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12643 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12644 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12645 := Call(__e, gen12644)

			gen12646 := Call(__e, gen12643, gen12645, Parse__shen_4_5non_1return_6)

			gen12647 := Call(__e, gen12642, gen12646)

			if True == gen12647 {
				gen12639 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5anysingle_6 := __e.Get(1)
					_ = Parse__shen_4_5anysingle_6
					gen12633 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12634 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12635 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12636 := Call(__e, gen12635)

					gen12637 := Call(__e, gen12634, gen12636, Parse__shen_4_5anysingle_6)

					gen12638 := Call(__e, gen12633, gen12637)

					if True == gen12638 {
						gen12630 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12631 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12632 := Call(__e, gen12631, Parse__shen_4_5anysingle_6)

						__e.TailApply(gen12630, gen12632, symshen_4skip)

						return

					} else {
						gen12629 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12629)

						return

					}

				}, 1)

				gen12640 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anysingle_6)

				gen12641 := Call(__e, gen12640, Parse__shen_4_5non_1return_6)

				__e.TailApply(gen12639, gen12641)

				return

			} else {
				gen12628 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12628)

				return

			}

		}, 1)

		gen12649 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5non_1return_6)

		gen12650 := Call(__e, gen12649, V1653)

		gen12651 := Call(__e, gen12648, gen12650)

		__e.TailApply(gen12627, gen12651)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5anysingle_6, gen12652)

	gen12679 := MakeNative(func(__e *ControlFlow) {
		V1655 := __e.Get(1)
		_ = V1655
		gen12675 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12676 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12677 := Call(__e, gen12676, V1655)

		gen12678 := Call(__e, gen12675, gen12677)

		if True == gen12678 {
			gen12672 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen12664 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				gen12665 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen12666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen12667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen12668 := Call(__e, gen12667, MakeNumber(13), Nil)

				gen12669 := Call(__e, gen12666, MakeNumber(10), gen12668)

				gen12670 := Call(__e, gen12665, Parse__X, gen12669)

				gen12671 := Call(__e, gen12664, gen12670)

				if True == gen12671 {
					gen12655 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12656 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12658 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen12659 := Call(__e, gen12658, V1655)

					gen12660 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen12661 := Call(__e, gen12660, V1655)

					gen12662 := Call(__e, gen12657, gen12659, gen12661)

					gen12663 := Call(__e, gen12656, gen12662)

					__e.TailApply(gen12655, gen12663, symshen_4skip)

					return

				} else {
					gen12654 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen12654)

					return

				}

			}, 1)

			gen12673 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12674 := Call(__e, gen12673, V1655)

			__e.TailApply(gen12672, gen12674)

			return

		} else {
			gen12653 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12653)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5non_1return_6, gen12679)

	gen12704 := MakeNative(func(__e *ControlFlow) {
		V1657 := __e.Get(1)
		_ = V1657
		gen12700 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12701 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12702 := Call(__e, gen12701, V1657)

		gen12703 := Call(__e, gen12700, gen12702)

		if True == gen12703 {
			gen12697 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen12691 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen12692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen12693 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen12694 := Call(__e, gen12693, MakeNumber(13), Nil)

				gen12695 := Call(__e, gen12692, MakeNumber(10), gen12694)

				gen12696 := Call(__e, gen12691, Parse__X, gen12695)

				if True == gen12696 {
					gen12682 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12683 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12684 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12685 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen12686 := Call(__e, gen12685, V1657)

					gen12687 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen12688 := Call(__e, gen12687, V1657)

					gen12689 := Call(__e, gen12684, gen12686, gen12688)

					gen12690 := Call(__e, gen12683, gen12689)

					__e.TailApply(gen12682, gen12690, symshen_4skip)

					return

				} else {
					gen12681 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen12681)

					return

				}

			}, 1)

			gen12698 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12699 := Call(__e, gen12698, V1657)

			__e.TailApply(gen12697, gen12699)

			return

		} else {
			gen12680 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12680)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5return_6, gen12704)

	gen12738 := MakeNative(func(__e *ControlFlow) {
		V1659 := __e.Get(1)
		_ = V1659
		gen12735 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5backslash_6 := __e.Get(1)
			_ = Parse__shen_4_5backslash_6
			gen12729 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12730 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12731 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12732 := Call(__e, gen12731)

			gen12733 := Call(__e, gen12730, gen12732, Parse__shen_4_5backslash_6)

			gen12734 := Call(__e, gen12729, gen12733)

			if True == gen12734 {
				gen12726 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5times_6 := __e.Get(1)
					_ = Parse__shen_4_5times_6
					gen12720 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12721 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12722 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12723 := Call(__e, gen12722)

					gen12724 := Call(__e, gen12721, gen12723, Parse__shen_4_5times_6)

					gen12725 := Call(__e, gen12720, gen12724)

					if True == gen12725 {
						gen12717 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5anymulti_6 := __e.Get(1)
							_ = Parse__shen_4_5anymulti_6
							gen12711 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen12712 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12713 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen12714 := Call(__e, gen12713)

							gen12715 := Call(__e, gen12712, gen12714, Parse__shen_4_5anymulti_6)

							gen12716 := Call(__e, gen12711, gen12715)

							if True == gen12716 {
								gen12708 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen12709 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen12710 := Call(__e, gen12709, Parse__shen_4_5anymulti_6)

								__e.TailApply(gen12708, gen12710, symshen_4skip)

								return

							} else {
								gen12707 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen12707)

								return

							}

						}, 1)

						gen12718 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anymulti_6)

						gen12719 := Call(__e, gen12718, Parse__shen_4_5times_6)

						__e.TailApply(gen12717, gen12719)

						return

					} else {
						gen12706 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12706)

						return

					}

				}, 1)

				gen12727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5times_6)

				gen12728 := Call(__e, gen12727, Parse__shen_4_5backslash_6)

				__e.TailApply(gen12726, gen12728)

				return

			} else {
				gen12705 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12705)

				return

			}

		}, 1)

		gen12736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

		gen12737 := Call(__e, gen12736, V1659)

		__e.TailApply(gen12735, gen12737)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5multiline_6, gen12738)

	gen12759 := MakeNative(func(__e *ControlFlow) {
		V1662 := __e.Get(1)
		_ = V1662
		gen12754 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12755 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12756 := Call(__e, gen12755, V1662)

		gen12757 := Call(__e, gen12754, gen12756)

		var gen12758 Obj

		if True == gen12757 {
			gen12750 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12752 := Call(__e, gen12751, V1662)

			gen12753 := Call(__e, gen12750, MakeNumber(42), gen12752)

			if True == gen12753 {
				gen12758 = True
			} else {
				gen12758 = False
			}

		} else {
			gen12758 = False
		}

		if True == gen12758 {
			gen12743 := MakeNative(func(__e *ControlFlow) {
				NewStream1660 := __e.Get(1)
				_ = NewStream1660
				gen12740 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen12741 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12742 := Call(__e, gen12741, NewStream1660)

				__e.TailApply(gen12740, gen12742, symshen_4skip)

				return

			}, 1)

			gen12744 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen12745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen12746 := Call(__e, gen12745, V1662)

			gen12747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen12748 := Call(__e, gen12747, V1662)

			gen12749 := Call(__e, gen12744, gen12746, gen12748)

			__e.TailApply(gen12743, gen12749)

			return

		} else {
			gen12739 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12739)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5times_6, gen12759)

	gen12845 := MakeNative(func(__e *ControlFlow) {
		V1664 := __e.Get(1)
		_ = V1664
		gen12820 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12816 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12817 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12818 := Call(__e, gen12817)

			gen12819 := Call(__e, gen12816, YaccParse, gen12818)

			if True == gen12819 {
				gen12791 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					gen12787 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12788 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12789 := Call(__e, gen12788)

					gen12790 := Call(__e, gen12787, YaccParse, gen12789)

					if True == gen12790 {
						gen12783 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen12784 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12785 := Call(__e, gen12784, V1664)

						gen12786 := Call(__e, gen12783, gen12785)

						if True == gen12786 {
							gen12780 := MakeNative(func(__e *ControlFlow) {
								Parse__X := __e.Get(1)
								_ = Parse__X
								gen12771 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5anymulti_6 := __e.Get(1)
									_ = Parse__shen_4_5anymulti_6
									gen12765 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen12766 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen12767 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen12768 := Call(__e, gen12767)

									gen12769 := Call(__e, gen12766, gen12768, Parse__shen_4_5anymulti_6)

									gen12770 := Call(__e, gen12765, gen12769)

									if True == gen12770 {
										gen12762 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen12763 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen12764 := Call(__e, gen12763, Parse__shen_4_5anymulti_6)

										__e.TailApply(gen12762, gen12764, symshen_4skip)

										return

									} else {
										gen12761 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen12761)

										return

									}

								}, 1)

								gen12772 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anymulti_6)

								gen12773 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen12774 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

								gen12775 := Call(__e, gen12774, V1664)

								gen12776 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen12777 := Call(__e, gen12776, V1664)

								gen12778 := Call(__e, gen12773, gen12775, gen12777)

								gen12779 := Call(__e, gen12772, gen12778)

								__e.TailApply(gen12771, gen12779)

								return

							}, 1)

							gen12781 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							gen12782 := Call(__e, gen12781, V1664)

							__e.TailApply(gen12780, gen12782)

							return

						} else {
							gen12760 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(gen12760)

							return

						}

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				gen12812 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5times_6 := __e.Get(1)
					_ = Parse__shen_4_5times_6
					gen12806 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12807 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12808 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12809 := Call(__e, gen12808)

					gen12810 := Call(__e, gen12807, gen12809, Parse__shen_4_5times_6)

					gen12811 := Call(__e, gen12806, gen12810)

					if True == gen12811 {
						gen12803 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5backslash_6 := __e.Get(1)
							_ = Parse__shen_4_5backslash_6
							gen12797 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen12798 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12799 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen12800 := Call(__e, gen12799)

							gen12801 := Call(__e, gen12798, gen12800, Parse__shen_4_5backslash_6)

							gen12802 := Call(__e, gen12797, gen12801)

							if True == gen12802 {
								gen12794 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen12795 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen12796 := Call(__e, gen12795, Parse__shen_4_5backslash_6)

								__e.TailApply(gen12794, gen12796, symshen_4skip)

								return

							} else {
								gen12793 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen12793)

								return

							}

						}, 1)

						gen12804 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5backslash_6)

						gen12805 := Call(__e, gen12804, Parse__shen_4_5times_6)

						__e.TailApply(gen12803, gen12805)

						return

					} else {
						gen12792 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12792)

						return

					}

				}, 1)

				gen12813 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5times_6)

				gen12814 := Call(__e, gen12813, V1664)

				gen12815 := Call(__e, gen12812, gen12814)

				__e.TailApply(gen12791, gen12815)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12841 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5comment_6 := __e.Get(1)
			_ = Parse__shen_4_5comment_6
			gen12835 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12836 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12837 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12838 := Call(__e, gen12837)

			gen12839 := Call(__e, gen12836, gen12838, Parse__shen_4_5comment_6)

			gen12840 := Call(__e, gen12835, gen12839)

			if True == gen12840 {
				gen12832 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5anymulti_6 := __e.Get(1)
					_ = Parse__shen_4_5anymulti_6
					gen12826 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12827 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12828 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12829 := Call(__e, gen12828)

					gen12830 := Call(__e, gen12827, gen12829, Parse__shen_4_5anymulti_6)

					gen12831 := Call(__e, gen12826, gen12830)

					if True == gen12831 {
						gen12823 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12824 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12825 := Call(__e, gen12824, Parse__shen_4_5anymulti_6)

						__e.TailApply(gen12823, gen12825, symshen_4skip)

						return

					} else {
						gen12822 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12822)

						return

					}

				}, 1)

				gen12833 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5anymulti_6)

				gen12834 := Call(__e, gen12833, Parse__shen_4_5comment_6)

				__e.TailApply(gen12832, gen12834)

				return

			} else {
				gen12821 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12821)

				return

			}

		}, 1)

		gen12842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comment_6)

		gen12843 := Call(__e, gen12842, V1664)

		gen12844 := Call(__e, gen12841, gen12843)

		__e.TailApply(gen12820, gen12844)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5anymulti_6, gen12845)

	gen12888 := MakeNative(func(__e *ControlFlow) {
		V1666 := __e.Get(1)
		_ = V1666
		gen12863 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen12859 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12860 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12861 := Call(__e, gen12860)

			gen12862 := Call(__e, gen12859, YaccParse, gen12861)

			if True == gen12862 {
				gen12856 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5whitespace_6 := __e.Get(1)
					_ = Parse__shen_4_5whitespace_6
					gen12850 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12851 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12852 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12853 := Call(__e, gen12852)

					gen12854 := Call(__e, gen12851, gen12853, Parse__shen_4_5whitespace_6)

					gen12855 := Call(__e, gen12850, gen12854)

					if True == gen12855 {
						gen12847 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12848 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12849 := Call(__e, gen12848, Parse__shen_4_5whitespace_6)

						__e.TailApply(gen12847, gen12849, symshen_4skip)

						return

					} else {
						gen12846 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12846)

						return

					}

				}, 1)

				gen12857 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespace_6)

				gen12858 := Call(__e, gen12857, V1666)

				__e.TailApply(gen12856, gen12858)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen12884 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5whitespace_6 := __e.Get(1)
			_ = Parse__shen_4_5whitespace_6
			gen12878 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen12879 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen12880 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen12881 := Call(__e, gen12880)

			gen12882 := Call(__e, gen12879, gen12881, Parse__shen_4_5whitespace_6)

			gen12883 := Call(__e, gen12878, gen12882)

			if True == gen12883 {
				gen12875 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5whitespaces_6 := __e.Get(1)
					_ = Parse__shen_4_5whitespaces_6
					gen12869 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen12870 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12871 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen12872 := Call(__e, gen12871)

					gen12873 := Call(__e, gen12870, gen12872, Parse__shen_4_5whitespaces_6)

					gen12874 := Call(__e, gen12869, gen12873)

					if True == gen12874 {
						gen12866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen12867 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen12868 := Call(__e, gen12867, Parse__shen_4_5whitespaces_6)

						__e.TailApply(gen12866, gen12868, symshen_4skip)

						return

					} else {
						gen12865 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen12865)

						return

					}

				}, 1)

				gen12876 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespaces_6)

				gen12877 := Call(__e, gen12876, Parse__shen_4_5whitespace_6)

				__e.TailApply(gen12875, gen12877)

				return

			} else {
				gen12864 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen12864)

				return

			}

		}, 1)

		gen12885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5whitespace_6)

		gen12886 := Call(__e, gen12885, V1666)

		gen12887 := Call(__e, gen12884, gen12886)

		__e.TailApply(gen12863, gen12887)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5whitespaces_6, gen12888)

	gen12919 := MakeNative(func(__e *ControlFlow) {
		V1668 := __e.Get(1)
		_ = V1668
		gen12915 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen12916 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen12917 := Call(__e, gen12916, V1668)

		gen12918 := Call(__e, gen12915, gen12917)

		if True == gen12918 {
			gen12912 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen12910 := MakeNative(func(__e *ControlFlow) {
					Parse__Case := __e.Get(1)
					_ = Parse__Case
					gen12908 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen12909 := Call(__e, gen12908, Parse__Case, MakeNumber(32))

					if True == gen12909 {
						__e.Return(True)
						return
					} else {
						gen12905 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen12906 := Call(__e, gen12905, Parse__Case, MakeNumber(13))

						var gen12907 Obj

						if True == gen12906 {
							gen12907 = True
						} else {
							gen12902 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12903 := Call(__e, gen12902, Parse__Case, MakeNumber(10))

							var gen12904 Obj

							if True == gen12903 {
								gen12904 = True
							} else {
								gen12900 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen12901 := Call(__e, gen12900, Parse__Case, MakeNumber(9))

								if True == gen12901 {
									gen12904 = True
								} else {
									gen12904 = False
								}

							}

							if True == gen12904 {
								gen12907 = True
							} else {
								gen12907 = False
							}

						}

						if True == gen12907 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				gen12911 := Call(__e, gen12910, Parse__X)

				if True == gen12911 {
					gen12891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12892 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12893 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen12894 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen12895 := Call(__e, gen12894, V1668)

					gen12896 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen12897 := Call(__e, gen12896, V1668)

					gen12898 := Call(__e, gen12893, gen12895, gen12897)

					gen12899 := Call(__e, gen12892, gen12898)

					__e.TailApply(gen12891, gen12899, symshen_4skip)

					return

				} else {
					gen12890 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen12890)

					return

				}

			}, 1)

			gen12913 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen12914 := Call(__e, gen12913, V1668)

			__e.TailApply(gen12912, gen12914)

			return

		} else {
			gen12889 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen12889)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5whitespace_6, gen12919)

	gen12976 := MakeNative(func(__e *ControlFlow) {
		V1670 := __e.Get(1)
		_ = V1670
		gen12974 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen12975 := Call(__e, gen12974, Nil, V1670)

		if True == gen12975 {
			__e.Return(Nil)
			return
		} else {
			gen12971 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen12972 := Call(__e, gen12971, V1670)

			var gen12973 Obj

			if True == gen12972 {
				gen12966 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen12967 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen12968 := Call(__e, gen12967, V1670)

				gen12969 := Call(__e, gen12966, gen12968)

				var gen12970 Obj

				if True == gen12969 {
					gen12959 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen12960 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen12961 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen12962 := Call(__e, gen12961, V1670)

					gen12963 := Call(__e, gen12960, gen12962)

					gen12964 := Call(__e, gen12959, gen12963)

					var gen12965 Obj

					if True == gen12964 {
						gen12950 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen12951 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen12952 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen12953 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen12954 := Call(__e, gen12953, V1670)

						gen12955 := Call(__e, gen12952, gen12954)

						gen12956 := Call(__e, gen12951, gen12955)

						gen12957 := Call(__e, gen12950, Nil, gen12956)

						var gen12958 Obj

						if True == gen12957 {
							gen12944 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen12945 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen12946 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen12947 := Call(__e, gen12946, V1670)

							gen12948 := Call(__e, gen12945, gen12947)

							gen12949 := Call(__e, gen12944, gen12948, symbar_b)

							if True == gen12949 {
								gen12958 = True
							} else {
								gen12958 = False
							}

						} else {
							gen12958 = False
						}

						if True == gen12958 {
							gen12965 = True
						} else {
							gen12965 = False
						}

					} else {
						gen12965 = False
					}

					if True == gen12965 {
						gen12970 = True
					} else {
						gen12970 = False
					}

				} else {
					gen12970 = False
				}

				if True == gen12970 {
					gen12973 = True
				} else {
					gen12973 = False
				}

			} else {
				gen12973 = False
			}

			if True == gen12973 {
				gen12935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen12936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen12937 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen12938 := Call(__e, gen12937, V1670)

				gen12939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen12940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen12941 := Call(__e, gen12940, V1670)

				gen12942 := Call(__e, gen12939, gen12941)

				gen12943 := Call(__e, gen12936, gen12938, gen12942)

				__e.TailApply(gen12935, symcons, gen12943)

				return

			} else {
				gen12933 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen12934 := Call(__e, gen12933, V1670)

				if True == gen12934 {
					gen12922 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen12923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen12924 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen12925 := Call(__e, gen12924, V1670)

					gen12926 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen12927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cons__form)

					gen12928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen12929 := Call(__e, gen12928, V1670)

					gen12930 := Call(__e, gen12927, gen12929)

					gen12931 := Call(__e, gen12926, gen12930, Nil)

					gen12932 := Call(__e, gen12923, gen12925, gen12931)

					__e.TailApply(gen12922, symcons, gen12932)

					return

				} else {
					if True == True {
						gen12921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen12921, symshen_4cons__form)

						return

					} else {
						gen12920 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen12920, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4cons__form, gen12976)

	gen13106 := MakeNative(func(__e *ControlFlow) {
		V1675 := __e.Get(1)
		_ = V1675
		V1676 := __e.Get(2)
		_ = V1676
		gen13103 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen13104 := Call(__e, gen13103, V1675)

		var gen13105 Obj

		if True == gen13104 {
			gen13098 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen13099 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen13100 := Call(__e, gen13099, V1675)

			gen13101 := Call(__e, gen13098, sym_3, gen13100)

			var gen13102 Obj

			if True == gen13101 {
				gen13093 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen13094 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen13095 := Call(__e, gen13094, V1675)

				gen13096 := Call(__e, gen13093, gen13095)

				var gen13097 Obj

				if True == gen13096 {
					gen13087 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen13088 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen13089 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen13090 := Call(__e, gen13089, V1675)

					gen13091 := Call(__e, gen13088, gen13090)

					gen13092 := Call(__e, gen13087, Nil, gen13091)

					if True == gen13092 {
						gen13097 = True
					} else {
						gen13097 = False
					}

				} else {
					gen13097 = False
				}

				if True == gen13097 {
					gen13102 = True
				} else {
					gen13102 = False
				}

			} else {
				gen13102 = False
			}

			if True == gen13102 {
				gen13105 = True
			} else {
				gen13105 = False
			}

		} else {
			gen13105 = False
		}

		if True == gen13105 {
			gen13080 := Call(__e, PrimNS1Value(symns2_1value), symappend)

			gen13081 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

			gen13082 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen13083 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen13084 := Call(__e, gen13083, V1675)

			gen13085 := Call(__e, gen13082, gen13084)

			gen13086 := Call(__e, gen13081, gen13085)

			__e.TailApply(gen13080, gen13086, V1676)

			return

		} else {
			gen13077 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen13078 := Call(__e, gen13077, V1675)

			var gen13079 Obj

			if True == gen13078 {
				gen13072 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen13073 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen13074 := Call(__e, gen13073, V1675)

				gen13075 := Call(__e, gen13072, sympackage, gen13074)

				var gen13076 Obj

				if True == gen13075 {
					gen13067 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen13068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen13069 := Call(__e, gen13068, V1675)

					gen13070 := Call(__e, gen13067, gen13069)

					var gen13071 Obj

					if True == gen13070 {
						gen13060 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen13061 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen13062 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen13063 := Call(__e, gen13062, V1675)

						gen13064 := Call(__e, gen13061, gen13063)

						gen13065 := Call(__e, gen13060, symnull, gen13064)

						var gen13066 Obj

						if True == gen13065 {
							gen13054 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen13055 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen13056 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen13057 := Call(__e, gen13056, V1675)

							gen13058 := Call(__e, gen13055, gen13057)

							gen13059 := Call(__e, gen13054, gen13058)

							if True == gen13059 {
								gen13066 = True
							} else {
								gen13066 = False
							}

						} else {
							gen13066 = False
						}

						if True == gen13066 {
							gen13071 = True
						} else {
							gen13071 = False
						}

					} else {
						gen13071 = False
					}

					if True == gen13071 {
						gen13076 = True
					} else {
						gen13076 = False
					}

				} else {
					gen13076 = False
				}

				if True == gen13076 {
					gen13079 = True
				} else {
					gen13079 = False
				}

			} else {
				gen13079 = False
			}

			if True == gen13079 {
				gen13047 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				gen13048 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen13049 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen13050 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen13051 := Call(__e, gen13050, V1675)

				gen13052 := Call(__e, gen13049, gen13051)

				gen13053 := Call(__e, gen13048, gen13052)

				__e.TailApply(gen13047, gen13053, V1676)

				return

			} else {
				gen13044 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen13045 := Call(__e, gen13044, V1675)

				var gen13046 Obj

				if True == gen13045 {
					gen13039 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen13040 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen13041 := Call(__e, gen13040, V1675)

					gen13042 := Call(__e, gen13039, sympackage, gen13041)

					var gen13043 Obj

					if True == gen13042 {
						gen13034 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen13035 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen13036 := Call(__e, gen13035, V1675)

						gen13037 := Call(__e, gen13034, gen13036)

						var gen13038 Obj

						if True == gen13037 {
							gen13028 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen13029 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen13030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen13031 := Call(__e, gen13030, V1675)

							gen13032 := Call(__e, gen13029, gen13031)

							gen13033 := Call(__e, gen13028, gen13032)

							if True == gen13033 {
								gen13038 = True
							} else {
								gen13038 = False
							}

						} else {
							gen13038 = False
						}

						if True == gen13038 {
							gen13043 = True
						} else {
							gen13043 = False
						}

					} else {
						gen13043 = False
					}

					if True == gen13043 {
						gen13046 = True
					} else {
						gen13046 = False
					}

				} else {
					gen13046 = False
				}

				if True == gen13046 {
					gen13019 := MakeNative(func(__e *ControlFlow) {
						ListofExceptions := __e.Get(1)
						_ = ListofExceptions
						gen13012 := MakeNative(func(__e *ControlFlow) {
							External := __e.Get(1)
							_ = External
							gen13001 := MakeNative(func(__e *ControlFlow) {
								PackageNameDot := __e.Get(1)
								_ = PackageNameDot
								gen12998 := MakeNative(func(__e *ControlFlow) {
									ExpPackageNameDot := __e.Get(1)
									_ = ExpPackageNameDot
									gen12989 := MakeNative(func(__e *ControlFlow) {
										Packaged := __e.Get(1)
										_ = Packaged
										gen12980 := MakeNative(func(__e *ControlFlow) {
											Internal := __e.Get(1)
											_ = Internal
											gen12979 := Call(__e, PrimNS1Value(symns2_1value), symappend)

											__e.TailApply(gen12979, Packaged, V1676)

											return

										}, 1)

										gen12981 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1internal)

										gen12982 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen12983 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen12984 := Call(__e, gen12983, V1675)

										gen12985 := Call(__e, gen12982, gen12984)

										gen12986 := Call(__e, PrimNS1Value(symns2_1value), symshen_4internal_1symbols)

										gen12987 := Call(__e, gen12986, ExpPackageNameDot, Packaged)

										gen12988 := Call(__e, gen12981, gen12985, gen12987)

										__e.TailApply(gen12980, gen12988)

										return

									}, 1)

									gen12990 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

									gen12991 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen12992 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen12993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen12994 := Call(__e, gen12993, V1675)

									gen12995 := Call(__e, gen12992, gen12994)

									gen12996 := Call(__e, gen12991, gen12995)

									gen12997 := Call(__e, gen12990, PackageNameDot, ListofExceptions, gen12996, ExpPackageNameDot)

									__e.TailApply(gen12989, gen12997)

									return

								}, 1)

								gen12999 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

								gen13000 := Call(__e, gen12999, PackageNameDot)

								__e.TailApply(gen12998, gen13000)

								return

							}, 1)

							gen13002 := Call(__e, PrimNS1Value(symns2_1value), symintern)

							gen13003 := Call(__e, PrimNS1Value(symns2_1value), symcn)

							gen13004 := Call(__e, PrimNS1Value(symns2_1value), symstr)

							gen13005 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen13006 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen13007 := Call(__e, gen13006, V1675)

							gen13008 := Call(__e, gen13005, gen13007)

							gen13009 := Call(__e, gen13004, gen13008)

							gen13010 := Call(__e, gen13003, gen13009, MakeString("."))

							gen13011 := Call(__e, gen13002, gen13010)

							__e.TailApply(gen13001, gen13011)

							return

						}, 1)

						gen13013 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1exceptions)

						gen13014 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen13015 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen13016 := Call(__e, gen13015, V1675)

						gen13017 := Call(__e, gen13014, gen13016)

						gen13018 := Call(__e, gen13013, ListofExceptions, gen13017)

						__e.TailApply(gen13012, gen13018)

						return

					}, 1)

					gen13020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

					gen13021 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen13022 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen13023 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen13024 := Call(__e, gen13023, V1675)

					gen13025 := Call(__e, gen13022, gen13024)

					gen13026 := Call(__e, gen13021, gen13025)

					gen13027 := Call(__e, gen13020, gen13026)

					__e.TailApply(gen13019, gen13027)

					return

				} else {
					if True == True {
						gen12978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						__e.TailApply(gen12978, V1675, V1676)

						return

					} else {
						gen12977 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen12977, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4package_1macro, gen13106)

	gen13120 := MakeNative(func(__e *ControlFlow) {
		V1679 := __e.Get(1)
		_ = V1679
		V1680 := __e.Get(2)
		_ = V1680
		gen13113 := MakeNative(func(__e *ControlFlow) {
			CurrExceptions := __e.Get(1)
			_ = CurrExceptions
			gen13110 := MakeNative(func(__e *ControlFlow) {
				AllExceptions := __e.Get(1)
				_ = AllExceptions
				gen13107 := Call(__e, PrimNS1Value(symns2_1value), symput)

				gen13108 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				gen13109 := Call(__e, gen13108, sym_dproperty_1vector_d)

				__e.TailApply(gen13107, V1680, symshen_4external_1symbols, AllExceptions, gen13109)

				return

			}, 1)

			gen13111 := Call(__e, PrimNS1Value(symns2_1value), symunion)

			gen13112 := Call(__e, gen13111, V1679, CurrExceptions)

			__e.TailApply(gen13110, gen13112)

			return

		}, 1)

		gen13117 := MakeNative(func(__e *ControlFlow) {
			gen13114 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen13115 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen13116 := Call(__e, gen13115, sym_dproperty_1vector_d)

			__e.TailApply(gen13114, V1680, symshen_4external_1symbols, gen13116)

			return

		}, 0)

		gen13118 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		gen13119 := Call(__e, PrimNS1Value(symtry_1catch), gen13117, gen13118)

		__e.TailApply(gen13113, gen13119)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1exceptions, gen13120)

	gen13132 := MakeNative(func(__e *ControlFlow) {
		V1683 := __e.Get(1)
		_ = V1683
		V1684 := __e.Get(2)
		_ = V1684
		gen13121 := Call(__e, PrimNS1Value(symns2_1value), symput)

		gen13122 := Call(__e, PrimNS1Value(symns2_1value), symunion)

		gen13126 := MakeNative(func(__e *ControlFlow) {
			gen13123 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen13124 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen13125 := Call(__e, gen13124, sym_dproperty_1vector_d)

			__e.TailApply(gen13123, V1683, symshen_4internal_1symbols, gen13125)

			return

		}, 0)

		gen13127 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		gen13128 := Call(__e, PrimNS1Value(symtry_1catch), gen13126, gen13127)

		gen13129 := Call(__e, gen13122, V1684, gen13128)

		gen13130 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen13131 := Call(__e, gen13130, sym_dproperty_1vector_d)

		__e.TailApply(gen13121, V1683, symshen_4internal_1symbols, gen13129, gen13131)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1internal, gen13132)

	gen13153 := MakeNative(func(__e *ControlFlow) {
		V1695 := __e.Get(1)
		_ = V1695
		V1696 := __e.Get(2)
		_ = V1696
		gen13150 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		gen13151 := Call(__e, gen13150, V1696)

		var gen13152 Obj

		if True == gen13151 {
			gen13146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

			gen13147 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

			gen13148 := Call(__e, gen13147, V1696)

			gen13149 := Call(__e, gen13146, V1695, gen13148)

			if True == gen13149 {
				gen13152 = True
			} else {
				gen13152 = False
			}

		} else {
			gen13152 = False
		}

		if True == gen13152 {
			gen13145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(gen13145, V1696, Nil)

			return

		} else {
			gen13143 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen13144 := Call(__e, gen13143, V1696)

			if True == gen13144 {
				gen13134 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				gen13135 := Call(__e, PrimNS1Value(symns2_1value), symshen_4internal_1symbols)

				gen13136 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen13137 := Call(__e, gen13136, V1696)

				gen13138 := Call(__e, gen13135, V1695, gen13137)

				gen13139 := Call(__e, PrimNS1Value(symns2_1value), symshen_4internal_1symbols)

				gen13140 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen13141 := Call(__e, gen13140, V1696)

				gen13142 := Call(__e, gen13139, V1695, gen13141)

				__e.TailApply(gen13134, gen13138, gen13142)

				return

			} else {
				if True == True {
					__e.Return(Nil)
					return
				} else {
					gen13133 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen13133, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4internal_1symbols, gen13153)

	gen13206 := MakeNative(func(__e *ControlFlow) {
		V1713 := __e.Get(1)
		_ = V1713
		V1714 := __e.Get(2)
		_ = V1714
		V1715 := __e.Get(3)
		_ = V1715
		V1716 := __e.Get(4)
		_ = V1716
		gen13204 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen13205 := Call(__e, gen13204, V1715)

		if True == gen13205 {
			gen13195 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen13196 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

			gen13197 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen13198 := Call(__e, gen13197, V1715)

			gen13199 := Call(__e, gen13196, V1713, V1714, gen13198, V1716)

			gen13200 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

			gen13201 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen13202 := Call(__e, gen13201, V1715)

			gen13203 := Call(__e, gen13200, V1713, V1714, gen13202, V1716)

			__e.TailApply(gen13195, gen13199, gen13203)

			return

		} else {
			gen13192 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sysfunc_2)

			gen13193 := Call(__e, gen13192, V1715)

			var gen13194 Obj

			if True == gen13193 {
				gen13194 = True
			} else {
				gen13189 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				gen13190 := Call(__e, gen13189, V1715)

				var gen13191 Obj

				if True == gen13190 {
					gen13191 = True
				} else {
					gen13186 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					gen13187 := Call(__e, gen13186, V1715, V1714)

					var gen13188 Obj

					if True == gen13187 {
						gen13188 = True
					} else {
						gen13183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4doubleunderline_2)

						gen13184 := Call(__e, gen13183, V1715)

						var gen13185 Obj

						if True == gen13184 {
							gen13185 = True
						} else {
							gen13181 := Call(__e, PrimNS1Value(symns2_1value), symshen_4singleunderline_2)

							gen13182 := Call(__e, gen13181, V1715)

							if True == gen13182 {
								gen13185 = True
							} else {
								gen13185 = False
							}

						}

						if True == gen13185 {
							gen13188 = True
						} else {
							gen13188 = False
						}

					}

					if True == gen13188 {
						gen13191 = True
					} else {
						gen13191 = False
					}

				}

				if True == gen13191 {
					gen13194 = True
				} else {
					gen13194 = False
				}

			}

			if True == gen13194 {
				__e.Return(V1715)
				return
			} else {
				gen13178 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

				gen13179 := Call(__e, gen13178, V1715)

				var gen13180 Obj

				if True == gen13179 {
					gen13174 := MakeNative(func(__e *ControlFlow) {
						ExplodeX := __e.Get(1)
						_ = ExplodeX
						gen13160 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						gen13161 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

						gen13162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen13163 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen13164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen13165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen13166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen13167 := Call(__e, gen13166, MakeString("."), Nil)

						gen13168 := Call(__e, gen13165, MakeString("n"), gen13167)

						gen13169 := Call(__e, gen13164, MakeString("e"), gen13168)

						gen13170 := Call(__e, gen13163, MakeString("h"), gen13169)

						gen13171 := Call(__e, gen13162, MakeString("s"), gen13170)

						gen13172 := Call(__e, gen13161, gen13171, ExplodeX)

						gen13173 := Call(__e, gen13160, gen13172)

						if True == gen13173 {
							gen13156 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen13157 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

							gen13158 := Call(__e, gen13157, V1716, ExplodeX)

							gen13159 := Call(__e, gen13156, gen13158)

							if True == gen13159 {
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

					gen13175 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

					gen13176 := Call(__e, gen13175, V1715)

					gen13177 := Call(__e, gen13174, gen13176)

					if True == gen13177 {
						gen13180 = True
					} else {
						gen13180 = False
					}

				} else {
					gen13180 = False
				}

				if True == gen13180 {
					gen13155 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					__e.TailApply(gen13155, V1713, V1715)

					return

				} else {
					if True == True {
						__e.Return(V1715)
						return
					} else {
						gen13154 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen13154, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 4)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4packageh, gen13206)

	return

}, 0)
