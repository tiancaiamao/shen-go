package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1072 := MakeNative(func(__e *ControlFlow) {
V3354 := __e.Get(1)
_ = V3354
__e.TailApply(V3354)
return
}, 1)

tmp1073 := Call(__e, ns2_1set, symthaw, tmp1072)


_ = tmp1073

tmp1074 := MakeNative(func(__e *ControlFlow) {
V3355 := __e.Get(1)
_ = V3355
tmp1075 := Call(__e, PrimFunc(symmacroexpand), V3355)


tmp1076 := Call(__e, PrimFunc(symshen_4find_1types), V3355)


tmp1077 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1075, tmp1076)


tmp1078 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1077)


__e.TailApply(PrimFunc(symeval_1kl), tmp1078)
return


}, 1)

tmp1079 := Call(__e, ns2_1set, symeval, tmp1074)


_ = tmp1079

tmp1080 := MakeNative(func(__e *ControlFlow) {
V3356 := __e.Get(1)
_ = V3356
tmp1087 := PrimEqual(symnull, V3356)

if True == tmp1087 {
__e.Return(Nil)
return
} else {
tmp1081 := MakeNative(func(__e *ControlFlow) {
tmp1082 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3356, symshen_4external_1symbols, tmp1082)
return


}, 0)

tmp1083 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
tmp1084 := Call(__e, PrimFunc(symshen_4app), V3356, MakeString(" does not exist.\n;"), symshen_4a)


tmp1085 := PrimStringConcat(MakeString("package "), tmp1084)

__e.Return(PrimSimpleError(tmp1085))
return


}, 1)

__e.TailApply(try_1catch, tmp1081, tmp1083)
return


}


}, 1)

tmp1088 := Call(__e, ns2_1set, symexternal, tmp1080)


_ = tmp1088

tmp1089 := MakeNative(func(__e *ControlFlow) {
V3357 := __e.Get(1)
_ = V3357
tmp1096 := PrimEqual(symnull, V3357)

if True == tmp1096 {
__e.Return(Nil)
return
} else {
tmp1090 := MakeNative(func(__e *ControlFlow) {
tmp1091 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3357, symshen_4internal_1symbols, tmp1091)
return


}, 0)

tmp1092 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
tmp1093 := Call(__e, PrimFunc(symshen_4app), V3357, MakeString(" does not exist.\n;"), symshen_4a)


tmp1094 := PrimStringConcat(MakeString("package "), tmp1093)

__e.Return(PrimSimpleError(tmp1094))
return


}, 1)

__e.TailApply(try_1catch, tmp1090, tmp1092)
return


}


}, 1)

tmp1097 := Call(__e, ns2_1set, syminternal, tmp1089)


_ = tmp1097

tmp1098 := MakeNative(func(__e *ControlFlow) {
V3358 := __e.Get(1)
_ = V3358
V3359 := __e.Get(2)
_ = V3359
tmp1100 := Call(__e, V3358, V3359)


if True == tmp1100 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3359)
return
}


}, 2)

tmp1101 := Call(__e, ns2_1set, symfail_1if, tmp1098)


_ = tmp1101

tmp1102 := MakeNative(func(__e *ControlFlow) {
V3360 := __e.Get(1)
_ = V3360
V3361 := __e.Get(2)
_ = V3361
__e.Return(PrimStringConcat(V3360, V3361))
return
}, 2)

tmp1103 := Call(__e, ns2_1set, sym_8s, tmp1102)


_ = tmp1103

tmp1104 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtc_d))
return
}, 0)

tmp1105 := Call(__e, ns2_1set, symtc_2, tmp1104)


_ = tmp1105

tmp1106 := MakeNative(func(__e *ControlFlow) {
V3362 := __e.Get(1)
_ = V3362
tmp1107 := MakeNative(func(__e *ControlFlow) {
tmp1108 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3362, symshen_4source, tmp1108)
return


}, 0)

tmp1109 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
tmp1110 := Call(__e, PrimFunc(symshen_4app), V3362, MakeString(" not found.\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp1110))
return


}, 1)

__e.TailApply(try_1catch, tmp1107, tmp1109)
return


}, 1)

tmp1111 := Call(__e, ns2_1set, symps, tmp1106)


_ = tmp1111

tmp1112 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstinput_d))
return
}, 0)

tmp1113 := Call(__e, ns2_1set, symstinput, tmp1112)


_ = tmp1113

tmp1114 := MakeNative(func(__e *ControlFlow) {
V3363 := __e.Get(1)
_ = V3363
tmp1115 := MakeNative(func(__e *ControlFlow) {
Vector := __e.Get(1)
_ = Vector
tmp1116 := MakeNative(func(__e *ControlFlow) {
ZeroStamp := __e.Get(1)
_ = ZeroStamp
tmp1117 := MakeNative(func(__e *ControlFlow) {
Standard := __e.Get(1)
_ = Standard
__e.Return(Standard)
return
}, 1)

tmp1121 := PrimEqual(V3363, MakeNumber(0))

var ifres1118 Obj

if True == tmp1121 {
ifres1118 = ZeroStamp


} else {
tmp1119 := Call(__e, PrimFunc(symfail))


tmp1120 := Call(__e, PrimFunc(symshen_4fillvector), ZeroStamp, MakeNumber(1), V3363, tmp1119)


ifres1118 = tmp1120


}

__e.TailApply(tmp1117, ifres1118)
return


}, 1)

tmp1122 := PrimVectorSet(Vector, MakeNumber(0), V3363)

__e.TailApply(tmp1116, tmp1122)
return


}, 1)

tmp1123 := PrimNumberAdd(V3363, MakeNumber(1))

tmp1124 := PrimAbsvector(tmp1123)

__e.TailApply(tmp1115, tmp1124)
return


}, 1)

tmp1125 := Call(__e, ns2_1set, symvector, tmp1114)


_ = tmp1125

tmp1126 := MakeNative(func(__e *ControlFlow) {
V3365 := __e.Get(1)
_ = V3365
V3366 := __e.Get(2)
_ = V3366
V3367 := __e.Get(3)
_ = V3367
V3368 := __e.Get(4)
_ = V3368
tmp1130 := PrimEqual(V3366, V3367)

if True == tmp1130 {
__e.Return(PrimVectorSet(V3365, V3367, V3368))
return
} else {
tmp1127 := PrimVectorSet(V3365, V3366, V3368)

tmp1128 := PrimNumberAdd(MakeNumber(1), V3366)

__e.TailApply(PrimFunc(symshen_4fillvector), tmp1127, tmp1128, V3367, V3368)
return


}


}, 4)

tmp1131 := Call(__e, ns2_1set, symshen_4fillvector, tmp1126)


_ = tmp1131

tmp1132 := MakeNative(func(__e *ControlFlow) {
V3369 := __e.Get(1)
_ = V3369
tmp1139 := PrimIsVector(V3369)

if True == tmp1139 {
tmp1134 := MakeNative(func(__e *ControlFlow) {
tmp1135 := PrimVectorGet(V3369, MakeNumber(0))

__e.Return(PrimGreatEqual(tmp1135, MakeNumber(0)))
return


}, 0)

tmp1136 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(False)
return
}, 1)

tmp1137 := Call(__e, try_1catch, tmp1134, tmp1136)


if True == tmp1137 {
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

tmp1140 := Call(__e, ns2_1set, symvector_2, tmp1132)


_ = tmp1140

tmp1141 := MakeNative(func(__e *ControlFlow) {
V3370 := __e.Get(1)
_ = V3370
V3371 := __e.Get(2)
_ = V3371
V3372 := __e.Get(3)
_ = V3372
tmp1143 := PrimEqual(V3371, MakeNumber(0))

if True == tmp1143 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
__e.Return(PrimVectorSet(V3370, V3371, V3372))
return
}


}, 3)

tmp1144 := Call(__e, ns2_1set, symvector_1_6, tmp1141)


_ = tmp1144

tmp1145 := MakeNative(func(__e *ControlFlow) {
V3373 := __e.Get(1)
_ = V3373
V3374 := __e.Get(2)
_ = V3374
tmp1152 := PrimEqual(V3374, MakeNumber(0))

if True == tmp1152 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
tmp1146 := MakeNative(func(__e *ControlFlow) {
VectorElement := __e.Get(1)
_ = VectorElement
tmp1148 := Call(__e, PrimFunc(symfail))


tmp1149 := PrimEqual(VectorElement, tmp1148)

if True == tmp1149 {
__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
return
} else {
__e.Return(VectorElement)
return
}


}, 1)

tmp1150 := PrimVectorGet(V3373, V3374)

__e.TailApply(tmp1146, tmp1150)
return


}


}, 2)

tmp1153 := Call(__e, ns2_1set, sym_5_1vector, tmp1145)


_ = tmp1153

