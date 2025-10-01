package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp13195 := MakeNative(func(__e *ControlFlow) {
V4258 := __e.Get(1)
_ = V4258
V4259 := __e.Get(2)
_ = V4259
tmp13196 := MakeNative(func(__e *ControlFlow) {
Vs := __e.Get(1)
_ = Vs
tmp13197 := MakeNative(func(__e *ControlFlow) {
A_d := __e.Get(1)
_ = A_d
tmp13198 := MakeNative(func(__e *ControlFlow) {
Curried := __e.Get(1)
_ = Curried
tmp13199 := MakeNative(func(__e *ControlFlow) {
V3660 := __e.Get(1)
_ = V3660
__e.Return(MakeNative(func(__e *ControlFlow) {
L3661 := __e.Get(1)
_ = L3661
__e.Return(MakeNative(func(__e *ControlFlow) {
K3662 := __e.Get(1)
_ = K3662
__e.Return(MakeNative(func(__e *ControlFlow) {
C3663 := __e.Get(1)
_ = C3663
tmp13200 := MakeNative(func(__e *ControlFlow) {
Out := __e.Get(1)
_ = Out
tmp13201 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13201

tmp13202 := Call(__e, PrimFunc(symshen_4deref), Vs, V3660)


tmp13203 := Call(__e, PrimFunc(symreceive), tmp13202)


tmp13204 := Call(__e, PrimFunc(symshen_4deref), A_d, V3660)


tmp13205 := Call(__e, PrimFunc(symreceive), tmp13204)


tmp13206 := MakeNative(func(__e *ControlFlow) {
tmp13207 := Call(__e, PrimFunc(symshen_4deref), Curried, V3660)


tmp13208 := Call(__e, PrimFunc(symreceive), tmp13207)


tmp13209 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), Out, V3660, L3661, K3662, C3663)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp13208, Out, V3660, L3661, K3662, tmp13209)
return


}, 0)

tmp13210 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp13203, tmp13205, Out, V3660, L3661, K3662, tmp13206)


__e.TailApply(PrimFunc(symshen_4gc), V3660, tmp13210)
return


}, 1)

tmp13211 := Call(__e, PrimFunc(symshen_4newpv), V3660)


__e.TailApply(tmp13200, tmp13211)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp13212 := Call(__e, PrimFunc(symshen_4reset_1prolog_1vector))


tmp13213 := Call(__e, tmp13199, tmp13212)


tmp13214 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp13215 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp13214)


tmp13216 := Call(__e, PrimFunc(sym_8v), True, tmp13215)


tmp13217 := Call(__e, tmp13213, tmp13216)


tmp13218 := Call(__e, tmp13217, MakeNumber(0))


tmp13219 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp13218, tmp13219)
return


}, 1)

tmp13220 := Call(__e, PrimFunc(symshen_4curry), V4258)


__e.TailApply(tmp13198, tmp13220)
return


}, 1)

tmp13221 := Call(__e, PrimFunc(symshen_4rectify_1type), V4259)


__e.TailApply(tmp13197, tmp13221)
return


}, 1)

tmp13222 := Call(__e, PrimFunc(symshen_4extract_1vars), V4259)


__e.TailApply(tmp13196, tmp13222)
return


}, 2)

tmp13223 := Call(__e, ns2_1set, symshen_4typecheck, tmp13195)


_ = tmp13223

tmp13224 := MakeNative(func(__e *ControlFlow) {
V4260 := __e.Get(1)
_ = V4260
V4261 := __e.Get(2)
_ = V4261
V4262 := __e.Get(3)
_ = V4262
V4263 := __e.Get(4)
_ = V4263
V4264 := __e.Get(5)
_ = V4264
V4265 := __e.Get(6)
_ = V4265
V4266 := __e.Get(7)
_ = V4266
tmp13225 := MakeNative(func(__e *ControlFlow) {
C3672 := __e.Get(1)
_ = C3672
tmp13243 := PrimEqual(C3672, False)

if True == tmp13243 {
tmp13241 := Call(__e, PrimFunc(symshen_4unlocked_2), V4264)


if True == tmp13241 {
tmp13226 := MakeNative(func(__e *ControlFlow) {
Tm3674 := __e.Get(1)
_ = Tm3674
tmp13238 := PrimIsPair(Tm3674)

if True == tmp13238 {
tmp13227 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp13228 := MakeNative(func(__e *ControlFlow) {
Vs := __e.Get(1)
_ = Vs
tmp13229 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp13230 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13230

tmp13231 := Call(__e, PrimFunc(symshen_4deref), X, V4263)


tmp13232 := Call(__e, PrimFunc(symsubst), tmp13231, V, V4261)


tmp13233 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), Vs, tmp13232, V4262, V4263, V4264, V4265, V4266)


__e.TailApply(PrimFunc(symshen_4gc), V4263, tmp13233)
return


}, 1)

tmp13234 := Call(__e, PrimFunc(symshen_4newpv), V4263)


__e.TailApply(tmp13229, tmp13234)
return


}, 1)

tmp13235 := PrimTail(Tm3674)

__e.TailApply(tmp13228, tmp13235)
return


}, 1)

tmp13236 := PrimHead(Tm3674)

__e.TailApply(tmp13227, tmp13236)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13239 := Call(__e, PrimFunc(symshen_4lazyderef), V4260, V4263)


__e.TailApply(tmp13226, tmp13239)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C3672)
return
}


}, 1)

tmp13251 := Call(__e, PrimFunc(symshen_4unlocked_2), V4264)


var ifres13244 Obj

if True == tmp13251 {
tmp13245 := MakeNative(func(__e *ControlFlow) {
Tm3673 := __e.Get(1)
_ = Tm3673
tmp13248 := PrimEqual(Tm3673, Nil)

if True == tmp13248 {
tmp13246 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13246

__e.TailApply(PrimFunc(symis_b), V4261, V4262, V4263, V4264, V4265, V4266)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13249 := Call(__e, PrimFunc(symshen_4lazyderef), V4260, V4263)


tmp13250 := Call(__e, tmp13245, tmp13249)


ifres13244 = tmp13250


} else {
ifres13244 = False


}

__e.TailApply(tmp13225, ifres13244)
return


}, 7)

tmp13252 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp13224)


_ = tmp13252

tmp13253 := MakeNative(func(__e *ControlFlow) {
V4267 := __e.Get(1)
_ = V4267
V4268 := __e.Get(2)
_ = V4268
V4269 := __e.Get(3)
_ = V4269
V4270 := __e.Get(4)
_ = V4270
V4271 := __e.Get(5)
_ = V4271
V4272 := __e.Get(6)
_ = V4272
tmp13254 := MakeNative(func(__e *ControlFlow) {
K3677 := __e.Get(1)
_ = K3677
tmp13255 := MakeNative(func(__e *ControlFlow) {
C3681 := __e.Get(1)
_ = C3681
tmp13268 := PrimEqual(C3681, False)

if True == tmp13268 {
tmp13256 := MakeNative(func(__e *ControlFlow) {
C3685 := __e.Get(1)
_ = C3685
tmp13258 := PrimEqual(C3685, False)

if True == tmp13258 {
__e.TailApply(PrimFunc(symshen_4unlock), V4270, K3677)
return
} else {
__e.Return(C3685)
return
}


}, 1)

tmp13266 := Call(__e, PrimFunc(symshen_4unlocked_2), V4270)


var ifres13259 Obj

if True == tmp13266 {
tmp13260 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13260

tmp13261 := PrimIntern(MakeString(":"))

tmp13262 := PrimCons(V4268, Nil)

tmp13263 := PrimCons(tmp13261, tmp13262)

tmp13264 := PrimCons(V4267, tmp13263)

tmp13265 := Call(__e, PrimFunc(symshen_4system_1S), tmp13264, Nil, V4269, V4270, K3677, V4272)


ifres13259 = tmp13265


} else {
ifres13259 = False


}

__e.TailApply(tmp13256, ifres13259)
return


} else {
__e.Return(C3681)
return
}


}, 1)

tmp13295 := Call(__e, PrimFunc(symshen_4unlocked_2), V4270)


var ifres13269 Obj

if True == tmp13295 {
tmp13270 := MakeNative(func(__e *ControlFlow) {
Tm3682 := __e.Get(1)
_ = Tm3682
tmp13292 := PrimIsPair(Tm3682)

if True == tmp13292 {
tmp13271 := MakeNative(func(__e *ControlFlow) {
Tm3683 := __e.Get(1)
_ = Tm3683
tmp13288 := PrimEqual(Tm3683, symdefine)

if True == tmp13288 {
tmp13272 := MakeNative(func(__e *ControlFlow) {
Tm3684 := __e.Get(1)
_ = Tm3684
tmp13284 := PrimIsPair(Tm3684)

if True == tmp13284 {
tmp13273 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp13274 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp13275 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13275

tmp13276 := MakeNative(func(__e *ControlFlow) {
tmp13277 := PrimValue(symshen_4_dspy_d)

tmp13278 := MakeNative(func(__e *ControlFlow) {
tmp13279 := PrimCons(F, X)

tmp13280 := PrimCons(symdefine, tmp13279)

__e.TailApply(PrimFunc(symshen_4t_d), tmp13280, V4268, V4269, V4270, K3677, V4272)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp13277, F, V4269, V4270, K3677, tmp13278)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4269, V4270, K3677, tmp13276)
return


}, 1)

tmp13281 := PrimTail(Tm3684)

__e.TailApply(tmp13274, tmp13281)
return


}, 1)

tmp13282 := PrimHead(Tm3684)

__e.TailApply(tmp13273, tmp13282)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13285 := PrimTail(Tm3682)

tmp13286 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13285, V4269)


__e.TailApply(tmp13272, tmp13286)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13289 := PrimHead(Tm3682)

tmp13290 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13289, V4269)


__e.TailApply(tmp13271, tmp13290)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13293 := Call(__e, PrimFunc(symshen_4lazyderef), V4267, V4269)


tmp13294 := Call(__e, tmp13270, tmp13293)


ifres13269 = tmp13294


} else {
ifres13269 = False


}

__e.TailApply(tmp13255, ifres13269)
return


}, 1)

tmp13296 := PrimNumberAdd(V4271, MakeNumber(1))

__e.TailApply(tmp13254, tmp13296)
return


}, 6)

tmp13297 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp13253)


_ = tmp13297

tmp13298 := MakeNative(func(__e *ControlFlow) {
V4273 := __e.Get(1)
_ = V4273
V4274 := __e.Get(2)
_ = V4274
V4275 := __e.Get(3)
_ = V4275
V4276 := __e.Get(4)
_ = V4276
V4277 := __e.Get(5)
_ = V4277
V4278 := __e.Get(6)
_ = V4278
tmp13299 := MakeNative(func(__e *ControlFlow) {
C3692 := __e.Get(1)
_ = C3692
tmp13316 := PrimEqual(C3692, False)

if True == tmp13316 {
tmp13314 := Call(__e, PrimFunc(symshen_4unlocked_2), V4276)


if True == tmp13314 {
tmp13300 := MakeNative(func(__e *ControlFlow) {
Tm3694 := __e.Get(1)
_ = Tm3694
tmp13311 := PrimEqual(Tm3694, True)

if True == tmp13311 {
tmp13301 := MakeNative(func(__e *ControlFlow) {
ShowF := __e.Get(1)
_ = ShowF
tmp13302 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13302

tmp13303 := Call(__e, PrimFunc(symshen_4deref), V4274, V4275)


tmp13304 := Call(__e, PrimFunc(symshen_4app), tmp13303, MakeString(")\n"), symshen_4a)


tmp13305 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp13304)

tmp13306 := Call(__e, PrimFunc(symstoutput))


tmp13307 := Call(__e, PrimFunc(sympr), tmp13305, tmp13306)


tmp13308 := Call(__e, PrimFunc(symis), ShowF, tmp13307, V4275, V4276, V4277, V4278)


__e.TailApply(PrimFunc(symshen_4gc), V4275, tmp13308)
return


}, 1)

tmp13309 := Call(__e, PrimFunc(symshen_4newpv), V4275)


__e.TailApply(tmp13301, tmp13309)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13312 := Call(__e, PrimFunc(symshen_4lazyderef), V4273, V4275)


__e.TailApply(tmp13300, tmp13312)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C3692)
return
}


}, 1)

tmp13324 := Call(__e, PrimFunc(symshen_4unlocked_2), V4276)


var ifres13317 Obj

if True == tmp13324 {
tmp13318 := MakeNative(func(__e *ControlFlow) {
Tm3693 := __e.Get(1)
_ = Tm3693
tmp13321 := PrimEqual(Tm3693, False)

if True == tmp13321 {
tmp13319 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13319

__e.TailApply(PrimFunc(symthaw), V4278)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13322 := Call(__e, PrimFunc(symshen_4lazyderef), V4273, V4275)


tmp13323 := Call(__e, tmp13318, tmp13322)


ifres13317 = tmp13323


} else {
ifres13317 = False


}

__e.TailApply(tmp13299, ifres13317)
return


}, 6)

tmp13325 := Call(__e, ns2_1set, symshen_4signal_1def, tmp13298)


_ = tmp13325

tmp13326 := MakeNative(func(__e *ControlFlow) {
V4279 := __e.Get(1)
_ = V4279
tmp13327 := Call(__e, PrimFunc(symshen_4curry_1type), V4279)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp13327)
return


}, 1)

tmp13328 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp13326)


_ = tmp13328

tmp13329 := MakeNative(func(__e *ControlFlow) {
V4280 := __e.Get(1)
_ = V4280
tmp13330 := MakeNative(func(__e *ControlFlow) {
tmp13331 := MakeNative(func(__e *ControlFlow) {
Demod := __e.Get(1)
_ = Demod
tmp13333 := PrimEqual(Demod, V4280)

if True == tmp13333 {
__e.Return(V4280)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), Demod)
return
}


}, 1)

tmp13334 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
__e.TailApply(PrimFunc(symshen_4demod), Y)
return
}, 1)

tmp13335 := Call(__e, PrimFunc(symshen_4walk), tmp13334, V4280)


__e.TailApply(tmp13331, tmp13335)
return


}, 0)

tmp13336 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(V4280)
return
}, 1)

__e.TailApply(try_1catch, tmp13330, tmp13336)
return


}, 1)

tmp13337 := Call(__e, ns2_1set, symshen_4demodulate, tmp13329)


_ = tmp13337

tmp13338 := MakeNative(func(__e *ControlFlow) {
V4281 := __e.Get(1)
_ = V4281
tmp13451 := PrimIsPair(V4281)

var ifres13424 Obj

if True == tmp13451 {
tmp13449 := PrimTail(V4281)

tmp13450 := PrimIsPair(tmp13449)

var ifres13426 Obj

if True == tmp13450 {
tmp13446 := PrimTail(V4281)

tmp13447 := PrimHead(tmp13446)

tmp13448 := PrimEqual(sym_1_1_6, tmp13447)

var ifres13428 Obj

if True == tmp13448 {
tmp13443 := PrimTail(V4281)

tmp13444 := PrimTail(tmp13443)

tmp13445 := PrimIsPair(tmp13444)

var ifres13430 Obj

if True == tmp13445 {
tmp13439 := PrimTail(V4281)

tmp13440 := PrimTail(tmp13439)

tmp13441 := PrimTail(tmp13440)

tmp13442 := PrimIsPair(tmp13441)

var ifres13432 Obj

if True == tmp13442 {
tmp13434 := PrimTail(V4281)

tmp13435 := PrimTail(tmp13434)

tmp13436 := PrimTail(tmp13435)

tmp13437 := PrimHead(tmp13436)

tmp13438 := PrimEqual(sym_1_1_6, tmp13437)

var ifres13433 Obj

if True == tmp13438 {
ifres13433 = True


} else {
ifres13433 = False


}

ifres13432 = ifres13433


} else {
ifres13432 = False


}

var ifres13431 Obj

if True == ifres13432 {
ifres13431 = True


} else {
ifres13431 = False


}

ifres13430 = ifres13431


} else {
ifres13430 = False


}

var ifres13429 Obj

if True == ifres13430 {
ifres13429 = True


} else {
ifres13429 = False


}

ifres13428 = ifres13429


} else {
ifres13428 = False


}

var ifres13427 Obj

if True == ifres13428 {
ifres13427 = True


} else {
ifres13427 = False


}

ifres13426 = ifres13427


} else {
ifres13426 = False


}

var ifres13425 Obj

if True == ifres13426 {
ifres13425 = True


} else {
ifres13425 = False


}

ifres13424 = ifres13425


} else {
ifres13424 = False


}

if True == ifres13424 {
tmp13339 := PrimHead(V4281)

tmp13340 := PrimTail(V4281)

tmp13341 := PrimTail(tmp13340)

tmp13342 := PrimCons(tmp13341, Nil)

tmp13343 := PrimCons(sym_1_1_6, tmp13342)

tmp13344 := PrimCons(tmp13339, tmp13343)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp13344)
return


} else {
tmp13422 := PrimIsPair(V4281)

var ifres13402 Obj

if True == tmp13422 {
tmp13420 := PrimTail(V4281)

tmp13421 := PrimIsPair(tmp13420)

var ifres13404 Obj

if True == tmp13421 {
tmp13417 := PrimTail(V4281)

tmp13418 := PrimHead(tmp13417)

tmp13419 := PrimEqual(sym_a_a_6, tmp13418)

var ifres13406 Obj

if True == tmp13419 {
tmp13414 := PrimTail(V4281)

tmp13415 := PrimTail(tmp13414)

tmp13416 := PrimIsPair(tmp13415)

var ifres13408 Obj

if True == tmp13416 {
tmp13410 := PrimTail(V4281)

tmp13411 := PrimTail(tmp13410)

tmp13412 := PrimTail(tmp13411)

tmp13413 := PrimEqual(Nil, tmp13412)

var ifres13409 Obj

if True == tmp13413 {
ifres13409 = True


} else {
ifres13409 = False


}

ifres13408 = ifres13409


} else {
ifres13408 = False


}

var ifres13407 Obj

if True == ifres13408 {
ifres13407 = True


} else {
ifres13407 = False


}

ifres13406 = ifres13407


} else {
ifres13406 = False


}

var ifres13405 Obj

if True == ifres13406 {
ifres13405 = True


} else {
ifres13405 = False


}

ifres13404 = ifres13405


} else {
ifres13404 = False


}

var ifres13403 Obj

if True == ifres13404 {
ifres13403 = True


} else {
ifres13403 = False


}

ifres13402 = ifres13403


} else {
ifres13402 = False


}

if True == ifres13402 {
tmp13345 := PrimHead(V4281)

tmp13346 := Call(__e, PrimFunc(symprotect), symA)


tmp13347 := PrimCons(tmp13346, Nil)

tmp13348 := PrimCons(sym_d, tmp13347)

tmp13349 := PrimCons(tmp13345, tmp13348)

tmp13350 := PrimCons(symboolean, Nil)

tmp13351 := PrimCons(symvector, tmp13350)

tmp13352 := PrimHead(V4281)

tmp13353 := PrimTail(V4281)

tmp13354 := PrimTail(tmp13353)

tmp13355 := PrimCons(sym_d, tmp13354)

tmp13356 := PrimCons(tmp13352, tmp13355)

tmp13357 := PrimCons(tmp13356, Nil)

tmp13358 := PrimCons(sym_1_1_6, tmp13357)

tmp13359 := PrimCons(tmp13351, tmp13358)

tmp13360 := PrimCons(tmp13359, Nil)

tmp13361 := PrimCons(sym_1_1_6, tmp13360)

tmp13362 := PrimCons(tmp13349, tmp13361)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp13362)
return


} else {
tmp13400 := PrimIsPair(V4281)

var ifres13373 Obj

if True == tmp13400 {
tmp13398 := PrimTail(V4281)

tmp13399 := PrimIsPair(tmp13398)

var ifres13375 Obj

if True == tmp13399 {
tmp13395 := PrimTail(V4281)

tmp13396 := PrimHead(tmp13395)

tmp13397 := PrimEqual(sym_d, tmp13396)

var ifres13377 Obj

if True == tmp13397 {
tmp13392 := PrimTail(V4281)

tmp13393 := PrimTail(tmp13392)

tmp13394 := PrimIsPair(tmp13393)

var ifres13379 Obj

if True == tmp13394 {
tmp13388 := PrimTail(V4281)

tmp13389 := PrimTail(tmp13388)

tmp13390 := PrimTail(tmp13389)

tmp13391 := PrimIsPair(tmp13390)

var ifres13381 Obj

if True == tmp13391 {
tmp13383 := PrimTail(V4281)

tmp13384 := PrimTail(tmp13383)

tmp13385 := PrimTail(tmp13384)

tmp13386 := PrimHead(tmp13385)

tmp13387 := PrimEqual(sym_d, tmp13386)

var ifres13382 Obj

if True == tmp13387 {
ifres13382 = True


} else {
ifres13382 = False


}

ifres13381 = ifres13382


} else {
ifres13381 = False


}

var ifres13380 Obj

if True == ifres13381 {
ifres13380 = True


} else {
ifres13380 = False


}

ifres13379 = ifres13380


} else {
ifres13379 = False


}

var ifres13378 Obj

if True == ifres13379 {
ifres13378 = True


} else {
ifres13378 = False


}

ifres13377 = ifres13378


} else {
ifres13377 = False


}

var ifres13376 Obj

if True == ifres13377 {
ifres13376 = True


} else {
ifres13376 = False


}

ifres13375 = ifres13376


} else {
ifres13375 = False


}

var ifres13374 Obj

if True == ifres13375 {
ifres13374 = True


} else {
ifres13374 = False


}

ifres13373 = ifres13374


} else {
ifres13373 = False


}

if True == ifres13373 {
tmp13363 := PrimHead(V4281)

tmp13364 := PrimTail(V4281)

tmp13365 := PrimTail(tmp13364)

tmp13366 := PrimCons(tmp13365, Nil)

tmp13367 := PrimCons(sym_d, tmp13366)

tmp13368 := PrimCons(tmp13363, tmp13367)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp13368)
return


} else {
tmp13371 := PrimIsPair(V4281)

if True == tmp13371 {
tmp13369 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
__e.TailApply(PrimFunc(symshen_4curry_1type), Z)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13369, V4281)
return


} else {
__e.Return(V4281)
return
}


}


}


}


}, 1)

tmp13452 := Call(__e, ns2_1set, symshen_4curry_1type, tmp13338)


_ = tmp13452

tmp13453 := MakeNative(func(__e *ControlFlow) {
V4282 := __e.Get(1)
_ = V4282
tmp13571 := PrimIsPair(V4282)

var ifres13563 Obj

if True == tmp13571 {
tmp13569 := PrimHead(V4282)

tmp13570 := PrimEqual(symdefine, tmp13569)

var ifres13565 Obj

if True == tmp13570 {
tmp13567 := PrimTail(V4282)

tmp13568 := PrimIsPair(tmp13567)

var ifres13566 Obj

if True == tmp13568 {
ifres13566 = True


} else {
ifres13566 = False


}

ifres13565 = ifres13566


} else {
ifres13565 = False


}

var ifres13564 Obj

if True == ifres13565 {
ifres13564 = True


} else {
ifres13564 = False


}

ifres13563 = ifres13564


} else {
ifres13563 = False


}

if True == ifres13563 {
__e.Return(V4282)
return
} else {
tmp13561 := PrimIsPair(V4282)

var ifres13542 Obj

if True == tmp13561 {
tmp13559 := PrimHead(V4282)

tmp13560 := PrimEqual(symtype, tmp13559)

var ifres13544 Obj

if True == tmp13560 {
tmp13557 := PrimTail(V4282)

tmp13558 := PrimIsPair(tmp13557)

var ifres13546 Obj

if True == tmp13558 {
tmp13554 := PrimTail(V4282)

tmp13555 := PrimTail(tmp13554)

tmp13556 := PrimIsPair(tmp13555)

var ifres13548 Obj

if True == tmp13556 {
tmp13550 := PrimTail(V4282)

tmp13551 := PrimTail(tmp13550)

tmp13552 := PrimTail(tmp13551)

tmp13553 := PrimEqual(Nil, tmp13552)

var ifres13549 Obj

if True == tmp13553 {
ifres13549 = True


} else {
ifres13549 = False


}

ifres13548 = ifres13549


} else {
ifres13548 = False


}

var ifres13547 Obj

if True == ifres13548 {
ifres13547 = True


} else {
ifres13547 = False


}

ifres13546 = ifres13547


} else {
ifres13546 = False


}

var ifres13545 Obj

if True == ifres13546 {
ifres13545 = True


} else {
ifres13545 = False


}

ifres13544 = ifres13545


} else {
ifres13544 = False


}

var ifres13543 Obj

if True == ifres13544 {
ifres13543 = True


} else {
ifres13543 = False


}

ifres13542 = ifres13543


} else {
ifres13542 = False


}

if True == ifres13542 {
tmp13454 := PrimTail(V4282)

tmp13455 := PrimHead(tmp13454)

tmp13456 := Call(__e, PrimFunc(symshen_4curry), tmp13455)


tmp13457 := PrimTail(V4282)

tmp13458 := PrimTail(tmp13457)

tmp13459 := PrimCons(tmp13456, tmp13458)

__e.Return(PrimCons(symtype, tmp13459))
return


} else {
tmp13540 := PrimIsPair(V4282)

var ifres13521 Obj

if True == tmp13540 {
tmp13538 := PrimHead(V4282)

tmp13539 := PrimEqual(syminput_7, tmp13538)

var ifres13523 Obj

if True == tmp13539 {
tmp13536 := PrimTail(V4282)

tmp13537 := PrimIsPair(tmp13536)

var ifres13525 Obj

if True == tmp13537 {
tmp13533 := PrimTail(V4282)

tmp13534 := PrimTail(tmp13533)

tmp13535 := PrimIsPair(tmp13534)

var ifres13527 Obj

if True == tmp13535 {
tmp13529 := PrimTail(V4282)

tmp13530 := PrimTail(tmp13529)

tmp13531 := PrimTail(tmp13530)

tmp13532 := PrimEqual(Nil, tmp13531)

var ifres13528 Obj

if True == tmp13532 {
ifres13528 = True


} else {
ifres13528 = False


}

ifres13527 = ifres13528


} else {
ifres13527 = False


}

var ifres13526 Obj

if True == ifres13527 {
ifres13526 = True


} else {
ifres13526 = False


}

ifres13525 = ifres13526


} else {
ifres13525 = False


}

var ifres13524 Obj

if True == ifres13525 {
ifres13524 = True


} else {
ifres13524 = False


}

ifres13523 = ifres13524


} else {
ifres13523 = False


}

var ifres13522 Obj

if True == ifres13523 {
ifres13522 = True


} else {
ifres13522 = False


}

ifres13521 = ifres13522


} else {
ifres13521 = False


}

if True == ifres13521 {
tmp13460 := PrimTail(V4282)

tmp13461 := PrimHead(tmp13460)

tmp13462 := PrimTail(V4282)

tmp13463 := PrimTail(tmp13462)

tmp13464 := PrimHead(tmp13463)

tmp13465 := Call(__e, PrimFunc(symshen_4curry), tmp13464)


tmp13466 := PrimCons(tmp13465, Nil)

tmp13467 := PrimCons(tmp13461, tmp13466)

__e.Return(PrimCons(syminput_7, tmp13467))
return


} else {
tmp13519 := PrimIsPair(V4282)

var ifres13515 Obj

if True == tmp13519 {
tmp13517 := PrimHead(V4282)

tmp13518 := Call(__e, PrimFunc(symshen_4special_2), tmp13517)


var ifres13516 Obj

if True == tmp13518 {
ifres13516 = True


} else {
ifres13516 = False


}

ifres13515 = ifres13516


} else {
ifres13515 = False


}

if True == ifres13515 {
tmp13468 := PrimHead(V4282)

tmp13469 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
__e.TailApply(PrimFunc(symshen_4curry), Y)
return
}, 1)

tmp13470 := PrimTail(V4282)

tmp13471 := Call(__e, PrimFunc(symmap), tmp13469, tmp13470)


__e.Return(PrimCons(tmp13468, tmp13471))
return


} else {
tmp13513 := PrimIsPair(V4282)

var ifres13509 Obj

if True == tmp13513 {
tmp13511 := PrimHead(V4282)

tmp13512 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp13511)


var ifres13510 Obj

if True == tmp13512 {
ifres13510 = True


} else {
ifres13510 = False


}

ifres13509 = ifres13510


} else {
ifres13509 = False


}

if True == ifres13509 {
__e.Return(V4282)
return
} else {
tmp13507 := PrimIsPair(V4282)

var ifres13498 Obj

if True == tmp13507 {
tmp13505 := PrimTail(V4282)

tmp13506 := PrimIsPair(tmp13505)

var ifres13500 Obj

if True == tmp13506 {
tmp13502 := PrimTail(V4282)

tmp13503 := PrimTail(tmp13502)

tmp13504 := PrimIsPair(tmp13503)

var ifres13501 Obj

if True == tmp13504 {
ifres13501 = True


} else {
ifres13501 = False


}

ifres13500 = ifres13501


} else {
ifres13500 = False


}

var ifres13499 Obj

if True == ifres13500 {
ifres13499 = True


} else {
ifres13499 = False


}

ifres13498 = ifres13499


} else {
ifres13498 = False


}

if True == ifres13498 {
tmp13472 := PrimHead(V4282)

tmp13473 := PrimTail(V4282)

tmp13474 := PrimHead(tmp13473)

tmp13475 := PrimCons(tmp13474, Nil)

tmp13476 := PrimCons(tmp13472, tmp13475)

tmp13477 := PrimTail(V4282)

tmp13478 := PrimTail(tmp13477)

tmp13479 := PrimCons(tmp13476, tmp13478)

__e.TailApply(PrimFunc(symshen_4curry), tmp13479)
return


} else {
tmp13496 := PrimIsPair(V4282)

var ifres13487 Obj

if True == tmp13496 {
tmp13494 := PrimTail(V4282)

tmp13495 := PrimIsPair(tmp13494)

var ifres13489 Obj

if True == tmp13495 {
tmp13491 := PrimTail(V4282)

tmp13492 := PrimTail(tmp13491)

tmp13493 := PrimEqual(Nil, tmp13492)

var ifres13490 Obj

if True == tmp13493 {
ifres13490 = True


} else {
ifres13490 = False


}

ifres13489 = ifres13490


} else {
ifres13489 = False


}

var ifres13488 Obj

if True == ifres13489 {
ifres13488 = True


} else {
ifres13488 = False


}

ifres13487 = ifres13488


} else {
ifres13487 = False


}

if True == ifres13487 {
tmp13480 := PrimHead(V4282)

tmp13481 := Call(__e, PrimFunc(symshen_4curry), tmp13480)


tmp13482 := PrimTail(V4282)

tmp13483 := PrimHead(tmp13482)

tmp13484 := Call(__e, PrimFunc(symshen_4curry), tmp13483)


tmp13485 := PrimCons(tmp13484, Nil)

__e.Return(PrimCons(tmp13481, tmp13485))
return


} else {
__e.Return(V4282)
return
}


}


}


}


}


}


}


}, 1)

