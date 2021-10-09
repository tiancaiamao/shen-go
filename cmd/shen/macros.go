package main

import . "github.com/tiancaiamao/shen-go/cora"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp8012 := MakeNative(func(__e *ControlFlow) {
		V748 := __e.Get(1)
		_ = V748
		tmp8013 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			tmp8016 := PrimEqual(V748, Y)

			if True == tmp8016 {
				__e.Return(V748)
				return
			} else {
				tmp8015 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symmacroexpand), Z)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symshen_4walk), tmp8015, Y)
				return

			}

		}, 1)

		tmp8017 := PrimNS3Value(sym_dmacros_d)

		tmp8018 := Call(__e, PrimNS2Value(symshen_4compose), tmp8017, V748)

		__e.TailApply(tmp8013, tmp8018)
		return

	}, 1)

	tmp8019 := Call(__e, PrimNS1Value(symns2_1set), symmacroexpand, tmp8012)

	_ = tmp8019

	tmp8020 := MakeNative(func(__e *ControlFlow) {
		V750 := __e.Get(1)
		_ = V750
		tmp8036 := PrimIsPair(V750)

		var ifres8028 Obj

		if True == tmp8036 {
			tmp8034 := PrimHead(V750)

			tmp8035 := PrimEqual(symerror, tmp8034)

			var ifres8030 Obj

			if True == tmp8035 {
				tmp8032 := PrimTail(V750)

				tmp8033 := PrimIsPair(tmp8032)

				var ifres8031 Obj

				if True == tmp8033 {
					ifres8031 = True

				} else {
					ifres8031 = False

				}

				ifres8030 = ifres8031

			} else {
				ifres8030 = False

			}

			var ifres8029 Obj

			if True == ifres8030 {
				ifres8029 = True

			} else {
				ifres8029 = False

			}

			ifres8028 = ifres8029

		} else {
			ifres8028 = False

		}

		if True == ifres8028 {
			tmp8022 := PrimTail(V750)

			tmp8023 := PrimHead(tmp8022)

			tmp8024 := PrimTail(V750)

			tmp8025 := PrimTail(tmp8024)

			tmp8026 := Call(__e, PrimNS2Value(symshen_4mkstr), tmp8023, tmp8025)

			tmp8027 := PrimCons(tmp8026, Nil)

			__e.Return(PrimCons(symsimple_1error, tmp8027))
			return

		} else {
			__e.Return(V750)
			return
		}

	}, 1)

	tmp8037 := Call(__e, PrimNS1Value(symns2_1set), symshen_4error_1macro, tmp8020)

	_ = tmp8037

	tmp8038 := MakeNative(func(__e *ControlFlow) {
		V752 := __e.Get(1)
		_ = V752
		tmp8076 := PrimIsPair(V752)

		var ifres8068 Obj

		if True == tmp8076 {
			tmp8074 := PrimHead(V752)

			tmp8075 := PrimEqual(symoutput, tmp8074)

			var ifres8070 Obj

			if True == tmp8075 {
				tmp8072 := PrimTail(V752)

				tmp8073 := PrimIsPair(tmp8072)

				var ifres8071 Obj

				if True == tmp8073 {
					ifres8071 = True

				} else {
					ifres8071 = False

				}

				ifres8070 = ifres8071

			} else {
				ifres8070 = False

			}

			var ifres8069 Obj

			if True == ifres8070 {
				ifres8069 = True

			} else {
				ifres8069 = False

			}

			ifres8068 = ifres8069

		} else {
			ifres8068 = False

		}

		if True == ifres8068 {
			tmp8060 := PrimTail(V752)

			tmp8061 := PrimHead(tmp8060)

			tmp8062 := PrimTail(V752)

			tmp8063 := PrimTail(tmp8062)

			tmp8064 := Call(__e, PrimNS2Value(symshen_4mkstr), tmp8061, tmp8063)

			tmp8065 := PrimCons(symstoutput, Nil)

			tmp8066 := PrimCons(tmp8065, Nil)

			tmp8067 := PrimCons(tmp8064, tmp8066)

			__e.Return(PrimCons(symshen_4prhush, tmp8067))
			return

		} else {
			tmp8059 := PrimIsPair(V752)

			var ifres8046 Obj

			if True == tmp8059 {
				tmp8057 := PrimHead(V752)

				tmp8058 := PrimEqual(sympr, tmp8057)

				var ifres8048 Obj

				if True == tmp8058 {
					tmp8055 := PrimTail(V752)

					tmp8056 := PrimIsPair(tmp8055)

					var ifres8050 Obj

					if True == tmp8056 {
						tmp8052 := PrimTail(V752)

						tmp8053 := PrimTail(tmp8052)

						tmp8054 := PrimEqual(Nil, tmp8053)

						var ifres8051 Obj

						if True == tmp8054 {
							ifres8051 = True

						} else {
							ifres8051 = False

						}

						ifres8050 = ifres8051

					} else {
						ifres8050 = False

					}

					var ifres8049 Obj

					if True == ifres8050 {
						ifres8049 = True

					} else {
						ifres8049 = False

					}

					ifres8048 = ifres8049

				} else {
					ifres8048 = False

				}

				var ifres8047 Obj

				if True == ifres8048 {
					ifres8047 = True

				} else {
					ifres8047 = False

				}

				ifres8046 = ifres8047

			} else {
				ifres8046 = False

			}

			if True == ifres8046 {
				tmp8041 := PrimTail(V752)

				tmp8042 := PrimHead(tmp8041)

				tmp8043 := PrimCons(symstoutput, Nil)

				tmp8044 := PrimCons(tmp8043, Nil)

				tmp8045 := PrimCons(tmp8042, tmp8044)

				__e.Return(PrimCons(sympr, tmp8045))
				return

			} else {
				__e.Return(V752)
				return
			}

		}

	}, 1)

	tmp8077 := Call(__e, PrimNS1Value(symns2_1set), symshen_4output_1macro, tmp8038)

	_ = tmp8077

	tmp8078 := MakeNative(func(__e *ControlFlow) {
		V754 := __e.Get(1)
		_ = V754
		tmp8092 := PrimIsPair(V754)

		var ifres8084 Obj

		if True == tmp8092 {
			tmp8090 := PrimHead(V754)

			tmp8091 := PrimEqual(symmake_1string, tmp8090)

			var ifres8086 Obj

			if True == tmp8091 {
				tmp8088 := PrimTail(V754)

				tmp8089 := PrimIsPair(tmp8088)

				var ifres8087 Obj

				if True == tmp8089 {
					ifres8087 = True

				} else {
					ifres8087 = False

				}

				ifres8086 = ifres8087

			} else {
				ifres8086 = False

			}

			var ifres8085 Obj

			if True == ifres8086 {
				ifres8085 = True

			} else {
				ifres8085 = False

			}

			ifres8084 = ifres8085

		} else {
			ifres8084 = False

		}

		if True == ifres8084 {
			tmp8080 := PrimTail(V754)

			tmp8081 := PrimHead(tmp8080)

			tmp8082 := PrimTail(V754)

			tmp8083 := PrimTail(tmp8082)

			__e.TailApply(PrimNS2Value(symshen_4mkstr), tmp8081, tmp8083)
			return

		} else {
			__e.Return(V754)
			return
		}

	}, 1)

	tmp8093 := Call(__e, PrimNS1Value(symns2_1set), symshen_4make_1string_1macro, tmp8078)

	_ = tmp8093

	tmp8094 := MakeNative(func(__e *ControlFlow) {
		V756 := __e.Get(1)
		_ = V756
		tmp8162 := PrimIsPair(V756)

		var ifres8154 Obj

		if True == tmp8162 {
			tmp8160 := PrimHead(V756)

			tmp8161 := PrimEqual(symlineread, tmp8160)

			var ifres8156 Obj

			if True == tmp8161 {
				tmp8158 := PrimTail(V756)

				tmp8159 := PrimEqual(Nil, tmp8158)

				var ifres8157 Obj

				if True == tmp8159 {
					ifres8157 = True

				} else {
					ifres8157 = False

				}

				ifres8156 = ifres8157

			} else {
				ifres8156 = False

			}

			var ifres8155 Obj

			if True == ifres8156 {
				ifres8155 = True

			} else {
				ifres8155 = False

			}

			ifres8154 = ifres8155

		} else {
			ifres8154 = False

		}

		if True == ifres8154 {
			tmp8152 := PrimCons(symstinput, Nil)

			tmp8153 := PrimCons(tmp8152, Nil)

			__e.Return(PrimCons(symlineread, tmp8153))
			return

		} else {
			tmp8151 := PrimIsPair(V756)

			var ifres8143 Obj

			if True == tmp8151 {
				tmp8149 := PrimHead(V756)

				tmp8150 := PrimEqual(syminput, tmp8149)

				var ifres8145 Obj

				if True == tmp8150 {
					tmp8147 := PrimTail(V756)

					tmp8148 := PrimEqual(Nil, tmp8147)

					var ifres8146 Obj

					if True == tmp8148 {
						ifres8146 = True

					} else {
						ifres8146 = False

					}

					ifres8145 = ifres8146

				} else {
					ifres8145 = False

				}

				var ifres8144 Obj

				if True == ifres8145 {
					ifres8144 = True

				} else {
					ifres8144 = False

				}

				ifres8143 = ifres8144

			} else {
				ifres8143 = False

			}

			if True == ifres8143 {
				tmp8141 := PrimCons(symstinput, Nil)

				tmp8142 := PrimCons(tmp8141, Nil)

				__e.Return(PrimCons(syminput, tmp8142))
				return

			} else {
				tmp8140 := PrimIsPair(V756)

				var ifres8132 Obj

				if True == tmp8140 {
					tmp8138 := PrimHead(V756)

					tmp8139 := PrimEqual(symread, tmp8138)

					var ifres8134 Obj

					if True == tmp8139 {
						tmp8136 := PrimTail(V756)

						tmp8137 := PrimEqual(Nil, tmp8136)

						var ifres8135 Obj

						if True == tmp8137 {
							ifres8135 = True

						} else {
							ifres8135 = False

						}

						ifres8134 = ifres8135

					} else {
						ifres8134 = False

					}

					var ifres8133 Obj

					if True == ifres8134 {
						ifres8133 = True

					} else {
						ifres8133 = False

					}

					ifres8132 = ifres8133

				} else {
					ifres8132 = False

				}

				if True == ifres8132 {
					tmp8130 := PrimCons(symstinput, Nil)

					tmp8131 := PrimCons(tmp8130, Nil)

					__e.Return(PrimCons(symread, tmp8131))
					return

				} else {
					tmp8129 := PrimIsPair(V756)

					var ifres8116 Obj

					if True == tmp8129 {
						tmp8127 := PrimHead(V756)

						tmp8128 := PrimEqual(syminput_7, tmp8127)

						var ifres8118 Obj

						if True == tmp8128 {
							tmp8125 := PrimTail(V756)

							tmp8126 := PrimIsPair(tmp8125)

							var ifres8120 Obj

							if True == tmp8126 {
								tmp8122 := PrimTail(V756)

								tmp8123 := PrimTail(tmp8122)

								tmp8124 := PrimEqual(Nil, tmp8123)

								var ifres8121 Obj

								if True == tmp8124 {
									ifres8121 = True

								} else {
									ifres8121 = False

								}

								ifres8120 = ifres8121

							} else {
								ifres8120 = False

							}

							var ifres8119 Obj

							if True == ifres8120 {
								ifres8119 = True

							} else {
								ifres8119 = False

							}

							ifres8118 = ifres8119

						} else {
							ifres8118 = False

						}

						var ifres8117 Obj

						if True == ifres8118 {
							ifres8117 = True

						} else {
							ifres8117 = False

						}

						ifres8116 = ifres8117

					} else {
						ifres8116 = False

					}

					if True == ifres8116 {
						tmp8111 := PrimTail(V756)

						tmp8112 := PrimHead(tmp8111)

						tmp8113 := PrimCons(symstinput, Nil)

						tmp8114 := PrimCons(tmp8113, Nil)

						tmp8115 := PrimCons(tmp8112, tmp8114)

						__e.Return(PrimCons(syminput_7, tmp8115))
						return

					} else {
						tmp8110 := PrimIsPair(V756)

						var ifres8102 Obj

						if True == tmp8110 {
							tmp8108 := PrimHead(V756)

							tmp8109 := PrimEqual(symread_1byte, tmp8108)

							var ifres8104 Obj

							if True == tmp8109 {
								tmp8106 := PrimTail(V756)

								tmp8107 := PrimEqual(Nil, tmp8106)

								var ifres8105 Obj

								if True == tmp8107 {
									ifres8105 = True

								} else {
									ifres8105 = False

								}

								ifres8104 = ifres8105

							} else {
								ifres8104 = False

							}

							var ifres8103 Obj

							if True == ifres8104 {
								ifres8103 = True

							} else {
								ifres8103 = False

							}

							ifres8102 = ifres8103

						} else {
							ifres8102 = False

						}

						if True == ifres8102 {
							tmp8100 := PrimCons(symstinput, Nil)

							tmp8101 := PrimCons(tmp8100, Nil)

							__e.Return(PrimCons(symread_1byte, tmp8101))
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

	tmp8163 := Call(__e, PrimNS1Value(symns2_1set), symshen_4input_1macro, tmp8094)

	_ = tmp8163

	tmp8164 := MakeNative(func(__e *ControlFlow) {
		V759 := __e.Get(1)
		_ = V759
		V760 := __e.Get(2)
		_ = V760
		tmp8171 := PrimEqual(Nil, V759)

		if True == tmp8171 {
			__e.Return(V760)
			return
		} else {
			tmp8170 := PrimIsPair(V759)

			if True == tmp8170 {
				tmp8167 := PrimTail(V759)

				tmp8168 := PrimHead(V759)

				tmp8169 := Call(__e, tmp8168, V760)

				__e.TailApply(PrimNS2Value(symshen_4compose), tmp8167, tmp8169)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4compose)
				return
			}

		}

	}, 2)

	tmp8172 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compose, tmp8164)

	_ = tmp8172

	tmp8173 := MakeNative(func(__e *ControlFlow) {
		V762 := __e.Get(1)
		_ = V762
		tmp8216 := PrimIsPair(V762)

		var ifres8197 Obj

		if True == tmp8216 {
			tmp8214 := PrimHead(V762)

			tmp8215 := PrimEqual(symcompile, tmp8214)

			var ifres8199 Obj

			if True == tmp8215 {
				tmp8212 := PrimTail(V762)

				tmp8213 := PrimIsPair(tmp8212)

				var ifres8201 Obj

				if True == tmp8213 {
					tmp8209 := PrimTail(V762)

					tmp8210 := PrimTail(tmp8209)

					tmp8211 := PrimIsPair(tmp8210)

					var ifres8203 Obj

					if True == tmp8211 {
						tmp8205 := PrimTail(V762)

						tmp8206 := PrimTail(tmp8205)

						tmp8207 := PrimTail(tmp8206)

						tmp8208 := PrimEqual(Nil, tmp8207)

						var ifres8204 Obj

						if True == tmp8208 {
							ifres8204 = True

						} else {
							ifres8204 = False

						}

						ifres8203 = ifres8204

					} else {
						ifres8203 = False

					}

					var ifres8202 Obj

					if True == ifres8203 {
						ifres8202 = True

					} else {
						ifres8202 = False

					}

					ifres8201 = ifres8202

				} else {
					ifres8201 = False

				}

				var ifres8200 Obj

				if True == ifres8201 {
					ifres8200 = True

				} else {
					ifres8200 = False

				}

				ifres8199 = ifres8200

			} else {
				ifres8199 = False

			}

			var ifres8198 Obj

			if True == ifres8199 {
				ifres8198 = True

			} else {
				ifres8198 = False

			}

			ifres8197 = ifres8198

		} else {
			ifres8197 = False

		}

		if True == ifres8197 {
			tmp8175 := PrimTail(V762)

			tmp8176 := PrimHead(tmp8175)

			tmp8177 := PrimTail(V762)

			tmp8178 := PrimTail(tmp8177)

			tmp8179 := PrimHead(tmp8178)

			tmp8180 := PrimCons(symE, Nil)

			tmp8181 := PrimCons(symcons_2, tmp8180)

			tmp8182 := PrimCons(symE, Nil)

			tmp8183 := PrimCons(MakeString("parse error here: ~S~%"), tmp8182)

			tmp8184 := PrimCons(symerror, tmp8183)

			tmp8185 := PrimCons(MakeString("parse error~%"), Nil)

			tmp8186 := PrimCons(symerror, tmp8185)

			tmp8187 := PrimCons(tmp8186, Nil)

			tmp8188 := PrimCons(tmp8184, tmp8187)

			tmp8189 := PrimCons(tmp8181, tmp8188)

			tmp8190 := PrimCons(symif, tmp8189)

			tmp8191 := PrimCons(tmp8190, Nil)

			tmp8192 := PrimCons(symE, tmp8191)

			tmp8193 := PrimCons(symlambda, tmp8192)

			tmp8194 := PrimCons(tmp8193, Nil)

			tmp8195 := PrimCons(tmp8179, tmp8194)

			tmp8196 := PrimCons(tmp8176, tmp8195)

			__e.Return(PrimCons(symcompile, tmp8196))
			return

		} else {
			__e.Return(V762)
			return
		}

	}, 1)

	tmp8217 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile_1macro, tmp8173)

	_ = tmp8217

	tmp8218 := MakeNative(func(__e *ControlFlow) {
		V764 := __e.Get(1)
		_ = V764
		tmp8240 := PrimIsPair(V764)

		var ifres8236 Obj

		if True == tmp8240 {
			tmp8238 := PrimHead(V764)

			tmp8239 := PrimEqual(symprolog_2, tmp8238)

			var ifres8237 Obj

			if True == tmp8239 {
				ifres8237 = True

			} else {
				ifres8237 = False

			}

			ifres8236 = ifres8237

		} else {
			ifres8236 = False

		}

		if True == ifres8236 {
			tmp8220 := PrimCons(symshen_4start_1new_1prolog_1process, Nil)

			tmp8221 := MakeNative(func(__e *ControlFlow) {
				Calls := __e.Get(1)
				_ = Calls
				tmp8222 := MakeNative(func(__e *ControlFlow) {
					Vs := __e.Get(1)
					_ = Vs
					tmp8223 := MakeNative(func(__e *ControlFlow) {
						External := __e.Get(1)
						_ = External
						tmp8224 := MakeNative(func(__e *ControlFlow) {
							PrologVs := __e.Get(1)
							_ = PrologVs
							__e.TailApply(PrimNS2Value(symshen_4locally_1bind_1prolog_1vs), symNPP, PrologVs, Calls)
							return
						}, 1)

						tmp8225 := Call(__e, PrimNS2Value(symdifference), Vs, External)

						__e.TailApply(tmp8224, tmp8225)
						return

					}, 1)

					tmp8226 := PrimTail(V764)

					tmp8227 := Call(__e, PrimNS2Value(symshen_4externally_1bound), tmp8226)

					__e.TailApply(tmp8223, tmp8227)
					return

				}, 1)

				tmp8228 := PrimTail(V764)

				tmp8229 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp8228)

				__e.TailApply(tmp8222, tmp8229)
				return

			}, 1)

			tmp8230 := PrimTail(V764)

			tmp8231 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), symNPP, tmp8230)

			tmp8232 := Call(__e, tmp8221, tmp8231)

			tmp8233 := PrimCons(tmp8232, Nil)

			tmp8234 := PrimCons(tmp8220, tmp8233)

			tmp8235 := PrimCons(symNPP, tmp8234)

			__e.Return(PrimCons(symlet, tmp8235))
			return

		} else {
			__e.Return(V764)
			return
		}

	}, 1)

	tmp8241 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1macro, tmp8218)

	_ = tmp8241

	tmp8242 := MakeNative(func(__e *ControlFlow) {
		V770 := __e.Get(1)
		_ = V770
		tmp8263 := PrimIsPair(V770)

		var ifres8250 Obj

		if True == tmp8263 {
			tmp8261 := PrimHead(V770)

			tmp8262 := PrimEqual(symreceive, tmp8261)

			var ifres8252 Obj

			if True == tmp8262 {
				tmp8259 := PrimTail(V770)

				tmp8260 := PrimIsPair(tmp8259)

				var ifres8254 Obj

				if True == tmp8260 {
					tmp8256 := PrimTail(V770)

					tmp8257 := PrimTail(tmp8256)

					tmp8258 := PrimEqual(Nil, tmp8257)

					var ifres8255 Obj

					if True == tmp8258 {
						ifres8255 = True

					} else {
						ifres8255 = False

					}

					ifres8254 = ifres8255

				} else {
					ifres8254 = False

				}

				var ifres8253 Obj

				if True == ifres8254 {
					ifres8253 = True

				} else {
					ifres8253 = False

				}

				ifres8252 = ifres8253

			} else {
				ifres8252 = False

			}

			var ifres8251 Obj

			if True == ifres8252 {
				ifres8251 = True

			} else {
				ifres8251 = False

			}

			ifres8250 = ifres8251

		} else {
			ifres8250 = False

		}

		if True == ifres8250 {
			__e.Return(PrimTail(V770))
			return
		} else {
			tmp8249 := PrimIsPair(V770)

			if True == tmp8249 {
				tmp8245 := PrimHead(V770)

				tmp8246 := Call(__e, PrimNS2Value(symshen_4externally_1bound), tmp8245)

				tmp8247 := PrimTail(V770)

				tmp8248 := Call(__e, PrimNS2Value(symshen_4externally_1bound), tmp8247)

				__e.TailApply(PrimNS2Value(symunion), tmp8246, tmp8248)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp8264 := Call(__e, PrimNS1Value(symns2_1set), symshen_4externally_1bound, tmp8242)

	_ = tmp8264

	tmp8265 := MakeNative(func(__e *ControlFlow) {
		V788 := __e.Get(1)
		_ = V788
		V789 := __e.Get(2)
		_ = V789
		V790 := __e.Get(3)
		_ = V790
		tmp8277 := PrimEqual(Nil, V789)

		if True == tmp8277 {
			__e.Return(V790)
			return
		} else {
			tmp8276 := PrimIsPair(V789)

			if True == tmp8276 {
				tmp8268 := PrimHead(V789)

				tmp8269 := PrimCons(V788, Nil)

				tmp8270 := PrimCons(symshen_4newpv, tmp8269)

				tmp8271 := PrimTail(V789)

				tmp8272 := Call(__e, PrimNS2Value(symshen_4locally_1bind_1prolog_1vs), V788, tmp8271, V790)

				tmp8273 := PrimCons(tmp8272, Nil)

				tmp8274 := PrimCons(tmp8270, tmp8273)

				tmp8275 := PrimCons(tmp8268, tmp8274)

				__e.Return(PrimCons(symlet, tmp8275))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error inp locally-bind-prolog-vs")))
				return
			}

		}

	}, 3)

	tmp8278 := Call(__e, PrimNS1Value(symns2_1set), symshen_4locally_1bind_1prolog_1vs, tmp8265)

	_ = tmp8278

	tmp8279 := MakeNative(func(__e *ControlFlow) {
		V803 := __e.Get(1)
		_ = V803
		V804 := __e.Get(2)
		_ = V804
		tmp8482 := PrimEqual(Nil, V804)

		if True == tmp8482 {
			__e.Return(True)
			return
		} else {
			tmp8481 := PrimIsPair(V804)

			var ifres8477 Obj

			if True == tmp8481 {
				tmp8479 := PrimHead(V804)

				tmp8480 := PrimEqual(sym_b, tmp8479)

				var ifres8478 Obj

				if True == tmp8480 {
					ifres8478 = True

				} else {
					ifres8478 = False

				}

				ifres8477 = ifres8478

			} else {
				ifres8477 = False

			}

			if True == ifres8477 {
				tmp8470 := PrimTail(V804)

				tmp8471 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8470)

				tmp8472 := PrimCons(tmp8471, Nil)

				tmp8473 := PrimCons(symfreeze, tmp8472)

				tmp8474 := PrimCons(tmp8473, Nil)

				tmp8475 := PrimCons(V803, tmp8474)

				tmp8476 := PrimCons(False, tmp8475)

				__e.Return(PrimCons(symcut, tmp8476))
				return

			} else {
				tmp8469 := PrimIsPair(V804)

				var ifres8449 Obj

				if True == tmp8469 {
					tmp8467 := PrimHead(V804)

					tmp8468 := PrimIsPair(tmp8467)

					var ifres8451 Obj

					if True == tmp8468 {
						tmp8464 := PrimHead(V804)

						tmp8465 := PrimHead(tmp8464)

						tmp8466 := PrimEqual(symwhen, tmp8465)

						var ifres8453 Obj

						if True == tmp8466 {
							tmp8461 := PrimHead(V804)

							tmp8462 := PrimTail(tmp8461)

							tmp8463 := PrimIsPair(tmp8462)

							var ifres8455 Obj

							if True == tmp8463 {
								tmp8457 := PrimHead(V804)

								tmp8458 := PrimTail(tmp8457)

								tmp8459 := PrimTail(tmp8458)

								tmp8460 := PrimEqual(Nil, tmp8459)

								var ifres8456 Obj

								if True == tmp8460 {
									ifres8456 = True

								} else {
									ifres8456 = False

								}

								ifres8455 = ifres8456

							} else {
								ifres8455 = False

							}

							var ifres8454 Obj

							if True == ifres8455 {
								ifres8454 = True

							} else {
								ifres8454 = False

							}

							ifres8453 = ifres8454

						} else {
							ifres8453 = False

						}

						var ifres8452 Obj

						if True == ifres8453 {
							ifres8452 = True

						} else {
							ifres8452 = False

						}

						ifres8451 = ifres8452

					} else {
						ifres8451 = False

					}

					var ifres8450 Obj

					if True == ifres8451 {
						ifres8450 = True

					} else {
						ifres8450 = False

					}

					ifres8449 = ifres8450

				} else {
					ifres8449 = False

				}

				if True == ifres8449 {
					tmp8438 := PrimHead(V804)

					tmp8439 := PrimTail(tmp8438)

					tmp8440 := PrimHead(tmp8439)

					tmp8441 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp8440, V803)

					tmp8442 := PrimTail(V804)

					tmp8443 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8442)

					tmp8444 := PrimCons(tmp8443, Nil)

					tmp8445 := PrimCons(symfreeze, tmp8444)

					tmp8446 := PrimCons(tmp8445, Nil)

					tmp8447 := PrimCons(V803, tmp8446)

					tmp8448 := PrimCons(tmp8441, tmp8447)

					__e.Return(PrimCons(symfwhen, tmp8448))
					return

				} else {
					tmp8437 := PrimIsPair(V804)

					var ifres8410 Obj

					if True == tmp8437 {
						tmp8435 := PrimHead(V804)

						tmp8436 := PrimIsPair(tmp8435)

						var ifres8412 Obj

						if True == tmp8436 {
							tmp8432 := PrimHead(V804)

							tmp8433 := PrimHead(tmp8432)

							tmp8434 := PrimEqual(symis, tmp8433)

							var ifres8414 Obj

							if True == tmp8434 {
								tmp8429 := PrimHead(V804)

								tmp8430 := PrimTail(tmp8429)

								tmp8431 := PrimIsPair(tmp8430)

								var ifres8416 Obj

								if True == tmp8431 {
									tmp8425 := PrimHead(V804)

									tmp8426 := PrimTail(tmp8425)

									tmp8427 := PrimTail(tmp8426)

									tmp8428 := PrimIsPair(tmp8427)

									var ifres8418 Obj

									if True == tmp8428 {
										tmp8420 := PrimHead(V804)

										tmp8421 := PrimTail(tmp8420)

										tmp8422 := PrimTail(tmp8421)

										tmp8423 := PrimTail(tmp8422)

										tmp8424 := PrimEqual(Nil, tmp8423)

										var ifres8419 Obj

										if True == tmp8424 {
											ifres8419 = True

										} else {
											ifres8419 = False

										}

										ifres8418 = ifres8419

									} else {
										ifres8418 = False

									}

									var ifres8417 Obj

									if True == ifres8418 {
										ifres8417 = True

									} else {
										ifres8417 = False

									}

									ifres8416 = ifres8417

								} else {
									ifres8416 = False

								}

								var ifres8415 Obj

								if True == ifres8416 {
									ifres8415 = True

								} else {
									ifres8415 = False

								}

								ifres8414 = ifres8415

							} else {
								ifres8414 = False

							}

							var ifres8413 Obj

							if True == ifres8414 {
								ifres8413 = True

							} else {
								ifres8413 = False

							}

							ifres8412 = ifres8413

						} else {
							ifres8412 = False

						}

						var ifres8411 Obj

						if True == ifres8412 {
							ifres8411 = True

						} else {
							ifres8411 = False

						}

						ifres8410 = ifres8411

					} else {
						ifres8410 = False

					}

					if True == ifres8410 {
						tmp8394 := PrimHead(V804)

						tmp8395 := PrimTail(tmp8394)

						tmp8396 := PrimHead(tmp8395)

						tmp8397 := PrimHead(V804)

						tmp8398 := PrimTail(tmp8397)

						tmp8399 := PrimTail(tmp8398)

						tmp8400 := PrimHead(tmp8399)

						tmp8401 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp8400, V803)

						tmp8402 := PrimTail(V804)

						tmp8403 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8402)

						tmp8404 := PrimCons(tmp8403, Nil)

						tmp8405 := PrimCons(symfreeze, tmp8404)

						tmp8406 := PrimCons(tmp8405, Nil)

						tmp8407 := PrimCons(V803, tmp8406)

						tmp8408 := PrimCons(tmp8401, tmp8407)

						tmp8409 := PrimCons(tmp8396, tmp8408)

						__e.Return(PrimCons(symbind, tmp8409))
						return

					} else {
						tmp8393 := PrimIsPair(V804)

						var ifres8373 Obj

						if True == tmp8393 {
							tmp8391 := PrimHead(V804)

							tmp8392 := PrimIsPair(tmp8391)

							var ifres8375 Obj

							if True == tmp8392 {
								tmp8388 := PrimHead(V804)

								tmp8389 := PrimHead(tmp8388)

								tmp8390 := PrimEqual(symreceive, tmp8389)

								var ifres8377 Obj

								if True == tmp8390 {
									tmp8385 := PrimHead(V804)

									tmp8386 := PrimTail(tmp8385)

									tmp8387 := PrimIsPair(tmp8386)

									var ifres8379 Obj

									if True == tmp8387 {
										tmp8381 := PrimHead(V804)

										tmp8382 := PrimTail(tmp8381)

										tmp8383 := PrimTail(tmp8382)

										tmp8384 := PrimEqual(Nil, tmp8383)

										var ifres8380 Obj

										if True == tmp8384 {
											ifres8380 = True

										} else {
											ifres8380 = False

										}

										ifres8379 = ifres8380

									} else {
										ifres8379 = False

									}

									var ifres8378 Obj

									if True == ifres8379 {
										ifres8378 = True

									} else {
										ifres8378 = False

									}

									ifres8377 = ifres8378

								} else {
									ifres8377 = False

								}

								var ifres8376 Obj

								if True == ifres8377 {
									ifres8376 = True

								} else {
									ifres8376 = False

								}

								ifres8375 = ifres8376

							} else {
								ifres8375 = False

							}

							var ifres8374 Obj

							if True == ifres8375 {
								ifres8374 = True

							} else {
								ifres8374 = False

							}

							ifres8373 = ifres8374

						} else {
							ifres8373 = False

						}

						if True == ifres8373 {
							tmp8372 := PrimTail(V804)

							__e.TailApply(PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8372)
							return

						} else {
							tmp8371 := PrimIsPair(V804)

							var ifres8344 Obj

							if True == tmp8371 {
								tmp8369 := PrimHead(V804)

								tmp8370 := PrimIsPair(tmp8369)

								var ifres8346 Obj

								if True == tmp8370 {
									tmp8366 := PrimHead(V804)

									tmp8367 := PrimHead(tmp8366)

									tmp8368 := PrimEqual(symbind, tmp8367)

									var ifres8348 Obj

									if True == tmp8368 {
										tmp8363 := PrimHead(V804)

										tmp8364 := PrimTail(tmp8363)

										tmp8365 := PrimIsPair(tmp8364)

										var ifres8350 Obj

										if True == tmp8365 {
											tmp8359 := PrimHead(V804)

											tmp8360 := PrimTail(tmp8359)

											tmp8361 := PrimTail(tmp8360)

											tmp8362 := PrimIsPair(tmp8361)

											var ifres8352 Obj

											if True == tmp8362 {
												tmp8354 := PrimHead(V804)

												tmp8355 := PrimTail(tmp8354)

												tmp8356 := PrimTail(tmp8355)

												tmp8357 := PrimTail(tmp8356)

												tmp8358 := PrimEqual(Nil, tmp8357)

												var ifres8353 Obj

												if True == tmp8358 {
													ifres8353 = True

												} else {
													ifres8353 = False

												}

												ifres8352 = ifres8353

											} else {
												ifres8352 = False

											}

											var ifres8351 Obj

											if True == ifres8352 {
												ifres8351 = True

											} else {
												ifres8351 = False

											}

											ifres8350 = ifres8351

										} else {
											ifres8350 = False

										}

										var ifres8349 Obj

										if True == ifres8350 {
											ifres8349 = True

										} else {
											ifres8349 = False

										}

										ifres8348 = ifres8349

									} else {
										ifres8348 = False

									}

									var ifres8347 Obj

									if True == ifres8348 {
										ifres8347 = True

									} else {
										ifres8347 = False

									}

									ifres8346 = ifres8347

								} else {
									ifres8346 = False

								}

								var ifres8345 Obj

								if True == ifres8346 {
									ifres8345 = True

								} else {
									ifres8345 = False

								}

								ifres8344 = ifres8345

							} else {
								ifres8344 = False

							}

							if True == ifres8344 {
								tmp8328 := PrimHead(V804)

								tmp8329 := PrimTail(tmp8328)

								tmp8330 := PrimHead(tmp8329)

								tmp8331 := PrimHead(V804)

								tmp8332 := PrimTail(tmp8331)

								tmp8333 := PrimTail(tmp8332)

								tmp8334 := PrimHead(tmp8333)

								tmp8335 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp8334, V803)

								tmp8336 := PrimTail(V804)

								tmp8337 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8336)

								tmp8338 := PrimCons(tmp8337, Nil)

								tmp8339 := PrimCons(symfreeze, tmp8338)

								tmp8340 := PrimCons(tmp8339, Nil)

								tmp8341 := PrimCons(V803, tmp8340)

								tmp8342 := PrimCons(tmp8335, tmp8341)

								tmp8343 := PrimCons(tmp8330, tmp8342)

								__e.Return(PrimCons(symbind, tmp8343))
								return

							} else {
								tmp8327 := PrimIsPair(V804)

								var ifres8307 Obj

								if True == tmp8327 {
									tmp8325 := PrimHead(V804)

									tmp8326 := PrimIsPair(tmp8325)

									var ifres8309 Obj

									if True == tmp8326 {
										tmp8322 := PrimHead(V804)

										tmp8323 := PrimHead(tmp8322)

										tmp8324 := PrimEqual(symfwhen, tmp8323)

										var ifres8311 Obj

										if True == tmp8324 {
											tmp8319 := PrimHead(V804)

											tmp8320 := PrimTail(tmp8319)

											tmp8321 := PrimIsPair(tmp8320)

											var ifres8313 Obj

											if True == tmp8321 {
												tmp8315 := PrimHead(V804)

												tmp8316 := PrimTail(tmp8315)

												tmp8317 := PrimTail(tmp8316)

												tmp8318 := PrimEqual(Nil, tmp8317)

												var ifres8314 Obj

												if True == tmp8318 {
													ifres8314 = True

												} else {
													ifres8314 = False

												}

												ifres8313 = ifres8314

											} else {
												ifres8313 = False

											}

											var ifres8312 Obj

											if True == ifres8313 {
												ifres8312 = True

											} else {
												ifres8312 = False

											}

											ifres8311 = ifres8312

										} else {
											ifres8311 = False

										}

										var ifres8310 Obj

										if True == ifres8311 {
											ifres8310 = True

										} else {
											ifres8310 = False

										}

										ifres8309 = ifres8310

									} else {
										ifres8309 = False

									}

									var ifres8308 Obj

									if True == ifres8309 {
										ifres8308 = True

									} else {
										ifres8308 = False

									}

									ifres8307 = ifres8308

								} else {
									ifres8307 = False

								}

								if True == ifres8307 {
									tmp8296 := PrimHead(V804)

									tmp8297 := PrimTail(tmp8296)

									tmp8298 := PrimHead(tmp8297)

									tmp8299 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp8298, V803)

									tmp8300 := PrimTail(V804)

									tmp8301 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8300)

									tmp8302 := PrimCons(tmp8301, Nil)

									tmp8303 := PrimCons(symfreeze, tmp8302)

									tmp8304 := PrimCons(tmp8303, Nil)

									tmp8305 := PrimCons(V803, tmp8304)

									tmp8306 := PrimCons(tmp8299, tmp8305)

									__e.Return(PrimCons(symfwhen, tmp8306))
									return

								} else {
									tmp8295 := PrimIsPair(V804)

									if True == tmp8295 {
										tmp8288 := PrimHead(V804)

										tmp8289 := PrimTail(V804)

										tmp8290 := Call(__e, PrimNS2Value(symshen_4bld_1prolog_1call), V803, tmp8289)

										tmp8291 := PrimCons(tmp8290, Nil)

										tmp8292 := PrimCons(symfreeze, tmp8291)

										tmp8293 := PrimCons(tmp8292, Nil)

										tmp8294 := PrimCons(V803, tmp8293)

										__e.TailApply(PrimNS2Value(symappend), tmp8288, tmp8294)
										return

									} else {
										__e.Return(PrimSimpleError(MakeString("implementation error in bld-prolog-call")))
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

	tmp8483 := Call(__e, PrimNS1Value(symns2_1set), symshen_4bld_1prolog_1call, tmp8279)

	_ = tmp8483

	tmp8484 := MakeNative(func(__e *ControlFlow) {
		V806 := __e.Get(1)
		_ = V806
		tmp8499 := PrimIsPair(V806)

		var ifres8491 Obj

		if True == tmp8499 {
			tmp8497 := PrimHead(V806)

			tmp8498 := PrimEqual(symdefprolog, tmp8497)

			var ifres8493 Obj

			if True == tmp8498 {
				tmp8495 := PrimTail(V806)

				tmp8496 := PrimIsPair(tmp8495)

				var ifres8494 Obj

				if True == tmp8496 {
					ifres8494 = True

				} else {
					ifres8494 = False

				}

				ifres8493 = ifres8494

			} else {
				ifres8493 = False

			}

			var ifres8492 Obj

			if True == ifres8493 {
				ifres8492 = True

			} else {
				ifres8492 = False

			}

			ifres8491 = ifres8492

		} else {
			ifres8491 = False

		}

		if True == ifres8491 {
			tmp8486 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimNS2Value(symshen_4_5defprolog_6), Y)
				return
			}, 1)

			tmp8487 := PrimTail(V806)

			tmp8488 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp8489 := PrimTail(V806)

				tmp8490 := PrimHead(tmp8489)

				__e.TailApply(PrimNS2Value(symshen_4prolog_1error), tmp8490, Y)
				return

			}, 1)

			__e.TailApply(PrimNS2Value(symcompile), tmp8486, tmp8487, tmp8488)
			return

		} else {
			__e.Return(V806)
			return
		}

	}, 1)

	tmp8500 := Call(__e, PrimNS1Value(symns2_1set), symshen_4defprolog_1macro, tmp8484)

	_ = tmp8500

	tmp8501 := MakeNative(func(__e *ControlFlow) {
		V808 := __e.Get(1)
		_ = V808
		tmp8530 := PrimIsPair(V808)

		var ifres8522 Obj

		if True == tmp8530 {
			tmp8528 := PrimHead(V808)

			tmp8529 := PrimEqual(symdatatype, tmp8528)

			var ifres8524 Obj

			if True == tmp8529 {
				tmp8526 := PrimTail(V808)

				tmp8527 := PrimIsPair(tmp8526)

				var ifres8525 Obj

				if True == tmp8527 {
					ifres8525 = True

				} else {
					ifres8525 = False

				}

				ifres8524 = ifres8525

			} else {
				ifres8524 = False

			}

			var ifres8523 Obj

			if True == ifres8524 {
				ifres8523 = True

			} else {
				ifres8523 = False

			}

			ifres8522 = ifres8523

		} else {
			ifres8522 = False

		}

		if True == ifres8522 {
			tmp8503 := PrimTail(V808)

			tmp8504 := PrimHead(tmp8503)

			tmp8505 := Call(__e, PrimNS2Value(symshen_4intern_1type), tmp8504)

			tmp8506 := PrimCons(symX, Nil)

			tmp8507 := PrimCons(symshen_4_5datatype_1rules_6, tmp8506)

			tmp8508 := PrimCons(tmp8507, Nil)

			tmp8509 := PrimCons(symX, tmp8508)

			tmp8510 := PrimCons(symlambda, tmp8509)

			tmp8511 := PrimTail(V808)

			tmp8512 := PrimTail(tmp8511)

			tmp8513 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp8512)

			tmp8514 := PrimCons(symshen_4datatype_1error, Nil)

			tmp8515 := PrimCons(symfunction, tmp8514)

			tmp8516 := PrimCons(tmp8515, Nil)

			tmp8517 := PrimCons(tmp8513, tmp8516)

			tmp8518 := PrimCons(tmp8510, tmp8517)

			tmp8519 := PrimCons(symcompile, tmp8518)

			tmp8520 := PrimCons(tmp8519, Nil)

			tmp8521 := PrimCons(tmp8505, tmp8520)

			__e.Return(PrimCons(symshen_4process_1datatype, tmp8521))
			return

		} else {
			__e.Return(V808)
			return
		}

	}, 1)

	tmp8531 := Call(__e, PrimNS1Value(symns2_1set), symshen_4datatype_1macro, tmp8501)

	_ = tmp8531

	tmp8532 := MakeNative(func(__e *ControlFlow) {
		V810 := __e.Get(1)
		_ = V810
		tmp8533 := PrimStr(V810)

		tmp8534 := PrimStringConcat(tmp8533, MakeString("#type"))

		__e.Return(PrimIntern(tmp8534))
		return

	}, 1)

	tmp8535 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intern_1type, tmp8532)

	_ = tmp8535

	tmp8536 := MakeNative(func(__e *ControlFlow) {
		V812 := __e.Get(1)
		_ = V812
		tmp8602 := PrimIsPair(V812)

		var ifres8583 Obj

		if True == tmp8602 {
			tmp8600 := PrimHead(V812)

			tmp8601 := PrimEqual(sym_8s, tmp8600)

			var ifres8585 Obj

			if True == tmp8601 {
				tmp8598 := PrimTail(V812)

				tmp8599 := PrimIsPair(tmp8598)

				var ifres8587 Obj

				if True == tmp8599 {
					tmp8595 := PrimTail(V812)

					tmp8596 := PrimTail(tmp8595)

					tmp8597 := PrimIsPair(tmp8596)

					var ifres8589 Obj

					if True == tmp8597 {
						tmp8591 := PrimTail(V812)

						tmp8592 := PrimTail(tmp8591)

						tmp8593 := PrimTail(tmp8592)

						tmp8594 := PrimIsPair(tmp8593)

						var ifres8590 Obj

						if True == tmp8594 {
							ifres8590 = True

						} else {
							ifres8590 = False

						}

						ifres8589 = ifres8590

					} else {
						ifres8589 = False

					}

					var ifres8588 Obj

					if True == ifres8589 {
						ifres8588 = True

					} else {
						ifres8588 = False

					}

					ifres8587 = ifres8588

				} else {
					ifres8587 = False

				}

				var ifres8586 Obj

				if True == ifres8587 {
					ifres8586 = True

				} else {
					ifres8586 = False

				}

				ifres8585 = ifres8586

			} else {
				ifres8585 = False

			}

			var ifres8584 Obj

			if True == ifres8585 {
				ifres8584 = True

			} else {
				ifres8584 = False

			}

			ifres8583 = ifres8584

		} else {
			ifres8583 = False

		}

		if True == ifres8583 {
			tmp8575 := PrimTail(V812)

			tmp8576 := PrimHead(tmp8575)

			tmp8577 := PrimTail(V812)

			tmp8578 := PrimTail(tmp8577)

			tmp8579 := PrimCons(sym_8s, tmp8578)

			tmp8580 := Call(__e, PrimNS2Value(symshen_4_8s_1macro), tmp8579)

			tmp8581 := PrimCons(tmp8580, Nil)

			tmp8582 := PrimCons(tmp8576, tmp8581)

			__e.Return(PrimCons(sym_8s, tmp8582))
			return

		} else {
			tmp8574 := PrimIsPair(V812)

			var ifres8550 Obj

			if True == tmp8574 {
				tmp8572 := PrimHead(V812)

				tmp8573 := PrimEqual(sym_8s, tmp8572)

				var ifres8552 Obj

				if True == tmp8573 {
					tmp8570 := PrimTail(V812)

					tmp8571 := PrimIsPair(tmp8570)

					var ifres8554 Obj

					if True == tmp8571 {
						tmp8567 := PrimTail(V812)

						tmp8568 := PrimTail(tmp8567)

						tmp8569 := PrimIsPair(tmp8568)

						var ifres8556 Obj

						if True == tmp8569 {
							tmp8563 := PrimTail(V812)

							tmp8564 := PrimTail(tmp8563)

							tmp8565 := PrimTail(tmp8564)

							tmp8566 := PrimEqual(Nil, tmp8565)

							var ifres8558 Obj

							if True == tmp8566 {
								tmp8560 := PrimTail(V812)

								tmp8561 := PrimHead(tmp8560)

								tmp8562 := PrimIsString(tmp8561)

								var ifres8559 Obj

								if True == tmp8562 {
									ifres8559 = True

								} else {
									ifres8559 = False

								}

								ifres8558 = ifres8559

							} else {
								ifres8558 = False

							}

							var ifres8557 Obj

							if True == ifres8558 {
								ifres8557 = True

							} else {
								ifres8557 = False

							}

							ifres8556 = ifres8557

						} else {
							ifres8556 = False

						}

						var ifres8555 Obj

						if True == ifres8556 {
							ifres8555 = True

						} else {
							ifres8555 = False

						}

						ifres8554 = ifres8555

					} else {
						ifres8554 = False

					}

					var ifres8553 Obj

					if True == ifres8554 {
						ifres8553 = True

					} else {
						ifres8553 = False

					}

					ifres8552 = ifres8553

				} else {
					ifres8552 = False

				}

				var ifres8551 Obj

				if True == ifres8552 {
					ifres8551 = True

				} else {
					ifres8551 = False

				}

				ifres8550 = ifres8551

			} else {
				ifres8550 = False

			}

			if True == ifres8550 {
				tmp8539 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					tmp8545 := Call(__e, PrimNS2Value(symlength), E)

					tmp8546 := PrimGreatThan(tmp8545, MakeNumber(1))

					if True == tmp8546 {
						tmp8541 := PrimTail(V812)

						tmp8542 := PrimTail(tmp8541)

						tmp8543 := Call(__e, PrimNS2Value(symappend), E, tmp8542)

						tmp8544 := PrimCons(sym_8s, tmp8543)

						__e.TailApply(PrimNS2Value(symshen_4_8s_1macro), tmp8544)
						return

					} else {
						__e.Return(V812)
						return
					}

				}, 1)

				tmp8547 := PrimTail(V812)

				tmp8548 := PrimHead(tmp8547)

				tmp8549 := Call(__e, PrimNS2Value(symexplode), tmp8548)

				__e.TailApply(tmp8539, tmp8549)
				return

			} else {
				__e.Return(V812)
				return
			}

		}

	}, 1)

	tmp8603 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_8s_1macro, tmp8536)

	_ = tmp8603

	tmp8604 := MakeNative(func(__e *ControlFlow) {
		V814 := __e.Get(1)
		_ = V814
		tmp8614 := PrimIsPair(V814)

		var ifres8610 Obj

		if True == tmp8614 {
			tmp8612 := PrimHead(V814)

			tmp8613 := PrimEqual(symsynonyms, tmp8612)

			var ifres8611 Obj

			if True == tmp8613 {
				ifres8611 = True

			} else {
				ifres8611 = False

			}

			ifres8610 = ifres8611

		} else {
			ifres8610 = False

		}

		if True == ifres8610 {
			tmp8606 := PrimTail(V814)

			tmp8607 := Call(__e, PrimNS2Value(symshen_4curry_1synonyms), tmp8606)

			tmp8608 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp8607)

			tmp8609 := PrimCons(tmp8608, Nil)

			__e.Return(PrimCons(symshen_4synonyms_1help, tmp8609))
			return

		} else {
			__e.Return(V814)
			return
		}

	}, 1)

	tmp8615 := Call(__e, PrimNS1Value(symns2_1set), symshen_4synonyms_1macro, tmp8604)

	_ = tmp8615

	tmp8616 := MakeNative(func(__e *ControlFlow) {
		V816 := __e.Get(1)
		_ = V816
		tmp8617 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4curry_1type), X)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symmap), tmp8617, V816)
		return

	}, 1)

	tmp8618 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1synonyms, tmp8616)

	_ = tmp8618

	tmp8619 := MakeNative(func(__e *ControlFlow) {
		V818 := __e.Get(1)
		_ = V818
		tmp8630 := PrimIsPair(V818)

		var ifres8622 Obj

		if True == tmp8630 {
			tmp8628 := PrimHead(V818)

			tmp8629 := PrimEqual(symnl, tmp8628)

			var ifres8624 Obj

			if True == tmp8629 {
				tmp8626 := PrimTail(V818)

				tmp8627 := PrimEqual(Nil, tmp8626)

				var ifres8625 Obj

				if True == tmp8627 {
					ifres8625 = True

				} else {
					ifres8625 = False

				}

				ifres8624 = ifres8625

			} else {
				ifres8624 = False

			}

			var ifres8623 Obj

			if True == ifres8624 {
				ifres8623 = True

			} else {
				ifres8623 = False

			}

			ifres8622 = ifres8623

		} else {
			ifres8622 = False

		}

		if True == ifres8622 {
			tmp8621 := PrimCons(MakeNumber(1), Nil)

			__e.Return(PrimCons(symnl, tmp8621))
			return

		} else {
			__e.Return(V818)
			return
		}

	}, 1)

	tmp8631 := Call(__e, PrimNS1Value(symns2_1set), symshen_4nl_1macro, tmp8619)

	_ = tmp8631

	tmp8632 := MakeNative(func(__e *ControlFlow) {
		V820 := __e.Get(1)
		_ = V820
		tmp8671 := PrimIsPair(V820)

		var ifres8644 Obj

		if True == tmp8671 {
			tmp8669 := PrimTail(V820)

			tmp8670 := PrimIsPair(tmp8669)

			var ifres8646 Obj

			if True == tmp8670 {
				tmp8666 := PrimTail(V820)

				tmp8667 := PrimTail(tmp8666)

				tmp8668 := PrimIsPair(tmp8667)

				var ifres8648 Obj

				if True == tmp8668 {
					tmp8662 := PrimTail(V820)

					tmp8663 := PrimTail(tmp8662)

					tmp8664 := PrimTail(tmp8663)

					tmp8665 := PrimIsPair(tmp8664)

					var ifres8650 Obj

					if True == tmp8665 {
						tmp8652 := PrimHead(V820)

						tmp8653 := PrimCons(symdo, Nil)

						tmp8654 := PrimCons(sym_d, tmp8653)

						tmp8655 := PrimCons(sym_7, tmp8654)

						tmp8656 := PrimCons(symor, tmp8655)

						tmp8657 := PrimCons(symand, tmp8656)

						tmp8658 := PrimCons(symappend, tmp8657)

						tmp8659 := PrimCons(sym_8v, tmp8658)

						tmp8660 := PrimCons(sym_8p, tmp8659)

						tmp8661 := Call(__e, PrimNS2Value(symelement_2), tmp8652, tmp8660)

						var ifres8651 Obj

						if True == tmp8661 {
							ifres8651 = True

						} else {
							ifres8651 = False

						}

						ifres8650 = ifres8651

					} else {
						ifres8650 = False

					}

					var ifres8649 Obj

					if True == ifres8650 {
						ifres8649 = True

					} else {
						ifres8649 = False

					}

					ifres8648 = ifres8649

				} else {
					ifres8648 = False

				}

				var ifres8647 Obj

				if True == ifres8648 {
					ifres8647 = True

				} else {
					ifres8647 = False

				}

				ifres8646 = ifres8647

			} else {
				ifres8646 = False

			}

			var ifres8645 Obj

			if True == ifres8646 {
				ifres8645 = True

			} else {
				ifres8645 = False

			}

			ifres8644 = ifres8645

		} else {
			ifres8644 = False

		}

		if True == ifres8644 {
			tmp8634 := PrimHead(V820)

			tmp8635 := PrimTail(V820)

			tmp8636 := PrimHead(tmp8635)

			tmp8637 := PrimHead(V820)

			tmp8638 := PrimTail(V820)

			tmp8639 := PrimTail(tmp8638)

			tmp8640 := PrimCons(tmp8637, tmp8639)

			tmp8641 := Call(__e, PrimNS2Value(symshen_4assoc_1macro), tmp8640)

			tmp8642 := PrimCons(tmp8641, Nil)

			tmp8643 := PrimCons(tmp8636, tmp8642)

			__e.Return(PrimCons(tmp8634, tmp8643))
			return

		} else {
			__e.Return(V820)
			return
		}

	}, 1)

	tmp8672 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1macro, tmp8632)

	_ = tmp8672

	tmp8673 := MakeNative(func(__e *ControlFlow) {
		V822 := __e.Get(1)
		_ = V822
		tmp8714 := PrimIsPair(V822)

		var ifres8688 Obj

		if True == tmp8714 {
			tmp8712 := PrimHead(V822)

			tmp8713 := PrimEqual(symlet, tmp8712)

			var ifres8690 Obj

			if True == tmp8713 {
				tmp8710 := PrimTail(V822)

				tmp8711 := PrimIsPair(tmp8710)

				var ifres8692 Obj

				if True == tmp8711 {
					tmp8707 := PrimTail(V822)

					tmp8708 := PrimTail(tmp8707)

					tmp8709 := PrimIsPair(tmp8708)

					var ifres8694 Obj

					if True == tmp8709 {
						tmp8703 := PrimTail(V822)

						tmp8704 := PrimTail(tmp8703)

						tmp8705 := PrimTail(tmp8704)

						tmp8706 := PrimIsPair(tmp8705)

						var ifres8696 Obj

						if True == tmp8706 {
							tmp8698 := PrimTail(V822)

							tmp8699 := PrimTail(tmp8698)

							tmp8700 := PrimTail(tmp8699)

							tmp8701 := PrimTail(tmp8700)

							tmp8702 := PrimIsPair(tmp8701)

							var ifres8697 Obj

							if True == tmp8702 {
								ifres8697 = True

							} else {
								ifres8697 = False

							}

							ifres8696 = ifres8697

						} else {
							ifres8696 = False

						}

						var ifres8695 Obj

						if True == ifres8696 {
							ifres8695 = True

						} else {
							ifres8695 = False

						}

						ifres8694 = ifres8695

					} else {
						ifres8694 = False

					}

					var ifres8693 Obj

					if True == ifres8694 {
						ifres8693 = True

					} else {
						ifres8693 = False

					}

					ifres8692 = ifres8693

				} else {
					ifres8692 = False

				}

				var ifres8691 Obj

				if True == ifres8692 {
					ifres8691 = True

				} else {
					ifres8691 = False

				}

				ifres8690 = ifres8691

			} else {
				ifres8690 = False

			}

			var ifres8689 Obj

			if True == ifres8690 {
				ifres8689 = True

			} else {
				ifres8689 = False

			}

			ifres8688 = ifres8689

		} else {
			ifres8688 = False

		}

		if True == ifres8688 {
			tmp8675 := PrimTail(V822)

			tmp8676 := PrimHead(tmp8675)

			tmp8677 := PrimTail(V822)

			tmp8678 := PrimTail(tmp8677)

			tmp8679 := PrimHead(tmp8678)

			tmp8680 := PrimTail(V822)

			tmp8681 := PrimTail(tmp8680)

			tmp8682 := PrimTail(tmp8681)

			tmp8683 := PrimCons(symlet, tmp8682)

			tmp8684 := Call(__e, PrimNS2Value(symshen_4let_1macro), tmp8683)

			tmp8685 := PrimCons(tmp8684, Nil)

			tmp8686 := PrimCons(tmp8679, tmp8685)

			tmp8687 := PrimCons(tmp8676, tmp8686)

			__e.Return(PrimCons(symlet, tmp8687))
			return

		} else {
			__e.Return(V822)
			return
		}

	}, 1)

	tmp8715 := Call(__e, PrimNS1Value(symns2_1set), symshen_4let_1macro, tmp8673)

	_ = tmp8715

	tmp8716 := MakeNative(func(__e *ControlFlow) {
		V824 := __e.Get(1)
		_ = V824
		tmp8767 := PrimIsPair(V824)

		var ifres8748 Obj

		if True == tmp8767 {
			tmp8765 := PrimHead(V824)

			tmp8766 := PrimEqual(sym_c_4, tmp8765)

			var ifres8750 Obj

			if True == tmp8766 {
				tmp8763 := PrimTail(V824)

				tmp8764 := PrimIsPair(tmp8763)

				var ifres8752 Obj

				if True == tmp8764 {
					tmp8760 := PrimTail(V824)

					tmp8761 := PrimTail(tmp8760)

					tmp8762 := PrimIsPair(tmp8761)

					var ifres8754 Obj

					if True == tmp8762 {
						tmp8756 := PrimTail(V824)

						tmp8757 := PrimTail(tmp8756)

						tmp8758 := PrimTail(tmp8757)

						tmp8759 := PrimIsPair(tmp8758)

						var ifres8755 Obj

						if True == tmp8759 {
							ifres8755 = True

						} else {
							ifres8755 = False

						}

						ifres8754 = ifres8755

					} else {
						ifres8754 = False

					}

					var ifres8753 Obj

					if True == ifres8754 {
						ifres8753 = True

					} else {
						ifres8753 = False

					}

					ifres8752 = ifres8753

				} else {
					ifres8752 = False

				}

				var ifres8751 Obj

				if True == ifres8752 {
					ifres8751 = True

				} else {
					ifres8751 = False

				}

				ifres8750 = ifres8751

			} else {
				ifres8750 = False

			}

			var ifres8749 Obj

			if True == ifres8750 {
				ifres8749 = True

			} else {
				ifres8749 = False

			}

			ifres8748 = ifres8749

		} else {
			ifres8748 = False

		}

		if True == ifres8748 {
			tmp8740 := PrimTail(V824)

			tmp8741 := PrimHead(tmp8740)

			tmp8742 := PrimTail(V824)

			tmp8743 := PrimTail(tmp8742)

			tmp8744 := PrimCons(sym_c_4, tmp8743)

			tmp8745 := Call(__e, PrimNS2Value(symshen_4abs_1macro), tmp8744)

			tmp8746 := PrimCons(tmp8745, Nil)

			tmp8747 := PrimCons(tmp8741, tmp8746)

			__e.Return(PrimCons(symlambda, tmp8747))
			return

		} else {
			tmp8739 := PrimIsPair(V824)

			var ifres8720 Obj

			if True == tmp8739 {
				tmp8737 := PrimHead(V824)

				tmp8738 := PrimEqual(sym_c_4, tmp8737)

				var ifres8722 Obj

				if True == tmp8738 {
					tmp8735 := PrimTail(V824)

					tmp8736 := PrimIsPair(tmp8735)

					var ifres8724 Obj

					if True == tmp8736 {
						tmp8732 := PrimTail(V824)

						tmp8733 := PrimTail(tmp8732)

						tmp8734 := PrimIsPair(tmp8733)

						var ifres8726 Obj

						if True == tmp8734 {
							tmp8728 := PrimTail(V824)

							tmp8729 := PrimTail(tmp8728)

							tmp8730 := PrimTail(tmp8729)

							tmp8731 := PrimEqual(Nil, tmp8730)

							var ifres8727 Obj

							if True == tmp8731 {
								ifres8727 = True

							} else {
								ifres8727 = False

							}

							ifres8726 = ifres8727

						} else {
							ifres8726 = False

						}

						var ifres8725 Obj

						if True == ifres8726 {
							ifres8725 = True

						} else {
							ifres8725 = False

						}

						ifres8724 = ifres8725

					} else {
						ifres8724 = False

					}

					var ifres8723 Obj

					if True == ifres8724 {
						ifres8723 = True

					} else {
						ifres8723 = False

					}

					ifres8722 = ifres8723

				} else {
					ifres8722 = False

				}

				var ifres8721 Obj

				if True == ifres8722 {
					ifres8721 = True

				} else {
					ifres8721 = False

				}

				ifres8720 = ifres8721

			} else {
				ifres8720 = False

			}

			if True == ifres8720 {
				tmp8719 := PrimTail(V824)

				__e.Return(PrimCons(symlambda, tmp8719))
				return

			} else {
				__e.Return(V824)
				return
			}

		}

	}, 1)

	tmp8768 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abs_1macro, tmp8716)

	_ = tmp8768

	tmp8769 := MakeNative(func(__e *ControlFlow) {
		V828 := __e.Get(1)
		_ = V828
		tmp8865 := PrimIsPair(V828)

		var ifres8847 Obj

		if True == tmp8865 {
			tmp8863 := PrimHead(V828)

			tmp8864 := PrimEqual(symcases, tmp8863)

			var ifres8849 Obj

			if True == tmp8864 {
				tmp8861 := PrimTail(V828)

				tmp8862 := PrimIsPair(tmp8861)

				var ifres8851 Obj

				if True == tmp8862 {
					tmp8858 := PrimTail(V828)

					tmp8859 := PrimHead(tmp8858)

					tmp8860 := PrimEqual(True, tmp8859)

					var ifres8853 Obj

					if True == tmp8860 {
						tmp8855 := PrimTail(V828)

						tmp8856 := PrimTail(tmp8855)

						tmp8857 := PrimIsPair(tmp8856)

						var ifres8854 Obj

						if True == tmp8857 {
							ifres8854 = True

						} else {
							ifres8854 = False

						}

						ifres8853 = ifres8854

					} else {
						ifres8853 = False

					}

					var ifres8852 Obj

					if True == ifres8853 {
						ifres8852 = True

					} else {
						ifres8852 = False

					}

					ifres8851 = ifres8852

				} else {
					ifres8851 = False

				}

				var ifres8850 Obj

				if True == ifres8851 {
					ifres8850 = True

				} else {
					ifres8850 = False

				}

				ifres8849 = ifres8850

			} else {
				ifres8849 = False

			}

			var ifres8848 Obj

			if True == ifres8849 {
				ifres8848 = True

			} else {
				ifres8848 = False

			}

			ifres8847 = ifres8848

		} else {
			ifres8847 = False

		}

		if True == ifres8847 {
			tmp8845 := PrimTail(V828)

			tmp8846 := PrimTail(tmp8845)

			__e.Return(PrimHead(tmp8846))
			return

		} else {
			tmp8844 := PrimIsPair(V828)

			var ifres8825 Obj

			if True == tmp8844 {
				tmp8842 := PrimHead(V828)

				tmp8843 := PrimEqual(symcases, tmp8842)

				var ifres8827 Obj

				if True == tmp8843 {
					tmp8840 := PrimTail(V828)

					tmp8841 := PrimIsPair(tmp8840)

					var ifres8829 Obj

					if True == tmp8841 {
						tmp8837 := PrimTail(V828)

						tmp8838 := PrimTail(tmp8837)

						tmp8839 := PrimIsPair(tmp8838)

						var ifres8831 Obj

						if True == tmp8839 {
							tmp8833 := PrimTail(V828)

							tmp8834 := PrimTail(tmp8833)

							tmp8835 := PrimTail(tmp8834)

							tmp8836 := PrimEqual(Nil, tmp8835)

							var ifres8832 Obj

							if True == tmp8836 {
								ifres8832 = True

							} else {
								ifres8832 = False

							}

							ifres8831 = ifres8832

						} else {
							ifres8831 = False

						}

						var ifres8830 Obj

						if True == ifres8831 {
							ifres8830 = True

						} else {
							ifres8830 = False

						}

						ifres8829 = ifres8830

					} else {
						ifres8829 = False

					}

					var ifres8828 Obj

					if True == ifres8829 {
						ifres8828 = True

					} else {
						ifres8828 = False

					}

					ifres8827 = ifres8828

				} else {
					ifres8827 = False

				}

				var ifres8826 Obj

				if True == ifres8827 {
					ifres8826 = True

				} else {
					ifres8826 = False

				}

				ifres8825 = ifres8826

			} else {
				ifres8825 = False

			}

			if True == ifres8825 {
				tmp8815 := PrimTail(V828)

				tmp8816 := PrimHead(tmp8815)

				tmp8817 := PrimTail(V828)

				tmp8818 := PrimTail(tmp8817)

				tmp8819 := PrimHead(tmp8818)

				tmp8820 := PrimCons(MakeString("error: cases exhausted"), Nil)

				tmp8821 := PrimCons(symsimple_1error, tmp8820)

				tmp8822 := PrimCons(tmp8821, Nil)

				tmp8823 := PrimCons(tmp8819, tmp8822)

				tmp8824 := PrimCons(tmp8816, tmp8823)

				__e.Return(PrimCons(symif, tmp8824))
				return

			} else {
				tmp8814 := PrimIsPair(V828)

				var ifres8801 Obj

				if True == tmp8814 {
					tmp8812 := PrimHead(V828)

					tmp8813 := PrimEqual(symcases, tmp8812)

					var ifres8803 Obj

					if True == tmp8813 {
						tmp8810 := PrimTail(V828)

						tmp8811 := PrimIsPair(tmp8810)

						var ifres8805 Obj

						if True == tmp8811 {
							tmp8807 := PrimTail(V828)

							tmp8808 := PrimTail(tmp8807)

							tmp8809 := PrimIsPair(tmp8808)

							var ifres8806 Obj

							if True == tmp8809 {
								ifres8806 = True

							} else {
								ifres8806 = False

							}

							ifres8805 = ifres8806

						} else {
							ifres8805 = False

						}

						var ifres8804 Obj

						if True == ifres8805 {
							ifres8804 = True

						} else {
							ifres8804 = False

						}

						ifres8803 = ifres8804

					} else {
						ifres8803 = False

					}

					var ifres8802 Obj

					if True == ifres8803 {
						ifres8802 = True

					} else {
						ifres8802 = False

					}

					ifres8801 = ifres8802

				} else {
					ifres8801 = False

				}

				if True == ifres8801 {
					tmp8788 := PrimTail(V828)

					tmp8789 := PrimHead(tmp8788)

					tmp8790 := PrimTail(V828)

					tmp8791 := PrimTail(tmp8790)

					tmp8792 := PrimHead(tmp8791)

					tmp8793 := PrimTail(V828)

					tmp8794 := PrimTail(tmp8793)

					tmp8795 := PrimTail(tmp8794)

					tmp8796 := PrimCons(symcases, tmp8795)

					tmp8797 := Call(__e, PrimNS2Value(symshen_4cases_1macro), tmp8796)

					tmp8798 := PrimCons(tmp8797, Nil)

					tmp8799 := PrimCons(tmp8792, tmp8798)

					tmp8800 := PrimCons(tmp8789, tmp8799)

					__e.Return(PrimCons(symif, tmp8800))
					return

				} else {
					tmp8787 := PrimIsPair(V828)

					var ifres8774 Obj

					if True == tmp8787 {
						tmp8785 := PrimHead(V828)

						tmp8786 := PrimEqual(symcases, tmp8785)

						var ifres8776 Obj

						if True == tmp8786 {
							tmp8783 := PrimTail(V828)

							tmp8784 := PrimIsPair(tmp8783)

							var ifres8778 Obj

							if True == tmp8784 {
								tmp8780 := PrimTail(V828)

								tmp8781 := PrimTail(tmp8780)

								tmp8782 := PrimEqual(Nil, tmp8781)

								var ifres8779 Obj

								if True == tmp8782 {
									ifres8779 = True

								} else {
									ifres8779 = False

								}

								ifres8778 = ifres8779

							} else {
								ifres8778 = False

							}

							var ifres8777 Obj

							if True == ifres8778 {
								ifres8777 = True

							} else {
								ifres8777 = False

							}

							ifres8776 = ifres8777

						} else {
							ifres8776 = False

						}

						var ifres8775 Obj

						if True == ifres8776 {
							ifres8775 = True

						} else {
							ifres8775 = False

						}

						ifres8774 = ifres8775

					} else {
						ifres8774 = False

					}

					if True == ifres8774 {
						__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
						return
					} else {
						__e.Return(V828)
						return
					}

				}

			}

		}

	}, 1)

	tmp8866 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cases_1macro, tmp8769)

	_ = tmp8866

	tmp8867 := MakeNative(func(__e *ControlFlow) {
		V830 := __e.Get(1)
		_ = V830
		tmp8915 := PrimIsPair(V830)

		var ifres8902 Obj

		if True == tmp8915 {
			tmp8913 := PrimHead(V830)

			tmp8914 := PrimEqual(symtime, tmp8913)

			var ifres8904 Obj

			if True == tmp8914 {
				tmp8911 := PrimTail(V830)

				tmp8912 := PrimIsPair(tmp8911)

				var ifres8906 Obj

				if True == tmp8912 {
					tmp8908 := PrimTail(V830)

					tmp8909 := PrimTail(tmp8908)

					tmp8910 := PrimEqual(Nil, tmp8909)

					var ifres8907 Obj

					if True == tmp8910 {
						ifres8907 = True

					} else {
						ifres8907 = False

					}

					ifres8906 = ifres8907

				} else {
					ifres8906 = False

				}

				var ifres8905 Obj

				if True == ifres8906 {
					ifres8905 = True

				} else {
					ifres8905 = False

				}

				ifres8904 = ifres8905

			} else {
				ifres8904 = False

			}

			var ifres8903 Obj

			if True == ifres8904 {
				ifres8903 = True

			} else {
				ifres8903 = False

			}

			ifres8902 = ifres8903

		} else {
			ifres8902 = False

		}

		if True == ifres8902 {
			tmp8869 := PrimCons(symrun, Nil)

			tmp8870 := PrimCons(symget_1time, tmp8869)

			tmp8871 := PrimTail(V830)

			tmp8872 := PrimHead(tmp8871)

			tmp8873 := PrimCons(symrun, Nil)

			tmp8874 := PrimCons(symget_1time, tmp8873)

			tmp8875 := PrimCons(symStart, Nil)

			tmp8876 := PrimCons(symFinish, tmp8875)

			tmp8877 := PrimCons(sym_1, tmp8876)

			tmp8878 := PrimCons(symTime, Nil)

			tmp8879 := PrimCons(symstr, tmp8878)

			tmp8880 := PrimCons(MakeString(" secs\n"), Nil)

			tmp8881 := PrimCons(tmp8879, tmp8880)

			tmp8882 := PrimCons(symcn, tmp8881)

			tmp8883 := PrimCons(tmp8882, Nil)

			tmp8884 := PrimCons(MakeString("\nrun time: "), tmp8883)

			tmp8885 := PrimCons(symcn, tmp8884)

			tmp8886 := PrimCons(symstoutput, Nil)

			tmp8887 := PrimCons(tmp8886, Nil)

			tmp8888 := PrimCons(tmp8885, tmp8887)

			tmp8889 := PrimCons(symshen_4prhush, tmp8888)

			tmp8890 := PrimCons(symResult, Nil)

			tmp8891 := PrimCons(tmp8889, tmp8890)

			tmp8892 := PrimCons(symMessage, tmp8891)

			tmp8893 := PrimCons(tmp8877, tmp8892)

			tmp8894 := PrimCons(symTime, tmp8893)

			tmp8895 := PrimCons(tmp8874, tmp8894)

			tmp8896 := PrimCons(symFinish, tmp8895)

			tmp8897 := PrimCons(tmp8872, tmp8896)

			tmp8898 := PrimCons(symResult, tmp8897)

			tmp8899 := PrimCons(tmp8870, tmp8898)

			tmp8900 := PrimCons(symStart, tmp8899)

			tmp8901 := PrimCons(symlet, tmp8900)

			__e.TailApply(PrimNS2Value(symshen_4let_1macro), tmp8901)
			return

		} else {
			__e.Return(V830)
			return
		}

	}, 1)

	tmp8916 := Call(__e, PrimNS1Value(symns2_1set), symshen_4timer_1macro, tmp8867)

	_ = tmp8916

	tmp8917 := MakeNative(func(__e *ControlFlow) {
		V832 := __e.Get(1)
		_ = V832
		tmp8924 := PrimIsPair(V832)

		if True == tmp8924 {
			tmp8919 := PrimHead(V832)

			tmp8920 := PrimTail(V832)

			tmp8921 := Call(__e, PrimNS2Value(symshen_4tuple_1up), tmp8920)

			tmp8922 := PrimCons(tmp8921, Nil)

			tmp8923 := PrimCons(tmp8919, tmp8922)

			__e.Return(PrimCons(sym_8p, tmp8923))
			return

		} else {
			__e.Return(V832)
			return
		}

	}, 1)

	tmp8925 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tuple_1up, tmp8917)

	_ = tmp8925

	tmp8926 := MakeNative(func(__e *ControlFlow) {
		V834 := __e.Get(1)
		_ = V834
		tmp9031 := PrimIsPair(V834)

		var ifres9005 Obj

		if True == tmp9031 {
			tmp9029 := PrimHead(V834)

			tmp9030 := PrimEqual(symput, tmp9029)

			var ifres9007 Obj

			if True == tmp9030 {
				tmp9027 := PrimTail(V834)

				tmp9028 := PrimIsPair(tmp9027)

				var ifres9009 Obj

				if True == tmp9028 {
					tmp9024 := PrimTail(V834)

					tmp9025 := PrimTail(tmp9024)

					tmp9026 := PrimIsPair(tmp9025)

					var ifres9011 Obj

					if True == tmp9026 {
						tmp9020 := PrimTail(V834)

						tmp9021 := PrimTail(tmp9020)

						tmp9022 := PrimTail(tmp9021)

						tmp9023 := PrimIsPair(tmp9022)

						var ifres9013 Obj

						if True == tmp9023 {
							tmp9015 := PrimTail(V834)

							tmp9016 := PrimTail(tmp9015)

							tmp9017 := PrimTail(tmp9016)

							tmp9018 := PrimTail(tmp9017)

							tmp9019 := PrimEqual(Nil, tmp9018)

							var ifres9014 Obj

							if True == tmp9019 {
								ifres9014 = True

							} else {
								ifres9014 = False

							}

							ifres9013 = ifres9014

						} else {
							ifres9013 = False

						}

						var ifres9012 Obj

						if True == ifres9013 {
							ifres9012 = True

						} else {
							ifres9012 = False

						}

						ifres9011 = ifres9012

					} else {
						ifres9011 = False

					}

					var ifres9010 Obj

					if True == ifres9011 {
						ifres9010 = True

					} else {
						ifres9010 = False

					}

					ifres9009 = ifres9010

				} else {
					ifres9009 = False

				}

				var ifres9008 Obj

				if True == ifres9009 {
					ifres9008 = True

				} else {
					ifres9008 = False

				}

				ifres9007 = ifres9008

			} else {
				ifres9007 = False

			}

			var ifres9006 Obj

			if True == ifres9007 {
				ifres9006 = True

			} else {
				ifres9006 = False

			}

			ifres9005 = ifres9006

		} else {
			ifres9005 = False

		}

		if True == ifres9005 {
			tmp8990 := PrimTail(V834)

			tmp8991 := PrimHead(tmp8990)

			tmp8992 := PrimTail(V834)

			tmp8993 := PrimTail(tmp8992)

			tmp8994 := PrimHead(tmp8993)

			tmp8995 := PrimTail(V834)

			tmp8996 := PrimTail(tmp8995)

			tmp8997 := PrimTail(tmp8996)

			tmp8998 := PrimHead(tmp8997)

			tmp8999 := PrimCons(sym_dproperty_1vector_d, Nil)

			tmp9000 := PrimCons(symvalue, tmp8999)

			tmp9001 := PrimCons(tmp9000, Nil)

			tmp9002 := PrimCons(tmp8998, tmp9001)

			tmp9003 := PrimCons(tmp8994, tmp9002)

			tmp9004 := PrimCons(tmp8991, tmp9003)

			__e.Return(PrimCons(symput, tmp9004))
			return

		} else {
			tmp8989 := PrimIsPair(V834)

			var ifres8970 Obj

			if True == tmp8989 {
				tmp8987 := PrimHead(V834)

				tmp8988 := PrimEqual(symget, tmp8987)

				var ifres8972 Obj

				if True == tmp8988 {
					tmp8985 := PrimTail(V834)

					tmp8986 := PrimIsPair(tmp8985)

					var ifres8974 Obj

					if True == tmp8986 {
						tmp8982 := PrimTail(V834)

						tmp8983 := PrimTail(tmp8982)

						tmp8984 := PrimIsPair(tmp8983)

						var ifres8976 Obj

						if True == tmp8984 {
							tmp8978 := PrimTail(V834)

							tmp8979 := PrimTail(tmp8978)

							tmp8980 := PrimTail(tmp8979)

							tmp8981 := PrimEqual(Nil, tmp8980)

							var ifres8977 Obj

							if True == tmp8981 {
								ifres8977 = True

							} else {
								ifres8977 = False

							}

							ifres8976 = ifres8977

						} else {
							ifres8976 = False

						}

						var ifres8975 Obj

						if True == ifres8976 {
							ifres8975 = True

						} else {
							ifres8975 = False

						}

						ifres8974 = ifres8975

					} else {
						ifres8974 = False

					}

					var ifres8973 Obj

					if True == ifres8974 {
						ifres8973 = True

					} else {
						ifres8973 = False

					}

					ifres8972 = ifres8973

				} else {
					ifres8972 = False

				}

				var ifres8971 Obj

				if True == ifres8972 {
					ifres8971 = True

				} else {
					ifres8971 = False

				}

				ifres8970 = ifres8971

			} else {
				ifres8970 = False

			}

			if True == ifres8970 {
				tmp8960 := PrimTail(V834)

				tmp8961 := PrimHead(tmp8960)

				tmp8962 := PrimTail(V834)

				tmp8963 := PrimTail(tmp8962)

				tmp8964 := PrimHead(tmp8963)

				tmp8965 := PrimCons(sym_dproperty_1vector_d, Nil)

				tmp8966 := PrimCons(symvalue, tmp8965)

				tmp8967 := PrimCons(tmp8966, Nil)

				tmp8968 := PrimCons(tmp8964, tmp8967)

				tmp8969 := PrimCons(tmp8961, tmp8968)

				__e.Return(PrimCons(symget, tmp8969))
				return

			} else {
				tmp8959 := PrimIsPair(V834)

				var ifres8940 Obj

				if True == tmp8959 {
					tmp8957 := PrimHead(V834)

					tmp8958 := PrimEqual(symunput, tmp8957)

					var ifres8942 Obj

					if True == tmp8958 {
						tmp8955 := PrimTail(V834)

						tmp8956 := PrimIsPair(tmp8955)

						var ifres8944 Obj

						if True == tmp8956 {
							tmp8952 := PrimTail(V834)

							tmp8953 := PrimTail(tmp8952)

							tmp8954 := PrimIsPair(tmp8953)

							var ifres8946 Obj

							if True == tmp8954 {
								tmp8948 := PrimTail(V834)

								tmp8949 := PrimTail(tmp8948)

								tmp8950 := PrimTail(tmp8949)

								tmp8951 := PrimEqual(Nil, tmp8950)

								var ifres8947 Obj

								if True == tmp8951 {
									ifres8947 = True

								} else {
									ifres8947 = False

								}

								ifres8946 = ifres8947

							} else {
								ifres8946 = False

							}

							var ifres8945 Obj

							if True == ifres8946 {
								ifres8945 = True

							} else {
								ifres8945 = False

							}

							ifres8944 = ifres8945

						} else {
							ifres8944 = False

						}

						var ifres8943 Obj

						if True == ifres8944 {
							ifres8943 = True

						} else {
							ifres8943 = False

						}

						ifres8942 = ifres8943

					} else {
						ifres8942 = False

					}

					var ifres8941 Obj

					if True == ifres8942 {
						ifres8941 = True

					} else {
						ifres8941 = False

					}

					ifres8940 = ifres8941

				} else {
					ifres8940 = False

				}

				if True == ifres8940 {
					tmp8930 := PrimTail(V834)

					tmp8931 := PrimHead(tmp8930)

					tmp8932 := PrimTail(V834)

					tmp8933 := PrimTail(tmp8932)

					tmp8934 := PrimHead(tmp8933)

					tmp8935 := PrimCons(sym_dproperty_1vector_d, Nil)

					tmp8936 := PrimCons(symvalue, tmp8935)

					tmp8937 := PrimCons(tmp8936, Nil)

					tmp8938 := PrimCons(tmp8934, tmp8937)

					tmp8939 := PrimCons(tmp8931, tmp8938)

					__e.Return(PrimCons(symunput, tmp8939))
					return

				} else {
					__e.Return(V834)
					return
				}

			}

		}

	}, 1)

	tmp9032 := Call(__e, PrimNS1Value(symns2_1set), symshen_4put_cget_1macro, tmp8926)

	_ = tmp9032

	tmp9033 := MakeNative(func(__e *ControlFlow) {
		V836 := __e.Get(1)
		_ = V836
		tmp9053 := PrimIsPair(V836)

		var ifres9040 Obj

		if True == tmp9053 {
			tmp9051 := PrimHead(V836)

			tmp9052 := PrimEqual(symfunction, tmp9051)

			var ifres9042 Obj

			if True == tmp9052 {
				tmp9049 := PrimTail(V836)

				tmp9050 := PrimIsPair(tmp9049)

				var ifres9044 Obj

				if True == tmp9050 {
					tmp9046 := PrimTail(V836)

					tmp9047 := PrimTail(tmp9046)

					tmp9048 := PrimEqual(Nil, tmp9047)

					var ifres9045 Obj

					if True == tmp9048 {
						ifres9045 = True

					} else {
						ifres9045 = False

					}

					ifres9044 = ifres9045

				} else {
					ifres9044 = False

				}

				var ifres9043 Obj

				if True == ifres9044 {
					ifres9043 = True

				} else {
					ifres9043 = False

				}

				ifres9042 = ifres9043

			} else {
				ifres9042 = False

			}

			var ifres9041 Obj

			if True == ifres9042 {
				ifres9041 = True

			} else {
				ifres9041 = False

			}

			ifres9040 = ifres9041

		} else {
			ifres9040 = False

		}

		if True == ifres9040 {
			tmp9035 := PrimTail(V836)

			tmp9036 := PrimHead(tmp9035)

			tmp9037 := PrimTail(V836)

			tmp9038 := PrimHead(tmp9037)

			tmp9039 := Call(__e, PrimNS2Value(symarity), tmp9038)

			__e.TailApply(PrimNS2Value(symshen_4function_1abstraction), tmp9036, tmp9039)
			return

		} else {
			__e.Return(V836)
			return
		}

	}, 1)

	tmp9054 := Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1macro, tmp9033)

	_ = tmp9054

	tmp9055 := MakeNative(func(__e *ControlFlow) {
		V839 := __e.Get(1)
		_ = V839
		V840 := __e.Get(2)
		_ = V840
		tmp9061 := PrimEqual(MakeNumber(0), V840)

		if True == tmp9061 {
			tmp9060 := Call(__e, PrimNS2Value(symshen_4app), V839, MakeString(" has no lambda form\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp9060))
			return

		} else {
			tmp9059 := PrimEqual(MakeNumber(-1), V840)

			if True == tmp9059 {
				tmp9058 := PrimCons(V839, Nil)

				__e.Return(PrimCons(symfunction, tmp9058))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4function_1abstraction_1help), V839, V840, Nil)
				return
			}

		}

	}, 2)

	tmp9062 := Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1abstraction, tmp9055)

	_ = tmp9062

	tmp9063 := MakeNative(func(__e *ControlFlow) {
		V844 := __e.Get(1)
		_ = V844
		V845 := __e.Get(2)
		_ = V845
		V846 := __e.Get(3)
		_ = V846
		tmp9073 := PrimEqual(MakeNumber(0), V845)

		if True == tmp9073 {
			__e.Return(PrimCons(V844, V846))
			return
		} else {
			tmp9065 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp9066 := PrimNumberSubtract(V845, MakeNumber(1))

				tmp9067 := PrimCons(X, Nil)

				tmp9068 := Call(__e, PrimNS2Value(symappend), V846, tmp9067)

				tmp9069 := Call(__e, PrimNS2Value(symshen_4function_1abstraction_1help), V844, tmp9066, tmp9068)

				tmp9070 := PrimCons(tmp9069, Nil)

				tmp9071 := PrimCons(X, tmp9070)

				__e.Return(PrimCons(sym_c_4, tmp9071))
				return

			}, 1)

			tmp9072 := Call(__e, PrimNS2Value(symgensym), symV)

			__e.TailApply(tmp9065, tmp9072)
			return

		}

	}, 3)

	tmp9074 := Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1abstraction_1help, tmp9063)

	_ = tmp9074

	tmp9075 := MakeNative(func(__e *ControlFlow) {
		V848 := __e.Get(1)
		_ = V848
		tmp9076 := MakeNative(func(__e *ControlFlow) {
			MacroReg := __e.Get(1)
			_ = MacroReg
			tmp9077 := MakeNative(func(__e *ControlFlow) {
				Pos := __e.Get(1)
				_ = Pos
				tmp9078 := MakeNative(func(__e *ControlFlow) {
					Remove1 := __e.Get(1)
					_ = Remove1
					tmp9079 := MakeNative(func(__e *ControlFlow) {
						Remove2 := __e.Get(1)
						_ = Remove2
						__e.Return(V848)
						return
					}, 1)

					tmp9080 := PrimNS3Value(sym_dmacros_d)

					tmp9081 := Call(__e, PrimNS2Value(symshen_4remove_1nth), Pos, tmp9080)

					tmp9082 := PrimNS3Set(sym_dmacros_d, tmp9081)

					__e.TailApply(tmp9079, tmp9082)
					return

				}, 1)

				tmp9083 := Call(__e, PrimNS2Value(symremove), V848, MacroReg)

				tmp9084 := PrimNS3Set(symshen_4_dmacroreg_d, tmp9083)

				__e.TailApply(tmp9078, tmp9084)
				return

			}, 1)

			tmp9085 := Call(__e, PrimNS2Value(symshen_4findpos), V848, MacroReg)

			__e.TailApply(tmp9077, tmp9085)
			return

		}, 1)

		tmp9086 := PrimNS3Value(symshen_4_dmacroreg_d)

		__e.TailApply(tmp9076, tmp9086)
		return

	}, 1)

	tmp9087 := Call(__e, PrimNS1Value(symns2_1set), symundefmacro, tmp9075)

	_ = tmp9087

	tmp9088 := MakeNative(func(__e *ControlFlow) {
		V858 := __e.Get(1)
		_ = V858
		V859 := __e.Get(2)
		_ = V859
		tmp9101 := PrimEqual(Nil, V859)

		if True == tmp9101 {
			tmp9100 := Call(__e, PrimNS2Value(symshen_4app), V858, MakeString(" is not a macro\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp9100))
			return

		} else {
			tmp9099 := PrimIsPair(V859)

			var ifres9095 Obj

			if True == tmp9099 {
				tmp9097 := PrimHead(V859)

				tmp9098 := PrimEqual(tmp9097, V858)

				var ifres9096 Obj

				if True == tmp9098 {
					ifres9096 = True

				} else {
					ifres9096 = False

				}

				ifres9095 = ifres9096

			} else {
				ifres9095 = False

			}

			if True == ifres9095 {
				__e.Return(MakeNumber(1))
				return
			} else {
				tmp9094 := PrimIsPair(V859)

				if True == tmp9094 {
					tmp9092 := PrimTail(V859)

					tmp9093 := Call(__e, PrimNS2Value(symshen_4findpos), V858, tmp9092)

					__e.Return(PrimNumberAdd(MakeNumber(1), tmp9093))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4findpos)
					return
				}

			}

		}

	}, 2)

	tmp9102 := Call(__e, PrimNS1Value(symns2_1set), symshen_4findpos, tmp9088)

	_ = tmp9102

	tmp9103 := MakeNative(func(__e *ControlFlow) {
		V864 := __e.Get(1)
		_ = V864
		V865 := __e.Get(2)
		_ = V865
		tmp9114 := PrimEqual(MakeNumber(1), V864)

		var ifres9111 Obj

		if True == tmp9114 {
			tmp9113 := PrimIsPair(V865)

			var ifres9112 Obj

			if True == tmp9113 {
				ifres9112 = True

			} else {
				ifres9112 = False

			}

			ifres9111 = ifres9112

		} else {
			ifres9111 = False

		}

		if True == ifres9111 {
			__e.Return(PrimTail(V865))
			return
		} else {
			tmp9110 := PrimIsPair(V865)

			if True == tmp9110 {
				tmp9106 := PrimHead(V865)

				tmp9107 := PrimNumberSubtract(V864, MakeNumber(1))

				tmp9108 := PrimTail(V865)

				tmp9109 := Call(__e, PrimNS2Value(symshen_4remove_1nth), tmp9107, tmp9108)

				__e.Return(PrimCons(tmp9106, tmp9109))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4remove_1nth)
				return
			}

		}

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4remove_1nth, tmp9103)
	return

}, 0)
