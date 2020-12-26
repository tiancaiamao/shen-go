package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen10796 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V451 := __args[0]
			_ = V451
			gen10795 := Call(__e, ShenFunc(sym_a), Nil, V451)

			if True == gen10795 {
				__e.Return(Nil)
				return
			} else {
				gen10793 := Call(__e, ShenFunc(symcons_2), V451)

				var gen10794 Obj
				if True == gen10793 {
					gen10791 := Call(__e, ShenFunc(symtl), V451)

					gen10792 := Call(__e, ShenFunc(symcons_2), gen10791)

					if True == gen10792 {
						gen10794 = True
					} else {
						gen10794 = False
					}

				} else {
					gen10794 = False
				}
				if True == gen10794 {
					gen10784 := Call(__e, ShenFunc(symhd), V451)

					gen10785 := Call(__e, ShenFunc(symtl), V451)

					gen10786 := Call(__e, ShenFunc(symhd), gen10785)

					gen10787 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

					gen10788 := Call(__e, ShenFunc(symput), gen10784, MakeSymbol("arity"), gen10786, gen10787)

					DecArity := gen10788
					_ = DecArity
					gen10789 := Call(__e, ShenFunc(symtl), V451)

					gen10790 := Call(__e, ShenFunc(symtl), gen10789)

					__e.TailApply(ShenFunc(symshen_4initialise__arity__table), gen10790)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.initialise_arity_table"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise_arity_table"), gen10796)

		gen10800 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V453 := __args[0]
			_ = V453
			gen10797 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(MakeNumber(-1))
				return
			}, 1)
			gen10799 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen10798 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V453, MakeSymbol("arity"), gen10798)

				return

			}, 0)
			__e.Return(Try(__e, gen10799).Catch(gen10797))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("arity"), gen10800)

		gen10807 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V455 := __args[0]
			_ = V455
			gen10801 := Call(__e, ShenFunc(symintern), MakeString("shen"))

			Shen := gen10801
			gen10802 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			gen10803 := Call(__e, ShenFunc(symget), Shen, MakeSymbol("shen.external-symbols"), gen10802)

			External := gen10803
			gen10804 := Call(__e, ShenFunc(symadjoin), V455, External)

			gen10805 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			gen10806 := Call(__e, ShenFunc(symput), Shen, MakeSymbol("shen.external-symbols"), gen10804, gen10805)

			Place := gen10806
			_ = Place
			__e.Return(V455)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("systemf"), gen10807)

		gen10809 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V458 := __args[0]
			_ = V458
			V459 := __args[1]
			_ = V459
			gen10808 := Call(__e, ShenFunc(symelement_2), V458, V459)

			if True == gen10808 {
				__e.Return(V459)
				return
			} else {
				__e.TailApply(ShenFunc(symcons), V458, V459)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("adjoin"), gen10809)

		gen10818 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V461 := __args[0]
			_ = V461
			gen10817 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), V461)

			if True == gen10817 {
				__e.Return(Nil)
				return
			} else {
				gen10816 := Call(__e, ShenFunc(sym_a), MakeSymbol("receive"), V461)

				if True == gen10816 {
					__e.Return(Nil)
					return
				} else {
					gen10810 := Call(__e, ShenFunc(symarity), V461)

					ArityF := gen10810
					gen10815 := Call(__e, ShenFunc(sym_a), ArityF, MakeNumber(-1))

					if True == gen10815 {
						__e.Return(Nil)
						return
					} else {
						gen10814 := Call(__e, ShenFunc(sym_a), ArityF, MakeNumber(0))

						if True == gen10814 {
							__e.Return(Nil)
							return
						} else {
							gen10811 := Call(__e, ShenFunc(symshen_4lambda_1form), V461, ArityF)

							gen10812 := Call(__e, ShenFunc(symeval_1kl), gen10811)

							gen10813 := Call(__e, ShenFunc(symcons), V461, gen10812)

							__e.TailApply(ShenFunc(symcons), gen10813, Nil)

							return

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lambda-form-entry"), gen10818)

		gen10826 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V464 := __args[0]
			_ = V464
			V465 := __args[1]
			_ = V465
			gen10825 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V465)

			if True == gen10825 {
				__e.Return(V464)
				return
			} else {
				gen10819 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

				X := gen10819
				gen10820 := Call(__e, ShenFunc(symshen_4add_1end), V464, X)

				gen10821 := Call(__e, ShenFunc(sym_1), V465, MakeNumber(1))

				gen10822 := Call(__e, ShenFunc(symshen_4lambda_1form), gen10820, gen10821)

				gen10823 := Call(__e, ShenFunc(symcons), gen10822, Nil)

				gen10824 := Call(__e, ShenFunc(symcons), X, gen10823)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen10824)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lambda-form"), gen10826)

		gen10830 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V468 := __args[0]
			_ = V468
			V469 := __args[1]
			_ = V469
			gen10829 := Call(__e, ShenFunc(symcons_2), V468)

			if True == gen10829 {
				gen10828 := Call(__e, ShenFunc(symcons), V469, Nil)

				__e.TailApply(ShenFunc(symappend), V468, gen10828)

				return

			} else {
				gen10827 := Call(__e, ShenFunc(symcons), V469, Nil)

				__e.TailApply(ShenFunc(symcons), V468, gen10827)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.add-end"), gen10830)

		gen10835 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V471 := __args[0]
			_ = V471
			gen10834 := Call(__e, ShenFunc(symcons_2), V471)

			if True == gen10834 {
				gen10831 := Call(__e, ShenFunc(symhd), V471)

				gen10832 := Call(__e, ShenFunc(symtl), V471)

				gen10833 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), gen10831, MakeSymbol("shen.lambda-form"), gen10832, gen10833)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.set-lambda-form-entry"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.set-lambda-form-entry"), gen10835)

		gen10838 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V473 := __args[0]
			_ = V473
			gen10836 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*special*"))

			gen10837 := Call(__e, ShenFunc(symcons), V473, gen10836)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*special*"), gen10837)

			__e.Return(V473)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("specialise"), gen10838)

		gen10841 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V475 := __args[0]
			_ = V475
			gen10839 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*special*"))

			gen10840 := Call(__e, ShenFunc(symremove), V475, gen10839)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*special*"), gen10840)

			__e.Return(V475)
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("unspecialise"), gen10841)

		return

	}, 0))
}
