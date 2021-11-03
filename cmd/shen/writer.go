package main

import . "github.com/tiancaiamao/shen-go/cora"

var WriterMain = MakeNative(func(__e *ControlFlow) {
	tmp963 := MakeNative(func(__e *ControlFlow) {
		V4723 := __e.Get(1)
		_ = V4723
		tmp964 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			tmp965 := MakeNative(func(__e *ControlFlow) {
				Print := __e.Get(1)
				_ = Print
				__e.Return(V4723)
				return
			}, 1)

			tmp966 := Call(__e, PrimNS2Value(symstoutput))

			tmp967 := Call(__e, PrimNS2Value(sympr), String, tmp966)

			__e.TailApply(tmp965, tmp967)
			return

		}, 1)

		tmp968 := Call(__e, PrimNS2Value(symshen_4insert), V4723, MakeString("~S"))

		__e.TailApply(tmp964, tmp968)
		return

	}, 1)

	tmp969 := Call(__e, PrimNS2Value(symdef), symprint, tmp963)

	_ = tmp969

	tmp970 := MakeNative(func(__e *ControlFlow) {
		V4724 := __e.Get(1)
		_ = V4724
		V4725 := __e.Get(2)
		_ = V4725
		tmp975 := PrimNS3Value(sym_dhush_d)

		if True == tmp975 {
			__e.Return(V4724)
			return
		} else {
			tmp974 := Call(__e, PrimNS2Value(symshen_4char_1stoutput_2), V4725)

			if True == tmp974 {
				__e.TailApply(PrimNS2Value(symshen_4write_1string), V4724, V4725)
				return
			} else {
				tmp973 := Call(__e, PrimNS2Value(symshen_4string_1_6byte), V4724, MakeNumber(0))

				__e.TailApply(PrimNS2Value(symshen_4write_1chars), V4724, V4725, tmp973, MakeNumber(1))
				return

			}

		}

	}, 2)

	tmp976 := Call(__e, PrimNS2Value(symdef), sympr, tmp970)

	_ = tmp976

	tmp977 := MakeNative(func(__e *ControlFlow) {
		V4726 := __e.Get(1)
		_ = V4726
		V4727 := __e.Get(2)
		_ = V4727
		tmp978 := MakeNative(func(__e *ControlFlow) {
			tmp979 := PrimPos(V4726, V4727)

			__e.Return(PrimStringToNumber(tmp979))
			return

		}, 0)

		tmp980 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(symshen_4eos)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp978, tmp980)
		return

	}, 2)

	tmp981 := Call(__e, PrimNS2Value(symdef), symshen_4string_1_6byte, tmp977)

	_ = tmp981

	tmp982 := MakeNative(func(__e *ControlFlow) {
		V4728 := __e.Get(1)
		_ = V4728
		V4729 := __e.Get(2)
		_ = V4729
		V4730 := __e.Get(3)
		_ = V4730
		V4731 := __e.Get(4)
		_ = V4731
		tmp987 := PrimEqual(symshen_4eos, V4730)

		if True == tmp987 {
			__e.Return(V4728)
			return
		} else {
			tmp984 := PrimWriteByte(V4730, V4729)

			_ = tmp984

			tmp985 := Call(__e, PrimNS2Value(symshen_4string_1_6byte), V4728, V4731)

			tmp986 := PrimNumberAdd(V4731, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4write_1chars), V4728, V4729, tmp985, tmp986)
			return

		}

	}, 4)

	tmp988 := Call(__e, PrimNS2Value(symdef), symshen_4write_1chars, tmp982)

	_ = tmp988

	tmp989 := MakeNative(func(__e *ControlFlow) {
		V4732 := __e.Get(1)
		_ = V4732
		V4733 := __e.Get(2)
		_ = V4733
		tmp994 := PrimIsString(V4732)

		if True == tmp994 {
			tmp993 := Call(__e, PrimNS2Value(symshen_4proc_1nl), V4732)

			__e.TailApply(PrimNS2Value(symshen_4mkstr_1l), tmp993, V4733)
			return

		} else {
			tmp991 := PrimCons(V4732, Nil)

			tmp992 := PrimCons(symshen_4proc_1nl, tmp991)

			__e.TailApply(PrimNS2Value(symshen_4mkstr_1r), tmp992, V4733)
			return

		}

	}, 2)

	tmp995 := Call(__e, PrimNS2Value(symdef), symshen_4mkstr, tmp989)

	_ = tmp995

	tmp996 := MakeNative(func(__e *ControlFlow) {
		V4738 := __e.Get(1)
		_ = V4738
		V4739 := __e.Get(2)
		_ = V4739
		tmp1003 := PrimEqual(Nil, V4739)

		if True == tmp1003 {
			__e.Return(V4738)
			return
		} else {
			tmp1002 := PrimIsPair(V4739)

			if True == tmp1002 {
				tmp999 := PrimHead(V4739)

				tmp1000 := Call(__e, PrimNS2Value(symshen_4insert_1l), tmp999, V4738)

				tmp1001 := PrimTail(V4739)

				__e.TailApply(PrimNS2Value(symshen_4mkstr_1l), tmp1000, tmp1001)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-l")))
				return
			}

		}

	}, 2)

	tmp1004 := Call(__e, PrimNS2Value(symdef), symshen_4mkstr_1l, tmp996)

	_ = tmp1004

	tmp1005 := MakeNative(func(__e *ControlFlow) {
		V4746 := __e.Get(1)
		_ = V4746
		V4747 := __e.Get(2)
		_ = V4747
		tmp1143 := PrimEqual(MakeString(""), V4747)

		if True == tmp1143 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp1142 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4747)

			var ifres1129 Obj

			if True == tmp1142 {
				tmp1140 := Call(__e, PrimNS2Value(symhdstr), V4747)

				tmp1141 := PrimEqual(MakeString("~"), tmp1140)

				var ifres1131 Obj

				if True == tmp1141 {
					tmp1138 := PrimTailString(V4747)

					tmp1139 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1138)

					var ifres1133 Obj

					if True == tmp1139 {
						tmp1135 := PrimTailString(V4747)

						tmp1136 := Call(__e, PrimNS2Value(symhdstr), tmp1135)

						tmp1137 := PrimEqual(MakeString("A"), tmp1136)

						var ifres1134 Obj

						if True == tmp1137 {
							ifres1134 = True

						} else {
							ifres1134 = False

						}

						ifres1133 = ifres1134

					} else {
						ifres1133 = False

					}

					var ifres1132 Obj

					if True == ifres1133 {
						ifres1132 = True

					} else {
						ifres1132 = False

					}

					ifres1131 = ifres1132

				} else {
					ifres1131 = False

				}

				var ifres1130 Obj

				if True == ifres1131 {
					ifres1130 = True

				} else {
					ifres1130 = False

				}

				ifres1129 = ifres1130

			} else {
				ifres1129 = False

			}

			if True == ifres1129 {
				tmp1124 := PrimTailString(V4747)

				tmp1125 := PrimTailString(tmp1124)

				tmp1126 := PrimCons(symshen_4a, Nil)

				tmp1127 := PrimCons(tmp1125, tmp1126)

				tmp1128 := PrimCons(V4746, tmp1127)

				__e.Return(PrimCons(symshen_4app, tmp1128))
				return

			} else {
				tmp1123 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4747)

				var ifres1110 Obj

				if True == tmp1123 {
					tmp1121 := Call(__e, PrimNS2Value(symhdstr), V4747)

					tmp1122 := PrimEqual(MakeString("~"), tmp1121)

					var ifres1112 Obj

					if True == tmp1122 {
						tmp1119 := PrimTailString(V4747)

						tmp1120 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1119)

						var ifres1114 Obj

						if True == tmp1120 {
							tmp1116 := PrimTailString(V4747)

							tmp1117 := Call(__e, PrimNS2Value(symhdstr), tmp1116)

							tmp1118 := PrimEqual(MakeString("R"), tmp1117)

							var ifres1115 Obj

							if True == tmp1118 {
								ifres1115 = True

							} else {
								ifres1115 = False

							}

							ifres1114 = ifres1115

						} else {
							ifres1114 = False

						}

						var ifres1113 Obj

						if True == ifres1114 {
							ifres1113 = True

						} else {
							ifres1113 = False

						}

						ifres1112 = ifres1113

					} else {
						ifres1112 = False

					}

					var ifres1111 Obj

					if True == ifres1112 {
						ifres1111 = True

					} else {
						ifres1111 = False

					}

					ifres1110 = ifres1111

				} else {
					ifres1110 = False

				}

				if True == ifres1110 {
					tmp1105 := PrimTailString(V4747)

					tmp1106 := PrimTailString(tmp1105)

					tmp1107 := PrimCons(symshen_4r, Nil)

					tmp1108 := PrimCons(tmp1106, tmp1107)

					tmp1109 := PrimCons(V4746, tmp1108)

					__e.Return(PrimCons(symshen_4app, tmp1109))
					return

				} else {
					tmp1104 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4747)

					var ifres1091 Obj

					if True == tmp1104 {
						tmp1102 := Call(__e, PrimNS2Value(symhdstr), V4747)

						tmp1103 := PrimEqual(MakeString("~"), tmp1102)

						var ifres1093 Obj

						if True == tmp1103 {
							tmp1100 := PrimTailString(V4747)

							tmp1101 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1100)

							var ifres1095 Obj

							if True == tmp1101 {
								tmp1097 := PrimTailString(V4747)

								tmp1098 := Call(__e, PrimNS2Value(symhdstr), tmp1097)

								tmp1099 := PrimEqual(MakeString("S"), tmp1098)

								var ifres1096 Obj

								if True == tmp1099 {
									ifres1096 = True

								} else {
									ifres1096 = False

								}

								ifres1095 = ifres1096

							} else {
								ifres1095 = False

							}

							var ifres1094 Obj

							if True == ifres1095 {
								ifres1094 = True

							} else {
								ifres1094 = False

							}

							ifres1093 = ifres1094

						} else {
							ifres1093 = False

						}

						var ifres1092 Obj

						if True == ifres1093 {
							ifres1092 = True

						} else {
							ifres1092 = False

						}

						ifres1091 = ifres1092

					} else {
						ifres1091 = False

					}

					if True == ifres1091 {
						tmp1086 := PrimTailString(V4747)

						tmp1087 := PrimTailString(tmp1086)

						tmp1088 := PrimCons(symshen_4s, Nil)

						tmp1089 := PrimCons(tmp1087, tmp1088)

						tmp1090 := PrimCons(V4746, tmp1089)

						__e.Return(PrimCons(symshen_4app, tmp1090))
						return

					} else {
						tmp1085 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4747)

						if True == tmp1085 {
							tmp1079 := Call(__e, PrimNS2Value(symhdstr), V4747)

							tmp1080 := PrimTailString(V4747)

							tmp1081 := Call(__e, PrimNS2Value(symshen_4insert_1l), V4746, tmp1080)

							tmp1082 := PrimCons(tmp1081, Nil)

							tmp1083 := PrimCons(tmp1079, tmp1082)

							tmp1084 := PrimCons(symcn, tmp1083)

							__e.TailApply(PrimNS2Value(symshen_4factor_1cn), tmp1084)
							return

						} else {
							tmp1078 := PrimIsPair(V4747)

							var ifres1059 Obj

							if True == tmp1078 {
								tmp1076 := PrimHead(V4747)

								tmp1077 := PrimEqual(symcn, tmp1076)

								var ifres1061 Obj

								if True == tmp1077 {
									tmp1074 := PrimTail(V4747)

									tmp1075 := PrimIsPair(tmp1074)

									var ifres1063 Obj

									if True == tmp1075 {
										tmp1071 := PrimTail(V4747)

										tmp1072 := PrimTail(tmp1071)

										tmp1073 := PrimIsPair(tmp1072)

										var ifres1065 Obj

										if True == tmp1073 {
											tmp1067 := PrimTail(V4747)

											tmp1068 := PrimTail(tmp1067)

											tmp1069 := PrimTail(tmp1068)

											tmp1070 := PrimEqual(Nil, tmp1069)

											var ifres1066 Obj

											if True == tmp1070 {
												ifres1066 = True

											} else {
												ifres1066 = False

											}

											ifres1065 = ifres1066

										} else {
											ifres1065 = False

										}

										var ifres1064 Obj

										if True == ifres1065 {
											ifres1064 = True

										} else {
											ifres1064 = False

										}

										ifres1063 = ifres1064

									} else {
										ifres1063 = False

									}

									var ifres1062 Obj

									if True == ifres1063 {
										ifres1062 = True

									} else {
										ifres1062 = False

									}

									ifres1061 = ifres1062

								} else {
									ifres1061 = False

								}

								var ifres1060 Obj

								if True == ifres1061 {
									ifres1060 = True

								} else {
									ifres1060 = False

								}

								ifres1059 = ifres1060

							} else {
								ifres1059 = False

							}

							if True == ifres1059 {
								tmp1051 := PrimTail(V4747)

								tmp1052 := PrimHead(tmp1051)

								tmp1053 := PrimTail(V4747)

								tmp1054 := PrimTail(tmp1053)

								tmp1055 := PrimHead(tmp1054)

								tmp1056 := Call(__e, PrimNS2Value(symshen_4insert_1l), V4746, tmp1055)

								tmp1057 := PrimCons(tmp1056, Nil)

								tmp1058 := PrimCons(tmp1052, tmp1057)

								__e.Return(PrimCons(symcn, tmp1058))
								return

							} else {
								tmp1050 := PrimIsPair(V4747)

								var ifres1024 Obj

								if True == tmp1050 {
									tmp1048 := PrimHead(V4747)

									tmp1049 := PrimEqual(symshen_4app, tmp1048)

									var ifres1026 Obj

									if True == tmp1049 {
										tmp1046 := PrimTail(V4747)

										tmp1047 := PrimIsPair(tmp1046)

										var ifres1028 Obj

										if True == tmp1047 {
											tmp1043 := PrimTail(V4747)

											tmp1044 := PrimTail(tmp1043)

											tmp1045 := PrimIsPair(tmp1044)

											var ifres1030 Obj

											if True == tmp1045 {
												tmp1039 := PrimTail(V4747)

												tmp1040 := PrimTail(tmp1039)

												tmp1041 := PrimTail(tmp1040)

												tmp1042 := PrimIsPair(tmp1041)

												var ifres1032 Obj

												if True == tmp1042 {
													tmp1034 := PrimTail(V4747)

													tmp1035 := PrimTail(tmp1034)

													tmp1036 := PrimTail(tmp1035)

													tmp1037 := PrimTail(tmp1036)

													tmp1038 := PrimEqual(Nil, tmp1037)

													var ifres1033 Obj

													if True == tmp1038 {
														ifres1033 = True

													} else {
														ifres1033 = False

													}

													ifres1032 = ifres1033

												} else {
													ifres1032 = False

												}

												var ifres1031 Obj

												if True == ifres1032 {
													ifres1031 = True

												} else {
													ifres1031 = False

												}

												ifres1030 = ifres1031

											} else {
												ifres1030 = False

											}

											var ifres1029 Obj

											if True == ifres1030 {
												ifres1029 = True

											} else {
												ifres1029 = False

											}

											ifres1028 = ifres1029

										} else {
											ifres1028 = False

										}

										var ifres1027 Obj

										if True == ifres1028 {
											ifres1027 = True

										} else {
											ifres1027 = False

										}

										ifres1026 = ifres1027

									} else {
										ifres1026 = False

									}

									var ifres1025 Obj

									if True == ifres1026 {
										ifres1025 = True

									} else {
										ifres1025 = False

									}

									ifres1024 = ifres1025

								} else {
									ifres1024 = False

								}

								if True == ifres1024 {
									tmp1013 := PrimTail(V4747)

									tmp1014 := PrimHead(tmp1013)

									tmp1015 := PrimTail(V4747)

									tmp1016 := PrimTail(tmp1015)

									tmp1017 := PrimHead(tmp1016)

									tmp1018 := Call(__e, PrimNS2Value(symshen_4insert_1l), V4746, tmp1017)

									tmp1019 := PrimTail(V4747)

									tmp1020 := PrimTail(tmp1019)

									tmp1021 := PrimTail(tmp1020)

									tmp1022 := PrimCons(tmp1018, tmp1021)

									tmp1023 := PrimCons(tmp1014, tmp1022)

									__e.Return(PrimCons(symshen_4app, tmp1023))
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

	tmp1144 := Call(__e, PrimNS2Value(symdef), symshen_4insert_1l, tmp1005)

	_ = tmp1144

	tmp1145 := MakeNative(func(__e *ControlFlow) {
		V4748 := __e.Get(1)
		_ = V4748
		tmp1230 := PrimIsPair(V4748)

		var ifres1161 Obj

		if True == tmp1230 {
			tmp1228 := PrimHead(V4748)

			tmp1229 := PrimEqual(symcn, tmp1228)

			var ifres1163 Obj

			if True == tmp1229 {
				tmp1226 := PrimTail(V4748)

				tmp1227 := PrimIsPair(tmp1226)

				var ifres1165 Obj

				if True == tmp1227 {
					tmp1223 := PrimTail(V4748)

					tmp1224 := PrimTail(tmp1223)

					tmp1225 := PrimIsPair(tmp1224)

					var ifres1167 Obj

					if True == tmp1225 {
						tmp1219 := PrimTail(V4748)

						tmp1220 := PrimTail(tmp1219)

						tmp1221 := PrimHead(tmp1220)

						tmp1222 := PrimIsPair(tmp1221)

						var ifres1169 Obj

						if True == tmp1222 {
							tmp1214 := PrimTail(V4748)

							tmp1215 := PrimTail(tmp1214)

							tmp1216 := PrimHead(tmp1215)

							tmp1217 := PrimHead(tmp1216)

							tmp1218 := PrimEqual(symcn, tmp1217)

							var ifres1171 Obj

							if True == tmp1218 {
								tmp1209 := PrimTail(V4748)

								tmp1210 := PrimTail(tmp1209)

								tmp1211 := PrimHead(tmp1210)

								tmp1212 := PrimTail(tmp1211)

								tmp1213 := PrimIsPair(tmp1212)

								var ifres1173 Obj

								if True == tmp1213 {
									tmp1203 := PrimTail(V4748)

									tmp1204 := PrimTail(tmp1203)

									tmp1205 := PrimHead(tmp1204)

									tmp1206 := PrimTail(tmp1205)

									tmp1207 := PrimTail(tmp1206)

									tmp1208 := PrimIsPair(tmp1207)

									var ifres1175 Obj

									if True == tmp1208 {
										tmp1196 := PrimTail(V4748)

										tmp1197 := PrimTail(tmp1196)

										tmp1198 := PrimHead(tmp1197)

										tmp1199 := PrimTail(tmp1198)

										tmp1200 := PrimTail(tmp1199)

										tmp1201 := PrimTail(tmp1200)

										tmp1202 := PrimEqual(Nil, tmp1201)

										var ifres1177 Obj

										if True == tmp1202 {
											tmp1192 := PrimTail(V4748)

											tmp1193 := PrimTail(tmp1192)

											tmp1194 := PrimTail(tmp1193)

											tmp1195 := PrimEqual(Nil, tmp1194)

											var ifres1179 Obj

											if True == tmp1195 {
												tmp1189 := PrimTail(V4748)

												tmp1190 := PrimHead(tmp1189)

												tmp1191 := PrimIsString(tmp1190)

												var ifres1181 Obj

												if True == tmp1191 {
													tmp1183 := PrimTail(V4748)

													tmp1184 := PrimTail(tmp1183)

													tmp1185 := PrimHead(tmp1184)

													tmp1186 := PrimTail(tmp1185)

													tmp1187 := PrimHead(tmp1186)

													tmp1188 := PrimIsString(tmp1187)

													var ifres1182 Obj

													if True == tmp1188 {
														ifres1182 = True

													} else {
														ifres1182 = False

													}

													ifres1181 = ifres1182

												} else {
													ifres1181 = False

												}

												var ifres1180 Obj

												if True == ifres1181 {
													ifres1180 = True

												} else {
													ifres1180 = False

												}

												ifres1179 = ifres1180

											} else {
												ifres1179 = False

											}

											var ifres1178 Obj

											if True == ifres1179 {
												ifres1178 = True

											} else {
												ifres1178 = False

											}

											ifres1177 = ifres1178

										} else {
											ifres1177 = False

										}

										var ifres1176 Obj

										if True == ifres1177 {
											ifres1176 = True

										} else {
											ifres1176 = False

										}

										ifres1175 = ifres1176

									} else {
										ifres1175 = False

									}

									var ifres1174 Obj

									if True == ifres1175 {
										ifres1174 = True

									} else {
										ifres1174 = False

									}

									ifres1173 = ifres1174

								} else {
									ifres1173 = False

								}

								var ifres1172 Obj

								if True == ifres1173 {
									ifres1172 = True

								} else {
									ifres1172 = False

								}

								ifres1171 = ifres1172

							} else {
								ifres1171 = False

							}

							var ifres1170 Obj

							if True == ifres1171 {
								ifres1170 = True

							} else {
								ifres1170 = False

							}

							ifres1169 = ifres1170

						} else {
							ifres1169 = False

						}

						var ifres1168 Obj

						if True == ifres1169 {
							ifres1168 = True

						} else {
							ifres1168 = False

						}

						ifres1167 = ifres1168

					} else {
						ifres1167 = False

					}

					var ifres1166 Obj

					if True == ifres1167 {
						ifres1166 = True

					} else {
						ifres1166 = False

					}

					ifres1165 = ifres1166

				} else {
					ifres1165 = False

				}

				var ifres1164 Obj

				if True == ifres1165 {
					ifres1164 = True

				} else {
					ifres1164 = False

				}

				ifres1163 = ifres1164

			} else {
				ifres1163 = False

			}

			var ifres1162 Obj

			if True == ifres1163 {
				ifres1162 = True

			} else {
				ifres1162 = False

			}

			ifres1161 = ifres1162

		} else {
			ifres1161 = False

		}

		if True == ifres1161 {
			tmp1147 := PrimTail(V4748)

			tmp1148 := PrimHead(tmp1147)

			tmp1149 := PrimTail(V4748)

			tmp1150 := PrimTail(tmp1149)

			tmp1151 := PrimHead(tmp1150)

			tmp1152 := PrimTail(tmp1151)

			tmp1153 := PrimHead(tmp1152)

			tmp1154 := PrimStringConcat(tmp1148, tmp1153)

			tmp1155 := PrimTail(V4748)

			tmp1156 := PrimTail(tmp1155)

			tmp1157 := PrimHead(tmp1156)

			tmp1158 := PrimTail(tmp1157)

			tmp1159 := PrimTail(tmp1158)

			tmp1160 := PrimCons(tmp1154, tmp1159)

			__e.Return(PrimCons(symcn, tmp1160))
			return

		} else {
			__e.Return(V4748)
			return
		}

	}, 1)

	tmp1231 := Call(__e, PrimNS2Value(symdef), symshen_4factor_1cn, tmp1145)

	_ = tmp1231

	tmp1232 := MakeNative(func(__e *ControlFlow) {
		V4751 := __e.Get(1)
		_ = V4751
		tmp1258 := PrimEqual(MakeString(""), V4751)

		if True == tmp1258 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp1257 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4751)

			var ifres1244 Obj

			if True == tmp1257 {
				tmp1255 := Call(__e, PrimNS2Value(symhdstr), V4751)

				tmp1256 := PrimEqual(MakeString("~"), tmp1255)

				var ifres1246 Obj

				if True == tmp1256 {
					tmp1253 := PrimTailString(V4751)

					tmp1254 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1253)

					var ifres1248 Obj

					if True == tmp1254 {
						tmp1250 := PrimTailString(V4751)

						tmp1251 := Call(__e, PrimNS2Value(symhdstr), tmp1250)

						tmp1252 := PrimEqual(MakeString("%"), tmp1251)

						var ifres1249 Obj

						if True == tmp1252 {
							ifres1249 = True

						} else {
							ifres1249 = False

						}

						ifres1248 = ifres1249

					} else {
						ifres1248 = False

					}

					var ifres1247 Obj

					if True == ifres1248 {
						ifres1247 = True

					} else {
						ifres1247 = False

					}

					ifres1246 = ifres1247

				} else {
					ifres1246 = False

				}

				var ifres1245 Obj

				if True == ifres1246 {
					ifres1245 = True

				} else {
					ifres1245 = False

				}

				ifres1244 = ifres1245

			} else {
				ifres1244 = False

			}

			if True == ifres1244 {
				tmp1240 := PrimNumberToString(MakeNumber(10))

				tmp1241 := PrimTailString(V4751)

				tmp1242 := PrimTailString(tmp1241)

				tmp1243 := Call(__e, PrimNS2Value(symshen_4proc_1nl), tmp1242)

				__e.Return(PrimStringConcat(tmp1240, tmp1243))
				return

			} else {
				tmp1239 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4751)

				if True == tmp1239 {
					tmp1236 := Call(__e, PrimNS2Value(symhdstr), V4751)

					tmp1237 := PrimTailString(V4751)

					tmp1238 := Call(__e, PrimNS2Value(symshen_4proc_1nl), tmp1237)

					__e.Return(PrimStringConcat(tmp1236, tmp1238))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.proc-nl")))
					return
				}

			}

		}

	}, 1)

	tmp1259 := Call(__e, PrimNS2Value(symdef), symshen_4proc_1nl, tmp1232)

	_ = tmp1259

	tmp1260 := MakeNative(func(__e *ControlFlow) {
		V4756 := __e.Get(1)
		_ = V4756
		V4757 := __e.Get(2)
		_ = V4757
		tmp1269 := PrimEqual(Nil, V4757)

		if True == tmp1269 {
			__e.Return(V4756)
			return
		} else {
			tmp1268 := PrimIsPair(V4757)

			if True == tmp1268 {
				tmp1263 := PrimHead(V4757)

				tmp1264 := PrimCons(V4756, Nil)

				tmp1265 := PrimCons(tmp1263, tmp1264)

				tmp1266 := PrimCons(symshen_4insert, tmp1265)

				tmp1267 := PrimTail(V4757)

				__e.TailApply(PrimNS2Value(symshen_4mkstr_1r), tmp1266, tmp1267)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-r")))
				return
			}

		}

	}, 2)

	tmp1270 := Call(__e, PrimNS2Value(symdef), symshen_4mkstr_1r, tmp1260)

	_ = tmp1270

	tmp1271 := MakeNative(func(__e *ControlFlow) {
		V4758 := __e.Get(1)
		_ = V4758
		V4759 := __e.Get(2)
		_ = V4759
		__e.TailApply(PrimNS2Value(symshen_4insert_1h), V4758, V4759, MakeString(""))
		return
	}, 2)

	tmp1272 := Call(__e, PrimNS2Value(symdef), symshen_4insert, tmp1271)

	_ = tmp1272

	tmp1273 := MakeNative(func(__e *ControlFlow) {
		V4768 := __e.Get(1)
		_ = V4768
		V4769 := __e.Get(2)
		_ = V4769
		V4770 := __e.Get(3)
		_ = V4770
		tmp1334 := PrimEqual(MakeString(""), V4769)

		if True == tmp1334 {
			__e.Return(V4770)
			return
		} else {
			tmp1333 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4769)

			var ifres1320 Obj

			if True == tmp1333 {
				tmp1331 := Call(__e, PrimNS2Value(symhdstr), V4769)

				tmp1332 := PrimEqual(MakeString("~"), tmp1331)

				var ifres1322 Obj

				if True == tmp1332 {
					tmp1329 := PrimTailString(V4769)

					tmp1330 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1329)

					var ifres1324 Obj

					if True == tmp1330 {
						tmp1326 := PrimTailString(V4769)

						tmp1327 := Call(__e, PrimNS2Value(symhdstr), tmp1326)

						tmp1328 := PrimEqual(MakeString("A"), tmp1327)

						var ifres1325 Obj

						if True == tmp1328 {
							ifres1325 = True

						} else {
							ifres1325 = False

						}

						ifres1324 = ifres1325

					} else {
						ifres1324 = False

					}

					var ifres1323 Obj

					if True == ifres1324 {
						ifres1323 = True

					} else {
						ifres1323 = False

					}

					ifres1322 = ifres1323

				} else {
					ifres1322 = False

				}

				var ifres1321 Obj

				if True == ifres1322 {
					ifres1321 = True

				} else {
					ifres1321 = False

				}

				ifres1320 = ifres1321

			} else {
				ifres1320 = False

			}

			if True == ifres1320 {
				tmp1317 := PrimTailString(V4769)

				tmp1318 := PrimTailString(tmp1317)

				tmp1319 := Call(__e, PrimNS2Value(symshen_4app), V4768, tmp1318, symshen_4a)

				__e.Return(PrimStringConcat(V4770, tmp1319))
				return

			} else {
				tmp1316 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4769)

				var ifres1303 Obj

				if True == tmp1316 {
					tmp1314 := Call(__e, PrimNS2Value(symhdstr), V4769)

					tmp1315 := PrimEqual(MakeString("~"), tmp1314)

					var ifres1305 Obj

					if True == tmp1315 {
						tmp1312 := PrimTailString(V4769)

						tmp1313 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1312)

						var ifres1307 Obj

						if True == tmp1313 {
							tmp1309 := PrimTailString(V4769)

							tmp1310 := Call(__e, PrimNS2Value(symhdstr), tmp1309)

							tmp1311 := PrimEqual(MakeString("R"), tmp1310)

							var ifres1308 Obj

							if True == tmp1311 {
								ifres1308 = True

							} else {
								ifres1308 = False

							}

							ifres1307 = ifres1308

						} else {
							ifres1307 = False

						}

						var ifres1306 Obj

						if True == ifres1307 {
							ifres1306 = True

						} else {
							ifres1306 = False

						}

						ifres1305 = ifres1306

					} else {
						ifres1305 = False

					}

					var ifres1304 Obj

					if True == ifres1305 {
						ifres1304 = True

					} else {
						ifres1304 = False

					}

					ifres1303 = ifres1304

				} else {
					ifres1303 = False

				}

				if True == ifres1303 {
					tmp1300 := PrimTailString(V4769)

					tmp1301 := PrimTailString(tmp1300)

					tmp1302 := Call(__e, PrimNS2Value(symshen_4app), V4768, tmp1301, symshen_4r)

					__e.Return(PrimStringConcat(V4770, tmp1302))
					return

				} else {
					tmp1299 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4769)

					var ifres1286 Obj

					if True == tmp1299 {
						tmp1297 := Call(__e, PrimNS2Value(symhdstr), V4769)

						tmp1298 := PrimEqual(MakeString("~"), tmp1297)

						var ifres1288 Obj

						if True == tmp1298 {
							tmp1295 := PrimTailString(V4769)

							tmp1296 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp1295)

							var ifres1290 Obj

							if True == tmp1296 {
								tmp1292 := PrimTailString(V4769)

								tmp1293 := Call(__e, PrimNS2Value(symhdstr), tmp1292)

								tmp1294 := PrimEqual(MakeString("S"), tmp1293)

								var ifres1291 Obj

								if True == tmp1294 {
									ifres1291 = True

								} else {
									ifres1291 = False

								}

								ifres1290 = ifres1291

							} else {
								ifres1290 = False

							}

							var ifres1289 Obj

							if True == ifres1290 {
								ifres1289 = True

							} else {
								ifres1289 = False

							}

							ifres1288 = ifres1289

						} else {
							ifres1288 = False

						}

						var ifres1287 Obj

						if True == ifres1288 {
							ifres1287 = True

						} else {
							ifres1287 = False

						}

						ifres1286 = ifres1287

					} else {
						ifres1286 = False

					}

					if True == ifres1286 {
						tmp1283 := PrimTailString(V4769)

						tmp1284 := PrimTailString(tmp1283)

						tmp1285 := Call(__e, PrimNS2Value(symshen_4app), V4768, tmp1284, symshen_4s)

						__e.Return(PrimStringConcat(V4770, tmp1285))
						return

					} else {
						tmp1282 := Call(__e, PrimNS2Value(symshen_4_7string_2), V4769)

						if True == tmp1282 {
							tmp1279 := PrimTailString(V4769)

							tmp1280 := Call(__e, PrimNS2Value(symhdstr), V4769)

							tmp1281 := PrimStringConcat(V4770, tmp1280)

							__e.TailApply(PrimNS2Value(symshen_4insert_1h), V4768, tmp1279, tmp1281)
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

	tmp1335 := Call(__e, PrimNS2Value(symdef), symshen_4insert_1h, tmp1273)

	_ = tmp1335

	tmp1336 := MakeNative(func(__e *ControlFlow) {
		V4771 := __e.Get(1)
		_ = V4771
		V4772 := __e.Get(2)
		_ = V4772
		V4773 := __e.Get(3)
		_ = V4773
		tmp1337 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), V4771, V4773)

		__e.Return(PrimStringConcat(tmp1337, V4772))
		return

	}, 3)

	tmp1338 := Call(__e, PrimNS2Value(symdef), symshen_4app, tmp1336)

	_ = tmp1338

	tmp1339 := MakeNative(func(__e *ControlFlow) {
		V4777 := __e.Get(1)
		_ = V4777
		V4778 := __e.Get(2)
		_ = V4778
		tmp1347 := Call(__e, PrimNS2Value(symfail))

		tmp1348 := PrimEqual(V4777, tmp1347)

		if True == tmp1348 {
			__e.Return(MakeString("..."))
			return
		} else {
			tmp1346 := Call(__e, PrimNS2Value(symshen_4list_2), V4777)

			if True == tmp1346 {
				__e.TailApply(PrimNS2Value(symshen_4list_1_6str), V4777, V4778)
				return
			} else {
				tmp1345 := PrimIsString(V4777)

				if True == tmp1345 {
					__e.TailApply(PrimNS2Value(symshen_4str_1_6str), V4777, V4778)
					return
				} else {
					tmp1344 := PrimIsVector(V4777)

					if True == tmp1344 {
						__e.TailApply(PrimNS2Value(symshen_4vector_1_6str), V4777, V4778)
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4atom_1_6str), V4777)
						return
					}

				}

			}

		}

	}, 2)

	tmp1349 := Call(__e, PrimNS2Value(symdef), symshen_4arg_1_6str, tmp1339)

	_ = tmp1349

	tmp1350 := MakeNative(func(__e *ControlFlow) {
		V4779 := __e.Get(1)
		_ = V4779
		V4780 := __e.Get(2)
		_ = V4780
		tmp1358 := PrimEqual(symshen_4r, V4780)

		if True == tmp1358 {
			tmp1355 := Call(__e, PrimNS2Value(symshen_4maxseq))

			tmp1356 := Call(__e, PrimNS2Value(symshen_4iter_1list), V4779, symshen_4r, tmp1355)

			tmp1357 := Call(__e, PrimNS2Value(sym_8s), tmp1356, MakeString(")"))

			__e.TailApply(PrimNS2Value(sym_8s), MakeString("("), tmp1357)
			return

		} else {
			tmp1352 := Call(__e, PrimNS2Value(symshen_4maxseq))

			tmp1353 := Call(__e, PrimNS2Value(symshen_4iter_1list), V4779, V4780, tmp1352)

			tmp1354 := Call(__e, PrimNS2Value(sym_8s), tmp1353, MakeString("]"))

			__e.TailApply(PrimNS2Value(sym_8s), MakeString("["), tmp1354)
			return

		}

	}, 2)

	tmp1359 := Call(__e, PrimNS2Value(symdef), symshen_4list_1_6str, tmp1350)

	_ = tmp1359

	tmp1360 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dmaximum_1print_1sequence_1size_d))
		return
	}, 0)

	tmp1361 := Call(__e, PrimNS2Value(symdef), symshen_4maxseq, tmp1360)

	_ = tmp1361

	tmp1362 := MakeNative(func(__e *ControlFlow) {
		V4791 := __e.Get(1)
		_ = V4791
		V4792 := __e.Get(2)
		_ = V4792
		V4793 := __e.Get(3)
		_ = V4793
		tmp1383 := PrimEqual(Nil, V4791)

		if True == tmp1383 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp1382 := PrimEqual(MakeNumber(0), V4793)

			if True == tmp1382 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				tmp1381 := PrimIsPair(V4791)

				var ifres1377 Obj

				if True == tmp1381 {
					tmp1379 := PrimTail(V4791)

					tmp1380 := PrimEqual(Nil, tmp1379)

					var ifres1378 Obj

					if True == tmp1380 {
						ifres1378 = True

					} else {
						ifres1378 = False

					}

					ifres1377 = ifres1378

				} else {
					ifres1377 = False

				}

				if True == ifres1377 {
					tmp1376 := PrimHead(V4791)

					__e.TailApply(PrimNS2Value(symshen_4arg_1_6str), tmp1376, V4792)
					return

				} else {
					tmp1375 := PrimIsPair(V4791)

					if True == tmp1375 {
						tmp1369 := PrimHead(V4791)

						tmp1370 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), tmp1369, V4792)

						tmp1371 := PrimTail(V4791)

						tmp1372 := PrimNumberSubtract(V4793, MakeNumber(1))

						tmp1373 := Call(__e, PrimNS2Value(symshen_4iter_1list), tmp1371, V4792, tmp1372)

						tmp1374 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp1373)

						__e.TailApply(PrimNS2Value(sym_8s), tmp1370, tmp1374)
						return

					} else {
						tmp1367 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), V4791, V4792)

						tmp1368 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp1367)

						__e.TailApply(PrimNS2Value(sym_8s), MakeString("|"), tmp1368)
						return

					}

				}

			}

		}

	}, 3)

	tmp1384 := Call(__e, PrimNS2Value(symdef), symshen_4iter_1list, tmp1362)

	_ = tmp1384

	tmp1385 := MakeNative(func(__e *ControlFlow) {
		V4796 := __e.Get(1)
		_ = V4796
		V4797 := __e.Get(2)
		_ = V4797
		tmp1390 := PrimEqual(symshen_4a, V4797)

		if True == tmp1390 {
			__e.Return(V4796)
			return
		} else {
			tmp1387 := PrimNumberToString(MakeNumber(34))

			tmp1388 := PrimNumberToString(MakeNumber(34))

			tmp1389 := Call(__e, PrimNS2Value(sym_8s), V4796, tmp1388)

			__e.TailApply(PrimNS2Value(sym_8s), tmp1387, tmp1389)
			return

		}

	}, 2)

	tmp1391 := Call(__e, PrimNS2Value(symdef), symshen_4str_1_6str, tmp1385)

	_ = tmp1391

	tmp1392 := MakeNative(func(__e *ControlFlow) {
		V4798 := __e.Get(1)
		_ = V4798
		V4799 := __e.Get(2)
		_ = V4799
		tmp1405 := Call(__e, PrimNS2Value(symshen_4print_1vector_2), V4798)

		if True == tmp1405 {
			tmp1403 := PrimVectorGet(V4798, MakeNumber(0))

			tmp1404 := Call(__e, PrimNS2Value(symfn), tmp1403)

			__e.TailApply(tmp1404, V4798)
			return

		} else {
			tmp1402 := Call(__e, PrimNS2Value(symvector_2), V4798)

			if True == tmp1402 {
				tmp1399 := Call(__e, PrimNS2Value(symshen_4maxseq))

				tmp1400 := Call(__e, PrimNS2Value(symshen_4iter_1vector), V4798, MakeNumber(1), V4799, tmp1399)

				tmp1401 := Call(__e, PrimNS2Value(sym_8s), tmp1400, MakeString(">"))

				__e.TailApply(PrimNS2Value(sym_8s), MakeString("<"), tmp1401)
				return

			} else {
				tmp1395 := Call(__e, PrimNS2Value(symshen_4maxseq))

				tmp1396 := Call(__e, PrimNS2Value(symshen_4iter_1vector), V4798, MakeNumber(0), V4799, tmp1395)

				tmp1397 := Call(__e, PrimNS2Value(sym_8s), tmp1396, MakeString(">>"))

				tmp1398 := Call(__e, PrimNS2Value(sym_8s), MakeString("<"), tmp1397)

				__e.TailApply(PrimNS2Value(sym_8s), MakeString("<"), tmp1398)
				return

			}

		}

	}, 2)

	tmp1406 := Call(__e, PrimNS2Value(symdef), symshen_4vector_1_6str, tmp1392)

	_ = tmp1406

	tmp1407 := MakeNative(func(__e *ControlFlow) {
		V4800 := __e.Get(1)
		_ = V4800
		tmp1408 := MakeNative(func(__e *ControlFlow) {
			Zero := __e.Get(1)
			_ = Zero
			tmp1415 := PrimEqual(Zero, symshen_4tuple)

			if True == tmp1415 {
				__e.Return(True)
				return
			} else {
				tmp1414 := PrimEqual(Zero, symshen_4pvar)

				if True == tmp1414 {
					__e.Return(True)
					return
				} else {
					tmp1412 := PrimIsNumber(Zero)

					tmp1413 := PrimNot(tmp1412)

					if True == tmp1413 {
						__e.TailApply(PrimNS2Value(symshen_4fbound_2), Zero)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}

		}, 1)

		tmp1416 := PrimVectorGet(V4800, MakeNumber(0))

		__e.TailApply(tmp1408, tmp1416)
		return

	}, 1)

	tmp1417 := Call(__e, PrimNS2Value(symdef), symshen_4print_1vector_2, tmp1407)

	_ = tmp1417

	tmp1418 := MakeNative(func(__e *ControlFlow) {
		V4801 := __e.Get(1)
		_ = V4801
		tmp1419 := Call(__e, PrimNS2Value(symarity), V4801)

		tmp1420 := PrimEqual(tmp1419, MakeNumber(-1))

		__e.Return(PrimNot(tmp1420))
		return

	}, 1)

	tmp1421 := Call(__e, PrimNS2Value(symdef), symshen_4fbound_2, tmp1418)

	_ = tmp1421

	tmp1422 := MakeNative(func(__e *ControlFlow) {
		V4802 := __e.Get(1)
		_ = V4802
		tmp1423 := PrimVectorGet(V4802, MakeNumber(1))

		tmp1424 := PrimVectorGet(V4802, MakeNumber(2))

		tmp1425 := Call(__e, PrimNS2Value(symshen_4app), tmp1424, MakeString(")"), symshen_4s)

		tmp1426 := PrimStringConcat(MakeString(" "), tmp1425)

		tmp1427 := Call(__e, PrimNS2Value(symshen_4app), tmp1423, tmp1426, symshen_4s)

		__e.Return(PrimStringConcat(MakeString("(@p "), tmp1427))
		return

	}, 1)

	tmp1428 := Call(__e, PrimNS2Value(symdef), symshen_4tuple, tmp1422)

	_ = tmp1428

	tmp1429 := MakeNative(func(__e *ControlFlow) {
		V4809 := __e.Get(1)
		_ = V4809
		V4810 := __e.Get(2)
		_ = V4810
		V4811 := __e.Get(3)
		_ = V4811
		V4812 := __e.Get(4)
		_ = V4812
		tmp1449 := PrimEqual(MakeNumber(0), V4812)

		if True == tmp1449 {
			__e.Return(MakeString("... etc"))
			return
		} else {
			tmp1431 := MakeNative(func(__e *ControlFlow) {
				Item := __e.Get(1)
				_ = Item
				tmp1432 := MakeNative(func(__e *ControlFlow) {
					Next := __e.Get(1)
					_ = Next
					tmp1441 := PrimEqual(Item, symshen_4out_1of_1bounds)

					if True == tmp1441 {
						__e.Return(MakeString(""))
						return
					} else {
						tmp1440 := PrimEqual(Next, symshen_4out_1of_1bounds)

						if True == tmp1440 {
							__e.TailApply(PrimNS2Value(symshen_4arg_1_6str), Item, V4811)
							return
						} else {
							tmp1435 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), Item, V4811)

							tmp1436 := PrimNumberAdd(V4810, MakeNumber(1))

							tmp1437 := PrimNumberSubtract(V4812, MakeNumber(1))

							tmp1438 := Call(__e, PrimNS2Value(symshen_4iter_1vector), V4809, tmp1436, V4811, tmp1437)

							tmp1439 := Call(__e, PrimNS2Value(sym_8s), MakeString(" "), tmp1438)

							__e.TailApply(PrimNS2Value(sym_8s), tmp1435, tmp1439)
							return

						}

					}

				}, 1)

				tmp1442 := MakeNative(func(__e *ControlFlow) {
					tmp1443 := PrimNumberAdd(V4810, MakeNumber(1))

					__e.Return(PrimVectorGet(V4809, tmp1443))
					return

				}, 0)

				tmp1444 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4out_1of_1bounds)
					return
				}, 1)

				tmp1445 := Call(__e, PrimNS1Value(symtry_1catch), tmp1442, tmp1444)

				__e.TailApply(tmp1432, tmp1445)
				return

			}, 1)

			tmp1446 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V4809, V4810))
				return
			}, 0)

			tmp1447 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4out_1of_1bounds)
				return
			}, 1)

			tmp1448 := Call(__e, PrimNS1Value(symtry_1catch), tmp1446, tmp1447)

			__e.TailApply(tmp1431, tmp1448)
			return

		}

	}, 4)

	tmp1450 := Call(__e, PrimNS2Value(symdef), symshen_4iter_1vector, tmp1429)

	_ = tmp1450

	tmp1451 := MakeNative(func(__e *ControlFlow) {
		V4813 := __e.Get(1)
		_ = V4813
		tmp1452 := MakeNative(func(__e *ControlFlow) {
			__e.Return(PrimStr(V4813))
			return
		}, 0)

		tmp1453 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.TailApply(PrimNS2Value(symshen_4funexstring))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp1452, tmp1453)
		return

	}, 1)

	tmp1454 := Call(__e, PrimNS2Value(symdef), symshen_4atom_1_6str, tmp1451)

	_ = tmp1454

	tmp1455 := MakeNative(func(__e *ControlFlow) {
		tmp1456 := PrimIntern(MakeString("x"))

		tmp1457 := Call(__e, PrimNS2Value(symgensym), tmp1456)

		tmp1458 := Call(__e, PrimNS2Value(symshen_4arg_1_6str), tmp1457, symshen_4a)

		tmp1459 := Call(__e, PrimNS2Value(sym_8s), tmp1458, MakeString("\x11"))

		tmp1460 := Call(__e, PrimNS2Value(sym_8s), MakeString("e"), tmp1459)

		tmp1461 := Call(__e, PrimNS2Value(sym_8s), MakeString("n"), tmp1460)

		tmp1462 := Call(__e, PrimNS2Value(sym_8s), MakeString("u"), tmp1461)

		tmp1463 := Call(__e, PrimNS2Value(sym_8s), MakeString("f"), tmp1462)

		__e.TailApply(PrimNS2Value(sym_8s), MakeString("\x10"), tmp1463)
		return

	}, 0)

	tmp1464 := Call(__e, PrimNS2Value(symdef), symshen_4funexstring, tmp1455)

	_ = tmp1464

	tmp1465 := MakeNative(func(__e *ControlFlow) {
		V4814 := __e.Get(1)
		_ = V4814
		tmp1469 := Call(__e, PrimNS2Value(symempty_2), V4814)

		if True == tmp1469 {
			__e.Return(True)
			return
		} else {
			tmp1468 := PrimIsPair(V4814)

			if True == tmp1468 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	__e.TailApply(PrimNS2Value(symdef), symshen_4list_2, tmp1465)
	return

}, 0)