tmp1154 := MakeNative(func(__e *ControlFlow) {
V3375 := __e.Get(1)
_ = V3375
tmp1158 := PrimIsInteger(V3375)

if True == tmp1158 {
tmp1156 := PrimGreatEqual(V3375, MakeNumber(0))

if True == tmp1156 {
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

tmp1159 := Call(__e, ns2_1set, symshen_4posint_2, tmp1154)


_ = tmp1159

tmp1160 := MakeNative(func(__e *ControlFlow) {
V3376 := __e.Get(1)
_ = V3376
__e.Return(PrimVectorGet(V3376, MakeNumber(0)))
return
}, 1)

tmp1161 := Call(__e, ns2_1set, symlimit, tmp1160)


_ = tmp1161

tmp1162 := MakeNative(func(__e *ControlFlow) {
V3380 := __e.Get(1)
_ = V3380
tmp1171 := Call(__e, PrimFunc(symshen_4_7string_2), V3380)


if True == tmp1171 {
tmp1167 := Call(__e, PrimFunc(symhdstr), V3380)


tmp1168 := PrimStringToNumber(tmp1167)

tmp1169 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1168)


if True == tmp1169 {
tmp1164 := PrimTailString(V3380)

tmp1165 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1164)


if True == tmp1165 {
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

tmp1172 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1162)


_ = tmp1172

tmp1173 := MakeNative(func(__e *ControlFlow) {
V3383 := __e.Get(1)
_ = V3383
tmp1188 := PrimEqual(MakeString(""), V3383)

if True == tmp1188 {
__e.Return(True)
return
} else {
tmp1186 := Call(__e, PrimFunc(symshen_4_7string_2), V3383)


if True == tmp1186 {
tmp1174 := MakeNative(func(__e *ControlFlow) {
N := __e.Get(1)
_ = N
tmp1182 := Call(__e, PrimFunc(symshen_4alpha_2), N)


var ifres1179 Obj

if True == tmp1182 {
ifres1179 = True


} else {
tmp1181 := Call(__e, PrimFunc(symshen_4digit_2), N)


var ifres1180 Obj

if True == tmp1181 {
ifres1180 = True


} else {
ifres1180 = False


}

ifres1179 = ifres1180


}

if True == ifres1179 {
tmp1176 := PrimTailString(V3383)

tmp1177 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1176)


if True == tmp1177 {
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

tmp1183 := Call(__e, PrimFunc(symhdstr), V3383)


tmp1184 := PrimStringToNumber(tmp1183)

__e.TailApply(tmp1174, tmp1184)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.alphanums?")))
return
}


}


}, 1)

tmp1189 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1173)


_ = tmp1189

tmp1190 := MakeNative(func(__e *ControlFlow) {
V3384 := __e.Get(1)
_ = V3384
tmp1202 := Call(__e, PrimFunc(symboolean_2), V3384)


var ifres1196 Obj

if True == tmp1202 {
ifres1196 = True


} else {
tmp1201 := PrimIsNumber(V3384)

var ifres1198 Obj

if True == tmp1201 {
ifres1198 = True


} else {
tmp1200 := PrimIsString(V3384)

var ifres1199 Obj

if True == tmp1200 {
ifres1199 = True


} else {
ifres1199 = False


}

ifres1198 = ifres1199


}

var ifres1197 Obj

if True == ifres1198 {
ifres1197 = True


} else {
ifres1197 = False


}

ifres1196 = ifres1197


}

if True == ifres1196 {
__e.Return(False)
return
} else {
tmp1191 := MakeNative(func(__e *ControlFlow) {
tmp1192 := MakeNative(func(__e *ControlFlow) {
String := __e.Get(1)
_ = String
__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), String)
return
}, 1)

tmp1193 := PrimStr(V3384)

__e.TailApply(tmp1192, tmp1193)
return


}, 0)

tmp1194 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1191, tmp1194)
return


}


}, 1)

tmp1203 := Call(__e, ns2_1set, symvariable_2, tmp1190)


_ = tmp1203

tmp1204 := MakeNative(func(__e *ControlFlow) {
V3387 := __e.Get(1)
_ = V3387
tmp1213 := Call(__e, PrimFunc(symshen_4_7string_2), V3387)


if True == tmp1213 {
tmp1209 := Call(__e, PrimFunc(symhdstr), V3387)


tmp1210 := PrimStringToNumber(tmp1209)

tmp1211 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1210)


if True == tmp1211 {
tmp1206 := PrimTailString(V3387)

tmp1207 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1206)


if True == tmp1207 {
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

tmp1214 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1204)


_ = tmp1214

tmp1215 := MakeNative(func(__e *ControlFlow) {
V3388 := __e.Get(1)
_ = V3388
tmp1216 := PrimValue(symshen_4_dgensym_d)

tmp1217 := PrimNumberAdd(MakeNumber(1), tmp1216)

tmp1218 := PrimSet(symshen_4_dgensym_d, tmp1217)

__e.TailApply(PrimFunc(symconcat), V3388, tmp1218)
return


}, 1)

tmp1219 := Call(__e, ns2_1set, symgensym, tmp1215)


_ = tmp1219

tmp1220 := MakeNative(func(__e *ControlFlow) {
V3389 := __e.Get(1)
_ = V3389
V3390 := __e.Get(2)
_ = V3390
tmp1221 := PrimStr(V3389)

tmp1222 := PrimStr(V3390)

tmp1223 := PrimStringConcat(tmp1221, tmp1222)

__e.Return(PrimIntern(tmp1223))
return


}, 2)

tmp1224 := Call(__e, ns2_1set, symconcat, tmp1220)


_ = tmp1224

tmp1225 := MakeNative(func(__e *ControlFlow) {
V3391 := __e.Get(1)
_ = V3391
V3392 := __e.Get(2)
_ = V3392
tmp1226 := MakeNative(func(__e *ControlFlow) {
Vector := __e.Get(1)
_ = Vector
tmp1227 := MakeNative(func(__e *ControlFlow) {
Tag := __e.Get(1)
_ = Tag
tmp1228 := MakeNative(func(__e *ControlFlow) {
Fst := __e.Get(1)
_ = Fst
tmp1229 := MakeNative(func(__e *ControlFlow) {
Snd := __e.Get(1)
_ = Snd
__e.Return(Vector)
return
}, 1)

tmp1230 := PrimVectorSet(Vector, MakeNumber(2), V3392)

__e.TailApply(tmp1229, tmp1230)
return


}, 1)

tmp1231 := PrimVectorSet(Vector, MakeNumber(1), V3391)

__e.TailApply(tmp1228, tmp1231)
return


}, 1)

tmp1232 := PrimVectorSet(Vector, MakeNumber(0), symshen_4tuple)

__e.TailApply(tmp1227, tmp1232)
return


}, 1)

tmp1233 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp1226, tmp1233)
return


}, 2)

tmp1234 := Call(__e, ns2_1set, sym_8p, tmp1225)


_ = tmp1234

tmp1235 := MakeNative(func(__e *ControlFlow) {
V3393 := __e.Get(1)
_ = V3393
__e.Return(PrimVectorGet(V3393, MakeNumber(1)))
return
}, 1)

tmp1236 := Call(__e, ns2_1set, symfst, tmp1235)


_ = tmp1236

tmp1237 := MakeNative(func(__e *ControlFlow) {
V3394 := __e.Get(1)
_ = V3394
__e.Return(PrimVectorGet(V3394, MakeNumber(2)))
return
}, 1)

tmp1238 := Call(__e, ns2_1set, symsnd, tmp1237)


_ = tmp1238

tmp1239 := MakeNative(func(__e *ControlFlow) {
V3395 := __e.Get(1)
_ = V3395
tmp1240 := MakeNative(func(__e *ControlFlow) {
tmp1245 := PrimIsVector(V3395)

if True == tmp1245 {
tmp1242 := PrimVectorGet(V3395, MakeNumber(0))

tmp1243 := PrimEqual(symshen_4tuple, tmp1242)

if True == tmp1243 {
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

tmp1246 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1240, tmp1246)
return


}, 1)

tmp1247 := Call(__e, ns2_1set, symtuple_2, tmp1239)


_ = tmp1247

tmp1248 := MakeNative(func(__e *ControlFlow) {
V3400 := __e.Get(1)
_ = V3400
V3401 := __e.Get(2)
_ = V3401
tmp1255 := PrimEqual(Nil, V3400)

if True == tmp1255 {
__e.Return(V3401)
return
} else {
tmp1253 := PrimIsPair(V3400)

if True == tmp1253 {
tmp1249 := PrimHead(V3400)

tmp1250 := PrimTail(V3400)

tmp1251 := Call(__e, PrimFunc(symappend), tmp1250, V3401)


__e.Return(PrimCons(tmp1249, tmp1251))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to append a non-list")))
return
}


}


}, 2)

tmp1256 := Call(__e, ns2_1set, symappend, tmp1248)


_ = tmp1256

tmp1257 := MakeNative(func(__e *ControlFlow) {
V3402 := __e.Get(1)
_ = V3402
V3403 := __e.Get(2)
_ = V3403
tmp1258 := MakeNative(func(__e *ControlFlow) {
Limit := __e.Get(1)
_ = Limit
tmp1259 := MakeNative(func(__e *ControlFlow) {
NewVector := __e.Get(1)
_ = NewVector
tmp1260 := MakeNative(func(__e *ControlFlow) {
X_7NewVector := __e.Get(1)
_ = X_7NewVector
tmp1262 := PrimEqual(Limit, MakeNumber(0))

if True == tmp1262 {
__e.Return(X_7NewVector)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3403, MakeNumber(1), Limit, X_7NewVector)
return
}


}, 1)

tmp1263 := Call(__e, PrimFunc(symvector_1_6), NewVector, MakeNumber(1), V3402)


__e.TailApply(tmp1260, tmp1263)
return


}, 1)

tmp1264 := PrimNumberAdd(Limit, MakeNumber(1))

tmp1265 := Call(__e, PrimFunc(symvector), tmp1264)


__e.TailApply(tmp1259, tmp1265)
return


}, 1)

tmp1266 := Call(__e, PrimFunc(symlimit), V3403)


__e.TailApply(tmp1258, tmp1266)
return


}, 2)

tmp1267 := Call(__e, ns2_1set, sym_8v, tmp1257)


_ = tmp1267

tmp1268 := MakeNative(func(__e *ControlFlow) {
V3405 := __e.Get(1)
_ = V3405
V3406 := __e.Get(2)
_ = V3406
V3407 := __e.Get(3)
_ = V3407
V3408 := __e.Get(4)
_ = V3408
tmp1274 := PrimEqual(V3406, V3407)

if True == tmp1274 {
tmp1269 := PrimNumberAdd(V3407, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3405, V3408, V3407, tmp1269)
return


} else {
tmp1270 := PrimNumberAdd(V3406, MakeNumber(1))

tmp1271 := PrimNumberAdd(V3406, MakeNumber(1))

tmp1272 := Call(__e, PrimFunc(symshen_4copyfromvector), V3405, V3408, V3406, tmp1271)


__e.TailApply(PrimFunc(symshen_4_8v_1help), V3405, tmp1270, V3407, tmp1272)
return


}


}, 4)

tmp1275 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1268)


_ = tmp1275

tmp1276 := MakeNative(func(__e *ControlFlow) {
V3409 := __e.Get(1)
_ = V3409
V3410 := __e.Get(2)
_ = V3410
V3411 := __e.Get(3)
_ = V3411
V3412 := __e.Get(4)
_ = V3412
tmp1277 := MakeNative(func(__e *ControlFlow) {
tmp1278 := Call(__e, PrimFunc(sym_5_1vector), V3409, V3411)


__e.TailApply(PrimFunc(symvector_1_6), V3410, V3412, tmp1278)
return


}, 0)

tmp1279 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(V3410)
return
}, 1)

__e.TailApply(try_1catch, tmp1277, tmp1279)
return


}, 4)

tmp1280 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1276)


_ = tmp1280

tmp1281 := MakeNative(func(__e *ControlFlow) {
V3413 := __e.Get(1)
_ = V3413
tmp1282 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3413, MakeNumber(1))
return
}, 0)

tmp1283 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n")))
return
}, 1)

__e.TailApply(try_1catch, tmp1282, tmp1283)
return


}, 1)

