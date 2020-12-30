package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen4866 := MakeNative(func(__e Evaluator) {
			V1430 := __e.Get(1)
			_ = V1430
			__e.TailApply(ShenFunc(symread_1byte), V1430)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-char-code"), gen4866)

		gen4868 := MakeNative(func(__e Evaluator) {
			V1432 := __e.Get(1)
			_ = V1432
			gen4867 := MakeNative(func(__e Evaluator) {
				S := __e.Get(1)
				_ = S
				__e.TailApply(ShenFunc(symread_1byte), S)

				return
			}, 1)
			__e.TailApply(ShenFunc(symshen_4read_1file_1as_1Xlist), V1432, gen4867)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-file-as-bytelist"), gen4868)

		gen4870 := MakeNative(func(__e Evaluator) {
			V1434 := __e.Get(1)
			_ = V1434
			gen4869 := MakeNative(func(__e Evaluator) {
				S := __e.Get(1)
				_ = S
				__e.TailApply(ShenFunc(symshen_4read_1char_1code), S)

				return
			}, 1)
			__e.TailApply(ShenFunc(symshen_4read_1file_1as_1Xlist), V1434, gen4869)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-file-as-charlist"), gen4870)

		gen4875 := MakeNative(func(__e Evaluator) {
			V1437 := __e.Get(1)
			_ = V1437
			V1438 := __e.Get(2)
			_ = V1438
			gen4871 := Call(__e, ShenFunc(symopen), V1437, MakeSymbol("in"))

			Stream := gen4871
			gen4872 := Call(__e, V1438, Stream)

			X := gen4872
			gen4873 := Call(__e, ShenFunc(symshen_4read_1file_1as_1Xlist_1help), Stream, V1438, X, Nil)

			Xs := gen4873
			gen4874 := Call(__e, ShenFunc(symclose), Stream)

			Close := gen4874
			_ = Close
			__e.TailApply(ShenFunc(symreverse), Xs)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-file-as-Xlist"), gen4875)

		gen4879 := MakeNative(func(__e Evaluator) {
			V1443 := __e.Get(1)
			_ = V1443
			V1444 := __e.Get(2)
			_ = V1444
			V1445 := __e.Get(3)
			_ = V1445
			V1446 := __e.Get(4)
			_ = V1446
			gen4878 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1445)

			if True == gen4878 {
				__e.Return(V1446)
				return
			} else {
				gen4876 := Call(__e, V1444, V1443)

				gen4877 := Call(__e, ShenFunc(symcons), V1445, V1446)

				__e.TailApply(ShenFunc(symshen_4read_1file_1as_1Xlist_1help), V1443, V1444, gen4876, gen4877)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-file-as-Xlist-help"), gen4879)

		gen4882 := MakeNative(func(__e Evaluator) {
			V1448 := __e.Get(1)
			_ = V1448
			gen4880 := Call(__e, ShenFunc(symopen), V1448, MakeSymbol("in"))

			Stream := gen4880
			gen4881 := Call(__e, ShenFunc(symshen_4read_1char_1code), Stream)

			__e.TailApply(ShenFunc(symshen_4rfas_1h), Stream, gen4881, MakeString(""))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-file-as-string"), gen4882)

		gen4887 := MakeNative(func(__e Evaluator) {
			V1452 := __e.Get(1)
			_ = V1452
			V1453 := __e.Get(2)
			_ = V1453
			V1454 := __e.Get(3)
			_ = V1454
			gen4886 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1453)

			if True == gen4886 {
				Call(__e, ShenFunc(symclose), V1452)
				__e.Return(V1454)
				return

			} else {
				gen4883 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1452)

				gen4884 := Call(__e, ShenFunc(symn_1_6string), V1453)

				gen4885 := Call(__e, ShenFunc(symcn), V1454, gen4884)

				__e.TailApply(ShenFunc(symshen_4rfas_1h), V1452, gen4883, gen4885)

				return

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rfas-h"), gen4887)

		gen4889 := MakeNative(func(__e Evaluator) {
			V1456 := __e.Get(1)
			_ = V1456
			gen4888 := Call(__e, ShenFunc(symread), V1456)

			__e.TailApply(ShenFunc(symeval_1kl), gen4888)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("input"), gen4889)

		gen4899 := MakeNative(func(__e Evaluator) {
			V1459 := __e.Get(1)
			_ = V1459
			V1460 := __e.Get(2)
			_ = V1460
			gen4890 := Call(__e, ShenFunc(symshen_4monotype), V1459)

			Mono_2 := gen4890
			_ = Mono_2
			gen4891 := Call(__e, ShenFunc(symread), V1460)

			Input := gen4891
			gen4896 := Call(__e, ShenFunc(symshen_4demodulate), V1459)

			gen4897 := Call(__e, ShenFunc(symshen_4typecheck), Input, gen4896)

			gen4898 := Call(__e, ShenFunc(sym_a), False, gen4897)

			if True == gen4898 {
				gen4892 := Call(__e, ShenFunc(symshen_4app), V1459, MakeString("\n"), MakeSymbol("shen.r"))

				gen4893 := Call(__e, ShenFunc(symcn), MakeString(" is not of type "), gen4892)

				gen4894 := Call(__e, ShenFunc(symshen_4app), Input, gen4893, MakeSymbol("shen.r"))

				gen4895 := Call(__e, ShenFunc(symcn), MakeString("type error: "), gen4894)

				__e.TailApply(ShenFunc(symsimple_1error), gen4895)

				return

			} else {
				__e.TailApply(ShenFunc(symeval_1kl), Input)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("input+"), gen4899)

		gen4905 := MakeNative(func(__e Evaluator) {
			V1462 := __e.Get(1)
			_ = V1462
			gen4904 := Call(__e, ShenFunc(symcons_2), V1462)

			if True == gen4904 {
				gen4903 := MakeNative(func(__e Evaluator) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(ShenFunc(symshen_4monotype), Z)

					return
				}, 1)
				__e.TailApply(ShenFunc(symmap), gen4903, V1462)

				return

			} else {
				gen4902 := Call(__e, ShenFunc(symvariable_2), V1462)

				if True == gen4902 {
					gen4900 := Call(__e, ShenFunc(symshen_4app), V1462, MakeString("\n"), MakeSymbol("shen.a"))

					gen4901 := Call(__e, ShenFunc(symcn), MakeString("input+ expects a monotype: not "), gen4900)

					__e.TailApply(ShenFunc(symsimple_1error), gen4901)

					return

				} else {
					__e.Return(V1462)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.monotype"), gen4905)

		gen4908 := MakeNative(func(__e Evaluator) {
			V1464 := __e.Get(1)
			_ = V1464
			gen4906 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1464)

			gen4907 := Call(__e, ShenFunc(symshen_4read_1loop), V1464, gen4906, Nil)

			__e.TailApply(ShenFunc(symhd), gen4907)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read"), gen4908)

		gen4909 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*it*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("it"), gen4909)

		gen4929 := MakeNative(func(__e Evaluator) {
			V1472 := __e.Get(1)
			_ = V1472
			V1473 := __e.Get(2)
			_ = V1473
			V1474 := __e.Get(3)
			_ = V1474
			gen4928 := Call(__e, ShenFunc(sym_a), MakeNumber(94), V1473)

			if True == gen4928 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("read aborted"))

				return
			} else {
				gen4927 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1473)

				if True == gen4927 {
					gen4926 := Call(__e, ShenFunc(symempty_2), V1474)

					if True == gen4926 {
						__e.TailApply(ShenFunc(symsimple_1error), MakeString("error: empty stream"))

						return
					} else {
						gen4924 := MakeNative(func(__e Evaluator) {
							X := __e.Get(1)
							_ = X
							__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

							return
						}, 1)
						gen4925 := MakeNative(func(__e Evaluator) {
							E := __e.Get(1)
							_ = E
							__e.Return(E)
							return
						}, 1)
						__e.TailApply(ShenFunc(symcompile), gen4924, V1474, gen4925)

						return

					}

				} else {
					gen4923 := Call(__e, ShenFunc(symshen_4terminator_2), V1473)

					if True == gen4923 {
						gen4913 := Call(__e, ShenFunc(symcons), V1473, Nil)

						gen4914 := Call(__e, ShenFunc(symappend), V1474, gen4913)

						AllChars := gen4914
						gen4915 := Call(__e, ShenFunc(symshen_4record_1it), AllChars)

						It := gen4915
						_ = It
						gen4916 := MakeNative(func(__e Evaluator) {
							X := __e.Get(1)
							_ = X
							__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

							return
						}, 1)
						gen4917 := MakeNative(func(__e Evaluator) {
							E := __e.Get(1)
							_ = E
							__e.Return(MakeSymbol("shen.nextbyte"))
							return
						}, 1)
						gen4918 := Call(__e, ShenFunc(symcompile), gen4916, AllChars, gen4917)

						Read := gen4918
						gen4921 := Call(__e, ShenFunc(sym_a), Read, MakeSymbol("shen.nextbyte"))

						var gen4922 Obj
						if True == gen4921 {
							gen4922 = True
						} else {
							gen4920 := Call(__e, ShenFunc(symempty_2), Read)

							if True == gen4920 {
								gen4922 = True
							} else {
								gen4922 = False
							}

						}
						if True == gen4922 {
							gen4919 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1472)

							__e.TailApply(ShenFunc(symshen_4read_1loop), V1472, gen4919, AllChars)

							return

						} else {
							__e.Return(Read)
							return
						}

					} else {
						gen4910 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1472)

						gen4911 := Call(__e, ShenFunc(symcons), V1473, Nil)

						gen4912 := Call(__e, ShenFunc(symappend), V1474, gen4911)

						__e.TailApply(ShenFunc(symshen_4read_1loop), V1472, gen4910, gen4912)

						return

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-loop"), gen4929)

		gen4937 := MakeNative(func(__e Evaluator) {
			V1476 := __e.Get(1)
			_ = V1476
			gen4930 := Call(__e, ShenFunc(symcons), MakeNumber(93), Nil)

			gen4931 := Call(__e, ShenFunc(symcons), MakeNumber(41), gen4930)

			gen4932 := Call(__e, ShenFunc(symcons), MakeNumber(34), gen4931)

			gen4933 := Call(__e, ShenFunc(symcons), MakeNumber(32), gen4932)

			gen4934 := Call(__e, ShenFunc(symcons), MakeNumber(13), gen4933)

			gen4935 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen4934)

			gen4936 := Call(__e, ShenFunc(symcons), MakeNumber(9), gen4935)

			__e.TailApply(ShenFunc(symelement_2), V1476, gen4936)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.terminator?"), gen4937)

		gen4939 := MakeNative(func(__e Evaluator) {
			V1478 := __e.Get(1)
			_ = V1478
			gen4938 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1478)

			__e.TailApply(ShenFunc(symshen_4lineread_1loop), gen4938, Nil, V1478)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("lineread"), gen4939)

		gen4964 := MakeNative(func(__e Evaluator) {
			V1483 := __e.Get(1)
			_ = V1483
			V1484 := __e.Get(2)
			_ = V1484
			V1485 := __e.Get(3)
			_ = V1485
			gen4963 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V1483)

			if True == gen4963 {
				gen4962 := Call(__e, ShenFunc(symempty_2), V1484)

				if True == gen4962 {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("empty stream"))

					return
				} else {
					gen4960 := MakeNative(func(__e Evaluator) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

						return
					}, 1)
					gen4961 := MakeNative(func(__e Evaluator) {
						E := __e.Get(1)
						_ = E
						__e.Return(E)
						return
					}, 1)
					__e.TailApply(ShenFunc(symcompile), gen4960, V1484, gen4961)

					return

				}

			} else {
				gen4958 := Call(__e, ShenFunc(symshen_4hat))

				gen4959 := Call(__e, ShenFunc(sym_a), V1483, gen4958)

				if True == gen4959 {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("line read aborted"))

					return
				} else {
					gen4953 := Call(__e, ShenFunc(symshen_4newline))

					gen4954 := Call(__e, ShenFunc(symshen_4carriage_1return))

					gen4955 := Call(__e, ShenFunc(symcons), gen4954, Nil)

					gen4956 := Call(__e, ShenFunc(symcons), gen4953, gen4955)

					gen4957 := Call(__e, ShenFunc(symelement_2), V1483, gen4956)

					if True == gen4957 {
						gen4943 := MakeNative(func(__e Evaluator) {
							X := __e.Get(1)
							_ = X
							__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

							return
						}, 1)
						gen4944 := MakeNative(func(__e Evaluator) {
							E := __e.Get(1)
							_ = E
							__e.Return(MakeSymbol("shen.nextline"))
							return
						}, 1)
						gen4945 := Call(__e, ShenFunc(symcompile), gen4943, V1484, gen4944)

						Line := gen4945
						gen4946 := Call(__e, ShenFunc(symshen_4record_1it), V1484)

						It := gen4946
						_ = It
						gen4951 := Call(__e, ShenFunc(sym_a), Line, MakeSymbol("shen.nextline"))

						var gen4952 Obj
						if True == gen4951 {
							gen4952 = True
						} else {
							gen4950 := Call(__e, ShenFunc(symempty_2), Line)

							if True == gen4950 {
								gen4952 = True
							} else {
								gen4952 = False
							}

						}
						if True == gen4952 {
							gen4947 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1485)

							gen4948 := Call(__e, ShenFunc(symcons), V1483, Nil)

							gen4949 := Call(__e, ShenFunc(symappend), V1484, gen4948)

							__e.TailApply(ShenFunc(symshen_4lineread_1loop), gen4947, gen4949, V1485)

							return

						} else {
							__e.Return(Line)
							return
						}

					} else {
						gen4940 := Call(__e, ShenFunc(symshen_4read_1char_1code), V1485)

						gen4941 := Call(__e, ShenFunc(symcons), V1483, Nil)

						gen4942 := Call(__e, ShenFunc(symappend), V1484, gen4941)

						__e.TailApply(ShenFunc(symshen_4lineread_1loop), gen4940, gen4942, V1485)

						return

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lineread-loop"), gen4964)

		gen4969 := MakeNative(func(__e Evaluator) {
			V1487 := __e.Get(1)
			_ = V1487
			gen4965 := Call(__e, ShenFunc(symshen_4trim_1whitespace), V1487)

			TrimLeft := gen4965
			gen4966 := Call(__e, ShenFunc(symreverse), TrimLeft)

			gen4967 := Call(__e, ShenFunc(symshen_4trim_1whitespace), gen4966)

			TrimRight := gen4967
			gen4968 := Call(__e, ShenFunc(symreverse), TrimRight)

			Trimmed := gen4968
			__e.TailApply(ShenFunc(symshen_4record_1it_1h), Trimmed)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-it"), gen4969)

		gen4979 := MakeNative(func(__e Evaluator) {
			V1489 := __e.Get(1)
			_ = V1489
			gen4977 := Call(__e, ShenFunc(symcons_2), V1489)

			var gen4978 Obj
			if True == gen4977 {
				gen4971 := Call(__e, ShenFunc(symhd), V1489)

				gen4972 := Call(__e, ShenFunc(symcons), MakeNumber(32), Nil)

				gen4973 := Call(__e, ShenFunc(symcons), MakeNumber(13), gen4972)

				gen4974 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen4973)

				gen4975 := Call(__e, ShenFunc(symcons), MakeNumber(9), gen4974)

				gen4976 := Call(__e, ShenFunc(symelement_2), gen4971, gen4975)

				if True == gen4976 {
					gen4978 = True
				} else {
					gen4978 = False
				}

			} else {
				gen4978 = False
			}
			if True == gen4978 {
				gen4970 := Call(__e, ShenFunc(symtl), V1489)

				__e.TailApply(ShenFunc(symshen_4trim_1whitespace), gen4970)

				return

			} else {
				__e.Return(V1489)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.trim-whitespace"), gen4979)

		gen4983 := MakeNative(func(__e Evaluator) {
			V1491 := __e.Get(1)
			_ = V1491
			gen4980 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symn_1_6string), X)

				return
			}, 1)
			gen4981 := Call(__e, ShenFunc(symmap), gen4980, V1491)

			gen4982 := Call(__e, ShenFunc(symshen_4cn_1all), gen4981)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*it*"), gen4982)

			__e.Return(V1491)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-it-h"), gen4983)

		gen4989 := MakeNative(func(__e Evaluator) {
			V1493 := __e.Get(1)
			_ = V1493
			gen4988 := Call(__e, ShenFunc(sym_a), Nil, V1493)

			if True == gen4988 {
				__e.Return(MakeString(""))
				return
			} else {
				gen4987 := Call(__e, ShenFunc(symcons_2), V1493)

				if True == gen4987 {
					gen4984 := Call(__e, ShenFunc(symhd), V1493)

					gen4985 := Call(__e, ShenFunc(symtl), V1493)

					gen4986 := Call(__e, ShenFunc(symshen_4cn_1all), gen4985)

					__e.TailApply(ShenFunc(symcn), gen4984, gen4986)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.cn-all"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cn-all"), gen4989)

		gen4993 := MakeNative(func(__e Evaluator) {
			V1495 := __e.Get(1)
			_ = V1495
			gen4990 := Call(__e, ShenFunc(symshen_4read_1file_1as_1charlist), V1495)

			Charlist := gen4990
			gen4991 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

				return
			}, 1)
			gen4992 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4read_1error), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symcompile), gen4991, Charlist, gen4992)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-file"), gen4993)

		gen4999 := MakeNative(func(__e Evaluator) {
			V1497 := __e.Get(1)
			_ = V1497
			gen4994 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symstring_1_6n), X)

				return
			}, 1)
			gen4995 := Call(__e, ShenFunc(symexplode), V1497)

			gen4996 := Call(__e, ShenFunc(symmap), gen4994, gen4995)

			Ns := gen4996
			gen4997 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

				return
			}, 1)
			gen4998 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4read_1error), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symcompile), gen4997, Ns, gen4998)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("read-from-string"), gen4999)

		gen5015 := MakeNative(func(__e Evaluator) {
			V1505 := __e.Get(1)
			_ = V1505
			gen5013 := Call(__e, ShenFunc(symcons_2), V1505)

			var gen5014 Obj
			if True == gen5013 {
				gen5010 := Call(__e, ShenFunc(symhd), V1505)

				gen5011 := Call(__e, ShenFunc(symcons_2), gen5010)

				var gen5012 Obj
				if True == gen5011 {
					gen5007 := Call(__e, ShenFunc(symtl), V1505)

					gen5008 := Call(__e, ShenFunc(symcons_2), gen5007)

					var gen5009 Obj
					if True == gen5008 {
						gen5004 := Call(__e, ShenFunc(symtl), V1505)

						gen5005 := Call(__e, ShenFunc(symtl), gen5004)

						gen5006 := Call(__e, ShenFunc(sym_a), Nil, gen5005)

						if True == gen5006 {
							gen5009 = True
						} else {
							gen5009 = False
						}

					} else {
						gen5009 = False
					}
					if True == gen5009 {
						gen5012 = True
					} else {
						gen5012 = False
					}

				} else {
					gen5012 = False
				}
				if True == gen5012 {
					gen5014 = True
				} else {
					gen5014 = False
				}

			} else {
				gen5014 = False
			}
			if True == gen5014 {
				gen5000 := Call(__e, ShenFunc(symhd), V1505)

				gen5001 := Call(__e, ShenFunc(symshen_4compress_150), MakeNumber(50), gen5000)

				gen5002 := Call(__e, ShenFunc(symshen_4app), gen5001, MakeString("\n"), MakeSymbol("shen.a"))

				gen5003 := Call(__e, ShenFunc(symcn), MakeString("read error here:\n\n "), gen5002)

				__e.TailApply(ShenFunc(symsimple_1error), gen5003)

				return

			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("read error\n"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-error"), gen5015)

		gen5024 := MakeNative(func(__e Evaluator) {
			V1512 := __e.Get(1)
			_ = V1512
			V1513 := __e.Get(2)
			_ = V1513
			gen5023 := Call(__e, ShenFunc(sym_a), Nil, V1513)

			if True == gen5023 {
				__e.Return(MakeString(""))
				return
			} else {
				gen5022 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V1512)

				if True == gen5022 {
					__e.Return(MakeString(""))
					return
				} else {
					gen5021 := Call(__e, ShenFunc(symcons_2), V1513)

					if True == gen5021 {
						gen5016 := Call(__e, ShenFunc(symhd), V1513)

						gen5017 := Call(__e, ShenFunc(symn_1_6string), gen5016)

						gen5018 := Call(__e, ShenFunc(sym_1), V1512, MakeNumber(1))

						gen5019 := Call(__e, ShenFunc(symtl), V1513)

						gen5020 := Call(__e, ShenFunc(symshen_4compress_150), gen5018, gen5019)

						__e.TailApply(ShenFunc(symcn), gen5017, gen5020)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.compress-50"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compress-50"), gen5024)

		gen5242 := MakeNative(func(__e Evaluator) {
			V1515 := __e.Get(1)
			_ = V1515
			gen5025 := Call(__e, ShenFunc(symshen_4_5lsb_6), V1515)

			Parse__shen_4_5lsb_6 := gen5025
			gen5044 := Call(__e, ShenFunc(symfail))

			gen5045 := Call(__e, ShenFunc(sym_a), gen5044, Parse__shen_4_5lsb_6)

			gen5046 := Call(__e, ShenFunc(symnot), gen5045)

			var gen5047 Obj
			if True == gen5046 {
				gen5026 := Call(__e, ShenFunc(symshen_4_5st__input1_6), Parse__shen_4_5lsb_6)

				Parse__shen_4_5st__input1_6 := gen5026
				gen5041 := Call(__e, ShenFunc(symfail))

				gen5042 := Call(__e, ShenFunc(sym_a), gen5041, Parse__shen_4_5st__input1_6)

				gen5043 := Call(__e, ShenFunc(symnot), gen5042)

				if True == gen5043 {
					gen5027 := Call(__e, ShenFunc(symshen_4_5rsb_6), Parse__shen_4_5st__input1_6)

					Parse__shen_4_5rsb_6 := gen5027
					gen5038 := Call(__e, ShenFunc(symfail))

					gen5039 := Call(__e, ShenFunc(sym_a), gen5038, Parse__shen_4_5rsb_6)

					gen5040 := Call(__e, ShenFunc(symnot), gen5039)

					if True == gen5040 {
						gen5028 := Call(__e, ShenFunc(symshen_4_5st__input2_6), Parse__shen_4_5rsb_6)

						Parse__shen_4_5st__input2_6 := gen5028
						gen5035 := Call(__e, ShenFunc(symfail))

						gen5036 := Call(__e, ShenFunc(sym_a), gen5035, Parse__shen_4_5st__input2_6)

						gen5037 := Call(__e, ShenFunc(symnot), gen5036)

						if True == gen5037 {
							gen5029 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input2_6)

							gen5030 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input1_6)

							gen5031 := Call(__e, ShenFunc(symshen_4cons__form), gen5030)

							gen5032 := Call(__e, ShenFunc(symmacroexpand), gen5031)

							gen5033 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input2_6)

							gen5034 := Call(__e, ShenFunc(symcons), gen5032, gen5033)

							gen5047 = Call(__e, ShenFunc(symshen_4pair), gen5029, gen5034)

						} else {
							gen5047 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen5047 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen5047 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen5047 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5047
			gen5240 := Call(__e, ShenFunc(symfail))

			gen5241 := Call(__e, ShenFunc(sym_a), YaccParse, gen5240)

			if True == gen5241 {
				gen5048 := Call(__e, ShenFunc(symshen_4_5lrb_6), V1515)

				Parse__shen_4_5lrb_6 := gen5048
				gen5066 := Call(__e, ShenFunc(symfail))

				gen5067 := Call(__e, ShenFunc(sym_a), gen5066, Parse__shen_4_5lrb_6)

				gen5068 := Call(__e, ShenFunc(symnot), gen5067)

				var gen5069 Obj
				if True == gen5068 {
					gen5049 := Call(__e, ShenFunc(symshen_4_5st__input1_6), Parse__shen_4_5lrb_6)

					Parse__shen_4_5st__input1_6 := gen5049
					gen5063 := Call(__e, ShenFunc(symfail))

					gen5064 := Call(__e, ShenFunc(sym_a), gen5063, Parse__shen_4_5st__input1_6)

					gen5065 := Call(__e, ShenFunc(symnot), gen5064)

					if True == gen5065 {
						gen5050 := Call(__e, ShenFunc(symshen_4_5rrb_6), Parse__shen_4_5st__input1_6)

						Parse__shen_4_5rrb_6 := gen5050
						gen5060 := Call(__e, ShenFunc(symfail))

						gen5061 := Call(__e, ShenFunc(sym_a), gen5060, Parse__shen_4_5rrb_6)

						gen5062 := Call(__e, ShenFunc(symnot), gen5061)

						if True == gen5062 {
							gen5051 := Call(__e, ShenFunc(symshen_4_5st__input2_6), Parse__shen_4_5rrb_6)

							Parse__shen_4_5st__input2_6 := gen5051
							gen5057 := Call(__e, ShenFunc(symfail))

							gen5058 := Call(__e, ShenFunc(sym_a), gen5057, Parse__shen_4_5st__input2_6)

							gen5059 := Call(__e, ShenFunc(symnot), gen5058)

							if True == gen5059 {
								gen5052 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input2_6)

								gen5053 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input1_6)

								gen5054 := Call(__e, ShenFunc(symmacroexpand), gen5053)

								gen5055 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input2_6)

								gen5056 := Call(__e, ShenFunc(symshen_4package_1macro), gen5054, gen5055)

								gen5069 = Call(__e, ShenFunc(symshen_4pair), gen5052, gen5056)

							} else {
								gen5069 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen5069 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen5069 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen5069 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen5069
				gen5238 := Call(__e, ShenFunc(symfail))

				gen5239 := Call(__e, ShenFunc(sym_a), YaccParse, gen5238)

				if True == gen5239 {
					gen5070 := Call(__e, ShenFunc(symshen_4_5lcurly_6), V1515)

					Parse__shen_4_5lcurly_6 := gen5070
					gen5078 := Call(__e, ShenFunc(symfail))

					gen5079 := Call(__e, ShenFunc(sym_a), gen5078, Parse__shen_4_5lcurly_6)

					gen5080 := Call(__e, ShenFunc(symnot), gen5079)

					var gen5081 Obj
					if True == gen5080 {
						gen5071 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5lcurly_6)

						Parse__shen_4_5st__input_6 := gen5071
						gen5075 := Call(__e, ShenFunc(symfail))

						gen5076 := Call(__e, ShenFunc(sym_a), gen5075, Parse__shen_4_5st__input_6)

						gen5077 := Call(__e, ShenFunc(symnot), gen5076)

						if True == gen5077 {
							gen5072 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

							gen5073 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

							gen5074 := Call(__e, ShenFunc(symcons), MakeSymbol("{"), gen5073)

							gen5081 = Call(__e, ShenFunc(symshen_4pair), gen5072, gen5074)

						} else {
							gen5081 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen5081 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen5081
					gen5236 := Call(__e, ShenFunc(symfail))

					gen5237 := Call(__e, ShenFunc(sym_a), YaccParse, gen5236)

					if True == gen5237 {
						gen5082 := Call(__e, ShenFunc(symshen_4_5rcurly_6), V1515)

						Parse__shen_4_5rcurly_6 := gen5082
						gen5090 := Call(__e, ShenFunc(symfail))

						gen5091 := Call(__e, ShenFunc(sym_a), gen5090, Parse__shen_4_5rcurly_6)

						gen5092 := Call(__e, ShenFunc(symnot), gen5091)

						var gen5093 Obj
						if True == gen5092 {
							gen5083 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5rcurly_6)

							Parse__shen_4_5st__input_6 := gen5083
							gen5087 := Call(__e, ShenFunc(symfail))

							gen5088 := Call(__e, ShenFunc(sym_a), gen5087, Parse__shen_4_5st__input_6)

							gen5089 := Call(__e, ShenFunc(symnot), gen5088)

							if True == gen5089 {
								gen5084 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

								gen5085 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

								gen5086 := Call(__e, ShenFunc(symcons), MakeSymbol("}"), gen5085)

								gen5093 = Call(__e, ShenFunc(symshen_4pair), gen5084, gen5086)

							} else {
								gen5093 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen5093 = Call(__e, ShenFunc(symfail))

						}
						YaccParse := gen5093
						gen5234 := Call(__e, ShenFunc(symfail))

						gen5235 := Call(__e, ShenFunc(sym_a), YaccParse, gen5234)

						if True == gen5235 {
							gen5094 := Call(__e, ShenFunc(symshen_4_5bar_6), V1515)

							Parse__shen_4_5bar_6 := gen5094
							gen5102 := Call(__e, ShenFunc(symfail))

							gen5103 := Call(__e, ShenFunc(sym_a), gen5102, Parse__shen_4_5bar_6)

							gen5104 := Call(__e, ShenFunc(symnot), gen5103)

							var gen5105 Obj
							if True == gen5104 {
								gen5095 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5bar_6)

								Parse__shen_4_5st__input_6 := gen5095
								gen5099 := Call(__e, ShenFunc(symfail))

								gen5100 := Call(__e, ShenFunc(sym_a), gen5099, Parse__shen_4_5st__input_6)

								gen5101 := Call(__e, ShenFunc(symnot), gen5100)

								if True == gen5101 {
									gen5096 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

									gen5097 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

									gen5098 := Call(__e, ShenFunc(symcons), MakeSymbol("bar!"), gen5097)

									gen5105 = Call(__e, ShenFunc(symshen_4pair), gen5096, gen5098)

								} else {
									gen5105 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen5105 = Call(__e, ShenFunc(symfail))

							}
							YaccParse := gen5105
							gen5232 := Call(__e, ShenFunc(symfail))

							gen5233 := Call(__e, ShenFunc(sym_a), YaccParse, gen5232)

							if True == gen5233 {
								gen5106 := Call(__e, ShenFunc(symshen_4_5semicolon_6), V1515)

								Parse__shen_4_5semicolon_6 := gen5106
								gen5114 := Call(__e, ShenFunc(symfail))

								gen5115 := Call(__e, ShenFunc(sym_a), gen5114, Parse__shen_4_5semicolon_6)

								gen5116 := Call(__e, ShenFunc(symnot), gen5115)

								var gen5117 Obj
								if True == gen5116 {
									gen5107 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5semicolon_6)

									Parse__shen_4_5st__input_6 := gen5107
									gen5111 := Call(__e, ShenFunc(symfail))

									gen5112 := Call(__e, ShenFunc(sym_a), gen5111, Parse__shen_4_5st__input_6)

									gen5113 := Call(__e, ShenFunc(symnot), gen5112)

									if True == gen5113 {
										gen5108 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

										gen5109 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

										gen5110 := Call(__e, ShenFunc(symcons), MakeSymbol(";"), gen5109)

										gen5117 = Call(__e, ShenFunc(symshen_4pair), gen5108, gen5110)

									} else {
										gen5117 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen5117 = Call(__e, ShenFunc(symfail))

								}
								YaccParse := gen5117
								gen5230 := Call(__e, ShenFunc(symfail))

								gen5231 := Call(__e, ShenFunc(sym_a), YaccParse, gen5230)

								if True == gen5231 {
									gen5118 := Call(__e, ShenFunc(symshen_4_5colon_6), V1515)

									Parse__shen_4_5colon_6 := gen5118
									gen5130 := Call(__e, ShenFunc(symfail))

									gen5131 := Call(__e, ShenFunc(sym_a), gen5130, Parse__shen_4_5colon_6)

									gen5132 := Call(__e, ShenFunc(symnot), gen5131)

									var gen5133 Obj
									if True == gen5132 {
										gen5119 := Call(__e, ShenFunc(symshen_4_5equal_6), Parse__shen_4_5colon_6)

										Parse__shen_4_5equal_6 := gen5119
										gen5127 := Call(__e, ShenFunc(symfail))

										gen5128 := Call(__e, ShenFunc(sym_a), gen5127, Parse__shen_4_5equal_6)

										gen5129 := Call(__e, ShenFunc(symnot), gen5128)

										if True == gen5129 {
											gen5120 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5equal_6)

											Parse__shen_4_5st__input_6 := gen5120
											gen5124 := Call(__e, ShenFunc(symfail))

											gen5125 := Call(__e, ShenFunc(sym_a), gen5124, Parse__shen_4_5st__input_6)

											gen5126 := Call(__e, ShenFunc(symnot), gen5125)

											if True == gen5126 {
												gen5121 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

												gen5122 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

												gen5123 := Call(__e, ShenFunc(symcons), MakeSymbol(":="), gen5122)

												gen5133 = Call(__e, ShenFunc(symshen_4pair), gen5121, gen5123)

											} else {
												gen5133 = Call(__e, ShenFunc(symfail))

											}

										} else {
											gen5133 = Call(__e, ShenFunc(symfail))

										}

									} else {
										gen5133 = Call(__e, ShenFunc(symfail))

									}
									YaccParse := gen5133
									gen5228 := Call(__e, ShenFunc(symfail))

									gen5229 := Call(__e, ShenFunc(sym_a), YaccParse, gen5228)

									if True == gen5229 {
										gen5134 := Call(__e, ShenFunc(symshen_4_5colon_6), V1515)

										Parse__shen_4_5colon_6 := gen5134
										gen5146 := Call(__e, ShenFunc(symfail))

										gen5147 := Call(__e, ShenFunc(sym_a), gen5146, Parse__shen_4_5colon_6)

										gen5148 := Call(__e, ShenFunc(symnot), gen5147)

										var gen5149 Obj
										if True == gen5148 {
											gen5135 := Call(__e, ShenFunc(symshen_4_5minus_6), Parse__shen_4_5colon_6)

											Parse__shen_4_5minus_6 := gen5135
											gen5143 := Call(__e, ShenFunc(symfail))

											gen5144 := Call(__e, ShenFunc(sym_a), gen5143, Parse__shen_4_5minus_6)

											gen5145 := Call(__e, ShenFunc(symnot), gen5144)

											if True == gen5145 {
												gen5136 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5minus_6)

												Parse__shen_4_5st__input_6 := gen5136
												gen5140 := Call(__e, ShenFunc(symfail))

												gen5141 := Call(__e, ShenFunc(sym_a), gen5140, Parse__shen_4_5st__input_6)

												gen5142 := Call(__e, ShenFunc(symnot), gen5141)

												if True == gen5142 {
													gen5137 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

													gen5138 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

													gen5139 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen5138)

													gen5149 = Call(__e, ShenFunc(symshen_4pair), gen5137, gen5139)

												} else {
													gen5149 = Call(__e, ShenFunc(symfail))

												}

											} else {
												gen5149 = Call(__e, ShenFunc(symfail))

											}

										} else {
											gen5149 = Call(__e, ShenFunc(symfail))

										}
										YaccParse := gen5149
										gen5226 := Call(__e, ShenFunc(symfail))

										gen5227 := Call(__e, ShenFunc(sym_a), YaccParse, gen5226)

										if True == gen5227 {
											gen5150 := Call(__e, ShenFunc(symshen_4_5colon_6), V1515)

											Parse__shen_4_5colon_6 := gen5150
											gen5158 := Call(__e, ShenFunc(symfail))

											gen5159 := Call(__e, ShenFunc(sym_a), gen5158, Parse__shen_4_5colon_6)

											gen5160 := Call(__e, ShenFunc(symnot), gen5159)

											var gen5161 Obj
											if True == gen5160 {
												gen5151 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5colon_6)

												Parse__shen_4_5st__input_6 := gen5151
												gen5155 := Call(__e, ShenFunc(symfail))

												gen5156 := Call(__e, ShenFunc(sym_a), gen5155, Parse__shen_4_5st__input_6)

												gen5157 := Call(__e, ShenFunc(symnot), gen5156)

												if True == gen5157 {
													gen5152 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

													gen5153 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

													gen5154 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen5153)

													gen5161 = Call(__e, ShenFunc(symshen_4pair), gen5152, gen5154)

												} else {
													gen5161 = Call(__e, ShenFunc(symfail))

												}

											} else {
												gen5161 = Call(__e, ShenFunc(symfail))

											}
											YaccParse := gen5161
											gen5224 := Call(__e, ShenFunc(symfail))

											gen5225 := Call(__e, ShenFunc(sym_a), YaccParse, gen5224)

											if True == gen5225 {
												gen5162 := Call(__e, ShenFunc(symshen_4_5comma_6), V1515)

												Parse__shen_4_5comma_6 := gen5162
												gen5171 := Call(__e, ShenFunc(symfail))

												gen5172 := Call(__e, ShenFunc(sym_a), gen5171, Parse__shen_4_5comma_6)

												gen5173 := Call(__e, ShenFunc(symnot), gen5172)

												var gen5174 Obj
												if True == gen5173 {
													gen5163 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5comma_6)

													Parse__shen_4_5st__input_6 := gen5163
													gen5168 := Call(__e, ShenFunc(symfail))

													gen5169 := Call(__e, ShenFunc(sym_a), gen5168, Parse__shen_4_5st__input_6)

													gen5170 := Call(__e, ShenFunc(symnot), gen5169)

													if True == gen5170 {
														gen5164 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

														gen5165 := Call(__e, ShenFunc(symintern), MakeString(","))

														gen5166 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

														gen5167 := Call(__e, ShenFunc(symcons), gen5165, gen5166)

														gen5174 = Call(__e, ShenFunc(symshen_4pair), gen5164, gen5167)

													} else {
														gen5174 = Call(__e, ShenFunc(symfail))

													}

												} else {
													gen5174 = Call(__e, ShenFunc(symfail))

												}
												YaccParse := gen5174
												gen5222 := Call(__e, ShenFunc(symfail))

												gen5223 := Call(__e, ShenFunc(sym_a), YaccParse, gen5222)

												if True == gen5223 {
													gen5175 := Call(__e, ShenFunc(symshen_4_5comment_6), V1515)

													Parse__shen_4_5comment_6 := gen5175
													gen5182 := Call(__e, ShenFunc(symfail))

													gen5183 := Call(__e, ShenFunc(sym_a), gen5182, Parse__shen_4_5comment_6)

													gen5184 := Call(__e, ShenFunc(symnot), gen5183)

													var gen5185 Obj
													if True == gen5184 {
														gen5176 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5comment_6)

														Parse__shen_4_5st__input_6 := gen5176
														gen5179 := Call(__e, ShenFunc(symfail))

														gen5180 := Call(__e, ShenFunc(sym_a), gen5179, Parse__shen_4_5st__input_6)

														gen5181 := Call(__e, ShenFunc(symnot), gen5180)

														if True == gen5181 {
															gen5177 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

															gen5178 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

															gen5185 = Call(__e, ShenFunc(symshen_4pair), gen5177, gen5178)

														} else {
															gen5185 = Call(__e, ShenFunc(symfail))

														}

													} else {
														gen5185 = Call(__e, ShenFunc(symfail))

													}
													YaccParse := gen5185
													gen5220 := Call(__e, ShenFunc(symfail))

													gen5221 := Call(__e, ShenFunc(sym_a), YaccParse, gen5220)

													if True == gen5221 {
														gen5186 := Call(__e, ShenFunc(symshen_4_5atom_6), V1515)

														Parse__shen_4_5atom_6 := gen5186
														gen5196 := Call(__e, ShenFunc(symfail))

														gen5197 := Call(__e, ShenFunc(sym_a), gen5196, Parse__shen_4_5atom_6)

														gen5198 := Call(__e, ShenFunc(symnot), gen5197)

														var gen5199 Obj
														if True == gen5198 {
															gen5187 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5atom_6)

															Parse__shen_4_5st__input_6 := gen5187
															gen5193 := Call(__e, ShenFunc(symfail))

															gen5194 := Call(__e, ShenFunc(sym_a), gen5193, Parse__shen_4_5st__input_6)

															gen5195 := Call(__e, ShenFunc(symnot), gen5194)

															if True == gen5195 {
																gen5188 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

																gen5189 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5atom_6)

																gen5190 := Call(__e, ShenFunc(symmacroexpand), gen5189)

																gen5191 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

																gen5192 := Call(__e, ShenFunc(symcons), gen5190, gen5191)

																gen5199 = Call(__e, ShenFunc(symshen_4pair), gen5188, gen5192)

															} else {
																gen5199 = Call(__e, ShenFunc(symfail))

															}

														} else {
															gen5199 = Call(__e, ShenFunc(symfail))

														}
														YaccParse := gen5199
														gen5218 := Call(__e, ShenFunc(symfail))

														gen5219 := Call(__e, ShenFunc(sym_a), YaccParse, gen5218)

														if True == gen5219 {
															gen5200 := Call(__e, ShenFunc(symshen_4_5whitespaces_6), V1515)

															Parse__shen_4_5whitespaces_6 := gen5200
															gen5207 := Call(__e, ShenFunc(symfail))

															gen5208 := Call(__e, ShenFunc(sym_a), gen5207, Parse__shen_4_5whitespaces_6)

															gen5209 := Call(__e, ShenFunc(symnot), gen5208)

															var gen5210 Obj
															if True == gen5209 {
																gen5201 := Call(__e, ShenFunc(symshen_4_5st__input_6), Parse__shen_4_5whitespaces_6)

																Parse__shen_4_5st__input_6 := gen5201
																gen5204 := Call(__e, ShenFunc(symfail))

																gen5205 := Call(__e, ShenFunc(sym_a), gen5204, Parse__shen_4_5st__input_6)

																gen5206 := Call(__e, ShenFunc(symnot), gen5205)

																if True == gen5206 {
																	gen5202 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

																	gen5203 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

																	gen5210 = Call(__e, ShenFunc(symshen_4pair), gen5202, gen5203)

																} else {
																	gen5210 = Call(__e, ShenFunc(symfail))

																}

															} else {
																gen5210 = Call(__e, ShenFunc(symfail))

															}
															YaccParse := gen5210
															gen5216 := Call(__e, ShenFunc(symfail))

															gen5217 := Call(__e, ShenFunc(sym_a), YaccParse, gen5216)

															if True == gen5217 {
																gen5211 := Call(__e, ShenFunc(sym_5e_6), V1515)

																Parse___5e_6 := gen5211
																gen5213 := Call(__e, ShenFunc(symfail))

																gen5214 := Call(__e, ShenFunc(sym_a), gen5213, Parse___5e_6)

																gen5215 := Call(__e, ShenFunc(symnot), gen5214)

																if True == gen5215 {
																	gen5212 := Call(__e, ShenFunc(symhd), Parse___5e_6)

																	__e.TailApply(ShenFunc(symshen_4pair), gen5212, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<st_input>"), gen5242)

		gen5252 := MakeNative(func(__e Evaluator) {
			V1518 := __e.Get(1)
			_ = V1518
			gen5249 := Call(__e, ShenFunc(symhd), V1518)

			gen5250 := Call(__e, ShenFunc(symcons_2), gen5249)

			var gen5251 Obj
			if True == gen5250 {
				gen5247 := Call(__e, ShenFunc(symshen_4hdhd), V1518)

				gen5248 := Call(__e, ShenFunc(sym_a), MakeNumber(91), gen5247)

				if True == gen5248 {
					gen5251 = True
				} else {
					gen5251 = False
				}

			} else {
				gen5251 = False
			}
			if True == gen5251 {
				gen5243 := Call(__e, ShenFunc(symshen_4tlhd), V1518)

				gen5244 := Call(__e, ShenFunc(symshen_4hdtl), V1518)

				gen5245 := Call(__e, ShenFunc(symshen_4pair), gen5243, gen5244)

				NewStream1516 := gen5245
				gen5246 := Call(__e, ShenFunc(symhd), NewStream1516)

				__e.TailApply(ShenFunc(symshen_4pair), gen5246, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<lsb>"), gen5252)

		gen5262 := MakeNative(func(__e Evaluator) {
			V1521 := __e.Get(1)
			_ = V1521
			gen5259 := Call(__e, ShenFunc(symhd), V1521)

			gen5260 := Call(__e, ShenFunc(symcons_2), gen5259)

			var gen5261 Obj
			if True == gen5260 {
				gen5257 := Call(__e, ShenFunc(symshen_4hdhd), V1521)

				gen5258 := Call(__e, ShenFunc(sym_a), MakeNumber(93), gen5257)

				if True == gen5258 {
					gen5261 = True
				} else {
					gen5261 = False
				}

			} else {
				gen5261 = False
			}
			if True == gen5261 {
				gen5253 := Call(__e, ShenFunc(symshen_4tlhd), V1521)

				gen5254 := Call(__e, ShenFunc(symshen_4hdtl), V1521)

				gen5255 := Call(__e, ShenFunc(symshen_4pair), gen5253, gen5254)

				NewStream1519 := gen5255
				gen5256 := Call(__e, ShenFunc(symhd), NewStream1519)

				__e.TailApply(ShenFunc(symshen_4pair), gen5256, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rsb>"), gen5262)

		gen5272 := MakeNative(func(__e Evaluator) {
			V1524 := __e.Get(1)
			_ = V1524
			gen5269 := Call(__e, ShenFunc(symhd), V1524)

			gen5270 := Call(__e, ShenFunc(symcons_2), gen5269)

			var gen5271 Obj
			if True == gen5270 {
				gen5267 := Call(__e, ShenFunc(symshen_4hdhd), V1524)

				gen5268 := Call(__e, ShenFunc(sym_a), MakeNumber(123), gen5267)

				if True == gen5268 {
					gen5271 = True
				} else {
					gen5271 = False
				}

			} else {
				gen5271 = False
			}
			if True == gen5271 {
				gen5263 := Call(__e, ShenFunc(symshen_4tlhd), V1524)

				gen5264 := Call(__e, ShenFunc(symshen_4hdtl), V1524)

				gen5265 := Call(__e, ShenFunc(symshen_4pair), gen5263, gen5264)

				NewStream1522 := gen5265
				gen5266 := Call(__e, ShenFunc(symhd), NewStream1522)

				__e.TailApply(ShenFunc(symshen_4pair), gen5266, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<lcurly>"), gen5272)

		gen5282 := MakeNative(func(__e Evaluator) {
			V1527 := __e.Get(1)
			_ = V1527
			gen5279 := Call(__e, ShenFunc(symhd), V1527)

			gen5280 := Call(__e, ShenFunc(symcons_2), gen5279)

			var gen5281 Obj
			if True == gen5280 {
				gen5277 := Call(__e, ShenFunc(symshen_4hdhd), V1527)

				gen5278 := Call(__e, ShenFunc(sym_a), MakeNumber(125), gen5277)

				if True == gen5278 {
					gen5281 = True
				} else {
					gen5281 = False
				}

			} else {
				gen5281 = False
			}
			if True == gen5281 {
				gen5273 := Call(__e, ShenFunc(symshen_4tlhd), V1527)

				gen5274 := Call(__e, ShenFunc(symshen_4hdtl), V1527)

				gen5275 := Call(__e, ShenFunc(symshen_4pair), gen5273, gen5274)

				NewStream1525 := gen5275
				gen5276 := Call(__e, ShenFunc(symhd), NewStream1525)

				__e.TailApply(ShenFunc(symshen_4pair), gen5276, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rcurly>"), gen5282)

		gen5292 := MakeNative(func(__e Evaluator) {
			V1530 := __e.Get(1)
			_ = V1530
			gen5289 := Call(__e, ShenFunc(symhd), V1530)

			gen5290 := Call(__e, ShenFunc(symcons_2), gen5289)

			var gen5291 Obj
			if True == gen5290 {
				gen5287 := Call(__e, ShenFunc(symshen_4hdhd), V1530)

				gen5288 := Call(__e, ShenFunc(sym_a), MakeNumber(124), gen5287)

				if True == gen5288 {
					gen5291 = True
				} else {
					gen5291 = False
				}

			} else {
				gen5291 = False
			}
			if True == gen5291 {
				gen5283 := Call(__e, ShenFunc(symshen_4tlhd), V1530)

				gen5284 := Call(__e, ShenFunc(symshen_4hdtl), V1530)

				gen5285 := Call(__e, ShenFunc(symshen_4pair), gen5283, gen5284)

				NewStream1528 := gen5285
				gen5286 := Call(__e, ShenFunc(symhd), NewStream1528)

				__e.TailApply(ShenFunc(symshen_4pair), gen5286, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<bar>"), gen5292)

		gen5302 := MakeNative(func(__e Evaluator) {
			V1533 := __e.Get(1)
			_ = V1533
			gen5299 := Call(__e, ShenFunc(symhd), V1533)

			gen5300 := Call(__e, ShenFunc(symcons_2), gen5299)

			var gen5301 Obj
			if True == gen5300 {
				gen5297 := Call(__e, ShenFunc(symshen_4hdhd), V1533)

				gen5298 := Call(__e, ShenFunc(sym_a), MakeNumber(59), gen5297)

				if True == gen5298 {
					gen5301 = True
				} else {
					gen5301 = False
				}

			} else {
				gen5301 = False
			}
			if True == gen5301 {
				gen5293 := Call(__e, ShenFunc(symshen_4tlhd), V1533)

				gen5294 := Call(__e, ShenFunc(symshen_4hdtl), V1533)

				gen5295 := Call(__e, ShenFunc(symshen_4pair), gen5293, gen5294)

				NewStream1531 := gen5295
				gen5296 := Call(__e, ShenFunc(symhd), NewStream1531)

				__e.TailApply(ShenFunc(symshen_4pair), gen5296, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<semicolon>"), gen5302)

		gen5312 := MakeNative(func(__e Evaluator) {
			V1536 := __e.Get(1)
			_ = V1536
			gen5309 := Call(__e, ShenFunc(symhd), V1536)

			gen5310 := Call(__e, ShenFunc(symcons_2), gen5309)

			var gen5311 Obj
			if True == gen5310 {
				gen5307 := Call(__e, ShenFunc(symshen_4hdhd), V1536)

				gen5308 := Call(__e, ShenFunc(sym_a), MakeNumber(58), gen5307)

				if True == gen5308 {
					gen5311 = True
				} else {
					gen5311 = False
				}

			} else {
				gen5311 = False
			}
			if True == gen5311 {
				gen5303 := Call(__e, ShenFunc(symshen_4tlhd), V1536)

				gen5304 := Call(__e, ShenFunc(symshen_4hdtl), V1536)

				gen5305 := Call(__e, ShenFunc(symshen_4pair), gen5303, gen5304)

				NewStream1534 := gen5305
				gen5306 := Call(__e, ShenFunc(symhd), NewStream1534)

				__e.TailApply(ShenFunc(symshen_4pair), gen5306, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<colon>"), gen5312)

		gen5322 := MakeNative(func(__e Evaluator) {
			V1539 := __e.Get(1)
			_ = V1539
			gen5319 := Call(__e, ShenFunc(symhd), V1539)

			gen5320 := Call(__e, ShenFunc(symcons_2), gen5319)

			var gen5321 Obj
			if True == gen5320 {
				gen5317 := Call(__e, ShenFunc(symshen_4hdhd), V1539)

				gen5318 := Call(__e, ShenFunc(sym_a), MakeNumber(44), gen5317)

				if True == gen5318 {
					gen5321 = True
				} else {
					gen5321 = False
				}

			} else {
				gen5321 = False
			}
			if True == gen5321 {
				gen5313 := Call(__e, ShenFunc(symshen_4tlhd), V1539)

				gen5314 := Call(__e, ShenFunc(symshen_4hdtl), V1539)

				gen5315 := Call(__e, ShenFunc(symshen_4pair), gen5313, gen5314)

				NewStream1537 := gen5315
				gen5316 := Call(__e, ShenFunc(symhd), NewStream1537)

				__e.TailApply(ShenFunc(symshen_4pair), gen5316, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<comma>"), gen5322)

		gen5332 := MakeNative(func(__e Evaluator) {
			V1542 := __e.Get(1)
			_ = V1542
			gen5329 := Call(__e, ShenFunc(symhd), V1542)

			gen5330 := Call(__e, ShenFunc(symcons_2), gen5329)

			var gen5331 Obj
			if True == gen5330 {
				gen5327 := Call(__e, ShenFunc(symshen_4hdhd), V1542)

				gen5328 := Call(__e, ShenFunc(sym_a), MakeNumber(61), gen5327)

				if True == gen5328 {
					gen5331 = True
				} else {
					gen5331 = False
				}

			} else {
				gen5331 = False
			}
			if True == gen5331 {
				gen5323 := Call(__e, ShenFunc(symshen_4tlhd), V1542)

				gen5324 := Call(__e, ShenFunc(symshen_4hdtl), V1542)

				gen5325 := Call(__e, ShenFunc(symshen_4pair), gen5323, gen5324)

				NewStream1540 := gen5325
				gen5326 := Call(__e, ShenFunc(symhd), NewStream1540)

				__e.TailApply(ShenFunc(symshen_4pair), gen5326, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<equal>"), gen5332)

		gen5342 := MakeNative(func(__e Evaluator) {
			V1545 := __e.Get(1)
			_ = V1545
			gen5339 := Call(__e, ShenFunc(symhd), V1545)

			gen5340 := Call(__e, ShenFunc(symcons_2), gen5339)

			var gen5341 Obj
			if True == gen5340 {
				gen5337 := Call(__e, ShenFunc(symshen_4hdhd), V1545)

				gen5338 := Call(__e, ShenFunc(sym_a), MakeNumber(45), gen5337)

				if True == gen5338 {
					gen5341 = True
				} else {
					gen5341 = False
				}

			} else {
				gen5341 = False
			}
			if True == gen5341 {
				gen5333 := Call(__e, ShenFunc(symshen_4tlhd), V1545)

				gen5334 := Call(__e, ShenFunc(symshen_4hdtl), V1545)

				gen5335 := Call(__e, ShenFunc(symshen_4pair), gen5333, gen5334)

				NewStream1543 := gen5335
				gen5336 := Call(__e, ShenFunc(symhd), NewStream1543)

				__e.TailApply(ShenFunc(symshen_4pair), gen5336, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<minus>"), gen5342)

		gen5352 := MakeNative(func(__e Evaluator) {
			V1548 := __e.Get(1)
			_ = V1548
			gen5349 := Call(__e, ShenFunc(symhd), V1548)

			gen5350 := Call(__e, ShenFunc(symcons_2), gen5349)

			var gen5351 Obj
			if True == gen5350 {
				gen5347 := Call(__e, ShenFunc(symshen_4hdhd), V1548)

				gen5348 := Call(__e, ShenFunc(sym_a), MakeNumber(40), gen5347)

				if True == gen5348 {
					gen5351 = True
				} else {
					gen5351 = False
				}

			} else {
				gen5351 = False
			}
			if True == gen5351 {
				gen5343 := Call(__e, ShenFunc(symshen_4tlhd), V1548)

				gen5344 := Call(__e, ShenFunc(symshen_4hdtl), V1548)

				gen5345 := Call(__e, ShenFunc(symshen_4pair), gen5343, gen5344)

				NewStream1546 := gen5345
				gen5346 := Call(__e, ShenFunc(symhd), NewStream1546)

				__e.TailApply(ShenFunc(symshen_4pair), gen5346, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<lrb>"), gen5352)

		gen5362 := MakeNative(func(__e Evaluator) {
			V1551 := __e.Get(1)
			_ = V1551
			gen5359 := Call(__e, ShenFunc(symhd), V1551)

			gen5360 := Call(__e, ShenFunc(symcons_2), gen5359)

			var gen5361 Obj
			if True == gen5360 {
				gen5357 := Call(__e, ShenFunc(symshen_4hdhd), V1551)

				gen5358 := Call(__e, ShenFunc(sym_a), MakeNumber(41), gen5357)

				if True == gen5358 {
					gen5361 = True
				} else {
					gen5361 = False
				}

			} else {
				gen5361 = False
			}
			if True == gen5361 {
				gen5353 := Call(__e, ShenFunc(symshen_4tlhd), V1551)

				gen5354 := Call(__e, ShenFunc(symshen_4hdtl), V1551)

				gen5355 := Call(__e, ShenFunc(symshen_4pair), gen5353, gen5354)

				NewStream1549 := gen5355
				gen5356 := Call(__e, ShenFunc(symhd), NewStream1549)

				__e.TailApply(ShenFunc(symshen_4pair), gen5356, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rrb>"), gen5362)

		gen5392 := MakeNative(func(__e Evaluator) {
			V1553 := __e.Get(1)
			_ = V1553
			gen5363 := Call(__e, ShenFunc(symshen_4_5str_6), V1553)

			Parse__shen_4_5str_6 := gen5363
			gen5367 := Call(__e, ShenFunc(symfail))

			gen5368 := Call(__e, ShenFunc(sym_a), gen5367, Parse__shen_4_5str_6)

			gen5369 := Call(__e, ShenFunc(symnot), gen5368)

			var gen5370 Obj
			if True == gen5369 {
				gen5364 := Call(__e, ShenFunc(symhd), Parse__shen_4_5str_6)

				gen5365 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5str_6)

				gen5366 := Call(__e, ShenFunc(symshen_4control_1chars), gen5365)

				gen5370 = Call(__e, ShenFunc(symshen_4pair), gen5364, gen5366)

			} else {
				gen5370 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5370
			gen5390 := Call(__e, ShenFunc(symfail))

			gen5391 := Call(__e, ShenFunc(sym_a), YaccParse, gen5390)

			if True == gen5391 {
				gen5371 := Call(__e, ShenFunc(symshen_4_5number_6), V1553)

				Parse__shen_4_5number_6 := gen5371
				gen5374 := Call(__e, ShenFunc(symfail))

				gen5375 := Call(__e, ShenFunc(sym_a), gen5374, Parse__shen_4_5number_6)

				gen5376 := Call(__e, ShenFunc(symnot), gen5375)

				var gen5377 Obj
				if True == gen5376 {
					gen5372 := Call(__e, ShenFunc(symhd), Parse__shen_4_5number_6)

					gen5373 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5number_6)

					gen5377 = Call(__e, ShenFunc(symshen_4pair), gen5372, gen5373)

				} else {
					gen5377 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen5377
				gen5388 := Call(__e, ShenFunc(symfail))

				gen5389 := Call(__e, ShenFunc(sym_a), YaccParse, gen5388)

				if True == gen5389 {
					gen5378 := Call(__e, ShenFunc(symshen_4_5sym_6), V1553)

					Parse__shen_4_5sym_6 := gen5378
					gen5385 := Call(__e, ShenFunc(symfail))

					gen5386 := Call(__e, ShenFunc(sym_a), gen5385, Parse__shen_4_5sym_6)

					gen5387 := Call(__e, ShenFunc(symnot), gen5386)

					if True == gen5387 {
						gen5379 := Call(__e, ShenFunc(symhd), Parse__shen_4_5sym_6)

						gen5382 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5sym_6)

						gen5383 := Call(__e, ShenFunc(sym_a), gen5382, MakeString("<>"))

						var gen5384 Obj
						if True == gen5383 {
							gen5381 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

							gen5384 = Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen5381)

						} else {
							gen5380 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5sym_6)

							gen5384 = Call(__e, ShenFunc(symintern), gen5380)

						}
						__e.TailApply(ShenFunc(symshen_4pair), gen5379, gen5384)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<atom>"), gen5392)

		gen5418 := MakeNative(func(__e Evaluator) {
			V1555 := __e.Get(1)
			_ = V1555
			gen5417 := Call(__e, ShenFunc(sym_a), Nil, V1555)

			if True == gen5417 {
				__e.Return(MakeString(""))
				return
			} else {
				gen5415 := Call(__e, ShenFunc(symcons_2), V1555)

				var gen5416 Obj
				if True == gen5415 {
					gen5412 := Call(__e, ShenFunc(symhd), V1555)

					gen5413 := Call(__e, ShenFunc(sym_a), MakeString("c"), gen5412)

					var gen5414 Obj
					if True == gen5413 {
						gen5409 := Call(__e, ShenFunc(symtl), V1555)

						gen5410 := Call(__e, ShenFunc(symcons_2), gen5409)

						var gen5411 Obj
						if True == gen5410 {
							gen5406 := Call(__e, ShenFunc(symtl), V1555)

							gen5407 := Call(__e, ShenFunc(symhd), gen5406)

							gen5408 := Call(__e, ShenFunc(sym_a), MakeString("#"), gen5407)

							if True == gen5408 {
								gen5411 = True
							} else {
								gen5411 = False
							}

						} else {
							gen5411 = False
						}
						if True == gen5411 {
							gen5414 = True
						} else {
							gen5414 = False
						}

					} else {
						gen5414 = False
					}
					if True == gen5414 {
						gen5416 = True
					} else {
						gen5416 = False
					}

				} else {
					gen5416 = False
				}
				if True == gen5416 {
					gen5397 := Call(__e, ShenFunc(symtl), V1555)

					gen5398 := Call(__e, ShenFunc(symtl), gen5397)

					gen5399 := Call(__e, ShenFunc(symshen_4code_1point), gen5398)

					CodePoint := gen5399
					gen5400 := Call(__e, ShenFunc(symtl), V1555)

					gen5401 := Call(__e, ShenFunc(symtl), gen5400)

					gen5402 := Call(__e, ShenFunc(symshen_4after_1codepoint), gen5401)

					AfterCodePoint := gen5402
					gen5403 := Call(__e, ShenFunc(symshen_4decimalise), CodePoint)

					gen5404 := Call(__e, ShenFunc(symn_1_6string), gen5403)

					gen5405 := Call(__e, ShenFunc(symshen_4control_1chars), AfterCodePoint)

					__e.TailApply(ShenFunc(sym_8s), gen5404, gen5405)

					return

				} else {
					gen5396 := Call(__e, ShenFunc(symcons_2), V1555)

					if True == gen5396 {
						gen5393 := Call(__e, ShenFunc(symhd), V1555)

						gen5394 := Call(__e, ShenFunc(symtl), V1555)

						gen5395 := Call(__e, ShenFunc(symshen_4control_1chars), gen5394)

						__e.TailApply(ShenFunc(sym_8s), gen5393, gen5395)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.control-chars"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.control-chars"), gen5418)

		gen5443 := MakeNative(func(__e Evaluator) {
			V1559 := __e.Get(1)
			_ = V1559
			gen5441 := Call(__e, ShenFunc(symcons_2), V1559)

			var gen5442 Obj
			if True == gen5441 {
				gen5439 := Call(__e, ShenFunc(symhd), V1559)

				gen5440 := Call(__e, ShenFunc(sym_a), MakeString(";"), gen5439)

				if True == gen5440 {
					gen5442 = True
				} else {
					gen5442 = False
				}

			} else {
				gen5442 = False
			}
			if True == gen5442 {
				__e.Return(MakeString(""))
				return
			} else {
				gen5437 := Call(__e, ShenFunc(symcons_2), V1559)

				var gen5438 Obj
				if True == gen5437 {
					gen5424 := Call(__e, ShenFunc(symhd), V1559)

					gen5425 := Call(__e, ShenFunc(symcons), MakeString("0"), Nil)

					gen5426 := Call(__e, ShenFunc(symcons), MakeString("9"), gen5425)

					gen5427 := Call(__e, ShenFunc(symcons), MakeString("8"), gen5426)

					gen5428 := Call(__e, ShenFunc(symcons), MakeString("7"), gen5427)

					gen5429 := Call(__e, ShenFunc(symcons), MakeString("6"), gen5428)

					gen5430 := Call(__e, ShenFunc(symcons), MakeString("5"), gen5429)

					gen5431 := Call(__e, ShenFunc(symcons), MakeString("4"), gen5430)

					gen5432 := Call(__e, ShenFunc(symcons), MakeString("3"), gen5431)

					gen5433 := Call(__e, ShenFunc(symcons), MakeString("2"), gen5432)

					gen5434 := Call(__e, ShenFunc(symcons), MakeString("1"), gen5433)

					gen5435 := Call(__e, ShenFunc(symcons), MakeString("0"), gen5434)

					gen5436 := Call(__e, ShenFunc(symelement_2), gen5424, gen5435)

					if True == gen5436 {
						gen5438 = True
					} else {
						gen5438 = False
					}

				} else {
					gen5438 = False
				}
				if True == gen5438 {
					gen5421 := Call(__e, ShenFunc(symhd), V1559)

					gen5422 := Call(__e, ShenFunc(symtl), V1559)

					gen5423 := Call(__e, ShenFunc(symshen_4code_1point), gen5422)

					__e.TailApply(ShenFunc(symcons), gen5421, gen5423)

					return

				} else {
					gen5419 := Call(__e, ShenFunc(symshen_4app), V1559, MakeString("\n"), MakeSymbol("shen.a"))

					gen5420 := Call(__e, ShenFunc(symcn), MakeString("code point parse error "), gen5419)

					__e.TailApply(ShenFunc(symsimple_1error), gen5420)

					return

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.code-point"), gen5443)

		gen5451 := MakeNative(func(__e Evaluator) {
			V1565 := __e.Get(1)
			_ = V1565
			gen5450 := Call(__e, ShenFunc(sym_a), Nil, V1565)

			if True == gen5450 {
				__e.Return(Nil)
				return
			} else {
				gen5448 := Call(__e, ShenFunc(symcons_2), V1565)

				var gen5449 Obj
				if True == gen5448 {
					gen5446 := Call(__e, ShenFunc(symhd), V1565)

					gen5447 := Call(__e, ShenFunc(sym_a), MakeString(";"), gen5446)

					if True == gen5447 {
						gen5449 = True
					} else {
						gen5449 = False
					}

				} else {
					gen5449 = False
				}
				if True == gen5449 {
					__e.TailApply(ShenFunc(symtl), V1565)

					return
				} else {
					gen5445 := Call(__e, ShenFunc(symcons_2), V1565)

					if True == gen5445 {
						gen5444 := Call(__e, ShenFunc(symtl), V1565)

						__e.TailApply(ShenFunc(symshen_4after_1codepoint), gen5444)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.after-codepoint"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.after-codepoint"), gen5451)

		gen5454 := MakeNative(func(__e Evaluator) {
			V1567 := __e.Get(1)
			_ = V1567
			gen5452 := Call(__e, ShenFunc(symshen_4digits_1_6integers), V1567)

			gen5453 := Call(__e, ShenFunc(symreverse), gen5452)

			__e.TailApply(ShenFunc(symshen_4pre), gen5453, MakeNumber(0))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.decimalise"), gen5454)

		gen5515 := MakeNative(func(__e Evaluator) {
			V1573 := __e.Get(1)
			_ = V1573
			gen5513 := Call(__e, ShenFunc(symcons_2), V1573)

			var gen5514 Obj
			if True == gen5513 {
				gen5511 := Call(__e, ShenFunc(symhd), V1573)

				gen5512 := Call(__e, ShenFunc(sym_a), MakeString("0"), gen5511)

				if True == gen5512 {
					gen5514 = True
				} else {
					gen5514 = False
				}

			} else {
				gen5514 = False
			}
			if True == gen5514 {
				gen5509 := Call(__e, ShenFunc(symtl), V1573)

				gen5510 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5509)

				__e.TailApply(ShenFunc(symcons), MakeNumber(0), gen5510)

				return

			} else {
				gen5507 := Call(__e, ShenFunc(symcons_2), V1573)

				var gen5508 Obj
				if True == gen5507 {
					gen5505 := Call(__e, ShenFunc(symhd), V1573)

					gen5506 := Call(__e, ShenFunc(sym_a), MakeString("1"), gen5505)

					if True == gen5506 {
						gen5508 = True
					} else {
						gen5508 = False
					}

				} else {
					gen5508 = False
				}
				if True == gen5508 {
					gen5503 := Call(__e, ShenFunc(symtl), V1573)

					gen5504 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5503)

					__e.TailApply(ShenFunc(symcons), MakeNumber(1), gen5504)

					return

				} else {
					gen5501 := Call(__e, ShenFunc(symcons_2), V1573)

					var gen5502 Obj
					if True == gen5501 {
						gen5499 := Call(__e, ShenFunc(symhd), V1573)

						gen5500 := Call(__e, ShenFunc(sym_a), MakeString("2"), gen5499)

						if True == gen5500 {
							gen5502 = True
						} else {
							gen5502 = False
						}

					} else {
						gen5502 = False
					}
					if True == gen5502 {
						gen5497 := Call(__e, ShenFunc(symtl), V1573)

						gen5498 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5497)

						__e.TailApply(ShenFunc(symcons), MakeNumber(2), gen5498)

						return

					} else {
						gen5495 := Call(__e, ShenFunc(symcons_2), V1573)

						var gen5496 Obj
						if True == gen5495 {
							gen5493 := Call(__e, ShenFunc(symhd), V1573)

							gen5494 := Call(__e, ShenFunc(sym_a), MakeString("3"), gen5493)

							if True == gen5494 {
								gen5496 = True
							} else {
								gen5496 = False
							}

						} else {
							gen5496 = False
						}
						if True == gen5496 {
							gen5491 := Call(__e, ShenFunc(symtl), V1573)

							gen5492 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5491)

							__e.TailApply(ShenFunc(symcons), MakeNumber(3), gen5492)

							return

						} else {
							gen5489 := Call(__e, ShenFunc(symcons_2), V1573)

							var gen5490 Obj
							if True == gen5489 {
								gen5487 := Call(__e, ShenFunc(symhd), V1573)

								gen5488 := Call(__e, ShenFunc(sym_a), MakeString("4"), gen5487)

								if True == gen5488 {
									gen5490 = True
								} else {
									gen5490 = False
								}

							} else {
								gen5490 = False
							}
							if True == gen5490 {
								gen5485 := Call(__e, ShenFunc(symtl), V1573)

								gen5486 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5485)

								__e.TailApply(ShenFunc(symcons), MakeNumber(4), gen5486)

								return

							} else {
								gen5483 := Call(__e, ShenFunc(symcons_2), V1573)

								var gen5484 Obj
								if True == gen5483 {
									gen5481 := Call(__e, ShenFunc(symhd), V1573)

									gen5482 := Call(__e, ShenFunc(sym_a), MakeString("5"), gen5481)

									if True == gen5482 {
										gen5484 = True
									} else {
										gen5484 = False
									}

								} else {
									gen5484 = False
								}
								if True == gen5484 {
									gen5479 := Call(__e, ShenFunc(symtl), V1573)

									gen5480 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5479)

									__e.TailApply(ShenFunc(symcons), MakeNumber(5), gen5480)

									return

								} else {
									gen5477 := Call(__e, ShenFunc(symcons_2), V1573)

									var gen5478 Obj
									if True == gen5477 {
										gen5475 := Call(__e, ShenFunc(symhd), V1573)

										gen5476 := Call(__e, ShenFunc(sym_a), MakeString("6"), gen5475)

										if True == gen5476 {
											gen5478 = True
										} else {
											gen5478 = False
										}

									} else {
										gen5478 = False
									}
									if True == gen5478 {
										gen5473 := Call(__e, ShenFunc(symtl), V1573)

										gen5474 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5473)

										__e.TailApply(ShenFunc(symcons), MakeNumber(6), gen5474)

										return

									} else {
										gen5471 := Call(__e, ShenFunc(symcons_2), V1573)

										var gen5472 Obj
										if True == gen5471 {
											gen5469 := Call(__e, ShenFunc(symhd), V1573)

											gen5470 := Call(__e, ShenFunc(sym_a), MakeString("7"), gen5469)

											if True == gen5470 {
												gen5472 = True
											} else {
												gen5472 = False
											}

										} else {
											gen5472 = False
										}
										if True == gen5472 {
											gen5467 := Call(__e, ShenFunc(symtl), V1573)

											gen5468 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5467)

											__e.TailApply(ShenFunc(symcons), MakeNumber(7), gen5468)

											return

										} else {
											gen5465 := Call(__e, ShenFunc(symcons_2), V1573)

											var gen5466 Obj
											if True == gen5465 {
												gen5463 := Call(__e, ShenFunc(symhd), V1573)

												gen5464 := Call(__e, ShenFunc(sym_a), MakeString("8"), gen5463)

												if True == gen5464 {
													gen5466 = True
												} else {
													gen5466 = False
												}

											} else {
												gen5466 = False
											}
											if True == gen5466 {
												gen5461 := Call(__e, ShenFunc(symtl), V1573)

												gen5462 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5461)

												__e.TailApply(ShenFunc(symcons), MakeNumber(8), gen5462)

												return

											} else {
												gen5459 := Call(__e, ShenFunc(symcons_2), V1573)

												var gen5460 Obj
												if True == gen5459 {
													gen5457 := Call(__e, ShenFunc(symhd), V1573)

													gen5458 := Call(__e, ShenFunc(sym_a), MakeString("9"), gen5457)

													if True == gen5458 {
														gen5460 = True
													} else {
														gen5460 = False
													}

												} else {
													gen5460 = False
												}
												if True == gen5460 {
													gen5455 := Call(__e, ShenFunc(symtl), V1573)

													gen5456 := Call(__e, ShenFunc(symshen_4digits_1_6integers), gen5455)

													__e.TailApply(ShenFunc(symcons), MakeNumber(9), gen5456)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.digits->integers"), gen5515)

		gen5528 := MakeNative(func(__e Evaluator) {
			V1575 := __e.Get(1)
			_ = V1575
			gen5516 := Call(__e, ShenFunc(symshen_4_5alpha_6), V1575)

			Parse__shen_4_5alpha_6 := gen5516
			gen5525 := Call(__e, ShenFunc(symfail))

			gen5526 := Call(__e, ShenFunc(sym_a), gen5525, Parse__shen_4_5alpha_6)

			gen5527 := Call(__e, ShenFunc(symnot), gen5526)

			if True == gen5527 {
				gen5517 := Call(__e, ShenFunc(symshen_4_5alphanums_6), Parse__shen_4_5alpha_6)

				Parse__shen_4_5alphanums_6 := gen5517
				gen5522 := Call(__e, ShenFunc(symfail))

				gen5523 := Call(__e, ShenFunc(sym_a), gen5522, Parse__shen_4_5alphanums_6)

				gen5524 := Call(__e, ShenFunc(symnot), gen5523)

				if True == gen5524 {
					gen5518 := Call(__e, ShenFunc(symhd), Parse__shen_4_5alphanums_6)

					gen5519 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alpha_6)

					gen5520 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alphanums_6)

					gen5521 := Call(__e, ShenFunc(sym_8s), gen5519, gen5520)

					__e.TailApply(ShenFunc(symshen_4pair), gen5518, gen5521)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<sym>"), gen5528)

		gen5549 := MakeNative(func(__e Evaluator) {
			V1577 := __e.Get(1)
			_ = V1577
			gen5529 := Call(__e, ShenFunc(symshen_4_5alphanum_6), V1577)

			Parse__shen_4_5alphanum_6 := gen5529
			gen5538 := Call(__e, ShenFunc(symfail))

			gen5539 := Call(__e, ShenFunc(sym_a), gen5538, Parse__shen_4_5alphanum_6)

			gen5540 := Call(__e, ShenFunc(symnot), gen5539)

			var gen5541 Obj
			if True == gen5540 {
				gen5530 := Call(__e, ShenFunc(symshen_4_5alphanums_6), Parse__shen_4_5alphanum_6)

				Parse__shen_4_5alphanums_6 := gen5530
				gen5535 := Call(__e, ShenFunc(symfail))

				gen5536 := Call(__e, ShenFunc(sym_a), gen5535, Parse__shen_4_5alphanums_6)

				gen5537 := Call(__e, ShenFunc(symnot), gen5536)

				if True == gen5537 {
					gen5531 := Call(__e, ShenFunc(symhd), Parse__shen_4_5alphanums_6)

					gen5532 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alphanum_6)

					gen5533 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alphanums_6)

					gen5534 := Call(__e, ShenFunc(sym_8s), gen5532, gen5533)

					gen5541 = Call(__e, ShenFunc(symshen_4pair), gen5531, gen5534)

				} else {
					gen5541 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen5541 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5541
			gen5547 := Call(__e, ShenFunc(symfail))

			gen5548 := Call(__e, ShenFunc(sym_a), YaccParse, gen5547)

			if True == gen5548 {
				gen5542 := Call(__e, ShenFunc(sym_5e_6), V1577)

				Parse___5e_6 := gen5542
				gen5544 := Call(__e, ShenFunc(symfail))

				gen5545 := Call(__e, ShenFunc(sym_a), gen5544, Parse___5e_6)

				gen5546 := Call(__e, ShenFunc(symnot), gen5545)

				if True == gen5546 {
					gen5543 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen5543, MakeString(""))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<alphanums>"), gen5549)

		gen5565 := MakeNative(func(__e Evaluator) {
			V1579 := __e.Get(1)
			_ = V1579
			gen5550 := Call(__e, ShenFunc(symshen_4_5alpha_6), V1579)

			Parse__shen_4_5alpha_6 := gen5550
			gen5553 := Call(__e, ShenFunc(symfail))

			gen5554 := Call(__e, ShenFunc(sym_a), gen5553, Parse__shen_4_5alpha_6)

			gen5555 := Call(__e, ShenFunc(symnot), gen5554)

			var gen5556 Obj
			if True == gen5555 {
				gen5551 := Call(__e, ShenFunc(symhd), Parse__shen_4_5alpha_6)

				gen5552 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5alpha_6)

				gen5556 = Call(__e, ShenFunc(symshen_4pair), gen5551, gen5552)

			} else {
				gen5556 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5556
			gen5563 := Call(__e, ShenFunc(symfail))

			gen5564 := Call(__e, ShenFunc(sym_a), YaccParse, gen5563)

			if True == gen5564 {
				gen5557 := Call(__e, ShenFunc(symshen_4_5num_6), V1579)

				Parse__shen_4_5num_6 := gen5557
				gen5560 := Call(__e, ShenFunc(symfail))

				gen5561 := Call(__e, ShenFunc(sym_a), gen5560, Parse__shen_4_5num_6)

				gen5562 := Call(__e, ShenFunc(symnot), gen5561)

				if True == gen5562 {
					gen5558 := Call(__e, ShenFunc(symhd), Parse__shen_4_5num_6)

					gen5559 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5num_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen5558, gen5559)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<alphanum>"), gen5565)

		gen5575 := MakeNative(func(__e Evaluator) {
			V1581 := __e.Get(1)
			_ = V1581
			gen5573 := Call(__e, ShenFunc(symhd), V1581)

			gen5574 := Call(__e, ShenFunc(symcons_2), gen5573)

			if True == gen5574 {
				gen5566 := Call(__e, ShenFunc(symshen_4hdhd), V1581)

				Parse__Char := gen5566
				gen5572 := Call(__e, ShenFunc(symshen_4numbyte_2), Parse__Char)

				if True == gen5572 {
					gen5567 := Call(__e, ShenFunc(symshen_4tlhd), V1581)

					gen5568 := Call(__e, ShenFunc(symshen_4hdtl), V1581)

					gen5569 := Call(__e, ShenFunc(symshen_4pair), gen5567, gen5568)

					gen5570 := Call(__e, ShenFunc(symhd), gen5569)

					gen5571 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

					__e.TailApply(ShenFunc(symshen_4pair), gen5570, gen5571)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<num>"), gen5575)

		gen5586 := MakeNative(func(__e Evaluator) {
			V1587 := __e.Get(1)
			_ = V1587
			gen5585 := Call(__e, ShenFunc(sym_a), MakeNumber(48), V1587)

			if True == gen5585 {
				__e.Return(True)
				return
			} else {
				gen5584 := Call(__e, ShenFunc(sym_a), MakeNumber(49), V1587)

				if True == gen5584 {
					__e.Return(True)
					return
				} else {
					gen5583 := Call(__e, ShenFunc(sym_a), MakeNumber(50), V1587)

					if True == gen5583 {
						__e.Return(True)
						return
					} else {
						gen5582 := Call(__e, ShenFunc(sym_a), MakeNumber(51), V1587)

						if True == gen5582 {
							__e.Return(True)
							return
						} else {
							gen5581 := Call(__e, ShenFunc(sym_a), MakeNumber(52), V1587)

							if True == gen5581 {
								__e.Return(True)
								return
							} else {
								gen5580 := Call(__e, ShenFunc(sym_a), MakeNumber(53), V1587)

								if True == gen5580 {
									__e.Return(True)
									return
								} else {
									gen5579 := Call(__e, ShenFunc(sym_a), MakeNumber(54), V1587)

									if True == gen5579 {
										__e.Return(True)
										return
									} else {
										gen5578 := Call(__e, ShenFunc(sym_a), MakeNumber(55), V1587)

										if True == gen5578 {
											__e.Return(True)
											return
										} else {
											gen5577 := Call(__e, ShenFunc(sym_a), MakeNumber(56), V1587)

											if True == gen5577 {
												__e.Return(True)
												return
											} else {
												gen5576 := Call(__e, ShenFunc(sym_a), MakeNumber(57), V1587)

												if True == gen5576 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.numbyte?"), gen5586)

		gen5596 := MakeNative(func(__e Evaluator) {
			V1589 := __e.Get(1)
			_ = V1589
			gen5594 := Call(__e, ShenFunc(symhd), V1589)

			gen5595 := Call(__e, ShenFunc(symcons_2), gen5594)

			if True == gen5595 {
				gen5587 := Call(__e, ShenFunc(symshen_4hdhd), V1589)

				Parse__Char := gen5587
				gen5593 := Call(__e, ShenFunc(symshen_4symbol_1code_2), Parse__Char)

				if True == gen5593 {
					gen5588 := Call(__e, ShenFunc(symshen_4tlhd), V1589)

					gen5589 := Call(__e, ShenFunc(symshen_4hdtl), V1589)

					gen5590 := Call(__e, ShenFunc(symshen_4pair), gen5588, gen5589)

					gen5591 := Call(__e, ShenFunc(symhd), gen5590)

					gen5592 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

					__e.TailApply(ShenFunc(symshen_4pair), gen5591, gen5592)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<alpha>"), gen5596)

		gen5618 := MakeNative(func(__e Evaluator) {
			V1591 := __e.Get(1)
			_ = V1591
			gen5617 := Call(__e, ShenFunc(sym_a), V1591, MakeNumber(126))

			if True == gen5617 {
				__e.Return(True)
				return
			} else {
				gen5614 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(94))

				var gen5615 Obj
				if True == gen5614 {
					gen5613 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(123))

					if True == gen5613 {
						gen5615 = True
					} else {
						gen5615 = False
					}

				} else {
					gen5615 = False
				}
				var gen5616 Obj
				if True == gen5615 {
					gen5616 = True
				} else {
					gen5610 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(59))

					var gen5611 Obj
					if True == gen5610 {
						gen5609 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(91))

						if True == gen5609 {
							gen5611 = True
						} else {
							gen5611 = False
						}

					} else {
						gen5611 = False
					}
					var gen5612 Obj
					if True == gen5611 {
						gen5612 = True
					} else {
						gen5606 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(41))

						var gen5607 Obj
						if True == gen5606 {
							gen5604 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(58))

							var gen5605 Obj
							if True == gen5604 {
								gen5602 := Call(__e, ShenFunc(sym_a), V1591, MakeNumber(44))

								gen5603 := Call(__e, ShenFunc(symnot), gen5602)

								if True == gen5603 {
									gen5605 = True
								} else {
									gen5605 = False
								}

							} else {
								gen5605 = False
							}
							if True == gen5605 {
								gen5607 = True
							} else {
								gen5607 = False
							}

						} else {
							gen5607 = False
						}
						var gen5608 Obj
						if True == gen5607 {
							gen5608 = True
						} else {
							gen5599 := Call(__e, ShenFunc(sym_6), V1591, MakeNumber(34))

							var gen5600 Obj
							if True == gen5599 {
								gen5598 := Call(__e, ShenFunc(sym_5), V1591, MakeNumber(40))

								if True == gen5598 {
									gen5600 = True
								} else {
									gen5600 = False
								}

							} else {
								gen5600 = False
							}
							var gen5601 Obj
							if True == gen5600 {
								gen5601 = True
							} else {
								gen5597 := Call(__e, ShenFunc(sym_a), V1591, MakeNumber(33))

								if True == gen5597 {
									gen5601 = True
								} else {
									gen5601 = False
								}

							}
							if True == gen5601 {
								gen5608 = True
							} else {
								gen5608 = False
							}

						}
						if True == gen5608 {
							gen5612 = True
						} else {
							gen5612 = False
						}

					}
					if True == gen5612 {
						gen5616 = True
					} else {
						gen5616 = False
					}

				}
				if True == gen5616 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.symbol-code?"), gen5618)

		gen5633 := MakeNative(func(__e Evaluator) {
			V1593 := __e.Get(1)
			_ = V1593
			gen5619 := Call(__e, ShenFunc(symshen_4_5dbq_6), V1593)

			Parse__shen_4_5dbq_6 := gen5619
			gen5630 := Call(__e, ShenFunc(symfail))

			gen5631 := Call(__e, ShenFunc(sym_a), gen5630, Parse__shen_4_5dbq_6)

			gen5632 := Call(__e, ShenFunc(symnot), gen5631)

			if True == gen5632 {
				gen5620 := Call(__e, ShenFunc(symshen_4_5strcontents_6), Parse__shen_4_5dbq_6)

				Parse__shen_4_5strcontents_6 := gen5620
				gen5627 := Call(__e, ShenFunc(symfail))

				gen5628 := Call(__e, ShenFunc(sym_a), gen5627, Parse__shen_4_5strcontents_6)

				gen5629 := Call(__e, ShenFunc(symnot), gen5628)

				if True == gen5629 {
					gen5621 := Call(__e, ShenFunc(symshen_4_5dbq_6), Parse__shen_4_5strcontents_6)

					Parse__shen_4_5dbq_6 := gen5621
					gen5624 := Call(__e, ShenFunc(symfail))

					gen5625 := Call(__e, ShenFunc(sym_a), gen5624, Parse__shen_4_5dbq_6)

					gen5626 := Call(__e, ShenFunc(symnot), gen5625)

					if True == gen5626 {
						gen5622 := Call(__e, ShenFunc(symhd), Parse__shen_4_5dbq_6)

						gen5623 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5strcontents_6)

						__e.TailApply(ShenFunc(symshen_4pair), gen5622, gen5623)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<str>"), gen5633)

		gen5642 := MakeNative(func(__e Evaluator) {
			V1595 := __e.Get(1)
			_ = V1595
			gen5640 := Call(__e, ShenFunc(symhd), V1595)

			gen5641 := Call(__e, ShenFunc(symcons_2), gen5640)

			if True == gen5641 {
				gen5634 := Call(__e, ShenFunc(symshen_4hdhd), V1595)

				Parse__Char := gen5634
				gen5639 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(34))

				if True == gen5639 {
					gen5635 := Call(__e, ShenFunc(symshen_4tlhd), V1595)

					gen5636 := Call(__e, ShenFunc(symshen_4hdtl), V1595)

					gen5637 := Call(__e, ShenFunc(symshen_4pair), gen5635, gen5636)

					gen5638 := Call(__e, ShenFunc(symhd), gen5637)

					__e.TailApply(ShenFunc(symshen_4pair), gen5638, Parse__Char)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<dbq>"), gen5642)

		gen5663 := MakeNative(func(__e Evaluator) {
			V1597 := __e.Get(1)
			_ = V1597
			gen5643 := Call(__e, ShenFunc(symshen_4_5strc_6), V1597)

			Parse__shen_4_5strc_6 := gen5643
			gen5652 := Call(__e, ShenFunc(symfail))

			gen5653 := Call(__e, ShenFunc(sym_a), gen5652, Parse__shen_4_5strc_6)

			gen5654 := Call(__e, ShenFunc(symnot), gen5653)

			var gen5655 Obj
			if True == gen5654 {
				gen5644 := Call(__e, ShenFunc(symshen_4_5strcontents_6), Parse__shen_4_5strc_6)

				Parse__shen_4_5strcontents_6 := gen5644
				gen5649 := Call(__e, ShenFunc(symfail))

				gen5650 := Call(__e, ShenFunc(sym_a), gen5649, Parse__shen_4_5strcontents_6)

				gen5651 := Call(__e, ShenFunc(symnot), gen5650)

				if True == gen5651 {
					gen5645 := Call(__e, ShenFunc(symhd), Parse__shen_4_5strcontents_6)

					gen5646 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5strc_6)

					gen5647 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5strcontents_6)

					gen5648 := Call(__e, ShenFunc(symcons), gen5646, gen5647)

					gen5655 = Call(__e, ShenFunc(symshen_4pair), gen5645, gen5648)

				} else {
					gen5655 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen5655 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5655
			gen5661 := Call(__e, ShenFunc(symfail))

			gen5662 := Call(__e, ShenFunc(sym_a), YaccParse, gen5661)

			if True == gen5662 {
				gen5656 := Call(__e, ShenFunc(sym_5e_6), V1597)

				Parse___5e_6 := gen5656
				gen5658 := Call(__e, ShenFunc(symfail))

				gen5659 := Call(__e, ShenFunc(sym_a), gen5658, Parse___5e_6)

				gen5660 := Call(__e, ShenFunc(symnot), gen5659)

				if True == gen5660 {
					gen5657 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen5657, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<strcontents>"), gen5663)

		gen5672 := MakeNative(func(__e Evaluator) {
			V1599 := __e.Get(1)
			_ = V1599
			gen5670 := Call(__e, ShenFunc(symhd), V1599)

			gen5671 := Call(__e, ShenFunc(symcons_2), gen5670)

			if True == gen5671 {
				gen5664 := Call(__e, ShenFunc(symshen_4hdhd), V1599)

				Parse__Char := gen5664
				gen5665 := Call(__e, ShenFunc(symshen_4tlhd), V1599)

				gen5666 := Call(__e, ShenFunc(symshen_4hdtl), V1599)

				gen5667 := Call(__e, ShenFunc(symshen_4pair), gen5665, gen5666)

				gen5668 := Call(__e, ShenFunc(symhd), gen5667)

				gen5669 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

				__e.TailApply(ShenFunc(symshen_4pair), gen5668, gen5669)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<byte>"), gen5672)

		gen5683 := MakeNative(func(__e Evaluator) {
			V1601 := __e.Get(1)
			_ = V1601
			gen5681 := Call(__e, ShenFunc(symhd), V1601)

			gen5682 := Call(__e, ShenFunc(symcons_2), gen5681)

			if True == gen5682 {
				gen5673 := Call(__e, ShenFunc(symshen_4hdhd), V1601)

				Parse__Char := gen5673
				gen5679 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(34))

				gen5680 := Call(__e, ShenFunc(symnot), gen5679)

				if True == gen5680 {
					gen5674 := Call(__e, ShenFunc(symshen_4tlhd), V1601)

					gen5675 := Call(__e, ShenFunc(symshen_4hdtl), V1601)

					gen5676 := Call(__e, ShenFunc(symshen_4pair), gen5674, gen5675)

					gen5677 := Call(__e, ShenFunc(symhd), gen5676)

					gen5678 := Call(__e, ShenFunc(symn_1_6string), Parse__Char)

					__e.TailApply(ShenFunc(symshen_4pair), gen5677, gen5678)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<strc>"), gen5683)

		gen5796 := MakeNative(func(__e Evaluator) {
			V1603 := __e.Get(1)
			_ = V1603
			gen5684 := Call(__e, ShenFunc(symshen_4_5minus_6), V1603)

			Parse__shen_4_5minus_6 := gen5684
			gen5692 := Call(__e, ShenFunc(symfail))

			gen5693 := Call(__e, ShenFunc(sym_a), gen5692, Parse__shen_4_5minus_6)

			gen5694 := Call(__e, ShenFunc(symnot), gen5693)

			var gen5695 Obj
			if True == gen5694 {
				gen5685 := Call(__e, ShenFunc(symshen_4_5number_6), Parse__shen_4_5minus_6)

				Parse__shen_4_5number_6 := gen5685
				gen5689 := Call(__e, ShenFunc(symfail))

				gen5690 := Call(__e, ShenFunc(sym_a), gen5689, Parse__shen_4_5number_6)

				gen5691 := Call(__e, ShenFunc(symnot), gen5690)

				if True == gen5691 {
					gen5686 := Call(__e, ShenFunc(symhd), Parse__shen_4_5number_6)

					gen5687 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5number_6)

					gen5688 := Call(__e, ShenFunc(sym_1), MakeNumber(0), gen5687)

					gen5695 = Call(__e, ShenFunc(symshen_4pair), gen5686, gen5688)

				} else {
					gen5695 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen5695 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5695
			gen5794 := Call(__e, ShenFunc(symfail))

			gen5795 := Call(__e, ShenFunc(sym_a), YaccParse, gen5794)

			if True == gen5795 {
				gen5696 := Call(__e, ShenFunc(symshen_4_5plus_6), V1603)

				Parse__shen_4_5plus_6 := gen5696
				gen5703 := Call(__e, ShenFunc(symfail))

				gen5704 := Call(__e, ShenFunc(sym_a), gen5703, Parse__shen_4_5plus_6)

				gen5705 := Call(__e, ShenFunc(symnot), gen5704)

				var gen5706 Obj
				if True == gen5705 {
					gen5697 := Call(__e, ShenFunc(symshen_4_5number_6), Parse__shen_4_5plus_6)

					Parse__shen_4_5number_6 := gen5697
					gen5700 := Call(__e, ShenFunc(symfail))

					gen5701 := Call(__e, ShenFunc(sym_a), gen5700, Parse__shen_4_5number_6)

					gen5702 := Call(__e, ShenFunc(symnot), gen5701)

					if True == gen5702 {
						gen5698 := Call(__e, ShenFunc(symhd), Parse__shen_4_5number_6)

						gen5699 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5number_6)

						gen5706 = Call(__e, ShenFunc(symshen_4pair), gen5698, gen5699)

					} else {
						gen5706 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen5706 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen5706
				gen5792 := Call(__e, ShenFunc(symfail))

				gen5793 := Call(__e, ShenFunc(sym_a), YaccParse, gen5792)

				if True == gen5793 {
					gen5707 := Call(__e, ShenFunc(symshen_4_5predigits_6), V1603)

					Parse__shen_4_5predigits_6 := gen5707
					gen5734 := Call(__e, ShenFunc(symfail))

					gen5735 := Call(__e, ShenFunc(sym_a), gen5734, Parse__shen_4_5predigits_6)

					gen5736 := Call(__e, ShenFunc(symnot), gen5735)

					var gen5737 Obj
					if True == gen5736 {
						gen5708 := Call(__e, ShenFunc(symshen_4_5stop_6), Parse__shen_4_5predigits_6)

						Parse__shen_4_5stop_6 := gen5708
						gen5731 := Call(__e, ShenFunc(symfail))

						gen5732 := Call(__e, ShenFunc(sym_a), gen5731, Parse__shen_4_5stop_6)

						gen5733 := Call(__e, ShenFunc(symnot), gen5732)

						if True == gen5733 {
							gen5709 := Call(__e, ShenFunc(symshen_4_5postdigits_6), Parse__shen_4_5stop_6)

							Parse__shen_4_5postdigits_6 := gen5709
							gen5728 := Call(__e, ShenFunc(symfail))

							gen5729 := Call(__e, ShenFunc(sym_a), gen5728, Parse__shen_4_5postdigits_6)

							gen5730 := Call(__e, ShenFunc(symnot), gen5729)

							if True == gen5730 {
								gen5710 := Call(__e, ShenFunc(symshen_4_5E_6), Parse__shen_4_5postdigits_6)

								Parse__shen_4_5E_6 := gen5710
								gen5725 := Call(__e, ShenFunc(symfail))

								gen5726 := Call(__e, ShenFunc(sym_a), gen5725, Parse__shen_4_5E_6)

								gen5727 := Call(__e, ShenFunc(symnot), gen5726)

								if True == gen5727 {
									gen5711 := Call(__e, ShenFunc(symshen_4_5log10_6), Parse__shen_4_5E_6)

									Parse__shen_4_5log10_6 := gen5711
									gen5722 := Call(__e, ShenFunc(symfail))

									gen5723 := Call(__e, ShenFunc(sym_a), gen5722, Parse__shen_4_5log10_6)

									gen5724 := Call(__e, ShenFunc(symnot), gen5723)

									if True == gen5724 {
										gen5712 := Call(__e, ShenFunc(symhd), Parse__shen_4_5log10_6)

										gen5713 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5log10_6)

										gen5714 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), gen5713)

										gen5715 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5predigits_6)

										gen5716 := Call(__e, ShenFunc(symreverse), gen5715)

										gen5717 := Call(__e, ShenFunc(symshen_4pre), gen5716, MakeNumber(0))

										gen5718 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5postdigits_6)

										gen5719 := Call(__e, ShenFunc(symshen_4post), gen5718, MakeNumber(1))

										gen5720 := Call(__e, ShenFunc(sym_7), gen5717, gen5719)

										gen5721 := Call(__e, ShenFunc(sym_d), gen5714, gen5720)

										gen5737 = Call(__e, ShenFunc(symshen_4pair), gen5712, gen5721)

									} else {
										gen5737 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen5737 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen5737 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen5737 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen5737 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen5737
					gen5790 := Call(__e, ShenFunc(symfail))

					gen5791 := Call(__e, ShenFunc(sym_a), YaccParse, gen5790)

					if True == gen5791 {
						gen5738 := Call(__e, ShenFunc(symshen_4_5digits_6), V1603)

						Parse__shen_4_5digits_6 := gen5738
						gen5754 := Call(__e, ShenFunc(symfail))

						gen5755 := Call(__e, ShenFunc(sym_a), gen5754, Parse__shen_4_5digits_6)

						gen5756 := Call(__e, ShenFunc(symnot), gen5755)

						var gen5757 Obj
						if True == gen5756 {
							gen5739 := Call(__e, ShenFunc(symshen_4_5E_6), Parse__shen_4_5digits_6)

							Parse__shen_4_5E_6 := gen5739
							gen5751 := Call(__e, ShenFunc(symfail))

							gen5752 := Call(__e, ShenFunc(sym_a), gen5751, Parse__shen_4_5E_6)

							gen5753 := Call(__e, ShenFunc(symnot), gen5752)

							if True == gen5753 {
								gen5740 := Call(__e, ShenFunc(symshen_4_5log10_6), Parse__shen_4_5E_6)

								Parse__shen_4_5log10_6 := gen5740
								gen5748 := Call(__e, ShenFunc(symfail))

								gen5749 := Call(__e, ShenFunc(sym_a), gen5748, Parse__shen_4_5log10_6)

								gen5750 := Call(__e, ShenFunc(symnot), gen5749)

								if True == gen5750 {
									gen5741 := Call(__e, ShenFunc(symhd), Parse__shen_4_5log10_6)

									gen5742 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5log10_6)

									gen5743 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), gen5742)

									gen5744 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

									gen5745 := Call(__e, ShenFunc(symreverse), gen5744)

									gen5746 := Call(__e, ShenFunc(symshen_4pre), gen5745, MakeNumber(0))

									gen5747 := Call(__e, ShenFunc(sym_d), gen5743, gen5746)

									gen5757 = Call(__e, ShenFunc(symshen_4pair), gen5741, gen5747)

								} else {
									gen5757 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen5757 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen5757 = Call(__e, ShenFunc(symfail))

						}
						YaccParse := gen5757
						gen5788 := Call(__e, ShenFunc(symfail))

						gen5789 := Call(__e, ShenFunc(sym_a), YaccParse, gen5788)

						if True == gen5789 {
							gen5758 := Call(__e, ShenFunc(symshen_4_5predigits_6), V1603)

							Parse__shen_4_5predigits_6 := gen5758
							gen5774 := Call(__e, ShenFunc(symfail))

							gen5775 := Call(__e, ShenFunc(sym_a), gen5774, Parse__shen_4_5predigits_6)

							gen5776 := Call(__e, ShenFunc(symnot), gen5775)

							var gen5777 Obj
							if True == gen5776 {
								gen5759 := Call(__e, ShenFunc(symshen_4_5stop_6), Parse__shen_4_5predigits_6)

								Parse__shen_4_5stop_6 := gen5759
								gen5771 := Call(__e, ShenFunc(symfail))

								gen5772 := Call(__e, ShenFunc(sym_a), gen5771, Parse__shen_4_5stop_6)

								gen5773 := Call(__e, ShenFunc(symnot), gen5772)

								if True == gen5773 {
									gen5760 := Call(__e, ShenFunc(symshen_4_5postdigits_6), Parse__shen_4_5stop_6)

									Parse__shen_4_5postdigits_6 := gen5760
									gen5768 := Call(__e, ShenFunc(symfail))

									gen5769 := Call(__e, ShenFunc(sym_a), gen5768, Parse__shen_4_5postdigits_6)

									gen5770 := Call(__e, ShenFunc(symnot), gen5769)

									if True == gen5770 {
										gen5761 := Call(__e, ShenFunc(symhd), Parse__shen_4_5postdigits_6)

										gen5762 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5predigits_6)

										gen5763 := Call(__e, ShenFunc(symreverse), gen5762)

										gen5764 := Call(__e, ShenFunc(symshen_4pre), gen5763, MakeNumber(0))

										gen5765 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5postdigits_6)

										gen5766 := Call(__e, ShenFunc(symshen_4post), gen5765, MakeNumber(1))

										gen5767 := Call(__e, ShenFunc(sym_7), gen5764, gen5766)

										gen5777 = Call(__e, ShenFunc(symshen_4pair), gen5761, gen5767)

									} else {
										gen5777 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen5777 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen5777 = Call(__e, ShenFunc(symfail))

							}
							YaccParse := gen5777
							gen5786 := Call(__e, ShenFunc(symfail))

							gen5787 := Call(__e, ShenFunc(sym_a), YaccParse, gen5786)

							if True == gen5787 {
								gen5778 := Call(__e, ShenFunc(symshen_4_5digits_6), V1603)

								Parse__shen_4_5digits_6 := gen5778
								gen5783 := Call(__e, ShenFunc(symfail))

								gen5784 := Call(__e, ShenFunc(sym_a), gen5783, Parse__shen_4_5digits_6)

								gen5785 := Call(__e, ShenFunc(symnot), gen5784)

								if True == gen5785 {
									gen5779 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

									gen5780 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

									gen5781 := Call(__e, ShenFunc(symreverse), gen5780)

									gen5782 := Call(__e, ShenFunc(symshen_4pre), gen5781, MakeNumber(0))

									__e.TailApply(ShenFunc(symshen_4pair), gen5779, gen5782)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<number>"), gen5796)

		gen5806 := MakeNative(func(__e Evaluator) {
			V1606 := __e.Get(1)
			_ = V1606
			gen5803 := Call(__e, ShenFunc(symhd), V1606)

			gen5804 := Call(__e, ShenFunc(symcons_2), gen5803)

			var gen5805 Obj
			if True == gen5804 {
				gen5801 := Call(__e, ShenFunc(symshen_4hdhd), V1606)

				gen5802 := Call(__e, ShenFunc(sym_a), MakeNumber(101), gen5801)

				if True == gen5802 {
					gen5805 = True
				} else {
					gen5805 = False
				}

			} else {
				gen5805 = False
			}
			if True == gen5805 {
				gen5797 := Call(__e, ShenFunc(symshen_4tlhd), V1606)

				gen5798 := Call(__e, ShenFunc(symshen_4hdtl), V1606)

				gen5799 := Call(__e, ShenFunc(symshen_4pair), gen5797, gen5798)

				NewStream1604 := gen5799
				gen5800 := Call(__e, ShenFunc(symhd), NewStream1604)

				__e.TailApply(ShenFunc(symshen_4pair), gen5800, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<E>"), gen5806)

		gen5831 := MakeNative(func(__e Evaluator) {
			V1608 := __e.Get(1)
			_ = V1608
			gen5807 := Call(__e, ShenFunc(symshen_4_5minus_6), V1608)

			Parse__shen_4_5minus_6 := gen5807
			gen5817 := Call(__e, ShenFunc(symfail))

			gen5818 := Call(__e, ShenFunc(sym_a), gen5817, Parse__shen_4_5minus_6)

			gen5819 := Call(__e, ShenFunc(symnot), gen5818)

			var gen5820 Obj
			if True == gen5819 {
				gen5808 := Call(__e, ShenFunc(symshen_4_5digits_6), Parse__shen_4_5minus_6)

				Parse__shen_4_5digits_6 := gen5808
				gen5814 := Call(__e, ShenFunc(symfail))

				gen5815 := Call(__e, ShenFunc(sym_a), gen5814, Parse__shen_4_5digits_6)

				gen5816 := Call(__e, ShenFunc(symnot), gen5815)

				if True == gen5816 {
					gen5809 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

					gen5810 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

					gen5811 := Call(__e, ShenFunc(symreverse), gen5810)

					gen5812 := Call(__e, ShenFunc(symshen_4pre), gen5811, MakeNumber(0))

					gen5813 := Call(__e, ShenFunc(sym_1), MakeNumber(0), gen5812)

					gen5820 = Call(__e, ShenFunc(symshen_4pair), gen5809, gen5813)

				} else {
					gen5820 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen5820 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5820
			gen5829 := Call(__e, ShenFunc(symfail))

			gen5830 := Call(__e, ShenFunc(sym_a), YaccParse, gen5829)

			if True == gen5830 {
				gen5821 := Call(__e, ShenFunc(symshen_4_5digits_6), V1608)

				Parse__shen_4_5digits_6 := gen5821
				gen5826 := Call(__e, ShenFunc(symfail))

				gen5827 := Call(__e, ShenFunc(sym_a), gen5826, Parse__shen_4_5digits_6)

				gen5828 := Call(__e, ShenFunc(symnot), gen5827)

				if True == gen5828 {
					gen5822 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

					gen5823 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

					gen5824 := Call(__e, ShenFunc(symreverse), gen5823)

					gen5825 := Call(__e, ShenFunc(symshen_4pre), gen5824, MakeNumber(0))

					__e.TailApply(ShenFunc(symshen_4pair), gen5822, gen5825)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<log10>"), gen5831)

		gen5840 := MakeNative(func(__e Evaluator) {
			V1610 := __e.Get(1)
			_ = V1610
			gen5838 := Call(__e, ShenFunc(symhd), V1610)

			gen5839 := Call(__e, ShenFunc(symcons_2), gen5838)

			if True == gen5839 {
				gen5832 := Call(__e, ShenFunc(symshen_4hdhd), V1610)

				Parse__Char := gen5832
				gen5837 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(43))

				if True == gen5837 {
					gen5833 := Call(__e, ShenFunc(symshen_4tlhd), V1610)

					gen5834 := Call(__e, ShenFunc(symshen_4hdtl), V1610)

					gen5835 := Call(__e, ShenFunc(symshen_4pair), gen5833, gen5834)

					gen5836 := Call(__e, ShenFunc(symhd), gen5835)

					__e.TailApply(ShenFunc(symshen_4pair), gen5836, Parse__Char)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<plus>"), gen5840)

		gen5849 := MakeNative(func(__e Evaluator) {
			V1612 := __e.Get(1)
			_ = V1612
			gen5847 := Call(__e, ShenFunc(symhd), V1612)

			gen5848 := Call(__e, ShenFunc(symcons_2), gen5847)

			if True == gen5848 {
				gen5841 := Call(__e, ShenFunc(symshen_4hdhd), V1612)

				Parse__Char := gen5841
				gen5846 := Call(__e, ShenFunc(sym_a), Parse__Char, MakeNumber(46))

				if True == gen5846 {
					gen5842 := Call(__e, ShenFunc(symshen_4tlhd), V1612)

					gen5843 := Call(__e, ShenFunc(symshen_4hdtl), V1612)

					gen5844 := Call(__e, ShenFunc(symshen_4pair), gen5842, gen5843)

					gen5845 := Call(__e, ShenFunc(symhd), gen5844)

					__e.TailApply(ShenFunc(symshen_4pair), gen5845, Parse__Char)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<stop>"), gen5849)

		gen5864 := MakeNative(func(__e Evaluator) {
			V1614 := __e.Get(1)
			_ = V1614
			gen5850 := Call(__e, ShenFunc(symshen_4_5digits_6), V1614)

			Parse__shen_4_5digits_6 := gen5850
			gen5853 := Call(__e, ShenFunc(symfail))

			gen5854 := Call(__e, ShenFunc(sym_a), gen5853, Parse__shen_4_5digits_6)

			gen5855 := Call(__e, ShenFunc(symnot), gen5854)

			var gen5856 Obj
			if True == gen5855 {
				gen5851 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

				gen5852 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

				gen5856 = Call(__e, ShenFunc(symshen_4pair), gen5851, gen5852)

			} else {
				gen5856 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5856
			gen5862 := Call(__e, ShenFunc(symfail))

			gen5863 := Call(__e, ShenFunc(sym_a), YaccParse, gen5862)

			if True == gen5863 {
				gen5857 := Call(__e, ShenFunc(sym_5e_6), V1614)

				Parse___5e_6 := gen5857
				gen5859 := Call(__e, ShenFunc(symfail))

				gen5860 := Call(__e, ShenFunc(sym_a), gen5859, Parse___5e_6)

				gen5861 := Call(__e, ShenFunc(symnot), gen5860)

				if True == gen5861 {
					gen5858 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen5858, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<predigits>"), gen5864)

		gen5871 := MakeNative(func(__e Evaluator) {
			V1616 := __e.Get(1)
			_ = V1616
			gen5865 := Call(__e, ShenFunc(symshen_4_5digits_6), V1616)

			Parse__shen_4_5digits_6 := gen5865
			gen5868 := Call(__e, ShenFunc(symfail))

			gen5869 := Call(__e, ShenFunc(sym_a), gen5868, Parse__shen_4_5digits_6)

			gen5870 := Call(__e, ShenFunc(symnot), gen5869)

			if True == gen5870 {
				gen5866 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

				gen5867 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen5866, gen5867)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<postdigits>"), gen5871)

		gen5894 := MakeNative(func(__e Evaluator) {
			V1618 := __e.Get(1)
			_ = V1618
			gen5872 := Call(__e, ShenFunc(symshen_4_5digit_6), V1618)

			Parse__shen_4_5digit_6 := gen5872
			gen5881 := Call(__e, ShenFunc(symfail))

			gen5882 := Call(__e, ShenFunc(sym_a), gen5881, Parse__shen_4_5digit_6)

			gen5883 := Call(__e, ShenFunc(symnot), gen5882)

			var gen5884 Obj
			if True == gen5883 {
				gen5873 := Call(__e, ShenFunc(symshen_4_5digits_6), Parse__shen_4_5digit_6)

				Parse__shen_4_5digits_6 := gen5873
				gen5878 := Call(__e, ShenFunc(symfail))

				gen5879 := Call(__e, ShenFunc(sym_a), gen5878, Parse__shen_4_5digits_6)

				gen5880 := Call(__e, ShenFunc(symnot), gen5879)

				if True == gen5880 {
					gen5874 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digits_6)

					gen5875 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digit_6)

					gen5876 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digits_6)

					gen5877 := Call(__e, ShenFunc(symcons), gen5875, gen5876)

					gen5884 = Call(__e, ShenFunc(symshen_4pair), gen5874, gen5877)

				} else {
					gen5884 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen5884 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5884
			gen5892 := Call(__e, ShenFunc(symfail))

			gen5893 := Call(__e, ShenFunc(sym_a), YaccParse, gen5892)

			if True == gen5893 {
				gen5885 := Call(__e, ShenFunc(symshen_4_5digit_6), V1618)

				Parse__shen_4_5digit_6 := gen5885
				gen5889 := Call(__e, ShenFunc(symfail))

				gen5890 := Call(__e, ShenFunc(sym_a), gen5889, Parse__shen_4_5digit_6)

				gen5891 := Call(__e, ShenFunc(symnot), gen5890)

				if True == gen5891 {
					gen5886 := Call(__e, ShenFunc(symhd), Parse__shen_4_5digit_6)

					gen5887 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5digit_6)

					gen5888 := Call(__e, ShenFunc(symcons), gen5887, Nil)

					__e.TailApply(ShenFunc(symshen_4pair), gen5886, gen5888)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<digits>"), gen5894)

		gen5904 := MakeNative(func(__e Evaluator) {
			V1620 := __e.Get(1)
			_ = V1620
			gen5902 := Call(__e, ShenFunc(symhd), V1620)

			gen5903 := Call(__e, ShenFunc(symcons_2), gen5902)

			if True == gen5903 {
				gen5895 := Call(__e, ShenFunc(symshen_4hdhd), V1620)

				Parse__X := gen5895
				gen5901 := Call(__e, ShenFunc(symshen_4numbyte_2), Parse__X)

				if True == gen5901 {
					gen5896 := Call(__e, ShenFunc(symshen_4tlhd), V1620)

					gen5897 := Call(__e, ShenFunc(symshen_4hdtl), V1620)

					gen5898 := Call(__e, ShenFunc(symshen_4pair), gen5896, gen5897)

					gen5899 := Call(__e, ShenFunc(symhd), gen5898)

					gen5900 := Call(__e, ShenFunc(symshen_4byte_1_6digit), Parse__X)

					__e.TailApply(ShenFunc(symshen_4pair), gen5899, gen5900)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<digit>"), gen5904)

		gen5915 := MakeNative(func(__e Evaluator) {
			V1622 := __e.Get(1)
			_ = V1622
			gen5914 := Call(__e, ShenFunc(sym_a), MakeNumber(48), V1622)

			if True == gen5914 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen5913 := Call(__e, ShenFunc(sym_a), MakeNumber(49), V1622)

				if True == gen5913 {
					__e.Return(MakeNumber(1))
					return
				} else {
					gen5912 := Call(__e, ShenFunc(sym_a), MakeNumber(50), V1622)

					if True == gen5912 {
						__e.Return(MakeNumber(2))
						return
					} else {
						gen5911 := Call(__e, ShenFunc(sym_a), MakeNumber(51), V1622)

						if True == gen5911 {
							__e.Return(MakeNumber(3))
							return
						} else {
							gen5910 := Call(__e, ShenFunc(sym_a), MakeNumber(52), V1622)

							if True == gen5910 {
								__e.Return(MakeNumber(4))
								return
							} else {
								gen5909 := Call(__e, ShenFunc(sym_a), MakeNumber(53), V1622)

								if True == gen5909 {
									__e.Return(MakeNumber(5))
									return
								} else {
									gen5908 := Call(__e, ShenFunc(sym_a), MakeNumber(54), V1622)

									if True == gen5908 {
										__e.Return(MakeNumber(6))
										return
									} else {
										gen5907 := Call(__e, ShenFunc(sym_a), MakeNumber(55), V1622)

										if True == gen5907 {
											__e.Return(MakeNumber(7))
											return
										} else {
											gen5906 := Call(__e, ShenFunc(sym_a), MakeNumber(56), V1622)

											if True == gen5906 {
												__e.Return(MakeNumber(8))
												return
											} else {
												gen5905 := Call(__e, ShenFunc(sym_a), MakeNumber(57), V1622)

												if True == gen5905 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.byte->digit"), gen5915)

		gen5924 := MakeNative(func(__e Evaluator) {
			V1627 := __e.Get(1)
			_ = V1627
			V1628 := __e.Get(2)
			_ = V1628
			gen5923 := Call(__e, ShenFunc(sym_a), Nil, V1627)

			if True == gen5923 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen5922 := Call(__e, ShenFunc(symcons_2), V1627)

				if True == gen5922 {
					gen5916 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), V1628)

					gen5917 := Call(__e, ShenFunc(symhd), V1627)

					gen5918 := Call(__e, ShenFunc(sym_d), gen5916, gen5917)

					gen5919 := Call(__e, ShenFunc(symtl), V1627)

					gen5920 := Call(__e, ShenFunc(sym_7), V1628, MakeNumber(1))

					gen5921 := Call(__e, ShenFunc(symshen_4pre), gen5919, gen5920)

					__e.TailApply(ShenFunc(sym_7), gen5918, gen5921)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.pre"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pre"), gen5924)

		gen5934 := MakeNative(func(__e Evaluator) {
			V1633 := __e.Get(1)
			_ = V1633
			V1634 := __e.Get(2)
			_ = V1634
			gen5933 := Call(__e, ShenFunc(sym_a), Nil, V1633)

			if True == gen5933 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen5932 := Call(__e, ShenFunc(symcons_2), V1633)

				if True == gen5932 {
					gen5925 := Call(__e, ShenFunc(sym_1), MakeNumber(0), V1634)

					gen5926 := Call(__e, ShenFunc(symshen_4expt), MakeNumber(10), gen5925)

					gen5927 := Call(__e, ShenFunc(symhd), V1633)

					gen5928 := Call(__e, ShenFunc(sym_d), gen5926, gen5927)

					gen5929 := Call(__e, ShenFunc(symtl), V1633)

					gen5930 := Call(__e, ShenFunc(sym_7), V1634, MakeNumber(1))

					gen5931 := Call(__e, ShenFunc(symshen_4post), gen5929, gen5930)

					__e.TailApply(ShenFunc(sym_7), gen5928, gen5931)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.post"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.post"), gen5934)

		gen5942 := MakeNative(func(__e Evaluator) {
			V1639 := __e.Get(1)
			_ = V1639
			V1640 := __e.Get(2)
			_ = V1640
			gen5941 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V1640)

			if True == gen5941 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen5940 := Call(__e, ShenFunc(sym_6), V1640, MakeNumber(0))

				if True == gen5940 {
					gen5938 := Call(__e, ShenFunc(sym_1), V1640, MakeNumber(1))

					gen5939 := Call(__e, ShenFunc(symshen_4expt), V1639, gen5938)

					__e.TailApply(ShenFunc(sym_d), V1639, gen5939)

					return

				} else {
					gen5935 := Call(__e, ShenFunc(sym_7), V1640, MakeNumber(1))

					gen5936 := Call(__e, ShenFunc(symshen_4expt), V1639, gen5935)

					gen5937 := Call(__e, ShenFunc(sym_c), gen5936, V1639)

					__e.TailApply(ShenFunc(sym_d), MakeNumber(1), gen5937)

					return

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.expt"), gen5942)

		gen5949 := MakeNative(func(__e Evaluator) {
			V1642 := __e.Get(1)
			_ = V1642
			gen5943 := Call(__e, ShenFunc(symshen_4_5st__input_6), V1642)

			Parse__shen_4_5st__input_6 := gen5943
			gen5946 := Call(__e, ShenFunc(symfail))

			gen5947 := Call(__e, ShenFunc(sym_a), gen5946, Parse__shen_4_5st__input_6)

			gen5948 := Call(__e, ShenFunc(symnot), gen5947)

			if True == gen5948 {
				gen5944 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

				gen5945 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen5944, gen5945)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<st_input1>"), gen5949)

		gen5956 := MakeNative(func(__e Evaluator) {
			V1644 := __e.Get(1)
			_ = V1644
			gen5950 := Call(__e, ShenFunc(symshen_4_5st__input_6), V1644)

			Parse__shen_4_5st__input_6 := gen5950
			gen5953 := Call(__e, ShenFunc(symfail))

			gen5954 := Call(__e, ShenFunc(sym_a), gen5953, Parse__shen_4_5st__input_6)

			gen5955 := Call(__e, ShenFunc(symnot), gen5954)

			if True == gen5955 {
				gen5951 := Call(__e, ShenFunc(symhd), Parse__shen_4_5st__input_6)

				gen5952 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5st__input_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen5951, gen5952)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<st_input2>"), gen5956)

		gen5970 := MakeNative(func(__e Evaluator) {
			V1646 := __e.Get(1)
			_ = V1646
			gen5957 := Call(__e, ShenFunc(symshen_4_5singleline_6), V1646)

			Parse__shen_4_5singleline_6 := gen5957
			gen5959 := Call(__e, ShenFunc(symfail))

			gen5960 := Call(__e, ShenFunc(sym_a), gen5959, Parse__shen_4_5singleline_6)

			gen5961 := Call(__e, ShenFunc(symnot), gen5960)

			var gen5962 Obj
			if True == gen5961 {
				gen5958 := Call(__e, ShenFunc(symhd), Parse__shen_4_5singleline_6)

				gen5962 = Call(__e, ShenFunc(symshen_4pair), gen5958, MakeSymbol("shen.skip"))

			} else {
				gen5962 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen5962
			gen5968 := Call(__e, ShenFunc(symfail))

			gen5969 := Call(__e, ShenFunc(sym_a), YaccParse, gen5968)

			if True == gen5969 {
				gen5963 := Call(__e, ShenFunc(symshen_4_5multiline_6), V1646)

				Parse__shen_4_5multiline_6 := gen5963
				gen5965 := Call(__e, ShenFunc(symfail))

				gen5966 := Call(__e, ShenFunc(sym_a), gen5965, Parse__shen_4_5multiline_6)

				gen5967 := Call(__e, ShenFunc(symnot), gen5966)

				if True == gen5967 {
					gen5964 := Call(__e, ShenFunc(symhd), Parse__shen_4_5multiline_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen5964, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<comment>"), gen5970)

		gen5988 := MakeNative(func(__e Evaluator) {
			V1648 := __e.Get(1)
			_ = V1648
			gen5971 := Call(__e, ShenFunc(symshen_4_5backslash_6), V1648)

			Parse__shen_4_5backslash_6 := gen5971
			gen5985 := Call(__e, ShenFunc(symfail))

			gen5986 := Call(__e, ShenFunc(sym_a), gen5985, Parse__shen_4_5backslash_6)

			gen5987 := Call(__e, ShenFunc(symnot), gen5986)

			if True == gen5987 {
				gen5972 := Call(__e, ShenFunc(symshen_4_5backslash_6), Parse__shen_4_5backslash_6)

				Parse__shen_4_5backslash_6 := gen5972
				gen5982 := Call(__e, ShenFunc(symfail))

				gen5983 := Call(__e, ShenFunc(sym_a), gen5982, Parse__shen_4_5backslash_6)

				gen5984 := Call(__e, ShenFunc(symnot), gen5983)

				if True == gen5984 {
					gen5973 := Call(__e, ShenFunc(symshen_4_5anysingle_6), Parse__shen_4_5backslash_6)

					Parse__shen_4_5anysingle_6 := gen5973
					gen5979 := Call(__e, ShenFunc(symfail))

					gen5980 := Call(__e, ShenFunc(sym_a), gen5979, Parse__shen_4_5anysingle_6)

					gen5981 := Call(__e, ShenFunc(symnot), gen5980)

					if True == gen5981 {
						gen5974 := Call(__e, ShenFunc(symshen_4_5return_6), Parse__shen_4_5anysingle_6)

						Parse__shen_4_5return_6 := gen5974
						gen5976 := Call(__e, ShenFunc(symfail))

						gen5977 := Call(__e, ShenFunc(sym_a), gen5976, Parse__shen_4_5return_6)

						gen5978 := Call(__e, ShenFunc(symnot), gen5977)

						if True == gen5978 {
							gen5975 := Call(__e, ShenFunc(symhd), Parse__shen_4_5return_6)

							__e.TailApply(ShenFunc(symshen_4pair), gen5975, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<singleline>"), gen5988)

		gen5998 := MakeNative(func(__e Evaluator) {
			V1651 := __e.Get(1)
			_ = V1651
			gen5995 := Call(__e, ShenFunc(symhd), V1651)

			gen5996 := Call(__e, ShenFunc(symcons_2), gen5995)

			var gen5997 Obj
			if True == gen5996 {
				gen5993 := Call(__e, ShenFunc(symshen_4hdhd), V1651)

				gen5994 := Call(__e, ShenFunc(sym_a), MakeNumber(92), gen5993)

				if True == gen5994 {
					gen5997 = True
				} else {
					gen5997 = False
				}

			} else {
				gen5997 = False
			}
			if True == gen5997 {
				gen5989 := Call(__e, ShenFunc(symshen_4tlhd), V1651)

				gen5990 := Call(__e, ShenFunc(symshen_4hdtl), V1651)

				gen5991 := Call(__e, ShenFunc(symshen_4pair), gen5989, gen5990)

				NewStream1649 := gen5991
				gen5992 := Call(__e, ShenFunc(symhd), NewStream1649)

				__e.TailApply(ShenFunc(symshen_4pair), gen5992, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<backslash>"), gen5998)

		gen6016 := MakeNative(func(__e Evaluator) {
			V1653 := __e.Get(1)
			_ = V1653
			gen5999 := Call(__e, ShenFunc(symshen_4_5non_1return_6), V1653)

			Parse__shen_4_5non_1return_6 := gen5999
			gen6005 := Call(__e, ShenFunc(symfail))

			gen6006 := Call(__e, ShenFunc(sym_a), gen6005, Parse__shen_4_5non_1return_6)

			gen6007 := Call(__e, ShenFunc(symnot), gen6006)

			var gen6008 Obj
			if True == gen6007 {
				gen6000 := Call(__e, ShenFunc(symshen_4_5anysingle_6), Parse__shen_4_5non_1return_6)

				Parse__shen_4_5anysingle_6 := gen6000
				gen6002 := Call(__e, ShenFunc(symfail))

				gen6003 := Call(__e, ShenFunc(sym_a), gen6002, Parse__shen_4_5anysingle_6)

				gen6004 := Call(__e, ShenFunc(symnot), gen6003)

				if True == gen6004 {
					gen6001 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anysingle_6)

					gen6008 = Call(__e, ShenFunc(symshen_4pair), gen6001, MakeSymbol("shen.skip"))

				} else {
					gen6008 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6008 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6008
			gen6014 := Call(__e, ShenFunc(symfail))

			gen6015 := Call(__e, ShenFunc(sym_a), YaccParse, gen6014)

			if True == gen6015 {
				gen6009 := Call(__e, ShenFunc(sym_5e_6), V1653)

				Parse___5e_6 := gen6009
				gen6011 := Call(__e, ShenFunc(symfail))

				gen6012 := Call(__e, ShenFunc(sym_a), gen6011, Parse___5e_6)

				gen6013 := Call(__e, ShenFunc(symnot), gen6012)

				if True == gen6013 {
					gen6010 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen6010, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<anysingle>"), gen6016)

		gen6028 := MakeNative(func(__e Evaluator) {
			V1655 := __e.Get(1)
			_ = V1655
			gen6026 := Call(__e, ShenFunc(symhd), V1655)

			gen6027 := Call(__e, ShenFunc(symcons_2), gen6026)

			if True == gen6027 {
				gen6017 := Call(__e, ShenFunc(symshen_4hdhd), V1655)

				Parse__X := gen6017
				gen6022 := Call(__e, ShenFunc(symcons), MakeNumber(13), Nil)

				gen6023 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen6022)

				gen6024 := Call(__e, ShenFunc(symelement_2), Parse__X, gen6023)

				gen6025 := Call(__e, ShenFunc(symnot), gen6024)

				if True == gen6025 {
					gen6018 := Call(__e, ShenFunc(symshen_4tlhd), V1655)

					gen6019 := Call(__e, ShenFunc(symshen_4hdtl), V1655)

					gen6020 := Call(__e, ShenFunc(symshen_4pair), gen6018, gen6019)

					gen6021 := Call(__e, ShenFunc(symhd), gen6020)

					__e.TailApply(ShenFunc(symshen_4pair), gen6021, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<non-return>"), gen6028)

		gen6039 := MakeNative(func(__e Evaluator) {
			V1657 := __e.Get(1)
			_ = V1657
			gen6037 := Call(__e, ShenFunc(symhd), V1657)

			gen6038 := Call(__e, ShenFunc(symcons_2), gen6037)

			if True == gen6038 {
				gen6029 := Call(__e, ShenFunc(symshen_4hdhd), V1657)

				Parse__X := gen6029
				gen6034 := Call(__e, ShenFunc(symcons), MakeNumber(13), Nil)

				gen6035 := Call(__e, ShenFunc(symcons), MakeNumber(10), gen6034)

				gen6036 := Call(__e, ShenFunc(symelement_2), Parse__X, gen6035)

				if True == gen6036 {
					gen6030 := Call(__e, ShenFunc(symshen_4tlhd), V1657)

					gen6031 := Call(__e, ShenFunc(symshen_4hdtl), V1657)

					gen6032 := Call(__e, ShenFunc(symshen_4pair), gen6030, gen6031)

					gen6033 := Call(__e, ShenFunc(symhd), gen6032)

					__e.TailApply(ShenFunc(symshen_4pair), gen6033, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<return>"), gen6039)

		gen6053 := MakeNative(func(__e Evaluator) {
			V1659 := __e.Get(1)
			_ = V1659
			gen6040 := Call(__e, ShenFunc(symshen_4_5backslash_6), V1659)

			Parse__shen_4_5backslash_6 := gen6040
			gen6050 := Call(__e, ShenFunc(symfail))

			gen6051 := Call(__e, ShenFunc(sym_a), gen6050, Parse__shen_4_5backslash_6)

			gen6052 := Call(__e, ShenFunc(symnot), gen6051)

			if True == gen6052 {
				gen6041 := Call(__e, ShenFunc(symshen_4_5times_6), Parse__shen_4_5backslash_6)

				Parse__shen_4_5times_6 := gen6041
				gen6047 := Call(__e, ShenFunc(symfail))

				gen6048 := Call(__e, ShenFunc(sym_a), gen6047, Parse__shen_4_5times_6)

				gen6049 := Call(__e, ShenFunc(symnot), gen6048)

				if True == gen6049 {
					gen6042 := Call(__e, ShenFunc(symshen_4_5anymulti_6), Parse__shen_4_5times_6)

					Parse__shen_4_5anymulti_6 := gen6042
					gen6044 := Call(__e, ShenFunc(symfail))

					gen6045 := Call(__e, ShenFunc(sym_a), gen6044, Parse__shen_4_5anymulti_6)

					gen6046 := Call(__e, ShenFunc(symnot), gen6045)

					if True == gen6046 {
						gen6043 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anymulti_6)

						__e.TailApply(ShenFunc(symshen_4pair), gen6043, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<multiline>"), gen6053)

		gen6063 := MakeNative(func(__e Evaluator) {
			V1662 := __e.Get(1)
			_ = V1662
			gen6060 := Call(__e, ShenFunc(symhd), V1662)

			gen6061 := Call(__e, ShenFunc(symcons_2), gen6060)

			var gen6062 Obj
			if True == gen6061 {
				gen6058 := Call(__e, ShenFunc(symshen_4hdhd), V1662)

				gen6059 := Call(__e, ShenFunc(sym_a), MakeNumber(42), gen6058)

				if True == gen6059 {
					gen6062 = True
				} else {
					gen6062 = False
				}

			} else {
				gen6062 = False
			}
			if True == gen6062 {
				gen6054 := Call(__e, ShenFunc(symshen_4tlhd), V1662)

				gen6055 := Call(__e, ShenFunc(symshen_4hdtl), V1662)

				gen6056 := Call(__e, ShenFunc(symshen_4pair), gen6054, gen6055)

				NewStream1660 := gen6056
				gen6057 := Call(__e, ShenFunc(symhd), NewStream1660)

				__e.TailApply(ShenFunc(symshen_4pair), gen6057, MakeSymbol("shen.skip"))

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<times>"), gen6063)

		gen6099 := MakeNative(func(__e Evaluator) {
			V1664 := __e.Get(1)
			_ = V1664
			gen6064 := Call(__e, ShenFunc(symshen_4_5comment_6), V1664)

			Parse__shen_4_5comment_6 := gen6064
			gen6070 := Call(__e, ShenFunc(symfail))

			gen6071 := Call(__e, ShenFunc(sym_a), gen6070, Parse__shen_4_5comment_6)

			gen6072 := Call(__e, ShenFunc(symnot), gen6071)

			var gen6073 Obj
			if True == gen6072 {
				gen6065 := Call(__e, ShenFunc(symshen_4_5anymulti_6), Parse__shen_4_5comment_6)

				Parse__shen_4_5anymulti_6 := gen6065
				gen6067 := Call(__e, ShenFunc(symfail))

				gen6068 := Call(__e, ShenFunc(sym_a), gen6067, Parse__shen_4_5anymulti_6)

				gen6069 := Call(__e, ShenFunc(symnot), gen6068)

				if True == gen6069 {
					gen6066 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anymulti_6)

					gen6073 = Call(__e, ShenFunc(symshen_4pair), gen6066, MakeSymbol("shen.skip"))

				} else {
					gen6073 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6073 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6073
			gen6097 := Call(__e, ShenFunc(symfail))

			gen6098 := Call(__e, ShenFunc(sym_a), YaccParse, gen6097)

			if True == gen6098 {
				gen6074 := Call(__e, ShenFunc(symshen_4_5times_6), V1664)

				Parse__shen_4_5times_6 := gen6074
				gen6080 := Call(__e, ShenFunc(symfail))

				gen6081 := Call(__e, ShenFunc(sym_a), gen6080, Parse__shen_4_5times_6)

				gen6082 := Call(__e, ShenFunc(symnot), gen6081)

				var gen6083 Obj
				if True == gen6082 {
					gen6075 := Call(__e, ShenFunc(symshen_4_5backslash_6), Parse__shen_4_5times_6)

					Parse__shen_4_5backslash_6 := gen6075
					gen6077 := Call(__e, ShenFunc(symfail))

					gen6078 := Call(__e, ShenFunc(sym_a), gen6077, Parse__shen_4_5backslash_6)

					gen6079 := Call(__e, ShenFunc(symnot), gen6078)

					if True == gen6079 {
						gen6076 := Call(__e, ShenFunc(symhd), Parse__shen_4_5backslash_6)

						gen6083 = Call(__e, ShenFunc(symshen_4pair), gen6076, MakeSymbol("shen.skip"))

					} else {
						gen6083 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen6083 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen6083
				gen6095 := Call(__e, ShenFunc(symfail))

				gen6096 := Call(__e, ShenFunc(sym_a), YaccParse, gen6095)

				if True == gen6096 {
					gen6093 := Call(__e, ShenFunc(symhd), V1664)

					gen6094 := Call(__e, ShenFunc(symcons_2), gen6093)

					if True == gen6094 {
						gen6084 := Call(__e, ShenFunc(symshen_4hdhd), V1664)

						Parse__X := gen6084
						_ = Parse__X
						gen6085 := Call(__e, ShenFunc(symshen_4tlhd), V1664)

						gen6086 := Call(__e, ShenFunc(symshen_4hdtl), V1664)

						gen6087 := Call(__e, ShenFunc(symshen_4pair), gen6085, gen6086)

						gen6088 := Call(__e, ShenFunc(symshen_4_5anymulti_6), gen6087)

						Parse__shen_4_5anymulti_6 := gen6088
						gen6090 := Call(__e, ShenFunc(symfail))

						gen6091 := Call(__e, ShenFunc(sym_a), gen6090, Parse__shen_4_5anymulti_6)

						gen6092 := Call(__e, ShenFunc(symnot), gen6091)

						if True == gen6092 {
							gen6089 := Call(__e, ShenFunc(symhd), Parse__shen_4_5anymulti_6)

							__e.TailApply(ShenFunc(symshen_4pair), gen6089, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<anymulti>"), gen6099)

		gen6117 := MakeNative(func(__e Evaluator) {
			V1666 := __e.Get(1)
			_ = V1666
			gen6100 := Call(__e, ShenFunc(symshen_4_5whitespace_6), V1666)

			Parse__shen_4_5whitespace_6 := gen6100
			gen6106 := Call(__e, ShenFunc(symfail))

			gen6107 := Call(__e, ShenFunc(sym_a), gen6106, Parse__shen_4_5whitespace_6)

			gen6108 := Call(__e, ShenFunc(symnot), gen6107)

			var gen6109 Obj
			if True == gen6108 {
				gen6101 := Call(__e, ShenFunc(symshen_4_5whitespaces_6), Parse__shen_4_5whitespace_6)

				Parse__shen_4_5whitespaces_6 := gen6101
				gen6103 := Call(__e, ShenFunc(symfail))

				gen6104 := Call(__e, ShenFunc(sym_a), gen6103, Parse__shen_4_5whitespaces_6)

				gen6105 := Call(__e, ShenFunc(symnot), gen6104)

				if True == gen6105 {
					gen6102 := Call(__e, ShenFunc(symhd), Parse__shen_4_5whitespaces_6)

					gen6109 = Call(__e, ShenFunc(symshen_4pair), gen6102, MakeSymbol("shen.skip"))

				} else {
					gen6109 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6109 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6109
			gen6115 := Call(__e, ShenFunc(symfail))

			gen6116 := Call(__e, ShenFunc(sym_a), YaccParse, gen6115)

			if True == gen6116 {
				gen6110 := Call(__e, ShenFunc(symshen_4_5whitespace_6), V1666)

				Parse__shen_4_5whitespace_6 := gen6110
				gen6112 := Call(__e, ShenFunc(symfail))

				gen6113 := Call(__e, ShenFunc(sym_a), gen6112, Parse__shen_4_5whitespace_6)

				gen6114 := Call(__e, ShenFunc(symnot), gen6113)

				if True == gen6114 {
					gen6111 := Call(__e, ShenFunc(symhd), Parse__shen_4_5whitespace_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen6111, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<whitespaces>"), gen6117)

		gen6132 := MakeNative(func(__e Evaluator) {
			V1668 := __e.Get(1)
			_ = V1668
			gen6130 := Call(__e, ShenFunc(symhd), V1668)

			gen6131 := Call(__e, ShenFunc(symcons_2), gen6130)

			if True == gen6131 {
				gen6118 := Call(__e, ShenFunc(symshen_4hdhd), V1668)

				Parse__X := gen6118
				Parse__Case := Parse__X
				gen6128 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(32))

				var gen6129 Obj
				if True == gen6128 {
					gen6129 = True
				} else {
					gen6126 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(13))

					var gen6127 Obj
					if True == gen6126 {
						gen6127 = True
					} else {
						gen6124 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(10))

						var gen6125 Obj
						if True == gen6124 {
							gen6125 = True
						} else {
							gen6123 := Call(__e, ShenFunc(sym_a), Parse__Case, MakeNumber(9))

							if True == gen6123 {
								gen6125 = True
							} else {
								gen6125 = False
							}

						}
						if True == gen6125 {
							gen6127 = True
						} else {
							gen6127 = False
						}

					}
					if True == gen6127 {
						gen6129 = True
					} else {
						gen6129 = False
					}

				}
				if True == gen6129 {
					gen6119 := Call(__e, ShenFunc(symshen_4tlhd), V1668)

					gen6120 := Call(__e, ShenFunc(symshen_4hdtl), V1668)

					gen6121 := Call(__e, ShenFunc(symshen_4pair), gen6119, gen6120)

					gen6122 := Call(__e, ShenFunc(symhd), gen6121)

					__e.TailApply(ShenFunc(symshen_4pair), gen6122, MakeSymbol("shen.skip"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<whitespace>"), gen6132)

		gen6161 := MakeNative(func(__e Evaluator) {
			V1670 := __e.Get(1)
			_ = V1670
			gen6160 := Call(__e, ShenFunc(sym_a), Nil, V1670)

			if True == gen6160 {
				__e.Return(Nil)
				return
			} else {
				gen6158 := Call(__e, ShenFunc(symcons_2), V1670)

				var gen6159 Obj
				if True == gen6158 {
					gen6155 := Call(__e, ShenFunc(symtl), V1670)

					gen6156 := Call(__e, ShenFunc(symcons_2), gen6155)

					var gen6157 Obj
					if True == gen6156 {
						gen6151 := Call(__e, ShenFunc(symtl), V1670)

						gen6152 := Call(__e, ShenFunc(symtl), gen6151)

						gen6153 := Call(__e, ShenFunc(symcons_2), gen6152)

						var gen6154 Obj
						if True == gen6153 {
							gen6146 := Call(__e, ShenFunc(symtl), V1670)

							gen6147 := Call(__e, ShenFunc(symtl), gen6146)

							gen6148 := Call(__e, ShenFunc(symtl), gen6147)

							gen6149 := Call(__e, ShenFunc(sym_a), Nil, gen6148)

							var gen6150 Obj
							if True == gen6149 {
								gen6143 := Call(__e, ShenFunc(symtl), V1670)

								gen6144 := Call(__e, ShenFunc(symhd), gen6143)

								gen6145 := Call(__e, ShenFunc(sym_a), gen6144, MakeSymbol("bar!"))

								if True == gen6145 {
									gen6150 = True
								} else {
									gen6150 = False
								}

							} else {
								gen6150 = False
							}
							if True == gen6150 {
								gen6154 = True
							} else {
								gen6154 = False
							}

						} else {
							gen6154 = False
						}
						if True == gen6154 {
							gen6157 = True
						} else {
							gen6157 = False
						}

					} else {
						gen6157 = False
					}
					if True == gen6157 {
						gen6159 = True
					} else {
						gen6159 = False
					}

				} else {
					gen6159 = False
				}
				if True == gen6159 {
					gen6139 := Call(__e, ShenFunc(symhd), V1670)

					gen6140 := Call(__e, ShenFunc(symtl), V1670)

					gen6141 := Call(__e, ShenFunc(symtl), gen6140)

					gen6142 := Call(__e, ShenFunc(symcons), gen6139, gen6141)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen6142)

					return

				} else {
					gen6138 := Call(__e, ShenFunc(symcons_2), V1670)

					if True == gen6138 {
						gen6133 := Call(__e, ShenFunc(symhd), V1670)

						gen6134 := Call(__e, ShenFunc(symtl), V1670)

						gen6135 := Call(__e, ShenFunc(symshen_4cons__form), gen6134)

						gen6136 := Call(__e, ShenFunc(symcons), gen6135, Nil)

						gen6137 := Call(__e, ShenFunc(symcons), gen6133, gen6136)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen6137)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.cons_form"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cons_form"), gen6161)

		gen6226 := MakeNative(func(__e Evaluator) {
			V1675 := __e.Get(1)
			_ = V1675
			V1676 := __e.Get(2)
			_ = V1676
			gen6224 := Call(__e, ShenFunc(symcons_2), V1675)

			var gen6225 Obj
			if True == gen6224 {
				gen6221 := Call(__e, ShenFunc(symhd), V1675)

				gen6222 := Call(__e, ShenFunc(sym_a), MakeSymbol("$"), gen6221)

				var gen6223 Obj
				if True == gen6222 {
					gen6218 := Call(__e, ShenFunc(symtl), V1675)

					gen6219 := Call(__e, ShenFunc(symcons_2), gen6218)

					var gen6220 Obj
					if True == gen6219 {
						gen6215 := Call(__e, ShenFunc(symtl), V1675)

						gen6216 := Call(__e, ShenFunc(symtl), gen6215)

						gen6217 := Call(__e, ShenFunc(sym_a), Nil, gen6216)

						if True == gen6217 {
							gen6220 = True
						} else {
							gen6220 = False
						}

					} else {
						gen6220 = False
					}
					if True == gen6220 {
						gen6223 = True
					} else {
						gen6223 = False
					}

				} else {
					gen6223 = False
				}
				if True == gen6223 {
					gen6225 = True
				} else {
					gen6225 = False
				}

			} else {
				gen6225 = False
			}
			if True == gen6225 {
				gen6212 := Call(__e, ShenFunc(symtl), V1675)

				gen6213 := Call(__e, ShenFunc(symhd), gen6212)

				gen6214 := Call(__e, ShenFunc(symexplode), gen6213)

				__e.TailApply(ShenFunc(symappend), gen6214, V1676)

				return

			} else {
				gen6210 := Call(__e, ShenFunc(symcons_2), V1675)

				var gen6211 Obj
				if True == gen6210 {
					gen6207 := Call(__e, ShenFunc(symhd), V1675)

					gen6208 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen6207)

					var gen6209 Obj
					if True == gen6208 {
						gen6204 := Call(__e, ShenFunc(symtl), V1675)

						gen6205 := Call(__e, ShenFunc(symcons_2), gen6204)

						var gen6206 Obj
						if True == gen6205 {
							gen6200 := Call(__e, ShenFunc(symtl), V1675)

							gen6201 := Call(__e, ShenFunc(symhd), gen6200)

							gen6202 := Call(__e, ShenFunc(sym_a), MakeSymbol("null"), gen6201)

							var gen6203 Obj
							if True == gen6202 {
								gen6197 := Call(__e, ShenFunc(symtl), V1675)

								gen6198 := Call(__e, ShenFunc(symtl), gen6197)

								gen6199 := Call(__e, ShenFunc(symcons_2), gen6198)

								if True == gen6199 {
									gen6203 = True
								} else {
									gen6203 = False
								}

							} else {
								gen6203 = False
							}
							if True == gen6203 {
								gen6206 = True
							} else {
								gen6206 = False
							}

						} else {
							gen6206 = False
						}
						if True == gen6206 {
							gen6209 = True
						} else {
							gen6209 = False
						}

					} else {
						gen6209 = False
					}
					if True == gen6209 {
						gen6211 = True
					} else {
						gen6211 = False
					}

				} else {
					gen6211 = False
				}
				if True == gen6211 {
					gen6194 := Call(__e, ShenFunc(symtl), V1675)

					gen6195 := Call(__e, ShenFunc(symtl), gen6194)

					gen6196 := Call(__e, ShenFunc(symtl), gen6195)

					__e.TailApply(ShenFunc(symappend), gen6196, V1676)

					return

				} else {
					gen6192 := Call(__e, ShenFunc(symcons_2), V1675)

					var gen6193 Obj
					if True == gen6192 {
						gen6189 := Call(__e, ShenFunc(symhd), V1675)

						gen6190 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen6189)

						var gen6191 Obj
						if True == gen6190 {
							gen6186 := Call(__e, ShenFunc(symtl), V1675)

							gen6187 := Call(__e, ShenFunc(symcons_2), gen6186)

							var gen6188 Obj
							if True == gen6187 {
								gen6183 := Call(__e, ShenFunc(symtl), V1675)

								gen6184 := Call(__e, ShenFunc(symtl), gen6183)

								gen6185 := Call(__e, ShenFunc(symcons_2), gen6184)

								if True == gen6185 {
									gen6188 = True
								} else {
									gen6188 = False
								}

							} else {
								gen6188 = False
							}
							if True == gen6188 {
								gen6191 = True
							} else {
								gen6191 = False
							}

						} else {
							gen6191 = False
						}
						if True == gen6191 {
							gen6193 = True
						} else {
							gen6193 = False
						}

					} else {
						gen6193 = False
					}
					if True == gen6193 {
						gen6162 := Call(__e, ShenFunc(symtl), V1675)

						gen6163 := Call(__e, ShenFunc(symtl), gen6162)

						gen6164 := Call(__e, ShenFunc(symhd), gen6163)

						gen6165 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), gen6164)

						ListofExceptions := gen6165
						gen6166 := Call(__e, ShenFunc(symtl), V1675)

						gen6167 := Call(__e, ShenFunc(symhd), gen6166)

						gen6168 := Call(__e, ShenFunc(symshen_4record_1exceptions), ListofExceptions, gen6167)

						External := gen6168
						_ = External
						gen6169 := Call(__e, ShenFunc(symtl), V1675)

						gen6170 := Call(__e, ShenFunc(symhd), gen6169)

						gen6171 := Call(__e, ShenFunc(symstr), gen6170)

						gen6172 := Call(__e, ShenFunc(symcn), gen6171, MakeString("."))

						gen6173 := Call(__e, ShenFunc(symintern), gen6172)

						PackageNameDot := gen6173
						gen6174 := Call(__e, ShenFunc(symexplode), PackageNameDot)

						ExpPackageNameDot := gen6174
						gen6175 := Call(__e, ShenFunc(symtl), V1675)

						gen6176 := Call(__e, ShenFunc(symtl), gen6175)

						gen6177 := Call(__e, ShenFunc(symtl), gen6176)

						gen6178 := Call(__e, ShenFunc(symshen_4packageh), PackageNameDot, ListofExceptions, gen6177, ExpPackageNameDot)

						Packaged := gen6178
						gen6179 := Call(__e, ShenFunc(symtl), V1675)

						gen6180 := Call(__e, ShenFunc(symhd), gen6179)

						gen6181 := Call(__e, ShenFunc(symshen_4internal_1symbols), ExpPackageNameDot, Packaged)

						gen6182 := Call(__e, ShenFunc(symshen_4record_1internal), gen6180, gen6181)

						Internal := gen6182
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.package-macro"), gen6226)

		gen6233 := MakeNative(func(__e Evaluator) {
			V1679 := __e.Get(1)
			_ = V1679
			V1680 := __e.Get(2)
			_ = V1680
			gen6227 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen6229 := MakeNative(func(__e Evaluator) {
				gen6228 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1680, MakeSymbol("shen.external-symbols"), gen6228)

				return

			}, 0)
			gen6230 := Try(__e, gen6229).Catch(gen6227)
			CurrExceptions := gen6230
			gen6231 := Call(__e, ShenFunc(symunion), V1679, CurrExceptions)

			AllExceptions := gen6231
			gen6232 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			__e.TailApply(ShenFunc(symput), V1680, MakeSymbol("shen.external-symbols"), AllExceptions, gen6232)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-exceptions"), gen6233)

		gen6240 := MakeNative(func(__e Evaluator) {
			V1683 := __e.Get(1)
			_ = V1683
			V1684 := __e.Get(2)
			_ = V1684
			gen6234 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen6236 := MakeNative(func(__e Evaluator) {
				gen6235 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1683, MakeSymbol("shen.internal-symbols"), gen6235)

				return

			}, 0)
			gen6237 := Try(__e, gen6236).Catch(gen6234)
			gen6238 := Call(__e, ShenFunc(symunion), V1684, gen6237)

			gen6239 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			__e.TailApply(ShenFunc(symput), V1683, MakeSymbol("shen.internal-symbols"), gen6238, gen6239)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-internal"), gen6240)

		gen6250 := MakeNative(func(__e Evaluator) {
			V1695 := __e.Get(1)
			_ = V1695
			V1696 := __e.Get(2)
			_ = V1696
			gen6248 := Call(__e, ShenFunc(symsymbol_2), V1696)

			var gen6249 Obj
			if True == gen6248 {
				gen6246 := Call(__e, ShenFunc(symexplode), V1696)

				gen6247 := Call(__e, ShenFunc(symshen_4prefix_2), V1695, gen6246)

				if True == gen6247 {
					gen6249 = True
				} else {
					gen6249 = False
				}

			} else {
				gen6249 = False
			}
			if True == gen6249 {
				__e.TailApply(ShenFunc(symcons), V1696, Nil)

				return
			} else {
				gen6245 := Call(__e, ShenFunc(symcons_2), V1696)

				if True == gen6245 {
					gen6241 := Call(__e, ShenFunc(symhd), V1696)

					gen6242 := Call(__e, ShenFunc(symshen_4internal_1symbols), V1695, gen6241)

					gen6243 := Call(__e, ShenFunc(symtl), V1696)

					gen6244 := Call(__e, ShenFunc(symshen_4internal_1symbols), V1695, gen6243)

					__e.TailApply(ShenFunc(symunion), gen6242, gen6244)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.internal-symbols"), gen6250)

		gen6278 := MakeNative(func(__e Evaluator) {
			V1713 := __e.Get(1)
			_ = V1713
			V1714 := __e.Get(2)
			_ = V1714
			V1715 := __e.Get(3)
			_ = V1715
			V1716 := __e.Get(4)
			_ = V1716
			gen6277 := Call(__e, ShenFunc(symcons_2), V1715)

			if True == gen6277 {
				gen6273 := Call(__e, ShenFunc(symhd), V1715)

				gen6274 := Call(__e, ShenFunc(symshen_4packageh), V1713, V1714, gen6273, V1716)

				gen6275 := Call(__e, ShenFunc(symtl), V1715)

				gen6276 := Call(__e, ShenFunc(symshen_4packageh), V1713, V1714, gen6275, V1716)

				__e.TailApply(ShenFunc(symcons), gen6274, gen6276)

				return

			} else {
				gen6271 := Call(__e, ShenFunc(symshen_4sysfunc_2), V1715)

				var gen6272 Obj
				if True == gen6271 {
					gen6272 = True
				} else {
					gen6269 := Call(__e, ShenFunc(symvariable_2), V1715)

					var gen6270 Obj
					if True == gen6269 {
						gen6270 = True
					} else {
						gen6267 := Call(__e, ShenFunc(symelement_2), V1715, V1714)

						var gen6268 Obj
						if True == gen6267 {
							gen6268 = True
						} else {
							gen6265 := Call(__e, ShenFunc(symshen_4doubleunderline_2), V1715)

							var gen6266 Obj
							if True == gen6265 {
								gen6266 = True
							} else {
								gen6264 := Call(__e, ShenFunc(symshen_4singleunderline_2), V1715)

								if True == gen6264 {
									gen6266 = True
								} else {
									gen6266 = False
								}

							}
							if True == gen6266 {
								gen6268 = True
							} else {
								gen6268 = False
							}

						}
						if True == gen6268 {
							gen6270 = True
						} else {
							gen6270 = False
						}

					}
					if True == gen6270 {
						gen6272 = True
					} else {
						gen6272 = False
					}

				}
				if True == gen6272 {
					__e.Return(V1715)
					return
				} else {
					gen6262 := Call(__e, ShenFunc(symsymbol_2), V1715)

					var gen6263 Obj
					if True == gen6262 {
						gen6251 := Call(__e, ShenFunc(symexplode), V1715)

						ExplodeX := gen6251
						gen6254 := Call(__e, ShenFunc(symcons), MakeString("."), Nil)

						gen6255 := Call(__e, ShenFunc(symcons), MakeString("n"), gen6254)

						gen6256 := Call(__e, ShenFunc(symcons), MakeString("e"), gen6255)

						gen6257 := Call(__e, ShenFunc(symcons), MakeString("h"), gen6256)

						gen6258 := Call(__e, ShenFunc(symcons), MakeString("s"), gen6257)

						gen6259 := Call(__e, ShenFunc(symshen_4prefix_2), gen6258, ExplodeX)

						gen6260 := Call(__e, ShenFunc(symnot), gen6259)

						var gen6261 Obj
						if True == gen6260 {
							gen6252 := Call(__e, ShenFunc(symshen_4prefix_2), V1716, ExplodeX)

							gen6253 := Call(__e, ShenFunc(symnot), gen6252)

							if True == gen6253 {
								gen6261 = True
							} else {
								gen6261 = False
							}

						} else {
							gen6261 = False
						}
						if True == gen6261 {
							gen6263 = True
						} else {
							gen6263 = False
						}

					} else {
						gen6263 = False
					}
					if True == gen6263 {
						__e.TailApply(ShenFunc(symconcat), V1713, V1715)

						return
					} else {
						__e.Return(V1715)
						return
					}

				}

			}

		}, 4)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.packageh"), gen6278)

		return

	}, 0))
}
