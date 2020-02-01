package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4repl Obj // shen.repl
var __defun__shen_4loop Obj // shen.loop
var __defun__shen_4toplevel_1display_1exception Obj // shen.toplevel-display-exception
var __defun__shen_4credits Obj // shen.credits
var __defun__shen_4initialise__environment Obj // shen.initialise_environment
var __defun__shen_4multiple_1set Obj // shen.multiple-set
var __defun__destroy Obj // destroy
var __defun__shen_4read_1evaluate_1print Obj // shen.read-evaluate-print
var __defun__shen_4retrieve_1from_1history_1if_1needed Obj // shen.retrieve-from-history-if-needed
var __defun__shen_4percent Obj // shen.percent
var __defun__shen_4exclamation Obj // shen.exclamation
var __defun__shen_4prbytes Obj // shen.prbytes
var __defun__shen_4update__history Obj // shen.update_history
var __defun__shen_4toplineread Obj // shen.toplineread
var __defun__shen_4toplineread__loop Obj // shen.toplineread_loop
var __defun__shen_4hat Obj // shen.hat
var __defun__shen_4newline Obj // shen.newline
var __defun__shen_4carriage_1return Obj // shen.carriage-return
var __defun__tc Obj // tc
var __defun__shen_4prompt Obj // shen.prompt
var __defun__shen_4toplevel Obj // shen.toplevel
var __defun__shen_4find_1past_1inputs Obj // shen.find-past-inputs
var __defun__shen_4make_1key Obj // shen.make-key
var __defun__shen_4trim_1gubbins Obj // shen.trim-gubbins
var __defun__shen_4space Obj // shen.space
var __defun__shen_4tab Obj // shen.tab
var __defun__shen_4left_1round Obj // shen.left-round
var __defun__shen_4find Obj // shen.find
var __defun__shen_4prefix_2 Obj // shen.prefix?
var __defun__shen_4print_1past_1inputs Obj // shen.print-past-inputs
var __defun__shen_4toplevel__evaluate Obj // shen.toplevel_evaluate
var __defun__shen_4typecheck_1and_1evaluate Obj // shen.typecheck-and-evaluate
var __defun__shen_4pretty_1type Obj // shen.pretty-type
var __defun__shen_4extract_1pvars Obj // shen.extract-pvars
var __defun__shen_4mult__subst Obj // shen.mult_subst

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg64 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
__ctx.Return(reg64)
return
}, 0))
__defun__shen_4repl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg65 := __e.Call(__defun__shen_4credits)
_ = reg65
__ctx.TailApply(__defun__shen_4loop)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.repl", value: __defun__shen_4repl})

__defun__shen_4loop = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg67 := __e.Call(__defun__shen_4initialise__environment)
_ = reg67
reg68 := __e.Call(__defun__shen_4prompt)
_ = reg68
reg69 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4read_1evaluate_1print)
return
}, 0)
reg71 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.TailApply(__defun__shen_4toplevel_1display_1exception, E)
return
}, 1)
reg73 := __e.Try(reg69).Catch(reg71)
_ = reg73
__ctx.TailApply(__defun__shen_4loop)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.loop", value: __defun__shen_4loop})

__defun__shen_4toplevel_1display_1exception = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2992 := __args[0]
_ = V2992
reg75 := PrimErrorToString(V2992)
reg76 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__pr, reg75, reg76)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.toplevel-display-exception", value: __defun__shen_4toplevel_1display_1exception})

__defun__shen_4credits = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg78 := MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n")
reg79 := __e.Call(__defun__stoutput)
reg80 := __e.Call(__defun__shen_4prhush, reg78, reg79)
_ = reg80
reg81 := MakeString("www.shenlanguage.org, ")
reg82 := MakeSymbol("*version*")
reg83 := PrimValue(reg82)
reg84 := MakeString("\n")
reg85 := MakeSymbol("shen.a")
reg86 := __e.Call(__defun__shen_4app, reg83, reg84, reg85)
reg87 := PrimStringConcat(reg81, reg86)
reg88 := __e.Call(__defun__stoutput)
reg89 := __e.Call(__defun__shen_4prhush, reg87, reg88)
_ = reg89
reg90 := MakeString("running under ")
reg91 := MakeSymbol("*language*")
reg92 := PrimValue(reg91)
reg93 := MakeString(", implementation: ")
reg94 := MakeSymbol("*implementation*")
reg95 := PrimValue(reg94)
reg96 := MakeString("")
reg97 := MakeSymbol("shen.a")
reg98 := __e.Call(__defun__shen_4app, reg95, reg96, reg97)
reg99 := PrimStringConcat(reg93, reg98)
reg100 := MakeSymbol("shen.a")
reg101 := __e.Call(__defun__shen_4app, reg92, reg99, reg100)
reg102 := PrimStringConcat(reg90, reg101)
reg103 := __e.Call(__defun__stoutput)
reg104 := __e.Call(__defun__shen_4prhush, reg102, reg103)
_ = reg104
reg105 := MakeString("\nport ")
reg106 := MakeSymbol("*port*")
reg107 := PrimValue(reg106)
reg108 := MakeString(" ported by ")
reg109 := MakeSymbol("*porters*")
reg110 := PrimValue(reg109)
reg111 := MakeString("\n")
reg112 := MakeSymbol("shen.a")
reg113 := __e.Call(__defun__shen_4app, reg110, reg111, reg112)
reg114 := PrimStringConcat(reg108, reg113)
reg115 := MakeSymbol("shen.a")
reg116 := __e.Call(__defun__shen_4app, reg107, reg114, reg115)
reg117 := PrimStringConcat(reg105, reg116)
reg118 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg117, reg118)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.credits", value: __defun__shen_4credits})