tmp1284 := Call(__e, ns2_1set, symhdv, tmp1281)


_ = tmp1284

tmp1285 := MakeNative(func(__e *ControlFlow) {
V3414 := __e.Get(1)
_ = V3414
tmp1286 := MakeNative(func(__e *ControlFlow) {
Limit := __e.Get(1)
_ = Limit
tmp1295 := PrimEqual(Limit, MakeNumber(0))

if True == tmp1295 {
__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
return
} else {
tmp1293 := PrimEqual(Limit, MakeNumber(1))

if True == tmp1293 {
__e.TailApply(PrimFunc(symvector), MakeNumber(0))
return
} else {
tmp1287 := MakeNative(func(__e *ControlFlow) {
NewVector := __e.Get(1)
_ = NewVector
tmp1288 := PrimNumberSubtract(Limit, MakeNumber(1))

tmp1289 := Call(__e, PrimFunc(symvector), tmp1288)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3414, MakeNumber(2), Limit, tmp1289)
return


}, 1)

tmp1290 := PrimNumberSubtract(Limit, MakeNumber(1))

tmp1291 := Call(__e, PrimFunc(symvector), tmp1290)


__e.TailApply(tmp1287, tmp1291)
return


}


}


}, 1)

tmp1296 := Call(__e, PrimFunc(symlimit), V3414)


__e.TailApply(tmp1286, tmp1296)
return


}, 1)

tmp1297 := Call(__e, ns2_1set, symtlv, tmp1285)


_ = tmp1297

tmp1298 := MakeNative(func(__e *ControlFlow) {
V3416 := __e.Get(1)
_ = V3416
V3417 := __e.Get(2)
_ = V3417
V3418 := __e.Get(3)
_ = V3418
V3419 := __e.Get(4)
_ = V3419
tmp1304 := PrimEqual(V3417, V3418)

if True == tmp1304 {
tmp1299 := PrimNumberSubtract(V3418, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3416, V3419, V3418, tmp1299)
return


} else {
tmp1300 := PrimNumberAdd(V3417, MakeNumber(1))

tmp1301 := PrimNumberSubtract(V3417, MakeNumber(1))

tmp1302 := Call(__e, PrimFunc(symshen_4copyfromvector), V3416, V3419, V3417, tmp1301)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3416, tmp1300, V3418, tmp1302)
return


}


}, 4)

tmp1305 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1298)


_ = tmp1305

tmp1306 := MakeNative(func(__e *ControlFlow) {
V3431 := __e.Get(1)
_ = V3431
V3432 := __e.Get(2)
_ = V3432
tmp1322 := PrimEqual(Nil, V3432)

if True == tmp1322 {
__e.Return(Nil)
return
} else {
tmp1320 := PrimIsPair(V3432)

var ifres1311 Obj

if True == tmp1320 {
tmp1318 := PrimHead(V3432)

tmp1319 := PrimIsPair(tmp1318)

var ifres1313 Obj

if True == tmp1319 {
tmp1315 := PrimHead(V3432)

tmp1316 := PrimHead(tmp1315)

tmp1317 := PrimEqual(V3431, tmp1316)

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

if True == ifres1311 {
__e.Return(PrimHead(V3432))
return
} else {
tmp1309 := PrimIsPair(V3432)

if True == tmp1309 {
tmp1307 := PrimTail(V3432)

__e.TailApply(PrimFunc(symassoc), V3431, tmp1307)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to search a non-list with assoc\n")))
return
}


}


}


}, 2)

tmp1323 := Call(__e, ns2_1set, symassoc, tmp1306)


_ = tmp1323

tmp1324 := MakeNative(func(__e *ControlFlow) {
V3435 := __e.Get(1)
_ = V3435
tmp1328 := PrimEqual(True, V3435)

if True == tmp1328 {
__e.Return(True)
return
} else {
tmp1326 := PrimEqual(False, V3435)

if True == tmp1326 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1329 := Call(__e, ns2_1set, symboolean_2, tmp1324)


_ = tmp1329

tmp1330 := MakeNative(func(__e *ControlFlow) {
V3436 := __e.Get(1)
_ = V3436
tmp1335 := PrimEqual(MakeNumber(0), V3436)

if True == tmp1335 {
__e.Return(MakeNumber(0))
return
} else {
tmp1331 := Call(__e, PrimFunc(symstoutput))


tmp1332 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1331)


_ = tmp1332

tmp1333 := PrimNumberSubtract(V3436, MakeNumber(1))

__e.TailApply(PrimFunc(symnl), tmp1333)
return


}


}, 1)

tmp1336 := Call(__e, ns2_1set, symnl, tmp1330)


_ = tmp1336

tmp1337 := MakeNative(func(__e *ControlFlow) {
V3443 := __e.Get(1)
_ = V3443
V3444 := __e.Get(2)
_ = V3444
tmp1348 := PrimEqual(Nil, V3443)

if True == tmp1348 {
__e.Return(Nil)
return
} else {
tmp1346 := PrimIsPair(V3443)

if True == tmp1346 {
tmp1343 := PrimHead(V3443)

tmp1344 := Call(__e, PrimFunc(symelement_2), tmp1343, V3444)


if True == tmp1344 {
tmp1338 := PrimTail(V3443)

__e.TailApply(PrimFunc(symdifference), tmp1338, V3444)
return


} else {
tmp1339 := PrimHead(V3443)

tmp1340 := PrimTail(V3443)

tmp1341 := Call(__e, PrimFunc(symdifference), tmp1340, V3444)


__e.Return(PrimCons(tmp1339, tmp1341))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the difference with a non-list\n")))
return
}


}


}, 2)

tmp1349 := Call(__e, ns2_1set, symdifference, tmp1337)


_ = tmp1349

tmp1350 := MakeNative(func(__e *ControlFlow) {
V3445 := __e.Get(1)
_ = V3445
V3446 := __e.Get(2)
_ = V3446
__e.Return(V3446)
return
}, 2)

tmp1351 := Call(__e, ns2_1set, symdo, tmp1350)


_ = tmp1351

tmp1352 := MakeNative(func(__e *ControlFlow) {
V3458 := __e.Get(1)
_ = V3458
V3459 := __e.Get(2)
_ = V3459
tmp1363 := PrimEqual(Nil, V3459)

if True == tmp1363 {
__e.Return(False)
return
} else {
tmp1361 := PrimIsPair(V3459)

var ifres1357 Obj

if True == tmp1361 {
tmp1359 := PrimHead(V3459)

tmp1360 := PrimEqual(V3458, tmp1359)

var ifres1358 Obj

if True == tmp1360 {
ifres1358 = True


} else {
ifres1358 = False


}

ifres1357 = ifres1358


} else {
ifres1357 = False


}

if True == ifres1357 {
__e.Return(True)
return
} else {
tmp1355 := PrimIsPair(V3459)

if True == tmp1355 {
tmp1353 := PrimTail(V3459)

__e.TailApply(PrimFunc(symelement_2), V3458, tmp1353)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find an element in a non-list\n")))
return
}


}


}


}, 2)

tmp1364 := Call(__e, ns2_1set, symelement_2, tmp1352)


_ = tmp1364

tmp1365 := MakeNative(func(__e *ControlFlow) {
V3462 := __e.Get(1)
_ = V3462
tmp1367 := PrimEqual(Nil, V3462)

if True == tmp1367 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1368 := Call(__e, ns2_1set, symempty_2, tmp1365)


_ = tmp1368

tmp1369 := MakeNative(func(__e *ControlFlow) {
V3463 := __e.Get(1)
_ = V3463
V3464 := __e.Get(2)
_ = V3464
tmp1370 := Call(__e, V3463, V3464)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3463, V3464, tmp1370)
return


}, 2)

tmp1371 := Call(__e, ns2_1set, symfix, tmp1369)


_ = tmp1371

tmp1372 := MakeNative(func(__e *ControlFlow) {
V3470 := __e.Get(1)
_ = V3470
V3471 := __e.Get(2)
_ = V3471
V3472 := __e.Get(3)
_ = V3472
tmp1375 := PrimEqual(V3471, V3472)

if True == tmp1375 {
__e.Return(V3472)
return
} else {
tmp1373 := Call(__e, V3470, V3472)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3470, V3472, tmp1373)
return


}


}, 3)

tmp1376 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1372)


_ = tmp1376

tmp1377 := MakeNative(func(__e *ControlFlow) {
V3473 := __e.Get(1)
_ = V3473
V3474 := __e.Get(2)
_ = V3474
V3475 := __e.Get(3)
_ = V3475
V3476 := __e.Get(4)
_ = V3476
tmp1378 := MakeNative(func(__e *ControlFlow) {
N := __e.Get(1)
_ = N
tmp1379 := MakeNative(func(__e *ControlFlow) {
Entry := __e.Get(1)
_ = Entry
tmp1380 := MakeNative(func(__e *ControlFlow) {
Change := __e.Get(1)
_ = Change
__e.Return(V3475)
return
}, 1)

tmp1381 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3473, V3474, V3475, Entry)


tmp1382 := Call(__e, PrimFunc(symvector_1_6), V3476, N, tmp1381)


__e.TailApply(tmp1380, tmp1382)
return


}, 1)

tmp1383 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3476, N)
return
}, 0)

tmp1384 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(Nil)
return
}, 1)

tmp1385 := Call(__e, try_1catch, tmp1383, tmp1384)


__e.TailApply(tmp1379, tmp1385)
return


}, 1)

tmp1386 := Call(__e, PrimFunc(symlimit), V3476)


tmp1387 := Call(__e, PrimFunc(symhash), V3473, tmp1386)


__e.TailApply(tmp1378, tmp1387)
return


}, 4)

tmp1388 := Call(__e, ns2_1set, symput, tmp1377)


_ = tmp1388

