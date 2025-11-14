package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12093 := MakeNative(func(__e *ControlFlow) {
V3016 := __e.Get(1)
_ = V3016
tmp12094 := MakeNative(func(__e *ControlFlow) {
W3017 := __e.Get(1)
_ = W3017
tmp12096 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3017)


if True == tmp12096 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3017)
return
}


}, 1)

tmp12116 := PrimIsPair(V3016)

var ifres12097 Obj

if True == tmp12116 {
tmp12098 := MakeNative(func(__e *ControlFlow) {
W3018 := __e.Get(1)
_ = W3018
tmp12099 := MakeNative(func(__e *ControlFlow) {
W3019 := __e.Get(1)
_ = W3019
tmp12100 := MakeNative(func(__e *ControlFlow) {
W3020 := __e.Get(1)
_ = W3020
tmp12110 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3020)


if True == tmp12110 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12101 := MakeNative(func(__e *ControlFlow) {
W3021 := __e.Get(1)
_ = W3021
tmp12102 := MakeNative(func(__e *ControlFlow) {
W3022 := __e.Get(1)
_ = W3022
tmp12103 := MakeNative(func(__e *ControlFlow) {
W3023 := __e.Get(1)
_ = W3023
tmp12104 := Call(__e, PrimFunc(symfn), W3018)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), W3018, tmp12104)
return


}, 1)

tmp12105 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W3018, W3021)


tmp12106 := Call(__e, tmp12103, tmp12105)


__e.TailApply(PrimFunc(symshen_4comb), W3022, tmp12106)
return


}, 1)

tmp12107 := Call(__e, PrimFunc(symshen_4in_1_6), W3020)


__e.TailApply(tmp12102, tmp12107)
return


}, 1)

tmp12108 := Call(__e, PrimFunc(symshen_4_5_1out), W3020)


__e.TailApply(tmp12101, tmp12108)
return


}


}, 1)

tmp12111 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3019)


__e.TailApply(tmp12100, tmp12111)
return


}, 1)

tmp12112 := Call(__e, PrimFunc(symtail), V3016)


__e.TailApply(tmp12099, tmp12112)
return


}, 1)

tmp12113 := Call(__e, PrimFunc(symhead), V3016)


tmp12114 := Call(__e, tmp12098, tmp12113)


ifres12097 = tmp12114


} else {
tmp12115 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12097 = tmp12115


}

__e.TailApply(tmp12094, ifres12097)
return


}, 1)

tmp12117 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12093)


_ = tmp12117

tmp12118 := MakeNative(func(__e *ControlFlow) {
V3024 := __e.Get(1)
_ = V3024
tmp12119 := MakeNative(func(__e *ControlFlow) {
W3025 := __e.Get(1)
_ = W3025
tmp12138 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3025)


if True == tmp12138 {
tmp12120 := MakeNative(func(__e *ControlFlow) {
W3032 := __e.Get(1)
_ = W3032
tmp12122 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3032)


if True == tmp12122 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3032)
return
}


}, 1)

tmp12123 := MakeNative(func(__e *ControlFlow) {
W3033 := __e.Get(1)
_ = W3033
tmp12134 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3033)


if True == tmp12134 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12124 := MakeNative(func(__e *ControlFlow) {
W3034 := __e.Get(1)
_ = W3034
tmp12125 := MakeNative(func(__e *ControlFlow) {
W3035 := __e.Get(1)
_ = W3035
tmp12130 := Call(__e, PrimFunc(symempty_2), W3034)


var ifres12126 Obj

if True == tmp12130 {
ifres12126 = Nil


} else {
tmp12127 := Call(__e, PrimFunc(symshen_4app), W3034, MakeString("\n ..."), symshen_4r)


tmp12128 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp12127)

tmp12129 := PrimSimpleError(tmp12128)

ifres12126 = tmp12129


}

__e.TailApply(PrimFunc(symshen_4comb), W3035, ifres12126)
return


}, 1)

tmp12131 := Call(__e, PrimFunc(symshen_4in_1_6), W3033)


__e.TailApply(tmp12125, tmp12131)
return


}, 1)

tmp12132 := Call(__e, PrimFunc(symshen_4_5_1out), W3033)


__e.TailApply(tmp12124, tmp12132)
return


}


}, 1)

tmp12135 := Call(__e, PrimFunc(sym_5_b_6), V3024)


tmp12136 := Call(__e, tmp12123, tmp12135)


__e.TailApply(tmp12120, tmp12136)
return


} else {
__e.Return(W3025)
return
}


}, 1)

tmp12139 := MakeNative(func(__e *ControlFlow) {
W3026 := __e.Get(1)
_ = W3026
tmp12154 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3026)


if True == tmp12154 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12140 := MakeNative(func(__e *ControlFlow) {
W3027 := __e.Get(1)
_ = W3027
tmp12141 := MakeNative(func(__e *ControlFlow) {
W3028 := __e.Get(1)
_ = W3028
tmp12142 := MakeNative(func(__e *ControlFlow) {
W3029 := __e.Get(1)
_ = W3029
tmp12149 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3029)


if True == tmp12149 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12143 := MakeNative(func(__e *ControlFlow) {
W3030 := __e.Get(1)
_ = W3030
tmp12144 := MakeNative(func(__e *ControlFlow) {
W3031 := __e.Get(1)
_ = W3031
tmp12145 := Call(__e, PrimFunc(symappend), W3027, W3030)


__e.TailApply(PrimFunc(symshen_4comb), W3031, tmp12145)
return


}, 1)

tmp12146 := Call(__e, PrimFunc(symshen_4in_1_6), W3029)


__e.TailApply(tmp12144, tmp12146)
return


}, 1)

tmp12147 := Call(__e, PrimFunc(symshen_4_5_1out), W3029)


__e.TailApply(tmp12143, tmp12147)
return


}


}, 1)

tmp12150 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3028)


__e.TailApply(tmp12142, tmp12150)
return


}, 1)

tmp12151 := Call(__e, PrimFunc(symshen_4in_1_6), W3026)


__e.TailApply(tmp12141, tmp12151)
return


}, 1)

tmp12152 := Call(__e, PrimFunc(symshen_4_5_1out), W3026)


__e.TailApply(tmp12140, tmp12152)
return


}


}, 1)

tmp12155 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3024)


tmp12156 := Call(__e, tmp12139, tmp12155)


__e.TailApply(tmp12119, tmp12156)
return


}, 1)

tmp12157 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12118)


_ = tmp12157

tmp12158 := MakeNative(func(__e *ControlFlow) {
V3036 := __e.Get(1)
_ = V3036
tmp12159 := MakeNative(func(__e *ControlFlow) {
W3037 := __e.Get(1)
_ = W3037
tmp12173 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3037)


if True == tmp12173 {
tmp12160 := MakeNative(func(__e *ControlFlow) {
W3041 := __e.Get(1)
_ = W3041
tmp12162 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3041)


if True == tmp12162 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3041)
return
}


}, 1)

tmp12163 := MakeNative(func(__e *ControlFlow) {
W3042 := __e.Get(1)
_ = W3042
tmp12169 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3042)


if True == tmp12169 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12164 := MakeNative(func(__e *ControlFlow) {
W3043 := __e.Get(1)
_ = W3043
tmp12165 := MakeNative(func(__e *ControlFlow) {
W3044 := __e.Get(1)
_ = W3044
__e.TailApply(PrimFunc(symshen_4comb), W3044, W3043)
return
}, 1)

tmp12166 := Call(__e, PrimFunc(symshen_4in_1_6), W3042)


__e.TailApply(tmp12165, tmp12166)
return


}, 1)

tmp12167 := Call(__e, PrimFunc(symshen_4_5_1out), W3042)


__e.TailApply(tmp12164, tmp12167)
return


}


}, 1)

tmp12170 := Call(__e, PrimFunc(symshen_4_5double_6), V3036)


tmp12171 := Call(__e, tmp12163, tmp12170)


__e.TailApply(tmp12160, tmp12171)
return


} else {
__e.Return(W3037)
return
}


}, 1)

tmp12174 := MakeNative(func(__e *ControlFlow) {
W3038 := __e.Get(1)
_ = W3038
tmp12180 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3038)


if True == tmp12180 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12175 := MakeNative(func(__e *ControlFlow) {
W3039 := __e.Get(1)
_ = W3039
tmp12176 := MakeNative(func(__e *ControlFlow) {
W3040 := __e.Get(1)
_ = W3040
__e.TailApply(PrimFunc(symshen_4comb), W3040, W3039)
return
}, 1)

tmp12177 := Call(__e, PrimFunc(symshen_4in_1_6), W3038)


__e.TailApply(tmp12176, tmp12177)
return


}, 1)

tmp12178 := Call(__e, PrimFunc(symshen_4_5_1out), W3038)


__e.TailApply(tmp12175, tmp12178)
return


}


}, 1)

tmp12181 := Call(__e, PrimFunc(symshen_4_5single_6), V3036)


tmp12182 := Call(__e, tmp12174, tmp12181)


__e.TailApply(tmp12159, tmp12182)
return


}, 1)

tmp12183 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12158)


_ = tmp12183

tmp12184 := MakeNative(func(__e *ControlFlow) {
V3045 := __e.Get(1)
_ = V3045
tmp12185 := MakeNative(func(__e *ControlFlow) {
W3046 := __e.Get(1)
_ = W3046
tmp12187 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3046)


if True == tmp12187 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3046)
return
}


}, 1)

tmp12188 := MakeNative(func(__e *ControlFlow) {
W3047 := __e.Get(1)
_ = W3047
tmp12226 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3047)


if True == tmp12226 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12189 := MakeNative(func(__e *ControlFlow) {
W3048 := __e.Get(1)
_ = W3048
tmp12190 := MakeNative(func(__e *ControlFlow) {
W3049 := __e.Get(1)
_ = W3049
tmp12191 := MakeNative(func(__e *ControlFlow) {
W3050 := __e.Get(1)
_ = W3050
tmp12221 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3050)


if True == tmp12221 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12192 := MakeNative(func(__e *ControlFlow) {
W3051 := __e.Get(1)
_ = W3051
tmp12193 := MakeNative(func(__e *ControlFlow) {
W3052 := __e.Get(1)
_ = W3052
tmp12194 := MakeNative(func(__e *ControlFlow) {
W3053 := __e.Get(1)
_ = W3053
tmp12216 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3053)


if True == tmp12216 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12195 := MakeNative(func(__e *ControlFlow) {
W3054 := __e.Get(1)
_ = W3054
tmp12196 := MakeNative(func(__e *ControlFlow) {
W3055 := __e.Get(1)
_ = W3055
tmp12212 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3055)


if True == tmp12212 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12197 := MakeNative(func(__e *ControlFlow) {
W3056 := __e.Get(1)
_ = W3056
tmp12198 := MakeNative(func(__e *ControlFlow) {
W3057 := __e.Get(1)
_ = W3057
tmp12199 := MakeNative(func(__e *ControlFlow) {
W3058 := __e.Get(1)
_ = W3058
tmp12207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3058)


if True == tmp12207 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12200 := MakeNative(func(__e *ControlFlow) {
W3059 := __e.Get(1)
_ = W3059
tmp12201 := PrimCons(W3056, Nil)

tmp12202 := PrimCons(W3051, tmp12201)

tmp12203 := PrimCons(W3048, tmp12202)

tmp12204 := PrimCons(tmp12203, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3059, tmp12204)
return


}, 1)

tmp12205 := Call(__e, PrimFunc(symshen_4in_1_6), W3058)


__e.TailApply(tmp12200, tmp12205)
return


}


}, 1)

tmp12208 := Call(__e, PrimFunc(symshen_4_5sc_6), W3057)


__e.TailApply(tmp12199, tmp12208)
return


}, 1)

tmp12209 := Call(__e, PrimFunc(symshen_4in_1_6), W3055)


__e.TailApply(tmp12198, tmp12209)
return


}, 1)

tmp12210 := Call(__e, PrimFunc(symshen_4_5_1out), W3055)


__e.TailApply(tmp12197, tmp12210)
return


}


}, 1)

tmp12213 := Call(__e, PrimFunc(symshen_4_5conc_6), W3054)


__e.TailApply(tmp12196, tmp12213)
return


}, 1)

tmp12214 := Call(__e, PrimFunc(symshen_4in_1_6), W3053)


__e.TailApply(tmp12195, tmp12214)
return


}


}, 1)

tmp12217 := Call(__e, PrimFunc(symshen_4_5sng_6), W3052)


__e.TailApply(tmp12194, tmp12217)
return


}, 1)

tmp12218 := Call(__e, PrimFunc(symshen_4in_1_6), W3050)


__e.TailApply(tmp12193, tmp12218)
return


}, 1)

tmp12219 := Call(__e, PrimFunc(symshen_4_5_1out), W3050)


__e.TailApply(tmp12192, tmp12219)
return


}


}, 1)

tmp12222 := Call(__e, PrimFunc(symshen_4_5prems_6), W3049)


__e.TailApply(tmp12191, tmp12222)
return


}, 1)

tmp12223 := Call(__e, PrimFunc(symshen_4in_1_6), W3047)


__e.TailApply(tmp12190, tmp12223)
return


}, 1)

tmp12224 := Call(__e, PrimFunc(symshen_4_5_1out), W3047)


__e.TailApply(tmp12189, tmp12224)
return


}


}, 1)

tmp12227 := Call(__e, PrimFunc(symshen_4_5sides_6), V3045)


tmp12228 := Call(__e, tmp12188, tmp12227)


__e.TailApply(tmp12185, tmp12228)
return


}, 1)

tmp12229 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12184)


_ = tmp12229

tmp12230 := MakeNative(func(__e *ControlFlow) {
V3060 := __e.Get(1)
_ = V3060
tmp12231 := MakeNative(func(__e *ControlFlow) {
W3061 := __e.Get(1)
_ = W3061
tmp12233 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3061)


if True == tmp12233 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3061)
return
}


}, 1)

tmp12234 := MakeNative(func(__e *ControlFlow) {
W3062 := __e.Get(1)
_ = W3062
tmp12271 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3062)


if True == tmp12271 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12235 := MakeNative(func(__e *ControlFlow) {
W3063 := __e.Get(1)
_ = W3063
tmp12236 := MakeNative(func(__e *ControlFlow) {
W3064 := __e.Get(1)
_ = W3064
tmp12237 := MakeNative(func(__e *ControlFlow) {
W3065 := __e.Get(1)
_ = W3065
tmp12266 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3065)


if True == tmp12266 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12238 := MakeNative(func(__e *ControlFlow) {
W3066 := __e.Get(1)
_ = W3066
tmp12239 := MakeNative(func(__e *ControlFlow) {
W3067 := __e.Get(1)
_ = W3067
tmp12240 := MakeNative(func(__e *ControlFlow) {
W3068 := __e.Get(1)
_ = W3068
tmp12261 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3068)


if True == tmp12261 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12241 := MakeNative(func(__e *ControlFlow) {
W3069 := __e.Get(1)
_ = W3069
tmp12242 := MakeNative(func(__e *ControlFlow) {
W3070 := __e.Get(1)
_ = W3070
tmp12257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3070)


if True == tmp12257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12243 := MakeNative(func(__e *ControlFlow) {
W3071 := __e.Get(1)
_ = W3071
tmp12244 := MakeNative(func(__e *ControlFlow) {
W3072 := __e.Get(1)
_ = W3072
tmp12245 := MakeNative(func(__e *ControlFlow) {
W3073 := __e.Get(1)
_ = W3073
tmp12252 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3073)


if True == tmp12252 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12246 := MakeNative(func(__e *ControlFlow) {
W3074 := __e.Get(1)
_ = W3074
tmp12247 := PrimCons(W3071, Nil)

tmp12248 := PrimCons(Nil, tmp12247)

tmp12249 := Call(__e, PrimFunc(symshen_4lr_1rule), W3063, W3066, tmp12248)


__e.TailApply(PrimFunc(symshen_4comb), W3074, tmp12249)
return


}, 1)

tmp12250 := Call(__e, PrimFunc(symshen_4in_1_6), W3073)


__e.TailApply(tmp12246, tmp12250)
return


}


}, 1)

tmp12253 := Call(__e, PrimFunc(symshen_4_5sc_6), W3072)


__e.TailApply(tmp12245, tmp12253)
return


}, 1)

tmp12254 := Call(__e, PrimFunc(symshen_4in_1_6), W3070)


__e.TailApply(tmp12244, tmp12254)
return


}, 1)

tmp12255 := Call(__e, PrimFunc(symshen_4_5_1out), W3070)


__e.TailApply(tmp12243, tmp12255)
return


}


}, 1)

tmp12258 := Call(__e, PrimFunc(symshen_4_5formula_6), W3069)


__e.TailApply(tmp12242, tmp12258)
return


}, 1)

tmp12259 := Call(__e, PrimFunc(symshen_4in_1_6), W3068)


__e.TailApply(tmp12241, tmp12259)
return


}


}, 1)

tmp12262 := Call(__e, PrimFunc(symshen_4_5dbl_6), W3067)


__e.TailApply(tmp12240, tmp12262)
return


}, 1)

tmp12263 := Call(__e, PrimFunc(symshen_4in_1_6), W3065)


__e.TailApply(tmp12239, tmp12263)
return


}, 1)

