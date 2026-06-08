package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
	tmp6979 := MakeNative(func(__e *ControlFlow) {
		V6915 := __e.Get(1)
		_ = V6915
		tmp6980 := MakeNative(func(__e *ControlFlow) {
			W6916 := __e.Get(1)
			_ = W6916
			__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V6915, W6916, W6916)
			return
		}, 1)

		tmp6981 := MakeNative(func(__e *ControlFlow) {
			Z6917 := __e.Get(1)
			_ = Z6917
			__e.Return(PrimTail(Z6917))
			return
		}, 1)

		tmp6982 := PrimValue(sym_dmacros_d)

		tmp6983 := Call(__e, PrimFunc(symmap), tmp6981, tmp6982)

		__e.TailApply(tmp6980, tmp6983)
		return

	}, 1)

	tmp6984 := Call(__e, ns2_1set, symmacroexpand, tmp6979)

	_ = tmp6984

	tmp6985 := MakeNative(func(__e *ControlFlow) {
		V6926 := __e.Get(1)
		_ = V6926
		V6927 := __e.Get(2)
		_ = V6927
		V6928 := __e.Get(3)
		_ = V6928
		tmp6995 := PrimEqual(Nil, V6927)

		if True == tmp6995 {
			__e.Return(V6926)
			return
		} else {
			tmp6993 := PrimIsPair(V6927)

			if True == tmp6993 {
				tmp6986 := MakeNative(func(__e *ControlFlow) {
					W6929 := __e.Get(1)
					_ = W6929
					tmp6989 := PrimEqual(V6926, W6929)

					if True == tmp6989 {
						tmp6987 := PrimTail(V6927)

						__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V6926, tmp6987, V6928)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W6929, V6928, V6928)
						return
					}

				}, 1)

				tmp6990 := PrimHead(V6927)

				tmp6991 := Call(__e, PrimFunc(symshen_4walk), tmp6990, V6926)

				__e.TailApply(tmp6986, tmp6991)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
				return
			}

		}

	}, 3)

	tmp6996 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp6985)

	_ = tmp6996

	tmp6997 := MakeNative(func(__e *ControlFlow) {
		V6930 := __e.Get(1)
		_ = V6930
		V6931 := __e.Get(2)
		_ = V6931
		tmp7001 := PrimIsPair(V6931)

		if True == tmp7001 {
			tmp6998 := MakeNative(func(__e *ControlFlow) {
				Z6932 := __e.Get(1)
				_ = Z6932
				__e.TailApply(PrimFunc(symshen_4walk), V6930, Z6932)
				return
			}, 1)

			tmp6999 := Call(__e, PrimFunc(symmap), tmp6998, V6931)

			__e.TailApply(V6930, tmp6999)
			return

		} else {
			__e.TailApply(V6930, V6931)
			return
		}

	}, 2)

	tmp7002 := Call(__e, ns2_1set, symshen_4walk, tmp6997)

	_ = tmp7002

	tmp7003 := MakeNative(func(__e *ControlFlow) {
		V6933 := __e.Get(1)
		_ = V6933
		tmp7004 := MakeNative(func(__e *ControlFlow) {
			GoTo6934 := __e.Get(1)
			_ = GoTo6934
			tmp7315 := PrimIsPair(V6933)

			if True == tmp7315 {
				tmp7005 := MakeNative(func(__e *ControlFlow) {
					Select6939 := __e.Get(1)
					_ = Select6939
					tmp7006 := MakeNative(func(__e *ControlFlow) {
						Select6940 := __e.Get(1)
						_ = Select6940
						tmp7311 := PrimEqual(symdefmacro, Select6939)

						var ifres7308 Obj

						if True == tmp7311 {
							tmp7310 := PrimIsPair(Select6940)

							var ifres7309 Obj

							if True == tmp7310 {
								ifres7309 = True

							} else {
								ifres7309 = False

							}

							ifres7308 = ifres7309

						} else {
							ifres7308 = False

						}

						if True == ifres7308 {
							tmp7007 := PrimHead(Select6940)

							tmp7008 := PrimTail(Select6940)

							__e.TailApply(PrimFunc(symshen_4process_1def), tmp7007, tmp7008)
							return

						} else {
							tmp7306 := PrimEqual(symdefcc, Select6939)

							if True == tmp7306 {
								__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select6940)
								return
							} else {
								tmp7304 := PrimEqual(symu_b, Select6939)

								var ifres7297 Obj

								if True == tmp7304 {
									tmp7303 := PrimIsPair(Select6940)

									var ifres7299 Obj

									if True == tmp7303 {
										tmp7301 := PrimTail(Select6940)

										tmp7302 := PrimEqual(Nil, tmp7301)

										var ifres7300 Obj

										if True == tmp7302 {
											ifres7300 = True

										} else {
											ifres7300 = False

										}

										ifres7299 = ifres7300

									} else {
										ifres7299 = False

									}

									var ifres7298 Obj

									if True == ifres7299 {
										ifres7298 = True

									} else {
										ifres7298 = False

									}

									ifres7297 = ifres7298

								} else {
									ifres7297 = False

								}

								if True == ifres7297 {
									tmp7009 := PrimHead(Select6940)

									tmp7010 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp7009)

									tmp7011 := PrimCons(tmp7010, Nil)

									__e.Return(PrimCons(symprotect, tmp7011))
									return

								} else {
									tmp7295 := PrimEqual(symerror, Select6939)

									var ifres7292 Obj

									if True == tmp7295 {
										tmp7294 := PrimIsPair(Select6940)

										var ifres7293 Obj

										if True == tmp7294 {
											ifres7293 = True

										} else {
											ifres7293 = False

										}

										ifres7292 = ifres7293

									} else {
										ifres7292 = False

									}

									if True == ifres7292 {
										tmp7012 := PrimHead(Select6940)

										tmp7013 := PrimTail(Select6940)

										tmp7014 := Call(__e, PrimFunc(symshen_4mkstr), tmp7012, tmp7013)

										tmp7015 := PrimCons(tmp7014, Nil)

										__e.Return(PrimCons(symsimple_1error, tmp7015))
										return

									} else {
										tmp7290 := PrimEqual(symoutput, Select6939)

										var ifres7287 Obj

										if True == tmp7290 {
											tmp7289 := PrimIsPair(Select6940)

											var ifres7288 Obj

											if True == tmp7289 {
												ifres7288 = True

											} else {
												ifres7288 = False

											}

											ifres7287 = ifres7288

										} else {
											ifres7287 = False

										}

										if True == ifres7287 {
											tmp7016 := PrimHead(Select6940)

											tmp7017 := PrimTail(Select6940)

											tmp7018 := Call(__e, PrimFunc(symshen_4mkstr), tmp7016, tmp7017)

											tmp7019 := PrimCons(symstoutput, Nil)

											tmp7020 := PrimCons(tmp7019, Nil)

											tmp7021 := PrimCons(tmp7018, tmp7020)

											__e.Return(PrimCons(sympr, tmp7021))
											return

										} else {
											tmp7285 := PrimEqual(sympr, Select6939)

											var ifres7278 Obj

											if True == tmp7285 {
												tmp7284 := PrimIsPair(Select6940)

												var ifres7280 Obj

												if True == tmp7284 {
													tmp7282 := PrimTail(Select6940)

													tmp7283 := PrimEqual(Nil, tmp7282)

													var ifres7281 Obj

													if True == tmp7283 {
														ifres7281 = True

													} else {
														ifres7281 = False

													}

													ifres7280 = ifres7281

												} else {
													ifres7280 = False

												}

												var ifres7279 Obj

												if True == ifres7280 {
													ifres7279 = True

												} else {
													ifres7279 = False

												}

												ifres7278 = ifres7279

											} else {
												ifres7278 = False

											}

											if True == ifres7278 {
												tmp7022 := PrimHead(Select6940)

												tmp7023 := PrimCons(symstoutput, Nil)

												tmp7024 := PrimCons(tmp7023, Nil)

												tmp7025 := PrimCons(tmp7022, tmp7024)

												__e.Return(PrimCons(sympr, tmp7025))
												return

											} else {
												tmp7276 := PrimEqual(symmake_1string, Select6939)

												var ifres7273 Obj

												if True == tmp7276 {
													tmp7275 := PrimIsPair(Select6940)

													var ifres7274 Obj

													if True == tmp7275 {
														ifres7274 = True

													} else {
														ifres7274 = False

													}

													ifres7273 = ifres7274

												} else {
													ifres7273 = False

												}

												if True == ifres7273 {
													tmp7026 := PrimHead(Select6940)

													tmp7027 := PrimTail(Select6940)

													__e.TailApply(PrimFunc(symshen_4mkstr), tmp7026, tmp7027)
													return

												} else {
													tmp7271 := PrimEqual(symlineread, Select6939)

													var ifres7268 Obj

													if True == tmp7271 {
														tmp7270 := PrimEqual(Nil, Select6940)

														var ifres7269 Obj

														if True == tmp7270 {
															ifres7269 = True

														} else {
															ifres7269 = False

														}

														ifres7268 = ifres7269

													} else {
														ifres7268 = False

													}

													if True == ifres7268 {
														tmp7028 := PrimCons(symstinput, Nil)

														tmp7029 := PrimCons(tmp7028, Nil)

														__e.Return(PrimCons(symlineread, tmp7029))
														return

													} else {
														tmp7266 := PrimEqual(syminput, Select6939)

														var ifres7263 Obj

														if True == tmp7266 {
															tmp7265 := PrimEqual(Nil, Select6940)

															var ifres7264 Obj

															if True == tmp7265 {
																ifres7264 = True

															} else {
																ifres7264 = False

															}

															ifres7263 = ifres7264

														} else {
															ifres7263 = False

														}

														if True == ifres7263 {
															tmp7030 := PrimCons(symstinput, Nil)

															tmp7031 := PrimCons(tmp7030, Nil)

															__e.Return(PrimCons(syminput, tmp7031))
															return

														} else {
															tmp7261 := PrimEqual(symread, Select6939)

															var ifres7258 Obj

															if True == tmp7261 {
																tmp7260 := PrimEqual(Nil, Select6940)

																var ifres7259 Obj

																if True == tmp7260 {
																	ifres7259 = True

																} else {
																	ifres7259 = False

																}

																ifres7258 = ifres7259

															} else {
																ifres7258 = False

															}

															if True == ifres7258 {
																tmp7032 := PrimCons(symstinput, Nil)

																tmp7033 := PrimCons(tmp7032, Nil)

																__e.Return(PrimCons(symread, tmp7033))
																return

															} else {
																tmp7256 := PrimEqual(syminput_7, Select6939)

																var ifres7249 Obj

																if True == tmp7256 {
																	tmp7255 := PrimIsPair(Select6940)

																	var ifres7251 Obj

																	if True == tmp7255 {
																		tmp7253 := PrimTail(Select6940)

																		tmp7254 := PrimEqual(Nil, tmp7253)

																		var ifres7252 Obj

																		if True == tmp7254 {
																			ifres7252 = True

																		} else {
																			ifres7252 = False

																		}

																		ifres7251 = ifres7252

																	} else {
																		ifres7251 = False

																	}

																	var ifres7250 Obj

																	if True == ifres7251 {
																		ifres7250 = True

																	} else {
																		ifres7250 = False

																	}

																	ifres7249 = ifres7250

																} else {
																	ifres7249 = False

																}

																if True == ifres7249 {
																	tmp7034 := PrimHead(Select6940)

																	tmp7035 := PrimCons(symstinput, Nil)

																	tmp7036 := PrimCons(tmp7035, Nil)

																	tmp7037 := PrimCons(tmp7034, tmp7036)

																	__e.Return(PrimCons(syminput_7, tmp7037))
																	return

																} else {
																	tmp7247 := PrimEqual(symread_1byte, Select6939)

																	var ifres7244 Obj

																	if True == tmp7247 {
																		tmp7246 := PrimEqual(Nil, Select6940)

																		var ifres7245 Obj

																		if True == tmp7246 {
																			ifres7245 = True

																		} else {
																			ifres7245 = False

																		}

																		ifres7244 = ifres7245

																	} else {
																		ifres7244 = False

																	}

																	if True == ifres7244 {
																		__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
																		return
																	} else {
																		tmp7242 := PrimEqual(symprolog_2, Select6939)

																		if True == tmp7242 {
																			__e.TailApply(PrimFunc(symshen_4call_1prolog), Select6940)
																			return
																		} else {
																			tmp7240 := PrimEqual(symdefprolog, Select6939)

																			var ifres7237 Obj

																			if True == tmp7240 {
																				tmp7239 := PrimIsPair(Select6940)

																				var ifres7238 Obj

																				if True == tmp7239 {
																					ifres7238 = True

																				} else {
																					ifres7238 = False

																				}

																				ifres7237 = ifres7238

																			} else {
																				ifres7237 = False

																			}

																			if True == ifres7237 {
																				tmp7038 := PrimHead(Select6940)

																				tmp7039 := PrimTail(Select6940)

																				__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp7038, tmp7039)
																				return

																			} else {
																				tmp7235 := PrimEqual(symdatatype, Select6939)

																				var ifres7232 Obj

																				if True == tmp7235 {
																					tmp7234 := PrimIsPair(Select6940)

																					var ifres7233 Obj

																					if True == tmp7234 {
																						ifres7233 = True

																					} else {
																						ifres7233 = False

																					}

																					ifres7232 = ifres7233

																				} else {
																					ifres7232 = False

																				}

																				if True == ifres7232 {
																					tmp7040 := PrimHead(Select6940)

																					tmp7041 := PrimTail(Select6940)

																					__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp7040, tmp7041)
																					return

																				} else {
																					tmp7230 := PrimEqual(sym_8s, Select6939)

																					if True == tmp7230 {
																						__e.TailApply(PrimFunc(symshen_4process_1_8s), V6933)
																						return
																					} else {
																						tmp7228 := PrimEqual(symsynonyms, Select6939)

																						if True == tmp7228 {
																							__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select6940)
																							return
																						} else {
																							tmp7226 := PrimEqual(symnl, Select6939)

																							var ifres7223 Obj

																							if True == tmp7226 {
																								tmp7225 := PrimEqual(Nil, Select6940)

																								var ifres7224 Obj

																								if True == tmp7225 {
																									ifres7224 = True

																								} else {
																									ifres7224 = False

																								}

																								ifres7223 = ifres7224

																							} else {
																								ifres7223 = False

																							}

																							if True == ifres7223 {
																								tmp7042 := PrimCons(MakeNumber(1), Nil)

																								__e.Return(PrimCons(symnl, tmp7042))
																								return

																							} else {
																								tmp7221 := PrimEqual(symlet, Select6939)

																								if True == tmp7221 {
																									__e.TailApply(PrimFunc(symshen_4process_1let), V6933)
																									return
																								} else {
																									tmp7219 := PrimEqual(sym_c_4, Select6939)

																									if True == tmp7219 {
																										__e.TailApply(PrimFunc(symshen_4process_1lambda), V6933)
																										return
																									} else {
																										tmp7217 := PrimEqual(symcases, Select6939)

																										if True == tmp7217 {
																											__e.TailApply(PrimFunc(symshen_4process_1cases), V6933)
																											return
																										} else {
																											tmp7215 := PrimEqual(symtime, Select6939)

																											var ifres7208 Obj

																											if True == tmp7215 {
																												tmp7214 := PrimIsPair(Select6940)

																												var ifres7210 Obj

																												if True == tmp7214 {
																													tmp7212 := PrimTail(Select6940)

																													tmp7213 := PrimEqual(Nil, tmp7212)

																													var ifres7211 Obj

																													if True == tmp7213 {
																														ifres7211 = True

																													} else {
																														ifres7211 = False

																													}

																													ifres7210 = ifres7211

																												} else {
																													ifres7210 = False

																												}

																												var ifres7209 Obj

																												if True == ifres7210 {
																													ifres7209 = True

																												} else {
																													ifres7209 = False

																												}

																												ifres7208 = ifres7209

																											} else {
																												ifres7208 = False

																											}

																											if True == ifres7208 {
																												tmp7043 := PrimHead(Select6940)

																												__e.TailApply(PrimFunc(symshen_4process_1time), tmp7043)
																												return

																											} else {
																												tmp7206 := PrimEqual(symput, Select6939)

																												var ifres7188 Obj

																												if True == tmp7206 {
																													tmp7205 := PrimIsPair(Select6940)

																													var ifres7190 Obj

																													if True == tmp7205 {
																														tmp7203 := PrimTail(Select6940)

																														tmp7204 := PrimIsPair(tmp7203)

																														var ifres7192 Obj

																														if True == tmp7204 {
																															tmp7200 := PrimTail(Select6940)

																															tmp7201 := PrimTail(tmp7200)

																															tmp7202 := PrimIsPair(tmp7201)

																															var ifres7194 Obj

																															if True == tmp7202 {
																																tmp7196 := PrimTail(Select6940)

																																tmp7197 := PrimTail(tmp7196)

																																tmp7198 := PrimTail(tmp7197)

																																tmp7199 := PrimEqual(Nil, tmp7198)

																																var ifres7195 Obj

																																if True == tmp7199 {
																																	ifres7195 = True

																																} else {
																																	ifres7195 = False

																																}

																																ifres7194 = ifres7195

																															} else {
																																ifres7194 = False

																															}

																															var ifres7193 Obj

																															if True == ifres7194 {
																																ifres7193 = True

																															} else {
																																ifres7193 = False

																															}

																															ifres7192 = ifres7193

																														} else {
																															ifres7192 = False

																														}

																														var ifres7191 Obj

																														if True == ifres7192 {
																															ifres7191 = True

																														} else {
																															ifres7191 = False

																														}

																														ifres7190 = ifres7191

																													} else {
																														ifres7190 = False

																													}

																													var ifres7189 Obj

																													if True == ifres7190 {
																														ifres7189 = True

																													} else {
																														ifres7189 = False

																													}

																													ifres7188 = ifres7189

																												} else {
																													ifres7188 = False

																												}

																												if True == ifres7188 {
																													tmp7044 := PrimHead(Select6940)

																													tmp7045 := PrimTail(Select6940)

																													tmp7046 := PrimHead(tmp7045)

																													tmp7047 := PrimTail(Select6940)

																													tmp7048 := PrimTail(tmp7047)

																													tmp7049 := PrimHead(tmp7048)

																													tmp7050 := PrimCons(sym_dproperty_1vector_d, Nil)

																													tmp7051 := PrimCons(symvalue, tmp7050)

																													tmp7052 := PrimCons(tmp7051, Nil)

																													tmp7053 := PrimCons(tmp7049, tmp7052)

																													tmp7054 := PrimCons(tmp7046, tmp7053)

																													tmp7055 := PrimCons(tmp7044, tmp7054)

																													__e.Return(PrimCons(symput, tmp7055))
																													return

																												} else {
																													tmp7186 := PrimEqual(symget, Select6939)

																													var ifres7174 Obj

																													if True == tmp7186 {
																														tmp7185 := PrimIsPair(Select6940)

																														var ifres7176 Obj

																														if True == tmp7185 {
																															tmp7183 := PrimTail(Select6940)

																															tmp7184 := PrimIsPair(tmp7183)

																															var ifres7178 Obj

																															if True == tmp7184 {
																																tmp7180 := PrimTail(Select6940)

																																tmp7181 := PrimTail(tmp7180)

																																tmp7182 := PrimEqual(Nil, tmp7181)

																																var ifres7179 Obj

																																if True == tmp7182 {
																																	ifres7179 = True

																																} else {
																																	ifres7179 = False

																																}

																																ifres7178 = ifres7179

																															} else {
																																ifres7178 = False

																															}

																															var ifres7177 Obj

																															if True == ifres7178 {
																																ifres7177 = True

																															} else {
																																ifres7177 = False

																															}

																															ifres7176 = ifres7177

																														} else {
																															ifres7176 = False

																														}

																														var ifres7175 Obj

																														if True == ifres7176 {
																															ifres7175 = True

																														} else {
																															ifres7175 = False

																														}

																														ifres7174 = ifres7175

																													} else {
																														ifres7174 = False

																													}

																													if True == ifres7174 {
																														tmp7056 := PrimHead(Select6940)

																														tmp7057 := PrimTail(Select6940)

																														tmp7058 := PrimHead(tmp7057)

																														tmp7059 := PrimCons(sym_dproperty_1vector_d, Nil)

																														tmp7060 := PrimCons(symvalue, tmp7059)

																														tmp7061 := PrimCons(tmp7060, Nil)

																														tmp7062 := PrimCons(tmp7058, tmp7061)

																														tmp7063 := PrimCons(tmp7056, tmp7062)

																														__e.Return(PrimCons(symget, tmp7063))
																														return

																													} else {
																														tmp7172 := PrimEqual(symunput, Select6939)

																														var ifres7160 Obj

																														if True == tmp7172 {
																															tmp7171 := PrimIsPair(Select6940)

																															var ifres7162 Obj

																															if True == tmp7171 {
																																tmp7169 := PrimTail(Select6940)

																																tmp7170 := PrimIsPair(tmp7169)

																																var ifres7164 Obj

																																if True == tmp7170 {
																																	tmp7166 := PrimTail(Select6940)

																																	tmp7167 := PrimTail(tmp7166)

																																	tmp7168 := PrimEqual(Nil, tmp7167)

																																	var ifres7165 Obj

																																	if True == tmp7168 {
																																		ifres7165 = True

																																	} else {
																																		ifres7165 = False

																																	}

																																	ifres7164 = ifres7165

																																} else {
																																	ifres7164 = False

																																}

																																var ifres7163 Obj

																																if True == ifres7164 {
																																	ifres7163 = True

																																} else {
																																	ifres7163 = False

																																}

																																ifres7162 = ifres7163

																															} else {
																																ifres7162 = False

																															}

																															var ifres7161 Obj

																															if True == ifres7162 {
																																ifres7161 = True

																															} else {
																																ifres7161 = False

																															}

																															ifres7160 = ifres7161

																														} else {
																															ifres7160 = False

																														}

																														if True == ifres7160 {
																															tmp7064 := PrimHead(Select6940)

																															tmp7065 := PrimTail(Select6940)

																															tmp7066 := PrimHead(tmp7065)

																															tmp7067 := PrimCons(sym_dproperty_1vector_d, Nil)

																															tmp7068 := PrimCons(symvalue, tmp7067)

																															tmp7069 := PrimCons(tmp7068, Nil)

																															tmp7070 := PrimCons(tmp7066, tmp7069)

																															tmp7071 := PrimCons(tmp7064, tmp7070)

																															__e.Return(PrimCons(symunput, tmp7071))
																															return

																														} else {
																															tmp7158 := PrimEqual(symshen_4_8c, Select6939)

																															var ifres7151 Obj

																															if True == tmp7158 {
																																tmp7157 := PrimIsPair(Select6940)

																																var ifres7153 Obj

																																if True == tmp7157 {
																																	tmp7155 := PrimTail(Select6940)

																																	tmp7156 := PrimEqual(Nil, tmp7155)

																																	var ifres7154 Obj

																																	if True == tmp7156 {
																																		ifres7154 = True

																																	} else {
																																		ifres7154 = False

																																	}

																																	ifres7153 = ifres7154

																																} else {
																																	ifres7153 = False

																																}

																																var ifres7152 Obj

																																if True == ifres7153 {
																																	ifres7152 = True

																																} else {
																																	ifres7152 = False

																																}

																																ifres7151 = ifres7152

																															} else {
																																ifres7151 = False

																															}

																															if True == ifres7151 {
																																tmp7072 := PrimHead(Select6940)

																																__e.TailApply(PrimFunc(symshen_4rcons__form), tmp7072)
																																return

																															} else {
																																tmp7073 := MakeNative(func(__e *ControlFlow) {
																																	GoTo6935 := __e.Get(1)
																																	_ = GoTo6935
																																	tmp7120 := PrimEqual(symshen_4_8ch, Select6939)

																																	if True == tmp7120 {
																																		tmp7118 := PrimIsPair(Select6940)

																																		if True == tmp7118 {
																																			tmp7074 := MakeNative(func(__e *ControlFlow) {
																																				Select6937 := __e.Get(1)
																																				_ = Select6937
																																				tmp7075 := MakeNative(func(__e *ControlFlow) {
																																					Select6938 := __e.Get(1)
																																					_ = Select6938
																																					tmp7114 := PrimIsPair(Select6937)

																																					var ifres7090 Obj

																																					if True == tmp7114 {
																																						tmp7112 := PrimTail(Select6937)

																																						tmp7113 := PrimIsPair(tmp7112)

																																						var ifres7092 Obj

																																						if True == tmp7113 {
																																							tmp7109 := PrimTail(Select6937)

																																							tmp7110 := PrimTail(tmp7109)

																																							tmp7111 := PrimIsPair(tmp7110)

																																							var ifres7094 Obj

																																							if True == tmp7111 {
																																								tmp7105 := PrimTail(Select6937)

																																								tmp7106 := PrimTail(tmp7105)

																																								tmp7107 := PrimTail(tmp7106)

																																								tmp7108 := PrimEqual(Nil, tmp7107)

																																								var ifres7096 Obj

																																								if True == tmp7108 {
																																									tmp7104 := PrimEqual(Nil, Select6938)

																																									var ifres7098 Obj

																																									if True == tmp7104 {
																																										tmp7100 := PrimTail(Select6937)

																																										tmp7101 := PrimHead(tmp7100)

																																										tmp7102 := PrimIntern(MakeString(":"))

																																										tmp7103 := PrimEqual(tmp7101, tmp7102)

																																										var ifres7099 Obj

																																										if True == tmp7103 {
																																											ifres7099 = True

																																										} else {
																																											ifres7099 = False

																																										}

																																										ifres7098 = ifres7099

																																									} else {
																																										ifres7098 = False

																																									}

																																									var ifres7097 Obj

																																									if True == ifres7098 {
																																										ifres7097 = True

																																									} else {
																																										ifres7097 = False

																																									}

																																									ifres7096 = ifres7097

																																								} else {
																																									ifres7096 = False

																																								}

																																								var ifres7095 Obj

																																								if True == ifres7096 {
																																									ifres7095 = True

																																								} else {
																																									ifres7095 = False

																																								}

																																								ifres7094 = ifres7095

																																							} else {
																																								ifres7094 = False

																																							}

																																							var ifres7093 Obj

																																							if True == ifres7094 {
																																								ifres7093 = True

																																							} else {
																																								ifres7093 = False

																																							}

																																							ifres7092 = ifres7093

																																						} else {
																																							ifres7092 = False

																																						}

																																						var ifres7091 Obj

																																						if True == ifres7092 {
																																							ifres7091 = True

																																						} else {
																																							ifres7091 = False

																																						}

																																						ifres7090 = ifres7091

																																					} else {
																																						ifres7090 = False

																																					}

																																					if True == ifres7090 {
																																						tmp7076 := PrimHead(Select6937)

																																						tmp7077 := PrimTail(Select6937)

																																						tmp7078 := PrimHead(tmp7077)

																																						tmp7079 := PrimTail(Select6937)

																																						tmp7080 := PrimTail(tmp7079)

																																						tmp7081 := PrimCons(sym_7, tmp7080)

																																						tmp7082 := PrimCons(tmp7081, Nil)

																																						tmp7083 := PrimCons(tmp7078, tmp7082)

																																						tmp7084 := PrimCons(tmp7076, tmp7083)

																																						tmp7085 := PrimCons(tmp7084, Nil)

																																						tmp7086 := PrimCons(sym_1, tmp7085)

																																						__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp7086)
																																						return

																																					} else {
																																						tmp7088 := PrimEqual(Nil, Select6938)

																																						if True == tmp7088 {
																																							__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), Select6937)
																																							return
																																						} else {
																																							__e.TailApply(PrimFunc(symthaw), GoTo6935)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp7115 := PrimTail(Select6940)

																																				__e.TailApply(tmp7075, tmp7115)
																																				return

																																			}, 1)

																																			tmp7116 := PrimHead(Select6940)

																																			__e.TailApply(tmp7074, tmp7116)
																																			return

																																		} else {
																																			__e.TailApply(PrimFunc(symthaw), GoTo6935)
																																			return
																																		}

																																	} else {
																																		__e.TailApply(PrimFunc(symthaw), GoTo6935)
																																		return
																																	}

																																}, 1)

																																tmp7121 := MakeNative(func(__e *ControlFlow) {
																																	tmp7149 := PrimIsPair(Select6940)

																																	var ifres7129 Obj

																																	if True == tmp7149 {
																																		tmp7147 := PrimTail(Select6940)

																																		tmp7148 := PrimIsPair(tmp7147)

																																		var ifres7131 Obj

																																		if True == tmp7148 {
																																			tmp7144 := PrimTail(Select6940)

																																			tmp7145 := PrimTail(tmp7144)

																																			tmp7146 := PrimIsPair(tmp7145)

																																			var ifres7133 Obj

																																			if True == tmp7146 {
																																				tmp7135 := PrimCons(symdo, Nil)

																																				tmp7136 := PrimCons(sym_d, tmp7135)

																																				tmp7137 := PrimCons(sym_7, tmp7136)

																																				tmp7138 := PrimCons(symor, tmp7137)

																																				tmp7139 := PrimCons(symand, tmp7138)

																																				tmp7140 := PrimCons(symappend, tmp7139)

																																				tmp7141 := PrimCons(sym_8v, tmp7140)

																																				tmp7142 := PrimCons(sym_8p, tmp7141)

																																				tmp7143 := Call(__e, PrimFunc(symelement_2), Select6939, tmp7142)

																																				var ifres7134 Obj

																																				if True == tmp7143 {
																																					ifres7134 = True

																																				} else {
																																					ifres7134 = False

																																				}

																																				ifres7133 = ifres7134

																																			} else {
																																				ifres7133 = False

																																			}

																																			var ifres7132 Obj

																																			if True == ifres7133 {
																																				ifres7132 = True

																																			} else {
																																				ifres7132 = False

																																			}

																																			ifres7131 = ifres7132

																																		} else {
																																			ifres7131 = False

																																		}

																																		var ifres7130 Obj

																																		if True == ifres7131 {
																																			ifres7130 = True

																																		} else {
																																			ifres7130 = False

																																		}

																																		ifres7129 = ifres7130

																																	} else {
																																		ifres7129 = False

																																	}

																																	if True == ifres7129 {
																																		tmp7122 := PrimHead(Select6940)

																																		tmp7123 := PrimTail(Select6940)

																																		tmp7124 := PrimCons(Select6939, tmp7123)

																																		tmp7125 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp7124)

																																		tmp7126 := PrimCons(tmp7125, Nil)

																																		tmp7127 := PrimCons(tmp7122, tmp7126)

																																		__e.Return(PrimCons(Select6939, tmp7127))
																																		return

																																	} else {
																																		__e.TailApply(PrimFunc(symthaw), GoTo6934)
																																		return
																																	}

																																}, 0)

																																__e.TailApply(tmp7073, tmp7121)
																																return

																															}

																														}

																													}

																												}

																											}

																										}

																									}

																								}

																							}

																						}

																					}

																				}

																			}

																		}

																	}

																}

															}

														}

													}

												}

											}

										}

									}

								}

							}

						}

					}, 1)

					tmp7312 := PrimTail(V6933)

					__e.TailApply(tmp7006, tmp7312)
					return

				}, 1)

				tmp7313 := PrimHead(V6933)

				__e.TailApply(tmp7005, tmp7313)
				return

			} else {
				__e.TailApply(PrimFunc(symthaw), GoTo6934)
				return
			}

		}, 1)

		tmp7316 := MakeNative(func(__e *ControlFlow) {
			__e.Return(V6933)
			return
		}, 0)

		__e.TailApply(tmp7004, tmp7316)
		return

	}, 1)

	tmp7317 := Call(__e, ns2_1set, symshen_4macros, tmp7003)

	_ = tmp7317

	tmp7318 := MakeNative(func(__e *ControlFlow) {
		V6941 := __e.Get(1)
		_ = V6941
		tmp7319 := MakeNative(func(__e *ControlFlow) {
			GoTo6942 := __e.Get(1)
			_ = GoTo6942
			tmp7353 := PrimIsPair(V6941)

			if True == tmp7353 {
				tmp7320 := MakeNative(func(__e *ControlFlow) {
					Select6943 := __e.Get(1)
					_ = Select6943
					tmp7321 := MakeNative(func(__e *ControlFlow) {
						Select6944 := __e.Get(1)
						_ = Select6944
						tmp7349 := PrimEqual(sym_7, Select6943)

						var ifres7342 Obj

						if True == tmp7349 {
							tmp7348 := PrimIsPair(Select6944)

							var ifres7344 Obj

							if True == tmp7348 {
								tmp7346 := PrimTail(Select6944)

								tmp7347 := PrimEqual(Nil, tmp7346)

								var ifres7345 Obj

								if True == tmp7347 {
									ifres7345 = True

								} else {
									ifres7345 = False

								}

								ifres7344 = ifres7345

							} else {
								ifres7344 = False

							}

							var ifres7343 Obj

							if True == ifres7344 {
								ifres7343 = True

							} else {
								ifres7343 = False

							}

							ifres7342 = ifres7343

						} else {
							ifres7342 = False

						}

						if True == ifres7342 {
							tmp7322 := PrimHead(Select6944)

							tmp7323 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp7322)

							tmp7324 := PrimCons(tmp7323, Nil)

							__e.Return(PrimCons(sym_7, tmp7324))
							return

						} else {
							tmp7340 := PrimEqual(sym_1, Select6943)

							var ifres7333 Obj

							if True == tmp7340 {
								tmp7339 := PrimIsPair(Select6944)

								var ifres7335 Obj

								if True == tmp7339 {
									tmp7337 := PrimTail(Select6944)

									tmp7338 := PrimEqual(Nil, tmp7337)

									var ifres7336 Obj

									if True == tmp7338 {
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

							if True == ifres7333 {
								tmp7325 := PrimHead(Select6944)

								tmp7326 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp7325)

								tmp7327 := PrimCons(tmp7326, Nil)

								__e.Return(PrimCons(sym_1, tmp7327))
								return

							} else {
								tmp7328 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select6943)

								tmp7329 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select6944)

								tmp7330 := PrimCons(tmp7329, Nil)

								tmp7331 := PrimCons(tmp7328, tmp7330)

								__e.Return(PrimCons(symcons, tmp7331))
								return

							}

						}

					}, 1)

					tmp7350 := PrimTail(V6941)

					__e.TailApply(tmp7321, tmp7350)
					return

				}, 1)

				tmp7351 := PrimHead(V6941)

				__e.TailApply(tmp7320, tmp7351)
				return

			} else {
				__e.TailApply(PrimFunc(symthaw), GoTo6942)
				return
			}

		}, 1)

		tmp7354 := MakeNative(func(__e *ControlFlow) {
			__e.Return(V6941)
			return
		}, 0)

		__e.TailApply(tmp7319, tmp7354)
		return

	}, 1)

	tmp7355 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp7318)

	_ = tmp7355

	tmp7356 := MakeNative(func(__e *ControlFlow) {
		V6945 := __e.Get(1)
		_ = V6945
		V6946 := __e.Get(2)
		_ = V6946
		tmp7357 := MakeNative(func(__e *ControlFlow) {
			W6947 := __e.Get(1)
			_ = W6947
			tmp7358 := MakeNative(func(__e *ControlFlow) {
				W6948 := __e.Get(1)
				_ = W6948
				tmp7359 := MakeNative(func(__e *ControlFlow) {
					W6949 := __e.Get(1)
					_ = W6949
					__e.Return(V6945)
					return
				}, 1)

				tmp7360 := Call(__e, PrimFunc(symfn), V6945)

				tmp7361 := Call(__e, PrimFunc(symshen_4record_1macro), V6945, tmp7360)

				__e.TailApply(tmp7359, tmp7361)
				return

			}, 1)

			tmp7362 := Call(__e, PrimFunc(symappend), V6946, W6947)

			tmp7363 := PrimCons(V6945, tmp7362)

			tmp7364 := PrimCons(symdefine, tmp7363)

			tmp7365 := Call(__e, PrimFunc(symeval), tmp7364)

			__e.TailApply(tmp7358, tmp7365)
			return

		}, 1)

		tmp7366 := PrimCons(symX, Nil)

		tmp7367 := PrimCons(sym_1_6, tmp7366)

		tmp7368 := PrimCons(symX, tmp7367)

		__e.TailApply(tmp7357, tmp7368)
		return

	}, 2)

	tmp7369 := Call(__e, ns2_1set, symshen_4process_1def, tmp7356)

	_ = tmp7369

	tmp7370 := MakeNative(func(__e *ControlFlow) {
		V6950 := __e.Get(1)
		_ = V6950
		tmp7410 := PrimIsPair(V6950)

		var ifres7384 Obj

		if True == tmp7410 {
			tmp7408 := PrimHead(V6950)

			tmp7409 := PrimEqual(symlet, tmp7408)

			var ifres7386 Obj

			if True == tmp7409 {
				tmp7406 := PrimTail(V6950)

				tmp7407 := PrimIsPair(tmp7406)

				var ifres7388 Obj

				if True == tmp7407 {
					tmp7403 := PrimTail(V6950)

					tmp7404 := PrimTail(tmp7403)

					tmp7405 := PrimIsPair(tmp7404)

					var ifres7390 Obj

					if True == tmp7405 {
						tmp7399 := PrimTail(V6950)

						tmp7400 := PrimTail(tmp7399)

						tmp7401 := PrimTail(tmp7400)

						tmp7402 := PrimIsPair(tmp7401)

						var ifres7392 Obj

						if True == tmp7402 {
							tmp7394 := PrimTail(V6950)

							tmp7395 := PrimTail(tmp7394)

							tmp7396 := PrimTail(tmp7395)

							tmp7397 := PrimTail(tmp7396)

							tmp7398 := PrimIsPair(tmp7397)

							var ifres7393 Obj

							if True == tmp7398 {
								ifres7393 = True

							} else {
								ifres7393 = False

							}

							ifres7392 = ifres7393

						} else {
							ifres7392 = False

						}

						var ifres7391 Obj

						if True == ifres7392 {
							ifres7391 = True

						} else {
							ifres7391 = False

						}

						ifres7390 = ifres7391

					} else {
						ifres7390 = False

					}

					var ifres7389 Obj

					if True == ifres7390 {
						ifres7389 = True

					} else {
						ifres7389 = False

					}

					ifres7388 = ifres7389

				} else {
					ifres7388 = False

				}

				var ifres7387 Obj

				if True == ifres7388 {
					ifres7387 = True

				} else {
					ifres7387 = False

				}

				ifres7386 = ifres7387

			} else {
				ifres7386 = False

			}

			var ifres7385 Obj

			if True == ifres7386 {
				ifres7385 = True

			} else {
				ifres7385 = False

			}

			ifres7384 = ifres7385

		} else {
			ifres7384 = False

		}

		if True == ifres7384 {
			tmp7371 := PrimTail(V6950)

			tmp7372 := PrimHead(tmp7371)

			tmp7373 := PrimTail(V6950)

			tmp7374 := PrimTail(tmp7373)

			tmp7375 := PrimHead(tmp7374)

			tmp7376 := PrimTail(V6950)

			tmp7377 := PrimTail(tmp7376)

			tmp7378 := PrimTail(tmp7377)

			tmp7379 := PrimCons(symlet, tmp7378)

			tmp7380 := PrimCons(tmp7379, Nil)

			tmp7381 := PrimCons(tmp7375, tmp7380)

			tmp7382 := PrimCons(tmp7372, tmp7381)

			__e.Return(PrimCons(symlet, tmp7382))
			return

		} else {
			__e.Return(V6950)
			return
		}

	}, 1)

	tmp7411 := Call(__e, ns2_1set, symshen_4process_1let, tmp7370)

	_ = tmp7411

	tmp7412 := MakeNative(func(__e *ControlFlow) {
		V6951 := __e.Get(1)
		_ = V6951
		tmp7413 := MakeNative(func(__e *ControlFlow) {
			GoTo6953 := __e.Get(1)
			_ = GoTo6953
			tmp7448 := PrimIsPair(V6951)

			if True == tmp7448 {
				tmp7414 := MakeNative(func(__e *ControlFlow) {
					Select6960 := __e.Get(1)
					_ = Select6960
					tmp7444 := PrimHead(V6951)

					tmp7445 := PrimEqual(sym_8s, tmp7444)

					if True == tmp7445 {
						tmp7442 := PrimIsPair(Select6960)

						if True == tmp7442 {
							tmp7415 := MakeNative(func(__e *ControlFlow) {
								Select6958 := __e.Get(1)
								_ = Select6958
								tmp7416 := MakeNative(func(__e *ControlFlow) {
									Select6959 := __e.Get(1)
									_ = Select6959
									tmp7438 := PrimIsPair(Select6959)

									if True == tmp7438 {
										tmp7417 := MakeNative(func(__e *ControlFlow) {
											Select6957 := __e.Get(1)
											_ = Select6957
											tmp7435 := PrimIsPair(Select6957)

											if True == tmp7435 {
												tmp7418 := PrimCons(sym_8s, Select6959)

												tmp7419 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp7418)

												tmp7420 := PrimCons(tmp7419, Nil)

												tmp7421 := PrimCons(Select6958, tmp7420)

												__e.Return(PrimCons(sym_8s, tmp7421))
												return

											} else {
												tmp7433 := PrimEqual(Nil, Select6957)

												var ifres7430 Obj

												if True == tmp7433 {
													tmp7432 := PrimIsString(Select6958)

													var ifres7431 Obj

													if True == tmp7432 {
														ifres7431 = True

													} else {
														ifres7431 = False

													}

													ifres7430 = ifres7431

												} else {
													ifres7430 = False

												}

												if True == ifres7430 {
													tmp7422 := MakeNative(func(__e *ControlFlow) {
														W6952 := __e.Get(1)
														_ = W6952
														tmp7426 := Call(__e, PrimFunc(symlength), W6952)

														tmp7427 := PrimGreatThan(tmp7426, MakeNumber(1))

														if True == tmp7427 {
															tmp7423 := Call(__e, PrimFunc(symappend), W6952, Select6959)

															tmp7424 := PrimCons(sym_8s, tmp7423)

															__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp7424)
															return

														} else {
															__e.Return(V6951)
															return
														}

													}, 1)

													tmp7428 := Call(__e, PrimFunc(symexplode), Select6958)

													__e.TailApply(tmp7422, tmp7428)
													return

												} else {
													__e.TailApply(PrimFunc(symthaw), GoTo6953)
													return
												}

											}

										}, 1)

										tmp7436 := PrimTail(Select6959)

										__e.TailApply(tmp7417, tmp7436)
										return

									} else {
										__e.TailApply(PrimFunc(symthaw), GoTo6953)
										return
									}

								}, 1)

								tmp7439 := PrimTail(Select6960)

								__e.TailApply(tmp7416, tmp7439)
								return

							}, 1)

							tmp7440 := PrimHead(Select6960)

							__e.TailApply(tmp7415, tmp7440)
							return

						} else {
							__e.TailApply(PrimFunc(symthaw), GoTo6953)
							return
						}

					} else {
						__e.TailApply(PrimFunc(symthaw), GoTo6953)
						return
					}

				}, 1)

				tmp7446 := PrimTail(V6951)

				__e.TailApply(tmp7414, tmp7446)
				return

			} else {
				__e.TailApply(PrimFunc(symthaw), GoTo6953)
				return
			}

		}, 1)

		tmp7449 := MakeNative(func(__e *ControlFlow) {
			__e.Return(V6951)
			return
		}, 0)

		__e.TailApply(tmp7413, tmp7449)
		return

	}, 1)

	tmp7450 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp7412)

	_ = tmp7450

	tmp7451 := MakeNative(func(__e *ControlFlow) {
		V6961 := __e.Get(1)
		_ = V6961
		V6962 := __e.Get(2)
		_ = V6962
		tmp7452 := MakeNative(func(__e *ControlFlow) {
			W6963 := __e.Get(1)
			_ = W6963
			tmp7453 := MakeNative(func(__e *ControlFlow) {
				W6964 := __e.Get(1)
				_ = W6964
				__e.Return(W6963)
				return
			}, 1)

			tmp7454 := MakeNative(func(__e *ControlFlow) {
				Z6965 := __e.Get(1)
				_ = Z6965
				__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z6965)
				return
			}, 1)

			tmp7455 := PrimCons(W6963, V6962)

			tmp7456 := Call(__e, PrimFunc(symcompile), tmp7454, tmp7455)

			__e.TailApply(tmp7453, tmp7456)
			return

		}, 1)

		tmp7457 := Call(__e, PrimFunc(symshen_4intern_1type), V6961)

		__e.TailApply(tmp7452, tmp7457)
		return

	}, 2)

	tmp7458 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp7451)

	_ = tmp7458

	tmp7459 := MakeNative(func(__e *ControlFlow) {
		V6966 := __e.Get(1)
		_ = V6966
		tmp7460 := PrimStr(V6966)

		tmp7461 := PrimStringConcat(tmp7460, MakeString("#type"))

		__e.Return(PrimIntern(tmp7461))
		return

	}, 1)

	tmp7462 := Call(__e, ns2_1set, symshen_4intern_1type, tmp7459)

	_ = tmp7462

	tmp7463 := MakeNative(func(__e *ControlFlow) {
		V6967 := __e.Get(1)
		_ = V6967
		tmp7464 := PrimValue(symshen_4_dsynonyms_d)

		tmp7465 := Call(__e, PrimFunc(symappend), V6967, tmp7464)

		tmp7466 := PrimSet(symshen_4_dsynonyms_d, tmp7465)

		__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp7466)
		return

	}, 1)

	tmp7467 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp7463)

	_ = tmp7467

	tmp7468 := MakeNative(func(__e *ControlFlow) {
		V6970 := __e.Get(1)
		_ = V6970
		tmp7518 := PrimIsPair(V6970)

		var ifres7479 Obj

		if True == tmp7518 {
			tmp7516 := PrimHead(V6970)

			tmp7517 := PrimEqual(symdefun, tmp7516)

			var ifres7481 Obj

			if True == tmp7517 {
				tmp7514 := PrimTail(V6970)

				tmp7515 := PrimIsPair(tmp7514)

				var ifres7483 Obj

				if True == tmp7515 {
					tmp7511 := PrimTail(V6970)

					tmp7512 := PrimTail(tmp7511)

					tmp7513 := PrimIsPair(tmp7512)

					var ifres7485 Obj

					if True == tmp7513 {
						tmp7507 := PrimTail(V6970)

						tmp7508 := PrimTail(tmp7507)

						tmp7509 := PrimHead(tmp7508)

						tmp7510 := PrimIsPair(tmp7509)

						var ifres7487 Obj

						if True == tmp7510 {
							tmp7502 := PrimTail(V6970)

							tmp7503 := PrimTail(tmp7502)

							tmp7504 := PrimHead(tmp7503)

							tmp7505 := PrimTail(tmp7504)

							tmp7506 := PrimEqual(Nil, tmp7505)

							var ifres7489 Obj

							if True == tmp7506 {
								tmp7498 := PrimTail(V6970)

								tmp7499 := PrimTail(tmp7498)

								tmp7500 := PrimTail(tmp7499)

								tmp7501 := PrimIsPair(tmp7500)

								var ifres7491 Obj

								if True == tmp7501 {
									tmp7493 := PrimTail(V6970)

									tmp7494 := PrimTail(tmp7493)

									tmp7495 := PrimTail(tmp7494)

									tmp7496 := PrimTail(tmp7495)

									tmp7497 := PrimEqual(Nil, tmp7496)

									var ifres7492 Obj

									if True == tmp7497 {
										ifres7492 = True

									} else {
										ifres7492 = False

									}

									ifres7491 = ifres7492

								} else {
									ifres7491 = False

								}

								var ifres7490 Obj

								if True == ifres7491 {
									ifres7490 = True

								} else {
									ifres7490 = False

								}

								ifres7489 = ifres7490

							} else {
								ifres7489 = False

							}

							var ifres7488 Obj

							if True == ifres7489 {
								ifres7488 = True

							} else {
								ifres7488 = False

							}

							ifres7487 = ifres7488

						} else {
							ifres7487 = False

						}

						var ifres7486 Obj

						if True == ifres7487 {
							ifres7486 = True

						} else {
							ifres7486 = False

						}

						ifres7485 = ifres7486

					} else {
						ifres7485 = False

					}

					var ifres7484 Obj

					if True == ifres7485 {
						ifres7484 = True

					} else {
						ifres7484 = False

					}

					ifres7483 = ifres7484

				} else {
					ifres7483 = False

				}

				var ifres7482 Obj

				if True == ifres7483 {
					ifres7482 = True

				} else {
					ifres7482 = False

				}

				ifres7481 = ifres7482

			} else {
				ifres7481 = False

			}

			var ifres7480 Obj

			if True == ifres7481 {
				ifres7480 = True

			} else {
				ifres7480 = False

			}

			ifres7479 = ifres7480

		} else {
			ifres7479 = False

		}

		if True == ifres7479 {
			tmp7469 := PrimTail(V6970)

			tmp7470 := PrimTail(tmp7469)

			tmp7471 := PrimHead(tmp7470)

			tmp7472 := PrimHead(tmp7471)

			tmp7473 := PrimTail(V6970)

			tmp7474 := PrimTail(tmp7473)

			tmp7475 := PrimTail(tmp7474)

			tmp7476 := PrimCons(tmp7472, tmp7475)

			tmp7477 := PrimCons(sym_c_4, tmp7476)

			__e.TailApply(PrimFunc(symeval), tmp7477)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lambda_1of_1defun)
			return
		}

	}, 1)

	tmp7519 := Call(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp7468)

	_ = tmp7519

	tmp7520 := MakeNative(func(__e *ControlFlow) {
		V6971 := __e.Get(1)
		_ = V6971
		tmp7521 := MakeNative(func(__e *ControlFlow) {
			W6972 := __e.Get(1)
			_ = W6972
			tmp7522 := MakeNative(func(__e *ControlFlow) {
				W6974 := __e.Get(1)
				_ = W6974
				tmp7523 := MakeNative(func(__e *ControlFlow) {
					W6975 := __e.Get(1)
					_ = W6975
					__e.Return(symsynonyms)
					return
				}, 1)

				tmp7524 := PrimSet(symshen_4_ddemodulation_1function_d, W6974)

				__e.TailApply(tmp7523, tmp7524)
				return

			}, 1)

			tmp7525 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W6972)

			tmp7526 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef), symshen_4demod, tmp7525)

			tmp7527 := Call(__e, PrimFunc(symshen_4lambda_1of_1defun), tmp7526)

			__e.TailApply(tmp7522, tmp7527)
			return

		}, 1)

		tmp7528 := MakeNative(func(__e *ControlFlow) {
			Z6973 := __e.Get(1)
			_ = Z6973
			__e.TailApply(PrimFunc(symshen_4curry_1type), Z6973)
			return
		}, 1)

		tmp7529 := Call(__e, PrimFunc(symmap), tmp7528, V6971)

		__e.TailApply(tmp7521, tmp7529)
		return

	}, 1)

	tmp7530 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp7520)

	_ = tmp7530

	tmp7531 := MakeNative(func(__e *ControlFlow) {
		V6978 := __e.Get(1)
		_ = V6978
		tmp7553 := PrimEqual(Nil, V6978)

		if True == tmp7553 {
			tmp7532 := MakeNative(func(__e *ControlFlow) {
				W6979 := __e.Get(1)
				_ = W6979
				tmp7533 := PrimCons(W6979, Nil)

				tmp7534 := PrimCons(sym_1_6, tmp7533)

				__e.Return(PrimCons(W6979, tmp7534))
				return

			}, 1)

			tmp7535 := Call(__e, PrimFunc(symgensym), symX)

			__e.TailApply(tmp7532, tmp7535)
			return

		} else {
			tmp7551 := PrimIsPair(V6978)

			var ifres7547 Obj

			if True == tmp7551 {
				tmp7549 := PrimTail(V6978)

				tmp7550 := PrimIsPair(tmp7549)

				var ifres7548 Obj

				if True == tmp7550 {
					ifres7548 = True

				} else {
					ifres7548 = False

				}

				ifres7547 = ifres7548

			} else {
				ifres7547 = False

			}

			if True == ifres7547 {
				tmp7536 := PrimHead(V6978)

				tmp7537 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7536)

				tmp7538 := PrimTail(V6978)

				tmp7539 := PrimHead(tmp7538)

				tmp7540 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7539)

				tmp7541 := PrimTail(V6978)

				tmp7542 := PrimTail(tmp7541)

				tmp7543 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp7542)

				tmp7544 := PrimCons(tmp7540, tmp7543)

				tmp7545 := PrimCons(sym_1_6, tmp7544)

				__e.Return(PrimCons(tmp7537, tmp7545))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
				return
			}

		}

	}, 1)

	tmp7554 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp7531)

	_ = tmp7554

	tmp7555 := MakeNative(func(__e *ControlFlow) {
		V6980 := __e.Get(1)
		_ = V6980
		tmp7556 := MakeNative(func(__e *ControlFlow) {
			GoTo6981 := __e.Get(1)
			_ = GoTo6981
			tmp7584 := PrimIsPair(V6980)

			if True == tmp7584 {
				tmp7557 := MakeNative(func(__e *ControlFlow) {
					Select6988 := __e.Get(1)
					_ = Select6988
					tmp7580 := PrimHead(V6980)

					tmp7581 := PrimEqual(sym_c_4, tmp7580)

					if True == tmp7581 {
						tmp7578 := PrimIsPair(Select6988)

						if True == tmp7578 {
							tmp7558 := MakeNative(func(__e *ControlFlow) {
								Select6986 := __e.Get(1)
								_ = Select6986
								tmp7559 := MakeNative(func(__e *ControlFlow) {
									Select6987 := __e.Get(1)
									_ = Select6987
									tmp7574 := PrimIsPair(Select6987)

									if True == tmp7574 {
										tmp7560 := MakeNative(func(__e *ControlFlow) {
											Select6985 := __e.Get(1)
											_ = Select6985
											tmp7571 := PrimIsPair(Select6985)

											if True == tmp7571 {
												tmp7561 := PrimCons(sym_c_4, Select6987)

												tmp7562 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp7561)

												tmp7563 := PrimCons(tmp7562, Nil)

												tmp7564 := PrimCons(Select6986, tmp7563)

												__e.Return(PrimCons(symlambda, tmp7564))
												return

											} else {
												tmp7569 := PrimEqual(Nil, Select6985)

												if True == tmp7569 {
													tmp7567 := PrimIsVariable(Select6986)

													if True == tmp7567 {
														__e.Return(PrimCons(symlambda, Select6988))
														return
													} else {
														tmp7565 := Call(__e, PrimFunc(symshen_4app), Select6986, MakeString(" is not a variable\n"), symshen_4s)

														__e.Return(PrimSimpleError(tmp7565))
														return

													}

												} else {
													__e.TailApply(PrimFunc(symthaw), GoTo6981)
													return
												}

											}

										}, 1)

										tmp7572 := PrimTail(Select6987)

										__e.TailApply(tmp7560, tmp7572)
										return

									} else {
										__e.TailApply(PrimFunc(symthaw), GoTo6981)
										return
									}

								}, 1)

								tmp7575 := PrimTail(Select6988)

								__e.TailApply(tmp7559, tmp7575)
								return

							}, 1)

							tmp7576 := PrimHead(Select6988)

							__e.TailApply(tmp7558, tmp7576)
							return

						} else {
							__e.TailApply(PrimFunc(symthaw), GoTo6981)
							return
						}

					} else {
						__e.TailApply(PrimFunc(symthaw), GoTo6981)
						return
					}

				}, 1)

				tmp7582 := PrimTail(V6980)

				__e.TailApply(tmp7557, tmp7582)
				return

			} else {
				__e.TailApply(PrimFunc(symthaw), GoTo6981)
				return
			}

		}, 1)

		tmp7585 := MakeNative(func(__e *ControlFlow) {
			__e.Return(V6980)
			return
		}, 0)

		__e.TailApply(tmp7556, tmp7585)
		return

	}, 1)

	tmp7586 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp7555)

	_ = tmp7586

	tmp7587 := MakeNative(func(__e *ControlFlow) {
		V6991 := __e.Get(1)
		_ = V6991
		tmp7588 := MakeNative(func(__e *ControlFlow) {
			GoTo6992 := __e.Get(1)
			_ = GoTo6992
			tmp7628 := PrimIsPair(V6991)

			if True == tmp7628 {
				tmp7589 := MakeNative(func(__e *ControlFlow) {
					Select7000 := __e.Get(1)
					_ = Select7000
					tmp7624 := PrimHead(V6991)

					tmp7625 := PrimEqual(symcases, tmp7624)

					if True == tmp7625 {
						tmp7622 := PrimIsPair(Select7000)

						if True == tmp7622 {
							tmp7590 := MakeNative(func(__e *ControlFlow) {
								Select6998 := __e.Get(1)
								_ = Select6998
								tmp7591 := MakeNative(func(__e *ControlFlow) {
									Select6999 := __e.Get(1)
									_ = Select6999
									tmp7618 := PrimEqual(True, Select6998)

									var ifres7615 Obj

									if True == tmp7618 {
										tmp7617 := PrimIsPair(Select6999)

										var ifres7616 Obj

										if True == tmp7617 {
											ifres7616 = True

										} else {
											ifres7616 = False

										}

										ifres7615 = ifres7616

									} else {
										ifres7615 = False

									}

									if True == ifres7615 {
										__e.Return(PrimHead(Select6999))
										return
									} else {
										tmp7592 := MakeNative(func(__e *ControlFlow) {
											GoTo6995 := __e.Get(1)
											_ = GoTo6995
											tmp7610 := PrimIsPair(Select6999)

											if True == tmp7610 {
												tmp7593 := MakeNative(func(__e *ControlFlow) {
													Select6996 := __e.Get(1)
													_ = Select6996
													tmp7594 := MakeNative(func(__e *ControlFlow) {
														Select6997 := __e.Get(1)
														_ = Select6997
														tmp7606 := PrimEqual(Nil, Select6997)

														if True == tmp7606 {
															tmp7595 := PrimCons(MakeString("error: cases exhausted"), Nil)

															tmp7596 := PrimCons(symsimple_1error, tmp7595)

															tmp7597 := PrimCons(tmp7596, Nil)

															tmp7598 := PrimCons(Select6996, tmp7597)

															tmp7599 := PrimCons(Select6998, tmp7598)

															__e.Return(PrimCons(symif, tmp7599))
															return

														} else {
															tmp7600 := PrimCons(symcases, Select6997)

															tmp7601 := Call(__e, PrimFunc(symshen_4process_1cases), tmp7600)

															tmp7602 := PrimCons(tmp7601, Nil)

															tmp7603 := PrimCons(Select6996, tmp7602)

															tmp7604 := PrimCons(Select6998, tmp7603)

															__e.Return(PrimCons(symif, tmp7604))
															return

														}

													}, 1)

													tmp7607 := PrimTail(Select6999)

													__e.TailApply(tmp7594, tmp7607)
													return

												}, 1)

												tmp7608 := PrimHead(Select6999)

												__e.TailApply(tmp7593, tmp7608)
												return

											} else {
												__e.TailApply(PrimFunc(symthaw), GoTo6995)
												return
											}

										}, 1)

										tmp7611 := MakeNative(func(__e *ControlFlow) {
											tmp7613 := PrimEqual(Nil, Select6999)

											if True == tmp7613 {
												__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
												return
											} else {
												__e.TailApply(PrimFunc(symthaw), GoTo6992)
												return
											}

										}, 0)

										__e.TailApply(tmp7592, tmp7611)
										return

									}

								}, 1)

								tmp7619 := PrimTail(Select7000)

								__e.TailApply(tmp7591, tmp7619)
								return

							}, 1)

							tmp7620 := PrimHead(Select7000)

							__e.TailApply(tmp7590, tmp7620)
							return

						} else {
							__e.TailApply(PrimFunc(symthaw), GoTo6992)
							return
						}

					} else {
						__e.TailApply(PrimFunc(symthaw), GoTo6992)
						return
					}

				}, 1)

				tmp7626 := PrimTail(V6991)

				__e.TailApply(tmp7589, tmp7626)
				return

			} else {
				__e.TailApply(PrimFunc(symthaw), GoTo6992)
				return
			}

		}, 1)

		tmp7629 := MakeNative(func(__e *ControlFlow) {
			__e.Return(V6991)
			return
		}, 0)

		__e.TailApply(tmp7588, tmp7629)
		return

	}, 1)

	tmp7630 := Call(__e, ns2_1set, symshen_4process_1cases, tmp7587)

	_ = tmp7630

	tmp7631 := MakeNative(func(__e *ControlFlow) {
		V7001 := __e.Get(1)
		_ = V7001
		tmp7632 := PrimCons(symrun, Nil)

		tmp7633 := PrimCons(symget_1time, tmp7632)

		tmp7634 := PrimCons(symrun, Nil)

		tmp7635 := PrimCons(symget_1time, tmp7634)

		tmp7636 := PrimCons(symStart, Nil)

		tmp7637 := PrimCons(symFinish, tmp7636)

		tmp7638 := PrimCons(sym_1, tmp7637)

		tmp7639 := PrimCons(symTime, Nil)

		tmp7640 := PrimCons(symstr, tmp7639)

		tmp7641 := PrimCons(MakeString(" secs\n"), Nil)

		tmp7642 := PrimCons(tmp7640, tmp7641)

		tmp7643 := PrimCons(symcn, tmp7642)

		tmp7644 := PrimCons(tmp7643, Nil)

		tmp7645 := PrimCons(MakeString("\nrun time: "), tmp7644)

		tmp7646 := PrimCons(symcn, tmp7645)

		tmp7647 := PrimCons(symstoutput, Nil)

		tmp7648 := PrimCons(tmp7647, Nil)

		tmp7649 := PrimCons(tmp7646, tmp7648)

		tmp7650 := PrimCons(sympr, tmp7649)

		tmp7651 := PrimCons(symResult, Nil)

		tmp7652 := PrimCons(tmp7650, tmp7651)

		tmp7653 := PrimCons(symMessage, tmp7652)

		tmp7654 := PrimCons(tmp7638, tmp7653)

		tmp7655 := PrimCons(symTime, tmp7654)

		tmp7656 := PrimCons(tmp7635, tmp7655)

		tmp7657 := PrimCons(symFinish, tmp7656)

		tmp7658 := PrimCons(V7001, tmp7657)

		tmp7659 := PrimCons(symResult, tmp7658)

		tmp7660 := PrimCons(tmp7633, tmp7659)

		tmp7661 := PrimCons(symStart, tmp7660)

		__e.Return(PrimCons(symlet, tmp7661))
		return

	}, 1)

	tmp7662 := Call(__e, ns2_1set, symshen_4process_1time, tmp7631)

	_ = tmp7662

	tmp7663 := MakeNative(func(__e *ControlFlow) {
		V7002 := __e.Get(1)
		_ = V7002
		tmp7689 := PrimIsPair(V7002)

		var ifres7674 Obj

		if True == tmp7689 {
			tmp7687 := PrimTail(V7002)

			tmp7688 := PrimIsPair(tmp7687)

			var ifres7676 Obj

			if True == tmp7688 {
				tmp7684 := PrimTail(V7002)

				tmp7685 := PrimTail(tmp7684)

				tmp7686 := PrimIsPair(tmp7685)

				var ifres7678 Obj

				if True == tmp7686 {
					tmp7680 := PrimTail(V7002)

					tmp7681 := PrimTail(tmp7680)

					tmp7682 := PrimTail(tmp7681)

					tmp7683 := PrimIsPair(tmp7682)

					var ifres7679 Obj

					if True == tmp7683 {
						ifres7679 = True

					} else {
						ifres7679 = False

					}

					ifres7678 = ifres7679

				} else {
					ifres7678 = False

				}

				var ifres7677 Obj

				if True == ifres7678 {
					ifres7677 = True

				} else {
					ifres7677 = False

				}

				ifres7676 = ifres7677

			} else {
				ifres7676 = False

			}

			var ifres7675 Obj

			if True == ifres7676 {
				ifres7675 = True

			} else {
				ifres7675 = False

			}

			ifres7674 = ifres7675

		} else {
			ifres7674 = False

		}

		if True == ifres7674 {
			tmp7664 := PrimHead(V7002)

			tmp7665 := PrimTail(V7002)

			tmp7666 := PrimHead(tmp7665)

			tmp7667 := PrimHead(V7002)

			tmp7668 := PrimTail(V7002)

			tmp7669 := PrimTail(tmp7668)

			tmp7670 := PrimCons(tmp7667, tmp7669)

			tmp7671 := PrimCons(tmp7670, Nil)

			tmp7672 := PrimCons(tmp7666, tmp7671)

			__e.Return(PrimCons(tmp7664, tmp7672))
			return

		} else {
			__e.Return(V7002)
			return
		}

	}, 1)

	tmp7690 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp7663)

	_ = tmp7690

	tmp7691 := MakeNative(func(__e *ControlFlow) {
		V7003 := __e.Get(1)
		_ = V7003
		tmp7692 := PrimStr(V7003)

		tmp7693 := Call(__e, PrimFunc(symshen_4mu_1h), tmp7692)

		__e.Return(PrimIntern(tmp7693))
		return

	}, 1)

	tmp7694 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp7691)

	_ = tmp7694

	tmp7695 := MakeNative(func(__e *ControlFlow) {
		V7004 := __e.Get(1)
		_ = V7004
		tmp7714 := PrimEqual(MakeString(""), V7004)

		if True == tmp7714 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp7712 := Call(__e, PrimFunc(symshen_4_7string_2), V7004)

			if True == tmp7712 {
				tmp7696 := MakeNative(func(__e *ControlFlow) {
					W7005 := __e.Get(1)
					_ = W7005
					tmp7697 := MakeNative(func(__e *ControlFlow) {
						W7006 := __e.Get(1)
						_ = W7006
						tmp7698 := MakeNative(func(__e *ControlFlow) {
							W7007 := __e.Get(1)
							_ = W7007
							tmp7699 := PrimTailString(V7004)

							tmp7700 := Call(__e, PrimFunc(symshen_4mu_1h), tmp7699)

							__e.TailApply(PrimFunc(sym_8s), W7007, tmp7700)
							return

						}, 1)

						tmp7707 := PrimGreatEqual(W7005, MakeNumber(97))

						var ifres7704 Obj

						if True == tmp7707 {
							tmp7706 := PrimLessEqual(W7005, MakeNumber(122))

							var ifres7705 Obj

							if True == tmp7706 {
								ifres7705 = True

							} else {
								ifres7705 = False

							}

							ifres7704 = ifres7705

						} else {
							ifres7704 = False

						}

						var ifres7701 Obj

						if True == ifres7704 {
							tmp7702 := PrimNumberToString(W7006)

							ifres7701 = tmp7702

						} else {
							tmp7703 := Call(__e, PrimFunc(symhdstr), V7004)

							ifres7701 = tmp7703

						}

						__e.TailApply(tmp7698, ifres7701)
						return

					}, 1)

					tmp7708 := PrimNumberSubtract(W7005, MakeNumber(32))

					__e.TailApply(tmp7697, tmp7708)
					return

				}, 1)

				tmp7709 := Call(__e, PrimFunc(symhdstr), V7004)

				tmp7710 := PrimStringToNumber(tmp7709)

				__e.TailApply(tmp7696, tmp7710)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4mu_1h)
				return
			}

		}

	}, 1)

	tmp7715 := Call(__e, ns2_1set, symshen_4mu_1h, tmp7695)

	_ = tmp7715

	tmp7716 := MakeNative(func(__e *ControlFlow) {
		V7008 := __e.Get(1)
		_ = V7008
		V7009 := __e.Get(2)
		_ = V7009
		tmp7717 := PrimValue(sym_dmacros_d)

		tmp7718 := Call(__e, PrimFunc(symshen_4update_1assoc), V7008, V7009, tmp7717)

		__e.Return(PrimSet(sym_dmacros_d, tmp7718))
		return

	}, 2)

	tmp7719 := Call(__e, ns2_1set, symshen_4record_1macro, tmp7716)

	_ = tmp7719

	tmp7720 := MakeNative(func(__e *ControlFlow) {
		V7019 := __e.Get(1)
		_ = V7019
		V7020 := __e.Get(2)
		_ = V7020
		V7021 := __e.Get(3)
		_ = V7021
		tmp7740 := PrimEqual(Nil, V7021)

		if True == tmp7740 {
			tmp7721 := PrimCons(V7019, V7020)

			__e.Return(PrimCons(tmp7721, Nil))
			return

		} else {
			tmp7722 := MakeNative(func(__e *ControlFlow) {
				GoTo7022 := __e.Get(1)
				_ = GoTo7022
				tmp7737 := PrimIsPair(V7021)

				if True == tmp7737 {
					tmp7723 := MakeNative(func(__e *ControlFlow) {
						Select7023 := __e.Get(1)
						_ = Select7023
						tmp7724 := MakeNative(func(__e *ControlFlow) {
							Select7024 := __e.Get(1)
							_ = Select7024
							tmp7733 := PrimIsPair(Select7023)

							var ifres7729 Obj

							if True == tmp7733 {
								tmp7731 := PrimHead(Select7023)

								tmp7732 := PrimEqual(V7019, tmp7731)

								var ifres7730 Obj

								if True == tmp7732 {
									ifres7730 = True

								} else {
									ifres7730 = False

								}

								ifres7729 = ifres7730

							} else {
								ifres7729 = False

							}

							if True == ifres7729 {
								tmp7725 := PrimHead(Select7023)

								tmp7726 := PrimCons(tmp7725, V7020)

								__e.Return(PrimCons(tmp7726, Select7024))
								return

							} else {
								tmp7727 := Call(__e, PrimFunc(symshen_4update_1assoc), V7019, V7020, Select7024)

								__e.Return(PrimCons(Select7023, tmp7727))
								return

							}

						}, 1)

						tmp7734 := PrimTail(V7021)

						__e.TailApply(tmp7724, tmp7734)
						return

					}, 1)

					tmp7735 := PrimHead(V7021)

					__e.TailApply(tmp7723, tmp7735)
					return

				} else {
					__e.TailApply(PrimFunc(symthaw), GoTo7022)
					return
				}

			}, 1)

			tmp7738 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
				return
			}, 0)

			__e.TailApply(tmp7722, tmp7738)
			return

		}

	}, 3)

	tmp7741 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp7720)

	_ = tmp7741

	tmp7742 := MakeNative(func(__e *ControlFlow) {
		tmp7750 := Call(__e, PrimFunc(symstinput))

		tmp7751 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp7750)

		if True == tmp7751 {
			tmp7743 := PrimCons(symstinput, Nil)

			tmp7744 := PrimCons(tmp7743, Nil)

			tmp7745 := PrimCons(symshen_4read_1unit_1string, tmp7744)

			tmp7746 := PrimCons(tmp7745, Nil)

			__e.Return(PrimCons(symstring_1_6n, tmp7746))
			return

		} else {
			tmp7747 := PrimCons(symstinput, Nil)

			tmp7748 := PrimCons(tmp7747, Nil)

			__e.Return(PrimCons(symread_1byte, tmp7748))
			return

		}

	}, 0)

	tmp7752 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp7742)

	_ = tmp7752

	tmp7753 := MakeNative(func(__e *ControlFlow) {
		V7025 := __e.Get(1)
		_ = V7025
		tmp7754 := MakeNative(func(__e *ControlFlow) {
			W7026 := __e.Get(1)
			_ = W7026
			tmp7755 := MakeNative(func(__e *ControlFlow) {
				W7027 := __e.Get(1)
				_ = W7027
				tmp7756 := MakeNative(func(__e *ControlFlow) {
					W7028 := __e.Get(1)
					_ = W7028
					tmp7757 := MakeNative(func(__e *ControlFlow) {
						W7029 := __e.Get(1)
						_ = W7029
						tmp7758 := MakeNative(func(__e *ControlFlow) {
							W7030 := __e.Get(1)
							_ = W7030
							tmp7759 := MakeNative(func(__e *ControlFlow) {
								W7032 := __e.Get(1)
								_ = W7032
								tmp7760 := MakeNative(func(__e *ControlFlow) {
									W7033 := __e.Get(1)
									_ = W7033
									tmp7761 := MakeNative(func(__e *ControlFlow) {
										W7034 := __e.Get(1)
										_ = W7034
										tmp7762 := MakeNative(func(__e *ControlFlow) {
											W7035 := __e.Get(1)
											_ = W7035
											tmp7763 := MakeNative(func(__e *ControlFlow) {
												W7036 := __e.Get(1)
												_ = W7036
												tmp7764 := MakeNative(func(__e *ControlFlow) {
													W7037 := __e.Get(1)
													_ = W7037
													tmp7765 := PrimCons(W7029, Nil)

													tmp7766 := PrimCons(W7028, tmp7765)

													tmp7767 := PrimCons(W7027, tmp7766)

													tmp7768 := PrimCons(W7026, tmp7767)

													__e.Return(PrimCons(W7037, tmp7768))
													return

												}, 1)

												tmp7769 := Call(__e, PrimFunc(symshen_4continue), W7032, W7030, W7033, W7034, W7035, W7036)

												tmp7770 := PrimCons(tmp7769, Nil)

												tmp7771 := PrimCons(W7036, tmp7770)

												tmp7772 := PrimCons(symlambda, tmp7771)

												tmp7773 := PrimCons(tmp7772, Nil)

												tmp7774 := PrimCons(W7035, tmp7773)

												tmp7775 := PrimCons(symlambda, tmp7774)

												tmp7776 := PrimCons(tmp7775, Nil)

												tmp7777 := PrimCons(W7034, tmp7776)

												tmp7778 := PrimCons(symlambda, tmp7777)

												tmp7779 := PrimCons(tmp7778, Nil)

												tmp7780 := PrimCons(W7033, tmp7779)

												tmp7781 := PrimCons(symlambda, tmp7780)

												__e.TailApply(tmp7764, tmp7781)
												return

											}, 1)

											tmp7782 := Call(__e, PrimFunc(symgensym), symC)

											__e.TailApply(tmp7763, tmp7782)
											return

										}, 1)

										tmp7783 := Call(__e, PrimFunc(symgensym), symK)

										__e.TailApply(tmp7762, tmp7783)
										return

									}, 1)

									tmp7784 := Call(__e, PrimFunc(symgensym), symL)

									__e.TailApply(tmp7761, tmp7784)
									return

								}, 1)

								tmp7785 := Call(__e, PrimFunc(symgensym), symV)

								__e.TailApply(tmp7760, tmp7785)
								return

							}, 1)

							tmp7786 := Call(__e, PrimFunc(symshen_4received), V7025)

							__e.TailApply(tmp7759, tmp7786)
							return

						}, 1)

						tmp7787 := MakeNative(func(__e *ControlFlow) {
							Z7031 := __e.Get(1)
							_ = Z7031
							__e.TailApply(PrimFunc(symshen_4_5body_6), Z7031)
							return
						}, 1)

						tmp7788 := Call(__e, PrimFunc(symcompile), tmp7787, V7025)

						__e.TailApply(tmp7758, tmp7788)
						return

					}, 1)

					tmp7789 := PrimCons(True, Nil)

					tmp7790 := PrimCons(symfreeze, tmp7789)

					__e.TailApply(tmp7757, tmp7790)
					return

				}, 1)

				__e.TailApply(tmp7756, MakeNumber(0))
				return

			}, 1)

			tmp7791 := PrimCons(MakeNumber(0), Nil)

			tmp7792 := PrimCons(symvector, tmp7791)

			tmp7793 := PrimCons(tmp7792, Nil)

			tmp7794 := PrimCons(MakeNumber(0), tmp7793)

			tmp7795 := PrimCons(True, tmp7794)

			tmp7796 := PrimCons(sym_8v, tmp7795)

			__e.TailApply(tmp7755, tmp7796)
			return

		}, 1)

		tmp7797 := PrimCons(symshen_4prolog_1vector, Nil)

		__e.TailApply(tmp7754, tmp7797)
		return

	}, 1)

	tmp7798 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp7753)

	_ = tmp7798

	tmp7799 := MakeNative(func(__e *ControlFlow) {
		V7040 := __e.Get(1)
		_ = V7040
		tmp7800 := MakeNative(func(__e *ControlFlow) {
			GoTo7041 := __e.Get(1)
			_ = GoTo7041
			tmp7817 := PrimIsPair(V7040)

			if True == tmp7817 {
				tmp7801 := MakeNative(func(__e *ControlFlow) {
					Select7042 := __e.Get(1)
					_ = Select7042
					tmp7802 := MakeNative(func(__e *ControlFlow) {
						Select7043 := __e.Get(1)
						_ = Select7043
						tmp7813 := PrimEqual(symreceive, Select7042)

						var ifres7806 Obj

						if True == tmp7813 {
							tmp7812 := PrimIsPair(Select7043)

							var ifres7808 Obj

							if True == tmp7812 {
								tmp7810 := PrimTail(Select7043)

								tmp7811 := PrimEqual(Nil, tmp7810)

								var ifres7809 Obj

								if True == tmp7811 {
									ifres7809 = True

								} else {
									ifres7809 = False

								}

								ifres7808 = ifres7809

							} else {
								ifres7808 = False

							}

							var ifres7807 Obj

							if True == ifres7808 {
								ifres7807 = True

							} else {
								ifres7807 = False

							}

							ifres7806 = ifres7807

						} else {
							ifres7806 = False

						}

						if True == ifres7806 {
							__e.Return(Select7043)
							return
						} else {
							tmp7803 := Call(__e, PrimFunc(symshen_4received), Select7042)

							tmp7804 := Call(__e, PrimFunc(symshen_4received), Select7043)

							__e.TailApply(PrimFunc(symunion), tmp7803, tmp7804)
							return

						}

					}, 1)

					tmp7814 := PrimTail(V7040)

					__e.TailApply(tmp7802, tmp7814)
					return

				}, 1)

				tmp7815 := PrimHead(V7040)

				__e.TailApply(tmp7801, tmp7815)
				return

			} else {
				__e.TailApply(PrimFunc(symthaw), GoTo7041)
				return
			}

		}, 1)

		tmp7818 := MakeNative(func(__e *ControlFlow) {
			__e.Return(Nil)
			return
		}, 0)

		__e.TailApply(tmp7800, tmp7818)
		return

	}, 1)

	tmp7819 := Call(__e, ns2_1set, symshen_4received, tmp7799)

	_ = tmp7819

	tmp7820 := MakeNative(func(__e *ControlFlow) {
		tmp7821 := MakeNative(func(__e *ControlFlow) {
			W7044 := __e.Get(1)
			_ = W7044
			tmp7822 := MakeNative(func(__e *ControlFlow) {
				W7045 := __e.Get(1)
				_ = W7045
				tmp7823 := MakeNative(func(__e *ControlFlow) {
					W7046 := __e.Get(1)
					_ = W7046
					__e.Return(W7046)
					return
				}, 1)

				tmp7824 := PrimVectorSet(W7044, MakeNumber(1), MakeNumber(2))

				__e.TailApply(tmp7823, tmp7824)
				return

			}, 1)

			tmp7825 := PrimVectorSet(W7044, MakeNumber(0), symshen_4print_1prolog_1vector)

			__e.TailApply(tmp7822, tmp7825)
			return

		}, 1)

		tmp7826 := PrimValue(symshen_4_dprolog_1memory_d)

		tmp7827 := PrimAbsvector(tmp7826)

		__e.TailApply(tmp7821, tmp7827)
		return

	}, 0)

	tmp7828 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp7820)

	_ = tmp7828

	tmp7829 := MakeNative(func(__e *ControlFlow) {
		V7047 := __e.Get(1)
		_ = V7047
		__e.Return(V7047)
		return
	}, 1)

	tmp7830 := Call(__e, ns2_1set, symreceive, tmp7829)

	_ = tmp7830

	tmp7831 := MakeNative(func(__e *ControlFlow) {
		V7048 := __e.Get(1)
		_ = V7048
		tmp7839 := PrimIsPair(V7048)

		if True == tmp7839 {
			tmp7832 := PrimHead(V7048)

			tmp7833 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7832)

			tmp7834 := PrimTail(V7048)

			tmp7835 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7834)

			tmp7836 := PrimCons(tmp7835, Nil)

			tmp7837 := PrimCons(tmp7833, tmp7836)

			__e.Return(PrimCons(symcons, tmp7837))
			return

		} else {
			__e.Return(V7048)
			return
		}

	}, 1)

	tmp7840 := Call(__e, ns2_1set, symshen_4rcons__form, tmp7831)

	_ = tmp7840

	tmp7841 := MakeNative(func(__e *ControlFlow) {
		V7049 := __e.Get(1)
		_ = V7049
		tmp7848 := PrimIsPair(V7049)

		if True == tmp7848 {
			tmp7842 := PrimHead(V7049)

			tmp7843 := PrimTail(V7049)

			tmp7844 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp7843)

			tmp7845 := PrimCons(tmp7844, Nil)

			tmp7846 := PrimCons(tmp7842, tmp7845)

			__e.Return(PrimCons(sym_8p, tmp7846))
			return

		} else {
			__e.Return(V7049)
			return
		}

	}, 1)

	tmp7849 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp7841)

	_ = tmp7849

	tmp7850 := MakeNative(func(__e *ControlFlow) {
		V7050 := __e.Get(1)
		_ = V7050
		tmp7851 := PrimValue(sym_dmacros_d)

		tmp7852 := Call(__e, PrimFunc(symassoc), V7050, tmp7851)

		tmp7853 := PrimValue(sym_dmacros_d)

		tmp7854 := Call(__e, PrimFunc(symremove), tmp7852, tmp7853)

		tmp7855 := PrimSet(sym_dmacros_d, tmp7854)

		_ = tmp7855

		__e.Return(V7050)
		return

	}, 1)

	__e.TailApply(ns2_1set, symundefmacro, tmp7850)
	return

}, 0)
