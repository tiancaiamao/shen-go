package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp16232 := MakeNative(func(__e *ControlFlow) {
V1141 := __e.Get(1)
_ = V1141
V1142 := __e.Get(2)
_ = V1142
tmp16233 := MakeNative(func(__e *ControlFlow) {
Compile := __e.Get(1)
_ = Compile
tmp16235 := Call(__e, PrimFunc(symshen_4parsed_2), Compile)


if True == tmp16235 {
__e.TailApply(PrimFunc(symshen_4objectcode), Compile)
return
} else {
__e.Return(PrimSimpleError(MakeString("parse failure\n")))
return
}


}, 1)

tmp16236 := PrimCons(symshen_4no_1action, Nil)

tmp16237 := PrimCons(V1142, tmp16236)

tmp16238 := Call(__e, V1141, tmp16237)


__e.TailApply(tmp16233, tmp16238)
return


}, 2)

tmp16239 := Call(__e, ns2_1set, symcompile, tmp16232)


_ = tmp16239

tmp16240 := MakeNative(func(__e *ControlFlow) {
V1147 := __e.Get(1)
_ = V1147
tmp16253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), V1147)


if True == tmp16253 {
__e.Return(False)
return
} else {
tmp16251 := PrimIsPair(V1147)

var ifres16247 Obj

if True == tmp16251 {
tmp16249 := PrimHead(V1147)

tmp16250 := PrimIsPair(tmp16249)

var ifres16248 Obj

if True == tmp16250 {
ifres16248 = True


} else {
ifres16248 = False


}

ifres16247 = ifres16248


} else {
ifres16247 = False


}

if True == ifres16247 {
tmp16241 := PrimHead(V1147)

tmp16242 := PrimSet(symshen_4_dresidue_d, tmp16241)

_ = tmp16242

tmp16243 := PrimHead(V1147)

tmp16244 := Call(__e, PrimFunc(symshen_4app), tmp16243, MakeString("\n ..."), symshen_4r)


tmp16245 := PrimStringConcat(MakeString("syntax error here: "), tmp16244)

__e.Return(PrimSimpleError(tmp16245))
return


} else {
__e.Return(True)
return
}


}


}, 1)

tmp16254 := Call(__e, ns2_1set, symshen_4parsed_2, tmp16240)


_ = tmp16254

tmp16255 := MakeNative(func(__e *ControlFlow) {
V1148 := __e.Get(1)
_ = V1148
tmp16256 := Call(__e, PrimFunc(symfail))


__e.Return(PrimEqual(V1148, tmp16256))
return


}, 1)

tmp16257 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp16255)


_ = tmp16257

tmp16258 := MakeNative(func(__e *ControlFlow) {
V1151 := __e.Get(1)
_ = V1151
tmp16271 := PrimIsPair(V1151)

var ifres16262 Obj

if True == tmp16271 {
tmp16269 := PrimTail(V1151)

tmp16270 := PrimIsPair(tmp16269)

var ifres16264 Obj

if True == tmp16270 {
tmp16266 := PrimTail(V1151)

tmp16267 := PrimTail(tmp16266)

tmp16268 := PrimEqual(Nil, tmp16267)

var ifres16265 Obj

if True == tmp16268 {
ifres16265 = True


} else {
ifres16265 = False


}

ifres16264 = ifres16265


} else {
ifres16264 = False


}

var ifres16263 Obj

if True == ifres16264 {
ifres16263 = True


} else {
ifres16263 = False


}

ifres16262 = ifres16263


} else {
ifres16262 = False


}

if True == ifres16262 {
tmp16259 := PrimTail(V1151)

__e.Return(PrimHead(tmp16259))
return


} else {
tmp16260 := Call(__e, PrimFunc(symshen_4app), V1151, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp16260))
return


}


}, 1)

tmp16272 := Call(__e, ns2_1set, symshen_4objectcode, tmp16258)


_ = tmp16272

tmp16273 := MakeNative(func(__e *ControlFlow) {
V1152 := __e.Get(1)
_ = V1152
tmp16274 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_5yacc_6), X)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16274, V1152)
return


}, 1)

tmp16275 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp16273)


_ = tmp16275

tmp16276 := MakeNative(func(__e *ControlFlow) {
V1153 := __e.Get(1)
_ = V1153
tmp16277 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16279 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16312 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1153)


var ifres16280 Obj

if True == tmp16312 {
tmp16281 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp16282 := MakeNative(func(__e *ControlFlow) {
News1034 := __e.Get(1)
_ = News1034
tmp16283 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5yaccsig_6 := __e.Get(1)
_ = Parseshen_4_5yaccsig_6
tmp16306 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5yaccsig_6)


if True == tmp16306 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16284 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5c_1rules_6 := __e.Get(1)
_ = Parseshen_4_5c_1rules_6
tmp16303 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5c_1rules_6)


if True == tmp16303 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16285 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5c_1rules_6)


tmp16286 := MakeNative(func(__e *ControlFlow) {
Stream := __e.Get(1)
_ = Stream
tmp16287 := MakeNative(func(__e *ControlFlow) {
Def := __e.Get(1)
_ = Def
__e.Return(Def)
return
}, 1)

tmp16288 := PrimCons(F, Nil)

tmp16289 := PrimCons(symdefine, tmp16288)

tmp16290 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5yaccsig_6)


tmp16291 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5yaccsig_6)


tmp16292 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5c_1rules_6)


tmp16293 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), tmp16291, Stream, tmp16292)


tmp16294 := PrimCons(tmp16293, Nil)

tmp16295 := PrimCons(sym_1_6, tmp16294)

tmp16296 := PrimCons(Stream, tmp16295)

tmp16297 := Call(__e, PrimFunc(symappend), tmp16290, tmp16296)


tmp16298 := Call(__e, PrimFunc(symappend), tmp16289, tmp16297)


__e.TailApply(tmp16287, tmp16298)
return


}, 1)

tmp16299 := Call(__e, PrimFunc(symprotect), symS)


tmp16300 := Call(__e, PrimFunc(symgensym), tmp16299)


tmp16301 := Call(__e, tmp16286, tmp16300)


__e.TailApply(PrimFunc(symshen_4comb), tmp16285, tmp16301)
return


}


}, 1)

tmp16304 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), Parseshen_4_5yaccsig_6)


__e.TailApply(tmp16284, tmp16304)
return


}


}, 1)

tmp16307 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), News1034)


__e.TailApply(tmp16283, tmp16307)
return


}, 1)

tmp16308 := Call(__e, PrimFunc(symshen_4tls), V1153)


__e.TailApply(tmp16282, tmp16308)
return


}, 1)

tmp16309 := Call(__e, PrimFunc(symshen_4hds), V1153)


tmp16310 := Call(__e, tmp16281, tmp16309)


ifres16280 = tmp16310


} else {
tmp16311 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16280 = tmp16311


}

__e.TailApply(tmp16277, ifres16280)
return


}, 1)

tmp16313 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp16276)


_ = tmp16313

tmp16314 := MakeNative(func(__e *ControlFlow) {
V1154 := __e.Get(1)
_ = V1154
tmp16315 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16326 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16326 {
tmp16316 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16318 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16318 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16319 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp16322 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp16322 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16320 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16320, Nil)
return


}


}, 1)

tmp16323 := Call(__e, PrimFunc(sym_5e_6), V1154)


tmp16324 := Call(__e, tmp16319, tmp16323)


__e.TailApply(tmp16316, tmp16324)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16397 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1154)


var ifres16327 Obj

if True == tmp16397 {
tmp16328 := MakeNative(func(__e *ControlFlow) {
LC := __e.Get(1)
_ = LC
tmp16329 := MakeNative(func(__e *ControlFlow) {
News1036 := __e.Get(1)
_ = News1036
tmp16392 := Call(__e, PrimFunc(symshen_4ccons_2), News1036)


if True == tmp16392 {
tmp16330 := MakeNative(func(__e *ControlFlow) {
SynCons := __e.Get(1)
_ = SynCons
tmp16387 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, symlist)


if True == tmp16387 {
tmp16331 := MakeNative(func(__e *ControlFlow) {
News1037 := __e.Get(1)
_ = News1037
tmp16384 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1037)


if True == tmp16384 {
tmp16332 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp16333 := MakeNative(func(__e *ControlFlow) {
News1038 := __e.Get(1)
_ = News1038
tmp16334 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5end_6 := __e.Get(1)
_ = Parseshen_4_5end_6
tmp16379 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)


if True == tmp16379 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16376 := Call(__e, PrimFunc(symshen_4tlstream), News1036)


tmp16377 := Call(__e, PrimFunc(symshen_4_ahd_2), tmp16376, sym_a_a_6)


if True == tmp16377 {
tmp16335 := MakeNative(func(__e *ControlFlow) {
News1039 := __e.Get(1)
_ = News1039
tmp16372 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1039)


if True == tmp16372 {
tmp16336 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp16337 := MakeNative(func(__e *ControlFlow) {
News1040 := __e.Get(1)
_ = News1040
tmp16368 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1040)


if True == tmp16368 {
tmp16338 := MakeNative(func(__e *ControlFlow) {
RC := __e.Get(1)
_ = RC
tmp16339 := MakeNative(func(__e *ControlFlow) {
News1041 := __e.Get(1)
_ = News1041
tmp16364 := PrimEqual(sym_i, LC)

var ifres16361 Obj

if True == tmp16364 {
tmp16363 := PrimEqual(sym_j, RC)

var ifres16362 Obj

if True == tmp16363 {
ifres16362 = True


} else {
ifres16362 = False


}

ifres16361 = ifres16362


} else {
ifres16361 = False


}

if True == ifres16361 {
tmp16340 := Call(__e, PrimFunc(symshen_4in_1_6), News1041)


tmp16341 := MakeNative(func(__e *ControlFlow) {
C := __e.Get(1)
_ = C
tmp16342 := PrimCons(A, Nil)

tmp16343 := PrimCons(symlist, tmp16342)

tmp16344 := Call(__e, PrimFunc(symprotect), C)


tmp16345 := PrimCons(tmp16344, Nil)

tmp16346 := PrimCons(tmp16343, tmp16345)

tmp16347 := PrimCons(symstr, tmp16346)

tmp16348 := PrimCons(A, Nil)

tmp16349 := PrimCons(symlist, tmp16348)

tmp16350 := PrimCons(B, Nil)

tmp16351 := PrimCons(tmp16349, tmp16350)

tmp16352 := PrimCons(symstr, tmp16351)

tmp16353 := PrimCons(sym_j, Nil)

tmp16354 := PrimCons(tmp16352, tmp16353)

tmp16355 := PrimCons(sym_1_1_6, tmp16354)

tmp16356 := PrimCons(tmp16347, tmp16355)

__e.Return(PrimCons(sym_i, tmp16356))
return


}, 1)

tmp16357 := Call(__e, PrimFunc(symgensym), symC)


tmp16358 := Call(__e, PrimFunc(symprotect), tmp16357)


tmp16359 := Call(__e, tmp16341, tmp16358)


__e.TailApply(PrimFunc(symshen_4comb), tmp16340, tmp16359)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16365 := Call(__e, PrimFunc(symshen_4tls), News1040)


__e.TailApply(tmp16339, tmp16365)
return


}, 1)

tmp16366 := Call(__e, PrimFunc(symshen_4hds), News1040)


__e.TailApply(tmp16338, tmp16366)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16369 := Call(__e, PrimFunc(symshen_4tls), News1039)


__e.TailApply(tmp16337, tmp16369)
return


}, 1)

tmp16370 := Call(__e, PrimFunc(symshen_4hds), News1039)


__e.TailApply(tmp16336, tmp16370)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16373 := Call(__e, PrimFunc(symshen_4tlstream), News1036)


tmp16374 := Call(__e, PrimFunc(symshen_4tls), tmp16373)


__e.TailApply(tmp16335, tmp16374)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp16380 := Call(__e, PrimFunc(symshen_4_5end_6), News1038)


__e.TailApply(tmp16334, tmp16380)
return


}, 1)

tmp16381 := Call(__e, PrimFunc(symshen_4tls), News1037)


__e.TailApply(tmp16333, tmp16381)
return


}, 1)

tmp16382 := Call(__e, PrimFunc(symshen_4hds), News1037)


__e.TailApply(tmp16332, tmp16382)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16385 := Call(__e, PrimFunc(symshen_4tls), SynCons)


__e.TailApply(tmp16331, tmp16385)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16388 := Call(__e, PrimFunc(symshen_4hds), News1036)


tmp16389 := Call(__e, PrimFunc(symshen_4_5_1out), News1036)


tmp16390 := Call(__e, PrimFunc(symshen_4comb), tmp16388, tmp16389)


__e.TailApply(tmp16330, tmp16390)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16393 := Call(__e, PrimFunc(symshen_4tls), V1154)


__e.TailApply(tmp16329, tmp16393)
return


}, 1)

tmp16394 := Call(__e, PrimFunc(symshen_4hds), V1154)


tmp16395 := Call(__e, tmp16328, tmp16394)


ifres16327 = tmp16395


} else {
tmp16396 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16327 = tmp16396


}

__e.TailApply(tmp16315, ifres16327)
return


}, 1)

tmp16398 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp16314)


_ = tmp16398

tmp16399 := MakeNative(func(__e *ControlFlow) {
V1155 := __e.Get(1)
_ = V1155
tmp16400 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16418 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16418 {
tmp16401 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16403 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16404 := MakeNative(func(__e *ControlFlow) {
Parse_5_b_6 := __e.Get(1)
_ = Parse_5_b_6
tmp16414 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5_b_6)


if True == tmp16414 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16405 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5_b_6)


tmp16411 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp16412 := Call(__e, PrimFunc(symempty_2), tmp16411)


var ifres16406 Obj

if True == tmp16412 {
ifres16406 = Nil


} else {
tmp16407 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp16408 := Call(__e, PrimFunc(symshen_4app), tmp16407, MakeString("\n ..."), symshen_4r)


tmp16409 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp16408)

tmp16410 := PrimSimpleError(tmp16409)

ifres16406 = tmp16410


}

__e.TailApply(PrimFunc(symshen_4comb), tmp16405, ifres16406)
return


}


}, 1)

tmp16415 := Call(__e, PrimFunc(sym_5_b_6), V1155)


tmp16416 := Call(__e, tmp16404, tmp16415)


__e.TailApply(tmp16401, tmp16416)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16419 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5c_1rule_6 := __e.Get(1)
_ = Parseshen_4_5c_1rule_6
tmp16429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5c_1rule_6)


if True == tmp16429 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16420 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5c_1rules_6 := __e.Get(1)
_ = Parseshen_4_5c_1rules_6
tmp16426 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5c_1rules_6)


if True == tmp16426 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16421 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5c_1rules_6)


tmp16422 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5c_1rule_6)


tmp16423 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5c_1rules_6)


tmp16424 := PrimCons(tmp16422, tmp16423)

__e.TailApply(PrimFunc(symshen_4comb), tmp16421, tmp16424)
return


}


}, 1)

