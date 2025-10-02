package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp13918 := MakeNative(func(__e *ControlFlow) {
V4484 := __e.Get(1)
_ = V4484
V4485 := __e.Get(2)
_ = V4485
tmp13919 := MakeNative(func(__e *ControlFlow) {
W4486 := __e.Get(1)
_ = W4486
tmp13920 := MakeNative(func(__e *ControlFlow) {
W4487 := __e.Get(1)
_ = W4487
tmp13921 := MakeNative(func(__e *ControlFlow) {
W4488 := __e.Get(1)
_ = W4488
tmp13922 := MakeNative(func(__e *ControlFlow) {
Z4489 := __e.Get(1)
_ = Z4489
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4490 := __e.Get(1)
_ = Z4490
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4491 := __e.Get(1)
_ = Z4491
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4492 := __e.Get(1)
_ = Z4492
tmp13923 := MakeNative(func(__e *ControlFlow) {
W4493 := __e.Get(1)
_ = W4493
tmp13924 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13924

tmp13925 := Call(__e, PrimFunc(symshen_4deref), W4486, Z4489)


tmp13926 := Call(__e, PrimFunc(symreceive), tmp13925)


tmp13927 := Call(__e, PrimFunc(symshen_4deref), W4487, Z4489)


tmp13928 := Call(__e, PrimFunc(symreceive), tmp13927)


tmp13929 := MakeNative(func(__e *ControlFlow) {
tmp13930 := Call(__e, PrimFunc(symshen_4deref), W4488, Z4489)


tmp13931 := Call(__e, PrimFunc(symreceive), tmp13930)


tmp13932 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), W4493, Z4489, Z4490, Z4491, Z4492)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp13931, W4493, Z4489, Z4490, Z4491, tmp13932)
return


}, 0)

tmp13933 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp13926, tmp13928, W4493, Z4489, Z4490, Z4491, tmp13929)


__e.TailApply(PrimFunc(symshen_4gc), Z4489, tmp13933)
return


}, 1)

tmp13934 := Call(__e, PrimFunc(symshen_4newpv), Z4489)


__e.TailApply(tmp13923, tmp13934)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp13935 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp13936 := Call(__e, tmp13922, tmp13935)


tmp13937 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp13938 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp13937)


tmp13939 := Call(__e, PrimFunc(sym_8v), True, tmp13938)


tmp13940 := Call(__e, tmp13936, tmp13939)


tmp13941 := Call(__e, tmp13940, MakeNumber(0))


tmp13942 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp13941, tmp13942)
return


}, 1)

tmp13943 := Call(__e, PrimFunc(symshen_4curry), V4484)


__e.TailApply(tmp13921, tmp13943)
return


}, 1)

tmp13944 := Call(__e, PrimFunc(symshen_4rectify_1type), V4485)


__e.TailApply(tmp13920, tmp13944)
return


}, 1)

tmp13945 := Call(__e, PrimFunc(symshen_4extract_1vars), V4485)


__e.TailApply(tmp13919, tmp13945)
return


}, 2)

tmp13946 := Call(__e, ns2_1set, symshen_4typecheck, tmp13918)


_ = tmp13946

tmp13947 := MakeNative(func(__e *ControlFlow) {
V4494 := __e.Get(1)
_ = V4494
V4495 := __e.Get(2)
_ = V4495
V4496 := __e.Get(3)
_ = V4496
V4497 := __e.Get(4)
_ = V4497
V4498 := __e.Get(5)
_ = V4498
V4499 := __e.Get(6)
_ = V4499
V4500 := __e.Get(7)
_ = V4500
tmp13948 := MakeNative(func(__e *ControlFlow) {
W4501 := __e.Get(1)
_ = W4501
tmp13966 := PrimEqual(W4501, False)

if True == tmp13966 {
tmp13964 := Call(__e, PrimFunc(symshen_4unlocked_2), V4498)


if True == tmp13964 {
tmp13949 := MakeNative(func(__e *ControlFlow) {
W4503 := __e.Get(1)
_ = W4503
tmp13961 := PrimIsPair(W4503)

if True == tmp13961 {
tmp13950 := MakeNative(func(__e *ControlFlow) {
W4504 := __e.Get(1)
_ = W4504
tmp13951 := MakeNative(func(__e *ControlFlow) {
W4505 := __e.Get(1)
_ = W4505
tmp13952 := MakeNative(func(__e *ControlFlow) {
W4506 := __e.Get(1)
_ = W4506
tmp13953 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13953

tmp13954 := Call(__e, PrimFunc(symshen_4deref), W4506, V4497)


tmp13955 := Call(__e, PrimFunc(symsubst), tmp13954, W4504, V4495)


tmp13956 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W4505, tmp13955, V4496, V4497, V4498, V4499, V4500)


__e.TailApply(PrimFunc(symshen_4gc), V4497, tmp13956)
return


}, 1)

tmp13957 := Call(__e, PrimFunc(symshen_4newpv), V4497)


__e.TailApply(tmp13952, tmp13957)
return


}, 1)

tmp13958 := PrimTail(W4503)

__e.TailApply(tmp13951, tmp13958)
return


}, 1)

tmp13959 := PrimHead(W4503)

__e.TailApply(tmp13950, tmp13959)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13962 := Call(__e, PrimFunc(symshen_4lazyderef), V4494, V4497)


__e.TailApply(tmp13949, tmp13962)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4501)
return
}


}, 1)

tmp13974 := Call(__e, PrimFunc(symshen_4unlocked_2), V4498)


var ifres13967 Obj

if True == tmp13974 {
tmp13968 := MakeNative(func(__e *ControlFlow) {
W4502 := __e.Get(1)
_ = W4502
tmp13971 := PrimEqual(W4502, Nil)

if True == tmp13971 {
tmp13969 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13969

__e.TailApply(PrimFunc(symis_b), V4495, V4496, V4497, V4498, V4499, V4500)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13972 := Call(__e, PrimFunc(symshen_4lazyderef), V4494, V4497)


tmp13973 := Call(__e, tmp13968, tmp13972)


ifres13967 = tmp13973


} else {
ifres13967 = False


}

__e.TailApply(tmp13948, ifres13967)
return


}, 7)

tmp13975 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp13947)


_ = tmp13975

tmp13976 := MakeNative(func(__e *ControlFlow) {
V4507 := __e.Get(1)
_ = V4507
V4508 := __e.Get(2)
_ = V4508
V4509 := __e.Get(3)
_ = V4509
V4510 := __e.Get(4)
_ = V4510
V4511 := __e.Get(5)
_ = V4511
V4512 := __e.Get(6)
_ = V4512
tmp13977 := MakeNative(func(__e *ControlFlow) {
W4513 := __e.Get(1)
_ = W4513
tmp13978 := MakeNative(func(__e *ControlFlow) {
W4514 := __e.Get(1)
_ = W4514
tmp13991 := PrimEqual(W4514, False)

if True == tmp13991 {
tmp13979 := MakeNative(func(__e *ControlFlow) {
W4520 := __e.Get(1)
_ = W4520
tmp13981 := PrimEqual(W4520, False)

if True == tmp13981 {
__e.TailApply(PrimFunc(symshen_4unlock), V4510, W4513)
return
} else {
__e.Return(W4520)
return
}


}, 1)

tmp13989 := Call(__e, PrimFunc(symshen_4unlocked_2), V4510)


var ifres13982 Obj

if True == tmp13989 {
tmp13983 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13983

tmp13984 := PrimIntern(MakeString(":"))

tmp13985 := PrimCons(V4508, Nil)

tmp13986 := PrimCons(tmp13984, tmp13985)

tmp13987 := PrimCons(V4507, tmp13986)

tmp13988 := Call(__e, PrimFunc(symshen_4system_1S), tmp13987, Nil, V4509, V4510, W4513, V4512)


ifres13982 = tmp13988


} else {
ifres13982 = False


}

__e.TailApply(tmp13979, ifres13982)
return


} else {
__e.Return(W4514)
return
}


}, 1)

tmp14020 := Call(__e, PrimFunc(symshen_4unlocked_2), V4510)


var ifres13992 Obj

if True == tmp14020 {
tmp13993 := MakeNative(func(__e *ControlFlow) {
W4515 := __e.Get(1)
_ = W4515
tmp14017 := PrimIsPair(W4515)

if True == tmp14017 {
tmp13994 := MakeNative(func(__e *ControlFlow) {
W4516 := __e.Get(1)
_ = W4516
tmp14013 := PrimEqual(W4516, symdefine)

if True == tmp14013 {
tmp13995 := MakeNative(func(__e *ControlFlow) {
W4517 := __e.Get(1)
_ = W4517
tmp14009 := PrimIsPair(W4517)

if True == tmp14009 {
tmp13996 := MakeNative(func(__e *ControlFlow) {
W4518 := __e.Get(1)
_ = W4518
tmp13997 := MakeNative(func(__e *ControlFlow) {
W4519 := __e.Get(1)
_ = W4519
tmp13998 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13998

tmp13999 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14000 := MakeNative(func(__e *ControlFlow) {
tmp14001 := MakeNative(func(__e *ControlFlow) {
tmp14002 := PrimValue(symshen_4_dspy_d)

tmp14003 := MakeNative(func(__e *ControlFlow) {
tmp14004 := PrimCons(W4518, W4519)

tmp14005 := PrimCons(symdefine, tmp14004)

__e.TailApply(PrimFunc(symshen_4t_d), tmp14005, V4508, V4509, V4510, W4513, V4512)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14002, W4518, V4509, V4510, W4513, tmp14003)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4509, V4510, W4513, tmp14001)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp13999, V4509, V4510, W4513, tmp14000)
return


}, 1)

tmp14006 := PrimTail(W4517)

__e.TailApply(tmp13997, tmp14006)
return


}, 1)

tmp14007 := PrimHead(W4517)

__e.TailApply(tmp13996, tmp14007)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14010 := PrimTail(W4515)

tmp14011 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14010, V4509)


__e.TailApply(tmp13995, tmp14011)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14014 := PrimHead(W4515)

tmp14015 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14014, V4509)


__e.TailApply(tmp13994, tmp14015)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14018 := Call(__e, PrimFunc(symshen_4lazyderef), V4507, V4509)


tmp14019 := Call(__e, tmp13993, tmp14018)


ifres13992 = tmp14019


} else {
ifres13992 = False


}

__e.TailApply(tmp13978, ifres13992)
return


}, 1)

tmp14021 := PrimNumberAdd(V4511, MakeNumber(1))

__e.TailApply(tmp13977, tmp14021)
return


}, 6)

tmp14022 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp13976)


_ = tmp14022

tmp14023 := MakeNative(func(__e *ControlFlow) {
V4521 := __e.Get(1)
_ = V4521
V4522 := __e.Get(2)
_ = V4522
V4523 := __e.Get(3)
_ = V4523
V4524 := __e.Get(4)
_ = V4524
V4525 := __e.Get(5)
_ = V4525
V4526 := __e.Get(6)
_ = V4526
tmp14024 := MakeNative(func(__e *ControlFlow) {
W4527 := __e.Get(1)
_ = W4527
tmp14041 := PrimEqual(W4527, False)

if True == tmp14041 {
tmp14039 := Call(__e, PrimFunc(symshen_4unlocked_2), V4524)


if True == tmp14039 {
tmp14025 := MakeNative(func(__e *ControlFlow) {
W4529 := __e.Get(1)
_ = W4529
tmp14036 := PrimEqual(W4529, True)

if True == tmp14036 {
tmp14026 := MakeNative(func(__e *ControlFlow) {
W4530 := __e.Get(1)
_ = W4530
tmp14027 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14027

tmp14028 := Call(__e, PrimFunc(symshen_4deref), V4522, V4523)


tmp14029 := Call(__e, PrimFunc(symshen_4app), tmp14028, MakeString(")\n"), symshen_4a)


tmp14030 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp14029)

tmp14031 := Call(__e, PrimFunc(symstoutput))


tmp14032 := Call(__e, PrimFunc(sympr), tmp14030, tmp14031)


tmp14033 := Call(__e, PrimFunc(symis), W4530, tmp14032, V4523, V4524, V4525, V4526)


__e.TailApply(PrimFunc(symshen_4gc), V4523, tmp14033)
return


}, 1)

tmp14034 := Call(__e, PrimFunc(symshen_4newpv), V4523)


__e.TailApply(tmp14026, tmp14034)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14037 := Call(__e, PrimFunc(symshen_4lazyderef), V4521, V4523)


__e.TailApply(tmp14025, tmp14037)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4527)
return
}


}, 1)

tmp14049 := Call(__e, PrimFunc(symshen_4unlocked_2), V4524)


var ifres14042 Obj

if True == tmp14049 {
tmp14043 := MakeNative(func(__e *ControlFlow) {
W4528 := __e.Get(1)
_ = W4528
tmp14046 := PrimEqual(W4528, False)

if True == tmp14046 {
tmp14044 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14044

__e.TailApply(PrimFunc(symthaw), V4526)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14047 := Call(__e, PrimFunc(symshen_4lazyderef), V4521, V4523)


tmp14048 := Call(__e, tmp14043, tmp14047)


ifres14042 = tmp14048


} else {
ifres14042 = False


}

__e.TailApply(tmp14024, ifres14042)
return


}, 6)

tmp14050 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14023)


_ = tmp14050

tmp14051 := MakeNative(func(__e *ControlFlow) {
V4531 := __e.Get(1)
_ = V4531
tmp14052 := Call(__e, PrimFunc(symshen_4curry_1type), V4531)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14052)
return


}, 1)

tmp14053 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14051)


_ = tmp14053

tmp14054 := MakeNative(func(__e *ControlFlow) {
V4532 := __e.Get(1)
_ = V4532
tmp14055 := MakeNative(func(__e *ControlFlow) {
tmp14056 := MakeNative(func(__e *ControlFlow) {
W4533 := __e.Get(1)
_ = W4533
tmp14058 := PrimEqual(W4533, V4532)

if True == tmp14058 {
__e.Return(V4532)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), W4533)
return
}


}, 1)

tmp14059 := MakeNative(func(__e *ControlFlow) {
Z4534 := __e.Get(1)
_ = Z4534
__e.TailApply(PrimFunc(symshen_4demod), Z4534)
return
}, 1)

tmp14060 := Call(__e, PrimFunc(symshen_4walk), tmp14059, V4532)


__e.TailApply(tmp14056, tmp14060)
return


}, 0)

tmp14061 := MakeNative(func(__e *ControlFlow) {
Z4535 := __e.Get(1)
_ = Z4535
__e.Return(V4532)
return
}, 1)

__e.TailApply(try_1catch, tmp14055, tmp14061)
return


}, 1)

tmp14062 := Call(__e, ns2_1set, symshen_4demodulate, tmp14054)


_ = tmp14062

tmp14063 := MakeNative(func(__e *ControlFlow) {
V4536 := __e.Get(1)
_ = V4536
tmp14187 := PrimIsPair(V4536)

var ifres14160 Obj

if True == tmp14187 {
tmp14185 := PrimTail(V4536)

tmp14186 := PrimIsPair(tmp14185)

var ifres14162 Obj

if True == tmp14186 {
tmp14182 := PrimTail(V4536)

tmp14183 := PrimHead(tmp14182)

tmp14184 := PrimEqual(sym_1_1_6, tmp14183)

var ifres14164 Obj

if True == tmp14184 {
tmp14179 := PrimTail(V4536)

tmp14180 := PrimTail(tmp14179)

tmp14181 := PrimIsPair(tmp14180)

var ifres14166 Obj

if True == tmp14181 {
tmp14175 := PrimTail(V4536)

tmp14176 := PrimTail(tmp14175)

tmp14177 := PrimTail(tmp14176)

tmp14178 := PrimIsPair(tmp14177)

var ifres14168 Obj

if True == tmp14178 {
tmp14170 := PrimTail(V4536)

tmp14171 := PrimTail(tmp14170)

tmp14172 := PrimTail(tmp14171)

tmp14173 := PrimHead(tmp14172)

tmp14174 := PrimEqual(sym_1_1_6, tmp14173)

var ifres14169 Obj

if True == tmp14174 {
ifres14169 = True


} else {
ifres14169 = False


}

ifres14168 = ifres14169


} else {
ifres14168 = False


}

var ifres14167 Obj

if True == ifres14168 {
ifres14167 = True


} else {
ifres14167 = False


}

ifres14166 = ifres14167


} else {
ifres14166 = False


}

var ifres14165 Obj

if True == ifres14166 {
ifres14165 = True


} else {
ifres14165 = False


}

ifres14164 = ifres14165


} else {
ifres14164 = False


}

var ifres14163 Obj

if True == ifres14164 {
ifres14163 = True


} else {
ifres14163 = False


}

ifres14162 = ifres14163


} else {
ifres14162 = False


}

var ifres14161 Obj

if True == ifres14162 {
ifres14161 = True


} else {
ifres14161 = False


}

ifres14160 = ifres14161


} else {
ifres14160 = False


}

if True == ifres14160 {
tmp14064 := PrimHead(V4536)

tmp14065 := PrimTail(V4536)

tmp14066 := PrimTail(tmp14065)

tmp14067 := PrimCons(tmp14066, Nil)

tmp14068 := PrimCons(sym_1_1_6, tmp14067)

tmp14069 := PrimCons(tmp14064, tmp14068)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14069)
return


} else {
tmp14158 := PrimIsPair(V4536)

var ifres14118 Obj

if True == tmp14158 {
tmp14156 := PrimHead(V4536)

tmp14157 := PrimIsPair(tmp14156)

var ifres14120 Obj

if True == tmp14157 {
tmp14153 := PrimHead(V4536)

tmp14154 := PrimHead(tmp14153)

tmp14155 := PrimEqual(symlist, tmp14154)

var ifres14122 Obj

if True == tmp14155 {
tmp14150 := PrimHead(V4536)

tmp14151 := PrimTail(tmp14150)

tmp14152 := PrimIsPair(tmp14151)

var ifres14124 Obj

if True == tmp14152 {
tmp14146 := PrimHead(V4536)

tmp14147 := PrimTail(tmp14146)

tmp14148 := PrimTail(tmp14147)

tmp14149 := PrimEqual(Nil, tmp14148)

var ifres14126 Obj

if True == tmp14149 {
tmp14144 := PrimTail(V4536)

tmp14145 := PrimIsPair(tmp14144)

var ifres14128 Obj

if True == tmp14145 {
tmp14141 := PrimTail(V4536)

tmp14142 := PrimHead(tmp14141)

tmp14143 := PrimEqual(sym_a_a_6, tmp14142)

var ifres14130 Obj

if True == tmp14143 {
tmp14138 := PrimTail(V4536)

tmp14139 := PrimTail(tmp14138)

tmp14140 := PrimIsPair(tmp14139)

var ifres14132 Obj

if True == tmp14140 {
tmp14134 := PrimTail(V4536)

tmp14135 := PrimTail(tmp14134)

tmp14136 := PrimTail(tmp14135)

tmp14137 := PrimEqual(Nil, tmp14136)

var ifres14133 Obj

if True == tmp14137 {
ifres14133 = True


} else {
ifres14133 = False


}

ifres14132 = ifres14133


} else {
ifres14132 = False


}

var ifres14131 Obj

if True == ifres14132 {
ifres14131 = True


} else {
ifres14131 = False


}

ifres14130 = ifres14131


} else {
ifres14130 = False


}

var ifres14129 Obj

if True == ifres14130 {
ifres14129 = True


} else {
ifres14129 = False


}

ifres14128 = ifres14129


} else {
ifres14128 = False


}

var ifres14127 Obj

if True == ifres14128 {
ifres14127 = True


} else {
ifres14127 = False


}

ifres14126 = ifres14127


} else {
ifres14126 = False


}

var ifres14125 Obj

if True == ifres14126 {
ifres14125 = True


} else {
ifres14125 = False


}

ifres14124 = ifres14125


} else {
ifres14124 = False


}

var ifres14123 Obj

if True == ifres14124 {
ifres14123 = True


} else {
ifres14123 = False


}

ifres14122 = ifres14123


} else {
ifres14122 = False


}

var ifres14121 Obj

if True == ifres14122 {
ifres14121 = True


} else {
ifres14121 = False


}

ifres14120 = ifres14121


} else {
ifres14120 = False


}

var ifres14119 Obj

if True == ifres14120 {
ifres14119 = True


} else {
ifres14119 = False


}

ifres14118 = ifres14119


} else {
ifres14118 = False


}

if True == ifres14118 {
tmp14070 := PrimHead(V4536)

tmp14071 := PrimHead(V4536)

tmp14072 := PrimTail(V4536)

tmp14073 := PrimTail(tmp14072)

tmp14074 := PrimCons(tmp14071, tmp14073)

tmp14075 := PrimCons(symstr, tmp14074)

tmp14076 := PrimCons(tmp14075, Nil)

tmp14077 := PrimCons(sym_1_1_6, tmp14076)

tmp14078 := PrimCons(tmp14070, tmp14077)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14078)
return


} else {
tmp14116 := PrimIsPair(V4536)

var ifres14089 Obj

if True == tmp14116 {
tmp14114 := PrimTail(V4536)

tmp14115 := PrimIsPair(tmp14114)

var ifres14091 Obj

if True == tmp14115 {
tmp14111 := PrimTail(V4536)

tmp14112 := PrimHead(tmp14111)

tmp14113 := PrimEqual(sym_d, tmp14112)

var ifres14093 Obj

if True == tmp14113 {
tmp14108 := PrimTail(V4536)

tmp14109 := PrimTail(tmp14108)

tmp14110 := PrimIsPair(tmp14109)

var ifres14095 Obj

if True == tmp14110 {
tmp14104 := PrimTail(V4536)

tmp14105 := PrimTail(tmp14104)

tmp14106 := PrimTail(tmp14105)

tmp14107 := PrimIsPair(tmp14106)

var ifres14097 Obj

if True == tmp14107 {
tmp14099 := PrimTail(V4536)

tmp14100 := PrimTail(tmp14099)

tmp14101 := PrimTail(tmp14100)

tmp14102 := PrimHead(tmp14101)

tmp14103 := PrimEqual(sym_d, tmp14102)

var ifres14098 Obj

if True == tmp14103 {
ifres14098 = True


} else {
ifres14098 = False


}

ifres14097 = ifres14098


} else {
ifres14097 = False


}

var ifres14096 Obj

if True == ifres14097 {
ifres14096 = True


} else {
ifres14096 = False


}

ifres14095 = ifres14096


} else {
ifres14095 = False


}

var ifres14094 Obj

if True == ifres14095 {
ifres14094 = True


} else {
ifres14094 = False


}

ifres14093 = ifres14094


} else {
ifres14093 = False


}

var ifres14092 Obj

if True == ifres14093 {
ifres14092 = True


} else {
ifres14092 = False


}

ifres14091 = ifres14092


} else {
ifres14091 = False


}

var ifres14090 Obj

if True == ifres14091 {
ifres14090 = True


} else {
ifres14090 = False


}

ifres14089 = ifres14090


} else {
ifres14089 = False


}

if True == ifres14089 {
tmp14079 := PrimHead(V4536)

tmp14080 := PrimTail(V4536)

tmp14081 := PrimTail(tmp14080)

tmp14082 := PrimCons(tmp14081, Nil)

tmp14083 := PrimCons(sym_d, tmp14082)

tmp14084 := PrimCons(tmp14079, tmp14083)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14084)
return


} else {
tmp14087 := PrimIsPair(V4536)

if True == tmp14087 {
tmp14085 := MakeNative(func(__e *ControlFlow) {
Z4537 := __e.Get(1)
_ = Z4537
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4537)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14085, V4536)
return


} else {
__e.Return(V4536)
return
}


}


}


}


}, 1)

tmp14188 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14063)


_ = tmp14188

tmp14189 := MakeNative(func(__e *ControlFlow) {
V4538 := __e.Get(1)
_ = V4538
tmp14307 := PrimIsPair(V4538)

var ifres14299 Obj

if True == tmp14307 {
tmp14305 := PrimHead(V4538)

tmp14306 := PrimEqual(symdefine, tmp14305)

var ifres14301 Obj

if True == tmp14306 {
tmp14303 := PrimTail(V4538)

tmp14304 := PrimIsPair(tmp14303)

var ifres14302 Obj

if True == tmp14304 {
ifres14302 = True


} else {
ifres14302 = False


}

ifres14301 = ifres14302


} else {
ifres14301 = False


}

var ifres14300 Obj

if True == ifres14301 {
ifres14300 = True


} else {
ifres14300 = False


}

ifres14299 = ifres14300


} else {
ifres14299 = False


}

if True == ifres14299 {
__e.Return(V4538)
return
} else {
tmp14297 := PrimIsPair(V4538)

var ifres14278 Obj

if True == tmp14297 {
tmp14295 := PrimHead(V4538)

tmp14296 := PrimEqual(symtype, tmp14295)

var ifres14280 Obj

if True == tmp14296 {
tmp14293 := PrimTail(V4538)

tmp14294 := PrimIsPair(tmp14293)

var ifres14282 Obj

if True == tmp14294 {
tmp14290 := PrimTail(V4538)

tmp14291 := PrimTail(tmp14290)

tmp14292 := PrimIsPair(tmp14291)

var ifres14284 Obj

if True == tmp14292 {
tmp14286 := PrimTail(V4538)

tmp14287 := PrimTail(tmp14286)

tmp14288 := PrimTail(tmp14287)

tmp14289 := PrimEqual(Nil, tmp14288)

var ifres14285 Obj

if True == tmp14289 {
ifres14285 = True


} else {
ifres14285 = False


}

ifres14284 = ifres14285


} else {
ifres14284 = False


}

var ifres14283 Obj

if True == ifres14284 {
ifres14283 = True


} else {
ifres14283 = False


}

ifres14282 = ifres14283


} else {
ifres14282 = False


}

var ifres14281 Obj

if True == ifres14282 {
ifres14281 = True


} else {
ifres14281 = False


}

ifres14280 = ifres14281


} else {
ifres14280 = False


}

var ifres14279 Obj

if True == ifres14280 {
ifres14279 = True


} else {
ifres14279 = False


}

ifres14278 = ifres14279


} else {
ifres14278 = False


}

if True == ifres14278 {
tmp14190 := PrimTail(V4538)

tmp14191 := PrimHead(tmp14190)

tmp14192 := Call(__e, PrimFunc(symshen_4curry), tmp14191)


tmp14193 := PrimTail(V4538)

tmp14194 := PrimTail(tmp14193)

tmp14195 := PrimCons(tmp14192, tmp14194)

__e.Return(PrimCons(symtype, tmp14195))
return


} else {
tmp14276 := PrimIsPair(V4538)

var ifres14257 Obj

if True == tmp14276 {
tmp14274 := PrimHead(V4538)

tmp14275 := PrimEqual(syminput_7, tmp14274)

var ifres14259 Obj

if True == tmp14275 {
tmp14272 := PrimTail(V4538)

tmp14273 := PrimIsPair(tmp14272)

var ifres14261 Obj

if True == tmp14273 {
tmp14269 := PrimTail(V4538)

tmp14270 := PrimTail(tmp14269)

tmp14271 := PrimIsPair(tmp14270)

var ifres14263 Obj

if True == tmp14271 {
tmp14265 := PrimTail(V4538)

tmp14266 := PrimTail(tmp14265)

tmp14267 := PrimTail(tmp14266)

tmp14268 := PrimEqual(Nil, tmp14267)

var ifres14264 Obj

if True == tmp14268 {
ifres14264 = True


} else {
ifres14264 = False


}

ifres14263 = ifres14264


} else {
ifres14263 = False


}

var ifres14262 Obj

if True == ifres14263 {
ifres14262 = True


} else {
ifres14262 = False


}

ifres14261 = ifres14262


} else {
ifres14261 = False


}

var ifres14260 Obj

if True == ifres14261 {
ifres14260 = True


} else {
ifres14260 = False


}

ifres14259 = ifres14260


} else {
ifres14259 = False


}

var ifres14258 Obj

if True == ifres14259 {
ifres14258 = True


} else {
ifres14258 = False


}

ifres14257 = ifres14258


} else {
ifres14257 = False


}

if True == ifres14257 {
tmp14196 := PrimTail(V4538)

tmp14197 := PrimHead(tmp14196)

tmp14198 := PrimTail(V4538)

tmp14199 := PrimTail(tmp14198)

tmp14200 := PrimHead(tmp14199)

tmp14201 := Call(__e, PrimFunc(symshen_4curry), tmp14200)


tmp14202 := PrimCons(tmp14201, Nil)

tmp14203 := PrimCons(tmp14197, tmp14202)

__e.Return(PrimCons(syminput_7, tmp14203))
return


} else {
tmp14255 := PrimIsPair(V4538)

var ifres14251 Obj

if True == tmp14255 {
tmp14253 := PrimHead(V4538)

tmp14254 := Call(__e, PrimFunc(symshen_4special_2), tmp14253)


var ifres14252 Obj

if True == tmp14254 {
ifres14252 = True


} else {
ifres14252 = False


}

ifres14251 = ifres14252


} else {
ifres14251 = False


}

if True == ifres14251 {
tmp14204 := PrimHead(V4538)

tmp14205 := MakeNative(func(__e *ControlFlow) {
Z4539 := __e.Get(1)
_ = Z4539
__e.TailApply(PrimFunc(symshen_4curry), Z4539)
return
}, 1)

tmp14206 := PrimTail(V4538)

tmp14207 := Call(__e, PrimFunc(symmap), tmp14205, tmp14206)


__e.Return(PrimCons(tmp14204, tmp14207))
return


} else {
tmp14249 := PrimIsPair(V4538)

var ifres14245 Obj

if True == tmp14249 {
tmp14247 := PrimHead(V4538)

tmp14248 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14247)


var ifres14246 Obj

if True == tmp14248 {
ifres14246 = True


} else {
ifres14246 = False


}

ifres14245 = ifres14246


} else {
ifres14245 = False


}

if True == ifres14245 {
__e.Return(V4538)
return
} else {
tmp14243 := PrimIsPair(V4538)

var ifres14234 Obj

if True == tmp14243 {
tmp14241 := PrimTail(V4538)

tmp14242 := PrimIsPair(tmp14241)

var ifres14236 Obj

if True == tmp14242 {
tmp14238 := PrimTail(V4538)

tmp14239 := PrimTail(tmp14238)

tmp14240 := PrimIsPair(tmp14239)

var ifres14237 Obj

if True == tmp14240 {
ifres14237 = True


} else {
ifres14237 = False


}

ifres14236 = ifres14237


} else {
ifres14236 = False


}

var ifres14235 Obj

if True == ifres14236 {
ifres14235 = True


} else {
ifres14235 = False


}

ifres14234 = ifres14235


} else {
ifres14234 = False


}

if True == ifres14234 {
tmp14208 := PrimHead(V4538)

tmp14209 := PrimTail(V4538)

tmp14210 := PrimHead(tmp14209)

tmp14211 := PrimCons(tmp14210, Nil)

tmp14212 := PrimCons(tmp14208, tmp14211)

tmp14213 := PrimTail(V4538)

tmp14214 := PrimTail(tmp14213)

tmp14215 := PrimCons(tmp14212, tmp14214)

__e.TailApply(PrimFunc(symshen_4curry), tmp14215)
return


} else {
tmp14232 := PrimIsPair(V4538)

var ifres14223 Obj

if True == tmp14232 {
tmp14230 := PrimTail(V4538)

tmp14231 := PrimIsPair(tmp14230)

var ifres14225 Obj

if True == tmp14231 {
tmp14227 := PrimTail(V4538)

tmp14228 := PrimTail(tmp14227)

tmp14229 := PrimEqual(Nil, tmp14228)

var ifres14226 Obj

if True == tmp14229 {
ifres14226 = True


} else {
ifres14226 = False


}

ifres14225 = ifres14226


} else {
ifres14225 = False


}

var ifres14224 Obj

if True == ifres14225 {
ifres14224 = True


} else {
ifres14224 = False


}

ifres14223 = ifres14224


} else {
ifres14223 = False


}

if True == ifres14223 {
tmp14216 := PrimHead(V4538)

tmp14217 := Call(__e, PrimFunc(symshen_4curry), tmp14216)


tmp14218 := PrimTail(V4538)

tmp14219 := PrimHead(tmp14218)

tmp14220 := Call(__e, PrimFunc(symshen_4curry), tmp14219)


tmp14221 := PrimCons(tmp14220, Nil)

__e.Return(PrimCons(tmp14217, tmp14221))
return


} else {
__e.Return(V4538)
return
}


}


}


}


}


}


}


}, 1)

tmp14308 := Call(__e, ns2_1set, symshen_4curry, tmp14189)


_ = tmp14308

tmp14309 := MakeNative(func(__e *ControlFlow) {
V4540 := __e.Get(1)
_ = V4540
tmp14310 := PrimValue(symshen_4_dspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4540, tmp14310)
return


}, 1)

tmp14311 := Call(__e, ns2_1set, symshen_4special_2, tmp14309)


_ = tmp14311

tmp14312 := MakeNative(func(__e *ControlFlow) {
V4541 := __e.Get(1)
_ = V4541
tmp14313 := PrimValue(symshen_4_dextraspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4541, tmp14313)
return


}, 1)

tmp14314 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14312)


_ = tmp14314

tmp14315 := MakeNative(func(__e *ControlFlow) {
V4542 := __e.Get(1)
_ = V4542
V4543 := __e.Get(2)
_ = V4543
V4544 := __e.Get(3)
_ = V4544
V4545 := __e.Get(4)
_ = V4545
V4546 := __e.Get(5)
_ = V4546
V4547 := __e.Get(6)
_ = V4547
tmp14316 := MakeNative(func(__e *ControlFlow) {
W4548 := __e.Get(1)
_ = W4548
tmp14317 := MakeNative(func(__e *ControlFlow) {
W4549 := __e.Get(1)
_ = W4549
tmp14375 := PrimEqual(W4549, False)

if True == tmp14375 {
tmp14318 := MakeNative(func(__e *ControlFlow) {
W4550 := __e.Get(1)
_ = W4550
tmp14337 := PrimEqual(W4550, False)

if True == tmp14337 {
tmp14319 := MakeNative(func(__e *ControlFlow) {
W4558 := __e.Get(1)
_ = W4558
tmp14329 := PrimEqual(W4558, False)

if True == tmp14329 {
tmp14320 := MakeNative(func(__e *ControlFlow) {
W4559 := __e.Get(1)
_ = W4559
tmp14322 := PrimEqual(W4559, False)

if True == tmp14322 {
__e.TailApply(PrimFunc(symshen_4unlock), V4545, W4548)
return
} else {
__e.Return(W4559)
return
}


}, 1)

tmp14327 := Call(__e, PrimFunc(symshen_4unlocked_2), V4545)


var ifres14323 Obj

if True == tmp14327 {
tmp14324 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14324

tmp14325 := PrimValue(symshen_4_ddatatypes_d)

tmp14326 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4542, V4543, tmp14325, V4544, V4545, W4548, V4547)


ifres14323 = tmp14326


} else {
ifres14323 = False


}

__e.TailApply(tmp14320, ifres14323)
return


} else {
__e.Return(W4558)
return
}


}, 1)

tmp14335 := Call(__e, PrimFunc(symshen_4unlocked_2), V4545)


var ifres14330 Obj

if True == tmp14335 {
tmp14331 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14331

tmp14332 := PrimValue(symshen_4_dspy_d)

tmp14333 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4542, V4543, V4544, V4545, W4548, V4547)
return
}, 0)

tmp14334 := Call(__e, PrimFunc(symwhen), tmp14332, V4544, V4545, W4548, tmp14333)


ifres14330 = tmp14334


} else {
ifres14330 = False


}

__e.TailApply(tmp14319, ifres14330)
return


} else {
__e.Return(W4550)
return
}


}, 1)

tmp14373 := Call(__e, PrimFunc(symshen_4unlocked_2), V4545)


var ifres14338 Obj

if True == tmp14373 {
tmp14339 := MakeNative(func(__e *ControlFlow) {
W4551 := __e.Get(1)
_ = W4551
tmp14370 := PrimIsPair(W4551)

if True == tmp14370 {
tmp14340 := MakeNative(func(__e *ControlFlow) {
W4552 := __e.Get(1)
_ = W4552
tmp14341 := MakeNative(func(__e *ControlFlow) {
W4553 := __e.Get(1)
_ = W4553
tmp14365 := PrimIsPair(W4553)

if True == tmp14365 {
tmp14342 := MakeNative(func(__e *ControlFlow) {
W4554 := __e.Get(1)
_ = W4554
tmp14343 := MakeNative(func(__e *ControlFlow) {
W4555 := __e.Get(1)
_ = W4555
tmp14360 := PrimIsPair(W4555)

if True == tmp14360 {
tmp14344 := MakeNative(func(__e *ControlFlow) {
W4556 := __e.Get(1)
_ = W4556
tmp14345 := MakeNative(func(__e *ControlFlow) {
W4557 := __e.Get(1)
_ = W4557
tmp14355 := PrimEqual(W4557, Nil)

if True == tmp14355 {
tmp14346 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14346

tmp14347 := Call(__e, PrimFunc(symshen_4deref), W4554, V4544)


tmp14348 := PrimIntern(MakeString(":"))

tmp14349 := PrimEqual(tmp14347, tmp14348)

tmp14350 := MakeNative(func(__e *ControlFlow) {
tmp14351 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14352 := MakeNative(func(__e *ControlFlow) {
tmp14353 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4552, W4556, V4543, V4544, V4545, W4548, V4547)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4544, V4545, W4548, tmp14353)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14351, V4544, V4545, W4548, tmp14352)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14349, V4544, V4545, W4548, tmp14350)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14356 := PrimTail(W4555)

tmp14357 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14356, V4544)


__e.TailApply(tmp14345, tmp14357)
return


}, 1)

tmp14358 := PrimHead(W4555)

__e.TailApply(tmp14344, tmp14358)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14361 := PrimTail(W4553)

tmp14362 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14361, V4544)


__e.TailApply(tmp14343, tmp14362)
return


}, 1)

tmp14363 := PrimHead(W4553)

__e.TailApply(tmp14342, tmp14363)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14366 := PrimTail(W4551)

tmp14367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14366, V4544)


__e.TailApply(tmp14341, tmp14367)
return


}, 1)

tmp14368 := PrimHead(W4551)

__e.TailApply(tmp14340, tmp14368)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14371 := Call(__e, PrimFunc(symshen_4lazyderef), V4542, V4544)


tmp14372 := Call(__e, tmp14339, tmp14371)


ifres14338 = tmp14372


} else {
ifres14338 = False


}

__e.TailApply(tmp14318, ifres14338)
return


} else {
__e.Return(W4549)
return
}


}, 1)

tmp14380 := Call(__e, PrimFunc(symshen_4unlocked_2), V4545)


var ifres14376 Obj

if True == tmp14380 {
tmp14377 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14377

tmp14378 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14379 := Call(__e, PrimFunc(symwhen), tmp14378, V4544, V4545, W4548, V4547)


ifres14376 = tmp14379


} else {
ifres14376 = False


}

__e.TailApply(tmp14317, ifres14376)
return


}, 1)

tmp14381 := PrimNumberAdd(V4546, MakeNumber(1))

__e.TailApply(tmp14316, tmp14381)
return


}, 6)

tmp14382 := Call(__e, ns2_1set, symshen_4system_1S, tmp14315)


_ = tmp14382

tmp14383 := MakeNative(func(__e *ControlFlow) {
V4566 := __e.Get(1)
_ = V4566
V4567 := __e.Get(2)
_ = V4567
V4568 := __e.Get(3)
_ = V4568
V4569 := __e.Get(4)
_ = V4569
V4570 := __e.Get(5)
_ = V4570
V4571 := __e.Get(6)
_ = V4571
tmp14384 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14384

tmp14385 := Call(__e, PrimFunc(symshen_4deref), V4566, V4568)


tmp14386 := Call(__e, PrimFunc(symshen_4show_1p), tmp14385)


_ = tmp14386

tmp14387 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp14387

tmp14388 := Call(__e, PrimFunc(symshen_4deref), V4567, V4568)


tmp14389 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14388, MakeNumber(1))


_ = tmp14389

tmp14390 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14390

__e.Return(False)
return


}, 6)

tmp14391 := Call(__e, ns2_1set, symshen_4show, tmp14383)


_ = tmp14391

tmp14392 := MakeNative(func(__e *ControlFlow) {
tmp14393 := MakeNative(func(__e *ControlFlow) {
W4572 := __e.Get(1)
_ = W4572
tmp14395 := PrimEqual(MakeNumber(1), W4572)

var ifres14394 Obj

if True == tmp14395 {
ifres14394 = MakeString("")


} else {
ifres14394 = MakeString("s")


}

tmp14396 := Call(__e, PrimFunc(symshen_4app), ifres14394, MakeString(" \n?- "), symshen_4a)


tmp14397 := PrimStringConcat(MakeString(" inference"), tmp14396)

tmp14398 := Call(__e, PrimFunc(symshen_4app), W4572, tmp14397, symshen_4a)


tmp14399 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp14398)

tmp14400 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14399, tmp14400)
return


}, 1)

tmp14401 := Call(__e, PrimFunc(syminferences))


__e.TailApply(tmp14393, tmp14401)
return


}, 0)

tmp14402 := Call(__e, ns2_1set, symshen_4line, tmp14392)


_ = tmp14402

tmp14403 := MakeNative(func(__e *ControlFlow) {
V4573 := __e.Get(1)
_ = V4573
tmp14435 := PrimIsPair(V4573)

var ifres14414 Obj

if True == tmp14435 {
tmp14433 := PrimTail(V4573)

tmp14434 := PrimIsPair(tmp14433)

var ifres14416 Obj

if True == tmp14434 {
tmp14430 := PrimTail(V4573)

tmp14431 := PrimTail(tmp14430)

tmp14432 := PrimIsPair(tmp14431)

var ifres14418 Obj

if True == tmp14432 {
tmp14426 := PrimTail(V4573)

tmp14427 := PrimTail(tmp14426)

tmp14428 := PrimTail(tmp14427)

tmp14429 := PrimEqual(Nil, tmp14428)

var ifres14420 Obj

if True == tmp14429 {
tmp14422 := PrimTail(V4573)

tmp14423 := PrimHead(tmp14422)

tmp14424 := PrimIntern(MakeString(":"))

tmp14425 := PrimEqual(tmp14423, tmp14424)

var ifres14421 Obj

if True == tmp14425 {
ifres14421 = True


} else {
ifres14421 = False


}

ifres14420 = ifres14421


} else {
ifres14420 = False


}

var ifres14419 Obj

if True == ifres14420 {
ifres14419 = True


} else {
ifres14419 = False


}

ifres14418 = ifres14419


} else {
ifres14418 = False


}

var ifres14417 Obj

if True == ifres14418 {
ifres14417 = True


} else {
ifres14417 = False


}

ifres14416 = ifres14417


} else {
ifres14416 = False


}

var ifres14415 Obj

if True == ifres14416 {
ifres14415 = True


} else {
ifres14415 = False


}

ifres14414 = ifres14415


} else {
ifres14414 = False


}

if True == ifres14414 {
tmp14404 := PrimHead(V4573)

tmp14405 := Call(__e, PrimFunc(symshen_4prterm), tmp14404)


_ = tmp14405

tmp14406 := Call(__e, PrimFunc(symstoutput))


tmp14407 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14406)


_ = tmp14407

tmp14408 := PrimTail(V4573)

tmp14409 := PrimTail(tmp14408)

tmp14410 := PrimHead(tmp14409)

tmp14411 := Call(__e, PrimFunc(symshen_4app), tmp14410, MakeString(""), symshen_4r)


tmp14412 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14411, tmp14412)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4573)
return
}


}, 1)

tmp14436 := Call(__e, ns2_1set, symshen_4show_1p, tmp14403)


_ = tmp14436

tmp14437 := MakeNative(func(__e *ControlFlow) {
V4574 := __e.Get(1)
_ = V4574
tmp14480 := PrimIsPair(V4574)

var ifres14461 Obj

if True == tmp14480 {
tmp14478 := PrimHead(V4574)

tmp14479 := PrimEqual(symcons, tmp14478)

var ifres14463 Obj

if True == tmp14479 {
tmp14476 := PrimTail(V4574)

tmp14477 := PrimIsPair(tmp14476)

var ifres14465 Obj

if True == tmp14477 {
tmp14473 := PrimTail(V4574)

tmp14474 := PrimTail(tmp14473)

tmp14475 := PrimIsPair(tmp14474)

var ifres14467 Obj

if True == tmp14475 {
tmp14469 := PrimTail(V4574)

tmp14470 := PrimTail(tmp14469)

tmp14471 := PrimTail(tmp14470)

tmp14472 := PrimEqual(Nil, tmp14471)

var ifres14468 Obj

if True == tmp14472 {
ifres14468 = True


} else {
ifres14468 = False


}

ifres14467 = ifres14468


} else {
ifres14467 = False


}

var ifres14466 Obj

if True == ifres14467 {
ifres14466 = True


} else {
ifres14466 = False


}

ifres14465 = ifres14466


} else {
ifres14465 = False


}

var ifres14464 Obj

if True == ifres14465 {
ifres14464 = True


} else {
ifres14464 = False


}

ifres14463 = ifres14464


} else {
ifres14463 = False


}

var ifres14462 Obj

if True == ifres14463 {
ifres14462 = True


} else {
ifres14462 = False


}

ifres14461 = ifres14462


} else {
ifres14461 = False


}

if True == ifres14461 {
tmp14438 := Call(__e, PrimFunc(symstoutput))


tmp14439 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14438)


_ = tmp14439

tmp14440 := PrimTail(V4574)

tmp14441 := PrimHead(tmp14440)

tmp14442 := Call(__e, PrimFunc(symshen_4prterm), tmp14441)


_ = tmp14442

tmp14443 := PrimTail(V4574)

tmp14444 := PrimTail(tmp14443)

tmp14445 := PrimHead(tmp14444)

tmp14446 := Call(__e, PrimFunc(symshen_4prtl), tmp14445)


_ = tmp14446

tmp14447 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp14447)
return


} else {
tmp14459 := PrimIsPair(V4574)

if True == tmp14459 {
tmp14448 := Call(__e, PrimFunc(symstoutput))


tmp14449 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14448)


_ = tmp14449

tmp14450 := PrimHead(V4574)

tmp14451 := Call(__e, PrimFunc(symshen_4prterm), tmp14450)


_ = tmp14451

tmp14452 := MakeNative(func(__e *ControlFlow) {
Z4575 := __e.Get(1)
_ = Z4575
tmp14453 := Call(__e, PrimFunc(symstoutput))


tmp14454 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14453)


_ = tmp14454

__e.TailApply(PrimFunc(symshen_4prterm), Z4575)
return


}, 1)

tmp14455 := PrimTail(V4574)

tmp14456 := Call(__e, PrimFunc(symmap), tmp14452, tmp14455)


_ = tmp14456

tmp14457 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14457)
return


} else {
__e.TailApply(PrimFunc(symprint), V4574)
return
}


}


}, 1)

tmp14481 := Call(__e, ns2_1set, symshen_4prterm, tmp14437)


_ = tmp14481

tmp14482 := MakeNative(func(__e *ControlFlow) {
V4576 := __e.Get(1)
_ = V4576
tmp14515 := PrimEqual(Nil, V4576)

if True == tmp14515 {
__e.Return(MakeString(""))
return
} else {
tmp14513 := PrimIsPair(V4576)

var ifres14494 Obj

if True == tmp14513 {
tmp14511 := PrimHead(V4576)

tmp14512 := PrimEqual(symcons, tmp14511)

var ifres14496 Obj

if True == tmp14512 {
tmp14509 := PrimTail(V4576)

tmp14510 := PrimIsPair(tmp14509)

var ifres14498 Obj

if True == tmp14510 {
tmp14506 := PrimTail(V4576)

tmp14507 := PrimTail(tmp14506)

tmp14508 := PrimIsPair(tmp14507)

var ifres14500 Obj

if True == tmp14508 {
tmp14502 := PrimTail(V4576)

tmp14503 := PrimTail(tmp14502)

tmp14504 := PrimTail(tmp14503)

tmp14505 := PrimEqual(Nil, tmp14504)

var ifres14501 Obj

if True == tmp14505 {
ifres14501 = True


} else {
ifres14501 = False


}

ifres14500 = ifres14501


} else {
ifres14500 = False


}

var ifres14499 Obj

if True == ifres14500 {
ifres14499 = True


} else {
ifres14499 = False


}

ifres14498 = ifres14499


} else {
ifres14498 = False


}

var ifres14497 Obj

if True == ifres14498 {
ifres14497 = True


} else {
ifres14497 = False


}

ifres14496 = ifres14497


} else {
ifres14496 = False


}

var ifres14495 Obj

if True == ifres14496 {
ifres14495 = True


} else {
ifres14495 = False


}

ifres14494 = ifres14495


} else {
ifres14494 = False


}

if True == ifres14494 {
tmp14483 := Call(__e, PrimFunc(symstoutput))


tmp14484 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14483)


_ = tmp14484

tmp14485 := PrimTail(V4576)

tmp14486 := PrimHead(tmp14485)

tmp14487 := Call(__e, PrimFunc(symshen_4prterm), tmp14486)


_ = tmp14487

tmp14488 := PrimTail(V4576)

tmp14489 := PrimTail(tmp14488)

tmp14490 := PrimHead(tmp14489)

__e.TailApply(PrimFunc(symshen_4prtl), tmp14490)
return


} else {
tmp14491 := Call(__e, PrimFunc(symstoutput))


tmp14492 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp14491)


_ = tmp14492

__e.TailApply(PrimFunc(symshen_4prterm), V4576)
return


}


}


}, 1)