__defun__shen_4initialise__environment = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg120 := MakeSymbol("shen.*call*")
reg121 := MakeNumber(0)
reg122 := MakeSymbol("shen.*infs*")
reg123 := MakeNumber(0)
reg124 := MakeSymbol("shen.*process-counter*")
reg125 := MakeNumber(0)
reg126 := MakeSymbol("shen.*catch*")
reg127 := MakeNumber(0)
reg128 := Nil;
reg129 := PrimCons(reg127, reg128)
reg130 := PrimCons(reg126, reg129)
reg131 := PrimCons(reg125, reg130)
reg132 := PrimCons(reg124, reg131)
reg133 := PrimCons(reg123, reg132)
reg134 := PrimCons(reg122, reg133)
reg135 := PrimCons(reg121, reg134)
reg136 := PrimCons(reg120, reg135)
__ctx.TailApply(__defun__shen_4multiple_1set, reg136)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.initialise_environment", value: __defun__shen_4initialise__environment})

__defun__shen_4multiple_1set = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2994 := __args[0]
_ = V2994
reg138 := Nil;
reg139 := PrimEqual(reg138, V2994)
if reg139 == True {
reg140 := Nil;
__ctx.Return(reg140)
return
} else {
reg141 := PrimIsPair(V2994)
var reg148 Obj
if reg141 == True {
reg142 := PrimTail(V2994)
reg143 := PrimIsPair(reg142)
var reg146 Obj
if reg143 == True {
reg144 := True;
reg146 = reg144
} else {
reg145 := False;
reg146 = reg145
}
reg148 = reg146
} else {
reg147 := False;
reg148 = reg147
}
if reg148 == True {
reg149 := PrimHead(V2994)
reg150 := PrimTail(V2994)
reg151 := PrimHead(reg150)
reg152 := PrimSet(reg149, reg151)
_ = reg152
reg153 := PrimTail(V2994)
reg154 := PrimTail(reg153)
__ctx.TailApply(__defun__shen_4multiple_1set, reg154)
return
} else {
reg156 := MakeSymbol("shen.multiple-set")
__ctx.TailApply(__defun__shen_4f__error, reg156)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.multiple-set", value: __defun__shen_4multiple_1set})

__defun__destroy = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2996 := __args[0]
_ = V2996
reg158 := MakeSymbol("symbol")
__ctx.TailApply(__defun__declare, V2996, reg158)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "destroy", value: __defun__destroy})

__defun__shen_4read_1evaluate_1print = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg160 := __e.Call(__defun__shen_4toplineread)
Lineread := reg160
_ = Lineread
reg161 := MakeSymbol("shen.*history*")
reg162 := PrimValue(reg161)
History := reg162
_ = History
reg163 := __e.Call(__defun__shen_4retrieve_1from_1history_1if_1needed, Lineread, History)
NewLineread := reg163
_ = NewLineread
reg164 := __e.Call(__defun__shen_4update__history, NewLineread, History)
NewHistory := reg164
_ = NewHistory
reg165 := __e.Call(__defun__fst, NewLineread)
Parsed := reg165
_ = Parsed
__ctx.TailApply(__defun__shen_4toplevel, Parsed)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.read-evaluate-print", value: __defun__shen_4read_1evaluate_1print})

