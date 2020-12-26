package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen11834 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V701 := __args[0]
			_ = V701
			gen11817 := Call(__e, ShenFunc(symget_1time), MakeSymbol("run"))

			Start := gen11817
			gen11818 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

			gen11819 := Call(__e, ShenFunc(symread_1file), V701)

			gen11820 := Call(__e, ShenFunc(symshen_4load_1help), gen11818, gen11819)

			Result := gen11820
			gen11821 := Call(__e, ShenFunc(symget_1time), MakeSymbol("run"))

			Finish := gen11821
			gen11822 := Call(__e, ShenFunc(sym_1), Finish, Start)

			Time := gen11822
			gen11823 := Call(__e, ShenFunc(symstr), Time)

			gen11824 := Call(__e, ShenFunc(symcn), gen11823, MakeString(" secs\n"))

			gen11825 := Call(__e, ShenFunc(symcn), MakeString("\nrun time: "), gen11824)

			gen11826 := Call(__e, ShenFunc(symstoutput))

			gen11827 := Call(__e, ShenFunc(symshen_4prhush), gen11825, gen11826)

			Message := gen11827
			_ = Message
			Load := Result
			_ = Load
			gen11832 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

			var gen11833 Obj
			if True == gen11832 {
				gen11828 := Call(__e, ShenFunc(syminferences))

				gen11829 := Call(__e, ShenFunc(symshen_4app), gen11828, MakeString(" inferences\n"), MakeSymbol("shen.a"))

				gen11830 := Call(__e, ShenFunc(symcn), MakeString("\ntypechecked in "), gen11829)

				gen11831 := Call(__e, ShenFunc(symstoutput))

				gen11833 = Call(__e, ShenFunc(symshen_4prhush), gen11830, gen11831)

			} else {
				gen11833 = MakeSymbol("shen.skip")
			}
			Infs := gen11833
			_ = Infs
			__e.Return(MakeSymbol("loaded"))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("load"), gen11834)

		gen11849 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V708 := __args[0]
			_ = V708
			V709 := __args[1]
			_ = V709
			gen11848 := Call(__e, ShenFunc(sym_a), False, V708)

			if True == gen11848 {
				gen11847 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					gen11844 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), X)

					gen11845 := Call(__e, ShenFunc(symshen_4app), gen11844, MakeString("\n"), MakeSymbol("shen.s"))

					gen11846 := Call(__e, ShenFunc(symstoutput))

					__e.TailApply(ShenFunc(symshen_4prhush), gen11845, gen11846)

					return

				}, 1)
				__e.TailApply(ShenFunc(symshen_4for_1each), gen11847, V709)

				return

			} else {
				gen11835 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4remove_1synonyms), X)

					return
				}, 1)
				gen11836 := Call(__e, ShenFunc(symmapcan), gen11835, V709)

				RemoveSynonyms := gen11836
				gen11837 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4typetable), X)

					return
				}, 1)
				gen11838 := Call(__e, ShenFunc(symmapcan), gen11837, RemoveSynonyms)

				Table := gen11838
				gen11839 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4assumetype), X)

					return
				}, 1)
				gen11840 := Call(__e, ShenFunc(symshen_4for_1each), gen11839, Table)

				Assume := gen11840
				_ = Assume
				gen11841 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.TailApply(ShenFunc(symshen_4unwind_1types), E, Table)

					return
				}, 1)
				gen11843 := MakeNative(func(__e Evaluator, __args ...Obj) {
					gen11842 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						__e.TailApply(ShenFunc(symshen_4typecheck_1and_1load), X)

						return
					}, 1)
					__e.TailApply(ShenFunc(symshen_4for_1each), gen11842, RemoveSynonyms)

					return

				}, 0)
				__e.Return(Try(__e, gen11843).Catch(gen11841))
				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.load-help"), gen11849)

		gen11854 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V711 := __args[0]
			_ = V711
			gen11852 := Call(__e, ShenFunc(symcons_2), V711)

			var gen11853 Obj
			if True == gen11852 {
				gen11850 := Call(__e, ShenFunc(symhd), V711)

				gen11851 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.synonyms-help"), gen11850)

				if True == gen11851 {
					gen11853 = True
				} else {
					gen11853 = False
				}

			} else {
				gen11853 = False
			}
			if True == gen11853 {
				Call(__e, ShenFunc(symeval), V711)
				__e.Return(Nil)
				return

			} else {
				__e.TailApply(ShenFunc(symcons), V711, Nil)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remove-synonyms"), gen11854)

		gen11856 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V713 := __args[0]
			_ = V713
			Call(__e, ShenFunc(symnl), MakeNumber(1))
			gen11855 := Call(__e, ShenFunc(symgensym), MakeSymbol("A"))

			__e.TailApply(ShenFunc(symshen_4typecheck_1and_1evaluate), V713, gen11855)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typecheck-and-load"), gen11856)

		gen11875 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V719 := __args[0]
			_ = V719
			gen11873 := Call(__e, ShenFunc(symcons_2), V719)

			var gen11874 Obj
			if True == gen11873 {
				gen11870 := Call(__e, ShenFunc(symhd), V719)

				gen11871 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), gen11870)

				var gen11872 Obj
				if True == gen11871 {
					gen11868 := Call(__e, ShenFunc(symtl), V719)

					gen11869 := Call(__e, ShenFunc(symcons_2), gen11868)

					if True == gen11869 {
						gen11872 = True
					} else {
						gen11872 = False
					}

				} else {
					gen11872 = False
				}
				if True == gen11872 {
					gen11874 = True
				} else {
					gen11874 = False
				}

			} else {
				gen11874 = False
			}
			if True == gen11874 {
				gen11857 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Y := __args[0]
					_ = Y
					__e.TailApply(ShenFunc(symshen_4_5sig_7rest_6), Y)

					return
				}, 1)
				gen11858 := Call(__e, ShenFunc(symtl), V719)

				gen11859 := Call(__e, ShenFunc(symtl), gen11858)

				gen11863 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					gen11860 := Call(__e, ShenFunc(symtl), V719)

					gen11861 := Call(__e, ShenFunc(symhd), gen11860)

					gen11862 := Call(__e, ShenFunc(symshen_4app), gen11861, MakeString(" lacks a proper signature.\n"), MakeSymbol("shen.a"))

					__e.TailApply(ShenFunc(symsimple_1error), gen11862)

					return

				}, 1)
				gen11864 := Call(__e, ShenFunc(symcompile), gen11857, gen11859, gen11863)

				Sig := gen11864
				gen11865 := Call(__e, ShenFunc(symtl), V719)

				gen11866 := Call(__e, ShenFunc(symhd), gen11865)

				gen11867 := Call(__e, ShenFunc(symcons), gen11866, Sig)

				__e.TailApply(ShenFunc(symcons), gen11867, Nil)

				return

			} else {
				__e.Return(Nil)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typetable"), gen11875)

		gen11879 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V721 := __args[0]
			_ = V721
			gen11878 := Call(__e, ShenFunc(symcons_2), V721)

			if True == gen11878 {
				gen11876 := Call(__e, ShenFunc(symhd), V721)

				gen11877 := Call(__e, ShenFunc(symtl), V721)

				__e.TailApply(ShenFunc(symdeclare), gen11876, gen11877)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.assumetype"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assumetype"), gen11879)

		gen11889 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V728 := __args[0]
			_ = V728
			V729 := __args[1]
			_ = V729
			gen11888 := Call(__e, ShenFunc(sym_a), Nil, V729)

			if True == gen11888 {
				gen11887 := Call(__e, ShenFunc(symerror_1to_1string), V728)

				__e.TailApply(ShenFunc(symsimple_1error), gen11887)

				return

			} else {
				gen11885 := Call(__e, ShenFunc(symcons_2), V729)

				var gen11886 Obj
				if True == gen11885 {
					gen11883 := Call(__e, ShenFunc(symhd), V729)

					gen11884 := Call(__e, ShenFunc(symcons_2), gen11883)

					if True == gen11884 {
						gen11886 = True
					} else {
						gen11886 = False
					}

				} else {
					gen11886 = False
				}
				if True == gen11886 {
					gen11880 := Call(__e, ShenFunc(symhd), V729)

					gen11881 := Call(__e, ShenFunc(symhd), gen11880)

					Call(__e, ShenFunc(symshen_4remtype), gen11881)

					gen11882 := Call(__e, ShenFunc(symtl), V729)

					__e.TailApply(ShenFunc(symshen_4unwind_1types), V728, gen11882)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.unwind-types"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.unwind-types"), gen11889)

		gen11892 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V731 := __args[0]
			_ = V731
			gen11890 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen11891 := Call(__e, ShenFunc(symshen_4removetype), V731, gen11890)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen11891)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remtype"), gen11892)

		gen11909 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V739 := __args[0]
			_ = V739
			V740 := __args[1]
			_ = V740
			gen11908 := Call(__e, ShenFunc(sym_a), Nil, V740)

			if True == gen11908 {
				__e.Return(Nil)
				return
			} else {
				gen11906 := Call(__e, ShenFunc(symcons_2), V740)

				var gen11907 Obj
				if True == gen11906 {
					gen11903 := Call(__e, ShenFunc(symhd), V740)

					gen11904 := Call(__e, ShenFunc(symcons_2), gen11903)

					var gen11905 Obj
					if True == gen11904 {
						gen11900 := Call(__e, ShenFunc(symhd), V740)

						gen11901 := Call(__e, ShenFunc(symhd), gen11900)

						gen11902 := Call(__e, ShenFunc(sym_a), gen11901, V739)

						if True == gen11902 {
							gen11905 = True
						} else {
							gen11905 = False
						}

					} else {
						gen11905 = False
					}
					if True == gen11905 {
						gen11907 = True
					} else {
						gen11907 = False
					}

				} else {
					gen11907 = False
				}
				if True == gen11907 {
					gen11897 := Call(__e, ShenFunc(symhd), V740)

					gen11898 := Call(__e, ShenFunc(symhd), gen11897)

					gen11899 := Call(__e, ShenFunc(symtl), V740)

					__e.TailApply(ShenFunc(symshen_4removetype), gen11898, gen11899)

					return

				} else {
					gen11896 := Call(__e, ShenFunc(symcons_2), V740)

					if True == gen11896 {
						gen11893 := Call(__e, ShenFunc(symhd), V740)

						gen11894 := Call(__e, ShenFunc(symtl), V740)

						gen11895 := Call(__e, ShenFunc(symshen_4removetype), V739, gen11894)

						__e.TailApply(ShenFunc(symcons), gen11893, gen11895)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.removetype"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.removetype"), gen11909)

		gen11920 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V742 := __args[0]
			_ = V742
			gen11910 := Call(__e, ShenFunc(symshen_4_5signature_6), V742)

			Parse__shen_4_5signature_6 := gen11910
			gen11917 := Call(__e, ShenFunc(symfail))

			gen11918 := Call(__e, ShenFunc(sym_a), gen11917, Parse__shen_4_5signature_6)

			gen11919 := Call(__e, ShenFunc(symnot), gen11918)

			if True == gen11919 {
				gen11911 := Call(__e, ShenFunc(sym_5_b_6), Parse__shen_4_5signature_6)

				Parse___5_b_6 := gen11911
				gen11914 := Call(__e, ShenFunc(symfail))

				gen11915 := Call(__e, ShenFunc(sym_a), gen11914, Parse___5_b_6)

				gen11916 := Call(__e, ShenFunc(symnot), gen11915)

				if True == gen11916 {
					gen11912 := Call(__e, ShenFunc(symhd), Parse___5_b_6)

					gen11913 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen11912, gen11913)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<sig+rest>"), gen11920)

		gen11926 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V745 := __args[0]
			_ = V745
			V746 := __args[1]
			_ = V746
			gen11921 := Call(__e, ShenFunc(symopen), V745, MakeSymbol("out"))

			Stream := gen11921
			gen11922 := Call(__e, ShenFunc(symstring_2), V746)

			var gen11923 Obj
			if True == gen11922 {
				gen11923 = Call(__e, ShenFunc(symshen_4app), V746, MakeString("\n\n"), MakeSymbol("shen.a"))

			} else {
				gen11923 = Call(__e, ShenFunc(symshen_4app), V746, MakeString("\n\n"), MakeSymbol("shen.s"))

			}
			String := gen11923
			gen11924 := Call(__e, ShenFunc(sympr), String, Stream)

			Write := gen11924
			_ = Write
			gen11925 := Call(__e, ShenFunc(symclose), Stream)

			Close := gen11925
			_ = Close
			__e.Return(V746)
			return

		}, 2)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("write-to-file"), gen11926)

		return

	}, 0))
}