tmp16427 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), Parseshen_4_5c_1rule_6)


__e.TailApply(tmp16420, tmp16427)
return


}


}, 1)

tmp16430 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V1155)


tmp16431 := Call(__e, tmp16419, tmp16430)


__e.TailApply(tmp16400, tmp16431)
return


}, 1)

tmp16432 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp16399)


_ = tmp16432

tmp16433 := MakeNative(func(__e *ControlFlow) {
V1156 := __e.Get(1)
_ = V1156
tmp16434 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16454 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16454 {
tmp16435 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16437 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16437 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16438 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5syntax_6 := __e.Get(1)
_ = Parseshen_4_5syntax_6
tmp16450 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5syntax_6)


if True == tmp16450 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16439 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp16447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp16447 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16440 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sc_6)


tmp16441 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5syntax_6)


tmp16442 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5syntax_6)


tmp16443 := Call(__e, PrimFunc(symshen_4autocomplete), tmp16442)


tmp16444 := PrimCons(tmp16443, Nil)

tmp16445 := PrimCons(tmp16441, tmp16444)

__e.TailApply(PrimFunc(symshen_4comb), tmp16440, tmp16445)
return


}


}, 1)

tmp16448 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5syntax_6)


__e.TailApply(tmp16439, tmp16448)
return


}


}, 1)

tmp16451 := Call(__e, PrimFunc(symshen_4_5syntax_6), V1156)


tmp16452 := Call(__e, tmp16438, tmp16451)


__e.TailApply(tmp16435, tmp16452)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16455 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5syntax_6 := __e.Get(1)
_ = Parseshen_4_5syntax_6
tmp16470 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5syntax_6)


if True == tmp16470 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16456 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5semantics_6 := __e.Get(1)
_ = Parseshen_4_5semantics_6
tmp16467 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5semantics_6)


if True == tmp16467 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16457 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp16464 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp16464 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16458 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sc_6)


tmp16459 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5syntax_6)


tmp16460 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5semantics_6)


tmp16461 := PrimCons(tmp16460, Nil)

tmp16462 := PrimCons(tmp16459, tmp16461)

__e.TailApply(PrimFunc(symshen_4comb), tmp16458, tmp16462)
return


}


}, 1)

tmp16465 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5semantics_6)


__e.TailApply(tmp16457, tmp16465)
return


}


}, 1)

tmp16468 := Call(__e, PrimFunc(symshen_4_5semantics_6), Parseshen_4_5syntax_6)


__e.TailApply(tmp16456, tmp16468)
return


}


}, 1)

tmp16471 := Call(__e, PrimFunc(symshen_4_5syntax_6), V1156)


tmp16472 := Call(__e, tmp16455, tmp16471)


__e.TailApply(tmp16434, tmp16472)
return


}, 1)

tmp16473 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp16433)


_ = tmp16473

tmp16474 := MakeNative(func(__e *ControlFlow) {
V1157 := __e.Get(1)
_ = V1157
tmp16503 := PrimIsPair(V1157)

var ifres16495 Obj

if True == tmp16503 {
tmp16501 := PrimTail(V1157)

tmp16502 := PrimEqual(Nil, tmp16501)

var ifres16497 Obj

if True == tmp16502 {
tmp16499 := PrimHead(V1157)

tmp16500 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16499)


var ifres16498 Obj

if True == tmp16500 {
ifres16498 = True


} else {
ifres16498 = False


}

ifres16497 = ifres16498


} else {
ifres16497 = False


}

var ifres16496 Obj

if True == ifres16497 {
ifres16496 = True


} else {
ifres16496 = False


}

ifres16495 = ifres16496


} else {
ifres16495 = False


}

if True == ifres16495 {
__e.Return(PrimHead(V1157))
return
} else {
tmp16493 := PrimIsPair(V1157)

var ifres16489 Obj

if True == tmp16493 {
tmp16491 := PrimHead(V1157)

tmp16492 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16491)


var ifres16490 Obj

if True == tmp16492 {
ifres16490 = True


} else {
ifres16490 = False


}

ifres16489 = ifres16490


} else {
ifres16489 = False


}

if True == ifres16489 {
tmp16475 := PrimHead(V1157)

tmp16476 := PrimTail(V1157)

tmp16477 := Call(__e, PrimFunc(symshen_4autocomplete), tmp16476)


tmp16478 := PrimCons(tmp16477, Nil)

tmp16479 := PrimCons(tmp16475, tmp16478)

__e.Return(PrimCons(symappend, tmp16479))
return


} else {
tmp16487 := PrimIsPair(V1157)

if True == tmp16487 {
tmp16480 := PrimHead(V1157)

tmp16481 := Call(__e, PrimFunc(symshen_4autocomplete), tmp16480)


tmp16482 := PrimTail(V1157)

tmp16483 := Call(__e, PrimFunc(symshen_4autocomplete), tmp16482)


tmp16484 := PrimCons(tmp16483, Nil)

tmp16485 := PrimCons(tmp16481, tmp16484)

__e.Return(PrimCons(symcons, tmp16485))
return


} else {
__e.Return(V1157)
return
}


}


}


}, 1)

tmp16504 := Call(__e, ns2_1set, symshen_4autocomplete, tmp16474)


_ = tmp16504

tmp16505 := MakeNative(func(__e *ControlFlow) {
V1158 := __e.Get(1)
_ = V1158
tmp16512 := PrimIsSymbol(V1158)

if True == tmp16512 {
tmp16507 := MakeNative(func(__e *ControlFlow) {
Explode := __e.Get(1)
_ = Explode
tmp16508 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), X)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16508, Explode)
return


}, 1)

tmp16509 := Call(__e, PrimFunc(symexplode), V1158)


tmp16510 := Call(__e, tmp16507, tmp16509)


if True == tmp16510 {
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

tmp16513 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp16505)


_ = tmp16513

tmp16514 := MakeNative(func(__e *ControlFlow) {
V1159 := __e.Get(1)
_ = V1159
tmp16515 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16535 {
tmp16516 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16527 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16527 {
tmp16517 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16519 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16519 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16520 := MakeNative(func(__e *ControlFlow) {
Parse_5_b_6 := __e.Get(1)
_ = Parse_5_b_6
tmp16523 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5_b_6)


if True == tmp16523 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16521 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5_b_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16521, False)
return


}


}, 1)

tmp16524 := Call(__e, PrimFunc(sym_5_b_6), V1159)


tmp16525 := Call(__e, tmp16520, tmp16524)


__e.TailApply(tmp16517, tmp16525)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16528 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5non_1terminal_1name_6 := __e.Get(1)
_ = Parseshen_4_5non_1terminal_1name_6
tmp16531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5non_1terminal_1name_6)


if True == tmp16531 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16529 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5non_1terminal_1name_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16529, True)
return


}


}, 1)

tmp16532 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V1159)


tmp16533 := Call(__e, tmp16528, tmp16532)


__e.TailApply(tmp16516, tmp16533)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16536 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5packagenames_6 := __e.Get(1)
_ = Parseshen_4_5packagenames_6
tmp16543 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5packagenames_6)


if True == tmp16543 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16537 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5non_1terminal_1name_6 := __e.Get(1)
_ = Parseshen_4_5non_1terminal_1name_6
tmp16540 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5non_1terminal_1name_6)


if True == tmp16540 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16538 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5non_1terminal_1name_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16538, True)
return


}


}, 1)

tmp16541 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), Parseshen_4_5packagenames_6)


__e.TailApply(tmp16537, tmp16541)
return


}


}, 1)

tmp16544 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V1159)


tmp16545 := Call(__e, tmp16536, tmp16544)


__e.TailApply(tmp16515, tmp16545)
return


}, 1)

tmp16546 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp16514)


_ = tmp16546

tmp16547 := MakeNative(func(__e *ControlFlow) {
V1160 := __e.Get(1)
_ = V1160
tmp16548 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16563 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16563 {
tmp16549 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16551 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16551 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16552 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5packagename_6 := __e.Get(1)
_ = Parseshen_4_5packagename_6
tmp16559 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5packagename_6)


if True == tmp16559 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16557 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5packagename_6, MakeString("."))


if True == tmp16557 {
tmp16553 := MakeNative(func(__e *ControlFlow) {
News1047 := __e.Get(1)
_ = News1047
tmp16554 := Call(__e, PrimFunc(symshen_4in_1_6), News1047)


__e.TailApply(PrimFunc(symshen_4comb), tmp16554, symshen_4skip)
return


}, 1)

tmp16555 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5packagename_6)


__e.TailApply(tmp16553, tmp16555)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp16560 := Call(__e, PrimFunc(symshen_4_5packagename_6), V1160)


tmp16561 := Call(__e, tmp16552, tmp16560)


__e.TailApply(tmp16549, tmp16561)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16564 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5packagename_6 := __e.Get(1)
_ = Parseshen_4_5packagename_6
tmp16575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5packagename_6)