tmp1389 := MakeNative(func(__e *ControlFlow) {
V3477 := __e.Get(1)
_ = V3477
V3478 := __e.Get(2)
_ = V3478
V3479 := __e.Get(3)
_ = V3479
tmp1390 := MakeNative(func(__e *ControlFlow) {
N := __e.Get(1)
_ = N
tmp1391 := MakeNative(func(__e *ControlFlow) {
Entry := __e.Get(1)
_ = Entry
tmp1392 := MakeNative(func(__e *ControlFlow) {
Change := __e.Get(1)
_ = Change
__e.Return(V3477)
return
}, 1)

tmp1393 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3477, V3478, Entry)


tmp1394 := Call(__e, PrimFunc(symvector_1_6), V3479, N, tmp1393)


__e.TailApply(tmp1392, tmp1394)
return


}, 1)

tmp1395 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3479, N)
return
}, 0)

tmp1396 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(Nil)
return
}, 1)

tmp1397 := Call(__e, try_1catch, tmp1395, tmp1396)


__e.TailApply(tmp1391, tmp1397)
return


}, 1)

tmp1398 := Call(__e, PrimFunc(symlimit), V3479)


tmp1399 := Call(__e, PrimFunc(symhash), V3477, tmp1398)


__e.TailApply(tmp1390, tmp1399)
return


}, 3)

tmp1400 := Call(__e, ns2_1set, symunput, tmp1389)


_ = tmp1400

tmp1401 := MakeNative(func(__e *ControlFlow) {
V3490 := __e.Get(1)
_ = V3490
V3491 := __e.Get(2)
_ = V3491
V3492 := __e.Get(3)
_ = V3492
tmp1445 := PrimEqual(Nil, V3492)

if True == tmp1445 {
__e.Return(Nil)
return
} else {
tmp1443 := PrimIsPair(V3492)

var ifres1408 Obj

if True == tmp1443 {
tmp1441 := PrimHead(V3492)

tmp1442 := PrimIsPair(tmp1441)

var ifres1410 Obj

if True == tmp1442 {
tmp1438 := PrimHead(V3492)

tmp1439 := PrimHead(tmp1438)

tmp1440 := PrimIsPair(tmp1439)

var ifres1412 Obj

if True == tmp1440 {
tmp1434 := PrimHead(V3492)

tmp1435 := PrimHead(tmp1434)

tmp1436 := PrimTail(tmp1435)

tmp1437 := PrimIsPair(tmp1436)

var ifres1414 Obj

if True == tmp1437 {
tmp1429 := PrimHead(V3492)

tmp1430 := PrimHead(tmp1429)

tmp1431 := PrimTail(tmp1430)

tmp1432 := PrimTail(tmp1431)

tmp1433 := PrimEqual(Nil, tmp1432)

var ifres1416 Obj

if True == tmp1433 {
tmp1424 := PrimHead(V3492)

tmp1425 := PrimHead(tmp1424)

tmp1426 := PrimTail(tmp1425)

tmp1427 := PrimHead(tmp1426)

tmp1428 := PrimEqual(V3491, tmp1427)

var ifres1418 Obj

if True == tmp1428 {
tmp1420 := PrimHead(V3492)

tmp1421 := PrimHead(tmp1420)

tmp1422 := PrimHead(tmp1421)

tmp1423 := PrimEqual(V3490, tmp1422)

var ifres1419 Obj

if True == tmp1423 {
ifres1419 = True


} else {
ifres1419 = False


}

ifres1418 = ifres1419


} else {
ifres1418 = False


}

var ifres1417 Obj

if True == ifres1418 {
ifres1417 = True


} else {
ifres1417 = False


}

ifres1416 = ifres1417


} else {
ifres1416 = False


}

var ifres1415 Obj

if True == ifres1416 {
ifres1415 = True


} else {
ifres1415 = False


}

ifres1414 = ifres1415


} else {
ifres1414 = False


}

var ifres1413 Obj

if True == ifres1414 {
ifres1413 = True


} else {
ifres1413 = False


}

ifres1412 = ifres1413


} else {
ifres1412 = False


}

var ifres1411 Obj

if True == ifres1412 {
ifres1411 = True


} else {
ifres1411 = False


}

ifres1410 = ifres1411


} else {
ifres1410 = False


}

var ifres1409 Obj

if True == ifres1410 {
ifres1409 = True


} else {
ifres1409 = False


}

ifres1408 = ifres1409


} else {
ifres1408 = False


}

if True == ifres1408 {
__e.Return(PrimTail(V3492))
return
} else {
tmp1406 := PrimIsPair(V3492)

if True == tmp1406 {
tmp1402 := PrimHead(V3492)

tmp1403 := PrimTail(V3492)

tmp1404 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3490, V3491, tmp1403)


__e.Return(PrimCons(tmp1402, tmp1404))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-pointer")))
return
}


}


}


}, 3)

tmp1446 := Call(__e, ns2_1set, symshen_4remove_1pointer, tmp1401)


_ = tmp1446

tmp1447 := MakeNative(func(__e *ControlFlow) {
V3505 := __e.Get(1)
_ = V3505
V3506 := __e.Get(2)
_ = V3506
V3507 := __e.Get(3)
_ = V3507
V3508 := __e.Get(4)
_ = V3508
tmp1498 := PrimEqual(Nil, V3508)

if True == tmp1498 {
tmp1448 := PrimCons(V3506, Nil)

tmp1449 := PrimCons(V3505, tmp1448)

tmp1450 := PrimCons(tmp1449, V3507)

__e.Return(PrimCons(tmp1450, Nil))
return


} else {
tmp1496 := PrimIsPair(V3508)

var ifres1461 Obj

if True == tmp1496 {
tmp1494 := PrimHead(V3508)

tmp1495 := PrimIsPair(tmp1494)

var ifres1463 Obj

if True == tmp1495 {
tmp1491 := PrimHead(V3508)

tmp1492 := PrimHead(tmp1491)

tmp1493 := PrimIsPair(tmp1492)

var ifres1465 Obj

if True == tmp1493 {
tmp1487 := PrimHead(V3508)

tmp1488 := PrimHead(tmp1487)

tmp1489 := PrimTail(tmp1488)

tmp1490 := PrimIsPair(tmp1489)

var ifres1467 Obj

if True == tmp1490 {
tmp1482 := PrimHead(V3508)

tmp1483 := PrimHead(tmp1482)

tmp1484 := PrimTail(tmp1483)

tmp1485 := PrimTail(tmp1484)

tmp1486 := PrimEqual(Nil, tmp1485)

var ifres1469 Obj

if True == tmp1486 {
tmp1477 := PrimHead(V3508)

tmp1478 := PrimHead(tmp1477)

tmp1479 := PrimTail(tmp1478)

tmp1480 := PrimHead(tmp1479)

tmp1481 := PrimEqual(V3506, tmp1480)

var ifres1471 Obj

if True == tmp1481 {
tmp1473 := PrimHead(V3508)

tmp1474 := PrimHead(tmp1473)

tmp1475 := PrimHead(tmp1474)

tmp1476 := PrimEqual(V3505, tmp1475)

var ifres1472 Obj

if True == tmp1476 {
ifres1472 = True


} else {
ifres1472 = False


}

ifres1471 = ifres1472


} else {
ifres1471 = False


}

var ifres1470 Obj

if True == ifres1471 {
ifres1470 = True


} else {
ifres1470 = False


}

ifres1469 = ifres1470


} else {
ifres1469 = False


}

var ifres1468 Obj

if True == ifres1469 {
ifres1468 = True


} else {
ifres1468 = False


}

ifres1467 = ifres1468


} else {
ifres1467 = False


}

var ifres1466 Obj

if True == ifres1467 {
ifres1466 = True


} else {
ifres1466 = False


}

ifres1465 = ifres1466


} else {
ifres1465 = False


}

var ifres1464 Obj

if True == ifres1465 {
ifres1464 = True


} else {
ifres1464 = False


}

ifres1463 = ifres1464


} else {
ifres1463 = False


}

var ifres1462 Obj

if True == ifres1463 {
ifres1462 = True


} else {
ifres1462 = False


}

ifres1461 = ifres1462


} else {
ifres1461 = False


}

if True == ifres1461 {
tmp1451 := PrimHead(V3508)

tmp1452 := PrimHead(tmp1451)

tmp1453 := PrimCons(tmp1452, V3507)

tmp1454 := PrimTail(V3508)

__e.Return(PrimCons(tmp1453, tmp1454))
return


} else {
tmp1459 := PrimIsPair(V3508)

if True == tmp1459 {
tmp1455 := PrimHead(V3508)

tmp1456 := PrimTail(V3508)

tmp1457 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3505, V3506, V3507, tmp1456)


__e.Return(PrimCons(tmp1455, tmp1457))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.change-pointer-value")))
return
}


}


}


}, 4)

tmp1499 := Call(__e, ns2_1set, symshen_4change_1pointer_1value, tmp1447)


_ = tmp1499

tmp1500 := MakeNative(func(__e *ControlFlow) {
V3509 := __e.Get(1)
_ = V3509
V3510 := __e.Get(2)
_ = V3510
V3511 := __e.Get(3)
_ = V3511
tmp1501 := MakeNative(func(__e *ControlFlow) {
N := __e.Get(1)
_ = N
tmp1502 := MakeNative(func(__e *ControlFlow) {
Entry := __e.Get(1)
_ = Entry
tmp1503 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp1509 := Call(__e, PrimFunc(symempty_2), Result)


if True == tmp1509 {
tmp1504 := Call(__e, PrimFunc(symshen_4app), V3509, MakeString("\n"), symshen_4s)


tmp1505 := PrimStringConcat(MakeString(" not found for "), tmp1504)

tmp1506 := Call(__e, PrimFunc(symshen_4app), V3510, tmp1505, symshen_4s)


tmp1507 := PrimStringConcat(MakeString("attribute "), tmp1506)

__e.Return(PrimSimpleError(tmp1507))
return


} else {
__e.Return(PrimTail(Result))
return
}


}, 1)

tmp1510 := PrimCons(V3510, Nil)

tmp1511 := PrimCons(V3509, tmp1510)

tmp1512 := Call(__e, PrimFunc(symassoc), tmp1511, Entry)


__e.TailApply(tmp1503, tmp1512)
return


}, 1)

tmp1513 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3511, N)
return
}, 0)

tmp1514 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
tmp1515 := Call(__e, PrimFunc(symshen_4app), V3510, MakeString("\n"), symshen_4s)


