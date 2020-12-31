package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen10978 := MakeNative(func(__e Evaluator) {
			V451 := __e.Get(1)
			_ = V451
			gen10977 := Call(__e, __e.Global(sym_a), Nil, V451)

			if True == gen10977 {
				__e.Return(Nil)
				return
			} else {
				gen10975 := Call(__e, __e.Global(symcons_2), V451)

				var gen10976 Obj
				if True == gen10975 {
					gen10973 := Call(__e, __e.Global(symtl), V451)

					gen10974 := Call(__e, __e.Global(symcons_2), gen10973)

					if True == gen10974 {
						gen10976 = True
					} else {
						gen10976 = False
					}

				} else {
					gen10976 = False
				}
				if True == gen10976 {
					gen10966 := Call(__e, __e.Global(symhd), V451)

					gen10967 := Call(__e, __e.Global(symtl), V451)

					gen10968 := Call(__e, __e.Global(symhd), gen10967)

					gen10969 := Call(__e, __e.Global(symvalue), MakeSymbol("*property-vector*"))

					gen10970 := Call(__e, __e.Global(symput), gen10966, MakeSymbol("arity"), gen10968, gen10969)

					DecArity := gen10970
					_ = DecArity
					gen10971 := Call(__e, __e.Global(symtl), V451)

					gen10972 := Call(__e, __e.Global(symtl), gen10971)

					__e.TailApply(__e.Global(symshen_4initialise__arity__table), gen10972)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.initialise_arity_table"))

					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.initialise_arity_table"), gen10978)

		gen10982 := MakeNative(func(__e Evaluator) {
			V453 := __e.Get(1)
			_ = V453
			gen10979 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(MakeNumber(-1))
				return
			}, 1)
			gen10981 := MakeNative(func(__e Evaluator) {
				gen10980 := Call(__e, __e.Global(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(__e.Global(symget), V453, MakeSymbol("arity"), gen10980)

				return

			}, 0)
			__e.Return(Try(__e, gen10981).Catch(gen10979))
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("arity"), gen10982)

		gen10989 := MakeNative(func(__e Evaluator) {
			V455 := __e.Get(1)
			_ = V455
			gen10983 := Call(__e, __e.Global(symintern), MakeString("shen"))

			Shen := gen10983
			gen10984 := Call(__e, __e.Global(symvalue), MakeSymbol("*property-vector*"))

			gen10985 := Call(__e, __e.Global(symget), Shen, MakeSymbol("shen.external-symbols"), gen10984)

			External := gen10985
			gen10986 := Call(__e, __e.Global(symadjoin), V455, External)

			gen10987 := Call(__e, __e.Global(symvalue), MakeSymbol("*property-vector*"))

			gen10988 := Call(__e, __e.Global(symput), Shen, MakeSymbol("shen.external-symbols"), gen10986, gen10987)

			Place := gen10988
			_ = Place
			__e.Return(V455)
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("systemf"), gen10989)

		gen10991 := MakeNative(func(__e Evaluator) {
			V458 := __e.Get(1)
			_ = V458
			V459 := __e.Get(2)
			_ = V459
			gen10990 := Call(__e, __e.Global(symelement_2), V458, V459)

			if True == gen10990 {
				__e.Return(V459)
				return
			} else {
				__e.TailApply(__e.Global(symcons), V458, V459)

				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("adjoin"), gen10991)

		gen11000 := MakeNative(func(__e Evaluator) {
			V461 := __e.Get(1)
			_ = V461
			gen10999 := Call(__e, __e.Global(sym_a), MakeSymbol("package"), V461)

			if True == gen10999 {
				__e.Return(Nil)
				return
			} else {
				gen10998 := Call(__e, __e.Global(sym_a), MakeSymbol("receive"), V461)

				if True == gen10998 {
					__e.Return(Nil)
					return
				} else {
					gen10992 := Call(__e, __e.Global(symarity), V461)

					ArityF := gen10992
					gen10997 := Call(__e, __e.Global(sym_a), ArityF, MakeNumber(-1))

					if True == gen10997 {
						__e.Return(Nil)
						return
					} else {
						gen10996 := Call(__e, __e.Global(sym_a), ArityF, MakeNumber(0))

						if True == gen10996 {
							__e.Return(Nil)
							return
						} else {
							gen10993 := Call(__e, __e.Global(symshen_4lambda_1form), V461, ArityF)

							gen10994 := Call(__e, __e.Global(symeval_1kl), gen10993)

							gen10995 := Call(__e, __e.Global(symcons), V461, gen10994)

							__e.TailApply(__e.Global(symcons), gen10995, Nil)

							return

						}

					}

				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.lambda-form-entry"), gen11000)

		gen11008 := MakeNative(func(__e Evaluator) {
			V464 := __e.Get(1)
			_ = V464
			V465 := __e.Get(2)
			_ = V465
			gen11007 := Call(__e, __e.Global(sym_a), MakeNumber(0), V465)

			if True == gen11007 {
				__e.Return(V464)
				return
			} else {
				gen11001 := Call(__e, __e.Global(symgensym), MakeSymbol("V"))

				X := gen11001
				gen11002 := Call(__e, __e.Global(symshen_4add_1end), V464, X)

				gen11003 := Call(__e, __e.Global(sym_1), V465, MakeNumber(1))

				gen11004 := Call(__e, __e.Global(symshen_4lambda_1form), gen11002, gen11003)

				gen11005 := Call(__e, __e.Global(symcons), gen11004, Nil)

				gen11006 := Call(__e, __e.Global(symcons), X, gen11005)

				__e.TailApply(__e.Global(symcons), MakeSymbol("lambda"), gen11006)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.lambda-form"), gen11008)

		gen11012 := MakeNative(func(__e Evaluator) {
			V468 := __e.Get(1)
			_ = V468
			V469 := __e.Get(2)
			_ = V469
			gen11011 := Call(__e, __e.Global(symcons_2), V468)

			if True == gen11011 {
				gen11010 := Call(__e, __e.Global(symcons), V469, Nil)

				__e.TailApply(__e.Global(symappend), V468, gen11010)

				return

			} else {
				gen11009 := Call(__e, __e.Global(symcons), V469, Nil)

				__e.TailApply(__e.Global(symcons), V468, gen11009)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.add-end"), gen11012)

		gen11017 := MakeNative(func(__e Evaluator) {
			V471 := __e.Get(1)
			_ = V471
			gen11016 := Call(__e, __e.Global(symcons_2), V471)

			if True == gen11016 {
				gen11013 := Call(__e, __e.Global(symhd), V471)

				gen11014 := Call(__e, __e.Global(symtl), V471)

				gen11015 := Call(__e, __e.Global(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(__e.Global(symput), gen11013, MakeSymbol("shen.lambda-form"), gen11014, gen11015)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.set-lambda-form-entry"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.set-lambda-form-entry"), gen11017)

		gen11020 := MakeNative(func(__e Evaluator) {
			V473 := __e.Get(1)
			_ = V473
			gen11018 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*special*"))

			gen11019 := Call(__e, __e.Global(symcons), V473, gen11018)

			Call(__e, __e.Global(symset), MakeSymbol("shen.*special*"), gen11019)

			__e.Return(V473)
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("specialise"), gen11020)

		gen11023 := MakeNative(func(__e Evaluator) {
			V475 := __e.Get(1)
			_ = V475
			gen11021 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*special*"))

			gen11022 := Call(__e, __e.Global(symremove), V475, gen11021)

			Call(__e, __e.Global(symset), MakeSymbol("shen.*special*"), gen11022)

			__e.Return(V475)
			return

		}, 1)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("unspecialise"), gen11023)

		return

	}, 0))
}