if True == tmp16575 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16573 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5packagename_6, MakeString("."))


if True == tmp16573 {
tmp16565 := MakeNative(func(__e *ControlFlow) {
News1046 := __e.Get(1)
_ = News1046
tmp16566 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5packagenames_6 := __e.Get(1)
_ = Parseshen_4_5packagenames_6
tmp16569 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5packagenames_6)


if True == tmp16569 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16567 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5packagenames_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16567, symshen_4skip)
return


}


}, 1)

tmp16570 := Call(__e, PrimFunc(symshen_4_5packagenames_6), News1046)


__e.TailApply(tmp16566, tmp16570)
return


}, 1)

tmp16571 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5packagename_6)


__e.TailApply(tmp16565, tmp16571)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp16576 := Call(__e, PrimFunc(symshen_4_5packagename_6), V1160)


tmp16577 := Call(__e, tmp16564, tmp16576)


__e.TailApply(tmp16548, tmp16577)
return


}, 1)

tmp16578 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp16547)


_ = tmp16578

tmp16579 := MakeNative(func(__e *ControlFlow) {
V1161 := __e.Get(1)
_ = V1161
tmp16580 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16591 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16591 {
tmp16581 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16584 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp16587 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp16587 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16585 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16585, symshen_4skip)
return


}


}, 1)

tmp16588 := Call(__e, PrimFunc(sym_5e_6), V1161)


tmp16589 := Call(__e, tmp16584, tmp16588)


__e.TailApply(tmp16581, tmp16589)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16592 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5packagechar_6 := __e.Get(1)
_ = Parseshen_4_5packagechar_6
tmp16599 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5packagechar_6)


if True == tmp16599 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16593 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5packagename_6 := __e.Get(1)
_ = Parseshen_4_5packagename_6
tmp16596 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5packagename_6)


if True == tmp16596 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16594 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5packagename_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16594, symshen_4skip)
return


}


}, 1)

tmp16597 := Call(__e, PrimFunc(symshen_4_5packagename_6), Parseshen_4_5packagechar_6)


__e.TailApply(tmp16593, tmp16597)
return


}


}, 1)

tmp16600 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V1161)


tmp16601 := Call(__e, tmp16592, tmp16600)


__e.TailApply(tmp16580, tmp16601)
return


}, 1)

tmp16602 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp16579)


_ = tmp16602

tmp16603 := MakeNative(func(__e *ControlFlow) {
V1162 := __e.Get(1)
_ = V1162
tmp16604 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16606 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16606 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16618 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1162)


var ifres16607 Obj

if True == tmp16618 {
tmp16608 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp16609 := MakeNative(func(__e *ControlFlow) {
News1050 := __e.Get(1)
_ = News1050
tmp16612 := PrimEqual(X, MakeString("."))

tmp16613 := PrimNot(tmp16612)

if True == tmp16613 {
tmp16610 := Call(__e, PrimFunc(symshen_4in_1_6), News1050)


__e.TailApply(PrimFunc(symshen_4comb), tmp16610, symshen_4skip)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16614 := Call(__e, PrimFunc(symshen_4tls), V1162)


__e.TailApply(tmp16609, tmp16614)
return


}, 1)

tmp16615 := Call(__e, PrimFunc(symshen_4hds), V1162)


tmp16616 := Call(__e, tmp16608, tmp16615)


ifres16607 = tmp16616


} else {
tmp16617 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16607 = tmp16617


}

__e.TailApply(tmp16604, ifres16607)
return


}, 1)

tmp16619 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp16603)


_ = tmp16619

tmp16620 := MakeNative(func(__e *ControlFlow) {
V1163 := __e.Get(1)
_ = V1163
tmp16621 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16623 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16623 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16644 := Call(__e, PrimFunc(symshen_4_ahd_2), V1163, MakeString("<"))


var ifres16624 Obj

if True == tmp16644 {
tmp16625 := MakeNative(func(__e *ControlFlow) {
News1052 := __e.Get(1)
_ = News1052
tmp16626 := MakeNative(func(__e *ControlFlow) {
Parse_5_b_6 := __e.Get(1)
_ = Parse_5_b_6
tmp16639 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5_b_6)


if True == tmp16639 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16629 := MakeNative(func(__e *ControlFlow) {
Reverse := __e.Get(1)
_ = Reverse
tmp16634 := PrimIsPair(Reverse)

if True == tmp16634 {
tmp16631 := PrimHead(Reverse)

tmp16632 := PrimEqual(tmp16631, MakeString(">"))

if True == tmp16632 {
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

tmp16635 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp16636 := Call(__e, PrimFunc(symreverse), tmp16635)


tmp16637 := Call(__e, tmp16629, tmp16636)


if True == tmp16637 {
tmp16627 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5_b_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp16627, symshen_4skip)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp16640 := Call(__e, PrimFunc(sym_5_b_6), News1052)


__e.TailApply(tmp16626, tmp16640)
return


}, 1)

tmp16641 := Call(__e, PrimFunc(symshen_4tls), V1163)


tmp16642 := Call(__e, tmp16625, tmp16641)


ifres16624 = tmp16642


} else {
tmp16643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16624 = tmp16643


}

__e.TailApply(tmp16621, ifres16624)
return


}, 1)

tmp16645 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp16620)


_ = tmp16645

tmp16646 := MakeNative(func(__e *ControlFlow) {
V1164 := __e.Get(1)
_ = V1164
tmp16647 := PrimIntern(MakeString(";"))

__e.Return(PrimEqual(V1164, tmp16647))
return


}, 1)

tmp16648 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp16646)


_ = tmp16648

tmp16649 := MakeNative(func(__e *ControlFlow) {
V1165 := __e.Get(1)
_ = V1165
tmp16650 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16652 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16663 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1165)


var ifres16653 Obj

if True == tmp16663 {
tmp16654 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp16655 := MakeNative(func(__e *ControlFlow) {
News1054 := __e.Get(1)
_ = News1054
tmp16658 := Call(__e, PrimFunc(symshen_4colon_1equal_2), X)


if True == tmp16658 {
tmp16656 := Call(__e, PrimFunc(symshen_4in_1_6), News1054)


__e.TailApply(PrimFunc(symshen_4comb), tmp16656, symshen_4skip)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16659 := Call(__e, PrimFunc(symshen_4tls), V1165)


__e.TailApply(tmp16655, tmp16659)
return


}, 1)

tmp16660 := Call(__e, PrimFunc(symshen_4hds), V1165)


tmp16661 := Call(__e, tmp16654, tmp16660)


ifres16653 = tmp16661


} else {
tmp16662 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16653 = tmp16662


}

__e.TailApply(tmp16650, ifres16653)
return


}, 1)

tmp16664 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp16649)


_ = tmp16664

tmp16665 := MakeNative(func(__e *ControlFlow) {
V1166 := __e.Get(1)
_ = V1166
tmp16666 := PrimIntern(MakeString(":="))

__e.Return(PrimEqual(tmp16666, V1166))
return


}, 1)

tmp16667 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp16665)


_ = tmp16667

tmp16668 := MakeNative(func(__e *ControlFlow) {
V1167 := __e.Get(1)
_ = V1167
tmp16669 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16682 {
tmp16670 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16672 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16672 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16673 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5syntax_1item_6 := __e.Get(1)
_ = Parseshen_4_5syntax_1item_6
tmp16678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5syntax_1item_6)


if True == tmp16678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16674 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5syntax_1item_6)


tmp16675 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5syntax_1item_6)


tmp16676 := PrimCons(tmp16675, Nil)

__e.TailApply(PrimFunc(symshen_4comb), tmp16674, tmp16676)
return


}


}, 1)

tmp16679 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V1167)


tmp16680 := Call(__e, tmp16673, tmp16679)


__e.TailApply(tmp16670, tmp16680)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16683 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5syntax_1item_6 := __e.Get(1)
_ = Parseshen_4_5syntax_1item_6
tmp16693 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5syntax_1item_6)


if True == tmp16693 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16684 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5syntax_6 := __e.Get(1)
_ = Parseshen_4_5syntax_6
tmp16690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5syntax_6)


if True == tmp16690 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16685 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5syntax_6)


tmp16686 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5syntax_1item_6)


tmp16687 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5syntax_6)


tmp16688 := PrimCons(tmp16686, tmp16687)

__e.TailApply(PrimFunc(symshen_4comb), tmp16685, tmp16688)
return


}


}, 1)

tmp16691 := Call(__e, PrimFunc(symshen_4_5syntax_6), Parseshen_4_5syntax_1item_6)


__e.TailApply(tmp16684, tmp16691)
return


}


}, 1)

tmp16694 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V1167)


tmp16695 := Call(__e, tmp16683, tmp16694)


__e.TailApply(tmp16669, tmp16695)
return


}, 1)

tmp16696 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp16668)


_ = tmp16696

tmp16697 := MakeNative(func(__e *ControlFlow) {
V1168 := __e.Get(1)
_ = V1168
tmp16698 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16700 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16700 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16711 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1168)


var ifres16701 Obj

if True == tmp16711 {
tmp16702 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp16703 := MakeNative(func(__e *ControlFlow) {
News1057 := __e.Get(1)
_ = News1057
tmp16706 := Call(__e, PrimFunc(symshen_4syntax_1item_2), X)


if True == tmp16706 {
tmp16704 := Call(__e, PrimFunc(symshen_4in_1_6), News1057)


__e.TailApply(PrimFunc(symshen_4comb), tmp16704, X)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16707 := Call(__e, PrimFunc(symshen_4tls), V1168)


__e.TailApply(tmp16703, tmp16707)
return


}, 1)

tmp16708 := Call(__e, PrimFunc(symshen_4hds), V1168)


tmp16709 := Call(__e, tmp16702, tmp16708)


ifres16701 = tmp16709


} else {
tmp16710 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16701 = tmp16710


}

__e.TailApply(tmp16698, ifres16701)
return


}, 1)

tmp16712 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp16697)


_ = tmp16712

tmp16713 := MakeNative(func(__e *ControlFlow) {
V1171 := __e.Get(1)
_ = V1171
tmp16749 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V1171)


if True == tmp16749 {
__e.Return(False)
return
} else {
tmp16747 := Call(__e, PrimFunc(symshen_4semicolon_2), V1171)


if True == tmp16747 {
__e.Return(False)
return
} else {
tmp16745 := Call(__e, PrimFunc(symatom_2), V1171)


if True == tmp16745 {
__e.Return(True)
return
} else {
tmp16743 := PrimIsPair(V1171)

var ifres16724 Obj

if True == tmp16743 {
tmp16741 := PrimHead(V1171)

tmp16742 := PrimEqual(symcons, tmp16741)

var ifres16726 Obj

if True == tmp16742 {
tmp16739 := PrimTail(V1171)

tmp16740 := PrimIsPair(tmp16739)

var ifres16728 Obj

if True == tmp16740 {
tmp16736 := PrimTail(V1171)

tmp16737 := PrimTail(tmp16736)

tmp16738 := PrimIsPair(tmp16737)

var ifres16730 Obj

if True == tmp16738 {
tmp16732 := PrimTail(V1171)

tmp16733 := PrimTail(tmp16732)

tmp16734 := PrimTail(tmp16733)

tmp16735 := PrimEqual(Nil, tmp16734)

var ifres16731 Obj

if True == tmp16735 {
ifres16731 = True


} else {
ifres16731 = False


}

ifres16730 = ifres16731


} else {
ifres16730 = False


}

var ifres16729 Obj

if True == ifres16730 {
ifres16729 = True


} else {
ifres16729 = False


}

ifres16728 = ifres16729


} else {
ifres16728 = False


}

var ifres16727 Obj

if True == ifres16728 {
ifres16727 = True


} else {
ifres16727 = False


}

ifres16726 = ifres16727


} else {
ifres16726 = False


}

var ifres16725 Obj

if True == ifres16726 {
ifres16725 = True


} else {
ifres16725 = False


}

ifres16724 = ifres16725


} else {
ifres16724 = False


}

if True == ifres16724 {
tmp16720 := PrimTail(V1171)

tmp16721 := PrimHead(tmp16720)

tmp16722 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp16721)


if True == tmp16722 {
tmp16715 := PrimTail(V1171)

tmp16716 := PrimTail(tmp16715)

tmp16717 := PrimHead(tmp16716)

tmp16718 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp16717)


if True == tmp16718 {
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
__e.Return(False)
return
}


}


}


}


}, 1)

