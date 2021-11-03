package main

import . "github.com/tiancaiamao/shen-go/cora"

var SysMain = MakeNative(func(__e *ControlFlow) {
	tmp94 := MakeNative(func(__e *ControlFlow) {
		V3354 := __e.Get(1)
		_ = V3354
		__e.TailApply(V3354)
		return
	}, 1)

	tmp95 := Call(__e, PrimNS2Value(symdef), symthaw, tmp94)

	_ = tmp95

	tmp96 := MakeNative(func(__e *ControlFlow) {
		V3355 := __e.Get(1)
		_ = V3355
		tmp97 := Call(__e, PrimNS2Value(symmacroexpand), V3355)

		tmp98 := Call(__e, PrimNS2Value(symshen_4find_1types), V3355)

		tmp99 := Call(__e, PrimNS2Value(symshen_4process_1applications), tmp97, tmp98)

		tmp100 := Call(__e, PrimNS2Value(symshen_4shen_1_6kl), tmp99)

		__e.TailApply(PrimNS2Value(symeval_1kl), tmp100)
		return

	}, 1)

	tmp101 := Call(__e, PrimNS2Value(symdef), symeval, tmp96)

	_ = tmp101

	tmp102 := MakeNative(func(__e *ControlFlow) {
		V3356 := __e.Get(1)
		_ = V3356
		tmp109 := PrimEqual(symnull, V3356)

		if True == tmp109 {
			__e.Return(Nil)
			return
		} else {
			tmp104 := MakeNative(func(__e *ControlFlow) {
				tmp105 := PrimNS3Value(sym_dproperty_1vector_d)

				__e.TailApply(PrimNS2Value(symget), V3356, symshen_4external_1symbols, tmp105)
				return

			}, 0)

			tmp106 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp107 := Call(__e, PrimNS2Value(symshen_4app), V3356, MakeString(" does not exist.\n;"), symshen_4a)

				tmp108 := PrimStringConcat(MakeString("package "), tmp107)

				__e.Return(PrimSimpleError(tmp108))
				return

			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp104, tmp106)
			return

		}

	}, 1)

	tmp110 := Call(__e, PrimNS2Value(symdef), symexternal, tmp102)

	_ = tmp110

	tmp111 := MakeNative(func(__e *ControlFlow) {
		V3357 := __e.Get(1)
		_ = V3357
		tmp118 := PrimEqual(symnull, V3357)

		if True == tmp118 {
			__e.Return(Nil)
			return
		} else {
			tmp113 := MakeNative(func(__e *ControlFlow) {
				tmp114 := PrimNS3Value(sym_dproperty_1vector_d)

				__e.TailApply(PrimNS2Value(symget), V3357, symshen_4internal_1symbols, tmp114)
				return

			}, 0)

			tmp115 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp116 := Call(__e, PrimNS2Value(symshen_4app), V3357, MakeString(" does not exist.\n;"), symshen_4a)

				tmp117 := PrimStringConcat(MakeString("package "), tmp116)

				__e.Return(PrimSimpleError(tmp117))
				return

			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp113, tmp115)
			return

		}

	}, 1)

	tmp119 := Call(__e, PrimNS2Value(symdef), syminternal, tmp111)

	_ = tmp119

	tmp120 := MakeNative(func(__e *ControlFlow) {
		V3358 := __e.Get(1)
		_ = V3358
		V3359 := __e.Get(2)
		_ = V3359
		tmp122 := Call(__e, V3358, V3359)

		if True == tmp122 {
			__e.TailApply(PrimNS2Value(symfail))
			return
		} else {
			__e.Return(V3359)
			return
		}

	}, 2)

	tmp123 := Call(__e, PrimNS2Value(symdef), symfail_1if, tmp120)

	_ = tmp123

	tmp124 := MakeNative(func(__e *ControlFlow) {
		V3360 := __e.Get(1)
		_ = V3360
		V3361 := __e.Get(2)
		_ = V3361
		__e.Return(PrimStringConcat(V3360, V3361))
		return
	}, 2)

	tmp125 := Call(__e, PrimNS2Value(symdef), sym_8s, tmp124)

	_ = tmp125

	tmp126 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dtc_d))
		return
	}, 0)

	tmp127 := Call(__e, PrimNS2Value(symdef), symtc_2, tmp126)

	_ = tmp127

	tmp128 := MakeNative(func(__e *ControlFlow) {
		V3362 := __e.Get(1)
		_ = V3362
		tmp129 := MakeNative(func(__e *ControlFlow) {
			tmp130 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V3362, symshen_4source, tmp130)
			return

		}, 0)

		tmp131 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp132 := Call(__e, PrimNS2Value(symshen_4app), V3362, MakeString(" not found.\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp132))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp129, tmp131)
		return

	}, 1)

	tmp133 := Call(__e, PrimNS2Value(symdef), symps, tmp128)

	_ = tmp133

	tmp134 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dstinput_d))
		return
	}, 0)

	tmp135 := Call(__e, PrimNS2Value(symdef), symstinput, tmp134)

	_ = tmp135

	tmp136 := MakeNative(func(__e *ControlFlow) {
		V3363 := __e.Get(1)
		_ = V3363
		tmp137 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp138 := MakeNative(func(__e *ControlFlow) {
				ZeroStamp := __e.Get(1)
				_ = ZeroStamp
				tmp139 := MakeNative(func(__e *ControlFlow) {
					Standard := __e.Get(1)
					_ = Standard
					__e.Return(Standard)
					return
				}, 1)

				tmp143 := PrimEqual(V3363, MakeNumber(0))

				var ifres140 Obj

				if True == tmp143 {
					ifres140 = ZeroStamp

				} else {
					tmp141 := Call(__e, PrimNS2Value(symfail))

					tmp142 := Call(__e, PrimNS2Value(symshen_4fillvector), ZeroStamp, MakeNumber(1), V3363, tmp141)

					ifres140 = tmp142

				}

				__e.TailApply(tmp139, ifres140)
				return

			}, 1)

			tmp144 := PrimVectorSet(Vector, MakeNumber(0), V3363)

			__e.TailApply(tmp138, tmp144)
			return

		}, 1)

		tmp145 := PrimNumberAdd(V3363, MakeNumber(1))

		tmp146 := PrimAbsvector(tmp145)

		__e.TailApply(tmp137, tmp146)
		return

	}, 1)

	tmp147 := Call(__e, PrimNS2Value(symdef), symvector, tmp136)

	_ = tmp147

	tmp148 := MakeNative(func(__e *ControlFlow) {
		V3365 := __e.Get(1)
		_ = V3365
		V3366 := __e.Get(2)
		_ = V3366
		V3367 := __e.Get(3)
		_ = V3367
		V3368 := __e.Get(4)
		_ = V3368
		tmp152 := PrimEqual(V3366, V3367)

		if True == tmp152 {
			__e.Return(PrimVectorSet(V3365, V3367, V3368))
			return
		} else {
			tmp150 := PrimVectorSet(V3365, V3366, V3368)

			tmp151 := PrimNumberAdd(MakeNumber(1), V3366)

			__e.TailApply(PrimNS2Value(symshen_4fillvector), tmp150, tmp151, V3367, V3368)
			return

		}

	}, 4)

	tmp153 := Call(__e, PrimNS2Value(symdef), symshen_4fillvector, tmp148)

	_ = tmp153

	tmp154 := MakeNative(func(__e *ControlFlow) {
		V3369 := __e.Get(1)
		_ = V3369
		tmp161 := PrimIsVector(V3369)

		if True == tmp161 {
			tmp157 := MakeNative(func(__e *ControlFlow) {
				tmp158 := PrimVectorGet(V3369, MakeNumber(0))

				__e.Return(PrimGreatEqual(tmp158, MakeNumber(0)))
				return

			}, 0)

			tmp159 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			tmp160 := Call(__e, PrimNS1Value(symtry_1catch), tmp157, tmp159)

			if True == tmp160 {
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

	tmp162 := Call(__e, PrimNS2Value(symdef), symvector_2, tmp154)

	_ = tmp162

	tmp163 := MakeNative(func(__e *ControlFlow) {
		V3370 := __e.Get(1)
		_ = V3370
		V3371 := __e.Get(2)
		_ = V3371
		V3372 := __e.Get(3)
		_ = V3372
		tmp165 := PrimEqual(V3371, MakeNumber(0))

		if True == tmp165 {
			__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
			return
		} else {
			__e.Return(PrimVectorSet(V3370, V3371, V3372))
			return
		}

	}, 3)

	tmp166 := Call(__e, PrimNS2Value(symdef), symvector_1_6, tmp163)

	_ = tmp166

	tmp167 := MakeNative(func(__e *ControlFlow) {
		V3373 := __e.Get(1)
		_ = V3373
		V3374 := __e.Get(2)
		_ = V3374
		tmp174 := PrimEqual(V3374, MakeNumber(0))

		if True == tmp174 {
			__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
			return
		} else {
			tmp169 := MakeNative(func(__e *ControlFlow) {
				VectorElement := __e.Get(1)
				_ = VectorElement
				tmp171 := Call(__e, PrimNS2Value(symfail))

				tmp172 := PrimEqual(VectorElement, tmp171)

				if True == tmp172 {
					__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
					return
				} else {
					__e.Return(VectorElement)
					return
				}

			}, 1)

			tmp173 := PrimVectorGet(V3373, V3374)

			__e.TailApply(tmp169, tmp173)
			return

		}

	}, 2)

	tmp175 := Call(__e, PrimNS2Value(symdef), sym_5_1vector, tmp167)

	_ = tmp175

	tmp176 := MakeNative(func(__e *ControlFlow) {
		V3375 := __e.Get(1)
		_ = V3375
		tmp180 := PrimIsInteger(V3375)

		if True == tmp180 {
			tmp179 := PrimGreatEqual(V3375, MakeNumber(0))

			if True == tmp179 {
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

	tmp181 := Call(__e, PrimNS2Value(symdef), symshen_4posint_2, tmp176)

	_ = tmp181

	tmp182 := MakeNative(func(__e *ControlFlow) {
		V3376 := __e.Get(1)
		_ = V3376
		__e.Return(PrimVectorGet(V3376, MakeNumber(0)))
		return
	}, 1)

	tmp183 := Call(__e, PrimNS2Value(symdef), symlimit, tmp182)

	_ = tmp183

	tmp184 := MakeNative(func(__e *ControlFlow) {
		V3380 := __e.Get(1)
		_ = V3380
		tmp193 := Call(__e, PrimNS2Value(symshen_4_7string_2), V3380)

		if True == tmp193 {
			tmp190 := Call(__e, PrimNS2Value(symhdstr), V3380)

			tmp191 := PrimStringToNumber(tmp190)

			tmp192 := Call(__e, PrimNS2Value(symshen_4alpha_2), tmp191)

			if True == tmp192 {
				tmp188 := PrimTailString(V3380)

				tmp189 := Call(__e, PrimNS2Value(symshen_4alphanums_2), tmp188)

				if True == tmp189 {
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
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-symbol?")))
			return
		}

	}, 1)

	tmp194 := Call(__e, PrimNS2Value(symdef), symshen_4analyse_1symbol_2, tmp184)

	_ = tmp194

	tmp195 := MakeNative(func(__e *ControlFlow) {
		V3383 := __e.Get(1)
		_ = V3383
		tmp210 := PrimEqual(MakeString(""), V3383)

		if True == tmp210 {
			__e.Return(True)
			return
		} else {
			tmp209 := Call(__e, PrimNS2Value(symshen_4_7string_2), V3383)

			if True == tmp209 {
				tmp198 := MakeNative(func(__e *ControlFlow) {
					N := __e.Get(1)
					_ = N
					tmp206 := Call(__e, PrimNS2Value(symshen_4alpha_2), N)

					var ifres203 Obj

					if True == tmp206 {
						ifres203 = True

					} else {
						tmp205 := Call(__e, PrimNS2Value(symshen_4digit_2), N)

						var ifres204 Obj

						if True == tmp205 {
							ifres204 = True

						} else {
							ifres204 = False

						}

						ifres203 = ifres204

					}

					if True == ifres203 {
						tmp201 := PrimTailString(V3383)

						tmp202 := Call(__e, PrimNS2Value(symshen_4alphanums_2), tmp201)

						if True == tmp202 {
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

				tmp207 := Call(__e, PrimNS2Value(symhdstr), V3383)

				tmp208 := PrimStringToNumber(tmp207)

				__e.TailApply(tmp198, tmp208)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.alphanums?")))
				return
			}

		}

	}, 1)

	tmp211 := Call(__e, PrimNS2Value(symdef), symshen_4alphanums_2, tmp195)

	_ = tmp211

	tmp212 := MakeNative(func(__e *ControlFlow) {
		V3384 := __e.Get(1)
		_ = V3384
		tmp224 := Call(__e, PrimNS2Value(symboolean_2), V3384)

		var ifres218 Obj

		if True == tmp224 {
			ifres218 = True

		} else {
			tmp223 := PrimIsNumber(V3384)

			var ifres220 Obj

			if True == tmp223 {
				ifres220 = True

			} else {
				tmp222 := PrimIsString(V3384)

				var ifres221 Obj

				if True == tmp222 {
					ifres221 = True

				} else {
					ifres221 = False

				}

				ifres220 = ifres221

			}

			var ifres219 Obj

			if True == ifres220 {
				ifres219 = True

			} else {
				ifres219 = False

			}

			ifres218 = ifres219

		}

		if True == ifres218 {
			__e.Return(False)
			return
		} else {
			tmp214 := MakeNative(func(__e *ControlFlow) {
				tmp215 := MakeNative(func(__e *ControlFlow) {
					String := __e.Get(1)
					_ = String
					__e.TailApply(PrimNS2Value(symshen_4analyse_1variable_2), String)
					return
				}, 1)

				tmp216 := PrimStr(V3384)

				__e.TailApply(tmp215, tmp216)
				return

			}, 0)

			tmp217 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp214, tmp217)
			return

		}

	}, 1)

	tmp225 := Call(__e, PrimNS2Value(symdef), symvariable_2, tmp212)

	_ = tmp225

	tmp226 := MakeNative(func(__e *ControlFlow) {
		V3387 := __e.Get(1)
		_ = V3387
		tmp235 := Call(__e, PrimNS2Value(symshen_4_7string_2), V3387)

		if True == tmp235 {
			tmp232 := Call(__e, PrimNS2Value(symhdstr), V3387)

			tmp233 := PrimStringToNumber(tmp232)

			tmp234 := Call(__e, PrimNS2Value(symshen_4uppercase_2), tmp233)

			if True == tmp234 {
				tmp230 := PrimTailString(V3387)

				tmp231 := Call(__e, PrimNS2Value(symshen_4alphanums_2), tmp230)

				if True == tmp231 {
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
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-variable?")))
			return
		}

	}, 1)

	tmp236 := Call(__e, PrimNS2Value(symdef), symshen_4analyse_1variable_2, tmp226)

	_ = tmp236

	tmp237 := MakeNative(func(__e *ControlFlow) {
		V3388 := __e.Get(1)
		_ = V3388
		tmp238 := PrimNS3Value(symshen_4_dgensym_d)

		tmp239 := PrimNumberAdd(MakeNumber(1), tmp238)

		tmp240 := PrimNS3Set(symshen_4_dgensym_d, tmp239)

		__e.TailApply(PrimNS2Value(symconcat), V3388, tmp240)
		return

	}, 1)

	tmp241 := Call(__e, PrimNS2Value(symdef), symgensym, tmp237)

	_ = tmp241

	tmp242 := MakeNative(func(__e *ControlFlow) {
		V3389 := __e.Get(1)
		_ = V3389
		V3390 := __e.Get(2)
		_ = V3390
		tmp243 := PrimStr(V3389)

		tmp244 := PrimStr(V3390)

		tmp245 := PrimStringConcat(tmp243, tmp244)

		__e.Return(PrimIntern(tmp245))
		return

	}, 2)

	tmp246 := Call(__e, PrimNS2Value(symdef), symconcat, tmp242)

	_ = tmp246

	tmp247 := MakeNative(func(__e *ControlFlow) {
		V3391 := __e.Get(1)
		_ = V3391
		V3392 := __e.Get(2)
		_ = V3392
		tmp248 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp249 := MakeNative(func(__e *ControlFlow) {
				Tag := __e.Get(1)
				_ = Tag
				tmp250 := MakeNative(func(__e *ControlFlow) {
					Fst := __e.Get(1)
					_ = Fst
					tmp251 := MakeNative(func(__e *ControlFlow) {
						Snd := __e.Get(1)
						_ = Snd
						__e.Return(Vector)
						return
					}, 1)

					tmp252 := PrimVectorSet(Vector, MakeNumber(2), V3392)

					__e.TailApply(tmp251, tmp252)
					return

				}, 1)

				tmp253 := PrimVectorSet(Vector, MakeNumber(1), V3391)

				__e.TailApply(tmp250, tmp253)
				return

			}, 1)

			tmp254 := PrimVectorSet(Vector, MakeNumber(0), symshen_4tuple)

			__e.TailApply(tmp249, tmp254)
			return

		}, 1)

		tmp255 := PrimAbsvector(MakeNumber(3))

		__e.TailApply(tmp248, tmp255)
		return

	}, 2)

	tmp256 := Call(__e, PrimNS2Value(symdef), sym_8p, tmp247)

	_ = tmp256

	tmp257 := MakeNative(func(__e *ControlFlow) {
		V3393 := __e.Get(1)
		_ = V3393
		__e.Return(PrimVectorGet(V3393, MakeNumber(1)))
		return
	}, 1)

	tmp258 := Call(__e, PrimNS2Value(symdef), symfst, tmp257)

	_ = tmp258

	tmp259 := MakeNative(func(__e *ControlFlow) {
		V3394 := __e.Get(1)
		_ = V3394
		__e.Return(PrimVectorGet(V3394, MakeNumber(2)))
		return
	}, 1)

	tmp260 := Call(__e, PrimNS2Value(symdef), symsnd, tmp259)

	_ = tmp260

	tmp261 := MakeNative(func(__e *ControlFlow) {
		V3395 := __e.Get(1)
		_ = V3395
		tmp262 := MakeNative(func(__e *ControlFlow) {
			tmp267 := PrimIsVector(V3395)

			if True == tmp267 {
				tmp265 := PrimVectorGet(V3395, MakeNumber(0))

				tmp266 := PrimEqual(symshen_4tuple, tmp265)

				if True == tmp266 {
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

		}, 0)

		tmp268 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp262, tmp268)
		return

	}, 1)

	tmp269 := Call(__e, PrimNS2Value(symdef), symtuple_2, tmp261)

	_ = tmp269

	tmp270 := MakeNative(func(__e *ControlFlow) {
		V3400 := __e.Get(1)
		_ = V3400
		V3401 := __e.Get(2)
		_ = V3401
		tmp277 := PrimEqual(Nil, V3400)

		if True == tmp277 {
			__e.Return(V3401)
			return
		} else {
			tmp276 := PrimIsPair(V3400)

			if True == tmp276 {
				tmp273 := PrimHead(V3400)

				tmp274 := PrimTail(V3400)

				tmp275 := Call(__e, PrimNS2Value(symappend), tmp274, V3401)

				__e.Return(PrimCons(tmp273, tmp275))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to append a non-list")))
				return
			}

		}

	}, 2)

	tmp278 := Call(__e, PrimNS2Value(symdef), symappend, tmp270)

	_ = tmp278

	tmp279 := MakeNative(func(__e *ControlFlow) {
		V3402 := __e.Get(1)
		_ = V3402
		V3403 := __e.Get(2)
		_ = V3403
		tmp280 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp281 := MakeNative(func(__e *ControlFlow) {
				NewVector := __e.Get(1)
				_ = NewVector
				tmp282 := MakeNative(func(__e *ControlFlow) {
					X_7NewVector := __e.Get(1)
					_ = X_7NewVector
					tmp284 := PrimEqual(Limit, MakeNumber(0))

					if True == tmp284 {
						__e.Return(X_7NewVector)
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4_8v_1help), V3403, MakeNumber(1), Limit, X_7NewVector)
						return
					}

				}, 1)

				tmp285 := Call(__e, PrimNS2Value(symvector_1_6), NewVector, MakeNumber(1), V3402)

				__e.TailApply(tmp282, tmp285)
				return

			}, 1)

			tmp286 := PrimNumberAdd(Limit, MakeNumber(1))

			tmp287 := Call(__e, PrimNS2Value(symvector), tmp286)

			__e.TailApply(tmp281, tmp287)
			return

		}, 1)

		tmp288 := Call(__e, PrimNS2Value(symlimit), V3403)

		__e.TailApply(tmp280, tmp288)
		return

	}, 2)

	tmp289 := Call(__e, PrimNS2Value(symdef), sym_8v, tmp279)

	_ = tmp289

	tmp290 := MakeNative(func(__e *ControlFlow) {
		V3405 := __e.Get(1)
		_ = V3405
		V3406 := __e.Get(2)
		_ = V3406
		V3407 := __e.Get(3)
		_ = V3407
		V3408 := __e.Get(4)
		_ = V3408
		tmp296 := PrimEqual(V3406, V3407)

		if True == tmp296 {
			tmp295 := PrimNumberAdd(V3407, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4copyfromvector), V3405, V3408, V3407, tmp295)
			return

		} else {
			tmp292 := PrimNumberAdd(V3406, MakeNumber(1))

			tmp293 := PrimNumberAdd(V3406, MakeNumber(1))

			tmp294 := Call(__e, PrimNS2Value(symshen_4copyfromvector), V3405, V3408, V3406, tmp293)

			__e.TailApply(PrimNS2Value(symshen_4_8v_1help), V3405, tmp292, V3407, tmp294)
			return

		}

	}, 4)

	tmp297 := Call(__e, PrimNS2Value(symdef), symshen_4_8v_1help, tmp290)

	_ = tmp297

	tmp298 := MakeNative(func(__e *ControlFlow) {
		V3409 := __e.Get(1)
		_ = V3409
		V3410 := __e.Get(2)
		_ = V3410
		V3411 := __e.Get(3)
		_ = V3411
		V3412 := __e.Get(4)
		_ = V3412
		tmp299 := MakeNative(func(__e *ControlFlow) {
			tmp300 := Call(__e, PrimNS2Value(sym_5_1vector), V3409, V3411)

			__e.TailApply(PrimNS2Value(symvector_1_6), V3410, V3412, tmp300)
			return

		}, 0)

		tmp301 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V3410)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp299, tmp301)
		return

	}, 4)

	tmp302 := Call(__e, PrimNS2Value(symdef), symshen_4copyfromvector, tmp298)

	_ = tmp302

	tmp303 := MakeNative(func(__e *ControlFlow) {
		V3413 := __e.Get(1)
		_ = V3413
		tmp304 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(sym_5_1vector), V3413, MakeNumber(1))
			return
		}, 0)

		tmp305 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n")))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp304, tmp305)
		return

	}, 1)

	tmp306 := Call(__e, PrimNS2Value(symdef), symhdv, tmp303)

	_ = tmp306

	tmp307 := MakeNative(func(__e *ControlFlow) {
		V3414 := __e.Get(1)
		_ = V3414
		tmp308 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp317 := PrimEqual(Limit, MakeNumber(0))

			if True == tmp317 {
				__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
				return
			} else {
				tmp316 := PrimEqual(Limit, MakeNumber(1))

				if True == tmp316 {
					__e.TailApply(PrimNS2Value(symvector), MakeNumber(0))
					return
				} else {
					tmp311 := MakeNative(func(__e *ControlFlow) {
						NewVector := __e.Get(1)
						_ = NewVector
						tmp312 := PrimNumberSubtract(Limit, MakeNumber(1))

						tmp313 := Call(__e, PrimNS2Value(symvector), tmp312)

						__e.TailApply(PrimNS2Value(symshen_4tlv_1help), V3414, MakeNumber(2), Limit, tmp313)
						return

					}, 1)

					tmp314 := PrimNumberSubtract(Limit, MakeNumber(1))

					tmp315 := Call(__e, PrimNS2Value(symvector), tmp314)

					__e.TailApply(tmp311, tmp315)
					return

				}

			}

		}, 1)

		tmp318 := Call(__e, PrimNS2Value(symlimit), V3414)

		__e.TailApply(tmp308, tmp318)
		return

	}, 1)

	tmp319 := Call(__e, PrimNS2Value(symdef), symtlv, tmp307)

	_ = tmp319

	tmp320 := MakeNative(func(__e *ControlFlow) {
		V3416 := __e.Get(1)
		_ = V3416
		V3417 := __e.Get(2)
		_ = V3417
		V3418 := __e.Get(3)
		_ = V3418
		V3419 := __e.Get(4)
		_ = V3419
		tmp326 := PrimEqual(V3417, V3418)

		if True == tmp326 {
			tmp325 := PrimNumberSubtract(V3418, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4copyfromvector), V3416, V3419, V3418, tmp325)
			return

		} else {
			tmp322 := PrimNumberAdd(V3417, MakeNumber(1))

			tmp323 := PrimNumberSubtract(V3417, MakeNumber(1))

			tmp324 := Call(__e, PrimNS2Value(symshen_4copyfromvector), V3416, V3419, V3417, tmp323)

			__e.TailApply(PrimNS2Value(symshen_4tlv_1help), V3416, tmp322, V3418, tmp324)
			return

		}

	}, 4)

	tmp327 := Call(__e, PrimNS2Value(symdef), symshen_4tlv_1help, tmp320)

	_ = tmp327

	tmp328 := MakeNative(func(__e *ControlFlow) {
		V3431 := __e.Get(1)
		_ = V3431
		V3432 := __e.Get(2)
		_ = V3432
		tmp344 := PrimEqual(Nil, V3432)

		if True == tmp344 {
			__e.Return(Nil)
			return
		} else {
			tmp343 := PrimIsPair(V3432)

			var ifres334 Obj

			if True == tmp343 {
				tmp341 := PrimHead(V3432)

				tmp342 := PrimIsPair(tmp341)

				var ifres336 Obj

				if True == tmp342 {
					tmp338 := PrimHead(V3432)

					tmp339 := PrimHead(tmp338)

					tmp340 := PrimEqual(V3431, tmp339)

					var ifres337 Obj

					if True == tmp340 {
						ifres337 = True

					} else {
						ifres337 = False

					}

					ifres336 = ifres337

				} else {
					ifres336 = False

				}

				var ifres335 Obj

				if True == ifres336 {
					ifres335 = True

				} else {
					ifres335 = False

				}

				ifres334 = ifres335

			} else {
				ifres334 = False

			}

			if True == ifres334 {
				__e.Return(PrimHead(V3432))
				return
			} else {
				tmp333 := PrimIsPair(V3432)

				if True == tmp333 {
					tmp332 := PrimTail(V3432)

					__e.TailApply(PrimNS2Value(symassoc), V3431, tmp332)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("attempt to search a non-list with assoc\n")))
					return
				}

			}

		}

	}, 2)

	tmp345 := Call(__e, PrimNS2Value(symdef), symassoc, tmp328)

	_ = tmp345

	tmp346 := MakeNative(func(__e *ControlFlow) {
		V3435 := __e.Get(1)
		_ = V3435
		tmp350 := PrimEqual(True, V3435)

		if True == tmp350 {
			__e.Return(True)
			return
		} else {
			tmp349 := PrimEqual(False, V3435)

			if True == tmp349 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp351 := Call(__e, PrimNS2Value(symdef), symboolean_2, tmp346)

	_ = tmp351

	tmp352 := MakeNative(func(__e *ControlFlow) {
		V3436 := __e.Get(1)
		_ = V3436
		tmp357 := PrimEqual(MakeNumber(0), V3436)

		if True == tmp357 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp354 := Call(__e, PrimNS2Value(symstoutput))

			tmp355 := Call(__e, PrimNS2Value(sympr), MakeString("\n"), tmp354)

			_ = tmp355

			tmp356 := PrimNumberSubtract(V3436, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symnl), tmp356)
			return

		}

	}, 1)

	tmp358 := Call(__e, PrimNS2Value(symdef), symnl, tmp352)

	_ = tmp358

	tmp359 := MakeNative(func(__e *ControlFlow) {
		V3443 := __e.Get(1)
		_ = V3443
		V3444 := __e.Get(2)
		_ = V3444
		tmp370 := PrimEqual(Nil, V3443)

		if True == tmp370 {
			__e.Return(Nil)
			return
		} else {
			tmp369 := PrimIsPair(V3443)

			if True == tmp369 {
				tmp367 := PrimHead(V3443)

				tmp368 := Call(__e, PrimNS2Value(symelement_2), tmp367, V3444)

				if True == tmp368 {
					tmp366 := PrimTail(V3443)

					__e.TailApply(PrimNS2Value(symdifference), tmp366, V3444)
					return

				} else {
					tmp363 := PrimHead(V3443)

					tmp364 := PrimTail(V3443)

					tmp365 := Call(__e, PrimNS2Value(symdifference), tmp364, V3444)

					__e.Return(PrimCons(tmp363, tmp365))
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to find the difference with a non-list\n")))
				return
			}

		}

	}, 2)

	tmp371 := Call(__e, PrimNS2Value(symdef), symdifference, tmp359)

	_ = tmp371

	tmp372 := MakeNative(func(__e *ControlFlow) {
		V3445 := __e.Get(1)
		_ = V3445
		V3446 := __e.Get(2)
		_ = V3446
		__e.Return(V3446)
		return
	}, 2)

	tmp373 := Call(__e, PrimNS2Value(symdef), symdo, tmp372)

	_ = tmp373

	tmp374 := MakeNative(func(__e *ControlFlow) {
		V3458 := __e.Get(1)
		_ = V3458
		V3459 := __e.Get(2)
		_ = V3459
		tmp385 := PrimEqual(Nil, V3459)

		if True == tmp385 {
			__e.Return(False)
			return
		} else {
			tmp384 := PrimIsPair(V3459)

			var ifres380 Obj

			if True == tmp384 {
				tmp382 := PrimHead(V3459)

				tmp383 := PrimEqual(V3458, tmp382)

				var ifres381 Obj

				if True == tmp383 {
					ifres381 = True

				} else {
					ifres381 = False

				}

				ifres380 = ifres381

			} else {
				ifres380 = False

			}

			if True == ifres380 {
				__e.Return(True)
				return
			} else {
				tmp379 := PrimIsPair(V3459)

				if True == tmp379 {
					tmp378 := PrimTail(V3459)

					__e.TailApply(PrimNS2Value(symelement_2), V3458, tmp378)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("attempt to find an element in a non-list\n")))
					return
				}

			}

		}

	}, 2)

	tmp386 := Call(__e, PrimNS2Value(symdef), symelement_2, tmp374)

	_ = tmp386

	tmp387 := MakeNative(func(__e *ControlFlow) {
		V3462 := __e.Get(1)
		_ = V3462
		tmp389 := PrimEqual(Nil, V3462)

		if True == tmp389 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp390 := Call(__e, PrimNS2Value(symdef), symempty_2, tmp387)

	_ = tmp390

	tmp391 := MakeNative(func(__e *ControlFlow) {
		V3463 := __e.Get(1)
		_ = V3463
		V3464 := __e.Get(2)
		_ = V3464
		tmp392 := Call(__e, V3463, V3464)

		__e.TailApply(PrimNS2Value(symshen_4fix_1help), V3463, V3464, tmp392)
		return

	}, 2)

	tmp393 := Call(__e, PrimNS2Value(symdef), symfix, tmp391)

	_ = tmp393

	tmp394 := MakeNative(func(__e *ControlFlow) {
		V3470 := __e.Get(1)
		_ = V3470
		V3471 := __e.Get(2)
		_ = V3471
		V3472 := __e.Get(3)
		_ = V3472
		tmp397 := PrimEqual(V3471, V3472)

		if True == tmp397 {
			__e.Return(V3472)
			return
		} else {
			tmp396 := Call(__e, V3470, V3472)

			__e.TailApply(PrimNS2Value(symshen_4fix_1help), V3470, V3472, tmp396)
			return

		}

	}, 3)

	tmp398 := Call(__e, PrimNS2Value(symdef), symshen_4fix_1help, tmp394)

	_ = tmp398

	tmp399 := MakeNative(func(__e *ControlFlow) {
		V3473 := __e.Get(1)
		_ = V3473
		V3474 := __e.Get(2)
		_ = V3474
		V3475 := __e.Get(3)
		_ = V3475
		V3476 := __e.Get(4)
		_ = V3476
		tmp400 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp401 := MakeNative(func(__e *ControlFlow) {
				Entry := __e.Get(1)
				_ = Entry
				tmp402 := MakeNative(func(__e *ControlFlow) {
					Change := __e.Get(1)
					_ = Change
					__e.Return(V3475)
					return
				}, 1)

				tmp403 := Call(__e, PrimNS2Value(symshen_4change_1pointer_1value), V3473, V3474, V3475, Entry)

				tmp404 := Call(__e, PrimNS2Value(symvector_1_6), V3476, N, tmp403)

				__e.TailApply(tmp402, tmp404)
				return

			}, 1)

			tmp405 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(sym_5_1vector), V3476, N)
				return
			}, 0)

			tmp406 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)

			tmp407 := Call(__e, PrimNS1Value(symtry_1catch), tmp405, tmp406)

			__e.TailApply(tmp401, tmp407)
			return

		}, 1)

		tmp408 := Call(__e, PrimNS2Value(symlimit), V3476)

		tmp409 := Call(__e, PrimNS2Value(symhash), V3473, tmp408)

		__e.TailApply(tmp400, tmp409)
		return

	}, 4)

	tmp410 := Call(__e, PrimNS2Value(symdef), symput, tmp399)

	_ = tmp410

	tmp411 := MakeNative(func(__e *ControlFlow) {
		V3477 := __e.Get(1)
		_ = V3477
		V3478 := __e.Get(2)
		_ = V3478
		V3479 := __e.Get(3)
		_ = V3479
		tmp412 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp413 := MakeNative(func(__e *ControlFlow) {
				Entry := __e.Get(1)
				_ = Entry
				tmp414 := MakeNative(func(__e *ControlFlow) {
					Change := __e.Get(1)
					_ = Change
					__e.Return(V3477)
					return
				}, 1)

				tmp415 := Call(__e, PrimNS2Value(symshen_4remove_1pointer), V3477, V3478, Entry)

				tmp416 := Call(__e, PrimNS2Value(symvector_1_6), V3479, N, tmp415)

				__e.TailApply(tmp414, tmp416)
				return

			}, 1)

			tmp417 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(sym_5_1vector), V3479, N)
				return
			}, 0)

			tmp418 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(Nil)
				return
			}, 1)

			tmp419 := Call(__e, PrimNS1Value(symtry_1catch), tmp417, tmp418)

			__e.TailApply(tmp413, tmp419)
			return

		}, 1)

		tmp420 := Call(__e, PrimNS2Value(symlimit), V3479)

		tmp421 := Call(__e, PrimNS2Value(symhash), V3477, tmp420)

		__e.TailApply(tmp412, tmp421)
		return

	}, 3)

	tmp422 := Call(__e, PrimNS2Value(symdef), symunput, tmp411)

	_ = tmp422

	tmp423 := MakeNative(func(__e *ControlFlow) {
		V3490 := __e.Get(1)
		_ = V3490
		V3491 := __e.Get(2)
		_ = V3491
		V3492 := __e.Get(3)
		_ = V3492
		tmp467 := PrimEqual(Nil, V3492)

		if True == tmp467 {
			__e.Return(Nil)
			return
		} else {
			tmp466 := PrimIsPair(V3492)

			var ifres431 Obj

			if True == tmp466 {
				tmp464 := PrimHead(V3492)

				tmp465 := PrimIsPair(tmp464)

				var ifres433 Obj

				if True == tmp465 {
					tmp461 := PrimHead(V3492)

					tmp462 := PrimHead(tmp461)

					tmp463 := PrimIsPair(tmp462)

					var ifres435 Obj

					if True == tmp463 {
						tmp457 := PrimHead(V3492)

						tmp458 := PrimHead(tmp457)

						tmp459 := PrimTail(tmp458)

						tmp460 := PrimIsPair(tmp459)

						var ifres437 Obj

						if True == tmp460 {
							tmp452 := PrimHead(V3492)

							tmp453 := PrimHead(tmp452)

							tmp454 := PrimTail(tmp453)

							tmp455 := PrimTail(tmp454)

							tmp456 := PrimEqual(Nil, tmp455)

							var ifres439 Obj

							if True == tmp456 {
								tmp447 := PrimHead(V3492)

								tmp448 := PrimHead(tmp447)

								tmp449 := PrimTail(tmp448)

								tmp450 := PrimHead(tmp449)

								tmp451 := PrimEqual(V3491, tmp450)

								var ifres441 Obj

								if True == tmp451 {
									tmp443 := PrimHead(V3492)

									tmp444 := PrimHead(tmp443)

									tmp445 := PrimHead(tmp444)

									tmp446 := PrimEqual(V3490, tmp445)

									var ifres442 Obj

									if True == tmp446 {
										ifres442 = True

									} else {
										ifres442 = False

									}

									ifres441 = ifres442

								} else {
									ifres441 = False

								}

								var ifres440 Obj

								if True == ifres441 {
									ifres440 = True

								} else {
									ifres440 = False

								}

								ifres439 = ifres440

							} else {
								ifres439 = False

							}

							var ifres438 Obj

							if True == ifres439 {
								ifres438 = True

							} else {
								ifres438 = False

							}

							ifres437 = ifres438

						} else {
							ifres437 = False

						}

						var ifres436 Obj

						if True == ifres437 {
							ifres436 = True

						} else {
							ifres436 = False

						}

						ifres435 = ifres436

					} else {
						ifres435 = False

					}

					var ifres434 Obj

					if True == ifres435 {
						ifres434 = True

					} else {
						ifres434 = False

					}

					ifres433 = ifres434

				} else {
					ifres433 = False

				}

				var ifres432 Obj

				if True == ifres433 {
					ifres432 = True

				} else {
					ifres432 = False

				}

				ifres431 = ifres432

			} else {
				ifres431 = False

			}

			if True == ifres431 {
				__e.Return(PrimTail(V3492))
				return
			} else {
				tmp430 := PrimIsPair(V3492)

				if True == tmp430 {
					tmp427 := PrimHead(V3492)

					tmp428 := PrimTail(V3492)

					tmp429 := Call(__e, PrimNS2Value(symshen_4remove_1pointer), V3490, V3491, tmp428)

					__e.Return(PrimCons(tmp427, tmp429))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-pointer")))
					return
				}

			}

		}

	}, 3)

	tmp468 := Call(__e, PrimNS2Value(symdef), symshen_4remove_1pointer, tmp423)

	_ = tmp468

	tmp469 := MakeNative(func(__e *ControlFlow) {
		V3505 := __e.Get(1)
		_ = V3505
		V3506 := __e.Get(2)
		_ = V3506
		V3507 := __e.Get(3)
		_ = V3507
		V3508 := __e.Get(4)
		_ = V3508
		tmp520 := PrimEqual(Nil, V3508)

		if True == tmp520 {
			tmp517 := PrimCons(V3506, Nil)

			tmp518 := PrimCons(V3505, tmp517)

			tmp519 := PrimCons(tmp518, V3507)

			__e.Return(PrimCons(tmp519, Nil))
			return

		} else {
			tmp516 := PrimIsPair(V3508)

			var ifres481 Obj

			if True == tmp516 {
				tmp514 := PrimHead(V3508)

				tmp515 := PrimIsPair(tmp514)

				var ifres483 Obj

				if True == tmp515 {
					tmp511 := PrimHead(V3508)

					tmp512 := PrimHead(tmp511)

					tmp513 := PrimIsPair(tmp512)

					var ifres485 Obj

					if True == tmp513 {
						tmp507 := PrimHead(V3508)

						tmp508 := PrimHead(tmp507)

						tmp509 := PrimTail(tmp508)

						tmp510 := PrimIsPair(tmp509)

						var ifres487 Obj

						if True == tmp510 {
							tmp502 := PrimHead(V3508)

							tmp503 := PrimHead(tmp502)

							tmp504 := PrimTail(tmp503)

							tmp505 := PrimTail(tmp504)

							tmp506 := PrimEqual(Nil, tmp505)

							var ifres489 Obj

							if True == tmp506 {
								tmp497 := PrimHead(V3508)

								tmp498 := PrimHead(tmp497)

								tmp499 := PrimTail(tmp498)

								tmp500 := PrimHead(tmp499)

								tmp501 := PrimEqual(V3506, tmp500)

								var ifres491 Obj

								if True == tmp501 {
									tmp493 := PrimHead(V3508)

									tmp494 := PrimHead(tmp493)

									tmp495 := PrimHead(tmp494)

									tmp496 := PrimEqual(V3505, tmp495)

									var ifres492 Obj

									if True == tmp496 {
										ifres492 = True

									} else {
										ifres492 = False

									}

									ifres491 = ifres492

								} else {
									ifres491 = False

								}

								var ifres490 Obj

								if True == ifres491 {
									ifres490 = True

								} else {
									ifres490 = False

								}

								ifres489 = ifres490

							} else {
								ifres489 = False

							}

							var ifres488 Obj

							if True == ifres489 {
								ifres488 = True

							} else {
								ifres488 = False

							}

							ifres487 = ifres488

						} else {
							ifres487 = False

						}

						var ifres486 Obj

						if True == ifres487 {
							ifres486 = True

						} else {
							ifres486 = False

						}

						ifres485 = ifres486

					} else {
						ifres485 = False

					}

					var ifres484 Obj

					if True == ifres485 {
						ifres484 = True

					} else {
						ifres484 = False

					}

					ifres483 = ifres484

				} else {
					ifres483 = False

				}

				var ifres482 Obj

				if True == ifres483 {
					ifres482 = True

				} else {
					ifres482 = False

				}

				ifres481 = ifres482

			} else {
				ifres481 = False

			}

			if True == ifres481 {
				tmp477 := PrimHead(V3508)

				tmp478 := PrimHead(tmp477)

				tmp479 := PrimCons(tmp478, V3507)

				tmp480 := PrimTail(V3508)

				__e.Return(PrimCons(tmp479, tmp480))
				return

			} else {
				tmp476 := PrimIsPair(V3508)

				if True == tmp476 {
					tmp473 := PrimHead(V3508)

					tmp474 := PrimTail(V3508)

					tmp475 := Call(__e, PrimNS2Value(symshen_4change_1pointer_1value), V3505, V3506, V3507, tmp474)

					__e.Return(PrimCons(tmp473, tmp475))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.change-pointer-value")))
					return
				}

			}

		}

	}, 4)

	tmp521 := Call(__e, PrimNS2Value(symdef), symshen_4change_1pointer_1value, tmp469)

	_ = tmp521

	tmp522 := MakeNative(func(__e *ControlFlow) {
		V3509 := __e.Get(1)
		_ = V3509
		V3510 := __e.Get(2)
		_ = V3510
		V3511 := __e.Get(3)
		_ = V3511
		tmp523 := MakeNative(func(__e *ControlFlow) {
			N := __e.Get(1)
			_ = N
			tmp524 := MakeNative(func(__e *ControlFlow) {
				Entry := __e.Get(1)
				_ = Entry
				tmp525 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp531 := Call(__e, PrimNS2Value(symempty_2), Result)

					if True == tmp531 {
						tmp527 := Call(__e, PrimNS2Value(symshen_4app), V3509, MakeString("\n"), symshen_4s)

						tmp528 := PrimStringConcat(MakeString(" not found for "), tmp527)

						tmp529 := Call(__e, PrimNS2Value(symshen_4app), V3510, tmp528, symshen_4s)

						tmp530 := PrimStringConcat(MakeString("attribute "), tmp529)

						__e.Return(PrimSimpleError(tmp530))
						return

					} else {
						__e.Return(PrimTail(Result))
						return
					}

				}, 1)

				tmp532 := PrimCons(V3510, Nil)

				tmp533 := PrimCons(V3509, tmp532)

				tmp534 := Call(__e, PrimNS2Value(symassoc), tmp533, Entry)

				__e.TailApply(tmp525, tmp534)
				return

			}, 1)

			tmp535 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(sym_5_1vector), V3511, N)
				return
			}, 0)

			tmp536 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp537 := Call(__e, PrimNS2Value(symshen_4app), V3510, MakeString("\n"), symshen_4s)

				tmp538 := PrimStringConcat(MakeString(" has no attributes: "), tmp537)

				tmp539 := Call(__e, PrimNS2Value(symshen_4app), V3509, tmp538, symshen_4a)

				__e.Return(PrimSimpleError(tmp539))
				return

			}, 1)

			tmp540 := Call(__e, PrimNS1Value(symtry_1catch), tmp535, tmp536)

			__e.TailApply(tmp524, tmp540)
			return

		}, 1)

		tmp541 := Call(__e, PrimNS2Value(symlimit), V3511)

		tmp542 := Call(__e, PrimNS2Value(symhash), V3509, tmp541)

		__e.TailApply(tmp523, tmp542)
		return

	}, 3)

	tmp543 := Call(__e, PrimNS2Value(symdef), symget, tmp522)

	_ = tmp543

	tmp544 := MakeNative(func(__e *ControlFlow) {
		V3512 := __e.Get(1)
		_ = V3512
		V3513 := __e.Get(2)
		_ = V3513
		tmp545 := MakeNative(func(__e *ControlFlow) {
			Hash := __e.Get(1)
			_ = Hash
			tmp547 := PrimEqual(Hash, MakeNumber(0))

			if True == tmp547 {
				__e.Return(MakeNumber(1))
				return
			} else {
				__e.Return(Hash)
				return
			}

		}, 1)

		tmp548 := Call(__e, PrimNS2Value(symshen_4hashkey), V3512)

		tmp549 := Call(__e, PrimNS2Value(symshen_4mod), tmp548, V3513)

		__e.TailApply(tmp545, tmp549)
		return

	}, 2)

	tmp550 := Call(__e, PrimNS2Value(symdef), symhash, tmp544)

	_ = tmp550

	tmp551 := MakeNative(func(__e *ControlFlow) {
		V3514 := __e.Get(1)
		_ = V3514
		tmp552 := MakeNative(func(__e *ControlFlow) {
			Ns := __e.Get(1)
			_ = Ns
			__e.TailApply(PrimNS2Value(symshen_4prodbutzero), Ns, MakeNumber(1))
			return
		}, 1)

		tmp553 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(PrimStringToNumber(X))
			return
		}, 1)

		tmp554 := Call(__e, PrimNS2Value(symexplode), V3514)

		tmp555 := Call(__e, PrimNS2Value(symmap), tmp553, tmp554)

		__e.TailApply(tmp552, tmp555)
		return

	}, 1)

	tmp556 := Call(__e, PrimNS2Value(symdef), symshen_4hashkey, tmp551)

	_ = tmp556

	tmp557 := MakeNative(func(__e *ControlFlow) {
		V3515 := __e.Get(1)
		_ = V3515
		V3516 := __e.Get(2)
		_ = V3516
		tmp576 := PrimEqual(Nil, V3515)

		if True == tmp576 {
			__e.Return(V3516)
			return
		} else {
			tmp575 := PrimIsPair(V3515)

			var ifres571 Obj

			if True == tmp575 {
				tmp573 := PrimHead(V3515)

				tmp574 := PrimEqual(MakeNumber(0), tmp573)

				var ifres572 Obj

				if True == tmp574 {
					ifres572 = True

				} else {
					ifres572 = False

				}

				ifres571 = ifres572

			} else {
				ifres571 = False

			}

			if True == ifres571 {
				tmp570 := PrimTail(V3515)

				__e.TailApply(PrimNS2Value(symshen_4prodbutzero), tmp570, V3516)
				return

			} else {
				tmp569 := PrimIsPair(V3515)

				if True == tmp569 {
					tmp568 := PrimGreatThan(V3516, MakeNumber(10000000000))

					if True == tmp568 {
						tmp565 := PrimTail(V3515)

						tmp566 := PrimHead(V3515)

						tmp567 := PrimNumberAdd(V3516, tmp566)

						__e.TailApply(PrimNS2Value(symshen_4prodbutzero), tmp565, tmp567)
						return

					} else {
						tmp562 := PrimTail(V3515)

						tmp563 := PrimHead(V3515)

						tmp564 := PrimNumberMultiply(V3516, tmp563)

						__e.TailApply(PrimNS2Value(symshen_4prodbutzero), tmp562, tmp564)
						return

					}

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4prodbutzero)
					return
				}

			}

		}

	}, 2)

	tmp577 := Call(__e, PrimNS2Value(symdef), symshen_4prodbutzero, tmp557)

	_ = tmp577

	tmp578 := MakeNative(func(__e *ControlFlow) {
		V3517 := __e.Get(1)
		_ = V3517
		V3518 := __e.Get(2)
		_ = V3518
		tmp579 := PrimCons(V3518, Nil)

		tmp580 := Call(__e, PrimNS2Value(symshen_4multiples), V3517, tmp579)

		__e.TailApply(PrimNS2Value(symshen_4modh), V3517, tmp580)
		return

	}, 2)

	tmp581 := Call(__e, PrimNS2Value(symdef), symshen_4mod, tmp578)

	_ = tmp581

	tmp582 := MakeNative(func(__e *ControlFlow) {
		V3523 := __e.Get(1)
		_ = V3523
		V3524 := __e.Get(2)
		_ = V3524
		tmp593 := PrimIsPair(V3524)

		var ifres589 Obj

		if True == tmp593 {
			tmp591 := PrimHead(V3524)

			tmp592 := PrimGreatThan(tmp591, V3523)

			var ifres590 Obj

			if True == tmp592 {
				ifres590 = True

			} else {
				ifres590 = False

			}

			ifres589 = ifres590

		} else {
			ifres589 = False

		}

		if True == ifres589 {
			__e.Return(PrimTail(V3524))
			return
		} else {
			tmp588 := PrimIsPair(V3524)

			if True == tmp588 {
				tmp585 := PrimHead(V3524)

				tmp586 := PrimNumberMultiply(MakeNumber(2), tmp585)

				tmp587 := PrimCons(tmp586, V3524)

				__e.TailApply(PrimNS2Value(symshen_4multiples), V3523, tmp587)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.multiples")))
				return
			}

		}

	}, 2)

	tmp594 := Call(__e, PrimNS2Value(symdef), symshen_4multiples, tmp582)

	_ = tmp594

	tmp595 := MakeNative(func(__e *ControlFlow) {
		V3531 := __e.Get(1)
		_ = V3531
		V3532 := __e.Get(2)
		_ = V3532
		tmp613 := PrimEqual(MakeNumber(0), V3531)

		if True == tmp613 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp612 := PrimEqual(Nil, V3532)

			if True == tmp612 {
				__e.Return(V3531)
				return
			} else {
				tmp611 := PrimIsPair(V3532)

				var ifres607 Obj

				if True == tmp611 {
					tmp609 := PrimHead(V3532)

					tmp610 := PrimGreatThan(tmp609, V3531)

					var ifres608 Obj

					if True == tmp610 {
						ifres608 = True

					} else {
						ifres608 = False

					}

					ifres607 = ifres608

				} else {
					ifres607 = False

				}

				if True == ifres607 {
					tmp605 := PrimTail(V3532)

					tmp606 := Call(__e, PrimNS2Value(symempty_2), tmp605)

					if True == tmp606 {
						__e.Return(V3531)
						return
					} else {
						tmp604 := PrimTail(V3532)

						__e.TailApply(PrimNS2Value(symshen_4modh), V3531, tmp604)
						return

					}

				} else {
					tmp602 := PrimIsPair(V3532)

					if True == tmp602 {
						tmp600 := PrimHead(V3532)

						tmp601 := PrimNumberSubtract(V3531, tmp600)

						__e.TailApply(PrimNS2Value(symshen_4modh), tmp601, V3532)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.modh")))
						return
					}

				}

			}

		}

	}, 2)

	tmp614 := Call(__e, PrimNS2Value(symdef), symshen_4modh, tmp595)

	_ = tmp614

	tmp615 := MakeNative(func(__e *ControlFlow) {
		V3535 := __e.Get(1)
		_ = V3535
		tmp622 := PrimEqual(Nil, V3535)

		if True == tmp622 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp621 := PrimIsPair(V3535)

			if True == tmp621 {
				tmp618 := PrimHead(V3535)

				tmp619 := PrimTail(V3535)

				tmp620 := Call(__e, PrimNS2Value(symsum), tmp619)

				__e.Return(PrimNumberAdd(tmp618, tmp620))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to sum a non-list\n")))
				return
			}

		}

	}, 1)

	tmp623 := Call(__e, PrimNS2Value(symdef), symsum, tmp615)

	_ = tmp623

	tmp624 := MakeNative(func(__e *ControlFlow) {
		V3540 := __e.Get(1)
		_ = V3540
		tmp626 := PrimIsPair(V3540)

		if True == tmp626 {
			__e.Return(PrimHead(V3540))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("head expects a non-empty list\n")))
			return
		}

	}, 1)

	tmp627 := Call(__e, PrimNS2Value(symdef), symhead, tmp624)

	_ = tmp627

	tmp628 := MakeNative(func(__e *ControlFlow) {
		V3545 := __e.Get(1)
		_ = V3545
		tmp630 := PrimIsPair(V3545)

		if True == tmp630 {
			__e.Return(PrimTail(V3545))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list\n")))
			return
		}

	}, 1)

	tmp631 := Call(__e, PrimNS2Value(symdef), symtail, tmp628)

	_ = tmp631

	tmp632 := MakeNative(func(__e *ControlFlow) {
		V3546 := __e.Get(1)
		_ = V3546
		__e.Return(PrimPos(V3546, MakeNumber(0)))
		return
	}, 1)

	tmp633 := Call(__e, PrimNS2Value(symdef), symhdstr, tmp632)

	_ = tmp633

	tmp634 := MakeNative(func(__e *ControlFlow) {
		V3553 := __e.Get(1)
		_ = V3553
		V3554 := __e.Get(2)
		_ = V3554
		tmp645 := PrimEqual(Nil, V3553)

		if True == tmp645 {
			__e.Return(Nil)
			return
		} else {
			tmp644 := PrimIsPair(V3553)

			if True == tmp644 {
				tmp642 := PrimHead(V3553)

				tmp643 := Call(__e, PrimNS2Value(symelement_2), tmp642, V3554)

				if True == tmp643 {
					tmp639 := PrimHead(V3553)

					tmp640 := PrimTail(V3553)

					tmp641 := Call(__e, PrimNS2Value(symintersection), tmp640, V3554)

					__e.Return(PrimCons(tmp639, tmp641))
					return

				} else {
					tmp638 := PrimTail(V3553)

					__e.TailApply(PrimNS2Value(symintersection), tmp638, V3554)
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n")))
				return
			}

		}

	}, 2)

	tmp646 := Call(__e, PrimNS2Value(symdef), symintersection, tmp634)

	_ = tmp646

	tmp647 := MakeNative(func(__e *ControlFlow) {
		V3555 := __e.Get(1)
		_ = V3555
		__e.TailApply(PrimNS2Value(symshen_4reverse_1help), V3555, Nil)
		return
	}, 1)

	tmp648 := Call(__e, PrimNS2Value(symdef), symreverse, tmp647)

	_ = tmp648

	tmp649 := MakeNative(func(__e *ControlFlow) {
		V3560 := __e.Get(1)
		_ = V3560
		V3561 := __e.Get(2)
		_ = V3561
		tmp656 := PrimEqual(Nil, V3560)

		if True == tmp656 {
			__e.Return(V3561)
			return
		} else {
			tmp655 := PrimIsPair(V3560)

			if True == tmp655 {
				tmp652 := PrimTail(V3560)

				tmp653 := PrimHead(V3560)

				tmp654 := PrimCons(tmp653, V3561)

				__e.TailApply(PrimNS2Value(symshen_4reverse_1help), tmp652, tmp654)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to reverse a non-list\n")))
				return
			}

		}

	}, 2)

	tmp657 := Call(__e, PrimNS2Value(symdef), symshen_4reverse_1help, tmp649)

	_ = tmp657

	tmp658 := MakeNative(func(__e *ControlFlow) {
		V3566 := __e.Get(1)
		_ = V3566
		V3567 := __e.Get(2)
		_ = V3567
		tmp669 := PrimEqual(Nil, V3566)

		if True == tmp669 {
			__e.Return(V3567)
			return
		} else {
			tmp668 := PrimIsPair(V3566)

			if True == tmp668 {
				tmp666 := PrimHead(V3566)

				tmp667 := Call(__e, PrimNS2Value(symelement_2), tmp666, V3567)

				if True == tmp667 {
					tmp665 := PrimTail(V3566)

					__e.TailApply(PrimNS2Value(symunion), tmp665, V3567)
					return

				} else {
					tmp662 := PrimHead(V3566)

					tmp663 := PrimTail(V3566)

					tmp664 := Call(__e, PrimNS2Value(symunion), tmp663, V3567)

					__e.Return(PrimCons(tmp662, tmp664))
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to find the union with a non-list\n")))
				return
			}

		}

	}, 2)

	tmp670 := Call(__e, PrimNS2Value(symdef), symunion, tmp658)

	_ = tmp670

	tmp671 := MakeNative(func(__e *ControlFlow) {
		V3568 := __e.Get(1)
		_ = V3568
		tmp672 := MakeNative(func(__e *ControlFlow) {
			Message := __e.Get(1)
			_ = Message
			tmp673 := MakeNative(func(__e *ControlFlow) {
				Y_1or_1N := __e.Get(1)
				_ = Y_1or_1N
				tmp674 := MakeNative(func(__e *ControlFlow) {
					Input := __e.Get(1)
					_ = Input
					tmp680 := PrimEqual(MakeString("y"), Input)

					if True == tmp680 {
						__e.Return(True)
						return
					} else {
						tmp679 := PrimEqual(MakeString("n"), Input)

						if True == tmp679 {
							__e.Return(False)
							return
						} else {
							tmp677 := Call(__e, PrimNS2Value(symstoutput))

							tmp678 := Call(__e, PrimNS2Value(sympr), MakeString("please answer y or n\n"), tmp677)

							_ = tmp678

							__e.TailApply(PrimNS2Value(symy_1or_1n_2), V3568)
							return

						}

					}

				}, 1)

				tmp681 := Call(__e, PrimNS2Value(symstinput))

				tmp682 := Call(__e, PrimNS2Value(symread), tmp681)

				tmp683 := Call(__e, PrimNS2Value(symshen_4app), tmp682, MakeString(""), symshen_4s)

				__e.TailApply(tmp674, tmp683)
				return

			}, 1)

			tmp684 := Call(__e, PrimNS2Value(symstoutput))

			tmp685 := Call(__e, PrimNS2Value(sympr), MakeString(" (y/n) "), tmp684)

			__e.TailApply(tmp673, tmp685)
			return

		}, 1)

		tmp686 := Call(__e, PrimNS2Value(symshen_4proc_1nl), V3568)

		tmp687 := Call(__e, PrimNS2Value(symstoutput))

		tmp688 := Call(__e, PrimNS2Value(sympr), tmp686, tmp687)

		__e.TailApply(tmp672, tmp688)
		return

	}, 1)

	tmp689 := Call(__e, PrimNS2Value(symdef), symy_1or_1n_2, tmp671)

	_ = tmp689

	tmp690 := MakeNative(func(__e *ControlFlow) {
		V3569 := __e.Get(1)
		_ = V3569
		if True == V3569 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}
	}, 1)

	tmp692 := Call(__e, PrimNS2Value(symdef), symnot, tmp690)

	_ = tmp692

	tmp693 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("")))
		return
	}, 0)

	tmp694 := Call(__e, PrimNS2Value(symdef), symabort, tmp693)

	_ = tmp694

	tmp695 := MakeNative(func(__e *ControlFlow) {
		V3575 := __e.Get(1)
		_ = V3575
		V3576 := __e.Get(2)
		_ = V3576
		V3577 := __e.Get(3)
		_ = V3577
		tmp703 := PrimEqual(V3576, V3577)

		if True == tmp703 {
			__e.Return(V3575)
			return
		} else {
			tmp702 := PrimIsPair(V3577)

			if True == tmp702 {
				tmp698 := PrimHead(V3577)

				tmp699 := Call(__e, PrimNS2Value(symsubst), V3575, V3576, tmp698)

				tmp700 := PrimTail(V3577)

				tmp701 := Call(__e, PrimNS2Value(symsubst), V3575, V3576, tmp700)

				__e.Return(PrimCons(tmp699, tmp701))
				return

			} else {
				__e.Return(V3577)
				return
			}

		}

	}, 3)

	tmp704 := Call(__e, PrimNS2Value(symdef), symsubst, tmp695)

	_ = tmp704

	tmp705 := MakeNative(func(__e *ControlFlow) {
		V3578 := __e.Get(1)
		_ = V3578
		tmp706 := Call(__e, PrimNS2Value(symshen_4app), V3578, MakeString(""), symshen_4a)

		__e.TailApply(PrimNS2Value(symshen_4explode_1h), tmp706)
		return

	}, 1)

	tmp707 := Call(__e, PrimNS2Value(symdef), symexplode, tmp705)

	_ = tmp707

	tmp708 := MakeNative(func(__e *ControlFlow) {
		V3581 := __e.Get(1)
		_ = V3581
		tmp715 := PrimEqual(MakeString(""), V3581)

		if True == tmp715 {
			__e.Return(Nil)
			return
		} else {
			tmp714 := Call(__e, PrimNS2Value(symshen_4_7string_2), V3581)

			if True == tmp714 {
				tmp711 := Call(__e, PrimNS2Value(symhdstr), V3581)

				tmp712 := PrimTailString(V3581)

				tmp713 := Call(__e, PrimNS2Value(symshen_4explode_1h), tmp712)

				__e.Return(PrimCons(tmp711, tmp713))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in explode-h")))
				return
			}

		}

	}, 1)

	tmp716 := Call(__e, PrimNS2Value(symdef), symshen_4explode_1h, tmp708)

	_ = tmp716

	tmp717 := MakeNative(func(__e *ControlFlow) {
		V3582 := __e.Get(1)
		_ = V3582
		tmp720 := PrimEqual(V3582, MakeString(""))

		var ifres718 Obj

		if True == tmp720 {
			ifres718 = MakeString("")

		} else {
			tmp719 := Call(__e, PrimNS2Value(symshen_4app), V3582, MakeString("/"), symshen_4a)

			ifres718 = tmp719

		}

		__e.Return(PrimNS3Set(sym_dhome_1directory_d, ifres718))
		return

	}, 1)

	tmp721 := Call(__e, PrimNS2Value(symdef), symcd, tmp717)

	_ = tmp721

	tmp722 := MakeNative(func(__e *ControlFlow) {
		V3583 := __e.Get(1)
		_ = V3583
		V3584 := __e.Get(2)
		_ = V3584
		__e.TailApply(PrimNS2Value(symshen_4map_1h), V3583, V3584, Nil)
		return
	}, 2)

	tmp723 := Call(__e, PrimNS2Value(symdef), symmap, tmp722)

	_ = tmp723

	tmp724 := MakeNative(func(__e *ControlFlow) {
		V3585 := __e.Get(1)
		_ = V3585
		V3586 := __e.Get(2)
		_ = V3586
		V3587 := __e.Get(3)
		_ = V3587
		tmp732 := PrimEqual(Nil, V3586)

		if True == tmp732 {
			__e.TailApply(PrimNS2Value(symreverse), V3587)
			return
		} else {
			tmp731 := PrimIsPair(V3586)

			if True == tmp731 {
				tmp727 := PrimTail(V3586)

				tmp728 := PrimHead(V3586)

				tmp729 := Call(__e, V3585, tmp728)

				tmp730 := PrimCons(tmp729, V3587)

				__e.TailApply(PrimNS2Value(symshen_4map_1h), V3585, tmp727, tmp730)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4map_1h)
				return
			}

		}

	}, 3)

	tmp733 := Call(__e, PrimNS2Value(symdef), symshen_4map_1h, tmp724)

	_ = tmp733

	tmp734 := MakeNative(func(__e *ControlFlow) {
		V3588 := __e.Get(1)
		_ = V3588
		__e.TailApply(PrimNS2Value(symshen_4length_1h), V3588, MakeNumber(0))
		return
	}, 1)

	tmp735 := Call(__e, PrimNS2Value(symdef), symlength, tmp734)

	_ = tmp735

	tmp736 := MakeNative(func(__e *ControlFlow) {
		V3593 := __e.Get(1)
		_ = V3593
		V3594 := __e.Get(2)
		_ = V3594
		tmp740 := PrimEqual(Nil, V3593)

		if True == tmp740 {
			__e.Return(V3594)
			return
		} else {
			tmp738 := PrimTail(V3593)

			tmp739 := PrimNumberAdd(V3594, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4length_1h), tmp738, tmp739)
			return

		}

	}, 2)

	tmp741 := Call(__e, PrimNS2Value(symdef), symshen_4length_1h, tmp736)

	_ = tmp741

	tmp742 := MakeNative(func(__e *ControlFlow) {
		V3600 := __e.Get(1)
		_ = V3600
		V3601 := __e.Get(2)
		_ = V3601
		tmp750 := PrimEqual(V3600, V3601)

		if True == tmp750 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp749 := PrimIsPair(V3601)

			if True == tmp749 {
				tmp745 := PrimHead(V3601)

				tmp746 := Call(__e, PrimNS2Value(symoccurrences), V3600, tmp745)

				tmp747 := PrimTail(V3601)

				tmp748 := Call(__e, PrimNS2Value(symoccurrences), V3600, tmp747)

				__e.Return(PrimNumberAdd(tmp746, tmp748))
				return

			} else {
				__e.Return(MakeNumber(0))
				return
			}

		}

	}, 2)

	tmp751 := Call(__e, PrimNS2Value(symdef), symoccurrences, tmp742)

	_ = tmp751

	tmp752 := MakeNative(func(__e *ControlFlow) {
		V3606 := __e.Get(1)
		_ = V3606
		V3607 := __e.Get(2)
		_ = V3607
		tmp765 := PrimEqual(MakeNumber(1), V3606)

		var ifres762 Obj

		if True == tmp765 {
			tmp764 := PrimIsPair(V3607)

			var ifres763 Obj

			if True == tmp764 {
				ifres763 = True

			} else {
				ifres763 = False

			}

			ifres762 = ifres763

		} else {
			ifres762 = False

		}

		if True == ifres762 {
			__e.Return(PrimHead(V3607))
			return
		} else {
			tmp761 := PrimIsPair(V3607)

			if True == tmp761 {
				tmp759 := PrimNumberSubtract(V3606, MakeNumber(1))

				tmp760 := PrimTail(V3607)

				__e.TailApply(PrimNS2Value(symnth), tmp759, tmp760)
				return

			} else {
				tmp755 := Call(__e, PrimNS2Value(symshen_4app), V3607, MakeString("\n"), symshen_4a)

				tmp756 := PrimStringConcat(MakeString(", "), tmp755)

				tmp757 := Call(__e, PrimNS2Value(symshen_4app), V3606, tmp756, symshen_4a)

				tmp758 := PrimStringConcat(MakeString("nth applied to "), tmp757)

				__e.Return(PrimSimpleError(tmp758))
				return

			}

		}

	}, 2)

	tmp766 := Call(__e, PrimNS2Value(symdef), symnth, tmp752)

	_ = tmp766

	tmp767 := MakeNative(func(__e *ControlFlow) {
		V3608 := __e.Get(1)
		_ = V3608
		tmp774 := PrimIsNumber(V3608)

		if True == tmp774 {
			tmp770 := MakeNative(func(__e *ControlFlow) {
				Abs := __e.Get(1)
				_ = Abs
				tmp771 := Call(__e, PrimNS2Value(symshen_4magless), Abs, MakeNumber(1))

				__e.TailApply(PrimNS2Value(symshen_4integer_1test_2), Abs, tmp771)
				return

			}, 1)

			tmp772 := Call(__e, PrimNS2Value(symshen_4abs), V3608)

			tmp773 := Call(__e, tmp770, tmp772)

			if True == tmp773 {
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

	tmp775 := Call(__e, PrimNS2Value(symdef), syminteger_2, tmp767)

	_ = tmp775

	tmp776 := MakeNative(func(__e *ControlFlow) {
		V3609 := __e.Get(1)
		_ = V3609
		tmp778 := PrimGreatThan(V3609, MakeNumber(0))

		if True == tmp778 {
			__e.Return(V3609)
			return
		} else {
			__e.Return(PrimNumberSubtract(MakeNumber(0), V3609))
			return
		}

	}, 1)

	tmp779 := Call(__e, PrimNS2Value(symdef), symshen_4abs, tmp776)

	_ = tmp779

	tmp780 := MakeNative(func(__e *ControlFlow) {
		V3610 := __e.Get(1)
		_ = V3610
		V3611 := __e.Get(2)
		_ = V3611
		tmp781 := MakeNative(func(__e *ControlFlow) {
			Nx2 := __e.Get(1)
			_ = Nx2
			tmp783 := PrimGreatThan(Nx2, V3610)

			if True == tmp783 {
				__e.Return(V3611)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4magless), V3610, Nx2)
				return
			}

		}, 1)

		tmp784 := PrimNumberMultiply(V3611, MakeNumber(2))

		__e.TailApply(tmp781, tmp784)
		return

	}, 2)

	tmp785 := Call(__e, PrimNS2Value(symdef), symshen_4magless, tmp780)

	_ = tmp785

	tmp786 := MakeNative(func(__e *ControlFlow) {
		V3615 := __e.Get(1)
		_ = V3615
		V3616 := __e.Get(2)
		_ = V3616
		tmp794 := PrimEqual(MakeNumber(0), V3615)

		if True == tmp794 {
			__e.Return(True)
			return
		} else {
			tmp793 := PrimGreatThan(MakeNumber(1), V3615)

			if True == tmp793 {
				__e.Return(False)
				return
			} else {
				tmp789 := MakeNative(func(__e *ControlFlow) {
					Abs_1N := __e.Get(1)
					_ = Abs_1N
					tmp791 := PrimGreatThan(MakeNumber(0), Abs_1N)

					if True == tmp791 {
						__e.Return(PrimIsInteger(V3615))
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4integer_1test_2), Abs_1N, V3616)
						return
					}

				}, 1)

				tmp792 := PrimNumberSubtract(V3615, V3616)

				__e.TailApply(tmp789, tmp792)
				return

			}

		}

	}, 2)

	tmp795 := Call(__e, PrimNS2Value(symdef), symshen_4integer_1test_2, tmp786)

	_ = tmp795

	tmp796 := MakeNative(func(__e *ControlFlow) {
		V3623 := __e.Get(1)
		_ = V3623
		V3624 := __e.Get(2)
		_ = V3624
		tmp804 := PrimEqual(Nil, V3624)

		if True == tmp804 {
			__e.Return(Nil)
			return
		} else {
			tmp803 := PrimIsPair(V3624)

			if True == tmp803 {
				tmp799 := PrimHead(V3624)

				tmp800 := Call(__e, V3623, tmp799)

				tmp801 := PrimTail(V3624)

				tmp802 := Call(__e, PrimNS2Value(symmapcan), V3623, tmp801)

				__e.TailApply(PrimNS2Value(symappend), tmp800, tmp802)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to mapcan over a non-list\n")))
				return
			}

		}

	}, 2)

	tmp805 := Call(__e, PrimNS2Value(symdef), symmapcan, tmp796)

	_ = tmp805

	tmp806 := MakeNative(func(__e *ControlFlow) {
		V3630 := __e.Get(1)
		_ = V3630
		V3631 := __e.Get(2)
		_ = V3631
		tmp808 := PrimEqual(V3630, V3631)

		if True == tmp808 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp809 := Call(__e, PrimNS2Value(symdef), sym_a_a, tmp806)

	_ = tmp809

	tmp810 := MakeNative(func(__e *ControlFlow) {
		V3632 := __e.Get(1)
		_ = V3632
		tmp820 := PrimIsSymbol(V3632)

		if True == tmp820 {
			tmp813 := MakeNative(func(__e *ControlFlow) {
				Val := __e.Get(1)
				_ = Val
				tmp815 := PrimEqual(Val, symshen_4this_1symbol_1is_1unbound)

				if True == tmp815 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}, 1)

			tmp816 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimNS3Value(V3632))
				return
			}, 0)

			tmp817 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4this_1symbol_1is_1unbound)
				return
			}, 1)

			tmp818 := Call(__e, PrimNS1Value(symtry_1catch), tmp816, tmp817)

			tmp819 := Call(__e, tmp813, tmp818)

			if True == tmp819 {
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

	tmp821 := Call(__e, PrimNS2Value(symdef), symbound_2, tmp810)

	_ = tmp821

	tmp822 := MakeNative(func(__e *ControlFlow) {
		V3633 := __e.Get(1)
		_ = V3633
		tmp828 := PrimEqual(MakeString(""), V3633)

		if True == tmp828 {
			__e.Return(Nil)
			return
		} else {
			tmp824 := PrimPos(V3633, MakeNumber(0))

			tmp825 := PrimStringToNumber(tmp824)

			tmp826 := PrimTailString(V3633)

			tmp827 := Call(__e, PrimNS2Value(symshen_4string_1_6bytes), tmp826)

			__e.Return(PrimCons(tmp825, tmp827))
			return

		}

	}, 1)

	tmp829 := Call(__e, PrimNS2Value(symdef), symshen_4string_1_6bytes, tmp822)

	_ = tmp829

	tmp830 := MakeNative(func(__e *ControlFlow) {
		V3634 := __e.Get(1)
		_ = V3634
		__e.Return(PrimNS3Set(symshen_4_dmaxinferences_d, V3634))
		return
	}, 1)

	tmp831 := Call(__e, PrimNS2Value(symdef), symmaxinferences, tmp830)

	_ = tmp831

	tmp832 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dinfs_d))
		return
	}, 0)

	tmp833 := Call(__e, PrimNS2Value(symdef), syminferences, tmp832)

	_ = tmp833

	tmp834 := MakeNative(func(__e *ControlFlow) {
		V3635 := __e.Get(1)
		_ = V3635
		__e.Return(V3635)
		return
	}, 1)

	tmp835 := Call(__e, PrimNS2Value(symdef), symprotect, tmp834)

	_ = tmp835

	tmp836 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dstoutput_d))
		return
	}, 0)

	tmp837 := Call(__e, PrimNS2Value(symdef), symstoutput, tmp836)

	_ = tmp837

	tmp838 := MakeNative(func(__e *ControlFlow) {
		V3636 := __e.Get(1)
		_ = V3636
		tmp839 := MakeNative(func(__e *ControlFlow) {
			Symbol := __e.Get(1)
			_ = Symbol
			tmp843 := PrimIsSymbol(Symbol)

			if True == tmp843 {
				__e.Return(Symbol)
				return
			} else {
				tmp841 := Call(__e, PrimNS2Value(symshen_4app), V3636, MakeString(" to a symbol"), symshen_4s)

				tmp842 := PrimStringConcat(MakeString("cannot intern "), tmp841)

				__e.Return(PrimSimpleError(tmp842))
				return

			}

		}, 1)

		tmp844 := PrimIntern(V3636)

		__e.TailApply(tmp839, tmp844)
		return

	}, 1)

	tmp845 := Call(__e, PrimNS2Value(symdef), symstring_1_6symbol, tmp838)

	_ = tmp845

	tmp846 := MakeNative(func(__e *ControlFlow) {
		V3639 := __e.Get(1)
		_ = V3639
		tmp850 := PrimEqual(sym_7, V3639)

		if True == tmp850 {
			__e.Return(PrimNS3Set(symshen_4_doptimise_d, True))
			return
		} else {
			tmp849 := PrimEqual(sym_1, V3639)

			if True == tmp849 {
				__e.Return(PrimNS3Set(symshen_4_doptimise_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp851 := Call(__e, PrimNS2Value(symdef), symoptimise, tmp846)

	_ = tmp851

	tmp852 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dos_d))
		return
	}, 0)

	tmp853 := Call(__e, PrimNS2Value(symdef), symos, tmp852)

	_ = tmp853

	tmp854 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dlanguage_d))
		return
	}, 0)

	tmp855 := Call(__e, PrimNS2Value(symdef), symlanguage, tmp854)

	_ = tmp855

	tmp856 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dversion_d))
		return
	}, 0)

	tmp857 := Call(__e, PrimNS2Value(symdef), symversion, tmp856)

	_ = tmp857

	tmp858 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dport_d))
		return
	}, 0)

	tmp859 := Call(__e, PrimNS2Value(symdef), symport, tmp858)

	_ = tmp859

	tmp860 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dporters_d))
		return
	}, 0)

	tmp861 := Call(__e, PrimNS2Value(symdef), symporters, tmp860)

	_ = tmp861

	tmp862 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dimplementation_d))
		return
	}, 0)

	tmp863 := Call(__e, PrimNS2Value(symdef), symimplementation, tmp862)

	_ = tmp863

	tmp864 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_drelease_d))
		return
	}, 0)

	tmp865 := Call(__e, PrimNS2Value(symdef), symrelease, tmp864)

	_ = tmp865

	tmp866 := MakeNative(func(__e *ControlFlow) {
		V3640 := __e.Get(1)
		_ = V3640
		tmp871 := PrimEqual(symnull, V3640)

		if True == tmp871 {
			__e.Return(True)
			return
		} else {
			tmp868 := MakeNative(func(__e *ControlFlow) {
				tmp869 := Call(__e, PrimNS2Value(symexternal), V3640)

				_ = tmp869

				__e.Return(True)
				return

			}, 0)

			tmp870 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp868, tmp870)
			return

		}

	}, 1)

	tmp872 := Call(__e, PrimNS2Value(symdef), sympackage_2, tmp866)

	_ = tmp872

	tmp873 := MakeNative(func(__e *ControlFlow) {
		V3641 := __e.Get(1)
		_ = V3641
		tmp874 := MakeNative(func(__e *ControlFlow) {
			Assoc := __e.Get(1)
			_ = Assoc
			tmp877 := Call(__e, PrimNS2Value(symempty_2), Assoc)

			if True == tmp877 {
				tmp876 := Call(__e, PrimNS2Value(symshen_4app), V3641, MakeString(" has no lambda expansion\n"), symshen_4a)

				__e.Return(PrimSimpleError(tmp876))
				return

			} else {
				__e.Return(PrimTail(Assoc))
				return
			}

		}, 1)

		tmp878 := PrimNS3Value(symshen_4_dlambdatable_d)

		tmp879 := Call(__e, PrimNS2Value(symassoc), V3641, tmp878)

		__e.TailApply(tmp874, tmp879)
		return

	}, 1)

	tmp880 := Call(__e, PrimNS2Value(symdef), symfn, tmp873)

	_ = tmp880

	tmp881 := MakeNative(func(__e *ControlFlow) {
		__e.Return(symshen_4fail_b)
		return
	}, 0)

	tmp882 := Call(__e, PrimNS2Value(symdef), symfail, tmp881)

	_ = tmp882

	tmp883 := MakeNative(func(__e *ControlFlow) {
		V3644 := __e.Get(1)
		_ = V3644
		tmp887 := PrimEqual(sym_7, V3644)

		if True == tmp887 {
			__e.Return(PrimNS3Set(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
			return
		} else {
			tmp886 := PrimEqual(sym_1, V3644)

			if True == tmp886 {
				__e.Return(PrimNS3Set(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
				return
			}

		}

	}, 1)

	tmp888 := Call(__e, PrimNS2Value(symdef), symenable_1type_1theory, tmp883)

	_ = tmp888

	tmp889 := MakeNative(func(__e *ControlFlow) {
		V3647 := __e.Get(1)
		_ = V3647
		tmp893 := PrimEqual(sym_7, V3647)

		if True == tmp893 {
			__e.Return(PrimNS3Set(symshen_4_dtc_d, True))
			return
		} else {
			tmp892 := PrimEqual(sym_1, V3647)

			if True == tmp892 {
				__e.Return(PrimNS3Set(symshen_4_dtc_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
				return
			}

		}

	}, 1)

	tmp894 := Call(__e, PrimNS2Value(symdef), symtc, tmp889)

	_ = tmp894

	tmp895 := MakeNative(func(__e *ControlFlow) {
		V3648 := __e.Get(1)
		_ = V3648
		tmp896 := PrimNS3Value(symshen_4_dsigf_d)

		tmp897 := Call(__e, PrimNS2Value(symshen_4unassoc), V3648, tmp896)

		_ = tmp897

		__e.Return(V3648)
		return

	}, 1)

	tmp898 := Call(__e, PrimNS2Value(symdef), symdestroy, tmp895)

	_ = tmp898

	tmp899 := MakeNative(func(__e *ControlFlow) {
		V3649 := __e.Get(1)
		_ = V3649
		V3650 := __e.Get(2)
		_ = V3650
		tmp900 := MakeNative(func(__e *ControlFlow) {
			Assoc := __e.Get(1)
			_ = Assoc
			tmp901 := MakeNative(func(__e *ControlFlow) {
				Remove := __e.Get(1)
				_ = Remove
				__e.Return(PrimNS3Set(symshen_4_dsigf_d, Remove))
				return
			}, 1)

			tmp902 := Call(__e, PrimNS2Value(symremove), Assoc, V3650)

			__e.TailApply(tmp901, tmp902)
			return

		}, 1)

		tmp903 := Call(__e, PrimNS2Value(symassoc), V3649, V3650)

		__e.TailApply(tmp900, tmp903)
		return

	}, 2)

	tmp904 := Call(__e, PrimNS2Value(symdef), symshen_4unassoc, tmp899)

	_ = tmp904

	tmp905 := MakeNative(func(__e *ControlFlow) {
		V3651 := __e.Get(1)
		_ = V3651
		tmp909 := Call(__e, PrimNS2Value(sympackage_2), V3651)

		if True == tmp909 {
			__e.Return(PrimNS3Set(symshen_4_dpackage_d, V3651))
			return
		} else {
			tmp907 := Call(__e, PrimNS2Value(symshen_4app), V3651, MakeString(" does not exist\n"), symshen_4a)

			tmp908 := PrimStringConcat(MakeString("package "), tmp907)

			__e.Return(PrimSimpleError(tmp908))
			return

		}

	}, 1)

	tmp910 := Call(__e, PrimNS2Value(symdef), symin_1package, tmp905)

	_ = tmp910

	tmp911 := MakeNative(func(__e *ControlFlow) {
		V3652 := __e.Get(1)
		_ = V3652
		V3653 := __e.Get(2)
		_ = V3653
		tmp912 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp913 := MakeNative(func(__e *ControlFlow) {
				String := __e.Get(1)
				_ = String
				tmp914 := MakeNative(func(__e *ControlFlow) {
					Write := __e.Get(1)
					_ = Write
					tmp915 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.Return(V3653)
						return
					}, 1)

					tmp916 := PrimCloseStream(Stream)

					__e.TailApply(tmp915, tmp916)
					return

				}, 1)

				tmp917 := Call(__e, PrimNS2Value(sympr), String, Stream)

				__e.TailApply(tmp914, tmp917)
				return

			}, 1)

			tmp921 := PrimIsString(V3653)

			var ifres918 Obj

			if True == tmp921 {
				tmp920 := Call(__e, PrimNS2Value(symshen_4app), V3653, MakeString("\n\n"), symshen_4a)

				ifres918 = tmp920

			} else {
				tmp919 := Call(__e, PrimNS2Value(symshen_4app), V3653, MakeString("\n\n"), symshen_4s)

				ifres918 = tmp919

			}

			__e.TailApply(tmp913, ifres918)
			return

		}, 1)

		tmp922 := PrimOpenStream(V3652, symout)

		__e.TailApply(tmp912, tmp922)
		return

	}, 2)

	tmp923 := Call(__e, PrimNS2Value(symdef), symwrite_1to_1file, tmp911)

	_ = tmp923

	tmp924 := MakeNative(func(__e *ControlFlow) {
		tmp925 := Call(__e, PrimNS2Value(symgensym), symshen_4t)

		__e.TailApply(PrimNS2Value(symshen_4freshterm), tmp925)
		return

	}, 0)

	tmp926 := Call(__e, PrimNS2Value(symdef), symfresh, tmp924)

	_ = tmp926

	tmp927 := MakeNative(func(__e *ControlFlow) {
		V3654 := __e.Get(1)
		_ = V3654
		V3655 := __e.Get(2)
		_ = V3655
		tmp928 := MakeNative(func(__e *ControlFlow) {
			AssertArity := __e.Get(1)
			_ = AssertArity
			tmp929 := MakeNative(func(__e *ControlFlow) {
				LambdaEntry := __e.Get(1)
				_ = LambdaEntry
				tmp930 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V3654)
					return
				}, 1)

				tmp931 := PrimNS3Value(symshen_4_dlambdatable_d)

				tmp932 := PrimCons(LambdaEntry, tmp931)

				tmp933 := PrimNS3Set(symshen_4_dlambdatable_d, tmp932)

				__e.TailApply(tmp930, tmp933)
				return

			}, 1)

			tmp934 := Call(__e, PrimNS2Value(symshen_4lambda_1entry), V3654)

			__e.TailApply(tmp929, tmp934)
			return

		}, 1)

		tmp935 := PrimNS3Value(sym_dproperty_1vector_d)

		tmp936 := Call(__e, PrimNS2Value(symput), V3654, symarity, V3655, tmp935)

		__e.TailApply(tmp928, tmp936)
		return

	}, 2)

	tmp937 := Call(__e, PrimNS2Value(symdef), symupdate_1lambda_1table, tmp927)

	_ = tmp937

	tmp938 := MakeNative(func(__e *ControlFlow) {
		V3658 := __e.Get(1)
		_ = V3658
		V3659 := __e.Get(2)
		_ = V3659
		tmp962 := PrimEqual(MakeNumber(0), V3659)

		if True == tmp962 {
			tmp956 := PrimNS3Value(symshen_4_dspecial_d)

			tmp957 := Call(__e, PrimNS2Value(symremove), V3658, tmp956)

			tmp958 := PrimNS3Set(symshen_4_dspecial_d, tmp957)

			_ = tmp958

			tmp959 := PrimNS3Value(symshen_4_dextraspecial_d)

			tmp960 := Call(__e, PrimNS2Value(symremove), V3658, tmp959)

			tmp961 := PrimNS3Set(symshen_4_dextraspecial_d, tmp960)

			_ = tmp961

			__e.Return(V3658)
			return

		} else {
			tmp955 := PrimEqual(MakeNumber(1), V3659)

			if True == tmp955 {
				tmp949 := PrimNS3Value(symshen_4_dspecial_d)

				tmp950 := Call(__e, PrimNS2Value(symadjoin), V3658, tmp949)

				tmp951 := PrimNS3Set(symshen_4_dspecial_d, tmp950)

				_ = tmp951

				tmp952 := PrimNS3Value(symshen_4_dextraspecial_d)

				tmp953 := Call(__e, PrimNS2Value(symremove), V3658, tmp952)

				tmp954 := PrimNS3Set(symshen_4_dextraspecial_d, tmp953)

				_ = tmp954

				__e.Return(V3658)
				return

			} else {
				tmp948 := PrimEqual(MakeNumber(2), V3659)

				if True == tmp948 {
					tmp942 := PrimNS3Value(symshen_4_dspecial_d)

					tmp943 := Call(__e, PrimNS2Value(symremove), V3658, tmp942)

					tmp944 := PrimNS3Set(symshen_4_dspecial_d, tmp943)

					_ = tmp944

					tmp945 := PrimNS3Value(symshen_4_dextraspecial_d)

					tmp946 := Call(__e, PrimNS2Value(symadjoin), V3658, tmp945)

					tmp947 := PrimNS3Set(symshen_4_dextraspecial_d, tmp946)

					_ = tmp947

					__e.Return(V3658)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n")))
					return
				}

			}

		}

	}, 2)

	__e.TailApply(PrimNS2Value(symdef), symspecialise, tmp938)
	return

}, 0)
