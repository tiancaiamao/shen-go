package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
	tmp75 := MakeNative(func(__e *ControlFlow) {
		V3743 := __e.Get(1)
		_ = V3743
		__e.TailApply(V3743)
		return
	}, 1)

	tmp76 := Call(__e, ns2_1set, symthaw, tmp75)

	_ = tmp76

	tmp77 := MakeNative(func(__e *ControlFlow) {
		V3744 := __e.Get(1)
		_ = V3744
		tmp78 := Call(__e, PrimFunc(symmacroexpand), V3744)

		tmp79 := Call(__e, PrimFunc(symshen_4find_1types), V3744)

		tmp80 := Call(__e, PrimFunc(symshen_4process_1applications), tmp78, tmp79)

		tmp81 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp80)

		__e.TailApply(PrimFunc(symeval_1kl), tmp81)
		return

	}, 1)

	tmp82 := Call(__e, ns2_1set, symeval, tmp77)

	_ = tmp82

	tmp83 := MakeNative(func(__e *ControlFlow) {
		V3745 := __e.Get(1)
		_ = V3745
		tmp90 := PrimEqual(symnull, V3745)

		if True == tmp90 {
			__e.Return(Nil)
			return
		} else {
			tmp84 := MakeNative(func(__e *ControlFlow) {
				tmp85 := PrimValue(sym_dproperty_1vector_d)

				__e.TailApply(PrimFunc(symget), V3745, symshen_4external_1symbols, tmp85)
				return

			}, 0)

			tmp86 := MakeNative(func(__e *ControlFlow) {
				Z3746 := __e.Get(1)
				_ = Z3746
				tmp87 := Call(__e, PrimFunc(symshen_4app), V3745, MakeString(" does not exist.\n;"), symshen_4a)

				tmp88 := PrimStringConcat(MakeString("package "), tmp87)

				__e.Return(PrimSimpleError(tmp88))
				return

			}, 1)

			__e.TailApply(try_1catch, tmp84, tmp86)
			return

		}

	}, 1)

	tmp91 := Call(__e, ns2_1set, symexternal, tmp83)

	_ = tmp91

	tmp92 := MakeNative(func(__e *ControlFlow) {
		V3747 := __e.Get(1)
		_ = V3747
		tmp99 := PrimEqual(symnull, V3747)

		if True == tmp99 {
			__e.Return(Nil)
			return
		} else {
			tmp93 := MakeNative(func(__e *ControlFlow) {
				tmp94 := PrimValue(sym_dproperty_1vector_d)

				__e.TailApply(PrimFunc(symget), V3747, symshen_4internal_1symbols, tmp94)
				return

			}, 0)

			tmp95 := MakeNative(func(__e *ControlFlow) {
				Z3748 := __e.Get(1)
				_ = Z3748
				tmp96 := Call(__e, PrimFunc(symshen_4app), V3747, MakeString(" does not exist.\n;"), symshen_4a)

				tmp97 := PrimStringConcat(MakeString("package "), tmp96)

				__e.Return(PrimSimpleError(tmp97))
				return

			}, 1)

			__e.TailApply(try_1catch, tmp93, tmp95)
			return

		}

	}, 1)

	tmp100 := Call(__e, ns2_1set, syminternal, tmp92)

	_ = tmp100

	tmp101 := MakeNative(func(__e *ControlFlow) {
		V3749 := __e.Get(1)
		_ = V3749
		V3750 := __e.Get(2)
		_ = V3750
		tmp103 := Call(__e, V3749, V3750)

		if True == tmp103 {
			__e.TailApply(PrimFunc(symfail))
			return
		} else {
			__e.Return(V3750)
			return
		}

	}, 2)

	tmp104 := Call(__e, ns2_1set, symfail_1if, tmp101)

	_ = tmp104

	tmp105 := MakeNative(func(__e *ControlFlow) {
		V3751 := __e.Get(1)
		_ = V3751
		V3752 := __e.Get(2)
		_ = V3752
		__e.Return(PrimStringConcat(V3751, V3752))
		return
	}, 2)

	tmp106 := Call(__e, ns2_1set, sym_8s, tmp105)

	_ = tmp106

	tmp107 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dtc_d))
		return
	}, 0)

	tmp108 := Call(__e, ns2_1set, symtc_2, tmp107)

	_ = tmp108

	tmp109 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_doccurs_d))
		return
	}, 0)

	tmp110 := Call(__e, ns2_1set, symoccurs_2, tmp109)

	_ = tmp110

	tmp111 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dfactorise_2_d))
		return
	}, 0)

	tmp112 := Call(__e, ns2_1set, symfactorise_2, tmp111)

	_ = tmp112

	tmp113 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dtracking_d))
		return
	}, 0)

	tmp114 := Call(__e, ns2_1set, symtracked, tmp113)

	_ = tmp114

	tmp115 := MakeNative(func(__e *ControlFlow) {
		V3753 := __e.Get(1)
		_ = V3753
		tmp116 := MakeNative(func(__e *ControlFlow) {
			tmp117 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symget), V3753, symshen_4source, tmp117)
			return

		}, 0)

		tmp118 := MakeNative(func(__e *ControlFlow) {
			Z3754 := __e.Get(1)
			_ = Z3754
			tmp119 := Call(__e, PrimFunc(symshen_4app), V3753, MakeString(" not found.\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp119))
			return

		}, 1)

		__e.TailApply(try_1catch, tmp116, tmp118)
		return

	}, 1)

	tmp120 := Call(__e, ns2_1set, symps, tmp115)

	_ = tmp120

	tmp121 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dstinput_d))
		return
	}, 0)

	tmp122 := Call(__e, ns2_1set, symstinput, tmp121)

	_ = tmp122

	tmp123 := MakeNative(func(__e *ControlFlow) {
		V3755 := __e.Get(1)
		_ = V3755
		tmp124 := MakeNative(func(__e *ControlFlow) {
			W3756 := __e.Get(1)
			_ = W3756
			tmp125 := MakeNative(func(__e *ControlFlow) {
				W3757 := __e.Get(1)
				_ = W3757
				tmp126 := MakeNative(func(__e *ControlFlow) {
					W3758 := __e.Get(1)
					_ = W3758
					__e.Return(W3758)
					return
				}, 1)

				tmp130 := PrimEqual(V3755, MakeNumber(0))

				var ifres127 Obj

				if True == tmp130 {
					ifres127 = W3757

				} else {
					tmp128 := Call(__e, PrimFunc(symfail))

					tmp129 := Call(__e, PrimFunc(symshen_4fillvector), W3757, MakeNumber(1), V3755, tmp128)

					ifres127 = tmp129

				}

				__e.TailApply(tmp126, ifres127)
				return

			}, 1)

			tmp131 := PrimVectorSet(W3756, MakeNumber(0), V3755)

			__e.TailApply(tmp125, tmp131)
			return

		}, 1)

		tmp132 := PrimNumberAdd(V3755, MakeNumber(1))

		tmp133 := PrimAbsvector(tmp132)

		__e.TailApply(tmp124, tmp133)
		return

	}, 1)

	tmp134 := Call(__e, ns2_1set, symvector, tmp123)

	_ = tmp134

	tmp135 := MakeNative(func(__e *ControlFlow) {
		V3760 := __e.Get(1)
		_ = V3760
		V3761 := __e.Get(2)
		_ = V3761
		V3762 := __e.Get(3)
		_ = V3762
		V3763 := __e.Get(4)
		_ = V3763
		tmp139 := PrimEqual(V3761, V3762)

		if True == tmp139 {
			__e.Return(PrimVectorSet(V3760, V3762, V3763))
			return
		} else {
			tmp136 := PrimVectorSet(V3760, V3761, V3763)

			tmp137 := PrimNumberAdd(MakeNumber(1), V3761)

			__e.TailApply(PrimFunc(symshen_4fillvector), tmp136, tmp137, V3762, V3763)
			return

		}

	}, 4)

	tmp140 := Call(__e, ns2_1set, symshen_4fillvector, tmp135)

	_ = tmp140

	tmp141 := MakeNative(func(__e *ControlFlow) {
		V3764 := __e.Get(1)
		_ = V3764
		tmp153 := PrimIsVector(V3764)

		if True == tmp153 {
			tmp143 := MakeNative(func(__e *ControlFlow) {
				W3765 := __e.Get(1)
				_ = W3765
				tmp147 := PrimIsNumber(W3765)

				if True == tmp147 {
					tmp145 := PrimGreatEqual(W3765, MakeNumber(0))

					if True == tmp145 {
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

			tmp148 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V3764, MakeNumber(0)))
				return
			}, 0)

			tmp149 := MakeNative(func(__e *ControlFlow) {
				Z3766 := __e.Get(1)
				_ = Z3766
				__e.Return(MakeNumber(-1))
				return
			}, 1)

			tmp150 := Call(__e, try_1catch, tmp148, tmp149)

			tmp151 := Call(__e, tmp143, tmp150)

			if True == tmp151 {
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

	tmp154 := Call(__e, ns2_1set, symvector_2, tmp141)

	_ = tmp154

	tmp155 := MakeNative(func(__e *ControlFlow) {
		V3767 := __e.Get(1)
		_ = V3767
		V3768 := __e.Get(2)
		_ = V3768
		V3769 := __e.Get(3)
		_ = V3769
		tmp157 := PrimEqual(V3768, MakeNumber(0))

		if True == tmp157 {
			__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
			return
		} else {
			__e.Return(PrimVectorSet(V3767, V3768, V3769))
			return
		}

	}, 3)

	tmp158 := Call(__e, ns2_1set, symvector_1_6, tmp155)

	_ = tmp158

	tmp159 := MakeNative(func(__e *ControlFlow) {
		V3770 := __e.Get(1)
		_ = V3770
		V3771 := __e.Get(2)
		_ = V3771
		tmp166 := PrimEqual(V3771, MakeNumber(0))

		if True == tmp166 {
			__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
			return
		} else {
			tmp160 := MakeNative(func(__e *ControlFlow) {
				W3772 := __e.Get(1)
				_ = W3772
				tmp162 := Call(__e, PrimFunc(symfail))

				tmp163 := PrimEqual(W3772, tmp162)

				if True == tmp163 {
					__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
					return
				} else {
					__e.Return(W3772)
					return
				}

			}, 1)

			tmp164 := PrimVectorGet(V3770, V3771)

			__e.TailApply(tmp160, tmp164)
			return

		}

	}, 2)

	tmp167 := Call(__e, ns2_1set, sym_5_1vector, tmp159)

	_ = tmp167

	tmp168 := MakeNative(func(__e *ControlFlow) {
		V3773 := __e.Get(1)
		_ = V3773
		tmp172 := PrimIsInteger(V3773)

		if True == tmp172 {
			tmp170 := PrimGreatEqual(V3773, MakeNumber(0))

			if True == tmp170 {
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

	tmp173 := Call(__e, ns2_1set, symshen_4posint_2, tmp168)

	_ = tmp173

	tmp174 := MakeNative(func(__e *ControlFlow) {
		V3774 := __e.Get(1)
		_ = V3774
		__e.Return(PrimVectorGet(V3774, MakeNumber(0)))
		return
	}, 1)

	tmp175 := Call(__e, ns2_1set, symlimit, tmp174)

	_ = tmp175

	tmp176 := MakeNative(func(__e *ControlFlow) {
		V3775 := __e.Get(1)
		_ = V3775
		tmp207 := Call(__e, PrimFunc(symboolean_2), V3775)

		var ifres192 Obj

		if True == tmp207 {
			ifres192 = True

		} else {
			tmp206 := PrimIsNumber(V3775)

			var ifres194 Obj

			if True == tmp206 {
				ifres194 = True

			} else {
				tmp205 := PrimIsString(V3775)

				var ifres196 Obj

				if True == tmp205 {
					ifres196 = True

				} else {
					tmp204 := PrimIsPair(V3775)

					var ifres198 Obj

					if True == tmp204 {
						ifres198 = True

					} else {
						tmp203 := Call(__e, PrimFunc(symempty_2), V3775)

						var ifres200 Obj

						if True == tmp203 {
							ifres200 = True

						} else {
							tmp202 := Call(__e, PrimFunc(symvector_2), V3775)

							var ifres201 Obj

							if True == tmp202 {
								ifres201 = True

							} else {
								ifres201 = False

							}

							ifres200 = ifres201

						}

						var ifres199 Obj

						if True == ifres200 {
							ifres199 = True

						} else {
							ifres199 = False

						}

						ifres198 = ifres199

					}

					var ifres197 Obj

					if True == ifres198 {
						ifres197 = True

					} else {
						ifres197 = False

					}

					ifres196 = ifres197

				}

				var ifres195 Obj

				if True == ifres196 {
					ifres195 = True

				} else {
					ifres195 = False

				}

				ifres194 = ifres195

			}

			var ifres193 Obj

			if True == ifres194 {
				ifres193 = True

			} else {
				ifres193 = False

			}

			ifres192 = ifres193

		}

		if True == ifres192 {
			__e.Return(False)
			return
		} else {
			tmp182 := PrimIntern(MakeString(":"))

			tmp183 := PrimIntern(MakeString(";"))

			tmp184 := PrimIntern(MakeString(","))

			tmp185 := PrimCons(tmp184, Nil)

			tmp186 := PrimCons(tmp183, tmp185)

			tmp187 := PrimCons(tmp182, tmp186)

			tmp188 := PrimCons(sym_j, tmp187)

			tmp189 := PrimCons(sym_i, tmp188)

			tmp190 := Call(__e, PrimFunc(symelement_2), V3775, tmp189)

			if True == tmp190 {
				__e.Return(True)
				return
			} else {
				tmp177 := MakeNative(func(__e *ControlFlow) {
					tmp178 := MakeNative(func(__e *ControlFlow) {
						W3776 := __e.Get(1)
						_ = W3776
						__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), W3776)
						return
					}, 1)

					tmp179 := PrimStr(V3775)

					__e.TailApply(tmp178, tmp179)
					return

				}, 0)

				tmp180 := MakeNative(func(__e *ControlFlow) {
					Z3777 := __e.Get(1)
					_ = Z3777
					__e.Return(False)
					return
				}, 1)

				__e.TailApply(try_1catch, tmp177, tmp180)
				return

			}

		}

	}, 1)

	tmp208 := Call(__e, ns2_1set, symsymbol_2, tmp176)

	_ = tmp208

	tmp209 := MakeNative(func(__e *ControlFlow) {
		V3780 := __e.Get(1)
		_ = V3780
		tmp218 := Call(__e, PrimFunc(symshen_4_7string_2), V3780)

		if True == tmp218 {
			tmp214 := Call(__e, PrimFunc(symhdstr), V3780)

			tmp215 := PrimStringToNumber(tmp214)

			tmp216 := Call(__e, PrimFunc(symshen_4alpha_2), tmp215)

			if True == tmp216 {
				tmp211 := PrimTailString(V3780)

				tmp212 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp211)

				if True == tmp212 {
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

	tmp219 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp209)

	_ = tmp219

	tmp220 := MakeNative(func(__e *ControlFlow) {
		V3783 := __e.Get(1)
		_ = V3783
		tmp235 := PrimEqual(MakeString(""), V3783)

		if True == tmp235 {
			__e.Return(True)
			return
		} else {
			tmp233 := Call(__e, PrimFunc(symshen_4_7string_2), V3783)

			if True == tmp233 {
				tmp221 := MakeNative(func(__e *ControlFlow) {
					W3784 := __e.Get(1)
					_ = W3784
					tmp229 := Call(__e, PrimFunc(symshen_4alpha_2), W3784)

					var ifres226 Obj

					if True == tmp229 {
						ifres226 = True

					} else {
						tmp228 := Call(__e, PrimFunc(symshen_4digit_2), W3784)

						var ifres227 Obj

						if True == tmp228 {
							ifres227 = True

						} else {
							ifres227 = False

						}

						ifres226 = ifres227

					}

					if True == ifres226 {
						tmp223 := PrimTailString(V3783)

						tmp224 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp223)

						if True == tmp224 {
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

				tmp230 := Call(__e, PrimFunc(symhdstr), V3783)

				tmp231 := PrimStringToNumber(tmp230)

				__e.TailApply(tmp221, tmp231)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.alphanums?")))
				return
			}

		}

	}, 1)

	tmp236 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp220)

	_ = tmp236

	tmp237 := MakeNative(func(__e *ControlFlow) {
		V3785 := __e.Get(1)
		_ = V3785
		tmp249 := Call(__e, PrimFunc(symboolean_2), V3785)

		var ifres243 Obj

		if True == tmp249 {
			ifres243 = True

		} else {
			tmp248 := PrimIsNumber(V3785)

			var ifres245 Obj

			if True == tmp248 {
				ifres245 = True

			} else {
				tmp247 := PrimIsString(V3785)

				var ifres246 Obj

				if True == tmp247 {
					ifres246 = True

				} else {
					ifres246 = False

				}

				ifres245 = ifres246

			}

			var ifres244 Obj

			if True == ifres245 {
				ifres244 = True

			} else {
				ifres244 = False

			}

			ifres243 = ifres244

		}

		if True == ifres243 {
			__e.Return(False)
			return
		} else {
			tmp238 := MakeNative(func(__e *ControlFlow) {
				tmp239 := MakeNative(func(__e *ControlFlow) {
					W3786 := __e.Get(1)
					_ = W3786
					__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), W3786)
					return
				}, 1)

				tmp240 := PrimStr(V3785)

				__e.TailApply(tmp239, tmp240)
				return

			}, 0)

			tmp241 := MakeNative(func(__e *ControlFlow) {
				Z3787 := __e.Get(1)
				_ = Z3787
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(try_1catch, tmp238, tmp241)
			return

		}

	}, 1)

	tmp250 := Call(__e, ns2_1set, symvariable_2, tmp237)

	_ = tmp250

	tmp251 := MakeNative(func(__e *ControlFlow) {
		V3790 := __e.Get(1)
		_ = V3790
		tmp260 := Call(__e, PrimFunc(symshen_4_7string_2), V3790)

		if True == tmp260 {
			tmp256 := Call(__e, PrimFunc(symhdstr), V3790)

			tmp257 := PrimStringToNumber(tmp256)

			tmp258 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp257)

			if True == tmp258 {
				tmp253 := PrimTailString(V3790)

				tmp254 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp253)

				if True == tmp254 {
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

	tmp261 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp251)

	_ = tmp261

	tmp262 := MakeNative(func(__e *ControlFlow) {
		V3791 := __e.Get(1)
		_ = V3791
		tmp263 := PrimValue(symshen_4_dgensym_d)

		tmp264 := PrimNumberAdd(MakeNumber(1), tmp263)

		tmp265 := PrimSet(symshen_4_dgensym_d, tmp264)

		__e.TailApply(PrimFunc(symconcat), V3791, tmp265)
		return

	}, 1)

	tmp266 := Call(__e, ns2_1set, symgensym, tmp262)

	_ = tmp266

	tmp267 := MakeNative(func(__e *ControlFlow) {
		V3792 := __e.Get(1)
		_ = V3792
		V3793 := __e.Get(2)
		_ = V3793
		tmp268 := PrimStr(V3792)

		tmp269 := PrimStr(V3793)

		tmp270 := PrimStringConcat(tmp268, tmp269)

		__e.Return(PrimIntern(tmp270))
		return

	}, 2)

	tmp271 := Call(__e, ns2_1set, symconcat, tmp267)

	_ = tmp271

	tmp272 := MakeNative(func(__e *ControlFlow) {
		V3794 := __e.Get(1)
		_ = V3794
		V3795 := __e.Get(2)
		_ = V3795
		tmp273 := MakeNative(func(__e *ControlFlow) {
			W3796 := __e.Get(1)
			_ = W3796
			tmp274 := MakeNative(func(__e *ControlFlow) {
				W3797 := __e.Get(1)
				_ = W3797
				tmp275 := MakeNative(func(__e *ControlFlow) {
					W3798 := __e.Get(1)
					_ = W3798
					tmp276 := MakeNative(func(__e *ControlFlow) {
						W3799 := __e.Get(1)
						_ = W3799
						__e.Return(W3796)
						return
					}, 1)

					tmp277 := PrimVectorSet(W3796, MakeNumber(2), V3795)

					__e.TailApply(tmp276, tmp277)
					return

				}, 1)

				tmp278 := PrimVectorSet(W3796, MakeNumber(1), V3794)

				__e.TailApply(tmp275, tmp278)
				return

			}, 1)

			tmp279 := PrimVectorSet(W3796, MakeNumber(0), symshen_4tuple)

			__e.TailApply(tmp274, tmp279)
			return

		}, 1)

		tmp280 := PrimAbsvector(MakeNumber(3))

		__e.TailApply(tmp273, tmp280)
		return

	}, 2)

	tmp281 := Call(__e, ns2_1set, sym_8p, tmp272)

	_ = tmp281

	tmp282 := MakeNative(func(__e *ControlFlow) {
		V3800 := __e.Get(1)
		_ = V3800
		__e.Return(PrimVectorGet(V3800, MakeNumber(1)))
		return
	}, 1)

	tmp283 := Call(__e, ns2_1set, symfst, tmp282)

	_ = tmp283

	tmp284 := MakeNative(func(__e *ControlFlow) {
		V3801 := __e.Get(1)
		_ = V3801
		__e.Return(PrimVectorGet(V3801, MakeNumber(2)))
		return
	}, 1)

	tmp285 := Call(__e, ns2_1set, symsnd, tmp284)

	_ = tmp285

	tmp286 := MakeNative(func(__e *ControlFlow) {
		V3802 := __e.Get(1)
		_ = V3802
		tmp293 := PrimIsVector(V3802)

		if True == tmp293 {
			tmp288 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V3802, MakeNumber(0)))
				return
			}, 0)

			tmp289 := MakeNative(func(__e *ControlFlow) {
				Z3803 := __e.Get(1)
				_ = Z3803
				__e.Return(symshen_4not_1tuple)
				return
			}, 1)

			tmp290 := Call(__e, try_1catch, tmp288, tmp289)

			tmp291 := PrimEqual(symshen_4tuple, tmp290)

			if True == tmp291 {
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

	tmp294 := Call(__e, ns2_1set, symtuple_2, tmp286)

	_ = tmp294

	tmp295 := MakeNative(func(__e *ControlFlow) {
		V3808 := __e.Get(1)
		_ = V3808
		V3809 := __e.Get(2)
		_ = V3809
		tmp302 := PrimEqual(Nil, V3808)

		if True == tmp302 {
			__e.Return(V3809)
			return
		} else {
			tmp300 := PrimIsPair(V3808)

			if True == tmp300 {
				tmp296 := PrimHead(V3808)

				tmp297 := PrimTail(V3808)

				tmp298 := Call(__e, PrimFunc(symappend), tmp297, V3809)

				__e.Return(PrimCons(tmp296, tmp298))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to append a non-list")))
				return
			}

		}

	}, 2)

	tmp303 := Call(__e, ns2_1set, symappend, tmp295)

	_ = tmp303

	tmp304 := MakeNative(func(__e *ControlFlow) {
		V3810 := __e.Get(1)
		_ = V3810
		V3811 := __e.Get(2)
		_ = V3811
		tmp305 := MakeNative(func(__e *ControlFlow) {
			W3812 := __e.Get(1)
			_ = W3812
			tmp306 := MakeNative(func(__e *ControlFlow) {
				W3813 := __e.Get(1)
				_ = W3813
				tmp307 := MakeNative(func(__e *ControlFlow) {
					W3814 := __e.Get(1)
					_ = W3814
					tmp309 := PrimEqual(W3812, MakeNumber(0))

					if True == tmp309 {
						__e.Return(W3814)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4_8v_1help), V3811, MakeNumber(1), W3812, W3814)
						return
					}

				}, 1)

				tmp310 := Call(__e, PrimFunc(symvector_1_6), W3813, MakeNumber(1), V3810)

				__e.TailApply(tmp307, tmp310)
				return

			}, 1)

			tmp311 := PrimNumberAdd(W3812, MakeNumber(1))

			tmp312 := Call(__e, PrimFunc(symvector), tmp311)

			__e.TailApply(tmp306, tmp312)
			return

		}, 1)

		tmp313 := Call(__e, PrimFunc(symlimit), V3811)

		__e.TailApply(tmp305, tmp313)
		return

	}, 2)

	tmp314 := Call(__e, ns2_1set, sym_8v, tmp304)

	_ = tmp314

	tmp315 := MakeNative(func(__e *ControlFlow) {
		V3816 := __e.Get(1)
		_ = V3816
		V3817 := __e.Get(2)
		_ = V3817
		V3818 := __e.Get(3)
		_ = V3818
		V3819 := __e.Get(4)
		_ = V3819
		tmp321 := PrimEqual(V3817, V3818)

		if True == tmp321 {
			tmp316 := PrimNumberAdd(V3818, MakeNumber(1))

			__e.TailApply(PrimFunc(symshen_4copyfromvector), V3816, V3819, V3818, tmp316)
			return

		} else {
			tmp317 := PrimNumberAdd(V3817, MakeNumber(1))

			tmp318 := PrimNumberAdd(V3817, MakeNumber(1))

			tmp319 := Call(__e, PrimFunc(symshen_4copyfromvector), V3816, V3819, V3817, tmp318)

			__e.TailApply(PrimFunc(symshen_4_8v_1help), V3816, tmp317, V3818, tmp319)
			return

		}

	}, 4)

	tmp322 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp315)

	_ = tmp322

	tmp323 := MakeNative(func(__e *ControlFlow) {
		V3820 := __e.Get(1)
		_ = V3820
		V3821 := __e.Get(2)
		_ = V3821
		V3822 := __e.Get(3)
		_ = V3822
		V3823 := __e.Get(4)
		_ = V3823
		tmp324 := MakeNative(func(__e *ControlFlow) {
			tmp325 := Call(__e, PrimFunc(sym_5_1vector), V3820, V3822)

			__e.TailApply(PrimFunc(symvector_1_6), V3821, V3823, tmp325)
			return

		}, 0)

		tmp326 := MakeNative(func(__e *ControlFlow) {
			Z3824 := __e.Get(1)
			_ = Z3824
			__e.Return(V3821)
			return
		}, 1)

		__e.TailApply(try_1catch, tmp324, tmp326)
		return

	}, 4)

	tmp327 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp323)

	_ = tmp327

	tmp328 := MakeNative(func(__e *ControlFlow) {
		V3825 := __e.Get(1)
		_ = V3825
		tmp329 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimFunc(sym_5_1vector), V3825, MakeNumber(1))
			return
		}, 0)

		tmp330 := MakeNative(func(__e *ControlFlow) {
			Z3826 := __e.Get(1)
			_ = Z3826
			__e.Return(PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n")))
			return
		}, 1)

		__e.TailApply(try_1catch, tmp329, tmp330)
		return

	}, 1)

	tmp331 := Call(__e, ns2_1set, symhdv, tmp328)

	_ = tmp331

	tmp332 := MakeNative(func(__e *ControlFlow) {
		V3827 := __e.Get(1)
		_ = V3827
		tmp333 := MakeNative(func(__e *ControlFlow) {
			W3828 := __e.Get(1)
			_ = W3828
			tmp342 := PrimEqual(W3828, MakeNumber(0))

			if True == tmp342 {
				__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
				return
			} else {
				tmp340 := PrimEqual(W3828, MakeNumber(1))

				if True == tmp340 {
					__e.TailApply(PrimFunc(symvector), MakeNumber(0))
					return
				} else {
					tmp334 := MakeNative(func(__e *ControlFlow) {
						W3829 := __e.Get(1)
						_ = W3829
						tmp335 := PrimNumberSubtract(W3828, MakeNumber(1))

						tmp336 := Call(__e, PrimFunc(symvector), tmp335)

						__e.TailApply(PrimFunc(symshen_4tlv_1help), V3827, MakeNumber(2), W3828, tmp336)
						return

					}, 1)

					tmp337 := PrimNumberSubtract(W3828, MakeNumber(1))

					tmp338 := Call(__e, PrimFunc(symvector), tmp337)

					__e.TailApply(tmp334, tmp338)
					return

				}

			}

		}, 1)

		tmp343 := Call(__e, PrimFunc(symlimit), V3827)

		__e.TailApply(tmp333, tmp343)
		return

	}, 1)

	tmp344 := Call(__e, ns2_1set, symtlv, tmp332)

	_ = tmp344

	tmp345 := MakeNative(func(__e *ControlFlow) {
		V3831 := __e.Get(1)
		_ = V3831
		V3832 := __e.Get(2)
		_ = V3832
		V3833 := __e.Get(3)
		_ = V3833
		V3834 := __e.Get(4)
		_ = V3834
		tmp351 := PrimEqual(V3832, V3833)

		if True == tmp351 {
			tmp346 := PrimNumberSubtract(V3833, MakeNumber(1))

			__e.TailApply(PrimFunc(symshen_4copyfromvector), V3831, V3834, V3833, tmp346)
			return

		} else {
			tmp347 := PrimNumberAdd(V3832, MakeNumber(1))

			tmp348 := PrimNumberSubtract(V3832, MakeNumber(1))

			tmp349 := Call(__e, PrimFunc(symshen_4copyfromvector), V3831, V3834, V3832, tmp348)

			__e.TailApply(PrimFunc(symshen_4tlv_1help), V3831, tmp347, V3833, tmp349)
			return

		}

	}, 4)

	tmp352 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp345)

	_ = tmp352

	tmp353 := MakeNative(func(__e *ControlFlow) {
		V3846 := __e.Get(1)
		_ = V3846
		V3847 := __e.Get(2)
		_ = V3847
		tmp369 := PrimEqual(Nil, V3847)

		if True == tmp369 {
			__e.Return(Nil)
			return
		} else {
			tmp367 := PrimIsPair(V3847)

			var ifres358 Obj

			if True == tmp367 {
				tmp365 := PrimHead(V3847)

				tmp366 := PrimIsPair(tmp365)

				var ifres360 Obj

				if True == tmp366 {
					tmp362 := PrimHead(V3847)

					tmp363 := PrimHead(tmp362)

					tmp364 := PrimEqual(V3846, tmp363)

					var ifres361 Obj

					if True == tmp364 {
						ifres361 = True

					} else {
						ifres361 = False

					}

					ifres360 = ifres361

				} else {
					ifres360 = False

				}

				var ifres359 Obj

				if True == ifres360 {
					ifres359 = True

				} else {
					ifres359 = False

				}

				ifres358 = ifres359

			} else {
				ifres358 = False

			}

			if True == ifres358 {
				__e.Return(PrimHead(V3847))
				return
			} else {
				tmp356 := PrimIsPair(V3847)

				if True == tmp356 {
					tmp354 := PrimTail(V3847)

					__e.TailApply(PrimFunc(symassoc), V3846, tmp354)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("attempt to search a non-list with assoc\n")))
					return
				}

			}

		}

	}, 2)

	tmp370 := Call(__e, ns2_1set, symassoc, tmp353)

	_ = tmp370

	tmp371 := MakeNative(func(__e *ControlFlow) {
		V3851 := __e.Get(1)
		_ = V3851
		V3852 := __e.Get(2)
		_ = V3852
		V3853 := __e.Get(3)
		_ = V3853
		tmp394 := PrimEqual(Nil, V3853)

		if True == tmp394 {
			tmp372 := PrimCons(V3851, V3852)

			__e.Return(PrimCons(tmp372, Nil))
			return

		} else {
			tmp392 := PrimIsPair(V3853)

			var ifres383 Obj

			if True == tmp392 {
				tmp390 := PrimHead(V3853)

				tmp391 := PrimIsPair(tmp390)

				var ifres385 Obj

				if True == tmp391 {
					tmp387 := PrimHead(V3853)

					tmp388 := PrimHead(tmp387)

					tmp389 := PrimEqual(V3851, tmp388)

					var ifres386 Obj

					if True == tmp389 {
						ifres386 = True

					} else {
						ifres386 = False

					}

					ifres385 = ifres386

				} else {
					ifres385 = False

				}

				var ifres384 Obj

				if True == ifres385 {
					ifres384 = True

				} else {
					ifres384 = False

				}

				ifres383 = ifres384

			} else {
				ifres383 = False

			}

			if True == ifres383 {
				tmp373 := PrimHead(V3853)

				tmp374 := PrimHead(tmp373)

				tmp375 := PrimCons(tmp374, V3852)

				tmp376 := PrimTail(V3853)

				__e.Return(PrimCons(tmp375, tmp376))
				return

			} else {
				tmp381 := PrimIsPair(V3853)

				if True == tmp381 {
					tmp377 := PrimHead(V3853)

					tmp378 := PrimTail(V3853)

					tmp379 := Call(__e, PrimFunc(symshen_4assoc_1set), V3851, V3852, tmp378)

					__e.Return(PrimCons(tmp377, tmp379))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assoc_1set)
					return
				}

			}

		}

	}, 3)

	tmp395 := Call(__e, ns2_1set, symshen_4assoc_1set, tmp371)

	_ = tmp395

	tmp396 := MakeNative(func(__e *ControlFlow) {
		V3857 := __e.Get(1)
		_ = V3857
		V3858 := __e.Get(2)
		_ = V3858
		tmp414 := PrimEqual(Nil, V3858)

		if True == tmp414 {
			__e.Return(Nil)
			return
		} else {
			tmp412 := PrimIsPair(V3858)

			var ifres403 Obj

			if True == tmp412 {
				tmp410 := PrimHead(V3858)

				tmp411 := PrimIsPair(tmp410)

				var ifres405 Obj

				if True == tmp411 {
					tmp407 := PrimHead(V3858)

					tmp408 := PrimHead(tmp407)

					tmp409 := PrimEqual(V3857, tmp408)

					var ifres406 Obj

					if True == tmp409 {
						ifres406 = True

					} else {
						ifres406 = False

					}

					ifres405 = ifres406

				} else {
					ifres405 = False

				}

				var ifres404 Obj

				if True == ifres405 {
					ifres404 = True

				} else {
					ifres404 = False

				}

				ifres403 = ifres404

			} else {
				ifres403 = False

			}

			if True == ifres403 {
				__e.Return(PrimTail(V3858))
				return
			} else {
				tmp401 := PrimIsPair(V3858)

				if True == tmp401 {
					tmp397 := PrimHead(V3858)

					tmp398 := PrimTail(V3858)

					tmp399 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3857, tmp398)

					__e.Return(PrimCons(tmp397, tmp399))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assoc_1rm)
					return
				}

			}

		}

	}, 2)

	tmp415 := Call(__e, ns2_1set, symshen_4assoc_1rm, tmp396)

	_ = tmp415

	tmp416 := MakeNative(func(__e *ControlFlow) {
		V3861 := __e.Get(1)
		_ = V3861
		tmp420 := PrimEqual(True, V3861)

		if True == tmp420 {
			__e.Return(True)
			return
		} else {
			tmp418 := PrimEqual(False, V3861)

			if True == tmp418 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp421 := Call(__e, ns2_1set, symboolean_2, tmp416)

	_ = tmp421

	tmp422 := MakeNative(func(__e *ControlFlow) {
		V3862 := __e.Get(1)
		_ = V3862
		tmp427 := PrimEqual(MakeNumber(0), V3862)

		if True == tmp427 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp423 := Call(__e, PrimFunc(symstoutput))

			tmp424 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp423)

			_ = tmp424

			tmp425 := PrimNumberSubtract(V3862, MakeNumber(1))

			__e.TailApply(PrimFunc(symnl), tmp425)
			return

		}

	}, 1)

	tmp428 := Call(__e, ns2_1set, symnl, tmp422)

	_ = tmp428

	tmp429 := MakeNative(func(__e *ControlFlow) {
		V3869 := __e.Get(1)
		_ = V3869
		V3870 := __e.Get(2)
		_ = V3870
		tmp440 := PrimEqual(Nil, V3869)

		if True == tmp440 {
			__e.Return(Nil)
			return
		} else {
			tmp438 := PrimIsPair(V3869)

			if True == tmp438 {
				tmp435 := PrimHead(V3869)

				tmp436 := Call(__e, PrimFunc(symelement_2), tmp435, V3870)

				if True == tmp436 {
					tmp430 := PrimTail(V3869)

					__e.TailApply(PrimFunc(symdifference), tmp430, V3870)
					return

				} else {
					tmp431 := PrimHead(V3869)

					tmp432 := PrimTail(V3869)

					tmp433 := Call(__e, PrimFunc(symdifference), tmp432, V3870)

					__e.Return(PrimCons(tmp431, tmp433))
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to find the difference with a non-list\n")))
				return
			}

		}

	}, 2)

	tmp441 := Call(__e, ns2_1set, symdifference, tmp429)

	_ = tmp441

	tmp442 := MakeNative(func(__e *ControlFlow) {
		V3871 := __e.Get(1)
		_ = V3871
		V3872 := __e.Get(2)
		_ = V3872
		__e.Return(V3872)
		return
	}, 2)

	tmp443 := Call(__e, ns2_1set, symdo, tmp442)

	_ = tmp443

	tmp444 := MakeNative(func(__e *ControlFlow) {
		V3884 := __e.Get(1)
		_ = V3884
		V3885 := __e.Get(2)
		_ = V3885
		tmp455 := PrimEqual(Nil, V3885)

		if True == tmp455 {
			__e.Return(False)
			return
		} else {
			tmp453 := PrimIsPair(V3885)

			var ifres449 Obj

			if True == tmp453 {
				tmp451 := PrimHead(V3885)

				tmp452 := PrimEqual(V3884, tmp451)

				var ifres450 Obj

				if True == tmp452 {
					ifres450 = True

				} else {
					ifres450 = False

				}

				ifres449 = ifres450

			} else {
				ifres449 = False

			}

			if True == ifres449 {
				__e.Return(True)
				return
			} else {
				tmp447 := PrimIsPair(V3885)

				if True == tmp447 {
					tmp445 := PrimTail(V3885)

					__e.TailApply(PrimFunc(symelement_2), V3884, tmp445)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("attempt to find an element in a non-list\n")))
					return
				}

			}

		}

	}, 2)

	tmp456 := Call(__e, ns2_1set, symelement_2, tmp444)

	_ = tmp456

	tmp457 := MakeNative(func(__e *ControlFlow) {
		V3888 := __e.Get(1)
		_ = V3888
		tmp459 := PrimEqual(Nil, V3888)

		if True == tmp459 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp460 := Call(__e, ns2_1set, symempty_2, tmp457)

	_ = tmp460

	tmp461 := MakeNative(func(__e *ControlFlow) {
		V3889 := __e.Get(1)
		_ = V3889
		V3890 := __e.Get(2)
		_ = V3890
		tmp462 := Call(__e, V3889, V3890)

		__e.TailApply(PrimFunc(symshen_4fix_1help), V3889, V3890, tmp462)
		return

	}, 2)

	tmp463 := Call(__e, ns2_1set, symfix, tmp461)

	_ = tmp463

	tmp464 := MakeNative(func(__e *ControlFlow) {
		V3896 := __e.Get(1)
		_ = V3896
		V3897 := __e.Get(2)
		_ = V3897
		V3898 := __e.Get(3)
		_ = V3898
		tmp467 := PrimEqual(V3897, V3898)

		if True == tmp467 {
			__e.Return(V3898)
			return
		} else {
			tmp465 := Call(__e, V3896, V3898)

			__e.TailApply(PrimFunc(symshen_4fix_1help), V3896, V3898, tmp465)
			return

		}

	}, 3)

	tmp468 := Call(__e, ns2_1set, symshen_4fix_1help, tmp464)

	_ = tmp468

	tmp469 := MakeNative(func(__e *ControlFlow) {
		V3899 := __e.Get(1)
		_ = V3899
		V3900 := __e.Get(2)
		_ = V3900
		V3901 := __e.Get(3)
		_ = V3901
		V3902 := __e.Get(4)
		_ = V3902
		tmp470 := MakeNative(func(__e *ControlFlow) {
			W3903 := __e.Get(1)
			_ = W3903
			tmp471 := MakeNative(func(__e *ControlFlow) {
				W3905 := __e.Get(1)
				_ = W3905
				tmp472 := MakeNative(func(__e *ControlFlow) {
					W3906 := __e.Get(1)
					_ = W3906
					__e.Return(V3901)
					return
				}, 1)

				tmp473 := Call(__e, PrimFunc(symshen_4dict_1_6), V3902, V3899, W3905)

				__e.TailApply(tmp472, tmp473)
				return

			}, 1)

			tmp474 := Call(__e, PrimFunc(symshen_4assoc_1set), V3900, V3901, W3903)

			__e.TailApply(tmp471, tmp474)
			return

		}, 1)

		tmp475 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimFunc(symshen_4_5_1dict), V3902, V3899)
			return
		}, 0)

		tmp476 := MakeNative(func(__e *ControlFlow) {
			Z3904 := __e.Get(1)
			_ = Z3904
			__e.Return(Nil)
			return
		}, 1)

		tmp477 := Call(__e, try_1catch, tmp475, tmp476)

		__e.TailApply(tmp470, tmp477)
		return

	}, 4)

	tmp478 := Call(__e, ns2_1set, symput, tmp469)

	_ = tmp478

	tmp479 := MakeNative(func(__e *ControlFlow) {
		V3907 := __e.Get(1)
		_ = V3907
		V3908 := __e.Get(2)
		_ = V3908
		V3909 := __e.Get(3)
		_ = V3909
		tmp480 := MakeNative(func(__e *ControlFlow) {
			W3910 := __e.Get(1)
			_ = W3910
			tmp481 := MakeNative(func(__e *ControlFlow) {
				W3912 := __e.Get(1)
				_ = W3912
				tmp482 := MakeNative(func(__e *ControlFlow) {
					W3913 := __e.Get(1)
					_ = W3913
					__e.Return(V3907)
					return
				}, 1)

				tmp483 := Call(__e, PrimFunc(symshen_4dict_1_6), V3909, V3907, W3912)

				__e.TailApply(tmp482, tmp483)
				return

			}, 1)

			tmp484 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3908, W3910)

			__e.TailApply(tmp481, tmp484)
			return

		}, 1)

		tmp485 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimFunc(symshen_4_5_1dict), V3909, V3907)
			return
		}, 0)

		tmp486 := MakeNative(func(__e *ControlFlow) {
			Z3911 := __e.Get(1)
			_ = Z3911
			__e.Return(Nil)
			return
		}, 1)

		tmp487 := Call(__e, try_1catch, tmp485, tmp486)

		__e.TailApply(tmp480, tmp487)
		return

	}, 3)

	tmp488 := Call(__e, ns2_1set, symunput, tmp479)

	_ = tmp488

	tmp489 := MakeNative(func(__e *ControlFlow) {
		V3914 := __e.Get(1)
		_ = V3914
		V3915 := __e.Get(2)
		_ = V3915
		V3916 := __e.Get(3)
		_ = V3916
		tmp490 := MakeNative(func(__e *ControlFlow) {
			W3917 := __e.Get(1)
			_ = W3917
			tmp491 := MakeNative(func(__e *ControlFlow) {
				W3919 := __e.Get(1)
				_ = W3919
				tmp497 := Call(__e, PrimFunc(symempty_2), W3919)

				if True == tmp497 {
					tmp492 := Call(__e, PrimFunc(symshen_4app), V3914, MakeString("\n"), symshen_4s)

					tmp493 := PrimStringConcat(MakeString(" not found for "), tmp492)

					tmp494 := Call(__e, PrimFunc(symshen_4app), V3915, tmp493, symshen_4s)

					tmp495 := PrimStringConcat(MakeString("attribute "), tmp494)

					__e.Return(PrimSimpleError(tmp495))
					return

				} else {
					__e.Return(PrimTail(W3919))
					return
				}

			}, 1)

			tmp498 := Call(__e, PrimFunc(symassoc), V3915, W3917)

			__e.TailApply(tmp491, tmp498)
			return

		}, 1)

		tmp499 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimFunc(symshen_4_5_1dict), V3916, V3914)
			return
		}, 0)

		tmp500 := MakeNative(func(__e *ControlFlow) {
			Z3918 := __e.Get(1)
			_ = Z3918
			tmp501 := Call(__e, PrimFunc(symshen_4app), V3915, MakeString("\n"), symshen_4s)

			tmp502 := PrimStringConcat(MakeString(" has no attributes: "), tmp501)

			tmp503 := Call(__e, PrimFunc(symshen_4app), V3914, tmp502, symshen_4a)

			__e.Return(PrimSimpleError(tmp503))
			return

		}, 1)

		tmp504 := Call(__e, try_1catch, tmp499, tmp500)

		__e.TailApply(tmp490, tmp504)
		return

	}, 3)

	tmp505 := Call(__e, ns2_1set, symget, tmp489)

	_ = tmp505

	tmp506 := MakeNative(func(__e *ControlFlow) {
		V3920 := __e.Get(1)
		_ = V3920
		V3921 := __e.Get(2)
		_ = V3921
		tmp507 := MakeNative(func(__e *ControlFlow) {
			W3922 := __e.Get(1)
			_ = W3922
			tmp509 := PrimEqual(W3922, MakeNumber(0))

			if True == tmp509 {
				__e.Return(MakeNumber(1))
				return
			} else {
				__e.Return(W3922)
				return
			}

		}, 1)

		tmp510 := Call(__e, PrimFunc(symshen_4hashkey), V3920)

		tmp511 := Call(__e, PrimFunc(symshen_4mod), tmp510, V3921)

		__e.TailApply(tmp507, tmp511)
		return

	}, 2)

	tmp512 := Call(__e, ns2_1set, symhash, tmp506)

	_ = tmp512

	tmp513 := MakeNative(func(__e *ControlFlow) {
		V3923 := __e.Get(1)
		_ = V3923
		tmp514 := MakeNative(func(__e *ControlFlow) {
			W3924 := __e.Get(1)
			_ = W3924
			__e.TailApply(PrimFunc(symshen_4prodbutzero), W3924, MakeNumber(1))
			return
		}, 1)

		tmp515 := MakeNative(func(__e *ControlFlow) {
			Z3925 := __e.Get(1)
			_ = Z3925
			__e.Return(PrimStringToNumber(Z3925))
			return
		}, 1)

		tmp516 := Call(__e, PrimFunc(symexplode), V3923)

		tmp517 := Call(__e, PrimFunc(symmap), tmp515, tmp516)

		__e.TailApply(tmp514, tmp517)
		return

	}, 1)

	tmp518 := Call(__e, ns2_1set, symshen_4hashkey, tmp513)

	_ = tmp518

	tmp519 := MakeNative(func(__e *ControlFlow) {
		V3926 := __e.Get(1)
		_ = V3926
		V3927 := __e.Get(2)
		_ = V3927
		tmp538 := PrimEqual(Nil, V3926)

		if True == tmp538 {
			__e.Return(V3927)
			return
		} else {
			tmp536 := PrimIsPair(V3926)

			var ifres532 Obj

			if True == tmp536 {
				tmp534 := PrimHead(V3926)

				tmp535 := PrimEqual(MakeNumber(0), tmp534)

				var ifres533 Obj

				if True == tmp535 {
					ifres533 = True

				} else {
					ifres533 = False

				}

				ifres532 = ifres533

			} else {
				ifres532 = False

			}

			if True == ifres532 {
				tmp520 := PrimTail(V3926)

				__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp520, V3927)
				return

			} else {
				tmp530 := PrimIsPair(V3926)

				if True == tmp530 {
					tmp528 := PrimGreatThan(V3927, MakeNumber(10000000000))

					if True == tmp528 {
						tmp521 := PrimTail(V3926)

						tmp522 := PrimHead(V3926)

						tmp523 := PrimNumberAdd(V3927, tmp522)

						__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp521, tmp523)
						return

					} else {
						tmp524 := PrimTail(V3926)

						tmp525 := PrimHead(V3926)

						tmp526 := PrimNumberMultiply(V3927, tmp525)

						__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp524, tmp526)
						return

					}

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prodbutzero)
					return
				}

			}

		}

	}, 2)

	tmp539 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp519)

	_ = tmp539

	tmp540 := MakeNative(func(__e *ControlFlow) {
		V3928 := __e.Get(1)
		_ = V3928
		V3929 := __e.Get(2)
		_ = V3929
		tmp541 := PrimCons(V3929, Nil)

		tmp542 := Call(__e, PrimFunc(symshen_4multiples), V3928, tmp541)

		__e.TailApply(PrimFunc(symshen_4modh), V3928, tmp542)
		return

	}, 2)

	tmp543 := Call(__e, ns2_1set, symshen_4mod, tmp540)

	_ = tmp543

	tmp544 := MakeNative(func(__e *ControlFlow) {
		V3934 := __e.Get(1)
		_ = V3934
		V3935 := __e.Get(2)
		_ = V3935
		tmp555 := PrimIsPair(V3935)

		var ifres551 Obj

		if True == tmp555 {
			tmp553 := PrimHead(V3935)

			tmp554 := PrimGreatThan(tmp553, V3934)

			var ifres552 Obj

			if True == tmp554 {
				ifres552 = True

			} else {
				ifres552 = False

			}

			ifres551 = ifres552

		} else {
			ifres551 = False

		}

		if True == ifres551 {
			__e.Return(PrimTail(V3935))
			return
		} else {
			tmp549 := PrimIsPair(V3935)

			if True == tmp549 {
				tmp545 := PrimHead(V3935)

				tmp546 := PrimNumberMultiply(MakeNumber(2), tmp545)

				tmp547 := PrimCons(tmp546, V3935)

				__e.TailApply(PrimFunc(symshen_4multiples), V3934, tmp547)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.multiples")))
				return
			}

		}

	}, 2)

	tmp556 := Call(__e, ns2_1set, symshen_4multiples, tmp544)

	_ = tmp556

	tmp557 := MakeNative(func(__e *ControlFlow) {
		V3942 := __e.Get(1)
		_ = V3942
		V3943 := __e.Get(2)
		_ = V3943
		tmp575 := PrimEqual(MakeNumber(0), V3942)

		if True == tmp575 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp573 := PrimEqual(Nil, V3943)

			if True == tmp573 {
				__e.Return(V3942)
				return
			} else {
				tmp571 := PrimIsPair(V3943)

				var ifres567 Obj

				if True == tmp571 {
					tmp569 := PrimHead(V3943)

					tmp570 := PrimGreatThan(tmp569, V3942)

					var ifres568 Obj

					if True == tmp570 {
						ifres568 = True

					} else {
						ifres568 = False

					}

					ifres567 = ifres568

				} else {
					ifres567 = False

				}

				if True == ifres567 {
					tmp560 := PrimTail(V3943)

					tmp561 := Call(__e, PrimFunc(symempty_2), tmp560)

					if True == tmp561 {
						__e.Return(V3942)
						return
					} else {
						tmp558 := PrimTail(V3943)

						__e.TailApply(PrimFunc(symshen_4modh), V3942, tmp558)
						return

					}

				} else {
					tmp565 := PrimIsPair(V3943)

					if True == tmp565 {
						tmp562 := PrimHead(V3943)

						tmp563 := PrimNumberSubtract(V3942, tmp562)

						__e.TailApply(PrimFunc(symshen_4modh), tmp563, V3943)
						return

					} else {
						__e.Return(PrimSimpleError(MakeString("implementation error in shen.modh")))
						return
					}

				}

			}

		}

	}, 2)

	tmp576 := Call(__e, ns2_1set, symshen_4modh, tmp557)

	_ = tmp576

	tmp577 := MakeNative(func(__e *ControlFlow) {
		V3946 := __e.Get(1)
		_ = V3946
		tmp584 := PrimEqual(Nil, V3946)

		if True == tmp584 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp582 := PrimIsPair(V3946)

			if True == tmp582 {
				tmp578 := PrimHead(V3946)

				tmp579 := PrimTail(V3946)

				tmp580 := Call(__e, PrimFunc(symsum), tmp579)

				__e.Return(PrimNumberAdd(tmp578, tmp580))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to sum a non-list\n")))
				return
			}

		}

	}, 1)

	tmp585 := Call(__e, ns2_1set, symsum, tmp577)

	_ = tmp585

	tmp586 := MakeNative(func(__e *ControlFlow) {
		V3951 := __e.Get(1)
		_ = V3951
		tmp588 := PrimIsPair(V3951)

		if True == tmp588 {
			__e.Return(PrimHead(V3951))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("head expects a non-empty list\n")))
			return
		}

	}, 1)

	tmp589 := Call(__e, ns2_1set, symhead, tmp586)

	_ = tmp589

	tmp590 := MakeNative(func(__e *ControlFlow) {
		V3956 := __e.Get(1)
		_ = V3956
		tmp592 := PrimIsPair(V3956)

		if True == tmp592 {
			__e.Return(PrimTail(V3956))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list\n")))
			return
		}

	}, 1)

	tmp593 := Call(__e, ns2_1set, symtail, tmp590)

	_ = tmp593

	tmp594 := MakeNative(func(__e *ControlFlow) {
		V3957 := __e.Get(1)
		_ = V3957
		__e.Return(PrimPos(V3957, MakeNumber(0)))
		return
	}, 1)

	tmp595 := Call(__e, ns2_1set, symhdstr, tmp594)

	_ = tmp595

	tmp596 := MakeNative(func(__e *ControlFlow) {
		V3964 := __e.Get(1)
		_ = V3964
		V3965 := __e.Get(2)
		_ = V3965
		tmp607 := PrimEqual(Nil, V3964)

		if True == tmp607 {
			__e.Return(Nil)
			return
		} else {
			tmp605 := PrimIsPair(V3964)

			if True == tmp605 {
				tmp602 := PrimHead(V3964)

				tmp603 := Call(__e, PrimFunc(symelement_2), tmp602, V3965)

				if True == tmp603 {
					tmp597 := PrimHead(V3964)

					tmp598 := PrimTail(V3964)

					tmp599 := Call(__e, PrimFunc(symintersection), tmp598, V3965)

					__e.Return(PrimCons(tmp597, tmp599))
					return

				} else {
					tmp600 := PrimTail(V3964)

					__e.TailApply(PrimFunc(symintersection), tmp600, V3965)
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n")))
				return
			}

		}

	}, 2)

	tmp608 := Call(__e, ns2_1set, symintersection, tmp596)

	_ = tmp608

	tmp609 := MakeNative(func(__e *ControlFlow) {
		V3966 := __e.Get(1)
		_ = V3966
		__e.TailApply(PrimFunc(symshen_4reverse_1help), V3966, Nil)
		return
	}, 1)

	tmp610 := Call(__e, ns2_1set, symreverse, tmp609)

	_ = tmp610

	tmp611 := MakeNative(func(__e *ControlFlow) {
		V3971 := __e.Get(1)
		_ = V3971
		V3972 := __e.Get(2)
		_ = V3972
		tmp618 := PrimEqual(Nil, V3971)

		if True == tmp618 {
			__e.Return(V3972)
			return
		} else {
			tmp616 := PrimIsPair(V3971)

			if True == tmp616 {
				tmp612 := PrimTail(V3971)

				tmp613 := PrimHead(V3971)

				tmp614 := PrimCons(tmp613, V3972)

				__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp612, tmp614)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to reverse a non-list\n")))
				return
			}

		}

	}, 2)

	tmp619 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp611)

	_ = tmp619

	tmp620 := MakeNative(func(__e *ControlFlow) {
		V3977 := __e.Get(1)
		_ = V3977
		V3978 := __e.Get(2)
		_ = V3978
		tmp631 := PrimEqual(Nil, V3977)

		if True == tmp631 {
			__e.Return(V3978)
			return
		} else {
			tmp629 := PrimIsPair(V3977)

			if True == tmp629 {
				tmp626 := PrimHead(V3977)

				tmp627 := Call(__e, PrimFunc(symelement_2), tmp626, V3978)

				if True == tmp627 {
					tmp621 := PrimTail(V3977)

					__e.TailApply(PrimFunc(symunion), tmp621, V3978)
					return

				} else {
					tmp622 := PrimHead(V3977)

					tmp623 := PrimTail(V3977)

					tmp624 := Call(__e, PrimFunc(symunion), tmp623, V3978)

					__e.Return(PrimCons(tmp622, tmp624))
					return

				}

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to find the union with a non-list\n")))
				return
			}

		}

	}, 2)

	tmp632 := Call(__e, ns2_1set, symunion, tmp620)

	_ = tmp632

	tmp633 := MakeNative(func(__e *ControlFlow) {
		V3979 := __e.Get(1)
		_ = V3979
		tmp634 := MakeNative(func(__e *ControlFlow) {
			W3980 := __e.Get(1)
			_ = W3980
			tmp635 := MakeNative(func(__e *ControlFlow) {
				W3981 := __e.Get(1)
				_ = W3981
				tmp636 := MakeNative(func(__e *ControlFlow) {
					W3982 := __e.Get(1)
					_ = W3982
					tmp642 := PrimEqual(MakeString("y"), W3982)

					if True == tmp642 {
						__e.Return(True)
						return
					} else {
						tmp640 := PrimEqual(MakeString("n"), W3982)

						if True == tmp640 {
							__e.Return(False)
							return
						} else {
							tmp637 := Call(__e, PrimFunc(symstoutput))

							tmp638 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp637)

							_ = tmp638

							__e.TailApply(PrimFunc(symy_1or_1n_2), V3979)
							return

						}

					}

				}, 1)

				tmp643 := Call(__e, PrimFunc(symstinput))

				tmp644 := Call(__e, PrimFunc(symread), tmp643)

				tmp645 := Call(__e, PrimFunc(symshen_4app), tmp644, MakeString(""), symshen_4s)

				__e.TailApply(tmp636, tmp645)
				return

			}, 1)

			tmp646 := Call(__e, PrimFunc(symstoutput))

			tmp647 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp646)

			__e.TailApply(tmp635, tmp647)
			return

		}, 1)

		tmp648 := Call(__e, PrimFunc(symshen_4proc_1nl), V3979)

		tmp649 := Call(__e, PrimFunc(symstoutput))

		tmp650 := Call(__e, PrimFunc(sympr), tmp648, tmp649)

		__e.TailApply(tmp634, tmp650)
		return

	}, 1)

	tmp651 := Call(__e, ns2_1set, symy_1or_1n_2, tmp633)

	_ = tmp651

	tmp652 := MakeNative(func(__e *ControlFlow) {
		V3983 := __e.Get(1)
		_ = V3983
		if True == V3983 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}
	}, 1)

	tmp654 := Call(__e, ns2_1set, symnot, tmp652)

	_ = tmp654

	tmp655 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("")))
		return
	}, 0)

	tmp656 := Call(__e, ns2_1set, symabort, tmp655)

	_ = tmp656

	tmp657 := MakeNative(func(__e *ControlFlow) {
		V3989 := __e.Get(1)
		_ = V3989
		V3990 := __e.Get(2)
		_ = V3990
		V3991 := __e.Get(3)
		_ = V3991
		tmp665 := PrimEqual(V3990, V3991)

		if True == tmp665 {
			__e.Return(V3989)
			return
		} else {
			tmp663 := PrimIsPair(V3991)

			if True == tmp663 {
				tmp658 := PrimHead(V3991)

				tmp659 := Call(__e, PrimFunc(symsubst), V3989, V3990, tmp658)

				tmp660 := PrimTail(V3991)

				tmp661 := Call(__e, PrimFunc(symsubst), V3989, V3990, tmp660)

				__e.Return(PrimCons(tmp659, tmp661))
				return

			} else {
				__e.Return(V3991)
				return
			}

		}

	}, 3)

	tmp666 := Call(__e, ns2_1set, symsubst, tmp657)

	_ = tmp666

	tmp667 := MakeNative(func(__e *ControlFlow) {
		V3992 := __e.Get(1)
		_ = V3992
		tmp668 := Call(__e, PrimFunc(symshen_4app), V3992, MakeString(""), symshen_4a)

		__e.TailApply(PrimFunc(symshen_4explode_1h), tmp668)
		return

	}, 1)

	tmp669 := Call(__e, ns2_1set, symexplode, tmp667)

	_ = tmp669

	tmp670 := MakeNative(func(__e *ControlFlow) {
		V3995 := __e.Get(1)
		_ = V3995
		tmp677 := PrimEqual(MakeString(""), V3995)

		if True == tmp677 {
			__e.Return(Nil)
			return
		} else {
			tmp675 := Call(__e, PrimFunc(symshen_4_7string_2), V3995)

			if True == tmp675 {
				tmp671 := Call(__e, PrimFunc(symhdstr), V3995)

				tmp672 := PrimTailString(V3995)

				tmp673 := Call(__e, PrimFunc(symshen_4explode_1h), tmp672)

				__e.Return(PrimCons(tmp671, tmp673))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in explode-h")))
				return
			}

		}

	}, 1)

	tmp678 := Call(__e, ns2_1set, symshen_4explode_1h, tmp670)

	_ = tmp678

	tmp679 := MakeNative(func(__e *ControlFlow) {
		V3996 := __e.Get(1)
		_ = V3996
		tmp682 := PrimEqual(V3996, MakeString(""))

		var ifres680 Obj

		if True == tmp682 {
			ifres680 = MakeString("")

		} else {
			tmp681 := Call(__e, PrimFunc(symshen_4app), V3996, MakeString("/"), symshen_4a)

			ifres680 = tmp681

		}

		__e.Return(PrimSet(sym_dhome_1directory_d, ifres680))
		return

	}, 1)

	tmp683 := Call(__e, ns2_1set, symcd, tmp679)

	_ = tmp683

	tmp684 := MakeNative(func(__e *ControlFlow) {
		V3997 := __e.Get(1)
		_ = V3997
		V3998 := __e.Get(2)
		_ = V3998
		tmp692 := PrimEqual(Nil, V3998)

		if True == tmp692 {
			__e.Return(True)
			return
		} else {
			tmp690 := PrimIsPair(V3998)

			if True == tmp690 {
				tmp685 := MakeNative(func(__e *ControlFlow) {
					W3999 := __e.Get(1)
					_ = W3999
					tmp686 := PrimTail(V3998)

					__e.TailApply(PrimFunc(symshen_4for_1each), V3997, tmp686)
					return

				}, 1)

				tmp687 := PrimHead(V3998)

				tmp688 := Call(__e, V3997, tmp687)

				__e.TailApply(tmp685, tmp688)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4for_1each)
				return
			}

		}

	}, 2)

	tmp693 := Call(__e, ns2_1set, symshen_4for_1each, tmp684)

	_ = tmp693

	tmp694 := MakeNative(func(__e *ControlFlow) {
		V4000 := __e.Get(1)
		_ = V4000
		V4001 := __e.Get(2)
		_ = V4001
		__e.TailApply(PrimFunc(symshen_4map_1h), V4000, V4001, Nil)
		return
	}, 2)

	tmp695 := Call(__e, ns2_1set, symmap, tmp694)

	_ = tmp695

	tmp696 := MakeNative(func(__e *ControlFlow) {
		V4002 := __e.Get(1)
		_ = V4002
		V4003 := __e.Get(2)
		_ = V4003
		V4004 := __e.Get(3)
		_ = V4004
		tmp704 := PrimEqual(Nil, V4003)

		if True == tmp704 {
			__e.TailApply(PrimFunc(symreverse), V4004)
			return
		} else {
			tmp702 := PrimIsPair(V4003)

			if True == tmp702 {
				tmp697 := PrimTail(V4003)

				tmp698 := PrimHead(V4003)

				tmp699 := Call(__e, V4002, tmp698)

				tmp700 := PrimCons(tmp699, V4004)

				__e.TailApply(PrimFunc(symshen_4map_1h), V4002, tmp697, tmp700)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4map_1h)
				return
			}

		}

	}, 3)

	tmp705 := Call(__e, ns2_1set, symshen_4map_1h, tmp696)

	_ = tmp705

	tmp706 := MakeNative(func(__e *ControlFlow) {
		V4005 := __e.Get(1)
		_ = V4005
		__e.TailApply(PrimFunc(symshen_4length_1h), V4005, MakeNumber(0))
		return
	}, 1)

	tmp707 := Call(__e, ns2_1set, symlength, tmp706)

	_ = tmp707

	tmp708 := MakeNative(func(__e *ControlFlow) {
		V4010 := __e.Get(1)
		_ = V4010
		V4011 := __e.Get(2)
		_ = V4011
		tmp712 := PrimEqual(Nil, V4010)

		if True == tmp712 {
			__e.Return(V4011)
			return
		} else {
			tmp709 := PrimTail(V4010)

			tmp710 := PrimNumberAdd(V4011, MakeNumber(1))

			__e.TailApply(PrimFunc(symshen_4length_1h), tmp709, tmp710)
			return

		}

	}, 2)

	tmp713 := Call(__e, ns2_1set, symshen_4length_1h, tmp708)

	_ = tmp713

	tmp714 := MakeNative(func(__e *ControlFlow) {
		V4017 := __e.Get(1)
		_ = V4017
		V4018 := __e.Get(2)
		_ = V4018
		tmp722 := PrimEqual(V4017, V4018)

		if True == tmp722 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp720 := PrimIsPair(V4018)

			if True == tmp720 {
				tmp715 := PrimHead(V4018)

				tmp716 := Call(__e, PrimFunc(symoccurrences), V4017, tmp715)

				tmp717 := PrimTail(V4018)

				tmp718 := Call(__e, PrimFunc(symoccurrences), V4017, tmp717)

				__e.Return(PrimNumberAdd(tmp716, tmp718))
				return

			} else {
				__e.Return(MakeNumber(0))
				return
			}

		}

	}, 2)

	tmp723 := Call(__e, ns2_1set, symoccurrences, tmp714)

	_ = tmp723

	tmp724 := MakeNative(func(__e *ControlFlow) {
		V4023 := __e.Get(1)
		_ = V4023
		V4024 := __e.Get(2)
		_ = V4024
		tmp737 := PrimEqual(MakeNumber(1), V4023)

		var ifres734 Obj

		if True == tmp737 {
			tmp736 := PrimIsPair(V4024)

			var ifres735 Obj

			if True == tmp736 {
				ifres735 = True

			} else {
				ifres735 = False

			}

			ifres734 = ifres735

		} else {
			ifres734 = False

		}

		if True == ifres734 {
			__e.Return(PrimHead(V4024))
			return
		} else {
			tmp732 := PrimIsPair(V4024)

			if True == tmp732 {
				tmp725 := PrimNumberSubtract(V4023, MakeNumber(1))

				tmp726 := PrimTail(V4024)

				__e.TailApply(PrimFunc(symnth), tmp725, tmp726)
				return

			} else {
				tmp727 := Call(__e, PrimFunc(symshen_4app), V4024, MakeString("\n"), symshen_4a)

				tmp728 := PrimStringConcat(MakeString(", "), tmp727)

				tmp729 := Call(__e, PrimFunc(symshen_4app), V4023, tmp728, symshen_4a)

				tmp730 := PrimStringConcat(MakeString("nth applied to "), tmp729)

				__e.Return(PrimSimpleError(tmp730))
				return

			}

		}

	}, 2)

	tmp738 := Call(__e, ns2_1set, symnth, tmp724)

	_ = tmp738

	tmp739 := MakeNative(func(__e *ControlFlow) {
		V4025 := __e.Get(1)
		_ = V4025
		tmp746 := PrimIsNumber(V4025)

		if True == tmp746 {
			tmp741 := MakeNative(func(__e *ControlFlow) {
				W4026 := __e.Get(1)
				_ = W4026
				tmp742 := Call(__e, PrimFunc(symshen_4magless), W4026, MakeNumber(1))

				__e.TailApply(PrimFunc(symshen_4integer_1test_2), W4026, tmp742)
				return

			}, 1)

			tmp743 := Call(__e, PrimFunc(symshen_4abs), V4025)

			tmp744 := Call(__e, tmp741, tmp743)

			if True == tmp744 {
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

	tmp747 := Call(__e, ns2_1set, syminteger_2, tmp739)

	_ = tmp747

	tmp748 := MakeNative(func(__e *ControlFlow) {
		V4027 := __e.Get(1)
		_ = V4027
		tmp750 := PrimGreatThan(V4027, MakeNumber(0))

		if True == tmp750 {
			__e.Return(V4027)
			return
		} else {
			__e.Return(PrimNumberSubtract(MakeNumber(0), V4027))
			return
		}

	}, 1)

	tmp751 := Call(__e, ns2_1set, symshen_4abs, tmp748)

	_ = tmp751

	tmp752 := MakeNative(func(__e *ControlFlow) {
		V4028 := __e.Get(1)
		_ = V4028
		V4029 := __e.Get(2)
		_ = V4029
		tmp753 := MakeNative(func(__e *ControlFlow) {
			W4030 := __e.Get(1)
			_ = W4030
			tmp755 := PrimGreatThan(W4030, V4028)

			if True == tmp755 {
				__e.Return(V4029)
				return
			} else {
				__e.TailApply(PrimFunc(symshen_4magless), V4028, W4030)
				return
			}

		}, 1)

		tmp756 := PrimNumberMultiply(V4029, MakeNumber(2))

		__e.TailApply(tmp753, tmp756)
		return

	}, 2)

	tmp757 := Call(__e, ns2_1set, symshen_4magless, tmp752)

	_ = tmp757

	tmp758 := MakeNative(func(__e *ControlFlow) {
		V4034 := __e.Get(1)
		_ = V4034
		V4035 := __e.Get(2)
		_ = V4035
		tmp766 := PrimEqual(MakeNumber(0), V4034)

		if True == tmp766 {
			__e.Return(True)
			return
		} else {
			tmp764 := PrimGreatThan(MakeNumber(1), V4034)

			if True == tmp764 {
				__e.Return(False)
				return
			} else {
				tmp759 := MakeNative(func(__e *ControlFlow) {
					W4036 := __e.Get(1)
					_ = W4036
					tmp761 := PrimGreatThan(MakeNumber(0), W4036)

					if True == tmp761 {
						__e.Return(PrimIsInteger(V4034))
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4integer_1test_2), W4036, V4035)
						return
					}

				}, 1)

				tmp762 := PrimNumberSubtract(V4034, V4035)

				__e.TailApply(tmp759, tmp762)
				return

			}

		}

	}, 2)

	tmp767 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp758)

	_ = tmp767

	tmp768 := MakeNative(func(__e *ControlFlow) {
		V4043 := __e.Get(1)
		_ = V4043
		V4044 := __e.Get(2)
		_ = V4044
		tmp776 := PrimEqual(Nil, V4044)

		if True == tmp776 {
			__e.Return(Nil)
			return
		} else {
			tmp774 := PrimIsPair(V4044)

			if True == tmp774 {
				tmp769 := PrimHead(V4044)

				tmp770 := Call(__e, V4043, tmp769)

				tmp771 := PrimTail(V4044)

				tmp772 := Call(__e, PrimFunc(symmapcan), V4043, tmp771)

				__e.TailApply(PrimFunc(symappend), tmp770, tmp772)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("attempt to mapcan over a non-list\n")))
				return
			}

		}

	}, 2)

	tmp777 := Call(__e, ns2_1set, symmapcan, tmp768)

	_ = tmp777

	tmp778 := MakeNative(func(__e *ControlFlow) {
		V4050 := __e.Get(1)
		_ = V4050
		V4051 := __e.Get(2)
		_ = V4051
		tmp780 := PrimEqual(V4050, V4051)

		if True == tmp780 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp781 := Call(__e, ns2_1set, sym_a_a, tmp778)

	_ = tmp781

	tmp782 := MakeNative(func(__e *ControlFlow) {
		V4052 := __e.Get(1)
		_ = V4052
		tmp792 := PrimIsSymbol(V4052)

		if True == tmp792 {
			tmp784 := MakeNative(func(__e *ControlFlow) {
				W4053 := __e.Get(1)
				_ = W4053
				tmp786 := PrimEqual(W4053, symshen_4this_1symbol_1is_1unbound)

				if True == tmp786 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}, 1)

			tmp787 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimValue(V4052))
				return
			}, 0)

			tmp788 := MakeNative(func(__e *ControlFlow) {
				Z4054 := __e.Get(1)
				_ = Z4054
				__e.Return(symshen_4this_1symbol_1is_1unbound)
				return
			}, 1)

			tmp789 := Call(__e, try_1catch, tmp787, tmp788)

			tmp790 := Call(__e, tmp784, tmp789)

			if True == tmp790 {
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

	tmp793 := Call(__e, ns2_1set, symbound_2, tmp782)

	_ = tmp793

	tmp794 := MakeNative(func(__e *ControlFlow) {
		V4055 := __e.Get(1)
		_ = V4055
		tmp800 := PrimEqual(MakeString(""), V4055)

		if True == tmp800 {
			__e.Return(Nil)
			return
		} else {
			tmp795 := PrimPos(V4055, MakeNumber(0))

			tmp796 := PrimStringToNumber(tmp795)

			tmp797 := PrimTailString(V4055)

			tmp798 := Call(__e, PrimFunc(symshen_4string_1_6bytes), tmp797)

			__e.Return(PrimCons(tmp796, tmp798))
			return

		}

	}, 1)

	tmp801 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp794)

	_ = tmp801

	tmp802 := MakeNative(func(__e *ControlFlow) {
		V4056 := __e.Get(1)
		_ = V4056
		tmp806 := PrimLessThan(V4056, MakeNumber(0))

		if True == tmp806 {
			__e.Return(PrimValue(symshen_4_dmaxinferences_d))
			return
		} else {
			tmp804 := PrimIsInteger(V4056)

			if True == tmp804 {
				__e.Return(PrimSet(symshen_4_dmaxinferences_d, V4056))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("maxinferences expects an integer value\n")))
				return
			}

		}

	}, 1)

	tmp807 := Call(__e, ns2_1set, symmaxinferences, tmp802)

	_ = tmp807

	tmp808 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dinfs_d))
		return
	}, 0)

	tmp809 := Call(__e, ns2_1set, syminferences, tmp808)

	_ = tmp809

	tmp810 := MakeNative(func(__e *ControlFlow) {
		V4057 := __e.Get(1)
		_ = V4057
		__e.Return(V4057)
		return
	}, 1)

	tmp811 := Call(__e, ns2_1set, symprotect, tmp810)

	_ = tmp811

	tmp812 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dsterror_d))
		return
	}, 0)

	tmp813 := Call(__e, ns2_1set, symsterror, tmp812)

	_ = tmp813

	tmp814 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dstoutput_d))
		return
	}, 0)

	tmp815 := Call(__e, ns2_1set, symstoutput, tmp814)

	_ = tmp815

	tmp816 := MakeNative(func(__e *ControlFlow) {
		V4058 := __e.Get(1)
		_ = V4058
		tmp817 := MakeNative(func(__e *ControlFlow) {
			W4059 := __e.Get(1)
			_ = W4059
			tmp821 := PrimIsSymbol(W4059)

			if True == tmp821 {
				__e.Return(W4059)
				return
			} else {
				tmp818 := Call(__e, PrimFunc(symshen_4app), V4058, MakeString(" to a symbol"), symshen_4s)

				tmp819 := PrimStringConcat(MakeString("cannot intern "), tmp818)

				__e.Return(PrimSimpleError(tmp819))
				return

			}

		}, 1)

		tmp822 := PrimIntern(V4058)

		__e.TailApply(tmp817, tmp822)
		return

	}, 1)

	tmp823 := Call(__e, ns2_1set, symstring_1_6symbol, tmp816)

	_ = tmp823

	tmp824 := MakeNative(func(__e *ControlFlow) {
		V4062 := __e.Get(1)
		_ = V4062
		tmp828 := PrimEqual(sym_7, V4062)

		if True == tmp828 {
			__e.Return(PrimSet(symshen_4_doptimise_d, True))
			return
		} else {
			tmp826 := PrimEqual(sym_1, V4062)

			if True == tmp826 {
				__e.Return(PrimSet(symshen_4_doptimise_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp829 := Call(__e, ns2_1set, symoptimise, tmp824)

	_ = tmp829

	tmp830 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dos_d))
		return
	}, 0)

	tmp831 := Call(__e, ns2_1set, symos, tmp830)

	_ = tmp831

	tmp832 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dlanguage_d))
		return
	}, 0)

	tmp833 := Call(__e, ns2_1set, symlanguage, tmp832)

	_ = tmp833

	tmp834 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dversion_d))
		return
	}, 0)

	tmp835 := Call(__e, ns2_1set, symversion, tmp834)

	_ = tmp835

	tmp836 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dport_d))
		return
	}, 0)

	tmp837 := Call(__e, ns2_1set, symport, tmp836)

	_ = tmp837

	tmp838 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dporters_d))
		return
	}, 0)

	tmp839 := Call(__e, ns2_1set, symporters, tmp838)

	_ = tmp839

	tmp840 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dimplementation_d))
		return
	}, 0)

	tmp841 := Call(__e, ns2_1set, symimplementation, tmp840)

	_ = tmp841

	tmp842 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_drelease_d))
		return
	}, 0)

	tmp843 := Call(__e, ns2_1set, symrelease, tmp842)

	_ = tmp843

	tmp844 := MakeNative(func(__e *ControlFlow) {
		V4063 := __e.Get(1)
		_ = V4063
		tmp849 := PrimEqual(symnull, V4063)

		if True == tmp849 {
			__e.Return(True)
			return
		} else {
			tmp845 := MakeNative(func(__e *ControlFlow) {
				tmp846 := Call(__e, PrimFunc(symexternal), V4063)

				_ = tmp846

				__e.Return(True)
				return

			}, 0)

			tmp847 := MakeNative(func(__e *ControlFlow) {
				Z4064 := __e.Get(1)
				_ = Z4064
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(try_1catch, tmp845, tmp847)
			return

		}

	}, 1)

	tmp850 := Call(__e, ns2_1set, sympackage_2, tmp844)

	_ = tmp850

	tmp851 := MakeNative(func(__e *ControlFlow) {
		__e.Return(sym_4_4_4)
		return
	}, 0)

	tmp852 := Call(__e, ns2_1set, symfail, tmp851)

	_ = tmp852

	tmp853 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_duserdefs_d))
		return
	}, 0)

	tmp854 := Call(__e, ns2_1set, symuserdefs, tmp853)

	_ = tmp854

	tmp855 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_doptimise_d))
		return
	}, 0)

	tmp856 := Call(__e, ns2_1set, symoptimise_2, tmp855)

	_ = tmp856

	tmp857 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dhush_d))
		return
	}, 0)

	tmp858 := Call(__e, ns2_1set, symhush_2, tmp857)

	_ = tmp858

	tmp859 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
		return
	}, 0)

	tmp860 := Call(__e, ns2_1set, symsystem_1S_2, tmp859)

	_ = tmp860

	tmp861 := MakeNative(func(__e *ControlFlow) {
		V4067 := __e.Get(1)
		_ = V4067
		tmp865 := PrimEqual(sym_7, V4067)

		if True == tmp865 {
			__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
			return
		} else {
			tmp863 := PrimEqual(sym_1, V4067)

			if True == tmp863 {
				__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
				return
			}

		}

	}, 1)

	tmp866 := Call(__e, ns2_1set, symenable_1type_1theory, tmp861)

	_ = tmp866

	tmp867 := MakeNative(func(__e *ControlFlow) {
		V4070 := __e.Get(1)
		_ = V4070
		tmp871 := PrimEqual(sym_7, V4070)

		if True == tmp871 {
			__e.Return(PrimSet(sym_dhush_d, True))
			return
		} else {
			tmp869 := PrimEqual(sym_1, V4070)

			if True == tmp869 {
				__e.Return(PrimSet(sym_dhush_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("hush expects a + or a -\n")))
				return
			}

		}

	}, 1)

	tmp872 := Call(__e, ns2_1set, symhush, tmp867)

	_ = tmp872

	tmp873 := MakeNative(func(__e *ControlFlow) {
		V4073 := __e.Get(1)
		_ = V4073
		tmp877 := PrimEqual(sym_7, V4073)

		if True == tmp877 {
			__e.Return(PrimSet(symshen_4_dtc_d, True))
			return
		} else {
			tmp875 := PrimEqual(sym_1, V4073)

			if True == tmp875 {
				__e.Return(PrimSet(symshen_4_dtc_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
				return
			}

		}

	}, 1)

	tmp878 := Call(__e, ns2_1set, symtc, tmp873)

	_ = tmp878

	tmp879 := MakeNative(func(__e *ControlFlow) {
		V4074 := __e.Get(1)
		_ = V4074
		tmp880 := PrimValue(symshen_4_dsigf_d)

		tmp881 := Call(__e, PrimFunc(symshen_4unassoc), V4074, tmp880)

		tmp882 := PrimSet(symshen_4_dsigf_d, tmp881)

		_ = tmp882

		__e.Return(V4074)
		return

	}, 1)

	tmp883 := Call(__e, ns2_1set, symdestroy, tmp879)

	_ = tmp883

	tmp884 := MakeNative(func(__e *ControlFlow) {
		V4084 := __e.Get(1)
		_ = V4084
		V4085 := __e.Get(2)
		_ = V4085
		tmp902 := PrimEqual(Nil, V4085)

		if True == tmp902 {
			__e.Return(Nil)
			return
		} else {
			tmp900 := PrimIsPair(V4085)

			var ifres891 Obj

			if True == tmp900 {
				tmp898 := PrimHead(V4085)

				tmp899 := PrimIsPair(tmp898)

				var ifres893 Obj

				if True == tmp899 {
					tmp895 := PrimHead(V4085)

					tmp896 := PrimHead(tmp895)

					tmp897 := PrimEqual(V4084, tmp896)

					var ifres894 Obj

					if True == tmp897 {
						ifres894 = True

					} else {
						ifres894 = False

					}

					ifres893 = ifres894

				} else {
					ifres893 = False

				}

				var ifres892 Obj

				if True == ifres893 {
					ifres892 = True

				} else {
					ifres892 = False

				}

				ifres891 = ifres892

			} else {
				ifres891 = False

			}

			if True == ifres891 {
				__e.Return(PrimTail(V4085))
				return
			} else {
				tmp889 := PrimIsPair(V4085)

				if True == tmp889 {
					tmp885 := PrimHead(V4085)

					tmp886 := PrimTail(V4085)

					tmp887 := Call(__e, PrimFunc(symshen_4unassoc), V4084, tmp886)

					__e.Return(PrimCons(tmp885, tmp887))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.unassoc")))
					return
				}

			}

		}

	}, 2)

	tmp903 := Call(__e, ns2_1set, symshen_4unassoc, tmp884)

	_ = tmp903

	tmp904 := MakeNative(func(__e *ControlFlow) {
		V4086 := __e.Get(1)
		_ = V4086
		tmp908 := Call(__e, PrimFunc(sympackage_2), V4086)

		if True == tmp908 {
			__e.Return(PrimSet(symshen_4_dpackage_d, V4086))
			return
		} else {
			tmp905 := Call(__e, PrimFunc(symshen_4app), V4086, MakeString(" does not exist\n"), symshen_4a)

			tmp906 := PrimStringConcat(MakeString("package "), tmp905)

			__e.Return(PrimSimpleError(tmp906))
			return

		}

	}, 1)

	tmp909 := Call(__e, ns2_1set, symin_1package, tmp904)

	_ = tmp909

	tmp910 := MakeNative(func(__e *ControlFlow) {
		V4087 := __e.Get(1)
		_ = V4087
		V4088 := __e.Get(2)
		_ = V4088
		tmp911 := MakeNative(func(__e *ControlFlow) {
			W4089 := __e.Get(1)
			_ = W4089
			tmp912 := MakeNative(func(__e *ControlFlow) {
				W4090 := __e.Get(1)
				_ = W4090
				tmp913 := MakeNative(func(__e *ControlFlow) {
					W4091 := __e.Get(1)
					_ = W4091
					tmp914 := MakeNative(func(__e *ControlFlow) {
						W4092 := __e.Get(1)
						_ = W4092
						__e.Return(V4088)
						return
					}, 1)

					tmp915 := PrimCloseStream(W4089)

					__e.TailApply(tmp914, tmp915)
					return

				}, 1)

				tmp916 := Call(__e, PrimFunc(sympr), W4090, W4089)

				__e.TailApply(tmp913, tmp916)
				return

			}, 1)

			tmp919 := PrimIsString(V4088)

			var ifres917 Obj

			if True == tmp919 {
				ifres917 = V4088

			} else {
				tmp918 := Call(__e, PrimFunc(symshen_4app), V4088, MakeString(""), symshen_4s)

				ifres917 = tmp918

			}

			__e.TailApply(tmp912, ifres917)
			return

		}, 1)

		tmp920 := PrimOpenStream(V4087, symout)

		__e.TailApply(tmp911, tmp920)
		return

	}, 2)

	tmp921 := Call(__e, ns2_1set, symwrite_1to_1file, tmp910)

	_ = tmp921

	tmp922 := MakeNative(func(__e *ControlFlow) {
		tmp923 := Call(__e, PrimFunc(symgensym), symshen_4t)

		__e.TailApply(PrimFunc(symshen_4freshterm), tmp923)
		return

	}, 0)

	tmp924 := Call(__e, ns2_1set, symfresh, tmp922)

	_ = tmp924

	tmp925 := MakeNative(func(__e *ControlFlow) {
		V4093 := __e.Get(1)
		_ = V4093
		V4094 := __e.Get(2)
		_ = V4094
		tmp926 := MakeNative(func(__e *ControlFlow) {
			W4095 := __e.Get(1)
			_ = W4095
			tmp927 := MakeNative(func(__e *ControlFlow) {
				W4096 := __e.Get(1)
				_ = W4096
				tmp928 := MakeNative(func(__e *ControlFlow) {
					W4097 := __e.Get(1)
					_ = W4097
					__e.Return(V4093)
					return
				}, 1)

				tmp929 := PrimCons(V4093, W4096)

				tmp930 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp929)

				__e.TailApply(tmp928, tmp930)
				return

			}, 1)

			tmp931 := Call(__e, PrimFunc(symshen_4lambda_1entry), V4093)

			__e.TailApply(tmp927, tmp931)
			return

		}, 1)

		tmp932 := PrimValue(sym_dproperty_1vector_d)

		tmp933 := Call(__e, PrimFunc(symput), V4093, symarity, V4094, tmp932)

		__e.TailApply(tmp926, tmp933)
		return

	}, 2)

	tmp934 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp925)

	_ = tmp934

	tmp935 := MakeNative(func(__e *ControlFlow) {
		V4100 := __e.Get(1)
		_ = V4100
		V4101 := __e.Get(2)
		_ = V4101
		tmp959 := PrimEqual(MakeNumber(0), V4101)

		if True == tmp959 {
			tmp936 := PrimValue(symshen_4_dspecial_d)

			tmp937 := Call(__e, PrimFunc(symremove), V4100, tmp936)

			tmp938 := PrimSet(symshen_4_dspecial_d, tmp937)

			_ = tmp938

			tmp939 := PrimValue(symshen_4_dextraspecial_d)

			tmp940 := Call(__e, PrimFunc(symremove), V4100, tmp939)

			tmp941 := PrimSet(symshen_4_dextraspecial_d, tmp940)

			_ = tmp941

			__e.Return(V4100)
			return

		} else {
			tmp957 := PrimEqual(MakeNumber(1), V4101)

			if True == tmp957 {
				tmp942 := PrimValue(symshen_4_dspecial_d)

				tmp943 := Call(__e, PrimFunc(symadjoin), V4100, tmp942)

				tmp944 := PrimSet(symshen_4_dspecial_d, tmp943)

				_ = tmp944

				tmp945 := PrimValue(symshen_4_dextraspecial_d)

				tmp946 := Call(__e, PrimFunc(symremove), V4100, tmp945)

				tmp947 := PrimSet(symshen_4_dextraspecial_d, tmp946)

				_ = tmp947

				__e.Return(V4100)
				return

			} else {
				tmp955 := PrimEqual(MakeNumber(2), V4101)

				if True == tmp955 {
					tmp948 := PrimValue(symshen_4_dspecial_d)

					tmp949 := Call(__e, PrimFunc(symremove), V4100, tmp948)

					tmp950 := PrimSet(symshen_4_dspecial_d, tmp949)

					_ = tmp950

					tmp951 := PrimValue(symshen_4_dextraspecial_d)

					tmp952 := Call(__e, PrimFunc(symadjoin), V4100, tmp951)

					tmp953 := PrimSet(symshen_4_dextraspecial_d, tmp952)

					_ = tmp953

					__e.Return(V4100)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n")))
					return
				}

			}

		}

	}, 2)

	tmp960 := Call(__e, ns2_1set, symspecialise, tmp935)

	_ = tmp960

	tmp961 := MakeNative(func(__e *ControlFlow) {
		V4102 := __e.Get(1)
		_ = V4102
		tmp962 := PrimValue(sym_dabsolute_d)

		tmp963 := PrimCons(V4102, tmp962)

		__e.Return(PrimSet(sym_dabsolute_d, tmp963))
		return

	}, 1)

	tmp964 := Call(__e, ns2_1set, symabsolute, tmp961)

	_ = tmp964

	tmp965 := MakeNative(func(__e *ControlFlow) {
		V4103 := __e.Get(1)
		_ = V4103
		tmp966 := PrimValue(sym_dabsolute_d)

		tmp967 := Call(__e, PrimFunc(symremove), V4103, tmp966)

		__e.Return(PrimSet(sym_dabsolute_d, tmp967))
		return

	}, 1)

	__e.TailApply(ns2_1set, symunabsolute, tmp965)
	return

}, 0)