tmp12264 := Call(__e, PrimFunc(symshen_4_5_1out), W3065)


__e.TailApply(tmp12238, tmp12264)
return


}


}, 1)

tmp12267 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3064)


__e.TailApply(tmp12237, tmp12267)
return


}, 1)

tmp12268 := Call(__e, PrimFunc(symshen_4in_1_6), W3062)


__e.TailApply(tmp12236, tmp12268)
return


}, 1)

tmp12269 := Call(__e, PrimFunc(symshen_4_5_1out), W3062)


__e.TailApply(tmp12235, tmp12269)
return


}


}, 1)

tmp12272 := Call(__e, PrimFunc(symshen_4_5sides_6), V3060)


tmp12273 := Call(__e, tmp12234, tmp12272)


__e.TailApply(tmp12231, tmp12273)
return


}, 1)

tmp12274 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12230)


_ = tmp12274

tmp12275 := MakeNative(func(__e *ControlFlow) {
V3075 := __e.Get(1)
_ = V3075
tmp12276 := MakeNative(func(__e *ControlFlow) {
W3076 := __e.Get(1)
_ = W3076
tmp12299 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3076)


if True == tmp12299 {
tmp12277 := MakeNative(func(__e *ControlFlow) {
W3085 := __e.Get(1)
_ = W3085
tmp12279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3085)


if True == tmp12279 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3085)
return
}


}, 1)

tmp12280 := MakeNative(func(__e *ControlFlow) {
W3086 := __e.Get(1)
_ = W3086
tmp12295 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3086)


if True == tmp12295 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12281 := MakeNative(func(__e *ControlFlow) {
W3087 := __e.Get(1)
_ = W3087
tmp12282 := MakeNative(func(__e *ControlFlow) {
W3088 := __e.Get(1)
_ = W3088
tmp12283 := MakeNative(func(__e *ControlFlow) {
W3089 := __e.Get(1)
_ = W3089
tmp12290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3089)


if True == tmp12290 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12284 := MakeNative(func(__e *ControlFlow) {
W3090 := __e.Get(1)
_ = W3090
tmp12285 := PrimCons(W3087, Nil)

tmp12286 := PrimCons(Nil, tmp12285)

tmp12287 := PrimCons(tmp12286, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3090, tmp12287)
return


}, 1)

tmp12288 := Call(__e, PrimFunc(symshen_4in_1_6), W3089)


__e.TailApply(tmp12284, tmp12288)
return


}


}, 1)

tmp12291 := Call(__e, PrimFunc(symshen_4_5sc_6), W3088)


__e.TailApply(tmp12283, tmp12291)
return


}, 1)

tmp12292 := Call(__e, PrimFunc(symshen_4in_1_6), W3086)


__e.TailApply(tmp12282, tmp12292)
return


}, 1)

tmp12293 := Call(__e, PrimFunc(symshen_4_5_1out), W3086)


__e.TailApply(tmp12281, tmp12293)
return


}


}, 1)

tmp12296 := Call(__e, PrimFunc(symshen_4_5formula_6), V3075)


tmp12297 := Call(__e, tmp12280, tmp12296)


__e.TailApply(tmp12277, tmp12297)
return


} else {
__e.Return(W3076)
return
}


}, 1)

tmp12300 := MakeNative(func(__e *ControlFlow) {
W3077 := __e.Get(1)
_ = W3077
tmp12323 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3077)


if True == tmp12323 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12301 := MakeNative(func(__e *ControlFlow) {
W3078 := __e.Get(1)
_ = W3078
tmp12302 := MakeNative(func(__e *ControlFlow) {
W3079 := __e.Get(1)
_ = W3079
tmp12303 := MakeNative(func(__e *ControlFlow) {
W3080 := __e.Get(1)
_ = W3080
tmp12318 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3080)


if True == tmp12318 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12304 := MakeNative(func(__e *ControlFlow) {
W3081 := __e.Get(1)
_ = W3081
tmp12305 := MakeNative(func(__e *ControlFlow) {
W3082 := __e.Get(1)
_ = W3082
tmp12314 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3082)


if True == tmp12314 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12306 := MakeNative(func(__e *ControlFlow) {
W3083 := __e.Get(1)
_ = W3083
tmp12307 := MakeNative(func(__e *ControlFlow) {
W3084 := __e.Get(1)
_ = W3084
tmp12308 := PrimCons(W3078, Nil)

tmp12309 := PrimCons(Nil, tmp12308)

tmp12310 := PrimCons(tmp12309, W3083)

__e.TailApply(PrimFunc(symshen_4comb), W3084, tmp12310)
return


}, 1)

tmp12311 := Call(__e, PrimFunc(symshen_4in_1_6), W3082)


__e.TailApply(tmp12307, tmp12311)
return


}, 1)

tmp12312 := Call(__e, PrimFunc(symshen_4_5_1out), W3082)


__e.TailApply(tmp12306, tmp12312)
return


}


}, 1)

tmp12315 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3081)


__e.TailApply(tmp12305, tmp12315)
return


}, 1)

tmp12316 := Call(__e, PrimFunc(symshen_4in_1_6), W3080)


__e.TailApply(tmp12304, tmp12316)
return


}


}, 1)

tmp12319 := Call(__e, PrimFunc(symshen_4_5sc_6), W3079)


__e.TailApply(tmp12303, tmp12319)
return


}, 1)

tmp12320 := Call(__e, PrimFunc(symshen_4in_1_6), W3077)


__e.TailApply(tmp12302, tmp12320)
return


}, 1)

tmp12321 := Call(__e, PrimFunc(symshen_4_5_1out), W3077)


__e.TailApply(tmp12301, tmp12321)
return


}


}, 1)

tmp12324 := Call(__e, PrimFunc(symshen_4_5formula_6), V3075)


tmp12325 := Call(__e, tmp12300, tmp12324)


__e.TailApply(tmp12276, tmp12325)
return


}, 1)

tmp12326 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12275)


_ = tmp12326

tmp12327 := MakeNative(func(__e *ControlFlow) {
V3091 := __e.Get(1)
_ = V3091
tmp12328 := MakeNative(func(__e *ControlFlow) {
W3092 := __e.Get(1)
_ = W3092
tmp12344 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3092)


if True == tmp12344 {
tmp12329 := MakeNative(func(__e *ControlFlow) {
W3100 := __e.Get(1)
_ = W3100
tmp12331 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3100)


if True == tmp12331 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3100)
return
}


}, 1)

tmp12332 := MakeNative(func(__e *ControlFlow) {
W3101 := __e.Get(1)
_ = W3101
tmp12340 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3101)


if True == tmp12340 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12333 := MakeNative(func(__e *ControlFlow) {
W3102 := __e.Get(1)
_ = W3102
tmp12334 := MakeNative(func(__e *ControlFlow) {
W3103 := __e.Get(1)
_ = W3103
tmp12335 := PrimCons(W3102, Nil)

tmp12336 := PrimCons(Nil, tmp12335)

__e.TailApply(PrimFunc(symshen_4comb), W3103, tmp12336)
return


}, 1)

tmp12337 := Call(__e, PrimFunc(symshen_4in_1_6), W3101)


__e.TailApply(tmp12334, tmp12337)
return


}, 1)

tmp12338 := Call(__e, PrimFunc(symshen_4_5_1out), W3101)


__e.TailApply(tmp12333, tmp12338)
return


}


}, 1)

tmp12341 := Call(__e, PrimFunc(symshen_4_5formula_6), V3091)


tmp12342 := Call(__e, tmp12332, tmp12341)


__e.TailApply(tmp12329, tmp12342)
return


} else {
__e.Return(W3092)
return
}


}, 1)

tmp12345 := MakeNative(func(__e *ControlFlow) {
W3093 := __e.Get(1)
_ = W3093
tmp12365 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3093)


if True == tmp12365 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12346 := MakeNative(func(__e *ControlFlow) {
W3094 := __e.Get(1)
_ = W3094
tmp12347 := MakeNative(func(__e *ControlFlow) {
W3095 := __e.Get(1)
_ = W3095
tmp12361 := Call(__e, PrimFunc(symshen_4hds_a_2), W3095, sym_6_6)


if True == tmp12361 {
tmp12348 := MakeNative(func(__e *ControlFlow) {
W3096 := __e.Get(1)
_ = W3096
tmp12349 := MakeNative(func(__e *ControlFlow) {
W3097 := __e.Get(1)
_ = W3097
tmp12357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3097)


if True == tmp12357 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12350 := MakeNative(func(__e *ControlFlow) {
W3098 := __e.Get(1)
_ = W3098
tmp12351 := MakeNative(func(__e *ControlFlow) {
W3099 := __e.Get(1)
_ = W3099
tmp12352 := PrimCons(W3098, Nil)

tmp12353 := PrimCons(W3094, tmp12352)

__e.TailApply(PrimFunc(symshen_4comb), W3099, tmp12353)
return


}, 1)

tmp12354 := Call(__e, PrimFunc(symshen_4in_1_6), W3097)


__e.TailApply(tmp12351, tmp12354)
return


}, 1)

tmp12355 := Call(__e, PrimFunc(symshen_4_5_1out), W3097)


__e.TailApply(tmp12350, tmp12355)
return


}


}, 1)

tmp12358 := Call(__e, PrimFunc(symshen_4_5formula_6), W3096)


__e.TailApply(tmp12349, tmp12358)
return


}, 1)

tmp12359 := Call(__e, PrimFunc(symtail), W3095)


__e.TailApply(tmp12348, tmp12359)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12362 := Call(__e, PrimFunc(symshen_4in_1_6), W3093)


__e.TailApply(tmp12347, tmp12362)
return


}, 1)

tmp12363 := Call(__e, PrimFunc(symshen_4_5_1out), W3093)


__e.TailApply(tmp12346, tmp12363)
return


}


}, 1)

tmp12366 := Call(__e, PrimFunc(symshen_4_5ass_6), V3091)


tmp12367 := Call(__e, tmp12345, tmp12366)


__e.TailApply(tmp12328, tmp12367)
return


}, 1)

tmp12368 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12327)


_ = tmp12368

tmp12369 := MakeNative(func(__e *ControlFlow) {
V3104 := __e.Get(1)
_ = V3104
tmp12370 := MakeNative(func(__e *ControlFlow) {
W3105 := __e.Get(1)
_ = W3105
tmp12382 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3105)


if True == tmp12382 {
tmp12371 := MakeNative(func(__e *ControlFlow) {
W3114 := __e.Get(1)
_ = W3114
tmp12373 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3114)


if True == tmp12373 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3114)
return
}


}, 1)

tmp12374 := MakeNative(func(__e *ControlFlow) {
W3115 := __e.Get(1)
_ = W3115
tmp12378 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3115)


if True == tmp12378 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12375 := MakeNative(func(__e *ControlFlow) {
W3116 := __e.Get(1)
_ = W3116
__e.TailApply(PrimFunc(symshen_4comb), W3116, Nil)
return
}, 1)

tmp12376 := Call(__e, PrimFunc(symshen_4in_1_6), W3115)


__e.TailApply(tmp12375, tmp12376)
return


}


}, 1)

tmp12379 := Call(__e, PrimFunc(sym_5e_6), V3104)


tmp12380 := Call(__e, tmp12374, tmp12379)


__e.TailApply(tmp12371, tmp12380)
return


} else {
__e.Return(W3105)
return
}


}, 1)

tmp12383 := MakeNative(func(__e *ControlFlow) {
W3106 := __e.Get(1)
_ = W3106
tmp12404 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3106)


if True == tmp12404 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12384 := MakeNative(func(__e *ControlFlow) {
W3107 := __e.Get(1)
_ = W3107
tmp12385 := MakeNative(func(__e *ControlFlow) {
W3108 := __e.Get(1)
_ = W3108
tmp12386 := MakeNative(func(__e *ControlFlow) {
W3109 := __e.Get(1)
_ = W3109
tmp12399 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3109)


if True == tmp12399 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12387 := MakeNative(func(__e *ControlFlow) {
W3110 := __e.Get(1)
_ = W3110
tmp12388 := MakeNative(func(__e *ControlFlow) {
W3111 := __e.Get(1)
_ = W3111
tmp12395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3111)


if True == tmp12395 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12389 := MakeNative(func(__e *ControlFlow) {
W3112 := __e.Get(1)
_ = W3112
tmp12390 := MakeNative(func(__e *ControlFlow) {
W3113 := __e.Get(1)
_ = W3113
tmp12391 := PrimCons(W3107, W3112)

__e.TailApply(PrimFunc(symshen_4comb), W3113, tmp12391)
return


}, 1)

tmp12392 := Call(__e, PrimFunc(symshen_4in_1_6), W3111)


__e.TailApply(tmp12390, tmp12392)
return


}, 1)

tmp12393 := Call(__e, PrimFunc(symshen_4_5_1out), W3111)


__e.TailApply(tmp12389, tmp12393)
return


}


}, 1)

tmp12396 := Call(__e, PrimFunc(symshen_4_5prems_6), W3110)


__e.TailApply(tmp12388, tmp12396)
return


}, 1)

tmp12397 := Call(__e, PrimFunc(symshen_4in_1_6), W3109)


__e.TailApply(tmp12387, tmp12397)
return


}


}, 1)

tmp12400 := Call(__e, PrimFunc(symshen_4_5sc_6), W3108)


__e.TailApply(tmp12386, tmp12400)
return


}, 1)

tmp12401 := Call(__e, PrimFunc(symshen_4in_1_6), W3106)


__e.TailApply(tmp12385, tmp12401)
return


}, 1)

tmp12402 := Call(__e, PrimFunc(symshen_4_5_1out), W3106)


__e.TailApply(tmp12384, tmp12402)
return


}


}, 1)

tmp12405 := Call(__e, PrimFunc(symshen_4_5prem_6), V3104)


tmp12406 := Call(__e, tmp12383, tmp12405)


__e.TailApply(tmp12370, tmp12406)
return


}, 1)

tmp12407 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12369)


_ = tmp12407

