package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
	tmp7856 := MakeNative(func(__e *ControlFlow) {
		V1215 := __e.Get(1)
		_ = V1215
		tmp7857 := MakeNative(func(__e *ControlFlow) {
			W1216 := __e.Get(1)
			_ = W1216
			tmp7858 := MakeNative(func(__e *ControlFlow) {
				W1217 := __e.Get(1)
				_ = W1217
				tmp7859 := MakeNative(func(__e *ControlFlow) {
					W1223 := __e.Get(1)
					_ = W1223
					__e.Return(symloaded)
					return
				}, 1)

				var ifres7860 Obj

				if True == W1216 {
					tmp7861 := Call(__e, PrimFunc(syminferences))

					tmp7862 := Call(__e, PrimFunc(symshen_4app), tmp7861, MakeString(" inferences\n"), symshen_4a)

					tmp7863 := PrimStringConcat(MakeString("\ntypechecked in "), tmp7862)

					tmp7864 := Call(__e, PrimFunc(symstoutput))

					tmp7865 := Call(__e, PrimFunc(sympr), tmp7863, tmp7864)

					ifres7860 = tmp7865

				} else {
					ifres7860 = symshen_4skip

				}

				__e.TailApply(tmp7859, ifres7860)
				return

			}, 1)

			tmp7866 := MakeNative(func(__e *ControlFlow) {
				W1218 := __e.Get(1)
				_ = W1218
				tmp7867 := MakeNative(func(__e *ControlFlow) {
					W1219 := __e.Get(1)
					_ = W1219
					tmp7868 := MakeNative(func(__e *ControlFlow) {
						W1220 := __e.Get(1)
						_ = W1220
						tmp7869 := MakeNative(func(__e *ControlFlow) {
							W1221 := __e.Get(1)
							_ = W1221
							tmp7870 := MakeNative(func(__e *ControlFlow) {
								W1222 := __e.Get(1)
								_ = W1222
								__e.Return(W1219)
								return
							}, 1)

							tmp7871 := PrimStr(W1221)

							tmp7872 := PrimStringConcat(tmp7871, MakeString(" secs\n"))

							tmp7873 := PrimStringConcat(MakeString("\nrun time: "), tmp7872)

							tmp7874 := Call(__e, PrimFunc(symstoutput))

							tmp7875 := Call(__e, PrimFunc(sympr), tmp7873, tmp7874)

							__e.TailApply(tmp7870, tmp7875)
							return

						}, 1)

						tmp7876 := PrimNumberSubtract(W1220, W1218)

						__e.TailApply(tmp7869, tmp7876)
						return

					}, 1)

					tmp7877 := PrimGetTime(symrun)

					__e.TailApply(tmp7868, tmp7877)
					return

				}, 1)

				tmp7878 := Call(__e, PrimFunc(symread_1file), V1215)

				tmp7879 := Call(__e, PrimFunc(symshen_4load_1help), W1216, tmp7878)

				__e.TailApply(tmp7867, tmp7879)
				return

			}, 1)

			tmp7880 := PrimGetTime(symrun)

			tmp7881 := Call(__e, tmp7866, tmp7880)

			__e.TailApply(tmp7858, tmp7881)
			return

		}, 1)

		tmp7882 := PrimValue(symshen_4_dtc_d)

		__e.TailApply(tmp7857, tmp7882)
		return

	}, 1)

	tmp7883 := Call(__e, ns2_1set, symload, tmp7856)

	_ = tmp7883

	tmp7884 := MakeNative(func(__e *ControlFlow) {
		V1226 := __e.Get(1)
		_ = V1226
		V1227 := __e.Get(2)
		_ = V1227
		tmp7886 := PrimEqual(False, V1226)

		if True == tmp7886 {
			__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V1227)
			return
		} else {
			__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V1227)
			return
		}

	}, 2)

	tmp7887 := Call(__e, ns2_1set, symshen_4load_1help, tmp7884)

	_ = tmp7887

	tmp7888 := MakeNative(func(__e *ControlFlow) {
		V1228 := __e.Get(1)
		_ = V1228
		tmp7889 := MakeNative(func(__e *ControlFlow) {
			Z1229 := __e.Get(1)
			_ = Z1229
			tmp7890 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z1229)

			tmp7891 := Call(__e, PrimFunc(symeval_1kl), tmp7890)

			tmp7892 := Call(__e, PrimFunc(symshen_4app), tmp7891, MakeString("\n"), symshen_4s)

			tmp7893 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), tmp7892, tmp7893)
			return

		}, 1)

		__e.TailApply(PrimFunc(symshen_4for_1each), tmp7889, V1228)
		return

	}, 1)

	tmp7894 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp7888)

	_ = tmp7894

	tmp7895 := MakeNative(func(__e *ControlFlow) {
		V1230 := __e.Get(1)
		_ = V1230
		tmp7896 := MakeNative(func(__e *ControlFlow) {
			W1231 := __e.Get(1)
			_ = W1231
			tmp7897 := MakeNative(func(__e *ControlFlow) {
				W1233 := __e.Get(1)
				_ = W1233
				tmp7898 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimFunc(symshen_4work_1through), V1230)
					return
				}, 0)

				tmp7899 := MakeNative(func(__e *ControlFlow) {
					Z1235 := __e.Get(1)
					_ = Z1235
					__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1235, W1231)
					return
				}, 1)

				__e.TailApply(try_1catch, tmp7898, tmp7899)
				return

			}, 1)

			tmp7900 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimFunc(symshen_4assumetypes), W1231)
				return
			}, 0)

			tmp7901 := MakeNative(func(__e *ControlFlow) {
				Z1234 := __e.Get(1)
				_ = Z1234
				__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1234, W1231)
				return
			}, 1)

			tmp7902 := Call(__e, try_1catch, tmp7900, tmp7901)

			__e.TailApply(tmp7897, tmp7902)
			return

		}, 1)

		tmp7903 := MakeNative(func(__e *ControlFlow) {
			Z1232 := __e.Get(1)
			_ = Z1232
			__e.TailApply(PrimFunc(symshen_4typetable), Z1232)
			return
		}, 1)

		tmp7904 := Call(__e, PrimFunc(symmapcan), tmp7903, V1230)

		__e.TailApply(tmp7896, tmp7904)
		return

	}, 1)

	tmp7905 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp7895)

	_ = tmp7905

	tmp7906 := MakeNative(func(__e *ControlFlow) {
		V1240 := __e.Get(1)
		_ = V1240
		tmp7951 := PrimIsPair(V1240)

		var ifres7932 Obj

		if True == tmp7951 {
			tmp7949 := PrimHead(V1240)

			tmp7950 := PrimEqual(symdefine, tmp7949)

			var ifres7934 Obj

			if True == tmp7950 {
				tmp7947 := PrimTail(V1240)

				tmp7948 := PrimIsPair(tmp7947)

				var ifres7936 Obj

				if True == tmp7948 {
					tmp7944 := PrimTail(V1240)

					tmp7945 := PrimTail(tmp7944)

					tmp7946 := PrimIsPair(tmp7945)

					var ifres7938 Obj

					if True == tmp7946 {
						tmp7940 := PrimTail(V1240)

						tmp7941 := PrimTail(tmp7940)

						tmp7942 := PrimHead(tmp7941)

						tmp7943 := PrimEqual(sym_i, tmp7942)

						var ifres7939 Obj

						if True == tmp7943 {
							ifres7939 = True

						} else {
							ifres7939 = False

						}

						ifres7938 = ifres7939

					} else {
						ifres7938 = False

					}

					var ifres7937 Obj

					if True == ifres7938 {
						ifres7937 = True

					} else {
						ifres7937 = False

					}

					ifres7936 = ifres7937

				} else {
					ifres7936 = False

				}

				var ifres7935 Obj

				if True == ifres7936 {
					ifres7935 = True

				} else {
					ifres7935 = False

				}

				ifres7934 = ifres7935

			} else {
				ifres7934 = False

			}

			var ifres7933 Obj

			if True == ifres7934 {
				ifres7933 = True

			} else {
				ifres7933 = False

			}

			ifres7932 = ifres7933

		} else {
			ifres7932 = False

		}

		if True == ifres7932 {
			tmp7907 := PrimTail(V1240)

			tmp7908 := PrimHead(tmp7907)

			tmp7909 := PrimTail(V1240)

			tmp7910 := PrimHead(tmp7909)

			tmp7911 := PrimTail(V1240)

			tmp7912 := PrimTail(tmp7911)

			tmp7913 := PrimTail(tmp7912)

			tmp7914 := Call(__e, PrimFunc(symshen_4type_1F), tmp7910, tmp7913)

			tmp7915 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp7914)

			tmp7916 := PrimCons(tmp7915, Nil)

			__e.Return(PrimCons(tmp7908, tmp7916))
			return

		} else {
			tmp7930 := PrimIsPair(V1240)

			var ifres7922 Obj

			if True == tmp7930 {
				tmp7928 := PrimHead(V1240)

				tmp7929 := PrimEqual(symdefine, tmp7928)

				var ifres7924 Obj

				if True == tmp7929 {
					tmp7926 := PrimTail(V1240)

					tmp7927 := PrimIsPair(tmp7926)

					var ifres7925 Obj

					if True == tmp7927 {
						ifres7925 = True

					} else {
						ifres7925 = False

					}

					ifres7924 = ifres7925

				} else {
					ifres7924 = False

				}

				var ifres7923 Obj

				if True == ifres7924 {
					ifres7923 = True

				} else {
					ifres7923 = False

				}

				ifres7922 = ifres7923

			} else {
				ifres7922 = False

			}

			if True == ifres7922 {
				tmp7917 := PrimTail(V1240)

				tmp7918 := PrimHead(tmp7917)

				tmp7919 := Call(__e, PrimFunc(symshen_4app), tmp7918, MakeString("\n"), symshen_4a)

				tmp7920 := PrimStringConcat(MakeString("missing { in "), tmp7919)

				__e.Return(PrimSimpleError(tmp7920))
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp7952 := Call(__e, ns2_1set, symshen_4typetable, tmp7906)

	_ = tmp7952

	tmp7953 := MakeNative(func(__e *ControlFlow) {
		V1247 := __e.Get(1)
		_ = V1247
		V1248 := __e.Get(2)
		_ = V1248
		tmp7966 := PrimIsPair(V1248)

		var ifres7962 Obj

		if True == tmp7966 {
			tmp7964 := PrimHead(V1248)

			tmp7965 := PrimEqual(sym_j, tmp7964)

			var ifres7963 Obj

			if True == tmp7965 {
				ifres7963 = True

			} else {
				ifres7963 = False

			}

			ifres7962 = ifres7963

		} else {
			ifres7962 = False

		}

		if True == ifres7962 {
			__e.Return(Nil)
			return
		} else {
			tmp7960 := PrimIsPair(V1248)

			if True == tmp7960 {
				tmp7954 := PrimHead(V1248)

				tmp7955 := PrimTail(V1248)

				tmp7956 := Call(__e, PrimFunc(symshen_4type_1F), V1247, tmp7955)

				__e.Return(PrimCons(tmp7954, tmp7956))
				return

			} else {
				tmp7957 := Call(__e, PrimFunc(symshen_4app), V1247, MakeString("\n"), symshen_4a)

				tmp7958 := PrimStringConcat(MakeString("missing } in "), tmp7957)

				__e.Return(PrimSimpleError(tmp7958))
				return

			}

		}

	}, 2)

	tmp7967 := Call(__e, ns2_1set, symshen_4type_1F, tmp7953)

	_ = tmp7967

	tmp7968 := MakeNative(func(__e *ControlFlow) {
		V1251 := __e.Get(1)
		_ = V1251
		tmp7982 := PrimEqual(Nil, V1251)

		if True == tmp7982 {
			__e.Return(Nil)
			return
		} else {
			tmp7980 := PrimIsPair(V1251)

			var ifres7976 Obj

			if True == tmp7980 {
				tmp7978 := PrimTail(V1251)

				tmp7979 := PrimIsPair(tmp7978)

				var ifres7977 Obj

				if True == tmp7979 {
					ifres7977 = True

				} else {
					ifres7977 = False

				}

				ifres7976 = ifres7977

			} else {
				ifres7976 = False

			}

			if True == ifres7976 {
				tmp7969 := PrimHead(V1251)

				tmp7970 := PrimTail(V1251)

				tmp7971 := PrimHead(tmp7970)

				tmp7972 := Call(__e, PrimFunc(symdeclare), tmp7969, tmp7971)

				_ = tmp7972

				tmp7973 := PrimTail(V1251)

				tmp7974 := PrimTail(tmp7973)

				__e.TailApply(PrimFunc(symshen_4assumetypes), tmp7974)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
				return
			}

		}

	}, 1)

	tmp7983 := Call(__e, ns2_1set, symshen_4assumetypes, tmp7968)

	_ = tmp7983

	tmp7984 := MakeNative(func(__e *ControlFlow) {
		V1256 := __e.Get(1)
		_ = V1256
		V1257 := __e.Get(2)
		_ = V1257
		tmp7995 := PrimIsPair(V1257)

		var ifres7991 Obj

		if True == tmp7995 {
			tmp7993 := PrimTail(V1257)

			tmp7994 := PrimIsPair(tmp7993)

			var ifres7992 Obj

			if True == tmp7994 {
				ifres7992 = True

			} else {
				ifres7992 = False

			}

			ifres7991 = ifres7992

		} else {
			ifres7991 = False

		}

		if True == ifres7991 {
			tmp7985 := PrimHead(V1257)

			tmp7986 := Call(__e, PrimFunc(symdestroy), tmp7985)

			_ = tmp7986

			tmp7987 := PrimTail(V1257)

			tmp7988 := PrimTail(tmp7987)

			__e.TailApply(PrimFunc(symshen_4unwind_1types), V1256, tmp7988)
			return

		} else {
			tmp7989 := PrimErrorToString(V1256)

			__e.Return(PrimSimpleError(tmp7989))
			return

		}

	}, 2)

	tmp7996 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp7984)

	_ = tmp7996

	tmp7997 := MakeNative(func(__e *ControlFlow) {
		V1260 := __e.Get(1)
		_ = V1260
		tmp8046 := PrimEqual(Nil, V1260)

		if True == tmp8046 {
			__e.Return(Nil)
			return
		} else {
			tmp8044 := PrimIsPair(V1260)

			var ifres8029 Obj

			if True == tmp8044 {
				tmp8042 := PrimTail(V1260)

				tmp8043 := PrimIsPair(tmp8042)

				var ifres8031 Obj

				if True == tmp8043 {
					tmp8039 := PrimTail(V1260)

					tmp8040 := PrimTail(tmp8039)

					tmp8041 := PrimIsPair(tmp8040)

					var ifres8033 Obj

					if True == tmp8041 {
						tmp8035 := PrimTail(V1260)

						tmp8036 := PrimHead(tmp8035)

						tmp8037 := PrimIntern(MakeString(":"))

						tmp8038 := PrimEqual(tmp8036, tmp8037)

						var ifres8034 Obj

						if True == tmp8038 {
							ifres8034 = True

						} else {
							ifres8034 = False

						}

						ifres8033 = ifres8034

					} else {
						ifres8033 = False

					}

					var ifres8032 Obj

					if True == ifres8033 {
						ifres8032 = True

					} else {
						ifres8032 = False

					}

					ifres8031 = ifres8032

				} else {
					ifres8031 = False

				}

				var ifres8030 Obj

				if True == ifres8031 {
					ifres8030 = True

				} else {
					ifres8030 = False

				}

				ifres8029 = ifres8030

			} else {
				ifres8029 = False

			}

			if True == ifres8029 {
				tmp7998 := MakeNative(func(__e *ControlFlow) {
					W1261 := __e.Get(1)
					_ = W1261
					tmp8014 := PrimEqual(W1261, False)

					if True == tmp8014 {
						__e.TailApply(PrimFunc(symshen_4type_1error))
						return
					} else {
						tmp7999 := MakeNative(func(__e *ControlFlow) {
							W1262 := __e.Get(1)
							_ = W1262
							tmp8000 := MakeNative(func(__e *ControlFlow) {
								W1263 := __e.Get(1)
								_ = W1263
								tmp8001 := PrimTail(V1260)

								tmp8002 := PrimTail(tmp8001)

								tmp8003 := PrimTail(tmp8002)

								__e.TailApply(PrimFunc(symshen_4work_1through), tmp8003)
								return

							}, 1)

							tmp8004 := Call(__e, PrimFunc(symshen_4pretty_1type), W1261)

							tmp8005 := Call(__e, PrimFunc(symshen_4app), tmp8004, MakeString("\n"), symshen_4r)

							tmp8006 := PrimStringConcat(MakeString(" : "), tmp8005)

							tmp8007 := Call(__e, PrimFunc(symshen_4app), W1262, tmp8006, symshen_4s)

							tmp8008 := Call(__e, PrimFunc(symstoutput))

							tmp8009 := Call(__e, PrimFunc(sympr), tmp8007, tmp8008)

							__e.TailApply(tmp8000, tmp8009)
							return

						}, 1)

						tmp8010 := PrimHead(V1260)

						tmp8011 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp8010)

						tmp8012 := Call(__e, PrimFunc(symeval_1kl), tmp8011)

						__e.TailApply(tmp7999, tmp8012)
						return

					}

				}, 1)

				tmp8015 := PrimHead(V1260)

				tmp8016 := PrimTail(V1260)

				tmp8017 := PrimTail(tmp8016)

				tmp8018 := PrimHead(tmp8017)

				tmp8019 := Call(__e, PrimFunc(symshen_4typecheck), tmp8015, tmp8018)

				__e.TailApply(tmp7998, tmp8019)
				return

			} else {
				tmp8027 := PrimIsPair(V1260)

				if True == tmp8027 {
					tmp8020 := PrimHead(V1260)

					tmp8021 := PrimIntern(MakeString(":"))

					tmp8022 := PrimTail(V1260)

					tmp8023 := PrimCons(symA, tmp8022)

					tmp8024 := PrimCons(tmp8021, tmp8023)

					tmp8025 := PrimCons(tmp8020, tmp8024)

					__e.TailApply(PrimFunc(symshen_4work_1through), tmp8025)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
					return
				}

			}

		}

	}, 1)

	tmp8047 := Call(__e, ns2_1set, symshen_4work_1through, tmp7997)

	_ = tmp8047

	tmp8048 := MakeNative(func(__e *ControlFlow) {
		V1265 := __e.Get(1)
		_ = V1265
		tmp8190 := PrimIsPair(V1265)

		var ifres8064 Obj

		if True == tmp8190 {
			tmp8188 := PrimHead(V1265)

			tmp8189 := PrimIsPair(tmp8188)

			var ifres8066 Obj

			if True == tmp8189 {
				tmp8185 := PrimHead(V1265)

				tmp8186 := PrimHead(tmp8185)

				tmp8187 := PrimEqual(symlist, tmp8186)

				var ifres8068 Obj

				if True == tmp8187 {
					tmp8182 := PrimHead(V1265)

					tmp8183 := PrimTail(tmp8182)

					tmp8184 := PrimIsPair(tmp8183)

					var ifres8070 Obj

					if True == tmp8184 {
						tmp8178 := PrimHead(V1265)

						tmp8179 := PrimTail(tmp8178)

						tmp8180 := PrimTail(tmp8179)

						tmp8181 := PrimEqual(Nil, tmp8180)

						var ifres8072 Obj

						if True == tmp8181 {
							tmp8176 := PrimTail(V1265)

							tmp8177 := PrimIsPair(tmp8176)

							var ifres8074 Obj

							if True == tmp8177 {
								tmp8173 := PrimTail(V1265)

								tmp8174 := PrimHead(tmp8173)

								tmp8175 := PrimEqual(sym_1_1_6, tmp8174)

								var ifres8076 Obj

								if True == tmp8175 {
									tmp8170 := PrimTail(V1265)

									tmp8171 := PrimTail(tmp8170)

									tmp8172 := PrimIsPair(tmp8171)

									var ifres8078 Obj

									if True == tmp8172 {
										tmp8166 := PrimTail(V1265)

										tmp8167 := PrimTail(tmp8166)

										tmp8168 := PrimHead(tmp8167)

										tmp8169 := PrimIsPair(tmp8168)

										var ifres8080 Obj

										if True == tmp8169 {
											tmp8161 := PrimTail(V1265)

											tmp8162 := PrimTail(tmp8161)

											tmp8163 := PrimHead(tmp8162)

											tmp8164 := PrimHead(tmp8163)

											tmp8165 := PrimEqual(symstr, tmp8164)

											var ifres8082 Obj

											if True == tmp8165 {
												tmp8156 := PrimTail(V1265)

												tmp8157 := PrimTail(tmp8156)

												tmp8158 := PrimHead(tmp8157)

												tmp8159 := PrimTail(tmp8158)

												tmp8160 := PrimIsPair(tmp8159)

												var ifres8084 Obj

												if True == tmp8160 {
													tmp8150 := PrimTail(V1265)

													tmp8151 := PrimTail(tmp8150)

													tmp8152 := PrimHead(tmp8151)

													tmp8153 := PrimTail(tmp8152)

													tmp8154 := PrimHead(tmp8153)

													tmp8155 := PrimIsPair(tmp8154)

													var ifres8086 Obj

													if True == tmp8155 {
														tmp8143 := PrimTail(V1265)

														tmp8144 := PrimTail(tmp8143)

														tmp8145 := PrimHead(tmp8144)

														tmp8146 := PrimTail(tmp8145)

														tmp8147 := PrimHead(tmp8146)

														tmp8148 := PrimHead(tmp8147)

														tmp8149 := PrimEqual(symlist, tmp8148)

														var ifres8088 Obj

														if True == tmp8149 {
															tmp8136 := PrimTail(V1265)

															tmp8137 := PrimTail(tmp8136)

															tmp8138 := PrimHead(tmp8137)

															tmp8139 := PrimTail(tmp8138)

															tmp8140 := PrimHead(tmp8139)

															tmp8141 := PrimTail(tmp8140)

															tmp8142 := PrimIsPair(tmp8141)

															var ifres8090 Obj

															if True == tmp8142 {
																tmp8128 := PrimTail(V1265)

																tmp8129 := PrimTail(tmp8128)

																tmp8130 := PrimHead(tmp8129)

																tmp8131 := PrimTail(tmp8130)

																tmp8132 := PrimHead(tmp8131)

																tmp8133 := PrimTail(tmp8132)

																tmp8134 := PrimTail(tmp8133)

																tmp8135 := PrimEqual(Nil, tmp8134)

																var ifres8092 Obj

																if True == tmp8135 {
																	tmp8122 := PrimTail(V1265)

																	tmp8123 := PrimTail(tmp8122)

																	tmp8124 := PrimHead(tmp8123)

																	tmp8125 := PrimTail(tmp8124)

																	tmp8126 := PrimTail(tmp8125)

																	tmp8127 := PrimIsPair(tmp8126)

																	var ifres8094 Obj

																	if True == tmp8127 {
																		tmp8115 := PrimTail(V1265)

																		tmp8116 := PrimTail(tmp8115)

																		tmp8117 := PrimHead(tmp8116)

																		tmp8118 := PrimTail(tmp8117)

																		tmp8119 := PrimTail(tmp8118)

																		tmp8120 := PrimTail(tmp8119)

																		tmp8121 := PrimEqual(Nil, tmp8120)

																		var ifres8096 Obj

																		if True == tmp8121 {
																			tmp8111 := PrimTail(V1265)

																			tmp8112 := PrimTail(tmp8111)

																			tmp8113 := PrimTail(tmp8112)

																			tmp8114 := PrimEqual(Nil, tmp8113)

																			var ifres8098 Obj

																			if True == tmp8114 {
																				tmp8100 := PrimHead(V1265)

																				tmp8101 := PrimTail(tmp8100)

																				tmp8102 := PrimHead(tmp8101)

																				tmp8103 := PrimTail(V1265)

																				tmp8104 := PrimTail(tmp8103)

																				tmp8105 := PrimHead(tmp8104)

																				tmp8106 := PrimTail(tmp8105)

																				tmp8107 := PrimHead(tmp8106)

																				tmp8108 := PrimTail(tmp8107)

																				tmp8109 := PrimHead(tmp8108)

																				tmp8110 := PrimEqual(tmp8102, tmp8109)

																				var ifres8099 Obj

																				if True == tmp8110 {
																					ifres8099 = True

																				} else {
																					ifres8099 = False

																				}

																				ifres8098 = ifres8099

																			} else {
																				ifres8098 = False

																			}

																			var ifres8097 Obj

																			if True == ifres8098 {
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

															var ifres8089 Obj

															if True == ifres8090 {
																ifres8089 = True

															} else {
																ifres8089 = False

															}

															ifres8088 = ifres8089

														} else {
															ifres8088 = False

														}

														var ifres8087 Obj

														if True == ifres8088 {
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

												var ifres8083 Obj

												if True == ifres8084 {
													ifres8083 = True

												} else {
													ifres8083 = False

												}

												ifres8082 = ifres8083

											} else {
												ifres8082 = False

											}

											var ifres8081 Obj

											if True == ifres8082 {
												ifres8081 = True

											} else {
												ifres8081 = False

											}

											ifres8080 = ifres8081

										} else {
											ifres8080 = False

										}

										var ifres8079 Obj

										if True == ifres8080 {
											ifres8079 = True

										} else {
											ifres8079 = False

										}

										ifres8078 = ifres8079

									} else {
										ifres8078 = False

									}

									var ifres8077 Obj

									if True == ifres8078 {
										ifres8077 = True

									} else {
										ifres8077 = False

									}

									ifres8076 = ifres8077

								} else {
									ifres8076 = False

								}

								var ifres8075 Obj

								if True == ifres8076 {
									ifres8075 = True

								} else {
									ifres8075 = False

								}

								ifres8074 = ifres8075

							} else {
								ifres8074 = False

							}

							var ifres8073 Obj

							if True == ifres8074 {
								ifres8073 = True

							} else {
								ifres8073 = False

							}

							ifres8072 = ifres8073

						} else {
							ifres8072 = False

						}

						var ifres8071 Obj

						if True == ifres8072 {
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

				var ifres8067 Obj

				if True == ifres8068 {
					ifres8067 = True

				} else {
					ifres8067 = False

				}

				ifres8066 = ifres8067

			} else {
				ifres8066 = False

			}

			var ifres8065 Obj

			if True == ifres8066 {
				ifres8065 = True

			} else {
				ifres8065 = False

			}

			ifres8064 = ifres8065

		} else {
			ifres8064 = False

		}

		if True == ifres8064 {
			tmp8049 := PrimTail(V1265)

			tmp8050 := PrimTail(tmp8049)

			tmp8051 := PrimHead(tmp8050)

			tmp8052 := PrimTail(tmp8051)

			tmp8053 := PrimHead(tmp8052)

			tmp8054 := PrimTail(V1265)

			tmp8055 := PrimTail(tmp8054)

			tmp8056 := PrimHead(tmp8055)

			tmp8057 := PrimTail(tmp8056)

			tmp8058 := PrimTail(tmp8057)

			tmp8059 := PrimCons(sym_a_a_6, tmp8058)

			__e.Return(PrimCons(tmp8053, tmp8059))
			return

		} else {
			tmp8062 := PrimIsPair(V1265)

			if True == tmp8062 {
				tmp8060 := MakeNative(func(__e *ControlFlow) {
					Z1266 := __e.Get(1)
					_ = Z1266
					__e.TailApply(PrimFunc(symshen_4pretty_1type), Z1266)
					return
				}, 1)

				__e.TailApply(PrimFunc(symmap), tmp8060, V1265)
				return

			} else {
				__e.Return(V1265)
				return
			}

		}

	}, 1)

	tmp8191 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp8048)

	_ = tmp8191

	tmp8192 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("type error\n")))
		return
	}, 0)

	tmp8193 := Call(__e, ns2_1set, symshen_4type_1error, tmp8192)

	_ = tmp8193

	tmp8194 := MakeNative(func(__e *ControlFlow) {
		V1267 := __e.Get(1)
		_ = V1267
		tmp8195 := MakeNative(func(__e *ControlFlow) {
			W1268 := __e.Get(1)
			_ = W1268
			tmp8196 := MakeNative(func(__e *ControlFlow) {
				W1269 := __e.Get(1)
				_ = W1269
				tmp8197 := MakeNative(func(__e *ControlFlow) {
					W1270 := __e.Get(1)
					_ = W1270
					tmp8198 := MakeNative(func(__e *ControlFlow) {
						W1271 := __e.Get(1)
						_ = W1271
						tmp8199 := MakeNative(func(__e *ControlFlow) {
							W1273 := __e.Get(1)
							_ = W1273
							__e.Return(W1268)
							return
						}, 1)

						tmp8200 := Call(__e, PrimFunc(symshen_4write_1kl), W1271, W1270)

						__e.TailApply(tmp8199, tmp8200)
						return

					}, 1)

					tmp8201 := MakeNative(func(__e *ControlFlow) {
						Z1272 := __e.Get(1)
						_ = Z1272
						tmp8202 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z1272)

						__e.TailApply(PrimFunc(symshen_4partial), tmp8202)
						return

					}, 1)

					tmp8203 := Call(__e, PrimFunc(symmap), tmp8201, W1269)

					__e.TailApply(tmp8198, tmp8203)
					return

				}, 1)

				tmp8204 := PrimOpenStream(W1268, symout)

				__e.TailApply(tmp8197, tmp8204)
				return

			}, 1)

			tmp8205 := Call(__e, PrimFunc(symread_1file), V1267)

			__e.TailApply(tmp8196, tmp8205)
			return

		}, 1)

		tmp8206 := Call(__e, PrimFunc(symshen_4klfile), V1267)

		__e.TailApply(tmp8195, tmp8206)
		return

	}, 1)

	tmp8207 := Call(__e, ns2_1set, symbootstrap, tmp8194)

	_ = tmp8207

	tmp8208 := MakeNative(func(__e *ControlFlow) {
		V1274 := __e.Get(1)
		_ = V1274
		tmp8231 := PrimIsPair(V1274)

		var ifres8218 Obj

		if True == tmp8231 {
			tmp8229 := PrimHead(V1274)

			tmp8230 := PrimEqual(symshen_4f_1error, tmp8229)

			var ifres8220 Obj

			if True == tmp8230 {
				tmp8227 := PrimTail(V1274)

				tmp8228 := PrimIsPair(tmp8227)

				var ifres8222 Obj

				if True == tmp8228 {
					tmp8224 := PrimTail(V1274)

					tmp8225 := PrimTail(tmp8224)

					tmp8226 := PrimEqual(Nil, tmp8225)

					var ifres8223 Obj

					if True == tmp8226 {
						ifres8223 = True

					} else {
						ifres8223 = False

					}

					ifres8222 = ifres8223

				} else {
					ifres8222 = False

				}

				var ifres8221 Obj

				if True == ifres8222 {
					ifres8221 = True

				} else {
					ifres8221 = False

				}

				ifres8220 = ifres8221

			} else {
				ifres8220 = False

			}

			var ifres8219 Obj

			if True == ifres8220 {
				ifres8219 = True

			} else {
				ifres8219 = False

			}

			ifres8218 = ifres8219

		} else {
			ifres8218 = False

		}

		if True == ifres8218 {
			tmp8209 := PrimTail(V1274)

			tmp8210 := PrimHead(tmp8209)

			tmp8211 := PrimStr(tmp8210)

			tmp8212 := PrimStringConcat(MakeString("partial function "), tmp8211)

			tmp8213 := PrimCons(tmp8212, Nil)

			__e.Return(PrimCons(symsimple_1error, tmp8213))
			return

		} else {
			tmp8216 := PrimIsPair(V1274)

			if True == tmp8216 {
				tmp8214 := MakeNative(func(__e *ControlFlow) {
					Z1275 := __e.Get(1)
					_ = Z1275
					__e.TailApply(PrimFunc(symshen_4partial), Z1275)
					return
				}, 1)

				__e.TailApply(PrimFunc(symmap), tmp8214, V1274)
				return

			} else {
				__e.Return(V1274)
				return
			}

		}

	}, 1)

	tmp8232 := Call(__e, ns2_1set, symshen_4partial, tmp8208)

	_ = tmp8232

	tmp8233 := MakeNative(func(__e *ControlFlow) {
		V1278 := __e.Get(1)
		_ = V1278
		V1279 := __e.Get(2)
		_ = V1279
		tmp8247 := PrimEqual(Nil, V1278)

		if True == tmp8247 {
			__e.Return(PrimCloseStream(V1279))
			return
		} else {
			tmp8245 := PrimIsPair(V1278)

			var ifres8241 Obj

			if True == tmp8245 {
				tmp8243 := PrimHead(V1278)

				tmp8244 := PrimIsPair(tmp8243)

				var ifres8242 Obj

				if True == tmp8244 {
					ifres8242 = True

				} else {
					ifres8242 = False

				}

				ifres8241 = ifres8242

			} else {
				ifres8241 = False

			}

			if True == ifres8241 {
				tmp8234 := PrimTail(V1278)

				tmp8235 := PrimHead(V1278)

				tmp8236 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp8235, V1279)

				_ = tmp8236

				__e.TailApply(PrimFunc(symshen_4write_1kl), tmp8234, V1279)
				return

			} else {
				tmp8239 := PrimIsPair(V1278)

				if True == tmp8239 {
					tmp8237 := PrimTail(V1278)

					__e.TailApply(PrimFunc(symshen_4write_1kl), tmp8237, V1279)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4write_1kl)
					return
				}

			}

		}

	}, 2)

	tmp8248 := Call(__e, ns2_1set, symshen_4write_1kl, tmp8233)

	_ = tmp8248

	tmp8249 := MakeNative(func(__e *ControlFlow) {
		V1282 := __e.Get(1)
		_ = V1282
		V1283 := __e.Get(2)
		_ = V1283
		tmp8289 := PrimIsPair(V1282)

		var ifres8252 Obj

		if True == tmp8289 {
			tmp8287 := PrimHead(V1282)

			tmp8288 := PrimEqual(symdefun, tmp8287)

			var ifres8254 Obj

			if True == tmp8288 {
				tmp8285 := PrimTail(V1282)

				tmp8286 := PrimIsPair(tmp8285)

				var ifres8256 Obj

				if True == tmp8286 {
					tmp8282 := PrimTail(V1282)

					tmp8283 := PrimHead(tmp8282)

					tmp8284 := PrimEqual(symfail, tmp8283)

					var ifres8258 Obj

					if True == tmp8284 {
						tmp8279 := PrimTail(V1282)

						tmp8280 := PrimTail(tmp8279)

						tmp8281 := PrimIsPair(tmp8280)

						var ifres8260 Obj

						if True == tmp8281 {
							tmp8275 := PrimTail(V1282)

							tmp8276 := PrimTail(tmp8275)

							tmp8277 := PrimHead(tmp8276)

							tmp8278 := PrimEqual(Nil, tmp8277)

							var ifres8262 Obj

							if True == tmp8278 {
								tmp8271 := PrimTail(V1282)

								tmp8272 := PrimTail(tmp8271)

								tmp8273 := PrimTail(tmp8272)

								tmp8274 := PrimIsPair(tmp8273)

								var ifres8264 Obj

								if True == tmp8274 {
									tmp8266 := PrimTail(V1282)

									tmp8267 := PrimTail(tmp8266)

									tmp8268 := PrimTail(tmp8267)

									tmp8269 := PrimTail(tmp8268)

									tmp8270 := PrimEqual(Nil, tmp8269)

									var ifres8265 Obj

									if True == tmp8270 {
										ifres8265 = True

									} else {
										ifres8265 = False

									}

									ifres8264 = ifres8265

								} else {
									ifres8264 = False

								}

								var ifres8263 Obj

								if True == ifres8264 {
									ifres8263 = True

								} else {
									ifres8263 = False

								}

								ifres8262 = ifres8263

							} else {
								ifres8262 = False

							}

							var ifres8261 Obj

							if True == ifres8262 {
								ifres8261 = True

							} else {
								ifres8261 = False

							}

							ifres8260 = ifres8261

						} else {
							ifres8260 = False

						}

						var ifres8259 Obj

						if True == ifres8260 {
							ifres8259 = True

						} else {
							ifres8259 = False

						}

						ifres8258 = ifres8259

					} else {
						ifres8258 = False

					}

					var ifres8257 Obj

					if True == ifres8258 {
						ifres8257 = True

					} else {
						ifres8257 = False

					}

					ifres8256 = ifres8257

				} else {
					ifres8256 = False

				}

				var ifres8255 Obj

				if True == ifres8256 {
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

		if True == ifres8252 {
			__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V1283)
			return
		} else {
			tmp8250 := Call(__e, PrimFunc(symshen_4app), V1282, MakeString("\n\n"), symshen_4r)

			__e.TailApply(PrimFunc(sympr), tmp8250, V1283)
			return

		}

	}, 2)

	tmp8290 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp8249)

	_ = tmp8290

	tmp8291 := MakeNative(func(__e *ControlFlow) {
		V1284 := __e.Get(1)
		_ = V1284
		tmp8300 := PrimEqual(MakeString(""), V1284)

		if True == tmp8300 {
			__e.Return(MakeString(".kl"))
			return
		} else {
			tmp8298 := PrimEqual(MakeString(".shen"), V1284)

			if True == tmp8298 {
				__e.Return(MakeString(".kl"))
				return
			} else {
				tmp8296 := Call(__e, PrimFunc(symshen_4_7string_2), V1284)

				if True == tmp8296 {
					tmp8292 := Call(__e, PrimFunc(symhdstr), V1284)

					tmp8293 := PrimTailString(V1284)

					tmp8294 := Call(__e, PrimFunc(symshen_4klfile), tmp8293)

					__e.TailApply(PrimFunc(sym_8s), tmp8292, tmp8294)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4klfile)
					return
				}

			}

		}

	}, 1)

	__e.TailApply(ns2_1set, symshen_4klfile, tmp8291)
	return

}, 0)
