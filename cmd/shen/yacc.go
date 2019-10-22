package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4yacc Obj // shen.yacc
var __defun__shen_4yacc_1_6shen Obj // shen.yacc->shen
var __defun__shen_4kill_1code Obj // shen.kill-code
var __defun__kill Obj // kill
var __defun__shen_4analyse_1kill Obj // shen.analyse-kill
var __defun__shen_4split__cc__rules Obj // shen.split_cc_rules
var __defun__shen_4split__cc__rule Obj // shen.split_cc_rule
var __defun__shen_4semantic_1completion_1warning Obj // shen.semantic-completion-warning
var __defun__shen_4default__semantics Obj // shen.default_semantics
var __defun__shen_4grammar__symbol_2 Obj // shen.grammar_symbol?
var __defun__shen_4yacc__cases Obj // shen.yacc_cases
var __defun__shen_4cc__body Obj // shen.cc_body
var __defun__shen_4syntax Obj // shen.syntax
var __defun__shen_4list_1stream Obj // shen.list-stream
var __defun__shen_4decons Obj // shen.decons
var __defun__shen_4insert_1runon Obj // shen.insert-runon
var __defun__shen_4strip_1pathname Obj // shen.strip-pathname
var __defun__shen_4recursive__descent Obj // shen.recursive_descent
var __defun__shen_4variable_1match Obj // shen.variable-match
var __defun__shen_4terminal_2 Obj // shen.terminal?
var __defun__shen_4jump__stream_2 Obj // shen.jump_stream?
var __defun__shen_4check__stream Obj // shen.check_stream
var __defun__shen_4jump__stream Obj // shen.jump_stream
var __defun__shen_4semantics Obj // shen.semantics
var __defun__shen_4pair Obj // shen.pair
var __defun__shen_4hdtl Obj // shen.hdtl
var __defun__shen_4hdhd Obj // shen.hdhd
var __defun__shen_4tlhd Obj // shen.tlhd
var __defun__shen_4snd_1or_1fail Obj // shen.snd-or-fail
var __defun__fail Obj // fail
var __defun___5_b_6 Obj // <!>
var __defun___5e_6 Obj // <e>

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg298913 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg298913)
return
}, 0))
__defun__shen_4yacc = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4275 := __args[0]
_ = V4275
reg298914 := PrimIsPair(V4275)
var reg298929 Obj
if reg298914 == True {
reg298915 := MakeSymbol("defcc")
reg298916 := PrimHead(V4275)
reg298917 := PrimEqual(reg298915, reg298916)
var reg298924 Obj
if reg298917 == True {
reg298918 := PrimTail(V4275)
reg298919 := PrimIsPair(reg298918)
var reg298922 Obj
if reg298919 == True {
reg298920 := True;
reg298922 = reg298920
} else {
reg298921 := False;
reg298922 = reg298921
}
reg298924 = reg298922
} else {
reg298923 := False;
reg298924 = reg298923
}
var reg298927 Obj
if reg298924 == True {
reg298925 := True;
reg298927 = reg298925
} else {
reg298926 := False;
reg298927 = reg298926
}
reg298929 = reg298927
} else {
reg298928 := False;
reg298929 = reg298928
}
if reg298929 == True {
reg298930 := PrimTail(V4275)
reg298931 := PrimHead(reg298930)
reg298932 := PrimTail(V4275)
reg298933 := PrimTail(reg298932)
__ctx.TailApply(__defun__shen_4yacc_1_6shen, reg298931, reg298933)
return
} else {
reg298935 := MakeSymbol("shen.yacc")
__ctx.TailApply(__defun__shen_4f__error, reg298935)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.yacc", value: __defun__shen_4yacc})

__defun__shen_4yacc_1_6shen = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4278 := __args[0]
_ = V4278
V4279 := __args[1]
_ = V4279
reg298937 := True;
reg298938 := Nil;
reg298939 := __e.Call(__defun__shen_4split__cc__rules, reg298937, V4279, reg298938)
CCRules := reg298939
_ = CCRules
reg298940 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4cc__body, X)
return
}, 1)
reg298942 := __e.Call(__defun__map, reg298940, CCRules)
CCBody := reg298942
_ = CCBody
reg298943 := __e.Call(__defun__shen_4yacc__cases, CCBody)
YaccCases := reg298943
_ = YaccCases
reg298944 := MakeSymbol("define")
reg298945 := MakeSymbol("Stream")
reg298946 := MakeSymbol("->")
reg298947 := __e.Call(__defun__shen_4kill_1code, YaccCases)
reg298948 := Nil;
reg298949 := PrimCons(reg298947, reg298948)
reg298950 := PrimCons(reg298946, reg298949)
reg298951 := PrimCons(reg298945, reg298950)
reg298952 := PrimCons(V4278, reg298951)
reg298953 := PrimCons(reg298944, reg298952)
__ctx.Return(reg298953)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.yacc->shen", value: __defun__shen_4yacc_1_6shen})

__defun__shen_4kill_1code = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4281 := __args[0]
_ = V4281
reg298954 := MakeSymbol("kill")
reg298955 := __e.Call(__defun__occurrences, reg298954, V4281)
reg298956 := MakeNumber(0)
reg298957 := PrimGreatThan(reg298955, reg298956)
if reg298957 == True {
reg298958 := MakeSymbol("trap-error")
reg298959 := MakeSymbol("lambda")
reg298960 := MakeSymbol("E")
reg298961 := MakeSymbol("shen.analyse-kill")
reg298962 := MakeSymbol("E")
reg298963 := Nil;
reg298964 := PrimCons(reg298962, reg298963)
reg298965 := PrimCons(reg298961, reg298964)
reg298966 := Nil;
reg298967 := PrimCons(reg298965, reg298966)
reg298968 := PrimCons(reg298960, reg298967)
reg298969 := PrimCons(reg298959, reg298968)
reg298970 := Nil;
reg298971 := PrimCons(reg298969, reg298970)
reg298972 := PrimCons(V4281, reg298971)
reg298973 := PrimCons(reg298958, reg298972)
__ctx.Return(reg298973)
return
} else {
__ctx.Return(V4281)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.kill-code", value: __defun__shen_4kill_1code})

__defun__kill = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg298974 := MakeString("yacc kill")
reg298975 := PrimSimpleError(reg298974)
__ctx.Return(reg298975)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "kill", value: __defun__kill})