tmp12408 := MakeNative(func(__e *ControlFlow) {
V3117 := __e.Get(1)
_ = V3117
tmp12409 := MakeNative(func(__e *ControlFlow) {
W3118 := __e.Get(1)
_ = W3118
tmp12451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3118)


if True == tmp12451 {
tmp12410 := MakeNative(func(__e *ControlFlow) {
W3120 := __e.Get(1)
_ = W3120
tmp12426 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3120)


if True == tmp12426 {
tmp12411 := MakeNative(func(__e *ControlFlow) {
W3128 := __e.Get(1)
_ = W3128
tmp12413 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3128)


if True == tmp12413 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3128)
return
}


}, 1)

tmp12414 := MakeNative(func(__e *ControlFlow) {
W3129 := __e.Get(1)
_ = W3129
tmp12422 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3129)


if True == tmp12422 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12415 := MakeNative(func(__e *ControlFlow) {
W3130 := __e.Get(1)
_ = W3130
tmp12416 := MakeNative(func(__e *ControlFlow) {
W3131 := __e.Get(1)
_ = W3131
tmp12417 := PrimCons(W3130, Nil)

tmp12418 := PrimCons(Nil, tmp12417)

__e.TailApply(PrimFunc(symshen_4comb), W3131, tmp12418)
return


}, 1)

tmp12419 := Call(__e, PrimFunc(symshen_4in_1_6), W3129)


__e.TailApply(tmp12416, tmp12419)
return


}, 1)

tmp12420 := Call(__e, PrimFunc(symshen_4_5_1out), W3129)


__e.TailApply(tmp12415, tmp12420)
return


}


}, 1)

tmp12423 := Call(__e, PrimFunc(symshen_4_5formula_6), V3117)


tmp12424 := Call(__e, tmp12414, tmp12423)


__e.TailApply(tmp12411, tmp12424)
return


} else {
__e.Return(W3120)
return
}


}, 1)

tmp12427 := MakeNative(func(__e *ControlFlow) {
W3121 := __e.Get(1)
_ = W3121
tmp12447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3121)


if True == tmp12447 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12428 := MakeNative(func(__e *ControlFlow) {
W3122 := __e.Get(1)
_ = W3122
tmp12429 := MakeNative(func(__e *ControlFlow) {
W3123 := __e.Get(1)
_ = W3123
tmp12443 := Call(__e, PrimFunc(symshen_4hds_a_2), W3123, sym_6_6)


if True == tmp12443 {
tmp12430 := MakeNative(func(__e *ControlFlow) {
W3124 := __e.Get(1)
_ = W3124
tmp12431 := MakeNative(func(__e *ControlFlow) {
W3125 := __e.Get(1)
_ = W3125
tmp12439 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3125)


if True == tmp12439 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12432 := MakeNative(func(__e *ControlFlow) {
W3126 := __e.Get(1)
_ = W3126
tmp12433 := MakeNative(func(__e *ControlFlow) {
W3127 := __e.Get(1)
_ = W3127
tmp12434 := PrimCons(W3126, Nil)

tmp12435 := PrimCons(W3122, tmp12434)

__e.TailApply(PrimFunc(symshen_4comb), W3127, tmp12435)
return


}, 1)

tmp12436 := Call(__e, PrimFunc(symshen_4in_1_6), W3125)


__e.TailApply(tmp12433, tmp12436)
return


}, 1)

tmp12437 := Call(__e, PrimFunc(symshen_4_5_1out), W3125)


__e.TailApply(tmp12432, tmp12437)
return


}


}, 1)

tmp12440 := Call(__e, PrimFunc(symshen_4_5formula_6), W3124)


__e.TailApply(tmp12431, tmp12440)
return


}, 1)

tmp12441 := Call(__e, PrimFunc(symtail), W3123)


__e.TailApply(tmp12430, tmp12441)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12444 := Call(__e, PrimFunc(symshen_4in_1_6), W3121)


__e.TailApply(tmp12429, tmp12444)
return


}, 1)

tmp12445 := Call(__e, PrimFunc(symshen_4_5_1out), W3121)


__e.TailApply(tmp12428, tmp12445)
return


}


}, 1)

tmp12448 := Call(__e, PrimFunc(symshen_4_5ass_6), V3117)


tmp12449 := Call(__e, tmp12427, tmp12448)


__e.TailApply(tmp12410, tmp12449)
return


} else {
__e.Return(W3118)
return
}


}, 1)

tmp12457 := Call(__e, PrimFunc(symshen_4hds_a_2), V3117, sym_b)


var ifres12452 Obj

if True == tmp12457 {
tmp12453 := MakeNative(func(__e *ControlFlow) {
W3119 := __e.Get(1)
_ = W3119
__e.TailApply(PrimFunc(symshen_4comb), W3119, sym_b)
return
}, 1)

tmp12454 := Call(__e, PrimFunc(symtail), V3117)


tmp12455 := Call(__e, tmp12453, tmp12454)


ifres12452 = tmp12455


} else {
tmp12456 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12452 = tmp12456


}

__e.TailApply(tmp12409, ifres12452)
return


}, 1)

tmp12458 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12408)


_ = tmp12458

tmp12459 := MakeNative(func(__e *ControlFlow) {
V3132 := __e.Get(1)
_ = V3132
tmp12460 := MakeNative(func(__e *ControlFlow) {
W3133 := __e.Get(1)
_ = W3133
tmp12485 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3133)


if True == tmp12485 {
tmp12461 := MakeNative(func(__e *ControlFlow) {
W3142 := __e.Get(1)
_ = W3142
tmp12473 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3142)


if True == tmp12473 {
tmp12462 := MakeNative(func(__e *ControlFlow) {
W3146 := __e.Get(1)
_ = W3146
tmp12464 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3146)


if True == tmp12464 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3146)
return
}


}, 1)

tmp12465 := MakeNative(func(__e *ControlFlow) {
W3147 := __e.Get(1)
_ = W3147
tmp12469 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3147)


if True == tmp12469 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12466 := MakeNative(func(__e *ControlFlow) {
W3148 := __e.Get(1)
_ = W3148
__e.TailApply(PrimFunc(symshen_4comb), W3148, Nil)
return
}, 1)

tmp12467 := Call(__e, PrimFunc(symshen_4in_1_6), W3147)


__e.TailApply(tmp12466, tmp12467)
return


}


}, 1)

tmp12470 := Call(__e, PrimFunc(sym_5e_6), V3132)


tmp12471 := Call(__e, tmp12465, tmp12470)


__e.TailApply(tmp12462, tmp12471)
return


} else {
__e.Return(W3142)
return
}


}, 1)

tmp12474 := MakeNative(func(__e *ControlFlow) {
W3143 := __e.Get(1)
_ = W3143
tmp12481 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3143)


if True == tmp12481 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12475 := MakeNative(func(__e *ControlFlow) {
W3144 := __e.Get(1)
_ = W3144
tmp12476 := MakeNative(func(__e *ControlFlow) {
W3145 := __e.Get(1)
_ = W3145
tmp12477 := PrimCons(W3144, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3145, tmp12477)
return


}, 1)

tmp12478 := Call(__e, PrimFunc(symshen_4in_1_6), W3143)


__e.TailApply(tmp12476, tmp12478)
return


}, 1)

tmp12479 := Call(__e, PrimFunc(symshen_4_5_1out), W3143)


__e.TailApply(tmp12475, tmp12479)
return


}


}, 1)

tmp12482 := Call(__e, PrimFunc(symshen_4_5formula_6), V3132)


tmp12483 := Call(__e, tmp12474, tmp12482)


__e.TailApply(tmp12461, tmp12483)
return


} else {
__e.Return(W3133)
return
}


}, 1)

tmp12486 := MakeNative(func(__e *ControlFlow) {
W3134 := __e.Get(1)
_ = W3134
tmp12507 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3134)


if True == tmp12507 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12487 := MakeNative(func(__e *ControlFlow) {
W3135 := __e.Get(1)
_ = W3135
tmp12488 := MakeNative(func(__e *ControlFlow) {
W3136 := __e.Get(1)
_ = W3136
tmp12489 := MakeNative(func(__e *ControlFlow) {
W3137 := __e.Get(1)
_ = W3137
tmp12502 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3137)


if True == tmp12502 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12490 := MakeNative(func(__e *ControlFlow) {
W3138 := __e.Get(1)
_ = W3138
tmp12491 := MakeNative(func(__e *ControlFlow) {
W3139 := __e.Get(1)
_ = W3139
tmp12498 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3139)


if True == tmp12498 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12492 := MakeNative(func(__e *ControlFlow) {
W3140 := __e.Get(1)
_ = W3140
tmp12493 := MakeNative(func(__e *ControlFlow) {
W3141 := __e.Get(1)
_ = W3141
tmp12494 := PrimCons(W3135, W3140)

__e.TailApply(PrimFunc(symshen_4comb), W3141, tmp12494)
return


}, 1)

tmp12495 := Call(__e, PrimFunc(symshen_4in_1_6), W3139)


__e.TailApply(tmp12493, tmp12495)
return


}, 1)

tmp12496 := Call(__e, PrimFunc(symshen_4_5_1out), W3139)


__e.TailApply(tmp12492, tmp12496)
return


}


}, 1)

tmp12499 := Call(__e, PrimFunc(symshen_4_5ass_6), W3138)


__e.TailApply(tmp12491, tmp12499)
return


}, 1)

tmp12500 := Call(__e, PrimFunc(symshen_4in_1_6), W3137)


__e.TailApply(tmp12490, tmp12500)
return


}


}, 1)

tmp12503 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W3136)


__e.TailApply(tmp12489, tmp12503)
return


}, 1)

tmp12504 := Call(__e, PrimFunc(symshen_4in_1_6), W3134)


__e.TailApply(tmp12488, tmp12504)
return


}, 1)

tmp12505 := Call(__e, PrimFunc(symshen_4_5_1out), W3134)


__e.TailApply(tmp12487, tmp12505)
return


}


}, 1)

tmp12508 := Call(__e, PrimFunc(symshen_4_5formula_6), V3132)


tmp12509 := Call(__e, tmp12486, tmp12508)


__e.TailApply(tmp12460, tmp12509)
return


}, 1)

tmp12510 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12459)


_ = tmp12510

tmp12511 := MakeNative(func(__e *ControlFlow) {
V3149 := __e.Get(1)
_ = V3149
tmp12512 := MakeNative(func(__e *ControlFlow) {
W3150 := __e.Get(1)
_ = W3150
tmp12514 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3150)


if True == tmp12514 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3150)
return
}


}, 1)

tmp12525 := PrimIsPair(V3149)

var ifres12515 Obj

if True == tmp12525 {
tmp12516 := MakeNative(func(__e *ControlFlow) {
W3151 := __e.Get(1)
_ = W3151
tmp12517 := MakeNative(func(__e *ControlFlow) {
W3152 := __e.Get(1)
_ = W3152
tmp12519 := PrimIntern(MakeString(","))

tmp12520 := PrimEqual(W3151, tmp12519)

if True == tmp12520 {
__e.TailApply(PrimFunc(symshen_4comb), W3152, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12521 := Call(__e, PrimFunc(symtail), V3149)


__e.TailApply(tmp12517, tmp12521)
return


}, 1)

tmp12522 := Call(__e, PrimFunc(symhead), V3149)


tmp12523 := Call(__e, tmp12516, tmp12522)


ifres12515 = tmp12523


} else {
tmp12524 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12515 = tmp12524


}

__e.TailApply(tmp12512, ifres12515)
return


}, 1)

tmp12526 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12511)


_ = tmp12526

tmp12527 := MakeNative(func(__e *ControlFlow) {
V3153 := __e.Get(1)
_ = V3153
tmp12528 := MakeNative(func(__e *ControlFlow) {
W3154 := __e.Get(1)
_ = W3154
tmp12542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3154)


if True == tmp12542 {
tmp12529 := MakeNative(func(__e *ControlFlow) {
W3163 := __e.Get(1)
_ = W3163
tmp12531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3163)


if True == tmp12531 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3163)
return
}


}, 1)

tmp12532 := MakeNative(func(__e *ControlFlow) {
W3164 := __e.Get(1)
_ = W3164
tmp12538 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3164)


if True == tmp12538 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12533 := MakeNative(func(__e *ControlFlow) {
W3165 := __e.Get(1)
_ = W3165
tmp12534 := MakeNative(func(__e *ControlFlow) {
W3166 := __e.Get(1)
_ = W3166
__e.TailApply(PrimFunc(symshen_4comb), W3166, W3165)
return
}, 1)

tmp12535 := Call(__e, PrimFunc(symshen_4in_1_6), W3164)


__e.TailApply(tmp12534, tmp12535)
return


}, 1)

tmp12536 := Call(__e, PrimFunc(symshen_4_5_1out), W3164)


__e.TailApply(tmp12533, tmp12536)
return


}


}, 1)

tmp12539 := Call(__e, PrimFunc(symshen_4_5expr_6), V3153)


tmp12540 := Call(__e, tmp12532, tmp12539)


__e.TailApply(tmp12529, tmp12540)
return


} else {
__e.Return(W3154)
return
}


}, 1)

tmp12543 := MakeNative(func(__e *ControlFlow) {
W3155 := __e.Get(1)
_ = W3155
tmp12569 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3155)


if True == tmp12569 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12544 := MakeNative(func(__e *ControlFlow) {
W3156 := __e.Get(1)
_ = W3156
tmp12545 := MakeNative(func(__e *ControlFlow) {
W3157 := __e.Get(1)
_ = W3157
tmp12546 := MakeNative(func(__e *ControlFlow) {
W3158 := __e.Get(1)
_ = W3158
tmp12564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3158)


if True == tmp12564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12547 := MakeNative(func(__e *ControlFlow) {
W3159 := __e.Get(1)
_ = W3159
tmp12548 := MakeNative(func(__e *ControlFlow) {
W3160 := __e.Get(1)
_ = W3160
tmp12560 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3160)


if True == tmp12560 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12549 := MakeNative(func(__e *ControlFlow) {
W3161 := __e.Get(1)
_ = W3161
tmp12550 := MakeNative(func(__e *ControlFlow) {
W3162 := __e.Get(1)
_ = W3162
tmp12551 := Call(__e, PrimFunc(symshen_4curry), W3156)


tmp12552 := PrimIntern(MakeString(":"))

tmp12553 := Call(__e, PrimFunc(symshen_4rectify_1type), W3161)


tmp12554 := PrimCons(tmp12553, Nil)

tmp12555 := PrimCons(tmp12552, tmp12554)

tmp12556 := PrimCons(tmp12551, tmp12555)

__e.TailApply(PrimFunc(symshen_4comb), W3162, tmp12556)
return


}, 1)

tmp12557 := Call(__e, PrimFunc(symshen_4in_1_6), W3160)


__e.TailApply(tmp12550, tmp12557)
return


}, 1)

tmp12558 := Call(__e, PrimFunc(symshen_4_5_1out), W3160)


__e.TailApply(tmp12549, tmp12558)
return


}


}, 1)

tmp12561 := Call(__e, PrimFunc(symshen_4_5type_6), W3159)


__e.TailApply(tmp12548, tmp12561)
return


}, 1)

tmp12562 := Call(__e, PrimFunc(symshen_4in_1_6), W3158)


__e.TailApply(tmp12547, tmp12562)
return


}


}, 1)

tmp12565 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W3157)


__e.TailApply(tmp12546, tmp12565)
return


}, 1)

tmp12566 := Call(__e, PrimFunc(symshen_4in_1_6), W3155)


__e.TailApply(tmp12545, tmp12566)
return


}, 1)

tmp12567 := Call(__e, PrimFunc(symshen_4_5_1out), W3155)


__e.TailApply(tmp12544, tmp12567)
return


}


}, 1)

tmp12570 := Call(__e, PrimFunc(symshen_4_5expr_6), V3153)


tmp12571 := Call(__e, tmp12543, tmp12570)


__e.TailApply(tmp12528, tmp12571)
return


}, 1)

tmp12572 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12527)


_ = tmp12572

tmp12573 := MakeNative(func(__e *ControlFlow) {
V3167 := __e.Get(1)
_ = V3167
tmp12574 := MakeNative(func(__e *ControlFlow) {
W3168 := __e.Get(1)
_ = W3168
tmp12576 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3168)


if True == tmp12576 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3168)
return
}


}, 1)

tmp12587 := PrimIsPair(V3167)

var ifres12577 Obj

