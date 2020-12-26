package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen8230 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1920 := __args[0]
			_ = V1920
			__e.TailApply(V1920)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("thaw"), gen8230)

		gen8236 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1922 := __args[0]
			_ = V1922
			gen8231 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Y := __args[0]
				_ = Y
				__e.TailApply(ShenFunc(symmacroexpand), Y)

				return
			}, 1)
			gen8232 := Call(__e, ShenFunc(symshen_4walk), gen8231, V1922)

			Macroexpand := gen8232
			gen8235 := Call(__e, ShenFunc(symshen_4packaged_2), Macroexpand)

			if True == gen8235 {
				gen8233 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Z := __args[0]
					_ = Z
					__e.TailApply(ShenFunc(symshen_4eval_1without_1macros), Z)

					return
				}, 1)
				gen8234 := Call(__e, ShenFunc(symshen_4package_1contents), Macroexpand)

				__e.TailApply(ShenFunc(symmap), gen8233, gen8234)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4eval_1without_1macros), Macroexpand)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("eval"), gen8236)

		gen8239 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1924 := __args[0]
			_ = V1924
			gen8237 := Call(__e, ShenFunc(symshen_4proc_1input_7), V1924)

			gen8238 := Call(__e, ShenFunc(symshen_4elim_1def), gen8237)

			__e.TailApply(ShenFunc(symeval_1kl), gen8238)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.eval-without-macros"), gen8239)

		gen8286 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1926 := __args[0]
			_ = V1926
			gen8284 := Call(__e, ShenFunc(symcons_2), V1926)

			var gen8285 Obj
			if True == gen8284 {
				gen8281 := Call(__e, ShenFunc(symhd), V1926)

				gen8282 := Call(__e, ShenFunc(sym_a), MakeSymbol("input+"), gen8281)

				var gen8283 Obj
				if True == gen8282 {
					gen8278 := Call(__e, ShenFunc(symtl), V1926)

					gen8279 := Call(__e, ShenFunc(symcons_2), gen8278)

					var gen8280 Obj
					if True == gen8279 {
						gen8274 := Call(__e, ShenFunc(symtl), V1926)

						gen8275 := Call(__e, ShenFunc(symtl), gen8274)

						gen8276 := Call(__e, ShenFunc(symcons_2), gen8275)

						var gen8277 Obj
						if True == gen8276 {
							gen8270 := Call(__e, ShenFunc(symtl), V1926)

							gen8271 := Call(__e, ShenFunc(symtl), gen8270)

							gen8272 := Call(__e, ShenFunc(symtl), gen8271)

							gen8273 := Call(__e, ShenFunc(sym_a), Nil, gen8272)

							if True == gen8273 {
								gen8277 = True
							} else {
								gen8277 = False
							}

						} else {
							gen8277 = False
						}
						if True == gen8277 {
							gen8280 = True
						} else {
							gen8280 = False
						}

					} else {
						gen8280 = False
					}
					if True == gen8280 {
						gen8283 = True
					} else {
						gen8283 = False
					}

				} else {
					gen8283 = False
				}
				if True == gen8283 {
					gen8285 = True
				} else {
					gen8285 = False
				}

			} else {
				gen8285 = False
			}
			if True == gen8285 {
				gen8264 := Call(__e, ShenFunc(symtl), V1926)

				gen8265 := Call(__e, ShenFunc(symhd), gen8264)

				gen8266 := Call(__e, ShenFunc(symshen_4rcons__form), gen8265)

				gen8267 := Call(__e, ShenFunc(symtl), V1926)

				gen8268 := Call(__e, ShenFunc(symtl), gen8267)

				gen8269 := Call(__e, ShenFunc(symcons), gen8266, gen8268)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("input+"), gen8269)

				return

			} else {
				gen8262 := Call(__e, ShenFunc(symcons_2), V1926)

				var gen8263 Obj
				if True == gen8262 {
					gen8259 := Call(__e, ShenFunc(symhd), V1926)

					gen8260 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.read+"), gen8259)

					var gen8261 Obj
					if True == gen8260 {
						gen8256 := Call(__e, ShenFunc(symtl), V1926)

						gen8257 := Call(__e, ShenFunc(symcons_2), gen8256)

						var gen8258 Obj
						if True == gen8257 {
							gen8252 := Call(__e, ShenFunc(symtl), V1926)

							gen8253 := Call(__e, ShenFunc(symtl), gen8252)

							gen8254 := Call(__e, ShenFunc(symcons_2), gen8253)

							var gen8255 Obj
							if True == gen8254 {
								gen8248 := Call(__e, ShenFunc(symtl), V1926)

								gen8249 := Call(__e, ShenFunc(symtl), gen8248)

								gen8250 := Call(__e, ShenFunc(symtl), gen8249)

								gen8251 := Call(__e, ShenFunc(sym_a), Nil, gen8250)

								if True == gen8251 {
									gen8255 = True
								} else {
									gen8255 = False
								}

							} else {
								gen8255 = False
							}
							if True == gen8255 {
								gen8258 = True
							} else {
								gen8258 = False
							}

						} else {
							gen8258 = False
						}
						if True == gen8258 {
							gen8261 = True
						} else {
							gen8261 = False
						}

					} else {
						gen8261 = False
					}
					if True == gen8261 {
						gen8263 = True
					} else {
						gen8263 = False
					}

				} else {
					gen8263 = False
				}
				if True == gen8263 {
					gen8242 := Call(__e, ShenFunc(symtl), V1926)

					gen8243 := Call(__e, ShenFunc(symhd), gen8242)

					gen8244 := Call(__e, ShenFunc(symshen_4rcons__form), gen8243)

					gen8245 := Call(__e, ShenFunc(symtl), V1926)

					gen8246 := Call(__e, ShenFunc(symtl), gen8245)

					gen8247 := Call(__e, ShenFunc(symcons), gen8244, gen8246)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.read+"), gen8247)

					return

				} else {
					gen8241 := Call(__e, ShenFunc(symcons_2), V1926)

					if True == gen8241 {
						gen8240 := MakeNative(func(__e Evaluator, __args ...Obj) {
							Z := __args[0]
							_ = Z
							__e.TailApply(ShenFunc(symshen_4proc_1input_7), Z)

							return
						}, 1)
						__e.TailApply(ShenFunc(symmap), gen8240, V1926)

						return

					} else {
						__e.Return(V1926)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.proc-input+"), gen8286)

		gen8329 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1928 := __args[0]
			_ = V1928
			gen8327 := Call(__e, ShenFunc(symcons_2), V1928)

			var gen8328 Obj
			if True == gen8327 {
				gen8324 := Call(__e, ShenFunc(symhd), V1928)

				gen8325 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), gen8324)

				var gen8326 Obj
				if True == gen8325 {
					gen8322 := Call(__e, ShenFunc(symtl), V1928)

					gen8323 := Call(__e, ShenFunc(symcons_2), gen8322)

					if True == gen8323 {
						gen8326 = True
					} else {
						gen8326 = False
					}

				} else {
					gen8326 = False
				}
				if True == gen8326 {
					gen8328 = True
				} else {
					gen8328 = False
				}

			} else {
				gen8328 = False
			}
			if True == gen8328 {
				gen8318 := Call(__e, ShenFunc(symtl), V1928)

				gen8319 := Call(__e, ShenFunc(symhd), gen8318)

				gen8320 := Call(__e, ShenFunc(symtl), V1928)

				gen8321 := Call(__e, ShenFunc(symtl), gen8320)

				__e.TailApply(ShenFunc(symshen_4shen_1_6kl), gen8319, gen8321)

				return

			} else {
				gen8316 := Call(__e, ShenFunc(symcons_2), V1928)

				var gen8317 Obj
				if True == gen8316 {
					gen8313 := Call(__e, ShenFunc(symhd), V1928)

					gen8314 := Call(__e, ShenFunc(sym_a), MakeSymbol("defmacro"), gen8313)

					var gen8315 Obj
					if True == gen8314 {
						gen8311 := Call(__e, ShenFunc(symtl), V1928)

						gen8312 := Call(__e, ShenFunc(symcons_2), gen8311)

						if True == gen8312 {
							gen8315 = True
						} else {
							gen8315 = False
						}

					} else {
						gen8315 = False
					}
					if True == gen8315 {
						gen8317 = True
					} else {
						gen8317 = False
					}

				} else {
					gen8317 = False
				}
				if True == gen8317 {
					gen8297 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), Nil)

					gen8298 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen8297)

					gen8299 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen8298)

					Default := gen8299
					gen8300 := Call(__e, ShenFunc(symtl), V1928)

					gen8301 := Call(__e, ShenFunc(symhd), gen8300)

					gen8302 := Call(__e, ShenFunc(symtl), V1928)

					gen8303 := Call(__e, ShenFunc(symtl), gen8302)

					gen8304 := Call(__e, ShenFunc(symappend), gen8303, Default)

					gen8305 := Call(__e, ShenFunc(symcons), gen8301, gen8304)

					gen8306 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen8305)

					gen8307 := Call(__e, ShenFunc(symshen_4elim_1def), gen8306)

					Def := gen8307
					gen8308 := Call(__e, ShenFunc(symtl), V1928)

					gen8309 := Call(__e, ShenFunc(symhd), gen8308)

					gen8310 := Call(__e, ShenFunc(symshen_4add_1macro), gen8309)

					MacroAdd := gen8310
					_ = MacroAdd
					__e.Return(Def)
					return

				} else {
					gen8295 := Call(__e, ShenFunc(symcons_2), V1928)

					var gen8296 Obj
					if True == gen8295 {
						gen8292 := Call(__e, ShenFunc(symhd), V1928)

						gen8293 := Call(__e, ShenFunc(sym_a), MakeSymbol("defcc"), gen8292)

						var gen8294 Obj
						if True == gen8293 {
							gen8290 := Call(__e, ShenFunc(symtl), V1928)

							gen8291 := Call(__e, ShenFunc(symcons_2), gen8290)

							if True == gen8291 {
								gen8294 = True
							} else {
								gen8294 = False
							}

						} else {
							gen8294 = False
						}
						if True == gen8294 {
							gen8296 = True
						} else {
							gen8296 = False
						}

					} else {
						gen8296 = False
					}
					if True == gen8296 {
						gen8289 := Call(__e, ShenFunc(symshen_4yacc), V1928)

						__e.TailApply(ShenFunc(symshen_4elim_1def), gen8289)

						return

					} else {
						gen8288 := Call(__e, ShenFunc(symcons_2), V1928)

						if True == gen8288 {
							gen8287 := MakeNative(func(__e Evaluator, __args ...Obj) {
								Z := __args[0]
								_ = Z
								__e.TailApply(ShenFunc(symshen_4elim_1def), Z)

								return
							}, 1)
							__e.TailApply(ShenFunc(symmap), gen8287, V1928)

							return

						} else {
							__e.Return(V1928)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.elim-def"), gen8329)

		gen8338 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1930 := __args[0]
			_ = V1930
			gen8330 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*macroreg*"))

			MacroReg := gen8330
			gen8331 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*macroreg*"))

			gen8332 := Call(__e, ShenFunc(symadjoin), V1930, gen8331)

			gen8333 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*macroreg*"), gen8332)

			NewMacroReg := gen8333
			gen8337 := Call(__e, ShenFunc(sym_a), MacroReg, NewMacroReg)

			if True == gen8337 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen8334 := Call(__e, ShenFunc(symfunction), V1930)

				gen8335 := Call(__e, ShenFunc(symvalue), MakeSymbol("*macros*"))

				gen8336 := Call(__e, ShenFunc(symcons), gen8334, gen8335)

				__e.TailApply(ShenFunc(symset), MakeSymbol("*macros*"), gen8336)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.add-macro"), gen8338)

		gen8350 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1938 := __args[0]
			_ = V1938
			gen8348 := Call(__e, ShenFunc(symcons_2), V1938)

			var gen8349 Obj
			if True == gen8348 {
				gen8345 := Call(__e, ShenFunc(symhd), V1938)

				gen8346 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen8345)

				var gen8347 Obj
				if True == gen8346 {
					gen8342 := Call(__e, ShenFunc(symtl), V1938)

					gen8343 := Call(__e, ShenFunc(symcons_2), gen8342)

					var gen8344 Obj
					if True == gen8343 {
						gen8339 := Call(__e, ShenFunc(symtl), V1938)

						gen8340 := Call(__e, ShenFunc(symtl), gen8339)

						gen8341 := Call(__e, ShenFunc(symcons_2), gen8340)

						if True == gen8341 {
							gen8344 = True
						} else {
							gen8344 = False
						}

					} else {
						gen8344 = False
					}
					if True == gen8344 {
						gen8347 = True
					} else {
						gen8347 = False
					}

				} else {
					gen8347 = False
				}
				if True == gen8347 {
					gen8349 = True
				} else {
					gen8349 = False
				}

			} else {
				gen8349 = False
			}
			if True == gen8349 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.packaged?"), gen8350)

		gen8356 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1940 := __args[0]
			_ = V1940
			gen8353 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				gen8351 := Call(__e, ShenFunc(symshen_4app), V1940, MakeString(" has not been used.\n"), MakeSymbol("shen.a"))

				gen8352 := Call(__e, ShenFunc(symcn), MakeString("package "), gen8351)

				__e.TailApply(ShenFunc(symsimple_1error), gen8352)

				return

			}, 1)
			gen8355 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8354 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1940, MakeSymbol("shen.external-symbols"), gen8354)

				return

			}, 0)
			__e.Return(Try(__e, gen8355).Catch(gen8353))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("external"), gen8356)

		gen8362 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1942 := __args[0]
			_ = V1942
			gen8359 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				gen8357 := Call(__e, ShenFunc(symshen_4app), V1942, MakeString(" has not been used.\n"), MakeSymbol("shen.a"))

				gen8358 := Call(__e, ShenFunc(symcn), MakeString("package "), gen8357)

				__e.TailApply(ShenFunc(symsimple_1error), gen8358)

				return

			}, 1)
			gen8361 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8360 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1942, MakeSymbol("shen.internal-symbols"), gen8360)

				return

			}, 0)
			__e.Return(Try(__e, gen8361).Catch(gen8359))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("internal"), gen8362)

		gen8405 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1946 := __args[0]
			_ = V1946
			gen8403 := Call(__e, ShenFunc(symcons_2), V1946)

			var gen8404 Obj
			if True == gen8403 {
				gen8400 := Call(__e, ShenFunc(symhd), V1946)

				gen8401 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen8400)

				var gen8402 Obj
				if True == gen8401 {
					gen8397 := Call(__e, ShenFunc(symtl), V1946)

					gen8398 := Call(__e, ShenFunc(symcons_2), gen8397)

					var gen8399 Obj
					if True == gen8398 {
						gen8393 := Call(__e, ShenFunc(symtl), V1946)

						gen8394 := Call(__e, ShenFunc(symhd), gen8393)

						gen8395 := Call(__e, ShenFunc(sym_a), MakeSymbol("null"), gen8394)

						var gen8396 Obj
						if True == gen8395 {
							gen8390 := Call(__e, ShenFunc(symtl), V1946)

							gen8391 := Call(__e, ShenFunc(symtl), gen8390)

							gen8392 := Call(__e, ShenFunc(symcons_2), gen8391)

							if True == gen8392 {
								gen8396 = True
							} else {
								gen8396 = False
							}

						} else {
							gen8396 = False
						}
						if True == gen8396 {
							gen8399 = True
						} else {
							gen8399 = False
						}

					} else {
						gen8399 = False
					}
					if True == gen8399 {
						gen8402 = True
					} else {
						gen8402 = False
					}

				} else {
					gen8402 = False
				}
				if True == gen8402 {
					gen8404 = True
				} else {
					gen8404 = False
				}

			} else {
				gen8404 = False
			}
			if True == gen8404 {
				gen8388 := Call(__e, ShenFunc(symtl), V1946)

				gen8389 := Call(__e, ShenFunc(symtl), gen8388)

				__e.TailApply(ShenFunc(symtl), gen8389)

				return

			} else {
				gen8386 := Call(__e, ShenFunc(symcons_2), V1946)

				var gen8387 Obj
				if True == gen8386 {
					gen8383 := Call(__e, ShenFunc(symhd), V1946)

					gen8384 := Call(__e, ShenFunc(sym_a), MakeSymbol("package"), gen8383)

					var gen8385 Obj
					if True == gen8384 {
						gen8380 := Call(__e, ShenFunc(symtl), V1946)

						gen8381 := Call(__e, ShenFunc(symcons_2), gen8380)

						var gen8382 Obj
						if True == gen8381 {
							gen8377 := Call(__e, ShenFunc(symtl), V1946)

							gen8378 := Call(__e, ShenFunc(symtl), gen8377)

							gen8379 := Call(__e, ShenFunc(symcons_2), gen8378)

							if True == gen8379 {
								gen8382 = True
							} else {
								gen8382 = False
							}

						} else {
							gen8382 = False
						}
						if True == gen8382 {
							gen8385 = True
						} else {
							gen8385 = False
						}

					} else {
						gen8385 = False
					}
					if True == gen8385 {
						gen8387 = True
					} else {
						gen8387 = False
					}

				} else {
					gen8387 = False
				}
				if True == gen8387 {
					gen8363 := Call(__e, ShenFunc(symtl), V1946)

					gen8364 := Call(__e, ShenFunc(symhd), gen8363)

					gen8365 := Call(__e, ShenFunc(symstr), gen8364)

					gen8366 := Call(__e, ShenFunc(symcn), gen8365, MakeString("."))

					gen8367 := Call(__e, ShenFunc(symintern), gen8366)

					PackageNameDot := gen8367
					gen8368 := Call(__e, ShenFunc(symexplode), PackageNameDot)

					ExpPackageNameDot := gen8368
					gen8369 := Call(__e, ShenFunc(symtl), V1946)

					gen8370 := Call(__e, ShenFunc(symhd), gen8369)

					gen8371 := Call(__e, ShenFunc(symtl), V1946)

					gen8372 := Call(__e, ShenFunc(symtl), gen8371)

					gen8373 := Call(__e, ShenFunc(symhd), gen8372)

					gen8374 := Call(__e, ShenFunc(symtl), V1946)

					gen8375 := Call(__e, ShenFunc(symtl), gen8374)

					gen8376 := Call(__e, ShenFunc(symtl), gen8375)

					__e.TailApply(ShenFunc(symshen_4packageh), gen8370, gen8373, gen8376, ExpPackageNameDot)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.package-contents"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.package-contents"), gen8405)

		gen8409 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1949 := __args[0]
			_ = V1949
			V1950 := __args[1]
			_ = V1950
			gen8408 := Call(__e, ShenFunc(symcons_2), V1950)

			if True == gen8408 {
				gen8406 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Z := __args[0]
					_ = Z
					__e.TailApply(ShenFunc(symshen_4walk), V1949, Z)

					return
				}, 1)
				gen8407 := Call(__e, ShenFunc(symmap), gen8406, V1950)

				__e.TailApply(V1949, gen8407)

				return

			} else {
				__e.TailApply(V1949, V1950)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.walk"), gen8409)

		gen8419 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1954 := __args[0]
			_ = V1954
			V1955 := __args[1]
			_ = V1955
			V1956 := __args[2]
			_ = V1956
			gen8410 := Call(__e, ShenFunc(symcons), Nil, Nil)

			gen8411 := Call(__e, ShenFunc(symcons), V1955, gen8410)

			gen8412 := Call(__e, V1954, gen8411)

			O := gen8412
			gen8416 := Call(__e, ShenFunc(symfail))

			gen8417 := Call(__e, ShenFunc(sym_a), gen8416, O)

			var gen8418 Obj
			if True == gen8417 {
				gen8418 = True
			} else {
				gen8413 := Call(__e, ShenFunc(symhd), O)

				gen8414 := Call(__e, ShenFunc(symempty_2), gen8413)

				gen8415 := Call(__e, ShenFunc(symnot), gen8414)

				if True == gen8415 {
					gen8418 = True
				} else {
					gen8418 = False
				}

			}
			if True == gen8418 {
				__e.TailApply(V1956, O)

				return
			} else {
				__e.TailApply(ShenFunc(symshen_4hdtl), O)

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("compile"), gen8419)

		gen8421 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1959 := __args[0]
			_ = V1959
			V1960 := __args[1]
			_ = V1960
			gen8420 := Call(__e, V1959, V1960)

			if True == gen8420 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				__e.Return(V1960)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fail-if"), gen8421)

		gen8422 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1963 := __args[0]
			_ = V1963
			V1964 := __args[1]
			_ = V1964
			__e.TailApply(ShenFunc(symcn), V1963, V1964)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("@s"), gen8422)

		gen8423 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tc?"), gen8423)

		gen8428 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1966 := __args[0]
			_ = V1966
			gen8425 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				gen8424 := Call(__e, ShenFunc(symshen_4app), V1966, MakeString(" not found.\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen8424)

				return

			}, 1)
			gen8427 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8426 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V1966, MakeSymbol("shen.source"), gen8426)

				return

			}, 0)
			__e.Return(Try(__e, gen8427).Catch(gen8425))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("ps"), gen8428)

		gen8429 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*stinput*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("stinput"), gen8429)

		gen8436 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1968 := __args[0]
			_ = V1968
			gen8430 := Call(__e, ShenFunc(sym_7), V1968, MakeNumber(1))

			gen8431 := Call(__e, ShenFunc(symabsvector), gen8430)

			Vector := gen8431
			gen8432 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(0), V1968)

			ZeroStamp := gen8432
			gen8434 := Call(__e, ShenFunc(sym_a), V1968, MakeNumber(0))

			var gen8435 Obj
			if True == gen8434 {
				gen8435 = ZeroStamp
			} else {
				gen8433 := Call(__e, ShenFunc(symfail))

				gen8435 = Call(__e, ShenFunc(symshen_4fillvector), ZeroStamp, MakeNumber(1), V1968, gen8433)

			}
			Standard := gen8435
			__e.Return(Standard)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("vector"), gen8436)

		gen8440 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1974 := __args[0]
			_ = V1974
			V1975 := __args[1]
			_ = V1975
			V1976 := __args[2]
			_ = V1976
			V1977 := __args[3]
			_ = V1977
			gen8439 := Call(__e, ShenFunc(sym_a), V1976, V1975)

			if True == gen8439 {
				__e.TailApply(ShenFunc(symaddress_1_6), V1974, V1976, V1977)

				return
			} else {
				gen8437 := Call(__e, ShenFunc(symaddress_1_6), V1974, V1975, V1977)

				gen8438 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1975)

				__e.TailApply(ShenFunc(symshen_4fillvector), gen8437, gen8438, V1976, V1977)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fillvector"), gen8440)

		gen8448 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1979 := __args[0]
			_ = V1979
			gen8447 := Call(__e, ShenFunc(symabsvector_2), V1979)

			if True == gen8447 {
				gen8441 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeNumber(-1))
					return
				}, 1)
				gen8442 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(sym_5_1address), V1979, MakeNumber(0))

					return
				}, 0)
				gen8443 := Try(__e, gen8442).Catch(gen8441)
				X := gen8443
				gen8445 := Call(__e, ShenFunc(symnumber_2), X)

				var gen8446 Obj
				if True == gen8445 {
					gen8444 := Call(__e, ShenFunc(sym_6_a), X, MakeNumber(0))

					if True == gen8444 {
						gen8446 = True
					} else {
						gen8446 = False
					}

				} else {
					gen8446 = False
				}
				if True == gen8446 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("vector?"), gen8448)

		gen8450 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1983 := __args[0]
			_ = V1983
			V1984 := __args[1]
			_ = V1984
			V1985 := __args[2]
			_ = V1985
			gen8449 := Call(__e, ShenFunc(sym_a), V1984, MakeNumber(0))

			if True == gen8449 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("cannot access 0th element of a vector\n"))

				return
			} else {
				__e.TailApply(ShenFunc(symaddress_1_6), V1983, V1984, V1985)

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("vector->"), gen8450)

		gen8455 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1988 := __args[0]
			_ = V1988
			V1989 := __args[1]
			_ = V1989
			gen8454 := Call(__e, ShenFunc(sym_a), V1989, MakeNumber(0))

			if True == gen8454 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("cannot access 0th element of a vector\n"))

				return
			} else {
				gen8451 := Call(__e, ShenFunc(sym_5_1address), V1988, V1989)

				VectorElement := gen8451
				gen8452 := Call(__e, ShenFunc(symfail))

				gen8453 := Call(__e, ShenFunc(sym_a), VectorElement, gen8452)

				if True == gen8453 {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("vector element not found\n"))

					return
				} else {
					__e.Return(VectorElement)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("<-vector"), gen8455)

		gen8458 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1991 := __args[0]
			_ = V1991
			gen8457 := Call(__e, ShenFunc(syminteger_2), V1991)

			if True == gen8457 {
				gen8456 := Call(__e, ShenFunc(sym_6_a), V1991, MakeNumber(0))

				if True == gen8456 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.posint?"), gen8458)

		gen8459 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1993 := __args[0]
			_ = V1993
			__e.TailApply(ShenFunc(sym_5_1address), V1993, MakeNumber(0))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("limit"), gen8459)

		gen8468 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1995 := __args[0]
			_ = V1995
			gen8466 := Call(__e, ShenFunc(symboolean_2), V1995)

			var gen8467 Obj
			if True == gen8466 {
				gen8467 = True
			} else {
				gen8464 := Call(__e, ShenFunc(symnumber_2), V1995)

				var gen8465 Obj
				if True == gen8464 {
					gen8465 = True
				} else {
					gen8463 := Call(__e, ShenFunc(symstring_2), V1995)

					if True == gen8463 {
						gen8465 = True
					} else {
						gen8465 = False
					}

				}
				if True == gen8465 {
					gen8467 = True
				} else {
					gen8467 = False
				}

			}
			if True == gen8467 {
				__e.Return(False)
				return
			} else {
				gen8460 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(False)
					return
				}, 1)
				gen8462 := MakeNative(func(__e Evaluator, __args ...Obj) {
					gen8461 := Call(__e, ShenFunc(symstr), V1995)

					String := gen8461
					__e.TailApply(ShenFunc(symshen_4analyse_1symbol_2), String)

					return

				}, 0)
				__e.Return(Try(__e, gen8462).Catch(gen8460))
				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("symbol?"), gen8468)

		gen8475 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1997 := __args[0]
			_ = V1997
			gen8474 := Call(__e, ShenFunc(sym_a), MakeString(""), V1997)

			if True == gen8474 {
				__e.Return(False)
				return
			} else {
				gen8473 := Call(__e, ShenFunc(symshen_4_7string_2), V1997)

				if True == gen8473 {
					gen8471 := Call(__e, ShenFunc(sympos), V1997, MakeNumber(0))

					gen8472 := Call(__e, ShenFunc(symshen_4alpha_2), gen8471)

					if True == gen8472 {
						gen8469 := Call(__e, ShenFunc(symtlstr), V1997)

						gen8470 := Call(__e, ShenFunc(symshen_4alphanums_2), gen8469)

						if True == gen8470 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.analyse-symbol?"), gen8475)

		gen8551 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1999 := __args[0]
			_ = V1999
			gen8476 := Call(__e, ShenFunc(symcons), MakeString("."), Nil)

			gen8477 := Call(__e, ShenFunc(symcons), MakeString("'"), gen8476)

			gen8478 := Call(__e, ShenFunc(symcons), MakeString("#"), gen8477)

			gen8479 := Call(__e, ShenFunc(symcons), MakeString("`"), gen8478)

			gen8480 := Call(__e, ShenFunc(symcons), MakeString(";"), gen8479)

			gen8481 := Call(__e, ShenFunc(symcons), MakeString(":"), gen8480)

			gen8482 := Call(__e, ShenFunc(symcons), MakeString("}"), gen8481)

			gen8483 := Call(__e, ShenFunc(symcons), MakeString("{"), gen8482)

			gen8484 := Call(__e, ShenFunc(symcons), MakeString("%"), gen8483)

			gen8485 := Call(__e, ShenFunc(symcons), MakeString("&"), gen8484)

			gen8486 := Call(__e, ShenFunc(symcons), MakeString("<"), gen8485)

			gen8487 := Call(__e, ShenFunc(symcons), MakeString(">"), gen8486)

			gen8488 := Call(__e, ShenFunc(symcons), MakeString("~"), gen8487)

			gen8489 := Call(__e, ShenFunc(symcons), MakeString("@"), gen8488)

			gen8490 := Call(__e, ShenFunc(symcons), MakeString("!"), gen8489)

			gen8491 := Call(__e, ShenFunc(symcons), MakeString("$"), gen8490)

			gen8492 := Call(__e, ShenFunc(symcons), MakeString("?"), gen8491)

			gen8493 := Call(__e, ShenFunc(symcons), MakeString("_"), gen8492)

			gen8494 := Call(__e, ShenFunc(symcons), MakeString("-"), gen8493)

			gen8495 := Call(__e, ShenFunc(symcons), MakeString("+"), gen8494)

			gen8496 := Call(__e, ShenFunc(symcons), MakeString("/"), gen8495)

			gen8497 := Call(__e, ShenFunc(symcons), MakeString("*"), gen8496)

			gen8498 := Call(__e, ShenFunc(symcons), MakeString("="), gen8497)

			gen8499 := Call(__e, ShenFunc(symcons), MakeString("z"), gen8498)

			gen8500 := Call(__e, ShenFunc(symcons), MakeString("y"), gen8499)

			gen8501 := Call(__e, ShenFunc(symcons), MakeString("x"), gen8500)

			gen8502 := Call(__e, ShenFunc(symcons), MakeString("w"), gen8501)

			gen8503 := Call(__e, ShenFunc(symcons), MakeString("v"), gen8502)

			gen8504 := Call(__e, ShenFunc(symcons), MakeString("u"), gen8503)

			gen8505 := Call(__e, ShenFunc(symcons), MakeString("t"), gen8504)

			gen8506 := Call(__e, ShenFunc(symcons), MakeString("s"), gen8505)

			gen8507 := Call(__e, ShenFunc(symcons), MakeString("r"), gen8506)

			gen8508 := Call(__e, ShenFunc(symcons), MakeString("q"), gen8507)

			gen8509 := Call(__e, ShenFunc(symcons), MakeString("p"), gen8508)

			gen8510 := Call(__e, ShenFunc(symcons), MakeString("o"), gen8509)

			gen8511 := Call(__e, ShenFunc(symcons), MakeString("n"), gen8510)

			gen8512 := Call(__e, ShenFunc(symcons), MakeString("m"), gen8511)

			gen8513 := Call(__e, ShenFunc(symcons), MakeString("l"), gen8512)

			gen8514 := Call(__e, ShenFunc(symcons), MakeString("k"), gen8513)

			gen8515 := Call(__e, ShenFunc(symcons), MakeString("j"), gen8514)

			gen8516 := Call(__e, ShenFunc(symcons), MakeString("i"), gen8515)

			gen8517 := Call(__e, ShenFunc(symcons), MakeString("h"), gen8516)

			gen8518 := Call(__e, ShenFunc(symcons), MakeString("g"), gen8517)

			gen8519 := Call(__e, ShenFunc(symcons), MakeString("f"), gen8518)

			gen8520 := Call(__e, ShenFunc(symcons), MakeString("e"), gen8519)

			gen8521 := Call(__e, ShenFunc(symcons), MakeString("d"), gen8520)

			gen8522 := Call(__e, ShenFunc(symcons), MakeString("c"), gen8521)

			gen8523 := Call(__e, ShenFunc(symcons), MakeString("b"), gen8522)

			gen8524 := Call(__e, ShenFunc(symcons), MakeString("a"), gen8523)

			gen8525 := Call(__e, ShenFunc(symcons), MakeString("Z"), gen8524)

			gen8526 := Call(__e, ShenFunc(symcons), MakeString("Y"), gen8525)

			gen8527 := Call(__e, ShenFunc(symcons), MakeString("X"), gen8526)

			gen8528 := Call(__e, ShenFunc(symcons), MakeString("W"), gen8527)

			gen8529 := Call(__e, ShenFunc(symcons), MakeString("V"), gen8528)

			gen8530 := Call(__e, ShenFunc(symcons), MakeString("U"), gen8529)

			gen8531 := Call(__e, ShenFunc(symcons), MakeString("T"), gen8530)

			gen8532 := Call(__e, ShenFunc(symcons), MakeString("S"), gen8531)

			gen8533 := Call(__e, ShenFunc(symcons), MakeString("R"), gen8532)

			gen8534 := Call(__e, ShenFunc(symcons), MakeString("Q"), gen8533)

			gen8535 := Call(__e, ShenFunc(symcons), MakeString("P"), gen8534)

			gen8536 := Call(__e, ShenFunc(symcons), MakeString("O"), gen8535)

			gen8537 := Call(__e, ShenFunc(symcons), MakeString("N"), gen8536)

			gen8538 := Call(__e, ShenFunc(symcons), MakeString("M"), gen8537)

			gen8539 := Call(__e, ShenFunc(symcons), MakeString("L"), gen8538)

			gen8540 := Call(__e, ShenFunc(symcons), MakeString("K"), gen8539)

			gen8541 := Call(__e, ShenFunc(symcons), MakeString("J"), gen8540)

			gen8542 := Call(__e, ShenFunc(symcons), MakeString("I"), gen8541)

			gen8543 := Call(__e, ShenFunc(symcons), MakeString("H"), gen8542)

			gen8544 := Call(__e, ShenFunc(symcons), MakeString("G"), gen8543)

			gen8545 := Call(__e, ShenFunc(symcons), MakeString("F"), gen8544)

			gen8546 := Call(__e, ShenFunc(symcons), MakeString("E"), gen8545)

			gen8547 := Call(__e, ShenFunc(symcons), MakeString("D"), gen8546)

			gen8548 := Call(__e, ShenFunc(symcons), MakeString("C"), gen8547)

			gen8549 := Call(__e, ShenFunc(symcons), MakeString("B"), gen8548)

			gen8550 := Call(__e, ShenFunc(symcons), MakeString("A"), gen8549)

			__e.TailApply(ShenFunc(symelement_2), V1999, gen8550)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.alpha?"), gen8551)

		gen8558 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2001 := __args[0]
			_ = V2001
			gen8557 := Call(__e, ShenFunc(sym_a), MakeString(""), V2001)

			if True == gen8557 {
				__e.Return(True)
				return
			} else {
				gen8556 := Call(__e, ShenFunc(symshen_4_7string_2), V2001)

				if True == gen8556 {
					gen8554 := Call(__e, ShenFunc(sympos), V2001, MakeNumber(0))

					gen8555 := Call(__e, ShenFunc(symshen_4alphanum_2), gen8554)

					if True == gen8555 {
						gen8552 := Call(__e, ShenFunc(symtlstr), V2001)

						gen8553 := Call(__e, ShenFunc(symshen_4alphanums_2), gen8552)

						if True == gen8553 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.alphanums?"), gen8558)

		gen8561 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2003 := __args[0]
			_ = V2003
			gen8560 := Call(__e, ShenFunc(symshen_4alpha_2), V2003)

			if True == gen8560 {
				__e.Return(True)
				return
			} else {
				gen8559 := Call(__e, ShenFunc(symshen_4digit_2), V2003)

				if True == gen8559 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.alphanum?"), gen8561)

		gen8572 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2005 := __args[0]
			_ = V2005
			gen8562 := Call(__e, ShenFunc(symcons), MakeString("0"), Nil)

			gen8563 := Call(__e, ShenFunc(symcons), MakeString("9"), gen8562)

			gen8564 := Call(__e, ShenFunc(symcons), MakeString("8"), gen8563)

			gen8565 := Call(__e, ShenFunc(symcons), MakeString("7"), gen8564)

			gen8566 := Call(__e, ShenFunc(symcons), MakeString("6"), gen8565)

			gen8567 := Call(__e, ShenFunc(symcons), MakeString("5"), gen8566)

			gen8568 := Call(__e, ShenFunc(symcons), MakeString("4"), gen8567)

			gen8569 := Call(__e, ShenFunc(symcons), MakeString("3"), gen8568)

			gen8570 := Call(__e, ShenFunc(symcons), MakeString("2"), gen8569)

			gen8571 := Call(__e, ShenFunc(symcons), MakeString("1"), gen8570)

			__e.TailApply(ShenFunc(symelement_2), V2005, gen8571)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.digit?"), gen8572)

		gen8581 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2007 := __args[0]
			_ = V2007
			gen8579 := Call(__e, ShenFunc(symboolean_2), V2007)

			var gen8580 Obj
			if True == gen8579 {
				gen8580 = True
			} else {
				gen8577 := Call(__e, ShenFunc(symnumber_2), V2007)

				var gen8578 Obj
				if True == gen8577 {
					gen8578 = True
				} else {
					gen8576 := Call(__e, ShenFunc(symstring_2), V2007)

					if True == gen8576 {
						gen8578 = True
					} else {
						gen8578 = False
					}

				}
				if True == gen8578 {
					gen8580 = True
				} else {
					gen8580 = False
				}

			}
			if True == gen8580 {
				__e.Return(False)
				return
			} else {
				gen8573 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(False)
					return
				}, 1)
				gen8575 := MakeNative(func(__e Evaluator, __args ...Obj) {
					gen8574 := Call(__e, ShenFunc(symstr), V2007)

					String := gen8574
					__e.TailApply(ShenFunc(symshen_4analyse_1variable_2), String)

					return

				}, 0)
				__e.Return(Try(__e, gen8575).Catch(gen8573))
				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("variable?"), gen8581)

		gen8587 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2009 := __args[0]
			_ = V2009
			gen8586 := Call(__e, ShenFunc(symshen_4_7string_2), V2009)

			if True == gen8586 {
				gen8584 := Call(__e, ShenFunc(sympos), V2009, MakeNumber(0))

				gen8585 := Call(__e, ShenFunc(symshen_4uppercase_2), gen8584)

				if True == gen8585 {
					gen8582 := Call(__e, ShenFunc(symtlstr), V2009)

					gen8583 := Call(__e, ShenFunc(symshen_4alphanums_2), gen8582)

					if True == gen8583 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.analyse-variable?"), gen8587)

		gen8614 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2011 := __args[0]
			_ = V2011
			gen8588 := Call(__e, ShenFunc(symcons), MakeString("Z"), Nil)

			gen8589 := Call(__e, ShenFunc(symcons), MakeString("Y"), gen8588)

			gen8590 := Call(__e, ShenFunc(symcons), MakeString("X"), gen8589)

			gen8591 := Call(__e, ShenFunc(symcons), MakeString("W"), gen8590)

			gen8592 := Call(__e, ShenFunc(symcons), MakeString("V"), gen8591)

			gen8593 := Call(__e, ShenFunc(symcons), MakeString("U"), gen8592)

			gen8594 := Call(__e, ShenFunc(symcons), MakeString("T"), gen8593)

			gen8595 := Call(__e, ShenFunc(symcons), MakeString("S"), gen8594)

			gen8596 := Call(__e, ShenFunc(symcons), MakeString("R"), gen8595)

			gen8597 := Call(__e, ShenFunc(symcons), MakeString("Q"), gen8596)

			gen8598 := Call(__e, ShenFunc(symcons), MakeString("P"), gen8597)

			gen8599 := Call(__e, ShenFunc(symcons), MakeString("O"), gen8598)

			gen8600 := Call(__e, ShenFunc(symcons), MakeString("N"), gen8599)

			gen8601 := Call(__e, ShenFunc(symcons), MakeString("M"), gen8600)

			gen8602 := Call(__e, ShenFunc(symcons), MakeString("L"), gen8601)

			gen8603 := Call(__e, ShenFunc(symcons), MakeString("K"), gen8602)

			gen8604 := Call(__e, ShenFunc(symcons), MakeString("J"), gen8603)

			gen8605 := Call(__e, ShenFunc(symcons), MakeString("I"), gen8604)

			gen8606 := Call(__e, ShenFunc(symcons), MakeString("H"), gen8605)

			gen8607 := Call(__e, ShenFunc(symcons), MakeString("G"), gen8606)

			gen8608 := Call(__e, ShenFunc(symcons), MakeString("F"), gen8607)

			gen8609 := Call(__e, ShenFunc(symcons), MakeString("E"), gen8608)

			gen8610 := Call(__e, ShenFunc(symcons), MakeString("D"), gen8609)

			gen8611 := Call(__e, ShenFunc(symcons), MakeString("C"), gen8610)

			gen8612 := Call(__e, ShenFunc(symcons), MakeString("B"), gen8611)

			gen8613 := Call(__e, ShenFunc(symcons), MakeString("A"), gen8612)

			__e.TailApply(ShenFunc(symelement_2), V2011, gen8613)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.uppercase?"), gen8614)

		gen8618 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2013 := __args[0]
			_ = V2013
			gen8615 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*gensym*"))

			gen8616 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen8615)

			gen8617 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*gensym*"), gen8616)

			__e.TailApply(ShenFunc(symconcat), V2013, gen8617)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("gensym"), gen8618)

		gen8622 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2016 := __args[0]
			_ = V2016
			V2017 := __args[1]
			_ = V2017
			gen8619 := Call(__e, ShenFunc(symstr), V2016)

			gen8620 := Call(__e, ShenFunc(symstr), V2017)

			gen8621 := Call(__e, ShenFunc(symcn), gen8619, gen8620)

			__e.TailApply(ShenFunc(symintern), gen8621)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("concat"), gen8622)

		gen8627 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2020 := __args[0]
			_ = V2020
			V2021 := __args[1]
			_ = V2021
			gen8623 := Call(__e, ShenFunc(symabsvector), MakeNumber(3))

			Vector := gen8623
			gen8624 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(0), MakeSymbol("shen.tuple"))

			Tag := gen8624
			_ = Tag
			gen8625 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(1), V2020)

			Fst := gen8625
			_ = Fst
			gen8626 := Call(__e, ShenFunc(symaddress_1_6), Vector, MakeNumber(2), V2021)

			Snd := gen8626
			_ = Snd
			__e.Return(Vector)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("@p"), gen8627)

		gen8628 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2023 := __args[0]
			_ = V2023
			__e.TailApply(ShenFunc(sym_5_1address), V2023, MakeNumber(1))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fst"), gen8628)

		gen8629 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2025 := __args[0]
			_ = V2025
			__e.TailApply(ShenFunc(sym_5_1address), V2025, MakeNumber(2))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("snd"), gen8629)

		gen8635 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2027 := __args[0]
			_ = V2027
			gen8634 := Call(__e, ShenFunc(symabsvector_2), V2027)

			if True == gen8634 {
				gen8630 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeSymbol("shen.not-tuple"))
					return
				}, 1)
				gen8631 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(sym_5_1address), V2027, MakeNumber(0))

					return
				}, 0)
				gen8632 := Try(__e, gen8631).Catch(gen8630)
				gen8633 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tuple"), gen8632)

				if True == gen8633 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("tuple?"), gen8635)

		gen8641 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2030 := __args[0]
			_ = V2030
			V2031 := __args[1]
			_ = V2031
			gen8640 := Call(__e, ShenFunc(sym_a), Nil, V2030)

			if True == gen8640 {
				__e.Return(V2031)
				return
			} else {
				gen8639 := Call(__e, ShenFunc(symcons_2), V2030)

				if True == gen8639 {
					gen8636 := Call(__e, ShenFunc(symhd), V2030)

					gen8637 := Call(__e, ShenFunc(symtl), V2030)

					gen8638 := Call(__e, ShenFunc(symappend), gen8637, V2031)

					__e.TailApply(ShenFunc(symcons), gen8636, gen8638)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("append"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("append"), gen8641)

		gen8647 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2034 := __args[0]
			_ = V2034
			V2035 := __args[1]
			_ = V2035
			gen8642 := Call(__e, ShenFunc(symlimit), V2035)

			Limit := gen8642
			gen8643 := Call(__e, ShenFunc(sym_7), Limit, MakeNumber(1))

			gen8644 := Call(__e, ShenFunc(symvector), gen8643)

			NewVector := gen8644
			gen8645 := Call(__e, ShenFunc(symvector_1_6), NewVector, MakeNumber(1), V2034)

			X_7NewVector := gen8645
			gen8646 := Call(__e, ShenFunc(sym_a), Limit, MakeNumber(0))

			if True == gen8646 {
				__e.Return(X_7NewVector)
				return
			} else {
				__e.TailApply(ShenFunc(symshen_4_8v_1help), V2035, MakeNumber(1), Limit, X_7NewVector)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("@v"), gen8647)

		gen8653 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2041 := __args[0]
			_ = V2041
			V2042 := __args[1]
			_ = V2042
			V2043 := __args[2]
			_ = V2043
			V2044 := __args[3]
			_ = V2044
			gen8652 := Call(__e, ShenFunc(sym_a), V2043, V2042)

			if True == gen8652 {
				gen8651 := Call(__e, ShenFunc(sym_7), V2043, MakeNumber(1))

				__e.TailApply(ShenFunc(symshen_4copyfromvector), V2041, V2044, V2043, gen8651)

				return

			} else {
				gen8648 := Call(__e, ShenFunc(sym_7), V2042, MakeNumber(1))

				gen8649 := Call(__e, ShenFunc(sym_7), V2042, MakeNumber(1))

				gen8650 := Call(__e, ShenFunc(symshen_4copyfromvector), V2041, V2044, V2042, gen8649)

				__e.TailApply(ShenFunc(symshen_4_8v_1help), V2041, gen8648, V2043, gen8650)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.@v-help"), gen8653)

		gen8657 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2049 := __args[0]
			_ = V2049
			V2050 := __args[1]
			_ = V2050
			V2051 := __args[2]
			_ = V2051
			V2052 := __args[3]
			_ = V2052
			gen8654 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(V2050)
				return
			}, 1)
			gen8656 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8655 := Call(__e, ShenFunc(sym_5_1vector), V2049, V2051)

				__e.TailApply(ShenFunc(symvector_1_6), V2050, V2052, gen8655)

				return

			}, 0)
			__e.Return(Try(__e, gen8656).Catch(gen8654))
			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copyfromvector"), gen8657)

		gen8662 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2054 := __args[0]
			_ = V2054
			gen8660 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				gen8658 := Call(__e, ShenFunc(symshen_4app), V2054, MakeString("\n"), MakeSymbol("shen.s"))

				gen8659 := Call(__e, ShenFunc(symcn), MakeString("hdv needs a non-empty vector as an argument; not "), gen8658)

				__e.TailApply(ShenFunc(symsimple_1error), gen8659)

				return

			}, 1)
			gen8661 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(sym_5_1vector), V2054, MakeNumber(1))

				return
			}, 0)
			__e.Return(Try(__e, gen8661).Catch(gen8660))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("hdv"), gen8662)

		gen8670 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2056 := __args[0]
			_ = V2056
			gen8663 := Call(__e, ShenFunc(symlimit), V2056)

			Limit := gen8663
			gen8669 := Call(__e, ShenFunc(sym_a), Limit, MakeNumber(0))

			if True == gen8669 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("cannot take the tail of the empty vector\n"))

				return
			} else {
				gen8668 := Call(__e, ShenFunc(sym_a), Limit, MakeNumber(1))

				if True == gen8668 {
					__e.TailApply(ShenFunc(symvector), MakeNumber(0))

					return
				} else {
					gen8664 := Call(__e, ShenFunc(sym_1), Limit, MakeNumber(1))

					gen8665 := Call(__e, ShenFunc(symvector), gen8664)

					NewVector := gen8665
					_ = NewVector
					gen8666 := Call(__e, ShenFunc(sym_1), Limit, MakeNumber(1))

					gen8667 := Call(__e, ShenFunc(symvector), gen8666)

					__e.TailApply(ShenFunc(symshen_4tlv_1help), V2056, MakeNumber(2), Limit, gen8667)

					return

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tlv"), gen8670)

		gen8676 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2062 := __args[0]
			_ = V2062
			V2063 := __args[1]
			_ = V2063
			V2064 := __args[2]
			_ = V2064
			V2065 := __args[3]
			_ = V2065
			gen8675 := Call(__e, ShenFunc(sym_a), V2064, V2063)

			if True == gen8675 {
				gen8674 := Call(__e, ShenFunc(sym_1), V2064, MakeNumber(1))

				__e.TailApply(ShenFunc(symshen_4copyfromvector), V2062, V2065, V2064, gen8674)

				return

			} else {
				gen8671 := Call(__e, ShenFunc(sym_7), V2063, MakeNumber(1))

				gen8672 := Call(__e, ShenFunc(sym_1), V2063, MakeNumber(1))

				gen8673 := Call(__e, ShenFunc(symshen_4copyfromvector), V2062, V2065, V2063, gen8672)

				__e.TailApply(ShenFunc(symshen_4tlv_1help), V2062, gen8671, V2064, gen8673)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tlv-help"), gen8676)

		gen8688 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2077 := __args[0]
			_ = V2077
			V2078 := __args[1]
			_ = V2078
			gen8687 := Call(__e, ShenFunc(sym_a), Nil, V2078)

			if True == gen8687 {
				__e.Return(Nil)
				return
			} else {
				gen8685 := Call(__e, ShenFunc(symcons_2), V2078)

				var gen8686 Obj
				if True == gen8685 {
					gen8682 := Call(__e, ShenFunc(symhd), V2078)

					gen8683 := Call(__e, ShenFunc(symcons_2), gen8682)

					var gen8684 Obj
					if True == gen8683 {
						gen8679 := Call(__e, ShenFunc(symhd), V2078)

						gen8680 := Call(__e, ShenFunc(symhd), gen8679)

						gen8681 := Call(__e, ShenFunc(sym_a), gen8680, V2077)

						if True == gen8681 {
							gen8684 = True
						} else {
							gen8684 = False
						}

					} else {
						gen8684 = False
					}
					if True == gen8684 {
						gen8686 = True
					} else {
						gen8686 = False
					}

				} else {
					gen8686 = False
				}
				if True == gen8686 {
					__e.TailApply(ShenFunc(symhd), V2078)

					return
				} else {
					gen8678 := Call(__e, ShenFunc(symcons_2), V2078)

					if True == gen8678 {
						gen8677 := Call(__e, ShenFunc(symtl), V2078)

						__e.TailApply(ShenFunc(symassoc), V2077, gen8677)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("assoc"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("assoc"), gen8688)

		gen8707 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2085 := __args[0]
			_ = V2085
			V2086 := __args[1]
			_ = V2086
			V2087 := __args[2]
			_ = V2087
			gen8706 := Call(__e, ShenFunc(sym_a), Nil, V2087)

			if True == gen8706 {
				gen8705 := Call(__e, ShenFunc(symcons), V2085, V2086)

				__e.TailApply(ShenFunc(symcons), gen8705, Nil)

				return

			} else {
				gen8703 := Call(__e, ShenFunc(symcons_2), V2087)

				var gen8704 Obj
				if True == gen8703 {
					gen8700 := Call(__e, ShenFunc(symhd), V2087)

					gen8701 := Call(__e, ShenFunc(symcons_2), gen8700)

					var gen8702 Obj
					if True == gen8701 {
						gen8697 := Call(__e, ShenFunc(symhd), V2087)

						gen8698 := Call(__e, ShenFunc(symhd), gen8697)

						gen8699 := Call(__e, ShenFunc(sym_a), gen8698, V2085)

						if True == gen8699 {
							gen8702 = True
						} else {
							gen8702 = False
						}

					} else {
						gen8702 = False
					}
					if True == gen8702 {
						gen8704 = True
					} else {
						gen8704 = False
					}

				} else {
					gen8704 = False
				}
				if True == gen8704 {
					gen8693 := Call(__e, ShenFunc(symhd), V2087)

					gen8694 := Call(__e, ShenFunc(symhd), gen8693)

					gen8695 := Call(__e, ShenFunc(symcons), gen8694, V2086)

					gen8696 := Call(__e, ShenFunc(symtl), V2087)

					__e.TailApply(ShenFunc(symcons), gen8695, gen8696)

					return

				} else {
					gen8692 := Call(__e, ShenFunc(symcons_2), V2087)

					if True == gen8692 {
						gen8689 := Call(__e, ShenFunc(symhd), V2087)

						gen8690 := Call(__e, ShenFunc(symtl), V2087)

						gen8691 := Call(__e, ShenFunc(symshen_4assoc_1set), V2085, V2086, gen8690)

						__e.TailApply(ShenFunc(symcons), gen8689, gen8691)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.assoc-set"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assoc-set"), gen8707)

		gen8721 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2093 := __args[0]
			_ = V2093
			V2094 := __args[1]
			_ = V2094
			gen8720 := Call(__e, ShenFunc(sym_a), Nil, V2094)

			if True == gen8720 {
				__e.Return(Nil)
				return
			} else {
				gen8718 := Call(__e, ShenFunc(symcons_2), V2094)

				var gen8719 Obj
				if True == gen8718 {
					gen8715 := Call(__e, ShenFunc(symhd), V2094)

					gen8716 := Call(__e, ShenFunc(symcons_2), gen8715)

					var gen8717 Obj
					if True == gen8716 {
						gen8712 := Call(__e, ShenFunc(symhd), V2094)

						gen8713 := Call(__e, ShenFunc(symhd), gen8712)

						gen8714 := Call(__e, ShenFunc(sym_a), gen8713, V2093)

						if True == gen8714 {
							gen8717 = True
						} else {
							gen8717 = False
						}

					} else {
						gen8717 = False
					}
					if True == gen8717 {
						gen8719 = True
					} else {
						gen8719 = False
					}

				} else {
					gen8719 = False
				}
				if True == gen8719 {
					__e.TailApply(ShenFunc(symtl), V2094)

					return
				} else {
					gen8711 := Call(__e, ShenFunc(symcons_2), V2094)

					if True == gen8711 {
						gen8708 := Call(__e, ShenFunc(symhd), V2094)

						gen8709 := Call(__e, ShenFunc(symtl), V2094)

						gen8710 := Call(__e, ShenFunc(symshen_4assoc_1rm), V2093, gen8709)

						__e.TailApply(ShenFunc(symcons), gen8708, gen8710)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.assoc-rm"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assoc-rm"), gen8721)

		gen8724 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2100 := __args[0]
			_ = V2100
			gen8723 := Call(__e, ShenFunc(sym_a), True, V2100)

			if True == gen8723 {
				__e.Return(True)
				return
			} else {
				gen8722 := Call(__e, ShenFunc(sym_a), False, V2100)

				if True == gen8722 {
					__e.Return(True)
					return
				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("boolean?"), gen8724)

		gen8728 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2102 := __args[0]
			_ = V2102
			gen8727 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V2102)

			if True == gen8727 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen8725 := Call(__e, ShenFunc(symstoutput))

				Call(__e, ShenFunc(symshen_4prhush), MakeString("\n"), gen8725)

				gen8726 := Call(__e, ShenFunc(sym_1), V2102, MakeNumber(1))

				__e.TailApply(ShenFunc(symnl), gen8726)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("nl"), gen8728)

		gen8737 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2107 := __args[0]
			_ = V2107
			V2108 := __args[1]
			_ = V2108
			gen8736 := Call(__e, ShenFunc(sym_a), Nil, V2107)

			if True == gen8736 {
				__e.Return(Nil)
				return
			} else {
				gen8735 := Call(__e, ShenFunc(symcons_2), V2107)

				if True == gen8735 {
					gen8733 := Call(__e, ShenFunc(symhd), V2107)

					gen8734 := Call(__e, ShenFunc(symelement_2), gen8733, V2108)

					if True == gen8734 {
						gen8732 := Call(__e, ShenFunc(symtl), V2107)

						__e.TailApply(ShenFunc(symdifference), gen8732, V2108)

						return

					} else {
						gen8729 := Call(__e, ShenFunc(symhd), V2107)

						gen8730 := Call(__e, ShenFunc(symtl), V2107)

						gen8731 := Call(__e, ShenFunc(symdifference), gen8730, V2108)

						__e.TailApply(ShenFunc(symcons), gen8729, gen8731)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("difference"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("difference"), gen8737)

		gen8738 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2111 := __args[0]
			_ = V2111
			V2112 := __args[1]
			_ = V2112
			__e.Return(V2112)
			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("do"), gen8738)

		gen8746 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2124 := __args[0]
			_ = V2124
			V2125 := __args[1]
			_ = V2125
			gen8745 := Call(__e, ShenFunc(sym_a), Nil, V2125)

			if True == gen8745 {
				__e.Return(False)
				return
			} else {
				gen8743 := Call(__e, ShenFunc(symcons_2), V2125)

				var gen8744 Obj
				if True == gen8743 {
					gen8741 := Call(__e, ShenFunc(symhd), V2125)

					gen8742 := Call(__e, ShenFunc(sym_a), gen8741, V2124)

					if True == gen8742 {
						gen8744 = True
					} else {
						gen8744 = False
					}

				} else {
					gen8744 = False
				}
				if True == gen8744 {
					__e.Return(True)
					return
				} else {
					gen8740 := Call(__e, ShenFunc(symcons_2), V2125)

					if True == gen8740 {
						gen8739 := Call(__e, ShenFunc(symtl), V2125)

						__e.TailApply(ShenFunc(symelement_2), V2124, gen8739)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("element?"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("element?"), gen8746)

		gen8748 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2131 := __args[0]
			_ = V2131
			gen8747 := Call(__e, ShenFunc(sym_a), Nil, V2131)

			if True == gen8747 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("empty?"), gen8748)

		gen8750 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2134 := __args[0]
			_ = V2134
			V2135 := __args[1]
			_ = V2135
			gen8749 := Call(__e, V2134, V2135)

			__e.TailApply(ShenFunc(symshen_4fix_1help), V2134, V2135, gen8749)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fix"), gen8750)

		gen8753 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2146 := __args[0]
			_ = V2146
			V2147 := __args[1]
			_ = V2147
			V2148 := __args[2]
			_ = V2148
			gen8752 := Call(__e, ShenFunc(sym_a), V2148, V2147)

			if True == gen8752 {
				__e.Return(V2148)
				return
			} else {
				gen8751 := Call(__e, V2146, V2148)

				__e.TailApply(ShenFunc(symshen_4fix_1help), V2146, V2148, gen8751)

				return

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fix-help"), gen8753)

		gen8759 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2153 := __args[0]
			_ = V2153
			V2154 := __args[1]
			_ = V2154
			V2155 := __args[2]
			_ = V2155
			V2156 := __args[3]
			_ = V2156
			gen8754 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen8755 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4_5_1dict), V2156, V2153)

				return
			}, 0)
			gen8756 := Try(__e, gen8755).Catch(gen8754)
			Curr := gen8756
			gen8757 := Call(__e, ShenFunc(symshen_4assoc_1set), V2154, V2155, Curr)

			Added := gen8757
			gen8758 := Call(__e, ShenFunc(symshen_4dict_1_6), V2156, V2153, Added)

			Update := gen8758
			_ = Update
			__e.Return(V2155)
			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("put"), gen8759)

		gen8765 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2160 := __args[0]
			_ = V2160
			V2161 := __args[1]
			_ = V2161
			V2162 := __args[2]
			_ = V2162
			gen8760 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen8761 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4_5_1dict), V2162, V2160)

				return
			}, 0)
			gen8762 := Try(__e, gen8761).Catch(gen8760)
			Curr := gen8762
			gen8763 := Call(__e, ShenFunc(symshen_4assoc_1rm), V2161, Curr)

			Removed := gen8763
			gen8764 := Call(__e, ShenFunc(symshen_4dict_1_6), V2162, V2160, Removed)

			Update := gen8764
			_ = Update
			__e.Return(V2160)
			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unput"), gen8765)

		gen8771 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2166 := __args[0]
			_ = V2166
			V2167 := __args[1]
			_ = V2167
			V2168 := __args[2]
			_ = V2168
			gen8766 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(Nil)
				return
			}, 1)
			gen8767 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4_5_1dict), V2168, V2166)

				return
			}, 0)
			gen8768 := Try(__e, gen8767).Catch(gen8766)
			Entry := gen8768
			gen8769 := Call(__e, ShenFunc(symassoc), V2167, Entry)

			Result := gen8769
			gen8770 := Call(__e, ShenFunc(symempty_2), Result)

			if True == gen8770 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("value not found\n"))

				return
			} else {
				__e.TailApply(ShenFunc(symtl), Result)

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("get"), gen8771)

		gen8776 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2171 := __args[0]
			_ = V2171
			V2172 := __args[1]
			_ = V2172
			gen8772 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symstring_1_6n), X)

				return
			}, 1)
			gen8773 := Call(__e, ShenFunc(symexplode), V2171)

			gen8774 := Call(__e, ShenFunc(symmap), gen8772, gen8773)

			gen8775 := Call(__e, ShenFunc(symsum), gen8774)

			__e.TailApply(ShenFunc(symshen_4mod), gen8775, V2172)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("hash"), gen8776)

		gen8779 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2175 := __args[0]
			_ = V2175
			V2176 := __args[1]
			_ = V2176
			gen8777 := Call(__e, ShenFunc(symcons), V2176, Nil)

			gen8778 := Call(__e, ShenFunc(symshen_4multiples), V2175, gen8777)

			__e.TailApply(ShenFunc(symshen_4modh), V2175, gen8778)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mod"), gen8779)

		gen8788 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2179 := __args[0]
			_ = V2179
			V2180 := __args[1]
			_ = V2180
			gen8786 := Call(__e, ShenFunc(symcons_2), V2180)

			var gen8787 Obj
			if True == gen8786 {
				gen8784 := Call(__e, ShenFunc(symhd), V2180)

				gen8785 := Call(__e, ShenFunc(sym_6), gen8784, V2179)

				if True == gen8785 {
					gen8787 = True
				} else {
					gen8787 = False
				}

			} else {
				gen8787 = False
			}
			if True == gen8787 {
				__e.TailApply(ShenFunc(symtl), V2180)

				return
			} else {
				gen8783 := Call(__e, ShenFunc(symcons_2), V2180)

				if True == gen8783 {
					gen8780 := Call(__e, ShenFunc(symhd), V2180)

					gen8781 := Call(__e, ShenFunc(sym_d), MakeNumber(2), gen8780)

					gen8782 := Call(__e, ShenFunc(symcons), gen8781, V2180)

					__e.TailApply(ShenFunc(symshen_4multiples), V2179, gen8782)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.multiples"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.multiples"), gen8788)

		gen8801 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2185 := __args[0]
			_ = V2185
			V2186 := __args[1]
			_ = V2186
			gen8800 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V2185)

			if True == gen8800 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen8799 := Call(__e, ShenFunc(sym_a), Nil, V2186)

				if True == gen8799 {
					__e.Return(V2185)
					return
				} else {
					gen8797 := Call(__e, ShenFunc(symcons_2), V2186)

					var gen8798 Obj
					if True == gen8797 {
						gen8795 := Call(__e, ShenFunc(symhd), V2186)

						gen8796 := Call(__e, ShenFunc(sym_6), gen8795, V2185)

						if True == gen8796 {
							gen8798 = True
						} else {
							gen8798 = False
						}

					} else {
						gen8798 = False
					}
					if True == gen8798 {
						gen8793 := Call(__e, ShenFunc(symtl), V2186)

						gen8794 := Call(__e, ShenFunc(symempty_2), gen8793)

						if True == gen8794 {
							__e.Return(V2185)
							return
						} else {
							gen8792 := Call(__e, ShenFunc(symtl), V2186)

							__e.TailApply(ShenFunc(symshen_4modh), V2185, gen8792)

							return

						}

					} else {
						gen8791 := Call(__e, ShenFunc(symcons_2), V2186)

						if True == gen8791 {
							gen8789 := Call(__e, ShenFunc(symhd), V2186)

							gen8790 := Call(__e, ShenFunc(sym_1), V2185, gen8789)

							__e.TailApply(ShenFunc(symshen_4modh), gen8790, V2186)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.modh"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.modh"), gen8801)

		gen8807 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2188 := __args[0]
			_ = V2188
			gen8806 := Call(__e, ShenFunc(sym_a), Nil, V2188)

			if True == gen8806 {
				__e.Return(MakeNumber(0))
				return
			} else {
				gen8805 := Call(__e, ShenFunc(symcons_2), V2188)

				if True == gen8805 {
					gen8802 := Call(__e, ShenFunc(symhd), V2188)

					gen8803 := Call(__e, ShenFunc(symtl), V2188)

					gen8804 := Call(__e, ShenFunc(symsum), gen8803)

					__e.TailApply(ShenFunc(sym_7), gen8802, gen8804)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("sum"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("sum"), gen8807)

		gen8809 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2196 := __args[0]
			_ = V2196
			gen8808 := Call(__e, ShenFunc(symcons_2), V2196)

			if True == gen8808 {
				__e.TailApply(ShenFunc(symhd), V2196)

				return
			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("head expects a non-empty list"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("head"), gen8809)

		gen8811 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2204 := __args[0]
			_ = V2204
			gen8810 := Call(__e, ShenFunc(symcons_2), V2204)

			if True == gen8810 {
				__e.TailApply(ShenFunc(symtl), V2204)

				return
			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("tail expects a non-empty list"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("tail"), gen8811)

		gen8812 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2206 := __args[0]
			_ = V2206
			__e.TailApply(ShenFunc(sympos), V2206, MakeNumber(0))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("hdstr"), gen8812)

		gen8821 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2211 := __args[0]
			_ = V2211
			V2212 := __args[1]
			_ = V2212
			gen8820 := Call(__e, ShenFunc(sym_a), Nil, V2211)

			if True == gen8820 {
				__e.Return(Nil)
				return
			} else {
				gen8819 := Call(__e, ShenFunc(symcons_2), V2211)

				if True == gen8819 {
					gen8817 := Call(__e, ShenFunc(symhd), V2211)

					gen8818 := Call(__e, ShenFunc(symelement_2), gen8817, V2212)

					if True == gen8818 {
						gen8814 := Call(__e, ShenFunc(symhd), V2211)

						gen8815 := Call(__e, ShenFunc(symtl), V2211)

						gen8816 := Call(__e, ShenFunc(symintersection), gen8815, V2212)

						__e.TailApply(ShenFunc(symcons), gen8814, gen8816)

						return

					} else {
						gen8813 := Call(__e, ShenFunc(symtl), V2211)

						__e.TailApply(ShenFunc(symintersection), gen8813, V2212)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("intersection"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("intersection"), gen8821)

		gen8822 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2214 := __args[0]
			_ = V2214
			__e.TailApply(ShenFunc(symshen_4reverse__help), V2214, Nil)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("reverse"), gen8822)

		gen8828 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2217 := __args[0]
			_ = V2217
			V2218 := __args[1]
			_ = V2218
			gen8827 := Call(__e, ShenFunc(sym_a), Nil, V2217)

			if True == gen8827 {
				__e.Return(V2218)
				return
			} else {
				gen8826 := Call(__e, ShenFunc(symcons_2), V2217)

				if True == gen8826 {
					gen8823 := Call(__e, ShenFunc(symtl), V2217)

					gen8824 := Call(__e, ShenFunc(symhd), V2217)

					gen8825 := Call(__e, ShenFunc(symcons), gen8824, V2218)

					__e.TailApply(ShenFunc(symshen_4reverse__help), gen8823, gen8825)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.reverse_help"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.reverse_help"), gen8828)

		gen8837 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2221 := __args[0]
			_ = V2221
			V2222 := __args[1]
			_ = V2222
			gen8836 := Call(__e, ShenFunc(sym_a), Nil, V2221)

			if True == gen8836 {
				__e.Return(V2222)
				return
			} else {
				gen8835 := Call(__e, ShenFunc(symcons_2), V2221)

				if True == gen8835 {
					gen8833 := Call(__e, ShenFunc(symhd), V2221)

					gen8834 := Call(__e, ShenFunc(symelement_2), gen8833, V2222)

					if True == gen8834 {
						gen8832 := Call(__e, ShenFunc(symtl), V2221)

						__e.TailApply(ShenFunc(symunion), gen8832, V2222)

						return

					} else {
						gen8829 := Call(__e, ShenFunc(symhd), V2221)

						gen8830 := Call(__e, ShenFunc(symtl), V2221)

						gen8831 := Call(__e, ShenFunc(symunion), gen8830, V2222)

						__e.TailApply(ShenFunc(symcons), gen8829, gen8831)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("union"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("union"), gen8837)

		gen8849 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2224 := __args[0]
			_ = V2224
			gen8838 := Call(__e, ShenFunc(symshen_4proc_1nl), V2224)

			gen8839 := Call(__e, ShenFunc(symstoutput))

			gen8840 := Call(__e, ShenFunc(symshen_4prhush), gen8838, gen8839)

			Message := gen8840
			_ = Message
			gen8841 := Call(__e, ShenFunc(symstoutput))

			gen8842 := Call(__e, ShenFunc(symshen_4prhush), MakeString(" (y/n) "), gen8841)

			Y_1or_1N := gen8842
			_ = Y_1or_1N
			gen8843 := Call(__e, ShenFunc(symstinput))

			gen8844 := Call(__e, ShenFunc(symread), gen8843)

			gen8845 := Call(__e, ShenFunc(symshen_4app), gen8844, MakeString(""), MakeSymbol("shen.s"))

			Input := gen8845
			gen8848 := Call(__e, ShenFunc(sym_a), MakeString("y"), Input)

			if True == gen8848 {
				__e.Return(True)
				return
			} else {
				gen8847 := Call(__e, ShenFunc(sym_a), MakeString("n"), Input)

				if True == gen8847 {
					__e.Return(False)
					return
				} else {
					gen8846 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), MakeString("please answer y or n\n"), gen8846)

					__e.TailApply(ShenFunc(symy_1or_1n_2), V2224)

					return

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("y-or-n?"), gen8849)

		gen8850 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2226 := __args[0]
			_ = V2226
			if True == V2226 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("not"), gen8850)

		gen8854 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2239 := __args[0]
			_ = V2239
			V2240 := __args[1]
			_ = V2240
			V2241 := __args[2]
			_ = V2241
			gen8853 := Call(__e, ShenFunc(sym_a), V2241, V2240)

			if True == gen8853 {
				__e.Return(V2239)
				return
			} else {
				gen8852 := Call(__e, ShenFunc(symcons_2), V2241)

				if True == gen8852 {
					gen8851 := MakeNative(func(__e Evaluator, __args ...Obj) {
						W := __args[0]
						_ = W
						__e.TailApply(ShenFunc(symsubst), V2239, V2240, W)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen8851, V2241)

					return

				} else {
					__e.Return(V2241)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("subst"), gen8854)

		gen8856 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2243 := __args[0]
			_ = V2243
			gen8855 := Call(__e, ShenFunc(symshen_4app), V2243, MakeString(""), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symshen_4explode_1h), gen8855)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("explode"), gen8856)

		gen8862 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2245 := __args[0]
			_ = V2245
			gen8861 := Call(__e, ShenFunc(sym_a), MakeString(""), V2245)

			if True == gen8861 {
				__e.Return(Nil)
				return
			} else {
				gen8860 := Call(__e, ShenFunc(symshen_4_7string_2), V2245)

				if True == gen8860 {
					gen8857 := Call(__e, ShenFunc(sympos), V2245, MakeNumber(0))

					gen8858 := Call(__e, ShenFunc(symtlstr), V2245)

					gen8859 := Call(__e, ShenFunc(symshen_4explode_1h), gen8858)

					__e.TailApply(ShenFunc(symcons), gen8857, gen8859)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.explode-h"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.explode-h"), gen8862)

		gen8865 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2247 := __args[0]
			_ = V2247
			gen8863 := Call(__e, ShenFunc(sym_a), V2247, MakeString(""))

			var gen8864 Obj
			if True == gen8863 {
				gen8864 = MakeString("")
			} else {
				gen8864 = Call(__e, ShenFunc(symshen_4app), V2247, MakeString("/"), MakeSymbol("shen.a"))

			}
			__e.TailApply(ShenFunc(symset), MakeSymbol("*home-directory*"), gen8864)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("cd"), gen8865)

		gen8871 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2250 := __args[0]
			_ = V2250
			V2251 := __args[1]
			_ = V2251
			gen8870 := Call(__e, ShenFunc(sym_a), Nil, V2251)

			if True == gen8870 {
				__e.Return(True)
				return
			} else {
				gen8869 := Call(__e, ShenFunc(symcons_2), V2251)

				if True == gen8869 {
					gen8866 := Call(__e, ShenFunc(symhd), V2251)

					gen8867 := Call(__e, V2250, gen8866)

					_ = gen8867
					gen8868 := Call(__e, ShenFunc(symtl), V2251)

					__e.TailApply(ShenFunc(symshen_4for_1each), V2250, gen8868)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.for-each"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.for-each"), gen8871)

		gen8878 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2256 := __args[0]
			_ = V2256
			V2257 := __args[1]
			_ = V2257
			gen8877 := Call(__e, ShenFunc(sym_a), Nil, V2257)

			if True == gen8877 {
				__e.Return(Nil)
				return
			} else {
				gen8876 := Call(__e, ShenFunc(symcons_2), V2257)

				if True == gen8876 {
					gen8872 := Call(__e, ShenFunc(symhd), V2257)

					gen8873 := Call(__e, V2256, gen8872)

					gen8874 := Call(__e, ShenFunc(symtl), V2257)

					gen8875 := Call(__e, ShenFunc(symmap), V2256, gen8874)

					__e.TailApply(ShenFunc(symcons), gen8873, gen8875)

					return

				} else {
					__e.TailApply(V2256, V2257)

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("map"), gen8878)

		gen8879 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2259 := __args[0]
			_ = V2259
			__e.TailApply(ShenFunc(symshen_4length_1h), V2259, MakeNumber(0))

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("length"), gen8879)

		gen8883 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2262 := __args[0]
			_ = V2262
			V2263 := __args[1]
			_ = V2263
			gen8882 := Call(__e, ShenFunc(sym_a), Nil, V2262)

			if True == gen8882 {
				__e.Return(V2263)
				return
			} else {
				gen8880 := Call(__e, ShenFunc(symtl), V2262)

				gen8881 := Call(__e, ShenFunc(sym_7), V2263, MakeNumber(1))

				__e.TailApply(ShenFunc(symshen_4length_1h), gen8880, gen8881)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.length-h"), gen8883)

		gen8890 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2275 := __args[0]
			_ = V2275
			V2276 := __args[1]
			_ = V2276
			gen8889 := Call(__e, ShenFunc(sym_a), V2276, V2275)

			if True == gen8889 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen8888 := Call(__e, ShenFunc(symcons_2), V2276)

				if True == gen8888 {
					gen8884 := Call(__e, ShenFunc(symhd), V2276)

					gen8885 := Call(__e, ShenFunc(symoccurrences), V2275, gen8884)

					gen8886 := Call(__e, ShenFunc(symtl), V2276)

					gen8887 := Call(__e, ShenFunc(symoccurrences), V2275, gen8886)

					__e.TailApply(ShenFunc(sym_7), gen8885, gen8887)

					return

				} else {
					__e.Return(MakeNumber(0))
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("occurrences"), gen8890)

		gen8901 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2283 := __args[0]
			_ = V2283
			V2284 := __args[1]
			_ = V2284
			gen8899 := Call(__e, ShenFunc(sym_a), MakeNumber(1), V2283)

			var gen8900 Obj
			if True == gen8899 {
				gen8898 := Call(__e, ShenFunc(symcons_2), V2284)

				if True == gen8898 {
					gen8900 = True
				} else {
					gen8900 = False
				}

			} else {
				gen8900 = False
			}
			if True == gen8900 {
				__e.TailApply(ShenFunc(symhd), V2284)

				return
			} else {
				gen8897 := Call(__e, ShenFunc(symcons_2), V2284)

				if True == gen8897 {
					gen8895 := Call(__e, ShenFunc(sym_1), V2283, MakeNumber(1))

					gen8896 := Call(__e, ShenFunc(symtl), V2284)

					__e.TailApply(ShenFunc(symnth), gen8895, gen8896)

					return

				} else {
					gen8891 := Call(__e, ShenFunc(symshen_4app), V2284, MakeString("\n"), MakeSymbol("shen.a"))

					gen8892 := Call(__e, ShenFunc(symcn), MakeString(", "), gen8891)

					gen8893 := Call(__e, ShenFunc(symshen_4app), V2283, gen8892, MakeSymbol("shen.a"))

					gen8894 := Call(__e, ShenFunc(symcn), MakeString("nth applied to "), gen8893)

					__e.TailApply(ShenFunc(symsimple_1error), gen8894)

					return

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("nth"), gen8901)

		gen8906 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2286 := __args[0]
			_ = V2286
			gen8905 := Call(__e, ShenFunc(symnumber_2), V2286)

			if True == gen8905 {
				gen8902 := Call(__e, ShenFunc(symshen_4abs), V2286)

				Abs := gen8902
				gen8903 := Call(__e, ShenFunc(symshen_4magless), Abs, MakeNumber(1))

				gen8904 := Call(__e, ShenFunc(symshen_4integer_1test_2), Abs, gen8903)

				if True == gen8904 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("integer?"), gen8906)

		gen8908 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2288 := __args[0]
			_ = V2288
			gen8907 := Call(__e, ShenFunc(sym_6), V2288, MakeNumber(0))

			if True == gen8907 {
				__e.Return(V2288)
				return
			} else {
				__e.TailApply(ShenFunc(sym_1), MakeNumber(0), V2288)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abs"), gen8908)

		gen8911 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2291 := __args[0]
			_ = V2291
			V2292 := __args[1]
			_ = V2292
			gen8909 := Call(__e, ShenFunc(sym_d), V2292, MakeNumber(2))

			Nx2 := gen8909
			gen8910 := Call(__e, ShenFunc(sym_6), Nx2, V2291)

			if True == gen8910 {
				__e.Return(V2292)
				return
			} else {
				__e.TailApply(ShenFunc(symshen_4magless), V2291, Nx2)

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.magless"), gen8911)

		gen8916 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2298 := __args[0]
			_ = V2298
			V2299 := __args[1]
			_ = V2299
			gen8915 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V2298)

			if True == gen8915 {
				__e.Return(True)
				return
			} else {
				gen8914 := Call(__e, ShenFunc(sym_6), MakeNumber(1), V2298)

				if True == gen8914 {
					__e.Return(False)
					return
				} else {
					gen8912 := Call(__e, ShenFunc(sym_1), V2298, V2299)

					Abs_1N := gen8912
					gen8913 := Call(__e, ShenFunc(sym_6), MakeNumber(0), Abs_1N)

					if True == gen8913 {
						__e.TailApply(ShenFunc(syminteger_2), V2298)

						return
					} else {
						__e.TailApply(ShenFunc(symshen_4integer_1test_2), Abs_1N, V2299)

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.integer-test?"), gen8916)

		gen8923 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2304 := __args[0]
			_ = V2304
			V2305 := __args[1]
			_ = V2305
			gen8922 := Call(__e, ShenFunc(sym_a), Nil, V2305)

			if True == gen8922 {
				__e.Return(Nil)
				return
			} else {
				gen8921 := Call(__e, ShenFunc(symcons_2), V2305)

				if True == gen8921 {
					gen8917 := Call(__e, ShenFunc(symhd), V2305)

					gen8918 := Call(__e, V2304, gen8917)

					gen8919 := Call(__e, ShenFunc(symtl), V2305)

					gen8920 := Call(__e, ShenFunc(symmapcan), V2304, gen8919)

					__e.TailApply(ShenFunc(symappend), gen8918, gen8920)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("mapcan"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("mapcan"), gen8923)

		gen8925 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2317 := __args[0]
			_ = V2317
			V2318 := __args[1]
			_ = V2318
			gen8924 := Call(__e, ShenFunc(sym_a), V2318, V2317)

			if True == gen8924 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("=="), gen8925)

		gen8926 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symsimple_1error), MakeString(""))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("abort"), gen8926)

		gen8933 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2320 := __args[0]
			_ = V2320
			gen8932 := Call(__e, ShenFunc(symsymbol_2), V2320)

			if True == gen8932 {
				gen8927 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeSymbol("shen.this-symbol-is-unbound"))
					return
				}, 1)
				gen8928 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(symvalue), V2320)

					return
				}, 0)
				gen8929 := Try(__e, gen8928).Catch(gen8927)
				Val := gen8929
				gen8930 := Call(__e, ShenFunc(sym_a), Val, MakeSymbol("shen.this-symbol-is-unbound"))

				var gen8931 Obj
				if True == gen8930 {
					gen8931 = False
				} else {
					gen8931 = True
				}
				if True == gen8931 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("bound?"), gen8933)

		gen8939 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2322 := __args[0]
			_ = V2322
			gen8938 := Call(__e, ShenFunc(sym_a), MakeString(""), V2322)

			if True == gen8938 {
				__e.Return(Nil)
				return
			} else {
				gen8934 := Call(__e, ShenFunc(sympos), V2322, MakeNumber(0))

				gen8935 := Call(__e, ShenFunc(symstring_1_6n), gen8934)

				gen8936 := Call(__e, ShenFunc(symtlstr), V2322)

				gen8937 := Call(__e, ShenFunc(symshen_4string_1_6bytes), gen8936)

				__e.TailApply(ShenFunc(symcons), gen8935, gen8937)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.string->bytes"), gen8939)

		gen8940 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2324 := __args[0]
			_ = V2324
			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*maxinferences*"), V2324)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("maxinferences"), gen8940)

		gen8941 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*infs*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("inferences"), gen8941)

		gen8942 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2326 := __args[0]
			_ = V2326
			__e.Return(V2326)
			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("protect"), gen8942)

		gen8943 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*stoutput*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("stoutput"), gen8943)

		gen8944 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*sterror*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("sterror"), gen8944)

		gen8949 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2328 := __args[0]
			_ = V2328
			gen8945 := Call(__e, ShenFunc(symintern), V2328)

			Symbol := gen8945
			gen8948 := Call(__e, ShenFunc(symsymbol_2), Symbol)

			if True == gen8948 {
				__e.Return(Symbol)
				return
			} else {
				gen8946 := Call(__e, ShenFunc(symshen_4app), V2328, MakeString(" to a symbol"), MakeSymbol("shen.s"))

				gen8947 := Call(__e, ShenFunc(symcn), MakeString("cannot intern "), gen8946)

				__e.TailApply(ShenFunc(symsimple_1error), gen8947)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("string->symbol"), gen8949)

		gen8952 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2334 := __args[0]
			_ = V2334
			gen8951 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V2334)

			if True == gen8951 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*optimise*"), True)

				return
			} else {
				gen8950 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V2334)

				if True == gen8950 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*optimise*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("optimise expects a + or a -.\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("optimise"), gen8952)

		gen8953 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*os*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("os"), gen8953)

		gen8954 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*language*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("language"), gen8954)

		gen8955 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*version*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("version"), gen8955)

		gen8956 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*port*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("port"), gen8956)

		gen8957 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*porters*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("porters"), gen8957)

		gen8958 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*implementation*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("implementation"), gen8958)

		gen8959 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("*release*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("release"), gen8959)

		gen8962 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2336 := __args[0]
			_ = V2336
			gen8960 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(False)
				return
			}, 1)
			gen8961 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Call(__e, ShenFunc(symexternal), V2336)
				__e.Return(True)
				return

			}, 0)
			__e.Return(Try(__e, gen8961).Catch(gen8960))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("package?"), gen8962)

		gen8963 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2338 := __args[0]
			_ = V2338
			__e.TailApply(ShenFunc(symshen_4lookup_1func), V2338)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("function"), gen8963)

		gen8968 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2340 := __args[0]
			_ = V2340
			gen8965 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				gen8964 := Call(__e, ShenFunc(symshen_4app), V2340, MakeString(" has no lambda expansion\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen8964)

				return

			}, 1)
			gen8967 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen8966 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V2340, MakeSymbol("shen.lambda-form"), gen8966)

				return

			}, 0)
			__e.Return(Try(__e, gen8967).Catch(gen8965))
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.lookup-func"), gen8968)

		return

	}, 0))
}