tmp16750 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp16713)


_ = tmp16750

tmp16751 := MakeNative(func(__e *ControlFlow) {
V1172 := __e.Get(1)
_ = V1172
tmp16752 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16772 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16772 {
tmp16753 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp16755 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp16755 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp16756 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5colon_1equal_6 := __e.Get(1)
_ = Parseshen_4_5colon_1equal_6
tmp16768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5colon_1equal_6)


if True == tmp16768 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16766 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), Parseshen_4_5colon_1equal_6)


if True == tmp16766 {
tmp16757 := MakeNative(func(__e *ControlFlow) {
Semantics := __e.Get(1)
_ = Semantics
tmp16758 := MakeNative(func(__e *ControlFlow) {
News1062 := __e.Get(1)
_ = News1062
tmp16761 := Call(__e, PrimFunc(symshen_4semicolon_2), Semantics)


tmp16762 := PrimNot(tmp16761)

if True == tmp16762 {
tmp16759 := Call(__e, PrimFunc(symshen_4in_1_6), News1062)


__e.TailApply(PrimFunc(symshen_4comb), tmp16759, Semantics)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16763 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5colon_1equal_6)


__e.TailApply(tmp16758, tmp16763)
return


}, 1)

tmp16764 := Call(__e, PrimFunc(symshen_4hds), Parseshen_4_5colon_1equal_6)


__e.TailApply(tmp16757, tmp16764)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp16769 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V1172)


tmp16770 := Call(__e, tmp16756, tmp16769)


__e.TailApply(tmp16753, tmp16770)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp16773 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5colon_1equal_6 := __e.Get(1)
_ = Parseshen_4_5colon_1equal_6
tmp16798 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5colon_1equal_6)


if True == tmp16798 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16796 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), Parseshen_4_5colon_1equal_6)


if True == tmp16796 {
tmp16774 := MakeNative(func(__e *ControlFlow) {
Semantics := __e.Get(1)
_ = Semantics
tmp16775 := MakeNative(func(__e *ControlFlow) {
News1059 := __e.Get(1)
_ = News1059
tmp16792 := Call(__e, PrimFunc(symshen_4_ahd_2), News1059, symwhere)


if True == tmp16792 {
tmp16776 := MakeNative(func(__e *ControlFlow) {
News1060 := __e.Get(1)
_ = News1060
tmp16789 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1060)


if True == tmp16789 {
tmp16777 := MakeNative(func(__e *ControlFlow) {
Guard := __e.Get(1)
_ = Guard
tmp16778 := MakeNative(func(__e *ControlFlow) {
News1061 := __e.Get(1)
_ = News1061
tmp16784 := Call(__e, PrimFunc(symshen_4semicolon_2), Semantics)


tmp16785 := PrimNot(tmp16784)

if True == tmp16785 {
tmp16779 := Call(__e, PrimFunc(symshen_4in_1_6), News1061)


tmp16780 := PrimCons(Semantics, Nil)

tmp16781 := PrimCons(Guard, tmp16780)

tmp16782 := PrimCons(symwhere, tmp16781)

__e.TailApply(PrimFunc(symshen_4comb), tmp16779, tmp16782)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16786 := Call(__e, PrimFunc(symshen_4tls), News1060)


__e.TailApply(tmp16778, tmp16786)
return


}, 1)

tmp16787 := Call(__e, PrimFunc(symshen_4hds), News1060)


__e.TailApply(tmp16777, tmp16787)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16790 := Call(__e, PrimFunc(symshen_4tls), News1059)


__e.TailApply(tmp16776, tmp16790)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16793 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5colon_1equal_6)


__e.TailApply(tmp16775, tmp16793)
return


}, 1)

tmp16794 := Call(__e, PrimFunc(symshen_4hds), Parseshen_4_5colon_1equal_6)


__e.TailApply(tmp16774, tmp16794)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp16799 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V1172)


tmp16800 := Call(__e, tmp16773, tmp16799)


__e.TailApply(tmp16752, tmp16800)
return


}, 1)

tmp16801 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp16751)


_ = tmp16801

tmp16802 := MakeNative(func(__e *ControlFlow) {
V1181 := __e.Get(1)
_ = V1181
V1182 := __e.Get(2)
_ = V1182
V1183 := __e.Get(3)
_ = V1183
tmp16810 := PrimEqual(Nil, V1183)

if True == tmp16810 {
__e.Return(PrimCons(symshen_4parse_1failure, Nil))
return
} else {
tmp16808 := PrimIsPair(V1183)

if True == tmp16808 {
tmp16803 := PrimHead(V1183)

tmp16804 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V1181, tmp16803, V1182)


tmp16805 := PrimTail(V1183)

tmp16806 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V1181, V1182, tmp16805)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp16804, tmp16806)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
return
}


}


}, 3)

tmp16811 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp16802)


_ = tmp16811

tmp16812 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp16813 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp16812)


_ = tmp16813

tmp16814 := MakeNative(func(__e *ControlFlow) {
V1184 := __e.Get(1)
_ = V1184
V1185 := __e.Get(2)
_ = V1185
tmp16815 := Call(__e, PrimFunc(symprotect), symResult)


tmp16816 := Call(__e, PrimFunc(symprotect), symResult)


tmp16817 := PrimCons(tmp16816, Nil)

tmp16818 := PrimCons(symshen_4parse_1failure_2, tmp16817)

tmp16819 := Call(__e, PrimFunc(symprotect), symResult)


tmp16820 := PrimCons(tmp16819, Nil)

tmp16821 := PrimCons(V1185, tmp16820)

tmp16822 := PrimCons(tmp16818, tmp16821)

tmp16823 := PrimCons(symif, tmp16822)

tmp16824 := PrimCons(tmp16823, Nil)

tmp16825 := PrimCons(V1184, tmp16824)

tmp16826 := PrimCons(tmp16815, tmp16825)

__e.Return(PrimCons(symlet, tmp16826))
return


}, 2)

tmp16827 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp16814)


_ = tmp16827

tmp16828 := MakeNative(func(__e *ControlFlow) {
V1192 := __e.Get(1)
_ = V1192
V1193 := __e.Get(2)
_ = V1193
V1194 := __e.Get(3)
_ = V1194
tmp16842 := PrimIsPair(V1193)

var ifres16833 Obj

if True == tmp16842 {
tmp16840 := PrimTail(V1193)

tmp16841 := PrimIsPair(tmp16840)

var ifres16835 Obj

if True == tmp16841 {
tmp16837 := PrimTail(V1193)

tmp16838 := PrimTail(tmp16837)

tmp16839 := PrimEqual(Nil, tmp16838)

var ifres16836 Obj

if True == tmp16839 {
ifres16836 = True


} else {
ifres16836 = False


}

ifres16835 = ifres16836


} else {
ifres16835 = False


}

var ifres16834 Obj

if True == ifres16835 {
ifres16834 = True


} else {
ifres16834 = False


}

ifres16833 = ifres16834


} else {
ifres16833 = False


}

if True == ifres16833 {
tmp16829 := PrimHead(V1193)

tmp16830 := PrimTail(V1193)

tmp16831 := PrimHead(tmp16830)

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V1192, V1194, tmp16829, tmp16831)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
return
}


}, 3)

tmp16843 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp16828)


_ = tmp16843

tmp16844 := MakeNative(func(__e *ControlFlow) {
V1203 := __e.Get(1)
_ = V1203
V1204 := __e.Get(2)
_ = V1204
V1205 := __e.Get(3)
_ = V1205
V1206 := __e.Get(4)
_ = V1206
tmp16908 := PrimEqual(Nil, V1205)

var ifres16886 Obj

if True == tmp16908 {
tmp16907 := PrimIsPair(V1206)

var ifres16888 Obj

if True == tmp16907 {
tmp16905 := PrimHead(V1206)

tmp16906 := PrimEqual(symwhere, tmp16905)

var ifres16890 Obj

if True == tmp16906 {
tmp16903 := PrimTail(V1206)

tmp16904 := PrimIsPair(tmp16903)

var ifres16892 Obj

if True == tmp16904 {
tmp16900 := PrimTail(V1206)

tmp16901 := PrimTail(tmp16900)

tmp16902 := PrimIsPair(tmp16901)

var ifres16894 Obj

if True == tmp16902 {
tmp16896 := PrimTail(V1206)

tmp16897 := PrimTail(tmp16896)

tmp16898 := PrimTail(tmp16897)

tmp16899 := PrimEqual(Nil, tmp16898)

var ifres16895 Obj

if True == tmp16899 {
ifres16895 = True


} else {
ifres16895 = False


}

ifres16894 = ifres16895


} else {
ifres16894 = False


}

var ifres16893 Obj

if True == ifres16894 {
ifres16893 = True


} else {
ifres16893 = False


}

ifres16892 = ifres16893


} else {
ifres16892 = False


}

var ifres16891 Obj

if True == ifres16892 {
ifres16891 = True


} else {
ifres16891 = False


}

ifres16890 = ifres16891


} else {
ifres16890 = False


}

var ifres16889 Obj

if True == ifres16890 {
ifres16889 = True


} else {
ifres16889 = False


}

ifres16888 = ifres16889


} else {
ifres16888 = False


}

var ifres16887 Obj

if True == ifres16888 {
ifres16887 = True


} else {
ifres16887 = False


}

ifres16886 = ifres16887


} else {
ifres16886 = False


}

if True == ifres16886 {
tmp16845 := PrimTail(V1206)

tmp16846 := PrimHead(tmp16845)

tmp16847 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp16846)


tmp16848 := PrimTail(V1206)

tmp16849 := PrimTail(tmp16848)

tmp16850 := PrimHead(tmp16849)

tmp16851 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V1203, V1204, Nil, tmp16850)


tmp16852 := PrimCons(symshen_4parse_1failure, Nil)

tmp16853 := PrimCons(tmp16852, Nil)

tmp16854 := PrimCons(tmp16851, tmp16853)

tmp16855 := PrimCons(tmp16847, tmp16854)

__e.Return(PrimCons(symif, tmp16855))
return


} else {
tmp16884 := PrimEqual(Nil, V1205)

if True == tmp16884 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V1203, V1204, V1206)
return
} else {
tmp16882 := PrimIsPair(V1205)

if True == tmp16882 {
tmp16879 := PrimHead(V1205)

tmp16880 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16879)


if True == tmp16880 {
tmp16856 := PrimHead(V1205)

tmp16857 := PrimTail(V1205)

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V1203, V1204, tmp16856, tmp16857, V1206)
return


} else {
tmp16876 := PrimHead(V1205)

tmp16877 := PrimIsVariable(tmp16876)

if True == tmp16877 {
tmp16858 := PrimHead(V1205)

tmp16859 := PrimTail(V1205)

__e.TailApply(PrimFunc(symshen_4variablecode), V1203, V1204, tmp16858, tmp16859, V1206)
return


} else {
tmp16873 := PrimHead(V1205)

tmp16874 := PrimEqual(sym__, tmp16873)

if True == tmp16874 {
tmp16860 := PrimHead(V1205)

tmp16861 := PrimTail(V1205)

__e.TailApply(PrimFunc(symshen_4wildcardcode), V1203, V1204, tmp16860, tmp16861, V1206)
return


} else {
tmp16870 := PrimHead(V1205)

tmp16871 := Call(__e, PrimFunc(symatom_2), tmp16870)


if True == tmp16871 {
tmp16862 := PrimHead(V1205)

tmp16863 := PrimTail(V1205)

__e.TailApply(PrimFunc(symshen_4terminalcode), V1203, V1204, tmp16862, tmp16863, V1206)
return


} else {
tmp16867 := PrimHead(V1205)

tmp16868 := PrimIsPair(tmp16867)

if True == tmp16868 {
tmp16864 := PrimHead(V1205)

tmp16865 := PrimTail(V1205)

__e.TailApply(PrimFunc(symshen_4conscode), V1203, V1204, tmp16864, tmp16865, V1206)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}, 4)

tmp16909 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp16844)


_ = tmp16909

tmp16910 := MakeNative(func(__e *ControlFlow) {
V1207 := __e.Get(1)
_ = V1207
V1208 := __e.Get(2)
_ = V1208
V1209 := __e.Get(3)
_ = V1209
V1210 := __e.Get(4)
_ = V1210
V1211 := __e.Get(5)
_ = V1211
tmp16911 := MakeNative(func(__e *ControlFlow) {
ApplyNonTerminal := __e.Get(1)
_ = ApplyNonTerminal
tmp16912 := PrimCons(V1208, Nil)

tmp16913 := PrimCons(V1209, tmp16912)

tmp16914 := PrimCons(ApplyNonTerminal, Nil)

tmp16915 := PrimCons(symshen_4parse_1failure_2, tmp16914)

tmp16916 := PrimCons(symshen_4parse_1failure, Nil)

tmp16917 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V1207, ApplyNonTerminal, V1210, V1211)


tmp16918 := PrimCons(tmp16917, Nil)

tmp16919 := PrimCons(tmp16916, tmp16918)

tmp16920 := PrimCons(tmp16915, tmp16919)

tmp16921 := PrimCons(symif, tmp16920)

tmp16922 := PrimCons(tmp16921, Nil)

tmp16923 := PrimCons(tmp16913, tmp16922)

tmp16924 := PrimCons(ApplyNonTerminal, tmp16923)

__e.Return(PrimCons(symlet, tmp16924))
return


}, 1)

tmp16925 := Call(__e, PrimFunc(symprotect), symParse)


tmp16926 := Call(__e, PrimFunc(symconcat), tmp16925, V1209)


__e.TailApply(tmp16911, tmp16926)
return


}, 5)

tmp16927 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp16910)