tmp14516 := Call(__e, ns2_1set, symshen_4prtl, tmp14482)


_ = tmp14516

tmp14517 := MakeNative(func(__e *ControlFlow) {
V4583 := __e.Get(1)
_ = V4583
V4584 := __e.Get(2)
_ = V4584
tmp14530 := PrimEqual(Nil, V4583)

if True == tmp14530 {
tmp14518 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp14518)
return


} else {
tmp14528 := PrimIsPair(V4583)

if True == tmp14528 {
tmp14519 := Call(__e, PrimFunc(symshen_4app), V4584, MakeString(". "), symshen_4a)


tmp14520 := Call(__e, PrimFunc(symstoutput))


tmp14521 := Call(__e, PrimFunc(sympr), tmp14519, tmp14520)


_ = tmp14521

tmp14522 := PrimHead(V4583)

tmp14523 := Call(__e, PrimFunc(symshen_4show_1p), tmp14522)


_ = tmp14523

tmp14524 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp14524

tmp14525 := PrimTail(V4583)

tmp14526 := PrimNumberAdd(V4584, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp14525, tmp14526)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
return
}


}


}, 2)

tmp14531 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp14517)


_ = tmp14531

tmp14532 := MakeNative(func(__e *ControlFlow) {
tmp14533 := MakeNative(func(__e *ControlFlow) {
W4585 := __e.Get(1)
_ = W4585
tmp14535 := PrimEqual(W4585, MakeNumber(94))

if True == tmp14535 {
__e.Return(PrimSimpleError(MakeString("input aborted\n")))
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 1)

tmp14536 := Call(__e, PrimFunc(symstinput))


tmp14537 := PrimReadByte(tmp14536)

__e.TailApply(tmp14533, tmp14537)
return


}, 0)

tmp14538 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp14532)


_ = tmp14538

tmp14539 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp14540 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp14539)


_ = tmp14540

tmp14541 := MakeNative(func(__e *ControlFlow) {
tmp14543 := Call(__e, PrimFunc(syminferences))


tmp14544 := PrimValue(symshen_4_dmaxinferences_d)

tmp14545 := PrimGreatThan(tmp14543, tmp14544)

if True == tmp14545 {
__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
return
} else {
__e.Return(False)
return
}


}, 0)

tmp14546 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp14541)


_ = tmp14546

