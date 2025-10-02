package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp9945 := MakeNative(func(__e *ControlFlow) {
V1212 := __e.Get(1)
_ = V1212
__e.TailApply(PrimFunc(symshen_4assert_d), V1212, symshen_4top)
return
}, 1)

tmp9946 := Call(__e, ns2_1set, symasserta, tmp9945)


_ = tmp9946

tmp9947 := MakeNative(func(__e *ControlFlow) {
V1213 := __e.Get(1)
_ = V1213
__e.TailApply(PrimFunc(symshen_4assert_d), V1213, symshen_4bottom)
return
}, 1)

tmp9948 := Call(__e, ns2_1set, symassertz, tmp9947)


_ = tmp9948

tmp9949 := MakeNative(func(__e *ControlFlow) {
V1214 := __e.Get(1)
_ = V1214
V1215 := __e.Get(2)
_ = V1215
tmp9983 := PrimIsPair(V1214)

var ifres9974 Obj

if True == tmp9983 {
tmp9981 := PrimTail(V1214)

tmp9982 := PrimIsPair(tmp9981)

var ifres9976 Obj

if True == tmp9982 {
tmp9978 := PrimTail(V1214)

tmp9979 := PrimHead(tmp9978)

tmp9980 := PrimEqual(sym_5_1_1, tmp9979)

var ifres9977 Obj

if True == tmp9980 {
ifres9977 = True


} else {
ifres9977 = False


}

ifres9976 = ifres9977


} else {
ifres9976 = False


}

var ifres9975 Obj

if True == ifres9976 {
ifres9975 = True


} else {
ifres9975 = False


}

ifres9974 = ifres9975


} else {
ifres9974 = False


}

if True == ifres9974 {
tmp9950 := MakeNative(func(__e *ControlFlow) {
W1216 := __e.Get(1)
_ = W1216
tmp9951 := MakeNative(func(__e *ControlFlow) {
W1217 := __e.Get(1)
_ = W1217
tmp9952 := MakeNative(func(__e *ControlFlow) {
W1218 := __e.Get(1)
_ = W1218
tmp9953 := MakeNative(func(__e *ControlFlow) {
W1219 := __e.Get(1)
_ = W1219
tmp9954 := MakeNative(func(__e *ControlFlow) {
W1220 := __e.Get(1)
_ = W1220
tmp9955 := MakeNative(func(__e *ControlFlow) {
W1221 := __e.Get(1)
_ = W1221
tmp9956 := MakeNative(func(__e *ControlFlow) {
W1222 := __e.Get(1)
_ = W1222
__e.Return(W1216)
return
}, 1)

tmp9957 := PrimTail(V1214)

tmp9958 := PrimTail(tmp9957)

tmp9959 := Call(__e, PrimFunc(symshen_4insert_1info), W1216, W1217, tmp9958, V1214, V1215)


__e.TailApply(tmp9956, tmp9959)
return


}, 1)

tmp9965 := PrimEqual(W1220, MakeNumber(-1))

var ifres9960 Obj

if True == tmp9965 {
tmp9961 := Call(__e, PrimFunc(symshen_4create_1skeleton), W1216, W1219)


tmp9962 := Call(__e, PrimFunc(symeval), tmp9961)


_ = tmp9962

tmp9963 := PrimValue(sym_dproperty_1vector_d)

tmp9964 := Call(__e, PrimFunc(symput), W1216, symshen_4dynamic, Nil, tmp9963)


ifres9960 = tmp9964


} else {
ifres9960 = symshen_4skip


}

__e.TailApply(tmp9955, ifres9960)
return


}, 1)

tmp9966 := Call(__e, PrimFunc(symarity), W1216)


__e.TailApply(tmp9954, tmp9966)
return


}, 1)

tmp9967 := Call(__e, PrimFunc(symshen_4parameters), W1218)


__e.TailApply(tmp9953, tmp9967)
return


}, 1)

tmp9968 := Call(__e, PrimFunc(symlength), W1217)


__e.TailApply(tmp9952, tmp9968)
return


}, 1)

tmp9969 := PrimHead(V1214)

tmp9970 := Call(__e, PrimFunc(symshen_4terms), tmp9969)


__e.TailApply(tmp9951, tmp9970)
return


}, 1)

tmp9971 := PrimHead(V1214)

tmp9972 := Call(__e, PrimFunc(symshen_4predicate), tmp9971)


__e.TailApply(tmp9950, tmp9972)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.assert*")))
return
}


}, 2)

tmp9984 := Call(__e, ns2_1set, symshen_4assert_d, tmp9949)


_ = tmp9984

tmp9985 := MakeNative(func(__e *ControlFlow) {
V1225 := __e.Get(1)
_ = V1225
tmp9987 := PrimIsPair(V1225)

if True == tmp9987 {
__e.Return(PrimHead(V1225))
return
} else {
__e.Return(V1225)
return
}


}, 1)

tmp9988 := Call(__e, ns2_1set, symshen_4predicate, tmp9985)


_ = tmp9988

tmp9989 := MakeNative(func(__e *ControlFlow) {
V1230 := __e.Get(1)
_ = V1230
tmp9991 := PrimIsPair(V1230)

if True == tmp9991 {
__e.Return(PrimTail(V1230))
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp9992 := Call(__e, ns2_1set, symshen_4terms, tmp9989)


_ = tmp9992

tmp9993 := MakeNative(func(__e *ControlFlow) {
V1231 := __e.Get(1)
_ = V1231
V1232 := __e.Get(2)
_ = V1232
tmp9994 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1231, V1232)


tmp9995 := PrimCons(V1231, tmp9994)

__e.Return(PrimCons(symdefprolog, tmp9995))
return


}, 2)

tmp9996 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp9993)


_ = tmp9996

tmp9997 := MakeNative(func(__e *ControlFlow) {
V1233 := __e.Get(1)
_ = V1233
V1234 := __e.Get(2)
_ = V1234
tmp9998 := Call(__e, PrimFunc(symshen_4cons_1form), V1234)


tmp9999 := PrimCons(symshen_4dynamic, Nil)

tmp10000 := PrimCons(V1233, tmp9999)

tmp10001 := PrimCons(symget, tmp10000)

tmp10002 := PrimCons(tmp10001, Nil)

tmp10003 := PrimCons(tmp9998, tmp10002)

tmp10004 := PrimCons(symshen_4call_1dynamic, tmp10003)

tmp10005 := PrimIntern(MakeString(";"))

tmp10006 := PrimCons(tmp10005, Nil)

tmp10007 := PrimCons(tmp10004, tmp10006)

tmp10008 := PrimCons(sym_5_1_1, tmp10007)

__e.TailApply(PrimFunc(symappend), V1234, tmp10008)
return


}, 2)

tmp10009 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp9997)


_ = tmp10009

tmp10010 := MakeNative(func(__e *ControlFlow) {
V1235 := __e.Get(1)
_ = V1235
V1236 := __e.Get(2)
_ = V1236
V1237 := __e.Get(3)
_ = V1237
V1238 := __e.Get(4)
_ = V1238
V1239 := __e.Get(5)
_ = V1239
tmp10011 := MakeNative(func(__e *ControlFlow) {
W1240 := __e.Get(1)
_ = W1240
tmp10012 := MakeNative(func(__e *ControlFlow) {
W1241 := __e.Get(1)
_ = W1241
tmp10013 := MakeNative(func(__e *ControlFlow) {
W1242 := __e.Get(1)
_ = W1242
tmp10014 := MakeNative(func(__e *ControlFlow) {
W1243 := __e.Get(1)
_ = W1243
tmp10015 := MakeNative(func(__e *ControlFlow) {
W1244 := __e.Get(1)
_ = W1244
tmp10016 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V1235, symshen_4dynamic, W1244, tmp10016)
return


}, 1)

tmp10021 := PrimEqual(V1239, symshen_4top)

var ifres10017 Obj

if True == tmp10021 {
tmp10018 := PrimCons(W1242, W1243)

ifres10017 = tmp10018


} else {
tmp10019 := PrimCons(W1242, Nil)

tmp10020 := Call(__e, PrimFunc(symappend), W1243, tmp10019)


ifres10017 = tmp10020


}

__e.TailApply(tmp10015, ifres10017)
return


}, 1)

tmp10022 := PrimValue(sym_dproperty_1vector_d)

tmp10023 := Call(__e, PrimFunc(symget), V1235, symshen_4dynamic, tmp10022)


__e.TailApply(tmp10014, tmp10023)
return


}, 1)

tmp10024 := Call(__e, PrimFunc(symfn), W1240)


tmp10025 := PrimCons(W1240, V1238)

tmp10026 := PrimCons(tmp10024, tmp10025)

__e.TailApply(tmp10013, tmp10026)
return


}, 1)

tmp10027 := PrimCons(W1240, Nil)

tmp10028 := PrimCons(symdefprolog, tmp10027)

tmp10029 := PrimCons(sym_5_1_1, V1237)

tmp10030 := Call(__e, PrimFunc(symappend), V1236, tmp10029)


tmp10031 := Call(__e, PrimFunc(symappend), tmp10028, tmp10030)


tmp10032 := Call(__e, PrimFunc(symeval), tmp10031)


__e.TailApply(tmp10012, tmp10032)
return


}, 1)

tmp10033 := Call(__e, PrimFunc(symgensym), symshen_4g)


__e.TailApply(tmp10011, tmp10033)
return


}, 5)

tmp10034 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10010)


_ = tmp10034

tmp10035 := MakeNative(func(__e *ControlFlow) {
tmp10036 := MakeNative(func(__e *ControlFlow) {
W1245 := __e.Get(1)
_ = W1245
tmp10037 := MakeNative(func(__e *ControlFlow) {
W1246 := __e.Get(1)
_ = W1246
__e.Return(W1246)
return
}, 1)

tmp10043 := Call(__e, PrimFunc(symempty_2), W1245)


var ifres10038 Obj

if True == tmp10043 {
tmp10039 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres10038 = tmp10039


} else {
tmp10040 := PrimTail(W1245)

tmp10041 := PrimSet(symshen_4_dnames_d, tmp10040)

_ = tmp10041

tmp10042 := PrimHead(W1245)

ifres10038 = tmp10042


}

__e.TailApply(tmp10037, ifres10038)
return


}, 1)

tmp10044 := PrimValue(symshen_4_dnames_d)

__e.TailApply(tmp10036, tmp10044)
return


}, 0)

tmp10045 := Call(__e, ns2_1set, symshen_4newname, tmp10035)


_ = tmp10045

tmp10046 := MakeNative(func(__e *ControlFlow) {
V1247 := __e.Get(1)
_ = V1247
V1248 := __e.Get(2)
_ = V1248
V1249 := __e.Get(3)
_ = V1249
V1250 := __e.Get(4)
_ = V1250
V1251 := __e.Get(5)
_ = V1251
V1252 := __e.Get(6)
_ = V1252
tmp10047 := MakeNative(func(__e *ControlFlow) {
W1253 := __e.Get(1)
_ = W1253
tmp10058 := PrimEqual(W1253, False)

if True == tmp10058 {
tmp10056 := Call(__e, PrimFunc(symshen_4unlocked_2), V1250)


if True == tmp10056 {
tmp10048 := MakeNative(func(__e *ControlFlow) {
W1257 := __e.Get(1)
_ = W1257
tmp10053 := PrimIsPair(W1257)

if True == tmp10053 {
tmp10049 := MakeNative(func(__e *ControlFlow) {
W1258 := __e.Get(1)
_ = W1258
tmp10050 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10050

__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1247, W1258, V1249, V1250, V1251, V1252)
return


}, 1)

tmp10051 := PrimTail(W1257)

__e.TailApply(tmp10049, tmp10051)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10054 := Call(__e, PrimFunc(symshen_4lazyderef), V1248, V1249)


__e.TailApply(tmp10048, tmp10054)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1253)
return
}


}, 1)

tmp10073 := Call(__e, PrimFunc(symshen_4unlocked_2), V1250)


var ifres10059 Obj

if True == tmp10073 {
tmp10060 := MakeNative(func(__e *ControlFlow) {
W1254 := __e.Get(1)
_ = W1254
tmp10070 := PrimIsPair(W1254)

if True == tmp10070 {
tmp10061 := MakeNative(func(__e *ControlFlow) {
W1255 := __e.Get(1)
_ = W1255
tmp10066 := PrimIsPair(W1255)

if True == tmp10066 {
tmp10062 := MakeNative(func(__e *ControlFlow) {
W1256 := __e.Get(1)
_ = W1256
tmp10063 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10063

__e.TailApply(PrimFunc(symshen_4callrec), W1256, V1247, V1249, V1250, V1251, V1252)
return


}, 1)

tmp10064 := PrimHead(W1255)

__e.TailApply(tmp10062, tmp10064)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10067 := PrimHead(W1254)

tmp10068 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10067, V1249)


__e.TailApply(tmp10061, tmp10068)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10071 := Call(__e, PrimFunc(symshen_4lazyderef), V1248, V1249)


tmp10072 := Call(__e, tmp10060, tmp10071)


ifres10059 = tmp10072


} else {
ifres10059 = False


}

__e.TailApply(tmp10047, ifres10059)
return


}, 6)

tmp10074 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10046)


_ = tmp10074

tmp10075 := MakeNative(func(__e *ControlFlow) {
V1259 := __e.Get(1)
_ = V1259
V1260 := __e.Get(2)
_ = V1260
V1261 := __e.Get(3)
_ = V1261
V1262 := __e.Get(4)
_ = V1262
V1263 := __e.Get(5)
_ = V1263
V1264 := __e.Get(6)
_ = V1264
tmp10085 := PrimEqual(Nil, V1260)

if True == tmp10085 {
tmp10076 := Call(__e, V1259, V1261)


tmp10077 := Call(__e, tmp10076, V1262)


tmp10078 := Call(__e, tmp10077, V1263)


__e.TailApply(tmp10078, V1264)
return


} else {
tmp10083 := PrimIsPair(V1260)

if True == tmp10083 {
tmp10079 := PrimHead(V1260)

tmp10080 := Call(__e, V1259, tmp10079)


tmp10081 := PrimTail(V1260)

__e.TailApply(PrimFunc(symshen_4callrec), tmp10080, tmp10081, V1261, V1262, V1263, V1264)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.callrec")))
return
}


}


}, 6)

tmp10086 := Call(__e, ns2_1set, symshen_4callrec, tmp10075)


_ = tmp10086

tmp10087 := MakeNative(func(__e *ControlFlow) {
V1265 := __e.Get(1)
_ = V1265
tmp10106 := PrimIsPair(V1265)

var ifres10097 Obj

if True == tmp10106 {
tmp10104 := PrimTail(V1265)

tmp10105 := PrimIsPair(tmp10104)

var ifres10099 Obj

if True == tmp10105 {
tmp10101 := PrimTail(V1265)

tmp10102 := PrimHead(tmp10101)

tmp10103 := PrimEqual(sym_5_1_1, tmp10102)

var ifres10100 Obj

if True == tmp10103 {
ifres10100 = True


} else {
ifres10100 = False


}

ifres10099 = ifres10100


} else {
ifres10099 = False


}

var ifres10098 Obj

if True == ifres10099 {
ifres10098 = True


} else {
ifres10098 = False


}

ifres10097 = ifres10098


} else {
ifres10097 = False


}

if True == ifres10097 {
tmp10088 := MakeNative(func(__e *ControlFlow) {
W1266 := __e.Get(1)
_ = W1266
tmp10089 := MakeNative(func(__e *ControlFlow) {
W1267 := __e.Get(1)
_ = W1267
tmp10090 := Call(__e, PrimFunc(symshen_4retract_1clause), V1265, W1267)


tmp10091 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), W1266, symshen_4dynamic, tmp10090, tmp10091)
return


}, 1)

tmp10092 := PrimValue(sym_dproperty_1vector_d)

tmp10093 := Call(__e, PrimFunc(symget), W1266, symshen_4dynamic, tmp10092)


__e.TailApply(tmp10089, tmp10093)
return


}, 1)

tmp10094 := PrimHead(V1265)

tmp10095 := Call(__e, PrimFunc(symshen_4predicate), tmp10094)


__e.TailApply(tmp10088, tmp10095)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function retract")))
return
}


}, 1)

tmp10107 := Call(__e, ns2_1set, symretract, tmp10087)


_ = tmp10107

tmp10108 := MakeNative(func(__e *ControlFlow) {
V1273 := __e.Get(1)
_ = V1273
V1274 := __e.Get(2)
_ = V1274
tmp10138 := PrimEqual(Nil, V1274)

if True == tmp10138 {
__e.Return(Nil)
return
} else {
tmp10136 := PrimIsPair(V1274)

var ifres10121 Obj

if True == tmp10136 {
tmp10134 := PrimHead(V1274)

tmp10135 := PrimIsPair(tmp10134)

var ifres10123 Obj

if True == tmp10135 {
tmp10131 := PrimHead(V1274)

tmp10132 := PrimTail(tmp10131)

tmp10133 := PrimIsPair(tmp10132)

var ifres10125 Obj

if True == tmp10133 {
tmp10127 := PrimHead(V1274)

tmp10128 := PrimTail(tmp10127)

tmp10129 := PrimTail(tmp10128)

tmp10130 := PrimEqual(V1273, tmp10129)

var ifres10126 Obj

if True == tmp10130 {
ifres10126 = True


} else {
ifres10126 = False


}

ifres10125 = ifres10126


} else {
ifres10125 = False


}

var ifres10124 Obj

if True == ifres10125 {
ifres10124 = True


} else {
ifres10124 = False


}

ifres10123 = ifres10124


} else {
ifres10123 = False


}

var ifres10122 Obj

if True == ifres10123 {
ifres10122 = True


} else {
ifres10122 = False


}

ifres10121 = ifres10122


} else {
ifres10121 = False


}

if True == ifres10121 {
tmp10109 := PrimHead(V1274)

tmp10110 := PrimTail(tmp10109)

tmp10111 := PrimHead(tmp10110)

tmp10112 := PrimValue(symshen_4_dnames_d)

tmp10113 := PrimCons(tmp10111, tmp10112)

tmp10114 := PrimSet(symshen_4_dnames_d, tmp10113)

_ = tmp10114

__e.Return(PrimTail(V1274))
return


} else {
tmp10119 := PrimIsPair(V1274)

if True == tmp10119 {
tmp10115 := PrimHead(V1274)

tmp10116 := PrimTail(V1274)

tmp10117 := Call(__e, PrimFunc(symshen_4retract_1clause), V1273, tmp10116)


__e.Return(PrimCons(tmp10115, tmp10117))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.retract-clause")))
return
}


}


}


}, 2)

tmp10139 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10108)


_ = tmp10139

tmp10140 := MakeNative(func(__e *ControlFlow) {
V1275 := __e.Get(1)
_ = V1275
V1276 := __e.Get(2)
_ = V1276
tmp10141 := MakeNative(func(__e *ControlFlow) {
Z1277 := __e.Get(1)
_ = Z1277
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1277)
return
}, 1)

tmp10142 := PrimCons(V1275, V1276)

__e.TailApply(PrimFunc(symcompile), tmp10141, tmp10142)
return


}, 2)

tmp10143 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10140)


_ = tmp10143

tmp10144 := MakeNative(func(__e *ControlFlow) {
V1278 := __e.Get(1)
_ = V1278
tmp10145 := MakeNative(func(__e *ControlFlow) {
W1279 := __e.Get(1)
_ = W1279
tmp10147 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1279)


if True == tmp10147 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1279)
return
}


}, 1)

tmp10169 := PrimIsPair(V1278)

var ifres10148 Obj

if True == tmp10169 {
tmp10149 := MakeNative(func(__e *ControlFlow) {
W1280 := __e.Get(1)
_ = W1280
tmp10150 := MakeNative(func(__e *ControlFlow) {
W1281 := __e.Get(1)
_ = W1281
tmp10151 := MakeNative(func(__e *ControlFlow) {
W1282 := __e.Get(1)
_ = W1282
tmp10163 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1282)


if True == tmp10163 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10152 := MakeNative(func(__e *ControlFlow) {
W1283 := __e.Get(1)
_ = W1283
tmp10153 := MakeNative(func(__e *ControlFlow) {
W1284 := __e.Get(1)
_ = W1284
tmp10154 := MakeNative(func(__e *ControlFlow) {
W1285 := __e.Get(1)
_ = W1285
tmp10155 := MakeNative(func(__e *ControlFlow) {
W1286 := __e.Get(1)
_ = W1286
__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), W1280, W1286)
return
}, 1)

tmp10156 := MakeNative(func(__e *ControlFlow) {
Z1287 := __e.Get(1)
_ = Z1287
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1287)
return
}, 1)

tmp10157 := Call(__e, PrimFunc(symmap), tmp10156, W1283)


__e.TailApply(tmp10155, tmp10157)
return


}, 1)

tmp10158 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), W1280, W1283)


tmp10159 := Call(__e, tmp10154, tmp10158)


__e.TailApply(PrimFunc(symshen_4comb), W1284, tmp10159)
return


}, 1)

tmp10160 := Call(__e, PrimFunc(symshen_4in_1_6), W1282)


__e.TailApply(tmp10153, tmp10160)
return


}, 1)

tmp10161 := Call(__e, PrimFunc(symshen_4_5_1out), W1282)


__e.TailApply(tmp10152, tmp10161)
return


}


}, 1)

tmp10164 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1281)


__e.TailApply(tmp10151, tmp10164)
return


}, 1)

tmp10165 := Call(__e, PrimFunc(symtail), V1278)


__e.TailApply(tmp10150, tmp10165)
return


}, 1)

tmp10166 := Call(__e, PrimFunc(symhead), V1278)


tmp10167 := Call(__e, tmp10149, tmp10166)


ifres10148 = tmp10167


} else {
tmp10168 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10148 = tmp10168


}

__e.TailApply(tmp10145, ifres10148)
return


}, 1)

tmp10170 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10144)


_ = tmp10170