__defun__shen_4retrieve_1from_1history_1if_1needed = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3008 := __args[0]
_ = V3008
V3009 := __args[1]
_ = V3009
reg167 := __e.Call(__defun__tuple_2, V3008)
var reg187 Obj
if reg167 == True {
reg168 := __e.Call(__defun__snd, V3008)
reg169 := PrimIsPair(reg168)
var reg182 Obj
if reg169 == True {
reg170 := __e.Call(__defun__snd, V3008)
reg171 := PrimHead(reg170)
reg172 := __e.Call(__defun__shen_4space)
reg173 := __e.Call(__defun__shen_4newline)
reg174 := Nil;
reg175 := PrimCons(reg173, reg174)
reg176 := PrimCons(reg172, reg175)
reg177 := __e.Call(__defun__element_2, reg171, reg176)
var reg180 Obj
if reg177 == True {
reg178 := True;
reg180 = reg178
} else {
reg179 := False;
reg180 = reg179
}
reg182 = reg180
} else {
reg181 := False;
reg182 = reg181
}
var reg185 Obj
if reg182 == True {
reg183 := True;
reg185 = reg183
} else {
reg184 := False;
reg185 = reg184
}
reg187 = reg185
} else {
reg186 := False;
reg187 = reg186
}
if reg187 == True {
reg188 := __e.Call(__defun__fst, V3008)
reg189 := __e.Call(__defun__snd, V3008)
reg190 := PrimTail(reg189)
reg191 := __e.Call(__defun___8p, reg188, reg190)
__ctx.TailApply(__defun__shen_4retrieve_1from_1history_1if_1needed, reg191, V3009)
return
} else {
reg193 := __e.Call(__defun__tuple_2, V3008)
var reg243 Obj
if reg193 == True {
reg194 := __e.Call(__defun__snd, V3008)
reg195 := PrimIsPair(reg194)
var reg238 Obj
if reg195 == True {
reg196 := __e.Call(__defun__snd, V3008)
reg197 := PrimTail(reg196)
reg198 := PrimIsPair(reg197)
var reg233 Obj
if reg198 == True {
reg199 := Nil;
reg200 := __e.Call(__defun__snd, V3008)
reg201 := PrimTail(reg200)
reg202 := PrimTail(reg201)
reg203 := PrimEqual(reg199, reg202)
var reg228 Obj
if reg203 == True {
reg204 := PrimIsPair(V3009)
var reg223 Obj
if reg204 == True {
reg205 := __e.Call(__defun__snd, V3008)
reg206 := PrimHead(reg205)
reg207 := __e.Call(__defun__shen_4exclamation)
reg208 := PrimEqual(reg206, reg207)
var reg218 Obj
if reg208 == True {
reg209 := __e.Call(__defun__snd, V3008)
reg210 := PrimTail(reg209)
reg211 := PrimHead(reg210)
reg212 := __e.Call(__defun__shen_4exclamation)
reg213 := PrimEqual(reg211, reg212)
var reg216 Obj
if reg213 == True {
reg214 := True;
reg216 = reg214
} else {
reg215 := False;
reg216 = reg215
}
reg218 = reg216
} else {
reg217 := False;
reg218 = reg217
}
var reg221 Obj
if reg218 == True {
reg219 := True;
reg221 = reg219
} else {
reg220 := False;
reg221 = reg220
}
reg223 = reg221
} else {
reg222 := False;
reg223 = reg222
}
var reg226 Obj
if reg223 == True {
reg224 := True;
reg226 = reg224
} else {
reg225 := False;
reg226 = reg225
}
reg228 = reg226
} else {
reg227 := False;
reg228 = reg227
}
var reg231 Obj
if reg228 == True {
reg229 := True;
reg231 = reg229
} else {
reg230 := False;
reg231 = reg230
}
reg233 = reg231
} else {
reg232 := False;
reg233 = reg232
}
var reg236 Obj
if reg233 == True {
reg234 := True;
reg236 = reg234
} else {
reg235 := False;
reg236 = reg235
}
reg238 = reg236
} else {
reg237 := False;
reg238 = reg237
}
var reg241 Obj
if reg238 == True {
reg239 := True;
reg241 = reg239
} else {
reg240 := False;
reg241 = reg240
}
reg243 = reg241
} else {
reg242 := False;
reg243 = reg242
}
if reg243 == True {
reg244 := PrimHead(V3009)
reg245 := __e.Call(__defun__snd, reg244)
reg246 := __e.Call(__defun__shen_4prbytes, reg245)
PastPrint := reg246
_ = PastPrint
reg247 := PrimHead(V3009)
__ctx.Return(reg247)
return
} else {
reg248 := __e.Call(__defun__tuple_2, V3008)
var reg264 Obj
if reg248 == True {
reg249 := __e.Call(__defun__snd, V3008)
reg250 := PrimIsPair(reg249)
var reg259 Obj
if reg250 == True {
reg251 := __e.Call(__defun__snd, V3008)
reg252 := PrimHead(reg251)
reg253 := __e.Call(__defun__shen_4exclamation)
reg254 := PrimEqual(reg252, reg253)
var reg257 Obj
if reg254 == True {
reg255 := True;
reg257 = reg255
} else {
reg256 := False;
reg257 = reg256
}
reg259 = reg257
} else {
reg258 := False;
reg259 = reg258
}
var reg262 Obj
if reg259 == True {
reg260 := True;
reg262 = reg260
} else {
reg261 := False;
reg262 = reg261
}
reg264 = reg262
} else {
reg263 := False;
reg264 = reg263
}
if reg264 == True {
reg265 := __e.Call(__defun__snd, V3008)
reg266 := PrimTail(reg265)
reg267 := __e.Call(__defun__shen_4make_1key, reg266, V3009)
Key_2 := reg267
_ = Key_2
reg268 := __e.Call(__defun__shen_4find_1past_1inputs, Key_2, V3009)
reg269 := __e.Call(__defun__head, reg268)
Find := reg269
_ = Find
reg270 := __e.Call(__defun__snd, Find)
reg271 := __e.Call(__defun__shen_4prbytes, reg270)
PastPrint := reg271
_ = PastPrint
__ctx.Return(Find)
return
} else {
reg272 := __e.Call(__defun__tuple_2, V3008)
var reg297 Obj
if reg272 == True {
reg273 := __e.Call(__defun__snd, V3008)
reg274 := PrimIsPair(reg273)
var reg292 Obj
if reg274 == True {
reg275 := Nil;
reg276 := __e.Call(__defun__snd, V3008)
reg277 := PrimTail(reg276)
reg278 := PrimEqual(reg275, reg277)
var reg287 Obj
if reg278 == True {
reg279 := __e.Call(__defun__snd, V3008)
reg280 := PrimHead(reg279)
reg281 := __e.Call(__defun__shen_4percent)
reg282 := PrimEqual(reg280, reg281)
var reg285 Obj
if reg282 == True {
reg283 := True;
reg285 = reg283
} else {
reg284 := False;
reg285 = reg284
}
reg287 = reg285
} else {
reg286 := False;
reg287 = reg286
}
var reg290 Obj
if reg287 == True {
reg288 := True;
reg290 = reg288
} else {
reg289 := False;
reg290 = reg289
}
reg292 = reg290
} else {
reg291 := False;
reg292 = reg291
}
var reg295 Obj
if reg292 == True {
reg293 := True;
reg295 = reg293
} else {
reg294 := False;
reg295 = reg294
}
reg297 = reg295
} else {
reg296 := False;
reg297 = reg296
}
if reg297 == True {
reg298 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg299 := True;
__ctx.Return(reg299)
return
}, 1)
reg300 := __e.Call(__defun__reverse, V3009)
reg301 := MakeNumber(0)
reg302 := __e.Call(__defun__shen_4print_1past_1inputs, reg298, reg300, reg301)
_ = reg302
__ctx.TailApply(__defun__abort)
return
} else {
reg304 := __e.Call(__defun__tuple_2, V3008)
var reg320 Obj
if reg304 == True {
reg305 := __e.Call(__defun__snd, V3008)
reg306 := PrimIsPair(reg305)
var reg315 Obj
if reg306 == True {
reg307 := __e.Call(__defun__snd, V3008)
reg308 := PrimHead(reg307)
reg309 := __e.Call(__defun__shen_4percent)
reg310 := PrimEqual(reg308, reg309)
var reg313 Obj
if reg310 == True {
reg311 := True;
reg313 = reg311
} else {
reg312 := False;
reg313 = reg312
}
reg315 = reg313
} else {
reg314 := False;
reg315 = reg314
}
var reg318 Obj
if reg315 == True {
reg316 := True;
reg318 = reg316
} else {
reg317 := False;
reg318 = reg317
}
reg320 = reg318
} else {
reg319 := False;
reg320 = reg319
}
if reg320 == True {
reg321 := __e.Call(__defun__snd, V3008)
reg322 := PrimTail(reg321)
reg323 := __e.Call(__defun__shen_4make_1key, reg322, V3009)
Key_2 := reg323
_ = Key_2
reg324 := __e.Call(__defun__reverse, V3009)
reg325 := MakeNumber(0)
reg326 := __e.Call(__defun__shen_4print_1past_1inputs, Key_2, reg324, reg325)
Pastprint := reg326
_ = Pastprint
__ctx.TailApply(__defun__abort)
return
} else {
__ctx.Return(V3008)
return
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.retrieve-from-history-if-needed", value: __defun__shen_4retrieve_1from_1history_1if_1needed})

__defun__shen_4percent = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg328 := MakeNumber(37)
__ctx.Return(reg328)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.percent", value: __defun__shen_4percent})

__defun__shen_4exclamation = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg329 := MakeNumber(33)
__ctx.Return(reg329)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.exclamation", value: __defun__shen_4exclamation})

__defun__shen_4prbytes = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3011 := __args[0]
_ = V3011
reg330 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Byte := __args[0]
_ = Byte
reg331 := PrimNumberToString(Byte)
reg332 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__pr, reg331, reg332)
return
}, 1)
reg334 := __e.Call(__defun__shen_4for_1each, reg330, V3011)
_ = reg334
reg335 := MakeNumber(1)
__ctx.TailApply(__defun__nl, reg335)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.prbytes", value: __defun__shen_4prbytes})