_ = tmp16927

tmp16928 := MakeNative(func(__e *ControlFlow) {
V1212 := __e.Get(1)
_ = V1212
V1213 := __e.Get(2)
_ = V1213
V1214 := __e.Get(3)
_ = V1214
V1215 := __e.Get(4)
_ = V1215
V1216 := __e.Get(5)
_ = V1216
tmp16929 := MakeNative(func(__e *ControlFlow) {
NewStream := __e.Get(1)
_ = NewStream
tmp16930 := PrimCons(V1213, Nil)

tmp16931 := PrimCons(symshen_4non_1empty_1stream_2, tmp16930)

tmp16932 := PrimCons(V1213, Nil)

tmp16933 := PrimCons(symshen_4hds, tmp16932)

tmp16934 := PrimCons(V1213, Nil)

tmp16935 := PrimCons(symshen_4tls, tmp16934)

tmp16936 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V1212, NewStream, V1215, V1216)


tmp16937 := PrimCons(tmp16936, Nil)

tmp16938 := PrimCons(tmp16935, tmp16937)

tmp16939 := PrimCons(NewStream, tmp16938)

tmp16940 := PrimCons(tmp16933, tmp16939)

tmp16941 := PrimCons(V1214, tmp16940)

tmp16942 := PrimCons(symlet, tmp16941)

tmp16943 := PrimCons(symshen_4parse_1failure, Nil)

tmp16944 := PrimCons(tmp16943, Nil)

tmp16945 := PrimCons(tmp16942, tmp16944)

tmp16946 := PrimCons(tmp16931, tmp16945)

__e.Return(PrimCons(symif, tmp16946))
return


}, 1)

tmp16947 := Call(__e, PrimFunc(symprotect), symNews)


tmp16948 := Call(__e, PrimFunc(symgensym), tmp16947)


__e.TailApply(tmp16929, tmp16948)
return


}, 5)

tmp16949 := Call(__e, ns2_1set, symshen_4variablecode, tmp16928)


_ = tmp16949

tmp16950 := MakeNative(func(__e *ControlFlow) {
V1217 := __e.Get(1)
_ = V1217
V1218 := __e.Get(2)
_ = V1218
V1219 := __e.Get(3)
_ = V1219
V1220 := __e.Get(4)
_ = V1220
V1221 := __e.Get(5)
_ = V1221
tmp16951 := MakeNative(func(__e *ControlFlow) {
NewStream := __e.Get(1)
_ = NewStream
tmp16952 := PrimCons(V1218, Nil)

tmp16953 := PrimCons(symshen_4non_1empty_1stream_2, tmp16952)

tmp16954 := PrimCons(V1218, Nil)

tmp16955 := PrimCons(symshen_4tls, tmp16954)

tmp16956 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V1217, NewStream, V1220, V1221)


tmp16957 := PrimCons(tmp16956, Nil)

tmp16958 := PrimCons(tmp16955, tmp16957)

tmp16959 := PrimCons(NewStream, tmp16958)

tmp16960 := PrimCons(symlet, tmp16959)

tmp16961 := PrimCons(symshen_4parse_1failure, Nil)

tmp16962 := PrimCons(tmp16961, Nil)

tmp16963 := PrimCons(tmp16960, tmp16962)

tmp16964 := PrimCons(tmp16953, tmp16963)

__e.Return(PrimCons(symif, tmp16964))
return


}, 1)

tmp16965 := Call(__e, PrimFunc(symprotect), symNews)


tmp16966 := Call(__e, PrimFunc(symgensym), tmp16965)


__e.TailApply(tmp16951, tmp16966)
return


}, 5)

tmp16967 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp16950)


_ = tmp16967

tmp16968 := MakeNative(func(__e *ControlFlow) {
V1222 := __e.Get(1)
_ = V1222
V1223 := __e.Get(2)
_ = V1223
V1224 := __e.Get(3)
_ = V1224
V1225 := __e.Get(4)
_ = V1225
V1226 := __e.Get(5)
_ = V1226
tmp16969 := MakeNative(func(__e *ControlFlow) {
NewStream := __e.Get(1)
_ = NewStream
tmp16970 := PrimCons(V1224, Nil)

tmp16971 := PrimCons(V1223, tmp16970)

tmp16972 := PrimCons(symshen_4_ahd_2, tmp16971)

tmp16973 := PrimCons(V1223, Nil)

tmp16974 := PrimCons(symshen_4tls, tmp16973)

tmp16975 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V1222, NewStream, V1225, V1226)


tmp16976 := PrimCons(tmp16975, Nil)

tmp16977 := PrimCons(tmp16974, tmp16976)

tmp16978 := PrimCons(NewStream, tmp16977)

tmp16979 := PrimCons(symlet, tmp16978)

tmp16980 := PrimCons(symshen_4parse_1failure, Nil)

tmp16981 := PrimCons(tmp16980, Nil)

tmp16982 := PrimCons(tmp16979, tmp16981)

tmp16983 := PrimCons(tmp16972, tmp16982)

__e.Return(PrimCons(symif, tmp16983))
return


}, 1)

tmp16984 := Call(__e, PrimFunc(symprotect), symNews)


tmp16985 := Call(__e, PrimFunc(symgensym), tmp16984)


__e.TailApply(tmp16969, tmp16985)
return


}, 5)

tmp16986 := Call(__e, ns2_1set, symshen_4terminalcode, tmp16968)


_ = tmp16986

tmp16987 := MakeNative(func(__e *ControlFlow) {
V1227 := __e.Get(1)
_ = V1227
V1228 := __e.Get(2)
_ = V1228
V1229 := __e.Get(3)
_ = V1229
V1230 := __e.Get(4)
_ = V1230
V1231 := __e.Get(5)
_ = V1231
tmp16988 := PrimCons(V1228, Nil)

tmp16989 := PrimCons(symshen_4ccons_2, tmp16988)

tmp16990 := Call(__e, PrimFunc(symprotect), symSynCons)


tmp16991 := PrimCons(V1228, Nil)

tmp16992 := PrimCons(symshen_4hds, tmp16991)

tmp16993 := PrimCons(V1228, Nil)

tmp16994 := PrimCons(symshen_4_5_1out, tmp16993)

tmp16995 := PrimCons(tmp16994, Nil)

tmp16996 := PrimCons(tmp16992, tmp16995)

tmp16997 := PrimCons(symshen_4comb, tmp16996)

tmp16998 := Call(__e, PrimFunc(symprotect), symSynCons)


tmp16999 := Call(__e, PrimFunc(symshen_4decons), V1229)


tmp17000 := PrimCons(symshen_4_5end_6, Nil)

tmp17001 := Call(__e, PrimFunc(symappend), tmp16999, tmp17000)


tmp17002 := PrimCons(V1228, Nil)

tmp17003 := PrimCons(symshen_4tlstream, tmp17002)

tmp17004 := PrimCons(V1231, Nil)

tmp17005 := PrimCons(V1230, tmp17004)

tmp17006 := PrimCons(tmp17003, tmp17005)

tmp17007 := PrimCons(symshen_4pushsemantics, tmp17006)

tmp17008 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V1227, tmp16998, tmp17001, tmp17007)


tmp17009 := PrimCons(tmp17008, Nil)

tmp17010 := PrimCons(tmp16997, tmp17009)

tmp17011 := PrimCons(tmp16990, tmp17010)

tmp17012 := PrimCons(symlet, tmp17011)

tmp17013 := PrimCons(symshen_4parse_1failure, Nil)

tmp17014 := PrimCons(tmp17013, Nil)

tmp17015 := PrimCons(tmp17012, tmp17014)

tmp17016 := PrimCons(tmp16989, tmp17015)

__e.Return(PrimCons(symif, tmp17016))
return


}, 5)

tmp17017 := Call(__e, ns2_1set, symshen_4conscode, tmp16987)


_ = tmp17017

tmp17018 := MakeNative(func(__e *ControlFlow) {
V1232 := __e.Get(1)
_ = V1232
tmp17045 := PrimIsPair(V1232)

var ifres17026 Obj

if True == tmp17045 {
tmp17043 := PrimHead(V1232)

tmp17044 := PrimEqual(symcons, tmp17043)

var ifres17028 Obj

if True == tmp17044 {
tmp17041 := PrimTail(V1232)

tmp17042 := PrimIsPair(tmp17041)

var ifres17030 Obj

if True == tmp17042 {
tmp17038 := PrimTail(V1232)

tmp17039 := PrimTail(tmp17038)

tmp17040 := PrimIsPair(tmp17039)

var ifres17032 Obj

if True == tmp17040 {
tmp17034 := PrimTail(V1232)

tmp17035 := PrimTail(tmp17034)

tmp17036 := PrimTail(tmp17035)

tmp17037 := PrimEqual(Nil, tmp17036)

var ifres17033 Obj

if True == tmp17037 {
ifres17033 = True


} else {
ifres17033 = False


}

ifres17032 = ifres17033


} else {
ifres17032 = False


}

var ifres17031 Obj

if True == ifres17032 {
ifres17031 = True


} else {
ifres17031 = False


}

ifres17030 = ifres17031


} else {
ifres17030 = False


}

var ifres17029 Obj

if True == ifres17030 {
ifres17029 = True


} else {
ifres17029 = False


}

ifres17028 = ifres17029


} else {
ifres17028 = False


}

var ifres17027 Obj

if True == ifres17028 {
ifres17027 = True


} else {
ifres17027 = False


}

ifres17026 = ifres17027


} else {
ifres17026 = False


}

if True == ifres17026 {
tmp17019 := PrimTail(V1232)

tmp17020 := PrimHead(tmp17019)

tmp17021 := PrimTail(V1232)

tmp17022 := PrimTail(tmp17021)

tmp17023 := PrimHead(tmp17022)

tmp17024 := Call(__e, PrimFunc(symshen_4decons), tmp17023)


__e.Return(PrimCons(tmp17020, tmp17024))
return


} else {
__e.Return(V1232)
return
}


}, 1)