tmp13572 := Call(__e, ns2_1set, symshen_4curry, tmp13453)


_ = tmp13572

tmp13573 := MakeNative(func(__e *ControlFlow) {
V4283 := __e.Get(1)
_ = V4283
tmp13574 := PrimValue(symshen_4_dspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4283, tmp13574)
return


}, 1)

tmp13575 := Call(__e, ns2_1set, symshen_4special_2, tmp13573)


_ = tmp13575

tmp13576 := MakeNative(func(__e *ControlFlow) {
V4284 := __e.Get(1)
_ = V4284
tmp13577 := PrimValue(symshen_4_dextraspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4284, tmp13577)
return


}, 1)

tmp13578 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp13576)


_ = tmp13578

tmp13579 := MakeNative(func(__e *ControlFlow) {
V4285 := __e.Get(1)
_ = V4285
V4286 := __e.Get(2)
_ = V4286
V4287 := __e.Get(3)
_ = V4287
V4288 := __e.Get(4)
_ = V4288
V4289 := __e.Get(5)
_ = V4289
V4290 := __e.Get(6)
_ = V4290
tmp13580 := MakeNative(func(__e *ControlFlow) {
K3697 := __e.Get(1)
_ = K3697
tmp13581 := MakeNative(func(__e *ControlFlow) {
C3701 := __e.Get(1)
_ = C3701
tmp13639 := PrimEqual(C3701, False)

if True == tmp13639 {
tmp13582 := MakeNative(func(__e *ControlFlow) {
C3702 := __e.Get(1)
_ = C3702
tmp13601 := PrimEqual(C3702, False)

if True == tmp13601 {
tmp13583 := MakeNative(func(__e *ControlFlow) {
C3707 := __e.Get(1)
_ = C3707
tmp13593 := PrimEqual(C3707, False)

if True == tmp13593 {
tmp13584 := MakeNative(func(__e *ControlFlow) {
C3708 := __e.Get(1)
_ = C3708
tmp13586 := PrimEqual(C3708, False)

if True == tmp13586 {
__e.TailApply(PrimFunc(symshen_4unlock), V4288, K3697)
return
} else {
__e.Return(C3708)
return
}


}, 1)

tmp13591 := Call(__e, PrimFunc(symshen_4unlocked_2), V4288)


var ifres13587 Obj

if True == tmp13591 {
tmp13588 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13588

tmp13589 := PrimValue(symshen_4_ddatatypes_d)

tmp13590 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4285, V4286, tmp13589, V4287, V4288, K3697, V4290)


ifres13587 = tmp13590


} else {
ifres13587 = False


}

__e.TailApply(tmp13584, ifres13587)
return


} else {
__e.Return(C3707)
return
}


}, 1)

tmp13599 := Call(__e, PrimFunc(symshen_4unlocked_2), V4288)


var ifres13594 Obj

if True == tmp13599 {
tmp13595 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13595

tmp13596 := PrimValue(symshen_4_dspy_d)

tmp13597 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4285, V4286, V4287, V4288, K3697, V4290)
return
}, 0)

tmp13598 := Call(__e, PrimFunc(symwhen), tmp13596, V4287, V4288, K3697, tmp13597)


ifres13594 = tmp13598


} else {
ifres13594 = False


}

__e.TailApply(tmp13583, ifres13594)
return


} else {
__e.Return(C3702)
return
}


}, 1)

tmp13637 := Call(__e, PrimFunc(symshen_4unlocked_2), V4288)


var ifres13602 Obj

if True == tmp13637 {
tmp13603 := MakeNative(func(__e *ControlFlow) {
Tm3703 := __e.Get(1)
_ = Tm3703
tmp13634 := PrimIsPair(Tm3703)

if True == tmp13634 {
tmp13604 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp13605 := MakeNative(func(__e *ControlFlow) {
Tm3704 := __e.Get(1)
_ = Tm3704
tmp13629 := PrimIsPair(Tm3704)

if True == tmp13629 {
tmp13606 := MakeNative(func(__e *ControlFlow) {
Colon := __e.Get(1)
_ = Colon
tmp13607 := MakeNative(func(__e *ControlFlow) {
Tm3705 := __e.Get(1)
_ = Tm3705
tmp13624 := PrimIsPair(Tm3705)

if True == tmp13624 {
tmp13608 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp13609 := MakeNative(func(__e *ControlFlow) {
Tm3706 := __e.Get(1)
_ = Tm3706
tmp13619 := PrimEqual(Tm3706, Nil)

if True == tmp13619 {
tmp13610 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13610

tmp13611 := Call(__e, PrimFunc(symshen_4deref), Colon, V4287)


tmp13612 := PrimIntern(MakeString(":"))

tmp13613 := PrimEqual(tmp13611, tmp13612)

tmp13614 := MakeNative(func(__e *ControlFlow) {
tmp13615 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp13616 := MakeNative(func(__e *ControlFlow) {
tmp13617 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, A, V4286, V4287, V4288, K3697, V4290)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4287, V4288, K3697, tmp13617)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp13615, V4287, V4288, K3697, tmp13616)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp13613, V4287, V4288, K3697, tmp13614)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13620 := PrimTail(Tm3705)

tmp13621 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13620, V4287)


__e.TailApply(tmp13609, tmp13621)
return


}, 1)

tmp13622 := PrimHead(Tm3705)

__e.TailApply(tmp13608, tmp13622)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13625 := PrimTail(Tm3704)

tmp13626 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13625, V4287)


__e.TailApply(tmp13607, tmp13626)
return


}, 1)

tmp13627 := PrimHead(Tm3704)

__e.TailApply(tmp13606, tmp13627)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13630 := PrimTail(Tm3703)

tmp13631 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13630, V4287)


__e.TailApply(tmp13605, tmp13631)
return


}, 1)

tmp13632 := PrimHead(Tm3703)

__e.TailApply(tmp13604, tmp13632)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13635 := Call(__e, PrimFunc(symshen_4lazyderef), V4285, V4287)


tmp13636 := Call(__e, tmp13603, tmp13635)


ifres13602 = tmp13636


} else {
ifres13602 = False


}

__e.TailApply(tmp13582, ifres13602)
return


} else {
__e.Return(C3701)
return
}


}, 1)

tmp13644 := Call(__e, PrimFunc(symshen_4unlocked_2), V4288)


var ifres13640 Obj

if True == tmp13644 {
tmp13641 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13641

tmp13642 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp13643 := Call(__e, PrimFunc(symwhen), tmp13642, V4287, V4288, K3697, V4290)


ifres13640 = tmp13643


} else {
ifres13640 = False


}

__e.TailApply(tmp13581, ifres13640)
return


}, 1)

tmp13645 := PrimNumberAdd(V4289, MakeNumber(1))

__e.TailApply(tmp13580, tmp13645)
return


}, 6)

tmp13646 := Call(__e, ns2_1set, symshen_4system_1S, tmp13579)


_ = tmp13646

tmp13647 := MakeNative(func(__e *ControlFlow) {
V4297 := __e.Get(1)
_ = V4297
V4298 := __e.Get(2)
_ = V4298
V4299 := __e.Get(3)
_ = V4299
V4300 := __e.Get(4)
_ = V4300
V4301 := __e.Get(5)
_ = V4301
V4302 := __e.Get(6)
_ = V4302
tmp13648 := Call(__e, PrimFunc(symshen_4line))


_ = tmp13648

tmp13649 := Call(__e, PrimFunc(symshen_4deref), V4297, V4299)


tmp13650 := Call(__e, PrimFunc(symshen_4show_1p), tmp13649)


_ = tmp13650

tmp13651 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp13651

tmp13652 := Call(__e, PrimFunc(symshen_4deref), V4298, V4299)


tmp13653 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp13652, MakeNumber(1))


_ = tmp13653

tmp13654 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp13654

__e.Return(False)
return


}, 6)

tmp13655 := Call(__e, ns2_1set, symshen_4show, tmp13647)


_ = tmp13655

tmp13656 := MakeNative(func(__e *ControlFlow) {
tmp13657 := MakeNative(func(__e *ControlFlow) {
Infs := __e.Get(1)
_ = Infs
tmp13659 := PrimEqual(MakeNumber(1), Infs)

var ifres13658 Obj

if True == tmp13659 {
ifres13658 = MakeString("")


} else {
ifres13658 = MakeString("s")


}

tmp13660 := Call(__e, PrimFunc(symshen_4app), ifres13658, MakeString(" \n?- "), symshen_4a)


tmp13661 := PrimStringConcat(MakeString(" inference"), tmp13660)

tmp13662 := Call(__e, PrimFunc(symshen_4app), Infs, tmp13661, symshen_4a)


tmp13663 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp13662)

tmp13664 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13663, tmp13664)
return


}, 1)

tmp13665 := Call(__e, PrimFunc(syminferences))


__e.TailApply(tmp13657, tmp13665)
return


}, 0)

tmp13666 := Call(__e, ns2_1set, symshen_4line, tmp13656)


_ = tmp13666

tmp13667 := MakeNative(func(__e *ControlFlow) {
V4303 := __e.Get(1)
_ = V4303
tmp13699 := PrimIsPair(V4303)

var ifres13678 Obj

if True == tmp13699 {
tmp13697 := PrimTail(V4303)

tmp13698 := PrimIsPair(tmp13697)

var ifres13680 Obj

if True == tmp13698 {
tmp13694 := PrimTail(V4303)

tmp13695 := PrimTail(tmp13694)

tmp13696 := PrimIsPair(tmp13695)

var ifres13682 Obj

if True == tmp13696 {
tmp13690 := PrimTail(V4303)

tmp13691 := PrimTail(tmp13690)

tmp13692 := PrimTail(tmp13691)

tmp13693 := PrimEqual(Nil, tmp13692)

var ifres13684 Obj

if True == tmp13693 {
tmp13686 := PrimTail(V4303)

tmp13687 := PrimHead(tmp13686)

tmp13688 := PrimIntern(MakeString(":"))

tmp13689 := PrimEqual(tmp13687, tmp13688)

var ifres13685 Obj

if True == tmp13689 {
ifres13685 = True


} else {
ifres13685 = False


}

ifres13684 = ifres13685


} else {
ifres13684 = False


}

var ifres13683 Obj

if True == ifres13684 {
ifres13683 = True


} else {
ifres13683 = False


}

ifres13682 = ifres13683


} else {
ifres13682 = False


}

var ifres13681 Obj

if True == ifres13682 {
ifres13681 = True


} else {
ifres13681 = False


}

ifres13680 = ifres13681


} else {
ifres13680 = False


}

var ifres13679 Obj

if True == ifres13680 {
ifres13679 = True


} else {
ifres13679 = False


}

ifres13678 = ifres13679


} else {
ifres13678 = False


}

if True == ifres13678 {
tmp13668 := PrimHead(V4303)

tmp13669 := Call(__e, PrimFunc(symshen_4prterm), tmp13668)


_ = tmp13669

tmp13670 := Call(__e, PrimFunc(symstoutput))


tmp13671 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp13670)


_ = tmp13671

tmp13672 := PrimTail(V4303)

tmp13673 := PrimTail(tmp13672)

tmp13674 := PrimHead(tmp13673)

tmp13675 := Call(__e, PrimFunc(symshen_4app), tmp13674, MakeString(""), symshen_4r)


tmp13676 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13675, tmp13676)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4303)
return
}


}, 1)

tmp13700 := Call(__e, ns2_1set, symshen_4show_1p, tmp13667)


_ = tmp13700

tmp13701 := MakeNative(func(__e *ControlFlow) {
V4304 := __e.Get(1)
_ = V4304
tmp13744 := PrimIsPair(V4304)

var ifres13725 Obj

if True == tmp13744 {
tmp13742 := PrimHead(V4304)

tmp13743 := PrimEqual(symcons, tmp13742)

var ifres13727 Obj

if True == tmp13743 {
tmp13740 := PrimTail(V4304)

tmp13741 := PrimIsPair(tmp13740)

var ifres13729 Obj

if True == tmp13741 {
tmp13737 := PrimTail(V4304)

tmp13738 := PrimTail(tmp13737)

tmp13739 := PrimIsPair(tmp13738)

var ifres13731 Obj

if True == tmp13739 {
tmp13733 := PrimTail(V4304)

tmp13734 := PrimTail(tmp13733)

tmp13735 := PrimTail(tmp13734)

tmp13736 := PrimEqual(Nil, tmp13735)

var ifres13732 Obj

if True == tmp13736 {
ifres13732 = True


} else {
ifres13732 = False


}

ifres13731 = ifres13732


} else {
ifres13731 = False


}

var ifres13730 Obj

if True == ifres13731 {
ifres13730 = True


} else {
ifres13730 = False


}

ifres13729 = ifres13730


} else {
ifres13729 = False


}

var ifres13728 Obj

if True == ifres13729 {
ifres13728 = True


} else {
ifres13728 = False


}

ifres13727 = ifres13728


} else {
ifres13727 = False


}

var ifres13726 Obj

if True == ifres13727 {
ifres13726 = True


} else {
ifres13726 = False


}

ifres13725 = ifres13726


} else {
ifres13725 = False


}

if True == ifres13725 {
tmp13702 := Call(__e, PrimFunc(symstoutput))


tmp13703 := Call(__e, PrimFunc(sympr), MakeString("["), tmp13702)


_ = tmp13703

tmp13704 := PrimTail(V4304)

tmp13705 := PrimHead(tmp13704)

tmp13706 := Call(__e, PrimFunc(symshen_4prterm), tmp13705)


_ = tmp13706

tmp13707 := PrimTail(V4304)

tmp13708 := PrimTail(tmp13707)

tmp13709 := PrimHead(tmp13708)

tmp13710 := Call(__e, PrimFunc(symshen_4prtl), tmp13709)


_ = tmp13710

tmp13711 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp13711)
return


} else {
tmp13723 := PrimIsPair(V4304)

if True == tmp13723 {
tmp13712 := Call(__e, PrimFunc(symstoutput))


tmp13713 := Call(__e, PrimFunc(sympr), MakeString("("), tmp13712)


_ = tmp13713

tmp13714 := PrimHead(V4304)

tmp13715 := Call(__e, PrimFunc(symshen_4prterm), tmp13714)


_ = tmp13715

tmp13716 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp13717 := Call(__e, PrimFunc(symstoutput))


tmp13718 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp13717)


_ = tmp13718

__e.TailApply(PrimFunc(symshen_4prterm), Y)
return


}, 1)

tmp13719 := PrimTail(V4304)

tmp13720 := Call(__e, PrimFunc(symmap), tmp13716, tmp13719)


_ = tmp13720

tmp13721 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp13721)
return


} else {
__e.TailApply(PrimFunc(symprint), V4304)
return
}


}


}, 1)

tmp13745 := Call(__e, ns2_1set, symshen_4prterm, tmp13701)


_ = tmp13745

tmp13746 := MakeNative(func(__e *ControlFlow) {
V4305 := __e.Get(1)
_ = V4305
tmp13779 := PrimEqual(Nil, V4305)

if True == tmp13779 {
__e.Return(MakeString(""))
return
} else {
tmp13777 := PrimIsPair(V4305)

var ifres13758 Obj

if True == tmp13777 {
tmp13775 := PrimHead(V4305)

tmp13776 := PrimEqual(symcons, tmp13775)

var ifres13760 Obj

if True == tmp13776 {
tmp13773 := PrimTail(V4305)

tmp13774 := PrimIsPair(tmp13773)

var ifres13762 Obj

if True == tmp13774 {
tmp13770 := PrimTail(V4305)

tmp13771 := PrimTail(tmp13770)

tmp13772 := PrimIsPair(tmp13771)

var ifres13764 Obj

if True == tmp13772 {
tmp13766 := PrimTail(V4305)

tmp13767 := PrimTail(tmp13766)

tmp13768 := PrimTail(tmp13767)

tmp13769 := PrimEqual(Nil, tmp13768)

var ifres13765 Obj

if True == tmp13769 {
ifres13765 = True


} else {
ifres13765 = False


}

ifres13764 = ifres13765


} else {
ifres13764 = False


}

var ifres13763 Obj

if True == ifres13764 {
ifres13763 = True


} else {
ifres13763 = False


}

ifres13762 = ifres13763


} else {
ifres13762 = False


}

var ifres13761 Obj

if True == ifres13762 {
ifres13761 = True


} else {
ifres13761 = False


}

ifres13760 = ifres13761


} else {
ifres13760 = False


}

var ifres13759 Obj

if True == ifres13760 {
ifres13759 = True


} else {
ifres13759 = False


}

ifres13758 = ifres13759


} else {
ifres13758 = False


}

if True == ifres13758 {
tmp13747 := Call(__e, PrimFunc(symstoutput))


tmp13748 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp13747)


_ = tmp13748

tmp13749 := PrimTail(V4305)

tmp13750 := PrimHead(tmp13749)

tmp13751 := Call(__e, PrimFunc(symshen_4prterm), tmp13750)


_ = tmp13751

tmp13752 := PrimTail(V4305)

tmp13753 := PrimTail(tmp13752)

tmp13754 := PrimHead(tmp13753)

__e.TailApply(PrimFunc(symshen_4prtl), tmp13754)
return


} else {
tmp13755 := Call(__e, PrimFunc(symstoutput))


tmp13756 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp13755)


_ = tmp13756

__e.TailApply(PrimFunc(symshen_4prterm), V4305)
return


}


}


}, 1)

tmp13780 := Call(__e, ns2_1set, symshen_4prtl, tmp13746)


_ = tmp13780

tmp13781 := MakeNative(func(__e *ControlFlow) {
V4312 := __e.Get(1)
_ = V4312
V4313 := __e.Get(2)
_ = V4313
tmp13794 := PrimEqual(Nil, V4312)

if True == tmp13794 {
tmp13782 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp13782)
return


} else {
tmp13792 := PrimIsPair(V4312)

if True == tmp13792 {
tmp13783 := Call(__e, PrimFunc(symshen_4app), V4313, MakeString(". "), symshen_4a)


tmp13784 := Call(__e, PrimFunc(symstoutput))


tmp13785 := Call(__e, PrimFunc(sympr), tmp13783, tmp13784)


_ = tmp13785

tmp13786 := PrimHead(V4312)

tmp13787 := Call(__e, PrimFunc(symshen_4show_1p), tmp13786)


_ = tmp13787

tmp13788 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp13788

tmp13789 := PrimTail(V4312)

tmp13790 := PrimNumberAdd(V4313, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp13789, tmp13790)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
return
}


}


}, 2)

tmp13795 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp13781)


_ = tmp13795

tmp13796 := MakeNative(func(__e *ControlFlow) {
tmp13797 := MakeNative(func(__e *ControlFlow) {
Byte := __e.Get(1)
_ = Byte
tmp13799 := PrimEqual(Byte, MakeNumber(94))

if True == tmp13799 {
__e.Return(PrimSimpleError(MakeString("input aborted\n")))
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 1)

tmp13800 := Call(__e, PrimFunc(symstinput))


tmp13801 := PrimReadByte(tmp13800)

__e.TailApply(tmp13797, tmp13801)
return


}, 0)

tmp13802 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp13796)


_ = tmp13802

tmp13803 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp13804 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp13803)


_ = tmp13804

tmp13805 := MakeNative(func(__e *ControlFlow) {
tmp13807 := Call(__e, PrimFunc(syminferences))


tmp13808 := PrimValue(symshen_4_dmaxinferences_d)

tmp13809 := PrimGreatThan(tmp13807, tmp13808)

if True == tmp13809 {
__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
return
} else {
__e.Return(False)
return
}


}, 0)

tmp13810 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp13805)


_ = tmp13810

