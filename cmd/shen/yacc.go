package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17216 := MakeNative(func(__e *ControlFlow) {
V107 := __e.Get(1)
_ = V107
V108 := __e.Get(2)
_ = V108
tmp17217 := MakeNative(func(__e *ControlFlow) {
W109 := __e.Get(1)
_ = W109
tmp17226 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W109)


if True == tmp17226 {
__e.Return(PrimSimpleError(MakeString("parse failure\n")))
return
} else {
tmp17223 := Call(__e, PrimFunc(symshen_4in_1_6), W109)


tmp17224 := PrimIsPair(tmp17223)

if True == tmp17224 {
tmp17218 := Call(__e, PrimFunc(symshen_4in_1_6), W109)


tmp17219 := PrimHead(tmp17218)

tmp17220 := Call(__e, PrimFunc(symshen_4app), tmp17219, MakeString(" ..."), symshen_4s)


tmp17221 := PrimStringConcat(MakeString("syntax error here: "), tmp17220)

__e.Return(PrimSimpleError(tmp17221))
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), W109)
return
}


}


}, 1)

tmp17227 := Call(__e, V107, V108)


__e.TailApply(tmp17217, tmp17227)
return


}, 2)

tmp17228 := Call(__e, ns2_1set, symcompile, tmp17216)


_ = tmp17228

tmp17229 := MakeNative(func(__e *ControlFlow) {
V110 := __e.Get(1)
_ = V110
tmp17230 := Call(__e, PrimFunc(symfail))


__e.Return(PrimEqual(V110, tmp17230))
return


}, 1)

tmp17231 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17229)


_ = tmp17231

tmp17232 := MakeNative(func(__e *ControlFlow) {
V113 := __e.Get(1)
_ = V113
tmp17245 := PrimIsPair(V113)

var ifres17236 Obj

if True == tmp17245 {
tmp17243 := PrimTail(V113)

tmp17244 := PrimIsPair(tmp17243)

var ifres17238 Obj

if True == tmp17244 {
tmp17240 := PrimTail(V113)

tmp17241 := PrimTail(tmp17240)

tmp17242 := PrimEqual(Nil, tmp17241)

var ifres17239 Obj

if True == tmp17242 {
ifres17239 = True


} else {
ifres17239 = False


}

ifres17238 = ifres17239


} else {
ifres17238 = False


}

var ifres17237 Obj

if True == ifres17238 {
ifres17237 = True


} else {
ifres17237 = False


}

ifres17236 = ifres17237


} else {
ifres17236 = False


}

if True == ifres17236 {
tmp17233 := PrimTail(V113)

__e.Return(PrimHead(tmp17233))
return


} else {
tmp17234 := Call(__e, PrimFunc(symshen_4app), V113, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp17234))
return


}


}, 1)

tmp17246 := Call(__e, ns2_1set, symshen_4objectcode, tmp17232)


_ = tmp17246

tmp17247 := MakeNative(func(__e *ControlFlow) {
V114 := __e.Get(1)
_ = V114
tmp17248 := MakeNative(func(__e *ControlFlow) {
Z115 := __e.Get(1)
_ = Z115
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z115)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17248, V114)
return


}, 1)

tmp17249 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17247)


_ = tmp17249

tmp17250 := MakeNative(func(__e *ControlFlow) {
V116 := __e.Get(1)
_ = V116
tmp17251 := MakeNative(func(__e *ControlFlow) {
W117 := __e.Get(1)
_ = W117
tmp17253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W117)


if True == tmp17253 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W117)
return
}


}, 1)

tmp17289 := PrimIsPair(V116)

var ifres17254 Obj

if True == tmp17289 {
tmp17255 := MakeNative(func(__e *ControlFlow) {
W118 := __e.Get(1)
_ = W118
tmp17256 := MakeNative(func(__e *ControlFlow) {
W119 := __e.Get(1)
_ = W119
tmp17257 := MakeNative(func(__e *ControlFlow) {
W120 := __e.Get(1)
_ = W120
tmp17283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W120)


if True == tmp17283 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17258 := MakeNative(func(__e *ControlFlow) {
W121 := __e.Get(1)
_ = W121
tmp17259 := MakeNative(func(__e *ControlFlow) {
W122 := __e.Get(1)
_ = W122
tmp17260 := MakeNative(func(__e *ControlFlow) {
W123 := __e.Get(1)
_ = W123
tmp17278 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W123)


if True == tmp17278 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17261 := MakeNative(func(__e *ControlFlow) {
W124 := __e.Get(1)
_ = W124
tmp17262 := MakeNative(func(__e *ControlFlow) {
W125 := __e.Get(1)
_ = W125
tmp17263 := MakeNative(func(__e *ControlFlow) {
W126 := __e.Get(1)
_ = W126
tmp17264 := MakeNative(func(__e *ControlFlow) {
W127 := __e.Get(1)
_ = W127
__e.Return(W127)
return
}, 1)

tmp17265 := PrimCons(W118, Nil)

tmp17266 := PrimCons(symdefine, tmp17265)

tmp17267 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), W121, W126, W124)


tmp17268 := PrimCons(tmp17267, Nil)

tmp17269 := PrimCons(sym_1_6, tmp17268)

tmp17270 := PrimCons(W126, tmp17269)

tmp17271 := Call(__e, PrimFunc(symappend), W121, tmp17270)


tmp17272 := Call(__e, PrimFunc(symappend), tmp17266, tmp17271)


__e.TailApply(tmp17264, tmp17272)
return


}, 1)

tmp17273 := Call(__e, PrimFunc(symgensym), symS)


tmp17274 := Call(__e, tmp17263, tmp17273)


__e.TailApply(PrimFunc(symshen_4comb), W125, tmp17274)
return


}, 1)

tmp17275 := Call(__e, PrimFunc(symshen_4in_1_6), W123)


__e.TailApply(tmp17262, tmp17275)
return


}, 1)

tmp17276 := Call(__e, PrimFunc(symshen_4_5_1out), W123)


__e.TailApply(tmp17261, tmp17276)
return


}


}, 1)

tmp17279 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W122)


__e.TailApply(tmp17260, tmp17279)
return


}, 1)

tmp17280 := Call(__e, PrimFunc(symshen_4in_1_6), W120)


__e.TailApply(tmp17259, tmp17280)
return


}, 1)

tmp17281 := Call(__e, PrimFunc(symshen_4_5_1out), W120)


__e.TailApply(tmp17258, tmp17281)
return


}


}, 1)

tmp17284 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), W119)


__e.TailApply(tmp17257, tmp17284)
return


}, 1)

tmp17285 := Call(__e, PrimFunc(symtail), V116)


__e.TailApply(tmp17256, tmp17285)
return


}, 1)

tmp17286 := Call(__e, PrimFunc(symhead), V116)


tmp17287 := Call(__e, tmp17255, tmp17286)


ifres17254 = tmp17287


} else {
tmp17288 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17254 = tmp17288


}

__e.TailApply(tmp17251, ifres17254)
return


}, 1)

tmp17290 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17250)


_ = tmp17290

tmp17291 := MakeNative(func(__e *ControlFlow) {
V128 := __e.Get(1)
_ = V128
tmp17292 := MakeNative(func(__e *ControlFlow) {
W129 := __e.Get(1)
_ = W129
tmp17304 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W129)


if True == tmp17304 {
tmp17293 := MakeNative(func(__e *ControlFlow) {
W144 := __e.Get(1)
_ = W144
tmp17295 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W144)


if True == tmp17295 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W144)
return
}


}, 1)

tmp17296 := MakeNative(func(__e *ControlFlow) {
W145 := __e.Get(1)
_ = W145
tmp17300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W145)


if True == tmp17300 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17297 := MakeNative(func(__e *ControlFlow) {
W146 := __e.Get(1)
_ = W146
__e.TailApply(PrimFunc(symshen_4comb), W146, Nil)
return
}, 1)

tmp17298 := Call(__e, PrimFunc(symshen_4in_1_6), W145)


__e.TailApply(tmp17297, tmp17298)
return


}


}, 1)

tmp17301 := Call(__e, PrimFunc(sym_5e_6), V128)


tmp17302 := Call(__e, tmp17296, tmp17301)


__e.TailApply(tmp17293, tmp17302)
return


} else {
__e.Return(W129)
return
}


}, 1)

tmp17367 := PrimIsPair(V128)

var ifres17305 Obj

