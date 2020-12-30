package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen2719 := MakeNative(func(__e Evaluator) {
			V1920 := __e.Get(1)
			_ = V1920
			__e.TailApply(V1920)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("thaw"), gen2719)

		gen2725 := MakeNative(func(__e Evaluator) {
			V1922 := __e.Get(1)
			_ = V1922
			gen2720 := MakeNative(func(__e Evaluator) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(ShenFunc(symmacroexpand), Y)

				return
			}, 1)
			gen2721 := Call(__e, ShenFunc(symshen_4walk), gen2720, V1922)

			Macroexpand := gen2721
			gen2724 := Call(__e, ShenFunc(symshen_4packaged_2), Macroexpand)

			if True == gen2724 {
				gen2722 := MakeNative(func(__e Evaluator) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(ShenFunc(symshen_4eval_1without_1macros), Z)

					return
				}, 1)
				gen2723 := Call(__e, ShenFunc(symshen_4package_1contents), Macroexpand)

				__e.TailApply(ShenFunc(symmap), gen2722, gen2723)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4eval_1without_1macros), Macroexpand)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("eval"), gen2725)

		gen2728 := MakeNative(func(__e Evaluator) {
			V1924 := __e.Get(1)
			_ = V1924
			gen2726 := Call(__e, ShenFunc(symshen_4proc_1input_7), V1924)

			gen2727 := Call(__e, ShenFunc(symshen_4elim_1def), gen2726)

			__e.TailApply(ShenFunc(symeval_1kl), gen2727)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.eval-without-macros"), gen2728)

		gen2775 := MakeNative(func(__e Evaluator) {
			V1926 := __e.Get(1)
			_ = V1926
			gen2773 := Call(__e, ShenFunc(symcons_2), V1926)

			var gen2774 Obj
			if True == gen2773 {
				gen2770 := Call(__e, ShenFunc(symhd), V1926)

				gen2771 := Call(__e, ShenFunc(sym_a), MakeSymbol("input+"), gen2770)

				var gen2772 Obj
				if True == gen2771 {
					gen2767 := Call(__e, ShenFunc(symtl), V1926)

					gen2768 := Call(__e, ShenFunc(symcons_2), gen2767)

					var gen2769 Obj
					if True == gen2768 {
						gen2763 := Call(__e, ShenFunc(symtl), V1926)

						gen2764 := Call(__e, ShenFunc(symtl), gen2763)

						gen2765 := Call(__e, ShenFunc(symcons_2), gen2764)

						var gen2766 Obj
						if True == gen2765 {
							gen2759 := Call(__e, ShenFunc(symtl), V1926)

							gen2760 := Call(__e, ShenFunc(symtl), gen2759)

							gen2761 := Call(__e, ShenFunc(symtl), gen2760)

							gen2762 := Call(__e, ShenFunc(sym_a), Nil, gen2761)

							if True == gen2762 {
								gen2766 = True
							} else {
								gen2766 = False
							}

						} else {
							gen2766 = False
						}
						if True == gen2766 {
							gen2769 = True
						} else {
							gen2769 = False
						}

					} else {
						gen2769 = False
					}
					if True == gen2769 {
						gen2772 = True
					} else {
						gen2772 = False
					}

				} else {
					gen2772 = False
				}
				if True == gen2772 {
					gen2774 = True
				} else {
					gen2774 = False
				}

			} else {
				gen2774 = False
			}
			if True == gen2774 {
				gen2753 := Call(__e, ShenFunc(symtl), V1926)

				gen2754 := Call(__e, ShenFunc(symhd), gen2753)

				gen2755 := Call(__e, ShenFunc(symshen_4rcons__form), gen2754)

				gen2756 := Call(__e, ShenFunc(symtl), V1926)

				gen2757 := Call(__e, ShenFunc(symtl), gen2756)

				gen2758 := Call(__e, ShenFunc(symcons), gen2755, gen2757)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("input+"), gen2758)

				return

			} else {
				gen2751 := Call(__e, ShenFunc(symcons_2), V1926)

				var gen2752 Obj
				if True == gen2751 {
					gen2748 := Call(__e, ShenFunc(symhd), V1926)

					gen2749 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.read+"), gen2748)

					var gen2750 Obj
					if True == gen2749 {
						gen2745 := Call(__e, ShenFunc(symtl), V1926)

						gen2746 := Call(__e, ShenFunc(symcons_2), gen2745)

						var gen2747 Obj
						if True == gen2746 {
							gen2741 := Call(__e, ShenFunc(symtl), V1926)

							gen2742 := Call(__e, ShenFunc(symtl), gen2741)

							gen2743 := Call(__e, ShenFunc(symcons_2), gen2742)

							var gen2744 Obj
							if True == gen2743 {
								gen2737 := Call(__e, ShenFunc(symtl), V1926)

								gen2738 := Call(__e, ShenFunc(symtl), gen2737)

								gen2739 := Call(__e, ShenFunc(symtl), gen2738)

								gen2740 := Call(__e, ShenFunc(sym_a), Nil, gen2739)

								if True == gen2740 {
									gen2744 = True
								} else {
									gen2744 = False
								}

							} else {
								gen2744 = False
							}
							if True == gen2744 {
								gen2747 = True
							} else {
								gen2747 = False
							}

						} else {
							gen2747 = False
						}
						if True == gen2747 {
							gen2750 = True
						} else {
							gen2750 = False
						}

					} else {
						gen2750 = False
					}
					if True == gen2750 {
						gen2752 = True
					} else {
						gen2752 = False
					}

				} else {
					gen2752 = False
				}
				if True == gen2752 {
					gen2731 := Call(__e, ShenFunc(symtl), V1926)

					gen2732 := Call(__e, ShenFunc(symhd), gen2731)

					gen2733 := Call(__e, ShenFunc(symshen_4rcons__form), gen2732)

					gen2734 := Call(__e, ShenFunc(symtl), V1926)

					gen2735 := Call(__e, ShenFunc(symtl), gen2734)

					gen2736 := Call(__e, ShenFunc(symcons), gen2733, gen2735)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.read+"), gen2736)

					return

				} else {
					gen2730 := Call(__e, ShenFunc(symcons_2), V1926)

					if True == gen2730 {
						gen2729 := MakeNative(func(__e Evaluator) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(ShenFunc(symshen_4proc_1input_7), Z)

							return
						}, 1)
						__e.TailApply(ShenFunc(symmap), gen2729, V1926)

						return

					} else {
						__e.Return(V1926)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.proc-input+"), gen2775)

		gen2818 := MakeNative(func(__e Evaluator) {
			V1928 := __e.Get(1)
			_ = V1928
			gen2816 := Call(__e, ShenFunc(symcons_2), V1928)

			var gen2817 Obj
			if True == gen2816 {
				gen2813 := Call(__e, ShenFunc(symhd), V1928)

				gen2814 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), gen2813)

				var gen2815 Obj
				if True == gen2814 {
					gen2811 := Call(__e, ShenFunc(symtl), V1928)

					gen2812 := Call(__e, ShenFunc(symcons_2), gen2811)

					if True == gen2812 {
						gen2815 = True
					} else {
						gen2815 = False
					}

				} else {
					gen2815 = False
				}
				if True == gen2815 {
					gen2817 = True
				} else {
					gen2817 = False
				}

			} else {
				gen2817 = False
			}
			if True == gen2817 {
				gen2807 := Call(__e, ShenFunc(symtl), V1928)

				gen2808 := Call(__e, ShenFunc(symhd), gen2807)

				gen2809 := Call(__e, ShenFunc(symtl), V1928)

				gen2810 := Call(__e, ShenFunc(symtl), gen2809)

				__e.TailApply(ShenFunc(symshen_4shen_1_6kl), gen2808, gen2810)

				return

			} else {
				gen2805 := Call(__e, ShenFunc(symcons_2), V1928)

				var gen2806 Obj
				if True == gen2805 {
					gen2802 := Call(__e, ShenFunc(symhd), V1928)

					gen2803 := Call(__e, ShenFunc(sym_a), MakeSymbol("defmacro"), gen2802)

					var gen2804 Obj
					if True == gen2803 {
						gen2800 := Call(__e, ShenFunc(symtl), V1928)

						gen2801 := Call(__e, ShenFunc(symcons_2), gen2800)

						if True == gen2801 {
							gen2804 = True
						} else {
							gen2804 = False
						}

					} else {
						gen2804 = False
					}
					if True == gen2804 {
						gen2806 = True
					} else {
						gen2806 = False
					}

				} else {
					gen2806 = False
				}
				if True == gen2806 {
					gen2786 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), Nil)

					gen2787 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen2786)

					gen2788 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen2787)

					Default := gen2788
					gen2789 := Call(__e, ShenFunc(symtl), V1928)

					gen2790 := Call(__e, ShenFunc(symhd), gen2789)

					gen2791 := Call(__e, ShenFunc(symtl), V1928)

					gen2792 := Call(__e, ShenFunc(symtl), gen2791)

					gen2793 := Call(__e, ShenFunc(symappend), gen2792, Default)

					gen2794 := Call(__e, ShenFunc(symcons), gen2790, gen2793)

					gen2795 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen2794)

					gen2796 := Call(__e, ShenFunc(symshen_4elim_1def), gen2795)

					Def := gen2796
					gen2797 := Call(__e, ShenFunc(symtl), V1928)

					gen2798 := Call(__e, ShenFunc(symhd), gen2797)

					gen2799 := Call(__e, ShenFunc(symshen_4add_1macro), gen2798)

					MacroAdd := gen2799
					_ = MacroAdd
					__e.Return(Def)
					return

				} else {
					gen2784 := Call(__e, ShenFunc(symcons_2), V1928)

					var gen2785 Obj
					if True == gen2784 {
						gen2781 := Call(__e, ShenFunc(symhd), V1928)

						gen2782 := Call(__e, ShenFunc(sym_a), MakeSymbol("defcc"), gen2781)

						var gen2783 Obj
						if True == gen2782 {
							gen2779 := Call(__e, ShenFunc(symtl), V1928)

							gen2780 := Call(__e, ShenFunc(symcons_2), gen2779)

							if True == gen2780 {
								gen2783 = True
							} else {
								gen2783 = False
							}

						} else {
							gen2783 = False
						}
						if True == gen2783 {
							gen2785 = True
						} else {
							gen2785 = False
						}

					} else {
						gen2785 = False
					}
					if True == gen2785 {
						gen2778 := Call(__e, ShenFunc(symshen_4yacc), V1928)

						__e.TailApply(ShenFunc(symshen_4elim_1def), gen2778)

						return

					} else {
						gen2777 := Call(__e, ShenFunc(symcons_2), V1928)

						if True == gen2777 {
							gen2776 := MakeNative(func(__e Evaluator) {
								Z := __e.Get(1)
								_ = Z
								__e.TailApply(ShenFunc(symshen_4elim_1def), Z)

								return
							}, 1)
							__e.TailApply(ShenFunc(symmap), gen2776, V1928)

							return

						} else {
							__e.Return(V1928)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.elim-def"), gen2818)

		gen2827 := MakeNative(func(__e Evaluator) {
			V1930 := __e.Get(1)
			_ = V1930
			gen2819 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*macroreg*"))

			MacroReg := gen2819
			gen2820 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*macroreg*"))

			gen2821 := Call(__e, ShenFunc(symadjoin), V1930, gen2820)

			gen2822 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*macroreg*"), gen2821)

			NewMacroReg := gen2822
			gen2826 := Call(__e, ShenFunc(sym_a), MacroReg, NewMacroReg)

			if True == gen2826 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen2823 := Call(__e, ShenFunc(symfunction), V1930)

				gen2824 := Call(__e, ShenFunc(symvalue), MakeSymbol("*macros*"))

				gen2825 := Call(__e, ShenFunc(symcons), gen2823, gen2824)

				__e.TailApply(ShenFunc(symset), MakeSymbol("*macros*"), gen2825)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.add-macro"), gen2827)

		gen2839 := MakeNative(func(__e Evaluator) {
			V1938 := __e.Get(1)
			_ = V1938
			gen2837 := Call(__e, ShenFunc(symcons_2), V1938)

			var gen2838 Obj
			if True == gen2837 {
				gen2834 := Call(__e, ShenFunc(symhd), V1938)

				gen2835 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen2834)

				var gen2836 Obj
				if True == gen2835 {
					gen2831 := Call(__e, ShenFunc(symtl), V1938)

					gen2832 := Call(__e, ShenFunc(symcons_2), gen2831)

					var gen2833 Obj
					if True == gen2832 {
						gen2828 := Call(__e, ShenFunc(symtl), V1938)

						gen2829 := Call(__e, ShenFunc(symtl), gen2828)

						gen2830 := Call(__e, ShenFunc(symcons_2), gen2829)

						if True == gen2830 {
							gen2833 = True
						} else {
							gen2833 = False
						}

					} else {
						gen2833 = False
					}
					if True == gen2833 {
						gen2836 = True
					} else {
						gen2836 = False
					}

				} else {
					gen2836 = False
				}
				if True == gen2836 {
					gen2838 = True
				} else {
					gen2838 = False
				}

			} else {
				gen2838 = False
			}
			if True == gen2838 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.packaged?"), gen2839)

		gen2845 := MakeNative(func(__e Evaluator) {
			V1940 := __e.Get(1)
			_ = V1940
			gen2842 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				gen2840 := Call(__e, ShenFunc(symshen_4app), V1940, MakeString(" has not been used.\n"), MakeSymbol("shen.a"))

				gen2841 := Call(__e, ShenFunc(symcn), MakeString("package "), gen2840)

				__e.TailApply(ShenFunc(symsimple_1error), gen2841)

				return

			}, 1)
			gen2844 := MakeNative(func(__e Evaluator) {
				gen2843 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1940, MakeSymbol("shen.external-symbols"), gen2843)

				return

			}, 0)
			__e.Return(Try(__e, gen2844).Catch(gen2842))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("external"), gen2845)

		gen2851 := MakeNative(func(__e Evaluator) {
			V1942 := __e.Get(1)
			_ = V1942
			gen2848 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				gen2846 := Call(__e, ShenFunc(symshen_4app), V1942, MakeString(" has not been used.\n"), MakeSymbol("shen.a"))

				gen2847 := Call(__e, ShenFunc(symcn), MakeString("package "), gen2846)

				__e.TailApply(ShenFunc(symsimple_1error), gen2847)

				return

			}, 1)
			gen2850 := MakeNative(func(__e Evaluator) {
				gen2849 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1942, MakeSymbol("shen.internal-symbols"), gen2849)

				return

			}, 0)
			__e.Return(Try(__e, gen2850).Catch(gen2848))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("internal"), gen2851)

		gen2894 := MakeNative(func(__e Evaluator) {
			V1946 := __e.Get(1)
			_ = V1946
			gen2892 := Call(__e, ShenFunc(symcons_2), V1946)

			var gen2893 Obj
			if True == gen2892 {
				gen2889 := Call(__e, ShenFunc(symhd), V1946)

				gen2890 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen2889)

				var gen2891 Obj
				if True == gen2890 {
					gen2886 := Call(__e, ShenFunc(symtl), V1946)

					gen2887 := Call(__e, ShenFunc(symcons_2), gen2886)

					var gen2888 Obj
					if True == gen2887 {
						gen2882 := Call(__e, ShenFunc(symtl), V1946)

						gen2883 := Call(__e, ShenFunc(symhd), gen2882)

						gen2884 := Call(__e, ShenFunc(sym_a), MakeSymbol("null"), gen2883)

						var gen2885 Obj
						if True == gen2884 {
							gen2879 := Call(__e, ShenFunc(symtl), V1946)

							gen2880 := Call(__e, ShenFunc(symtl), gen2879)

							gen2881 := Call(__e, ShenFunc(symcons_2), gen2880)

							if True == gen2881 {
								gen2885 = True
							} else {
								gen2885 = False
							}

						} else {
							gen2885 = False
						}
						if True == gen2885 {
							gen2888 = True
						} else {
							gen2888 = False
						}

					} else {
						gen2888 = False
					}
					if True == gen2888 {
						gen2891 = True
					} else {
						gen2891 = False
					}

				} else {
					gen2891 = False
				}
				if True == gen2891 {
					gen2893 = True
				} else {
					gen2893 = False
				}

			} else {
				gen2893 = False
			}
			if True == gen2893 {
				gen2877 := Call(__e, ShenFunc(symtl), V1946)

				gen2878 := Call(__e, ShenFunc(symtl), gen2877)

				__e.TailApply(ShenFunc(symtl), gen2878)

				return

			} else {
				gen2875 := Call(__e, ShenFunc(symcons_2), V1946)

				var gen2876 Obj
				if True == gen2875 {
					gen2872 := Call(__e, ShenFunc(symhd), V1946)

					gen2873 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen2872)

					var gen2874 Obj
					if True == gen2873 {
						gen2869 := Call(__e, ShenFunc(symtl), V1946)

						gen2870 := Call(__e, ShenFunc(symcons_2), gen2869)

						var gen2871 Obj
						if True == gen2870 {
							gen2866 := Call(__e, ShenFunc(symtl), V1946)

							gen2867 := Call(__e, ShenFunc(symtl), gen2866)

							gen2868 := Call(__e, ShenFunc(symcons_2), gen2867)

							if True == gen2868 {
								gen2871 = True
							} else {
								gen2871 = False
							}

						} else {
							gen2871 = False
						}
						if True == gen2871 {
							gen2874 = True
						} else {
							gen2874 = False
						}

					} else {
						gen2874 = False
					}
					if True == gen2874 {
						gen2876 = True
					} else {
						gen2876 = False
					}

				} else {
					gen2876 = False
				}
				if True == gen2876 {
					gen2852 := Call(__e, ShenFunc(symtl), V1946)

					gen2853 := Call(__e, ShenFunc(symhd), gen2852)

					gen2854 := Call(__e, ShenFunc(symstr), gen2853)

					gen2855 := Call(__e, ShenFunc(symcn), gen2854, MakeString("."))

					gen2856 := Call(__e, ShenFunc(symintern), gen2855)

					PackageNameDot := gen2856
					gen2857 := Call(__e, ShenFunc(symexplode), PackageNameDot)

					ExpPackageNameDot := gen2857
					gen2858 := Call(__e, ShenFunc(symtl), V1946)

					gen2859 := Call(__e, ShenFunc(symhd), gen2858)

					gen2860 := Call(__e, ShenFunc(symtl), V1946)

					gen2861 := Call(__e, ShenFunc(symtl), gen2860)

					gen2862 := Call(__e, ShenFunc(symhd), gen2861)

					gen2863 := Call(__e, ShenFunc(symtl), V1946)

					gen2864 := Call(__e, ShenFunc(symtl), gen2863)

					gen2865 := Call(__e, ShenFunc(symtl), gen2864)

					__e.TailApply(ShenFunc(symshen_4packageh), gen2859, gen2862, gen2865, ExpPackageNameDot)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.package-contents"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.package-contents"), gen2894)

		gen2898 := MakeNative(func(__e Evaluator) {
			V1949 := __e.Get(1)
			_ = V1949
			V1950 := __e.Get(2)
			_ = V1950
			gen2897 := Call(__e, ShenFunc(symcons_2), V1950)

			if True == gen2897 {
				gen2895 := MakeNative(func(__e Evaluator) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(ShenFunc(symshen_4walk), V1949, Z)

					return
				}, 1)
				gen2896 := Call(__e, ShenFunc(symmap), gen2895, V1950)

				__e.TailApply(V1949, gen2896)

				return

			} else {
				__e.TailApply(V1949, V1950)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.walk"), gen2898)

		gen2908 := MakeNative(func(__e Evaluator) {
			V1954 := __e.Get(1)
			_ = V1954
			V1955 := __e.Get(2)
			_ = V1955
			V1956 := __e.Get(3)
			_ = V1956
			gen2899 := Call(__e, ShenFunc(symcons), Nil, Nil)

			gen2900 := Call(__e, ShenFunc(symcons), V1955, gen2899)

			gen2901 := Call(__e, V1954, gen2900)

			O := gen2901
			gen2905 := Call(__e, ShenFunc(symfail))

			gen2906 := Call(__e, ShenFunc(sym_a), gen2905, O)

			var gen2907 Obj
			if True == gen2906 {
				gen2907 = True
			} else {
				gen2902 := Call(__e, ShenFunc(symhd), O)

				gen2903 := Call(__e, ShenFunc(symempty_2), gen2902)

				gen2904 := Call(__e, ShenFunc(symnot), gen2903)

				if True == gen2904 {
					gen2907 = True
				} else {
					gen2907 = False
				}

			}
			if True == gen2907 {
				__e.TailApply(V1956, O)

				return
			} else {
				__e.TailApply(ShenFunc(symshen_4hdtl), O)

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("compile"), gen2908)

		gen2910 := MakeNative(func(__e Evaluator) {
			V1959 := __e.Get(1)
			_ = V1959
			V1960 := __e.Get(2)
			_ = V1960
			gen2909 := Call(__e, V1959, V1960)

			if True == gen2909 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				__e.Return(V1960)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fail-if"), gen2910)

		gen2911 := MakeNative(func(__e Evaluator) {
			V1963 := __e.Get(1)
			_ = V1963
			V1964 := __e.Get(2)
			_ = V1964
			__e.TailApply(ShenFunc(symcn), V1963, V1964)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("@s"), gen2911)

		gen2912 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tc?"), gen2912)

		gen2917 := MakeNative(func(__e Evaluator) {
			V1966 := __e.Get(1)
			_ = V1966
			gen2914 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				gen2913 := Call(__e, ShenFunc(symshen_4app), V1966, MakeString(" not found.\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen2913)

				return

			}, 1)
			gen2916 := MakeNative(func(__e Evaluator) {
				gen2915 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1966, MakeSymbol("shen.source"), gen2915)

				return

			}, 0)
			__e.Return(Try(__e, gen2916).Catch(gen2914))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("ps"), gen2917)

		gen2918 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*stinput*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("stinput"), gen2918)

		gen2925 := MakeNative(func(__e Evaluator) {
			V1968 := __e.Get(1)
			_ = V1968
			gen2919 := Call(__e, ShenFunc(sym_7), V1968, MakeNumber(1))

			gen2920 := Call(__e, ShenFunc(symabsvector), gen2919)

			Vector := gen2920
			gen2921 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(0), V1968)

			ZeroStamp := gen2921
			gen2923 := Call(__e, ShenFunc(sym_a), V1968, MakeNumber(0))

			var gen2924 Obj
			if True == gen2923 {
				gen2924 = ZeroStamp
			} else {
				gen2922 := Call(__e, ShenFunc(symfail))

				gen2924 = Call(__e, ShenFunc(symshen_4fillvector), ZeroStamp, MakeNumber(1), V1968, gen2922)

			}
			Standard := gen2924
			__e.Return(Standard)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("vector"), gen2925)

		gen2929 := MakeNative(func(__e Evaluator) {
			V1974 := __e.Get(1)
			_ = V1974
			V1975 := __e.Get(2)
			_ = V1975
			V1976 := __e.Get(3)
			_ = V1976
			V1977 := __e.Get(4)
			_ = V1977
			gen2928 := Call(__e, ShenFunc(sym_a), V1976, V1975)

			if True == gen2928 {
				__e.TailApply(ShenFunc(symaddress_1_6), V1974, V1976, V1977)

				return
			} else {
				gen2926 := Call(__e, ShenFunc(symaddress_1_6), V1974, V1975, V1977)

				gen2927 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1975)

				__e.TailApply(ShenFunc(symshen_4fillvector), gen2926, gen2927, V1976, V1977)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fillvector"), gen2929)

		gen2937 := MakeNative(func(__e Evaluator) {
			V1979 := __e.Get(1)
			_ = V1979
			gen2936 := Call(__e, ShenFunc(symabsvector_2), V1979)

			if True == gen2936 {
				gen2930 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeNumber(-1))
					return
				}, 1)
				gen2931 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(sym_5_1address), V1979, MakeNumber(0))

					return
				}, 0)
				gen2932 := Try(__e, gen2931).Catch(gen2930)
				X := gen2932
				gen2934 := Call(__e, ShenFunc(symnumber_2), X)

				var gen2935 Obj
				if True == gen2934 {
					gen2933 := Call(__e, ShenFunc(sym_6_a), X, MakeNumber(0))

					if True == gen2933 {
						gen2935 = True
					} else {
						gen2935 = False
					}

				} else {
					gen2935 = False
				}
				if True == gen2935 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("vector?"), gen2937)

		gen2939 := MakeNative(func(__e Evaluator) {
			V1983 := __e.Get(1)
			_ = V1983
			V1984 := __e.Get(2)
			_ = V1984
			V1985 := __e.Get(3)
			_ = V1985
			gen2938 := Call(__e, ShenFunc(sym_a), V1984, MakeNumber(0))

			if True == gen2938 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("cannot access 0th element of a vector\n"))

				return
			} else {
				__e.TailApply(ShenFunc(symaddress_1_6), V1983, V1984, V1985)

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("vector->"), gen2939)

		gen2944 := MakeNative(func(__e Evaluator) {
			V1988 := __e.Get(1)
			_ = V1988
			V1989 := __e.Get(2)
			_ = V1989
			gen2943 := Call(__e, ShenFunc(sym_a), V1989, MakeNumber(0))

			if True == gen2943 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("cannot access 0th element of a vector\n"))

				return
			} else {
				gen2940 := Call(__e, ShenFunc(sym_5_1address), V1988, V1989)

				VectorElement := gen2940
				gen2941 := Call(__e, ShenFunc(symfail))

				gen2942 := Call(__e, ShenFunc(sym_a), VectorElement, gen2941)

				if True == gen2942 {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("vector element not found\n"))

					return
				} else {
					__e.Return(VectorElement)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("<-vector"), gen2944)

		gen2947 := MakeNative(func(__e Evaluator) {
			V1991 := __e.Get(1)
			_ = V1991
			gen2946 := Call(__e, ShenFunc(syminteger_2), V1991)

			if True == gen2946 {
				gen2945 := Call(__e, ShenFunc(sym_6_a), V1991, MakeNumber(0))

				if True == gen2945 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.posint?"), gen2947)

		gen2948 := MakeNative(func(__e Evaluator) {
			V1993 := __e.Get(1)
			_ = V1993
			__e.TailApply(ShenFunc(sym_5_1address), V1993, MakeNumber(0))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("limit"), gen2948)

		gen2957 := MakeNative(func(__e Evaluator) {
			V1995 := __e.Get(1)
			_ = V1995
			gen2955 := Call(__e, ShenFunc(symboolean_2), V1995)

			var gen2956 Obj
			if True == gen2955 {
				gen2956 = True
			} else {
				gen2953 := Call(__e, ShenFunc(symnumber_2), V1995)

				var gen2954 Obj
				if True == gen2953 {
					gen2954 = True
				} else {
					gen2952 := Call(__e, ShenFunc(symstring_2), V1995)

					if True == gen2952 {
						gen2954 = True
					} else {
						gen2954 = False
					}

				}
				if True == gen2954 {
					gen2956 = True
				} else {
					gen2956 = False
				}

			}
			if True == gen2956 {
				__e.Return(False)
				return
			} else {
				gen2949 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(False)
					return
				}, 1)
				gen2951 := MakeNative(func(__e Evaluator) {
					gen2950 := Call(__e, ShenFunc(symstr), V1995)

					String := gen2950
					__e.TailApply(ShenFunc(symshen_4analyse_1symbol_2), String)

					return

				}, 0)
				__e.Return(Try(__e, gen2951).Catch(gen2949))
				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("symbol?"), gen2957)

		gen2964 := MakeNative(func(__e Evaluator) {
			V1997 := __e.Get(1)
			_ = V1997
			gen2963 := Call(__e, ShenFunc(sym_a), MakeString(""), V1997)

			if True == gen2963 {
				__e.Return(False)
				return
			} else {
				gen2962 := Call(__e, ShenFunc(symshen_4_7string_2), V1997)

				if True == gen2962 {
					gen2960 := Call(__e, ShenFunc(sympos), V1997, MakeNumber(0))

					gen2961 := Call(__e, ShenFunc(symshen_4alpha_2), gen2960)

					if True == gen2961 {
						gen2958 := Call(__e, ShenFunc(symtlstr), V1997)

						gen2959 := Call(__e, ShenFunc(symshen_4alphanums_2), gen2958)

						if True == gen2959 {
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

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.analyse-symbol?"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.analyse-symbol?"), gen2964)

		gen3040 := MakeNative(func(__e Evaluator) {
			V1999 := __e.Get(1)
			_ = V1999
			gen2965 := Call(__e, ShenFunc(symcons), MakeString("."), Nil)

			gen2966 := Call(__e, ShenFunc(symcons), MakeString("'"), gen2965)

			gen2967 := Call(__e, ShenFunc(symcons), MakeString("#"), gen2966)

			gen2968 := Call(__e, ShenFunc(symcons), MakeString("`"), gen2967)

			gen2969 := Call(__e, ShenFunc(symcons), MakeString(";"), gen2968)

			gen2970 := Call(__e, ShenFunc(symcons), MakeString(":"), gen2969)

			gen2971 := Call(__e, ShenFunc(symcons), MakeString("}"), gen2970)

			gen2972 := Call(__e, ShenFunc(symcons), MakeString("{"), gen2971)

			gen2973 := Call(__e, ShenFunc(symcons), MakeString("%"), gen2972)

			gen2974 := Call(__e, ShenFunc(symcons), MakeString("&"), gen2973)

			gen2975 := Call(__e, ShenFunc(symcons), MakeString("<"), gen2974)

			gen2976 := Call(__e, ShenFunc(symcons), MakeString(">"), gen2975)

			gen2977 := Call(__e, ShenFunc(symcons), MakeString("~"), gen2976)

			gen2978 := Call(__e, ShenFunc(symcons), MakeString("@"), gen2977)

			gen2979 := Call(__e, ShenFunc(symcons), MakeString("!"), gen2978)

			gen2980 := Call(__e, ShenFunc(symcons), MakeString("$"), gen2979)

			gen2981 := Call(__e, ShenFunc(symcons), MakeString("?"), gen2980)

			gen2982 := Call(__e, ShenFunc(symcons), MakeString("_"), gen2981)

			gen2983 := Call(__e, ShenFunc(symcons), MakeString("-"), gen2982)

			gen2984 := Call(__e, ShenFunc(symcons), MakeString("+"), gen2983)

			gen2985 := Call(__e, ShenFunc(symcons), MakeString("/"), gen2984)

			gen2986 := Call(__e, ShenFunc(symcons), MakeString("*"), gen2985)

			gen2987 := Call(__e, ShenFunc(symcons), MakeString("="), gen2986)

			gen2988 := Call(__e, ShenFunc(symcons), MakeString("z"), gen2987)

			gen2989 := Call(__e, ShenFunc(symcons), MakeString("y"), gen2988)

			gen2990 := Call(__e, ShenFunc(symcons), MakeString("x"), gen2989)

			gen2991 := Call(__e, ShenFunc(symcons), MakeString("w"), gen2990)

			gen2992 := Call(__e, ShenFunc(symcons), MakeString("v"), gen2991)

			gen2993 := Call(__e, ShenFunc(symcons), MakeString("u"), gen2992)

			gen2994 := Call(__e, ShenFunc(symcons), MakeString("t"), gen2993)

			gen2995 := Call(__e, ShenFunc(symcons), MakeString("s"), gen2994)

			gen2996 := Call(__e, ShenFunc(symcons), MakeString("r"), gen2995)

			gen2997 := Call(__e, ShenFunc(symcons), MakeString("q"), gen2996)

			gen2998 := Call(__e, ShenFunc(symcons), MakeString("p"), gen2997)

			gen2999 := Call(__e, ShenFunc(symcons), MakeString("o"), gen2998)

			gen3000 := Call(__e, ShenFunc(symcons), MakeString("n"), gen2999)

			gen3001 := Call(__e, ShenFunc(symcons), MakeString("m"), gen3000)

			gen3002 := Call(__e, ShenFunc(symcons), MakeString("l"), gen3001)

			gen3003 := Call(__e, ShenFunc(symcons), MakeString("k"), gen3002)

			gen3004 := Call(__e, ShenFunc(symcons), MakeString("j"), gen3003)

			gen3005 := Call(__e, ShenFunc(symcons), MakeString("i"), gen3004)

			gen3006 := Call(__e, ShenFunc(symcons), MakeString("h"), gen3005)

			gen3007 := Call(__e, ShenFunc(symcons), MakeString("g"), gen3006)

			gen3008 := Call(__e, ShenFunc(symcons), MakeString("f"), gen3007)

			gen3009 := Call(__e, ShenFunc(symcons), MakeString("e"), gen3008)

			gen3010 := Call(__e, ShenFunc(symcons), MakeString("d"), gen3009)

			gen3011 := Call(__e, ShenFunc(symcons), MakeString("c"), gen3010)

			gen3012 := Call(__e, ShenFunc(symcons), MakeString("b"), gen3011)

			gen3013 := Call(__e, ShenFunc(symcons), MakeString("a"), gen3012)

			gen3014 := Call(__e, ShenFunc(symcons), MakeString("Z"), gen3013)

			gen3015 := Call(__e, ShenFunc(symcons), MakeString("Y"), gen3014)

			gen3016 := Call(__e, ShenFunc(symcons), MakeString("X"), gen3015)

			gen3017 := Call(__e, ShenFunc(symcons), MakeString("W"), gen3016)

			gen3018 := Call(__e, ShenFunc(symcons), MakeString("V"), gen3017)

			gen3019 := Call(__e, ShenFunc(symcons), MakeString("U"), gen3018)

			gen3020 := Call(__e, ShenFunc(symcons), MakeString("T"), gen3019)

			gen3021 := Call(__e, ShenFunc(symcons), MakeString("S"), gen3020)

			gen3022 := Call(__e, ShenFunc(symcons), MakeString("R"), gen3021)

			gen3023 := Call(__e, ShenFunc(symcons), MakeString("Q"), gen3022)

			gen3024 := Call(__e, ShenFunc(symcons), MakeString("P"), gen3023)

			gen3025 := Call(__e, ShenFunc(symcons), MakeString("O"), gen3024)

			gen3026 := Call(__e, ShenFunc(symcons), MakeString("N"), gen3025)

			gen3027 := Call(__e, ShenFunc(symcons), MakeString("M"), gen3026)

			gen3028 := Call(__e, ShenFunc(symcons), MakeString("L"), gen3027)

			gen3029 := Call(__e, ShenFunc(symcons), MakeString("K"), gen3028)

			gen3030 := Call(__e, ShenFunc(symcons), MakeString("J"), gen3029)

			gen3031 := Call(__e, ShenFunc(symcons), MakeString("I"), gen3030)

			gen3032 := Call(__e, ShenFunc(symcons), MakeString("H"), gen3031)

			gen3033 := Call(__e, ShenFunc(symcons), MakeString("G"), gen3032)

			gen3034 := Call(__e, ShenFunc(symcons), MakeString("F"), gen3033)

			gen3035 := Call(__e, ShenFunc(symcons), MakeString("E"), gen3034)

			gen3036 := Call(__e, ShenFunc(symcons), MakeString("D"), gen3035)

			gen3037 := Call(__e, ShenFunc(symcons), MakeString("C"), gen3036)

			gen3038 := Call(__e, ShenFunc(symcons), MakeString("B"), gen3037)

			gen3039 := Call(__e, ShenFunc(symcons), MakeString("A"), gen3038)

			__e.TailApply(ShenFunc(symelement_2), V1999, gen3039)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.alpha?"), gen3040)

		gen3047 := MakeNative(func(__e Evaluator) {
			V2001 := __e.Get(1)
			_ = V2001
			gen3046 := Call(__e, ShenFunc(sym_a), MakeString(""), V2001)

			if True == gen3046 {
				__e.Return(True)
				return
			} else {
				gen3045 := Call(__e, ShenFunc(symshen_4_7string_2), V2001)

				if True == gen3045 {
					gen3043 := Call(__e, ShenFunc(sympos), V2001, MakeNumber(0))

					gen3044 := Call(__e, ShenFunc(symshen_4alphanum_2), gen3043)

					if True == gen3044 {
						gen3041 := Call(__e, ShenFunc(symtlstr), V2001)

						gen3042 := Call(__e, ShenFunc(symshen_4alphanums_2), gen3041)

						if True == gen3042 {
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

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.alphanums?"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.alphanums?"), gen3047)

		gen3050 := MakeNative(func(__e Evaluator) {
			V2003 := __e.Get(1)
			_ = V2003
			gen3049 := Call(__e, ShenFunc(symshen_4alpha_2), V2003)

			if True == gen3049 {
				__e.Return(True)
				return
			} else {
				gen3048 := Call(__e, ShenFunc(symshen_4digit_2), V2003)

				if True == gen3048 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.alphanum?"), gen3050)

		gen3061 := MakeNative(func(__e Evaluator) {
			V2005 := __e.Get(1)
			_ = V2005
			gen3051 := Call(__e, ShenFunc(symcons), MakeString("0"), Nil)

			gen3052 := Call(__e, ShenFunc(symcons), MakeString("9"), gen3051)

			gen3053 := Call(__e, ShenFunc(symcons), MakeString("8"), gen3052)

			gen3054 := Call(__e, ShenFunc(symcons), MakeString("7"), gen3053)

			gen3055 := Call(__e, ShenFunc(symcons), MakeString("6"), gen3054)

			gen3056 := Call(__e, ShenFunc(symcons), MakeString("5"), gen3055)

			gen3057 := Call(__e, ShenFunc(symcons), MakeString("4"), gen3056)

			gen3058 := Call(__e, ShenFunc(symcons), MakeString("3"), gen3057)

			gen3059 := Call(__e, ShenFunc(symcons), MakeString("2"), gen3058)

			gen3060 := Call(__e, ShenFunc(symcons), MakeString("1"), gen3059)

			__e.TailApply(ShenFunc(symelement_2), V2005, gen3060)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.digit?"), gen3061)

		gen3070 := MakeNative(func(__e Evaluator) {
			V2007 := __e.Get(1)
			_ = V2007
			gen3068 := Call(__e, ShenFunc(symboolean_2), V2007)

			var gen3069 Obj
			if True == gen3068 {
				gen3069 = True
			} else {
				gen3066 := Call(__e, ShenFunc(symnumber_2), V2007)

				var gen3067 Obj
				if True == gen3066 {
					gen3067 = True
				} else {
					gen3065 := Call(__e, ShenFunc(symstring_2), V2007)

					if True == gen3065 {
						gen3067 = True
					} else {
						gen3067 = False
					}

				}
				if True == gen3067 {
					gen3069 = True
				} else {
					gen3069 = False
				}

			}
			if True == gen3069 {
				__e.Return(False)
				return
			} else {
				gen3062 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(False)
					return
				}, 1)
				gen3064 := MakeNative(func(__e Evaluator) {
					gen3063 := Call(__e, ShenFunc(symstr), V2007)

					String := gen3063
					__e.TailApply(ShenFunc(symshen_4analyse_1variable_2), String)

					return

				}, 0)
				__e.Return(Try(__e, gen3064).Catch(gen3062))
				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("variable?"), gen3070)

		gen3076 := MakeNative(func(__e Evaluator) {
			V2009 := __e.Get(1)
			_ = V2009
			gen3075 := Call(__e, ShenFunc(symshen_4_7string_2), V2009)

			if True == gen3075 {
				gen3073 := Call(__e, ShenFunc(sympos), V2009, MakeNumber(0))

				gen3074 := Call(__e, ShenFunc(symshen_4uppercase_2), gen3073)

				if True == gen3074 {
					gen3071 := Call(__e, ShenFunc(symtlstr), V2009)

					gen3072 := Call(__e, ShenFunc(symshen_4alphanums_2), gen3071)

					if True == gen3072 {
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

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.analyse-variable?"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.analyse-variable?"), gen3076)

		gen3103 := MakeNative(func(__e Evaluator) {
			V2011 := __e.Get(1)
			_ = V2011
			gen3077 := Call(__e, ShenFunc(symcons), MakeString("Z"), Nil)

			gen3078 := Call(__e, ShenFunc(symcons), MakeString("Y"), gen3077)

			gen3079 := Call(__e, ShenFunc(symcons), MakeString("X"), gen3078)

			gen3080 := Call(__e, ShenFunc(symcons), MakeString("W"), gen3079)

			gen3081 := Call(__e, ShenFunc(symcons), MakeString("V"), gen3080)

			gen3082 := Call(__e, ShenFunc(symcons), MakeString("U"), gen3081)

			gen3083 := Call(__e, ShenFunc(symcons), MakeString("T"), gen3082)

			gen3084 := Call(__e, ShenFunc(symcons), MakeString("S"), gen3083)

			gen3085 := Call(__e, ShenFunc(symcons), MakeString("R"), gen3084)

			gen3086 := Call(__e, ShenFunc(symcons), MakeString("Q"), gen3085)

			gen3087 := Call(__e, ShenFunc(symcons), MakeString("P"), gen3086)

			gen3088 := Call(__e, ShenFunc(symcons), MakeString("O"), gen3087)

			gen3089 := Call(__e, ShenFunc(symcons), MakeString("N"), gen3088)

			gen3090 := Call(__e, ShenFunc(symcons), MakeString("M"), gen3089)

			gen3091 := Call(__e, ShenFunc(symcons), MakeString("L"), gen3090)

			gen3092 := Call(__e, ShenFunc(symcons), MakeString("K"), gen3091)

			gen3093 := Call(__e, ShenFunc(symcons), MakeString("J"), gen3092)

			gen3094 := Call(__e, ShenFunc(symcons), MakeString("I"), gen3093)

			gen3095 := Call(__e, ShenFunc(symcons), MakeString("H"), gen3094)

			gen3096 := Call(__e, ShenFunc(symcons), MakeString("G"), gen3095)

			gen3097 := Call(__e, ShenFunc(symcons), MakeString("F"), gen3096)

			gen3098 := Call(__e, ShenFunc(symcons), MakeString("E"), gen3097)

			gen3099 := Call(__e, ShenFunc(symcons), MakeString("D"), gen3098)

			gen3100 := Call(__e, ShenFunc(symcons), MakeString("C"), gen3099)

			gen3101 := Call(__e, ShenFunc(symcons), MakeString("B"), gen3100)

			gen3102 := Call(__e, ShenFunc(symcons), MakeString("A"), gen3101)

			__e.TailApply(ShenFunc(symelement_2), V2011, gen3102)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.uppercase?"), gen3103)

		gen3107 := MakeNative(func(__e Evaluator) {
			V2013 := __e.Get(1)
			_ = V2013
			gen3104 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*gensym*"))

			gen3105 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen3104)

			gen3106 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*gensym*"), gen3105)

			__e.TailApply(ShenFunc(symconcat), V2013, gen3106)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("gensym"), gen3107)

		gen3111 := MakeNative(func(__e Evaluator) {
			V2016 := __e.Get(1)
			_ = V2016
			V2017 := __e.Get(2)
			_ = V2017
			gen3108 := Call(__e, ShenFunc(symstr), V2016)

			gen3109 := Call(__e, ShenFunc(symstr), V2017)

			gen3110 := Call(__e, ShenFunc(symcn), gen3108, gen3109)

			__e.TailApply(ShenFunc(symintern), gen3110)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("concat"), gen3111)

		gen3116 := MakeNative(func(__e Evaluator) {
			V2020 := __e.Get(1)
			_ = V2020
			V2021 := __e.Get(2)
			_ = V2021
			gen3112 := Call(__e, ShenFunc(symabsvector), MakeNumber(3))

			Vector := gen3112
			gen3113 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(0), MakeSymbol("shen.tuple"))

			Tag := gen3113
			_ = Tag
			gen3114 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(1), V2020)

			Fst := gen3114
			_ = Fst
			gen3115 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(2), V2021)

			Snd := gen3115
			_ = Snd
			__e.Return(Vector)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("@p"), gen3116)

		gen3117 := MakeNative(func(__e Evaluator) {
			V2023 := __e.Get(1)
			_ = V2023
			__e.TailApply(ShenFunc(sym_5_1address), V2023, MakeNumber(1))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fst"), gen3117)

		gen3118 := MakeNative(func(__e Evaluator) {
			V2025 := __e.Get(1)
			_ = V2025
			__e.TailApply(ShenFunc(sym_5_1address), V2025, MakeNumber(2))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("snd"), gen3118)

		gen3124 := MakeNative(func(__e Evaluator) {
			V2027 := __e.Get(1)
			_ = V2027
			gen3123 := Call(__e, ShenFunc(symabsvector_2), V2027)

			if True == gen3123 {
				gen3119 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeSymbol("shen.not-tuple"))
					return
				}, 1)
				gen3120 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(sym_5_1address), V2027, MakeNumber(0))

					return
				}, 0)
				gen3121 := Try(__e, gen3120).Catch(gen3119)
				gen3122 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tuple"), gen3121)

				if True == gen3122 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("tuple?"), gen3124)

		gen3130 := MakeNative(func(__e Evaluator) {
			V2030 := __e.Get(1)
			_ = V2030
			V2031 := __e.Get(2)
			_ = V2031
			gen3129 := Call(__e, ShenFunc(sym_a), Nil, V2030)

			if True == gen3129 {
				__e.Return(V2031)
				return
			} else {
				gen3128 := Call(__e, ShenFunc(symcons_2), V2030)

				if True == gen3128 {
					gen3125 := Call(__e, ShenFunc(symhd), V2030)

					gen3126 := Call(__e, ShenFunc(symtl), V2030)

					gen3127 := Call(__e, ShenFunc(symappend), gen3126, V2031)

					__e.TailApply(ShenFunc(symcons), gen3125, gen3127)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("append"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("append"), gen3130)

		gen3136 := MakeNative(func(__e Evaluator) {
			V2034 := __e.Get(1)
			_ = V2034
			V2035 := __e.Get(2)
			_ = V2035
			gen3131 := Call(__e, ShenFunc(symlimit), V2035)

			Limit := gen3131
			gen3132 := Call(__e, ShenFunc(sym_7), Limit, MakeNumber(1))

			gen3133 := Call(__e, ShenFunc(symvector), gen3132)

			NewVector := gen3133
			gen3134 := Call(__e, ShenFunc(symvector_1_6), NewVector, MakeNumber(1), V2034)

			X_7NewVector := gen3134
			gen3135 := Call(__e, ShenFunc(sym_a), Limit, MakeNumber(0))

			if True == gen3135 {
				__e.Return(X_7NewVector)
				return
			} else {
				__e.TailApply(ShenFunc(symshen_4_8v_1help), V2035, MakeNumber(1), Limit, X_7NewVector)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("@v"), gen3136)

		gen3142 := MakeNative(func(__e Evaluator) {
			V2041 := __e.Get(1)
			_ = V2041
			V2042 := __e.Get(2)
			_ = V2042
			V2043 := __e.Get(3)
			_ = V2043
			V2044 := __e.Get(4)
			_ = V2044
			gen3141 := Call(__e, ShenFunc(sym_a), V2043, V2042)

			if True == gen3141 {
				gen3140 := Call(__e, ShenFunc(sym_7), V2043, MakeNumber(1))

				__e.TailApply(ShenFunc(symshen_4copyfromvector), V2041, V2044, V2043, gen3140)

				return

			} else {
				gen3137 := Call(__e, ShenFunc(sym_7), V2042, MakeNumber(1))

				gen3138 := Call(__e, ShenFunc(sym_7), V2042, MakeNumber(1))

				gen3139 := Call(__e, ShenFunc(symshen_4copyfromvector), V2041, V2044, V2042, gen3138)

				__e.TailApply(ShenFunc(symshen_4_8v_1help), V2041, gen3137, V2043, gen3139)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.@v-help"), gen3142)

		gen3146 := MakeNative(func(__e Evaluator) {
			V2049 := __e.Get(1)
			_ = V2049
			V2050 := __e.Get(2)
			_ = V2050
			V2051 := __e.Get(3)
			_ = V2051
			V2052 := __e.Get(4)
			_ = V2052
			gen3143 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(V2050)
				return
			}, 1)
			gen3145 := MakeNative(func(__e Evaluator) {
				gen3144 := Call(__e, ShenFunc(sym_5_1vector), V2049, V2051)

				__e.TailApply(ShenFunc(symvector_1_6), V2050, V2052, gen3144)

				return

			}, 0)
			__e.Return(Try(__e, gen3145).Catch(gen3143))
			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copyfromvector"), gen3146)

		gen3151 := MakeNative(func(__e Evaluator) {
			V2054 := __e.Get(1)
			_ = V2054
			gen3149 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				gen3147 := Call(__e, ShenFunc(symshen_4app), V2054, MakeString("\n"), MakeSymbol("shen.s"))

				gen3148 := Call(__e, ShenFunc(symcn), MakeString("hdv needs a non-empty vector as an argument; not "), gen3147)

				__e.TailApply(ShenFunc(symsimple_1error), gen3148)

				return

			}, 1)
			gen3150 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(sym_5_1vector), V2054, MakeNumber(1))

				return
			}, 0)
			__e.Return(Try(__e, gen3150).Catch(gen3149))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("hdv"), gen3151)

		gen3159 := MakeNative(func(__e Evaluator) {
			V2056 := __e.Get(1)
			_ = V2056
			gen3152 := Call(__e, ShenFunc(symlimit), V2056)

			Limit := gen3152
			gen3158 := Call(__e, ShenFunc(sym_a), Limit, MakeNumber(0))

			if True == gen3158 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("cannot take the tail of the empty vector\n"))

				return
			} else {
				gen3157 := Call(__e, ShenFunc(sym_a), Limit, MakeNumber(1))

				if True == gen3157 {
					__e.TailApply(ShenFunc(symvector), MakeNumber(0))

					return
				} else {
					gen3153 := Call(__e, ShenFunc(sym_1), Limit, MakeNumber(1))

					gen3154 := Call(__e, ShenFunc(symvector), gen3153)

					NewVector := gen3154
					_ = NewVector
					gen3155 := Call(__e, ShenFunc(sym_1), Limit, MakeNumber(1))

					gen3156 := Call(__e, ShenFunc(symvector), gen3155)

					__e.TailApply(ShenFunc(symshen_4tlv_1help), V2056, MakeNumber(2), Limit, gen3156)

					return

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tlv"), gen3159)

		gen3165 := MakeNative(func(__e Evaluator) {
			V2062 := __e.Get(1)
			_ = V2062
			V2063 := __e.Get(2)
			_ = V2063
			V2064 := __e.Get(3)
			_ = V2064
			V2065 := __e.Get(4)
			_ = V2065
			gen3164 := Call(__e, ShenFunc(sym_a), V2064, V2063)

			if True == gen3164 {
				gen3163 := Call(__e, ShenFunc(sym_1), V2064, MakeNumber(1))

				__e.TailApply(ShenFunc(symshen_4copyfromvector), V2062, V2065, V2064, gen3163)

				return

			} else {
				gen3160 := Call(__e, ShenFunc(sym_7), V2063, MakeNumber(1))

				gen3161 := Call(__e, ShenFunc(sym_1), V2063, MakeNumber(1))

				gen3162 := Call(__e, ShenFunc(symshen_4copyfromvector), V2062, V2065, V2063, gen3161)

				__e.TailApply(ShenFunc(symshen_4tlv_1help), V2062, gen3160, V2064, gen3162)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tlv-help"), gen3165)

		gen3177 := MakeNative(func(__e Evaluator) {
			V2077 := __e.Get(1)
			_ = V2077
			V2078 := __e.Get(2)
			_ = V2078
			gen3176 := Call(__e, ShenFunc(sym_a), Nil, V2078)

			if True == gen3176 {
				__e.Return(Nil)
				return
			} else {
				gen3174 := Call(__e, ShenFunc(symcons_2), V2078)

				var gen3175 Obj
				if True == gen3174 {
					gen3171 := Call(__e, ShenFunc(symhd), V2078)

					gen3172 := Call(__e, ShenFunc(symcons_2), gen3171)

					var gen3173 Obj
					if True == gen3172 {
						gen3168 := Call(__e, ShenFunc(symhd), V2078)

						gen3169 := Call(__e, ShenFunc(symhd), gen3168)

						gen3170 := Call(__e, ShenFunc(sym_a), gen3169, V2077)

						if True == gen3170 {
							gen3173 = True
						} else {
							gen3173 = False
						}

					} else {
						gen3173 = False
					}
					if True == gen3173 {
						gen3175 = True
					} else {
						gen3175 = False
					}

				} else {
					gen3175 = False
				}
				if True == gen3175 {
					__e.TailApply(ShenFunc(symhd), V2078)

					return
				} else {
					gen3167 := Call(__e, ShenFunc(symcons_2), V2078)

					if True == gen3167 {
						gen3166 := Call(__e, ShenFunc(symtl), V2078)

						__e.TailApply(ShenFunc(symassoc), V2077, gen3166)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("assoc"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("assoc"), gen3177)

		gen3196 := MakeNative(func(__e Evaluator) {
			V2085 := __e.Get(1)
			_ = V2085
			V2086 := __e.Get(2)
			_ = V2086
			V2087 := __e.Get(3)
			_ = V2087
			gen3195 := Call(__e, ShenFunc(sym_a), Nil, V2087)

			if True == gen3195 {
				gen3194 := Call(__e, ShenFunc(symcons), V2085, V2086)

				__e.TailApply(ShenFunc(symcons), gen3194, Nil)

				return

			} else {
				gen3192 := Call(__e, ShenFunc(symcons_2), V2087)

				var gen3193 Obj
				if True == gen3192 {
					gen3189 := Call(__e, ShenFunc(symhd), V2087)

					gen3190 := Call(__e, ShenFunc(symcons_2), gen3189)

					var gen3191 Obj
					if True == gen3190 {
						gen3186 := Call(__e, ShenFunc(symhd), V2087)

						gen3187 := Call(__e, ShenFunc(symhd), gen3186)

						gen3188 := Call(__e, ShenFunc(sym_a), gen3187, V2085)

						if True == gen3188 {
							gen3191 = True
						} else {
							gen3191 = False
						}

					} else {
						gen3191 = False
					}
					if True == gen3191 {
						gen3193 = True
					} else {
						gen3193 = False
					}

				} else {
					gen3193 = False
				}
				if True == gen3193 {
					gen3182 := Call(__e, ShenFunc(symhd), V2087)

					gen3183 := Call(__e, ShenFunc(symhd), gen3182)

					gen3184 := Call(__e, ShenFunc(symcons), gen3183, V2086)

					gen3185 := Call(__e, ShenFunc(symtl), V2087)

					__e.TailApply(ShenFunc(symcons), gen3184, gen3185)

					return

				} else {
					gen3181 := Call(__e, ShenFunc(symcons_2), V2087)

					if True == gen3181 {
						gen3178 := Call(__e, ShenFunc(symhd), V2087)

						gen3179 := Call(__e, ShenFunc(symtl), V2087)

						gen3180 := Call(__e, ShenFunc(symshen_4assoc_1set), V2085, V2086, gen3179)

						__e.TailApply(ShenFunc(symcons), gen3178, gen3180)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.assoc-set"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assoc-set"), gen3196)

		gen3210 := MakeNative(func(__e Evaluator) {
			V2093 := __e.Get(1)
			_ = V2093
			V2094 := __e.Get(2)
			_ = V2094
			gen3209 := Call(__e, ShenFunc(sym_a), Nil, V2094)

			if True == gen3209 {
				__e.Return(Nil)
				return
			} else {
				gen3207 := Call(__e, ShenFunc(symcons_2), V2094)

				var gen3208 Obj
				if True == gen3207 {
					gen3204 := Call(__e, ShenFunc(symhd), V2094)

					gen3205 := Call(__e, ShenFunc(symcons_2), gen3204)

					var gen3206 Obj
					if True == gen3205 {
						gen3201 := Call(__e, ShenFunc(symhd), V2094)

						gen3202 := Call(__e, ShenFunc(symhd), gen3201)

						gen3203 := Call(__e, ShenFunc(sym_a), gen3202, V2093)

						if True == gen3203 {
							gen3206 = True
						} else {
							gen3206 = False
						}

					} else {
						gen3206 = False
					}
					if True == gen3206 {
						gen3208 = True
					} else {
						gen3208 = False
					}

				} else {
					gen3208 = False
				}
				if True == gen3208 {
					__e.TailApply(ShenFunc(symtl), V2094)

					return
				} else {
					gen3200 := Call(__e, ShenFunc(symcons_2), V2094)

					if True == gen3200 {
						gen3197 := Call(__e, ShenFunc(symhd), V2094)

						gen3198 := Call(__e, ShenFunc(symtl), V2094)

						gen3199 := Call(__e, ShenFunc(symshen_4assoc_1rm), V2093, gen3198)

						__e.TailApply(ShenFunc(symcons), gen3197, gen3199)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.assoc-rm"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assoc-rm"), gen3210)

		gen3213 := MakeNative(func(__e Evaluator) {
			V2100 := __e.Get(1)
			_ = V2100
			gen3212 := Call(__e, ShenFunc(sym_a), True, V2100)

			if True == gen3212 {
				__e.Return(True)
				return
			} else {
				gen3211 := Call(__e, ShenFunc(sym_a), False, V2100)

				if True == gen3211 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("boolean?"), gen3213)

		gen3217 := MakeNative(func(__e Evaluator) {
			V2102 := __e.Get(1)
			_ = V2102
			gen3216 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V2102)

			if True == gen3216 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen3214 := Call(__e, ShenFunc(symstoutput))

				Call(__e, ShenFunc(symshen_4prhush), MakeString("\n"), gen3214)

				gen3215 := Call(__e, ShenFunc(sym_1), V2102, MakeNumber(1))

				__e.TailApply(ShenFunc(symnl), gen3215)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("nl"), gen3217)

		gen3226 := MakeNative(func(__e Evaluator) {
			V2107 := __e.Get(1)
			_ = V2107
			V2108 := __e.Get(2)
			_ = V2108
			gen3225 := Call(__e, ShenFunc(sym_a), Nil, V2107)

			if True == gen3225 {
				__e.Return(Nil)
				return
			} else {
				gen3224 := Call(__e, ShenFunc(symcons_2), V2107)

				if True == gen3224 {
					gen3222 := Call(__e, ShenFunc(symhd), V2107)

					gen3223 := Call(__e, ShenFunc(symelement_2), gen3222, V2108)

					if True == gen3223 {
						gen3221 := Call(__e, ShenFunc(symtl), V2107)

						__e.TailApply(ShenFunc(symdifference), gen3221, V2108)

						return

					} else {
						gen3218 := Call(__e, ShenFunc(symhd), V2107)

						gen3219 := Call(__e, ShenFunc(symtl), V2107)

						gen3220 := Call(__e, ShenFunc(symdifference), gen3219, V2108)

						__e.TailApply(ShenFunc(symcons), gen3218, gen3220)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("difference"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("difference"), gen3226)

		gen3227 := MakeNative(func(__e Evaluator) {
			V2111 := __e.Get(1)
			_ = V2111
			V2112 := __e.Get(2)
			_ = V2112
			__e.Return(V2112)
			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("do"), gen3227)

		gen3235 := MakeNative(func(__e Evaluator) {
			V2124 := __e.Get(1)
			_ = V2124
			V2125 := __e.Get(2)
			_ = V2125
			gen3234 := Call(__e, ShenFunc(sym_a), Nil, V2125)

			if True == gen3234 {
				__e.Return(False)
				return
			} else {
				gen3232 := Call(__e, ShenFunc(symcons_2), V2125)

				var gen3233 Obj
				if True == gen3232 {
					gen3230 := Call(__e, ShenFunc(symhd), V2125)

					gen3231 := Call(__e, ShenFunc(sym_a), gen3230, V2124)

					if True == gen3231 {
						gen3233 = True
					} else {
						gen3233 = False
					}

				} else {
					gen3233 = False
				}
				if True == gen3233 {
					__e.Return(True)
					return
				} else {
					gen3229 := Call(__e, ShenFunc(symcons_2), V2125)

					if True == gen3229 {
						gen3228 := Call(__e, ShenFunc(symtl), V2125)

						__e.TailApply(ShenFunc(symelement_2), V2124, gen3228)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("element?"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("element?"), gen3235)

		gen3237 := MakeNative(func(__e Evaluator) {
			V2131 := __e.Get(1)
			_ = V2131
			gen3236 := Call(__e, ShenFunc(sym_a), Nil, V2131)

			if True == gen3236 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("empty?"), gen3237)

		gen3239 := MakeNative(func(__e Evaluator) {
			V2134 := __e.Get(1)
			_ = V2134
			V2135 := __e.Get(2)
			_ = V2135
			gen3238 := Call(__e, V2134, V2135)

			__e.TailApply(ShenFunc(symshen_4fix_1help), V2134, V2135, gen3238)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fix"), gen3239)

		gen3242 := MakeNative(func(__e Evaluator) {
			V2146 := __e.Get(1)
			_ = V2146
			V2147 := __e.Get(2)
			_ = V2147
			V2148 := __e.Get(3)
			_ = V2148
			gen3241 := Call(__e, ShenFunc(sym_a), V2148, V2147)

			if True == gen3241 {
				__e.Return(V2148)
				return
			} else {
				gen3240 := Call(__e, V2146, V2148)

				__e.TailApply(ShenFunc(symshen_4fix_1help), V2146, V2148, gen3240)

				return

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fix-help"), gen3242)

		gen3248 := MakeNative(func(__e Evaluator) {
			V2153 := __e.Get(1)
			_ = V2153
			V2154 := __e.Get(2)
			_ = V2154
			V2155 := __e.Get(3)
			_ = V2155
			V2156 := __e.Get(4)
			_ = V2156
			gen3243 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen3244 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(symshen_4_5_1dict), V2156, V2153)

				return
			}, 0)
			gen3245 := Try(__e, gen3244).Catch(gen3243)
			Curr := gen3245
			gen3246 := Call(__e, ShenFunc(symshen_4assoc_1set), V2154, V2155, Curr)

			Added := gen3246
			gen3247 := Call(__e, ShenFunc(symshen_4dict_1_6), V2156, V2153, Added)

			Update := gen3247
			_ = Update
			__e.Return(V2155)
			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("put"), gen3248)

		gen3254 := MakeNative(func(__e Evaluator) {
			V2160 := __e.Get(1)
			_ = V2160
			V2161 := __e.Get(2)
			_ = V2161
			V2162 := __e.Get(3)
			_ = V2162
			gen3249 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen3250 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(symshen_4_5_1dict), V2162, V2160)

				return
			}, 0)
			gen3251 := Try(__e, gen3250).Catch(gen3249)
			Curr := gen3251
			gen3252 := Call(__e, ShenFunc(symshen_4assoc_1rm), V2161, Curr)

			Removed := gen3252
			gen3253 := Call(__e, ShenFunc(symshen_4dict_1_6), V2162, V2160, Removed)

			Update := gen3253
			_ = Update
			__e.Return(V2160)
			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unput"), gen3254)

		gen3260 := MakeNative(func(__e Evaluator) {
			V2166 := __e.Get(1)
			_ = V2166
			V2167 := __e.Get(2)
			_ = V2167
			V2168 := __e.Get(3)
			_ = V2168
			gen3255 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen3256 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(symshen_4_5_1dict), V2168, V2166)

				return
			}, 0)
			gen3257 := Try(__e, gen3256).Catch(gen3255)
			Entry := gen3257
			gen3258 := Call(__e, ShenFunc(symassoc), V2167, Entry)

			Result := gen3258
			gen3259 := Call(__e, ShenFunc(symempty_2), Result)

			if True == gen3259 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("value not found\n"))

				return
			} else {
				__e.TailApply(ShenFunc(symtl), Result)

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("get"), gen3260)

		gen3265 := MakeNative(func(__e Evaluator) {
			V2171 := __e.Get(1)
			_ = V2171
			V2172 := __e.Get(2)
			_ = V2172
			gen3261 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symstring_1_6n), X)

				return
			}, 1)
			gen3262 := Call(__e, ShenFunc(symexplode), V2171)

			gen3263 := Call(__e, ShenFunc(symmap), gen3261, gen3262)

			gen3264 := Call(__e, ShenFunc(symsum), gen3263)

			__e.TailApply(ShenFunc(symshen_4mod), gen3264, V2172)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("hash"), gen3265)

		gen3268 := MakeNative(func(__e Evaluator) {
			V2175 := __e.Get(1)
			_ = V2175
			V2176 := __e.Get(2)
			_ = V2176
			gen3266 := Call(__e, ShenFunc(symcons), V2176, Nil)

			gen3267 := Call(__e, ShenFunc(symshen_4multiples), V2175, gen3266)

			__e.TailApply(ShenFunc(symshen_4modh), V2175, gen3267)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mod"), gen3268)

		gen3277 := MakeNative(func(__e Evaluator) {
			V2179 := __e.Get(1)
			_ = V2179
			V2180 := __e.Get(2)
			_ = V2180
			gen3275 := Call(__e, ShenFunc(symcons_2), V2180)

			var gen3276 Obj
			if True == gen3275 {
				gen3273 := Call(__e, ShenFunc(symhd), V2180)

				gen3274 := Call(__e, ShenFunc(sym_6), gen3273, V2179)

				if True == gen3274 {
					gen3276 = True
				} else {
					gen3276 = False
				}

			} else {
				gen3276 = False
			}
			if True == gen3276 {
				__e.TailApply(ShenFunc(symtl), V2180)

				return
			} else {
				gen3272 := Call(__e, ShenFunc(symcons_2), V2180)

				if True == gen3272 {
					gen3269 := Call(__e, ShenFunc(symhd), V2180)

					gen3270 := Call(__e, ShenFunc(sym_d), MakeNumber(2), gen3269)

					gen3271 := Call(__e, ShenFunc(symcons), gen3270, V2180)

					__e.TailApply(ShenFunc(symshen_4multiples), V2179, gen3271)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.multiples"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.multiples"), gen3277)

		gen3290 := MakeNative(func(__e Evaluator) {
			V2185 := __e.Get(1)
			_ = V2185
			V2186 := __e.Get(2)
			_ = V2186
			gen3289 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V2185)

			if True == gen3289 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen3288 := Call(__e, ShenFunc(sym_a), Nil, V2186)

				if True == gen3288 {
					__e.Return(V2185)
					return
				} else {
					gen3286 := Call(__e, ShenFunc(symcons_2), V2186)

					var gen3287 Obj
					if True == gen3286 {
						gen3284 := Call(__e, ShenFunc(symhd), V2186)

						gen3285 := Call(__e, ShenFunc(sym_6), gen3284, V2185)

						if True == gen3285 {
							gen3287 = True
						} else {
							gen3287 = False
						}

					} else {
						gen3287 = False
					}
					if True == gen3287 {
						gen3282 := Call(__e, ShenFunc(symtl), V2186)

						gen3283 := Call(__e, ShenFunc(symempty_2), gen3282)

						if True == gen3283 {
							__e.Return(V2185)
							return
						} else {
							gen3281 := Call(__e, ShenFunc(symtl), V2186)

							__e.TailApply(ShenFunc(symshen_4modh), V2185, gen3281)

							return

						}

					} else {
						gen3280 := Call(__e, ShenFunc(symcons_2), V2186)

						if True == gen3280 {
							gen3278 := Call(__e, ShenFunc(symhd), V2186)

							gen3279 := Call(__e, ShenFunc(sym_1), V2185, gen3278)

							__e.TailApply(ShenFunc(symshen_4modh), gen3279, V2186)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.modh"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.modh"), gen3290)

		gen3296 := MakeNative(func(__e Evaluator) {
			V2188 := __e.Get(1)
			_ = V2188
			gen3295 := Call(__e, ShenFunc(sym_a), Nil, V2188)

			if True == gen3295 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen3294 := Call(__e, ShenFunc(symcons_2), V2188)

				if True == gen3294 {
					gen3291 := Call(__e, ShenFunc(symhd), V2188)

					gen3292 := Call(__e, ShenFunc(symtl), V2188)

					gen3293 := Call(__e, ShenFunc(symsum), gen3292)

					__e.TailApply(ShenFunc(sym_7), gen3291, gen3293)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("sum"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("sum"), gen3296)

		gen3298 := MakeNative(func(__e Evaluator) {
			V2196 := __e.Get(1)
			_ = V2196
			gen3297 := Call(__e, ShenFunc(symcons_2), V2196)

			if True == gen3297 {
				__e.TailApply(ShenFunc(symhd), V2196)

				return
			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("head expects a non-empty list"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("head"), gen3298)

		gen3300 := MakeNative(func(__e Evaluator) {
			V2204 := __e.Get(1)
			_ = V2204
			gen3299 := Call(__e, ShenFunc(symcons_2), V2204)

			if True == gen3299 {
				__e.TailApply(ShenFunc(symtl), V2204)

				return
			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("tail expects a non-empty list"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tail"), gen3300)

		gen3301 := MakeNative(func(__e Evaluator) {
			V2206 := __e.Get(1)
			_ = V2206
			__e.TailApply(ShenFunc(sympos), V2206, MakeNumber(0))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("hdstr"), gen3301)

		gen3310 := MakeNative(func(__e Evaluator) {
			V2211 := __e.Get(1)
			_ = V2211
			V2212 := __e.Get(2)
			_ = V2212
			gen3309 := Call(__e, ShenFunc(sym_a), Nil, V2211)

			if True == gen3309 {
				__e.Return(Nil)
				return
			} else {
				gen3308 := Call(__e, ShenFunc(symcons_2), V2211)

				if True == gen3308 {
					gen3306 := Call(__e, ShenFunc(symhd), V2211)

					gen3307 := Call(__e, ShenFunc(symelement_2), gen3306, V2212)

					if True == gen3307 {
						gen3303 := Call(__e, ShenFunc(symhd), V2211)

						gen3304 := Call(__e, ShenFunc(symtl), V2211)

						gen3305 := Call(__e, ShenFunc(symintersection), gen3304, V2212)

						__e.TailApply(ShenFunc(symcons), gen3303, gen3305)

						return

					} else {
						gen3302 := Call(__e, ShenFunc(symtl), V2211)

						__e.TailApply(ShenFunc(symintersection), gen3302, V2212)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("intersection"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("intersection"), gen3310)

		gen3311 := MakeNative(func(__e Evaluator) {
			V2214 := __e.Get(1)
			_ = V2214
			__e.TailApply(ShenFunc(symshen_4reverse__help), V2214, Nil)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("reverse"), gen3311)

		gen3317 := MakeNative(func(__e Evaluator) {
			V2217 := __e.Get(1)
			_ = V2217
			V2218 := __e.Get(2)
			_ = V2218
			gen3316 := Call(__e, ShenFunc(sym_a), Nil, V2217)

			if True == gen3316 {
				__e.Return(V2218)
				return
			} else {
				gen3315 := Call(__e, ShenFunc(symcons_2), V2217)

				if True == gen3315 {
					gen3312 := Call(__e, ShenFunc(symtl), V2217)

					gen3313 := Call(__e, ShenFunc(symhd), V2217)

					gen3314 := Call(__e, ShenFunc(symcons), gen3313, V2218)

					__e.TailApply(ShenFunc(symshen_4reverse__help), gen3312, gen3314)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.reverse_help"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.reverse_help"), gen3317)

		gen3326 := MakeNative(func(__e Evaluator) {
			V2221 := __e.Get(1)
			_ = V2221
			V2222 := __e.Get(2)
			_ = V2222
			gen3325 := Call(__e, ShenFunc(sym_a), Nil, V2221)

			if True == gen3325 {
				__e.Return(V2222)
				return
			} else {
				gen3324 := Call(__e, ShenFunc(symcons_2), V2221)

				if True == gen3324 {
					gen3322 := Call(__e, ShenFunc(symhd), V2221)

					gen3323 := Call(__e, ShenFunc(symelement_2), gen3322, V2222)

					if True == gen3323 {
						gen3321 := Call(__e, ShenFunc(symtl), V2221)

						__e.TailApply(ShenFunc(symunion), gen3321, V2222)

						return

					} else {
						gen3318 := Call(__e, ShenFunc(symhd), V2221)

						gen3319 := Call(__e, ShenFunc(symtl), V2221)

						gen3320 := Call(__e, ShenFunc(symunion), gen3319, V2222)

						__e.TailApply(ShenFunc(symcons), gen3318, gen3320)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("union"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("union"), gen3326)

		gen3338 := MakeNative(func(__e Evaluator) {
			V2224 := __e.Get(1)
			_ = V2224
			gen3327 := Call(__e, ShenFunc(symshen_4proc_1nl), V2224)

			gen3328 := Call(__e, ShenFunc(symstoutput))

			gen3329 := Call(__e, ShenFunc(symshen_4prhush), gen3327, gen3328)

			Message := gen3329
			_ = Message
			gen3330 := Call(__e, ShenFunc(symstoutput))

			gen3331 := Call(__e, ShenFunc(symshen_4prhush), MakeString(" (y/n) "), gen3330)

			Y_1or_1N := gen3331
			_ = Y_1or_1N
			gen3332 := Call(__e, ShenFunc(symstinput))

			gen3333 := Call(__e, ShenFunc(symread), gen3332)

			gen3334 := Call(__e, ShenFunc(symshen_4app), gen3333, MakeString(""), MakeSymbol("shen.s"))

			Input := gen3334
			gen3337 := Call(__e, ShenFunc(sym_a), MakeString("y"), Input)

			if True == gen3337 {
				__e.Return(True)
				return
			} else {
				gen3336 := Call(__e, ShenFunc(sym_a), MakeString("n"), Input)

				if True == gen3336 {
					__e.Return(False)
					return
				} else {
					gen3335 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), MakeString("please answer y or n\n"), gen3335)

					__e.TailApply(ShenFunc(symy_1or_1n_2), V2224)

					return

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("y-or-n?"), gen3338)

		gen3339 := MakeNative(func(__e Evaluator) {
			V2226 := __e.Get(1)
			_ = V2226
			if True == V2226 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("not"), gen3339)

		gen3343 := MakeNative(func(__e Evaluator) {
			V2239 := __e.Get(1)
			_ = V2239
			V2240 := __e.Get(2)
			_ = V2240
			V2241 := __e.Get(3)
			_ = V2241
			gen3342 := Call(__e, ShenFunc(sym_a), V2241, V2240)

			if True == gen3342 {
				__e.Return(V2239)
				return
			} else {
				gen3341 := Call(__e, ShenFunc(symcons_2), V2241)

				if True == gen3341 {
					gen3340 := MakeNative(func(__e Evaluator) {
						W := __e.Get(1)
						_ = W
						__e.TailApply(ShenFunc(symsubst), V2239, V2240, W)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen3340, V2241)

					return

				} else {
					__e.Return(V2241)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("subst"), gen3343)

		gen3345 := MakeNative(func(__e Evaluator) {
			V2243 := __e.Get(1)
			_ = V2243
			gen3344 := Call(__e, ShenFunc(symshen_4app), V2243, MakeString(""), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symshen_4explode_1h), gen3344)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("explode"), gen3345)

		gen3351 := MakeNative(func(__e Evaluator) {
			V2245 := __e.Get(1)
			_ = V2245
			gen3350 := Call(__e, ShenFunc(sym_a), MakeString(""), V2245)

			if True == gen3350 {
				__e.Return(Nil)
				return
			} else {
				gen3349 := Call(__e, ShenFunc(symshen_4_7string_2), V2245)

				if True == gen3349 {
					gen3346 := Call(__e, ShenFunc(sympos), V2245, MakeNumber(0))

					gen3347 := Call(__e, ShenFunc(symtlstr), V2245)

					gen3348 := Call(__e, ShenFunc(symshen_4explode_1h), gen3347)

					__e.TailApply(ShenFunc(symcons), gen3346, gen3348)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.explode-h"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.explode-h"), gen3351)

		gen3354 := MakeNative(func(__e Evaluator) {
			V2247 := __e.Get(1)
			_ = V2247
			gen3352 := Call(__e, ShenFunc(sym_a), V2247, MakeString(""))

			var gen3353 Obj
			if True == gen3352 {
				gen3353 = MakeString("")
			} else {
				gen3353 = Call(__e, ShenFunc(symshen_4app), V2247, MakeString("/"), MakeSymbol("shen.a"))

			}
			__e.TailApply(ShenFunc(symset), MakeSymbol("*home-directory*"), gen3353)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("cd"), gen3354)

		gen3360 := MakeNative(func(__e Evaluator) {
			V2250 := __e.Get(1)
			_ = V2250
			V2251 := __e.Get(2)
			_ = V2251
			gen3359 := Call(__e, ShenFunc(sym_a), Nil, V2251)

			if True == gen3359 {
				__e.Return(True)
				return
			} else {
				gen3358 := Call(__e, ShenFunc(symcons_2), V2251)

				if True == gen3358 {
					gen3355 := Call(__e, ShenFunc(symhd), V2251)

					gen3356 := Call(__e, V2250, gen3355)

					_ = gen3356
					gen3357 := Call(__e, ShenFunc(symtl), V2251)

					__e.TailApply(ShenFunc(symshen_4for_1each), V2250, gen3357)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.for-each"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.for-each"), gen3360)

		gen3367 := MakeNative(func(__e Evaluator) {
			V2256 := __e.Get(1)
			_ = V2256
			V2257 := __e.Get(2)
			_ = V2257
			gen3366 := Call(__e, ShenFunc(sym_a), Nil, V2257)

			if True == gen3366 {
				__e.Return(Nil)
				return
			} else {
				gen3365 := Call(__e, ShenFunc(symcons_2), V2257)

				if True == gen3365 {
					gen3361 := Call(__e, ShenFunc(symhd), V2257)

					gen3362 := Call(__e, V2256, gen3361)

					gen3363 := Call(__e, ShenFunc(symtl), V2257)

					gen3364 := Call(__e, ShenFunc(symmap), V2256, gen3363)

					__e.TailApply(ShenFunc(symcons), gen3362, gen3364)

					return

				} else {
					__e.TailApply(V2256, V2257)

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("map"), gen3367)

		gen3368 := MakeNative(func(__e Evaluator) {
			V2259 := __e.Get(1)
			_ = V2259
			__e.TailApply(ShenFunc(symshen_4length_1h), V2259, MakeNumber(0))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("length"), gen3368)

		gen3372 := MakeNative(func(__e Evaluator) {
			V2262 := __e.Get(1)
			_ = V2262
			V2263 := __e.Get(2)
			_ = V2263
			gen3371 := Call(__e, ShenFunc(sym_a), Nil, V2262)

			if True == gen3371 {
				__e.Return(V2263)
				return
			} else {
				gen3369 := Call(__e, ShenFunc(symtl), V2262)

				gen3370 := Call(__e, ShenFunc(sym_7), V2263, MakeNumber(1))

				__e.TailApply(ShenFunc(symshen_4length_1h), gen3369, gen3370)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.length-h"), gen3372)

		gen3379 := MakeNative(func(__e Evaluator) {
			V2275 := __e.Get(1)
			_ = V2275
			V2276 := __e.Get(2)
			_ = V2276
			gen3378 := Call(__e, ShenFunc(sym_a), V2276, V2275)

			if True == gen3378 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen3377 := Call(__e, ShenFunc(symcons_2), V2276)

				if True == gen3377 {
					gen3373 := Call(__e, ShenFunc(symhd), V2276)

					gen3374 := Call(__e, ShenFunc(symoccurrences), V2275, gen3373)

					gen3375 := Call(__e, ShenFunc(symtl), V2276)

					gen3376 := Call(__e, ShenFunc(symoccurrences), V2275, gen3375)

					__e.TailApply(ShenFunc(sym_7), gen3374, gen3376)

					return

				} else {
					__e.Return(MakeNumber(0))
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("occurrences"), gen3379)

		gen3390 := MakeNative(func(__e Evaluator) {
			V2283 := __e.Get(1)
			_ = V2283
			V2284 := __e.Get(2)
			_ = V2284
			gen3388 := Call(__e, ShenFunc(sym_a), MakeNumber(1), V2283)

			var gen3389 Obj
			if True == gen3388 {
				gen3387 := Call(__e, ShenFunc(symcons_2), V2284)

				if True == gen3387 {
					gen3389 = True
				} else {
					gen3389 = False
				}

			} else {
				gen3389 = False
			}
			if True == gen3389 {
				__e.TailApply(ShenFunc(symhd), V2284)

				return
			} else {
				gen3386 := Call(__e, ShenFunc(symcons_2), V2284)

				if True == gen3386 {
					gen3384 := Call(__e, ShenFunc(sym_1), V2283, MakeNumber(1))

					gen3385 := Call(__e, ShenFunc(symtl), V2284)

					__e.TailApply(ShenFunc(symnth), gen3384, gen3385)

					return

				} else {
					gen3380 := Call(__e, ShenFunc(symshen_4app), V2284, MakeString("\n"), MakeSymbol("shen.a"))

					gen3381 := Call(__e, ShenFunc(symcn), MakeString(", "), gen3380)

					gen3382 := Call(__e, ShenFunc(symshen_4app), V2283, gen3381, MakeSymbol("shen.a"))

					gen3383 := Call(__e, ShenFunc(symcn), MakeString("nth applied to "), gen3382)

					__e.TailApply(ShenFunc(symsimple_1error), gen3383)

					return

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("nth"), gen3390)

		gen3395 := MakeNative(func(__e Evaluator) {
			V2286 := __e.Get(1)
			_ = V2286
			gen3394 := Call(__e, ShenFunc(symnumber_2), V2286)

			if True == gen3394 {
				gen3391 := Call(__e, ShenFunc(symshen_4abs), V2286)

				Abs := gen3391
				gen3392 := Call(__e, ShenFunc(symshen_4magless), Abs, MakeNumber(1))

				gen3393 := Call(__e, ShenFunc(symshen_4integer_1test_2), Abs, gen3392)

				if True == gen3393 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("integer?"), gen3395)

		gen3397 := MakeNative(func(__e Evaluator) {
			V2288 := __e.Get(1)
			_ = V2288
			gen3396 := Call(__e, ShenFunc(sym_6), V2288, MakeNumber(0))

			if True == gen3396 {
				__e.Return(V2288)
				return
			} else {
				__e.TailApply(ShenFunc(sym_1), MakeNumber(0), V2288)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abs"), gen3397)

		gen3400 := MakeNative(func(__e Evaluator) {
			V2291 := __e.Get(1)
			_ = V2291
			V2292 := __e.Get(2)
			_ = V2292
			gen3398 := Call(__e, ShenFunc(sym_d), V2292, MakeNumber(2))

			Nx2 := gen3398
			gen3399 := Call(__e, ShenFunc(sym_6), Nx2, V2291)

			if True == gen3399 {
				__e.Return(V2292)
				return
			} else {
				__e.TailApply(ShenFunc(symshen_4magless), V2291, Nx2)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.magless"), gen3400)

		gen3405 := MakeNative(func(__e Evaluator) {
			V2298 := __e.Get(1)
			_ = V2298
			V2299 := __e.Get(2)
			_ = V2299
			gen3404 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V2298)

			if True == gen3404 {
				__e.Return(True)
				return
			} else {
				gen3403 := Call(__e, ShenFunc(sym_6), MakeNumber(1), V2298)

				if True == gen3403 {
					__e.Return(False)
					return
				} else {
					gen3401 := Call(__e, ShenFunc(sym_1), V2298, V2299)

					Abs_1N := gen3401
					gen3402 := Call(__e, ShenFunc(sym_6), MakeNumber(0), Abs_1N)

					if True == gen3402 {
						__e.TailApply(ShenFunc(syminteger_2), V2298)

						return
					} else {
						__e.TailApply(ShenFunc(symshen_4integer_1test_2), Abs_1N, V2299)

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.integer-test?"), gen3405)

		gen3412 := MakeNative(func(__e Evaluator) {
			V2304 := __e.Get(1)
			_ = V2304
			V2305 := __e.Get(2)
			_ = V2305
			gen3411 := Call(__e, ShenFunc(sym_a), Nil, V2305)

			if True == gen3411 {
				__e.Return(Nil)
				return
			} else {
				gen3410 := Call(__e, ShenFunc(symcons_2), V2305)

				if True == gen3410 {
					gen3406 := Call(__e, ShenFunc(symhd), V2305)

					gen3407 := Call(__e, V2304, gen3406)

					gen3408 := Call(__e, ShenFunc(symtl), V2305)

					gen3409 := Call(__e, ShenFunc(symmapcan), V2304, gen3408)

					__e.TailApply(ShenFunc(symappend), gen3407, gen3409)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("mapcan"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("mapcan"), gen3412)

		gen3414 := MakeNative(func(__e Evaluator) {
			V2317 := __e.Get(1)
			_ = V2317
			V2318 := __e.Get(2)
			_ = V2318
			gen3413 := Call(__e, ShenFunc(sym_a), V2318, V2317)

			if True == gen3413 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("=="), gen3414)

		gen3415 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symsimple_1error), MakeString(""))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("abort"), gen3415)

		gen3422 := MakeNative(func(__e Evaluator) {
			V2320 := __e.Get(1)
			_ = V2320
			gen3421 := Call(__e, ShenFunc(symsymbol_2), V2320)

			if True == gen3421 {
				gen3416 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeSymbol("shen.this-symbol-is-unbound"))
					return
				}, 1)
				gen3417 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(symvalue), V2320)

					return
				}, 0)
				gen3418 := Try(__e, gen3417).Catch(gen3416)
				Val := gen3418
				gen3419 := Call(__e, ShenFunc(sym_a), Val, MakeSymbol("shen.this-symbol-is-unbound"))

				var gen3420 Obj
				if True == gen3419 {
					gen3420 = False
				} else {
					gen3420 = True
				}
				if True == gen3420 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("bound?"), gen3422)

		gen3428 := MakeNative(func(__e Evaluator) {
			V2322 := __e.Get(1)
			_ = V2322
			gen3427 := Call(__e, ShenFunc(sym_a), MakeString(""), V2322)

			if True == gen3427 {
				__e.Return(Nil)
				return
			} else {
				gen3423 := Call(__e, ShenFunc(sympos), V2322, MakeNumber(0))

				gen3424 := Call(__e, ShenFunc(symstring_1_6n), gen3423)

				gen3425 := Call(__e, ShenFunc(symtlstr), V2322)

				gen3426 := Call(__e, ShenFunc(symshen_4string_1_6bytes), gen3425)

				__e.TailApply(ShenFunc(symcons), gen3424, gen3426)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.string->bytes"), gen3428)

		gen3429 := MakeNative(func(__e Evaluator) {
			V2324 := __e.Get(1)
			_ = V2324
			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*maxinferences*"), V2324)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("maxinferences"), gen3429)

		gen3430 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*infs*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("inferences"), gen3430)

		gen3431 := MakeNative(func(__e Evaluator) {
			V2326 := __e.Get(1)
			_ = V2326
			__e.Return(V2326)
			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("protect"), gen3431)

		gen3432 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*stoutput*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("stoutput"), gen3432)

		gen3433 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*sterror*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("sterror"), gen3433)

		gen3438 := MakeNative(func(__e Evaluator) {
			V2328 := __e.Get(1)
			_ = V2328
			gen3434 := Call(__e, ShenFunc(symintern), V2328)

			Symbol := gen3434
			gen3437 := Call(__e, ShenFunc(symsymbol_2), Symbol)

			if True == gen3437 {
				__e.Return(Symbol)
				return
			} else {
				gen3435 := Call(__e, ShenFunc(symshen_4app), V2328, MakeString(" to a symbol"), MakeSymbol("shen.s"))

				gen3436 := Call(__e, ShenFunc(symcn), MakeString("cannot intern "), gen3435)

				__e.TailApply(ShenFunc(symsimple_1error), gen3436)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("string->symbol"), gen3438)

		gen3441 := MakeNative(func(__e Evaluator) {
			V2334 := __e.Get(1)
			_ = V2334
			gen3440 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V2334)

			if True == gen3440 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*optimise*"), True)

				return
			} else {
				gen3439 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V2334)

				if True == gen3439 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*optimise*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("optimise expects a + or a -.\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("optimise"), gen3441)

		gen3442 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*os*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("os"), gen3442)

		gen3443 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*language*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("language"), gen3443)

		gen3444 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*version*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("version"), gen3444)

		gen3445 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*port*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("port"), gen3445)

		gen3446 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*porters*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("porters"), gen3446)

		gen3447 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*implementation*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("implementation"), gen3447)

		gen3448 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*release*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("release"), gen3448)

		gen3451 := MakeNative(func(__e Evaluator) {
			V2336 := __e.Get(1)
			_ = V2336
			gen3449 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)
			gen3450 := MakeNative(func(__e Evaluator) {
				Call(__e, ShenFunc(symexternal), V2336)
				__e.Return(True)
				return

			}, 0)
			__e.Return(Try(__e, gen3450).Catch(gen3449))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("package?"), gen3451)

		gen3452 := MakeNative(func(__e Evaluator) {
			V2338 := __e.Get(1)
			_ = V2338
			__e.TailApply(ShenFunc(symshen_4lookup_1func), V2338)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("function"), gen3452)

		gen3457 := MakeNative(func(__e Evaluator) {
			V2340 := __e.Get(1)
			_ = V2340
			gen3454 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				gen3453 := Call(__e, ShenFunc(symshen_4app), V2340, MakeString(" has no lambda expansion\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen3453)

				return

			}, 1)
			gen3456 := MakeNative(func(__e Evaluator) {
				gen3455 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V2340, MakeSymbol("shen.lambda-form"), gen3455)

				return

			}, 0)
			__e.Return(Try(__e, gen3456).Catch(gen3454))
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.lookup-func"), gen3457)

		return

	}, 0))
}