tmp13811 := MakeNative(func(__e *ControlFlow) {
V4314 := __e.Get(1)
_ = V4314
V4315 := __e.Get(2)
_ = V4315
V4316 := __e.Get(3)
_ = V4316
V4317 := __e.Get(4)
_ = V4317
V4318 := __e.Get(5)
_ = V4318
V4319 := __e.Get(6)
_ = V4319
V4320 := __e.Get(7)
_ = V4320
tmp13812 := MakeNative(func(__e *ControlFlow) {
K3712 := __e.Get(1)
_ = K3712
tmp13813 := MakeNative(func(__e *ControlFlow) {
C3717 := __e.Get(1)
_ = C3717
tmp14693 := PrimEqual(C3717, False)

if True == tmp14693 {
tmp13814 := MakeNative(func(__e *ControlFlow) {
C3718 := __e.Get(1)
_ = C3718
tmp14683 := PrimEqual(C3718, False)

if True == tmp14683 {
tmp13815 := MakeNative(func(__e *ControlFlow) {
C3719 := __e.Get(1)
_ = C3719
tmp14677 := PrimEqual(C3719, False)

if True == tmp14677 {
tmp13816 := MakeNative(func(__e *ControlFlow) {
C3720 := __e.Get(1)
_ = C3720
tmp14658 := PrimEqual(C3720, False)

if True == tmp14658 {
tmp13817 := MakeNative(func(__e *ControlFlow) {
C3723 := __e.Get(1)
_ = C3723
tmp14631 := PrimEqual(C3723, False)

if True == tmp14631 {
tmp13818 := MakeNative(func(__e *ControlFlow) {
C3728 := __e.Get(1)
_ = C3728
tmp14596 := PrimEqual(C3728, False)

if True == tmp14596 {
tmp13819 := MakeNative(func(__e *ControlFlow) {
C3732 := __e.Get(1)
_ = C3732
tmp14565 := PrimEqual(C3732, False)

if True == tmp14565 {
tmp13820 := MakeNative(func(__e *ControlFlow) {
C3736 := __e.Get(1)
_ = C3736
tmp14480 := PrimEqual(C3736, False)

if True == tmp14480 {
tmp13821 := MakeNative(func(__e *ControlFlow) {
C3750 := __e.Get(1)
_ = C3750
tmp14374 := PrimEqual(C3750, False)

if True == tmp14374 {
tmp13822 := MakeNative(func(__e *ControlFlow) {
C3766 := __e.Get(1)
_ = C3766
tmp14289 := PrimEqual(C3766, False)

if True == tmp14289 {
tmp13823 := MakeNative(func(__e *ControlFlow) {
C3780 := __e.Get(1)
_ = C3780
tmp14246 := PrimEqual(C3780, False)

if True == tmp14246 {
tmp13824 := MakeNative(func(__e *ControlFlow) {
C3788 := __e.Get(1)
_ = C3788
tmp14122 := PrimEqual(C3788, False)

if True == tmp14122 {
tmp13825 := MakeNative(func(__e *ControlFlow) {
C3804 := __e.Get(1)
_ = C3804
tmp14058 := PrimEqual(C3804, False)

if True == tmp14058 {
tmp13826 := MakeNative(func(__e *ControlFlow) {
C3811 := __e.Get(1)
_ = C3811
tmp13970 := PrimEqual(C3811, False)

if True == tmp13970 {
tmp13827 := MakeNative(func(__e *ControlFlow) {
C3825 := __e.Get(1)
_ = C3825
tmp13932 := PrimEqual(C3825, False)

if True == tmp13932 {
tmp13828 := MakeNative(func(__e *ControlFlow) {
C3831 := __e.Get(1)
_ = C3831
tmp13893 := PrimEqual(C3831, False)

if True == tmp13893 {
tmp13829 := MakeNative(func(__e *ControlFlow) {
C3837 := __e.Get(1)
_ = C3837
tmp13855 := PrimEqual(C3837, False)

if True == tmp13855 {
tmp13830 := MakeNative(func(__e *ControlFlow) {
C3843 := __e.Get(1)
_ = C3843
tmp13844 := PrimEqual(C3843, False)

if True == tmp13844 {
tmp13831 := MakeNative(func(__e *ControlFlow) {
C3844 := __e.Get(1)
_ = C3844
tmp13833 := PrimEqual(C3844, False)

if True == tmp13833 {
__e.TailApply(PrimFunc(symshen_4unlock), V4318, K3712)
return
} else {
__e.Return(C3844)
return
}


}, 1)

tmp13842 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres13834 Obj

if True == tmp13842 {
tmp13835 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13835

tmp13836 := PrimIntern(MakeString(":"))

tmp13837 := PrimCons(V4315, Nil)

tmp13838 := PrimCons(tmp13836, tmp13837)

tmp13839 := PrimCons(V4314, tmp13838)

tmp13840 := PrimValue(symshen_4_ddatatypes_d)

tmp13841 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp13839, V4316, tmp13840, V4317, V4318, K3712, V4320)


ifres13834 = tmp13841


} else {
ifres13834 = False


}

__e.TailApply(tmp13831, ifres13834)
return


} else {
__e.Return(C3843)
return
}


}, 1)

tmp13853 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres13845 Obj

if True == tmp13853 {
tmp13846 := MakeNative(func(__e *ControlFlow) {
Normalised := __e.Get(1)
_ = Normalised
tmp13847 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13847

tmp13848 := MakeNative(func(__e *ControlFlow) {
tmp13849 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4314, V4315, Normalised, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4317, V4318, K3712, tmp13849)
return


}, 0)

tmp13850 := Call(__e, PrimFunc(symshen_4l_1rules), V4316, Normalised, False, V4317, V4318, K3712, tmp13848)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp13850)
return


}, 1)

tmp13851 := Call(__e, PrimFunc(symshen_4newpv), V4317)


tmp13852 := Call(__e, tmp13846, tmp13851)


ifres13845 = tmp13852


} else {
ifres13845 = False


}

__e.TailApply(tmp13830, ifres13845)
return


} else {
__e.Return(C3837)
return
}


}, 1)

tmp13891 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres13856 Obj

if True == tmp13891 {
tmp13857 := MakeNative(func(__e *ControlFlow) {
Tm3838 := __e.Get(1)
_ = Tm3838
tmp13888 := PrimIsPair(Tm3838)

if True == tmp13888 {
tmp13858 := MakeNative(func(__e *ControlFlow) {
Tm3839 := __e.Get(1)
_ = Tm3839
tmp13884 := PrimEqual(Tm3839, symset)

if True == tmp13884 {
tmp13859 := MakeNative(func(__e *ControlFlow) {
Tm3840 := __e.Get(1)
_ = Tm3840
tmp13880 := PrimIsPair(Tm3840)

if True == tmp13880 {
tmp13860 := MakeNative(func(__e *ControlFlow) {
Var := __e.Get(1)
_ = Var
tmp13861 := MakeNative(func(__e *ControlFlow) {
Tm3841 := __e.Get(1)
_ = Tm3841
tmp13875 := PrimIsPair(Tm3841)

if True == tmp13875 {
tmp13862 := MakeNative(func(__e *ControlFlow) {
Val := __e.Get(1)
_ = Val
tmp13863 := MakeNative(func(__e *ControlFlow) {
Tm3842 := __e.Get(1)
_ = Tm3842
tmp13870 := PrimEqual(Tm3842, Nil)

if True == tmp13870 {
tmp13864 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13864

tmp13865 := MakeNative(func(__e *ControlFlow) {
tmp13866 := PrimCons(Var, Nil)

tmp13867 := PrimCons(symvalue, tmp13866)

tmp13868 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), Val, V4315, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp13867, V4315, V4316, V4317, V4318, K3712, tmp13868)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), Var, symsymbol, V4316, V4317, V4318, K3712, tmp13865)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13871 := PrimTail(Tm3841)

tmp13872 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13871, V4317)


__e.TailApply(tmp13863, tmp13872)
return


}, 1)

tmp13873 := PrimHead(Tm3841)

__e.TailApply(tmp13862, tmp13873)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13876 := PrimTail(Tm3840)

tmp13877 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13876, V4317)


__e.TailApply(tmp13861, tmp13877)
return


}, 1)

tmp13878 := PrimHead(Tm3840)

__e.TailApply(tmp13860, tmp13878)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13881 := PrimTail(Tm3838)

tmp13882 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13881, V4317)


__e.TailApply(tmp13859, tmp13882)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13885 := PrimHead(Tm3838)

tmp13886 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13885, V4317)


__e.TailApply(tmp13858, tmp13886)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13889 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp13890 := Call(__e, tmp13857, tmp13889)


ifres13856 = tmp13890


} else {
ifres13856 = False


}

__e.TailApply(tmp13829, ifres13856)
return


} else {
__e.Return(C3831)
return
}


}, 1)

tmp13930 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres13894 Obj

if True == tmp13930 {
tmp13895 := MakeNative(func(__e *ControlFlow) {
Tm3832 := __e.Get(1)
_ = Tm3832
tmp13927 := PrimIsPair(Tm3832)

if True == tmp13927 {
tmp13896 := MakeNative(func(__e *ControlFlow) {
Tm3833 := __e.Get(1)
_ = Tm3833
tmp13923 := PrimEqual(Tm3833, syminput_7)

if True == tmp13923 {
tmp13897 := MakeNative(func(__e *ControlFlow) {
Tm3834 := __e.Get(1)
_ = Tm3834
tmp13919 := PrimIsPair(Tm3834)

if True == tmp13919 {
tmp13898 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp13899 := MakeNative(func(__e *ControlFlow) {
Tm3835 := __e.Get(1)
_ = Tm3835
tmp13914 := PrimIsPair(Tm3835)

if True == tmp13914 {
tmp13900 := MakeNative(func(__e *ControlFlow) {
Stream := __e.Get(1)
_ = Stream
tmp13901 := MakeNative(func(__e *ControlFlow) {
Tm3836 := __e.Get(1)
_ = Tm3836
tmp13909 := PrimEqual(Tm3836, Nil)

if True == tmp13909 {
tmp13902 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13902

tmp13903 := Call(__e, PrimFunc(symshen_4deref), A, V4317)


tmp13904 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp13903)


tmp13905 := MakeNative(func(__e *ControlFlow) {
tmp13906 := PrimCons(symin, Nil)

tmp13907 := PrimCons(symstream, tmp13906)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), Stream, tmp13907, V4316, V4317, V4318, K3712, V4320)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V4315, tmp13904, V4317, V4318, K3712, tmp13905)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13910 := PrimTail(Tm3835)

tmp13911 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13910, V4317)


__e.TailApply(tmp13901, tmp13911)
return


}, 1)

tmp13912 := PrimHead(Tm3835)

__e.TailApply(tmp13900, tmp13912)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13915 := PrimTail(Tm3834)

tmp13916 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13915, V4317)


__e.TailApply(tmp13899, tmp13916)
return


}, 1)

tmp13917 := PrimHead(Tm3834)

__e.TailApply(tmp13898, tmp13917)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13920 := PrimTail(Tm3832)

tmp13921 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13920, V4317)


__e.TailApply(tmp13897, tmp13921)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13924 := PrimHead(Tm3832)

tmp13925 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13924, V4317)


__e.TailApply(tmp13896, tmp13925)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13928 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp13929 := Call(__e, tmp13895, tmp13928)


ifres13894 = tmp13929


} else {
ifres13894 = False


}

__e.TailApply(tmp13828, ifres13894)
return


} else {
__e.Return(C3825)
return
}


}, 1)

tmp13968 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres13933 Obj

if True == tmp13968 {
tmp13934 := MakeNative(func(__e *ControlFlow) {
Tm3826 := __e.Get(1)
_ = Tm3826
tmp13965 := PrimIsPair(Tm3826)

if True == tmp13965 {
tmp13935 := MakeNative(func(__e *ControlFlow) {
Tm3827 := __e.Get(1)
_ = Tm3827
tmp13961 := PrimEqual(Tm3827, symtype)

if True == tmp13961 {
tmp13936 := MakeNative(func(__e *ControlFlow) {
Tm3828 := __e.Get(1)
_ = Tm3828
tmp13957 := PrimIsPair(Tm3828)

if True == tmp13957 {
tmp13937 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp13938 := MakeNative(func(__e *ControlFlow) {
Tm3829 := __e.Get(1)
_ = Tm3829
tmp13952 := PrimIsPair(Tm3829)

if True == tmp13952 {
tmp13939 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp13940 := MakeNative(func(__e *ControlFlow) {
Tm3830 := __e.Get(1)
_ = Tm3830
tmp13947 := PrimEqual(Tm3830, Nil)

if True == tmp13947 {
tmp13941 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13941

tmp13942 := MakeNative(func(__e *ControlFlow) {
tmp13943 := Call(__e, PrimFunc(symshen_4deref), A, V4317)


tmp13944 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp13943)


tmp13945 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, V4315, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp13944, V4315, V4317, V4318, K3712, tmp13945)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4317, V4318, K3712, tmp13942)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13948 := PrimTail(Tm3829)

tmp13949 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13948, V4317)


__e.TailApply(tmp13940, tmp13949)
return


}, 1)

tmp13950 := PrimHead(Tm3829)

__e.TailApply(tmp13939, tmp13950)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13953 := PrimTail(Tm3828)

tmp13954 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13953, V4317)


__e.TailApply(tmp13938, tmp13954)
return


}, 1)

tmp13955 := PrimHead(Tm3828)

__e.TailApply(tmp13937, tmp13955)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13958 := PrimTail(Tm3826)

tmp13959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13958, V4317)


__e.TailApply(tmp13936, tmp13959)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13962 := PrimHead(Tm3826)

tmp13963 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13962, V4317)


__e.TailApply(tmp13935, tmp13963)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13966 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp13967 := Call(__e, tmp13934, tmp13966)


ifres13933 = tmp13967


} else {
ifres13933 = False


}

__e.TailApply(tmp13827, ifres13933)
return


} else {
__e.Return(C3811)
return
}


}, 1)

tmp14056 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres13971 Obj

if True == tmp14056 {
tmp13972 := MakeNative(func(__e *ControlFlow) {
Tm3812 := __e.Get(1)
_ = Tm3812
tmp14053 := PrimIsPair(Tm3812)

if True == tmp14053 {
tmp13973 := MakeNative(func(__e *ControlFlow) {
Tm3813 := __e.Get(1)
_ = Tm3813
tmp14049 := PrimEqual(Tm3813, symopen)

if True == tmp14049 {
tmp13974 := MakeNative(func(__e *ControlFlow) {
Tm3814 := __e.Get(1)
_ = Tm3814
tmp14045 := PrimIsPair(Tm3814)

if True == tmp14045 {
tmp13975 := MakeNative(func(__e *ControlFlow) {
File := __e.Get(1)
_ = File
tmp13976 := MakeNative(func(__e *ControlFlow) {
Tm3815 := __e.Get(1)
_ = Tm3815
tmp14040 := PrimIsPair(Tm3815)

if True == tmp14040 {
tmp13977 := MakeNative(func(__e *ControlFlow) {
V3709 := __e.Get(1)
_ = V3709
tmp13978 := MakeNative(func(__e *ControlFlow) {
Tm3816 := __e.Get(1)
_ = Tm3816
tmp14035 := PrimEqual(Tm3816, Nil)

if True == tmp14035 {
tmp13979 := MakeNative(func(__e *ControlFlow) {
Tm3817 := __e.Get(1)
_ = Tm3817
tmp13980 := MakeNative(func(__e *ControlFlow) {
GoTo3818 := __e.Get(1)
_ = GoTo3818
tmp14024 := PrimIsPair(Tm3817)

if True == tmp14024 {
tmp13981 := MakeNative(func(__e *ControlFlow) {
Tm3819 := __e.Get(1)
_ = Tm3819
tmp13982 := MakeNative(func(__e *ControlFlow) {
GoTo3820 := __e.Get(1)
_ = GoTo3820
tmp13986 := PrimEqual(Tm3819, symstream)

if True == tmp13986 {
__e.TailApply(PrimFunc(symthaw), GoTo3820)
return
} else {
tmp13984 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3819)


if True == tmp13984 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3819, symstream, V4317, GoTo3820)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13987 := MakeNative(func(__e *ControlFlow) {
tmp13988 := MakeNative(func(__e *ControlFlow) {
Tm3821 := __e.Get(1)
_ = Tm3821
tmp13989 := MakeNative(func(__e *ControlFlow) {
GoTo3822 := __e.Get(1)
_ = GoTo3822
tmp14009 := PrimIsPair(Tm3821)

if True == tmp14009 {
tmp13990 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp13991 := MakeNative(func(__e *ControlFlow) {
Tm3823 := __e.Get(1)
_ = Tm3823
tmp13992 := MakeNative(func(__e *ControlFlow) {
GoTo3824 := __e.Get(1)
_ = GoTo3824
tmp13996 := PrimEqual(Tm3823, Nil)

if True == tmp13996 {
__e.TailApply(PrimFunc(symthaw), GoTo3824)
return
} else {
tmp13994 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3823)


if True == tmp13994 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3823, Nil, V4317, GoTo3824)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13997 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3822, D)
return
}, 0)

__e.TailApply(tmp13992, tmp13997)
return


}, 1)

tmp13998 := PrimTail(Tm3821)

tmp13999 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13998, V4317)


__e.TailApply(tmp13991, tmp13999)
return


}, 1)

tmp14000 := PrimHead(Tm3821)

__e.TailApply(tmp13990, tmp14000)
return


} else {
tmp14007 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3821)


if True == tmp14007 {
tmp14001 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp14002 := PrimCons(D, Nil)

tmp14003 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3822, D)
return
}, 0)

tmp14004 := Call(__e, PrimFunc(symshen_4bind_b), Tm3821, tmp14002, V4317, tmp14003)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14004)
return


}, 1)

tmp14005 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14001, tmp14005)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14010 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
__e.TailApply(GoTo3818, D)
return
}, 1)

__e.TailApply(tmp13989, tmp14010)
return


}, 1)

tmp14011 := PrimTail(Tm3817)

tmp14012 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14011, V4317)


__e.TailApply(tmp13988, tmp14012)
return


}, 0)

__e.TailApply(tmp13982, tmp13987)
return


}, 1)

tmp14013 := PrimHead(Tm3817)

tmp14014 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14013, V4317)


__e.TailApply(tmp13981, tmp14014)
return


} else {
tmp14022 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3817)


if True == tmp14022 {
tmp14015 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp14016 := PrimCons(D, Nil)

tmp14017 := PrimCons(symstream, tmp14016)

tmp14018 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3818, D)
return
}, 0)

tmp14019 := Call(__e, PrimFunc(symshen_4bind_b), Tm3817, tmp14017, V4317, tmp14018)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14019)
return


}, 1)

tmp14020 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14015, tmp14020)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14025 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp14026 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14026

tmp14027 := MakeNative(func(__e *ControlFlow) {
tmp14028 := Call(__e, PrimFunc(symshen_4lazyderef), D, V4317)


tmp14029 := PrimCons(symout, Nil)

tmp14030 := PrimCons(symin, tmp14029)

tmp14031 := Call(__e, PrimFunc(symelement_2), tmp14028, tmp14030)


tmp14032 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), File, symstring, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14031, V4317, V4318, K3712, tmp14032)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V3709, D, V4317, V4318, K3712, tmp14027)
return


}, 1)

__e.TailApply(tmp13980, tmp14025)
return


}, 1)

tmp14033 := Call(__e, PrimFunc(symshen_4lazyderef), V4315, V4317)


__e.TailApply(tmp13979, tmp14033)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14036 := PrimTail(Tm3815)

tmp14037 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14036, V4317)


__e.TailApply(tmp13978, tmp14037)
return


}, 1)

tmp14038 := PrimHead(Tm3815)

__e.TailApply(tmp13977, tmp14038)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14041 := PrimTail(Tm3814)

tmp14042 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14041, V4317)


__e.TailApply(tmp13976, tmp14042)
return


}, 1)

tmp14043 := PrimHead(Tm3814)

__e.TailApply(tmp13975, tmp14043)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14046 := PrimTail(Tm3812)

tmp14047 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14046, V4317)


__e.TailApply(tmp13974, tmp14047)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14050 := PrimHead(Tm3812)

tmp14051 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14050, V4317)


__e.TailApply(tmp13973, tmp14051)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14054 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14055 := Call(__e, tmp13972, tmp14054)


ifres13971 = tmp14055


} else {
ifres13971 = False


}

__e.TailApply(tmp13826, ifres13971)
return


} else {
__e.Return(C3804)
return
}


}, 1)

tmp14120 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14059 Obj

if True == tmp14120 {
tmp14060 := MakeNative(func(__e *ControlFlow) {
Tm3805 := __e.Get(1)
_ = Tm3805
tmp14117 := PrimIsPair(Tm3805)

if True == tmp14117 {
tmp14061 := MakeNative(func(__e *ControlFlow) {
Tm3806 := __e.Get(1)
_ = Tm3806
tmp14113 := PrimEqual(Tm3806, symlet)

if True == tmp14113 {
tmp14062 := MakeNative(func(__e *ControlFlow) {
Tm3807 := __e.Get(1)
_ = Tm3807
tmp14109 := PrimIsPair(Tm3807)

if True == tmp14109 {
tmp14063 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14064 := MakeNative(func(__e *ControlFlow) {
Tm3808 := __e.Get(1)
_ = Tm3808
tmp14104 := PrimIsPair(Tm3808)

if True == tmp14104 {
tmp14065 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14066 := MakeNative(func(__e *ControlFlow) {
Tm3809 := __e.Get(1)
_ = Tm3809
tmp14099 := PrimIsPair(Tm3809)

if True == tmp14099 {
tmp14067 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
tmp14068 := MakeNative(func(__e *ControlFlow) {
Tm3810 := __e.Get(1)
_ = Tm3810
tmp14094 := PrimEqual(Tm3810, Nil)

if True == tmp14094 {
tmp14069 := MakeNative(func(__e *ControlFlow) {
W := __e.Get(1)
_ = W
tmp14070 := MakeNative(func(__e *ControlFlow) {
New := __e.Get(1)
_ = New
tmp14071 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14072 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14072

tmp14073 := MakeNative(func(__e *ControlFlow) {
tmp14074 := Call(__e, PrimFunc(symshen_4lazyderef), X, V4317)


tmp14075 := Call(__e, PrimFunc(symshen_4freshterm), tmp14074)


tmp14076 := MakeNative(func(__e *ControlFlow) {
tmp14077 := Call(__e, PrimFunc(symshen_4lazyderef), X, V4317)


tmp14078 := Call(__e, PrimFunc(symshen_4lazyderef), New, V4317)


tmp14079 := Call(__e, PrimFunc(symshen_4lazyderef), Z, V4317)


tmp14080 := Call(__e, PrimFunc(symshen_4beta), tmp14077, tmp14078, tmp14079)


tmp14081 := MakeNative(func(__e *ControlFlow) {
tmp14082 := PrimIntern(MakeString(":"))

tmp14083 := PrimCons(B, Nil)

tmp14084 := PrimCons(tmp14082, tmp14083)

tmp14085 := PrimCons(New, tmp14084)

tmp14086 := PrimCons(tmp14085, V4316)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W, V4315, tmp14086, V4317, V4318, K3712, V4320)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W, tmp14080, V4317, V4318, K3712, tmp14081)
return


}, 0)

__e.TailApply(PrimFunc(symbind), New, tmp14075, V4317, V4318, K3712, tmp14076)
return


}, 0)

tmp14087 := Call(__e, PrimFunc(symshen_4system_1S_1h), Y, B, V4316, V4317, V4318, K3712, tmp14073)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14087)
return


}, 1)

tmp14088 := Call(__e, PrimFunc(symshen_4newpv), V4317)


tmp14089 := Call(__e, tmp14071, tmp14088)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14089)
return


}, 1)

tmp14090 := Call(__e, PrimFunc(symshen_4newpv), V4317)


tmp14091 := Call(__e, tmp14070, tmp14090)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14091)
return


}, 1)

tmp14092 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14069, tmp14092)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14095 := PrimTail(Tm3809)

tmp14096 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14095, V4317)


__e.TailApply(tmp14068, tmp14096)
return


}, 1)

tmp14097 := PrimHead(Tm3809)

__e.TailApply(tmp14067, tmp14097)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14100 := PrimTail(Tm3808)

tmp14101 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14100, V4317)


__e.TailApply(tmp14066, tmp14101)
return


}, 1)

tmp14102 := PrimHead(Tm3808)

__e.TailApply(tmp14065, tmp14102)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14105 := PrimTail(Tm3807)

tmp14106 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14105, V4317)


__e.TailApply(tmp14064, tmp14106)
return


}, 1)

tmp14107 := PrimHead(Tm3807)

__e.TailApply(tmp14063, tmp14107)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14110 := PrimTail(Tm3805)

tmp14111 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14110, V4317)


__e.TailApply(tmp14062, tmp14111)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14114 := PrimHead(Tm3805)

tmp14115 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14114, V4317)


__e.TailApply(tmp14061, tmp14115)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14118 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14119 := Call(__e, tmp14060, tmp14118)


ifres14059 = tmp14119


} else {
ifres14059 = False


}

__e.TailApply(tmp13825, ifres14059)
return


} else {
__e.Return(C3788)
return
}


}, 1)

tmp14244 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14123 Obj

if True == tmp14244 {
tmp14124 := MakeNative(func(__e *ControlFlow) {
Tm3789 := __e.Get(1)
_ = Tm3789
tmp14241 := PrimIsPair(Tm3789)

if True == tmp14241 {
tmp14125 := MakeNative(func(__e *ControlFlow) {
Tm3790 := __e.Get(1)
_ = Tm3790
tmp14237 := PrimEqual(Tm3790, symlambda)

if True == tmp14237 {
tmp14126 := MakeNative(func(__e *ControlFlow) {
Tm3791 := __e.Get(1)
_ = Tm3791
tmp14233 := PrimIsPair(Tm3791)

if True == tmp14233 {
tmp14127 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14128 := MakeNative(func(__e *ControlFlow) {
Tm3792 := __e.Get(1)
_ = Tm3792
tmp14228 := PrimIsPair(Tm3792)

if True == tmp14228 {
tmp14129 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14130 := MakeNative(func(__e *ControlFlow) {
Tm3793 := __e.Get(1)
_ = Tm3793
tmp14223 := PrimEqual(Tm3793, Nil)

if True == tmp14223 {
tmp14131 := MakeNative(func(__e *ControlFlow) {
Tm3794 := __e.Get(1)
_ = Tm3794
tmp14132 := MakeNative(func(__e *ControlFlow) {
GoTo3795 := __e.Get(1)
_ = GoTo3795
tmp14199 := PrimIsPair(Tm3794)

if True == tmp14199 {
tmp14133 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14134 := MakeNative(func(__e *ControlFlow) {
Tm3796 := __e.Get(1)
_ = Tm3796
tmp14135 := MakeNative(func(__e *ControlFlow) {
GoTo3797 := __e.Get(1)
_ = GoTo3797
tmp14179 := PrimIsPair(Tm3796)

if True == tmp14179 {
tmp14136 := MakeNative(func(__e *ControlFlow) {
Tm3798 := __e.Get(1)
_ = Tm3798
tmp14137 := MakeNative(func(__e *ControlFlow) {
GoTo3799 := __e.Get(1)
_ = GoTo3799
tmp14141 := PrimEqual(Tm3798, sym_1_1_6)

if True == tmp14141 {
__e.TailApply(PrimFunc(symthaw), GoTo3799)
return
} else {
tmp14139 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3798)


if True == tmp14139 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3798, sym_1_1_6, V4317, GoTo3799)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14142 := MakeNative(func(__e *ControlFlow) {
tmp14143 := MakeNative(func(__e *ControlFlow) {
Tm3800 := __e.Get(1)
_ = Tm3800
tmp14144 := MakeNative(func(__e *ControlFlow) {
GoTo3801 := __e.Get(1)
_ = GoTo3801
tmp14164 := PrimIsPair(Tm3800)

if True == tmp14164 {
tmp14145 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14146 := MakeNative(func(__e *ControlFlow) {
Tm3802 := __e.Get(1)
_ = Tm3802
tmp14147 := MakeNative(func(__e *ControlFlow) {
GoTo3803 := __e.Get(1)
_ = GoTo3803
tmp14151 := PrimEqual(Tm3802, Nil)

if True == tmp14151 {
__e.TailApply(PrimFunc(symthaw), GoTo3803)
return
} else {
tmp14149 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3802)


if True == tmp14149 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3802, Nil, V4317, GoTo3803)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14152 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3801, B)
return
}, 0)

__e.TailApply(tmp14147, tmp14152)
return


}, 1)

tmp14153 := PrimTail(Tm3800)

tmp14154 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14153, V4317)


__e.TailApply(tmp14146, tmp14154)
return


}, 1)

tmp14155 := PrimHead(Tm3800)

__e.TailApply(tmp14145, tmp14155)
return


} else {
tmp14162 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3800)


if True == tmp14162 {
tmp14156 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14157 := PrimCons(B, Nil)

tmp14158 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3801, B)
return
}, 0)

tmp14159 := Call(__e, PrimFunc(symshen_4bind_b), Tm3800, tmp14157, V4317, tmp14158)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14159)
return


}, 1)

tmp14160 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14156, tmp14160)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14165 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
__e.TailApply(GoTo3797, B)
return
}, 1)

__e.TailApply(tmp14144, tmp14165)
return


}, 1)

tmp14166 := PrimTail(Tm3796)

tmp14167 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14166, V4317)


__e.TailApply(tmp14143, tmp14167)
return


}, 0)

__e.TailApply(tmp14137, tmp14142)
return


}, 1)

tmp14168 := PrimHead(Tm3796)

tmp14169 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14168, V4317)


__e.TailApply(tmp14136, tmp14169)
return


} else {
tmp14177 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3796)


if True == tmp14177 {
tmp14170 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14171 := PrimCons(B, Nil)

tmp14172 := PrimCons(sym_1_1_6, tmp14171)

tmp14173 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3797, B)
return
}, 0)

tmp14174 := Call(__e, PrimFunc(symshen_4bind_b), Tm3796, tmp14172, V4317, tmp14173)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14174)
return


}, 1)

tmp14175 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14170, tmp14175)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14180 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14181 := Call(__e, GoTo3795, A)


__e.TailApply(tmp14181, B)
return


}, 1)

__e.TailApply(tmp14135, tmp14180)
return


}, 1)

tmp14182 := PrimTail(Tm3794)

tmp14183 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14182, V4317)


__e.TailApply(tmp14134, tmp14183)
return


}, 1)

tmp14184 := PrimHead(Tm3794)

__e.TailApply(tmp14133, tmp14184)
return


} else {
tmp14197 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3794)


if True == tmp14197 {
tmp14185 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14186 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14187 := PrimCons(B, Nil)

tmp14188 := PrimCons(sym_1_1_6, tmp14187)

tmp14189 := PrimCons(A, tmp14188)

tmp14190 := MakeNative(func(__e *ControlFlow) {
tmp14191 := Call(__e, GoTo3795, A)


__e.TailApply(tmp14191, B)
return


}, 0)

tmp14192 := Call(__e, PrimFunc(symshen_4bind_b), Tm3794, tmp14189, V4317, tmp14190)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14192)
return


}, 1)

tmp14193 := Call(__e, PrimFunc(symshen_4newpv), V4317)


tmp14194 := Call(__e, tmp14186, tmp14193)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14194)
return


}, 1)

tmp14195 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14185, tmp14195)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14200 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
__e.Return(MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14201 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
tmp14202 := MakeNative(func(__e *ControlFlow) {
New := __e.Get(1)
_ = New
tmp14203 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14203

tmp14204 := Call(__e, PrimFunc(symshen_4lazyderef), X, V4317)


tmp14205 := Call(__e, PrimFunc(symshen_4freshterm), tmp14204)


tmp14206 := MakeNative(func(__e *ControlFlow) {
tmp14207 := Call(__e, PrimFunc(symshen_4lazyderef), X, V4317)


tmp14208 := Call(__e, PrimFunc(symshen_4deref), New, V4317)


tmp14209 := Call(__e, PrimFunc(symshen_4deref), Y, V4317)


tmp14210 := Call(__e, PrimFunc(symshen_4beta), tmp14207, tmp14208, tmp14209)


tmp14211 := MakeNative(func(__e *ControlFlow) {
tmp14212 := PrimIntern(MakeString(":"))

tmp14213 := PrimCons(A, Nil)

tmp14214 := PrimCons(tmp14212, tmp14213)

tmp14215 := PrimCons(New, tmp14214)

tmp14216 := PrimCons(tmp14215, V4316)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), Z, B, tmp14216, V4317, V4318, K3712, V4320)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z, tmp14210, V4317, V4318, K3712, tmp14211)
return


}, 0)

tmp14217 := Call(__e, PrimFunc(symbind), New, tmp14205, V4317, V4318, K3712, tmp14206)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14217)
return


}, 1)

tmp14218 := Call(__e, PrimFunc(symshen_4newpv), V4317)


tmp14219 := Call(__e, tmp14202, tmp14218)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14219)
return


}, 1)

tmp14220 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14201, tmp14220)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14132, tmp14200)
return


}, 1)

tmp14221 := Call(__e, PrimFunc(symshen_4lazyderef), V4315, V4317)


__e.TailApply(tmp14131, tmp14221)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14224 := PrimTail(Tm3792)

tmp14225 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14224, V4317)


__e.TailApply(tmp14130, tmp14225)
return


}, 1)

tmp14226 := PrimHead(Tm3792)

__e.TailApply(tmp14129, tmp14226)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14229 := PrimTail(Tm3791)

tmp14230 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14229, V4317)


__e.TailApply(tmp14128, tmp14230)
return


}, 1)

tmp14231 := PrimHead(Tm3791)

__e.TailApply(tmp14127, tmp14231)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14234 := PrimTail(Tm3789)

tmp14235 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14234, V4317)


__e.TailApply(tmp14126, tmp14235)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14238 := PrimHead(Tm3789)

tmp14239 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14238, V4317)


__e.TailApply(tmp14125, tmp14239)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14242 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14243 := Call(__e, tmp14124, tmp14242)


ifres14123 = tmp14243


} else {
ifres14123 = False


}

__e.TailApply(tmp13824, ifres14123)
return


} else {
__e.Return(C3780)
return
}


}, 1)

tmp14287 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14247 Obj

if True == tmp14287 {
tmp14248 := MakeNative(func(__e *ControlFlow) {
Tm3781 := __e.Get(1)
_ = Tm3781
tmp14284 := PrimIsPair(Tm3781)

if True == tmp14284 {
tmp14249 := MakeNative(func(__e *ControlFlow) {
Tm3782 := __e.Get(1)
_ = Tm3782
tmp14280 := PrimEqual(Tm3782, sym_8s)

if True == tmp14280 {
tmp14250 := MakeNative(func(__e *ControlFlow) {
Tm3783 := __e.Get(1)
_ = Tm3783
tmp14276 := PrimIsPair(Tm3783)

if True == tmp14276 {
tmp14251 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14252 := MakeNative(func(__e *ControlFlow) {
Tm3784 := __e.Get(1)
_ = Tm3784
tmp14271 := PrimIsPair(Tm3784)

if True == tmp14271 {
tmp14253 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14254 := MakeNative(func(__e *ControlFlow) {
Tm3785 := __e.Get(1)
_ = Tm3785
tmp14266 := PrimEqual(Tm3785, Nil)

if True == tmp14266 {
tmp14255 := MakeNative(func(__e *ControlFlow) {
Tm3786 := __e.Get(1)
_ = Tm3786
tmp14256 := MakeNative(func(__e *ControlFlow) {
GoTo3787 := __e.Get(1)
_ = GoTo3787
tmp14260 := PrimEqual(Tm3786, symstring)

if True == tmp14260 {
__e.TailApply(PrimFunc(symthaw), GoTo3787)
return
} else {
tmp14258 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3786)


if True == tmp14258 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3786, symstring, V4317, GoTo3787)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14261 := MakeNative(func(__e *ControlFlow) {
tmp14262 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14262

tmp14263 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), Y, symstring, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, symstring, V4316, V4317, V4318, K3712, tmp14263)
return


}, 0)

__e.TailApply(tmp14256, tmp14261)
return


}, 1)

tmp14264 := Call(__e, PrimFunc(symshen_4lazyderef), V4315, V4317)


__e.TailApply(tmp14255, tmp14264)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14267 := PrimTail(Tm3784)

tmp14268 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14267, V4317)


__e.TailApply(tmp14254, tmp14268)
return


}, 1)

tmp14269 := PrimHead(Tm3784)

__e.TailApply(tmp14253, tmp14269)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14272 := PrimTail(Tm3783)

tmp14273 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14272, V4317)


__e.TailApply(tmp14252, tmp14273)
return


}, 1)

tmp14274 := PrimHead(Tm3783)

__e.TailApply(tmp14251, tmp14274)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14277 := PrimTail(Tm3781)

tmp14278 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14277, V4317)


__e.TailApply(tmp14250, tmp14278)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14281 := PrimHead(Tm3781)

tmp14282 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14281, V4317)


__e.TailApply(tmp14249, tmp14282)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14285 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14286 := Call(__e, tmp14248, tmp14285)


ifres14247 = tmp14286


} else {
ifres14247 = False


}

__e.TailApply(tmp13823, ifres14247)
return


} else {
__e.Return(C3766)
return
}


}, 1)

tmp14372 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14290 Obj

if True == tmp14372 {
tmp14291 := MakeNative(func(__e *ControlFlow) {
Tm3767 := __e.Get(1)
_ = Tm3767
tmp14369 := PrimIsPair(Tm3767)

if True == tmp14369 {
tmp14292 := MakeNative(func(__e *ControlFlow) {
Tm3768 := __e.Get(1)
_ = Tm3768
tmp14365 := PrimEqual(Tm3768, sym_8v)

if True == tmp14365 {
tmp14293 := MakeNative(func(__e *ControlFlow) {
Tm3769 := __e.Get(1)
_ = Tm3769
tmp14361 := PrimIsPair(Tm3769)

if True == tmp14361 {
tmp14294 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14295 := MakeNative(func(__e *ControlFlow) {
Tm3770 := __e.Get(1)
_ = Tm3770
tmp14356 := PrimIsPair(Tm3770)

if True == tmp14356 {
tmp14296 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14297 := MakeNative(func(__e *ControlFlow) {
Tm3771 := __e.Get(1)
_ = Tm3771
tmp14351 := PrimEqual(Tm3771, Nil)

if True == tmp14351 {
tmp14298 := MakeNative(func(__e *ControlFlow) {
Tm3772 := __e.Get(1)
_ = Tm3772
tmp14299 := MakeNative(func(__e *ControlFlow) {
GoTo3773 := __e.Get(1)
_ = GoTo3773
tmp14343 := PrimIsPair(Tm3772)

if True == tmp14343 {
tmp14300 := MakeNative(func(__e *ControlFlow) {
Tm3774 := __e.Get(1)
_ = Tm3774
tmp14301 := MakeNative(func(__e *ControlFlow) {
GoTo3775 := __e.Get(1)
_ = GoTo3775
tmp14305 := PrimEqual(Tm3774, symvector)

if True == tmp14305 {
__e.TailApply(PrimFunc(symthaw), GoTo3775)
return
} else {
tmp14303 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3774)


if True == tmp14303 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3774, symvector, V4317, GoTo3775)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14306 := MakeNative(func(__e *ControlFlow) {
tmp14307 := MakeNative(func(__e *ControlFlow) {
Tm3776 := __e.Get(1)
_ = Tm3776
tmp14308 := MakeNative(func(__e *ControlFlow) {
GoTo3777 := __e.Get(1)
_ = GoTo3777
tmp14328 := PrimIsPair(Tm3776)

if True == tmp14328 {
tmp14309 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14310 := MakeNative(func(__e *ControlFlow) {
Tm3778 := __e.Get(1)
_ = Tm3778
tmp14311 := MakeNative(func(__e *ControlFlow) {
GoTo3779 := __e.Get(1)
_ = GoTo3779
tmp14315 := PrimEqual(Tm3778, Nil)

if True == tmp14315 {
__e.TailApply(PrimFunc(symthaw), GoTo3779)
return
} else {
tmp14313 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3778)


if True == tmp14313 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3778, Nil, V4317, GoTo3779)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14316 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3777, A)
return
}, 0)

__e.TailApply(tmp14311, tmp14316)
return


}, 1)

tmp14317 := PrimTail(Tm3776)

tmp14318 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14317, V4317)


__e.TailApply(tmp14310, tmp14318)
return


}, 1)

tmp14319 := PrimHead(Tm3776)

__e.TailApply(tmp14309, tmp14319)
return


} else {
tmp14326 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3776)


if True == tmp14326 {
tmp14320 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14321 := PrimCons(A, Nil)

tmp14322 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3777, A)
return
}, 0)

tmp14323 := Call(__e, PrimFunc(symshen_4bind_b), Tm3776, tmp14321, V4317, tmp14322)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14323)
return


}, 1)

tmp14324 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14320, tmp14324)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14329 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
__e.TailApply(GoTo3773, A)
return
}, 1)

__e.TailApply(tmp14308, tmp14329)
return


}, 1)

tmp14330 := PrimTail(Tm3772)

tmp14331 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14330, V4317)


__e.TailApply(tmp14307, tmp14331)
return


}, 0)

__e.TailApply(tmp14301, tmp14306)
return


}, 1)

tmp14332 := PrimHead(Tm3772)

tmp14333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14332, V4317)


__e.TailApply(tmp14300, tmp14333)
return


} else {
tmp14341 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3772)


if True == tmp14341 {
tmp14334 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14335 := PrimCons(A, Nil)

tmp14336 := PrimCons(symvector, tmp14335)

tmp14337 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3773, A)
return
}, 0)

tmp14338 := Call(__e, PrimFunc(symshen_4bind_b), Tm3772, tmp14336, V4317, tmp14337)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14338)
return


}, 1)

tmp14339 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14334, tmp14339)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14344 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14345 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14345

tmp14346 := MakeNative(func(__e *ControlFlow) {
tmp14347 := PrimCons(A, Nil)

tmp14348 := PrimCons(symvector, tmp14347)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), Y, tmp14348, V4316, V4317, V4318, K3712, V4320)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, A, V4316, V4317, V4318, K3712, tmp14346)
return


}, 1)

__e.TailApply(tmp14299, tmp14344)
return


}, 1)

tmp14349 := Call(__e, PrimFunc(symshen_4lazyderef), V4315, V4317)


__e.TailApply(tmp14298, tmp14349)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14352 := PrimTail(Tm3770)

tmp14353 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14352, V4317)


__e.TailApply(tmp14297, tmp14353)
return


}, 1)

tmp14354 := PrimHead(Tm3770)

__e.TailApply(tmp14296, tmp14354)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14357 := PrimTail(Tm3769)

tmp14358 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14357, V4317)


__e.TailApply(tmp14295, tmp14358)
return


}, 1)

tmp14359 := PrimHead(Tm3769)

__e.TailApply(tmp14294, tmp14359)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14362 := PrimTail(Tm3767)

tmp14363 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14362, V4317)


__e.TailApply(tmp14293, tmp14363)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14366 := PrimHead(Tm3767)

tmp14367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14366, V4317)


__e.TailApply(tmp14292, tmp14367)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14370 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14371 := Call(__e, tmp14291, tmp14370)


ifres14290 = tmp14371


} else {
ifres14290 = False


}

__e.TailApply(tmp13822, ifres14290)
return


} else {
__e.Return(C3750)
return
}


}, 1)

tmp14478 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14375 Obj

if True == tmp14478 {
tmp14376 := MakeNative(func(__e *ControlFlow) {
Tm3751 := __e.Get(1)
_ = Tm3751
tmp14475 := PrimIsPair(Tm3751)

if True == tmp14475 {
tmp14377 := MakeNative(func(__e *ControlFlow) {
Tm3752 := __e.Get(1)
_ = Tm3752
tmp14471 := PrimEqual(Tm3752, sym_8p)

if True == tmp14471 {
tmp14378 := MakeNative(func(__e *ControlFlow) {
Tm3753 := __e.Get(1)
_ = Tm3753
tmp14467 := PrimIsPair(Tm3753)

if True == tmp14467 {
tmp14379 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14380 := MakeNative(func(__e *ControlFlow) {
Tm3754 := __e.Get(1)
_ = Tm3754
tmp14462 := PrimIsPair(Tm3754)

if True == tmp14462 {
tmp14381 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14382 := MakeNative(func(__e *ControlFlow) {
Tm3755 := __e.Get(1)
_ = Tm3755
tmp14457 := PrimEqual(Tm3755, Nil)

if True == tmp14457 {
tmp14383 := MakeNative(func(__e *ControlFlow) {
Tm3756 := __e.Get(1)
_ = Tm3756
tmp14384 := MakeNative(func(__e *ControlFlow) {
GoTo3757 := __e.Get(1)
_ = GoTo3757
tmp14451 := PrimIsPair(Tm3756)

if True == tmp14451 {
tmp14385 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14386 := MakeNative(func(__e *ControlFlow) {
Tm3758 := __e.Get(1)
_ = Tm3758
tmp14387 := MakeNative(func(__e *ControlFlow) {
GoTo3759 := __e.Get(1)
_ = GoTo3759
tmp14431 := PrimIsPair(Tm3758)

if True == tmp14431 {
tmp14388 := MakeNative(func(__e *ControlFlow) {
Tm3760 := __e.Get(1)
_ = Tm3760
tmp14389 := MakeNative(func(__e *ControlFlow) {
GoTo3761 := __e.Get(1)
_ = GoTo3761
tmp14393 := PrimEqual(Tm3760, sym_d)

if True == tmp14393 {
__e.TailApply(PrimFunc(symthaw), GoTo3761)
return
} else {
tmp14391 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3760)


if True == tmp14391 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3760, sym_d, V4317, GoTo3761)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14394 := MakeNative(func(__e *ControlFlow) {
tmp14395 := MakeNative(func(__e *ControlFlow) {
Tm3762 := __e.Get(1)
_ = Tm3762
tmp14396 := MakeNative(func(__e *ControlFlow) {
GoTo3763 := __e.Get(1)
_ = GoTo3763
tmp14416 := PrimIsPair(Tm3762)

if True == tmp14416 {
tmp14397 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14398 := MakeNative(func(__e *ControlFlow) {
Tm3764 := __e.Get(1)
_ = Tm3764
tmp14399 := MakeNative(func(__e *ControlFlow) {
GoTo3765 := __e.Get(1)
_ = GoTo3765
tmp14403 := PrimEqual(Tm3764, Nil)

if True == tmp14403 {
__e.TailApply(PrimFunc(symthaw), GoTo3765)
return
} else {
tmp14401 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3764)


if True == tmp14401 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3764, Nil, V4317, GoTo3765)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14404 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3763, B)
return
}, 0)

__e.TailApply(tmp14399, tmp14404)
return


}, 1)

tmp14405 := PrimTail(Tm3762)

tmp14406 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14405, V4317)


__e.TailApply(tmp14398, tmp14406)
return


}, 1)

tmp14407 := PrimHead(Tm3762)

__e.TailApply(tmp14397, tmp14407)
return


} else {
tmp14414 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3762)


if True == tmp14414 {
tmp14408 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14409 := PrimCons(B, Nil)

tmp14410 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3763, B)
return
}, 0)

tmp14411 := Call(__e, PrimFunc(symshen_4bind_b), Tm3762, tmp14409, V4317, tmp14410)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14411)
return


}, 1)

tmp14412 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14408, tmp14412)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14417 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
__e.TailApply(GoTo3759, B)
return
}, 1)

__e.TailApply(tmp14396, tmp14417)
return


}, 1)

tmp14418 := PrimTail(Tm3758)

tmp14419 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14418, V4317)


__e.TailApply(tmp14395, tmp14419)
return


}, 0)

__e.TailApply(tmp14389, tmp14394)
return


}, 1)

tmp14420 := PrimHead(Tm3758)

tmp14421 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14420, V4317)


__e.TailApply(tmp14388, tmp14421)
return


} else {
tmp14429 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3758)


if True == tmp14429 {
tmp14422 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14423 := PrimCons(B, Nil)

tmp14424 := PrimCons(sym_d, tmp14423)

tmp14425 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3759, B)
return
}, 0)

tmp14426 := Call(__e, PrimFunc(symshen_4bind_b), Tm3758, tmp14424, V4317, tmp14425)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14426)
return


}, 1)

tmp14427 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14422, tmp14427)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14432 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14433 := Call(__e, GoTo3757, A)


__e.TailApply(tmp14433, B)
return


}, 1)

__e.TailApply(tmp14387, tmp14432)
return


}, 1)

tmp14434 := PrimTail(Tm3756)

tmp14435 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14434, V4317)


__e.TailApply(tmp14386, tmp14435)
return


}, 1)

tmp14436 := PrimHead(Tm3756)

__e.TailApply(tmp14385, tmp14436)
return


} else {
tmp14449 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3756)


if True == tmp14449 {
tmp14437 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14438 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14439 := PrimCons(B, Nil)

tmp14440 := PrimCons(sym_d, tmp14439)

tmp14441 := PrimCons(A, tmp14440)

tmp14442 := MakeNative(func(__e *ControlFlow) {
tmp14443 := Call(__e, GoTo3757, A)


__e.TailApply(tmp14443, B)
return


}, 0)

tmp14444 := Call(__e, PrimFunc(symshen_4bind_b), Tm3756, tmp14441, V4317, tmp14442)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14444)
return


}, 1)

tmp14445 := Call(__e, PrimFunc(symshen_4newpv), V4317)


tmp14446 := Call(__e, tmp14438, tmp14445)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14446)
return


}, 1)

tmp14447 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14437, tmp14447)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14452 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
__e.Return(MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14453 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14453

tmp14454 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), Y, B, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, A, V4316, V4317, V4318, K3712, tmp14454)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14384, tmp14452)
return


}, 1)

