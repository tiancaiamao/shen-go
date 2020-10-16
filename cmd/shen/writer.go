package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__pr Obj                          // pr
var __defun__shen_4prh Obj                   // shen.prh
var __defun__shen_4write_1char_1and_1inc Obj // shen.write-char-and-inc
var __defun__print Obj                       // print
var __defun__shen_4prhush Obj                // shen.prhush
var __defun__shen_4mkstr Obj                 // shen.mkstr
var __defun__shen_4mkstr_1l Obj              // shen.mkstr-l
var __defun__shen_4insert_1l Obj             // shen.insert-l
var __defun__shen_4factor_1cn Obj            // shen.factor-cn
var __defun__shen_4proc_1nl Obj              // shen.proc-nl
var __defun__shen_4mkstr_1r Obj              // shen.mkstr-r
var __defun__shen_4insert Obj                // shen.insert
var __defun__shen_4insert_1h Obj             // shen.insert-h
var __defun__shen_4app Obj                   // shen.app
var __defun__shen_4arg_1_6str Obj            // shen.arg->str
var __defun__shen_4list_1_6str Obj           // shen.list->str
var __defun__shen_4maxseq Obj                // shen.maxseq
var __defun__shen_4iter_1list Obj            // shen.iter-list
var __defun__shen_4str_1_6str Obj            // shen.str->str
var __defun__shen_4vector_1_6str Obj         // shen.vector->str
var __defun__shen_4empty_1absvector_2 Obj    // shen.empty-absvector?
var __defun__shen_4print_1vector_2 Obj       // shen.print-vector?
var __defun__shen_4fbound_2 Obj              // shen.fbound?
var __defun__shen_4tuple Obj                 // shen.tuple
var __defun__shen_4dictionary Obj            // shen.dictionary
var __defun__shen_4iter_1vector Obj          // shen.iter-vector
var __defun__shen_4atom_1_6str Obj           // shen.atom->str
var __defun__shen_4funexstring Obj           // shen.funexstring
var __defun__shen_4list_2 Obj                // shen.list?

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg16656 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__ctx.Return(reg16656)
		return
	}, 0))
	__defun__pr = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4564 := __args[0]
		_ = V4564
		V4565 := __args[1]
		_ = V4565
		reg16657 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg16658 := MakeNumber(0)
			__ctx.TailApply(__defun__shen_4prh, V4564, V4565, reg16658)
			return
		}, 0)
		reg16660 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			__ctx.Return(V4564)
			return
		}, 1)
		reg16661 := __e.Try(reg16657).Catch(reg16660)
		__ctx.Return(reg16661)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "pr", value: __defun__pr})

	__defun__shen_4prh = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4569 := __args[0]
		_ = V4569
		V4570 := __args[1]
		_ = V4570
		V4571 := __args[2]
		_ = V4571
		reg16662 := __e.Call(__defun__shen_4write_1char_1and_1inc, V4569, V4570, V4571)
		__ctx.TailApply(__defun__shen_4prh, V4569, V4570, reg16662)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.prh", value: __defun__shen_4prh})

	__defun__shen_4write_1char_1and_1inc = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4575 := __args[0]
		_ = V4575
		V4576 := __args[1]
		_ = V4576
		V4577 := __args[2]
		_ = V4577
		reg16664 := PrimPos(V4575, V4577)
		reg16665 := PrimStringToNumber(reg16664)
		reg16666 := PrimWriteByte(reg16665, V4576)
		_ = reg16666
		reg16667 := MakeNumber(1)
		reg16668 := PrimNumberAdd(V4577, reg16667)
		__ctx.Return(reg16668)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.write-char-and-inc", value: __defun__shen_4write_1char_1and_1inc})

	__defun__print = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4579 := __args[0]
		_ = V4579
		reg16669 := MakeString("~S")
		reg16670 := __e.Call(__defun__shen_4insert, V4579, reg16669)
		String := reg16670
		_ = String
		reg16671 := __e.Call(__defun__stoutput)
		reg16672 := __e.Call(__defun__shen_4prhush, String, reg16671)
		Print := reg16672
		_ = Print
		__ctx.Return(V4579)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "print", value: __defun__print})

	__defun__shen_4prhush = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4582 := __args[0]
		_ = V4582
		V4583 := __args[1]
		_ = V4583
		reg16673 := MakeSymbol("*hush*")
		reg16674 := PrimValue(reg16673)
		if reg16674 == True {
			__ctx.Return(V4582)
			return
		} else {
			__ctx.TailApply(__defun__pr, V4582, V4583)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.prhush", value: __defun__shen_4prhush})

	__defun__shen_4mkstr = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4586 := __args[0]
		_ = V4586
		V4587 := __args[1]
		_ = V4587
		reg16676 := PrimIsString(V4586)
		if reg16676 == True {
			reg16677 := __e.Call(__defun__shen_4proc_1nl, V4586)
			__ctx.TailApply(__defun__shen_4mkstr_1l, reg16677, V4587)
			return
		} else {
			reg16679 := MakeSymbol("shen.proc-nl")
			reg16680 := Nil
			reg16681 := PrimCons(V4586, reg16680)
			reg16682 := PrimCons(reg16679, reg16681)
			__ctx.TailApply(__defun__shen_4mkstr_1r, reg16682, V4587)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.mkstr", value: __defun__shen_4mkstr})

	__defun__shen_4mkstr_1l = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4590 := __args[0]
		_ = V4590
		V4591 := __args[1]
		_ = V4591
		reg16684 := Nil
		reg16685 := PrimEqual(reg16684, V4591)
		if reg16685 == True {
			__ctx.Return(V4590)
			return
		} else {
			reg16686 := PrimIsPair(V4591)
			if reg16686 == True {
				reg16687 := PrimHead(V4591)
				reg16688 := __e.Call(__defun__shen_4insert_1l, reg16687, V4590)
				reg16689 := PrimTail(V4591)
				__ctx.TailApply(__defun__shen_4mkstr_1l, reg16688, reg16689)
				return
			} else {
				reg16691 := MakeSymbol("shen.mkstr-l")
				__ctx.TailApply(__defun__shen_4f__error, reg16691)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.mkstr-l", value: __defun__shen_4mkstr_1l})

	__defun__shen_4insert_1l = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4596 := __args[0]
		_ = V4596
		V4597 := __args[1]
		_ = V4597
		reg16693 := MakeString("")
		reg16694 := PrimEqual(reg16693, V4597)
		if reg16694 == True {
			reg16695 := MakeString("")
			__ctx.Return(reg16695)
			return
		} else {
			reg16696 := __e.Call(__defun__shen_4_7string_2, V4597)
			var reg16722 Obj
			if reg16696 == True {
				reg16697 := MakeString("~")
				reg16698 := MakeNumber(0)
				reg16699 := PrimPos(V4597, reg16698)
				reg16700 := PrimEqual(reg16697, reg16699)
				var reg16717 Obj
				if reg16700 == True {
					reg16701 := PrimTailString(V4597)
					reg16702 := __e.Call(__defun__shen_4_7string_2, reg16701)
					var reg16712 Obj
					if reg16702 == True {
						reg16703 := MakeString("A")
						reg16704 := PrimTailString(V4597)
						reg16705 := MakeNumber(0)
						reg16706 := PrimPos(reg16704, reg16705)
						reg16707 := PrimEqual(reg16703, reg16706)
						var reg16710 Obj
						if reg16707 == True {
							reg16708 := True
							reg16710 = reg16708
						} else {
							reg16709 := False
							reg16710 = reg16709
						}
						reg16712 = reg16710
					} else {
						reg16711 := False
						reg16712 = reg16711
					}
					var reg16715 Obj
					if reg16712 == True {
						reg16713 := True
						reg16715 = reg16713
					} else {
						reg16714 := False
						reg16715 = reg16714
					}
					reg16717 = reg16715
				} else {
					reg16716 := False
					reg16717 = reg16716
				}
				var reg16720 Obj
				if reg16717 == True {
					reg16718 := True
					reg16720 = reg16718
				} else {
					reg16719 := False
					reg16720 = reg16719
				}
				reg16722 = reg16720
			} else {
				reg16721 := False
				reg16722 = reg16721
			}
			if reg16722 == True {
				reg16723 := MakeSymbol("shen.app")
				reg16724 := PrimTailString(V4597)
				reg16725 := PrimTailString(reg16724)
				reg16726 := MakeSymbol("shen.a")
				reg16727 := Nil
				reg16728 := PrimCons(reg16726, reg16727)
				reg16729 := PrimCons(reg16725, reg16728)
				reg16730 := PrimCons(V4596, reg16729)
				reg16731 := PrimCons(reg16723, reg16730)
				__ctx.Return(reg16731)
				return
			} else {
				reg16732 := __e.Call(__defun__shen_4_7string_2, V4597)
				var reg16758 Obj
				if reg16732 == True {
					reg16733 := MakeString("~")
					reg16734 := MakeNumber(0)
					reg16735 := PrimPos(V4597, reg16734)
					reg16736 := PrimEqual(reg16733, reg16735)
					var reg16753 Obj
					if reg16736 == True {
						reg16737 := PrimTailString(V4597)
						reg16738 := __e.Call(__defun__shen_4_7string_2, reg16737)
						var reg16748 Obj
						if reg16738 == True {
							reg16739 := MakeString("R")
							reg16740 := PrimTailString(V4597)
							reg16741 := MakeNumber(0)
							reg16742 := PrimPos(reg16740, reg16741)
							reg16743 := PrimEqual(reg16739, reg16742)
							var reg16746 Obj
							if reg16743 == True {
								reg16744 := True
								reg16746 = reg16744
							} else {
								reg16745 := False
								reg16746 = reg16745
							}
							reg16748 = reg16746
						} else {
							reg16747 := False
							reg16748 = reg16747
						}
						var reg16751 Obj
						if reg16748 == True {
							reg16749 := True
							reg16751 = reg16749
						} else {
							reg16750 := False
							reg16751 = reg16750
						}
						reg16753 = reg16751
					} else {
						reg16752 := False
						reg16753 = reg16752
					}
					var reg16756 Obj
					if reg16753 == True {
						reg16754 := True
						reg16756 = reg16754
					} else {
						reg16755 := False
						reg16756 = reg16755
					}
					reg16758 = reg16756
				} else {
					reg16757 := False
					reg16758 = reg16757
				}
				if reg16758 == True {
					reg16759 := MakeSymbol("shen.app")
					reg16760 := PrimTailString(V4597)
					reg16761 := PrimTailString(reg16760)
					reg16762 := MakeSymbol("shen.r")
					reg16763 := Nil
					reg16764 := PrimCons(reg16762, reg16763)
					reg16765 := PrimCons(reg16761, reg16764)
					reg16766 := PrimCons(V4596, reg16765)
					reg16767 := PrimCons(reg16759, reg16766)
					__ctx.Return(reg16767)
					return
				} else {
					reg16768 := __e.Call(__defun__shen_4_7string_2, V4597)
					var reg16794 Obj
					if reg16768 == True {
						reg16769 := MakeString("~")
						reg16770 := MakeNumber(0)
						reg16771 := PrimPos(V4597, reg16770)
						reg16772 := PrimEqual(reg16769, reg16771)
						var reg16789 Obj
						if reg16772 == True {
							reg16773 := PrimTailString(V4597)
							reg16774 := __e.Call(__defun__shen_4_7string_2, reg16773)
							var reg16784 Obj
							if reg16774 == True {
								reg16775 := MakeString("S")
								reg16776 := PrimTailString(V4597)
								reg16777 := MakeNumber(0)
								reg16778 := PrimPos(reg16776, reg16777)
								reg16779 := PrimEqual(reg16775, reg16778)
								var reg16782 Obj
								if reg16779 == True {
									reg16780 := True
									reg16782 = reg16780
								} else {
									reg16781 := False
									reg16782 = reg16781
								}
								reg16784 = reg16782
							} else {
								reg16783 := False
								reg16784 = reg16783
							}
							var reg16787 Obj
							if reg16784 == True {
								reg16785 := True
								reg16787 = reg16785
							} else {
								reg16786 := False
								reg16787 = reg16786
							}
							reg16789 = reg16787
						} else {
							reg16788 := False
							reg16789 = reg16788
						}
						var reg16792 Obj
						if reg16789 == True {
							reg16790 := True
							reg16792 = reg16790
						} else {
							reg16791 := False
							reg16792 = reg16791
						}
						reg16794 = reg16792
					} else {
						reg16793 := False
						reg16794 = reg16793
					}
					if reg16794 == True {
						reg16795 := MakeSymbol("shen.app")
						reg16796 := PrimTailString(V4597)
						reg16797 := PrimTailString(reg16796)
						reg16798 := MakeSymbol("shen.s")
						reg16799 := Nil
						reg16800 := PrimCons(reg16798, reg16799)
						reg16801 := PrimCons(reg16797, reg16800)
						reg16802 := PrimCons(V4596, reg16801)
						reg16803 := PrimCons(reg16795, reg16802)
						__ctx.Return(reg16803)
						return
					} else {
						reg16804 := __e.Call(__defun__shen_4_7string_2, V4597)
						if reg16804 == True {
							reg16805 := MakeSymbol("cn")
							reg16806 := MakeNumber(0)
							reg16807 := PrimPos(V4597, reg16806)
							reg16808 := PrimTailString(V4597)
							reg16809 := __e.Call(__defun__shen_4insert_1l, V4596, reg16808)
							reg16810 := Nil
							reg16811 := PrimCons(reg16809, reg16810)
							reg16812 := PrimCons(reg16807, reg16811)
							reg16813 := PrimCons(reg16805, reg16812)
							__ctx.TailApply(__defun__shen_4factor_1cn, reg16813)
							return
						} else {
							reg16815 := PrimIsPair(V4597)
							var reg16848 Obj
							if reg16815 == True {
								reg16816 := MakeSymbol("cn")
								reg16817 := PrimHead(V4597)
								reg16818 := PrimEqual(reg16816, reg16817)
								var reg16843 Obj
								if reg16818 == True {
									reg16819 := PrimTail(V4597)
									reg16820 := PrimIsPair(reg16819)
									var reg16838 Obj
									if reg16820 == True {
										reg16821 := PrimTail(V4597)
										reg16822 := PrimTail(reg16821)
										reg16823 := PrimIsPair(reg16822)
										var reg16833 Obj
										if reg16823 == True {
											reg16824 := Nil
											reg16825 := PrimTail(V4597)
											reg16826 := PrimTail(reg16825)
											reg16827 := PrimTail(reg16826)
											reg16828 := PrimEqual(reg16824, reg16827)
											var reg16831 Obj
											if reg16828 == True {
												reg16829 := True
												reg16831 = reg16829
											} else {
												reg16830 := False
												reg16831 = reg16830
											}
											reg16833 = reg16831
										} else {
											reg16832 := False
											reg16833 = reg16832
										}
										var reg16836 Obj
										if reg16833 == True {
											reg16834 := True
											reg16836 = reg16834
										} else {
											reg16835 := False
											reg16836 = reg16835
										}
										reg16838 = reg16836
									} else {
										reg16837 := False
										reg16838 = reg16837
									}
									var reg16841 Obj
									if reg16838 == True {
										reg16839 := True
										reg16841 = reg16839
									} else {
										reg16840 := False
										reg16841 = reg16840
									}
									reg16843 = reg16841
								} else {
									reg16842 := False
									reg16843 = reg16842
								}
								var reg16846 Obj
								if reg16843 == True {
									reg16844 := True
									reg16846 = reg16844
								} else {
									reg16845 := False
									reg16846 = reg16845
								}
								reg16848 = reg16846
							} else {
								reg16847 := False
								reg16848 = reg16847
							}
							if reg16848 == True {
								reg16849 := MakeSymbol("cn")
								reg16850 := PrimTail(V4597)
								reg16851 := PrimHead(reg16850)
								reg16852 := PrimTail(V4597)
								reg16853 := PrimTail(reg16852)
								reg16854 := PrimHead(reg16853)
								reg16855 := __e.Call(__defun__shen_4insert_1l, V4596, reg16854)
								reg16856 := Nil
								reg16857 := PrimCons(reg16855, reg16856)
								reg16858 := PrimCons(reg16851, reg16857)
								reg16859 := PrimCons(reg16849, reg16858)
								__ctx.Return(reg16859)
								return
							} else {
								reg16860 := PrimIsPair(V4597)
								var reg16903 Obj
								if reg16860 == True {
									reg16861 := MakeSymbol("shen.app")
									reg16862 := PrimHead(V4597)
									reg16863 := PrimEqual(reg16861, reg16862)
									var reg16898 Obj
									if reg16863 == True {
										reg16864 := PrimTail(V4597)
										reg16865 := PrimIsPair(reg16864)
										var reg16893 Obj
										if reg16865 == True {
											reg16866 := PrimTail(V4597)
											reg16867 := PrimTail(reg16866)
											reg16868 := PrimIsPair(reg16867)
											var reg16888 Obj
											if reg16868 == True {
												reg16869 := PrimTail(V4597)
												reg16870 := PrimTail(reg16869)
												reg16871 := PrimTail(reg16870)
												reg16872 := PrimIsPair(reg16871)
												var reg16883 Obj
												if reg16872 == True {
													reg16873 := Nil
													reg16874 := PrimTail(V4597)
													reg16875 := PrimTail(reg16874)
													reg16876 := PrimTail(reg16875)
													reg16877 := PrimTail(reg16876)
													reg16878 := PrimEqual(reg16873, reg16877)
													var reg16881 Obj
													if reg16878 == True {
														reg16879 := True
														reg16881 = reg16879
													} else {
														reg16880 := False
														reg16881 = reg16880
													}
													reg16883 = reg16881
												} else {
													reg16882 := False
													reg16883 = reg16882
												}
												var reg16886 Obj
												if reg16883 == True {
													reg16884 := True
													reg16886 = reg16884
												} else {
													reg16885 := False
													reg16886 = reg16885
												}
												reg16888 = reg16886
											} else {
												reg16887 := False
												reg16888 = reg16887
											}
											var reg16891 Obj
											if reg16888 == True {
												reg16889 := True
												reg16891 = reg16889
											} else {
												reg16890 := False
												reg16891 = reg16890
											}
											reg16893 = reg16891
										} else {
											reg16892 := False
											reg16893 = reg16892
										}
										var reg16896 Obj
										if reg16893 == True {
											reg16894 := True
											reg16896 = reg16894
										} else {
											reg16895 := False
											reg16896 = reg16895
										}
										reg16898 = reg16896
									} else {
										reg16897 := False
										reg16898 = reg16897
									}
									var reg16901 Obj
									if reg16898 == True {
										reg16899 := True
										reg16901 = reg16899
									} else {
										reg16900 := False
										reg16901 = reg16900
									}
									reg16903 = reg16901
								} else {
									reg16902 := False
									reg16903 = reg16902
								}
								if reg16903 == True {
									reg16904 := MakeSymbol("shen.app")
									reg16905 := PrimTail(V4597)
									reg16906 := PrimHead(reg16905)
									reg16907 := PrimTail(V4597)
									reg16908 := PrimTail(reg16907)
									reg16909 := PrimHead(reg16908)
									reg16910 := __e.Call(__defun__shen_4insert_1l, V4596, reg16909)
									reg16911 := PrimTail(V4597)
									reg16912 := PrimTail(reg16911)
									reg16913 := PrimTail(reg16912)
									reg16914 := PrimCons(reg16910, reg16913)
									reg16915 := PrimCons(reg16906, reg16914)
									reg16916 := PrimCons(reg16904, reg16915)
									__ctx.Return(reg16916)
									return
								} else {
									reg16917 := MakeSymbol("shen.insert-l")
									__ctx.TailApply(__defun__shen_4f__error, reg16917)
									return
								}
							}
						}
					}
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.insert-l", value: __defun__shen_4insert_1l})

	__defun__shen_4factor_1cn = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4599 := __args[0]
		_ = V4599
		reg16919 := PrimIsPair(V4599)
		var reg17025 Obj
		if reg16919 == True {
			reg16920 := MakeSymbol("cn")
			reg16921 := PrimHead(V4599)
			reg16922 := PrimEqual(reg16920, reg16921)
			var reg17020 Obj
			if reg16922 == True {
				reg16923 := PrimTail(V4599)
				reg16924 := PrimIsPair(reg16923)
				var reg17015 Obj
				if reg16924 == True {
					reg16925 := PrimTail(V4599)
					reg16926 := PrimTail(reg16925)
					reg16927 := PrimIsPair(reg16926)
					var reg17010 Obj
					if reg16927 == True {
						reg16928 := PrimTail(V4599)
						reg16929 := PrimTail(reg16928)
						reg16930 := PrimHead(reg16929)
						reg16931 := PrimIsPair(reg16930)
						var reg17005 Obj
						if reg16931 == True {
							reg16932 := MakeSymbol("cn")
							reg16933 := PrimTail(V4599)
							reg16934 := PrimTail(reg16933)
							reg16935 := PrimHead(reg16934)
							reg16936 := PrimHead(reg16935)
							reg16937 := PrimEqual(reg16932, reg16936)
							var reg17000 Obj
							if reg16937 == True {
								reg16938 := PrimTail(V4599)
								reg16939 := PrimTail(reg16938)
								reg16940 := PrimHead(reg16939)
								reg16941 := PrimTail(reg16940)
								reg16942 := PrimIsPair(reg16941)
								var reg16995 Obj
								if reg16942 == True {
									reg16943 := PrimTail(V4599)
									reg16944 := PrimTail(reg16943)
									reg16945 := PrimHead(reg16944)
									reg16946 := PrimTail(reg16945)
									reg16947 := PrimTail(reg16946)
									reg16948 := PrimIsPair(reg16947)
									var reg16990 Obj
									if reg16948 == True {
										reg16949 := Nil
										reg16950 := PrimTail(V4599)
										reg16951 := PrimTail(reg16950)
										reg16952 := PrimHead(reg16951)
										reg16953 := PrimTail(reg16952)
										reg16954 := PrimTail(reg16953)
										reg16955 := PrimTail(reg16954)
										reg16956 := PrimEqual(reg16949, reg16955)
										var reg16985 Obj
										if reg16956 == True {
											reg16957 := Nil
											reg16958 := PrimTail(V4599)
											reg16959 := PrimTail(reg16958)
											reg16960 := PrimTail(reg16959)
											reg16961 := PrimEqual(reg16957, reg16960)
											var reg16980 Obj
											if reg16961 == True {
												reg16962 := PrimTail(V4599)
												reg16963 := PrimHead(reg16962)
												reg16964 := PrimIsString(reg16963)
												var reg16975 Obj
												if reg16964 == True {
													reg16965 := PrimTail(V4599)
													reg16966 := PrimTail(reg16965)
													reg16967 := PrimHead(reg16966)
													reg16968 := PrimTail(reg16967)
													reg16969 := PrimHead(reg16968)
													reg16970 := PrimIsString(reg16969)
													var reg16973 Obj
													if reg16970 == True {
														reg16971 := True
														reg16973 = reg16971
													} else {
														reg16972 := False
														reg16973 = reg16972
													}
													reg16975 = reg16973
												} else {
													reg16974 := False
													reg16975 = reg16974
												}
												var reg16978 Obj
												if reg16975 == True {
													reg16976 := True
													reg16978 = reg16976
												} else {
													reg16977 := False
													reg16978 = reg16977
												}
												reg16980 = reg16978
											} else {
												reg16979 := False
												reg16980 = reg16979
											}
											var reg16983 Obj
											if reg16980 == True {
												reg16981 := True
												reg16983 = reg16981
											} else {
												reg16982 := False
												reg16983 = reg16982
											}
											reg16985 = reg16983
										} else {
											reg16984 := False
											reg16985 = reg16984
										}
										var reg16988 Obj
										if reg16985 == True {
											reg16986 := True
											reg16988 = reg16986
										} else {
											reg16987 := False
											reg16988 = reg16987
										}
										reg16990 = reg16988
									} else {
										reg16989 := False
										reg16990 = reg16989
									}
									var reg16993 Obj
									if reg16990 == True {
										reg16991 := True
										reg16993 = reg16991
									} else {
										reg16992 := False
										reg16993 = reg16992
									}
									reg16995 = reg16993
								} else {
									reg16994 := False
									reg16995 = reg16994
								}
								var reg16998 Obj
								if reg16995 == True {
									reg16996 := True
									reg16998 = reg16996
								} else {
									reg16997 := False
									reg16998 = reg16997
								}
								reg17000 = reg16998
							} else {
								reg16999 := False
								reg17000 = reg16999
							}
							var reg17003 Obj
							if reg17000 == True {
								reg17001 := True
								reg17003 = reg17001
							} else {
								reg17002 := False
								reg17003 = reg17002
							}
							reg17005 = reg17003
						} else {
							reg17004 := False
							reg17005 = reg17004
						}
						var reg17008 Obj
						if reg17005 == True {
							reg17006 := True
							reg17008 = reg17006
						} else {
							reg17007 := False
							reg17008 = reg17007
						}
						reg17010 = reg17008
					} else {
						reg17009 := False
						reg17010 = reg17009
					}
					var reg17013 Obj
					if reg17010 == True {
						reg17011 := True
						reg17013 = reg17011
					} else {
						reg17012 := False
						reg17013 = reg17012
					}
					reg17015 = reg17013
				} else {
					reg17014 := False
					reg17015 = reg17014
				}
				var reg17018 Obj
				if reg17015 == True {
					reg17016 := True
					reg17018 = reg17016
				} else {
					reg17017 := False
					reg17018 = reg17017
				}
				reg17020 = reg17018
			} else {
				reg17019 := False
				reg17020 = reg17019
			}
			var reg17023 Obj
			if reg17020 == True {
				reg17021 := True
				reg17023 = reg17021
			} else {
				reg17022 := False
				reg17023 = reg17022
			}
			reg17025 = reg17023
		} else {
			reg17024 := False
			reg17025 = reg17024
		}
		if reg17025 == True {
			reg17026 := MakeSymbol("cn")
			reg17027 := PrimTail(V4599)
			reg17028 := PrimHead(reg17027)
			reg17029 := PrimTail(V4599)
			reg17030 := PrimTail(reg17029)
			reg17031 := PrimHead(reg17030)
			reg17032 := PrimTail(reg17031)
			reg17033 := PrimHead(reg17032)
			reg17034 := PrimStringConcat(reg17028, reg17033)
			reg17035 := PrimTail(V4599)
			reg17036 := PrimTail(reg17035)
			reg17037 := PrimHead(reg17036)
			reg17038 := PrimTail(reg17037)
			reg17039 := PrimTail(reg17038)
			reg17040 := PrimCons(reg17034, reg17039)
			reg17041 := PrimCons(reg17026, reg17040)
			__ctx.Return(reg17041)
			return
		} else {
			__ctx.Return(V4599)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.factor-cn", value: __defun__shen_4factor_1cn})

	__defun__shen_4proc_1nl = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4601 := __args[0]
		_ = V4601
		reg17042 := MakeString("")
		reg17043 := PrimEqual(reg17042, V4601)
		if reg17043 == True {
			reg17044 := MakeString("")
			__ctx.Return(reg17044)
			return
		} else {
			reg17045 := __e.Call(__defun__shen_4_7string_2, V4601)
			var reg17071 Obj
			if reg17045 == True {
				reg17046 := MakeString("~")
				reg17047 := MakeNumber(0)
				reg17048 := PrimPos(V4601, reg17047)
				reg17049 := PrimEqual(reg17046, reg17048)
				var reg17066 Obj
				if reg17049 == True {
					reg17050 := PrimTailString(V4601)
					reg17051 := __e.Call(__defun__shen_4_7string_2, reg17050)
					var reg17061 Obj
					if reg17051 == True {
						reg17052 := MakeString("%")
						reg17053 := PrimTailString(V4601)
						reg17054 := MakeNumber(0)
						reg17055 := PrimPos(reg17053, reg17054)
						reg17056 := PrimEqual(reg17052, reg17055)
						var reg17059 Obj
						if reg17056 == True {
							reg17057 := True
							reg17059 = reg17057
						} else {
							reg17058 := False
							reg17059 = reg17058
						}
						reg17061 = reg17059
					} else {
						reg17060 := False
						reg17061 = reg17060
					}
					var reg17064 Obj
					if reg17061 == True {
						reg17062 := True
						reg17064 = reg17062
					} else {
						reg17063 := False
						reg17064 = reg17063
					}
					reg17066 = reg17064
				} else {
					reg17065 := False
					reg17066 = reg17065
				}
				var reg17069 Obj
				if reg17066 == True {
					reg17067 := True
					reg17069 = reg17067
				} else {
					reg17068 := False
					reg17069 = reg17068
				}
				reg17071 = reg17069
			} else {
				reg17070 := False
				reg17071 = reg17070
			}
			if reg17071 == True {
				reg17072 := MakeNumber(10)
				reg17073 := PrimNumberToString(reg17072)
				reg17074 := PrimTailString(V4601)
				reg17075 := PrimTailString(reg17074)
				reg17076 := __e.Call(__defun__shen_4proc_1nl, reg17075)
				reg17077 := PrimStringConcat(reg17073, reg17076)
				__ctx.Return(reg17077)
				return
			} else {
				reg17078 := __e.Call(__defun__shen_4_7string_2, V4601)
				if reg17078 == True {
					reg17079 := MakeNumber(0)
					reg17080 := PrimPos(V4601, reg17079)
					reg17081 := PrimTailString(V4601)
					reg17082 := __e.Call(__defun__shen_4proc_1nl, reg17081)
					reg17083 := PrimStringConcat(reg17080, reg17082)
					__ctx.Return(reg17083)
					return
				} else {
					reg17084 := MakeSymbol("shen.proc-nl")
					__ctx.TailApply(__defun__shen_4f__error, reg17084)
					return
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.proc-nl", value: __defun__shen_4proc_1nl})

	__defun__shen_4mkstr_1r = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4604 := __args[0]
		_ = V4604
		V4605 := __args[1]
		_ = V4605
		reg17086 := Nil
		reg17087 := PrimEqual(reg17086, V4605)
		if reg17087 == True {
			__ctx.Return(V4604)
			return
		} else {
			reg17088 := PrimIsPair(V4605)
			if reg17088 == True {
				reg17089 := MakeSymbol("shen.insert")
				reg17090 := PrimHead(V4605)
				reg17091 := Nil
				reg17092 := PrimCons(V4604, reg17091)
				reg17093 := PrimCons(reg17090, reg17092)
				reg17094 := PrimCons(reg17089, reg17093)
				reg17095 := PrimTail(V4605)
				__ctx.TailApply(__defun__shen_4mkstr_1r, reg17094, reg17095)
				return
			} else {
				reg17097 := MakeSymbol("shen.mkstr-r")
				__ctx.TailApply(__defun__shen_4f__error, reg17097)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.mkstr-r", value: __defun__shen_4mkstr_1r})

	__defun__shen_4insert = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4608 := __args[0]
		_ = V4608
		V4609 := __args[1]
		_ = V4609
		reg17099 := MakeString("")
		__ctx.TailApply(__defun__shen_4insert_1h, V4608, V4609, reg17099)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.insert", value: __defun__shen_4insert})

	__defun__shen_4insert_1h = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4615 := __args[0]
		_ = V4615
		V4616 := __args[1]
		_ = V4616
		V4617 := __args[2]
		_ = V4617
		reg17101 := MakeString("")
		reg17102 := PrimEqual(reg17101, V4616)
		if reg17102 == True {
			__ctx.Return(V4617)
			return
		} else {
			reg17103 := __e.Call(__defun__shen_4_7string_2, V4616)
			var reg17129 Obj
			if reg17103 == True {
				reg17104 := MakeString("~")
				reg17105 := MakeNumber(0)
				reg17106 := PrimPos(V4616, reg17105)
				reg17107 := PrimEqual(reg17104, reg17106)
				var reg17124 Obj
				if reg17107 == True {
					reg17108 := PrimTailString(V4616)
					reg17109 := __e.Call(__defun__shen_4_7string_2, reg17108)
					var reg17119 Obj
					if reg17109 == True {
						reg17110 := MakeString("A")
						reg17111 := PrimTailString(V4616)
						reg17112 := MakeNumber(0)
						reg17113 := PrimPos(reg17111, reg17112)
						reg17114 := PrimEqual(reg17110, reg17113)
						var reg17117 Obj
						if reg17114 == True {
							reg17115 := True
							reg17117 = reg17115
						} else {
							reg17116 := False
							reg17117 = reg17116
						}
						reg17119 = reg17117
					} else {
						reg17118 := False
						reg17119 = reg17118
					}
					var reg17122 Obj
					if reg17119 == True {
						reg17120 := True
						reg17122 = reg17120
					} else {
						reg17121 := False
						reg17122 = reg17121
					}
					reg17124 = reg17122
				} else {
					reg17123 := False
					reg17124 = reg17123
				}
				var reg17127 Obj
				if reg17124 == True {
					reg17125 := True
					reg17127 = reg17125
				} else {
					reg17126 := False
					reg17127 = reg17126
				}
				reg17129 = reg17127
			} else {
				reg17128 := False
				reg17129 = reg17128
			}
			if reg17129 == True {
				reg17130 := PrimTailString(V4616)
				reg17131 := PrimTailString(reg17130)
				reg17132 := MakeSymbol("shen.a")
				reg17133 := __e.Call(__defun__shen_4app, V4615, reg17131, reg17132)
				reg17134 := PrimStringConcat(V4617, reg17133)
				__ctx.Return(reg17134)
				return
			} else {
				reg17135 := __e.Call(__defun__shen_4_7string_2, V4616)
				var reg17161 Obj
				if reg17135 == True {
					reg17136 := MakeString("~")
					reg17137 := MakeNumber(0)
					reg17138 := PrimPos(V4616, reg17137)
					reg17139 := PrimEqual(reg17136, reg17138)
					var reg17156 Obj
					if reg17139 == True {
						reg17140 := PrimTailString(V4616)
						reg17141 := __e.Call(__defun__shen_4_7string_2, reg17140)
						var reg17151 Obj
						if reg17141 == True {
							reg17142 := MakeString("R")
							reg17143 := PrimTailString(V4616)
							reg17144 := MakeNumber(0)
							reg17145 := PrimPos(reg17143, reg17144)
							reg17146 := PrimEqual(reg17142, reg17145)
							var reg17149 Obj
							if reg17146 == True {
								reg17147 := True
								reg17149 = reg17147
							} else {
								reg17148 := False
								reg17149 = reg17148
							}
							reg17151 = reg17149
						} else {
							reg17150 := False
							reg17151 = reg17150
						}
						var reg17154 Obj
						if reg17151 == True {
							reg17152 := True
							reg17154 = reg17152
						} else {
							reg17153 := False
							reg17154 = reg17153
						}
						reg17156 = reg17154
					} else {
						reg17155 := False
						reg17156 = reg17155
					}
					var reg17159 Obj
					if reg17156 == True {
						reg17157 := True
						reg17159 = reg17157
					} else {
						reg17158 := False
						reg17159 = reg17158
					}
					reg17161 = reg17159
				} else {
					reg17160 := False
					reg17161 = reg17160
				}
				if reg17161 == True {
					reg17162 := PrimTailString(V4616)
					reg17163 := PrimTailString(reg17162)
					reg17164 := MakeSymbol("shen.r")
					reg17165 := __e.Call(__defun__shen_4app, V4615, reg17163, reg17164)
					reg17166 := PrimStringConcat(V4617, reg17165)
					__ctx.Return(reg17166)
					return
				} else {
					reg17167 := __e.Call(__defun__shen_4_7string_2, V4616)
					var reg17193 Obj
					if reg17167 == True {
						reg17168 := MakeString("~")
						reg17169 := MakeNumber(0)
						reg17170 := PrimPos(V4616, reg17169)
						reg17171 := PrimEqual(reg17168, reg17170)
						var reg17188 Obj
						if reg17171 == True {
							reg17172 := PrimTailString(V4616)
							reg17173 := __e.Call(__defun__shen_4_7string_2, reg17172)
							var reg17183 Obj
							if reg17173 == True {
								reg17174 := MakeString("S")
								reg17175 := PrimTailString(V4616)
								reg17176 := MakeNumber(0)
								reg17177 := PrimPos(reg17175, reg17176)
								reg17178 := PrimEqual(reg17174, reg17177)
								var reg17181 Obj
								if reg17178 == True {
									reg17179 := True
									reg17181 = reg17179
								} else {
									reg17180 := False
									reg17181 = reg17180
								}
								reg17183 = reg17181
							} else {
								reg17182 := False
								reg17183 = reg17182
							}
							var reg17186 Obj
							if reg17183 == True {
								reg17184 := True
								reg17186 = reg17184
							} else {
								reg17185 := False
								reg17186 = reg17185
							}
							reg17188 = reg17186
						} else {
							reg17187 := False
							reg17188 = reg17187
						}
						var reg17191 Obj
						if reg17188 == True {
							reg17189 := True
							reg17191 = reg17189
						} else {
							reg17190 := False
							reg17191 = reg17190
						}
						reg17193 = reg17191
					} else {
						reg17192 := False
						reg17193 = reg17192
					}
					if reg17193 == True {
						reg17194 := PrimTailString(V4616)
						reg17195 := PrimTailString(reg17194)
						reg17196 := MakeSymbol("shen.s")
						reg17197 := __e.Call(__defun__shen_4app, V4615, reg17195, reg17196)
						reg17198 := PrimStringConcat(V4617, reg17197)
						__ctx.Return(reg17198)
						return
					} else {
						reg17199 := __e.Call(__defun__shen_4_7string_2, V4616)
						if reg17199 == True {
							reg17200 := PrimTailString(V4616)
							reg17201 := MakeNumber(0)
							reg17202 := PrimPos(V4616, reg17201)
							reg17203 := PrimStringConcat(V4617, reg17202)
							__ctx.TailApply(__defun__shen_4insert_1h, V4615, reg17200, reg17203)
							return
						} else {
							reg17205 := MakeSymbol("shen.insert-h")
							__ctx.TailApply(__defun__shen_4f__error, reg17205)
							return
						}
					}
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.insert-h", value: __defun__shen_4insert_1h})

	__defun__shen_4app = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4621 := __args[0]
		_ = V4621
		V4622 := __args[1]
		_ = V4622
		V4623 := __args[2]
		_ = V4623
		reg17207 := __e.Call(__defun__shen_4arg_1_6str, V4621, V4623)
		reg17208 := PrimStringConcat(reg17207, V4622)
		__ctx.Return(reg17208)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.app", value: __defun__shen_4app})

	__defun__shen_4arg_1_6str = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4631 := __args[0]
		_ = V4631
		V4632 := __args[1]
		_ = V4632
		reg17209 := __e.Call(__defun__fail)
		reg17210 := PrimEqual(V4631, reg17209)
		if reg17210 == True {
			reg17211 := MakeString("...")
			__ctx.Return(reg17211)
			return
		} else {
			reg17212 := __e.Call(__defun__shen_4list_2, V4631)
			if reg17212 == True {
				__ctx.TailApply(__defun__shen_4list_1_6str, V4631, V4632)
				return
			} else {
				reg17214 := PrimIsString(V4631)
				if reg17214 == True {
					__ctx.TailApply(__defun__shen_4str_1_6str, V4631, V4632)
					return
				} else {
					reg17216 := PrimIsVector(V4631)
					if reg17216 == True {
						__ctx.TailApply(__defun__shen_4vector_1_6str, V4631, V4632)
						return
					} else {
						__ctx.TailApply(__defun__shen_4atom_1_6str, V4631)
						return
					}
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.arg->str", value: __defun__shen_4arg_1_6str})

	__defun__shen_4list_1_6str = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4635 := __args[0]
		_ = V4635
		V4636 := __args[1]
		_ = V4636
		reg17219 := MakeSymbol("shen.r")
		reg17220 := PrimEqual(reg17219, V4636)
		if reg17220 == True {
			reg17221 := MakeString("(")
			reg17222 := MakeSymbol("shen.r")
			reg17223 := __e.Call(__defun__shen_4maxseq)
			reg17224 := __e.Call(__defun__shen_4iter_1list, V4635, reg17222, reg17223)
			reg17225 := MakeString(")")
			reg17226 := __e.Call(__defun___8s, reg17224, reg17225)
			__ctx.TailApply(__defun___8s, reg17221, reg17226)
			return
		} else {
			reg17228 := MakeString("[")
			reg17229 := __e.Call(__defun__shen_4maxseq)
			reg17230 := __e.Call(__defun__shen_4iter_1list, V4635, V4636, reg17229)
			reg17231 := MakeString("]")
			reg17232 := __e.Call(__defun___8s, reg17230, reg17231)
			__ctx.TailApply(__defun___8s, reg17228, reg17232)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.list->str", value: __defun__shen_4list_1_6str})

	__defun__shen_4maxseq = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg17234 := MakeSymbol("*maximum-print-sequence-size*")
		reg17235 := PrimValue(reg17234)
		__ctx.Return(reg17235)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.maxseq", value: __defun__shen_4maxseq})

	__defun__shen_4iter_1list = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4650 := __args[0]
		_ = V4650
		V4651 := __args[1]
		_ = V4651
		V4652 := __args[2]
		_ = V4652
		reg17236 := Nil
		reg17237 := PrimEqual(reg17236, V4650)
		if reg17237 == True {
			reg17238 := MakeString("")
			__ctx.Return(reg17238)
			return
		} else {
			reg17239 := MakeNumber(0)
			reg17240 := PrimEqual(reg17239, V4652)
			if reg17240 == True {
				reg17241 := MakeString("... etc")
				__ctx.Return(reg17241)
				return
			} else {
				reg17242 := PrimIsPair(V4650)
				var reg17250 Obj
				if reg17242 == True {
					reg17243 := Nil
					reg17244 := PrimTail(V4650)
					reg17245 := PrimEqual(reg17243, reg17244)
					var reg17248 Obj
					if reg17245 == True {
						reg17246 := True
						reg17248 = reg17246
					} else {
						reg17247 := False
						reg17248 = reg17247
					}
					reg17250 = reg17248
				} else {
					reg17249 := False
					reg17250 = reg17249
				}
				if reg17250 == True {
					reg17251 := PrimHead(V4650)
					__ctx.TailApply(__defun__shen_4arg_1_6str, reg17251, V4651)
					return
				} else {
					reg17253 := PrimIsPair(V4650)
					if reg17253 == True {
						reg17254 := PrimHead(V4650)
						reg17255 := __e.Call(__defun__shen_4arg_1_6str, reg17254, V4651)
						reg17256 := MakeString(" ")
						reg17257 := PrimTail(V4650)
						reg17258 := MakeNumber(1)
						reg17259 := PrimNumberSubtract(V4652, reg17258)
						reg17260 := __e.Call(__defun__shen_4iter_1list, reg17257, V4651, reg17259)
						reg17261 := __e.Call(__defun___8s, reg17256, reg17260)
						__ctx.TailApply(__defun___8s, reg17255, reg17261)
						return
					} else {
						reg17263 := MakeString("|")
						reg17264 := MakeString(" ")
						reg17265 := __e.Call(__defun__shen_4arg_1_6str, V4650, V4651)
						reg17266 := __e.Call(__defun___8s, reg17264, reg17265)
						__ctx.TailApply(__defun___8s, reg17263, reg17266)
						return
					}
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.iter-list", value: __defun__shen_4iter_1list})

	__defun__shen_4str_1_6str = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4659 := __args[0]
		_ = V4659
		V4660 := __args[1]
		_ = V4660
		reg17268 := MakeSymbol("shen.a")
		reg17269 := PrimEqual(reg17268, V4660)
		if reg17269 == True {
			__ctx.Return(V4659)
			return
		} else {
			reg17270 := MakeNumber(34)
			reg17271 := PrimNumberToString(reg17270)
			reg17272 := MakeNumber(34)
			reg17273 := PrimNumberToString(reg17272)
			reg17274 := __e.Call(__defun___8s, V4659, reg17273)
			__ctx.TailApply(__defun___8s, reg17271, reg17274)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.str->str", value: __defun__shen_4str_1_6str})

	__defun__shen_4vector_1_6str = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4663 := __args[0]
		_ = V4663
		V4664 := __args[1]
		_ = V4664
		reg17277 := __e.Call(__defun__shen_4print_1vector_2, V4663)
		if reg17277 == True {
			reg17278 := MakeNumber(0)
			reg17279 := PrimVectorGet(V4663, reg17278)
			reg17280 := __e.Call(__defun__function, reg17279)
			f17276 := reg17280
			_ = f17276
			__ctx.TailApply(f17276, V4663)
			return
		} else {
			reg17282 := __e.Call(__defun__vector_2, V4663)
			if reg17282 == True {
				reg17283 := MakeString("<")
				reg17284 := MakeNumber(1)
				reg17285 := __e.Call(__defun__shen_4maxseq)
				reg17286 := __e.Call(__defun__shen_4iter_1vector, V4663, reg17284, V4664, reg17285)
				reg17287 := MakeString(">")
				reg17288 := __e.Call(__defun___8s, reg17286, reg17287)
				__ctx.TailApply(__defun___8s, reg17283, reg17288)
				return
			} else {
				reg17290 := MakeString("<")
				reg17291 := MakeString("<")
				reg17292 := MakeNumber(0)
				reg17293 := __e.Call(__defun__shen_4maxseq)
				reg17294 := __e.Call(__defun__shen_4iter_1vector, V4663, reg17292, V4664, reg17293)
				reg17295 := MakeString(">>")
				reg17296 := __e.Call(__defun___8s, reg17294, reg17295)
				reg17297 := __e.Call(__defun___8s, reg17291, reg17296)
				__ctx.TailApply(__defun___8s, reg17290, reg17297)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.vector->str", value: __defun__shen_4vector_1_6str})

	__defun__shen_4empty_1absvector_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4666 := __args[0]
		_ = V4666
		reg17299 := MakeSymbol("shen.*empty-absvector*")
		reg17300 := PrimValue(reg17299)
		reg17301 := PrimEqual(V4666, reg17300)
		__ctx.Return(reg17301)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.empty-absvector?", value: __defun__shen_4empty_1absvector_2})

	__defun__shen_4print_1vector_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4668 := __args[0]
		_ = V4668
		reg17302 := __e.Call(__defun__shen_4empty_1absvector_2, V4668)
		reg17303 := PrimNot(reg17302)
		if reg17303 == True {
			reg17304 := MakeNumber(0)
			reg17305 := PrimVectorGet(V4668, reg17304)
			First := reg17305
			_ = First
			reg17306 := MakeSymbol("shen.tuple")
			reg17307 := PrimEqual(First, reg17306)
			var reg17334 Obj
			if reg17307 == True {
				reg17308 := True
				reg17334 = reg17308
			} else {
				reg17309 := MakeSymbol("shen.pvar")
				reg17310 := PrimEqual(First, reg17309)
				var reg17330 Obj
				if reg17310 == True {
					reg17311 := True
					reg17330 = reg17311
				} else {
					reg17312 := MakeSymbol("shen.dictionary")
					reg17313 := PrimEqual(First, reg17312)
					var reg17326 Obj
					if reg17313 == True {
						reg17314 := True
						reg17326 = reg17314
					} else {
						reg17315 := PrimIsNumber(First)
						reg17316 := PrimNot(reg17315)
						var reg17322 Obj
						if reg17316 == True {
							reg17317 := __e.Call(__defun__shen_4fbound_2, First)
							var reg17320 Obj
							if reg17317 == True {
								reg17318 := True
								reg17320 = reg17318
							} else {
								reg17319 := False
								reg17320 = reg17319
							}
							reg17322 = reg17320
						} else {
							reg17321 := False
							reg17322 = reg17321
						}
						var reg17325 Obj
						if reg17322 == True {
							reg17323 := True
							reg17325 = reg17323
						} else {
							reg17324 := False
							reg17325 = reg17324
						}
						reg17326 = reg17325
					}
					var reg17329 Obj
					if reg17326 == True {
						reg17327 := True
						reg17329 = reg17327
					} else {
						reg17328 := False
						reg17329 = reg17328
					}
					reg17330 = reg17329
				}
				var reg17333 Obj
				if reg17330 == True {
					reg17331 := True
					reg17333 = reg17331
				} else {
					reg17332 := False
					reg17333 = reg17332
				}
				reg17334 = reg17333
			}
			if reg17334 == True {
				reg17335 := True
				__ctx.Return(reg17335)
				return
			} else {
				reg17336 := False
				__ctx.Return(reg17336)
				return
			}
		} else {
			reg17337 := False
			__ctx.Return(reg17337)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.print-vector?", value: __defun__shen_4print_1vector_2})

	__defun__shen_4fbound_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4670 := __args[0]
		_ = V4670
		reg17338 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg17339 := __e.Call(__defun__shen_4lookup_1func, V4670)
			_ = reg17339
			reg17340 := True
			__ctx.Return(reg17340)
			return
		}, 0)
		reg17341 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg17342 := False
			__ctx.Return(reg17342)
			return
		}, 1)
		reg17343 := __e.Try(reg17338).Catch(reg17341)
		__ctx.Return(reg17343)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.fbound?", value: __defun__shen_4fbound_2})

	__defun__shen_4tuple = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4672 := __args[0]
		_ = V4672
		reg17344 := MakeString("(@p ")
		reg17345 := MakeNumber(1)
		reg17346 := PrimVectorGet(V4672, reg17345)
		reg17347 := MakeString(" ")
		reg17348 := MakeNumber(2)
		reg17349 := PrimVectorGet(V4672, reg17348)
		reg17350 := MakeString(")")
		reg17351 := MakeSymbol("shen.s")
		reg17352 := __e.Call(__defun__shen_4app, reg17349, reg17350, reg17351)
		reg17353 := PrimStringConcat(reg17347, reg17352)
		reg17354 := MakeSymbol("shen.s")
		reg17355 := __e.Call(__defun__shen_4app, reg17346, reg17353, reg17354)
		reg17356 := PrimStringConcat(reg17344, reg17355)
		__ctx.Return(reg17356)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.tuple", value: __defun__shen_4tuple})

	__defun__shen_4dictionary = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4674 := __args[0]
		_ = V4674
		reg17357 := MakeString("(dict ...)")
		__ctx.Return(reg17357)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dictionary", value: __defun__shen_4dictionary})

	__defun__shen_4iter_1vector = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4685 := __args[0]
		_ = V4685
		V4686 := __args[1]
		_ = V4686
		V4687 := __args[2]
		_ = V4687
		V4688 := __args[3]
		_ = V4688
		reg17358 := MakeNumber(0)
		reg17359 := PrimEqual(reg17358, V4688)
		if reg17359 == True {
			reg17360 := MakeString("... etc")
			__ctx.Return(reg17360)
			return
		} else {
			reg17361 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg17362 := PrimVectorGet(V4685, V4686)
				__ctx.Return(reg17362)
				return
			}, 0)
			reg17363 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg17364 := MakeSymbol("shen.out-of-bounds")
				__ctx.Return(reg17364)
				return
			}, 1)
			reg17365 := __e.Try(reg17361).Catch(reg17363)
			Item := reg17365
			_ = Item
			reg17366 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg17367 := MakeNumber(1)
				reg17368 := PrimNumberAdd(V4686, reg17367)
				reg17369 := PrimVectorGet(V4685, reg17368)
				__ctx.Return(reg17369)
				return
			}, 0)
			reg17370 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg17371 := MakeSymbol("shen.out-of-bounds")
				__ctx.Return(reg17371)
				return
			}, 1)
			reg17372 := __e.Try(reg17366).Catch(reg17370)
			Next := reg17372
			_ = Next
			reg17373 := MakeSymbol("shen.out-of-bounds")
			reg17374 := PrimEqual(Item, reg17373)
			if reg17374 == True {
				reg17375 := MakeString("")
				__ctx.Return(reg17375)
				return
			} else {
				reg17376 := MakeSymbol("shen.out-of-bounds")
				reg17377 := PrimEqual(Next, reg17376)
				if reg17377 == True {
					__ctx.TailApply(__defun__shen_4arg_1_6str, Item, V4687)
					return
				} else {
					reg17379 := __e.Call(__defun__shen_4arg_1_6str, Item, V4687)
					reg17380 := MakeString(" ")
					reg17381 := MakeNumber(1)
					reg17382 := PrimNumberAdd(V4686, reg17381)
					reg17383 := MakeNumber(1)
					reg17384 := PrimNumberSubtract(V4688, reg17383)
					reg17385 := __e.Call(__defun__shen_4iter_1vector, V4685, reg17382, V4687, reg17384)
					reg17386 := __e.Call(__defun___8s, reg17380, reg17385)
					__ctx.TailApply(__defun___8s, reg17379, reg17386)
					return
				}
			}
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.iter-vector", value: __defun__shen_4iter_1vector})

	__defun__shen_4atom_1_6str = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4690 := __args[0]
		_ = V4690
		reg17388 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg17389 := PrimStr(V4690)
			__ctx.Return(reg17389)
			return
		}, 0)
		reg17390 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			__ctx.TailApply(__defun__shen_4funexstring)
			return
		}, 1)
		reg17392 := __e.Try(reg17388).Catch(reg17390)
		__ctx.Return(reg17392)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.atom->str", value: __defun__shen_4atom_1_6str})

	__defun__shen_4funexstring = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg17393 := MakeString("\x10")
		reg17394 := MakeString("f")
		reg17395 := MakeString("u")
		reg17396 := MakeString("n")
		reg17397 := MakeString("e")
		reg17398 := MakeString("x")
		reg17399 := PrimIntern(reg17398)
		reg17400 := __e.Call(__defun__gensym, reg17399)
		reg17401 := MakeSymbol("shen.a")
		reg17402 := __e.Call(__defun__shen_4arg_1_6str, reg17400, reg17401)
		reg17403 := MakeString("\x11")
		reg17404 := __e.Call(__defun___8s, reg17402, reg17403)
		reg17405 := __e.Call(__defun___8s, reg17397, reg17404)
		reg17406 := __e.Call(__defun___8s, reg17396, reg17405)
		reg17407 := __e.Call(__defun___8s, reg17395, reg17406)
		reg17408 := __e.Call(__defun___8s, reg17394, reg17407)
		__ctx.TailApply(__defun___8s, reg17393, reg17408)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.funexstring", value: __defun__shen_4funexstring})

	__defun__shen_4list_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4692 := __args[0]
		_ = V4692
		reg17410 := __e.Call(__defun__empty_2, V4692)
		if reg17410 == True {
			reg17411 := True
			__ctx.Return(reg17411)
			return
		} else {
			reg17412 := PrimIsPair(V4692)
			if reg17412 == True {
				reg17413 := True
				__ctx.Return(reg17413)
				return
			} else {
				reg17414 := False
				__ctx.Return(reg17414)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.list?", value: __defun__shen_4list_2})

}