if True == tmp12587 {
tmp12578 := MakeNative(func(__e *ControlFlow) {
W3169 := __e.Get(1)
_ = W3169
tmp12579 := MakeNative(func(__e *ControlFlow) {
W3170 := __e.Get(1)
_ = W3170
tmp12581 := PrimIntern(MakeString(":"))

tmp12582 := PrimEqual(W3169, tmp12581)

if True == tmp12582 {
__e.TailApply(PrimFunc(symshen_4comb), W3170, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12583 := Call(__e, PrimFunc(symtail), V3167)


__e.TailApply(tmp12579, tmp12583)
return


}, 1)

tmp12584 := Call(__e, PrimFunc(symhead), V3167)


tmp12585 := Call(__e, tmp12578, tmp12584)


ifres12577 = tmp12585


} else {
tmp12586 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12577 = tmp12586


}

__e.TailApply(tmp12574, ifres12577)
return


}, 1)

tmp12588 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12573)


_ = tmp12588

tmp12589 := MakeNative(func(__e *ControlFlow) {
V3171 := __e.Get(1)
_ = V3171
tmp12590 := MakeNative(func(__e *ControlFlow) {
W3172 := __e.Get(1)
_ = W3172
tmp12602 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3172)


if True == tmp12602 {
tmp12591 := MakeNative(func(__e *ControlFlow) {
W3179 := __e.Get(1)
_ = W3179
tmp12593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3179)


if True == tmp12593 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3179)
return
}


}, 1)

tmp12594 := MakeNative(func(__e *ControlFlow) {
W3180 := __e.Get(1)
_ = W3180
tmp12598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3180)


if True == tmp12598 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12595 := MakeNative(func(__e *ControlFlow) {
W3181 := __e.Get(1)
_ = W3181
__e.TailApply(PrimFunc(symshen_4comb), W3181, Nil)
return
}, 1)

tmp12596 := Call(__e, PrimFunc(symshen_4in_1_6), W3180)


__e.TailApply(tmp12595, tmp12596)
return


}


}, 1)

tmp12599 := Call(__e, PrimFunc(sym_5e_6), V3171)


tmp12600 := Call(__e, tmp12594, tmp12599)


__e.TailApply(tmp12591, tmp12600)
return


} else {
__e.Return(W3172)
return
}


}, 1)

tmp12603 := MakeNative(func(__e *ControlFlow) {
W3173 := __e.Get(1)
_ = W3173
tmp12618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3173)


if True == tmp12618 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12604 := MakeNative(func(__e *ControlFlow) {
W3174 := __e.Get(1)
_ = W3174
tmp12605 := MakeNative(func(__e *ControlFlow) {
W3175 := __e.Get(1)
_ = W3175
tmp12606 := MakeNative(func(__e *ControlFlow) {
W3176 := __e.Get(1)
_ = W3176
tmp12613 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3176)


if True == tmp12613 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12607 := MakeNative(func(__e *ControlFlow) {
W3177 := __e.Get(1)
_ = W3177
tmp12608 := MakeNative(func(__e *ControlFlow) {
W3178 := __e.Get(1)
_ = W3178
tmp12609 := PrimCons(W3174, W3177)

__e.TailApply(PrimFunc(symshen_4comb), W3178, tmp12609)
return


}, 1)

tmp12610 := Call(__e, PrimFunc(symshen_4in_1_6), W3176)


__e.TailApply(tmp12608, tmp12610)
return


}, 1)

tmp12611 := Call(__e, PrimFunc(symshen_4_5_1out), W3176)


__e.TailApply(tmp12607, tmp12611)
return


}


}, 1)

tmp12614 := Call(__e, PrimFunc(symshen_4_5sides_6), W3175)


__e.TailApply(tmp12606, tmp12614)
return


}, 1)

tmp12615 := Call(__e, PrimFunc(symshen_4in_1_6), W3173)


__e.TailApply(tmp12605, tmp12615)
return


}, 1)

tmp12616 := Call(__e, PrimFunc(symshen_4_5_1out), W3173)


__e.TailApply(tmp12604, tmp12616)
return


}


}, 1)

tmp12619 := Call(__e, PrimFunc(symshen_4_5side_6), V3171)


tmp12620 := Call(__e, tmp12603, tmp12619)


__e.TailApply(tmp12590, tmp12620)
return


}, 1)

tmp12621 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12589)


_ = tmp12621

tmp12622 := MakeNative(func(__e *ControlFlow) {
V3182 := __e.Get(1)
_ = V3182
tmp12623 := MakeNative(func(__e *ControlFlow) {
W3183 := __e.Get(1)
_ = W3183
tmp12673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3183)


if True == tmp12673 {
tmp12624 := MakeNative(func(__e *ControlFlow) {
W3187 := __e.Get(1)
_ = W3187
tmp12650 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3187)


if True == tmp12650 {
tmp12625 := MakeNative(func(__e *ControlFlow) {
W3193 := __e.Get(1)
_ = W3193
tmp12627 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3193)


if True == tmp12627 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3193)
return
}


}, 1)

tmp12648 := Call(__e, PrimFunc(symshen_4hds_a_2), V3182, symshen_4let_b)


var ifres12628 Obj

if True == tmp12648 {
tmp12629 := MakeNative(func(__e *ControlFlow) {
W3194 := __e.Get(1)
_ = W3194
tmp12644 := PrimIsPair(W3194)

if True == tmp12644 {
tmp12630 := MakeNative(func(__e *ControlFlow) {
W3195 := __e.Get(1)
_ = W3195
tmp12631 := MakeNative(func(__e *ControlFlow) {
W3196 := __e.Get(1)
_ = W3196
tmp12640 := PrimIsPair(W3196)

if True == tmp12640 {
tmp12632 := MakeNative(func(__e *ControlFlow) {
W3197 := __e.Get(1)
_ = W3197
tmp12633 := MakeNative(func(__e *ControlFlow) {
W3198 := __e.Get(1)
_ = W3198
tmp12634 := PrimCons(W3197, Nil)

tmp12635 := PrimCons(W3195, tmp12634)

tmp12636 := PrimCons(symshen_4let_b, tmp12635)

__e.TailApply(PrimFunc(symshen_4comb), W3198, tmp12636)
return


}, 1)

tmp12637 := Call(__e, PrimFunc(symtail), W3196)


__e.TailApply(tmp12633, tmp12637)
return


}, 1)

tmp12638 := Call(__e, PrimFunc(symhead), W3196)


__e.TailApply(tmp12632, tmp12638)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12641 := Call(__e, PrimFunc(symtail), W3194)


__e.TailApply(tmp12631, tmp12641)
return


}, 1)

tmp12642 := Call(__e, PrimFunc(symhead), W3194)


__e.TailApply(tmp12630, tmp12642)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12645 := Call(__e, PrimFunc(symtail), V3182)


tmp12646 := Call(__e, tmp12629, tmp12645)


ifres12628 = tmp12646


} else {
tmp12647 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12628 = tmp12647


}

__e.TailApply(tmp12625, ifres12628)
return


} else {
__e.Return(W3187)
return
}


}, 1)

tmp12671 := Call(__e, PrimFunc(symshen_4hds_a_2), V3182, symlet)


var ifres12651 Obj

if True == tmp12671 {
tmp12652 := MakeNative(func(__e *ControlFlow) {
W3188 := __e.Get(1)
_ = W3188
tmp12667 := PrimIsPair(W3188)

if True == tmp12667 {
tmp12653 := MakeNative(func(__e *ControlFlow) {
W3189 := __e.Get(1)
_ = W3189
tmp12654 := MakeNative(func(__e *ControlFlow) {
W3190 := __e.Get(1)
_ = W3190
tmp12663 := PrimIsPair(W3190)

if True == tmp12663 {
tmp12655 := MakeNative(func(__e *ControlFlow) {
W3191 := __e.Get(1)
_ = W3191
tmp12656 := MakeNative(func(__e *ControlFlow) {
W3192 := __e.Get(1)
_ = W3192
tmp12657 := PrimCons(W3191, Nil)

tmp12658 := PrimCons(W3189, tmp12657)

tmp12659 := PrimCons(symlet, tmp12658)

__e.TailApply(PrimFunc(symshen_4comb), W3192, tmp12659)
return


}, 1)

tmp12660 := Call(__e, PrimFunc(symtail), W3190)


__e.TailApply(tmp12656, tmp12660)
return


}, 1)

tmp12661 := Call(__e, PrimFunc(symhead), W3190)


__e.TailApply(tmp12655, tmp12661)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12664 := Call(__e, PrimFunc(symtail), W3188)


__e.TailApply(tmp12654, tmp12664)
return


}, 1)

tmp12665 := Call(__e, PrimFunc(symhead), W3188)


__e.TailApply(tmp12653, tmp12665)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12668 := Call(__e, PrimFunc(symtail), V3182)


tmp12669 := Call(__e, tmp12652, tmp12668)


ifres12651 = tmp12669


} else {
tmp12670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12651 = tmp12670


}

__e.TailApply(tmp12624, ifres12651)
return


} else {
__e.Return(W3183)
return
}


}, 1)

tmp12687 := Call(__e, PrimFunc(symshen_4hds_a_2), V3182, symif)


var ifres12674 Obj

if True == tmp12687 {
tmp12675 := MakeNative(func(__e *ControlFlow) {
W3184 := __e.Get(1)
_ = W3184
tmp12683 := PrimIsPair(W3184)

if True == tmp12683 {
tmp12676 := MakeNative(func(__e *ControlFlow) {
W3185 := __e.Get(1)
_ = W3185
tmp12677 := MakeNative(func(__e *ControlFlow) {
W3186 := __e.Get(1)
_ = W3186
tmp12678 := PrimCons(W3185, Nil)

tmp12679 := PrimCons(symif, tmp12678)

__e.TailApply(PrimFunc(symshen_4comb), W3186, tmp12679)
return


}, 1)

tmp12680 := Call(__e, PrimFunc(symtail), W3184)


__e.TailApply(tmp12677, tmp12680)
return


}, 1)

tmp12681 := Call(__e, PrimFunc(symhead), W3184)


__e.TailApply(tmp12676, tmp12681)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12684 := Call(__e, PrimFunc(symtail), V3182)


tmp12685 := Call(__e, tmp12675, tmp12684)


ifres12674 = tmp12685


} else {
tmp12686 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12674 = tmp12686


}

__e.TailApply(tmp12623, ifres12674)
return


}, 1)

tmp12688 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12622)


_ = tmp12688

tmp12689 := MakeNative(func(__e *ControlFlow) {
V3205 := __e.Get(1)
_ = V3205
V3206 := __e.Get(2)
_ = V3206
V3207 := __e.Get(3)
_ = V3207
tmp12724 := PrimIsPair(V3207)

var ifres12711 Obj

if True == tmp12724 {
tmp12722 := PrimHead(V3207)

tmp12723 := PrimEqual(Nil, tmp12722)

var ifres12713 Obj

if True == tmp12723 {
tmp12720 := PrimTail(V3207)

tmp12721 := PrimIsPair(tmp12720)

var ifres12715 Obj

if True == tmp12721 {
tmp12717 := PrimTail(V3207)

tmp12718 := PrimTail(tmp12717)

tmp12719 := PrimEqual(Nil, tmp12718)

var ifres12716 Obj

if True == tmp12719 {
ifres12716 = True


} else {
ifres12716 = False


}

ifres12715 = ifres12716


} else {
ifres12715 = False


}

var ifres12714 Obj

if True == ifres12715 {
ifres12714 = True


} else {
ifres12714 = False


}

ifres12713 = ifres12714


} else {
ifres12713 = False


}

var ifres12712 Obj

if True == ifres12713 {
ifres12712 = True


} else {
ifres12712 = False


}

ifres12711 = ifres12712


} else {
ifres12711 = False


}

if True == ifres12711 {
tmp12690 := MakeNative(func(__e *ControlFlow) {
W3208 := __e.Get(1)
_ = W3208
tmp12691 := MakeNative(func(__e *ControlFlow) {
W3209 := __e.Get(1)
_ = W3209
tmp12692 := MakeNative(func(__e *ControlFlow) {
W3210 := __e.Get(1)
_ = W3210
tmp12693 := MakeNative(func(__e *ControlFlow) {
W3211 := __e.Get(1)
_ = W3211
tmp12694 := MakeNative(func(__e *ControlFlow) {
W3212 := __e.Get(1)
_ = W3212
tmp12695 := PrimCons(W3211, Nil)

__e.Return(PrimCons(W3212, tmp12695))
return


}, 1)

tmp12696 := PrimCons(V3207, Nil)

tmp12697 := PrimCons(V3206, tmp12696)

tmp12698 := PrimCons(V3205, tmp12697)

__e.TailApply(tmp12694, tmp12698)
return


}, 1)

tmp12699 := PrimCons(W3210, Nil)

tmp12700 := PrimCons(W3209, Nil)

tmp12701 := PrimCons(tmp12699, tmp12700)

tmp12702 := PrimCons(V3205, tmp12701)

__e.TailApply(tmp12693, tmp12702)
return


}, 1)

tmp12703 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3206)


tmp12704 := PrimCons(W3208, Nil)

tmp12705 := PrimCons(tmp12703, tmp12704)

__e.TailApply(tmp12692, tmp12705)
return


}, 1)

tmp12706 := PrimTail(V3207)

tmp12707 := PrimCons(W3208, Nil)

tmp12708 := PrimCons(tmp12706, tmp12707)

__e.TailApply(tmp12691, tmp12708)
return


}, 1)

tmp12709 := Call(__e, PrimFunc(symgensym), symP)


__e.TailApply(tmp12690, tmp12709)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
return
}


}, 3)

tmp12725 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12689)


_ = tmp12725

tmp12726 := MakeNative(func(__e *ControlFlow) {
V3215 := __e.Get(1)
_ = V3215
tmp12755 := PrimEqual(Nil, V3215)

if True == tmp12755 {
__e.Return(Nil)
return
} else {
tmp12753 := PrimIsPair(V3215)

var ifres12733 Obj

if True == tmp12753 {
tmp12751 := PrimHead(V3215)

tmp12752 := PrimIsPair(tmp12751)

var ifres12735 Obj

if True == tmp12752 {
tmp12748 := PrimHead(V3215)

tmp12749 := PrimHead(tmp12748)

tmp12750 := PrimEqual(Nil, tmp12749)

var ifres12737 Obj

if True == tmp12750 {
tmp12745 := PrimHead(V3215)

tmp12746 := PrimTail(tmp12745)

tmp12747 := PrimIsPair(tmp12746)

var ifres12739 Obj

if True == tmp12747 {
tmp12741 := PrimHead(V3215)

tmp12742 := PrimTail(tmp12741)

tmp12743 := PrimTail(tmp12742)

tmp12744 := PrimEqual(Nil, tmp12743)

var ifres12740 Obj

if True == tmp12744 {
ifres12740 = True


} else {
ifres12740 = False


}

ifres12739 = ifres12740


} else {
ifres12739 = False


}

var ifres12738 Obj

if True == ifres12739 {
ifres12738 = True


} else {
ifres12738 = False


}

ifres12737 = ifres12738


} else {
ifres12737 = False


}

var ifres12736 Obj

if True == ifres12737 {
ifres12736 = True


} else {
ifres12736 = False


}

ifres12735 = ifres12736


} else {
ifres12735 = False


}

var ifres12734 Obj

if True == ifres12735 {
ifres12734 = True


} else {
ifres12734 = False


}

ifres12733 = ifres12734


} else {
ifres12733 = False


}

if True == ifres12733 {
tmp12727 := PrimHead(V3215)

tmp12728 := PrimTail(tmp12727)

tmp12729 := PrimHead(tmp12728)

tmp12730 := PrimTail(V3215)

tmp12731 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp12730)


__e.Return(PrimCons(tmp12729, tmp12731))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
return
}


}


}, 1)

tmp12756 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp12726)


_ = tmp12756

tmp12757 := MakeNative(func(__e *ControlFlow) {
V3216 := __e.Get(1)
_ = V3216
tmp12758 := MakeNative(func(__e *ControlFlow) {
W3217 := __e.Get(1)
_ = W3217
tmp12760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3217)


if True == tmp12760 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3217)
return
}


}, 1)

tmp12772 := PrimIsPair(V3216)

var ifres12761 Obj

if True == tmp12772 {
tmp12762 := MakeNative(func(__e *ControlFlow) {
W3218 := __e.Get(1)
_ = W3218
tmp12763 := MakeNative(func(__e *ControlFlow) {
W3219 := __e.Get(1)
_ = W3219
tmp12766 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W3218)


tmp12767 := PrimNot(tmp12766)

if True == tmp12767 {
tmp12764 := Call(__e, PrimFunc(symmacroexpand), W3218)


__e.TailApply(PrimFunc(symshen_4comb), W3219, tmp12764)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12768 := Call(__e, PrimFunc(symtail), V3216)


__e.TailApply(tmp12763, tmp12768)
return


}, 1)

tmp12769 := Call(__e, PrimFunc(symhead), V3216)


tmp12770 := Call(__e, tmp12762, tmp12769)


ifres12761 = tmp12770


} else {
tmp12771 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12761 = tmp12771


}

__e.TailApply(tmp12758, ifres12761)
return


}, 1)

tmp12773 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12757)


_ = tmp12773

tmp12774 := MakeNative(func(__e *ControlFlow) {
V3220 := __e.Get(1)
_ = V3220
tmp12781 := PrimIntern(MakeString(";"))

tmp12782 := PrimIntern(MakeString(","))

tmp12783 := PrimIntern(MakeString(":"))

tmp12784 := PrimCons(sym_5_1_1, Nil)

tmp12785 := PrimCons(tmp12783, tmp12784)

tmp12786 := PrimCons(tmp12782, tmp12785)

tmp12787 := PrimCons(tmp12781, tmp12786)

tmp12788 := PrimCons(sym_6_6, tmp12787)

tmp12789 := Call(__e, PrimFunc(symelement_2), V3220, tmp12788)


if True == tmp12789 {
__e.Return(True)
return
} else {
tmp12779 := Call(__e, PrimFunc(symshen_4sng_2), V3220)


var ifres12776 Obj

if True == tmp12779 {
ifres12776 = True


} else {
tmp12778 := Call(__e, PrimFunc(symshen_4dbl_2), V3220)


var ifres12777 Obj

if True == tmp12778 {
ifres12777 = True


} else {
ifres12777 = False


}

ifres12776 = ifres12777


}

if True == ifres12776 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12790 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12774)


_ = tmp12790

tmp12791 := MakeNative(func(__e *ControlFlow) {
V3221 := __e.Get(1)
_ = V3221
tmp12792 := MakeNative(func(__e *ControlFlow) {
W3222 := __e.Get(1)
_ = W3222
tmp12794 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3222)


if True == tmp12794 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3222)
return
}


}, 1)

tmp12795 := MakeNative(func(__e *ControlFlow) {
W3223 := __e.Get(1)
_ = W3223
tmp12801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3223)


if True == tmp12801 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12796 := MakeNative(func(__e *ControlFlow) {
W3224 := __e.Get(1)
_ = W3224
tmp12797 := MakeNative(func(__e *ControlFlow) {
W3225 := __e.Get(1)
_ = W3225
__e.TailApply(PrimFunc(symshen_4comb), W3225, W3224)
return
}, 1)

tmp12798 := Call(__e, PrimFunc(symshen_4in_1_6), W3223)


__e.TailApply(tmp12797, tmp12798)
return


}, 1)

tmp12799 := Call(__e, PrimFunc(symshen_4_5_1out), W3223)


__e.TailApply(tmp12796, tmp12799)
return


}


}, 1)

tmp12802 := Call(__e, PrimFunc(symshen_4_5expr_6), V3221)


tmp12803 := Call(__e, tmp12795, tmp12802)


__e.TailApply(tmp12792, tmp12803)
return


}, 1)