tmp1516 := PrimStringConcat(MakeString(" has no attributes: "), tmp1515)

tmp1517 := Call(__e, PrimFunc(symshen_4app), V3509, tmp1516, symshen_4a)


__e.Return(PrimSimpleError(tmp1517))
return


}, 1)

tmp1518 := Call(__e, try_1catch, tmp1513, tmp1514)


__e.TailApply(tmp1502, tmp1518)
return


}, 1)

tmp1519 := Call(__e, PrimFunc(symlimit), V3511)


tmp1520 := Call(__e, PrimFunc(symhash), V3509, tmp1519)


__e.TailApply(tmp1501, tmp1520)
return


}, 3)

tmp1521 := Call(__e, ns2_1set, symget, tmp1500)


_ = tmp1521

tmp1522 := MakeNative(func(__e *ControlFlow) {
V3512 := __e.Get(1)
_ = V3512
V3513 := __e.Get(2)
_ = V3513
tmp1523 := MakeNative(func(__e *ControlFlow) {
Hash := __e.Get(1)
_ = Hash
tmp1525 := PrimEqual(Hash, MakeNumber(0))

if True == tmp1525 {
__e.Return(MakeNumber(1))
return
} else {
__e.Return(Hash)
return
}


}, 1)

tmp1526 := Call(__e, PrimFunc(symshen_4hashkey), V3512)


tmp1527 := Call(__e, PrimFunc(symshen_4mod), tmp1526, V3513)


__e.TailApply(tmp1523, tmp1527)
return


}, 2)

tmp1528 := Call(__e, ns2_1set, symhash, tmp1522)


_ = tmp1528

tmp1529 := MakeNative(func(__e *ControlFlow) {
V3514 := __e.Get(1)
_ = V3514
tmp1530 := MakeNative(func(__e *ControlFlow) {
Ns := __e.Get(1)
_ = Ns
__e.TailApply(PrimFunc(symshen_4prodbutzero), Ns, MakeNumber(1))
return
}, 1)

tmp1531 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(PrimStringToNumber(X))
return
}, 1)

tmp1532 := Call(__e, PrimFunc(symexplode), V3514)


tmp1533 := Call(__e, PrimFunc(symmap), tmp1531, tmp1532)


__e.TailApply(tmp1530, tmp1533)
return


}, 1)

tmp1534 := Call(__e, ns2_1set, symshen_4hashkey, tmp1529)


_ = tmp1534

tmp1535 := MakeNative(func(__e *ControlFlow) {
V3515 := __e.Get(1)
_ = V3515
V3516 := __e.Get(2)
_ = V3516
tmp1554 := PrimEqual(Nil, V3515)

if True == tmp1554 {
__e.Return(V3516)
return
} else {
tmp1552 := PrimIsPair(V3515)

var ifres1548 Obj

if True == tmp1552 {
tmp1550 := PrimHead(V3515)

tmp1551 := PrimEqual(MakeNumber(0), tmp1550)

var ifres1549 Obj

if True == tmp1551 {
ifres1549 = True


} else {
ifres1549 = False


}

ifres1548 = ifres1549


} else {
ifres1548 = False


}

if True == ifres1548 {
tmp1536 := PrimTail(V3515)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1536, V3516)
return


} else {
tmp1546 := PrimIsPair(V3515)

if True == tmp1546 {
tmp1544 := PrimGreatThan(V3516, MakeNumber(10000000000))

if True == tmp1544 {
tmp1537 := PrimTail(V3515)

tmp1538 := PrimHead(V3515)

tmp1539 := PrimNumberAdd(V3516, tmp1538)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1537, tmp1539)
return


} else {
tmp1540 := PrimTail(V3515)

tmp1541 := PrimHead(V3515)

tmp1542 := PrimNumberMultiply(V3516, tmp1541)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1540, tmp1542)
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prodbutzero)
return
}


}


}


}, 2)

tmp1555 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1535)


_ = tmp1555

tmp1556 := MakeNative(func(__e *ControlFlow) {
V3517 := __e.Get(1)
_ = V3517
V3518 := __e.Get(2)
_ = V3518
tmp1557 := PrimCons(V3518, Nil)

tmp1558 := Call(__e, PrimFunc(symshen_4multiples), V3517, tmp1557)


__e.TailApply(PrimFunc(symshen_4modh), V3517, tmp1558)
return


}, 2)

tmp1559 := Call(__e, ns2_1set, symshen_4mod, tmp1556)


_ = tmp1559

tmp1560 := MakeNative(func(__e *ControlFlow) {
V3523 := __e.Get(1)
_ = V3523
V3524 := __e.Get(2)
_ = V3524
tmp1571 := PrimIsPair(V3524)

var ifres1567 Obj

if True == tmp1571 {
tmp1569 := PrimHead(V3524)

tmp1570 := PrimGreatThan(tmp1569, V3523)

var ifres1568 Obj

if True == tmp1570 {
ifres1568 = True


} else {
ifres1568 = False


}

ifres1567 = ifres1568


} else {
ifres1567 = False


}

if True == ifres1567 {
__e.Return(PrimTail(V3524))
return
} else {
tmp1565 := PrimIsPair(V3524)

if True == tmp1565 {
tmp1561 := PrimHead(V3524)

tmp1562 := PrimNumberMultiply(MakeNumber(2), tmp1561)

tmp1563 := PrimCons(tmp1562, V3524)

__e.TailApply(PrimFunc(symshen_4multiples), V3523, tmp1563)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.multiples")))
return
}


}


}, 2)

tmp1572 := Call(__e, ns2_1set, symshen_4multiples, tmp1560)


_ = tmp1572

tmp1573 := MakeNative(func(__e *ControlFlow) {
V3531 := __e.Get(1)
_ = V3531
V3532 := __e.Get(2)
_ = V3532
tmp1591 := PrimEqual(MakeNumber(0), V3531)

if True == tmp1591 {
__e.Return(MakeNumber(0))
return
} else {
tmp1589 := PrimEqual(Nil, V3532)

if True == tmp1589 {
__e.Return(V3531)
return
} else {
tmp1587 := PrimIsPair(V3532)

var ifres1583 Obj

if True == tmp1587 {
tmp1585 := PrimHead(V3532)

tmp1586 := PrimGreatThan(tmp1585, V3531)

var ifres1584 Obj

if True == tmp1586 {
ifres1584 = True


} else {
ifres1584 = False


}

ifres1583 = ifres1584


} else {
ifres1583 = False


}

if True == ifres1583 {
tmp1576 := PrimTail(V3532)

tmp1577 := Call(__e, PrimFunc(symempty_2), tmp1576)


if True == tmp1577 {
__e.Return(V3531)
return
} else {
tmp1574 := PrimTail(V3532)

__e.TailApply(PrimFunc(symshen_4modh), V3531, tmp1574)
return


}


} else {
tmp1581 := PrimIsPair(V3532)

if True == tmp1581 {
tmp1578 := PrimHead(V3532)

tmp1579 := PrimNumberSubtract(V3531, tmp1578)

__e.TailApply(PrimFunc(symshen_4modh), tmp1579, V3532)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.modh")))
return
}


}


}


}


}, 2)

tmp1592 := Call(__e, ns2_1set, symshen_4modh, tmp1573)


_ = tmp1592

tmp1593 := MakeNative(func(__e *ControlFlow) {
V3535 := __e.Get(1)
_ = V3535
tmp1600 := PrimEqual(Nil, V3535)

if True == tmp1600 {
__e.Return(MakeNumber(0))
return
} else {
tmp1598 := PrimIsPair(V3535)

if True == tmp1598 {
tmp1594 := PrimHead(V3535)

tmp1595 := PrimTail(V3535)

tmp1596 := Call(__e, PrimFunc(symsum), tmp1595)


__e.Return(PrimNumberAdd(tmp1594, tmp1596))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to sum a non-list\n")))
return
}


}


}, 1)

tmp1601 := Call(__e, ns2_1set, symsum, tmp1593)


_ = tmp1601

tmp1602 := MakeNative(func(__e *ControlFlow) {
V3540 := __e.Get(1)
_ = V3540
tmp1604 := PrimIsPair(V3540)

if True == tmp1604 {
__e.Return(PrimHead(V3540))
return
} else {
__e.Return(PrimSimpleError(MakeString("head expects a non-empty list\n")))
return
}


}, 1)

tmp1605 := Call(__e, ns2_1set, symhead, tmp1602)


_ = tmp1605

tmp1606 := MakeNative(func(__e *ControlFlow) {
V3545 := __e.Get(1)
_ = V3545
tmp1608 := PrimIsPair(V3545)

if True == tmp1608 {
__e.Return(PrimTail(V3545))
return
} else {
__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list\n")))
return
}


}, 1)

tmp1609 := Call(__e, ns2_1set, symtail, tmp1606)


_ = tmp1609

tmp1610 := MakeNative(func(__e *ControlFlow) {
V3546 := __e.Get(1)
_ = V3546
__e.Return(PrimPos(V3546, MakeNumber(0)))
return
}, 1)

tmp1611 := Call(__e, ns2_1set, symhdstr, tmp1610)


_ = tmp1611

tmp1612 := MakeNative(func(__e *ControlFlow) {
V3553 := __e.Get(1)
_ = V3553
V3554 := __e.Get(2)
_ = V3554
tmp1623 := PrimEqual(Nil, V3553)

if True == tmp1623 {
__e.Return(Nil)
return
} else {
tmp1621 := PrimIsPair(V3553)

if True == tmp1621 {
tmp1618 := PrimHead(V3553)

tmp1619 := Call(__e, PrimFunc(symelement_2), tmp1618, V3554)


if True == tmp1619 {
tmp1613 := PrimHead(V3553)

tmp1614 := PrimTail(V3553)

tmp1615 := Call(__e, PrimFunc(symintersection), tmp1614, V3554)


__e.Return(PrimCons(tmp1613, tmp1615))
return


} else {
tmp1616 := PrimTail(V3553)

__e.TailApply(PrimFunc(symintersection), tmp1616, V3554)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n")))
return
}


}


}, 2)

tmp1624 := Call(__e, ns2_1set, symintersection, tmp1612)


_ = tmp1624

tmp1625 := MakeNative(func(__e *ControlFlow) {
V3555 := __e.Get(1)
_ = V3555
__e.TailApply(PrimFunc(symshen_4reverse_1help), V3555, Nil)
return
}, 1)

tmp1626 := Call(__e, ns2_1set, symreverse, tmp1625)


_ = tmp1626

tmp1627 := MakeNative(func(__e *ControlFlow) {
V3560 := __e.Get(1)
_ = V3560
V3561 := __e.Get(2)
_ = V3561
tmp1634 := PrimEqual(Nil, V3560)

if True == tmp1634 {
__e.Return(V3561)
return
} else {
tmp1632 := PrimIsPair(V3560)

if True == tmp1632 {
tmp1628 := PrimTail(V3560)

tmp1629 := PrimHead(V3560)

tmp1630 := PrimCons(tmp1629, V3561)

__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1628, tmp1630)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to reverse a non-list\n")))
return
}


}


}, 2)

