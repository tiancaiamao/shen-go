package main

import . "github.com/tiancaiamao/shen-go/cora"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
	tmp7171 := MakeNative(func(__e *ControlFlow) {
		V1786 := __e.Get(1)
		_ = V1786
		tmp7172 := MakeNative(func(__e *ControlFlow) {
			Fs := __e.Get(1)
			_ = Fs
			__e.TailApply(PrimNS2Value(symshen_4macroexpand_1h), V1786, Fs, Fs)
			return
		}, 1)

		tmp7173 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(PrimTail(X))
			return
		}, 1)

		tmp7174 := PrimNS3Value(sym_dmacros_d)

		tmp7175 := Call(__e, PrimNS2Value(symmap), tmp7173, tmp7174)

		__e.TailApply(tmp7172, tmp7175)
		return

	}, 1)

	tmp7176 := Call(__e, PrimNS2Value(symdef), symmacroexpand, tmp7171)

	_ = tmp7176

	tmp7177 := MakeNative(func(__e *ControlFlow) {
		V1795 := __e.Get(1)
		_ = V1795
		V1796 := __e.Get(2)
		_ = V1796
		V1797 := __e.Get(3)
		_ = V1797
		tmp7187 := PrimEqual(Nil, V1796)

		if True == tmp7187 {
			__e.Return(V1795)
			return
		} else {
			tmp7186 := PrimIsPair(V1796)

			if True == tmp7186 {
				tmp7180 := MakeNative(func(__e *ControlFlow) {
					Y := __e.Get(1)
					_ = Y
					tmp7183 := PrimEqual(V1795, Y)

					if True == tmp7183 {
						tmp7182 := PrimTail(V1796)

						__e.TailApply(PrimNS2Value(symshen_4macroexpand_1h), V1795, tmp7182, V1797)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4macroexpand_1h), Y, V1797, V1797)
						return
					}

				}, 1)

				tmp7184 := PrimHead(V1796)

				tmp7185 := Call(__e, PrimNS2Value(symshen_4walk), tmp7184, V1795)

				__e.TailApply(tmp7180, tmp7185)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
				return
			}

		}

	}, 3)

	tmp7188 := Call(__e, PrimNS2Value(symdef), symshen_4macroexpand_1h, tmp7177)

	_ = tmp7188

	tmp7189 := MakeNative(func(__e *ControlFlow) {
		V1798 := __e.Get(1)
		_ = V1798
		V1799 := __e.Get(2)
		_ = V1799
		tmp7193 := PrimIsPair(V1799)

		if True == tmp7193 {
			tmp7191 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimNS2Value(symshen_4walk), V1798, Z)
				return
			}, 1)

			tmp7192 := Call(__e, PrimNS2Value(symmap), tmp7191, V1799)

			__e.TailApply(V1798, tmp7192)
			return

		} else {
			__e.TailApply(V1798, V1799)
			return
		}

	}, 2)

	tmp7194 := Call(__e, PrimNS2Value(symdef), symshen_4walk, tmp7189)

	_ = tmp7194

	tmp7195 := MakeNative(func(__e *ControlFlow) {
		V1800 := __e.Get(1)
		_ = V1800
		tmp7228 := PrimIsPair(V1800)

		var ifres7220 Obj

		if True == tmp7228 {
			tmp7226 := PrimHead(V1800)

			tmp7227 := PrimEqual(symdefmacro, tmp7226)

			var ifres7222 Obj

			if True == tmp7227 {
				tmp7224 := PrimTail(V1800)

				tmp7225 := PrimIsPair(tmp7224)

				var ifres7223 Obj

				if True == tmp7225 {
					ifres7223 = True

				} else {
					ifres7223 = False

				}

				ifres7222 = ifres7223

			} else {
				ifres7222 = False

			}

			var ifres7221 Obj

			if True == ifres7222 {
				ifres7221 = True

			} else {
				ifres7221 = False

			}

			ifres7220 = ifres7221

		} else {
			ifres7220 = False

		}

		if True == ifres7220 {
			tmp7197 := MakeNative(func(__e *ControlFlow) {
				Default := __e.Get(1)
				_ = Default
				tmp7198 := MakeNative(func(__e *ControlFlow) {
					Def := __e.Get(1)
					_ = Def
					tmp7199 := MakeNative(func(__e *ControlFlow) {
						Record := __e.Get(1)
						_ = Record
						tmp7200 := PrimTail(V1800)

						__e.Return(PrimHead(tmp7200))
						return

					}, 1)

					tmp7201 := PrimTail(V1800)

					tmp7202 := PrimHead(tmp7201)

					tmp7203 := PrimTail(V1800)

					tmp7204 := PrimHead(tmp7203)

					tmp7205 := Call(__e, PrimNS2Value(symfn), tmp7204)

					tmp7206 := Call(__e, PrimNS2Value(symshen_4record_1macro), tmp7202, tmp7205)

					__e.TailApply(tmp7199, tmp7206)
					return

				}, 1)

				tmp7207 := PrimTail(V1800)

				tmp7208 := PrimHead(tmp7207)

				tmp7209 := PrimTail(V1800)

				tmp7210 := PrimTail(tmp7209)

				tmp7211 := Call(__e, PrimNS2Value(symappend), tmp7210, Default)

				tmp7212 := PrimCons(tmp7208, tmp7211)

				tmp7213 := PrimCons(symdefine, tmp7212)

				tmp7214 := Call(__e, PrimNS2Value(symeval), tmp7213)

				__e.TailApply(tmp7198, tmp7214)
				return

			}, 1)

			tmp7215 := Call(__e, PrimNS2Value(symprotect), symX)

			tmp7216 := Call(__e, PrimNS2Value(symprotect), symX)

			tmp7217 := PrimCons(tmp7216, Nil)

			tmp7218 := PrimCons(sym_1_6, tmp7217)

			tmp7219 := PrimCons(tmp7215, tmp7218)

			__e.TailApply(tmp7197, tmp7219)
			return

		} else {
			__e.Return(V1800)
			return
		}

	}, 1)

	tmp7229 := Call(__e, PrimNS2Value(symdef), symshen_4defmacro_1macro, tmp7195)

	_ = tmp7229

	tmp7230 := MakeNative(func(__e *ControlFlow) {
		V1801 := __e.Get(1)
		_ = V1801
		tmp7249 := PrimIsPair(V1801)

		var ifres7236 Obj

		if True == tmp7249 {
			tmp7247 := PrimHead(V1801)

			tmp7248 := PrimEqual(symu_b, tmp7247)

			var ifres7238 Obj

			if True == tmp7248 {
				tmp7245 := PrimTail(V1801)

				tmp7246 := PrimIsPair(tmp7245)

				var ifres7240 Obj

				if True == tmp7246 {
					tmp7242 := PrimTail(V1801)

					tmp7243 := PrimTail(tmp7242)

					tmp7244 := PrimEqual(Nil, tmp7243)

					var ifres7241 Obj

					if True == tmp7244 {
						ifres7241 = True

					} else {
						ifres7241 = False

					}

					ifres7240 = ifres7241

				} else {
					ifres7240 = False

				}

				var ifres7239 Obj

				if True == ifres7240 {
					ifres7239 = True

				} else {
					ifres7239 = False

				}

				ifres7238 = ifres7239

			} else {
				ifres7238 = False

			}

			var ifres7237 Obj

			if True == ifres7238 {
				ifres7237 = True

			} else {
				ifres7237 = False

			}

			ifres7236 = ifres7237

		} else {
			ifres7236 = False

		}

		if True == ifres7236 {
			tmp7232 := PrimTail(V1801)

			tmp7233 := PrimHead(tmp7232)

			tmp7234 := Call(__e, PrimNS2Value(symshen_4make_1uppercase), tmp7233)

			tmp7235 := PrimCons(tmp7234, Nil)

			__e.Return(PrimCons(symprotect, tmp7235))
			return

		} else {
			__e.Return(V1801)
			return
		}

	}, 1)

	tmp7250 := Call(__e, PrimNS2Value(symdef), symshen_4u_b_1macro, tmp7230)

	_ = tmp7250

	tmp7251 := MakeNative(func(__e *ControlFlow) {
		V1802 := __e.Get(1)
		_ = V1802
		tmp7252 := PrimStr(V1802)

		tmp7253 := Call(__e, PrimNS2Value(symshen_4mu_1h), tmp7252)

		__e.Return(PrimIntern(tmp7253))
		return

	}, 1)

	tmp7254 := Call(__e, PrimNS2Value(symdef), symshen_4make_1uppercase, tmp7251)

	_ = tmp7254

	tmp7255 := MakeNative(func(__e *ControlFlow) {
		V1803 := __e.Get(1)
		_ = V1803
		tmp7274 := PrimEqual(MakeString(""), V1803)

		if True == tmp7274 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp7273 := Call(__e, PrimNS2Value(symshen_4_7string_2), V1803)

			if True == tmp7273 {
				tmp7258 := MakeNative(func(__e *ControlFlow) {
					ASCII := __e.Get(1)
					_ = ASCII
					tmp7259 := MakeNative(func(__e *ControlFlow) {
						ASCII_132 := __e.Get(1)
						_ = ASCII_132
						tmp7260 := MakeNative(func(__e *ControlFlow) {
							Upper := __e.Get(1)
							_ = Upper
							tmp7261 := PrimTailString(V1803)

							tmp7262 := Call(__e, PrimNS2Value(symshen_4mu_1h), tmp7261)

							__e.TailApply(PrimNS2Value(sym_8s), Upper, tmp7262)
							return

						}, 1)

						tmp7269 := PrimGreatEqual(ASCII, MakeNumber(97))

						var ifres7266 Obj

						if True == tmp7269 {
							tmp7268 := PrimLessEqual(ASCII, MakeNumber(122))

							var ifres7267 Obj

							if True == tmp7268 {
								ifres7267 = True

							} else {
								ifres7267 = False

							}

							ifres7266 = ifres7267

						} else {
							ifres7266 = False

						}

						var ifres7263 Obj

						if True == ifres7266 {
							tmp7265 := PrimNumberToString(ASCII_132)

							ifres7263 = tmp7265

						} else {
							tmp7264 := Call(__e, PrimNS2Value(symhdstr), V1803)

							ifres7263 = tmp7264

						}

						__e.TailApply(tmp7260, ifres7263)
						return

					}, 1)

					tmp7270 := PrimNumberSubtract(ASCII, MakeNumber(32))

					__e.TailApply(tmp7259, tmp7270)
					return

				}, 1)

				tmp7271 := Call(__e, PrimNS2Value(symhdstr), V1803)

				tmp7272 := PrimStringToNumber(tmp7271)

				__e.TailApply(tmp7258, tmp7272)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4mu_1h)
				return
			}

		}

	}, 1)

	tmp7275 := Call(__e, PrimNS2Value(symdef), symshen_4mu_1h, tmp7255)

	_ = tmp7275

	tmp7276 := MakeNative(func(__e *ControlFlow) {
		V1804 := __e.Get(1)
		_ = V1804
		V1805 := __e.Get(2)
		_ = V1805
		tmp7277 := PrimNS3Value(sym_dmacros_d)

		tmp7278 := Call(__e, PrimNS2Value(symshen_4update_1assoc), V1804, V1805, tmp7277)

		__e.Return(PrimNS3Set(sym_dmacros_d, tmp7278))
		return

	}, 2)

	tmp7279 := Call(__e, PrimNS2Value(symdef), symshen_4record_1macro, tmp7276)

	_ = tmp7279

	tmp7280 := MakeNative(func(__e *ControlFlow) {
		V1815 := __e.Get(1)
		_ = V1815
		V1816 := __e.Get(2)
		_ = V1816
		V1817 := __e.Get(3)
		_ = V1817
		tmp7303 := PrimEqual(Nil, V1817)

		if True == tmp7303 {
			tmp7302 := PrimCons(V1815, V1816)

			__e.Return(PrimCons(tmp7302, Nil))
			return

		} else {
			tmp7301 := PrimIsPair(V1817)

			var ifres7292 Obj

			if True == tmp7301 {
				tmp7299 := PrimHead(V1817)

				tmp7300 := PrimIsPair(tmp7299)

				var ifres7294 Obj

				if True == tmp7300 {
					tmp7296 := PrimHead(V1817)

					tmp7297 := PrimHead(tmp7296)

					tmp7298 := PrimEqual(V1815, tmp7297)

					var ifres7295 Obj

					if True == tmp7298 {
						ifres7295 = True

					} else {
						ifres7295 = False

					}

					ifres7294 = ifres7295

				} else {
					ifres7294 = False

				}

				var ifres7293 Obj

				if True == ifres7294 {
					ifres7293 = True

				} else {
					ifres7293 = False

				}

				ifres7292 = ifres7293

			} else {
				ifres7292 = False

			}

			if True == ifres7292 {
				tmp7288 := PrimHead(V1817)

				tmp7289 := PrimHead(tmp7288)

				tmp7290 := PrimCons(tmp7289, V1816)

				tmp7291 := PrimTail(V1817)

				__e.Return(PrimCons(tmp7290, tmp7291))
				return

			} else {
				tmp7287 := PrimIsPair(V1817)

				if True == tmp7287 {
					tmp7284 := PrimHead(V1817)

					tmp7285 := PrimTail(V1817)

					tmp7286 := Call(__e, PrimNS2Value(symshen_4update_1assoc), V1815, V1816, tmp7285)

					__e.Return(PrimCons(tmp7284, tmp7286))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
					return
				}

			}

		}

	}, 3)

	tmp7304 := Call(__e, PrimNS2Value(symdef), symshen_4update_1assoc, tmp7280)

	_ = tmp7304

	tmp7305 := MakeNative(func(__e *ControlFlow) {
		V1818 := __e.Get(1)
		_ = V1818
		tmp7321 := PrimIsPair(V1818)

		var ifres7313 Obj

		if True == tmp7321 {
			tmp7319 := PrimHead(V1818)

			tmp7320 := PrimEqual(symerror, tmp7319)

			var ifres7315 Obj

			if True == tmp7320 {
				tmp7317 := PrimTail(V1818)

				tmp7318 := PrimIsPair(tmp7317)

				var ifres7316 Obj

				if True == tmp7318 {
					ifres7316 = True

				} else {
					ifres7316 = False

				}

				ifres7315 = ifres7316

			} else {
				ifres7315 = False

			}

			var ifres7314 Obj

			if True == ifres7315 {
				ifres7314 = True

			} else {
				ifres7314 = False

			}

			ifres7313 = ifres7314

		} else {
			ifres7313 = False

		}

		if True == ifres7313 {
			tmp7307 := PrimTail(V1818)

			tmp7308 := PrimHead(tmp7307)

			tmp7309 := PrimTail(V1818)

			tmp7310 := PrimTail(tmp7309)

			tmp7311 := Call(__e, PrimNS2Value(symshen_4mkstr), tmp7308, tmp7310)

			tmp7312 := PrimCons(tmp7311, Nil)

			__e.Return(PrimCons(symsimple_1error, tmp7312))
			return

		} else {
			__e.Return(V1818)
			return
		}

	}, 1)

	tmp7322 := Call(__e, PrimNS2Value(symdef), symshen_4error_1macro, tmp7305)

	_ = tmp7322

	tmp7323 := MakeNative(func(__e *ControlFlow) {
		V1819 := __e.Get(1)
		_ = V1819
		tmp7361 := PrimIsPair(V1819)

		var ifres7353 Obj

		if True == tmp7361 {
			tmp7359 := PrimHead(V1819)

			tmp7360 := PrimEqual(symoutput, tmp7359)

			var ifres7355 Obj

			if True == tmp7360 {
				tmp7357 := PrimTail(V1819)

				tmp7358 := PrimIsPair(tmp7357)

				var ifres7356 Obj

				if True == tmp7358 {
					ifres7356 = True

				} else {
					ifres7356 = False

				}

				ifres7355 = ifres7356

			} else {
				ifres7355 = False

			}

			var ifres7354 Obj

			if True == ifres7355 {
				ifres7354 = True

			} else {
				ifres7354 = False

			}

			ifres7353 = ifres7354

		} else {
			ifres7353 = False

		}

		if True == ifres7353 {
			tmp7345 := PrimTail(V1819)

			tmp7346 := PrimHead(tmp7345)

			tmp7347 := PrimTail(V1819)

			tmp7348 := PrimTail(tmp7347)

			tmp7349 := Call(__e, PrimNS2Value(symshen_4mkstr), tmp7346, tmp7348)

			tmp7350 := PrimCons(symstoutput, Nil)

			tmp7351 := PrimCons(tmp7350, Nil)

			tmp7352 := PrimCons(tmp7349, tmp7351)

			__e.Return(PrimCons(sympr, tmp7352))
			return

		} else {
			tmp7344 := PrimIsPair(V1819)

			var ifres7331 Obj

			if True == tmp7344 {
				tmp7342 := PrimHead(V1819)

				tmp7343 := PrimEqual(sympr, tmp7342)

				var ifres7333 Obj

				if True == tmp7343 {
					tmp7340 := PrimTail(V1819)

					tmp7341 := PrimIsPair(tmp7340)

					var ifres7335 Obj

					if True == tmp7341 {
						tmp7337 := PrimTail(V1819)

						tmp7338 := PrimTail(tmp7337)

						tmp7339 := PrimEqual(Nil, tmp7338)

						var ifres7336 Obj

						if True == tmp7339 {
							ifres7336 = True

						} else {
							ifres7336 = False

						}

						ifres7335 = ifres7336

					} else {
						ifres7335 = False

					}

					var ifres7334 Obj

					if True == ifres7335 {
						ifres7334 = True

					} else {
						ifres7334 = False

					}

					ifres7333 = ifres7334

				} else {
					ifres7333 = False

				}

				var ifres7332 Obj

				if True == ifres7333 {
					ifres7332 = True

				} else {
					ifres7332 = False

				}

				ifres7331 = ifres7332

			} else {
				ifres7331 = False

			}

			if True == ifres7331 {
				tmp7326 := PrimTail(V1819)

				tmp7327 := PrimHead(tmp7326)

				tmp7328 := PrimCons(symstoutput, Nil)

				tmp7329 := PrimCons(tmp7328, Nil)

				tmp7330 := PrimCons(tmp7327, tmp7329)

				__e.Return(PrimCons(sympr, tmp7330))
				return

			} else {
				__e.Return(V1819)
				return
			}

		}

	}, 1)

	tmp7362 := Call(__e, PrimNS2Value(symdef), symshen_4output_1macro, tmp7323)

	_ = tmp7362

	tmp7363 := MakeNative(func(__e *ControlFlow) {
		V1820 := __e.Get(1)
		_ = V1820
		tmp7377 := PrimIsPair(V1820)

		var ifres7369 Obj

		if True == tmp7377 {
			tmp7375 := PrimHead(V1820)

			tmp7376 := PrimEqual(symmake_1string, tmp7375)

			var ifres7371 Obj

			if True == tmp7376 {
				tmp7373 := PrimTail(V1820)

				tmp7374 := PrimIsPair(tmp7373)

				var ifres7372 Obj

				if True == tmp7374 {
					ifres7372 = True

				} else {
					ifres7372 = False

				}

				ifres7371 = ifres7372

			} else {
				ifres7371 = False

			}

			var ifres7370 Obj

			if True == ifres7371 {
				ifres7370 = True

			} else {
				ifres7370 = False

			}

			ifres7369 = ifres7370

		} else {
			ifres7369 = False

		}

		if True == ifres7369 {
			tmp7365 := PrimTail(V1820)

			tmp7366 := PrimHead(tmp7365)

			tmp7367 := PrimTail(V1820)

			tmp7368 := PrimTail(tmp7367)

			__e.TailApply(PrimNS2Value(symshen_4mkstr), tmp7366, tmp7368)
			return

		} else {
			__e.Return(V1820)
			return
		}

	}, 1)

	tmp7378 := Call(__e, PrimNS2Value(symdef), symshen_4make_1string_1macro, tmp7363)

	_ = tmp7378

	tmp7379 := MakeNative(func(__e *ControlFlow) {
		V1821 := __e.Get(1)
		_ = V1821
		tmp7454 := PrimIsPair(V1821)

		var ifres7446 Obj

		if True == tmp7454 {
			tmp7452 := PrimHead(V1821)

			tmp7453 := PrimEqual(symlineread, tmp7452)

			var ifres7448 Obj

			if True == tmp7453 {
				tmp7450 := PrimTail(V1821)

				tmp7451 := PrimEqual(Nil, tmp7450)

				var ifres7449 Obj

				if True == tmp7451 {
					ifres7449 = True

				} else {
					ifres7449 = False

				}

				ifres7448 = ifres7449

			} else {
				ifres7448 = False

			}

			var ifres7447 Obj

			if True == ifres7448 {
				ifres7447 = True

			} else {
				ifres7447 = False

			}

			ifres7446 = ifres7447

		} else {
			ifres7446 = False

		}

		if True == ifres7446 {
			tmp7444 := PrimCons(symstinput, Nil)

			tmp7445 := PrimCons(tmp7444, Nil)

			__e.Return(PrimCons(symlineread, tmp7445))
			return

		} else {
			tmp7443 := PrimIsPair(V1821)

			var ifres7435 Obj

			if True == tmp7443 {
				tmp7441 := PrimHead(V1821)

				tmp7442 := PrimEqual(syminput, tmp7441)

				var ifres7437 Obj

				if True == tmp7442 {
					tmp7439 := PrimTail(V1821)

					tmp7440 := PrimEqual(Nil, tmp7439)

					var ifres7438 Obj

					if True == tmp7440 {
						ifres7438 = True

					} else {
						ifres7438 = False

					}

					ifres7437 = ifres7438

				} else {
					ifres7437 = False

				}

				var ifres7436 Obj

				if True == ifres7437 {
					ifres7436 = True

				} else {
					ifres7436 = False

				}

				ifres7435 = ifres7436

			} else {
				ifres7435 = False

			}

			if True == ifres7435 {
				tmp7433 := PrimCons(symstinput, Nil)

				tmp7434 := PrimCons(tmp7433, Nil)

				__e.Return(PrimCons(syminput, tmp7434))
				return

			} else {
				tmp7432 := PrimIsPair(V1821)

				var ifres7424 Obj

				if True == tmp7432 {
					tmp7430 := PrimHead(V1821)

					tmp7431 := PrimEqual(symread, tmp7430)

					var ifres7426 Obj

					if True == tmp7431 {
						tmp7428 := PrimTail(V1821)

						tmp7429 := PrimEqual(Nil, tmp7428)

						var ifres7427 Obj

						if True == tmp7429 {
							ifres7427 = True

						} else {
							ifres7427 = False

						}

						ifres7426 = ifres7427

					} else {
						ifres7426 = False

					}

					var ifres7425 Obj

					if True == ifres7426 {
						ifres7425 = True

					} else {
						ifres7425 = False

					}

					ifres7424 = ifres7425

				} else {
					ifres7424 = False

				}

				if True == ifres7424 {
					tmp7422 := PrimCons(symstinput, Nil)

					tmp7423 := PrimCons(tmp7422, Nil)

					__e.Return(PrimCons(symread, tmp7423))
					return

				} else {
					tmp7421 := PrimIsPair(V1821)

					var ifres7408 Obj

					if True == tmp7421 {
						tmp7419 := PrimHead(V1821)

						tmp7420 := PrimEqual(syminput_7, tmp7419)

						var ifres7410 Obj

						if True == tmp7420 {
							tmp7417 := PrimTail(V1821)

							tmp7418 := PrimIsPair(tmp7417)

							var ifres7412 Obj

							if True == tmp7418 {
								tmp7414 := PrimTail(V1821)

								tmp7415 := PrimTail(tmp7414)

								tmp7416 := PrimEqual(Nil, tmp7415)

								var ifres7413 Obj

								if True == tmp7416 {
									ifres7413 = True

								} else {
									ifres7413 = False

								}

								ifres7412 = ifres7413

							} else {
								ifres7412 = False

							}

							var ifres7411 Obj

							if True == ifres7412 {
								ifres7411 = True

							} else {
								ifres7411 = False

							}

							ifres7410 = ifres7411

						} else {
							ifres7410 = False

						}

						var ifres7409 Obj

						if True == ifres7410 {
							ifres7409 = True

						} else {
							ifres7409 = False

						}

						ifres7408 = ifres7409

					} else {
						ifres7408 = False

					}

					if True == ifres7408 {
						tmp7403 := PrimTail(V1821)

						tmp7404 := PrimHead(tmp7403)

						tmp7405 := PrimCons(symstinput, Nil)

						tmp7406 := PrimCons(tmp7405, Nil)

						tmp7407 := PrimCons(tmp7404, tmp7406)

						__e.Return(PrimCons(syminput_7, tmp7407))
						return

					} else {
						tmp7402 := PrimIsPair(V1821)

						var ifres7394 Obj

						if True == tmp7402 {
							tmp7400 := PrimHead(V1821)

							tmp7401 := PrimEqual(symread_1byte, tmp7400)

							var ifres7396 Obj

							if True == tmp7401 {
								tmp7398 := PrimTail(V1821)

								tmp7399 := PrimEqual(Nil, tmp7398)

								var ifres7397 Obj

								if True == tmp7399 {
									ifres7397 = True

								} else {
									ifres7397 = False

								}

								ifres7396 = ifres7397

							} else {
								ifres7396 = False

							}

							var ifres7395 Obj

							if True == ifres7396 {
								ifres7395 = True

							} else {
								ifres7395 = False

							}

							ifres7394 = ifres7395

						} else {
							ifres7394 = False

						}

						if True == ifres7394 {
							tmp7392 := Call(__e, PrimNS2Value(symstinput))

							tmp7393 := Call(__e, PrimNS2Value(symshen_4char_1stinput_2), tmp7392)

							if True == tmp7393 {
								tmp7388 := PrimCons(symstinput, Nil)

								tmp7389 := PrimCons(tmp7388, Nil)

								tmp7390 := PrimCons(symshen_4read_1unit_1string, tmp7389)

								tmp7391 := PrimCons(tmp7390, Nil)

								__e.Return(PrimCons(symstring_1_6n, tmp7391))
								return

							} else {
								tmp7386 := PrimCons(symstinput, Nil)

								tmp7387 := PrimCons(tmp7386, Nil)

								__e.Return(PrimCons(symread_1byte, tmp7387))
								return

							}

						} else {
							__e.Return(V1821)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp7455 := Call(__e, PrimNS2Value(symdef), symshen_4input_1macro, tmp7379)

	_ = tmp7455

	tmp7456 := MakeNative(func(__e *ControlFlow) {
		V1822 := __e.Get(1)
		_ = V1822
		tmp7463 := PrimIsPair(V1822)

		var ifres7459 Obj

		if True == tmp7463 {
			tmp7461 := PrimHead(V1822)

			tmp7462 := PrimEqual(symdefcc, tmp7461)

			var ifres7460 Obj

			if True == tmp7462 {
				ifres7460 = True

			} else {
				ifres7460 = False

			}

			ifres7459 = ifres7460

		} else {
			ifres7459 = False

		}

		if True == ifres7459 {
			tmp7458 := PrimTail(V1822)

			__e.TailApply(PrimNS2Value(symshen_4yacc_1_6shen), tmp7458)
			return

		} else {
			__e.Return(V1822)
			return
		}

	}, 1)

	tmp7464 := Call(__e, PrimNS2Value(symdef), symshen_4defcc_1macro, tmp7456)

	_ = tmp7464

	tmp7465 := MakeNative(func(__e *ControlFlow) {
		V1823 := __e.Get(1)
		_ = V1823
		tmp7472 := PrimIsPair(V1823)

		var ifres7468 Obj

		if True == tmp7472 {
			tmp7470 := PrimHead(V1823)

			tmp7471 := PrimEqual(symprolog_2, tmp7470)

			var ifres7469 Obj

			if True == tmp7471 {
				ifres7469 = True

			} else {
				ifres7469 = False

			}

			ifres7468 = ifres7469

		} else {
			ifres7468 = False

		}

		if True == ifres7468 {
			tmp7467 := PrimTail(V1823)

			__e.TailApply(PrimNS2Value(symshen_4call_1prolog), tmp7467)
			return

		} else {
			__e.Return(V1823)
			return
		}

	}, 1)

	tmp7473 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1macro, tmp7465)

	_ = tmp7473

	tmp7474 := MakeNative(func(__e *ControlFlow) {
		V1824 := __e.Get(1)
		_ = V1824
		tmp7475 := MakeNative(func(__e *ControlFlow) {
			Bindings := __e.Get(1)
			_ = Bindings
			tmp7476 := MakeNative(func(__e *ControlFlow) {
				Lock := __e.Get(1)
				_ = Lock
				tmp7477 := MakeNative(func(__e *ControlFlow) {
					Key := __e.Get(1)
					_ = Key
					tmp7478 := MakeNative(func(__e *ControlFlow) {
						Continuation := __e.Get(1)
						_ = Continuation
						tmp7479 := MakeNative(func(__e *ControlFlow) {
							CLiterals := __e.Get(1)
							_ = CLiterals
							tmp7480 := MakeNative(func(__e *ControlFlow) {
								Received := __e.Get(1)
								_ = Received
								tmp7481 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp7482 := MakeNative(func(__e *ControlFlow) {
										L := __e.Get(1)
										_ = L
										tmp7483 := MakeNative(func(__e *ControlFlow) {
											K := __e.Get(1)
											_ = K
											tmp7484 := MakeNative(func(__e *ControlFlow) {
												C := __e.Get(1)
												_ = C
												tmp7485 := MakeNative(func(__e *ControlFlow) {
													Lambda := __e.Get(1)
													_ = Lambda
													tmp7486 := PrimCons(Continuation, Nil)

													tmp7487 := PrimCons(Key, tmp7486)

													tmp7488 := PrimCons(Lock, tmp7487)

													tmp7489 := PrimCons(Bindings, tmp7488)

													__e.Return(PrimCons(Lambda, tmp7489))
													return

												}, 1)

												tmp7490 := Call(__e, PrimNS2Value(symshen_4continue), Received, CLiterals, B, L, K, C)

												tmp7491 := PrimCons(tmp7490, Nil)

												tmp7492 := PrimCons(C, tmp7491)

												tmp7493 := PrimCons(symlambda, tmp7492)

												tmp7494 := PrimCons(tmp7493, Nil)

												tmp7495 := PrimCons(K, tmp7494)

												tmp7496 := PrimCons(symlambda, tmp7495)

												tmp7497 := PrimCons(tmp7496, Nil)

												tmp7498 := PrimCons(L, tmp7497)

												tmp7499 := PrimCons(symlambda, tmp7498)

												tmp7500 := PrimCons(tmp7499, Nil)

												tmp7501 := PrimCons(B, tmp7500)

												tmp7502 := PrimCons(symlambda, tmp7501)

												__e.TailApply(tmp7485, tmp7502)
												return

											}, 1)

											tmp7503 := Call(__e, PrimNS2Value(symprotect), symC)

											tmp7504 := Call(__e, PrimNS2Value(symgensym), tmp7503)

											__e.TailApply(tmp7484, tmp7504)
											return

										}, 1)

										tmp7505 := Call(__e, PrimNS2Value(symprotect), symK)

										tmp7506 := Call(__e, PrimNS2Value(symgensym), tmp7505)

										__e.TailApply(tmp7483, tmp7506)
										return

									}, 1)

									tmp7507 := Call(__e, PrimNS2Value(symprotect), symL)

									tmp7508 := Call(__e, PrimNS2Value(symgensym), tmp7507)

									__e.TailApply(tmp7482, tmp7508)
									return

								}, 1)

								tmp7509 := Call(__e, PrimNS2Value(symprotect), symV)

								tmp7510 := Call(__e, PrimNS2Value(symgensym), tmp7509)

								__e.TailApply(tmp7481, tmp7510)
								return

							}, 1)

							tmp7511 := Call(__e, PrimNS2Value(symshen_4received), V1824)

							__e.TailApply(tmp7480, tmp7511)
							return

						}, 1)

						tmp7512 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							__e.TailApply(PrimNS2Value(symshen_4_5body_6), X)
							return
						}, 1)

						tmp7513 := Call(__e, PrimNS2Value(symcompile), tmp7512, V1824)

						__e.TailApply(tmp7479, tmp7513)
						return

					}, 1)

					tmp7514 := PrimCons(True, Nil)

					tmp7515 := PrimCons(symfreeze, tmp7514)

					__e.TailApply(tmp7478, tmp7515)
					return

				}, 1)

				__e.TailApply(tmp7477, MakeNumber(0))
				return

			}, 1)

			tmp7516 := PrimCons(MakeNumber(0), Nil)

			tmp7517 := PrimCons(symvector, tmp7516)

			tmp7518 := PrimCons(tmp7517, Nil)

			tmp7519 := PrimCons(MakeNumber(0), tmp7518)

			tmp7520 := PrimCons(True, tmp7519)

			tmp7521 := PrimCons(sym_8v, tmp7520)

			__e.TailApply(tmp7476, tmp7521)
			return

		}, 1)

		tmp7522 := PrimCons(symshen_4reset_1prolog_1vector, Nil)

		__e.TailApply(tmp7475, tmp7522)
		return

	}, 1)

	tmp7523 := Call(__e, PrimNS2Value(symdef), symshen_4call_1prolog, tmp7474)

	_ = tmp7523

	tmp7524 := MakeNative(func(__e *ControlFlow) {
		V1827 := __e.Get(1)
		_ = V1827
		tmp7545 := PrimIsPair(V1827)

		var ifres7532 Obj

		if True == tmp7545 {
			tmp7543 := PrimHead(V1827)

			tmp7544 := PrimEqual(symreceive, tmp7543)

			var ifres7534 Obj

			if True == tmp7544 {
				tmp7541 := PrimTail(V1827)

				tmp7542 := PrimIsPair(tmp7541)

				var ifres7536 Obj

				if True == tmp7542 {
					tmp7538 := PrimTail(V1827)

					tmp7539 := PrimTail(tmp7538)

					tmp7540 := PrimEqual(Nil, tmp7539)

					var ifres7537 Obj

					if True == tmp7540 {
						ifres7537 = True

					} else {
						ifres7537 = False

					}

					ifres7536 = ifres7537

				} else {
					ifres7536 = False

				}

				var ifres7535 Obj

				if True == ifres7536 {
					ifres7535 = True

				} else {
					ifres7535 = False

				}

				ifres7534 = ifres7535

			} else {
				ifres7534 = False

			}

			var ifres7533 Obj

			if True == ifres7534 {
				ifres7533 = True

			} else {
				ifres7533 = False

			}

			ifres7532 = ifres7533

		} else {
			ifres7532 = False

		}

		if True == ifres7532 {
			__e.Return(PrimTail(V1827))
			return
		} else {
			tmp7531 := PrimIsPair(V1827)

			if True == tmp7531 {
				tmp7527 := PrimHead(V1827)

				tmp7528 := Call(__e, PrimNS2Value(symshen_4received), tmp7527)

				tmp7529 := PrimTail(V1827)

				tmp7530 := Call(__e, PrimNS2Value(symshen_4received), tmp7529)

				__e.TailApply(PrimNS2Value(symunion), tmp7528, tmp7530)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp7546 := Call(__e, PrimNS2Value(symdef), symshen_4received, tmp7524)

	_ = tmp7546

	tmp7547 := MakeNative(func(__e *ControlFlow) {
		tmp7548 := PrimNS3Value(symshen_4_dprolog_1vector_d)

		__e.Return(PrimVectorSet(tmp7548, MakeNumber(1), MakeNumber(2)))
		return

	}, 0)

	tmp7549 := Call(__e, PrimNS2Value(symdef), symshen_4reset_1prolog_1vector, tmp7547)

	_ = tmp7549

	tmp7550 := MakeNative(func(__e *ControlFlow) {
		V1828 := __e.Get(1)
		_ = V1828
		__e.Return(V1828)
		return
	}, 1)

	tmp7551 := Call(__e, PrimNS2Value(symdef), symreceive, tmp7550)

	_ = tmp7551

	tmp7552 := MakeNative(func(__e *ControlFlow) {
		V1829 := __e.Get(1)
		_ = V1829
		tmp7566 := PrimIsPair(V1829)

		var ifres7558 Obj

		if True == tmp7566 {
			tmp7564 := PrimHead(V1829)

			tmp7565 := PrimEqual(symdefprolog, tmp7564)

			var ifres7560 Obj

			if True == tmp7565 {
				tmp7562 := PrimTail(V1829)

				tmp7563 := PrimIsPair(tmp7562)

				var ifres7561 Obj

				if True == tmp7563 {
					ifres7561 = True

				} else {
					ifres7561 = False

				}

				ifres7560 = ifres7561

			} else {
				ifres7560 = False

			}

			var ifres7559 Obj

			if True == ifres7560 {
				ifres7559 = True

			} else {
				ifres7559 = False

			}

			ifres7558 = ifres7559

		} else {
			ifres7558 = False

		}

		if True == ifres7558 {
			tmp7554 := PrimTail(V1829)

			tmp7555 := PrimHead(tmp7554)

			tmp7556 := PrimTail(V1829)

			tmp7557 := PrimTail(tmp7556)

			__e.TailApply(PrimNS2Value(symshen_4compile_1prolog), tmp7555, tmp7557)
			return

		} else {
			__e.Return(V1829)
			return
		}

	}, 1)

	tmp7567 := Call(__e, PrimNS2Value(symdef), symshen_4defprolog_1macro, tmp7552)

	_ = tmp7567

	tmp7568 := MakeNative(func(__e *ControlFlow) {
		V1830 := __e.Get(1)
		_ = V1830
		tmp7588 := PrimIsPair(V1830)

		var ifres7580 Obj

		if True == tmp7588 {
			tmp7586 := PrimHead(V1830)

			tmp7587 := PrimEqual(symdatatype, tmp7586)

			var ifres7582 Obj

			if True == tmp7587 {
				tmp7584 := PrimTail(V1830)

				tmp7585 := PrimIsPair(tmp7584)

				var ifres7583 Obj

				if True == tmp7585 {
					ifres7583 = True

				} else {
					ifres7583 = False

				}

				ifres7582 = ifres7583

			} else {
				ifres7582 = False

			}

			var ifres7581 Obj

			if True == ifres7582 {
				ifres7581 = True

			} else {
				ifres7581 = False

			}

			ifres7580 = ifres7581

		} else {
			ifres7580 = False

		}

		if True == ifres7580 {
			tmp7570 := MakeNative(func(__e *ControlFlow) {
				D := __e.Get(1)
				_ = D
				tmp7571 := MakeNative(func(__e *ControlFlow) {
					Compile := __e.Get(1)
					_ = Compile
					__e.Return(D)
					return
				}, 1)

				tmp7572 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4_5datatype_6), X)
					return
				}, 1)

				tmp7573 := PrimTail(V1830)

				tmp7574 := PrimTail(tmp7573)

				tmp7575 := PrimCons(D, tmp7574)

				tmp7576 := Call(__e, PrimNS2Value(symcompile), tmp7572, tmp7575)

				__e.TailApply(tmp7571, tmp7576)
				return

			}, 1)

			tmp7577 := PrimTail(V1830)

			tmp7578 := PrimHead(tmp7577)

			tmp7579 := Call(__e, PrimNS2Value(symshen_4intern_1type), tmp7578)

			__e.TailApply(tmp7570, tmp7579)
			return

		} else {
			__e.Return(V1830)
			return
		}

	}, 1)

	tmp7589 := Call(__e, PrimNS2Value(symdef), symshen_4datatype_1macro, tmp7568)

	_ = tmp7589

	tmp7590 := MakeNative(func(__e *ControlFlow) {
		V1831 := __e.Get(1)
		_ = V1831
		tmp7598 := PrimIsPair(V1831)

		if True == tmp7598 {
			tmp7592 := PrimHead(V1831)

			tmp7593 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp7592)

			tmp7594 := PrimTail(V1831)

			tmp7595 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp7594)

			tmp7596 := PrimCons(tmp7595, Nil)

			tmp7597 := PrimCons(tmp7593, tmp7596)

			__e.Return(PrimCons(symcons, tmp7597))
			return

		} else {
			__e.Return(V1831)
			return
		}

	}, 1)

	tmp7599 := Call(__e, PrimNS2Value(symdef), symshen_4rcons__form, tmp7590)

	_ = tmp7599

	tmp7600 := MakeNative(func(__e *ControlFlow) {
		V1832 := __e.Get(1)
		_ = V1832
		tmp7601 := PrimStr(V1832)

		tmp7602 := PrimStringConcat(tmp7601, MakeString("#type"))

		__e.Return(PrimIntern(tmp7602))
		return

	}, 1)

	tmp7603 := Call(__e, PrimNS2Value(symdef), symshen_4intern_1type, tmp7600)

	_ = tmp7603

	tmp7604 := MakeNative(func(__e *ControlFlow) {
		V1833 := __e.Get(1)
		_ = V1833
		tmp7670 := PrimIsPair(V1833)

		var ifres7651 Obj

		if True == tmp7670 {
			tmp7668 := PrimHead(V1833)

			tmp7669 := PrimEqual(sym_8s, tmp7668)

			var ifres7653 Obj

			if True == tmp7669 {
				tmp7666 := PrimTail(V1833)

				tmp7667 := PrimIsPair(tmp7666)

				var ifres7655 Obj

				if True == tmp7667 {
					tmp7663 := PrimTail(V1833)

					tmp7664 := PrimTail(tmp7663)

					tmp7665 := PrimIsPair(tmp7664)

					var ifres7657 Obj

					if True == tmp7665 {
						tmp7659 := PrimTail(V1833)

						tmp7660 := PrimTail(tmp7659)

						tmp7661 := PrimTail(tmp7660)

						tmp7662 := PrimIsPair(tmp7661)

						var ifres7658 Obj

						if True == tmp7662 {
							ifres7658 = True

						} else {
							ifres7658 = False

						}

						ifres7657 = ifres7658

					} else {
						ifres7657 = False

					}

					var ifres7656 Obj

					if True == ifres7657 {
						ifres7656 = True

					} else {
						ifres7656 = False

					}

					ifres7655 = ifres7656

				} else {
					ifres7655 = False

				}

				var ifres7654 Obj

				if True == ifres7655 {
					ifres7654 = True

				} else {
					ifres7654 = False

				}

				ifres7653 = ifres7654

			} else {
				ifres7653 = False

			}

			var ifres7652 Obj

			if True == ifres7653 {
				ifres7652 = True

			} else {
				ifres7652 = False

			}

			ifres7651 = ifres7652

		} else {
			ifres7651 = False

		}

		if True == ifres7651 {
			tmp7643 := PrimTail(V1833)

			tmp7644 := PrimHead(tmp7643)

			tmp7645 := PrimTail(V1833)

			tmp7646 := PrimTail(tmp7645)

			tmp7647 := PrimCons(sym_8s, tmp7646)

			tmp7648 := Call(__e, PrimNS2Value(symshen_4_8s_1macro), tmp7647)

			tmp7649 := PrimCons(tmp7648, Nil)

			tmp7650 := PrimCons(tmp7644, tmp7649)

			__e.Return(PrimCons(sym_8s, tmp7650))
			return

		} else {
			tmp7642 := PrimIsPair(V1833)

			var ifres7618 Obj

			if True == tmp7642 {
				tmp7640 := PrimHead(V1833)

				tmp7641 := PrimEqual(sym_8s, tmp7640)

				var ifres7620 Obj

				if True == tmp7641 {
					tmp7638 := PrimTail(V1833)

					tmp7639 := PrimIsPair(tmp7638)

					var ifres7622 Obj

					if True == tmp7639 {
						tmp7635 := PrimTail(V1833)

						tmp7636 := PrimTail(tmp7635)

						tmp7637 := PrimIsPair(tmp7636)

						var ifres7624 Obj

						if True == tmp7637 {
							tmp7631 := PrimTail(V1833)

							tmp7632 := PrimTail(tmp7631)

							tmp7633 := PrimTail(tmp7632)

							tmp7634 := PrimEqual(Nil, tmp7633)

							var ifres7626 Obj

							if True == tmp7634 {
								tmp7628 := PrimTail(V1833)

								tmp7629 := PrimHead(tmp7628)

								tmp7630 := PrimIsString(tmp7629)

								var ifres7627 Obj

								if True == tmp7630 {
									ifres7627 = True

								} else {
									ifres7627 = False

								}

								ifres7626 = ifres7627

							} else {
								ifres7626 = False

							}

							var ifres7625 Obj

							if True == ifres7626 {
								ifres7625 = True

							} else {
								ifres7625 = False

							}

							ifres7624 = ifres7625

						} else {
							ifres7624 = False

						}

						var ifres7623 Obj

						if True == ifres7624 {
							ifres7623 = True

						} else {
							ifres7623 = False

						}

						ifres7622 = ifres7623

					} else {
						ifres7622 = False

					}

					var ifres7621 Obj

					if True == ifres7622 {
						ifres7621 = True

					} else {
						ifres7621 = False

					}

					ifres7620 = ifres7621

				} else {
					ifres7620 = False

				}

				var ifres7619 Obj

				if True == ifres7620 {
					ifres7619 = True

				} else {
					ifres7619 = False

				}

				ifres7618 = ifres7619

			} else {
				ifres7618 = False

			}

			if True == ifres7618 {
				tmp7607 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					tmp7613 := Call(__e, PrimNS2Value(symlength), E)

					tmp7614 := PrimGreatThan(tmp7613, MakeNumber(1))

					if True == tmp7614 {
						tmp7609 := PrimTail(V1833)

						tmp7610 := PrimTail(tmp7609)

						tmp7611 := Call(__e, PrimNS2Value(symappend), E, tmp7610)

						tmp7612 := PrimCons(sym_8s, tmp7611)

						__e.TailApply(PrimNS2Value(symshen_4_8s_1macro), tmp7612)
						return

					} else {
						__e.Return(V1833)
						return
					}

				}, 1)

				tmp7615 := PrimTail(V1833)

				tmp7616 := PrimHead(tmp7615)

				tmp7617 := Call(__e, PrimNS2Value(symexplode), tmp7616)

				__e.TailApply(tmp7607, tmp7617)
				return

			} else {
				__e.Return(V1833)
				return
			}

		}

	}, 1)

	tmp7671 := Call(__e, PrimNS2Value(symdef), symshen_4_8s_1macro, tmp7604)

	_ = tmp7671

	tmp7672 := MakeNative(func(__e *ControlFlow) {
		V1834 := __e.Get(1)
		_ = V1834
		tmp7681 := PrimIsPair(V1834)

		var ifres7677 Obj

		if True == tmp7681 {
			tmp7679 := PrimHead(V1834)

			tmp7680 := PrimEqual(symsynonyms, tmp7679)

			var ifres7678 Obj

			if True == tmp7680 {
				ifres7678 = True

			} else {
				ifres7678 = False

			}

			ifres7677 = ifres7678

		} else {
			ifres7677 = False

		}

		if True == ifres7677 {
			tmp7674 := PrimTail(V1834)

			tmp7675 := PrimNS3Value(symshen_4_dsynonyms_d)

			tmp7676 := Call(__e, PrimNS2Value(symappend), tmp7674, tmp7675)

			__e.TailApply(PrimNS2Value(symshen_4synonyms_1h), tmp7676)
			return

		} else {
			__e.Return(V1834)
			return
		}

	}, 1)

	tmp7682 := Call(__e, PrimNS2Value(symdef), symshen_4synonyms_1macro, tmp7672)

	_ = tmp7682

	tmp7683 := MakeNative(func(__e *ControlFlow) {
		V1835 := __e.Get(1)
		_ = V1835
		tmp7684 := MakeNative(func(__e *ControlFlow) {
			CurryTypes := __e.Get(1)
			_ = CurryTypes
			tmp7685 := MakeNative(func(__e *ControlFlow) {
				Eval := __e.Get(1)
				_ = Eval
				__e.Return(symsynonyms)
				return
			}, 1)

			tmp7686 := Call(__e, PrimNS2Value(symshen_4compile_1synonyms), CurryTypes)

			tmp7687 := PrimCons(symshen_4demod, tmp7686)

			tmp7688 := PrimCons(symdefine, tmp7687)

			tmp7689 := Call(__e, PrimNS2Value(symeval), tmp7688)

			__e.TailApply(tmp7685, tmp7689)
			return

		}, 1)

		tmp7690 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4curry_1type), X)
			return
		}, 1)

		tmp7691 := Call(__e, PrimNS2Value(symmap), tmp7690, V1835)

		__e.TailApply(tmp7684, tmp7691)
		return

	}, 1)

	tmp7692 := Call(__e, PrimNS2Value(symdef), symshen_4synonyms_1h, tmp7683)

	_ = tmp7692

	tmp7693 := MakeNative(func(__e *ControlFlow) {
		V1838 := __e.Get(1)
		_ = V1838
		tmp7716 := PrimEqual(Nil, V1838)

		if True == tmp7716 {
			tmp7711 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp7712 := PrimCons(X, Nil)

				tmp7713 := PrimCons(sym_1_6, tmp7712)

				__e.Return(PrimCons(X, tmp7713))
				return

			}, 1)

			tmp7714 := Call(__e, PrimNS2Value(symprotect), symX)

			tmp7715 := Call(__e, PrimNS2Value(symgensym), tmp7714)

			__e.TailApply(tmp7711, tmp7715)
			return

		} else {
			tmp7710 := PrimIsPair(V1838)

			var ifres7706 Obj

			if True == tmp7710 {
				tmp7708 := PrimTail(V1838)

				tmp7709 := PrimIsPair(tmp7708)

				var ifres7707 Obj

				if True == tmp7709 {
					ifres7707 = True

				} else {
					ifres7707 = False

				}

				ifres7706 = ifres7707

			} else {
				ifres7706 = False

			}

			if True == ifres7706 {
				tmp7696 := PrimHead(V1838)

				tmp7697 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp7696)

				tmp7698 := PrimTail(V1838)

				tmp7699 := PrimHead(tmp7698)

				tmp7700 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp7699)

				tmp7701 := PrimTail(V1838)

				tmp7702 := PrimTail(tmp7701)

				tmp7703 := Call(__e, PrimNS2Value(symshen_4compile_1synonyms), tmp7702)

				tmp7704 := PrimCons(tmp7700, tmp7703)

				tmp7705 := PrimCons(sym_1_6, tmp7704)

				__e.Return(PrimCons(tmp7697, tmp7705))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
				return
			}

		}

	}, 1)

	tmp7717 := Call(__e, PrimNS2Value(symdef), symshen_4compile_1synonyms, tmp7693)

	_ = tmp7717

	tmp7718 := MakeNative(func(__e *ControlFlow) {
		V1839 := __e.Get(1)
		_ = V1839
		tmp7729 := PrimIsPair(V1839)

		var ifres7721 Obj

		if True == tmp7729 {
			tmp7727 := PrimHead(V1839)

			tmp7728 := PrimEqual(symnl, tmp7727)

			var ifres7723 Obj

			if True == tmp7728 {
				tmp7725 := PrimTail(V1839)

				tmp7726 := PrimEqual(Nil, tmp7725)

				var ifres7724 Obj

				if True == tmp7726 {
					ifres7724 = True

				} else {
					ifres7724 = False

				}

				ifres7723 = ifres7724

			} else {
				ifres7723 = False

			}

			var ifres7722 Obj

			if True == ifres7723 {
				ifres7722 = True

			} else {
				ifres7722 = False

			}

			ifres7721 = ifres7722

		} else {
			ifres7721 = False

		}

		if True == ifres7721 {
			tmp7720 := PrimCons(MakeNumber(1), Nil)

			__e.Return(PrimCons(symnl, tmp7720))
			return

		} else {
			__e.Return(V1839)
			return
		}

	}, 1)

	tmp7730 := Call(__e, PrimNS2Value(symdef), symshen_4nl_1macro, tmp7718)

	_ = tmp7730

	tmp7731 := MakeNative(func(__e *ControlFlow) {
		V1840 := __e.Get(1)
		_ = V1840
		tmp7770 := PrimIsPair(V1840)

		var ifres7743 Obj

		if True == tmp7770 {
			tmp7768 := PrimTail(V1840)

			tmp7769 := PrimIsPair(tmp7768)

			var ifres7745 Obj

			if True == tmp7769 {
				tmp7765 := PrimTail(V1840)

				tmp7766 := PrimTail(tmp7765)

				tmp7767 := PrimIsPair(tmp7766)

				var ifres7747 Obj

				if True == tmp7767 {
					tmp7761 := PrimTail(V1840)

					tmp7762 := PrimTail(tmp7761)

					tmp7763 := PrimTail(tmp7762)

					tmp7764 := PrimIsPair(tmp7763)

					var ifres7749 Obj

					if True == tmp7764 {
						tmp7751 := PrimHead(V1840)

						tmp7752 := PrimCons(symdo, Nil)

						tmp7753 := PrimCons(sym_d, tmp7752)

						tmp7754 := PrimCons(sym_7, tmp7753)

						tmp7755 := PrimCons(symor, tmp7754)

						tmp7756 := PrimCons(symand, tmp7755)

						tmp7757 := PrimCons(symappend, tmp7756)

						tmp7758 := PrimCons(sym_8v, tmp7757)

						tmp7759 := PrimCons(sym_8p, tmp7758)

						tmp7760 := Call(__e, PrimNS2Value(symelement_2), tmp7751, tmp7759)

						var ifres7750 Obj

						if True == tmp7760 {
							ifres7750 = True

						} else {
							ifres7750 = False

						}

						ifres7749 = ifres7750

					} else {
						ifres7749 = False

					}

					var ifres7748 Obj

					if True == ifres7749 {
						ifres7748 = True

					} else {
						ifres7748 = False

					}

					ifres7747 = ifres7748

				} else {
					ifres7747 = False

				}

				var ifres7746 Obj

				if True == ifres7747 {
					ifres7746 = True

				} else {
					ifres7746 = False

				}

				ifres7745 = ifres7746

			} else {
				ifres7745 = False

			}

			var ifres7744 Obj

			if True == ifres7745 {
				ifres7744 = True

			} else {
				ifres7744 = False

			}

			ifres7743 = ifres7744

		} else {
			ifres7743 = False

		}

		if True == ifres7743 {
			tmp7733 := PrimHead(V1840)

			tmp7734 := PrimTail(V1840)

			tmp7735 := PrimHead(tmp7734)

			tmp7736 := PrimHead(V1840)

			tmp7737 := PrimTail(V1840)

			tmp7738 := PrimTail(tmp7737)

			tmp7739 := PrimCons(tmp7736, tmp7738)

			tmp7740 := Call(__e, PrimNS2Value(symshen_4assoc_1macro), tmp7739)

			tmp7741 := PrimCons(tmp7740, Nil)

			tmp7742 := PrimCons(tmp7735, tmp7741)

			__e.Return(PrimCons(tmp7733, tmp7742))
			return

		} else {
			__e.Return(V1840)
			return
		}

	}, 1)

	tmp7771 := Call(__e, PrimNS2Value(symdef), symshen_4assoc_1macro, tmp7731)

	_ = tmp7771

	tmp7772 := MakeNative(func(__e *ControlFlow) {
		V1841 := __e.Get(1)
		_ = V1841
		tmp7848 := PrimIsPair(V1841)

		var ifres7822 Obj

		if True == tmp7848 {
			tmp7846 := PrimHead(V1841)

			tmp7847 := PrimEqual(symlet, tmp7846)

			var ifres7824 Obj

			if True == tmp7847 {
				tmp7844 := PrimTail(V1841)

				tmp7845 := PrimIsPair(tmp7844)

				var ifres7826 Obj

				if True == tmp7845 {
					tmp7841 := PrimTail(V1841)

					tmp7842 := PrimTail(tmp7841)

					tmp7843 := PrimIsPair(tmp7842)

					var ifres7828 Obj

					if True == tmp7843 {
						tmp7837 := PrimTail(V1841)

						tmp7838 := PrimTail(tmp7837)

						tmp7839 := PrimTail(tmp7838)

						tmp7840 := PrimIsPair(tmp7839)

						var ifres7830 Obj

						if True == tmp7840 {
							tmp7832 := PrimTail(V1841)

							tmp7833 := PrimTail(tmp7832)

							tmp7834 := PrimTail(tmp7833)

							tmp7835 := PrimTail(tmp7834)

							tmp7836 := PrimIsPair(tmp7835)

							var ifres7831 Obj

							if True == tmp7836 {
								ifres7831 = True

							} else {
								ifres7831 = False

							}

							ifres7830 = ifres7831

						} else {
							ifres7830 = False

						}

						var ifres7829 Obj

						if True == ifres7830 {
							ifres7829 = True

						} else {
							ifres7829 = False

						}

						ifres7828 = ifres7829

					} else {
						ifres7828 = False

					}

					var ifres7827 Obj

					if True == ifres7828 {
						ifres7827 = True

					} else {
						ifres7827 = False

					}

					ifres7826 = ifres7827

				} else {
					ifres7826 = False

				}

				var ifres7825 Obj

				if True == ifres7826 {
					ifres7825 = True

				} else {
					ifres7825 = False

				}

				ifres7824 = ifres7825

			} else {
				ifres7824 = False

			}

			var ifres7823 Obj

			if True == ifres7824 {
				ifres7823 = True

			} else {
				ifres7823 = False

			}

			ifres7822 = ifres7823

		} else {
			ifres7822 = False

		}

		if True == ifres7822 {
			tmp7809 := PrimTail(V1841)

			tmp7810 := PrimHead(tmp7809)

			tmp7811 := PrimTail(V1841)

			tmp7812 := PrimTail(tmp7811)

			tmp7813 := PrimHead(tmp7812)

			tmp7814 := PrimTail(V1841)

			tmp7815 := PrimTail(tmp7814)

			tmp7816 := PrimTail(tmp7815)

			tmp7817 := PrimCons(symlet, tmp7816)

			tmp7818 := Call(__e, PrimNS2Value(symshen_4let_1macro), tmp7817)

			tmp7819 := PrimCons(tmp7818, Nil)

			tmp7820 := PrimCons(tmp7813, tmp7819)

			tmp7821 := PrimCons(tmp7810, tmp7820)

			__e.Return(PrimCons(symlet, tmp7821))
			return

		} else {
			tmp7808 := PrimIsPair(V1841)

			var ifres7782 Obj

			if True == tmp7808 {
				tmp7806 := PrimHead(V1841)

				tmp7807 := PrimEqual(symlet, tmp7806)

				var ifres7784 Obj

				if True == tmp7807 {
					tmp7804 := PrimTail(V1841)

					tmp7805 := PrimIsPair(tmp7804)

					var ifres7786 Obj

					if True == tmp7805 {
						tmp7801 := PrimTail(V1841)

						tmp7802 := PrimTail(tmp7801)

						tmp7803 := PrimIsPair(tmp7802)

						var ifres7788 Obj

						if True == tmp7803 {
							tmp7797 := PrimTail(V1841)

							tmp7798 := PrimTail(tmp7797)

							tmp7799 := PrimTail(tmp7798)

							tmp7800 := PrimIsPair(tmp7799)

							var ifres7790 Obj

							if True == tmp7800 {
								tmp7792 := PrimTail(V1841)

								tmp7793 := PrimTail(tmp7792)

								tmp7794 := PrimTail(tmp7793)

								tmp7795 := PrimTail(tmp7794)

								tmp7796 := PrimEqual(Nil, tmp7795)

								var ifres7791 Obj

								if True == tmp7796 {
									ifres7791 = True

								} else {
									ifres7791 = False

								}

								ifres7790 = ifres7791

							} else {
								ifres7790 = False

							}

							var ifres7789 Obj

							if True == ifres7790 {
								ifres7789 = True

							} else {
								ifres7789 = False

							}

							ifres7788 = ifres7789

						} else {
							ifres7788 = False

						}

						var ifres7787 Obj

						if True == ifres7788 {
							ifres7787 = True

						} else {
							ifres7787 = False

						}

						ifres7786 = ifres7787

					} else {
						ifres7786 = False

					}

					var ifres7785 Obj

					if True == ifres7786 {
						ifres7785 = True

					} else {
						ifres7785 = False

					}

					ifres7784 = ifres7785

				} else {
					ifres7784 = False

				}

				var ifres7783 Obj

				if True == ifres7784 {
					ifres7783 = True

				} else {
					ifres7783 = False

				}

				ifres7782 = ifres7783

			} else {
				ifres7782 = False

			}

			if True == ifres7782 {
				tmp7779 := PrimTail(V1841)

				tmp7780 := PrimHead(tmp7779)

				tmp7781 := PrimIsVariable(tmp7780)

				if True == tmp7781 {
					__e.Return(V1841)
					return
				} else {
					tmp7776 := PrimTail(V1841)

					tmp7777 := PrimHead(tmp7776)

					tmp7778 := Call(__e, PrimNS2Value(symshen_4app), tmp7777, MakeString(" is not a variable\n"), symshen_4s)

					__e.Return(PrimSimpleError(tmp7778))
					return

				}

			} else {
				__e.Return(V1841)
				return
			}

		}

	}, 1)

	tmp7849 := Call(__e, PrimNS2Value(symdef), symshen_4let_1macro, tmp7772)

	_ = tmp7849

	tmp7850 := MakeNative(func(__e *ControlFlow) {
		V1842 := __e.Get(1)
		_ = V1842
		tmp7908 := PrimIsPair(V1842)

		var ifres7889 Obj

		if True == tmp7908 {
			tmp7906 := PrimHead(V1842)

			tmp7907 := PrimEqual(sym_c_4, tmp7906)

			var ifres7891 Obj

			if True == tmp7907 {
				tmp7904 := PrimTail(V1842)

				tmp7905 := PrimIsPair(tmp7904)

				var ifres7893 Obj

				if True == tmp7905 {
					tmp7901 := PrimTail(V1842)

					tmp7902 := PrimTail(tmp7901)

					tmp7903 := PrimIsPair(tmp7902)

					var ifres7895 Obj

					if True == tmp7903 {
						tmp7897 := PrimTail(V1842)

						tmp7898 := PrimTail(tmp7897)

						tmp7899 := PrimTail(tmp7898)

						tmp7900 := PrimIsPair(tmp7899)

						var ifres7896 Obj

						if True == tmp7900 {
							ifres7896 = True

						} else {
							ifres7896 = False

						}

						ifres7895 = ifres7896

					} else {
						ifres7895 = False

					}

					var ifres7894 Obj

					if True == ifres7895 {
						ifres7894 = True

					} else {
						ifres7894 = False

					}

					ifres7893 = ifres7894

				} else {
					ifres7893 = False

				}

				var ifres7892 Obj

				if True == ifres7893 {
					ifres7892 = True

				} else {
					ifres7892 = False

				}

				ifres7891 = ifres7892

			} else {
				ifres7891 = False

			}

			var ifres7890 Obj

			if True == ifres7891 {
				ifres7890 = True

			} else {
				ifres7890 = False

			}

			ifres7889 = ifres7890

		} else {
			ifres7889 = False

		}

		if True == ifres7889 {
			tmp7881 := PrimTail(V1842)

			tmp7882 := PrimHead(tmp7881)

			tmp7883 := PrimTail(V1842)

			tmp7884 := PrimTail(tmp7883)

			tmp7885 := PrimCons(sym_c_4, tmp7884)

			tmp7886 := Call(__e, PrimNS2Value(symshen_4abs_1macro), tmp7885)

			tmp7887 := PrimCons(tmp7886, Nil)

			tmp7888 := PrimCons(tmp7882, tmp7887)

			__e.Return(PrimCons(symlambda, tmp7888))
			return

		} else {
			tmp7880 := PrimIsPair(V1842)

			var ifres7861 Obj

			if True == tmp7880 {
				tmp7878 := PrimHead(V1842)

				tmp7879 := PrimEqual(sym_c_4, tmp7878)

				var ifres7863 Obj

				if True == tmp7879 {
					tmp7876 := PrimTail(V1842)

					tmp7877 := PrimIsPair(tmp7876)

					var ifres7865 Obj

					if True == tmp7877 {
						tmp7873 := PrimTail(V1842)

						tmp7874 := PrimTail(tmp7873)

						tmp7875 := PrimIsPair(tmp7874)

						var ifres7867 Obj

						if True == tmp7875 {
							tmp7869 := PrimTail(V1842)

							tmp7870 := PrimTail(tmp7869)

							tmp7871 := PrimTail(tmp7870)

							tmp7872 := PrimEqual(Nil, tmp7871)

							var ifres7868 Obj

							if True == tmp7872 {
								ifres7868 = True

							} else {
								ifres7868 = False

							}

							ifres7867 = ifres7868

						} else {
							ifres7867 = False

						}

						var ifres7866 Obj

						if True == ifres7867 {
							ifres7866 = True

						} else {
							ifres7866 = False

						}

						ifres7865 = ifres7866

					} else {
						ifres7865 = False

					}

					var ifres7864 Obj

					if True == ifres7865 {
						ifres7864 = True

					} else {
						ifres7864 = False

					}

					ifres7863 = ifres7864

				} else {
					ifres7863 = False

				}

				var ifres7862 Obj

				if True == ifres7863 {
					ifres7862 = True

				} else {
					ifres7862 = False

				}

				ifres7861 = ifres7862

			} else {
				ifres7861 = False

			}

			if True == ifres7861 {
				tmp7858 := PrimTail(V1842)

				tmp7859 := PrimHead(tmp7858)

				tmp7860 := PrimIsVariable(tmp7859)

				if True == tmp7860 {
					tmp7857 := PrimTail(V1842)

					__e.Return(PrimCons(symlambda, tmp7857))
					return

				} else {
					tmp7854 := PrimTail(V1842)

					tmp7855 := PrimHead(tmp7854)

					tmp7856 := Call(__e, PrimNS2Value(symshen_4app), tmp7855, MakeString(" is not a variable\n"), symshen_4s)

					__e.Return(PrimSimpleError(tmp7856))
					return

				}

			} else {
				__e.Return(V1842)
				return
			}

		}

	}, 1)

	tmp7909 := Call(__e, PrimNS2Value(symdef), symshen_4abs_1macro, tmp7850)

	_ = tmp7909

	tmp7910 := MakeNative(func(__e *ControlFlow) {
		V1845 := __e.Get(1)
		_ = V1845
		tmp8006 := PrimIsPair(V1845)

		var ifres7988 Obj

		if True == tmp8006 {
			tmp8004 := PrimHead(V1845)

			tmp8005 := PrimEqual(symcases, tmp8004)

			var ifres7990 Obj

			if True == tmp8005 {
				tmp8002 := PrimTail(V1845)

				tmp8003 := PrimIsPair(tmp8002)

				var ifres7992 Obj

				if True == tmp8003 {
					tmp7999 := PrimTail(V1845)

					tmp8000 := PrimHead(tmp7999)

					tmp8001 := PrimEqual(True, tmp8000)

					var ifres7994 Obj

					if True == tmp8001 {
						tmp7996 := PrimTail(V1845)

						tmp7997 := PrimTail(tmp7996)

						tmp7998 := PrimIsPair(tmp7997)

						var ifres7995 Obj

						if True == tmp7998 {
							ifres7995 = True

						} else {
							ifres7995 = False

						}

						ifres7994 = ifres7995

					} else {
						ifres7994 = False

					}

					var ifres7993 Obj

					if True == ifres7994 {
						ifres7993 = True

					} else {
						ifres7993 = False

					}

					ifres7992 = ifres7993

				} else {
					ifres7992 = False

				}

				var ifres7991 Obj

				if True == ifres7992 {
					ifres7991 = True

				} else {
					ifres7991 = False

				}

				ifres7990 = ifres7991

			} else {
				ifres7990 = False

			}

			var ifres7989 Obj

			if True == ifres7990 {
				ifres7989 = True

			} else {
				ifres7989 = False

			}

			ifres7988 = ifres7989

		} else {
			ifres7988 = False

		}

		if True == ifres7988 {
			tmp7986 := PrimTail(V1845)

			tmp7987 := PrimTail(tmp7986)

			__e.Return(PrimHead(tmp7987))
			return

		} else {
			tmp7985 := PrimIsPair(V1845)

			var ifres7966 Obj

			if True == tmp7985 {
				tmp7983 := PrimHead(V1845)

				tmp7984 := PrimEqual(symcases, tmp7983)

				var ifres7968 Obj

				if True == tmp7984 {
					tmp7981 := PrimTail(V1845)

					tmp7982 := PrimIsPair(tmp7981)

					var ifres7970 Obj

					if True == tmp7982 {
						tmp7978 := PrimTail(V1845)

						tmp7979 := PrimTail(tmp7978)

						tmp7980 := PrimIsPair(tmp7979)

						var ifres7972 Obj

						if True == tmp7980 {
							tmp7974 := PrimTail(V1845)

							tmp7975 := PrimTail(tmp7974)

							tmp7976 := PrimTail(tmp7975)

							tmp7977 := PrimEqual(Nil, tmp7976)

							var ifres7973 Obj

							if True == tmp7977 {
								ifres7973 = True

							} else {
								ifres7973 = False

							}

							ifres7972 = ifres7973

						} else {
							ifres7972 = False

						}

						var ifres7971 Obj

						if True == ifres7972 {
							ifres7971 = True

						} else {
							ifres7971 = False

						}

						ifres7970 = ifres7971

					} else {
						ifres7970 = False

					}

					var ifres7969 Obj

					if True == ifres7970 {
						ifres7969 = True

					} else {
						ifres7969 = False

					}

					ifres7968 = ifres7969

				} else {
					ifres7968 = False

				}

				var ifres7967 Obj

				if True == ifres7968 {
					ifres7967 = True

				} else {
					ifres7967 = False

				}

				ifres7966 = ifres7967

			} else {
				ifres7966 = False

			}

			if True == ifres7966 {
				tmp7956 := PrimTail(V1845)

				tmp7957 := PrimHead(tmp7956)

				tmp7958 := PrimTail(V1845)

				tmp7959 := PrimTail(tmp7958)

				tmp7960 := PrimHead(tmp7959)

				tmp7961 := PrimCons(MakeString("error: cases exhausted"), Nil)

				tmp7962 := PrimCons(symsimple_1error, tmp7961)

				tmp7963 := PrimCons(tmp7962, Nil)

				tmp7964 := PrimCons(tmp7960, tmp7963)

				tmp7965 := PrimCons(tmp7957, tmp7964)

				__e.Return(PrimCons(symif, tmp7965))
				return

			} else {
				tmp7955 := PrimIsPair(V1845)

				var ifres7942 Obj

				if True == tmp7955 {
					tmp7953 := PrimHead(V1845)

					tmp7954 := PrimEqual(symcases, tmp7953)

					var ifres7944 Obj

					if True == tmp7954 {
						tmp7951 := PrimTail(V1845)

						tmp7952 := PrimIsPair(tmp7951)

						var ifres7946 Obj

						if True == tmp7952 {
							tmp7948 := PrimTail(V1845)

							tmp7949 := PrimTail(tmp7948)

							tmp7950 := PrimIsPair(tmp7949)

							var ifres7947 Obj

							if True == tmp7950 {
								ifres7947 = True

							} else {
								ifres7947 = False

							}

							ifres7946 = ifres7947

						} else {
							ifres7946 = False

						}

						var ifres7945 Obj

						if True == ifres7946 {
							ifres7945 = True

						} else {
							ifres7945 = False

						}

						ifres7944 = ifres7945

					} else {
						ifres7944 = False

					}

					var ifres7943 Obj

					if True == ifres7944 {
						ifres7943 = True

					} else {
						ifres7943 = False

					}

					ifres7942 = ifres7943

				} else {
					ifres7942 = False

				}

				if True == ifres7942 {
					tmp7929 := PrimTail(V1845)

					tmp7930 := PrimHead(tmp7929)

					tmp7931 := PrimTail(V1845)

					tmp7932 := PrimTail(tmp7931)

					tmp7933 := PrimHead(tmp7932)

					tmp7934 := PrimTail(V1845)

					tmp7935 := PrimTail(tmp7934)

					tmp7936 := PrimTail(tmp7935)

					tmp7937 := PrimCons(symcases, tmp7936)

					tmp7938 := Call(__e, PrimNS2Value(symshen_4cases_1macro), tmp7937)

					tmp7939 := PrimCons(tmp7938, Nil)

					tmp7940 := PrimCons(tmp7933, tmp7939)

					tmp7941 := PrimCons(tmp7930, tmp7940)

					__e.Return(PrimCons(symif, tmp7941))
					return

				} else {
					tmp7928 := PrimIsPair(V1845)

					var ifres7915 Obj

					if True == tmp7928 {
						tmp7926 := PrimHead(V1845)

						tmp7927 := PrimEqual(symcases, tmp7926)

						var ifres7917 Obj

						if True == tmp7927 {
							tmp7924 := PrimTail(V1845)

							tmp7925 := PrimIsPair(tmp7924)

							var ifres7919 Obj

							if True == tmp7925 {
								tmp7921 := PrimTail(V1845)

								tmp7922 := PrimTail(tmp7921)

								tmp7923 := PrimEqual(Nil, tmp7922)

								var ifres7920 Obj

								if True == tmp7923 {
									ifres7920 = True

								} else {
									ifres7920 = False

								}

								ifres7919 = ifres7920

							} else {
								ifres7919 = False

							}

							var ifres7918 Obj

							if True == ifres7919 {
								ifres7918 = True

							} else {
								ifres7918 = False

							}

							ifres7917 = ifres7918

						} else {
							ifres7917 = False

						}

						var ifres7916 Obj

						if True == ifres7917 {
							ifres7916 = True

						} else {
							ifres7916 = False

						}

						ifres7915 = ifres7916

					} else {
						ifres7915 = False

					}

					if True == ifres7915 {
						__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
						return
					} else {
						__e.Return(V1845)
						return
					}

				}

			}

		}

	}, 1)

	tmp8007 := Call(__e, PrimNS2Value(symdef), symshen_4cases_1macro, tmp7910)

	_ = tmp8007

	tmp8008 := MakeNative(func(__e *ControlFlow) {
		V1846 := __e.Get(1)
		_ = V1846
		tmp8065 := PrimIsPair(V1846)

		var ifres8052 Obj

		if True == tmp8065 {
			tmp8063 := PrimHead(V1846)

			tmp8064 := PrimEqual(symtime, tmp8063)

			var ifres8054 Obj

			if True == tmp8064 {
				tmp8061 := PrimTail(V1846)

				tmp8062 := PrimIsPair(tmp8061)

				var ifres8056 Obj

				if True == tmp8062 {
					tmp8058 := PrimTail(V1846)

					tmp8059 := PrimTail(tmp8058)

					tmp8060 := PrimEqual(Nil, tmp8059)

					var ifres8057 Obj

					if True == tmp8060 {
						ifres8057 = True

					} else {
						ifres8057 = False

					}

					ifres8056 = ifres8057

				} else {
					ifres8056 = False

				}

				var ifres8055 Obj

				if True == ifres8056 {
					ifres8055 = True

				} else {
					ifres8055 = False

				}

				ifres8054 = ifres8055

			} else {
				ifres8054 = False

			}

			var ifres8053 Obj

			if True == ifres8054 {
				ifres8053 = True

			} else {
				ifres8053 = False

			}

			ifres8052 = ifres8053

		} else {
			ifres8052 = False

		}

		if True == ifres8052 {
			tmp8010 := Call(__e, PrimNS2Value(symprotect), symStart)

			tmp8011 := PrimCons(symrun, Nil)

			tmp8012 := PrimCons(symget_1time, tmp8011)

			tmp8013 := Call(__e, PrimNS2Value(symprotect), symResult)

			tmp8014 := PrimTail(V1846)

			tmp8015 := PrimHead(tmp8014)

			tmp8016 := Call(__e, PrimNS2Value(symprotect), symFinish)

			tmp8017 := PrimCons(symrun, Nil)

			tmp8018 := PrimCons(symget_1time, tmp8017)

			tmp8019 := Call(__e, PrimNS2Value(symprotect), symTime)

			tmp8020 := Call(__e, PrimNS2Value(symprotect), symFinish)

			tmp8021 := Call(__e, PrimNS2Value(symprotect), symStart)

			tmp8022 := PrimCons(tmp8021, Nil)

			tmp8023 := PrimCons(tmp8020, tmp8022)

			tmp8024 := PrimCons(sym_1, tmp8023)

			tmp8025 := Call(__e, PrimNS2Value(symprotect), symMessage)

			tmp8026 := Call(__e, PrimNS2Value(symprotect), symTime)

			tmp8027 := PrimCons(tmp8026, Nil)

			tmp8028 := PrimCons(symstr, tmp8027)

			tmp8029 := PrimCons(MakeString(" secs\n"), Nil)

			tmp8030 := PrimCons(tmp8028, tmp8029)

			tmp8031 := PrimCons(symcn, tmp8030)

			tmp8032 := PrimCons(tmp8031, Nil)

			tmp8033 := PrimCons(MakeString("\nrun time: "), tmp8032)

			tmp8034 := PrimCons(symcn, tmp8033)

			tmp8035 := PrimCons(symstoutput, Nil)

			tmp8036 := PrimCons(tmp8035, Nil)

			tmp8037 := PrimCons(tmp8034, tmp8036)

			tmp8038 := PrimCons(sympr, tmp8037)

			tmp8039 := Call(__e, PrimNS2Value(symprotect), symResult)

			tmp8040 := PrimCons(tmp8039, Nil)

			tmp8041 := PrimCons(tmp8038, tmp8040)

			tmp8042 := PrimCons(tmp8025, tmp8041)

			tmp8043 := PrimCons(tmp8024, tmp8042)

			tmp8044 := PrimCons(tmp8019, tmp8043)

			tmp8045 := PrimCons(tmp8018, tmp8044)

			tmp8046 := PrimCons(tmp8016, tmp8045)

			tmp8047 := PrimCons(tmp8015, tmp8046)

			tmp8048 := PrimCons(tmp8013, tmp8047)

			tmp8049 := PrimCons(tmp8012, tmp8048)

			tmp8050 := PrimCons(tmp8010, tmp8049)

			tmp8051 := PrimCons(symlet, tmp8050)

			__e.TailApply(PrimNS2Value(symshen_4let_1macro), tmp8051)
			return

		} else {
			__e.Return(V1846)
			return
		}

	}, 1)

	tmp8066 := Call(__e, PrimNS2Value(symdef), symshen_4timer_1macro, tmp8008)

	_ = tmp8066

	tmp8067 := MakeNative(func(__e *ControlFlow) {
		V1847 := __e.Get(1)
		_ = V1847
		tmp8074 := PrimIsPair(V1847)

		if True == tmp8074 {
			tmp8069 := PrimHead(V1847)

			tmp8070 := PrimTail(V1847)

			tmp8071 := Call(__e, PrimNS2Value(symshen_4tuple_1up), tmp8070)

			tmp8072 := PrimCons(tmp8071, Nil)

			tmp8073 := PrimCons(tmp8069, tmp8072)

			__e.Return(PrimCons(sym_8p, tmp8073))
			return

		} else {
			__e.Return(V1847)
			return
		}

	}, 1)

	tmp8075 := Call(__e, PrimNS2Value(symdef), symshen_4tuple_1up, tmp8067)

	_ = tmp8075

	tmp8076 := MakeNative(func(__e *ControlFlow) {
		V1848 := __e.Get(1)
		_ = V1848
		tmp8181 := PrimIsPair(V1848)

		var ifres8155 Obj

		if True == tmp8181 {
			tmp8179 := PrimHead(V1848)

			tmp8180 := PrimEqual(symput, tmp8179)

			var ifres8157 Obj

			if True == tmp8180 {
				tmp8177 := PrimTail(V1848)

				tmp8178 := PrimIsPair(tmp8177)

				var ifres8159 Obj

				if True == tmp8178 {
					tmp8174 := PrimTail(V1848)

					tmp8175 := PrimTail(tmp8174)

					tmp8176 := PrimIsPair(tmp8175)

					var ifres8161 Obj

					if True == tmp8176 {
						tmp8170 := PrimTail(V1848)

						tmp8171 := PrimTail(tmp8170)

						tmp8172 := PrimTail(tmp8171)

						tmp8173 := PrimIsPair(tmp8172)

						var ifres8163 Obj

						if True == tmp8173 {
							tmp8165 := PrimTail(V1848)

							tmp8166 := PrimTail(tmp8165)

							tmp8167 := PrimTail(tmp8166)

							tmp8168 := PrimTail(tmp8167)

							tmp8169 := PrimEqual(Nil, tmp8168)

							var ifres8164 Obj

							if True == tmp8169 {
								ifres8164 = True

							} else {
								ifres8164 = False

							}

							ifres8163 = ifres8164

						} else {
							ifres8163 = False

						}

						var ifres8162 Obj

						if True == ifres8163 {
							ifres8162 = True

						} else {
							ifres8162 = False

						}

						ifres8161 = ifres8162

					} else {
						ifres8161 = False

					}

					var ifres8160 Obj

					if True == ifres8161 {
						ifres8160 = True

					} else {
						ifres8160 = False

					}

					ifres8159 = ifres8160

				} else {
					ifres8159 = False

				}

				var ifres8158 Obj

				if True == ifres8159 {
					ifres8158 = True

				} else {
					ifres8158 = False

				}

				ifres8157 = ifres8158

			} else {
				ifres8157 = False

			}

			var ifres8156 Obj

			if True == ifres8157 {
				ifres8156 = True

			} else {
				ifres8156 = False

			}

			ifres8155 = ifres8156

		} else {
			ifres8155 = False

		}

		if True == ifres8155 {
			tmp8140 := PrimTail(V1848)

			tmp8141 := PrimHead(tmp8140)

			tmp8142 := PrimTail(V1848)

			tmp8143 := PrimTail(tmp8142)

			tmp8144 := PrimHead(tmp8143)

			tmp8145 := PrimTail(V1848)

			tmp8146 := PrimTail(tmp8145)

			tmp8147 := PrimTail(tmp8146)

			tmp8148 := PrimHead(tmp8147)

			tmp8149 := PrimCons(sym_dproperty_1vector_d, Nil)

			tmp8150 := PrimCons(symvalue, tmp8149)

			tmp8151 := PrimCons(tmp8150, Nil)

			tmp8152 := PrimCons(tmp8148, tmp8151)

			tmp8153 := PrimCons(tmp8144, tmp8152)

			tmp8154 := PrimCons(tmp8141, tmp8153)

			__e.Return(PrimCons(symput, tmp8154))
			return

		} else {
			tmp8139 := PrimIsPair(V1848)

			var ifres8120 Obj

			if True == tmp8139 {
				tmp8137 := PrimHead(V1848)

				tmp8138 := PrimEqual(symget, tmp8137)

				var ifres8122 Obj

				if True == tmp8138 {
					tmp8135 := PrimTail(V1848)

					tmp8136 := PrimIsPair(tmp8135)

					var ifres8124 Obj

					if True == tmp8136 {
						tmp8132 := PrimTail(V1848)

						tmp8133 := PrimTail(tmp8132)

						tmp8134 := PrimIsPair(tmp8133)

						var ifres8126 Obj

						if True == tmp8134 {
							tmp8128 := PrimTail(V1848)

							tmp8129 := PrimTail(tmp8128)

							tmp8130 := PrimTail(tmp8129)

							tmp8131 := PrimEqual(Nil, tmp8130)

							var ifres8127 Obj

							if True == tmp8131 {
								ifres8127 = True

							} else {
								ifres8127 = False

							}

							ifres8126 = ifres8127

						} else {
							ifres8126 = False

						}

						var ifres8125 Obj

						if True == ifres8126 {
							ifres8125 = True

						} else {
							ifres8125 = False

						}

						ifres8124 = ifres8125

					} else {
						ifres8124 = False

					}

					var ifres8123 Obj

					if True == ifres8124 {
						ifres8123 = True

					} else {
						ifres8123 = False

					}

					ifres8122 = ifres8123

				} else {
					ifres8122 = False

				}

				var ifres8121 Obj

				if True == ifres8122 {
					ifres8121 = True

				} else {
					ifres8121 = False

				}

				ifres8120 = ifres8121

			} else {
				ifres8120 = False

			}

			if True == ifres8120 {
				tmp8110 := PrimTail(V1848)

				tmp8111 := PrimHead(tmp8110)

				tmp8112 := PrimTail(V1848)

				tmp8113 := PrimTail(tmp8112)

				tmp8114 := PrimHead(tmp8113)

				tmp8115 := PrimCons(sym_dproperty_1vector_d, Nil)

				tmp8116 := PrimCons(symvalue, tmp8115)

				tmp8117 := PrimCons(tmp8116, Nil)

				tmp8118 := PrimCons(tmp8114, tmp8117)

				tmp8119 := PrimCons(tmp8111, tmp8118)

				__e.Return(PrimCons(symget, tmp8119))
				return

			} else {
				tmp8109 := PrimIsPair(V1848)

				var ifres8090 Obj

				if True == tmp8109 {
					tmp8107 := PrimHead(V1848)

					tmp8108 := PrimEqual(symunput, tmp8107)

					var ifres8092 Obj

					if True == tmp8108 {
						tmp8105 := PrimTail(V1848)

						tmp8106 := PrimIsPair(tmp8105)

						var ifres8094 Obj

						if True == tmp8106 {
							tmp8102 := PrimTail(V1848)

							tmp8103 := PrimTail(tmp8102)

							tmp8104 := PrimIsPair(tmp8103)

							var ifres8096 Obj

							if True == tmp8104 {
								tmp8098 := PrimTail(V1848)

								tmp8099 := PrimTail(tmp8098)

								tmp8100 := PrimTail(tmp8099)

								tmp8101 := PrimEqual(Nil, tmp8100)

								var ifres8097 Obj

								if True == tmp8101 {
									ifres8097 = True

								} else {
									ifres8097 = False

								}

								ifres8096 = ifres8097

							} else {
								ifres8096 = False

							}

							var ifres8095 Obj

							if True == ifres8096 {
								ifres8095 = True

							} else {
								ifres8095 = False

							}

							ifres8094 = ifres8095

						} else {
							ifres8094 = False

						}

						var ifres8093 Obj

						if True == ifres8094 {
							ifres8093 = True

						} else {
							ifres8093 = False

						}

						ifres8092 = ifres8093

					} else {
						ifres8092 = False

					}

					var ifres8091 Obj

					if True == ifres8092 {
						ifres8091 = True

					} else {
						ifres8091 = False

					}

					ifres8090 = ifres8091

				} else {
					ifres8090 = False

				}

				if True == ifres8090 {
					tmp8080 := PrimTail(V1848)

					tmp8081 := PrimHead(tmp8080)

					tmp8082 := PrimTail(V1848)

					tmp8083 := PrimTail(tmp8082)

					tmp8084 := PrimHead(tmp8083)

					tmp8085 := PrimCons(sym_dproperty_1vector_d, Nil)

					tmp8086 := PrimCons(symvalue, tmp8085)

					tmp8087 := PrimCons(tmp8086, Nil)

					tmp8088 := PrimCons(tmp8084, tmp8087)

					tmp8089 := PrimCons(tmp8081, tmp8088)

					__e.Return(PrimCons(symunput, tmp8089))
					return

				} else {
					__e.Return(V1848)
					return
				}

			}

		}

	}, 1)

	tmp8182 := Call(__e, PrimNS2Value(symdef), symshen_4put_cget_1macro, tmp8076)

	_ = tmp8182

	tmp8183 := MakeNative(func(__e *ControlFlow) {
		V1849 := __e.Get(1)
		_ = V1849
		tmp8184 := PrimNS3Value(sym_dmacros_d)

		tmp8185 := Call(__e, PrimNS2Value(symassoc), V1849, tmp8184)

		tmp8186 := PrimNS3Value(sym_dmacros_d)

		tmp8187 := Call(__e, PrimNS2Value(symremove), tmp8185, tmp8186)

		tmp8188 := PrimNS3Set(sym_dmacros_d, tmp8187)

		_ = tmp8188

		__e.Return(V1849)
		return

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), symundefmacro, tmp8183)
	return

}, 0)