tmp10171 := MakeNative(func(__e *ControlFlow) {
V1290 := __e.Get(1)
_ = V1290
V1291 := __e.Get(2)
_ = V1291
tmp10215 := PrimIsPair(V1291)

var ifres10196 Obj

if True == tmp10215 {
tmp10213 := PrimHead(V1291)

tmp10214 := PrimIsPair(tmp10213)

var ifres10198 Obj

if True == tmp10214 {
tmp10210 := PrimHead(V1291)

tmp10211 := PrimTail(tmp10210)

tmp10212 := PrimIsPair(tmp10211)

var ifres10200 Obj

if True == tmp10212 {
tmp10206 := PrimHead(V1291)

tmp10207 := PrimTail(tmp10206)

tmp10208 := PrimTail(tmp10207)

tmp10209 := PrimEqual(Nil, tmp10208)

var ifres10202 Obj

if True == tmp10209 {
tmp10204 := PrimTail(V1291)

tmp10205 := PrimEqual(Nil, tmp10204)

var ifres10203 Obj

if True == tmp10205 {
ifres10203 = True


} else {
ifres10203 = False


}

ifres10202 = ifres10203


} else {
ifres10202 = False


}

var ifres10201 Obj

if True == ifres10202 {
ifres10201 = True


} else {
ifres10201 = False


}

ifres10200 = ifres10201


} else {
ifres10200 = False


}

var ifres10199 Obj

if True == ifres10200 {
ifres10199 = True


} else {
ifres10199 = False


}

ifres10198 = ifres10199


} else {
ifres10198 = False


}

var ifres10197 Obj

if True == ifres10198 {
ifres10197 = True


} else {
ifres10197 = False


}

ifres10196 = ifres10197


} else {
ifres10196 = False


}

if True == ifres10196 {
tmp10172 := PrimHead(V1291)

tmp10173 := PrimHead(tmp10172)

__e.TailApply(PrimFunc(symlength), tmp10173)
return


} else {
tmp10194 := PrimIsPair(V1291)

var ifres10179 Obj

if True == tmp10194 {
tmp10192 := PrimHead(V1291)

tmp10193 := PrimIsPair(tmp10192)

var ifres10181 Obj

if True == tmp10193 {
tmp10189 := PrimHead(V1291)

tmp10190 := PrimTail(tmp10189)

tmp10191 := PrimIsPair(tmp10190)

var ifres10183 Obj

if True == tmp10191 {
tmp10185 := PrimHead(V1291)

tmp10186 := PrimTail(tmp10185)

tmp10187 := PrimTail(tmp10186)

tmp10188 := PrimEqual(Nil, tmp10187)

var ifres10184 Obj

if True == tmp10188 {
ifres10184 = True


} else {
ifres10184 = False


}

ifres10183 = ifres10184


} else {
ifres10183 = False


}

var ifres10182 Obj

if True == ifres10183 {
ifres10182 = True


} else {
ifres10182 = False


}

ifres10181 = ifres10182


} else {
ifres10181 = False


}

var ifres10180 Obj

if True == ifres10181 {
ifres10180 = True


} else {
ifres10180 = False


}

ifres10179 = ifres10180


} else {
ifres10179 = False


}

if True == ifres10179 {
tmp10174 := PrimHead(V1291)

tmp10175 := PrimHead(tmp10174)

tmp10176 := Call(__e, PrimFunc(symlength), tmp10175)


tmp10177 := PrimTail(V1291)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1290, tmp10176, tmp10177)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prolog-arity-check")))
return
}


}


}, 2)

tmp10216 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10171)


_ = tmp10216

tmp10217 := MakeNative(func(__e *ControlFlow) {
V1296 := __e.Get(1)
_ = V1296
V1297 := __e.Get(2)
_ = V1297
V1298 := __e.Get(3)
_ = V1298
tmp10233 := PrimEqual(Nil, V1298)

if True == tmp10233 {
__e.Return(V1297)
return
} else {
tmp10231 := PrimIsPair(V1298)

var ifres10227 Obj

if True == tmp10231 {
tmp10229 := PrimHead(V1298)

tmp10230 := PrimIsPair(tmp10229)

var ifres10228 Obj

if True == tmp10230 {
ifres10228 = True


} else {
ifres10228 = False


}

ifres10227 = ifres10228


} else {
ifres10227 = False


}

if True == ifres10227 {
tmp10222 := PrimHead(V1298)

tmp10223 := PrimHead(tmp10222)

tmp10224 := Call(__e, PrimFunc(symlength), tmp10223)


tmp10225 := PrimEqual(V1297, tmp10224)

if True == tmp10225 {
tmp10218 := PrimTail(V1298)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1296, V1297, tmp10218)
return


} else {
tmp10219 := Call(__e, PrimFunc(symshen_4app), V1296, MakeString("\n"), symshen_4a)


tmp10220 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp10219)

__e.Return(PrimSimpleError(tmp10220))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.pac-h")))
return
}


}


}, 3)

tmp10234 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10217)


_ = tmp10234

tmp10235 := MakeNative(func(__e *ControlFlow) {
V1299 := __e.Get(1)
_ = V1299
tmp10236 := MakeNative(func(__e *ControlFlow) {
W1300 := __e.Get(1)
_ = W1300
tmp10255 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1300)


if True == tmp10255 {
tmp10237 := MakeNative(func(__e *ControlFlow) {
W1307 := __e.Get(1)
_ = W1307
tmp10239 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1307)


if True == tmp10239 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1307)
return
}


}, 1)

tmp10240 := MakeNative(func(__e *ControlFlow) {
W1308 := __e.Get(1)
_ = W1308
tmp10251 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1308)


if True == tmp10251 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10241 := MakeNative(func(__e *ControlFlow) {
W1309 := __e.Get(1)
_ = W1309
tmp10242 := MakeNative(func(__e *ControlFlow) {
W1310 := __e.Get(1)
_ = W1310
tmp10247 := Call(__e, PrimFunc(symempty_2), W1309)


var ifres10243 Obj

if True == tmp10247 {
ifres10243 = Nil


} else {
tmp10244 := Call(__e, PrimFunc(symshen_4app), W1309, MakeString("\n ..."), symshen_4r)


tmp10245 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp10244)

tmp10246 := PrimSimpleError(tmp10245)

ifres10243 = tmp10246


}

__e.TailApply(PrimFunc(symshen_4comb), W1310, ifres10243)
return


}, 1)

tmp10248 := Call(__e, PrimFunc(symshen_4in_1_6), W1308)


__e.TailApply(tmp10242, tmp10248)
return


}, 1)

tmp10249 := Call(__e, PrimFunc(symshen_4_5_1out), W1308)


__e.TailApply(tmp10241, tmp10249)
return


}


}, 1)

tmp10252 := Call(__e, PrimFunc(sym_5_b_6), V1299)


tmp10253 := Call(__e, tmp10240, tmp10252)


__e.TailApply(tmp10237, tmp10253)
return


} else {
__e.Return(W1300)
return
}


}, 1)

tmp10256 := MakeNative(func(__e *ControlFlow) {
W1301 := __e.Get(1)
_ = W1301
tmp10271 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1301)


if True == tmp10271 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10257 := MakeNative(func(__e *ControlFlow) {
W1302 := __e.Get(1)
_ = W1302
tmp10258 := MakeNative(func(__e *ControlFlow) {
W1303 := __e.Get(1)
_ = W1303
tmp10259 := MakeNative(func(__e *ControlFlow) {
W1304 := __e.Get(1)
_ = W1304
tmp10266 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1304)


if True == tmp10266 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10260 := MakeNative(func(__e *ControlFlow) {
W1305 := __e.Get(1)
_ = W1305
tmp10261 := MakeNative(func(__e *ControlFlow) {
W1306 := __e.Get(1)
_ = W1306
tmp10262 := PrimCons(W1302, W1305)

__e.TailApply(PrimFunc(symshen_4comb), W1306, tmp10262)
return


}, 1)

tmp10263 := Call(__e, PrimFunc(symshen_4in_1_6), W1304)


__e.TailApply(tmp10261, tmp10263)
return


}, 1)

tmp10264 := Call(__e, PrimFunc(symshen_4_5_1out), W1304)


__e.TailApply(tmp10260, tmp10264)
return


}


}, 1)

tmp10267 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1303)


__e.TailApply(tmp10259, tmp10267)
return


}, 1)

tmp10268 := Call(__e, PrimFunc(symshen_4in_1_6), W1301)


__e.TailApply(tmp10258, tmp10268)
return


}, 1)

tmp10269 := Call(__e, PrimFunc(symshen_4_5_1out), W1301)


__e.TailApply(tmp10257, tmp10269)
return


}


}, 1)

tmp10272 := Call(__e, PrimFunc(symshen_4_5clause_6), V1299)


tmp10273 := Call(__e, tmp10256, tmp10272)


__e.TailApply(tmp10236, tmp10273)
return


}, 1)

tmp10274 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10235)


_ = tmp10274

tmp10275 := MakeNative(func(__e *ControlFlow) {
V1311 := __e.Get(1)
_ = V1311
tmp10291 := PrimIsPair(V1311)

var ifres10282 Obj

if True == tmp10291 {
tmp10289 := PrimTail(V1311)

tmp10290 := PrimIsPair(tmp10289)

var ifres10284 Obj

if True == tmp10290 {
tmp10286 := PrimTail(V1311)

tmp10287 := PrimTail(tmp10286)

tmp10288 := PrimEqual(Nil, tmp10287)

var ifres10285 Obj

if True == tmp10288 {
ifres10285 = True


} else {
ifres10285 = False


}

ifres10284 = ifres10285


} else {
ifres10284 = False


}

var ifres10283 Obj

if True == ifres10284 {
ifres10283 = True


} else {
ifres10283 = False


}

ifres10282 = ifres10283


} else {
ifres10282 = False


}

if True == ifres10282 {
tmp10276 := PrimHead(V1311)

tmp10277 := PrimTail(V1311)

tmp10278 := PrimHead(tmp10277)

tmp10279 := Call(__e, PrimFunc(sym_8p), tmp10276, tmp10278)


tmp10280 := Call(__e, PrimFunc(symshen_4linearise), tmp10279)


__e.TailApply(PrimFunc(symshen_4lch), tmp10280)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.linearise-clause")))
return
}


}, 1)

tmp10292 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10275)


_ = tmp10292

tmp10293 := MakeNative(func(__e *ControlFlow) {
V1312 := __e.Get(1)
_ = V1312
tmp10299 := Call(__e, PrimFunc(symtuple_2), V1312)


if True == tmp10299 {
tmp10294 := Call(__e, PrimFunc(symfst), V1312)


tmp10295 := Call(__e, PrimFunc(symsnd), V1312)


tmp10296 := Call(__e, PrimFunc(symshen_4lchh), tmp10295)


tmp10297 := PrimCons(tmp10296, Nil)

__e.Return(PrimCons(tmp10294, tmp10297))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.lch")))
return
}


}, 1)

tmp10300 := Call(__e, ns2_1set, symshen_4lch, tmp10293)


_ = tmp10300

tmp10301 := MakeNative(func(__e *ControlFlow) {
V1313 := __e.Get(1)
_ = V1313
tmp10364 := PrimIsPair(V1313)

var ifres10313 Obj

if True == tmp10364 {
tmp10362 := PrimHead(V1313)

tmp10363 := PrimEqual(symwhere, tmp10362)

var ifres10315 Obj

if True == tmp10363 {
tmp10360 := PrimTail(V1313)

tmp10361 := PrimIsPair(tmp10360)

var ifres10317 Obj

if True == tmp10361 {
tmp10357 := PrimTail(V1313)

tmp10358 := PrimHead(tmp10357)

tmp10359 := PrimIsPair(tmp10358)

var ifres10319 Obj

if True == tmp10359 {
tmp10353 := PrimTail(V1313)

tmp10354 := PrimHead(tmp10353)

tmp10355 := PrimHead(tmp10354)

tmp10356 := PrimEqual(sym_a, tmp10355)

var ifres10321 Obj

if True == tmp10356 {
tmp10349 := PrimTail(V1313)

tmp10350 := PrimHead(tmp10349)

tmp10351 := PrimTail(tmp10350)

tmp10352 := PrimIsPair(tmp10351)

var ifres10323 Obj

if True == tmp10352 {
tmp10344 := PrimTail(V1313)

tmp10345 := PrimHead(tmp10344)

tmp10346 := PrimTail(tmp10345)

tmp10347 := PrimTail(tmp10346)

tmp10348 := PrimIsPair(tmp10347)

var ifres10325 Obj

if True == tmp10348 {
tmp10338 := PrimTail(V1313)

tmp10339 := PrimHead(tmp10338)

tmp10340 := PrimTail(tmp10339)

tmp10341 := PrimTail(tmp10340)

tmp10342 := PrimTail(tmp10341)

tmp10343 := PrimEqual(Nil, tmp10342)

var ifres10327 Obj

if True == tmp10343 {
tmp10335 := PrimTail(V1313)

tmp10336 := PrimTail(tmp10335)

tmp10337 := PrimIsPair(tmp10336)

var ifres10329 Obj

if True == tmp10337 {
tmp10331 := PrimTail(V1313)

tmp10332 := PrimTail(tmp10331)

tmp10333 := PrimTail(tmp10332)

tmp10334 := PrimEqual(Nil, tmp10333)

var ifres10330 Obj

if True == tmp10334 {
ifres10330 = True


} else {
ifres10330 = False


}

ifres10329 = ifres10330


} else {
ifres10329 = False


}

var ifres10328 Obj

if True == ifres10329 {
ifres10328 = True


} else {
ifres10328 = False


}

ifres10327 = ifres10328


} else {
ifres10327 = False


}

var ifres10326 Obj

if True == ifres10327 {
ifres10326 = True


} else {
ifres10326 = False


}

ifres10325 = ifres10326


} else {
ifres10325 = False


}

var ifres10324 Obj

if True == ifres10325 {
ifres10324 = True


} else {
ifres10324 = False


}

ifres10323 = ifres10324


} else {
ifres10323 = False


}

var ifres10322 Obj

if True == ifres10323 {
ifres10322 = True


} else {
ifres10322 = False


}

ifres10321 = ifres10322


} else {
ifres10321 = False


}

var ifres10320 Obj

if True == ifres10321 {
ifres10320 = True


} else {
ifres10320 = False


}

ifres10319 = ifres10320


} else {
ifres10319 = False


}

var ifres10318 Obj

if True == ifres10319 {
ifres10318 = True


} else {
ifres10318 = False


}

ifres10317 = ifres10318


} else {
ifres10317 = False


}

var ifres10316 Obj

if True == ifres10317 {
ifres10316 = True


} else {
ifres10316 = False


}

ifres10315 = ifres10316


} else {
ifres10315 = False


}

var ifres10314 Obj

if True == ifres10315 {
ifres10314 = True


} else {
ifres10314 = False


}

ifres10313 = ifres10314


} else {
ifres10313 = False


}

if True == ifres10313 {
tmp10303 := PrimValue(symshen_4_doccurs_d)

var ifres10302 Obj

if True == tmp10303 {
ifres10302 = symis_b


} else {
ifres10302 = symis


}

tmp10304 := PrimTail(V1313)

tmp10305 := PrimHead(tmp10304)

tmp10306 := PrimTail(tmp10305)

tmp10307 := PrimCons(ifres10302, tmp10306)

tmp10308 := PrimTail(V1313)

tmp10309 := PrimTail(tmp10308)

tmp10310 := PrimHead(tmp10309)

tmp10311 := Call(__e, PrimFunc(symshen_4lchh), tmp10310)


__e.Return(PrimCons(tmp10307, tmp10311))
return


} else {
__e.Return(V1313)
return
}


}, 1)

tmp10365 := Call(__e, ns2_1set, symshen_4lchh, tmp10301)


_ = tmp10365

tmp10366 := MakeNative(func(__e *ControlFlow) {
V1314 := __e.Get(1)
_ = V1314
tmp10367 := MakeNative(func(__e *ControlFlow) {
W1315 := __e.Get(1)
_ = W1315
tmp10369 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1315)


if True == tmp10369 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1315)
return
}


}, 1)

tmp10370 := MakeNative(func(__e *ControlFlow) {
W1316 := __e.Get(1)
_ = W1316
tmp10396 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1316)


if True == tmp10396 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10371 := MakeNative(func(__e *ControlFlow) {
W1317 := __e.Get(1)
_ = W1317
tmp10372 := MakeNative(func(__e *ControlFlow) {
W1318 := __e.Get(1)
_ = W1318
tmp10392 := Call(__e, PrimFunc(symshen_4hds_a_2), W1318, sym_5_1_1)


if True == tmp10392 {
tmp10373 := MakeNative(func(__e *ControlFlow) {
W1319 := __e.Get(1)
_ = W1319
tmp10374 := MakeNative(func(__e *ControlFlow) {
W1320 := __e.Get(1)
_ = W1320
tmp10388 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1320)


if True == tmp10388 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10375 := MakeNative(func(__e *ControlFlow) {
W1321 := __e.Get(1)
_ = W1321
tmp10376 := MakeNative(func(__e *ControlFlow) {
W1322 := __e.Get(1)
_ = W1322
tmp10377 := MakeNative(func(__e *ControlFlow) {
W1323 := __e.Get(1)
_ = W1323
tmp10383 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1323)


if True == tmp10383 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10378 := MakeNative(func(__e *ControlFlow) {
W1324 := __e.Get(1)
_ = W1324
tmp10379 := PrimCons(W1321, Nil)

tmp10380 := PrimCons(W1317, tmp10379)

__e.TailApply(PrimFunc(symshen_4comb), W1324, tmp10380)
return


}, 1)

tmp10381 := Call(__e, PrimFunc(symshen_4in_1_6), W1323)


__e.TailApply(tmp10378, tmp10381)
return


}


}, 1)

tmp10384 := Call(__e, PrimFunc(symshen_4_5sc_6), W1322)


__e.TailApply(tmp10377, tmp10384)
return


}, 1)

tmp10385 := Call(__e, PrimFunc(symshen_4in_1_6), W1320)


__e.TailApply(tmp10376, tmp10385)
return


}, 1)

tmp10386 := Call(__e, PrimFunc(symshen_4_5_1out), W1320)


__e.TailApply(tmp10375, tmp10386)
return


}


}, 1)

tmp10389 := Call(__e, PrimFunc(symshen_4_5body_6), W1319)


__e.TailApply(tmp10374, tmp10389)
return


}, 1)

tmp10390 := Call(__e, PrimFunc(symtail), W1318)


__e.TailApply(tmp10373, tmp10390)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10393 := Call(__e, PrimFunc(symshen_4in_1_6), W1316)


__e.TailApply(tmp10372, tmp10393)
return


}, 1)

tmp10394 := Call(__e, PrimFunc(symshen_4_5_1out), W1316)


__e.TailApply(tmp10371, tmp10394)
return


}


}, 1)

tmp10397 := Call(__e, PrimFunc(symshen_4_5head_6), V1314)


tmp10398 := Call(__e, tmp10370, tmp10397)


__e.TailApply(tmp10367, tmp10398)
return


}, 1)

tmp10399 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10366)


_ = tmp10399

tmp10400 := MakeNative(func(__e *ControlFlow) {
V1325 := __e.Get(1)
_ = V1325
tmp10401 := MakeNative(func(__e *ControlFlow) {
W1326 := __e.Get(1)
_ = W1326
tmp10413 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1326)


if True == tmp10413 {
tmp10402 := MakeNative(func(__e *ControlFlow) {
W1333 := __e.Get(1)
_ = W1333
tmp10404 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1333)


if True == tmp10404 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1333)
return
}


}, 1)

tmp10405 := MakeNative(func(__e *ControlFlow) {
W1334 := __e.Get(1)
_ = W1334
tmp10409 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1334)


if True == tmp10409 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10406 := MakeNative(func(__e *ControlFlow) {
W1335 := __e.Get(1)
_ = W1335
__e.TailApply(PrimFunc(symshen_4comb), W1335, Nil)
return
}, 1)

tmp10407 := Call(__e, PrimFunc(symshen_4in_1_6), W1334)


__e.TailApply(tmp10406, tmp10407)
return


}


}, 1)

tmp10410 := Call(__e, PrimFunc(sym_5e_6), V1325)


tmp10411 := Call(__e, tmp10405, tmp10410)


__e.TailApply(tmp10402, tmp10411)
return


} else {
__e.Return(W1326)
return
}


}, 1)

tmp10414 := MakeNative(func(__e *ControlFlow) {
W1327 := __e.Get(1)
_ = W1327
tmp10429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1327)


if True == tmp10429 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10415 := MakeNative(func(__e *ControlFlow) {
W1328 := __e.Get(1)
_ = W1328
tmp10416 := MakeNative(func(__e *ControlFlow) {
W1329 := __e.Get(1)
_ = W1329
tmp10417 := MakeNative(func(__e *ControlFlow) {
W1330 := __e.Get(1)
_ = W1330
tmp10424 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1330)


if True == tmp10424 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10418 := MakeNative(func(__e *ControlFlow) {
W1331 := __e.Get(1)
_ = W1331
tmp10419 := MakeNative(func(__e *ControlFlow) {
W1332 := __e.Get(1)
_ = W1332
tmp10420 := PrimCons(W1328, W1331)

__e.TailApply(PrimFunc(symshen_4comb), W1332, tmp10420)
return


}, 1)

tmp10421 := Call(__e, PrimFunc(symshen_4in_1_6), W1330)


__e.TailApply(tmp10419, tmp10421)
return


}, 1)

tmp10422 := Call(__e, PrimFunc(symshen_4_5_1out), W1330)


__e.TailApply(tmp10418, tmp10422)
return


}


}, 1)

tmp10425 := Call(__e, PrimFunc(symshen_4_5head_6), W1329)


__e.TailApply(tmp10417, tmp10425)
return


}, 1)

tmp10426 := Call(__e, PrimFunc(symshen_4in_1_6), W1327)


__e.TailApply(tmp10416, tmp10426)
return


}, 1)

tmp10427 := Call(__e, PrimFunc(symshen_4_5_1out), W1327)


__e.TailApply(tmp10415, tmp10427)
return


}


}, 1)

tmp10430 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1325)


tmp10431 := Call(__e, tmp10414, tmp10430)


__e.TailApply(tmp10401, tmp10431)
return


}, 1)

tmp10432 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10400)


_ = tmp10432

