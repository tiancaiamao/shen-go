package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen9681 := MakeNative(func(__e Evaluator) {
			V4594 := __e.Get(1)
			_ = V4594
			V4595 := __e.Get(2)
			_ = V4595
			gen9679 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(V4594)
				return
			}, 1)
			gen9680 := MakeNative(func(__e Evaluator) {
				__e.TailApply(__e.Global(symshen_4prh), V4594, V4595, MakeNumber(0))

				return
			}, 0)
			__e.Return(Try(__e, gen9680).Catch(gen9679))
			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("pr"), gen9681)

		gen9683 := MakeNative(func(__e Evaluator) {
			V4599 := __e.Get(1)
			_ = V4599
			V4600 := __e.Get(2)
			_ = V4600
			V4601 := __e.Get(3)
			_ = V4601
			gen9682 := Call(__e, __e.Global(symshen_4write_1char_1and_1inc), V4599, V4600, V4601)

			__e.TailApply(__e.Global(symshen_4prh), V4599, V4600, gen9682)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.prh"), gen9683)

		gen9686 := MakeNative(func(__e Evaluator) {
			V4605 := __e.Get(1)
			_ = V4605
			V4606 := __e.Get(2)
			_ = V4606
			V4607 := __e.Get(3)
			_ = V4607
			gen9684 := Call(__e, __e.Global(sympos), V4605, V4607)

			gen9685 := Call(__e, __e.Global(symstring_1_6n), gen9684)

			Call(__e, __e.Global(symwrite_1byte), gen9685, V4606)

			__e.TailApply(__e.Global(sym_7), V4607, MakeNumber(1))

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.write-char-and-inc"), gen9686)

		gen9690 := MakeNative(func(__e Evaluator) {
			V4609 := __e.Get(1)
			_ = V4609
			gen9687 := Call(__e, __e.Global(symshen_4insert), V4609, MakeString("~S"))

			String := gen9687
			gen9688 := Call(__e, __e.Global(symstoutput))

			gen9689 := Call(__e, __e.Global(symshen_4prhush), String, gen9688)

			Print := gen9689
			_ = Print
			__e.Return(V4609)
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("print"), gen9690)

		gen9692 := MakeNative(func(__e Evaluator) {
			V4612 := __e.Get(1)
			_ = V4612
			V4613 := __e.Get(2)
			_ = V4613
			gen9691 := Call(__e, __e.Global(symvalue), MakeSymbol("*hush*"))

			if True == gen9691 {
				__e.Return(V4612)
				return
			} else {
				__e.TailApply(__e.Global(sympr), V4612, V4613)

				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.prhush"), gen9692)

		gen9697 := MakeNative(func(__e Evaluator) {
			V4616 := __e.Get(1)
			_ = V4616
			V4617 := __e.Get(2)
			_ = V4617
			gen9696 := Call(__e, __e.Global(symstring_2), V4616)

			if True == gen9696 {
				gen9695 := Call(__e, __e.Global(symshen_4proc_1nl), V4616)

				__e.TailApply(__e.Global(symshen_4mkstr_1l), gen9695, V4617)

				return

			} else {
				gen9693 := Call(__e, __e.Global(symcons), V4616, Nil)

				gen9694 := Call(__e, __e.Global(symcons), MakeSymbol("shen.proc-nl"), gen9693)

				__e.TailApply(__e.Global(symshen_4mkstr_1r), gen9694, V4617)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.mkstr"), gen9697)

		gen9703 := MakeNative(func(__e Evaluator) {
			V4620 := __e.Get(1)
			_ = V4620
			V4621 := __e.Get(2)
			_ = V4621
			gen9702 := Call(__e, __e.Global(sym_a), Nil, V4621)

			if True == gen9702 {
				__e.Return(V4620)
				return
			} else {
				gen9701 := Call(__e, __e.Global(symcons_2), V4621)

				if True == gen9701 {
					gen9698 := Call(__e, __e.Global(symhd), V4621)

					gen9699 := Call(__e, __e.Global(symshen_4insert_1l), gen9698, V4620)

					gen9700 := Call(__e, __e.Global(symtl), V4621)

					__e.TailApply(__e.Global(symshen_4mkstr_1l), gen9699, gen9700)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.mkstr-l"))

					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.mkstr-l"), gen9703)

		gen9817 := MakeNative(func(__e Evaluator) {
			V4626 := __e.Get(1)
			_ = V4626
			V4627 := __e.Get(2)
			_ = V4627
			gen9816 := Call(__e, __e.Global(sym_a), MakeString(""), V4627)

			if True == gen9816 {
				__e.Return(MakeString(""))
				return
			} else {
				gen9814 := Call(__e, __e.Global(symshen_4_7string_2), V4627)

				var gen9815 Obj
				if True == gen9814 {
					gen9811 := Call(__e, __e.Global(sympos), V4627, MakeNumber(0))

					gen9812 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9811)

					var gen9813 Obj
					if True == gen9812 {
						gen9808 := Call(__e, __e.Global(symtlstr), V4627)

						gen9809 := Call(__e, __e.Global(symshen_4_7string_2), gen9808)

						var gen9810 Obj
						if True == gen9809 {
							gen9805 := Call(__e, __e.Global(symtlstr), V4627)

							gen9806 := Call(__e, __e.Global(sympos), gen9805, MakeNumber(0))

							gen9807 := Call(__e, __e.Global(sym_a), MakeString("A"), gen9806)

							if True == gen9807 {
								gen9810 = True
							} else {
								gen9810 = False
							}

						} else {
							gen9810 = False
						}
						if True == gen9810 {
							gen9813 = True
						} else {
							gen9813 = False
						}

					} else {
						gen9813 = False
					}
					if True == gen9813 {
						gen9815 = True
					} else {
						gen9815 = False
					}

				} else {
					gen9815 = False
				}
				if True == gen9815 {
					gen9800 := Call(__e, __e.Global(symtlstr), V4627)

					gen9801 := Call(__e, __e.Global(symtlstr), gen9800)

					gen9802 := Call(__e, __e.Global(symcons), MakeSymbol("shen.a"), Nil)

					gen9803 := Call(__e, __e.Global(symcons), gen9801, gen9802)

					gen9804 := Call(__e, __e.Global(symcons), V4626, gen9803)

					__e.TailApply(__e.Global(symcons), MakeSymbol("shen.app"), gen9804)

					return

				} else {
					gen9798 := Call(__e, __e.Global(symshen_4_7string_2), V4627)

					var gen9799 Obj
					if True == gen9798 {
						gen9795 := Call(__e, __e.Global(sympos), V4627, MakeNumber(0))

						gen9796 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9795)

						var gen9797 Obj
						if True == gen9796 {
							gen9792 := Call(__e, __e.Global(symtlstr), V4627)

							gen9793 := Call(__e, __e.Global(symshen_4_7string_2), gen9792)

							var gen9794 Obj
							if True == gen9793 {
								gen9789 := Call(__e, __e.Global(symtlstr), V4627)

								gen9790 := Call(__e, __e.Global(sympos), gen9789, MakeNumber(0))

								gen9791 := Call(__e, __e.Global(sym_a), MakeString("R"), gen9790)

								if True == gen9791 {
									gen9794 = True
								} else {
									gen9794 = False
								}

							} else {
								gen9794 = False
							}
							if True == gen9794 {
								gen9797 = True
							} else {
								gen9797 = False
							}

						} else {
							gen9797 = False
						}
						if True == gen9797 {
							gen9799 = True
						} else {
							gen9799 = False
						}

					} else {
						gen9799 = False
					}
					if True == gen9799 {
						gen9784 := Call(__e, __e.Global(symtlstr), V4627)

						gen9785 := Call(__e, __e.Global(symtlstr), gen9784)

						gen9786 := Call(__e, __e.Global(symcons), MakeSymbol("shen.r"), Nil)

						gen9787 := Call(__e, __e.Global(symcons), gen9785, gen9786)

						gen9788 := Call(__e, __e.Global(symcons), V4626, gen9787)

						__e.TailApply(__e.Global(symcons), MakeSymbol("shen.app"), gen9788)

						return

					} else {
						gen9782 := Call(__e, __e.Global(symshen_4_7string_2), V4627)

						var gen9783 Obj
						if True == gen9782 {
							gen9779 := Call(__e, __e.Global(sympos), V4627, MakeNumber(0))

							gen9780 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9779)

							var gen9781 Obj
							if True == gen9780 {
								gen9776 := Call(__e, __e.Global(symtlstr), V4627)

								gen9777 := Call(__e, __e.Global(symshen_4_7string_2), gen9776)

								var gen9778 Obj
								if True == gen9777 {
									gen9773 := Call(__e, __e.Global(symtlstr), V4627)

									gen9774 := Call(__e, __e.Global(sympos), gen9773, MakeNumber(0))

									gen9775 := Call(__e, __e.Global(sym_a), MakeString("S"), gen9774)

									if True == gen9775 {
										gen9778 = True
									} else {
										gen9778 = False
									}

								} else {
									gen9778 = False
								}
								if True == gen9778 {
									gen9781 = True
								} else {
									gen9781 = False
								}

							} else {
								gen9781 = False
							}
							if True == gen9781 {
								gen9783 = True
							} else {
								gen9783 = False
							}

						} else {
							gen9783 = False
						}
						if True == gen9783 {
							gen9768 := Call(__e, __e.Global(symtlstr), V4627)

							gen9769 := Call(__e, __e.Global(symtlstr), gen9768)

							gen9770 := Call(__e, __e.Global(symcons), MakeSymbol("shen.s"), Nil)

							gen9771 := Call(__e, __e.Global(symcons), gen9769, gen9770)

							gen9772 := Call(__e, __e.Global(symcons), V4626, gen9771)

							__e.TailApply(__e.Global(symcons), MakeSymbol("shen.app"), gen9772)

							return

						} else {
							gen9767 := Call(__e, __e.Global(symshen_4_7string_2), V4627)

							if True == gen9767 {
								gen9761 := Call(__e, __e.Global(sympos), V4627, MakeNumber(0))

								gen9762 := Call(__e, __e.Global(symtlstr), V4627)

								gen9763 := Call(__e, __e.Global(symshen_4insert_1l), V4626, gen9762)

								gen9764 := Call(__e, __e.Global(symcons), gen9763, Nil)

								gen9765 := Call(__e, __e.Global(symcons), gen9761, gen9764)

								gen9766 := Call(__e, __e.Global(symcons), MakeSymbol("cn"), gen9765)

								__e.TailApply(__e.Global(symshen_4factor_1cn), gen9766)

								return

							} else {
								gen9759 := Call(__e, __e.Global(symcons_2), V4627)

								var gen9760 Obj
								if True == gen9759 {
									gen9756 := Call(__e, __e.Global(symhd), V4627)

									gen9757 := Call(__e, __e.Global(sym_a), MakeSymbol("cn"), gen9756)

									var gen9758 Obj
									if True == gen9757 {
										gen9753 := Call(__e, __e.Global(symtl), V4627)

										gen9754 := Call(__e, __e.Global(symcons_2), gen9753)

										var gen9755 Obj
										if True == gen9754 {
											gen9749 := Call(__e, __e.Global(symtl), V4627)

											gen9750 := Call(__e, __e.Global(symtl), gen9749)

											gen9751 := Call(__e, __e.Global(symcons_2), gen9750)

											var gen9752 Obj
											if True == gen9751 {
												gen9745 := Call(__e, __e.Global(symtl), V4627)

												gen9746 := Call(__e, __e.Global(symtl), gen9745)

												gen9747 := Call(__e, __e.Global(symtl), gen9746)

												gen9748 := Call(__e, __e.Global(sym_a), Nil, gen9747)

												if True == gen9748 {
													gen9752 = True
												} else {
													gen9752 = False
												}

											} else {
												gen9752 = False
											}
											if True == gen9752 {
												gen9755 = True
											} else {
												gen9755 = False
											}

										} else {
											gen9755 = False
										}
										if True == gen9755 {
											gen9758 = True
										} else {
											gen9758 = False
										}

									} else {
										gen9758 = False
									}
									if True == gen9758 {
										gen9760 = True
									} else {
										gen9760 = False
									}

								} else {
									gen9760 = False
								}
								if True == gen9760 {
									gen9737 := Call(__e, __e.Global(symtl), V4627)

									gen9738 := Call(__e, __e.Global(symhd), gen9737)

									gen9739 := Call(__e, __e.Global(symtl), V4627)

									gen9740 := Call(__e, __e.Global(symtl), gen9739)

									gen9741 := Call(__e, __e.Global(symhd), gen9740)

									gen9742 := Call(__e, __e.Global(symshen_4insert_1l), V4626, gen9741)

									gen9743 := Call(__e, __e.Global(symcons), gen9742, Nil)

									gen9744 := Call(__e, __e.Global(symcons), gen9738, gen9743)

									__e.TailApply(__e.Global(symcons), MakeSymbol("cn"), gen9744)

									return

								} else {
									gen9735 := Call(__e, __e.Global(symcons_2), V4627)

									var gen9736 Obj
									if True == gen9735 {
										gen9732 := Call(__e, __e.Global(symhd), V4627)

										gen9733 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.app"), gen9732)

										var gen9734 Obj
										if True == gen9733 {
											gen9729 := Call(__e, __e.Global(symtl), V4627)

											gen9730 := Call(__e, __e.Global(symcons_2), gen9729)

											var gen9731 Obj
											if True == gen9730 {
												gen9725 := Call(__e, __e.Global(symtl), V4627)

												gen9726 := Call(__e, __e.Global(symtl), gen9725)

												gen9727 := Call(__e, __e.Global(symcons_2), gen9726)

												var gen9728 Obj
												if True == gen9727 {
													gen9720 := Call(__e, __e.Global(symtl), V4627)

													gen9721 := Call(__e, __e.Global(symtl), gen9720)

													gen9722 := Call(__e, __e.Global(symtl), gen9721)

													gen9723 := Call(__e, __e.Global(symcons_2), gen9722)

													var gen9724 Obj
													if True == gen9723 {
														gen9715 := Call(__e, __e.Global(symtl), V4627)

														gen9716 := Call(__e, __e.Global(symtl), gen9715)

														gen9717 := Call(__e, __e.Global(symtl), gen9716)

														gen9718 := Call(__e, __e.Global(symtl), gen9717)

														gen9719 := Call(__e, __e.Global(sym_a), Nil, gen9718)

														if True == gen9719 {
															gen9724 = True
														} else {
															gen9724 = False
														}

													} else {
														gen9724 = False
													}
													if True == gen9724 {
														gen9728 = True
													} else {
														gen9728 = False
													}

												} else {
													gen9728 = False
												}
												if True == gen9728 {
													gen9731 = True
												} else {
													gen9731 = False
												}

											} else {
												gen9731 = False
											}
											if True == gen9731 {
												gen9734 = True
											} else {
												gen9734 = False
											}

										} else {
											gen9734 = False
										}
										if True == gen9734 {
											gen9736 = True
										} else {
											gen9736 = False
										}

									} else {
										gen9736 = False
									}
									if True == gen9736 {
										gen9704 := Call(__e, __e.Global(symtl), V4627)

										gen9705 := Call(__e, __e.Global(symhd), gen9704)

										gen9706 := Call(__e, __e.Global(symtl), V4627)

										gen9707 := Call(__e, __e.Global(symtl), gen9706)

										gen9708 := Call(__e, __e.Global(symhd), gen9707)

										gen9709 := Call(__e, __e.Global(symshen_4insert_1l), V4626, gen9708)

										gen9710 := Call(__e, __e.Global(symtl), V4627)

										gen9711 := Call(__e, __e.Global(symtl), gen9710)

										gen9712 := Call(__e, __e.Global(symtl), gen9711)

										gen9713 := Call(__e, __e.Global(symcons), gen9709, gen9712)

										gen9714 := Call(__e, __e.Global(symcons), gen9705, gen9713)

										__e.TailApply(__e.Global(symcons), MakeSymbol("shen.app"), gen9714)

										return

									} else {
										__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.insert-l"))

										return
									}

								}

							}

						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.insert-l"), gen9817)

		gen9891 := MakeNative(func(__e Evaluator) {
			V4629 := __e.Get(1)
			_ = V4629
			gen9889 := Call(__e, __e.Global(symcons_2), V4629)

			var gen9890 Obj
			if True == gen9889 {
				gen9886 := Call(__e, __e.Global(symhd), V4629)

				gen9887 := Call(__e, __e.Global(sym_a), MakeSymbol("cn"), gen9886)

				var gen9888 Obj
				if True == gen9887 {
					gen9883 := Call(__e, __e.Global(symtl), V4629)

					gen9884 := Call(__e, __e.Global(symcons_2), gen9883)

					var gen9885 Obj
					if True == gen9884 {
						gen9879 := Call(__e, __e.Global(symtl), V4629)

						gen9880 := Call(__e, __e.Global(symtl), gen9879)

						gen9881 := Call(__e, __e.Global(symcons_2), gen9880)

						var gen9882 Obj
						if True == gen9881 {
							gen9874 := Call(__e, __e.Global(symtl), V4629)

							gen9875 := Call(__e, __e.Global(symtl), gen9874)

							gen9876 := Call(__e, __e.Global(symhd), gen9875)

							gen9877 := Call(__e, __e.Global(symcons_2), gen9876)

							var gen9878 Obj
							if True == gen9877 {
								gen9868 := Call(__e, __e.Global(symtl), V4629)

								gen9869 := Call(__e, __e.Global(symtl), gen9868)

								gen9870 := Call(__e, __e.Global(symhd), gen9869)

								gen9871 := Call(__e, __e.Global(symhd), gen9870)

								gen9872 := Call(__e, __e.Global(sym_a), MakeSymbol("cn"), gen9871)

								var gen9873 Obj
								if True == gen9872 {
									gen9862 := Call(__e, __e.Global(symtl), V4629)

									gen9863 := Call(__e, __e.Global(symtl), gen9862)

									gen9864 := Call(__e, __e.Global(symhd), gen9863)

									gen9865 := Call(__e, __e.Global(symtl), gen9864)

									gen9866 := Call(__e, __e.Global(symcons_2), gen9865)

									var gen9867 Obj
									if True == gen9866 {
										gen9855 := Call(__e, __e.Global(symtl), V4629)

										gen9856 := Call(__e, __e.Global(symtl), gen9855)

										gen9857 := Call(__e, __e.Global(symhd), gen9856)

										gen9858 := Call(__e, __e.Global(symtl), gen9857)

										gen9859 := Call(__e, __e.Global(symtl), gen9858)

										gen9860 := Call(__e, __e.Global(symcons_2), gen9859)

										var gen9861 Obj
										if True == gen9860 {
											gen9847 := Call(__e, __e.Global(symtl), V4629)

											gen9848 := Call(__e, __e.Global(symtl), gen9847)

											gen9849 := Call(__e, __e.Global(symhd), gen9848)

											gen9850 := Call(__e, __e.Global(symtl), gen9849)

											gen9851 := Call(__e, __e.Global(symtl), gen9850)

											gen9852 := Call(__e, __e.Global(symtl), gen9851)

											gen9853 := Call(__e, __e.Global(sym_a), Nil, gen9852)

											var gen9854 Obj
											if True == gen9853 {
												gen9842 := Call(__e, __e.Global(symtl), V4629)

												gen9843 := Call(__e, __e.Global(symtl), gen9842)

												gen9844 := Call(__e, __e.Global(symtl), gen9843)

												gen9845 := Call(__e, __e.Global(sym_a), Nil, gen9844)

												var gen9846 Obj
												if True == gen9845 {
													gen9838 := Call(__e, __e.Global(symtl), V4629)

													gen9839 := Call(__e, __e.Global(symhd), gen9838)

													gen9840 := Call(__e, __e.Global(symstring_2), gen9839)

													var gen9841 Obj
													if True == gen9840 {
														gen9832 := Call(__e, __e.Global(symtl), V4629)

														gen9833 := Call(__e, __e.Global(symtl), gen9832)

														gen9834 := Call(__e, __e.Global(symhd), gen9833)

														gen9835 := Call(__e, __e.Global(symtl), gen9834)

														gen9836 := Call(__e, __e.Global(symhd), gen9835)

														gen9837 := Call(__e, __e.Global(symstring_2), gen9836)

														if True == gen9837 {
															gen9841 = True
														} else {
															gen9841 = False
														}

													} else {
														gen9841 = False
													}
													if True == gen9841 {
														gen9846 = True
													} else {
														gen9846 = False
													}

												} else {
													gen9846 = False
												}
												if True == gen9846 {
													gen9854 = True
												} else {
													gen9854 = False
												}

											} else {
												gen9854 = False
											}
											if True == gen9854 {
												gen9861 = True
											} else {
												gen9861 = False
											}

										} else {
											gen9861 = False
										}
										if True == gen9861 {
											gen9867 = True
										} else {
											gen9867 = False
										}

									} else {
										gen9867 = False
									}
									if True == gen9867 {
										gen9873 = True
									} else {
										gen9873 = False
									}

								} else {
									gen9873 = False
								}
								if True == gen9873 {
									gen9878 = True
								} else {
									gen9878 = False
								}

							} else {
								gen9878 = False
							}
							if True == gen9878 {
								gen9882 = True
							} else {
								gen9882 = False
							}

						} else {
							gen9882 = False
						}
						if True == gen9882 {
							gen9885 = True
						} else {
							gen9885 = False
						}

					} else {
						gen9885 = False
					}
					if True == gen9885 {
						gen9888 = True
					} else {
						gen9888 = False
					}

				} else {
					gen9888 = False
				}
				if True == gen9888 {
					gen9890 = True
				} else {
					gen9890 = False
				}

			} else {
				gen9890 = False
			}
			if True == gen9890 {
				gen9818 := Call(__e, __e.Global(symtl), V4629)

				gen9819 := Call(__e, __e.Global(symhd), gen9818)

				gen9820 := Call(__e, __e.Global(symtl), V4629)

				gen9821 := Call(__e, __e.Global(symtl), gen9820)

				gen9822 := Call(__e, __e.Global(symhd), gen9821)

				gen9823 := Call(__e, __e.Global(symtl), gen9822)

				gen9824 := Call(__e, __e.Global(symhd), gen9823)

				gen9825 := Call(__e, __e.Global(symcn), gen9819, gen9824)

				gen9826 := Call(__e, __e.Global(symtl), V4629)

				gen9827 := Call(__e, __e.Global(symtl), gen9826)

				gen9828 := Call(__e, __e.Global(symhd), gen9827)

				gen9829 := Call(__e, __e.Global(symtl), gen9828)

				gen9830 := Call(__e, __e.Global(symtl), gen9829)

				gen9831 := Call(__e, __e.Global(symcons), gen9825, gen9830)

				__e.TailApply(__e.Global(symcons), MakeSymbol("cn"), gen9831)

				return

			} else {
				__e.Return(V4629)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.factor-cn"), gen9891)

		gen9912 := MakeNative(func(__e Evaluator) {
			V4631 := __e.Get(1)
			_ = V4631
			gen9911 := Call(__e, __e.Global(sym_a), MakeString(""), V4631)

			if True == gen9911 {
				__e.Return(MakeString(""))
				return
			} else {
				gen9909 := Call(__e, __e.Global(symshen_4_7string_2), V4631)

				var gen9910 Obj
				if True == gen9909 {
					gen9906 := Call(__e, __e.Global(sympos), V4631, MakeNumber(0))

					gen9907 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9906)

					var gen9908 Obj
					if True == gen9907 {
						gen9903 := Call(__e, __e.Global(symtlstr), V4631)

						gen9904 := Call(__e, __e.Global(symshen_4_7string_2), gen9903)

						var gen9905 Obj
						if True == gen9904 {
							gen9900 := Call(__e, __e.Global(symtlstr), V4631)

							gen9901 := Call(__e, __e.Global(sympos), gen9900, MakeNumber(0))

							gen9902 := Call(__e, __e.Global(sym_a), MakeString("%"), gen9901)

							if True == gen9902 {
								gen9905 = True
							} else {
								gen9905 = False
							}

						} else {
							gen9905 = False
						}
						if True == gen9905 {
							gen9908 = True
						} else {
							gen9908 = False
						}

					} else {
						gen9908 = False
					}
					if True == gen9908 {
						gen9910 = True
					} else {
						gen9910 = False
					}

				} else {
					gen9910 = False
				}
				if True == gen9910 {
					gen9896 := Call(__e, __e.Global(symn_1_6string), MakeNumber(10))

					gen9897 := Call(__e, __e.Global(symtlstr), V4631)

					gen9898 := Call(__e, __e.Global(symtlstr), gen9897)

					gen9899 := Call(__e, __e.Global(symshen_4proc_1nl), gen9898)

					__e.TailApply(__e.Global(symcn), gen9896, gen9899)

					return

				} else {
					gen9895 := Call(__e, __e.Global(symshen_4_7string_2), V4631)

					if True == gen9895 {
						gen9892 := Call(__e, __e.Global(sympos), V4631, MakeNumber(0))

						gen9893 := Call(__e, __e.Global(symtlstr), V4631)

						gen9894 := Call(__e, __e.Global(symshen_4proc_1nl), gen9893)

						__e.TailApply(__e.Global(symcn), gen9892, gen9894)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.proc-nl"))

						return
					}

				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.proc-nl"), gen9912)

		gen9920 := MakeNative(func(__e Evaluator) {
			V4634 := __e.Get(1)
			_ = V4634
			V4635 := __e.Get(2)
			_ = V4635
			gen9919 := Call(__e, __e.Global(sym_a), Nil, V4635)

			if True == gen9919 {
				__e.Return(V4634)
				return
			} else {
				gen9918 := Call(__e, __e.Global(symcons_2), V4635)

				if True == gen9918 {
					gen9913 := Call(__e, __e.Global(symhd), V4635)

					gen9914 := Call(__e, __e.Global(symcons), V4634, Nil)

					gen9915 := Call(__e, __e.Global(symcons), gen9913, gen9914)

					gen9916 := Call(__e, __e.Global(symcons), MakeSymbol("shen.insert"), gen9915)

					gen9917 := Call(__e, __e.Global(symtl), V4635)

					__e.TailApply(__e.Global(symshen_4mkstr_1r), gen9916, gen9917)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.mkstr-r"))

					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.mkstr-r"), gen9920)

		gen9921 := MakeNative(func(__e Evaluator) {
			V4638 := __e.Get(1)
			_ = V4638
			V4639 := __e.Get(2)
			_ = V4639
			__e.TailApply(__e.Global(symshen_4insert_1h), V4638, V4639, MakeString(""))

			return
		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.insert"), gen9921)

		gen9969 := MakeNative(func(__e Evaluator) {
			V4645 := __e.Get(1)
			_ = V4645
			V4646 := __e.Get(2)
			_ = V4646
			V4647 := __e.Get(3)
			_ = V4647
			gen9968 := Call(__e, __e.Global(sym_a), MakeString(""), V4646)

			if True == gen9968 {
				__e.Return(V4647)
				return
			} else {
				gen9966 := Call(__e, __e.Global(symshen_4_7string_2), V4646)

				var gen9967 Obj
				if True == gen9966 {
					gen9963 := Call(__e, __e.Global(sympos), V4646, MakeNumber(0))

					gen9964 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9963)

					var gen9965 Obj
					if True == gen9964 {
						gen9960 := Call(__e, __e.Global(symtlstr), V4646)

						gen9961 := Call(__e, __e.Global(symshen_4_7string_2), gen9960)

						var gen9962 Obj
						if True == gen9961 {
							gen9957 := Call(__e, __e.Global(symtlstr), V4646)

							gen9958 := Call(__e, __e.Global(sympos), gen9957, MakeNumber(0))

							gen9959 := Call(__e, __e.Global(sym_a), MakeString("A"), gen9958)

							if True == gen9959 {
								gen9962 = True
							} else {
								gen9962 = False
							}

						} else {
							gen9962 = False
						}
						if True == gen9962 {
							gen9965 = True
						} else {
							gen9965 = False
						}

					} else {
						gen9965 = False
					}
					if True == gen9965 {
						gen9967 = True
					} else {
						gen9967 = False
					}

				} else {
					gen9967 = False
				}
				if True == gen9967 {
					gen9954 := Call(__e, __e.Global(symtlstr), V4646)

					gen9955 := Call(__e, __e.Global(symtlstr), gen9954)

					gen9956 := Call(__e, __e.Global(symshen_4app), V4645, gen9955, MakeSymbol("shen.a"))

					__e.TailApply(__e.Global(symcn), V4647, gen9956)

					return

				} else {
					gen9952 := Call(__e, __e.Global(symshen_4_7string_2), V4646)

					var gen9953 Obj
					if True == gen9952 {
						gen9949 := Call(__e, __e.Global(sympos), V4646, MakeNumber(0))

						gen9950 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9949)

						var gen9951 Obj
						if True == gen9950 {
							gen9946 := Call(__e, __e.Global(symtlstr), V4646)

							gen9947 := Call(__e, __e.Global(symshen_4_7string_2), gen9946)

							var gen9948 Obj
							if True == gen9947 {
								gen9943 := Call(__e, __e.Global(symtlstr), V4646)

								gen9944 := Call(__e, __e.Global(sympos), gen9943, MakeNumber(0))

								gen9945 := Call(__e, __e.Global(sym_a), MakeString("R"), gen9944)

								if True == gen9945 {
									gen9948 = True
								} else {
									gen9948 = False
								}

							} else {
								gen9948 = False
							}
							if True == gen9948 {
								gen9951 = True
							} else {
								gen9951 = False
							}

						} else {
							gen9951 = False
						}
						if True == gen9951 {
							gen9953 = True
						} else {
							gen9953 = False
						}

					} else {
						gen9953 = False
					}
					if True == gen9953 {
						gen9940 := Call(__e, __e.Global(symtlstr), V4646)

						gen9941 := Call(__e, __e.Global(symtlstr), gen9940)

						gen9942 := Call(__e, __e.Global(symshen_4app), V4645, gen9941, MakeSymbol("shen.r"))

						__e.TailApply(__e.Global(symcn), V4647, gen9942)

						return

					} else {
						gen9938 := Call(__e, __e.Global(symshen_4_7string_2), V4646)

						var gen9939 Obj
						if True == gen9938 {
							gen9935 := Call(__e, __e.Global(sympos), V4646, MakeNumber(0))

							gen9936 := Call(__e, __e.Global(sym_a), MakeString("~"), gen9935)

							var gen9937 Obj
							if True == gen9936 {
								gen9932 := Call(__e, __e.Global(symtlstr), V4646)

								gen9933 := Call(__e, __e.Global(symshen_4_7string_2), gen9932)

								var gen9934 Obj
								if True == gen9933 {
									gen9929 := Call(__e, __e.Global(symtlstr), V4646)

									gen9930 := Call(__e, __e.Global(sympos), gen9929, MakeNumber(0))

									gen9931 := Call(__e, __e.Global(sym_a), MakeString("S"), gen9930)

									if True == gen9931 {
										gen9934 = True
									} else {
										gen9934 = False
									}

								} else {
									gen9934 = False
								}
								if True == gen9934 {
									gen9937 = True
								} else {
									gen9937 = False
								}

							} else {
								gen9937 = False
							}
							if True == gen9937 {
								gen9939 = True
							} else {
								gen9939 = False
							}

						} else {
							gen9939 = False
						}
						if True == gen9939 {
							gen9926 := Call(__e, __e.Global(symtlstr), V4646)

							gen9927 := Call(__e, __e.Global(symtlstr), gen9926)

							gen9928 := Call(__e, __e.Global(symshen_4app), V4645, gen9927, MakeSymbol("shen.s"))

							__e.TailApply(__e.Global(symcn), V4647, gen9928)

							return

						} else {
							gen9925 := Call(__e, __e.Global(symshen_4_7string_2), V4646)

							if True == gen9925 {
								gen9922 := Call(__e, __e.Global(symtlstr), V4646)

								gen9923 := Call(__e, __e.Global(sympos), V4646, MakeNumber(0))

								gen9924 := Call(__e, __e.Global(symcn), V4647, gen9923)

								__e.TailApply(__e.Global(symshen_4insert_1h), V4645, gen9922, gen9924)

								return

							} else {
								__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.insert-h"))

								return
							}

						}

					}

				}

			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.insert-h"), gen9969)

		gen9971 := MakeNative(func(__e Evaluator) {
			V4651 := __e.Get(1)
			_ = V4651
			V4652 := __e.Get(2)
			_ = V4652
			V4653 := __e.Get(3)
			_ = V4653
			gen9970 := Call(__e, __e.Global(symshen_4arg_1_6str), V4651, V4653)

			__e.TailApply(__e.Global(symcn), gen9970, V4652)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.app"), gen9971)

		gen9977 := MakeNative(func(__e Evaluator) {
			V4661 := __e.Get(1)
			_ = V4661
			V4662 := __e.Get(2)
			_ = V4662
			gen9975 := Call(__e, __e.Global(symfail))

			gen9976 := Call(__e, __e.Global(sym_a), V4661, gen9975)

			if True == gen9976 {
				__e.Return(MakeString("..."))
				return
			} else {
				gen9974 := Call(__e, __e.Global(symshen_4list_2), V4661)

				if True == gen9974 {
					__e.TailApply(__e.Global(symshen_4list_1_6str), V4661, V4662)

					return
				} else {
					gen9973 := Call(__e, __e.Global(symstring_2), V4661)

					if True == gen9973 {
						__e.TailApply(__e.Global(symshen_4str_1_6str), V4661, V4662)

						return
					} else {
						gen9972 := Call(__e, __e.Global(symabsvector_2), V4661)

						if True == gen9972 {
							__e.TailApply(__e.Global(symshen_4vector_1_6str), V4661, V4662)

							return
						} else {
							__e.TailApply(__e.Global(symshen_4atom_1_6str), V4661)

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.arg->str"), gen9977)

		gen9985 := MakeNative(func(__e Evaluator) {
			V4665 := __e.Get(1)
			_ = V4665
			V4666 := __e.Get(2)
			_ = V4666
			gen9984 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.r"), V4666)

			if True == gen9984 {
				gen9981 := Call(__e, __e.Global(symshen_4maxseq))

				gen9982 := Call(__e, __e.Global(symshen_4iter_1list), V4665, MakeSymbol("shen.r"), gen9981)

				gen9983 := Call(__e, __e.Global(sym_8s), gen9982, MakeString(")"))

				__e.TailApply(__e.Global(sym_8s), MakeString("("), gen9983)

				return

			} else {
				gen9978 := Call(__e, __e.Global(symshen_4maxseq))

				gen9979 := Call(__e, __e.Global(symshen_4iter_1list), V4665, V4666, gen9978)

				gen9980 := Call(__e, __e.Global(sym_8s), gen9979, MakeString("]"))

				__e.TailApply(__e.Global(sym_8s), MakeString("["), gen9980)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.list->str"), gen9985)

		gen9986 := MakeNative(func(__e Evaluator) {
			__e.TailApply(__e.Global(symvalue), MakeSymbol("*maximum-print-sequence-size*"))

			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.maxseq"), gen9986)

		gen10003 := MakeNative(func(__e Evaluator) {
			V4680 := __e.Get(1)
			_ = V4680
			V4681 := __e.Get(2)
			_ = V4681
			V4682 := __e.Get(3)
			_ = V4682
			gen10002 := Call(__e, __e.Global(sym_a), Nil, V4680)

			if True == gen10002 {
				__e.Return(MakeString(""))
				return
			} else {
				gen10001 := Call(__e, __e.Global(sym_a), MakeNumber(0), V4682)

				if True == gen10001 {
					__e.Return(MakeString("... etc"))
					return
				} else {
					gen9999 := Call(__e, __e.Global(symcons_2), V4680)

					var gen10000 Obj
					if True == gen9999 {
						gen9997 := Call(__e, __e.Global(symtl), V4680)

						gen9998 := Call(__e, __e.Global(sym_a), Nil, gen9997)

						if True == gen9998 {
							gen10000 = True
						} else {
							gen10000 = False
						}

					} else {
						gen10000 = False
					}
					if True == gen10000 {
						gen9996 := Call(__e, __e.Global(symhd), V4680)

						__e.TailApply(__e.Global(symshen_4arg_1_6str), gen9996, V4681)

						return

					} else {
						gen9995 := Call(__e, __e.Global(symcons_2), V4680)

						if True == gen9995 {
							gen9989 := Call(__e, __e.Global(symhd), V4680)

							gen9990 := Call(__e, __e.Global(symshen_4arg_1_6str), gen9989, V4681)

							gen9991 := Call(__e, __e.Global(symtl), V4680)

							gen9992 := Call(__e, __e.Global(sym_1), V4682, MakeNumber(1))

							gen9993 := Call(__e, __e.Global(symshen_4iter_1list), gen9991, V4681, gen9992)

							gen9994 := Call(__e, __e.Global(sym_8s), MakeString(" "), gen9993)

							__e.TailApply(__e.Global(sym_8s), gen9990, gen9994)

							return

						} else {
							gen9987 := Call(__e, __e.Global(symshen_4arg_1_6str), V4680, V4681)

							gen9988 := Call(__e, __e.Global(sym_8s), MakeString(" "), gen9987)

							__e.TailApply(__e.Global(sym_8s), MakeString("|"), gen9988)

							return

						}

					}

				}

			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.iter-list"), gen10003)

		gen10008 := MakeNative(func(__e Evaluator) {
			V4689 := __e.Get(1)
			_ = V4689
			V4690 := __e.Get(2)
			_ = V4690
			gen10007 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.a"), V4690)

			if True == gen10007 {
				__e.Return(V4689)
				return
			} else {
				gen10004 := Call(__e, __e.Global(symn_1_6string), MakeNumber(34))

				gen10005 := Call(__e, __e.Global(symn_1_6string), MakeNumber(34))

				gen10006 := Call(__e, __e.Global(sym_8s), V4689, gen10005)

				__e.TailApply(__e.Global(sym_8s), gen10004, gen10006)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.str->str"), gen10008)

		gen10020 := MakeNative(func(__e Evaluator) {
			V4693 := __e.Get(1)
			_ = V4693
			V4694 := __e.Get(2)
			_ = V4694
			gen10019 := Call(__e, __e.Global(symshen_4print_1vector_2), V4693)

			if True == gen10019 {
				gen10017 := Call(__e, __e.Global(sym_5_1address), V4693, MakeNumber(0))

				gen10018 := Call(__e, __e.Global(symfunction), gen10017)

				f32 := gen10018
				__e.TailApply(f32, V4693)

				return

			} else {
				gen10016 := Call(__e, __e.Global(symvector_2), V4693)

				if True == gen10016 {
					gen10013 := Call(__e, __e.Global(symshen_4maxseq))

					gen10014 := Call(__e, __e.Global(symshen_4iter_1vector), V4693, MakeNumber(1), V4694, gen10013)

					gen10015 := Call(__e, __e.Global(sym_8s), gen10014, MakeString(">"))

					__e.TailApply(__e.Global(sym_8s), MakeString("<"), gen10015)

					return

				} else {
					gen10009 := Call(__e, __e.Global(symshen_4maxseq))

					gen10010 := Call(__e, __e.Global(symshen_4iter_1vector), V4693, MakeNumber(0), V4694, gen10009)

					gen10011 := Call(__e, __e.Global(sym_8s), gen10010, MakeString(">>"))

					gen10012 := Call(__e, __e.Global(sym_8s), MakeString("<"), gen10011)

					__e.TailApply(__e.Global(sym_8s), MakeString("<"), gen10012)

					return

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.vector->str"), gen10020)

		gen10022 := MakeNative(func(__e Evaluator) {
			V4696 := __e.Get(1)
			_ = V4696
			gen10021 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*empty-absvector*"))

			__e.TailApply(__e.Global(sym_a), V4696, gen10021)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.empty-absvector?"), gen10022)

		gen10036 := MakeNative(func(__e Evaluator) {
			V4698 := __e.Get(1)
			_ = V4698
			gen10034 := Call(__e, __e.Global(symshen_4empty_1absvector_2), V4698)

			gen10035 := Call(__e, __e.Global(symnot), gen10034)

			if True == gen10035 {
				gen10023 := Call(__e, __e.Global(sym_5_1address), V4698, MakeNumber(0))

				First := gen10023
				gen10032 := Call(__e, __e.Global(sym_a), First, MakeSymbol("shen.tuple"))

				var gen10033 Obj
				if True == gen10032 {
					gen10033 = True
				} else {
					gen10030 := Call(__e, __e.Global(sym_a), First, MakeSymbol("shen.pvar"))

					var gen10031 Obj
					if True == gen10030 {
						gen10031 = True
					} else {
						gen10028 := Call(__e, __e.Global(sym_a), First, MakeSymbol("shen.dictionary"))

						var gen10029 Obj
						if True == gen10028 {
							gen10029 = True
						} else {
							gen10025 := Call(__e, __e.Global(symnumber_2), First)

							gen10026 := Call(__e, __e.Global(symnot), gen10025)

							var gen10027 Obj
							if True == gen10026 {
								gen10024 := Call(__e, __e.Global(symshen_4fbound_2), First)

								if True == gen10024 {
									gen10027 = True
								} else {
									gen10027 = False
								}

							} else {
								gen10027 = False
							}
							if True == gen10027 {
								gen10029 = True
							} else {
								gen10029 = False
							}

						}
						if True == gen10029 {
							gen10031 = True
						} else {
							gen10031 = False
						}

					}
					if True == gen10031 {
						gen10033 = True
					} else {
						gen10033 = False
					}

				}
				if True == gen10033 {
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
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.print-vector?"), gen10036)

		gen10039 := MakeNative(func(__e Evaluator) {
			V4700 := __e.Get(1)
			_ = V4700
			gen10037 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)
			gen10038 := MakeNative(func(__e Evaluator) {
				Call(__e, __e.Global(symshen_4lookup_1func), V4700)
				__e.Return(True)
				return

			}, 0)
			__e.Return(Try(__e, gen10038).Catch(gen10037))
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.fbound?"), gen10039)

		gen10045 := MakeNative(func(__e Evaluator) {
			V4702 := __e.Get(1)
			_ = V4702
			gen10040 := Call(__e, __e.Global(sym_5_1address), V4702, MakeNumber(1))

			gen10041 := Call(__e, __e.Global(sym_5_1address), V4702, MakeNumber(2))

			gen10042 := Call(__e, __e.Global(symshen_4app), gen10041, MakeString(")"), MakeSymbol("shen.s"))

			gen10043 := Call(__e, __e.Global(symcn), MakeString(" "), gen10042)

			gen10044 := Call(__e, __e.Global(symshen_4app), gen10040, gen10043, MakeSymbol("shen.s"))

			__e.TailApply(__e.Global(symcn), MakeString("(@p "), gen10044)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.tuple"), gen10045)

		gen10046 := MakeNative(func(__e Evaluator) {
			V4704 := __e.Get(1)
			_ = V4704
			__e.Return(MakeString("(dict ...)"))
			return
		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dictionary"), gen10046)

		gen10062 := MakeNative(func(__e Evaluator) {
			V4715 := __e.Get(1)
			_ = V4715
			V4716 := __e.Get(2)
			_ = V4716
			V4717 := __e.Get(3)
			_ = V4717
			V4718 := __e.Get(4)
			_ = V4718
			gen10061 := Call(__e, __e.Global(sym_a), MakeNumber(0), V4718)

			if True == gen10061 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				gen10047 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeSymbol("shen.out-of-bounds"))
					return
				}, 1)
				gen10048 := MakeNative(func(__e Evaluator) {
					__e.TailApply(__e.Global(sym_5_1address), V4715, V4716)

					return
				}, 0)
				gen10049 := Try(__e, gen10048).Catch(gen10047)
				Item := gen10049
				gen10050 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeSymbol("shen.out-of-bounds"))
					return
				}, 1)
				gen10052 := MakeNative(func(__e Evaluator) {
					gen10051 := Call(__e, __e.Global(sym_7), V4716, MakeNumber(1))

					__e.TailApply(__e.Global(sym_5_1address), V4715, gen10051)

					return

				}, 0)
				gen10053 := Try(__e, gen10052).Catch(gen10050)
				Next := gen10053
				gen10060 := Call(__e, __e.Global(sym_a), Item, MakeSymbol("shen.out-of-bounds"))

				if True == gen10060 {
					__e.Return(MakeString(""))
					return
				} else {
					gen10059 := Call(__e, __e.Global(sym_a), Next, MakeSymbol("shen.out-of-bounds"))

					if True == gen10059 {
						__e.TailApply(__e.Global(symshen_4arg_1_6str), Item, V4717)

						return
					} else {
						gen10054 := Call(__e, __e.Global(symshen_4arg_1_6str), Item, V4717)

						gen10055 := Call(__e, __e.Global(sym_7), V4716, MakeNumber(1))

						gen10056 := Call(__e, __e.Global(sym_1), V4718, MakeNumber(1))

						gen10057 := Call(__e, __e.Global(symshen_4iter_1vector), V4715, gen10055, V4717, gen10056)

						gen10058 := Call(__e, __e.Global(sym_8s), MakeString(" "), gen10057)

						__e.TailApply(__e.Global(sym_8s), gen10054, gen10058)

						return

					}

				}

			}

		}, 4)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.iter-vector"), gen10062)

		gen10065 := MakeNative(func(__e Evaluator) {
			V4720 := __e.Get(1)
			_ = V4720
			gen10063 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.TailApply(__e.Global(symshen_4funexstring))

				return
			}, 1)
			gen10064 := MakeNative(func(__e Evaluator) {
				__e.TailApply(__e.Global(symstr), V4720)

				return
			}, 0)
			__e.Return(Try(__e, gen10064).Catch(gen10063))
			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.atom->str"), gen10065)

		gen10074 := MakeNative(func(__e Evaluator) {
			gen10066 := Call(__e, __e.Global(symintern), MakeString("x"))

			gen10067 := Call(__e, __e.Global(symgensym), gen10066)

			gen10068 := Call(__e, __e.Global(symshen_4arg_1_6str), gen10067, MakeSymbol("shen.a"))

			gen10069 := Call(__e, __e.Global(sym_8s), gen10068, MakeString("\x11"))

			gen10070 := Call(__e, __e.Global(sym_8s), MakeString("e"), gen10069)

			gen10071 := Call(__e, __e.Global(sym_8s), MakeString("n"), gen10070)

			gen10072 := Call(__e, __e.Global(sym_8s), MakeString("u"), gen10071)

			gen10073 := Call(__e, __e.Global(sym_8s), MakeString("f"), gen10072)

			__e.TailApply(__e.Global(sym_8s), MakeString("\x10"), gen10073)

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.funexstring"), gen10074)

		gen10077 := MakeNative(func(__e Evaluator) {
			V4722 := __e.Get(1)
			_ = V4722
			gen10076 := Call(__e, __e.Global(symempty_2), V4722)

			if True == gen10076 {
				__e.Return(True)
				return
			} else {
				gen10075 := Call(__e, __e.Global(symcons_2), V4722)

				if True == gen10075 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.list?"), gen10077)

		return

	}, 0))
}