__defun__shen_4update__history = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3014 := __args[0]
_ = V3014
V3015 := __args[1]
_ = V3015
reg337 := MakeSymbol("shen.*history*")
reg338 := PrimCons(V3014, V3015)
reg339 := PrimSet(reg337, reg338)
__ctx.Return(reg339)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.update_history", value: __defun__shen_4update__history})

__defun__shen_4toplineread = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg340 := __e.Call(__defun__stinput)
reg341 := __e.Call(__defun__shen_4read_1char_1code, reg340)
reg342 := Nil;
__ctx.TailApply(__defun__shen_4toplineread__loop, reg341, reg342)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.toplineread", value: __defun__shen_4toplineread})

__defun__shen_4toplineread__loop = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3019 := __args[0]
_ = V3019
V3020 := __args[1]
_ = V3020
reg344 := __e.Call(__defun__shen_4hat)
reg345 := PrimEqual(V3019, reg344)
if reg345 == True {
reg346 := MakeString("line read aborted")
reg347 := PrimSimpleError(reg346)
__ctx.Return(reg347)
return
} else {
reg348 := __e.Call(__defun__shen_4newline)
reg349 := __e.Call(__defun__shen_4carriage_1return)
reg350 := Nil;
reg351 := PrimCons(reg349, reg350)
reg352 := PrimCons(reg348, reg351)
reg353 := __e.Call(__defun__element_2, V3019, reg352)
if reg353 == True {
reg354 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg356 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg357 := MakeSymbol("shen.nextline")
__ctx.Return(reg357)
return
}, 1)
reg358 := __e.Call(__defun__compile, reg354, V3020, reg356)
Line := reg358
_ = Line
reg359 := __e.Call(__defun__shen_4record_1it, V3020)
It := reg359
_ = It
reg360 := MakeSymbol("shen.nextline")
reg361 := PrimEqual(Line, reg360)
var reg367 Obj
if reg361 == True {
reg362 := True;
reg367 = reg362
} else {
reg363 := __e.Call(__defun__empty_2, Line)
var reg366 Obj
if reg363 == True {
reg364 := True;
reg366 = reg364
} else {
reg365 := False;
reg366 = reg365
}
reg367 = reg366
}
if reg367 == True {
reg368 := __e.Call(__defun__stinput)
reg369 := __e.Call(__defun__shen_4read_1char_1code, reg368)
reg370 := Nil;
reg371 := PrimCons(V3019, reg370)
reg372 := __e.Call(__defun__append, V3020, reg371)
__ctx.TailApply(__defun__shen_4toplineread__loop, reg369, reg372)
return
} else {
__ctx.TailApply(__defun___8p, Line, V3020)
return
}
} else {
reg375 := __e.Call(__defun__stinput)
reg376 := __e.Call(__defun__shen_4read_1char_1code, reg375)
reg377 := MakeNumber(-1)
reg378 := PrimEqual(V3019, reg377)
var reg382 Obj
if reg378 == True {
reg382 = V3020
} else {
reg379 := Nil;
reg380 := PrimCons(V3019, reg379)
reg381 := __e.Call(__defun__append, V3020, reg380)
reg382 = reg381
}
__ctx.TailApply(__defun__shen_4toplineread__loop, reg376, reg382)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.toplineread_loop", value: __defun__shen_4toplineread__loop})

__defun__shen_4hat = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg384 := MakeNumber(94)
__ctx.Return(reg384)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.hat", value: __defun__shen_4hat})

__defun__shen_4newline = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg385 := MakeNumber(10)
__ctx.Return(reg385)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.newline", value: __defun__shen_4newline})

__defun__shen_4carriage_1return = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg386 := MakeNumber(13)
__ctx.Return(reg386)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.carriage-return", value: __defun__shen_4carriage_1return})

__defun__tc = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3026 := __args[0]
_ = V3026
reg387 := MakeSymbol("+")
reg388 := PrimEqual(reg387, V3026)
if reg388 == True {
reg389 := MakeSymbol("shen.*tc*")
reg390 := True;
reg391 := PrimSet(reg389, reg390)
__ctx.Return(reg391)
return
} else {
reg392 := MakeSymbol("-")
reg393 := PrimEqual(reg392, V3026)
if reg393 == True {
reg394 := MakeSymbol("shen.*tc*")
reg395 := False;
reg396 := PrimSet(reg394, reg395)
__ctx.Return(reg396)
return
} else {
reg397 := MakeString("tc expects a + or -")
reg398 := PrimSimpleError(reg397)
__ctx.Return(reg398)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "tc", value: __defun__tc})

__defun__shen_4prompt = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg399 := MakeSymbol("shen.*tc*")
reg400 := PrimValue(reg399)
if reg400 == True {
reg401 := MakeString("\n\n(")
reg402 := MakeSymbol("shen.*history*")
reg403 := PrimValue(reg402)
reg404 := __e.Call(__defun__length, reg403)
reg405 := MakeString("+) ")
reg406 := MakeSymbol("shen.a")
reg407 := __e.Call(__defun__shen_4app, reg404, reg405, reg406)
reg408 := PrimStringConcat(reg401, reg407)
reg409 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg408, reg409)
return
} else {
reg411 := MakeString("\n\n(")
reg412 := MakeSymbol("shen.*history*")
reg413 := PrimValue(reg412)
reg414 := __e.Call(__defun__length, reg413)
reg415 := MakeString("-) ")
reg416 := MakeSymbol("shen.a")
reg417 := __e.Call(__defun__shen_4app, reg414, reg415, reg416)
reg418 := PrimStringConcat(reg411, reg417)
reg419 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg418, reg419)
return
}
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.prompt", value: __defun__shen_4prompt})