tmp10433 := MakeNative(func(__e *ControlFlow) {
V1336 := __e.Get(1)
_ = V1336
tmp10434 := MakeNative(func(__e *ControlFlow) {
W1337 := __e.Get(1)
_ = W1337
tmp10622 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1337)


if True == tmp10622 {
tmp10435 := MakeNative(func(__e *ControlFlow) {
W1340 := __e.Get(1)
_ = W1340
tmp10609 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1340)


if True == tmp10609 {
tmp10436 := MakeNative(func(__e *ControlFlow) {
W1343 := __e.Get(1)
_ = W1343
tmp10570 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1343)


if True == tmp10570 {
tmp10437 := MakeNative(func(__e *ControlFlow) {
W1355 := __e.Get(1)
_ = W1355
tmp10540 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1355)


if True == tmp10540 {
tmp10438 := MakeNative(func(__e *ControlFlow) {
W1364 := __e.Get(1)
_ = W1364
tmp10510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1364)


if True == tmp10510 {
tmp10439 := MakeNative(func(__e *ControlFlow) {
W1373 := __e.Get(1)
_ = W1373
tmp10476 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1373)


if True == tmp10476 {
tmp10440 := MakeNative(func(__e *ControlFlow) {
W1383 := __e.Get(1)
_ = W1383
tmp10442 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1383)


if True == tmp10442 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1383)
return
}


}, 1)

tmp10474 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10443 Obj

if True == tmp10474 {
tmp10444 := MakeNative(func(__e *ControlFlow) {
W1384 := __e.Get(1)
_ = W1384
tmp10445 := MakeNative(func(__e *ControlFlow) {
W1385 := __e.Get(1)
_ = W1385
tmp10469 := Call(__e, PrimFunc(symshen_4hds_a_2), W1384, symmode)


if True == tmp10469 {
tmp10446 := MakeNative(func(__e *ControlFlow) {
W1386 := __e.Get(1)
_ = W1386
tmp10447 := MakeNative(func(__e *ControlFlow) {
W1387 := __e.Get(1)
_ = W1387
tmp10465 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1387)


if True == tmp10465 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10448 := MakeNative(func(__e *ControlFlow) {
W1388 := __e.Get(1)
_ = W1388
tmp10449 := MakeNative(func(__e *ControlFlow) {
W1389 := __e.Get(1)
_ = W1389
tmp10461 := Call(__e, PrimFunc(symshen_4hds_a_2), W1389, sym_1)


if True == tmp10461 {
tmp10450 := MakeNative(func(__e *ControlFlow) {
W1390 := __e.Get(1)
_ = W1390
tmp10451 := MakeNative(func(__e *ControlFlow) {
W1391 := __e.Get(1)
_ = W1391
tmp10457 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1391)


if True == tmp10457 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10452 := MakeNative(func(__e *ControlFlow) {
W1392 := __e.Get(1)
_ = W1392
tmp10453 := PrimCons(W1388, Nil)

tmp10454 := PrimCons(symshen_4_1m, tmp10453)

__e.TailApply(PrimFunc(symshen_4comb), W1385, tmp10454)
return


}, 1)

tmp10455 := Call(__e, PrimFunc(symshen_4in_1_6), W1391)


__e.TailApply(tmp10452, tmp10455)
return


}


}, 1)

tmp10458 := Call(__e, PrimFunc(sym_5end_6), W1390)


__e.TailApply(tmp10451, tmp10458)
return


}, 1)

tmp10459 := Call(__e, PrimFunc(symtail), W1389)


__e.TailApply(tmp10450, tmp10459)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10462 := Call(__e, PrimFunc(symshen_4in_1_6), W1387)


__e.TailApply(tmp10449, tmp10462)
return


}, 1)

tmp10463 := Call(__e, PrimFunc(symshen_4_5_1out), W1387)


__e.TailApply(tmp10448, tmp10463)
return


}


}, 1)

tmp10466 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1386)


__e.TailApply(tmp10447, tmp10466)
return


}, 1)

tmp10467 := Call(__e, PrimFunc(symtail), W1384)


__e.TailApply(tmp10446, tmp10467)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10470 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10445, tmp10470)
return


}, 1)

tmp10471 := Call(__e, PrimFunc(symhead), V1336)


tmp10472 := Call(__e, tmp10444, tmp10471)


ifres10443 = tmp10472


} else {
tmp10473 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10443 = tmp10473


}

__e.TailApply(tmp10440, ifres10443)
return


} else {
__e.Return(W1373)
return
}


}, 1)

tmp10508 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10477 Obj

if True == tmp10508 {
tmp10478 := MakeNative(func(__e *ControlFlow) {
W1374 := __e.Get(1)
_ = W1374
tmp10479 := MakeNative(func(__e *ControlFlow) {
W1375 := __e.Get(1)
_ = W1375
tmp10503 := Call(__e, PrimFunc(symshen_4hds_a_2), W1374, symmode)


if True == tmp10503 {
tmp10480 := MakeNative(func(__e *ControlFlow) {
W1376 := __e.Get(1)
_ = W1376
tmp10481 := MakeNative(func(__e *ControlFlow) {
W1377 := __e.Get(1)
_ = W1377
tmp10499 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1377)


if True == tmp10499 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10482 := MakeNative(func(__e *ControlFlow) {
W1378 := __e.Get(1)
_ = W1378
tmp10483 := MakeNative(func(__e *ControlFlow) {
W1379 := __e.Get(1)
_ = W1379
tmp10495 := Call(__e, PrimFunc(symshen_4hds_a_2), W1379, sym_7)


if True == tmp10495 {
tmp10484 := MakeNative(func(__e *ControlFlow) {
W1380 := __e.Get(1)
_ = W1380
tmp10485 := MakeNative(func(__e *ControlFlow) {
W1381 := __e.Get(1)
_ = W1381
tmp10491 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1381)


if True == tmp10491 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10486 := MakeNative(func(__e *ControlFlow) {
W1382 := __e.Get(1)
_ = W1382
tmp10487 := PrimCons(W1378, Nil)

tmp10488 := PrimCons(symshen_4_7m, tmp10487)

__e.TailApply(PrimFunc(symshen_4comb), W1375, tmp10488)
return


}, 1)

tmp10489 := Call(__e, PrimFunc(symshen_4in_1_6), W1381)


__e.TailApply(tmp10486, tmp10489)
return


}


}, 1)

tmp10492 := Call(__e, PrimFunc(sym_5end_6), W1380)


__e.TailApply(tmp10485, tmp10492)
return


}, 1)

tmp10493 := Call(__e, PrimFunc(symtail), W1379)


__e.TailApply(tmp10484, tmp10493)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10496 := Call(__e, PrimFunc(symshen_4in_1_6), W1377)


__e.TailApply(tmp10483, tmp10496)
return


}, 1)

tmp10497 := Call(__e, PrimFunc(symshen_4_5_1out), W1377)


__e.TailApply(tmp10482, tmp10497)
return


}


}, 1)

tmp10500 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1376)


__e.TailApply(tmp10481, tmp10500)
return


}, 1)

tmp10501 := Call(__e, PrimFunc(symtail), W1374)


__e.TailApply(tmp10480, tmp10501)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10504 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10479, tmp10504)
return


}, 1)

tmp10505 := Call(__e, PrimFunc(symhead), V1336)


tmp10506 := Call(__e, tmp10478, tmp10505)


ifres10477 = tmp10506


} else {
tmp10507 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10477 = tmp10507


}

__e.TailApply(tmp10439, ifres10477)
return


} else {
__e.Return(W1364)
return
}


}, 1)

tmp10538 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10511 Obj

if True == tmp10538 {
tmp10512 := MakeNative(func(__e *ControlFlow) {
W1365 := __e.Get(1)
_ = W1365
tmp10513 := MakeNative(func(__e *ControlFlow) {
W1366 := __e.Get(1)
_ = W1366
tmp10533 := Call(__e, PrimFunc(symshen_4hds_a_2), W1365, sym_1)


if True == tmp10533 {
tmp10514 := MakeNative(func(__e *ControlFlow) {
W1367 := __e.Get(1)
_ = W1367
tmp10515 := MakeNative(func(__e *ControlFlow) {
W1368 := __e.Get(1)
_ = W1368
tmp10529 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1368)


if True == tmp10529 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10516 := MakeNative(func(__e *ControlFlow) {
W1369 := __e.Get(1)
_ = W1369
tmp10517 := MakeNative(func(__e *ControlFlow) {
W1370 := __e.Get(1)
_ = W1370
tmp10518 := MakeNative(func(__e *ControlFlow) {
W1371 := __e.Get(1)
_ = W1371
tmp10524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1371)


if True == tmp10524 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10519 := MakeNative(func(__e *ControlFlow) {
W1372 := __e.Get(1)
_ = W1372
tmp10520 := PrimCons(W1369, Nil)

tmp10521 := PrimCons(symshen_4_1m, tmp10520)

__e.TailApply(PrimFunc(symshen_4comb), W1366, tmp10521)
return


}, 1)

tmp10522 := Call(__e, PrimFunc(symshen_4in_1_6), W1371)


__e.TailApply(tmp10519, tmp10522)
return


}


}, 1)

tmp10525 := Call(__e, PrimFunc(sym_5end_6), W1370)


__e.TailApply(tmp10518, tmp10525)
return


}, 1)

tmp10526 := Call(__e, PrimFunc(symshen_4in_1_6), W1368)


__e.TailApply(tmp10517, tmp10526)
return


}, 1)

tmp10527 := Call(__e, PrimFunc(symshen_4_5_1out), W1368)


__e.TailApply(tmp10516, tmp10527)
return


}


}, 1)

tmp10530 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1367)


__e.TailApply(tmp10515, tmp10530)
return


}, 1)

tmp10531 := Call(__e, PrimFunc(symtail), W1365)


__e.TailApply(tmp10514, tmp10531)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10534 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10513, tmp10534)
return


}, 1)

tmp10535 := Call(__e, PrimFunc(symhead), V1336)


tmp10536 := Call(__e, tmp10512, tmp10535)


ifres10511 = tmp10536


} else {
tmp10537 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10511 = tmp10537


}

__e.TailApply(tmp10438, ifres10511)
return


} else {
__e.Return(W1355)
return
}


}, 1)

tmp10568 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10541 Obj

if True == tmp10568 {
tmp10542 := MakeNative(func(__e *ControlFlow) {
W1356 := __e.Get(1)
_ = W1356
tmp10543 := MakeNative(func(__e *ControlFlow) {
W1357 := __e.Get(1)
_ = W1357
tmp10563 := Call(__e, PrimFunc(symshen_4hds_a_2), W1356, sym_7)


if True == tmp10563 {
tmp10544 := MakeNative(func(__e *ControlFlow) {
W1358 := __e.Get(1)
_ = W1358
tmp10545 := MakeNative(func(__e *ControlFlow) {
W1359 := __e.Get(1)
_ = W1359
tmp10559 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1359)


if True == tmp10559 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10546 := MakeNative(func(__e *ControlFlow) {
W1360 := __e.Get(1)
_ = W1360
tmp10547 := MakeNative(func(__e *ControlFlow) {
W1361 := __e.Get(1)
_ = W1361
tmp10548 := MakeNative(func(__e *ControlFlow) {
W1362 := __e.Get(1)
_ = W1362
tmp10554 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1362)


if True == tmp10554 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10549 := MakeNative(func(__e *ControlFlow) {
W1363 := __e.Get(1)
_ = W1363
tmp10550 := PrimCons(W1360, Nil)

tmp10551 := PrimCons(symshen_4_7m, tmp10550)

__e.TailApply(PrimFunc(symshen_4comb), W1357, tmp10551)
return


}, 1)

tmp10552 := Call(__e, PrimFunc(symshen_4in_1_6), W1362)


__e.TailApply(tmp10549, tmp10552)
return


}


}, 1)

tmp10555 := Call(__e, PrimFunc(sym_5end_6), W1361)


__e.TailApply(tmp10548, tmp10555)
return


}, 1)

tmp10556 := Call(__e, PrimFunc(symshen_4in_1_6), W1359)


__e.TailApply(tmp10547, tmp10556)
return


}, 1)

tmp10557 := Call(__e, PrimFunc(symshen_4_5_1out), W1359)


__e.TailApply(tmp10546, tmp10557)
return


}


}, 1)

tmp10560 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1358)


__e.TailApply(tmp10545, tmp10560)
return


}, 1)

tmp10561 := Call(__e, PrimFunc(symtail), W1356)


__e.TailApply(tmp10544, tmp10561)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10564 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10543, tmp10564)
return


}, 1)

tmp10565 := Call(__e, PrimFunc(symhead), V1336)


tmp10566 := Call(__e, tmp10542, tmp10565)


ifres10541 = tmp10566


} else {
tmp10567 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10541 = tmp10567


}

__e.TailApply(tmp10437, ifres10541)
return


} else {
__e.Return(W1343)
return
}


}, 1)

tmp10607 := Call(__e, PrimFunc(symshen_4ccons_2), V1336)


var ifres10571 Obj

if True == tmp10607 {
tmp10572 := MakeNative(func(__e *ControlFlow) {
W1344 := __e.Get(1)
_ = W1344
tmp10573 := MakeNative(func(__e *ControlFlow) {
W1345 := __e.Get(1)
_ = W1345
tmp10602 := Call(__e, PrimFunc(symshen_4hds_a_2), W1344, symcons)


if True == tmp10602 {
tmp10574 := MakeNative(func(__e *ControlFlow) {
W1346 := __e.Get(1)
_ = W1346
tmp10575 := MakeNative(func(__e *ControlFlow) {
W1347 := __e.Get(1)
_ = W1347
tmp10598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1347)


if True == tmp10598 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10576 := MakeNative(func(__e *ControlFlow) {
W1348 := __e.Get(1)
_ = W1348
tmp10577 := MakeNative(func(__e *ControlFlow) {
W1349 := __e.Get(1)
_ = W1349
tmp10578 := MakeNative(func(__e *ControlFlow) {
W1350 := __e.Get(1)
_ = W1350
tmp10593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1350)


if True == tmp10593 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10579 := MakeNative(func(__e *ControlFlow) {
W1351 := __e.Get(1)
_ = W1351
tmp10580 := MakeNative(func(__e *ControlFlow) {
W1352 := __e.Get(1)
_ = W1352
tmp10581 := MakeNative(func(__e *ControlFlow) {
W1353 := __e.Get(1)
_ = W1353
tmp10588 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1353)


if True == tmp10588 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10582 := MakeNative(func(__e *ControlFlow) {
W1354 := __e.Get(1)
_ = W1354
tmp10583 := PrimCons(W1351, Nil)

tmp10584 := PrimCons(W1348, tmp10583)

tmp10585 := PrimCons(symcons, tmp10584)

__e.TailApply(PrimFunc(symshen_4comb), W1345, tmp10585)
return


}, 1)

tmp10586 := Call(__e, PrimFunc(symshen_4in_1_6), W1353)


__e.TailApply(tmp10582, tmp10586)
return


}


}, 1)

tmp10589 := Call(__e, PrimFunc(sym_5end_6), W1352)


__e.TailApply(tmp10581, tmp10589)
return


}, 1)

tmp10590 := Call(__e, PrimFunc(symshen_4in_1_6), W1350)


__e.TailApply(tmp10580, tmp10590)
return


}, 1)

tmp10591 := Call(__e, PrimFunc(symshen_4_5_1out), W1350)


__e.TailApply(tmp10579, tmp10591)
return


}


}, 1)

tmp10594 := Call(__e, PrimFunc(symshen_4_5hterm2_6), W1349)


__e.TailApply(tmp10578, tmp10594)
return


}, 1)

tmp10595 := Call(__e, PrimFunc(symshen_4in_1_6), W1347)


__e.TailApply(tmp10577, tmp10595)
return


}, 1)

tmp10596 := Call(__e, PrimFunc(symshen_4_5_1out), W1347)


__e.TailApply(tmp10576, tmp10596)
return


}


}, 1)

tmp10599 := Call(__e, PrimFunc(symshen_4_5hterm1_6), W1346)


__e.TailApply(tmp10575, tmp10599)
return


}, 1)

tmp10600 := Call(__e, PrimFunc(symtail), W1344)


__e.TailApply(tmp10574, tmp10600)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10603 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10573, tmp10603)
return


}, 1)

tmp10604 := Call(__e, PrimFunc(symhead), V1336)


tmp10605 := Call(__e, tmp10572, tmp10604)


ifres10571 = tmp10605


} else {
tmp10606 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10571 = tmp10606


}

__e.TailApply(tmp10436, ifres10571)
return


} else {
__e.Return(W1340)
return
}


}, 1)

tmp10620 := PrimIsPair(V1336)

var ifres10610 Obj

if True == tmp10620 {
tmp10611 := MakeNative(func(__e *ControlFlow) {
W1341 := __e.Get(1)
_ = W1341
tmp10612 := MakeNative(func(__e *ControlFlow) {
W1342 := __e.Get(1)
_ = W1342
tmp10614 := PrimIntern(MakeString(":"))

tmp10615 := PrimEqual(W1341, tmp10614)

if True == tmp10615 {
__e.TailApply(PrimFunc(symshen_4comb), W1342, W1341)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10616 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10612, tmp10616)
return


}, 1)

tmp10617 := Call(__e, PrimFunc(symhead), V1336)


tmp10618 := Call(__e, tmp10611, tmp10617)


ifres10610 = tmp10618


} else {
tmp10619 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10610 = tmp10619


}

__e.TailApply(tmp10435, ifres10610)
return


} else {
__e.Return(W1337)
return
}


}, 1)

tmp10636 := PrimIsPair(V1336)

var ifres10623 Obj

if True == tmp10636 {
tmp10624 := MakeNative(func(__e *ControlFlow) {
W1338 := __e.Get(1)
_ = W1338
tmp10625 := MakeNative(func(__e *ControlFlow) {
W1339 := __e.Get(1)
_ = W1339
tmp10631 := Call(__e, PrimFunc(symatom_2), W1338)


var ifres10627 Obj

if True == tmp10631 {
tmp10629 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), W1338)


tmp10630 := PrimNot(tmp10629)

var ifres10628 Obj

if True == tmp10630 {
ifres10628 = True


} else {
ifres10628 = False


}

ifres10627 = ifres10628


} else {
ifres10627 = False


}

if True == ifres10627 {
__e.TailApply(PrimFunc(symshen_4comb), W1339, W1338)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10632 := Call(__e, PrimFunc(symtail), V1336)


__e.TailApply(tmp10625, tmp10632)
return


}, 1)

tmp10633 := Call(__e, PrimFunc(symhead), V1336)


tmp10634 := Call(__e, tmp10624, tmp10633)


ifres10623 = tmp10634


} else {
tmp10635 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10623 = tmp10635


}

__e.TailApply(tmp10434, ifres10623)
return


}, 1)

tmp10637 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10433)


_ = tmp10637

tmp10638 := MakeNative(func(__e *ControlFlow) {
V1393 := __e.Get(1)
_ = V1393
tmp10639 := PrimIntern(MakeString(";"))

tmp10640 := PrimCons(sym_5_1_1, Nil)

tmp10641 := PrimCons(tmp10639, tmp10640)

__e.TailApply(PrimFunc(symelement_2), V1393, tmp10641)
return


}, 1)

tmp10642 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10638)


_ = tmp10642

tmp10643 := MakeNative(func(__e *ControlFlow) {
V1394 := __e.Get(1)
_ = V1394
tmp10656 := PrimIsSymbol(V1394)

if True == tmp10656 {
__e.Return(True)
return
} else {
tmp10654 := PrimIsString(V1394)

var ifres10645 Obj

if True == tmp10654 {
ifres10645 = True


} else {
tmp10653 := Call(__e, PrimFunc(symboolean_2), V1394)


var ifres10647 Obj

if True == tmp10653 {
ifres10647 = True


} else {
tmp10652 := PrimIsNumber(V1394)

var ifres10649 Obj

if True == tmp10652 {
ifres10649 = True


} else {
tmp10651 := Call(__e, PrimFunc(symempty_2), V1394)


var ifres10650 Obj

if True == tmp10651 {
ifres10650 = True


} else {
ifres10650 = False


}

ifres10649 = ifres10650


}

var ifres10648 Obj

if True == ifres10649 {
ifres10648 = True


} else {
ifres10648 = False


}

ifres10647 = ifres10648


}

var ifres10646 Obj

if True == ifres10647 {
ifres10646 = True


} else {
ifres10646 = False


}

ifres10645 = ifres10646


}

if True == ifres10645 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10657 := Call(__e, ns2_1set, symatom_2, tmp10643)


_ = tmp10657

tmp10658 := MakeNative(func(__e *ControlFlow) {
V1395 := __e.Get(1)
_ = V1395
tmp10659 := MakeNative(func(__e *ControlFlow) {
W1396 := __e.Get(1)
_ = W1396
tmp10661 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1396)


if True == tmp10661 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1396)
return
}


}, 1)

tmp10662 := MakeNative(func(__e *ControlFlow) {
W1397 := __e.Get(1)
_ = W1397
tmp10668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1397)


if True == tmp10668 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10663 := MakeNative(func(__e *ControlFlow) {
W1398 := __e.Get(1)
_ = W1398
tmp10664 := MakeNative(func(__e *ControlFlow) {
W1399 := __e.Get(1)
_ = W1399
__e.TailApply(PrimFunc(symshen_4comb), W1399, W1398)
return
}, 1)

tmp10665 := Call(__e, PrimFunc(symshen_4in_1_6), W1397)


__e.TailApply(tmp10664, tmp10665)
return


}, 1)

tmp10666 := Call(__e, PrimFunc(symshen_4_5_1out), W1397)


__e.TailApply(tmp10663, tmp10666)
return


}


}, 1)

tmp10669 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1395)


tmp10670 := Call(__e, tmp10662, tmp10669)


__e.TailApply(tmp10659, tmp10670)
return


}, 1)

tmp10671 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10658)


_ = tmp10671

tmp10672 := MakeNative(func(__e *ControlFlow) {
V1400 := __e.Get(1)
_ = V1400
tmp10673 := MakeNative(func(__e *ControlFlow) {
W1401 := __e.Get(1)
_ = W1401
tmp10675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1401)


if True == tmp10675 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1401)
return
}


}, 1)

tmp10676 := MakeNative(func(__e *ControlFlow) {
W1402 := __e.Get(1)
_ = W1402
tmp10682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1402)


if True == tmp10682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10677 := MakeNative(func(__e *ControlFlow) {
W1403 := __e.Get(1)
_ = W1403
tmp10678 := MakeNative(func(__e *ControlFlow) {
W1404 := __e.Get(1)
_ = W1404
__e.TailApply(PrimFunc(symshen_4comb), W1404, W1403)
return
}, 1)

tmp10679 := Call(__e, PrimFunc(symshen_4in_1_6), W1402)


__e.TailApply(tmp10678, tmp10679)
return


}, 1)

tmp10680 := Call(__e, PrimFunc(symshen_4_5_1out), W1402)


__e.TailApply(tmp10677, tmp10680)
return


}


}, 1)

tmp10683 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1400)


tmp10684 := Call(__e, tmp10676, tmp10683)


__e.TailApply(tmp10673, tmp10684)
return


}, 1)

tmp10685 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10672)


_ = tmp10685

tmp10686 := MakeNative(func(__e *ControlFlow) {
V1405 := __e.Get(1)
_ = V1405
tmp10687 := MakeNative(func(__e *ControlFlow) {
W1406 := __e.Get(1)
_ = W1406
tmp10699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1406)


if True == tmp10699 {
tmp10688 := MakeNative(func(__e *ControlFlow) {
W1413 := __e.Get(1)
_ = W1413
tmp10690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1413)


if True == tmp10690 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1413)
return
}


}, 1)

tmp10691 := MakeNative(func(__e *ControlFlow) {
W1414 := __e.Get(1)
_ = W1414
tmp10695 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1414)


if True == tmp10695 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10692 := MakeNative(func(__e *ControlFlow) {
W1415 := __e.Get(1)
_ = W1415
__e.TailApply(PrimFunc(symshen_4comb), W1415, Nil)
return
}, 1)

tmp10693 := Call(__e, PrimFunc(symshen_4in_1_6), W1414)


__e.TailApply(tmp10692, tmp10693)
return


}


}, 1)

tmp10696 := Call(__e, PrimFunc(sym_5e_6), V1405)


tmp10697 := Call(__e, tmp10691, tmp10696)


__e.TailApply(tmp10688, tmp10697)
return


} else {
__e.Return(W1406)
return
}


}, 1)

tmp10700 := MakeNative(func(__e *ControlFlow) {
W1407 := __e.Get(1)
_ = W1407
tmp10715 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1407)


if True == tmp10715 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10701 := MakeNative(func(__e *ControlFlow) {
W1408 := __e.Get(1)
_ = W1408
tmp10702 := MakeNative(func(__e *ControlFlow) {
W1409 := __e.Get(1)
_ = W1409
tmp10703 := MakeNative(func(__e *ControlFlow) {
W1410 := __e.Get(1)
_ = W1410
tmp10710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1410)


if True == tmp10710 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10704 := MakeNative(func(__e *ControlFlow) {
W1411 := __e.Get(1)
_ = W1411
tmp10705 := MakeNative(func(__e *ControlFlow) {
W1412 := __e.Get(1)
_ = W1412
tmp10706 := PrimCons(W1408, W1411)

__e.TailApply(PrimFunc(symshen_4comb), W1412, tmp10706)
return


}, 1)

tmp10707 := Call(__e, PrimFunc(symshen_4in_1_6), W1410)


__e.TailApply(tmp10705, tmp10707)
return


}, 1)

tmp10708 := Call(__e, PrimFunc(symshen_4_5_1out), W1410)


__e.TailApply(tmp10704, tmp10708)
return


}


}, 1)

tmp10711 := Call(__e, PrimFunc(symshen_4_5body_6), W1409)


__e.TailApply(tmp10703, tmp10711)
return


}, 1)

tmp10712 := Call(__e, PrimFunc(symshen_4in_1_6), W1407)


__e.TailApply(tmp10702, tmp10712)
return


}, 1)

tmp10713 := Call(__e, PrimFunc(symshen_4_5_1out), W1407)


__e.TailApply(tmp10701, tmp10713)
return


}


}, 1)

tmp10716 := Call(__e, PrimFunc(symshen_4_5literal_6), V1405)