tmp14547 := MakeNative(func(__e *ControlFlow) {
V4586 := __e.Get(1)
_ = V4586
V4587 := __e.Get(2)
_ = V4587
V4588 := __e.Get(3)
_ = V4588
V4589 := __e.Get(4)
_ = V4589
V4590 := __e.Get(5)
_ = V4590
V4591 := __e.Get(6)
_ = V4591
V4592 := __e.Get(7)
_ = V4592
tmp14548 := MakeNative(func(__e *ControlFlow) {
W4593 := __e.Get(1)
_ = W4593
tmp14549 := MakeNative(func(__e *ControlFlow) {
W4594 := __e.Get(1)
_ = W4594
tmp15463 := PrimEqual(W4594, False)

if True == tmp15463 {
tmp14550 := MakeNative(func(__e *ControlFlow) {
W4595 := __e.Get(1)
_ = W4595
tmp15453 := PrimEqual(W4595, False)

if True == tmp15453 {
tmp14551 := MakeNative(func(__e *ControlFlow) {
W4596 := __e.Get(1)
_ = W4596
tmp15447 := PrimEqual(W4596, False)

if True == tmp15447 {
tmp14552 := MakeNative(func(__e *ControlFlow) {
W4597 := __e.Get(1)
_ = W4597
tmp15428 := PrimEqual(W4597, False)

if True == tmp15428 {
tmp14553 := MakeNative(func(__e *ControlFlow) {
W4601 := __e.Get(1)
_ = W4601
tmp15395 := PrimEqual(W4601, False)

if True == tmp15395 {
tmp14554 := MakeNative(func(__e *ControlFlow) {
W4607 := __e.Get(1)
_ = W4607
tmp15368 := PrimEqual(W4607, False)

if True == tmp15368 {
tmp14555 := MakeNative(func(__e *ControlFlow) {
W4613 := __e.Get(1)
_ = W4613
tmp15333 := PrimEqual(W4613, False)

if True == tmp15333 {
tmp14556 := MakeNative(func(__e *ControlFlow) {
W4620 := __e.Get(1)
_ = W4620
tmp15302 := PrimEqual(W4620, False)

if True == tmp15302 {
tmp14557 := MakeNative(func(__e *ControlFlow) {
W4627 := __e.Get(1)
_ = W4627
tmp15217 := PrimEqual(W4627, False)

if True == tmp15217 {
tmp14558 := MakeNative(func(__e *ControlFlow) {
W4648 := __e.Get(1)
_ = W4648
tmp15111 := PrimEqual(W4648, False)

if True == tmp15111 {
tmp14559 := MakeNative(func(__e *ControlFlow) {
W4676 := __e.Get(1)
_ = W4676
tmp15026 := PrimEqual(W4676, False)

if True == tmp15026 {
tmp14560 := MakeNative(func(__e *ControlFlow) {
W4697 := __e.Get(1)
_ = W4697
tmp14983 := PrimEqual(W4697, False)

if True == tmp14983 {
tmp14561 := MakeNative(func(__e *ControlFlow) {
W4707 := __e.Get(1)
_ = W4707
tmp14859 := PrimEqual(W4707, False)

if True == tmp14859 {
tmp14562 := MakeNative(func(__e *ControlFlow) {
W4737 := __e.Get(1)
_ = W4737
tmp14795 := PrimEqual(W4737, False)

if True == tmp14795 {
tmp14563 := MakeNative(func(__e *ControlFlow) {
W4750 := __e.Get(1)
_ = W4750
tmp14707 := PrimEqual(W4750, False)

if True == tmp14707 {
tmp14564 := MakeNative(func(__e *ControlFlow) {
W4771 := __e.Get(1)
_ = W4771
tmp14669 := PrimEqual(W4771, False)

if True == tmp14669 {
tmp14565 := MakeNative(func(__e *ControlFlow) {
W4779 := __e.Get(1)
_ = W4779
tmp14630 := PrimEqual(W4779, False)

if True == tmp14630 {
tmp14566 := MakeNative(func(__e *ControlFlow) {
W4787 := __e.Get(1)
_ = W4787
tmp14592 := PrimEqual(W4787, False)

if True == tmp14592 {
tmp14567 := MakeNative(func(__e *ControlFlow) {
W4795 := __e.Get(1)
_ = W4795
tmp14581 := PrimEqual(W4795, False)

if True == tmp14581 {
tmp14568 := MakeNative(func(__e *ControlFlow) {
W4797 := __e.Get(1)
_ = W4797
tmp14570 := PrimEqual(W4797, False)

if True == tmp14570 {
__e.TailApply(PrimFunc(symshen_4unlock), V4590, W4593)
return
} else {
__e.Return(W4797)
return
}


}, 1)

tmp14579 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14571 Obj

if True == tmp14579 {
tmp14572 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14572

tmp14573 := PrimIntern(MakeString(":"))

tmp14574 := PrimCons(V4587, Nil)

tmp14575 := PrimCons(tmp14573, tmp14574)

tmp14576 := PrimCons(V4586, tmp14575)

tmp14577 := PrimValue(symshen_4_ddatatypes_d)

tmp14578 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp14576, V4588, tmp14577, V4589, V4590, W4593, V4592)


ifres14571 = tmp14578


} else {
ifres14571 = False


}

__e.TailApply(tmp14568, ifres14571)
return


} else {
__e.Return(W4795)
return
}


}, 1)

tmp14590 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14582 Obj

if True == tmp14590 {
tmp14583 := MakeNative(func(__e *ControlFlow) {
W4796 := __e.Get(1)
_ = W4796
tmp14584 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14584

tmp14585 := MakeNative(func(__e *ControlFlow) {
tmp14586 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4586, V4587, W4796, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4589, V4590, W4593, tmp14586)
return


}, 0)

tmp14587 := Call(__e, PrimFunc(symshen_4l_1rules), V4588, W4796, False, V4589, V4590, W4593, tmp14585)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14587)
return


}, 1)

tmp14588 := Call(__e, PrimFunc(symshen_4newpv), V4589)


tmp14589 := Call(__e, tmp14583, tmp14588)


ifres14582 = tmp14589


} else {
ifres14582 = False


}

__e.TailApply(tmp14567, ifres14582)
return


} else {
__e.Return(W4787)
return
}


}, 1)

tmp14628 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14593 Obj

if True == tmp14628 {
tmp14594 := MakeNative(func(__e *ControlFlow) {
W4788 := __e.Get(1)
_ = W4788
tmp14625 := PrimIsPair(W4788)

if True == tmp14625 {
tmp14595 := MakeNative(func(__e *ControlFlow) {
W4789 := __e.Get(1)
_ = W4789
tmp14621 := PrimEqual(W4789, symset)

if True == tmp14621 {
tmp14596 := MakeNative(func(__e *ControlFlow) {
W4790 := __e.Get(1)
_ = W4790
tmp14617 := PrimIsPair(W4790)

if True == tmp14617 {
tmp14597 := MakeNative(func(__e *ControlFlow) {
W4791 := __e.Get(1)
_ = W4791
tmp14598 := MakeNative(func(__e *ControlFlow) {
W4792 := __e.Get(1)
_ = W4792
tmp14612 := PrimIsPair(W4792)

if True == tmp14612 {
tmp14599 := MakeNative(func(__e *ControlFlow) {
W4793 := __e.Get(1)
_ = W4793
tmp14600 := MakeNative(func(__e *ControlFlow) {
W4794 := __e.Get(1)
_ = W4794
tmp14607 := PrimEqual(W4794, Nil)

if True == tmp14607 {
tmp14601 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14601

tmp14602 := MakeNative(func(__e *ControlFlow) {
tmp14603 := PrimCons(W4791, Nil)

tmp14604 := PrimCons(symvalue, tmp14603)

tmp14605 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4793, V4587, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp14604, V4587, V4588, V4589, V4590, W4593, tmp14605)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4791, symsymbol, V4588, V4589, V4590, W4593, tmp14602)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14608 := PrimTail(W4792)

tmp14609 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14608, V4589)


__e.TailApply(tmp14600, tmp14609)
return


}, 1)

tmp14610 := PrimHead(W4792)

__e.TailApply(tmp14599, tmp14610)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14613 := PrimTail(W4790)

tmp14614 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14613, V4589)


__e.TailApply(tmp14598, tmp14614)
return


}, 1)

tmp14615 := PrimHead(W4790)

__e.TailApply(tmp14597, tmp14615)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14618 := PrimTail(W4788)

tmp14619 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14618, V4589)


__e.TailApply(tmp14596, tmp14619)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14622 := PrimHead(W4788)

tmp14623 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14622, V4589)


__e.TailApply(tmp14595, tmp14623)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14626 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp14627 := Call(__e, tmp14594, tmp14626)


ifres14593 = tmp14627


} else {
ifres14593 = False


}

__e.TailApply(tmp14566, ifres14593)
return


} else {
__e.Return(W4779)
return
}


}, 1)

tmp14667 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14631 Obj

if True == tmp14667 {
tmp14632 := MakeNative(func(__e *ControlFlow) {
W4780 := __e.Get(1)
_ = W4780
tmp14664 := PrimIsPair(W4780)

if True == tmp14664 {
tmp14633 := MakeNative(func(__e *ControlFlow) {
W4781 := __e.Get(1)
_ = W4781
tmp14660 := PrimEqual(W4781, syminput_7)

if True == tmp14660 {
tmp14634 := MakeNative(func(__e *ControlFlow) {
W4782 := __e.Get(1)
_ = W4782
tmp14656 := PrimIsPair(W4782)

if True == tmp14656 {
tmp14635 := MakeNative(func(__e *ControlFlow) {
W4783 := __e.Get(1)
_ = W4783
tmp14636 := MakeNative(func(__e *ControlFlow) {
W4784 := __e.Get(1)
_ = W4784
tmp14651 := PrimIsPair(W4784)

if True == tmp14651 {
tmp14637 := MakeNative(func(__e *ControlFlow) {
W4785 := __e.Get(1)
_ = W4785
tmp14638 := MakeNative(func(__e *ControlFlow) {
W4786 := __e.Get(1)
_ = W4786
tmp14646 := PrimEqual(W4786, Nil)

if True == tmp14646 {
tmp14639 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14639

tmp14640 := Call(__e, PrimFunc(symshen_4deref), W4783, V4589)


tmp14641 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14640)


tmp14642 := MakeNative(func(__e *ControlFlow) {
tmp14643 := PrimCons(symin, Nil)

tmp14644 := PrimCons(symstream, tmp14643)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4785, tmp14644, V4588, V4589, V4590, W4593, V4592)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V4587, tmp14641, V4589, V4590, W4593, tmp14642)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14647 := PrimTail(W4784)

tmp14648 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14647, V4589)


__e.TailApply(tmp14638, tmp14648)
return


}, 1)

tmp14649 := PrimHead(W4784)

__e.TailApply(tmp14637, tmp14649)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14652 := PrimTail(W4782)

tmp14653 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14652, V4589)


__e.TailApply(tmp14636, tmp14653)
return


}, 1)

tmp14654 := PrimHead(W4782)

__e.TailApply(tmp14635, tmp14654)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14657 := PrimTail(W4780)

tmp14658 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14657, V4589)


__e.TailApply(tmp14634, tmp14658)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14661 := PrimHead(W4780)

tmp14662 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14661, V4589)


__e.TailApply(tmp14633, tmp14662)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14665 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp14666 := Call(__e, tmp14632, tmp14665)


ifres14631 = tmp14666


} else {
ifres14631 = False


}

__e.TailApply(tmp14565, ifres14631)
return


} else {
__e.Return(W4771)
return
}


}, 1)

tmp14705 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14670 Obj

if True == tmp14705 {
tmp14671 := MakeNative(func(__e *ControlFlow) {
W4772 := __e.Get(1)
_ = W4772
tmp14702 := PrimIsPair(W4772)

if True == tmp14702 {
tmp14672 := MakeNative(func(__e *ControlFlow) {
W4773 := __e.Get(1)
_ = W4773
tmp14698 := PrimEqual(W4773, symtype)

if True == tmp14698 {
tmp14673 := MakeNative(func(__e *ControlFlow) {
W4774 := __e.Get(1)
_ = W4774
tmp14694 := PrimIsPair(W4774)

if True == tmp14694 {
tmp14674 := MakeNative(func(__e *ControlFlow) {
W4775 := __e.Get(1)
_ = W4775
tmp14675 := MakeNative(func(__e *ControlFlow) {
W4776 := __e.Get(1)
_ = W4776
tmp14689 := PrimIsPair(W4776)

if True == tmp14689 {
tmp14676 := MakeNative(func(__e *ControlFlow) {
W4777 := __e.Get(1)
_ = W4777
tmp14677 := MakeNative(func(__e *ControlFlow) {
W4778 := __e.Get(1)
_ = W4778
tmp14684 := PrimEqual(W4778, Nil)

if True == tmp14684 {
tmp14678 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14678

tmp14679 := MakeNative(func(__e *ControlFlow) {
tmp14680 := Call(__e, PrimFunc(symshen_4deref), W4777, V4589)


tmp14681 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14680)


tmp14682 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4775, V4587, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp14681, V4587, V4589, V4590, W4593, tmp14682)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4589, V4590, W4593, tmp14679)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14685 := PrimTail(W4776)

tmp14686 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14685, V4589)


__e.TailApply(tmp14677, tmp14686)
return


}, 1)

tmp14687 := PrimHead(W4776)

__e.TailApply(tmp14676, tmp14687)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14690 := PrimTail(W4774)

tmp14691 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14690, V4589)


__e.TailApply(tmp14675, tmp14691)
return


}, 1)

tmp14692 := PrimHead(W4774)

__e.TailApply(tmp14674, tmp14692)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14695 := PrimTail(W4772)

tmp14696 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14695, V4589)


__e.TailApply(tmp14673, tmp14696)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14699 := PrimHead(W4772)

tmp14700 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14699, V4589)


__e.TailApply(tmp14672, tmp14700)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14703 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp14704 := Call(__e, tmp14671, tmp14703)


ifres14670 = tmp14704


} else {
ifres14670 = False


}

__e.TailApply(tmp14564, ifres14670)
return


} else {
__e.Return(W4750)
return
}


}, 1)

tmp14793 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14708 Obj

if True == tmp14793 {
tmp14709 := MakeNative(func(__e *ControlFlow) {
W4751 := __e.Get(1)
_ = W4751
tmp14790 := PrimIsPair(W4751)

if True == tmp14790 {
tmp14710 := MakeNative(func(__e *ControlFlow) {
W4752 := __e.Get(1)
_ = W4752
tmp14786 := PrimEqual(W4752, symopen)

if True == tmp14786 {
tmp14711 := MakeNative(func(__e *ControlFlow) {
W4753 := __e.Get(1)
_ = W4753
tmp14782 := PrimIsPair(W4753)

if True == tmp14782 {
tmp14712 := MakeNative(func(__e *ControlFlow) {
W4754 := __e.Get(1)
_ = W4754
tmp14713 := MakeNative(func(__e *ControlFlow) {
W4755 := __e.Get(1)
_ = W4755
tmp14777 := PrimIsPair(W4755)

if True == tmp14777 {
tmp14714 := MakeNative(func(__e *ControlFlow) {
W4756 := __e.Get(1)
_ = W4756
tmp14715 := MakeNative(func(__e *ControlFlow) {
W4757 := __e.Get(1)
_ = W4757
tmp14772 := PrimEqual(W4757, Nil)

if True == tmp14772 {
tmp14716 := MakeNative(func(__e *ControlFlow) {
W4758 := __e.Get(1)
_ = W4758
tmp14717 := MakeNative(func(__e *ControlFlow) {
W4759 := __e.Get(1)
_ = W4759
tmp14761 := PrimIsPair(W4758)

if True == tmp14761 {
tmp14718 := MakeNative(func(__e *ControlFlow) {
W4761 := __e.Get(1)
_ = W4761
tmp14719 := MakeNative(func(__e *ControlFlow) {
W4762 := __e.Get(1)
_ = W4762
tmp14723 := PrimEqual(W4761, symstream)

if True == tmp14723 {
__e.TailApply(PrimFunc(symthaw), W4762)
return
} else {
tmp14721 := Call(__e, PrimFunc(symshen_4pvar_2), W4761)


if True == tmp14721 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4761, symstream, V4589, W4762)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14724 := MakeNative(func(__e *ControlFlow) {
tmp14725 := MakeNative(func(__e *ControlFlow) {
W4763 := __e.Get(1)
_ = W4763
tmp14726 := MakeNative(func(__e *ControlFlow) {
W4764 := __e.Get(1)
_ = W4764
tmp14746 := PrimIsPair(W4763)

if True == tmp14746 {
tmp14727 := MakeNative(func(__e *ControlFlow) {
W4766 := __e.Get(1)
_ = W4766
tmp14728 := MakeNative(func(__e *ControlFlow) {
W4767 := __e.Get(1)
_ = W4767
tmp14729 := MakeNative(func(__e *ControlFlow) {
W4768 := __e.Get(1)
_ = W4768
tmp14733 := PrimEqual(W4767, Nil)

if True == tmp14733 {
__e.TailApply(PrimFunc(symthaw), W4768)
return
} else {
tmp14731 := Call(__e, PrimFunc(symshen_4pvar_2), W4767)


if True == tmp14731 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4767, Nil, V4589, W4768)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14734 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4764, W4766)
return
}, 0)

__e.TailApply(tmp14729, tmp14734)
return


}, 1)

tmp14735 := PrimTail(W4763)

tmp14736 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14735, V4589)


__e.TailApply(tmp14728, tmp14736)
return


}, 1)

tmp14737 := PrimHead(W4763)

__e.TailApply(tmp14727, tmp14737)
return


} else {
tmp14744 := Call(__e, PrimFunc(symshen_4pvar_2), W4763)


if True == tmp14744 {
tmp14738 := MakeNative(func(__e *ControlFlow) {
W4769 := __e.Get(1)
_ = W4769
tmp14739 := PrimCons(W4769, Nil)

tmp14740 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4764, W4769)
return
}, 0)

tmp14741 := Call(__e, PrimFunc(symshen_4bind_b), W4763, tmp14739, V4589, tmp14740)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14741)
return


}, 1)

tmp14742 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14738, tmp14742)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14747 := MakeNative(func(__e *ControlFlow) {
Z4765 := __e.Get(1)
_ = Z4765
__e.TailApply(W4759, Z4765)
return
}, 1)

__e.TailApply(tmp14726, tmp14747)
return


}, 1)

tmp14748 := PrimTail(W4758)

tmp14749 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14748, V4589)


__e.TailApply(tmp14725, tmp14749)
return


}, 0)

__e.TailApply(tmp14719, tmp14724)
return


}, 1)

tmp14750 := PrimHead(W4758)

tmp14751 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14750, V4589)


__e.TailApply(tmp14718, tmp14751)
return


} else {
tmp14759 := Call(__e, PrimFunc(symshen_4pvar_2), W4758)


if True == tmp14759 {
tmp14752 := MakeNative(func(__e *ControlFlow) {
W4770 := __e.Get(1)
_ = W4770
tmp14753 := PrimCons(W4770, Nil)

tmp14754 := PrimCons(symstream, tmp14753)

tmp14755 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4759, W4770)
return
}, 0)

tmp14756 := Call(__e, PrimFunc(symshen_4bind_b), W4758, tmp14754, V4589, tmp14755)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14756)
return


}, 1)

tmp14757 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14752, tmp14757)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14762 := MakeNative(func(__e *ControlFlow) {
Z4760 := __e.Get(1)
_ = Z4760
tmp14763 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14763

tmp14764 := MakeNative(func(__e *ControlFlow) {
tmp14765 := Call(__e, PrimFunc(symshen_4lazyderef), Z4760, V4589)


tmp14766 := PrimCons(symout, Nil)

tmp14767 := PrimCons(symin, tmp14766)

tmp14768 := Call(__e, PrimFunc(symelement_2), tmp14765, tmp14767)


tmp14769 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4754, symstring, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14768, V4589, V4590, W4593, tmp14769)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W4756, Z4760, V4589, V4590, W4593, tmp14764)
return


}, 1)

__e.TailApply(tmp14717, tmp14762)
return


}, 1)

tmp14770 := Call(__e, PrimFunc(symshen_4lazyderef), V4587, V4589)


__e.TailApply(tmp14716, tmp14770)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14773 := PrimTail(W4755)

tmp14774 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14773, V4589)


__e.TailApply(tmp14715, tmp14774)
return


}, 1)

tmp14775 := PrimHead(W4755)

__e.TailApply(tmp14714, tmp14775)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14778 := PrimTail(W4753)

tmp14779 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14778, V4589)


__e.TailApply(tmp14713, tmp14779)
return


}, 1)

tmp14780 := PrimHead(W4753)

__e.TailApply(tmp14712, tmp14780)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14783 := PrimTail(W4751)

tmp14784 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14783, V4589)


__e.TailApply(tmp14711, tmp14784)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14787 := PrimHead(W4751)

tmp14788 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14787, V4589)


__e.TailApply(tmp14710, tmp14788)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14791 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp14792 := Call(__e, tmp14709, tmp14791)


ifres14708 = tmp14792


} else {
ifres14708 = False


}

__e.TailApply(tmp14563, ifres14708)
return


} else {
__e.Return(W4737)
return
}


}, 1)

tmp14857 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14796 Obj

if True == tmp14857 {
tmp14797 := MakeNative(func(__e *ControlFlow) {
W4738 := __e.Get(1)
_ = W4738
tmp14854 := PrimIsPair(W4738)

if True == tmp14854 {
tmp14798 := MakeNative(func(__e *ControlFlow) {
W4739 := __e.Get(1)
_ = W4739
tmp14850 := PrimEqual(W4739, symlet)

if True == tmp14850 {
tmp14799 := MakeNative(func(__e *ControlFlow) {
W4740 := __e.Get(1)
_ = W4740
tmp14846 := PrimIsPair(W4740)

if True == tmp14846 {
tmp14800 := MakeNative(func(__e *ControlFlow) {
W4741 := __e.Get(1)
_ = W4741
tmp14801 := MakeNative(func(__e *ControlFlow) {
W4742 := __e.Get(1)
_ = W4742
tmp14841 := PrimIsPair(W4742)

if True == tmp14841 {
tmp14802 := MakeNative(func(__e *ControlFlow) {
W4743 := __e.Get(1)
_ = W4743
tmp14803 := MakeNative(func(__e *ControlFlow) {
W4744 := __e.Get(1)
_ = W4744
tmp14836 := PrimIsPair(W4744)

if True == tmp14836 {
tmp14804 := MakeNative(func(__e *ControlFlow) {
W4745 := __e.Get(1)
_ = W4745
tmp14805 := MakeNative(func(__e *ControlFlow) {
W4746 := __e.Get(1)
_ = W4746
tmp14831 := PrimEqual(W4746, Nil)

if True == tmp14831 {
tmp14806 := MakeNative(func(__e *ControlFlow) {
W4747 := __e.Get(1)
_ = W4747
tmp14807 := MakeNative(func(__e *ControlFlow) {
W4748 := __e.Get(1)
_ = W4748
tmp14808 := MakeNative(func(__e *ControlFlow) {
W4749 := __e.Get(1)
_ = W4749
tmp14809 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14809

tmp14810 := MakeNative(func(__e *ControlFlow) {
tmp14811 := Call(__e, PrimFunc(symshen_4lazyderef), W4741, V4589)


tmp14812 := Call(__e, PrimFunc(symshen_4freshterm), tmp14811)


tmp14813 := MakeNative(func(__e *ControlFlow) {
tmp14814 := Call(__e, PrimFunc(symshen_4lazyderef), W4741, V4589)


tmp14815 := Call(__e, PrimFunc(symshen_4lazyderef), W4748, V4589)


tmp14816 := Call(__e, PrimFunc(symshen_4lazyderef), W4745, V4589)


tmp14817 := Call(__e, PrimFunc(symshen_4beta), tmp14814, tmp14815, tmp14816)


tmp14818 := MakeNative(func(__e *ControlFlow) {
tmp14819 := PrimIntern(MakeString(":"))

tmp14820 := PrimCons(W4749, Nil)

tmp14821 := PrimCons(tmp14819, tmp14820)

tmp14822 := PrimCons(W4748, tmp14821)

tmp14823 := PrimCons(tmp14822, V4588)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4747, V4587, tmp14823, V4589, V4590, W4593, V4592)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4747, tmp14817, V4589, V4590, W4593, tmp14818)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4748, tmp14812, V4589, V4590, W4593, tmp14813)
return


}, 0)

tmp14824 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4743, W4749, V4588, V4589, V4590, W4593, tmp14810)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14824)
return


}, 1)

tmp14825 := Call(__e, PrimFunc(symshen_4newpv), V4589)


tmp14826 := Call(__e, tmp14808, tmp14825)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14826)
return


}, 1)

tmp14827 := Call(__e, PrimFunc(symshen_4newpv), V4589)


tmp14828 := Call(__e, tmp14807, tmp14827)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14828)
return


}, 1)

tmp14829 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14806, tmp14829)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14832 := PrimTail(W4744)

tmp14833 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14832, V4589)


__e.TailApply(tmp14805, tmp14833)
return


}, 1)

tmp14834 := PrimHead(W4744)

__e.TailApply(tmp14804, tmp14834)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14837 := PrimTail(W4742)

tmp14838 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14837, V4589)


__e.TailApply(tmp14803, tmp14838)
return


}, 1)

tmp14839 := PrimHead(W4742)

__e.TailApply(tmp14802, tmp14839)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14842 := PrimTail(W4740)

tmp14843 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14842, V4589)


__e.TailApply(tmp14801, tmp14843)
return


}, 1)

tmp14844 := PrimHead(W4740)

__e.TailApply(tmp14800, tmp14844)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14847 := PrimTail(W4738)

tmp14848 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14847, V4589)


__e.TailApply(tmp14799, tmp14848)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14851 := PrimHead(W4738)

tmp14852 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14851, V4589)


__e.TailApply(tmp14798, tmp14852)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14855 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp14856 := Call(__e, tmp14797, tmp14855)


ifres14796 = tmp14856


} else {
ifres14796 = False


}

__e.TailApply(tmp14562, ifres14796)
return


} else {
__e.Return(W4707)
return
}


}, 1)

tmp14981 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14860 Obj

if True == tmp14981 {
tmp14861 := MakeNative(func(__e *ControlFlow) {
W4708 := __e.Get(1)
_ = W4708
tmp14978 := PrimIsPair(W4708)

if True == tmp14978 {
tmp14862 := MakeNative(func(__e *ControlFlow) {
W4709 := __e.Get(1)
_ = W4709
tmp14974 := PrimEqual(W4709, symlambda)

if True == tmp14974 {
tmp14863 := MakeNative(func(__e *ControlFlow) {
W4710 := __e.Get(1)
_ = W4710
tmp14970 := PrimIsPair(W4710)

if True == tmp14970 {
tmp14864 := MakeNative(func(__e *ControlFlow) {
W4711 := __e.Get(1)
_ = W4711
tmp14865 := MakeNative(func(__e *ControlFlow) {
W4712 := __e.Get(1)
_ = W4712
tmp14965 := PrimIsPair(W4712)

if True == tmp14965 {
tmp14866 := MakeNative(func(__e *ControlFlow) {
W4713 := __e.Get(1)
_ = W4713
tmp14867 := MakeNative(func(__e *ControlFlow) {
W4714 := __e.Get(1)
_ = W4714
tmp14960 := PrimEqual(W4714, Nil)

if True == tmp14960 {
tmp14868 := MakeNative(func(__e *ControlFlow) {
W4715 := __e.Get(1)
_ = W4715
tmp14869 := MakeNative(func(__e *ControlFlow) {
W4716 := __e.Get(1)
_ = W4716
tmp14936 := PrimIsPair(W4715)

if True == tmp14936 {
tmp14870 := MakeNative(func(__e *ControlFlow) {
W4721 := __e.Get(1)
_ = W4721
tmp14871 := MakeNative(func(__e *ControlFlow) {
W4722 := __e.Get(1)
_ = W4722
tmp14872 := MakeNative(func(__e *ControlFlow) {
W4723 := __e.Get(1)
_ = W4723
tmp14916 := PrimIsPair(W4722)

if True == tmp14916 {
tmp14873 := MakeNative(func(__e *ControlFlow) {
W4725 := __e.Get(1)
_ = W4725
tmp14874 := MakeNative(func(__e *ControlFlow) {
W4726 := __e.Get(1)
_ = W4726
tmp14878 := PrimEqual(W4725, sym_1_1_6)

if True == tmp14878 {
__e.TailApply(PrimFunc(symthaw), W4726)
return
} else {
tmp14876 := Call(__e, PrimFunc(symshen_4pvar_2), W4725)


if True == tmp14876 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4725, sym_1_1_6, V4589, W4726)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14879 := MakeNative(func(__e *ControlFlow) {
tmp14880 := MakeNative(func(__e *ControlFlow) {
W4727 := __e.Get(1)
_ = W4727
tmp14881 := MakeNative(func(__e *ControlFlow) {
W4728 := __e.Get(1)
_ = W4728
tmp14901 := PrimIsPair(W4727)

if True == tmp14901 {
tmp14882 := MakeNative(func(__e *ControlFlow) {
W4730 := __e.Get(1)
_ = W4730
tmp14883 := MakeNative(func(__e *ControlFlow) {
W4731 := __e.Get(1)
_ = W4731
tmp14884 := MakeNative(func(__e *ControlFlow) {
W4732 := __e.Get(1)
_ = W4732
tmp14888 := PrimEqual(W4731, Nil)

if True == tmp14888 {
__e.TailApply(PrimFunc(symthaw), W4732)
return
} else {
tmp14886 := Call(__e, PrimFunc(symshen_4pvar_2), W4731)


if True == tmp14886 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4731, Nil, V4589, W4732)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14889 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4728, W4730)
return
}, 0)

__e.TailApply(tmp14884, tmp14889)
return


}, 1)

tmp14890 := PrimTail(W4727)

tmp14891 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14890, V4589)


__e.TailApply(tmp14883, tmp14891)
return


}, 1)

tmp14892 := PrimHead(W4727)

__e.TailApply(tmp14882, tmp14892)
return


} else {
tmp14899 := Call(__e, PrimFunc(symshen_4pvar_2), W4727)


if True == tmp14899 {
tmp14893 := MakeNative(func(__e *ControlFlow) {
W4733 := __e.Get(1)
_ = W4733
tmp14894 := PrimCons(W4733, Nil)

tmp14895 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4728, W4733)
return
}, 0)

tmp14896 := Call(__e, PrimFunc(symshen_4bind_b), W4727, tmp14894, V4589, tmp14895)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14896)
return


}, 1)

tmp14897 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14893, tmp14897)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14902 := MakeNative(func(__e *ControlFlow) {
Z4729 := __e.Get(1)
_ = Z4729
__e.TailApply(W4723, Z4729)
return
}, 1)

__e.TailApply(tmp14881, tmp14902)
return


}, 1)

tmp14903 := PrimTail(W4722)

tmp14904 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14903, V4589)


__e.TailApply(tmp14880, tmp14904)
return


}, 0)

__e.TailApply(tmp14874, tmp14879)
return


}, 1)

tmp14905 := PrimHead(W4722)

tmp14906 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14905, V4589)


__e.TailApply(tmp14873, tmp14906)
return


} else {
tmp14914 := Call(__e, PrimFunc(symshen_4pvar_2), W4722)


if True == tmp14914 {
tmp14907 := MakeNative(func(__e *ControlFlow) {
W4734 := __e.Get(1)
_ = W4734
tmp14908 := PrimCons(W4734, Nil)

tmp14909 := PrimCons(sym_1_1_6, tmp14908)

tmp14910 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4723, W4734)
return
}, 0)

tmp14911 := Call(__e, PrimFunc(symshen_4bind_b), W4722, tmp14909, V4589, tmp14910)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14911)
return


}, 1)

tmp14912 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14907, tmp14912)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14917 := MakeNative(func(__e *ControlFlow) {
Z4724 := __e.Get(1)
_ = Z4724
tmp14918 := Call(__e, W4716, W4721)


__e.TailApply(tmp14918, Z4724)
return


}, 1)

__e.TailApply(tmp14872, tmp14917)
return


}, 1)

tmp14919 := PrimTail(W4715)

tmp14920 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14919, V4589)


__e.TailApply(tmp14871, tmp14920)
return


}, 1)

tmp14921 := PrimHead(W4715)

__e.TailApply(tmp14870, tmp14921)
return


} else {
tmp14934 := Call(__e, PrimFunc(symshen_4pvar_2), W4715)


if True == tmp14934 {
tmp14922 := MakeNative(func(__e *ControlFlow) {
W4735 := __e.Get(1)
_ = W4735
tmp14923 := MakeNative(func(__e *ControlFlow) {
W4736 := __e.Get(1)
_ = W4736
tmp14924 := PrimCons(W4736, Nil)

tmp14925 := PrimCons(sym_1_1_6, tmp14924)

tmp14926 := PrimCons(W4735, tmp14925)

tmp14927 := MakeNative(func(__e *ControlFlow) {
tmp14928 := Call(__e, W4716, W4735)


__e.TailApply(tmp14928, W4736)
return


}, 0)

tmp14929 := Call(__e, PrimFunc(symshen_4bind_b), W4715, tmp14926, V4589, tmp14927)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14929)
return


}, 1)

tmp14930 := Call(__e, PrimFunc(symshen_4newpv), V4589)


tmp14931 := Call(__e, tmp14923, tmp14930)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14931)
return


}, 1)

tmp14932 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14922, tmp14932)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14937 := MakeNative(func(__e *ControlFlow) {
Z4717 := __e.Get(1)
_ = Z4717
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4718 := __e.Get(1)
_ = Z4718
tmp14938 := MakeNative(func(__e *ControlFlow) {
W4719 := __e.Get(1)
_ = W4719
tmp14939 := MakeNative(func(__e *ControlFlow) {
W4720 := __e.Get(1)
_ = W4720
tmp14940 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14940

tmp14941 := Call(__e, PrimFunc(symshen_4lazyderef), W4711, V4589)


tmp14942 := Call(__e, PrimFunc(symshen_4freshterm), tmp14941)


tmp14943 := MakeNative(func(__e *ControlFlow) {
tmp14944 := Call(__e, PrimFunc(symshen_4lazyderef), W4711, V4589)


tmp14945 := Call(__e, PrimFunc(symshen_4deref), W4720, V4589)


tmp14946 := Call(__e, PrimFunc(symshen_4deref), W4713, V4589)


tmp14947 := Call(__e, PrimFunc(symshen_4beta), tmp14944, tmp14945, tmp14946)


tmp14948 := MakeNative(func(__e *ControlFlow) {
tmp14949 := PrimIntern(MakeString(":"))

tmp14950 := PrimCons(Z4717, Nil)

tmp14951 := PrimCons(tmp14949, tmp14950)

tmp14952 := PrimCons(W4720, tmp14951)

tmp14953 := PrimCons(tmp14952, V4588)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4719, Z4718, tmp14953, V4589, V4590, W4593, V4592)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4719, tmp14947, V4589, V4590, W4593, tmp14948)
return


}, 0)

tmp14954 := Call(__e, PrimFunc(symbind), W4720, tmp14942, V4589, V4590, W4593, tmp14943)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14954)
return


}, 1)

tmp14955 := Call(__e, PrimFunc(symshen_4newpv), V4589)


tmp14956 := Call(__e, tmp14939, tmp14955)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp14956)
return


}, 1)

tmp14957 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp14938, tmp14957)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14869, tmp14937)
return


}, 1)

tmp14958 := Call(__e, PrimFunc(symshen_4lazyderef), V4587, V4589)


__e.TailApply(tmp14868, tmp14958)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14961 := PrimTail(W4712)

tmp14962 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14961, V4589)


__e.TailApply(tmp14867, tmp14962)
return


}, 1)

tmp14963 := PrimHead(W4712)

__e.TailApply(tmp14866, tmp14963)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14966 := PrimTail(W4710)

tmp14967 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14966, V4589)


__e.TailApply(tmp14865, tmp14967)
return


}, 1)

tmp14968 := PrimHead(W4710)

__e.TailApply(tmp14864, tmp14968)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14971 := PrimTail(W4708)

tmp14972 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14971, V4589)


__e.TailApply(tmp14863, tmp14972)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14975 := PrimHead(W4708)

tmp14976 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14975, V4589)


__e.TailApply(tmp14862, tmp14976)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14979 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp14980 := Call(__e, tmp14861, tmp14979)


ifres14860 = tmp14980


} else {
ifres14860 = False


}

__e.TailApply(tmp14561, ifres14860)
return


} else {
__e.Return(W4697)
return
}


}, 1)

tmp15024 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres14984 Obj

if True == tmp15024 {
tmp14985 := MakeNative(func(__e *ControlFlow) {
W4698 := __e.Get(1)
_ = W4698
tmp15021 := PrimIsPair(W4698)

if True == tmp15021 {
tmp14986 := MakeNative(func(__e *ControlFlow) {
W4699 := __e.Get(1)
_ = W4699
tmp15017 := PrimEqual(W4699, sym_8s)

if True == tmp15017 {
tmp14987 := MakeNative(func(__e *ControlFlow) {
W4700 := __e.Get(1)
_ = W4700
tmp15013 := PrimIsPair(W4700)

if True == tmp15013 {
tmp14988 := MakeNative(func(__e *ControlFlow) {
W4701 := __e.Get(1)
_ = W4701
tmp14989 := MakeNative(func(__e *ControlFlow) {
W4702 := __e.Get(1)
_ = W4702
tmp15008 := PrimIsPair(W4702)

if True == tmp15008 {
tmp14990 := MakeNative(func(__e *ControlFlow) {
W4703 := __e.Get(1)
_ = W4703
tmp14991 := MakeNative(func(__e *ControlFlow) {
W4704 := __e.Get(1)
_ = W4704
tmp15003 := PrimEqual(W4704, Nil)

if True == tmp15003 {
tmp14992 := MakeNative(func(__e *ControlFlow) {
W4705 := __e.Get(1)
_ = W4705
tmp14993 := MakeNative(func(__e *ControlFlow) {
W4706 := __e.Get(1)
_ = W4706
tmp14997 := PrimEqual(W4705, symstring)

if True == tmp14997 {
__e.TailApply(PrimFunc(symthaw), W4706)
return
} else {
tmp14995 := Call(__e, PrimFunc(symshen_4pvar_2), W4705)


if True == tmp14995 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4705, symstring, V4589, W4706)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14998 := MakeNative(func(__e *ControlFlow) {
tmp14999 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14999

tmp15000 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4703, symstring, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4701, symstring, V4588, V4589, V4590, W4593, tmp15000)
return


}, 0)

__e.TailApply(tmp14993, tmp14998)
return


}, 1)

tmp15001 := Call(__e, PrimFunc(symshen_4lazyderef), V4587, V4589)


__e.TailApply(tmp14992, tmp15001)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15004 := PrimTail(W4702)

tmp15005 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15004, V4589)


__e.TailApply(tmp14991, tmp15005)
return


}, 1)

tmp15006 := PrimHead(W4702)

__e.TailApply(tmp14990, tmp15006)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15009 := PrimTail(W4700)

tmp15010 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15009, V4589)


__e.TailApply(tmp14989, tmp15010)
return


}, 1)

tmp15011 := PrimHead(W4700)

__e.TailApply(tmp14988, tmp15011)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15014 := PrimTail(W4698)

tmp15015 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15014, V4589)


__e.TailApply(tmp14987, tmp15015)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15018 := PrimHead(W4698)

tmp15019 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15018, V4589)


__e.TailApply(tmp14986, tmp15019)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15022 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15023 := Call(__e, tmp14985, tmp15022)


ifres14984 = tmp15023


} else {
ifres14984 = False


}

__e.TailApply(tmp14560, ifres14984)
return


} else {
__e.Return(W4676)
return
}


}, 1)

tmp15109 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15027 Obj

if True == tmp15109 {
tmp15028 := MakeNative(func(__e *ControlFlow) {
W4677 := __e.Get(1)
_ = W4677
tmp15106 := PrimIsPair(W4677)

if True == tmp15106 {
tmp15029 := MakeNative(func(__e *ControlFlow) {
W4678 := __e.Get(1)
_ = W4678
tmp15102 := PrimEqual(W4678, sym_8v)

if True == tmp15102 {
tmp15030 := MakeNative(func(__e *ControlFlow) {
W4679 := __e.Get(1)
_ = W4679
tmp15098 := PrimIsPair(W4679)

if True == tmp15098 {
tmp15031 := MakeNative(func(__e *ControlFlow) {
W4680 := __e.Get(1)
_ = W4680
tmp15032 := MakeNative(func(__e *ControlFlow) {
W4681 := __e.Get(1)
_ = W4681
tmp15093 := PrimIsPair(W4681)

if True == tmp15093 {
tmp15033 := MakeNative(func(__e *ControlFlow) {
W4682 := __e.Get(1)
_ = W4682
tmp15034 := MakeNative(func(__e *ControlFlow) {
W4683 := __e.Get(1)
_ = W4683
tmp15088 := PrimEqual(W4683, Nil)

if True == tmp15088 {
tmp15035 := MakeNative(func(__e *ControlFlow) {
W4684 := __e.Get(1)
_ = W4684
tmp15036 := MakeNative(func(__e *ControlFlow) {
W4685 := __e.Get(1)
_ = W4685
tmp15080 := PrimIsPair(W4684)

if True == tmp15080 {
tmp15037 := MakeNative(func(__e *ControlFlow) {
W4687 := __e.Get(1)
_ = W4687
tmp15038 := MakeNative(func(__e *ControlFlow) {
W4688 := __e.Get(1)
_ = W4688
tmp15042 := PrimEqual(W4687, symvector)

if True == tmp15042 {
__e.TailApply(PrimFunc(symthaw), W4688)
return
} else {
tmp15040 := Call(__e, PrimFunc(symshen_4pvar_2), W4687)


if True == tmp15040 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4687, symvector, V4589, W4688)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15043 := MakeNative(func(__e *ControlFlow) {
tmp15044 := MakeNative(func(__e *ControlFlow) {
W4689 := __e.Get(1)
_ = W4689
tmp15045 := MakeNative(func(__e *ControlFlow) {
W4690 := __e.Get(1)
_ = W4690
tmp15065 := PrimIsPair(W4689)

if True == tmp15065 {
tmp15046 := MakeNative(func(__e *ControlFlow) {
W4692 := __e.Get(1)
_ = W4692
tmp15047 := MakeNative(func(__e *ControlFlow) {
W4693 := __e.Get(1)
_ = W4693
tmp15048 := MakeNative(func(__e *ControlFlow) {
W4694 := __e.Get(1)
_ = W4694
tmp15052 := PrimEqual(W4693, Nil)

if True == tmp15052 {
__e.TailApply(PrimFunc(symthaw), W4694)
return
} else {
tmp15050 := Call(__e, PrimFunc(symshen_4pvar_2), W4693)


if True == tmp15050 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4693, Nil, V4589, W4694)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15053 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4690, W4692)
return
}, 0)

__e.TailApply(tmp15048, tmp15053)
return


}, 1)

tmp15054 := PrimTail(W4689)

tmp15055 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15054, V4589)


__e.TailApply(tmp15047, tmp15055)
return


}, 1)

tmp15056 := PrimHead(W4689)

__e.TailApply(tmp15046, tmp15056)
return


} else {
tmp15063 := Call(__e, PrimFunc(symshen_4pvar_2), W4689)


if True == tmp15063 {
tmp15057 := MakeNative(func(__e *ControlFlow) {
W4695 := __e.Get(1)
_ = W4695
tmp15058 := PrimCons(W4695, Nil)

tmp15059 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4690, W4695)
return
}, 0)

tmp15060 := Call(__e, PrimFunc(symshen_4bind_b), W4689, tmp15058, V4589, tmp15059)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15060)
return


}, 1)

tmp15061 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15057, tmp15061)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15066 := MakeNative(func(__e *ControlFlow) {
Z4691 := __e.Get(1)
_ = Z4691
__e.TailApply(W4685, Z4691)
return
}, 1)

__e.TailApply(tmp15045, tmp15066)
return


}, 1)

tmp15067 := PrimTail(W4684)

tmp15068 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15067, V4589)


__e.TailApply(tmp15044, tmp15068)
return


}, 0)

__e.TailApply(tmp15038, tmp15043)
return


}, 1)

tmp15069 := PrimHead(W4684)

tmp15070 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15069, V4589)


__e.TailApply(tmp15037, tmp15070)
return


} else {
tmp15078 := Call(__e, PrimFunc(symshen_4pvar_2), W4684)


if True == tmp15078 {
tmp15071 := MakeNative(func(__e *ControlFlow) {
W4696 := __e.Get(1)
_ = W4696
tmp15072 := PrimCons(W4696, Nil)

tmp15073 := PrimCons(symvector, tmp15072)

tmp15074 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4685, W4696)
return
}, 0)

tmp15075 := Call(__e, PrimFunc(symshen_4bind_b), W4684, tmp15073, V4589, tmp15074)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15075)
return


}, 1)

tmp15076 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15071, tmp15076)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15081 := MakeNative(func(__e *ControlFlow) {
Z4686 := __e.Get(1)
_ = Z4686
tmp15082 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15082

tmp15083 := MakeNative(func(__e *ControlFlow) {
tmp15084 := PrimCons(Z4686, Nil)

tmp15085 := PrimCons(symvector, tmp15084)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4682, tmp15085, V4588, V4589, V4590, W4593, V4592)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4680, Z4686, V4588, V4589, V4590, W4593, tmp15083)
return


}, 1)

__e.TailApply(tmp15036, tmp15081)
return


}, 1)

tmp15086 := Call(__e, PrimFunc(symshen_4lazyderef), V4587, V4589)


__e.TailApply(tmp15035, tmp15086)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15089 := PrimTail(W4681)

tmp15090 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15089, V4589)


__e.TailApply(tmp15034, tmp15090)
return


}, 1)

tmp15091 := PrimHead(W4681)

__e.TailApply(tmp15033, tmp15091)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15094 := PrimTail(W4679)

tmp15095 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15094, V4589)


__e.TailApply(tmp15032, tmp15095)
return


}, 1)

tmp15096 := PrimHead(W4679)

__e.TailApply(tmp15031, tmp15096)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15099 := PrimTail(W4677)

tmp15100 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15099, V4589)


__e.TailApply(tmp15030, tmp15100)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15103 := PrimHead(W4677)

tmp15104 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15103, V4589)


__e.TailApply(tmp15029, tmp15104)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15107 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15108 := Call(__e, tmp15028, tmp15107)


ifres15027 = tmp15108


} else {
ifres15027 = False


}

__e.TailApply(tmp14559, ifres15027)
return


} else {
__e.Return(W4648)
return
}


}, 1)

tmp15215 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15112 Obj

if True == tmp15215 {
tmp15113 := MakeNative(func(__e *ControlFlow) {
W4649 := __e.Get(1)
_ = W4649
tmp15212 := PrimIsPair(W4649)

if True == tmp15212 {
tmp15114 := MakeNative(func(__e *ControlFlow) {
W4650 := __e.Get(1)
_ = W4650
tmp15208 := PrimEqual(W4650, sym_8p)

if True == tmp15208 {
tmp15115 := MakeNative(func(__e *ControlFlow) {
W4651 := __e.Get(1)
_ = W4651
tmp15204 := PrimIsPair(W4651)

if True == tmp15204 {
tmp15116 := MakeNative(func(__e *ControlFlow) {
W4652 := __e.Get(1)
_ = W4652
tmp15117 := MakeNative(func(__e *ControlFlow) {
W4653 := __e.Get(1)
_ = W4653
tmp15199 := PrimIsPair(W4653)

if True == tmp15199 {
tmp15118 := MakeNative(func(__e *ControlFlow) {
W4654 := __e.Get(1)
_ = W4654
tmp15119 := MakeNative(func(__e *ControlFlow) {
W4655 := __e.Get(1)
_ = W4655
tmp15194 := PrimEqual(W4655, Nil)

if True == tmp15194 {
tmp15120 := MakeNative(func(__e *ControlFlow) {
W4656 := __e.Get(1)
_ = W4656
tmp15121 := MakeNative(func(__e *ControlFlow) {
W4657 := __e.Get(1)
_ = W4657
tmp15188 := PrimIsPair(W4656)

if True == tmp15188 {
tmp15122 := MakeNative(func(__e *ControlFlow) {
W4660 := __e.Get(1)
_ = W4660
tmp15123 := MakeNative(func(__e *ControlFlow) {
W4661 := __e.Get(1)
_ = W4661
tmp15124 := MakeNative(func(__e *ControlFlow) {
W4662 := __e.Get(1)
_ = W4662
tmp15168 := PrimIsPair(W4661)

if True == tmp15168 {
tmp15125 := MakeNative(func(__e *ControlFlow) {
W4664 := __e.Get(1)
_ = W4664
tmp15126 := MakeNative(func(__e *ControlFlow) {
W4665 := __e.Get(1)
_ = W4665
tmp15130 := PrimEqual(W4664, sym_d)

if True == tmp15130 {
__e.TailApply(PrimFunc(symthaw), W4665)
return
} else {
tmp15128 := Call(__e, PrimFunc(symshen_4pvar_2), W4664)


if True == tmp15128 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4664, sym_d, V4589, W4665)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15131 := MakeNative(func(__e *ControlFlow) {
tmp15132 := MakeNative(func(__e *ControlFlow) {
W4666 := __e.Get(1)
_ = W4666
tmp15133 := MakeNative(func(__e *ControlFlow) {
W4667 := __e.Get(1)
_ = W4667
tmp15153 := PrimIsPair(W4666)

if True == tmp15153 {
tmp15134 := MakeNative(func(__e *ControlFlow) {
W4669 := __e.Get(1)
_ = W4669
tmp15135 := MakeNative(func(__e *ControlFlow) {
W4670 := __e.Get(1)
_ = W4670
tmp15136 := MakeNative(func(__e *ControlFlow) {
W4671 := __e.Get(1)
_ = W4671
tmp15140 := PrimEqual(W4670, Nil)

if True == tmp15140 {
__e.TailApply(PrimFunc(symthaw), W4671)
return
} else {
tmp15138 := Call(__e, PrimFunc(symshen_4pvar_2), W4670)


if True == tmp15138 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4670, Nil, V4589, W4671)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15141 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4667, W4669)
return
}, 0)

__e.TailApply(tmp15136, tmp15141)
return


}, 1)

tmp15142 := PrimTail(W4666)

tmp15143 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15142, V4589)


__e.TailApply(tmp15135, tmp15143)
return


}, 1)

tmp15144 := PrimHead(W4666)

__e.TailApply(tmp15134, tmp15144)
return


} else {
tmp15151 := Call(__e, PrimFunc(symshen_4pvar_2), W4666)


if True == tmp15151 {
tmp15145 := MakeNative(func(__e *ControlFlow) {
W4672 := __e.Get(1)
_ = W4672
tmp15146 := PrimCons(W4672, Nil)

tmp15147 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4667, W4672)
return
}, 0)

tmp15148 := Call(__e, PrimFunc(symshen_4bind_b), W4666, tmp15146, V4589, tmp15147)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15148)
return


}, 1)

tmp15149 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15145, tmp15149)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15154 := MakeNative(func(__e *ControlFlow) {
Z4668 := __e.Get(1)
_ = Z4668
__e.TailApply(W4662, Z4668)
return
}, 1)

__e.TailApply(tmp15133, tmp15154)
return


}, 1)

tmp15155 := PrimTail(W4661)

tmp15156 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15155, V4589)


__e.TailApply(tmp15132, tmp15156)
return


}, 0)

__e.TailApply(tmp15126, tmp15131)
return


}, 1)

tmp15157 := PrimHead(W4661)

tmp15158 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15157, V4589)


__e.TailApply(tmp15125, tmp15158)
return


} else {
tmp15166 := Call(__e, PrimFunc(symshen_4pvar_2), W4661)


if True == tmp15166 {
tmp15159 := MakeNative(func(__e *ControlFlow) {
W4673 := __e.Get(1)
_ = W4673
tmp15160 := PrimCons(W4673, Nil)

tmp15161 := PrimCons(sym_d, tmp15160)

tmp15162 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4662, W4673)
return
}, 0)

tmp15163 := Call(__e, PrimFunc(symshen_4bind_b), W4661, tmp15161, V4589, tmp15162)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15163)
return


}, 1)

tmp15164 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15159, tmp15164)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15169 := MakeNative(func(__e *ControlFlow) {
Z4663 := __e.Get(1)
_ = Z4663
tmp15170 := Call(__e, W4657, W4660)


__e.TailApply(tmp15170, Z4663)
return


}, 1)

__e.TailApply(tmp15124, tmp15169)
return


}, 1)

tmp15171 := PrimTail(W4656)

tmp15172 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15171, V4589)


__e.TailApply(tmp15123, tmp15172)
return


}, 1)

tmp15173 := PrimHead(W4656)

__e.TailApply(tmp15122, tmp15173)
return


} else {
tmp15186 := Call(__e, PrimFunc(symshen_4pvar_2), W4656)


if True == tmp15186 {
tmp15174 := MakeNative(func(__e *ControlFlow) {
W4674 := __e.Get(1)
_ = W4674
tmp15175 := MakeNative(func(__e *ControlFlow) {
W4675 := __e.Get(1)
_ = W4675
tmp15176 := PrimCons(W4675, Nil)

tmp15177 := PrimCons(sym_d, tmp15176)

tmp15178 := PrimCons(W4674, tmp15177)

tmp15179 := MakeNative(func(__e *ControlFlow) {
tmp15180 := Call(__e, W4657, W4674)


__e.TailApply(tmp15180, W4675)
return


}, 0)

tmp15181 := Call(__e, PrimFunc(symshen_4bind_b), W4656, tmp15178, V4589, tmp15179)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15181)
return


}, 1)

tmp15182 := Call(__e, PrimFunc(symshen_4newpv), V4589)


tmp15183 := Call(__e, tmp15175, tmp15182)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15183)
return


}, 1)

tmp15184 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15174, tmp15184)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15189 := MakeNative(func(__e *ControlFlow) {
Z4658 := __e.Get(1)
_ = Z4658
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4659 := __e.Get(1)
_ = Z4659
tmp15190 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15190

tmp15191 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4654, Z4659, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4652, Z4658, V4588, V4589, V4590, W4593, tmp15191)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15121, tmp15189)
return


}, 1)

tmp15192 := Call(__e, PrimFunc(symshen_4lazyderef), V4587, V4589)


__e.TailApply(tmp15120, tmp15192)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15195 := PrimTail(W4653)

tmp15196 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15195, V4589)


__e.TailApply(tmp15119, tmp15196)
return


}, 1)

tmp15197 := PrimHead(W4653)

__e.TailApply(tmp15118, tmp15197)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15200 := PrimTail(W4651)

tmp15201 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15200, V4589)


__e.TailApply(tmp15117, tmp15201)
return


}, 1)

tmp15202 := PrimHead(W4651)

__e.TailApply(tmp15116, tmp15202)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15205 := PrimTail(W4649)

tmp15206 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15205, V4589)


__e.TailApply(tmp15115, tmp15206)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15209 := PrimHead(W4649)

tmp15210 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15209, V4589)


__e.TailApply(tmp15114, tmp15210)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15213 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15214 := Call(__e, tmp15113, tmp15213)


ifres15112 = tmp15214


} else {
ifres15112 = False


}

__e.TailApply(tmp14558, ifres15112)
return


} else {
__e.Return(W4627)
return
}


}, 1)

tmp15300 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15218 Obj

if True == tmp15300 {
tmp15219 := MakeNative(func(__e *ControlFlow) {
W4628 := __e.Get(1)
_ = W4628
tmp15297 := PrimIsPair(W4628)

if True == tmp15297 {
tmp15220 := MakeNative(func(__e *ControlFlow) {
W4629 := __e.Get(1)
_ = W4629
tmp15293 := PrimEqual(W4629, symcons)

if True == tmp15293 {
tmp15221 := MakeNative(func(__e *ControlFlow) {
W4630 := __e.Get(1)
_ = W4630
tmp15289 := PrimIsPair(W4630)

if True == tmp15289 {
tmp15222 := MakeNative(func(__e *ControlFlow) {
W4631 := __e.Get(1)
_ = W4631
tmp15223 := MakeNative(func(__e *ControlFlow) {
W4632 := __e.Get(1)
_ = W4632
tmp15284 := PrimIsPair(W4632)

if True == tmp15284 {
tmp15224 := MakeNative(func(__e *ControlFlow) {
W4633 := __e.Get(1)
_ = W4633
tmp15225 := MakeNative(func(__e *ControlFlow) {
W4634 := __e.Get(1)
_ = W4634
tmp15279 := PrimEqual(W4634, Nil)

if True == tmp15279 {
tmp15226 := MakeNative(func(__e *ControlFlow) {
W4635 := __e.Get(1)
_ = W4635
tmp15227 := MakeNative(func(__e *ControlFlow) {
W4636 := __e.Get(1)
_ = W4636
tmp15271 := PrimIsPair(W4635)

if True == tmp15271 {
tmp15228 := MakeNative(func(__e *ControlFlow) {
W4638 := __e.Get(1)
_ = W4638
tmp15229 := MakeNative(func(__e *ControlFlow) {
W4639 := __e.Get(1)
_ = W4639
tmp15233 := PrimEqual(W4638, symlist)

if True == tmp15233 {
__e.TailApply(PrimFunc(symthaw), W4639)
return
} else {
tmp15231 := Call(__e, PrimFunc(symshen_4pvar_2), W4638)


if True == tmp15231 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4638, symlist, V4589, W4639)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15234 := MakeNative(func(__e *ControlFlow) {
tmp15235 := MakeNative(func(__e *ControlFlow) {
W4640 := __e.Get(1)
_ = W4640
tmp15236 := MakeNative(func(__e *ControlFlow) {
W4641 := __e.Get(1)
_ = W4641
tmp15256 := PrimIsPair(W4640)

if True == tmp15256 {
tmp15237 := MakeNative(func(__e *ControlFlow) {
W4643 := __e.Get(1)
_ = W4643
tmp15238 := MakeNative(func(__e *ControlFlow) {
W4644 := __e.Get(1)
_ = W4644
tmp15239 := MakeNative(func(__e *ControlFlow) {
W4645 := __e.Get(1)
_ = W4645
tmp15243 := PrimEqual(W4644, Nil)

if True == tmp15243 {
__e.TailApply(PrimFunc(symthaw), W4645)
return
} else {
tmp15241 := Call(__e, PrimFunc(symshen_4pvar_2), W4644)


if True == tmp15241 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4644, Nil, V4589, W4645)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15244 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4641, W4643)
return
}, 0)

__e.TailApply(tmp15239, tmp15244)
return


}, 1)

tmp15245 := PrimTail(W4640)

tmp15246 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15245, V4589)


__e.TailApply(tmp15238, tmp15246)
return


}, 1)

tmp15247 := PrimHead(W4640)

__e.TailApply(tmp15237, tmp15247)
return


} else {
tmp15254 := Call(__e, PrimFunc(symshen_4pvar_2), W4640)


if True == tmp15254 {
tmp15248 := MakeNative(func(__e *ControlFlow) {
W4646 := __e.Get(1)
_ = W4646
tmp15249 := PrimCons(W4646, Nil)

tmp15250 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4641, W4646)
return
}, 0)

tmp15251 := Call(__e, PrimFunc(symshen_4bind_b), W4640, tmp15249, V4589, tmp15250)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15251)
return


}, 1)

tmp15252 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15248, tmp15252)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15257 := MakeNative(func(__e *ControlFlow) {
Z4642 := __e.Get(1)
_ = Z4642
__e.TailApply(W4636, Z4642)
return
}, 1)

__e.TailApply(tmp15236, tmp15257)
return


}, 1)

tmp15258 := PrimTail(W4635)

tmp15259 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15258, V4589)


__e.TailApply(tmp15235, tmp15259)
return


}, 0)

__e.TailApply(tmp15229, tmp15234)
return


}, 1)

tmp15260 := PrimHead(W4635)

tmp15261 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15260, V4589)


__e.TailApply(tmp15228, tmp15261)
return


} else {
tmp15269 := Call(__e, PrimFunc(symshen_4pvar_2), W4635)


if True == tmp15269 {
tmp15262 := MakeNative(func(__e *ControlFlow) {
W4647 := __e.Get(1)
_ = W4647
tmp15263 := PrimCons(W4647, Nil)

tmp15264 := PrimCons(symlist, tmp15263)

tmp15265 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4636, W4647)
return
}, 0)

tmp15266 := Call(__e, PrimFunc(symshen_4bind_b), W4635, tmp15264, V4589, tmp15265)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15266)
return


}, 1)

tmp15267 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15262, tmp15267)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15272 := MakeNative(func(__e *ControlFlow) {
Z4637 := __e.Get(1)
_ = Z4637
tmp15273 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15273

tmp15274 := MakeNative(func(__e *ControlFlow) {
tmp15275 := PrimCons(Z4637, Nil)

tmp15276 := PrimCons(symlist, tmp15275)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4633, tmp15276, V4588, V4589, V4590, W4593, V4592)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4631, Z4637, V4588, V4589, V4590, W4593, tmp15274)
return


}, 1)

__e.TailApply(tmp15227, tmp15272)
return


}, 1)

tmp15277 := Call(__e, PrimFunc(symshen_4lazyderef), V4587, V4589)


__e.TailApply(tmp15226, tmp15277)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15280 := PrimTail(W4632)

tmp15281 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15280, V4589)


__e.TailApply(tmp15225, tmp15281)
return


}, 1)

tmp15282 := PrimHead(W4632)

__e.TailApply(tmp15224, tmp15282)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15285 := PrimTail(W4630)

tmp15286 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15285, V4589)


__e.TailApply(tmp15223, tmp15286)
return


}, 1)

tmp15287 := PrimHead(W4630)

__e.TailApply(tmp15222, tmp15287)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15290 := PrimTail(W4628)

tmp15291 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15290, V4589)


__e.TailApply(tmp15221, tmp15291)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15294 := PrimHead(W4628)

tmp15295 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15294, V4589)


__e.TailApply(tmp15220, tmp15295)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15298 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15299 := Call(__e, tmp15219, tmp15298)


ifres15218 = tmp15299


} else {
ifres15218 = False


}

__e.TailApply(tmp14557, ifres15218)
return


} else {
__e.Return(W4620)
return
}


}, 1)

tmp15331 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15303 Obj

if True == tmp15331 {
tmp15304 := MakeNative(func(__e *ControlFlow) {
W4621 := __e.Get(1)
_ = W4621
tmp15328 := PrimIsPair(W4621)

if True == tmp15328 {
tmp15305 := MakeNative(func(__e *ControlFlow) {
W4622 := __e.Get(1)
_ = W4622
tmp15306 := MakeNative(func(__e *ControlFlow) {
W4623 := __e.Get(1)
_ = W4623
tmp15323 := PrimIsPair(W4623)

if True == tmp15323 {
tmp15307 := MakeNative(func(__e *ControlFlow) {
W4624 := __e.Get(1)
_ = W4624
tmp15308 := MakeNative(func(__e *ControlFlow) {
W4625 := __e.Get(1)
_ = W4625
tmp15318 := PrimEqual(W4625, Nil)

if True == tmp15318 {
tmp15309 := MakeNative(func(__e *ControlFlow) {
W4626 := __e.Get(1)
_ = W4626
tmp15310 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15310

tmp15311 := PrimCons(V4587, Nil)

tmp15312 := PrimCons(sym_1_1_6, tmp15311)

tmp15313 := PrimCons(W4626, tmp15312)

tmp15314 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4624, W4626, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

tmp15315 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4622, tmp15313, V4588, V4589, V4590, W4593, tmp15314)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15315)
return


}, 1)

tmp15316 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15309, tmp15316)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15319 := PrimTail(W4623)

tmp15320 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15319, V4589)


__e.TailApply(tmp15308, tmp15320)
return


}, 1)

tmp15321 := PrimHead(W4623)

__e.TailApply(tmp15307, tmp15321)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15324 := PrimTail(W4621)

tmp15325 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15324, V4589)


__e.TailApply(tmp15306, tmp15325)
return


}, 1)

tmp15326 := PrimHead(W4621)

__e.TailApply(tmp15305, tmp15326)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15329 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15330 := Call(__e, tmp15304, tmp15329)


ifres15303 = tmp15330


} else {
ifres15303 = False


}

__e.TailApply(tmp14556, ifres15303)
return


} else {
__e.Return(W4613)
return
}


}, 1)

tmp15366 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15334 Obj

if True == tmp15366 {
tmp15335 := MakeNative(func(__e *ControlFlow) {
W4614 := __e.Get(1)
_ = W4614
tmp15363 := PrimIsPair(W4614)

if True == tmp15363 {
tmp15336 := MakeNative(func(__e *ControlFlow) {
W4615 := __e.Get(1)
_ = W4615
tmp15337 := MakeNative(func(__e *ControlFlow) {
W4616 := __e.Get(1)
_ = W4616
tmp15358 := PrimIsPair(W4616)

if True == tmp15358 {
tmp15338 := MakeNative(func(__e *ControlFlow) {
W4617 := __e.Get(1)
_ = W4617
tmp15339 := MakeNative(func(__e *ControlFlow) {
W4618 := __e.Get(1)
_ = W4618
tmp15353 := PrimEqual(W4618, Nil)

if True == tmp15353 {
tmp15340 := MakeNative(func(__e *ControlFlow) {
W4619 := __e.Get(1)
_ = W4619
tmp15341 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15341

tmp15342 := Call(__e, PrimFunc(symshen_4lazyderef), W4615, V4589)


tmp15343 := PrimIsPair(tmp15342)

tmp15344 := PrimNot(tmp15343)

tmp15345 := MakeNative(func(__e *ControlFlow) {
tmp15346 := PrimCons(V4587, Nil)

tmp15347 := PrimCons(sym_1_1_6, tmp15346)

tmp15348 := PrimCons(W4619, tmp15347)

tmp15349 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4617, W4619, V4588, V4589, V4590, W4593, V4592)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4615, tmp15348, V4589, V4590, W4593, tmp15349)
return


}, 0)

tmp15350 := Call(__e, PrimFunc(symwhen), tmp15344, V4589, V4590, W4593, tmp15345)


__e.TailApply(PrimFunc(symshen_4gc), V4589, tmp15350)
return


}, 1)

tmp15351 := Call(__e, PrimFunc(symshen_4newpv), V4589)


__e.TailApply(tmp15340, tmp15351)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15354 := PrimTail(W4616)

tmp15355 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15354, V4589)


__e.TailApply(tmp15339, tmp15355)
return


}, 1)

tmp15356 := PrimHead(W4616)

__e.TailApply(tmp15338, tmp15356)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15359 := PrimTail(W4614)

tmp15360 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15359, V4589)


__e.TailApply(tmp15337, tmp15360)
return


}, 1)

tmp15361 := PrimHead(W4614)

__e.TailApply(tmp15336, tmp15361)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15364 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15365 := Call(__e, tmp15335, tmp15364)


ifres15334 = tmp15365


} else {
ifres15334 = False


}

__e.TailApply(tmp14555, ifres15334)
return


} else {
__e.Return(W4607)
return
}


}, 1)

tmp15393 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15369 Obj

if True == tmp15393 {
tmp15370 := MakeNative(func(__e *ControlFlow) {
W4608 := __e.Get(1)
_ = W4608
tmp15390 := PrimIsPair(W4608)

if True == tmp15390 {
tmp15371 := MakeNative(func(__e *ControlFlow) {
W4609 := __e.Get(1)
_ = W4609
tmp15386 := PrimEqual(W4609, symfn)

if True == tmp15386 {
tmp15372 := MakeNative(func(__e *ControlFlow) {
W4610 := __e.Get(1)
_ = W4610
tmp15382 := PrimIsPair(W4610)

if True == tmp15382 {
tmp15373 := MakeNative(func(__e *ControlFlow) {
W4611 := __e.Get(1)
_ = W4611
tmp15374 := MakeNative(func(__e *ControlFlow) {
W4612 := __e.Get(1)
_ = W4612
tmp15377 := PrimEqual(W4612, Nil)

if True == tmp15377 {
tmp15375 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15375

__e.TailApply(PrimFunc(symshen_4lookupsig), W4611, V4587, V4589, V4590, W4593, V4592)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15378 := PrimTail(W4610)

tmp15379 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15378, V4589)


__e.TailApply(tmp15374, tmp15379)
return


}, 1)

tmp15380 := PrimHead(W4610)

__e.TailApply(tmp15373, tmp15380)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15383 := PrimTail(W4608)

tmp15384 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15383, V4589)


__e.TailApply(tmp15372, tmp15384)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15387 := PrimHead(W4608)

tmp15388 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15387, V4589)


__e.TailApply(tmp15371, tmp15388)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15391 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15392 := Call(__e, tmp15370, tmp15391)


ifres15369 = tmp15392


} else {
ifres15369 = False


}

__e.TailApply(tmp14554, ifres15369)
return


} else {
__e.Return(W4601)
return
}


}, 1)

tmp15426 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15396 Obj

if True == tmp15426 {
tmp15397 := MakeNative(func(__e *ControlFlow) {
W4602 := __e.Get(1)
_ = W4602
tmp15423 := PrimIsPair(W4602)

if True == tmp15423 {
tmp15398 := MakeNative(func(__e *ControlFlow) {
W4603 := __e.Get(1)
_ = W4603
tmp15419 := PrimEqual(W4603, symfn)

if True == tmp15419 {
tmp15399 := MakeNative(func(__e *ControlFlow) {
W4604 := __e.Get(1)
_ = W4604
tmp15415 := PrimIsPair(W4604)

if True == tmp15415 {
tmp15400 := MakeNative(func(__e *ControlFlow) {
W4605 := __e.Get(1)
_ = W4605
tmp15401 := MakeNative(func(__e *ControlFlow) {
W4606 := __e.Get(1)
_ = W4606
tmp15410 := PrimEqual(W4606, Nil)

if True == tmp15410 {
tmp15402 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15402

tmp15403 := Call(__e, PrimFunc(symshen_4deref), W4605, V4589)


tmp15404 := Call(__e, PrimFunc(symarity), tmp15403)


tmp15405 := PrimEqual(tmp15404, MakeNumber(0))

tmp15406 := MakeNative(func(__e *ControlFlow) {
tmp15407 := MakeNative(func(__e *ControlFlow) {
tmp15408 := PrimCons(W4605, Nil)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15408, V4587, V4588, V4589, V4590, W4593, V4592)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4589, V4590, W4593, tmp15407)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15405, V4589, V4590, W4593, tmp15406)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15411 := PrimTail(W4604)

tmp15412 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15411, V4589)


__e.TailApply(tmp15401, tmp15412)
return


}, 1)

tmp15413 := PrimHead(W4604)

__e.TailApply(tmp15400, tmp15413)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15416 := PrimTail(W4602)

tmp15417 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15416, V4589)


__e.TailApply(tmp15399, tmp15417)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15420 := PrimHead(W4602)

tmp15421 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15420, V4589)


__e.TailApply(tmp15398, tmp15421)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15424 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15425 := Call(__e, tmp15397, tmp15424)


ifres15396 = tmp15425


} else {
ifres15396 = False


}

__e.TailApply(tmp14553, ifres15396)
return


} else {
__e.Return(W4597)
return
}


}, 1)

tmp15445 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15429 Obj

if True == tmp15445 {
tmp15430 := MakeNative(func(__e *ControlFlow) {
W4598 := __e.Get(1)
_ = W4598
tmp15442 := PrimIsPair(W4598)

if True == tmp15442 {
tmp15431 := MakeNative(func(__e *ControlFlow) {
W4599 := __e.Get(1)
_ = W4599
tmp15432 := MakeNative(func(__e *ControlFlow) {
W4600 := __e.Get(1)
_ = W4600
tmp15437 := PrimEqual(W4600, Nil)

if True == tmp15437 {
tmp15433 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15433

tmp15434 := PrimCons(V4587, Nil)

tmp15435 := PrimCons(sym_1_1_6, tmp15434)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4599, tmp15435, V4589, V4590, W4593, V4592)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15438 := PrimTail(W4598)

tmp15439 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15438, V4589)


__e.TailApply(tmp15432, tmp15439)
return


}, 1)

tmp15440 := PrimHead(W4598)

__e.TailApply(tmp15431, tmp15440)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15443 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15444 := Call(__e, tmp15430, tmp15443)


ifres15429 = tmp15444


} else {
ifres15429 = False


}

__e.TailApply(tmp14552, ifres15429)
return


} else {
__e.Return(W4596)
return
}


}, 1)

tmp15451 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15448 Obj

if True == tmp15451 {
tmp15449 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15449

tmp15450 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4586, V4587, V4588, V4589, V4590, W4593, V4592)


ifres15448 = tmp15450


} else {
ifres15448 = False


}

__e.TailApply(tmp14551, ifres15448)
return


} else {
__e.Return(W4595)
return
}


}, 1)

tmp15461 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15454 Obj

if True == tmp15461 {
tmp15455 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15455

tmp15456 := Call(__e, PrimFunc(symshen_4lazyderef), V4586, V4589)


tmp15457 := PrimIsPair(tmp15456)

tmp15458 := PrimNot(tmp15457)

tmp15459 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4586, V4587, V4589, V4590, W4593, V4592)
return
}, 0)

tmp15460 := Call(__e, PrimFunc(symwhen), tmp15458, V4589, V4590, W4593, tmp15459)


ifres15454 = tmp15460


} else {
ifres15454 = False


}

__e.TailApply(tmp14550, ifres15454)
return


} else {
__e.Return(W4594)
return
}


}, 1)

tmp15473 := Call(__e, PrimFunc(symshen_4unlocked_2), V4590)


var ifres15464 Obj

if True == tmp15473 {
tmp15465 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15465

tmp15466 := PrimValue(symshen_4_dspy_d)

tmp15467 := MakeNative(func(__e *ControlFlow) {
tmp15468 := PrimIntern(MakeString(":"))

tmp15469 := PrimCons(V4587, Nil)

tmp15470 := PrimCons(tmp15468, tmp15469)

tmp15471 := PrimCons(V4586, tmp15470)

__e.TailApply(PrimFunc(symshen_4show), tmp15471, V4588, V4589, V4590, W4593, V4592)
return


}, 0)

tmp15472 := Call(__e, PrimFunc(symwhen), tmp15466, V4589, V4590, W4593, tmp15467)


ifres15464 = tmp15472


} else {
ifres15464 = False


}

__e.TailApply(tmp14549, ifres15464)
return


}, 1)

tmp15474 := PrimNumberAdd(V4591, MakeNumber(1))

__e.TailApply(tmp14548, tmp15474)
return


}, 7)

tmp15475 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp14547)


_ = tmp15475

tmp15476 := MakeNative(func(__e *ControlFlow) {
V4798 := __e.Get(1)
_ = V4798
V4799 := __e.Get(2)
_ = V4799
V4800 := __e.Get(3)
_ = V4800
V4801 := __e.Get(4)
_ = V4801
V4802 := __e.Get(5)
_ = V4802
V4803 := __e.Get(6)
_ = V4803
tmp15477 := MakeNative(func(__e *ControlFlow) {
W4804 := __e.Get(1)
_ = W4804
tmp15585 := PrimEqual(W4804, False)

if True == tmp15585 {
tmp15478 := MakeNative(func(__e *ControlFlow) {
W4807 := __e.Get(1)
_ = W4807
tmp15569 := PrimEqual(W4807, False)

if True == tmp15569 {
tmp15479 := MakeNative(func(__e *ControlFlow) {
W4810 := __e.Get(1)
_ = W4810
tmp15553 := PrimEqual(W4810, False)

if True == tmp15553 {
tmp15480 := MakeNative(func(__e *ControlFlow) {
W4813 := __e.Get(1)
_ = W4813
tmp15537 := PrimEqual(W4813, False)

if True == tmp15537 {
tmp15535 := Call(__e, PrimFunc(symshen_4unlocked_2), V4801)


if True == tmp15535 {
tmp15481 := MakeNative(func(__e *ControlFlow) {
W4816 := __e.Get(1)
_ = W4816
tmp15532 := PrimEqual(W4816, Nil)

if True == tmp15532 {
tmp15482 := MakeNative(func(__e *ControlFlow) {
W4817 := __e.Get(1)
_ = W4817
tmp15483 := MakeNative(func(__e *ControlFlow) {
W4818 := __e.Get(1)
_ = W4818
tmp15527 := PrimIsPair(W4817)

if True == tmp15527 {
tmp15484 := MakeNative(func(__e *ControlFlow) {
W4820 := __e.Get(1)
_ = W4820
tmp15485 := MakeNative(func(__e *ControlFlow) {
W4821 := __e.Get(1)
_ = W4821
tmp15489 := PrimEqual(W4820, symlist)

if True == tmp15489 {
__e.TailApply(PrimFunc(symthaw), W4821)
return
} else {
tmp15487 := Call(__e, PrimFunc(symshen_4pvar_2), W4820)


if True == tmp15487 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4820, symlist, V4800, W4821)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15490 := MakeNative(func(__e *ControlFlow) {
tmp15491 := MakeNative(func(__e *ControlFlow) {
W4822 := __e.Get(1)
_ = W4822
tmp15492 := MakeNative(func(__e *ControlFlow) {
W4823 := __e.Get(1)
_ = W4823
tmp15512 := PrimIsPair(W4822)

if True == tmp15512 {
tmp15493 := MakeNative(func(__e *ControlFlow) {
W4825 := __e.Get(1)
_ = W4825
tmp15494 := MakeNative(func(__e *ControlFlow) {
W4826 := __e.Get(1)
_ = W4826
tmp15495 := MakeNative(func(__e *ControlFlow) {
W4827 := __e.Get(1)
_ = W4827
tmp15499 := PrimEqual(W4826, Nil)

if True == tmp15499 {
__e.TailApply(PrimFunc(symthaw), W4827)
return
} else {
tmp15497 := Call(__e, PrimFunc(symshen_4pvar_2), W4826)


if True == tmp15497 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4826, Nil, V4800, W4827)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15500 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4823, W4825)
return
}, 0)

__e.TailApply(tmp15495, tmp15500)
return


}, 1)

tmp15501 := PrimTail(W4822)

tmp15502 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15501, V4800)


__e.TailApply(tmp15494, tmp15502)
return


}, 1)

tmp15503 := PrimHead(W4822)

__e.TailApply(tmp15493, tmp15503)
return


} else {
tmp15510 := Call(__e, PrimFunc(symshen_4pvar_2), W4822)


if True == tmp15510 {
tmp15504 := MakeNative(func(__e *ControlFlow) {
W4828 := __e.Get(1)
_ = W4828
tmp15505 := PrimCons(W4828, Nil)

tmp15506 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4823, W4828)
return
}, 0)

tmp15507 := Call(__e, PrimFunc(symshen_4bind_b), W4822, tmp15505, V4800, tmp15506)


__e.TailApply(PrimFunc(symshen_4gc), V4800, tmp15507)
return


}, 1)

tmp15508 := Call(__e, PrimFunc(symshen_4newpv), V4800)


__e.TailApply(tmp15504, tmp15508)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15513 := MakeNative(func(__e *ControlFlow) {
Z4824 := __e.Get(1)
_ = Z4824
__e.TailApply(W4818, Z4824)
return
}, 1)

__e.TailApply(tmp15492, tmp15513)
return


}, 1)

tmp15514 := PrimTail(W4817)

tmp15515 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15514, V4800)


__e.TailApply(tmp15491, tmp15515)
return


}, 0)

__e.TailApply(tmp15485, tmp15490)
return


}, 1)

tmp15516 := PrimHead(W4817)

tmp15517 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15516, V4800)


__e.TailApply(tmp15484, tmp15517)
return


} else {
tmp15525 := Call(__e, PrimFunc(symshen_4pvar_2), W4817)


if True == tmp15525 {
tmp15518 := MakeNative(func(__e *ControlFlow) {
W4829 := __e.Get(1)
_ = W4829
tmp15519 := PrimCons(W4829, Nil)

tmp15520 := PrimCons(symlist, tmp15519)

tmp15521 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4818, W4829)
return
}, 0)

tmp15522 := Call(__e, PrimFunc(symshen_4bind_b), W4817, tmp15520, V4800, tmp15521)


__e.TailApply(PrimFunc(symshen_4gc), V4800, tmp15522)
return


}, 1)

tmp15523 := Call(__e, PrimFunc(symshen_4newpv), V4800)


__e.TailApply(tmp15518, tmp15523)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15528 := MakeNative(func(__e *ControlFlow) {
Z4819 := __e.Get(1)
_ = Z4819
tmp15529 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15529

__e.TailApply(PrimFunc(symthaw), V4803)
return


}, 1)

__e.TailApply(tmp15483, tmp15528)
return


}, 1)

tmp15530 := Call(__e, PrimFunc(symshen_4lazyderef), V4799, V4800)


__e.TailApply(tmp15482, tmp15530)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15533 := Call(__e, PrimFunc(symshen_4lazyderef), V4798, V4800)


__e.TailApply(tmp15481, tmp15533)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4813)
return
}


}, 1)

tmp15551 := Call(__e, PrimFunc(symshen_4unlocked_2), V4801)


var ifres15538 Obj

if True == tmp15551 {
tmp15539 := MakeNative(func(__e *ControlFlow) {
W4814 := __e.Get(1)
_ = W4814
tmp15540 := MakeNative(func(__e *ControlFlow) {
W4815 := __e.Get(1)
_ = W4815
tmp15544 := PrimEqual(W4814, symsymbol)

if True == tmp15544 {
__e.TailApply(PrimFunc(symthaw), W4815)
return
} else {
tmp15542 := Call(__e, PrimFunc(symshen_4pvar_2), W4814)


if True == tmp15542 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4814, symsymbol, V4800, W4815)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15545 := MakeNative(func(__e *ControlFlow) {
tmp15546 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15546

tmp15547 := Call(__e, PrimFunc(symshen_4lazyderef), V4798, V4800)


tmp15548 := PrimIsSymbol(tmp15547)

__e.TailApply(PrimFunc(symwhen), tmp15548, V4800, V4801, V4802, V4803)
return


}, 0)

__e.TailApply(tmp15540, tmp15545)
return


}, 1)

tmp15549 := Call(__e, PrimFunc(symshen_4lazyderef), V4799, V4800)


tmp15550 := Call(__e, tmp15539, tmp15549)


ifres15538 = tmp15550


} else {
ifres15538 = False


}

__e.TailApply(tmp15480, ifres15538)
return


} else {
__e.Return(W4810)
return
}


}, 1)

tmp15567 := Call(__e, PrimFunc(symshen_4unlocked_2), V4801)


var ifres15554 Obj

if True == tmp15567 {
tmp15555 := MakeNative(func(__e *ControlFlow) {
W4811 := __e.Get(1)
_ = W4811
tmp15556 := MakeNative(func(__e *ControlFlow) {
W4812 := __e.Get(1)
_ = W4812
tmp15560 := PrimEqual(W4811, symstring)

if True == tmp15560 {
__e.TailApply(PrimFunc(symthaw), W4812)
return
} else {
tmp15558 := Call(__e, PrimFunc(symshen_4pvar_2), W4811)


if True == tmp15558 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4811, symstring, V4800, W4812)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15561 := MakeNative(func(__e *ControlFlow) {
tmp15562 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15562

tmp15563 := Call(__e, PrimFunc(symshen_4lazyderef), V4798, V4800)


tmp15564 := PrimIsString(tmp15563)

__e.TailApply(PrimFunc(symwhen), tmp15564, V4800, V4801, V4802, V4803)
return


}, 0)

__e.TailApply(tmp15556, tmp15561)
return


}, 1)

tmp15565 := Call(__e, PrimFunc(symshen_4lazyderef), V4799, V4800)


tmp15566 := Call(__e, tmp15555, tmp15565)


ifres15554 = tmp15566


} else {
ifres15554 = False


}

__e.TailApply(tmp15479, ifres15554)
return


} else {
__e.Return(W4807)
return
}


}, 1)

tmp15583 := Call(__e, PrimFunc(symshen_4unlocked_2), V4801)


var ifres15570 Obj

if True == tmp15583 {
tmp15571 := MakeNative(func(__e *ControlFlow) {
W4808 := __e.Get(1)
_ = W4808
tmp15572 := MakeNative(func(__e *ControlFlow) {
W4809 := __e.Get(1)
_ = W4809
tmp15576 := PrimEqual(W4808, symboolean)

if True == tmp15576 {
__e.TailApply(PrimFunc(symthaw), W4809)
return
} else {
tmp15574 := Call(__e, PrimFunc(symshen_4pvar_2), W4808)


if True == tmp15574 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4808, symboolean, V4800, W4809)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15577 := MakeNative(func(__e *ControlFlow) {
tmp15578 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15578

tmp15579 := Call(__e, PrimFunc(symshen_4lazyderef), V4798, V4800)


tmp15580 := Call(__e, PrimFunc(symboolean_2), tmp15579)


__e.TailApply(PrimFunc(symwhen), tmp15580, V4800, V4801, V4802, V4803)
return


}, 0)

__e.TailApply(tmp15572, tmp15577)
return


}, 1)

tmp15581 := Call(__e, PrimFunc(symshen_4lazyderef), V4799, V4800)


tmp15582 := Call(__e, tmp15571, tmp15581)


ifres15570 = tmp15582


} else {
ifres15570 = False


}

__e.TailApply(tmp15478, ifres15570)
return


} else {
__e.Return(W4804)
return
}


}, 1)

tmp15599 := Call(__e, PrimFunc(symshen_4unlocked_2), V4801)


var ifres15586 Obj

if True == tmp15599 {
tmp15587 := MakeNative(func(__e *ControlFlow) {
W4805 := __e.Get(1)
_ = W4805
tmp15588 := MakeNative(func(__e *ControlFlow) {
W4806 := __e.Get(1)
_ = W4806
tmp15592 := PrimEqual(W4805, symnumber)

if True == tmp15592 {
__e.TailApply(PrimFunc(symthaw), W4806)
return
} else {
tmp15590 := Call(__e, PrimFunc(symshen_4pvar_2), W4805)


if True == tmp15590 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4805, symnumber, V4800, W4806)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15593 := MakeNative(func(__e *ControlFlow) {
tmp15594 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15594

tmp15595 := Call(__e, PrimFunc(symshen_4lazyderef), V4798, V4800)


tmp15596 := PrimIsNumber(tmp15595)

__e.TailApply(PrimFunc(symwhen), tmp15596, V4800, V4801, V4802, V4803)
return


}, 0)

__e.TailApply(tmp15588, tmp15593)
return


}, 1)

tmp15597 := Call(__e, PrimFunc(symshen_4lazyderef), V4799, V4800)


tmp15598 := Call(__e, tmp15587, tmp15597)


ifres15586 = tmp15598


} else {
ifres15586 = False


}

__e.TailApply(tmp15477, ifres15586)
return


}, 6)

tmp15600 := Call(__e, ns2_1set, symshen_4primitive, tmp15476)


_ = tmp15600

tmp15601 := MakeNative(func(__e *ControlFlow) {
V4830 := __e.Get(1)
_ = V4830
V4831 := __e.Get(2)
_ = V4831
V4832 := __e.Get(3)
_ = V4832
V4833 := __e.Get(4)
_ = V4833
V4834 := __e.Get(5)
_ = V4834
V4835 := __e.Get(6)
_ = V4835
V4836 := __e.Get(7)
_ = V4836
tmp15602 := MakeNative(func(__e *ControlFlow) {
W4837 := __e.Get(1)
_ = W4837
tmp15613 := PrimEqual(W4837, False)

if True == tmp15613 {
tmp15611 := Call(__e, PrimFunc(symshen_4unlocked_2), V4834)


if True == tmp15611 {
tmp15603 := MakeNative(func(__e *ControlFlow) {
W4846 := __e.Get(1)
_ = W4846
tmp15608 := PrimIsPair(W4846)

if True == tmp15608 {
tmp15604 := MakeNative(func(__e *ControlFlow) {
W4847 := __e.Get(1)
_ = W4847
tmp15605 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15605

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V4830, V4831, W4847, V4833, V4834, V4835, V4836)
return


}, 1)

tmp15606 := PrimTail(W4846)

__e.TailApply(tmp15604, tmp15606)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15609 := Call(__e, PrimFunc(symshen_4lazyderef), V4832, V4833)


__e.TailApply(tmp15603, tmp15609)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4837)
return
}


}, 1)

tmp15655 := Call(__e, PrimFunc(symshen_4unlocked_2), V4834)


var ifres15614 Obj

if True == tmp15655 {
tmp15615 := MakeNative(func(__e *ControlFlow) {
W4838 := __e.Get(1)
_ = W4838
tmp15652 := PrimIsPair(W4838)

if True == tmp15652 {
tmp15616 := MakeNative(func(__e *ControlFlow) {
W4839 := __e.Get(1)
_ = W4839
tmp15648 := PrimIsPair(W4839)

if True == tmp15648 {
tmp15617 := MakeNative(func(__e *ControlFlow) {
W4840 := __e.Get(1)
_ = W4840
tmp15618 := MakeNative(func(__e *ControlFlow) {
W4841 := __e.Get(1)
_ = W4841
tmp15643 := PrimIsPair(W4841)

if True == tmp15643 {
tmp15619 := MakeNative(func(__e *ControlFlow) {
W4842 := __e.Get(1)
_ = W4842
tmp15620 := MakeNative(func(__e *ControlFlow) {
W4843 := __e.Get(1)
_ = W4843
tmp15638 := PrimIsPair(W4843)

if True == tmp15638 {
tmp15621 := MakeNative(func(__e *ControlFlow) {
W4844 := __e.Get(1)
_ = W4844
tmp15622 := MakeNative(func(__e *ControlFlow) {
W4845 := __e.Get(1)
_ = W4845
tmp15633 := PrimEqual(W4845, Nil)

if True == tmp15633 {
tmp15623 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15623

tmp15624 := Call(__e, PrimFunc(symshen_4deref), W4842, V4833)


tmp15625 := PrimIntern(MakeString(":"))

tmp15626 := PrimEqual(tmp15624, tmp15625)

tmp15627 := MakeNative(func(__e *ControlFlow) {
tmp15628 := Call(__e, PrimFunc(symshen_4deref), V4830, V4833)


tmp15629 := Call(__e, PrimFunc(symshen_4deref), W4840, V4833)


tmp15630 := PrimEqual(tmp15628, tmp15629)

tmp15631 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V4831, W4844, V4833, V4834, V4835, V4836)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15630, V4833, V4834, V4835, tmp15631)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15626, V4833, V4834, V4835, tmp15627)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15634 := PrimTail(W4843)

tmp15635 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15634, V4833)


__e.TailApply(tmp15622, tmp15635)
return


}, 1)

tmp15636 := PrimHead(W4843)

__e.TailApply(tmp15621, tmp15636)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15639 := PrimTail(W4841)

tmp15640 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15639, V4833)


__e.TailApply(tmp15620, tmp15640)
return


}, 1)

tmp15641 := PrimHead(W4841)

__e.TailApply(tmp15619, tmp15641)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15644 := PrimTail(W4839)

tmp15645 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15644, V4833)


__e.TailApply(tmp15618, tmp15645)
return


}, 1)

tmp15646 := PrimHead(W4839)

__e.TailApply(tmp15617, tmp15646)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15649 := PrimHead(W4838)

tmp15650 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15649, V4833)


__e.TailApply(tmp15616, tmp15650)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15653 := Call(__e, PrimFunc(symshen_4lazyderef), V4832, V4833)


tmp15654 := Call(__e, tmp15615, tmp15653)


ifres15614 = tmp15654


} else {
ifres15614 = False


}

__e.TailApply(tmp15602, ifres15614)
return


}, 7)

tmp15656 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15601)


_ = tmp15656

tmp15657 := MakeNative(func(__e *ControlFlow) {
V4848 := __e.Get(1)
_ = V4848
V4849 := __e.Get(2)
_ = V4849
V4850 := __e.Get(3)
_ = V4850
V4851 := __e.Get(4)
_ = V4851
V4852 := __e.Get(5)
_ = V4852
V4853 := __e.Get(6)
_ = V4853
tmp15662 := Call(__e, PrimFunc(symshen_4unlocked_2), V4851)


if True == tmp15662 {
tmp15658 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15658

tmp15659 := PrimValue(symshen_4_dsigf_d)

tmp15660 := Call(__e, PrimFunc(symassoc), V4848, tmp15659)


__e.TailApply(PrimFunc(symshen_4sigf), tmp15660, V4849, V4850, V4851, V4852, V4853)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15663 := Call(__e, ns2_1set, symshen_4lookupsig, tmp15657)


_ = tmp15663

tmp15664 := MakeNative(func(__e *ControlFlow) {
V4868 := __e.Get(1)
_ = V4868
V4869 := __e.Get(2)
_ = V4869
V4870 := __e.Get(3)
_ = V4870
V4871 := __e.Get(4)
_ = V4871
V4872 := __e.Get(5)
_ = V4872
V4873 := __e.Get(6)
_ = V4873
tmp15671 := PrimIsPair(V4868)

if True == tmp15671 {
tmp15665 := PrimTail(V4868)

tmp15666 := Call(__e, tmp15665, V4869)


tmp15667 := Call(__e, tmp15666, V4870)


tmp15668 := Call(__e, tmp15667, V4871)


tmp15669 := Call(__e, tmp15668, V4872)


__e.TailApply(tmp15669, V4873)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15672 := Call(__e, ns2_1set, symshen_4sigf, tmp15664)


_ = tmp15672

tmp15673 := MakeNative(func(__e *ControlFlow) {
V4874 := __e.Get(1)
_ = V4874
tmp15674 := MakeNative(func(__e *ControlFlow) {
W4875 := __e.Get(1)
_ = W4875
tmp15675 := MakeNative(func(__e *ControlFlow) {
W4876 := __e.Get(1)
_ = W4876
tmp15676 := MakeNative(func(__e *ControlFlow) {
W4877 := __e.Get(1)
_ = W4877
tmp15677 := MakeNative(func(__e *ControlFlow) {
W4878 := __e.Get(1)
_ = W4878
__e.Return(W4878)
return
}, 1)

tmp15678 := PrimValue(symshen_4_dgensym_d)

tmp15679 := PrimNumberAdd(MakeNumber(1), tmp15678)

tmp15680 := PrimSet(symshen_4_dgensym_d, tmp15679)

tmp15681 := PrimVectorSet(W4877, MakeNumber(2), tmp15680)

__e.TailApply(tmp15677, tmp15681)
return


}, 1)

tmp15682 := PrimVectorSet(W4876, MakeNumber(1), V4874)

__e.TailApply(tmp15676, tmp15682)
return


}, 1)

tmp15683 := PrimVectorSet(W4875, MakeNumber(0), symshen_4print_1freshterm)

__e.TailApply(tmp15675, tmp15683)
return


}, 1)

tmp15684 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp15674, tmp15684)
return


}, 1)

tmp15685 := Call(__e, ns2_1set, symshen_4freshterm, tmp15673)


_ = tmp15685

tmp15686 := MakeNative(func(__e *ControlFlow) {
V4879 := __e.Get(1)
_ = V4879
tmp15687 := PrimVectorGet(V4879, MakeNumber(1))

tmp15688 := PrimStr(tmp15687)

__e.Return(PrimStringConcat(MakeString("&&"), tmp15688))
return


}, 1)

tmp15689 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp15686)


_ = tmp15689

tmp15690 := MakeNative(func(__e *ControlFlow) {
V4880 := __e.Get(1)
_ = V4880
V4881 := __e.Get(2)
_ = V4881
V4882 := __e.Get(3)
_ = V4882
V4883 := __e.Get(4)
_ = V4883
V4884 := __e.Get(5)
_ = V4884
V4885 := __e.Get(6)
_ = V4885
V4886 := __e.Get(7)
_ = V4886
tmp15691 := MakeNative(func(__e *ControlFlow) {
W4887 := __e.Get(1)
_ = W4887
tmp15702 := PrimEqual(W4887, False)

if True == tmp15702 {
tmp15700 := Call(__e, PrimFunc(symshen_4unlocked_2), V4884)


if True == tmp15700 {
tmp15692 := MakeNative(func(__e *ControlFlow) {
W4891 := __e.Get(1)
_ = W4891
tmp15697 := PrimIsPair(W4891)

if True == tmp15697 {
tmp15693 := MakeNative(func(__e *ControlFlow) {
W4892 := __e.Get(1)
_ = W4892
tmp15694 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15694

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V4880, V4881, W4892, V4883, V4884, V4885, V4886)
return


}, 1)

tmp15695 := PrimTail(W4891)

__e.TailApply(tmp15693, tmp15695)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15698 := Call(__e, PrimFunc(symshen_4lazyderef), V4882, V4883)


__e.TailApply(tmp15692, tmp15698)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4887)
return
}


}, 1)

tmp15722 := Call(__e, PrimFunc(symshen_4unlocked_2), V4884)


var ifres15703 Obj

if True == tmp15722 {
tmp15704 := MakeNative(func(__e *ControlFlow) {
W4888 := __e.Get(1)
_ = W4888
tmp15719 := PrimIsPair(W4888)

if True == tmp15719 {
tmp15705 := MakeNative(func(__e *ControlFlow) {
W4889 := __e.Get(1)
_ = W4889
tmp15715 := PrimIsPair(W4889)

if True == tmp15715 {
tmp15706 := MakeNative(func(__e *ControlFlow) {
W4890 := __e.Get(1)
_ = W4890
tmp15707 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15707

tmp15708 := Call(__e, PrimFunc(symshen_4deref), W4890, V4883)


tmp15709 := Call(__e, PrimFunc(symshen_4deref), V4880, V4883)


tmp15710 := Call(__e, tmp15708, tmp15709)


tmp15711 := Call(__e, PrimFunc(symshen_4deref), V4881, V4883)


tmp15712 := Call(__e, tmp15710, tmp15711)


__e.TailApply(PrimFunc(symcall), tmp15712, V4883, V4884, V4885, V4886)
return


}, 1)

tmp15713 := PrimTail(W4889)

__e.TailApply(tmp15706, tmp15713)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15716 := PrimHead(W4888)

tmp15717 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15716, V4883)


__e.TailApply(tmp15705, tmp15717)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15720 := Call(__e, PrimFunc(symshen_4lazyderef), V4882, V4883)


tmp15721 := Call(__e, tmp15704, tmp15720)


ifres15703 = tmp15721


} else {
ifres15703 = False


}

__e.TailApply(tmp15691, ifres15703)
return


}, 7)

tmp15723 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp15690)


_ = tmp15723

tmp15724 := MakeNative(func(__e *ControlFlow) {
V4893 := __e.Get(1)
_ = V4893
V4894 := __e.Get(2)
_ = V4894
V4895 := __e.Get(3)
_ = V4895
V4896 := __e.Get(4)
_ = V4896
V4897 := __e.Get(5)
_ = V4897
V4898 := __e.Get(6)
_ = V4898
V4899 := __e.Get(7)
_ = V4899
tmp15725 := MakeNative(func(__e *ControlFlow) {
W4900 := __e.Get(1)
_ = W4900
tmp15726 := MakeNative(func(__e *ControlFlow) {
W4901 := __e.Get(1)
_ = W4901
tmp16156 := PrimEqual(W4901, False)

if True == tmp16156 {
tmp15727 := MakeNative(func(__e *ControlFlow) {
W4904 := __e.Get(1)
_ = W4904
tmp16056 := PrimEqual(W4904, False)

if True == tmp16056 {
tmp15728 := MakeNative(func(__e *ControlFlow) {
W4924 := __e.Get(1)
_ = W4924
tmp15951 := PrimEqual(W4924, False)

if True == tmp15951 {
tmp15729 := MakeNative(func(__e *ControlFlow) {
W4946 := __e.Get(1)
_ = W4946
tmp15870 := PrimEqual(W4946, False)

if True == tmp15870 {
tmp15730 := MakeNative(func(__e *ControlFlow) {
W4962 := __e.Get(1)
_ = W4962
tmp15770 := PrimEqual(W4962, False)

if True == tmp15770 {
tmp15731 := MakeNative(func(__e *ControlFlow) {
W4982 := __e.Get(1)
_ = W4982
tmp15733 := PrimEqual(W4982, False)

if True == tmp15733 {
__e.TailApply(PrimFunc(symshen_4unlock), V4897, W4900)
return
} else {
__e.Return(W4982)
return
}


}, 1)

tmp15768 := Call(__e, PrimFunc(symshen_4unlocked_2), V4897)


var ifres15734 Obj

if True == tmp15768 {
tmp15735 := MakeNative(func(__e *ControlFlow) {
W4983 := __e.Get(1)
_ = W4983
tmp15765 := PrimIsPair(W4983)

if True == tmp15765 {
tmp15736 := MakeNative(func(__e *ControlFlow) {
W4984 := __e.Get(1)
_ = W4984
tmp15737 := MakeNative(func(__e *ControlFlow) {
W4985 := __e.Get(1)
_ = W4985
tmp15738 := MakeNative(func(__e *ControlFlow) {
W4986 := __e.Get(1)
_ = W4986
tmp15739 := MakeNative(func(__e *ControlFlow) {
W4987 := __e.Get(1)
_ = W4987
tmp15757 := PrimIsPair(W4986)

if True == tmp15757 {
tmp15740 := MakeNative(func(__e *ControlFlow) {
W4990 := __e.Get(1)
_ = W4990
tmp15741 := MakeNative(func(__e *ControlFlow) {
W4991 := __e.Get(1)
_ = W4991
tmp15742 := Call(__e, W4987, W4990)


__e.TailApply(tmp15742, W4991)
return


}, 1)

tmp15743 := PrimTail(W4986)

__e.TailApply(tmp15741, tmp15743)
return


}, 1)

tmp15744 := PrimHead(W4986)

__e.TailApply(tmp15740, tmp15744)
return


} else {
tmp15755 := Call(__e, PrimFunc(symshen_4pvar_2), W4986)


if True == tmp15755 {
tmp15745 := MakeNative(func(__e *ControlFlow) {
W4992 := __e.Get(1)
_ = W4992
tmp15746 := MakeNative(func(__e *ControlFlow) {
W4993 := __e.Get(1)
_ = W4993
tmp15747 := PrimCons(W4992, W4993)

tmp15748 := MakeNative(func(__e *ControlFlow) {
tmp15749 := Call(__e, W4987, W4992)


__e.TailApply(tmp15749, W4993)
return


}, 0)

tmp15750 := Call(__e, PrimFunc(symshen_4bind_b), W4986, tmp15747, V4896, tmp15748)


__e.TailApply(PrimFunc(symshen_4gc), V4896, tmp15750)
return


}, 1)

tmp15751 := Call(__e, PrimFunc(symshen_4newpv), V4896)


tmp15752 := Call(__e, tmp15746, tmp15751)


__e.TailApply(PrimFunc(symshen_4gc), V4896, tmp15752)
return


}, 1)

tmp15753 := Call(__e, PrimFunc(symshen_4newpv), V4896)


__e.TailApply(tmp15745, tmp15753)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15758 := MakeNative(func(__e *ControlFlow) {
Z4988 := __e.Get(1)
_ = Z4988
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4989 := __e.Get(1)
_ = Z4989
tmp15759 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15759

tmp15760 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), W4985, Z4989, V4895, V4896, V4897, W4900, V4899)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z4988, W4984, V4896, V4897, W4900, tmp15760)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15739, tmp15758)
return


}, 1)

tmp15761 := Call(__e, PrimFunc(symshen_4lazyderef), V4894, V4896)


__e.TailApply(tmp15738, tmp15761)
return


}, 1)

tmp15762 := PrimTail(W4983)

__e.TailApply(tmp15737, tmp15762)
return


}, 1)

tmp15763 := PrimHead(W4983)

__e.TailApply(tmp15736, tmp15763)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15766 := Call(__e, PrimFunc(symshen_4lazyderef), V4893, V4896)


tmp15767 := Call(__e, tmp15735, tmp15766)


ifres15734 = tmp15767


} else {
ifres15734 = False


}

__e.TailApply(tmp15731, ifres15734)
return


} else {
__e.Return(W4962)
return
}


}, 1)

tmp15868 := Call(__e, PrimFunc(symshen_4unlocked_2), V4897)


var ifres15771 Obj

if True == tmp15868 {
tmp15772 := MakeNative(func(__e *ControlFlow) {
W4963 := __e.Get(1)
_ = W4963
tmp15865 := PrimIsPair(W4963)

if True == tmp15865 {
tmp15773 := MakeNative(func(__e *ControlFlow) {
W4964 := __e.Get(1)
_ = W4964
tmp15861 := PrimIsPair(W4964)

if True == tmp15861 {
tmp15774 := MakeNative(func(__e *ControlFlow) {
W4965 := __e.Get(1)
_ = W4965
tmp15857 := PrimIsPair(W4965)

if True == tmp15857 {
tmp15775 := MakeNative(func(__e *ControlFlow) {
W4966 := __e.Get(1)
_ = W4966
tmp15853 := PrimEqual(W4966, sym_8v)

if True == tmp15853 {
tmp15776 := MakeNative(func(__e *ControlFlow) {
W4967 := __e.Get(1)
_ = W4967
tmp15849 := PrimIsPair(W4967)

if True == tmp15849 {
tmp15777 := MakeNative(func(__e *ControlFlow) {
W4968 := __e.Get(1)
_ = W4968
tmp15778 := MakeNative(func(__e *ControlFlow) {
W4969 := __e.Get(1)
_ = W4969
tmp15844 := PrimIsPair(W4969)

if True == tmp15844 {
tmp15779 := MakeNative(func(__e *ControlFlow) {
W4970 := __e.Get(1)
_ = W4970
tmp15780 := MakeNative(func(__e *ControlFlow) {
W4971 := __e.Get(1)
_ = W4971
tmp15839 := PrimEqual(W4971, Nil)

if True == tmp15839 {
tmp15781 := MakeNative(func(__e *ControlFlow) {
W4972 := __e.Get(1)
_ = W4972
tmp15835 := PrimIsPair(W4972)

if True == tmp15835 {
tmp15782 := MakeNative(func(__e *ControlFlow) {
W4973 := __e.Get(1)
_ = W4973
tmp15783 := MakeNative(func(__e *ControlFlow) {
W4974 := __e.Get(1)
_ = W4974
tmp15830 := PrimIsPair(W4974)

if True == tmp15830 {
tmp15784 := MakeNative(func(__e *ControlFlow) {
W4975 := __e.Get(1)
_ = W4975
tmp15826 := PrimIsPair(W4975)

if True == tmp15826 {
tmp15785 := MakeNative(func(__e *ControlFlow) {
W4976 := __e.Get(1)
_ = W4976
tmp15822 := PrimEqual(W4976, symvector)

if True == tmp15822 {
tmp15786 := MakeNative(func(__e *ControlFlow) {
W4977 := __e.Get(1)
_ = W4977
tmp15818 := PrimIsPair(W4977)

if True == tmp15818 {
tmp15787 := MakeNative(func(__e *ControlFlow) {
W4978 := __e.Get(1)
_ = W4978
tmp15788 := MakeNative(func(__e *ControlFlow) {
W4979 := __e.Get(1)
_ = W4979
tmp15813 := PrimEqual(W4979, Nil)

if True == tmp15813 {
tmp15789 := MakeNative(func(__e *ControlFlow) {
W4980 := __e.Get(1)
_ = W4980
tmp15809 := PrimEqual(W4980, Nil)

if True == tmp15809 {
tmp15790 := MakeNative(func(__e *ControlFlow) {
W4981 := __e.Get(1)
_ = W4981
tmp15791 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15791

tmp15792 := Call(__e, PrimFunc(symshen_4deref), W4973, V4896)


tmp15793 := PrimIntern(MakeString(":"))

tmp15794 := PrimEqual(tmp15792, tmp15793)

tmp15795 := MakeNative(func(__e *ControlFlow) {
tmp15796 := MakeNative(func(__e *ControlFlow) {
tmp15797 := PrimCons(W4978, Nil)

tmp15798 := PrimCons(W4973, tmp15797)

tmp15799 := PrimCons(W4968, tmp15798)

tmp15800 := PrimCons(W4978, Nil)

tmp15801 := PrimCons(symvector, tmp15800)

tmp15802 := PrimCons(tmp15801, Nil)

tmp15803 := PrimCons(W4973, tmp15802)

tmp15804 := PrimCons(W4970, tmp15803)

tmp15805 := PrimCons(tmp15804, W4981)

tmp15806 := PrimCons(tmp15799, tmp15805)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15806, V4894, True, V4896, V4897, W4900, V4899)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4896, V4897, W4900, tmp15796)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15794, V4896, V4897, W4900, tmp15795)
return


}, 1)

tmp15807 := PrimTail(W4963)

__e.TailApply(tmp15790, tmp15807)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15810 := PrimTail(W4974)

tmp15811 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15810, V4896)


__e.TailApply(tmp15789, tmp15811)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15814 := PrimTail(W4977)

tmp15815 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15814, V4896)


__e.TailApply(tmp15788, tmp15815)
return


}, 1)

tmp15816 := PrimHead(W4977)

__e.TailApply(tmp15787, tmp15816)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15819 := PrimTail(W4975)

tmp15820 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15819, V4896)


__e.TailApply(tmp15786, tmp15820)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15823 := PrimHead(W4975)

tmp15824 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15823, V4896)


__e.TailApply(tmp15785, tmp15824)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15827 := PrimHead(W4974)

tmp15828 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15827, V4896)


__e.TailApply(tmp15784, tmp15828)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15831 := PrimTail(W4972)

tmp15832 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15831, V4896)


__e.TailApply(tmp15783, tmp15832)
return


}, 1)

tmp15833 := PrimHead(W4972)

__e.TailApply(tmp15782, tmp15833)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15836 := PrimTail(W4964)

tmp15837 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15836, V4896)


__e.TailApply(tmp15781, tmp15837)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15840 := PrimTail(W4969)

tmp15841 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15840, V4896)


__e.TailApply(tmp15780, tmp15841)
return


}, 1)

tmp15842 := PrimHead(W4969)

__e.TailApply(tmp15779, tmp15842)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15845 := PrimTail(W4967)

tmp15846 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15845, V4896)


__e.TailApply(tmp15778, tmp15846)
return


}, 1)

tmp15847 := PrimHead(W4967)

__e.TailApply(tmp15777, tmp15847)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15850 := PrimTail(W4965)

tmp15851 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15850, V4896)


__e.TailApply(tmp15776, tmp15851)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15854 := PrimHead(W4965)

tmp15855 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15854, V4896)


__e.TailApply(tmp15775, tmp15855)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15858 := PrimHead(W4964)

tmp15859 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15858, V4896)


__e.TailApply(tmp15774, tmp15859)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15862 := PrimHead(W4963)

tmp15863 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15862, V4896)


__e.TailApply(tmp15773, tmp15863)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15866 := Call(__e, PrimFunc(symshen_4lazyderef), V4893, V4896)


tmp15867 := Call(__e, tmp15772, tmp15866)


ifres15771 = tmp15867


} else {
ifres15771 = False


}

__e.TailApply(tmp15730, ifres15771)
return


} else {
__e.Return(W4946)
return
}


}, 1)

tmp15949 := Call(__e, PrimFunc(symshen_4unlocked_2), V4897)


var ifres15871 Obj

if True == tmp15949 {
tmp15872 := MakeNative(func(__e *ControlFlow) {
W4947 := __e.Get(1)
_ = W4947
tmp15946 := PrimIsPair(W4947)

if True == tmp15946 {
tmp15873 := MakeNative(func(__e *ControlFlow) {
W4948 := __e.Get(1)
_ = W4948
tmp15942 := PrimIsPair(W4948)

if True == tmp15942 {
tmp15874 := MakeNative(func(__e *ControlFlow) {
W4949 := __e.Get(1)
_ = W4949
tmp15938 := PrimIsPair(W4949)

if True == tmp15938 {
tmp15875 := MakeNative(func(__e *ControlFlow) {
W4950 := __e.Get(1)
_ = W4950
tmp15934 := PrimEqual(W4950, sym_8s)

if True == tmp15934 {
tmp15876 := MakeNative(func(__e *ControlFlow) {
W4951 := __e.Get(1)
_ = W4951
tmp15930 := PrimIsPair(W4951)

if True == tmp15930 {
tmp15877 := MakeNative(func(__e *ControlFlow) {
W4952 := __e.Get(1)
_ = W4952
tmp15878 := MakeNative(func(__e *ControlFlow) {
W4953 := __e.Get(1)
_ = W4953
tmp15925 := PrimIsPair(W4953)

if True == tmp15925 {
tmp15879 := MakeNative(func(__e *ControlFlow) {
W4954 := __e.Get(1)
_ = W4954
tmp15880 := MakeNative(func(__e *ControlFlow) {
W4955 := __e.Get(1)
_ = W4955
tmp15920 := PrimEqual(W4955, Nil)

if True == tmp15920 {
tmp15881 := MakeNative(func(__e *ControlFlow) {
W4956 := __e.Get(1)
_ = W4956
tmp15916 := PrimIsPair(W4956)

if True == tmp15916 {
tmp15882 := MakeNative(func(__e *ControlFlow) {
W4957 := __e.Get(1)
_ = W4957
tmp15883 := MakeNative(func(__e *ControlFlow) {
W4958 := __e.Get(1)
_ = W4958
tmp15911 := PrimIsPair(W4958)

if True == tmp15911 {
tmp15884 := MakeNative(func(__e *ControlFlow) {
W4959 := __e.Get(1)
_ = W4959
tmp15907 := PrimEqual(W4959, symstring)

if True == tmp15907 {
tmp15885 := MakeNative(func(__e *ControlFlow) {
W4960 := __e.Get(1)
_ = W4960
tmp15903 := PrimEqual(W4960, Nil)

if True == tmp15903 {
tmp15886 := MakeNative(func(__e *ControlFlow) {
W4961 := __e.Get(1)
_ = W4961
tmp15887 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15887

tmp15888 := Call(__e, PrimFunc(symshen_4deref), W4957, V4896)


tmp15889 := PrimIntern(MakeString(":"))

tmp15890 := PrimEqual(tmp15888, tmp15889)

tmp15891 := MakeNative(func(__e *ControlFlow) {
tmp15892 := MakeNative(func(__e *ControlFlow) {
tmp15893 := PrimCons(symstring, Nil)

tmp15894 := PrimCons(W4957, tmp15893)

tmp15895 := PrimCons(W4952, tmp15894)

tmp15896 := PrimCons(symstring, Nil)

tmp15897 := PrimCons(W4957, tmp15896)

tmp15898 := PrimCons(W4954, tmp15897)

tmp15899 := PrimCons(tmp15898, W4961)

tmp15900 := PrimCons(tmp15895, tmp15899)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15900, V4894, True, V4896, V4897, W4900, V4899)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4896, V4897, W4900, tmp15892)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15890, V4896, V4897, W4900, tmp15891)
return


}, 1)

tmp15901 := PrimTail(W4947)

__e.TailApply(tmp15886, tmp15901)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15904 := PrimTail(W4958)

tmp15905 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15904, V4896)


__e.TailApply(tmp15885, tmp15905)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15908 := PrimHead(W4958)

tmp15909 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15908, V4896)


__e.TailApply(tmp15884, tmp15909)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15912 := PrimTail(W4956)

tmp15913 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15912, V4896)


__e.TailApply(tmp15883, tmp15913)
return


}, 1)

tmp15914 := PrimHead(W4956)

__e.TailApply(tmp15882, tmp15914)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15917 := PrimTail(W4948)

tmp15918 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15917, V4896)


__e.TailApply(tmp15881, tmp15918)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15921 := PrimTail(W4953)

tmp15922 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15921, V4896)


__e.TailApply(tmp15880, tmp15922)
return


}, 1)

tmp15923 := PrimHead(W4953)

__e.TailApply(tmp15879, tmp15923)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15926 := PrimTail(W4951)

tmp15927 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15926, V4896)


__e.TailApply(tmp15878, tmp15927)
return


}, 1)

tmp15928 := PrimHead(W4951)

__e.TailApply(tmp15877, tmp15928)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15931 := PrimTail(W4949)

tmp15932 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15931, V4896)


__e.TailApply(tmp15876, tmp15932)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15935 := PrimHead(W4949)

tmp15936 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15935, V4896)


__e.TailApply(tmp15875, tmp15936)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15939 := PrimHead(W4948)

tmp15940 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15939, V4896)


__e.TailApply(tmp15874, tmp15940)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15943 := PrimHead(W4947)

tmp15944 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15943, V4896)


__e.TailApply(tmp15873, tmp15944)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15947 := Call(__e, PrimFunc(symshen_4lazyderef), V4893, V4896)


tmp15948 := Call(__e, tmp15872, tmp15947)


ifres15871 = tmp15948


} else {
ifres15871 = False


}

__e.TailApply(tmp15729, ifres15871)
return


} else {
__e.Return(W4924)
return
}


}, 1)

tmp16054 := Call(__e, PrimFunc(symshen_4unlocked_2), V4897)


var ifres15952 Obj

if True == tmp16054 {
tmp15953 := MakeNative(func(__e *ControlFlow) {
W4925 := __e.Get(1)
_ = W4925
tmp16051 := PrimIsPair(W4925)

if True == tmp16051 {
tmp15954 := MakeNative(func(__e *ControlFlow) {
W4926 := __e.Get(1)
_ = W4926
tmp16047 := PrimIsPair(W4926)

if True == tmp16047 {
tmp15955 := MakeNative(func(__e *ControlFlow) {
W4927 := __e.Get(1)
_ = W4927
tmp16043 := PrimIsPair(W4927)

if True == tmp16043 {
tmp15956 := MakeNative(func(__e *ControlFlow) {
W4928 := __e.Get(1)
_ = W4928
tmp16039 := PrimEqual(W4928, sym_8p)

if True == tmp16039 {
tmp15957 := MakeNative(func(__e *ControlFlow) {
W4929 := __e.Get(1)
_ = W4929
tmp16035 := PrimIsPair(W4929)

if True == tmp16035 {
tmp15958 := MakeNative(func(__e *ControlFlow) {
W4930 := __e.Get(1)
_ = W4930
tmp15959 := MakeNative(func(__e *ControlFlow) {
W4931 := __e.Get(1)
_ = W4931
tmp16030 := PrimIsPair(W4931)

if True == tmp16030 {
tmp15960 := MakeNative(func(__e *ControlFlow) {
W4932 := __e.Get(1)
_ = W4932
tmp15961 := MakeNative(func(__e *ControlFlow) {
W4933 := __e.Get(1)
_ = W4933
tmp16025 := PrimEqual(W4933, Nil)

if True == tmp16025 {
tmp15962 := MakeNative(func(__e *ControlFlow) {
W4934 := __e.Get(1)
_ = W4934
tmp16021 := PrimIsPair(W4934)

if True == tmp16021 {
tmp15963 := MakeNative(func(__e *ControlFlow) {
W4935 := __e.Get(1)
_ = W4935
tmp15964 := MakeNative(func(__e *ControlFlow) {
W4936 := __e.Get(1)
_ = W4936
tmp16016 := PrimIsPair(W4936)

if True == tmp16016 {
tmp15965 := MakeNative(func(__e *ControlFlow) {
W4937 := __e.Get(1)
_ = W4937
tmp16012 := PrimIsPair(W4937)

if True == tmp16012 {
tmp15966 := MakeNative(func(__e *ControlFlow) {
W4938 := __e.Get(1)
_ = W4938
tmp15967 := MakeNative(func(__e *ControlFlow) {
W4939 := __e.Get(1)
_ = W4939
tmp16007 := PrimIsPair(W4939)

if True == tmp16007 {
tmp15968 := MakeNative(func(__e *ControlFlow) {
W4940 := __e.Get(1)
_ = W4940
tmp16003 := PrimEqual(W4940, sym_d)

if True == tmp16003 {
tmp15969 := MakeNative(func(__e *ControlFlow) {
W4941 := __e.Get(1)
_ = W4941
tmp15999 := PrimIsPair(W4941)

if True == tmp15999 {
tmp15970 := MakeNative(func(__e *ControlFlow) {
W4942 := __e.Get(1)
_ = W4942
tmp15971 := MakeNative(func(__e *ControlFlow) {
W4943 := __e.Get(1)
_ = W4943
tmp15994 := PrimEqual(W4943, Nil)

if True == tmp15994 {
tmp15972 := MakeNative(func(__e *ControlFlow) {
W4944 := __e.Get(1)
_ = W4944
tmp15990 := PrimEqual(W4944, Nil)

if True == tmp15990 {
tmp15973 := MakeNative(func(__e *ControlFlow) {
W4945 := __e.Get(1)
_ = W4945
tmp15974 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15974

tmp15975 := Call(__e, PrimFunc(symshen_4deref), W4935, V4896)


tmp15976 := PrimIntern(MakeString(":"))

tmp15977 := PrimEqual(tmp15975, tmp15976)

tmp15978 := MakeNative(func(__e *ControlFlow) {
tmp15979 := MakeNative(func(__e *ControlFlow) {
tmp15980 := PrimCons(W4938, Nil)

tmp15981 := PrimCons(W4935, tmp15980)

tmp15982 := PrimCons(W4930, tmp15981)

tmp15983 := PrimCons(W4942, Nil)

tmp15984 := PrimCons(W4935, tmp15983)

tmp15985 := PrimCons(W4932, tmp15984)

tmp15986 := PrimCons(tmp15985, W4945)

tmp15987 := PrimCons(tmp15982, tmp15986)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15987, V4894, True, V4896, V4897, W4900, V4899)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4896, V4897, W4900, tmp15979)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15977, V4896, V4897, W4900, tmp15978)
return


}, 1)

tmp15988 := PrimTail(W4925)

__e.TailApply(tmp15973, tmp15988)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15991 := PrimTail(W4936)

tmp15992 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15991, V4896)


__e.TailApply(tmp15972, tmp15992)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15995 := PrimTail(W4941)

tmp15996 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15995, V4896)


__e.TailApply(tmp15971, tmp15996)
return


}, 1)

tmp15997 := PrimHead(W4941)

__e.TailApply(tmp15970, tmp15997)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16000 := PrimTail(W4939)

tmp16001 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16000, V4896)


__e.TailApply(tmp15969, tmp16001)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16004 := PrimHead(W4939)

tmp16005 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16004, V4896)


__e.TailApply(tmp15968, tmp16005)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16008 := PrimTail(W4937)

tmp16009 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16008, V4896)


__e.TailApply(tmp15967, tmp16009)
return


}, 1)

tmp16010 := PrimHead(W4937)

__e.TailApply(tmp15966, tmp16010)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16013 := PrimHead(W4936)

tmp16014 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16013, V4896)


__e.TailApply(tmp15965, tmp16014)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16017 := PrimTail(W4934)

tmp16018 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16017, V4896)


__e.TailApply(tmp15964, tmp16018)
return


}, 1)

tmp16019 := PrimHead(W4934)

__e.TailApply(tmp15963, tmp16019)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16022 := PrimTail(W4926)

tmp16023 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16022, V4896)


__e.TailApply(tmp15962, tmp16023)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16026 := PrimTail(W4931)

tmp16027 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16026, V4896)


__e.TailApply(tmp15961, tmp16027)
return


}, 1)

tmp16028 := PrimHead(W4931)

__e.TailApply(tmp15960, tmp16028)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16031 := PrimTail(W4929)

tmp16032 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16031, V4896)


__e.TailApply(tmp15959, tmp16032)
return


}, 1)

tmp16033 := PrimHead(W4929)

__e.TailApply(tmp15958, tmp16033)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16036 := PrimTail(W4927)

tmp16037 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16036, V4896)


__e.TailApply(tmp15957, tmp16037)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16040 := PrimHead(W4927)

tmp16041 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16040, V4896)


__e.TailApply(tmp15956, tmp16041)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16044 := PrimHead(W4926)

tmp16045 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16044, V4896)


__e.TailApply(tmp15955, tmp16045)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16048 := PrimHead(W4925)

tmp16049 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16048, V4896)


__e.TailApply(tmp15954, tmp16049)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16052 := Call(__e, PrimFunc(symshen_4lazyderef), V4893, V4896)


tmp16053 := Call(__e, tmp15953, tmp16052)


ifres15952 = tmp16053


} else {
ifres15952 = False


}

__e.TailApply(tmp15728, ifres15952)
return


} else {
__e.Return(W4904)
return
}


}, 1)

tmp16154 := Call(__e, PrimFunc(symshen_4unlocked_2), V4897)


var ifres16057 Obj

if True == tmp16154 {
tmp16058 := MakeNative(func(__e *ControlFlow) {
W4905 := __e.Get(1)
_ = W4905
tmp16151 := PrimIsPair(W4905)

if True == tmp16151 {
tmp16059 := MakeNative(func(__e *ControlFlow) {
W4906 := __e.Get(1)
_ = W4906
tmp16147 := PrimIsPair(W4906)

if True == tmp16147 {
tmp16060 := MakeNative(func(__e *ControlFlow) {
W4907 := __e.Get(1)
_ = W4907
tmp16143 := PrimIsPair(W4907)

if True == tmp16143 {
tmp16061 := MakeNative(func(__e *ControlFlow) {
W4908 := __e.Get(1)
_ = W4908
tmp16139 := PrimEqual(W4908, symcons)

if True == tmp16139 {
tmp16062 := MakeNative(func(__e *ControlFlow) {
W4909 := __e.Get(1)
_ = W4909
tmp16135 := PrimIsPair(W4909)

if True == tmp16135 {
tmp16063 := MakeNative(func(__e *ControlFlow) {
W4910 := __e.Get(1)
_ = W4910
tmp16064 := MakeNative(func(__e *ControlFlow) {
W4911 := __e.Get(1)
_ = W4911
tmp16130 := PrimIsPair(W4911)

if True == tmp16130 {
tmp16065 := MakeNative(func(__e *ControlFlow) {
W4912 := __e.Get(1)
_ = W4912
tmp16066 := MakeNative(func(__e *ControlFlow) {
W4913 := __e.Get(1)
_ = W4913
tmp16125 := PrimEqual(W4913, Nil)

if True == tmp16125 {
tmp16067 := MakeNative(func(__e *ControlFlow) {
W4914 := __e.Get(1)
_ = W4914
tmp16121 := PrimIsPair(W4914)

if True == tmp16121 {
tmp16068 := MakeNative(func(__e *ControlFlow) {
W4915 := __e.Get(1)
_ = W4915
tmp16069 := MakeNative(func(__e *ControlFlow) {
W4916 := __e.Get(1)
_ = W4916
tmp16116 := PrimIsPair(W4916)

if True == tmp16116 {
tmp16070 := MakeNative(func(__e *ControlFlow) {
W4917 := __e.Get(1)
_ = W4917
tmp16112 := PrimIsPair(W4917)

if True == tmp16112 {
tmp16071 := MakeNative(func(__e *ControlFlow) {
W4918 := __e.Get(1)
_ = W4918
tmp16108 := PrimEqual(W4918, symlist)

if True == tmp16108 {
tmp16072 := MakeNative(func(__e *ControlFlow) {
W4919 := __e.Get(1)
_ = W4919
tmp16104 := PrimIsPair(W4919)

if True == tmp16104 {
tmp16073 := MakeNative(func(__e *ControlFlow) {
W4920 := __e.Get(1)
_ = W4920
tmp16074 := MakeNative(func(__e *ControlFlow) {
W4921 := __e.Get(1)
_ = W4921
tmp16099 := PrimEqual(W4921, Nil)

if True == tmp16099 {
tmp16075 := MakeNative(func(__e *ControlFlow) {
W4922 := __e.Get(1)
_ = W4922
tmp16095 := PrimEqual(W4922, Nil)

if True == tmp16095 {
tmp16076 := MakeNative(func(__e *ControlFlow) {
W4923 := __e.Get(1)
_ = W4923
tmp16077 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16077

tmp16078 := Call(__e, PrimFunc(symshen_4deref), W4915, V4896)


tmp16079 := PrimIntern(MakeString(":"))

tmp16080 := PrimEqual(tmp16078, tmp16079)

tmp16081 := MakeNative(func(__e *ControlFlow) {
tmp16082 := MakeNative(func(__e *ControlFlow) {
tmp16083 := PrimCons(W4920, Nil)

tmp16084 := PrimCons(W4915, tmp16083)

tmp16085 := PrimCons(W4910, tmp16084)

tmp16086 := PrimCons(W4920, Nil)

tmp16087 := PrimCons(symlist, tmp16086)

tmp16088 := PrimCons(tmp16087, Nil)

tmp16089 := PrimCons(W4915, tmp16088)

tmp16090 := PrimCons(W4912, tmp16089)

tmp16091 := PrimCons(tmp16090, W4923)

tmp16092 := PrimCons(tmp16085, tmp16091)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16092, V4894, True, V4896, V4897, W4900, V4899)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4896, V4897, W4900, tmp16082)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16080, V4896, V4897, W4900, tmp16081)
return


}, 1)

tmp16093 := PrimTail(W4905)

__e.TailApply(tmp16076, tmp16093)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16096 := PrimTail(W4916)

tmp16097 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16096, V4896)


__e.TailApply(tmp16075, tmp16097)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16100 := PrimTail(W4919)

tmp16101 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16100, V4896)


__e.TailApply(tmp16074, tmp16101)
return


}, 1)

tmp16102 := PrimHead(W4919)

__e.TailApply(tmp16073, tmp16102)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16105 := PrimTail(W4917)

tmp16106 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16105, V4896)


__e.TailApply(tmp16072, tmp16106)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16109 := PrimHead(W4917)

tmp16110 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16109, V4896)


__e.TailApply(tmp16071, tmp16110)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16113 := PrimHead(W4916)

tmp16114 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16113, V4896)


__e.TailApply(tmp16070, tmp16114)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16117 := PrimTail(W4914)

tmp16118 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16117, V4896)


__e.TailApply(tmp16069, tmp16118)
return


}, 1)

tmp16119 := PrimHead(W4914)

__e.TailApply(tmp16068, tmp16119)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16122 := PrimTail(W4906)

tmp16123 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16122, V4896)


__e.TailApply(tmp16067, tmp16123)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16126 := PrimTail(W4911)

tmp16127 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16126, V4896)


__e.TailApply(tmp16066, tmp16127)
return


}, 1)

tmp16128 := PrimHead(W4911)

__e.TailApply(tmp16065, tmp16128)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16131 := PrimTail(W4909)

tmp16132 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16131, V4896)


__e.TailApply(tmp16064, tmp16132)
return


}, 1)

tmp16133 := PrimHead(W4909)

__e.TailApply(tmp16063, tmp16133)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16136 := PrimTail(W4907)

tmp16137 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16136, V4896)


__e.TailApply(tmp16062, tmp16137)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16140 := PrimHead(W4907)

tmp16141 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16140, V4896)


__e.TailApply(tmp16061, tmp16141)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16144 := PrimHead(W4906)

tmp16145 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16144, V4896)


__e.TailApply(tmp16060, tmp16145)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16148 := PrimHead(W4905)

tmp16149 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16148, V4896)


__e.TailApply(tmp16059, tmp16149)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16152 := Call(__e, PrimFunc(symshen_4lazyderef), V4893, V4896)


tmp16153 := Call(__e, tmp16058, tmp16152)


ifres16057 = tmp16153


} else {
ifres16057 = False


}

__e.TailApply(tmp15727, ifres16057)
return


} else {
__e.Return(W4901)
return
}


}, 1)

tmp16169 := Call(__e, PrimFunc(symshen_4unlocked_2), V4897)


var ifres16157 Obj

if True == tmp16169 {
tmp16158 := MakeNative(func(__e *ControlFlow) {
W4902 := __e.Get(1)
_ = W4902
tmp16166 := PrimEqual(W4902, Nil)

if True == tmp16166 {
tmp16159 := MakeNative(func(__e *ControlFlow) {
W4903 := __e.Get(1)
_ = W4903
tmp16163 := PrimEqual(W4903, True)

if True == tmp16163 {
tmp16160 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16160

tmp16161 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V4894, Nil, V4896, V4897, W4900, V4899)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4896, V4897, W4900, tmp16161)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16164 := Call(__e, PrimFunc(symshen_4lazyderef), V4895, V4896)


__e.TailApply(tmp16159, tmp16164)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16167 := Call(__e, PrimFunc(symshen_4lazyderef), V4893, V4896)


tmp16168 := Call(__e, tmp16158, tmp16167)


ifres16157 = tmp16168


} else {
ifres16157 = False


}

__e.TailApply(tmp15726, ifres16157)
return


}, 1)

tmp16170 := PrimNumberAdd(V4898, MakeNumber(1))

__e.TailApply(tmp15725, tmp16170)
return


}, 7)

tmp16171 := Call(__e, ns2_1set, symshen_4l_1rules, tmp15724)


_ = tmp16171

tmp16172 := MakeNative(func(__e *ControlFlow) {
V4994 := __e.Get(1)
_ = V4994
V4995 := __e.Get(2)
_ = V4995
V4996 := __e.Get(3)
_ = V4996
V4997 := __e.Get(4)
_ = V4997
V4998 := __e.Get(5)
_ = V4998
V4999 := __e.Get(6)
_ = V4999
tmp16173 := MakeNative(func(__e *ControlFlow) {
W5000 := __e.Get(1)
_ = W5000
tmp16174 := MakeNative(func(__e *ControlFlow) {
W5001 := __e.Get(1)
_ = W5001
tmp16176 := PrimEqual(W5001, False)

if True == tmp16176 {
__e.TailApply(PrimFunc(symshen_4unlock), V4997, W5000)
return
} else {
__e.Return(W5001)
return
}


}, 1)

tmp16224 := Call(__e, PrimFunc(symshen_4unlocked_2), V4997)


var ifres16177 Obj

if True == tmp16224 {
tmp16178 := MakeNative(func(__e *ControlFlow) {
W5002 := __e.Get(1)
_ = W5002
tmp16221 := PrimIsPair(W5002)

if True == tmp16221 {
tmp16179 := MakeNative(func(__e *ControlFlow) {
W5003 := __e.Get(1)
_ = W5003
tmp16217 := PrimEqual(W5003, symdefine)

if True == tmp16217 {
tmp16180 := MakeNative(func(__e *ControlFlow) {
W5004 := __e.Get(1)
_ = W5004
tmp16213 := PrimIsPair(W5004)

if True == tmp16213 {
tmp16181 := MakeNative(func(__e *ControlFlow) {
W5005 := __e.Get(1)
_ = W5005
tmp16182 := MakeNative(func(__e *ControlFlow) {
W5006 := __e.Get(1)
_ = W5006
tmp16183 := MakeNative(func(__e *ControlFlow) {
W5007 := __e.Get(1)
_ = W5007
tmp16184 := MakeNative(func(__e *ControlFlow) {
W5008 := __e.Get(1)
_ = W5008
tmp16185 := MakeNative(func(__e *ControlFlow) {
W5009 := __e.Get(1)
_ = W5009
tmp16186 := MakeNative(func(__e *ControlFlow) {
W5010 := __e.Get(1)
_ = W5010
tmp16187 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16187

tmp16188 := MakeNative(func(__e *ControlFlow) {
tmp16189 := PrimCons(W5005, W5006)

tmp16190 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16189)


tmp16191 := MakeNative(func(__e *ControlFlow) {
tmp16192 := Call(__e, PrimFunc(symshen_4lazyderef), W5007, V4996)


tmp16193 := Call(__e, PrimFunc(symfst), tmp16192)


tmp16194 := MakeNative(func(__e *ControlFlow) {
tmp16195 := Call(__e, PrimFunc(symshen_4lazyderef), W5007, V4996)


tmp16196 := Call(__e, PrimFunc(symsnd), tmp16195)


tmp16197 := MakeNative(func(__e *ControlFlow) {
tmp16198 := Call(__e, PrimFunc(symshen_4deref), W5010, V4996)


tmp16199 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16198)


tmp16200 := MakeNative(func(__e *ControlFlow) {
tmp16201 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), W5010, V4995, V4996, V4997, W5000, V4999)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), W5005, W5008, W5009, MakeNumber(1), V4996, V4997, W5000, tmp16201)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5009, tmp16199, V4996, V4997, W5000, tmp16200)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5008, tmp16196, V4996, V4997, W5000, tmp16197)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5010, tmp16193, V4996, V4997, W5000, tmp16194)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5007, tmp16190, V4996, V4997, W5000, tmp16191)
return


}, 0)

tmp16202 := Call(__e, PrimFunc(symshen_4cut), V4996, V4997, W5000, tmp16188)


__e.TailApply(PrimFunc(symshen_4gc), V4996, tmp16202)
return


}, 1)

tmp16203 := Call(__e, PrimFunc(symshen_4newpv), V4996)


tmp16204 := Call(__e, tmp16186, tmp16203)


__e.TailApply(PrimFunc(symshen_4gc), V4996, tmp16204)
return


}, 1)

tmp16205 := Call(__e, PrimFunc(symshen_4newpv), V4996)


tmp16206 := Call(__e, tmp16185, tmp16205)


__e.TailApply(PrimFunc(symshen_4gc), V4996, tmp16206)
return


}, 1)

tmp16207 := Call(__e, PrimFunc(symshen_4newpv), V4996)


tmp16208 := Call(__e, tmp16184, tmp16207)


__e.TailApply(PrimFunc(symshen_4gc), V4996, tmp16208)
return


}, 1)

tmp16209 := Call(__e, PrimFunc(symshen_4newpv), V4996)


__e.TailApply(tmp16183, tmp16209)
return


}, 1)

tmp16210 := PrimTail(W5004)

__e.TailApply(tmp16182, tmp16210)
return


}, 1)

tmp16211 := PrimHead(W5004)

__e.TailApply(tmp16181, tmp16211)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16214 := PrimTail(W5002)

tmp16215 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16214, V4996)


__e.TailApply(tmp16180, tmp16215)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16218 := PrimHead(W5002)

tmp16219 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16218, V4996)


__e.TailApply(tmp16179, tmp16219)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16222 := Call(__e, PrimFunc(symshen_4lazyderef), V4994, V4996)


tmp16223 := Call(__e, tmp16178, tmp16222)


ifres16177 = tmp16223


} else {
ifres16177 = False


}

__e.TailApply(tmp16174, ifres16177)
return


}, 1)

tmp16225 := PrimNumberAdd(V4998, MakeNumber(1))

__e.TailApply(tmp16173, tmp16225)
return


}, 6)

tmp16226 := Call(__e, ns2_1set, symshen_4t_d, tmp16172)


_ = tmp16226

tmp16227 := MakeNative(func(__e *ControlFlow) {
V5011 := __e.Get(1)
_ = V5011
tmp16228 := MakeNative(func(__e *ControlFlow) {
Z5012 := __e.Get(1)
_ = Z5012
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5012)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16228, V5011)
return


}, 1)

tmp16229 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16227)


_ = tmp16229

tmp16230 := MakeNative(func(__e *ControlFlow) {
V5013 := __e.Get(1)
_ = V5013
tmp16231 := MakeNative(func(__e *ControlFlow) {
W5014 := __e.Get(1)
_ = W5014
tmp16233 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5014)


if True == tmp16233 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5014)
return
}


}, 1)

tmp16266 := PrimIsPair(V5013)

var ifres16234 Obj

if True == tmp16266 {
tmp16235 := MakeNative(func(__e *ControlFlow) {
W5015 := __e.Get(1)
_ = W5015
tmp16262 := Call(__e, PrimFunc(symshen_4hds_a_2), W5015, sym_i)


if True == tmp16262 {
tmp16236 := MakeNative(func(__e *ControlFlow) {
W5016 := __e.Get(1)
_ = W5016
tmp16237 := MakeNative(func(__e *ControlFlow) {
W5017 := __e.Get(1)
_ = W5017
tmp16258 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5017)


if True == tmp16258 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16238 := MakeNative(func(__e *ControlFlow) {
W5018 := __e.Get(1)
_ = W5018
tmp16239 := MakeNative(func(__e *ControlFlow) {
W5019 := __e.Get(1)
_ = W5019
tmp16254 := Call(__e, PrimFunc(symshen_4hds_a_2), W5019, sym_j)


if True == tmp16254 {
tmp16240 := MakeNative(func(__e *ControlFlow) {
W5020 := __e.Get(1)
_ = W5020
tmp16241 := MakeNative(func(__e *ControlFlow) {
W5021 := __e.Get(1)
_ = W5021
tmp16250 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5021)


if True == tmp16250 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16242 := MakeNative(func(__e *ControlFlow) {
W5022 := __e.Get(1)
_ = W5022
tmp16243 := MakeNative(func(__e *ControlFlow) {
W5023 := __e.Get(1)
_ = W5023
tmp16244 := MakeNative(func(__e *ControlFlow) {
W5024 := __e.Get(1)
_ = W5024
__e.TailApply(PrimFunc(sym_8p), W5024, W5022)
return
}, 1)

tmp16245 := Call(__e, PrimFunc(symshen_4rectify_1type), W5018)


tmp16246 := Call(__e, tmp16244, tmp16245)


__e.TailApply(PrimFunc(symshen_4comb), W5023, tmp16246)
return


}, 1)

tmp16247 := Call(__e, PrimFunc(symshen_4in_1_6), W5021)


__e.TailApply(tmp16243, tmp16247)
return


}, 1)

tmp16248 := Call(__e, PrimFunc(symshen_4_5_1out), W5021)


__e.TailApply(tmp16242, tmp16248)
return


}


}, 1)

tmp16251 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5020)


__e.TailApply(tmp16241, tmp16251)
return


}, 1)

tmp16252 := Call(__e, PrimFunc(symtail), W5019)


__e.TailApply(tmp16240, tmp16252)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16255 := Call(__e, PrimFunc(symshen_4in_1_6), W5017)


__e.TailApply(tmp16239, tmp16255)
return


}, 1)

tmp16256 := Call(__e, PrimFunc(symshen_4_5_1out), W5017)


__e.TailApply(tmp16238, tmp16256)
return


}


}, 1)

tmp16259 := Call(__e, PrimFunc(symshen_4_5signature_6), W5016)


__e.TailApply(tmp16237, tmp16259)
return


}, 1)

tmp16260 := Call(__e, PrimFunc(symtail), W5015)


__e.TailApply(tmp16236, tmp16260)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16263 := Call(__e, PrimFunc(symtail), V5013)


tmp16264 := Call(__e, tmp16235, tmp16263)


ifres16234 = tmp16264


} else {
tmp16265 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16234 = tmp16265


}

__e.TailApply(tmp16231, ifres16234)
return


}, 1)

tmp16267 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16230)


_ = tmp16267

tmp16268 := MakeNative(func(__e *ControlFlow) {
V5025 := __e.Get(1)
_ = V5025
tmp16269 := MakeNative(func(__e *ControlFlow) {
W5026 := __e.Get(1)
_ = W5026
tmp16270 := MakeNative(func(__e *ControlFlow) {
W5027 := __e.Get(1)
_ = W5027
__e.TailApply(PrimFunc(symshen_4freshen_1type), W5027, V5025)
return
}, 1)

tmp16271 := MakeNative(func(__e *ControlFlow) {
Z5028 := __e.Get(1)
_ = Z5028
tmp16272 := Call(__e, PrimFunc(symconcat), sym_e, Z5028)


tmp16273 := Call(__e, PrimFunc(symshen_4freshterm), tmp16272)


__e.Return(PrimCons(Z5028, tmp16273))
return


}, 1)

tmp16274 := Call(__e, PrimFunc(symmap), tmp16271, W5026)


__e.TailApply(tmp16270, tmp16274)
return


}, 1)

tmp16275 := Call(__e, PrimFunc(symshen_4extract_1vars), V5025)


__e.TailApply(tmp16269, tmp16275)
return


}, 1)

tmp16276 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16268)


_ = tmp16276

tmp16277 := MakeNative(func(__e *ControlFlow) {
V5029 := __e.Get(1)
_ = V5029
V5030 := __e.Get(2)
_ = V5030
tmp16291 := PrimEqual(Nil, V5029)

if True == tmp16291 {
__e.Return(V5030)
return
} else {
tmp16289 := PrimIsPair(V5029)

var ifres16285 Obj

if True == tmp16289 {
tmp16287 := PrimHead(V5029)

tmp16288 := PrimIsPair(tmp16287)

var ifres16286 Obj

if True == tmp16288 {
ifres16286 = True


} else {
ifres16286 = False


}

ifres16285 = ifres16286


} else {
ifres16285 = False


}

if True == ifres16285 {
tmp16278 := PrimTail(V5029)

tmp16279 := PrimHead(V5029)

tmp16280 := PrimTail(tmp16279)

tmp16281 := PrimHead(V5029)

tmp16282 := PrimHead(tmp16281)

tmp16283 := Call(__e, PrimFunc(symsubst), tmp16280, tmp16282, V5030)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16278, tmp16283)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen-type")))
return
}


}


}, 2)

tmp16292 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16277)


_ = tmp16292

tmp16293 := MakeNative(func(__e *ControlFlow) {
V5031 := __e.Get(1)
_ = V5031
tmp16294 := MakeNative(func(__e *ControlFlow) {
W5032 := __e.Get(1)
_ = W5032
tmp16309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5032)


if True == tmp16309 {
tmp16295 := MakeNative(func(__e *ControlFlow) {
W5039 := __e.Get(1)
_ = W5039
tmp16297 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5039)


if True == tmp16297 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5039)
return
}


}, 1)

tmp16298 := MakeNative(func(__e *ControlFlow) {
W5040 := __e.Get(1)
_ = W5040
tmp16305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5040)


if True == tmp16305 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16299 := MakeNative(func(__e *ControlFlow) {
W5041 := __e.Get(1)
_ = W5041
tmp16300 := MakeNative(func(__e *ControlFlow) {
W5042 := __e.Get(1)
_ = W5042
tmp16301 := PrimCons(W5041, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W5042, tmp16301)
return


}, 1)

tmp16302 := Call(__e, PrimFunc(symshen_4in_1_6), W5040)


__e.TailApply(tmp16300, tmp16302)
return


}, 1)

tmp16303 := Call(__e, PrimFunc(symshen_4_5_1out), W5040)


__e.TailApply(tmp16299, tmp16303)
return


}


}, 1)

tmp16306 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5031)


tmp16307 := Call(__e, tmp16298, tmp16306)


__e.TailApply(tmp16295, tmp16307)
return


} else {
__e.Return(W5032)
return
}


}, 1)

tmp16310 := MakeNative(func(__e *ControlFlow) {
W5033 := __e.Get(1)
_ = W5033
tmp16325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5033)


if True == tmp16325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16311 := MakeNative(func(__e *ControlFlow) {
W5034 := __e.Get(1)
_ = W5034
tmp16312 := MakeNative(func(__e *ControlFlow) {
W5035 := __e.Get(1)
_ = W5035
tmp16313 := MakeNative(func(__e *ControlFlow) {
W5036 := __e.Get(1)
_ = W5036
tmp16320 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5036)


if True == tmp16320 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16314 := MakeNative(func(__e *ControlFlow) {
W5037 := __e.Get(1)
_ = W5037
tmp16315 := MakeNative(func(__e *ControlFlow) {
W5038 := __e.Get(1)
_ = W5038
tmp16316 := PrimCons(W5034, W5037)

__e.TailApply(PrimFunc(symshen_4comb), W5038, tmp16316)
return


}, 1)

tmp16317 := Call(__e, PrimFunc(symshen_4in_1_6), W5036)


__e.TailApply(tmp16315, tmp16317)
return


}, 1)

tmp16318 := Call(__e, PrimFunc(symshen_4_5_1out), W5036)


__e.TailApply(tmp16314, tmp16318)
return


}


}, 1)

tmp16321 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5035)


__e.TailApply(tmp16313, tmp16321)
return


}, 1)

tmp16322 := Call(__e, PrimFunc(symshen_4in_1_6), W5033)


__e.TailApply(tmp16312, tmp16322)
return


}, 1)

tmp16323 := Call(__e, PrimFunc(symshen_4_5_1out), W5033)


__e.TailApply(tmp16311, tmp16323)
return


}


}, 1)

tmp16326 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5031)


tmp16327 := Call(__e, tmp16310, tmp16326)


__e.TailApply(tmp16294, tmp16327)
return


}, 1)

tmp16328 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16293)


_ = tmp16328

tmp16329 := MakeNative(func(__e *ControlFlow) {
V5043 := __e.Get(1)
_ = V5043
tmp16330 := MakeNative(func(__e *ControlFlow) {
W5044 := __e.Get(1)
_ = W5044
tmp16416 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5044)


if True == tmp16416 {
tmp16331 := MakeNative(func(__e *ControlFlow) {
W5054 := __e.Get(1)
_ = W5054
tmp16380 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5054)


if True == tmp16380 {
tmp16332 := MakeNative(func(__e *ControlFlow) {
W5064 := __e.Get(1)
_ = W5064
tmp16357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5064)


if True == tmp16357 {
tmp16333 := MakeNative(func(__e *ControlFlow) {
W5071 := __e.Get(1)
_ = W5071
tmp16335 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5071)


if True == tmp16335 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5071)
return
}


}, 1)

tmp16336 := MakeNative(func(__e *ControlFlow) {
W5072 := __e.Get(1)
_ = W5072
tmp16353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5072)


if True == tmp16353 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16337 := MakeNative(func(__e *ControlFlow) {
W5073 := __e.Get(1)
_ = W5073
tmp16338 := MakeNative(func(__e *ControlFlow) {
W5074 := __e.Get(1)
_ = W5074
tmp16349 := Call(__e, PrimFunc(symshen_4hds_a_2), W5074, sym_1_6)


if True == tmp16349 {
tmp16339 := MakeNative(func(__e *ControlFlow) {
W5075 := __e.Get(1)
_ = W5075
tmp16346 := PrimIsPair(W5075)

if True == tmp16346 {
tmp16340 := MakeNative(func(__e *ControlFlow) {
W5076 := __e.Get(1)
_ = W5076
tmp16341 := MakeNative(func(__e *ControlFlow) {
W5077 := __e.Get(1)
_ = W5077
tmp16342 := Call(__e, PrimFunc(sym_8p), W5073, W5076)


__e.TailApply(PrimFunc(symshen_4comb), W5077, tmp16342)
return


}, 1)

tmp16343 := Call(__e, PrimFunc(symtail), W5075)


__e.TailApply(tmp16341, tmp16343)
return


}, 1)

tmp16344 := Call(__e, PrimFunc(symhead), W5075)


__e.TailApply(tmp16340, tmp16344)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16347 := Call(__e, PrimFunc(symtail), W5074)


__e.TailApply(tmp16339, tmp16347)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16350 := Call(__e, PrimFunc(symshen_4in_1_6), W5072)


__e.TailApply(tmp16338, tmp16350)
return


}, 1)

tmp16351 := Call(__e, PrimFunc(symshen_4_5_1out), W5072)


__e.TailApply(tmp16337, tmp16351)
return


}


}, 1)

tmp16354 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5043)


tmp16355 := Call(__e, tmp16336, tmp16354)


__e.TailApply(tmp16333, tmp16355)
return


} else {
__e.Return(W5064)
return
}


}, 1)

tmp16358 := MakeNative(func(__e *ControlFlow) {
W5065 := __e.Get(1)
_ = W5065
tmp16376 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5065)


if True == tmp16376 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16359 := MakeNative(func(__e *ControlFlow) {
W5066 := __e.Get(1)
_ = W5066
tmp16360 := MakeNative(func(__e *ControlFlow) {
W5067 := __e.Get(1)
_ = W5067
tmp16372 := Call(__e, PrimFunc(symshen_4hds_a_2), W5067, sym_5_1)


if True == tmp16372 {
tmp16361 := MakeNative(func(__e *ControlFlow) {
W5068 := __e.Get(1)
_ = W5068
tmp16369 := PrimIsPair(W5068)

if True == tmp16369 {
tmp16362 := MakeNative(func(__e *ControlFlow) {
W5069 := __e.Get(1)
_ = W5069
tmp16363 := MakeNative(func(__e *ControlFlow) {
W5070 := __e.Get(1)
_ = W5070
tmp16364 := Call(__e, PrimFunc(symshen_4correct), W5069)


tmp16365 := Call(__e, PrimFunc(sym_8p), W5066, tmp16364)


__e.TailApply(PrimFunc(symshen_4comb), W5070, tmp16365)
return


}, 1)

tmp16366 := Call(__e, PrimFunc(symtail), W5068)


__e.TailApply(tmp16363, tmp16366)
return


}, 1)

tmp16367 := Call(__e, PrimFunc(symhead), W5068)


__e.TailApply(tmp16362, tmp16367)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16370 := Call(__e, PrimFunc(symtail), W5067)


__e.TailApply(tmp16361, tmp16370)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16373 := Call(__e, PrimFunc(symshen_4in_1_6), W5065)


__e.TailApply(tmp16360, tmp16373)
return


}, 1)

tmp16374 := Call(__e, PrimFunc(symshen_4_5_1out), W5065)


__e.TailApply(tmp16359, tmp16374)
return


}


}, 1)

tmp16377 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5043)


tmp16378 := Call(__e, tmp16358, tmp16377)


__e.TailApply(tmp16332, tmp16378)
return


} else {
__e.Return(W5054)
return
}


}, 1)

tmp16381 := MakeNative(func(__e *ControlFlow) {
W5055 := __e.Get(1)
_ = W5055
tmp16412 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5055)


if True == tmp16412 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16382 := MakeNative(func(__e *ControlFlow) {
W5056 := __e.Get(1)
_ = W5056
tmp16383 := MakeNative(func(__e *ControlFlow) {
W5057 := __e.Get(1)
_ = W5057
tmp16408 := Call(__e, PrimFunc(symshen_4hds_a_2), W5057, sym_5_1)


if True == tmp16408 {
tmp16384 := MakeNative(func(__e *ControlFlow) {
W5058 := __e.Get(1)
_ = W5058
tmp16405 := PrimIsPair(W5058)

if True == tmp16405 {
tmp16385 := MakeNative(func(__e *ControlFlow) {
W5059 := __e.Get(1)
_ = W5059
tmp16386 := MakeNative(func(__e *ControlFlow) {
W5060 := __e.Get(1)
_ = W5060
tmp16401 := Call(__e, PrimFunc(symshen_4hds_a_2), W5060, symwhere)


if True == tmp16401 {
tmp16387 := MakeNative(func(__e *ControlFlow) {
W5061 := __e.Get(1)
_ = W5061
tmp16398 := PrimIsPair(W5061)

if True == tmp16398 {
tmp16388 := MakeNative(func(__e *ControlFlow) {
W5062 := __e.Get(1)
_ = W5062
tmp16389 := MakeNative(func(__e *ControlFlow) {
W5063 := __e.Get(1)
_ = W5063
tmp16390 := PrimCons(W5059, Nil)

tmp16391 := PrimCons(W5062, tmp16390)

tmp16392 := PrimCons(symwhere, tmp16391)

tmp16393 := Call(__e, PrimFunc(symshen_4correct), tmp16392)


tmp16394 := Call(__e, PrimFunc(sym_8p), W5056, tmp16393)


__e.TailApply(PrimFunc(symshen_4comb), W5063, tmp16394)
return


}, 1)

tmp16395 := Call(__e, PrimFunc(symtail), W5061)


__e.TailApply(tmp16389, tmp16395)
return


}, 1)

tmp16396 := Call(__e, PrimFunc(symhead), W5061)


__e.TailApply(tmp16388, tmp16396)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16399 := Call(__e, PrimFunc(symtail), W5060)


__e.TailApply(tmp16387, tmp16399)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16402 := Call(__e, PrimFunc(symtail), W5058)


__e.TailApply(tmp16386, tmp16402)
return


}, 1)

tmp16403 := Call(__e, PrimFunc(symhead), W5058)


__e.TailApply(tmp16385, tmp16403)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16406 := Call(__e, PrimFunc(symtail), W5057)


__e.TailApply(tmp16384, tmp16406)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16409 := Call(__e, PrimFunc(symshen_4in_1_6), W5055)


__e.TailApply(tmp16383, tmp16409)
return


}, 1)

tmp16410 := Call(__e, PrimFunc(symshen_4_5_1out), W5055)


__e.TailApply(tmp16382, tmp16410)
return


}


}, 1)

tmp16413 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5043)


tmp16414 := Call(__e, tmp16381, tmp16413)


__e.TailApply(tmp16331, tmp16414)
return


} else {
__e.Return(W5044)
return
}


}, 1)

tmp16417 := MakeNative(func(__e *ControlFlow) {
W5045 := __e.Get(1)
_ = W5045
tmp16447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5045)


if True == tmp16447 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16418 := MakeNative(func(__e *ControlFlow) {
W5046 := __e.Get(1)
_ = W5046
tmp16419 := MakeNative(func(__e *ControlFlow) {
W5047 := __e.Get(1)
_ = W5047
tmp16443 := Call(__e, PrimFunc(symshen_4hds_a_2), W5047, sym_1_6)


if True == tmp16443 {
tmp16420 := MakeNative(func(__e *ControlFlow) {
W5048 := __e.Get(1)
_ = W5048
tmp16440 := PrimIsPair(W5048)

if True == tmp16440 {
tmp16421 := MakeNative(func(__e *ControlFlow) {
W5049 := __e.Get(1)
_ = W5049
tmp16422 := MakeNative(func(__e *ControlFlow) {
W5050 := __e.Get(1)
_ = W5050
tmp16436 := Call(__e, PrimFunc(symshen_4hds_a_2), W5050, symwhere)


if True == tmp16436 {
tmp16423 := MakeNative(func(__e *ControlFlow) {
W5051 := __e.Get(1)
_ = W5051
tmp16433 := PrimIsPair(W5051)

if True == tmp16433 {
tmp16424 := MakeNative(func(__e *ControlFlow) {
W5052 := __e.Get(1)
_ = W5052
tmp16425 := MakeNative(func(__e *ControlFlow) {
W5053 := __e.Get(1)
_ = W5053
tmp16426 := PrimCons(W5049, Nil)

tmp16427 := PrimCons(W5052, tmp16426)

tmp16428 := PrimCons(symwhere, tmp16427)

tmp16429 := Call(__e, PrimFunc(sym_8p), W5046, tmp16428)


__e.TailApply(PrimFunc(symshen_4comb), W5053, tmp16429)
return


}, 1)

tmp16430 := Call(__e, PrimFunc(symtail), W5051)


__e.TailApply(tmp16425, tmp16430)
return


}, 1)

tmp16431 := Call(__e, PrimFunc(symhead), W5051)


__e.TailApply(tmp16424, tmp16431)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16434 := Call(__e, PrimFunc(symtail), W5050)


__e.TailApply(tmp16423, tmp16434)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16437 := Call(__e, PrimFunc(symtail), W5048)


__e.TailApply(tmp16422, tmp16437)
return


}, 1)

tmp16438 := Call(__e, PrimFunc(symhead), W5048)


__e.TailApply(tmp16421, tmp16438)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16441 := Call(__e, PrimFunc(symtail), W5047)


__e.TailApply(tmp16420, tmp16441)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16444 := Call(__e, PrimFunc(symshen_4in_1_6), W5045)


__e.TailApply(tmp16419, tmp16444)
return


}, 1)

tmp16445 := Call(__e, PrimFunc(symshen_4_5_1out), W5045)


__e.TailApply(tmp16418, tmp16445)
return


}


}, 1)

tmp16448 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5043)


tmp16449 := Call(__e, tmp16417, tmp16448)


__e.TailApply(tmp16330, tmp16449)
return


}, 1)

tmp16450 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16329)


_ = tmp16450

tmp16451 := MakeNative(func(__e *ControlFlow) {
V5078 := __e.Get(1)
_ = V5078
tmp16599 := PrimIsPair(V5078)

var ifres16543 Obj

if True == tmp16599 {
tmp16597 := PrimHead(V5078)

tmp16598 := PrimEqual(symwhere, tmp16597)

var ifres16545 Obj

if True == tmp16598 {
tmp16595 := PrimTail(V5078)

tmp16596 := PrimIsPair(tmp16595)

var ifres16547 Obj

if True == tmp16596 {
tmp16592 := PrimTail(V5078)

tmp16593 := PrimTail(tmp16592)

tmp16594 := PrimIsPair(tmp16593)

var ifres16549 Obj

if True == tmp16594 {
tmp16588 := PrimTail(V5078)

tmp16589 := PrimTail(tmp16588)

tmp16590 := PrimHead(tmp16589)

tmp16591 := PrimIsPair(tmp16590)

var ifres16551 Obj

if True == tmp16591 {
tmp16583 := PrimTail(V5078)

tmp16584 := PrimTail(tmp16583)

tmp16585 := PrimHead(tmp16584)

tmp16586 := PrimHead(tmp16585)

tmp16587 := PrimEqual(symfail_1if, tmp16586)

var ifres16553 Obj

if True == tmp16587 {
tmp16578 := PrimTail(V5078)

tmp16579 := PrimTail(tmp16578)

tmp16580 := PrimHead(tmp16579)

tmp16581 := PrimTail(tmp16580)

tmp16582 := PrimIsPair(tmp16581)

var ifres16555 Obj

if True == tmp16582 {
tmp16572 := PrimTail(V5078)

tmp16573 := PrimTail(tmp16572)

tmp16574 := PrimHead(tmp16573)

tmp16575 := PrimTail(tmp16574)

tmp16576 := PrimTail(tmp16575)

tmp16577 := PrimIsPair(tmp16576)

var ifres16557 Obj

if True == tmp16577 {
tmp16565 := PrimTail(V5078)

tmp16566 := PrimTail(tmp16565)

tmp16567 := PrimHead(tmp16566)

tmp16568 := PrimTail(tmp16567)

tmp16569 := PrimTail(tmp16568)

tmp16570 := PrimTail(tmp16569)

tmp16571 := PrimEqual(Nil, tmp16570)

var ifres16559 Obj

if True == tmp16571 {
tmp16561 := PrimTail(V5078)

tmp16562 := PrimTail(tmp16561)

tmp16563 := PrimTail(tmp16562)

tmp16564 := PrimEqual(Nil, tmp16563)

var ifres16560 Obj

if True == tmp16564 {
ifres16560 = True


} else {
ifres16560 = False


}

ifres16559 = ifres16560


} else {
ifres16559 = False


}

var ifres16558 Obj

if True == ifres16559 {
ifres16558 = True


} else {
ifres16558 = False


}

ifres16557 = ifres16558


} else {
ifres16557 = False


}

var ifres16556 Obj

if True == ifres16557 {
ifres16556 = True


} else {
ifres16556 = False


}

ifres16555 = ifres16556


} else {
ifres16555 = False


}

var ifres16554 Obj

if True == ifres16555 {
ifres16554 = True


} else {
ifres16554 = False


}

ifres16553 = ifres16554


} else {
ifres16553 = False


}

var ifres16552 Obj

if True == ifres16553 {
ifres16552 = True


} else {
ifres16552 = False


}

ifres16551 = ifres16552


} else {
ifres16551 = False


}

var ifres16550 Obj

if True == ifres16551 {
ifres16550 = True


} else {
ifres16550 = False


}

ifres16549 = ifres16550


} else {
ifres16549 = False


}

var ifres16548 Obj

if True == ifres16549 {
ifres16548 = True


} else {
ifres16548 = False


}

ifres16547 = ifres16548


} else {
ifres16547 = False


}

var ifres16546 Obj

if True == ifres16547 {
ifres16546 = True


} else {
ifres16546 = False


}

ifres16545 = ifres16546


} else {
ifres16545 = False


}

var ifres16544 Obj

if True == ifres16545 {
ifres16544 = True


} else {
ifres16544 = False


}

ifres16543 = ifres16544


} else {
ifres16543 = False


}

if True == ifres16543 {
tmp16452 := PrimTail(V5078)

tmp16453 := PrimHead(tmp16452)

tmp16454 := PrimTail(V5078)

tmp16455 := PrimTail(tmp16454)

tmp16456 := PrimHead(tmp16455)

tmp16457 := PrimTail(tmp16456)

tmp16458 := PrimCons(tmp16457, Nil)

tmp16459 := PrimCons(symnot, tmp16458)

tmp16460 := PrimCons(tmp16459, Nil)

tmp16461 := PrimCons(tmp16453, tmp16460)

tmp16462 := PrimCons(symand, tmp16461)

tmp16463 := PrimTail(V5078)

tmp16464 := PrimTail(tmp16463)

tmp16465 := PrimHead(tmp16464)

tmp16466 := PrimTail(tmp16465)

tmp16467 := PrimTail(tmp16466)

tmp16468 := PrimCons(tmp16462, tmp16467)

__e.Return(PrimCons(symwhere, tmp16468))
return


} else {
tmp16541 := PrimIsPair(V5078)

var ifres16522 Obj

if True == tmp16541 {
tmp16539 := PrimHead(V5078)

tmp16540 := PrimEqual(symwhere, tmp16539)

var ifres16524 Obj

if True == tmp16540 {
tmp16537 := PrimTail(V5078)

tmp16538 := PrimIsPair(tmp16537)

var ifres16526 Obj

if True == tmp16538 {
tmp16534 := PrimTail(V5078)

tmp16535 := PrimTail(tmp16534)

tmp16536 := PrimIsPair(tmp16535)

var ifres16528 Obj

if True == tmp16536 {
tmp16530 := PrimTail(V5078)

tmp16531 := PrimTail(tmp16530)

tmp16532 := PrimTail(tmp16531)

tmp16533 := PrimEqual(Nil, tmp16532)

var ifres16529 Obj

if True == tmp16533 {
ifres16529 = True


} else {
ifres16529 = False


}

ifres16528 = ifres16529


} else {
ifres16528 = False


}

var ifres16527 Obj

if True == ifres16528 {
ifres16527 = True


} else {
ifres16527 = False


}

ifres16526 = ifres16527


} else {
ifres16526 = False


}

var ifres16525 Obj

if True == ifres16526 {
ifres16525 = True


} else {
ifres16525 = False


}

ifres16524 = ifres16525


} else {
ifres16524 = False


}

var ifres16523 Obj

if True == ifres16524 {
ifres16523 = True


} else {
ifres16523 = False


}

ifres16522 = ifres16523


} else {
ifres16522 = False


}

if True == ifres16522 {
tmp16469 := PrimTail(V5078)

tmp16470 := PrimHead(tmp16469)

tmp16471 := PrimTail(V5078)

tmp16472 := PrimTail(tmp16471)

tmp16473 := PrimHead(tmp16472)

tmp16474 := PrimCons(symfail, Nil)

tmp16475 := PrimCons(tmp16474, Nil)

tmp16476 := PrimCons(tmp16473, tmp16475)

tmp16477 := PrimCons(sym_a, tmp16476)

tmp16478 := PrimCons(tmp16477, Nil)

tmp16479 := PrimCons(symnot, tmp16478)

tmp16480 := PrimCons(tmp16479, Nil)

tmp16481 := PrimCons(tmp16470, tmp16480)

tmp16482 := PrimCons(symand, tmp16481)

tmp16483 := PrimTail(V5078)

tmp16484 := PrimTail(tmp16483)

tmp16485 := PrimCons(tmp16482, tmp16484)

__e.Return(PrimCons(symwhere, tmp16485))
return


} else {
tmp16520 := PrimIsPair(V5078)

var ifres16501 Obj

if True == tmp16520 {
tmp16518 := PrimHead(V5078)

tmp16519 := PrimEqual(symfail_1if, tmp16518)

var ifres16503 Obj

if True == tmp16519 {
tmp16516 := PrimTail(V5078)

tmp16517 := PrimIsPair(tmp16516)

var ifres16505 Obj

if True == tmp16517 {
tmp16513 := PrimTail(V5078)

tmp16514 := PrimTail(tmp16513)

tmp16515 := PrimIsPair(tmp16514)

var ifres16507 Obj

if True == tmp16515 {
tmp16509 := PrimTail(V5078)

tmp16510 := PrimTail(tmp16509)

tmp16511 := PrimTail(tmp16510)

tmp16512 := PrimEqual(Nil, tmp16511)

var ifres16508 Obj

if True == tmp16512 {
ifres16508 = True


} else {
ifres16508 = False


}

ifres16507 = ifres16508


} else {
ifres16507 = False


}

var ifres16506 Obj

if True == ifres16507 {
ifres16506 = True


} else {
ifres16506 = False


}

ifres16505 = ifres16506


} else {
ifres16505 = False


}

var ifres16504 Obj

if True == ifres16505 {
ifres16504 = True


} else {
ifres16504 = False


}

ifres16503 = ifres16504


} else {
ifres16503 = False


}

var ifres16502 Obj

if True == ifres16503 {
ifres16502 = True


} else {
ifres16502 = False


}

ifres16501 = ifres16502


} else {
ifres16501 = False


}

if True == ifres16501 {
tmp16486 := PrimTail(V5078)

tmp16487 := PrimCons(tmp16486, Nil)

tmp16488 := PrimCons(symnot, tmp16487)

tmp16489 := PrimTail(V5078)

tmp16490 := PrimTail(tmp16489)

tmp16491 := PrimCons(tmp16488, tmp16490)

__e.Return(PrimCons(symwhere, tmp16491))
return


} else {
tmp16492 := PrimCons(symfail, Nil)

tmp16493 := PrimCons(tmp16492, Nil)

tmp16494 := PrimCons(V5078, tmp16493)

tmp16495 := PrimCons(sym_a, tmp16494)

tmp16496 := PrimCons(tmp16495, Nil)

tmp16497 := PrimCons(symnot, tmp16496)

tmp16498 := PrimCons(V5078, Nil)

tmp16499 := PrimCons(tmp16497, tmp16498)

__e.Return(PrimCons(symwhere, tmp16499))
return


}


}


}


}, 1)

tmp16600 := Call(__e, ns2_1set, symshen_4correct, tmp16451)


_ = tmp16600

tmp16601 := MakeNative(func(__e *ControlFlow) {
V5079 := __e.Get(1)
_ = V5079
V5080 := __e.Get(2)
_ = V5080
V5081 := __e.Get(3)
_ = V5081
V5082 := __e.Get(4)
_ = V5082
V5083 := __e.Get(5)
_ = V5083
V5084 := __e.Get(6)
_ = V5084
V5085 := __e.Get(7)
_ = V5085
V5086 := __e.Get(8)
_ = V5086
tmp16602 := MakeNative(func(__e *ControlFlow) {
W5087 := __e.Get(1)
_ = W5087
tmp16603 := MakeNative(func(__e *ControlFlow) {
W5088 := __e.Get(1)
_ = W5088
tmp16633 := PrimEqual(W5088, False)

if True == tmp16633 {
tmp16604 := MakeNative(func(__e *ControlFlow) {
W5090 := __e.Get(1)
_ = W5090
tmp16606 := PrimEqual(W5090, False)

if True == tmp16606 {
__e.TailApply(PrimFunc(symshen_4unlock), V5084, W5087)
return
} else {
__e.Return(W5090)
return
}


}, 1)

tmp16631 := Call(__e, PrimFunc(symshen_4unlocked_2), V5084)


var ifres16607 Obj

if True == tmp16631 {
tmp16608 := MakeNative(func(__e *ControlFlow) {
W5091 := __e.Get(1)
_ = W5091
tmp16628 := PrimIsPair(W5091)

if True == tmp16628 {
tmp16609 := MakeNative(func(__e *ControlFlow) {
W5092 := __e.Get(1)
_ = W5092
tmp16610 := MakeNative(func(__e *ControlFlow) {
W5093 := __e.Get(1)
_ = W5093
tmp16611 := MakeNative(func(__e *ControlFlow) {
W5094 := __e.Get(1)
_ = W5094
tmp16612 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16612

tmp16613 := Call(__e, PrimFunc(symshen_4deref), W5092, V5083)


tmp16614 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16613)


tmp16615 := MakeNative(func(__e *ControlFlow) {
tmp16616 := Call(__e, PrimFunc(symshen_4lazyderef), W5094, V5083)


tmp16617 := Call(__e, PrimFunc(symfst), tmp16616)


tmp16618 := Call(__e, PrimFunc(symshen_4lazyderef), W5094, V5083)


tmp16619 := Call(__e, PrimFunc(symsnd), tmp16618)


tmp16620 := MakeNative(func(__e *ControlFlow) {
tmp16621 := MakeNative(func(__e *ControlFlow) {
tmp16622 := PrimNumberAdd(V5082, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5079, W5093, V5081, tmp16622, V5083, V5084, W5087, V5086)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5083, V5084, W5087, tmp16621)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5079, V5082, tmp16617, tmp16619, V5081, V5083, V5084, W5087, tmp16620)
return


}, 0)

tmp16623 := Call(__e, PrimFunc(symbind), W5094, tmp16614, V5083, V5084, W5087, tmp16615)


__e.TailApply(PrimFunc(symshen_4gc), V5083, tmp16623)
return


}, 1)

tmp16624 := Call(__e, PrimFunc(symshen_4newpv), V5083)


__e.TailApply(tmp16611, tmp16624)
return


}, 1)

tmp16625 := PrimTail(W5091)

__e.TailApply(tmp16610, tmp16625)
return


}, 1)

tmp16626 := PrimHead(W5091)

__e.TailApply(tmp16609, tmp16626)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16629 := Call(__e, PrimFunc(symshen_4lazyderef), V5080, V5083)


tmp16630 := Call(__e, tmp16608, tmp16629)


ifres16607 = tmp16630


} else {
ifres16607 = False


}

__e.TailApply(tmp16604, ifres16607)
return


} else {
__e.Return(W5088)
return
}


}, 1)

tmp16641 := Call(__e, PrimFunc(symshen_4unlocked_2), V5084)


var ifres16634 Obj

if True == tmp16641 {
tmp16635 := MakeNative(func(__e *ControlFlow) {
W5089 := __e.Get(1)
_ = W5089
tmp16638 := PrimEqual(W5089, Nil)

if True == tmp16638 {
tmp16636 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16636

__e.TailApply(PrimFunc(symthaw), V5086)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16639 := Call(__e, PrimFunc(symshen_4lazyderef), V5080, V5083)


tmp16640 := Call(__e, tmp16635, tmp16639)


ifres16634 = tmp16640


} else {
ifres16634 = False


}

__e.TailApply(tmp16603, ifres16634)
return


}, 1)

tmp16642 := PrimNumberAdd(V5085, MakeNumber(1))

__e.TailApply(tmp16602, tmp16642)
return


}, 8)

tmp16643 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16601)


_ = tmp16643

tmp16644 := MakeNative(func(__e *ControlFlow) {
V5095 := __e.Get(1)
_ = V5095
tmp16657 := Call(__e, PrimFunc(symtuple_2), V5095)


if True == tmp16657 {
tmp16645 := MakeNative(func(__e *ControlFlow) {
W5096 := __e.Get(1)
_ = W5096
tmp16646 := MakeNative(func(__e *ControlFlow) {
W5097 := __e.Get(1)
_ = W5097
tmp16647 := Call(__e, PrimFunc(symfst), V5095)


tmp16648 := Call(__e, PrimFunc(symshen_4freshen), W5097, tmp16647)


tmp16649 := Call(__e, PrimFunc(symsnd), V5095)


tmp16650 := Call(__e, PrimFunc(symshen_4freshen), W5097, tmp16649)


__e.TailApply(PrimFunc(sym_8p), tmp16648, tmp16650)
return


}, 1)

tmp16651 := MakeNative(func(__e *ControlFlow) {
Z5098 := __e.Get(1)
_ = Z5098
tmp16652 := Call(__e, PrimFunc(symshen_4freshterm), Z5098)


__e.Return(PrimCons(Z5098, tmp16652))
return


}, 1)

tmp16653 := Call(__e, PrimFunc(symmap), tmp16651, W5096)


__e.TailApply(tmp16646, tmp16653)
return


}, 1)

tmp16654 := Call(__e, PrimFunc(symfst), V5095)


tmp16655 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16654)


__e.TailApply(tmp16645, tmp16655)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen-rule")))
return
}


}, 1)

tmp16658 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16644)


_ = tmp16658

tmp16659 := MakeNative(func(__e *ControlFlow) {
V5099 := __e.Get(1)
_ = V5099
V5100 := __e.Get(2)
_ = V5100
tmp16673 := PrimEqual(Nil, V5099)

if True == tmp16673 {
__e.Return(V5100)
return
} else {
tmp16671 := PrimIsPair(V5099)

var ifres16667 Obj

if True == tmp16671 {
tmp16669 := PrimHead(V5099)

tmp16670 := PrimIsPair(tmp16669)

var ifres16668 Obj

if True == tmp16670 {
ifres16668 = True


} else {
ifres16668 = False


}

ifres16667 = ifres16668


} else {
ifres16667 = False


}

if True == ifres16667 {
tmp16660 := PrimTail(V5099)

tmp16661 := PrimHead(V5099)

tmp16662 := PrimHead(tmp16661)

tmp16663 := PrimHead(V5099)

tmp16664 := PrimTail(tmp16663)

tmp16665 := Call(__e, PrimFunc(symshen_4beta), tmp16662, tmp16664, V5100)


__e.TailApply(PrimFunc(symshen_4freshen), tmp16660, tmp16665)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen")))
return
}


}


}, 2)

tmp16674 := Call(__e, ns2_1set, symshen_4freshen, tmp16659)


_ = tmp16674

tmp16675 := MakeNative(func(__e *ControlFlow) {
V5101 := __e.Get(1)
_ = V5101
V5102 := __e.Get(2)
_ = V5102
V5103 := __e.Get(3)
_ = V5103
V5104 := __e.Get(4)
_ = V5104
V5105 := __e.Get(5)
_ = V5105
V5106 := __e.Get(6)
_ = V5106
V5107 := __e.Get(7)
_ = V5107
V5108 := __e.Get(8)
_ = V5108
V5109 := __e.Get(9)
_ = V5109
tmp16676 := MakeNative(func(__e *ControlFlow) {
W5110 := __e.Get(1)
_ = W5110
tmp16689 := PrimEqual(W5110, False)

if True == tmp16689 {
tmp16687 := Call(__e, PrimFunc(symshen_4unlocked_2), V5107)


if True == tmp16687 {
tmp16677 := MakeNative(func(__e *ControlFlow) {
W5111 := __e.Get(1)
_ = W5111
tmp16678 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16678

tmp16679 := Call(__e, PrimFunc(symshen_4app), V5101, MakeString("\n"), symshen_4a)


tmp16680 := PrimStringConcat(MakeString(" of "), tmp16679)

tmp16681 := Call(__e, PrimFunc(symshen_4app), V5102, tmp16680, symshen_4a)


tmp16682 := PrimStringConcat(MakeString("type error in rule "), tmp16681)

tmp16683 := PrimSimpleError(tmp16682)

tmp16684 := Call(__e, PrimFunc(symbind), W5111, tmp16683, V5106, V5107, V5108, V5109)


__e.TailApply(PrimFunc(symshen_4gc), V5106, tmp16684)
return


}, 1)

tmp16685 := Call(__e, PrimFunc(symshen_4newpv), V5106)


__e.TailApply(tmp16677, tmp16685)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5110)
return
}


}, 1)

tmp16693 := Call(__e, PrimFunc(symshen_4unlocked_2), V5107)


var ifres16690 Obj

if True == tmp16693 {
tmp16691 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16691

tmp16692 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5103, V5104, V5105, V5106, V5107, V5108, V5109)


ifres16690 = tmp16692


} else {
ifres16690 = False


}

__e.TailApply(tmp16676, ifres16690)
return


}, 9)

tmp16694 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16675)


_ = tmp16694

tmp16695 := MakeNative(func(__e *ControlFlow) {
V5112 := __e.Get(1)
_ = V5112
V5113 := __e.Get(2)
_ = V5113
V5114 := __e.Get(3)
_ = V5114
V5115 := __e.Get(4)
_ = V5115
V5116 := __e.Get(5)
_ = V5116
V5117 := __e.Get(6)
_ = V5117
V5118 := __e.Get(7)
_ = V5118
tmp16696 := MakeNative(func(__e *ControlFlow) {
W5119 := __e.Get(1)
_ = W5119
tmp16697 := MakeNative(func(__e *ControlFlow) {
W5120 := __e.Get(1)
_ = W5120
tmp16720 := PrimEqual(W5120, False)

if True == tmp16720 {
tmp16698 := MakeNative(func(__e *ControlFlow) {
W5127 := __e.Get(1)
_ = W5127
tmp16700 := PrimEqual(W5127, False)

if True == tmp16700 {
__e.TailApply(PrimFunc(symshen_4unlock), V5116, W5119)
return
} else {
__e.Return(W5127)
return
}


}, 1)

tmp16718 := Call(__e, PrimFunc(symshen_4unlocked_2), V5116)


var ifres16701 Obj

if True == tmp16718 {
tmp16702 := MakeNative(func(__e *ControlFlow) {
W5128 := __e.Get(1)
_ = W5128
tmp16703 := MakeNative(func(__e *ControlFlow) {
W5129 := __e.Get(1)
_ = W5129
tmp16704 := MakeNative(func(__e *ControlFlow) {
W5130 := __e.Get(1)
_ = W5130
tmp16705 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16705

tmp16706 := Call(__e, PrimFunc(symshen_4freshterms), V5112)


tmp16707 := MakeNative(func(__e *ControlFlow) {
tmp16708 := MakeNative(func(__e *ControlFlow) {
tmp16709 := MakeNative(func(__e *ControlFlow) {
tmp16710 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5113, W5129, W5130, V5115, V5116, W5119, V5118)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5112, V5114, W5130, V5115, V5116, W5119, tmp16710)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5115, V5116, W5119, tmp16709)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5112, V5114, W5128, W5129, V5115, V5116, W5119, tmp16708)
return


}, 0)

tmp16711 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16706, W5128, V5115, V5116, W5119, tmp16707)


__e.TailApply(PrimFunc(symshen_4gc), V5115, tmp16711)
return


}, 1)

tmp16712 := Call(__e, PrimFunc(symshen_4newpv), V5115)


tmp16713 := Call(__e, tmp16704, tmp16712)


__e.TailApply(PrimFunc(symshen_4gc), V5115, tmp16713)
return


}, 1)

tmp16714 := Call(__e, PrimFunc(symshen_4newpv), V5115)


tmp16715 := Call(__e, tmp16703, tmp16714)


__e.TailApply(PrimFunc(symshen_4gc), V5115, tmp16715)
return


}, 1)

tmp16716 := Call(__e, PrimFunc(symshen_4newpv), V5115)


tmp16717 := Call(__e, tmp16702, tmp16716)


ifres16701 = tmp16717


} else {
ifres16701 = False


}

__e.TailApply(tmp16698, ifres16701)
return


} else {
__e.Return(W5120)
return
}


}, 1)

tmp16750 := Call(__e, PrimFunc(symshen_4unlocked_2), V5116)


var ifres16721 Obj

if True == tmp16750 {
tmp16722 := MakeNative(func(__e *ControlFlow) {
W5121 := __e.Get(1)
_ = W5121
tmp16747 := PrimEqual(W5121, Nil)

if True == tmp16747 {
tmp16723 := MakeNative(func(__e *ControlFlow) {
W5122 := __e.Get(1)
_ = W5122
tmp16744 := PrimIsPair(W5122)

if True == tmp16744 {
tmp16724 := MakeNative(func(__e *ControlFlow) {
W5123 := __e.Get(1)
_ = W5123
tmp16740 := PrimEqual(W5123, sym_1_1_6)

if True == tmp16740 {
tmp16725 := MakeNative(func(__e *ControlFlow) {
W5124 := __e.Get(1)
_ = W5124
tmp16736 := PrimIsPair(W5124)

if True == tmp16736 {
tmp16726 := MakeNative(func(__e *ControlFlow) {
W5125 := __e.Get(1)
_ = W5125
tmp16727 := MakeNative(func(__e *ControlFlow) {
W5126 := __e.Get(1)
_ = W5126
tmp16731 := PrimEqual(W5126, Nil)

if True == tmp16731 {
tmp16728 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16728

tmp16729 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5113, W5125, Nil, V5115, V5116, W5119, V5118)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5115, V5116, W5119, tmp16729)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16732 := PrimTail(W5124)

tmp16733 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16732, V5115)


__e.TailApply(tmp16727, tmp16733)
return


}, 1)

tmp16734 := PrimHead(W5124)

__e.TailApply(tmp16726, tmp16734)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16737 := PrimTail(W5122)

tmp16738 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16737, V5115)


__e.TailApply(tmp16725, tmp16738)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16741 := PrimHead(W5122)

tmp16742 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16741, V5115)


__e.TailApply(tmp16724, tmp16742)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16745 := Call(__e, PrimFunc(symshen_4lazyderef), V5114, V5115)


__e.TailApply(tmp16723, tmp16745)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16748 := Call(__e, PrimFunc(symshen_4lazyderef), V5112, V5115)


tmp16749 := Call(__e, tmp16722, tmp16748)


ifres16721 = tmp16749


} else {
ifres16721 = False


}

__e.TailApply(tmp16697, ifres16721)
return


}, 1)

tmp16751 := PrimNumberAdd(V5117, MakeNumber(1))

__e.TailApply(tmp16696, tmp16751)
return


}, 7)

tmp16752 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16695)


_ = tmp16752

tmp16753 := MakeNative(func(__e *ControlFlow) {
V5131 := __e.Get(1)
_ = V5131
V5132 := __e.Get(2)
_ = V5132
V5133 := __e.Get(3)
_ = V5133
V5134 := __e.Get(4)
_ = V5134
V5135 := __e.Get(5)
_ = V5135
V5136 := __e.Get(6)
_ = V5136
V5137 := __e.Get(7)
_ = V5137
tmp16754 := MakeNative(func(__e *ControlFlow) {
W5138 := __e.Get(1)
_ = W5138
tmp16907 := PrimEqual(W5138, False)

if True == tmp16907 {
tmp16905 := Call(__e, PrimFunc(symshen_4unlocked_2), V5135)


if True == tmp16905 {
tmp16755 := MakeNative(func(__e *ControlFlow) {
W5142 := __e.Get(1)
_ = W5142
tmp16902 := PrimIsPair(W5142)

if True == tmp16902 {
tmp16756 := MakeNative(func(__e *ControlFlow) {
W5143 := __e.Get(1)
_ = W5143
tmp16757 := MakeNative(func(__e *ControlFlow) {
W5144 := __e.Get(1)
_ = W5144
tmp16758 := MakeNative(func(__e *ControlFlow) {
W5145 := __e.Get(1)
_ = W5145
tmp16897 := PrimIsPair(W5145)

if True == tmp16897 {
tmp16759 := MakeNative(func(__e *ControlFlow) {
W5146 := __e.Get(1)
_ = W5146
tmp16760 := MakeNative(func(__e *ControlFlow) {
W5147 := __e.Get(1)
_ = W5147
tmp16892 := PrimIsPair(W5147)

if True == tmp16892 {
tmp16761 := MakeNative(func(__e *ControlFlow) {
W5148 := __e.Get(1)
_ = W5148
tmp16888 := PrimEqual(W5148, sym_1_1_6)

if True == tmp16888 {
tmp16762 := MakeNative(func(__e *ControlFlow) {
W5149 := __e.Get(1)
_ = W5149
tmp16884 := PrimIsPair(W5149)

if True == tmp16884 {
tmp16763 := MakeNative(func(__e *ControlFlow) {
W5150 := __e.Get(1)
_ = W5150
tmp16764 := MakeNative(func(__e *ControlFlow) {
W5151 := __e.Get(1)
_ = W5151
tmp16879 := PrimEqual(W5151, Nil)

if True == tmp16879 {
tmp16765 := MakeNative(func(__e *ControlFlow) {
W5152 := __e.Get(1)
_ = W5152
tmp16766 := MakeNative(func(__e *ControlFlow) {
W5153 := __e.Get(1)
_ = W5153
tmp16870 := PrimIsPair(W5152)

if True == tmp16870 {
tmp16767 := MakeNative(func(__e *ControlFlow) {
W5158 := __e.Get(1)
_ = W5158
tmp16768 := MakeNative(func(__e *ControlFlow) {
W5159 := __e.Get(1)
_ = W5159
tmp16838 := PrimIsPair(W5158)

if True == tmp16838 {
tmp16769 := MakeNative(func(__e *ControlFlow) {
W5164 := __e.Get(1)
_ = W5164
tmp16770 := MakeNative(func(__e *ControlFlow) {
W5165 := __e.Get(1)
_ = W5165
tmp16771 := MakeNative(func(__e *ControlFlow) {
W5166 := __e.Get(1)
_ = W5166
tmp16813 := PrimIsPair(W5165)

if True == tmp16813 {
tmp16772 := MakeNative(func(__e *ControlFlow) {
W5169 := __e.Get(1)
_ = W5169
tmp16773 := MakeNative(func(__e *ControlFlow) {
W5170 := __e.Get(1)
_ = W5170
tmp16774 := MakeNative(func(__e *ControlFlow) {
W5171 := __e.Get(1)
_ = W5171
tmp16794 := PrimIsPair(W5170)

if True == tmp16794 {
tmp16775 := MakeNative(func(__e *ControlFlow) {
W5173 := __e.Get(1)
_ = W5173
tmp16776 := MakeNative(func(__e *ControlFlow) {
W5174 := __e.Get(1)
_ = W5174
tmp16777 := MakeNative(func(__e *ControlFlow) {
W5175 := __e.Get(1)
_ = W5175
tmp16781 := PrimEqual(W5174, Nil)

if True == tmp16781 {
__e.TailApply(PrimFunc(symthaw), W5175)
return
} else {
tmp16779 := Call(__e, PrimFunc(symshen_4pvar_2), W5174)


if True == tmp16779 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5174, Nil, V5134, W5175)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16782 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5171, W5173)
return
}, 0)

__e.TailApply(tmp16777, tmp16782)
return


}, 1)

tmp16783 := PrimTail(W5170)

tmp16784 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16783, V5134)


__e.TailApply(tmp16776, tmp16784)
return


}, 1)

tmp16785 := PrimHead(W5170)

__e.TailApply(tmp16775, tmp16785)
return


} else {
tmp16792 := Call(__e, PrimFunc(symshen_4pvar_2), W5170)


if True == tmp16792 {
tmp16786 := MakeNative(func(__e *ControlFlow) {
W5176 := __e.Get(1)
_ = W5176
tmp16787 := PrimCons(W5176, Nil)

tmp16788 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5171, W5176)
return
}, 0)

tmp16789 := Call(__e, PrimFunc(symshen_4bind_b), W5170, tmp16787, V5134, tmp16788)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16789)
return


}, 1)

tmp16790 := Call(__e, PrimFunc(symshen_4newpv), V5134)


__e.TailApply(tmp16786, tmp16790)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16795 := MakeNative(func(__e *ControlFlow) {
Z5172 := __e.Get(1)
_ = Z5172
tmp16796 := Call(__e, W5166, W5169)


__e.TailApply(tmp16796, Z5172)
return


}, 1)

__e.TailApply(tmp16774, tmp16795)
return


}, 1)

tmp16797 := PrimTail(W5165)

tmp16798 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16797, V5134)


__e.TailApply(tmp16773, tmp16798)
return


}, 1)

tmp16799 := PrimHead(W5165)

__e.TailApply(tmp16772, tmp16799)
return


} else {
tmp16811 := Call(__e, PrimFunc(symshen_4pvar_2), W5165)


if True == tmp16811 {
tmp16800 := MakeNative(func(__e *ControlFlow) {
W5177 := __e.Get(1)
_ = W5177
tmp16801 := MakeNative(func(__e *ControlFlow) {
W5178 := __e.Get(1)
_ = W5178
tmp16802 := PrimCons(W5178, Nil)

tmp16803 := PrimCons(W5177, tmp16802)

tmp16804 := MakeNative(func(__e *ControlFlow) {
tmp16805 := Call(__e, W5166, W5177)


__e.TailApply(tmp16805, W5178)
return


}, 0)

tmp16806 := Call(__e, PrimFunc(symshen_4bind_b), W5165, tmp16803, V5134, tmp16804)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16806)
return


}, 1)

tmp16807 := Call(__e, PrimFunc(symshen_4newpv), V5134)


tmp16808 := Call(__e, tmp16801, tmp16807)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16808)
return


}, 1)

tmp16809 := Call(__e, PrimFunc(symshen_4newpv), V5134)


__e.TailApply(tmp16800, tmp16809)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16814 := MakeNative(func(__e *ControlFlow) {
Z5167 := __e.Get(1)
_ = Z5167
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5168 := __e.Get(1)
_ = Z5168
tmp16815 := Call(__e, W5159, W5164)


tmp16816 := Call(__e, tmp16815, Z5167)


__e.TailApply(tmp16816, Z5168)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16771, tmp16814)
return


}, 1)

tmp16817 := PrimTail(W5158)

tmp16818 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16817, V5134)


__e.TailApply(tmp16770, tmp16818)
return


}, 1)

tmp16819 := PrimHead(W5158)

__e.TailApply(tmp16769, tmp16819)
return


} else {
tmp16836 := Call(__e, PrimFunc(symshen_4pvar_2), W5158)


if True == tmp16836 {
tmp16820 := MakeNative(func(__e *ControlFlow) {
W5179 := __e.Get(1)
_ = W5179
tmp16821 := MakeNative(func(__e *ControlFlow) {
W5180 := __e.Get(1)
_ = W5180
tmp16822 := MakeNative(func(__e *ControlFlow) {
W5181 := __e.Get(1)
_ = W5181
tmp16823 := PrimCons(W5181, Nil)

tmp16824 := PrimCons(W5180, tmp16823)

tmp16825 := PrimCons(W5179, tmp16824)

tmp16826 := MakeNative(func(__e *ControlFlow) {
tmp16827 := Call(__e, W5159, W5179)


tmp16828 := Call(__e, tmp16827, W5180)


__e.TailApply(tmp16828, W5181)
return


}, 0)

tmp16829 := Call(__e, PrimFunc(symshen_4bind_b), W5158, tmp16825, V5134, tmp16826)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16829)
return


}, 1)

tmp16830 := Call(__e, PrimFunc(symshen_4newpv), V5134)


tmp16831 := Call(__e, tmp16822, tmp16830)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16831)
return


}, 1)

tmp16832 := Call(__e, PrimFunc(symshen_4newpv), V5134)


tmp16833 := Call(__e, tmp16821, tmp16832)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16833)
return


}, 1)

tmp16834 := Call(__e, PrimFunc(symshen_4newpv), V5134)


__e.TailApply(tmp16820, tmp16834)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16839 := MakeNative(func(__e *ControlFlow) {
Z5160 := __e.Get(1)
_ = Z5160
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5161 := __e.Get(1)
_ = Z5161
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5162 := __e.Get(1)
_ = Z5162
tmp16840 := MakeNative(func(__e *ControlFlow) {
W5163 := __e.Get(1)
_ = W5163
tmp16841 := Call(__e, W5153, Z5160)


tmp16842 := Call(__e, tmp16841, Z5161)


tmp16843 := Call(__e, tmp16842, Z5162)


__e.TailApply(tmp16843, W5163)
return


}, 1)

tmp16844 := PrimTail(W5152)

__e.TailApply(tmp16840, tmp16844)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16768, tmp16839)
return


}, 1)

tmp16845 := PrimHead(W5152)

tmp16846 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16845, V5134)


__e.TailApply(tmp16767, tmp16846)
return


} else {
tmp16868 := Call(__e, PrimFunc(symshen_4pvar_2), W5152)


if True == tmp16868 {
tmp16847 := MakeNative(func(__e *ControlFlow) {
W5182 := __e.Get(1)
_ = W5182
tmp16848 := MakeNative(func(__e *ControlFlow) {
W5183 := __e.Get(1)
_ = W5183
tmp16849 := MakeNative(func(__e *ControlFlow) {
W5184 := __e.Get(1)
_ = W5184
tmp16850 := MakeNative(func(__e *ControlFlow) {
W5185 := __e.Get(1)
_ = W5185
tmp16851 := PrimCons(W5184, Nil)

tmp16852 := PrimCons(W5183, tmp16851)

tmp16853 := PrimCons(W5182, tmp16852)

tmp16854 := PrimCons(tmp16853, W5185)

tmp16855 := MakeNative(func(__e *ControlFlow) {
tmp16856 := Call(__e, W5153, W5182)


tmp16857 := Call(__e, tmp16856, W5183)


tmp16858 := Call(__e, tmp16857, W5184)


__e.TailApply(tmp16858, W5185)
return


}, 0)

tmp16859 := Call(__e, PrimFunc(symshen_4bind_b), W5152, tmp16854, V5134, tmp16855)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16859)
return


}, 1)

tmp16860 := Call(__e, PrimFunc(symshen_4newpv), V5134)


tmp16861 := Call(__e, tmp16850, tmp16860)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16861)
return


}, 1)

tmp16862 := Call(__e, PrimFunc(symshen_4newpv), V5134)


tmp16863 := Call(__e, tmp16849, tmp16862)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16863)
return


}, 1)

tmp16864 := Call(__e, PrimFunc(symshen_4newpv), V5134)


tmp16865 := Call(__e, tmp16848, tmp16864)


__e.TailApply(PrimFunc(symshen_4gc), V5134, tmp16865)
return


}, 1)

tmp16866 := Call(__e, PrimFunc(symshen_4newpv), V5134)


__e.TailApply(tmp16847, tmp16866)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16871 := MakeNative(func(__e *ControlFlow) {
Z5154 := __e.Get(1)
_ = Z5154
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5155 := __e.Get(1)
_ = Z5155
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5156 := __e.Get(1)
_ = Z5156
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5157 := __e.Get(1)
_ = Z5157
tmp16872 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16872

tmp16873 := MakeNative(func(__e *ControlFlow) {
tmp16874 := MakeNative(func(__e *ControlFlow) {
tmp16875 := PrimIntern(MakeString(":"))

tmp16876 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), W5144, W5150, Z5157, V5134, V5135, V5136, V5137)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5155, tmp16875, V5134, V5135, V5136, tmp16876)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5143, Z5154, V5134, V5135, V5136, tmp16874)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5146, Z5156, V5134, V5135, V5136, tmp16873)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16766, tmp16871)
return


}, 1)

tmp16877 := Call(__e, PrimFunc(symshen_4lazyderef), V5133, V5134)


__e.TailApply(tmp16765, tmp16877)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16880 := PrimTail(W5149)

tmp16881 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16880, V5134)


__e.TailApply(tmp16764, tmp16881)
return


}, 1)

tmp16882 := PrimHead(W5149)

__e.TailApply(tmp16763, tmp16882)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16885 := PrimTail(W5147)

tmp16886 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16885, V5134)


__e.TailApply(tmp16762, tmp16886)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16889 := PrimHead(W5147)

tmp16890 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16889, V5134)


__e.TailApply(tmp16761, tmp16890)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16893 := PrimTail(W5145)

tmp16894 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16893, V5134)


__e.TailApply(tmp16760, tmp16894)
return


}, 1)

tmp16895 := PrimHead(W5145)

__e.TailApply(tmp16759, tmp16895)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16898 := Call(__e, PrimFunc(symshen_4lazyderef), V5132, V5134)


__e.TailApply(tmp16758, tmp16898)
return


}, 1)

tmp16899 := PrimTail(W5142)

__e.TailApply(tmp16757, tmp16899)
return


}, 1)

tmp16900 := PrimHead(W5142)

__e.TailApply(tmp16756, tmp16900)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16903 := Call(__e, PrimFunc(symshen_4lazyderef), V5131, V5134)


__e.TailApply(tmp16755, tmp16903)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5138)
return
}


}, 1)

tmp16923 := Call(__e, PrimFunc(symshen_4unlocked_2), V5135)


var ifres16908 Obj

if True == tmp16923 {
tmp16909 := MakeNative(func(__e *ControlFlow) {
W5139 := __e.Get(1)
_ = W5139
tmp16920 := PrimEqual(W5139, Nil)

if True == tmp16920 {
tmp16910 := MakeNative(func(__e *ControlFlow) {
W5140 := __e.Get(1)
_ = W5140
tmp16911 := MakeNative(func(__e *ControlFlow) {
W5141 := __e.Get(1)
_ = W5141
tmp16915 := PrimEqual(W5140, Nil)

if True == tmp16915 {
__e.TailApply(PrimFunc(symthaw), W5141)
return
} else {
tmp16913 := Call(__e, PrimFunc(symshen_4pvar_2), W5140)


if True == tmp16913 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5140, Nil, V5134, W5141)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16916 := MakeNative(func(__e *ControlFlow) {
tmp16917 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16917

__e.TailApply(PrimFunc(symthaw), V5137)
return


}, 0)

__e.TailApply(tmp16911, tmp16916)
return


}, 1)

tmp16918 := Call(__e, PrimFunc(symshen_4lazyderef), V5133, V5134)


__e.TailApply(tmp16910, tmp16918)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16921 := Call(__e, PrimFunc(symshen_4lazyderef), V5131, V5134)


tmp16922 := Call(__e, tmp16909, tmp16921)


ifres16908 = tmp16922


} else {
ifres16908 = False


}

__e.TailApply(tmp16754, ifres16908)
return


}, 7)

tmp16924 := Call(__e, ns2_1set, symshen_4myassume, tmp16753)


_ = tmp16924

tmp16925 := MakeNative(func(__e *ControlFlow) {
V5188 := __e.Get(1)
_ = V5188
tmp16948 := PrimEqual(Nil, V5188)

if True == tmp16948 {
__e.Return(Nil)
return
} else {
tmp16946 := PrimIsPair(V5188)

var ifres16942 Obj

if True == tmp16946 {
tmp16944 := PrimHead(V5188)

tmp16945 := PrimIsPair(tmp16944)

var ifres16943 Obj

if True == tmp16945 {
ifres16943 = True


} else {
ifres16943 = False


}

ifres16942 = ifres16943


} else {
ifres16942 = False


}

if True == ifres16942 {
tmp16926 := PrimHead(V5188)

tmp16927 := PrimTail(V5188)

tmp16928 := Call(__e, PrimFunc(symappend), tmp16926, tmp16927)


__e.TailApply(PrimFunc(symshen_4freshterms), tmp16928)
return


} else {
tmp16940 := PrimIsPair(V5188)

var ifres16936 Obj

if True == tmp16940 {
tmp16938 := PrimHead(V5188)

tmp16939 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp16938)


var ifres16937 Obj

if True == tmp16939 {
ifres16937 = True


} else {
ifres16937 = False


}

ifres16936 = ifres16937


} else {
ifres16936 = False


}

if True == ifres16936 {
tmp16929 := PrimHead(V5188)

tmp16930 := PrimTail(V5188)

tmp16931 := Call(__e, PrimFunc(symshen_4freshterms), tmp16930)


__e.TailApply(PrimFunc(symadjoin), tmp16929, tmp16931)
return


} else {
tmp16934 := PrimIsPair(V5188)

if True == tmp16934 {
tmp16932 := PrimTail(V5188)

__e.TailApply(PrimFunc(symshen_4freshterms), tmp16932)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshterms")))
return
}


}


}


}


}, 1)

tmp16949 := Call(__e, ns2_1set, symshen_4freshterms, tmp16925)


_ = tmp16949

tmp16950 := MakeNative(func(__e *ControlFlow) {
V5189 := __e.Get(1)
_ = V5189
V5190 := __e.Get(2)
_ = V5190
V5191 := __e.Get(3)
_ = V5191
V5192 := __e.Get(4)
_ = V5192
V5193 := __e.Get(5)
_ = V5193
V5194 := __e.Get(6)
_ = V5194
tmp16951 := MakeNative(func(__e *ControlFlow) {
W5195 := __e.Get(1)
_ = W5195
tmp17075 := PrimEqual(W5195, False)

if True == tmp17075 {
tmp17073 := Call(__e, PrimFunc(symshen_4unlocked_2), V5192)


if True == tmp17073 {
tmp16952 := MakeNative(func(__e *ControlFlow) {
W5199 := __e.Get(1)
_ = W5199
tmp17070 := PrimIsPair(W5199)

if True == tmp17070 {
tmp16953 := MakeNative(func(__e *ControlFlow) {
W5200 := __e.Get(1)
_ = W5200
tmp16954 := MakeNative(func(__e *ControlFlow) {
W5201 := __e.Get(1)
_ = W5201
tmp16955 := MakeNative(func(__e *ControlFlow) {
W5202 := __e.Get(1)
_ = W5202
tmp16956 := MakeNative(func(__e *ControlFlow) {
W5203 := __e.Get(1)
_ = W5203
tmp17060 := PrimIsPair(W5202)

if True == tmp17060 {
tmp16957 := MakeNative(func(__e *ControlFlow) {
W5208 := __e.Get(1)
_ = W5208
tmp16958 := MakeNative(func(__e *ControlFlow) {
W5209 := __e.Get(1)
_ = W5209
tmp17028 := PrimIsPair(W5208)

if True == tmp17028 {
tmp16959 := MakeNative(func(__e *ControlFlow) {
W5214 := __e.Get(1)
_ = W5214
tmp16960 := MakeNative(func(__e *ControlFlow) {
W5215 := __e.Get(1)
_ = W5215
tmp16961 := MakeNative(func(__e *ControlFlow) {
W5216 := __e.Get(1)
_ = W5216
tmp17003 := PrimIsPair(W5215)

if True == tmp17003 {
tmp16962 := MakeNative(func(__e *ControlFlow) {
W5219 := __e.Get(1)
_ = W5219
tmp16963 := MakeNative(func(__e *ControlFlow) {
W5220 := __e.Get(1)
_ = W5220
tmp16964 := MakeNative(func(__e *ControlFlow) {
W5221 := __e.Get(1)
_ = W5221
tmp16984 := PrimIsPair(W5220)

if True == tmp16984 {
tmp16965 := MakeNative(func(__e *ControlFlow) {
W5223 := __e.Get(1)
_ = W5223
tmp16966 := MakeNative(func(__e *ControlFlow) {
W5224 := __e.Get(1)
_ = W5224
tmp16967 := MakeNative(func(__e *ControlFlow) {
W5225 := __e.Get(1)
_ = W5225
tmp16971 := PrimEqual(W5224, Nil)

if True == tmp16971 {
__e.TailApply(PrimFunc(symthaw), W5225)
return
} else {
tmp16969 := Call(__e, PrimFunc(symshen_4pvar_2), W5224)


if True == tmp16969 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5224, Nil, V5191, W5225)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16972 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5221, W5223)
return
}, 0)

__e.TailApply(tmp16967, tmp16972)
return


}, 1)

tmp16973 := PrimTail(W5220)

tmp16974 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16973, V5191)


__e.TailApply(tmp16966, tmp16974)
return


}, 1)

tmp16975 := PrimHead(W5220)

__e.TailApply(tmp16965, tmp16975)
return


} else {
tmp16982 := Call(__e, PrimFunc(symshen_4pvar_2), W5220)


if True == tmp16982 {
tmp16976 := MakeNative(func(__e *ControlFlow) {
W5226 := __e.Get(1)
_ = W5226
tmp16977 := PrimCons(W5226, Nil)

tmp16978 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5221, W5226)
return
}, 0)

tmp16979 := Call(__e, PrimFunc(symshen_4bind_b), W5220, tmp16977, V5191, tmp16978)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp16979)
return


}, 1)

tmp16980 := Call(__e, PrimFunc(symshen_4newpv), V5191)


__e.TailApply(tmp16976, tmp16980)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16985 := MakeNative(func(__e *ControlFlow) {
Z5222 := __e.Get(1)
_ = Z5222
tmp16986 := Call(__e, W5216, W5219)


__e.TailApply(tmp16986, Z5222)
return


}, 1)

__e.TailApply(tmp16964, tmp16985)
return


}, 1)

tmp16987 := PrimTail(W5215)

tmp16988 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16987, V5191)


__e.TailApply(tmp16963, tmp16988)
return


}, 1)

tmp16989 := PrimHead(W5215)

__e.TailApply(tmp16962, tmp16989)
return


} else {
tmp17001 := Call(__e, PrimFunc(symshen_4pvar_2), W5215)


if True == tmp17001 {
tmp16990 := MakeNative(func(__e *ControlFlow) {
W5227 := __e.Get(1)
_ = W5227
tmp16991 := MakeNative(func(__e *ControlFlow) {
W5228 := __e.Get(1)
_ = W5228
tmp16992 := PrimCons(W5228, Nil)

tmp16993 := PrimCons(W5227, tmp16992)

tmp16994 := MakeNative(func(__e *ControlFlow) {
tmp16995 := Call(__e, W5216, W5227)


__e.TailApply(tmp16995, W5228)
return


}, 0)

tmp16996 := Call(__e, PrimFunc(symshen_4bind_b), W5215, tmp16993, V5191, tmp16994)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp16996)
return


}, 1)

tmp16997 := Call(__e, PrimFunc(symshen_4newpv), V5191)


tmp16998 := Call(__e, tmp16991, tmp16997)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp16998)
return


}, 1)

tmp16999 := Call(__e, PrimFunc(symshen_4newpv), V5191)


__e.TailApply(tmp16990, tmp16999)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17004 := MakeNative(func(__e *ControlFlow) {
Z5217 := __e.Get(1)
_ = Z5217
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5218 := __e.Get(1)
_ = Z5218
tmp17005 := Call(__e, W5209, W5214)


tmp17006 := Call(__e, tmp17005, Z5217)


__e.TailApply(tmp17006, Z5218)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16961, tmp17004)
return


}, 1)

tmp17007 := PrimTail(W5208)

tmp17008 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17007, V5191)


__e.TailApply(tmp16960, tmp17008)
return


}, 1)

tmp17009 := PrimHead(W5208)

__e.TailApply(tmp16959, tmp17009)
return


} else {
tmp17026 := Call(__e, PrimFunc(symshen_4pvar_2), W5208)


if True == tmp17026 {
tmp17010 := MakeNative(func(__e *ControlFlow) {
W5229 := __e.Get(1)
_ = W5229
tmp17011 := MakeNative(func(__e *ControlFlow) {
W5230 := __e.Get(1)
_ = W5230
tmp17012 := MakeNative(func(__e *ControlFlow) {
W5231 := __e.Get(1)
_ = W5231
tmp17013 := PrimCons(W5231, Nil)

tmp17014 := PrimCons(W5230, tmp17013)

tmp17015 := PrimCons(W5229, tmp17014)

tmp17016 := MakeNative(func(__e *ControlFlow) {
tmp17017 := Call(__e, W5209, W5229)


tmp17018 := Call(__e, tmp17017, W5230)


__e.TailApply(tmp17018, W5231)
return


}, 0)

tmp17019 := Call(__e, PrimFunc(symshen_4bind_b), W5208, tmp17015, V5191, tmp17016)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17019)
return


}, 1)

tmp17020 := Call(__e, PrimFunc(symshen_4newpv), V5191)


tmp17021 := Call(__e, tmp17012, tmp17020)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17021)
return


}, 1)

tmp17022 := Call(__e, PrimFunc(symshen_4newpv), V5191)


tmp17023 := Call(__e, tmp17011, tmp17022)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17023)
return


}, 1)

tmp17024 := Call(__e, PrimFunc(symshen_4newpv), V5191)


__e.TailApply(tmp17010, tmp17024)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17029 := MakeNative(func(__e *ControlFlow) {
Z5210 := __e.Get(1)
_ = Z5210
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5211 := __e.Get(1)
_ = Z5211
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5212 := __e.Get(1)
_ = Z5212
tmp17030 := MakeNative(func(__e *ControlFlow) {
W5213 := __e.Get(1)
_ = W5213
tmp17031 := Call(__e, W5203, Z5210)


tmp17032 := Call(__e, tmp17031, Z5211)


tmp17033 := Call(__e, tmp17032, Z5212)


__e.TailApply(tmp17033, W5213)
return


}, 1)

tmp17034 := PrimTail(W5202)

__e.TailApply(tmp17030, tmp17034)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16958, tmp17029)
return


}, 1)

tmp17035 := PrimHead(W5202)

tmp17036 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17035, V5191)


__e.TailApply(tmp16957, tmp17036)
return


} else {
tmp17058 := Call(__e, PrimFunc(symshen_4pvar_2), W5202)


if True == tmp17058 {
tmp17037 := MakeNative(func(__e *ControlFlow) {
W5232 := __e.Get(1)
_ = W5232
tmp17038 := MakeNative(func(__e *ControlFlow) {
W5233 := __e.Get(1)
_ = W5233
tmp17039 := MakeNative(func(__e *ControlFlow) {
W5234 := __e.Get(1)
_ = W5234
tmp17040 := MakeNative(func(__e *ControlFlow) {
W5235 := __e.Get(1)
_ = W5235
tmp17041 := PrimCons(W5234, Nil)

tmp17042 := PrimCons(W5233, tmp17041)

tmp17043 := PrimCons(W5232, tmp17042)

tmp17044 := PrimCons(tmp17043, W5235)

tmp17045 := MakeNative(func(__e *ControlFlow) {
tmp17046 := Call(__e, W5203, W5232)


tmp17047 := Call(__e, tmp17046, W5233)


tmp17048 := Call(__e, tmp17047, W5234)


__e.TailApply(tmp17048, W5235)
return


}, 0)

tmp17049 := Call(__e, PrimFunc(symshen_4bind_b), W5202, tmp17044, V5191, tmp17045)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17049)
return


}, 1)

tmp17050 := Call(__e, PrimFunc(symshen_4newpv), V5191)


tmp17051 := Call(__e, tmp17040, tmp17050)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17051)
return


}, 1)

tmp17052 := Call(__e, PrimFunc(symshen_4newpv), V5191)


tmp17053 := Call(__e, tmp17039, tmp17052)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17053)
return


}, 1)

tmp17054 := Call(__e, PrimFunc(symshen_4newpv), V5191)


tmp17055 := Call(__e, tmp17038, tmp17054)


__e.TailApply(PrimFunc(symshen_4gc), V5191, tmp17055)
return


}, 1)

tmp17056 := Call(__e, PrimFunc(symshen_4newpv), V5191)


__e.TailApply(tmp17037, tmp17056)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17061 := MakeNative(func(__e *ControlFlow) {
Z5204 := __e.Get(1)
_ = Z5204
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5205 := __e.Get(1)
_ = Z5205
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5206 := __e.Get(1)
_ = Z5206
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5207 := __e.Get(1)
_ = Z5207
tmp17062 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17062

tmp17063 := MakeNative(func(__e *ControlFlow) {
tmp17064 := PrimIntern(MakeString(":"))

tmp17065 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), W5201, Z5207, V5191, V5192, V5193, V5194)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5205, tmp17064, V5191, V5192, V5193, tmp17065)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5204, W5200, V5191, V5192, V5193, tmp17063)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16956, tmp17061)
return


}, 1)

tmp17066 := Call(__e, PrimFunc(symshen_4lazyderef), V5190, V5191)


__e.TailApply(tmp16955, tmp17066)
return


}, 1)

tmp17067 := PrimTail(W5199)

__e.TailApply(tmp16954, tmp17067)
return


}, 1)

tmp17068 := PrimHead(W5199)

__e.TailApply(tmp16953, tmp17068)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17071 := Call(__e, PrimFunc(symshen_4lazyderef), V5189, V5191)


__e.TailApply(tmp16952, tmp17071)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5195)
return
}


}, 1)

tmp17091 := Call(__e, PrimFunc(symshen_4unlocked_2), V5192)


var ifres17076 Obj

if True == tmp17091 {
tmp17077 := MakeNative(func(__e *ControlFlow) {
W5196 := __e.Get(1)
_ = W5196
tmp17088 := PrimEqual(W5196, Nil)

if True == tmp17088 {
tmp17078 := MakeNative(func(__e *ControlFlow) {
W5197 := __e.Get(1)
_ = W5197
tmp17079 := MakeNative(func(__e *ControlFlow) {
W5198 := __e.Get(1)
_ = W5198
tmp17083 := PrimEqual(W5197, Nil)

if True == tmp17083 {
__e.TailApply(PrimFunc(symthaw), W5198)
return
} else {
tmp17081 := Call(__e, PrimFunc(symshen_4pvar_2), W5197)


if True == tmp17081 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5197, Nil, V5191, W5198)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17084 := MakeNative(func(__e *ControlFlow) {
tmp17085 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17085

__e.TailApply(PrimFunc(symthaw), V5194)
return


}, 0)

__e.TailApply(tmp17079, tmp17084)
return


}, 1)

tmp17086 := Call(__e, PrimFunc(symshen_4lazyderef), V5190, V5191)


__e.TailApply(tmp17078, tmp17086)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17089 := Call(__e, PrimFunc(symshen_4lazyderef), V5189, V5191)


tmp17090 := Call(__e, tmp17077, tmp17089)


ifres17076 = tmp17090


} else {
ifres17076 = False


}

__e.TailApply(tmp16951, ifres17076)
return


}, 6)

tmp17092 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp16950)


_ = tmp17092

tmp17093 := MakeNative(func(__e *ControlFlow) {
V5236 := __e.Get(1)
_ = V5236
V5237 := __e.Get(2)
_ = V5237
V5238 := __e.Get(3)
_ = V5238
V5239 := __e.Get(4)
_ = V5239
V5240 := __e.Get(5)
_ = V5240
V5241 := __e.Get(6)
_ = V5241
V5242 := __e.Get(7)
_ = V5242
tmp17094 := MakeNative(func(__e *ControlFlow) {
W5243 := __e.Get(1)
_ = W5243
tmp17095 := MakeNative(func(__e *ControlFlow) {
W5244 := __e.Get(1)
_ = W5244
tmp17105 := PrimEqual(W5244, False)

if True == tmp17105 {
tmp17096 := MakeNative(func(__e *ControlFlow) {
W5253 := __e.Get(1)
_ = W5253
tmp17098 := PrimEqual(W5253, False)

if True == tmp17098 {
__e.TailApply(PrimFunc(symshen_4unlock), V5240, W5243)
return
} else {
__e.Return(W5253)
return
}


}, 1)

tmp17103 := Call(__e, PrimFunc(symshen_4unlocked_2), V5240)


var ifres17099 Obj

if True == tmp17103 {
tmp17100 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17100

tmp17101 := Call(__e, PrimFunc(symshen_4curry), V5236)


tmp17102 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17101, V5237, V5238, V5239, V5240, W5243, V5242)


ifres17099 = tmp17102


} else {
ifres17099 = False


}

__e.TailApply(tmp17096, ifres17099)
return


} else {
__e.Return(W5244)
return
}


}, 1)

tmp17150 := Call(__e, PrimFunc(symshen_4unlocked_2), V5240)


var ifres17106 Obj

if True == tmp17150 {
tmp17107 := MakeNative(func(__e *ControlFlow) {
W5245 := __e.Get(1)
_ = W5245
tmp17147 := PrimIsPair(W5245)

if True == tmp17147 {
tmp17108 := MakeNative(func(__e *ControlFlow) {
W5246 := __e.Get(1)
_ = W5246
tmp17143 := PrimEqual(W5246, symwhere)

if True == tmp17143 {
tmp17109 := MakeNative(func(__e *ControlFlow) {
W5247 := __e.Get(1)
_ = W5247
tmp17139 := PrimIsPair(W5247)

if True == tmp17139 {
tmp17110 := MakeNative(func(__e *ControlFlow) {
W5248 := __e.Get(1)
_ = W5248
tmp17111 := MakeNative(func(__e *ControlFlow) {
W5249 := __e.Get(1)
_ = W5249
tmp17134 := PrimIsPair(W5249)

if True == tmp17134 {
tmp17112 := MakeNative(func(__e *ControlFlow) {
W5250 := __e.Get(1)
_ = W5250
tmp17113 := MakeNative(func(__e *ControlFlow) {
W5251 := __e.Get(1)
_ = W5251
tmp17129 := PrimEqual(W5251, Nil)

if True == tmp17129 {
tmp17114 := MakeNative(func(__e *ControlFlow) {
W5252 := __e.Get(1)
_ = W5252
tmp17115 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17115

tmp17116 := MakeNative(func(__e *ControlFlow) {
tmp17117 := Call(__e, PrimFunc(symshen_4curry), W5248)


tmp17118 := MakeNative(func(__e *ControlFlow) {
tmp17119 := MakeNative(func(__e *ControlFlow) {
tmp17120 := MakeNative(func(__e *ControlFlow) {
tmp17121 := PrimIntern(MakeString(":"))

tmp17122 := PrimCons(symverified, Nil)

tmp17123 := PrimCons(tmp17121, tmp17122)

tmp17124 := PrimCons(W5252, tmp17123)

tmp17125 := PrimCons(tmp17124, V5238)

__e.TailApply(PrimFunc(symshen_4t_d_1correct), W5250, V5237, tmp17125, V5239, V5240, W5243, V5242)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5239, V5240, W5243, tmp17120)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5252, symboolean, V5238, V5239, V5240, W5243, tmp17119)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5252, tmp17117, V5239, V5240, W5243, tmp17118)
return


}, 0)

tmp17126 := Call(__e, PrimFunc(symshen_4cut), V5239, V5240, W5243, tmp17116)


__e.TailApply(PrimFunc(symshen_4gc), V5239, tmp17126)
return


}, 1)

tmp17127 := Call(__e, PrimFunc(symshen_4newpv), V5239)


__e.TailApply(tmp17114, tmp17127)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17130 := PrimTail(W5249)

tmp17131 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17130, V5239)


__e.TailApply(tmp17113, tmp17131)
return


}, 1)

tmp17132 := PrimHead(W5249)

__e.TailApply(tmp17112, tmp17132)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17135 := PrimTail(W5247)

tmp17136 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17135, V5239)


__e.TailApply(tmp17111, tmp17136)
return


}, 1)

tmp17137 := PrimHead(W5247)

__e.TailApply(tmp17110, tmp17137)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17140 := PrimTail(W5245)

tmp17141 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17140, V5239)


__e.TailApply(tmp17109, tmp17141)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17144 := PrimHead(W5245)

tmp17145 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17144, V5239)


__e.TailApply(tmp17108, tmp17145)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17148 := Call(__e, PrimFunc(symshen_4lazyderef), V5236, V5239)


tmp17149 := Call(__e, tmp17107, tmp17148)


ifres17106 = tmp17149


} else {
ifres17106 = False


}

__e.TailApply(tmp17095, ifres17106)
return


}, 1)

tmp17151 := PrimNumberAdd(V5241, MakeNumber(1))

__e.TailApply(tmp17094, tmp17151)
return


}, 7)

tmp17152 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17093)


_ = tmp17152

tmp17153 := MakeNative(func(__e *ControlFlow) {
V5254 := __e.Get(1)
_ = V5254
V5255 := __e.Get(2)
_ = V5255
V5256 := __e.Get(3)
_ = V5256
V5257 := __e.Get(4)
_ = V5257
V5258 := __e.Get(5)
_ = V5258
V5259 := __e.Get(6)
_ = V5259
V5260 := __e.Get(7)
_ = V5260
V5261 := __e.Get(8)
_ = V5261
tmp17154 := MakeNative(func(__e *ControlFlow) {
W5262 := __e.Get(1)
_ = W5262
tmp17196 := PrimEqual(W5262, False)

if True == tmp17196 {
tmp17194 := Call(__e, PrimFunc(symshen_4unlocked_2), V5259)


if True == tmp17194 {
tmp17155 := MakeNative(func(__e *ControlFlow) {
W5264 := __e.Get(1)
_ = W5264
tmp17191 := PrimIsPair(W5264)

if True == tmp17191 {
tmp17156 := MakeNative(func(__e *ControlFlow) {
W5265 := __e.Get(1)
_ = W5265
tmp17157 := MakeNative(func(__e *ControlFlow) {
W5266 := __e.Get(1)
_ = W5266
tmp17158 := MakeNative(func(__e *ControlFlow) {
W5267 := __e.Get(1)
_ = W5267
tmp17186 := PrimIsPair(W5267)

if True == tmp17186 {
tmp17159 := MakeNative(func(__e *ControlFlow) {
W5268 := __e.Get(1)
_ = W5268
tmp17160 := MakeNative(func(__e *ControlFlow) {
W5269 := __e.Get(1)
_ = W5269
tmp17181 := PrimIsPair(W5269)

if True == tmp17181 {
tmp17161 := MakeNative(func(__e *ControlFlow) {
W5270 := __e.Get(1)
_ = W5270
tmp17177 := PrimEqual(W5270, sym_1_1_6)

if True == tmp17177 {
tmp17162 := MakeNative(func(__e *ControlFlow) {
W5271 := __e.Get(1)
_ = W5271
tmp17173 := PrimIsPair(W5271)

if True == tmp17173 {
tmp17163 := MakeNative(func(__e *ControlFlow) {
W5272 := __e.Get(1)
_ = W5272
tmp17164 := MakeNative(func(__e *ControlFlow) {
W5273 := __e.Get(1)
_ = W5273
tmp17168 := PrimEqual(W5273, Nil)

if True == tmp17168 {
tmp17165 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17165

tmp17166 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W5266, W5272, V5256, V5257, V5258, V5259, V5260, V5261)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5265, W5268, V5256, V5258, V5259, V5260, tmp17166)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17169 := PrimTail(W5271)

tmp17170 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17169, V5258)


__e.TailApply(tmp17164, tmp17170)
return


}, 1)

tmp17171 := PrimHead(W5271)

__e.TailApply(tmp17163, tmp17171)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17174 := PrimTail(W5269)

tmp17175 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17174, V5258)


__e.TailApply(tmp17162, tmp17175)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17178 := PrimHead(W5269)

tmp17179 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17178, V5258)


__e.TailApply(tmp17161, tmp17179)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17182 := PrimTail(W5267)

tmp17183 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17182, V5258)


__e.TailApply(tmp17160, tmp17183)
return


}, 1)

tmp17184 := PrimHead(W5267)

__e.TailApply(tmp17159, tmp17184)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17187 := Call(__e, PrimFunc(symshen_4lazyderef), V5255, V5258)


__e.TailApply(tmp17158, tmp17187)
return


}, 1)

tmp17188 := PrimTail(W5264)

__e.TailApply(tmp17157, tmp17188)
return


}, 1)

tmp17189 := PrimHead(W5264)

__e.TailApply(tmp17156, tmp17189)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17192 := Call(__e, PrimFunc(symshen_4lazyderef), V5254, V5258)


__e.TailApply(tmp17155, tmp17192)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5262)
return
}


}, 1)

tmp17204 := Call(__e, PrimFunc(symshen_4unlocked_2), V5259)


var ifres17197 Obj

if True == tmp17204 {
tmp17198 := MakeNative(func(__e *ControlFlow) {
W5263 := __e.Get(1)
_ = W5263
tmp17201 := PrimEqual(W5263, Nil)

if True == tmp17201 {
tmp17199 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17199

__e.TailApply(PrimFunc(symis_b), V5255, V5257, V5258, V5259, V5260, V5261)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17202 := Call(__e, PrimFunc(symshen_4lazyderef), V5254, V5258)


tmp17203 := Call(__e, tmp17198, tmp17202)


ifres17197 = tmp17203


} else {
ifres17197 = False


}

__e.TailApply(tmp17154, ifres17197)
return


}, 8)

tmp17205 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17153)


_ = tmp17205

tmp17206 := MakeNative(func(__e *ControlFlow) {
V5274 := __e.Get(1)
_ = V5274
tmp17215 := PrimIsVector(V5274)

if True == tmp17215 {
tmp17212 := PrimIsString(V5274)

tmp17213 := PrimNot(tmp17212)

var ifres17208 Obj

if True == tmp17213 {
tmp17210 := PrimVectorGet(V5274, MakeNumber(0))

tmp17211 := PrimEqual(tmp17210, symshen_4print_1freshterm)

var ifres17209 Obj

if True == tmp17211 {
ifres17209 = True


} else {
ifres17209 = False


}

ifres17208 = ifres17209


} else {
ifres17208 = False


}

if True == ifres17208 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17206)
return




}, 0)