__defun__shen_4toplevel = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3028 := __args[0]
_ = V3028
reg421 := MakeSymbol("shen.*tc*")
reg422 := PrimValue(reg421)
__ctx.TailApply(__defun__shen_4toplevel__evaluate, V3028, reg422)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.toplevel", value: __defun__shen_4toplevel})

__defun__shen_4find_1past_1inputs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3031 := __args[0]
_ = V3031
V3032 := __args[1]
_ = V3032
reg424 := __e.Call(__defun__shen_4find, V3031, V3032)
F := reg424
_ = F
reg425 := __e.Call(__defun__empty_2, F)
if reg425 == True {
reg426 := MakeString("input not found\n")
reg427 := PrimSimpleError(reg426)
__ctx.Return(reg427)
return
} else {
__ctx.Return(F)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.find-past-inputs", value: __defun__shen_4find_1past_1inputs})

__defun__shen_4make_1key = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3035 := __args[0]
_ = V3035
V3036 := __args[1]
_ = V3036
reg428 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg430 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg431 := PrimIsPair(E)
if reg431 == True {
reg432 := MakeString("parse error here: ")
reg433 := MakeString("\n")
reg434 := MakeSymbol("shen.s")
reg435 := __e.Call(__defun__shen_4app, E, reg433, reg434)
reg436 := PrimStringConcat(reg432, reg435)
reg437 := PrimSimpleError(reg436)
__ctx.Return(reg437)
return
} else {
reg438 := MakeString("parse error\n")
reg439 := PrimSimpleError(reg438)
__ctx.Return(reg439)
return
}
}, 1)
reg440 := __e.Call(__defun__compile, reg428, V3035, reg430)
reg441 := PrimHead(reg440)
Atom := reg441
_ = Atom
reg442 := PrimIsInteger(Atom)
if reg442 == True {
reg443 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg444 := MakeNumber(1)
reg445 := PrimNumberAdd(Atom, reg444)
reg446 := __e.Call(__defun__reverse, V3036)
reg447 := __e.Call(__defun__nth, reg445, reg446)
reg448 := PrimEqual(X, reg447)
__ctx.Return(reg448)
return
}, 1)
__ctx.Return(reg443)
return
} else {
reg449 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg450 := __e.Call(__defun__snd, X)
reg451 := __e.Call(__defun__shen_4trim_1gubbins, reg450)
__ctx.TailApply(__defun__shen_4prefix_2, V3035, reg451)
return
}, 1)
__ctx.Return(reg449)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.make-key", value: __defun__shen_4make_1key})

