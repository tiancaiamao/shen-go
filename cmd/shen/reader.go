package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen6817 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1430 := __args[0]
			_ = V1430
			__e.TailApply(ShenFunc(symread_1byte), V1430)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-char-code"), gen6817)

		gen6819 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1432 := __args[0]
			_ = V1432
			gen6818 := MakeNative(func(__e Evaluator, __args ...Obj) {
				S := __args[0]
				_ = S
				__e.TailApply(ShenFunc(symread_1byte), S)

				return
			}, 1)
			__e.TailApply(ShenFunc(symshen_4read_1file_1as_1Xlist), V1432, gen6818)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-file-as-bytelist"), gen6819)

		gen6821 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1434 := __args[0]
			_ = V1434
			gen6820 := MakeNative(func(__e Evaluator, __args ...Obj) {
				S := __args[0]
				_ = S
				__e.TailApply(ShenFunc(symshen_4read_1char_1code), S)

				return
			}, 1)
			__e.TailApply(ShenFunc(symshen_4read_1file_1as_1Xlist), V1434, gen6820)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-file-as-charlist"), gen6821)

		gen6826 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1437 := __args[0]
			_ = V1437
			V1438 := __args[1]
			_ = V1438
			gen6822 := Call(__e, ShenFunc(symopen), V1437, MakeSymbol("in"))

			Stream := gen6822
			gen6823 := Call(__e, V1438, Stream)

			X := gen6823
			gen6824 := Call(__e, ShenFunc(symshen_4read_1file_1as_1Xlist_1help), Stream, V1438, X, Nil)

			Xs := gen6824
			gen6825 := Call(__e, ShenFunc(symclose), Stream)

			Close := gen6825
			_ = Close
			__e.TailApply(ShenFunc(symreverse), Xs)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-file-as-Xlist"), gen6826)

		gen6830 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1443 := __args[0]
			_ = V1443
			V1444 := __args[1]
			_ = V1444
			V1445 := __args[2]
			_ = V1445
			V1446 := __args[3]
			_ = V1446
			gen6829 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1445)

			if True == gen6829 {
				__e.Return(V1446)
				return
			} else {
				gen6827 := Call(__e, V1444, V1443)

				gen6828 := Call(__e, ShenFunc(symcons), V1445, V1446)

				__e.TailApply(ShenFunc(symshen_4read_1file_1as_1Xlist_1help), V1443, V1444, gen6827, gen6828)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-file-as-Xlist-help"), gen6830)

		gen6833 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1448 := __args[0]
			_ = V1448
			gen6831 := Call(__e, ShenFunc(symopen), V1448, MakeSymbol("in"))

			Stream := gen6831
			gen6832 := Call(__e, ShenFunc(symshen_4read_1char_1code), Stream)

			__e.TailApply(ShenFunc(symshen_4rfas_1h), Stream, gen6832, MakeString(""))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-file-as-string"), gen6833)

		gen6838 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1452 := __args[0]
			_ = V1452
			V1453 := __args[1]
			_ = V1453
			V1454 := __args[2]
			_ = V1454
			gen6837 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1453)

			if True == gen6837 {
				Call(__e, ShenFunc(symclose), V1452)
				__e.Return(V1454)
				return

			} else {
				gen6834 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1452)

				gen6835 := Call(__e, ShenFunc(symn_1_6string), V1453)

				gen6836 := Call(__e, ShenFunc(symcn), V1454, gen6835)

				__e.TailApply(ShenFunc(symshen_4rfas_1h), V1452, gen6834, gen6836)

				return

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rfas-h"), gen6838)

		gen6840 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1456 := __args[0]
			_ = V1456
			gen6839 := Call(__e, ShenFunc(symread), V1456)

			__e.TailApply(ShenFunc(symeval_1kl), gen6839)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("input"), gen6840)

		gen6850 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1459 := __args[0]
			_ = V1459
			V1460 := __args[1]
			_ = V1460
			gen6841 := Call(__e, ShenFunc(symshen_4monotype), V1459)

			Mono_2 := gen6841
			_ = Mono_2
			gen6842 := Call(__e, ShenFunc(symread), V1460)

			Input := gen6842
			gen6847 := Call(__e, ShenFunc(symshen_4demodulate), V1459)

			gen6848 := Call(__e, ShenFunc(symshen_4typecheck), Input, gen6847)

			gen6849 := Call(__e, ShenFunc(sym_a), False, gen6848)

			if True == gen6849 {
				gen6843 := Call(__e, ShenFunc(symshen_4app), V1459, MakeString("\n"), MakeSymbol("shen.r"))

				gen6844 := Call(__e, ShenFunc(symcn), MakeString(" is not of type "), gen6843)

				gen6845 := Call(__e, ShenFunc(symshen_4app), Input, gen6844, MakeSymbol("shen.r"))

				gen6846 := Call(__e, ShenFunc(symcn), MakeString("type error: "), gen6845)

				__e.TailApply(ShenFunc(symsimple_1error), gen6846)

				return

			} else {
				__e.TailApply(ShenFunc(symeval_1kl), Input)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("input+"), gen6850)

		gen6856 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1462 := __args[0]
			_ = V1462
			gen6855 := Call(__e, ShenFunc(symcons_2), V1462)

			if True == gen6855 {
				gen6854 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Z := __args[0]
					_ = Z
					__e.TailApply(ShenFunc(symshen_4monotype), Z)

					return
				}, 1)
				__e.TailApply(ShenFunc(symmap), gen6854, V1462)

				return

			} else {
				gen6853 := Call(__e, ShenFunc(symvariable_2), V1462)

				if True == gen6853 {
					gen6851 := Call(__e, ShenFunc(symshen_4app), V1462, MakeString("\n"), MakeSymbol("shen.a"))

					gen6852 := Call(__e, ShenFunc(symcn), MakeString("input+ expects a monotype: not "), gen6851)

					__e.TailApply(ShenFunc(symsimple_1error), gen6852)

					return

				} else {
					__e.Return(V1462)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.monotype"), gen6856)

		gen6859 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1464 := __args[0]
			_ = V1464
			gen6857 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1464)

			gen6858 := Call(__e, ShenFunc(symshen_4read_1loop), V1464, gen6857, Nil)

			__e.TailApply(ShenFunc(symhd), gen6858)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read"), gen6859)

		gen6860 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*it*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("it"), gen6860)

		gen6880 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1472 := __args[0]
			_ = V1472
			V1473 := __args[1]
			_ = V1473
			V1474 := __args[2]
			_ = V1474
			gen6879 := Call(__e, ShenFunc(sym_a), MakeNumber(94), V1473)

			if True == gen6879 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("read aborted"))

				return
			} else {
				gen6878 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1473)

				if True == gen6878 {
					gen6877 := Call(__e, ShenFunc(symempty_2), V1474)

					if True == gen6877 {
						__e.TailApply(ShenFunc(symsimple_1error), MakeString("error: empty stream"))

						return
					} else {
						gen6875 := MakeNative(func(__e Evaluator, __args ...Obj) {
							X := __args[0]
							_ = X
							__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

							return
						}, 1)
						gen6876 := MakeNative(func(__e Evaluator, __args ...Obj) {
							E := __args[0]
							_ = E
							__e.Return(E)
							return
						}, 1)
						__e.TailApply(ShenFunc(symcompile), gen6875, V1474, gen6876)

						return

					}

				} else {
					gen6874 := Call(__e, ShenFunc(symshen_4terminator_2), V1473)

					if True == gen6874 {
						gen6864 := Call(__e, ShenFunc(symcons), V1473, Nil)

						gen6865 := Call(__e, ShenFunc(symappend), V1474, gen6864)

						AllChars := gen6865
						gen6866 := Call(__e, ShenFunc(symshen_4record_1it), AllChars)

						It := gen6866
						_ = It
						gen6867 := MakeNative(func(__e Evaluator, __args ...Obj) {
							X := __args[0]
							_ = X
							__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

							return
						}, 1)
						gen6868 := MakeNative(func(__e Evaluator, __args ...Obj) {
							E := __args[0]
							_ = E
							__e.Return(MakeSymbol("shen.nextbyte"))
							return
						}, 1)
						gen6869 := Call(__e, ShenFunc(symcompile), gen6867, AllChars, gen6868)

						Read := gen6869
						gen6872 := Call(__e, ShenFunc(sym_a), Read, MakeSymbol("shen.nextbyte"))

						var gen6873 Obj
						if True == gen6872 {
							gen6873 = True
						} else {
							gen6871 := Call(__e, ShenFunc(symempty_2), Read)

							if True == gen6871 {
								gen6873 = True
							} else {
								gen6873 = False
							}

						}
						if True == gen6873 {
							gen6870 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1472)

							__e.TailApply(ShenFunc(symshen_4read_1loop), V1472, gen6870, AllChars)

							return

						} else {
							__e.Return(Read)
							return
						}

					} else {
						gen6861 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1472)

						gen6862 := Call(__e, ShenFunc(symcons), V1473, Nil)

						gen6863 := Call(__e, ShenFunc(symappend), V1474, gen6862)

						__e.TailApply(ShenFunc(symshen_4read_1loop), V1472, gen6861, gen6863)

						return

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-loop"), gen6880)

		gen6888 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1476 := __args[0]
			_ = V1476
			gen6881 := Call(__e, ShenFunc(symcons), MakeNumber(93), Nil)

			gen6882 := Call(__e, ShenFunc(symcons), MakeNumber(41), gen6881)

			gen6883 := Call(__e, ShenFunc(symcons), MakeNumber(34), gen6882)

			gen6884 := Call(__e, ShenFunc(symcons), MakeNumber(32), gen6883)

			gen6885 := Call(__e, ShenFunc(symcons), MakeNumber(13), gen6884)

			gen6886 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen6885)

			gen6887 := Call(__e, ShenFunc(symcons), MakeNumber(9), gen6886)

			__e.TailApply(ShenFunc(symelement_2), V1476, gen6887)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.terminator?"), gen6888)

		gen6890 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1478 := __args[0]
			_ = V1478
			gen6889 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1478)

			__e.TailApply(ShenFunc(symshen_4lineread_1loop), gen6889, Nil, V1478)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("lineread"), gen6890)

		gen6915 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1483 := __args[0]
			_ = V1483
			V1484 := __args[1]
			_ = V1484
			V1485 := __args[2]
			_ = V1485
			gen6914 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1483)

			if True == gen6914 {
				gen6913 := Call(__e, ShenFunc(symempty_2), V1484)

				if True == gen6913 {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("empty stream"))

					return
				} else {
					gen6911 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

						return
					}, 1)
					gen6912 := MakeNative(func(__e Evaluator, __args ...Obj) {
						E := __args[0]
						_ = E
						__e.Return(E)
						return
					}, 1)
					__e.TailApply(ShenFunc(symcompile), gen6911, V1484, gen6912)

					return

				}

			} else {
				gen6909 := Call(__e, ShenFunc(symshen_4hat))

				gen6910 := Call(__e, ShenFunc(sym_a), V1483, gen6909)

				if True == gen6910 {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("line read aborted"))

					return
				} else {
					gen6904 := Call(__e, ShenFunc(symshen_4newline))

					gen6905 := Call(__e, ShenFunc(symshen_4carriage_1return))

					gen6906 := Call(__e, ShenFunc(symcons), gen6905, Nil)

					gen6907 := Call(__e, ShenFunc(symcons), gen6904, gen6906)

					gen6908 := Call(__e, ShenFunc(symelement_2), V1483, gen6907)

					if True == gen6908 {
						gen6894 := MakeNative(func(__e Evaluator, __args ...Obj) {
							X := __args[0]
							_ = X
							__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

							return
						}, 1)
						gen6895 := MakeNative(func(__e Evaluator, __args ...Obj) {
							E := __args[0]
							_ = E
							__e.Return(MakeSymbol("shen.nextline"))
							return
						}, 1)
						gen6896 := Call(__e, ShenFunc(symcompile), gen6894, V1484, gen6895)

						Line := gen6896
						_ = Line
						gen6897 := Call(__e, ShenFunc(symshen_4record_1it), V1484)

						It := gen6897
						_ = It
						gen6902 := Call(__e, ShenFunc(sym_a), Line, MakeSymbol("shen.nextline"))

						var gen6903 Obj
						if True == gen6902 {
							gen6903 = True
						} else {
							gen6901 := Call(__e, ShenFunc(symempty_2), Line)

							if True == gen6901 {
								gen6903 = True
							} else {
								gen6903 = False
							}

						}
						if True == gen6903 {
							gen6898 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1485)

							gen6899 := Call(__e, ShenFunc(symcons), V1483, Nil)

							gen6900 := Call(__e, ShenFunc(symappend), V1484, gen6899)

							__e.TailApply(ShenFunc(symshen_4lineread_1loop), gen6898, gen6900, V1485)

							return

						} else {
							__e.Return(Line)
							return
						}

					} else {
						gen6891 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1485)

						gen6892 := Call(__e, ShenFunc(symcons), V1483, Nil)

						gen6893 := Call(__e, ShenFunc(symappend), V1484, gen6892)

						__e.TailApply(ShenFunc(symshen_4lineread_1loop), gen6891, gen6893, V1485)

						return

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lineread-loop"), gen6915)

		gen6920 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1487 := __args[0]
			_ = V1487
			gen6916 := Call(__e, ShenFunc(symshen_4trim_1whitespace), V1487)

			TrimLeft := gen6916
			gen6917 := Call(__e, ShenFunc(symreverse), TrimLeft)

			gen6918 := Call(__e, ShenFunc(symshen_4trim_1whitespace), gen6917)

			TrimRight := gen6918
			gen6919 := Call(__e, ShenFunc(symreverse), TrimRight)

			Trimmed := gen6919
			__e.TailApply(ShenFunc(symshen_4record_1it_1h), Trimmed)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-it"), gen6920)

		gen6930 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1489 := __args[0]
			_ = V1489
			gen6928 := Call(__e, ShenFunc(symcons_2), V1489)

			var gen6929 Obj
			if True == gen6928 {
				gen6922 := Call(__e, ShenFunc(symhd), V1489)

				gen6923 := Call(__e, ShenFunc(symcons), MakeNumber(32), Nil)

				gen6924 := Call(__e, ShenFunc(symcons), MakeNumber(13), gen6923)

				gen6925 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen6924)

				gen6926 := Call(__e, ShenFunc(symcons), MakeNumber(9), gen6925)

				gen6927 := Call(__e, ShenFunc(symelement_2), gen6922, gen6926)

				if True == gen6927 {
					gen6929 = True
				} else {
					gen6929 = False
				}

			} else {
				gen6929 = False
			}
			if True == gen6929 {
				gen6921 := Call(__e, ShenFunc(symtl), V1489)

				__e.TailApply(ShenFunc(symshen_4trim_1whitespace), gen6921)

				return

			} else {
				__e.Return(V1489)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.trim-whitespace"), gen6930)

		gen6934 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1491 := __args[0]
			_ = V1491
			gen6931 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symn_1_6string), X)

				return
			}, 1)
			gen6932 := Call(__e, ShenFunc(symmap), gen6931, V1491)

			gen6933 := Call(__e, ShenFunc(symshen_4cn_1all), gen6932)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*it*"), gen6933)

			__e.Return(V1491)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-it-h"), gen6934)

		gen6940 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1493 := __args[0]
			_ = V1493
			gen6939 := Call(__e, ShenFunc(sym_a), Nil, V1493)

			if True == gen6939 {
				__e.Return(MakeString(""))
				return
			} else {
				gen6938 := Call(__e, ShenFunc(symcons_2), V1493)

				if True == gen6938 {
					gen6935 := Call(__e, ShenFunc(symhd), V1493)

					gen6936 := Call(__e, ShenFunc(symtl), V1493)

					gen6937 := Call(__e, ShenFunc(symshen_4cn_1all), gen6936)

					__e.TailApply(ShenFunc(symcn), gen6935, gen6937)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.cn-all"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cn-all"), gen6940)

		gen6944 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1495 := __args[0]
			_ = V1495
			gen6941 := Call(__e, ShenFunc(symshen_4read_1file_1as_1charlist), V1495)

			Charlist := gen6941
			gen6942 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

				return
			}, 1)
			gen6943 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4read_1error), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symcompile), gen6942, Charlist, gen6943)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-file"), gen6944)

		gen6950 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1497 := __args[0]
			_ = V1497
			gen6945 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symstring_1_6n), X)

				return
			}, 1)
			gen6946 := Call(__e, ShenFunc(symexplode), V1497)

			gen6947 := Call(__e, ShenFunc(symmap), gen6945, gen6946)

			Ns := gen6947
			gen6948 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

				return
			}, 1)
			gen6949 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4read_1error), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symcompile), gen6948, Ns, gen6949)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-from-string"), gen6950)

		gen6966 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1505 := __args[0]
			_ = V1505
			gen6964 := Call(__e, ShenFunc(symcons_2), V1505)

			var gen6965 Obj
			if True == gen6964 {
				gen6961 := Call(__e, ShenFunc(symhd), V1505)

				gen6962 := Call(__e, ShenFunc(symcons_2), gen6961)

				var gen6963 Obj
				if True == gen6962 {
					gen6958 := Call(__e, ShenFunc(symtl), V1505)

					gen6959 := Call(__e, ShenFunc(symcons_2), gen6958)

					var gen6960 Obj
					if True == gen6959 {
						gen6955 := Call(__e, ShenFunc(symtl), V1505)

						gen6956 := Call(__e, ShenFunc(symtl), gen6955)

						gen6957 := Call(__e, ShenFunc(sym_a), Nil, gen6956)

						if True == gen6957 {
							gen6960 = True
						} else {
							gen6960 = False
						}

					} else {
						gen6960 = False
					}
					if True == gen6960 {
						gen6963 = True
					} else {
						gen6963 = False
					}

				} else {
					gen6963 = False
				}
				if True == gen6963 {
					gen6965 = True
				} else {
					gen6965 = False
				}

			} else {
				gen6965 = False
			}
			if True == gen6965 {
				gen6951 := Call(__e, ShenFunc(symhd), V1505)

				gen6952 := Call(__e, ShenFunc(symshen_4compress_150), MakeNumber(50), gen6951)

				gen6953 := Call(__e, ShenFunc(symshen_4app), gen6952, MakeString("\n"), MakeSymbol("shen.a"))

				gen6954 := Call(__e, ShenFunc(symcn), MakeString("read error here:\n\n "), gen6953)

				__e.TailApply(ShenFunc(symsimple_1error), gen6954)

				return

			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("read error\n"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-error"), gen6966)

		gen6975 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1512 := __args[0]
			_ = V1512
			V1513 := __args[1]
			_ = V1513
			gen6974 := Call(__e, ShenFunc(sym_a), Nil, V1513)

			if True == gen6974 {
				__e.Return(MakeString(""))
				return
			} else {
				gen6973 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V1512)

				if True == gen6973 {
					__e.Return(MakeString(""))
					return
				} else {
					gen6972 := Call(__e, ShenFunc(symcons_2), V1513)

					if True == gen6972 {
						gen6967 := Call(__e, ShenFunc(symhd), V1513)

						gen6968 := Call(__e, ShenFunc(symn_1_6string), gen6967)

						gen6969 := Call(__e, ShenFunc(sym_1), V1512, MakeNumber(1))

						gen6970 := Call(__e, ShenFunc(symtl), V1513)

						gen6971 := Call(__e, ShenFunc(symshen_4compress_150), gen6969, gen6970)

						__e.TailApply(ShenFunc(symcn), gen6968, gen6971)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.compress-50"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compress-50"), gen6975)

		gen7193 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1515 := __args[0]
			_ = V1515
			gen6976 := Call(__e, ShenFunc(symshen_4_5lsb_6), V1515)

			Parse__shen_4_5lsb_6 := gen6976
			gen6995 := Call(__e, ShenFunc(symfail))

			gen6996 := Call(__e, ShenFunc(sym_a), gen6995, Parse__shen_4_5lsb_6)

			gen6997 := Call(__e, ShenFunc(symnot), gen6996)

			var gen6998 Obj
			if True == gen6997 {
				gen6977 := Call(__e, ShenFunc(symshen_4_5st__input1_6), Parse__shen_4_5lsb_6)

				Parse__shen_4_5st__input1_6 := gen6977
				gen6992 := Call(__e, ShenFunc(symfail))

				gen6993 := Call(__e, ShenFunc(sym_a), gen6992, Parse__shen_4_5st__input1_6)

				gen6994 := Call(__e, ShenFunc(symnot), gen6993)

				if True == gen6994 {
					gen6978 := Call(__e, ShenFunc(symshen_4_5rsb_6), Parse__shen_4_5st__input1_6)

					Parse__shen_4_5rsb_6 := gen6978
					gen6989 := Call(__e, ShenFunc(symfail))

					gen6990 := Call(__e, ShenFunc(sym_a), gen6989, Parse__shen_4_5rsb_6)

					gen6991 := Call(__e, ShenFunc(symnot), gen6990)

					if True == gen6991 {
						gen6979 := Call(__e, ShenFunc(symshen_4_5st__input2_6), Parse__shen_4_5rsb_6)

						Parse__shen_4_5st__input2_6 := gen6979
						gen6986 := Call(__e, ShenFunc(symfail))

						gen6987 := Call(__e, ShenFunc(sym_a), gen6986, Parse__shen_4_5st__input2_6)

						gen6988 := Call(__e, ShenFunc(symnot), gen6987)

						if True == gen6988 {
							gen6980 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input2_6)

							gen6981 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input1_6)

							gen6982 := Call(__e, ShenFunc(symshen_4cons__form), gen6981)

							gen6983 := Call(__e, ShenFunc(symmacroexpand), gen6982)

							gen6984 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input2_6)

							gen6985 := Call(__e, ShenFunc(symcons), gen6983, gen6984)

							gen6998 = Call(__e, ShenFunc(symshen_4pair), gen6980, gen6985)

						} else {
							gen6998 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen6998 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen6998 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6998 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6998
			gen7191 := Call(__e, ShenFunc(symfail))

			gen7192 := Call(__e, ShenFunc(sym_a), YaccParse, gen7191)

			if True == gen7192 {
				gen6999 := Call(__e, ShenFunc(symshen_4_5lrb_6), V1515)

				Parse__shen_4_5lrb_6 := gen6999
				gen7017 := Call(__e, ShenFunc(symfail))

				gen7018 := Call(__e, ShenFunc(sym_a), gen7017, Parse__shen_4_5lrb_6)

				gen7019 := Call(__e, ShenFunc(symnot), gen7018)

				var gen7020 Obj
				if True == gen7019 {
					gen7000 := Call(__e, ShenFunc(symshen_4_5st__input1_6), Parse__shen_4_5lrb_6)

					Parse__shen_4_5st__input1_6 := gen7000
					gen7014 := Call(__e, ShenFunc(symfail))

					gen7015 := Call(__e, ShenFunc(sym_a), gen7014, Parse__shen_4_5st__input1_6)

					gen7016 := Call(__e, ShenFunc(symnot), gen7015)

					if True == gen7016 {
						gen7001 := Call(__e, ShenFunc(symshen_4_5rrb_6), Parse__shen_4_5st__input1_6)

						Parse__shen_4_5rrb_6 := gen7001
						gen7011 := Call(__e, ShenFunc(symfail))

						gen7012 := Call(__e, ShenFunc(sym_a), gen7011, Parse__shen_4_5rrb_6)

						gen7013 := Call(__e, ShenFunc(symnot), gen7012)

						if True == gen7013 {
							gen7002 := Call(__e, ShenFunc(symshen_4_5st__input2_6), Parse__shen_4_5rrb_6)

							Parse__shen_4_5st__input2_6 := gen7002
							gen7008 := Call(__e, ShenFunc(symfail))

							gen7009 := Call(__e, ShenFunc(sym_a), gen7008, Parse__shen_4_5st__input2_6)

							gen7010 := Call(__e, ShenFunc(symnot), gen7009)

							if True == gen7010 {
								gen7003 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input2_6)

								gen7004 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input1_6)

								gen7005 := Call(__e, ShenFunc(symmacroexpand), gen7004)

								gen7006 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input2_6)

								gen7007 := Call(__e, ShenFunc(symshen_4package_1macro), gen7005, gen7006)

								gen7020 = Call(__e, ShenFunc(symshen_4pair), gen7003, gen7007)

							} else {
								gen7020 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen7020 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen7020 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen7020 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen7020
				gen7189 := Call(__e, ShenFunc(symfail))

				gen7190 := Call(__e, ShenFunc(sym_a), YaccParse, gen7189)

				if True == gen7190 {
					gen7021 := Call(__e, ShenFunc(symshen_4_5lcurly_6), V1515)

					Parse__shen_4_5lcurly_6 := gen7021
					gen7029 := Call(__e, ShenFunc(symfail))

					gen7030 := Call(__e, ShenFunc(sym_a), gen7029, Parse__shen_4_5lcurly_6)

					gen7031 := Call(__e, ShenFunc(symnot), gen7030)

					var gen7032 Obj
					if True == gen7031 {
						gen7022 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5lcurly_6)

						Parse__shen_4_5st__input_6 := gen7022
						gen7026 := Call(__e, ShenFunc(symfail))

						gen7027 := Call(__e, ShenFunc(sym_a), gen7026, Parse__shen_4_5st__input_6)

						gen7028 := Call(__e, ShenFunc(symnot), gen7027)

						if True == gen7028 {
							gen7023 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

							gen7024 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

							gen7025 := Call(__e, ShenFunc(symcons), MakeSymbol("{"), gen7024)

							gen7032 = Call(__e, ShenFunc(symshen_4pair), gen7023, gen7025)

						} else {
							gen7032 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen7032 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen7032
					gen7187 := Call(__e, ShenFunc(symfail))

					gen7188 := Call(__e, ShenFunc(sym_a), YaccParse, gen7187)

					if True == gen7188 {
						gen7033 := Call(__e, ShenFunc(symshen_4_5rcurly_6), V1515)

						Parse__shen_4_5rcurly_6 := gen7033
						gen7041 := Call(__e, ShenFunc(symfail))

						gen7042 := Call(__e, ShenFunc(sym_a), gen7041, Parse__shen_4_5rcurly_6)

						gen7043 := Call(__e, ShenFunc(symnot), gen7042)

						var gen7044 Obj
						if True == gen7043 {
							gen7034 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5rcurly_6)

							Parse__shen_4_5st__input_6 := gen7034
							gen7038 := Call(__e, ShenFunc(symfail))

							gen7039 := Call(__e, ShenFunc(sym_a), gen7038, Parse__shen_4_5st__input_6)

							gen7040 := Call(__e, ShenFunc(symnot), gen7039)

							if True == gen7040 {
								gen7035 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

								gen7036 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

								gen7037 := Call(__e, ShenFunc(symcons), MakeSymbol("}"), gen7036)

								gen7044 = Call(__e, ShenFunc(symshen_4pair), gen7035, gen7037)

							} else {
								gen7044 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen7044 = Call(__e, ShenFunc(symfail))

						}
						YaccParse := gen7044
						gen7185 := Call(__e, ShenFunc(symfail))

						gen7186 := Call(__e, ShenFunc(sym_a), YaccParse, gen7185)

						if True == gen7186 {
							gen7045 := Call(__e, ShenFunc(symshen_4_5bar_6), V1515)

							Parse__shen_4_5bar_6 := gen7045
							gen7053 := Call(__e, ShenFunc(symfail))

							gen7054 := Call(__e, ShenFunc(sym_a), gen7053, Parse__shen_4_5bar_6)

							gen7055 := Call(__e, ShenFunc(symnot), gen7054)

							var gen7056 Obj
							if True == gen7055 {
								gen7046 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5bar_6)

								Parse__shen_4_5st__input_6 := gen7046
								gen7050 := Call(__e, ShenFunc(symfail))

								gen7051 := Call(__e, ShenFunc(sym_a), gen7050, Parse__shen_4_5st__input_6)

								gen7052 := Call(__e, ShenFunc(symnot), gen7051)

								if True == gen7052 {
									gen7047 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

									gen7048 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

									gen7049 := Call(__e, ShenFunc(symcons), MakeSymbol("bar!"), gen7048)

									gen7056 = Call(__e, ShenFunc(symshen_4pair), gen7047, gen7049)

								} else {
									gen7056 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen7056 = Call(__e, ShenFunc(symfail))

							}
							YaccParse := gen7056
							gen7183 := Call(__e, ShenFunc(symfail))

							gen7184 := Call(__e, ShenFunc(sym_a), YaccParse, gen7183)

							if True == gen7184 {
								gen7057 := Call(__e, ShenFunc(symshen_4_5semicolon_6), V1515)

								Parse__shen_4_5semicolon_6 := gen7057
								gen7065 := Call(__e, ShenFunc(symfail))

								gen7066 := Call(__e, ShenFunc(sym_a), gen7065, Parse__shen_4_5semicolon_6)

								gen7067 := Call(__e, ShenFunc(symnot), gen7066)

								var gen7068 Obj
								if True == gen7067 {
									gen7058 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5semicolon_6)

									Parse__shen_4_5st__input_6 := gen7058
									gen7062 := Call(__e, ShenFunc(symfail))

									gen7063 := Call(__e, ShenFunc(sym_a), gen7062, Parse__shen_4_5st__input_6)

									gen7064 := Call(__e, ShenFunc(symnot), gen7063)

									if True == gen7064 {
										gen7059 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

										gen7060 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

										gen7061 := Call(__e, ShenFunc(symcons), MakeSymbol(";"), gen7060)

										gen7068 = Call(__e, ShenFunc(symshen_4pair), gen7059, gen7061)

									} else {
										gen7068 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen7068 = Call(__e, ShenFunc(symfail))

								}
								YaccParse := gen7068
								gen7181 := Call(__e, ShenFunc(symfail))

								gen7182 := Call(__e, ShenFunc(sym_a), YaccParse, gen7181)

								if True == gen7182 {
									gen7069 := Call(__e, ShenFunc(symshen_4_5colon_6), V1515)

									Parse__shen_4_5colon_6 := gen7069
									gen7081 := Call(__e, ShenFunc(symfail))

									gen7082 := Call(__e, ShenFunc(sym_a), gen7081, Parse__shen_4_5colon_6)

									gen7083 := Call(__e, ShenFunc(symnot), gen7082)

									var gen7084 Obj
									if True == gen7083 {
										gen7070 := Call(__e, ShenFunc(symshen_4_5equal_6), Parse__shen_4_5colon_6)

										Parse__shen_4_5equal_6 := gen7070
										gen7078 := Call(__e, ShenFunc(symfail))

										gen7079 := Call(__e, ShenFunc(sym_a), gen7078, Parse__shen_4_5equal_6)

										gen7080 := Call(__e, ShenFunc(symnot), gen7079)

										if True == gen7080 {
											gen7071 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5equal_6)

											Parse__shen_4_5st__input_6 := gen7071
											gen7075 := Call(__e, ShenFunc(symfail))

											gen7076 := Call(__e, ShenFunc(sym_a), gen7075, Parse__shen_4_5st__input_6)

											gen7077 := Call(__e, ShenFunc(symnot), gen7076)

											if True == gen7077 {
												gen7072 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

												gen7073 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

												gen7074 := Call(__e, ShenFunc(symcons), MakeSymbol(":="), gen7073)

												gen7084 = Call(__e, ShenFunc(symshen_4pair), gen7072, gen7074)

											} else {
												gen7084 = Call(__e, ShenFunc(symfail))

											}

										} else {
											gen7084 = Call(__e, ShenFunc(symfail))

										}

									} else {
										gen7084 = Call(__e, ShenFunc(symfail))

									}
									YaccParse := gen7084
									gen7179 := Call(__e, ShenFunc(symfail))

									gen7180 := Call(__e, ShenFunc(sym_a), YaccParse, gen7179)

									if True == gen7180 {
										gen7085 := Call(__e, ShenFunc(symshen_4_5colon_6), V1515)

										Parse__shen_4_5colon_6 := gen7085
										gen7097 := Call(__e, ShenFunc(symfail))

										gen7098 := Call(__e, ShenFunc(sym_a), gen7097, Parse__shen_4_5colon_6)

										gen7099 := Call(__e, ShenFunc(symnot), gen7098)

										var gen7100 Obj
										if True == gen7099 {
											gen7086 := Call(__e, ShenFunc(symshen_4_5minus_6), Parse__shen_4_5colon_6)

											Parse__shen_4_5minus_6 := gen7086
											gen7094 := Call(__e, ShenFunc(symfail))

											gen7095 := Call(__e, ShenFunc(sym_a), gen7094, Parse__shen_4_5minus_6)

											gen7096 := Call(__e, ShenFunc(symnot), gen7095)

											if True == gen7096 {
												gen7087 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5minus_6)

												Parse__shen_4_5st__input_6 := gen7087
												gen7091 := Call(__e, ShenFunc(symfail))

												gen7092 := Call(__e, ShenFunc(sym_a), gen7091, Parse__shen_4_5st__input_6)

												gen7093 := Call(__e, ShenFunc(symnot), gen7092)

												if True == gen7093 {
													gen7088 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

													gen7089 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

													gen7090 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen7089)

													gen7100 = Call(__e, ShenFunc(symshen_4pair), gen7088, gen7090)

												} else {
													gen7100 = Call(__e, ShenFunc(symfail))

												}

											} else {
												gen7100 = Call(__e, ShenFunc(symfail))

											}

										} else {
											gen7100 = Call(__e, ShenFunc(symfail))

										}
										YaccParse := gen7100
										gen7177 := Call(__e, ShenFunc(symfail))

										gen7178 := Call(__e, ShenFunc(sym_a), YaccParse, gen7177)

										if True == gen7178 {
											gen7101 := Call(__e, ShenFunc(symshen_4_5colon_6), V1515)

											Parse__shen_4_5colon_6 := gen7101
											gen7109 := Call(__e, ShenFunc(symfail))

											gen7110 := Call(__e, ShenFunc(sym_a), gen7109, Parse__shen_4_5colon_6)

											gen7111 := Call(__e, ShenFunc(symnot), gen7110)

											var gen7112 Obj
											if True == gen7111 {
												gen7102 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5colon_6)

												Parse__shen_4_5st__input_6 := gen7102
												gen7106 := Call(__e, ShenFunc(symfail))

												gen7107 := Call(__e, ShenFunc(sym_a), gen7106, Parse__shen_4_5st__input_6)

												gen7108 := Call(__e, ShenFunc(symnot), gen7107)

												if True == gen7108 {
													gen7103 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

													gen7104 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

													gen7105 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen7104)

													gen7112 = Call(__e, ShenFunc(symshen_4pair), gen7103, gen7105)

												} else {
													gen7112 = Call(__e, ShenFunc(symfail))

												}

											} else {
												gen7112 = Call(__e, ShenFunc(symfail))

											}
											YaccParse := gen7112
											gen7175 := Call(__e, ShenFunc(symfail))

											gen7176 := Call(__e, ShenFunc(sym_a), YaccParse, gen7175)

											if True == gen7176 {
												gen7113 := Call(__e, ShenFunc(symshen_4_5comma_6), V1515)

												Parse__shen_4_5comma_6 := gen7113
												gen7122 := Call(__e, ShenFunc(symfail))

												gen7123 := Call(__e, ShenFunc(sym_a), gen7122, Parse__shen_4_5comma_6)

												gen7124 := Call(__e, ShenFunc(symnot), gen7123)

												var gen7125 Obj
												if True == gen7124 {
													gen7114 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5comma_6)

													Parse__shen_4_5st__input_6 := gen7114
													gen7119 := Call(__e, ShenFunc(symfail))

													gen7120 := Call(__e, ShenFunc(sym_a), gen7119, Parse__shen_4_5st__input_6)

													gen7121 := Call(__e, ShenFunc(symnot), gen7120)

													if True == gen7121 {
														gen7115 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

														gen7116 := Call(__e, ShenFunc(symintern), MakeString(","))

														gen7117 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

														gen7118 := Call(__e, ShenFunc(symcons), gen7116, gen7117)

														gen7125 = Call(__e, ShenFunc(symshen_4pair), gen7115, gen7118)

													} else {
														gen7125 = Call(__e, ShenFunc(symfail))

													}

												} else {
													gen7125 = Call(__e, ShenFunc(symfail))

												}
												YaccParse := gen7125
												gen7173 := Call(__e, ShenFunc(symfail))

												gen7174 := Call(__e, ShenFunc(sym_a), YaccParse, gen7173)

												if True == gen7174 {
													gen7126 := Call(__e, ShenFunc(symshen_4_5comment_6), V1515)

													Parse__shen_4_5comment_6 := gen7126
													gen7133 := Call(__e, ShenFunc(symfail))

													gen7134 := Call(__e, ShenFunc(sym_a), gen7133, Parse__shen_4_5comment_6)

													gen7135 := Call(__e, ShenFunc(symnot), gen7134)

													var gen7136 Obj
													if True == gen7135 {
														gen7127 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5comment_6)

														Parse__shen_4_5st__input_6 := gen7127
														gen7130 := Call(__e, ShenFunc(symfail))

														gen7131 := Call(__e, ShenFunc(sym_a), gen7130, Parse__shen_4_5st__input_6)

														gen7132 := Call(__e, ShenFunc(symnot), gen7131)

														if True == gen7132 {
															gen7128 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

															gen7129 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

															gen7136 = Call(__e, ShenFunc(symshen_4pair), gen7128, gen7129)

														} else {
															gen7136 = Call(__e, ShenFunc(symfail))

														}

													} else {
														gen7136 = Call(__e, ShenFunc(symfail))

													}
													YaccParse := gen7136
													gen7171 := Call(__e, ShenFunc(symfail))

													gen7172 := Call(__e, ShenFunc(sym_a), YaccParse, gen7171)

													if True == gen7172 {
														gen7137 := Call(__e, ShenFunc(symshen_4_5atom_6), V1515)

														Parse__shen_4_5atom_6 := gen7137
														gen7147 := Call(__e, ShenFunc(symfail))

														gen7148 := Call(__e, ShenFunc(sym_a), gen7147, Parse__shen_4_5atom_6)

														gen7149 := Call(__e, ShenFunc(symnot), gen7148)

														var gen7150 Obj
														if True == gen7149 {
															gen7138 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5atom_6)

															Parse__shen_4_5st__input_6 := gen7138
															gen7144 := Call(__e, ShenFunc(symfail))

															gen7145 := Call(__e, ShenFunc(sym_a), gen7144, Parse__shen_4_5st__input_6)

															gen7146 := Call(__e, ShenFunc(symnot), gen7145)

															if True == gen7146 {
																gen7139 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

																gen7140 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5atom_6)

																gen7141 := Call(__e, ShenFunc(symmacroexpand), gen7140)

																gen7142 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

																gen7143 := Call(__e, ShenFunc(symcons), gen7141, gen7142)

																gen7150 = Call(__e, ShenFunc(symshen_4pair), gen7139, gen7143)

															} else {
																gen7150 = Call(__e, ShenFunc(symfail))

															}

														} else {
															gen7150 = Call(__e, ShenFunc(symfail))

														}
														YaccParse := gen7150
														gen7169 := Call(__e, ShenFunc(symfail))

														gen7170 := Call(__e, ShenFunc(sym_a), YaccParse, gen7169)

														if True == gen7170 {
															gen7151 := Call(__e, ShenFunc(symshen_4_5whitespaces_6), V1515)

															Parse__shen_4_5whitespaces_6 := gen7151
															gen7158 := Call(__e, ShenFunc(symfail))

															gen7159 := Call(__e, ShenFunc(sym_a), gen7158, Parse__shen_4_5whitespaces_6)

															gen7160 := Call(__e, ShenFunc(symnot), gen7159)

															var gen7161 Obj
															if True == gen7160 {
																gen7152 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5whitespaces_6)

																Parse__shen_4_5st__input_6 := gen7152
																gen7155 := Call(__e, ShenFunc(symfail))

																gen7156 := Call(__e, ShenFunc(sym_a), gen7155, Parse__shen_4_5st__input_6)

																gen7157 := Call(__e, ShenFunc(symnot), gen7156)

																if True == gen7157 {
																	gen7153 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

																	gen7154 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

																	gen7161 = Call(__e, ShenFunc(symshen_4pair), gen7153, gen7154)

																} else {
																	gen7161 = Call(__e, ShenFunc(symfail))

																}

															} else {
																gen7161 = Call(__e, ShenFunc(symfail))

															}
															YaccParse := gen7161
															gen7167 := Call(__e, ShenFunc(symfail))

															gen7168 := Call(__e, ShenFunc(sym_a), YaccParse, gen7167)

															if True == gen7168 {
																gen7162 := Call(__e, ShenFunc(sym_5e_6), V1515)

																Parse___5e_6 := gen7162
																gen7164 := Call(__e, ShenFunc(symfail))

																gen7165 := Call(__e, ShenFunc(sym_a), gen7164, Parse___5e_6)

																gen7166 := Call(__e, ShenFunc(symnot), gen7165)

																if True == gen7166 {
																	gen7163 := Call(__e, ShenFunc(symhd), Parse___5e_6)

																	__e.TailApply(ShenFunc(symshen_4pair), gen7163, Nil)

																	return

																} else {
																	__e.TailApply(ShenFunc(symfail))

																	return
																}

															} else {
																__e.Return(YaccParse)
																return
															}

														} else {
															__e.Return(YaccParse)
															return
														}

													} else {
														__e.Return(YaccParse)
														return
													}

												} else {
													__e.Return(YaccParse)
													return
												}

											} else {
												__e.Return(YaccParse)
												return
											}

										} else {
											__e.Return(YaccParse)
											return
										}

									} else {
										__e.Return(YaccParse)
										return
									}

								} else {
									__e.Return(YaccParse)
									return
								}

							} else {
								__e.Return(YaccParse)
								return
							}

						} else {
							__e.Return(YaccParse)
							return
						}

					} else {
						__e.Return(YaccParse)
						return
					}

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<st_input>"), gen7193)

		gen7203 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1518 := __args[0]
			_ = V1518
			gen7200 := Call(__e, ShenFunc(symhd), V1518)

			gen7201 := Call(__e, ShenFunc(symcons_2), gen7200)

			var gen7202 Obj
			if True == gen7201 {
				gen7198 := Call(__e, ShenFunc(symshen_4hdhd), V1518)

				gen7199 := Call(__e, ShenFunc(sym_a), MakeNumber(91), gen7198)

				if True == gen7199 {
					gen7202 = True
				} else {
					gen7202 = False
				}

			} else {
				gen7202 = False
			}
			if True == gen7202 {
				gen7194 := Call(__e, ShenFunc(symshen_4tlhd), V1518)

				gen7195 := Call(__e, ShenFunc(symshen_4hdtl), V1518)

				gen7196 := Call(__e, ShenFunc(symshen_4pair), gen7194, gen7195)

				NewStream1516 := gen7196
				gen7197 := Call(__e, ShenFunc(symhd), NewStream1516)

				__e.TailApply(ShenFunc(symshen_4pair), gen7197, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<lsb>"), gen7203)

		gen7213 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1521 := __args[0]
			_ = V1521
			gen7210 := Call(__e, ShenFunc(symhd), V1521)

			gen7211 := Call(__e, ShenFunc(symcons_2), gen7210)

			var gen7212 Obj
			if True == gen7211 {
				gen7208 := Call(__e, ShenFunc(symshen_4hdhd), V1521)

				gen7209 := Call(__e, ShenFunc(sym_a), MakeNumber(93), gen7208)

				if True == gen7209 {
					gen7212 = True
				} else {
					gen7212 = False
				}

			} else {
				gen7212 = False
			}
			if True == gen7212 {
				gen7204 := Call(__e, ShenFunc(symshen_4tlhd), V1521)

				gen7205 := Call(__e, ShenFunc(symshen_4hdtl), V1521)

				gen7206 := Call(__e, ShenFunc(symshen_4pair), gen7204, gen7205)

				NewStream1519 := gen7206
				gen7207 := Call(__e, ShenFunc(symhd), NewStream1519)

				__e.TailApply(ShenFunc(symshen_4pair), gen7207, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rsb>"), gen7213)

		gen7223 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1524 := __args[0]
			_ = V1524
			gen7220 := Call(__e, ShenFunc(symhd), V1524)

			gen7221 := Call(__e, ShenFunc(symcons_2), gen7220)

			var gen7222 Obj
			if True == gen7221 {
				gen7218 := Call(__e, ShenFunc(symshen_4hdhd), V1524)

				gen7219 := Call(__e, ShenFunc(sym_a), MakeNumber(123), gen7218)

				if True == gen7219 {
					gen7222 = True
				} else {
					gen7222 = False
				}

			} else {
				gen7222 = False
			}
			if True == gen7222 {
				gen7214 := Call(__e, ShenFunc(symshen_4tlhd), V1524)

				gen7215 := Call(__e, ShenFunc(symshen_4hdtl), V1524)

				gen7216 := Call(__e, ShenFunc(symshen_4pair), gen7214, gen7215)

				NewStream1522 := gen7216
				gen7217 := Call(__e, ShenFunc(symhd), NewStream1522)

				__e.TailApply(ShenFunc(symshen_4pair), gen7217, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<lcurly>"), gen7223)

		gen7233 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1527 := __args[0]
			_ = V1527
			gen7230 := Call(__e, ShenFunc(symhd), V1527)

			gen7231 := Call(__e, ShenFunc(symcons_2), gen7230)

			var gen7232 Obj
			if True == gen7231 {
				gen7228 := Call(__e, ShenFunc(symshen_4hdhd), V1527)

				gen7229 := Call(__e, ShenFunc(sym_a), MakeNumber(125), gen7228)

				if True == gen7229 {
					gen7232 = True
				} else {
					gen7232 = False
				}

			} else {
				gen7232 = False
			}
			if True == gen7232 {
				gen7224 := Call(__e, ShenFunc(symshen_4tlhd), V1527)

				gen7225 := Call(__e, ShenFunc(symshen_4hdtl), V1527)

				gen7226 := Call(__e, ShenFunc(symshen_4pair), gen7224, gen7225)

				NewStream1525 := gen7226
				gen7227 := Call(__e, ShenFunc(symhd), NewStream1525)

				__e.TailApply(ShenFunc(symshen_4pair), gen7227, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rcurly>"), gen7233)

		gen7243 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1530 := __args[0]
			_ = V1530
			gen7240 := Call(__e, ShenFunc(symhd), V1530)

			gen7241 := Call(__e, ShenFunc(symcons_2), gen7240)

			var gen7242 Obj
			if True == gen7241 {
				gen7238 := Call(__e, ShenFunc(symshen_4hdhd), V1530)

				gen7239 := Call(__e, ShenFunc(sym_a), MakeNumber(124), gen7238)

				if True == gen7239 {
					gen7242 = True
				} else {
					gen7242 = False
				}

			} else {
				gen7242 = False
			}
			if True == gen7242 {
				gen7234 := Call(__e, ShenFunc(symshen_4tlhd), V1530)

				gen7235 := Call(__e, ShenFunc(symshen_4hdtl), V1530)

				gen7236 := Call(__e, ShenFunc(symshen_4pair), gen7234, gen7235)

				NewStream1528 := gen7236
				gen7237 := Call(__e, ShenFunc(symhd), NewStream1528)

				__e.TailApply(ShenFunc(symshen_4pair), gen7237, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<bar>"), gen7243)

		gen7253 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1533 := __args[0]
			_ = V1533
			gen7250 := Call(__e, ShenFunc(symhd), V1533)

			gen7251 := Call(__e, ShenFunc(symcons_2), gen7250)

			var gen7252 Obj
			if True == gen7251 {
				gen7248 := Call(__e, ShenFunc(symshen_4hdhd), V1533)

				gen7249 := Call(__e, ShenFunc(sym_a), MakeNumber(59), gen7248)

				if True == gen7249 {
					gen7252 = True
				} else {
					gen7252 = False
				}

			} else {
				gen7252 = False
			}
			if True == gen7252 {
				gen7244 := Call(__e, ShenFunc(symshen_4tlhd), V1533)

				gen7245 := Call(__e, ShenFunc(symshen_4hdtl), V1533)

				gen7246 := Call(__e, ShenFunc(symshen_4pair), gen7244, gen7245)

				NewStream1531 := gen7246
				gen7247 := Call(__e, ShenFunc(symhd), NewStream1531)

				__e.TailApply(ShenFunc(symshen_4pair), gen7247, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<semicolon>"), gen7253)

		gen7263 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1536 := __args[0]
			_ = V1536
			gen7260 := Call(__e, ShenFunc(symhd), V1536)

			gen7261 := Call(__e, ShenFunc(symcons_2), gen7260)

			var gen7262 Obj
			if True == gen7261 {
				gen7258 := Call(__e, ShenFunc(symshen_4hdhd), V1536)

				gen7259 := Call(__e, ShenFunc(sym_a), MakeNumber(58), gen7258)

				if True == gen7259 {
					gen7262 = True
				} else {
					gen7262 = False
				}

			} else {
				gen7262 = False
			}
			if True == gen7262 {
				gen7254 := Call(__e, ShenFunc(symshen_4tlhd), V1536)

				gen7255 := Call(__e, ShenFunc(symshen_4hdtl), V1536)

				gen7256 := Call(__e, ShenFunc(symshen_4pair), gen7254, gen7255)

				NewStream1534 := gen7256
				gen7257 := Call(__e, ShenFunc(symhd), NewStream1534)

				__e.TailApply(ShenFunc(symshen_4pair), gen7257, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<colon>"), gen7263)

		gen7273 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1539 := __args[0]
			_ = V1539
			gen7270 := Call(__e, ShenFunc(symhd), V1539)

			gen7271 := Call(__e, ShenFunc(symcons_2), gen7270)

			var gen7272 Obj
			if True == gen7271 {
				gen7268 := Call(__e, ShenFunc(symshen_4hdhd), V1539)

				gen7269 := Call(__e, ShenFunc(sym_a), MakeNumber(44), gen7268)

				if True == gen7269 {
					gen7272 = True
				} else {
					gen7272 = False
				}

			} else {
				gen7272 = False
			}
			if True == gen7272 {
				gen7264 := Call(__e, ShenFunc(symshen_4tlhd), V1539)

				gen7265 := Call(__e, ShenFunc(symshen_4hdtl), V1539)

				gen7266 := Call(__e, ShenFunc(symshen_4pair), gen7264, gen7265)

				NewStream1537 := gen7266
				gen7267 := Call(__e, ShenFunc(symhd), NewStream1537)

				__e.TailApply(ShenFunc(symshen_4pair), gen7267, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<comma>"), gen7273)

		gen7283 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1542 := __args[0]
			_ = V1542
			gen7280 := Call(__e, ShenFunc(symhd), V1542)

			gen7281 := Call(__e, ShenFunc(symcons_2), gen7280)

			var gen7282 Obj
			if True == gen7281 {
				gen7278 := Call(__e, ShenFunc(symshen_4hdhd), V1542)

				gen7279 := Call(__e, ShenFunc(sym_a), MakeNumber(61), gen7278)

				if True == gen7279 {
					gen7282 = True
				} else {
					gen7282 = False
				}

			} else {
				gen7282 = False
			}
			if True == gen7282 {
				gen7274 := Call(__e, ShenFunc(symshen_4tlhd), V1542)

				gen7275 := Call(__e, ShenFunc(symshen_4hdtl), V1542)

				gen7276 := Call(__e, ShenFunc(symshen_4pair), gen7274, gen7275)

				NewStream1540 := gen7276
				gen7277 := Call(__e, ShenFunc(symhd), NewStream1540)

				__e.TailApply(ShenFunc(symshen_4pair), gen7277, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<equal>"), gen7283)

		gen7293 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1545 := __args[0]
			_ = V1545
			gen7290 := Call(__e, ShenFunc(symhd), V1545)

			gen7291 := Call(__e, ShenFunc(symcons_2), gen7290)

			var gen7292 Obj
			if True == gen7291 {
				gen7288 := Call(__e, ShenFunc(symshen_4hdhd), V1545)

				gen7289 := Call(__e, ShenFunc(sym_a), MakeNumber(45), gen7288)

				if True == gen7289 {
					gen7292 = True
				} else {
					gen7292 = False
				}

			} else {
				gen7292 = False
			}
			if True == gen7292 {
				gen7284 := Call(__e, ShenFunc(symshen_4tlhd), V1545)

				gen7285 := Call(__e, ShenFunc(symshen_4hdtl), V1545)

				gen7286 := Call(__e, ShenFunc(symshen_4pair), gen7284, gen7285)

				NewStream1543 := gen7286
				gen7287 := Call(__e, ShenFunc(symhd), NewStream1543)

				__e.TailApply(ShenFunc(symshen_4pair), gen7287, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<minus>"), gen7293)

		gen7303 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1548 := __args[0]
			_ = V1548
			gen7300 := Call(__e, ShenFunc(symhd), V1548)

			gen7301 := Call(__e, ShenFunc(symcons_2), gen7300)

			var gen7302 Obj
			if True == gen7301 {
				gen7298 := Call(__e, ShenFunc(symshen_4hdhd), V1548)

				gen7299 := Call(__e, ShenFunc(sym_a), MakeNumber(40), gen7298)

				if True == gen7299 {
					gen7302 = True
				} else {
					gen7302 = False
				}

			} else {
				gen7302 = False
			}
			if True == gen7302 {
				gen7294 := Call(__e, ShenFunc(symshen_4tlhd), V1548)

				gen7295 := Call(__e, ShenFunc(symshen_4hdtl), V1548)

				gen7296 := Call(__e, ShenFunc(symshen_4pair), gen7294, gen7295)

				NewStream1546 := gen7296
				gen7297 := Call(__e, ShenFunc(symhd), NewStream1546)

				__e.TailApply(ShenFunc(symshen_4pair), gen7297, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<lrb>"), gen7303)

		gen7313 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1551 := __args[0]
			_ = V1551
			gen7310 := Call(__e, ShenFunc(symhd), V1551)

			gen7311 := Call(__e, ShenFunc(symcons_2), gen7310)

			var gen7312 Obj
			if True == gen7311 {
				gen7308 := Call(__e, ShenFunc(symshen_4hdhd), V1551)

				gen7309 := Call(__e, ShenFunc(sym_a), MakeNumber(41), gen7308)

				if True == gen7309 {
					gen7312 = True
				} else {
					gen7312 = False
				}

			} else {
				gen7312 = False
			}
			if True == gen7312 {
				gen7304 := Call(__e, ShenFunc(symshen_4tlhd), V1551)

				gen7305 := Call(__e, ShenFunc(symshen_4hdtl), V1551)

				gen7306 := Call(__e, ShenFunc(symshen_4pair), gen7304, gen7305)

				NewStream1549 := gen7306
				gen7307 := Call(__e, ShenFunc(symhd), NewStream1549)

				__e.TailApply(ShenFunc(symshen_4pair), gen7307, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rrb>"), gen7313)

		gen7343 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1553 := __args[0]
			_ = V1553
			gen7314 := Call(__e, ShenFunc(symshen_4_5str_6), V1553)

			Parse__shen_4_5str_6 := gen7314
			gen7318 := Call(__e, ShenFunc(symfail))

			gen7319 := Call(__e, ShenFunc(sym_a), gen7318, Parse__shen_4_5str_6)

			gen7320 := Call(__e, ShenFunc(symnot), gen7319)

			var gen7321 Obj
			if True == gen7320 {
				gen7315 := Call(__e, ShenFunc(symhd), Parse__shen_4_5str_6)

				gen7316 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5str_6)

				gen7317 := Call(__e, ShenFunc(symshen_4control_1chars), gen7316)

				gen7321 = Call(__e, ShenFunc(symshen_4pair), gen7315, gen7317)

			} else {
				gen7321 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7321
			gen7341 := Call(__e, ShenFunc(symfail))

			gen7342 := Call(__e, ShenFunc(sym_a), YaccParse, gen7341)

			if True == gen7342 {
				gen7322 := Call(__e, ShenFunc(symshen_4_5number_6), V1553)

				Parse__shen_4_5number_6 := gen7322
				gen7325 := Call(__e, ShenFunc(symfail))

				gen7326 := Call(__e, ShenFunc(sym_a), gen7325, Parse__shen_4_5number_6)

				gen7327 := Call(__e, ShenFunc(symnot), gen7326)

				var gen7328 Obj
				if True == gen7327 {
					gen7323 := Call(__e, ShenFunc(symhd), Parse__shen_4_5number_6)

					gen7324 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5number_6)

					gen7328 = Call(__e, ShenFunc(symshen_4pair), gen7323, gen7324)

				} else {
					gen7328 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen7328
				gen7339 := Call(__e, ShenFunc(symfail))

				gen7340 := Call(__e, ShenFunc(sym_a), YaccParse, gen7339)

				if True == gen7340 {
					gen7329 := Call(__e, ShenFunc(symshen_4_5sym_6), V1553)

					Parse__shen_4_5sym_6 := gen7329
					gen7336 := Call(__e, ShenFunc(symfail))

					gen7337 := Call(__e, ShenFunc(sym_a), gen7336, Parse__shen_4_5sym_6)

					gen7338 := Call(__e, ShenFunc(symnot), gen7337)

					if True == gen7338 {
						gen7330 := Call(__e, ShenFunc(symhd), Parse__shen_4_5sym_6)

						gen7333 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5sym_6)

						gen7334 := Call(__e, ShenFunc(sym_a), gen7333, MakeString("<>"))

						var gen7335 Obj
						if True == gen7334 {
							gen7332 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

							gen7335 = Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen7332)

						} else {
							gen7331 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5sym_6)

							gen7335 = Call(__e, ShenFunc(symintern), gen7331)

						}
						__e.TailApply(ShenFunc(symshen_4pair), gen7330, gen7335)

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<atom>"), gen7343)

		gen7369 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1555 := __args[0]
			_ = V1555
			gen7368 := Call(__e, ShenFunc(sym_a), Nil, V1555)

			if True == gen7368 {
				__e.Return(MakeString(""))
				return
			} else {
				gen7366 := Call(__e, ShenFunc(symcons_2), V1555)

				var gen7367 Obj
				if True == gen7366 {
					gen7363 := Call(__e, ShenFunc(symhd), V1555)

					gen7364 := Call(__e, ShenFunc(sym_a), MakeString("c"), gen7363)

					var gen7365 Obj
					if True == gen7364 {
						gen7360 := Call(__e, ShenFunc(symtl), V1555)

						gen7361 := Call(__e, ShenFunc(symcons_2), gen7360)

						var gen7362 Obj
						if True == gen7361 {
							gen7357 := Call(__e, ShenFunc(symtl), V1555)

							gen7358 := Call(__e, ShenFunc(symhd), gen7357)

							gen7359 := Call(__e, ShenFunc(sym_a), MakeString("#"), gen7358)

							if True == gen7359 {
								gen7362 = True
							} else {
								gen7362 = False
							}

						} else {
							gen7362 = False
						}
						if True == gen7362 {
							gen7365 = True
						} else {
							gen7365 = False
						}

					} else {
						gen7365 = False
					}
					if True == gen7365 {
						gen7367 = True
					} else {
						gen7367 = False
					}

				} else {
					gen7367 = False
				}
				if True == gen7367 {
					gen7348 := Call(__e, ShenFunc(symtl), V1555)

					gen7349 := Call(__e, ShenFunc(symtl), gen7348)

					gen7350 := Call(__e, ShenFunc(symshen_4code_1point), gen7349)

					CodePoint := gen7350
					gen7351 := Call(__e, ShenFunc(symtl), V1555)

					gen7352 := Call(__e, ShenFunc(symtl), gen7351)

					gen7353 := Call(__e, ShenFunc(symshen_4after_1codepoint), gen7352)

					AfterCodePoint := gen7353
					gen7354 := Call(__e, ShenFunc(symshen_4decimalise), CodePoint)

					gen7355 := Call(__e, ShenFunc(symn_1_6string), gen7354)

					gen7356 := Call(__e, ShenFunc(symshen_4control_1chars), AfterCodePoint)

					__e.TailApply(ShenFunc(sym_8s), gen7355, gen7356)

					return

				} else {
					gen7347 := Call(__e, ShenFunc(symcons_2), V1555)

					if True == gen7347 {
						gen7344 := Call(__e, ShenFunc(symhd), V1555)

						gen7345 := Call(__e, ShenFunc(symtl), V1555)

						gen7346 := Call(__e, ShenFunc(symshen_4control_1chars), gen7345)

						__e.TailApply(ShenFunc(sym_8s), gen7344, gen7346)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.control-chars"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.control-chars"), gen7369)

		gen7394 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1559 := __args[0]
			_ = V1559
			gen7392 := Call(__e, ShenFunc(symcons_2), V1559)

			var gen7393 Obj
			if True == gen7392 {
				gen7390 := Call(__e, ShenFunc(symhd), V1559)

				gen7391 := Call(__e, ShenFunc(sym_a), MakeString(";"), gen7390)

				if True == gen7391 {
					gen7393 = True
				} else {
					gen7393 = False
				}

			} else {
				gen7393 = False
			}
			if True == gen7393 {
				__e.Return(MakeString(""))
				return
			} else {
				gen7388 := Call(__e, ShenFunc(symcons_2), V1559)

				var gen7389 Obj
				if True == gen7388 {
					gen7375 := Call(__e, ShenFunc(symhd), V1559)

					gen7376 := Call(__e, ShenFunc(symcons), MakeString("0"), Nil)

					gen7377 := Call(__e, ShenFunc(symcons), MakeString("9"), gen7376)

					gen7378 := Call(__e, ShenFunc(symcons), MakeString("8"), gen7377)

					gen7379 := Call(__e, ShenFunc(symcons), MakeString("7"), gen7378)

					gen7380 := Call(__e, ShenFunc(symcons), MakeString("6"), gen7379)

					gen7381 := Call(__e, ShenFunc(symcons), MakeString("5"), gen7380)

					gen7382 := Call(__e, ShenFunc(symcons), MakeString("4"), gen7381)

					gen7383 := Call(__e, ShenFunc(symcons), MakeString("3"), gen7382)

					gen7384 := Call(__e, ShenFunc(symcons), MakeString("2"), gen7383)

					gen7385 := Call(__e, ShenFunc(symcons), MakeString("1"), gen7384)

					gen7386 := Call(__e, ShenFunc(symcons), MakeString("0"), gen7385)

					gen7387 := Call(__e, ShenFunc(symelement_2), gen7375, gen7386)

					if True == gen7387 {
						gen7389 = True
					} else {
						gen7389 = False
					}

				} else {
					gen7389 = False
				}
				if True == gen7389 {
					gen7372 := Call(__e, ShenFunc(symhd), V1559)

					gen7373 := Call(__e, ShenFunc(symtl), V1559)

					gen7374 := Call(__e, ShenFunc(symshen_4code_1point), gen7373)

					__e.TailApply(ShenFunc(symcons), gen7372, gen7374)

					return

				} else {
					gen7370 := Call(__e, ShenFunc(symshen_4app), V1559, MakeString("\n"), MakeSymbol("shen.a"))

					gen7371 := Call(__e, ShenFunc(symcn), MakeString("code point parse error "), gen7370)

					__e.TailApply(ShenFunc(symsimple_1error), gen7371)

					return

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.code-point"), gen7394)

		gen7402 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1565 := __args[0]
			_ = V1565
			gen7401 := Call(__e, ShenFunc(sym_a), Nil, V1565)

			if True == gen7401 {
				__e.Return(Nil)
				return
			} else {
				gen7399 := Call(__e, ShenFunc(symcons_2), V1565)

				var gen7400 Obj
				if True == gen7399 {
					gen7397 := Call(__e, ShenFunc(symhd), V1565)

					gen7398 := Call(__e, ShenFunc(sym_a), MakeString(";"), gen7397)

					if True == gen7398 {
						gen7400 = True
					} else {
						gen7400 = False
					}

				} else {
					gen7400 = False
				}
				if True == gen7400 {
					__e.TailApply(ShenFunc(symtl), V1565)

					return
				} else {
					gen7396 := Call(__e, ShenFunc(symcons_2), V1565)

					if True == gen7396 {
						gen7395 := Call(__e, ShenFunc(symtl), V1565)

						__e.TailApply(ShenFunc(symshen_4after_1codepoint), gen7395)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.after-codepoint"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.after-codepoint"), gen7402)

		gen7405 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1567 := __args[0]
			_ = V1567
			gen7403 := Call(__e, ShenFunc(symshen_4digits_1_6integers), V1567)

			gen7404 := Call(__e, ShenFunc(symreverse), gen7403)

			__e.TailApply(ShenFunc(symshen_4pre), gen7404, MakeNumber(0))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.decimalise"), gen7405)

		gen7466 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1573 := __args[0]
			_ = V1573
			gen7464 := Call(__e, ShenFunc(symcons_2), V1573)

			var gen7465 Obj
			if True == gen7464 {
				gen7462 := Call(__e, ShenFunc(symhd), V1573)

				gen7463 := Call(__e, ShenFunc(sym_a), MakeString("0"), gen7462)

				if True == gen7463 {
					gen7465 = True
				} else {
					gen7465 = False
				}

			} else {
				gen7465 = False
			}
			if True == gen7465 {
				gen7460 := Call(__e, ShenFunc(symtl), V1573)

				gen7461 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7460)

				__e.TailApply(ShenFunc(symcons), MakeNumber(0), gen7461)

				return

			} else {
				gen7458 := Call(__e, ShenFunc(symcons_2), V1573)

				var gen7459 Obj
				if True == gen7458 {
					gen7456 := Call(__e, ShenFunc(symhd), V1573)

					gen7457 := Call(__e, ShenFunc(sym_a), MakeString("1"), gen7456)

					if True == gen7457 {
						gen7459 = True
					} else {
						gen7459 = False
					}

				} else {
					gen7459 = False
				}
				if True == gen7459 {
					gen7454 := Call(__e, ShenFunc(symtl), V1573)

					gen7455 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7454)

					__e.TailApply(ShenFunc(symcons), MakeNumber(1), gen7455)

					return

				} else {
					gen7452 := Call(__e, ShenFunc(symcons_2), V1573)

					var gen7453 Obj
					if True == gen7452 {
						gen7450 := Call(__e, ShenFunc(symhd), V1573)

						gen7451 := Call(__e, ShenFunc(sym_a), MakeString("2"), gen7450)

						if True == gen7451 {
							gen7453 = True
						} else {
							gen7453 = False
						}

					} else {
						gen7453 = False
					}
					if True == gen7453 {
						gen7448 := Call(__e, ShenFunc(symtl), V1573)

						gen7449 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7448)

						__e.TailApply(ShenFunc(symcons), MakeNumber(2), gen7449)

						return

					} else {
						gen7446 := Call(__e, ShenFunc(symcons_2), V1573)

						var gen7447 Obj
						if True == gen7446 {
							gen7444 := Call(__e, ShenFunc(symhd), V1573)

							gen7445 := Call(__e, ShenFunc(sym_a), MakeString("3"), gen7444)

							if True == gen7445 {
								gen7447 = True
							} else {
								gen7447 = False
							}

						} else {
							gen7447 = False
						}
						if True == gen7447 {
							gen7442 := Call(__e, ShenFunc(symtl), V1573)

							gen7443 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7442)

							__e.TailApply(ShenFunc(symcons), MakeNumber(3), gen7443)

							return

						} else {
							gen7440 := Call(__e, ShenFunc(symcons_2), V1573)

							var gen7441 Obj
							if True == gen7440 {
								gen7438 := Call(__e, ShenFunc(symhd), V1573)

								gen7439 := Call(__e, ShenFunc(sym_a), MakeString("4"), gen7438)

								if True == gen7439 {
									gen7441 = True
								} else {
									gen7441 = False
								}

							} else {
								gen7441 = False
							}
							if True == gen7441 {
								gen7436 := Call(__e, ShenFunc(symtl), V1573)

								gen7437 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7436)

								__e.TailApply(ShenFunc(symcons), MakeNumber(4), gen7437)

								return

							} else {
								gen7434 := Call(__e, ShenFunc(symcons_2), V1573)

								var gen7435 Obj
								if True == gen7434 {
									gen7432 := Call(__e, ShenFunc(symhd), V1573)

									gen7433 := Call(__e, ShenFunc(sym_a), MakeString("5"), gen7432)

									if True == gen7433 {
										gen7435 = True
									} else {
										gen7435 = False
									}

								} else {
									gen7435 = False
								}
								if True == gen7435 {
									gen7430 := Call(__e, ShenFunc(symtl), V1573)

									gen7431 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7430)

									__e.TailApply(ShenFunc(symcons), MakeNumber(5), gen7431)

									return

								} else {
									gen7428 := Call(__e, ShenFunc(symcons_2), V1573)

									var gen7429 Obj
									if True == gen7428 {
										gen7426 := Call(__e, ShenFunc(symhd), V1573)

										gen7427 := Call(__e, ShenFunc(sym_a), MakeString("6"), gen7426)

										if True == gen7427 {
											gen7429 = True
										} else {
											gen7429 = False
										}

									} else {
										gen7429 = False
									}
									if True == gen7429 {
										gen7424 := Call(__e, ShenFunc(symtl), V1573)

										gen7425 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7424)

										__e.TailApply(ShenFunc(symcons), MakeNumber(6), gen7425)

										return

									} else {
										gen7422 := Call(__e, ShenFunc(symcons_2), V1573)

										var gen7423 Obj
										if True == gen7422 {
											gen7420 := Call(__e, ShenFunc(symhd), V1573)

											gen7421 := Call(__e, ShenFunc(sym_a), MakeString("7"), gen7420)

											if True == gen7421 {
												gen7423 = True
											} else {
												gen7423 = False
											}

										} else {
											gen7423 = False
										}
										if True == gen7423 {
											gen7418 := Call(__e, ShenFunc(symtl), V1573)

											gen7419 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7418)

											__e.TailApply(ShenFunc(symcons), MakeNumber(7), gen7419)

											return

										} else {
											gen7416 := Call(__e, ShenFunc(symcons_2), V1573)

											var gen7417 Obj
											if True == gen7416 {
												gen7414 := Call(__e, ShenFunc(symhd), V1573)

												gen7415 := Call(__e, ShenFunc(sym_a), MakeString("8"), gen7414)

												if True == gen7415 {
													gen7417 = True
												} else {
													gen7417 = False
												}

											} else {
												gen7417 = False
											}
											if True == gen7417 {
												gen7412 := Call(__e, ShenFunc(symtl), V1573)

												gen7413 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7412)

												__e.TailApply(ShenFunc(symcons), MakeNumber(8), gen7413)

												return

											} else {
												gen7410 := Call(__e, ShenFunc(symcons_2), V1573)

												var gen7411 Obj
												if True == gen7410 {
													gen7408 := Call(__e, ShenFunc(symhd), V1573)

													gen7409 := Call(__e, ShenFunc(sym_a), MakeString("9"), gen7408)

													if True == gen7409 {
														gen7411 = True
													} else {
														gen7411 = False
													}

												} else {
													gen7411 = False
												}
												if True == gen7411 {
													gen7406 := Call(__e, ShenFunc(symtl), V1573)

													gen7407 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen7406)

													__e.TailApply(ShenFunc(symcons), MakeNumber(9), gen7407)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.digits->integers"), gen7466)

		gen7479 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1575 := __args[0]
			_ = V1575
			gen7467 := Call(__e, ShenFunc(symshen_4_5alpha_6), V1575)

			Parse__shen_4_5alpha_6 := gen7467
			gen7476 := Call(__e, ShenFunc(symfail))

			gen7477 := Call(__e, ShenFunc(sym_a), gen7476, Parse__shen_4_5alpha_6)

			gen7478 := Call(__e, ShenFunc(symnot), gen7477)

			if True == gen7478 {
				gen7468 := Call(__e, ShenFunc(symshen_4_5alphanums_6), Parse__shen_4_5alpha_6)

				Parse__shen_4_5alphanums_6 := gen7468
				gen7473 := Call(__e, ShenFunc(symfail))

				gen7474 := Call(__e, ShenFunc(sym_a), gen7473, Parse__shen_4_5alphanums_6)

				gen7475 := Call(__e, ShenFunc(symnot), gen7474)

				if True == gen7475 {
					gen7469 := Call(__e, ShenFunc(symhd), Parse__shen_4_5alphanums_6)

					gen7470 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alpha_6)

					gen7471 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alphanums_6)

					gen7472 := Call(__e, ShenFunc(sym_8s), gen7470, gen7471)

					__e.TailApply(ShenFunc(symshen_4pair), gen7469, gen7472)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<sym>"), gen7479)

		gen7500 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1577 := __args[0]
			_ = V1577
			gen7480 := Call(__e, ShenFunc(symshen_4_5alphanum_6), V1577)

			Parse__shen_4_5alphanum_6 := gen7480
			gen7489 := Call(__e, ShenFunc(symfail))

			gen7490 := Call(__e, ShenFunc(sym_a), gen7489, Parse__shen_4_5alphanum_6)

			gen7491 := Call(__e, ShenFunc(symnot), gen7490)

			var gen7492 Obj
			if True == gen7491 {
				gen7481 := Call(__e, ShenFunc(symshen_4_5alphanums_6), Parse__shen_4_5alphanum_6)

				Parse__shen_4_5alphanums_6 := gen7481
				gen7486 := Call(__e, ShenFunc(symfail))

				gen7487 := Call(__e, ShenFunc(sym_a), gen7486, Parse__shen_4_5alphanums_6)

				gen7488 := Call(__e, ShenFunc(symnot), gen7487)

				if True == gen7488 {
					gen7482 := Call(__e, ShenFunc(symhd), Parse__shen_4_5alphanums_6)

					gen7483 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alphanum_6)

					gen7484 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alphanums_6)

					gen7485 := Call(__e, ShenFunc(sym_8s), gen7483, gen7484)

					gen7492 = Call(__e, ShenFunc(symshen_4pair), gen7482, gen7485)

				} else {
					gen7492 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen7492 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7492
			gen7498 := Call(__e, ShenFunc(symfail))

			gen7499 := Call(__e, ShenFunc(sym_a), YaccParse, gen7498)

			if True == gen7499 {
				gen7493 := Call(__e, ShenFunc(sym_5e_6), V1577)

				Parse___5e_6 := gen7493
				gen7495 := Call(__e, ShenFunc(symfail))

				gen7496 := Call(__e, ShenFunc(sym_a), gen7495, Parse___5e_6)

				gen7497 := Call(__e, ShenFunc(symnot), gen7496)

				if True == gen7497 {
					gen7494 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen7494, MakeString(""))

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<alphanums>"), gen7500)

		gen7516 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1579 := __args[0]
			_ = V1579
			gen7501 := Call(__e, ShenFunc(symshen_4_5alpha_6), V1579)

			Parse__shen_4_5alpha_6 := gen7501
			gen7504 := Call(__e, ShenFunc(symfail))

			gen7505 := Call(__e, ShenFunc(sym_a), gen7504, Parse__shen_4_5alpha_6)

			gen7506 := Call(__e, ShenFunc(symnot), gen7505)

			var gen7507 Obj
			if True == gen7506 {
				gen7502 := Call(__e, ShenFunc(symhd), Parse__shen_4_5alpha_6)

				gen7503 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alpha_6)

				gen7507 = Call(__e, ShenFunc(symshen_4pair), gen7502, gen7503)

			} else {
				gen7507 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7507
			gen7514 := Call(__e, ShenFunc(symfail))

			gen7515 := Call(__e, ShenFunc(sym_a), YaccParse, gen7514)

			if True == gen7515 {
				gen7508 := Call(__e, ShenFunc(symshen_4_5num_6), V1579)

				Parse__shen_4_5num_6 := gen7508
				gen7511 := Call(__e, ShenFunc(symfail))

				gen7512 := Call(__e, ShenFunc(sym_a), gen7511, Parse__shen_4_5num_6)

				gen7513 := Call(__e, ShenFunc(symnot), gen7512)

				if True == gen7513 {
					gen7509 := Call(__e, ShenFunc(symhd), Parse__shen_4_5num_6)

					gen7510 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5num_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen7509, gen7510)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<alphanum>"), gen7516)

		gen7526 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1581 := __args[0]
			_ = V1581
			gen7524 := Call(__e, ShenFunc(symhd), V1581)

			gen7525 := Call(__e, ShenFunc(symcons_2), gen7524)

			if True == gen7525 {
				gen7517 := Call(__e, ShenFunc(symshen_4hdhd), V1581)

				Parse__Char := gen7517
				gen7523 := Call(__e, ShenFunc(symshen_4numbyte_2), Parse__Char)

				if True == gen7523 {
					gen7518 := Call(__e, ShenFunc(symshen_4tlhd), V1581)

					gen7519 := Call(__e, ShenFunc(symshen_4hdtl), V1581)

					gen7520 := Call(__e, ShenFunc(symshen_4pair), gen7518, gen7519)

					gen7521 := Call(__e, ShenFunc(symhd), gen7520)

					gen7522 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

					__e.TailApply(ShenFunc(symshen_4pair), gen7521, gen7522)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<num>"), gen7526)

		gen7537 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1587 := __args[0]
			_ = V1587
			gen7536 := Call(__e, ShenFunc(sym_a), MakeNumber(48), V1587)

			if True == gen7536 {
				__e.Return(True)
				return
			} else {
				gen7535 := Call(__e, ShenFunc(sym_a), MakeNumber(49), V1587)

				if True == gen7535 {
					__e.Return(True)
					return
				} else {
					gen7534 := Call(__e, ShenFunc(sym_a), MakeNumber(50), V1587)

					if True == gen7534 {
						__e.Return(True)
						return
					} else {
						gen7533 := Call(__e, ShenFunc(sym_a), MakeNumber(51), V1587)

						if True == gen7533 {
							__e.Return(True)
							return
						} else {
							gen7532 := Call(__e, ShenFunc(sym_a), MakeNumber(52), V1587)

							if True == gen7532 {
								__e.Return(True)
								return
							} else {
								gen7531 := Call(__e, ShenFunc(sym_a), MakeNumber(53), V1587)

								if True == gen7531 {
									__e.Return(True)
									return
								} else {
									gen7530 := Call(__e, ShenFunc(sym_a), MakeNumber(54), V1587)

									if True == gen7530 {
										__e.Return(True)
										return
									} else {
										gen7529 := Call(__e, ShenFunc(sym_a), MakeNumber(55), V1587)

										if True == gen7529 {
											__e.Return(True)
											return
										} else {
											gen7528 := Call(__e, ShenFunc(sym_a), MakeNumber(56), V1587)

											if True == gen7528 {
												__e.Return(True)
												return
											} else {
												gen7527 := Call(__e, ShenFunc(sym_a), MakeNumber(57), V1587)

												if True == gen7527 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.numbyte?"), gen7537)

		gen7547 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1589 := __args[0]
			_ = V1589
			gen7545 := Call(__e, ShenFunc(symhd), V1589)

			gen7546 := Call(__e, ShenFunc(symcons_2), gen7545)

			if True == gen7546 {
				gen7538 := Call(__e, ShenFunc(symshen_4hdhd), V1589)

				Parse__Char := gen7538
				gen7544 := Call(__e, ShenFunc(symshen_4symbol_1code_2), Parse__Char)

				if True == gen7544 {
					gen7539 := Call(__e, ShenFunc(symshen_4tlhd), V1589)

					gen7540 := Call(__e, ShenFunc(symshen_4hdtl), V1589)

					gen7541 := Call(__e, ShenFunc(symshen_4pair), gen7539, gen7540)

					gen7542 := Call(__e, ShenFunc(symhd), gen7541)

					gen7543 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

					__e.TailApply(ShenFunc(symshen_4pair), gen7542, gen7543)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<alpha>"), gen7547)

		gen7569 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1591 := __args[0]
			_ = V1591
			gen7568 := Call(__e, ShenFunc(sym_a), V1591, MakeNumber(126))

			if True == gen7568 {
				__e.Return(True)
				return
			} else {
				gen7565 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(94))

				var gen7566 Obj
				if True == gen7565 {
					gen7564 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(123))

					if True == gen7564 {
						gen7566 = True
					} else {
						gen7566 = False
					}

				} else {
					gen7566 = False
				}
				var gen7567 Obj
				if True == gen7566 {
					gen7567 = True
				} else {
					gen7561 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(59))

					var gen7562 Obj
					if True == gen7561 {
						gen7560 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(91))

						if True == gen7560 {
							gen7562 = True
						} else {
							gen7562 = False
						}

					} else {
						gen7562 = False
					}
					var gen7563 Obj
					if True == gen7562 {
						gen7563 = True
					} else {
						gen7557 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(41))

						var gen7558 Obj
						if True == gen7557 {
							gen7555 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(58))

							var gen7556 Obj
							if True == gen7555 {
								gen7553 := Call(__e, ShenFunc(sym_a), V1591, MakeNumber(44))

								gen7554 := Call(__e, ShenFunc(symnot), gen7553)

								if True == gen7554 {
									gen7556 = True
								} else {
									gen7556 = False
								}

							} else {
								gen7556 = False
							}
							if True == gen7556 {
								gen7558 = True
							} else {
								gen7558 = False
							}

						} else {
							gen7558 = False
						}
						var gen7559 Obj
						if True == gen7558 {
							gen7559 = True
						} else {
							gen7550 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(34))

							var gen7551 Obj
							if True == gen7550 {
								gen7549 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(40))

								if True == gen7549 {
									gen7551 = True
								} else {
									gen7551 = False
								}

							} else {
								gen7551 = False
							}
							var gen7552 Obj
							if True == gen7551 {
								gen7552 = True
							} else {
								gen7548 := Call(__e, ShenFunc(sym_a), V1591, MakeNumber(33))

								if True == gen7548 {
									gen7552 = True
								} else {
									gen7552 = False
								}

							}
							if True == gen7552 {
								gen7559 = True
							} else {
								gen7559 = False
							}

						}
						if True == gen7559 {
							gen7563 = True
						} else {
							gen7563 = False
						}

					}
					if True == gen7563 {
						gen7567 = True
					} else {
						gen7567 = False
					}

				}
				if True == gen7567 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.symbol-code?"), gen7569)

		gen7584 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1593 := __args[0]
			_ = V1593
			gen7570 := Call(__e, ShenFunc(symshen_4_5dbq_6), V1593)

			Parse__shen_4_5dbq_6 := gen7570
			gen7581 := Call(__e, ShenFunc(symfail))

			gen7582 := Call(__e, ShenFunc(sym_a), gen7581, Parse__shen_4_5dbq_6)

			gen7583 := Call(__e, ShenFunc(symnot), gen7582)

			if True == gen7583 {
				gen7571 := Call(__e, ShenFunc(symshen_4_5strcontents_6), Parse__shen_4_5dbq_6)

				Parse__shen_4_5strcontents_6 := gen7571
				gen7578 := Call(__e, ShenFunc(symfail))

				gen7579 := Call(__e, ShenFunc(sym_a), gen7578, Parse__shen_4_5strcontents_6)

				gen7580 := Call(__e, ShenFunc(symnot), gen7579)

				if True == gen7580 {
					gen7572 := Call(__e, ShenFunc(symshen_4_5dbq_6), Parse__shen_4_5strcontents_6)

					Parse__shen_4_5dbq_6 := gen7572
					gen7575 := Call(__e, ShenFunc(symfail))

					gen7576 := Call(__e, ShenFunc(sym_a), gen7575, Parse__shen_4_5dbq_6)

					gen7577 := Call(__e, ShenFunc(symnot), gen7576)

					if True == gen7577 {
						gen7573 := Call(__e, ShenFunc(symhd), Parse__shen_4_5dbq_6)

						gen7574 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5strcontents_6)

						__e.TailApply(ShenFunc(symshen_4pair), gen7573, gen7574)

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<str>"), gen7584)

		gen7593 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1595 := __args[0]
			_ = V1595
			gen7591 := Call(__e, ShenFunc(symhd), V1595)

			gen7592 := Call(__e, ShenFunc(symcons_2), gen7591)

			if True == gen7592 {
				gen7585 := Call(__e, ShenFunc(symshen_4hdhd), V1595)

				Parse__Char := gen7585
				gen7590 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(34))

				if True == gen7590 {
					gen7586 := Call(__e, ShenFunc(symshen_4tlhd), V1595)

					gen7587 := Call(__e, ShenFunc(symshen_4hdtl), V1595)

					gen7588 := Call(__e, ShenFunc(symshen_4pair), gen7586, gen7587)

					gen7589 := Call(__e, ShenFunc(symhd), gen7588)

					__e.TailApply(ShenFunc(symshen_4pair), gen7589, Parse__Char)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<dbq>"), gen7593)

		gen7614 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1597 := __args[0]
			_ = V1597
			gen7594 := Call(__e, ShenFunc(symshen_4_5strc_6), V1597)

			Parse__shen_4_5strc_6 := gen7594
			gen7603 := Call(__e, ShenFunc(symfail))

			gen7604 := Call(__e, ShenFunc(sym_a), gen7603, Parse__shen_4_5strc_6)

			gen7605 := Call(__e, ShenFunc(symnot), gen7604)

			var gen7606 Obj
			if True == gen7605 {
				gen7595 := Call(__e, ShenFunc(symshen_4_5strcontents_6), Parse__shen_4_5strc_6)

				Parse__shen_4_5strcontents_6 := gen7595
				gen7600 := Call(__e, ShenFunc(symfail))

				gen7601 := Call(__e, ShenFunc(sym_a), gen7600, Parse__shen_4_5strcontents_6)

				gen7602 := Call(__e, ShenFunc(symnot), gen7601)

				if True == gen7602 {
					gen7596 := Call(__e, ShenFunc(symhd), Parse__shen_4_5strcontents_6)

					gen7597 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5strc_6)

					gen7598 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5strcontents_6)

					gen7599 := Call(__e, ShenFunc(symcons), gen7597, gen7598)

					gen7606 = Call(__e, ShenFunc(symshen_4pair), gen7596, gen7599)

				} else {
					gen7606 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen7606 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7606
			gen7612 := Call(__e, ShenFunc(symfail))

			gen7613 := Call(__e, ShenFunc(sym_a), YaccParse, gen7612)

			if True == gen7613 {
				gen7607 := Call(__e, ShenFunc(sym_5e_6), V1597)

				Parse___5e_6 := gen7607
				gen7609 := Call(__e, ShenFunc(symfail))

				gen7610 := Call(__e, ShenFunc(sym_a), gen7609, Parse___5e_6)

				gen7611 := Call(__e, ShenFunc(symnot), gen7610)

				if True == gen7611 {
					gen7608 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen7608, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<strcontents>"), gen7614)

		gen7623 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1599 := __args[0]
			_ = V1599
			gen7621 := Call(__e, ShenFunc(symhd), V1599)

			gen7622 := Call(__e, ShenFunc(symcons_2), gen7621)

			if True == gen7622 {
				gen7615 := Call(__e, ShenFunc(symshen_4hdhd), V1599)

				Parse__Char := gen7615
				gen7616 := Call(__e, ShenFunc(symshen_4tlhd), V1599)

				gen7617 := Call(__e, ShenFunc(symshen_4hdtl), V1599)

				gen7618 := Call(__e, ShenFunc(symshen_4pair), gen7616, gen7617)

				gen7619 := Call(__e, ShenFunc(symhd), gen7618)

				gen7620 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

				__e.TailApply(ShenFunc(symshen_4pair), gen7619, gen7620)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<byte>"), gen7623)

		gen7634 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1601 := __args[0]
			_ = V1601
			gen7632 := Call(__e, ShenFunc(symhd), V1601)

			gen7633 := Call(__e, ShenFunc(symcons_2), gen7632)

			if True == gen7633 {
				gen7624 := Call(__e, ShenFunc(symshen_4hdhd), V1601)

				Parse__Char := gen7624
				gen7630 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(34))

				gen7631 := Call(__e, ShenFunc(symnot), gen7630)

				if True == gen7631 {
					gen7625 := Call(__e, ShenFunc(symshen_4tlhd), V1601)

					gen7626 := Call(__e, ShenFunc(symshen_4hdtl), V1601)

					gen7627 := Call(__e, ShenFunc(symshen_4pair), gen7625, gen7626)

					gen7628 := Call(__e, ShenFunc(symhd), gen7627)

					gen7629 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

					__e.TailApply(ShenFunc(symshen_4pair), gen7628, gen7629)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<strc>"), gen7634)

		gen7747 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1603 := __args[0]
			_ = V1603
			gen7635 := Call(__e, ShenFunc(symshen_4_5minus_6), V1603)

			Parse__shen_4_5minus_6 := gen7635
			gen7643 := Call(__e, ShenFunc(symfail))

			gen7644 := Call(__e, ShenFunc(sym_a), gen7643, Parse__shen_4_5minus_6)

			gen7645 := Call(__e, ShenFunc(symnot), gen7644)

			var gen7646 Obj
			if True == gen7645 {
				gen7636 := Call(__e, ShenFunc(symshen_4_5number_6), Parse__shen_4_5minus_6)

				Parse__shen_4_5number_6 := gen7636
				gen7640 := Call(__e, ShenFunc(symfail))

				gen7641 := Call(__e, ShenFunc(sym_a), gen7640, Parse__shen_4_5number_6)

				gen7642 := Call(__e, ShenFunc(symnot), gen7641)

				if True == gen7642 {
					gen7637 := Call(__e, ShenFunc(symhd), Parse__shen_4_5number_6)

					gen7638 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5number_6)

					gen7639 := Call(__e, ShenFunc(sym_1), MakeNumber(0), gen7638)

					gen7646 = Call(__e, ShenFunc(symshen_4pair), gen7637, gen7639)

				} else {
					gen7646 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen7646 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7646
			gen7745 := Call(__e, ShenFunc(symfail))

			gen7746 := Call(__e, ShenFunc(sym_a), YaccParse, gen7745)

			if True == gen7746 {
				gen7647 := Call(__e, ShenFunc(symshen_4_5plus_6), V1603)

				Parse__shen_4_5plus_6 := gen7647
				gen7654 := Call(__e, ShenFunc(symfail))

				gen7655 := Call(__e, ShenFunc(sym_a), gen7654, Parse__shen_4_5plus_6)

				gen7656 := Call(__e, ShenFunc(symnot), gen7655)

				var gen7657 Obj
				if True == gen7656 {
					gen7648 := Call(__e, ShenFunc(symshen_4_5number_6), Parse__shen_4_5plus_6)

					Parse__shen_4_5number_6 := gen7648
					gen7651 := Call(__e, ShenFunc(symfail))

					gen7652 := Call(__e, ShenFunc(sym_a), gen7651, Parse__shen_4_5number_6)

					gen7653 := Call(__e, ShenFunc(symnot), gen7652)

					if True == gen7653 {
						gen7649 := Call(__e, ShenFunc(symhd), Parse__shen_4_5number_6)

						gen7650 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5number_6)

						gen7657 = Call(__e, ShenFunc(symshen_4pair), gen7649, gen7650)

					} else {
						gen7657 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen7657 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen7657
				gen7743 := Call(__e, ShenFunc(symfail))

				gen7744 := Call(__e, ShenFunc(sym_a), YaccParse, gen7743)

				if True == gen7744 {
					gen7658 := Call(__e, ShenFunc(symshen_4_5predigits_6), V1603)

					Parse__shen_4_5predigits_6 := gen7658
					gen7685 := Call(__e, ShenFunc(symfail))

					gen7686 := Call(__e, ShenFunc(sym_a), gen7685, Parse__shen_4_5predigits_6)

					gen7687 := Call(__e, ShenFunc(symnot), gen7686)

					var gen7688 Obj
					if True == gen7687 {
						gen7659 := Call(__e, ShenFunc(symshen_4_5stop_6), Parse__shen_4_5predigits_6)

						Parse__shen_4_5stop_6 := gen7659
						gen7682 := Call(__e, ShenFunc(symfail))

						gen7683 := Call(__e, ShenFunc(sym_a), gen7682, Parse__shen_4_5stop_6)

						gen7684 := Call(__e, ShenFunc(symnot), gen7683)

						if True == gen7684 {
							gen7660 := Call(__e, ShenFunc(symshen_4_5postdigits_6), Parse__shen_4_5stop_6)

							Parse__shen_4_5postdigits_6 := gen7660
							gen7679 := Call(__e, ShenFunc(symfail))

							gen7680 := Call(__e, ShenFunc(sym_a), gen7679, Parse__shen_4_5postdigits_6)

							gen7681 := Call(__e, ShenFunc(symnot), gen7680)

							if True == gen7681 {
								gen7661 := Call(__e, ShenFunc(symshen_4_5E_6), Parse__shen_4_5postdigits_6)

								Parse__shen_4_5E_6 := gen7661
								gen7676 := Call(__e, ShenFunc(symfail))

								gen7677 := Call(__e, ShenFunc(sym_a), gen7676, Parse__shen_4_5E_6)

								gen7678 := Call(__e, ShenFunc(symnot), gen7677)

								if True == gen7678 {
									gen7662 := Call(__e, ShenFunc(symshen_4_5log10_6), Parse__shen_4_5E_6)

									Parse__shen_4_5log10_6 := gen7662
									gen7673 := Call(__e, ShenFunc(symfail))

									gen7674 := Call(__e, ShenFunc(sym_a), gen7673, Parse__shen_4_5log10_6)

									gen7675 := Call(__e, ShenFunc(symnot), gen7674)

									if True == gen7675 {
										gen7663 := Call(__e, ShenFunc(symhd), Parse__shen_4_5log10_6)

										gen7664 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5log10_6)

										gen7665 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), gen7664)

										gen7666 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5predigits_6)

										gen7667 := Call(__e, ShenFunc(symreverse), gen7666)

										gen7668 := Call(__e, ShenFunc(symshen_4pre), gen7667, MakeNumber(0))

										gen7669 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5postdigits_6)

										gen7670 := Call(__e, ShenFunc(symshen_4post), gen7669, MakeNumber(1))

										gen7671 := Call(__e, ShenFunc(sym_7), gen7668, gen7670)

										gen7672 := Call(__e, ShenFunc(sym_d), gen7665, gen7671)

										gen7688 = Call(__e, ShenFunc(symshen_4pair), gen7663, gen7672)

									} else {
										gen7688 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen7688 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen7688 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen7688 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen7688 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen7688
					gen7741 := Call(__e, ShenFunc(symfail))

					gen7742 := Call(__e, ShenFunc(sym_a), YaccParse, gen7741)

					if True == gen7742 {
						gen7689 := Call(__e, ShenFunc(symshen_4_5digits_6), V1603)

						Parse__shen_4_5digits_6 := gen7689
						gen7705 := Call(__e, ShenFunc(symfail))

						gen7706 := Call(__e, ShenFunc(sym_a), gen7705, Parse__shen_4_5digits_6)

						gen7707 := Call(__e, ShenFunc(symnot), gen7706)

						var gen7708 Obj
						if True == gen7707 {
							gen7690 := Call(__e, ShenFunc(symshen_4_5E_6), Parse__shen_4_5digits_6)

							Parse__shen_4_5E_6 := gen7690
							gen7702 := Call(__e, ShenFunc(symfail))

							gen7703 := Call(__e, ShenFunc(sym_a), gen7702, Parse__shen_4_5E_6)

							gen7704 := Call(__e, ShenFunc(symnot), gen7703)

							if True == gen7704 {
								gen7691 := Call(__e, ShenFunc(symshen_4_5log10_6), Parse__shen_4_5E_6)

								Parse__shen_4_5log10_6 := gen7691
								gen7699 := Call(__e, ShenFunc(symfail))

								gen7700 := Call(__e, ShenFunc(sym_a), gen7699, Parse__shen_4_5log10_6)

								gen7701 := Call(__e, ShenFunc(symnot), gen7700)

								if True == gen7701 {
									gen7692 := Call(__e, ShenFunc(symhd), Parse__shen_4_5log10_6)

									gen7693 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5log10_6)

									gen7694 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), gen7693)

									gen7695 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

									gen7696 := Call(__e, ShenFunc(symreverse), gen7695)

									gen7697 := Call(__e, ShenFunc(symshen_4pre), gen7696, MakeNumber(0))

									gen7698 := Call(__e, ShenFunc(sym_d), gen7694, gen7697)

									gen7708 = Call(__e, ShenFunc(symshen_4pair), gen7692, gen7698)

								} else {
									gen7708 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen7708 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen7708 = Call(__e, ShenFunc(symfail))

						}
						YaccParse := gen7708
						gen7739 := Call(__e, ShenFunc(symfail))

						gen7740 := Call(__e, ShenFunc(sym_a), YaccParse, gen7739)

						if True == gen7740 {
							gen7709 := Call(__e, ShenFunc(symshen_4_5predigits_6), V1603)

							Parse__shen_4_5predigits_6 := gen7709
							gen7725 := Call(__e, ShenFunc(symfail))

							gen7726 := Call(__e, ShenFunc(sym_a), gen7725, Parse__shen_4_5predigits_6)

							gen7727 := Call(__e, ShenFunc(symnot), gen7726)

							var gen7728 Obj
							if True == gen7727 {
								gen7710 := Call(__e, ShenFunc(symshen_4_5stop_6), Parse__shen_4_5predigits_6)

								Parse__shen_4_5stop_6 := gen7710
								gen7722 := Call(__e, ShenFunc(symfail))

								gen7723 := Call(__e, ShenFunc(sym_a), gen7722, Parse__shen_4_5stop_6)

								gen7724 := Call(__e, ShenFunc(symnot), gen7723)

								if True == gen7724 {
									gen7711 := Call(__e, ShenFunc(symshen_4_5postdigits_6), Parse__shen_4_5stop_6)

									Parse__shen_4_5postdigits_6 := gen7711
									gen7719 := Call(__e, ShenFunc(symfail))

									gen7720 := Call(__e, ShenFunc(sym_a), gen7719, Parse__shen_4_5postdigits_6)

									gen7721 := Call(__e, ShenFunc(symnot), gen7720)

									if True == gen7721 {
										gen7712 := Call(__e, ShenFunc(symhd), Parse__shen_4_5postdigits_6)

										gen7713 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5predigits_6)

										gen7714 := Call(__e, ShenFunc(symreverse), gen7713)

										gen7715 := Call(__e, ShenFunc(symshen_4pre), gen7714, MakeNumber(0))

										gen7716 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5postdigits_6)

										gen7717 := Call(__e, ShenFunc(symshen_4post), gen7716, MakeNumber(1))

										gen7718 := Call(__e, ShenFunc(sym_7), gen7715, gen7717)

										gen7728 = Call(__e, ShenFunc(symshen_4pair), gen7712, gen7718)

									} else {
										gen7728 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen7728 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen7728 = Call(__e, ShenFunc(symfail))

							}
							YaccParse := gen7728
							gen7737 := Call(__e, ShenFunc(symfail))

							gen7738 := Call(__e, ShenFunc(sym_a), YaccParse, gen7737)

							if True == gen7738 {
								gen7729 := Call(__e, ShenFunc(symshen_4_5digits_6), V1603)

								Parse__shen_4_5digits_6 := gen7729
								gen7734 := Call(__e, ShenFunc(symfail))

								gen7735 := Call(__e, ShenFunc(sym_a), gen7734, Parse__shen_4_5digits_6)

								gen7736 := Call(__e, ShenFunc(symnot), gen7735)

								if True == gen7736 {
									gen7730 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

									gen7731 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

									gen7732 := Call(__e, ShenFunc(symreverse), gen7731)

									gen7733 := Call(__e, ShenFunc(symshen_4pre), gen7732, MakeNumber(0))

									__e.TailApply(ShenFunc(symshen_4pair), gen7730, gen7733)

									return

								} else {
									__e.TailApply(ShenFunc(symfail))

									return
								}

							} else {
								__e.Return(YaccParse)
								return
							}

						} else {
							__e.Return(YaccParse)
							return
						}

					} else {
						__e.Return(YaccParse)
						return
					}

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<number>"), gen7747)

		gen7757 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1606 := __args[0]
			_ = V1606
			gen7754 := Call(__e, ShenFunc(symhd), V1606)

			gen7755 := Call(__e, ShenFunc(symcons_2), gen7754)

			var gen7756 Obj
			if True == gen7755 {
				gen7752 := Call(__e, ShenFunc(symshen_4hdhd), V1606)

				gen7753 := Call(__e, ShenFunc(sym_a), MakeNumber(101), gen7752)

				if True == gen7753 {
					gen7756 = True
				} else {
					gen7756 = False
				}

			} else {
				gen7756 = False
			}
			if True == gen7756 {
				gen7748 := Call(__e, ShenFunc(symshen_4tlhd), V1606)

				gen7749 := Call(__e, ShenFunc(symshen_4hdtl), V1606)

				gen7750 := Call(__e, ShenFunc(symshen_4pair), gen7748, gen7749)

				NewStream1604 := gen7750
				gen7751 := Call(__e, ShenFunc(symhd), NewStream1604)

				__e.TailApply(ShenFunc(symshen_4pair), gen7751, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<E>"), gen7757)

		gen7782 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1608 := __args[0]
			_ = V1608
			gen7758 := Call(__e, ShenFunc(symshen_4_5minus_6), V1608)

			Parse__shen_4_5minus_6 := gen7758
			gen7768 := Call(__e, ShenFunc(symfail))

			gen7769 := Call(__e, ShenFunc(sym_a), gen7768, Parse__shen_4_5minus_6)

			gen7770 := Call(__e, ShenFunc(symnot), gen7769)

			var gen7771 Obj
			if True == gen7770 {
				gen7759 := Call(__e, ShenFunc(symshen_4_5digits_6), Parse__shen_4_5minus_6)

				Parse__shen_4_5digits_6 := gen7759
				gen7765 := Call(__e, ShenFunc(symfail))

				gen7766 := Call(__e, ShenFunc(sym_a), gen7765, Parse__shen_4_5digits_6)

				gen7767 := Call(__e, ShenFunc(symnot), gen7766)

				if True == gen7767 {
					gen7760 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

					gen7761 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

					gen7762 := Call(__e, ShenFunc(symreverse), gen7761)

					gen7763 := Call(__e, ShenFunc(symshen_4pre), gen7762, MakeNumber(0))

					gen7764 := Call(__e, ShenFunc(sym_1), MakeNumber(0), gen7763)

					gen7771 = Call(__e, ShenFunc(symshen_4pair), gen7760, gen7764)

				} else {
					gen7771 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen7771 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7771
			gen7780 := Call(__e, ShenFunc(symfail))

			gen7781 := Call(__e, ShenFunc(sym_a), YaccParse, gen7780)

			if True == gen7781 {
				gen7772 := Call(__e, ShenFunc(symshen_4_5digits_6), V1608)

				Parse__shen_4_5digits_6 := gen7772
				gen7777 := Call(__e, ShenFunc(symfail))

				gen7778 := Call(__e, ShenFunc(sym_a), gen7777, Parse__shen_4_5digits_6)

				gen7779 := Call(__e, ShenFunc(symnot), gen7778)

				if True == gen7779 {
					gen7773 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

					gen7774 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

					gen7775 := Call(__e, ShenFunc(symreverse), gen7774)

					gen7776 := Call(__e, ShenFunc(symshen_4pre), gen7775, MakeNumber(0))

					__e.TailApply(ShenFunc(symshen_4pair), gen7773, gen7776)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<log10>"), gen7782)

		gen7791 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1610 := __args[0]
			_ = V1610
			gen7789 := Call(__e, ShenFunc(symhd), V1610)

			gen7790 := Call(__e, ShenFunc(symcons_2), gen7789)

			if True == gen7790 {
				gen7783 := Call(__e, ShenFunc(symshen_4hdhd), V1610)

				Parse__Char := gen7783
				gen7788 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(43))

				if True == gen7788 {
					gen7784 := Call(__e, ShenFunc(symshen_4tlhd), V1610)

					gen7785 := Call(__e, ShenFunc(symshen_4hdtl), V1610)

					gen7786 := Call(__e, ShenFunc(symshen_4pair), gen7784, gen7785)

					gen7787 := Call(__e, ShenFunc(symhd), gen7786)

					__e.TailApply(ShenFunc(symshen_4pair), gen7787, Parse__Char)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<plus>"), gen7791)

		gen7800 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1612 := __args[0]
			_ = V1612
			gen7798 := Call(__e, ShenFunc(symhd), V1612)

			gen7799 := Call(__e, ShenFunc(symcons_2), gen7798)

			if True == gen7799 {
				gen7792 := Call(__e, ShenFunc(symshen_4hdhd), V1612)

				Parse__Char := gen7792
				gen7797 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(46))

				if True == gen7797 {
					gen7793 := Call(__e, ShenFunc(symshen_4tlhd), V1612)

					gen7794 := Call(__e, ShenFunc(symshen_4hdtl), V1612)

					gen7795 := Call(__e, ShenFunc(symshen_4pair), gen7793, gen7794)

					gen7796 := Call(__e, ShenFunc(symhd), gen7795)

					__e.TailApply(ShenFunc(symshen_4pair), gen7796, Parse__Char)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<stop>"), gen7800)

		gen7815 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1614 := __args[0]
			_ = V1614
			gen7801 := Call(__e, ShenFunc(symshen_4_5digits_6), V1614)

			Parse__shen_4_5digits_6 := gen7801
			gen7804 := Call(__e, ShenFunc(symfail))

			gen7805 := Call(__e, ShenFunc(sym_a), gen7804, Parse__shen_4_5digits_6)

			gen7806 := Call(__e, ShenFunc(symnot), gen7805)

			var gen7807 Obj
			if True == gen7806 {
				gen7802 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

				gen7803 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

				gen7807 = Call(__e, ShenFunc(symshen_4pair), gen7802, gen7803)

			} else {
				gen7807 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7807
			gen7813 := Call(__e, ShenFunc(symfail))

			gen7814 := Call(__e, ShenFunc(sym_a), YaccParse, gen7813)

			if True == gen7814 {
				gen7808 := Call(__e, ShenFunc(sym_5e_6), V1614)

				Parse___5e_6 := gen7808
				gen7810 := Call(__e, ShenFunc(symfail))

				gen7811 := Call(__e, ShenFunc(sym_a), gen7810, Parse___5e_6)

				gen7812 := Call(__e, ShenFunc(symnot), gen7811)

				if True == gen7812 {
					gen7809 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen7809, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<predigits>"), gen7815)

		gen7822 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1616 := __args[0]
			_ = V1616
			gen7816 := Call(__e, ShenFunc(symshen_4_5digits_6), V1616)

			Parse__shen_4_5digits_6 := gen7816
			gen7819 := Call(__e, ShenFunc(symfail))

			gen7820 := Call(__e, ShenFunc(sym_a), gen7819, Parse__shen_4_5digits_6)

			gen7821 := Call(__e, ShenFunc(symnot), gen7820)

			if True == gen7821 {
				gen7817 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

				gen7818 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen7817, gen7818)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<postdigits>"), gen7822)

		gen7845 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1618 := __args[0]
			_ = V1618
			gen7823 := Call(__e, ShenFunc(symshen_4_5digit_6), V1618)

			Parse__shen_4_5digit_6 := gen7823
			gen7832 := Call(__e, ShenFunc(symfail))

			gen7833 := Call(__e, ShenFunc(sym_a), gen7832, Parse__shen_4_5digit_6)

			gen7834 := Call(__e, ShenFunc(symnot), gen7833)

			var gen7835 Obj
			if True == gen7834 {
				gen7824 := Call(__e, ShenFunc(symshen_4_5digits_6), Parse__shen_4_5digit_6)

				Parse__shen_4_5digits_6 := gen7824
				gen7829 := Call(__e, ShenFunc(symfail))

				gen7830 := Call(__e, ShenFunc(sym_a), gen7829, Parse__shen_4_5digits_6)

				gen7831 := Call(__e, ShenFunc(symnot), gen7830)

				if True == gen7831 {
					gen7825 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

					gen7826 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digit_6)

					gen7827 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

					gen7828 := Call(__e, ShenFunc(symcons), gen7826, gen7827)

					gen7835 = Call(__e, ShenFunc(symshen_4pair), gen7825, gen7828)

				} else {
					gen7835 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen7835 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7835
			gen7843 := Call(__e, ShenFunc(symfail))

			gen7844 := Call(__e, ShenFunc(sym_a), YaccParse, gen7843)

			if True == gen7844 {
				gen7836 := Call(__e, ShenFunc(symshen_4_5digit_6), V1618)

				Parse__shen_4_5digit_6 := gen7836
				gen7840 := Call(__e, ShenFunc(symfail))

				gen7841 := Call(__e, ShenFunc(sym_a), gen7840, Parse__shen_4_5digit_6)

				gen7842 := Call(__e, ShenFunc(symnot), gen7841)

				if True == gen7842 {
					gen7837 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digit_6)

					gen7838 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digit_6)

					gen7839 := Call(__e, ShenFunc(symcons), gen7838, Nil)

					__e.TailApply(ShenFunc(symshen_4pair), gen7837, gen7839)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<digits>"), gen7845)

		gen7855 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1620 := __args[0]
			_ = V1620
			gen7853 := Call(__e, ShenFunc(symhd), V1620)

			gen7854 := Call(__e, ShenFunc(symcons_2), gen7853)

			if True == gen7854 {
				gen7846 := Call(__e, ShenFunc(symshen_4hdhd), V1620)

				Parse__X := gen7846
				gen7852 := Call(__e, ShenFunc(symshen_4numbyte_2), Parse__X)

				if True == gen7852 {
					gen7847 := Call(__e, ShenFunc(symshen_4tlhd), V1620)

					gen7848 := Call(__e, ShenFunc(symshen_4hdtl), V1620)

					gen7849 := Call(__e, ShenFunc(symshen_4pair), gen7847, gen7848)

					gen7850 := Call(__e, ShenFunc(symhd), gen7849)

					gen7851 := Call(__e, ShenFunc(symshen_4byte_1_6digit), Parse__X)

					__e.TailApply(ShenFunc(symshen_4pair), gen7850, gen7851)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<digit>"), gen7855)

		gen7866 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1622 := __args[0]
			_ = V1622
			gen7865 := Call(__e, ShenFunc(sym_a), MakeNumber(48), V1622)

			if True == gen7865 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen7864 := Call(__e, ShenFunc(sym_a), MakeNumber(49), V1622)

				if True == gen7864 {
					__e.Return(MakeNumber(1))
					return
				} else {
					gen7863 := Call(__e, ShenFunc(sym_a), MakeNumber(50), V1622)

					if True == gen7863 {
						__e.Return(MakeNumber(2))
						return
					} else {
						gen7862 := Call(__e, ShenFunc(sym_a), MakeNumber(51), V1622)

						if True == gen7862 {
							__e.Return(MakeNumber(3))
							return
						} else {
							gen7861 := Call(__e, ShenFunc(sym_a), MakeNumber(52), V1622)

							if True == gen7861 {
								__e.Return(MakeNumber(4))
								return
							} else {
								gen7860 := Call(__e, ShenFunc(sym_a), MakeNumber(53), V1622)

								if True == gen7860 {
									__e.Return(MakeNumber(5))
									return
								} else {
									gen7859 := Call(__e, ShenFunc(sym_a), MakeNumber(54), V1622)

									if True == gen7859 {
										__e.Return(MakeNumber(6))
										return
									} else {
										gen7858 := Call(__e, ShenFunc(sym_a), MakeNumber(55), V1622)

										if True == gen7858 {
											__e.Return(MakeNumber(7))
											return
										} else {
											gen7857 := Call(__e, ShenFunc(sym_a), MakeNumber(56), V1622)

											if True == gen7857 {
												__e.Return(MakeNumber(8))
												return
											} else {
												gen7856 := Call(__e, ShenFunc(sym_a), MakeNumber(57), V1622)

												if True == gen7856 {
													__e.Return(MakeNumber(9))
													return
												} else {
													__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.byte->digit"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.byte->digit"), gen7866)

		gen7875 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1627 := __args[0]
			_ = V1627
			V1628 := __args[1]
			_ = V1628
			gen7874 := Call(__e, ShenFunc(sym_a), Nil, V1627)

			if True == gen7874 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen7873 := Call(__e, ShenFunc(symcons_2), V1627)

				if True == gen7873 {
					gen7867 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), V1628)

					gen7868 := Call(__e, ShenFunc(symhd), V1627)

					gen7869 := Call(__e, ShenFunc(sym_d), gen7867, gen7868)

					gen7870 := Call(__e, ShenFunc(symtl), V1627)

					gen7871 := Call(__e, ShenFunc(sym_7), V1628, MakeNumber(1))

					gen7872 := Call(__e, ShenFunc(symshen_4pre), gen7870, gen7871)

					__e.TailApply(ShenFunc(sym_7), gen7869, gen7872)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.pre"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pre"), gen7875)

		gen7885 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1633 := __args[0]
			_ = V1633
			V1634 := __args[1]
			_ = V1634
			gen7884 := Call(__e, ShenFunc(sym_a), Nil, V1633)

			if True == gen7884 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen7883 := Call(__e, ShenFunc(symcons_2), V1633)

				if True == gen7883 {
					gen7876 := Call(__e, ShenFunc(sym_1), MakeNumber(0), V1634)

					gen7877 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), gen7876)

					gen7878 := Call(__e, ShenFunc(symhd), V1633)

					gen7879 := Call(__e, ShenFunc(sym_d), gen7877, gen7878)

					gen7880 := Call(__e, ShenFunc(symtl), V1633)

					gen7881 := Call(__e, ShenFunc(sym_7), V1634, MakeNumber(1))

					gen7882 := Call(__e, ShenFunc(symshen_4post), gen7880, gen7881)

					__e.TailApply(ShenFunc(sym_7), gen7879, gen7882)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.post"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.post"), gen7885)

		gen7893 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1639 := __args[0]
			_ = V1639
			V1640 := __args[1]
			_ = V1640
			gen7892 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V1640)

			if True == gen7892 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen7891 := Call(__e, ShenFunc(sym_6), V1640, MakeNumber(0))

				if True == gen7891 {
					gen7889 := Call(__e, ShenFunc(sym_1), V1640, MakeNumber(1))

					gen7890 := Call(__e, ShenFunc(symshen_4expt), V1639, gen7889)

					__e.TailApply(ShenFunc(sym_d), V1639, gen7890)

					return

				} else {
					gen7886 := Call(__e, ShenFunc(sym_7), V1640, MakeNumber(1))

					gen7887 := Call(__e, ShenFunc(symshen_4expt), V1639, gen7886)

					gen7888 := Call(__e, ShenFunc(sym_c), gen7887, V1639)

					__e.TailApply(ShenFunc(sym_d), MakeNumber(1), gen7888)

					return

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.expt"), gen7893)

		gen7900 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1642 := __args[0]
			_ = V1642
			gen7894 := Call(__e, ShenFunc(symshen_4_5st__input_6), V1642)

			Parse__shen_4_5st__input_6 := gen7894
			gen7897 := Call(__e, ShenFunc(symfail))

			gen7898 := Call(__e, ShenFunc(sym_a), gen7897, Parse__shen_4_5st__input_6)

			gen7899 := Call(__e, ShenFunc(symnot), gen7898)

			if True == gen7899 {
				gen7895 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

				gen7896 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen7895, gen7896)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<st_input1>"), gen7900)

		gen7907 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1644 := __args[0]
			_ = V1644
			gen7901 := Call(__e, ShenFunc(symshen_4_5st__input_6), V1644)

			Parse__shen_4_5st__input_6 := gen7901
			gen7904 := Call(__e, ShenFunc(symfail))

			gen7905 := Call(__e, ShenFunc(sym_a), gen7904, Parse__shen_4_5st__input_6)

			gen7906 := Call(__e, ShenFunc(symnot), gen7905)

			if True == gen7906 {
				gen7902 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

				gen7903 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen7902, gen7903)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<st_input2>"), gen7907)

		gen7921 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1646 := __args[0]
			_ = V1646
			gen7908 := Call(__e, ShenFunc(symshen_4_5singleline_6), V1646)

			Parse__shen_4_5singleline_6 := gen7908
			gen7910 := Call(__e, ShenFunc(symfail))

			gen7911 := Call(__e, ShenFunc(sym_a), gen7910, Parse__shen_4_5singleline_6)

			gen7912 := Call(__e, ShenFunc(symnot), gen7911)

			var gen7913 Obj
			if True == gen7912 {
				gen7909 := Call(__e, ShenFunc(symhd), Parse__shen_4_5singleline_6)

				gen7913 = Call(__e, ShenFunc(symshen_4pair), gen7909, MakeSymbol("shen.skip"))

			} else {
				gen7913 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7913
			gen7919 := Call(__e, ShenFunc(symfail))

			gen7920 := Call(__e, ShenFunc(sym_a), YaccParse, gen7919)

			if True == gen7920 {
				gen7914 := Call(__e, ShenFunc(symshen_4_5multiline_6), V1646)

				Parse__shen_4_5multiline_6 := gen7914
				gen7916 := Call(__e, ShenFunc(symfail))

				gen7917 := Call(__e, ShenFunc(sym_a), gen7916, Parse__shen_4_5multiline_6)

				gen7918 := Call(__e, ShenFunc(symnot), gen7917)

				if True == gen7918 {
					gen7915 := Call(__e, ShenFunc(symhd), Parse__shen_4_5multiline_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen7915, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<comment>"), gen7921)

		gen7939 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1648 := __args[0]
			_ = V1648
			gen7922 := Call(__e, ShenFunc(symshen_4_5backslash_6), V1648)

			Parse__shen_4_5backslash_6 := gen7922
			gen7936 := Call(__e, ShenFunc(symfail))

			gen7937 := Call(__e, ShenFunc(sym_a), gen7936, Parse__shen_4_5backslash_6)

			gen7938 := Call(__e, ShenFunc(symnot), gen7937)

			if True == gen7938 {
				gen7923 := Call(__e, ShenFunc(symshen_4_5backslash_6), Parse__shen_4_5backslash_6)

				Parse__shen_4_5backslash_6 := gen7923
				gen7933 := Call(__e, ShenFunc(symfail))

				gen7934 := Call(__e, ShenFunc(sym_a), gen7933, Parse__shen_4_5backslash_6)

				gen7935 := Call(__e, ShenFunc(symnot), gen7934)

				if True == gen7935 {
					gen7924 := Call(__e, ShenFunc(symshen_4_5anysingle_6), Parse__shen_4_5backslash_6)

					Parse__shen_4_5anysingle_6 := gen7924
					gen7930 := Call(__e, ShenFunc(symfail))

					gen7931 := Call(__e, ShenFunc(sym_a), gen7930, Parse__shen_4_5anysingle_6)

					gen7932 := Call(__e, ShenFunc(symnot), gen7931)

					if True == gen7932 {
						gen7925 := Call(__e, ShenFunc(symshen_4_5return_6), Parse__shen_4_5anysingle_6)

						Parse__shen_4_5return_6 := gen7925
						gen7927 := Call(__e, ShenFunc(symfail))

						gen7928 := Call(__e, ShenFunc(sym_a), gen7927, Parse__shen_4_5return_6)

						gen7929 := Call(__e, ShenFunc(symnot), gen7928)

						if True == gen7929 {
							gen7926 := Call(__e, ShenFunc(symhd), Parse__shen_4_5return_6)

							__e.TailApply(ShenFunc(symshen_4pair), gen7926, MakeSymbol("shen.skip"))

							return

						} else {
							__e.TailApply(ShenFunc(symfail))

							return
						}

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<singleline>"), gen7939)

		gen7949 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1651 := __args[0]
			_ = V1651
			gen7946 := Call(__e, ShenFunc(symhd), V1651)

			gen7947 := Call(__e, ShenFunc(symcons_2), gen7946)

			var gen7948 Obj
			if True == gen7947 {
				gen7944 := Call(__e, ShenFunc(symshen_4hdhd), V1651)

				gen7945 := Call(__e, ShenFunc(sym_a), MakeNumber(92), gen7944)

				if True == gen7945 {
					gen7948 = True
				} else {
					gen7948 = False
				}

			} else {
				gen7948 = False
			}
			if True == gen7948 {
				gen7940 := Call(__e, ShenFunc(symshen_4tlhd), V1651)

				gen7941 := Call(__e, ShenFunc(symshen_4hdtl), V1651)

				gen7942 := Call(__e, ShenFunc(symshen_4pair), gen7940, gen7941)

				NewStream1649 := gen7942
				gen7943 := Call(__e, ShenFunc(symhd), NewStream1649)

				__e.TailApply(ShenFunc(symshen_4pair), gen7943, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<backslash>"), gen7949)

		gen7967 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1653 := __args[0]
			_ = V1653
			gen7950 := Call(__e, ShenFunc(symshen_4_5non_1return_6), V1653)

			Parse__shen_4_5non_1return_6 := gen7950
			gen7956 := Call(__e, ShenFunc(symfail))

			gen7957 := Call(__e, ShenFunc(sym_a), gen7956, Parse__shen_4_5non_1return_6)

			gen7958 := Call(__e, ShenFunc(symnot), gen7957)

			var gen7959 Obj
			if True == gen7958 {
				gen7951 := Call(__e, ShenFunc(symshen_4_5anysingle_6), Parse__shen_4_5non_1return_6)

				Parse__shen_4_5anysingle_6 := gen7951
				gen7953 := Call(__e, ShenFunc(symfail))

				gen7954 := Call(__e, ShenFunc(sym_a), gen7953, Parse__shen_4_5anysingle_6)

				gen7955 := Call(__e, ShenFunc(symnot), gen7954)

				if True == gen7955 {
					gen7952 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anysingle_6)

					gen7959 = Call(__e, ShenFunc(symshen_4pair), gen7952, MakeSymbol("shen.skip"))

				} else {
					gen7959 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen7959 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen7959
			gen7965 := Call(__e, ShenFunc(symfail))

			gen7966 := Call(__e, ShenFunc(sym_a), YaccParse, gen7965)

			if True == gen7966 {
				gen7960 := Call(__e, ShenFunc(sym_5e_6), V1653)

				Parse___5e_6 := gen7960
				gen7962 := Call(__e, ShenFunc(symfail))

				gen7963 := Call(__e, ShenFunc(sym_a), gen7962, Parse___5e_6)

				gen7964 := Call(__e, ShenFunc(symnot), gen7963)

				if True == gen7964 {
					gen7961 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen7961, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<anysingle>"), gen7967)

		gen7979 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1655 := __args[0]
			_ = V1655
			gen7977 := Call(__e, ShenFunc(symhd), V1655)

			gen7978 := Call(__e, ShenFunc(symcons_2), gen7977)

			if True == gen7978 {
				gen7968 := Call(__e, ShenFunc(symshen_4hdhd), V1655)

				Parse__X := gen7968
				gen7973 := Call(__e, ShenFunc(symcons), MakeNumber(13), Nil)

				gen7974 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen7973)

				gen7975 := Call(__e, ShenFunc(symelement_2), Parse__X, gen7974)

				gen7976 := Call(__e, ShenFunc(symnot), gen7975)

				if True == gen7976 {
					gen7969 := Call(__e, ShenFunc(symshen_4tlhd), V1655)

					gen7970 := Call(__e, ShenFunc(symshen_4hdtl), V1655)

					gen7971 := Call(__e, ShenFunc(symshen_4pair), gen7969, gen7970)

					gen7972 := Call(__e, ShenFunc(symhd), gen7971)

					__e.TailApply(ShenFunc(symshen_4pair), gen7972, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<non-return>"), gen7979)

		gen7990 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1657 := __args[0]
			_ = V1657
			gen7988 := Call(__e, ShenFunc(symhd), V1657)

			gen7989 := Call(__e, ShenFunc(symcons_2), gen7988)

			if True == gen7989 {
				gen7980 := Call(__e, ShenFunc(symshen_4hdhd), V1657)

				Parse__X := gen7980
				gen7985 := Call(__e, ShenFunc(symcons), MakeNumber(13), Nil)

				gen7986 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen7985)

				gen7987 := Call(__e, ShenFunc(symelement_2), Parse__X, gen7986)

				if True == gen7987 {
					gen7981 := Call(__e, ShenFunc(symshen_4tlhd), V1657)

					gen7982 := Call(__e, ShenFunc(symshen_4hdtl), V1657)

					gen7983 := Call(__e, ShenFunc(symshen_4pair), gen7981, gen7982)

					gen7984 := Call(__e, ShenFunc(symhd), gen7983)

					__e.TailApply(ShenFunc(symshen_4pair), gen7984, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<return>"), gen7990)

		gen8004 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1659 := __args[0]
			_ = V1659
			gen7991 := Call(__e, ShenFunc(symshen_4_5backslash_6), V1659)

			Parse__shen_4_5backslash_6 := gen7991
			gen8001 := Call(__e, ShenFunc(symfail))

			gen8002 := Call(__e, ShenFunc(sym_a), gen8001, Parse__shen_4_5backslash_6)

			gen8003 := Call(__e, ShenFunc(symnot), gen8002)

			if True == gen8003 {
				gen7992 := Call(__e, ShenFunc(symshen_4_5times_6), Parse__shen_4_5backslash_6)

				Parse__shen_4_5times_6 := gen7992
				gen7998 := Call(__e, ShenFunc(symfail))

				gen7999 := Call(__e, ShenFunc(sym_a), gen7998, Parse__shen_4_5times_6)

				gen8000 := Call(__e, ShenFunc(symnot), gen7999)

				if True == gen8000 {
					gen7993 := Call(__e, ShenFunc(symshen_4_5anymulti_6), Parse__shen_4_5times_6)

					Parse__shen_4_5anymulti_6 := gen7993
					gen7995 := Call(__e, ShenFunc(symfail))

					gen7996 := Call(__e, ShenFunc(sym_a), gen7995, Parse__shen_4_5anymulti_6)

					gen7997 := Call(__e, ShenFunc(symnot), gen7996)

					if True == gen7997 {
						gen7994 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anymulti_6)

						__e.TailApply(ShenFunc(symshen_4pair), gen7994, MakeSymbol("shen.skip"))

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<multiline>"), gen8004)

		gen8014 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1662 := __args[0]
			_ = V1662
			gen8011 := Call(__e, ShenFunc(symhd), V1662)

			gen8012 := Call(__e, ShenFunc(symcons_2), gen8011)

			var gen8013 Obj
			if True == gen8012 {
				gen8009 := Call(__e, ShenFunc(symshen_4hdhd), V1662)

				gen8010 := Call(__e, ShenFunc(sym_a), MakeNumber(42), gen8009)

				if True == gen8010 {
					gen8013 = True
				} else {
					gen8013 = False
				}

			} else {
				gen8013 = False
			}
			if True == gen8013 {
				gen8005 := Call(__e, ShenFunc(symshen_4tlhd), V1662)

				gen8006 := Call(__e, ShenFunc(symshen_4hdtl), V1662)

				gen8007 := Call(__e, ShenFunc(symshen_4pair), gen8005, gen8006)

				NewStream1660 := gen8007
				gen8008 := Call(__e, ShenFunc(symhd), NewStream1660)

				__e.TailApply(ShenFunc(symshen_4pair), gen8008, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<times>"), gen8014)

		gen8050 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1664 := __args[0]
			_ = V1664
			gen8015 := Call(__e, ShenFunc(symshen_4_5comment_6), V1664)

			Parse__shen_4_5comment_6 := gen8015
			gen8021 := Call(__e, ShenFunc(symfail))

			gen8022 := Call(__e, ShenFunc(sym_a), gen8021, Parse__shen_4_5comment_6)

			gen8023 := Call(__e, ShenFunc(symnot), gen8022)

			var gen8024 Obj
			if True == gen8023 {
				gen8016 := Call(__e, ShenFunc(symshen_4_5anymulti_6), Parse__shen_4_5comment_6)

				Parse__shen_4_5anymulti_6 := gen8016
				gen8018 := Call(__e, ShenFunc(symfail))

				gen8019 := Call(__e, ShenFunc(sym_a), gen8018, Parse__shen_4_5anymulti_6)

				gen8020 := Call(__e, ShenFunc(symnot), gen8019)

				if True == gen8020 {
					gen8017 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anymulti_6)

					gen8024 = Call(__e, ShenFunc(symshen_4pair), gen8017, MakeSymbol("shen.skip"))

				} else {
					gen8024 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen8024 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen8024
			gen8048 := Call(__e, ShenFunc(symfail))

			gen8049 := Call(__e, ShenFunc(sym_a), YaccParse, gen8048)

			if True == gen8049 {
				gen8025 := Call(__e, ShenFunc(symshen_4_5times_6), V1664)

				Parse__shen_4_5times_6 := gen8025
				gen8031 := Call(__e, ShenFunc(symfail))

				gen8032 := Call(__e, ShenFunc(sym_a), gen8031, Parse__shen_4_5times_6)

				gen8033 := Call(__e, ShenFunc(symnot), gen8032)

				var gen8034 Obj
				if True == gen8033 {
					gen8026 := Call(__e, ShenFunc(symshen_4_5backslash_6), Parse__shen_4_5times_6)

					Parse__shen_4_5backslash_6 := gen8026
					gen8028 := Call(__e, ShenFunc(symfail))

					gen8029 := Call(__e, ShenFunc(sym_a), gen8028, Parse__shen_4_5backslash_6)

					gen8030 := Call(__e, ShenFunc(symnot), gen8029)

					if True == gen8030 {
						gen8027 := Call(__e, ShenFunc(symhd), Parse__shen_4_5backslash_6)

						gen8034 = Call(__e, ShenFunc(symshen_4pair), gen8027, MakeSymbol("shen.skip"))

					} else {
						gen8034 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen8034 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen8034
				gen8046 := Call(__e, ShenFunc(symfail))

				gen8047 := Call(__e, ShenFunc(sym_a), YaccParse, gen8046)

				if True == gen8047 {
					gen8044 := Call(__e, ShenFunc(symhd), V1664)

					gen8045 := Call(__e, ShenFunc(symcons_2), gen8044)

					if True == gen8045 {
						gen8035 := Call(__e, ShenFunc(symshen_4hdhd), V1664)

						Parse__X := gen8035
						_ = Parse__X
						gen8036 := Call(__e, ShenFunc(symshen_4tlhd), V1664)

						gen8037 := Call(__e, ShenFunc(symshen_4hdtl), V1664)

						gen8038 := Call(__e, ShenFunc(symshen_4pair), gen8036, gen8037)

						gen8039 := Call(__e, ShenFunc(symshen_4_5anymulti_6), gen8038)

						Parse__shen_4_5anymulti_6 := gen8039
						gen8041 := Call(__e, ShenFunc(symfail))

						gen8042 := Call(__e, ShenFunc(sym_a), gen8041, Parse__shen_4_5anymulti_6)

						gen8043 := Call(__e, ShenFunc(symnot), gen8042)

						if True == gen8043 {
							gen8040 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anymulti_6)

							__e.TailApply(ShenFunc(symshen_4pair), gen8040, MakeSymbol("shen.skip"))

							return

						} else {
							__e.TailApply(ShenFunc(symfail))

							return
						}

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<anymulti>"), gen8050)

		gen8068 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1666 := __args[0]
			_ = V1666
			gen8051 := Call(__e, ShenFunc(symshen_4_5whitespace_6), V1666)

			Parse__shen_4_5whitespace_6 := gen8051
			gen8057 := Call(__e, ShenFunc(symfail))

			gen8058 := Call(__e, ShenFunc(sym_a), gen8057, Parse__shen_4_5whitespace_6)

			gen8059 := Call(__e, ShenFunc(symnot), gen8058)

			var gen8060 Obj
			if True == gen8059 {
				gen8052 := Call(__e, ShenFunc(symshen_4_5whitespaces_6), Parse__shen_4_5whitespace_6)

				Parse__shen_4_5whitespaces_6 := gen8052
				gen8054 := Call(__e, ShenFunc(symfail))

				gen8055 := Call(__e, ShenFunc(sym_a), gen8054, Parse__shen_4_5whitespaces_6)

				gen8056 := Call(__e, ShenFunc(symnot), gen8055)

				if True == gen8056 {
					gen8053 := Call(__e, ShenFunc(symhd), Parse__shen_4_5whitespaces_6)

					gen8060 = Call(__e, ShenFunc(symshen_4pair), gen8053, MakeSymbol("shen.skip"))

				} else {
					gen8060 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen8060 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen8060
			gen8066 := Call(__e, ShenFunc(symfail))

			gen8067 := Call(__e, ShenFunc(sym_a), YaccParse, gen8066)

			if True == gen8067 {
				gen8061 := Call(__e, ShenFunc(symshen_4_5whitespace_6), V1666)

				Parse__shen_4_5whitespace_6 := gen8061
				gen8063 := Call(__e, ShenFunc(symfail))

				gen8064 := Call(__e, ShenFunc(sym_a), gen8063, Parse__shen_4_5whitespace_6)

				gen8065 := Call(__e, ShenFunc(symnot), gen8064)

				if True == gen8065 {
					gen8062 := Call(__e, ShenFunc(symhd), Parse__shen_4_5whitespace_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen8062, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<whitespaces>"), gen8068)

		gen8083 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1668 := __args[0]
			_ = V1668
			gen8081 := Call(__e, ShenFunc(symhd), V1668)

			gen8082 := Call(__e, ShenFunc(symcons_2), gen8081)

			if True == gen8082 {
				gen8069 := Call(__e, ShenFunc(symshen_4hdhd), V1668)

				Parse__X := gen8069
				Parse__Case := Parse__X
				gen8079 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(32))

				var gen8080 Obj
				if True == gen8079 {
					gen8080 = True
				} else {
					gen8077 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(13))

					var gen8078 Obj
					if True == gen8077 {
						gen8078 = True
					} else {
						gen8075 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(10))

						var gen8076 Obj
						if True == gen8075 {
							gen8076 = True
						} else {
							gen8074 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(9))

							if True == gen8074 {
								gen8076 = True
							} else {
								gen8076 = False
							}

						}
						if True == gen8076 {
							gen8078 = True
						} else {
							gen8078 = False
						}

					}
					if True == gen8078 {
						gen8080 = True
					} else {
						gen8080 = False
					}

				}
				if True == gen8080 {
					gen8070 := Call(__e, ShenFunc(symshen_4tlhd), V1668)

					gen8071 := Call(__e, ShenFunc(symshen_4hdtl), V1668)

					gen8072 := Call(__e, ShenFunc(symshen_4pair), gen8070, gen8071)

					gen8073 := Call(__e, ShenFunc(symhd), gen8072)

					__e.TailApply(ShenFunc(symshen_4pair), gen8073, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<whitespace>"), gen8083)

		gen8112 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1670 := __args[0]
			_ = V1670
			gen8111 := Call(__e, ShenFunc(sym_a), Nil, V1670)

			if True == gen8111 {
				__e.Return(Nil)
				return
			} else {
				gen8109 := Call(__e, ShenFunc(symcons_2), V1670)

				var gen8110 Obj
				if True == gen8109 {
					gen8106 := Call(__e, ShenFunc(symtl), V1670)

					gen8107 := Call(__e, ShenFunc(symcons_2), gen8106)

					var gen8108 Obj
					if True == gen8107 {
						gen8102 := Call(__e, ShenFunc(symtl), V1670)

						gen8103 := Call(__e, ShenFunc(symtl), gen8102)

						gen8104 := Call(__e, ShenFunc(symcons_2), gen8103)

						var gen8105 Obj
						if True == gen8104 {
							gen8097 := Call(__e, ShenFunc(symtl), V1670)

							gen8098 := Call(__e, ShenFunc(symtl), gen8097)

							gen8099 := Call(__e, ShenFunc(symtl), gen8098)

							gen8100 := Call(__e, ShenFunc(sym_a), Nil, gen8099)

							var gen8101 Obj
							if True == gen8100 {
								gen8094 := Call(__e, ShenFunc(symtl), V1670)

								gen8095 := Call(__e, ShenFunc(symhd), gen8094)

								gen8096 := Call(__e, ShenFunc(sym_a), gen8095, MakeSymbol("bar!"))

								if True == gen8096 {
									gen8101 = True
								} else {
									gen8101 = False
								}

							} else {
								gen8101 = False
							}
							if True == gen8101 {
								gen8105 = True
							} else {
								gen8105 = False
							}

						} else {
							gen8105 = False
						}
						if True == gen8105 {
							gen8108 = True
						} else {
							gen8108 = False
						}

					} else {
						gen8108 = False
					}
					if True == gen8108 {
						gen8110 = True
					} else {
						gen8110 = False
					}

				} else {
					gen8110 = False
				}
				if True == gen8110 {
					gen8090 := Call(__e, ShenFunc(symhd), V1670)

					gen8091 := Call(__e, ShenFunc(symtl), V1670)

					gen8092 := Call(__e, ShenFunc(symtl), gen8091)

					gen8093 := Call(__e, ShenFunc(symcons), gen8090, gen8092)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen8093)

					return

				} else {
					gen8089 := Call(__e, ShenFunc(symcons_2), V1670)

					if True == gen8089 {
						gen8084 := Call(__e, ShenFunc(symhd), V1670)

						gen8085 := Call(__e, ShenFunc(symtl), V1670)

						gen8086 := Call(__e, ShenFunc(symshen_4cons__form), gen8085)

						gen8087 := Call(__e, ShenFunc(symcons), gen8086, Nil)

						gen8088 := Call(__e, ShenFunc(symcons), gen8084, gen8087)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen8088)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.cons_form"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cons_form"), gen8112)

		gen8177 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1675 := __args[0]
			_ = V1675
			V1676 := __args[1]
			_ = V1676
			gen8175 := Call(__e, ShenFunc(symcons_2), V1675)

			var gen8176 Obj
			if True == gen8175 {
				gen8172 := Call(__e, ShenFunc(symhd), V1675)

				gen8173 := Call(__e, ShenFunc(sym_a), MakeSymbol("$"), gen8172)

				var gen8174 Obj
				if True == gen8173 {
					gen8169 := Call(__e, ShenFunc(symtl), V1675)

					gen8170 := Call(__e, ShenFunc(symcons_2), gen8169)

					var gen8171 Obj
					if True == gen8170 {
						gen8166 := Call(__e, ShenFunc(symtl), V1675)

						gen8167 := Call(__e, ShenFunc(symtl), gen8166)

						gen8168 := Call(__e, ShenFunc(sym_a), Nil, gen8167)

						if True == gen8168 {
							gen8171 = True
						} else {
							gen8171 = False
						}

					} else {
						gen8171 = False
					}
					if True == gen8171 {
						gen8174 = True
					} else {
						gen8174 = False
					}

				} else {
					gen8174 = False
				}
				if True == gen8174 {
					gen8176 = True
				} else {
					gen8176 = False
				}

			} else {
				gen8176 = False
			}
			if True == gen8176 {
				gen8163 := Call(__e, ShenFunc(symtl), V1675)

				gen8164 := Call(__e, ShenFunc(symhd), gen8163)

				gen8165 := Call(__e, ShenFunc(symexplode), gen8164)

				__e.TailApply(ShenFunc(symappend), gen8165, V1676)

				return

			} else {
				gen8161 := Call(__e, ShenFunc(symcons_2), V1675)

				var gen8162 Obj
				if True == gen8161 {
					gen8158 := Call(__e, ShenFunc(symhd), V1675)

					gen8159 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen8158)

					var gen8160 Obj
					if True == gen8159 {
						gen8155 := Call(__e, ShenFunc(symtl), V1675)

						gen8156 := Call(__e, ShenFunc(symcons_2), gen8155)

						var gen8157 Obj
						if True == gen8156 {
							gen8151 := Call(__e, ShenFunc(symtl), V1675)

							gen8152 := Call(__e, ShenFunc(symhd), gen8151)

							gen8153 := Call(__e, ShenFunc(sym_a), MakeSymbol("null"), gen8152)

							var gen8154 Obj
							if True == gen8153 {
								gen8148 := Call(__e, ShenFunc(symtl), V1675)

								gen8149 := Call(__e, ShenFunc(symtl), gen8148)

								gen8150 := Call(__e, ShenFunc(symcons_2), gen8149)

								if True == gen8150 {
									gen8154 = True
								} else {
									gen8154 = False
								}

							} else {
								gen8154 = False
							}
							if True == gen8154 {
								gen8157 = True
							} else {
								gen8157 = False
							}

						} else {
							gen8157 = False
						}
						if True == gen8157 {
							gen8160 = True
						} else {
							gen8160 = False
						}

					} else {
						gen8160 = False
					}
					if True == gen8160 {
						gen8162 = True
					} else {
						gen8162 = False
					}

				} else {
					gen8162 = False
				}
				if True == gen8162 {
					gen8145 := Call(__e, ShenFunc(symtl), V1675)

					gen8146 := Call(__e, ShenFunc(symtl), gen8145)

					gen8147 := Call(__e, ShenFunc(symtl), gen8146)

					__e.TailApply(ShenFunc(symappend), gen8147, V1676)

					return

				} else {
					gen8143 := Call(__e, ShenFunc(symcons_2), V1675)

					var gen8144 Obj
					if True == gen8143 {
						gen8140 := Call(__e, ShenFunc(symhd), V1675)

						gen8141 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen8140)

						var gen8142 Obj
						if True == gen8141 {
							gen8137 := Call(__e, ShenFunc(symtl), V1675)

							gen8138 := Call(__e, ShenFunc(symcons_2), gen8137)

							var gen8139 Obj
							if True == gen8138 {
								gen8134 := Call(__e, ShenFunc(symtl), V1675)

								gen8135 := Call(__e, ShenFunc(symtl), gen8134)

								gen8136 := Call(__e, ShenFunc(symcons_2), gen8135)

								if True == gen8136 {
									gen8139 = True
								} else {
									gen8139 = False
								}

							} else {
								gen8139 = False
							}
							if True == gen8139 {
								gen8142 = True
							} else {
								gen8142 = False
							}

						} else {
							gen8142 = False
						}
						if True == gen8142 {
							gen8144 = True
						} else {
							gen8144 = False
						}

					} else {
						gen8144 = False
					}
					if True == gen8144 {
						gen8113 := Call(__e, ShenFunc(symtl), V1675)

						gen8114 := Call(__e, ShenFunc(symtl), gen8113)

						gen8115 := Call(__e, ShenFunc(symhd), gen8114)

						gen8116 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), gen8115)

						ListofExceptions := gen8116
						gen8117 := Call(__e, ShenFunc(symtl), V1675)

						gen8118 := Call(__e, ShenFunc(symhd), gen8117)

						gen8119 := Call(__e, ShenFunc(symshen_4record_1exceptions), ListofExceptions, gen8118)

						External := gen8119
						_ = External
						gen8120 := Call(__e, ShenFunc(symtl), V1675)

						gen8121 := Call(__e, ShenFunc(symhd), gen8120)

						gen8122 := Call(__e, ShenFunc(symstr), gen8121)

						gen8123 := Call(__e, ShenFunc(symcn), gen8122, MakeString("."))

						gen8124 := Call(__e, ShenFunc(symintern), gen8123)

						PackageNameDot := gen8124
						gen8125 := Call(__e, ShenFunc(symexplode), PackageNameDot)

						ExpPackageNameDot := gen8125
						gen8126 := Call(__e, ShenFunc(symtl), V1675)

						gen8127 := Call(__e, ShenFunc(symtl), gen8126)

						gen8128 := Call(__e, ShenFunc(symtl), gen8127)

						gen8129 := Call(__e, ShenFunc(symshen_4packageh), PackageNameDot, ListofExceptions, gen8128, ExpPackageNameDot)

						Packaged := gen8129
						gen8130 := Call(__e, ShenFunc(symtl), V1675)

						gen8131 := Call(__e, ShenFunc(symhd), gen8130)

						gen8132 := Call(__e, ShenFunc(symshen_4internal_1symbols), ExpPackageNameDot, Packaged)

						gen8133 := Call(__e, ShenFunc(symshen_4record_1internal), gen8131, gen8132)

						Internal := gen8133
						_ = Internal
						__e.TailApply(ShenFunc(symappend), Packaged, V1676)

						return

					} else {
						__e.TailApply(ShenFunc(symcons), V1675, V1676)

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.package-macro"), gen8177)

		gen8184 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1679 := __args[0]
			_ = V1679
			V1680 := __args[1]
			_ = V1680
			gen8178 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen8180 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8179 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1680, MakeSymbol("shen.external-symbols"), gen8179)

				return

			}, 0)
			gen8181 := Try(__e, gen8180).Catch(gen8178)
			CurrExceptions := gen8181
			gen8182 := Call(__e, ShenFunc(symunion), V1679, CurrExceptions)

			AllExceptions := gen8182
			gen8183 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			__e.TailApply(ShenFunc(symput), V1680, MakeSymbol("shen.external-symbols"), AllExceptions, gen8183)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-exceptions"), gen8184)

		gen8191 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1683 := __args[0]
			_ = V1683
			V1684 := __args[1]
			_ = V1684
			gen8185 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen8187 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8186 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1683, MakeSymbol("shen.internal-symbols"), gen8186)

				return

			}, 0)
			gen8188 := Try(__e, gen8187).Catch(gen8185)
			gen8189 := Call(__e, ShenFunc(symunion), V1684, gen8188)

			gen8190 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			__e.TailApply(ShenFunc(symput), V1683, MakeSymbol("shen.internal-symbols"), gen8189, gen8190)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-internal"), gen8191)

		gen8201 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1695 := __args[0]
			_ = V1695
			V1696 := __args[1]
			_ = V1696
			gen8199 := Call(__e, ShenFunc(symsymbol_2), V1696)

			var gen8200 Obj
			if True == gen8199 {
				gen8197 := Call(__e, ShenFunc(symexplode), V1696)

				gen8198 := Call(__e, ShenFunc(symshen_4prefix_2), V1695, gen8197)

				if True == gen8198 {
					gen8200 = True
				} else {
					gen8200 = False
				}

			} else {
				gen8200 = False
			}
			if True == gen8200 {
				__e.TailApply(ShenFunc(symcons), V1696, Nil)

				return
			} else {
				gen8196 := Call(__e, ShenFunc(symcons_2), V1696)

				if True == gen8196 {
					gen8192 := Call(__e, ShenFunc(symhd), V1696)

					gen8193 := Call(__e, ShenFunc(symshen_4internal_1symbols), V1695, gen8192)

					gen8194 := Call(__e, ShenFunc(symtl), V1696)

					gen8195 := Call(__e, ShenFunc(symshen_4internal_1symbols), V1695, gen8194)

					__e.TailApply(ShenFunc(symunion), gen8193, gen8195)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.internal-symbols"), gen8201)

		gen8229 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1713 := __args[0]
			_ = V1713
			V1714 := __args[1]
			_ = V1714
			V1715 := __args[2]
			_ = V1715
			V1716 := __args[3]
			_ = V1716
			gen8228 := Call(__e, ShenFunc(symcons_2), V1715)

			if True == gen8228 {
				gen8224 := Call(__e, ShenFunc(symhd), V1715)

				gen8225 := Call(__e, ShenFunc(symshen_4packageh), V1713, V1714, gen8224, V1716)

				gen8226 := Call(__e, ShenFunc(symtl), V1715)

				gen8227 := Call(__e, ShenFunc(symshen_4packageh), V1713, V1714, gen8226, V1716)

				__e.TailApply(ShenFunc(symcons), gen8225, gen8227)

				return

			} else {
				gen8222 := Call(__e, ShenFunc(symshen_4sysfunc_2), V1715)

				var gen8223 Obj
				if True == gen8222 {
					gen8223 = True
				} else {
					gen8220 := Call(__e, ShenFunc(symvariable_2), V1715)

					var gen8221 Obj
					if True == gen8220 {
						gen8221 = True
					} else {
						gen8218 := Call(__e, ShenFunc(symelement_2), V1715, V1714)

						var gen8219 Obj
						if True == gen8218 {
							gen8219 = True
						} else {
							gen8216 := Call(__e, ShenFunc(symshen_4doubleunderline_2), V1715)

							var gen8217 Obj
							if True == gen8216 {
								gen8217 = True
							} else {
								gen8215 := Call(__e, ShenFunc(symshen_4singleunderline_2), V1715)

								if True == gen8215 {
									gen8217 = True
								} else {
									gen8217 = False
								}

							}
							if True == gen8217 {
								gen8219 = True
							} else {
								gen8219 = False
							}

						}
						if True == gen8219 {
							gen8221 = True
						} else {
							gen8221 = False
						}

					}
					if True == gen8221 {
						gen8223 = True
					} else {
						gen8223 = False
					}

				}
				if True == gen8223 {
					__e.Return(V1715)
					return
				} else {
					gen8213 := Call(__e, ShenFunc(symsymbol_2), V1715)

					var gen8214 Obj
					if True == gen8213 {
						gen8202 := Call(__e, ShenFunc(symexplode), V1715)

						ExplodeX := gen8202
						gen8205 := Call(__e, ShenFunc(symcons), MakeString("."), Nil)

						gen8206 := Call(__e, ShenFunc(symcons), MakeString("n"), gen8205)

						gen8207 := Call(__e, ShenFunc(symcons), MakeString("e"), gen8206)

						gen8208 := Call(__e, ShenFunc(symcons), MakeString("h"), gen8207)

						gen8209 := Call(__e, ShenFunc(symcons), MakeString("s"), gen8208)

						gen8210 := Call(__e, ShenFunc(symshen_4prefix_2), gen8209, ExplodeX)

						gen8211 := Call(__e, ShenFunc(symnot), gen8210)

						var gen8212 Obj
						if True == gen8211 {
							gen8203 := Call(__e, ShenFunc(symshen_4prefix_2), V1716, ExplodeX)

							gen8204 := Call(__e, ShenFunc(symnot), gen8203)

							if True == gen8204 {
								gen8212 = True
							} else {
								gen8212 = False
							}

						} else {
							gen8212 = False
						}
						if True == gen8212 {
							gen8214 = True
						} else {
							gen8214 = False
						}

					} else {
						gen8214 = False
					}
					if True == gen8214 {
						__e.TailApply(ShenFunc(symconcat), V1713, V1715)

						return
					} else {
						__e.Return(V1715)
						return
					}

				}

			}

		}, 4)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.packageh"), gen8229)

		return

	}, 0))
}
