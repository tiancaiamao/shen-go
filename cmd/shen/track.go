package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen8979 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3135 := __args[0]
			_ = V3135
			gen8969 := Call(__e, ShenFunc(symshen_4app), V3135, MakeString(";\n"), MakeSymbol("shen.a"))

			gen8970 := Call(__e, ShenFunc(symcn), MakeString("partial function "), gen8969)

			gen8971 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen8970, gen8971)

			gen8976 := Call(__e, ShenFunc(symshen_4tracked_2), V3135)

			gen8977 := Call(__e, ShenFunc(symnot), gen8976)

			var gen8978 Obj
			if True == gen8977 {
				gen8973 := Call(__e, ShenFunc(symshen_4app), V3135, MakeString("? "), MakeSymbol("shen.a"))

				gen8974 := Call(__e, ShenFunc(symcn), MakeString("track "), gen8973)

				gen8975 := Call(__e, ShenFunc(symy_1or_1n_2), gen8974)

				if True == gen8975 {
					gen8978 = True
				} else {
					gen8978 = False
				}

			} else {
				gen8978 = False
			}
			if True == gen8978 {
				gen8972 := Call(__e, ShenFunc(symps), V3135)

				Call(__e, ShenFunc(symshen_4track_1function), gen8972)

			} else {
				MakeSymbol("shen.ok")
			}

			__e.TailApply(ShenFunc(symsimple_1error), MakeString("aborted"))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.f_error"), gen8979)

		gen8981 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3137 := __args[0]
			_ = V3137
			gen8980 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tracking*"))

			__e.TailApply(ShenFunc(symelement_2), V3137, gen8980)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tracked?"), gen8981)

		gen8983 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3139 := __args[0]
			_ = V3139
			gen8982 := Call(__e, ShenFunc(symps), V3139)

			Source := gen8982
			__e.TailApply(ShenFunc(symshen_4track_1function), Source)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("track"), gen8983)

		gen9029 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3141 := __args[0]
			_ = V3141
			gen9027 := Call(__e, ShenFunc(symcons_2), V3141)

			var gen9028 Obj
			if True == gen9027 {
				gen9024 := Call(__e, ShenFunc(symhd), V3141)

				gen9025 := Call(__e, ShenFunc(sym_a), MakeSymbol("defun"), gen9024)

				var gen9026 Obj
				if True == gen9025 {
					gen9021 := Call(__e, ShenFunc(symtl), V3141)

					gen9022 := Call(__e, ShenFunc(symcons_2), gen9021)

					var gen9023 Obj
					if True == gen9022 {
						gen9017 := Call(__e, ShenFunc(symtl), V3141)

						gen9018 := Call(__e, ShenFunc(symtl), gen9017)

						gen9019 := Call(__e, ShenFunc(symcons_2), gen9018)

						var gen9020 Obj
						if True == gen9019 {
							gen9012 := Call(__e, ShenFunc(symtl), V3141)

							gen9013 := Call(__e, ShenFunc(symtl), gen9012)

							gen9014 := Call(__e, ShenFunc(symtl), gen9013)

							gen9015 := Call(__e, ShenFunc(symcons_2), gen9014)

							var gen9016 Obj
							if True == gen9015 {
								gen9007 := Call(__e, ShenFunc(symtl), V3141)

								gen9008 := Call(__e, ShenFunc(symtl), gen9007)

								gen9009 := Call(__e, ShenFunc(symtl), gen9008)

								gen9010 := Call(__e, ShenFunc(symtl), gen9009)

								gen9011 := Call(__e, ShenFunc(sym_a), Nil, gen9010)

								if True == gen9011 {
									gen9016 = True
								} else {
									gen9016 = False
								}

							} else {
								gen9016 = False
							}
							if True == gen9016 {
								gen9020 = True
							} else {
								gen9020 = False
							}

						} else {
							gen9020 = False
						}
						if True == gen9020 {
							gen9023 = True
						} else {
							gen9023 = False
						}

					} else {
						gen9023 = False
					}
					if True == gen9023 {
						gen9026 = True
					} else {
						gen9026 = False
					}

				} else {
					gen9026 = False
				}
				if True == gen9026 {
					gen9028 = True
				} else {
					gen9028 = False
				}

			} else {
				gen9028 = False
			}
			if True == gen9028 {
				gen8984 := Call(__e, ShenFunc(symtl), V3141)

				gen8985 := Call(__e, ShenFunc(symhd), gen8984)

				gen8986 := Call(__e, ShenFunc(symtl), V3141)

				gen8987 := Call(__e, ShenFunc(symtl), gen8986)

				gen8988 := Call(__e, ShenFunc(symhd), gen8987)

				gen8989 := Call(__e, ShenFunc(symtl), V3141)

				gen8990 := Call(__e, ShenFunc(symhd), gen8989)

				gen8991 := Call(__e, ShenFunc(symtl), V3141)

				gen8992 := Call(__e, ShenFunc(symtl), gen8991)

				gen8993 := Call(__e, ShenFunc(symhd), gen8992)

				gen8994 := Call(__e, ShenFunc(symtl), V3141)

				gen8995 := Call(__e, ShenFunc(symtl), gen8994)

				gen8996 := Call(__e, ShenFunc(symtl), gen8995)

				gen8997 := Call(__e, ShenFunc(symhd), gen8996)

				gen8998 := Call(__e, ShenFunc(symshen_4insert_1tracking_1code), gen8990, gen8993, gen8997)

				gen8999 := Call(__e, ShenFunc(symcons), gen8998, Nil)

				gen9000 := Call(__e, ShenFunc(symcons), gen8988, gen8999)

				gen9001 := Call(__e, ShenFunc(symcons), gen8985, gen9000)

				gen9002 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen9001)

				KL := gen9002
				gen9003 := Call(__e, ShenFunc(symeval_1kl), KL)

				Ob := gen9003
				gen9004 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tracking*"))

				gen9005 := Call(__e, ShenFunc(symcons), Ob, gen9004)

				gen9006 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*tracking*"), gen9005)

				Tr := gen9006
				_ = Tr
				__e.Return(Ob)
				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.track-function"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.track-function"), gen9029)

		gen9082 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3145 := __args[0]
			_ = V3145
			V3146 := __args[1]
			_ = V3146
			V3147 := __args[2]
			_ = V3147
			gen9030 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9031 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9030)

			gen9032 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

			gen9033 := Call(__e, ShenFunc(symcons), gen9031, gen9032)

			gen9034 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen9033)

			gen9035 := Call(__e, ShenFunc(symcons), gen9034, Nil)

			gen9036 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), gen9035)

			gen9037 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen9036)

			gen9038 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9039 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9038)

			gen9040 := Call(__e, ShenFunc(symshen_4cons__form), V3146)

			gen9041 := Call(__e, ShenFunc(symcons), gen9040, Nil)

			gen9042 := Call(__e, ShenFunc(symcons), V3145, gen9041)

			gen9043 := Call(__e, ShenFunc(symcons), gen9039, gen9042)

			gen9044 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.input-track"), gen9043)

			gen9045 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.terpri-or-read-char"), Nil)

			gen9046 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9047 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9046)

			gen9048 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

			gen9049 := Call(__e, ShenFunc(symcons), V3145, gen9048)

			gen9050 := Call(__e, ShenFunc(symcons), gen9047, gen9049)

			gen9051 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.output-track"), gen9050)

			gen9052 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9053 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9052)

			gen9054 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

			gen9055 := Call(__e, ShenFunc(symcons), gen9053, gen9054)

			gen9056 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen9055)

			gen9057 := Call(__e, ShenFunc(symcons), gen9056, Nil)

			gen9058 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), gen9057)

			gen9059 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen9058)

			gen9060 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.terpri-or-read-char"), Nil)

			gen9061 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

			gen9062 := Call(__e, ShenFunc(symcons), gen9060, gen9061)

			gen9063 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9062)

			gen9064 := Call(__e, ShenFunc(symcons), gen9063, Nil)

			gen9065 := Call(__e, ShenFunc(symcons), gen9059, gen9064)

			gen9066 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9065)

			gen9067 := Call(__e, ShenFunc(symcons), gen9066, Nil)

			gen9068 := Call(__e, ShenFunc(symcons), gen9051, gen9067)

			gen9069 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9068)

			gen9070 := Call(__e, ShenFunc(symcons), gen9069, Nil)

			gen9071 := Call(__e, ShenFunc(symcons), V3147, gen9070)

			gen9072 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen9071)

			gen9073 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9072)

			gen9074 := Call(__e, ShenFunc(symcons), gen9073, Nil)

			gen9075 := Call(__e, ShenFunc(symcons), gen9045, gen9074)

			gen9076 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9075)

			gen9077 := Call(__e, ShenFunc(symcons), gen9076, Nil)

			gen9078 := Call(__e, ShenFunc(symcons), gen9044, gen9077)

			gen9079 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9078)

			gen9080 := Call(__e, ShenFunc(symcons), gen9079, Nil)

			gen9081 := Call(__e, ShenFunc(symcons), gen9037, gen9080)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen9081)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-tracking-code"), gen9082)

		gen9085 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3153 := __args[0]
			_ = V3153
			gen9084 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V3153)

			if True == gen9084 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*step*"), True)

				return
			} else {
				gen9083 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V3153)

				if True == gen9083 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*step*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("step expects a + or a -.\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("step"), gen9085)

		gen9088 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3159 := __args[0]
			_ = V3159
			gen9087 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V3159)

			if True == gen9087 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*spy*"), True)

				return
			} else {
				gen9086 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V3159)

				if True == gen9086 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*spy*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("spy expects a + or a -.\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("spy"), gen9088)

		gen9092 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen9091 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*step*"))

			if True == gen9091 {
				gen9089 := Call(__e, ShenFunc(symvalue), MakeSymbol("*stinput*"))

				gen9090 := Call(__e, ShenFunc(symread_1byte), gen9089)

				__e.TailApply(ShenFunc(symshen_4check_1byte), gen9090)

				return

			} else {
				__e.TailApply(ShenFunc(symnl), MakeNumber(1))

				return
			}

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.terpri-or-read-char"), gen9092)

		gen9095 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3165 := __args[0]
			_ = V3165
			gen9093 := Call(__e, ShenFunc(symshen_4hat))

			gen9094 := Call(__e, ShenFunc(sym_a), V3165, gen9093)

			if True == gen9094 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("aborted"))

				return
			} else {
				__e.Return(True)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.check-byte"), gen9095)

		gen9107 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3169 := __args[0]
			_ = V3169
			V3170 := __args[1]
			_ = V3170
			V3171 := __args[2]
			_ = V3171
			gen9096 := Call(__e, ShenFunc(symshen_4spaces), V3169)

			gen9097 := Call(__e, ShenFunc(symshen_4spaces), V3169)

			gen9098 := Call(__e, ShenFunc(symshen_4app), gen9097, MakeString(""), MakeSymbol("shen.a"))

			gen9099 := Call(__e, ShenFunc(symcn), MakeString(" \n"), gen9098)

			gen9100 := Call(__e, ShenFunc(symshen_4app), V3170, gen9099, MakeSymbol("shen.a"))

			gen9101 := Call(__e, ShenFunc(symcn), MakeString("> Inputs to "), gen9100)

			gen9102 := Call(__e, ShenFunc(symshen_4app), V3169, gen9101, MakeSymbol("shen.a"))

			gen9103 := Call(__e, ShenFunc(symcn), MakeString("<"), gen9102)

			gen9104 := Call(__e, ShenFunc(symshen_4app), gen9096, gen9103, MakeSymbol("shen.a"))

			gen9105 := Call(__e, ShenFunc(symcn), MakeString("\n"), gen9104)

			gen9106 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen9105, gen9106)

			__e.TailApply(ShenFunc(symshen_4recursively_1print), V3171)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.input-track"), gen9107)

		gen9114 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3173 := __args[0]
			_ = V3173
			gen9113 := Call(__e, ShenFunc(sym_a), Nil, V3173)

			if True == gen9113 {
				gen9112 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), MakeString(" ==>"), gen9112)

				return

			} else {
				gen9111 := Call(__e, ShenFunc(symcons_2), V3173)

				if True == gen9111 {
					gen9108 := Call(__e, ShenFunc(symhd), V3173)

					Call(__e, ShenFunc(symprint), gen9108)

					gen9109 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), MakeString(", "), gen9109)

					gen9110 := Call(__e, ShenFunc(symtl), V3173)

					__e.TailApply(ShenFunc(symshen_4recursively_1print), gen9110)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.recursively-print"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.recursively-print"), gen9114)

		gen9118 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3175 := __args[0]
			_ = V3175
			gen9117 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V3175)

			if True == gen9117 {
				__e.Return(MakeString(""))
				return
			} else {
				gen9115 := Call(__e, ShenFunc(sym_1), V3175, MakeNumber(1))

				gen9116 := Call(__e, ShenFunc(symshen_4spaces), gen9115)

				__e.TailApply(ShenFunc(symcn), MakeString(" "), gen9116)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.spaces"), gen9118)

		gen9132 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3179 := __args[0]
			_ = V3179
			V3180 := __args[1]
			_ = V3180
			V3181 := __args[2]
			_ = V3181
			gen9119 := Call(__e, ShenFunc(symshen_4spaces), V3179)

			gen9120 := Call(__e, ShenFunc(symshen_4spaces), V3179)

			gen9121 := Call(__e, ShenFunc(symshen_4app), V3181, MakeString(""), MakeSymbol("shen.s"))

			gen9122 := Call(__e, ShenFunc(symcn), MakeString("==> "), gen9121)

			gen9123 := Call(__e, ShenFunc(symshen_4app), gen9120, gen9122, MakeSymbol("shen.a"))

			gen9124 := Call(__e, ShenFunc(symcn), MakeString(" \n"), gen9123)

			gen9125 := Call(__e, ShenFunc(symshen_4app), V3180, gen9124, MakeSymbol("shen.a"))

			gen9126 := Call(__e, ShenFunc(symcn), MakeString("> Output from "), gen9125)

			gen9127 := Call(__e, ShenFunc(symshen_4app), V3179, gen9126, MakeSymbol("shen.a"))

			gen9128 := Call(__e, ShenFunc(symcn), MakeString("<"), gen9127)

			gen9129 := Call(__e, ShenFunc(symshen_4app), gen9119, gen9128, MakeSymbol("shen.a"))

			gen9130 := Call(__e, ShenFunc(symcn), MakeString("\n"), gen9129)

			gen9131 := Call(__e, ShenFunc(symstoutput))

			__e.TailApply(ShenFunc(symshen_4prhush), gen9130, gen9131)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.output-track"), gen9132)

		gen9137 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3183 := __args[0]
			_ = V3183
			gen9133 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tracking*"))

			Tracking := gen9133
			gen9134 := Call(__e, ShenFunc(symremove), V3183, Tracking)

			gen9135 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*tracking*"), gen9134)

			_ = gen9135
			gen9136 := Call(__e, ShenFunc(symps), V3183)

			__e.TailApply(ShenFunc(symeval), gen9136)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("untrack"), gen9137)

		gen9139 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3185 := __args[0]
			_ = V3185
			gen9138 := Call(__e, ShenFunc(symps), V3185)

			__e.TailApply(ShenFunc(symshen_4profile_1help), gen9138)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("profile"), gen9139)

		gen9199 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3191 := __args[0]
			_ = V3191
			gen9197 := Call(__e, ShenFunc(symcons_2), V3191)

			var gen9198 Obj
			if True == gen9197 {
				gen9194 := Call(__e, ShenFunc(symhd), V3191)

				gen9195 := Call(__e, ShenFunc(sym_a), MakeSymbol("defun"), gen9194)

				var gen9196 Obj
				if True == gen9195 {
					gen9191 := Call(__e, ShenFunc(symtl), V3191)

					gen9192 := Call(__e, ShenFunc(symcons_2), gen9191)

					var gen9193 Obj
					if True == gen9192 {
						gen9187 := Call(__e, ShenFunc(symtl), V3191)

						gen9188 := Call(__e, ShenFunc(symtl), gen9187)

						gen9189 := Call(__e, ShenFunc(symcons_2), gen9188)

						var gen9190 Obj
						if True == gen9189 {
							gen9182 := Call(__e, ShenFunc(symtl), V3191)

							gen9183 := Call(__e, ShenFunc(symtl), gen9182)

							gen9184 := Call(__e, ShenFunc(symtl), gen9183)

							gen9185 := Call(__e, ShenFunc(symcons_2), gen9184)

							var gen9186 Obj
							if True == gen9185 {
								gen9177 := Call(__e, ShenFunc(symtl), V3191)

								gen9178 := Call(__e, ShenFunc(symtl), gen9177)

								gen9179 := Call(__e, ShenFunc(symtl), gen9178)

								gen9180 := Call(__e, ShenFunc(symtl), gen9179)

								gen9181 := Call(__e, ShenFunc(sym_a), Nil, gen9180)

								if True == gen9181 {
									gen9186 = True
								} else {
									gen9186 = False
								}

							} else {
								gen9186 = False
							}
							if True == gen9186 {
								gen9190 = True
							} else {
								gen9190 = False
							}

						} else {
							gen9190 = False
						}
						if True == gen9190 {
							gen9193 = True
						} else {
							gen9193 = False
						}

					} else {
						gen9193 = False
					}
					if True == gen9193 {
						gen9196 = True
					} else {
						gen9196 = False
					}

				} else {
					gen9196 = False
				}
				if True == gen9196 {
					gen9198 = True
				} else {
					gen9198 = False
				}

			} else {
				gen9198 = False
			}
			if True == gen9198 {
				gen9140 := Call(__e, ShenFunc(symgensym), MakeSymbol("shen.f"))

				G := gen9140
				gen9141 := Call(__e, ShenFunc(symtl), V3191)

				gen9142 := Call(__e, ShenFunc(symhd), gen9141)

				gen9143 := Call(__e, ShenFunc(symtl), V3191)

				gen9144 := Call(__e, ShenFunc(symtl), gen9143)

				gen9145 := Call(__e, ShenFunc(symhd), gen9144)

				gen9146 := Call(__e, ShenFunc(symtl), V3191)

				gen9147 := Call(__e, ShenFunc(symhd), gen9146)

				gen9148 := Call(__e, ShenFunc(symtl), V3191)

				gen9149 := Call(__e, ShenFunc(symtl), gen9148)

				gen9150 := Call(__e, ShenFunc(symhd), gen9149)

				gen9151 := Call(__e, ShenFunc(symtl), V3191)

				gen9152 := Call(__e, ShenFunc(symtl), gen9151)

				gen9153 := Call(__e, ShenFunc(symhd), gen9152)

				gen9154 := Call(__e, ShenFunc(symcons), G, gen9153)

				gen9155 := Call(__e, ShenFunc(symshen_4profile_1func), gen9147, gen9150, gen9154)

				gen9156 := Call(__e, ShenFunc(symcons), gen9155, Nil)

				gen9157 := Call(__e, ShenFunc(symcons), gen9145, gen9156)

				gen9158 := Call(__e, ShenFunc(symcons), gen9142, gen9157)

				gen9159 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen9158)

				Profile := gen9159
				gen9160 := Call(__e, ShenFunc(symtl), V3191)

				gen9161 := Call(__e, ShenFunc(symtl), gen9160)

				gen9162 := Call(__e, ShenFunc(symhd), gen9161)

				gen9163 := Call(__e, ShenFunc(symtl), V3191)

				gen9164 := Call(__e, ShenFunc(symhd), gen9163)

				gen9165 := Call(__e, ShenFunc(symtl), V3191)

				gen9166 := Call(__e, ShenFunc(symtl), gen9165)

				gen9167 := Call(__e, ShenFunc(symtl), gen9166)

				gen9168 := Call(__e, ShenFunc(symhd), gen9167)

				gen9169 := Call(__e, ShenFunc(symsubst), G, gen9164, gen9168)

				gen9170 := Call(__e, ShenFunc(symcons), gen9169, Nil)

				gen9171 := Call(__e, ShenFunc(symcons), gen9162, gen9170)

				gen9172 := Call(__e, ShenFunc(symcons), G, gen9171)

				gen9173 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen9172)

				Def := gen9173
				gen9174 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), Profile)

				CompileProfile := gen9174
				_ = CompileProfile
				gen9175 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), Def)

				CompileG := gen9175
				_ = CompileG
				gen9176 := Call(__e, ShenFunc(symtl), V3191)

				__e.TailApply(ShenFunc(symhd), gen9176)

				return

			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("Cannot profile.\n"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.profile-help"), gen9199)

		gen9200 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3193 := __args[0]
			_ = V3193
			__e.TailApply(ShenFunc(symuntrack), V3193)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unprofile"), gen9200)

		gen9231 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3197 := __args[0]
			_ = V3197
			V3198 := __args[1]
			_ = V3198
			V3199 := __args[2]
			_ = V3199
			gen9201 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

			gen9202 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen9201)

			gen9203 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

			gen9204 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen9203)

			gen9205 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), Nil)

			gen9206 := Call(__e, ShenFunc(symcons), gen9204, gen9205)

			gen9207 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen9206)

			gen9208 := Call(__e, ShenFunc(symcons), V3197, Nil)

			gen9209 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.get-profile"), gen9208)

			gen9210 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), Nil)

			gen9211 := Call(__e, ShenFunc(symcons), gen9209, gen9210)

			gen9212 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen9211)

			gen9213 := Call(__e, ShenFunc(symcons), gen9212, Nil)

			gen9214 := Call(__e, ShenFunc(symcons), V3197, gen9213)

			gen9215 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.put-profile"), gen9214)

			gen9216 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

			gen9217 := Call(__e, ShenFunc(symcons), gen9215, gen9216)

			gen9218 := Call(__e, ShenFunc(symcons), MakeSymbol("Record"), gen9217)

			gen9219 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9218)

			gen9220 := Call(__e, ShenFunc(symcons), gen9219, Nil)

			gen9221 := Call(__e, ShenFunc(symcons), gen9207, gen9220)

			gen9222 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), gen9221)

			gen9223 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9222)

			gen9224 := Call(__e, ShenFunc(symcons), gen9223, Nil)

			gen9225 := Call(__e, ShenFunc(symcons), V3199, gen9224)

			gen9226 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen9225)

			gen9227 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9226)

			gen9228 := Call(__e, ShenFunc(symcons), gen9227, Nil)

			gen9229 := Call(__e, ShenFunc(symcons), gen9202, gen9228)

			gen9230 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), gen9229)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen9230)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.profile-func"), gen9231)

		gen9234 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3201 := __args[0]
			_ = V3201
			gen9232 := Call(__e, ShenFunc(symshen_4get_1profile), V3201)

			Results := gen9232
			gen9233 := Call(__e, ShenFunc(symshen_4put_1profile), V3201, MakeNumber(0))

			Initialise := gen9233
			_ = Initialise
			__e.TailApply(ShenFunc(sym_8p), V3201, Results)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("profile-results"), gen9234)

		gen9238 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3203 := __args[0]
			_ = V3203
			gen9235 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(MakeNumber(0))
				return
			}, 1)
			gen9237 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen9236 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V3203, MakeSymbol("profile"), gen9236)

				return

			}, 0)
			__e.Return(Try(__e, gen9237).Catch(gen9235))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.get-profile"), gen9238)

		gen9240 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3206 := __args[0]
			_ = V3206
			V3207 := __args[1]
			_ = V3207
			gen9239 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			__e.TailApply(ShenFunc(symput), V3206, MakeSymbol("profile"), V3207, gen9239)

			return

		}, 2)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.put-profile"), gen9240)

		return

	}, 0))
}