if True == tmp17367 {
tmp17306 := MakeNative(func(__e *ControlFlow) {
W130 := __e.Get(1)
_ = W130
tmp17307 := MakeNative(func(__e *ControlFlow) {
W131 := __e.Get(1)
_ = W131
tmp17362 := Call(__e, PrimFunc(symshen_4ccons_2), W131)


if True == tmp17362 {
tmp17308 := MakeNative(func(__e *ControlFlow) {
W132 := __e.Get(1)
_ = W132
tmp17309 := MakeNative(func(__e *ControlFlow) {
W133 := __e.Get(1)
_ = W133
tmp17358 := Call(__e, PrimFunc(symshen_4hds_a_2), W132, symlist)


if True == tmp17358 {
tmp17310 := MakeNative(func(__e *ControlFlow) {
W134 := __e.Get(1)
_ = W134
tmp17355 := PrimIsPair(W134)

if True == tmp17355 {
tmp17311 := MakeNative(func(__e *ControlFlow) {
W135 := __e.Get(1)
_ = W135
tmp17312 := MakeNative(func(__e *ControlFlow) {
W136 := __e.Get(1)
_ = W136
tmp17313 := MakeNative(func(__e *ControlFlow) {
W137 := __e.Get(1)
_ = W137
tmp17350 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W137)


if True == tmp17350 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17314 := MakeNative(func(__e *ControlFlow) {
W138 := __e.Get(1)
_ = W138
tmp17347 := Call(__e, PrimFunc(symshen_4hds_a_2), W133, sym_a_a_6)


if True == tmp17347 {
tmp17315 := MakeNative(func(__e *ControlFlow) {
W139 := __e.Get(1)
_ = W139
tmp17344 := PrimIsPair(W139)

if True == tmp17344 {
tmp17316 := MakeNative(func(__e *ControlFlow) {
W140 := __e.Get(1)
_ = W140
tmp17317 := MakeNative(func(__e *ControlFlow) {
W141 := __e.Get(1)
_ = W141
tmp17340 := PrimIsPair(W141)

if True == tmp17340 {
tmp17318 := MakeNative(func(__e *ControlFlow) {
W142 := __e.Get(1)
_ = W142
tmp17319 := MakeNative(func(__e *ControlFlow) {
W143 := __e.Get(1)
_ = W143
tmp17336 := PrimEqual(sym_i, W130)

var ifres17333 Obj

if True == tmp17336 {
tmp17335 := PrimEqual(sym_j, W142)

var ifres17334 Obj

if True == tmp17335 {
ifres17334 = True


} else {
ifres17334 = False


}

ifres17333 = ifres17334


} else {
ifres17333 = False


}

if True == ifres17333 {
tmp17320 := PrimCons(W135, Nil)

tmp17321 := PrimCons(symlist, tmp17320)

tmp17322 := PrimCons(W135, Nil)

tmp17323 := PrimCons(symlist, tmp17322)

tmp17324 := PrimCons(W140, Nil)

tmp17325 := PrimCons(tmp17323, tmp17324)

tmp17326 := PrimCons(symstr, tmp17325)

tmp17327 := PrimCons(sym_j, Nil)

tmp17328 := PrimCons(tmp17326, tmp17327)

tmp17329 := PrimCons(sym_1_1_6, tmp17328)

tmp17330 := PrimCons(tmp17321, tmp17329)

tmp17331 := PrimCons(sym_i, tmp17330)

__e.TailApply(PrimFunc(symshen_4comb), W143, tmp17331)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17337 := Call(__e, PrimFunc(symtail), W141)


__e.TailApply(tmp17319, tmp17337)
return


}, 1)

tmp17338 := Call(__e, PrimFunc(symhead), W141)


__e.TailApply(tmp17318, tmp17338)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17341 := Call(__e, PrimFunc(symtail), W139)


__e.TailApply(tmp17317, tmp17341)
return


}, 1)

tmp17342 := Call(__e, PrimFunc(symhead), W139)


__e.TailApply(tmp17316, tmp17342)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17345 := Call(__e, PrimFunc(symtail), W133)


__e.TailApply(tmp17315, tmp17345)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17348 := Call(__e, PrimFunc(symshen_4in_1_6), W137)


__e.TailApply(tmp17314, tmp17348)
return


}


}, 1)

tmp17351 := Call(__e, PrimFunc(sym_5end_6), W136)


__e.TailApply(tmp17313, tmp17351)
return


}, 1)

tmp17352 := Call(__e, PrimFunc(symtail), W134)


__e.TailApply(tmp17312, tmp17352)
return


}, 1)

tmp17353 := Call(__e, PrimFunc(symhead), W134)


__e.TailApply(tmp17311, tmp17353)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17356 := Call(__e, PrimFunc(symtail), W132)


__e.TailApply(tmp17310, tmp17356)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17359 := Call(__e, PrimFunc(symtail), W131)


__e.TailApply(tmp17309, tmp17359)
return


}, 1)

tmp17360 := Call(__e, PrimFunc(symhead), W131)


__e.TailApply(tmp17308, tmp17360)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17363 := Call(__e, PrimFunc(symtail), V128)


__e.TailApply(tmp17307, tmp17363)
return


}, 1)

tmp17364 := Call(__e, PrimFunc(symhead), V128)


tmp17365 := Call(__e, tmp17306, tmp17364)


ifres17305 = tmp17365


} else {
tmp17366 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17305 = tmp17366


}

__e.TailApply(tmp17292, ifres17305)
return


}, 1)

tmp17368 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17291)


_ = tmp17368

tmp17369 := MakeNative(func(__e *ControlFlow) {
V147 := __e.Get(1)
_ = V147
tmp17370 := MakeNative(func(__e *ControlFlow) {
W148 := __e.Get(1)
_ = W148
tmp17389 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W148)


if True == tmp17389 {
tmp17371 := MakeNative(func(__e *ControlFlow) {
W155 := __e.Get(1)
_ = W155
tmp17373 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W155)


if True == tmp17373 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W155)
return
}


}, 1)

tmp17374 := MakeNative(func(__e *ControlFlow) {
W156 := __e.Get(1)
_ = W156
tmp17385 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W156)


if True == tmp17385 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17375 := MakeNative(func(__e *ControlFlow) {
W157 := __e.Get(1)
_ = W157
tmp17376 := MakeNative(func(__e *ControlFlow) {
W158 := __e.Get(1)
_ = W158
tmp17381 := Call(__e, PrimFunc(symempty_2), W157)


var ifres17377 Obj

if True == tmp17381 {
ifres17377 = Nil


} else {
tmp17378 := Call(__e, PrimFunc(symshen_4app), W157, MakeString("\n ..."), symshen_4r)


tmp17379 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp17378)

tmp17380 := PrimSimpleError(tmp17379)

ifres17377 = tmp17380


}

__e.TailApply(PrimFunc(symshen_4comb), W158, ifres17377)
return


}, 1)

tmp17382 := Call(__e, PrimFunc(symshen_4in_1_6), W156)


__e.TailApply(tmp17376, tmp17382)
return


}, 1)

tmp17383 := Call(__e, PrimFunc(symshen_4_5_1out), W156)


__e.TailApply(tmp17375, tmp17383)
return


}


}, 1)

tmp17386 := Call(__e, PrimFunc(sym_5_b_6), V147)


tmp17387 := Call(__e, tmp17374, tmp17386)


__e.TailApply(tmp17371, tmp17387)
return


} else {
__e.Return(W148)
return
}


}, 1)

tmp17390 := MakeNative(func(__e *ControlFlow) {
W149 := __e.Get(1)
_ = W149
tmp17405 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W149)


if True == tmp17405 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17391 := MakeNative(func(__e *ControlFlow) {
W150 := __e.Get(1)
_ = W150
tmp17392 := MakeNative(func(__e *ControlFlow) {
W151 := __e.Get(1)
_ = W151
tmp17393 := MakeNative(func(__e *ControlFlow) {
W152 := __e.Get(1)
_ = W152
tmp17400 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W152)


if True == tmp17400 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17394 := MakeNative(func(__e *ControlFlow) {
W153 := __e.Get(1)
_ = W153
tmp17395 := MakeNative(func(__e *ControlFlow) {
W154 := __e.Get(1)
_ = W154
tmp17396 := PrimCons(W150, W153)

__e.TailApply(PrimFunc(symshen_4comb), W154, tmp17396)
return


}, 1)

tmp17397 := Call(__e, PrimFunc(symshen_4in_1_6), W152)


__e.TailApply(tmp17395, tmp17397)
return


}, 1)

tmp17398 := Call(__e, PrimFunc(symshen_4_5_1out), W152)


__e.TailApply(tmp17394, tmp17398)
return


}


}, 1)

tmp17401 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W151)


__e.TailApply(tmp17393, tmp17401)
return


}, 1)

tmp17402 := Call(__e, PrimFunc(symshen_4in_1_6), W149)


__e.TailApply(tmp17392, tmp17402)
return


}, 1)

tmp17403 := Call(__e, PrimFunc(symshen_4_5_1out), W149)


__e.TailApply(tmp17391, tmp17403)
return


}


}, 1)

tmp17406 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V147)


tmp17407 := Call(__e, tmp17390, tmp17406)


__e.TailApply(tmp17370, tmp17407)
return


}, 1)

tmp17408 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17369)


_ = tmp17408

tmp17409 := MakeNative(func(__e *ControlFlow) {
V159 := __e.Get(1)
_ = V159
tmp17410 := MakeNative(func(__e *ControlFlow) {
W160 := __e.Get(1)
_ = W160
tmp17433 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W160)


if True == tmp17433 {
tmp17411 := MakeNative(func(__e *ControlFlow) {
W169 := __e.Get(1)
_ = W169
tmp17413 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W169)


if True == tmp17413 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W169)
return
}


}, 1)

tmp17414 := MakeNative(func(__e *ControlFlow) {
W170 := __e.Get(1)
_ = W170
tmp17429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W170)


if True == tmp17429 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17415 := MakeNative(func(__e *ControlFlow) {
W171 := __e.Get(1)
_ = W171
tmp17416 := MakeNative(func(__e *ControlFlow) {
W172 := __e.Get(1)
_ = W172
tmp17417 := MakeNative(func(__e *ControlFlow) {
W173 := __e.Get(1)
_ = W173
tmp17424 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W173)


if True == tmp17424 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17418 := MakeNative(func(__e *ControlFlow) {
W174 := __e.Get(1)
_ = W174
tmp17419 := Call(__e, PrimFunc(symshen_4autocomplete), W171)


tmp17420 := PrimCons(tmp17419, Nil)

tmp17421 := PrimCons(W171, tmp17420)

__e.TailApply(PrimFunc(symshen_4comb), W174, tmp17421)
return


}, 1)

tmp17422 := Call(__e, PrimFunc(symshen_4in_1_6), W173)


__e.TailApply(tmp17418, tmp17422)
return


}


}, 1)

tmp17425 := Call(__e, PrimFunc(symshen_4_5sc_6), W172)


__e.TailApply(tmp17417, tmp17425)
return


}, 1)

tmp17426 := Call(__e, PrimFunc(symshen_4in_1_6), W170)


__e.TailApply(tmp17416, tmp17426)
return


}, 1)

tmp17427 := Call(__e, PrimFunc(symshen_4_5_1out), W170)


__e.TailApply(tmp17415, tmp17427)
return


}


}, 1)

tmp17430 := Call(__e, PrimFunc(symshen_4_5syntax_6), V159)


tmp17431 := Call(__e, tmp17414, tmp17430)


__e.TailApply(tmp17411, tmp17431)
return


} else {
__e.Return(W160)
return
}


}, 1)

tmp17434 := MakeNative(func(__e *ControlFlow) {
W161 := __e.Get(1)
_ = W161
tmp17456 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W161)


if True == tmp17456 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17435 := MakeNative(func(__e *ControlFlow) {
W162 := __e.Get(1)
_ = W162
tmp17436 := MakeNative(func(__e *ControlFlow) {
W163 := __e.Get(1)
_ = W163
tmp17437 := MakeNative(func(__e *ControlFlow) {
W164 := __e.Get(1)
_ = W164
tmp17451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W164)


if True == tmp17451 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17438 := MakeNative(func(__e *ControlFlow) {
W165 := __e.Get(1)
_ = W165
tmp17439 := MakeNative(func(__e *ControlFlow) {
W166 := __e.Get(1)
_ = W166
tmp17440 := MakeNative(func(__e *ControlFlow) {
W167 := __e.Get(1)
_ = W167
tmp17446 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W167)


if True == tmp17446 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17441 := MakeNative(func(__e *ControlFlow) {
W168 := __e.Get(1)
_ = W168
tmp17442 := PrimCons(W165, Nil)

tmp17443 := PrimCons(W162, tmp17442)

__e.TailApply(PrimFunc(symshen_4comb), W168, tmp17443)
return


}, 1)

tmp17444 := Call(__e, PrimFunc(symshen_4in_1_6), W167)


__e.TailApply(tmp17441, tmp17444)
return


}


}, 1)

tmp17447 := Call(__e, PrimFunc(symshen_4_5sc_6), W166)


__e.TailApply(tmp17440, tmp17447)
return


}, 1)

tmp17448 := Call(__e, PrimFunc(symshen_4in_1_6), W164)


__e.TailApply(tmp17439, tmp17448)
return


}, 1)

tmp17449 := Call(__e, PrimFunc(symshen_4_5_1out), W164)


__e.TailApply(tmp17438, tmp17449)
return


}


}, 1)

tmp17452 := Call(__e, PrimFunc(symshen_4_5semantics_6), W163)


__e.TailApply(tmp17437, tmp17452)
return


}, 1)

tmp17453 := Call(__e, PrimFunc(symshen_4in_1_6), W161)


__e.TailApply(tmp17436, tmp17453)
return


}, 1)

tmp17454 := Call(__e, PrimFunc(symshen_4_5_1out), W161)


__e.TailApply(tmp17435, tmp17454)
return


}


}, 1)

tmp17457 := Call(__e, PrimFunc(symshen_4_5syntax_6), V159)


tmp17458 := Call(__e, tmp17434, tmp17457)


__e.TailApply(tmp17410, tmp17458)
return


}, 1)

tmp17459 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17409)


_ = tmp17459