tmp10717 := Call(__e, tmp10700, tmp10716)


__e.TailApply(tmp10687, tmp10717)
return


}, 1)

tmp10718 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10686)


_ = tmp10718

tmp10719 := MakeNative(func(__e *ControlFlow) {
V1416 := __e.Get(1)
_ = V1416
tmp10720 := MakeNative(func(__e *ControlFlow) {
W1417 := __e.Get(1)
_ = W1417
tmp10747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1417)


if True == tmp10747 {
tmp10721 := MakeNative(func(__e *ControlFlow) {
W1419 := __e.Get(1)
_ = W1419
tmp10723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1419)


if True == tmp10723 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1419)
return
}


}, 1)

tmp10745 := Call(__e, PrimFunc(symshen_4ccons_2), V1416)


var ifres10724 Obj

if True == tmp10745 {
tmp10725 := MakeNative(func(__e *ControlFlow) {
W1420 := __e.Get(1)
_ = W1420
tmp10726 := MakeNative(func(__e *ControlFlow) {
W1421 := __e.Get(1)
_ = W1421
tmp10727 := MakeNative(func(__e *ControlFlow) {
W1422 := __e.Get(1)
_ = W1422
tmp10739 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1422)


if True == tmp10739 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10728 := MakeNative(func(__e *ControlFlow) {
W1423 := __e.Get(1)
_ = W1423
tmp10729 := MakeNative(func(__e *ControlFlow) {
W1424 := __e.Get(1)
_ = W1424
tmp10730 := MakeNative(func(__e *ControlFlow) {
W1425 := __e.Get(1)
_ = W1425
tmp10734 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1425)


if True == tmp10734 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10731 := MakeNative(func(__e *ControlFlow) {
W1426 := __e.Get(1)
_ = W1426
__e.TailApply(PrimFunc(symshen_4comb), W1421, W1423)
return
}, 1)

tmp10732 := Call(__e, PrimFunc(symshen_4in_1_6), W1425)


__e.TailApply(tmp10731, tmp10732)
return


}


}, 1)

tmp10735 := Call(__e, PrimFunc(sym_5end_6), W1424)


__e.TailApply(tmp10730, tmp10735)
return


}, 1)

tmp10736 := Call(__e, PrimFunc(symshen_4in_1_6), W1422)


__e.TailApply(tmp10729, tmp10736)
return


}, 1)

tmp10737 := Call(__e, PrimFunc(symshen_4_5_1out), W1422)


__e.TailApply(tmp10728, tmp10737)
return


}


}, 1)

tmp10740 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1420)


__e.TailApply(tmp10727, tmp10740)
return


}, 1)

tmp10741 := Call(__e, PrimFunc(symtail), V1416)


__e.TailApply(tmp10726, tmp10741)
return


}, 1)

tmp10742 := Call(__e, PrimFunc(symhead), V1416)


tmp10743 := Call(__e, tmp10725, tmp10742)


ifres10724 = tmp10743


} else {
tmp10744 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10724 = tmp10744


}

__e.TailApply(tmp10721, ifres10724)
return


} else {
__e.Return(W1417)
return
}


}, 1)

tmp10753 := Call(__e, PrimFunc(symshen_4hds_a_2), V1416, sym_b)


var ifres10748 Obj

if True == tmp10753 {
tmp10749 := MakeNative(func(__e *ControlFlow) {
W1418 := __e.Get(1)
_ = W1418
__e.TailApply(PrimFunc(symshen_4comb), W1418, sym_b)
return
}, 1)

tmp10750 := Call(__e, PrimFunc(symtail), V1416)


tmp10751 := Call(__e, tmp10749, tmp10750)


ifres10748 = tmp10751


} else {
tmp10752 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10748 = tmp10752


}

__e.TailApply(tmp10720, ifres10748)
return


}, 1)

tmp10754 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10719)


_ = tmp10754

tmp10755 := MakeNative(func(__e *ControlFlow) {
V1427 := __e.Get(1)
_ = V1427
tmp10756 := MakeNative(func(__e *ControlFlow) {
W1428 := __e.Get(1)
_ = W1428
tmp10768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1428)


if True == tmp10768 {
tmp10757 := MakeNative(func(__e *ControlFlow) {
W1435 := __e.Get(1)
_ = W1435
tmp10759 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1435)


if True == tmp10759 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1435)
return
}


}, 1)

tmp10760 := MakeNative(func(__e *ControlFlow) {
W1436 := __e.Get(1)
_ = W1436
tmp10764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1436)


if True == tmp10764 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10761 := MakeNative(func(__e *ControlFlow) {
W1437 := __e.Get(1)
_ = W1437
__e.TailApply(PrimFunc(symshen_4comb), W1437, Nil)
return
}, 1)

tmp10762 := Call(__e, PrimFunc(symshen_4in_1_6), W1436)


__e.TailApply(tmp10761, tmp10762)
return


}


}, 1)

tmp10765 := Call(__e, PrimFunc(sym_5e_6), V1427)


tmp10766 := Call(__e, tmp10760, tmp10765)


__e.TailApply(tmp10757, tmp10766)
return


} else {
__e.Return(W1428)
return
}


}, 1)

tmp10769 := MakeNative(func(__e *ControlFlow) {
W1429 := __e.Get(1)
_ = W1429
tmp10784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1429)


if True == tmp10784 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10770 := MakeNative(func(__e *ControlFlow) {
W1430 := __e.Get(1)
_ = W1430
tmp10771 := MakeNative(func(__e *ControlFlow) {
W1431 := __e.Get(1)
_ = W1431
tmp10772 := MakeNative(func(__e *ControlFlow) {
W1432 := __e.Get(1)
_ = W1432
tmp10779 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1432)


if True == tmp10779 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10773 := MakeNative(func(__e *ControlFlow) {
W1433 := __e.Get(1)
_ = W1433
tmp10774 := MakeNative(func(__e *ControlFlow) {
W1434 := __e.Get(1)
_ = W1434
tmp10775 := PrimCons(W1430, W1433)

__e.TailApply(PrimFunc(symshen_4comb), W1434, tmp10775)
return


}, 1)

tmp10776 := Call(__e, PrimFunc(symshen_4in_1_6), W1432)


__e.TailApply(tmp10774, tmp10776)
return


}, 1)

tmp10777 := Call(__e, PrimFunc(symshen_4_5_1out), W1432)


__e.TailApply(tmp10773, tmp10777)
return


}


}, 1)

tmp10780 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1431)


__e.TailApply(tmp10772, tmp10780)
return


}, 1)

tmp10781 := Call(__e, PrimFunc(symshen_4in_1_6), W1429)


__e.TailApply(tmp10771, tmp10781)
return


}, 1)

tmp10782 := Call(__e, PrimFunc(symshen_4_5_1out), W1429)


__e.TailApply(tmp10770, tmp10782)
return


}


}, 1)

tmp10785 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1427)


tmp10786 := Call(__e, tmp10769, tmp10785)


__e.TailApply(tmp10756, tmp10786)
return


}, 1)

tmp10787 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10755)


_ = tmp10787

tmp10788 := MakeNative(func(__e *ControlFlow) {
V1438 := __e.Get(1)
_ = V1438
tmp10789 := MakeNative(func(__e *ControlFlow) {
W1439 := __e.Get(1)
_ = W1439
tmp10829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1439)


if True == tmp10829 {
tmp10790 := MakeNative(func(__e *ControlFlow) {
W1443 := __e.Get(1)
_ = W1443
tmp10817 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1443)


if True == tmp10817 {
tmp10791 := MakeNative(func(__e *ControlFlow) {
W1446 := __e.Get(1)
_ = W1446
tmp10793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1446)


if True == tmp10793 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1446)
return
}


}, 1)

tmp10815 := Call(__e, PrimFunc(symshen_4ccons_2), V1438)


var ifres10794 Obj

if True == tmp10815 {
tmp10795 := MakeNative(func(__e *ControlFlow) {
W1447 := __e.Get(1)
_ = W1447
tmp10796 := MakeNative(func(__e *ControlFlow) {
W1448 := __e.Get(1)
_ = W1448
tmp10797 := MakeNative(func(__e *ControlFlow) {
W1449 := __e.Get(1)
_ = W1449
tmp10809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1449)


if True == tmp10809 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10798 := MakeNative(func(__e *ControlFlow) {
W1450 := __e.Get(1)
_ = W1450
tmp10799 := MakeNative(func(__e *ControlFlow) {
W1451 := __e.Get(1)
_ = W1451
tmp10800 := MakeNative(func(__e *ControlFlow) {
W1452 := __e.Get(1)
_ = W1452
tmp10804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1452)


if True == tmp10804 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10801 := MakeNative(func(__e *ControlFlow) {
W1453 := __e.Get(1)
_ = W1453
__e.TailApply(PrimFunc(symshen_4comb), W1448, W1450)
return
}, 1)

tmp10802 := Call(__e, PrimFunc(symshen_4in_1_6), W1452)


__e.TailApply(tmp10801, tmp10802)
return


}


}, 1)

tmp10805 := Call(__e, PrimFunc(sym_5end_6), W1451)


__e.TailApply(tmp10800, tmp10805)
return


}, 1)

tmp10806 := Call(__e, PrimFunc(symshen_4in_1_6), W1449)


__e.TailApply(tmp10799, tmp10806)
return


}, 1)

tmp10807 := Call(__e, PrimFunc(symshen_4_5_1out), W1449)


__e.TailApply(tmp10798, tmp10807)
return


}


}, 1)

tmp10810 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1447)


__e.TailApply(tmp10797, tmp10810)
return


}, 1)

tmp10811 := Call(__e, PrimFunc(symtail), V1438)


__e.TailApply(tmp10796, tmp10811)
return


}, 1)

tmp10812 := Call(__e, PrimFunc(symhead), V1438)


tmp10813 := Call(__e, tmp10795, tmp10812)


ifres10794 = tmp10813


} else {
tmp10814 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10794 = tmp10814


}

__e.TailApply(tmp10791, ifres10794)
return


} else {
__e.Return(W1443)
return
}


}, 1)

tmp10827 := PrimIsPair(V1438)

var ifres10818 Obj

if True == tmp10827 {
tmp10819 := MakeNative(func(__e *ControlFlow) {
W1444 := __e.Get(1)
_ = W1444
tmp10820 := MakeNative(func(__e *ControlFlow) {
W1445 := __e.Get(1)
_ = W1445
tmp10822 := Call(__e, PrimFunc(symatom_2), W1444)


if True == tmp10822 {
__e.TailApply(PrimFunc(symshen_4comb), W1445, W1444)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10823 := Call(__e, PrimFunc(symtail), V1438)


__e.TailApply(tmp10820, tmp10823)
return


}, 1)

tmp10824 := Call(__e, PrimFunc(symhead), V1438)


tmp10825 := Call(__e, tmp10819, tmp10824)


ifres10818 = tmp10825


} else {
tmp10826 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10818 = tmp10826


}

__e.TailApply(tmp10790, ifres10818)
return


} else {
__e.Return(W1439)
return
}


}, 1)

tmp10830 := MakeNative(func(__e *ControlFlow) {
W1440 := __e.Get(1)
_ = W1440
tmp10836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1440)


if True == tmp10836 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10831 := MakeNative(func(__e *ControlFlow) {
W1441 := __e.Get(1)
_ = W1441
tmp10832 := MakeNative(func(__e *ControlFlow) {
W1442 := __e.Get(1)
_ = W1442
__e.TailApply(PrimFunc(symshen_4comb), W1442, W1441)
return
}, 1)

tmp10833 := Call(__e, PrimFunc(symshen_4in_1_6), W1440)


__e.TailApply(tmp10832, tmp10833)
return


}, 1)

tmp10834 := Call(__e, PrimFunc(symshen_4_5_1out), W1440)


__e.TailApply(tmp10831, tmp10834)
return


}


}, 1)

tmp10837 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1438)


tmp10838 := Call(__e, tmp10830, tmp10837)


__e.TailApply(tmp10789, tmp10838)
return


}, 1)

tmp10839 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10788)


_ = tmp10839

tmp10840 := MakeNative(func(__e *ControlFlow) {
V1454 := __e.Get(1)
_ = V1454
tmp10841 := MakeNative(func(__e *ControlFlow) {
W1455 := __e.Get(1)
_ = W1455
tmp10843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1455)


if True == tmp10843 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1455)
return
}


}, 1)

tmp10854 := PrimIsPair(V1454)

var ifres10844 Obj

if True == tmp10854 {
tmp10845 := MakeNative(func(__e *ControlFlow) {
W1456 := __e.Get(1)
_ = W1456
tmp10846 := MakeNative(func(__e *ControlFlow) {
W1457 := __e.Get(1)
_ = W1457
tmp10849 := PrimEqual(W1456, sym__)

if True == tmp10849 {
tmp10847 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W1457, tmp10847)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10850 := Call(__e, PrimFunc(symtail), V1454)


__e.TailApply(tmp10846, tmp10850)
return


}, 1)

tmp10851 := Call(__e, PrimFunc(symhead), V1454)


tmp10852 := Call(__e, tmp10845, tmp10851)


ifres10844 = tmp10852


} else {
tmp10853 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10844 = tmp10853


}

__e.TailApply(tmp10841, ifres10844)
return


}, 1)

tmp10855 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10840)


_ = tmp10855

tmp10856 := MakeNative(func(__e *ControlFlow) {
V1458 := __e.Get(1)
_ = V1458
tmp10857 := MakeNative(func(__e *ControlFlow) {
W1459 := __e.Get(1)
_ = W1459
tmp10859 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1459)


if True == tmp10859 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1459)
return
}


}, 1)

tmp10869 := PrimIsPair(V1458)

var ifres10860 Obj

if True == tmp10869 {
tmp10861 := MakeNative(func(__e *ControlFlow) {
W1460 := __e.Get(1)
_ = W1460
tmp10862 := MakeNative(func(__e *ControlFlow) {
W1461 := __e.Get(1)
_ = W1461
tmp10864 := Call(__e, PrimFunc(symshen_4semicolon_2), W1460)


if True == tmp10864 {
__e.TailApply(PrimFunc(symshen_4comb), W1461, W1460)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10865 := Call(__e, PrimFunc(symtail), V1458)


__e.TailApply(tmp10862, tmp10865)
return


}, 1)

tmp10866 := Call(__e, PrimFunc(symhead), V1458)


tmp10867 := Call(__e, tmp10861, tmp10866)


ifres10860 = tmp10867


} else {
tmp10868 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10860 = tmp10868


}

__e.TailApply(tmp10857, ifres10860)
return


}, 1)

tmp10870 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10856)


_ = tmp10870

tmp10871 := MakeNative(func(__e *ControlFlow) {
V1462 := __e.Get(1)
_ = V1462
V1463 := __e.Get(2)
_ = V1463
tmp10872 := MakeNative(func(__e *ControlFlow) {
W1464 := __e.Get(1)
_ = W1464
tmp10873 := MakeNative(func(__e *ControlFlow) {
W1465 := __e.Get(1)
_ = W1465
tmp10874 := MakeNative(func(__e *ControlFlow) {
W1466 := __e.Get(1)
_ = W1466
tmp10875 := MakeNative(func(__e *ControlFlow) {
W1467 := __e.Get(1)
_ = W1467
tmp10876 := MakeNative(func(__e *ControlFlow) {
W1468 := __e.Get(1)
_ = W1468
tmp10877 := MakeNative(func(__e *ControlFlow) {
W1469 := __e.Get(1)
_ = W1469
tmp10878 := MakeNative(func(__e *ControlFlow) {
W1470 := __e.Get(1)
_ = W1470
tmp10879 := MakeNative(func(__e *ControlFlow) {
W1471 := __e.Get(1)
_ = W1471
tmp10880 := MakeNative(func(__e *ControlFlow) {
W1472 := __e.Get(1)
_ = W1472
__e.Return(W1472)
return
}, 1)

tmp10881 := PrimCons(sym_1_6, Nil)

tmp10882 := PrimCons(W1467, tmp10881)

tmp10883 := PrimCons(W1466, tmp10882)

tmp10884 := PrimCons(W1465, tmp10883)

tmp10885 := PrimCons(W1464, tmp10884)

tmp10886 := PrimCons(W1471, Nil)

tmp10887 := Call(__e, PrimFunc(symappend), tmp10885, tmp10886)


tmp10888 := Call(__e, PrimFunc(symappend), W1468, tmp10887)


tmp10889 := PrimCons(V1462, tmp10888)

tmp10890 := PrimCons(symdefine, tmp10889)

__e.TailApply(tmp10880, tmp10890)
return


}, 1)

var ifres10891 Obj

if True == W1469 {
tmp10892 := PrimCons(MakeNumber(1), Nil)

tmp10893 := PrimCons(W1466, tmp10892)

tmp10894 := PrimCons(sym_7, tmp10893)

tmp10895 := PrimCons(W1470, Nil)

tmp10896 := PrimCons(tmp10894, tmp10895)

tmp10897 := PrimCons(W1466, tmp10896)

tmp10898 := PrimCons(symlet, tmp10897)

ifres10891 = tmp10898


} else {
ifres10891 = W1470


}

__e.TailApply(tmp10879, ifres10891)
return


}, 1)

tmp10899 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1463, W1468, W1464, W1465, W1466, W1467, W1469)


__e.TailApply(tmp10878, tmp10899)
return


}, 1)

tmp10900 := Call(__e, PrimFunc(symshen_4hascut_2), V1463)


__e.TailApply(tmp10877, tmp10900)
return


}, 1)

tmp10901 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1463)


__e.TailApply(tmp10876, tmp10901)
return


}, 1)

tmp10902 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp10875, tmp10902)
return


}, 1)

tmp10903 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp10874, tmp10903)
return


}, 1)

tmp10904 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp10873, tmp10904)
return


}, 1)

tmp10905 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp10872, tmp10905)
return


}, 2)

tmp10906 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10871)


_ = tmp10906

tmp10907 := MakeNative(func(__e *ControlFlow) {
V1475 := __e.Get(1)
_ = V1475
tmp10917 := PrimEqual(sym_b, V1475)

if True == tmp10917 {
__e.Return(True)
return
} else {
tmp10915 := PrimIsPair(V1475)

if True == tmp10915 {
tmp10912 := PrimHead(V1475)

tmp10913 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10912)


if True == tmp10913 {
__e.Return(True)
return
} else {
tmp10909 := PrimTail(V1475)

tmp10910 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10909)


if True == tmp10910 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 1)

tmp10918 := Call(__e, ns2_1set, symshen_4hascut_2, tmp10907)


_ = tmp10918

tmp10919 := MakeNative(func(__e *ControlFlow) {
V1480 := __e.Get(1)
_ = V1480
tmp10928 := PrimIsPair(V1480)

var ifres10924 Obj

if True == tmp10928 {
tmp10926 := PrimHead(V1480)

tmp10927 := PrimIsPair(tmp10926)

var ifres10925 Obj

if True == tmp10927 {
ifres10925 = True


} else {
ifres10925 = False


}

ifres10924 = ifres10925


} else {
ifres10924 = False


}

if True == ifres10924 {
tmp10920 := PrimHead(V1480)

tmp10921 := PrimHead(tmp10920)

tmp10922 := Call(__e, PrimFunc(symlength), tmp10921)


__e.TailApply(PrimFunc(symshen_4parameters), tmp10922)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prolog-parameters")))
return
}


}, 1)

tmp10929 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp10919)


_ = tmp10929

tmp10930 := MakeNative(func(__e *ControlFlow) {
V1501 := __e.Get(1)
_ = V1501
V1502 := __e.Get(2)
_ = V1502
V1503 := __e.Get(3)
_ = V1503
V1504 := __e.Get(4)
_ = V1504
V1505 := __e.Get(5)
_ = V1505
V1506 := __e.Get(6)
_ = V1506
V1507 := __e.Get(7)
_ = V1507
tmp11023 := PrimEqual(Nil, V1501)

var ifres11020 Obj

if True == tmp11023 {
tmp11022 := PrimEqual(True, V1507)

var ifres11021 Obj

if True == tmp11022 {
ifres11021 = True


} else {
ifres11021 = False


}

ifres11020 = ifres11021


} else {
ifres11020 = False


}

if True == ifres11020 {
tmp10931 := PrimCons(V1505, Nil)

tmp10932 := PrimCons(V1504, tmp10931)

__e.Return(PrimCons(symshen_4unlock, tmp10932))
return


} else {
tmp11018 := PrimIsPair(V1501)

var ifres10996 Obj

if True == tmp11018 {
tmp11016 := PrimHead(V1501)

tmp11017 := PrimIsPair(tmp11016)

var ifres10998 Obj

if True == tmp11017 {
tmp11013 := PrimHead(V1501)

tmp11014 := PrimTail(tmp11013)

tmp11015 := PrimIsPair(tmp11014)

var ifres11000 Obj

if True == tmp11015 {
tmp11009 := PrimHead(V1501)

tmp11010 := PrimTail(tmp11009)

tmp11011 := PrimTail(tmp11010)

tmp11012 := PrimEqual(Nil, tmp11011)

var ifres11002 Obj

if True == tmp11012 {
tmp11007 := PrimTail(V1501)

tmp11008 := PrimEqual(Nil, tmp11007)

var ifres11004 Obj

if True == tmp11008 {
tmp11006 := PrimEqual(False, V1507)

var ifres11005 Obj

if True == tmp11006 {
ifres11005 = True


} else {
ifres11005 = False


}

ifres11004 = ifres11005


} else {
ifres11004 = False


}

var ifres11003 Obj

if True == ifres11004 {
ifres11003 = True


} else {
ifres11003 = False


}

ifres11002 = ifres11003


} else {
ifres11002 = False


}

var ifres11001 Obj

if True == ifres11002 {
ifres11001 = True


} else {
ifres11001 = False


}

ifres11000 = ifres11001


} else {
ifres11000 = False


}

var ifres10999 Obj

if True == ifres11000 {
ifres10999 = True


} else {
ifres10999 = False


}

ifres10998 = ifres10999


} else {
ifres10998 = False


}

var ifres10997 Obj

if True == ifres10998 {
ifres10997 = True


} else {
ifres10997 = False


}

ifres10996 = ifres10997


} else {
ifres10996 = False


}

if True == ifres10996 {
tmp10933 := MakeNative(func(__e *ControlFlow) {
W1508 := __e.Get(1)
_ = W1508
tmp10934 := PrimCons(V1504, Nil)

tmp10935 := PrimCons(symshen_4unlocked_2, tmp10934)

tmp10936 := PrimHead(V1501)

tmp10937 := PrimHead(tmp10936)

tmp10938 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10937, V1502, V1503, W1508)


tmp10939 := PrimCons(False, Nil)

tmp10940 := PrimCons(tmp10938, tmp10939)

tmp10941 := PrimCons(tmp10935, tmp10940)

__e.Return(PrimCons(symif, tmp10941))
return


}, 1)

tmp10942 := PrimHead(V1501)

tmp10943 := PrimHead(tmp10942)

tmp10944 := PrimHead(V1501)

tmp10945 := PrimTail(tmp10944)

tmp10946 := PrimHead(tmp10945)

tmp10947 := Call(__e, PrimFunc(symshen_4continue), tmp10943, tmp10946, V1503, V1504, V1505, V1506)


__e.TailApply(tmp10933, tmp10947)
return


} else {
tmp10994 := PrimIsPair(V1501)

var ifres10979 Obj

if True == tmp10994 {
tmp10992 := PrimHead(V1501)

tmp10993 := PrimIsPair(tmp10992)

var ifres10981 Obj

if True == tmp10993 {
tmp10989 := PrimHead(V1501)

tmp10990 := PrimTail(tmp10989)

tmp10991 := PrimIsPair(tmp10990)

var ifres10983 Obj

if True == tmp10991 {
tmp10985 := PrimHead(V1501)

tmp10986 := PrimTail(tmp10985)

tmp10987 := PrimTail(tmp10986)

tmp10988 := PrimEqual(Nil, tmp10987)

var ifres10984 Obj

if True == tmp10988 {
ifres10984 = True


} else {
ifres10984 = False


}

ifres10983 = ifres10984


} else {
ifres10983 = False


}

var ifres10982 Obj

if True == ifres10983 {
ifres10982 = True


} else {
ifres10982 = False


}

ifres10981 = ifres10982


} else {
ifres10981 = False


}

var ifres10980 Obj

if True == ifres10981 {
ifres10980 = True


} else {
ifres10980 = False


}

ifres10979 = ifres10980


} else {
ifres10979 = False


}

if True == ifres10979 {
tmp10948 := MakeNative(func(__e *ControlFlow) {
W1509 := __e.Get(1)
_ = W1509
tmp10949 := MakeNative(func(__e *ControlFlow) {
W1510 := __e.Get(1)
_ = W1510
tmp10950 := PrimCons(V1504, Nil)

tmp10951 := PrimCons(symshen_4unlocked_2, tmp10950)

tmp10952 := PrimHead(V1501)

tmp10953 := PrimHead(tmp10952)

tmp10954 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10953, V1502, V1503, W1510)


tmp10955 := PrimCons(False, Nil)

tmp10956 := PrimCons(tmp10954, tmp10955)

tmp10957 := PrimCons(tmp10951, tmp10956)

tmp10958 := PrimCons(symif, tmp10957)

tmp10959 := PrimCons(False, Nil)

tmp10960 := PrimCons(W1509, tmp10959)

tmp10961 := PrimCons(sym_a, tmp10960)

tmp10962 := PrimTail(V1501)

tmp10963 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp10962, V1502, V1503, V1504, V1505, V1506, V1507)


tmp10964 := PrimCons(W1509, Nil)

tmp10965 := PrimCons(tmp10963, tmp10964)

tmp10966 := PrimCons(tmp10961, tmp10965)

tmp10967 := PrimCons(symif, tmp10966)

tmp10968 := PrimCons(tmp10967, Nil)

tmp10969 := PrimCons(tmp10958, tmp10968)

tmp10970 := PrimCons(W1509, tmp10969)

__e.Return(PrimCons(symlet, tmp10970))
return


}, 1)

tmp10971 := PrimHead(V1501)

tmp10972 := PrimHead(tmp10971)

tmp10973 := PrimHead(V1501)

tmp10974 := PrimTail(tmp10973)

tmp10975 := PrimHead(tmp10974)

tmp10976 := Call(__e, PrimFunc(symshen_4continue), tmp10972, tmp10975, V1503, V1504, V1505, V1506)


__e.TailApply(tmp10949, tmp10976)
return


}, 1)

tmp10977 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp10948, tmp10977)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
return
}


}


}


}, 7)

