package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__pr Obj // pr
var __defun__shen_4prh Obj // shen.prh
var __defun__shen_4write_1char_1and_1inc Obj // shen.write-char-and-inc
var __defun__print Obj // print
var __defun__shen_4prhush Obj // shen.prhush
var __defun__shen_4mkstr Obj // shen.mkstr
var __defun__shen_4mkstr_1l Obj // shen.mkstr-l
var __defun__shen_4insert_1l Obj // shen.insert-l
var __defun__shen_4factor_1cn Obj // shen.factor-cn
var __defun__shen_4proc_1nl Obj // shen.proc-nl
var __defun__shen_4mkstr_1r Obj // shen.mkstr-r
var __defun__shen_4insert Obj // shen.insert
var __defun__shen_4insert_1h Obj // shen.insert-h
var __defun__shen_4app Obj // shen.app
var __defun__shen_4arg_1_6str Obj // shen.arg->str
var __defun__shen_4list_1_6str Obj // shen.list->str
var __defun__shen_4maxseq Obj // shen.maxseq
var __defun__shen_4iter_1list Obj // shen.iter-list
var __defun__shen_4str_1_6str Obj // shen.str->str
var __defun__shen_4vector_1_6str Obj // shen.vector->str
var __defun__shen_4print_1vector_2 Obj // shen.print-vector?
var __defun__shen_4fbound_2 Obj // shen.fbound?
var __defun__shen_4tuple Obj // shen.tuple
var __defun__shen_4dictionary Obj // shen.dictionary
var __defun__shen_4iter_1vector Obj // shen.iter-vector
var __defun__shen_4atom_1_6str Obj // shen.atom->str
var __defun__shen_4funexstring Obj // shen.funexstring
var __defun__shen_4list_2 Obj // shen.list?

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308027 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg308027)
return
}, 0))
__defun__pr = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4147 := __args[0]
_ = V4147
V4148 := __args[1]
_ = V4148
reg308028 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308029 := MakeNumber(0)
__ctx.TailApply(__defun__shen_4prh, V4147, V4148, reg308029)
return
}, 0)
reg308031 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.Return(V4147)
return
}, 1)
reg308032 := __e.Try(reg308028).Catch(reg308031)
__ctx.Return(reg308032)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "pr", value: __defun__pr})

__defun__shen_4prh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4152 := __args[0]
_ = V4152
V4153 := __args[1]
_ = V4153
V4154 := __args[2]
_ = V4154
reg308033 := __e.Call(__defun__shen_4write_1char_1and_1inc, V4152, V4153, V4154)
__ctx.TailApply(__defun__shen_4prh, V4152, V4153, reg308033)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.prh", value: __defun__shen_4prh})

__defun__shen_4write_1char_1and_1inc = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4158 := __args[0]
_ = V4158
V4159 := __args[1]
_ = V4159
V4160 := __args[2]
_ = V4160
reg308035 := PrimPos(V4158, V4160)
reg308036 := PrimStringToNumber(reg308035)
reg308037 := PrimWriteByte(reg308036, V4159)
_ = reg308037
reg308038 := MakeNumber(1)
reg308039 := PrimNumberAdd(V4160, reg308038)
__ctx.Return(reg308039)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.write-char-and-inc", value: __defun__shen_4write_1char_1and_1inc})

__defun__print = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4162 := __args[0]
_ = V4162
reg308040 := MakeString("~S")
reg308041 := __e.Call(__defun__shen_4insert, V4162, reg308040)
String := reg308041
_ = String
reg308042 := __e.Call(__defun__stoutput)
reg308043 := __e.Call(__defun__shen_4prhush, String, reg308042)
Print := reg308043
_ = Print
__ctx.Return(V4162)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "print", value: __defun__print})