tmp17046 := Call(__e, ns2_1set, symshen_4decons, tmp17018)


_ = tmp17046

tmp17047 := MakeNative(func(__e *ControlFlow) {
V1239 := __e.Get(1)
_ = V1239
tmp17064 := PrimIsPair(V1239)

var ifres17051 Obj

if True == tmp17064 {
tmp17062 := PrimHead(V1239)

tmp17063 := PrimIsPair(tmp17062)

var ifres17053 Obj

if True == tmp17063 {
tmp17060 := PrimTail(V1239)

tmp17061 := PrimIsPair(tmp17060)

var ifres17055 Obj

if True == tmp17061 {
tmp17057 := PrimTail(V1239)

tmp17058 := PrimTail(tmp17057)

tmp17059 := PrimEqual(Nil, tmp17058)

var ifres17056 Obj

if True == tmp17059 {
ifres17056 = True


} else {
ifres17056 = False


}

ifres17055 = ifres17056


} else {
ifres17055 = False


}

var ifres17054 Obj

if True == ifres17055 {
ifres17054 = True


} else {
ifres17054 = False


}

ifres17053 = ifres17054


} else {
ifres17053 = False


}

var ifres17052 Obj

if True == ifres17053 {
ifres17052 = True


} else {
ifres17052 = False


}

ifres17051 = ifres17052


} else {
ifres17051 = False


}

if True == ifres17051 {
tmp17048 := PrimHead(V1239)

tmp17049 := PrimHead(tmp17048)

__e.Return(PrimIsPair(tmp17049))
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17065 := Call(__e, ns2_1set, symshen_4ccons_2, tmp17047)


_ = tmp17065

tmp17066 := MakeNative(func(__e *ControlFlow) {
V1248 := __e.Get(1)
_ = V1248
tmp17072 := PrimIsPair(V1248)

var ifres17068 Obj

if True == tmp17072 {
tmp17070 := PrimHead(V1248)

tmp17071 := PrimIsPair(tmp17070)

var ifres17069 Obj

if True == tmp17071 {
ifres17069 = True


} else {
ifres17069 = False


}

ifres17068 = ifres17069


} else {
ifres17068 = False


}

if True == ifres17068 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp17073 := Call(__e, ns2_1set, symshen_4non_1empty_1stream_2, tmp17066)


_ = tmp17073

tmp17074 := MakeNative(func(__e *ControlFlow) {
V1249 := __e.Get(1)
_ = V1249
tmp17075 := PrimHead(V1249)

__e.Return(PrimHead(tmp17075))
return


}, 1)

tmp17076 := Call(__e, ns2_1set, symshen_4hds, tmp17074)


_ = tmp17076

tmp17077 := MakeNative(func(__e *ControlFlow) {
V1254 := __e.Get(1)
_ = V1254
tmp17095 := PrimIsPair(V1254)

var ifres17082 Obj

if True == tmp17095 {
tmp17093 := PrimHead(V1254)

tmp17094 := PrimIsPair(tmp17093)

var ifres17084 Obj

if True == tmp17094 {
tmp17091 := PrimTail(V1254)

tmp17092 := PrimIsPair(tmp17091)

var ifres17086 Obj

if True == tmp17092 {
tmp17088 := PrimTail(V1254)

tmp17089 := PrimTail(tmp17088)

tmp17090 := PrimEqual(Nil, tmp17089)

var ifres17087 Obj

if True == tmp17090 {
ifres17087 = True


} else {
ifres17087 = False


}

ifres17086 = ifres17087


} else {
ifres17086 = False


}

var ifres17085 Obj

if True == ifres17086 {
ifres17085 = True


} else {
ifres17085 = False


}

ifres17084 = ifres17085


} else {
ifres17084 = False


}

var ifres17083 Obj

if True == ifres17084 {
ifres17083 = True


} else {
ifres17083 = False


}

ifres17082 = ifres17083


} else {
ifres17082 = False


}

if True == ifres17082 {
tmp17078 := PrimHead(V1254)

tmp17079 := PrimHead(tmp17078)

tmp17080 := PrimTail(V1254)

__e.Return(PrimCons(tmp17079, tmp17080))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.hdstream\n")))
return
}


}, 1)

tmp17096 := Call(__e, ns2_1set, symshen_4hdstream, tmp17077)


_ = tmp17096

tmp17097 := MakeNative(func(__e *ControlFlow) {
V1255 := __e.Get(1)
_ = V1255
V1256 := __e.Get(2)
_ = V1256
tmp17098 := PrimCons(V1256, Nil)

__e.Return(PrimCons(V1255, tmp17098))
return


}, 2)

tmp17099 := Call(__e, ns2_1set, symshen_4comb, tmp17097)


_ = tmp17099

tmp17100 := MakeNative(func(__e *ControlFlow) {
V1261 := __e.Get(1)
_ = V1261
tmp17118 := PrimIsPair(V1261)

var ifres17105 Obj

if True == tmp17118 {
tmp17116 := PrimHead(V1261)

tmp17117 := PrimIsPair(tmp17116)

var ifres17107 Obj

if True == tmp17117 {
tmp17114 := PrimTail(V1261)

tmp17115 := PrimIsPair(tmp17114)

var ifres17109 Obj

if True == tmp17115 {
tmp17111 := PrimTail(V1261)

tmp17112 := PrimTail(tmp17111)

tmp17113 := PrimEqual(Nil, tmp17112)

var ifres17110 Obj

if True == tmp17113 {
ifres17110 = True


} else {
ifres17110 = False


}

ifres17109 = ifres17110


} else {
ifres17109 = False


}

var ifres17108 Obj

if True == ifres17109 {
ifres17108 = True


} else {
ifres17108 = False


}

ifres17107 = ifres17108


} else {
ifres17107 = False


}

var ifres17106 Obj

if True == ifres17107 {
ifres17106 = True


} else {
ifres17106 = False


}

ifres17105 = ifres17106


} else {
ifres17105 = False


}

if True == ifres17105 {
tmp17101 := PrimHead(V1261)

tmp17102 := PrimTail(tmp17101)

tmp17103 := PrimTail(V1261)

__e.Return(PrimCons(tmp17102, tmp17103))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.tlstream\n")))
return
}


}, 1)

tmp17119 := Call(__e, ns2_1set, symshen_4tlstream, tmp17100)


_ = tmp17119

tmp17120 := MakeNative(func(__e *ControlFlow) {
V1271 := __e.Get(1)
_ = V1271
V1272 := __e.Get(2)
_ = V1272
tmp17131 := PrimIsPair(V1271)

var ifres17122 Obj

if True == tmp17131 {
tmp17129 := PrimHead(V1271)

tmp17130 := PrimIsPair(tmp17129)

var ifres17124 Obj

if True == tmp17130 {
tmp17126 := PrimHead(V1271)

tmp17127 := PrimHead(tmp17126)

tmp17128 := PrimEqual(tmp17127, V1272)

var ifres17125 Obj

if True == tmp17128 {
ifres17125 = True


} else {
ifres17125 = False


}

ifres17124 = ifres17125


} else {
ifres17124 = False


}

var ifres17123 Obj

if True == ifres17124 {
ifres17123 = True


} else {
ifres17123 = False


}

ifres17122 = ifres17123


} else {
ifres17122 = False


}

if True == ifres17122 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp17132 := Call(__e, ns2_1set, symshen_4_ahd_2, tmp17120)


_ = tmp17132

tmp17133 := MakeNative(func(__e *ControlFlow) {
V1277 := __e.Get(1)
_ = V1277
tmp17151 := PrimIsPair(V1277)

var ifres17138 Obj

if True == tmp17151 {
tmp17149 := PrimHead(V1277)

tmp17150 := PrimIsPair(tmp17149)

var ifres17140 Obj

if True == tmp17150 {
tmp17147 := PrimTail(V1277)

tmp17148 := PrimIsPair(tmp17147)

var ifres17142 Obj

if True == tmp17148 {
tmp17144 := PrimTail(V1277)

tmp17145 := PrimTail(tmp17144)

tmp17146 := PrimEqual(Nil, tmp17145)

var ifres17143 Obj

if True == tmp17146 {
ifres17143 = True


} else {
ifres17143 = False


}

ifres17142 = ifres17143


} else {
ifres17142 = False


}

var ifres17141 Obj

if True == ifres17142 {
ifres17141 = True


} else {
ifres17141 = False


}

ifres17140 = ifres17141


} else {
ifres17140 = False


}

var ifres17139 Obj

if True == ifres17140 {
ifres17139 = True


} else {
ifres17139 = False


}

ifres17138 = ifres17139


} else {
ifres17138 = False


}

if True == ifres17138 {
tmp17134 := PrimHead(V1277)

tmp17135 := PrimTail(tmp17134)

tmp17136 := PrimTail(V1277)

__e.Return(PrimCons(tmp17135, tmp17136))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.tls\n")))
return
}


}, 1)

tmp17152 := Call(__e, ns2_1set, symshen_4tls, tmp17133)


_ = tmp17152

tmp17153 := MakeNative(func(__e *ControlFlow) {
V1280 := __e.Get(1)
_ = V1280
V1281 := __e.Get(2)
_ = V1281
V1282 := __e.Get(3)
_ = V1282
tmp17198 := PrimIsPair(V1282)

var ifres17172 Obj

if True == tmp17198 {
tmp17196 := PrimHead(V1282)

tmp17197 := PrimEqual(symshen_4pushsemantics, tmp17196)

var ifres17174 Obj

if True == tmp17197 {
tmp17194 := PrimTail(V1282)

tmp17195 := PrimIsPair(tmp17194)

var ifres17176 Obj

if True == tmp17195 {
tmp17191 := PrimTail(V1282)

tmp17192 := PrimTail(tmp17191)

tmp17193 := PrimIsPair(tmp17192)

var ifres17178 Obj

if True == tmp17193 {
tmp17187 := PrimTail(V1282)

tmp17188 := PrimTail(tmp17187)

tmp17189 := PrimTail(tmp17188)

tmp17190 := PrimIsPair(tmp17189)

var ifres17180 Obj

if True == tmp17190 {
tmp17182 := PrimTail(V1282)

tmp17183 := PrimTail(tmp17182)

tmp17184 := PrimTail(tmp17183)

tmp17185 := PrimTail(tmp17184)

tmp17186 := PrimEqual(Nil, tmp17185)

var ifres17181 Obj

if True == tmp17186 {
ifres17181 = True


} else {
ifres17181 = False


}

ifres17180 = ifres17181


} else {
ifres17180 = False


}

var ifres17179 Obj

if True == ifres17180 {
ifres17179 = True


} else {
ifres17179 = False


}

ifres17178 = ifres17179


} else {
ifres17178 = False


}

var ifres17177 Obj

if True == ifres17178 {
ifres17177 = True


} else {
ifres17177 = False


}

ifres17176 = ifres17177


} else {
ifres17176 = False


}

var ifres17175 Obj

if True == ifres17176 {
ifres17175 = True


} else {
ifres17175 = False


}

ifres17174 = ifres17175


} else {
ifres17174 = False


}

var ifres17173 Obj

if True == ifres17174 {
ifres17173 = True


} else {
ifres17173 = False


}

ifres17172 = ifres17173


} else {
ifres17172 = False


}

if True == ifres17172 {
tmp17154 := PrimTail(V1282)

tmp17155 := PrimHead(tmp17154)

tmp17156 := PrimTail(V1282)

tmp17157 := PrimTail(tmp17156)

tmp17158 := PrimHead(tmp17157)

tmp17159 := PrimTail(V1282)

tmp17160 := PrimTail(tmp17159)

tmp17161 := PrimTail(tmp17160)

tmp17162 := PrimHead(tmp17161)

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V1280, tmp17155, tmp17158, tmp17162)
return


} else {
tmp17163 := MakeNative(func(__e *ControlFlow) {
Process := __e.Get(1)
_ = Process
tmp17164 := MakeNative(func(__e *ControlFlow) {
Annotate := __e.Get(1)
_ = Annotate
tmp17165 := PrimCons(V1281, Nil)

tmp17166 := PrimCons(symshen_4in_1_6, tmp17165)

tmp17167 := PrimCons(Annotate, Nil)

tmp17168 := PrimCons(tmp17166, tmp17167)

__e.Return(PrimCons(symshen_4comb, tmp17168))
return


}, 1)

tmp17169 := Call(__e, PrimFunc(symshen_4use_1type_1info), V1280, Process)


__e.TailApply(tmp17164, tmp17169)
return


}, 1)

tmp17170 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V1282)


__e.TailApply(tmp17163, tmp17170)
return


}


}, 3)

tmp17199 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp17153)


_ = tmp17199

tmp17200 := MakeNative(func(__e *ControlFlow) {
V1286 := __e.Get(1)
_ = V1286
V1287 := __e.Get(2)
_ = V1287
tmp17420 := PrimIsPair(V1286)

var ifres17209 Obj

if True == tmp17420 {
tmp17418 := PrimHead(V1286)

tmp17419 := PrimEqual(sym_i, tmp17418)

var ifres17211 Obj

if True == tmp17419 {
tmp17416 := PrimTail(V1286)

tmp17417 := PrimIsPair(tmp17416)

var ifres17213 Obj

if True == tmp17417 {
tmp17413 := PrimTail(V1286)

tmp17414 := PrimHead(tmp17413)

tmp17415 := PrimIsPair(tmp17414)

var ifres17215 Obj

if True == tmp17415 {
tmp17409 := PrimTail(V1286)

tmp17410 := PrimHead(tmp17409)

tmp17411 := PrimHead(tmp17410)

tmp17412 := PrimEqual(symstr, tmp17411)

var ifres17217 Obj

if True == tmp17412 {
tmp17405 := PrimTail(V1286)

tmp17406 := PrimHead(tmp17405)

tmp17407 := PrimTail(tmp17406)

tmp17408 := PrimIsPair(tmp17407)

var ifres17219 Obj

if True == tmp17408 {
tmp17400 := PrimTail(V1286)

tmp17401 := PrimHead(tmp17400)

tmp17402 := PrimTail(tmp17401)

tmp17403 := PrimHead(tmp17402)

tmp17404 := PrimIsPair(tmp17403)

var ifres17221 Obj

if True == tmp17404 {
tmp17394 := PrimTail(V1286)

tmp17395 := PrimHead(tmp17394)

tmp17396 := PrimTail(tmp17395)

tmp17397 := PrimHead(tmp17396)

tmp17398 := PrimHead(tmp17397)

tmp17399 := PrimEqual(symlist, tmp17398)

var ifres17223 Obj

if True == tmp17399 {
tmp17388 := PrimTail(V1286)

tmp17389 := PrimHead(tmp17388)

tmp17390 := PrimTail(tmp17389)

tmp17391 := PrimHead(tmp17390)

tmp17392 := PrimTail(tmp17391)

tmp17393 := PrimIsPair(tmp17392)

var ifres17225 Obj

if True == tmp17393 {
tmp17381 := PrimTail(V1286)

tmp17382 := PrimHead(tmp17381)

tmp17383 := PrimTail(tmp17382)

tmp17384 := PrimHead(tmp17383)

tmp17385 := PrimTail(tmp17384)

tmp17386 := PrimTail(tmp17385)

tmp17387 := PrimEqual(Nil, tmp17386)

var ifres17227 Obj

if True == tmp17387 {
tmp17376 := PrimTail(V1286)

tmp17377 := PrimHead(tmp17376)

tmp17378 := PrimTail(tmp17377)

tmp17379 := PrimTail(tmp17378)

tmp17380 := PrimIsPair(tmp17379)

var ifres17229 Obj

if True == tmp17380 {
tmp17370 := PrimTail(V1286)

tmp17371 := PrimHead(tmp17370)

tmp17372 := PrimTail(tmp17371)

tmp17373 := PrimTail(tmp17372)

tmp17374 := PrimTail(tmp17373)

tmp17375 := PrimEqual(Nil, tmp17374)

var ifres17231 Obj

if True == tmp17375 {
tmp17367 := PrimTail(V1286)

tmp17368 := PrimTail(tmp17367)

tmp17369 := PrimIsPair(tmp17368)

var ifres17233 Obj

if True == tmp17369 {
tmp17363 := PrimTail(V1286)

tmp17364 := PrimTail(tmp17363)

tmp17365 := PrimHead(tmp17364)

tmp17366 := PrimEqual(sym_1_1_6, tmp17365)

var ifres17235 Obj

if True == tmp17366 {
tmp17359 := PrimTail(V1286)

tmp17360 := PrimTail(tmp17359)

tmp17361 := PrimTail(tmp17360)

tmp17362 := PrimIsPair(tmp17361)

var ifres17237 Obj

if True == tmp17362 {
tmp17354 := PrimTail(V1286)

tmp17355 := PrimTail(tmp17354)

tmp17356 := PrimTail(tmp17355)

tmp17357 := PrimHead(tmp17356)

tmp17358 := PrimIsPair(tmp17357)

var ifres17239 Obj

if True == tmp17358 {
tmp17348 := PrimTail(V1286)

tmp17349 := PrimTail(tmp17348)

tmp17350 := PrimTail(tmp17349)

tmp17351 := PrimHead(tmp17350)

tmp17352 := PrimHead(tmp17351)

tmp17353 := PrimEqual(symstr, tmp17352)

var ifres17241 Obj

if True == tmp17353 {
tmp17342 := PrimTail(V1286)

tmp17343 := PrimTail(tmp17342)

tmp17344 := PrimTail(tmp17343)

tmp17345 := PrimHead(tmp17344)

tmp17346 := PrimTail(tmp17345)

tmp17347 := PrimIsPair(tmp17346)

var ifres17243 Obj

if True == tmp17347 {
tmp17335 := PrimTail(V1286)

tmp17336 := PrimTail(tmp17335)

tmp17337 := PrimTail(tmp17336)

tmp17338 := PrimHead(tmp17337)

tmp17339 := PrimTail(tmp17338)

tmp17340 := PrimHead(tmp17339)

tmp17341 := PrimIsPair(tmp17340)

var ifres17245 Obj

if True == tmp17341 {
tmp17327 := PrimTail(V1286)

tmp17328 := PrimTail(tmp17327)

tmp17329 := PrimTail(tmp17328)

tmp17330 := PrimHead(tmp17329)

tmp17331 := PrimTail(tmp17330)

tmp17332 := PrimHead(tmp17331)

tmp17333 := PrimHead(tmp17332)

tmp17334 := PrimEqual(symlist, tmp17333)

var ifres17247 Obj

if True == tmp17334 {
tmp17319 := PrimTail(V1286)

tmp17320 := PrimTail(tmp17319)

tmp17321 := PrimTail(tmp17320)

tmp17322 := PrimHead(tmp17321)

tmp17323 := PrimTail(tmp17322)

tmp17324 := PrimHead(tmp17323)

tmp17325 := PrimTail(tmp17324)

tmp17326 := PrimIsPair(tmp17325)

var ifres17249 Obj

if True == tmp17326 {
tmp17310 := PrimTail(V1286)

tmp17311 := PrimTail(tmp17310)

tmp17312 := PrimTail(tmp17311)

tmp17313 := PrimHead(tmp17312)

tmp17314 := PrimTail(tmp17313)

tmp17315 := PrimHead(tmp17314)

tmp17316 := PrimTail(tmp17315)

tmp17317 := PrimTail(tmp17316)

tmp17318 := PrimEqual(Nil, tmp17317)

var ifres17251 Obj

if True == tmp17318 {
tmp17303 := PrimTail(V1286)

tmp17304 := PrimTail(tmp17303)

tmp17305 := PrimTail(tmp17304)

tmp17306 := PrimHead(tmp17305)

tmp17307 := PrimTail(tmp17306)

tmp17308 := PrimTail(tmp17307)

tmp17309 := PrimIsPair(tmp17308)

var ifres17253 Obj

if True == tmp17309 {
tmp17295 := PrimTail(V1286)

tmp17296 := PrimTail(tmp17295)

tmp17297 := PrimTail(tmp17296)

tmp17298 := PrimHead(tmp17297)

tmp17299 := PrimTail(tmp17298)

tmp17300 := PrimTail(tmp17299)

tmp17301 := PrimTail(tmp17300)

tmp17302 := PrimEqual(Nil, tmp17301)

var ifres17255 Obj

if True == tmp17302 {
tmp17290 := PrimTail(V1286)

tmp17291 := PrimTail(tmp17290)

tmp17292 := PrimTail(tmp17291)

tmp17293 := PrimTail(tmp17292)

tmp17294 := PrimIsPair(tmp17293)

var ifres17257 Obj

if True == tmp17294 {
tmp17284 := PrimTail(V1286)

tmp17285 := PrimTail(tmp17284)

tmp17286 := PrimTail(tmp17285)

tmp17287 := PrimTail(tmp17286)

tmp17288 := PrimHead(tmp17287)

tmp17289 := PrimEqual(sym_j, tmp17288)

var ifres17259 Obj

if True == tmp17289 {
tmp17278 := PrimTail(V1286)

tmp17279 := PrimTail(tmp17278)

tmp17280 := PrimTail(tmp17279)

tmp17281 := PrimTail(tmp17280)

tmp17282 := PrimTail(tmp17281)

tmp17283 := PrimEqual(Nil, tmp17282)

var ifres17261 Obj

if True == tmp17283 {
tmp17263 := PrimTail(V1286)

tmp17264 := PrimHead(tmp17263)

tmp17265 := PrimTail(tmp17264)

tmp17266 := PrimHead(tmp17265)

tmp17267 := PrimTail(tmp17266)

tmp17268 := PrimHead(tmp17267)

tmp17269 := PrimTail(V1286)

tmp17270 := PrimTail(tmp17269)

tmp17271 := PrimTail(tmp17270)

tmp17272 := PrimHead(tmp17271)

tmp17273 := PrimTail(tmp17272)

tmp17274 := PrimHead(tmp17273)

tmp17275 := PrimTail(tmp17274)

tmp17276 := PrimHead(tmp17275)

tmp17277 := PrimEqual(tmp17268, tmp17276)

var ifres17262 Obj

if True == tmp17277 {
ifres17262 = True


} else {
ifres17262 = False


}

ifres17261 = ifres17262


} else {
ifres17261 = False


}

var ifres17260 Obj

if True == ifres17261 {
ifres17260 = True


} else {
ifres17260 = False


}

ifres17259 = ifres17260


} else {
ifres17259 = False


}

var ifres17258 Obj

if True == ifres17259 {
ifres17258 = True


} else {
ifres17258 = False


}

ifres17257 = ifres17258


} else {
ifres17257 = False


}

var ifres17256 Obj

if True == ifres17257 {
ifres17256 = True


} else {
ifres17256 = False


}

ifres17255 = ifres17256


} else {
ifres17255 = False


}

var ifres17254 Obj

if True == ifres17255 {
ifres17254 = True


} else {
ifres17254 = False


}

ifres17253 = ifres17254


} else {
ifres17253 = False


}

var ifres17252 Obj

if True == ifres17253 {
ifres17252 = True


} else {
ifres17252 = False


}

ifres17251 = ifres17252


} else {
ifres17251 = False


}

var ifres17250 Obj

if True == ifres17251 {
ifres17250 = True


} else {
ifres17250 = False


}

ifres17249 = ifres17250


} else {
ifres17249 = False


}

var ifres17248 Obj

if True == ifres17249 {
ifres17248 = True


} else {
ifres17248 = False


}

ifres17247 = ifres17248


} else {
ifres17247 = False


}

var ifres17246 Obj

if True == ifres17247 {
ifres17246 = True


} else {
ifres17246 = False


}

ifres17245 = ifres17246


} else {
ifres17245 = False


}

var ifres17244 Obj

if True == ifres17245 {
ifres17244 = True


} else {
ifres17244 = False


}

ifres17243 = ifres17244


} else {
ifres17243 = False


}

var ifres17242 Obj

if True == ifres17243 {
ifres17242 = True


} else {
ifres17242 = False


}

ifres17241 = ifres17242


} else {
ifres17241 = False


}

var ifres17240 Obj

if True == ifres17241 {
ifres17240 = True


} else {
ifres17240 = False


}

ifres17239 = ifres17240


} else {
ifres17239 = False


}

var ifres17238 Obj

if True == ifres17239 {
ifres17238 = True


} else {
ifres17238 = False


}

ifres17237 = ifres17238


} else {
ifres17237 = False


}

var ifres17236 Obj

if True == ifres17237 {
ifres17236 = True


} else {
ifres17236 = False


}

ifres17235 = ifres17236


} else {
ifres17235 = False


}

var ifres17234 Obj

if True == ifres17235 {
ifres17234 = True


} else {
ifres17234 = False


}

ifres17233 = ifres17234


} else {
ifres17233 = False


}

var ifres17232 Obj

if True == ifres17233 {
ifres17232 = True


} else {
ifres17232 = False


}

ifres17231 = ifres17232


} else {
ifres17231 = False


}

var ifres17230 Obj

if True == ifres17231 {
ifres17230 = True


} else {
ifres17230 = False


}

ifres17229 = ifres17230


} else {
ifres17229 = False


}

var ifres17228 Obj

if True == ifres17229 {
ifres17228 = True


} else {
ifres17228 = False


}

ifres17227 = ifres17228


} else {
ifres17227 = False


}

var ifres17226 Obj

if True == ifres17227 {
ifres17226 = True


} else {
ifres17226 = False


}

ifres17225 = ifres17226


} else {
ifres17225 = False


}

var ifres17224 Obj

if True == ifres17225 {
ifres17224 = True


} else {
ifres17224 = False


}

ifres17223 = ifres17224


} else {
ifres17223 = False


}

var ifres17222 Obj

if True == ifres17223 {
ifres17222 = True


} else {
ifres17222 = False


}

ifres17221 = ifres17222


} else {
ifres17221 = False


}

var ifres17220 Obj

if True == ifres17221 {
ifres17220 = True


} else {
ifres17220 = False


}

ifres17219 = ifres17220


} else {
ifres17219 = False


}

var ifres17218 Obj

if True == ifres17219 {
ifres17218 = True


} else {
ifres17218 = False


}

ifres17217 = ifres17218


} else {
ifres17217 = False


}

var ifres17216 Obj

if True == ifres17217 {
ifres17216 = True


} else {
ifres17216 = False


}

ifres17215 = ifres17216


} else {
ifres17215 = False


}

var ifres17214 Obj

if True == ifres17215 {
ifres17214 = True


} else {
ifres17214 = False


}

ifres17213 = ifres17214


} else {
ifres17213 = False


}

var ifres17212 Obj

if True == ifres17213 {
ifres17212 = True


} else {
ifres17212 = False


}

ifres17211 = ifres17212


} else {
ifres17211 = False


}

var ifres17210 Obj

if True == ifres17211 {
ifres17210 = True


} else {
ifres17210 = False


}

ifres17209 = ifres17210


} else {
ifres17209 = False


}

if True == ifres17209 {
tmp17201 := PrimTail(V1286)

tmp17202 := PrimTail(tmp17201)

tmp17203 := PrimTail(tmp17202)

tmp17204 := PrimHead(tmp17203)

tmp17205 := PrimTail(tmp17204)

tmp17206 := PrimTail(tmp17205)

tmp17207 := PrimCons(V1287, tmp17206)

__e.Return(PrimCons(symtype, tmp17207))
return


} else {
__e.Return(V1287)
return
}


}, 2)

tmp17421 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp17200)


_ = tmp17421

tmp17422 := MakeNative(func(__e *ControlFlow) {
V1288 := __e.Get(1)
_ = V1288
tmp17430 := PrimIsPair(V1288)

if True == tmp17430 {
tmp17423 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp17423, V1288)
return


} else {
tmp17428 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V1288)


if True == tmp17428 {
tmp17424 := Call(__e, PrimFunc(symprotect), symParse)


tmp17425 := Call(__e, PrimFunc(symconcat), tmp17424, V1288)


tmp17426 := PrimCons(tmp17425, Nil)

__e.Return(PrimCons(symshen_4_5_1out, tmp17426))
return


} else {
__e.Return(V1288)
return
}


}


}, 1)

tmp17431 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp17422)