tmp11024 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp10930)


_ = tmp11024

tmp11025 := MakeNative(func(__e *ControlFlow) {
V1511 := __e.Get(1)
_ = V1511
V1512 := __e.Get(2)
_ = V1512
tmp11030 := Call(__e, PrimFunc(symshen_4locked_2), V1511)


var ifres11027 Obj

if True == tmp11030 {
tmp11029 := Call(__e, PrimFunc(symshen_4fits_2), V1512, V1511)


var ifres11028 Obj

if True == tmp11029 {
ifres11028 = True


} else {
ifres11028 = False


}

ifres11027 = ifres11028


} else {
ifres11027 = False


}

if True == ifres11027 {
__e.TailApply(PrimFunc(symshen_4openlock), V1511)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11031 := Call(__e, ns2_1set, symshen_4unlock, tmp11025)


_ = tmp11031

tmp11032 := MakeNative(func(__e *ControlFlow) {
V1513 := __e.Get(1)
_ = V1513
tmp11033 := Call(__e, PrimFunc(symshen_4unlocked_2), V1513)


__e.Return(PrimNot(tmp11033))
return


}, 1)

tmp11034 := Call(__e, ns2_1set, symshen_4locked_2, tmp11032)


_ = tmp11034

tmp11035 := MakeNative(func(__e *ControlFlow) {
V1514 := __e.Get(1)
_ = V1514
__e.Return(PrimVectorGet(V1514, MakeNumber(1)))
return
}, 1)

tmp11036 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11035)


_ = tmp11036

tmp11037 := MakeNative(func(__e *ControlFlow) {
V1515 := __e.Get(1)
_ = V1515
tmp11038 := PrimVectorSet(V1515, MakeNumber(1), True)

_ = tmp11038

__e.Return(False)
return


}, 1)

tmp11039 := Call(__e, ns2_1set, symshen_4openlock, tmp11037)


_ = tmp11039

tmp11040 := MakeNative(func(__e *ControlFlow) {
V1516 := __e.Get(1)
_ = V1516
V1517 := __e.Get(2)
_ = V1517
tmp11041 := PrimVectorGet(V1517, MakeNumber(2))

__e.Return(PrimEqual(V1516, tmp11041))
return


}, 2)

tmp11042 := Call(__e, ns2_1set, symshen_4fits_2, tmp11040)


_ = tmp11042

tmp11043 := MakeNative(func(__e *ControlFlow) {
V1520 := __e.Get(1)
_ = V1520
V1521 := __e.Get(2)
_ = V1521
V1522 := __e.Get(3)
_ = V1522
V1523 := __e.Get(4)
_ = V1523
tmp11044 := MakeNative(func(__e *ControlFlow) {
W1524 := __e.Get(1)
_ = W1524
tmp11049 := PrimEqual(W1524, False)

var ifres11046 Obj

if True == tmp11049 {
tmp11048 := Call(__e, PrimFunc(symshen_4unlocked_2), V1521)


var ifres11047 Obj

if True == tmp11048 {
ifres11047 = True


} else {
ifres11047 = False


}

ifres11046 = ifres11047


} else {
ifres11046 = False


}

if True == ifres11046 {
__e.TailApply(PrimFunc(symshen_4lock), V1522, V1521)
return
} else {
__e.Return(W1524)
return
}


}, 1)

tmp11050 := Call(__e, PrimFunc(symthaw), V1523)


__e.TailApply(tmp11044, tmp11050)
return


}, 4)

tmp11051 := Call(__e, ns2_1set, symshen_4cut, tmp11043)


_ = tmp11051

tmp11052 := MakeNative(func(__e *ControlFlow) {
V1525 := __e.Get(1)
_ = V1525
V1526 := __e.Get(2)
_ = V1526
tmp11053 := MakeNative(func(__e *ControlFlow) {
W1527 := __e.Get(1)
_ = W1527
tmp11054 := MakeNative(func(__e *ControlFlow) {
W1528 := __e.Get(1)
_ = W1528
__e.Return(False)
return
}, 1)

tmp11055 := PrimVectorSet(V1526, MakeNumber(2), V1525)

__e.TailApply(tmp11054, tmp11055)
return


}, 1)

tmp11056 := PrimVectorSet(V1526, MakeNumber(1), False)

__e.TailApply(tmp11053, tmp11056)
return


}, 2)

tmp11057 := Call(__e, ns2_1set, symshen_4lock, tmp11052)


_ = tmp11057

tmp11058 := MakeNative(func(__e *ControlFlow) {
V1529 := __e.Get(1)
_ = V1529
V1530 := __e.Get(2)
_ = V1530
V1531 := __e.Get(3)
_ = V1531
V1532 := __e.Get(4)
_ = V1532
V1533 := __e.Get(5)
_ = V1533
V1534 := __e.Get(6)
_ = V1534
tmp11059 := MakeNative(func(__e *ControlFlow) {
W1535 := __e.Get(1)
_ = W1535
tmp11060 := MakeNative(func(__e *ControlFlow) {
W1536 := __e.Get(1)
_ = W1536
tmp11061 := MakeNative(func(__e *ControlFlow) {
W1537 := __e.Get(1)
_ = W1537
tmp11062 := MakeNative(func(__e *ControlFlow) {
W1538 := __e.Get(1)
_ = W1538
__e.TailApply(PrimFunc(symshen_4stpart), W1537, W1538, V1531)
return
}, 1)

tmp11063 := PrimCons(symshen_4incinfs, Nil)

tmp11064 := Call(__e, PrimFunc(symshen_4compile_1body), V1530, V1531, V1532, V1533, V1534)


tmp11065 := PrimCons(tmp11064, Nil)

tmp11066 := PrimCons(tmp11063, tmp11065)

tmp11067 := PrimCons(symdo, tmp11066)

__e.TailApply(tmp11062, tmp11067)
return


}, 1)

tmp11068 := Call(__e, PrimFunc(symdifference), W1536, W1535)


__e.TailApply(tmp11061, tmp11068)
return


}, 1)

tmp11069 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1530)


__e.TailApply(tmp11060, tmp11069)
return


}, 1)

tmp11070 := Call(__e, PrimFunc(symshen_4extract_1vars), V1529)


__e.TailApply(tmp11059, tmp11070)
return


}, 6)

tmp11071 := Call(__e, ns2_1set, symshen_4continue, tmp11058)


_ = tmp11071

tmp11072 := MakeNative(func(__e *ControlFlow) {
V1541 := __e.Get(1)
_ = V1541
tmp11107 := PrimIsPair(V1541)

var ifres11088 Obj

if True == tmp11107 {
tmp11105 := PrimHead(V1541)

tmp11106 := PrimEqual(symlambda, tmp11105)

var ifres11090 Obj

if True == tmp11106 {
tmp11103 := PrimTail(V1541)

tmp11104 := PrimIsPair(tmp11103)

var ifres11092 Obj

if True == tmp11104 {
tmp11100 := PrimTail(V1541)

tmp11101 := PrimTail(tmp11100)

tmp11102 := PrimIsPair(tmp11101)

var ifres11094 Obj

if True == tmp11102 {
tmp11096 := PrimTail(V1541)

tmp11097 := PrimTail(tmp11096)

tmp11098 := PrimTail(tmp11097)

tmp11099 := PrimEqual(Nil, tmp11098)

var ifres11095 Obj

if True == tmp11099 {
ifres11095 = True


} else {
ifres11095 = False


}

ifres11094 = ifres11095


} else {
ifres11094 = False


}

var ifres11093 Obj

if True == ifres11094 {
ifres11093 = True


} else {
ifres11093 = False


}

ifres11092 = ifres11093


} else {
ifres11092 = False


}

var ifres11091 Obj

if True == ifres11092 {
ifres11091 = True


} else {
ifres11091 = False


}

ifres11090 = ifres11091


} else {
ifres11090 = False


}

var ifres11089 Obj

if True == ifres11090 {
ifres11089 = True


} else {
ifres11089 = False


}

ifres11088 = ifres11089


} else {
ifres11088 = False


}

if True == ifres11088 {
tmp11073 := PrimTail(V1541)

tmp11074 := PrimHead(tmp11073)

tmp11075 := PrimTail(V1541)

tmp11076 := PrimTail(tmp11075)

tmp11077 := PrimHead(tmp11076)

tmp11078 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11077)


__e.TailApply(PrimFunc(symremove), tmp11074, tmp11078)
return


} else {
tmp11086 := PrimIsPair(V1541)

if True == tmp11086 {
tmp11079 := PrimHead(V1541)

tmp11080 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11079)


tmp11081 := PrimTail(V1541)

tmp11082 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11081)


__e.TailApply(PrimFunc(symunion), tmp11080, tmp11082)
return


} else {
tmp11084 := PrimIsVariable(V1541)

if True == tmp11084 {
__e.Return(PrimCons(V1541, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp11108 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11072)


_ = tmp11108

tmp11109 := MakeNative(func(__e *ControlFlow) {
V1558 := __e.Get(1)
_ = V1558
V1559 := __e.Get(2)
_ = V1559
V1560 := __e.Get(3)
_ = V1560
V1561 := __e.Get(4)
_ = V1561
V1562 := __e.Get(5)
_ = V1562
tmp11144 := PrimEqual(Nil, V1558)

if True == tmp11144 {
tmp11110 := PrimCons(V1562, Nil)

__e.Return(PrimCons(symthaw, tmp11110))
return


} else {
tmp11142 := PrimIsPair(V1558)

var ifres11138 Obj

if True == tmp11142 {
tmp11140 := PrimHead(V1558)

tmp11141 := PrimEqual(sym_b, tmp11140)

var ifres11139 Obj

if True == tmp11141 {
ifres11139 = True


} else {
ifres11139 = False


}

ifres11138 = ifres11139


} else {
ifres11138 = False


}

if True == ifres11138 {
tmp11111 := PrimCons(symshen_4cut, Nil)

tmp11112 := PrimTail(V1558)

tmp11113 := PrimCons(tmp11111, tmp11112)

__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11113, V1559, V1560, V1561, V1562)
return


} else {
tmp11136 := PrimIsPair(V1558)

var ifres11132 Obj

if True == tmp11136 {
tmp11134 := PrimTail(V1558)

tmp11135 := PrimEqual(Nil, tmp11134)

var ifres11133 Obj

if True == tmp11135 {
ifres11133 = True


} else {
ifres11133 = False


}

ifres11132 = ifres11133


} else {
ifres11132 = False


}

if True == ifres11132 {
tmp11114 := PrimHead(V1558)

tmp11115 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11114, V1559)


tmp11116 := PrimCons(V1562, Nil)

tmp11117 := PrimCons(V1561, tmp11116)

tmp11118 := PrimCons(V1560, tmp11117)

tmp11119 := PrimCons(V1559, tmp11118)

__e.TailApply(PrimFunc(symappend), tmp11115, tmp11119)
return


} else {
tmp11130 := PrimIsPair(V1558)

if True == tmp11130 {
tmp11120 := MakeNative(func(__e *ControlFlow) {
W1563 := __e.Get(1)
_ = W1563
tmp11121 := PrimTail(V1558)

tmp11122 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11121, V1559, V1560, V1561, V1562)


tmp11123 := PrimCons(tmp11122, Nil)

tmp11124 := PrimCons(V1561, tmp11123)

tmp11125 := PrimCons(V1560, tmp11124)

tmp11126 := PrimCons(V1559, tmp11125)

__e.TailApply(PrimFunc(symappend), W1563, tmp11126)
return


}, 1)

tmp11127 := PrimHead(V1558)

tmp11128 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11127, V1559)


__e.TailApply(tmp11120, tmp11128)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
return
}


}


}


}


}, 5)

tmp11145 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11109)


_ = tmp11145

tmp11146 := MakeNative(func(__e *ControlFlow) {
V1580 := __e.Get(1)
_ = V1580
V1581 := __e.Get(2)
_ = V1581
V1582 := __e.Get(3)
_ = V1582
V1583 := __e.Get(4)
_ = V1583
V1584 := __e.Get(5)
_ = V1584
tmp11170 := PrimEqual(Nil, V1580)

if True == tmp11170 {
__e.Return(V1584)
return
} else {
tmp11168 := PrimIsPair(V1580)

var ifres11164 Obj

if True == tmp11168 {
tmp11166 := PrimHead(V1580)

tmp11167 := PrimEqual(sym_b, tmp11166)

var ifres11165 Obj

if True == tmp11167 {
ifres11165 = True


} else {
ifres11165 = False


}

ifres11164 = ifres11165


} else {
ifres11164 = False


}

if True == ifres11164 {
tmp11147 := PrimCons(symshen_4cut, Nil)

tmp11148 := PrimTail(V1580)

tmp11149 := PrimCons(tmp11147, tmp11148)

__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11149, V1581, V1582, V1583, V1584)
return


} else {
tmp11162 := PrimIsPair(V1580)

if True == tmp11162 {
tmp11150 := MakeNative(func(__e *ControlFlow) {
W1585 := __e.Get(1)
_ = W1585
tmp11151 := PrimTail(V1580)

tmp11152 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11151, V1581, V1582, V1583, V1584)


tmp11153 := PrimCons(tmp11152, Nil)

tmp11154 := PrimCons(V1583, tmp11153)

tmp11155 := PrimCons(V1582, tmp11154)

tmp11156 := PrimCons(V1581, tmp11155)

tmp11157 := Call(__e, PrimFunc(symappend), W1585, tmp11156)


tmp11158 := PrimCons(tmp11157, Nil)

__e.Return(PrimCons(symfreeze, tmp11158))
return


}, 1)

tmp11159 := PrimHead(V1580)

tmp11160 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11159, V1581)


__e.TailApply(tmp11150, tmp11160)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
return
}


}


}


}, 5)

tmp11171 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11146)


_ = tmp11171

tmp11172 := MakeNative(func(__e *ControlFlow) {
V1590 := __e.Get(1)
_ = V1590
V1591 := __e.Get(2)
_ = V1591
tmp11187 := PrimIsPair(V1590)

var ifres11183 Obj

if True == tmp11187 {
tmp11185 := PrimHead(V1590)

tmp11186 := PrimEqual(symfork, tmp11185)

var ifres11184 Obj

if True == tmp11186 {
ifres11184 = True


} else {
ifres11184 = False


}

ifres11183 = ifres11184


} else {
ifres11183 = False


}

if True == ifres11183 {
tmp11173 := PrimTail(V1590)

tmp11174 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11173, V1591)


tmp11175 := PrimCons(tmp11174, Nil)

__e.Return(PrimCons(symfork, tmp11175))
return


} else {
tmp11181 := PrimIsPair(V1590)

if True == tmp11181 {
tmp11176 := PrimHead(V1590)

tmp11177 := MakeNative(func(__e *ControlFlow) {
Z1592 := __e.Get(1)
_ = Z1592
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1592, V1591)
return
}, 1)

tmp11178 := PrimTail(V1590)

tmp11179 := Call(__e, PrimFunc(symmap), tmp11177, tmp11178)


__e.Return(PrimCons(tmp11176, tmp11179))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
return
}


}


}, 2)

tmp11188 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11172)


_ = tmp11188

tmp11189 := MakeNative(func(__e *ControlFlow) {
V1599 := __e.Get(1)
_ = V1599
V1600 := __e.Get(2)
_ = V1600
tmp11199 := PrimEqual(Nil, V1599)

if True == tmp11199 {
__e.Return(Nil)
return
} else {
tmp11197 := PrimIsPair(V1599)

if True == tmp11197 {
tmp11190 := PrimHead(V1599)

tmp11191 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11190, V1600)


tmp11192 := PrimTail(V1599)

tmp11193 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11192, V1600)


tmp11194 := PrimCons(tmp11193, Nil)

tmp11195 := PrimCons(tmp11191, tmp11194)

__e.Return(PrimCons(symcons, tmp11195))
return


} else {
__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
return
}


}


}, 2)

tmp11200 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11189)


_ = tmp11200

tmp11201 := MakeNative(func(__e *ControlFlow) {
V1603 := __e.Get(1)
_ = V1603
V1604 := __e.Get(2)
_ = V1604
tmp11233 := PrimIsPair(V1603)

var ifres11214 Obj

if True == tmp11233 {
tmp11231 := PrimHead(V1603)

tmp11232 := PrimEqual(symcons, tmp11231)

var ifres11216 Obj

if True == tmp11232 {
tmp11229 := PrimTail(V1603)

tmp11230 := PrimIsPair(tmp11229)

var ifres11218 Obj

if True == tmp11230 {
tmp11226 := PrimTail(V1603)

tmp11227 := PrimTail(tmp11226)

tmp11228 := PrimIsPair(tmp11227)

var ifres11220 Obj

if True == tmp11228 {
tmp11222 := PrimTail(V1603)

tmp11223 := PrimTail(tmp11222)

tmp11224 := PrimTail(tmp11223)

tmp11225 := PrimEqual(Nil, tmp11224)

var ifres11221 Obj

if True == tmp11225 {
ifres11221 = True


} else {
ifres11221 = False


}

ifres11220 = ifres11221


} else {
ifres11220 = False


}

var ifres11219 Obj

if True == ifres11220 {
ifres11219 = True


} else {
ifres11219 = False


}

ifres11218 = ifres11219


} else {
ifres11218 = False


}

var ifres11217 Obj

if True == ifres11218 {
ifres11217 = True


} else {
ifres11217 = False


}

ifres11216 = ifres11217


} else {
ifres11216 = False


}

var ifres11215 Obj

if True == ifres11216 {
ifres11215 = True


} else {
ifres11215 = False


}

ifres11214 = ifres11215


} else {
ifres11214 = False


}

if True == ifres11214 {
tmp11202 := PrimTail(V1603)

tmp11203 := PrimHead(tmp11202)

tmp11204 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11203, V1604)


tmp11205 := PrimTail(V1603)

tmp11206 := PrimTail(tmp11205)

tmp11207 := PrimHead(tmp11206)

tmp11208 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11207, V1604)


tmp11209 := PrimCons(tmp11208, Nil)

tmp11210 := PrimCons(tmp11204, tmp11209)

__e.Return(PrimCons(symcons, tmp11210))
return


} else {
tmp11212 := PrimIsPair(V1603)

if True == tmp11212 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1603, V1604, Nil)
return
} else {
__e.Return(V1603)
return
}


}


}, 2)

tmp11234 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11201)


_ = tmp11234

