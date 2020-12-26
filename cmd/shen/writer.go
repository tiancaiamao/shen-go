package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen18866 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4594 := __args[0]
			_ = V4594
			V4595 := __args[1]
			_ = V4595
			gen18864 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(V4594)
				return
			}, 1)
			gen18865 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4prh), V4594, V4595, MakeNumber(0))

				return
			}, 0)
			__e.Return(Try(__e, gen18865).Catch(gen18864))
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("pr"), gen18866)

		gen18868 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4599 := __args[0]
			_ = V4599
			V4600 := __args[1]
			_ = V4600
			V4601 := __args[2]
			_ = V4601
			gen18867 := Call(__e, ShenFunc(symshen_4write_1char_1and_1inc), V4599, V4600, V4601)

			__e.TailApply(ShenFunc(symshen_4prh), V4599, V4600, gen18867)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prh"), gen18868)

		gen18871 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4605 := __args[0]
			_ = V4605
			V4606 := __args[1]
			_ = V4606
			V4607 := __args[2]
			_ = V4607
			gen18869 := Call(__e, ShenFunc(sympos), V4605, V4607)

			gen18870 := Call(__e, ShenFunc(symstring_1_6n), gen18869)

			Call(__e, ShenFunc(symwrite_1byte), gen18870, V4606)

			__e.TailApply(ShenFunc(sym_7), V4607, MakeNumber(1))

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.write-char-and-inc"), gen18871)

		gen18875 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4609 := __args[0]
			_ = V4609
			gen18872 := Call(__e, ShenFunc(symshen_4insert), V4609, MakeString("~S"))

			String := gen18872
			gen18873 := Call(__e, ShenFunc(symstoutput))

			gen18874 := Call(__e, ShenFunc(symshen_4prhush), String, gen18873)

			Print := gen18874
			_ = Print
			__e.Return(V4609)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("print"), gen18875)

		gen18877 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4612 := __args[0]
			_ = V4612
			V4613 := __args[1]
			_ = V4613
			gen18876 := Call(__e, ShenFunc(symvalue), MakeSymbol("*hush*"))

			if True == gen18876 {
				__e.Return(V4612)
				return
			} else {
				__e.TailApply(ShenFunc(sympr), V4612, V4613)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prhush"), gen18877)

		gen18882 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4616 := __args[0]
			_ = V4616
			V4617 := __args[1]
			_ = V4617
			gen18881 := Call(__e, ShenFunc(symstring_2), V4616)

			if True == gen18881 {
				gen18880 := Call(__e, ShenFunc(symshen_4proc_1nl), V4616)

				__e.TailApply(ShenFunc(symshen_4mkstr_1l), gen18880, V4617)

				return

			} else {
				gen18878 := Call(__e, ShenFunc(symcons), V4616, Nil)

				gen18879 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.proc-nl"), gen18878)

				__e.TailApply(ShenFunc(symshen_4mkstr_1r), gen18879, V4617)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mkstr"), gen18882)

		gen18888 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4620 := __args[0]
			_ = V4620
			V4621 := __args[1]
			_ = V4621
			gen18887 := Call(__e, ShenFunc(sym_a), Nil, V4621)

			if True == gen18887 {
				__e.Return(V4620)
				return
			} else {
				gen18886 := Call(__e, ShenFunc(symcons_2), V4621)

				if True == gen18886 {
					gen18883 := Call(__e, ShenFunc(symhd), V4621)

					gen18884 := Call(__e, ShenFunc(symshen_4insert_1l), gen18883, V4620)

					gen18885 := Call(__e, ShenFunc(symtl), V4621)

					__e.TailApply(ShenFunc(symshen_4mkstr_1l), gen18884, gen18885)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.mkstr-l"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mkstr-l"), gen18888)

		gen19002 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4626 := __args[0]
			_ = V4626
			V4627 := __args[1]
			_ = V4627
			gen19001 := Call(__e, ShenFunc(sym_a), MakeString(""), V4627)

			if True == gen19001 {
				__e.Return(MakeString(""))
				return
			} else {
				gen18999 := Call(__e, ShenFunc(symshen_4_7string_2), V4627)

				var gen19000 Obj
				if True == gen18999 {
					gen18996 := Call(__e, ShenFunc(sympos), V4627, MakeNumber(0))

					gen18997 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen18996)

					var gen18998 Obj
					if True == gen18997 {
						gen18993 := Call(__e, ShenFunc(symtlstr), V4627)

						gen18994 := Call(__e, ShenFunc(symshen_4_7string_2), gen18993)

						var gen18995 Obj
						if True == gen18994 {
							gen18990 := Call(__e, ShenFunc(symtlstr), V4627)

							gen18991 := Call(__e, ShenFunc(sympos), gen18990, MakeNumber(0))

							gen18992 := Call(__e, ShenFunc(sym_a), MakeString("A"), gen18991)

							if True == gen18992 {
								gen18995 = True
							} else {
								gen18995 = False
							}

						} else {
							gen18995 = False
						}
						if True == gen18995 {
							gen18998 = True
						} else {
							gen18998 = False
						}

					} else {
						gen18998 = False
					}
					if True == gen18998 {
						gen19000 = True
					} else {
						gen19000 = False
					}

				} else {
					gen19000 = False
				}
				if True == gen19000 {
					gen18985 := Call(__e, ShenFunc(symtlstr), V4627)

					gen18986 := Call(__e, ShenFunc(symtlstr), gen18985)

					gen18987 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), Nil)

					gen18988 := Call(__e, ShenFunc(symcons), gen18986, gen18987)

					gen18989 := Call(__e, ShenFunc(symcons), V4626, gen18988)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.app"), gen18989)

					return

				} else {
					gen18983 := Call(__e, ShenFunc(symshen_4_7string_2), V4627)

					var gen18984 Obj
					if True == gen18983 {
						gen18980 := Call(__e, ShenFunc(sympos), V4627, MakeNumber(0))

						gen18981 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen18980)

						var gen18982 Obj
						if True == gen18981 {
							gen18977 := Call(__e, ShenFunc(symtlstr), V4627)

							gen18978 := Call(__e, ShenFunc(symshen_4_7string_2), gen18977)

							var gen18979 Obj
							if True == gen18978 {
								gen18974 := Call(__e, ShenFunc(symtlstr), V4627)

								gen18975 := Call(__e, ShenFunc(sympos), gen18974, MakeNumber(0))

								gen18976 := Call(__e, ShenFunc(sym_a), MakeString("R"), gen18975)

								if True == gen18976 {
									gen18979 = True
								} else {
									gen18979 = False
								}

							} else {
								gen18979 = False
							}
							if True == gen18979 {
								gen18982 = True
							} else {
								gen18982 = False
							}

						} else {
							gen18982 = False
						}
						if True == gen18982 {
							gen18984 = True
						} else {
							gen18984 = False
						}

					} else {
						gen18984 = False
					}
					if True == gen18984 {
						gen18969 := Call(__e, ShenFunc(symtlstr), V4627)

						gen18970 := Call(__e, ShenFunc(symtlstr), gen18969)

						gen18971 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.r"), Nil)

						gen18972 := Call(__e, ShenFunc(symcons), gen18970, gen18971)

						gen18973 := Call(__e, ShenFunc(symcons), V4626, gen18972)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.app"), gen18973)

						return

					} else {
						gen18967 := Call(__e, ShenFunc(symshen_4_7string_2), V4627)

						var gen18968 Obj
						if True == gen18967 {
							gen18964 := Call(__e, ShenFunc(sympos), V4627, MakeNumber(0))

							gen18965 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen18964)

							var gen18966 Obj
							if True == gen18965 {
								gen18961 := Call(__e, ShenFunc(symtlstr), V4627)

								gen18962 := Call(__e, ShenFunc(symshen_4_7string_2), gen18961)

								var gen18963 Obj
								if True == gen18962 {
									gen18958 := Call(__e, ShenFunc(symtlstr), V4627)

									gen18959 := Call(__e, ShenFunc(sympos), gen18958, MakeNumber(0))

									gen18960 := Call(__e, ShenFunc(sym_a), MakeString("S"), gen18959)

									if True == gen18960 {
										gen18963 = True
									} else {
										gen18963 = False
									}

								} else {
									gen18963 = False
								}
								if True == gen18963 {
									gen18966 = True
								} else {
									gen18966 = False
								}

							} else {
								gen18966 = False
							}
							if True == gen18966 {
								gen18968 = True
							} else {
								gen18968 = False
							}

						} else {
							gen18968 = False
						}
						if True == gen18968 {
							gen18953 := Call(__e, ShenFunc(symtlstr), V4627)

							gen18954 := Call(__e, ShenFunc(symtlstr), gen18953)

							gen18955 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.s"), Nil)

							gen18956 := Call(__e, ShenFunc(symcons), gen18954, gen18955)

							gen18957 := Call(__e, ShenFunc(symcons), V4626, gen18956)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.app"), gen18957)

							return

						} else {
							gen18952 := Call(__e, ShenFunc(symshen_4_7string_2), V4627)

							if True == gen18952 {
								gen18946 := Call(__e, ShenFunc(sympos), V4627, MakeNumber(0))

								gen18947 := Call(__e, ShenFunc(symtlstr), V4627)

								gen18948 := Call(__e, ShenFunc(symshen_4insert_1l), V4626, gen18947)

								gen18949 := Call(__e, ShenFunc(symcons), gen18948, Nil)

								gen18950 := Call(__e, ShenFunc(symcons), gen18946, gen18949)

								gen18951 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen18950)

								__e.TailApply(ShenFunc(symshen_4factor_1cn), gen18951)

								return

							} else {
								gen18944 := Call(__e, ShenFunc(symcons_2), V4627)

								var gen18945 Obj
								if True == gen18944 {
									gen18941 := Call(__e, ShenFunc(symhd), V4627)

									gen18942 := Call(__e, ShenFunc(sym_a), MakeSymbol("cn"), gen18941)

									var gen18943 Obj
									if True == gen18942 {
										gen18938 := Call(__e, ShenFunc(symtl), V4627)

										gen18939 := Call(__e, ShenFunc(symcons_2), gen18938)

										var gen18940 Obj
										if True == gen18939 {
											gen18934 := Call(__e, ShenFunc(symtl), V4627)

											gen18935 := Call(__e, ShenFunc(symtl), gen18934)

											gen18936 := Call(__e, ShenFunc(symcons_2), gen18935)

											var gen18937 Obj
											if True == gen18936 {
												gen18930 := Call(__e, ShenFunc(symtl), V4627)

												gen18931 := Call(__e, ShenFunc(symtl), gen18930)

												gen18932 := Call(__e, ShenFunc(symtl), gen18931)

												gen18933 := Call(__e, ShenFunc(sym_a), Nil, gen18932)

												if True == gen18933 {
													gen18937 = True
												} else {
													gen18937 = False
												}

											} else {
												gen18937 = False
											}
											if True == gen18937 {
												gen18940 = True
											} else {
												gen18940 = False
											}

										} else {
											gen18940 = False
										}
										if True == gen18940 {
											gen18943 = True
										} else {
											gen18943 = False
										}

									} else {
										gen18943 = False
									}
									if True == gen18943 {
										gen18945 = True
									} else {
										gen18945 = False
									}

								} else {
									gen18945 = False
								}
								if True == gen18945 {
									gen18922 := Call(__e, ShenFunc(symtl), V4627)

									gen18923 := Call(__e, ShenFunc(symhd), gen18922)

									gen18924 := Call(__e, ShenFunc(symtl), V4627)

									gen18925 := Call(__e, ShenFunc(symtl), gen18924)

									gen18926 := Call(__e, ShenFunc(symhd), gen18925)

									gen18927 := Call(__e, ShenFunc(symshen_4insert_1l), V4626, gen18926)

									gen18928 := Call(__e, ShenFunc(symcons), gen18927, Nil)

									gen18929 := Call(__e, ShenFunc(symcons), gen18923, gen18928)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("cn"), gen18929)

									return

								} else {
									gen18920 := Call(__e, ShenFunc(symcons_2), V4627)

									var gen18921 Obj
									if True == gen18920 {
										gen18917 := Call(__e, ShenFunc(symhd), V4627)

										gen18918 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.app"), gen18917)

										var gen18919 Obj
										if True == gen18918 {
											gen18914 := Call(__e, ShenFunc(symtl), V4627)

											gen18915 := Call(__e, ShenFunc(symcons_2), gen18914)

											var gen18916 Obj
											if True == gen18915 {
												gen18910 := Call(__e, ShenFunc(symtl), V4627)

												gen18911 := Call(__e, ShenFunc(symtl), gen18910)

												gen18912 := Call(__e, ShenFunc(symcons_2), gen18911)

												var gen18913 Obj
												if True == gen18912 {
													gen18905 := Call(__e, ShenFunc(symtl), V4627)

													gen18906 := Call(__e, ShenFunc(symtl), gen18905)

													gen18907 := Call(__e, ShenFunc(symtl), gen18906)

													gen18908 := Call(__e, ShenFunc(symcons_2), gen18907)

													var gen18909 Obj
													if True == gen18908 {
														gen18900 := Call(__e, ShenFunc(symtl), V4627)

														gen18901 := Call(__e, ShenFunc(symtl), gen18900)

														gen18902 := Call(__e, ShenFunc(symtl), gen18901)

														gen18903 := Call(__e, ShenFunc(symtl), gen18902)

														gen18904 := Call(__e, ShenFunc(sym_a), Nil, gen18903)

														if True == gen18904 {
															gen18909 = True
														} else {
															gen18909 = False
														}

													} else {
														gen18909 = False
													}
													if True == gen18909 {
														gen18913 = True
													} else {
														gen18913 = False
													}

												} else {
													gen18913 = False
												}
												if True == gen18913 {
													gen18916 = True
												} else {
													gen18916 = False
												}

											} else {
												gen18916 = False
											}
											if True == gen18916 {
												gen18919 = True
											} else {
												gen18919 = False
											}

										} else {
											gen18919 = False
										}
										if True == gen18919 {
											gen18921 = True
										} else {
											gen18921 = False
										}

									} else {
										gen18921 = False
									}
									if True == gen18921 {
										gen18889 := Call(__e, ShenFunc(symtl), V4627)

										gen18890 := Call(__e, ShenFunc(symhd), gen18889)

										gen18891 := Call(__e, ShenFunc(symtl), V4627)

										gen18892 := Call(__e, ShenFunc(symtl), gen18891)

										gen18893 := Call(__e, ShenFunc(symhd), gen18892)

										gen18894 := Call(__e, ShenFunc(symshen_4insert_1l), V4626, gen18893)

										gen18895 := Call(__e, ShenFunc(symtl), V4627)

										gen18896 := Call(__e, ShenFunc(symtl), gen18895)

										gen18897 := Call(__e, ShenFunc(symtl), gen18896)

										gen18898 := Call(__e, ShenFunc(symcons), gen18894, gen18897)

										gen18899 := Call(__e, ShenFunc(symcons), gen18890, gen18898)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.app"), gen18899)

										return

									} else {
										__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.insert-l"))

										return
									}

								}

							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-l"), gen19002)

		gen19076 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4629 := __args[0]
			_ = V4629
			gen19074 := Call(__e, ShenFunc(symcons_2), V4629)

			var gen19075 Obj
			if True == gen19074 {
				gen19071 := Call(__e, ShenFunc(symhd), V4629)

				gen19072 := Call(__e, ShenFunc(sym_a), MakeSymbol("cn"), gen19071)

				var gen19073 Obj
				if True == gen19072 {
					gen19068 := Call(__e, ShenFunc(symtl), V4629)

					gen19069 := Call(__e, ShenFunc(symcons_2), gen19068)

					var gen19070 Obj
					if True == gen19069 {
						gen19064 := Call(__e, ShenFunc(symtl), V4629)

						gen19065 := Call(__e, ShenFunc(symtl), gen19064)

						gen19066 := Call(__e, ShenFunc(symcons_2), gen19065)

						var gen19067 Obj
						if True == gen19066 {
							gen19059 := Call(__e, ShenFunc(symtl), V4629)

							gen19060 := Call(__e, ShenFunc(symtl), gen19059)

							gen19061 := Call(__e, ShenFunc(symhd), gen19060)

							gen19062 := Call(__e, ShenFunc(symcons_2), gen19061)

							var gen19063 Obj
							if True == gen19062 {
								gen19053 := Call(__e, ShenFunc(symtl), V4629)

								gen19054 := Call(__e, ShenFunc(symtl), gen19053)

								gen19055 := Call(__e, ShenFunc(symhd), gen19054)

								gen19056 := Call(__e, ShenFunc(symhd), gen19055)

								gen19057 := Call(__e, ShenFunc(sym_a), MakeSymbol("cn"), gen19056)

								var gen19058 Obj
								if True == gen19057 {
									gen19047 := Call(__e, ShenFunc(symtl), V4629)

									gen19048 := Call(__e, ShenFunc(symtl), gen19047)

									gen19049 := Call(__e, ShenFunc(symhd), gen19048)

									gen19050 := Call(__e, ShenFunc(symtl), gen19049)

									gen19051 := Call(__e, ShenFunc(symcons_2), gen19050)

									var gen19052 Obj
									if True == gen19051 {
										gen19040 := Call(__e, ShenFunc(symtl), V4629)

										gen19041 := Call(__e, ShenFunc(symtl), gen19040)

										gen19042 := Call(__e, ShenFunc(symhd), gen19041)

										gen19043 := Call(__e, ShenFunc(symtl), gen19042)

										gen19044 := Call(__e, ShenFunc(symtl), gen19043)

										gen19045 := Call(__e, ShenFunc(symcons_2), gen19044)

										var gen19046 Obj
										if True == gen19045 {
											gen19032 := Call(__e, ShenFunc(symtl), V4629)

											gen19033 := Call(__e, ShenFunc(symtl), gen19032)

											gen19034 := Call(__e, ShenFunc(symhd), gen19033)

											gen19035 := Call(__e, ShenFunc(symtl), gen19034)

											gen19036 := Call(__e, ShenFunc(symtl), gen19035)

											gen19037 := Call(__e, ShenFunc(symtl), gen19036)

											gen19038 := Call(__e, ShenFunc(sym_a), Nil, gen19037)

											var gen19039 Obj
											if True == gen19038 {
												gen19027 := Call(__e, ShenFunc(symtl), V4629)

												gen19028 := Call(__e, ShenFunc(symtl), gen19027)

												gen19029 := Call(__e, ShenFunc(symtl), gen19028)

												gen19030 := Call(__e, ShenFunc(sym_a), Nil, gen19029)

												var gen19031 Obj
												if True == gen19030 {
													gen19023 := Call(__e, ShenFunc(symtl), V4629)

													gen19024 := Call(__e, ShenFunc(symhd), gen19023)

													gen19025 := Call(__e, ShenFunc(symstring_2), gen19024)

													var gen19026 Obj
													if True == gen19025 {
														gen19017 := Call(__e, ShenFunc(symtl), V4629)

														gen19018 := Call(__e, ShenFunc(symtl), gen19017)

														gen19019 := Call(__e, ShenFunc(symhd), gen19018)

														gen19020 := Call(__e, ShenFunc(symtl), gen19019)

														gen19021 := Call(__e, ShenFunc(symhd), gen19020)

														gen19022 := Call(__e, ShenFunc(symstring_2), gen19021)

														if True == gen19022 {
															gen19026 = True
														} else {
															gen19026 = False
														}

													} else {
														gen19026 = False
													}
													if True == gen19026 {
														gen19031 = True
													} else {
														gen19031 = False
													}

												} else {
													gen19031 = False
												}
												if True == gen19031 {
													gen19039 = True
												} else {
													gen19039 = False
												}

											} else {
												gen19039 = False
											}
											if True == gen19039 {
												gen19046 = True
											} else {
												gen19046 = False
											}

										} else {
											gen19046 = False
										}
										if True == gen19046 {
											gen19052 = True
										} else {
											gen19052 = False
										}

									} else {
										gen19052 = False
									}
									if True == gen19052 {
										gen19058 = True
									} else {
										gen19058 = False
									}

								} else {
									gen19058 = False
								}
								if True == gen19058 {
									gen19063 = True
								} else {
									gen19063 = False
								}

							} else {
								gen19063 = False
							}
							if True == gen19063 {
								gen19067 = True
							} else {
								gen19067 = False
							}

						} else {
							gen19067 = False
						}
						if True == gen19067 {
							gen19070 = True
						} else {
							gen19070 = False
						}

					} else {
						gen19070 = False
					}
					if True == gen19070 {
						gen19073 = True
					} else {
						gen19073 = False
					}

				} else {
					gen19073 = False
				}
				if True == gen19073 {
					gen19075 = True
				} else {
					gen19075 = False
				}

			} else {
				gen19075 = False
			}
			if True == gen19075 {
				gen19003 := Call(__e, ShenFunc(symtl), V4629)

				gen19004 := Call(__e, ShenFunc(symhd), gen19003)

				gen19005 := Call(__e, ShenFunc(symtl), V4629)

				gen19006 := Call(__e, ShenFunc(symtl), gen19005)

				gen19007 := Call(__e, ShenFunc(symhd), gen19006)

				gen19008 := Call(__e, ShenFunc(symtl), gen19007)

				gen19009 := Call(__e, ShenFunc(symhd), gen19008)

				gen19010 := Call(__e, ShenFunc(symcn), gen19004, gen19009)

				gen19011 := Call(__e, ShenFunc(symtl), V4629)

				gen19012 := Call(__e, ShenFunc(symtl), gen19011)

				gen19013 := Call(__e, ShenFunc(symhd), gen19012)

				gen19014 := Call(__e, ShenFunc(symtl), gen19013)

				gen19015 := Call(__e, ShenFunc(symtl), gen19014)

				gen19016 := Call(__e, ShenFunc(symcons), gen19010, gen19015)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("cn"), gen19016)

				return

			} else {
				__e.Return(V4629)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.factor-cn"), gen19076)

		gen19097 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4631 := __args[0]
			_ = V4631
			gen19096 := Call(__e, ShenFunc(sym_a), MakeString(""), V4631)

			if True == gen19096 {
				__e.Return(MakeString(""))
				return
			} else {
				gen19094 := Call(__e, ShenFunc(symshen_4_7string_2), V4631)

				var gen19095 Obj
				if True == gen19094 {
					gen19091 := Call(__e, ShenFunc(sympos), V4631, MakeNumber(0))

					gen19092 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen19091)

					var gen19093 Obj
					if True == gen19092 {
						gen19088 := Call(__e, ShenFunc(symtlstr), V4631)

						gen19089 := Call(__e, ShenFunc(symshen_4_7string_2), gen19088)

						var gen19090 Obj
						if True == gen19089 {
							gen19085 := Call(__e, ShenFunc(symtlstr), V4631)

							gen19086 := Call(__e, ShenFunc(sympos), gen19085, MakeNumber(0))

							gen19087 := Call(__e, ShenFunc(sym_a), MakeString("%"), gen19086)

							if True == gen19087 {
								gen19090 = True
							} else {
								gen19090 = False
							}

						} else {
							gen19090 = False
						}
						if True == gen19090 {
							gen19093 = True
						} else {
							gen19093 = False
						}

					} else {
						gen19093 = False
					}
					if True == gen19093 {
						gen19095 = True
					} else {
						gen19095 = False
					}

				} else {
					gen19095 = False
				}
				if True == gen19095 {
					gen19081 := Call(__e, ShenFunc(symn_1_6string), MakeNumber(10))

					gen19082 := Call(__e, ShenFunc(symtlstr), V4631)

					gen19083 := Call(__e, ShenFunc(symtlstr), gen19082)

					gen19084 := Call(__e, ShenFunc(symshen_4proc_1nl), gen19083)

					__e.TailApply(ShenFunc(symcn), gen19081, gen19084)

					return

				} else {
					gen19080 := Call(__e, ShenFunc(symshen_4_7string_2), V4631)

					if True == gen19080 {
						gen19077 := Call(__e, ShenFunc(sympos), V4631, MakeNumber(0))

						gen19078 := Call(__e, ShenFunc(symtlstr), V4631)

						gen19079 := Call(__e, ShenFunc(symshen_4proc_1nl), gen19078)

						__e.TailApply(ShenFunc(symcn), gen19077, gen19079)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.proc-nl"))

						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.proc-nl"), gen19097)

		gen19105 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4634 := __args[0]
			_ = V4634
			V4635 := __args[1]
			_ = V4635
			gen19104 := Call(__e, ShenFunc(sym_a), Nil, V4635)

			if True == gen19104 {
				__e.Return(V4634)
				return
			} else {
				gen19103 := Call(__e, ShenFunc(symcons_2), V4635)

				if True == gen19103 {
					gen19098 := Call(__e, ShenFunc(symhd), V4635)

					gen19099 := Call(__e, ShenFunc(symcons), V4634, Nil)

					gen19100 := Call(__e, ShenFunc(symcons), gen19098, gen19099)

					gen19101 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.insert"), gen19100)

					gen19102 := Call(__e, ShenFunc(symtl), V4635)

					__e.TailApply(ShenFunc(symshen_4mkstr_1r), gen19101, gen19102)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.mkstr-r"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mkstr-r"), gen19105)

		gen19106 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4638 := __args[0]
			_ = V4638
			V4639 := __args[1]
			_ = V4639
			__e.TailApply(ShenFunc(symshen_4insert_1h), V4638, V4639, MakeString(""))

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert"), gen19106)

		gen19154 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4645 := __args[0]
			_ = V4645
			V4646 := __args[1]
			_ = V4646
			V4647 := __args[2]
			_ = V4647
			gen19153 := Call(__e, ShenFunc(sym_a), MakeString(""), V4646)

			if True == gen19153 {
				__e.Return(V4647)
				return
			} else {
				gen19151 := Call(__e, ShenFunc(symshen_4_7string_2), V4646)

				var gen19152 Obj
				if True == gen19151 {
					gen19148 := Call(__e, ShenFunc(sympos), V4646, MakeNumber(0))

					gen19149 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen19148)

					var gen19150 Obj
					if True == gen19149 {
						gen19145 := Call(__e, ShenFunc(symtlstr), V4646)

						gen19146 := Call(__e, ShenFunc(symshen_4_7string_2), gen19145)

						var gen19147 Obj
						if True == gen19146 {
							gen19142 := Call(__e, ShenFunc(symtlstr), V4646)

							gen19143 := Call(__e, ShenFunc(sympos), gen19142, MakeNumber(0))

							gen19144 := Call(__e, ShenFunc(sym_a), MakeString("A"), gen19143)

							if True == gen19144 {
								gen19147 = True
							} else {
								gen19147 = False
							}

						} else {
							gen19147 = False
						}
						if True == gen19147 {
							gen19150 = True
						} else {
							gen19150 = False
						}

					} else {
						gen19150 = False
					}
					if True == gen19150 {
						gen19152 = True
					} else {
						gen19152 = False
					}

				} else {
					gen19152 = False
				}
				if True == gen19152 {
					gen19139 := Call(__e, ShenFunc(symtlstr), V4646)

					gen19140 := Call(__e, ShenFunc(symtlstr), gen19139)

					gen19141 := Call(__e, ShenFunc(symshen_4app), V4645, gen19140, MakeSymbol("shen.a"))

					__e.TailApply(ShenFunc(symcn), V4647, gen19141)

					return

				} else {
					gen19137 := Call(__e, ShenFunc(symshen_4_7string_2), V4646)

					var gen19138 Obj
					if True == gen19137 {
						gen19134 := Call(__e, ShenFunc(sympos), V4646, MakeNumber(0))

						gen19135 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen19134)

						var gen19136 Obj
						if True == gen19135 {
							gen19131 := Call(__e, ShenFunc(symtlstr), V4646)

							gen19132 := Call(__e, ShenFunc(symshen_4_7string_2), gen19131)

							var gen19133 Obj
							if True == gen19132 {
								gen19128 := Call(__e, ShenFunc(symtlstr), V4646)

								gen19129 := Call(__e, ShenFunc(sympos), gen19128, MakeNumber(0))

								gen19130 := Call(__e, ShenFunc(sym_a), MakeString("R"), gen19129)

								if True == gen19130 {
									gen19133 = True
								} else {
									gen19133 = False
								}

							} else {
								gen19133 = False
							}
							if True == gen19133 {
								gen19136 = True
							} else {
								gen19136 = False
							}

						} else {
							gen19136 = False
						}
						if True == gen19136 {
							gen19138 = True
						} else {
							gen19138 = False
						}

					} else {
						gen19138 = False
					}
					if True == gen19138 {
						gen19125 := Call(__e, ShenFunc(symtlstr), V4646)

						gen19126 := Call(__e, ShenFunc(symtlstr), gen19125)

						gen19127 := Call(__e, ShenFunc(symshen_4app), V4645, gen19126, MakeSymbol("shen.r"))

						__e.TailApply(ShenFunc(symcn), V4647, gen19127)

						return

					} else {
						gen19123 := Call(__e, ShenFunc(symshen_4_7string_2), V4646)

						var gen19124 Obj
						if True == gen19123 {
							gen19120 := Call(__e, ShenFunc(sympos), V4646, MakeNumber(0))

							gen19121 := Call(__e, ShenFunc(sym_a), MakeString("~"), gen19120)

							var gen19122 Obj
							if True == gen19121 {
								gen19117 := Call(__e, ShenFunc(symtlstr), V4646)

								gen19118 := Call(__e, ShenFunc(symshen_4_7string_2), gen19117)

								var gen19119 Obj
								if True == gen19118 {
									gen19114 := Call(__e, ShenFunc(symtlstr), V4646)

									gen19115 := Call(__e, ShenFunc(sympos), gen19114, MakeNumber(0))

									gen19116 := Call(__e, ShenFunc(sym_a), MakeString("S"), gen19115)

									if True == gen19116 {
										gen19119 = True
									} else {
										gen19119 = False
									}

								} else {
									gen19119 = False
								}
								if True == gen19119 {
									gen19122 = True
								} else {
									gen19122 = False
								}

							} else {
								gen19122 = False
							}
							if True == gen19122 {
								gen19124 = True
							} else {
								gen19124 = False
							}

						} else {
							gen19124 = False
						}
						if True == gen19124 {
							gen19111 := Call(__e, ShenFunc(symtlstr), V4646)

							gen19112 := Call(__e, ShenFunc(symtlstr), gen19111)

							gen19113 := Call(__e, ShenFunc(symshen_4app), V4645, gen19112, MakeSymbol("shen.s"))

							__e.TailApply(ShenFunc(symcn), V4647, gen19113)

							return

						} else {
							gen19110 := Call(__e, ShenFunc(symshen_4_7string_2), V4646)

							if True == gen19110 {
								gen19107 := Call(__e, ShenFunc(symtlstr), V4646)

								gen19108 := Call(__e, ShenFunc(sympos), V4646, MakeNumber(0))

								gen19109 := Call(__e, ShenFunc(symcn), V4647, gen19108)

								__e.TailApply(ShenFunc(symshen_4insert_1h), V4645, gen19107, gen19109)

								return

							} else {
								__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.insert-h"))

								return
							}

						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-h"), gen19154)

		gen19156 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4651 := __args[0]
			_ = V4651
			V4652 := __args[1]
			_ = V4652
			V4653 := __args[2]
			_ = V4653
			gen19155 := Call(__e, ShenFunc(symshen_4arg_1_6str), V4651, V4653)

			__e.TailApply(ShenFunc(symcn), gen19155, V4652)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.app"), gen19156)

		gen19162 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4661 := __args[0]
			_ = V4661
			V4662 := __args[1]
			_ = V4662
			gen19160 := Call(__e, ShenFunc(symfail))

			gen19161 := Call(__e, ShenFunc(sym_a), V4661, gen19160)

			if True == gen19161 {
				__e.Return(MakeString("..."))
				return
			} else {
				gen19159 := Call(__e, ShenFunc(symshen_4list_2), V4661)

				if True == gen19159 {
					__e.TailApply(ShenFunc(symshen_4list_1_6str), V4661, V4662)

					return
				} else {
					gen19158 := Call(__e, ShenFunc(symstring_2), V4661)

					if True == gen19158 {
						__e.TailApply(ShenFunc(symshen_4str_1_6str), V4661, V4662)

						return
					} else {
						gen19157 := Call(__e, ShenFunc(symabsvector_2), V4661)

						if True == gen19157 {
							__e.TailApply(ShenFunc(symshen_4vector_1_6str), V4661, V4662)

							return
						} else {
							__e.TailApply(ShenFunc(symshen_4atom_1_6str), V4661)

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.arg->str"), gen19162)

		gen19170 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4665 := __args[0]
			_ = V4665
			V4666 := __args[1]
			_ = V4666
			gen19169 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.r"), V4666)

			if True == gen19169 {
				gen19166 := Call(__e, ShenFunc(symshen_4maxseq))

				gen19167 := Call(__e, ShenFunc(symshen_4iter_1list), V4665, MakeSymbol("shen.r"), gen19166)

				gen19168 := Call(__e, ShenFunc(sym_8s), gen19167, MakeString(")"))

				__e.TailApply(ShenFunc(sym_8s), MakeString("("), gen19168)

				return

			} else {
				gen19163 := Call(__e, ShenFunc(symshen_4maxseq))

				gen19164 := Call(__e, ShenFunc(symshen_4iter_1list), V4665, V4666, gen19163)

				gen19165 := Call(__e, ShenFunc(sym_8s), gen19164, MakeString("]"))

				__e.TailApply(ShenFunc(sym_8s), MakeString("["), gen19165)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.list->str"), gen19170)

		gen19171 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*maximum-print-sequence-size*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.maxseq"), gen19171)

		gen19188 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4680 := __args[0]
			_ = V4680
			V4681 := __args[1]
			_ = V4681
			V4682 := __args[2]
			_ = V4682
			gen19187 := Call(__e, ShenFunc(sym_a), Nil, V4680)

			if True == gen19187 {
				__e.Return(MakeString(""))
				return
			} else {
				gen19186 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V4682)

				if True == gen19186 {
					__e.Return(MakeString("... etc"))
					return
				} else {
					gen19184 := Call(__e, ShenFunc(symcons_2), V4680)

					var gen19185 Obj
					if True == gen19184 {
						gen19182 := Call(__e, ShenFunc(symtl), V4680)

						gen19183 := Call(__e, ShenFunc(sym_a), Nil, gen19182)

						if True == gen19183 {
							gen19185 = True
						} else {
							gen19185 = False
						}

					} else {
						gen19185 = False
					}
					if True == gen19185 {
						gen19181 := Call(__e, ShenFunc(symhd), V4680)

						__e.TailApply(ShenFunc(symshen_4arg_1_6str), gen19181, V4681)

						return

					} else {
						gen19180 := Call(__e, ShenFunc(symcons_2), V4680)

						if True == gen19180 {
							gen19174 := Call(__e, ShenFunc(symhd), V4680)

							gen19175 := Call(__e, ShenFunc(symshen_4arg_1_6str), gen19174, V4681)

							gen19176 := Call(__e, ShenFunc(symtl), V4680)

							gen19177 := Call(__e, ShenFunc(sym_1), V4682, MakeNumber(1))

							gen19178 := Call(__e, ShenFunc(symshen_4iter_1list), gen19176, V4681, gen19177)

							gen19179 := Call(__e, ShenFunc(sym_8s), MakeString(" "), gen19178)

							__e.TailApply(ShenFunc(sym_8s), gen19175, gen19179)

							return

						} else {
							gen19172 := Call(__e, ShenFunc(symshen_4arg_1_6str), V4680, V4681)

							gen19173 := Call(__e, ShenFunc(sym_8s), MakeString(" "), gen19172)

							__e.TailApply(ShenFunc(sym_8s), MakeString("|"), gen19173)

							return

						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.iter-list"), gen19188)

		gen19193 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4689 := __args[0]
			_ = V4689
			V4690 := __args[1]
			_ = V4690
			gen19192 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.a"), V4690)

			if True == gen19192 {
				__e.Return(V4689)
				return
			} else {
				gen19189 := Call(__e, ShenFunc(symn_1_6string), MakeNumber(34))

				gen19190 := Call(__e, ShenFunc(symn_1_6string), MakeNumber(34))

				gen19191 := Call(__e, ShenFunc(sym_8s), V4689, gen19190)

				__e.TailApply(ShenFunc(sym_8s), gen19189, gen19191)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.str->str"), gen19193)

		gen19205 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4693 := __args[0]
			_ = V4693
			V4694 := __args[1]
			_ = V4694
			gen19204 := Call(__e, ShenFunc(symshen_4print_1vector_2), V4693)

			if True == gen19204 {
				gen19202 := Call(__e, ShenFunc(sym_5_1address), V4693, MakeNumber(0))

				gen19203 := Call(__e, ShenFunc(symfunction), gen19202)

				f35 := gen19203
				__e.TailApply(f35, V4693)

				return

			} else {
				gen19201 := Call(__e, ShenFunc(symvector_2), V4693)

				if True == gen19201 {
					gen19198 := Call(__e, ShenFunc(symshen_4maxseq))

					gen19199 := Call(__e, ShenFunc(symshen_4iter_1vector), V4693, MakeNumber(1), V4694, gen19198)

					gen19200 := Call(__e, ShenFunc(sym_8s), gen19199, MakeString(">"))

					__e.TailApply(ShenFunc(sym_8s), MakeString("<"), gen19200)

					return

				} else {
					gen19194 := Call(__e, ShenFunc(symshen_4maxseq))

					gen19195 := Call(__e, ShenFunc(symshen_4iter_1vector), V4693, MakeNumber(0), V4694, gen19194)

					gen19196 := Call(__e, ShenFunc(sym_8s), gen19195, MakeString(">>"))

					gen19197 := Call(__e, ShenFunc(sym_8s), MakeString("<"), gen19196)

					__e.TailApply(ShenFunc(sym_8s), MakeString("<"), gen19197)

					return

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.vector->str"), gen19205)

		gen19207 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4696 := __args[0]
			_ = V4696
			gen19206 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*empty-absvector*"))

			__e.TailApply(ShenFunc(sym_a), V4696, gen19206)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.empty-absvector?"), gen19207)

		gen19221 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4698 := __args[0]
			_ = V4698
			gen19219 := Call(__e, ShenFunc(symshen_4empty_1absvector_2), V4698)

			gen19220 := Call(__e, ShenFunc(symnot), gen19219)

			if True == gen19220 {
				gen19208 := Call(__e, ShenFunc(sym_5_1address), V4698, MakeNumber(0))

				First := gen19208
				gen19217 := Call(__e, ShenFunc(sym_a), First, MakeSymbol("shen.tuple"))

				var gen19218 Obj
				if True == gen19217 {
					gen19218 = True
				} else {
					gen19215 := Call(__e, ShenFunc(sym_a), First, MakeSymbol("shen.pvar"))

					var gen19216 Obj
					if True == gen19215 {
						gen19216 = True
					} else {
						gen19213 := Call(__e, ShenFunc(sym_a), First, MakeSymbol("shen.dictionary"))

						var gen19214 Obj
						if True == gen19213 {
							gen19214 = True
						} else {
							gen19210 := Call(__e, ShenFunc(symnumber_2), First)

							gen19211 := Call(__e, ShenFunc(symnot), gen19210)

							var gen19212 Obj
							if True == gen19211 {
								gen19209 := Call(__e, ShenFunc(symshen_4fbound_2), First)

								if True == gen19209 {
									gen19212 = True
								} else {
									gen19212 = False
								}

							} else {
								gen19212 = False
							}
							if True == gen19212 {
								gen19214 = True
							} else {
								gen19214 = False
							}

						}
						if True == gen19214 {
							gen19216 = True
						} else {
							gen19216 = False
						}

					}
					if True == gen19216 {
						gen19218 = True
					} else {
						gen19218 = False
					}

				}
				if True == gen19218 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.print-vector?"), gen19221)

		gen19224 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4700 := __args[0]
			_ = V4700
			gen19222 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(False)
				return
			}, 1)
			gen19223 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Call(__e, ShenFunc(symshen_4lookup_1func), V4700)
				__e.Return(True)
				return

			}, 0)
			__e.Return(Try(__e, gen19223).Catch(gen19222))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fbound?"), gen19224)

		gen19230 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4702 := __args[0]
			_ = V4702
			gen19225 := Call(__e, ShenFunc(sym_5_1address), V4702, MakeNumber(1))

			gen19226 := Call(__e, ShenFunc(sym_5_1address), V4702, MakeNumber(2))

			gen19227 := Call(__e, ShenFunc(symshen_4app), gen19226, MakeString(")"), MakeSymbol("shen.s"))

			gen19228 := Call(__e, ShenFunc(symcn), MakeString(" "), gen19227)

			gen19229 := Call(__e, ShenFunc(symshen_4app), gen19225, gen19228, MakeSymbol("shen.s"))

			__e.TailApply(ShenFunc(symcn), MakeString("(@p "), gen19229)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tuple"), gen19230)

		gen19231 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4704 := __args[0]
			_ = V4704
			__e.Return(MakeString("(dict ...)"))
			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dictionary"), gen19231)

		gen19247 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4715 := __args[0]
			_ = V4715
			V4716 := __args[1]
			_ = V4716
			V4717 := __args[2]
			_ = V4717
			V4718 := __args[3]
			_ = V4718
			gen19246 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V4718)

			if True == gen19246 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				gen19232 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeSymbol("shen.out-of-bounds"))
					return
				}, 1)
				gen19233 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(sym_5_1address), V4715, V4716)

					return
				}, 0)
				gen19234 := Try(__e, gen19233).Catch(gen19232)
				Item := gen19234
				gen19235 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeSymbol("shen.out-of-bounds"))
					return
				}, 1)
				gen19237 := MakeNative(func(__e Evaluator, __args ...Obj) {
					gen19236 := Call(__e, ShenFunc(sym_7), V4716, MakeNumber(1))

					__e.TailApply(ShenFunc(sym_5_1address), V4715, gen19236)

					return

				}, 0)
				gen19238 := Try(__e, gen19237).Catch(gen19235)
				Next := gen19238
				gen19245 := Call(__e, ShenFunc(sym_a), Item, MakeSymbol("shen.out-of-bounds"))

				if True == gen19245 {
					__e.Return(MakeString(""))
					return
				} else {
					gen19244 := Call(__e, ShenFunc(sym_a), Next, MakeSymbol("shen.out-of-bounds"))

					if True == gen19244 {
						__e.TailApply(ShenFunc(symshen_4arg_1_6str), Item, V4717)

						return
					} else {
						gen19239 := Call(__e, ShenFunc(symshen_4arg_1_6str), Item, V4717)

						gen19240 := Call(__e, ShenFunc(sym_7), V4716, MakeNumber(1))

						gen19241 := Call(__e, ShenFunc(sym_1), V4718, MakeNumber(1))

						gen19242 := Call(__e, ShenFunc(symshen_4iter_1vector), V4715, gen19240, V4717, gen19241)

						gen19243 := Call(__e, ShenFunc(sym_8s), MakeString(" "), gen19242)

						__e.TailApply(ShenFunc(sym_8s), gen19239, gen19243)

						return

					}

				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.iter-vector"), gen19247)

		gen19250 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4720 := __args[0]
			_ = V4720
			gen19248 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.TailApply(ShenFunc(symshen_4funexstring))

				return
			}, 1)
			gen19249 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symstr), V4720)

				return
			}, 0)
			__e.Return(Try(__e, gen19249).Catch(gen19248))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.atom->str"), gen19250)

		gen19259 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen19251 := Call(__e, ShenFunc(symintern), MakeString("x"))

			gen19252 := Call(__e, ShenFunc(symgensym), gen19251)

			gen19253 := Call(__e, ShenFunc(symshen_4arg_1_6str), gen19252, MakeSymbol("shen.a"))

			gen19254 := Call(__e, ShenFunc(sym_8s), gen19253, MakeString("\x11"))

			gen19255 := Call(__e, ShenFunc(sym_8s), MakeString("e"), gen19254)

			gen19256 := Call(__e, ShenFunc(sym_8s), MakeString("n"), gen19255)

			gen19257 := Call(__e, ShenFunc(sym_8s), MakeString("u"), gen19256)

			gen19258 := Call(__e, ShenFunc(sym_8s), MakeString("f"), gen19257)

			__e.TailApply(ShenFunc(sym_8s), MakeString("\x10"), gen19258)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.funexstring"), gen19259)

		gen19262 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4722 := __args[0]
			_ = V4722
			gen19261 := Call(__e, ShenFunc(symempty_2), V4722)

			if True == gen19261 {
				__e.Return(True)
				return
			} else {
				gen19260 := Call(__e, ShenFunc(symcons_2), V4722)

				if True == gen19260 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.list?"), gen19262)

		return

	}, 0))
}