tmp14455 := Call(__e, PrimFunc(symshen_4lazyderef), V4315, V4317)


__e.TailApply(tmp14383, tmp14455)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14458 := PrimTail(Tm3754)

tmp14459 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14458, V4317)


__e.TailApply(tmp14382, tmp14459)
return


}, 1)

tmp14460 := PrimHead(Tm3754)

__e.TailApply(tmp14381, tmp14460)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14463 := PrimTail(Tm3753)

tmp14464 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14463, V4317)


__e.TailApply(tmp14380, tmp14464)
return


}, 1)

tmp14465 := PrimHead(Tm3753)

__e.TailApply(tmp14379, tmp14465)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14468 := PrimTail(Tm3751)

tmp14469 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14468, V4317)


__e.TailApply(tmp14378, tmp14469)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14472 := PrimHead(Tm3751)

tmp14473 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14472, V4317)


__e.TailApply(tmp14377, tmp14473)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14476 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14477 := Call(__e, tmp14376, tmp14476)


ifres14375 = tmp14477


} else {
ifres14375 = False


}

__e.TailApply(tmp13821, ifres14375)
return


} else {
__e.Return(C3736)
return
}


}, 1)

tmp14563 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14481 Obj

if True == tmp14563 {
tmp14482 := MakeNative(func(__e *ControlFlow) {
Tm3737 := __e.Get(1)
_ = Tm3737
tmp14560 := PrimIsPair(Tm3737)

if True == tmp14560 {
tmp14483 := MakeNative(func(__e *ControlFlow) {
Tm3738 := __e.Get(1)
_ = Tm3738
tmp14556 := PrimEqual(Tm3738, symcons)

if True == tmp14556 {
tmp14484 := MakeNative(func(__e *ControlFlow) {
Tm3739 := __e.Get(1)
_ = Tm3739
tmp14552 := PrimIsPair(Tm3739)

if True == tmp14552 {
tmp14485 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14486 := MakeNative(func(__e *ControlFlow) {
Tm3740 := __e.Get(1)
_ = Tm3740
tmp14547 := PrimIsPair(Tm3740)

if True == tmp14547 {
tmp14487 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14488 := MakeNative(func(__e *ControlFlow) {
Tm3741 := __e.Get(1)
_ = Tm3741
tmp14542 := PrimEqual(Tm3741, Nil)

if True == tmp14542 {
tmp14489 := MakeNative(func(__e *ControlFlow) {
Tm3742 := __e.Get(1)
_ = Tm3742
tmp14490 := MakeNative(func(__e *ControlFlow) {
GoTo3743 := __e.Get(1)
_ = GoTo3743
tmp14534 := PrimIsPair(Tm3742)

if True == tmp14534 {
tmp14491 := MakeNative(func(__e *ControlFlow) {
Tm3744 := __e.Get(1)
_ = Tm3744
tmp14492 := MakeNative(func(__e *ControlFlow) {
GoTo3745 := __e.Get(1)
_ = GoTo3745
tmp14496 := PrimEqual(Tm3744, symlist)

if True == tmp14496 {
__e.TailApply(PrimFunc(symthaw), GoTo3745)
return
} else {
tmp14494 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3744)


if True == tmp14494 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3744, symlist, V4317, GoTo3745)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14497 := MakeNative(func(__e *ControlFlow) {
tmp14498 := MakeNative(func(__e *ControlFlow) {
Tm3746 := __e.Get(1)
_ = Tm3746
tmp14499 := MakeNative(func(__e *ControlFlow) {
GoTo3747 := __e.Get(1)
_ = GoTo3747
tmp14519 := PrimIsPair(Tm3746)

if True == tmp14519 {
tmp14500 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14501 := MakeNative(func(__e *ControlFlow) {
Tm3748 := __e.Get(1)
_ = Tm3748
tmp14502 := MakeNative(func(__e *ControlFlow) {
GoTo3749 := __e.Get(1)
_ = GoTo3749
tmp14506 := PrimEqual(Tm3748, Nil)

if True == tmp14506 {
__e.TailApply(PrimFunc(symthaw), GoTo3749)
return
} else {
tmp14504 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3748)


if True == tmp14504 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3748, Nil, V4317, GoTo3749)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14507 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3747, A)
return
}, 0)

__e.TailApply(tmp14502, tmp14507)
return


}, 1)

tmp14508 := PrimTail(Tm3746)

tmp14509 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14508, V4317)


__e.TailApply(tmp14501, tmp14509)
return


}, 1)

tmp14510 := PrimHead(Tm3746)

__e.TailApply(tmp14500, tmp14510)
return


} else {
tmp14517 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3746)


if True == tmp14517 {
tmp14511 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14512 := PrimCons(A, Nil)

tmp14513 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3747, A)
return
}, 0)

tmp14514 := Call(__e, PrimFunc(symshen_4bind_b), Tm3746, tmp14512, V4317, tmp14513)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14514)
return


}, 1)

tmp14515 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14511, tmp14515)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14520 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
__e.TailApply(GoTo3743, A)
return
}, 1)

__e.TailApply(tmp14499, tmp14520)
return


}, 1)

tmp14521 := PrimTail(Tm3742)

tmp14522 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14521, V4317)


__e.TailApply(tmp14498, tmp14522)
return


}, 0)

__e.TailApply(tmp14492, tmp14497)
return


}, 1)

tmp14523 := PrimHead(Tm3742)

tmp14524 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14523, V4317)


__e.TailApply(tmp14491, tmp14524)
return


} else {
tmp14532 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3742)


if True == tmp14532 {
tmp14525 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14526 := PrimCons(A, Nil)

tmp14527 := PrimCons(symlist, tmp14526)

tmp14528 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3743, A)
return
}, 0)

tmp14529 := Call(__e, PrimFunc(symshen_4bind_b), Tm3742, tmp14527, V4317, tmp14528)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14529)
return


}, 1)

tmp14530 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14525, tmp14530)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14535 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14536 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14536

tmp14537 := MakeNative(func(__e *ControlFlow) {
tmp14538 := PrimCons(A, Nil)

tmp14539 := PrimCons(symlist, tmp14538)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), Y, tmp14539, V4316, V4317, V4318, K3712, V4320)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, A, V4316, V4317, V4318, K3712, tmp14537)
return


}, 1)

__e.TailApply(tmp14490, tmp14535)
return


}, 1)

tmp14540 := Call(__e, PrimFunc(symshen_4lazyderef), V4315, V4317)


__e.TailApply(tmp14489, tmp14540)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14543 := PrimTail(Tm3740)

tmp14544 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14543, V4317)


__e.TailApply(tmp14488, tmp14544)
return


}, 1)

tmp14545 := PrimHead(Tm3740)

__e.TailApply(tmp14487, tmp14545)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14548 := PrimTail(Tm3739)

tmp14549 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14548, V4317)


__e.TailApply(tmp14486, tmp14549)
return


}, 1)

tmp14550 := PrimHead(Tm3739)

__e.TailApply(tmp14485, tmp14550)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14553 := PrimTail(Tm3737)

tmp14554 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14553, V4317)


__e.TailApply(tmp14484, tmp14554)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14557 := PrimHead(Tm3737)

tmp14558 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14557, V4317)


__e.TailApply(tmp14483, tmp14558)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14561 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14562 := Call(__e, tmp14482, tmp14561)


ifres14481 = tmp14562


} else {
ifres14481 = False


}

__e.TailApply(tmp13820, ifres14481)
return


} else {
__e.Return(C3732)
return
}


}, 1)

tmp14594 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14566 Obj

if True == tmp14594 {
tmp14567 := MakeNative(func(__e *ControlFlow) {
Tm3733 := __e.Get(1)
_ = Tm3733
tmp14591 := PrimIsPair(Tm3733)

if True == tmp14591 {
tmp14568 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp14569 := MakeNative(func(__e *ControlFlow) {
Tm3734 := __e.Get(1)
_ = Tm3734
tmp14586 := PrimIsPair(Tm3734)

if True == tmp14586 {
tmp14570 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14571 := MakeNative(func(__e *ControlFlow) {
Tm3735 := __e.Get(1)
_ = Tm3735
tmp14581 := PrimEqual(Tm3735, Nil)

if True == tmp14581 {
tmp14572 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14573 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14573

tmp14574 := PrimCons(V4315, Nil)

tmp14575 := PrimCons(sym_1_1_6, tmp14574)

tmp14576 := PrimCons(B, tmp14575)

tmp14577 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, B, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

tmp14578 := Call(__e, PrimFunc(symshen_4system_1S_1h), F, tmp14576, V4316, V4317, V4318, K3712, tmp14577)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14578)
return


}, 1)

tmp14579 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14572, tmp14579)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14582 := PrimTail(Tm3734)

tmp14583 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14582, V4317)


__e.TailApply(tmp14571, tmp14583)
return


}, 1)

tmp14584 := PrimHead(Tm3734)

__e.TailApply(tmp14570, tmp14584)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14587 := PrimTail(Tm3733)

tmp14588 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14587, V4317)


__e.TailApply(tmp14569, tmp14588)
return


}, 1)

tmp14589 := PrimHead(Tm3733)

__e.TailApply(tmp14568, tmp14589)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14592 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14593 := Call(__e, tmp14567, tmp14592)


ifres14566 = tmp14593


} else {
ifres14566 = False


}

__e.TailApply(tmp13819, ifres14566)
return


} else {
__e.Return(C3728)
return
}


}, 1)

tmp14629 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14597 Obj

if True == tmp14629 {
tmp14598 := MakeNative(func(__e *ControlFlow) {
Tm3729 := __e.Get(1)
_ = Tm3729
tmp14626 := PrimIsPair(Tm3729)

if True == tmp14626 {
tmp14599 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp14600 := MakeNative(func(__e *ControlFlow) {
Tm3730 := __e.Get(1)
_ = Tm3730
tmp14621 := PrimIsPair(Tm3730)

if True == tmp14621 {
tmp14601 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp14602 := MakeNative(func(__e *ControlFlow) {
Tm3731 := __e.Get(1)
_ = Tm3731
tmp14616 := PrimEqual(Tm3731, Nil)

if True == tmp14616 {
tmp14603 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14604 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14604

tmp14605 := Call(__e, PrimFunc(symshen_4lazyderef), F, V4317)


tmp14606 := PrimIsPair(tmp14605)

tmp14607 := PrimNot(tmp14606)

tmp14608 := MakeNative(func(__e *ControlFlow) {
tmp14609 := PrimCons(V4315, Nil)

tmp14610 := PrimCons(sym_1_1_6, tmp14609)

tmp14611 := PrimCons(B, tmp14610)

tmp14612 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), X, B, V4316, V4317, V4318, K3712, V4320)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), F, tmp14611, V4317, V4318, K3712, tmp14612)
return


}, 0)

tmp14613 := Call(__e, PrimFunc(symwhen), tmp14607, V4317, V4318, K3712, tmp14608)


__e.TailApply(PrimFunc(symshen_4gc), V4317, tmp14613)
return


}, 1)

tmp14614 := Call(__e, PrimFunc(symshen_4newpv), V4317)


__e.TailApply(tmp14603, tmp14614)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14617 := PrimTail(Tm3730)

tmp14618 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14617, V4317)


__e.TailApply(tmp14602, tmp14618)
return


}, 1)

tmp14619 := PrimHead(Tm3730)

__e.TailApply(tmp14601, tmp14619)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14622 := PrimTail(Tm3729)

tmp14623 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14622, V4317)


__e.TailApply(tmp14600, tmp14623)
return


}, 1)

tmp14624 := PrimHead(Tm3729)

__e.TailApply(tmp14599, tmp14624)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14627 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14628 := Call(__e, tmp14598, tmp14627)


ifres14597 = tmp14628


} else {
ifres14597 = False


}

__e.TailApply(tmp13818, ifres14597)
return


} else {
__e.Return(C3723)
return
}


}, 1)

tmp14656 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14632 Obj

if True == tmp14656 {
tmp14633 := MakeNative(func(__e *ControlFlow) {
Tm3724 := __e.Get(1)
_ = Tm3724
tmp14653 := PrimIsPair(Tm3724)

if True == tmp14653 {
tmp14634 := MakeNative(func(__e *ControlFlow) {
Tm3725 := __e.Get(1)
_ = Tm3725
tmp14649 := PrimEqual(Tm3725, symfn)

if True == tmp14649 {
tmp14635 := MakeNative(func(__e *ControlFlow) {
Tm3726 := __e.Get(1)
_ = Tm3726
tmp14645 := PrimIsPair(Tm3726)

if True == tmp14645 {
tmp14636 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp14637 := MakeNative(func(__e *ControlFlow) {
Tm3727 := __e.Get(1)
_ = Tm3727
tmp14640 := PrimEqual(Tm3727, Nil)

if True == tmp14640 {
tmp14638 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14638

__e.TailApply(PrimFunc(symshen_4lookupsig), F, V4315, V4317, V4318, K3712, V4320)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14641 := PrimTail(Tm3726)

tmp14642 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14641, V4317)


__e.TailApply(tmp14637, tmp14642)
return


}, 1)

tmp14643 := PrimHead(Tm3726)

__e.TailApply(tmp14636, tmp14643)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14646 := PrimTail(Tm3724)

tmp14647 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14646, V4317)


__e.TailApply(tmp14635, tmp14647)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14650 := PrimHead(Tm3724)

tmp14651 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14650, V4317)


__e.TailApply(tmp14634, tmp14651)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14654 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14655 := Call(__e, tmp14633, tmp14654)


ifres14632 = tmp14655


} else {
ifres14632 = False


}

__e.TailApply(tmp13817, ifres14632)
return


} else {
__e.Return(C3720)
return
}


}, 1)

tmp14675 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14659 Obj

if True == tmp14675 {
tmp14660 := MakeNative(func(__e *ControlFlow) {
Tm3721 := __e.Get(1)
_ = Tm3721
tmp14672 := PrimIsPair(Tm3721)

if True == tmp14672 {
tmp14661 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp14662 := MakeNative(func(__e *ControlFlow) {
Tm3722 := __e.Get(1)
_ = Tm3722
tmp14667 := PrimEqual(Tm3722, Nil)

if True == tmp14667 {
tmp14663 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14663

tmp14664 := PrimCons(V4315, Nil)

tmp14665 := PrimCons(sym_1_1_6, tmp14664)

__e.TailApply(PrimFunc(symshen_4lookupsig), F, tmp14665, V4317, V4318, K3712, V4320)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14668 := PrimTail(Tm3721)

tmp14669 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14668, V4317)


__e.TailApply(tmp14662, tmp14669)
return


}, 1)

tmp14670 := PrimHead(Tm3721)

__e.TailApply(tmp14661, tmp14670)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14673 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14674 := Call(__e, tmp14660, tmp14673)


ifres14659 = tmp14674


} else {
ifres14659 = False


}

__e.TailApply(tmp13816, ifres14659)
return


} else {
__e.Return(C3719)
return
}


}, 1)

tmp14681 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14678 Obj

if True == tmp14681 {
tmp14679 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14679

tmp14680 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4314, V4315, V4316, V4317, V4318, K3712, V4320)


ifres14678 = tmp14680


} else {
ifres14678 = False


}

__e.TailApply(tmp13815, ifres14678)
return


} else {
__e.Return(C3718)
return
}


}, 1)

tmp14691 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14684 Obj

if True == tmp14691 {
tmp14685 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14685

tmp14686 := Call(__e, PrimFunc(symshen_4lazyderef), V4314, V4317)


tmp14687 := PrimIsPair(tmp14686)

tmp14688 := PrimNot(tmp14687)

tmp14689 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4314, V4315, V4317, V4318, K3712, V4320)
return
}, 0)

tmp14690 := Call(__e, PrimFunc(symwhen), tmp14688, V4317, V4318, K3712, tmp14689)


ifres14684 = tmp14690


} else {
ifres14684 = False


}

__e.TailApply(tmp13814, ifres14684)
return


} else {
__e.Return(C3717)
return
}


}, 1)

tmp14703 := Call(__e, PrimFunc(symshen_4unlocked_2), V4318)


var ifres14694 Obj

if True == tmp14703 {
tmp14695 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14695

tmp14696 := PrimValue(symshen_4_dspy_d)

tmp14697 := MakeNative(func(__e *ControlFlow) {
tmp14698 := PrimIntern(MakeString(":"))

tmp14699 := PrimCons(V4315, Nil)

tmp14700 := PrimCons(tmp14698, tmp14699)

tmp14701 := PrimCons(V4314, tmp14700)

__e.TailApply(PrimFunc(symshen_4show), tmp14701, V4316, V4317, V4318, K3712, V4320)
return


}, 0)

tmp14702 := Call(__e, PrimFunc(symwhen), tmp14696, V4317, V4318, K3712, tmp14697)


ifres14694 = tmp14702


} else {
ifres14694 = False


}

__e.TailApply(tmp13813, ifres14694)
return


}, 1)

tmp14704 := PrimNumberAdd(V4319, MakeNumber(1))

__e.TailApply(tmp13812, tmp14704)
return


}, 7)

tmp14705 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp13811)


_ = tmp14705

tmp14706 := MakeNative(func(__e *ControlFlow) {
V4321 := __e.Get(1)
_ = V4321
V4322 := __e.Get(2)
_ = V4322
V4323 := __e.Get(3)
_ = V4323
V4324 := __e.Get(4)
_ = V4324
V4325 := __e.Get(5)
_ = V4325
V4326 := __e.Get(6)
_ = V4326
tmp14707 := MakeNative(func(__e *ControlFlow) {
C3851 := __e.Get(1)
_ = C3851
tmp14815 := PrimEqual(C3851, False)

if True == tmp14815 {
tmp14708 := MakeNative(func(__e *ControlFlow) {
C3854 := __e.Get(1)
_ = C3854
tmp14799 := PrimEqual(C3854, False)

if True == tmp14799 {
tmp14709 := MakeNative(func(__e *ControlFlow) {
C3857 := __e.Get(1)
_ = C3857
tmp14783 := PrimEqual(C3857, False)

if True == tmp14783 {
tmp14710 := MakeNative(func(__e *ControlFlow) {
C3860 := __e.Get(1)
_ = C3860
tmp14767 := PrimEqual(C3860, False)

if True == tmp14767 {
tmp14765 := Call(__e, PrimFunc(symshen_4unlocked_2), V4324)


if True == tmp14765 {
tmp14711 := MakeNative(func(__e *ControlFlow) {
Tm3863 := __e.Get(1)
_ = Tm3863
tmp14762 := PrimEqual(Tm3863, Nil)

if True == tmp14762 {
tmp14712 := MakeNative(func(__e *ControlFlow) {
Tm3864 := __e.Get(1)
_ = Tm3864
tmp14713 := MakeNative(func(__e *ControlFlow) {
GoTo3865 := __e.Get(1)
_ = GoTo3865
tmp14757 := PrimIsPair(Tm3864)

if True == tmp14757 {
tmp14714 := MakeNative(func(__e *ControlFlow) {
Tm3866 := __e.Get(1)
_ = Tm3866
tmp14715 := MakeNative(func(__e *ControlFlow) {
GoTo3867 := __e.Get(1)
_ = GoTo3867
tmp14719 := PrimEqual(Tm3866, symlist)

if True == tmp14719 {
__e.TailApply(PrimFunc(symthaw), GoTo3867)
return
} else {
tmp14717 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3866)


if True == tmp14717 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3866, symlist, V4323, GoTo3867)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14720 := MakeNative(func(__e *ControlFlow) {
tmp14721 := MakeNative(func(__e *ControlFlow) {
Tm3868 := __e.Get(1)
_ = Tm3868
tmp14722 := MakeNative(func(__e *ControlFlow) {
GoTo3869 := __e.Get(1)
_ = GoTo3869
tmp14742 := PrimIsPair(Tm3868)

if True == tmp14742 {
tmp14723 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14724 := MakeNative(func(__e *ControlFlow) {
Tm3870 := __e.Get(1)
_ = Tm3870
tmp14725 := MakeNative(func(__e *ControlFlow) {
GoTo3871 := __e.Get(1)
_ = GoTo3871
tmp14729 := PrimEqual(Tm3870, Nil)

if True == tmp14729 {
__e.TailApply(PrimFunc(symthaw), GoTo3871)
return
} else {
tmp14727 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3870)


if True == tmp14727 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3870, Nil, V4323, GoTo3871)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14730 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3869, A)
return
}, 0)

__e.TailApply(tmp14725, tmp14730)
return


}, 1)

tmp14731 := PrimTail(Tm3868)

tmp14732 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14731, V4323)


__e.TailApply(tmp14724, tmp14732)
return


}, 1)

tmp14733 := PrimHead(Tm3868)

__e.TailApply(tmp14723, tmp14733)
return


} else {
tmp14740 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3868)


if True == tmp14740 {
tmp14734 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14735 := PrimCons(A, Nil)

tmp14736 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3869, A)
return
}, 0)

tmp14737 := Call(__e, PrimFunc(symshen_4bind_b), Tm3868, tmp14735, V4323, tmp14736)


__e.TailApply(PrimFunc(symshen_4gc), V4323, tmp14737)
return


}, 1)

tmp14738 := Call(__e, PrimFunc(symshen_4newpv), V4323)


__e.TailApply(tmp14734, tmp14738)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14743 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
__e.TailApply(GoTo3865, A)
return
}, 1)

__e.TailApply(tmp14722, tmp14743)
return


}, 1)

tmp14744 := PrimTail(Tm3864)

tmp14745 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14744, V4323)


__e.TailApply(tmp14721, tmp14745)
return


}, 0)

__e.TailApply(tmp14715, tmp14720)
return


}, 1)

tmp14746 := PrimHead(Tm3864)

tmp14747 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14746, V4323)


__e.TailApply(tmp14714, tmp14747)
return


} else {
tmp14755 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3864)


if True == tmp14755 {
tmp14748 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14749 := PrimCons(A, Nil)

tmp14750 := PrimCons(symlist, tmp14749)

tmp14751 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(GoTo3865, A)
return
}, 0)

tmp14752 := Call(__e, PrimFunc(symshen_4bind_b), Tm3864, tmp14750, V4323, tmp14751)


__e.TailApply(PrimFunc(symshen_4gc), V4323, tmp14752)
return


}, 1)

tmp14753 := Call(__e, PrimFunc(symshen_4newpv), V4323)


__e.TailApply(tmp14748, tmp14753)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14758 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp14759 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14759

__e.TailApply(PrimFunc(symthaw), V4326)
return


}, 1)

__e.TailApply(tmp14713, tmp14758)
return


}, 1)

tmp14760 := Call(__e, PrimFunc(symshen_4lazyderef), V4322, V4323)


__e.TailApply(tmp14712, tmp14760)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14763 := Call(__e, PrimFunc(symshen_4lazyderef), V4321, V4323)


__e.TailApply(tmp14711, tmp14763)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C3860)
return
}


}, 1)

tmp14781 := Call(__e, PrimFunc(symshen_4unlocked_2), V4324)


var ifres14768 Obj

if True == tmp14781 {
tmp14769 := MakeNative(func(__e *ControlFlow) {
Tm3861 := __e.Get(1)
_ = Tm3861
tmp14770 := MakeNative(func(__e *ControlFlow) {
GoTo3862 := __e.Get(1)
_ = GoTo3862
tmp14774 := PrimEqual(Tm3861, symsymbol)

if True == tmp14774 {
__e.TailApply(PrimFunc(symthaw), GoTo3862)
return
} else {
tmp14772 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3861)


if True == tmp14772 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3861, symsymbol, V4323, GoTo3862)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14775 := MakeNative(func(__e *ControlFlow) {
tmp14776 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14776

tmp14777 := Call(__e, PrimFunc(symshen_4lazyderef), V4321, V4323)


tmp14778 := PrimIsSymbol(tmp14777)

__e.TailApply(PrimFunc(symwhen), tmp14778, V4323, V4324, V4325, V4326)
return


}, 0)

__e.TailApply(tmp14770, tmp14775)
return


}, 1)

tmp14779 := Call(__e, PrimFunc(symshen_4lazyderef), V4322, V4323)


tmp14780 := Call(__e, tmp14769, tmp14779)


ifres14768 = tmp14780


} else {
ifres14768 = False


}

__e.TailApply(tmp14710, ifres14768)
return


} else {
__e.Return(C3857)
return
}


}, 1)