tmp12804 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12791)


_ = tmp12804

tmp12805 := MakeNative(func(__e *ControlFlow) {
V3226 := __e.Get(1)
_ = V3226
tmp12806 := MakeNative(func(__e *ControlFlow) {
W3227 := __e.Get(1)
_ = W3227
tmp12808 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3227)


if True == tmp12808 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3227)
return
}


}, 1)

tmp12818 := PrimIsPair(V3226)

var ifres12809 Obj

if True == tmp12818 {
tmp12810 := MakeNative(func(__e *ControlFlow) {
W3228 := __e.Get(1)
_ = W3228
tmp12811 := MakeNative(func(__e *ControlFlow) {
W3229 := __e.Get(1)
_ = W3229
tmp12813 := Call(__e, PrimFunc(symshen_4dbl_2), W3228)


if True == tmp12813 {
__e.TailApply(PrimFunc(symshen_4comb), W3229, W3228)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12814 := Call(__e, PrimFunc(symtail), V3226)


__e.TailApply(tmp12811, tmp12814)
return


}, 1)

tmp12815 := Call(__e, PrimFunc(symhead), V3226)


tmp12816 := Call(__e, tmp12810, tmp12815)


ifres12809 = tmp12816


} else {
tmp12817 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12809 = tmp12817


}

__e.TailApply(tmp12806, ifres12809)
return


}, 1)

tmp12819 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12805)


_ = tmp12819

tmp12820 := MakeNative(func(__e *ControlFlow) {
V3230 := __e.Get(1)
_ = V3230
tmp12821 := MakeNative(func(__e *ControlFlow) {
W3231 := __e.Get(1)
_ = W3231
tmp12823 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3231)


if True == tmp12823 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3231)
return
}


}, 1)

tmp12833 := PrimIsPair(V3230)

var ifres12824 Obj

if True == tmp12833 {
tmp12825 := MakeNative(func(__e *ControlFlow) {
W3232 := __e.Get(1)
_ = W3232
tmp12826 := MakeNative(func(__e *ControlFlow) {
W3233 := __e.Get(1)
_ = W3233
tmp12828 := Call(__e, PrimFunc(symshen_4sng_2), W3232)


if True == tmp12828 {
__e.TailApply(PrimFunc(symshen_4comb), W3233, W3232)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12829 := Call(__e, PrimFunc(symtail), V3230)


__e.TailApply(tmp12826, tmp12829)
return


}, 1)

tmp12830 := Call(__e, PrimFunc(symhead), V3230)


tmp12831 := Call(__e, tmp12825, tmp12830)


ifres12824 = tmp12831


} else {
tmp12832 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12824 = tmp12832


}

__e.TailApply(tmp12821, ifres12824)
return


}, 1)

tmp12834 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12820)


_ = tmp12834

