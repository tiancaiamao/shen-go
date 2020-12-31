package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen13713 := MakeNative(func(__e Evaluator) {
			V3210 := __e.Get(1)
			_ = V3210
			V3211 := __e.Get(2)
			_ = V3211
			gen13682 := Call(__e, __e.Global(symcons), V3210, V3211)

			gen13683 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen13684 := Call(__e, __e.Global(symcons), gen13682, gen13683)

			gen13685 := Call(__e, __e.Global(symset), MakeSymbol("shen.*signedfuncs*"), gen13684)

			Record := gen13685
			_ = Record
			gen13686 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(MakeSymbol("shen.skip"))
				return
			}, 1)
			gen13687 := MakeNative(func(__e Evaluator) {
				__e.TailApply(__e.Global(symshen_4variancy_1test), V3210, V3211)

				return
			}, 0)
			gen13688 := Try(__e, gen13687).Catch(gen13686)
			Variancy := gen13688
			_ = Variancy
			gen13689 := Call(__e, __e.Global(symshen_4demodulate), V3211)

			gen13690 := Call(__e, __e.Global(symshen_4rcons__form), gen13689)

			Type := gen13690
			gen13691 := Call(__e, __e.Global(symconcat), MakeSymbol("shen.type-signature-of-"), V3210)

			F_d := gen13691
			gen13692 := Call(__e, __e.Global(symshen_4parameters), MakeNumber(1))

			Parameters := gen13692
			gen13693 := Call(__e, __e.Global(symcons), MakeSymbol("X"), Nil)

			gen13694 := Call(__e, __e.Global(symcons), F_d, gen13693)

			gen13695 := Call(__e, __e.Global(symcons), Type, Nil)

			gen13696 := Call(__e, __e.Global(symcons), MakeSymbol("X"), gen13695)

			gen13697 := Call(__e, __e.Global(symcons), MakeSymbol("unify!"), gen13696)

			gen13698 := Call(__e, __e.Global(symcons), gen13697, Nil)

			gen13699 := Call(__e, __e.Global(symcons), gen13698, Nil)

			gen13700 := Call(__e, __e.Global(symcons), MakeSymbol(":-"), gen13699)

			gen13701 := Call(__e, __e.Global(symcons), gen13694, gen13700)

			Clause := gen13701
			gen13702 := Call(__e, __e.Global(symshen_4aum), Clause, Parameters)

			AUM__instruction := gen13702
			gen13703 := Call(__e, __e.Global(symshen_4aum__to__shen), AUM__instruction)

			Code := gen13703
			gen13704 := Call(__e, __e.Global(symcons), MakeSymbol("Continuation"), Nil)

			gen13705 := Call(__e, __e.Global(symcons), MakeSymbol("ProcessN"), gen13704)

			gen13706 := Call(__e, __e.Global(symcons), Code, Nil)

			gen13707 := Call(__e, __e.Global(symcons), MakeSymbol("->"), gen13706)

			gen13708 := Call(__e, __e.Global(symappend), gen13705, gen13707)

			gen13709 := Call(__e, __e.Global(symappend), Parameters, gen13708)

			gen13710 := Call(__e, __e.Global(symcons), F_d, gen13709)

			gen13711 := Call(__e, __e.Global(symcons), MakeSymbol("define"), gen13710)

			ShenDef := gen13711
			gen13712 := Call(__e, __e.Global(symshen_4eval_1without_1macros), ShenDef)

			Eval := gen13712
			_ = Eval
			__e.Return(V3210)
			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("declare"), gen13713)

		gen13717 := MakeNative(func(__e Evaluator) {
			V3213 := __e.Get(1)
			_ = V3213
			gen13714 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*demodulation-function*"))

			gen13715 := Call(__e, __e.Global(symshen_4walk), gen13714, V3213)

			Demod := gen13715
			gen13716 := Call(__e, __e.Global(sym_a), Demod, V3213)

			if True == gen13716 {
				__e.Return(V3213)
				return
			} else {
				__e.TailApply(__e.Global(symshen_4demodulate), Demod)

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.demodulate"), gen13717)

		gen13725 := MakeNative(func(__e Evaluator) {
			V3216 := __e.Get(1)
			_ = V3216
			V3217 := __e.Get(2)
			_ = V3217
			gen13718 := Call(__e, __e.Global(symshen_4typecheck), V3216, MakeSymbol("B"))

			TypeF := gen13718
			gen13723 := Call(__e, __e.Global(sym_a), MakeSymbol("symbol"), TypeF)

			var gen13724 Obj
			if True == gen13723 {
				gen13724 = MakeSymbol("shen.skip")
			} else {
				gen13722 := Call(__e, __e.Global(symshen_4variant_2), TypeF, V3217)

				if True == gen13722 {
					gen13724 = MakeSymbol("shen.skip")
				} else {
					gen13719 := Call(__e, __e.Global(symshen_4app), V3216, MakeString(" may create errors\n"), MakeSymbol("shen.a"))

					gen13720 := Call(__e, __e.Global(symcn), MakeString("warning: changing the type of "), gen13719)

					gen13721 := Call(__e, __e.Global(symstoutput))

					gen13724 = Call(__e, __e.Global(symshen_4prhush), gen13720, gen13721)

				}

			}
			Check := gen13724
			_ = Check
			__e.Return(MakeSymbol("shen.skip"))
			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.variancy-test"), gen13725)

		gen13766 := MakeNative(func(__e Evaluator) {
			V3230 := __e.Get(1)
			_ = V3230
			V3231 := __e.Get(2)
			_ = V3231
			gen13765 := Call(__e, __e.Global(sym_a), V3231, V3230)

			if True == gen13765 {
				__e.Return(True)
				return
			} else {
				gen13763 := Call(__e, __e.Global(symcons_2), V3230)

				var gen13764 Obj
				if True == gen13763 {
					gen13761 := Call(__e, __e.Global(symcons_2), V3231)

					var gen13762 Obj
					if True == gen13761 {
						gen13758 := Call(__e, __e.Global(symhd), V3231)

						gen13759 := Call(__e, __e.Global(symhd), V3230)

						gen13760 := Call(__e, __e.Global(sym_a), gen13758, gen13759)

						if True == gen13760 {
							gen13762 = True
						} else {
							gen13762 = False
						}

					} else {
						gen13762 = False
					}
					if True == gen13762 {
						gen13764 = True
					} else {
						gen13764 = False
					}

				} else {
					gen13764 = False
				}
				if True == gen13764 {
					gen13756 := Call(__e, __e.Global(symtl), V3230)

					gen13757 := Call(__e, __e.Global(symtl), V3231)

					__e.TailApply(__e.Global(symshen_4variant_2), gen13756, gen13757)

					return

				} else {
					gen13754 := Call(__e, __e.Global(symcons_2), V3230)

					var gen13755 Obj
					if True == gen13754 {
						gen13752 := Call(__e, __e.Global(symcons_2), V3231)

						var gen13753 Obj
						if True == gen13752 {
							gen13749 := Call(__e, __e.Global(symhd), V3230)

							gen13750 := Call(__e, __e.Global(symshen_4pvar_2), gen13749)

							var gen13751 Obj
							if True == gen13750 {
								gen13747 := Call(__e, __e.Global(symhd), V3231)

								gen13748 := Call(__e, __e.Global(symvariable_2), gen13747)

								if True == gen13748 {
									gen13751 = True
								} else {
									gen13751 = False
								}

							} else {
								gen13751 = False
							}
							if True == gen13751 {
								gen13753 = True
							} else {
								gen13753 = False
							}

						} else {
							gen13753 = False
						}
						if True == gen13753 {
							gen13755 = True
						} else {
							gen13755 = False
						}

					} else {
						gen13755 = False
					}
					if True == gen13755 {
						gen13741 := Call(__e, __e.Global(symhd), V3230)

						gen13742 := Call(__e, __e.Global(symtl), V3230)

						gen13743 := Call(__e, __e.Global(symsubst), MakeSymbol("shen.a"), gen13741, gen13742)

						gen13744 := Call(__e, __e.Global(symhd), V3231)

						gen13745 := Call(__e, __e.Global(symtl), V3231)

						gen13746 := Call(__e, __e.Global(symsubst), MakeSymbol("shen.a"), gen13744, gen13745)

						__e.TailApply(__e.Global(symshen_4variant_2), gen13743, gen13746)

						return

					} else {
						gen13739 := Call(__e, __e.Global(symcons_2), V3230)

						var gen13740 Obj
						if True == gen13739 {
							gen13736 := Call(__e, __e.Global(symhd), V3230)

							gen13737 := Call(__e, __e.Global(symcons_2), gen13736)

							var gen13738 Obj
							if True == gen13737 {
								gen13734 := Call(__e, __e.Global(symcons_2), V3231)

								var gen13735 Obj
								if True == gen13734 {
									gen13732 := Call(__e, __e.Global(symhd), V3231)

									gen13733 := Call(__e, __e.Global(symcons_2), gen13732)

									if True == gen13733 {
										gen13735 = True
									} else {
										gen13735 = False
									}

								} else {
									gen13735 = False
								}
								if True == gen13735 {
									gen13738 = True
								} else {
									gen13738 = False
								}

							} else {
								gen13738 = False
							}
							if True == gen13738 {
								gen13740 = True
							} else {
								gen13740 = False
							}

						} else {
							gen13740 = False
						}
						if True == gen13740 {
							gen13726 := Call(__e, __e.Global(symhd), V3230)

							gen13727 := Call(__e, __e.Global(symtl), V3230)

							gen13728 := Call(__e, __e.Global(symappend), gen13726, gen13727)

							gen13729 := Call(__e, __e.Global(symhd), V3231)

							gen13730 := Call(__e, __e.Global(symtl), V3231)

							gen13731 := Call(__e, __e.Global(symappend), gen13729, gen13730)

							__e.TailApply(__e.Global(symshen_4variant_2), gen13728, gen13731)

							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.variant?"), gen13766)

		gen13771 := MakeNative(func(__e Evaluator) {
			V3236 := __e.Get(1)
			_ = V3236
			V3237 := __e.Get(2)
			_ = V3237
			V3238 := __e.Get(3)
			_ = V3238
			gen13767 := Call(__e, __e.Global(symshen_4newpv), V3237)

			A := gen13767
			Call(__e, __e.Global(symshen_4incinfs))
			gen13768 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13769 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13768)

			gen13770 := Call(__e, __e.Global(symcons), A, gen13769)

			__e.TailApply(__e.Global(symunify_b), V3236, gen13770, V3237, V3238)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-absvector?"), gen13771)

		gen13783 := MakeNative(func(__e Evaluator) {
			V3246 := __e.Get(1)
			_ = V3246
			V3247 := __e.Get(2)
			_ = V3247
			V3248 := __e.Get(3)
			_ = V3248
			gen13772 := Call(__e, __e.Global(symshen_4newpv), V3247)

			A := gen13772
			Call(__e, __e.Global(symshen_4incinfs))
			gen13773 := Call(__e, __e.Global(symcons), A, Nil)

			gen13774 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13773)

			gen13775 := Call(__e, __e.Global(symcons), A, Nil)

			gen13776 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13775)

			gen13777 := Call(__e, __e.Global(symcons), gen13776, Nil)

			gen13778 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13777)

			gen13779 := Call(__e, __e.Global(symcons), gen13774, gen13778)

			gen13780 := Call(__e, __e.Global(symcons), gen13779, Nil)

			gen13781 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13780)

			gen13782 := Call(__e, __e.Global(symcons), A, gen13781)

			__e.TailApply(__e.Global(symunify_b), V3246, gen13782, V3247, V3248)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-adjoin"), gen13783)

		gen13790 := MakeNative(func(__e Evaluator) {
			V3256 := __e.Get(1)
			_ = V3256
			V3257 := __e.Get(2)
			_ = V3257
			V3258 := __e.Get(3)
			_ = V3258
			Call(__e, __e.Global(symshen_4incinfs))
			gen13784 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13785 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13784)

			gen13786 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), gen13785)

			gen13787 := Call(__e, __e.Global(symcons), gen13786, Nil)

			gen13788 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13787)

			gen13789 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), gen13788)

			__e.TailApply(__e.Global(symunify_b), V3256, gen13789, V3257, V3258)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-and"), gen13790)

		gen13801 := MakeNative(func(__e Evaluator) {
			V3266 := __e.Get(1)
			_ = V3266
			V3267 := __e.Get(2)
			_ = V3267
			V3268 := __e.Get(3)
			_ = V3268
			gen13791 := Call(__e, __e.Global(symshen_4newpv), V3267)

			A := gen13791
			Call(__e, __e.Global(symshen_4incinfs))
			gen13792 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen13793 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13792)

			gen13794 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen13793)

			gen13795 := Call(__e, __e.Global(symcons), gen13794, Nil)

			gen13796 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13795)

			gen13797 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen13796)

			gen13798 := Call(__e, __e.Global(symcons), gen13797, Nil)

			gen13799 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13798)

			gen13800 := Call(__e, __e.Global(symcons), A, gen13799)

			__e.TailApply(__e.Global(symunify_b), V3266, gen13800, V3267, V3268)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-shen.app"), gen13801)

		gen13815 := MakeNative(func(__e Evaluator) {
			V3276 := __e.Get(1)
			_ = V3276
			V3277 := __e.Get(2)
			_ = V3277
			V3278 := __e.Get(3)
			_ = V3278
			gen13802 := Call(__e, __e.Global(symshen_4newpv), V3277)

			A := gen13802
			Call(__e, __e.Global(symshen_4incinfs))
			gen13803 := Call(__e, __e.Global(symcons), A, Nil)

			gen13804 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13803)

			gen13805 := Call(__e, __e.Global(symcons), A, Nil)

			gen13806 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13805)

			gen13807 := Call(__e, __e.Global(symcons), A, Nil)

			gen13808 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13807)

			gen13809 := Call(__e, __e.Global(symcons), gen13808, Nil)

			gen13810 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13809)

			gen13811 := Call(__e, __e.Global(symcons), gen13806, gen13810)

			gen13812 := Call(__e, __e.Global(symcons), gen13811, Nil)

			gen13813 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13812)

			gen13814 := Call(__e, __e.Global(symcons), gen13804, gen13813)

			__e.TailApply(__e.Global(symunify_b), V3276, gen13814, V3277, V3278)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-append"), gen13815)

		gen13820 := MakeNative(func(__e Evaluator) {
			V3286 := __e.Get(1)
			_ = V3286
			V3287 := __e.Get(2)
			_ = V3287
			V3288 := __e.Get(3)
			_ = V3288
			gen13816 := Call(__e, __e.Global(symshen_4newpv), V3287)

			A := gen13816
			Call(__e, __e.Global(symshen_4incinfs))
			gen13817 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen13818 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13817)

			gen13819 := Call(__e, __e.Global(symcons), A, gen13818)

			__e.TailApply(__e.Global(symunify_b), V3286, gen13819, V3287, V3288)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-arity"), gen13820)

		gen13834 := MakeNative(func(__e Evaluator) {
			V3296 := __e.Get(1)
			_ = V3296
			V3297 := __e.Get(2)
			_ = V3297
			V3298 := __e.Get(3)
			_ = V3298
			gen13821 := Call(__e, __e.Global(symshen_4newpv), V3297)

			A := gen13821
			Call(__e, __e.Global(symshen_4incinfs))
			gen13822 := Call(__e, __e.Global(symcons), A, Nil)

			gen13823 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13822)

			gen13824 := Call(__e, __e.Global(symcons), gen13823, Nil)

			gen13825 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13824)

			gen13826 := Call(__e, __e.Global(symcons), A, Nil)

			gen13827 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13826)

			gen13828 := Call(__e, __e.Global(symcons), gen13827, Nil)

			gen13829 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13828)

			gen13830 := Call(__e, __e.Global(symcons), gen13825, gen13829)

			gen13831 := Call(__e, __e.Global(symcons), gen13830, Nil)

			gen13832 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13831)

			gen13833 := Call(__e, __e.Global(symcons), A, gen13832)

			__e.TailApply(__e.Global(symunify_b), V3296, gen13833, V3297, V3298)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-assoc"), gen13834)

		gen13839 := MakeNative(func(__e Evaluator) {
			V3306 := __e.Get(1)
			_ = V3306
			V3307 := __e.Get(2)
			_ = V3307
			V3308 := __e.Get(3)
			_ = V3308
			gen13835 := Call(__e, __e.Global(symshen_4newpv), V3307)

			A := gen13835
			Call(__e, __e.Global(symshen_4incinfs))
			gen13836 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13837 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13836)

			gen13838 := Call(__e, __e.Global(symcons), A, gen13837)

			__e.TailApply(__e.Global(symunify_b), V3306, gen13838, V3307, V3308)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-boolean?"), gen13839)

		gen13843 := MakeNative(func(__e Evaluator) {
			V3316 := __e.Get(1)
			_ = V3316
			V3317 := __e.Get(2)
			_ = V3317
			V3318 := __e.Get(3)
			_ = V3318
			Call(__e, __e.Global(symshen_4incinfs))
			gen13840 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13841 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13840)

			gen13842 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen13841)

			__e.TailApply(__e.Global(symunify_b), V3316, gen13842, V3317, V3318)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-bound?"), gen13843)

		gen13847 := MakeNative(func(__e Evaluator) {
			V3326 := __e.Get(1)
			_ = V3326
			V3327 := __e.Get(2)
			_ = V3327
			V3328 := __e.Get(3)
			_ = V3328
			Call(__e, __e.Global(symshen_4incinfs))
			gen13844 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen13845 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13844)

			gen13846 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen13845)

			__e.TailApply(__e.Global(symunify_b), V3326, gen13846, V3327, V3328)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-cd"), gen13847)

		gen13857 := MakeNative(func(__e Evaluator) {
			V3336 := __e.Get(1)
			_ = V3336
			V3337 := __e.Get(2)
			_ = V3337
			V3338 := __e.Get(3)
			_ = V3338
			gen13848 := Call(__e, __e.Global(symshen_4newpv), V3337)

			A := gen13848
			gen13849 := Call(__e, __e.Global(symshen_4newpv), V3337)

			B := gen13849
			Call(__e, __e.Global(symshen_4incinfs))
			gen13850 := Call(__e, __e.Global(symcons), A, Nil)

			gen13851 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen13850)

			gen13852 := Call(__e, __e.Global(symcons), B, Nil)

			gen13853 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13852)

			gen13854 := Call(__e, __e.Global(symcons), gen13853, Nil)

			gen13855 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13854)

			gen13856 := Call(__e, __e.Global(symcons), gen13851, gen13855)

			__e.TailApply(__e.Global(symunify_b), V3336, gen13856, V3337, V3338)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-close"), gen13857)

		gen13864 := MakeNative(func(__e Evaluator) {
			V3346 := __e.Get(1)
			_ = V3346
			V3347 := __e.Get(2)
			_ = V3347
			V3348 := __e.Get(3)
			_ = V3348
			Call(__e, __e.Global(symshen_4incinfs))
			gen13858 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen13859 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13858)

			gen13860 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen13859)

			gen13861 := Call(__e, __e.Global(symcons), gen13860, Nil)

			gen13862 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13861)

			gen13863 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen13862)

			__e.TailApply(__e.Global(symunify_b), V3346, gen13863, V3347, V3348)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-cn"), gen13864)

		gen13882 := MakeNative(func(__e Evaluator) {
			V3356 := __e.Get(1)
			_ = V3356
			V3357 := __e.Get(2)
			_ = V3357
			V3358 := __e.Get(3)
			_ = V3358
			gen13865 := Call(__e, __e.Global(symshen_4newpv), V3357)

			A := gen13865
			gen13866 := Call(__e, __e.Global(symshen_4newpv), V3357)

			B := gen13866
			Call(__e, __e.Global(symshen_4incinfs))
			gen13867 := Call(__e, __e.Global(symcons), B, Nil)

			gen13868 := Call(__e, __e.Global(symcons), MakeSymbol("shen.==>"), gen13867)

			gen13869 := Call(__e, __e.Global(symcons), A, gen13868)

			gen13870 := Call(__e, __e.Global(symcons), B, Nil)

			gen13871 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13870)

			gen13872 := Call(__e, __e.Global(symcons), A, gen13871)

			gen13873 := Call(__e, __e.Global(symcons), B, Nil)

			gen13874 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13873)

			gen13875 := Call(__e, __e.Global(symcons), gen13872, gen13874)

			gen13876 := Call(__e, __e.Global(symcons), gen13875, Nil)

			gen13877 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13876)

			gen13878 := Call(__e, __e.Global(symcons), A, gen13877)

			gen13879 := Call(__e, __e.Global(symcons), gen13878, Nil)

			gen13880 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13879)

			gen13881 := Call(__e, __e.Global(symcons), gen13869, gen13880)

			__e.TailApply(__e.Global(symunify_b), V3356, gen13881, V3357, V3358)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-compile"), gen13882)

		gen13887 := MakeNative(func(__e Evaluator) {
			V3366 := __e.Get(1)
			_ = V3366
			V3367 := __e.Get(2)
			_ = V3367
			V3368 := __e.Get(3)
			_ = V3368
			gen13883 := Call(__e, __e.Global(symshen_4newpv), V3367)

			A := gen13883
			Call(__e, __e.Global(symshen_4incinfs))
			gen13884 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13885 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13884)

			gen13886 := Call(__e, __e.Global(symcons), A, gen13885)

			__e.TailApply(__e.Global(symunify_b), V3366, gen13886, V3367, V3368)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-cons?"), gen13887)

		gen13896 := MakeNative(func(__e Evaluator) {
			V3376 := __e.Get(1)
			_ = V3376
			V3377 := __e.Get(2)
			_ = V3377
			V3378 := __e.Get(3)
			_ = V3378
			gen13888 := Call(__e, __e.Global(symshen_4newpv), V3377)

			A := gen13888
			gen13889 := Call(__e, __e.Global(symshen_4newpv), V3377)

			B := gen13889
			Call(__e, __e.Global(symshen_4incinfs))
			gen13890 := Call(__e, __e.Global(symcons), B, Nil)

			gen13891 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13890)

			gen13892 := Call(__e, __e.Global(symcons), A, gen13891)

			gen13893 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen13894 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13893)

			gen13895 := Call(__e, __e.Global(symcons), gen13892, gen13894)

			__e.TailApply(__e.Global(symunify_b), V3376, gen13895, V3377, V3378)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-destroy"), gen13896)

		gen13910 := MakeNative(func(__e Evaluator) {
			V3386 := __e.Get(1)
			_ = V3386
			V3387 := __e.Get(2)
			_ = V3387
			V3388 := __e.Get(3)
			_ = V3388
			gen13897 := Call(__e, __e.Global(symshen_4newpv), V3387)

			A := gen13897
			Call(__e, __e.Global(symshen_4incinfs))
			gen13898 := Call(__e, __e.Global(symcons), A, Nil)

			gen13899 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13898)

			gen13900 := Call(__e, __e.Global(symcons), A, Nil)

			gen13901 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13900)

			gen13902 := Call(__e, __e.Global(symcons), A, Nil)

			gen13903 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13902)

			gen13904 := Call(__e, __e.Global(symcons), gen13903, Nil)

			gen13905 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13904)

			gen13906 := Call(__e, __e.Global(symcons), gen13901, gen13905)

			gen13907 := Call(__e, __e.Global(symcons), gen13906, Nil)

			gen13908 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13907)

			gen13909 := Call(__e, __e.Global(symcons), gen13899, gen13908)

			__e.TailApply(__e.Global(symunify_b), V3386, gen13909, V3387, V3388)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-difference"), gen13910)

		gen13919 := MakeNative(func(__e Evaluator) {
			V3396 := __e.Get(1)
			_ = V3396
			V3397 := __e.Get(2)
			_ = V3397
			V3398 := __e.Get(3)
			_ = V3398
			gen13911 := Call(__e, __e.Global(symshen_4newpv), V3397)

			A := gen13911
			gen13912 := Call(__e, __e.Global(symshen_4newpv), V3397)

			B := gen13912
			Call(__e, __e.Global(symshen_4incinfs))
			gen13913 := Call(__e, __e.Global(symcons), B, Nil)

			gen13914 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13913)

			gen13915 := Call(__e, __e.Global(symcons), B, gen13914)

			gen13916 := Call(__e, __e.Global(symcons), gen13915, Nil)

			gen13917 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13916)

			gen13918 := Call(__e, __e.Global(symcons), A, gen13917)

			__e.TailApply(__e.Global(symunify_b), V3396, gen13918, V3397, V3398)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-do"), gen13919)

		gen13929 := MakeNative(func(__e Evaluator) {
			V3406 := __e.Get(1)
			_ = V3406
			V3407 := __e.Get(2)
			_ = V3407
			V3408 := __e.Get(3)
			_ = V3408
			gen13920 := Call(__e, __e.Global(symshen_4newpv), V3407)

			A := gen13920
			gen13921 := Call(__e, __e.Global(symshen_4newpv), V3407)

			B := gen13921
			Call(__e, __e.Global(symshen_4incinfs))
			gen13922 := Call(__e, __e.Global(symcons), A, Nil)

			gen13923 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13922)

			gen13924 := Call(__e, __e.Global(symcons), B, Nil)

			gen13925 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13924)

			gen13926 := Call(__e, __e.Global(symcons), gen13925, Nil)

			gen13927 := Call(__e, __e.Global(symcons), MakeSymbol("shen.==>"), gen13926)

			gen13928 := Call(__e, __e.Global(symcons), gen13923, gen13927)

			__e.TailApply(__e.Global(symunify_b), V3406, gen13928, V3407, V3408)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-<e>"), gen13929)

		gen13938 := MakeNative(func(__e Evaluator) {
			V3416 := __e.Get(1)
			_ = V3416
			V3417 := __e.Get(2)
			_ = V3417
			V3418 := __e.Get(3)
			_ = V3418
			gen13930 := Call(__e, __e.Global(symshen_4newpv), V3417)

			A := gen13930
			Call(__e, __e.Global(symshen_4incinfs))
			gen13931 := Call(__e, __e.Global(symcons), A, Nil)

			gen13932 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13931)

			gen13933 := Call(__e, __e.Global(symcons), A, Nil)

			gen13934 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13933)

			gen13935 := Call(__e, __e.Global(symcons), gen13934, Nil)

			gen13936 := Call(__e, __e.Global(symcons), MakeSymbol("shen.==>"), gen13935)

			gen13937 := Call(__e, __e.Global(symcons), gen13932, gen13936)

			__e.TailApply(__e.Global(symunify_b), V3416, gen13937, V3417, V3418)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-<!>"), gen13938)

		gen13948 := MakeNative(func(__e Evaluator) {
			V3426 := __e.Get(1)
			_ = V3426
			V3427 := __e.Get(2)
			_ = V3427
			V3428 := __e.Get(3)
			_ = V3428
			gen13939 := Call(__e, __e.Global(symshen_4newpv), V3427)

			A := gen13939
			Call(__e, __e.Global(symshen_4incinfs))
			gen13940 := Call(__e, __e.Global(symcons), A, Nil)

			gen13941 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13940)

			gen13942 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13943 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13942)

			gen13944 := Call(__e, __e.Global(symcons), gen13941, gen13943)

			gen13945 := Call(__e, __e.Global(symcons), gen13944, Nil)

			gen13946 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13945)

			gen13947 := Call(__e, __e.Global(symcons), A, gen13946)

			__e.TailApply(__e.Global(symunify_b), V3426, gen13947, V3427, V3428)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-element?"), gen13948)

		gen13953 := MakeNative(func(__e Evaluator) {
			V3436 := __e.Get(1)
			_ = V3436
			V3437 := __e.Get(2)
			_ = V3437
			V3438 := __e.Get(3)
			_ = V3438
			gen13949 := Call(__e, __e.Global(symshen_4newpv), V3437)

			A := gen13949
			Call(__e, __e.Global(symshen_4incinfs))
			gen13950 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13951 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13950)

			gen13952 := Call(__e, __e.Global(symcons), A, gen13951)

			__e.TailApply(__e.Global(symunify_b), V3436, gen13952, V3437, V3438)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-empty?"), gen13953)

		gen13957 := MakeNative(func(__e Evaluator) {
			V3446 := __e.Get(1)
			_ = V3446
			V3447 := __e.Get(2)
			_ = V3447
			V3448 := __e.Get(3)
			_ = V3448
			Call(__e, __e.Global(symshen_4incinfs))
			gen13954 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13955 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13954)

			gen13956 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen13955)

			__e.TailApply(__e.Global(symunify_b), V3446, gen13956, V3447, V3448)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-enable-type-theory"), gen13957)

		gen13963 := MakeNative(func(__e Evaluator) {
			V3456 := __e.Get(1)
			_ = V3456
			V3457 := __e.Get(2)
			_ = V3457
			V3458 := __e.Get(3)
			_ = V3458
			Call(__e, __e.Global(symshen_4incinfs))
			gen13958 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen13959 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13958)

			gen13960 := Call(__e, __e.Global(symcons), gen13959, Nil)

			gen13961 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13960)

			gen13962 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen13961)

			__e.TailApply(__e.Global(symunify_b), V3456, gen13962, V3457, V3458)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-external"), gen13963)

		gen13967 := MakeNative(func(__e Evaluator) {
			V3466 := __e.Get(1)
			_ = V3466
			V3467 := __e.Get(2)
			_ = V3467
			V3468 := __e.Get(3)
			_ = V3468
			Call(__e, __e.Global(symshen_4incinfs))
			gen13964 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen13965 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13964)

			gen13966 := Call(__e, __e.Global(symcons), MakeSymbol("exception"), gen13965)

			__e.TailApply(__e.Global(symunify_b), V3466, gen13966, V3467, V3468)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-error-to-string"), gen13967)

		gen13974 := MakeNative(func(__e Evaluator) {
			V3476 := __e.Get(1)
			_ = V3476
			V3477 := __e.Get(2)
			_ = V3477
			V3478 := __e.Get(3)
			_ = V3478
			gen13968 := Call(__e, __e.Global(symshen_4newpv), V3477)

			A := gen13968
			Call(__e, __e.Global(symshen_4incinfs))
			gen13969 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen13970 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen13969)

			gen13971 := Call(__e, __e.Global(symcons), gen13970, Nil)

			gen13972 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13971)

			gen13973 := Call(__e, __e.Global(symcons), A, gen13972)

			__e.TailApply(__e.Global(symunify_b), V3476, gen13973, V3477, V3478)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-explode"), gen13974)

		gen13977 := MakeNative(func(__e Evaluator) {
			V3486 := __e.Get(1)
			_ = V3486
			V3487 := __e.Get(2)
			_ = V3487
			V3488 := __e.Get(3)
			_ = V3488
			Call(__e, __e.Global(symshen_4incinfs))
			gen13975 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen13976 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13975)

			__e.TailApply(__e.Global(symunify_b), V3486, gen13976, V3487, V3488)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-fail"), gen13977)

		gen13987 := MakeNative(func(__e Evaluator) {
			V3496 := __e.Get(1)
			_ = V3496
			V3497 := __e.Get(2)
			_ = V3497
			V3498 := __e.Get(3)
			_ = V3498
			Call(__e, __e.Global(symshen_4incinfs))
			gen13978 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen13979 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13978)

			gen13980 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen13979)

			gen13981 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen13982 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13981)

			gen13983 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen13982)

			gen13984 := Call(__e, __e.Global(symcons), gen13983, Nil)

			gen13985 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13984)

			gen13986 := Call(__e, __e.Global(symcons), gen13980, gen13985)

			__e.TailApply(__e.Global(symunify_b), V3496, gen13986, V3497, V3498)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-fail-if"), gen13987)

		gen13998 := MakeNative(func(__e Evaluator) {
			V3506 := __e.Get(1)
			_ = V3506
			V3507 := __e.Get(2)
			_ = V3507
			V3508 := __e.Get(3)
			_ = V3508
			gen13988 := Call(__e, __e.Global(symshen_4newpv), V3507)

			A := gen13988
			Call(__e, __e.Global(symshen_4incinfs))
			gen13989 := Call(__e, __e.Global(symcons), A, Nil)

			gen13990 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13989)

			gen13991 := Call(__e, __e.Global(symcons), A, gen13990)

			gen13992 := Call(__e, __e.Global(symcons), A, Nil)

			gen13993 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13992)

			gen13994 := Call(__e, __e.Global(symcons), A, gen13993)

			gen13995 := Call(__e, __e.Global(symcons), gen13994, Nil)

			gen13996 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen13995)

			gen13997 := Call(__e, __e.Global(symcons), gen13991, gen13996)

			__e.TailApply(__e.Global(symunify_b), V3506, gen13997, V3507, V3508)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-fix"), gen13998)

		gen14005 := MakeNative(func(__e Evaluator) {
			V3516 := __e.Get(1)
			_ = V3516
			V3517 := __e.Get(2)
			_ = V3517
			V3518 := __e.Get(3)
			_ = V3518
			gen13999 := Call(__e, __e.Global(symshen_4newpv), V3517)

			A := gen13999
			Call(__e, __e.Global(symshen_4incinfs))
			gen14000 := Call(__e, __e.Global(symcons), A, Nil)

			gen14001 := Call(__e, __e.Global(symcons), MakeSymbol("lazy"), gen14000)

			gen14002 := Call(__e, __e.Global(symcons), gen14001, Nil)

			gen14003 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14002)

			gen14004 := Call(__e, __e.Global(symcons), A, gen14003)

			__e.TailApply(__e.Global(symunify_b), V3516, gen14004, V3517, V3518)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-freeze"), gen14005)

		gen14014 := MakeNative(func(__e Evaluator) {
			V3526 := __e.Get(1)
			_ = V3526
			V3527 := __e.Get(2)
			_ = V3527
			V3528 := __e.Get(3)
			_ = V3528
			gen14006 := Call(__e, __e.Global(symshen_4newpv), V3527)

			B := gen14006
			gen14007 := Call(__e, __e.Global(symshen_4newpv), V3527)

			A := gen14007
			Call(__e, __e.Global(symshen_4incinfs))
			gen14008 := Call(__e, __e.Global(symcons), B, Nil)

			gen14009 := Call(__e, __e.Global(symcons), MakeSymbol("*"), gen14008)

			gen14010 := Call(__e, __e.Global(symcons), A, gen14009)

			gen14011 := Call(__e, __e.Global(symcons), A, Nil)

			gen14012 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14011)

			gen14013 := Call(__e, __e.Global(symcons), gen14010, gen14012)

			__e.TailApply(__e.Global(symunify_b), V3526, gen14013, V3527, V3528)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-fst"), gen14014)

		gen14026 := MakeNative(func(__e Evaluator) {
			V3536 := __e.Get(1)
			_ = V3536
			V3537 := __e.Get(2)
			_ = V3537
			V3538 := __e.Get(3)
			_ = V3538
			gen14015 := Call(__e, __e.Global(symshen_4newpv), V3537)

			A := gen14015
			gen14016 := Call(__e, __e.Global(symshen_4newpv), V3537)

			B := gen14016
			Call(__e, __e.Global(symshen_4incinfs))
			gen14017 := Call(__e, __e.Global(symcons), B, Nil)

			gen14018 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14017)

			gen14019 := Call(__e, __e.Global(symcons), A, gen14018)

			gen14020 := Call(__e, __e.Global(symcons), B, Nil)

			gen14021 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14020)

			gen14022 := Call(__e, __e.Global(symcons), A, gen14021)

			gen14023 := Call(__e, __e.Global(symcons), gen14022, Nil)

			gen14024 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14023)

			gen14025 := Call(__e, __e.Global(symcons), gen14019, gen14024)

			__e.TailApply(__e.Global(symunify_b), V3536, gen14025, V3537, V3538)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-function"), gen14026)

		gen14030 := MakeNative(func(__e Evaluator) {
			V3546 := __e.Get(1)
			_ = V3546
			V3547 := __e.Get(2)
			_ = V3547
			V3548 := __e.Get(3)
			_ = V3548
			Call(__e, __e.Global(symshen_4incinfs))
			gen14027 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14028 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14027)

			gen14029 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14028)

			__e.TailApply(__e.Global(symunify_b), V3546, gen14029, V3547, V3548)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-gensym"), gen14030)

		gen14040 := MakeNative(func(__e Evaluator) {
			V3556 := __e.Get(1)
			_ = V3556
			V3557 := __e.Get(2)
			_ = V3557
			V3558 := __e.Get(3)
			_ = V3558
			gen14031 := Call(__e, __e.Global(symshen_4newpv), V3557)

			A := gen14031
			Call(__e, __e.Global(symshen_4incinfs))
			gen14032 := Call(__e, __e.Global(symcons), A, Nil)

			gen14033 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14032)

			gen14034 := Call(__e, __e.Global(symcons), A, Nil)

			gen14035 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14034)

			gen14036 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14035)

			gen14037 := Call(__e, __e.Global(symcons), gen14036, Nil)

			gen14038 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14037)

			gen14039 := Call(__e, __e.Global(symcons), gen14033, gen14038)

			__e.TailApply(__e.Global(symunify_b), V3556, gen14039, V3557, V3558)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-<-vector"), gen14040)

		gen14055 := MakeNative(func(__e Evaluator) {
			V3566 := __e.Get(1)
			_ = V3566
			V3567 := __e.Get(2)
			_ = V3567
			V3568 := __e.Get(3)
			_ = V3568
			gen14041 := Call(__e, __e.Global(symshen_4newpv), V3567)

			A := gen14041
			Call(__e, __e.Global(symshen_4incinfs))
			gen14042 := Call(__e, __e.Global(symcons), A, Nil)

			gen14043 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14042)

			gen14044 := Call(__e, __e.Global(symcons), A, Nil)

			gen14045 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14044)

			gen14046 := Call(__e, __e.Global(symcons), gen14045, Nil)

			gen14047 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14046)

			gen14048 := Call(__e, __e.Global(symcons), A, gen14047)

			gen14049 := Call(__e, __e.Global(symcons), gen14048, Nil)

			gen14050 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14049)

			gen14051 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14050)

			gen14052 := Call(__e, __e.Global(symcons), gen14051, Nil)

			gen14053 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14052)

			gen14054 := Call(__e, __e.Global(symcons), gen14043, gen14053)

			__e.TailApply(__e.Global(symunify_b), V3566, gen14054, V3567, V3568)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-vector->"), gen14055)

		gen14062 := MakeNative(func(__e Evaluator) {
			V3576 := __e.Get(1)
			_ = V3576
			V3577 := __e.Get(2)
			_ = V3577
			V3578 := __e.Get(3)
			_ = V3578
			gen14056 := Call(__e, __e.Global(symshen_4newpv), V3577)

			A := gen14056
			Call(__e, __e.Global(symshen_4incinfs))
			gen14057 := Call(__e, __e.Global(symcons), A, Nil)

			gen14058 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14057)

			gen14059 := Call(__e, __e.Global(symcons), gen14058, Nil)

			gen14060 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14059)

			gen14061 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14060)

			__e.TailApply(__e.Global(symunify_b), V3576, gen14061, V3577, V3578)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-vector"), gen14062)

		gen14066 := MakeNative(func(__e Evaluator) {
			V3586 := __e.Get(1)
			_ = V3586
			V3587 := __e.Get(2)
			_ = V3587
			V3588 := __e.Get(3)
			_ = V3588
			Call(__e, __e.Global(symshen_4incinfs))
			gen14063 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14064 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14063)

			gen14065 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14064)

			__e.TailApply(__e.Global(symunify_b), V3586, gen14065, V3587, V3588)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-get-time"), gen14066)

		gen14074 := MakeNative(func(__e Evaluator) {
			V3596 := __e.Get(1)
			_ = V3596
			V3597 := __e.Get(2)
			_ = V3597
			V3598 := __e.Get(3)
			_ = V3598
			gen14067 := Call(__e, __e.Global(symshen_4newpv), V3597)

			A := gen14067
			Call(__e, __e.Global(symshen_4incinfs))
			gen14068 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14069 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14068)

			gen14070 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14069)

			gen14071 := Call(__e, __e.Global(symcons), gen14070, Nil)

			gen14072 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14071)

			gen14073 := Call(__e, __e.Global(symcons), A, gen14072)

			__e.TailApply(__e.Global(symunify_b), V3596, gen14073, V3597, V3598)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-hash"), gen14074)

		gen14081 := MakeNative(func(__e Evaluator) {
			V3606 := __e.Get(1)
			_ = V3606
			V3607 := __e.Get(2)
			_ = V3607
			V3608 := __e.Get(3)
			_ = V3608
			gen14075 := Call(__e, __e.Global(symshen_4newpv), V3607)

			A := gen14075
			Call(__e, __e.Global(symshen_4incinfs))
			gen14076 := Call(__e, __e.Global(symcons), A, Nil)

			gen14077 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14076)

			gen14078 := Call(__e, __e.Global(symcons), A, Nil)

			gen14079 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14078)

			gen14080 := Call(__e, __e.Global(symcons), gen14077, gen14079)

			__e.TailApply(__e.Global(symunify_b), V3606, gen14080, V3607, V3608)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-head"), gen14081)

		gen14088 := MakeNative(func(__e Evaluator) {
			V3616 := __e.Get(1)
			_ = V3616
			V3617 := __e.Get(2)
			_ = V3617
			V3618 := __e.Get(3)
			_ = V3618
			gen14082 := Call(__e, __e.Global(symshen_4newpv), V3617)

			A := gen14082
			Call(__e, __e.Global(symshen_4incinfs))
			gen14083 := Call(__e, __e.Global(symcons), A, Nil)

			gen14084 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14083)

			gen14085 := Call(__e, __e.Global(symcons), A, Nil)

			gen14086 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14085)

			gen14087 := Call(__e, __e.Global(symcons), gen14084, gen14086)

			__e.TailApply(__e.Global(symunify_b), V3616, gen14087, V3617, V3618)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-hdv"), gen14088)

		gen14092 := MakeNative(func(__e Evaluator) {
			V3626 := __e.Get(1)
			_ = V3626
			V3627 := __e.Get(2)
			_ = V3627
			V3628 := __e.Get(3)
			_ = V3628
			Call(__e, __e.Global(symshen_4incinfs))
			gen14089 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14090 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14089)

			gen14091 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14090)

			__e.TailApply(__e.Global(symunify_b), V3626, gen14091, V3627, V3628)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-hdstr"), gen14092)

		gen14103 := MakeNative(func(__e Evaluator) {
			V3636 := __e.Get(1)
			_ = V3636
			V3637 := __e.Get(2)
			_ = V3637
			V3638 := __e.Get(3)
			_ = V3638
			gen14093 := Call(__e, __e.Global(symshen_4newpv), V3637)

			A := gen14093
			Call(__e, __e.Global(symshen_4incinfs))
			gen14094 := Call(__e, __e.Global(symcons), A, Nil)

			gen14095 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14094)

			gen14096 := Call(__e, __e.Global(symcons), A, gen14095)

			gen14097 := Call(__e, __e.Global(symcons), gen14096, Nil)

			gen14098 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14097)

			gen14099 := Call(__e, __e.Global(symcons), A, gen14098)

			gen14100 := Call(__e, __e.Global(symcons), gen14099, Nil)

			gen14101 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14100)

			gen14102 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), gen14101)

			__e.TailApply(__e.Global(symunify_b), V3636, gen14102, V3637, V3638)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-if"), gen14103)

		gen14106 := MakeNative(func(__e Evaluator) {
			V3646 := __e.Get(1)
			_ = V3646
			V3647 := __e.Get(2)
			_ = V3647
			V3648 := __e.Get(3)
			_ = V3648
			Call(__e, __e.Global(symshen_4incinfs))
			gen14104 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14105 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14104)

			__e.TailApply(__e.Global(symunify_b), V3646, gen14105, V3647, V3648)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-it"), gen14106)

		gen14109 := MakeNative(func(__e Evaluator) {
			V3656 := __e.Get(1)
			_ = V3656
			V3657 := __e.Get(2)
			_ = V3657
			V3658 := __e.Get(3)
			_ = V3658
			Call(__e, __e.Global(symshen_4incinfs))
			gen14107 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14108 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14107)

			__e.TailApply(__e.Global(symunify_b), V3656, gen14108, V3657, V3658)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-implementation"), gen14109)

		gen14117 := MakeNative(func(__e Evaluator) {
			V3666 := __e.Get(1)
			_ = V3666
			V3667 := __e.Get(2)
			_ = V3667
			V3668 := __e.Get(3)
			_ = V3668
			Call(__e, __e.Global(symshen_4incinfs))
			gen14110 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14111 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14110)

			gen14112 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14113 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14112)

			gen14114 := Call(__e, __e.Global(symcons), gen14113, Nil)

			gen14115 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14114)

			gen14116 := Call(__e, __e.Global(symcons), gen14111, gen14115)

			__e.TailApply(__e.Global(symunify_b), V3666, gen14116, V3667, V3668)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-include"), gen14117)

		gen14125 := MakeNative(func(__e Evaluator) {
			V3676 := __e.Get(1)
			_ = V3676
			V3677 := __e.Get(2)
			_ = V3677
			V3678 := __e.Get(3)
			_ = V3678
			Call(__e, __e.Global(symshen_4incinfs))
			gen14118 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14119 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14118)

			gen14120 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14121 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14120)

			gen14122 := Call(__e, __e.Global(symcons), gen14121, Nil)

			gen14123 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14122)

			gen14124 := Call(__e, __e.Global(symcons), gen14119, gen14123)

			__e.TailApply(__e.Global(symunify_b), V3676, gen14124, V3677, V3678)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-include-all-but"), gen14125)

		gen14128 := MakeNative(func(__e Evaluator) {
			V3686 := __e.Get(1)
			_ = V3686
			V3687 := __e.Get(2)
			_ = V3687
			V3688 := __e.Get(3)
			_ = V3688
			Call(__e, __e.Global(symshen_4incinfs))
			gen14126 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14127 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14126)

			__e.TailApply(__e.Global(symunify_b), V3686, gen14127, V3687, V3688)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-inferences"), gen14128)

		gen14136 := MakeNative(func(__e Evaluator) {
			V3696 := __e.Get(1)
			_ = V3696
			V3697 := __e.Get(2)
			_ = V3697
			V3698 := __e.Get(3)
			_ = V3698
			gen14129 := Call(__e, __e.Global(symshen_4newpv), V3697)

			A := gen14129
			Call(__e, __e.Global(symshen_4incinfs))
			gen14130 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14131 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14130)

			gen14132 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14131)

			gen14133 := Call(__e, __e.Global(symcons), gen14132, Nil)

			gen14134 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14133)

			gen14135 := Call(__e, __e.Global(symcons), A, gen14134)

			__e.TailApply(__e.Global(symunify_b), V3696, gen14135, V3697, V3698)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-shen.insert"), gen14136)

		gen14141 := MakeNative(func(__e Evaluator) {
			V3706 := __e.Get(1)
			_ = V3706
			V3707 := __e.Get(2)
			_ = V3707
			V3708 := __e.Get(3)
			_ = V3708
			gen14137 := Call(__e, __e.Global(symshen_4newpv), V3707)

			A := gen14137
			Call(__e, __e.Global(symshen_4incinfs))
			gen14138 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14139 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14138)

			gen14140 := Call(__e, __e.Global(symcons), A, gen14139)

			__e.TailApply(__e.Global(symunify_b), V3706, gen14140, V3707, V3708)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-integer?"), gen14141)

		gen14147 := MakeNative(func(__e Evaluator) {
			V3716 := __e.Get(1)
			_ = V3716
			V3717 := __e.Get(2)
			_ = V3717
			V3718 := __e.Get(3)
			_ = V3718
			Call(__e, __e.Global(symshen_4incinfs))
			gen14142 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14143 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14142)

			gen14144 := Call(__e, __e.Global(symcons), gen14143, Nil)

			gen14145 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14144)

			gen14146 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14145)

			__e.TailApply(__e.Global(symunify_b), V3716, gen14146, V3717, V3718)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-internal"), gen14147)

		gen14161 := MakeNative(func(__e Evaluator) {
			V3726 := __e.Get(1)
			_ = V3726
			V3727 := __e.Get(2)
			_ = V3727
			V3728 := __e.Get(3)
			_ = V3728
			gen14148 := Call(__e, __e.Global(symshen_4newpv), V3727)

			A := gen14148
			Call(__e, __e.Global(symshen_4incinfs))
			gen14149 := Call(__e, __e.Global(symcons), A, Nil)

			gen14150 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14149)

			gen14151 := Call(__e, __e.Global(symcons), A, Nil)

			gen14152 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14151)

			gen14153 := Call(__e, __e.Global(symcons), A, Nil)

			gen14154 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14153)

			gen14155 := Call(__e, __e.Global(symcons), gen14154, Nil)

			gen14156 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14155)

			gen14157 := Call(__e, __e.Global(symcons), gen14152, gen14156)

			gen14158 := Call(__e, __e.Global(symcons), gen14157, Nil)

			gen14159 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14158)

			gen14160 := Call(__e, __e.Global(symcons), gen14150, gen14159)

			__e.TailApply(__e.Global(symunify_b), V3726, gen14160, V3727, V3728)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-intersection"), gen14161)

		gen14165 := MakeNative(func(__e Evaluator) {
			V3736 := __e.Get(1)
			_ = V3736
			V3737 := __e.Get(2)
			_ = V3737
			V3738 := __e.Get(3)
			_ = V3738
			gen14162 := Call(__e, __e.Global(symshen_4newpv), V3737)

			A := gen14162
			Call(__e, __e.Global(symshen_4incinfs))
			gen14163 := Call(__e, __e.Global(symcons), A, Nil)

			gen14164 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14163)

			__e.TailApply(__e.Global(symunify_b), V3736, gen14164, V3737, V3738)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-kill"), gen14165)

		gen14168 := MakeNative(func(__e Evaluator) {
			V3746 := __e.Get(1)
			_ = V3746
			V3747 := __e.Get(2)
			_ = V3747
			V3748 := __e.Get(3)
			_ = V3748
			Call(__e, __e.Global(symshen_4incinfs))
			gen14166 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14167 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14166)

			__e.TailApply(__e.Global(symunify_b), V3746, gen14167, V3747, V3748)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-language"), gen14168)

		gen14175 := MakeNative(func(__e Evaluator) {
			V3756 := __e.Get(1)
			_ = V3756
			V3757 := __e.Get(2)
			_ = V3757
			V3758 := __e.Get(3)
			_ = V3758
			gen14169 := Call(__e, __e.Global(symshen_4newpv), V3757)

			A := gen14169
			Call(__e, __e.Global(symshen_4incinfs))
			gen14170 := Call(__e, __e.Global(symcons), A, Nil)

			gen14171 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14170)

			gen14172 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14173 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14172)

			gen14174 := Call(__e, __e.Global(symcons), gen14171, gen14173)

			__e.TailApply(__e.Global(symunify_b), V3756, gen14174, V3757, V3758)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-length"), gen14175)

		gen14182 := MakeNative(func(__e Evaluator) {
			V3766 := __e.Get(1)
			_ = V3766
			V3767 := __e.Get(2)
			_ = V3767
			V3768 := __e.Get(3)
			_ = V3768
			gen14176 := Call(__e, __e.Global(symshen_4newpv), V3767)

			A := gen14176
			Call(__e, __e.Global(symshen_4incinfs))
			gen14177 := Call(__e, __e.Global(symcons), A, Nil)

			gen14178 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14177)

			gen14179 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14180 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14179)

			gen14181 := Call(__e, __e.Global(symcons), gen14178, gen14180)

			__e.TailApply(__e.Global(symunify_b), V3766, gen14181, V3767, V3768)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-limit"), gen14182)

		gen14186 := MakeNative(func(__e Evaluator) {
			V3776 := __e.Get(1)
			_ = V3776
			V3777 := __e.Get(2)
			_ = V3777
			V3778 := __e.Get(3)
			_ = V3778
			Call(__e, __e.Global(symshen_4incinfs))
			gen14183 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14184 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14183)

			gen14185 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14184)

			__e.TailApply(__e.Global(symunify_b), V3776, gen14185, V3777, V3778)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-load"), gen14186)

		gen14202 := MakeNative(func(__e Evaluator) {
			V3786 := __e.Get(1)
			_ = V3786
			V3787 := __e.Get(2)
			_ = V3787
			V3788 := __e.Get(3)
			_ = V3788
			gen14187 := Call(__e, __e.Global(symshen_4newpv), V3787)

			A := gen14187
			gen14188 := Call(__e, __e.Global(symshen_4newpv), V3787)

			B := gen14188
			Call(__e, __e.Global(symshen_4incinfs))
			gen14189 := Call(__e, __e.Global(symcons), B, Nil)

			gen14190 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14189)

			gen14191 := Call(__e, __e.Global(symcons), A, gen14190)

			gen14192 := Call(__e, __e.Global(symcons), A, Nil)

			gen14193 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14192)

			gen14194 := Call(__e, __e.Global(symcons), B, Nil)

			gen14195 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14194)

			gen14196 := Call(__e, __e.Global(symcons), gen14195, Nil)

			gen14197 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14196)

			gen14198 := Call(__e, __e.Global(symcons), gen14193, gen14197)

			gen14199 := Call(__e, __e.Global(symcons), gen14198, Nil)

			gen14200 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14199)

			gen14201 := Call(__e, __e.Global(symcons), gen14191, gen14200)

			__e.TailApply(__e.Global(symunify_b), V3786, gen14201, V3787, V3788)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-map"), gen14202)

		gen14220 := MakeNative(func(__e Evaluator) {
			V3796 := __e.Get(1)
			_ = V3796
			V3797 := __e.Get(2)
			_ = V3797
			V3798 := __e.Get(3)
			_ = V3798
			gen14203 := Call(__e, __e.Global(symshen_4newpv), V3797)

			A := gen14203
			gen14204 := Call(__e, __e.Global(symshen_4newpv), V3797)

			B := gen14204
			Call(__e, __e.Global(symshen_4incinfs))
			gen14205 := Call(__e, __e.Global(symcons), B, Nil)

			gen14206 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14205)

			gen14207 := Call(__e, __e.Global(symcons), gen14206, Nil)

			gen14208 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14207)

			gen14209 := Call(__e, __e.Global(symcons), A, gen14208)

			gen14210 := Call(__e, __e.Global(symcons), A, Nil)

			gen14211 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14210)

			gen14212 := Call(__e, __e.Global(symcons), B, Nil)

			gen14213 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14212)

			gen14214 := Call(__e, __e.Global(symcons), gen14213, Nil)

			gen14215 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14214)

			gen14216 := Call(__e, __e.Global(symcons), gen14211, gen14215)

			gen14217 := Call(__e, __e.Global(symcons), gen14216, Nil)

			gen14218 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14217)

			gen14219 := Call(__e, __e.Global(symcons), gen14209, gen14218)

			__e.TailApply(__e.Global(symunify_b), V3796, gen14219, V3797, V3798)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-mapcan"), gen14220)

		gen14224 := MakeNative(func(__e Evaluator) {
			V3806 := __e.Get(1)
			_ = V3806
			V3807 := __e.Get(2)
			_ = V3807
			V3808 := __e.Get(3)
			_ = V3808
			Call(__e, __e.Global(symshen_4incinfs))
			gen14221 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14222 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14221)

			gen14223 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14222)

			__e.TailApply(__e.Global(symunify_b), V3806, gen14223, V3807, V3808)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-maxinferences"), gen14224)

		gen14228 := MakeNative(func(__e Evaluator) {
			V3816 := __e.Get(1)
			_ = V3816
			V3817 := __e.Get(2)
			_ = V3817
			V3818 := __e.Get(3)
			_ = V3818
			Call(__e, __e.Global(symshen_4incinfs))
			gen14225 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14226 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14225)

			gen14227 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14226)

			__e.TailApply(__e.Global(symunify_b), V3816, gen14227, V3817, V3818)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-n->string"), gen14228)

		gen14232 := MakeNative(func(__e Evaluator) {
			V3826 := __e.Get(1)
			_ = V3826
			V3827 := __e.Get(2)
			_ = V3827
			V3828 := __e.Get(3)
			_ = V3828
			Call(__e, __e.Global(symshen_4incinfs))
			gen14229 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14230 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14229)

			gen14231 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14230)

			__e.TailApply(__e.Global(symunify_b), V3826, gen14231, V3827, V3828)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-nl"), gen14232)

		gen14236 := MakeNative(func(__e Evaluator) {
			V3836 := __e.Get(1)
			_ = V3836
			V3837 := __e.Get(2)
			_ = V3837
			V3838 := __e.Get(3)
			_ = V3838
			Call(__e, __e.Global(symshen_4incinfs))
			gen14233 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14234 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14233)

			gen14235 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), gen14234)

			__e.TailApply(__e.Global(symunify_b), V3836, gen14235, V3837, V3838)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-not"), gen14236)

		gen14246 := MakeNative(func(__e Evaluator) {
			V3846 := __e.Get(1)
			_ = V3846
			V3847 := __e.Get(2)
			_ = V3847
			V3848 := __e.Get(3)
			_ = V3848
			gen14237 := Call(__e, __e.Global(symshen_4newpv), V3847)

			A := gen14237
			Call(__e, __e.Global(symshen_4incinfs))
			gen14238 := Call(__e, __e.Global(symcons), A, Nil)

			gen14239 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14238)

			gen14240 := Call(__e, __e.Global(symcons), A, Nil)

			gen14241 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14240)

			gen14242 := Call(__e, __e.Global(symcons), gen14239, gen14241)

			gen14243 := Call(__e, __e.Global(symcons), gen14242, Nil)

			gen14244 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14243)

			gen14245 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14244)

			__e.TailApply(__e.Global(symunify_b), V3846, gen14245, V3847, V3848)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-nth"), gen14246)

		gen14251 := MakeNative(func(__e Evaluator) {
			V3856 := __e.Get(1)
			_ = V3856
			V3857 := __e.Get(2)
			_ = V3857
			V3858 := __e.Get(3)
			_ = V3858
			gen14247 := Call(__e, __e.Global(symshen_4newpv), V3857)

			A := gen14247
			Call(__e, __e.Global(symshen_4incinfs))
			gen14248 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14249 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14248)

			gen14250 := Call(__e, __e.Global(symcons), A, gen14249)

			__e.TailApply(__e.Global(symunify_b), V3856, gen14250, V3857, V3858)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-number?"), gen14251)

		gen14260 := MakeNative(func(__e Evaluator) {
			V3866 := __e.Get(1)
			_ = V3866
			V3867 := __e.Get(2)
			_ = V3867
			V3868 := __e.Get(3)
			_ = V3868
			gen14252 := Call(__e, __e.Global(symshen_4newpv), V3867)

			A := gen14252
			gen14253 := Call(__e, __e.Global(symshen_4newpv), V3867)

			B := gen14253
			Call(__e, __e.Global(symshen_4incinfs))
			gen14254 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14255 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14254)

			gen14256 := Call(__e, __e.Global(symcons), B, gen14255)

			gen14257 := Call(__e, __e.Global(symcons), gen14256, Nil)

			gen14258 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14257)

			gen14259 := Call(__e, __e.Global(symcons), A, gen14258)

			__e.TailApply(__e.Global(symunify_b), V3866, gen14259, V3867, V3868)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-occurrences"), gen14260)

		gen14264 := MakeNative(func(__e Evaluator) {
			V3876 := __e.Get(1)
			_ = V3876
			V3877 := __e.Get(2)
			_ = V3877
			V3878 := __e.Get(3)
			_ = V3878
			Call(__e, __e.Global(symshen_4incinfs))
			gen14261 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14262 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14261)

			gen14263 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14262)

			__e.TailApply(__e.Global(symunify_b), V3876, gen14263, V3877, V3878)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-occurs-check"), gen14264)

		gen14268 := MakeNative(func(__e Evaluator) {
			V3886 := __e.Get(1)
			_ = V3886
			V3887 := __e.Get(2)
			_ = V3887
			V3888 := __e.Get(3)
			_ = V3888
			Call(__e, __e.Global(symshen_4incinfs))
			gen14265 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14266 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14265)

			gen14267 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14266)

			__e.TailApply(__e.Global(symunify_b), V3886, gen14267, V3887, V3888)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-optimise"), gen14268)

		gen14275 := MakeNative(func(__e Evaluator) {
			V3896 := __e.Get(1)
			_ = V3896
			V3897 := __e.Get(2)
			_ = V3897
			V3898 := __e.Get(3)
			_ = V3898
			Call(__e, __e.Global(symshen_4incinfs))
			gen14269 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14270 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14269)

			gen14271 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), gen14270)

			gen14272 := Call(__e, __e.Global(symcons), gen14271, Nil)

			gen14273 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14272)

			gen14274 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), gen14273)

			__e.TailApply(__e.Global(symunify_b), V3896, gen14274, V3897, V3898)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-or"), gen14275)

		gen14278 := MakeNative(func(__e Evaluator) {
			V3906 := __e.Get(1)
			_ = V3906
			V3907 := __e.Get(2)
			_ = V3907
			V3908 := __e.Get(3)
			_ = V3908
			Call(__e, __e.Global(symshen_4incinfs))
			gen14276 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14277 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14276)

			__e.TailApply(__e.Global(symunify_b), V3906, gen14277, V3907, V3908)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-os"), gen14278)

		gen14282 := MakeNative(func(__e Evaluator) {
			V3916 := __e.Get(1)
			_ = V3916
			V3917 := __e.Get(2)
			_ = V3917
			V3918 := __e.Get(3)
			_ = V3918
			Call(__e, __e.Global(symshen_4incinfs))
			gen14279 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14280 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14279)

			gen14281 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14280)

			__e.TailApply(__e.Global(symunify_b), V3916, gen14281, V3917, V3918)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-package?"), gen14282)

		gen14285 := MakeNative(func(__e Evaluator) {
			V3926 := __e.Get(1)
			_ = V3926
			V3927 := __e.Get(2)
			_ = V3927
			V3928 := __e.Get(3)
			_ = V3928
			Call(__e, __e.Global(symshen_4incinfs))
			gen14283 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14284 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14283)

			__e.TailApply(__e.Global(symunify_b), V3926, gen14284, V3927, V3928)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-port"), gen14285)

		gen14288 := MakeNative(func(__e Evaluator) {
			V3936 := __e.Get(1)
			_ = V3936
			V3937 := __e.Get(2)
			_ = V3937
			V3938 := __e.Get(3)
			_ = V3938
			Call(__e, __e.Global(symshen_4incinfs))
			gen14286 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14287 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14286)

			__e.TailApply(__e.Global(symunify_b), V3936, gen14287, V3937, V3938)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-porters"), gen14288)

		gen14295 := MakeNative(func(__e Evaluator) {
			V3946 := __e.Get(1)
			_ = V3946
			V3947 := __e.Get(2)
			_ = V3947
			V3948 := __e.Get(3)
			_ = V3948
			Call(__e, __e.Global(symshen_4incinfs))
			gen14289 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14290 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14289)

			gen14291 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14290)

			gen14292 := Call(__e, __e.Global(symcons), gen14291, Nil)

			gen14293 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14292)

			gen14294 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14293)

			__e.TailApply(__e.Global(symunify_b), V3946, gen14294, V3947, V3948)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-pos"), gen14295)

		gen14304 := MakeNative(func(__e Evaluator) {
			V3956 := __e.Get(1)
			_ = V3956
			V3957 := __e.Get(2)
			_ = V3957
			V3958 := __e.Get(3)
			_ = V3958
			Call(__e, __e.Global(symshen_4incinfs))
			gen14296 := Call(__e, __e.Global(symcons), MakeSymbol("out"), Nil)

			gen14297 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14296)

			gen14298 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14299 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14298)

			gen14300 := Call(__e, __e.Global(symcons), gen14297, gen14299)

			gen14301 := Call(__e, __e.Global(symcons), gen14300, Nil)

			gen14302 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14301)

			gen14303 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14302)

			__e.TailApply(__e.Global(symunify_b), V3956, gen14303, V3957, V3958)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-pr"), gen14304)

		gen14309 := MakeNative(func(__e Evaluator) {
			V3966 := __e.Get(1)
			_ = V3966
			V3967 := __e.Get(2)
			_ = V3967
			V3968 := __e.Get(3)
			_ = V3968
			gen14305 := Call(__e, __e.Global(symshen_4newpv), V3967)

			A := gen14305
			Call(__e, __e.Global(symshen_4incinfs))
			gen14306 := Call(__e, __e.Global(symcons), A, Nil)

			gen14307 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14306)

			gen14308 := Call(__e, __e.Global(symcons), A, gen14307)

			__e.TailApply(__e.Global(symunify_b), V3966, gen14308, V3967, V3968)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-print"), gen14309)

		gen14321 := MakeNative(func(__e Evaluator) {
			V3976 := __e.Get(1)
			_ = V3976
			V3977 := __e.Get(2)
			_ = V3977
			V3978 := __e.Get(3)
			_ = V3978
			gen14310 := Call(__e, __e.Global(symshen_4newpv), V3977)

			A := gen14310
			gen14311 := Call(__e, __e.Global(symshen_4newpv), V3977)

			B := gen14311
			Call(__e, __e.Global(symshen_4incinfs))
			gen14312 := Call(__e, __e.Global(symcons), B, Nil)

			gen14313 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14312)

			gen14314 := Call(__e, __e.Global(symcons), A, gen14313)

			gen14315 := Call(__e, __e.Global(symcons), B, Nil)

			gen14316 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14315)

			gen14317 := Call(__e, __e.Global(symcons), A, gen14316)

			gen14318 := Call(__e, __e.Global(symcons), gen14317, Nil)

			gen14319 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14318)

			gen14320 := Call(__e, __e.Global(symcons), gen14314, gen14319)

			__e.TailApply(__e.Global(symunify_b), V3976, gen14320, V3977, V3978)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-profile"), gen14321)

		gen14329 := MakeNative(func(__e Evaluator) {
			V3986 := __e.Get(1)
			_ = V3986
			V3987 := __e.Get(2)
			_ = V3987
			V3988 := __e.Get(3)
			_ = V3988
			Call(__e, __e.Global(symshen_4incinfs))
			gen14322 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14323 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14322)

			gen14324 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14325 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14324)

			gen14326 := Call(__e, __e.Global(symcons), gen14325, Nil)

			gen14327 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14326)

			gen14328 := Call(__e, __e.Global(symcons), gen14323, gen14327)

			__e.TailApply(__e.Global(symunify_b), V3986, gen14328, V3987, V3988)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-preclude"), gen14329)

		gen14333 := MakeNative(func(__e Evaluator) {
			V3996 := __e.Get(1)
			_ = V3996
			V3997 := __e.Get(2)
			_ = V3997
			V3998 := __e.Get(3)
			_ = V3998
			Call(__e, __e.Global(symshen_4incinfs))
			gen14330 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14331 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14330)

			gen14332 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14331)

			__e.TailApply(__e.Global(symunify_b), V3996, gen14332, V3997, V3998)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-shen.proc-nl"), gen14333)

		gen14348 := MakeNative(func(__e Evaluator) {
			V4006 := __e.Get(1)
			_ = V4006
			V4007 := __e.Get(2)
			_ = V4007
			V4008 := __e.Get(3)
			_ = V4008
			gen14334 := Call(__e, __e.Global(symshen_4newpv), V4007)

			A := gen14334
			gen14335 := Call(__e, __e.Global(symshen_4newpv), V4007)

			B := gen14335
			Call(__e, __e.Global(symshen_4incinfs))
			gen14336 := Call(__e, __e.Global(symcons), B, Nil)

			gen14337 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14336)

			gen14338 := Call(__e, __e.Global(symcons), A, gen14337)

			gen14339 := Call(__e, __e.Global(symcons), B, Nil)

			gen14340 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14339)

			gen14341 := Call(__e, __e.Global(symcons), A, gen14340)

			gen14342 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14343 := Call(__e, __e.Global(symcons), MakeSymbol("*"), gen14342)

			gen14344 := Call(__e, __e.Global(symcons), gen14341, gen14343)

			gen14345 := Call(__e, __e.Global(symcons), gen14344, Nil)

			gen14346 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14345)

			gen14347 := Call(__e, __e.Global(symcons), gen14338, gen14346)

			__e.TailApply(__e.Global(symunify_b), V4006, gen14347, V4007, V4008)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-profile-results"), gen14348)

		gen14352 := MakeNative(func(__e Evaluator) {
			V4016 := __e.Get(1)
			_ = V4016
			V4017 := __e.Get(2)
			_ = V4017
			V4018 := __e.Get(3)
			_ = V4018
			Call(__e, __e.Global(symshen_4incinfs))
			gen14349 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14350 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14349)

			gen14351 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14350)

			__e.TailApply(__e.Global(symunify_b), V4016, gen14351, V4017, V4018)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-protect"), gen14352)

		gen14360 := MakeNative(func(__e Evaluator) {
			V4026 := __e.Get(1)
			_ = V4026
			V4027 := __e.Get(2)
			_ = V4027
			V4028 := __e.Get(3)
			_ = V4028
			Call(__e, __e.Global(symshen_4incinfs))
			gen14353 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14354 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14353)

			gen14355 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14356 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14355)

			gen14357 := Call(__e, __e.Global(symcons), gen14356, Nil)

			gen14358 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14357)

			gen14359 := Call(__e, __e.Global(symcons), gen14354, gen14358)

			__e.TailApply(__e.Global(symunify_b), V4026, gen14359, V4027, V4028)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-preclude-all-but"), gen14360)

		gen14369 := MakeNative(func(__e Evaluator) {
			V4036 := __e.Get(1)
			_ = V4036
			V4037 := __e.Get(2)
			_ = V4037
			V4038 := __e.Get(3)
			_ = V4038
			Call(__e, __e.Global(symshen_4incinfs))
			gen14361 := Call(__e, __e.Global(symcons), MakeSymbol("out"), Nil)

			gen14362 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14361)

			gen14363 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14364 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14363)

			gen14365 := Call(__e, __e.Global(symcons), gen14362, gen14364)

			gen14366 := Call(__e, __e.Global(symcons), gen14365, Nil)

			gen14367 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14366)

			gen14368 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14367)

			__e.TailApply(__e.Global(symunify_b), V4036, gen14368, V4037, V4038)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-shen.prhush"), gen14369)

		gen14375 := MakeNative(func(__e Evaluator) {
			V4046 := __e.Get(1)
			_ = V4046
			V4047 := __e.Get(2)
			_ = V4047
			V4048 := __e.Get(3)
			_ = V4048
			Call(__e, __e.Global(symshen_4incinfs))
			gen14370 := Call(__e, __e.Global(symcons), MakeSymbol("unit"), Nil)

			gen14371 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14370)

			gen14372 := Call(__e, __e.Global(symcons), gen14371, Nil)

			gen14373 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14372)

			gen14374 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14373)

			__e.TailApply(__e.Global(symunify_b), V4046, gen14374, V4047, V4048)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-ps"), gen14375)

		gen14381 := MakeNative(func(__e Evaluator) {
			V4056 := __e.Get(1)
			_ = V4056
			V4057 := __e.Get(2)
			_ = V4057
			V4058 := __e.Get(3)
			_ = V4058
			Call(__e, __e.Global(symshen_4incinfs))
			gen14376 := Call(__e, __e.Global(symcons), MakeSymbol("in"), Nil)

			gen14377 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14376)

			gen14378 := Call(__e, __e.Global(symcons), MakeSymbol("unit"), Nil)

			gen14379 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14378)

			gen14380 := Call(__e, __e.Global(symcons), gen14377, gen14379)

			__e.TailApply(__e.Global(symunify_b), V4056, gen14380, V4057, V4058)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-read"), gen14381)

		gen14387 := MakeNative(func(__e Evaluator) {
			V4066 := __e.Get(1)
			_ = V4066
			V4067 := __e.Get(2)
			_ = V4067
			V4068 := __e.Get(3)
			_ = V4068
			Call(__e, __e.Global(symshen_4incinfs))
			gen14382 := Call(__e, __e.Global(symcons), MakeSymbol("in"), Nil)

			gen14383 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14382)

			gen14384 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14385 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14384)

			gen14386 := Call(__e, __e.Global(symcons), gen14383, gen14385)

			__e.TailApply(__e.Global(symunify_b), V4066, gen14386, V4067, V4068)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-read-byte"), gen14387)

		gen14393 := MakeNative(func(__e Evaluator) {
			V4076 := __e.Get(1)
			_ = V4076
			V4077 := __e.Get(2)
			_ = V4077
			V4078 := __e.Get(3)
			_ = V4078
			Call(__e, __e.Global(symshen_4incinfs))
			gen14388 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14389 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14388)

			gen14390 := Call(__e, __e.Global(symcons), gen14389, Nil)

			gen14391 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14390)

			gen14392 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14391)

			__e.TailApply(__e.Global(symunify_b), V4076, gen14392, V4077, V4078)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-read-file-as-bytelist"), gen14393)

		gen14397 := MakeNative(func(__e Evaluator) {
			V4086 := __e.Get(1)
			_ = V4086
			V4087 := __e.Get(2)
			_ = V4087
			V4088 := __e.Get(3)
			_ = V4088
			Call(__e, __e.Global(symshen_4incinfs))
			gen14394 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14395 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14394)

			gen14396 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14395)

			__e.TailApply(__e.Global(symunify_b), V4086, gen14396, V4087, V4088)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-read-file-as-string"), gen14397)

		gen14403 := MakeNative(func(__e Evaluator) {
			V4096 := __e.Get(1)
			_ = V4096
			V4097 := __e.Get(2)
			_ = V4097
			V4098 := __e.Get(3)
			_ = V4098
			Call(__e, __e.Global(symshen_4incinfs))
			gen14398 := Call(__e, __e.Global(symcons), MakeSymbol("unit"), Nil)

			gen14399 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14398)

			gen14400 := Call(__e, __e.Global(symcons), gen14399, Nil)

			gen14401 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14400)

			gen14402 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14401)

			__e.TailApply(__e.Global(symunify_b), V4096, gen14402, V4097, V4098)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-read-file"), gen14403)

		gen14409 := MakeNative(func(__e Evaluator) {
			V4106 := __e.Get(1)
			_ = V4106
			V4107 := __e.Get(2)
			_ = V4107
			V4108 := __e.Get(3)
			_ = V4108
			Call(__e, __e.Global(symshen_4incinfs))
			gen14404 := Call(__e, __e.Global(symcons), MakeSymbol("unit"), Nil)

			gen14405 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14404)

			gen14406 := Call(__e, __e.Global(symcons), gen14405, Nil)

			gen14407 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14406)

			gen14408 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14407)

			__e.TailApply(__e.Global(symunify_b), V4106, gen14408, V4107, V4108)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-read-from-string"), gen14409)

		gen14412 := MakeNative(func(__e Evaluator) {
			V4116 := __e.Get(1)
			_ = V4116
			V4117 := __e.Get(2)
			_ = V4117
			V4118 := __e.Get(3)
			_ = V4118
			Call(__e, __e.Global(symshen_4incinfs))
			gen14410 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14411 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14410)

			__e.TailApply(__e.Global(symunify_b), V4116, gen14411, V4117, V4118)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-release"), gen14412)

		gen14424 := MakeNative(func(__e Evaluator) {
			V4126 := __e.Get(1)
			_ = V4126
			V4127 := __e.Get(2)
			_ = V4127
			V4128 := __e.Get(3)
			_ = V4128
			gen14413 := Call(__e, __e.Global(symshen_4newpv), V4127)

			A := gen14413
			Call(__e, __e.Global(symshen_4incinfs))
			gen14414 := Call(__e, __e.Global(symcons), A, Nil)

			gen14415 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14414)

			gen14416 := Call(__e, __e.Global(symcons), A, Nil)

			gen14417 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14416)

			gen14418 := Call(__e, __e.Global(symcons), gen14417, Nil)

			gen14419 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14418)

			gen14420 := Call(__e, __e.Global(symcons), gen14415, gen14419)

			gen14421 := Call(__e, __e.Global(symcons), gen14420, Nil)

			gen14422 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14421)

			gen14423 := Call(__e, __e.Global(symcons), A, gen14422)

			__e.TailApply(__e.Global(symunify_b), V4126, gen14423, V4127, V4128)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-remove"), gen14424)

		gen14433 := MakeNative(func(__e Evaluator) {
			V4136 := __e.Get(1)
			_ = V4136
			V4137 := __e.Get(2)
			_ = V4137
			V4138 := __e.Get(3)
			_ = V4138
			gen14425 := Call(__e, __e.Global(symshen_4newpv), V4137)

			A := gen14425
			Call(__e, __e.Global(symshen_4incinfs))
			gen14426 := Call(__e, __e.Global(symcons), A, Nil)

			gen14427 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14426)

			gen14428 := Call(__e, __e.Global(symcons), A, Nil)

			gen14429 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14428)

			gen14430 := Call(__e, __e.Global(symcons), gen14429, Nil)

			gen14431 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14430)

			gen14432 := Call(__e, __e.Global(symcons), gen14427, gen14431)

			__e.TailApply(__e.Global(symunify_b), V4136, gen14432, V4137, V4138)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-reverse"), gen14433)

		gen14438 := MakeNative(func(__e Evaluator) {
			V4146 := __e.Get(1)
			_ = V4146
			V4147 := __e.Get(2)
			_ = V4147
			V4148 := __e.Get(3)
			_ = V4148
			gen14434 := Call(__e, __e.Global(symshen_4newpv), V4147)

			A := gen14434
			Call(__e, __e.Global(symshen_4incinfs))
			gen14435 := Call(__e, __e.Global(symcons), A, Nil)

			gen14436 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14435)

			gen14437 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14436)

			__e.TailApply(__e.Global(symunify_b), V4146, gen14437, V4147, V4148)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-simple-error"), gen14438)

		gen14447 := MakeNative(func(__e Evaluator) {
			V4156 := __e.Get(1)
			_ = V4156
			V4157 := __e.Get(2)
			_ = V4157
			V4158 := __e.Get(3)
			_ = V4158
			gen14439 := Call(__e, __e.Global(symshen_4newpv), V4157)

			A := gen14439
			gen14440 := Call(__e, __e.Global(symshen_4newpv), V4157)

			B := gen14440
			Call(__e, __e.Global(symshen_4incinfs))
			gen14441 := Call(__e, __e.Global(symcons), B, Nil)

			gen14442 := Call(__e, __e.Global(symcons), MakeSymbol("*"), gen14441)

			gen14443 := Call(__e, __e.Global(symcons), A, gen14442)

			gen14444 := Call(__e, __e.Global(symcons), B, Nil)

			gen14445 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14444)

			gen14446 := Call(__e, __e.Global(symcons), gen14443, gen14445)

			__e.TailApply(__e.Global(symunify_b), V4156, gen14446, V4157, V4158)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-snd"), gen14447)

		gen14451 := MakeNative(func(__e Evaluator) {
			V4166 := __e.Get(1)
			_ = V4166
			V4167 := __e.Get(2)
			_ = V4167
			V4168 := __e.Get(3)
			_ = V4168
			Call(__e, __e.Global(symshen_4incinfs))
			gen14448 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14449 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14448)

			gen14450 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14449)

			__e.TailApply(__e.Global(symunify_b), V4166, gen14450, V4167, V4168)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-specialise"), gen14451)

		gen14455 := MakeNative(func(__e Evaluator) {
			V4176 := __e.Get(1)
			_ = V4176
			V4177 := __e.Get(2)
			_ = V4177
			V4178 := __e.Get(3)
			_ = V4178
			Call(__e, __e.Global(symshen_4incinfs))
			gen14452 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14453 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14452)

			gen14454 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14453)

			__e.TailApply(__e.Global(symunify_b), V4176, gen14454, V4177, V4178)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-spy"), gen14455)

		gen14459 := MakeNative(func(__e Evaluator) {
			V4186 := __e.Get(1)
			_ = V4186
			V4187 := __e.Get(2)
			_ = V4187
			V4188 := __e.Get(3)
			_ = V4188
			Call(__e, __e.Global(symshen_4incinfs))
			gen14456 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14457 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14456)

			gen14458 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14457)

			__e.TailApply(__e.Global(symunify_b), V4186, gen14458, V4187, V4188)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-step"), gen14459)

		gen14464 := MakeNative(func(__e Evaluator) {
			V4196 := __e.Get(1)
			_ = V4196
			V4197 := __e.Get(2)
			_ = V4197
			V4198 := __e.Get(3)
			_ = V4198
			Call(__e, __e.Global(symshen_4incinfs))
			gen14460 := Call(__e, __e.Global(symcons), MakeSymbol("in"), Nil)

			gen14461 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14460)

			gen14462 := Call(__e, __e.Global(symcons), gen14461, Nil)

			gen14463 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14462)

			__e.TailApply(__e.Global(symunify_b), V4196, gen14463, V4197, V4198)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-stinput"), gen14464)

		gen14469 := MakeNative(func(__e Evaluator) {
			V4206 := __e.Get(1)
			_ = V4206
			V4207 := __e.Get(2)
			_ = V4207
			V4208 := __e.Get(3)
			_ = V4208
			Call(__e, __e.Global(symshen_4incinfs))
			gen14465 := Call(__e, __e.Global(symcons), MakeSymbol("out"), Nil)

			gen14466 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14465)

			gen14467 := Call(__e, __e.Global(symcons), gen14466, Nil)

			gen14468 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14467)

			__e.TailApply(__e.Global(symunify_b), V4206, gen14468, V4207, V4208)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-sterror"), gen14469)

		gen14474 := MakeNative(func(__e Evaluator) {
			V4216 := __e.Get(1)
			_ = V4216
			V4217 := __e.Get(2)
			_ = V4217
			V4218 := __e.Get(3)
			_ = V4218
			Call(__e, __e.Global(symshen_4incinfs))
			gen14470 := Call(__e, __e.Global(symcons), MakeSymbol("out"), Nil)

			gen14471 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14470)

			gen14472 := Call(__e, __e.Global(symcons), gen14471, Nil)

			gen14473 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14472)

			__e.TailApply(__e.Global(symunify_b), V4216, gen14473, V4217, V4218)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-stoutput"), gen14474)

		gen14479 := MakeNative(func(__e Evaluator) {
			V4226 := __e.Get(1)
			_ = V4226
			V4227 := __e.Get(2)
			_ = V4227
			V4228 := __e.Get(3)
			_ = V4228
			gen14475 := Call(__e, __e.Global(symshen_4newpv), V4227)

			A := gen14475
			Call(__e, __e.Global(symshen_4incinfs))
			gen14476 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14477 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14476)

			gen14478 := Call(__e, __e.Global(symcons), A, gen14477)

			__e.TailApply(__e.Global(symunify_b), V4226, gen14478, V4227, V4228)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-string?"), gen14479)

		gen14484 := MakeNative(func(__e Evaluator) {
			V4236 := __e.Get(1)
			_ = V4236
			V4237 := __e.Get(2)
			_ = V4237
			V4238 := __e.Get(3)
			_ = V4238
			gen14480 := Call(__e, __e.Global(symshen_4newpv), V4237)

			A := gen14480
			Call(__e, __e.Global(symshen_4incinfs))
			gen14481 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14482 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14481)

			gen14483 := Call(__e, __e.Global(symcons), A, gen14482)

			__e.TailApply(__e.Global(symunify_b), V4236, gen14483, V4237, V4238)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-str"), gen14484)

		gen14488 := MakeNative(func(__e Evaluator) {
			V4246 := __e.Get(1)
			_ = V4246
			V4247 := __e.Get(2)
			_ = V4247
			V4248 := __e.Get(3)
			_ = V4248
			Call(__e, __e.Global(symshen_4incinfs))
			gen14485 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14486 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14485)

			gen14487 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14486)

			__e.TailApply(__e.Global(symunify_b), V4246, gen14487, V4247, V4248)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-string->n"), gen14488)

		gen14492 := MakeNative(func(__e Evaluator) {
			V4256 := __e.Get(1)
			_ = V4256
			V4257 := __e.Get(2)
			_ = V4257
			V4258 := __e.Get(3)
			_ = V4258
			Call(__e, __e.Global(symshen_4incinfs))
			gen14489 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14490 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14489)

			gen14491 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14490)

			__e.TailApply(__e.Global(symunify_b), V4256, gen14491, V4257, V4258)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-string->symbol"), gen14492)

		gen14498 := MakeNative(func(__e Evaluator) {
			V4266 := __e.Get(1)
			_ = V4266
			V4267 := __e.Get(2)
			_ = V4267
			V4268 := __e.Get(3)
			_ = V4268
			Call(__e, __e.Global(symshen_4incinfs))
			gen14493 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14494 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14493)

			gen14495 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14496 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14495)

			gen14497 := Call(__e, __e.Global(symcons), gen14494, gen14496)

			__e.TailApply(__e.Global(symunify_b), V4266, gen14497, V4267, V4268)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-sum"), gen14498)

		gen14503 := MakeNative(func(__e Evaluator) {
			V4276 := __e.Get(1)
			_ = V4276
			V4277 := __e.Get(2)
			_ = V4277
			V4278 := __e.Get(3)
			_ = V4278
			gen14499 := Call(__e, __e.Global(symshen_4newpv), V4277)

			A := gen14499
			Call(__e, __e.Global(symshen_4incinfs))
			gen14500 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14501 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14500)

			gen14502 := Call(__e, __e.Global(symcons), A, gen14501)

			__e.TailApply(__e.Global(symunify_b), V4276, gen14502, V4277, V4278)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-symbol?"), gen14503)

		gen14507 := MakeNative(func(__e Evaluator) {
			V4286 := __e.Get(1)
			_ = V4286
			V4287 := __e.Get(2)
			_ = V4287
			V4288 := __e.Get(3)
			_ = V4288
			Call(__e, __e.Global(symshen_4incinfs))
			gen14504 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14505 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14504)

			gen14506 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14505)

			__e.TailApply(__e.Global(symunify_b), V4286, gen14506, V4287, V4288)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-systemf"), gen14507)

		gen14516 := MakeNative(func(__e Evaluator) {
			V4296 := __e.Get(1)
			_ = V4296
			V4297 := __e.Get(2)
			_ = V4297
			V4298 := __e.Get(3)
			_ = V4298
			gen14508 := Call(__e, __e.Global(symshen_4newpv), V4297)

			A := gen14508
			Call(__e, __e.Global(symshen_4incinfs))
			gen14509 := Call(__e, __e.Global(symcons), A, Nil)

			gen14510 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14509)

			gen14511 := Call(__e, __e.Global(symcons), A, Nil)

			gen14512 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14511)

			gen14513 := Call(__e, __e.Global(symcons), gen14512, Nil)

			gen14514 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14513)

			gen14515 := Call(__e, __e.Global(symcons), gen14510, gen14514)

			__e.TailApply(__e.Global(symunify_b), V4296, gen14515, V4297, V4298)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-tail"), gen14516)

		gen14520 := MakeNative(func(__e Evaluator) {
			V4306 := __e.Get(1)
			_ = V4306
			V4307 := __e.Get(2)
			_ = V4307
			V4308 := __e.Get(3)
			_ = V4308
			Call(__e, __e.Global(symshen_4incinfs))
			gen14517 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14518 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14517)

			gen14519 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14518)

			__e.TailApply(__e.Global(symunify_b), V4306, gen14519, V4307, V4308)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-tlstr"), gen14520)

		gen14529 := MakeNative(func(__e Evaluator) {
			V4316 := __e.Get(1)
			_ = V4316
			V4317 := __e.Get(2)
			_ = V4317
			V4318 := __e.Get(3)
			_ = V4318
			gen14521 := Call(__e, __e.Global(symshen_4newpv), V4317)

			A := gen14521
			Call(__e, __e.Global(symshen_4incinfs))
			gen14522 := Call(__e, __e.Global(symcons), A, Nil)

			gen14523 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14522)

			gen14524 := Call(__e, __e.Global(symcons), A, Nil)

			gen14525 := Call(__e, __e.Global(symcons), MakeSymbol("vector"), gen14524)

			gen14526 := Call(__e, __e.Global(symcons), gen14525, Nil)

			gen14527 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14526)

			gen14528 := Call(__e, __e.Global(symcons), gen14523, gen14527)

			__e.TailApply(__e.Global(symunify_b), V4316, gen14528, V4317, V4318)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-tlv"), gen14529)

		gen14533 := MakeNative(func(__e Evaluator) {
			V4326 := __e.Get(1)
			_ = V4326
			V4327 := __e.Get(2)
			_ = V4327
			V4328 := __e.Get(3)
			_ = V4328
			Call(__e, __e.Global(symshen_4incinfs))
			gen14530 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14531 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14530)

			gen14532 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14531)

			__e.TailApply(__e.Global(symunify_b), V4326, gen14532, V4327, V4328)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-tc"), gen14533)

		gen14536 := MakeNative(func(__e Evaluator) {
			V4336 := __e.Get(1)
			_ = V4336
			V4337 := __e.Get(2)
			_ = V4337
			V4338 := __e.Get(3)
			_ = V4338
			Call(__e, __e.Global(symshen_4incinfs))
			gen14534 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14535 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14534)

			__e.TailApply(__e.Global(symunify_b), V4336, gen14535, V4337, V4338)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-tc?"), gen14536)

		gen14543 := MakeNative(func(__e Evaluator) {
			V4346 := __e.Get(1)
			_ = V4346
			V4347 := __e.Get(2)
			_ = V4347
			V4348 := __e.Get(3)
			_ = V4348
			gen14537 := Call(__e, __e.Global(symshen_4newpv), V4347)

			A := gen14537
			Call(__e, __e.Global(symshen_4incinfs))
			gen14538 := Call(__e, __e.Global(symcons), A, Nil)

			gen14539 := Call(__e, __e.Global(symcons), MakeSymbol("lazy"), gen14538)

			gen14540 := Call(__e, __e.Global(symcons), A, Nil)

			gen14541 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14540)

			gen14542 := Call(__e, __e.Global(symcons), gen14539, gen14541)

			__e.TailApply(__e.Global(symunify_b), V4346, gen14542, V4347, V4348)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-thaw"), gen14543)

		gen14547 := MakeNative(func(__e Evaluator) {
			V4356 := __e.Get(1)
			_ = V4356
			V4357 := __e.Get(2)
			_ = V4357
			V4358 := __e.Get(3)
			_ = V4358
			Call(__e, __e.Global(symshen_4incinfs))
			gen14544 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14545 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14544)

			gen14546 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14545)

			__e.TailApply(__e.Global(symunify_b), V4356, gen14546, V4357, V4358)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-track"), gen14547)

		gen14558 := MakeNative(func(__e Evaluator) {
			V4366 := __e.Get(1)
			_ = V4366
			V4367 := __e.Get(2)
			_ = V4367
			V4368 := __e.Get(3)
			_ = V4368
			gen14548 := Call(__e, __e.Global(symshen_4newpv), V4367)

			A := gen14548
			Call(__e, __e.Global(symshen_4incinfs))
			gen14549 := Call(__e, __e.Global(symcons), A, Nil)

			gen14550 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14549)

			gen14551 := Call(__e, __e.Global(symcons), MakeSymbol("exception"), gen14550)

			gen14552 := Call(__e, __e.Global(symcons), A, Nil)

			gen14553 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14552)

			gen14554 := Call(__e, __e.Global(symcons), gen14551, gen14553)

			gen14555 := Call(__e, __e.Global(symcons), gen14554, Nil)

			gen14556 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14555)

			gen14557 := Call(__e, __e.Global(symcons), A, gen14556)

			__e.TailApply(__e.Global(symunify_b), V4366, gen14557, V4367, V4368)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-trap-error"), gen14558)

		gen14563 := MakeNative(func(__e Evaluator) {
			V4376 := __e.Get(1)
			_ = V4376
			V4377 := __e.Get(2)
			_ = V4377
			V4378 := __e.Get(3)
			_ = V4378
			gen14559 := Call(__e, __e.Global(symshen_4newpv), V4377)

			A := gen14559
			Call(__e, __e.Global(symshen_4incinfs))
			gen14560 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14561 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14560)

			gen14562 := Call(__e, __e.Global(symcons), A, gen14561)

			__e.TailApply(__e.Global(symunify_b), V4376, gen14562, V4377, V4378)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-tuple?"), gen14563)

		gen14567 := MakeNative(func(__e Evaluator) {
			V4386 := __e.Get(1)
			_ = V4386
			V4387 := __e.Get(2)
			_ = V4387
			V4388 := __e.Get(3)
			_ = V4388
			Call(__e, __e.Global(symshen_4incinfs))
			gen14564 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14565 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14564)

			gen14566 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14565)

			__e.TailApply(__e.Global(symunify_b), V4386, gen14566, V4387, V4388)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-undefmacro"), gen14567)

		gen14581 := MakeNative(func(__e Evaluator) {
			V4396 := __e.Get(1)
			_ = V4396
			V4397 := __e.Get(2)
			_ = V4397
			V4398 := __e.Get(3)
			_ = V4398
			gen14568 := Call(__e, __e.Global(symshen_4newpv), V4397)

			A := gen14568
			Call(__e, __e.Global(symshen_4incinfs))
			gen14569 := Call(__e, __e.Global(symcons), A, Nil)

			gen14570 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14569)

			gen14571 := Call(__e, __e.Global(symcons), A, Nil)

			gen14572 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14571)

			gen14573 := Call(__e, __e.Global(symcons), A, Nil)

			gen14574 := Call(__e, __e.Global(symcons), MakeSymbol("list"), gen14573)

			gen14575 := Call(__e, __e.Global(symcons), gen14574, Nil)

			gen14576 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14575)

			gen14577 := Call(__e, __e.Global(symcons), gen14572, gen14576)

			gen14578 := Call(__e, __e.Global(symcons), gen14577, Nil)

			gen14579 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14578)

			gen14580 := Call(__e, __e.Global(symcons), gen14570, gen14579)

			__e.TailApply(__e.Global(symunify_b), V4396, gen14580, V4397, V4398)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-union"), gen14581)

		gen14593 := MakeNative(func(__e Evaluator) {
			V4406 := __e.Get(1)
			_ = V4406
			V4407 := __e.Get(2)
			_ = V4407
			V4408 := __e.Get(3)
			_ = V4408
			gen14582 := Call(__e, __e.Global(symshen_4newpv), V4407)

			A := gen14582
			gen14583 := Call(__e, __e.Global(symshen_4newpv), V4407)

			B := gen14583
			Call(__e, __e.Global(symshen_4incinfs))
			gen14584 := Call(__e, __e.Global(symcons), B, Nil)

			gen14585 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14584)

			gen14586 := Call(__e, __e.Global(symcons), A, gen14585)

			gen14587 := Call(__e, __e.Global(symcons), B, Nil)

			gen14588 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14587)

			gen14589 := Call(__e, __e.Global(symcons), A, gen14588)

			gen14590 := Call(__e, __e.Global(symcons), gen14589, Nil)

			gen14591 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14590)

			gen14592 := Call(__e, __e.Global(symcons), gen14586, gen14591)

			__e.TailApply(__e.Global(symunify_b), V4406, gen14592, V4407, V4408)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-unprofile"), gen14593)

		gen14597 := MakeNative(func(__e Evaluator) {
			V4416 := __e.Get(1)
			_ = V4416
			V4417 := __e.Get(2)
			_ = V4417
			V4418 := __e.Get(3)
			_ = V4418
			Call(__e, __e.Global(symshen_4incinfs))
			gen14594 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14595 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14594)

			gen14596 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14595)

			__e.TailApply(__e.Global(symunify_b), V4416, gen14596, V4417, V4418)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-untrack"), gen14597)

		gen14601 := MakeNative(func(__e Evaluator) {
			V4426 := __e.Get(1)
			_ = V4426
			V4427 := __e.Get(2)
			_ = V4427
			V4428 := __e.Get(3)
			_ = V4428
			Call(__e, __e.Global(symshen_4incinfs))
			gen14598 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), Nil)

			gen14599 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14598)

			gen14600 := Call(__e, __e.Global(symcons), MakeSymbol("symbol"), gen14599)

			__e.TailApply(__e.Global(symunify_b), V4426, gen14600, V4427, V4428)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-unspecialise"), gen14601)

		gen14606 := MakeNative(func(__e Evaluator) {
			V4436 := __e.Get(1)
			_ = V4436
			V4437 := __e.Get(2)
			_ = V4437
			V4438 := __e.Get(3)
			_ = V4438
			gen14602 := Call(__e, __e.Global(symshen_4newpv), V4437)

			A := gen14602
			Call(__e, __e.Global(symshen_4incinfs))
			gen14603 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14604 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14603)

			gen14605 := Call(__e, __e.Global(symcons), A, gen14604)

			__e.TailApply(__e.Global(symunify_b), V4436, gen14605, V4437, V4438)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-variable?"), gen14606)

		gen14611 := MakeNative(func(__e Evaluator) {
			V4446 := __e.Get(1)
			_ = V4446
			V4447 := __e.Get(2)
			_ = V4447
			V4448 := __e.Get(3)
			_ = V4448
			gen14607 := Call(__e, __e.Global(symshen_4newpv), V4447)

			A := gen14607
			Call(__e, __e.Global(symshen_4incinfs))
			gen14608 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14609 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14608)

			gen14610 := Call(__e, __e.Global(symcons), A, gen14609)

			__e.TailApply(__e.Global(symunify_b), V4446, gen14610, V4447, V4448)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-vector?"), gen14611)

		gen14614 := MakeNative(func(__e Evaluator) {
			V4456 := __e.Get(1)
			_ = V4456
			V4457 := __e.Get(2)
			_ = V4457
			V4458 := __e.Get(3)
			_ = V4458
			Call(__e, __e.Global(symshen_4incinfs))
			gen14612 := Call(__e, __e.Global(symcons), MakeSymbol("string"), Nil)

			gen14613 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14612)

			__e.TailApply(__e.Global(symunify_b), V4456, gen14613, V4457, V4458)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-version"), gen14614)

		gen14622 := MakeNative(func(__e Evaluator) {
			V4466 := __e.Get(1)
			_ = V4466
			V4467 := __e.Get(2)
			_ = V4467
			V4468 := __e.Get(3)
			_ = V4468
			gen14615 := Call(__e, __e.Global(symshen_4newpv), V4467)

			A := gen14615
			Call(__e, __e.Global(symshen_4incinfs))
			gen14616 := Call(__e, __e.Global(symcons), A, Nil)

			gen14617 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14616)

			gen14618 := Call(__e, __e.Global(symcons), A, gen14617)

			gen14619 := Call(__e, __e.Global(symcons), gen14618, Nil)

			gen14620 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14619)

			gen14621 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14620)

			__e.TailApply(__e.Global(symunify_b), V4466, gen14621, V4467, V4468)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-write-to-file"), gen14622)

		gen14631 := MakeNative(func(__e Evaluator) {
			V4476 := __e.Get(1)
			_ = V4476
			V4477 := __e.Get(2)
			_ = V4477
			V4478 := __e.Get(3)
			_ = V4478
			Call(__e, __e.Global(symshen_4incinfs))
			gen14623 := Call(__e, __e.Global(symcons), MakeSymbol("out"), Nil)

			gen14624 := Call(__e, __e.Global(symcons), MakeSymbol("stream"), gen14623)

			gen14625 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14626 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14625)

			gen14627 := Call(__e, __e.Global(symcons), gen14624, gen14626)

			gen14628 := Call(__e, __e.Global(symcons), gen14627, Nil)

			gen14629 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14628)

			gen14630 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14629)

			__e.TailApply(__e.Global(symunify_b), V4476, gen14630, V4477, V4478)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-write-byte"), gen14631)

		gen14635 := MakeNative(func(__e Evaluator) {
			V4486 := __e.Get(1)
			_ = V4486
			V4487 := __e.Get(2)
			_ = V4487
			V4488 := __e.Get(3)
			_ = V4488
			Call(__e, __e.Global(symshen_4incinfs))
			gen14632 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14633 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14632)

			gen14634 := Call(__e, __e.Global(symcons), MakeSymbol("string"), gen14633)

			__e.TailApply(__e.Global(symunify_b), V4486, gen14634, V4487, V4488)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-y-or-n?"), gen14635)

		gen14642 := MakeNative(func(__e Evaluator) {
			V4496 := __e.Get(1)
			_ = V4496
			V4497 := __e.Get(2)
			_ = V4497
			V4498 := __e.Get(3)
			_ = V4498
			Call(__e, __e.Global(symshen_4incinfs))
			gen14636 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14637 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14636)

			gen14638 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14637)

			gen14639 := Call(__e, __e.Global(symcons), gen14638, Nil)

			gen14640 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14639)

			gen14641 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14640)

			__e.TailApply(__e.Global(symunify_b), V4496, gen14641, V4497, V4498)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of->"), gen14642)

		gen14649 := MakeNative(func(__e Evaluator) {
			V4506 := __e.Get(1)
			_ = V4506
			V4507 := __e.Get(2)
			_ = V4507
			V4508 := __e.Get(3)
			_ = V4508
			Call(__e, __e.Global(symshen_4incinfs))
			gen14643 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14644 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14643)

			gen14645 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14644)

			gen14646 := Call(__e, __e.Global(symcons), gen14645, Nil)

			gen14647 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14646)

			gen14648 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14647)

			__e.TailApply(__e.Global(symunify_b), V4506, gen14648, V4507, V4508)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-<"), gen14649)

		gen14656 := MakeNative(func(__e Evaluator) {
			V4516 := __e.Get(1)
			_ = V4516
			V4517 := __e.Get(2)
			_ = V4517
			V4518 := __e.Get(3)
			_ = V4518
			Call(__e, __e.Global(symshen_4incinfs))
			gen14650 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14651 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14650)

			gen14652 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14651)

			gen14653 := Call(__e, __e.Global(symcons), gen14652, Nil)

			gen14654 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14653)

			gen14655 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14654)

			__e.TailApply(__e.Global(symunify_b), V4516, gen14655, V4517, V4518)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of->="), gen14656)

		gen14663 := MakeNative(func(__e Evaluator) {
			V4526 := __e.Get(1)
			_ = V4526
			V4527 := __e.Get(2)
			_ = V4527
			V4528 := __e.Get(3)
			_ = V4528
			Call(__e, __e.Global(symshen_4incinfs))
			gen14657 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14658 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14657)

			gen14659 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14658)

			gen14660 := Call(__e, __e.Global(symcons), gen14659, Nil)

			gen14661 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14660)

			gen14662 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14661)

			__e.TailApply(__e.Global(symunify_b), V4526, gen14662, V4527, V4528)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-<="), gen14663)

		gen14671 := MakeNative(func(__e Evaluator) {
			V4536 := __e.Get(1)
			_ = V4536
			V4537 := __e.Get(2)
			_ = V4537
			V4538 := __e.Get(3)
			_ = V4538
			gen14664 := Call(__e, __e.Global(symshen_4newpv), V4537)

			A := gen14664
			Call(__e, __e.Global(symshen_4incinfs))
			gen14665 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14666 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14665)

			gen14667 := Call(__e, __e.Global(symcons), A, gen14666)

			gen14668 := Call(__e, __e.Global(symcons), gen14667, Nil)

			gen14669 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14668)

			gen14670 := Call(__e, __e.Global(symcons), A, gen14669)

			__e.TailApply(__e.Global(symunify_b), V4536, gen14670, V4537, V4538)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-="), gen14671)

		gen14678 := MakeNative(func(__e Evaluator) {
			V4546 := __e.Get(1)
			_ = V4546
			V4547 := __e.Get(2)
			_ = V4547
			V4548 := __e.Get(3)
			_ = V4548
			Call(__e, __e.Global(symshen_4incinfs))
			gen14672 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14673 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14672)

			gen14674 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14673)

			gen14675 := Call(__e, __e.Global(symcons), gen14674, Nil)

			gen14676 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14675)

			gen14677 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14676)

			__e.TailApply(__e.Global(symunify_b), V4546, gen14677, V4547, V4548)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-+"), gen14678)

		gen14685 := MakeNative(func(__e Evaluator) {
			V4556 := __e.Get(1)
			_ = V4556
			V4557 := __e.Get(2)
			_ = V4557
			V4558 := __e.Get(3)
			_ = V4558
			Call(__e, __e.Global(symshen_4incinfs))
			gen14679 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14680 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14679)

			gen14681 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14680)

			gen14682 := Call(__e, __e.Global(symcons), gen14681, Nil)

			gen14683 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14682)

			gen14684 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14683)

			__e.TailApply(__e.Global(symunify_b), V4556, gen14684, V4557, V4558)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-/"), gen14685)

		gen14692 := MakeNative(func(__e Evaluator) {
			V4566 := __e.Get(1)
			_ = V4566
			V4567 := __e.Get(2)
			_ = V4567
			V4568 := __e.Get(3)
			_ = V4568
			Call(__e, __e.Global(symshen_4incinfs))
			gen14686 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14687 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14686)

			gen14688 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14687)

			gen14689 := Call(__e, __e.Global(symcons), gen14688, Nil)

			gen14690 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14689)

			gen14691 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14690)

			__e.TailApply(__e.Global(symunify_b), V4566, gen14691, V4567, V4568)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of--"), gen14692)

		gen14699 := MakeNative(func(__e Evaluator) {
			V4576 := __e.Get(1)
			_ = V4576
			V4577 := __e.Get(2)
			_ = V4577
			V4578 := __e.Get(3)
			_ = V4578
			Call(__e, __e.Global(symshen_4incinfs))
			gen14693 := Call(__e, __e.Global(symcons), MakeSymbol("number"), Nil)

			gen14694 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14693)

			gen14695 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14694)

			gen14696 := Call(__e, __e.Global(symcons), gen14695, Nil)

			gen14697 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14696)

			gen14698 := Call(__e, __e.Global(symcons), MakeSymbol("number"), gen14697)

			__e.TailApply(__e.Global(symunify_b), V4576, gen14698, V4577, V4578)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.type-signature-of-*"), gen14699)

		gen14708 := MakeNative(func(__e Evaluator) {
			V4586 := __e.Get(1)
			_ = V4586
			V4587 := __e.Get(2)
			_ = V4587
			V4588 := __e.Get(3)
			_ = V4588
			gen14700 := Call(__e, __e.Global(symshen_4newpv), V4587)

			A := gen14700
			gen14701 := Call(__e, __e.Global(symshen_4newpv), V4587)

			B := gen14701
			Call(__e, __e.Global(symshen_4incinfs))
			gen14702 := Call(__e, __e.Global(symcons), MakeSymbol("boolean"), Nil)

			gen14703 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14702)

			gen14704 := Call(__e, __e.Global(symcons), B, gen14703)

			gen14705 := Call(__e, __e.Global(symcons), gen14704, Nil)

			gen14706 := Call(__e, __e.Global(symcons), MakeSymbol("-->"), gen14705)

			gen14707 := Call(__e, __e.Global(symcons), A, gen14706)

			__e.TailApply(__e.Global(symunify_b), V4586, gen14707, V4587, V4588)

			return

		}, 3)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.type-signature-of-=="), gen14708)

		return

	}, 0))
}