tmp14797 := Call(__e, PrimFunc(symshen_4unlocked_2), V4324)


var ifres14784 Obj

if True == tmp14797 {
tmp14785 := MakeNative(func(__e *ControlFlow) {
Tm3858 := __e.Get(1)
_ = Tm3858
tmp14786 := MakeNative(func(__e *ControlFlow) {
GoTo3859 := __e.Get(1)
_ = GoTo3859
tmp14790 := PrimEqual(Tm3858, symstring)

if True == tmp14790 {
__e.TailApply(PrimFunc(symthaw), GoTo3859)
return
} else {
tmp14788 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3858)


if True == tmp14788 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3858, symstring, V4323, GoTo3859)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14791 := MakeNative(func(__e *ControlFlow) {
tmp14792 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14792

tmp14793 := Call(__e, PrimFunc(symshen_4lazyderef), V4321, V4323)


tmp14794 := PrimIsString(tmp14793)

__e.TailApply(PrimFunc(symwhen), tmp14794, V4323, V4324, V4325, V4326)
return


}, 0)

__e.TailApply(tmp14786, tmp14791)
return


}, 1)

tmp14795 := Call(__e, PrimFunc(symshen_4lazyderef), V4322, V4323)


tmp14796 := Call(__e, tmp14785, tmp14795)


ifres14784 = tmp14796


} else {
ifres14784 = False


}

__e.TailApply(tmp14709, ifres14784)
return


} else {
__e.Return(C3854)
return
}


}, 1)

tmp14813 := Call(__e, PrimFunc(symshen_4unlocked_2), V4324)


var ifres14800 Obj

if True == tmp14813 {
tmp14801 := MakeNative(func(__e *ControlFlow) {
Tm3855 := __e.Get(1)
_ = Tm3855
tmp14802 := MakeNative(func(__e *ControlFlow) {
GoTo3856 := __e.Get(1)
_ = GoTo3856
tmp14806 := PrimEqual(Tm3855, symboolean)

if True == tmp14806 {
__e.TailApply(PrimFunc(symthaw), GoTo3856)
return
} else {
tmp14804 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3855)


if True == tmp14804 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3855, symboolean, V4323, GoTo3856)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14807 := MakeNative(func(__e *ControlFlow) {
tmp14808 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14808

tmp14809 := Call(__e, PrimFunc(symshen_4lazyderef), V4321, V4323)


tmp14810 := Call(__e, PrimFunc(symboolean_2), tmp14809)


__e.TailApply(PrimFunc(symwhen), tmp14810, V4323, V4324, V4325, V4326)
return


}, 0)

__e.TailApply(tmp14802, tmp14807)
return


}, 1)

tmp14811 := Call(__e, PrimFunc(symshen_4lazyderef), V4322, V4323)


tmp14812 := Call(__e, tmp14801, tmp14811)


ifres14800 = tmp14812


} else {
ifres14800 = False


}

__e.TailApply(tmp14708, ifres14800)
return


} else {
__e.Return(C3851)
return
}


}, 1)

tmp14829 := Call(__e, PrimFunc(symshen_4unlocked_2), V4324)


var ifres14816 Obj

if True == tmp14829 {
tmp14817 := MakeNative(func(__e *ControlFlow) {
Tm3852 := __e.Get(1)
_ = Tm3852
tmp14818 := MakeNative(func(__e *ControlFlow) {
GoTo3853 := __e.Get(1)
_ = GoTo3853
tmp14822 := PrimEqual(Tm3852, symnumber)

if True == tmp14822 {
__e.TailApply(PrimFunc(symthaw), GoTo3853)
return
} else {
tmp14820 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3852)


if True == tmp14820 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm3852, symnumber, V4323, GoTo3853)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14823 := MakeNative(func(__e *ControlFlow) {
tmp14824 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14824

tmp14825 := Call(__e, PrimFunc(symshen_4lazyderef), V4321, V4323)


tmp14826 := PrimIsNumber(tmp14825)

__e.TailApply(PrimFunc(symwhen), tmp14826, V4323, V4324, V4325, V4326)
return


}, 0)

__e.TailApply(tmp14818, tmp14823)
return


}, 1)

tmp14827 := Call(__e, PrimFunc(symshen_4lazyderef), V4322, V4323)


tmp14828 := Call(__e, tmp14817, tmp14827)


ifres14816 = tmp14828


} else {
ifres14816 = False


}

__e.TailApply(tmp14707, ifres14816)
return


}, 6)

tmp14830 := Call(__e, ns2_1set, symshen_4primitive, tmp14706)


_ = tmp14830