tmp12835 := MakeNative(func(__e *ControlFlow) {
V3234 := __e.Get(1)
_ = V3234
tmp12840 := PrimIsSymbol(V3234)

if True == tmp12840 {
tmp12837 := PrimStr(V3234)

tmp12838 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp12837)


if True == tmp12838 {
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

tmp12841 := Call(__e, ns2_1set, symshen_4sng_2, tmp12835)


_ = tmp12841

tmp12842 := MakeNative(func(__e *ControlFlow) {
V3237 := __e.Get(1)
_ = V3237
tmp12851 := PrimEqual(MakeString("___"), V3237)

if True == tmp12851 {
__e.Return(True)
return
} else {
tmp12849 := Call(__e, PrimFunc(symshen_4_7string_2), V3237)


var ifres12845 Obj

if True == tmp12849 {
tmp12847 := Call(__e, PrimFunc(symhdstr), V3237)


tmp12848 := PrimEqual(MakeString("_"), tmp12847)

var ifres12846 Obj

if True == tmp12848 {
ifres12846 = True


} else {
ifres12846 = False


}

ifres12845 = ifres12846


} else {
ifres12845 = False


}

if True == ifres12845 {
tmp12843 := PrimTailString(V3237)

__e.TailApply(PrimFunc(symshen_4sng_1h_2), tmp12843)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12852 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp12842)


_ = tmp12852

tmp12853 := MakeNative(func(__e *ControlFlow) {
V3238 := __e.Get(1)
_ = V3238
tmp12858 := PrimIsSymbol(V3238)

if True == tmp12858 {
tmp12855 := PrimStr(V3238)

tmp12856 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp12855)


if True == tmp12856 {
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

tmp12859 := Call(__e, ns2_1set, symshen_4dbl_2, tmp12853)


_ = tmp12859

tmp12860 := MakeNative(func(__e *ControlFlow) {
V3241 := __e.Get(1)
_ = V3241
tmp12869 := PrimEqual(MakeString("==="), V3241)

if True == tmp12869 {
__e.Return(True)
return
} else {
tmp12867 := Call(__e, PrimFunc(symshen_4_7string_2), V3241)


var ifres12863 Obj

if True == tmp12867 {
tmp12865 := Call(__e, PrimFunc(symhdstr), V3241)


tmp12866 := PrimEqual(MakeString("="), tmp12865)

var ifres12864 Obj

if True == tmp12866 {
ifres12864 = True


} else {
ifres12864 = False


}

ifres12863 = ifres12864


} else {
ifres12863 = False


}

if True == ifres12863 {
tmp12861 := PrimTailString(V3241)

__e.TailApply(PrimFunc(symshen_4dbl_1h_2), tmp12861)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12870 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp12860)


_ = tmp12870

tmp12871 := MakeNative(func(__e *ControlFlow) {
V3242 := __e.Get(1)
_ = V3242
V3243 := __e.Get(2)
_ = V3243
tmp12872 := PrimValue(symshen_4_ddatatypes_d)

tmp12873 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3242, V3243, tmp12872)


tmp12874 := PrimSet(symshen_4_ddatatypes_d, tmp12873)

_ = tmp12874

tmp12875 := PrimValue(symshen_4_dalldatatypes_d)

tmp12876 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3242, V3243, tmp12875)


tmp12877 := PrimSet(symshen_4_dalldatatypes_d, tmp12876)

_ = tmp12877

__e.Return(V3242)
return


}, 2)

tmp12878 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12871)


_ = tmp12878

tmp12879 := MakeNative(func(__e *ControlFlow) {
V3244 := __e.Get(1)
_ = V3244
V3245 := __e.Get(2)
_ = V3245
tmp12880 := MakeNative(func(__e *ControlFlow) {
W3246 := __e.Get(1)
_ = W3246
tmp12881 := PrimCons(V3244, W3246)

tmp12882 := PrimCons(symdefprolog, tmp12881)

__e.TailApply(PrimFunc(symeval), tmp12882)
return


}, 1)

tmp12883 := MakeNative(func(__e *ControlFlow) {
Z3247 := __e.Get(1)
_ = Z3247
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3247)
return
}, 1)

tmp12884 := Call(__e, PrimFunc(symmapcan), tmp12883, V3245)


__e.TailApply(tmp12880, tmp12884)
return


}, 2)

tmp12885 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12879)


_ = tmp12885

tmp12886 := MakeNative(func(__e *ControlFlow) {
V3250 := __e.Get(1)
_ = V3250
tmp12963 := PrimIsPair(V3250)

var ifres12927 Obj

if True == tmp12963 {
tmp12961 := PrimTail(V3250)

tmp12962 := PrimIsPair(tmp12961)

var ifres12929 Obj

if True == tmp12962 {
tmp12958 := PrimTail(V3250)

tmp12959 := PrimTail(tmp12958)

tmp12960 := PrimIsPair(tmp12959)

var ifres12931 Obj

if True == tmp12960 {
tmp12954 := PrimTail(V3250)

tmp12955 := PrimTail(tmp12954)

tmp12956 := PrimHead(tmp12955)

tmp12957 := PrimIsPair(tmp12956)

var ifres12933 Obj

if True == tmp12957 {
tmp12949 := PrimTail(V3250)

tmp12950 := PrimTail(tmp12949)

tmp12951 := PrimHead(tmp12950)

tmp12952 := PrimTail(tmp12951)

tmp12953 := PrimIsPair(tmp12952)

var ifres12935 Obj

if True == tmp12953 {
tmp12943 := PrimTail(V3250)

tmp12944 := PrimTail(tmp12943)

tmp12945 := PrimHead(tmp12944)

tmp12946 := PrimTail(tmp12945)

tmp12947 := PrimTail(tmp12946)

tmp12948 := PrimEqual(Nil, tmp12947)

var ifres12937 Obj

if True == tmp12948 {
tmp12939 := PrimTail(V3250)

tmp12940 := PrimTail(tmp12939)

tmp12941 := PrimTail(tmp12940)

tmp12942 := PrimEqual(Nil, tmp12941)

var ifres12938 Obj

if True == tmp12942 {
ifres12938 = True


} else {
ifres12938 = False


}

ifres12937 = ifres12938


} else {
ifres12937 = False


}

var ifres12936 Obj

if True == ifres12937 {
ifres12936 = True


} else {
ifres12936 = False


}

ifres12935 = ifres12936


} else {
ifres12935 = False


}

var ifres12934 Obj

if True == ifres12935 {
ifres12934 = True


} else {
ifres12934 = False


}

ifres12933 = ifres12934


} else {
ifres12933 = False


}

var ifres12932 Obj

if True == ifres12933 {
ifres12932 = True


} else {
ifres12932 = False


}

ifres12931 = ifres12932


} else {
ifres12931 = False


}

var ifres12930 Obj

if True == ifres12931 {
ifres12930 = True


} else {
ifres12930 = False


}

ifres12929 = ifres12930


} else {
ifres12929 = False


}

var ifres12928 Obj

if True == ifres12929 {
ifres12928 = True


} else {
ifres12928 = False


}

ifres12927 = ifres12928


} else {
ifres12927 = False


}

if True == ifres12927 {
tmp12887 := MakeNative(func(__e *ControlFlow) {
W3251 := __e.Get(1)
_ = W3251
tmp12888 := MakeNative(func(__e *ControlFlow) {
W3252 := __e.Get(1)
_ = W3252
tmp12889 := MakeNative(func(__e *ControlFlow) {
W3253 := __e.Get(1)
_ = W3253
tmp12890 := MakeNative(func(__e *ControlFlow) {
W3254 := __e.Get(1)
_ = W3254
tmp12891 := MakeNative(func(__e *ControlFlow) {
W3255 := __e.Get(1)
_ = W3255
tmp12892 := PrimCons(sym_5_1_1, Nil)

tmp12893 := PrimIntern(MakeString(";"))

tmp12894 := PrimCons(tmp12893, Nil)

tmp12895 := Call(__e, PrimFunc(symappend), W3255, tmp12894)


tmp12896 := Call(__e, PrimFunc(symappend), tmp12892, tmp12895)


__e.TailApply(PrimFunc(symappend), W3254, tmp12896)
return


}, 1)

tmp12897 := PrimTail(V3250)

tmp12898 := PrimTail(tmp12897)

tmp12899 := PrimHead(tmp12898)

tmp12900 := PrimHead(tmp12899)

tmp12901 := PrimHead(V3250)

tmp12902 := PrimTail(V3250)

tmp12903 := PrimHead(tmp12902)

tmp12904 := Call(__e, PrimFunc(symshen_4goals), W3251, tmp12900, tmp12901, tmp12903, W3252, W3253)


__e.TailApply(tmp12891, tmp12904)
return


}, 1)

tmp12905 := PrimTail(V3250)

tmp12906 := PrimTail(tmp12905)

tmp12907 := PrimHead(tmp12906)

tmp12908 := PrimTail(tmp12907)

tmp12909 := PrimHead(tmp12908)

tmp12910 := Call(__e, PrimFunc(symshen_4compile_1consequent), tmp12909, W3252)


__e.TailApply(tmp12890, tmp12910)
return


}, 1)

tmp12911 := PrimTail(V3250)

tmp12912 := PrimTail(tmp12911)

tmp12913 := PrimHead(tmp12912)

tmp12914 := PrimTail(tmp12913)

tmp12915 := PrimHead(tmp12914)

tmp12916 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp12915)


__e.TailApply(tmp12889, tmp12916)
return


}, 1)

tmp12917 := PrimTail(V3250)

tmp12918 := PrimTail(tmp12917)

tmp12919 := PrimHead(tmp12918)

tmp12920 := PrimHead(tmp12919)

tmp12921 := Call(__e, PrimFunc(symlength), tmp12920)


tmp12922 := Call(__e, PrimFunc(symshen_4nvars), tmp12921)


tmp12923 := PrimCons(symDelta, Nil)

tmp12924 := Call(__e, PrimFunc(symappend), tmp12922, tmp12923)


__e.TailApply(tmp12888, tmp12924)
return


}, 1)

tmp12925 := Call(__e, PrimFunc(symshen_4extract_1vars), V3250)


__e.TailApply(tmp12887, tmp12925)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rule->clause")))
return
}


}, 1)

tmp12964 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp12886)


_ = tmp12964

tmp12965 := MakeNative(func(__e *ControlFlow) {
V3262 := __e.Get(1)
_ = V3262
V3263 := __e.Get(2)
_ = V3263
tmp12970 := PrimIsPair(V3263)

if True == tmp12970 {
tmp12966 := Call(__e, PrimFunc(symshen_4optimise_1typing), V3262)


tmp12967 := PrimHead(V3263)

tmp12968 := PrimCons(tmp12967, Nil)

__e.Return(PrimCons(tmp12966, tmp12968))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-consequent")))
return
}


}, 2)

tmp12971 := Call(__e, ns2_1set, symshen_4compile_1consequent, tmp12965)


_ = tmp12971

tmp12972 := MakeNative(func(__e *ControlFlow) {
V3264 := __e.Get(1)
_ = V3264
tmp12977 := PrimEqual(MakeNumber(0), V3264)

if True == tmp12977 {
__e.Return(Nil)
return
} else {
tmp12973 := Call(__e, PrimFunc(symgensym), symV)


tmp12974 := PrimNumberSubtract(V3264, MakeNumber(1))

tmp12975 := Call(__e, PrimFunc(symshen_4nvars), tmp12974)


__e.Return(PrimCons(tmp12973, tmp12975))
return


}


}, 1)

tmp12978 := Call(__e, ns2_1set, symshen_4nvars, tmp12972)


_ = tmp12978

tmp12979 := MakeNative(func(__e *ControlFlow) {
V3265 := __e.Get(1)
_ = V3265
tmp13019 := PrimIsPair(V3265)

var ifres12998 Obj

if True == tmp13019 {
tmp13017 := PrimTail(V3265)

tmp13018 := PrimIsPair(tmp13017)

var ifres13000 Obj

if True == tmp13018 {
tmp13014 := PrimTail(V3265)

tmp13015 := PrimTail(tmp13014)

tmp13016 := PrimIsPair(tmp13015)

var ifres13002 Obj

if True == tmp13016 {
tmp13010 := PrimTail(V3265)

tmp13011 := PrimTail(tmp13010)

tmp13012 := PrimTail(tmp13011)

tmp13013 := PrimEqual(Nil, tmp13012)

var ifres13004 Obj

if True == tmp13013 {
tmp13006 := PrimTail(V3265)

tmp13007 := PrimHead(tmp13006)

tmp13008 := PrimIntern(MakeString(":"))

tmp13009 := PrimEqual(tmp13007, tmp13008)

var ifres13005 Obj

if True == tmp13009 {
ifres13005 = True


} else {
ifres13005 = False


}

ifres13004 = ifres13005


} else {
ifres13004 = False


}

var ifres13003 Obj

if True == ifres13004 {
ifres13003 = True


} else {
ifres13003 = False


}

ifres13002 = ifres13003


} else {
ifres13002 = False


}

var ifres13001 Obj

if True == ifres13002 {
ifres13001 = True


} else {
ifres13001 = False


}

ifres13000 = ifres13001


} else {
ifres13000 = False


}

var ifres12999 Obj

if True == ifres13000 {
ifres12999 = True


} else {
ifres12999 = False


}

ifres12998 = ifres12999


} else {
ifres12998 = False


}

if True == ifres12998 {
tmp12980 := MakeNative(func(__e *ControlFlow) {
W3266 := __e.Get(1)
_ = W3266
__e.TailApply(PrimFunc(symshen_4cons_1form_1with_1modes), W3266)
return
}, 1)

tmp12981 := PrimHead(V3265)

tmp12982 := PrimTail(V3265)

tmp12983 := PrimHead(tmp12982)

tmp12984 := PrimTail(V3265)

tmp12985 := PrimTail(tmp12984)

tmp12986 := PrimCons(sym_7, tmp12985)

tmp12987 := PrimCons(tmp12986, Nil)

tmp12988 := PrimCons(tmp12983, tmp12987)

tmp12989 := PrimCons(tmp12981, tmp12988)

tmp12990 := PrimCons(tmp12989, Nil)

tmp12991 := PrimCons(sym_1, tmp12990)

tmp12992 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp12991)


__e.TailApply(tmp12980, tmp12992)
return


} else {
tmp12993 := MakeNative(func(__e *ControlFlow) {
W3267 := __e.Get(1)
_ = W3267
__e.TailApply(PrimFunc(symshen_4cons_1form_1with_1modes), W3267)
return
}, 1)

tmp12994 := PrimCons(V3265, Nil)

tmp12995 := PrimCons(sym_7, tmp12994)

tmp12996 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp12995)


__e.TailApply(tmp12993, tmp12996)
return


}


}, 1)

tmp13020 := Call(__e, ns2_1set, symshen_4optimise_1typing, tmp12979)


_ = tmp13020

tmp13021 := MakeNative(func(__e *ControlFlow) {
V3268 := __e.Get(1)
_ = V3268
tmp13064 := PrimIsPair(V3268)

var ifres13051 Obj

if True == tmp13064 {
tmp13062 := PrimHead(V3268)

tmp13063 := PrimEqual(sym_7, tmp13062)

var ifres13053 Obj

if True == tmp13063 {
tmp13060 := PrimTail(V3268)

tmp13061 := PrimIsPair(tmp13060)

var ifres13055 Obj

if True == tmp13061 {
tmp13057 := PrimTail(V3268)

tmp13058 := PrimTail(tmp13057)

tmp13059 := PrimEqual(Nil, tmp13058)

var ifres13056 Obj

if True == tmp13059 {
ifres13056 = True


} else {
ifres13056 = False


}

ifres13055 = ifres13056


} else {
ifres13055 = False


}

var ifres13054 Obj

if True == ifres13055 {
ifres13054 = True


} else {
ifres13054 = False


}

ifres13053 = ifres13054


} else {
ifres13053 = False


}

var ifres13052 Obj

if True == ifres13053 {
ifres13052 = True


} else {
ifres13052 = False


}

ifres13051 = ifres13052


} else {
ifres13051 = False


}

if True == ifres13051 {
tmp13022 := PrimTail(V3268)

tmp13023 := PrimHead(tmp13022)

tmp13024 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp13023)


tmp13025 := PrimCons(sym_7, Nil)

tmp13026 := PrimCons(tmp13024, tmp13025)

__e.Return(PrimCons(symmode, tmp13026))
return


} else {
tmp13049 := PrimIsPair(V3268)

var ifres13036 Obj

if True == tmp13049 {
tmp13047 := PrimHead(V3268)

tmp13048 := PrimEqual(sym_1, tmp13047)

var ifres13038 Obj

if True == tmp13048 {
tmp13045 := PrimTail(V3268)

tmp13046 := PrimIsPair(tmp13045)

var ifres13040 Obj

if True == tmp13046 {
tmp13042 := PrimTail(V3268)

tmp13043 := PrimTail(tmp13042)

tmp13044 := PrimEqual(Nil, tmp13043)

var ifres13041 Obj

if True == tmp13044 {
ifres13041 = True


} else {
ifres13041 = False


}

ifres13040 = ifres13041


} else {
ifres13040 = False


}

var ifres13039 Obj

if True == ifres13040 {
ifres13039 = True


} else {
ifres13039 = False


}

ifres13038 = ifres13039


} else {
ifres13038 = False


}

var ifres13037 Obj

if True == ifres13038 {
ifres13037 = True


} else {
ifres13037 = False


}

ifres13036 = ifres13037


} else {
ifres13036 = False


}

if True == ifres13036 {
tmp13027 := PrimTail(V3268)

tmp13028 := PrimHead(tmp13027)

tmp13029 := Call(__e, PrimFunc(symshen_4expand_1mode_1forms), tmp13028)


tmp13030 := PrimCons(sym_1, Nil)

tmp13031 := PrimCons(tmp13029, tmp13030)

__e.Return(PrimCons(symmode, tmp13031))
return


} else {
tmp13034 := PrimIsPair(V3268)

if True == tmp13034 {
tmp13032 := MakeNative(func(__e *ControlFlow) {
Z3269 := __e.Get(1)
_ = Z3269
__e.TailApply(PrimFunc(symshen_4expand_1mode_1forms), Z3269)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13032, V3268)
return


} else {
__e.Return(V3268)
return
}


}


}


}, 1)

tmp13065 := Call(__e, ns2_1set, symshen_4expand_1mode_1forms, tmp13021)


_ = tmp13065

tmp13066 := MakeNative(func(__e *ControlFlow) {
V3270 := __e.Get(1)
_ = V3270
tmp13118 := PrimIsPair(V3270)

var ifres13099 Obj

if True == tmp13118 {
tmp13116 := PrimHead(V3270)

tmp13117 := PrimEqual(symmode, tmp13116)

var ifres13101 Obj

if True == tmp13117 {
tmp13114 := PrimTail(V3270)

tmp13115 := PrimIsPair(tmp13114)

var ifres13103 Obj

if True == tmp13115 {
tmp13111 := PrimTail(V3270)

tmp13112 := PrimTail(tmp13111)

tmp13113 := PrimIsPair(tmp13112)

var ifres13105 Obj

if True == tmp13113 {
tmp13107 := PrimTail(V3270)

tmp13108 := PrimTail(tmp13107)

tmp13109 := PrimTail(tmp13108)

tmp13110 := PrimEqual(Nil, tmp13109)

var ifres13106 Obj

if True == tmp13110 {
ifres13106 = True


} else {
ifres13106 = False


}

ifres13105 = ifres13106


} else {
ifres13105 = False


}

var ifres13104 Obj

if True == ifres13105 {
ifres13104 = True


} else {
ifres13104 = False


}

ifres13103 = ifres13104


} else {
ifres13103 = False


}

var ifres13102 Obj

if True == ifres13103 {
ifres13102 = True


} else {
ifres13102 = False


}

ifres13101 = ifres13102


} else {
ifres13101 = False


}

var ifres13100 Obj

if True == ifres13101 {
ifres13100 = True


} else {
ifres13100 = False


}

ifres13099 = ifres13100


} else {
ifres13099 = False


}

if True == ifres13099 {
tmp13067 := PrimTail(V3270)

tmp13068 := PrimTail(tmp13067)

tmp13069 := PrimHead(tmp13068)

tmp13070 := PrimTail(V3270)

tmp13071 := PrimHead(tmp13070)

tmp13072 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp13071)


tmp13073 := PrimCons(tmp13072, Nil)

__e.Return(PrimCons(tmp13069, tmp13073))
return


} else {
tmp13097 := PrimIsPair(V3270)

var ifres13084 Obj

if True == tmp13097 {
tmp13095 := PrimHead(V3270)

tmp13096 := PrimEqual(symbar_b, tmp13095)

var ifres13086 Obj

if True == tmp13096 {
tmp13093 := PrimTail(V3270)

tmp13094 := PrimIsPair(tmp13093)

var ifres13088 Obj

if True == tmp13094 {
tmp13090 := PrimTail(V3270)

tmp13091 := PrimTail(tmp13090)

tmp13092 := PrimEqual(Nil, tmp13091)

var ifres13089 Obj

if True == tmp13092 {
ifres13089 = True


} else {
ifres13089 = False


}

ifres13088 = ifres13089


} else {
ifres13088 = False


}

var ifres13087 Obj

if True == ifres13088 {
ifres13087 = True


} else {
ifres13087 = False


}

ifres13086 = ifres13087


} else {
ifres13086 = False


}

var ifres13085 Obj

if True == ifres13086 {
ifres13085 = True


} else {
ifres13085 = False


}

ifres13084 = ifres13085


} else {
ifres13084 = False


}

if True == ifres13084 {
tmp13074 := PrimTail(V3270)

__e.Return(PrimHead(tmp13074))
return


} else {
tmp13082 := PrimIsPair(V3270)

if True == tmp13082 {
tmp13075 := PrimHead(V3270)

tmp13076 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp13075)


tmp13077 := PrimTail(V3270)

tmp13078 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp13077)


tmp13079 := PrimCons(tmp13078, Nil)

tmp13080 := PrimCons(tmp13076, tmp13079)

__e.Return(PrimCons(symcons, tmp13080))
return


} else {
__e.Return(V3270)
return
}


}


}


}, 1)

tmp13119 := Call(__e, ns2_1set, symshen_4cons_1form_1with_1modes, tmp13066)


_ = tmp13119

tmp13120 := MakeNative(func(__e *ControlFlow) {
V3271 := __e.Get(1)
_ = V3271
V3272 := __e.Get(2)
_ = V3272
V3273 := __e.Get(3)
_ = V3273
V3274 := __e.Get(4)
_ = V3274
V3275 := __e.Get(5)
_ = V3275
V3276 := __e.Get(6)
_ = V3276
tmp13121 := MakeNative(func(__e *ControlFlow) {
W3277 := __e.Get(1)
_ = W3277
tmp13122 := MakeNative(func(__e *ControlFlow) {
W3278 := __e.Get(1)
_ = W3278
tmp13123 := MakeNative(func(__e *ControlFlow) {
W3279 := __e.Get(1)
_ = W3279
tmp13124 := Call(__e, PrimFunc(symappend), W3278, W3279)


__e.TailApply(PrimFunc(symappend), W3277, tmp13124)
return


}, 1)

tmp13125 := Call(__e, PrimFunc(symshen_4compile_1premises), V3274, V3275)


__e.TailApply(tmp13123, tmp13125)
return


}, 1)

tmp13126 := Call(__e, PrimFunc(symshen_4compile_1side_1conditions), V3273)


__e.TailApply(tmp13122, tmp13126)
return


}, 1)

tmp13127 := Call(__e, PrimFunc(symshen_4compile_1assumptions), V3272, V3271, V3275, V3276)


__e.TailApply(tmp13121, tmp13127)
return


}, 6)

tmp13128 := Call(__e, ns2_1set, symshen_4goals, tmp13120)


_ = tmp13128

tmp13129 := MakeNative(func(__e *ControlFlow) {
V3294 := __e.Get(1)
_ = V3294
V3295 := __e.Get(2)
_ = V3295
V3296 := __e.Get(3)
_ = V3296
V3297 := __e.Get(4)
_ = V3297
tmp13152 := PrimEqual(Nil, V3294)

if True == tmp13152 {
__e.Return(Nil)
return
} else {
tmp13150 := PrimIsPair(V3294)

var ifres13143 Obj

if True == tmp13150 {
tmp13149 := PrimIsPair(V3296)

var ifres13145 Obj

if True == tmp13149 {
tmp13147 := PrimTail(V3296)

tmp13148 := PrimIsPair(tmp13147)

var ifres13146 Obj

if True == tmp13148 {
ifres13146 = True


} else {
ifres13146 = False


}

ifres13145 = ifres13146


} else {
ifres13145 = False


}

var ifres13144 Obj

if True == ifres13145 {
ifres13144 = True


} else {
ifres13144 = False


}

ifres13143 = ifres13144


} else {
ifres13143 = False


}

if True == ifres13143 {
tmp13130 := MakeNative(func(__e *ControlFlow) {
W3298 := __e.Get(1)
_ = W3298
tmp13131 := PrimHead(V3294)

tmp13132 := PrimHead(V3296)

tmp13133 := PrimTail(V3296)

tmp13134 := PrimHead(tmp13133)

tmp13135 := Call(__e, PrimFunc(symshen_4compile_1assumption), tmp13131, tmp13132, tmp13134, V3295, V3297)


tmp13136 := PrimTail(V3294)

tmp13137 := PrimTail(V3296)

tmp13138 := Call(__e, PrimFunc(symshen_4compile_1assumptions), tmp13136, V3295, tmp13137, W3298)


__e.Return(PrimCons(tmp13135, tmp13138))
return


}, 1)

tmp13139 := PrimHead(V3294)

tmp13140 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13139)


tmp13141 := Call(__e, PrimFunc(symappend), tmp13140, V3297)


__e.TailApply(tmp13130, tmp13141)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-assumptions")))
return
}


}


}, 4)