tmp11235 := MakeNative(func(__e *ControlFlow) {
V1613 := __e.Get(1)
_ = V1613
V1614 := __e.Get(2)
_ = V1614
V1615 := __e.Get(3)
_ = V1615
tmp11329 := PrimIsPair(V1613)

var ifres11316 Obj

if True == tmp11329 {
tmp11327 := PrimHead(V1613)

tmp11328 := PrimEqual(MakeNumber(0), tmp11327)

var ifres11318 Obj

if True == tmp11328 {
tmp11325 := PrimTail(V1613)

tmp11326 := PrimIsPair(tmp11325)

var ifres11320 Obj

if True == tmp11326 {
tmp11322 := PrimTail(V1613)

tmp11323 := PrimTail(tmp11322)

tmp11324 := PrimEqual(Nil, tmp11323)

var ifres11321 Obj

if True == tmp11324 {
ifres11321 = True


} else {
ifres11321 = False


}

ifres11320 = ifres11321


} else {
ifres11320 = False


}

var ifres11319 Obj

if True == ifres11320 {
ifres11319 = True


} else {
ifres11319 = False


}

ifres11318 = ifres11319


} else {
ifres11318 = False


}

var ifres11317 Obj

if True == ifres11318 {
ifres11317 = True


} else {
ifres11317 = False


}

ifres11316 = ifres11317


} else {
ifres11316 = False


}

if True == ifres11316 {
tmp11242 := PrimTail(V1613)

tmp11243 := PrimHead(tmp11242)

tmp11244 := PrimIsVariable(tmp11243)

if True == tmp11244 {
tmp11236 := PrimTail(V1613)

__e.Return(PrimHead(tmp11236))
return


} else {
tmp11237 := PrimTail(V1613)

tmp11238 := PrimHead(tmp11237)

tmp11239 := Call(__e, PrimFunc(symshen_4app), tmp11238, MakeString("\n"), symshen_4s)


tmp11240 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp11239)

__e.Return(PrimSimpleError(tmp11240))
return


}


} else {
tmp11314 := PrimIsPair(V1613)

var ifres11301 Obj

if True == tmp11314 {
tmp11312 := PrimHead(V1613)

tmp11313 := PrimEqual(MakeNumber(1), tmp11312)

var ifres11303 Obj

if True == tmp11313 {
tmp11310 := PrimTail(V1613)

tmp11311 := PrimIsPair(tmp11310)

var ifres11305 Obj

if True == tmp11311 {
tmp11307 := PrimTail(V1613)

tmp11308 := PrimTail(tmp11307)

tmp11309 := PrimEqual(Nil, tmp11308)

var ifres11306 Obj

if True == tmp11309 {
ifres11306 = True


} else {
ifres11306 = False


}

ifres11305 = ifres11306


} else {
ifres11305 = False


}

var ifres11304 Obj

if True == ifres11305 {
ifres11304 = True


} else {
ifres11304 = False


}

ifres11303 = ifres11304


} else {
ifres11303 = False


}

var ifres11302 Obj

if True == ifres11303 {
ifres11302 = True


} else {
ifres11302 = False


}

ifres11301 = ifres11302


} else {
ifres11301 = False


}

if True == ifres11301 {
tmp11254 := PrimTail(V1613)

tmp11255 := PrimHead(tmp11254)

tmp11256 := PrimIsVariable(tmp11255)

if True == tmp11256 {
tmp11245 := PrimTail(V1613)

tmp11246 := PrimHead(tmp11245)

tmp11247 := PrimCons(V1614, Nil)

tmp11248 := PrimCons(tmp11246, tmp11247)

__e.Return(PrimCons(symshen_4lazyderef, tmp11248))
return


} else {
tmp11249 := PrimTail(V1613)

tmp11250 := PrimHead(tmp11249)

tmp11251 := Call(__e, PrimFunc(symshen_4app), tmp11250, MakeString("\n"), symshen_4s)


tmp11252 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp11251)

__e.Return(PrimSimpleError(tmp11252))
return


}


} else {
tmp11298 := Call(__e, PrimFunc(symelement_2), V1613, V1615)


tmp11299 := PrimNot(tmp11298)

var ifres11295 Obj

if True == tmp11299 {
tmp11297 := PrimIsVariable(V1613)

var ifres11296 Obj

if True == tmp11297 {
ifres11296 = True


} else {
ifres11296 = False


}

ifres11295 = ifres11296


} else {
ifres11295 = False


}

if True == ifres11295 {
tmp11257 := PrimCons(V1614, Nil)

tmp11258 := PrimCons(V1613, tmp11257)

__e.Return(PrimCons(symshen_4deref, tmp11258))
return


} else {
tmp11293 := PrimIsPair(V1613)

var ifres11274 Obj

if True == tmp11293 {
tmp11291 := PrimHead(V1613)

tmp11292 := PrimEqual(symlambda, tmp11291)

var ifres11276 Obj

if True == tmp11292 {
tmp11289 := PrimTail(V1613)

tmp11290 := PrimIsPair(tmp11289)

var ifres11278 Obj

if True == tmp11290 {
tmp11286 := PrimTail(V1613)

tmp11287 := PrimTail(tmp11286)

tmp11288 := PrimIsPair(tmp11287)

var ifres11280 Obj

if True == tmp11288 {
tmp11282 := PrimTail(V1613)

tmp11283 := PrimTail(tmp11282)

tmp11284 := PrimTail(tmp11283)

tmp11285 := PrimEqual(Nil, tmp11284)

var ifres11281 Obj

if True == tmp11285 {
ifres11281 = True


} else {
ifres11281 = False


}

ifres11280 = ifres11281


} else {
ifres11280 = False


}

var ifres11279 Obj

if True == ifres11280 {
ifres11279 = True


} else {
ifres11279 = False


}

ifres11278 = ifres11279


} else {
ifres11278 = False


}

var ifres11277 Obj

if True == ifres11278 {
ifres11277 = True


} else {
ifres11277 = False


}

ifres11276 = ifres11277


} else {
ifres11276 = False


}

var ifres11275 Obj

if True == ifres11276 {
ifres11275 = True


} else {
ifres11275 = False


}

ifres11274 = ifres11275


} else {
ifres11274 = False


}

if True == ifres11274 {
tmp11259 := PrimTail(V1613)

tmp11260 := PrimHead(tmp11259)

tmp11261 := PrimTail(V1613)

tmp11262 := PrimTail(tmp11261)

tmp11263 := PrimHead(tmp11262)

tmp11264 := PrimTail(V1613)

tmp11265 := PrimHead(tmp11264)

tmp11266 := PrimCons(tmp11265, V1615)

tmp11267 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11263, V1614, tmp11266)


tmp11268 := PrimCons(tmp11267, Nil)

tmp11269 := PrimCons(tmp11260, tmp11268)

__e.Return(PrimCons(symlambda, tmp11269))
return


} else {
tmp11272 := PrimIsPair(V1613)

if True == tmp11272 {
tmp11270 := MakeNative(func(__e *ControlFlow) {
Z1616 := __e.Get(1)
_ = Z1616
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1616, V1614, V1615)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11270, V1613)
return


} else {
__e.Return(V1613)
return
}


}


}


}


}


}, 3)

tmp11330 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11235)


_ = tmp11330

tmp11331 := MakeNative(func(__e *ControlFlow) {
V1634 := __e.Get(1)
_ = V1634
V1635 := __e.Get(2)
_ = V1635
V1636 := __e.Get(3)
_ = V1636
V1637 := __e.Get(4)
_ = V1637
V1638 := __e.Get(5)
_ = V1638
tmp11507 := PrimEqual(Nil, V1635)

var ifres11504 Obj

if True == tmp11507 {
tmp11506 := PrimEqual(Nil, V1636)

var ifres11505 Obj

if True == tmp11506 {
ifres11505 = True


} else {
ifres11505 = False


}

ifres11504 = ifres11505


} else {
ifres11504 = False


}

if True == ifres11504 {
__e.Return(V1638)
return
} else {
tmp11502 := PrimIsPair(V1635)

var ifres11482 Obj

if True == tmp11502 {
tmp11500 := PrimHead(V1635)

tmp11501 := PrimIsPair(tmp11500)

var ifres11484 Obj

if True == tmp11501 {
tmp11497 := PrimHead(V1635)

tmp11498 := PrimHead(tmp11497)

tmp11499 := PrimEqual(symshen_4_7m, tmp11498)

var ifres11486 Obj

if True == tmp11499 {
tmp11494 := PrimHead(V1635)

tmp11495 := PrimTail(tmp11494)

tmp11496 := PrimIsPair(tmp11495)

var ifres11488 Obj

if True == tmp11496 {
tmp11490 := PrimHead(V1635)

tmp11491 := PrimTail(tmp11490)

tmp11492 := PrimTail(tmp11491)

tmp11493 := PrimEqual(Nil, tmp11492)

var ifres11489 Obj

if True == tmp11493 {
ifres11489 = True


} else {
ifres11489 = False


}

ifres11488 = ifres11489


} else {
ifres11488 = False


}

var ifres11487 Obj

if True == ifres11488 {
ifres11487 = True


} else {
ifres11487 = False


}

ifres11486 = ifres11487


} else {
ifres11486 = False


}

var ifres11485 Obj

if True == ifres11486 {
ifres11485 = True


} else {
ifres11485 = False


}

ifres11484 = ifres11485


} else {
ifres11484 = False


}

var ifres11483 Obj

if True == ifres11484 {
ifres11483 = True


} else {
ifres11483 = False


}

ifres11482 = ifres11483


} else {
ifres11482 = False


}

if True == ifres11482 {
tmp11332 := PrimHead(V1635)

tmp11333 := PrimTail(tmp11332)

tmp11334 := PrimHead(tmp11333)

tmp11335 := PrimTail(V1635)

tmp11336 := PrimCons(V1634, tmp11335)

tmp11337 := PrimCons(tmp11334, tmp11336)

tmp11338 := PrimCons(symshen_4_7m, tmp11337)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1634, tmp11338, V1636, V1637, V1638)
return


} else {
tmp11480 := PrimIsPair(V1635)

var ifres11460 Obj

if True == tmp11480 {
tmp11478 := PrimHead(V1635)

tmp11479 := PrimIsPair(tmp11478)

var ifres11462 Obj

if True == tmp11479 {
tmp11475 := PrimHead(V1635)

tmp11476 := PrimHead(tmp11475)

tmp11477 := PrimEqual(symshen_4_1m, tmp11476)

var ifres11464 Obj

if True == tmp11477 {
tmp11472 := PrimHead(V1635)

tmp11473 := PrimTail(tmp11472)

tmp11474 := PrimIsPair(tmp11473)

var ifres11466 Obj

if True == tmp11474 {
tmp11468 := PrimHead(V1635)

tmp11469 := PrimTail(tmp11468)

tmp11470 := PrimTail(tmp11469)

tmp11471 := PrimEqual(Nil, tmp11470)

var ifres11467 Obj

if True == tmp11471 {
ifres11467 = True


} else {
ifres11467 = False


}

ifres11466 = ifres11467


} else {
ifres11466 = False


}

var ifres11465 Obj

if True == ifres11466 {
ifres11465 = True


} else {
ifres11465 = False


}

ifres11464 = ifres11465


} else {
ifres11464 = False


}

var ifres11463 Obj

if True == ifres11464 {
ifres11463 = True


} else {
ifres11463 = False


}

ifres11462 = ifres11463


} else {
ifres11462 = False


}

var ifres11461 Obj

if True == ifres11462 {
ifres11461 = True


} else {
ifres11461 = False


}

ifres11460 = ifres11461


} else {
ifres11460 = False


}

if True == ifres11460 {
tmp11339 := PrimHead(V1635)

tmp11340 := PrimTail(tmp11339)

tmp11341 := PrimHead(tmp11340)

tmp11342 := PrimTail(V1635)

tmp11343 := PrimCons(V1634, tmp11342)

tmp11344 := PrimCons(tmp11341, tmp11343)

tmp11345 := PrimCons(symshen_4_1m, tmp11344)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1634, tmp11345, V1636, V1637, V1638)
return


} else {
tmp11458 := PrimIsPair(V1635)

var ifres11454 Obj

if True == tmp11458 {
tmp11456 := PrimHead(V1635)

tmp11457 := PrimEqual(symshen_4_1m, tmp11456)

var ifres11455 Obj

if True == tmp11457 {
ifres11455 = True


} else {
ifres11455 = False


}

ifres11454 = ifres11455


} else {
ifres11454 = False


}

if True == ifres11454 {
tmp11346 := PrimTail(V1635)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11346, V1636, V1637, V1638)
return


} else {
tmp11452 := PrimIsPair(V1635)

var ifres11448 Obj

if True == tmp11452 {
tmp11450 := PrimHead(V1635)

tmp11451 := PrimEqual(symshen_4_7m, tmp11450)

var ifres11449 Obj

if True == tmp11451 {
ifres11449 = True


} else {
ifres11449 = False


}

ifres11448 = ifres11449


} else {
ifres11448 = False


}

if True == ifres11448 {
tmp11347 := PrimTail(V1635)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11347, V1636, V1637, V1638)
return


} else {
tmp11446 := PrimIsPair(V1635)

var ifres11439 Obj

if True == tmp11446 {
tmp11445 := PrimIsPair(V1636)

var ifres11441 Obj

if True == tmp11445 {
tmp11443 := PrimHead(V1635)

tmp11444 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11443)


var ifres11442 Obj

if True == tmp11444 {
ifres11442 = True


} else {
ifres11442 = False


}

ifres11441 = ifres11442


} else {
ifres11441 = False


}

var ifres11440 Obj

if True == ifres11441 {
ifres11440 = True


} else {
ifres11440 = False


}

ifres11439 = ifres11440


} else {
ifres11439 = False


}

if True == ifres11439 {
tmp11348 := PrimTail(V1635)

tmp11349 := PrimTail(V1636)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1634, tmp11348, tmp11349, V1637, V1638)
return


} else {
tmp11437 := PrimIsPair(V1635)

var ifres11433 Obj

if True == tmp11437 {
tmp11435 := PrimHead(V1635)

tmp11436 := PrimIsVariable(tmp11435)

var ifres11434 Obj

if True == tmp11436 {
ifres11434 = True


} else {
ifres11434 = False


}

ifres11433 = ifres11434


} else {
ifres11433 = False


}

if True == ifres11433 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V1634, V1635, V1636, V1637, V1638)
return
} else {
tmp11431 := PrimEqual(symshen_4_1m, V1634)

var ifres11424 Obj

if True == tmp11431 {
tmp11430 := PrimIsPair(V1635)

var ifres11426 Obj

if True == tmp11430 {
tmp11428 := PrimHead(V1635)

tmp11429 := Call(__e, PrimFunc(symatom_2), tmp11428)


var ifres11427 Obj

if True == tmp11429 {
ifres11427 = True


} else {
ifres11427 = False


}

ifres11426 = ifres11427


} else {
ifres11426 = False


}

var ifres11425 Obj

if True == ifres11426 {
ifres11425 = True


} else {
ifres11425 = False


}

ifres11424 = ifres11425


} else {
ifres11424 = False


}

if True == ifres11424 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V1635, V1636, V1637, V1638)
return
} else {
tmp11422 := PrimEqual(symshen_4_1m, V1634)

var ifres11392 Obj

if True == tmp11422 {
tmp11421 := PrimIsPair(V1635)

var ifres11394 Obj

if True == tmp11421 {
tmp11419 := PrimHead(V1635)

tmp11420 := PrimIsPair(tmp11419)

var ifres11396 Obj

if True == tmp11420 {
tmp11416 := PrimHead(V1635)

tmp11417 := PrimHead(tmp11416)

tmp11418 := PrimEqual(symcons, tmp11417)

var ifres11398 Obj

if True == tmp11418 {
tmp11413 := PrimHead(V1635)

tmp11414 := PrimTail(tmp11413)

tmp11415 := PrimIsPair(tmp11414)

var ifres11400 Obj

if True == tmp11415 {
tmp11409 := PrimHead(V1635)

tmp11410 := PrimTail(tmp11409)

tmp11411 := PrimTail(tmp11410)

tmp11412 := PrimIsPair(tmp11411)

var ifres11402 Obj

if True == tmp11412 {
tmp11404 := PrimHead(V1635)

tmp11405 := PrimTail(tmp11404)

tmp11406 := PrimTail(tmp11405)

tmp11407 := PrimTail(tmp11406)

tmp11408 := PrimEqual(Nil, tmp11407)

var ifres11403 Obj

if True == tmp11408 {
ifres11403 = True


} else {
ifres11403 = False


}

ifres11402 = ifres11403


} else {
ifres11402 = False


}

var ifres11401 Obj

if True == ifres11402 {
ifres11401 = True


} else {
ifres11401 = False


}

ifres11400 = ifres11401


} else {
ifres11400 = False


}

var ifres11399 Obj

if True == ifres11400 {
ifres11399 = True


} else {
ifres11399 = False


}

ifres11398 = ifres11399


} else {
ifres11398 = False


}

var ifres11397 Obj

if True == ifres11398 {
ifres11397 = True


} else {
ifres11397 = False


}

ifres11396 = ifres11397


} else {
ifres11396 = False


}

var ifres11395 Obj

if True == ifres11396 {
ifres11395 = True


} else {
ifres11395 = False


}

ifres11394 = ifres11395


} else {
ifres11394 = False


}

var ifres11393 Obj

if True == ifres11394 {
ifres11393 = True


} else {
ifres11393 = False


}

ifres11392 = ifres11393


} else {
ifres11392 = False


}

if True == ifres11392 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V1635, V1636, V1637, V1638)
return
} else {
tmp11390 := PrimEqual(symshen_4_7m, V1634)

var ifres11383 Obj

if True == tmp11390 {
tmp11389 := PrimIsPair(V1635)

var ifres11385 Obj

if True == tmp11389 {
tmp11387 := PrimHead(V1635)

tmp11388 := Call(__e, PrimFunc(symatom_2), tmp11387)


var ifres11386 Obj

if True == tmp11388 {
ifres11386 = True


} else {
ifres11386 = False


}

ifres11385 = ifres11386


} else {
ifres11385 = False


}

var ifres11384 Obj

if True == ifres11385 {
ifres11384 = True


} else {
ifres11384 = False


}

ifres11383 = ifres11384


} else {
ifres11383 = False


}

if True == ifres11383 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V1635, V1636, V1637, V1638)
return
} else {
tmp11381 := PrimEqual(symshen_4_7m, V1634)

var ifres11351 Obj

if True == tmp11381 {
tmp11380 := PrimIsPair(V1635)

var ifres11353 Obj

if True == tmp11380 {
tmp11378 := PrimHead(V1635)

tmp11379 := PrimIsPair(tmp11378)

var ifres11355 Obj

if True == tmp11379 {
tmp11375 := PrimHead(V1635)

tmp11376 := PrimHead(tmp11375)

tmp11377 := PrimEqual(symcons, tmp11376)

var ifres11357 Obj

if True == tmp11377 {
tmp11372 := PrimHead(V1635)

tmp11373 := PrimTail(tmp11372)

tmp11374 := PrimIsPair(tmp11373)

var ifres11359 Obj

if True == tmp11374 {
tmp11368 := PrimHead(V1635)

tmp11369 := PrimTail(tmp11368)

tmp11370 := PrimTail(tmp11369)

tmp11371 := PrimIsPair(tmp11370)

var ifres11361 Obj

if True == tmp11371 {
tmp11363 := PrimHead(V1635)

tmp11364 := PrimTail(tmp11363)

tmp11365 := PrimTail(tmp11364)

tmp11366 := PrimTail(tmp11365)

tmp11367 := PrimEqual(Nil, tmp11366)

var ifres11362 Obj

if True == tmp11367 {
ifres11362 = True


} else {
ifres11362 = False


}

ifres11361 = ifres11362


} else {
ifres11361 = False


}

var ifres11360 Obj

if True == ifres11361 {
ifres11360 = True


} else {
ifres11360 = False


}

ifres11359 = ifres11360


} else {
ifres11359 = False


}

var ifres11358 Obj

if True == ifres11359 {
ifres11358 = True


} else {
ifres11358 = False


}

ifres11357 = ifres11358


} else {
ifres11357 = False


}

var ifres11356 Obj

if True == ifres11357 {
ifres11356 = True


} else {
ifres11356 = False


}

ifres11355 = ifres11356


} else {
ifres11355 = False


}

var ifres11354 Obj

if True == ifres11355 {
ifres11354 = True


} else {
ifres11354 = False


}

ifres11353 = ifres11354


} else {
ifres11353 = False


}

var ifres11352 Obj

if True == ifres11353 {
ifres11352 = True


} else {
ifres11352 = False


}

ifres11351 = ifres11352


} else {
ifres11351 = False


}

if True == ifres11351 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V1635, V1636, V1637, V1638)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-head")))
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


}, 5)

tmp11508 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11331)


_ = tmp11508

tmp11509 := MakeNative(func(__e *ControlFlow) {
V1649 := __e.Get(1)
_ = V1649
V1650 := __e.Get(2)
_ = V1650
V1651 := __e.Get(3)
_ = V1651
V1652 := __e.Get(4)
_ = V1652
V1653 := __e.Get(5)
_ = V1653
tmp11530 := PrimIsPair(V1650)

var ifres11527 Obj

if True == tmp11530 {
tmp11529 := PrimIsPair(V1651)

var ifres11528 Obj

if True == tmp11529 {
ifres11528 = True


} else {
ifres11528 = False


}

ifres11527 = ifres11528


} else {
ifres11527 = False


}

if True == ifres11527 {
tmp11524 := PrimHead(V1651)

tmp11525 := PrimIsVariable(tmp11524)

if True == tmp11525 {
tmp11510 := PrimTail(V1650)

tmp11511 := PrimTail(V1651)

tmp11512 := PrimHead(V1651)

tmp11513 := PrimHead(V1650)

tmp11514 := Call(__e, PrimFunc(symsubst), tmp11512, tmp11513, V1653)


__e.TailApply(PrimFunc(symshen_4compile_1head), V1649, tmp11510, tmp11511, V1652, tmp11514)
return


} else {
tmp11515 := PrimHead(V1650)

tmp11516 := PrimHead(V1651)

tmp11517 := PrimTail(V1650)

tmp11518 := PrimTail(V1651)

tmp11519 := Call(__e, PrimFunc(symshen_4compile_1head), V1649, tmp11517, tmp11518, V1652, V1653)


tmp11520 := PrimCons(tmp11519, Nil)

tmp11521 := PrimCons(tmp11516, tmp11520)

tmp11522 := PrimCons(tmp11515, tmp11521)

__e.Return(PrimCons(symlet, tmp11522))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
return
}


}, 5)

tmp11531 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11509)


_ = tmp11531

tmp11532 := MakeNative(func(__e *ControlFlow) {
V1662 := __e.Get(1)
_ = V1662
V1663 := __e.Get(2)
_ = V1663
V1664 := __e.Get(3)
_ = V1664
V1665 := __e.Get(4)
_ = V1665
tmp11557 := PrimIsPair(V1662)

var ifres11554 Obj

if True == tmp11557 {
tmp11556 := PrimIsPair(V1663)

var ifres11555 Obj

if True == tmp11556 {
ifres11555 = True


} else {
ifres11555 = False


}

ifres11554 = ifres11555


} else {
ifres11554 = False


}

if True == ifres11554 {
tmp11533 := MakeNative(func(__e *ControlFlow) {
W1666 := __e.Get(1)
_ = W1666
tmp11534 := PrimHead(V1663)

tmp11535 := PrimCons(V1664, Nil)

tmp11536 := PrimCons(tmp11534, tmp11535)

tmp11537 := PrimCons(symshen_4lazyderef, tmp11536)

tmp11538 := PrimHead(V1662)

tmp11539 := PrimCons(tmp11538, Nil)

tmp11540 := PrimCons(W1666, tmp11539)

tmp11541 := PrimCons(sym_a, tmp11540)

tmp11542 := PrimTail(V1662)

tmp11543 := PrimTail(V1663)

tmp11544 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11542, tmp11543, V1664, V1665)


tmp11545 := PrimCons(False, Nil)

tmp11546 := PrimCons(tmp11544, tmp11545)

tmp11547 := PrimCons(tmp11541, tmp11546)

tmp11548 := PrimCons(symif, tmp11547)

tmp11549 := PrimCons(tmp11548, Nil)

tmp11550 := PrimCons(tmp11537, tmp11549)

tmp11551 := PrimCons(W1666, tmp11550)

__e.Return(PrimCons(symlet, tmp11551))
return


}, 1)

tmp11552 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11533, tmp11552)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
return
}


}, 4)

tmp11558 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11532)


_ = tmp11558