tmp17460 := MakeNative(func(__e *ControlFlow) {
V175 := __e.Get(1)
_ = V175
tmp17489 := PrimIsPair(V175)

var ifres17481 Obj

if True == tmp17489 {
tmp17487 := PrimTail(V175)

tmp17488 := PrimEqual(Nil, tmp17487)

var ifres17483 Obj

if True == tmp17488 {
tmp17485 := PrimHead(V175)

tmp17486 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17485)


var ifres17484 Obj

if True == tmp17486 {
ifres17484 = True


} else {
ifres17484 = False


}

ifres17483 = ifres17484


} else {
ifres17483 = False


}

var ifres17482 Obj

if True == ifres17483 {
ifres17482 = True


} else {
ifres17482 = False


}

ifres17481 = ifres17482


} else {
ifres17481 = False


}

if True == ifres17481 {
__e.Return(PrimHead(V175))
return
} else {
tmp17479 := PrimIsPair(V175)

var ifres17475 Obj

if True == tmp17479 {
tmp17477 := PrimHead(V175)

tmp17478 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17477)


var ifres17476 Obj

if True == tmp17478 {
ifres17476 = True


} else {
ifres17476 = False


}

ifres17475 = ifres17476


} else {
ifres17475 = False


}

if True == ifres17475 {
tmp17461 := PrimHead(V175)

tmp17462 := PrimTail(V175)

tmp17463 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17462)


tmp17464 := PrimCons(tmp17463, Nil)

tmp17465 := PrimCons(tmp17461, tmp17464)

__e.Return(PrimCons(symappend, tmp17465))
return


} else {
tmp17473 := PrimIsPair(V175)

if True == tmp17473 {
tmp17466 := PrimHead(V175)

tmp17467 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17466)


tmp17468 := PrimTail(V175)

tmp17469 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17468)


tmp17470 := PrimCons(tmp17469, Nil)

tmp17471 := PrimCons(tmp17467, tmp17470)

__e.Return(PrimCons(symcons, tmp17471))
return


} else {
__e.Return(V175)
return
}


}


}


}, 1)

tmp17490 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17460)


_ = tmp17490

tmp17491 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17498 := PrimIsSymbol(V176)

if True == tmp17498 {
tmp17493 := MakeNative(func(__e *ControlFlow) {
W177 := __e.Get(1)
_ = W177
tmp17494 := MakeNative(func(__e *ControlFlow) {
Z178 := __e.Get(1)
_ = Z178
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z178)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17494, W177)
return


}, 1)

tmp17495 := Call(__e, PrimFunc(symexplode), V176)


tmp17496 := Call(__e, tmp17493, tmp17495)


if True == tmp17496 {
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

tmp17499 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17491)


_ = tmp17499

tmp17500 := MakeNative(func(__e *ControlFlow) {
V179 := __e.Get(1)
_ = V179
tmp17501 := MakeNative(func(__e *ControlFlow) {
W180 := __e.Get(1)
_ = W180
tmp17523 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W180)


if True == tmp17523 {
tmp17502 := MakeNative(func(__e *ControlFlow) {
W185 := __e.Get(1)
_ = W185
tmp17514 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W185)


if True == tmp17514 {
tmp17503 := MakeNative(func(__e *ControlFlow) {
W188 := __e.Get(1)
_ = W188
tmp17505 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W188)


if True == tmp17505 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W188)
return
}


}, 1)

tmp17506 := MakeNative(func(__e *ControlFlow) {
W189 := __e.Get(1)
_ = W189
tmp17510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W189)


if True == tmp17510 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17507 := MakeNative(func(__e *ControlFlow) {
W190 := __e.Get(1)
_ = W190
__e.TailApply(PrimFunc(symshen_4comb), W190, False)
return
}, 1)

tmp17508 := Call(__e, PrimFunc(symshen_4in_1_6), W189)


__e.TailApply(tmp17507, tmp17508)
return


}


}, 1)

tmp17511 := Call(__e, PrimFunc(sym_5_b_6), V179)


tmp17512 := Call(__e, tmp17506, tmp17511)


__e.TailApply(tmp17503, tmp17512)
return


} else {
__e.Return(W185)
return
}


}, 1)

tmp17515 := MakeNative(func(__e *ControlFlow) {
W186 := __e.Get(1)
_ = W186
tmp17519 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W186)


if True == tmp17519 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17516 := MakeNative(func(__e *ControlFlow) {
W187 := __e.Get(1)
_ = W187
__e.TailApply(PrimFunc(symshen_4comb), W187, True)
return
}, 1)

tmp17517 := Call(__e, PrimFunc(symshen_4in_1_6), W186)


__e.TailApply(tmp17516, tmp17517)
return


}


}, 1)

tmp17520 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V179)


tmp17521 := Call(__e, tmp17515, tmp17520)


__e.TailApply(tmp17502, tmp17521)
return


} else {
__e.Return(W180)
return
}


}, 1)

tmp17524 := MakeNative(func(__e *ControlFlow) {
W181 := __e.Get(1)
_ = W181
tmp17534 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W181)


if True == tmp17534 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17525 := MakeNative(func(__e *ControlFlow) {
W182 := __e.Get(1)
_ = W182
tmp17526 := MakeNative(func(__e *ControlFlow) {
W183 := __e.Get(1)
_ = W183
tmp17530 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W183)


if True == tmp17530 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17527 := MakeNative(func(__e *ControlFlow) {
W184 := __e.Get(1)
_ = W184
__e.TailApply(PrimFunc(symshen_4comb), W184, True)
return
}, 1)

tmp17528 := Call(__e, PrimFunc(symshen_4in_1_6), W183)


__e.TailApply(tmp17527, tmp17528)
return


}


}, 1)

tmp17531 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), W182)


__e.TailApply(tmp17526, tmp17531)
return


}, 1)

tmp17532 := Call(__e, PrimFunc(symshen_4in_1_6), W181)


__e.TailApply(tmp17525, tmp17532)
return


}


}, 1)

tmp17535 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V179)


tmp17536 := Call(__e, tmp17524, tmp17535)


__e.TailApply(tmp17501, tmp17536)
return


}, 1)

tmp17537 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17500)


_ = tmp17537

tmp17538 := MakeNative(func(__e *ControlFlow) {
V191 := __e.Get(1)
_ = V191
tmp17539 := MakeNative(func(__e *ControlFlow) {
W192 := __e.Get(1)
_ = W192
tmp17555 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W192)


if True == tmp17555 {
tmp17540 := MakeNative(func(__e *ControlFlow) {
W198 := __e.Get(1)
_ = W198
tmp17542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W198)


if True == tmp17542 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W198)
return
}


}, 1)

tmp17543 := MakeNative(func(__e *ControlFlow) {
W199 := __e.Get(1)
_ = W199
tmp17551 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W199)


if True == tmp17551 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17544 := MakeNative(func(__e *ControlFlow) {
W200 := __e.Get(1)
_ = W200
tmp17548 := Call(__e, PrimFunc(symshen_4hds_a_2), W200, MakeString("."))


if True == tmp17548 {
tmp17545 := MakeNative(func(__e *ControlFlow) {
W201 := __e.Get(1)
_ = W201
__e.TailApply(PrimFunc(symshen_4comb), W201, symshen_4skip)
return
}, 1)

tmp17546 := Call(__e, PrimFunc(symtail), W200)


__e.TailApply(tmp17545, tmp17546)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17549 := Call(__e, PrimFunc(symshen_4in_1_6), W199)


__e.TailApply(tmp17544, tmp17549)
return


}


}, 1)

tmp17552 := Call(__e, PrimFunc(symshen_4_5packagename_6), V191)


tmp17553 := Call(__e, tmp17543, tmp17552)


__e.TailApply(tmp17540, tmp17553)
return


} else {
__e.Return(W192)
return
}


}, 1)

tmp17556 := MakeNative(func(__e *ControlFlow) {
W193 := __e.Get(1)
_ = W193
tmp17570 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W193)


if True == tmp17570 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17557 := MakeNative(func(__e *ControlFlow) {
W194 := __e.Get(1)
_ = W194
tmp17567 := Call(__e, PrimFunc(symshen_4hds_a_2), W194, MakeString("."))


if True == tmp17567 {
tmp17558 := MakeNative(func(__e *ControlFlow) {
W195 := __e.Get(1)
_ = W195
tmp17559 := MakeNative(func(__e *ControlFlow) {
W196 := __e.Get(1)
_ = W196
tmp17563 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W196)


if True == tmp17563 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17560 := MakeNative(func(__e *ControlFlow) {
W197 := __e.Get(1)
_ = W197
__e.TailApply(PrimFunc(symshen_4comb), W197, symshen_4skip)
return
}, 1)

tmp17561 := Call(__e, PrimFunc(symshen_4in_1_6), W196)


__e.TailApply(tmp17560, tmp17561)
return


}


}, 1)

tmp17564 := Call(__e, PrimFunc(symshen_4_5packagenames_6), W195)


__e.TailApply(tmp17559, tmp17564)
return


}, 1)

tmp17565 := Call(__e, PrimFunc(symtail), W194)


__e.TailApply(tmp17558, tmp17565)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17568 := Call(__e, PrimFunc(symshen_4in_1_6), W193)


__e.TailApply(tmp17557, tmp17568)
return


}


}, 1)

tmp17571 := Call(__e, PrimFunc(symshen_4_5packagename_6), V191)


tmp17572 := Call(__e, tmp17556, tmp17571)


__e.TailApply(tmp17539, tmp17572)
return


}, 1)

tmp17573 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17538)


_ = tmp17573

tmp17574 := MakeNative(func(__e *ControlFlow) {
V202 := __e.Get(1)
_ = V202
tmp17575 := MakeNative(func(__e *ControlFlow) {
W203 := __e.Get(1)
_ = W203
tmp17587 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W203)


if True == tmp17587 {
tmp17576 := MakeNative(func(__e *ControlFlow) {
W208 := __e.Get(1)
_ = W208
tmp17578 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W208)


if True == tmp17578 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W208)
return
}


}, 1)

tmp17579 := MakeNative(func(__e *ControlFlow) {
W209 := __e.Get(1)
_ = W209
tmp17583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W209)


if True == tmp17583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17580 := MakeNative(func(__e *ControlFlow) {
W210 := __e.Get(1)
_ = W210
__e.TailApply(PrimFunc(symshen_4comb), W210, symshen_4skip)
return
}, 1)

tmp17581 := Call(__e, PrimFunc(symshen_4in_1_6), W209)


__e.TailApply(tmp17580, tmp17581)
return


}


}, 1)

tmp17584 := Call(__e, PrimFunc(sym_5e_6), V202)


tmp17585 := Call(__e, tmp17579, tmp17584)


__e.TailApply(tmp17576, tmp17585)
return


} else {
__e.Return(W203)
return
}


}, 1)

tmp17588 := MakeNative(func(__e *ControlFlow) {
W204 := __e.Get(1)
_ = W204
tmp17598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W204)


if True == tmp17598 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17589 := MakeNative(func(__e *ControlFlow) {
W205 := __e.Get(1)
_ = W205
tmp17590 := MakeNative(func(__e *ControlFlow) {
W206 := __e.Get(1)
_ = W206
tmp17594 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W206)


if True == tmp17594 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17591 := MakeNative(func(__e *ControlFlow) {
W207 := __e.Get(1)
_ = W207
__e.TailApply(PrimFunc(symshen_4comb), W207, symshen_4skip)
return
}, 1)

tmp17592 := Call(__e, PrimFunc(symshen_4in_1_6), W206)


__e.TailApply(tmp17591, tmp17592)
return


}


}, 1)

tmp17595 := Call(__e, PrimFunc(symshen_4_5packagename_6), W205)


__e.TailApply(tmp17590, tmp17595)
return


}, 1)

tmp17596 := Call(__e, PrimFunc(symshen_4in_1_6), W204)


__e.TailApply(tmp17589, tmp17596)
return


}


}, 1)

tmp17599 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V202)


tmp17600 := Call(__e, tmp17588, tmp17599)


__e.TailApply(tmp17575, tmp17600)
return


}, 1)

tmp17601 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17574)


_ = tmp17601