tmp13153 := Call(__e, ns2_1set, symshen_4compile_1assumptions, tmp13129)


_ = tmp13153

tmp13154 := MakeNative(func(__e *ControlFlow) {
V3299 := __e.Get(1)
_ = V3299
V3300 := __e.Get(2)
_ = V3300
V3301 := __e.Get(3)
_ = V3301
V3302 := __e.Get(4)
_ = V3302
V3303 := __e.Get(5)
_ = V3303
tmp13155 := MakeNative(func(__e *ControlFlow) {
W3304 := __e.Get(1)
_ = W3304
tmp13156 := MakeNative(func(__e *ControlFlow) {
W3305 := __e.Get(1)
_ = W3305
tmp13157 := PrimCons(V3301, V3302)

tmp13158 := PrimCons(Nil, tmp13157)

tmp13159 := PrimCons(V3300, tmp13158)

__e.Return(PrimCons(W3304, tmp13159))
return


}, 1)

tmp13160 := Call(__e, PrimFunc(symshen_4compile_1search_1procedure), W3304, V3299, V3300, V3301, V3302, V3303)


__e.TailApply(tmp13156, tmp13160)
return


}, 1)

tmp13161 := Call(__e, PrimFunc(symgensym), symshen_4search)


__e.TailApply(tmp13155, tmp13161)
return


}, 5)

tmp13162 := Call(__e, ns2_1set, symshen_4compile_1assumption, tmp13154)


_ = tmp13162

tmp13163 := MakeNative(func(__e *ControlFlow) {
V3306 := __e.Get(1)
_ = V3306
V3307 := __e.Get(2)
_ = V3307
V3308 := __e.Get(3)
_ = V3308
V3309 := __e.Get(4)
_ = V3309
V3310 := __e.Get(5)
_ = V3310
V3311 := __e.Get(6)
_ = V3311
tmp13164 := MakeNative(func(__e *ControlFlow) {
W3312 := __e.Get(1)
_ = W3312
tmp13165 := MakeNative(func(__e *ControlFlow) {
W3313 := __e.Get(1)
_ = W3313
tmp13166 := MakeNative(func(__e *ControlFlow) {
W3314 := __e.Get(1)
_ = W3314
tmp13167 := Call(__e, PrimFunc(symappend), W3313, W3314)


tmp13168 := PrimCons(V3306, tmp13167)

tmp13169 := PrimCons(symdefprolog, tmp13168)

__e.TailApply(PrimFunc(symeval), tmp13169)
return


}, 1)

tmp13170 := Call(__e, PrimFunc(symshen_4keep_1looking), V3306, V3308, W3312, V3309, V3310)


__e.TailApply(tmp13166, tmp13170)
return


}, 1)

tmp13171 := Call(__e, PrimFunc(symshen_4foundit_b), V3307, V3308, W3312, V3309, V3310, V3311)


__e.TailApply(tmp13165, tmp13171)
return


}, 1)

tmp13172 := Call(__e, PrimFunc(symgensym), symPrevious)


__e.TailApply(tmp13164, tmp13172)
return


}, 6)

tmp13173 := Call(__e, ns2_1set, symshen_4compile_1search_1procedure, tmp13163)


_ = tmp13173

tmp13174 := MakeNative(func(__e *ControlFlow) {
V3315 := __e.Get(1)
_ = V3315
V3316 := __e.Get(2)
_ = V3316
V3317 := __e.Get(3)
_ = V3317
V3318 := __e.Get(4)
_ = V3318
V3319 := __e.Get(5)
_ = V3319
V3320 := __e.Get(6)
_ = V3320
tmp13175 := MakeNative(func(__e *ControlFlow) {
W3321 := __e.Get(1)
_ = W3321
tmp13176 := MakeNative(func(__e *ControlFlow) {
W3322 := __e.Get(1)
_ = W3322
tmp13177 := MakeNative(func(__e *ControlFlow) {
W3323 := __e.Get(1)
_ = W3323
tmp13178 := MakeNative(func(__e *ControlFlow) {
W3324 := __e.Get(1)
_ = W3324
tmp13179 := PrimCons(sym_5_1_1, Nil)

tmp13180 := PrimIntern(MakeString(";"))

tmp13181 := PrimCons(tmp13180, Nil)

tmp13182 := Call(__e, PrimFunc(symappend), W3324, tmp13181)


tmp13183 := Call(__e, PrimFunc(symappend), tmp13179, tmp13182)


__e.TailApply(PrimFunc(symappend), W3323, tmp13183)
return


}, 1)

tmp13184 := Call(__e, PrimFunc(symshen_4body_1foundit_b), V3316, V3317, V3318, W3322)


__e.TailApply(tmp13178, tmp13184)
return


}, 1)

tmp13185 := Call(__e, PrimFunc(symshen_4head_1foundit_b), V3315, V3316, V3317, V3318, V3319, W3322)


__e.TailApply(tmp13177, tmp13185)
return


}, 1)

tmp13186 := Call(__e, PrimFunc(symshen_4tabulate_1passive), W3321)


__e.TailApply(tmp13176, tmp13186)
return


}, 1)

tmp13187 := Call(__e, PrimFunc(symshen_4passive), V3315, V3320)


__e.TailApply(tmp13175, tmp13187)
return


}, 6)

tmp13188 := Call(__e, ns2_1set, symshen_4foundit_b, tmp13174)


_ = tmp13188

tmp13189 := MakeNative(func(__e *ControlFlow) {
V3325 := __e.Get(1)
_ = V3325
V3326 := __e.Get(2)
_ = V3326
V3327 := __e.Get(3)
_ = V3327
V3328 := __e.Get(4)
_ = V3328
V3329 := __e.Get(5)
_ = V3329
tmp13190 := MakeNative(func(__e *ControlFlow) {
W3330 := __e.Get(1)
_ = W3330
tmp13191 := MakeNative(func(__e *ControlFlow) {
W3331 := __e.Get(1)
_ = W3331
tmp13192 := MakeNative(func(__e *ControlFlow) {
W3332 := __e.Get(1)
_ = W3332
tmp13193 := PrimCons(sym_5_1_1, Nil)

tmp13194 := PrimIntern(MakeString(";"))

tmp13195 := PrimCons(tmp13194, Nil)

tmp13196 := Call(__e, PrimFunc(symappend), W3332, tmp13195)


tmp13197 := Call(__e, PrimFunc(symappend), tmp13193, tmp13196)


__e.TailApply(PrimFunc(symappend), W3331, tmp13197)
return


}, 1)

tmp13198 := PrimCons(V3327, Nil)

tmp13199 := PrimCons(W3330, tmp13198)

tmp13200 := PrimCons(symcons, tmp13199)

tmp13201 := PrimCons(V3328, V3329)

tmp13202 := PrimCons(tmp13200, tmp13201)

tmp13203 := PrimCons(V3326, tmp13202)

tmp13204 := PrimCons(V3325, tmp13203)

tmp13205 := PrimCons(tmp13204, Nil)

__e.TailApply(tmp13192, tmp13205)
return


}, 1)

tmp13206 := PrimCons(V3326, Nil)

tmp13207 := PrimCons(W3330, tmp13206)

tmp13208 := PrimCons(symcons, tmp13207)

tmp13209 := PrimCons(tmp13208, Nil)

tmp13210 := PrimCons(sym_1, tmp13209)

tmp13211 := PrimCons(V3328, V3329)

tmp13212 := PrimCons(V3327, tmp13211)

tmp13213 := PrimCons(tmp13210, tmp13212)

__e.TailApply(tmp13191, tmp13213)
return


}, 1)

tmp13214 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp13190, tmp13214)
return


}, 5)

tmp13215 := Call(__e, ns2_1set, symshen_4keep_1looking, tmp13189)


_ = tmp13215

