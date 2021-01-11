package main

import . "github.com/tiancaiamao/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen19897 := MakeNative(func(__e *ControlFlow) {
		V4594 := __e.Get(1)
		_ = V4594
		V4595 := __e.Get(2)
		_ = V4595
		gen19895 := MakeNative(func(__e *ControlFlow) {
			gen19894 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prh)

			__e.TailApply(gen19894, V4594, V4595, MakeNumber(0))

			return

		}, 0)

		gen19896 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V4594)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen19895, gen19896)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), sympr, gen19897)

	gen19901 := MakeNative(func(__e *ControlFlow) {
		V4599 := __e.Get(1)
		_ = V4599
		V4600 := __e.Get(2)
		_ = V4600
		V4601 := __e.Get(3)
		_ = V4601
		gen19898 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prh)

		gen19899 := Call(__e, PrimNS1Value(symns2_1value), symshen_4write_1char_1and_1inc)

		gen19900 := Call(__e, gen19899, V4599, V4600, V4601)

		__e.TailApply(gen19898, V4599, V4600, gen19900)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4prh, gen19901)

	gen19908 := MakeNative(func(__e *ControlFlow) {
		V4605 := __e.Get(1)
		_ = V4605
		V4606 := __e.Get(2)
		_ = V4606
		V4607 := __e.Get(3)
		_ = V4607
		gen19902 := Call(__e, PrimNS1Value(symns2_1value), symwrite_1byte)

		gen19903 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

		gen19904 := Call(__e, PrimNS1Value(symns2_1value), sympos)

		gen19905 := Call(__e, gen19904, V4605, V4607)

		gen19906 := Call(__e, gen19903, gen19905)

		Call(__e, gen19902, gen19906, V4606)

		gen19907 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		__e.TailApply(gen19907, V4607, MakeNumber(1))

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4write_1char_1and_1inc, gen19908)

	gen19917 := MakeNative(func(__e *ControlFlow) {
		V4609 := __e.Get(1)
		_ = V4609
		gen19914 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			gen19909 := MakeNative(func(__e *ControlFlow) {
				Print := __e.Get(1)
				_ = Print
				__e.Return(V4609)
				return
			}, 1)

			gen19910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			gen19911 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen19912 := Call(__e, gen19911)

			gen19913 := Call(__e, gen19910, String, gen19912)

			__e.TailApply(gen19909, gen19913)

			return

		}, 1)

		gen19915 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert)

		gen19916 := Call(__e, gen19915, V4609, MakeString("~S"))

		__e.TailApply(gen19914, gen19916)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symprint, gen19917)

	gen19921 := MakeNative(func(__e *ControlFlow) {
		V4612 := __e.Get(1)
		_ = V4612
		V4613 := __e.Get(2)
		_ = V4613
		gen19919 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen19920 := Call(__e, gen19919, sym_dhush_d)

		if True == gen19920 {
			__e.Return(V4612)
			return
		} else {
			gen19918 := Call(__e, PrimNS1Value(symns2_1value), sympr)

			__e.TailApply(gen19918, V4612, V4613)

			return

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4prhush, gen19921)

	gen19933 := MakeNative(func(__e *ControlFlow) {
		V4616 := __e.Get(1)
		_ = V4616
		V4617 := __e.Get(2)
		_ = V4617
		gen19931 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

		gen19932 := Call(__e, gen19931, V4616)

		if True == gen19932 {
			gen19928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1l)

			gen19929 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

			gen19930 := Call(__e, gen19929, V4616)

			__e.TailApply(gen19928, gen19930, V4617)

			return

		} else {
			if True == True {
				gen19923 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1r)

				gen19924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen19925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen19926 := Call(__e, gen19925, V4616, Nil)

				gen19927 := Call(__e, gen19924, symshen_4proc_1nl, gen19926)

				__e.TailApply(gen19923, gen19927, V4617)

				return

			} else {
				gen19922 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen19922, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr, gen19933)

	gen19947 := MakeNative(func(__e *ControlFlow) {
		V4620 := __e.Get(1)
		_ = V4620
		V4621 := __e.Get(2)
		_ = V4621
		gen19945 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen19946 := Call(__e, gen19945, Nil, V4621)

		if True == gen19946 {
			__e.Return(V4620)
			return
		} else {
			gen19943 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen19944 := Call(__e, gen19943, V4621)

			if True == gen19944 {
				gen19936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1l)

				gen19937 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

				gen19938 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19939 := Call(__e, gen19938, V4621)

				gen19940 := Call(__e, gen19937, gen19939, V4620)

				gen19941 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen19942 := Call(__e, gen19941, V4621)

				__e.TailApply(gen19936, gen19940, gen19942)

				return

			} else {
				if True == True {
					gen19935 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen19935, symshen_4mkstr_1l)

					return

				} else {
					gen19934 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen19934, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr_1l, gen19947)

	gen20164 := MakeNative(func(__e *ControlFlow) {
		V4626 := __e.Get(1)
		_ = V4626
		V4627 := __e.Get(2)
		_ = V4627
		gen20162 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20163 := Call(__e, gen20162, MakeString(""), V4627)

		if True == gen20163 {
			__e.Return(MakeString(""))
			return
		} else {
			gen20159 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			gen20160 := Call(__e, gen20159, V4627)

			var gen20161 Obj

			if True == gen20160 {
				gen20154 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20155 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen20156 := Call(__e, gen20155, V4627, MakeNumber(0))

				gen20157 := Call(__e, gen20154, MakeString("~"), gen20156)

				var gen20158 Obj

				if True == gen20157 {
					gen20149 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					gen20150 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20151 := Call(__e, gen20150, V4627)

					gen20152 := Call(__e, gen20149, gen20151)

					var gen20153 Obj

					if True == gen20152 {
						gen20143 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20144 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						gen20145 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20146 := Call(__e, gen20145, V4627)

						gen20147 := Call(__e, gen20144, gen20146, MakeNumber(0))

						gen20148 := Call(__e, gen20143, MakeString("A"), gen20147)

						if True == gen20148 {
							gen20153 = True
						} else {
							gen20153 = False
						}

					} else {
						gen20153 = False
					}

					if True == gen20153 {
						gen20158 = True
					} else {
						gen20158 = False
					}

				} else {
					gen20158 = False
				}

				if True == gen20158 {
					gen20161 = True
				} else {
					gen20161 = False
				}

			} else {
				gen20161 = False
			}

			if True == gen20161 {
				gen20132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20133 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20134 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20135 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen20136 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen20137 := Call(__e, gen20136, V4627)

				gen20138 := Call(__e, gen20135, gen20137)

				gen20139 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20140 := Call(__e, gen20139, symshen_4a, Nil)

				gen20141 := Call(__e, gen20134, gen20138, gen20140)

				gen20142 := Call(__e, gen20133, V4626, gen20141)

				__e.TailApply(gen20132, symshen_4app, gen20142)

				return

			} else {
				gen20129 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				gen20130 := Call(__e, gen20129, V4627)

				var gen20131 Obj

				if True == gen20130 {
					gen20124 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen20125 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					gen20126 := Call(__e, gen20125, V4627, MakeNumber(0))

					gen20127 := Call(__e, gen20124, MakeString("~"), gen20126)

					var gen20128 Obj

					if True == gen20127 {
						gen20119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						gen20120 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20121 := Call(__e, gen20120, V4627)

						gen20122 := Call(__e, gen20119, gen20121)

						var gen20123 Obj

						if True == gen20122 {
							gen20113 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen20114 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							gen20115 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							gen20116 := Call(__e, gen20115, V4627)

							gen20117 := Call(__e, gen20114, gen20116, MakeNumber(0))

							gen20118 := Call(__e, gen20113, MakeString("R"), gen20117)

							if True == gen20118 {
								gen20123 = True
							} else {
								gen20123 = False
							}

						} else {
							gen20123 = False
						}

						if True == gen20123 {
							gen20128 = True
						} else {
							gen20128 = False
						}

					} else {
						gen20128 = False
					}

					if True == gen20128 {
						gen20131 = True
					} else {
						gen20131 = False
					}

				} else {
					gen20131 = False
				}

				if True == gen20131 {
					gen20102 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20103 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20105 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20106 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20107 := Call(__e, gen20106, V4627)

					gen20108 := Call(__e, gen20105, gen20107)

					gen20109 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20110 := Call(__e, gen20109, symshen_4r, Nil)

					gen20111 := Call(__e, gen20104, gen20108, gen20110)

					gen20112 := Call(__e, gen20103, V4626, gen20111)

					__e.TailApply(gen20102, symshen_4app, gen20112)

					return

				} else {
					gen20099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					gen20100 := Call(__e, gen20099, V4627)

					var gen20101 Obj

					if True == gen20100 {
						gen20094 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20095 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						gen20096 := Call(__e, gen20095, V4627, MakeNumber(0))

						gen20097 := Call(__e, gen20094, MakeString("~"), gen20096)

						var gen20098 Obj

						if True == gen20097 {
							gen20089 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

							gen20090 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							gen20091 := Call(__e, gen20090, V4627)

							gen20092 := Call(__e, gen20089, gen20091)

							var gen20093 Obj

							if True == gen20092 {
								gen20083 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen20084 := Call(__e, PrimNS1Value(symns2_1value), sympos)

								gen20085 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

								gen20086 := Call(__e, gen20085, V4627)

								gen20087 := Call(__e, gen20084, gen20086, MakeNumber(0))

								gen20088 := Call(__e, gen20083, MakeString("S"), gen20087)

								if True == gen20088 {
									gen20093 = True
								} else {
									gen20093 = False
								}

							} else {
								gen20093 = False
							}

							if True == gen20093 {
								gen20098 = True
							} else {
								gen20098 = False
							}

						} else {
							gen20098 = False
						}

						if True == gen20098 {
							gen20101 = True
						} else {
							gen20101 = False
						}

					} else {
						gen20101 = False
					}

					if True == gen20101 {
						gen20072 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20073 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20074 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20075 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20076 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20077 := Call(__e, gen20076, V4627)

						gen20078 := Call(__e, gen20075, gen20077)

						gen20079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20080 := Call(__e, gen20079, symshen_4s, Nil)

						gen20081 := Call(__e, gen20074, gen20078, gen20080)

						gen20082 := Call(__e, gen20073, V4626, gen20081)

						__e.TailApply(gen20072, symshen_4app, gen20082)

						return

					} else {
						gen20070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						gen20071 := Call(__e, gen20070, V4627)

						if True == gen20071 {
							gen20057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4factor_1cn)

							gen20058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen20059 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen20060 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							gen20061 := Call(__e, gen20060, V4627, MakeNumber(0))

							gen20062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen20063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

							gen20064 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							gen20065 := Call(__e, gen20064, V4627)

							gen20066 := Call(__e, gen20063, V4626, gen20065)

							gen20067 := Call(__e, gen20062, gen20066, Nil)

							gen20068 := Call(__e, gen20059, gen20061, gen20067)

							gen20069 := Call(__e, gen20058, symcn, gen20068)

							__e.TailApply(gen20057, gen20069)

							return

						} else {
							gen20054 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen20055 := Call(__e, gen20054, V4627)

							var gen20056 Obj

							if True == gen20055 {
								gen20049 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen20050 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen20051 := Call(__e, gen20050, V4627)

								gen20052 := Call(__e, gen20049, symcn, gen20051)

								var gen20053 Obj

								if True == gen20052 {
									gen20044 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen20045 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen20046 := Call(__e, gen20045, V4627)

									gen20047 := Call(__e, gen20044, gen20046)

									var gen20048 Obj

									if True == gen20047 {
										gen20037 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen20038 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20039 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20040 := Call(__e, gen20039, V4627)

										gen20041 := Call(__e, gen20038, gen20040)

										gen20042 := Call(__e, gen20037, gen20041)

										var gen20043 Obj

										if True == gen20042 {
											gen20029 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen20030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen20031 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen20032 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen20033 := Call(__e, gen20032, V4627)

											gen20034 := Call(__e, gen20031, gen20033)

											gen20035 := Call(__e, gen20030, gen20034)

											gen20036 := Call(__e, gen20029, Nil, gen20035)

											if True == gen20036 {
												gen20043 = True
											} else {
												gen20043 = False
											}

										} else {
											gen20043 = False
										}

										if True == gen20043 {
											gen20048 = True
										} else {
											gen20048 = False
										}

									} else {
										gen20048 = False
									}

									if True == gen20048 {
										gen20053 = True
									} else {
										gen20053 = False
									}

								} else {
									gen20053 = False
								}

								if True == gen20053 {
									gen20056 = True
								} else {
									gen20056 = False
								}

							} else {
								gen20056 = False
							}

							if True == gen20056 {
								gen20012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen20013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen20014 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen20015 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20016 := Call(__e, gen20015, V4627)

								gen20017 := Call(__e, gen20014, gen20016)

								gen20018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen20019 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

								gen20020 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen20021 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20022 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20023 := Call(__e, gen20022, V4627)

								gen20024 := Call(__e, gen20021, gen20023)

								gen20025 := Call(__e, gen20020, gen20024)

								gen20026 := Call(__e, gen20019, V4626, gen20025)

								gen20027 := Call(__e, gen20018, gen20026, Nil)

								gen20028 := Call(__e, gen20013, gen20017, gen20027)

								__e.TailApply(gen20012, symcn, gen20028)

								return

							} else {
								gen20009 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen20010 := Call(__e, gen20009, V4627)

								var gen20011 Obj

								if True == gen20010 {
									gen20004 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen20005 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen20006 := Call(__e, gen20005, V4627)

									gen20007 := Call(__e, gen20004, symshen_4app, gen20006)

									var gen20008 Obj

									if True == gen20007 {
										gen19999 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen20000 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20001 := Call(__e, gen20000, V4627)

										gen20002 := Call(__e, gen19999, gen20001)

										var gen20003 Obj

										if True == gen20002 {
											gen19992 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen19993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen19994 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen19995 := Call(__e, gen19994, V4627)

											gen19996 := Call(__e, gen19993, gen19995)

											gen19997 := Call(__e, gen19992, gen19996)

											var gen19998 Obj

											if True == gen19997 {
												gen19983 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												gen19984 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen19985 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen19986 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen19987 := Call(__e, gen19986, V4627)

												gen19988 := Call(__e, gen19985, gen19987)

												gen19989 := Call(__e, gen19984, gen19988)

												gen19990 := Call(__e, gen19983, gen19989)

												var gen19991 Obj

												if True == gen19990 {
													gen19973 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													gen19974 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen19975 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen19976 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen19977 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen19978 := Call(__e, gen19977, V4627)

													gen19979 := Call(__e, gen19976, gen19978)

													gen19980 := Call(__e, gen19975, gen19979)

													gen19981 := Call(__e, gen19974, gen19980)

													gen19982 := Call(__e, gen19973, Nil, gen19981)

													if True == gen19982 {
														gen19991 = True
													} else {
														gen19991 = False
													}

												} else {
													gen19991 = False
												}

												if True == gen19991 {
													gen19998 = True
												} else {
													gen19998 = False
												}

											} else {
												gen19998 = False
											}

											if True == gen19998 {
												gen20003 = True
											} else {
												gen20003 = False
											}

										} else {
											gen20003 = False
										}

										if True == gen20003 {
											gen20008 = True
										} else {
											gen20008 = False
										}

									} else {
										gen20008 = False
									}

									if True == gen20008 {
										gen20011 = True
									} else {
										gen20011 = False
									}

								} else {
									gen20011 = False
								}

								if True == gen20011 {
									gen19950 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen19951 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen19952 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen19953 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen19954 := Call(__e, gen19953, V4627)

									gen19955 := Call(__e, gen19952, gen19954)

									gen19956 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen19957 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

									gen19958 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen19959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen19960 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen19961 := Call(__e, gen19960, V4627)

									gen19962 := Call(__e, gen19959, gen19961)

									gen19963 := Call(__e, gen19958, gen19962)

									gen19964 := Call(__e, gen19957, V4626, gen19963)

									gen19965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen19966 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen19967 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen19968 := Call(__e, gen19967, V4627)

									gen19969 := Call(__e, gen19966, gen19968)

									gen19970 := Call(__e, gen19965, gen19969)

									gen19971 := Call(__e, gen19956, gen19964, gen19970)

									gen19972 := Call(__e, gen19951, gen19955, gen19971)

									__e.TailApply(gen19950, symshen_4app, gen19972)

									return

								} else {
									if True == True {
										gen19949 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

										__e.TailApply(gen19949, symshen_4insert_1l)

										return

									} else {
										gen19948 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

										__e.TailApply(gen19948, MakeString("no cond match"))

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

	Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1l, gen20164)

	gen20302 := MakeNative(func(__e *ControlFlow) {
		V4629 := __e.Get(1)
		_ = V4629
		gen20299 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen20300 := Call(__e, gen20299, V4629)

		var gen20301 Obj

		if True == gen20300 {
			gen20294 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20295 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20296 := Call(__e, gen20295, V4629)

			gen20297 := Call(__e, gen20294, symcn, gen20296)

			var gen20298 Obj

			if True == gen20297 {
				gen20289 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20290 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20291 := Call(__e, gen20290, V4629)

				gen20292 := Call(__e, gen20289, gen20291)

				var gen20293 Obj

				if True == gen20292 {
					gen20282 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen20283 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20284 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20285 := Call(__e, gen20284, V4629)

					gen20286 := Call(__e, gen20283, gen20285)

					gen20287 := Call(__e, gen20282, gen20286)

					var gen20288 Obj

					if True == gen20287 {
						gen20273 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen20274 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen20275 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20276 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20277 := Call(__e, gen20276, V4629)

						gen20278 := Call(__e, gen20275, gen20277)

						gen20279 := Call(__e, gen20274, gen20278)

						gen20280 := Call(__e, gen20273, gen20279)

						var gen20281 Obj

						if True == gen20280 {
							gen20262 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen20263 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen20264 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen20265 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen20266 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen20267 := Call(__e, gen20266, V4629)

							gen20268 := Call(__e, gen20265, gen20267)

							gen20269 := Call(__e, gen20264, gen20268)

							gen20270 := Call(__e, gen20263, gen20269)

							gen20271 := Call(__e, gen20262, symcn, gen20270)

							var gen20272 Obj

							if True == gen20271 {
								gen20251 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen20252 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20253 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen20254 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20255 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20256 := Call(__e, gen20255, V4629)

								gen20257 := Call(__e, gen20254, gen20256)

								gen20258 := Call(__e, gen20253, gen20257)

								gen20259 := Call(__e, gen20252, gen20258)

								gen20260 := Call(__e, gen20251, gen20259)

								var gen20261 Obj

								if True == gen20260 {
									gen20238 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen20239 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen20240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen20241 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen20242 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen20243 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen20244 := Call(__e, gen20243, V4629)

									gen20245 := Call(__e, gen20242, gen20244)

									gen20246 := Call(__e, gen20241, gen20245)

									gen20247 := Call(__e, gen20240, gen20246)

									gen20248 := Call(__e, gen20239, gen20247)

									gen20249 := Call(__e, gen20238, gen20248)

									var gen20250 Obj

									if True == gen20249 {
										gen20223 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen20224 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20225 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20226 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20227 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen20228 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20229 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen20230 := Call(__e, gen20229, V4629)

										gen20231 := Call(__e, gen20228, gen20230)

										gen20232 := Call(__e, gen20227, gen20231)

										gen20233 := Call(__e, gen20226, gen20232)

										gen20234 := Call(__e, gen20225, gen20233)

										gen20235 := Call(__e, gen20224, gen20234)

										gen20236 := Call(__e, gen20223, Nil, gen20235)

										var gen20237 Obj

										if True == gen20236 {
											gen20214 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen20215 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen20216 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen20217 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen20218 := Call(__e, gen20217, V4629)

											gen20219 := Call(__e, gen20216, gen20218)

											gen20220 := Call(__e, gen20215, gen20219)

											gen20221 := Call(__e, gen20214, Nil, gen20220)

											var gen20222 Obj

											if True == gen20221 {
												gen20207 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

												gen20208 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen20209 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen20210 := Call(__e, gen20209, V4629)

												gen20211 := Call(__e, gen20208, gen20210)

												gen20212 := Call(__e, gen20207, gen20211)

												var gen20213 Obj

												if True == gen20212 {
													gen20195 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

													gen20196 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													gen20197 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen20198 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													gen20199 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen20200 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													gen20201 := Call(__e, gen20200, V4629)

													gen20202 := Call(__e, gen20199, gen20201)

													gen20203 := Call(__e, gen20198, gen20202)

													gen20204 := Call(__e, gen20197, gen20203)

													gen20205 := Call(__e, gen20196, gen20204)

													gen20206 := Call(__e, gen20195, gen20205)

													if True == gen20206 {
														gen20213 = True
													} else {
														gen20213 = False
													}

												} else {
													gen20213 = False
												}

												if True == gen20213 {
													gen20222 = True
												} else {
													gen20222 = False
												}

											} else {
												gen20222 = False
											}

											if True == gen20222 {
												gen20237 = True
											} else {
												gen20237 = False
											}

										} else {
											gen20237 = False
										}

										if True == gen20237 {
											gen20250 = True
										} else {
											gen20250 = False
										}

									} else {
										gen20250 = False
									}

									if True == gen20250 {
										gen20261 = True
									} else {
										gen20261 = False
									}

								} else {
									gen20261 = False
								}

								if True == gen20261 {
									gen20272 = True
								} else {
									gen20272 = False
								}

							} else {
								gen20272 = False
							}

							if True == gen20272 {
								gen20281 = True
							} else {
								gen20281 = False
							}

						} else {
							gen20281 = False
						}

						if True == gen20281 {
							gen20288 = True
						} else {
							gen20288 = False
						}

					} else {
						gen20288 = False
					}

					if True == gen20288 {
						gen20293 = True
					} else {
						gen20293 = False
					}

				} else {
					gen20293 = False
				}

				if True == gen20293 {
					gen20298 = True
				} else {
					gen20298 = False
				}

			} else {
				gen20298 = False
			}

			if True == gen20298 {
				gen20301 = True
			} else {
				gen20301 = False
			}

		} else {
			gen20301 = False
		}

		if True == gen20301 {
			gen20166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20167 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20168 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen20169 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20171 := Call(__e, gen20170, V4629)

			gen20172 := Call(__e, gen20169, gen20171)

			gen20173 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20174 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20175 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20176 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20177 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20178 := Call(__e, gen20177, V4629)

			gen20179 := Call(__e, gen20176, gen20178)

			gen20180 := Call(__e, gen20175, gen20179)

			gen20181 := Call(__e, gen20174, gen20180)

			gen20182 := Call(__e, gen20173, gen20181)

			gen20183 := Call(__e, gen20168, gen20172, gen20182)

			gen20184 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20185 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20186 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20188 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20189 := Call(__e, gen20188, V4629)

			gen20190 := Call(__e, gen20187, gen20189)

			gen20191 := Call(__e, gen20186, gen20190)

			gen20192 := Call(__e, gen20185, gen20191)

			gen20193 := Call(__e, gen20184, gen20192)

			gen20194 := Call(__e, gen20167, gen20183, gen20193)

			__e.TailApply(gen20166, symcn, gen20194)

			return

		} else {
			if True == True {
				__e.Return(V4629)
				return
			} else {
				gen20165 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20165, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4factor_1cn, gen20302)

	gen20344 := MakeNative(func(__e *ControlFlow) {
		V4631 := __e.Get(1)
		_ = V4631
		gen20342 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20343 := Call(__e, gen20342, MakeString(""), V4631)

		if True == gen20343 {
			__e.Return(MakeString(""))
			return
		} else {
			gen20339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			gen20340 := Call(__e, gen20339, V4631)

			var gen20341 Obj

			if True == gen20340 {
				gen20334 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20335 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen20336 := Call(__e, gen20335, V4631, MakeNumber(0))

				gen20337 := Call(__e, gen20334, MakeString("~"), gen20336)

				var gen20338 Obj

				if True == gen20337 {
					gen20329 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					gen20330 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20331 := Call(__e, gen20330, V4631)

					gen20332 := Call(__e, gen20329, gen20331)

					var gen20333 Obj

					if True == gen20332 {
						gen20323 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20324 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						gen20325 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20326 := Call(__e, gen20325, V4631)

						gen20327 := Call(__e, gen20324, gen20326, MakeNumber(0))

						gen20328 := Call(__e, gen20323, MakeString("%"), gen20327)

						if True == gen20328 {
							gen20333 = True
						} else {
							gen20333 = False
						}

					} else {
						gen20333 = False
					}

					if True == gen20333 {
						gen20338 = True
					} else {
						gen20338 = False
					}

				} else {
					gen20338 = False
				}

				if True == gen20338 {
					gen20341 = True
				} else {
					gen20341 = False
				}

			} else {
				gen20341 = False
			}

			if True == gen20341 {
				gen20314 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen20315 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				gen20316 := Call(__e, gen20315, MakeNumber(10))

				gen20317 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

				gen20318 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen20319 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen20320 := Call(__e, gen20319, V4631)

				gen20321 := Call(__e, gen20318, gen20320)

				gen20322 := Call(__e, gen20317, gen20321)

				__e.TailApply(gen20314, gen20316, gen20322)

				return

			} else {
				gen20312 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				gen20313 := Call(__e, gen20312, V4631)

				if True == gen20313 {
					gen20305 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen20306 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					gen20307 := Call(__e, gen20306, V4631, MakeNumber(0))

					gen20308 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

					gen20309 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20310 := Call(__e, gen20309, V4631)

					gen20311 := Call(__e, gen20308, gen20310)

					__e.TailApply(gen20305, gen20307, gen20311)

					return

				} else {
					if True == True {
						gen20304 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen20304, symshen_4proc_1nl)

						return

					} else {
						gen20303 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen20303, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4proc_1nl, gen20344)

	gen20362 := MakeNative(func(__e *ControlFlow) {
		V4634 := __e.Get(1)
		_ = V4634
		V4635 := __e.Get(2)
		_ = V4635
		gen20360 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20361 := Call(__e, gen20360, Nil, V4635)

		if True == gen20361 {
			__e.Return(V4634)
			return
		} else {
			gen20358 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen20359 := Call(__e, gen20358, V4635)

			if True == gen20359 {
				gen20347 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1r)

				gen20348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20349 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20350 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen20351 := Call(__e, gen20350, V4635)

				gen20352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20353 := Call(__e, gen20352, V4634, Nil)

				gen20354 := Call(__e, gen20349, gen20351, gen20353)

				gen20355 := Call(__e, gen20348, symshen_4insert, gen20354)

				gen20356 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20357 := Call(__e, gen20356, V4635)

				__e.TailApply(gen20347, gen20355, gen20357)

				return

			} else {
				if True == True {
					gen20346 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen20346, symshen_4mkstr_1r)

					return

				} else {
					gen20345 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen20345, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr_1r, gen20362)

	gen20364 := MakeNative(func(__e *ControlFlow) {
		V4638 := __e.Get(1)
		_ = V4638
		V4639 := __e.Get(2)
		_ = V4639
		gen20363 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1h)

		__e.TailApply(gen20363, V4638, V4639, MakeString(""))

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4insert, gen20364)

	gen20456 := MakeNative(func(__e *ControlFlow) {
		V4645 := __e.Get(1)
		_ = V4645
		V4646 := __e.Get(2)
		_ = V4646
		V4647 := __e.Get(3)
		_ = V4647
		gen20454 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20455 := Call(__e, gen20454, MakeString(""), V4646)

		if True == gen20455 {
			__e.Return(V4647)
			return
		} else {
			gen20451 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			gen20452 := Call(__e, gen20451, V4646)

			var gen20453 Obj

			if True == gen20452 {
				gen20446 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20447 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen20448 := Call(__e, gen20447, V4646, MakeNumber(0))

				gen20449 := Call(__e, gen20446, MakeString("~"), gen20448)

				var gen20450 Obj

				if True == gen20449 {
					gen20441 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					gen20442 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20443 := Call(__e, gen20442, V4646)

					gen20444 := Call(__e, gen20441, gen20443)

					var gen20445 Obj

					if True == gen20444 {
						gen20435 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20436 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						gen20437 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20438 := Call(__e, gen20437, V4646)

						gen20439 := Call(__e, gen20436, gen20438, MakeNumber(0))

						gen20440 := Call(__e, gen20435, MakeString("A"), gen20439)

						if True == gen20440 {
							gen20445 = True
						} else {
							gen20445 = False
						}

					} else {
						gen20445 = False
					}

					if True == gen20445 {
						gen20450 = True
					} else {
						gen20450 = False
					}

				} else {
					gen20450 = False
				}

				if True == gen20450 {
					gen20453 = True
				} else {
					gen20453 = False
				}

			} else {
				gen20453 = False
			}

			if True == gen20453 {
				gen20428 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen20429 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen20430 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen20431 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen20432 := Call(__e, gen20431, V4646)

				gen20433 := Call(__e, gen20430, gen20432)

				gen20434 := Call(__e, gen20429, V4645, gen20433, symshen_4a)

				__e.TailApply(gen20428, V4647, gen20434)

				return

			} else {
				gen20425 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				gen20426 := Call(__e, gen20425, V4646)

				var gen20427 Obj

				if True == gen20426 {
					gen20420 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen20421 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					gen20422 := Call(__e, gen20421, V4646, MakeNumber(0))

					gen20423 := Call(__e, gen20420, MakeString("~"), gen20422)

					var gen20424 Obj

					if True == gen20423 {
						gen20415 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						gen20416 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20417 := Call(__e, gen20416, V4646)

						gen20418 := Call(__e, gen20415, gen20417)

						var gen20419 Obj

						if True == gen20418 {
							gen20409 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen20410 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							gen20411 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							gen20412 := Call(__e, gen20411, V4646)

							gen20413 := Call(__e, gen20410, gen20412, MakeNumber(0))

							gen20414 := Call(__e, gen20409, MakeString("R"), gen20413)

							if True == gen20414 {
								gen20419 = True
							} else {
								gen20419 = False
							}

						} else {
							gen20419 = False
						}

						if True == gen20419 {
							gen20424 = True
						} else {
							gen20424 = False
						}

					} else {
						gen20424 = False
					}

					if True == gen20424 {
						gen20427 = True
					} else {
						gen20427 = False
					}

				} else {
					gen20427 = False
				}

				if True == gen20427 {
					gen20402 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen20403 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen20404 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20405 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen20406 := Call(__e, gen20405, V4646)

					gen20407 := Call(__e, gen20404, gen20406)

					gen20408 := Call(__e, gen20403, V4645, gen20407, symshen_4r)

					__e.TailApply(gen20402, V4647, gen20408)

					return

				} else {
					gen20399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					gen20400 := Call(__e, gen20399, V4646)

					var gen20401 Obj

					if True == gen20400 {
						gen20394 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20395 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						gen20396 := Call(__e, gen20395, V4646, MakeNumber(0))

						gen20397 := Call(__e, gen20394, MakeString("~"), gen20396)

						var gen20398 Obj

						if True == gen20397 {
							gen20389 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

							gen20390 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							gen20391 := Call(__e, gen20390, V4646)

							gen20392 := Call(__e, gen20389, gen20391)

							var gen20393 Obj

							if True == gen20392 {
								gen20383 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen20384 := Call(__e, PrimNS1Value(symns2_1value), sympos)

								gen20385 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

								gen20386 := Call(__e, gen20385, V4646)

								gen20387 := Call(__e, gen20384, gen20386, MakeNumber(0))

								gen20388 := Call(__e, gen20383, MakeString("S"), gen20387)

								if True == gen20388 {
									gen20393 = True
								} else {
									gen20393 = False
								}

							} else {
								gen20393 = False
							}

							if True == gen20393 {
								gen20398 = True
							} else {
								gen20398 = False
							}

						} else {
							gen20398 = False
						}

						if True == gen20398 {
							gen20401 = True
						} else {
							gen20401 = False
						}

					} else {
						gen20401 = False
					}

					if True == gen20401 {
						gen20376 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						gen20377 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						gen20378 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20379 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						gen20380 := Call(__e, gen20379, V4646)

						gen20381 := Call(__e, gen20378, gen20380)

						gen20382 := Call(__e, gen20377, V4645, gen20381, symshen_4s)

						__e.TailApply(gen20376, V4647, gen20382)

						return

					} else {
						gen20374 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						gen20375 := Call(__e, gen20374, V4646)

						if True == gen20375 {
							gen20367 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1h)

							gen20368 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							gen20369 := Call(__e, gen20368, V4646)

							gen20370 := Call(__e, PrimNS1Value(symns2_1value), symcn)

							gen20371 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							gen20372 := Call(__e, gen20371, V4646, MakeNumber(0))

							gen20373 := Call(__e, gen20370, V4647, gen20372)

							__e.TailApply(gen20367, V4645, gen20369, gen20373)

							return

						} else {
							if True == True {
								gen20366 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

								__e.TailApply(gen20366, symshen_4insert_1h)

								return

							} else {
								gen20365 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

								__e.TailApply(gen20365, MakeString("no cond match"))

								return

							}
						}

					}

				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1h, gen20456)

	gen20460 := MakeNative(func(__e *ControlFlow) {
		V4651 := __e.Get(1)
		_ = V4651
		V4652 := __e.Get(2)
		_ = V4652
		V4653 := __e.Get(3)
		_ = V4653
		gen20457 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen20458 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

		gen20459 := Call(__e, gen20458, V4651, V4653)

		__e.TailApply(gen20457, gen20459, V4652)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4app, gen20460)

	gen20476 := MakeNative(func(__e *ControlFlow) {
		V4661 := __e.Get(1)
		_ = V4661
		V4662 := __e.Get(2)
		_ = V4662
		gen20472 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20473 := Call(__e, PrimNS1Value(symns2_1value), symfail)

		gen20474 := Call(__e, gen20473)

		gen20475 := Call(__e, gen20472, V4661, gen20474)

		if True == gen20475 {
			__e.Return(MakeString("..."))
			return
		} else {
			gen20470 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list_2)

			gen20471 := Call(__e, gen20470, V4661)

			if True == gen20471 {
				gen20469 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list_1_6str)

				__e.TailApply(gen20469, V4661, V4662)

				return

			} else {
				gen20467 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

				gen20468 := Call(__e, gen20467, V4661)

				if True == gen20468 {
					gen20466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4str_1_6str)

					__e.TailApply(gen20466, V4661, V4662)

					return

				} else {
					gen20464 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

					gen20465 := Call(__e, gen20464, V4661)

					if True == gen20465 {
						gen20463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4vector_1_6str)

						__e.TailApply(gen20463, V4661, V4662)

						return

					} else {
						if True == True {
							gen20462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4atom_1_6str)

							__e.TailApply(gen20462, V4661)

							return

						} else {
							gen20461 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen20461, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4arg_1_6str, gen20476)

	gen20494 := MakeNative(func(__e *ControlFlow) {
		V4665 := __e.Get(1)
		_ = V4665
		V4666 := __e.Get(2)
		_ = V4666
		gen20492 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20493 := Call(__e, gen20492, symshen_4r, V4666)

		if True == gen20493 {
			gen20485 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			gen20486 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			gen20487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1list)

			gen20488 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

			gen20489 := Call(__e, gen20488)

			gen20490 := Call(__e, gen20487, V4665, symshen_4r, gen20489)

			gen20491 := Call(__e, gen20486, gen20490, MakeString(")"))

			__e.TailApply(gen20485, MakeString("("), gen20491)

			return

		} else {
			if True == True {
				gen20478 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20479 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1list)

				gen20481 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

				gen20482 := Call(__e, gen20481)

				gen20483 := Call(__e, gen20480, V4665, V4666, gen20482)

				gen20484 := Call(__e, gen20479, gen20483, MakeString("]"))

				__e.TailApply(gen20478, MakeString("["), gen20484)

				return

			} else {
				gen20477 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20477, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4list_1_6str, gen20494)

	gen20496 := MakeNative(func(__e *ControlFlow) {
		gen20495 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen20495, sym_dmaximum_1print_1sequence_1size_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4maxseq, gen20496)

	gen20532 := MakeNative(func(__e *ControlFlow) {
		V4680 := __e.Get(1)
		_ = V4680
		V4681 := __e.Get(2)
		_ = V4681
		V4682 := __e.Get(3)
		_ = V4682
		gen20530 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20531 := Call(__e, gen20530, Nil, V4680)

		if True == gen20531 {
			__e.Return(MakeString(""))
			return
		} else {
			gen20528 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20529 := Call(__e, gen20528, MakeNumber(0), V4682)

			if True == gen20529 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				gen20525 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20526 := Call(__e, gen20525, V4680)

				var gen20527 Obj

				if True == gen20526 {
					gen20521 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen20522 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20523 := Call(__e, gen20522, V4680)

					gen20524 := Call(__e, gen20521, Nil, gen20523)

					if True == gen20524 {
						gen20527 = True
					} else {
						gen20527 = False
					}

				} else {
					gen20527 = False
				}

				if True == gen20527 {
					gen20518 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

					gen20519 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen20520 := Call(__e, gen20519, V4680)

					__e.TailApply(gen20518, gen20520, V4681)

					return

				} else {
					gen20516 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen20517 := Call(__e, gen20516, V4680)

					if True == gen20517 {
						gen20503 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						gen20504 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

						gen20505 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen20506 := Call(__e, gen20505, V4680)

						gen20507 := Call(__e, gen20504, gen20506, V4681)

						gen20508 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						gen20509 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1list)

						gen20510 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20511 := Call(__e, gen20510, V4680)

						gen20512 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						gen20513 := Call(__e, gen20512, V4682, MakeNumber(1))

						gen20514 := Call(__e, gen20509, gen20511, V4681, gen20513)

						gen20515 := Call(__e, gen20508, MakeString(" "), gen20514)

						__e.TailApply(gen20503, gen20507, gen20515)

						return

					} else {
						if True == True {
							gen20498 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

							gen20499 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

							gen20500 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

							gen20501 := Call(__e, gen20500, V4680, V4681)

							gen20502 := Call(__e, gen20499, MakeString(" "), gen20501)

							__e.TailApply(gen20498, MakeString("|"), gen20502)

							return

						} else {
							gen20497 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen20497, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4iter_1list, gen20532)

	gen20543 := MakeNative(func(__e *ControlFlow) {
		V4689 := __e.Get(1)
		_ = V4689
		V4690 := __e.Get(2)
		_ = V4690
		gen20541 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20542 := Call(__e, gen20541, symshen_4a, V4690)

		if True == gen20542 {
			__e.Return(V4689)
			return
		} else {
			if True == True {
				gen20534 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20535 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				gen20536 := Call(__e, gen20535, MakeNumber(34))

				gen20537 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20538 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				gen20539 := Call(__e, gen20538, MakeNumber(34))

				gen20540 := Call(__e, gen20537, V4689, gen20539)

				__e.TailApply(gen20534, gen20536, gen20540)

				return

			} else {
				gen20533 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20533, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4str_1_6str, gen20543)

	gen20568 := MakeNative(func(__e *ControlFlow) {
		V4693 := __e.Get(1)
		_ = V4693
		V4694 := __e.Get(2)
		_ = V4694
		gen20566 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1vector_2)

		gen20567 := Call(__e, gen20566, V4693)

		if True == gen20567 {
			gen20562 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

			gen20563 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			gen20564 := Call(__e, gen20563, V4693, MakeNumber(0))

			gen20565 := Call(__e, gen20562, gen20564)

			__e.TailApply(gen20565, V4693)

			return

		} else {
			gen20560 := Call(__e, PrimNS1Value(symns2_1value), symvector_2)

			gen20561 := Call(__e, gen20560, V4693)

			if True == gen20561 {
				gen20553 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20554 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20555 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1vector)

				gen20556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

				gen20557 := Call(__e, gen20556)

				gen20558 := Call(__e, gen20555, V4693, MakeNumber(1), V4694, gen20557)

				gen20559 := Call(__e, gen20554, gen20558, MakeString(">"))

				__e.TailApply(gen20553, MakeString("<"), gen20559)

				return

			} else {
				gen20544 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20545 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20546 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				gen20547 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1vector)

				gen20548 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

				gen20549 := Call(__e, gen20548)

				gen20550 := Call(__e, gen20547, V4693, MakeNumber(0), V4694, gen20549)

				gen20551 := Call(__e, gen20546, gen20550, MakeString(">>"))

				gen20552 := Call(__e, gen20545, MakeString("<"), gen20551)

				__e.TailApply(gen20544, MakeString("<"), gen20552)

				return

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4vector_1_6str, gen20568)

	gen20572 := MakeNative(func(__e *ControlFlow) {
		V4696 := __e.Get(1)
		_ = V4696
		gen20569 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20570 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen20571 := Call(__e, gen20570, symshen_4_dempty_1absvector_d)

		__e.TailApply(gen20569, V4696, gen20571)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4empty_1absvector_2, gen20572)

	gen20596 := MakeNative(func(__e *ControlFlow) {
		V4698 := __e.Get(1)
		_ = V4698
		gen20592 := Call(__e, PrimNS1Value(symns2_1value), symnot)

		gen20593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4empty_1absvector_2)

		gen20594 := Call(__e, gen20593, V4698)

		gen20595 := Call(__e, gen20592, gen20594)

		if True == gen20595 {
			gen20588 := MakeNative(func(__e *ControlFlow) {
				First := __e.Get(1)
				_ = First
				gen20586 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20587 := Call(__e, gen20586, First, symshen_4tuple)

				if True == gen20587 {
					__e.Return(True)
					return
				} else {
					gen20583 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen20584 := Call(__e, gen20583, First, symshen_4pvar)

					var gen20585 Obj

					if True == gen20584 {
						gen20585 = True
					} else {
						gen20580 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20581 := Call(__e, gen20580, First, symshen_4dictionary)

						var gen20582 Obj

						if True == gen20581 {
							gen20582 = True
						} else {
							gen20575 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen20576 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

							gen20577 := Call(__e, gen20576, First)

							gen20578 := Call(__e, gen20575, gen20577)

							var gen20579 Obj

							if True == gen20578 {
								gen20573 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fbound_2)

								gen20574 := Call(__e, gen20573, First)

								if True == gen20574 {
									gen20579 = True
								} else {
									gen20579 = False
								}

							} else {
								gen20579 = False
							}

							if True == gen20579 {
								gen20582 = True
							} else {
								gen20582 = False
							}

						}

						if True == gen20582 {
							gen20585 = True
						} else {
							gen20585 = False
						}

					}

					if True == gen20585 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}, 1)

			gen20589 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			gen20590 := Call(__e, gen20589, V4698, MakeNumber(0))

			gen20591 := Call(__e, gen20588, gen20590)

			if True == gen20591 {
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

	Call(__e, PrimNS1Value(symns2_1set), symshen_4print_1vector_2, gen20596)

	gen20600 := MakeNative(func(__e *ControlFlow) {
		V4700 := __e.Get(1)
		_ = V4700
		gen20598 := MakeNative(func(__e *ControlFlow) {
			gen20597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lookup_1func)

			Call(__e, gen20597, V4700)

			__e.Return(True)
			return

		}, 0)

		gen20599 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen20598, gen20599)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4fbound_2, gen20600)

	gen20612 := MakeNative(func(__e *ControlFlow) {
		V4702 := __e.Get(1)
		_ = V4702
		gen20601 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen20602 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen20603 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		gen20604 := Call(__e, gen20603, V4702, MakeNumber(1))

		gen20605 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen20606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen20607 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		gen20608 := Call(__e, gen20607, V4702, MakeNumber(2))

		gen20609 := Call(__e, gen20606, gen20608, MakeString(")"), symshen_4s)

		gen20610 := Call(__e, gen20605, MakeString(" "), gen20609)

		gen20611 := Call(__e, gen20602, gen20604, gen20610, symshen_4s)

		__e.TailApply(gen20601, MakeString("(@p "), gen20611)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4tuple, gen20612)

	gen20613 := MakeNative(func(__e *ControlFlow) {
		V4704 := __e.Get(1)
		_ = V4704
		__e.Return(MakeString("(dict ...)"))
		return
	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dictionary, gen20613)

	gen20645 := MakeNative(func(__e *ControlFlow) {
		V4715 := __e.Get(1)
		_ = V4715
		V4716 := __e.Get(2)
		_ = V4716
		V4717 := __e.Get(3)
		_ = V4717
		V4718 := __e.Get(4)
		_ = V4718
		gen20643 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20644 := Call(__e, gen20643, MakeNumber(0), V4718)

		if True == gen20644 {
			__e.Return(MakeString("... etc"))
			return
		} else {
			if True == True {
				gen20638 := MakeNative(func(__e *ControlFlow) {
					Item := __e.Get(1)
					_ = Item
					gen20631 := MakeNative(func(__e *ControlFlow) {
						Next := __e.Get(1)
						_ = Next
						gen20629 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20630 := Call(__e, gen20629, Item, symshen_4out_1of_1bounds)

						if True == gen20630 {
							__e.Return(MakeString(""))
							return
						} else {
							gen20627 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen20628 := Call(__e, gen20627, Next, symshen_4out_1of_1bounds)

							if True == gen20628 {
								gen20626 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

								__e.TailApply(gen20626, Item, V4717)

								return

							} else {
								gen20615 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

								gen20616 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

								gen20617 := Call(__e, gen20616, Item, V4717)

								gen20618 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

								gen20619 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1vector)

								gen20620 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

								gen20621 := Call(__e, gen20620, V4716, MakeNumber(1))

								gen20622 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

								gen20623 := Call(__e, gen20622, V4718, MakeNumber(1))

								gen20624 := Call(__e, gen20619, V4715, gen20621, V4717, gen20623)

								gen20625 := Call(__e, gen20618, MakeString(" "), gen20624)

								__e.TailApply(gen20615, gen20617, gen20625)

								return

							}

						}

					}, 1)

					gen20635 := MakeNative(func(__e *ControlFlow) {
						gen20632 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

						gen20633 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

						gen20634 := Call(__e, gen20633, V4716, MakeNumber(1))

						__e.TailApply(gen20632, V4715, gen20634)

						return

					}, 0)

					gen20636 := MakeNative(func(__e *ControlFlow) {
						E := __e.Get(1)
						_ = E
						__e.Return(symshen_4out_1of_1bounds)
						return
					}, 1)

					gen20637 := Call(__e, PrimNS1Value(symtry_1catch), gen20635, gen20636)

					__e.TailApply(gen20631, gen20637)

					return

				}, 1)

				gen20640 := MakeNative(func(__e *ControlFlow) {
					gen20639 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

					__e.TailApply(gen20639, V4715, V4716)

					return

				}, 0)

				gen20641 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4out_1of_1bounds)
					return
				}, 1)

				gen20642 := Call(__e, PrimNS1Value(symtry_1catch), gen20640, gen20641)

				__e.TailApply(gen20638, gen20642)

				return

			} else {
				gen20614 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20614, MakeString("no cond match"))

				return

			}
		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4iter_1vector, gen20645)

	gen20650 := MakeNative(func(__e *ControlFlow) {
		V4720 := __e.Get(1)
		_ = V4720
		gen20647 := MakeNative(func(__e *ControlFlow) {
			gen20646 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			__e.TailApply(gen20646, V4720)

			return

		}, 0)

		gen20649 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen20648 := Call(__e, PrimNS1Value(symns2_1value), symshen_4funexstring)

			__e.TailApply(gen20648)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen20647, gen20649)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4atom_1_6str, gen20650)

	gen20668 := MakeNative(func(__e *ControlFlow) {
		gen20651 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		gen20652 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		gen20653 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		gen20654 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		gen20655 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		gen20656 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		gen20657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

		gen20658 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		gen20659 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		gen20660 := Call(__e, gen20659, MakeString("x"))

		gen20661 := Call(__e, gen20658, gen20660)

		gen20662 := Call(__e, gen20657, gen20661, symshen_4a)

		gen20663 := Call(__e, gen20656, gen20662, MakeString("\x11"))

		gen20664 := Call(__e, gen20655, MakeString("e"), gen20663)

		gen20665 := Call(__e, gen20654, MakeString("n"), gen20664)

		gen20666 := Call(__e, gen20653, MakeString("u"), gen20665)

		gen20667 := Call(__e, gen20652, MakeString("f"), gen20666)

		__e.TailApply(gen20651, MakeString("\x10"), gen20667)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4funexstring, gen20668)

	gen20673 := MakeNative(func(__e *ControlFlow) {
		V4722 := __e.Get(1)
		_ = V4722
		gen20671 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

		gen20672 := Call(__e, gen20671, V4722)

		if True == gen20672 {
			__e.Return(True)
			return
		} else {
			gen20669 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen20670 := Call(__e, gen20669, V4722)

			if True == gen20670 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4list_2, gen20673)

	return

}, 0)