tmp11559 := MakeNative(func(__e *ControlFlow) {
V1675 := __e.Get(1)
_ = V1675
V1676 := __e.Get(2)
_ = V1676
V1677 := __e.Get(3)
_ = V1677
V1678 := __e.Get(4)
_ = V1678
tmp11624 := PrimIsPair(V1675)

var ifres11594 Obj

if True == tmp11624 {
tmp11622 := PrimHead(V1675)

tmp11623 := PrimIsPair(tmp11622)

var ifres11596 Obj

if True == tmp11623 {
tmp11619 := PrimHead(V1675)

tmp11620 := PrimHead(tmp11619)

tmp11621 := PrimEqual(symcons, tmp11620)

var ifres11598 Obj

if True == tmp11621 {
tmp11616 := PrimHead(V1675)

tmp11617 := PrimTail(tmp11616)

tmp11618 := PrimIsPair(tmp11617)

var ifres11600 Obj

if True == tmp11618 {
tmp11612 := PrimHead(V1675)

tmp11613 := PrimTail(tmp11612)

tmp11614 := PrimTail(tmp11613)

tmp11615 := PrimIsPair(tmp11614)

var ifres11602 Obj

if True == tmp11615 {
tmp11607 := PrimHead(V1675)

tmp11608 := PrimTail(tmp11607)

tmp11609 := PrimTail(tmp11608)

tmp11610 := PrimTail(tmp11609)

tmp11611 := PrimEqual(Nil, tmp11610)

var ifres11604 Obj

if True == tmp11611 {
tmp11606 := PrimIsPair(V1676)

var ifres11605 Obj

if True == tmp11606 {
ifres11605 = True


} else {
ifres11605 = False


}

ifres11604 = ifres11605


} else {
ifres11604 = False


}

var ifres11603 Obj

if True == ifres11604 {
ifres11603 = True


} else {
ifres11603 = False


}

ifres11602 = ifres11603


} else {
ifres11602 = False


}

var ifres11601 Obj

if True == ifres11602 {
ifres11601 = True


} else {
ifres11601 = False


}

ifres11600 = ifres11601


} else {
ifres11600 = False


}

var ifres11599 Obj

if True == ifres11600 {
ifres11599 = True


} else {
ifres11599 = False


}

ifres11598 = ifres11599


} else {
ifres11598 = False


}

var ifres11597 Obj

if True == ifres11598 {
ifres11597 = True


} else {
ifres11597 = False


}

ifres11596 = ifres11597


} else {
ifres11596 = False


}

var ifres11595 Obj

if True == ifres11596 {
ifres11595 = True


} else {
ifres11595 = False


}

ifres11594 = ifres11595


} else {
ifres11594 = False


}

if True == ifres11594 {
tmp11560 := MakeNative(func(__e *ControlFlow) {
W1679 := __e.Get(1)
_ = W1679
tmp11561 := PrimHead(V1676)

tmp11562 := PrimCons(V1677, Nil)

tmp11563 := PrimCons(tmp11561, tmp11562)

tmp11564 := PrimCons(symshen_4lazyderef, tmp11563)

tmp11565 := PrimCons(W1679, Nil)

tmp11566 := PrimCons(symcons_2, tmp11565)

tmp11567 := PrimHead(V1675)

tmp11568 := PrimTail(tmp11567)

tmp11569 := PrimHead(tmp11568)

tmp11570 := PrimHead(V1675)

tmp11571 := PrimTail(tmp11570)

tmp11572 := PrimTail(tmp11571)

tmp11573 := PrimHead(tmp11572)

tmp11574 := PrimTail(V1675)

tmp11575 := PrimCons(tmp11573, tmp11574)

tmp11576 := PrimCons(tmp11569, tmp11575)

tmp11577 := PrimCons(W1679, Nil)

tmp11578 := PrimCons(symhd, tmp11577)

tmp11579 := PrimCons(W1679, Nil)

tmp11580 := PrimCons(symtl, tmp11579)

tmp11581 := PrimTail(V1676)

tmp11582 := PrimCons(tmp11580, tmp11581)

tmp11583 := PrimCons(tmp11578, tmp11582)

tmp11584 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11576, tmp11583, V1677, V1678)


tmp11585 := PrimCons(False, Nil)

tmp11586 := PrimCons(tmp11584, tmp11585)

tmp11587 := PrimCons(tmp11566, tmp11586)

tmp11588 := PrimCons(symif, tmp11587)

tmp11589 := PrimCons(tmp11588, Nil)

tmp11590 := PrimCons(tmp11564, tmp11589)

tmp11591 := PrimCons(W1679, tmp11590)

__e.Return(PrimCons(symlet, tmp11591))
return


}, 1)

tmp11592 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11560, tmp11592)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
return
}


}, 4)

tmp11625 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11559)


_ = tmp11625

tmp11626 := MakeNative(func(__e *ControlFlow) {
V1688 := __e.Get(1)
_ = V1688
V1689 := __e.Get(2)
_ = V1689
V1690 := __e.Get(3)
_ = V1690
V1691 := __e.Get(4)
_ = V1691
tmp11672 := PrimIsPair(V1688)

var ifres11669 Obj

if True == tmp11672 {
tmp11671 := PrimIsPair(V1689)

var ifres11670 Obj

if True == tmp11671 {
ifres11670 = True


} else {
ifres11670 = False


}

ifres11669 = ifres11670


} else {
ifres11669 = False


}

if True == ifres11669 {
tmp11627 := MakeNative(func(__e *ControlFlow) {
W1692 := __e.Get(1)
_ = W1692
tmp11628 := MakeNative(func(__e *ControlFlow) {
W1693 := __e.Get(1)
_ = W1693
tmp11629 := PrimHead(V1689)

tmp11630 := PrimCons(V1690, Nil)

tmp11631 := PrimCons(tmp11629, tmp11630)

tmp11632 := PrimCons(symshen_4lazyderef, tmp11631)

tmp11633 := PrimTail(V1688)

tmp11634 := PrimTail(V1689)

tmp11635 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11633, tmp11634, V1690, V1691)


tmp11636 := PrimCons(tmp11635, Nil)

tmp11637 := PrimCons(symfreeze, tmp11636)

tmp11638 := PrimHead(V1688)

tmp11639 := PrimCons(tmp11638, Nil)

tmp11640 := PrimCons(W1692, tmp11639)

tmp11641 := PrimCons(sym_a, tmp11640)

tmp11642 := PrimCons(W1693, Nil)

tmp11643 := PrimCons(symthaw, tmp11642)

tmp11644 := PrimCons(W1692, Nil)

tmp11645 := PrimCons(symshen_4pvar_2, tmp11644)

tmp11646 := PrimHead(V1688)

tmp11647 := Call(__e, PrimFunc(symshen_4demode), tmp11646)


tmp11648 := PrimCons(W1693, Nil)

tmp11649 := PrimCons(V1690, tmp11648)

tmp11650 := PrimCons(tmp11647, tmp11649)

tmp11651 := PrimCons(W1692, tmp11650)

tmp11652 := PrimCons(symshen_4bind_b, tmp11651)

tmp11653 := PrimCons(False, Nil)

tmp11654 := PrimCons(tmp11652, tmp11653)

tmp11655 := PrimCons(tmp11645, tmp11654)

tmp11656 := PrimCons(symif, tmp11655)

tmp11657 := PrimCons(tmp11656, Nil)

tmp11658 := PrimCons(tmp11643, tmp11657)

tmp11659 := PrimCons(tmp11641, tmp11658)

tmp11660 := PrimCons(symif, tmp11659)

tmp11661 := PrimCons(tmp11660, Nil)

tmp11662 := PrimCons(tmp11637, tmp11661)

tmp11663 := PrimCons(W1693, tmp11662)

tmp11664 := PrimCons(tmp11632, tmp11663)

tmp11665 := PrimCons(W1692, tmp11664)

__e.Return(PrimCons(symlet, tmp11665))
return


}, 1)

tmp11666 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11628, tmp11666)
return


}, 1)

tmp11667 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11627, tmp11667)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
return
}


}, 4)

tmp11673 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11626)


_ = tmp11673

tmp11674 := MakeNative(func(__e *ControlFlow) {
V1702 := __e.Get(1)
_ = V1702
V1703 := __e.Get(2)
_ = V1703
V1704 := __e.Get(3)
_ = V1704
V1705 := __e.Get(4)
_ = V1705
tmp11770 := PrimIsPair(V1702)

var ifres11740 Obj

if True == tmp11770 {
tmp11768 := PrimHead(V1702)

tmp11769 := PrimIsPair(tmp11768)

var ifres11742 Obj

if True == tmp11769 {
tmp11765 := PrimHead(V1702)

tmp11766 := PrimHead(tmp11765)

tmp11767 := PrimEqual(symcons, tmp11766)

var ifres11744 Obj

if True == tmp11767 {
tmp11762 := PrimHead(V1702)

tmp11763 := PrimTail(tmp11762)

tmp11764 := PrimIsPair(tmp11763)

var ifres11746 Obj

if True == tmp11764 {
tmp11758 := PrimHead(V1702)

tmp11759 := PrimTail(tmp11758)

tmp11760 := PrimTail(tmp11759)

tmp11761 := PrimIsPair(tmp11760)

var ifres11748 Obj

if True == tmp11761 {
tmp11753 := PrimHead(V1702)

tmp11754 := PrimTail(tmp11753)

tmp11755 := PrimTail(tmp11754)

tmp11756 := PrimTail(tmp11755)

tmp11757 := PrimEqual(Nil, tmp11756)

var ifres11750 Obj

if True == tmp11757 {
tmp11752 := PrimIsPair(V1703)

var ifres11751 Obj

if True == tmp11752 {
ifres11751 = True


} else {
ifres11751 = False


}

ifres11750 = ifres11751


} else {
ifres11750 = False


}

var ifres11749 Obj

if True == ifres11750 {
ifres11749 = True


} else {
ifres11749 = False


}

ifres11748 = ifres11749


} else {
ifres11748 = False


}

var ifres11747 Obj

if True == ifres11748 {
ifres11747 = True


} else {
ifres11747 = False


}

ifres11746 = ifres11747


} else {
ifres11746 = False


}

var ifres11745 Obj

if True == ifres11746 {
ifres11745 = True


} else {
ifres11745 = False


}

ifres11744 = ifres11745


} else {
ifres11744 = False


}

var ifres11743 Obj

if True == ifres11744 {
ifres11743 = True


} else {
ifres11743 = False


}

ifres11742 = ifres11743


} else {
ifres11742 = False


}

var ifres11741 Obj

if True == ifres11742 {
ifres11741 = True


} else {
ifres11741 = False


}

ifres11740 = ifres11741


} else {
ifres11740 = False


}

if True == ifres11740 {
tmp11675 := MakeNative(func(__e *ControlFlow) {
W1706 := __e.Get(1)
_ = W1706
tmp11676 := MakeNative(func(__e *ControlFlow) {
W1707 := __e.Get(1)
_ = W1707
tmp11677 := MakeNative(func(__e *ControlFlow) {
W1708 := __e.Get(1)
_ = W1708
tmp11678 := MakeNative(func(__e *ControlFlow) {
W1709 := __e.Get(1)
_ = W1709
tmp11679 := MakeNative(func(__e *ControlFlow) {
W1710 := __e.Get(1)
_ = W1710
tmp11680 := PrimHead(V1703)

tmp11681 := PrimCons(V1704, Nil)

tmp11682 := PrimCons(tmp11680, tmp11681)

tmp11683 := PrimCons(symshen_4lazyderef, tmp11682)

tmp11684 := PrimTail(V1702)

tmp11685 := PrimTail(V1703)

tmp11686 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11684, tmp11685, V1704, V1705)


tmp11687 := Call(__e, PrimFunc(symshen_4goto), W1708, tmp11686)


tmp11688 := PrimCons(W1706, Nil)

tmp11689 := PrimCons(symcons_2, tmp11688)

tmp11690 := PrimHead(V1702)

tmp11691 := PrimTail(tmp11690)

tmp11692 := PrimCons(W1706, Nil)

tmp11693 := PrimCons(symhd, tmp11692)

tmp11694 := PrimCons(W1706, Nil)

tmp11695 := PrimCons(symtl, tmp11694)

tmp11696 := PrimCons(tmp11695, Nil)

tmp11697 := PrimCons(tmp11693, tmp11696)

tmp11698 := Call(__e, PrimFunc(symshen_4invoke), W1707, W1708)


tmp11699 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11691, tmp11697, V1704, tmp11698)


tmp11700 := PrimCons(W1706, Nil)

tmp11701 := PrimCons(symshen_4pvar_2, tmp11700)

tmp11702 := Call(__e, PrimFunc(symshen_4demode), W1709)


tmp11703 := Call(__e, PrimFunc(symshen_4invoke), W1707, W1708)


tmp11704 := PrimCons(tmp11703, Nil)

tmp11705 := PrimCons(symfreeze, tmp11704)

tmp11706 := PrimCons(tmp11705, Nil)

tmp11707 := PrimCons(V1704, tmp11706)

tmp11708 := PrimCons(tmp11702, tmp11707)

tmp11709 := PrimCons(W1706, tmp11708)

tmp11710 := PrimCons(symshen_4bind_b, tmp11709)

tmp11711 := Call(__e, PrimFunc(symshen_4stpart), W1710, tmp11710, V1704)


tmp11712 := PrimCons(False, Nil)

tmp11713 := PrimCons(tmp11711, tmp11712)

tmp11714 := PrimCons(tmp11701, tmp11713)

tmp11715 := PrimCons(symif, tmp11714)

tmp11716 := PrimCons(tmp11715, Nil)

tmp11717 := PrimCons(tmp11699, tmp11716)

tmp11718 := PrimCons(tmp11689, tmp11717)

tmp11719 := PrimCons(symif, tmp11718)

tmp11720 := PrimCons(tmp11719, Nil)

tmp11721 := PrimCons(tmp11687, tmp11720)

tmp11722 := PrimCons(W1707, tmp11721)

tmp11723 := PrimCons(tmp11683, tmp11722)

tmp11724 := PrimCons(W1706, tmp11723)

__e.Return(PrimCons(symlet, tmp11724))
return


}, 1)

tmp11725 := Call(__e, PrimFunc(symshen_4extract_1vars), W1709)


__e.TailApply(tmp11679, tmp11725)
return


}, 1)

tmp11726 := PrimHead(V1702)

tmp11727 := Call(__e, PrimFunc(symshen_4tame), tmp11726)


__e.TailApply(tmp11678, tmp11727)
return


}, 1)

tmp11728 := PrimHead(V1702)

tmp11729 := PrimTail(tmp11728)

tmp11730 := PrimHead(tmp11729)

tmp11731 := PrimHead(V1702)

tmp11732 := PrimTail(tmp11731)

tmp11733 := PrimTail(tmp11732)

tmp11734 := PrimHead(tmp11733)

tmp11735 := PrimCons(tmp11730, tmp11734)

tmp11736 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11735)


__e.TailApply(tmp11677, tmp11736)
return


}, 1)

tmp11737 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11676, tmp11737)
return


}, 1)

tmp11738 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11675, tmp11738)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
return
}


}, 4)

tmp11771 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11674)


_ = tmp11771

tmp11772 := MakeNative(func(__e *ControlFlow) {
V1711 := __e.Get(1)
_ = V1711
tmp11809 := PrimIsPair(V1711)

var ifres11796 Obj

if True == tmp11809 {
tmp11807 := PrimHead(V1711)

tmp11808 := PrimEqual(symshen_4_7m, tmp11807)

var ifres11798 Obj

if True == tmp11808 {
tmp11805 := PrimTail(V1711)

tmp11806 := PrimIsPair(tmp11805)

var ifres11800 Obj

if True == tmp11806 {
tmp11802 := PrimTail(V1711)

tmp11803 := PrimTail(tmp11802)

tmp11804 := PrimEqual(Nil, tmp11803)

var ifres11801 Obj

if True == tmp11804 {
ifres11801 = True


} else {
ifres11801 = False


}

ifres11800 = ifres11801


} else {
ifres11800 = False


}

var ifres11799 Obj

if True == ifres11800 {
ifres11799 = True


} else {
ifres11799 = False


}

ifres11798 = ifres11799


} else {
ifres11798 = False


}

var ifres11797 Obj

if True == ifres11798 {
ifres11797 = True


} else {
ifres11797 = False


}

ifres11796 = ifres11797


} else {
ifres11796 = False


}

if True == ifres11796 {
tmp11773 := PrimTail(V1711)

tmp11774 := PrimHead(tmp11773)

__e.TailApply(PrimFunc(symshen_4demode), tmp11774)
return


} else {
tmp11794 := PrimIsPair(V1711)

var ifres11781 Obj

if True == tmp11794 {
tmp11792 := PrimHead(V1711)

tmp11793 := PrimEqual(symshen_4_1m, tmp11792)

var ifres11783 Obj

if True == tmp11793 {
tmp11790 := PrimTail(V1711)

tmp11791 := PrimIsPair(tmp11790)

var ifres11785 Obj

if True == tmp11791 {
tmp11787 := PrimTail(V1711)

tmp11788 := PrimTail(tmp11787)

tmp11789 := PrimEqual(Nil, tmp11788)

var ifres11786 Obj

if True == tmp11789 {
ifres11786 = True


} else {
ifres11786 = False


}

ifres11785 = ifres11786


} else {
ifres11785 = False


}

var ifres11784 Obj

if True == ifres11785 {
ifres11784 = True


} else {
ifres11784 = False


}

ifres11783 = ifres11784


} else {
ifres11783 = False


}

var ifres11782 Obj

if True == ifres11783 {
ifres11782 = True


} else {
ifres11782 = False


}

ifres11781 = ifres11782


} else {
ifres11781 = False


}

if True == ifres11781 {
tmp11775 := PrimTail(V1711)

tmp11776 := PrimHead(tmp11775)

__e.TailApply(PrimFunc(symshen_4demode), tmp11776)
return


} else {
tmp11779 := PrimIsPair(V1711)

if True == tmp11779 {
tmp11777 := MakeNative(func(__e *ControlFlow) {
Z1712 := __e.Get(1)
_ = Z1712
__e.TailApply(PrimFunc(symshen_4demode), Z1712)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11777, V1711)
return


} else {
__e.Return(V1711)
return
}


}


}


}, 1)

tmp11810 := Call(__e, ns2_1set, symshen_4demode, tmp11772)


_ = tmp11810

tmp11811 := MakeNative(func(__e *ControlFlow) {
V1713 := __e.Get(1)
_ = V1713
tmp11816 := Call(__e, PrimFunc(symshen_4wildcard_2), V1713)


if True == tmp11816 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp11814 := PrimIsPair(V1713)

if True == tmp11814 {
tmp11812 := MakeNative(func(__e *ControlFlow) {
Z1714 := __e.Get(1)
_ = Z1714
__e.TailApply(PrimFunc(symshen_4tame), Z1714)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11812, V1713)
return


} else {
__e.Return(V1713)
return
}


}


}, 1)

tmp11817 := Call(__e, ns2_1set, symshen_4tame, tmp11811)


_ = tmp11817

tmp11818 := MakeNative(func(__e *ControlFlow) {
V1715 := __e.Get(1)
_ = V1715
V1716 := __e.Get(2)
_ = V1716
tmp11821 := PrimEqual(Nil, V1715)

if True == tmp11821 {
tmp11819 := PrimCons(V1716, Nil)

__e.Return(PrimCons(symfreeze, tmp11819))
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V1715, V1716)
return
}


}, 2)

tmp11822 := Call(__e, ns2_1set, symshen_4goto, tmp11818)


_ = tmp11822

tmp11823 := MakeNative(func(__e *ControlFlow) {
V1717 := __e.Get(1)
_ = V1717
V1718 := __e.Get(2)
_ = V1718
tmp11832 := PrimEqual(Nil, V1717)

if True == tmp11832 {
__e.Return(V1718)
return
} else {
tmp11830 := PrimIsPair(V1717)

if True == tmp11830 {
tmp11824 := PrimHead(V1717)

tmp11825 := PrimTail(V1717)

tmp11826 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11825, V1718)


tmp11827 := PrimCons(tmp11826, Nil)

tmp11828 := PrimCons(tmp11824, tmp11827)

__e.Return(PrimCons(symlambda, tmp11828))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.goto-h")))
return
}


}


}, 2)

tmp11833 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11823)


_ = tmp11833

tmp11834 := MakeNative(func(__e *ControlFlow) {
V1719 := __e.Get(1)
_ = V1719
V1720 := __e.Get(2)
_ = V1720
tmp11837 := PrimEqual(Nil, V1720)

if True == tmp11837 {
tmp11835 := PrimCons(V1719, Nil)

__e.Return(PrimCons(symthaw, tmp11835))
return


} else {
__e.Return(PrimCons(V1719, V1720))
return
}


}, 2)

tmp11838 := Call(__e, ns2_1set, symshen_4invoke, tmp11834)


_ = tmp11838

tmp11839 := MakeNative(func(__e *ControlFlow) {
V1721 := __e.Get(1)
_ = V1721
__e.Return(PrimEqual(V1721, sym__))
return
}, 1)

tmp11840 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11839)


_ = tmp11840