tmp17602 := MakeNative(func(__e *ControlFlow) {
V211 := __e.Get(1)
_ = V211
tmp17603 := MakeNative(func(__e *ControlFlow) {
W212 := __e.Get(1)
_ = W212
tmp17605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W212)


if True == tmp17605 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W212)
return
}


}, 1)

tmp17616 := PrimIsPair(V211)

var ifres17606 Obj

if True == tmp17616 {
tmp17607 := MakeNative(func(__e *ControlFlow) {
W213 := __e.Get(1)
_ = W213
tmp17608 := MakeNative(func(__e *ControlFlow) {
W214 := __e.Get(1)
_ = W214
tmp17610 := PrimEqual(W213, MakeString("."))

tmp17611 := PrimNot(tmp17610)

if True == tmp17611 {
__e.TailApply(PrimFunc(symshen_4comb), W214, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17612 := Call(__e, PrimFunc(symtail), V211)


__e.TailApply(tmp17608, tmp17612)
return


}, 1)

tmp17613 := Call(__e, PrimFunc(symhead), V211)


tmp17614 := Call(__e, tmp17607, tmp17613)


ifres17606 = tmp17614


} else {
tmp17615 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17606 = tmp17615


}

__e.TailApply(tmp17603, ifres17606)
return


}, 1)

tmp17617 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17602)


_ = tmp17617

tmp17618 := MakeNative(func(__e *ControlFlow) {
V215 := __e.Get(1)
_ = V215
tmp17619 := MakeNative(func(__e *ControlFlow) {
W216 := __e.Get(1)
_ = W216
tmp17621 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W216)


if True == tmp17621 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W216)
return
}


}, 1)

tmp17644 := Call(__e, PrimFunc(symshen_4hds_a_2), V215, MakeString("<"))


var ifres17622 Obj

if True == tmp17644 {
tmp17623 := MakeNative(func(__e *ControlFlow) {
W217 := __e.Get(1)
_ = W217
tmp17624 := MakeNative(func(__e *ControlFlow) {
W218 := __e.Get(1)
_ = W218
tmp17639 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W218)


if True == tmp17639 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17625 := MakeNative(func(__e *ControlFlow) {
W219 := __e.Get(1)
_ = W219
tmp17626 := MakeNative(func(__e *ControlFlow) {
W220 := __e.Get(1)
_ = W220
tmp17628 := MakeNative(func(__e *ControlFlow) {
W221 := __e.Get(1)
_ = W221
tmp17633 := PrimIsPair(W221)

if True == tmp17633 {
tmp17630 := PrimHead(W221)

tmp17631 := PrimEqual(tmp17630, MakeString(">"))

if True == tmp17631 {
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

tmp17634 := Call(__e, PrimFunc(symreverse), W219)


tmp17635 := Call(__e, tmp17628, tmp17634)


if True == tmp17635 {
__e.TailApply(PrimFunc(symshen_4comb), W220, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17636 := Call(__e, PrimFunc(symshen_4in_1_6), W218)


__e.TailApply(tmp17626, tmp17636)
return


}, 1)

tmp17637 := Call(__e, PrimFunc(symshen_4_5_1out), W218)


__e.TailApply(tmp17625, tmp17637)
return


}


}, 1)

tmp17640 := Call(__e, PrimFunc(sym_5_b_6), W217)


__e.TailApply(tmp17624, tmp17640)
return


}, 1)

tmp17641 := Call(__e, PrimFunc(symtail), V215)


tmp17642 := Call(__e, tmp17623, tmp17641)


ifres17622 = tmp17642


} else {
tmp17643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17622 = tmp17643


}

__e.TailApply(tmp17619, ifres17622)
return


}, 1)

tmp17645 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17618)


_ = tmp17645

tmp17646 := MakeNative(func(__e *ControlFlow) {
V222 := __e.Get(1)
_ = V222
tmp17647 := PrimIntern(MakeString(";"))

__e.Return(PrimEqual(V222, tmp17647))
return


}, 1)

tmp17648 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17646)


_ = tmp17648

tmp17649 := MakeNative(func(__e *ControlFlow) {
V223 := __e.Get(1)
_ = V223
tmp17650 := MakeNative(func(__e *ControlFlow) {
W224 := __e.Get(1)
_ = W224
tmp17652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W224)


if True == tmp17652 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W224)
return
}


}, 1)

tmp17662 := PrimIsPair(V223)

var ifres17653 Obj

if True == tmp17662 {
tmp17654 := MakeNative(func(__e *ControlFlow) {
W225 := __e.Get(1)
_ = W225
tmp17655 := MakeNative(func(__e *ControlFlow) {
W226 := __e.Get(1)
_ = W226
tmp17657 := Call(__e, PrimFunc(symshen_4colon_1equal_2), W225)


if True == tmp17657 {
__e.TailApply(PrimFunc(symshen_4comb), W226, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17658 := Call(__e, PrimFunc(symtail), V223)


__e.TailApply(tmp17655, tmp17658)
return


}, 1)

tmp17659 := Call(__e, PrimFunc(symhead), V223)


tmp17660 := Call(__e, tmp17654, tmp17659)


ifres17653 = tmp17660


} else {
tmp17661 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17653 = tmp17661


}

__e.TailApply(tmp17650, ifres17653)
return


}, 1)

tmp17663 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17649)


_ = tmp17663

tmp17664 := MakeNative(func(__e *ControlFlow) {
V227 := __e.Get(1)
_ = V227
tmp17665 := PrimIntern(MakeString(":="))

__e.Return(PrimEqual(tmp17665, V227))
return


}, 1)

tmp17666 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17664)


_ = tmp17666

tmp17667 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17668 := MakeNative(func(__e *ControlFlow) {
W229 := __e.Get(1)
_ = W229
tmp17683 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W229)


if True == tmp17683 {
tmp17669 := MakeNative(func(__e *ControlFlow) {
W236 := __e.Get(1)
_ = W236
tmp17671 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W236)


if True == tmp17671 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W236)
return
}


}, 1)

tmp17672 := MakeNative(func(__e *ControlFlow) {
W237 := __e.Get(1)
_ = W237
tmp17679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W237)


if True == tmp17679 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17673 := MakeNative(func(__e *ControlFlow) {
W238 := __e.Get(1)
_ = W238
tmp17674 := MakeNative(func(__e *ControlFlow) {
W239 := __e.Get(1)
_ = W239
tmp17675 := PrimCons(W238, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W239, tmp17675)
return


}, 1)

tmp17676 := Call(__e, PrimFunc(symshen_4in_1_6), W237)


__e.TailApply(tmp17674, tmp17676)
return


}, 1)

tmp17677 := Call(__e, PrimFunc(symshen_4_5_1out), W237)


__e.TailApply(tmp17673, tmp17677)
return


}


}, 1)

tmp17680 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V228)


tmp17681 := Call(__e, tmp17672, tmp17680)


__e.TailApply(tmp17669, tmp17681)
return


} else {
__e.Return(W229)
return
}


}, 1)

tmp17684 := MakeNative(func(__e *ControlFlow) {
W230 := __e.Get(1)
_ = W230
tmp17699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W230)


if True == tmp17699 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17685 := MakeNative(func(__e *ControlFlow) {
W231 := __e.Get(1)
_ = W231
tmp17686 := MakeNative(func(__e *ControlFlow) {
W232 := __e.Get(1)
_ = W232
tmp17687 := MakeNative(func(__e *ControlFlow) {
W233 := __e.Get(1)
_ = W233
tmp17694 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W233)


if True == tmp17694 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17688 := MakeNative(func(__e *ControlFlow) {
W234 := __e.Get(1)
_ = W234
tmp17689 := MakeNative(func(__e *ControlFlow) {
W235 := __e.Get(1)
_ = W235
tmp17690 := PrimCons(W231, W234)

__e.TailApply(PrimFunc(symshen_4comb), W235, tmp17690)
return


}, 1)

tmp17691 := Call(__e, PrimFunc(symshen_4in_1_6), W233)


__e.TailApply(tmp17689, tmp17691)
return


}, 1)

tmp17692 := Call(__e, PrimFunc(symshen_4_5_1out), W233)


__e.TailApply(tmp17688, tmp17692)
return


}


}, 1)

tmp17695 := Call(__e, PrimFunc(symshen_4_5syntax_6), W232)


__e.TailApply(tmp17687, tmp17695)
return


}, 1)

tmp17696 := Call(__e, PrimFunc(symshen_4in_1_6), W230)


__e.TailApply(tmp17686, tmp17696)
return


}, 1)

tmp17697 := Call(__e, PrimFunc(symshen_4_5_1out), W230)


__e.TailApply(tmp17685, tmp17697)
return


}


}, 1)

tmp17700 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V228)


tmp17701 := Call(__e, tmp17684, tmp17700)


__e.TailApply(tmp17668, tmp17701)
return


}, 1)

tmp17702 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17667)


_ = tmp17702

tmp17703 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17704 := MakeNative(func(__e *ControlFlow) {
W241 := __e.Get(1)
_ = W241
tmp17706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W241)


if True == tmp17706 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W241)
return
}


}, 1)

tmp17716 := PrimIsPair(V240)

var ifres17707 Obj

if True == tmp17716 {
tmp17708 := MakeNative(func(__e *ControlFlow) {
W242 := __e.Get(1)
_ = W242
tmp17709 := MakeNative(func(__e *ControlFlow) {
W243 := __e.Get(1)
_ = W243
tmp17711 := Call(__e, PrimFunc(symshen_4syntax_1item_2), W242)


if True == tmp17711 {
__e.TailApply(PrimFunc(symshen_4comb), W243, W242)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17712 := Call(__e, PrimFunc(symtail), V240)


__e.TailApply(tmp17709, tmp17712)
return


}, 1)

tmp17713 := Call(__e, PrimFunc(symhead), V240)


tmp17714 := Call(__e, tmp17708, tmp17713)


ifres17707 = tmp17714


} else {
tmp17715 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17707 = tmp17715


}

__e.TailApply(tmp17704, ifres17707)
return


}, 1)

tmp17717 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17703)


_ = tmp17717

