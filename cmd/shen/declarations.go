package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen22423 := MakeNative(func(__e *ControlFlow) {
		V451 := __e.Get(1)
		_ = V451
		gen22421 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22422 := Call(__e, gen22421, Nil, V451)

		if True == gen22422 {
			__e.Return(Nil)
			return
		} else {
			gen22418 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen22419 := Call(__e, gen22418, V451)

			var gen22420 Obj

			if True == gen22419 {
				gen22414 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen22415 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22416 := Call(__e, gen22415, V451)

				gen22417 := Call(__e, gen22414, gen22416)

				if True == gen22417 {
					gen22420 = True
				} else {
					gen22420 = False
				}

			} else {
				gen22420 = False
			}

			if True == gen22420 {
				gen22403 := MakeNative(func(__e *ControlFlow) {
					DecArity := __e.Get(1)
					_ = DecArity
					gen22398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4initialise__arity__table)

					gen22399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22400 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22401 := Call(__e, gen22400, V451)

					gen22402 := Call(__e, gen22399, gen22401)

					__e.TailApply(gen22398, gen22402)

					return

				}, 1)

				gen22404 := Call(__e, PrimNS1Value(symns2_1value), symput)

				gen22405 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22406 := Call(__e, gen22405, V451)

				gen22407 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22408 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22409 := Call(__e, gen22408, V451)

				gen22410 := Call(__e, gen22407, gen22409)

				gen22411 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				gen22412 := Call(__e, gen22411, sym_dproperty_1vector_d)

				gen22413 := Call(__e, gen22404, gen22406, symarity, gen22410, gen22412)

				__e.TailApply(gen22403, gen22413)

				return

			} else {
				if True == True {
					gen22397 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen22397, symshen_4initialise__arity__table)

					return

				} else {
					gen22396 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen22396, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise__arity__table, gen22423)

	gen22429 := MakeNative(func(__e *ControlFlow) {
		V453 := __e.Get(1)
		_ = V453
		gen22427 := MakeNative(func(__e *ControlFlow) {
			gen22424 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen22425 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen22426 := Call(__e, gen22425, sym_dproperty_1vector_d)

			__e.TailApply(gen22424, V453, symarity, gen22426)

			return

		}, 0)

		gen22428 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(-1))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen22427, gen22428)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symarity, gen22429)

	gen22445 := MakeNative(func(__e *ControlFlow) {
		V455 := __e.Get(1)
		_ = V455
		gen22442 := MakeNative(func(__e *ControlFlow) {
			Shen := __e.Get(1)
			_ = Shen
			gen22437 := MakeNative(func(__e *ControlFlow) {
				External := __e.Get(1)
				_ = External
				gen22430 := MakeNative(func(__e *ControlFlow) {
					Place := __e.Get(1)
					_ = Place
					__e.Return(V455)
					return
				}, 1)

				gen22431 := Call(__e, PrimNS1Value(symns2_1value), symput)

				gen22432 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

				gen22433 := Call(__e, gen22432, V455, External)

				gen22434 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				gen22435 := Call(__e, gen22434, sym_dproperty_1vector_d)

				gen22436 := Call(__e, gen22431, Shen, symshen_4external_1symbols, gen22433, gen22435)

				__e.TailApply(gen22430, gen22436)

				return

			}, 1)

			gen22438 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen22439 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen22440 := Call(__e, gen22439, sym_dproperty_1vector_d)

			gen22441 := Call(__e, gen22438, Shen, symshen_4external_1symbols, gen22440)

			__e.TailApply(gen22437, gen22441)

			return

		}, 1)

		gen22443 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		gen22444 := Call(__e, gen22443, MakeString("shen"))

		__e.TailApply(gen22442, gen22444)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symsystemf, gen22445)

	gen22449 := MakeNative(func(__e *ControlFlow) {
		V458 := __e.Get(1)
		_ = V458
		V459 := __e.Get(2)
		_ = V459
		gen22447 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen22448 := Call(__e, gen22447, V458, V459)

		if True == gen22448 {
			__e.Return(V459)
			return
		} else {
			gen22446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(gen22446, V458, V459)

			return

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symadjoin, gen22449)

	gen22469 := MakeNative(func(__e *ControlFlow) {
		V461 := __e.Get(1)
		_ = V461
		gen22467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22468 := Call(__e, gen22467, sympackage, V461)

		if True == gen22468 {
			__e.Return(Nil)
			return
		} else {
			gen22465 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen22466 := Call(__e, gen22465, symreceive, V461)

			if True == gen22466 {
				__e.Return(Nil)
				return
			} else {
				if True == True {
					gen22462 := MakeNative(func(__e *ControlFlow) {
						ArityF := __e.Get(1)
						_ = ArityF
						gen22460 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen22461 := Call(__e, gen22460, ArityF, MakeNumber(-1))

						if True == gen22461 {
							__e.Return(Nil)
							return
						} else {
							gen22458 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen22459 := Call(__e, gen22458, ArityF, MakeNumber(0))

							if True == gen22459 {
								__e.Return(Nil)
								return
							} else {
								gen22451 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen22452 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen22453 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

								gen22454 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1form)

								gen22455 := Call(__e, gen22454, V461, ArityF)

								gen22456 := Call(__e, gen22453, gen22455)

								gen22457 := Call(__e, gen22452, V461, gen22456)

								__e.TailApply(gen22451, gen22457, Nil)

								return

							}

						}

					}, 1)

					gen22463 := Call(__e, PrimNS1Value(symns2_1value), symarity)

					gen22464 := Call(__e, gen22463, V461)

					__e.TailApply(gen22462, gen22464)

					return

				} else {
					gen22450 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen22450, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1form_1entry, gen22469)

	gen22487 := MakeNative(func(__e *ControlFlow) {
		V464 := __e.Get(1)
		_ = V464
		V465 := __e.Get(2)
		_ = V465
		gen22485 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22486 := Call(__e, gen22485, MakeNumber(0), V465)

		if True == gen22486 {
			__e.Return(V464)
			return
		} else {
			if True == True {
				gen22482 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen22471 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22472 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22474 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1form)

					gen22475 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add_1end)

					gen22476 := Call(__e, gen22475, V464, X)

					gen22477 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					gen22478 := Call(__e, gen22477, V465, MakeNumber(1))

					gen22479 := Call(__e, gen22474, gen22476, gen22478)

					gen22480 := Call(__e, gen22473, gen22479, Nil)

					gen22481 := Call(__e, gen22472, X, gen22480)

					__e.TailApply(gen22471, symlambda, gen22481)

					return

				}, 1)

				gen22483 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				gen22484 := Call(__e, gen22483, symV)

				__e.TailApply(gen22482, gen22484)

				return

			} else {
				gen22470 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen22470, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1form, gen22487)

	gen22497 := MakeNative(func(__e *ControlFlow) {
		V468 := __e.Get(1)
		_ = V468
		V469 := __e.Get(2)
		_ = V469
		gen22495 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen22496 := Call(__e, gen22495, V468)

		if True == gen22496 {
			gen22492 := Call(__e, PrimNS1Value(symns2_1value), symappend)

			gen22493 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22494 := Call(__e, gen22493, V469, Nil)

			__e.TailApply(gen22492, V468, gen22494)

			return

		} else {
			if True == True {
				gen22489 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22490 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22491 := Call(__e, gen22490, V469, Nil)

				__e.TailApply(gen22489, V468, gen22491)

				return

			} else {
				gen22488 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen22488, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4add_1end, gen22497)

	gen22509 := MakeNative(func(__e *ControlFlow) {
		V471 := __e.Get(1)
		_ = V471
		gen22507 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen22508 := Call(__e, gen22507, V471)

		if True == gen22508 {
			gen22500 := Call(__e, PrimNS1Value(symns2_1value), symput)

			gen22501 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22502 := Call(__e, gen22501, V471)

			gen22503 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22504 := Call(__e, gen22503, V471)

			gen22505 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen22506 := Call(__e, gen22505, sym_dproperty_1vector_d)

			__e.TailApply(gen22500, gen22502, symshen_4lambda_1form, gen22504, gen22506)

			return

		} else {
			if True == True {
				gen22499 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen22499, symshen_4set_1lambda_1form_1entry)

				return

			} else {
				gen22498 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen22498, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4set_1lambda_1form_1entry, gen22509)

	gen22515 := MakeNative(func(__e *ControlFlow) {
		V473 := __e.Get(1)
		_ = V473
		gen22510 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen22511 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen22512 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen22513 := Call(__e, gen22512, symshen_4_dspecial_d)

		gen22514 := Call(__e, gen22511, V473, gen22513)

		Call(__e, gen22510, symshen_4_dspecial_d, gen22514)

		__e.Return(V473)
		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symspecialise, gen22515)

	gen22521 := MakeNative(func(__e *ControlFlow) {
		V475 := __e.Get(1)
		_ = V475
		gen22516 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen22517 := Call(__e, PrimNS1Value(symns2_1value), symremove)

		gen22518 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen22519 := Call(__e, gen22518, symshen_4_dspecial_d)

		gen22520 := Call(__e, gen22517, V475, gen22519)

		Call(__e, gen22516, symshen_4_dspecial_d, gen22520)

		__e.Return(V475)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symunspecialise, gen22521)

	return

}, 0)