tmp13216 := MakeNative(func(__e *ControlFlow) {
V3337 := __e.Get(1)
_ = V3337
V3338 := __e.Get(2)
_ = V3338
tmp13224 := PrimIsPair(V3337)

if True == tmp13224 {
tmp13217 := PrimHead(V3337)

tmp13218 := Call(__e, PrimFunc(symshen_4passive), tmp13217, V3338)


tmp13219 := PrimTail(V3337)

tmp13220 := Call(__e, PrimFunc(symshen_4passive), tmp13219, V3338)


__e.TailApply(PrimFunc(symunion), tmp13218, tmp13220)
return


} else {
tmp13222 := Call(__e, PrimFunc(symshen_4passive_2), V3337, V3338)


if True == tmp13222 {
__e.Return(PrimCons(V3337, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}, 2)

tmp13225 := Call(__e, ns2_1set, symshen_4passive, tmp13216)


_ = tmp13225

tmp13226 := MakeNative(func(__e *ControlFlow) {
V3339 := __e.Get(1)
_ = V3339
V3340 := __e.Get(2)
_ = V3340
tmp13230 := Call(__e, PrimFunc(symelement_2), V3339, V3340)


tmp13231 := PrimNot(tmp13230)

if True == tmp13231 {
tmp13228 := PrimIsVariable(V3339)

if True == tmp13228 {
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


}, 2)

tmp13232 := Call(__e, ns2_1set, symshen_4passive_2, tmp13226)


_ = tmp13232

tmp13233 := MakeNative(func(__e *ControlFlow) {
V3341 := __e.Get(1)
_ = V3341
tmp13234 := MakeNative(func(__e *ControlFlow) {
Z3342 := __e.Get(1)
_ = Z3342
tmp13235 := Call(__e, PrimFunc(symgensym), symV)


__e.Return(PrimCons(Z3342, tmp13235))
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp13234, V3341)
return


}, 1)

tmp13236 := Call(__e, ns2_1set, symshen_4tabulate_1passive, tmp13233)


_ = tmp13236

tmp13237 := MakeNative(func(__e *ControlFlow) {
V3343 := __e.Get(1)
_ = V3343
V3344 := __e.Get(2)
_ = V3344
V3345 := __e.Get(3)
_ = V3345
V3346 := __e.Get(4)
_ = V3346
V3347 := __e.Get(5)
_ = V3347
V3348 := __e.Get(6)
_ = V3348
tmp13238 := MakeNative(func(__e *ControlFlow) {
W3349 := __e.Get(1)
_ = W3349
tmp13239 := Call(__e, PrimFunc(symshen_4optimise_1typing), V3343)


tmp13240 := PrimCons(V3344, Nil)

tmp13241 := PrimCons(tmp13239, tmp13240)

tmp13242 := PrimCons(symcons, tmp13241)

tmp13243 := PrimCons(tmp13242, Nil)

tmp13244 := PrimCons(sym_1, tmp13243)

tmp13245 := PrimCons(V3346, W3349)

tmp13246 := PrimCons(V3345, tmp13245)

__e.Return(PrimCons(tmp13244, tmp13246))
return


}, 1)

tmp13247 := Call(__e, PrimFunc(symshen_4optimise_1passive), V3347, V3348)


__e.TailApply(tmp13238, tmp13247)
return


}, 6)

tmp13248 := Call(__e, ns2_1set, symshen_4head_1foundit_b, tmp13237)


_ = tmp13248

tmp13249 := MakeNative(func(__e *ControlFlow) {
V3350 := __e.Get(1)
_ = V3350
V3351 := __e.Get(2)
_ = V3351
tmp13250 := MakeNative(func(__e *ControlFlow) {
Z3352 := __e.Get(1)
_ = Z3352
__e.TailApply(PrimFunc(symshen_4optimise_1passive_1h), Z3352, V3351)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13250, V3350)
return


}, 2)

tmp13251 := Call(__e, ns2_1set, symshen_4optimise_1passive, tmp13249)


_ = tmp13251

tmp13252 := MakeNative(func(__e *ControlFlow) {
V3353 := __e.Get(1)
_ = V3353
V3354 := __e.Get(2)
_ = V3354
tmp13253 := MakeNative(func(__e *ControlFlow) {
W3355 := __e.Get(1)
_ = W3355
tmp13255 := Call(__e, PrimFunc(symempty_2), W3355)


if True == tmp13255 {
__e.Return(V3353)
return
} else {
__e.Return(PrimTail(W3355))
return
}


}, 1)

tmp13256 := Call(__e, PrimFunc(symassoc), V3353, V3354)


__e.TailApply(tmp13253, tmp13256)
return


}, 2)

tmp13257 := Call(__e, ns2_1set, symshen_4optimise_1passive_1h, tmp13252)


_ = tmp13257

tmp13258 := MakeNative(func(__e *ControlFlow) {
V3364 := __e.Get(1)
_ = V3364
V3365 := __e.Get(2)
_ = V3365
V3366 := __e.Get(3)
_ = V3366
V3367 := __e.Get(4)
_ = V3367
tmp13285 := PrimEqual(Nil, V3367)

if True == tmp13285 {
tmp13259 := PrimCons(V3365, Nil)

tmp13260 := PrimCons(MakeNumber(1), tmp13259)

tmp13261 := PrimCons(V3364, Nil)

tmp13262 := PrimCons(MakeNumber(1), tmp13261)

tmp13263 := PrimCons(tmp13262, Nil)

tmp13264 := PrimCons(tmp13260, tmp13263)

tmp13265 := PrimCons(symappend, tmp13264)

tmp13266 := PrimCons(tmp13265, Nil)

tmp13267 := PrimCons(V3366, tmp13266)

tmp13268 := PrimCons(symbind, tmp13267)

__e.Return(PrimCons(tmp13268, Nil))
return


} else {
tmp13283 := PrimIsPair(V3367)

var ifres13279 Obj

if True == tmp13283 {
tmp13281 := PrimHead(V3367)

tmp13282 := PrimIsPair(tmp13281)

var ifres13280 Obj

if True == tmp13282 {
ifres13280 = True


} else {
ifres13280 = False


}

ifres13279 = ifres13280


} else {
ifres13279 = False


}

if True == ifres13279 {
tmp13269 := PrimHead(V3367)

tmp13270 := PrimTail(tmp13269)

tmp13271 := PrimHead(V3367)

tmp13272 := PrimHead(tmp13271)

tmp13273 := PrimCons(tmp13272, Nil)

tmp13274 := PrimCons(tmp13270, tmp13273)

tmp13275 := PrimCons(symbind, tmp13274)

tmp13276 := PrimTail(V3367)

tmp13277 := Call(__e, PrimFunc(symshen_4body_1foundit_b), V3364, V3365, V3366, tmp13276)


__e.Return(PrimCons(tmp13275, tmp13277))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.body-foundit!")))
return
}


}


}, 4)

tmp13286 := Call(__e, ns2_1set, symshen_4body_1foundit_b, tmp13258)


_ = tmp13286

tmp13287 := MakeNative(func(__e *ControlFlow) {
V3368 := __e.Get(1)
_ = V3368
tmp13288 := MakeNative(func(__e *ControlFlow) {
Z3369 := __e.Get(1)
_ = Z3369
__e.TailApply(PrimFunc(symshen_4compile_1side_1condition), Z3369)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13288, V3368)
return


}, 1)

tmp13289 := Call(__e, ns2_1set, symshen_4compile_1side_1conditions, tmp13287)


_ = tmp13289

tmp13290 := MakeNative(func(__e *ControlFlow) {
V3372 := __e.Get(1)
_ = V3372
tmp13350 := PrimIsPair(V3372)

var ifres13331 Obj

if True == tmp13350 {
tmp13348 := PrimHead(V3372)

tmp13349 := PrimEqual(symlet, tmp13348)

var ifres13333 Obj

if True == tmp13349 {
tmp13346 := PrimTail(V3372)

tmp13347 := PrimIsPair(tmp13346)

var ifres13335 Obj

if True == tmp13347 {
tmp13343 := PrimTail(V3372)

tmp13344 := PrimTail(tmp13343)

tmp13345 := PrimIsPair(tmp13344)

var ifres13337 Obj

if True == tmp13345 {
tmp13339 := PrimTail(V3372)

tmp13340 := PrimTail(tmp13339)

tmp13341 := PrimTail(tmp13340)

tmp13342 := PrimEqual(Nil, tmp13341)

var ifres13338 Obj

if True == tmp13342 {
ifres13338 = True


} else {
ifres13338 = False


}

ifres13337 = ifres13338


} else {
ifres13337 = False


}

var ifres13336 Obj

if True == ifres13337 {
ifres13336 = True


} else {
ifres13336 = False


}

ifres13335 = ifres13336


} else {
ifres13335 = False


}

var ifres13334 Obj

if True == ifres13335 {
ifres13334 = True


} else {
ifres13334 = False


}

ifres13333 = ifres13334


} else {
ifres13333 = False


}

var ifres13332 Obj

if True == ifres13333 {
ifres13332 = True


} else {
ifres13332 = False


}

ifres13331 = ifres13332


} else {
ifres13331 = False


}

if True == ifres13331 {
tmp13291 := PrimTail(V3372)

__e.Return(PrimCons(symis, tmp13291))
return


} else {
tmp13329 := PrimIsPair(V3372)

var ifres13310 Obj

if True == tmp13329 {
tmp13327 := PrimHead(V3372)

tmp13328 := PrimEqual(symshen_4let_b, tmp13327)

var ifres13312 Obj

if True == tmp13328 {
tmp13325 := PrimTail(V3372)

tmp13326 := PrimIsPair(tmp13325)

var ifres13314 Obj

if True == tmp13326 {
tmp13322 := PrimTail(V3372)

tmp13323 := PrimTail(tmp13322)

tmp13324 := PrimIsPair(tmp13323)

var ifres13316 Obj

if True == tmp13324 {
tmp13318 := PrimTail(V3372)

tmp13319 := PrimTail(tmp13318)

tmp13320 := PrimTail(tmp13319)

tmp13321 := PrimEqual(Nil, tmp13320)

var ifres13317 Obj

if True == tmp13321 {
ifres13317 = True


} else {
ifres13317 = False


}

ifres13316 = ifres13317


} else {
ifres13316 = False


}

var ifres13315 Obj

if True == ifres13316 {
ifres13315 = True


} else {
ifres13315 = False


}

ifres13314 = ifres13315


} else {
ifres13314 = False


}

var ifres13313 Obj

if True == ifres13314 {
ifres13313 = True


} else {
ifres13313 = False


}

ifres13312 = ifres13313


} else {
ifres13312 = False


}

var ifres13311 Obj

if True == ifres13312 {
ifres13311 = True


} else {
ifres13311 = False


}

ifres13310 = ifres13311


} else {
ifres13310 = False


}

if True == ifres13310 {
tmp13292 := PrimTail(V3372)

__e.Return(PrimCons(symis_b, tmp13292))
return


} else {
tmp13308 := PrimIsPair(V3372)

var ifres13295 Obj

if True == tmp13308 {
tmp13306 := PrimHead(V3372)

tmp13307 := PrimEqual(symif, tmp13306)

var ifres13297 Obj

if True == tmp13307 {
tmp13304 := PrimTail(V3372)

tmp13305 := PrimIsPair(tmp13304)

var ifres13299 Obj

if True == tmp13305 {
tmp13301 := PrimTail(V3372)

tmp13302 := PrimTail(tmp13301)

tmp13303 := PrimEqual(Nil, tmp13302)

var ifres13300 Obj

if True == tmp13303 {
ifres13300 = True


} else {
ifres13300 = False


}

ifres13299 = ifres13300


} else {
ifres13299 = False


}

var ifres13298 Obj

if True == ifres13299 {
ifres13298 = True


} else {
ifres13298 = False


}

ifres13297 = ifres13298


} else {
ifres13297 = False


}

var ifres13296 Obj

if True == ifres13297 {
ifres13296 = True


} else {
ifres13296 = False


}

ifres13295 = ifres13296


} else {
ifres13295 = False


}

if True == ifres13295 {
tmp13293 := PrimTail(V3372)

__e.Return(PrimCons(symwhen, tmp13293))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-side-condition")))
return
}


}


}


}, 1)

tmp13351 := Call(__e, ns2_1set, symshen_4compile_1side_1condition, tmp13290)


_ = tmp13351

tmp13352 := MakeNative(func(__e *ControlFlow) {
V3373 := __e.Get(1)
_ = V3373
V3374 := __e.Get(2)
_ = V3374
tmp13353 := MakeNative(func(__e *ControlFlow) {
W3375 := __e.Get(1)
_ = W3375
tmp13354 := MakeNative(func(__e *ControlFlow) {
Z3376 := __e.Get(1)
_ = Z3376
__e.TailApply(PrimFunc(symshen_4compile_1premise), Z3376, W3375)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13354, V3373)
return


}, 1)

tmp13355 := Call(__e, PrimFunc(symreverse), V3374)


tmp13356 := PrimHead(tmp13355)

__e.TailApply(tmp13353, tmp13356)
return


}, 2)

tmp13357 := Call(__e, ns2_1set, symshen_4compile_1premises, tmp13352)


_ = tmp13357

tmp13358 := MakeNative(func(__e *ControlFlow) {
V3383 := __e.Get(1)
_ = V3383
V3384 := __e.Get(2)
_ = V3384
tmp13375 := PrimEqual(sym_b, V3383)

if True == tmp13375 {
__e.Return(sym_b)
return
} else {
tmp13373 := PrimIsPair(V3383)

var ifres13364 Obj

if True == tmp13373 {
tmp13371 := PrimTail(V3383)

tmp13372 := PrimIsPair(tmp13371)

var ifres13366 Obj

if True == tmp13372 {
tmp13368 := PrimTail(V3383)

tmp13369 := PrimTail(tmp13368)

tmp13370 := PrimEqual(Nil, tmp13369)

var ifres13367 Obj

if True == tmp13370 {
ifres13367 = True


} else {
ifres13367 = False


}

ifres13366 = ifres13367


} else {
ifres13366 = False


}

var ifres13365 Obj

if True == ifres13366 {
ifres13365 = True


} else {
ifres13365 = False


}

ifres13364 = ifres13365


} else {
ifres13364 = False


}

if True == ifres13364 {
tmp13359 := PrimHead(V3383)

tmp13360 := Call(__e, PrimFunc(symreverse), tmp13359)


tmp13361 := PrimTail(V3383)

tmp13362 := PrimHead(tmp13361)

__e.TailApply(PrimFunc(symshen_4compile_1premise_1h), tmp13360, tmp13362, V3384)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.premise")))
return
}


}


}, 2)

tmp13376 := Call(__e, ns2_1set, symshen_4compile_1premise, tmp13358)


_ = tmp13376

tmp13377 := MakeNative(func(__e *ControlFlow) {
V3391 := __e.Get(1)
_ = V3391
V3392 := __e.Get(2)
_ = V3392
V3393 := __e.Get(3)
_ = V3393
tmp13390 := PrimEqual(Nil, V3391)

if True == tmp13390 {
tmp13378 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), V3392)


tmp13379 := PrimCons(V3393, Nil)

tmp13380 := PrimCons(tmp13378, tmp13379)

__e.Return(PrimCons(symshen_4system_1S, tmp13380))
return


} else {
tmp13388 := PrimIsPair(V3391)

if True == tmp13388 {
tmp13381 := PrimTail(V3391)

tmp13382 := PrimHead(V3391)

tmp13383 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp13382)


tmp13384 := PrimCons(V3393, Nil)

tmp13385 := PrimCons(tmp13383, tmp13384)

tmp13386 := PrimCons(symcons, tmp13385)

__e.TailApply(PrimFunc(symshen_4compile_1premise_1h), tmp13381, V3392, tmp13386)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-premise-h")))
return
}


}


}, 3)

tmp13391 := Call(__e, ns2_1set, symshen_4compile_1premise_1h, tmp13377)


_ = tmp13391

tmp13392 := MakeNative(func(__e *ControlFlow) {
V3394 := __e.Get(1)
_ = V3394
tmp13416 := PrimIsPair(V3394)

var ifres13403 Obj

if True == tmp13416 {
tmp13414 := PrimHead(V3394)

tmp13415 := PrimEqual(symbar_b, tmp13414)

var ifres13405 Obj

if True == tmp13415 {
tmp13412 := PrimTail(V3394)

tmp13413 := PrimIsPair(tmp13412)

var ifres13407 Obj

if True == tmp13413 {
tmp13409 := PrimTail(V3394)

tmp13410 := PrimTail(tmp13409)

tmp13411 := PrimEqual(Nil, tmp13410)

var ifres13408 Obj

if True == tmp13411 {
ifres13408 = True


} else {
ifres13408 = False


}

ifres13407 = ifres13408


} else {
ifres13407 = False


}

var ifres13406 Obj

if True == ifres13407 {
ifres13406 = True


} else {
ifres13406 = False


}

ifres13405 = ifres13406


} else {
ifres13405 = False


}

var ifres13404 Obj

if True == ifres13405 {
ifres13404 = True


} else {
ifres13404 = False


}

ifres13403 = ifres13404


} else {
ifres13403 = False


}

if True == ifres13403 {
tmp13393 := PrimTail(V3394)

__e.Return(PrimHead(tmp13393))
return


} else {
tmp13401 := PrimIsPair(V3394)

if True == tmp13401 {
tmp13394 := PrimHead(V3394)

tmp13395 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp13394)


tmp13396 := PrimTail(V3394)

tmp13397 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp13396)


tmp13398 := PrimCons(tmp13397, Nil)

tmp13399 := PrimCons(tmp13395, tmp13398)

__e.Return(PrimCons(symcons, tmp13399))
return


} else {
__e.Return(V3394)
return
}


}


}, 1)

tmp13417 := Call(__e, ns2_1set, symshen_4cons_1form_1no_1modes, tmp13392)


_ = tmp13417

tmp13418 := MakeNative(func(__e *ControlFlow) {
V3395 := __e.Get(1)
_ = V3395
tmp13419 := MakeNative(func(__e *ControlFlow) {
W3396 := __e.Get(1)
_ = W3396
tmp13420 := MakeNative(func(__e *ControlFlow) {
W3398 := __e.Get(1)
_ = W3398
tmp13421 := MakeNative(func(__e *ControlFlow) {
W3399 := __e.Get(1)
_ = W3399
tmp13422 := MakeNative(func(__e *ControlFlow) {
W3400 := __e.Get(1)
_ = W3400
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3400)
return
}, 1)

tmp13423 := PrimSet(symshen_4_ddatatypes_d, W3399)

__e.TailApply(tmp13422, tmp13423)
return


}, 1)

tmp13424 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3396, W3398)


__e.TailApply(tmp13421, tmp13424)
return


}, 1)

tmp13425 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp13420, tmp13425)
return


}, 1)

tmp13426 := MakeNative(func(__e *ControlFlow) {
Z3397 := __e.Get(1)
_ = Z3397
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3397)
return
}, 1)

tmp13427 := Call(__e, PrimFunc(symmap), tmp13426, V3395)


__e.TailApply(tmp13419, tmp13427)
return


}, 1)

tmp13428 := Call(__e, ns2_1set, sympreclude, tmp13418)


_ = tmp13428

tmp13429 := MakeNative(func(__e *ControlFlow) {
V3405 := __e.Get(1)
_ = V3405
V3406 := __e.Get(2)
_ = V3406
tmp13436 := PrimEqual(Nil, V3405)

if True == tmp13436 {
__e.Return(V3406)
return
} else {
tmp13434 := PrimIsPair(V3405)

if True == tmp13434 {
tmp13430 := PrimTail(V3405)

tmp13431 := PrimHead(V3405)

tmp13432 := Call(__e, PrimFunc(symshen_4unassoc), tmp13431, V3406)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13430, tmp13432)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
return
}


}


}, 2)

tmp13437 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13429)


_ = tmp13437

tmp13438 := MakeNative(func(__e *ControlFlow) {
V3407 := __e.Get(1)
_ = V3407
tmp13439 := MakeNative(func(__e *ControlFlow) {
Z3408 := __e.Get(1)
_ = Z3408
__e.Return(PrimHead(Z3408))
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13439, V3407)
return


}, 1)

tmp13440 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13438)


_ = tmp13440

tmp13441 := MakeNative(func(__e *ControlFlow) {
V3409 := __e.Get(1)
_ = V3409
tmp13442 := MakeNative(func(__e *ControlFlow) {
W3410 := __e.Get(1)
_ = W3410
tmp13443 := MakeNative(func(__e *ControlFlow) {
W3412 := __e.Get(1)
_ = W3412
tmp13444 := MakeNative(func(__e *ControlFlow) {
W3414 := __e.Get(1)
_ = W3414
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3414)
return
}, 1)

tmp13445 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp13444, tmp13445)
return


}, 1)

tmp13446 := MakeNative(func(__e *ControlFlow) {
Z3413 := __e.Get(1)
_ = Z3413
tmp13447 := Call(__e, PrimFunc(symfn), Z3413)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3413, tmp13447)
return


}, 1)

tmp13448 := Call(__e, PrimFunc(symmap), tmp13446, W3410)


__e.TailApply(tmp13443, tmp13448)
return


}, 1)

tmp13449 := MakeNative(func(__e *ControlFlow) {
Z3411 := __e.Get(1)
_ = Z3411
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3411)
return
}, 1)

tmp13450 := Call(__e, PrimFunc(symmap), tmp13449, V3409)


__e.TailApply(tmp13442, tmp13450)
return


}, 1)

tmp13451 := Call(__e, ns2_1set, syminclude, tmp13441)


_ = tmp13451

tmp13452 := MakeNative(func(__e *ControlFlow) {
V3415 := __e.Get(1)
_ = V3415
tmp13453 := MakeNative(func(__e *ControlFlow) {
W3416 := __e.Get(1)
_ = W3416
tmp13454 := MakeNative(func(__e *ControlFlow) {
W3417 := __e.Get(1)
_ = W3417
tmp13455 := MakeNative(func(__e *ControlFlow) {
W3419 := __e.Get(1)
_ = W3419
tmp13456 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13456)
return


}, 1)

tmp13457 := MakeNative(func(__e *ControlFlow) {
Z3420 := __e.Get(1)
_ = Z3420
tmp13458 := Call(__e, PrimFunc(symfn), Z3420)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3420, tmp13458)
return


}, 1)

tmp13459 := Call(__e, PrimFunc(symmap), tmp13457, W3417)


__e.TailApply(tmp13455, tmp13459)
return


}, 1)

tmp13460 := MakeNative(func(__e *ControlFlow) {
Z3418 := __e.Get(1)
_ = Z3418
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3418)
return
}, 1)

tmp13461 := Call(__e, PrimFunc(symmap), tmp13460, V3415)


__e.TailApply(tmp13454, tmp13461)
return


}, 1)

tmp13462 := PrimSet(symshen_4_ddatatypes_d, Nil)

__e.TailApply(tmp13453, tmp13462)
return


}, 1)

tmp13463 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13452)


_ = tmp13463

tmp13464 := MakeNative(func(__e *ControlFlow) {
V3421 := __e.Get(1)
_ = V3421
tmp13465 := MakeNative(func(__e *ControlFlow) {
W3422 := __e.Get(1)
_ = W3422
tmp13466 := MakeNative(func(__e *ControlFlow) {
W3424 := __e.Get(1)
_ = W3424
tmp13467 := MakeNative(func(__e *ControlFlow) {
W3425 := __e.Get(1)
_ = W3425
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3425)
return
}, 1)

tmp13468 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3422, W3424)


tmp13469 := PrimSet(symshen_4_ddatatypes_d, tmp13468)

__e.TailApply(tmp13467, tmp13469)
return


}, 1)

tmp13470 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(tmp13466, tmp13470)
return


}, 1)

tmp13471 := MakeNative(func(__e *ControlFlow) {
Z3423 := __e.Get(1)
_ = Z3423
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3423)
return
}, 1)

tmp13472 := Call(__e, PrimFunc(symmap), tmp13471, V3421)


__e.TailApply(tmp13465, tmp13472)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13464)
return




}, 0)