tmp17718 := MakeNative(func(__e *ControlFlow) {
V246 := __e.Get(1)
_ = V246
tmp17754 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V246)


if True == tmp17754 {
__e.Return(False)
return
} else {
tmp17752 := Call(__e, PrimFunc(symshen_4semicolon_2), V246)


if True == tmp17752 {
__e.Return(False)
return
} else {
tmp17750 := Call(__e, PrimFunc(symatom_2), V246)


if True == tmp17750 {
__e.Return(True)
return
} else {
tmp17748 := PrimIsPair(V246)

var ifres17729 Obj

if True == tmp17748 {
tmp17746 := PrimHead(V246)

tmp17747 := PrimEqual(symcons, tmp17746)

var ifres17731 Obj

if True == tmp17747 {
tmp17744 := PrimTail(V246)

tmp17745 := PrimIsPair(tmp17744)

var ifres17733 Obj

if True == tmp17745 {
tmp17741 := PrimTail(V246)

tmp17742 := PrimTail(tmp17741)

tmp17743 := PrimIsPair(tmp17742)

var ifres17735 Obj

if True == tmp17743 {
tmp17737 := PrimTail(V246)

tmp17738 := PrimTail(tmp17737)

tmp17739 := PrimTail(tmp17738)

tmp17740 := PrimEqual(Nil, tmp17739)

var ifres17736 Obj

if True == tmp17740 {
ifres17736 = True


} else {
ifres17736 = False


}

ifres17735 = ifres17736


} else {
ifres17735 = False


}

var ifres17734 Obj

if True == ifres17735 {
ifres17734 = True


} else {
ifres17734 = False


}

ifres17733 = ifres17734


} else {
ifres17733 = False


}

var ifres17732 Obj

if True == ifres17733 {
ifres17732 = True


} else {
ifres17732 = False


}

ifres17731 = ifres17732


} else {
ifres17731 = False


}

var ifres17730 Obj

if True == ifres17731 {
ifres17730 = True


} else {
ifres17730 = False


}

ifres17729 = ifres17730


} else {
ifres17729 = False


}

if True == ifres17729 {
tmp17725 := PrimTail(V246)

tmp17726 := PrimHead(tmp17725)

tmp17727 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17726)


if True == tmp17727 {
tmp17720 := PrimTail(V246)

tmp17721 := PrimTail(tmp17720)

tmp17722 := PrimHead(tmp17721)

tmp17723 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17722)


if True == tmp17723 {
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

tmp17755 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17718)


_ = tmp17755

tmp17756 := MakeNative(func(__e *ControlFlow) {
V247 := __e.Get(1)
_ = V247
tmp17757 := MakeNative(func(__e *ControlFlow) {
W248 := __e.Get(1)
_ = W248
tmp17778 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W248)


if True == tmp17778 {
tmp17758 := MakeNative(func(__e *ControlFlow) {
W256 := __e.Get(1)
_ = W256
tmp17760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W256)


if True == tmp17760 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W256)
return
}


}, 1)

tmp17761 := MakeNative(func(__e *ControlFlow) {
W257 := __e.Get(1)
_ = W257
tmp17774 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W257)


if True == tmp17774 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17762 := MakeNative(func(__e *ControlFlow) {
W258 := __e.Get(1)
_ = W258
tmp17771 := PrimIsPair(W258)

if True == tmp17771 {
tmp17763 := MakeNative(func(__e *ControlFlow) {
W259 := __e.Get(1)
_ = W259
tmp17764 := MakeNative(func(__e *ControlFlow) {
W260 := __e.Get(1)
_ = W260
tmp17766 := Call(__e, PrimFunc(symshen_4semicolon_2), W259)


tmp17767 := PrimNot(tmp17766)

if True == tmp17767 {
__e.TailApply(PrimFunc(symshen_4comb), W260, W259)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17768 := Call(__e, PrimFunc(symtail), W258)


__e.TailApply(tmp17764, tmp17768)
return


}, 1)

tmp17769 := Call(__e, PrimFunc(symhead), W258)


__e.TailApply(tmp17763, tmp17769)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17772 := Call(__e, PrimFunc(symshen_4in_1_6), W257)


__e.TailApply(tmp17762, tmp17772)
return


}


}, 1)

tmp17775 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V247)


tmp17776 := Call(__e, tmp17761, tmp17775)


__e.TailApply(tmp17758, tmp17776)
return


} else {
__e.Return(W248)
return
}


}, 1)

tmp17779 := MakeNative(func(__e *ControlFlow) {
W249 := __e.Get(1)
_ = W249
tmp17805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W249)


if True == tmp17805 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17780 := MakeNative(func(__e *ControlFlow) {
W250 := __e.Get(1)
_ = W250
tmp17802 := PrimIsPair(W250)

if True == tmp17802 {
tmp17781 := MakeNative(func(__e *ControlFlow) {
W251 := __e.Get(1)
_ = W251
tmp17782 := MakeNative(func(__e *ControlFlow) {
W252 := __e.Get(1)
_ = W252
tmp17798 := Call(__e, PrimFunc(symshen_4hds_a_2), W252, symwhere)


if True == tmp17798 {
tmp17783 := MakeNative(func(__e *ControlFlow) {
W253 := __e.Get(1)
_ = W253
tmp17795 := PrimIsPair(W253)

if True == tmp17795 {
tmp17784 := MakeNative(func(__e *ControlFlow) {
W254 := __e.Get(1)
_ = W254
tmp17785 := MakeNative(func(__e *ControlFlow) {
W255 := __e.Get(1)
_ = W255
tmp17790 := Call(__e, PrimFunc(symshen_4semicolon_2), W251)


tmp17791 := PrimNot(tmp17790)

if True == tmp17791 {
tmp17786 := PrimCons(W251, Nil)

tmp17787 := PrimCons(W254, tmp17786)

tmp17788 := PrimCons(symwhere, tmp17787)

__e.TailApply(PrimFunc(symshen_4comb), W255, tmp17788)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17792 := Call(__e, PrimFunc(symtail), W253)


__e.TailApply(tmp17785, tmp17792)
return


}, 1)

tmp17793 := Call(__e, PrimFunc(symhead), W253)


__e.TailApply(tmp17784, tmp17793)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17796 := Call(__e, PrimFunc(symtail), W252)


__e.TailApply(tmp17783, tmp17796)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17799 := Call(__e, PrimFunc(symtail), W250)


__e.TailApply(tmp17782, tmp17799)
return


}, 1)

tmp17800 := Call(__e, PrimFunc(symhead), W250)


__e.TailApply(tmp17781, tmp17800)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17803 := Call(__e, PrimFunc(symshen_4in_1_6), W249)


__e.TailApply(tmp17780, tmp17803)
return


}


}, 1)

tmp17806 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V247)


tmp17807 := Call(__e, tmp17779, tmp17806)


__e.TailApply(tmp17757, tmp17807)
return


}, 1)

tmp17808 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp17756)


_ = tmp17808

tmp17809 := MakeNative(func(__e *ControlFlow) {
V269 := __e.Get(1)
_ = V269
V270 := __e.Get(2)
_ = V270
V271 := __e.Get(3)
_ = V271
tmp17817 := PrimEqual(Nil, V271)

if True == tmp17817 {
__e.Return(PrimCons(symshen_4parse_1failure, Nil))
return
} else {
tmp17815 := PrimIsPair(V271)

if True == tmp17815 {
tmp17810 := PrimHead(V271)

tmp17811 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V269, tmp17810, V270)


tmp17812 := PrimTail(V271)

tmp17813 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V269, V270, tmp17812)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp17811, tmp17813)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
return
}


}


}, 3)

tmp17818 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp17809)


_ = tmp17818

tmp17819 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp17820 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp17819)


_ = tmp17820

tmp17821 := MakeNative(func(__e *ControlFlow) {
V272 := __e.Get(1)
_ = V272
V273 := __e.Get(2)
_ = V273
tmp17822 := PrimCons(symResult, Nil)

tmp17823 := PrimCons(symshen_4parse_1failure_2, tmp17822)

tmp17824 := PrimCons(symResult, Nil)

tmp17825 := PrimCons(V273, tmp17824)

tmp17826 := PrimCons(tmp17823, tmp17825)

tmp17827 := PrimCons(symif, tmp17826)

tmp17828 := PrimCons(tmp17827, Nil)

tmp17829 := PrimCons(V272, tmp17828)

tmp17830 := PrimCons(symResult, tmp17829)

__e.Return(PrimCons(symlet, tmp17830))
return


}, 2)

tmp17831 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp17821)


_ = tmp17831

tmp17832 := MakeNative(func(__e *ControlFlow) {
V280 := __e.Get(1)
_ = V280
V281 := __e.Get(2)
_ = V281
V282 := __e.Get(3)
_ = V282
tmp17846 := PrimIsPair(V281)

var ifres17837 Obj

if True == tmp17846 {
tmp17844 := PrimTail(V281)

tmp17845 := PrimIsPair(tmp17844)

var ifres17839 Obj

if True == tmp17845 {
tmp17841 := PrimTail(V281)

tmp17842 := PrimTail(tmp17841)

tmp17843 := PrimEqual(Nil, tmp17842)

var ifres17840 Obj

if True == tmp17843 {
ifres17840 = True


} else {
ifres17840 = False


}

ifres17839 = ifres17840


} else {
ifres17839 = False


}

var ifres17838 Obj

if True == ifres17839 {
ifres17838 = True


} else {
ifres17838 = False


}

ifres17837 = ifres17838


} else {
ifres17837 = False


}

if True == ifres17837 {
tmp17833 := PrimHead(V281)

tmp17834 := PrimTail(V281)

tmp17835 := PrimHead(tmp17834)

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V280, V282, tmp17833, tmp17835)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
return
}


}, 3)

tmp17847 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp17832)


_ = tmp17847

tmp17848 := MakeNative(func(__e *ControlFlow) {
V291 := __e.Get(1)
_ = V291
V292 := __e.Get(2)
_ = V292
V293 := __e.Get(3)
_ = V293
V294 := __e.Get(4)
_ = V294
tmp17912 := PrimEqual(Nil, V293)

var ifres17890 Obj

if True == tmp17912 {
tmp17911 := PrimIsPair(V294)

var ifres17892 Obj

if True == tmp17911 {
tmp17909 := PrimHead(V294)

tmp17910 := PrimEqual(symwhere, tmp17909)

var ifres17894 Obj

if True == tmp17910 {
tmp17907 := PrimTail(V294)

tmp17908 := PrimIsPair(tmp17907)

var ifres17896 Obj

if True == tmp17908 {
tmp17904 := PrimTail(V294)

tmp17905 := PrimTail(tmp17904)

tmp17906 := PrimIsPair(tmp17905)

var ifres17898 Obj

if True == tmp17906 {
tmp17900 := PrimTail(V294)

tmp17901 := PrimTail(tmp17900)

tmp17902 := PrimTail(tmp17901)

tmp17903 := PrimEqual(Nil, tmp17902)

var ifres17899 Obj

if True == tmp17903 {
ifres17899 = True


} else {
ifres17899 = False


}

ifres17898 = ifres17899


} else {
ifres17898 = False


}

var ifres17897 Obj

if True == ifres17898 {
ifres17897 = True


} else {
ifres17897 = False


}

ifres17896 = ifres17897


} else {
ifres17896 = False


}

var ifres17895 Obj

if True == ifres17896 {
ifres17895 = True


} else {
ifres17895 = False


}

ifres17894 = ifres17895


} else {
ifres17894 = False


}

var ifres17893 Obj

if True == ifres17894 {
ifres17893 = True


} else {
ifres17893 = False


}

ifres17892 = ifres17893


} else {
ifres17892 = False


}

var ifres17891 Obj

if True == ifres17892 {
ifres17891 = True


} else {
ifres17891 = False


}

ifres17890 = ifres17891


} else {
ifres17890 = False


}

if True == ifres17890 {
tmp17849 := PrimTail(V294)

tmp17850 := PrimHead(tmp17849)

tmp17851 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp17850)


tmp17852 := PrimTail(V294)

tmp17853 := PrimTail(tmp17852)

tmp17854 := PrimHead(tmp17853)

tmp17855 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V291, V292, Nil, tmp17854)


tmp17856 := PrimCons(symshen_4parse_1failure, Nil)

tmp17857 := PrimCons(tmp17856, Nil)

tmp17858 := PrimCons(tmp17855, tmp17857)

tmp17859 := PrimCons(tmp17851, tmp17858)

__e.Return(PrimCons(symif, tmp17859))
return


} else {
tmp17888 := PrimEqual(Nil, V293)

if True == tmp17888 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V291, V292, V294)
return
} else {
tmp17886 := PrimIsPair(V293)

if True == tmp17886 {
tmp17883 := PrimHead(V293)

tmp17884 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17883)


if True == tmp17884 {
tmp17860 := PrimHead(V293)

tmp17861 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V291, V292, tmp17860, tmp17861, V294)
return


} else {
tmp17880 := PrimHead(V293)

tmp17881 := PrimIsVariable(tmp17880)

if True == tmp17881 {
tmp17862 := PrimHead(V293)

tmp17863 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4variablecode), V291, V292, tmp17862, tmp17863, V294)
return


} else {
tmp17877 := PrimHead(V293)

tmp17878 := PrimEqual(sym__, tmp17877)

if True == tmp17878 {
tmp17864 := PrimHead(V293)

tmp17865 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4wildcardcode), V291, V292, tmp17864, tmp17865, V294)
return


} else {
tmp17874 := PrimHead(V293)

tmp17875 := Call(__e, PrimFunc(symatom_2), tmp17874)


if True == tmp17875 {
tmp17866 := PrimHead(V293)

tmp17867 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4terminalcode), V291, V292, tmp17866, tmp17867, V294)
return


} else {
tmp17871 := PrimHead(V293)

tmp17872 := PrimIsPair(tmp17871)

if True == tmp17872 {
tmp17868 := PrimHead(V293)

tmp17869 := PrimTail(V293)

__e.TailApply(PrimFunc(symshen_4conscode), V291, V292, tmp17868, tmp17869, V294)
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

tmp17913 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp17848)


_ = tmp17913

tmp17914 := MakeNative(func(__e *ControlFlow) {
V295 := __e.Get(1)
_ = V295
V296 := __e.Get(2)
_ = V296
V297 := __e.Get(3)
_ = V297
V298 := __e.Get(4)
_ = V298
V299 := __e.Get(5)
_ = V299
tmp17915 := MakeNative(func(__e *ControlFlow) {
W300 := __e.Get(1)
_ = W300
tmp17916 := MakeNative(func(__e *ControlFlow) {
W301 := __e.Get(1)
_ = W301
tmp17917 := MakeNative(func(__e *ControlFlow) {
W302 := __e.Get(1)
_ = W302
tmp17918 := PrimCons(V296, Nil)

tmp17919 := PrimCons(V297, tmp17918)

tmp17920 := PrimCons(W300, Nil)

tmp17921 := PrimCons(symshen_4parse_1failure_2, tmp17920)

tmp17922 := PrimCons(symshen_4parse_1failure, Nil)

tmp17923 := MakeNative(func(__e *ControlFlow) {
W303 := __e.Get(1)
_ = W303
tmp17933 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V297, V299)


var ifres17930 Obj

if True == tmp17933 {
ifres17930 = True


} else {
tmp17932 := Call(__e, PrimFunc(symshen_4occurs_1check_2), W301, V299)


var ifres17931 Obj

if True == tmp17932 {
ifres17931 = True


} else {
ifres17931 = False


}

ifres17930 = ifres17931


}

if True == ifres17930 {
tmp17924 := PrimCons(W300, Nil)

tmp17925 := PrimCons(symshen_4_5_1out, tmp17924)

tmp17926 := PrimCons(W303, Nil)

tmp17927 := PrimCons(tmp17925, tmp17926)

tmp17928 := PrimCons(W301, tmp17927)

__e.Return(PrimCons(symlet, tmp17928))
return


} else {
__e.Return(W303)
return
}


}, 1)

tmp17934 := PrimCons(W300, Nil)

tmp17935 := PrimCons(symshen_4in_1_6, tmp17934)

tmp17936 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V295, W302, V298, V299)


tmp17937 := PrimCons(tmp17936, Nil)

tmp17938 := PrimCons(tmp17935, tmp17937)

tmp17939 := PrimCons(W302, tmp17938)

tmp17940 := PrimCons(symlet, tmp17939)

tmp17941 := Call(__e, tmp17923, tmp17940)


tmp17942 := PrimCons(tmp17941, Nil)

tmp17943 := PrimCons(tmp17922, tmp17942)

tmp17944 := PrimCons(tmp17921, tmp17943)

tmp17945 := PrimCons(symif, tmp17944)

tmp17946 := PrimCons(tmp17945, Nil)

tmp17947 := PrimCons(tmp17919, tmp17946)

tmp17948 := PrimCons(W300, tmp17947)

__e.Return(PrimCons(symlet, tmp17948))
return


}, 1)

tmp17949 := Call(__e, PrimFunc(symconcat), symRemainder, V297)


__e.TailApply(tmp17917, tmp17949)
return


}, 1)

tmp17950 := Call(__e, PrimFunc(symconcat), symAction, V297)


__e.TailApply(tmp17916, tmp17950)
return


}, 1)

tmp17951 := Call(__e, PrimFunc(symconcat), symParse, V297)


__e.TailApply(tmp17915, tmp17951)
return


}, 5)