tmp1635 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1627)


_ = tmp1635

tmp1636 := MakeNative(func(__e *ControlFlow) {
V3566 := __e.Get(1)
_ = V3566
V3567 := __e.Get(2)
_ = V3567
tmp1647 := PrimEqual(Nil, V3566)

if True == tmp1647 {
__e.Return(V3567)
return
} else {
tmp1645 := PrimIsPair(V3566)

if True == tmp1645 {
tmp1642 := PrimHead(V3566)

tmp1643 := Call(__e, PrimFunc(symelement_2), tmp1642, V3567)


if True == tmp1643 {
tmp1637 := PrimTail(V3566)

__e.TailApply(PrimFunc(symunion), tmp1637, V3567)
return


} else {
tmp1638 := PrimHead(V3566)

tmp1639 := PrimTail(V3566)

tmp1640 := Call(__e, PrimFunc(symunion), tmp1639, V3567)


__e.Return(PrimCons(tmp1638, tmp1640))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the union with a non-list\n")))
return
}


}


}, 2)

tmp1648 := Call(__e, ns2_1set, symunion, tmp1636)


_ = tmp1648

tmp1649 := MakeNative(func(__e *ControlFlow) {
V3568 := __e.Get(1)
_ = V3568
tmp1650 := MakeNative(func(__e *ControlFlow) {
Message := __e.Get(1)
_ = Message
tmp1651 := MakeNative(func(__e *ControlFlow) {
Y_1or_1N := __e.Get(1)
_ = Y_1or_1N
tmp1652 := MakeNative(func(__e *ControlFlow) {
Input := __e.Get(1)
_ = Input
tmp1658 := PrimEqual(MakeString("y"), Input)

if True == tmp1658 {
__e.Return(True)
return
} else {
tmp1656 := PrimEqual(MakeString("n"), Input)

if True == tmp1656 {
__e.Return(False)
return
} else {
tmp1653 := Call(__e, PrimFunc(symstoutput))


tmp1654 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1653)


_ = tmp1654

__e.TailApply(PrimFunc(symy_1or_1n_2), V3568)
return


}


}


}, 1)

tmp1659 := Call(__e, PrimFunc(symstinput))


tmp1660 := Call(__e, PrimFunc(symread), tmp1659)


tmp1661 := Call(__e, PrimFunc(symshen_4app), tmp1660, MakeString(""), symshen_4s)


__e.TailApply(tmp1652, tmp1661)
return


}, 1)

tmp1662 := Call(__e, PrimFunc(symstoutput))


tmp1663 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1662)


__e.TailApply(tmp1651, tmp1663)
return


}, 1)

tmp1664 := Call(__e, PrimFunc(symshen_4proc_1nl), V3568)


tmp1665 := Call(__e, PrimFunc(symstoutput))


tmp1666 := Call(__e, PrimFunc(sympr), tmp1664, tmp1665)


__e.TailApply(tmp1650, tmp1666)
return


}, 1)

tmp1667 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1649)


_ = tmp1667

tmp1668 := MakeNative(func(__e *ControlFlow) {
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

tmp1670 := Call(__e, ns2_1set, symnot, tmp1668)


_ = tmp1670

tmp1671 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("")))
return
}, 0)

tmp1672 := Call(__e, ns2_1set, symabort, tmp1671)


_ = tmp1672

tmp1673 := MakeNative(func(__e *ControlFlow) {
V3575 := __e.Get(1)
_ = V3575
V3576 := __e.Get(2)
_ = V3576
V3577 := __e.Get(3)
_ = V3577
tmp1681 := PrimEqual(V3576, V3577)

if True == tmp1681 {
__e.Return(V3575)
return
} else {
tmp1679 := PrimIsPair(V3577)

if True == tmp1679 {
tmp1674 := PrimHead(V3577)

tmp1675 := Call(__e, PrimFunc(symsubst), V3575, V3576, tmp1674)


tmp1676 := PrimTail(V3577)

tmp1677 := Call(__e, PrimFunc(symsubst), V3575, V3576, tmp1676)


__e.Return(PrimCons(tmp1675, tmp1677))
return


} else {
__e.Return(V3577)
return
}


}


}, 3)

tmp1682 := Call(__e, ns2_1set, symsubst, tmp1673)


_ = tmp1682

tmp1683 := MakeNative(func(__e *ControlFlow) {
V3578 := __e.Get(1)
_ = V3578
tmp1684 := Call(__e, PrimFunc(symshen_4app), V3578, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1684)
return


}, 1)

tmp1685 := Call(__e, ns2_1set, symexplode, tmp1683)


_ = tmp1685

tmp1686 := MakeNative(func(__e *ControlFlow) {
V3581 := __e.Get(1)
_ = V3581
tmp1693 := PrimEqual(MakeString(""), V3581)

if True == tmp1693 {
__e.Return(Nil)
return
} else {
tmp1691 := Call(__e, PrimFunc(symshen_4_7string_2), V3581)


if True == tmp1691 {
tmp1687 := Call(__e, PrimFunc(symhdstr), V3581)


tmp1688 := PrimTailString(V3581)

tmp1689 := Call(__e, PrimFunc(symshen_4explode_1h), tmp1688)


__e.Return(PrimCons(tmp1687, tmp1689))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in explode-h")))
return
}


}


}, 1)

tmp1694 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1686)


_ = tmp1694

tmp1695 := MakeNative(func(__e *ControlFlow) {
V3582 := __e.Get(1)
_ = V3582
tmp1698 := PrimEqual(V3582, MakeString(""))

var ifres1696 Obj

if True == tmp1698 {
ifres1696 = MakeString("")


} else {
tmp1697 := Call(__e, PrimFunc(symshen_4app), V3582, MakeString("/"), symshen_4a)


ifres1696 = tmp1697


}

__e.Return(PrimSet(sym_dhome_1directory_d, ifres1696))
return


}, 1)

tmp1699 := Call(__e, ns2_1set, symcd, tmp1695)


_ = tmp1699

tmp1700 := MakeNative(func(__e *ControlFlow) {
V3583 := __e.Get(1)
_ = V3583
V3584 := __e.Get(2)
_ = V3584
__e.TailApply(PrimFunc(symshen_4map_1h), V3583, V3584, Nil)
return
}, 2)

tmp1701 := Call(__e, ns2_1set, symmap, tmp1700)


_ = tmp1701

tmp1702 := MakeNative(func(__e *ControlFlow) {
V3585 := __e.Get(1)
_ = V3585
V3586 := __e.Get(2)
_ = V3586
V3587 := __e.Get(3)
_ = V3587
tmp1710 := PrimEqual(Nil, V3586)

if True == tmp1710 {
__e.TailApply(PrimFunc(symreverse), V3587)
return
} else {
tmp1708 := PrimIsPair(V3586)

if True == tmp1708 {
tmp1703 := PrimTail(V3586)

tmp1704 := PrimHead(V3586)

tmp1705 := Call(__e, V3585, tmp1704)


tmp1706 := PrimCons(tmp1705, V3587)

__e.TailApply(PrimFunc(symshen_4map_1h), V3585, tmp1703, tmp1706)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4map_1h)
return
}


}


}, 3)

tmp1711 := Call(__e, ns2_1set, symshen_4map_1h, tmp1702)


_ = tmp1711

tmp1712 := MakeNative(func(__e *ControlFlow) {
V3588 := __e.Get(1)
_ = V3588
__e.TailApply(PrimFunc(symshen_4length_1h), V3588, MakeNumber(0))
return
}, 1)

tmp1713 := Call(__e, ns2_1set, symlength, tmp1712)


_ = tmp1713

tmp1714 := MakeNative(func(__e *ControlFlow) {
V3593 := __e.Get(1)
_ = V3593
V3594 := __e.Get(2)
_ = V3594
tmp1718 := PrimEqual(Nil, V3593)

if True == tmp1718 {
__e.Return(V3594)
return
} else {
tmp1715 := PrimTail(V3593)

tmp1716 := PrimNumberAdd(V3594, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4length_1h), tmp1715, tmp1716)
return


}


}, 2)

tmp1719 := Call(__e, ns2_1set, symshen_4length_1h, tmp1714)


_ = tmp1719

tmp1720 := MakeNative(func(__e *ControlFlow) {
V3600 := __e.Get(1)
_ = V3600
V3601 := __e.Get(2)
_ = V3601
tmp1728 := PrimEqual(V3600, V3601)

if True == tmp1728 {
__e.Return(MakeNumber(1))
return
} else {
tmp1726 := PrimIsPair(V3601)

if True == tmp1726 {
tmp1721 := PrimHead(V3601)

tmp1722 := Call(__e, PrimFunc(symoccurrences), V3600, tmp1721)


tmp1723 := PrimTail(V3601)

tmp1724 := Call(__e, PrimFunc(symoccurrences), V3600, tmp1723)


__e.Return(PrimNumberAdd(tmp1722, tmp1724))
return


} else {
__e.Return(MakeNumber(0))
return
}


}


}, 2)

tmp1729 := Call(__e, ns2_1set, symoccurrences, tmp1720)


_ = tmp1729