_ = tmp17431

tmp17432 := MakeNative(func(__e *ControlFlow) {
V1293 := __e.Get(1)
_ = V1293
tmp17444 := PrimIsPair(V1293)

var ifres17435 Obj

if True == tmp17444 {
tmp17442 := PrimTail(V1293)

tmp17443 := PrimIsPair(tmp17442)

var ifres17437 Obj

if True == tmp17443 {
tmp17439 := PrimTail(V1293)

tmp17440 := PrimTail(tmp17439)

tmp17441 := PrimEqual(Nil, tmp17440)

var ifres17438 Obj

if True == tmp17441 {
ifres17438 = True


} else {
ifres17438 = False


}

ifres17437 = ifres17438


} else {
ifres17437 = False


}

var ifres17436 Obj

if True == ifres17437 {
ifres17436 = True


} else {
ifres17436 = False


}

ifres17435 = ifres17436


} else {
ifres17435 = False


}

if True == ifres17435 {
tmp17433 := PrimTail(V1293)

__e.Return(PrimHead(tmp17433))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.<-out\n")))
return
}


}, 1)

tmp17445 := Call(__e, ns2_1set, symshen_4_5_1out, tmp17432)


_ = tmp17445

tmp17446 := MakeNative(func(__e *ControlFlow) {
V1298 := __e.Get(1)
_ = V1298
tmp17457 := PrimIsPair(V1298)

var ifres17448 Obj

if True == tmp17457 {
tmp17455 := PrimTail(V1298)

tmp17456 := PrimIsPair(tmp17455)

var ifres17450 Obj

if True == tmp17456 {
tmp17452 := PrimTail(V1298)

tmp17453 := PrimTail(tmp17452)

tmp17454 := PrimEqual(Nil, tmp17453)

var ifres17451 Obj

if True == tmp17454 {
ifres17451 = True


} else {
ifres17451 = False


}

ifres17450 = ifres17451


} else {
ifres17450 = False


}

var ifres17449 Obj

if True == ifres17450 {
ifres17449 = True


} else {
ifres17449 = False


}

ifres17448 = ifres17449


} else {
ifres17448 = False


}

if True == ifres17448 {
__e.Return(PrimHead(V1298))
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.in->\n")))
return
}


}, 1)

