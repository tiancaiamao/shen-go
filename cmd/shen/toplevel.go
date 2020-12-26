package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen15837 := MakeNative(func(__e Evaluator, __args ...Obj) {
			Call(__e, ShenFunc(symshen_4credits))
			__e.TailApply(ShenFunc(symshen_4loop))

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.repl"), gen15837)

		gen15840 := MakeNative(func(__e Evaluator, __args ...Obj) {
			Call(__e, ShenFunc(symshen_4initialise__environment))
			Call(__e, ShenFunc(symshen_4prompt))
			gen15838 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.TailApply(ShenFunc(symshen_4toplevel_1display_1exception), E)

				return
			}, 1)
			gen15839 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4read_1evaluate_1print))

				return
			}, 0)
			Try(__e, gen15839).Catch(gen15838)

			__e.TailApply(ShenFunc(symshen_4loop))

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.loop"), gen15840)

		gen15843 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3022 := __args[0]
			_ = V3022
			gen15841 := Call(__e, ShenFunc(symerror_1to_1string), V3022)

			gen15842 := Call(__e, ShenFunc(symstoutput))

			__e.TailApply(ShenFunc(sympr), gen15841, gen15842)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.toplevel-display-exception"), gen15843)

		gen15863 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen15844 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n"), gen15844)

			gen15845 := Call(__e, ShenFunc(symvalue), MakeSymbol("*version*"))

			gen15846 := Call(__e, ShenFunc(symshen_4app), gen15845, MakeString("\n"), MakeSymbol("shen.a"))

			gen15847 := Call(__e, ShenFunc(symcn), MakeString("www.shenlanguage.org, "), gen15846)

			gen15848 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen15847, gen15848)

			gen15849 := Call(__e, ShenFunc(symvalue), MakeSymbol("*language*"))

			gen15850 := Call(__e, ShenFunc(symvalue), MakeSymbol("*implementation*"))

			gen15851 := Call(__e, ShenFunc(symshen_4app), gen15850, MakeString(""), MakeSymbol("shen.a"))

			gen15852 := Call(__e, ShenFunc(symcn), MakeString(", implementation: "), gen15851)

			gen15853 := Call(__e, ShenFunc(symshen_4app), gen15849, gen15852, MakeSymbol("shen.a"))

			gen15854 := Call(__e, ShenFunc(symcn), MakeString("running under "), gen15853)

			gen15855 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen15854, gen15855)

			gen15856 := Call(__e, ShenFunc(symvalue), MakeSymbol("*port*"))

			gen15857 := Call(__e, ShenFunc(symvalue), MakeSymbol("*porters*"))

			gen15858 := Call(__e, ShenFunc(symshen_4app), gen15857, MakeString("\n"), MakeSymbol("shen.a"))

			gen15859 := Call(__e, ShenFunc(symcn), MakeString(" ported by "), gen15858)

			gen15860 := Call(__e, ShenFunc(symshen_4app), gen15856, gen15859, MakeSymbol("shen.a"))

			gen15861 := Call(__e, ShenFunc(symcn), MakeString("\nport "), gen15860)

			gen15862 := Call(__e, ShenFunc(symstoutput))

			__e.TailApply(ShenFunc(symshen_4prhush), gen15861, gen15862)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.credits"), gen15863)

		gen15872 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen15864 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

			gen15865 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*catch*"), gen15864)

			gen15866 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15865)

			gen15867 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*process-counter*"), gen15866)

			gen15868 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15867)

			gen15869 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*infs*"), gen15868)

			gen15870 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15869)

			gen15871 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), gen15870)

			__e.TailApply(ShenFunc(symshen_4multiple_1set), gen15871)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise_environment"), gen15872)

		gen15883 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3024 := __args[0]
			_ = V3024
			gen15882 := Call(__e, ShenFunc(sym_a), Nil, V3024)

			if True == gen15882 {
				__e.Return(Nil)
				return
			} else {
				gen15880 := Call(__e, ShenFunc(symcons_2), V3024)

				var gen15881 Obj
				if True == gen15880 {
					gen15878 := Call(__e, ShenFunc(symtl), V3024)

					gen15879 := Call(__e, ShenFunc(symcons_2), gen15878)

					if True == gen15879 {
						gen15881 = True
					} else {
						gen15881 = False
					}

				} else {
					gen15881 = False
				}
				if True == gen15881 {
					gen15873 := Call(__e, ShenFunc(symhd), V3024)

					gen15874 := Call(__e, ShenFunc(symtl), V3024)

					gen15875 := Call(__e, ShenFunc(symhd), gen15874)

					Call(__e, ShenFunc(symset), gen15873, gen15875)

					gen15876 := Call(__e, ShenFunc(symtl), V3024)

					gen15877 := Call(__e, ShenFunc(symtl), gen15876)

					__e.TailApply(ShenFunc(symshen_4multiple_1set), gen15877)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.multiple-set"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.multiple-set"), gen15883)

		gen15884 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3026 := __args[0]
			_ = V3026
			__e.TailApply(ShenFunc(symdeclare), V3026, MakeSymbol("symbol"))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("destroy"), gen15884)

		gen15890 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen15885 := Call(__e, ShenFunc(symshen_4toplineread))

			Lineread := gen15885
			gen15886 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*history*"))

			History := gen15886
			gen15887 := Call(__e, ShenFunc(symshen_4retrieve_1from_1history_1if_1needed), Lineread, History)

			NewLineread := gen15887
			gen15888 := Call(__e, ShenFunc(symshen_4update__history), NewLineread, History)

			NewHistory := gen15888
			_ = NewHistory
			gen15889 := Call(__e, ShenFunc(symfst), NewLineread)

			Parsed := gen15889
			__e.TailApply(ShenFunc(symshen_4toplevel), Parsed)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.read-evaluate-print"), gen15890)

		gen15981 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3038 := __args[0]
			_ = V3038
			V3039 := __args[1]
			_ = V3039
			gen15979 := Call(__e, ShenFunc(symtuple_2), V3038)

			var gen15980 Obj
			if True == gen15979 {
				gen15976 := Call(__e, ShenFunc(symsnd), V3038)

				gen15977 := Call(__e, ShenFunc(symcons_2), gen15976)

				var gen15978 Obj
				if True == gen15977 {
					gen15969 := Call(__e, ShenFunc(symsnd), V3038)

					gen15970 := Call(__e, ShenFunc(symhd), gen15969)

					gen15971 := Call(__e, ShenFunc(symshen_4space))

					gen15972 := Call(__e, ShenFunc(symshen_4newline))

					gen15973 := Call(__e, ShenFunc(symcons), gen15972, Nil)

					gen15974 := Call(__e, ShenFunc(symcons), gen15971, gen15973)

					gen15975 := Call(__e, ShenFunc(symelement_2), gen15970, gen15974)

					if True == gen15975 {
						gen15978 = True
					} else {
						gen15978 = False
					}

				} else {
					gen15978 = False
				}
				if True == gen15978 {
					gen15980 = True
				} else {
					gen15980 = False
				}

			} else {
				gen15980 = False
			}
			if True == gen15980 {
				gen15965 := Call(__e, ShenFunc(symfst), V3038)

				gen15966 := Call(__e, ShenFunc(symsnd), V3038)

				gen15967 := Call(__e, ShenFunc(symtl), gen15966)

				gen15968 := Call(__e, ShenFunc(sym_8p), gen15965, gen15967)

				__e.TailApply(ShenFunc(symshen_4retrieve_1from_1history_1if_1needed), gen15968, V3039)

				return

			} else {
				gen15963 := Call(__e, ShenFunc(symtuple_2), V3038)

				var gen15964 Obj
				if True == gen15963 {
					gen15960 := Call(__e, ShenFunc(symsnd), V3038)

					gen15961 := Call(__e, ShenFunc(symcons_2), gen15960)

					var gen15962 Obj
					if True == gen15961 {
						gen15956 := Call(__e, ShenFunc(symsnd), V3038)

						gen15957 := Call(__e, ShenFunc(symtl), gen15956)

						gen15958 := Call(__e, ShenFunc(symcons_2), gen15957)

						var gen15959 Obj
						if True == gen15958 {
							gen15951 := Call(__e, ShenFunc(symsnd), V3038)

							gen15952 := Call(__e, ShenFunc(symtl), gen15951)

							gen15953 := Call(__e, ShenFunc(symtl), gen15952)

							gen15954 := Call(__e, ShenFunc(sym_a), Nil, gen15953)

							var gen15955 Obj
							if True == gen15954 {
								gen15949 := Call(__e, ShenFunc(symcons_2), V3039)

								var gen15950 Obj
								if True == gen15949 {
									gen15944 := Call(__e, ShenFunc(symsnd), V3038)

									gen15945 := Call(__e, ShenFunc(symhd), gen15944)

									gen15946 := Call(__e, ShenFunc(symshen_4exclamation))

									gen15947 := Call(__e, ShenFunc(sym_a), gen15945, gen15946)

									var gen15948 Obj
									if True == gen15947 {
										gen15939 := Call(__e, ShenFunc(symsnd), V3038)

										gen15940 := Call(__e, ShenFunc(symtl), gen15939)

										gen15941 := Call(__e, ShenFunc(symhd), gen15940)

										gen15942 := Call(__e, ShenFunc(symshen_4exclamation))

										gen15943 := Call(__e, ShenFunc(sym_a), gen15941, gen15942)

										if True == gen15943 {
											gen15948 = True
										} else {
											gen15948 = False
										}

									} else {
										gen15948 = False
									}
									if True == gen15948 {
										gen15950 = True
									} else {
										gen15950 = False
									}

								} else {
									gen15950 = False
								}
								if True == gen15950 {
									gen15955 = True
								} else {
									gen15955 = False
								}

							} else {
								gen15955 = False
							}
							if True == gen15955 {
								gen15959 = True
							} else {
								gen15959 = False
							}

						} else {
							gen15959 = False
						}
						if True == gen15959 {
							gen15962 = True
						} else {
							gen15962 = False
						}

					} else {
						gen15962 = False
					}
					if True == gen15962 {
						gen15964 = True
					} else {
						gen15964 = False
					}

				} else {
					gen15964 = False
				}
				if True == gen15964 {
					gen15936 := Call(__e, ShenFunc(symhd), V3039)

					gen15937 := Call(__e, ShenFunc(symsnd), gen15936)

					gen15938 := Call(__e, ShenFunc(symshen_4prbytes), gen15937)

					PastPrint := gen15938
					_ = PastPrint
					__e.TailApply(ShenFunc(symhd), V3039)

					return

				} else {
					gen15934 := Call(__e, ShenFunc(symtuple_2), V3038)

					var gen15935 Obj
					if True == gen15934 {
						gen15931 := Call(__e, ShenFunc(symsnd), V3038)

						gen15932 := Call(__e, ShenFunc(symcons_2), gen15931)

						var gen15933 Obj
						if True == gen15932 {
							gen15927 := Call(__e, ShenFunc(symsnd), V3038)

							gen15928 := Call(__e, ShenFunc(symhd), gen15927)

							gen15929 := Call(__e, ShenFunc(symshen_4exclamation))

							gen15930 := Call(__e, ShenFunc(sym_a), gen15928, gen15929)

							if True == gen15930 {
								gen15933 = True
							} else {
								gen15933 = False
							}

						} else {
							gen15933 = False
						}
						if True == gen15933 {
							gen15935 = True
						} else {
							gen15935 = False
						}

					} else {
						gen15935 = False
					}
					if True == gen15935 {
						gen15920 := Call(__e, ShenFunc(symsnd), V3038)

						gen15921 := Call(__e, ShenFunc(symtl), gen15920)

						gen15922 := Call(__e, ShenFunc(symshen_4make_1key), gen15921, V3039)

						Key_2 := gen15922
						gen15923 := Call(__e, ShenFunc(symshen_4find_1past_1inputs), Key_2, V3039)

						gen15924 := Call(__e, ShenFunc(symhead), gen15923)

						Find := gen15924
						gen15925 := Call(__e, ShenFunc(symsnd), Find)

						gen15926 := Call(__e, ShenFunc(symshen_4prbytes), gen15925)

						PastPrint := gen15926
						_ = PastPrint
						__e.Return(Find)
						return

					} else {
						gen15918 := Call(__e, ShenFunc(symtuple_2), V3038)

						var gen15919 Obj
						if True == gen15918 {
							gen15915 := Call(__e, ShenFunc(symsnd), V3038)

							gen15916 := Call(__e, ShenFunc(symcons_2), gen15915)

							var gen15917 Obj
							if True == gen15916 {
								gen15911 := Call(__e, ShenFunc(symsnd), V3038)

								gen15912 := Call(__e, ShenFunc(symtl), gen15911)

								gen15913 := Call(__e, ShenFunc(sym_a), Nil, gen15912)

								var gen15914 Obj
								if True == gen15913 {
									gen15907 := Call(__e, ShenFunc(symsnd), V3038)

									gen15908 := Call(__e, ShenFunc(symhd), gen15907)

									gen15909 := Call(__e, ShenFunc(symshen_4percent))

									gen15910 := Call(__e, ShenFunc(sym_a), gen15908, gen15909)

									if True == gen15910 {
										gen15914 = True
									} else {
										gen15914 = False
									}

								} else {
									gen15914 = False
								}
								if True == gen15914 {
									gen15917 = True
								} else {
									gen15917 = False
								}

							} else {
								gen15917 = False
							}
							if True == gen15917 {
								gen15919 = True
							} else {
								gen15919 = False
							}

						} else {
							gen15919 = False
						}
						if True == gen15919 {
							gen15905 := MakeNative(func(__e Evaluator, __args ...Obj) {
								X := __args[0]
								_ = X
								__e.Return(True)
								return
							}, 1)
							gen15906 := Call(__e, ShenFunc(symreverse), V3039)

							Call(__e, ShenFunc(symshen_4print_1past_1inputs), gen15905, gen15906, MakeNumber(0))

							__e.TailApply(ShenFunc(symabort))

							return

						} else {
							gen15903 := Call(__e, ShenFunc(symtuple_2), V3038)

							var gen15904 Obj
							if True == gen15903 {
								gen15900 := Call(__e, ShenFunc(symsnd), V3038)

								gen15901 := Call(__e, ShenFunc(symcons_2), gen15900)

								var gen15902 Obj
								if True == gen15901 {
									gen15896 := Call(__e, ShenFunc(symsnd), V3038)

									gen15897 := Call(__e, ShenFunc(symhd), gen15896)

									gen15898 := Call(__e, ShenFunc(symshen_4percent))

									gen15899 := Call(__e, ShenFunc(sym_a), gen15897, gen15898)

									if True == gen15899 {
										gen15902 = True
									} else {
										gen15902 = False
									}

								} else {
									gen15902 = False
								}
								if True == gen15902 {
									gen15904 = True
								} else {
									gen15904 = False
								}

							} else {
								gen15904 = False
							}
							if True == gen15904 {
								gen15891 := Call(__e, ShenFunc(symsnd), V3038)

								gen15892 := Call(__e, ShenFunc(symtl), gen15891)

								gen15893 := Call(__e, ShenFunc(symshen_4make_1key), gen15892, V3039)

								Key_2 := gen15893
								gen15894 := Call(__e, ShenFunc(symreverse), V3039)

								gen15895 := Call(__e, ShenFunc(symshen_4print_1past_1inputs), Key_2, gen15894, MakeNumber(0))

								Pastprint := gen15895
								_ = Pastprint
								__e.TailApply(ShenFunc(symabort))

								return

							} else {
								__e.Return(V3038)
								return
							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.retrieve-from-history-if-needed"), gen15981)

		gen15982 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(37))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.percent"), gen15982)

		gen15983 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(33))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.exclamation"), gen15983)

		gen15987 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3041 := __args[0]
			_ = V3041
			gen15986 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Byte := __args[0]
				_ = Byte
				gen15984 := Call(__e, ShenFunc(symn_1_6string), Byte)

				gen15985 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(sympr), gen15984, gen15985)

				return

			}, 1)
			Call(__e, ShenFunc(symshen_4for_1each), gen15986, V3041)

			__e.TailApply(ShenFunc(symnl), MakeNumber(1))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prbytes"), gen15987)

		gen15989 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3044 := __args[0]
			_ = V3044
			V3045 := __args[1]
			_ = V3045
			gen15988 := Call(__e, ShenFunc(symcons), V3044, V3045)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*history*"), gen15988)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.update_history"), gen15989)

		gen15992 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen15990 := Call(__e, ShenFunc(symstinput))

			gen15991 := Call(__e, ShenFunc(symshen_4read_1char_1code), gen15990)

			__e.TailApply(ShenFunc(symshen_4toplineread__loop), gen15991, Nil)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.toplineread"), gen15992)

		gen16016 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3049 := __args[0]
			_ = V3049
			V3050 := __args[1]
			_ = V3050
			gen16014 := Call(__e, ShenFunc(symshen_4hat))

			gen16015 := Call(__e, ShenFunc(sym_a), V3049, gen16014)

			if True == gen16015 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("line read aborted"))

				return
			} else {
				gen16009 := Call(__e, ShenFunc(symshen_4newline))

				gen16010 := Call(__e, ShenFunc(symshen_4carriage_1return))

				gen16011 := Call(__e, ShenFunc(symcons), gen16010, Nil)

				gen16012 := Call(__e, ShenFunc(symcons), gen16009, gen16011)

				gen16013 := Call(__e, ShenFunc(symelement_2), V3049, gen16012)

				if True == gen16013 {
					gen15998 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

						return
					}, 1)
					gen15999 := MakeNative(func(__e Evaluator, __args ...Obj) {
						E := __args[0]
						_ = E
						__e.Return(MakeSymbol("shen.nextline"))
						return
					}, 1)
					gen16000 := Call(__e, ShenFunc(symcompile), gen15998, V3050, gen15999)

					Line := gen16000
					gen16001 := Call(__e, ShenFunc(symshen_4record_1it), V3050)

					It := gen16001
					_ = It
					gen16007 := Call(__e, ShenFunc(sym_a), Line, MakeSymbol("shen.nextline"))

					var gen16008 Obj
					if True == gen16007 {
						gen16008 = True
					} else {
						gen16006 := Call(__e, ShenFunc(symempty_2), Line)

						if True == gen16006 {
							gen16008 = True
						} else {
							gen16008 = False
						}

					}
					if True == gen16008 {
						gen16002 := Call(__e, ShenFunc(symstinput))

						gen16003 := Call(__e, ShenFunc(symshen_4read_1char_1code), gen16002)

						gen16004 := Call(__e, ShenFunc(symcons), V3049, Nil)

						gen16005 := Call(__e, ShenFunc(symappend), V3050, gen16004)

						__e.TailApply(ShenFunc(symshen_4toplineread__loop), gen16003, gen16005)

						return

					} else {
						__e.TailApply(ShenFunc(sym_8p), Line, V3050)

						return
					}

				} else {
					gen15993 := Call(__e, ShenFunc(symstinput))

					gen15994 := Call(__e, ShenFunc(symshen_4read_1char_1code), gen15993)

					gen15996 := Call(__e, ShenFunc(sym_a), V3049, MakeNumber(-1))

					var gen15997 Obj
					if True == gen15996 {
						gen15997 = V3050
					} else {
						gen15995 := Call(__e, ShenFunc(symcons), V3049, Nil)

						gen15997 = Call(__e, ShenFunc(symappend), V3050, gen15995)

					}
					__e.TailApply(ShenFunc(symshen_4toplineread__loop), gen15994, gen15997)

					return

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.toplineread_loop"), gen16016)

		gen16017 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(94))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.hat"), gen16017)

		gen16018 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(10))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newline"), gen16018)

		gen16019 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(13))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.carriage-return"), gen16019)

		gen16022 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3056 := __args[0]
			_ = V3056
			gen16021 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V3056)

			if True == gen16021 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*tc*"), True)

				return
			} else {
				gen16020 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V3056)

				if True == gen16020 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*tc*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("tc expects a + or -"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tc"), gen16022)

		gen16034 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen16033 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

			if True == gen16033 {
				gen16028 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*history*"))

				gen16029 := Call(__e, ShenFunc(symlength), gen16028)

				gen16030 := Call(__e, ShenFunc(symshen_4app), gen16029, MakeString("+) "), MakeSymbol("shen.a"))

				gen16031 := Call(__e, ShenFunc(symcn), MakeString("\n\n("), gen16030)

				gen16032 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen16031, gen16032)

				return

			} else {
				gen16023 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*history*"))

				gen16024 := Call(__e, ShenFunc(symlength), gen16023)

				gen16025 := Call(__e, ShenFunc(symshen_4app), gen16024, MakeString("-) "), MakeSymbol("shen.a"))

				gen16026 := Call(__e, ShenFunc(symcn), MakeString("\n\n("), gen16025)

				gen16027 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen16026, gen16027)

				return

			}

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prompt"), gen16034)

		gen16036 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3058 := __args[0]
			_ = V3058
			gen16035 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

			__e.TailApply(ShenFunc(symshen_4toplevel__evaluate), V3058, gen16035)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.toplevel"), gen16036)

		gen16039 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3061 := __args[0]
			_ = V3061
			V3062 := __args[1]
			_ = V3062
			gen16037 := Call(__e, ShenFunc(symshen_4find), V3061, V3062)

			F := gen16037
			gen16038 := Call(__e, ShenFunc(symempty_2), F)

			if True == gen16038 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("input not found\n"))

				return
			} else {
				__e.Return(F)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.find-past-inputs"), gen16039)

		gen16053 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3065 := __args[0]
			_ = V3065
			V3066 := __args[1]
			_ = V3066
			gen16040 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5st__input_6), X)

				return
			}, 1)
			gen16044 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				gen16043 := Call(__e, ShenFunc(symcons_2), E)

				if True == gen16043 {
					gen16041 := Call(__e, ShenFunc(symshen_4app), E, MakeString("\n"), MakeSymbol("shen.s"))

					gen16042 := Call(__e, ShenFunc(symcn), MakeString("parse error here: "), gen16041)

					__e.TailApply(ShenFunc(symsimple_1error), gen16042)

					return

				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("parse error\n"))

					return
				}

			}, 1)
			gen16045 := Call(__e, ShenFunc(symcompile), gen16040, V3065, gen16044)

			gen16046 := Call(__e, ShenFunc(symhd), gen16045)

			Atom := gen16046
			gen16052 := Call(__e, ShenFunc(syminteger_2), Atom)

			if True == gen16052 {
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					gen16049 := Call(__e, ShenFunc(sym_7), Atom, MakeNumber(1))

					gen16050 := Call(__e, ShenFunc(symreverse), V3066)

					gen16051 := Call(__e, ShenFunc(symnth), gen16049, gen16050)

					__e.TailApply(ShenFunc(sym_a), X, gen16051)

					return

				}, 1))
				return
			} else {
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					gen16047 := Call(__e, ShenFunc(symsnd), X)

					gen16048 := Call(__e, ShenFunc(symshen_4trim_1gubbins), gen16047)

					__e.TailApply(ShenFunc(symshen_4prefix_2), V3065, gen16048)

					return

				}, 1))
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.make-key"), gen16053)

		gen16084 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3068 := __args[0]
			_ = V3068
			gen16082 := Call(__e, ShenFunc(symcons_2), V3068)

			var gen16083 Obj
			if True == gen16082 {
				gen16079 := Call(__e, ShenFunc(symhd), V3068)

				gen16080 := Call(__e, ShenFunc(symshen_4space))

				gen16081 := Call(__e, ShenFunc(sym_a), gen16079, gen16080)

				if True == gen16081 {
					gen16083 = True
				} else {
					gen16083 = False
				}

			} else {
				gen16083 = False
			}
			if True == gen16083 {
				gen16078 := Call(__e, ShenFunc(symtl), V3068)

				__e.TailApply(ShenFunc(symshen_4trim_1gubbins), gen16078)

				return

			} else {
				gen16076 := Call(__e, ShenFunc(symcons_2), V3068)

				var gen16077 Obj
				if True == gen16076 {
					gen16073 := Call(__e, ShenFunc(symhd), V3068)

					gen16074 := Call(__e, ShenFunc(symshen_4newline))

					gen16075 := Call(__e, ShenFunc(sym_a), gen16073, gen16074)

					if True == gen16075 {
						gen16077 = True
					} else {
						gen16077 = False
					}

				} else {
					gen16077 = False
				}
				if True == gen16077 {
					gen16072 := Call(__e, ShenFunc(symtl), V3068)

					__e.TailApply(ShenFunc(symshen_4trim_1gubbins), gen16072)

					return

				} else {
					gen16070 := Call(__e, ShenFunc(symcons_2), V3068)

					var gen16071 Obj
					if True == gen16070 {
						gen16067 := Call(__e, ShenFunc(symhd), V3068)

						gen16068 := Call(__e, ShenFunc(symshen_4carriage_1return))

						gen16069 := Call(__e, ShenFunc(sym_a), gen16067, gen16068)

						if True == gen16069 {
							gen16071 = True
						} else {
							gen16071 = False
						}

					} else {
						gen16071 = False
					}
					if True == gen16071 {
						gen16066 := Call(__e, ShenFunc(symtl), V3068)

						__e.TailApply(ShenFunc(symshen_4trim_1gubbins), gen16066)

						return

					} else {
						gen16064 := Call(__e, ShenFunc(symcons_2), V3068)

						var gen16065 Obj
						if True == gen16064 {
							gen16061 := Call(__e, ShenFunc(symhd), V3068)

							gen16062 := Call(__e, ShenFunc(symshen_4tab))

							gen16063 := Call(__e, ShenFunc(sym_a), gen16061, gen16062)

							if True == gen16063 {
								gen16065 = True
							} else {
								gen16065 = False
							}

						} else {
							gen16065 = False
						}
						if True == gen16065 {
							gen16060 := Call(__e, ShenFunc(symtl), V3068)

							__e.TailApply(ShenFunc(symshen_4trim_1gubbins), gen16060)

							return

						} else {
							gen16058 := Call(__e, ShenFunc(symcons_2), V3068)

							var gen16059 Obj
							if True == gen16058 {
								gen16055 := Call(__e, ShenFunc(symhd), V3068)

								gen16056 := Call(__e, ShenFunc(symshen_4left_1round))

								gen16057 := Call(__e, ShenFunc(sym_a), gen16055, gen16056)

								if True == gen16057 {
									gen16059 = True
								} else {
									gen16059 = False
								}

							} else {
								gen16059 = False
							}
							if True == gen16059 {
								gen16054 := Call(__e, ShenFunc(symtl), V3068)

								__e.TailApply(ShenFunc(symshen_4trim_1gubbins), gen16054)

								return

							} else {
								__e.Return(V3068)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.trim-gubbins"), gen16084)

		gen16085 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(32))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.space"), gen16085)

		gen16086 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(9))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tab"), gen16086)

		gen16087 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeNumber(40))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.left-round"), gen16087)

		gen16098 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3077 := __args[0]
			_ = V3077
			V3078 := __args[1]
			_ = V3078
			gen16097 := Call(__e, ShenFunc(sym_a), Nil, V3078)

			if True == gen16097 {
				__e.Return(Nil)
				return
			} else {
				gen16095 := Call(__e, ShenFunc(symcons_2), V3078)

				var gen16096 Obj
				if True == gen16095 {
					gen16093 := Call(__e, ShenFunc(symhd), V3078)

					gen16094 := Call(__e, V3077, gen16093)

					if True == gen16094 {
						gen16096 = True
					} else {
						gen16096 = False
					}

				} else {
					gen16096 = False
				}
				if True == gen16096 {
					gen16090 := Call(__e, ShenFunc(symhd), V3078)

					gen16091 := Call(__e, ShenFunc(symtl), V3078)

					gen16092 := Call(__e, ShenFunc(symshen_4find), V3077, gen16091)

					__e.TailApply(ShenFunc(symcons), gen16090, gen16092)

					return

				} else {
					gen16089 := Call(__e, ShenFunc(symcons_2), V3078)

					if True == gen16089 {
						gen16088 := Call(__e, ShenFunc(symtl), V3078)

						__e.TailApply(ShenFunc(symshen_4find), V3077, gen16088)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.find"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.find"), gen16098)

		gen16109 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3092 := __args[0]
			_ = V3092
			V3093 := __args[1]
			_ = V3093
			gen16108 := Call(__e, ShenFunc(sym_a), Nil, V3092)

			if True == gen16108 {
				__e.Return(True)
				return
			} else {
				gen16106 := Call(__e, ShenFunc(symcons_2), V3092)

				var gen16107 Obj
				if True == gen16106 {
					gen16104 := Call(__e, ShenFunc(symcons_2), V3093)

					var gen16105 Obj
					if True == gen16104 {
						gen16101 := Call(__e, ShenFunc(symhd), V3093)

						gen16102 := Call(__e, ShenFunc(symhd), V3092)

						gen16103 := Call(__e, ShenFunc(sym_a), gen16101, gen16102)

						if True == gen16103 {
							gen16105 = True
						} else {
							gen16105 = False
						}

					} else {
						gen16105 = False
					}
					if True == gen16105 {
						gen16107 = True
					} else {
						gen16107 = False
					}

				} else {
					gen16107 = False
				}
				if True == gen16107 {
					gen16099 := Call(__e, ShenFunc(symtl), V3092)

					gen16100 := Call(__e, ShenFunc(symtl), V3093)

					__e.TailApply(ShenFunc(symshen_4prefix_2), gen16099, gen16100)

					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prefix?"), gen16109)

		gen16128 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3105 := __args[0]
			_ = V3105
			V3106 := __args[1]
			_ = V3106
			V3107 := __args[2]
			_ = V3107
			gen16127 := Call(__e, ShenFunc(sym_a), Nil, V3106)

			if True == gen16127 {
				__e.Return(MakeSymbol("_"))
				return
			} else {
				gen16125 := Call(__e, ShenFunc(symcons_2), V3106)

				var gen16126 Obj
				if True == gen16125 {
					gen16122 := Call(__e, ShenFunc(symhd), V3106)

					gen16123 := Call(__e, V3105, gen16122)

					gen16124 := Call(__e, ShenFunc(symnot), gen16123)

					if True == gen16124 {
						gen16126 = True
					} else {
						gen16126 = False
					}

				} else {
					gen16126 = False
				}
				if True == gen16126 {
					gen16120 := Call(__e, ShenFunc(symtl), V3106)

					gen16121 := Call(__e, ShenFunc(sym_7), V3107, MakeNumber(1))

					__e.TailApply(ShenFunc(symshen_4print_1past_1inputs), V3105, gen16120, gen16121)

					return

				} else {
					gen16118 := Call(__e, ShenFunc(symcons_2), V3106)

					var gen16119 Obj
					if True == gen16118 {
						gen16116 := Call(__e, ShenFunc(symhd), V3106)

						gen16117 := Call(__e, ShenFunc(symtuple_2), gen16116)

						if True == gen16117 {
							gen16119 = True
						} else {
							gen16119 = False
						}

					} else {
						gen16119 = False
					}
					if True == gen16119 {
						gen16110 := Call(__e, ShenFunc(symshen_4app), V3107, MakeString(". "), MakeSymbol("shen.a"))

						gen16111 := Call(__e, ShenFunc(symstoutput))

						Call(__e, ShenFunc(symshen_4prhush), gen16110, gen16111)

						gen16112 := Call(__e, ShenFunc(symhd), V3106)

						gen16113 := Call(__e, ShenFunc(symsnd), gen16112)

						Call(__e, ShenFunc(symshen_4prbytes), gen16113)

						gen16114 := Call(__e, ShenFunc(symtl), V3106)

						gen16115 := Call(__e, ShenFunc(sym_7), V3107, MakeNumber(1))

						__e.TailApply(ShenFunc(symshen_4print_1past_1inputs), V3105, gen16114, gen16115)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.print-past-inputs"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.print-past-inputs"), gen16128)

		gen16175 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3110 := __args[0]
			_ = V3110
			V3111 := __args[1]
			_ = V3111
			gen16173 := Call(__e, ShenFunc(symcons_2), V3110)

			var gen16174 Obj
			if True == gen16173 {
				gen16170 := Call(__e, ShenFunc(symtl), V3110)

				gen16171 := Call(__e, ShenFunc(symcons_2), gen16170)

				var gen16172 Obj
				if True == gen16171 {
					gen16166 := Call(__e, ShenFunc(symtl), V3110)

					gen16167 := Call(__e, ShenFunc(symhd), gen16166)

					gen16168 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen16167)

					var gen16169 Obj
					if True == gen16168 {
						gen16162 := Call(__e, ShenFunc(symtl), V3110)

						gen16163 := Call(__e, ShenFunc(symtl), gen16162)

						gen16164 := Call(__e, ShenFunc(symcons_2), gen16163)

						var gen16165 Obj
						if True == gen16164 {
							gen16157 := Call(__e, ShenFunc(symtl), V3110)

							gen16158 := Call(__e, ShenFunc(symtl), gen16157)

							gen16159 := Call(__e, ShenFunc(symtl), gen16158)

							gen16160 := Call(__e, ShenFunc(sym_a), Nil, gen16159)

							var gen16161 Obj
							if True == gen16160 {
								gen16156 := Call(__e, ShenFunc(sym_a), True, V3111)

								if True == gen16156 {
									gen16161 = True
								} else {
									gen16161 = False
								}

							} else {
								gen16161 = False
							}
							if True == gen16161 {
								gen16165 = True
							} else {
								gen16165 = False
							}

						} else {
							gen16165 = False
						}
						if True == gen16165 {
							gen16169 = True
						} else {
							gen16169 = False
						}

					} else {
						gen16169 = False
					}
					if True == gen16169 {
						gen16172 = True
					} else {
						gen16172 = False
					}

				} else {
					gen16172 = False
				}
				if True == gen16172 {
					gen16174 = True
				} else {
					gen16174 = False
				}

			} else {
				gen16174 = False
			}
			if True == gen16174 {
				gen16152 := Call(__e, ShenFunc(symhd), V3110)

				gen16153 := Call(__e, ShenFunc(symtl), V3110)

				gen16154 := Call(__e, ShenFunc(symtl), gen16153)

				gen16155 := Call(__e, ShenFunc(symhd), gen16154)

				__e.TailApply(ShenFunc(symshen_4typecheck_1and_1evaluate), gen16152, gen16155)

				return

			} else {
				gen16150 := Call(__e, ShenFunc(symcons_2), V3110)

				var gen16151 Obj
				if True == gen16150 {
					gen16148 := Call(__e, ShenFunc(symtl), V3110)

					gen16149 := Call(__e, ShenFunc(symcons_2), gen16148)

					if True == gen16149 {
						gen16151 = True
					} else {
						gen16151 = False
					}

				} else {
					gen16151 = False
				}
				if True == gen16151 {
					gen16145 := Call(__e, ShenFunc(symhd), V3110)

					gen16146 := Call(__e, ShenFunc(symcons), gen16145, Nil)

					Call(__e, ShenFunc(symshen_4toplevel__evaluate), gen16146, V3111)

					Call(__e, ShenFunc(symnl), MakeNumber(1))
					gen16147 := Call(__e, ShenFunc(symtl), V3110)

					__e.TailApply(ShenFunc(symshen_4toplevel__evaluate), gen16147, V3111)

					return

				} else {
					gen16143 := Call(__e, ShenFunc(symcons_2), V3110)

					var gen16144 Obj
					if True == gen16143 {
						gen16140 := Call(__e, ShenFunc(symtl), V3110)

						gen16141 := Call(__e, ShenFunc(sym_a), Nil, gen16140)

						var gen16142 Obj
						if True == gen16141 {
							gen16139 := Call(__e, ShenFunc(sym_a), True, V3111)

							if True == gen16139 {
								gen16142 = True
							} else {
								gen16142 = False
							}

						} else {
							gen16142 = False
						}
						if True == gen16142 {
							gen16144 = True
						} else {
							gen16144 = False
						}

					} else {
						gen16144 = False
					}
					if True == gen16144 {
						gen16137 := Call(__e, ShenFunc(symhd), V3110)

						gen16138 := Call(__e, ShenFunc(symgensym), MakeSymbol("A"))

						__e.TailApply(ShenFunc(symshen_4typecheck_1and_1evaluate), gen16137, gen16138)

						return

					} else {
						gen16135 := Call(__e, ShenFunc(symcons_2), V3110)

						var gen16136 Obj
						if True == gen16135 {
							gen16132 := Call(__e, ShenFunc(symtl), V3110)

							gen16133 := Call(__e, ShenFunc(sym_a), Nil, gen16132)

							var gen16134 Obj
							if True == gen16133 {
								gen16131 := Call(__e, ShenFunc(sym_a), False, V3111)

								if True == gen16131 {
									gen16134 = True
								} else {
									gen16134 = False
								}

							} else {
								gen16134 = False
							}
							if True == gen16134 {
								gen16136 = True
							} else {
								gen16136 = False
							}

						} else {
							gen16136 = False
						}
						if True == gen16136 {
							gen16129 := Call(__e, ShenFunc(symhd), V3110)

							gen16130 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), gen16129)

							Eval := gen16130
							__e.TailApply(ShenFunc(symprint), Eval)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.toplevel_evaluate"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.toplevel_evaluate"), gen16175)

		gen16184 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3114 := __args[0]
			_ = V3114
			V3115 := __args[1]
			_ = V3115
			gen16176 := Call(__e, ShenFunc(symshen_4typecheck), V3114, V3115)

			Typecheck := gen16176
			gen16183 := Call(__e, ShenFunc(sym_a), Typecheck, False)

			if True == gen16183 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("type error\n"))

				return
			} else {
				gen16177 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), V3114)

				Eval := gen16177
				gen16178 := Call(__e, ShenFunc(symshen_4pretty_1type), Typecheck)

				Type := gen16178
				gen16179 := Call(__e, ShenFunc(symshen_4app), Type, MakeString(""), MakeSymbol("shen.r"))

				gen16180 := Call(__e, ShenFunc(symcn), MakeString(" : "), gen16179)

				gen16181 := Call(__e, ShenFunc(symshen_4app), Eval, gen16180, MakeSymbol("shen.s"))

				gen16182 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen16181, gen16182)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typecheck-and-evaluate"), gen16184)

		gen16187 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3117 := __args[0]
			_ = V3117
			gen16185 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*alphabet*"))

			gen16186 := Call(__e, ShenFunc(symshen_4extract_1pvars), V3117)

			__e.TailApply(ShenFunc(symshen_4mult__subst), gen16185, gen16186, V3117)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pretty-type"), gen16187)

		gen16194 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3123 := __args[0]
			_ = V3123
			gen16193 := Call(__e, ShenFunc(symshen_4pvar_2), V3123)

			if True == gen16193 {
				__e.TailApply(ShenFunc(symcons), V3123, Nil)

				return
			} else {
				gen16192 := Call(__e, ShenFunc(symcons_2), V3123)

				if True == gen16192 {
					gen16188 := Call(__e, ShenFunc(symhd), V3123)

					gen16189 := Call(__e, ShenFunc(symshen_4extract_1pvars), gen16188)

					gen16190 := Call(__e, ShenFunc(symtl), V3123)

					gen16191 := Call(__e, ShenFunc(symshen_4extract_1pvars), gen16190)

					__e.TailApply(ShenFunc(symunion), gen16189, gen16191)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extract-pvars"), gen16194)

		gen16205 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3131 := __args[0]
			_ = V3131
			V3132 := __args[1]
			_ = V3132
			V3133 := __args[2]
			_ = V3133
			gen16204 := Call(__e, ShenFunc(sym_a), Nil, V3131)

			if True == gen16204 {
				__e.Return(V3133)
				return
			} else {
				gen16203 := Call(__e, ShenFunc(sym_a), Nil, V3132)

				if True == gen16203 {
					__e.Return(V3133)
					return
				} else {
					gen16201 := Call(__e, ShenFunc(symcons_2), V3131)

					var gen16202 Obj
					if True == gen16201 {
						gen16200 := Call(__e, ShenFunc(symcons_2), V3132)

						if True == gen16200 {
							gen16202 = True
						} else {
							gen16202 = False
						}

					} else {
						gen16202 = False
					}
					if True == gen16202 {
						gen16195 := Call(__e, ShenFunc(symtl), V3131)

						gen16196 := Call(__e, ShenFunc(symtl), V3132)

						gen16197 := Call(__e, ShenFunc(symhd), V3131)

						gen16198 := Call(__e, ShenFunc(symhd), V3132)

						gen16199 := Call(__e, ShenFunc(symsubst), gen16197, gen16198, V3133)

						__e.TailApply(ShenFunc(symshen_4mult__subst), gen16195, gen16196, gen16199)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.mult_subst"))

						return
					}

				}

			}

		}, 3)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.mult_subst"), gen16205)

		return

	}, 0))
}