tmp1730 := MakeNative(func(__e *ControlFlow) {
V3606 := __e.Get(1)
_ = V3606
V3607 := __e.Get(2)
_ = V3607
tmp1743 := PrimEqual(MakeNumber(1), V3606)

var ifres1740 Obj

if True == tmp1743 {
tmp1742 := PrimIsPair(V3607)

var ifres1741 Obj

if True == tmp1742 {
ifres1741 = True


} else {
ifres1741 = False


}

ifres1740 = ifres1741


} else {
ifres1740 = False


}

if True == ifres1740 {
__e.Return(PrimHead(V3607))
return
} else {
tmp1738 := PrimIsPair(V3607)

if True == tmp1738 {
tmp1731 := PrimNumberSubtract(V3606, MakeNumber(1))

tmp1732 := PrimTail(V3607)

__e.TailApply(PrimFunc(symnth), tmp1731, tmp1732)
return


} else {
tmp1733 := Call(__e, PrimFunc(symshen_4app), V3607, MakeString("\n"), symshen_4a)


tmp1734 := PrimStringConcat(MakeString(", "), tmp1733)

tmp1735 := Call(__e, PrimFunc(symshen_4app), V3606, tmp1734, symshen_4a)


tmp1736 := PrimStringConcat(MakeString("nth applied to "), tmp1735)

__e.Return(PrimSimpleError(tmp1736))
return


}


}


}, 2)

tmp1744 := Call(__e, ns2_1set, symnth, tmp1730)


_ = tmp1744

tmp1745 := MakeNative(func(__e *ControlFlow) {
V3608 := __e.Get(1)
_ = V3608
tmp1752 := PrimIsNumber(V3608)

if True == tmp1752 {
tmp1747 := MakeNative(func(__e *ControlFlow) {
Abs := __e.Get(1)
_ = Abs
tmp1748 := Call(__e, PrimFunc(symshen_4magless), Abs, MakeNumber(1))


__e.TailApply(PrimFunc(symshen_4integer_1test_2), Abs, tmp1748)
return


}, 1)

tmp1749 := Call(__e, PrimFunc(symshen_4abs), V3608)


tmp1750 := Call(__e, tmp1747, tmp1749)


if True == tmp1750 {
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

tmp1753 := Call(__e, ns2_1set, syminteger_2, tmp1745)


_ = tmp1753

tmp1754 := MakeNative(func(__e *ControlFlow) {
V3609 := __e.Get(1)
_ = V3609
tmp1756 := PrimGreatThan(V3609, MakeNumber(0))

if True == tmp1756 {
__e.Return(V3609)
return
} else {
__e.Return(PrimNumberSubtract(MakeNumber(0), V3609))
return
}


}, 1)

tmp1757 := Call(__e, ns2_1set, symshen_4abs, tmp1754)


_ = tmp1757

tmp1758 := MakeNative(func(__e *ControlFlow) {
V3610 := __e.Get(1)
_ = V3610
V3611 := __e.Get(2)
_ = V3611
tmp1759 := MakeNative(func(__e *ControlFlow) {
Nx2 := __e.Get(1)
_ = Nx2
tmp1761 := PrimGreatThan(Nx2, V3610)

if True == tmp1761 {
__e.Return(V3611)
return
} else {
__e.TailApply(PrimFunc(symshen_4magless), V3610, Nx2)
return
}


}, 1)

tmp1762 := PrimNumberMultiply(V3611, MakeNumber(2))

__e.TailApply(tmp1759, tmp1762)
return


}, 2)

tmp1763 := Call(__e, ns2_1set, symshen_4magless, tmp1758)


_ = tmp1763

tmp1764 := MakeNative(func(__e *ControlFlow) {
V3615 := __e.Get(1)
_ = V3615
V3616 := __e.Get(2)
_ = V3616
tmp1772 := PrimEqual(MakeNumber(0), V3615)

if True == tmp1772 {
__e.Return(True)
return
} else {
tmp1770 := PrimGreatThan(MakeNumber(1), V3615)

if True == tmp1770 {
__e.Return(False)
return
} else {
tmp1765 := MakeNative(func(__e *ControlFlow) {
Abs_1N := __e.Get(1)
_ = Abs_1N
tmp1767 := PrimGreatThan(MakeNumber(0), Abs_1N)

if True == tmp1767 {
__e.Return(PrimIsInteger(V3615))
return
} else {
__e.TailApply(PrimFunc(symshen_4integer_1test_2), Abs_1N, V3616)
return
}


}, 1)

tmp1768 := PrimNumberSubtract(V3615, V3616)

__e.TailApply(tmp1765, tmp1768)
return


}


}


}, 2)

tmp1773 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1764)


_ = tmp1773

tmp1774 := MakeNative(func(__e *ControlFlow) {
V3623 := __e.Get(1)
_ = V3623
V3624 := __e.Get(2)
_ = V3624
tmp1782 := PrimEqual(Nil, V3624)

if True == tmp1782 {
__e.Return(Nil)
return
} else {
tmp1780 := PrimIsPair(V3624)

if True == tmp1780 {
tmp1775 := PrimHead(V3624)

tmp1776 := Call(__e, V3623, tmp1775)


tmp1777 := PrimTail(V3624)

tmp1778 := Call(__e, PrimFunc(symmapcan), V3623, tmp1777)


__e.TailApply(PrimFunc(symappend), tmp1776, tmp1778)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to mapcan over a non-list\n")))
return
}


}


}, 2)

tmp1783 := Call(__e, ns2_1set, symmapcan, tmp1774)


_ = tmp1783

tmp1784 := MakeNative(func(__e *ControlFlow) {
V3630 := __e.Get(1)
_ = V3630
V3631 := __e.Get(2)
_ = V3631
tmp1786 := PrimEqual(V3630, V3631)

if True == tmp1786 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1787 := Call(__e, ns2_1set, sym_a_a, tmp1784)


_ = tmp1787

tmp1788 := MakeNative(func(__e *ControlFlow) {
V3632 := __e.Get(1)
_ = V3632
tmp1798 := PrimIsSymbol(V3632)

if True == tmp1798 {
tmp1790 := MakeNative(func(__e *ControlFlow) {
Val := __e.Get(1)
_ = Val
tmp1792 := PrimEqual(Val, symshen_4this_1symbol_1is_1unbound)

if True == tmp1792 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}


}, 1)

tmp1793 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(V3632))
return
}, 0)

tmp1794 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1795 := Call(__e, try_1catch, tmp1793, tmp1794)


tmp1796 := Call(__e, tmp1790, tmp1795)


if True == tmp1796 {
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

tmp1799 := Call(__e, ns2_1set, symbound_2, tmp1788)


_ = tmp1799

tmp1800 := MakeNative(func(__e *ControlFlow) {
V3633 := __e.Get(1)
_ = V3633
tmp1806 := PrimEqual(MakeString(""), V3633)

if True == tmp1806 {
__e.Return(Nil)
return
} else {
tmp1801 := PrimPos(V3633, MakeNumber(0))

tmp1802 := PrimStringToNumber(tmp1801)

tmp1803 := PrimTailString(V3633)

tmp1804 := Call(__e, PrimFunc(symshen_4string_1_6bytes), tmp1803)


__e.Return(PrimCons(tmp1802, tmp1804))
return


}


}, 1)

tmp1807 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1800)


_ = tmp1807

tmp1808 := MakeNative(func(__e *ControlFlow) {
V3634 := __e.Get(1)
_ = V3634
__e.Return(PrimSet(symshen_4_dmaxinferences_d, V3634))
return
}, 1)

tmp1809 := Call(__e, ns2_1set, symmaxinferences, tmp1808)


_ = tmp1809

tmp1810 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dinfs_d))
return
}, 0)

tmp1811 := Call(__e, ns2_1set, syminferences, tmp1810)


_ = tmp1811

tmp1812 := MakeNative(func(__e *ControlFlow) {
V3635 := __e.Get(1)
_ = V3635
__e.Return(V3635)
return
}, 1)

tmp1813 := Call(__e, ns2_1set, symprotect, tmp1812)


_ = tmp1813

tmp1814 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstoutput_d))
return
}, 0)

tmp1815 := Call(__e, ns2_1set, symstoutput, tmp1814)


_ = tmp1815

tmp1816 := MakeNative(func(__e *ControlFlow) {
V3636 := __e.Get(1)
_ = V3636
tmp1817 := MakeNative(func(__e *ControlFlow) {
Symbol := __e.Get(1)
_ = Symbol
tmp1821 := PrimIsSymbol(Symbol)

if True == tmp1821 {
__e.Return(Symbol)
return
} else {
tmp1818 := Call(__e, PrimFunc(symshen_4app), V3636, MakeString(" to a symbol"), symshen_4s)


tmp1819 := PrimStringConcat(MakeString("cannot intern "), tmp1818)

__e.Return(PrimSimpleError(tmp1819))
return


}


}, 1)

tmp1822 := PrimIntern(V3636)

__e.TailApply(tmp1817, tmp1822)
return


}, 1)

tmp1823 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1816)


_ = tmp1823

tmp1824 := MakeNative(func(__e *ControlFlow) {
V3639 := __e.Get(1)
_ = V3639
tmp1828 := PrimEqual(sym_7, V3639)

if True == tmp1828 {
__e.Return(PrimSet(symshen_4_doptimise_d, True))
return
} else {
tmp1826 := PrimEqual(sym_1, V3639)

if True == tmp1826 {
__e.Return(PrimSet(symshen_4_doptimise_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
return
}


}


}, 1)

tmp1829 := Call(__e, ns2_1set, symoptimise, tmp1824)


_ = tmp1829

tmp1830 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dos_d))
return
}, 0)

tmp1831 := Call(__e, ns2_1set, symos, tmp1830)


_ = tmp1831

tmp1832 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dlanguage_d))
return
}, 0)

tmp1833 := Call(__e, ns2_1set, symlanguage, tmp1832)


_ = tmp1833

tmp1834 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dversion_d))
return
}, 0)

tmp1835 := Call(__e, ns2_1set, symversion, tmp1834)


_ = tmp1835

tmp1836 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dport_d))
return
}, 0)

tmp1837 := Call(__e, ns2_1set, symport, tmp1836)


_ = tmp1837

tmp1838 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dporters_d))
return
}, 0)

tmp1839 := Call(__e, ns2_1set, symporters, tmp1838)


_ = tmp1839

tmp1840 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dimplementation_d))
return
}, 0)

tmp1841 := Call(__e, ns2_1set, symimplementation, tmp1840)


_ = tmp1841

tmp1842 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_drelease_d))
return
}, 0)

tmp1843 := Call(__e, ns2_1set, symrelease, tmp1842)


_ = tmp1843