__defun__shen_4analyse_1kill = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4283 := __args[0]
_ = V4283
reg298976 := PrimErrorToString(V4283)
String := reg298976
_ = String
reg298977 := MakeString("yacc kill")
reg298978 := PrimEqual(String, reg298977)
if reg298978 == True {
__ctx.TailApply(__defun__fail)
return
} else {
__ctx.Return(V4283)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.analyse-kill", value: __defun__shen_4analyse_1kill})

__defun__shen_4split__cc__rules = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4289 := __args[0]
_ = V4289
V4290 := __args[1]
_ = V4290
V4291 := __args[2]
_ = V4291
reg298980 := Nil;
reg298981 := PrimEqual(reg298980, V4290)
var reg298988 Obj
if reg298981 == True {
reg298982 := Nil;
reg298983 := PrimEqual(reg298982, V4291)
var reg298986 Obj
if reg298983 == True {
reg298984 := True;
reg298986 = reg298984
} else {
reg298985 := False;
reg298986 = reg298985
}
reg298988 = reg298986
} else {
reg298987 := False;
reg298988 = reg298987
}
if reg298988 == True {
reg298989 := Nil;
__ctx.Return(reg298989)
return
} else {
reg298990 := Nil;
reg298991 := PrimEqual(reg298990, V4290)
if reg298991 == True {
reg298992 := __e.Call(__defun__reverse, V4291)
reg298993 := Nil;
reg298994 := __e.Call(__defun__shen_4split__cc__rule, V4289, reg298992, reg298993)
reg298995 := Nil;
reg298996 := PrimCons(reg298994, reg298995)
__ctx.Return(reg298996)
return
} else {
reg298997 := PrimIsPair(V4290)
var reg299005 Obj
if reg298997 == True {
reg298998 := MakeSymbol(";")
reg298999 := PrimHead(V4290)
reg299000 := PrimEqual(reg298998, reg298999)
var reg299003 Obj
if reg299000 == True {
reg299001 := True;
reg299003 = reg299001
} else {
reg299002 := False;
reg299003 = reg299002
}
reg299005 = reg299003
} else {
reg299004 := False;
reg299005 = reg299004
}
if reg299005 == True {
reg299006 := __e.Call(__defun__reverse, V4291)
reg299007 := Nil;
reg299008 := __e.Call(__defun__shen_4split__cc__rule, V4289, reg299006, reg299007)
reg299009 := PrimTail(V4290)
reg299010 := Nil;
reg299011 := __e.Call(__defun__shen_4split__cc__rules, V4289, reg299009, reg299010)
reg299012 := PrimCons(reg299008, reg299011)
__ctx.Return(reg299012)
return
} else {
reg299013 := PrimIsPair(V4290)
if reg299013 == True {
reg299014 := PrimTail(V4290)
reg299015 := PrimHead(V4290)
reg299016 := PrimCons(reg299015, V4291)
__ctx.TailApply(__defun__shen_4split__cc__rules, V4289, reg299014, reg299016)
return
} else {
reg299018 := MakeSymbol("shen.split_cc_rules")
__ctx.TailApply(__defun__shen_4f__error, reg299018)
return
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.split_cc_rules", value: __defun__shen_4split__cc__rules})

__defun__shen_4split__cc__rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4299 := __args[0]
_ = V4299
V4300 := __args[1]
_ = V4300
V4301 := __args[2]
_ = V4301
reg299020 := PrimIsPair(V4300)
var reg299044 Obj
if reg299020 == True {
reg299021 := MakeSymbol(":=")
reg299022 := PrimHead(V4300)
reg299023 := PrimEqual(reg299021, reg299022)
var reg299039 Obj
if reg299023 == True {
reg299024 := PrimTail(V4300)
reg299025 := PrimIsPair(reg299024)
var reg299034 Obj
if reg299025 == True {
reg299026 := Nil;
reg299027 := PrimTail(V4300)
reg299028 := PrimTail(reg299027)
reg299029 := PrimEqual(reg299026, reg299028)
var reg299032 Obj
if reg299029 == True {
reg299030 := True;
reg299032 = reg299030
} else {
reg299031 := False;
reg299032 = reg299031
}
reg299034 = reg299032
} else {
reg299033 := False;
reg299034 = reg299033
}
var reg299037 Obj
if reg299034 == True {
reg299035 := True;
reg299037 = reg299035
} else {
reg299036 := False;
reg299037 = reg299036
}
reg299039 = reg299037
} else {
reg299038 := False;
reg299039 = reg299038
}
var reg299042 Obj
if reg299039 == True {
reg299040 := True;
reg299042 = reg299040
} else {
reg299041 := False;
reg299042 = reg299041
}
reg299044 = reg299042
} else {
reg299043 := False;
reg299044 = reg299043
}
if reg299044 == True {
reg299045 := __e.Call(__defun__reverse, V4301)
reg299046 := PrimTail(V4300)
reg299047 := PrimCons(reg299045, reg299046)
__ctx.Return(reg299047)
return
} else {
reg299048 := PrimIsPair(V4300)
var reg299101 Obj
if reg299048 == True {
reg299049 := MakeSymbol(":=")
reg299050 := PrimHead(V4300)
reg299051 := PrimEqual(reg299049, reg299050)
var reg299096 Obj
if reg299051 == True {
reg299052 := PrimTail(V4300)
reg299053 := PrimIsPair(reg299052)
var reg299091 Obj
if reg299053 == True {
reg299054 := PrimTail(V4300)
reg299055 := PrimTail(reg299054)
reg299056 := PrimIsPair(reg299055)
var reg299086 Obj
if reg299056 == True {
reg299057 := MakeSymbol("where")
reg299058 := PrimTail(V4300)
reg299059 := PrimTail(reg299058)
reg299060 := PrimHead(reg299059)
reg299061 := PrimEqual(reg299057, reg299060)
var reg299081 Obj
if reg299061 == True {
reg299062 := PrimTail(V4300)
reg299063 := PrimTail(reg299062)
reg299064 := PrimTail(reg299063)
reg299065 := PrimIsPair(reg299064)
var reg299076 Obj
if reg299065 == True {
reg299066 := Nil;
reg299067 := PrimTail(V4300)
reg299068 := PrimTail(reg299067)
reg299069 := PrimTail(reg299068)
reg299070 := PrimTail(reg299069)
reg299071 := PrimEqual(reg299066, reg299070)
var reg299074 Obj
if reg299071 == True {
reg299072 := True;
reg299074 = reg299072
} else {
reg299073 := False;
reg299074 = reg299073
}
reg299076 = reg299074
} else {
reg299075 := False;
reg299076 = reg299075
}
var reg299079 Obj
if reg299076 == True {
reg299077 := True;
reg299079 = reg299077
} else {
reg299078 := False;
reg299079 = reg299078
}
reg299081 = reg299079
} else {
reg299080 := False;
reg299081 = reg299080
}
var reg299084 Obj
if reg299081 == True {
reg299082 := True;
reg299084 = reg299082
} else {
reg299083 := False;
reg299084 = reg299083
}
reg299086 = reg299084
} else {
reg299085 := False;
reg299086 = reg299085
}
var reg299089 Obj
if reg299086 == True {
reg299087 := True;
reg299089 = reg299087
} else {
reg299088 := False;
reg299089 = reg299088
}
reg299091 = reg299089
} else {
reg299090 := False;
reg299091 = reg299090
}
var reg299094 Obj
if reg299091 == True {
reg299092 := True;
reg299094 = reg299092
} else {
reg299093 := False;
reg299094 = reg299093
}
reg299096 = reg299094
} else {
reg299095 := False;
reg299096 = reg299095
}
var reg299099 Obj
if reg299096 == True {
reg299097 := True;
reg299099 = reg299097
} else {
reg299098 := False;
reg299099 = reg299098
}
reg299101 = reg299099
} else {
reg299100 := False;
reg299101 = reg299100
}
if reg299101 == True {
reg299102 := __e.Call(__defun__reverse, V4301)
reg299103 := MakeSymbol("where")
reg299104 := PrimTail(V4300)
reg299105 := PrimTail(reg299104)
reg299106 := PrimTail(reg299105)
reg299107 := PrimHead(reg299106)
reg299108 := PrimTail(V4300)
reg299109 := PrimHead(reg299108)
reg299110 := Nil;
reg299111 := PrimCons(reg299109, reg299110)
reg299112 := PrimCons(reg299107, reg299111)
reg299113 := PrimCons(reg299103, reg299112)
reg299114 := Nil;
reg299115 := PrimCons(reg299113, reg299114)
reg299116 := PrimCons(reg299102, reg299115)
__ctx.Return(reg299116)
return
} else {
reg299117 := Nil;
reg299118 := PrimEqual(reg299117, V4300)
if reg299118 == True {
reg299119 := __e.Call(__defun__shen_4semantic_1completion_1warning, V4299, V4301)
_ = reg299119
reg299120 := MakeSymbol(":=")
reg299121 := __e.Call(__defun__reverse, V4301)
reg299122 := __e.Call(__defun__shen_4default__semantics, reg299121)
reg299123 := Nil;
reg299124 := PrimCons(reg299122, reg299123)
reg299125 := PrimCons(reg299120, reg299124)
__ctx.TailApply(__defun__shen_4split__cc__rule, V4299, reg299125, V4301)
return
} else {
reg299127 := PrimIsPair(V4300)
if reg299127 == True {
reg299128 := PrimTail(V4300)
reg299129 := PrimHead(V4300)
reg299130 := PrimCons(reg299129, V4301)
__ctx.TailApply(__defun__shen_4split__cc__rule, V4299, reg299128, reg299130)
return
} else {
reg299132 := MakeSymbol("shen.split_cc_rule")
__ctx.TailApply(__defun__shen_4f__error, reg299132)
return
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.split_cc_rule", value: __defun__shen_4split__cc__rule})

__defun__shen_4semantic_1completion_1warning = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4312 := __args[0]
_ = V4312
V4313 := __args[1]
_ = V4313
reg299134 := True;
reg299135 := PrimEqual(reg299134, V4312)
if reg299135 == True {
reg299136 := MakeString("warning: ")
reg299137 := __e.Call(__defun__stoutput)
reg299138 := __e.Call(__defun__shen_4prhush, reg299136, reg299137)
_ = reg299138
reg299139 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg299140 := MakeString(" ")
reg299141 := MakeSymbol("shen.a")
reg299142 := __e.Call(__defun__shen_4app, X, reg299140, reg299141)
reg299143 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg299142, reg299143)
return
}, 1)
reg299145 := __e.Call(__defun__reverse, V4313)
reg299146 := __e.Call(__defun__shen_4for_1each, reg299139, reg299145)
_ = reg299146
reg299147 := MakeString("has no semantics.\n")
reg299148 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg299147, reg299148)
return
} else {
reg299150 := MakeSymbol("shen.skip")
__ctx.Return(reg299150)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.semantic-completion-warning", value: __defun__shen_4semantic_1completion_1warning})

__defun__shen_4default__semantics = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4315 := __args[0]
_ = V4315
reg299151 := Nil;
reg299152 := PrimEqual(reg299151, V4315)
if reg299152 == True {
reg299153 := Nil;
__ctx.Return(reg299153)
return
} else {
reg299154 := PrimIsPair(V4315)
var reg299169 Obj
if reg299154 == True {
reg299155 := Nil;
reg299156 := PrimTail(V4315)
reg299157 := PrimEqual(reg299155, reg299156)
var reg299164 Obj
if reg299157 == True {
reg299158 := PrimHead(V4315)
reg299159 := __e.Call(__defun__shen_4grammar__symbol_2, reg299158)
var reg299162 Obj
if reg299159 == True {
reg299160 := True;
reg299162 = reg299160
} else {
reg299161 := False;
reg299162 = reg299161
}
reg299164 = reg299162
} else {
reg299163 := False;
reg299164 = reg299163
}
var reg299167 Obj
if reg299164 == True {
reg299165 := True;
reg299167 = reg299165
} else {
reg299166 := False;
reg299167 = reg299166
}
reg299169 = reg299167
} else {
reg299168 := False;
reg299169 = reg299168
}
if reg299169 == True {
reg299170 := PrimHead(V4315)
__ctx.Return(reg299170)
return
} else {
reg299171 := PrimIsPair(V4315)
var reg299178 Obj
if reg299171 == True {
reg299172 := PrimHead(V4315)
reg299173 := __e.Call(__defun__shen_4grammar__symbol_2, reg299172)
var reg299176 Obj
if reg299173 == True {
reg299174 := True;
reg299176 = reg299174
} else {
reg299175 := False;
reg299176 = reg299175
}
reg299178 = reg299176
} else {
reg299177 := False;
reg299178 = reg299177
}
if reg299178 == True {
reg299179 := MakeSymbol("append")
reg299180 := PrimHead(V4315)
reg299181 := PrimTail(V4315)
reg299182 := __e.Call(__defun__shen_4default__semantics, reg299181)
reg299183 := Nil;
reg299184 := PrimCons(reg299182, reg299183)
reg299185 := PrimCons(reg299180, reg299184)
reg299186 := PrimCons(reg299179, reg299185)
__ctx.Return(reg299186)
return
} else {
reg299187 := PrimIsPair(V4315)
if reg299187 == True {
reg299188 := MakeSymbol("cons")
reg299189 := PrimHead(V4315)
reg299190 := PrimTail(V4315)
reg299191 := __e.Call(__defun__shen_4default__semantics, reg299190)
reg299192 := Nil;
reg299193 := PrimCons(reg299191, reg299192)
reg299194 := PrimCons(reg299189, reg299193)
reg299195 := PrimCons(reg299188, reg299194)
__ctx.Return(reg299195)
return
} else {
reg299196 := MakeSymbol("shen.default_semantics")
__ctx.TailApply(__defun__shen_4f__error, reg299196)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.default_semantics", value: __defun__shen_4default__semantics})

__defun__shen_4grammar__symbol_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4317 := __args[0]
_ = V4317
reg299198 := PrimIsSymbol(V4317)
if reg299198 == True {
reg299199 := __e.Call(__defun__explode, V4317)
reg299200 := __e.Call(__defun__shen_4strip_1pathname, reg299199)
Cs := reg299200
_ = Cs
reg299201 := PrimHead(Cs)
reg299202 := MakeString("<")
reg299203 := PrimEqual(reg299201, reg299202)
var reg299212 Obj
if reg299203 == True {
reg299204 := __e.Call(__defun__reverse, Cs)
reg299205 := PrimHead(reg299204)
reg299206 := MakeString(">")
reg299207 := PrimEqual(reg299205, reg299206)
var reg299210 Obj
if reg299207 == True {
reg299208 := True;
reg299210 = reg299208
} else {
reg299209 := False;
reg299210 = reg299209
}
reg299212 = reg299210
} else {
reg299211 := False;
reg299212 = reg299211
}
if reg299212 == True {
reg299213 := True;
__ctx.Return(reg299213)
return
} else {
reg299214 := False;
__ctx.Return(reg299214)
return
}
} else {
reg299215 := False;
__ctx.Return(reg299215)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.grammar_symbol?", value: __defun__shen_4grammar__symbol_2})

__defun__shen_4yacc__cases = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4319 := __args[0]
_ = V4319
reg299216 := PrimIsPair(V4319)
var reg299224 Obj
if reg299216 == True {
reg299217 := Nil;
reg299218 := PrimTail(V4319)
reg299219 := PrimEqual(reg299217, reg299218)
var reg299222 Obj
if reg299219 == True {
reg299220 := True;
reg299222 = reg299220
} else {
reg299221 := False;
reg299222 = reg299221
}
reg299224 = reg299222
} else {
reg299223 := False;
reg299224 = reg299223
}
if reg299224 == True {
reg299225 := PrimHead(V4319)
__ctx.Return(reg299225)
return
} else {
reg299226 := PrimIsPair(V4319)
if reg299226 == True {
reg299227 := MakeSymbol("YaccParse")
P := reg299227
_ = P
reg299228 := MakeSymbol("let")
reg299229 := PrimHead(V4319)
reg299230 := MakeSymbol("if")
reg299231 := MakeSymbol("=")
reg299232 := MakeSymbol("fail")
reg299233 := Nil;
reg299234 := PrimCons(reg299232, reg299233)
reg299235 := Nil;
reg299236 := PrimCons(reg299234, reg299235)
reg299237 := PrimCons(P, reg299236)
reg299238 := PrimCons(reg299231, reg299237)
reg299239 := PrimTail(V4319)
reg299240 := __e.Call(__defun__shen_4yacc__cases, reg299239)
reg299241 := Nil;
reg299242 := PrimCons(P, reg299241)
reg299243 := PrimCons(reg299240, reg299242)
reg299244 := PrimCons(reg299238, reg299243)
reg299245 := PrimCons(reg299230, reg299244)
reg299246 := Nil;
reg299247 := PrimCons(reg299245, reg299246)
reg299248 := PrimCons(reg299229, reg299247)
reg299249 := PrimCons(P, reg299248)
reg299250 := PrimCons(reg299228, reg299249)
__ctx.Return(reg299250)
return
} else {
reg299251 := MakeSymbol("shen.yacc_cases")
__ctx.TailApply(__defun__shen_4f__error, reg299251)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.yacc_cases", value: __defun__shen_4yacc__cases})

__defun__shen_4cc__body = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4321 := __args[0]
_ = V4321
reg299253 := PrimIsPair(V4321)
var reg299269 Obj
if reg299253 == True {
reg299254 := PrimTail(V4321)
reg299255 := PrimIsPair(reg299254)
var reg299264 Obj
if reg299255 == True {
reg299256 := Nil;
reg299257 := PrimTail(V4321)
reg299258 := PrimTail(reg299257)
reg299259 := PrimEqual(reg299256, reg299258)
var reg299262 Obj
if reg299259 == True {
reg299260 := True;
reg299262 = reg299260
} else {
reg299261 := False;
reg299262 = reg299261
}
reg299264 = reg299262
} else {
reg299263 := False;
reg299264 = reg299263
}
var reg299267 Obj
if reg299264 == True {
reg299265 := True;
reg299267 = reg299265
} else {
reg299266 := False;
reg299267 = reg299266
}
reg299269 = reg299267
} else {
reg299268 := False;
reg299269 = reg299268
}
if reg299269 == True {
reg299270 := PrimHead(V4321)
reg299271 := MakeSymbol("Stream")
reg299272 := PrimTail(V4321)
reg299273 := PrimHead(reg299272)
__ctx.TailApply(__defun__shen_4syntax, reg299270, reg299271, reg299273)
return
} else {
reg299275 := MakeSymbol("shen.cc_body")
__ctx.TailApply(__defun__shen_4f__error, reg299275)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.cc_body", value: __defun__shen_4cc__body})

__defun__shen_4syntax = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4325 := __args[0]
_ = V4325
V4326 := __args[1]
_ = V4326
V4327 := __args[2]
_ = V4327
reg299277 := Nil;
reg299278 := PrimEqual(reg299277, V4325)
var reg299317 Obj
if reg299278 == True {
reg299279 := PrimIsPair(V4327)
var reg299312 Obj
if reg299279 == True {
reg299280 := MakeSymbol("where")
reg299281 := PrimHead(V4327)
reg299282 := PrimEqual(reg299280, reg299281)
var reg299307 Obj
if reg299282 == True {
reg299283 := PrimTail(V4327)
reg299284 := PrimIsPair(reg299283)
var reg299302 Obj
if reg299284 == True {
reg299285 := PrimTail(V4327)
reg299286 := PrimTail(reg299285)
reg299287 := PrimIsPair(reg299286)
var reg299297 Obj
if reg299287 == True {
reg299288 := Nil;
reg299289 := PrimTail(V4327)
reg299290 := PrimTail(reg299289)
reg299291 := PrimTail(reg299290)
reg299292 := PrimEqual(reg299288, reg299291)
var reg299295 Obj
if reg299292 == True {
reg299293 := True;
reg299295 = reg299293
} else {
reg299294 := False;
reg299295 = reg299294
}
reg299297 = reg299295
} else {
reg299296 := False;
reg299297 = reg299296
}
var reg299300 Obj
if reg299297 == True {
reg299298 := True;
reg299300 = reg299298
} else {
reg299299 := False;
reg299300 = reg299299
}
reg299302 = reg299300
} else {
reg299301 := False;
reg299302 = reg299301
}
var reg299305 Obj
if reg299302 == True {
reg299303 := True;
reg299305 = reg299303
} else {
reg299304 := False;
reg299305 = reg299304
}
reg299307 = reg299305
} else {
reg299306 := False;
reg299307 = reg299306
}
var reg299310 Obj
if reg299307 == True {
reg299308 := True;
reg299310 = reg299308
} else {
reg299309 := False;
reg299310 = reg299309
}
reg299312 = reg299310
} else {
reg299311 := False;
reg299312 = reg299311
}
var reg299315 Obj
if reg299312 == True {
reg299313 := True;
reg299315 = reg299313
} else {
reg299314 := False;
reg299315 = reg299314
}
reg299317 = reg299315
} else {
reg299316 := False;
reg299317 = reg299316
}
if reg299317 == True {
reg299318 := MakeSymbol("if")
reg299319 := PrimTail(V4327)
reg299320 := PrimHead(reg299319)
reg299321 := __e.Call(__defun__shen_4semantics, reg299320)
reg299322 := MakeSymbol("shen.pair")
reg299323 := MakeSymbol("hd")
reg299324 := Nil;
reg299325 := PrimCons(V4326, reg299324)
reg299326 := PrimCons(reg299323, reg299325)
reg299327 := PrimTail(V4327)
reg299328 := PrimTail(reg299327)
reg299329 := PrimHead(reg299328)
reg299330 := __e.Call(__defun__shen_4semantics, reg299329)
reg299331 := Nil;
reg299332 := PrimCons(reg299330, reg299331)
reg299333 := PrimCons(reg299326, reg299332)
reg299334 := PrimCons(reg299322, reg299333)
reg299335 := MakeSymbol("fail")
reg299336 := Nil;
reg299337 := PrimCons(reg299335, reg299336)
reg299338 := Nil;
reg299339 := PrimCons(reg299337, reg299338)
reg299340 := PrimCons(reg299334, reg299339)
reg299341 := PrimCons(reg299321, reg299340)
reg299342 := PrimCons(reg299318, reg299341)
__ctx.Return(reg299342)
return
} else {
reg299343 := Nil;
reg299344 := PrimEqual(reg299343, V4325)
if reg299344 == True {
reg299345 := MakeSymbol("shen.pair")
reg299346 := MakeSymbol("hd")
reg299347 := Nil;
reg299348 := PrimCons(V4326, reg299347)
reg299349 := PrimCons(reg299346, reg299348)
reg299350 := __e.Call(__defun__shen_4semantics, V4327)
reg299351 := Nil;
reg299352 := PrimCons(reg299350, reg299351)
reg299353 := PrimCons(reg299349, reg299352)
reg299354 := PrimCons(reg299345, reg299353)
__ctx.Return(reg299354)
return
} else {
reg299355 := PrimIsPair(V4325)
if reg299355 == True {
reg299356 := PrimHead(V4325)
reg299357 := __e.Call(__defun__shen_4grammar__symbol_2, reg299356)
if reg299357 == True {
__ctx.TailApply(__defun__shen_4recursive__descent, V4325, V4326, V4327)
return
} else {
reg299359 := PrimHead(V4325)
reg299360 := PrimIsVariable(reg299359)
if reg299360 == True {
__ctx.TailApply(__defun__shen_4variable_1match, V4325, V4326, V4327)
return
} else {
reg299362 := PrimHead(V4325)
reg299363 := __e.Call(__defun__shen_4jump__stream_2, reg299362)
if reg299363 == True {
__ctx.TailApply(__defun__shen_4jump__stream, V4325, V4326, V4327)
return
} else {
reg299365 := PrimHead(V4325)
reg299366 := __e.Call(__defun__shen_4terminal_2, reg299365)
if reg299366 == True {
__ctx.TailApply(__defun__shen_4check__stream, V4325, V4326, V4327)
return
} else {
reg299368 := PrimHead(V4325)
reg299369 := PrimIsPair(reg299368)
if reg299369 == True {
reg299370 := PrimHead(V4325)
reg299371 := __e.Call(__defun__shen_4decons, reg299370)
reg299372 := PrimTail(V4325)
__ctx.TailApply(__defun__shen_4list_1stream, reg299371, reg299372, V4326, V4327)
return
} else {
reg299374 := PrimHead(V4325)
reg299375 := MakeString(" is not legal syntax\n")
reg299376 := MakeSymbol("shen.a")
reg299377 := __e.Call(__defun__shen_4app, reg299374, reg299375, reg299376)
reg299378 := PrimSimpleError(reg299377)
__ctx.Return(reg299378)
return
}
}
}
}
}
} else {
reg299379 := MakeSymbol("shen.syntax")
__ctx.TailApply(__defun__shen_4f__error, reg299379)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.syntax", value: __defun__shen_4syntax})

__defun__shen_4list_1stream = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4332 := __args[0]
_ = V4332
V4333 := __args[1]
_ = V4333
V4334 := __args[2]
_ = V4334
V4335 := __args[3]
_ = V4335
reg299381 := MakeSymbol("and")
reg299382 := MakeSymbol("cons?")
reg299383 := MakeSymbol("hd")
reg299384 := Nil;
reg299385 := PrimCons(V4334, reg299384)
reg299386 := PrimCons(reg299383, reg299385)
reg299387 := Nil;
reg299388 := PrimCons(reg299386, reg299387)
reg299389 := PrimCons(reg299382, reg299388)
reg299390 := MakeSymbol("cons?")
reg299391 := MakeSymbol("shen.hdhd")
reg299392 := Nil;
reg299393 := PrimCons(V4334, reg299392)
reg299394 := PrimCons(reg299391, reg299393)
reg299395 := Nil;
reg299396 := PrimCons(reg299394, reg299395)
reg299397 := PrimCons(reg299390, reg299396)
reg299398 := Nil;
reg299399 := PrimCons(reg299397, reg299398)
reg299400 := PrimCons(reg299389, reg299399)
reg299401 := PrimCons(reg299381, reg299400)
Test := reg299401
_ = Test
reg299402 := MakeSymbol("shen.place")
reg299403 := __e.Call(__defun__gensym, reg299402)
Placeholder := reg299403
_ = Placeholder
reg299404 := MakeSymbol("shen.pair")
reg299405 := MakeSymbol("shen.tlhd")
reg299406 := Nil;
reg299407 := PrimCons(V4334, reg299406)
reg299408 := PrimCons(reg299405, reg299407)
reg299409 := MakeSymbol("shen.hdtl")
reg299410 := Nil;
reg299411 := PrimCons(V4334, reg299410)
reg299412 := PrimCons(reg299409, reg299411)
reg299413 := Nil;
reg299414 := PrimCons(reg299412, reg299413)
reg299415 := PrimCons(reg299408, reg299414)
reg299416 := PrimCons(reg299404, reg299415)
reg299417 := __e.Call(__defun__shen_4syntax, V4333, reg299416, V4335)
RunOn := reg299417
_ = RunOn
reg299418 := MakeSymbol("shen.pair")
reg299419 := MakeSymbol("shen.hdhd")
reg299420 := Nil;
reg299421 := PrimCons(V4334, reg299420)
reg299422 := PrimCons(reg299419, reg299421)
reg299423 := MakeSymbol("shen.hdtl")
reg299424 := Nil;
reg299425 := PrimCons(V4334, reg299424)
reg299426 := PrimCons(reg299423, reg299425)
reg299427 := Nil;
reg299428 := PrimCons(reg299426, reg299427)
reg299429 := PrimCons(reg299422, reg299428)
reg299430 := PrimCons(reg299418, reg299429)
reg299431 := __e.Call(__defun__shen_4syntax, V4332, reg299430, Placeholder)
reg299432 := __e.Call(__defun__shen_4insert_1runon, RunOn, Placeholder, reg299431)
Action := reg299432
_ = Action
reg299433 := MakeSymbol("if")
reg299434 := MakeSymbol("fail")
reg299435 := Nil;
reg299436 := PrimCons(reg299434, reg299435)
reg299437 := Nil;
reg299438 := PrimCons(reg299436, reg299437)
reg299439 := PrimCons(Action, reg299438)
reg299440 := PrimCons(Test, reg299439)
reg299441 := PrimCons(reg299433, reg299440)
__ctx.Return(reg299441)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.list-stream", value: __defun__shen_4list_1stream})

__defun__shen_4decons = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4337 := __args[0]
_ = V4337
reg299442 := PrimIsPair(V4337)
var reg299485 Obj
if reg299442 == True {
reg299443 := MakeSymbol("cons")
reg299444 := PrimHead(V4337)
reg299445 := PrimEqual(reg299443, reg299444)
var reg299480 Obj
if reg299445 == True {
reg299446 := PrimTail(V4337)
reg299447 := PrimIsPair(reg299446)
var reg299475 Obj
if reg299447 == True {
reg299448 := PrimTail(V4337)
reg299449 := PrimTail(reg299448)
reg299450 := PrimIsPair(reg299449)
var reg299470 Obj
if reg299450 == True {
reg299451 := Nil;
reg299452 := PrimTail(V4337)
reg299453 := PrimTail(reg299452)
reg299454 := PrimHead(reg299453)
reg299455 := PrimEqual(reg299451, reg299454)
var reg299465 Obj
if reg299455 == True {
reg299456 := Nil;
reg299457 := PrimTail(V4337)
reg299458 := PrimTail(reg299457)
reg299459 := PrimTail(reg299458)
reg299460 := PrimEqual(reg299456, reg299459)
var reg299463 Obj
if reg299460 == True {
reg299461 := True;
reg299463 = reg299461
} else {
reg299462 := False;
reg299463 = reg299462
}
reg299465 = reg299463
} else {
reg299464 := False;
reg299465 = reg299464
}
var reg299468 Obj
if reg299465 == True {
reg299466 := True;
reg299468 = reg299466
} else {
reg299467 := False;
reg299468 = reg299467
}
reg299470 = reg299468
} else {
reg299469 := False;
reg299470 = reg299469
}
var reg299473 Obj
if reg299470 == True {
reg299471 := True;
reg299473 = reg299471
} else {
reg299472 := False;
reg299473 = reg299472
}
reg299475 = reg299473
} else {
reg299474 := False;
reg299475 = reg299474
}
var reg299478 Obj
if reg299475 == True {
reg299476 := True;
reg299478 = reg299476
} else {
reg299477 := False;
reg299478 = reg299477
}
reg299480 = reg299478
} else {
reg299479 := False;
reg299480 = reg299479
}
var reg299483 Obj
if reg299480 == True {
reg299481 := True;
reg299483 = reg299481
} else {
reg299482 := False;
reg299483 = reg299482
}
reg299485 = reg299483
} else {
reg299484 := False;
reg299485 = reg299484
}
if reg299485 == True {
reg299486 := PrimTail(V4337)
reg299487 := PrimHead(reg299486)
reg299488 := Nil;
reg299489 := PrimCons(reg299487, reg299488)
__ctx.Return(reg299489)
return
} else {
reg299490 := PrimIsPair(V4337)
var reg299523 Obj
if reg299490 == True {
reg299491 := MakeSymbol("cons")
reg299492 := PrimHead(V4337)
reg299493 := PrimEqual(reg299491, reg299492)
var reg299518 Obj
if reg299493 == True {
reg299494 := PrimTail(V4337)
reg299495 := PrimIsPair(reg299494)
var reg299513 Obj
if reg299495 == True {
reg299496 := PrimTail(V4337)
reg299497 := PrimTail(reg299496)
reg299498 := PrimIsPair(reg299497)
var reg299508 Obj
if reg299498 == True {
reg299499 := Nil;
reg299500 := PrimTail(V4337)
reg299501 := PrimTail(reg299500)
reg299502 := PrimTail(reg299501)
reg299503 := PrimEqual(reg299499, reg299502)
var reg299506 Obj
if reg299503 == True {
reg299504 := True;
reg299506 = reg299504
} else {
reg299505 := False;
reg299506 = reg299505
}
reg299508 = reg299506
} else {
reg299507 := False;
reg299508 = reg299507
}
var reg299511 Obj
if reg299508 == True {
reg299509 := True;
reg299511 = reg299509
} else {
reg299510 := False;
reg299511 = reg299510
}
reg299513 = reg299511
} else {
reg299512 := False;
reg299513 = reg299512
}
var reg299516 Obj
if reg299513 == True {
reg299514 := True;
reg299516 = reg299514
} else {
reg299515 := False;
reg299516 = reg299515
}
reg299518 = reg299516
} else {
reg299517 := False;
reg299518 = reg299517
}
var reg299521 Obj
if reg299518 == True {
reg299519 := True;
reg299521 = reg299519
} else {
reg299520 := False;
reg299521 = reg299520
}
reg299523 = reg299521
} else {
reg299522 := False;
reg299523 = reg299522
}
if reg299523 == True {
reg299524 := PrimTail(V4337)
reg299525 := PrimHead(reg299524)
reg299526 := PrimTail(V4337)
reg299527 := PrimTail(reg299526)
reg299528 := PrimHead(reg299527)
reg299529 := __e.Call(__defun__shen_4decons, reg299528)
reg299530 := PrimCons(reg299525, reg299529)
__ctx.Return(reg299530)
return
} else {
__ctx.Return(V4337)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.decons", value: __defun__shen_4decons})

__defun__shen_4insert_1runon = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4352 := __args[0]
_ = V4352
V4353 := __args[1]
_ = V4353
V4354 := __args[2]
_ = V4354
reg299531 := PrimIsPair(V4354)
var reg299573 Obj
if reg299531 == True {
reg299532 := MakeSymbol("shen.pair")
reg299533 := PrimHead(V4354)
reg299534 := PrimEqual(reg299532, reg299533)
var reg299568 Obj
if reg299534 == True {
reg299535 := PrimTail(V4354)
reg299536 := PrimIsPair(reg299535)
var reg299563 Obj
if reg299536 == True {
reg299537 := PrimTail(V4354)
reg299538 := PrimTail(reg299537)
reg299539 := PrimIsPair(reg299538)
var reg299558 Obj
if reg299539 == True {
reg299540 := Nil;
reg299541 := PrimTail(V4354)
reg299542 := PrimTail(reg299541)
reg299543 := PrimTail(reg299542)
reg299544 := PrimEqual(reg299540, reg299543)
var reg299553 Obj
if reg299544 == True {
reg299545 := PrimTail(V4354)
reg299546 := PrimTail(reg299545)
reg299547 := PrimHead(reg299546)
reg299548 := PrimEqual(reg299547, V4353)
var reg299551 Obj
if reg299548 == True {
reg299549 := True;
reg299551 = reg299549
} else {
reg299550 := False;
reg299551 = reg299550
}
reg299553 = reg299551
} else {
reg299552 := False;
reg299553 = reg299552
}
var reg299556 Obj
if reg299553 == True {
reg299554 := True;
reg299556 = reg299554
} else {
reg299555 := False;
reg299556 = reg299555
}
reg299558 = reg299556
} else {
reg299557 := False;
reg299558 = reg299557
}
var reg299561 Obj
if reg299558 == True {
reg299559 := True;
reg299561 = reg299559
} else {
reg299560 := False;
reg299561 = reg299560
}
reg299563 = reg299561
} else {
reg299562 := False;
reg299563 = reg299562
}
var reg299566 Obj
if reg299563 == True {
reg299564 := True;
reg299566 = reg299564
} else {
reg299565 := False;
reg299566 = reg299565
}
reg299568 = reg299566
} else {
reg299567 := False;
reg299568 = reg299567
}
var reg299571 Obj
if reg299568 == True {
reg299569 := True;
reg299571 = reg299569
} else {
reg299570 := False;
reg299571 = reg299570
}
reg299573 = reg299571
} else {
reg299572 := False;
reg299573 = reg299572
}
if reg299573 == True {
__ctx.Return(V4352)
return
} else {
reg299574 := PrimIsPair(V4354)
if reg299574 == True {
reg299575 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4insert_1runon, V4352, V4353, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg299575, V4354)
return
} else {
__ctx.Return(V4354)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.insert-runon", value: __defun__shen_4insert_1runon})

__defun__shen_4strip_1pathname = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4360 := __args[0]
_ = V4360
reg299578 := MakeString(".")
reg299579 := __e.Call(__defun__element_2, reg299578, V4360)
reg299580 := PrimNot(reg299579)
if reg299580 == True {
__ctx.Return(V4360)
return
} else {
reg299581 := PrimIsPair(V4360)
if reg299581 == True {
reg299582 := PrimTail(V4360)
__ctx.TailApply(__defun__shen_4strip_1pathname, reg299582)
return
} else {
reg299584 := MakeSymbol("shen.strip-pathname")
__ctx.TailApply(__defun__shen_4f__error, reg299584)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.strip-pathname", value: __defun__shen_4strip_1pathname})

__defun__shen_4recursive__descent = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4364 := __args[0]
_ = V4364
V4365 := __args[1]
_ = V4365
V4366 := __args[2]
_ = V4366
reg299586 := PrimIsPair(V4364)
if reg299586 == True {
reg299587 := PrimHead(V4364)
reg299588 := Nil;
reg299589 := PrimCons(V4365, reg299588)
reg299590 := PrimCons(reg299587, reg299589)
Test := reg299590
_ = Test
reg299591 := PrimTail(V4364)
reg299592 := MakeSymbol("Parse_")
reg299593 := PrimHead(V4364)
reg299594 := __e.Call(__defun__concat, reg299592, reg299593)
reg299595 := __e.Call(__defun__shen_4syntax, reg299591, reg299594, V4366)
Action := reg299595
_ = Action
reg299596 := MakeSymbol("fail")
reg299597 := Nil;
reg299598 := PrimCons(reg299596, reg299597)
Else := reg299598
_ = Else
reg299599 := MakeSymbol("let")
reg299600 := MakeSymbol("Parse_")
reg299601 := PrimHead(V4364)
reg299602 := __e.Call(__defun__concat, reg299600, reg299601)
reg299603 := MakeSymbol("if")
reg299604 := MakeSymbol("not")
reg299605 := MakeSymbol("=")
reg299606 := MakeSymbol("fail")
reg299607 := Nil;
reg299608 := PrimCons(reg299606, reg299607)
reg299609 := MakeSymbol("Parse_")
reg299610 := PrimHead(V4364)
reg299611 := __e.Call(__defun__concat, reg299609, reg299610)
reg299612 := Nil;
reg299613 := PrimCons(reg299611, reg299612)
reg299614 := PrimCons(reg299608, reg299613)
reg299615 := PrimCons(reg299605, reg299614)
reg299616 := Nil;
reg299617 := PrimCons(reg299615, reg299616)
reg299618 := PrimCons(reg299604, reg299617)
reg299619 := Nil;
reg299620 := PrimCons(Else, reg299619)
reg299621 := PrimCons(Action, reg299620)
reg299622 := PrimCons(reg299618, reg299621)
reg299623 := PrimCons(reg299603, reg299622)
reg299624 := Nil;
reg299625 := PrimCons(reg299623, reg299624)
reg299626 := PrimCons(Test, reg299625)
reg299627 := PrimCons(reg299602, reg299626)
reg299628 := PrimCons(reg299599, reg299627)
__ctx.Return(reg299628)
return
} else {
reg299629 := MakeSymbol("shen.recursive_descent")
__ctx.TailApply(__defun__shen_4f__error, reg299629)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.recursive_descent", value: __defun__shen_4recursive__descent})

__defun__shen_4variable_1match = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4370 := __args[0]
_ = V4370
V4371 := __args[1]
_ = V4371
V4372 := __args[2]
_ = V4372
reg299631 := PrimIsPair(V4370)
if reg299631 == True {
reg299632 := MakeSymbol("cons?")
reg299633 := MakeSymbol("hd")
reg299634 := Nil;
reg299635 := PrimCons(V4371, reg299634)
reg299636 := PrimCons(reg299633, reg299635)
reg299637 := Nil;
reg299638 := PrimCons(reg299636, reg299637)
reg299639 := PrimCons(reg299632, reg299638)
Test := reg299639
_ = Test
reg299640 := MakeSymbol("let")
reg299641 := MakeSymbol("Parse_")
reg299642 := PrimHead(V4370)
reg299643 := __e.Call(__defun__concat, reg299641, reg299642)
reg299644 := MakeSymbol("shen.hdhd")
reg299645 := Nil;
reg299646 := PrimCons(V4371, reg299645)
reg299647 := PrimCons(reg299644, reg299646)
reg299648 := PrimTail(V4370)
reg299649 := MakeSymbol("shen.pair")
reg299650 := MakeSymbol("shen.tlhd")
reg299651 := Nil;
reg299652 := PrimCons(V4371, reg299651)
reg299653 := PrimCons(reg299650, reg299652)
reg299654 := MakeSymbol("shen.hdtl")
reg299655 := Nil;
reg299656 := PrimCons(V4371, reg299655)
reg299657 := PrimCons(reg299654, reg299656)
reg299658 := Nil;
reg299659 := PrimCons(reg299657, reg299658)
reg299660 := PrimCons(reg299653, reg299659)
reg299661 := PrimCons(reg299649, reg299660)
reg299662 := __e.Call(__defun__shen_4syntax, reg299648, reg299661, V4372)
reg299663 := Nil;
reg299664 := PrimCons(reg299662, reg299663)
reg299665 := PrimCons(reg299647, reg299664)
reg299666 := PrimCons(reg299643, reg299665)
reg299667 := PrimCons(reg299640, reg299666)
Action := reg299667
_ = Action
reg299668 := MakeSymbol("fail")
reg299669 := Nil;
reg299670 := PrimCons(reg299668, reg299669)
Else := reg299670
_ = Else
reg299671 := MakeSymbol("if")
reg299672 := Nil;
reg299673 := PrimCons(Else, reg299672)
reg299674 := PrimCons(Action, reg299673)
reg299675 := PrimCons(Test, reg299674)
reg299676 := PrimCons(reg299671, reg299675)
__ctx.Return(reg299676)
return
} else {
reg299677 := MakeSymbol("shen.variable-match")
__ctx.TailApply(__defun__shen_4f__error, reg299677)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.variable-match", value: __defun__shen_4variable_1match})

__defun__shen_4terminal_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4382 := __args[0]
_ = V4382
reg299679 := PrimIsPair(V4382)
if reg299679 == True {
reg299680 := False;
__ctx.Return(reg299680)
return
} else {
reg299681 := PrimIsVariable(V4382)
if reg299681 == True {
reg299682 := False;
__ctx.Return(reg299682)
return
} else {
reg299683 := True;
__ctx.Return(reg299683)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.terminal?", value: __defun__shen_4terminal_2})

__defun__shen_4jump__stream_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4388 := __args[0]
_ = V4388
reg299684 := MakeSymbol("_")
reg299685 := PrimEqual(V4388, reg299684)
if reg299685 == True {
reg299686 := True;
__ctx.Return(reg299686)
return
} else {
reg299687 := False;
__ctx.Return(reg299687)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.jump_stream?", value: __defun__shen_4jump__stream_2})

__defun__shen_4check__stream = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4392 := __args[0]
_ = V4392
V4393 := __args[1]
_ = V4393
V4394 := __args[2]
_ = V4394
reg299688 := PrimIsPair(V4392)
if reg299688 == True {
reg299689 := MakeSymbol("and")
reg299690 := MakeSymbol("cons?")
reg299691 := MakeSymbol("hd")
reg299692 := Nil;
reg299693 := PrimCons(V4393, reg299692)
reg299694 := PrimCons(reg299691, reg299693)
reg299695 := Nil;
reg299696 := PrimCons(reg299694, reg299695)
reg299697 := PrimCons(reg299690, reg299696)
reg299698 := MakeSymbol("=")
reg299699 := PrimHead(V4392)
reg299700 := MakeSymbol("shen.hdhd")
reg299701 := Nil;
reg299702 := PrimCons(V4393, reg299701)
reg299703 := PrimCons(reg299700, reg299702)
reg299704 := Nil;
reg299705 := PrimCons(reg299703, reg299704)
reg299706 := PrimCons(reg299699, reg299705)
reg299707 := PrimCons(reg299698, reg299706)
reg299708 := Nil;
reg299709 := PrimCons(reg299707, reg299708)
reg299710 := PrimCons(reg299697, reg299709)
reg299711 := PrimCons(reg299689, reg299710)
Test := reg299711
_ = Test
reg299712 := MakeSymbol("NewStream")
reg299713 := __e.Call(__defun__gensym, reg299712)
NewStr := reg299713
_ = NewStr
reg299714 := MakeSymbol("let")
reg299715 := MakeSymbol("shen.pair")
reg299716 := MakeSymbol("shen.tlhd")
reg299717 := Nil;
reg299718 := PrimCons(V4393, reg299717)
reg299719 := PrimCons(reg299716, reg299718)
reg299720 := MakeSymbol("shen.hdtl")
reg299721 := Nil;
reg299722 := PrimCons(V4393, reg299721)
reg299723 := PrimCons(reg299720, reg299722)
reg299724 := Nil;
reg299725 := PrimCons(reg299723, reg299724)
reg299726 := PrimCons(reg299719, reg299725)
reg299727 := PrimCons(reg299715, reg299726)
reg299728 := PrimTail(V4392)
reg299729 := __e.Call(__defun__shen_4syntax, reg299728, NewStr, V4394)
reg299730 := Nil;
reg299731 := PrimCons(reg299729, reg299730)
reg299732 := PrimCons(reg299727, reg299731)
reg299733 := PrimCons(NewStr, reg299732)
reg299734 := PrimCons(reg299714, reg299733)
Action := reg299734
_ = Action
reg299735 := MakeSymbol("fail")
reg299736 := Nil;
reg299737 := PrimCons(reg299735, reg299736)
Else := reg299737
_ = Else
reg299738 := MakeSymbol("if")
reg299739 := Nil;
reg299740 := PrimCons(Else, reg299739)
reg299741 := PrimCons(Action, reg299740)
reg299742 := PrimCons(Test, reg299741)
reg299743 := PrimCons(reg299738, reg299742)
__ctx.Return(reg299743)
return
} else {
reg299744 := MakeSymbol("shen.check_stream")
__ctx.TailApply(__defun__shen_4f__error, reg299744)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.check_stream", value: __defun__shen_4check__stream})

__defun__shen_4jump__stream = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4398 := __args[0]
_ = V4398
V4399 := __args[1]
_ = V4399
V4400 := __args[2]
_ = V4400
reg299746 := PrimIsPair(V4398)
if reg299746 == True {
reg299747 := MakeSymbol("cons?")
reg299748 := MakeSymbol("hd")
reg299749 := Nil;
reg299750 := PrimCons(V4399, reg299749)
reg299751 := PrimCons(reg299748, reg299750)
reg299752 := Nil;
reg299753 := PrimCons(reg299751, reg299752)
reg299754 := PrimCons(reg299747, reg299753)
Test := reg299754
_ = Test
reg299755 := PrimTail(V4398)
reg299756 := MakeSymbol("shen.pair")
reg299757 := MakeSymbol("shen.tlhd")
reg299758 := Nil;
reg299759 := PrimCons(V4399, reg299758)
reg299760 := PrimCons(reg299757, reg299759)
reg299761 := MakeSymbol("shen.hdtl")
reg299762 := Nil;
reg299763 := PrimCons(V4399, reg299762)
reg299764 := PrimCons(reg299761, reg299763)
reg299765 := Nil;
reg299766 := PrimCons(reg299764, reg299765)
reg299767 := PrimCons(reg299760, reg299766)
reg299768 := PrimCons(reg299756, reg299767)
reg299769 := __e.Call(__defun__shen_4syntax, reg299755, reg299768, V4400)
Action := reg299769
_ = Action
reg299770 := MakeSymbol("fail")
reg299771 := Nil;
reg299772 := PrimCons(reg299770, reg299771)
Else := reg299772
_ = Else
reg299773 := MakeSymbol("if")
reg299774 := Nil;
reg299775 := PrimCons(Else, reg299774)
reg299776 := PrimCons(Action, reg299775)
reg299777 := PrimCons(Test, reg299776)
reg299778 := PrimCons(reg299773, reg299777)
__ctx.Return(reg299778)
return
} else {
reg299779 := MakeSymbol("shen.jump_stream")
__ctx.TailApply(__defun__shen_4f__error, reg299779)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.jump_stream", value: __defun__shen_4jump__stream})

__defun__shen_4semantics = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4402 := __args[0]
_ = V4402
reg299781 := Nil;
reg299782 := PrimEqual(reg299781, V4402)
if reg299782 == True {
reg299783 := Nil;
__ctx.Return(reg299783)
return
} else {
reg299784 := __e.Call(__defun__shen_4grammar__symbol_2, V4402)
if reg299784 == True {
reg299785 := MakeSymbol("shen.hdtl")
reg299786 := MakeSymbol("Parse_")
reg299787 := __e.Call(__defun__concat, reg299786, V4402)
reg299788 := Nil;
reg299789 := PrimCons(reg299787, reg299788)
reg299790 := PrimCons(reg299785, reg299789)
__ctx.Return(reg299790)
return
} else {
reg299791 := PrimIsVariable(V4402)
if reg299791 == True {
reg299792 := MakeSymbol("Parse_")
__ctx.TailApply(__defun__concat, reg299792, V4402)
return
} else {
reg299794 := PrimIsPair(V4402)
if reg299794 == True {
reg299795 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4semantics, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg299795, V4402)
return
} else {
__ctx.Return(V4402)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.semantics", value: __defun__shen_4semantics})

__defun__shen_4pair = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4405 := __args[0]
_ = V4405
V4406 := __args[1]
_ = V4406
reg299798 := Nil;
reg299799 := PrimCons(V4406, reg299798)
reg299800 := PrimCons(V4405, reg299799)
__ctx.Return(reg299800)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.pair", value: __defun__shen_4pair})

__defun__shen_4hdtl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4408 := __args[0]
_ = V4408
reg299801 := PrimTail(V4408)
reg299802 := PrimHead(reg299801)
__ctx.Return(reg299802)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.hdtl", value: __defun__shen_4hdtl})

__defun__shen_4hdhd = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4410 := __args[0]
_ = V4410
reg299803 := PrimHead(V4410)
reg299804 := PrimHead(reg299803)
__ctx.Return(reg299804)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.hdhd", value: __defun__shen_4hdhd})

__defun__shen_4tlhd = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4412 := __args[0]
_ = V4412
reg299805 := PrimHead(V4412)
reg299806 := PrimTail(reg299805)
__ctx.Return(reg299806)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.tlhd", value: __defun__shen_4tlhd})

__defun__shen_4snd_1or_1fail = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4420 := __args[0]
_ = V4420
reg299807 := PrimIsPair(V4420)
var reg299823 Obj
if reg299807 == True {
reg299808 := PrimTail(V4420)
reg299809 := PrimIsPair(reg299808)
var reg299818 Obj
if reg299809 == True {
reg299810 := Nil;
reg299811 := PrimTail(V4420)
reg299812 := PrimTail(reg299811)
reg299813 := PrimEqual(reg299810, reg299812)
var reg299816 Obj
if reg299813 == True {
reg299814 := True;
reg299816 = reg299814
} else {
reg299815 := False;
reg299816 = reg299815
}
reg299818 = reg299816
} else {
reg299817 := False;
reg299818 = reg299817
}
var reg299821 Obj
if reg299818 == True {
reg299819 := True;
reg299821 = reg299819
} else {
reg299820 := False;
reg299821 = reg299820
}
reg299823 = reg299821
} else {
reg299822 := False;
reg299823 = reg299822
}
if reg299823 == True {
reg299824 := PrimTail(V4420)
reg299825 := PrimHead(reg299824)
__ctx.Return(reg299825)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.snd-or-fail", value: __defun__shen_4snd_1or_1fail})

__defun__fail = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg299827 := MakeSymbol("...")
__ctx.Return(reg299827)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "fail", value: __defun__fail})

__defun___5_b_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4428 := __args[0]
_ = V4428
reg299828 := PrimIsPair(V4428)
var reg299844 Obj
if reg299828 == True {
reg299829 := PrimTail(V4428)
reg299830 := PrimIsPair(reg299829)
var reg299839 Obj
if reg299830 == True {
reg299831 := Nil;
reg299832 := PrimTail(V4428)
reg299833 := PrimTail(reg299832)
reg299834 := PrimEqual(reg299831, reg299833)
var reg299837 Obj
if reg299834 == True {
reg299835 := True;
reg299837 = reg299835
} else {
reg299836 := False;
reg299837 = reg299836
}
reg299839 = reg299837
} else {
reg299838 := False;
reg299839 = reg299838
}
var reg299842 Obj
if reg299839 == True {
reg299840 := True;
reg299842 = reg299840
} else {
reg299841 := False;
reg299842 = reg299841
}
reg299844 = reg299842
} else {
reg299843 := False;
reg299844 = reg299843
}
if reg299844 == True {
reg299845 := Nil;
reg299846 := PrimHead(V4428)
reg299847 := Nil;
reg299848 := PrimCons(reg299846, reg299847)
reg299849 := PrimCons(reg299845, reg299848)
__ctx.Return(reg299849)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "<!>", value: __defun___5_b_6})

__defun___5e_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4434 := __args[0]
_ = V4434
reg299851 := PrimIsPair(V4434)
var reg299867 Obj
if reg299851 == True {
reg299852 := PrimTail(V4434)
reg299853 := PrimIsPair(reg299852)
var reg299862 Obj
if reg299853 == True {
reg299854 := Nil;
reg299855 := PrimTail(V4434)
reg299856 := PrimTail(reg299855)
reg299857 := PrimEqual(reg299854, reg299856)
var reg299860 Obj
if reg299857 == True {
reg299858 := True;
reg299860 = reg299858
} else {
reg299859 := False;
reg299860 = reg299859
}
reg299862 = reg299860
} else {
reg299861 := False;
reg299862 = reg299861
}
var reg299865 Obj
if reg299862 == True {
reg299863 := True;
reg299865 = reg299863
} else {
reg299864 := False;
reg299865 = reg299864
}
reg299867 = reg299865
} else {
reg299866 := False;
reg299867 = reg299866
}
if reg299867 == True {
reg299868 := PrimHead(V4434)
reg299869 := Nil;
reg299870 := Nil;
reg299871 := PrimCons(reg299869, reg299870)
reg299872 := PrimCons(reg299868, reg299871)
__ctx.Return(reg299872)
return
} else {
reg299873 := MakeSymbol("<e>")
__ctx.TailApply(__defun__shen_4f__error, reg299873)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "<e>", value: __defun___5e_6})

}