__defun__shen_4trim_1gubbins = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3038 := __args[0]
_ = V3038
reg453 := PrimIsPair(V3038)
var reg461 Obj
if reg453 == True {
reg454 := PrimHead(V3038)
reg455 := __e.Call(__defun__shen_4space)
reg456 := PrimEqual(reg454, reg455)
var reg459 Obj
if reg456 == True {
reg457 := True;
reg459 = reg457
} else {
reg458 := False;
reg459 = reg458
}
reg461 = reg459
} else {
reg460 := False;
reg461 = reg460
}
if reg461 == True {
reg462 := PrimTail(V3038)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg462)
return
} else {
reg464 := PrimIsPair(V3038)
var reg472 Obj
if reg464 == True {
reg465 := PrimHead(V3038)
reg466 := __e.Call(__defun__shen_4newline)
reg467 := PrimEqual(reg465, reg466)
var reg470 Obj
if reg467 == True {
reg468 := True;
reg470 = reg468
} else {
reg469 := False;
reg470 = reg469
}
reg472 = reg470
} else {
reg471 := False;
reg472 = reg471
}
if reg472 == True {
reg473 := PrimTail(V3038)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg473)
return
} else {
reg475 := PrimIsPair(V3038)
var reg483 Obj
if reg475 == True {
reg476 := PrimHead(V3038)
reg477 := __e.Call(__defun__shen_4carriage_1return)
reg478 := PrimEqual(reg476, reg477)
var reg481 Obj
if reg478 == True {
reg479 := True;
reg481 = reg479
} else {
reg480 := False;
reg481 = reg480
}
reg483 = reg481
} else {
reg482 := False;
reg483 = reg482
}
if reg483 == True {
reg484 := PrimTail(V3038)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg484)
return
} else {
reg486 := PrimIsPair(V3038)
var reg494 Obj
if reg486 == True {
reg487 := PrimHead(V3038)
reg488 := __e.Call(__defun__shen_4tab)
reg489 := PrimEqual(reg487, reg488)
var reg492 Obj
if reg489 == True {
reg490 := True;
reg492 = reg490
} else {
reg491 := False;
reg492 = reg491
}
reg494 = reg492
} else {
reg493 := False;
reg494 = reg493
}
if reg494 == True {
reg495 := PrimTail(V3038)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg495)
return
} else {
reg497 := PrimIsPair(V3038)
var reg505 Obj
if reg497 == True {
reg498 := PrimHead(V3038)
reg499 := __e.Call(__defun__shen_4left_1round)
reg500 := PrimEqual(reg498, reg499)
var reg503 Obj
if reg500 == True {
reg501 := True;
reg503 = reg501
} else {
reg502 := False;
reg503 = reg502
}
reg505 = reg503
} else {
reg504 := False;
reg505 = reg504
}
if reg505 == True {
reg506 := PrimTail(V3038)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg506)
return
} else {
__ctx.Return(V3038)
return
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.trim-gubbins", value: __defun__shen_4trim_1gubbins})

__defun__shen_4space = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg508 := MakeNumber(32)
__ctx.Return(reg508)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.space", value: __defun__shen_4space})

__defun__shen_4tab = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg509 := MakeNumber(9)
__ctx.Return(reg509)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.tab", value: __defun__shen_4tab})

__defun__shen_4left_1round = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg510 := MakeNumber(40)
__ctx.Return(reg510)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.left-round", value: __defun__shen_4left_1round})

__defun__shen_4find = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3047 := __args[0]
_ = V3047
V3048 := __args[1]
_ = V3048
reg511 := Nil;
reg512 := PrimEqual(reg511, V3048)
if reg512 == True {
reg513 := Nil;
__ctx.Return(reg513)
return
} else {
reg514 := PrimIsPair(V3048)
var reg521 Obj
if reg514 == True {
reg515 := PrimHead(V3048)
reg516 := __e.Call(V3047, reg515)
var reg519 Obj
if reg516 == True {
reg517 := True;
reg519 = reg517
} else {
reg518 := False;
reg519 = reg518
}
reg521 = reg519
} else {
reg520 := False;
reg521 = reg520
}
if reg521 == True {
reg522 := PrimHead(V3048)
reg523 := PrimTail(V3048)
reg524 := __e.Call(__defun__shen_4find, V3047, reg523)
reg525 := PrimCons(reg522, reg524)
__ctx.Return(reg525)
return
} else {
reg526 := PrimIsPair(V3048)
if reg526 == True {
reg527 := PrimTail(V3048)
__ctx.TailApply(__defun__shen_4find, V3047, reg527)
return
} else {
reg529 := MakeSymbol("shen.find")
__ctx.TailApply(__defun__shen_4f__error, reg529)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.find", value: __defun__shen_4find})

__defun__shen_4prefix_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3062 := __args[0]
_ = V3062
V3063 := __args[1]
_ = V3063
reg531 := Nil;
reg532 := PrimEqual(reg531, V3062)
if reg532 == True {
reg533 := True;
__ctx.Return(reg533)
return
} else {
reg534 := PrimIsPair(V3062)
var reg548 Obj
if reg534 == True {
reg535 := PrimIsPair(V3063)
var reg543 Obj
if reg535 == True {
reg536 := PrimHead(V3063)
reg537 := PrimHead(V3062)
reg538 := PrimEqual(reg536, reg537)
var reg541 Obj
if reg538 == True {
reg539 := True;
reg541 = reg539
} else {
reg540 := False;
reg541 = reg540
}
reg543 = reg541
} else {
reg542 := False;
reg543 = reg542
}
var reg546 Obj
if reg543 == True {
reg544 := True;
reg546 = reg544
} else {
reg545 := False;
reg546 = reg545
}
reg548 = reg546
} else {
reg547 := False;
reg548 = reg547
}
if reg548 == True {
reg549 := PrimTail(V3062)
reg550 := PrimTail(V3063)
__ctx.TailApply(__defun__shen_4prefix_2, reg549, reg550)
return
} else {
reg552 := False;
__ctx.Return(reg552)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.prefix?", value: __defun__shen_4prefix_2})

__defun__shen_4print_1past_1inputs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3075 := __args[0]
_ = V3075
V3076 := __args[1]
_ = V3076
V3077 := __args[2]
_ = V3077
reg553 := Nil;
reg554 := PrimEqual(reg553, V3076)
if reg554 == True {
reg555 := MakeSymbol("_")
__ctx.Return(reg555)
return
} else {
reg556 := PrimIsPair(V3076)
var reg564 Obj
if reg556 == True {
reg557 := PrimHead(V3076)
reg558 := __e.Call(V3075, reg557)
reg559 := PrimNot(reg558)
var reg562 Obj
if reg559 == True {
reg560 := True;
reg562 = reg560
} else {
reg561 := False;
reg562 = reg561
}
reg564 = reg562
} else {
reg563 := False;
reg564 = reg563
}
if reg564 == True {
reg565 := PrimTail(V3076)
reg566 := MakeNumber(1)
reg567 := PrimNumberAdd(V3077, reg566)
__ctx.TailApply(__defun__shen_4print_1past_1inputs, V3075, reg565, reg567)
return
} else {
reg569 := PrimIsPair(V3076)
var reg576 Obj
if reg569 == True {
reg570 := PrimHead(V3076)
reg571 := __e.Call(__defun__tuple_2, reg570)
var reg574 Obj
if reg571 == True {
reg572 := True;
reg574 = reg572
} else {
reg573 := False;
reg574 = reg573
}
reg576 = reg574
} else {
reg575 := False;
reg576 = reg575
}
if reg576 == True {
reg577 := MakeString(". ")
reg578 := MakeSymbol("shen.a")
reg579 := __e.Call(__defun__shen_4app, V3077, reg577, reg578)
reg580 := __e.Call(__defun__stoutput)
reg581 := __e.Call(__defun__shen_4prhush, reg579, reg580)
_ = reg581
reg582 := PrimHead(V3076)
reg583 := __e.Call(__defun__snd, reg582)
reg584 := __e.Call(__defun__shen_4prbytes, reg583)
_ = reg584
reg585 := PrimTail(V3076)
reg586 := MakeNumber(1)
reg587 := PrimNumberAdd(V3077, reg586)
__ctx.TailApply(__defun__shen_4print_1past_1inputs, V3075, reg585, reg587)
return
} else {
reg589 := MakeSymbol("shen.print-past-inputs")
__ctx.TailApply(__defun__shen_4f__error, reg589)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.print-past-inputs", value: __defun__shen_4print_1past_1inputs})

__defun__shen_4toplevel__evaluate = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3080 := __args[0]
_ = V3080
V3081 := __args[1]
_ = V3081
reg591 := PrimIsPair(V3080)
var reg632 Obj
if reg591 == True {
reg592 := PrimTail(V3080)
reg593 := PrimIsPair(reg592)
var reg627 Obj
if reg593 == True {
reg594 := MakeSymbol(":")
reg595 := PrimTail(V3080)
reg596 := PrimHead(reg595)
reg597 := PrimEqual(reg594, reg596)
var reg622 Obj
if reg597 == True {
reg598 := PrimTail(V3080)
reg599 := PrimTail(reg598)
reg600 := PrimIsPair(reg599)
var reg617 Obj
if reg600 == True {
reg601 := Nil;
reg602 := PrimTail(V3080)
reg603 := PrimTail(reg602)
reg604 := PrimTail(reg603)
reg605 := PrimEqual(reg601, reg604)
var reg612 Obj
if reg605 == True {
reg606 := True;
reg607 := PrimEqual(reg606, V3081)
var reg610 Obj
if reg607 == True {
reg608 := True;
reg610 = reg608
} else {
reg609 := False;
reg610 = reg609
}
reg612 = reg610
} else {
reg611 := False;
reg612 = reg611
}
var reg615 Obj
if reg612 == True {
reg613 := True;
reg615 = reg613
} else {
reg614 := False;
reg615 = reg614
}
reg617 = reg615
} else {
reg616 := False;
reg617 = reg616
}
var reg620 Obj
if reg617 == True {
reg618 := True;
reg620 = reg618
} else {
reg619 := False;
reg620 = reg619
}
reg622 = reg620
} else {
reg621 := False;
reg622 = reg621
}
var reg625 Obj
if reg622 == True {
reg623 := True;
reg625 = reg623
} else {
reg624 := False;
reg625 = reg624
}
reg627 = reg625
} else {
reg626 := False;
reg627 = reg626
}
var reg630 Obj
if reg627 == True {
reg628 := True;
reg630 = reg628
} else {
reg629 := False;
reg630 = reg629
}
reg632 = reg630
} else {
reg631 := False;
reg632 = reg631
}
if reg632 == True {
reg633 := PrimHead(V3080)
reg634 := PrimTail(V3080)
reg635 := PrimTail(reg634)
reg636 := PrimHead(reg635)
__ctx.TailApply(__defun__shen_4typecheck_1and_1evaluate, reg633, reg636)
return
} else {
reg638 := PrimIsPair(V3080)
var reg645 Obj
if reg638 == True {
reg639 := PrimTail(V3080)
reg640 := PrimIsPair(reg639)
var reg643 Obj
if reg640 == True {
reg641 := True;
reg643 = reg641
} else {
reg642 := False;
reg643 = reg642
}
reg645 = reg643
} else {
reg644 := False;
reg645 = reg644
}
if reg645 == True {
reg646 := PrimHead(V3080)
reg647 := Nil;
reg648 := PrimCons(reg646, reg647)
reg649 := __e.Call(__defun__shen_4toplevel__evaluate, reg648, V3081)
_ = reg649
reg650 := MakeNumber(1)
reg651 := __e.Call(__defun__nl, reg650)
_ = reg651
reg652 := PrimTail(V3080)
__ctx.TailApply(__defun__shen_4toplevel__evaluate, reg652, V3081)
return
} else {
reg654 := PrimIsPair(V3080)
var reg669 Obj
if reg654 == True {
reg655 := Nil;
reg656 := PrimTail(V3080)
reg657 := PrimEqual(reg655, reg656)
var reg664 Obj
if reg657 == True {
reg658 := True;
reg659 := PrimEqual(reg658, V3081)
var reg662 Obj
if reg659 == True {
reg660 := True;
reg662 = reg660
} else {
reg661 := False;
reg662 = reg661
}
reg664 = reg662
} else {
reg663 := False;
reg664 = reg663
}
var reg667 Obj
if reg664 == True {
reg665 := True;
reg667 = reg665
} else {
reg666 := False;
reg667 = reg666
}
reg669 = reg667
} else {
reg668 := False;
reg669 = reg668
}
if reg669 == True {
reg670 := PrimHead(V3080)
reg671 := MakeSymbol("A")
reg672 := __e.Call(__defun__gensym, reg671)
__ctx.TailApply(__defun__shen_4typecheck_1and_1evaluate, reg670, reg672)
return
} else {
reg674 := PrimIsPair(V3080)
var reg689 Obj
if reg674 == True {
reg675 := Nil;
reg676 := PrimTail(V3080)
reg677 := PrimEqual(reg675, reg676)
var reg684 Obj
if reg677 == True {
reg678 := False;
reg679 := PrimEqual(reg678, V3081)
var reg682 Obj
if reg679 == True {
reg680 := True;
reg682 = reg680
} else {
reg681 := False;
reg682 = reg681
}
reg684 = reg682
} else {
reg683 := False;
reg684 = reg683
}
var reg687 Obj
if reg684 == True {
reg685 := True;
reg687 = reg685
} else {
reg686 := False;
reg687 = reg686
}
reg689 = reg687
} else {
reg688 := False;
reg689 = reg688
}
if reg689 == True {
reg690 := PrimHead(V3080)
reg691 := __e.Call(__defun__shen_4eval_1without_1macros, reg690)
Eval := reg691
_ = Eval
__ctx.TailApply(__defun__print, Eval)
return
} else {
reg693 := MakeSymbol("shen.toplevel_evaluate")
__ctx.TailApply(__defun__shen_4f__error, reg693)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.toplevel_evaluate", value: __defun__shen_4toplevel__evaluate})

__defun__shen_4typecheck_1and_1evaluate = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3084 := __args[0]
_ = V3084
V3085 := __args[1]
_ = V3085
reg695 := __e.Call(__defun__shen_4typecheck, V3084, V3085)
Typecheck := reg695
_ = Typecheck
reg696 := False;
reg697 := PrimEqual(Typecheck, reg696)
if reg697 == True {
reg698 := MakeString("type error\n")
reg699 := PrimSimpleError(reg698)
__ctx.Return(reg699)
return
} else {
reg700 := __e.Call(__defun__shen_4eval_1without_1macros, V3084)
Eval := reg700
_ = Eval
reg701 := __e.Call(__defun__shen_4pretty_1type, Typecheck)
Type := reg701
_ = Type
reg702 := MakeString(" : ")
reg703 := MakeString("")
reg704 := MakeSymbol("shen.r")
reg705 := __e.Call(__defun__shen_4app, Type, reg703, reg704)
reg706 := PrimStringConcat(reg702, reg705)
reg707 := MakeSymbol("shen.s")
reg708 := __e.Call(__defun__shen_4app, Eval, reg706, reg707)
reg709 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg708, reg709)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.typecheck-and-evaluate", value: __defun__shen_4typecheck_1and_1evaluate})

__defun__shen_4pretty_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3087 := __args[0]
_ = V3087
reg711 := MakeSymbol("shen.*alphabet*")
reg712 := PrimValue(reg711)
reg713 := __e.Call(__defun__shen_4extract_1pvars, V3087)
__ctx.TailApply(__defun__shen_4mult__subst, reg712, reg713, V3087)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.pretty-type", value: __defun__shen_4pretty_1type})

__defun__shen_4extract_1pvars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3093 := __args[0]
_ = V3093
reg715 := __e.Call(__defun__shen_4pvar_2, V3093)
if reg715 == True {
reg716 := Nil;
reg717 := PrimCons(V3093, reg716)
__ctx.Return(reg717)
return
} else {
reg718 := PrimIsPair(V3093)
if reg718 == True {
reg719 := PrimHead(V3093)
reg720 := __e.Call(__defun__shen_4extract_1pvars, reg719)
reg721 := PrimTail(V3093)
reg722 := __e.Call(__defun__shen_4extract_1pvars, reg721)
__ctx.TailApply(__defun__union, reg720, reg722)
return
} else {
reg724 := Nil;
__ctx.Return(reg724)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.extract-pvars", value: __defun__shen_4extract_1pvars})

__defun__shen_4mult__subst = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3101 := __args[0]
_ = V3101
V3102 := __args[1]
_ = V3102
V3103 := __args[2]
_ = V3103
reg725 := Nil;
reg726 := PrimEqual(reg725, V3101)
if reg726 == True {
__ctx.Return(V3103)
return
} else {
reg727 := Nil;
reg728 := PrimEqual(reg727, V3102)
if reg728 == True {
__ctx.Return(V3103)
return
} else {
reg729 := PrimIsPair(V3101)
var reg735 Obj
if reg729 == True {
reg730 := PrimIsPair(V3102)
var reg733 Obj
if reg730 == True {
reg731 := True;
reg733 = reg731
} else {
reg732 := False;
reg733 = reg732
}
reg735 = reg733
} else {
reg734 := False;
reg735 = reg734
}
if reg735 == True {
reg736 := PrimTail(V3101)
reg737 := PrimTail(V3102)
reg738 := PrimHead(V3101)
reg739 := PrimHead(V3102)
reg740 := __e.Call(__defun__subst, reg738, reg739, V3103)
__ctx.TailApply(__defun__shen_4mult__subst, reg736, reg737, reg740)
return
} else {
reg742 := MakeSymbol("shen.mult_subst")
__ctx.TailApply(__defun__shen_4f__error, reg742)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.mult_subst", value: __defun__shen_4mult__subst})

}