tmp1844 := MakeNative(func(__e *ControlFlow) {
V3640 := __e.Get(1)
_ = V3640
tmp1849 := PrimEqual(symnull, V3640)

if True == tmp1849 {
__e.Return(True)
return
} else {
tmp1845 := MakeNative(func(__e *ControlFlow) {
tmp1846 := Call(__e, PrimFunc(symexternal), V3640)


_ = tmp1846

__e.Return(True)
return


}, 0)

tmp1847 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1845, tmp1847)
return


}


}, 1)

tmp1850 := Call(__e, ns2_1set, sympackage_2, tmp1844)


_ = tmp1850

tmp1851 := MakeNative(func(__e *ControlFlow) {
V3641 := __e.Get(1)
_ = V3641
tmp1852 := MakeNative(func(__e *ControlFlow) {
Assoc := __e.Get(1)
_ = Assoc
tmp1855 := Call(__e, PrimFunc(symempty_2), Assoc)


if True == tmp1855 {
tmp1853 := Call(__e, PrimFunc(symshen_4app), V3641, MakeString(" has no lambda expansion\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp1853))
return


} else {
__e.Return(PrimTail(Assoc))
return
}


}, 1)

tmp1856 := PrimValue(symshen_4_dlambdatable_d)

tmp1857 := Call(__e, PrimFunc(symassoc), V3641, tmp1856)


__e.TailApply(tmp1852, tmp1857)
return


}, 1)

tmp1858 := Call(__e, ns2_1set, symfn, tmp1851)


_ = tmp1858

tmp1859 := MakeNative(func(__e *ControlFlow) {
__e.Return(sym_4_4_4)
return
}, 0)

tmp1860 := Call(__e, ns2_1set, symfail, tmp1859)


_ = tmp1860

tmp1861 := MakeNative(func(__e *ControlFlow) {
V3644 := __e.Get(1)
_ = V3644
tmp1865 := PrimEqual(sym_7, V3644)

if True == tmp1865 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
return
} else {
tmp1863 := PrimEqual(sym_1, V3644)

if True == tmp1863 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
return
}


}


}, 1)

tmp1866 := Call(__e, ns2_1set, symenable_1type_1theory, tmp1861)


_ = tmp1866

tmp1867 := MakeNative(func(__e *ControlFlow) {
V3647 := __e.Get(1)
_ = V3647
tmp1871 := PrimEqual(sym_7, V3647)

if True == tmp1871 {
__e.Return(PrimSet(symshen_4_dtc_d, True))
return
} else {
tmp1869 := PrimEqual(sym_1, V3647)

if True == tmp1869 {
__e.Return(PrimSet(symshen_4_dtc_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
return
}


}


}, 1)

tmp1872 := Call(__e, ns2_1set, symtc, tmp1867)


_ = tmp1872

tmp1873 := MakeNative(func(__e *ControlFlow) {
V3648 := __e.Get(1)
_ = V3648
tmp1874 := PrimValue(symshen_4_dsigf_d)

tmp1875 := Call(__e, PrimFunc(symshen_4unassoc), V3648, tmp1874)


_ = tmp1875

__e.Return(V3648)
return


}, 1)

tmp1876 := Call(__e, ns2_1set, symdestroy, tmp1873)


_ = tmp1876

tmp1877 := MakeNative(func(__e *ControlFlow) {
V3649 := __e.Get(1)
_ = V3649
V3650 := __e.Get(2)
_ = V3650
tmp1878 := MakeNative(func(__e *ControlFlow) {
Assoc := __e.Get(1)
_ = Assoc
tmp1879 := MakeNative(func(__e *ControlFlow) {
Remove := __e.Get(1)
_ = Remove
__e.Return(PrimSet(symshen_4_dsigf_d, Remove))
return
}, 1)

tmp1880 := Call(__e, PrimFunc(symremove), Assoc, V3650)


__e.TailApply(tmp1879, tmp1880)
return


}, 1)

tmp1881 := Call(__e, PrimFunc(symassoc), V3649, V3650)


__e.TailApply(tmp1878, tmp1881)
return


}, 2)

tmp1882 := Call(__e, ns2_1set, symshen_4unassoc, tmp1877)


_ = tmp1882

tmp1883 := MakeNative(func(__e *ControlFlow) {
V3651 := __e.Get(1)
_ = V3651
tmp1887 := Call(__e, PrimFunc(sympackage_2), V3651)


if True == tmp1887 {
__e.Return(PrimSet(symshen_4_dpackage_d, V3651))
return
} else {
tmp1884 := Call(__e, PrimFunc(symshen_4app), V3651, MakeString(" does not exist\n"), symshen_4a)


tmp1885 := PrimStringConcat(MakeString("package "), tmp1884)

__e.Return(PrimSimpleError(tmp1885))
return


}


}, 1)

tmp1888 := Call(__e, ns2_1set, symin_1package, tmp1883)


_ = tmp1888

tmp1889 := MakeNative(func(__e *ControlFlow) {
V3652 := __e.Get(1)
_ = V3652
V3653 := __e.Get(2)
_ = V3653
tmp1890 := MakeNative(func(__e *ControlFlow) {
Stream := __e.Get(1)
_ = Stream
tmp1891 := MakeNative(func(__e *ControlFlow) {
String := __e.Get(1)
_ = String
tmp1892 := MakeNative(func(__e *ControlFlow) {
Write := __e.Get(1)
_ = Write
tmp1893 := MakeNative(func(__e *ControlFlow) {
Close := __e.Get(1)
_ = Close
__e.Return(V3653)
return
}, 1)

tmp1894 := PrimCloseStream(Stream)

__e.TailApply(tmp1893, tmp1894)
return


}, 1)

tmp1895 := Call(__e, PrimFunc(sympr), String, Stream)


__e.TailApply(tmp1892, tmp1895)
return


}, 1)

tmp1899 := PrimIsString(V3653)

var ifres1896 Obj

if True == tmp1899 {
tmp1897 := Call(__e, PrimFunc(symshen_4app), V3653, MakeString("\n\n"), symshen_4a)


ifres1896 = tmp1897


} else {
tmp1898 := Call(__e, PrimFunc(symshen_4app), V3653, MakeString("\n\n"), symshen_4s)


ifres1896 = tmp1898


}

__e.TailApply(tmp1891, ifres1896)
return


}, 1)

tmp1900 := PrimOpenStream(V3652, symout)

__e.TailApply(tmp1890, tmp1900)
return


}, 2)

tmp1901 := Call(__e, ns2_1set, symwrite_1to_1file, tmp1889)


_ = tmp1901

tmp1902 := MakeNative(func(__e *ControlFlow) {
tmp1903 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp1903)
return


}, 0)

tmp1904 := Call(__e, ns2_1set, symfresh, tmp1902)


_ = tmp1904

tmp1905 := MakeNative(func(__e *ControlFlow) {
V3654 := __e.Get(1)
_ = V3654
V3655 := __e.Get(2)
_ = V3655
tmp1906 := MakeNative(func(__e *ControlFlow) {
AssertArity := __e.Get(1)
_ = AssertArity
tmp1907 := MakeNative(func(__e *ControlFlow) {
LambdaEntry := __e.Get(1)
_ = LambdaEntry
tmp1908 := MakeNative(func(__e *ControlFlow) {
Update := __e.Get(1)
_ = Update
__e.Return(V3654)
return
}, 1)

tmp1909 := PrimValue(symshen_4_dlambdatable_d)

tmp1910 := PrimCons(LambdaEntry, tmp1909)

tmp1911 := PrimSet(symshen_4_dlambdatable_d, tmp1910)

__e.TailApply(tmp1908, tmp1911)
return


}, 1)

tmp1912 := Call(__e, PrimFunc(symshen_4lambda_1entry), V3654)


__e.TailApply(tmp1907, tmp1912)
return


}, 1)

tmp1913 := PrimValue(sym_dproperty_1vector_d)

tmp1914 := Call(__e, PrimFunc(symput), V3654, symarity, V3655, tmp1913)


__e.TailApply(tmp1906, tmp1914)
return


}, 2)

tmp1915 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp1905)


_ = tmp1915

tmp1916 := MakeNative(func(__e *ControlFlow) {
V3658 := __e.Get(1)
_ = V3658
V3659 := __e.Get(2)
_ = V3659
tmp1940 := PrimEqual(MakeNumber(0), V3659)

if True == tmp1940 {
tmp1917 := PrimValue(symshen_4_dspecial_d)

tmp1918 := Call(__e, PrimFunc(symremove), V3658, tmp1917)


tmp1919 := PrimSet(symshen_4_dspecial_d, tmp1918)

_ = tmp1919

tmp1920 := PrimValue(symshen_4_dextraspecial_d)

tmp1921 := Call(__e, PrimFunc(symremove), V3658, tmp1920)


tmp1922 := PrimSet(symshen_4_dextraspecial_d, tmp1921)

_ = tmp1922

__e.Return(V3658)
return


} else {
tmp1938 := PrimEqual(MakeNumber(1), V3659)

if True == tmp1938 {
tmp1923 := PrimValue(symshen_4_dspecial_d)

tmp1924 := Call(__e, PrimFunc(symadjoin), V3658, tmp1923)


tmp1925 := PrimSet(symshen_4_dspecial_d, tmp1924)

_ = tmp1925

tmp1926 := PrimValue(symshen_4_dextraspecial_d)

tmp1927 := Call(__e, PrimFunc(symremove), V3658, tmp1926)


tmp1928 := PrimSet(symshen_4_dextraspecial_d, tmp1927)

_ = tmp1928

__e.Return(V3658)
return


} else {
tmp1936 := PrimEqual(MakeNumber(2), V3659)

if True == tmp1936 {
tmp1929 := PrimValue(symshen_4_dspecial_d)

tmp1930 := Call(__e, PrimFunc(symremove), V3658, tmp1929)


tmp1931 := PrimSet(symshen_4_dspecial_d, tmp1930)

_ = tmp1931

tmp1932 := PrimValue(symshen_4_dextraspecial_d)

tmp1933 := Call(__e, PrimFunc(symadjoin), V3658, tmp1932)


tmp1934 := PrimSet(symshen_4_dextraspecial_d, tmp1933)

_ = tmp1934

__e.Return(V3658)
return


} else {
__e.Return(PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n")))
return
}


}


}


}, 2)

__e.TailApply(ns2_1set, symspecialise, tmp1916)
return




}, 0)