tmp17952 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp17914)


_ = tmp17952

tmp17953 := MakeNative(func(__e *ControlFlow) {
V304 := __e.Get(1)
_ = V304
V305 := __e.Get(2)
_ = V305
V306 := __e.Get(3)
_ = V306
V307 := __e.Get(4)
_ = V307
V308 := __e.Get(5)
_ = V308
tmp17954 := MakeNative(func(__e *ControlFlow) {
W309 := __e.Get(1)
_ = W309
tmp17955 := PrimCons(V305, Nil)

tmp17956 := PrimCons(symcons_2, tmp17955)

tmp17957 := MakeNative(func(__e *ControlFlow) {
W310 := __e.Get(1)
_ = W310
tmp17964 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V306, V308)


if True == tmp17964 {
tmp17958 := PrimCons(V305, Nil)

tmp17959 := PrimCons(symhead, tmp17958)

tmp17960 := PrimCons(W310, Nil)

tmp17961 := PrimCons(tmp17959, tmp17960)

tmp17962 := PrimCons(V306, tmp17961)

__e.Return(PrimCons(symlet, tmp17962))
return


} else {
__e.Return(W310)
return
}


}, 1)

tmp17965 := PrimCons(V305, Nil)

tmp17966 := PrimCons(symtail, tmp17965)

tmp17967 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V304, W309, V307, V308)


tmp17968 := PrimCons(tmp17967, Nil)

tmp17969 := PrimCons(tmp17966, tmp17968)

tmp17970 := PrimCons(W309, tmp17969)

tmp17971 := PrimCons(symlet, tmp17970)

tmp17972 := Call(__e, tmp17957, tmp17971)


tmp17973 := PrimCons(symshen_4parse_1failure, Nil)

tmp17974 := PrimCons(tmp17973, Nil)

tmp17975 := PrimCons(tmp17972, tmp17974)

tmp17976 := PrimCons(tmp17956, tmp17975)

__e.Return(PrimCons(symif, tmp17976))
return


}, 1)

tmp17977 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp17954, tmp17977)
return


}, 5)

tmp17978 := Call(__e, ns2_1set, symshen_4variablecode, tmp17953)


_ = tmp17978

tmp17979 := MakeNative(func(__e *ControlFlow) {
V311 := __e.Get(1)
_ = V311
V312 := __e.Get(2)
_ = V312
V313 := __e.Get(3)
_ = V313
V314 := __e.Get(4)
_ = V314
V315 := __e.Get(5)
_ = V315
tmp17980 := MakeNative(func(__e *ControlFlow) {
W316 := __e.Get(1)
_ = W316
tmp17981 := PrimCons(V312, Nil)

tmp17982 := PrimCons(symcons_2, tmp17981)

tmp17983 := PrimCons(V312, Nil)

tmp17984 := PrimCons(symtail, tmp17983)

tmp17985 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V311, W316, V314, V315)


tmp17986 := PrimCons(tmp17985, Nil)

tmp17987 := PrimCons(tmp17984, tmp17986)

tmp17988 := PrimCons(W316, tmp17987)

tmp17989 := PrimCons(symlet, tmp17988)

tmp17990 := PrimCons(symshen_4parse_1failure, Nil)

tmp17991 := PrimCons(tmp17990, Nil)

tmp17992 := PrimCons(tmp17989, tmp17991)

tmp17993 := PrimCons(tmp17982, tmp17992)

__e.Return(PrimCons(symif, tmp17993))
return


}, 1)

tmp17994 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp17980, tmp17994)
return


}, 5)

tmp17995 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp17979)


_ = tmp17995

tmp17996 := MakeNative(func(__e *ControlFlow) {
V317 := __e.Get(1)
_ = V317
V318 := __e.Get(2)
_ = V318
V319 := __e.Get(3)
_ = V319
V320 := __e.Get(4)
_ = V320
V321 := __e.Get(5)
_ = V321
tmp17997 := MakeNative(func(__e *ControlFlow) {
W322 := __e.Get(1)
_ = W322
tmp17998 := PrimCons(V319, Nil)

tmp17999 := PrimCons(V318, tmp17998)

tmp18000 := PrimCons(symshen_4hds_a_2, tmp17999)

tmp18001 := PrimCons(V318, Nil)

tmp18002 := PrimCons(symtail, tmp18001)

tmp18003 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V317, W322, V320, V321)


tmp18004 := PrimCons(tmp18003, Nil)

tmp18005 := PrimCons(tmp18002, tmp18004)

tmp18006 := PrimCons(W322, tmp18005)

tmp18007 := PrimCons(symlet, tmp18006)

tmp18008 := PrimCons(symshen_4parse_1failure, Nil)

tmp18009 := PrimCons(tmp18008, Nil)

tmp18010 := PrimCons(tmp18007, tmp18009)

tmp18011 := PrimCons(tmp18000, tmp18010)

__e.Return(PrimCons(symif, tmp18011))
return


}, 1)

tmp18012 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp17997, tmp18012)
return


}, 5)

tmp18013 := Call(__e, ns2_1set, symshen_4terminalcode, tmp17996)


_ = tmp18013

tmp18014 := MakeNative(func(__e *ControlFlow) {
V330 := __e.Get(1)
_ = V330
V331 := __e.Get(2)
_ = V331
tmp18020 := PrimIsPair(V330)

var ifres18016 Obj

if True == tmp18020 {
tmp18018 := PrimHead(V330)

tmp18019 := PrimEqual(tmp18018, V331)

var ifres18017 Obj

if True == tmp18019 {
ifres18017 = True


} else {
ifres18017 = False


}

ifres18016 = ifres18017


} else {
ifres18016 = False


}

if True == ifres18016 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18021 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18014)


_ = tmp18021

tmp18022 := MakeNative(func(__e *ControlFlow) {
V332 := __e.Get(1)
_ = V332
V333 := __e.Get(2)
_ = V333
V334 := __e.Get(3)
_ = V334
V335 := __e.Get(4)
_ = V335
V336 := __e.Get(5)
_ = V336
tmp18023 := MakeNative(func(__e *ControlFlow) {
W337 := __e.Get(1)
_ = W337
tmp18024 := MakeNative(func(__e *ControlFlow) {
W338 := __e.Get(1)
_ = W338
tmp18025 := MakeNative(func(__e *ControlFlow) {
W339 := __e.Get(1)
_ = W339
tmp18026 := PrimCons(V333, Nil)

tmp18027 := PrimCons(symshen_4ccons_2, tmp18026)

tmp18028 := PrimCons(V333, Nil)

tmp18029 := PrimCons(symhead, tmp18028)

tmp18030 := PrimCons(V333, Nil)

tmp18031 := PrimCons(symtail, tmp18030)

tmp18032 := Call(__e, PrimFunc(symshen_4decons), V334)


tmp18033 := PrimCons(sym_5end_6, Nil)

tmp18034 := Call(__e, PrimFunc(symappend), tmp18032, tmp18033)


tmp18035 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V332, W339, V335, V336)


tmp18036 := PrimCons(tmp18035, Nil)

tmp18037 := PrimCons(symshen_4processed, tmp18036)

tmp18038 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V332, W338, tmp18034, tmp18037)


tmp18039 := PrimCons(tmp18038, Nil)

tmp18040 := PrimCons(tmp18031, tmp18039)

tmp18041 := PrimCons(W339, tmp18040)

tmp18042 := PrimCons(tmp18029, tmp18041)

tmp18043 := PrimCons(W338, tmp18042)

tmp18044 := PrimCons(symlet, tmp18043)

tmp18045 := PrimCons(symshen_4parse_1failure, Nil)

tmp18046 := PrimCons(tmp18045, Nil)

tmp18047 := PrimCons(tmp18044, tmp18046)

tmp18048 := PrimCons(tmp18027, tmp18047)

__e.Return(PrimCons(symif, tmp18048))
return


}, 1)

tmp18049 := Call(__e, PrimFunc(symgensym), symTl)


__e.TailApply(tmp18025, tmp18049)
return


}, 1)

tmp18050 := Call(__e, PrimFunc(symgensym), symHd)


__e.TailApply(tmp18024, tmp18050)
return


}, 1)

tmp18051 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18023, tmp18051)
return


}, 5)