tmp11841 := MakeNative(func(__e *ControlFlow) {
V1722 := __e.Get(1)
_ = V1722
tmp11842 := MakeNative(func(__e *ControlFlow) {
tmp11847 := PrimIsVector(V1722)

if True == tmp11847 {
tmp11844 := PrimVectorGet(V1722, MakeNumber(0))

tmp11845 := PrimEqual(tmp11844, symshen_4pvar)

if True == tmp11845 {
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

tmp11848 := MakeNative(func(__e *ControlFlow) {
Z1723 := __e.Get(1)
_ = Z1723
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp11842, tmp11848)
return


}, 1)

tmp11849 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11841)


_ = tmp11849

tmp11850 := MakeNative(func(__e *ControlFlow) {
V1724 := __e.Get(1)
_ = V1724
V1725 := __e.Get(2)
_ = V1725
tmp11857 := Call(__e, PrimFunc(symshen_4pvar_2), V1724)


if True == tmp11857 {
tmp11851 := MakeNative(func(__e *ControlFlow) {
W1726 := __e.Get(1)
_ = W1726
tmp11853 := PrimEqual(W1726, symshen_4_1null_1)

if True == tmp11853 {
__e.Return(V1724)
return
} else {
__e.TailApply(PrimFunc(symshen_4lazyderef), W1726, V1725)
return
}


}, 1)

tmp11854 := PrimVectorGet(V1724, MakeNumber(1))

tmp11855 := PrimVectorGet(V1725, tmp11854)

__e.TailApply(tmp11851, tmp11855)
return


} else {
__e.Return(V1724)
return
}


}, 2)

tmp11858 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11850)


_ = tmp11858

tmp11859 := MakeNative(func(__e *ControlFlow) {
V1727 := __e.Get(1)
_ = V1727
V1728 := __e.Get(2)
_ = V1728
tmp11872 := PrimIsPair(V1727)

if True == tmp11872 {
tmp11860 := PrimHead(V1727)

tmp11861 := Call(__e, PrimFunc(symshen_4deref), tmp11860, V1728)


tmp11862 := PrimTail(V1727)

tmp11863 := Call(__e, PrimFunc(symshen_4deref), tmp11862, V1728)


__e.Return(PrimCons(tmp11861, tmp11863))
return


} else {
tmp11870 := Call(__e, PrimFunc(symshen_4pvar_2), V1727)


if True == tmp11870 {
tmp11864 := MakeNative(func(__e *ControlFlow) {
W1729 := __e.Get(1)
_ = W1729
tmp11866 := PrimEqual(W1729, symshen_4_1null_1)

if True == tmp11866 {
__e.Return(V1727)
return
} else {
__e.TailApply(PrimFunc(symshen_4deref), W1729, V1728)
return
}


}, 1)

tmp11867 := PrimVectorGet(V1727, MakeNumber(1))

tmp11868 := PrimVectorGet(V1728, tmp11867)

__e.TailApply(tmp11864, tmp11868)
return


} else {
__e.Return(V1727)
return
}


}


}, 2)

tmp11873 := Call(__e, ns2_1set, symshen_4deref, tmp11859)


_ = tmp11873

tmp11874 := MakeNative(func(__e *ControlFlow) {
V1730 := __e.Get(1)
_ = V1730
V1731 := __e.Get(2)
_ = V1731
V1732 := __e.Get(3)
_ = V1732
V1733 := __e.Get(4)
_ = V1733
tmp11875 := MakeNative(func(__e *ControlFlow) {
W1734 := __e.Get(1)
_ = W1734
tmp11876 := MakeNative(func(__e *ControlFlow) {
W1735 := __e.Get(1)
_ = W1735
tmp11878 := PrimEqual(W1735, False)

if True == tmp11878 {
__e.TailApply(PrimFunc(symshen_4unwind), V1730, V1732, W1735)
return
} else {
__e.Return(W1735)
return
}


}, 1)

tmp11879 := Call(__e, PrimFunc(symthaw), V1733)


__e.TailApply(tmp11876, tmp11879)
return


}, 1)

tmp11880 := Call(__e, PrimFunc(symshen_4bindv), V1730, V1731, V1732)


__e.TailApply(tmp11875, tmp11880)
return


}, 4)

tmp11881 := Call(__e, ns2_1set, symshen_4bind_b, tmp11874)


_ = tmp11881

tmp11882 := MakeNative(func(__e *ControlFlow) {
V1736 := __e.Get(1)
_ = V1736
V1737 := __e.Get(2)
_ = V1737
V1738 := __e.Get(3)
_ = V1738
tmp11883 := PrimVectorGet(V1736, MakeNumber(1))

__e.Return(PrimVectorSet(V1738, tmp11883, V1737))
return


}, 3)

tmp11884 := Call(__e, ns2_1set, symshen_4bindv, tmp11882)


_ = tmp11884

tmp11885 := MakeNative(func(__e *ControlFlow) {
V1739 := __e.Get(1)
_ = V1739
V1740 := __e.Get(2)
_ = V1740
V1741 := __e.Get(3)
_ = V1741
tmp11886 := PrimVectorGet(V1739, MakeNumber(1))

tmp11887 := PrimVectorSet(V1740, tmp11886, symshen_4_1null_1)

_ = tmp11887

__e.Return(V1741)
return


}, 3)

tmp11888 := Call(__e, ns2_1set, symshen_4unwind, tmp11885)


_ = tmp11888

tmp11889 := MakeNative(func(__e *ControlFlow) {
V1750 := __e.Get(1)
_ = V1750
V1751 := __e.Get(2)
_ = V1751
V1752 := __e.Get(3)
_ = V1752
tmp11904 := PrimEqual(Nil, V1750)

if True == tmp11904 {
__e.Return(V1751)
return
} else {
tmp11902 := PrimIsPair(V1750)

if True == tmp11902 {
tmp11890 := PrimHead(V1750)

tmp11891 := PrimCons(V1752, Nil)

tmp11892 := PrimCons(symshen_4newpv, tmp11891)

tmp11893 := PrimTail(V1750)

tmp11894 := Call(__e, PrimFunc(symshen_4stpart), tmp11893, V1751, V1752)


tmp11895 := PrimCons(tmp11894, Nil)

tmp11896 := PrimCons(V1752, tmp11895)

tmp11897 := PrimCons(symshen_4gc, tmp11896)

tmp11898 := PrimCons(tmp11897, Nil)

tmp11899 := PrimCons(tmp11892, tmp11898)

tmp11900 := PrimCons(tmp11890, tmp11899)

__e.Return(PrimCons(symlet, tmp11900))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
return
}


}


}, 3)

tmp11905 := Call(__e, ns2_1set, symshen_4stpart, tmp11889)


_ = tmp11905

tmp11906 := MakeNative(func(__e *ControlFlow) {
V1753 := __e.Get(1)
_ = V1753
V1754 := __e.Get(2)
_ = V1754
tmp11911 := PrimEqual(V1754, False)

if True == tmp11911 {
tmp11907 := MakeNative(func(__e *ControlFlow) {
W1755 := __e.Get(1)
_ = W1755
tmp11908 := Call(__e, PrimFunc(symshen_4decrement_1ticket), W1755, V1753)


_ = tmp11908

__e.Return(V1754)
return


}, 1)

tmp11909 := Call(__e, PrimFunc(symshen_4ticket_1number), V1753)


__e.TailApply(tmp11907, tmp11909)
return


} else {
__e.Return(V1754)
return
}


}, 2)

tmp11912 := Call(__e, ns2_1set, symshen_4gc, tmp11906)


_ = tmp11912

tmp11913 := MakeNative(func(__e *ControlFlow) {
V1756 := __e.Get(1)
_ = V1756
V1757 := __e.Get(2)
_ = V1757
tmp11914 := PrimNumberSubtract(V1756, MakeNumber(1))

__e.Return(PrimVectorSet(V1757, MakeNumber(1), tmp11914))
return


}, 2)

tmp11915 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp11913)


_ = tmp11915

tmp11916 := MakeNative(func(__e *ControlFlow) {
V1758 := __e.Get(1)
_ = V1758
tmp11917 := MakeNative(func(__e *ControlFlow) {
W1759 := __e.Get(1)
_ = W1759
tmp11918 := MakeNative(func(__e *ControlFlow) {
W1760 := __e.Get(1)
_ = W1760
tmp11919 := MakeNative(func(__e *ControlFlow) {
W1761 := __e.Get(1)
_ = W1761
__e.Return(W1760)
return
}, 1)

tmp11920 := Call(__e, PrimFunc(symshen_4nextticket), V1758, W1759)


__e.TailApply(tmp11919, tmp11920)
return


}, 1)

tmp11921 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), W1759)


__e.TailApply(tmp11918, tmp11921)
return


}, 1)

tmp11922 := Call(__e, PrimFunc(symshen_4ticket_1number), V1758)


__e.TailApply(tmp11917, tmp11922)
return


}, 1)

tmp11923 := Call(__e, ns2_1set, symshen_4newpv, tmp11916)


_ = tmp11923

tmp11924 := MakeNative(func(__e *ControlFlow) {
V1762 := __e.Get(1)
_ = V1762
__e.Return(PrimVectorGet(V1762, MakeNumber(1)))
return
}, 1)

tmp11925 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp11924)


_ = tmp11925

tmp11926 := MakeNative(func(__e *ControlFlow) {
V1763 := __e.Get(1)
_ = V1763
V1764 := __e.Get(2)
_ = V1764
tmp11927 := MakeNative(func(__e *ControlFlow) {
W1765 := __e.Get(1)
_ = W1765
tmp11928 := PrimNumberAdd(V1764, MakeNumber(1))

__e.Return(PrimVectorSet(W1765, MakeNumber(1), tmp11928))
return


}, 1)

tmp11929 := PrimVectorSet(V1763, V1764, symshen_4_1null_1)

__e.TailApply(tmp11927, tmp11929)
return


}, 2)

tmp11930 := Call(__e, ns2_1set, symshen_4nextticket, tmp11926)


_ = tmp11930

tmp11931 := MakeNative(func(__e *ControlFlow) {
V1766 := __e.Get(1)
_ = V1766
tmp11932 := PrimAbsvector(MakeNumber(2))

tmp11933 := PrimVectorSet(tmp11932, MakeNumber(0), symshen_4pvar)

__e.Return(PrimVectorSet(tmp11933, MakeNumber(1), V1766))
return


}, 1)

tmp11934 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11931)


_ = tmp11934

tmp11935 := MakeNative(func(__e *ControlFlow) {
V1767 := __e.Get(1)
_ = V1767
tmp11936 := PrimVectorGet(V1767, MakeNumber(1))

tmp11937 := Call(__e, PrimFunc(symshen_4app), tmp11936, MakeString(""), symshen_4a)


__e.Return(PrimStringConcat(MakeString("Var"), tmp11937))
return


}, 1)

tmp11938 := Call(__e, ns2_1set, symshen_4pvar, tmp11935)


_ = tmp11938

tmp11939 := MakeNative(func(__e *ControlFlow) {
tmp11940 := PrimValue(symshen_4_dinfs_d)

tmp11941 := PrimNumberAdd(MakeNumber(1), tmp11940)

__e.Return(PrimSet(symshen_4_dinfs_d, tmp11941))
return


}, 0)

tmp11942 := Call(__e, ns2_1set, symshen_4incinfs, tmp11939)


_ = tmp11942

tmp11943 := MakeNative(func(__e *ControlFlow) {
V1768 := __e.Get(1)
_ = V1768
tmp11950 := PrimIsInteger(V1768)

var ifres11947 Obj

if True == tmp11950 {
tmp11949 := PrimGreatThan(V1768, MakeNumber(0))

var ifres11948 Obj

if True == tmp11949 {
ifres11948 = True


} else {
ifres11948 = False


}

ifres11947 = ifres11948


} else {
ifres11947 = False


}

if True == ifres11947 {
__e.Return(PrimSet(symshen_4_dsize_1prolog_1vector_d, V1768))
return
} else {
tmp11944 := Call(__e, PrimFunc(symshen_4app), V1768, MakeString(""), symshen_4a)


tmp11945 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp11944)

__e.Return(PrimSimpleError(tmp11945))
return


}


}, 1)

tmp11951 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11943)


_ = tmp11951

tmp11952 := MakeNative(func(__e *ControlFlow) {
V1780 := __e.Get(1)
_ = V1780
V1781 := __e.Get(2)
_ = V1781
V1782 := __e.Get(3)
_ = V1782
V1783 := __e.Get(4)
_ = V1783
tmp11982 := PrimEqual(V1780, V1781)

if True == tmp11982 {
__e.TailApply(PrimFunc(symthaw), V1783)
return
} else {
tmp11980 := Call(__e, PrimFunc(symshen_4pvar_2), V1780)


var ifres11975 Obj

if True == tmp11980 {
tmp11977 := Call(__e, PrimFunc(symshen_4deref), V1781, V1782)


tmp11978 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1780, tmp11977)


tmp11979 := PrimNot(tmp11978)

var ifres11976 Obj

if True == tmp11979 {
ifres11976 = True


} else {
ifres11976 = False


}

ifres11975 = ifres11976


} else {
ifres11975 = False


}

if True == ifres11975 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1780, V1781, V1782, V1783)
return
} else {
tmp11973 := Call(__e, PrimFunc(symshen_4pvar_2), V1781)


var ifres11968 Obj

if True == tmp11973 {
tmp11970 := Call(__e, PrimFunc(symshen_4deref), V1780, V1782)


tmp11971 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1781, tmp11970)


tmp11972 := PrimNot(tmp11971)

var ifres11969 Obj

if True == tmp11972 {
ifres11969 = True


} else {
ifres11969 = False


}

ifres11968 = ifres11969


} else {
ifres11968 = False


}

if True == ifres11968 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1781, V1780, V1782, V1783)
return
} else {
tmp11966 := PrimIsPair(V1780)

var ifres11963 Obj

if True == tmp11966 {
tmp11965 := PrimIsPair(V1781)

var ifres11964 Obj

if True == tmp11965 {
ifres11964 = True


} else {
ifres11964 = False


}

ifres11963 = ifres11964


} else {
ifres11963 = False


}

if True == ifres11963 {
tmp11953 := PrimHead(V1780)

tmp11954 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11953, V1782)


tmp11955 := PrimHead(V1781)

tmp11956 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11955, V1782)


tmp11957 := MakeNative(func(__e *ControlFlow) {
tmp11958 := PrimTail(V1780)

tmp11959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11958, V1782)


tmp11960 := PrimTail(V1781)

tmp11961 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11960, V1782)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11959, tmp11961, V1782, V1783)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11954, tmp11956, V1782, tmp11957)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp11983 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp11952)


_ = tmp11983

tmp11984 := MakeNative(func(__e *ControlFlow) {
V1795 := __e.Get(1)
_ = V1795
V1796 := __e.Get(2)
_ = V1796
V1797 := __e.Get(3)
_ = V1797
V1798 := __e.Get(4)
_ = V1798
tmp12004 := PrimEqual(V1795, V1796)

if True == tmp12004 {
__e.TailApply(PrimFunc(symthaw), V1798)
return
} else {
tmp12002 := Call(__e, PrimFunc(symshen_4pvar_2), V1795)


if True == tmp12002 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1795, V1796, V1797, V1798)
return
} else {
tmp12000 := Call(__e, PrimFunc(symshen_4pvar_2), V1796)


if True == tmp12000 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1796, V1795, V1797, V1798)
return
} else {
tmp11998 := PrimIsPair(V1795)

var ifres11995 Obj

if True == tmp11998 {
tmp11997 := PrimIsPair(V1796)

var ifres11996 Obj

if True == tmp11997 {
ifres11996 = True


} else {
ifres11996 = False


}

ifres11995 = ifres11996


} else {
ifres11995 = False


}

if True == ifres11995 {
tmp11985 := PrimHead(V1795)

tmp11986 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11985, V1797)


tmp11987 := PrimHead(V1796)

tmp11988 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11987, V1797)


tmp11989 := MakeNative(func(__e *ControlFlow) {
tmp11990 := PrimTail(V1795)

tmp11991 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11990, V1797)


tmp11992 := PrimTail(V1796)

tmp11993 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11992, V1797)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11991, tmp11993, V1797, V1798)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11986, tmp11988, V1797, tmp11989)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12005 := Call(__e, ns2_1set, symshen_4lzy_a, tmp11984)


_ = tmp12005

tmp12006 := MakeNative(func(__e *ControlFlow) {
V1804 := __e.Get(1)
_ = V1804
V1805 := __e.Get(2)
_ = V1805
tmp12016 := PrimEqual(V1804, V1805)

if True == tmp12016 {
__e.Return(True)
return
} else {
tmp12014 := PrimIsPair(V1805)

if True == tmp12014 {
tmp12011 := PrimHead(V1805)

tmp12012 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1804, tmp12011)


if True == tmp12012 {
__e.Return(True)
return
} else {
tmp12008 := PrimTail(V1805)

tmp12009 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1804, tmp12008)


if True == tmp12009 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 2)

tmp12017 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12006)


_ = tmp12017

tmp12018 := MakeNative(func(__e *ControlFlow) {
V1806 := __e.Get(1)
_ = V1806
V1807 := __e.Get(2)
_ = V1807
V1808 := __e.Get(3)
_ = V1808
V1809 := __e.Get(4)
_ = V1809
V1810 := __e.Get(5)
_ = V1810
tmp12019 := Call(__e, V1806, V1807)


tmp12020 := Call(__e, tmp12019, V1808)


tmp12021 := Call(__e, tmp12020, V1809)


__e.TailApply(tmp12021, V1810)
return


}, 5)

tmp12022 := Call(__e, ns2_1set, symcall, tmp12018)


_ = tmp12022

tmp12023 := MakeNative(func(__e *ControlFlow) {
V1817 := __e.Get(1)
_ = V1817
V1818 := __e.Get(2)
_ = V1818
V1819 := __e.Get(3)
_ = V1819
V1820 := __e.Get(4)
_ = V1820
V1821 := __e.Get(5)
_ = V1821
__e.TailApply(PrimFunc(symshen_4deref), V1817, V1818)
return
}, 5)

tmp12024 := Call(__e, ns2_1set, symreturn, tmp12023)


_ = tmp12024

tmp12025 := MakeNative(func(__e *ControlFlow) {
V1828 := __e.Get(1)
_ = V1828
V1829 := __e.Get(2)
_ = V1829
V1830 := __e.Get(3)
_ = V1830
V1831 := __e.Get(4)
_ = V1831
V1832 := __e.Get(5)
_ = V1832
if True == V1828 {
__e.TailApply(PrimFunc(symthaw), V1832)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp12027 := Call(__e, ns2_1set, symwhen, tmp12025)


_ = tmp12027

tmp12028 := MakeNative(func(__e *ControlFlow) {
V1833 := __e.Get(1)
_ = V1833
V1834 := __e.Get(2)
_ = V1834
V1835 := __e.Get(3)
_ = V1835
V1836 := __e.Get(4)
_ = V1836
V1837 := __e.Get(5)
_ = V1837
V1838 := __e.Get(6)
_ = V1838
tmp12029 := Call(__e, PrimFunc(symshen_4lazyderef), V1833, V1835)


tmp12030 := Call(__e, PrimFunc(symshen_4lazyderef), V1834, V1835)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12029, tmp12030, V1835, V1838)
return


}, 6)

tmp12031 := Call(__e, ns2_1set, symis, tmp12028)


_ = tmp12031

tmp12032 := MakeNative(func(__e *ControlFlow) {
V1839 := __e.Get(1)
_ = V1839
V1840 := __e.Get(2)
_ = V1840
V1841 := __e.Get(3)
_ = V1841
V1842 := __e.Get(4)
_ = V1842
V1843 := __e.Get(5)
_ = V1843
V1844 := __e.Get(6)
_ = V1844
tmp12033 := Call(__e, PrimFunc(symshen_4lazyderef), V1839, V1841)


tmp12034 := Call(__e, PrimFunc(symshen_4lazyderef), V1840, V1841)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12033, tmp12034, V1841, V1844)
return


}, 6)

tmp12035 := Call(__e, ns2_1set, symis_b, tmp12032)


_ = tmp12035

tmp12036 := MakeNative(func(__e *ControlFlow) {
V1849 := __e.Get(1)
_ = V1849
V1850 := __e.Get(2)
_ = V1850
V1851 := __e.Get(3)
_ = V1851
V1852 := __e.Get(4)
_ = V1852
V1853 := __e.Get(5)
_ = V1853
V1854 := __e.Get(6)
_ = V1854
__e.TailApply(PrimFunc(symshen_4bind_b), V1849, V1850, V1851, V1854)
return
}, 6)

tmp12037 := Call(__e, ns2_1set, symbind, tmp12036)


_ = tmp12037

tmp12038 := MakeNative(func(__e *ControlFlow) {
V1855 := __e.Get(1)
_ = V1855
V1856 := __e.Get(2)
_ = V1856
V1857 := __e.Get(3)
_ = V1857
V1858 := __e.Get(4)
_ = V1858
V1859 := __e.Get(5)
_ = V1859
tmp12040 := Call(__e, PrimFunc(symshen_4lazyderef), V1855, V1856)


tmp12041 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12040)


if True == tmp12041 {
__e.TailApply(PrimFunc(symthaw), V1859)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12042 := Call(__e, ns2_1set, symvar_2, tmp12038)


_ = tmp12042

tmp12043 := MakeNative(func(__e *ControlFlow) {
V1862 := __e.Get(1)
_ = V1862
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12044 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12043)


_ = tmp12044

tmp12045 := MakeNative(func(__e *ControlFlow) {
V1881 := __e.Get(1)
_ = V1881
V1882 := __e.Get(2)
_ = V1882
V1883 := __e.Get(3)
_ = V1883
V1884 := __e.Get(4)
_ = V1884
V1885 := __e.Get(5)
_ = V1885
tmp12058 := PrimEqual(Nil, V1881)

if True == tmp12058 {
__e.Return(False)
return
} else {
tmp12056 := PrimIsPair(V1881)

if True == tmp12056 {
tmp12046 := MakeNative(func(__e *ControlFlow) {
W1886 := __e.Get(1)
_ = W1886
tmp12049 := PrimEqual(W1886, False)

if True == tmp12049 {
tmp12047 := PrimTail(V1881)

__e.TailApply(PrimFunc(symfork), tmp12047, V1882, V1883, V1884, V1885)
return


} else {
__e.Return(W1886)
return
}


}, 1)

tmp12050 := PrimHead(V1881)

tmp12051 := Call(__e, tmp12050, V1882)


tmp12052 := Call(__e, tmp12051, V1883)


tmp12053 := Call(__e, tmp12052, V1884)


tmp12054 := Call(__e, tmp12053, V1885)


__e.TailApply(tmp12046, tmp12054)
return


} else {
__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
return
}


}


}, 5)

tmp12059 := Call(__e, ns2_1set, symfork, tmp12045)


_ = tmp12059

tmp12060 := MakeNative(func(__e *ControlFlow) {
V1887 := __e.Get(1)
_ = V1887
V1888 := __e.Get(2)
_ = V1888
V1889 := __e.Get(3)
_ = V1889
V1890 := __e.Get(4)
_ = V1890
V1891 := __e.Get(5)
_ = V1891
V1892 := __e.Get(6)
_ = V1892
V1893 := __e.Get(7)
_ = V1893
tmp12067 := Call(__e, PrimFunc(symshen_4unlocked_2), V1891)


if True == tmp12067 {
tmp12061 := MakeNative(func(__e *ControlFlow) {
W1894 := __e.Get(1)
_ = W1894
tmp12062 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12062

tmp12063 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V1887, V1888, V1889, W1894, V1890, V1891, V1892, V1893)
return
}, 0)

tmp12064 := Call(__e, PrimFunc(symis), W1894, Nil, V1890, V1891, V1892, tmp12063)


__e.TailApply(PrimFunc(symshen_4gc), V1890, tmp12064)
return


}, 1)

tmp12065 := Call(__e, PrimFunc(symshen_4newpv), V1890)


__e.TailApply(tmp12061, tmp12065)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12068 := Call(__e, ns2_1set, symfindall, tmp12060)


_ = tmp12068

tmp12069 := MakeNative(func(__e *ControlFlow) {
V1895 := __e.Get(1)
_ = V1895
V1896 := __e.Get(2)
_ = V1896
V1897 := __e.Get(3)
_ = V1897
V1898 := __e.Get(4)
_ = V1898
V1899 := __e.Get(5)
_ = V1899
V1900 := __e.Get(6)
_ = V1900
V1901 := __e.Get(7)
_ = V1901
V1902 := __e.Get(8)
_ = V1902
tmp12070 := MakeNative(func(__e *ControlFlow) {
W1903 := __e.Get(1)
_ = W1903
tmp12075 := PrimEqual(W1903, False)

if True == tmp12075 {
tmp12073 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


if True == tmp12073 {
tmp12071 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12071

__e.TailApply(PrimFunc(symis_b), V1897, V1898, V1899, V1900, V1901, V1902)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1903)
return
}


}, 1)

tmp12080 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


var ifres12076 Obj

if True == tmp12080 {
tmp12077 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12077

tmp12078 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V1895, V1898, V1899, V1900, V1901, V1902)
return
}, 0)

tmp12079 := Call(__e, PrimFunc(symcall), V1896, V1899, V1900, V1901, tmp12078)


ifres12076 = tmp12079


} else {
ifres12076 = False


}

__e.TailApply(tmp12070, ifres12076)
return


}, 8)

tmp12081 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12069)


_ = tmp12081

tmp12082 := MakeNative(func(__e *ControlFlow) {
V1910 := __e.Get(1)
_ = V1910
V1911 := __e.Get(2)
_ = V1911
V1912 := __e.Get(3)
_ = V1912
V1913 := __e.Get(4)
_ = V1913
V1914 := __e.Get(5)
_ = V1914
V1915 := __e.Get(6)
_ = V1915
tmp12083 := Call(__e, PrimFunc(symshen_4deref), V1910, V1912)


tmp12084 := Call(__e, PrimFunc(symshen_4lazyderef), V1911, V1912)


tmp12085 := PrimCons(tmp12083, tmp12084)

tmp12086 := Call(__e, PrimFunc(symshen_4bindv), V1911, tmp12085, V1912)


_ = tmp12086

__e.Return(False)
return


}, 6)

tmp12087 := Call(__e, ns2_1set, symshen_4overbind, tmp12082)


_ = tmp12087

tmp12088 := MakeNative(func(__e *ControlFlow) {
V1918 := __e.Get(1)
_ = V1918
tmp12092 := PrimEqual(sym_7, V1918)

if True == tmp12092 {
__e.Return(PrimSet(symshen_4_doccurs_d, True))
return
} else {
tmp12090 := PrimEqual(sym_1, V1918)

if True == tmp12090 {
__e.Return(PrimSet(symshen_4_doccurs_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp12088)
return




}, 0)