tmp17458 := Call(__e, ns2_1set, symshen_4in_1_6, tmp17446)


_ = tmp17458

tmp17459 := MakeNative(func(__e *ControlFlow) {
V1303 := __e.Get(1)
_ = V1303
tmp17472 := PrimIsPair(V1303)

var ifres17463 Obj

if True == tmp17472 {
tmp17470 := PrimTail(V1303)

tmp17471 := PrimIsPair(tmp17470)

var ifres17465 Obj

if True == tmp17471 {
tmp17467 := PrimTail(V1303)

tmp17468 := PrimTail(tmp17467)

tmp17469 := PrimEqual(Nil, tmp17468)

var ifres17466 Obj

if True == tmp17469 {
ifres17466 = True


} else {
ifres17466 = False


}

ifres17465 = ifres17466


} else {
ifres17465 = False


}

var ifres17464 Obj

if True == ifres17465 {
ifres17464 = True


} else {
ifres17464 = False


}

ifres17463 = ifres17464


} else {
ifres17463 = False


}

if True == ifres17463 {
tmp17460 := PrimHead(V1303)

tmp17461 := PrimCons(tmp17460, Nil)

__e.Return(PrimCons(Nil, tmp17461))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in <!>\n")))
return
}


}, 1)

tmp17473 := Call(__e, ns2_1set, sym_5_b_6, tmp17459)


_ = tmp17473

tmp17474 := MakeNative(func(__e *ControlFlow) {
V1308 := __e.Get(1)
_ = V1308
tmp17487 := PrimIsPair(V1308)

var ifres17478 Obj

if True == tmp17487 {
tmp17485 := PrimTail(V1308)

tmp17486 := PrimIsPair(tmp17485)

var ifres17480 Obj

if True == tmp17486 {
tmp17482 := PrimTail(V1308)

tmp17483 := PrimTail(tmp17482)

tmp17484 := PrimEqual(Nil, tmp17483)

var ifres17481 Obj

if True == tmp17484 {
ifres17481 = True


} else {
ifres17481 = False


}

ifres17480 = ifres17481


} else {
ifres17480 = False


}

var ifres17479 Obj

if True == ifres17480 {
ifres17479 = True


} else {
ifres17479 = False


}

ifres17478 = ifres17479


} else {
ifres17478 = False


}

if True == ifres17478 {
tmp17475 := PrimHead(V1308)

tmp17476 := PrimCons(Nil, Nil)

__e.Return(PrimCons(tmp17475, tmp17476))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in <e>\n")))
return
}


}, 1)

tmp17488 := Call(__e, ns2_1set, sym_5e_6, tmp17474)


_ = tmp17488

tmp17489 := MakeNative(func(__e *ControlFlow) {
V1311 := __e.Get(1)
_ = V1311
tmp17504 := PrimIsPair(V1311)

var ifres17491 Obj

if True == tmp17504 {
tmp17502 := PrimHead(V1311)

tmp17503 := PrimEqual(Nil, tmp17502)

var ifres17493 Obj

if True == tmp17503 {
tmp17500 := PrimTail(V1311)

tmp17501 := PrimIsPair(tmp17500)

var ifres17495 Obj

if True == tmp17501 {
tmp17497 := PrimTail(V1311)

tmp17498 := PrimTail(tmp17497)

tmp17499 := PrimEqual(Nil, tmp17498)

var ifres17496 Obj

if True == tmp17499 {
ifres17496 = True


} else {
ifres17496 = False


}

ifres17495 = ifres17496


} else {
ifres17495 = False


}

var ifres17494 Obj

if True == ifres17495 {
ifres17494 = True


} else {
ifres17494 = False


}

ifres17493 = ifres17494


} else {
ifres17493 = False


}

var ifres17492 Obj

if True == ifres17493 {
ifres17492 = True


} else {
ifres17492 = False


}

ifres17491 = ifres17492


} else {
ifres17491 = False


}

if True == ifres17491 {
__e.Return(V1311)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, symshen_4_5end_6, tmp17489)
return




}, 0)