tmp14831 := MakeNative(func(__e *ControlFlow) {
V4327 := __e.Get(1)
_ = V4327
V4328 := __e.Get(2)
_ = V4328
V4329 := __e.Get(3)
_ = V4329
V4330 := __e.Get(4)
_ = V4330
V4331 := __e.Get(5)
_ = V4331
V4332 := __e.Get(6)
_ = V4332
V4333 := __e.Get(7)
_ = V4333
tmp14832 := MakeNative(func(__e *ControlFlow) {
C3879 := __e.Get(1)
_ = C3879
tmp14843 := PrimEqual(C3879, False)

if True == tmp14843 {
tmp14841 := Call(__e, PrimFunc(symshen_4unlocked_2), V4331)


if True == tmp14841 {
tmp14833 := MakeNative(func(__e *ControlFlow) {
Tm3885 := __e.Get(1)
_ = Tm3885
tmp14838 := PrimIsPair(Tm3885)

if True == tmp14838 {
tmp14834 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp14835 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14835

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V4327, V4328, Hyp, V4330, V4331, V4332, V4333)
return


}, 1)

tmp14836 := PrimTail(Tm3885)

__e.TailApply(tmp14834, tmp14836)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14839 := Call(__e, PrimFunc(symshen_4lazyderef), V4329, V4330)


__e.TailApply(tmp14833, tmp14839)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C3879)
return
}


}, 1)

tmp14885 := Call(__e, PrimFunc(symshen_4unlocked_2), V4331)


var ifres14844 Obj

if True == tmp14885 {
tmp14845 := MakeNative(func(__e *ControlFlow) {
Tm3880 := __e.Get(1)
_ = Tm3880
tmp14882 := PrimIsPair(Tm3880)

if True == tmp14882 {
tmp14846 := MakeNative(func(__e *ControlFlow) {
Tm3881 := __e.Get(1)
_ = Tm3881
tmp14878 := PrimIsPair(Tm3881)

if True == tmp14878 {
tmp14847 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp14848 := MakeNative(func(__e *ControlFlow) {
Tm3882 := __e.Get(1)
_ = Tm3882
tmp14873 := PrimIsPair(Tm3882)

if True == tmp14873 {
tmp14849 := MakeNative(func(__e *ControlFlow) {
Colon := __e.Get(1)
_ = Colon
tmp14850 := MakeNative(func(__e *ControlFlow) {
Tm3883 := __e.Get(1)
_ = Tm3883
tmp14868 := PrimIsPair(Tm3883)

if True == tmp14868 {
tmp14851 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp14852 := MakeNative(func(__e *ControlFlow) {
Tm3884 := __e.Get(1)
_ = Tm3884
tmp14863 := PrimEqual(Tm3884, Nil)

if True == tmp14863 {
tmp14853 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14853

tmp14854 := Call(__e, PrimFunc(symshen_4deref), Colon, V4330)


tmp14855 := PrimIntern(MakeString(":"))

tmp14856 := PrimEqual(tmp14854, tmp14855)

tmp14857 := MakeNative(func(__e *ControlFlow) {
tmp14858 := Call(__e, PrimFunc(symshen_4deref), V4327, V4330)


tmp14859 := Call(__e, PrimFunc(symshen_4deref), Y, V4330)


tmp14860 := PrimEqual(tmp14858, tmp14859)

tmp14861 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V4328, B, V4330, V4331, V4332, V4333)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14860, V4330, V4331, V4332, tmp14861)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14856, V4330, V4331, V4332, tmp14857)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14864 := PrimTail(Tm3883)

tmp14865 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14864, V4330)


__e.TailApply(tmp14852, tmp14865)
return


}, 1)

tmp14866 := PrimHead(Tm3883)

__e.TailApply(tmp14851, tmp14866)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14869 := PrimTail(Tm3882)

tmp14870 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14869, V4330)


__e.TailApply(tmp14850, tmp14870)
return


}, 1)

tmp14871 := PrimHead(Tm3882)

__e.TailApply(tmp14849, tmp14871)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14874 := PrimTail(Tm3881)

tmp14875 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14874, V4330)


__e.TailApply(tmp14848, tmp14875)
return


}, 1)

tmp14876 := PrimHead(Tm3881)

__e.TailApply(tmp14847, tmp14876)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14879 := PrimHead(Tm3880)

tmp14880 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14879, V4330)


__e.TailApply(tmp14846, tmp14880)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14883 := Call(__e, PrimFunc(symshen_4lazyderef), V4329, V4330)


tmp14884 := Call(__e, tmp14845, tmp14883)


ifres14844 = tmp14884


} else {
ifres14844 = False


}

__e.TailApply(tmp14832, ifres14844)
return


}, 7)

tmp14886 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp14831)


_ = tmp14886

tmp14887 := MakeNative(func(__e *ControlFlow) {
V4334 := __e.Get(1)
_ = V4334
V4335 := __e.Get(2)
_ = V4335
V4336 := __e.Get(3)
_ = V4336
V4337 := __e.Get(4)
_ = V4337
V4338 := __e.Get(5)
_ = V4338
V4339 := __e.Get(6)
_ = V4339
tmp14892 := Call(__e, PrimFunc(symshen_4unlocked_2), V4337)


if True == tmp14892 {
tmp14888 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14888

tmp14889 := PrimValue(symshen_4_dsigf_d)

tmp14890 := Call(__e, PrimFunc(symassoc), V4334, tmp14889)


__e.TailApply(PrimFunc(symshen_4sigf), tmp14890, V4335, V4336, V4337, V4338, V4339)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp14893 := Call(__e, ns2_1set, symshen_4lookupsig, tmp14887)


_ = tmp14893

tmp14894 := MakeNative(func(__e *ControlFlow) {
V4354 := __e.Get(1)
_ = V4354
V4355 := __e.Get(2)
_ = V4355
V4356 := __e.Get(3)
_ = V4356
V4357 := __e.Get(4)
_ = V4357
V4358 := __e.Get(5)
_ = V4358
V4359 := __e.Get(6)
_ = V4359
tmp14901 := PrimIsPair(V4354)

if True == tmp14901 {
tmp14895 := PrimTail(V4354)

tmp14896 := Call(__e, tmp14895, V4355)


tmp14897 := Call(__e, tmp14896, V4356)


tmp14898 := Call(__e, tmp14897, V4357)


tmp14899 := Call(__e, tmp14898, V4358)


__e.TailApply(tmp14899, V4359)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp14902 := Call(__e, ns2_1set, symshen_4sigf, tmp14894)


_ = tmp14902

tmp14903 := MakeNative(func(__e *ControlFlow) {
V4360 := __e.Get(1)
_ = V4360
tmp14904 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp14905 := MakeNative(func(__e *ControlFlow) {
V0 := __e.Get(1)
_ = V0
tmp14906 := MakeNative(func(__e *ControlFlow) {
V1 := __e.Get(1)
_ = V1
tmp14907 := MakeNative(func(__e *ControlFlow) {
V2 := __e.Get(1)
_ = V2
__e.Return(V2)
return
}, 1)

tmp14908 := PrimValue(symshen_4_dgensym_d)

tmp14909 := PrimNumberAdd(MakeNumber(1), tmp14908)

tmp14910 := PrimSet(symshen_4_dgensym_d, tmp14909)

tmp14911 := PrimVectorSet(V1, MakeNumber(2), tmp14910)

__e.TailApply(tmp14907, tmp14911)
return


}, 1)

tmp14912 := PrimVectorSet(V0, MakeNumber(1), V4360)

__e.TailApply(tmp14906, tmp14912)
return


}, 1)

tmp14913 := PrimVectorSet(V, MakeNumber(0), symshen_4print_1freshterm)

__e.TailApply(tmp14905, tmp14913)
return


}, 1)

tmp14914 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp14904, tmp14914)
return


}, 1)

tmp14915 := Call(__e, ns2_1set, symshen_4freshterm, tmp14903)


_ = tmp14915

tmp14916 := MakeNative(func(__e *ControlFlow) {
V4361 := __e.Get(1)
_ = V4361
tmp14917 := PrimVectorGet(V4361, MakeNumber(1))

tmp14918 := PrimStr(tmp14917)

__e.Return(PrimStringConcat(MakeString("&&"), tmp14918))
return


}, 1)

tmp14919 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp14916)


_ = tmp14919

tmp14920 := MakeNative(func(__e *ControlFlow) {
V4362 := __e.Get(1)
_ = V4362
V4363 := __e.Get(2)
_ = V4363
V4364 := __e.Get(3)
_ = V4364
V4365 := __e.Get(4)
_ = V4365
V4366 := __e.Get(5)
_ = V4366
V4367 := __e.Get(6)
_ = V4367
V4368 := __e.Get(7)
_ = V4368
tmp14921 := MakeNative(func(__e *ControlFlow) {
C3899 := __e.Get(1)
_ = C3899
tmp14932 := PrimEqual(C3899, False)

if True == tmp14932 {
tmp14930 := Call(__e, PrimFunc(symshen_4unlocked_2), V4366)


if True == tmp14930 {
tmp14922 := MakeNative(func(__e *ControlFlow) {
Tm3902 := __e.Get(1)
_ = Tm3902
tmp14927 := PrimIsPair(Tm3902)

if True == tmp14927 {
tmp14923 := MakeNative(func(__e *ControlFlow) {
Ds := __e.Get(1)
_ = Ds
tmp14924 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14924

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V4362, V4363, Ds, V4365, V4366, V4367, V4368)
return


}, 1)

tmp14925 := PrimTail(Tm3902)

__e.TailApply(tmp14923, tmp14925)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14928 := Call(__e, PrimFunc(symshen_4lazyderef), V4364, V4365)


__e.TailApply(tmp14922, tmp14928)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C3899)
return
}


}, 1)

tmp14952 := Call(__e, PrimFunc(symshen_4unlocked_2), V4366)


var ifres14933 Obj

if True == tmp14952 {
tmp14934 := MakeNative(func(__e *ControlFlow) {
Tm3900 := __e.Get(1)
_ = Tm3900
tmp14949 := PrimIsPair(Tm3900)

if True == tmp14949 {
tmp14935 := MakeNative(func(__e *ControlFlow) {
Tm3901 := __e.Get(1)
_ = Tm3901
tmp14945 := PrimIsPair(Tm3901)

if True == tmp14945 {
tmp14936 := MakeNative(func(__e *ControlFlow) {
Fn := __e.Get(1)
_ = Fn
tmp14937 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14937

tmp14938 := Call(__e, PrimFunc(symshen_4deref), Fn, V4365)


tmp14939 := Call(__e, PrimFunc(symshen_4deref), V4362, V4365)


tmp14940 := Call(__e, tmp14938, tmp14939)


tmp14941 := Call(__e, PrimFunc(symshen_4deref), V4363, V4365)


tmp14942 := Call(__e, tmp14940, tmp14941)


__e.TailApply(PrimFunc(symcall), tmp14942, V4365, V4366, V4367, V4368)
return


}, 1)

tmp14943 := PrimTail(Tm3901)

__e.TailApply(tmp14936, tmp14943)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14946 := PrimHead(Tm3900)

tmp14947 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14946, V4365)


__e.TailApply(tmp14935, tmp14947)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14950 := Call(__e, PrimFunc(symshen_4lazyderef), V4364, V4365)


tmp14951 := Call(__e, tmp14934, tmp14950)


ifres14933 = tmp14951


} else {
ifres14933 = False


}

__e.TailApply(tmp14921, ifres14933)
return


}, 7)

tmp14953 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp14920)


_ = tmp14953

tmp14954 := MakeNative(func(__e *ControlFlow) {
V4369 := __e.Get(1)
_ = V4369
V4370 := __e.Get(2)
_ = V4370
V4371 := __e.Get(3)
_ = V4371
V4372 := __e.Get(4)
_ = V4372
V4373 := __e.Get(5)
_ = V4373
V4374 := __e.Get(6)
_ = V4374
V4375 := __e.Get(7)
_ = V4375
tmp14955 := MakeNative(func(__e *ControlFlow) {
K3905 := __e.Get(1)
_ = K3905
tmp14956 := MakeNative(func(__e *ControlFlow) {
C3910 := __e.Get(1)
_ = C3910
tmp15386 := PrimEqual(C3910, False)

if True == tmp15386 {
tmp14957 := MakeNative(func(__e *ControlFlow) {
C3913 := __e.Get(1)
_ = C3913
tmp15286 := PrimEqual(C3913, False)

if True == tmp15286 {
tmp14958 := MakeNative(func(__e *ControlFlow) {
C3928 := __e.Get(1)
_ = C3928
tmp15181 := PrimEqual(C3928, False)

if True == tmp15181 {
tmp14959 := MakeNative(func(__e *ControlFlow) {
C3944 := __e.Get(1)
_ = C3944
tmp15100 := PrimEqual(C3944, False)

if True == tmp15100 {
tmp14960 := MakeNative(func(__e *ControlFlow) {
C3956 := __e.Get(1)
_ = C3956
tmp15000 := PrimEqual(C3956, False)

if True == tmp15000 {
tmp14961 := MakeNative(func(__e *ControlFlow) {
C3971 := __e.Get(1)
_ = C3971
tmp14963 := PrimEqual(C3971, False)

if True == tmp14963 {
__e.TailApply(PrimFunc(symshen_4unlock), V4373, K3905)
return
} else {
__e.Return(C3971)
return
}


}, 1)

tmp14998 := Call(__e, PrimFunc(symshen_4unlocked_2), V4373)


var ifres14964 Obj

if True == tmp14998 {
tmp14965 := MakeNative(func(__e *ControlFlow) {
Tm3972 := __e.Get(1)
_ = Tm3972
tmp14995 := PrimIsPair(Tm3972)

if True == tmp14995 {
tmp14966 := MakeNative(func(__e *ControlFlow) {
P := __e.Get(1)
_ = P
tmp14967 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp14968 := MakeNative(func(__e *ControlFlow) {
Tm3973 := __e.Get(1)
_ = Tm3973
tmp14969 := MakeNative(func(__e *ControlFlow) {
GoTo3974 := __e.Get(1)
_ = GoTo3974
tmp14987 := PrimIsPair(Tm3973)

if True == tmp14987 {
tmp14970 := MakeNative(func(__e *ControlFlow) {
Q := __e.Get(1)
_ = Q
tmp14971 := MakeNative(func(__e *ControlFlow) {
Normalised := __e.Get(1)
_ = Normalised
tmp14972 := Call(__e, GoTo3974, Q)


__e.TailApply(tmp14972, Normalised)
return


}, 1)

tmp14973 := PrimTail(Tm3973)

__e.TailApply(tmp14971, tmp14973)
return


}, 1)

tmp14974 := PrimHead(Tm3973)

__e.TailApply(tmp14970, tmp14974)
return


} else {
tmp14985 := Call(__e, PrimFunc(symshen_4pvar_2), Tm3973)


if True == tmp14985 {
tmp14975 := MakeNative(func(__e *ControlFlow) {
Q := __e.Get(1)
_ = Q
tmp14976 := MakeNative(func(__e *ControlFlow) {
Normalised := __e.Get(1)
_ = Normalised
tmp14977 := PrimCons(Q, Normalised)

tmp14978 := MakeNative(func(__e *ControlFlow) {
tmp14979 := Call(__e, GoTo3974, Q)


__e.TailApply(tmp14979, Normalised)
return


}, 0)

tmp14980 := Call(__e, PrimFunc(symshen_4bind_b), Tm3973, tmp14977, V4372, tmp14978)


__e.TailApply(PrimFunc(symshen_4gc), V4372, tmp14980)
return


}, 1)

tmp14981 := Call(__e, PrimFunc(symshen_4newpv), V4372)


tmp14982 := Call(__e, tmp14976, tmp14981)


__e.TailApply(PrimFunc(symshen_4gc), V4372, tmp14982)
return


}, 1)

tmp14983 := Call(__e, PrimFunc(symshen_4newpv), V4372)


__e.TailApply(tmp14975, tmp14983)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14988 := MakeNative(func(__e *ControlFlow) {
Q := __e.Get(1)
_ = Q
__e.Return(MakeNative(func(__e *ControlFlow) {
Normalised := __e.Get(1)
_ = Normalised
tmp14989 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14989

tmp14990 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), Hyp, Normalised, V4371, V4372, V4373, K3905, V4375)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Q, P, V4372, V4373, K3905, tmp14990)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14969, tmp14988)
return


}, 1)

tmp14991 := Call(__e, PrimFunc(symshen_4lazyderef), V4370, V4372)


__e.TailApply(tmp14968, tmp14991)
return


}, 1)

tmp14992 := PrimTail(Tm3972)

__e.TailApply(tmp14967, tmp14992)
return


}, 1)

tmp14993 := PrimHead(Tm3972)

__e.TailApply(tmp14966, tmp14993)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14996 := Call(__e, PrimFunc(symshen_4lazyderef), V4369, V4372)


tmp14997 := Call(__e, tmp14965, tmp14996)


ifres14964 = tmp14997


} else {
ifres14964 = False


}

__e.TailApply(tmp14961, ifres14964)
return


} else {
__e.Return(C3956)
return
}


}, 1)

tmp15098 := Call(__e, PrimFunc(symshen_4unlocked_2), V4373)


var ifres15001 Obj

if True == tmp15098 {
tmp15002 := MakeNative(func(__e *ControlFlow) {
Tm3957 := __e.Get(1)
_ = Tm3957
tmp15095 := PrimIsPair(Tm3957)

if True == tmp15095 {
tmp15003 := MakeNative(func(__e *ControlFlow) {
Tm3958 := __e.Get(1)
_ = Tm3958
tmp15091 := PrimIsPair(Tm3958)

if True == tmp15091 {
tmp15004 := MakeNative(func(__e *ControlFlow) {
Tm3959 := __e.Get(1)
_ = Tm3959
tmp15087 := PrimIsPair(Tm3959)

if True == tmp15087 {
tmp15005 := MakeNative(func(__e *ControlFlow) {
Tm3960 := __e.Get(1)
_ = Tm3960
tmp15083 := PrimEqual(Tm3960, sym_8v)

if True == tmp15083 {
tmp15006 := MakeNative(func(__e *ControlFlow) {
Tm3961 := __e.Get(1)
_ = Tm3961
tmp15079 := PrimIsPair(Tm3961)

if True == tmp15079 {
tmp15007 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp15008 := MakeNative(func(__e *ControlFlow) {
Tm3962 := __e.Get(1)
_ = Tm3962
tmp15074 := PrimIsPair(Tm3962)

if True == tmp15074 {
tmp15009 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp15010 := MakeNative(func(__e *ControlFlow) {
Tm3963 := __e.Get(1)
_ = Tm3963
tmp15069 := PrimEqual(Tm3963, Nil)

if True == tmp15069 {
tmp15011 := MakeNative(func(__e *ControlFlow) {
Tm3964 := __e.Get(1)
_ = Tm3964
tmp15065 := PrimIsPair(Tm3964)

if True == tmp15065 {
tmp15012 := MakeNative(func(__e *ControlFlow) {
Colon := __e.Get(1)
_ = Colon
tmp15013 := MakeNative(func(__e *ControlFlow) {
Tm3965 := __e.Get(1)
_ = Tm3965
tmp15060 := PrimIsPair(Tm3965)

if True == tmp15060 {
tmp15014 := MakeNative(func(__e *ControlFlow) {
Tm3966 := __e.Get(1)
_ = Tm3966
tmp15056 := PrimIsPair(Tm3966)

if True == tmp15056 {
tmp15015 := MakeNative(func(__e *ControlFlow) {
Tm3967 := __e.Get(1)
_ = Tm3967
tmp15052 := PrimEqual(Tm3967, symvector)

if True == tmp15052 {
tmp15016 := MakeNative(func(__e *ControlFlow) {
Tm3968 := __e.Get(1)
_ = Tm3968
tmp15048 := PrimIsPair(Tm3968)

if True == tmp15048 {
tmp15017 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp15018 := MakeNative(func(__e *ControlFlow) {
Tm3969 := __e.Get(1)
_ = Tm3969
tmp15043 := PrimEqual(Tm3969, Nil)

if True == tmp15043 {
tmp15019 := MakeNative(func(__e *ControlFlow) {
Tm3970 := __e.Get(1)
_ = Tm3970
tmp15039 := PrimEqual(Tm3970, Nil)

if True == tmp15039 {
tmp15020 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp15021 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15021

tmp15022 := Call(__e, PrimFunc(symshen_4deref), Colon, V4372)


tmp15023 := PrimIntern(MakeString(":"))

tmp15024 := PrimEqual(tmp15022, tmp15023)

tmp15025 := MakeNative(func(__e *ControlFlow) {
tmp15026 := MakeNative(func(__e *ControlFlow) {
tmp15027 := PrimCons(A, Nil)

tmp15028 := PrimCons(Colon, tmp15027)

tmp15029 := PrimCons(X, tmp15028)

tmp15030 := PrimCons(A, Nil)

tmp15031 := PrimCons(symvector, tmp15030)

tmp15032 := PrimCons(tmp15031, Nil)

tmp15033 := PrimCons(Colon, tmp15032)

tmp15034 := PrimCons(Y, tmp15033)

tmp15035 := PrimCons(tmp15034, Hyp)

tmp15036 := PrimCons(tmp15029, tmp15035)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15036, V4370, True, V4372, V4373, K3905, V4375)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4372, V4373, K3905, tmp15026)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15024, V4372, V4373, K3905, tmp15025)
return


}, 1)

tmp15037 := PrimTail(Tm3957)

__e.TailApply(tmp15020, tmp15037)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15040 := PrimTail(Tm3965)

tmp15041 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15040, V4372)


__e.TailApply(tmp15019, tmp15041)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15044 := PrimTail(Tm3968)

tmp15045 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15044, V4372)


__e.TailApply(tmp15018, tmp15045)
return


}, 1)

tmp15046 := PrimHead(Tm3968)

__e.TailApply(tmp15017, tmp15046)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15049 := PrimTail(Tm3966)

tmp15050 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15049, V4372)


__e.TailApply(tmp15016, tmp15050)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15053 := PrimHead(Tm3966)

tmp15054 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15053, V4372)


__e.TailApply(tmp15015, tmp15054)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15057 := PrimHead(Tm3965)

tmp15058 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15057, V4372)


__e.TailApply(tmp15014, tmp15058)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15061 := PrimTail(Tm3964)

tmp15062 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15061, V4372)


__e.TailApply(tmp15013, tmp15062)
return


}, 1)

tmp15063 := PrimHead(Tm3964)

__e.TailApply(tmp15012, tmp15063)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15066 := PrimTail(Tm3958)

tmp15067 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15066, V4372)


__e.TailApply(tmp15011, tmp15067)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15070 := PrimTail(Tm3962)

tmp15071 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15070, V4372)


__e.TailApply(tmp15010, tmp15071)
return


}, 1)

tmp15072 := PrimHead(Tm3962)

__e.TailApply(tmp15009, tmp15072)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15075 := PrimTail(Tm3961)

tmp15076 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15075, V4372)


__e.TailApply(tmp15008, tmp15076)
return


}, 1)

tmp15077 := PrimHead(Tm3961)

__e.TailApply(tmp15007, tmp15077)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15080 := PrimTail(Tm3959)

tmp15081 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15080, V4372)


__e.TailApply(tmp15006, tmp15081)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15084 := PrimHead(Tm3959)

tmp15085 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15084, V4372)


__e.TailApply(tmp15005, tmp15085)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15088 := PrimHead(Tm3958)

tmp15089 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15088, V4372)


__e.TailApply(tmp15004, tmp15089)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15092 := PrimHead(Tm3957)

tmp15093 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15092, V4372)


__e.TailApply(tmp15003, tmp15093)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15096 := Call(__e, PrimFunc(symshen_4lazyderef), V4369, V4372)


tmp15097 := Call(__e, tmp15002, tmp15096)


ifres15001 = tmp15097


} else {
ifres15001 = False


}

__e.TailApply(tmp14960, ifres15001)
return


} else {
__e.Return(C3944)
return
}


}, 1)

tmp15179 := Call(__e, PrimFunc(symshen_4unlocked_2), V4373)


var ifres15101 Obj

if True == tmp15179 {
tmp15102 := MakeNative(func(__e *ControlFlow) {
Tm3945 := __e.Get(1)
_ = Tm3945
tmp15176 := PrimIsPair(Tm3945)

if True == tmp15176 {
tmp15103 := MakeNative(func(__e *ControlFlow) {
Tm3946 := __e.Get(1)
_ = Tm3946
tmp15172 := PrimIsPair(Tm3946)

if True == tmp15172 {
tmp15104 := MakeNative(func(__e *ControlFlow) {
Tm3947 := __e.Get(1)
_ = Tm3947
tmp15168 := PrimIsPair(Tm3947)

if True == tmp15168 {
tmp15105 := MakeNative(func(__e *ControlFlow) {
Tm3948 := __e.Get(1)
_ = Tm3948
tmp15164 := PrimEqual(Tm3948, sym_8s)

if True == tmp15164 {
tmp15106 := MakeNative(func(__e *ControlFlow) {
Tm3949 := __e.Get(1)
_ = Tm3949
tmp15160 := PrimIsPair(Tm3949)

if True == tmp15160 {
tmp15107 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp15108 := MakeNative(func(__e *ControlFlow) {
Tm3950 := __e.Get(1)
_ = Tm3950
tmp15155 := PrimIsPair(Tm3950)

if True == tmp15155 {
tmp15109 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp15110 := MakeNative(func(__e *ControlFlow) {
Tm3951 := __e.Get(1)
_ = Tm3951
tmp15150 := PrimEqual(Tm3951, Nil)

if True == tmp15150 {
tmp15111 := MakeNative(func(__e *ControlFlow) {
Tm3952 := __e.Get(1)
_ = Tm3952
tmp15146 := PrimIsPair(Tm3952)

if True == tmp15146 {
tmp15112 := MakeNative(func(__e *ControlFlow) {
Colon := __e.Get(1)
_ = Colon
tmp15113 := MakeNative(func(__e *ControlFlow) {
Tm3953 := __e.Get(1)
_ = Tm3953
tmp15141 := PrimIsPair(Tm3953)

if True == tmp15141 {
tmp15114 := MakeNative(func(__e *ControlFlow) {
Tm3954 := __e.Get(1)
_ = Tm3954
tmp15137 := PrimEqual(Tm3954, symstring)

if True == tmp15137 {
tmp15115 := MakeNative(func(__e *ControlFlow) {
Tm3955 := __e.Get(1)
_ = Tm3955
tmp15133 := PrimEqual(Tm3955, Nil)

if True == tmp15133 {
tmp15116 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp15117 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15117

tmp15118 := Call(__e, PrimFunc(symshen_4deref), Colon, V4372)


tmp15119 := PrimIntern(MakeString(":"))

tmp15120 := PrimEqual(tmp15118, tmp15119)

tmp15121 := MakeNative(func(__e *ControlFlow) {
tmp15122 := MakeNative(func(__e *ControlFlow) {
tmp15123 := PrimCons(symstring, Nil)

tmp15124 := PrimCons(Colon, tmp15123)

tmp15125 := PrimCons(X, tmp15124)

tmp15126 := PrimCons(symstring, Nil)

tmp15127 := PrimCons(Colon, tmp15126)

tmp15128 := PrimCons(Y, tmp15127)

tmp15129 := PrimCons(tmp15128, Hyp)

tmp15130 := PrimCons(tmp15125, tmp15129)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15130, V4370, True, V4372, V4373, K3905, V4375)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4372, V4373, K3905, tmp15122)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15120, V4372, V4373, K3905, tmp15121)
return


}, 1)

tmp15131 := PrimTail(Tm3945)

__e.TailApply(tmp15116, tmp15131)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15134 := PrimTail(Tm3953)

tmp15135 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15134, V4372)


__e.TailApply(tmp15115, tmp15135)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15138 := PrimHead(Tm3953)

tmp15139 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15138, V4372)


__e.TailApply(tmp15114, tmp15139)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15142 := PrimTail(Tm3952)

tmp15143 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15142, V4372)


__e.TailApply(tmp15113, tmp15143)
return


}, 1)

tmp15144 := PrimHead(Tm3952)

__e.TailApply(tmp15112, tmp15144)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15147 := PrimTail(Tm3946)

tmp15148 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15147, V4372)


__e.TailApply(tmp15111, tmp15148)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15151 := PrimTail(Tm3950)

tmp15152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15151, V4372)


__e.TailApply(tmp15110, tmp15152)
return


}, 1)

tmp15153 := PrimHead(Tm3950)

__e.TailApply(tmp15109, tmp15153)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15156 := PrimTail(Tm3949)

tmp15157 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15156, V4372)


__e.TailApply(tmp15108, tmp15157)
return


}, 1)

tmp15158 := PrimHead(Tm3949)

__e.TailApply(tmp15107, tmp15158)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15161 := PrimTail(Tm3947)

tmp15162 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15161, V4372)


__e.TailApply(tmp15106, tmp15162)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15165 := PrimHead(Tm3947)

tmp15166 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15165, V4372)


__e.TailApply(tmp15105, tmp15166)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15169 := PrimHead(Tm3946)

tmp15170 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15169, V4372)


__e.TailApply(tmp15104, tmp15170)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15173 := PrimHead(Tm3945)

tmp15174 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15173, V4372)


__e.TailApply(tmp15103, tmp15174)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15177 := Call(__e, PrimFunc(symshen_4lazyderef), V4369, V4372)


tmp15178 := Call(__e, tmp15102, tmp15177)


ifres15101 = tmp15178


} else {
ifres15101 = False


}

__e.TailApply(tmp14959, ifres15101)
return


} else {
__e.Return(C3928)
return
}


}, 1)

tmp15284 := Call(__e, PrimFunc(symshen_4unlocked_2), V4373)


var ifres15182 Obj

if True == tmp15284 {
tmp15183 := MakeNative(func(__e *ControlFlow) {
Tm3929 := __e.Get(1)
_ = Tm3929
tmp15281 := PrimIsPair(Tm3929)

if True == tmp15281 {
tmp15184 := MakeNative(func(__e *ControlFlow) {
Tm3930 := __e.Get(1)
_ = Tm3930
tmp15277 := PrimIsPair(Tm3930)

if True == tmp15277 {
tmp15185 := MakeNative(func(__e *ControlFlow) {
Tm3931 := __e.Get(1)
_ = Tm3931
tmp15273 := PrimIsPair(Tm3931)

if True == tmp15273 {
tmp15186 := MakeNative(func(__e *ControlFlow) {
Tm3932 := __e.Get(1)
_ = Tm3932
tmp15269 := PrimEqual(Tm3932, sym_8p)

if True == tmp15269 {
tmp15187 := MakeNative(func(__e *ControlFlow) {
Tm3933 := __e.Get(1)
_ = Tm3933
tmp15265 := PrimIsPair(Tm3933)

if True == tmp15265 {
tmp15188 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp15189 := MakeNative(func(__e *ControlFlow) {
Tm3934 := __e.Get(1)
_ = Tm3934
tmp15260 := PrimIsPair(Tm3934)

if True == tmp15260 {
tmp15190 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp15191 := MakeNative(func(__e *ControlFlow) {
Tm3935 := __e.Get(1)
_ = Tm3935
tmp15255 := PrimEqual(Tm3935, Nil)

if True == tmp15255 {
tmp15192 := MakeNative(func(__e *ControlFlow) {
Tm3936 := __e.Get(1)
_ = Tm3936
tmp15251 := PrimIsPair(Tm3936)

if True == tmp15251 {
tmp15193 := MakeNative(func(__e *ControlFlow) {
Colon := __e.Get(1)
_ = Colon
tmp15194 := MakeNative(func(__e *ControlFlow) {
Tm3937 := __e.Get(1)
_ = Tm3937
tmp15246 := PrimIsPair(Tm3937)

if True == tmp15246 {
tmp15195 := MakeNative(func(__e *ControlFlow) {
Tm3938 := __e.Get(1)
_ = Tm3938
tmp15242 := PrimIsPair(Tm3938)

if True == tmp15242 {
tmp15196 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp15197 := MakeNative(func(__e *ControlFlow) {
Tm3939 := __e.Get(1)
_ = Tm3939
tmp15237 := PrimIsPair(Tm3939)

if True == tmp15237 {
tmp15198 := MakeNative(func(__e *ControlFlow) {
Tm3940 := __e.Get(1)
_ = Tm3940
tmp15233 := PrimEqual(Tm3940, sym_d)

if True == tmp15233 {
tmp15199 := MakeNative(func(__e *ControlFlow) {
Tm3941 := __e.Get(1)
_ = Tm3941
tmp15229 := PrimIsPair(Tm3941)

if True == tmp15229 {
tmp15200 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp15201 := MakeNative(func(__e *ControlFlow) {
Tm3942 := __e.Get(1)
_ = Tm3942
tmp15224 := PrimEqual(Tm3942, Nil)

if True == tmp15224 {
tmp15202 := MakeNative(func(__e *ControlFlow) {
Tm3943 := __e.Get(1)
_ = Tm3943
tmp15220 := PrimEqual(Tm3943, Nil)

if True == tmp15220 {
tmp15203 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp15204 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15204

tmp15205 := Call(__e, PrimFunc(symshen_4deref), Colon, V4372)


tmp15206 := PrimIntern(MakeString(":"))

tmp15207 := PrimEqual(tmp15205, tmp15206)

tmp15208 := MakeNative(func(__e *ControlFlow) {
tmp15209 := MakeNative(func(__e *ControlFlow) {
tmp15210 := PrimCons(A, Nil)

tmp15211 := PrimCons(Colon, tmp15210)

tmp15212 := PrimCons(X, tmp15211)

tmp15213 := PrimCons(B, Nil)

tmp15214 := PrimCons(Colon, tmp15213)

tmp15215 := PrimCons(Y, tmp15214)

tmp15216 := PrimCons(tmp15215, Hyp)

tmp15217 := PrimCons(tmp15212, tmp15216)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15217, V4370, True, V4372, V4373, K3905, V4375)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4372, V4373, K3905, tmp15209)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15207, V4372, V4373, K3905, tmp15208)
return


}, 1)

tmp15218 := PrimTail(Tm3929)

__e.TailApply(tmp15203, tmp15218)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15221 := PrimTail(Tm3937)

tmp15222 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15221, V4372)


__e.TailApply(tmp15202, tmp15222)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15225 := PrimTail(Tm3941)

tmp15226 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15225, V4372)


__e.TailApply(tmp15201, tmp15226)
return


}, 1)

tmp15227 := PrimHead(Tm3941)

__e.TailApply(tmp15200, tmp15227)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15230 := PrimTail(Tm3939)

tmp15231 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15230, V4372)


__e.TailApply(tmp15199, tmp15231)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15234 := PrimHead(Tm3939)

tmp15235 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15234, V4372)


__e.TailApply(tmp15198, tmp15235)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15238 := PrimTail(Tm3938)

tmp15239 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15238, V4372)


__e.TailApply(tmp15197, tmp15239)
return


}, 1)

tmp15240 := PrimHead(Tm3938)

__e.TailApply(tmp15196, tmp15240)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15243 := PrimHead(Tm3937)

tmp15244 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15243, V4372)


__e.TailApply(tmp15195, tmp15244)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15247 := PrimTail(Tm3936)

tmp15248 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15247, V4372)


__e.TailApply(tmp15194, tmp15248)
return


}, 1)

tmp15249 := PrimHead(Tm3936)

__e.TailApply(tmp15193, tmp15249)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15252 := PrimTail(Tm3930)

tmp15253 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15252, V4372)


__e.TailApply(tmp15192, tmp15253)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15256 := PrimTail(Tm3934)

tmp15257 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15256, V4372)


__e.TailApply(tmp15191, tmp15257)
return


}, 1)

tmp15258 := PrimHead(Tm3934)

__e.TailApply(tmp15190, tmp15258)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15261 := PrimTail(Tm3933)

tmp15262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15261, V4372)


__e.TailApply(tmp15189, tmp15262)
return


}, 1)

tmp15263 := PrimHead(Tm3933)

__e.TailApply(tmp15188, tmp15263)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15266 := PrimTail(Tm3931)

tmp15267 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15266, V4372)


__e.TailApply(tmp15187, tmp15267)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15270 := PrimHead(Tm3931)

tmp15271 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15270, V4372)


__e.TailApply(tmp15186, tmp15271)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15274 := PrimHead(Tm3930)

tmp15275 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15274, V4372)


__e.TailApply(tmp15185, tmp15275)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15278 := PrimHead(Tm3929)

tmp15279 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15278, V4372)


__e.TailApply(tmp15184, tmp15279)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15282 := Call(__e, PrimFunc(symshen_4lazyderef), V4369, V4372)


tmp15283 := Call(__e, tmp15183, tmp15282)


ifres15182 = tmp15283


} else {
ifres15182 = False


}

__e.TailApply(tmp14958, ifres15182)
return


} else {
__e.Return(C3913)
return
}


}, 1)

tmp15384 := Call(__e, PrimFunc(symshen_4unlocked_2), V4373)


var ifres15287 Obj

if True == tmp15384 {
tmp15288 := MakeNative(func(__e *ControlFlow) {
Tm3914 := __e.Get(1)
_ = Tm3914
tmp15381 := PrimIsPair(Tm3914)

if True == tmp15381 {
tmp15289 := MakeNative(func(__e *ControlFlow) {
Tm3915 := __e.Get(1)
_ = Tm3915
tmp15377 := PrimIsPair(Tm3915)

if True == tmp15377 {
tmp15290 := MakeNative(func(__e *ControlFlow) {
Tm3916 := __e.Get(1)
_ = Tm3916
tmp15373 := PrimIsPair(Tm3916)

if True == tmp15373 {
tmp15291 := MakeNative(func(__e *ControlFlow) {
Tm3917 := __e.Get(1)
_ = Tm3917
tmp15369 := PrimEqual(Tm3917, symcons)

if True == tmp15369 {
tmp15292 := MakeNative(func(__e *ControlFlow) {
Tm3918 := __e.Get(1)
_ = Tm3918
tmp15365 := PrimIsPair(Tm3918)

if True == tmp15365 {
tmp15293 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp15294 := MakeNative(func(__e *ControlFlow) {
Tm3919 := __e.Get(1)
_ = Tm3919
tmp15360 := PrimIsPair(Tm3919)

if True == tmp15360 {
tmp15295 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp15296 := MakeNative(func(__e *ControlFlow) {
Tm3920 := __e.Get(1)
_ = Tm3920
tmp15355 := PrimEqual(Tm3920, Nil)

if True == tmp15355 {
tmp15297 := MakeNative(func(__e *ControlFlow) {
Tm3921 := __e.Get(1)
_ = Tm3921
tmp15351 := PrimIsPair(Tm3921)

if True == tmp15351 {
tmp15298 := MakeNative(func(__e *ControlFlow) {
Colon := __e.Get(1)
_ = Colon
tmp15299 := MakeNative(func(__e *ControlFlow) {
Tm3922 := __e.Get(1)
_ = Tm3922
tmp15346 := PrimIsPair(Tm3922)

if True == tmp15346 {
tmp15300 := MakeNative(func(__e *ControlFlow) {
Tm3923 := __e.Get(1)
_ = Tm3923
tmp15342 := PrimIsPair(Tm3923)

if True == tmp15342 {
tmp15301 := MakeNative(func(__e *ControlFlow) {
Tm3924 := __e.Get(1)
_ = Tm3924
tmp15338 := PrimEqual(Tm3924, symlist)

if True == tmp15338 {
tmp15302 := MakeNative(func(__e *ControlFlow) {
Tm3925 := __e.Get(1)
_ = Tm3925
tmp15334 := PrimIsPair(Tm3925)

if True == tmp15334 {
tmp15303 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp15304 := MakeNative(func(__e *ControlFlow) {
Tm3926 := __e.Get(1)
_ = Tm3926
tmp15329 := PrimEqual(Tm3926, Nil)

if True == tmp15329 {
tmp15305 := MakeNative(func(__e *ControlFlow) {
Tm3927 := __e.Get(1)
_ = Tm3927
tmp15325 := PrimEqual(Tm3927, Nil)

if True == tmp15325 {
tmp15306 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp15307 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15307

tmp15308 := Call(__e, PrimFunc(symshen_4deref), Colon, V4372)


tmp15309 := PrimIntern(MakeString(":"))

tmp15310 := PrimEqual(tmp15308, tmp15309)

tmp15311 := MakeNative(func(__e *ControlFlow) {
tmp15312 := MakeNative(func(__e *ControlFlow) {
tmp15313 := PrimCons(A, Nil)

tmp15314 := PrimCons(Colon, tmp15313)

tmp15315 := PrimCons(X, tmp15314)

tmp15316 := PrimCons(A, Nil)

tmp15317 := PrimCons(symlist, tmp15316)

tmp15318 := PrimCons(tmp15317, Nil)

tmp15319 := PrimCons(Colon, tmp15318)

tmp15320 := PrimCons(Y, tmp15319)

tmp15321 := PrimCons(tmp15320, Hyp)

tmp15322 := PrimCons(tmp15315, tmp15321)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15322, V4370, True, V4372, V4373, K3905, V4375)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4372, V4373, K3905, tmp15312)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15310, V4372, V4373, K3905, tmp15311)
return


}, 1)

tmp15323 := PrimTail(Tm3914)

__e.TailApply(tmp15306, tmp15323)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15326 := PrimTail(Tm3922)

tmp15327 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15326, V4372)


__e.TailApply(tmp15305, tmp15327)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15330 := PrimTail(Tm3925)

tmp15331 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15330, V4372)


__e.TailApply(tmp15304, tmp15331)
return


}, 1)

tmp15332 := PrimHead(Tm3925)

__e.TailApply(tmp15303, tmp15332)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15335 := PrimTail(Tm3923)

tmp15336 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15335, V4372)


__e.TailApply(tmp15302, tmp15336)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15339 := PrimHead(Tm3923)

tmp15340 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15339, V4372)


__e.TailApply(tmp15301, tmp15340)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15343 := PrimHead(Tm3922)

tmp15344 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15343, V4372)


__e.TailApply(tmp15300, tmp15344)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15347 := PrimTail(Tm3921)

tmp15348 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15347, V4372)


__e.TailApply(tmp15299, tmp15348)
return


}, 1)

tmp15349 := PrimHead(Tm3921)

__e.TailApply(tmp15298, tmp15349)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15352 := PrimTail(Tm3915)

tmp15353 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15352, V4372)


__e.TailApply(tmp15297, tmp15353)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15356 := PrimTail(Tm3919)

tmp15357 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15356, V4372)


__e.TailApply(tmp15296, tmp15357)
return


}, 1)

tmp15358 := PrimHead(Tm3919)

__e.TailApply(tmp15295, tmp15358)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15361 := PrimTail(Tm3918)

tmp15362 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15361, V4372)


__e.TailApply(tmp15294, tmp15362)
return


}, 1)

tmp15363 := PrimHead(Tm3918)

__e.TailApply(tmp15293, tmp15363)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15366 := PrimTail(Tm3916)

tmp15367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15366, V4372)


__e.TailApply(tmp15292, tmp15367)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15370 := PrimHead(Tm3916)

tmp15371 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15370, V4372)


__e.TailApply(tmp15291, tmp15371)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15374 := PrimHead(Tm3915)

tmp15375 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15374, V4372)


__e.TailApply(tmp15290, tmp15375)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15378 := PrimHead(Tm3914)

tmp15379 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15378, V4372)


__e.TailApply(tmp15289, tmp15379)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15382 := Call(__e, PrimFunc(symshen_4lazyderef), V4369, V4372)


tmp15383 := Call(__e, tmp15288, tmp15382)


ifres15287 = tmp15383


} else {
ifres15287 = False


}

__e.TailApply(tmp14957, ifres15287)
return


} else {
__e.Return(C3910)
return
}


}, 1)

tmp15399 := Call(__e, PrimFunc(symshen_4unlocked_2), V4373)


var ifres15387 Obj

if True == tmp15399 {
tmp15388 := MakeNative(func(__e *ControlFlow) {
Tm3911 := __e.Get(1)
_ = Tm3911
tmp15396 := PrimEqual(Tm3911, Nil)

if True == tmp15396 {
tmp15389 := MakeNative(func(__e *ControlFlow) {
Tm3912 := __e.Get(1)
_ = Tm3912
tmp15393 := PrimEqual(Tm3912, True)

if True == tmp15393 {
tmp15390 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15390

tmp15391 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V4370, Nil, V4372, V4373, K3905, V4375)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4372, V4373, K3905, tmp15391)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15394 := Call(__e, PrimFunc(symshen_4lazyderef), V4371, V4372)


__e.TailApply(tmp15389, tmp15394)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15397 := Call(__e, PrimFunc(symshen_4lazyderef), V4369, V4372)


tmp15398 := Call(__e, tmp15388, tmp15397)


ifres15387 = tmp15398


} else {
ifres15387 = False


}

__e.TailApply(tmp14956, ifres15387)
return


}, 1)

tmp15400 := PrimNumberAdd(V4374, MakeNumber(1))

__e.TailApply(tmp14955, tmp15400)
return


}, 7)

tmp15401 := Call(__e, ns2_1set, symshen_4l_1rules, tmp14954)


_ = tmp15401

tmp15402 := MakeNative(func(__e *ControlFlow) {
V4376 := __e.Get(1)
_ = V4376
V4377 := __e.Get(2)
_ = V4377
V4378 := __e.Get(3)
_ = V4378
V4379 := __e.Get(4)
_ = V4379
V4380 := __e.Get(5)
_ = V4380
V4381 := __e.Get(6)
_ = V4381
tmp15403 := MakeNative(func(__e *ControlFlow) {
K3977 := __e.Get(1)
_ = K3977
tmp15404 := MakeNative(func(__e *ControlFlow) {
C3981 := __e.Get(1)
_ = C3981
tmp15406 := PrimEqual(C3981, False)

if True == tmp15406 {
__e.TailApply(PrimFunc(symshen_4unlock), V4379, K3977)
return
} else {
__e.Return(C3981)
return
}


}, 1)

tmp15454 := Call(__e, PrimFunc(symshen_4unlocked_2), V4379)


var ifres15407 Obj

if True == tmp15454 {
tmp15408 := MakeNative(func(__e *ControlFlow) {
Tm3982 := __e.Get(1)
_ = Tm3982
tmp15451 := PrimIsPair(Tm3982)

if True == tmp15451 {
tmp15409 := MakeNative(func(__e *ControlFlow) {
Tm3983 := __e.Get(1)
_ = Tm3983
tmp15447 := PrimEqual(Tm3983, symdefine)

if True == tmp15447 {
tmp15410 := MakeNative(func(__e *ControlFlow) {
Tm3984 := __e.Get(1)
_ = Tm3984
tmp15443 := PrimIsPair(Tm3984)

if True == tmp15443 {
tmp15411 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp15412 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp15413 := MakeNative(func(__e *ControlFlow) {
SigxRules := __e.Get(1)
_ = SigxRules
tmp15414 := MakeNative(func(__e *ControlFlow) {
Rules := __e.Get(1)
_ = Rules
tmp15415 := MakeNative(func(__e *ControlFlow) {
FreshSig := __e.Get(1)
_ = FreshSig
tmp15416 := MakeNative(func(__e *ControlFlow) {
Sig := __e.Get(1)
_ = Sig
tmp15417 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15417

tmp15418 := MakeNative(func(__e *ControlFlow) {
tmp15419 := PrimCons(F, X)

tmp15420 := Call(__e, PrimFunc(symshen_4sigxrules), tmp15419)


tmp15421 := MakeNative(func(__e *ControlFlow) {
tmp15422 := Call(__e, PrimFunc(symshen_4lazyderef), SigxRules, V4378)


tmp15423 := Call(__e, PrimFunc(symfst), tmp15422)


tmp15424 := MakeNative(func(__e *ControlFlow) {
tmp15425 := Call(__e, PrimFunc(symshen_4lazyderef), SigxRules, V4378)


tmp15426 := Call(__e, PrimFunc(symsnd), tmp15425)


tmp15427 := MakeNative(func(__e *ControlFlow) {
tmp15428 := Call(__e, PrimFunc(symshen_4deref), Sig, V4378)


tmp15429 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp15428)


tmp15430 := MakeNative(func(__e *ControlFlow) {
tmp15431 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), Sig, V4377, V4378, V4379, K3977, V4381)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), F, Rules, FreshSig, MakeNumber(1), V4378, V4379, K3977, tmp15431)
return


}, 0)

__e.TailApply(PrimFunc(symbind), FreshSig, tmp15429, V4378, V4379, K3977, tmp15430)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Rules, tmp15426, V4378, V4379, K3977, tmp15427)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Sig, tmp15423, V4378, V4379, K3977, tmp15424)
return


}, 0)

__e.TailApply(PrimFunc(symbind), SigxRules, tmp15420, V4378, V4379, K3977, tmp15421)
return


}, 0)

tmp15432 := Call(__e, PrimFunc(symshen_4cut), V4378, V4379, K3977, tmp15418)


__e.TailApply(PrimFunc(symshen_4gc), V4378, tmp15432)
return


}, 1)

tmp15433 := Call(__e, PrimFunc(symshen_4newpv), V4378)


tmp15434 := Call(__e, tmp15416, tmp15433)


__e.TailApply(PrimFunc(symshen_4gc), V4378, tmp15434)
return


}, 1)

tmp15435 := Call(__e, PrimFunc(symshen_4newpv), V4378)


tmp15436 := Call(__e, tmp15415, tmp15435)


__e.TailApply(PrimFunc(symshen_4gc), V4378, tmp15436)
return


}, 1)

tmp15437 := Call(__e, PrimFunc(symshen_4newpv), V4378)


tmp15438 := Call(__e, tmp15414, tmp15437)


__e.TailApply(PrimFunc(symshen_4gc), V4378, tmp15438)
return


}, 1)

tmp15439 := Call(__e, PrimFunc(symshen_4newpv), V4378)


__e.TailApply(tmp15413, tmp15439)
return


}, 1)

tmp15440 := PrimTail(Tm3984)

__e.TailApply(tmp15412, tmp15440)
return


}, 1)

tmp15441 := PrimHead(Tm3984)

__e.TailApply(tmp15411, tmp15441)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15444 := PrimTail(Tm3982)

tmp15445 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15444, V4378)


__e.TailApply(tmp15410, tmp15445)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15448 := PrimHead(Tm3982)

tmp15449 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15448, V4378)


__e.TailApply(tmp15409, tmp15449)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15452 := Call(__e, PrimFunc(symshen_4lazyderef), V4376, V4378)


tmp15453 := Call(__e, tmp15408, tmp15452)


ifres15407 = tmp15453


} else {
ifres15407 = False


}

__e.TailApply(tmp15404, ifres15407)
return


}, 1)

tmp15455 := PrimNumberAdd(V4380, MakeNumber(1))

__e.TailApply(tmp15403, tmp15455)
return


}, 6)

tmp15456 := Call(__e, ns2_1set, symshen_4t_d, tmp15402)


_ = tmp15456

tmp15457 := MakeNative(func(__e *ControlFlow) {
V4382 := __e.Get(1)
_ = V4382
tmp15458 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), X)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp15458, V4382)
return


}, 1)

tmp15459 := Call(__e, ns2_1set, symshen_4sigxrules, tmp15457)


_ = tmp15459

tmp15460 := MakeNative(func(__e *ControlFlow) {
V4383 := __e.Get(1)
_ = V4383
tmp15461 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp15493 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V4383)


var ifres15464 Obj

if True == tmp15493 {
tmp15465 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp15466 := MakeNative(func(__e *ControlFlow) {
News3986 := __e.Get(1)
_ = News3986
tmp15488 := Call(__e, PrimFunc(symshen_4_ahd_2), News3986, sym_i)


if True == tmp15488 {
tmp15467 := MakeNative(func(__e *ControlFlow) {
News3987 := __e.Get(1)
_ = News3987
tmp15468 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5signature_6 := __e.Get(1)
_ = Parseshen_4_5signature_6
tmp15484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5signature_6)


if True == tmp15484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15482 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5signature_6, sym_j)


if True == tmp15482 {
tmp15469 := MakeNative(func(__e *ControlFlow) {
News3988 := __e.Get(1)
_ = News3988
tmp15470 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rules_d_6 := __e.Get(1)
_ = Parseshen_4_5rules_d_6
tmp15478 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rules_d_6)


if True == tmp15478 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15471 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5rules_d_6)


tmp15472 := MakeNative(func(__e *ControlFlow) {
Rectified := __e.Get(1)
_ = Rectified
tmp15473 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rules_d_6)


__e.TailApply(PrimFunc(sym_8p), Rectified, tmp15473)
return


}, 1)

tmp15474 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5signature_6)


tmp15475 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15474)


tmp15476 := Call(__e, tmp15472, tmp15475)


__e.TailApply(PrimFunc(symshen_4comb), tmp15471, tmp15476)
return


}


}, 1)

tmp15479 := Call(__e, PrimFunc(symshen_4_5rules_d_6), News3988)


__e.TailApply(tmp15470, tmp15479)
return


}, 1)

tmp15480 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5signature_6)


__e.TailApply(tmp15469, tmp15480)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp15485 := Call(__e, PrimFunc(symshen_4_5signature_6), News3987)


__e.TailApply(tmp15468, tmp15485)
return


}, 1)

tmp15486 := Call(__e, PrimFunc(symshen_4tls), News3986)


__e.TailApply(tmp15467, tmp15486)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15489 := Call(__e, PrimFunc(symshen_4tls), V4383)


__e.TailApply(tmp15466, tmp15489)
return


}, 1)

tmp15490 := Call(__e, PrimFunc(symshen_4hds), V4383)


tmp15491 := Call(__e, tmp15465, tmp15490)


ifres15464 = tmp15491


} else {
tmp15492 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres15464 = tmp15492


}

__e.TailApply(tmp15461, ifres15464)
return


}, 1)

tmp15494 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp15460)


_ = tmp15494

tmp15495 := MakeNative(func(__e *ControlFlow) {
V4384 := __e.Get(1)
_ = V4384
tmp15496 := MakeNative(func(__e *ControlFlow) {
Vs := __e.Get(1)
_ = Vs
tmp15497 := MakeNative(func(__e *ControlFlow) {
Assoc := __e.Get(1)
_ = Assoc
__e.TailApply(PrimFunc(symshen_4freshen_1type), Assoc, V4384)
return
}, 1)

tmp15498 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp15499 := Call(__e, PrimFunc(symconcat), sym_e, V)


tmp15500 := Call(__e, PrimFunc(symshen_4freshterm), tmp15499)


__e.Return(PrimCons(V, tmp15500))
return


}, 1)

tmp15501 := Call(__e, PrimFunc(symmap), tmp15498, Vs)


__e.TailApply(tmp15497, tmp15501)
return


}, 1)

tmp15502 := Call(__e, PrimFunc(symshen_4extract_1vars), V4384)


__e.TailApply(tmp15496, tmp15502)
return


}, 1)

tmp15503 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp15495)


_ = tmp15503

tmp15504 := MakeNative(func(__e *ControlFlow) {
V4385 := __e.Get(1)
_ = V4385
V4386 := __e.Get(2)
_ = V4386
tmp15518 := PrimEqual(Nil, V4385)

if True == tmp15518 {
__e.Return(V4386)
return
} else {
tmp15516 := PrimIsPair(V4385)

var ifres15512 Obj

if True == tmp15516 {
tmp15514 := PrimHead(V4385)

tmp15515 := PrimIsPair(tmp15514)

var ifres15513 Obj

if True == tmp15515 {
ifres15513 = True


} else {
ifres15513 = False


}

ifres15512 = ifres15513


} else {
ifres15512 = False


}

if True == ifres15512 {
tmp15505 := PrimTail(V4385)

tmp15506 := PrimHead(V4385)

tmp15507 := PrimTail(tmp15506)

tmp15508 := PrimHead(V4385)

tmp15509 := PrimHead(tmp15508)

tmp15510 := Call(__e, PrimFunc(symsubst), tmp15507, tmp15509, V4386)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp15505, tmp15510)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
return
}


}


}, 2)

tmp15519 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp15504)


_ = tmp15519

tmp15520 := MakeNative(func(__e *ControlFlow) {
V4387 := __e.Get(1)
_ = V4387
tmp15521 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15535 {
tmp15522 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15524 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp15525 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rule_d_6 := __e.Get(1)
_ = Parseshen_4_5rule_d_6
tmp15531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rule_d_6)


if True == tmp15531 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15526 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5rule_d_6)


tmp15527 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rule_d_6)


tmp15528 := Call(__e, PrimFunc(symshen_4linearise), tmp15527)


tmp15529 := PrimCons(tmp15528, Nil)

__e.TailApply(PrimFunc(symshen_4comb), tmp15526, tmp15529)
return


}


}, 1)

tmp15532 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4387)


tmp15533 := Call(__e, tmp15525, tmp15532)


__e.TailApply(tmp15522, tmp15533)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp15536 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rule_d_6 := __e.Get(1)
_ = Parseshen_4_5rule_d_6
tmp15547 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rule_d_6)


if True == tmp15547 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15537 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rules_d_6 := __e.Get(1)
_ = Parseshen_4_5rules_d_6
tmp15544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rules_d_6)


if True == tmp15544 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15538 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5rules_d_6)


tmp15539 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rule_d_6)


tmp15540 := Call(__e, PrimFunc(symshen_4linearise), tmp15539)


tmp15541 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rules_d_6)


tmp15542 := PrimCons(tmp15540, tmp15541)

__e.TailApply(PrimFunc(symshen_4comb), tmp15538, tmp15542)
return


}


}, 1)

tmp15545 := Call(__e, PrimFunc(symshen_4_5rules_d_6), Parseshen_4_5rule_d_6)


__e.TailApply(tmp15537, tmp15545)
return


}


}, 1)

tmp15548 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4387)


tmp15549 := Call(__e, tmp15536, tmp15548)


__e.TailApply(tmp15521, tmp15549)
return


}, 1)

tmp15550 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp15520)


_ = tmp15550

tmp15551 := MakeNative(func(__e *ControlFlow) {
V4388 := __e.Get(1)
_ = V4388
tmp15552 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15632 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15632 {
tmp15553 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15598 {
tmp15554 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15577 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15577 {
tmp15555 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp15557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp15557 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp15558 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp15573 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp15573 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15571 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)


if True == tmp15571 {
tmp15559 := MakeNative(func(__e *ControlFlow) {
News4001 := __e.Get(1)
_ = News4001
tmp15568 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News4001)


if True == tmp15568 {
tmp15560 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp15561 := MakeNative(func(__e *ControlFlow) {
News4002 := __e.Get(1)
_ = News4002
tmp15562 := Call(__e, PrimFunc(symshen_4in_1_6), News4002)


tmp15563 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp15564 := Call(__e, PrimFunc(sym_8p), tmp15563, Action)


__e.TailApply(PrimFunc(symshen_4comb), tmp15562, tmp15564)
return


}, 1)

tmp15565 := Call(__e, PrimFunc(symshen_4tls), News4001)


__e.TailApply(tmp15561, tmp15565)
return


}, 1)

tmp15566 := Call(__e, PrimFunc(symshen_4hds), News4001)


__e.TailApply(tmp15560, tmp15566)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15569 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp15559, tmp15569)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp15574 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4388)


tmp15575 := Call(__e, tmp15558, tmp15574)


__e.TailApply(tmp15555, tmp15575)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp15578 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp15594 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp15594 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15592 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)


if True == tmp15592 {
tmp15579 := MakeNative(func(__e *ControlFlow) {
News3999 := __e.Get(1)
_ = News3999
tmp15589 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3999)


if True == tmp15589 {
tmp15580 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp15581 := MakeNative(func(__e *ControlFlow) {
News4000 := __e.Get(1)
_ = News4000
tmp15582 := Call(__e, PrimFunc(symshen_4in_1_6), News4000)


tmp15583 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp15584 := Call(__e, PrimFunc(symshen_4correct), Action)


tmp15585 := Call(__e, PrimFunc(sym_8p), tmp15583, tmp15584)


__e.TailApply(PrimFunc(symshen_4comb), tmp15582, tmp15585)
return


}, 1)

tmp15586 := Call(__e, PrimFunc(symshen_4tls), News3999)


__e.TailApply(tmp15581, tmp15586)
return


}, 1)

tmp15587 := Call(__e, PrimFunc(symshen_4hds), News3999)


__e.TailApply(tmp15580, tmp15587)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15590 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp15579, tmp15590)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp15595 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4388)


tmp15596 := Call(__e, tmp15578, tmp15595)


__e.TailApply(tmp15554, tmp15596)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp15599 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp15628 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp15628 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15626 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)


if True == tmp15626 {
tmp15600 := MakeNative(func(__e *ControlFlow) {
News3995 := __e.Get(1)
_ = News3995
tmp15623 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3995)


if True == tmp15623 {
tmp15601 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp15602 := MakeNative(func(__e *ControlFlow) {
News3996 := __e.Get(1)
_ = News3996
tmp15619 := Call(__e, PrimFunc(symshen_4_ahd_2), News3996, symwhere)


if True == tmp15619 {
tmp15603 := MakeNative(func(__e *ControlFlow) {
News3997 := __e.Get(1)
_ = News3997
tmp15616 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3997)


if True == tmp15616 {
tmp15604 := MakeNative(func(__e *ControlFlow) {
Guard := __e.Get(1)
_ = Guard
tmp15605 := MakeNative(func(__e *ControlFlow) {
News3998 := __e.Get(1)
_ = News3998
tmp15606 := Call(__e, PrimFunc(symshen_4in_1_6), News3998)


tmp15607 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp15608 := PrimCons(Action, Nil)

tmp15609 := PrimCons(Guard, tmp15608)

tmp15610 := PrimCons(symwhere, tmp15609)

tmp15611 := Call(__e, PrimFunc(symshen_4correct), tmp15610)


tmp15612 := Call(__e, PrimFunc(sym_8p), tmp15607, tmp15611)


__e.TailApply(PrimFunc(symshen_4comb), tmp15606, tmp15612)
return


}, 1)

tmp15613 := Call(__e, PrimFunc(symshen_4tls), News3997)


__e.TailApply(tmp15605, tmp15613)
return


}, 1)

tmp15614 := Call(__e, PrimFunc(symshen_4hds), News3997)


__e.TailApply(tmp15604, tmp15614)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15617 := Call(__e, PrimFunc(symshen_4tls), News3996)


__e.TailApply(tmp15603, tmp15617)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15620 := Call(__e, PrimFunc(symshen_4tls), News3995)


__e.TailApply(tmp15602, tmp15620)
return


}, 1)

tmp15621 := Call(__e, PrimFunc(symshen_4hds), News3995)


__e.TailApply(tmp15601, tmp15621)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15624 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp15600, tmp15624)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp15629 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4388)


tmp15630 := Call(__e, tmp15599, tmp15629)


__e.TailApply(tmp15553, tmp15630)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp15633 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp15661 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp15661 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15659 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)


if True == tmp15659 {
tmp15634 := MakeNative(func(__e *ControlFlow) {
News3991 := __e.Get(1)
_ = News3991
tmp15656 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3991)


if True == tmp15656 {
tmp15635 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp15636 := MakeNative(func(__e *ControlFlow) {
News3992 := __e.Get(1)
_ = News3992
tmp15652 := Call(__e, PrimFunc(symshen_4_ahd_2), News3992, symwhere)


if True == tmp15652 {
tmp15637 := MakeNative(func(__e *ControlFlow) {
News3993 := __e.Get(1)
_ = News3993
tmp15649 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3993)


if True == tmp15649 {
tmp15638 := MakeNative(func(__e *ControlFlow) {
Guard := __e.Get(1)
_ = Guard
tmp15639 := MakeNative(func(__e *ControlFlow) {
News3994 := __e.Get(1)
_ = News3994
tmp15640 := Call(__e, PrimFunc(symshen_4in_1_6), News3994)


tmp15641 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp15642 := PrimCons(Action, Nil)

tmp15643 := PrimCons(Guard, tmp15642)

tmp15644 := PrimCons(symwhere, tmp15643)

tmp15645 := Call(__e, PrimFunc(sym_8p), tmp15641, tmp15644)


__e.TailApply(PrimFunc(symshen_4comb), tmp15640, tmp15645)
return


}, 1)

tmp15646 := Call(__e, PrimFunc(symshen_4tls), News3993)


__e.TailApply(tmp15639, tmp15646)
return


}, 1)

tmp15647 := Call(__e, PrimFunc(symshen_4hds), News3993)


__e.TailApply(tmp15638, tmp15647)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15650 := Call(__e, PrimFunc(symshen_4tls), News3992)


__e.TailApply(tmp15637, tmp15650)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15653 := Call(__e, PrimFunc(symshen_4tls), News3991)


__e.TailApply(tmp15636, tmp15653)
return


}, 1)

tmp15654 := Call(__e, PrimFunc(symshen_4hds), News3991)


__e.TailApply(tmp15635, tmp15654)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15657 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp15634, tmp15657)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp15662 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4388)


tmp15663 := Call(__e, tmp15633, tmp15662)


__e.TailApply(tmp15552, tmp15663)
return


}, 1)

tmp15664 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp15551)


_ = tmp15664

tmp15665 := MakeNative(func(__e *ControlFlow) {
V4389 := __e.Get(1)
_ = V4389
tmp15813 := PrimIsPair(V4389)

var ifres15757 Obj

if True == tmp15813 {
tmp15811 := PrimHead(V4389)

tmp15812 := PrimEqual(symwhere, tmp15811)

var ifres15759 Obj

if True == tmp15812 {
tmp15809 := PrimTail(V4389)

tmp15810 := PrimIsPair(tmp15809)

var ifres15761 Obj

if True == tmp15810 {
tmp15806 := PrimTail(V4389)

tmp15807 := PrimTail(tmp15806)

tmp15808 := PrimIsPair(tmp15807)

var ifres15763 Obj

if True == tmp15808 {
tmp15802 := PrimTail(V4389)

tmp15803 := PrimTail(tmp15802)

tmp15804 := PrimHead(tmp15803)

tmp15805 := PrimIsPair(tmp15804)

var ifres15765 Obj

if True == tmp15805 {
tmp15797 := PrimTail(V4389)

tmp15798 := PrimTail(tmp15797)

tmp15799 := PrimHead(tmp15798)

tmp15800 := PrimHead(tmp15799)

tmp15801 := PrimEqual(symfail_1if, tmp15800)

var ifres15767 Obj

if True == tmp15801 {
tmp15792 := PrimTail(V4389)

tmp15793 := PrimTail(tmp15792)

tmp15794 := PrimHead(tmp15793)

tmp15795 := PrimTail(tmp15794)

tmp15796 := PrimIsPair(tmp15795)

var ifres15769 Obj

if True == tmp15796 {
tmp15786 := PrimTail(V4389)

tmp15787 := PrimTail(tmp15786)

tmp15788 := PrimHead(tmp15787)

tmp15789 := PrimTail(tmp15788)

tmp15790 := PrimTail(tmp15789)

tmp15791 := PrimIsPair(tmp15790)

var ifres15771 Obj

if True == tmp15791 {
tmp15779 := PrimTail(V4389)

tmp15780 := PrimTail(tmp15779)

tmp15781 := PrimHead(tmp15780)

tmp15782 := PrimTail(tmp15781)

tmp15783 := PrimTail(tmp15782)

tmp15784 := PrimTail(tmp15783)

tmp15785 := PrimEqual(Nil, tmp15784)

var ifres15773 Obj

if True == tmp15785 {
tmp15775 := PrimTail(V4389)

tmp15776 := PrimTail(tmp15775)

tmp15777 := PrimTail(tmp15776)

tmp15778 := PrimEqual(Nil, tmp15777)

var ifres15774 Obj

if True == tmp15778 {
ifres15774 = True


} else {
ifres15774 = False


}

ifres15773 = ifres15774


} else {
ifres15773 = False


}

var ifres15772 Obj

if True == ifres15773 {
ifres15772 = True


} else {
ifres15772 = False


}

ifres15771 = ifres15772


} else {
ifres15771 = False


}

var ifres15770 Obj

if True == ifres15771 {
ifres15770 = True


} else {
ifres15770 = False


}

ifres15769 = ifres15770


} else {
ifres15769 = False


}

var ifres15768 Obj

if True == ifres15769 {
ifres15768 = True


} else {
ifres15768 = False


}

ifres15767 = ifres15768


} else {
ifres15767 = False


}

var ifres15766 Obj

if True == ifres15767 {
ifres15766 = True


} else {
ifres15766 = False


}

ifres15765 = ifres15766


} else {
ifres15765 = False


}

var ifres15764 Obj

if True == ifres15765 {
ifres15764 = True


} else {
ifres15764 = False


}

ifres15763 = ifres15764


} else {
ifres15763 = False


}

var ifres15762 Obj

if True == ifres15763 {
ifres15762 = True


} else {
ifres15762 = False


}

ifres15761 = ifres15762


} else {
ifres15761 = False


}

var ifres15760 Obj

if True == ifres15761 {
ifres15760 = True


} else {
ifres15760 = False


}

ifres15759 = ifres15760


} else {
ifres15759 = False


}

var ifres15758 Obj

if True == ifres15759 {
ifres15758 = True


} else {
ifres15758 = False


}

ifres15757 = ifres15758


} else {
ifres15757 = False


}

if True == ifres15757 {
tmp15666 := PrimTail(V4389)

tmp15667 := PrimHead(tmp15666)

tmp15668 := PrimTail(V4389)

tmp15669 := PrimTail(tmp15668)

tmp15670 := PrimHead(tmp15669)

tmp15671 := PrimTail(tmp15670)

tmp15672 := PrimCons(tmp15671, Nil)

tmp15673 := PrimCons(symnot, tmp15672)

tmp15674 := PrimCons(tmp15673, Nil)

tmp15675 := PrimCons(tmp15667, tmp15674)

tmp15676 := PrimCons(symand, tmp15675)

tmp15677 := PrimTail(V4389)

tmp15678 := PrimTail(tmp15677)

tmp15679 := PrimHead(tmp15678)

tmp15680 := PrimTail(tmp15679)

tmp15681 := PrimTail(tmp15680)

tmp15682 := PrimCons(tmp15676, tmp15681)

__e.Return(PrimCons(symwhere, tmp15682))
return


} else {
tmp15755 := PrimIsPair(V4389)

var ifres15736 Obj

if True == tmp15755 {
tmp15753 := PrimHead(V4389)

tmp15754 := PrimEqual(symwhere, tmp15753)

var ifres15738 Obj

if True == tmp15754 {
tmp15751 := PrimTail(V4389)

tmp15752 := PrimIsPair(tmp15751)

var ifres15740 Obj

if True == tmp15752 {
tmp15748 := PrimTail(V4389)

tmp15749 := PrimTail(tmp15748)

tmp15750 := PrimIsPair(tmp15749)

var ifres15742 Obj

if True == tmp15750 {
tmp15744 := PrimTail(V4389)

tmp15745 := PrimTail(tmp15744)

tmp15746 := PrimTail(tmp15745)

tmp15747 := PrimEqual(Nil, tmp15746)

var ifres15743 Obj

if True == tmp15747 {
ifres15743 = True


} else {
ifres15743 = False


}

ifres15742 = ifres15743


} else {
ifres15742 = False


}

var ifres15741 Obj

if True == ifres15742 {
ifres15741 = True


} else {
ifres15741 = False


}

ifres15740 = ifres15741


} else {
ifres15740 = False


}

var ifres15739 Obj

if True == ifres15740 {
ifres15739 = True


} else {
ifres15739 = False


}

ifres15738 = ifres15739


} else {
ifres15738 = False


}

var ifres15737 Obj

if True == ifres15738 {
ifres15737 = True


} else {
ifres15737 = False


}

ifres15736 = ifres15737


} else {
ifres15736 = False


}

if True == ifres15736 {
tmp15683 := PrimTail(V4389)

tmp15684 := PrimHead(tmp15683)

tmp15685 := PrimTail(V4389)

tmp15686 := PrimTail(tmp15685)

tmp15687 := PrimHead(tmp15686)

tmp15688 := PrimCons(symfail, Nil)

tmp15689 := PrimCons(tmp15688, Nil)

tmp15690 := PrimCons(tmp15687, tmp15689)

tmp15691 := PrimCons(sym_a, tmp15690)

tmp15692 := PrimCons(tmp15691, Nil)

tmp15693 := PrimCons(symnot, tmp15692)

tmp15694 := PrimCons(tmp15693, Nil)

tmp15695 := PrimCons(tmp15684, tmp15694)

tmp15696 := PrimCons(symand, tmp15695)

tmp15697 := PrimTail(V4389)

tmp15698 := PrimTail(tmp15697)

tmp15699 := PrimCons(tmp15696, tmp15698)

__e.Return(PrimCons(symwhere, tmp15699))
return


} else {
tmp15734 := PrimIsPair(V4389)

var ifres15715 Obj

if True == tmp15734 {
tmp15732 := PrimHead(V4389)

tmp15733 := PrimEqual(symfail_1if, tmp15732)

var ifres15717 Obj

if True == tmp15733 {
tmp15730 := PrimTail(V4389)

tmp15731 := PrimIsPair(tmp15730)

var ifres15719 Obj

if True == tmp15731 {
tmp15727 := PrimTail(V4389)

tmp15728 := PrimTail(tmp15727)

tmp15729 := PrimIsPair(tmp15728)

var ifres15721 Obj

if True == tmp15729 {
tmp15723 := PrimTail(V4389)

tmp15724 := PrimTail(tmp15723)

tmp15725 := PrimTail(tmp15724)

tmp15726 := PrimEqual(Nil, tmp15725)

var ifres15722 Obj

if True == tmp15726 {
ifres15722 = True


} else {
ifres15722 = False


}

ifres15721 = ifres15722


} else {
ifres15721 = False


}

var ifres15720 Obj

if True == ifres15721 {
ifres15720 = True


} else {
ifres15720 = False


}

ifres15719 = ifres15720


} else {
ifres15719 = False


}

var ifres15718 Obj

if True == ifres15719 {
ifres15718 = True


} else {
ifres15718 = False


}

ifres15717 = ifres15718


} else {
ifres15717 = False


}

var ifres15716 Obj

if True == ifres15717 {
ifres15716 = True


} else {
ifres15716 = False


}

ifres15715 = ifres15716


} else {
ifres15715 = False


}

if True == ifres15715 {
tmp15700 := PrimTail(V4389)

tmp15701 := PrimCons(tmp15700, Nil)

tmp15702 := PrimCons(symnot, tmp15701)

tmp15703 := PrimTail(V4389)

tmp15704 := PrimTail(tmp15703)

tmp15705 := PrimCons(tmp15702, tmp15704)

__e.Return(PrimCons(symwhere, tmp15705))
return


} else {
tmp15706 := PrimCons(symfail, Nil)

tmp15707 := PrimCons(tmp15706, Nil)

tmp15708 := PrimCons(V4389, tmp15707)

tmp15709 := PrimCons(sym_a, tmp15708)

tmp15710 := PrimCons(tmp15709, Nil)

tmp15711 := PrimCons(symnot, tmp15710)

tmp15712 := PrimCons(V4389, Nil)

tmp15713 := PrimCons(tmp15711, tmp15712)

__e.Return(PrimCons(symwhere, tmp15713))
return


}


}


}


}, 1)

tmp15814 := Call(__e, ns2_1set, symshen_4correct, tmp15665)


_ = tmp15814

tmp15815 := MakeNative(func(__e *ControlFlow) {
V4390 := __e.Get(1)
_ = V4390
V4391 := __e.Get(2)
_ = V4391
V4392 := __e.Get(3)
_ = V4392
V4393 := __e.Get(4)
_ = V4393
V4394 := __e.Get(5)
_ = V4394
V4395 := __e.Get(6)
_ = V4395
V4396 := __e.Get(7)
_ = V4396
V4397 := __e.Get(8)
_ = V4397
tmp15816 := MakeNative(func(__e *ControlFlow) {
K4005 := __e.Get(1)
_ = K4005
tmp15817 := MakeNative(func(__e *ControlFlow) {
C4011 := __e.Get(1)
_ = C4011
tmp15847 := PrimEqual(C4011, False)

if True == tmp15847 {
tmp15818 := MakeNative(func(__e *ControlFlow) {
C4013 := __e.Get(1)
_ = C4013
tmp15820 := PrimEqual(C4013, False)

if True == tmp15820 {
__e.TailApply(PrimFunc(symshen_4unlock), V4395, K4005)
return
} else {
__e.Return(C4013)
return
}


}, 1)

tmp15845 := Call(__e, PrimFunc(symshen_4unlocked_2), V4395)


var ifres15821 Obj

if True == tmp15845 {
tmp15822 := MakeNative(func(__e *ControlFlow) {
Tm4014 := __e.Get(1)
_ = Tm4014
tmp15842 := PrimIsPair(Tm4014)

if True == tmp15842 {
tmp15823 := MakeNative(func(__e *ControlFlow) {
Rule := __e.Get(1)
_ = Rule
tmp15824 := MakeNative(func(__e *ControlFlow) {
Rules := __e.Get(1)
_ = Rules
tmp15825 := MakeNative(func(__e *ControlFlow) {
Fresh := __e.Get(1)
_ = Fresh
tmp15826 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15826

tmp15827 := Call(__e, PrimFunc(symshen_4deref), Rule, V4394)


tmp15828 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp15827)


tmp15829 := MakeNative(func(__e *ControlFlow) {
tmp15830 := Call(__e, PrimFunc(symshen_4lazyderef), Fresh, V4394)


tmp15831 := Call(__e, PrimFunc(symfst), tmp15830)


tmp15832 := Call(__e, PrimFunc(symshen_4lazyderef), Fresh, V4394)


tmp15833 := Call(__e, PrimFunc(symsnd), tmp15832)


tmp15834 := MakeNative(func(__e *ControlFlow) {
tmp15835 := MakeNative(func(__e *ControlFlow) {
tmp15836 := PrimNumberAdd(V4393, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4t_d_1rules), V4390, Rules, V4392, tmp15836, V4394, V4395, K4005, V4397)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4394, V4395, K4005, tmp15835)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V4390, V4393, tmp15831, tmp15833, V4392, V4394, V4395, K4005, tmp15834)
return


}, 0)

tmp15837 := Call(__e, PrimFunc(symbind), Fresh, tmp15828, V4394, V4395, K4005, tmp15829)


__e.TailApply(PrimFunc(symshen_4gc), V4394, tmp15837)
return


}, 1)

tmp15838 := Call(__e, PrimFunc(symshen_4newpv), V4394)


__e.TailApply(tmp15825, tmp15838)
return


}, 1)

tmp15839 := PrimTail(Tm4014)

__e.TailApply(tmp15824, tmp15839)
return


}, 1)

tmp15840 := PrimHead(Tm4014)

__e.TailApply(tmp15823, tmp15840)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15843 := Call(__e, PrimFunc(symshen_4lazyderef), V4391, V4394)


tmp15844 := Call(__e, tmp15822, tmp15843)


ifres15821 = tmp15844


} else {
ifres15821 = False


}

__e.TailApply(tmp15818, ifres15821)
return


} else {
__e.Return(C4011)
return
}


}, 1)

tmp15855 := Call(__e, PrimFunc(symshen_4unlocked_2), V4395)


var ifres15848 Obj

if True == tmp15855 {
tmp15849 := MakeNative(func(__e *ControlFlow) {
Tm4012 := __e.Get(1)
_ = Tm4012
tmp15852 := PrimEqual(Tm4012, Nil)

if True == tmp15852 {
tmp15850 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15850

__e.TailApply(PrimFunc(symthaw), V4397)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15853 := Call(__e, PrimFunc(symshen_4lazyderef), V4391, V4394)


tmp15854 := Call(__e, tmp15849, tmp15853)


ifres15848 = tmp15854


} else {
ifres15848 = False


}

__e.TailApply(tmp15817, ifres15848)
return


}, 1)

tmp15856 := PrimNumberAdd(V4396, MakeNumber(1))

__e.TailApply(tmp15816, tmp15856)
return


}, 8)

tmp15857 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp15815)


_ = tmp15857

tmp15858 := MakeNative(func(__e *ControlFlow) {
V4398 := __e.Get(1)
_ = V4398
tmp15871 := Call(__e, PrimFunc(symtuple_2), V4398)


if True == tmp15871 {
tmp15859 := MakeNative(func(__e *ControlFlow) {
Vs := __e.Get(1)
_ = Vs
tmp15860 := MakeNative(func(__e *ControlFlow) {
Assoc := __e.Get(1)
_ = Assoc
tmp15861 := Call(__e, PrimFunc(symfst), V4398)


tmp15862 := Call(__e, PrimFunc(symshen_4freshen), Assoc, tmp15861)


tmp15863 := Call(__e, PrimFunc(symsnd), V4398)


tmp15864 := Call(__e, PrimFunc(symshen_4freshen), Assoc, tmp15863)


__e.TailApply(PrimFunc(sym_8p), tmp15862, tmp15864)
return


}, 1)

tmp15865 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp15866 := Call(__e, PrimFunc(symshen_4freshterm), V)


__e.Return(PrimCons(V, tmp15866))
return


}, 1)

tmp15867 := Call(__e, PrimFunc(symmap), tmp15865, Vs)


__e.TailApply(tmp15860, tmp15867)
return


}, 1)

tmp15868 := Call(__e, PrimFunc(symfst), V4398)


tmp15869 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp15868)


__e.TailApply(tmp15859, tmp15869)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
return
}


}, 1)

tmp15872 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp15858)


_ = tmp15872

tmp15873 := MakeNative(func(__e *ControlFlow) {
V4399 := __e.Get(1)
_ = V4399
V4400 := __e.Get(2)
_ = V4400
tmp15887 := PrimEqual(Nil, V4399)

if True == tmp15887 {
__e.Return(V4400)
return
} else {
tmp15885 := PrimIsPair(V4399)

var ifres15881 Obj

if True == tmp15885 {
tmp15883 := PrimHead(V4399)

tmp15884 := PrimIsPair(tmp15883)

var ifres15882 Obj

if True == tmp15884 {
ifres15882 = True


} else {
ifres15882 = False


}

ifres15881 = ifres15882


} else {
ifres15881 = False


}

if True == ifres15881 {
tmp15874 := PrimTail(V4399)

tmp15875 := PrimHead(V4399)

tmp15876 := PrimHead(tmp15875)

tmp15877 := PrimHead(V4399)

tmp15878 := PrimTail(tmp15877)

tmp15879 := Call(__e, PrimFunc(symshen_4beta), tmp15876, tmp15878, V4400)


__e.TailApply(PrimFunc(symshen_4freshen), tmp15874, tmp15879)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
return
}


}


}, 2)

tmp15888 := Call(__e, ns2_1set, symshen_4freshen, tmp15873)


_ = tmp15888

tmp15889 := MakeNative(func(__e *ControlFlow) {
V4401 := __e.Get(1)
_ = V4401
V4402 := __e.Get(2)
_ = V4402
V4403 := __e.Get(3)
_ = V4403
V4404 := __e.Get(4)
_ = V4404
V4405 := __e.Get(5)
_ = V4405
V4406 := __e.Get(6)
_ = V4406
V4407 := __e.Get(7)
_ = V4407
V4408 := __e.Get(8)
_ = V4408
V4409 := __e.Get(9)
_ = V4409
tmp15890 := MakeNative(func(__e *ControlFlow) {
C4024 := __e.Get(1)
_ = C4024
tmp15903 := PrimEqual(C4024, False)

if True == tmp15903 {
tmp15901 := Call(__e, PrimFunc(symshen_4unlocked_2), V4407)


if True == tmp15901 {
tmp15891 := MakeNative(func(__e *ControlFlow) {
Err := __e.Get(1)
_ = Err
tmp15892 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15892

tmp15893 := Call(__e, PrimFunc(symshen_4app), V4401, MakeString("\n"), symshen_4a)


tmp15894 := PrimStringConcat(MakeString(" of "), tmp15893)

tmp15895 := Call(__e, PrimFunc(symshen_4app), V4402, tmp15894, symshen_4a)


tmp15896 := PrimStringConcat(MakeString("type error in rule "), tmp15895)

tmp15897 := PrimSimpleError(tmp15896)

tmp15898 := Call(__e, PrimFunc(symbind), Err, tmp15897, V4406, V4407, V4408, V4409)


__e.TailApply(PrimFunc(symshen_4gc), V4406, tmp15898)
return


}, 1)

tmp15899 := Call(__e, PrimFunc(symshen_4newpv), V4406)


__e.TailApply(tmp15891, tmp15899)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C4024)
return
}


}, 1)

tmp15907 := Call(__e, PrimFunc(symshen_4unlocked_2), V4407)


var ifres15904 Obj

if True == tmp15907 {
tmp15905 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15905

tmp15906 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V4403, V4404, V4405, V4406, V4407, V4408, V4409)


ifres15904 = tmp15906


} else {
ifres15904 = False


}

__e.TailApply(tmp15890, ifres15904)
return


}, 9)

tmp15908 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp15889)


_ = tmp15908

tmp15909 := MakeNative(func(__e *ControlFlow) {
V4410 := __e.Get(1)
_ = V4410
V4411 := __e.Get(2)
_ = V4411
V4412 := __e.Get(3)
_ = V4412
V4413 := __e.Get(4)
_ = V4413
V4414 := __e.Get(5)
_ = V4414
V4415 := __e.Get(6)
_ = V4415
V4416 := __e.Get(7)
_ = V4416
tmp15910 := MakeNative(func(__e *ControlFlow) {
K4027 := __e.Get(1)
_ = K4027
tmp15911 := MakeNative(func(__e *ControlFlow) {
C4032 := __e.Get(1)
_ = C4032
tmp15928 := PrimEqual(C4032, False)

if True == tmp15928 {
tmp15912 := MakeNative(func(__e *ControlFlow) {
C4038 := __e.Get(1)
_ = C4038
tmp15914 := PrimEqual(C4038, False)

if True == tmp15914 {
__e.TailApply(PrimFunc(symshen_4unlock), V4414, K4027)
return
} else {
__e.Return(C4038)
return
}


}, 1)

tmp15926 := Call(__e, PrimFunc(symshen_4unlocked_2), V4414)


var ifres15915 Obj

if True == tmp15926 {
tmp15916 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp15917 := MakeNative(func(__e *ControlFlow) {
Hyps := __e.Get(1)
_ = Hyps
tmp15918 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15918

tmp15919 := MakeNative(func(__e *ControlFlow) {
tmp15920 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V4411, B, Hyps, V4413, V4414, K4027, V4416)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4413, V4414, K4027, tmp15920)
return


}, 0)

tmp15921 := Call(__e, PrimFunc(symshen_4t_d_1integrity), V4410, V4412, Hyps, B, V4413, V4414, K4027, tmp15919)


__e.TailApply(PrimFunc(symshen_4gc), V4413, tmp15921)
return


}, 1)

tmp15922 := Call(__e, PrimFunc(symshen_4newpv), V4413)


tmp15923 := Call(__e, tmp15917, tmp15922)


__e.TailApply(PrimFunc(symshen_4gc), V4413, tmp15923)
return


}, 1)

tmp15924 := Call(__e, PrimFunc(symshen_4newpv), V4413)


tmp15925 := Call(__e, tmp15916, tmp15924)


ifres15915 = tmp15925


} else {
ifres15915 = False


}

__e.TailApply(tmp15912, ifres15915)
return


} else {
__e.Return(C4032)
return
}


}, 1)

tmp15958 := Call(__e, PrimFunc(symshen_4unlocked_2), V4414)


var ifres15929 Obj

if True == tmp15958 {
tmp15930 := MakeNative(func(__e *ControlFlow) {
Tm4033 := __e.Get(1)
_ = Tm4033
tmp15955 := PrimEqual(Tm4033, Nil)

if True == tmp15955 {
tmp15931 := MakeNative(func(__e *ControlFlow) {
Tm4034 := __e.Get(1)
_ = Tm4034
tmp15952 := PrimIsPair(Tm4034)

if True == tmp15952 {
tmp15932 := MakeNative(func(__e *ControlFlow) {
Tm4035 := __e.Get(1)
_ = Tm4035
tmp15948 := PrimEqual(Tm4035, sym_1_1_6)

if True == tmp15948 {
tmp15933 := MakeNative(func(__e *ControlFlow) {
Tm4036 := __e.Get(1)
_ = Tm4036
tmp15944 := PrimIsPair(Tm4036)

if True == tmp15944 {
tmp15934 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp15935 := MakeNative(func(__e *ControlFlow) {
Tm4037 := __e.Get(1)
_ = Tm4037
tmp15939 := PrimEqual(Tm4037, Nil)

if True == tmp15939 {
tmp15936 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15936

tmp15937 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V4411, A, Nil, V4413, V4414, K4027, V4416)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4413, V4414, K4027, tmp15937)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15940 := PrimTail(Tm4036)

tmp15941 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15940, V4413)


__e.TailApply(tmp15935, tmp15941)
return


}, 1)

tmp15942 := PrimHead(Tm4036)

__e.TailApply(tmp15934, tmp15942)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15945 := PrimTail(Tm4034)

tmp15946 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15945, V4413)


__e.TailApply(tmp15933, tmp15946)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15949 := PrimHead(Tm4034)

tmp15950 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15949, V4413)


__e.TailApply(tmp15932, tmp15950)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15953 := Call(__e, PrimFunc(symshen_4lazyderef), V4412, V4413)


__e.TailApply(tmp15931, tmp15953)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15956 := Call(__e, PrimFunc(symshen_4lazyderef), V4410, V4413)


tmp15957 := Call(__e, tmp15930, tmp15956)


ifres15929 = tmp15957


} else {
ifres15929 = False


}

__e.TailApply(tmp15911, ifres15929)
return


}, 1)

tmp15959 := PrimNumberAdd(V4415, MakeNumber(1))

__e.TailApply(tmp15910, tmp15959)
return


}, 7)

tmp15960 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp15909)


_ = tmp15960

tmp15961 := MakeNative(func(__e *ControlFlow) {
V4417 := __e.Get(1)
_ = V4417
V4418 := __e.Get(2)
_ = V4418
V4419 := __e.Get(3)
_ = V4419
V4420 := __e.Get(4)
_ = V4420
V4421 := __e.Get(5)
_ = V4421
V4422 := __e.Get(6)
_ = V4422
V4423 := __e.Get(7)
_ = V4423
tmp15962 := MakeNative(func(__e *ControlFlow) {
K4041 := __e.Get(1)
_ = K4041
tmp15963 := MakeNative(func(__e *ControlFlow) {
C4046 := __e.Get(1)
_ = C4046
tmp15973 := PrimEqual(C4046, False)

if True == tmp15973 {
tmp15964 := MakeNative(func(__e *ControlFlow) {
C4052 := __e.Get(1)
_ = C4052
tmp15966 := PrimEqual(C4052, False)

if True == tmp15966 {
__e.TailApply(PrimFunc(symshen_4unlock), V4421, K4041)
return
} else {
__e.Return(C4052)
return
}


}, 1)

tmp15971 := Call(__e, PrimFunc(symshen_4unlocked_2), V4421)


var ifres15967 Obj

if True == tmp15971 {
tmp15968 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15968

tmp15969 := Call(__e, PrimFunc(symshen_4curry), V4417)


tmp15970 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp15969, V4418, V4419, V4420, V4421, K4041, V4423)


ifres15967 = tmp15970


} else {
ifres15967 = False


}

__e.TailApply(tmp15964, ifres15967)
return


} else {
__e.Return(C4046)
return
}


}, 1)

tmp16018 := Call(__e, PrimFunc(symshen_4unlocked_2), V4421)


var ifres15974 Obj

if True == tmp16018 {
tmp15975 := MakeNative(func(__e *ControlFlow) {
Tm4047 := __e.Get(1)
_ = Tm4047
tmp16015 := PrimIsPair(Tm4047)

if True == tmp16015 {
tmp15976 := MakeNative(func(__e *ControlFlow) {
Tm4048 := __e.Get(1)
_ = Tm4048
tmp16011 := PrimEqual(Tm4048, symwhere)

if True == tmp16011 {
tmp15977 := MakeNative(func(__e *ControlFlow) {
Tm4049 := __e.Get(1)
_ = Tm4049
tmp16007 := PrimIsPair(Tm4049)

if True == tmp16007 {
tmp15978 := MakeNative(func(__e *ControlFlow) {
G := __e.Get(1)
_ = G
tmp15979 := MakeNative(func(__e *ControlFlow) {
Tm4050 := __e.Get(1)
_ = Tm4050
tmp16002 := PrimIsPair(Tm4050)

if True == tmp16002 {
tmp15980 := MakeNative(func(__e *ControlFlow) {
R := __e.Get(1)
_ = R
tmp15981 := MakeNative(func(__e *ControlFlow) {
Tm4051 := __e.Get(1)
_ = Tm4051
tmp15997 := PrimEqual(Tm4051, Nil)

if True == tmp15997 {
tmp15982 := MakeNative(func(__e *ControlFlow) {
CurryG := __e.Get(1)
_ = CurryG
tmp15983 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15983

tmp15984 := MakeNative(func(__e *ControlFlow) {
tmp15985 := Call(__e, PrimFunc(symshen_4curry), G)


tmp15986 := MakeNative(func(__e *ControlFlow) {
tmp15987 := MakeNative(func(__e *ControlFlow) {
tmp15988 := MakeNative(func(__e *ControlFlow) {
tmp15989 := PrimIntern(MakeString(":"))

tmp15990 := PrimCons(symverified, Nil)

tmp15991 := PrimCons(tmp15989, tmp15990)

tmp15992 := PrimCons(CurryG, tmp15991)

tmp15993 := PrimCons(tmp15992, V4419)

__e.TailApply(PrimFunc(symshen_4t_d_1correct), R, V4418, tmp15993, V4420, V4421, K4041, V4423)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4420, V4421, K4041, tmp15988)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), CurryG, symboolean, V4419, V4420, V4421, K4041, tmp15987)
return


}, 0)

__e.TailApply(PrimFunc(symbind), CurryG, tmp15985, V4420, V4421, K4041, tmp15986)
return


}, 0)

tmp15994 := Call(__e, PrimFunc(symshen_4cut), V4420, V4421, K4041, tmp15984)


__e.TailApply(PrimFunc(symshen_4gc), V4420, tmp15994)
return


}, 1)

tmp15995 := Call(__e, PrimFunc(symshen_4newpv), V4420)


__e.TailApply(tmp15982, tmp15995)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15998 := PrimTail(Tm4050)

tmp15999 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15998, V4420)


__e.TailApply(tmp15981, tmp15999)
return


}, 1)

tmp16000 := PrimHead(Tm4050)

__e.TailApply(tmp15980, tmp16000)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16003 := PrimTail(Tm4049)

tmp16004 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16003, V4420)


__e.TailApply(tmp15979, tmp16004)
return


}, 1)

tmp16005 := PrimHead(Tm4049)

__e.TailApply(tmp15978, tmp16005)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16008 := PrimTail(Tm4047)

tmp16009 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16008, V4420)


__e.TailApply(tmp15977, tmp16009)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16012 := PrimHead(Tm4047)

tmp16013 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16012, V4420)


__e.TailApply(tmp15976, tmp16013)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16016 := Call(__e, PrimFunc(symshen_4lazyderef), V4417, V4420)


tmp16017 := Call(__e, tmp15975, tmp16016)


ifres15974 = tmp16017


} else {
ifres15974 = False


}

__e.TailApply(tmp15963, ifres15974)
return


}, 1)

tmp16019 := PrimNumberAdd(V4422, MakeNumber(1))

__e.TailApply(tmp15962, tmp16019)
return


}, 7)

tmp16020 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp15961)


_ = tmp16020

tmp16021 := MakeNative(func(__e *ControlFlow) {
V4424 := __e.Get(1)
_ = V4424
V4425 := __e.Get(2)
_ = V4425
V4426 := __e.Get(3)
_ = V4426
V4427 := __e.Get(4)
_ = V4427
V4428 := __e.Get(5)
_ = V4428
V4429 := __e.Get(6)
_ = V4429
V4430 := __e.Get(7)
_ = V4430
V4431 := __e.Get(8)
_ = V4431
tmp16022 := MakeNative(func(__e *ControlFlow) {
K4056 := __e.Get(1)
_ = K4056
tmp16023 := MakeNative(func(__e *ControlFlow) {
C4062 := __e.Get(1)
_ = C4062
tmp16102 := PrimEqual(C4062, False)

if True == tmp16102 {
tmp16024 := MakeNative(func(__e *ControlFlow) {
C4066 := __e.Get(1)
_ = C4066
tmp16026 := PrimEqual(C4066, False)

if True == tmp16026 {
__e.TailApply(PrimFunc(symshen_4unlock), V4429, K4056)
return
} else {
__e.Return(C4066)
return
}


}, 1)

tmp16100 := Call(__e, PrimFunc(symshen_4unlocked_2), V4429)


var ifres16027 Obj

if True == tmp16100 {
tmp16028 := MakeNative(func(__e *ControlFlow) {
Tm4067 := __e.Get(1)
_ = Tm4067
tmp16097 := PrimIsPair(Tm4067)

if True == tmp16097 {
tmp16029 := MakeNative(func(__e *ControlFlow) {
P := __e.Get(1)
_ = P
tmp16030 := MakeNative(func(__e *ControlFlow) {
Ps := __e.Get(1)
_ = Ps
tmp16031 := MakeNative(func(__e *ControlFlow) {
Tm4068 := __e.Get(1)
_ = Tm4068
tmp16092 := PrimIsPair(Tm4068)

if True == tmp16092 {
tmp16032 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp16033 := MakeNative(func(__e *ControlFlow) {
Tm4069 := __e.Get(1)
_ = Tm4069
tmp16087 := PrimIsPair(Tm4069)

if True == tmp16087 {
tmp16034 := MakeNative(func(__e *ControlFlow) {
Tm4070 := __e.Get(1)
_ = Tm4070
tmp16083 := PrimEqual(Tm4070, sym_1_1_6)

if True == tmp16083 {
tmp16035 := MakeNative(func(__e *ControlFlow) {
Tm4071 := __e.Get(1)
_ = Tm4071
tmp16079 := PrimIsPair(Tm4071)

if True == tmp16079 {
tmp16036 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp16037 := MakeNative(func(__e *ControlFlow) {
Tm4072 := __e.Get(1)
_ = Tm4072
tmp16074 := PrimEqual(Tm4072, Nil)

if True == tmp16074 {
tmp16038 := MakeNative(func(__e *ControlFlow) {
Tm4073 := __e.Get(1)
_ = Tm4073
tmp16039 := MakeNative(func(__e *ControlFlow) {
GoTo4074 := __e.Get(1)
_ = GoTo4074
tmp16057 := PrimIsPair(Tm4073)

if True == tmp16057 {
tmp16040 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp16041 := MakeNative(func(__e *ControlFlow) {
Hyps := __e.Get(1)
_ = Hyps
tmp16042 := Call(__e, GoTo4074, Hyp)


__e.TailApply(tmp16042, Hyps)
return


}, 1)

tmp16043 := PrimTail(Tm4073)

__e.TailApply(tmp16041, tmp16043)
return


}, 1)

tmp16044 := PrimHead(Tm4073)

__e.TailApply(tmp16040, tmp16044)
return


} else {
tmp16055 := Call(__e, PrimFunc(symshen_4pvar_2), Tm4073)


if True == tmp16055 {
tmp16045 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp16046 := MakeNative(func(__e *ControlFlow) {
Hyps := __e.Get(1)
_ = Hyps
tmp16047 := PrimCons(Hyp, Hyps)

tmp16048 := MakeNative(func(__e *ControlFlow) {
tmp16049 := Call(__e, GoTo4074, Hyp)


__e.TailApply(tmp16049, Hyps)
return


}, 0)

tmp16050 := Call(__e, PrimFunc(symshen_4bind_b), Tm4073, tmp16047, V4428, tmp16048)


__e.TailApply(PrimFunc(symshen_4gc), V4428, tmp16050)
return


}, 1)

tmp16051 := Call(__e, PrimFunc(symshen_4newpv), V4428)


tmp16052 := Call(__e, tmp16046, tmp16051)


__e.TailApply(PrimFunc(symshen_4gc), V4428, tmp16052)
return


}, 1)

tmp16053 := Call(__e, PrimFunc(symshen_4newpv), V4428)


__e.TailApply(tmp16045, tmp16053)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16058 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
__e.Return(MakeNative(func(__e *ControlFlow) {
Hyps := __e.Get(1)
_ = Hyps
tmp16059 := MakeNative(func(__e *ControlFlow) {
PHyps := __e.Get(1)
_ = PHyps
tmp16060 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16060

tmp16061 := PrimIntern(MakeString(":"))

tmp16062 := PrimCons(A, Nil)

tmp16063 := PrimCons(tmp16061, tmp16062)

tmp16064 := PrimCons(P, tmp16063)

tmp16065 := MakeNative(func(__e *ControlFlow) {
tmp16066 := MakeNative(func(__e *ControlFlow) {
tmp16067 := MakeNative(func(__e *ControlFlow) {
tmp16068 := MakeNative(func(__e *ControlFlow) {
tmp16069 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), Ps, B, Hyps, V4427, V4428, V4429, K4056, V4431)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4428, V4429, K4056, tmp16069)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), P, A, PHyps, V4428, V4429, K4056, tmp16068)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4428, V4429, K4056, tmp16067)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4p_1hyps), P, PHyps, V4428, V4429, K4056, tmp16066)
return


}, 0)

tmp16070 := Call(__e, PrimFunc(symbind), Hyp, tmp16064, V4428, V4429, K4056, tmp16065)


__e.TailApply(PrimFunc(symshen_4gc), V4428, tmp16070)
return


}, 1)

tmp16071 := Call(__e, PrimFunc(symshen_4newpv), V4428)


__e.TailApply(tmp16059, tmp16071)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16039, tmp16058)
return


}, 1)

tmp16072 := Call(__e, PrimFunc(symshen_4lazyderef), V4426, V4428)


__e.TailApply(tmp16038, tmp16072)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16075 := PrimTail(Tm4071)

tmp16076 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16075, V4428)


__e.TailApply(tmp16037, tmp16076)
return


}, 1)

tmp16077 := PrimHead(Tm4071)

__e.TailApply(tmp16036, tmp16077)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16080 := PrimTail(Tm4069)

tmp16081 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16080, V4428)


__e.TailApply(tmp16035, tmp16081)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16084 := PrimHead(Tm4069)

tmp16085 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16084, V4428)


__e.TailApply(tmp16034, tmp16085)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16088 := PrimTail(Tm4068)

tmp16089 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16088, V4428)


__e.TailApply(tmp16033, tmp16089)
return


}, 1)

tmp16090 := PrimHead(Tm4068)

__e.TailApply(tmp16032, tmp16090)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16093 := Call(__e, PrimFunc(symshen_4lazyderef), V4425, V4428)


__e.TailApply(tmp16031, tmp16093)
return


}, 1)

tmp16094 := PrimTail(Tm4067)

__e.TailApply(tmp16030, tmp16094)
return


}, 1)

tmp16095 := PrimHead(Tm4067)

__e.TailApply(tmp16029, tmp16095)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16098 := Call(__e, PrimFunc(symshen_4lazyderef), V4424, V4428)


tmp16099 := Call(__e, tmp16028, tmp16098)


ifres16027 = tmp16099


} else {
ifres16027 = False


}

__e.TailApply(tmp16024, ifres16027)
return


} else {
__e.Return(C4062)
return
}


}, 1)

tmp16118 := Call(__e, PrimFunc(symshen_4unlocked_2), V4429)


var ifres16103 Obj

if True == tmp16118 {
tmp16104 := MakeNative(func(__e *ControlFlow) {
Tm4063 := __e.Get(1)
_ = Tm4063
tmp16115 := PrimEqual(Tm4063, Nil)

if True == tmp16115 {
tmp16105 := MakeNative(func(__e *ControlFlow) {
Tm4064 := __e.Get(1)
_ = Tm4064
tmp16106 := MakeNative(func(__e *ControlFlow) {
GoTo4065 := __e.Get(1)
_ = GoTo4065
tmp16110 := PrimEqual(Tm4064, Nil)

if True == tmp16110 {
__e.TailApply(PrimFunc(symthaw), GoTo4065)
return
} else {
tmp16108 := Call(__e, PrimFunc(symshen_4pvar_2), Tm4064)


if True == tmp16108 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm4064, Nil, V4428, GoTo4065)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16111 := MakeNative(func(__e *ControlFlow) {
tmp16112 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16112

__e.TailApply(PrimFunc(symis_b), V4425, V4427, V4428, V4429, K4056, V4431)
return


}, 0)

__e.TailApply(tmp16106, tmp16111)
return


}, 1)

tmp16113 := Call(__e, PrimFunc(symshen_4lazyderef), V4426, V4428)


__e.TailApply(tmp16105, tmp16113)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16116 := Call(__e, PrimFunc(symshen_4lazyderef), V4424, V4428)


tmp16117 := Call(__e, tmp16104, tmp16116)


ifres16103 = tmp16117


} else {
ifres16103 = False


}

__e.TailApply(tmp16023, ifres16103)
return


}, 1)

tmp16119 := PrimNumberAdd(V4430, MakeNumber(1))

__e.TailApply(tmp16022, tmp16119)
return


}, 8)

tmp16120 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp16021)


_ = tmp16120

tmp16121 := MakeNative(func(__e *ControlFlow) {
V4432 := __e.Get(1)
_ = V4432
V4433 := __e.Get(2)
_ = V4433
V4434 := __e.Get(3)
_ = V4434
V4435 := __e.Get(4)
_ = V4435
V4436 := __e.Get(5)
_ = V4436
V4437 := __e.Get(6)
_ = V4437
tmp16122 := MakeNative(func(__e *ControlFlow) {
K4077 := __e.Get(1)
_ = K4077
tmp16123 := MakeNative(func(__e *ControlFlow) {
C4081 := __e.Get(1)
_ = C4081
tmp16156 := PrimEqual(C4081, False)

if True == tmp16156 {
tmp16124 := MakeNative(func(__e *ControlFlow) {
C4082 := __e.Get(1)
_ = C4082
tmp16133 := PrimEqual(C4082, False)

if True == tmp16133 {
tmp16125 := MakeNative(func(__e *ControlFlow) {
C4084 := __e.Get(1)
_ = C4084
tmp16127 := PrimEqual(C4084, False)

if True == tmp16127 {
__e.TailApply(PrimFunc(symshen_4unlock), V4435, K4077)
return
} else {
__e.Return(C4084)
return
}


}, 1)

tmp16131 := Call(__e, PrimFunc(symshen_4unlocked_2), V4435)


var ifres16128 Obj

if True == tmp16131 {
tmp16129 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16129

tmp16130 := Call(__e, PrimFunc(symbind), V4433, Nil, V4434, V4435, K4077, V4437)


ifres16128 = tmp16130


} else {
ifres16128 = False


}

__e.TailApply(tmp16125, ifres16128)
return


} else {
__e.Return(C4082)
return
}


}, 1)

tmp16154 := Call(__e, PrimFunc(symshen_4unlocked_2), V4435)


var ifres16134 Obj

if True == tmp16154 {
tmp16135 := MakeNative(func(__e *ControlFlow) {
Tm4083 := __e.Get(1)
_ = Tm4083
tmp16151 := PrimIsPair(Tm4083)

if True == tmp16151 {
tmp16136 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp16137 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp16138 := MakeNative(func(__e *ControlFlow) {
XHyps := __e.Get(1)
_ = XHyps
tmp16139 := MakeNative(func(__e *ControlFlow) {
YHyps := __e.Get(1)
_ = YHyps
tmp16140 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16140

tmp16141 := MakeNative(func(__e *ControlFlow) {
tmp16142 := MakeNative(func(__e *ControlFlow) {
tmp16143 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4join), XHyps, YHyps, V4433, V4434, V4435, K4077, V4437)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4p_1hyps), Y, YHyps, V4434, V4435, K4077, tmp16143)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4p_1hyps), X, XHyps, V4434, V4435, K4077, tmp16142)
return


}, 0)

tmp16144 := Call(__e, PrimFunc(symshen_4cut), V4434, V4435, K4077, tmp16141)


__e.TailApply(PrimFunc(symshen_4gc), V4434, tmp16144)
return


}, 1)

tmp16145 := Call(__e, PrimFunc(symshen_4newpv), V4434)


tmp16146 := Call(__e, tmp16139, tmp16145)


__e.TailApply(PrimFunc(symshen_4gc), V4434, tmp16146)
return


}, 1)

tmp16147 := Call(__e, PrimFunc(symshen_4newpv), V4434)


__e.TailApply(tmp16138, tmp16147)
return


}, 1)

tmp16148 := PrimTail(Tm4083)

__e.TailApply(tmp16137, tmp16148)
return


}, 1)

tmp16149 := PrimHead(Tm4083)

__e.TailApply(tmp16136, tmp16149)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16152 := Call(__e, PrimFunc(symshen_4lazyderef), V4432, V4434)


tmp16153 := Call(__e, tmp16135, tmp16152)


ifres16134 = tmp16153


} else {
ifres16134 = False


}

__e.TailApply(tmp16124, ifres16134)
return


} else {
__e.Return(C4081)
return
}


}, 1)

tmp16172 := Call(__e, PrimFunc(symshen_4unlocked_2), V4435)


var ifres16157 Obj

if True == tmp16172 {
tmp16158 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp16159 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16159

tmp16160 := Call(__e, PrimFunc(symshen_4deref), V4432, V4434)


tmp16161 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp16160)


tmp16162 := MakeNative(func(__e *ControlFlow) {
tmp16163 := MakeNative(func(__e *ControlFlow) {
tmp16164 := PrimIntern(MakeString(":"))

tmp16165 := PrimCons(A, Nil)

tmp16166 := PrimCons(tmp16164, tmp16165)

tmp16167 := PrimCons(V4432, tmp16166)

tmp16168 := PrimCons(tmp16167, Nil)

__e.TailApply(PrimFunc(symbind), V4433, tmp16168, V4434, V4435, K4077, V4437)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4434, V4435, K4077, tmp16163)
return


}, 0)

tmp16169 := Call(__e, PrimFunc(symwhen), tmp16161, V4434, V4435, K4077, tmp16162)


__e.TailApply(PrimFunc(symshen_4gc), V4434, tmp16169)
return


}, 1)

tmp16170 := Call(__e, PrimFunc(symshen_4newpv), V4434)


tmp16171 := Call(__e, tmp16158, tmp16170)


ifres16157 = tmp16171


} else {
ifres16157 = False


}

__e.TailApply(tmp16123, ifres16157)
return


}, 1)

tmp16173 := PrimNumberAdd(V4436, MakeNumber(1))

__e.TailApply(tmp16122, tmp16173)
return


}, 6)

tmp16174 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp16121)


_ = tmp16174

tmp16175 := MakeNative(func(__e *ControlFlow) {
V4438 := __e.Get(1)
_ = V4438
tmp16184 := PrimIsVector(V4438)

if True == tmp16184 {
tmp16181 := PrimIsString(V4438)

tmp16182 := PrimNot(tmp16181)

var ifres16177 Obj

if True == tmp16182 {
tmp16179 := PrimVectorGet(V4438, MakeNumber(0))

tmp16180 := PrimEqual(tmp16179, symshen_4print_1freshterm)

var ifres16178 Obj

if True == tmp16180 {
ifres16178 = True


} else {
ifres16178 = False


}

ifres16177 = ifres16178


} else {
ifres16177 = False


}

if True == ifres16177 {
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

tmp16185 := Call(__e, ns2_1set, symshen_4freshterm_2, tmp16175)


_ = tmp16185

tmp16186 := MakeNative(func(__e *ControlFlow) {
V4439 := __e.Get(1)
_ = V4439
V4440 := __e.Get(2)
_ = V4440
V4441 := __e.Get(3)
_ = V4441
V4442 := __e.Get(4)
_ = V4442
V4443 := __e.Get(5)
_ = V4443
V4444 := __e.Get(6)
_ = V4444
V4445 := __e.Get(7)
_ = V4445
tmp16187 := MakeNative(func(__e *ControlFlow) {
C4092 := __e.Get(1)
_ = C4092
tmp16223 := PrimEqual(C4092, False)

if True == tmp16223 {
tmp16221 := Call(__e, PrimFunc(symshen_4unlocked_2), V4443)


if True == tmp16221 {
tmp16188 := MakeNative(func(__e *ControlFlow) {
Tm4094 := __e.Get(1)
_ = Tm4094
tmp16218 := PrimIsPair(Tm4094)

if True == tmp16218 {
tmp16189 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp16190 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp16191 := MakeNative(func(__e *ControlFlow) {
Tm4095 := __e.Get(1)
_ = Tm4095
tmp16192 := MakeNative(func(__e *ControlFlow) {
GoTo4096 := __e.Get(1)
_ = GoTo4096
tmp16210 := PrimIsPair(Tm4095)

if True == tmp16210 {
tmp16193 := MakeNative(func(__e *ControlFlow) {
X_d := __e.Get(1)
_ = X_d
tmp16194 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
tmp16195 := Call(__e, GoTo4096, X_d)


__e.TailApply(tmp16195, Z)
return


}, 1)

tmp16196 := PrimTail(Tm4095)

__e.TailApply(tmp16194, tmp16196)
return


}, 1)

tmp16197 := PrimHead(Tm4095)

__e.TailApply(tmp16193, tmp16197)
return


} else {
tmp16208 := Call(__e, PrimFunc(symshen_4pvar_2), Tm4095)


if True == tmp16208 {
tmp16198 := MakeNative(func(__e *ControlFlow) {
X_d := __e.Get(1)
_ = X_d
tmp16199 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
tmp16200 := PrimCons(X_d, Z)

tmp16201 := MakeNative(func(__e *ControlFlow) {
tmp16202 := Call(__e, GoTo4096, X_d)


__e.TailApply(tmp16202, Z)
return


}, 0)

tmp16203 := Call(__e, PrimFunc(symshen_4bind_b), Tm4095, tmp16200, V4442, tmp16201)


__e.TailApply(PrimFunc(symshen_4gc), V4442, tmp16203)
return


}, 1)

tmp16204 := Call(__e, PrimFunc(symshen_4newpv), V4442)


tmp16205 := Call(__e, tmp16199, tmp16204)


__e.TailApply(PrimFunc(symshen_4gc), V4442, tmp16205)
return


}, 1)

tmp16206 := Call(__e, PrimFunc(symshen_4newpv), V4442)


__e.TailApply(tmp16198, tmp16206)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16211 := MakeNative(func(__e *ControlFlow) {
X_d := __e.Get(1)
_ = X_d
__e.Return(MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
tmp16212 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16212

tmp16213 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4join), Y, V4440, Z, V4442, V4443, V4444, V4445)
return
}, 0)

__e.TailApply(PrimFunc(symbind), X_d, X, V4442, V4443, V4444, tmp16213)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16192, tmp16211)
return


}, 1)

tmp16214 := Call(__e, PrimFunc(symshen_4lazyderef), V4441, V4442)


__e.TailApply(tmp16191, tmp16214)
return


}, 1)

tmp16215 := PrimTail(Tm4094)

__e.TailApply(tmp16190, tmp16215)
return


}, 1)

tmp16216 := PrimHead(Tm4094)

__e.TailApply(tmp16189, tmp16216)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16219 := Call(__e, PrimFunc(symshen_4lazyderef), V4439, V4442)


__e.TailApply(tmp16188, tmp16219)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(C4092)
return
}


}, 1)

tmp16231 := Call(__e, PrimFunc(symshen_4unlocked_2), V4443)


var ifres16224 Obj

if True == tmp16231 {
tmp16225 := MakeNative(func(__e *ControlFlow) {
Tm4093 := __e.Get(1)
_ = Tm4093
tmp16228 := PrimEqual(Tm4093, Nil)

if True == tmp16228 {
tmp16226 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16226

__e.TailApply(PrimFunc(symbind), V4441, V4440, V4442, V4443, V4444, V4445)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16229 := Call(__e, PrimFunc(symshen_4lazyderef), V4439, V4442)


tmp16230 := Call(__e, tmp16225, tmp16229)


ifres16224 = tmp16230


} else {
ifres16224 = False


}

__e.TailApply(tmp16187, ifres16224)
return


}, 7)

__e.TailApply(ns2_1set, symshen_4join, tmp16186)
return




}, 0)