tmp18052 := Call(__e, ns2_1set, symshen_4conscode, tmp18022)


_ = tmp18052

tmp18053 := MakeNative(func(__e *ControlFlow) {
V348 := __e.Get(1)
_ = V348
tmp18059 := PrimIsPair(V348)

var ifres18055 Obj

if True == tmp18059 {
tmp18057 := PrimHead(V348)

tmp18058 := PrimIsPair(tmp18057)

var ifres18056 Obj

if True == tmp18058 {
ifres18056 = True


} else {
ifres18056 = False


}

ifres18055 = ifres18056


} else {
ifres18055 = False


}

if True == ifres18055 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp18060 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18053)


_ = tmp18060

tmp18061 := MakeNative(func(__e *ControlFlow) {
V349 := __e.Get(1)
_ = V349
tmp18088 := PrimIsPair(V349)

var ifres18069 Obj

if True == tmp18088 {
tmp18086 := PrimHead(V349)

tmp18087 := PrimEqual(symcons, tmp18086)

var ifres18071 Obj

if True == tmp18087 {
tmp18084 := PrimTail(V349)

tmp18085 := PrimIsPair(tmp18084)

var ifres18073 Obj

if True == tmp18085 {
tmp18081 := PrimTail(V349)

tmp18082 := PrimTail(tmp18081)

tmp18083 := PrimIsPair(tmp18082)

var ifres18075 Obj

if True == tmp18083 {
tmp18077 := PrimTail(V349)

tmp18078 := PrimTail(tmp18077)

tmp18079 := PrimTail(tmp18078)

tmp18080 := PrimEqual(Nil, tmp18079)

var ifres18076 Obj

if True == tmp18080 {
ifres18076 = True


} else {
ifres18076 = False


}

ifres18075 = ifres18076


} else {
ifres18075 = False


}

var ifres18074 Obj

if True == ifres18075 {
ifres18074 = True


} else {
ifres18074 = False


}

ifres18073 = ifres18074


} else {
ifres18073 = False


}

var ifres18072 Obj

if True == ifres18073 {
ifres18072 = True


} else {
ifres18072 = False


}

ifres18071 = ifres18072


} else {
ifres18071 = False


}

var ifres18070 Obj

if True == ifres18071 {
ifres18070 = True


} else {
ifres18070 = False


}

ifres18069 = ifres18070


} else {
ifres18069 = False


}

if True == ifres18069 {
tmp18062 := PrimTail(V349)

tmp18063 := PrimHead(tmp18062)

tmp18064 := PrimTail(V349)

tmp18065 := PrimTail(tmp18064)

tmp18066 := PrimHead(tmp18065)

tmp18067 := Call(__e, PrimFunc(symshen_4decons), tmp18066)


__e.Return(PrimCons(tmp18063, tmp18067))
return


} else {
__e.Return(V349)
return
}


}, 1)

tmp18089 := Call(__e, ns2_1set, symshen_4decons, tmp18061)


_ = tmp18089

tmp18090 := MakeNative(func(__e *ControlFlow) {
V350 := __e.Get(1)
_ = V350
V351 := __e.Get(2)
_ = V351
tmp18091 := PrimCons(V351, Nil)

__e.Return(PrimCons(V350, tmp18091))
return


}, 2)

tmp18092 := Call(__e, ns2_1set, symshen_4comb, tmp18090)


_ = tmp18092

tmp18093 := MakeNative(func(__e *ControlFlow) {
V356 := __e.Get(1)
_ = V356
V357 := __e.Get(2)
_ = V357
V358 := __e.Get(3)
_ = V358
tmp18115 := PrimIsPair(V358)

var ifres18102 Obj

if True == tmp18115 {
tmp18113 := PrimHead(V358)

tmp18114 := PrimEqual(symshen_4processed, tmp18113)

var ifres18104 Obj

if True == tmp18114 {
tmp18111 := PrimTail(V358)

tmp18112 := PrimIsPair(tmp18111)

var ifres18106 Obj

if True == tmp18112 {
tmp18108 := PrimTail(V358)

tmp18109 := PrimTail(tmp18108)

tmp18110 := PrimEqual(Nil, tmp18109)

var ifres18107 Obj

if True == tmp18110 {
ifres18107 = True


} else {
ifres18107 = False


}

ifres18106 = ifres18107


} else {
ifres18106 = False


}

var ifres18105 Obj

if True == ifres18106 {
ifres18105 = True


} else {
ifres18105 = False


}

ifres18104 = ifres18105


} else {
ifres18104 = False


}

var ifres18103 Obj

if True == ifres18104 {
ifres18103 = True


} else {
ifres18103 = False


}

ifres18102 = ifres18103


} else {
ifres18102 = False


}

if True == ifres18102 {
tmp18094 := PrimTail(V358)

__e.Return(PrimHead(tmp18094))
return


} else {
tmp18095 := MakeNative(func(__e *ControlFlow) {
W359 := __e.Get(1)
_ = W359
tmp18096 := MakeNative(func(__e *ControlFlow) {
W360 := __e.Get(1)
_ = W360
tmp18097 := PrimCons(W360, Nil)

tmp18098 := PrimCons(V357, tmp18097)

__e.Return(PrimCons(symshen_4comb, tmp18098))
return


}, 1)

tmp18099 := Call(__e, PrimFunc(symshen_4use_1type_1info), V356, W359)


__e.TailApply(tmp18096, tmp18099)
return


}, 1)

tmp18100 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V358)


__e.TailApply(tmp18095, tmp18100)
return


}


}, 3)

tmp18116 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18093)


_ = tmp18116