__defun__shen_4prhush = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4165 := __args[0]
_ = V4165
V4166 := __args[1]
_ = V4166
reg308044 := MakeSymbol("*hush*")
reg308045 := PrimValue(reg308044)
if reg308045 == True {
__ctx.Return(V4165)
return
} else {
__ctx.TailApply(__defun__pr, V4165, V4166)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.prhush", value: __defun__shen_4prhush})

__defun__shen_4mkstr = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4169 := __args[0]
_ = V4169
V4170 := __args[1]
_ = V4170
reg308047 := PrimIsString(V4169)
if reg308047 == True {
reg308048 := __e.Call(__defun__shen_4proc_1nl, V4169)
__ctx.TailApply(__defun__shen_4mkstr_1l, reg308048, V4170)
return
} else {
reg308050 := MakeSymbol("shen.proc-nl")
reg308051 := Nil;
reg308052 := PrimCons(V4169, reg308051)
reg308053 := PrimCons(reg308050, reg308052)
__ctx.TailApply(__defun__shen_4mkstr_1r, reg308053, V4170)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.mkstr", value: __defun__shen_4mkstr})

__defun__shen_4mkstr_1l = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4173 := __args[0]
_ = V4173
V4174 := __args[1]
_ = V4174
reg308055 := Nil;
reg308056 := PrimEqual(reg308055, V4174)
if reg308056 == True {
__ctx.Return(V4173)
return
} else {
reg308057 := PrimIsPair(V4174)
if reg308057 == True {
reg308058 := PrimHead(V4174)
reg308059 := __e.Call(__defun__shen_4insert_1l, reg308058, V4173)
reg308060 := PrimTail(V4174)
__ctx.TailApply(__defun__shen_4mkstr_1l, reg308059, reg308060)
return
} else {
reg308062 := MakeSymbol("shen.mkstr-l")
__ctx.TailApply(__defun__shen_4f__error, reg308062)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.mkstr-l", value: __defun__shen_4mkstr_1l})

__defun__shen_4insert_1l = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4179 := __args[0]
_ = V4179
V4180 := __args[1]
_ = V4180
reg308064 := MakeString("")
reg308065 := PrimEqual(reg308064, V4180)
if reg308065 == True {
reg308066 := MakeString("")
__ctx.Return(reg308066)
return
} else {
reg308067 := __e.Call(__defun__shen_4_7string_2, V4180)
var reg308093 Obj
if reg308067 == True {
reg308068 := MakeString("~")
reg308069 := MakeNumber(0)
reg308070 := PrimPos(V4180, reg308069)
reg308071 := PrimEqual(reg308068, reg308070)
var reg308088 Obj
if reg308071 == True {
reg308072 := PrimTailString(V4180)
reg308073 := __e.Call(__defun__shen_4_7string_2, reg308072)
var reg308083 Obj
if reg308073 == True {
reg308074 := MakeString("A")
reg308075 := PrimTailString(V4180)
reg308076 := MakeNumber(0)
reg308077 := PrimPos(reg308075, reg308076)
reg308078 := PrimEqual(reg308074, reg308077)
var reg308081 Obj
if reg308078 == True {
reg308079 := True;
reg308081 = reg308079
} else {
reg308080 := False;
reg308081 = reg308080
}
reg308083 = reg308081
} else {
reg308082 := False;
reg308083 = reg308082
}
var reg308086 Obj
if reg308083 == True {
reg308084 := True;
reg308086 = reg308084
} else {
reg308085 := False;
reg308086 = reg308085
}
reg308088 = reg308086
} else {
reg308087 := False;
reg308088 = reg308087
}
var reg308091 Obj
if reg308088 == True {
reg308089 := True;
reg308091 = reg308089
} else {
reg308090 := False;
reg308091 = reg308090
}
reg308093 = reg308091
} else {
reg308092 := False;
reg308093 = reg308092
}
if reg308093 == True {
reg308094 := MakeSymbol("shen.app")
reg308095 := PrimTailString(V4180)
reg308096 := PrimTailString(reg308095)
reg308097 := MakeSymbol("shen.a")
reg308098 := Nil;
reg308099 := PrimCons(reg308097, reg308098)
reg308100 := PrimCons(reg308096, reg308099)
reg308101 := PrimCons(V4179, reg308100)
reg308102 := PrimCons(reg308094, reg308101)
__ctx.Return(reg308102)
return
} else {
reg308103 := __e.Call(__defun__shen_4_7string_2, V4180)
var reg308129 Obj
if reg308103 == True {
reg308104 := MakeString("~")
reg308105 := MakeNumber(0)
reg308106 := PrimPos(V4180, reg308105)
reg308107 := PrimEqual(reg308104, reg308106)
var reg308124 Obj
if reg308107 == True {
reg308108 := PrimTailString(V4180)
reg308109 := __e.Call(__defun__shen_4_7string_2, reg308108)
var reg308119 Obj
if reg308109 == True {
reg308110 := MakeString("R")
reg308111 := PrimTailString(V4180)
reg308112 := MakeNumber(0)
reg308113 := PrimPos(reg308111, reg308112)
reg308114 := PrimEqual(reg308110, reg308113)
var reg308117 Obj
if reg308114 == True {
reg308115 := True;
reg308117 = reg308115
} else {
reg308116 := False;
reg308117 = reg308116
}
reg308119 = reg308117
} else {
reg308118 := False;
reg308119 = reg308118
}
var reg308122 Obj
if reg308119 == True {
reg308120 := True;
reg308122 = reg308120
} else {
reg308121 := False;
reg308122 = reg308121
}
reg308124 = reg308122
} else {
reg308123 := False;
reg308124 = reg308123
}
var reg308127 Obj
if reg308124 == True {
reg308125 := True;
reg308127 = reg308125
} else {
reg308126 := False;
reg308127 = reg308126
}
reg308129 = reg308127
} else {
reg308128 := False;
reg308129 = reg308128
}
if reg308129 == True {
reg308130 := MakeSymbol("shen.app")
reg308131 := PrimTailString(V4180)
reg308132 := PrimTailString(reg308131)
reg308133 := MakeSymbol("shen.r")
reg308134 := Nil;
reg308135 := PrimCons(reg308133, reg308134)
reg308136 := PrimCons(reg308132, reg308135)
reg308137 := PrimCons(V4179, reg308136)
reg308138 := PrimCons(reg308130, reg308137)
__ctx.Return(reg308138)
return
} else {
reg308139 := __e.Call(__defun__shen_4_7string_2, V4180)
var reg308165 Obj
if reg308139 == True {
reg308140 := MakeString("~")
reg308141 := MakeNumber(0)
reg308142 := PrimPos(V4180, reg308141)
reg308143 := PrimEqual(reg308140, reg308142)
var reg308160 Obj
if reg308143 == True {
reg308144 := PrimTailString(V4180)
reg308145 := __e.Call(__defun__shen_4_7string_2, reg308144)
var reg308155 Obj
if reg308145 == True {
reg308146 := MakeString("S")
reg308147 := PrimTailString(V4180)
reg308148 := MakeNumber(0)
reg308149 := PrimPos(reg308147, reg308148)
reg308150 := PrimEqual(reg308146, reg308149)
var reg308153 Obj
if reg308150 == True {
reg308151 := True;
reg308153 = reg308151
} else {
reg308152 := False;
reg308153 = reg308152
}
reg308155 = reg308153
} else {
reg308154 := False;
reg308155 = reg308154
}
var reg308158 Obj
if reg308155 == True {
reg308156 := True;
reg308158 = reg308156
} else {
reg308157 := False;
reg308158 = reg308157
}
reg308160 = reg308158
} else {
reg308159 := False;
reg308160 = reg308159
}
var reg308163 Obj
if reg308160 == True {
reg308161 := True;
reg308163 = reg308161
} else {
reg308162 := False;
reg308163 = reg308162
}
reg308165 = reg308163
} else {
reg308164 := False;
reg308165 = reg308164
}
if reg308165 == True {
reg308166 := MakeSymbol("shen.app")
reg308167 := PrimTailString(V4180)
reg308168 := PrimTailString(reg308167)
reg308169 := MakeSymbol("shen.s")
reg308170 := Nil;
reg308171 := PrimCons(reg308169, reg308170)
reg308172 := PrimCons(reg308168, reg308171)
reg308173 := PrimCons(V4179, reg308172)
reg308174 := PrimCons(reg308166, reg308173)
__ctx.Return(reg308174)
return
} else {
reg308175 := __e.Call(__defun__shen_4_7string_2, V4180)
if reg308175 == True {
reg308176 := MakeSymbol("cn")
reg308177 := MakeNumber(0)
reg308178 := PrimPos(V4180, reg308177)
reg308179 := PrimTailString(V4180)
reg308180 := __e.Call(__defun__shen_4insert_1l, V4179, reg308179)
reg308181 := Nil;
reg308182 := PrimCons(reg308180, reg308181)
reg308183 := PrimCons(reg308178, reg308182)
reg308184 := PrimCons(reg308176, reg308183)
__ctx.TailApply(__defun__shen_4factor_1cn, reg308184)
return
} else {
reg308186 := PrimIsPair(V4180)
var reg308219 Obj
if reg308186 == True {
reg308187 := MakeSymbol("cn")
reg308188 := PrimHead(V4180)
reg308189 := PrimEqual(reg308187, reg308188)
var reg308214 Obj
if reg308189 == True {
reg308190 := PrimTail(V4180)
reg308191 := PrimIsPair(reg308190)
var reg308209 Obj
if reg308191 == True {
reg308192 := PrimTail(V4180)
reg308193 := PrimTail(reg308192)
reg308194 := PrimIsPair(reg308193)
var reg308204 Obj
if reg308194 == True {
reg308195 := Nil;
reg308196 := PrimTail(V4180)
reg308197 := PrimTail(reg308196)
reg308198 := PrimTail(reg308197)
reg308199 := PrimEqual(reg308195, reg308198)
var reg308202 Obj
if reg308199 == True {
reg308200 := True;
reg308202 = reg308200
} else {
reg308201 := False;
reg308202 = reg308201
}
reg308204 = reg308202
} else {
reg308203 := False;
reg308204 = reg308203
}
var reg308207 Obj
if reg308204 == True {
reg308205 := True;
reg308207 = reg308205
} else {
reg308206 := False;
reg308207 = reg308206
}
reg308209 = reg308207
} else {
reg308208 := False;
reg308209 = reg308208
}
var reg308212 Obj
if reg308209 == True {
reg308210 := True;
reg308212 = reg308210
} else {
reg308211 := False;
reg308212 = reg308211
}
reg308214 = reg308212
} else {
reg308213 := False;
reg308214 = reg308213
}
var reg308217 Obj
if reg308214 == True {
reg308215 := True;
reg308217 = reg308215
} else {
reg308216 := False;
reg308217 = reg308216
}
reg308219 = reg308217
} else {
reg308218 := False;
reg308219 = reg308218
}
if reg308219 == True {
reg308220 := MakeSymbol("cn")
reg308221 := PrimTail(V4180)
reg308222 := PrimHead(reg308221)
reg308223 := PrimTail(V4180)
reg308224 := PrimTail(reg308223)
reg308225 := PrimHead(reg308224)
reg308226 := __e.Call(__defun__shen_4insert_1l, V4179, reg308225)
reg308227 := Nil;
reg308228 := PrimCons(reg308226, reg308227)
reg308229 := PrimCons(reg308222, reg308228)
reg308230 := PrimCons(reg308220, reg308229)
__ctx.Return(reg308230)
return
} else {
reg308231 := PrimIsPair(V4180)
var reg308274 Obj
if reg308231 == True {
reg308232 := MakeSymbol("shen.app")
reg308233 := PrimHead(V4180)
reg308234 := PrimEqual(reg308232, reg308233)
var reg308269 Obj
if reg308234 == True {
reg308235 := PrimTail(V4180)
reg308236 := PrimIsPair(reg308235)
var reg308264 Obj
if reg308236 == True {
reg308237 := PrimTail(V4180)
reg308238 := PrimTail(reg308237)
reg308239 := PrimIsPair(reg308238)
var reg308259 Obj
if reg308239 == True {
reg308240 := PrimTail(V4180)
reg308241 := PrimTail(reg308240)
reg308242 := PrimTail(reg308241)
reg308243 := PrimIsPair(reg308242)
var reg308254 Obj
if reg308243 == True {
reg308244 := Nil;
reg308245 := PrimTail(V4180)
reg308246 := PrimTail(reg308245)
reg308247 := PrimTail(reg308246)
reg308248 := PrimTail(reg308247)
reg308249 := PrimEqual(reg308244, reg308248)
var reg308252 Obj
if reg308249 == True {
reg308250 := True;
reg308252 = reg308250
} else {
reg308251 := False;
reg308252 = reg308251
}
reg308254 = reg308252
} else {
reg308253 := False;
reg308254 = reg308253
}
var reg308257 Obj
if reg308254 == True {
reg308255 := True;
reg308257 = reg308255
} else {
reg308256 := False;
reg308257 = reg308256
}
reg308259 = reg308257
} else {
reg308258 := False;
reg308259 = reg308258
}
var reg308262 Obj
if reg308259 == True {
reg308260 := True;
reg308262 = reg308260
} else {
reg308261 := False;
reg308262 = reg308261
}
reg308264 = reg308262
} else {
reg308263 := False;
reg308264 = reg308263
}
var reg308267 Obj
if reg308264 == True {
reg308265 := True;
reg308267 = reg308265
} else {
reg308266 := False;
reg308267 = reg308266
}
reg308269 = reg308267
} else {
reg308268 := False;
reg308269 = reg308268
}
var reg308272 Obj
if reg308269 == True {
reg308270 := True;
reg308272 = reg308270
} else {
reg308271 := False;
reg308272 = reg308271
}
reg308274 = reg308272
} else {
reg308273 := False;
reg308274 = reg308273
}
if reg308274 == True {
reg308275 := MakeSymbol("shen.app")
reg308276 := PrimTail(V4180)
reg308277 := PrimHead(reg308276)
reg308278 := PrimTail(V4180)
reg308279 := PrimTail(reg308278)
reg308280 := PrimHead(reg308279)
reg308281 := __e.Call(__defun__shen_4insert_1l, V4179, reg308280)
reg308282 := PrimTail(V4180)
reg308283 := PrimTail(reg308282)
reg308284 := PrimTail(reg308283)
reg308285 := PrimCons(reg308281, reg308284)
reg308286 := PrimCons(reg308277, reg308285)
reg308287 := PrimCons(reg308275, reg308286)
__ctx.Return(reg308287)
return
} else {
reg308288 := MakeSymbol("shen.insert-l")
__ctx.TailApply(__defun__shen_4f__error, reg308288)
return
}
}
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.insert-l", value: __defun__shen_4insert_1l})

__defun__shen_4factor_1cn = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4182 := __args[0]
_ = V4182
reg308290 := PrimIsPair(V4182)
var reg308396 Obj
if reg308290 == True {
reg308291 := MakeSymbol("cn")
reg308292 := PrimHead(V4182)
reg308293 := PrimEqual(reg308291, reg308292)
var reg308391 Obj
if reg308293 == True {
reg308294 := PrimTail(V4182)
reg308295 := PrimIsPair(reg308294)
var reg308386 Obj
if reg308295 == True {
reg308296 := PrimTail(V4182)
reg308297 := PrimTail(reg308296)
reg308298 := PrimIsPair(reg308297)
var reg308381 Obj
if reg308298 == True {
reg308299 := PrimTail(V4182)
reg308300 := PrimTail(reg308299)
reg308301 := PrimHead(reg308300)
reg308302 := PrimIsPair(reg308301)
var reg308376 Obj
if reg308302 == True {
reg308303 := MakeSymbol("cn")
reg308304 := PrimTail(V4182)
reg308305 := PrimTail(reg308304)
reg308306 := PrimHead(reg308305)
reg308307 := PrimHead(reg308306)
reg308308 := PrimEqual(reg308303, reg308307)
var reg308371 Obj
if reg308308 == True {
reg308309 := PrimTail(V4182)
reg308310 := PrimTail(reg308309)
reg308311 := PrimHead(reg308310)
reg308312 := PrimTail(reg308311)
reg308313 := PrimIsPair(reg308312)
var reg308366 Obj
if reg308313 == True {
reg308314 := PrimTail(V4182)
reg308315 := PrimTail(reg308314)
reg308316 := PrimHead(reg308315)
reg308317 := PrimTail(reg308316)
reg308318 := PrimTail(reg308317)
reg308319 := PrimIsPair(reg308318)
var reg308361 Obj
if reg308319 == True {
reg308320 := Nil;
reg308321 := PrimTail(V4182)
reg308322 := PrimTail(reg308321)
reg308323 := PrimHead(reg308322)
reg308324 := PrimTail(reg308323)
reg308325 := PrimTail(reg308324)
reg308326 := PrimTail(reg308325)
reg308327 := PrimEqual(reg308320, reg308326)
var reg308356 Obj
if reg308327 == True {
reg308328 := Nil;
reg308329 := PrimTail(V4182)
reg308330 := PrimTail(reg308329)
reg308331 := PrimTail(reg308330)
reg308332 := PrimEqual(reg308328, reg308331)
var reg308351 Obj
if reg308332 == True {
reg308333 := PrimTail(V4182)
reg308334 := PrimHead(reg308333)
reg308335 := PrimIsString(reg308334)
var reg308346 Obj
if reg308335 == True {
reg308336 := PrimTail(V4182)
reg308337 := PrimTail(reg308336)
reg308338 := PrimHead(reg308337)
reg308339 := PrimTail(reg308338)
reg308340 := PrimHead(reg308339)
reg308341 := PrimIsString(reg308340)
var reg308344 Obj
if reg308341 == True {
reg308342 := True;
reg308344 = reg308342
} else {
reg308343 := False;
reg308344 = reg308343
}
reg308346 = reg308344
} else {
reg308345 := False;
reg308346 = reg308345
}
var reg308349 Obj
if reg308346 == True {
reg308347 := True;
reg308349 = reg308347
} else {
reg308348 := False;
reg308349 = reg308348
}
reg308351 = reg308349
} else {
reg308350 := False;
reg308351 = reg308350
}
var reg308354 Obj
if reg308351 == True {
reg308352 := True;
reg308354 = reg308352
} else {
reg308353 := False;
reg308354 = reg308353
}
reg308356 = reg308354
} else {
reg308355 := False;
reg308356 = reg308355
}
var reg308359 Obj
if reg308356 == True {
reg308357 := True;
reg308359 = reg308357
} else {
reg308358 := False;
reg308359 = reg308358
}
reg308361 = reg308359
} else {
reg308360 := False;
reg308361 = reg308360
}
var reg308364 Obj
if reg308361 == True {
reg308362 := True;
reg308364 = reg308362
} else {
reg308363 := False;
reg308364 = reg308363
}
reg308366 = reg308364
} else {
reg308365 := False;
reg308366 = reg308365
}
var reg308369 Obj
if reg308366 == True {
reg308367 := True;
reg308369 = reg308367
} else {
reg308368 := False;
reg308369 = reg308368
}
reg308371 = reg308369
} else {
reg308370 := False;
reg308371 = reg308370
}
var reg308374 Obj
if reg308371 == True {
reg308372 := True;
reg308374 = reg308372
} else {
reg308373 := False;
reg308374 = reg308373
}
reg308376 = reg308374
} else {
reg308375 := False;
reg308376 = reg308375
}
var reg308379 Obj
if reg308376 == True {
reg308377 := True;
reg308379 = reg308377
} else {
reg308378 := False;
reg308379 = reg308378
}
reg308381 = reg308379
} else {
reg308380 := False;
reg308381 = reg308380
}
var reg308384 Obj
if reg308381 == True {
reg308382 := True;
reg308384 = reg308382
} else {
reg308383 := False;
reg308384 = reg308383
}
reg308386 = reg308384
} else {
reg308385 := False;
reg308386 = reg308385
}
var reg308389 Obj
if reg308386 == True {
reg308387 := True;
reg308389 = reg308387
} else {
reg308388 := False;
reg308389 = reg308388
}
reg308391 = reg308389
} else {
reg308390 := False;
reg308391 = reg308390
}
var reg308394 Obj
if reg308391 == True {
reg308392 := True;
reg308394 = reg308392
} else {
reg308393 := False;
reg308394 = reg308393
}
reg308396 = reg308394
} else {
reg308395 := False;
reg308396 = reg308395
}
if reg308396 == True {
reg308397 := MakeSymbol("cn")
reg308398 := PrimTail(V4182)
reg308399 := PrimHead(reg308398)
reg308400 := PrimTail(V4182)
reg308401 := PrimTail(reg308400)
reg308402 := PrimHead(reg308401)
reg308403 := PrimTail(reg308402)
reg308404 := PrimHead(reg308403)
reg308405 := PrimStringConcat(reg308399, reg308404)
reg308406 := PrimTail(V4182)
reg308407 := PrimTail(reg308406)
reg308408 := PrimHead(reg308407)
reg308409 := PrimTail(reg308408)
reg308410 := PrimTail(reg308409)
reg308411 := PrimCons(reg308405, reg308410)
reg308412 := PrimCons(reg308397, reg308411)
__ctx.Return(reg308412)
return
} else {
__ctx.Return(V4182)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.factor-cn", value: __defun__shen_4factor_1cn})

__defun__shen_4proc_1nl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4184 := __args[0]
_ = V4184
reg308413 := MakeString("")
reg308414 := PrimEqual(reg308413, V4184)
if reg308414 == True {
reg308415 := MakeString("")
__ctx.Return(reg308415)
return
} else {
reg308416 := __e.Call(__defun__shen_4_7string_2, V4184)
var reg308442 Obj
if reg308416 == True {
reg308417 := MakeString("~")
reg308418 := MakeNumber(0)
reg308419 := PrimPos(V4184, reg308418)
reg308420 := PrimEqual(reg308417, reg308419)
var reg308437 Obj
if reg308420 == True {
reg308421 := PrimTailString(V4184)
reg308422 := __e.Call(__defun__shen_4_7string_2, reg308421)
var reg308432 Obj
if reg308422 == True {
reg308423 := MakeString("%")
reg308424 := PrimTailString(V4184)
reg308425 := MakeNumber(0)
reg308426 := PrimPos(reg308424, reg308425)
reg308427 := PrimEqual(reg308423, reg308426)
var reg308430 Obj
if reg308427 == True {
reg308428 := True;
reg308430 = reg308428
} else {
reg308429 := False;
reg308430 = reg308429
}
reg308432 = reg308430
} else {
reg308431 := False;
reg308432 = reg308431
}
var reg308435 Obj
if reg308432 == True {
reg308433 := True;
reg308435 = reg308433
} else {
reg308434 := False;
reg308435 = reg308434
}
reg308437 = reg308435
} else {
reg308436 := False;
reg308437 = reg308436
}
var reg308440 Obj
if reg308437 == True {
reg308438 := True;
reg308440 = reg308438
} else {
reg308439 := False;
reg308440 = reg308439
}
reg308442 = reg308440
} else {
reg308441 := False;
reg308442 = reg308441
}
if reg308442 == True {
reg308443 := MakeNumber(10)
reg308444 := PrimNumberToString(reg308443)
reg308445 := PrimTailString(V4184)
reg308446 := PrimTailString(reg308445)
reg308447 := __e.Call(__defun__shen_4proc_1nl, reg308446)
reg308448 := PrimStringConcat(reg308444, reg308447)
__ctx.Return(reg308448)
return
} else {
reg308449 := __e.Call(__defun__shen_4_7string_2, V4184)
if reg308449 == True {
reg308450 := MakeNumber(0)
reg308451 := PrimPos(V4184, reg308450)
reg308452 := PrimTailString(V4184)
reg308453 := __e.Call(__defun__shen_4proc_1nl, reg308452)
reg308454 := PrimStringConcat(reg308451, reg308453)
__ctx.Return(reg308454)
return
} else {
reg308455 := MakeSymbol("shen.proc-nl")
__ctx.TailApply(__defun__shen_4f__error, reg308455)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.proc-nl", value: __defun__shen_4proc_1nl})

__defun__shen_4mkstr_1r = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4187 := __args[0]
_ = V4187
V4188 := __args[1]
_ = V4188
reg308457 := Nil;
reg308458 := PrimEqual(reg308457, V4188)
if reg308458 == True {
__ctx.Return(V4187)
return
} else {
reg308459 := PrimIsPair(V4188)
if reg308459 == True {
reg308460 := MakeSymbol("shen.insert")
reg308461 := PrimHead(V4188)
reg308462 := Nil;
reg308463 := PrimCons(V4187, reg308462)
reg308464 := PrimCons(reg308461, reg308463)
reg308465 := PrimCons(reg308460, reg308464)
reg308466 := PrimTail(V4188)
__ctx.TailApply(__defun__shen_4mkstr_1r, reg308465, reg308466)
return
} else {
reg308468 := MakeSymbol("shen.mkstr-r")
__ctx.TailApply(__defun__shen_4f__error, reg308468)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.mkstr-r", value: __defun__shen_4mkstr_1r})

__defun__shen_4insert = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4191 := __args[0]
_ = V4191
V4192 := __args[1]
_ = V4192
reg308470 := MakeString("")
__ctx.TailApply(__defun__shen_4insert_1h, V4191, V4192, reg308470)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.insert", value: __defun__shen_4insert})

__defun__shen_4insert_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4198 := __args[0]
_ = V4198
V4199 := __args[1]
_ = V4199
V4200 := __args[2]
_ = V4200
reg308472 := MakeString("")
reg308473 := PrimEqual(reg308472, V4199)
if reg308473 == True {
__ctx.Return(V4200)
return
} else {
reg308474 := __e.Call(__defun__shen_4_7string_2, V4199)
var reg308500 Obj
if reg308474 == True {
reg308475 := MakeString("~")
reg308476 := MakeNumber(0)
reg308477 := PrimPos(V4199, reg308476)
reg308478 := PrimEqual(reg308475, reg308477)
var reg308495 Obj
if reg308478 == True {
reg308479 := PrimTailString(V4199)
reg308480 := __e.Call(__defun__shen_4_7string_2, reg308479)
var reg308490 Obj
if reg308480 == True {
reg308481 := MakeString("A")
reg308482 := PrimTailString(V4199)
reg308483 := MakeNumber(0)
reg308484 := PrimPos(reg308482, reg308483)
reg308485 := PrimEqual(reg308481, reg308484)
var reg308488 Obj
if reg308485 == True {
reg308486 := True;
reg308488 = reg308486
} else {
reg308487 := False;
reg308488 = reg308487
}
reg308490 = reg308488
} else {
reg308489 := False;
reg308490 = reg308489
}
var reg308493 Obj
if reg308490 == True {
reg308491 := True;
reg308493 = reg308491
} else {
reg308492 := False;
reg308493 = reg308492
}
reg308495 = reg308493
} else {
reg308494 := False;
reg308495 = reg308494
}
var reg308498 Obj
if reg308495 == True {
reg308496 := True;
reg308498 = reg308496
} else {
reg308497 := False;
reg308498 = reg308497
}
reg308500 = reg308498
} else {
reg308499 := False;
reg308500 = reg308499
}
if reg308500 == True {
reg308501 := PrimTailString(V4199)
reg308502 := PrimTailString(reg308501)
reg308503 := MakeSymbol("shen.a")
reg308504 := __e.Call(__defun__shen_4app, V4198, reg308502, reg308503)
reg308505 := PrimStringConcat(V4200, reg308504)
__ctx.Return(reg308505)
return
} else {
reg308506 := __e.Call(__defun__shen_4_7string_2, V4199)
var reg308532 Obj
if reg308506 == True {
reg308507 := MakeString("~")
reg308508 := MakeNumber(0)
reg308509 := PrimPos(V4199, reg308508)
reg308510 := PrimEqual(reg308507, reg308509)
var reg308527 Obj
if reg308510 == True {
reg308511 := PrimTailString(V4199)
reg308512 := __e.Call(__defun__shen_4_7string_2, reg308511)
var reg308522 Obj
if reg308512 == True {
reg308513 := MakeString("R")
reg308514 := PrimTailString(V4199)
reg308515 := MakeNumber(0)
reg308516 := PrimPos(reg308514, reg308515)
reg308517 := PrimEqual(reg308513, reg308516)
var reg308520 Obj
if reg308517 == True {
reg308518 := True;
reg308520 = reg308518
} else {
reg308519 := False;
reg308520 = reg308519
}
reg308522 = reg308520
} else {
reg308521 := False;
reg308522 = reg308521
}
var reg308525 Obj
if reg308522 == True {
reg308523 := True;
reg308525 = reg308523
} else {
reg308524 := False;
reg308525 = reg308524
}
reg308527 = reg308525
} else {
reg308526 := False;
reg308527 = reg308526
}
var reg308530 Obj
if reg308527 == True {
reg308528 := True;
reg308530 = reg308528
} else {
reg308529 := False;
reg308530 = reg308529
}
reg308532 = reg308530
} else {
reg308531 := False;
reg308532 = reg308531
}
if reg308532 == True {
reg308533 := PrimTailString(V4199)
reg308534 := PrimTailString(reg308533)
reg308535 := MakeSymbol("shen.r")
reg308536 := __e.Call(__defun__shen_4app, V4198, reg308534, reg308535)
reg308537 := PrimStringConcat(V4200, reg308536)
__ctx.Return(reg308537)
return
} else {
reg308538 := __e.Call(__defun__shen_4_7string_2, V4199)
var reg308564 Obj
if reg308538 == True {
reg308539 := MakeString("~")
reg308540 := MakeNumber(0)
reg308541 := PrimPos(V4199, reg308540)
reg308542 := PrimEqual(reg308539, reg308541)
var reg308559 Obj
if reg308542 == True {
reg308543 := PrimTailString(V4199)
reg308544 := __e.Call(__defun__shen_4_7string_2, reg308543)
var reg308554 Obj
if reg308544 == True {
reg308545 := MakeString("S")
reg308546 := PrimTailString(V4199)
reg308547 := MakeNumber(0)
reg308548 := PrimPos(reg308546, reg308547)
reg308549 := PrimEqual(reg308545, reg308548)
var reg308552 Obj
if reg308549 == True {
reg308550 := True;
reg308552 = reg308550
} else {
reg308551 := False;
reg308552 = reg308551
}
reg308554 = reg308552
} else {
reg308553 := False;
reg308554 = reg308553
}
var reg308557 Obj
if reg308554 == True {
reg308555 := True;
reg308557 = reg308555
} else {
reg308556 := False;
reg308557 = reg308556
}
reg308559 = reg308557
} else {
reg308558 := False;
reg308559 = reg308558
}
var reg308562 Obj
if reg308559 == True {
reg308560 := True;
reg308562 = reg308560
} else {
reg308561 := False;
reg308562 = reg308561
}
reg308564 = reg308562
} else {
reg308563 := False;
reg308564 = reg308563
}
if reg308564 == True {
reg308565 := PrimTailString(V4199)
reg308566 := PrimTailString(reg308565)
reg308567 := MakeSymbol("shen.s")
reg308568 := __e.Call(__defun__shen_4app, V4198, reg308566, reg308567)
reg308569 := PrimStringConcat(V4200, reg308568)
__ctx.Return(reg308569)
return
} else {
reg308570 := __e.Call(__defun__shen_4_7string_2, V4199)
if reg308570 == True {
reg308571 := PrimTailString(V4199)
reg308572 := MakeNumber(0)
reg308573 := PrimPos(V4199, reg308572)
reg308574 := PrimStringConcat(V4200, reg308573)
__ctx.TailApply(__defun__shen_4insert_1h, V4198, reg308571, reg308574)
return
} else {
reg308576 := MakeSymbol("shen.insert-h")
__ctx.TailApply(__defun__shen_4f__error, reg308576)
return
}
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.insert-h", value: __defun__shen_4insert_1h})

__defun__shen_4app = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4204 := __args[0]
_ = V4204
V4205 := __args[1]
_ = V4205
V4206 := __args[2]
_ = V4206
reg308578 := __e.Call(__defun__shen_4arg_1_6str, V4204, V4206)
reg308579 := PrimStringConcat(reg308578, V4205)
__ctx.Return(reg308579)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.app", value: __defun__shen_4app})

__defun__shen_4arg_1_6str = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4214 := __args[0]
_ = V4214
V4215 := __args[1]
_ = V4215
reg308580 := __e.Call(__defun__fail)
reg308581 := PrimEqual(V4214, reg308580)
if reg308581 == True {
reg308582 := MakeString("...")
__ctx.Return(reg308582)
return
} else {
reg308583 := __e.Call(__defun__shen_4list_2, V4214)
if reg308583 == True {
__ctx.TailApply(__defun__shen_4list_1_6str, V4214, V4215)
return
} else {
reg308585 := PrimIsString(V4214)
if reg308585 == True {
__ctx.TailApply(__defun__shen_4str_1_6str, V4214, V4215)
return
} else {
reg308587 := PrimIsVector(V4214)
if reg308587 == True {
__ctx.TailApply(__defun__shen_4vector_1_6str, V4214, V4215)
return
} else {
__ctx.TailApply(__defun__shen_4atom_1_6str, V4214)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.arg->str", value: __defun__shen_4arg_1_6str})

__defun__shen_4list_1_6str = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4218 := __args[0]
_ = V4218
V4219 := __args[1]
_ = V4219
reg308590 := MakeSymbol("shen.r")
reg308591 := PrimEqual(reg308590, V4219)
if reg308591 == True {
reg308592 := MakeString("(")
reg308593 := MakeSymbol("shen.r")
reg308594 := __e.Call(__defun__shen_4maxseq)
reg308595 := __e.Call(__defun__shen_4iter_1list, V4218, reg308593, reg308594)
reg308596 := MakeString(")")
reg308597 := __e.Call(__defun___8s, reg308595, reg308596)
__ctx.TailApply(__defun___8s, reg308592, reg308597)
return
} else {
reg308599 := MakeString("[")
reg308600 := __e.Call(__defun__shen_4maxseq)
reg308601 := __e.Call(__defun__shen_4iter_1list, V4218, V4219, reg308600)
reg308602 := MakeString("]")
reg308603 := __e.Call(__defun___8s, reg308601, reg308602)
__ctx.TailApply(__defun___8s, reg308599, reg308603)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.list->str", value: __defun__shen_4list_1_6str})

__defun__shen_4maxseq = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308605 := MakeSymbol("*maximum-print-sequence-size*")
reg308606 := PrimValue(reg308605)
__ctx.Return(reg308606)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.maxseq", value: __defun__shen_4maxseq})

__defun__shen_4iter_1list = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4233 := __args[0]
_ = V4233
V4234 := __args[1]
_ = V4234
V4235 := __args[2]
_ = V4235
reg308607 := Nil;
reg308608 := PrimEqual(reg308607, V4233)
if reg308608 == True {
reg308609 := MakeString("")
__ctx.Return(reg308609)
return
} else {
reg308610 := MakeNumber(0)
reg308611 := PrimEqual(reg308610, V4235)
if reg308611 == True {
reg308612 := MakeString("... etc")
__ctx.Return(reg308612)
return
} else {
reg308613 := PrimIsPair(V4233)
var reg308621 Obj
if reg308613 == True {
reg308614 := Nil;
reg308615 := PrimTail(V4233)
reg308616 := PrimEqual(reg308614, reg308615)
var reg308619 Obj
if reg308616 == True {
reg308617 := True;
reg308619 = reg308617
} else {
reg308618 := False;
reg308619 = reg308618
}
reg308621 = reg308619
} else {
reg308620 := False;
reg308621 = reg308620
}
if reg308621 == True {
reg308622 := PrimHead(V4233)
__ctx.TailApply(__defun__shen_4arg_1_6str, reg308622, V4234)
return
} else {
reg308624 := PrimIsPair(V4233)
if reg308624 == True {
reg308625 := PrimHead(V4233)
reg308626 := __e.Call(__defun__shen_4arg_1_6str, reg308625, V4234)
reg308627 := MakeString(" ")
reg308628 := PrimTail(V4233)
reg308629 := MakeNumber(1)
reg308630 := PrimNumberSubtract(V4235, reg308629)
reg308631 := __e.Call(__defun__shen_4iter_1list, reg308628, V4234, reg308630)
reg308632 := __e.Call(__defun___8s, reg308627, reg308631)
__ctx.TailApply(__defun___8s, reg308626, reg308632)
return
} else {
reg308634 := MakeString("|")
reg308635 := MakeString(" ")
reg308636 := __e.Call(__defun__shen_4arg_1_6str, V4233, V4234)
reg308637 := __e.Call(__defun___8s, reg308635, reg308636)
__ctx.TailApply(__defun___8s, reg308634, reg308637)
return
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.iter-list", value: __defun__shen_4iter_1list})

__defun__shen_4str_1_6str = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4242 := __args[0]
_ = V4242
V4243 := __args[1]
_ = V4243
reg308639 := MakeSymbol("shen.a")
reg308640 := PrimEqual(reg308639, V4243)
if reg308640 == True {
__ctx.Return(V4242)
return
} else {
reg308641 := MakeNumber(34)
reg308642 := PrimNumberToString(reg308641)
reg308643 := MakeNumber(34)
reg308644 := PrimNumberToString(reg308643)
reg308645 := __e.Call(__defun___8s, V4242, reg308644)
__ctx.TailApply(__defun___8s, reg308642, reg308645)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.str->str", value: __defun__shen_4str_1_6str})

__defun__shen_4vector_1_6str = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4246 := __args[0]
_ = V4246
V4247 := __args[1]
_ = V4247
reg308648 := __e.Call(__defun__shen_4print_1vector_2, V4246)
if reg308648 == True {
reg308649 := MakeNumber(0)
reg308650 := PrimVectorGet(V4246, reg308649)
reg308651 := __e.Call(__defun__function, reg308650)
f308647 := reg308651
_ = f308647
__ctx.TailApply(f308647, V4246)
return
} else {
reg308653 := __e.Call(__defun__vector_2, V4246)
if reg308653 == True {
reg308654 := MakeString("<")
reg308655 := MakeNumber(1)
reg308656 := __e.Call(__defun__shen_4maxseq)
reg308657 := __e.Call(__defun__shen_4iter_1vector, V4246, reg308655, V4247, reg308656)
reg308658 := MakeString(">")
reg308659 := __e.Call(__defun___8s, reg308657, reg308658)
__ctx.TailApply(__defun___8s, reg308654, reg308659)
return
} else {
reg308661 := MakeString("<")
reg308662 := MakeString("<")
reg308663 := MakeNumber(0)
reg308664 := __e.Call(__defun__shen_4maxseq)
reg308665 := __e.Call(__defun__shen_4iter_1vector, V4246, reg308663, V4247, reg308664)
reg308666 := MakeString(">>")
reg308667 := __e.Call(__defun___8s, reg308665, reg308666)
reg308668 := __e.Call(__defun___8s, reg308662, reg308667)
__ctx.TailApply(__defun___8s, reg308661, reg308668)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.vector->str", value: __defun__shen_4vector_1_6str})

__defun__shen_4print_1vector_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4249 := __args[0]
_ = V4249
reg308670 := MakeNumber(0)
reg308671 := PrimVectorGet(V4249, reg308670)
Zero := reg308671
_ = Zero
reg308672 := MakeSymbol("shen.tuple")
reg308673 := PrimEqual(Zero, reg308672)
if reg308673 == True {
reg308674 := True;
__ctx.Return(reg308674)
return
} else {
reg308675 := MakeSymbol("shen.pvar")
reg308676 := PrimEqual(Zero, reg308675)
if reg308676 == True {
reg308677 := True;
__ctx.Return(reg308677)
return
} else {
reg308678 := MakeSymbol("shen.dictionary")
reg308679 := PrimEqual(Zero, reg308678)
if reg308679 == True {
reg308680 := True;
__ctx.Return(reg308680)
return
} else {
reg308681 := PrimIsNumber(Zero)
reg308682 := PrimNot(reg308681)
if reg308682 == True {
__ctx.TailApply(__defun__shen_4fbound_2, Zero)
return
} else {
reg308684 := False;
__ctx.Return(reg308684)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.print-vector?", value: __defun__shen_4print_1vector_2})

__defun__shen_4fbound_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4251 := __args[0]
_ = V4251
reg308685 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308686 := __e.Call(__defun__shen_4lookup_1func, V4251)
_ = reg308686
reg308687 := True;
__ctx.Return(reg308687)
return
}, 0)
reg308688 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg308689 := False;
__ctx.Return(reg308689)
return
}, 1)
reg308690 := __e.Try(reg308685).Catch(reg308688)
__ctx.Return(reg308690)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.fbound?", value: __defun__shen_4fbound_2})

__defun__shen_4tuple = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4253 := __args[0]
_ = V4253
reg308691 := MakeString("(@p ")
reg308692 := MakeNumber(1)
reg308693 := PrimVectorGet(V4253, reg308692)
reg308694 := MakeString(" ")
reg308695 := MakeNumber(2)
reg308696 := PrimVectorGet(V4253, reg308695)
reg308697 := MakeString(")")
reg308698 := MakeSymbol("shen.s")
reg308699 := __e.Call(__defun__shen_4app, reg308696, reg308697, reg308698)
reg308700 := PrimStringConcat(reg308694, reg308699)
reg308701 := MakeSymbol("shen.s")
reg308702 := __e.Call(__defun__shen_4app, reg308693, reg308700, reg308701)
reg308703 := PrimStringConcat(reg308691, reg308702)
__ctx.Return(reg308703)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.tuple", value: __defun__shen_4tuple})

__defun__shen_4dictionary = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4255 := __args[0]
_ = V4255
reg308704 := MakeString("(dict ...)")
__ctx.Return(reg308704)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dictionary", value: __defun__shen_4dictionary})

__defun__shen_4iter_1vector = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4266 := __args[0]
_ = V4266
V4267 := __args[1]
_ = V4267
V4268 := __args[2]
_ = V4268
V4269 := __args[3]
_ = V4269
reg308705 := MakeNumber(0)
reg308706 := PrimEqual(reg308705, V4269)
if reg308706 == True {
reg308707 := MakeString("... etc")
__ctx.Return(reg308707)
return
} else {
reg308708 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308709 := PrimVectorGet(V4266, V4267)
__ctx.Return(reg308709)
return
}, 0)
reg308710 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg308711 := MakeSymbol("shen.out-of-bounds")
__ctx.Return(reg308711)
return
}, 1)
reg308712 := __e.Try(reg308708).Catch(reg308710)
Item := reg308712
_ = Item
reg308713 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308714 := MakeNumber(1)
reg308715 := PrimNumberAdd(V4267, reg308714)
reg308716 := PrimVectorGet(V4266, reg308715)
__ctx.Return(reg308716)
return
}, 0)
reg308717 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg308718 := MakeSymbol("shen.out-of-bounds")
__ctx.Return(reg308718)
return
}, 1)
reg308719 := __e.Try(reg308713).Catch(reg308717)
Next := reg308719
_ = Next
reg308720 := MakeSymbol("shen.out-of-bounds")
reg308721 := PrimEqual(Item, reg308720)
if reg308721 == True {
reg308722 := MakeString("")
__ctx.Return(reg308722)
return
} else {
reg308723 := MakeSymbol("shen.out-of-bounds")
reg308724 := PrimEqual(Next, reg308723)
if reg308724 == True {
__ctx.TailApply(__defun__shen_4arg_1_6str, Item, V4268)
return
} else {
reg308726 := __e.Call(__defun__shen_4arg_1_6str, Item, V4268)
reg308727 := MakeString(" ")
reg308728 := MakeNumber(1)
reg308729 := PrimNumberAdd(V4267, reg308728)
reg308730 := MakeNumber(1)
reg308731 := PrimNumberSubtract(V4269, reg308730)
reg308732 := __e.Call(__defun__shen_4iter_1vector, V4266, reg308729, V4268, reg308731)
reg308733 := __e.Call(__defun___8s, reg308727, reg308732)
__ctx.TailApply(__defun___8s, reg308726, reg308733)
return
}
}
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.iter-vector", value: __defun__shen_4iter_1vector})

__defun__shen_4atom_1_6str = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4271 := __args[0]
_ = V4271
reg308735 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308736 := PrimStr(V4271)
__ctx.Return(reg308736)
return
}, 0)
reg308737 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.TailApply(__defun__shen_4funexstring)
return
}, 1)
reg308739 := __e.Try(reg308735).Catch(reg308737)
__ctx.Return(reg308739)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.atom->str", value: __defun__shen_4atom_1_6str})

__defun__shen_4funexstring = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308740 := MakeString("\x10")
reg308741 := MakeString("f")
reg308742 := MakeString("u")
reg308743 := MakeString("n")
reg308744 := MakeString("e")
reg308745 := MakeString("x")
reg308746 := PrimIntern(reg308745)
reg308747 := __e.Call(__defun__gensym, reg308746)
reg308748 := MakeSymbol("shen.a")
reg308749 := __e.Call(__defun__shen_4arg_1_6str, reg308747, reg308748)
reg308750 := MakeString("\x11")
reg308751 := __e.Call(__defun___8s, reg308749, reg308750)
reg308752 := __e.Call(__defun___8s, reg308744, reg308751)
reg308753 := __e.Call(__defun___8s, reg308743, reg308752)
reg308754 := __e.Call(__defun___8s, reg308742, reg308753)
reg308755 := __e.Call(__defun___8s, reg308741, reg308754)
__ctx.TailApply(__defun___8s, reg308740, reg308755)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.funexstring", value: __defun__shen_4funexstring})

__defun__shen_4list_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4273 := __args[0]
_ = V4273
reg308757 := __e.Call(__defun__empty_2, V4273)
if reg308757 == True {
reg308758 := True;
__ctx.Return(reg308758)
return
} else {
reg308759 := PrimIsPair(V4273)
if reg308759 == True {
reg308760 := True;
__ctx.Return(reg308760)
return
} else {
reg308761 := False;
__ctx.Return(reg308761)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.list?", value: __defun__shen_4list_2})

}
