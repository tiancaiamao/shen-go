package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen9586 := MakeNative(func(__e Evaluator) {
			V701 := __e.Get(1)
			_ = V701
			gen9569 := Call(__e, __e.Global(symget_1time), MakeSymbol("run"))

			Start := gen9569
			gen9570 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*tc*"))

			gen9571 := Call(__e, __e.Global(symread_1file), V701)

			gen9572 := Call(__e, __e.Global(symshen_4load_1help), gen9570, gen9571)

			Result := gen9572
			gen9573 := Call(__e, __e.Global(symget_1time), MakeSymbol("run"))

			Finish := gen9573
			gen9574 := Call(__e, __e.Global(sym_1), Finish, Start)

			Time := gen9574
			gen9575 := Call(__e, __e.Global(symstr), Time)

			gen9576 := Call(__e, __e.Global(symcn), gen9575, MakeString(" secs\n"))

			gen9577 := Call(__e, __e.Global(symcn), MakeString("\nrun time: "), gen9576)

			gen9578 := Call(__e, __e.Global(symstoutput))

			gen9579 := Call(__e, __e.Global(symshen_4prhush), gen9577, gen9578)

			Message := gen9579
			_ = Message
			Load := Result
			_ = Load
			gen9584 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*tc*"))

			var gen9585 Obj
			if True == gen9584 {
				gen9580 := Call(__e, __e.Global(syminferences))

				gen9581 := Call(__e, __e.Global(symshen_4app), gen9580, MakeString(" inferences\n"), MakeSymbol("shen.a"))

				gen9582 := Call(__e, __e.Global(symcn), MakeString("\ntypechecked in "), gen9581)

				gen9583 := Call(__e, __e.Global(symstoutput))

				gen9585 = Call(__e, __e.Global(symshen_4prhush), gen9582, gen9583)

			} else {
				gen9585 = MakeSymbol("shen.skip")
			}
			Infs := gen9585
			_ = Infs
			__e.Return(MakeSymbol("loaded"))
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("load"), gen9586)

		gen9601 := MakeNative(func(__e Evaluator) {
			V708 := __e.Get(1)
			_ = V708
			V709 := __e.Get(2)
			_ = V709
			gen9600 := Call(__e, __e.Global(sym_a), False, V708)

			if True == gen9600 {
				gen9599 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					gen9596 := Call(__e, __e.Global(symshen_4eval_1without_1macros), X)

					gen9597 := Call(__e, __e.Global(symshen_4app), gen9596, MakeString("\n"), MakeSymbol("shen.s"))

					gen9598 := Call(__e, __e.Global(symstoutput))

					__e.TailApply(__e.Global(symshen_4prhush), gen9597, gen9598)

					return

				}, 1)
				__e.TailApply(__e.Global(symshen_4for_1each), gen9599, V709)

				return

			} else {
				gen9587 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(__e.Global(symshen_4remove_1synonyms), X)

					return
				}, 1)
				gen9588 := Call(__e, __e.Global(symmapcan), gen9587, V709)

				RemoveSynonyms := gen9588
				gen9589 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(__e.Global(symshen_4typetable), X)

					return
				}, 1)
				gen9590 := Call(__e, __e.Global(symmapcan), gen9589, RemoveSynonyms)

				Table := gen9590
				gen9591 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(__e.Global(symshen_4assumetype), X)

					return
				}, 1)
				gen9592 := Call(__e, __e.Global(symshen_4for_1each), gen9591, Table)

				Assume := gen9592
				_ = Assume
				gen9593 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.TailApply(__e.Global(symshen_4unwind_1types), E, Table)

					return
				}, 1)
				gen9595 := MakeNative(func(__e Evaluator) {
					gen9594 := MakeNative(func(__e Evaluator) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(__e.Global(symshen_4typecheck_1and_1load), X)

						return
					}, 1)
					__e.TailApply(__e.Global(symshen_4for_1each), gen9594, RemoveSynonyms)

					return

				}, 0)
				__e.Return(Try(__e, gen9595).Catch(gen9593))
				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.load-help"), gen9601)

		gen9606 := MakeNative(func(__e Evaluator) {
			V711 := __e.Get(1)
			_ = V711
			gen9604 := Call(__e, __e.Global(symcons_2), V711)

			var gen9605 Obj
			if True == gen9604 {
				gen9602 := Call(__e, __e.Global(symhd), V711)

				gen9603 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.synonyms-help"), gen9602)

				if True == gen9603 {
					gen9605 = True
				} else {
					gen9605 = False
				}

			} else {
				gen9605 = False
			}
			if True == gen9605 {
				Call(__e, __e.Global(symeval), V711)
				__e.Return(Nil)
				return

			} else {
				__e.TailApply(__e.Global(symcons), V711, Nil)

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.remove-synonyms"), gen9606)

		gen9608 := MakeNative(func(__e Evaluator) {
			V713 := __e.Get(1)
			_ = V713
			Call(__e, __e.Global(symnl), MakeNumber(1))
			gen9607 := Call(__e, __e.Global(symgensym), MakeSymbol("A"))

			__e.TailApply(__e.Global(symshen_4typecheck_1and_1evaluate), V713, gen9607)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.typecheck-and-load"), gen9608)

		gen9627 := MakeNative(func(__e Evaluator) {
			V719 := __e.Get(1)
			_ = V719
			gen9625 := Call(__e, __e.Global(symcons_2), V719)

			var gen9626 Obj
			if True == gen9625 {
				gen9622 := Call(__e, __e.Global(symhd), V719)

				gen9623 := Call(__e, __e.Global(sym_a), MakeSymbol("define"), gen9622)

				var gen9624 Obj
				if True == gen9623 {
					gen9620 := Call(__e, __e.Global(symtl), V719)

					gen9621 := Call(__e, __e.Global(symcons_2), gen9620)

					if True == gen9621 {
						gen9624 = True
					} else {
						gen9624 = False
					}

				} else {
					gen9624 = False
				}
				if True == gen9624 {
					gen9626 = True
				} else {
					gen9626 = False
				}

			} else {
				gen9626 = False
			}
			if True == gen9626 {
				gen9609 := MakeNative(func(__e Evaluator) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(__e.Global(symshen_4_5sig_7rest_6), Y)

					return
				}, 1)
				gen9610 := Call(__e, __e.Global(symtl), V719)

				gen9611 := Call(__e, __e.Global(symtl), gen9610)

				gen9615 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					gen9612 := Call(__e, __e.Global(symtl), V719)

					gen9613 := Call(__e, __e.Global(symhd), gen9612)

					gen9614 := Call(__e, __e.Global(symshen_4app), gen9613, MakeString(" lacks a proper signature.\n"), MakeSymbol("shen.a"))

					__e.TailApply(__e.Global(symsimple_1error), gen9614)

					return

				}, 1)
				gen9616 := Call(__e, __e.Global(symcompile), gen9609, gen9611, gen9615)

				Sig := gen9616
				gen9617 := Call(__e, __e.Global(symtl), V719)

				gen9618 := Call(__e, __e.Global(symhd), gen9617)

				gen9619 := Call(__e, __e.Global(symcons), gen9618, Sig)

				__e.TailApply(__e.Global(symcons), gen9619, Nil)

				return

			} else {
				__e.Return(Nil)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.typetable"), gen9627)

		gen9631 := MakeNative(func(__e Evaluator) {
			V721 := __e.Get(1)
			_ = V721
			gen9630 := Call(__e, __e.Global(symcons_2), V721)

			if True == gen9630 {
				gen9628 := Call(__e, __e.Global(symhd), V721)

				gen9629 := Call(__e, __e.Global(symtl), V721)

				__e.TailApply(__e.Global(symdeclare), gen9628, gen9629)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.assumetype"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.assumetype"), gen9631)

		gen9641 := MakeNative(func(__e Evaluator) {
			V728 := __e.Get(1)
			_ = V728
			V729 := __e.Get(2)
			_ = V729
			gen9640 := Call(__e, __e.Global(sym_a), Nil, V729)

			if True == gen9640 {
				gen9639 := Call(__e, __e.Global(symerror_1to_1string), V728)

				__e.TailApply(__e.Global(symsimple_1error), gen9639)

				return

			} else {
				gen9637 := Call(__e, __e.Global(symcons_2), V729)

				var gen9638 Obj
				if True == gen9637 {
					gen9635 := Call(__e, __e.Global(symhd), V729)

					gen9636 := Call(__e, __e.Global(symcons_2), gen9635)

					if True == gen9636 {
						gen9638 = True
					} else {
						gen9638 = False
					}

				} else {
					gen9638 = False
				}
				if True == gen9638 {
					gen9632 := Call(__e, __e.Global(symhd), V729)

					gen9633 := Call(__e, __e.Global(symhd), gen9632)

					Call(__e, __e.Global(symshen_4remtype), gen9633)

					gen9634 := Call(__e, __e.Global(symtl), V729)

					__e.TailApply(__e.Global(symshen_4unwind_1types), V728, gen9634)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.unwind-types"))

					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.unwind-types"), gen9641)

		gen9644 := MakeNative(func(__e Evaluator) {
			V731 := __e.Get(1)
			_ = V731
			gen9642 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen9643 := Call(__e, __e.Global(symshen_4removetype), V731, gen9642)

			__e.TailApply(__e.Global(symset), MakeSymbol("shen.*signedfuncs*"), gen9643)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.remtype"), gen9644)

		gen9661 := MakeNative(func(__e Evaluator) {
			V739 := __e.Get(1)
			_ = V739
			V740 := __e.Get(2)
			_ = V740
			gen9660 := Call(__e, __e.Global(sym_a), Nil, V740)

			if True == gen9660 {
				__e.Return(Nil)
				return
			} else {
				gen9658 := Call(__e, __e.Global(symcons_2), V740)

				var gen9659 Obj
				if True == gen9658 {
					gen9655 := Call(__e, __e.Global(symhd), V740)

					gen9656 := Call(__e, __e.Global(symcons_2), gen9655)

					var gen9657 Obj
					if True == gen9656 {
						gen9652 := Call(__e, __e.Global(symhd), V740)

						gen9653 := Call(__e, __e.Global(symhd), gen9652)

						gen9654 := Call(__e, __e.Global(sym_a), gen9653, V739)

						if True == gen9654 {
							gen9657 = True
						} else {
							gen9657 = False
						}

					} else {
						gen9657 = False
					}
					if True == gen9657 {
						gen9659 = True
					} else {
						gen9659 = False
					}

				} else {
					gen9659 = False
				}
				if True == gen9659 {
					gen9649 := Call(__e, __e.Global(symhd), V740)

					gen9650 := Call(__e, __e.Global(symhd), gen9649)

					gen9651 := Call(__e, __e.Global(symtl), V740)

					__e.TailApply(__e.Global(symshen_4removetype), gen9650, gen9651)

					return

				} else {
					gen9648 := Call(__e, __e.Global(symcons_2), V740)

					if True == gen9648 {
						gen9645 := Call(__e, __e.Global(symhd), V740)

						gen9646 := Call(__e, __e.Global(symtl), V740)

						gen9647 := Call(__e, __e.Global(symshen_4removetype), V739, gen9646)

						__e.TailApply(__e.Global(symcons), gen9645, gen9647)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.removetype"))

						return
					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.removetype"), gen9661)

		gen9672 := MakeNative(func(__e Evaluator) {
			V742 := __e.Get(1)
			_ = V742
			gen9662 := Call(__e, __e.Global(symshen_4_5signature_6), V742)

			Parse__shen_4_5signature_6 := gen9662
			gen9669 := Call(__e, __e.Global(symfail))

			gen9670 := Call(__e, __e.Global(sym_a), gen9669, Parse__shen_4_5signature_6)

			gen9671 := Call(__e, __e.Global(symnot), gen9670)

			if True == gen9671 {
				gen9663 := Call(__e, __e.Global(sym_5_b_6), Parse__shen_4_5signature_6)

				Parse___5_b_6 := gen9663
				gen9666 := Call(__e, __e.Global(symfail))

				gen9667 := Call(__e, __e.Global(sym_a), gen9666, Parse___5_b_6)

				gen9668 := Call(__e, __e.Global(symnot), gen9667)

				if True == gen9668 {
					gen9664 := Call(__e, __e.Global(symhd), Parse___5_b_6)

					gen9665 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5signature_6)

					__e.TailApply(__e.Global(symshen_4pair), gen9664, gen9665)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<sig+rest>"), gen9672)

		gen9678 := MakeNative(func(__e Evaluator) {
			V745 := __e.Get(1)
			_ = V745
			V746 := __e.Get(2)
			_ = V746
			gen9673 := Call(__e, __e.Global(symopen), V745, MakeSymbol("out"))

			Stream := gen9673
			gen9674 := Call(__e, __e.Global(symstring_2), V746)

			var gen9675 Obj
			if True == gen9674 {
				gen9675 = Call(__e, __e.Global(symshen_4app), V746, MakeString("\n\n"), MakeSymbol("shen.a"))

			} else {
				gen9675 = Call(__e, __e.Global(symshen_4app), V746, MakeString("\n\n"), MakeSymbol("shen.s"))

			}
			String := gen9675
			gen9676 := Call(__e, __e.Global(sympr), String, Stream)

			Write := gen9676
			_ = Write
			gen9677 := Call(__e, __e.Global(symclose), Stream)

			Close := gen9677
			_ = Close
			__e.Return(V746)
			return

		}, 2)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("write-to-file"), gen9678)

		return

	}, 0))
}
