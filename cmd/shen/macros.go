package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen5933 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V748 := __args[0]
			_ = V748
			gen5929 := Call(__e, ShenFunc(symvalue), MakeSymbol("*macros*"))

			gen5930 := Call(__e, ShenFunc(symshen_4compose), gen5929, V748)

			Y := gen5930
			gen5932 := Call(__e, ShenFunc(sym_a), V748, Y)

			if True == gen5932 {
				__e.Return(V748)
				return
			} else {
				gen5931 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Z := __args[0]
					_ = Z
					__e.TailApply(ShenFunc(symmacroexpand), Z)

					return
				}, 1)
				__e.TailApply(ShenFunc(symshen_4walk), gen5931, Y)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("macroexpand"), gen5933)

		gen5947 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V750 := __args[0]
			_ = V750
			gen5945 := Call(__e, ShenFunc(symcons_2), V750)

			var gen5946 Obj
			if True == gen5945 {
				gen5942 := Call(__e, ShenFunc(symhd), V750)

				gen5943 := Call(__e, ShenFunc(sym_a), MakeSymbol("error"), gen5942)

				var gen5944 Obj
				if True == gen5943 {
					gen5940 := Call(__e, ShenFunc(symtl), V750)

					gen5941 := Call(__e, ShenFunc(symcons_2), gen5940)

					if True == gen5941 {
						gen5944 = True
					} else {
						gen5944 = False
					}

				} else {
					gen5944 = False
				}
				if True == gen5944 {
					gen5946 = True
				} else {
					gen5946 = False
				}

			} else {
				gen5946 = False
			}
			if True == gen5946 {
				gen5934 := Call(__e, ShenFunc(symtl), V750)

				gen5935 := Call(__e, ShenFunc(symhd), gen5934)

				gen5936 := Call(__e, ShenFunc(symtl), V750)

				gen5937 := Call(__e, ShenFunc(symtl), gen5936)

				gen5938 := Call(__e, ShenFunc(symshen_4mkstr), gen5935, gen5937)

				gen5939 := Call(__e, ShenFunc(symcons), gen5938, Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("simple-error"), gen5939)

				return

			} else {
				__e.Return(V750)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.error-macro"), gen5947)

		gen5979 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V752 := __args[0]
			_ = V752
			gen5977 := Call(__e, ShenFunc(symcons_2), V752)

			var gen5978 Obj
			if True == gen5977 {
				gen5974 := Call(__e, ShenFunc(symhd), V752)

				gen5975 := Call(__e, ShenFunc(sym_a), MakeSymbol("output"), gen5974)

				var gen5976 Obj
				if True == gen5975 {
					gen5972 := Call(__e, ShenFunc(symtl), V752)

					gen5973 := Call(__e, ShenFunc(symcons_2), gen5972)

					if True == gen5973 {
						gen5976 = True
					} else {
						gen5976 = False
					}

				} else {
					gen5976 = False
				}
				if True == gen5976 {
					gen5978 = True
				} else {
					gen5978 = False
				}

			} else {
				gen5978 = False
			}
			if True == gen5978 {
				gen5964 := Call(__e, ShenFunc(symtl), V752)

				gen5965 := Call(__e, ShenFunc(symhd), gen5964)

				gen5966 := Call(__e, ShenFunc(symtl), V752)

				gen5967 := Call(__e, ShenFunc(symtl), gen5966)

				gen5968 := Call(__e, ShenFunc(symshen_4mkstr), gen5965, gen5967)

				gen5969 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), Nil)

				gen5970 := Call(__e, ShenFunc(symcons), gen5969, Nil)

				gen5971 := Call(__e, ShenFunc(symcons), gen5968, gen5970)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.prhush"), gen5971)

				return

			} else {
				gen5962 := Call(__e, ShenFunc(symcons_2), V752)

				var gen5963 Obj
				if True == gen5962 {
					gen5959 := Call(__e, ShenFunc(symhd), V752)

					gen5960 := Call(__e, ShenFunc(sym_a), MakeSymbol("pr"), gen5959)

					var gen5961 Obj
					if True == gen5960 {
						gen5956 := Call(__e, ShenFunc(symtl), V752)

						gen5957 := Call(__e, ShenFunc(symcons_2), gen5956)

						var gen5958 Obj
						if True == gen5957 {
							gen5953 := Call(__e, ShenFunc(symtl), V752)

							gen5954 := Call(__e, ShenFunc(symtl), gen5953)

							gen5955 := Call(__e, ShenFunc(sym_a), Nil, gen5954)

							if True == gen5955 {
								gen5958 = True
							} else {
								gen5958 = False
							}

						} else {
							gen5958 = False
						}
						if True == gen5958 {
							gen5961 = True
						} else {
							gen5961 = False
						}

					} else {
						gen5961 = False
					}
					if True == gen5961 {
						gen5963 = True
					} else {
						gen5963 = False
					}

				} else {
					gen5963 = False
				}
				if True == gen5963 {
					gen5948 := Call(__e, ShenFunc(symtl), V752)

					gen5949 := Call(__e, ShenFunc(symhd), gen5948)

					gen5950 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), Nil)

					gen5951 := Call(__e, ShenFunc(symcons), gen5950, Nil)

					gen5952 := Call(__e, ShenFunc(symcons), gen5949, gen5951)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("pr"), gen5952)

					return

				} else {
					__e.Return(V752)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.output-macro"), gen5979)

		gen5991 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V754 := __args[0]
			_ = V754
			gen5989 := Call(__e, ShenFunc(symcons_2), V754)

			var gen5990 Obj
			if True == gen5989 {
				gen5986 := Call(__e, ShenFunc(symhd), V754)

				gen5987 := Call(__e, ShenFunc(sym_a), MakeSymbol("make-string"), gen5986)

				var gen5988 Obj
				if True == gen5987 {
					gen5984 := Call(__e, ShenFunc(symtl), V754)

					gen5985 := Call(__e, ShenFunc(symcons_2), gen5984)

					if True == gen5985 {
						gen5988 = True
					} else {
						gen5988 = False
					}

				} else {
					gen5988 = False
				}
				if True == gen5988 {
					gen5990 = True
				} else {
					gen5990 = False
				}

			} else {
				gen5990 = False
			}
			if True == gen5990 {
				gen5980 := Call(__e, ShenFunc(symtl), V754)

				gen5981 := Call(__e, ShenFunc(symhd), gen5980)

				gen5982 := Call(__e, ShenFunc(symtl), V754)

				gen5983 := Call(__e, ShenFunc(symtl), gen5982)

				__e.TailApply(ShenFunc(symshen_4mkstr), gen5981, gen5983)

				return

			} else {
				__e.Return(V754)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.make-string-macro"), gen5991)

		gen6044 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V756 := __args[0]
			_ = V756
			gen6042 := Call(__e, ShenFunc(symcons_2), V756)

			var gen6043 Obj
			if True == gen6042 {
				gen6039 := Call(__e, ShenFunc(symhd), V756)

				gen6040 := Call(__e, ShenFunc(sym_a), MakeSymbol("lineread"), gen6039)

				var gen6041 Obj
				if True == gen6040 {
					gen6037 := Call(__e, ShenFunc(symtl), V756)

					gen6038 := Call(__e, ShenFunc(sym_a), Nil, gen6037)

					if True == gen6038 {
						gen6041 = True
					} else {
						gen6041 = False
					}

				} else {
					gen6041 = False
				}
				if True == gen6041 {
					gen6043 = True
				} else {
					gen6043 = False
				}

			} else {
				gen6043 = False
			}
			if True == gen6043 {
				gen6035 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

				gen6036 := Call(__e, ShenFunc(symcons), gen6035, Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("lineread"), gen6036)

				return

			} else {
				gen6033 := Call(__e, ShenFunc(symcons_2), V756)

				var gen6034 Obj
				if True == gen6033 {
					gen6030 := Call(__e, ShenFunc(symhd), V756)

					gen6031 := Call(__e, ShenFunc(sym_a), MakeSymbol("input"), gen6030)

					var gen6032 Obj
					if True == gen6031 {
						gen6028 := Call(__e, ShenFunc(symtl), V756)

						gen6029 := Call(__e, ShenFunc(sym_a), Nil, gen6028)

						if True == gen6029 {
							gen6032 = True
						} else {
							gen6032 = False
						}

					} else {
						gen6032 = False
					}
					if True == gen6032 {
						gen6034 = True
					} else {
						gen6034 = False
					}

				} else {
					gen6034 = False
				}
				if True == gen6034 {
					gen6026 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

					gen6027 := Call(__e, ShenFunc(symcons), gen6026, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("input"), gen6027)

					return

				} else {
					gen6024 := Call(__e, ShenFunc(symcons_2), V756)

					var gen6025 Obj
					if True == gen6024 {
						gen6021 := Call(__e, ShenFunc(symhd), V756)

						gen6022 := Call(__e, ShenFunc(sym_a), MakeSymbol("read"), gen6021)

						var gen6023 Obj
						if True == gen6022 {
							gen6019 := Call(__e, ShenFunc(symtl), V756)

							gen6020 := Call(__e, ShenFunc(sym_a), Nil, gen6019)

							if True == gen6020 {
								gen6023 = True
							} else {
								gen6023 = False
							}

						} else {
							gen6023 = False
						}
						if True == gen6023 {
							gen6025 = True
						} else {
							gen6025 = False
						}

					} else {
						gen6025 = False
					}
					if True == gen6025 {
						gen6017 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

						gen6018 := Call(__e, ShenFunc(symcons), gen6017, Nil)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("read"), gen6018)

						return

					} else {
						gen6015 := Call(__e, ShenFunc(symcons_2), V756)

						var gen6016 Obj
						if True == gen6015 {
							gen6012 := Call(__e, ShenFunc(symhd), V756)

							gen6013 := Call(__e, ShenFunc(sym_a), MakeSymbol("input+"), gen6012)

							var gen6014 Obj
							if True == gen6013 {
								gen6009 := Call(__e, ShenFunc(symtl), V756)

								gen6010 := Call(__e, ShenFunc(symcons_2), gen6009)

								var gen6011 Obj
								if True == gen6010 {
									gen6006 := Call(__e, ShenFunc(symtl), V756)

									gen6007 := Call(__e, ShenFunc(symtl), gen6006)

									gen6008 := Call(__e, ShenFunc(sym_a), Nil, gen6007)

									if True == gen6008 {
										gen6011 = True
									} else {
										gen6011 = False
									}

								} else {
									gen6011 = False
								}
								if True == gen6011 {
									gen6014 = True
								} else {
									gen6014 = False
								}

							} else {
								gen6014 = False
							}
							if True == gen6014 {
								gen6016 = True
							} else {
								gen6016 = False
							}

						} else {
							gen6016 = False
						}
						if True == gen6016 {
							gen6001 := Call(__e, ShenFunc(symtl), V756)

							gen6002 := Call(__e, ShenFunc(symhd), gen6001)

							gen6003 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

							gen6004 := Call(__e, ShenFunc(symcons), gen6003, Nil)

							gen6005 := Call(__e, ShenFunc(symcons), gen6002, gen6004)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("input+"), gen6005)

							return

						} else {
							gen5999 := Call(__e, ShenFunc(symcons_2), V756)

							var gen6000 Obj
							if True == gen5999 {
								gen5996 := Call(__e, ShenFunc(symhd), V756)

								gen5997 := Call(__e, ShenFunc(sym_a), MakeSymbol("read-byte"), gen5996)

								var gen5998 Obj
								if True == gen5997 {
									gen5994 := Call(__e, ShenFunc(symtl), V756)

									gen5995 := Call(__e, ShenFunc(sym_a), Nil, gen5994)

									if True == gen5995 {
										gen5998 = True
									} else {
										gen5998 = False
									}

								} else {
									gen5998 = False
								}
								if True == gen5998 {
									gen6000 = True
								} else {
									gen6000 = False
								}

							} else {
								gen6000 = False
							}
							if True == gen6000 {
								gen5992 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

								gen5993 := Call(__e, ShenFunc(symcons), gen5992, Nil)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("read-byte"), gen5993)

								return

							} else {
								__e.Return(V756)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.input-macro"), gen6044)

		gen6050 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V759 := __args[0]
			_ = V759
			V760 := __args[1]
			_ = V760
			gen6049 := Call(__e, ShenFunc(sym_a), Nil, V759)

			if True == gen6049 {
				__e.Return(V760)
				return
			} else {
				gen6048 := Call(__e, ShenFunc(symcons_2), V759)

				if True == gen6048 {
					gen6045 := Call(__e, ShenFunc(symtl), V759)

					gen6046 := Call(__e, ShenFunc(symhd), V759)

					f31 := gen6046
					gen6047 := Call(__e, f31, V760)

					__e.TailApply(ShenFunc(symshen_4compose), gen6045, gen6047)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.compose"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compose"), gen6050)

		gen6089 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V762 := __args[0]
			_ = V762
			gen6087 := Call(__e, ShenFunc(symcons_2), V762)

			var gen6088 Obj
			if True == gen6087 {
				gen6084 := Call(__e, ShenFunc(symhd), V762)

				gen6085 := Call(__e, ShenFunc(sym_a), MakeSymbol("compile"), gen6084)

				var gen6086 Obj
				if True == gen6085 {
					gen6081 := Call(__e, ShenFunc(symtl), V762)

					gen6082 := Call(__e, ShenFunc(symcons_2), gen6081)

					var gen6083 Obj
					if True == gen6082 {
						gen6077 := Call(__e, ShenFunc(symtl), V762)

						gen6078 := Call(__e, ShenFunc(symtl), gen6077)

						gen6079 := Call(__e, ShenFunc(symcons_2), gen6078)

						var gen6080 Obj
						if True == gen6079 {
							gen6073 := Call(__e, ShenFunc(symtl), V762)

							gen6074 := Call(__e, ShenFunc(symtl), gen6073)

							gen6075 := Call(__e, ShenFunc(symtl), gen6074)

							gen6076 := Call(__e, ShenFunc(sym_a), Nil, gen6075)

							if True == gen6076 {
								gen6080 = True
							} else {
								gen6080 = False
							}

						} else {
							gen6080 = False
						}
						if True == gen6080 {
							gen6083 = True
						} else {
							gen6083 = False
						}

					} else {
						gen6083 = False
					}
					if True == gen6083 {
						gen6086 = True
					} else {
						gen6086 = False
					}

				} else {
					gen6086 = False
				}
				if True == gen6086 {
					gen6088 = True
				} else {
					gen6088 = False
				}

			} else {
				gen6088 = False
			}
			if True == gen6088 {
				gen6051 := Call(__e, ShenFunc(symtl), V762)

				gen6052 := Call(__e, ShenFunc(symhd), gen6051)

				gen6053 := Call(__e, ShenFunc(symtl), V762)

				gen6054 := Call(__e, ShenFunc(symtl), gen6053)

				gen6055 := Call(__e, ShenFunc(symhd), gen6054)

				gen6056 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), Nil)

				gen6057 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen6056)

				gen6058 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), Nil)

				gen6059 := Call(__e, ShenFunc(symcons), MakeString("parse error here: ~S~%"), gen6058)

				gen6060 := Call(__e, ShenFunc(symcons), MakeSymbol("error"), gen6059)

				gen6061 := Call(__e, ShenFunc(symcons), MakeString("parse error~%"), Nil)

				gen6062 := Call(__e, ShenFunc(symcons), MakeSymbol("error"), gen6061)

				gen6063 := Call(__e, ShenFunc(symcons), gen6062, Nil)

				gen6064 := Call(__e, ShenFunc(symcons), gen6060, gen6063)

				gen6065 := Call(__e, ShenFunc(symcons), gen6057, gen6064)

				gen6066 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen6065)

				gen6067 := Call(__e, ShenFunc(symcons), gen6066, Nil)

				gen6068 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), gen6067)

				gen6069 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen6068)

				gen6070 := Call(__e, ShenFunc(symcons), gen6069, Nil)

				gen6071 := Call(__e, ShenFunc(symcons), gen6055, gen6070)

				gen6072 := Call(__e, ShenFunc(symcons), gen6052, gen6071)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("compile"), gen6072)

				return

			} else {
				__e.Return(V762)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile-macro"), gen6089)

		gen6106 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V764 := __args[0]
			_ = V764
			gen6104 := Call(__e, ShenFunc(symcons_2), V764)

			var gen6105 Obj
			if True == gen6104 {
				gen6102 := Call(__e, ShenFunc(symhd), V764)

				gen6103 := Call(__e, ShenFunc(sym_a), MakeSymbol("prolog?"), gen6102)

				if True == gen6103 {
					gen6105 = True
				} else {
					gen6105 = False
				}

			} else {
				gen6105 = False
			}
			if True == gen6105 {
				gen6090 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.start-new-prolog-process"), Nil)

				gen6091 := Call(__e, ShenFunc(symtl), V764)

				gen6092 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), MakeSymbol("NPP"), gen6091)

				Calls := gen6092
				gen6093 := Call(__e, ShenFunc(symtl), V764)

				gen6094 := Call(__e, ShenFunc(symshen_4extract__vars), gen6093)

				Vs := gen6094
				gen6095 := Call(__e, ShenFunc(symtl), V764)

				gen6096 := Call(__e, ShenFunc(symshen_4externally_1bound), gen6095)

				External := gen6096
				gen6097 := Call(__e, ShenFunc(symdifference), Vs, External)

				PrologVs := gen6097
				gen6098 := Call(__e, ShenFunc(symshen_4locally_1bind_1prolog_1vs), MakeSymbol("NPP"), PrologVs, Calls)

				gen6099 := Call(__e, ShenFunc(symcons), gen6098, Nil)

				gen6100 := Call(__e, ShenFunc(symcons), gen6090, gen6099)

				gen6101 := Call(__e, ShenFunc(symcons), MakeSymbol("NPP"), gen6100)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen6101)

				return

			} else {
				__e.Return(V764)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-macro"), gen6106)

		gen6123 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V770 := __args[0]
			_ = V770
			gen6121 := Call(__e, ShenFunc(symcons_2), V770)

			var gen6122 Obj
			if True == gen6121 {
				gen6118 := Call(__e, ShenFunc(symhd), V770)

				gen6119 := Call(__e, ShenFunc(sym_a), MakeSymbol("receive"), gen6118)

				var gen6120 Obj
				if True == gen6119 {
					gen6115 := Call(__e, ShenFunc(symtl), V770)

					gen6116 := Call(__e, ShenFunc(symcons_2), gen6115)

					var gen6117 Obj
					if True == gen6116 {
						gen6112 := Call(__e, ShenFunc(symtl), V770)

						gen6113 := Call(__e, ShenFunc(symtl), gen6112)

						gen6114 := Call(__e, ShenFunc(sym_a), Nil, gen6113)

						if True == gen6114 {
							gen6117 = True
						} else {
							gen6117 = False
						}

					} else {
						gen6117 = False
					}
					if True == gen6117 {
						gen6120 = True
					} else {
						gen6120 = False
					}

				} else {
					gen6120 = False
				}
				if True == gen6120 {
					gen6122 = True
				} else {
					gen6122 = False
				}

			} else {
				gen6122 = False
			}
			if True == gen6122 {
				__e.TailApply(ShenFunc(symtl), V770)

				return
			} else {
				gen6111 := Call(__e, ShenFunc(symcons_2), V770)

				if True == gen6111 {
					gen6107 := Call(__e, ShenFunc(symhd), V770)

					gen6108 := Call(__e, ShenFunc(symshen_4externally_1bound), gen6107)

					gen6109 := Call(__e, ShenFunc(symtl), V770)

					gen6110 := Call(__e, ShenFunc(symshen_4externally_1bound), gen6109)

					__e.TailApply(ShenFunc(symunion), gen6108, gen6110)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.externally-bound"), gen6123)

		gen6134 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V788 := __args[0]
			_ = V788
			V789 := __args[1]
			_ = V789
			V790 := __args[2]
			_ = V790
			gen6133 := Call(__e, ShenFunc(sym_a), Nil, V789)

			if True == gen6133 {
				__e.Return(V790)
				return
			} else {
				gen6132 := Call(__e, ShenFunc(symcons_2), V789)

				if True == gen6132 {
					gen6124 := Call(__e, ShenFunc(symhd), V789)

					gen6125 := Call(__e, ShenFunc(symcons), V788, Nil)

					gen6126 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.newpv"), gen6125)

					gen6127 := Call(__e, ShenFunc(symtl), V789)

					gen6128 := Call(__e, ShenFunc(symshen_4locally_1bind_1prolog_1vs), V788, gen6127, V790)

					gen6129 := Call(__e, ShenFunc(symcons), gen6128, Nil)

					gen6130 := Call(__e, ShenFunc(symcons), gen6126, gen6129)

					gen6131 := Call(__e, ShenFunc(symcons), gen6124, gen6130)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen6131)

					return

				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("implementation error inp locally-bind-prolog-vs"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.locally-bind-prolog-vs"), gen6134)

		gen6307 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V803 := __args[0]
			_ = V803
			V804 := __args[1]
			_ = V804
			gen6306 := Call(__e, ShenFunc(sym_a), Nil, V804)

			if True == gen6306 {
				__e.Return(True)
				return
			} else {
				gen6304 := Call(__e, ShenFunc(symcons_2), V804)

				var gen6305 Obj
				if True == gen6304 {
					gen6302 := Call(__e, ShenFunc(symhd), V804)

					gen6303 := Call(__e, ShenFunc(sym_a), MakeSymbol("!"), gen6302)

					if True == gen6303 {
						gen6305 = True
					} else {
						gen6305 = False
					}

				} else {
					gen6305 = False
				}
				if True == gen6305 {
					gen6295 := Call(__e, ShenFunc(symtl), V804)

					gen6296 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen6295)

					gen6297 := Call(__e, ShenFunc(symcons), gen6296, Nil)

					gen6298 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen6297)

					gen6299 := Call(__e, ShenFunc(symcons), gen6298, Nil)

					gen6300 := Call(__e, ShenFunc(symcons), V803, gen6299)

					gen6301 := Call(__e, ShenFunc(symcons), False, gen6300)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("cut"), gen6301)

					return

				} else {
					gen6293 := Call(__e, ShenFunc(symcons_2), V804)

					var gen6294 Obj
					if True == gen6293 {
						gen6290 := Call(__e, ShenFunc(symhd), V804)

						gen6291 := Call(__e, ShenFunc(symcons_2), gen6290)

						var gen6292 Obj
						if True == gen6291 {
							gen6286 := Call(__e, ShenFunc(symhd), V804)

							gen6287 := Call(__e, ShenFunc(symhd), gen6286)

							gen6288 := Call(__e, ShenFunc(sym_a), MakeSymbol("when"), gen6287)

							var gen6289 Obj
							if True == gen6288 {
								gen6282 := Call(__e, ShenFunc(symhd), V804)

								gen6283 := Call(__e, ShenFunc(symtl), gen6282)

								gen6284 := Call(__e, ShenFunc(symcons_2), gen6283)

								var gen6285 Obj
								if True == gen6284 {
									gen6278 := Call(__e, ShenFunc(symhd), V804)

									gen6279 := Call(__e, ShenFunc(symtl), gen6278)

									gen6280 := Call(__e, ShenFunc(symtl), gen6279)

									gen6281 := Call(__e, ShenFunc(sym_a), Nil, gen6280)

									if True == gen6281 {
										gen6285 = True
									} else {
										gen6285 = False
									}

								} else {
									gen6285 = False
								}
								if True == gen6285 {
									gen6289 = True
								} else {
									gen6289 = False
								}

							} else {
								gen6289 = False
							}
							if True == gen6289 {
								gen6292 = True
							} else {
								gen6292 = False
							}

						} else {
							gen6292 = False
						}
						if True == gen6292 {
							gen6294 = True
						} else {
							gen6294 = False
						}

					} else {
						gen6294 = False
					}
					if True == gen6294 {
						gen6267 := Call(__e, ShenFunc(symhd), V804)

						gen6268 := Call(__e, ShenFunc(symtl), gen6267)

						gen6269 := Call(__e, ShenFunc(symhd), gen6268)

						gen6270 := Call(__e, ShenFunc(symshen_4insert_1deref), gen6269, V803)

						gen6271 := Call(__e, ShenFunc(symtl), V804)

						gen6272 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen6271)

						gen6273 := Call(__e, ShenFunc(symcons), gen6272, Nil)

						gen6274 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen6273)

						gen6275 := Call(__e, ShenFunc(symcons), gen6274, Nil)

						gen6276 := Call(__e, ShenFunc(symcons), V803, gen6275)

						gen6277 := Call(__e, ShenFunc(symcons), gen6270, gen6276)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen6277)

						return

					} else {
						gen6265 := Call(__e, ShenFunc(symcons_2), V804)

						var gen6266 Obj
						if True == gen6265 {
							gen6262 := Call(__e, ShenFunc(symhd), V804)

							gen6263 := Call(__e, ShenFunc(symcons_2), gen6262)

							var gen6264 Obj
							if True == gen6263 {
								gen6258 := Call(__e, ShenFunc(symhd), V804)

								gen6259 := Call(__e, ShenFunc(symhd), gen6258)

								gen6260 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen6259)

								var gen6261 Obj
								if True == gen6260 {
									gen6254 := Call(__e, ShenFunc(symhd), V804)

									gen6255 := Call(__e, ShenFunc(symtl), gen6254)

									gen6256 := Call(__e, ShenFunc(symcons_2), gen6255)

									var gen6257 Obj
									if True == gen6256 {
										gen6249 := Call(__e, ShenFunc(symhd), V804)

										gen6250 := Call(__e, ShenFunc(symtl), gen6249)

										gen6251 := Call(__e, ShenFunc(symtl), gen6250)

										gen6252 := Call(__e, ShenFunc(symcons_2), gen6251)

										var gen6253 Obj
										if True == gen6252 {
											gen6244 := Call(__e, ShenFunc(symhd), V804)

											gen6245 := Call(__e, ShenFunc(symtl), gen6244)

											gen6246 := Call(__e, ShenFunc(symtl), gen6245)

											gen6247 := Call(__e, ShenFunc(symtl), gen6246)

											gen6248 := Call(__e, ShenFunc(sym_a), Nil, gen6247)

											if True == gen6248 {
												gen6253 = True
											} else {
												gen6253 = False
											}

										} else {
											gen6253 = False
										}
										if True == gen6253 {
											gen6257 = True
										} else {
											gen6257 = False
										}

									} else {
										gen6257 = False
									}
									if True == gen6257 {
										gen6261 = True
									} else {
										gen6261 = False
									}

								} else {
									gen6261 = False
								}
								if True == gen6261 {
									gen6264 = True
								} else {
									gen6264 = False
								}

							} else {
								gen6264 = False
							}
							if True == gen6264 {
								gen6266 = True
							} else {
								gen6266 = False
							}

						} else {
							gen6266 = False
						}
						if True == gen6266 {
							gen6228 := Call(__e, ShenFunc(symhd), V804)

							gen6229 := Call(__e, ShenFunc(symtl), gen6228)

							gen6230 := Call(__e, ShenFunc(symhd), gen6229)

							gen6231 := Call(__e, ShenFunc(symhd), V804)

							gen6232 := Call(__e, ShenFunc(symtl), gen6231)

							gen6233 := Call(__e, ShenFunc(symtl), gen6232)

							gen6234 := Call(__e, ShenFunc(symhd), gen6233)

							gen6235 := Call(__e, ShenFunc(symshen_4insert_1deref), gen6234, V803)

							gen6236 := Call(__e, ShenFunc(symtl), V804)

							gen6237 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen6236)

							gen6238 := Call(__e, ShenFunc(symcons), gen6237, Nil)

							gen6239 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen6238)

							gen6240 := Call(__e, ShenFunc(symcons), gen6239, Nil)

							gen6241 := Call(__e, ShenFunc(symcons), V803, gen6240)

							gen6242 := Call(__e, ShenFunc(symcons), gen6235, gen6241)

							gen6243 := Call(__e, ShenFunc(symcons), gen6230, gen6242)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen6243)

							return

						} else {
							gen6226 := Call(__e, ShenFunc(symcons_2), V804)

							var gen6227 Obj
							if True == gen6226 {
								gen6223 := Call(__e, ShenFunc(symhd), V804)

								gen6224 := Call(__e, ShenFunc(symcons_2), gen6223)

								var gen6225 Obj
								if True == gen6224 {
									gen6219 := Call(__e, ShenFunc(symhd), V804)

									gen6220 := Call(__e, ShenFunc(symhd), gen6219)

									gen6221 := Call(__e, ShenFunc(sym_a), MakeSymbol("receive"), gen6220)

									var gen6222 Obj
									if True == gen6221 {
										gen6215 := Call(__e, ShenFunc(symhd), V804)

										gen6216 := Call(__e, ShenFunc(symtl), gen6215)

										gen6217 := Call(__e, ShenFunc(symcons_2), gen6216)

										var gen6218 Obj
										if True == gen6217 {
											gen6211 := Call(__e, ShenFunc(symhd), V804)

											gen6212 := Call(__e, ShenFunc(symtl), gen6211)

											gen6213 := Call(__e, ShenFunc(symtl), gen6212)

											gen6214 := Call(__e, ShenFunc(sym_a), Nil, gen6213)

											if True == gen6214 {
												gen6218 = True
											} else {
												gen6218 = False
											}

										} else {
											gen6218 = False
										}
										if True == gen6218 {
											gen6222 = True
										} else {
											gen6222 = False
										}

									} else {
										gen6222 = False
									}
									if True == gen6222 {
										gen6225 = True
									} else {
										gen6225 = False
									}

								} else {
									gen6225 = False
								}
								if True == gen6225 {
									gen6227 = True
								} else {
									gen6227 = False
								}

							} else {
								gen6227 = False
							}
							if True == gen6227 {
								gen6210 := Call(__e, ShenFunc(symtl), V804)

								__e.TailApply(ShenFunc(symshen_4bld_1prolog_1call), V803, gen6210)

								return

							} else {
								gen6208 := Call(__e, ShenFunc(symcons_2), V804)

								var gen6209 Obj
								if True == gen6208 {
									gen6205 := Call(__e, ShenFunc(symhd), V804)

									gen6206 := Call(__e, ShenFunc(symcons_2), gen6205)

									var gen6207 Obj
									if True == gen6206 {
										gen6201 := Call(__e, ShenFunc(symhd), V804)

										gen6202 := Call(__e, ShenFunc(symhd), gen6201)

										gen6203 := Call(__e, ShenFunc(sym_a), MakeSymbol("bind"), gen6202)

										var gen6204 Obj
										if True == gen6203 {
											gen6197 := Call(__e, ShenFunc(symhd), V804)

											gen6198 := Call(__e, ShenFunc(symtl), gen6197)

											gen6199 := Call(__e, ShenFunc(symcons_2), gen6198)

											var gen6200 Obj
											if True == gen6199 {
												gen6192 := Call(__e, ShenFunc(symhd), V804)

												gen6193 := Call(__e, ShenFunc(symtl), gen6192)

												gen6194 := Call(__e, ShenFunc(symtl), gen6193)

												gen6195 := Call(__e, ShenFunc(symcons_2), gen6194)

												var gen6196 Obj
												if True == gen6195 {
													gen6187 := Call(__e, ShenFunc(symhd), V804)

													gen6188 := Call(__e, ShenFunc(symtl), gen6187)

													gen6189 := Call(__e, ShenFunc(symtl), gen6188)

													gen6190 := Call(__e, ShenFunc(symtl), gen6189)

													gen6191 := Call(__e, ShenFunc(sym_a), Nil, gen6190)

													if True == gen6191 {
														gen6196 = True
													} else {
														gen6196 = False
													}

												} else {
													gen6196 = False
												}
												if True == gen6196 {
													gen6200 = True
												} else {
													gen6200 = False
												}

											} else {
												gen6200 = False
											}
											if True == gen6200 {
												gen6204 = True
											} else {
												gen6204 = False
											}

										} else {
											gen6204 = False
										}
										if True == gen6204 {
											gen6207 = True
										} else {
											gen6207 = False
										}

									} else {
										gen6207 = False
									}
									if True == gen6207 {
										gen6209 = True
									} else {
										gen6209 = False
									}

								} else {
									gen6209 = False
								}
								if True == gen6209 {
									gen6171 := Call(__e, ShenFunc(symhd), V804)

									gen6172 := Call(__e, ShenFunc(symtl), gen6171)

									gen6173 := Call(__e, ShenFunc(symhd), gen6172)

									gen6174 := Call(__e, ShenFunc(symhd), V804)

									gen6175 := Call(__e, ShenFunc(symtl), gen6174)

									gen6176 := Call(__e, ShenFunc(symtl), gen6175)

									gen6177 := Call(__e, ShenFunc(symhd), gen6176)

									gen6178 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen6177, V803)

									gen6179 := Call(__e, ShenFunc(symtl), V804)

									gen6180 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen6179)

									gen6181 := Call(__e, ShenFunc(symcons), gen6180, Nil)

									gen6182 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen6181)

									gen6183 := Call(__e, ShenFunc(symcons), gen6182, Nil)

									gen6184 := Call(__e, ShenFunc(symcons), V803, gen6183)

									gen6185 := Call(__e, ShenFunc(symcons), gen6178, gen6184)

									gen6186 := Call(__e, ShenFunc(symcons), gen6173, gen6185)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen6186)

									return

								} else {
									gen6169 := Call(__e, ShenFunc(symcons_2), V804)

									var gen6170 Obj
									if True == gen6169 {
										gen6166 := Call(__e, ShenFunc(symhd), V804)

										gen6167 := Call(__e, ShenFunc(symcons_2), gen6166)

										var gen6168 Obj
										if True == gen6167 {
											gen6162 := Call(__e, ShenFunc(symhd), V804)

											gen6163 := Call(__e, ShenFunc(symhd), gen6162)

											gen6164 := Call(__e, ShenFunc(sym_a), MakeSymbol("fwhen"), gen6163)

											var gen6165 Obj
											if True == gen6164 {
												gen6158 := Call(__e, ShenFunc(symhd), V804)

												gen6159 := Call(__e, ShenFunc(symtl), gen6158)

												gen6160 := Call(__e, ShenFunc(symcons_2), gen6159)

												var gen6161 Obj
												if True == gen6160 {
													gen6154 := Call(__e, ShenFunc(symhd), V804)

													gen6155 := Call(__e, ShenFunc(symtl), gen6154)

													gen6156 := Call(__e, ShenFunc(symtl), gen6155)

													gen6157 := Call(__e, ShenFunc(sym_a), Nil, gen6156)

													if True == gen6157 {
														gen6161 = True
													} else {
														gen6161 = False
													}

												} else {
													gen6161 = False
												}
												if True == gen6161 {
													gen6165 = True
												} else {
													gen6165 = False
												}

											} else {
												gen6165 = False
											}
											if True == gen6165 {
												gen6168 = True
											} else {
												gen6168 = False
											}

										} else {
											gen6168 = False
										}
										if True == gen6168 {
											gen6170 = True
										} else {
											gen6170 = False
										}

									} else {
										gen6170 = False
									}
									if True == gen6170 {
										gen6143 := Call(__e, ShenFunc(symhd), V804)

										gen6144 := Call(__e, ShenFunc(symtl), gen6143)

										gen6145 := Call(__e, ShenFunc(symhd), gen6144)

										gen6146 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen6145, V803)

										gen6147 := Call(__e, ShenFunc(symtl), V804)

										gen6148 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen6147)

										gen6149 := Call(__e, ShenFunc(symcons), gen6148, Nil)

										gen6150 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen6149)

										gen6151 := Call(__e, ShenFunc(symcons), gen6150, Nil)

										gen6152 := Call(__e, ShenFunc(symcons), V803, gen6151)

										gen6153 := Call(__e, ShenFunc(symcons), gen6146, gen6152)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen6153)

										return

									} else {
										gen6142 := Call(__e, ShenFunc(symcons_2), V804)

										if True == gen6142 {
											gen6135 := Call(__e, ShenFunc(symhd), V804)

											gen6136 := Call(__e, ShenFunc(symtl), V804)

											gen6137 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen6136)

											gen6138 := Call(__e, ShenFunc(symcons), gen6137, Nil)

											gen6139 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen6138)

											gen6140 := Call(__e, ShenFunc(symcons), gen6139, Nil)

											gen6141 := Call(__e, ShenFunc(symcons), V803, gen6140)

											__e.TailApply(ShenFunc(symappend), gen6135, gen6141)

											return

										} else {
											__e.TailApply(ShenFunc(symsimple_1error), MakeString("implementation error in bld-prolog-call"))

											return
										}

									}

								}

							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.bld-prolog-call"), gen6307)

		gen6320 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V806 := __args[0]
			_ = V806
			gen6318 := Call(__e, ShenFunc(symcons_2), V806)

			var gen6319 Obj
			if True == gen6318 {
				gen6315 := Call(__e, ShenFunc(symhd), V806)

				gen6316 := Call(__e, ShenFunc(sym_a), MakeSymbol("defprolog"), gen6315)

				var gen6317 Obj
				if True == gen6316 {
					gen6313 := Call(__e, ShenFunc(symtl), V806)

					gen6314 := Call(__e, ShenFunc(symcons_2), gen6313)

					if True == gen6314 {
						gen6317 = True
					} else {
						gen6317 = False
					}

				} else {
					gen6317 = False
				}
				if True == gen6317 {
					gen6319 = True
				} else {
					gen6319 = False
				}

			} else {
				gen6319 = False
			}
			if True == gen6319 {
				gen6308 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Y := __args[0]
					_ = Y
					__e.TailApply(ShenFunc(symshen_4_5defprolog_6), Y)

					return
				}, 1)
				gen6309 := Call(__e, ShenFunc(symtl), V806)

				gen6312 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Y := __args[0]
					_ = Y
					gen6310 := Call(__e, ShenFunc(symtl), V806)

					gen6311 := Call(__e, ShenFunc(symhd), gen6310)

					__e.TailApply(ShenFunc(symshen_4prolog_1error), gen6311, Y)

					return

				}, 1)
				__e.TailApply(ShenFunc(symcompile), gen6308, gen6309, gen6312)

				return

			} else {
				__e.Return(V806)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.defprolog-macro"), gen6320)

		gen6347 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V808 := __args[0]
			_ = V808
			gen6345 := Call(__e, ShenFunc(symcons_2), V808)

			var gen6346 Obj
			if True == gen6345 {
				gen6342 := Call(__e, ShenFunc(symhd), V808)

				gen6343 := Call(__e, ShenFunc(sym_a), MakeSymbol("datatype"), gen6342)

				var gen6344 Obj
				if True == gen6343 {
					gen6340 := Call(__e, ShenFunc(symtl), V808)

					gen6341 := Call(__e, ShenFunc(symcons_2), gen6340)

					if True == gen6341 {
						gen6344 = True
					} else {
						gen6344 = False
					}

				} else {
					gen6344 = False
				}
				if True == gen6344 {
					gen6346 = True
				} else {
					gen6346 = False
				}

			} else {
				gen6346 = False
			}
			if True == gen6346 {
				gen6321 := Call(__e, ShenFunc(symtl), V808)

				gen6322 := Call(__e, ShenFunc(symhd), gen6321)

				gen6323 := Call(__e, ShenFunc(symshen_4intern_1type), gen6322)

				gen6324 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), Nil)

				gen6325 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.<datatype-rules>"), gen6324)

				gen6326 := Call(__e, ShenFunc(symcons), gen6325, Nil)

				gen6327 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen6326)

				gen6328 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen6327)

				gen6329 := Call(__e, ShenFunc(symtl), V808)

				gen6330 := Call(__e, ShenFunc(symtl), gen6329)

				gen6331 := Call(__e, ShenFunc(symshen_4rcons__form), gen6330)

				gen6332 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.datatype-error"), Nil)

				gen6333 := Call(__e, ShenFunc(symcons), MakeSymbol("function"), gen6332)

				gen6334 := Call(__e, ShenFunc(symcons), gen6333, Nil)

				gen6335 := Call(__e, ShenFunc(symcons), gen6331, gen6334)

				gen6336 := Call(__e, ShenFunc(symcons), gen6328, gen6335)

				gen6337 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen6336)

				gen6338 := Call(__e, ShenFunc(symcons), gen6337, Nil)

				gen6339 := Call(__e, ShenFunc(symcons), gen6323, gen6338)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.process-datatype"), gen6339)

				return

			} else {
				__e.Return(V808)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.datatype-macro"), gen6347)

		gen6350 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V810 := __args[0]
			_ = V810
			gen6348 := Call(__e, ShenFunc(symstr), V810)

			gen6349 := Call(__e, ShenFunc(symcn), gen6348, MakeString("#type"))

			__e.TailApply(ShenFunc(symintern), gen6349)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intern-type"), gen6350)

		gen6404 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V812 := __args[0]
			_ = V812
			gen6402 := Call(__e, ShenFunc(symcons_2), V812)

			var gen6403 Obj
			if True == gen6402 {
				gen6399 := Call(__e, ShenFunc(symhd), V812)

				gen6400 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen6399)

				var gen6401 Obj
				if True == gen6400 {
					gen6396 := Call(__e, ShenFunc(symtl), V812)

					gen6397 := Call(__e, ShenFunc(symcons_2), gen6396)

					var gen6398 Obj
					if True == gen6397 {
						gen6392 := Call(__e, ShenFunc(symtl), V812)

						gen6393 := Call(__e, ShenFunc(symtl), gen6392)

						gen6394 := Call(__e, ShenFunc(symcons_2), gen6393)

						var gen6395 Obj
						if True == gen6394 {
							gen6388 := Call(__e, ShenFunc(symtl), V812)

							gen6389 := Call(__e, ShenFunc(symtl), gen6388)

							gen6390 := Call(__e, ShenFunc(symtl), gen6389)

							gen6391 := Call(__e, ShenFunc(symcons_2), gen6390)

							if True == gen6391 {
								gen6395 = True
							} else {
								gen6395 = False
							}

						} else {
							gen6395 = False
						}
						if True == gen6395 {
							gen6398 = True
						} else {
							gen6398 = False
						}

					} else {
						gen6398 = False
					}
					if True == gen6398 {
						gen6401 = True
					} else {
						gen6401 = False
					}

				} else {
					gen6401 = False
				}
				if True == gen6401 {
					gen6403 = True
				} else {
					gen6403 = False
				}

			} else {
				gen6403 = False
			}
			if True == gen6403 {
				gen6380 := Call(__e, ShenFunc(symtl), V812)

				gen6381 := Call(__e, ShenFunc(symhd), gen6380)

				gen6382 := Call(__e, ShenFunc(symtl), V812)

				gen6383 := Call(__e, ShenFunc(symtl), gen6382)

				gen6384 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen6383)

				gen6385 := Call(__e, ShenFunc(symshen_4_8s_1macro), gen6384)

				gen6386 := Call(__e, ShenFunc(symcons), gen6385, Nil)

				gen6387 := Call(__e, ShenFunc(symcons), gen6381, gen6386)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("@s"), gen6387)

				return

			} else {
				gen6378 := Call(__e, ShenFunc(symcons_2), V812)

				var gen6379 Obj
				if True == gen6378 {
					gen6375 := Call(__e, ShenFunc(symhd), V812)

					gen6376 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen6375)

					var gen6377 Obj
					if True == gen6376 {
						gen6372 := Call(__e, ShenFunc(symtl), V812)

						gen6373 := Call(__e, ShenFunc(symcons_2), gen6372)

						var gen6374 Obj
						if True == gen6373 {
							gen6368 := Call(__e, ShenFunc(symtl), V812)

							gen6369 := Call(__e, ShenFunc(symtl), gen6368)

							gen6370 := Call(__e, ShenFunc(symcons_2), gen6369)

							var gen6371 Obj
							if True == gen6370 {
								gen6363 := Call(__e, ShenFunc(symtl), V812)

								gen6364 := Call(__e, ShenFunc(symtl), gen6363)

								gen6365 := Call(__e, ShenFunc(symtl), gen6364)

								gen6366 := Call(__e, ShenFunc(sym_a), Nil, gen6365)

								var gen6367 Obj
								if True == gen6366 {
									gen6360 := Call(__e, ShenFunc(symtl), V812)

									gen6361 := Call(__e, ShenFunc(symhd), gen6360)

									gen6362 := Call(__e, ShenFunc(symstring_2), gen6361)

									if True == gen6362 {
										gen6367 = True
									} else {
										gen6367 = False
									}

								} else {
									gen6367 = False
								}
								if True == gen6367 {
									gen6371 = True
								} else {
									gen6371 = False
								}

							} else {
								gen6371 = False
							}
							if True == gen6371 {
								gen6374 = True
							} else {
								gen6374 = False
							}

						} else {
							gen6374 = False
						}
						if True == gen6374 {
							gen6377 = True
						} else {
							gen6377 = False
						}

					} else {
						gen6377 = False
					}
					if True == gen6377 {
						gen6379 = True
					} else {
						gen6379 = False
					}

				} else {
					gen6379 = False
				}
				if True == gen6379 {
					gen6351 := Call(__e, ShenFunc(symtl), V812)

					gen6352 := Call(__e, ShenFunc(symhd), gen6351)

					gen6353 := Call(__e, ShenFunc(symexplode), gen6352)

					E := gen6353
					gen6358 := Call(__e, ShenFunc(symlength), E)

					gen6359 := Call(__e, ShenFunc(sym_6), gen6358, MakeNumber(1))

					if True == gen6359 {
						gen6354 := Call(__e, ShenFunc(symtl), V812)

						gen6355 := Call(__e, ShenFunc(symtl), gen6354)

						gen6356 := Call(__e, ShenFunc(symappend), E, gen6355)

						gen6357 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen6356)

						__e.TailApply(ShenFunc(symshen_4_8s_1macro), gen6357)

						return

					} else {
						__e.Return(V812)
						return
					}

				} else {
					__e.Return(V812)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.@s-macro"), gen6404)

		gen6413 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V814 := __args[0]
			_ = V814
			gen6411 := Call(__e, ShenFunc(symcons_2), V814)

			var gen6412 Obj
			if True == gen6411 {
				gen6409 := Call(__e, ShenFunc(symhd), V814)

				gen6410 := Call(__e, ShenFunc(sym_a), MakeSymbol("synonyms"), gen6409)

				if True == gen6410 {
					gen6412 = True
				} else {
					gen6412 = False
				}

			} else {
				gen6412 = False
			}
			if True == gen6412 {
				gen6405 := Call(__e, ShenFunc(symtl), V814)

				gen6406 := Call(__e, ShenFunc(symshen_4curry_1synonyms), gen6405)

				gen6407 := Call(__e, ShenFunc(symshen_4rcons__form), gen6406)

				gen6408 := Call(__e, ShenFunc(symcons), gen6407, Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.synonyms-help"), gen6408)

				return

			} else {
				__e.Return(V814)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.synonyms-macro"), gen6413)

		gen6415 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V816 := __args[0]
			_ = V816
			gen6414 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4curry_1type), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symmap), gen6414, V816)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry-synonyms"), gen6415)

		gen6424 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V818 := __args[0]
			_ = V818
			gen6422 := Call(__e, ShenFunc(symcons_2), V818)

			var gen6423 Obj
			if True == gen6422 {
				gen6419 := Call(__e, ShenFunc(symhd), V818)

				gen6420 := Call(__e, ShenFunc(sym_a), MakeSymbol("nl"), gen6419)

				var gen6421 Obj
				if True == gen6420 {
					gen6417 := Call(__e, ShenFunc(symtl), V818)

					gen6418 := Call(__e, ShenFunc(sym_a), Nil, gen6417)

					if True == gen6418 {
						gen6421 = True
					} else {
						gen6421 = False
					}

				} else {
					gen6421 = False
				}
				if True == gen6421 {
					gen6423 = True
				} else {
					gen6423 = False
				}

			} else {
				gen6423 = False
			}
			if True == gen6423 {
				gen6416 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("nl"), gen6416)

				return

			} else {
				__e.Return(V818)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.nl-macro"), gen6424)

		gen6459 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V820 := __args[0]
			_ = V820
			gen6457 := Call(__e, ShenFunc(symcons_2), V820)

			var gen6458 Obj
			if True == gen6457 {
				gen6454 := Call(__e, ShenFunc(symtl), V820)

				gen6455 := Call(__e, ShenFunc(symcons_2), gen6454)

				var gen6456 Obj
				if True == gen6455 {
					gen6450 := Call(__e, ShenFunc(symtl), V820)

					gen6451 := Call(__e, ShenFunc(symtl), gen6450)

					gen6452 := Call(__e, ShenFunc(symcons_2), gen6451)

					var gen6453 Obj
					if True == gen6452 {
						gen6445 := Call(__e, ShenFunc(symtl), V820)

						gen6446 := Call(__e, ShenFunc(symtl), gen6445)

						gen6447 := Call(__e, ShenFunc(symtl), gen6446)

						gen6448 := Call(__e, ShenFunc(symcons_2), gen6447)

						var gen6449 Obj
						if True == gen6448 {
							gen6435 := Call(__e, ShenFunc(symhd), V820)

							gen6436 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), Nil)

							gen6437 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen6436)

							gen6438 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen6437)

							gen6439 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen6438)

							gen6440 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen6439)

							gen6441 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen6440)

							gen6442 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen6441)

							gen6443 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen6442)

							gen6444 := Call(__e, ShenFunc(symelement_2), gen6435, gen6443)

							if True == gen6444 {
								gen6449 = True
							} else {
								gen6449 = False
							}

						} else {
							gen6449 = False
						}
						if True == gen6449 {
							gen6453 = True
						} else {
							gen6453 = False
						}

					} else {
						gen6453 = False
					}
					if True == gen6453 {
						gen6456 = True
					} else {
						gen6456 = False
					}

				} else {
					gen6456 = False
				}
				if True == gen6456 {
					gen6458 = True
				} else {
					gen6458 = False
				}

			} else {
				gen6458 = False
			}
			if True == gen6458 {
				gen6425 := Call(__e, ShenFunc(symhd), V820)

				gen6426 := Call(__e, ShenFunc(symtl), V820)

				gen6427 := Call(__e, ShenFunc(symhd), gen6426)

				gen6428 := Call(__e, ShenFunc(symhd), V820)

				gen6429 := Call(__e, ShenFunc(symtl), V820)

				gen6430 := Call(__e, ShenFunc(symtl), gen6429)

				gen6431 := Call(__e, ShenFunc(symcons), gen6428, gen6430)

				gen6432 := Call(__e, ShenFunc(symshen_4assoc_1macro), gen6431)

				gen6433 := Call(__e, ShenFunc(symcons), gen6432, Nil)

				gen6434 := Call(__e, ShenFunc(symcons), gen6427, gen6433)

				__e.TailApply(ShenFunc(symcons), gen6425, gen6434)

				return

			} else {
				__e.Return(V820)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assoc-macro"), gen6459)

		gen6495 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V822 := __args[0]
			_ = V822
			gen6493 := Call(__e, ShenFunc(symcons_2), V822)

			var gen6494 Obj
			if True == gen6493 {
				gen6490 := Call(__e, ShenFunc(symhd), V822)

				gen6491 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen6490)

				var gen6492 Obj
				if True == gen6491 {
					gen6487 := Call(__e, ShenFunc(symtl), V822)

					gen6488 := Call(__e, ShenFunc(symcons_2), gen6487)

					var gen6489 Obj
					if True == gen6488 {
						gen6483 := Call(__e, ShenFunc(symtl), V822)

						gen6484 := Call(__e, ShenFunc(symtl), gen6483)

						gen6485 := Call(__e, ShenFunc(symcons_2), gen6484)

						var gen6486 Obj
						if True == gen6485 {
							gen6478 := Call(__e, ShenFunc(symtl), V822)

							gen6479 := Call(__e, ShenFunc(symtl), gen6478)

							gen6480 := Call(__e, ShenFunc(symtl), gen6479)

							gen6481 := Call(__e, ShenFunc(symcons_2), gen6480)

							var gen6482 Obj
							if True == gen6481 {
								gen6473 := Call(__e, ShenFunc(symtl), V822)

								gen6474 := Call(__e, ShenFunc(symtl), gen6473)

								gen6475 := Call(__e, ShenFunc(symtl), gen6474)

								gen6476 := Call(__e, ShenFunc(symtl), gen6475)

								gen6477 := Call(__e, ShenFunc(symcons_2), gen6476)

								if True == gen6477 {
									gen6482 = True
								} else {
									gen6482 = False
								}

							} else {
								gen6482 = False
							}
							if True == gen6482 {
								gen6486 = True
							} else {
								gen6486 = False
							}

						} else {
							gen6486 = False
						}
						if True == gen6486 {
							gen6489 = True
						} else {
							gen6489 = False
						}

					} else {
						gen6489 = False
					}
					if True == gen6489 {
						gen6492 = True
					} else {
						gen6492 = False
					}

				} else {
					gen6492 = False
				}
				if True == gen6492 {
					gen6494 = True
				} else {
					gen6494 = False
				}

			} else {
				gen6494 = False
			}
			if True == gen6494 {
				gen6460 := Call(__e, ShenFunc(symtl), V822)

				gen6461 := Call(__e, ShenFunc(symhd), gen6460)

				gen6462 := Call(__e, ShenFunc(symtl), V822)

				gen6463 := Call(__e, ShenFunc(symtl), gen6462)

				gen6464 := Call(__e, ShenFunc(symhd), gen6463)

				gen6465 := Call(__e, ShenFunc(symtl), V822)

				gen6466 := Call(__e, ShenFunc(symtl), gen6465)

				gen6467 := Call(__e, ShenFunc(symtl), gen6466)

				gen6468 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen6467)

				gen6469 := Call(__e, ShenFunc(symshen_4let_1macro), gen6468)

				gen6470 := Call(__e, ShenFunc(symcons), gen6469, Nil)

				gen6471 := Call(__e, ShenFunc(symcons), gen6464, gen6470)

				gen6472 := Call(__e, ShenFunc(symcons), gen6461, gen6471)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen6472)

				return

			} else {
				__e.Return(V822)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.let-macro"), gen6495)

		gen6537 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V824 := __args[0]
			_ = V824
			gen6535 := Call(__e, ShenFunc(symcons_2), V824)

			var gen6536 Obj
			if True == gen6535 {
				gen6532 := Call(__e, ShenFunc(symhd), V824)

				gen6533 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen6532)

				var gen6534 Obj
				if True == gen6533 {
					gen6529 := Call(__e, ShenFunc(symtl), V824)

					gen6530 := Call(__e, ShenFunc(symcons_2), gen6529)

					var gen6531 Obj
					if True == gen6530 {
						gen6525 := Call(__e, ShenFunc(symtl), V824)

						gen6526 := Call(__e, ShenFunc(symtl), gen6525)

						gen6527 := Call(__e, ShenFunc(symcons_2), gen6526)

						var gen6528 Obj
						if True == gen6527 {
							gen6521 := Call(__e, ShenFunc(symtl), V824)

							gen6522 := Call(__e, ShenFunc(symtl), gen6521)

							gen6523 := Call(__e, ShenFunc(symtl), gen6522)

							gen6524 := Call(__e, ShenFunc(symcons_2), gen6523)

							if True == gen6524 {
								gen6528 = True
							} else {
								gen6528 = False
							}

						} else {
							gen6528 = False
						}
						if True == gen6528 {
							gen6531 = True
						} else {
							gen6531 = False
						}

					} else {
						gen6531 = False
					}
					if True == gen6531 {
						gen6534 = True
					} else {
						gen6534 = False
					}

				} else {
					gen6534 = False
				}
				if True == gen6534 {
					gen6536 = True
				} else {
					gen6536 = False
				}

			} else {
				gen6536 = False
			}
			if True == gen6536 {
				gen6513 := Call(__e, ShenFunc(symtl), V824)

				gen6514 := Call(__e, ShenFunc(symhd), gen6513)

				gen6515 := Call(__e, ShenFunc(symtl), V824)

				gen6516 := Call(__e, ShenFunc(symtl), gen6515)

				gen6517 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen6516)

				gen6518 := Call(__e, ShenFunc(symshen_4abs_1macro), gen6517)

				gen6519 := Call(__e, ShenFunc(symcons), gen6518, Nil)

				gen6520 := Call(__e, ShenFunc(symcons), gen6514, gen6519)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen6520)

				return

			} else {
				gen6511 := Call(__e, ShenFunc(symcons_2), V824)

				var gen6512 Obj
				if True == gen6511 {
					gen6508 := Call(__e, ShenFunc(symhd), V824)

					gen6509 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen6508)

					var gen6510 Obj
					if True == gen6509 {
						gen6505 := Call(__e, ShenFunc(symtl), V824)

						gen6506 := Call(__e, ShenFunc(symcons_2), gen6505)

						var gen6507 Obj
						if True == gen6506 {
							gen6501 := Call(__e, ShenFunc(symtl), V824)

							gen6502 := Call(__e, ShenFunc(symtl), gen6501)

							gen6503 := Call(__e, ShenFunc(symcons_2), gen6502)

							var gen6504 Obj
							if True == gen6503 {
								gen6497 := Call(__e, ShenFunc(symtl), V824)

								gen6498 := Call(__e, ShenFunc(symtl), gen6497)

								gen6499 := Call(__e, ShenFunc(symtl), gen6498)

								gen6500 := Call(__e, ShenFunc(sym_a), Nil, gen6499)

								if True == gen6500 {
									gen6504 = True
								} else {
									gen6504 = False
								}

							} else {
								gen6504 = False
							}
							if True == gen6504 {
								gen6507 = True
							} else {
								gen6507 = False
							}

						} else {
							gen6507 = False
						}
						if True == gen6507 {
							gen6510 = True
						} else {
							gen6510 = False
						}

					} else {
						gen6510 = False
					}
					if True == gen6510 {
						gen6512 = True
					} else {
						gen6512 = False
					}

				} else {
					gen6512 = False
				}
				if True == gen6512 {
					gen6496 := Call(__e, ShenFunc(symtl), V824)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen6496)

					return

				} else {
					__e.Return(V824)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abs-macro"), gen6537)

		gen6616 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V828 := __args[0]
			_ = V828
			gen6614 := Call(__e, ShenFunc(symcons_2), V828)

			var gen6615 Obj
			if True == gen6614 {
				gen6611 := Call(__e, ShenFunc(symhd), V828)

				gen6612 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen6611)

				var gen6613 Obj
				if True == gen6612 {
					gen6608 := Call(__e, ShenFunc(symtl), V828)

					gen6609 := Call(__e, ShenFunc(symcons_2), gen6608)

					var gen6610 Obj
					if True == gen6609 {
						gen6604 := Call(__e, ShenFunc(symtl), V828)

						gen6605 := Call(__e, ShenFunc(symhd), gen6604)

						gen6606 := Call(__e, ShenFunc(sym_a), True, gen6605)

						var gen6607 Obj
						if True == gen6606 {
							gen6601 := Call(__e, ShenFunc(symtl), V828)

							gen6602 := Call(__e, ShenFunc(symtl), gen6601)

							gen6603 := Call(__e, ShenFunc(symcons_2), gen6602)

							if True == gen6603 {
								gen6607 = True
							} else {
								gen6607 = False
							}

						} else {
							gen6607 = False
						}
						if True == gen6607 {
							gen6610 = True
						} else {
							gen6610 = False
						}

					} else {
						gen6610 = False
					}
					if True == gen6610 {
						gen6613 = True
					} else {
						gen6613 = False
					}

				} else {
					gen6613 = False
				}
				if True == gen6613 {
					gen6615 = True
				} else {
					gen6615 = False
				}

			} else {
				gen6615 = False
			}
			if True == gen6615 {
				gen6599 := Call(__e, ShenFunc(symtl), V828)

				gen6600 := Call(__e, ShenFunc(symtl), gen6599)

				__e.TailApply(ShenFunc(symhd), gen6600)

				return

			} else {
				gen6597 := Call(__e, ShenFunc(symcons_2), V828)

				var gen6598 Obj
				if True == gen6597 {
					gen6594 := Call(__e, ShenFunc(symhd), V828)

					gen6595 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen6594)

					var gen6596 Obj
					if True == gen6595 {
						gen6591 := Call(__e, ShenFunc(symtl), V828)

						gen6592 := Call(__e, ShenFunc(symcons_2), gen6591)

						var gen6593 Obj
						if True == gen6592 {
							gen6587 := Call(__e, ShenFunc(symtl), V828)

							gen6588 := Call(__e, ShenFunc(symtl), gen6587)

							gen6589 := Call(__e, ShenFunc(symcons_2), gen6588)

							var gen6590 Obj
							if True == gen6589 {
								gen6583 := Call(__e, ShenFunc(symtl), V828)

								gen6584 := Call(__e, ShenFunc(symtl), gen6583)

								gen6585 := Call(__e, ShenFunc(symtl), gen6584)

								gen6586 := Call(__e, ShenFunc(sym_a), Nil, gen6585)

								if True == gen6586 {
									gen6590 = True
								} else {
									gen6590 = False
								}

							} else {
								gen6590 = False
							}
							if True == gen6590 {
								gen6593 = True
							} else {
								gen6593 = False
							}

						} else {
							gen6593 = False
						}
						if True == gen6593 {
							gen6596 = True
						} else {
							gen6596 = False
						}

					} else {
						gen6596 = False
					}
					if True == gen6596 {
						gen6598 = True
					} else {
						gen6598 = False
					}

				} else {
					gen6598 = False
				}
				if True == gen6598 {
					gen6573 := Call(__e, ShenFunc(symtl), V828)

					gen6574 := Call(__e, ShenFunc(symhd), gen6573)

					gen6575 := Call(__e, ShenFunc(symtl), V828)

					gen6576 := Call(__e, ShenFunc(symtl), gen6575)

					gen6577 := Call(__e, ShenFunc(symhd), gen6576)

					gen6578 := Call(__e, ShenFunc(symcons), MakeString("error: cases exhausted"), Nil)

					gen6579 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen6578)

					gen6580 := Call(__e, ShenFunc(symcons), gen6579, Nil)

					gen6581 := Call(__e, ShenFunc(symcons), gen6577, gen6580)

					gen6582 := Call(__e, ShenFunc(symcons), gen6574, gen6581)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen6582)

					return

				} else {
					gen6571 := Call(__e, ShenFunc(symcons_2), V828)

					var gen6572 Obj
					if True == gen6571 {
						gen6568 := Call(__e, ShenFunc(symhd), V828)

						gen6569 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen6568)

						var gen6570 Obj
						if True == gen6569 {
							gen6565 := Call(__e, ShenFunc(symtl), V828)

							gen6566 := Call(__e, ShenFunc(symcons_2), gen6565)

							var gen6567 Obj
							if True == gen6566 {
								gen6562 := Call(__e, ShenFunc(symtl), V828)

								gen6563 := Call(__e, ShenFunc(symtl), gen6562)

								gen6564 := Call(__e, ShenFunc(symcons_2), gen6563)

								if True == gen6564 {
									gen6567 = True
								} else {
									gen6567 = False
								}

							} else {
								gen6567 = False
							}
							if True == gen6567 {
								gen6570 = True
							} else {
								gen6570 = False
							}

						} else {
							gen6570 = False
						}
						if True == gen6570 {
							gen6572 = True
						} else {
							gen6572 = False
						}

					} else {
						gen6572 = False
					}
					if True == gen6572 {
						gen6549 := Call(__e, ShenFunc(symtl), V828)

						gen6550 := Call(__e, ShenFunc(symhd), gen6549)

						gen6551 := Call(__e, ShenFunc(symtl), V828)

						gen6552 := Call(__e, ShenFunc(symtl), gen6551)

						gen6553 := Call(__e, ShenFunc(symhd), gen6552)

						gen6554 := Call(__e, ShenFunc(symtl), V828)

						gen6555 := Call(__e, ShenFunc(symtl), gen6554)

						gen6556 := Call(__e, ShenFunc(symtl), gen6555)

						gen6557 := Call(__e, ShenFunc(symcons), MakeSymbol("cases"), gen6556)

						gen6558 := Call(__e, ShenFunc(symshen_4cases_1macro), gen6557)

						gen6559 := Call(__e, ShenFunc(symcons), gen6558, Nil)

						gen6560 := Call(__e, ShenFunc(symcons), gen6553, gen6559)

						gen6561 := Call(__e, ShenFunc(symcons), gen6550, gen6560)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen6561)

						return

					} else {
						gen6547 := Call(__e, ShenFunc(symcons_2), V828)

						var gen6548 Obj
						if True == gen6547 {
							gen6544 := Call(__e, ShenFunc(symhd), V828)

							gen6545 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen6544)

							var gen6546 Obj
							if True == gen6545 {
								gen6541 := Call(__e, ShenFunc(symtl), V828)

								gen6542 := Call(__e, ShenFunc(symcons_2), gen6541)

								var gen6543 Obj
								if True == gen6542 {
									gen6538 := Call(__e, ShenFunc(symtl), V828)

									gen6539 := Call(__e, ShenFunc(symtl), gen6538)

									gen6540 := Call(__e, ShenFunc(sym_a), Nil, gen6539)

									if True == gen6540 {
										gen6543 = True
									} else {
										gen6543 = False
									}

								} else {
									gen6543 = False
								}
								if True == gen6543 {
									gen6546 = True
								} else {
									gen6546 = False
								}

							} else {
								gen6546 = False
							}
							if True == gen6546 {
								gen6548 = True
							} else {
								gen6548 = False
							}

						} else {
							gen6548 = False
						}
						if True == gen6548 {
							__e.TailApply(ShenFunc(symsimple_1error), MakeString("error: odd number of case elements\n"))

							return
						} else {
							__e.Return(V828)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cases-macro"), gen6616)

		gen6661 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V830 := __args[0]
			_ = V830
			gen6659 := Call(__e, ShenFunc(symcons_2), V830)

			var gen6660 Obj
			if True == gen6659 {
				gen6656 := Call(__e, ShenFunc(symhd), V830)

				gen6657 := Call(__e, ShenFunc(sym_a), MakeSymbol("time"), gen6656)

				var gen6658 Obj
				if True == gen6657 {
					gen6653 := Call(__e, ShenFunc(symtl), V830)

					gen6654 := Call(__e, ShenFunc(symcons_2), gen6653)

					var gen6655 Obj
					if True == gen6654 {
						gen6650 := Call(__e, ShenFunc(symtl), V830)

						gen6651 := Call(__e, ShenFunc(symtl), gen6650)

						gen6652 := Call(__e, ShenFunc(sym_a), Nil, gen6651)

						if True == gen6652 {
							gen6655 = True
						} else {
							gen6655 = False
						}

					} else {
						gen6655 = False
					}
					if True == gen6655 {
						gen6658 = True
					} else {
						gen6658 = False
					}

				} else {
					gen6658 = False
				}
				if True == gen6658 {
					gen6660 = True
				} else {
					gen6660 = False
				}

			} else {
				gen6660 = False
			}
			if True == gen6660 {
				gen6617 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

				gen6618 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen6617)

				gen6619 := Call(__e, ShenFunc(symtl), V830)

				gen6620 := Call(__e, ShenFunc(symhd), gen6619)

				gen6621 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

				gen6622 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen6621)

				gen6623 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), Nil)

				gen6624 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), gen6623)

				gen6625 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen6624)

				gen6626 := Call(__e, ShenFunc(symcons), MakeSymbol("Time"), Nil)

				gen6627 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen6626)

				gen6628 := Call(__e, ShenFunc(symcons), MakeString(" secs\n"), Nil)

				gen6629 := Call(__e, ShenFunc(symcons), gen6627, gen6628)

				gen6630 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen6629)

				gen6631 := Call(__e, ShenFunc(symcons), gen6630, Nil)

				gen6632 := Call(__e, ShenFunc(symcons), MakeString("\nrun time: "), gen6631)

				gen6633 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen6632)

				gen6634 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), Nil)

				gen6635 := Call(__e, ShenFunc(symcons), gen6634, Nil)

				gen6636 := Call(__e, ShenFunc(symcons), gen6633, gen6635)

				gen6637 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.prhush"), gen6636)

				gen6638 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

				gen6639 := Call(__e, ShenFunc(symcons), gen6637, gen6638)

				gen6640 := Call(__e, ShenFunc(symcons), MakeSymbol("Message"), gen6639)

				gen6641 := Call(__e, ShenFunc(symcons), gen6625, gen6640)

				gen6642 := Call(__e, ShenFunc(symcons), MakeSymbol("Time"), gen6641)

				gen6643 := Call(__e, ShenFunc(symcons), gen6622, gen6642)

				gen6644 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), gen6643)

				gen6645 := Call(__e, ShenFunc(symcons), gen6620, gen6644)

				gen6646 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen6645)

				gen6647 := Call(__e, ShenFunc(symcons), gen6618, gen6646)

				gen6648 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), gen6647)

				gen6649 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen6648)

				__e.TailApply(ShenFunc(symshen_4let_1macro), gen6649)

				return

			} else {
				__e.Return(V830)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.timer-macro"), gen6661)

		gen6668 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V832 := __args[0]
			_ = V832
			gen6667 := Call(__e, ShenFunc(symcons_2), V832)

			if True == gen6667 {
				gen6662 := Call(__e, ShenFunc(symhd), V832)

				gen6663 := Call(__e, ShenFunc(symtl), V832)

				gen6664 := Call(__e, ShenFunc(symshen_4tuple_1up), gen6663)

				gen6665 := Call(__e, ShenFunc(symcons), gen6664, Nil)

				gen6666 := Call(__e, ShenFunc(symcons), gen6662, gen6665)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("@p"), gen6666)

				return

			} else {
				__e.Return(V832)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tuple-up"), gen6668)

		gen6758 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V834 := __args[0]
			_ = V834
			gen6756 := Call(__e, ShenFunc(symcons_2), V834)

			var gen6757 Obj
			if True == gen6756 {
				gen6753 := Call(__e, ShenFunc(symhd), V834)

				gen6754 := Call(__e, ShenFunc(sym_a), MakeSymbol("put"), gen6753)

				var gen6755 Obj
				if True == gen6754 {
					gen6750 := Call(__e, ShenFunc(symtl), V834)

					gen6751 := Call(__e, ShenFunc(symcons_2), gen6750)

					var gen6752 Obj
					if True == gen6751 {
						gen6746 := Call(__e, ShenFunc(symtl), V834)

						gen6747 := Call(__e, ShenFunc(symtl), gen6746)

						gen6748 := Call(__e, ShenFunc(symcons_2), gen6747)

						var gen6749 Obj
						if True == gen6748 {
							gen6741 := Call(__e, ShenFunc(symtl), V834)

							gen6742 := Call(__e, ShenFunc(symtl), gen6741)

							gen6743 := Call(__e, ShenFunc(symtl), gen6742)

							gen6744 := Call(__e, ShenFunc(symcons_2), gen6743)

							var gen6745 Obj
							if True == gen6744 {
								gen6736 := Call(__e, ShenFunc(symtl), V834)

								gen6737 := Call(__e, ShenFunc(symtl), gen6736)

								gen6738 := Call(__e, ShenFunc(symtl), gen6737)

								gen6739 := Call(__e, ShenFunc(symtl), gen6738)

								gen6740 := Call(__e, ShenFunc(sym_a), Nil, gen6739)

								if True == gen6740 {
									gen6745 = True
								} else {
									gen6745 = False
								}

							} else {
								gen6745 = False
							}
							if True == gen6745 {
								gen6749 = True
							} else {
								gen6749 = False
							}

						} else {
							gen6749 = False
						}
						if True == gen6749 {
							gen6752 = True
						} else {
							gen6752 = False
						}

					} else {
						gen6752 = False
					}
					if True == gen6752 {
						gen6755 = True
					} else {
						gen6755 = False
					}

				} else {
					gen6755 = False
				}
				if True == gen6755 {
					gen6757 = True
				} else {
					gen6757 = False
				}

			} else {
				gen6757 = False
			}
			if True == gen6757 {
				gen6721 := Call(__e, ShenFunc(symtl), V834)

				gen6722 := Call(__e, ShenFunc(symhd), gen6721)

				gen6723 := Call(__e, ShenFunc(symtl), V834)

				gen6724 := Call(__e, ShenFunc(symtl), gen6723)

				gen6725 := Call(__e, ShenFunc(symhd), gen6724)

				gen6726 := Call(__e, ShenFunc(symtl), V834)

				gen6727 := Call(__e, ShenFunc(symtl), gen6726)

				gen6728 := Call(__e, ShenFunc(symtl), gen6727)

				gen6729 := Call(__e, ShenFunc(symhd), gen6728)

				gen6730 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), Nil)

				gen6731 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen6730)

				gen6732 := Call(__e, ShenFunc(symcons), gen6731, Nil)

				gen6733 := Call(__e, ShenFunc(symcons), gen6729, gen6732)

				gen6734 := Call(__e, ShenFunc(symcons), gen6725, gen6733)

				gen6735 := Call(__e, ShenFunc(symcons), gen6722, gen6734)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("put"), gen6735)

				return

			} else {
				gen6719 := Call(__e, ShenFunc(symcons_2), V834)

				var gen6720 Obj
				if True == gen6719 {
					gen6716 := Call(__e, ShenFunc(symhd), V834)

					gen6717 := Call(__e, ShenFunc(sym_a), MakeSymbol("get"), gen6716)

					var gen6718 Obj
					if True == gen6717 {
						gen6713 := Call(__e, ShenFunc(symtl), V834)

						gen6714 := Call(__e, ShenFunc(symcons_2), gen6713)

						var gen6715 Obj
						if True == gen6714 {
							gen6709 := Call(__e, ShenFunc(symtl), V834)

							gen6710 := Call(__e, ShenFunc(symtl), gen6709)

							gen6711 := Call(__e, ShenFunc(symcons_2), gen6710)

							var gen6712 Obj
							if True == gen6711 {
								gen6705 := Call(__e, ShenFunc(symtl), V834)

								gen6706 := Call(__e, ShenFunc(symtl), gen6705)

								gen6707 := Call(__e, ShenFunc(symtl), gen6706)

								gen6708 := Call(__e, ShenFunc(sym_a), Nil, gen6707)

								if True == gen6708 {
									gen6712 = True
								} else {
									gen6712 = False
								}

							} else {
								gen6712 = False
							}
							if True == gen6712 {
								gen6715 = True
							} else {
								gen6715 = False
							}

						} else {
							gen6715 = False
						}
						if True == gen6715 {
							gen6718 = True
						} else {
							gen6718 = False
						}

					} else {
						gen6718 = False
					}
					if True == gen6718 {
						gen6720 = True
					} else {
						gen6720 = False
					}

				} else {
					gen6720 = False
				}
				if True == gen6720 {
					gen6695 := Call(__e, ShenFunc(symtl), V834)

					gen6696 := Call(__e, ShenFunc(symhd), gen6695)

					gen6697 := Call(__e, ShenFunc(symtl), V834)

					gen6698 := Call(__e, ShenFunc(symtl), gen6697)

					gen6699 := Call(__e, ShenFunc(symhd), gen6698)

					gen6700 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), Nil)

					gen6701 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen6700)

					gen6702 := Call(__e, ShenFunc(symcons), gen6701, Nil)

					gen6703 := Call(__e, ShenFunc(symcons), gen6699, gen6702)

					gen6704 := Call(__e, ShenFunc(symcons), gen6696, gen6703)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("get"), gen6704)

					return

				} else {
					gen6693 := Call(__e, ShenFunc(symcons_2), V834)

					var gen6694 Obj
					if True == gen6693 {
						gen6690 := Call(__e, ShenFunc(symhd), V834)

						gen6691 := Call(__e, ShenFunc(sym_a), MakeSymbol("unput"), gen6690)

						var gen6692 Obj
						if True == gen6691 {
							gen6687 := Call(__e, ShenFunc(symtl), V834)

							gen6688 := Call(__e, ShenFunc(symcons_2), gen6687)

							var gen6689 Obj
							if True == gen6688 {
								gen6683 := Call(__e, ShenFunc(symtl), V834)

								gen6684 := Call(__e, ShenFunc(symtl), gen6683)

								gen6685 := Call(__e, ShenFunc(symcons_2), gen6684)

								var gen6686 Obj
								if True == gen6685 {
									gen6679 := Call(__e, ShenFunc(symtl), V834)

									gen6680 := Call(__e, ShenFunc(symtl), gen6679)

									gen6681 := Call(__e, ShenFunc(symtl), gen6680)

									gen6682 := Call(__e, ShenFunc(sym_a), Nil, gen6681)

									if True == gen6682 {
										gen6686 = True
									} else {
										gen6686 = False
									}

								} else {
									gen6686 = False
								}
								if True == gen6686 {
									gen6689 = True
								} else {
									gen6689 = False
								}

							} else {
								gen6689 = False
							}
							if True == gen6689 {
								gen6692 = True
							} else {
								gen6692 = False
							}

						} else {
							gen6692 = False
						}
						if True == gen6692 {
							gen6694 = True
						} else {
							gen6694 = False
						}

					} else {
						gen6694 = False
					}
					if True == gen6694 {
						gen6669 := Call(__e, ShenFunc(symtl), V834)

						gen6670 := Call(__e, ShenFunc(symhd), gen6669)

						gen6671 := Call(__e, ShenFunc(symtl), V834)

						gen6672 := Call(__e, ShenFunc(symtl), gen6671)

						gen6673 := Call(__e, ShenFunc(symhd), gen6672)

						gen6674 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), Nil)

						gen6675 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen6674)

						gen6676 := Call(__e, ShenFunc(symcons), gen6675, Nil)

						gen6677 := Call(__e, ShenFunc(symcons), gen6673, gen6676)

						gen6678 := Call(__e, ShenFunc(symcons), gen6670, gen6677)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("unput"), gen6678)

						return

					} else {
						__e.Return(V834)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.put/get-macro"), gen6758)

		gen6775 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V836 := __args[0]
			_ = V836
			gen6773 := Call(__e, ShenFunc(symcons_2), V836)

			var gen6774 Obj
			if True == gen6773 {
				gen6770 := Call(__e, ShenFunc(symhd), V836)

				gen6771 := Call(__e, ShenFunc(sym_a), MakeSymbol("function"), gen6770)

				var gen6772 Obj
				if True == gen6771 {
					gen6767 := Call(__e, ShenFunc(symtl), V836)

					gen6768 := Call(__e, ShenFunc(symcons_2), gen6767)

					var gen6769 Obj
					if True == gen6768 {
						gen6764 := Call(__e, ShenFunc(symtl), V836)

						gen6765 := Call(__e, ShenFunc(symtl), gen6764)

						gen6766 := Call(__e, ShenFunc(sym_a), Nil, gen6765)

						if True == gen6766 {
							gen6769 = True
						} else {
							gen6769 = False
						}

					} else {
						gen6769 = False
					}
					if True == gen6769 {
						gen6772 = True
					} else {
						gen6772 = False
					}

				} else {
					gen6772 = False
				}
				if True == gen6772 {
					gen6774 = True
				} else {
					gen6774 = False
				}

			} else {
				gen6774 = False
			}
			if True == gen6774 {
				gen6759 := Call(__e, ShenFunc(symtl), V836)

				gen6760 := Call(__e, ShenFunc(symhd), gen6759)

				gen6761 := Call(__e, ShenFunc(symtl), V836)

				gen6762 := Call(__e, ShenFunc(symhd), gen6761)

				gen6763 := Call(__e, ShenFunc(symarity), gen6762)

				__e.TailApply(ShenFunc(symshen_4function_1abstraction), gen6760, gen6763)

				return

			} else {
				__e.Return(V836)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.function-macro"), gen6775)

		gen6780 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V839 := __args[0]
			_ = V839
			V840 := __args[1]
			_ = V840
			gen6779 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V840)

			if True == gen6779 {
				gen6778 := Call(__e, ShenFunc(symshen_4app), V839, MakeString(" has no lambda form\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen6778)

				return

			} else {
				gen6777 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V840)

				if True == gen6777 {
					gen6776 := Call(__e, ShenFunc(symcons), V839, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("function"), gen6776)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4function_1abstraction_1help), V839, V840, Nil)

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.function-abstraction"), gen6780)

		gen6789 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V844 := __args[0]
			_ = V844
			V845 := __args[1]
			_ = V845
			V846 := __args[2]
			_ = V846
			gen6788 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V845)

			if True == gen6788 {
				__e.TailApply(ShenFunc(symcons), V844, V846)

				return
			} else {
				gen6781 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

				X := gen6781
				gen6782 := Call(__e, ShenFunc(sym_1), V845, MakeNumber(1))

				gen6783 := Call(__e, ShenFunc(symcons), X, Nil)

				gen6784 := Call(__e, ShenFunc(symappend), V846, gen6783)

				gen6785 := Call(__e, ShenFunc(symshen_4function_1abstraction_1help), V844, gen6782, gen6784)

				gen6786 := Call(__e, ShenFunc(symcons), gen6785, Nil)

				gen6787 := Call(__e, ShenFunc(symcons), X, gen6786)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("/."), gen6787)

				return

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.function-abstraction-help"), gen6789)

		gen6797 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V848 := __args[0]
			_ = V848
			gen6790 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*macroreg*"))

			MacroReg := gen6790
			gen6791 := Call(__e, ShenFunc(symshen_4findpos), V848, MacroReg)

			Pos := gen6791
			gen6792 := Call(__e, ShenFunc(symremove), V848, MacroReg)

			gen6793 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*macroreg*"), gen6792)

			Remove1 := gen6793
			_ = Remove1
			gen6794 := Call(__e, ShenFunc(symvalue), MakeSymbol("*macros*"))

			gen6795 := Call(__e, ShenFunc(symshen_4remove_1nth), Pos, gen6794)

			gen6796 := Call(__e, ShenFunc(symset), MakeSymbol("*macros*"), gen6795)

			Remove2 := gen6796
			_ = Remove2
			__e.Return(V848)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("undefmacro"), gen6797)

		gen6807 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V858 := __args[0]
			_ = V858
			V859 := __args[1]
			_ = V859
			gen6806 := Call(__e, ShenFunc(sym_a), Nil, V859)

			if True == gen6806 {
				gen6805 := Call(__e, ShenFunc(symshen_4app), V858, MakeString(" is not a macro\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen6805)

				return

			} else {
				gen6803 := Call(__e, ShenFunc(symcons_2), V859)

				var gen6804 Obj
				if True == gen6803 {
					gen6801 := Call(__e, ShenFunc(symhd), V859)

					gen6802 := Call(__e, ShenFunc(sym_a), gen6801, V858)

					if True == gen6802 {
						gen6804 = True
					} else {
						gen6804 = False
					}

				} else {
					gen6804 = False
				}
				if True == gen6804 {
					__e.Return(MakeNumber(1))
					return
				} else {
					gen6800 := Call(__e, ShenFunc(symcons_2), V859)

					if True == gen6800 {
						gen6798 := Call(__e, ShenFunc(symtl), V859)

						gen6799 := Call(__e, ShenFunc(symshen_4findpos), V858, gen6798)

						__e.TailApply(ShenFunc(sym_7), MakeNumber(1), gen6799)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.findpos"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.findpos"), gen6807)

		gen6816 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V864 := __args[0]
			_ = V864
			V865 := __args[1]
			_ = V865
			gen6814 := Call(__e, ShenFunc(sym_a), MakeNumber(1), V864)

			var gen6815 Obj
			if True == gen6814 {
				gen6813 := Call(__e, ShenFunc(symcons_2), V865)

				if True == gen6813 {
					gen6815 = True
				} else {
					gen6815 = False
				}

			} else {
				gen6815 = False
			}
			if True == gen6815 {
				__e.TailApply(ShenFunc(symtl), V865)

				return
			} else {
				gen6812 := Call(__e, ShenFunc(symcons_2), V865)

				if True == gen6812 {
					gen6808 := Call(__e, ShenFunc(symhd), V865)

					gen6809 := Call(__e, ShenFunc(sym_1), V864, MakeNumber(1))

					gen6810 := Call(__e, ShenFunc(symtl), V865)

					gen6811 := Call(__e, ShenFunc(symshen_4remove_1nth), gen6809, gen6810)

					__e.TailApply(ShenFunc(symcons), gen6808, gen6811)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.remove-nth"))

					return
				}

			}

		}, 2)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.remove-nth"), gen6816)

		return

	}, 0))
}
