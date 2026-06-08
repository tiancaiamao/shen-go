package main

import . "github.com/tiancaiamao/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
	tmp968 := MakeNative(func(__e *ControlFlow) {
		V6779 := __e.Get(1)
		_ = V6779
		tmp969 := MakeNative(func(__e *ControlFlow) {
			W6780 := __e.Get(1)
			_ = W6780
			tmp970 := MakeNative(func(__e *ControlFlow) {
				W6781 := __e.Get(1)
				_ = W6781
				__e.Return(V6779)
				return
			}, 1)

			tmp971 := Call(__e, PrimFunc(symstoutput))

			tmp972 := Call(__e, PrimFunc(sympr), W6780, tmp971)

			__e.TailApply(tmp970, tmp972)
			return

		}, 1)

		tmp973 := Call(__e, PrimFunc(symshen_4insert), V6779, MakeString("~S"))

		__e.TailApply(tmp969, tmp973)
		return

	}, 1)

	tmp974 := Call(__e, ns2_1set, symprint, tmp968)

	_ = tmp974

	tmp975 := MakeNative(func(__e *ControlFlow) {
		V6782 := __e.Get(1)
		_ = V6782
		V6783 := __e.Get(2)
		_ = V6783
		tmp980 := PrimValue(sym_dhush_d)

		if True == tmp980 {
			__e.Return(V6782)
			return
		} else {
			tmp978 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V6783)

			if True == tmp978 {
				__e.TailApply(PrimFunc(symshen_4write_1string), V6782, V6783)
				return
			} else {
				tmp976 := Call(__e, PrimFunc(symshen_4string_1_6byte), V6782, MakeNumber(0))

				__e.TailApply(PrimFunc(symshen_4write_1chars), V6782, V6783, tmp976, MakeNumber(1))
				return

			}

		}

	}, 2)

	tmp981 := Call(__e, ns2_1set, sympr, tmp975)

	_ = tmp981

	tmp982 := MakeNative(func(__e *ControlFlow) {
		V6784 := __e.Get(1)
		_ = V6784
		V6785 := __e.Get(2)
		_ = V6785
		tmp983 := MakeNative(func(__e *ControlFlow) {
			tmp984 := PrimPos(V6784, V6785)

			__e.Return(PrimStringToNumber(tmp984))
			return

		}, 0)

		tmp985 := MakeNative(func(__e *ControlFlow) {
			Z6786 := __e.Get(1)
			_ = Z6786
			__e.Return(symshen_4eos)
			return
		}, 1)

		__e.TailApply(try_1catch, tmp983, tmp985)
		return

	}, 2)

	tmp986 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp982)

	_ = tmp986

	tmp987 := MakeNative(func(__e *ControlFlow) {
		V6787 := __e.Get(1)
		_ = V6787
		V6788 := __e.Get(2)
		_ = V6788
		V6789 := __e.Get(3)
		_ = V6789
		V6790 := __e.Get(4)
		_ = V6790
		tmp992 := PrimEqual(symshen_4eos, V6789)

		if True == tmp992 {
			__e.Return(V6787)
			return
		} else {
			tmp988 := PrimWriteByte(V6789, V6788)

			_ = tmp988

			tmp989 := Call(__e, PrimFunc(symshen_4string_1_6byte), V6787, V6790)

			tmp990 := PrimNumberAdd(V6790, MakeNumber(1))

			__e.TailApply(PrimFunc(symshen_4write_1chars), V6787, V6788, tmp989, tmp990)
			return

		}

	}, 4)

	tmp993 := Call(__e, ns2_1set, symshen_4write_1chars, tmp987)

	_ = tmp993

	tmp994 := MakeNative(func(__e *ControlFlow) {
		V6791 := __e.Get(1)
		_ = V6791
		V6792 := __e.Get(2)
		_ = V6792
		tmp999 := PrimIsString(V6791)

		if True == tmp999 {
			tmp995 := Call(__e, PrimFunc(symshen_4proc_1nl), V6791)

			__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp995, V6792)
			return

		} else {
			tmp996 := PrimCons(V6791, Nil)

			tmp997 := PrimCons(symshen_4proc_1nl, tmp996)

			__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp997, V6792)
			return

		}

	}, 2)

	tmp1000 := Call(__e, ns2_1set, symshen_4mkstr, tmp994)

	_ = tmp1000

	tmp1001 := MakeNative(func(__e *ControlFlow) {
		V6797 := __e.Get(1)
		_ = V6797
		V6798 := __e.Get(2)
		_ = V6798
		tmp1008 := PrimEqual(Nil, V6798)

		if True == tmp1008 {
			__e.Return(V6797)
			return
		} else {
			tmp1006 := PrimIsPair(V6798)

			if True == tmp1006 {
				tmp1002 := PrimHead(V6798)

				tmp1003 := Call(__e, PrimFunc(symshen_4insert_1l), tmp1002, V6797)

				tmp1004 := PrimTail(V6798)

				__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp1003, tmp1004)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-l")))
				return
			}

		}

	}, 2)

	tmp1009 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp1001)

	_ = tmp1009

	tmp1010 := MakeNative(func(__e *ControlFlow) {
		V6805 := __e.Get(1)
		_ = V6805
		V6806 := __e.Get(2)
		_ = V6806
		tmp1148 := PrimEqual(MakeString(""), V6806)

		if True == tmp1148 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp1146 := Call(__e, PrimFunc(symshen_4_7string_2), V6806)

			var ifres1133 Obj

			if True == tmp1146 {
				tmp1144 := Call(__e, PrimFunc(symhdstr), V6806)

				tmp1145 := PrimEqual(MakeString("~"), tmp1144)

				var ifres1135 Obj

				if True == tmp1145 {
					tmp1142 := PrimTailString(V6806)

					tmp1143 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1142)

					var ifres1137 Obj

					if True == tmp1143 {
						tmp1139 := PrimTailString(V6806)

						tmp1140 := Call(__e, PrimFunc(symhdstr), tmp1139)

						tmp1141 := PrimEqual(MakeString("A"), tmp1140)

						var ifres1138 Obj

						if True == tmp1141 {
							ifres1138 = True

						} else {
							ifres1138 = False

						}

						ifres1137 = ifres1138

					} else {
						ifres1137 = False

					}

					var ifres1136 Obj

					if True == ifres1137 {
						ifres1136 = True

					} else {
						ifres1136 = False

					}

					ifres1135 = ifres1136

				} else {
					ifres1135 = False

				}

				var ifres1134 Obj

				if True == ifres1135 {
					ifres1134 = True

				} else {
					ifres1134 = False

				}

				ifres1133 = ifres1134

			} else {
				ifres1133 = False

			}

			if True == ifres1133 {
				tmp1011 := PrimTailString(V6806)

				tmp1012 := PrimTailString(tmp1011)

				tmp1013 := PrimCons(symshen_4a, Nil)

				tmp1014 := PrimCons(tmp1012, tmp1013)

				tmp1015 := PrimCons(V6805, tmp1014)

				__e.Return(PrimCons(symshen_4app, tmp1015))
				return

			} else {
				tmp1131 := Call(__e, PrimFunc(symshen_4_7string_2), V6806)

				var ifres1118 Obj

				if True == tmp1131 {
					tmp1129 := Call(__e, PrimFunc(symhdstr), V6806)

					tmp1130 := PrimEqual(MakeString("~"), tmp1129)

					var ifres1120 Obj

					if True == tmp1130 {
						tmp1127 := PrimTailString(V6806)

						tmp1128 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1127)

						var ifres1122 Obj

						if True == tmp1128 {
							tmp1124 := PrimTailString(V6806)

							tmp1125 := Call(__e, PrimFunc(symhdstr), tmp1124)

							tmp1126 := PrimEqual(MakeString("R"), tmp1125)

							var ifres1123 Obj

							if True == tmp1126 {
								ifres1123 = True

							} else {
								ifres1123 = False

							}

							ifres1122 = ifres1123

						} else {
							ifres1122 = False

						}

						var ifres1121 Obj

						if True == ifres1122 {
							ifres1121 = True

						} else {
							ifres1121 = False

						}

						ifres1120 = ifres1121

					} else {
						ifres1120 = False

					}

					var ifres1119 Obj

					if True == ifres1120 {
						ifres1119 = True

					} else {
						ifres1119 = False

					}

					ifres1118 = ifres1119

				} else {
					ifres1118 = False

				}

				if True == ifres1118 {
					tmp1016 := PrimTailString(V6806)

					tmp1017 := PrimTailString(tmp1016)

					tmp1018 := PrimCons(symshen_4r, Nil)

					tmp1019 := PrimCons(tmp1017, tmp1018)

					tmp1020 := PrimCons(V6805, tmp1019)

					__e.Return(PrimCons(symshen_4app, tmp1020))
					return

				} else {
					tmp1116 := Call(__e, PrimFunc(symshen_4_7string_2), V6806)

					var ifres1103 Obj

					if True == tmp1116 {
						tmp1114 := Call(__e, PrimFunc(symhdstr), V6806)

						tmp1115 := PrimEqual(MakeString("~"), tmp1114)

						var ifres1105 Obj

						if True == tmp1115 {
							tmp1112 := PrimTailString(V6806)

							tmp1113 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1112)

							var ifres1107 Obj

							if True == tmp1113 {
								tmp1109 := PrimTailString(V6806)

								tmp1110 := Call(__e, PrimFunc(symhdstr), tmp1109)

								tmp1111 := PrimEqual(MakeString("S"), tmp1110)

								var ifres1108 Obj

								if True == tmp1111 {
									ifres1108 = True

								} else {
									ifres1108 = False

								}

								ifres1107 = ifres1108

							} else {
								ifres1107 = False

							}

							var ifres1106 Obj

							if True == ifres1107 {
								ifres1106 = True

							} else {
								ifres1106 = False

							}

							ifres1105 = ifres1106

						} else {
							ifres1105 = False

						}

						var ifres1104 Obj

						if True == ifres1105 {
							ifres1104 = True

						} else {
							ifres1104 = False

						}

						ifres1103 = ifres1104

					} else {
						ifres1103 = False

					}

					if True == ifres1103 {
						tmp1021 := PrimTailString(V6806)

						tmp1022 := PrimTailString(tmp1021)

						tmp1023 := PrimCons(symshen_4s, Nil)

						tmp1024 := PrimCons(tmp1022, tmp1023)

						tmp1025 := PrimCons(V6805, tmp1024)

						__e.Return(PrimCons(symshen_4app, tmp1025))
						return

					} else {
						tmp1101 := Call(__e, PrimFunc(symshen_4_7string_2), V6806)

						if True == tmp1101 {
							tmp1026 := Call(__e, PrimFunc(symhdstr), V6806)

							tmp1027 := PrimTailString(V6806)

							tmp1028 := Call(__e, PrimFunc(symshen_4insert_1l), V6805, tmp1027)

							tmp1029 := PrimCons(tmp1028, Nil)

							tmp1030 := PrimCons(tmp1026, tmp1029)

							tmp1031 := PrimCons(symcn, tmp1030)

							__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp1031)
							return

						} else {
							tmp1099 := PrimIsPair(V6806)

							var ifres1080 Obj

							if True == tmp1099 {
								tmp1097 := PrimHead(V6806)

								tmp1098 := PrimEqual(symcn, tmp1097)

								var ifres1082 Obj

								if True == tmp1098 {
									tmp1095 := PrimTail(V6806)

									tmp1096 := PrimIsPair(tmp1095)

									var ifres1084 Obj

									if True == tmp1096 {
										tmp1092 := PrimTail(V6806)

										tmp1093 := PrimTail(tmp1092)

										tmp1094 := PrimIsPair(tmp1093)

										var ifres1086 Obj

										if True == tmp1094 {
											tmp1088 := PrimTail(V6806)

											tmp1089 := PrimTail(tmp1088)

											tmp1090 := PrimTail(tmp1089)

											tmp1091 := PrimEqual(Nil, tmp1090)

											var ifres1087 Obj

											if True == tmp1091 {
												ifres1087 = True

											} else {
												ifres1087 = False

											}

											ifres1086 = ifres1087

										} else {
											ifres1086 = False

										}

										var ifres1085 Obj

										if True == ifres1086 {
											ifres1085 = True

										} else {
											ifres1085 = False

										}

										ifres1084 = ifres1085

									} else {
										ifres1084 = False

									}

									var ifres1083 Obj

									if True == ifres1084 {
										ifres1083 = True

									} else {
										ifres1083 = False

									}

									ifres1082 = ifres1083

								} else {
									ifres1082 = False

								}

								var ifres1081 Obj

								if True == ifres1082 {
									ifres1081 = True

								} else {
									ifres1081 = False

								}

								ifres1080 = ifres1081

							} else {
								ifres1080 = False

							}

							if True == ifres1080 {
								tmp1032 := PrimTail(V6806)

								tmp1033 := PrimHead(tmp1032)

								tmp1034 := PrimTail(V6806)

								tmp1035 := PrimTail(tmp1034)

								tmp1036 := PrimHead(tmp1035)

								tmp1037 := Call(__e, PrimFunc(symshen_4insert_1l), V6805, tmp1036)

								tmp1038 := PrimCons(tmp1037, Nil)

								tmp1039 := PrimCons(tmp1033, tmp1038)

								__e.Return(PrimCons(symcn, tmp1039))
								return

							} else {
								tmp1078 := PrimIsPair(V6806)

								var ifres1052 Obj

								if True == tmp1078 {
									tmp1076 := PrimHead(V6806)

									tmp1077 := PrimEqual(symshen_4app, tmp1076)

									var ifres1054 Obj

									if True == tmp1077 {
										tmp1074 := PrimTail(V6806)

										tmp1075 := PrimIsPair(tmp1074)

										var ifres1056 Obj

										if True == tmp1075 {
											tmp1071 := PrimTail(V6806)

											tmp1072 := PrimTail(tmp1071)

											tmp1073 := PrimIsPair(tmp1072)

											var ifres1058 Obj

											if True == tmp1073 {
												tmp1067 := PrimTail(V6806)

												tmp1068 := PrimTail(tmp1067)

												tmp1069 := PrimTail(tmp1068)

												tmp1070 := PrimIsPair(tmp1069)

												var ifres1060 Obj

												if True == tmp1070 {
													tmp1062 := PrimTail(V6806)

													tmp1063 := PrimTail(tmp1062)

													tmp1064 := PrimTail(tmp1063)

													tmp1065 := PrimTail(tmp1064)

													tmp1066 := PrimEqual(Nil, tmp1065)

													var ifres1061 Obj

													if True == tmp1066 {
														ifres1061 = True

													} else {
														ifres1061 = False

													}

													ifres1060 = ifres1061

												} else {
													ifres1060 = False

												}

												var ifres1059 Obj

												if True == ifres1060 {
													ifres1059 = True

												} else {
													ifres1059 = False

												}

												ifres1058 = ifres1059

											} else {
												ifres1058 = False

											}

											var ifres1057 Obj

											if True == ifres1058 {
												ifres1057 = True

											} else {
												ifres1057 = False

											}

											ifres1056 = ifres1057

										} else {
											ifres1056 = False

										}

										var ifres1055 Obj

										if True == ifres1056 {
											ifres1055 = True

										} else {
											ifres1055 = False

										}

										ifres1054 = ifres1055

									} else {
										ifres1054 = False

									}

									var ifres1053 Obj

									if True == ifres1054 {
										ifres1053 = True

									} else {
										ifres1053 = False

									}

									ifres1052 = ifres1053

								} else {
									ifres1052 = False

								}

								if True == ifres1052 {
									tmp1040 := PrimTail(V6806)

									tmp1041 := PrimHead(tmp1040)

									tmp1042 := PrimTail(V6806)

									tmp1043 := PrimTail(tmp1042)

									tmp1044 := PrimHead(tmp1043)

									tmp1045 := Call(__e, PrimFunc(symshen_4insert_1l), V6805, tmp1044)

									tmp1046 := PrimTail(V6806)

									tmp1047 := PrimTail(tmp1046)

									tmp1048 := PrimTail(tmp1047)

									tmp1049 := PrimCons(tmp1045, tmp1048)

									tmp1050 := PrimCons(tmp1041, tmp1049)

									__e.Return(PrimCons(symshen_4app, tmp1050))
									return

								} else {
									__e.Return(PrimSimpleError(MakeString("implementation error in shen.insert-l")))
									return
								}

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp1149 := Call(__e, ns2_1set, symshen_4insert_1l, tmp1010)

	_ = tmp1149

	tmp1150 := MakeNative(func(__e *ControlFlow) {
		V6807 := __e.Get(1)
		_ = V6807
		tmp1235 := PrimIsPair(V6807)

		var ifres1166 Obj

		if True == tmp1235 {
			tmp1233 := PrimHead(V6807)

			tmp1234 := PrimEqual(symcn, tmp1233)

			var ifres1168 Obj

			if True == tmp1234 {
				tmp1231 := PrimTail(V6807)

				tmp1232 := PrimIsPair(tmp1231)

				var ifres1170 Obj

				if True == tmp1232 {
					tmp1228 := PrimTail(V6807)

					tmp1229 := PrimTail(tmp1228)

					tmp1230 := PrimIsPair(tmp1229)

					var ifres1172 Obj

					if True == tmp1230 {
						tmp1224 := PrimTail(V6807)

						tmp1225 := PrimTail(tmp1224)

						tmp1226 := PrimHead(tmp1225)

						tmp1227 := PrimIsPair(tmp1226)

						var ifres1174 Obj

						if True == tmp1227 {
							tmp1219 := PrimTail(V6807)

							tmp1220 := PrimTail(tmp1219)

							tmp1221 := PrimHead(tmp1220)

							tmp1222 := PrimHead(tmp1221)

							tmp1223 := PrimEqual(symcn, tmp1222)

							var ifres1176 Obj

							if True == tmp1223 {
								tmp1214 := PrimTail(V6807)

								tmp1215 := PrimTail(tmp1214)

								tmp1216 := PrimHead(tmp1215)

								tmp1217 := PrimTail(tmp1216)

								tmp1218 := PrimIsPair(tmp1217)

								var ifres1178 Obj

								if True == tmp1218 {
									tmp1208 := PrimTail(V6807)

									tmp1209 := PrimTail(tmp1208)

									tmp1210 := PrimHead(tmp1209)

									tmp1211 := PrimTail(tmp1210)

									tmp1212 := PrimTail(tmp1211)

									tmp1213 := PrimIsPair(tmp1212)

									var ifres1180 Obj

									if True == tmp1213 {
										tmp1201 := PrimTail(V6807)

										tmp1202 := PrimTail(tmp1201)

										tmp1203 := PrimHead(tmp1202)

										tmp1204 := PrimTail(tmp1203)

										tmp1205 := PrimTail(tmp1204)

										tmp1206 := PrimTail(tmp1205)

										tmp1207 := PrimEqual(Nil, tmp1206)

										var ifres1182 Obj

										if True == tmp1207 {
											tmp1197 := PrimTail(V6807)

											tmp1198 := PrimTail(tmp1197)

											tmp1199 := PrimTail(tmp1198)

											tmp1200 := PrimEqual(Nil, tmp1199)

											var ifres1184 Obj

											if True == tmp1200 {
												tmp1194 := PrimTail(V6807)

												tmp1195 := PrimHead(tmp1194)

												tmp1196 := PrimIsString(tmp1195)

												var ifres1186 Obj

												if True == tmp1196 {
													tmp1188 := PrimTail(V6807)

													tmp1189 := PrimTail(tmp1188)

													tmp1190 := PrimHead(tmp1189)

													tmp1191 := PrimTail(tmp1190)

													tmp1192 := PrimHead(tmp1191)

													tmp1193 := PrimIsString(tmp1192)

													var ifres1187 Obj

													if True == tmp1193 {
														ifres1187 = True

													} else {
														ifres1187 = False

													}

													ifres1186 = ifres1187

												} else {
													ifres1186 = False

												}

												var ifres1185 Obj

												if True == ifres1186 {
													ifres1185 = True

												} else {
													ifres1185 = False

												}

												ifres1184 = ifres1185

											} else {
												ifres1184 = False

											}

											var ifres1183 Obj

											if True == ifres1184 {
												ifres1183 = True

											} else {
												ifres1183 = False

											}

											ifres1182 = ifres1183

										} else {
											ifres1182 = False

										}

										var ifres1181 Obj

										if True == ifres1182 {
											ifres1181 = True

										} else {
											ifres1181 = False

										}

										ifres1180 = ifres1181

									} else {
										ifres1180 = False

									}

									var ifres1179 Obj

									if True == ifres1180 {
										ifres1179 = True

									} else {
										ifres1179 = False

									}

									ifres1178 = ifres1179

								} else {
									ifres1178 = False

								}

								var ifres1177 Obj

								if True == ifres1178 {
									ifres1177 = True

								} else {
									ifres1177 = False

								}

								ifres1176 = ifres1177

							} else {
								ifres1176 = False

							}

							var ifres1175 Obj

							if True == ifres1176 {
								ifres1175 = True

							} else {
								ifres1175 = False

							}

							ifres1174 = ifres1175

						} else {
							ifres1174 = False

						}

						var ifres1173 Obj

						if True == ifres1174 {
							ifres1173 = True

						} else {
							ifres1173 = False

						}

						ifres1172 = ifres1173

					} else {
						ifres1172 = False

					}

					var ifres1171 Obj

					if True == ifres1172 {
						ifres1171 = True

					} else {
						ifres1171 = False

					}

					ifres1170 = ifres1171

				} else {
					ifres1170 = False

				}

				var ifres1169 Obj

				if True == ifres1170 {
					ifres1169 = True

				} else {
					ifres1169 = False

				}

				ifres1168 = ifres1169

			} else {
				ifres1168 = False

			}

			var ifres1167 Obj

			if True == ifres1168 {
				ifres1167 = True

			} else {
				ifres1167 = False

			}

			ifres1166 = ifres1167

		} else {
			ifres1166 = False

		}

		if True == ifres1166 {
			tmp1151 := PrimTail(V6807)

			tmp1152 := PrimHead(tmp1151)

			tmp1153 := PrimTail(V6807)

			tmp1154 := PrimTail(tmp1153)

			tmp1155 := PrimHead(tmp1154)

			tmp1156 := PrimTail(tmp1155)

			tmp1157 := PrimHead(tmp1156)

			tmp1158 := PrimStringConcat(tmp1152, tmp1157)

			tmp1159 := PrimTail(V6807)

			tmp1160 := PrimTail(tmp1159)

			tmp1161 := PrimHead(tmp1160)

			tmp1162 := PrimTail(tmp1161)

			tmp1163 := PrimTail(tmp1162)

			tmp1164 := PrimCons(tmp1158, tmp1163)

			__e.Return(PrimCons(symcn, tmp1164))
			return

		} else {
			__e.Return(V6807)
			return
		}

	}, 1)

	tmp1236 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp1150)

	_ = tmp1236

	tmp1237 := MakeNative(func(__e *ControlFlow) {
		V6810 := __e.Get(1)
		_ = V6810
		tmp1263 := PrimEqual(MakeString(""), V6810)

		if True == tmp1263 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp1261 := Call(__e, PrimFunc(symshen_4_7string_2), V6810)

			var ifres1248 Obj

			if True == tmp1261 {
				tmp1259 := Call(__e, PrimFunc(symhdstr), V6810)

				tmp1260 := PrimEqual(MakeString("~"), tmp1259)

				var ifres1250 Obj

				if True == tmp1260 {
					tmp1257 := PrimTailString(V6810)

					tmp1258 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1257)

					var ifres1252 Obj

					if True == tmp1258 {
						tmp1254 := PrimTailString(V6810)

						tmp1255 := Call(__e, PrimFunc(symhdstr), tmp1254)

						tmp1256 := PrimEqual(MakeString("%"), tmp1255)

						var ifres1253 Obj

						if True == tmp1256 {
							ifres1253 = True

						} else {
							ifres1253 = False

						}

						ifres1252 = ifres1253

					} else {
						ifres1252 = False

					}

					var ifres1251 Obj

					if True == ifres1252 {
						ifres1251 = True

					} else {
						ifres1251 = False

					}

					ifres1250 = ifres1251

				} else {
					ifres1250 = False

				}

				var ifres1249 Obj

				if True == ifres1250 {
					ifres1249 = True

				} else {
					ifres1249 = False

				}

				ifres1248 = ifres1249

			} else {
				ifres1248 = False

			}

			if True == ifres1248 {
				tmp1238 := PrimNumberToString(MakeNumber(10))

				tmp1239 := PrimTailString(V6810)

				tmp1240 := PrimTailString(tmp1239)

				tmp1241 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp1240)

				__e.Return(PrimStringConcat(tmp1238, tmp1241))
				return

			} else {
				tmp1246 := Call(__e, PrimFunc(symshen_4_7string_2), V6810)

				if True == tmp1246 {
					tmp1242 := Call(__e, PrimFunc(symhdstr), V6810)

					tmp1243 := PrimTailString(V6810)

					tmp1244 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp1243)

					__e.Return(PrimStringConcat(tmp1242, tmp1244))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.proc-nl")))
					return
				}

			}

		}

	}, 1)

	tmp1264 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp1237)

	_ = tmp1264

	tmp1265 := MakeNative(func(__e *ControlFlow) {
		V6815 := __e.Get(1)
		_ = V6815
		V6816 := __e.Get(2)
		_ = V6816
		tmp1274 := PrimEqual(Nil, V6816)

		if True == tmp1274 {
			__e.Return(V6815)
			return
		} else {
			tmp1272 := PrimIsPair(V6816)

			if True == tmp1272 {
				tmp1266 := PrimHead(V6816)

				tmp1267 := PrimCons(V6815, Nil)

				tmp1268 := PrimCons(tmp1266, tmp1267)

				tmp1269 := PrimCons(symshen_4insert, tmp1268)

				tmp1270 := PrimTail(V6816)

				__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp1269, tmp1270)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-r")))
				return
			}

		}

	}, 2)

	tmp1275 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp1265)

	_ = tmp1275

	tmp1276 := MakeNative(func(__e *ControlFlow) {
		V6817 := __e.Get(1)
		_ = V6817
		V6818 := __e.Get(2)
		_ = V6818
		__e.TailApply(PrimFunc(symshen_4insert_1h), V6817, V6818, MakeString(""))
		return
	}, 2)

	tmp1277 := Call(__e, ns2_1set, symshen_4insert, tmp1276)

	_ = tmp1277

	tmp1278 := MakeNative(func(__e *ControlFlow) {
		V6827 := __e.Get(1)
		_ = V6827
		V6828 := __e.Get(2)
		_ = V6828
		V6829 := __e.Get(3)
		_ = V6829
		tmp1339 := PrimEqual(MakeString(""), V6828)

		if True == tmp1339 {
			__e.Return(V6829)
			return
		} else {
			tmp1337 := Call(__e, PrimFunc(symshen_4_7string_2), V6828)

			var ifres1324 Obj

			if True == tmp1337 {
				tmp1335 := Call(__e, PrimFunc(symhdstr), V6828)

				tmp1336 := PrimEqual(MakeString("~"), tmp1335)

				var ifres1326 Obj

				if True == tmp1336 {
					tmp1333 := PrimTailString(V6828)

					tmp1334 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1333)

					var ifres1328 Obj

					if True == tmp1334 {
						tmp1330 := PrimTailString(V6828)

						tmp1331 := Call(__e, PrimFunc(symhdstr), tmp1330)

						tmp1332 := PrimEqual(MakeString("A"), tmp1331)

						var ifres1329 Obj

						if True == tmp1332 {
							ifres1329 = True

						} else {
							ifres1329 = False

						}

						ifres1328 = ifres1329

					} else {
						ifres1328 = False

					}

					var ifres1327 Obj

					if True == ifres1328 {
						ifres1327 = True

					} else {
						ifres1327 = False

					}

					ifres1326 = ifres1327

				} else {
					ifres1326 = False

				}

				var ifres1325 Obj

				if True == ifres1326 {
					ifres1325 = True

				} else {
					ifres1325 = False

				}

				ifres1324 = ifres1325

			} else {
				ifres1324 = False

			}

			if True == ifres1324 {
				tmp1279 := PrimTailString(V6828)

				tmp1280 := PrimTailString(tmp1279)

				tmp1281 := Call(__e, PrimFunc(symshen_4app), V6827, tmp1280, symshen_4a)

				__e.Return(PrimStringConcat(V6829, tmp1281))
				return

			} else {
				tmp1322 := Call(__e, PrimFunc(symshen_4_7string_2), V6828)

				var ifres1309 Obj

				if True == tmp1322 {
					tmp1320 := Call(__e, PrimFunc(symhdstr), V6828)

					tmp1321 := PrimEqual(MakeString("~"), tmp1320)

					var ifres1311 Obj

					if True == tmp1321 {
						tmp1318 := PrimTailString(V6828)

						tmp1319 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1318)

						var ifres1313 Obj

						if True == tmp1319 {
							tmp1315 := PrimTailString(V6828)

							tmp1316 := Call(__e, PrimFunc(symhdstr), tmp1315)

							tmp1317 := PrimEqual(MakeString("R"), tmp1316)

							var ifres1314 Obj

							if True == tmp1317 {
								ifres1314 = True

							} else {
								ifres1314 = False

							}

							ifres1313 = ifres1314

						} else {
							ifres1313 = False

						}

						var ifres1312 Obj

						if True == ifres1313 {
							ifres1312 = True

						} else {
							ifres1312 = False

						}

						ifres1311 = ifres1312

					} else {
						ifres1311 = False

					}

					var ifres1310 Obj

					if True == ifres1311 {
						ifres1310 = True

					} else {
						ifres1310 = False

					}

					ifres1309 = ifres1310

				} else {
					ifres1309 = False

				}

				if True == ifres1309 {
					tmp1282 := PrimTailString(V6828)

					tmp1283 := PrimTailString(tmp1282)

					tmp1284 := Call(__e, PrimFunc(symshen_4app), V6827, tmp1283, symshen_4r)

					__e.Return(PrimStringConcat(V6829, tmp1284))
					return

				} else {
					tmp1307 := Call(__e, PrimFunc(symshen_4_7string_2), V6828)

					var ifres1294 Obj

					if True == tmp1307 {
						tmp1305 := Call(__e, PrimFunc(symhdstr), V6828)

						tmp1306 := PrimEqual(MakeString("~"), tmp1305)

						var ifres1296 Obj

						if True == tmp1306 {
							tmp1303 := PrimTailString(V6828)

							tmp1304 := Call(__e, PrimFunc(symshen_4_7string_2), tmp1303)

							var ifres1298 Obj

							if True == tmp1304 {
								tmp1300 := PrimTailString(V6828)

								tmp1301 := Call(__e, PrimFunc(symhdstr), tmp1300)

								tmp1302 := PrimEqual(MakeString("S"), tmp1301)

								var ifres1299 Obj

								if True == tmp1302 {
									ifres1299 = True

								} else {
									ifres1299 = False

								}

								ifres1298 = ifres1299

							} else {
								ifres1298 = False

							}

							var ifres1297 Obj

							if True == ifres1298 {
								ifres1297 = True

							} else {
								ifres1297 = False

							}

							ifres1296 = ifres1297

						} else {
							ifres1296 = False

						}

						var ifres1295 Obj

						if True == ifres1296 {
							ifres1295 = True

						} else {
							ifres1295 = False

						}

						ifres1294 = ifres1295

					} else {
						ifres1294 = False

					}

					if True == ifres1294 {
						tmp1285 := PrimTailString(V6828)

						tmp1286 := PrimTailString(tmp1285)

						tmp1287 := Call(__e, PrimFunc(symshen_4app), V6827, tmp1286, symshen_4s)

						__e.Return(PrimStringConcat(V6829, tmp1287))
						return

					} else {
						tmp1292 := Call(__e, PrimFunc(symshen_4_7string_2), V6828)

						if True == tmp1292 {
							tmp1288 := PrimTailString(V6828)

							tmp1289 := Call(__e, PrimFunc(symhdstr), V6828)

							tmp1290 := PrimStringConcat(V6829, tmp1289)

							__e.TailApply(PrimFunc(symshen_4insert_1h), V6827, tmp1288, tmp1290)
							return

						} else {
							__e.Return(PrimSimpleError(MakeString("implementation error in shen.insert-h")))
							return
						}

					}

				}

			}

		}

	}, 3)

	tmp1340 := Call(__e, ns2_1set, symshen_4insert_1h, tmp1278)

	_ = tmp1340

	tmp1341 := MakeNative(func(__e *ControlFlow) {
		V6830 := __e.Get(1)
		_ = V6830
		V6831 := __e.Get(2)
		_ = V6831
		V6832 := __e.Get(3)
		_ = V6832
		tmp1342 := Call(__e, PrimFunc(symshen_4arg_1_6str), V6830, V6832)

		__e.Return(PrimStringConcat(tmp1342, V6831))
		return

	}, 3)

	tmp1343 := Call(__e, ns2_1set, symshen_4app, tmp1341)

	_ = tmp1343

	tmp1344 := MakeNative(func(__e *ControlFlow) {
		V6836 := __e.Get(1)
		_ = V6836
		V6837 := __e.Get(2)
		_ = V6837
		tmp1352 := Call(__e, PrimFunc(symfail))

		tmp1353 := PrimEqual(V6836, tmp1352)

		if True == tmp1353 {
			__e.Return(MakeString("..."))
			return
		} else {
			tmp1350 := Call(__e, PrimFunc(symshen_4list_2), V6836)

			if True == tmp1350 {
				__e.TailApply(PrimFunc(symshen_4list_1_6str), V6836, V6837)
				return
			} else {
				tmp1348 := PrimIsString(V6836)

				if True == tmp1348 {
					__e.TailApply(PrimFunc(symshen_4str_1_6str), V6836, V6837)
					return
				} else {
					tmp1346 := PrimIsVector(V6836)

					if True == tmp1346 {
						__e.TailApply(PrimFunc(symshen_4vector_1_6str), V6836, V6837)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4atom_1_6str), V6836)
						return
					}

				}

			}

		}

	}, 2)

	tmp1354 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp1344)

	_ = tmp1354

	tmp1355 := MakeNative(func(__e *ControlFlow) {
		V6838 := __e.Get(1)
		_ = V6838
		V6839 := __e.Get(2)
		_ = V6839
		tmp1363 := PrimEqual(symshen_4r, V6839)

		if True == tmp1363 {
			tmp1356 := Call(__e, PrimFunc(symshen_4maxseq))

			tmp1357 := Call(__e, PrimFunc(symshen_4iter_1list), V6838, symshen_4r, tmp1356)

			tmp1358 := Call(__e, PrimFunc(sym_8s), tmp1357, MakeString(")"))

			__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp1358)
			return

		} else {
			tmp1359 := Call(__e, PrimFunc(symshen_4maxseq))

			tmp1360 := Call(__e, PrimFunc(symshen_4iter_1list), V6838, V6839, tmp1359)

			tmp1361 := Call(__e, PrimFunc(sym_8s), tmp1360, MakeString("]"))

			__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp1361)
			return

		}

	}, 2)

	tmp1364 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp1355)

	_ = tmp1364

	tmp1365 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dmaximum_1print_1sequence_1size_d))
		return
	}, 0)

	tmp1366 := Call(__e, ns2_1set, symshen_4maxseq, tmp1365)

	_ = tmp1366

	tmp1367 := MakeNative(func(__e *ControlFlow) {
		V6850 := __e.Get(1)
		_ = V6850
		V6851 := __e.Get(2)
		_ = V6851
		V6852 := __e.Get(3)
		_ = V6852
		tmp1388 := PrimEqual(Nil, V6850)

		if True == tmp1388 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp1386 := PrimEqual(MakeNumber(0), V6852)

			if True == tmp1386 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				tmp1384 := PrimIsPair(V6850)

				var ifres1380 Obj

				if True == tmp1384 {
					tmp1382 := PrimTail(V6850)

					tmp1383 := PrimEqual(Nil, tmp1382)

					var ifres1381 Obj

					if True == tmp1383 {
						ifres1381 = True

					} else {
						ifres1381 = False

					}

					ifres1380 = ifres1381

				} else {
					ifres1380 = False

				}

				if True == ifres1380 {
					tmp1368 := PrimHead(V6850)

					__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp1368, V6851)
					return

				} else {
					tmp1378 := PrimIsPair(V6850)

					if True == tmp1378 {
						tmp1369 := PrimHead(V6850)

						tmp1370 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp1369, V6851)

						tmp1371 := PrimTail(V6850)

						tmp1372 := PrimNumberSubtract(V6852, MakeNumber(1))

						tmp1373 := Call(__e, PrimFunc(symshen_4iter_1list), tmp1371, V6851, tmp1372)

						tmp1374 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp1373)

						__e.TailApply(PrimFunc(sym_8s), tmp1370, tmp1374)
						return

					} else {
						tmp1375 := Call(__e, PrimFunc(symshen_4arg_1_6str), V6850, V6851)

						tmp1376 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp1375)

						__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp1376)
						return

					}

				}

			}

		}

	}, 3)

	tmp1389 := Call(__e, ns2_1set, symshen_4iter_1list, tmp1367)

	_ = tmp1389

	tmp1390 := MakeNative(func(__e *ControlFlow) {
		V6855 := __e.Get(1)
		_ = V6855
		V6856 := __e.Get(2)
		_ = V6856
		tmp1395 := PrimEqual(symshen_4a, V6856)

		if True == tmp1395 {
			__e.Return(V6855)
			return
		} else {
			tmp1391 := PrimNumberToString(MakeNumber(34))

			tmp1392 := PrimNumberToString(MakeNumber(34))

			tmp1393 := Call(__e, PrimFunc(sym_8s), V6855, tmp1392)

			__e.TailApply(PrimFunc(sym_8s), tmp1391, tmp1393)
			return

		}

	}, 2)

	tmp1396 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp1390)

	_ = tmp1396

	tmp1397 := MakeNative(func(__e *ControlFlow) {
		V6857 := __e.Get(1)
		_ = V6857
		V6858 := __e.Get(2)
		_ = V6858
		tmp1410 := Call(__e, PrimFunc(symshen_4print_1vector_2), V6857)

		if True == tmp1410 {
			tmp1398 := PrimVectorGet(V6857, MakeNumber(0))

			tmp1399 := Call(__e, PrimFunc(symfn), tmp1398)

			__e.TailApply(tmp1399, V6857)
			return

		} else {
			tmp1408 := Call(__e, PrimFunc(symvector_2), V6857)

			if True == tmp1408 {
				tmp1400 := Call(__e, PrimFunc(symshen_4maxseq))

				tmp1401 := Call(__e, PrimFunc(symshen_4iter_1vector), V6857, MakeNumber(1), V6858, tmp1400)

				tmp1402 := Call(__e, PrimFunc(sym_8s), tmp1401, MakeString(">"))

				__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp1402)
				return

			} else {
				tmp1403 := Call(__e, PrimFunc(symshen_4maxseq))

				tmp1404 := Call(__e, PrimFunc(symshen_4iter_1vector), V6857, MakeNumber(0), V6858, tmp1403)

				tmp1405 := Call(__e, PrimFunc(sym_8s), tmp1404, MakeString(">>"))

				tmp1406 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp1405)

				__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp1406)
				return

			}

		}

	}, 2)

	tmp1411 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp1397)

	_ = tmp1411

	tmp1412 := MakeNative(func(__e *ControlFlow) {
		V6859 := __e.Get(1)
		_ = V6859
		tmp1413 := PrimValue(symshen_4_dempty_1absvector_d)

		__e.Return(PrimEqual(V6859, tmp1413))
		return

	}, 1)

	tmp1414 := Call(__e, ns2_1set, symshen_4empty_1absvector_2, tmp1412)

	_ = tmp1414

	tmp1415 := MakeNative(func(__e *ControlFlow) {
		V6860 := __e.Get(1)
		_ = V6860
		tmp1435 := Call(__e, PrimFunc(symshen_4empty_1absvector_2), V6860)

		tmp1436 := PrimNot(tmp1435)

		if True == tmp1436 {
			tmp1417 := MakeNative(func(__e *ControlFlow) {
				W6861 := __e.Get(1)
				_ = W6861
				tmp1431 := PrimEqual(W6861, symshen_4tuple)

				if True == tmp1431 {
					__e.Return(True)
					return
				} else {
					tmp1429 := PrimEqual(W6861, symshen_4pvar)

					var ifres1419 Obj

					if True == tmp1429 {
						ifres1419 = True

					} else {
						tmp1428 := PrimEqual(W6861, symshen_4dictionary)

						var ifres1421 Obj

						if True == tmp1428 {
							ifres1421 = True

						} else {
							tmp1426 := PrimIsNumber(W6861)

							tmp1427 := PrimNot(tmp1426)

							var ifres1423 Obj

							if True == tmp1427 {
								tmp1425 := Call(__e, PrimFunc(symshen_4fbound_2), W6861)

								var ifres1424 Obj

								if True == tmp1425 {
									ifres1424 = True

								} else {
									ifres1424 = False

								}

								ifres1423 = ifres1424

							} else {
								ifres1423 = False

							}

							var ifres1422 Obj

							if True == ifres1423 {
								ifres1422 = True

							} else {
								ifres1422 = False

							}

							ifres1421 = ifres1422

						}

						var ifres1420 Obj

						if True == ifres1421 {
							ifres1420 = True

						} else {
							ifres1420 = False

						}

						ifres1419 = ifres1420

					}

					if True == ifres1419 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}, 1)

			tmp1432 := PrimVectorGet(V6860, MakeNumber(0))

			tmp1433 := Call(__e, tmp1417, tmp1432)

			if True == tmp1433 {
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

	tmp1437 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp1415)

	_ = tmp1437

	tmp1438 := MakeNative(func(__e *ControlFlow) {
		V6862 := __e.Get(1)
		_ = V6862
		tmp1439 := Call(__e, PrimFunc(symarity), V6862)

		tmp1440 := PrimEqual(tmp1439, MakeNumber(-1))

		__e.Return(PrimNot(tmp1440))
		return

	}, 1)

	tmp1441 := Call(__e, ns2_1set, symshen_4fbound_2, tmp1438)

	_ = tmp1441

	tmp1442 := MakeNative(func(__e *ControlFlow) {
		V6863 := __e.Get(1)
		_ = V6863
		tmp1443 := PrimVectorGet(V6863, MakeNumber(1))

		tmp1444 := PrimVectorGet(V6863, MakeNumber(2))

		tmp1445 := Call(__e, PrimFunc(symshen_4app), tmp1444, MakeString(")"), symshen_4s)

		tmp1446 := PrimStringConcat(MakeString(" "), tmp1445)

		tmp1447 := Call(__e, PrimFunc(symshen_4app), tmp1443, tmp1446, symshen_4s)

		__e.Return(PrimStringConcat(MakeString("(@p "), tmp1447))
		return

	}, 1)

	tmp1448 := Call(__e, ns2_1set, symshen_4tuple, tmp1442)

	_ = tmp1448

	tmp1449 := MakeNative(func(__e *ControlFlow) {
		V6864 := __e.Get(1)
		_ = V6864
		__e.Return(MakeString("(dict ...)"))
		return
	}, 1)

	tmp1450 := Call(__e, ns2_1set, symshen_4dictionary, tmp1449)

	_ = tmp1450

	tmp1451 := MakeNative(func(__e *ControlFlow) {
		V6871 := __e.Get(1)
		_ = V6871
		V6872 := __e.Get(2)
		_ = V6872
		V6873 := __e.Get(3)
		_ = V6873
		V6874 := __e.Get(4)
		_ = V6874
		tmp1471 := PrimEqual(MakeNumber(0), V6874)

		if True == tmp1471 {
			__e.Return(MakeString("... etc"))
			return
		} else {
			tmp1452 := MakeNative(func(__e *ControlFlow) {
				W6875 := __e.Get(1)
				_ = W6875
				tmp1453 := MakeNative(func(__e *ControlFlow) {
					W6877 := __e.Get(1)
					_ = W6877
					tmp1462 := PrimEqual(W6875, symshen_4out_1of_1bounds)

					if True == tmp1462 {
						__e.Return(MakeString(""))
						return
					} else {
						tmp1460 := PrimEqual(W6877, symshen_4out_1of_1bounds)

						if True == tmp1460 {
							__e.TailApply(PrimFunc(symshen_4arg_1_6str), W6875, V6873)
							return
						} else {
							tmp1454 := Call(__e, PrimFunc(symshen_4arg_1_6str), W6875, V6873)

							tmp1455 := PrimNumberAdd(V6872, MakeNumber(1))

							tmp1456 := PrimNumberSubtract(V6874, MakeNumber(1))

							tmp1457 := Call(__e, PrimFunc(symshen_4iter_1vector), V6871, tmp1455, V6873, tmp1456)

							tmp1458 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp1457)

							__e.TailApply(PrimFunc(sym_8s), tmp1454, tmp1458)
							return

						}

					}

				}, 1)

				tmp1463 := MakeNative(func(__e *ControlFlow) {
					tmp1464 := PrimNumberAdd(V6872, MakeNumber(1))

					__e.Return(PrimVectorGet(V6871, tmp1464))
					return

				}, 0)

				tmp1465 := MakeNative(func(__e *ControlFlow) {
					Z6878 := __e.Get(1)
					_ = Z6878
					__e.Return(symshen_4out_1of_1bounds)
					return
				}, 1)

				tmp1466 := Call(__e, try_1catch, tmp1463, tmp1465)

				__e.TailApply(tmp1453, tmp1466)
				return

			}, 1)

			tmp1467 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V6871, V6872))
				return
			}, 0)

			tmp1468 := MakeNative(func(__e *ControlFlow) {
				Z6876 := __e.Get(1)
				_ = Z6876
				__e.Return(symshen_4out_1of_1bounds)
				return
			}, 1)

			tmp1469 := Call(__e, try_1catch, tmp1467, tmp1468)

			__e.TailApply(tmp1452, tmp1469)
			return

		}

	}, 4)

	tmp1472 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp1451)

	_ = tmp1472

	tmp1473 := MakeNative(func(__e *ControlFlow) {
		V6879 := __e.Get(1)
		_ = V6879
		tmp1474 := MakeNative(func(__e *ControlFlow) {
			__e.Return(PrimStr(V6879))
			return
		}, 0)

		tmp1475 := MakeNative(func(__e *ControlFlow) {
			Z6880 := __e.Get(1)
			_ = Z6880
			__e.TailApply(PrimFunc(symshen_4funexstring))
			return
		}, 1)

		__e.TailApply(try_1catch, tmp1474, tmp1475)
		return

	}, 1)

	tmp1476 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp1473)

	_ = tmp1476

	tmp1477 := MakeNative(func(__e *ControlFlow) {
		tmp1478 := PrimIntern(MakeString("x"))

		tmp1479 := Call(__e, PrimFunc(symgensym), tmp1478)

		tmp1480 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp1479, symshen_4a)

		tmp1481 := Call(__e, PrimFunc(sym_8s), tmp1480, MakeString("\x11"))

		tmp1482 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp1481)

		tmp1483 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp1482)

		tmp1484 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp1483)

		tmp1485 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp1484)

		__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp1485)
		return

	}, 0)

	tmp1486 := Call(__e, ns2_1set, symshen_4funexstring, tmp1477)

	_ = tmp1486

	tmp1487 := MakeNative(func(__e *ControlFlow) {
		V6881 := __e.Get(1)
		_ = V6881
		tmp1491 := Call(__e, PrimFunc(symempty_2), V6881)

		if True == tmp1491 {
			__e.Return(True)
			return
		} else {
			tmp1489 := PrimIsPair(V6881)

			if True == tmp1489 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	__e.TailApply(ns2_1set, symshen_4list_2, tmp1487)
	return

}, 0)