tmp18117 := MakeNative(func(__e *ControlFlow) {
V364 := __e.Get(1)
_ = V364
V365 := __e.Get(2)
_ = V365
tmp18305 := PrimIsPair(V364)

var ifres18126 Obj

if True == tmp18305 {
tmp18303 := PrimHead(V364)

tmp18304 := PrimEqual(sym_i, tmp18303)

var ifres18128 Obj

if True == tmp18304 {
tmp18301 := PrimTail(V364)

tmp18302 := PrimIsPair(tmp18301)

var ifres18130 Obj

if True == tmp18302 {
tmp18298 := PrimTail(V364)

tmp18299 := PrimHead(tmp18298)

tmp18300 := PrimIsPair(tmp18299)

var ifres18132 Obj

if True == tmp18300 {
tmp18294 := PrimTail(V364)

tmp18295 := PrimHead(tmp18294)

tmp18296 := PrimHead(tmp18295)

tmp18297 := PrimEqual(symlist, tmp18296)

var ifres18134 Obj

if True == tmp18297 {
tmp18290 := PrimTail(V364)

tmp18291 := PrimHead(tmp18290)

tmp18292 := PrimTail(tmp18291)

tmp18293 := PrimIsPair(tmp18292)

var ifres18136 Obj

if True == tmp18293 {
tmp18285 := PrimTail(V364)

tmp18286 := PrimHead(tmp18285)

tmp18287 := PrimTail(tmp18286)

tmp18288 := PrimTail(tmp18287)

tmp18289 := PrimEqual(Nil, tmp18288)

var ifres18138 Obj

if True == tmp18289 {
tmp18282 := PrimTail(V364)

tmp18283 := PrimTail(tmp18282)

tmp18284 := PrimIsPair(tmp18283)

var ifres18140 Obj

if True == tmp18284 {
tmp18278 := PrimTail(V364)

tmp18279 := PrimTail(tmp18278)

tmp18280 := PrimHead(tmp18279)

tmp18281 := PrimEqual(sym_1_1_6, tmp18280)

var ifres18142 Obj

if True == tmp18281 {
tmp18274 := PrimTail(V364)

tmp18275 := PrimTail(tmp18274)

tmp18276 := PrimTail(tmp18275)

tmp18277 := PrimIsPair(tmp18276)

var ifres18144 Obj

if True == tmp18277 {
tmp18269 := PrimTail(V364)

tmp18270 := PrimTail(tmp18269)

tmp18271 := PrimTail(tmp18270)

tmp18272 := PrimHead(tmp18271)

tmp18273 := PrimIsPair(tmp18272)

var ifres18146 Obj

if True == tmp18273 {
tmp18263 := PrimTail(V364)

tmp18264 := PrimTail(tmp18263)

tmp18265 := PrimTail(tmp18264)

tmp18266 := PrimHead(tmp18265)

tmp18267 := PrimHead(tmp18266)

tmp18268 := PrimEqual(symstr, tmp18267)

var ifres18148 Obj

if True == tmp18268 {
tmp18257 := PrimTail(V364)

tmp18258 := PrimTail(tmp18257)

tmp18259 := PrimTail(tmp18258)

tmp18260 := PrimHead(tmp18259)

tmp18261 := PrimTail(tmp18260)

tmp18262 := PrimIsPair(tmp18261)

var ifres18150 Obj

if True == tmp18262 {
tmp18250 := PrimTail(V364)

tmp18251 := PrimTail(tmp18250)

tmp18252 := PrimTail(tmp18251)

tmp18253 := PrimHead(tmp18252)

tmp18254 := PrimTail(tmp18253)

tmp18255 := PrimHead(tmp18254)

tmp18256 := PrimIsPair(tmp18255)

var ifres18152 Obj

if True == tmp18256 {
tmp18242 := PrimTail(V364)

tmp18243 := PrimTail(tmp18242)

tmp18244 := PrimTail(tmp18243)

tmp18245 := PrimHead(tmp18244)

tmp18246 := PrimTail(tmp18245)

tmp18247 := PrimHead(tmp18246)

tmp18248 := PrimHead(tmp18247)

tmp18249 := PrimEqual(symlist, tmp18248)

var ifres18154 Obj

if True == tmp18249 {
tmp18234 := PrimTail(V364)

tmp18235 := PrimTail(tmp18234)

tmp18236 := PrimTail(tmp18235)

tmp18237 := PrimHead(tmp18236)

tmp18238 := PrimTail(tmp18237)

tmp18239 := PrimHead(tmp18238)

tmp18240 := PrimTail(tmp18239)

tmp18241 := PrimIsPair(tmp18240)

var ifres18156 Obj

if True == tmp18241 {
tmp18225 := PrimTail(V364)

tmp18226 := PrimTail(tmp18225)

tmp18227 := PrimTail(tmp18226)

tmp18228 := PrimHead(tmp18227)

tmp18229 := PrimTail(tmp18228)

tmp18230 := PrimHead(tmp18229)

tmp18231 := PrimTail(tmp18230)

tmp18232 := PrimTail(tmp18231)

tmp18233 := PrimEqual(Nil, tmp18232)

var ifres18158 Obj

if True == tmp18233 {
tmp18218 := PrimTail(V364)

tmp18219 := PrimTail(tmp18218)

tmp18220 := PrimTail(tmp18219)

tmp18221 := PrimHead(tmp18220)

tmp18222 := PrimTail(tmp18221)

tmp18223 := PrimTail(tmp18222)

tmp18224 := PrimIsPair(tmp18223)

var ifres18160 Obj

if True == tmp18224 {
tmp18210 := PrimTail(V364)

tmp18211 := PrimTail(tmp18210)

tmp18212 := PrimTail(tmp18211)

tmp18213 := PrimHead(tmp18212)

tmp18214 := PrimTail(tmp18213)

tmp18215 := PrimTail(tmp18214)

tmp18216 := PrimTail(tmp18215)

tmp18217 := PrimEqual(Nil, tmp18216)

var ifres18162 Obj

if True == tmp18217 {
tmp18205 := PrimTail(V364)

tmp18206 := PrimTail(tmp18205)

tmp18207 := PrimTail(tmp18206)

tmp18208 := PrimTail(tmp18207)

tmp18209 := PrimIsPair(tmp18208)

var ifres18164 Obj

if True == tmp18209 {
tmp18199 := PrimTail(V364)

tmp18200 := PrimTail(tmp18199)

tmp18201 := PrimTail(tmp18200)

tmp18202 := PrimTail(tmp18201)

tmp18203 := PrimHead(tmp18202)

tmp18204 := PrimEqual(sym_j, tmp18203)

var ifres18166 Obj

if True == tmp18204 {
tmp18193 := PrimTail(V364)

tmp18194 := PrimTail(tmp18193)

tmp18195 := PrimTail(tmp18194)

tmp18196 := PrimTail(tmp18195)

tmp18197 := PrimTail(tmp18196)

tmp18198 := PrimEqual(Nil, tmp18197)

var ifres18168 Obj

if True == tmp18198 {
tmp18180 := PrimTail(V364)

tmp18181 := PrimHead(tmp18180)

tmp18182 := PrimTail(tmp18181)

tmp18183 := PrimHead(tmp18182)

tmp18184 := PrimTail(V364)

tmp18185 := PrimTail(tmp18184)

tmp18186 := PrimTail(tmp18185)

tmp18187 := PrimHead(tmp18186)

tmp18188 := PrimTail(tmp18187)

tmp18189 := PrimHead(tmp18188)

tmp18190 := PrimTail(tmp18189)

tmp18191 := PrimHead(tmp18190)

tmp18192 := PrimEqual(tmp18183, tmp18191)

var ifres18170 Obj

if True == tmp18192 {
tmp18172 := PrimTail(V364)

tmp18173 := PrimTail(tmp18172)

tmp18174 := PrimTail(tmp18173)

tmp18175 := PrimHead(tmp18174)

tmp18176 := PrimTail(tmp18175)

tmp18177 := PrimTail(tmp18176)

tmp18178 := PrimHead(tmp18177)

tmp18179 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18178)


var ifres18171 Obj

if True == tmp18179 {
ifres18171 = True


} else {
ifres18171 = False


}

ifres18170 = ifres18171


} else {
ifres18170 = False


}

var ifres18169 Obj

if True == ifres18170 {
ifres18169 = True


} else {
ifres18169 = False


}

ifres18168 = ifres18169


} else {
ifres18168 = False


}

var ifres18167 Obj

if True == ifres18168 {
ifres18167 = True


} else {
ifres18167 = False


}

ifres18166 = ifres18167


} else {
ifres18166 = False


}

var ifres18165 Obj

if True == ifres18166 {
ifres18165 = True


} else {
ifres18165 = False


}

ifres18164 = ifres18165


} else {
ifres18164 = False


}

var ifres18163 Obj

if True == ifres18164 {
ifres18163 = True


} else {
ifres18163 = False


}

ifres18162 = ifres18163


} else {
ifres18162 = False


}

var ifres18161 Obj

if True == ifres18162 {
ifres18161 = True


} else {
ifres18161 = False


}

ifres18160 = ifres18161


} else {
ifres18160 = False


}

var ifres18159 Obj

if True == ifres18160 {
ifres18159 = True


} else {
ifres18159 = False


}

ifres18158 = ifres18159


} else {
ifres18158 = False


}

var ifres18157 Obj

if True == ifres18158 {
ifres18157 = True


} else {
ifres18157 = False


}

ifres18156 = ifres18157


} else {
ifres18156 = False


}

var ifres18155 Obj

if True == ifres18156 {
ifres18155 = True


} else {
ifres18155 = False


}

ifres18154 = ifres18155


} else {
ifres18154 = False


}

var ifres18153 Obj

if True == ifres18154 {
ifres18153 = True


} else {
ifres18153 = False


}

ifres18152 = ifres18153


} else {
ifres18152 = False


}

var ifres18151 Obj

if True == ifres18152 {
ifres18151 = True


} else {
ifres18151 = False


}

ifres18150 = ifres18151


} else {
ifres18150 = False


}

var ifres18149 Obj

if True == ifres18150 {
ifres18149 = True


} else {
ifres18149 = False


}

ifres18148 = ifres18149


} else {
ifres18148 = False


}

var ifres18147 Obj

if True == ifres18148 {
ifres18147 = True


} else {
ifres18147 = False


}

ifres18146 = ifres18147


} else {
ifres18146 = False


}

var ifres18145 Obj

if True == ifres18146 {
ifres18145 = True


} else {
ifres18145 = False


}

ifres18144 = ifres18145


} else {
ifres18144 = False


}

var ifres18143 Obj

if True == ifres18144 {
ifres18143 = True


} else {
ifres18143 = False


}

ifres18142 = ifres18143


} else {
ifres18142 = False


}

var ifres18141 Obj

if True == ifres18142 {
ifres18141 = True


} else {
ifres18141 = False


}

ifres18140 = ifres18141


} else {
ifres18140 = False


}

var ifres18139 Obj

if True == ifres18140 {
ifres18139 = True


} else {
ifres18139 = False


}

ifres18138 = ifres18139


} else {
ifres18138 = False


}

var ifres18137 Obj

if True == ifres18138 {
ifres18137 = True


} else {
ifres18137 = False


}

ifres18136 = ifres18137


} else {
ifres18136 = False


}

var ifres18135 Obj

if True == ifres18136 {
ifres18135 = True


} else {
ifres18135 = False


}

ifres18134 = ifres18135


} else {
ifres18134 = False


}

var ifres18133 Obj

if True == ifres18134 {
ifres18133 = True


} else {
ifres18133 = False


}

ifres18132 = ifres18133


} else {
ifres18132 = False


}

var ifres18131 Obj

if True == ifres18132 {
ifres18131 = True


} else {
ifres18131 = False


}

ifres18130 = ifres18131


} else {
ifres18130 = False


}

var ifres18129 Obj

if True == ifres18130 {
ifres18129 = True


} else {
ifres18129 = False


}

ifres18128 = ifres18129


} else {
ifres18128 = False


}

var ifres18127 Obj

if True == ifres18128 {
ifres18127 = True


} else {
ifres18127 = False


}

ifres18126 = ifres18127


} else {
ifres18126 = False


}

if True == ifres18126 {
tmp18118 := PrimTail(V364)

tmp18119 := PrimTail(tmp18118)

tmp18120 := PrimTail(tmp18119)

tmp18121 := PrimHead(tmp18120)

tmp18122 := PrimTail(tmp18121)

tmp18123 := PrimTail(tmp18122)

tmp18124 := PrimCons(V365, tmp18123)

__e.Return(PrimCons(symtype, tmp18124))
return


} else {
__e.Return(V365)
return
}


}, 2)

tmp18306 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18117)


_ = tmp18306

tmp18307 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18317 := PrimIsVariable(V368)

if True == tmp18317 {
__e.Return(False)
return
} else {
tmp18315 := PrimIsPair(V368)

if True == tmp18315 {
tmp18312 := PrimHead(V368)

tmp18313 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18312)


if True == tmp18313 {
tmp18309 := PrimTail(V368)

tmp18310 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18309)


if True == tmp18310 {
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
__e.Return(True)
return
}


}


}, 1)

tmp18318 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18307)


_ = tmp18318

tmp18319 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
tmp18345 := PrimIsPair(V369)

var ifres18327 Obj

if True == tmp18345 {
tmp18343 := PrimHead(V369)

tmp18344 := PrimEqual(symprotect, tmp18343)

var ifres18329 Obj

if True == tmp18344 {
tmp18341 := PrimTail(V369)

tmp18342 := PrimIsPair(tmp18341)

var ifres18331 Obj

if True == tmp18342 {
tmp18338 := PrimTail(V369)

tmp18339 := PrimTail(tmp18338)

tmp18340 := PrimEqual(Nil, tmp18339)

var ifres18333 Obj

if True == tmp18340 {
tmp18335 := PrimTail(V369)

tmp18336 := PrimHead(tmp18335)

tmp18337 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18336)


var ifres18334 Obj

if True == tmp18337 {
ifres18334 = True


} else {
ifres18334 = False


}

ifres18333 = ifres18334


} else {
ifres18333 = False


}

var ifres18332 Obj

if True == ifres18333 {
ifres18332 = True


} else {
ifres18332 = False


}

ifres18331 = ifres18332


} else {
ifres18331 = False


}

var ifres18330 Obj

if True == ifres18331 {
ifres18330 = True


} else {
ifres18330 = False


}

ifres18329 = ifres18330


} else {
ifres18329 = False


}

var ifres18328 Obj

if True == ifres18329 {
ifres18328 = True


} else {
ifres18328 = False


}

ifres18327 = ifres18328


} else {
ifres18327 = False


}

if True == ifres18327 {
tmp18320 := PrimTail(V369)

__e.Return(PrimHead(tmp18320))
return


} else {
tmp18325 := PrimIsPair(V369)

if True == tmp18325 {
tmp18321 := MakeNative(func(__e *ControlFlow) {
Z370 := __e.Get(1)
_ = Z370
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z370)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18321, V369)
return


} else {
tmp18323 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V369)


if True == tmp18323 {
__e.TailApply(PrimFunc(symconcat), symAction, V369)
return
} else {
__e.Return(V369)
return
}


}


}


}, 1)

tmp18346 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18319)


_ = tmp18346

tmp18347 := MakeNative(func(__e *ControlFlow) {
V371 := __e.Get(1)
_ = V371
tmp18348 := PrimTail(V371)

__e.Return(PrimHead(tmp18348))
return


}, 1)

tmp18349 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18347)


_ = tmp18349

tmp18350 := MakeNative(func(__e *ControlFlow) {
V372 := __e.Get(1)
_ = V372
__e.Return(PrimHead(V372))
return
}, 1)

tmp18351 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18350)


_ = tmp18351

tmp18352 := MakeNative(func(__e *ControlFlow) {
V373 := __e.Get(1)
_ = V373
tmp18353 := PrimCons(V373, Nil)

__e.Return(PrimCons(Nil, tmp18353))
return


}, 1)

tmp18354 := Call(__e, ns2_1set, sym_5_b_6, tmp18352)


_ = tmp18354

tmp18355 := MakeNative(func(__e *ControlFlow) {
V374 := __e.Get(1)
_ = V374
tmp18356 := PrimCons(Nil, Nil)

__e.Return(PrimCons(V374, tmp18356))
return


}, 1)

tmp18357 := Call(__e, ns2_1set, sym_5e_6, tmp18355)


_ = tmp18357

tmp18358 := MakeNative(func(__e *ControlFlow) {
V377 := __e.Get(1)
_ = V377
tmp18361 := PrimEqual(Nil, V377)

if True == tmp18361 {
tmp18359 := PrimCons(Nil, Nil)

__e.Return(PrimCons(Nil, tmp18359))
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18358)
return




}, 0)

