package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4typecheck Obj // shen.typecheck
var __defun__shen_4curry Obj // shen.curry
var __defun__shen_4special_2 Obj // shen.special?
var __defun__shen_4extraspecial_2 Obj // shen.extraspecial?
var __defun__shen_4t_d Obj // shen.t*
var __defun__shen_4type_1theory_1enabled_2 Obj // shen.type-theory-enabled?
var __defun__enable_1type_1theory Obj // enable-type-theory
var __defun__shen_4prolog_1failure Obj // shen.prolog-failure
var __defun__shen_4maxinfexceeded_2 Obj // shen.maxinfexceeded?
var __defun__shen_4errormaxinfs Obj // shen.errormaxinfs
var __defun__shen_4udefs_d Obj // shen.udefs*
var __defun__shen_4th_d Obj // shen.th*
var __defun__shen_4t_d_1hyps Obj // shen.t*-hyps
var __defun__shen_4show Obj // shen.show
var __defun__shen_4line Obj // shen.line
var __defun__shen_4show_1p Obj // shen.show-p
var __defun__shen_4show_1assumptions Obj // shen.show-assumptions
var __defun__shen_4pause_1for_1user Obj // shen.pause-for-user
var __defun__shen_4typedf_2 Obj // shen.typedf?
var __defun__shen_4sigf Obj // shen.sigf
var __defun__shen_4placeholder Obj // shen.placeholder
var __defun__shen_4base Obj // shen.base
var __defun__shen_4by__hypothesis Obj // shen.by_hypothesis
var __defun__shen_4t_d_1def Obj // shen.t*-def
var __defun__shen_4t_d_1defh Obj // shen.t*-defh
var __defun__shen_4t_d_1defhh Obj // shen.t*-defhh
var __defun__shen_4memo Obj // shen.memo
var __defun__shen_4_5sig_7rules_6 Obj // shen.<sig+rules>
var __defun__shen_4_5non_1ll_1rules_6 Obj // shen.<non-ll-rules>
var __defun__shen_4ue Obj // shen.ue
var __defun__shen_4ue_1sig Obj // shen.ue-sig
var __defun__shen_4ues Obj // shen.ues
var __defun__shen_4ue_2 Obj // shen.ue?
var __defun__shen_4ue_1h_2 Obj // shen.ue-h?
var __defun__shen_4t_d_1rules Obj // shen.t*-rules
var __defun__shen_4t_d_1rule Obj // shen.t*-rule
var __defun__shen_4placeholders Obj // shen.placeholders
var __defun__shen_4newhyps Obj // shen.newhyps
var __defun__shen_4patthyps Obj // shen.patthyps
var __defun__shen_4result_1type Obj // shen.result-type
var __defun__shen_4t_d_1patterns Obj // shen.t*-patterns
var __defun__shen_4t_d_1action Obj // shen.t*-action
var __defun__findall Obj // findall
var __defun__shen_4findallhelp Obj // shen.findallhelp
var __defun__shen_4remember Obj // shen.remember

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19190 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
__ctx.Return(reg19190)
return
}, 0))
__defun__shen_4typecheck = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2701 := __args[0]
_ = V2701
V2702 := __args[1]
_ = V2702
reg19191 := __e.Call(__defun__shen_4curry, V2701)
Curry := reg19191
_ = Curry
reg19192 := __e.Call(__defun__shen_4start_1new_1prolog_1process)
ProcessN := reg19192
_ = ProcessN
reg19193 := __e.Call(__defun__shen_4curry_1type, V2702)
reg19194 := __e.Call(__defun__shen_4demodulate, reg19193)
reg19195 := __e.Call(__defun__shen_4insert_1prolog_1variables, reg19194, ProcessN)
Type := reg19195
_ = Type
reg19196 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19197 := MakeSymbol("shen.void")
__ctx.TailApply(__defun__return, Type, ProcessN, reg19197)
return
}, 0)
Continuation := reg19196
_ = Continuation
reg19199 := MakeSymbol(":")
reg19200 := Nil;
reg19201 := PrimCons(Type, reg19200)
reg19202 := PrimCons(reg19199, reg19201)
reg19203 := PrimCons(Curry, reg19202)
reg19204 := Nil;
__ctx.TailApply(__defun__shen_4t_d, reg19203, reg19204, ProcessN, Continuation)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.typecheck", value: __defun__shen_4typecheck})

__defun__shen_4curry = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2704 := __args[0]
_ = V2704
reg19206 := PrimIsPair(V2704)
var reg19213 Obj
if reg19206 == True {
reg19207 := PrimHead(V2704)
reg19208 := __e.Call(__defun__shen_4special_2, reg19207)
var reg19211 Obj
if reg19208 == True {
reg19209 := True;
reg19211 = reg19209
} else {
reg19210 := False;
reg19211 = reg19210
}
reg19213 = reg19211
} else {
reg19212 := False;
reg19213 = reg19212
}
if reg19213 == True {
reg19214 := PrimHead(V2704)
reg19215 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4curry, Y)
return
}, 1)
reg19217 := PrimTail(V2704)
reg19218 := __e.Call(__defun__map, reg19215, reg19217)
reg19219 := PrimCons(reg19214, reg19218)
__ctx.Return(reg19219)
return
} else {
reg19220 := PrimIsPair(V2704)
var reg19234 Obj
if reg19220 == True {
reg19221 := PrimTail(V2704)
reg19222 := PrimIsPair(reg19221)
var reg19229 Obj
if reg19222 == True {
reg19223 := PrimHead(V2704)
reg19224 := __e.Call(__defun__shen_4extraspecial_2, reg19223)
var reg19227 Obj
if reg19224 == True {
reg19225 := True;
reg19227 = reg19225
} else {
reg19226 := False;
reg19227 = reg19226
}
reg19229 = reg19227
} else {
reg19228 := False;
reg19229 = reg19228
}
var reg19232 Obj
if reg19229 == True {
reg19230 := True;
reg19232 = reg19230
} else {
reg19231 := False;
reg19232 = reg19231
}
reg19234 = reg19232
} else {
reg19233 := False;
reg19234 = reg19233
}
if reg19234 == True {
__ctx.Return(V2704)
return
} else {
reg19235 := PrimIsPair(V2704)
var reg19268 Obj
if reg19235 == True {
reg19236 := MakeSymbol("type")
reg19237 := PrimHead(V2704)
reg19238 := PrimEqual(reg19236, reg19237)
var reg19263 Obj
if reg19238 == True {
reg19239 := PrimTail(V2704)
reg19240 := PrimIsPair(reg19239)
var reg19258 Obj
if reg19240 == True {
reg19241 := PrimTail(V2704)
reg19242 := PrimTail(reg19241)
reg19243 := PrimIsPair(reg19242)
var reg19253 Obj
if reg19243 == True {
reg19244 := Nil;
reg19245 := PrimTail(V2704)
reg19246 := PrimTail(reg19245)
reg19247 := PrimTail(reg19246)
reg19248 := PrimEqual(reg19244, reg19247)
var reg19251 Obj
if reg19248 == True {
reg19249 := True;
reg19251 = reg19249
} else {
reg19250 := False;
reg19251 = reg19250
}
reg19253 = reg19251
} else {
reg19252 := False;
reg19253 = reg19252
}
var reg19256 Obj
if reg19253 == True {
reg19254 := True;
reg19256 = reg19254
} else {
reg19255 := False;
reg19256 = reg19255
}
reg19258 = reg19256
} else {
reg19257 := False;
reg19258 = reg19257
}
var reg19261 Obj
if reg19258 == True {
reg19259 := True;
reg19261 = reg19259
} else {
reg19260 := False;
reg19261 = reg19260
}
reg19263 = reg19261
} else {
reg19262 := False;
reg19263 = reg19262
}
var reg19266 Obj
if reg19263 == True {
reg19264 := True;
reg19266 = reg19264
} else {
reg19265 := False;
reg19266 = reg19265
}
reg19268 = reg19266
} else {
reg19267 := False;
reg19268 = reg19267
}
if reg19268 == True {
reg19269 := MakeSymbol("type")
reg19270 := PrimTail(V2704)
reg19271 := PrimHead(reg19270)
reg19272 := __e.Call(__defun__shen_4curry, reg19271)
reg19273 := PrimTail(V2704)
reg19274 := PrimTail(reg19273)
reg19275 := PrimCons(reg19272, reg19274)
reg19276 := PrimCons(reg19269, reg19275)
__ctx.Return(reg19276)
return
} else {
reg19277 := PrimIsPair(V2704)
var reg19292 Obj
if reg19277 == True {
reg19278 := PrimTail(V2704)
reg19279 := PrimIsPair(reg19278)
var reg19287 Obj
if reg19279 == True {
reg19280 := PrimTail(V2704)
reg19281 := PrimTail(reg19280)
reg19282 := PrimIsPair(reg19281)
var reg19285 Obj
if reg19282 == True {
reg19283 := True;
reg19285 = reg19283
} else {
reg19284 := False;
reg19285 = reg19284
}
reg19287 = reg19285
} else {
reg19286 := False;
reg19287 = reg19286
}
var reg19290 Obj
if reg19287 == True {
reg19288 := True;
reg19290 = reg19288
} else {
reg19289 := False;
reg19290 = reg19289
}
reg19292 = reg19290
} else {
reg19291 := False;
reg19292 = reg19291
}
if reg19292 == True {
reg19293 := PrimHead(V2704)
reg19294 := PrimTail(V2704)
reg19295 := PrimHead(reg19294)
reg19296 := Nil;
reg19297 := PrimCons(reg19295, reg19296)
reg19298 := PrimCons(reg19293, reg19297)
reg19299 := PrimTail(V2704)
reg19300 := PrimTail(reg19299)
reg19301 := PrimCons(reg19298, reg19300)
__ctx.TailApply(__defun__shen_4curry, reg19301)
return
} else {
reg19303 := PrimIsPair(V2704)
var reg19319 Obj
if reg19303 == True {
reg19304 := PrimTail(V2704)
reg19305 := PrimIsPair(reg19304)
var reg19314 Obj
if reg19305 == True {
reg19306 := Nil;
reg19307 := PrimTail(V2704)
reg19308 := PrimTail(reg19307)
reg19309 := PrimEqual(reg19306, reg19308)
var reg19312 Obj
if reg19309 == True {
reg19310 := True;
reg19312 = reg19310
} else {
reg19311 := False;
reg19312 = reg19311
}
reg19314 = reg19312
} else {
reg19313 := False;
reg19314 = reg19313
}
var reg19317 Obj
if reg19314 == True {
reg19315 := True;
reg19317 = reg19315
} else {
reg19316 := False;
reg19317 = reg19316
}
reg19319 = reg19317
} else {
reg19318 := False;
reg19319 = reg19318
}
if reg19319 == True {
reg19320 := PrimHead(V2704)
reg19321 := __e.Call(__defun__shen_4curry, reg19320)
reg19322 := PrimTail(V2704)
reg19323 := PrimHead(reg19322)
reg19324 := __e.Call(__defun__shen_4curry, reg19323)
reg19325 := Nil;
reg19326 := PrimCons(reg19324, reg19325)
reg19327 := PrimCons(reg19321, reg19326)
__ctx.Return(reg19327)
return
} else {
__ctx.Return(V2704)
return
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.curry", value: __defun__shen_4curry})

__defun__shen_4special_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2706 := __args[0]
_ = V2706
reg19328 := MakeSymbol("shen.*special*")
reg19329 := PrimValue(reg19328)
__ctx.TailApply(__defun__element_2, V2706, reg19329)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.special?", value: __defun__shen_4special_2})

__defun__shen_4extraspecial_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2708 := __args[0]
_ = V2708
reg19331 := MakeSymbol("shen.*extraspecial*")
reg19332 := PrimValue(reg19331)
__ctx.TailApply(__defun__element_2, V2708, reg19332)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.extraspecial?", value: __defun__shen_4extraspecial_2})

__defun__shen_4t_d = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2713 := __args[0]
_ = V2713
V2714 := __args[1]
_ = V2714
V2715 := __args[2]
_ = V2715
V2716 := __args[3]
_ = V2716
reg19334 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg19334
_ = Throwcontrol
reg19335 := __e.Call(__defun__shen_4newpv, V2715)
Error := reg19335
_ = Error
reg19336 := __e.Call(__defun__shen_4incinfs)
_ = reg19336
reg19337 := __e.Call(__defun__shen_4maxinfexceeded_2)
reg19338 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19339 := __e.Call(__defun__shen_4errormaxinfs)
__ctx.TailApply(__defun__bind, Error, reg19339, V2715, V2716)
return
}, 0)
reg19341 := __e.Call(__defun__fwhen, reg19337, V2715, reg19338)
Case := reg19341
_ = Case
reg19342 := False;
reg19343 := PrimEqual(Case, reg19342)
var reg19403 Obj
if reg19343 == True {
reg19344 := __e.Call(__defun__shen_4lazyderef, V2713, V2715)
V2693 := reg19344
_ = V2693
reg19345 := MakeSymbol("fail")
reg19346 := PrimEqual(reg19345, V2693)
var reg19352 Obj
if reg19346 == True {
reg19347 := __e.Call(__defun__shen_4incinfs)
_ = reg19347
reg19348 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4prolog_1failure, V2715, V2716)
return
}, 0)
reg19350 := __e.Call(__defun__cut, Throwcontrol, V2715, reg19348)
reg19352 = reg19350
} else {
reg19351 := False;
reg19352 = reg19351
}
Case := reg19352
_ = Case
reg19353 := False;
reg19354 := PrimEqual(Case, reg19353)
var reg19402 Obj
if reg19354 == True {
reg19355 := __e.Call(__defun__shen_4lazyderef, V2713, V2715)
V2694 := reg19355
_ = V2694
reg19356 := PrimIsPair(V2694)
var reg19389 Obj
if reg19356 == True {
reg19357 := PrimHead(V2694)
X := reg19357
_ = X
reg19358 := PrimTail(V2694)
reg19359 := __e.Call(__defun__shen_4lazyderef, reg19358, V2715)
V2695 := reg19359
_ = V2695
reg19360 := PrimIsPair(V2695)
var reg19387 Obj
if reg19360 == True {
reg19361 := PrimHead(V2695)
reg19362 := __e.Call(__defun__shen_4lazyderef, reg19361, V2715)
V2696 := reg19362
_ = V2696
reg19363 := MakeSymbol(":")
reg19364 := PrimEqual(reg19363, V2696)
var reg19385 Obj
if reg19364 == True {
reg19365 := PrimTail(V2695)
reg19366 := __e.Call(__defun__shen_4lazyderef, reg19365, V2715)
V2697 := reg19366
_ = V2697
reg19367 := PrimIsPair(V2697)
var reg19383 Obj
if reg19367 == True {
reg19368 := PrimHead(V2697)
A := reg19368
_ = A
reg19369 := PrimTail(V2697)
reg19370 := __e.Call(__defun__shen_4lazyderef, reg19369, V2715)
V2698 := reg19370
_ = V2698
reg19371 := Nil;
reg19372 := PrimEqual(reg19371, V2698)
var reg19381 Obj
if reg19372 == True {
reg19373 := __e.Call(__defun__shen_4incinfs)
_ = reg19373
reg19374 := __e.Call(__defun__shen_4type_1theory_1enabled_2)
reg19375 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19376 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, X, A, V2714, V2715, V2716)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2715, reg19376)
return
}, 0)
reg19379 := __e.Call(__defun__fwhen, reg19374, V2715, reg19375)
reg19381 = reg19379
} else {
reg19380 := False;
reg19381 = reg19380
}
reg19383 = reg19381
} else {
reg19382 := False;
reg19383 = reg19382
}
reg19385 = reg19383
} else {
reg19384 := False;
reg19385 = reg19384
}
reg19387 = reg19385
} else {
reg19386 := False;
reg19387 = reg19386
}
reg19389 = reg19387
} else {
reg19388 := False;
reg19389 = reg19388
}
Case := reg19389
_ = Case
reg19390 := False;
reg19391 := PrimEqual(Case, reg19390)
var reg19401 Obj
if reg19391 == True {
reg19392 := __e.Call(__defun__shen_4newpv, V2715)
Datatypes := reg19392
_ = Datatypes
reg19393 := __e.Call(__defun__shen_4incinfs)
_ = reg19393
reg19394 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19395 := MakeSymbol("shen.*datatypes*")
reg19396 := PrimValue(reg19395)
reg19397 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4udefs_d, V2713, V2714, Datatypes, V2715, V2716)
return
}, 0)
__ctx.TailApply(__defun__bind, Datatypes, reg19396, V2715, reg19397)
return
}, 0)
reg19400 := __e.Call(__defun__shen_4show, V2713, V2714, V2715, reg19394)
reg19401 = reg19400
} else {
reg19401 = Case
}
reg19402 = reg19401
} else {
reg19402 = Case
}
reg19403 = reg19402
} else {
reg19403 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg19403)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.t*", value: __defun__shen_4t_d})

__defun__shen_4type_1theory_1enabled_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19405 := MakeSymbol("shen.*shen-type-theory-enabled?*")
reg19406 := PrimValue(reg19405)
__ctx.Return(reg19406)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.type-theory-enabled?", value: __defun__shen_4type_1theory_1enabled_2})

__defun__enable_1type_1theory = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2722 := __args[0]
_ = V2722
reg19407 := MakeSymbol("+")
reg19408 := PrimEqual(reg19407, V2722)
if reg19408 == True {
reg19409 := MakeSymbol("shen.*shen-type-theory-enabled?*")
reg19410 := True;
reg19411 := PrimSet(reg19409, reg19410)
__ctx.Return(reg19411)
return
} else {
reg19412 := MakeSymbol("-")
reg19413 := PrimEqual(reg19412, V2722)
if reg19413 == True {
reg19414 := MakeSymbol("shen.*shen-type-theory-enabled?*")
reg19415 := False;
reg19416 := PrimSet(reg19414, reg19415)
__ctx.Return(reg19416)
return
} else {
reg19417 := MakeString("enable-type-theory expects a + or a -\n")
reg19418 := PrimSimpleError(reg19417)
__ctx.Return(reg19418)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "enable-type-theory", value: __defun__enable_1type_1theory})

__defun__shen_4prolog_1failure = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2733 := __args[0]
_ = V2733
V2734 := __args[1]
_ = V2734
reg19419 := False;
__ctx.Return(reg19419)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.prolog-failure", value: __defun__shen_4prolog_1failure})

__defun__shen_4maxinfexceeded_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19420 := __e.Call(__defun__inferences)
reg19421 := MakeSymbol("shen.*maxinferences*")
reg19422 := PrimValue(reg19421)
reg19423 := PrimGreatThan(reg19420, reg19422)
__ctx.Return(reg19423)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.maxinfexceeded?", value: __defun__shen_4maxinfexceeded_2})

__defun__shen_4errormaxinfs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19424 := MakeString("maximum inferences exceeded~%")
reg19425 := PrimSimpleError(reg19424)
__ctx.Return(reg19425)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.errormaxinfs", value: __defun__shen_4errormaxinfs})

__defun__shen_4udefs_d = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2740 := __args[0]
_ = V2740
V2741 := __args[1]
_ = V2741
V2742 := __args[2]
_ = V2742
V2743 := __args[3]
_ = V2743
V2744 := __args[4]
_ = V2744
reg19426 := __e.Call(__defun__shen_4lazyderef, V2742, V2743)
V2689 := reg19426
_ = V2689
reg19427 := PrimIsPair(V2689)
var reg19436 Obj
if reg19427 == True {
reg19428 := PrimHead(V2689)
D := reg19428
_ = D
reg19429 := __e.Call(__defun__shen_4incinfs)
_ = reg19429
reg19430 := Nil;
reg19431 := PrimCons(V2741, reg19430)
reg19432 := PrimCons(V2740, reg19431)
reg19433 := PrimCons(D, reg19432)
reg19434 := __e.Call(__defun__call, reg19433, V2743, V2744)
reg19436 = reg19434
} else {
reg19435 := False;
reg19436 = reg19435
}
Case := reg19436
_ = Case
reg19437 := False;
reg19438 := PrimEqual(Case, reg19437)
if reg19438 == True {
reg19439 := __e.Call(__defun__shen_4lazyderef, V2742, V2743)
V2690 := reg19439
_ = V2690
reg19440 := PrimIsPair(V2690)
if reg19440 == True {
reg19441 := PrimTail(V2690)
Ds := reg19441
_ = Ds
reg19442 := __e.Call(__defun__shen_4incinfs)
_ = reg19442
__ctx.TailApply(__defun__shen_4udefs_d, V2740, V2741, Ds, V2743, V2744)
return
} else {
reg19444 := False;
__ctx.Return(reg19444)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.udefs*", value: __defun__shen_4udefs_d})

__defun__shen_4th_d = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2750 := __args[0]
_ = V2750
V2751 := __args[1]
_ = V2751
V2752 := __args[2]
_ = V2752
V2753 := __args[3]
_ = V2753
V2754 := __args[4]
_ = V2754
reg19445 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg19445
_ = Throwcontrol
reg19446 := __e.Call(__defun__shen_4incinfs)
_ = reg19446
reg19447 := MakeSymbol(":")
reg19448 := Nil;
reg19449 := PrimCons(V2751, reg19448)
reg19450 := PrimCons(reg19447, reg19449)
reg19451 := PrimCons(V2750, reg19450)
reg19452 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19453 := False;
__ctx.TailApply(__defun__fwhen, reg19453, V2753, V2754)
return
}, 0)
reg19455 := __e.Call(__defun__shen_4show, reg19451, V2752, V2753, reg19452)
Case := reg19455
_ = Case
reg19456 := False;
reg19457 := PrimEqual(Case, reg19456)
var reg20876 Obj
if reg19457 == True {
reg19458 := __e.Call(__defun__shen_4newpv, V2753)
F := reg19458
_ = F
reg19459 := __e.Call(__defun__shen_4incinfs)
_ = reg19459
reg19460 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
reg19461 := __e.Call(__defun__shen_4typedf_2, reg19460)
reg19462 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19463 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
reg19464 := __e.Call(__defun__shen_4sigf, reg19463)
reg19465 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19466 := Nil;
reg19467 := PrimCons(V2751, reg19466)
reg19468 := PrimCons(F, reg19467)
__ctx.TailApply(__defun__call, reg19468, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, F, reg19464, V2753, reg19465)
return
}, 0)
reg19471 := __e.Call(__defun__fwhen, reg19461, V2753, reg19462)
Case := reg19471
_ = Case
reg19472 := False;
reg19473 := PrimEqual(Case, reg19472)
var reg20875 Obj
if reg19473 == True {
reg19474 := __e.Call(__defun__shen_4incinfs)
_ = reg19474
reg19475 := __e.Call(__defun__shen_4base, V2750, V2751, V2753, V2754)
Case := reg19475
_ = Case
reg19476 := False;
reg19477 := PrimEqual(Case, reg19476)
var reg20874 Obj
if reg19477 == True {
reg19478 := __e.Call(__defun__shen_4incinfs)
_ = reg19478
reg19479 := __e.Call(__defun__shen_4by__hypothesis, V2750, V2751, V2752, V2753, V2754)
Case := reg19479
_ = Case
reg19480 := False;
reg19481 := PrimEqual(Case, reg19480)
var reg20873 Obj
if reg19481 == True {
reg19482 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2585 := reg19482
_ = V2585
reg19483 := PrimIsPair(V2585)
var reg19498 Obj
if reg19483 == True {
reg19484 := PrimHead(V2585)
F := reg19484
_ = F
reg19485 := PrimTail(V2585)
reg19486 := __e.Call(__defun__shen_4lazyderef, reg19485, V2753)
V2586 := reg19486
_ = V2586
reg19487 := Nil;
reg19488 := PrimEqual(reg19487, V2586)
var reg19496 Obj
if reg19488 == True {
reg19489 := __e.Call(__defun__shen_4incinfs)
_ = reg19489
reg19490 := MakeSymbol("-->")
reg19491 := Nil;
reg19492 := PrimCons(V2751, reg19491)
reg19493 := PrimCons(reg19490, reg19492)
reg19494 := __e.Call(__defun__shen_4th_d, F, reg19493, V2752, V2753, V2754)
reg19496 = reg19494
} else {
reg19495 := False;
reg19496 = reg19495
}
reg19498 = reg19496
} else {
reg19497 := False;
reg19498 = reg19497
}
Case := reg19498
_ = Case
reg19499 := False;
reg19500 := PrimEqual(Case, reg19499)
var reg20872 Obj
if reg19500 == True {
reg19501 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2587 := reg19501
_ = V2587
reg19502 := PrimIsPair(V2587)
var reg19527 Obj
if reg19502 == True {
reg19503 := PrimHead(V2587)
F := reg19503
_ = F
reg19504 := PrimTail(V2587)
reg19505 := __e.Call(__defun__shen_4lazyderef, reg19504, V2753)
V2588 := reg19505
_ = V2588
reg19506 := PrimIsPair(V2588)
var reg19525 Obj
if reg19506 == True {
reg19507 := PrimHead(V2588)
X := reg19507
_ = X
reg19508 := PrimTail(V2588)
reg19509 := __e.Call(__defun__shen_4lazyderef, reg19508, V2753)
V2589 := reg19509
_ = V2589
reg19510 := Nil;
reg19511 := PrimEqual(reg19510, V2589)
var reg19523 Obj
if reg19511 == True {
reg19512 := __e.Call(__defun__shen_4newpv, V2753)
B := reg19512
_ = B
reg19513 := __e.Call(__defun__shen_4incinfs)
_ = reg19513
reg19514 := MakeSymbol("-->")
reg19515 := Nil;
reg19516 := PrimCons(V2751, reg19515)
reg19517 := PrimCons(reg19514, reg19516)
reg19518 := PrimCons(B, reg19517)
reg19519 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, X, B, V2752, V2753, V2754)
return
}, 0)
reg19521 := __e.Call(__defun__shen_4th_d, F, reg19518, V2752, V2753, reg19519)
reg19523 = reg19521
} else {
reg19522 := False;
reg19523 = reg19522
}
reg19525 = reg19523
} else {
reg19524 := False;
reg19525 = reg19524
}
reg19527 = reg19525
} else {
reg19526 := False;
reg19527 = reg19526
}
Case := reg19527
_ = Case
reg19528 := False;
reg19529 := PrimEqual(Case, reg19528)
var reg20871 Obj
if reg19529 == True {
reg19530 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2590 := reg19530
_ = V2590
reg19531 := PrimIsPair(V2590)
var reg19685 Obj
if reg19531 == True {
reg19532 := PrimHead(V2590)
reg19533 := __e.Call(__defun__shen_4lazyderef, reg19532, V2753)
V2591 := reg19533
_ = V2591
reg19534 := MakeSymbol("cons")
reg19535 := PrimEqual(reg19534, V2591)
var reg19683 Obj
if reg19535 == True {
reg19536 := PrimTail(V2590)
reg19537 := __e.Call(__defun__shen_4lazyderef, reg19536, V2753)
V2592 := reg19537
_ = V2592
reg19538 := PrimIsPair(V2592)
var reg19681 Obj
if reg19538 == True {
reg19539 := PrimHead(V2592)
X := reg19539
_ = X
reg19540 := PrimTail(V2592)
reg19541 := __e.Call(__defun__shen_4lazyderef, reg19540, V2753)
V2593 := reg19541
_ = V2593
reg19542 := PrimIsPair(V2593)
var reg19679 Obj
if reg19542 == True {
reg19543 := PrimHead(V2593)
Y := reg19543
_ = Y
reg19544 := PrimTail(V2593)
reg19545 := __e.Call(__defun__shen_4lazyderef, reg19544, V2753)
V2594 := reg19545
_ = V2594
reg19546 := Nil;
reg19547 := PrimEqual(reg19546, V2594)
var reg19677 Obj
if reg19547 == True {
reg19548 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2595 := reg19548
_ = V2595
reg19549 := PrimIsPair(V2595)
var reg19675 Obj
if reg19549 == True {
reg19550 := PrimHead(V2595)
reg19551 := __e.Call(__defun__shen_4lazyderef, reg19550, V2753)
V2596 := reg19551
_ = V2596
reg19552 := MakeSymbol("list")
reg19553 := PrimEqual(reg19552, V2596)
var reg19656 Obj
if reg19553 == True {
reg19554 := PrimTail(V2595)
reg19555 := __e.Call(__defun__shen_4lazyderef, reg19554, V2753)
V2597 := reg19555
_ = V2597
reg19556 := PrimIsPair(V2597)
var reg19601 Obj
if reg19556 == True {
reg19557 := PrimHead(V2597)
A := reg19557
_ = A
reg19558 := PrimTail(V2597)
reg19559 := __e.Call(__defun__shen_4lazyderef, reg19558, V2753)
V2598 := reg19559
_ = V2598
reg19560 := Nil;
reg19561 := PrimEqual(reg19560, V2598)
var reg19584 Obj
if reg19561 == True {
reg19562 := __e.Call(__defun__shen_4incinfs)
_ = reg19562
reg19563 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19564 := MakeSymbol("list")
reg19565 := Nil;
reg19566 := PrimCons(A, reg19565)
reg19567 := PrimCons(reg19564, reg19566)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19567, V2752, V2753, V2754)
return
}, 0)
reg19569 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19563)
reg19584 = reg19569
} else {
reg19570 := __e.Call(__defun__shen_4pvar_2, V2598)
var reg19583 Obj
if reg19570 == True {
reg19571 := Nil;
reg19572 := __e.Call(__defun__shen_4bindv, V2598, reg19571, V2753)
_ = reg19572
reg19573 := __e.Call(__defun__shen_4incinfs)
_ = reg19573
reg19574 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19575 := MakeSymbol("list")
reg19576 := Nil;
reg19577 := PrimCons(A, reg19576)
reg19578 := PrimCons(reg19575, reg19577)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19578, V2752, V2753, V2754)
return
}, 0)
reg19580 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19574)
Result := reg19580
_ = Result
reg19581 := __e.Call(__defun__shen_4unbindv, V2598, V2753)
_ = reg19581
reg19583 = Result
} else {
reg19582 := False;
reg19583 = reg19582
}
reg19584 = reg19583
}
reg19601 = reg19584
} else {
reg19585 := __e.Call(__defun__shen_4pvar_2, V2597)
var reg19600 Obj
if reg19585 == True {
reg19586 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19586
_ = A
reg19587 := Nil;
reg19588 := PrimCons(A, reg19587)
reg19589 := __e.Call(__defun__shen_4bindv, V2597, reg19588, V2753)
_ = reg19589
reg19590 := __e.Call(__defun__shen_4incinfs)
_ = reg19590
reg19591 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19592 := MakeSymbol("list")
reg19593 := Nil;
reg19594 := PrimCons(A, reg19593)
reg19595 := PrimCons(reg19592, reg19594)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19595, V2752, V2753, V2754)
return
}, 0)
reg19597 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19591)
Result := reg19597
_ = Result
reg19598 := __e.Call(__defun__shen_4unbindv, V2597, V2753)
_ = reg19598
reg19600 = Result
} else {
reg19599 := False;
reg19600 = reg19599
}
reg19601 = reg19600
}
reg19656 = reg19601
} else {
reg19602 := __e.Call(__defun__shen_4pvar_2, V2596)
var reg19655 Obj
if reg19602 == True {
reg19603 := MakeSymbol("list")
reg19604 := __e.Call(__defun__shen_4bindv, V2596, reg19603, V2753)
_ = reg19604
reg19605 := PrimTail(V2595)
reg19606 := __e.Call(__defun__shen_4lazyderef, reg19605, V2753)
V2599 := reg19606
_ = V2599
reg19607 := PrimIsPair(V2599)
var reg19652 Obj
if reg19607 == True {
reg19608 := PrimHead(V2599)
A := reg19608
_ = A
reg19609 := PrimTail(V2599)
reg19610 := __e.Call(__defun__shen_4lazyderef, reg19609, V2753)
V2600 := reg19610
_ = V2600
reg19611 := Nil;
reg19612 := PrimEqual(reg19611, V2600)
var reg19635 Obj
if reg19612 == True {
reg19613 := __e.Call(__defun__shen_4incinfs)
_ = reg19613
reg19614 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19615 := MakeSymbol("list")
reg19616 := Nil;
reg19617 := PrimCons(A, reg19616)
reg19618 := PrimCons(reg19615, reg19617)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19618, V2752, V2753, V2754)
return
}, 0)
reg19620 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19614)
reg19635 = reg19620
} else {
reg19621 := __e.Call(__defun__shen_4pvar_2, V2600)
var reg19634 Obj
if reg19621 == True {
reg19622 := Nil;
reg19623 := __e.Call(__defun__shen_4bindv, V2600, reg19622, V2753)
_ = reg19623
reg19624 := __e.Call(__defun__shen_4incinfs)
_ = reg19624
reg19625 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19626 := MakeSymbol("list")
reg19627 := Nil;
reg19628 := PrimCons(A, reg19627)
reg19629 := PrimCons(reg19626, reg19628)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19629, V2752, V2753, V2754)
return
}, 0)
reg19631 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19625)
Result := reg19631
_ = Result
reg19632 := __e.Call(__defun__shen_4unbindv, V2600, V2753)
_ = reg19632
reg19634 = Result
} else {
reg19633 := False;
reg19634 = reg19633
}
reg19635 = reg19634
}
reg19652 = reg19635
} else {
reg19636 := __e.Call(__defun__shen_4pvar_2, V2599)
var reg19651 Obj
if reg19636 == True {
reg19637 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19637
_ = A
reg19638 := Nil;
reg19639 := PrimCons(A, reg19638)
reg19640 := __e.Call(__defun__shen_4bindv, V2599, reg19639, V2753)
_ = reg19640
reg19641 := __e.Call(__defun__shen_4incinfs)
_ = reg19641
reg19642 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19643 := MakeSymbol("list")
reg19644 := Nil;
reg19645 := PrimCons(A, reg19644)
reg19646 := PrimCons(reg19643, reg19645)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19646, V2752, V2753, V2754)
return
}, 0)
reg19648 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19642)
Result := reg19648
_ = Result
reg19649 := __e.Call(__defun__shen_4unbindv, V2599, V2753)
_ = reg19649
reg19651 = Result
} else {
reg19650 := False;
reg19651 = reg19650
}
reg19652 = reg19651
}
Result := reg19652
_ = Result
reg19653 := __e.Call(__defun__shen_4unbindv, V2596, V2753)
_ = reg19653
reg19655 = Result
} else {
reg19654 := False;
reg19655 = reg19654
}
reg19656 = reg19655
}
reg19675 = reg19656
} else {
reg19657 := __e.Call(__defun__shen_4pvar_2, V2595)
var reg19674 Obj
if reg19657 == True {
reg19658 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19658
_ = A
reg19659 := MakeSymbol("list")
reg19660 := Nil;
reg19661 := PrimCons(A, reg19660)
reg19662 := PrimCons(reg19659, reg19661)
reg19663 := __e.Call(__defun__shen_4bindv, V2595, reg19662, V2753)
_ = reg19663
reg19664 := __e.Call(__defun__shen_4incinfs)
_ = reg19664
reg19665 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19666 := MakeSymbol("list")
reg19667 := Nil;
reg19668 := PrimCons(A, reg19667)
reg19669 := PrimCons(reg19666, reg19668)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19669, V2752, V2753, V2754)
return
}, 0)
reg19671 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19665)
Result := reg19671
_ = Result
reg19672 := __e.Call(__defun__shen_4unbindv, V2595, V2753)
_ = reg19672
reg19674 = Result
} else {
reg19673 := False;
reg19674 = reg19673
}
reg19675 = reg19674
}
reg19677 = reg19675
} else {
reg19676 := False;
reg19677 = reg19676
}
reg19679 = reg19677
} else {
reg19678 := False;
reg19679 = reg19678
}
reg19681 = reg19679
} else {
reg19680 := False;
reg19681 = reg19680
}
reg19683 = reg19681
} else {
reg19682 := False;
reg19683 = reg19682
}
reg19685 = reg19683
} else {
reg19684 := False;
reg19685 = reg19684
}
Case := reg19685
_ = Case
reg19686 := False;
reg19687 := PrimEqual(Case, reg19686)
var reg20870 Obj
if reg19687 == True {
reg19688 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2601 := reg19688
_ = V2601
reg19689 := PrimIsPair(V2601)
var reg19836 Obj
if reg19689 == True {
reg19690 := PrimHead(V2601)
reg19691 := __e.Call(__defun__shen_4lazyderef, reg19690, V2753)
V2602 := reg19691
_ = V2602
reg19692 := MakeSymbol("@p")
reg19693 := PrimEqual(reg19692, V2602)
var reg19834 Obj
if reg19693 == True {
reg19694 := PrimTail(V2601)
reg19695 := __e.Call(__defun__shen_4lazyderef, reg19694, V2753)
V2603 := reg19695
_ = V2603
reg19696 := PrimIsPair(V2603)
var reg19832 Obj
if reg19696 == True {
reg19697 := PrimHead(V2603)
X := reg19697
_ = X
reg19698 := PrimTail(V2603)
reg19699 := __e.Call(__defun__shen_4lazyderef, reg19698, V2753)
V2604 := reg19699
_ = V2604
reg19700 := PrimIsPair(V2604)
var reg19830 Obj
if reg19700 == True {
reg19701 := PrimHead(V2604)
Y := reg19701
_ = Y
reg19702 := PrimTail(V2604)
reg19703 := __e.Call(__defun__shen_4lazyderef, reg19702, V2753)
V2605 := reg19703
_ = V2605
reg19704 := Nil;
reg19705 := PrimEqual(reg19704, V2605)
var reg19828 Obj
if reg19705 == True {
reg19706 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2606 := reg19706
_ = V2606
reg19707 := PrimIsPair(V2606)
var reg19826 Obj
if reg19707 == True {
reg19708 := PrimHead(V2606)
A := reg19708
_ = A
reg19709 := PrimTail(V2606)
reg19710 := __e.Call(__defun__shen_4lazyderef, reg19709, V2753)
V2607 := reg19710
_ = V2607
reg19711 := PrimIsPair(V2607)
var reg19809 Obj
if reg19711 == True {
reg19712 := PrimHead(V2607)
reg19713 := __e.Call(__defun__shen_4lazyderef, reg19712, V2753)
V2608 := reg19713
_ = V2608
reg19714 := MakeSymbol("*")
reg19715 := PrimEqual(reg19714, V2608)
var reg19794 Obj
if reg19715 == True {
reg19716 := PrimTail(V2607)
reg19717 := __e.Call(__defun__shen_4lazyderef, reg19716, V2753)
V2609 := reg19717
_ = V2609
reg19718 := PrimIsPair(V2609)
var reg19751 Obj
if reg19718 == True {
reg19719 := PrimHead(V2609)
B := reg19719
_ = B
reg19720 := PrimTail(V2609)
reg19721 := __e.Call(__defun__shen_4lazyderef, reg19720, V2753)
V2610 := reg19721
_ = V2610
reg19722 := Nil;
reg19723 := PrimEqual(reg19722, V2610)
var reg19738 Obj
if reg19723 == True {
reg19724 := __e.Call(__defun__shen_4incinfs)
_ = reg19724
reg19725 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19727 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19725)
reg19738 = reg19727
} else {
reg19728 := __e.Call(__defun__shen_4pvar_2, V2610)
var reg19737 Obj
if reg19728 == True {
reg19729 := Nil;
reg19730 := __e.Call(__defun__shen_4bindv, V2610, reg19729, V2753)
_ = reg19730
reg19731 := __e.Call(__defun__shen_4incinfs)
_ = reg19731
reg19732 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19734 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19732)
Result := reg19734
_ = Result
reg19735 := __e.Call(__defun__shen_4unbindv, V2610, V2753)
_ = reg19735
reg19737 = Result
} else {
reg19736 := False;
reg19737 = reg19736
}
reg19738 = reg19737
}
reg19751 = reg19738
} else {
reg19739 := __e.Call(__defun__shen_4pvar_2, V2609)
var reg19750 Obj
if reg19739 == True {
reg19740 := __e.Call(__defun__shen_4newpv, V2753)
B := reg19740
_ = B
reg19741 := Nil;
reg19742 := PrimCons(B, reg19741)
reg19743 := __e.Call(__defun__shen_4bindv, V2609, reg19742, V2753)
_ = reg19743
reg19744 := __e.Call(__defun__shen_4incinfs)
_ = reg19744
reg19745 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19747 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19745)
Result := reg19747
_ = Result
reg19748 := __e.Call(__defun__shen_4unbindv, V2609, V2753)
_ = reg19748
reg19750 = Result
} else {
reg19749 := False;
reg19750 = reg19749
}
reg19751 = reg19750
}
reg19794 = reg19751
} else {
reg19752 := __e.Call(__defun__shen_4pvar_2, V2608)
var reg19793 Obj
if reg19752 == True {
reg19753 := MakeSymbol("*")
reg19754 := __e.Call(__defun__shen_4bindv, V2608, reg19753, V2753)
_ = reg19754
reg19755 := PrimTail(V2607)
reg19756 := __e.Call(__defun__shen_4lazyderef, reg19755, V2753)
V2611 := reg19756
_ = V2611
reg19757 := PrimIsPair(V2611)
var reg19790 Obj
if reg19757 == True {
reg19758 := PrimHead(V2611)
B := reg19758
_ = B
reg19759 := PrimTail(V2611)
reg19760 := __e.Call(__defun__shen_4lazyderef, reg19759, V2753)
V2612 := reg19760
_ = V2612
reg19761 := Nil;
reg19762 := PrimEqual(reg19761, V2612)
var reg19777 Obj
if reg19762 == True {
reg19763 := __e.Call(__defun__shen_4incinfs)
_ = reg19763
reg19764 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19766 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19764)
reg19777 = reg19766
} else {
reg19767 := __e.Call(__defun__shen_4pvar_2, V2612)
var reg19776 Obj
if reg19767 == True {
reg19768 := Nil;
reg19769 := __e.Call(__defun__shen_4bindv, V2612, reg19768, V2753)
_ = reg19769
reg19770 := __e.Call(__defun__shen_4incinfs)
_ = reg19770
reg19771 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19773 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19771)
Result := reg19773
_ = Result
reg19774 := __e.Call(__defun__shen_4unbindv, V2612, V2753)
_ = reg19774
reg19776 = Result
} else {
reg19775 := False;
reg19776 = reg19775
}
reg19777 = reg19776
}
reg19790 = reg19777
} else {
reg19778 := __e.Call(__defun__shen_4pvar_2, V2611)
var reg19789 Obj
if reg19778 == True {
reg19779 := __e.Call(__defun__shen_4newpv, V2753)
B := reg19779
_ = B
reg19780 := Nil;
reg19781 := PrimCons(B, reg19780)
reg19782 := __e.Call(__defun__shen_4bindv, V2611, reg19781, V2753)
_ = reg19782
reg19783 := __e.Call(__defun__shen_4incinfs)
_ = reg19783
reg19784 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19786 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19784)
Result := reg19786
_ = Result
reg19787 := __e.Call(__defun__shen_4unbindv, V2611, V2753)
_ = reg19787
reg19789 = Result
} else {
reg19788 := False;
reg19789 = reg19788
}
reg19790 = reg19789
}
Result := reg19790
_ = Result
reg19791 := __e.Call(__defun__shen_4unbindv, V2608, V2753)
_ = reg19791
reg19793 = Result
} else {
reg19792 := False;
reg19793 = reg19792
}
reg19794 = reg19793
}
reg19809 = reg19794
} else {
reg19795 := __e.Call(__defun__shen_4pvar_2, V2607)
var reg19808 Obj
if reg19795 == True {
reg19796 := __e.Call(__defun__shen_4newpv, V2753)
B := reg19796
_ = B
reg19797 := MakeSymbol("*")
reg19798 := Nil;
reg19799 := PrimCons(B, reg19798)
reg19800 := PrimCons(reg19797, reg19799)
reg19801 := __e.Call(__defun__shen_4bindv, V2607, reg19800, V2753)
_ = reg19801
reg19802 := __e.Call(__defun__shen_4incinfs)
_ = reg19802
reg19803 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19805 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19803)
Result := reg19805
_ = Result
reg19806 := __e.Call(__defun__shen_4unbindv, V2607, V2753)
_ = reg19806
reg19808 = Result
} else {
reg19807 := False;
reg19808 = reg19807
}
reg19809 = reg19808
}
reg19826 = reg19809
} else {
reg19810 := __e.Call(__defun__shen_4pvar_2, V2606)
var reg19825 Obj
if reg19810 == True {
reg19811 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19811
_ = A
reg19812 := __e.Call(__defun__shen_4newpv, V2753)
B := reg19812
_ = B
reg19813 := MakeSymbol("*")
reg19814 := Nil;
reg19815 := PrimCons(B, reg19814)
reg19816 := PrimCons(reg19813, reg19815)
reg19817 := PrimCons(A, reg19816)
reg19818 := __e.Call(__defun__shen_4bindv, V2606, reg19817, V2753)
_ = reg19818
reg19819 := __e.Call(__defun__shen_4incinfs)
_ = reg19819
reg19820 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V2752, V2753, V2754)
return
}, 0)
reg19822 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19820)
Result := reg19822
_ = Result
reg19823 := __e.Call(__defun__shen_4unbindv, V2606, V2753)
_ = reg19823
reg19825 = Result
} else {
reg19824 := False;
reg19825 = reg19824
}
reg19826 = reg19825
}
reg19828 = reg19826
} else {
reg19827 := False;
reg19828 = reg19827
}
reg19830 = reg19828
} else {
reg19829 := False;
reg19830 = reg19829
}
reg19832 = reg19830
} else {
reg19831 := False;
reg19832 = reg19831
}
reg19834 = reg19832
} else {
reg19833 := False;
reg19834 = reg19833
}
reg19836 = reg19834
} else {
reg19835 := False;
reg19836 = reg19835
}
Case := reg19836
_ = Case
reg19837 := False;
reg19838 := PrimEqual(Case, reg19837)
var reg20869 Obj
if reg19838 == True {
reg19839 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2613 := reg19839
_ = V2613
reg19840 := PrimIsPair(V2613)
var reg19994 Obj
if reg19840 == True {
reg19841 := PrimHead(V2613)
reg19842 := __e.Call(__defun__shen_4lazyderef, reg19841, V2753)
V2614 := reg19842
_ = V2614
reg19843 := MakeSymbol("@v")
reg19844 := PrimEqual(reg19843, V2614)
var reg19992 Obj
if reg19844 == True {
reg19845 := PrimTail(V2613)
reg19846 := __e.Call(__defun__shen_4lazyderef, reg19845, V2753)
V2615 := reg19846
_ = V2615
reg19847 := PrimIsPair(V2615)
var reg19990 Obj
if reg19847 == True {
reg19848 := PrimHead(V2615)
X := reg19848
_ = X
reg19849 := PrimTail(V2615)
reg19850 := __e.Call(__defun__shen_4lazyderef, reg19849, V2753)
V2616 := reg19850
_ = V2616
reg19851 := PrimIsPair(V2616)
var reg19988 Obj
if reg19851 == True {
reg19852 := PrimHead(V2616)
Y := reg19852
_ = Y
reg19853 := PrimTail(V2616)
reg19854 := __e.Call(__defun__shen_4lazyderef, reg19853, V2753)
V2617 := reg19854
_ = V2617
reg19855 := Nil;
reg19856 := PrimEqual(reg19855, V2617)
var reg19986 Obj
if reg19856 == True {
reg19857 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2618 := reg19857
_ = V2618
reg19858 := PrimIsPair(V2618)
var reg19984 Obj
if reg19858 == True {
reg19859 := PrimHead(V2618)
reg19860 := __e.Call(__defun__shen_4lazyderef, reg19859, V2753)
V2619 := reg19860
_ = V2619
reg19861 := MakeSymbol("vector")
reg19862 := PrimEqual(reg19861, V2619)
var reg19965 Obj
if reg19862 == True {
reg19863 := PrimTail(V2618)
reg19864 := __e.Call(__defun__shen_4lazyderef, reg19863, V2753)
V2620 := reg19864
_ = V2620
reg19865 := PrimIsPair(V2620)
var reg19910 Obj
if reg19865 == True {
reg19866 := PrimHead(V2620)
A := reg19866
_ = A
reg19867 := PrimTail(V2620)
reg19868 := __e.Call(__defun__shen_4lazyderef, reg19867, V2753)
V2621 := reg19868
_ = V2621
reg19869 := Nil;
reg19870 := PrimEqual(reg19869, V2621)
var reg19893 Obj
if reg19870 == True {
reg19871 := __e.Call(__defun__shen_4incinfs)
_ = reg19871
reg19872 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19873 := MakeSymbol("vector")
reg19874 := Nil;
reg19875 := PrimCons(A, reg19874)
reg19876 := PrimCons(reg19873, reg19875)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19876, V2752, V2753, V2754)
return
}, 0)
reg19878 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19872)
reg19893 = reg19878
} else {
reg19879 := __e.Call(__defun__shen_4pvar_2, V2621)
var reg19892 Obj
if reg19879 == True {
reg19880 := Nil;
reg19881 := __e.Call(__defun__shen_4bindv, V2621, reg19880, V2753)
_ = reg19881
reg19882 := __e.Call(__defun__shen_4incinfs)
_ = reg19882
reg19883 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19884 := MakeSymbol("vector")
reg19885 := Nil;
reg19886 := PrimCons(A, reg19885)
reg19887 := PrimCons(reg19884, reg19886)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19887, V2752, V2753, V2754)
return
}, 0)
reg19889 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19883)
Result := reg19889
_ = Result
reg19890 := __e.Call(__defun__shen_4unbindv, V2621, V2753)
_ = reg19890
reg19892 = Result
} else {
reg19891 := False;
reg19892 = reg19891
}
reg19893 = reg19892
}
reg19910 = reg19893
} else {
reg19894 := __e.Call(__defun__shen_4pvar_2, V2620)
var reg19909 Obj
if reg19894 == True {
reg19895 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19895
_ = A
reg19896 := Nil;
reg19897 := PrimCons(A, reg19896)
reg19898 := __e.Call(__defun__shen_4bindv, V2620, reg19897, V2753)
_ = reg19898
reg19899 := __e.Call(__defun__shen_4incinfs)
_ = reg19899
reg19900 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19901 := MakeSymbol("vector")
reg19902 := Nil;
reg19903 := PrimCons(A, reg19902)
reg19904 := PrimCons(reg19901, reg19903)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19904, V2752, V2753, V2754)
return
}, 0)
reg19906 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19900)
Result := reg19906
_ = Result
reg19907 := __e.Call(__defun__shen_4unbindv, V2620, V2753)
_ = reg19907
reg19909 = Result
} else {
reg19908 := False;
reg19909 = reg19908
}
reg19910 = reg19909
}
reg19965 = reg19910
} else {
reg19911 := __e.Call(__defun__shen_4pvar_2, V2619)
var reg19964 Obj
if reg19911 == True {
reg19912 := MakeSymbol("vector")
reg19913 := __e.Call(__defun__shen_4bindv, V2619, reg19912, V2753)
_ = reg19913
reg19914 := PrimTail(V2618)
reg19915 := __e.Call(__defun__shen_4lazyderef, reg19914, V2753)
V2622 := reg19915
_ = V2622
reg19916 := PrimIsPair(V2622)
var reg19961 Obj
if reg19916 == True {
reg19917 := PrimHead(V2622)
A := reg19917
_ = A
reg19918 := PrimTail(V2622)
reg19919 := __e.Call(__defun__shen_4lazyderef, reg19918, V2753)
V2623 := reg19919
_ = V2623
reg19920 := Nil;
reg19921 := PrimEqual(reg19920, V2623)
var reg19944 Obj
if reg19921 == True {
reg19922 := __e.Call(__defun__shen_4incinfs)
_ = reg19922
reg19923 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19924 := MakeSymbol("vector")
reg19925 := Nil;
reg19926 := PrimCons(A, reg19925)
reg19927 := PrimCons(reg19924, reg19926)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19927, V2752, V2753, V2754)
return
}, 0)
reg19929 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19923)
reg19944 = reg19929
} else {
reg19930 := __e.Call(__defun__shen_4pvar_2, V2623)
var reg19943 Obj
if reg19930 == True {
reg19931 := Nil;
reg19932 := __e.Call(__defun__shen_4bindv, V2623, reg19931, V2753)
_ = reg19932
reg19933 := __e.Call(__defun__shen_4incinfs)
_ = reg19933
reg19934 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19935 := MakeSymbol("vector")
reg19936 := Nil;
reg19937 := PrimCons(A, reg19936)
reg19938 := PrimCons(reg19935, reg19937)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19938, V2752, V2753, V2754)
return
}, 0)
reg19940 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19934)
Result := reg19940
_ = Result
reg19941 := __e.Call(__defun__shen_4unbindv, V2623, V2753)
_ = reg19941
reg19943 = Result
} else {
reg19942 := False;
reg19943 = reg19942
}
reg19944 = reg19943
}
reg19961 = reg19944
} else {
reg19945 := __e.Call(__defun__shen_4pvar_2, V2622)
var reg19960 Obj
if reg19945 == True {
reg19946 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19946
_ = A
reg19947 := Nil;
reg19948 := PrimCons(A, reg19947)
reg19949 := __e.Call(__defun__shen_4bindv, V2622, reg19948, V2753)
_ = reg19949
reg19950 := __e.Call(__defun__shen_4incinfs)
_ = reg19950
reg19951 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19952 := MakeSymbol("vector")
reg19953 := Nil;
reg19954 := PrimCons(A, reg19953)
reg19955 := PrimCons(reg19952, reg19954)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19955, V2752, V2753, V2754)
return
}, 0)
reg19957 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19951)
Result := reg19957
_ = Result
reg19958 := __e.Call(__defun__shen_4unbindv, V2622, V2753)
_ = reg19958
reg19960 = Result
} else {
reg19959 := False;
reg19960 = reg19959
}
reg19961 = reg19960
}
Result := reg19961
_ = Result
reg19962 := __e.Call(__defun__shen_4unbindv, V2619, V2753)
_ = reg19962
reg19964 = Result
} else {
reg19963 := False;
reg19964 = reg19963
}
reg19965 = reg19964
}
reg19984 = reg19965
} else {
reg19966 := __e.Call(__defun__shen_4pvar_2, V2618)
var reg19983 Obj
if reg19966 == True {
reg19967 := __e.Call(__defun__shen_4newpv, V2753)
A := reg19967
_ = A
reg19968 := MakeSymbol("vector")
reg19969 := Nil;
reg19970 := PrimCons(A, reg19969)
reg19971 := PrimCons(reg19968, reg19970)
reg19972 := __e.Call(__defun__shen_4bindv, V2618, reg19971, V2753)
_ = reg19972
reg19973 := __e.Call(__defun__shen_4incinfs)
_ = reg19973
reg19974 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19975 := MakeSymbol("vector")
reg19976 := Nil;
reg19977 := PrimCons(A, reg19976)
reg19978 := PrimCons(reg19975, reg19977)
__ctx.TailApply(__defun__shen_4th_d, Y, reg19978, V2752, V2753, V2754)
return
}, 0)
reg19980 := __e.Call(__defun__shen_4th_d, X, A, V2752, V2753, reg19974)
Result := reg19980
_ = Result
reg19981 := __e.Call(__defun__shen_4unbindv, V2618, V2753)
_ = reg19981
reg19983 = Result
} else {
reg19982 := False;
reg19983 = reg19982
}
reg19984 = reg19983
}
reg19986 = reg19984
} else {
reg19985 := False;
reg19986 = reg19985
}
reg19988 = reg19986
} else {
reg19987 := False;
reg19988 = reg19987
}
reg19990 = reg19988
} else {
reg19989 := False;
reg19990 = reg19989
}
reg19992 = reg19990
} else {
reg19991 := False;
reg19992 = reg19991
}
reg19994 = reg19992
} else {
reg19993 := False;
reg19994 = reg19993
}
Case := reg19994
_ = Case
reg19995 := False;
reg19996 := PrimEqual(Case, reg19995)
var reg20868 Obj
if reg19996 == True {
reg19997 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2624 := reg19997
_ = V2624
reg19998 := PrimIsPair(V2624)
var reg20046 Obj
if reg19998 == True {
reg19999 := PrimHead(V2624)
reg20000 := __e.Call(__defun__shen_4lazyderef, reg19999, V2753)
V2625 := reg20000
_ = V2625
reg20001 := MakeSymbol("@s")
reg20002 := PrimEqual(reg20001, V2625)
var reg20044 Obj
if reg20002 == True {
reg20003 := PrimTail(V2624)
reg20004 := __e.Call(__defun__shen_4lazyderef, reg20003, V2753)
V2626 := reg20004
_ = V2626
reg20005 := PrimIsPair(V2626)
var reg20042 Obj
if reg20005 == True {
reg20006 := PrimHead(V2626)
X := reg20006
_ = X
reg20007 := PrimTail(V2626)
reg20008 := __e.Call(__defun__shen_4lazyderef, reg20007, V2753)
V2627 := reg20008
_ = V2627
reg20009 := PrimIsPair(V2627)
var reg20040 Obj
if reg20009 == True {
reg20010 := PrimHead(V2627)
Y := reg20010
_ = Y
reg20011 := PrimTail(V2627)
reg20012 := __e.Call(__defun__shen_4lazyderef, reg20011, V2753)
V2628 := reg20012
_ = V2628
reg20013 := Nil;
reg20014 := PrimEqual(reg20013, V2628)
var reg20038 Obj
if reg20014 == True {
reg20015 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2629 := reg20015
_ = V2629
reg20016 := MakeSymbol("string")
reg20017 := PrimEqual(reg20016, V2629)
var reg20036 Obj
if reg20017 == True {
reg20018 := __e.Call(__defun__shen_4incinfs)
_ = reg20018
reg20019 := MakeSymbol("string")
reg20020 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20021 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, Y, reg20021, V2752, V2753, V2754)
return
}, 0)
reg20023 := __e.Call(__defun__shen_4th_d, X, reg20019, V2752, V2753, reg20020)
reg20036 = reg20023
} else {
reg20024 := __e.Call(__defun__shen_4pvar_2, V2629)
var reg20035 Obj
if reg20024 == True {
reg20025 := MakeSymbol("string")
reg20026 := __e.Call(__defun__shen_4bindv, V2629, reg20025, V2753)
_ = reg20026
reg20027 := __e.Call(__defun__shen_4incinfs)
_ = reg20027
reg20028 := MakeSymbol("string")
reg20029 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20030 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, Y, reg20030, V2752, V2753, V2754)
return
}, 0)
reg20032 := __e.Call(__defun__shen_4th_d, X, reg20028, V2752, V2753, reg20029)
Result := reg20032
_ = Result
reg20033 := __e.Call(__defun__shen_4unbindv, V2629, V2753)
_ = reg20033
reg20035 = Result
} else {
reg20034 := False;
reg20035 = reg20034
}
reg20036 = reg20035
}
reg20038 = reg20036
} else {
reg20037 := False;
reg20038 = reg20037
}
reg20040 = reg20038
} else {
reg20039 := False;
reg20040 = reg20039
}
reg20042 = reg20040
} else {
reg20041 := False;
reg20042 = reg20041
}
reg20044 = reg20042
} else {
reg20043 := False;
reg20044 = reg20043
}
reg20046 = reg20044
} else {
reg20045 := False;
reg20046 = reg20045
}
Case := reg20046
_ = Case
reg20047 := False;
reg20048 := PrimEqual(Case, reg20047)
var reg20867 Obj
if reg20048 == True {
reg20049 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2630 := reg20049
_ = V2630
reg20050 := PrimIsPair(V2630)
var reg20333 Obj
if reg20050 == True {
reg20051 := PrimHead(V2630)
reg20052 := __e.Call(__defun__shen_4lazyderef, reg20051, V2753)
V2631 := reg20052
_ = V2631
reg20053 := MakeSymbol("lambda")
reg20054 := PrimEqual(reg20053, V2631)
var reg20331 Obj
if reg20054 == True {
reg20055 := PrimTail(V2630)
reg20056 := __e.Call(__defun__shen_4lazyderef, reg20055, V2753)
V2632 := reg20056
_ = V2632
reg20057 := PrimIsPair(V2632)
var reg20329 Obj
if reg20057 == True {
reg20058 := PrimHead(V2632)
X := reg20058
_ = X
reg20059 := PrimTail(V2632)
reg20060 := __e.Call(__defun__shen_4lazyderef, reg20059, V2753)
V2633 := reg20060
_ = V2633
reg20061 := PrimIsPair(V2633)
var reg20327 Obj
if reg20061 == True {
reg20062 := PrimHead(V2633)
Y := reg20062
_ = Y
reg20063 := PrimTail(V2633)
reg20064 := __e.Call(__defun__shen_4lazyderef, reg20063, V2753)
V2634 := reg20064
_ = V2634
reg20065 := Nil;
reg20066 := PrimEqual(reg20065, V2634)
var reg20325 Obj
if reg20066 == True {
reg20067 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2635 := reg20067
_ = V2635
reg20068 := PrimIsPair(V2635)
var reg20323 Obj
if reg20068 == True {
reg20069 := PrimHead(V2635)
A := reg20069
_ = A
reg20070 := PrimTail(V2635)
reg20071 := __e.Call(__defun__shen_4lazyderef, reg20070, V2753)
V2636 := reg20071
_ = V2636
reg20072 := PrimIsPair(V2636)
var reg20289 Obj
if reg20072 == True {
reg20073 := PrimHead(V2636)
reg20074 := __e.Call(__defun__shen_4lazyderef, reg20073, V2753)
V2637 := reg20074
_ = V2637
reg20075 := MakeSymbol("-->")
reg20076 := PrimEqual(reg20075, V2637)
var reg20257 Obj
if reg20076 == True {
reg20077 := PrimTail(V2636)
reg20078 := __e.Call(__defun__shen_4lazyderef, reg20077, V2753)
V2638 := reg20078
_ = V2638
reg20079 := PrimIsPair(V2638)
var reg20163 Obj
if reg20079 == True {
reg20080 := PrimHead(V2638)
B := reg20080
_ = B
reg20081 := PrimTail(V2638)
reg20082 := __e.Call(__defun__shen_4lazyderef, reg20081, V2753)
V2639 := reg20082
_ = V2639
reg20083 := Nil;
reg20084 := PrimEqual(reg20083, V2639)
var reg20133 Obj
if reg20084 == True {
reg20085 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20085
_ = Z
reg20086 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20086
_ = X_e_e
reg20087 := __e.Call(__defun__shen_4incinfs)
_ = reg20087
reg20088 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20089 := __e.Call(__defun__shen_4placeholder)
reg20090 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20091 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20092 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20093 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20094 := __e.Call(__defun__shen_4ebr, reg20091, reg20092, reg20093)
reg20095 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20096 := MakeSymbol(":")
reg20097 := Nil;
reg20098 := PrimCons(A, reg20097)
reg20099 := PrimCons(reg20096, reg20098)
reg20100 := PrimCons(X_e_e, reg20099)
reg20101 := PrimCons(reg20100, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20101, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20094, V2753, reg20095)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20089, V2753, reg20090)
return
}, 0)
reg20105 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20088)
reg20133 = reg20105
} else {
reg20106 := __e.Call(__defun__shen_4pvar_2, V2639)
var reg20132 Obj
if reg20106 == True {
reg20107 := Nil;
reg20108 := __e.Call(__defun__shen_4bindv, V2639, reg20107, V2753)
_ = reg20108
reg20109 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20109
_ = Z
reg20110 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20110
_ = X_e_e
reg20111 := __e.Call(__defun__shen_4incinfs)
_ = reg20111
reg20112 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20113 := __e.Call(__defun__shen_4placeholder)
reg20114 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20115 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20116 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20117 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20118 := __e.Call(__defun__shen_4ebr, reg20115, reg20116, reg20117)
reg20119 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20120 := MakeSymbol(":")
reg20121 := Nil;
reg20122 := PrimCons(A, reg20121)
reg20123 := PrimCons(reg20120, reg20122)
reg20124 := PrimCons(X_e_e, reg20123)
reg20125 := PrimCons(reg20124, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20125, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20118, V2753, reg20119)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20113, V2753, reg20114)
return
}, 0)
reg20129 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20112)
Result := reg20129
_ = Result
reg20130 := __e.Call(__defun__shen_4unbindv, V2639, V2753)
_ = reg20130
reg20132 = Result
} else {
reg20131 := False;
reg20132 = reg20131
}
reg20133 = reg20132
}
reg20163 = reg20133
} else {
reg20134 := __e.Call(__defun__shen_4pvar_2, V2638)
var reg20162 Obj
if reg20134 == True {
reg20135 := __e.Call(__defun__shen_4newpv, V2753)
B := reg20135
_ = B
reg20136 := Nil;
reg20137 := PrimCons(B, reg20136)
reg20138 := __e.Call(__defun__shen_4bindv, V2638, reg20137, V2753)
_ = reg20138
reg20139 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20139
_ = Z
reg20140 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20140
_ = X_e_e
reg20141 := __e.Call(__defun__shen_4incinfs)
_ = reg20141
reg20142 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20143 := __e.Call(__defun__shen_4placeholder)
reg20144 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20145 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20146 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20147 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20148 := __e.Call(__defun__shen_4ebr, reg20145, reg20146, reg20147)
reg20149 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20150 := MakeSymbol(":")
reg20151 := Nil;
reg20152 := PrimCons(A, reg20151)
reg20153 := PrimCons(reg20150, reg20152)
reg20154 := PrimCons(X_e_e, reg20153)
reg20155 := PrimCons(reg20154, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20155, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20148, V2753, reg20149)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20143, V2753, reg20144)
return
}, 0)
reg20159 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20142)
Result := reg20159
_ = Result
reg20160 := __e.Call(__defun__shen_4unbindv, V2638, V2753)
_ = reg20160
reg20162 = Result
} else {
reg20161 := False;
reg20162 = reg20161
}
reg20163 = reg20162
}
reg20257 = reg20163
} else {
reg20164 := __e.Call(__defun__shen_4pvar_2, V2637)
var reg20256 Obj
if reg20164 == True {
reg20165 := MakeSymbol("-->")
reg20166 := __e.Call(__defun__shen_4bindv, V2637, reg20165, V2753)
_ = reg20166
reg20167 := PrimTail(V2636)
reg20168 := __e.Call(__defun__shen_4lazyderef, reg20167, V2753)
V2640 := reg20168
_ = V2640
reg20169 := PrimIsPair(V2640)
var reg20253 Obj
if reg20169 == True {
reg20170 := PrimHead(V2640)
B := reg20170
_ = B
reg20171 := PrimTail(V2640)
reg20172 := __e.Call(__defun__shen_4lazyderef, reg20171, V2753)
V2641 := reg20172
_ = V2641
reg20173 := Nil;
reg20174 := PrimEqual(reg20173, V2641)
var reg20223 Obj
if reg20174 == True {
reg20175 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20175
_ = Z
reg20176 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20176
_ = X_e_e
reg20177 := __e.Call(__defun__shen_4incinfs)
_ = reg20177
reg20178 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20179 := __e.Call(__defun__shen_4placeholder)
reg20180 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20181 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20182 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20183 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20184 := __e.Call(__defun__shen_4ebr, reg20181, reg20182, reg20183)
reg20185 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20186 := MakeSymbol(":")
reg20187 := Nil;
reg20188 := PrimCons(A, reg20187)
reg20189 := PrimCons(reg20186, reg20188)
reg20190 := PrimCons(X_e_e, reg20189)
reg20191 := PrimCons(reg20190, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20191, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20184, V2753, reg20185)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20179, V2753, reg20180)
return
}, 0)
reg20195 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20178)
reg20223 = reg20195
} else {
reg20196 := __e.Call(__defun__shen_4pvar_2, V2641)
var reg20222 Obj
if reg20196 == True {
reg20197 := Nil;
reg20198 := __e.Call(__defun__shen_4bindv, V2641, reg20197, V2753)
_ = reg20198
reg20199 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20199
_ = Z
reg20200 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20200
_ = X_e_e
reg20201 := __e.Call(__defun__shen_4incinfs)
_ = reg20201
reg20202 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20203 := __e.Call(__defun__shen_4placeholder)
reg20204 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20205 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20206 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20207 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20208 := __e.Call(__defun__shen_4ebr, reg20205, reg20206, reg20207)
reg20209 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20210 := MakeSymbol(":")
reg20211 := Nil;
reg20212 := PrimCons(A, reg20211)
reg20213 := PrimCons(reg20210, reg20212)
reg20214 := PrimCons(X_e_e, reg20213)
reg20215 := PrimCons(reg20214, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20215, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20208, V2753, reg20209)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20203, V2753, reg20204)
return
}, 0)
reg20219 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20202)
Result := reg20219
_ = Result
reg20220 := __e.Call(__defun__shen_4unbindv, V2641, V2753)
_ = reg20220
reg20222 = Result
} else {
reg20221 := False;
reg20222 = reg20221
}
reg20223 = reg20222
}
reg20253 = reg20223
} else {
reg20224 := __e.Call(__defun__shen_4pvar_2, V2640)
var reg20252 Obj
if reg20224 == True {
reg20225 := __e.Call(__defun__shen_4newpv, V2753)
B := reg20225
_ = B
reg20226 := Nil;
reg20227 := PrimCons(B, reg20226)
reg20228 := __e.Call(__defun__shen_4bindv, V2640, reg20227, V2753)
_ = reg20228
reg20229 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20229
_ = Z
reg20230 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20230
_ = X_e_e
reg20231 := __e.Call(__defun__shen_4incinfs)
_ = reg20231
reg20232 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20233 := __e.Call(__defun__shen_4placeholder)
reg20234 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20235 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20236 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20237 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20238 := __e.Call(__defun__shen_4ebr, reg20235, reg20236, reg20237)
reg20239 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20240 := MakeSymbol(":")
reg20241 := Nil;
reg20242 := PrimCons(A, reg20241)
reg20243 := PrimCons(reg20240, reg20242)
reg20244 := PrimCons(X_e_e, reg20243)
reg20245 := PrimCons(reg20244, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20245, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20238, V2753, reg20239)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20233, V2753, reg20234)
return
}, 0)
reg20249 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20232)
Result := reg20249
_ = Result
reg20250 := __e.Call(__defun__shen_4unbindv, V2640, V2753)
_ = reg20250
reg20252 = Result
} else {
reg20251 := False;
reg20252 = reg20251
}
reg20253 = reg20252
}
Result := reg20253
_ = Result
reg20254 := __e.Call(__defun__shen_4unbindv, V2637, V2753)
_ = reg20254
reg20256 = Result
} else {
reg20255 := False;
reg20256 = reg20255
}
reg20257 = reg20256
}
reg20289 = reg20257
} else {
reg20258 := __e.Call(__defun__shen_4pvar_2, V2636)
var reg20288 Obj
if reg20258 == True {
reg20259 := __e.Call(__defun__shen_4newpv, V2753)
B := reg20259
_ = B
reg20260 := MakeSymbol("-->")
reg20261 := Nil;
reg20262 := PrimCons(B, reg20261)
reg20263 := PrimCons(reg20260, reg20262)
reg20264 := __e.Call(__defun__shen_4bindv, V2636, reg20263, V2753)
_ = reg20264
reg20265 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20265
_ = Z
reg20266 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20266
_ = X_e_e
reg20267 := __e.Call(__defun__shen_4incinfs)
_ = reg20267
reg20268 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20269 := __e.Call(__defun__shen_4placeholder)
reg20270 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20271 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20272 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20273 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20274 := __e.Call(__defun__shen_4ebr, reg20271, reg20272, reg20273)
reg20275 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20276 := MakeSymbol(":")
reg20277 := Nil;
reg20278 := PrimCons(A, reg20277)
reg20279 := PrimCons(reg20276, reg20278)
reg20280 := PrimCons(X_e_e, reg20279)
reg20281 := PrimCons(reg20280, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20281, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20274, V2753, reg20275)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20269, V2753, reg20270)
return
}, 0)
reg20285 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20268)
Result := reg20285
_ = Result
reg20286 := __e.Call(__defun__shen_4unbindv, V2636, V2753)
_ = reg20286
reg20288 = Result
} else {
reg20287 := False;
reg20288 = reg20287
}
reg20289 = reg20288
}
reg20323 = reg20289
} else {
reg20290 := __e.Call(__defun__shen_4pvar_2, V2635)
var reg20322 Obj
if reg20290 == True {
reg20291 := __e.Call(__defun__shen_4newpv, V2753)
A := reg20291
_ = A
reg20292 := __e.Call(__defun__shen_4newpv, V2753)
B := reg20292
_ = B
reg20293 := MakeSymbol("-->")
reg20294 := Nil;
reg20295 := PrimCons(B, reg20294)
reg20296 := PrimCons(reg20293, reg20295)
reg20297 := PrimCons(A, reg20296)
reg20298 := __e.Call(__defun__shen_4bindv, V2635, reg20297, V2753)
_ = reg20298
reg20299 := __e.Call(__defun__shen_4newpv, V2753)
Z := reg20299
_ = Z
reg20300 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20300
_ = X_e_e
reg20301 := __e.Call(__defun__shen_4incinfs)
_ = reg20301
reg20302 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20303 := __e.Call(__defun__shen_4placeholder)
reg20304 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20305 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20306 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20307 := __e.Call(__defun__shen_4lazyderef, Y, V2753)
reg20308 := __e.Call(__defun__shen_4ebr, reg20305, reg20306, reg20307)
reg20309 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20310 := MakeSymbol(":")
reg20311 := Nil;
reg20312 := PrimCons(A, reg20311)
reg20313 := PrimCons(reg20310, reg20312)
reg20314 := PrimCons(X_e_e, reg20313)
reg20315 := PrimCons(reg20314, V2752)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg20315, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg20308, V2753, reg20309)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20303, V2753, reg20304)
return
}, 0)
reg20319 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20302)
Result := reg20319
_ = Result
reg20320 := __e.Call(__defun__shen_4unbindv, V2635, V2753)
_ = reg20320
reg20322 = Result
} else {
reg20321 := False;
reg20322 = reg20321
}
reg20323 = reg20322
}
reg20325 = reg20323
} else {
reg20324 := False;
reg20325 = reg20324
}
reg20327 = reg20325
} else {
reg20326 := False;
reg20327 = reg20326
}
reg20329 = reg20327
} else {
reg20328 := False;
reg20329 = reg20328
}
reg20331 = reg20329
} else {
reg20330 := False;
reg20331 = reg20330
}
reg20333 = reg20331
} else {
reg20332 := False;
reg20333 = reg20332
}
Case := reg20333
_ = Case
reg20334 := False;
reg20335 := PrimEqual(Case, reg20334)
var reg20866 Obj
if reg20335 == True {
reg20336 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2642 := reg20336
_ = V2642
reg20337 := PrimIsPair(V2642)
var reg20391 Obj
if reg20337 == True {
reg20338 := PrimHead(V2642)
reg20339 := __e.Call(__defun__shen_4lazyderef, reg20338, V2753)
V2643 := reg20339
_ = V2643
reg20340 := MakeSymbol("let")
reg20341 := PrimEqual(reg20340, V2643)
var reg20389 Obj
if reg20341 == True {
reg20342 := PrimTail(V2642)
reg20343 := __e.Call(__defun__shen_4lazyderef, reg20342, V2753)
V2644 := reg20343
_ = V2644
reg20344 := PrimIsPair(V2644)
var reg20387 Obj
if reg20344 == True {
reg20345 := PrimHead(V2644)
X := reg20345
_ = X
reg20346 := PrimTail(V2644)
reg20347 := __e.Call(__defun__shen_4lazyderef, reg20346, V2753)
V2645 := reg20347
_ = V2645
reg20348 := PrimIsPair(V2645)
var reg20385 Obj
if reg20348 == True {
reg20349 := PrimHead(V2645)
Y := reg20349
_ = Y
reg20350 := PrimTail(V2645)
reg20351 := __e.Call(__defun__shen_4lazyderef, reg20350, V2753)
V2646 := reg20351
_ = V2646
reg20352 := PrimIsPair(V2646)
var reg20383 Obj
if reg20352 == True {
reg20353 := PrimHead(V2646)
Z := reg20353
_ = Z
reg20354 := PrimTail(V2646)
reg20355 := __e.Call(__defun__shen_4lazyderef, reg20354, V2753)
V2647 := reg20355
_ = V2647
reg20356 := Nil;
reg20357 := PrimEqual(reg20356, V2647)
var reg20381 Obj
if reg20357 == True {
reg20358 := __e.Call(__defun__shen_4newpv, V2753)
W := reg20358
_ = W
reg20359 := __e.Call(__defun__shen_4newpv, V2753)
X_e_e := reg20359
_ = X_e_e
reg20360 := __e.Call(__defun__shen_4newpv, V2753)
B := reg20360
_ = B
reg20361 := __e.Call(__defun__shen_4incinfs)
_ = reg20361
reg20362 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20363 := __e.Call(__defun__shen_4placeholder)
reg20364 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20365 := __e.Call(__defun__shen_4lazyderef, X_e_e, V2753)
reg20366 := __e.Call(__defun__shen_4lazyderef, X, V2753)
reg20367 := __e.Call(__defun__shen_4lazyderef, Z, V2753)
reg20368 := __e.Call(__defun__shen_4ebr, reg20365, reg20366, reg20367)
reg20369 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20370 := MakeSymbol(":")
reg20371 := Nil;
reg20372 := PrimCons(B, reg20371)
reg20373 := PrimCons(reg20370, reg20372)
reg20374 := PrimCons(X_e_e, reg20373)
reg20375 := PrimCons(reg20374, V2752)
__ctx.TailApply(__defun__shen_4th_d, W, V2751, reg20375, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__bind, W, reg20368, V2753, reg20369)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg20363, V2753, reg20364)
return
}, 0)
reg20379 := __e.Call(__defun__shen_4th_d, Y, B, V2752, V2753, reg20362)
reg20381 = reg20379
} else {
reg20380 := False;
reg20381 = reg20380
}
reg20383 = reg20381
} else {
reg20382 := False;
reg20383 = reg20382
}
reg20385 = reg20383
} else {
reg20384 := False;
reg20385 = reg20384
}
reg20387 = reg20385
} else {
reg20386 := False;
reg20387 = reg20386
}
reg20389 = reg20387
} else {
reg20388 := False;
reg20389 = reg20388
}
reg20391 = reg20389
} else {
reg20390 := False;
reg20391 = reg20390
}
Case := reg20391
_ = Case
reg20392 := False;
reg20393 := PrimEqual(Case, reg20392)
var reg20865 Obj
if reg20393 == True {
reg20394 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2648 := reg20394
_ = V2648
reg20395 := PrimIsPair(V2648)
var reg20605 Obj
if reg20395 == True {
reg20396 := PrimHead(V2648)
reg20397 := __e.Call(__defun__shen_4lazyderef, reg20396, V2753)
V2649 := reg20397
_ = V2649
reg20398 := MakeSymbol("open")
reg20399 := PrimEqual(reg20398, V2649)
var reg20603 Obj
if reg20399 == True {
reg20400 := PrimTail(V2648)
reg20401 := __e.Call(__defun__shen_4lazyderef, reg20400, V2753)
V2650 := reg20401
_ = V2650
reg20402 := PrimIsPair(V2650)
var reg20601 Obj
if reg20402 == True {
reg20403 := PrimHead(V2650)
FileName := reg20403
_ = FileName
reg20404 := PrimTail(V2650)
reg20405 := __e.Call(__defun__shen_4lazyderef, reg20404, V2753)
V2651 := reg20405
_ = V2651
reg20406 := PrimIsPair(V2651)
var reg20599 Obj
if reg20406 == True {
reg20407 := PrimHead(V2651)
Direction2581 := reg20407
_ = Direction2581
reg20408 := PrimTail(V2651)
reg20409 := __e.Call(__defun__shen_4lazyderef, reg20408, V2753)
V2652 := reg20409
_ = V2652
reg20410 := Nil;
reg20411 := PrimEqual(reg20410, V2652)
var reg20597 Obj
if reg20411 == True {
reg20412 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2653 := reg20412
_ = V2653
reg20413 := PrimIsPair(V2653)
var reg20595 Obj
if reg20413 == True {
reg20414 := PrimHead(V2653)
reg20415 := __e.Call(__defun__shen_4lazyderef, reg20414, V2753)
V2654 := reg20415
_ = V2654
reg20416 := MakeSymbol("stream")
reg20417 := PrimEqual(reg20416, V2654)
var reg20568 Obj
if reg20417 == True {
reg20418 := PrimTail(V2653)
reg20419 := __e.Call(__defun__shen_4lazyderef, reg20418, V2753)
V2655 := reg20419
_ = V2655
reg20420 := PrimIsPair(V2655)
var reg20489 Obj
if reg20420 == True {
reg20421 := PrimHead(V2655)
Direction := reg20421
_ = Direction
reg20422 := PrimTail(V2655)
reg20423 := __e.Call(__defun__shen_4lazyderef, reg20422, V2753)
V2656 := reg20423
_ = V2656
reg20424 := Nil;
reg20425 := PrimEqual(reg20424, V2656)
var reg20464 Obj
if reg20425 == True {
reg20426 := __e.Call(__defun__shen_4incinfs)
_ = reg20426
reg20427 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20428 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20429 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20430 := MakeSymbol("in")
reg20431 := MakeSymbol("out")
reg20432 := Nil;
reg20433 := PrimCons(reg20431, reg20432)
reg20434 := PrimCons(reg20430, reg20433)
reg20435 := __e.Call(__defun__element_2, reg20429, reg20434)
reg20436 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20437 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20437, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20435, V2753, reg20436)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20428)
return
}, 0)
reg20441 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20427)
reg20464 = reg20441
} else {
reg20442 := __e.Call(__defun__shen_4pvar_2, V2656)
var reg20463 Obj
if reg20442 == True {
reg20443 := Nil;
reg20444 := __e.Call(__defun__shen_4bindv, V2656, reg20443, V2753)
_ = reg20444
reg20445 := __e.Call(__defun__shen_4incinfs)
_ = reg20445
reg20446 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20447 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20448 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20449 := MakeSymbol("in")
reg20450 := MakeSymbol("out")
reg20451 := Nil;
reg20452 := PrimCons(reg20450, reg20451)
reg20453 := PrimCons(reg20449, reg20452)
reg20454 := __e.Call(__defun__element_2, reg20448, reg20453)
reg20455 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20456 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20456, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20454, V2753, reg20455)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20447)
return
}, 0)
reg20460 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20446)
Result := reg20460
_ = Result
reg20461 := __e.Call(__defun__shen_4unbindv, V2656, V2753)
_ = reg20461
reg20463 = Result
} else {
reg20462 := False;
reg20463 = reg20462
}
reg20464 = reg20463
}
reg20489 = reg20464
} else {
reg20465 := __e.Call(__defun__shen_4pvar_2, V2655)
var reg20488 Obj
if reg20465 == True {
reg20466 := __e.Call(__defun__shen_4newpv, V2753)
Direction := reg20466
_ = Direction
reg20467 := Nil;
reg20468 := PrimCons(Direction, reg20467)
reg20469 := __e.Call(__defun__shen_4bindv, V2655, reg20468, V2753)
_ = reg20469
reg20470 := __e.Call(__defun__shen_4incinfs)
_ = reg20470
reg20471 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20472 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20473 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20474 := MakeSymbol("in")
reg20475 := MakeSymbol("out")
reg20476 := Nil;
reg20477 := PrimCons(reg20475, reg20476)
reg20478 := PrimCons(reg20474, reg20477)
reg20479 := __e.Call(__defun__element_2, reg20473, reg20478)
reg20480 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20481 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20481, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20479, V2753, reg20480)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20472)
return
}, 0)
reg20485 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20471)
Result := reg20485
_ = Result
reg20486 := __e.Call(__defun__shen_4unbindv, V2655, V2753)
_ = reg20486
reg20488 = Result
} else {
reg20487 := False;
reg20488 = reg20487
}
reg20489 = reg20488
}
reg20568 = reg20489
} else {
reg20490 := __e.Call(__defun__shen_4pvar_2, V2654)
var reg20567 Obj
if reg20490 == True {
reg20491 := MakeSymbol("stream")
reg20492 := __e.Call(__defun__shen_4bindv, V2654, reg20491, V2753)
_ = reg20492
reg20493 := PrimTail(V2653)
reg20494 := __e.Call(__defun__shen_4lazyderef, reg20493, V2753)
V2657 := reg20494
_ = V2657
reg20495 := PrimIsPair(V2657)
var reg20564 Obj
if reg20495 == True {
reg20496 := PrimHead(V2657)
Direction := reg20496
_ = Direction
reg20497 := PrimTail(V2657)
reg20498 := __e.Call(__defun__shen_4lazyderef, reg20497, V2753)
V2658 := reg20498
_ = V2658
reg20499 := Nil;
reg20500 := PrimEqual(reg20499, V2658)
var reg20539 Obj
if reg20500 == True {
reg20501 := __e.Call(__defun__shen_4incinfs)
_ = reg20501
reg20502 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20503 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20504 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20505 := MakeSymbol("in")
reg20506 := MakeSymbol("out")
reg20507 := Nil;
reg20508 := PrimCons(reg20506, reg20507)
reg20509 := PrimCons(reg20505, reg20508)
reg20510 := __e.Call(__defun__element_2, reg20504, reg20509)
reg20511 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20512 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20512, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20510, V2753, reg20511)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20503)
return
}, 0)
reg20516 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20502)
reg20539 = reg20516
} else {
reg20517 := __e.Call(__defun__shen_4pvar_2, V2658)
var reg20538 Obj
if reg20517 == True {
reg20518 := Nil;
reg20519 := __e.Call(__defun__shen_4bindv, V2658, reg20518, V2753)
_ = reg20519
reg20520 := __e.Call(__defun__shen_4incinfs)
_ = reg20520
reg20521 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20522 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20523 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20524 := MakeSymbol("in")
reg20525 := MakeSymbol("out")
reg20526 := Nil;
reg20527 := PrimCons(reg20525, reg20526)
reg20528 := PrimCons(reg20524, reg20527)
reg20529 := __e.Call(__defun__element_2, reg20523, reg20528)
reg20530 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20531 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20531, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20529, V2753, reg20530)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20522)
return
}, 0)
reg20535 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20521)
Result := reg20535
_ = Result
reg20536 := __e.Call(__defun__shen_4unbindv, V2658, V2753)
_ = reg20536
reg20538 = Result
} else {
reg20537 := False;
reg20538 = reg20537
}
reg20539 = reg20538
}
reg20564 = reg20539
} else {
reg20540 := __e.Call(__defun__shen_4pvar_2, V2657)
var reg20563 Obj
if reg20540 == True {
reg20541 := __e.Call(__defun__shen_4newpv, V2753)
Direction := reg20541
_ = Direction
reg20542 := Nil;
reg20543 := PrimCons(Direction, reg20542)
reg20544 := __e.Call(__defun__shen_4bindv, V2657, reg20543, V2753)
_ = reg20544
reg20545 := __e.Call(__defun__shen_4incinfs)
_ = reg20545
reg20546 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20547 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20548 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20549 := MakeSymbol("in")
reg20550 := MakeSymbol("out")
reg20551 := Nil;
reg20552 := PrimCons(reg20550, reg20551)
reg20553 := PrimCons(reg20549, reg20552)
reg20554 := __e.Call(__defun__element_2, reg20548, reg20553)
reg20555 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20556 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20556, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20554, V2753, reg20555)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20547)
return
}, 0)
reg20560 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20546)
Result := reg20560
_ = Result
reg20561 := __e.Call(__defun__shen_4unbindv, V2657, V2753)
_ = reg20561
reg20563 = Result
} else {
reg20562 := False;
reg20563 = reg20562
}
reg20564 = reg20563
}
Result := reg20564
_ = Result
reg20565 := __e.Call(__defun__shen_4unbindv, V2654, V2753)
_ = reg20565
reg20567 = Result
} else {
reg20566 := False;
reg20567 = reg20566
}
reg20568 = reg20567
}
reg20595 = reg20568
} else {
reg20569 := __e.Call(__defun__shen_4pvar_2, V2653)
var reg20594 Obj
if reg20569 == True {
reg20570 := __e.Call(__defun__shen_4newpv, V2753)
Direction := reg20570
_ = Direction
reg20571 := MakeSymbol("stream")
reg20572 := Nil;
reg20573 := PrimCons(Direction, reg20572)
reg20574 := PrimCons(reg20571, reg20573)
reg20575 := __e.Call(__defun__shen_4bindv, V2653, reg20574, V2753)
_ = reg20575
reg20576 := __e.Call(__defun__shen_4incinfs)
_ = reg20576
reg20577 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20578 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20579 := __e.Call(__defun__shen_4lazyderef, Direction, V2753)
reg20580 := MakeSymbol("in")
reg20581 := MakeSymbol("out")
reg20582 := Nil;
reg20583 := PrimCons(reg20581, reg20582)
reg20584 := PrimCons(reg20580, reg20583)
reg20585 := __e.Call(__defun__element_2, reg20579, reg20584)
reg20586 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20587 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg20587, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg20585, V2753, reg20586)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20578)
return
}, 0)
reg20591 := __e.Call(__defun__unify_b, Direction, Direction2581, V2753, reg20577)
Result := reg20591
_ = Result
reg20592 := __e.Call(__defun__shen_4unbindv, V2653, V2753)
_ = reg20592
reg20594 = Result
} else {
reg20593 := False;
reg20594 = reg20593
}
reg20595 = reg20594
}
reg20597 = reg20595
} else {
reg20596 := False;
reg20597 = reg20596
}
reg20599 = reg20597
} else {
reg20598 := False;
reg20599 = reg20598
}
reg20601 = reg20599
} else {
reg20600 := False;
reg20601 = reg20600
}
reg20603 = reg20601
} else {
reg20602 := False;
reg20603 = reg20602
}
reg20605 = reg20603
} else {
reg20604 := False;
reg20605 = reg20604
}
Case := reg20605
_ = Case
reg20606 := False;
reg20607 := PrimEqual(Case, reg20606)
var reg20864 Obj
if reg20607 == True {
reg20608 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2659 := reg20608
_ = V2659
reg20609 := PrimIsPair(V2659)
var reg20641 Obj
if reg20609 == True {
reg20610 := PrimHead(V2659)
reg20611 := __e.Call(__defun__shen_4lazyderef, reg20610, V2753)
V2660 := reg20611
_ = V2660
reg20612 := MakeSymbol("type")
reg20613 := PrimEqual(reg20612, V2660)
var reg20639 Obj
if reg20613 == True {
reg20614 := PrimTail(V2659)
reg20615 := __e.Call(__defun__shen_4lazyderef, reg20614, V2753)
V2661 := reg20615
_ = V2661
reg20616 := PrimIsPair(V2661)
var reg20637 Obj
if reg20616 == True {
reg20617 := PrimHead(V2661)
X := reg20617
_ = X
reg20618 := PrimTail(V2661)
reg20619 := __e.Call(__defun__shen_4lazyderef, reg20618, V2753)
V2662 := reg20619
_ = V2662
reg20620 := PrimIsPair(V2662)
var reg20635 Obj
if reg20620 == True {
reg20621 := PrimHead(V2662)
A := reg20621
_ = A
reg20622 := PrimTail(V2662)
reg20623 := __e.Call(__defun__shen_4lazyderef, reg20622, V2753)
V2663 := reg20623
_ = V2663
reg20624 := Nil;
reg20625 := PrimEqual(reg20624, V2663)
var reg20633 Obj
if reg20625 == True {
reg20626 := __e.Call(__defun__shen_4incinfs)
_ = reg20626
reg20627 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20628 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, X, A, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__unify, A, V2751, V2753, reg20628)
return
}, 0)
reg20631 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20627)
reg20633 = reg20631
} else {
reg20632 := False;
reg20633 = reg20632
}
reg20635 = reg20633
} else {
reg20634 := False;
reg20635 = reg20634
}
reg20637 = reg20635
} else {
reg20636 := False;
reg20637 = reg20636
}
reg20639 = reg20637
} else {
reg20638 := False;
reg20639 = reg20638
}
reg20641 = reg20639
} else {
reg20640 := False;
reg20641 = reg20640
}
Case := reg20641
_ = Case
reg20642 := False;
reg20643 := PrimEqual(Case, reg20642)
var reg20863 Obj
if reg20643 == True {
reg20644 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2664 := reg20644
_ = V2664
reg20645 := PrimIsPair(V2664)
var reg20685 Obj
if reg20645 == True {
reg20646 := PrimHead(V2664)
reg20647 := __e.Call(__defun__shen_4lazyderef, reg20646, V2753)
V2665 := reg20647
_ = V2665
reg20648 := MakeSymbol("input+")
reg20649 := PrimEqual(reg20648, V2665)
var reg20683 Obj
if reg20649 == True {
reg20650 := PrimTail(V2664)
reg20651 := __e.Call(__defun__shen_4lazyderef, reg20650, V2753)
V2666 := reg20651
_ = V2666
reg20652 := PrimIsPair(V2666)
var reg20681 Obj
if reg20652 == True {
reg20653 := PrimHead(V2666)
A := reg20653
_ = A
reg20654 := PrimTail(V2666)
reg20655 := __e.Call(__defun__shen_4lazyderef, reg20654, V2753)
V2667 := reg20655
_ = V2667
reg20656 := PrimIsPair(V2667)
var reg20679 Obj
if reg20656 == True {
reg20657 := PrimHead(V2667)
Stream := reg20657
_ = Stream
reg20658 := PrimTail(V2667)
reg20659 := __e.Call(__defun__shen_4lazyderef, reg20658, V2753)
V2668 := reg20659
_ = V2668
reg20660 := Nil;
reg20661 := PrimEqual(reg20660, V2668)
var reg20677 Obj
if reg20661 == True {
reg20662 := __e.Call(__defun__shen_4newpv, V2753)
C := reg20662
_ = C
reg20663 := __e.Call(__defun__shen_4incinfs)
_ = reg20663
reg20664 := __e.Call(__defun__shen_4lazyderef, A, V2753)
reg20665 := __e.Call(__defun__shen_4demodulate, reg20664)
reg20666 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20667 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20668 := MakeSymbol("stream")
reg20669 := MakeSymbol("in")
reg20670 := Nil;
reg20671 := PrimCons(reg20669, reg20670)
reg20672 := PrimCons(reg20668, reg20671)
__ctx.TailApply(__defun__shen_4th_d, Stream, reg20672, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__unify, V2751, C, V2753, reg20667)
return
}, 0)
reg20675 := __e.Call(__defun__bind, C, reg20665, V2753, reg20666)
reg20677 = reg20675
} else {
reg20676 := False;
reg20677 = reg20676
}
reg20679 = reg20677
} else {
reg20678 := False;
reg20679 = reg20678
}
reg20681 = reg20679
} else {
reg20680 := False;
reg20681 = reg20680
}
reg20683 = reg20681
} else {
reg20682 := False;
reg20683 = reg20682
}
reg20685 = reg20683
} else {
reg20684 := False;
reg20685 = reg20684
}
Case := reg20685
_ = Case
reg20686 := False;
reg20687 := PrimEqual(Case, reg20686)
var reg20862 Obj
if reg20687 == True {
reg20688 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2669 := reg20688
_ = V2669
reg20689 := PrimIsPair(V2669)
var reg20730 Obj
if reg20689 == True {
reg20690 := PrimHead(V2669)
reg20691 := __e.Call(__defun__shen_4lazyderef, reg20690, V2753)
V2670 := reg20691
_ = V2670
reg20692 := MakeSymbol("set")
reg20693 := PrimEqual(reg20692, V2670)
var reg20728 Obj
if reg20693 == True {
reg20694 := PrimTail(V2669)
reg20695 := __e.Call(__defun__shen_4lazyderef, reg20694, V2753)
V2671 := reg20695
_ = V2671
reg20696 := PrimIsPair(V2671)
var reg20726 Obj
if reg20696 == True {
reg20697 := PrimHead(V2671)
Var := reg20697
_ = Var
reg20698 := PrimTail(V2671)
reg20699 := __e.Call(__defun__shen_4lazyderef, reg20698, V2753)
V2672 := reg20699
_ = V2672
reg20700 := PrimIsPair(V2672)
var reg20724 Obj
if reg20700 == True {
reg20701 := PrimHead(V2672)
Val := reg20701
_ = Val
reg20702 := PrimTail(V2672)
reg20703 := __e.Call(__defun__shen_4lazyderef, reg20702, V2753)
V2673 := reg20703
_ = V2673
reg20704 := Nil;
reg20705 := PrimEqual(reg20704, V2673)
var reg20722 Obj
if reg20705 == True {
reg20706 := __e.Call(__defun__shen_4incinfs)
_ = reg20706
reg20707 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20708 := MakeSymbol("symbol")
reg20709 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20710 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20711 := MakeSymbol("value")
reg20712 := Nil;
reg20713 := PrimCons(Var, reg20712)
reg20714 := PrimCons(reg20711, reg20713)
reg20715 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Val, V2751, V2752, V2753, V2754)
return
}, 0)
__ctx.TailApply(__defun__shen_4th_d, reg20714, V2751, V2752, V2753, reg20715)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2753, reg20710)
return
}, 0)
__ctx.TailApply(__defun__shen_4th_d, Var, reg20708, V2752, V2753, reg20709)
return
}, 0)
reg20720 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20707)
reg20722 = reg20720
} else {
reg20721 := False;
reg20722 = reg20721
}
reg20724 = reg20722
} else {
reg20723 := False;
reg20724 = reg20723
}
reg20726 = reg20724
} else {
reg20725 := False;
reg20726 = reg20725
}
reg20728 = reg20726
} else {
reg20727 := False;
reg20728 = reg20727
}
reg20730 = reg20728
} else {
reg20729 := False;
reg20730 = reg20729
}
Case := reg20730
_ = Case
reg20731 := False;
reg20732 := PrimEqual(Case, reg20731)
var reg20861 Obj
if reg20732 == True {
reg20733 := __e.Call(__defun__shen_4newpv, V2753)
NewHyp := reg20733
_ = NewHyp
reg20734 := __e.Call(__defun__shen_4incinfs)
_ = reg20734
reg20735 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, V2750, V2751, NewHyp, V2753, V2754)
return
}, 0)
reg20737 := __e.Call(__defun__shen_4t_d_1hyps, V2752, NewHyp, V2753, reg20735)
Case := reg20737
_ = Case
reg20738 := False;
reg20739 := PrimEqual(Case, reg20738)
var reg20860 Obj
if reg20739 == True {
reg20740 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2674 := reg20740
_ = V2674
reg20741 := PrimIsPair(V2674)
var reg20763 Obj
if reg20741 == True {
reg20742 := PrimHead(V2674)
reg20743 := __e.Call(__defun__shen_4lazyderef, reg20742, V2753)
V2675 := reg20743
_ = V2675
reg20744 := MakeSymbol("define")
reg20745 := PrimEqual(reg20744, V2675)
var reg20761 Obj
if reg20745 == True {
reg20746 := PrimTail(V2674)
reg20747 := __e.Call(__defun__shen_4lazyderef, reg20746, V2753)
V2676 := reg20747
_ = V2676
reg20748 := PrimIsPair(V2676)
var reg20759 Obj
if reg20748 == True {
reg20749 := PrimHead(V2676)
F := reg20749
_ = F
reg20750 := PrimTail(V2676)
X := reg20750
_ = X
reg20751 := __e.Call(__defun__shen_4incinfs)
_ = reg20751
reg20752 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20753 := MakeSymbol("define")
reg20754 := PrimCons(F, X)
reg20755 := PrimCons(reg20753, reg20754)
__ctx.TailApply(__defun__shen_4t_d_1def, reg20755, V2751, V2752, V2753, V2754)
return
}, 0)
reg20757 := __e.Call(__defun__cut, Throwcontrol, V2753, reg20752)
reg20759 = reg20757
} else {
reg20758 := False;
reg20759 = reg20758
}
reg20761 = reg20759
} else {
reg20760 := False;
reg20761 = reg20760
}
reg20763 = reg20761
} else {
reg20762 := False;
reg20763 = reg20762
}
Case := reg20763
_ = Case
reg20764 := False;
reg20765 := PrimEqual(Case, reg20764)
var reg20859 Obj
if reg20765 == True {
reg20766 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2677 := reg20766
_ = V2677
reg20767 := PrimIsPair(V2677)
var reg20789 Obj
if reg20767 == True {
reg20768 := PrimHead(V2677)
reg20769 := __e.Call(__defun__shen_4lazyderef, reg20768, V2753)
V2678 := reg20769
_ = V2678
reg20770 := MakeSymbol("defmacro")
reg20771 := PrimEqual(reg20770, V2678)
var reg20787 Obj
if reg20771 == True {
reg20772 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2679 := reg20772
_ = V2679
reg20773 := MakeSymbol("unit")
reg20774 := PrimEqual(reg20773, V2679)
var reg20785 Obj
if reg20774 == True {
reg20775 := __e.Call(__defun__shen_4incinfs)
_ = reg20775
reg20776 := __e.Call(__defun__cut, Throwcontrol, V2753, V2754)
reg20785 = reg20776
} else {
reg20777 := __e.Call(__defun__shen_4pvar_2, V2679)
var reg20784 Obj
if reg20777 == True {
reg20778 := MakeSymbol("unit")
reg20779 := __e.Call(__defun__shen_4bindv, V2679, reg20778, V2753)
_ = reg20779
reg20780 := __e.Call(__defun__shen_4incinfs)
_ = reg20780
reg20781 := __e.Call(__defun__cut, Throwcontrol, V2753, V2754)
Result := reg20781
_ = Result
reg20782 := __e.Call(__defun__shen_4unbindv, V2679, V2753)
_ = reg20782
reg20784 = Result
} else {
reg20783 := False;
reg20784 = reg20783
}
reg20785 = reg20784
}
reg20787 = reg20785
} else {
reg20786 := False;
reg20787 = reg20786
}
reg20789 = reg20787
} else {
reg20788 := False;
reg20789 = reg20788
}
Case := reg20789
_ = Case
reg20790 := False;
reg20791 := PrimEqual(Case, reg20790)
var reg20858 Obj
if reg20791 == True {
reg20792 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2680 := reg20792
_ = V2680
reg20793 := PrimIsPair(V2680)
var reg20815 Obj
if reg20793 == True {
reg20794 := PrimHead(V2680)
reg20795 := __e.Call(__defun__shen_4lazyderef, reg20794, V2753)
V2681 := reg20795
_ = V2681
reg20796 := MakeSymbol("shen.process-datatype")
reg20797 := PrimEqual(reg20796, V2681)
var reg20813 Obj
if reg20797 == True {
reg20798 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2682 := reg20798
_ = V2682
reg20799 := MakeSymbol("symbol")
reg20800 := PrimEqual(reg20799, V2682)
var reg20811 Obj
if reg20800 == True {
reg20801 := __e.Call(__defun__shen_4incinfs)
_ = reg20801
reg20802 := __e.Call(__defun__thaw, V2754)
reg20811 = reg20802
} else {
reg20803 := __e.Call(__defun__shen_4pvar_2, V2682)
var reg20810 Obj
if reg20803 == True {
reg20804 := MakeSymbol("symbol")
reg20805 := __e.Call(__defun__shen_4bindv, V2682, reg20804, V2753)
_ = reg20805
reg20806 := __e.Call(__defun__shen_4incinfs)
_ = reg20806
reg20807 := __e.Call(__defun__thaw, V2754)
Result := reg20807
_ = Result
reg20808 := __e.Call(__defun__shen_4unbindv, V2682, V2753)
_ = reg20808
reg20810 = Result
} else {
reg20809 := False;
reg20810 = reg20809
}
reg20811 = reg20810
}
reg20813 = reg20811
} else {
reg20812 := False;
reg20813 = reg20812
}
reg20815 = reg20813
} else {
reg20814 := False;
reg20815 = reg20814
}
Case := reg20815
_ = Case
reg20816 := False;
reg20817 := PrimEqual(Case, reg20816)
var reg20857 Obj
if reg20817 == True {
reg20818 := __e.Call(__defun__shen_4lazyderef, V2750, V2753)
V2683 := reg20818
_ = V2683
reg20819 := PrimIsPair(V2683)
var reg20841 Obj
if reg20819 == True {
reg20820 := PrimHead(V2683)
reg20821 := __e.Call(__defun__shen_4lazyderef, reg20820, V2753)
V2684 := reg20821
_ = V2684
reg20822 := MakeSymbol("shen.synonyms-help")
reg20823 := PrimEqual(reg20822, V2684)
var reg20839 Obj
if reg20823 == True {
reg20824 := __e.Call(__defun__shen_4lazyderef, V2751, V2753)
V2685 := reg20824
_ = V2685
reg20825 := MakeSymbol("symbol")
reg20826 := PrimEqual(reg20825, V2685)
var reg20837 Obj
if reg20826 == True {
reg20827 := __e.Call(__defun__shen_4incinfs)
_ = reg20827
reg20828 := __e.Call(__defun__thaw, V2754)
reg20837 = reg20828
} else {
reg20829 := __e.Call(__defun__shen_4pvar_2, V2685)
var reg20836 Obj
if reg20829 == True {
reg20830 := MakeSymbol("symbol")
reg20831 := __e.Call(__defun__shen_4bindv, V2685, reg20830, V2753)
_ = reg20831
reg20832 := __e.Call(__defun__shen_4incinfs)
_ = reg20832
reg20833 := __e.Call(__defun__thaw, V2754)
Result := reg20833
_ = Result
reg20834 := __e.Call(__defun__shen_4unbindv, V2685, V2753)
_ = reg20834
reg20836 = Result
} else {
reg20835 := False;
reg20836 = reg20835
}
reg20837 = reg20836
}
reg20839 = reg20837
} else {
reg20838 := False;
reg20839 = reg20838
}
reg20841 = reg20839
} else {
reg20840 := False;
reg20841 = reg20840
}
Case := reg20841
_ = Case
reg20842 := False;
reg20843 := PrimEqual(Case, reg20842)
var reg20856 Obj
if reg20843 == True {
reg20844 := __e.Call(__defun__shen_4newpv, V2753)
Datatypes := reg20844
_ = Datatypes
reg20845 := __e.Call(__defun__shen_4incinfs)
_ = reg20845
reg20846 := MakeSymbol("shen.*datatypes*")
reg20847 := PrimValue(reg20846)
reg20848 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg20849 := MakeSymbol(":")
reg20850 := Nil;
reg20851 := PrimCons(V2751, reg20850)
reg20852 := PrimCons(reg20849, reg20851)
reg20853 := PrimCons(V2750, reg20852)
__ctx.TailApply(__defun__shen_4udefs_d, reg20853, V2752, Datatypes, V2753, V2754)
return
}, 0)
reg20855 := __e.Call(__defun__bind, Datatypes, reg20847, V2753, reg20848)
reg20856 = reg20855
} else {
reg20856 = Case
}
reg20857 = reg20856
} else {
reg20857 = Case
}
reg20858 = reg20857
} else {
reg20858 = Case
}
reg20859 = reg20858
} else {
reg20859 = Case
}
reg20860 = reg20859
} else {
reg20860 = Case
}
reg20861 = reg20860
} else {
reg20861 = Case
}
reg20862 = reg20861
} else {
reg20862 = Case
}
reg20863 = reg20862
} else {
reg20863 = Case
}
reg20864 = reg20863
} else {
reg20864 = Case
}
reg20865 = reg20864
} else {
reg20865 = Case
}
reg20866 = reg20865
} else {
reg20866 = Case
}
reg20867 = reg20866
} else {
reg20867 = Case
}
reg20868 = reg20867
} else {
reg20868 = Case
}
reg20869 = reg20868
} else {
reg20869 = Case
}
reg20870 = reg20869
} else {
reg20870 = Case
}
reg20871 = reg20870
} else {
reg20871 = Case
}
reg20872 = reg20871
} else {
reg20872 = Case
}
reg20873 = reg20872
} else {
reg20873 = Case
}
reg20874 = reg20873
} else {
reg20874 = Case
}
reg20875 = reg20874
} else {
reg20875 = Case
}
reg20876 = reg20875
} else {
reg20876 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg20876)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.th*", value: __defun__shen_4th_d})

__defun__shen_4t_d_1hyps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2759 := __args[0]
_ = V2759
V2760 := __args[1]
_ = V2760
V2761 := __args[2]
_ = V2761
V2762 := __args[3]
_ = V2762
reg20878 := __e.Call(__defun__shen_4lazyderef, V2759, V2761)
V2496 := reg20878
_ = V2496
reg20879 := PrimIsPair(V2496)
var reg21417 Obj
if reg20879 == True {
reg20880 := PrimHead(V2496)
reg20881 := __e.Call(__defun__shen_4lazyderef, reg20880, V2761)
V2497 := reg20881
_ = V2497
reg20882 := PrimIsPair(V2497)
var reg21415 Obj
if reg20882 == True {
reg20883 := PrimHead(V2497)
reg20884 := __e.Call(__defun__shen_4lazyderef, reg20883, V2761)
V2498 := reg20884
_ = V2498
reg20885 := PrimIsPair(V2498)
var reg21413 Obj
if reg20885 == True {
reg20886 := PrimHead(V2498)
reg20887 := __e.Call(__defun__shen_4lazyderef, reg20886, V2761)
V2499 := reg20887
_ = V2499
reg20888 := MakeSymbol("cons")
reg20889 := PrimEqual(reg20888, V2499)
var reg21411 Obj
if reg20889 == True {
reg20890 := PrimTail(V2498)
reg20891 := __e.Call(__defun__shen_4lazyderef, reg20890, V2761)
V2500 := reg20891
_ = V2500
reg20892 := PrimIsPair(V2500)
var reg21409 Obj
if reg20892 == True {
reg20893 := PrimHead(V2500)
X := reg20893
_ = X
reg20894 := PrimTail(V2500)
reg20895 := __e.Call(__defun__shen_4lazyderef, reg20894, V2761)
V2501 := reg20895
_ = V2501
reg20896 := PrimIsPair(V2501)
var reg21407 Obj
if reg20896 == True {
reg20897 := PrimHead(V2501)
Y := reg20897
_ = Y
reg20898 := PrimTail(V2501)
reg20899 := __e.Call(__defun__shen_4lazyderef, reg20898, V2761)
V2502 := reg20899
_ = V2502
reg20900 := Nil;
reg20901 := PrimEqual(reg20900, V2502)
var reg21405 Obj
if reg20901 == True {
reg20902 := PrimTail(V2497)
reg20903 := __e.Call(__defun__shen_4lazyderef, reg20902, V2761)
V2503 := reg20903
_ = V2503
reg20904 := PrimIsPair(V2503)
var reg21403 Obj
if reg20904 == True {
reg20905 := PrimHead(V2503)
reg20906 := __e.Call(__defun__shen_4lazyderef, reg20905, V2761)
V2504 := reg20906
_ = V2504
reg20907 := MakeSymbol(":")
reg20908 := PrimEqual(reg20907, V2504)
var reg21401 Obj
if reg20908 == True {
reg20909 := PrimTail(V2503)
reg20910 := __e.Call(__defun__shen_4lazyderef, reg20909, V2761)
V2505 := reg20910
_ = V2505
reg20911 := PrimIsPair(V2505)
var reg21399 Obj
if reg20911 == True {
reg20912 := PrimHead(V2505)
reg20913 := __e.Call(__defun__shen_4lazyderef, reg20912, V2761)
V2506 := reg20913
_ = V2506
reg20914 := PrimIsPair(V2506)
var reg21397 Obj
if reg20914 == True {
reg20915 := PrimHead(V2506)
reg20916 := __e.Call(__defun__shen_4lazyderef, reg20915, V2761)
V2507 := reg20916
_ = V2507
reg20917 := MakeSymbol("list")
reg20918 := PrimEqual(reg20917, V2507)
var reg21327 Obj
if reg20918 == True {
reg20919 := PrimTail(V2506)
reg20920 := __e.Call(__defun__shen_4lazyderef, reg20919, V2761)
V2508 := reg20920
_ = V2508
reg20921 := PrimIsPair(V2508)
var reg21119 Obj
if reg20921 == True {
reg20922 := PrimHead(V2508)
A := reg20922
_ = A
reg20923 := PrimTail(V2508)
reg20924 := __e.Call(__defun__shen_4lazyderef, reg20923, V2761)
V2509 := reg20924
_ = V2509
reg20925 := Nil;
reg20926 := PrimEqual(reg20925, V2509)
var reg21051 Obj
if reg20926 == True {
reg20927 := PrimTail(V2505)
reg20928 := __e.Call(__defun__shen_4lazyderef, reg20927, V2761)
V2510 := reg20928
_ = V2510
reg20929 := Nil;
reg20930 := PrimEqual(reg20929, V2510)
var reg20985 Obj
if reg20930 == True {
reg20931 := PrimTail(V2496)
Hyp := reg20931
_ = Hyp
reg20932 := __e.Call(__defun__shen_4incinfs)
_ = reg20932
reg20933 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg20934 := MakeSymbol(":")
reg20935 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg20936 := Nil;
reg20937 := PrimCons(reg20935, reg20936)
reg20938 := PrimCons(reg20934, reg20937)
reg20939 := PrimCons(reg20933, reg20938)
reg20940 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg20941 := MakeSymbol(":")
reg20942 := MakeSymbol("list")
reg20943 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg20944 := Nil;
reg20945 := PrimCons(reg20943, reg20944)
reg20946 := PrimCons(reg20942, reg20945)
reg20947 := Nil;
reg20948 := PrimCons(reg20946, reg20947)
reg20949 := PrimCons(reg20941, reg20948)
reg20950 := PrimCons(reg20940, reg20949)
reg20951 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg20952 := PrimCons(reg20950, reg20951)
reg20953 := PrimCons(reg20939, reg20952)
reg20954 := __e.Call(__defun__bind, V2760, reg20953, V2761, V2762)
reg20985 = reg20954
} else {
reg20955 := __e.Call(__defun__shen_4pvar_2, V2510)
var reg20984 Obj
if reg20955 == True {
reg20956 := Nil;
reg20957 := __e.Call(__defun__shen_4bindv, V2510, reg20956, V2761)
_ = reg20957
reg20958 := PrimTail(V2496)
Hyp := reg20958
_ = Hyp
reg20959 := __e.Call(__defun__shen_4incinfs)
_ = reg20959
reg20960 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg20961 := MakeSymbol(":")
reg20962 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg20963 := Nil;
reg20964 := PrimCons(reg20962, reg20963)
reg20965 := PrimCons(reg20961, reg20964)
reg20966 := PrimCons(reg20960, reg20965)
reg20967 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg20968 := MakeSymbol(":")
reg20969 := MakeSymbol("list")
reg20970 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg20971 := Nil;
reg20972 := PrimCons(reg20970, reg20971)
reg20973 := PrimCons(reg20969, reg20972)
reg20974 := Nil;
reg20975 := PrimCons(reg20973, reg20974)
reg20976 := PrimCons(reg20968, reg20975)
reg20977 := PrimCons(reg20967, reg20976)
reg20978 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg20979 := PrimCons(reg20977, reg20978)
reg20980 := PrimCons(reg20966, reg20979)
reg20981 := __e.Call(__defun__bind, V2760, reg20980, V2761, V2762)
Result := reg20981
_ = Result
reg20982 := __e.Call(__defun__shen_4unbindv, V2510, V2761)
_ = reg20982
reg20984 = Result
} else {
reg20983 := False;
reg20984 = reg20983
}
reg20985 = reg20984
}
reg21051 = reg20985
} else {
reg20986 := __e.Call(__defun__shen_4pvar_2, V2509)
var reg21050 Obj
if reg20986 == True {
reg20987 := Nil;
reg20988 := __e.Call(__defun__shen_4bindv, V2509, reg20987, V2761)
_ = reg20988
reg20989 := PrimTail(V2505)
reg20990 := __e.Call(__defun__shen_4lazyderef, reg20989, V2761)
V2511 := reg20990
_ = V2511
reg20991 := Nil;
reg20992 := PrimEqual(reg20991, V2511)
var reg21047 Obj
if reg20992 == True {
reg20993 := PrimTail(V2496)
Hyp := reg20993
_ = Hyp
reg20994 := __e.Call(__defun__shen_4incinfs)
_ = reg20994
reg20995 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg20996 := MakeSymbol(":")
reg20997 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg20998 := Nil;
reg20999 := PrimCons(reg20997, reg20998)
reg21000 := PrimCons(reg20996, reg20999)
reg21001 := PrimCons(reg20995, reg21000)
reg21002 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21003 := MakeSymbol(":")
reg21004 := MakeSymbol("list")
reg21005 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21006 := Nil;
reg21007 := PrimCons(reg21005, reg21006)
reg21008 := PrimCons(reg21004, reg21007)
reg21009 := Nil;
reg21010 := PrimCons(reg21008, reg21009)
reg21011 := PrimCons(reg21003, reg21010)
reg21012 := PrimCons(reg21002, reg21011)
reg21013 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21014 := PrimCons(reg21012, reg21013)
reg21015 := PrimCons(reg21001, reg21014)
reg21016 := __e.Call(__defun__bind, V2760, reg21015, V2761, V2762)
reg21047 = reg21016
} else {
reg21017 := __e.Call(__defun__shen_4pvar_2, V2511)
var reg21046 Obj
if reg21017 == True {
reg21018 := Nil;
reg21019 := __e.Call(__defun__shen_4bindv, V2511, reg21018, V2761)
_ = reg21019
reg21020 := PrimTail(V2496)
Hyp := reg21020
_ = Hyp
reg21021 := __e.Call(__defun__shen_4incinfs)
_ = reg21021
reg21022 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21023 := MakeSymbol(":")
reg21024 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21025 := Nil;
reg21026 := PrimCons(reg21024, reg21025)
reg21027 := PrimCons(reg21023, reg21026)
reg21028 := PrimCons(reg21022, reg21027)
reg21029 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21030 := MakeSymbol(":")
reg21031 := MakeSymbol("list")
reg21032 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21033 := Nil;
reg21034 := PrimCons(reg21032, reg21033)
reg21035 := PrimCons(reg21031, reg21034)
reg21036 := Nil;
reg21037 := PrimCons(reg21035, reg21036)
reg21038 := PrimCons(reg21030, reg21037)
reg21039 := PrimCons(reg21029, reg21038)
reg21040 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21041 := PrimCons(reg21039, reg21040)
reg21042 := PrimCons(reg21028, reg21041)
reg21043 := __e.Call(__defun__bind, V2760, reg21042, V2761, V2762)
Result := reg21043
_ = Result
reg21044 := __e.Call(__defun__shen_4unbindv, V2511, V2761)
_ = reg21044
reg21046 = Result
} else {
reg21045 := False;
reg21046 = reg21045
}
reg21047 = reg21046
}
Result := reg21047
_ = Result
reg21048 := __e.Call(__defun__shen_4unbindv, V2509, V2761)
_ = reg21048
reg21050 = Result
} else {
reg21049 := False;
reg21050 = reg21049
}
reg21051 = reg21050
}
reg21119 = reg21051
} else {
reg21052 := __e.Call(__defun__shen_4pvar_2, V2508)
var reg21118 Obj
if reg21052 == True {
reg21053 := __e.Call(__defun__shen_4newpv, V2761)
A := reg21053
_ = A
reg21054 := Nil;
reg21055 := PrimCons(A, reg21054)
reg21056 := __e.Call(__defun__shen_4bindv, V2508, reg21055, V2761)
_ = reg21056
reg21057 := PrimTail(V2505)
reg21058 := __e.Call(__defun__shen_4lazyderef, reg21057, V2761)
V2512 := reg21058
_ = V2512
reg21059 := Nil;
reg21060 := PrimEqual(reg21059, V2512)
var reg21115 Obj
if reg21060 == True {
reg21061 := PrimTail(V2496)
Hyp := reg21061
_ = Hyp
reg21062 := __e.Call(__defun__shen_4incinfs)
_ = reg21062
reg21063 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21064 := MakeSymbol(":")
reg21065 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21066 := Nil;
reg21067 := PrimCons(reg21065, reg21066)
reg21068 := PrimCons(reg21064, reg21067)
reg21069 := PrimCons(reg21063, reg21068)
reg21070 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21071 := MakeSymbol(":")
reg21072 := MakeSymbol("list")
reg21073 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21074 := Nil;
reg21075 := PrimCons(reg21073, reg21074)
reg21076 := PrimCons(reg21072, reg21075)
reg21077 := Nil;
reg21078 := PrimCons(reg21076, reg21077)
reg21079 := PrimCons(reg21071, reg21078)
reg21080 := PrimCons(reg21070, reg21079)
reg21081 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21082 := PrimCons(reg21080, reg21081)
reg21083 := PrimCons(reg21069, reg21082)
reg21084 := __e.Call(__defun__bind, V2760, reg21083, V2761, V2762)
reg21115 = reg21084
} else {
reg21085 := __e.Call(__defun__shen_4pvar_2, V2512)
var reg21114 Obj
if reg21085 == True {
reg21086 := Nil;
reg21087 := __e.Call(__defun__shen_4bindv, V2512, reg21086, V2761)
_ = reg21087
reg21088 := PrimTail(V2496)
Hyp := reg21088
_ = Hyp
reg21089 := __e.Call(__defun__shen_4incinfs)
_ = reg21089
reg21090 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21091 := MakeSymbol(":")
reg21092 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21093 := Nil;
reg21094 := PrimCons(reg21092, reg21093)
reg21095 := PrimCons(reg21091, reg21094)
reg21096 := PrimCons(reg21090, reg21095)
reg21097 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21098 := MakeSymbol(":")
reg21099 := MakeSymbol("list")
reg21100 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21101 := Nil;
reg21102 := PrimCons(reg21100, reg21101)
reg21103 := PrimCons(reg21099, reg21102)
reg21104 := Nil;
reg21105 := PrimCons(reg21103, reg21104)
reg21106 := PrimCons(reg21098, reg21105)
reg21107 := PrimCons(reg21097, reg21106)
reg21108 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21109 := PrimCons(reg21107, reg21108)
reg21110 := PrimCons(reg21096, reg21109)
reg21111 := __e.Call(__defun__bind, V2760, reg21110, V2761, V2762)
Result := reg21111
_ = Result
reg21112 := __e.Call(__defun__shen_4unbindv, V2512, V2761)
_ = reg21112
reg21114 = Result
} else {
reg21113 := False;
reg21114 = reg21113
}
reg21115 = reg21114
}
Result := reg21115
_ = Result
reg21116 := __e.Call(__defun__shen_4unbindv, V2508, V2761)
_ = reg21116
reg21118 = Result
} else {
reg21117 := False;
reg21118 = reg21117
}
reg21119 = reg21118
}
reg21327 = reg21119
} else {
reg21120 := __e.Call(__defun__shen_4pvar_2, V2507)
var reg21326 Obj
if reg21120 == True {
reg21121 := MakeSymbol("list")
reg21122 := __e.Call(__defun__shen_4bindv, V2507, reg21121, V2761)
_ = reg21122
reg21123 := PrimTail(V2506)
reg21124 := __e.Call(__defun__shen_4lazyderef, reg21123, V2761)
V2513 := reg21124
_ = V2513
reg21125 := PrimIsPair(V2513)
var reg21323 Obj
if reg21125 == True {
reg21126 := PrimHead(V2513)
A := reg21126
_ = A
reg21127 := PrimTail(V2513)
reg21128 := __e.Call(__defun__shen_4lazyderef, reg21127, V2761)
V2514 := reg21128
_ = V2514
reg21129 := Nil;
reg21130 := PrimEqual(reg21129, V2514)
var reg21255 Obj
if reg21130 == True {
reg21131 := PrimTail(V2505)
reg21132 := __e.Call(__defun__shen_4lazyderef, reg21131, V2761)
V2515 := reg21132
_ = V2515
reg21133 := Nil;
reg21134 := PrimEqual(reg21133, V2515)
var reg21189 Obj
if reg21134 == True {
reg21135 := PrimTail(V2496)
Hyp := reg21135
_ = Hyp
reg21136 := __e.Call(__defun__shen_4incinfs)
_ = reg21136
reg21137 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21138 := MakeSymbol(":")
reg21139 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21140 := Nil;
reg21141 := PrimCons(reg21139, reg21140)
reg21142 := PrimCons(reg21138, reg21141)
reg21143 := PrimCons(reg21137, reg21142)
reg21144 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21145 := MakeSymbol(":")
reg21146 := MakeSymbol("list")
reg21147 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21148 := Nil;
reg21149 := PrimCons(reg21147, reg21148)
reg21150 := PrimCons(reg21146, reg21149)
reg21151 := Nil;
reg21152 := PrimCons(reg21150, reg21151)
reg21153 := PrimCons(reg21145, reg21152)
reg21154 := PrimCons(reg21144, reg21153)
reg21155 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21156 := PrimCons(reg21154, reg21155)
reg21157 := PrimCons(reg21143, reg21156)
reg21158 := __e.Call(__defun__bind, V2760, reg21157, V2761, V2762)
reg21189 = reg21158
} else {
reg21159 := __e.Call(__defun__shen_4pvar_2, V2515)
var reg21188 Obj
if reg21159 == True {
reg21160 := Nil;
reg21161 := __e.Call(__defun__shen_4bindv, V2515, reg21160, V2761)
_ = reg21161
reg21162 := PrimTail(V2496)
Hyp := reg21162
_ = Hyp
reg21163 := __e.Call(__defun__shen_4incinfs)
_ = reg21163
reg21164 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21165 := MakeSymbol(":")
reg21166 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21167 := Nil;
reg21168 := PrimCons(reg21166, reg21167)
reg21169 := PrimCons(reg21165, reg21168)
reg21170 := PrimCons(reg21164, reg21169)
reg21171 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21172 := MakeSymbol(":")
reg21173 := MakeSymbol("list")
reg21174 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21175 := Nil;
reg21176 := PrimCons(reg21174, reg21175)
reg21177 := PrimCons(reg21173, reg21176)
reg21178 := Nil;
reg21179 := PrimCons(reg21177, reg21178)
reg21180 := PrimCons(reg21172, reg21179)
reg21181 := PrimCons(reg21171, reg21180)
reg21182 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21183 := PrimCons(reg21181, reg21182)
reg21184 := PrimCons(reg21170, reg21183)
reg21185 := __e.Call(__defun__bind, V2760, reg21184, V2761, V2762)
Result := reg21185
_ = Result
reg21186 := __e.Call(__defun__shen_4unbindv, V2515, V2761)
_ = reg21186
reg21188 = Result
} else {
reg21187 := False;
reg21188 = reg21187
}
reg21189 = reg21188
}
reg21255 = reg21189
} else {
reg21190 := __e.Call(__defun__shen_4pvar_2, V2514)
var reg21254 Obj
if reg21190 == True {
reg21191 := Nil;
reg21192 := __e.Call(__defun__shen_4bindv, V2514, reg21191, V2761)
_ = reg21192
reg21193 := PrimTail(V2505)
reg21194 := __e.Call(__defun__shen_4lazyderef, reg21193, V2761)
V2516 := reg21194
_ = V2516
reg21195 := Nil;
reg21196 := PrimEqual(reg21195, V2516)
var reg21251 Obj
if reg21196 == True {
reg21197 := PrimTail(V2496)
Hyp := reg21197
_ = Hyp
reg21198 := __e.Call(__defun__shen_4incinfs)
_ = reg21198
reg21199 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21200 := MakeSymbol(":")
reg21201 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21202 := Nil;
reg21203 := PrimCons(reg21201, reg21202)
reg21204 := PrimCons(reg21200, reg21203)
reg21205 := PrimCons(reg21199, reg21204)
reg21206 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21207 := MakeSymbol(":")
reg21208 := MakeSymbol("list")
reg21209 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21210 := Nil;
reg21211 := PrimCons(reg21209, reg21210)
reg21212 := PrimCons(reg21208, reg21211)
reg21213 := Nil;
reg21214 := PrimCons(reg21212, reg21213)
reg21215 := PrimCons(reg21207, reg21214)
reg21216 := PrimCons(reg21206, reg21215)
reg21217 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21218 := PrimCons(reg21216, reg21217)
reg21219 := PrimCons(reg21205, reg21218)
reg21220 := __e.Call(__defun__bind, V2760, reg21219, V2761, V2762)
reg21251 = reg21220
} else {
reg21221 := __e.Call(__defun__shen_4pvar_2, V2516)
var reg21250 Obj
if reg21221 == True {
reg21222 := Nil;
reg21223 := __e.Call(__defun__shen_4bindv, V2516, reg21222, V2761)
_ = reg21223
reg21224 := PrimTail(V2496)
Hyp := reg21224
_ = Hyp
reg21225 := __e.Call(__defun__shen_4incinfs)
_ = reg21225
reg21226 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21227 := MakeSymbol(":")
reg21228 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21229 := Nil;
reg21230 := PrimCons(reg21228, reg21229)
reg21231 := PrimCons(reg21227, reg21230)
reg21232 := PrimCons(reg21226, reg21231)
reg21233 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21234 := MakeSymbol(":")
reg21235 := MakeSymbol("list")
reg21236 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21237 := Nil;
reg21238 := PrimCons(reg21236, reg21237)
reg21239 := PrimCons(reg21235, reg21238)
reg21240 := Nil;
reg21241 := PrimCons(reg21239, reg21240)
reg21242 := PrimCons(reg21234, reg21241)
reg21243 := PrimCons(reg21233, reg21242)
reg21244 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21245 := PrimCons(reg21243, reg21244)
reg21246 := PrimCons(reg21232, reg21245)
reg21247 := __e.Call(__defun__bind, V2760, reg21246, V2761, V2762)
Result := reg21247
_ = Result
reg21248 := __e.Call(__defun__shen_4unbindv, V2516, V2761)
_ = reg21248
reg21250 = Result
} else {
reg21249 := False;
reg21250 = reg21249
}
reg21251 = reg21250
}
Result := reg21251
_ = Result
reg21252 := __e.Call(__defun__shen_4unbindv, V2514, V2761)
_ = reg21252
reg21254 = Result
} else {
reg21253 := False;
reg21254 = reg21253
}
reg21255 = reg21254
}
reg21323 = reg21255
} else {
reg21256 := __e.Call(__defun__shen_4pvar_2, V2513)
var reg21322 Obj
if reg21256 == True {
reg21257 := __e.Call(__defun__shen_4newpv, V2761)
A := reg21257
_ = A
reg21258 := Nil;
reg21259 := PrimCons(A, reg21258)
reg21260 := __e.Call(__defun__shen_4bindv, V2513, reg21259, V2761)
_ = reg21260
reg21261 := PrimTail(V2505)
reg21262 := __e.Call(__defun__shen_4lazyderef, reg21261, V2761)
V2517 := reg21262
_ = V2517
reg21263 := Nil;
reg21264 := PrimEqual(reg21263, V2517)
var reg21319 Obj
if reg21264 == True {
reg21265 := PrimTail(V2496)
Hyp := reg21265
_ = Hyp
reg21266 := __e.Call(__defun__shen_4incinfs)
_ = reg21266
reg21267 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21268 := MakeSymbol(":")
reg21269 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21270 := Nil;
reg21271 := PrimCons(reg21269, reg21270)
reg21272 := PrimCons(reg21268, reg21271)
reg21273 := PrimCons(reg21267, reg21272)
reg21274 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21275 := MakeSymbol(":")
reg21276 := MakeSymbol("list")
reg21277 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21278 := Nil;
reg21279 := PrimCons(reg21277, reg21278)
reg21280 := PrimCons(reg21276, reg21279)
reg21281 := Nil;
reg21282 := PrimCons(reg21280, reg21281)
reg21283 := PrimCons(reg21275, reg21282)
reg21284 := PrimCons(reg21274, reg21283)
reg21285 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21286 := PrimCons(reg21284, reg21285)
reg21287 := PrimCons(reg21273, reg21286)
reg21288 := __e.Call(__defun__bind, V2760, reg21287, V2761, V2762)
reg21319 = reg21288
} else {
reg21289 := __e.Call(__defun__shen_4pvar_2, V2517)
var reg21318 Obj
if reg21289 == True {
reg21290 := Nil;
reg21291 := __e.Call(__defun__shen_4bindv, V2517, reg21290, V2761)
_ = reg21291
reg21292 := PrimTail(V2496)
Hyp := reg21292
_ = Hyp
reg21293 := __e.Call(__defun__shen_4incinfs)
_ = reg21293
reg21294 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21295 := MakeSymbol(":")
reg21296 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21297 := Nil;
reg21298 := PrimCons(reg21296, reg21297)
reg21299 := PrimCons(reg21295, reg21298)
reg21300 := PrimCons(reg21294, reg21299)
reg21301 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21302 := MakeSymbol(":")
reg21303 := MakeSymbol("list")
reg21304 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21305 := Nil;
reg21306 := PrimCons(reg21304, reg21305)
reg21307 := PrimCons(reg21303, reg21306)
reg21308 := Nil;
reg21309 := PrimCons(reg21307, reg21308)
reg21310 := PrimCons(reg21302, reg21309)
reg21311 := PrimCons(reg21301, reg21310)
reg21312 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21313 := PrimCons(reg21311, reg21312)
reg21314 := PrimCons(reg21300, reg21313)
reg21315 := __e.Call(__defun__bind, V2760, reg21314, V2761, V2762)
Result := reg21315
_ = Result
reg21316 := __e.Call(__defun__shen_4unbindv, V2517, V2761)
_ = reg21316
reg21318 = Result
} else {
reg21317 := False;
reg21318 = reg21317
}
reg21319 = reg21318
}
Result := reg21319
_ = Result
reg21320 := __e.Call(__defun__shen_4unbindv, V2513, V2761)
_ = reg21320
reg21322 = Result
} else {
reg21321 := False;
reg21322 = reg21321
}
reg21323 = reg21322
}
Result := reg21323
_ = Result
reg21324 := __e.Call(__defun__shen_4unbindv, V2507, V2761)
_ = reg21324
reg21326 = Result
} else {
reg21325 := False;
reg21326 = reg21325
}
reg21327 = reg21326
}
reg21397 = reg21327
} else {
reg21328 := __e.Call(__defun__shen_4pvar_2, V2506)
var reg21396 Obj
if reg21328 == True {
reg21329 := __e.Call(__defun__shen_4newpv, V2761)
A := reg21329
_ = A
reg21330 := MakeSymbol("list")
reg21331 := Nil;
reg21332 := PrimCons(A, reg21331)
reg21333 := PrimCons(reg21330, reg21332)
reg21334 := __e.Call(__defun__shen_4bindv, V2506, reg21333, V2761)
_ = reg21334
reg21335 := PrimTail(V2505)
reg21336 := __e.Call(__defun__shen_4lazyderef, reg21335, V2761)
V2518 := reg21336
_ = V2518
reg21337 := Nil;
reg21338 := PrimEqual(reg21337, V2518)
var reg21393 Obj
if reg21338 == True {
reg21339 := PrimTail(V2496)
Hyp := reg21339
_ = Hyp
reg21340 := __e.Call(__defun__shen_4incinfs)
_ = reg21340
reg21341 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21342 := MakeSymbol(":")
reg21343 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21344 := Nil;
reg21345 := PrimCons(reg21343, reg21344)
reg21346 := PrimCons(reg21342, reg21345)
reg21347 := PrimCons(reg21341, reg21346)
reg21348 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21349 := MakeSymbol(":")
reg21350 := MakeSymbol("list")
reg21351 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21352 := Nil;
reg21353 := PrimCons(reg21351, reg21352)
reg21354 := PrimCons(reg21350, reg21353)
reg21355 := Nil;
reg21356 := PrimCons(reg21354, reg21355)
reg21357 := PrimCons(reg21349, reg21356)
reg21358 := PrimCons(reg21348, reg21357)
reg21359 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21360 := PrimCons(reg21358, reg21359)
reg21361 := PrimCons(reg21347, reg21360)
reg21362 := __e.Call(__defun__bind, V2760, reg21361, V2761, V2762)
reg21393 = reg21362
} else {
reg21363 := __e.Call(__defun__shen_4pvar_2, V2518)
var reg21392 Obj
if reg21363 == True {
reg21364 := Nil;
reg21365 := __e.Call(__defun__shen_4bindv, V2518, reg21364, V2761)
_ = reg21365
reg21366 := PrimTail(V2496)
Hyp := reg21366
_ = Hyp
reg21367 := __e.Call(__defun__shen_4incinfs)
_ = reg21367
reg21368 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21369 := MakeSymbol(":")
reg21370 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21371 := Nil;
reg21372 := PrimCons(reg21370, reg21371)
reg21373 := PrimCons(reg21369, reg21372)
reg21374 := PrimCons(reg21368, reg21373)
reg21375 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21376 := MakeSymbol(":")
reg21377 := MakeSymbol("list")
reg21378 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21379 := Nil;
reg21380 := PrimCons(reg21378, reg21379)
reg21381 := PrimCons(reg21377, reg21380)
reg21382 := Nil;
reg21383 := PrimCons(reg21381, reg21382)
reg21384 := PrimCons(reg21376, reg21383)
reg21385 := PrimCons(reg21375, reg21384)
reg21386 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21387 := PrimCons(reg21385, reg21386)
reg21388 := PrimCons(reg21374, reg21387)
reg21389 := __e.Call(__defun__bind, V2760, reg21388, V2761, V2762)
Result := reg21389
_ = Result
reg21390 := __e.Call(__defun__shen_4unbindv, V2518, V2761)
_ = reg21390
reg21392 = Result
} else {
reg21391 := False;
reg21392 = reg21391
}
reg21393 = reg21392
}
Result := reg21393
_ = Result
reg21394 := __e.Call(__defun__shen_4unbindv, V2506, V2761)
_ = reg21394
reg21396 = Result
} else {
reg21395 := False;
reg21396 = reg21395
}
reg21397 = reg21396
}
reg21399 = reg21397
} else {
reg21398 := False;
reg21399 = reg21398
}
reg21401 = reg21399
} else {
reg21400 := False;
reg21401 = reg21400
}
reg21403 = reg21401
} else {
reg21402 := False;
reg21403 = reg21402
}
reg21405 = reg21403
} else {
reg21404 := False;
reg21405 = reg21404
}
reg21407 = reg21405
} else {
reg21406 := False;
reg21407 = reg21406
}
reg21409 = reg21407
} else {
reg21408 := False;
reg21409 = reg21408
}
reg21411 = reg21409
} else {
reg21410 := False;
reg21411 = reg21410
}
reg21413 = reg21411
} else {
reg21412 := False;
reg21413 = reg21412
}
reg21415 = reg21413
} else {
reg21414 := False;
reg21415 = reg21414
}
reg21417 = reg21415
} else {
reg21416 := False;
reg21417 = reg21416
}
Case := reg21417
_ = Case
reg21418 := False;
reg21419 := PrimEqual(Case, reg21418)
if reg21419 == True {
reg21420 := __e.Call(__defun__shen_4lazyderef, V2759, V2761)
V2519 := reg21420
_ = V2519
reg21421 := PrimIsPair(V2519)
var reg21971 Obj
if reg21421 == True {
reg21422 := PrimHead(V2519)
reg21423 := __e.Call(__defun__shen_4lazyderef, reg21422, V2761)
V2520 := reg21423
_ = V2520
reg21424 := PrimIsPair(V2520)
var reg21969 Obj
if reg21424 == True {
reg21425 := PrimHead(V2520)
reg21426 := __e.Call(__defun__shen_4lazyderef, reg21425, V2761)
V2521 := reg21426
_ = V2521
reg21427 := PrimIsPair(V2521)
var reg21967 Obj
if reg21427 == True {
reg21428 := PrimHead(V2521)
reg21429 := __e.Call(__defun__shen_4lazyderef, reg21428, V2761)
V2522 := reg21429
_ = V2522
reg21430 := MakeSymbol("@p")
reg21431 := PrimEqual(reg21430, V2522)
var reg21965 Obj
if reg21431 == True {
reg21432 := PrimTail(V2521)
reg21433 := __e.Call(__defun__shen_4lazyderef, reg21432, V2761)
V2523 := reg21433
_ = V2523
reg21434 := PrimIsPair(V2523)
var reg21963 Obj
if reg21434 == True {
reg21435 := PrimHead(V2523)
X := reg21435
_ = X
reg21436 := PrimTail(V2523)
reg21437 := __e.Call(__defun__shen_4lazyderef, reg21436, V2761)
V2524 := reg21437
_ = V2524
reg21438 := PrimIsPair(V2524)
var reg21961 Obj
if reg21438 == True {
reg21439 := PrimHead(V2524)
Y := reg21439
_ = Y
reg21440 := PrimTail(V2524)
reg21441 := __e.Call(__defun__shen_4lazyderef, reg21440, V2761)
V2525 := reg21441
_ = V2525
reg21442 := Nil;
reg21443 := PrimEqual(reg21442, V2525)
var reg21959 Obj
if reg21443 == True {
reg21444 := PrimTail(V2520)
reg21445 := __e.Call(__defun__shen_4lazyderef, reg21444, V2761)
V2526 := reg21445
_ = V2526
reg21446 := PrimIsPair(V2526)
var reg21957 Obj
if reg21446 == True {
reg21447 := PrimHead(V2526)
reg21448 := __e.Call(__defun__shen_4lazyderef, reg21447, V2761)
V2527 := reg21448
_ = V2527
reg21449 := MakeSymbol(":")
reg21450 := PrimEqual(reg21449, V2527)
var reg21955 Obj
if reg21450 == True {
reg21451 := PrimTail(V2526)
reg21452 := __e.Call(__defun__shen_4lazyderef, reg21451, V2761)
V2528 := reg21452
_ = V2528
reg21453 := PrimIsPair(V2528)
var reg21953 Obj
if reg21453 == True {
reg21454 := PrimHead(V2528)
reg21455 := __e.Call(__defun__shen_4lazyderef, reg21454, V2761)
V2529 := reg21455
_ = V2529
reg21456 := PrimIsPair(V2529)
var reg21951 Obj
if reg21456 == True {
reg21457 := PrimHead(V2529)
A := reg21457
_ = A
reg21458 := PrimTail(V2529)
reg21459 := __e.Call(__defun__shen_4lazyderef, reg21458, V2761)
V2530 := reg21459
_ = V2530
reg21460 := PrimIsPair(V2530)
var reg21887 Obj
if reg21460 == True {
reg21461 := PrimHead(V2530)
reg21462 := __e.Call(__defun__shen_4lazyderef, reg21461, V2761)
V2531 := reg21462
_ = V2531
reg21463 := MakeSymbol("*")
reg21464 := PrimEqual(reg21463, V2531)
var reg21825 Obj
if reg21464 == True {
reg21465 := PrimTail(V2530)
reg21466 := __e.Call(__defun__shen_4lazyderef, reg21465, V2761)
V2532 := reg21466
_ = V2532
reg21467 := PrimIsPair(V2532)
var reg21641 Obj
if reg21467 == True {
reg21468 := PrimHead(V2532)
B := reg21468
_ = B
reg21469 := PrimTail(V2532)
reg21470 := __e.Call(__defun__shen_4lazyderef, reg21469, V2761)
V2533 := reg21470
_ = V2533
reg21471 := Nil;
reg21472 := PrimEqual(reg21471, V2533)
var reg21581 Obj
if reg21472 == True {
reg21473 := PrimTail(V2528)
reg21474 := __e.Call(__defun__shen_4lazyderef, reg21473, V2761)
V2534 := reg21474
_ = V2534
reg21475 := Nil;
reg21476 := PrimEqual(reg21475, V2534)
var reg21523 Obj
if reg21476 == True {
reg21477 := PrimTail(V2519)
Hyp := reg21477
_ = Hyp
reg21478 := __e.Call(__defun__shen_4incinfs)
_ = reg21478
reg21479 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21480 := MakeSymbol(":")
reg21481 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21482 := Nil;
reg21483 := PrimCons(reg21481, reg21482)
reg21484 := PrimCons(reg21480, reg21483)
reg21485 := PrimCons(reg21479, reg21484)
reg21486 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21487 := MakeSymbol(":")
reg21488 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21489 := Nil;
reg21490 := PrimCons(reg21488, reg21489)
reg21491 := PrimCons(reg21487, reg21490)
reg21492 := PrimCons(reg21486, reg21491)
reg21493 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21494 := PrimCons(reg21492, reg21493)
reg21495 := PrimCons(reg21485, reg21494)
reg21496 := __e.Call(__defun__bind, V2760, reg21495, V2761, V2762)
reg21523 = reg21496
} else {
reg21497 := __e.Call(__defun__shen_4pvar_2, V2534)
var reg21522 Obj
if reg21497 == True {
reg21498 := Nil;
reg21499 := __e.Call(__defun__shen_4bindv, V2534, reg21498, V2761)
_ = reg21499
reg21500 := PrimTail(V2519)
Hyp := reg21500
_ = Hyp
reg21501 := __e.Call(__defun__shen_4incinfs)
_ = reg21501
reg21502 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21503 := MakeSymbol(":")
reg21504 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21505 := Nil;
reg21506 := PrimCons(reg21504, reg21505)
reg21507 := PrimCons(reg21503, reg21506)
reg21508 := PrimCons(reg21502, reg21507)
reg21509 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21510 := MakeSymbol(":")
reg21511 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21512 := Nil;
reg21513 := PrimCons(reg21511, reg21512)
reg21514 := PrimCons(reg21510, reg21513)
reg21515 := PrimCons(reg21509, reg21514)
reg21516 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21517 := PrimCons(reg21515, reg21516)
reg21518 := PrimCons(reg21508, reg21517)
reg21519 := __e.Call(__defun__bind, V2760, reg21518, V2761, V2762)
Result := reg21519
_ = Result
reg21520 := __e.Call(__defun__shen_4unbindv, V2534, V2761)
_ = reg21520
reg21522 = Result
} else {
reg21521 := False;
reg21522 = reg21521
}
reg21523 = reg21522
}
reg21581 = reg21523
} else {
reg21524 := __e.Call(__defun__shen_4pvar_2, V2533)
var reg21580 Obj
if reg21524 == True {
reg21525 := Nil;
reg21526 := __e.Call(__defun__shen_4bindv, V2533, reg21525, V2761)
_ = reg21526
reg21527 := PrimTail(V2528)
reg21528 := __e.Call(__defun__shen_4lazyderef, reg21527, V2761)
V2535 := reg21528
_ = V2535
reg21529 := Nil;
reg21530 := PrimEqual(reg21529, V2535)
var reg21577 Obj
if reg21530 == True {
reg21531 := PrimTail(V2519)
Hyp := reg21531
_ = Hyp
reg21532 := __e.Call(__defun__shen_4incinfs)
_ = reg21532
reg21533 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21534 := MakeSymbol(":")
reg21535 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21536 := Nil;
reg21537 := PrimCons(reg21535, reg21536)
reg21538 := PrimCons(reg21534, reg21537)
reg21539 := PrimCons(reg21533, reg21538)
reg21540 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21541 := MakeSymbol(":")
reg21542 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21543 := Nil;
reg21544 := PrimCons(reg21542, reg21543)
reg21545 := PrimCons(reg21541, reg21544)
reg21546 := PrimCons(reg21540, reg21545)
reg21547 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21548 := PrimCons(reg21546, reg21547)
reg21549 := PrimCons(reg21539, reg21548)
reg21550 := __e.Call(__defun__bind, V2760, reg21549, V2761, V2762)
reg21577 = reg21550
} else {
reg21551 := __e.Call(__defun__shen_4pvar_2, V2535)
var reg21576 Obj
if reg21551 == True {
reg21552 := Nil;
reg21553 := __e.Call(__defun__shen_4bindv, V2535, reg21552, V2761)
_ = reg21553
reg21554 := PrimTail(V2519)
Hyp := reg21554
_ = Hyp
reg21555 := __e.Call(__defun__shen_4incinfs)
_ = reg21555
reg21556 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21557 := MakeSymbol(":")
reg21558 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21559 := Nil;
reg21560 := PrimCons(reg21558, reg21559)
reg21561 := PrimCons(reg21557, reg21560)
reg21562 := PrimCons(reg21556, reg21561)
reg21563 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21564 := MakeSymbol(":")
reg21565 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21566 := Nil;
reg21567 := PrimCons(reg21565, reg21566)
reg21568 := PrimCons(reg21564, reg21567)
reg21569 := PrimCons(reg21563, reg21568)
reg21570 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21571 := PrimCons(reg21569, reg21570)
reg21572 := PrimCons(reg21562, reg21571)
reg21573 := __e.Call(__defun__bind, V2760, reg21572, V2761, V2762)
Result := reg21573
_ = Result
reg21574 := __e.Call(__defun__shen_4unbindv, V2535, V2761)
_ = reg21574
reg21576 = Result
} else {
reg21575 := False;
reg21576 = reg21575
}
reg21577 = reg21576
}
Result := reg21577
_ = Result
reg21578 := __e.Call(__defun__shen_4unbindv, V2533, V2761)
_ = reg21578
reg21580 = Result
} else {
reg21579 := False;
reg21580 = reg21579
}
reg21581 = reg21580
}
reg21641 = reg21581
} else {
reg21582 := __e.Call(__defun__shen_4pvar_2, V2532)
var reg21640 Obj
if reg21582 == True {
reg21583 := __e.Call(__defun__shen_4newpv, V2761)
B := reg21583
_ = B
reg21584 := Nil;
reg21585 := PrimCons(B, reg21584)
reg21586 := __e.Call(__defun__shen_4bindv, V2532, reg21585, V2761)
_ = reg21586
reg21587 := PrimTail(V2528)
reg21588 := __e.Call(__defun__shen_4lazyderef, reg21587, V2761)
V2536 := reg21588
_ = V2536
reg21589 := Nil;
reg21590 := PrimEqual(reg21589, V2536)
var reg21637 Obj
if reg21590 == True {
reg21591 := PrimTail(V2519)
Hyp := reg21591
_ = Hyp
reg21592 := __e.Call(__defun__shen_4incinfs)
_ = reg21592
reg21593 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21594 := MakeSymbol(":")
reg21595 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21596 := Nil;
reg21597 := PrimCons(reg21595, reg21596)
reg21598 := PrimCons(reg21594, reg21597)
reg21599 := PrimCons(reg21593, reg21598)
reg21600 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21601 := MakeSymbol(":")
reg21602 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21603 := Nil;
reg21604 := PrimCons(reg21602, reg21603)
reg21605 := PrimCons(reg21601, reg21604)
reg21606 := PrimCons(reg21600, reg21605)
reg21607 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21608 := PrimCons(reg21606, reg21607)
reg21609 := PrimCons(reg21599, reg21608)
reg21610 := __e.Call(__defun__bind, V2760, reg21609, V2761, V2762)
reg21637 = reg21610
} else {
reg21611 := __e.Call(__defun__shen_4pvar_2, V2536)
var reg21636 Obj
if reg21611 == True {
reg21612 := Nil;
reg21613 := __e.Call(__defun__shen_4bindv, V2536, reg21612, V2761)
_ = reg21613
reg21614 := PrimTail(V2519)
Hyp := reg21614
_ = Hyp
reg21615 := __e.Call(__defun__shen_4incinfs)
_ = reg21615
reg21616 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21617 := MakeSymbol(":")
reg21618 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21619 := Nil;
reg21620 := PrimCons(reg21618, reg21619)
reg21621 := PrimCons(reg21617, reg21620)
reg21622 := PrimCons(reg21616, reg21621)
reg21623 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21624 := MakeSymbol(":")
reg21625 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21626 := Nil;
reg21627 := PrimCons(reg21625, reg21626)
reg21628 := PrimCons(reg21624, reg21627)
reg21629 := PrimCons(reg21623, reg21628)
reg21630 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21631 := PrimCons(reg21629, reg21630)
reg21632 := PrimCons(reg21622, reg21631)
reg21633 := __e.Call(__defun__bind, V2760, reg21632, V2761, V2762)
Result := reg21633
_ = Result
reg21634 := __e.Call(__defun__shen_4unbindv, V2536, V2761)
_ = reg21634
reg21636 = Result
} else {
reg21635 := False;
reg21636 = reg21635
}
reg21637 = reg21636
}
Result := reg21637
_ = Result
reg21638 := __e.Call(__defun__shen_4unbindv, V2532, V2761)
_ = reg21638
reg21640 = Result
} else {
reg21639 := False;
reg21640 = reg21639
}
reg21641 = reg21640
}
reg21825 = reg21641
} else {
reg21642 := __e.Call(__defun__shen_4pvar_2, V2531)
var reg21824 Obj
if reg21642 == True {
reg21643 := MakeSymbol("*")
reg21644 := __e.Call(__defun__shen_4bindv, V2531, reg21643, V2761)
_ = reg21644
reg21645 := PrimTail(V2530)
reg21646 := __e.Call(__defun__shen_4lazyderef, reg21645, V2761)
V2537 := reg21646
_ = V2537
reg21647 := PrimIsPair(V2537)
var reg21821 Obj
if reg21647 == True {
reg21648 := PrimHead(V2537)
B := reg21648
_ = B
reg21649 := PrimTail(V2537)
reg21650 := __e.Call(__defun__shen_4lazyderef, reg21649, V2761)
V2538 := reg21650
_ = V2538
reg21651 := Nil;
reg21652 := PrimEqual(reg21651, V2538)
var reg21761 Obj
if reg21652 == True {
reg21653 := PrimTail(V2528)
reg21654 := __e.Call(__defun__shen_4lazyderef, reg21653, V2761)
V2539 := reg21654
_ = V2539
reg21655 := Nil;
reg21656 := PrimEqual(reg21655, V2539)
var reg21703 Obj
if reg21656 == True {
reg21657 := PrimTail(V2519)
Hyp := reg21657
_ = Hyp
reg21658 := __e.Call(__defun__shen_4incinfs)
_ = reg21658
reg21659 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21660 := MakeSymbol(":")
reg21661 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21662 := Nil;
reg21663 := PrimCons(reg21661, reg21662)
reg21664 := PrimCons(reg21660, reg21663)
reg21665 := PrimCons(reg21659, reg21664)
reg21666 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21667 := MakeSymbol(":")
reg21668 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21669 := Nil;
reg21670 := PrimCons(reg21668, reg21669)
reg21671 := PrimCons(reg21667, reg21670)
reg21672 := PrimCons(reg21666, reg21671)
reg21673 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21674 := PrimCons(reg21672, reg21673)
reg21675 := PrimCons(reg21665, reg21674)
reg21676 := __e.Call(__defun__bind, V2760, reg21675, V2761, V2762)
reg21703 = reg21676
} else {
reg21677 := __e.Call(__defun__shen_4pvar_2, V2539)
var reg21702 Obj
if reg21677 == True {
reg21678 := Nil;
reg21679 := __e.Call(__defun__shen_4bindv, V2539, reg21678, V2761)
_ = reg21679
reg21680 := PrimTail(V2519)
Hyp := reg21680
_ = Hyp
reg21681 := __e.Call(__defun__shen_4incinfs)
_ = reg21681
reg21682 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21683 := MakeSymbol(":")
reg21684 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21685 := Nil;
reg21686 := PrimCons(reg21684, reg21685)
reg21687 := PrimCons(reg21683, reg21686)
reg21688 := PrimCons(reg21682, reg21687)
reg21689 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21690 := MakeSymbol(":")
reg21691 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21692 := Nil;
reg21693 := PrimCons(reg21691, reg21692)
reg21694 := PrimCons(reg21690, reg21693)
reg21695 := PrimCons(reg21689, reg21694)
reg21696 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21697 := PrimCons(reg21695, reg21696)
reg21698 := PrimCons(reg21688, reg21697)
reg21699 := __e.Call(__defun__bind, V2760, reg21698, V2761, V2762)
Result := reg21699
_ = Result
reg21700 := __e.Call(__defun__shen_4unbindv, V2539, V2761)
_ = reg21700
reg21702 = Result
} else {
reg21701 := False;
reg21702 = reg21701
}
reg21703 = reg21702
}
reg21761 = reg21703
} else {
reg21704 := __e.Call(__defun__shen_4pvar_2, V2538)
var reg21760 Obj
if reg21704 == True {
reg21705 := Nil;
reg21706 := __e.Call(__defun__shen_4bindv, V2538, reg21705, V2761)
_ = reg21706
reg21707 := PrimTail(V2528)
reg21708 := __e.Call(__defun__shen_4lazyderef, reg21707, V2761)
V2540 := reg21708
_ = V2540
reg21709 := Nil;
reg21710 := PrimEqual(reg21709, V2540)
var reg21757 Obj
if reg21710 == True {
reg21711 := PrimTail(V2519)
Hyp := reg21711
_ = Hyp
reg21712 := __e.Call(__defun__shen_4incinfs)
_ = reg21712
reg21713 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21714 := MakeSymbol(":")
reg21715 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21716 := Nil;
reg21717 := PrimCons(reg21715, reg21716)
reg21718 := PrimCons(reg21714, reg21717)
reg21719 := PrimCons(reg21713, reg21718)
reg21720 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21721 := MakeSymbol(":")
reg21722 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21723 := Nil;
reg21724 := PrimCons(reg21722, reg21723)
reg21725 := PrimCons(reg21721, reg21724)
reg21726 := PrimCons(reg21720, reg21725)
reg21727 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21728 := PrimCons(reg21726, reg21727)
reg21729 := PrimCons(reg21719, reg21728)
reg21730 := __e.Call(__defun__bind, V2760, reg21729, V2761, V2762)
reg21757 = reg21730
} else {
reg21731 := __e.Call(__defun__shen_4pvar_2, V2540)
var reg21756 Obj
if reg21731 == True {
reg21732 := Nil;
reg21733 := __e.Call(__defun__shen_4bindv, V2540, reg21732, V2761)
_ = reg21733
reg21734 := PrimTail(V2519)
Hyp := reg21734
_ = Hyp
reg21735 := __e.Call(__defun__shen_4incinfs)
_ = reg21735
reg21736 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21737 := MakeSymbol(":")
reg21738 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21739 := Nil;
reg21740 := PrimCons(reg21738, reg21739)
reg21741 := PrimCons(reg21737, reg21740)
reg21742 := PrimCons(reg21736, reg21741)
reg21743 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21744 := MakeSymbol(":")
reg21745 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21746 := Nil;
reg21747 := PrimCons(reg21745, reg21746)
reg21748 := PrimCons(reg21744, reg21747)
reg21749 := PrimCons(reg21743, reg21748)
reg21750 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21751 := PrimCons(reg21749, reg21750)
reg21752 := PrimCons(reg21742, reg21751)
reg21753 := __e.Call(__defun__bind, V2760, reg21752, V2761, V2762)
Result := reg21753
_ = Result
reg21754 := __e.Call(__defun__shen_4unbindv, V2540, V2761)
_ = reg21754
reg21756 = Result
} else {
reg21755 := False;
reg21756 = reg21755
}
reg21757 = reg21756
}
Result := reg21757
_ = Result
reg21758 := __e.Call(__defun__shen_4unbindv, V2538, V2761)
_ = reg21758
reg21760 = Result
} else {
reg21759 := False;
reg21760 = reg21759
}
reg21761 = reg21760
}
reg21821 = reg21761
} else {
reg21762 := __e.Call(__defun__shen_4pvar_2, V2537)
var reg21820 Obj
if reg21762 == True {
reg21763 := __e.Call(__defun__shen_4newpv, V2761)
B := reg21763
_ = B
reg21764 := Nil;
reg21765 := PrimCons(B, reg21764)
reg21766 := __e.Call(__defun__shen_4bindv, V2537, reg21765, V2761)
_ = reg21766
reg21767 := PrimTail(V2528)
reg21768 := __e.Call(__defun__shen_4lazyderef, reg21767, V2761)
V2541 := reg21768
_ = V2541
reg21769 := Nil;
reg21770 := PrimEqual(reg21769, V2541)
var reg21817 Obj
if reg21770 == True {
reg21771 := PrimTail(V2519)
Hyp := reg21771
_ = Hyp
reg21772 := __e.Call(__defun__shen_4incinfs)
_ = reg21772
reg21773 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21774 := MakeSymbol(":")
reg21775 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21776 := Nil;
reg21777 := PrimCons(reg21775, reg21776)
reg21778 := PrimCons(reg21774, reg21777)
reg21779 := PrimCons(reg21773, reg21778)
reg21780 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21781 := MakeSymbol(":")
reg21782 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21783 := Nil;
reg21784 := PrimCons(reg21782, reg21783)
reg21785 := PrimCons(reg21781, reg21784)
reg21786 := PrimCons(reg21780, reg21785)
reg21787 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21788 := PrimCons(reg21786, reg21787)
reg21789 := PrimCons(reg21779, reg21788)
reg21790 := __e.Call(__defun__bind, V2760, reg21789, V2761, V2762)
reg21817 = reg21790
} else {
reg21791 := __e.Call(__defun__shen_4pvar_2, V2541)
var reg21816 Obj
if reg21791 == True {
reg21792 := Nil;
reg21793 := __e.Call(__defun__shen_4bindv, V2541, reg21792, V2761)
_ = reg21793
reg21794 := PrimTail(V2519)
Hyp := reg21794
_ = Hyp
reg21795 := __e.Call(__defun__shen_4incinfs)
_ = reg21795
reg21796 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21797 := MakeSymbol(":")
reg21798 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21799 := Nil;
reg21800 := PrimCons(reg21798, reg21799)
reg21801 := PrimCons(reg21797, reg21800)
reg21802 := PrimCons(reg21796, reg21801)
reg21803 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21804 := MakeSymbol(":")
reg21805 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21806 := Nil;
reg21807 := PrimCons(reg21805, reg21806)
reg21808 := PrimCons(reg21804, reg21807)
reg21809 := PrimCons(reg21803, reg21808)
reg21810 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21811 := PrimCons(reg21809, reg21810)
reg21812 := PrimCons(reg21802, reg21811)
reg21813 := __e.Call(__defun__bind, V2760, reg21812, V2761, V2762)
Result := reg21813
_ = Result
reg21814 := __e.Call(__defun__shen_4unbindv, V2541, V2761)
_ = reg21814
reg21816 = Result
} else {
reg21815 := False;
reg21816 = reg21815
}
reg21817 = reg21816
}
Result := reg21817
_ = Result
reg21818 := __e.Call(__defun__shen_4unbindv, V2537, V2761)
_ = reg21818
reg21820 = Result
} else {
reg21819 := False;
reg21820 = reg21819
}
reg21821 = reg21820
}
Result := reg21821
_ = Result
reg21822 := __e.Call(__defun__shen_4unbindv, V2531, V2761)
_ = reg21822
reg21824 = Result
} else {
reg21823 := False;
reg21824 = reg21823
}
reg21825 = reg21824
}
reg21887 = reg21825
} else {
reg21826 := __e.Call(__defun__shen_4pvar_2, V2530)
var reg21886 Obj
if reg21826 == True {
reg21827 := __e.Call(__defun__shen_4newpv, V2761)
B := reg21827
_ = B
reg21828 := MakeSymbol("*")
reg21829 := Nil;
reg21830 := PrimCons(B, reg21829)
reg21831 := PrimCons(reg21828, reg21830)
reg21832 := __e.Call(__defun__shen_4bindv, V2530, reg21831, V2761)
_ = reg21832
reg21833 := PrimTail(V2528)
reg21834 := __e.Call(__defun__shen_4lazyderef, reg21833, V2761)
V2542 := reg21834
_ = V2542
reg21835 := Nil;
reg21836 := PrimEqual(reg21835, V2542)
var reg21883 Obj
if reg21836 == True {
reg21837 := PrimTail(V2519)
Hyp := reg21837
_ = Hyp
reg21838 := __e.Call(__defun__shen_4incinfs)
_ = reg21838
reg21839 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21840 := MakeSymbol(":")
reg21841 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21842 := Nil;
reg21843 := PrimCons(reg21841, reg21842)
reg21844 := PrimCons(reg21840, reg21843)
reg21845 := PrimCons(reg21839, reg21844)
reg21846 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21847 := MakeSymbol(":")
reg21848 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21849 := Nil;
reg21850 := PrimCons(reg21848, reg21849)
reg21851 := PrimCons(reg21847, reg21850)
reg21852 := PrimCons(reg21846, reg21851)
reg21853 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21854 := PrimCons(reg21852, reg21853)
reg21855 := PrimCons(reg21845, reg21854)
reg21856 := __e.Call(__defun__bind, V2760, reg21855, V2761, V2762)
reg21883 = reg21856
} else {
reg21857 := __e.Call(__defun__shen_4pvar_2, V2542)
var reg21882 Obj
if reg21857 == True {
reg21858 := Nil;
reg21859 := __e.Call(__defun__shen_4bindv, V2542, reg21858, V2761)
_ = reg21859
reg21860 := PrimTail(V2519)
Hyp := reg21860
_ = Hyp
reg21861 := __e.Call(__defun__shen_4incinfs)
_ = reg21861
reg21862 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21863 := MakeSymbol(":")
reg21864 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21865 := Nil;
reg21866 := PrimCons(reg21864, reg21865)
reg21867 := PrimCons(reg21863, reg21866)
reg21868 := PrimCons(reg21862, reg21867)
reg21869 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21870 := MakeSymbol(":")
reg21871 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21872 := Nil;
reg21873 := PrimCons(reg21871, reg21872)
reg21874 := PrimCons(reg21870, reg21873)
reg21875 := PrimCons(reg21869, reg21874)
reg21876 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21877 := PrimCons(reg21875, reg21876)
reg21878 := PrimCons(reg21868, reg21877)
reg21879 := __e.Call(__defun__bind, V2760, reg21878, V2761, V2762)
Result := reg21879
_ = Result
reg21880 := __e.Call(__defun__shen_4unbindv, V2542, V2761)
_ = reg21880
reg21882 = Result
} else {
reg21881 := False;
reg21882 = reg21881
}
reg21883 = reg21882
}
Result := reg21883
_ = Result
reg21884 := __e.Call(__defun__shen_4unbindv, V2530, V2761)
_ = reg21884
reg21886 = Result
} else {
reg21885 := False;
reg21886 = reg21885
}
reg21887 = reg21886
}
reg21951 = reg21887
} else {
reg21888 := __e.Call(__defun__shen_4pvar_2, V2529)
var reg21950 Obj
if reg21888 == True {
reg21889 := __e.Call(__defun__shen_4newpv, V2761)
A := reg21889
_ = A
reg21890 := __e.Call(__defun__shen_4newpv, V2761)
B := reg21890
_ = B
reg21891 := MakeSymbol("*")
reg21892 := Nil;
reg21893 := PrimCons(B, reg21892)
reg21894 := PrimCons(reg21891, reg21893)
reg21895 := PrimCons(A, reg21894)
reg21896 := __e.Call(__defun__shen_4bindv, V2529, reg21895, V2761)
_ = reg21896
reg21897 := PrimTail(V2528)
reg21898 := __e.Call(__defun__shen_4lazyderef, reg21897, V2761)
V2543 := reg21898
_ = V2543
reg21899 := Nil;
reg21900 := PrimEqual(reg21899, V2543)
var reg21947 Obj
if reg21900 == True {
reg21901 := PrimTail(V2519)
Hyp := reg21901
_ = Hyp
reg21902 := __e.Call(__defun__shen_4incinfs)
_ = reg21902
reg21903 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21904 := MakeSymbol(":")
reg21905 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21906 := Nil;
reg21907 := PrimCons(reg21905, reg21906)
reg21908 := PrimCons(reg21904, reg21907)
reg21909 := PrimCons(reg21903, reg21908)
reg21910 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21911 := MakeSymbol(":")
reg21912 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21913 := Nil;
reg21914 := PrimCons(reg21912, reg21913)
reg21915 := PrimCons(reg21911, reg21914)
reg21916 := PrimCons(reg21910, reg21915)
reg21917 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21918 := PrimCons(reg21916, reg21917)
reg21919 := PrimCons(reg21909, reg21918)
reg21920 := __e.Call(__defun__bind, V2760, reg21919, V2761, V2762)
reg21947 = reg21920
} else {
reg21921 := __e.Call(__defun__shen_4pvar_2, V2543)
var reg21946 Obj
if reg21921 == True {
reg21922 := Nil;
reg21923 := __e.Call(__defun__shen_4bindv, V2543, reg21922, V2761)
_ = reg21923
reg21924 := PrimTail(V2519)
Hyp := reg21924
_ = Hyp
reg21925 := __e.Call(__defun__shen_4incinfs)
_ = reg21925
reg21926 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg21927 := MakeSymbol(":")
reg21928 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg21929 := Nil;
reg21930 := PrimCons(reg21928, reg21929)
reg21931 := PrimCons(reg21927, reg21930)
reg21932 := PrimCons(reg21926, reg21931)
reg21933 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg21934 := MakeSymbol(":")
reg21935 := __e.Call(__defun__shen_4lazyderef, B, V2761)
reg21936 := Nil;
reg21937 := PrimCons(reg21935, reg21936)
reg21938 := PrimCons(reg21934, reg21937)
reg21939 := PrimCons(reg21933, reg21938)
reg21940 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg21941 := PrimCons(reg21939, reg21940)
reg21942 := PrimCons(reg21932, reg21941)
reg21943 := __e.Call(__defun__bind, V2760, reg21942, V2761, V2762)
Result := reg21943
_ = Result
reg21944 := __e.Call(__defun__shen_4unbindv, V2543, V2761)
_ = reg21944
reg21946 = Result
} else {
reg21945 := False;
reg21946 = reg21945
}
reg21947 = reg21946
}
Result := reg21947
_ = Result
reg21948 := __e.Call(__defun__shen_4unbindv, V2529, V2761)
_ = reg21948
reg21950 = Result
} else {
reg21949 := False;
reg21950 = reg21949
}
reg21951 = reg21950
}
reg21953 = reg21951
} else {
reg21952 := False;
reg21953 = reg21952
}
reg21955 = reg21953
} else {
reg21954 := False;
reg21955 = reg21954
}
reg21957 = reg21955
} else {
reg21956 := False;
reg21957 = reg21956
}
reg21959 = reg21957
} else {
reg21958 := False;
reg21959 = reg21958
}
reg21961 = reg21959
} else {
reg21960 := False;
reg21961 = reg21960
}
reg21963 = reg21961
} else {
reg21962 := False;
reg21963 = reg21962
}
reg21965 = reg21963
} else {
reg21964 := False;
reg21965 = reg21964
}
reg21967 = reg21965
} else {
reg21966 := False;
reg21967 = reg21966
}
reg21969 = reg21967
} else {
reg21968 := False;
reg21969 = reg21968
}
reg21971 = reg21969
} else {
reg21970 := False;
reg21971 = reg21970
}
Case := reg21971
_ = Case
reg21972 := False;
reg21973 := PrimEqual(Case, reg21972)
if reg21973 == True {
reg21974 := __e.Call(__defun__shen_4lazyderef, V2759, V2761)
V2544 := reg21974
_ = V2544
reg21975 := PrimIsPair(V2544)
var reg22513 Obj
if reg21975 == True {
reg21976 := PrimHead(V2544)
reg21977 := __e.Call(__defun__shen_4lazyderef, reg21976, V2761)
V2545 := reg21977
_ = V2545
reg21978 := PrimIsPair(V2545)
var reg22511 Obj
if reg21978 == True {
reg21979 := PrimHead(V2545)
reg21980 := __e.Call(__defun__shen_4lazyderef, reg21979, V2761)
V2546 := reg21980
_ = V2546
reg21981 := PrimIsPair(V2546)
var reg22509 Obj
if reg21981 == True {
reg21982 := PrimHead(V2546)
reg21983 := __e.Call(__defun__shen_4lazyderef, reg21982, V2761)
V2547 := reg21983
_ = V2547
reg21984 := MakeSymbol("@v")
reg21985 := PrimEqual(reg21984, V2547)
var reg22507 Obj
if reg21985 == True {
reg21986 := PrimTail(V2546)
reg21987 := __e.Call(__defun__shen_4lazyderef, reg21986, V2761)
V2548 := reg21987
_ = V2548
reg21988 := PrimIsPair(V2548)
var reg22505 Obj
if reg21988 == True {
reg21989 := PrimHead(V2548)
X := reg21989
_ = X
reg21990 := PrimTail(V2548)
reg21991 := __e.Call(__defun__shen_4lazyderef, reg21990, V2761)
V2549 := reg21991
_ = V2549
reg21992 := PrimIsPair(V2549)
var reg22503 Obj
if reg21992 == True {
reg21993 := PrimHead(V2549)
Y := reg21993
_ = Y
reg21994 := PrimTail(V2549)
reg21995 := __e.Call(__defun__shen_4lazyderef, reg21994, V2761)
V2550 := reg21995
_ = V2550
reg21996 := Nil;
reg21997 := PrimEqual(reg21996, V2550)
var reg22501 Obj
if reg21997 == True {
reg21998 := PrimTail(V2545)
reg21999 := __e.Call(__defun__shen_4lazyderef, reg21998, V2761)
V2551 := reg21999
_ = V2551
reg22000 := PrimIsPair(V2551)
var reg22499 Obj
if reg22000 == True {
reg22001 := PrimHead(V2551)
reg22002 := __e.Call(__defun__shen_4lazyderef, reg22001, V2761)
V2552 := reg22002
_ = V2552
reg22003 := MakeSymbol(":")
reg22004 := PrimEqual(reg22003, V2552)
var reg22497 Obj
if reg22004 == True {
reg22005 := PrimTail(V2551)
reg22006 := __e.Call(__defun__shen_4lazyderef, reg22005, V2761)
V2553 := reg22006
_ = V2553
reg22007 := PrimIsPair(V2553)
var reg22495 Obj
if reg22007 == True {
reg22008 := PrimHead(V2553)
reg22009 := __e.Call(__defun__shen_4lazyderef, reg22008, V2761)
V2554 := reg22009
_ = V2554
reg22010 := PrimIsPair(V2554)
var reg22493 Obj
if reg22010 == True {
reg22011 := PrimHead(V2554)
reg22012 := __e.Call(__defun__shen_4lazyderef, reg22011, V2761)
V2555 := reg22012
_ = V2555
reg22013 := MakeSymbol("vector")
reg22014 := PrimEqual(reg22013, V2555)
var reg22423 Obj
if reg22014 == True {
reg22015 := PrimTail(V2554)
reg22016 := __e.Call(__defun__shen_4lazyderef, reg22015, V2761)
V2556 := reg22016
_ = V2556
reg22017 := PrimIsPair(V2556)
var reg22215 Obj
if reg22017 == True {
reg22018 := PrimHead(V2556)
A := reg22018
_ = A
reg22019 := PrimTail(V2556)
reg22020 := __e.Call(__defun__shen_4lazyderef, reg22019, V2761)
V2557 := reg22020
_ = V2557
reg22021 := Nil;
reg22022 := PrimEqual(reg22021, V2557)
var reg22147 Obj
if reg22022 == True {
reg22023 := PrimTail(V2553)
reg22024 := __e.Call(__defun__shen_4lazyderef, reg22023, V2761)
V2558 := reg22024
_ = V2558
reg22025 := Nil;
reg22026 := PrimEqual(reg22025, V2558)
var reg22081 Obj
if reg22026 == True {
reg22027 := PrimTail(V2544)
Hyp := reg22027
_ = Hyp
reg22028 := __e.Call(__defun__shen_4incinfs)
_ = reg22028
reg22029 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22030 := MakeSymbol(":")
reg22031 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22032 := Nil;
reg22033 := PrimCons(reg22031, reg22032)
reg22034 := PrimCons(reg22030, reg22033)
reg22035 := PrimCons(reg22029, reg22034)
reg22036 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22037 := MakeSymbol(":")
reg22038 := MakeSymbol("vector")
reg22039 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22040 := Nil;
reg22041 := PrimCons(reg22039, reg22040)
reg22042 := PrimCons(reg22038, reg22041)
reg22043 := Nil;
reg22044 := PrimCons(reg22042, reg22043)
reg22045 := PrimCons(reg22037, reg22044)
reg22046 := PrimCons(reg22036, reg22045)
reg22047 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22048 := PrimCons(reg22046, reg22047)
reg22049 := PrimCons(reg22035, reg22048)
reg22050 := __e.Call(__defun__bind, V2760, reg22049, V2761, V2762)
reg22081 = reg22050
} else {
reg22051 := __e.Call(__defun__shen_4pvar_2, V2558)
var reg22080 Obj
if reg22051 == True {
reg22052 := Nil;
reg22053 := __e.Call(__defun__shen_4bindv, V2558, reg22052, V2761)
_ = reg22053
reg22054 := PrimTail(V2544)
Hyp := reg22054
_ = Hyp
reg22055 := __e.Call(__defun__shen_4incinfs)
_ = reg22055
reg22056 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22057 := MakeSymbol(":")
reg22058 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22059 := Nil;
reg22060 := PrimCons(reg22058, reg22059)
reg22061 := PrimCons(reg22057, reg22060)
reg22062 := PrimCons(reg22056, reg22061)
reg22063 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22064 := MakeSymbol(":")
reg22065 := MakeSymbol("vector")
reg22066 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22067 := Nil;
reg22068 := PrimCons(reg22066, reg22067)
reg22069 := PrimCons(reg22065, reg22068)
reg22070 := Nil;
reg22071 := PrimCons(reg22069, reg22070)
reg22072 := PrimCons(reg22064, reg22071)
reg22073 := PrimCons(reg22063, reg22072)
reg22074 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22075 := PrimCons(reg22073, reg22074)
reg22076 := PrimCons(reg22062, reg22075)
reg22077 := __e.Call(__defun__bind, V2760, reg22076, V2761, V2762)
Result := reg22077
_ = Result
reg22078 := __e.Call(__defun__shen_4unbindv, V2558, V2761)
_ = reg22078
reg22080 = Result
} else {
reg22079 := False;
reg22080 = reg22079
}
reg22081 = reg22080
}
reg22147 = reg22081
} else {
reg22082 := __e.Call(__defun__shen_4pvar_2, V2557)
var reg22146 Obj
if reg22082 == True {
reg22083 := Nil;
reg22084 := __e.Call(__defun__shen_4bindv, V2557, reg22083, V2761)
_ = reg22084
reg22085 := PrimTail(V2553)
reg22086 := __e.Call(__defun__shen_4lazyderef, reg22085, V2761)
V2559 := reg22086
_ = V2559
reg22087 := Nil;
reg22088 := PrimEqual(reg22087, V2559)
var reg22143 Obj
if reg22088 == True {
reg22089 := PrimTail(V2544)
Hyp := reg22089
_ = Hyp
reg22090 := __e.Call(__defun__shen_4incinfs)
_ = reg22090
reg22091 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22092 := MakeSymbol(":")
reg22093 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22094 := Nil;
reg22095 := PrimCons(reg22093, reg22094)
reg22096 := PrimCons(reg22092, reg22095)
reg22097 := PrimCons(reg22091, reg22096)
reg22098 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22099 := MakeSymbol(":")
reg22100 := MakeSymbol("vector")
reg22101 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22102 := Nil;
reg22103 := PrimCons(reg22101, reg22102)
reg22104 := PrimCons(reg22100, reg22103)
reg22105 := Nil;
reg22106 := PrimCons(reg22104, reg22105)
reg22107 := PrimCons(reg22099, reg22106)
reg22108 := PrimCons(reg22098, reg22107)
reg22109 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22110 := PrimCons(reg22108, reg22109)
reg22111 := PrimCons(reg22097, reg22110)
reg22112 := __e.Call(__defun__bind, V2760, reg22111, V2761, V2762)
reg22143 = reg22112
} else {
reg22113 := __e.Call(__defun__shen_4pvar_2, V2559)
var reg22142 Obj
if reg22113 == True {
reg22114 := Nil;
reg22115 := __e.Call(__defun__shen_4bindv, V2559, reg22114, V2761)
_ = reg22115
reg22116 := PrimTail(V2544)
Hyp := reg22116
_ = Hyp
reg22117 := __e.Call(__defun__shen_4incinfs)
_ = reg22117
reg22118 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22119 := MakeSymbol(":")
reg22120 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22121 := Nil;
reg22122 := PrimCons(reg22120, reg22121)
reg22123 := PrimCons(reg22119, reg22122)
reg22124 := PrimCons(reg22118, reg22123)
reg22125 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22126 := MakeSymbol(":")
reg22127 := MakeSymbol("vector")
reg22128 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22129 := Nil;
reg22130 := PrimCons(reg22128, reg22129)
reg22131 := PrimCons(reg22127, reg22130)
reg22132 := Nil;
reg22133 := PrimCons(reg22131, reg22132)
reg22134 := PrimCons(reg22126, reg22133)
reg22135 := PrimCons(reg22125, reg22134)
reg22136 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22137 := PrimCons(reg22135, reg22136)
reg22138 := PrimCons(reg22124, reg22137)
reg22139 := __e.Call(__defun__bind, V2760, reg22138, V2761, V2762)
Result := reg22139
_ = Result
reg22140 := __e.Call(__defun__shen_4unbindv, V2559, V2761)
_ = reg22140
reg22142 = Result
} else {
reg22141 := False;
reg22142 = reg22141
}
reg22143 = reg22142
}
Result := reg22143
_ = Result
reg22144 := __e.Call(__defun__shen_4unbindv, V2557, V2761)
_ = reg22144
reg22146 = Result
} else {
reg22145 := False;
reg22146 = reg22145
}
reg22147 = reg22146
}
reg22215 = reg22147
} else {
reg22148 := __e.Call(__defun__shen_4pvar_2, V2556)
var reg22214 Obj
if reg22148 == True {
reg22149 := __e.Call(__defun__shen_4newpv, V2761)
A := reg22149
_ = A
reg22150 := Nil;
reg22151 := PrimCons(A, reg22150)
reg22152 := __e.Call(__defun__shen_4bindv, V2556, reg22151, V2761)
_ = reg22152
reg22153 := PrimTail(V2553)
reg22154 := __e.Call(__defun__shen_4lazyderef, reg22153, V2761)
V2560 := reg22154
_ = V2560
reg22155 := Nil;
reg22156 := PrimEqual(reg22155, V2560)
var reg22211 Obj
if reg22156 == True {
reg22157 := PrimTail(V2544)
Hyp := reg22157
_ = Hyp
reg22158 := __e.Call(__defun__shen_4incinfs)
_ = reg22158
reg22159 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22160 := MakeSymbol(":")
reg22161 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22162 := Nil;
reg22163 := PrimCons(reg22161, reg22162)
reg22164 := PrimCons(reg22160, reg22163)
reg22165 := PrimCons(reg22159, reg22164)
reg22166 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22167 := MakeSymbol(":")
reg22168 := MakeSymbol("vector")
reg22169 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22170 := Nil;
reg22171 := PrimCons(reg22169, reg22170)
reg22172 := PrimCons(reg22168, reg22171)
reg22173 := Nil;
reg22174 := PrimCons(reg22172, reg22173)
reg22175 := PrimCons(reg22167, reg22174)
reg22176 := PrimCons(reg22166, reg22175)
reg22177 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22178 := PrimCons(reg22176, reg22177)
reg22179 := PrimCons(reg22165, reg22178)
reg22180 := __e.Call(__defun__bind, V2760, reg22179, V2761, V2762)
reg22211 = reg22180
} else {
reg22181 := __e.Call(__defun__shen_4pvar_2, V2560)
var reg22210 Obj
if reg22181 == True {
reg22182 := Nil;
reg22183 := __e.Call(__defun__shen_4bindv, V2560, reg22182, V2761)
_ = reg22183
reg22184 := PrimTail(V2544)
Hyp := reg22184
_ = Hyp
reg22185 := __e.Call(__defun__shen_4incinfs)
_ = reg22185
reg22186 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22187 := MakeSymbol(":")
reg22188 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22189 := Nil;
reg22190 := PrimCons(reg22188, reg22189)
reg22191 := PrimCons(reg22187, reg22190)
reg22192 := PrimCons(reg22186, reg22191)
reg22193 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22194 := MakeSymbol(":")
reg22195 := MakeSymbol("vector")
reg22196 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22197 := Nil;
reg22198 := PrimCons(reg22196, reg22197)
reg22199 := PrimCons(reg22195, reg22198)
reg22200 := Nil;
reg22201 := PrimCons(reg22199, reg22200)
reg22202 := PrimCons(reg22194, reg22201)
reg22203 := PrimCons(reg22193, reg22202)
reg22204 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22205 := PrimCons(reg22203, reg22204)
reg22206 := PrimCons(reg22192, reg22205)
reg22207 := __e.Call(__defun__bind, V2760, reg22206, V2761, V2762)
Result := reg22207
_ = Result
reg22208 := __e.Call(__defun__shen_4unbindv, V2560, V2761)
_ = reg22208
reg22210 = Result
} else {
reg22209 := False;
reg22210 = reg22209
}
reg22211 = reg22210
}
Result := reg22211
_ = Result
reg22212 := __e.Call(__defun__shen_4unbindv, V2556, V2761)
_ = reg22212
reg22214 = Result
} else {
reg22213 := False;
reg22214 = reg22213
}
reg22215 = reg22214
}
reg22423 = reg22215
} else {
reg22216 := __e.Call(__defun__shen_4pvar_2, V2555)
var reg22422 Obj
if reg22216 == True {
reg22217 := MakeSymbol("vector")
reg22218 := __e.Call(__defun__shen_4bindv, V2555, reg22217, V2761)
_ = reg22218
reg22219 := PrimTail(V2554)
reg22220 := __e.Call(__defun__shen_4lazyderef, reg22219, V2761)
V2561 := reg22220
_ = V2561
reg22221 := PrimIsPair(V2561)
var reg22419 Obj
if reg22221 == True {
reg22222 := PrimHead(V2561)
A := reg22222
_ = A
reg22223 := PrimTail(V2561)
reg22224 := __e.Call(__defun__shen_4lazyderef, reg22223, V2761)
V2562 := reg22224
_ = V2562
reg22225 := Nil;
reg22226 := PrimEqual(reg22225, V2562)
var reg22351 Obj
if reg22226 == True {
reg22227 := PrimTail(V2553)
reg22228 := __e.Call(__defun__shen_4lazyderef, reg22227, V2761)
V2563 := reg22228
_ = V2563
reg22229 := Nil;
reg22230 := PrimEqual(reg22229, V2563)
var reg22285 Obj
if reg22230 == True {
reg22231 := PrimTail(V2544)
Hyp := reg22231
_ = Hyp
reg22232 := __e.Call(__defun__shen_4incinfs)
_ = reg22232
reg22233 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22234 := MakeSymbol(":")
reg22235 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22236 := Nil;
reg22237 := PrimCons(reg22235, reg22236)
reg22238 := PrimCons(reg22234, reg22237)
reg22239 := PrimCons(reg22233, reg22238)
reg22240 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22241 := MakeSymbol(":")
reg22242 := MakeSymbol("vector")
reg22243 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22244 := Nil;
reg22245 := PrimCons(reg22243, reg22244)
reg22246 := PrimCons(reg22242, reg22245)
reg22247 := Nil;
reg22248 := PrimCons(reg22246, reg22247)
reg22249 := PrimCons(reg22241, reg22248)
reg22250 := PrimCons(reg22240, reg22249)
reg22251 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22252 := PrimCons(reg22250, reg22251)
reg22253 := PrimCons(reg22239, reg22252)
reg22254 := __e.Call(__defun__bind, V2760, reg22253, V2761, V2762)
reg22285 = reg22254
} else {
reg22255 := __e.Call(__defun__shen_4pvar_2, V2563)
var reg22284 Obj
if reg22255 == True {
reg22256 := Nil;
reg22257 := __e.Call(__defun__shen_4bindv, V2563, reg22256, V2761)
_ = reg22257
reg22258 := PrimTail(V2544)
Hyp := reg22258
_ = Hyp
reg22259 := __e.Call(__defun__shen_4incinfs)
_ = reg22259
reg22260 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22261 := MakeSymbol(":")
reg22262 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22263 := Nil;
reg22264 := PrimCons(reg22262, reg22263)
reg22265 := PrimCons(reg22261, reg22264)
reg22266 := PrimCons(reg22260, reg22265)
reg22267 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22268 := MakeSymbol(":")
reg22269 := MakeSymbol("vector")
reg22270 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22271 := Nil;
reg22272 := PrimCons(reg22270, reg22271)
reg22273 := PrimCons(reg22269, reg22272)
reg22274 := Nil;
reg22275 := PrimCons(reg22273, reg22274)
reg22276 := PrimCons(reg22268, reg22275)
reg22277 := PrimCons(reg22267, reg22276)
reg22278 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22279 := PrimCons(reg22277, reg22278)
reg22280 := PrimCons(reg22266, reg22279)
reg22281 := __e.Call(__defun__bind, V2760, reg22280, V2761, V2762)
Result := reg22281
_ = Result
reg22282 := __e.Call(__defun__shen_4unbindv, V2563, V2761)
_ = reg22282
reg22284 = Result
} else {
reg22283 := False;
reg22284 = reg22283
}
reg22285 = reg22284
}
reg22351 = reg22285
} else {
reg22286 := __e.Call(__defun__shen_4pvar_2, V2562)
var reg22350 Obj
if reg22286 == True {
reg22287 := Nil;
reg22288 := __e.Call(__defun__shen_4bindv, V2562, reg22287, V2761)
_ = reg22288
reg22289 := PrimTail(V2553)
reg22290 := __e.Call(__defun__shen_4lazyderef, reg22289, V2761)
V2564 := reg22290
_ = V2564
reg22291 := Nil;
reg22292 := PrimEqual(reg22291, V2564)
var reg22347 Obj
if reg22292 == True {
reg22293 := PrimTail(V2544)
Hyp := reg22293
_ = Hyp
reg22294 := __e.Call(__defun__shen_4incinfs)
_ = reg22294
reg22295 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22296 := MakeSymbol(":")
reg22297 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22298 := Nil;
reg22299 := PrimCons(reg22297, reg22298)
reg22300 := PrimCons(reg22296, reg22299)
reg22301 := PrimCons(reg22295, reg22300)
reg22302 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22303 := MakeSymbol(":")
reg22304 := MakeSymbol("vector")
reg22305 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22306 := Nil;
reg22307 := PrimCons(reg22305, reg22306)
reg22308 := PrimCons(reg22304, reg22307)
reg22309 := Nil;
reg22310 := PrimCons(reg22308, reg22309)
reg22311 := PrimCons(reg22303, reg22310)
reg22312 := PrimCons(reg22302, reg22311)
reg22313 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22314 := PrimCons(reg22312, reg22313)
reg22315 := PrimCons(reg22301, reg22314)
reg22316 := __e.Call(__defun__bind, V2760, reg22315, V2761, V2762)
reg22347 = reg22316
} else {
reg22317 := __e.Call(__defun__shen_4pvar_2, V2564)
var reg22346 Obj
if reg22317 == True {
reg22318 := Nil;
reg22319 := __e.Call(__defun__shen_4bindv, V2564, reg22318, V2761)
_ = reg22319
reg22320 := PrimTail(V2544)
Hyp := reg22320
_ = Hyp
reg22321 := __e.Call(__defun__shen_4incinfs)
_ = reg22321
reg22322 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22323 := MakeSymbol(":")
reg22324 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22325 := Nil;
reg22326 := PrimCons(reg22324, reg22325)
reg22327 := PrimCons(reg22323, reg22326)
reg22328 := PrimCons(reg22322, reg22327)
reg22329 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22330 := MakeSymbol(":")
reg22331 := MakeSymbol("vector")
reg22332 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22333 := Nil;
reg22334 := PrimCons(reg22332, reg22333)
reg22335 := PrimCons(reg22331, reg22334)
reg22336 := Nil;
reg22337 := PrimCons(reg22335, reg22336)
reg22338 := PrimCons(reg22330, reg22337)
reg22339 := PrimCons(reg22329, reg22338)
reg22340 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22341 := PrimCons(reg22339, reg22340)
reg22342 := PrimCons(reg22328, reg22341)
reg22343 := __e.Call(__defun__bind, V2760, reg22342, V2761, V2762)
Result := reg22343
_ = Result
reg22344 := __e.Call(__defun__shen_4unbindv, V2564, V2761)
_ = reg22344
reg22346 = Result
} else {
reg22345 := False;
reg22346 = reg22345
}
reg22347 = reg22346
}
Result := reg22347
_ = Result
reg22348 := __e.Call(__defun__shen_4unbindv, V2562, V2761)
_ = reg22348
reg22350 = Result
} else {
reg22349 := False;
reg22350 = reg22349
}
reg22351 = reg22350
}
reg22419 = reg22351
} else {
reg22352 := __e.Call(__defun__shen_4pvar_2, V2561)
var reg22418 Obj
if reg22352 == True {
reg22353 := __e.Call(__defun__shen_4newpv, V2761)
A := reg22353
_ = A
reg22354 := Nil;
reg22355 := PrimCons(A, reg22354)
reg22356 := __e.Call(__defun__shen_4bindv, V2561, reg22355, V2761)
_ = reg22356
reg22357 := PrimTail(V2553)
reg22358 := __e.Call(__defun__shen_4lazyderef, reg22357, V2761)
V2565 := reg22358
_ = V2565
reg22359 := Nil;
reg22360 := PrimEqual(reg22359, V2565)
var reg22415 Obj
if reg22360 == True {
reg22361 := PrimTail(V2544)
Hyp := reg22361
_ = Hyp
reg22362 := __e.Call(__defun__shen_4incinfs)
_ = reg22362
reg22363 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22364 := MakeSymbol(":")
reg22365 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22366 := Nil;
reg22367 := PrimCons(reg22365, reg22366)
reg22368 := PrimCons(reg22364, reg22367)
reg22369 := PrimCons(reg22363, reg22368)
reg22370 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22371 := MakeSymbol(":")
reg22372 := MakeSymbol("vector")
reg22373 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22374 := Nil;
reg22375 := PrimCons(reg22373, reg22374)
reg22376 := PrimCons(reg22372, reg22375)
reg22377 := Nil;
reg22378 := PrimCons(reg22376, reg22377)
reg22379 := PrimCons(reg22371, reg22378)
reg22380 := PrimCons(reg22370, reg22379)
reg22381 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22382 := PrimCons(reg22380, reg22381)
reg22383 := PrimCons(reg22369, reg22382)
reg22384 := __e.Call(__defun__bind, V2760, reg22383, V2761, V2762)
reg22415 = reg22384
} else {
reg22385 := __e.Call(__defun__shen_4pvar_2, V2565)
var reg22414 Obj
if reg22385 == True {
reg22386 := Nil;
reg22387 := __e.Call(__defun__shen_4bindv, V2565, reg22386, V2761)
_ = reg22387
reg22388 := PrimTail(V2544)
Hyp := reg22388
_ = Hyp
reg22389 := __e.Call(__defun__shen_4incinfs)
_ = reg22389
reg22390 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22391 := MakeSymbol(":")
reg22392 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22393 := Nil;
reg22394 := PrimCons(reg22392, reg22393)
reg22395 := PrimCons(reg22391, reg22394)
reg22396 := PrimCons(reg22390, reg22395)
reg22397 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22398 := MakeSymbol(":")
reg22399 := MakeSymbol("vector")
reg22400 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22401 := Nil;
reg22402 := PrimCons(reg22400, reg22401)
reg22403 := PrimCons(reg22399, reg22402)
reg22404 := Nil;
reg22405 := PrimCons(reg22403, reg22404)
reg22406 := PrimCons(reg22398, reg22405)
reg22407 := PrimCons(reg22397, reg22406)
reg22408 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22409 := PrimCons(reg22407, reg22408)
reg22410 := PrimCons(reg22396, reg22409)
reg22411 := __e.Call(__defun__bind, V2760, reg22410, V2761, V2762)
Result := reg22411
_ = Result
reg22412 := __e.Call(__defun__shen_4unbindv, V2565, V2761)
_ = reg22412
reg22414 = Result
} else {
reg22413 := False;
reg22414 = reg22413
}
reg22415 = reg22414
}
Result := reg22415
_ = Result
reg22416 := __e.Call(__defun__shen_4unbindv, V2561, V2761)
_ = reg22416
reg22418 = Result
} else {
reg22417 := False;
reg22418 = reg22417
}
reg22419 = reg22418
}
Result := reg22419
_ = Result
reg22420 := __e.Call(__defun__shen_4unbindv, V2555, V2761)
_ = reg22420
reg22422 = Result
} else {
reg22421 := False;
reg22422 = reg22421
}
reg22423 = reg22422
}
reg22493 = reg22423
} else {
reg22424 := __e.Call(__defun__shen_4pvar_2, V2554)
var reg22492 Obj
if reg22424 == True {
reg22425 := __e.Call(__defun__shen_4newpv, V2761)
A := reg22425
_ = A
reg22426 := MakeSymbol("vector")
reg22427 := Nil;
reg22428 := PrimCons(A, reg22427)
reg22429 := PrimCons(reg22426, reg22428)
reg22430 := __e.Call(__defun__shen_4bindv, V2554, reg22429, V2761)
_ = reg22430
reg22431 := PrimTail(V2553)
reg22432 := __e.Call(__defun__shen_4lazyderef, reg22431, V2761)
V2566 := reg22432
_ = V2566
reg22433 := Nil;
reg22434 := PrimEqual(reg22433, V2566)
var reg22489 Obj
if reg22434 == True {
reg22435 := PrimTail(V2544)
Hyp := reg22435
_ = Hyp
reg22436 := __e.Call(__defun__shen_4incinfs)
_ = reg22436
reg22437 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22438 := MakeSymbol(":")
reg22439 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22440 := Nil;
reg22441 := PrimCons(reg22439, reg22440)
reg22442 := PrimCons(reg22438, reg22441)
reg22443 := PrimCons(reg22437, reg22442)
reg22444 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22445 := MakeSymbol(":")
reg22446 := MakeSymbol("vector")
reg22447 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22448 := Nil;
reg22449 := PrimCons(reg22447, reg22448)
reg22450 := PrimCons(reg22446, reg22449)
reg22451 := Nil;
reg22452 := PrimCons(reg22450, reg22451)
reg22453 := PrimCons(reg22445, reg22452)
reg22454 := PrimCons(reg22444, reg22453)
reg22455 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22456 := PrimCons(reg22454, reg22455)
reg22457 := PrimCons(reg22443, reg22456)
reg22458 := __e.Call(__defun__bind, V2760, reg22457, V2761, V2762)
reg22489 = reg22458
} else {
reg22459 := __e.Call(__defun__shen_4pvar_2, V2566)
var reg22488 Obj
if reg22459 == True {
reg22460 := Nil;
reg22461 := __e.Call(__defun__shen_4bindv, V2566, reg22460, V2761)
_ = reg22461
reg22462 := PrimTail(V2544)
Hyp := reg22462
_ = Hyp
reg22463 := __e.Call(__defun__shen_4incinfs)
_ = reg22463
reg22464 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22465 := MakeSymbol(":")
reg22466 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22467 := Nil;
reg22468 := PrimCons(reg22466, reg22467)
reg22469 := PrimCons(reg22465, reg22468)
reg22470 := PrimCons(reg22464, reg22469)
reg22471 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22472 := MakeSymbol(":")
reg22473 := MakeSymbol("vector")
reg22474 := __e.Call(__defun__shen_4lazyderef, A, V2761)
reg22475 := Nil;
reg22476 := PrimCons(reg22474, reg22475)
reg22477 := PrimCons(reg22473, reg22476)
reg22478 := Nil;
reg22479 := PrimCons(reg22477, reg22478)
reg22480 := PrimCons(reg22472, reg22479)
reg22481 := PrimCons(reg22471, reg22480)
reg22482 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22483 := PrimCons(reg22481, reg22482)
reg22484 := PrimCons(reg22470, reg22483)
reg22485 := __e.Call(__defun__bind, V2760, reg22484, V2761, V2762)
Result := reg22485
_ = Result
reg22486 := __e.Call(__defun__shen_4unbindv, V2566, V2761)
_ = reg22486
reg22488 = Result
} else {
reg22487 := False;
reg22488 = reg22487
}
reg22489 = reg22488
}
Result := reg22489
_ = Result
reg22490 := __e.Call(__defun__shen_4unbindv, V2554, V2761)
_ = reg22490
reg22492 = Result
} else {
reg22491 := False;
reg22492 = reg22491
}
reg22493 = reg22492
}
reg22495 = reg22493
} else {
reg22494 := False;
reg22495 = reg22494
}
reg22497 = reg22495
} else {
reg22496 := False;
reg22497 = reg22496
}
reg22499 = reg22497
} else {
reg22498 := False;
reg22499 = reg22498
}
reg22501 = reg22499
} else {
reg22500 := False;
reg22501 = reg22500
}
reg22503 = reg22501
} else {
reg22502 := False;
reg22503 = reg22502
}
reg22505 = reg22503
} else {
reg22504 := False;
reg22505 = reg22504
}
reg22507 = reg22505
} else {
reg22506 := False;
reg22507 = reg22506
}
reg22509 = reg22507
} else {
reg22508 := False;
reg22509 = reg22508
}
reg22511 = reg22509
} else {
reg22510 := False;
reg22511 = reg22510
}
reg22513 = reg22511
} else {
reg22512 := False;
reg22513 = reg22512
}
Case := reg22513
_ = Case
reg22514 := False;
reg22515 := PrimEqual(Case, reg22514)
if reg22515 == True {
reg22516 := __e.Call(__defun__shen_4lazyderef, V2759, V2761)
V2567 := reg22516
_ = V2567
reg22517 := PrimIsPair(V2567)
var reg22682 Obj
if reg22517 == True {
reg22518 := PrimHead(V2567)
reg22519 := __e.Call(__defun__shen_4lazyderef, reg22518, V2761)
V2568 := reg22519
_ = V2568
reg22520 := PrimIsPair(V2568)
var reg22680 Obj
if reg22520 == True {
reg22521 := PrimHead(V2568)
reg22522 := __e.Call(__defun__shen_4lazyderef, reg22521, V2761)
V2569 := reg22522
_ = V2569
reg22523 := PrimIsPair(V2569)
var reg22678 Obj
if reg22523 == True {
reg22524 := PrimHead(V2569)
reg22525 := __e.Call(__defun__shen_4lazyderef, reg22524, V2761)
V2570 := reg22525
_ = V2570
reg22526 := MakeSymbol("@s")
reg22527 := PrimEqual(reg22526, V2570)
var reg22676 Obj
if reg22527 == True {
reg22528 := PrimTail(V2569)
reg22529 := __e.Call(__defun__shen_4lazyderef, reg22528, V2761)
V2571 := reg22529
_ = V2571
reg22530 := PrimIsPair(V2571)
var reg22674 Obj
if reg22530 == True {
reg22531 := PrimHead(V2571)
X := reg22531
_ = X
reg22532 := PrimTail(V2571)
reg22533 := __e.Call(__defun__shen_4lazyderef, reg22532, V2761)
V2572 := reg22533
_ = V2572
reg22534 := PrimIsPair(V2572)
var reg22672 Obj
if reg22534 == True {
reg22535 := PrimHead(V2572)
Y := reg22535
_ = Y
reg22536 := PrimTail(V2572)
reg22537 := __e.Call(__defun__shen_4lazyderef, reg22536, V2761)
V2573 := reg22537
_ = V2573
reg22538 := Nil;
reg22539 := PrimEqual(reg22538, V2573)
var reg22670 Obj
if reg22539 == True {
reg22540 := PrimTail(V2568)
reg22541 := __e.Call(__defun__shen_4lazyderef, reg22540, V2761)
V2574 := reg22541
_ = V2574
reg22542 := PrimIsPair(V2574)
var reg22668 Obj
if reg22542 == True {
reg22543 := PrimHead(V2574)
reg22544 := __e.Call(__defun__shen_4lazyderef, reg22543, V2761)
V2575 := reg22544
_ = V2575
reg22545 := MakeSymbol(":")
reg22546 := PrimEqual(reg22545, V2575)
var reg22666 Obj
if reg22546 == True {
reg22547 := PrimTail(V2574)
reg22548 := __e.Call(__defun__shen_4lazyderef, reg22547, V2761)
V2576 := reg22548
_ = V2576
reg22549 := PrimIsPair(V2576)
var reg22664 Obj
if reg22549 == True {
reg22550 := PrimHead(V2576)
reg22551 := __e.Call(__defun__shen_4lazyderef, reg22550, V2761)
V2577 := reg22551
_ = V2577
reg22552 := MakeSymbol("string")
reg22553 := PrimEqual(reg22552, V2577)
var reg22662 Obj
if reg22553 == True {
reg22554 := PrimTail(V2576)
reg22555 := __e.Call(__defun__shen_4lazyderef, reg22554, V2761)
V2578 := reg22555
_ = V2578
reg22556 := Nil;
reg22557 := PrimEqual(reg22556, V2578)
var reg22604 Obj
if reg22557 == True {
reg22558 := PrimTail(V2567)
Hyp := reg22558
_ = Hyp
reg22559 := __e.Call(__defun__shen_4incinfs)
_ = reg22559
reg22560 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22561 := MakeSymbol(":")
reg22562 := MakeSymbol("string")
reg22563 := Nil;
reg22564 := PrimCons(reg22562, reg22563)
reg22565 := PrimCons(reg22561, reg22564)
reg22566 := PrimCons(reg22560, reg22565)
reg22567 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22568 := MakeSymbol(":")
reg22569 := MakeSymbol("string")
reg22570 := Nil;
reg22571 := PrimCons(reg22569, reg22570)
reg22572 := PrimCons(reg22568, reg22571)
reg22573 := PrimCons(reg22567, reg22572)
reg22574 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22575 := PrimCons(reg22573, reg22574)
reg22576 := PrimCons(reg22566, reg22575)
reg22577 := __e.Call(__defun__bind, V2760, reg22576, V2761, V2762)
reg22604 = reg22577
} else {
reg22578 := __e.Call(__defun__shen_4pvar_2, V2578)
var reg22603 Obj
if reg22578 == True {
reg22579 := Nil;
reg22580 := __e.Call(__defun__shen_4bindv, V2578, reg22579, V2761)
_ = reg22580
reg22581 := PrimTail(V2567)
Hyp := reg22581
_ = Hyp
reg22582 := __e.Call(__defun__shen_4incinfs)
_ = reg22582
reg22583 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22584 := MakeSymbol(":")
reg22585 := MakeSymbol("string")
reg22586 := Nil;
reg22587 := PrimCons(reg22585, reg22586)
reg22588 := PrimCons(reg22584, reg22587)
reg22589 := PrimCons(reg22583, reg22588)
reg22590 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22591 := MakeSymbol(":")
reg22592 := MakeSymbol("string")
reg22593 := Nil;
reg22594 := PrimCons(reg22592, reg22593)
reg22595 := PrimCons(reg22591, reg22594)
reg22596 := PrimCons(reg22590, reg22595)
reg22597 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22598 := PrimCons(reg22596, reg22597)
reg22599 := PrimCons(reg22589, reg22598)
reg22600 := __e.Call(__defun__bind, V2760, reg22599, V2761, V2762)
Result := reg22600
_ = Result
reg22601 := __e.Call(__defun__shen_4unbindv, V2578, V2761)
_ = reg22601
reg22603 = Result
} else {
reg22602 := False;
reg22603 = reg22602
}
reg22604 = reg22603
}
reg22662 = reg22604
} else {
reg22605 := __e.Call(__defun__shen_4pvar_2, V2577)
var reg22661 Obj
if reg22605 == True {
reg22606 := MakeSymbol("string")
reg22607 := __e.Call(__defun__shen_4bindv, V2577, reg22606, V2761)
_ = reg22607
reg22608 := PrimTail(V2576)
reg22609 := __e.Call(__defun__shen_4lazyderef, reg22608, V2761)
V2579 := reg22609
_ = V2579
reg22610 := Nil;
reg22611 := PrimEqual(reg22610, V2579)
var reg22658 Obj
if reg22611 == True {
reg22612 := PrimTail(V2567)
Hyp := reg22612
_ = Hyp
reg22613 := __e.Call(__defun__shen_4incinfs)
_ = reg22613
reg22614 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22615 := MakeSymbol(":")
reg22616 := MakeSymbol("string")
reg22617 := Nil;
reg22618 := PrimCons(reg22616, reg22617)
reg22619 := PrimCons(reg22615, reg22618)
reg22620 := PrimCons(reg22614, reg22619)
reg22621 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22622 := MakeSymbol(":")
reg22623 := MakeSymbol("string")
reg22624 := Nil;
reg22625 := PrimCons(reg22623, reg22624)
reg22626 := PrimCons(reg22622, reg22625)
reg22627 := PrimCons(reg22621, reg22626)
reg22628 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22629 := PrimCons(reg22627, reg22628)
reg22630 := PrimCons(reg22620, reg22629)
reg22631 := __e.Call(__defun__bind, V2760, reg22630, V2761, V2762)
reg22658 = reg22631
} else {
reg22632 := __e.Call(__defun__shen_4pvar_2, V2579)
var reg22657 Obj
if reg22632 == True {
reg22633 := Nil;
reg22634 := __e.Call(__defun__shen_4bindv, V2579, reg22633, V2761)
_ = reg22634
reg22635 := PrimTail(V2567)
Hyp := reg22635
_ = Hyp
reg22636 := __e.Call(__defun__shen_4incinfs)
_ = reg22636
reg22637 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22638 := MakeSymbol(":")
reg22639 := MakeSymbol("string")
reg22640 := Nil;
reg22641 := PrimCons(reg22639, reg22640)
reg22642 := PrimCons(reg22638, reg22641)
reg22643 := PrimCons(reg22637, reg22642)
reg22644 := __e.Call(__defun__shen_4lazyderef, Y, V2761)
reg22645 := MakeSymbol(":")
reg22646 := MakeSymbol("string")
reg22647 := Nil;
reg22648 := PrimCons(reg22646, reg22647)
reg22649 := PrimCons(reg22645, reg22648)
reg22650 := PrimCons(reg22644, reg22649)
reg22651 := __e.Call(__defun__shen_4lazyderef, Hyp, V2761)
reg22652 := PrimCons(reg22650, reg22651)
reg22653 := PrimCons(reg22643, reg22652)
reg22654 := __e.Call(__defun__bind, V2760, reg22653, V2761, V2762)
Result := reg22654
_ = Result
reg22655 := __e.Call(__defun__shen_4unbindv, V2579, V2761)
_ = reg22655
reg22657 = Result
} else {
reg22656 := False;
reg22657 = reg22656
}
reg22658 = reg22657
}
Result := reg22658
_ = Result
reg22659 := __e.Call(__defun__shen_4unbindv, V2577, V2761)
_ = reg22659
reg22661 = Result
} else {
reg22660 := False;
reg22661 = reg22660
}
reg22662 = reg22661
}
reg22664 = reg22662
} else {
reg22663 := False;
reg22664 = reg22663
}
reg22666 = reg22664
} else {
reg22665 := False;
reg22666 = reg22665
}
reg22668 = reg22666
} else {
reg22667 := False;
reg22668 = reg22667
}
reg22670 = reg22668
} else {
reg22669 := False;
reg22670 = reg22669
}
reg22672 = reg22670
} else {
reg22671 := False;
reg22672 = reg22671
}
reg22674 = reg22672
} else {
reg22673 := False;
reg22674 = reg22673
}
reg22676 = reg22674
} else {
reg22675 := False;
reg22676 = reg22675
}
reg22678 = reg22676
} else {
reg22677 := False;
reg22678 = reg22677
}
reg22680 = reg22678
} else {
reg22679 := False;
reg22680 = reg22679
}
reg22682 = reg22680
} else {
reg22681 := False;
reg22682 = reg22681
}
Case := reg22682
_ = Case
reg22683 := False;
reg22684 := PrimEqual(Case, reg22683)
if reg22684 == True {
reg22685 := __e.Call(__defun__shen_4lazyderef, V2759, V2761)
V2580 := reg22685
_ = V2580
reg22686 := PrimIsPair(V2580)
if reg22686 == True {
reg22687 := PrimHead(V2580)
X := reg22687
_ = X
reg22688 := PrimTail(V2580)
Hyp := reg22688
_ = Hyp
reg22689 := __e.Call(__defun__shen_4newpv, V2761)
NewHyps := reg22689
_ = NewHyps
reg22690 := __e.Call(__defun__shen_4incinfs)
_ = reg22690
reg22691 := __e.Call(__defun__shen_4lazyderef, X, V2761)
reg22692 := __e.Call(__defun__shen_4lazyderef, NewHyps, V2761)
reg22693 := PrimCons(reg22691, reg22692)
reg22694 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4t_d_1hyps, Hyp, NewHyps, V2761, V2762)
return
}, 0)
__ctx.TailApply(__defun__bind, V2760, reg22693, V2761, reg22694)
return
} else {
reg22697 := False;
__ctx.Return(reg22697)
return
}
} else {
__ctx.Return(Case)
return
}
} else {
__ctx.Return(Case)
return
}
} else {
__ctx.Return(Case)
return
}
} else {
__ctx.Return(Case)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.t*-hyps", value: __defun__shen_4t_d_1hyps})

__defun__shen_4show = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2779 := __args[0]
_ = V2779
V2780 := __args[1]
_ = V2780
V2781 := __args[2]
_ = V2781
V2782 := __args[3]
_ = V2782
reg22698 := MakeSymbol("shen.*spy*")
reg22699 := PrimValue(reg22698)
if reg22699 == True {
reg22700 := __e.Call(__defun__shen_4line)
_ = reg22700
reg22701 := __e.Call(__defun__shen_4deref, V2779, V2781)
reg22702 := __e.Call(__defun__shen_4show_1p, reg22701)
_ = reg22702
reg22703 := MakeNumber(1)
reg22704 := __e.Call(__defun__nl, reg22703)
_ = reg22704
reg22705 := MakeNumber(1)
reg22706 := __e.Call(__defun__nl, reg22705)
_ = reg22706
reg22707 := __e.Call(__defun__shen_4deref, V2780, V2781)
reg22708 := MakeNumber(1)
reg22709 := __e.Call(__defun__shen_4show_1assumptions, reg22707, reg22708)
_ = reg22709
reg22710 := MakeString("\n> ")
reg22711 := __e.Call(__defun__stoutput)
reg22712 := __e.Call(__defun__shen_4prhush, reg22710, reg22711)
_ = reg22712
reg22713 := __e.Call(__defun__shen_4pause_1for_1user)
_ = reg22713
__ctx.TailApply(__defun__thaw, V2782)
return
} else {
__ctx.TailApply(__defun__thaw, V2782)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.show", value: __defun__shen_4show})

__defun__shen_4line = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg22716 := __e.Call(__defun__inferences)
Infs := reg22716
_ = Infs
reg22717 := MakeString("____________________________________________________________ ")
reg22718 := MakeString(" inference")
reg22719 := MakeNumber(1)
reg22720 := PrimEqual(reg22719, Infs)
var reg22723 Obj
if reg22720 == True {
reg22721 := MakeString("")
reg22723 = reg22721
} else {
reg22722 := MakeString("s")
reg22723 = reg22722
}
reg22724 := MakeString(" \n?- ")
reg22725 := MakeSymbol("shen.a")
reg22726 := __e.Call(__defun__shen_4app, reg22723, reg22724, reg22725)
reg22727 := PrimStringConcat(reg22718, reg22726)
reg22728 := MakeSymbol("shen.a")
reg22729 := __e.Call(__defun__shen_4app, Infs, reg22727, reg22728)
reg22730 := PrimStringConcat(reg22717, reg22729)
reg22731 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg22730, reg22731)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.line", value: __defun__shen_4line})

__defun__shen_4show_1p = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2784 := __args[0]
_ = V2784
reg22733 := PrimIsPair(V2784)
var reg22767 Obj
if reg22733 == True {
reg22734 := PrimTail(V2784)
reg22735 := PrimIsPair(reg22734)
var reg22762 Obj
if reg22735 == True {
reg22736 := MakeSymbol(":")
reg22737 := PrimTail(V2784)
reg22738 := PrimHead(reg22737)
reg22739 := PrimEqual(reg22736, reg22738)
var reg22757 Obj
if reg22739 == True {
reg22740 := PrimTail(V2784)
reg22741 := PrimTail(reg22740)
reg22742 := PrimIsPair(reg22741)
var reg22752 Obj
if reg22742 == True {
reg22743 := Nil;
reg22744 := PrimTail(V2784)
reg22745 := PrimTail(reg22744)
reg22746 := PrimTail(reg22745)
reg22747 := PrimEqual(reg22743, reg22746)
var reg22750 Obj
if reg22747 == True {
reg22748 := True;
reg22750 = reg22748
} else {
reg22749 := False;
reg22750 = reg22749
}
reg22752 = reg22750
} else {
reg22751 := False;
reg22752 = reg22751
}
var reg22755 Obj
if reg22752 == True {
reg22753 := True;
reg22755 = reg22753
} else {
reg22754 := False;
reg22755 = reg22754
}
reg22757 = reg22755
} else {
reg22756 := False;
reg22757 = reg22756
}
var reg22760 Obj
if reg22757 == True {
reg22758 := True;
reg22760 = reg22758
} else {
reg22759 := False;
reg22760 = reg22759
}
reg22762 = reg22760
} else {
reg22761 := False;
reg22762 = reg22761
}
var reg22765 Obj
if reg22762 == True {
reg22763 := True;
reg22765 = reg22763
} else {
reg22764 := False;
reg22765 = reg22764
}
reg22767 = reg22765
} else {
reg22766 := False;
reg22767 = reg22766
}
if reg22767 == True {
reg22768 := PrimHead(V2784)
reg22769 := MakeString(" : ")
reg22770 := PrimTail(V2784)
reg22771 := PrimTail(reg22770)
reg22772 := PrimHead(reg22771)
reg22773 := MakeString("")
reg22774 := MakeSymbol("shen.r")
reg22775 := __e.Call(__defun__shen_4app, reg22772, reg22773, reg22774)
reg22776 := PrimStringConcat(reg22769, reg22775)
reg22777 := MakeSymbol("shen.r")
reg22778 := __e.Call(__defun__shen_4app, reg22768, reg22776, reg22777)
reg22779 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg22778, reg22779)
return
} else {
reg22781 := MakeString("")
reg22782 := MakeSymbol("shen.r")
reg22783 := __e.Call(__defun__shen_4app, V2784, reg22781, reg22782)
reg22784 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg22783, reg22784)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.show-p", value: __defun__shen_4show_1p})

__defun__shen_4show_1assumptions = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2789 := __args[0]
_ = V2789
V2790 := __args[1]
_ = V2790
reg22786 := Nil;
reg22787 := PrimEqual(reg22786, V2789)
if reg22787 == True {
reg22788 := MakeSymbol("shen.skip")
__ctx.Return(reg22788)
return
} else {
reg22789 := PrimIsPair(V2789)
if reg22789 == True {
reg22790 := MakeString(". ")
reg22791 := MakeSymbol("shen.a")
reg22792 := __e.Call(__defun__shen_4app, V2790, reg22790, reg22791)
reg22793 := __e.Call(__defun__stoutput)
reg22794 := __e.Call(__defun__shen_4prhush, reg22792, reg22793)
_ = reg22794
reg22795 := PrimHead(V2789)
reg22796 := __e.Call(__defun__shen_4show_1p, reg22795)
_ = reg22796
reg22797 := MakeNumber(1)
reg22798 := __e.Call(__defun__nl, reg22797)
_ = reg22798
reg22799 := PrimTail(V2789)
reg22800 := MakeNumber(1)
reg22801 := PrimNumberAdd(V2790, reg22800)
__ctx.TailApply(__defun__shen_4show_1assumptions, reg22799, reg22801)
return
} else {
reg22803 := MakeSymbol("shen.show-assumptions")
__ctx.TailApply(__defun__shen_4f__error, reg22803)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.show-assumptions", value: __defun__shen_4show_1assumptions})

__defun__shen_4pause_1for_1user = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg22805 := __e.Call(__defun__stinput)
reg22806 := PrimReadByte(reg22805)
Byte := reg22806
_ = Byte
reg22807 := MakeNumber(94)
reg22808 := PrimEqual(Byte, reg22807)
if reg22808 == True {
reg22809 := MakeString("input aborted\n")
reg22810 := PrimSimpleError(reg22809)
__ctx.Return(reg22810)
return
} else {
reg22811 := MakeNumber(1)
__ctx.TailApply(__defun__nl, reg22811)
return
}
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.pause-for-user", value: __defun__shen_4pause_1for_1user})

__defun__shen_4typedf_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2792 := __args[0]
_ = V2792
reg22813 := MakeSymbol("shen.*signedfuncs*")
reg22814 := PrimValue(reg22813)
reg22815 := __e.Call(__defun__assoc, V2792, reg22814)
reg22816 := PrimIsPair(reg22815)
__ctx.Return(reg22816)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.typedf?", value: __defun__shen_4typedf_2})

__defun__shen_4sigf = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2794 := __args[0]
_ = V2794
reg22817 := MakeSymbol("shen.type-signature-of-")
__ctx.TailApply(__defun__concat, reg22817, V2794)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sigf", value: __defun__shen_4sigf})

__defun__shen_4placeholder = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg22819 := MakeSymbol("&&")
__ctx.TailApply(__defun__gensym, reg22819)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.placeholder", value: __defun__shen_4placeholder})

__defun__shen_4base = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2799 := __args[0]
_ = V2799
V2800 := __args[1]
_ = V2800
V2801 := __args[2]
_ = V2801
V2802 := __args[3]
_ = V2802
reg22821 := __e.Call(__defun__shen_4lazyderef, V2800, V2801)
V2483 := reg22821
_ = V2483
reg22822 := MakeSymbol("number")
reg22823 := PrimEqual(reg22822, V2483)
var reg22838 Obj
if reg22823 == True {
reg22824 := __e.Call(__defun__shen_4incinfs)
_ = reg22824
reg22825 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22826 := PrimIsNumber(reg22825)
reg22827 := __e.Call(__defun__fwhen, reg22826, V2801, V2802)
reg22838 = reg22827
} else {
reg22828 := __e.Call(__defun__shen_4pvar_2, V2483)
var reg22837 Obj
if reg22828 == True {
reg22829 := MakeSymbol("number")
reg22830 := __e.Call(__defun__shen_4bindv, V2483, reg22829, V2801)
_ = reg22830
reg22831 := __e.Call(__defun__shen_4incinfs)
_ = reg22831
reg22832 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22833 := PrimIsNumber(reg22832)
reg22834 := __e.Call(__defun__fwhen, reg22833, V2801, V2802)
Result := reg22834
_ = Result
reg22835 := __e.Call(__defun__shen_4unbindv, V2483, V2801)
_ = reg22835
reg22837 = Result
} else {
reg22836 := False;
reg22837 = reg22836
}
reg22838 = reg22837
}
Case := reg22838
_ = Case
reg22839 := False;
reg22840 := PrimEqual(Case, reg22839)
if reg22840 == True {
reg22841 := __e.Call(__defun__shen_4lazyderef, V2800, V2801)
V2484 := reg22841
_ = V2484
reg22842 := MakeSymbol("boolean")
reg22843 := PrimEqual(reg22842, V2484)
var reg22858 Obj
if reg22843 == True {
reg22844 := __e.Call(__defun__shen_4incinfs)
_ = reg22844
reg22845 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22846 := __e.Call(__defun__boolean_2, reg22845)
reg22847 := __e.Call(__defun__fwhen, reg22846, V2801, V2802)
reg22858 = reg22847
} else {
reg22848 := __e.Call(__defun__shen_4pvar_2, V2484)
var reg22857 Obj
if reg22848 == True {
reg22849 := MakeSymbol("boolean")
reg22850 := __e.Call(__defun__shen_4bindv, V2484, reg22849, V2801)
_ = reg22850
reg22851 := __e.Call(__defun__shen_4incinfs)
_ = reg22851
reg22852 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22853 := __e.Call(__defun__boolean_2, reg22852)
reg22854 := __e.Call(__defun__fwhen, reg22853, V2801, V2802)
Result := reg22854
_ = Result
reg22855 := __e.Call(__defun__shen_4unbindv, V2484, V2801)
_ = reg22855
reg22857 = Result
} else {
reg22856 := False;
reg22857 = reg22856
}
reg22858 = reg22857
}
Case := reg22858
_ = Case
reg22859 := False;
reg22860 := PrimEqual(Case, reg22859)
if reg22860 == True {
reg22861 := __e.Call(__defun__shen_4lazyderef, V2800, V2801)
V2485 := reg22861
_ = V2485
reg22862 := MakeSymbol("string")
reg22863 := PrimEqual(reg22862, V2485)
var reg22878 Obj
if reg22863 == True {
reg22864 := __e.Call(__defun__shen_4incinfs)
_ = reg22864
reg22865 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22866 := PrimIsString(reg22865)
reg22867 := __e.Call(__defun__fwhen, reg22866, V2801, V2802)
reg22878 = reg22867
} else {
reg22868 := __e.Call(__defun__shen_4pvar_2, V2485)
var reg22877 Obj
if reg22868 == True {
reg22869 := MakeSymbol("string")
reg22870 := __e.Call(__defun__shen_4bindv, V2485, reg22869, V2801)
_ = reg22870
reg22871 := __e.Call(__defun__shen_4incinfs)
_ = reg22871
reg22872 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22873 := PrimIsString(reg22872)
reg22874 := __e.Call(__defun__fwhen, reg22873, V2801, V2802)
Result := reg22874
_ = Result
reg22875 := __e.Call(__defun__shen_4unbindv, V2485, V2801)
_ = reg22875
reg22877 = Result
} else {
reg22876 := False;
reg22877 = reg22876
}
reg22878 = reg22877
}
Case := reg22878
_ = Case
reg22879 := False;
reg22880 := PrimEqual(Case, reg22879)
if reg22880 == True {
reg22881 := __e.Call(__defun__shen_4lazyderef, V2800, V2801)
V2486 := reg22881
_ = V2486
reg22882 := MakeSymbol("symbol")
reg22883 := PrimEqual(reg22882, V2486)
var reg22908 Obj
if reg22883 == True {
reg22884 := __e.Call(__defun__shen_4incinfs)
_ = reg22884
reg22885 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22886 := PrimIsSymbol(reg22885)
reg22887 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg22888 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22889 := __e.Call(__defun__shen_4ue_2, reg22888)
reg22890 := PrimNot(reg22889)
__ctx.TailApply(__defun__fwhen, reg22890, V2801, V2802)
return
}, 0)
reg22892 := __e.Call(__defun__fwhen, reg22886, V2801, reg22887)
reg22908 = reg22892
} else {
reg22893 := __e.Call(__defun__shen_4pvar_2, V2486)
var reg22907 Obj
if reg22893 == True {
reg22894 := MakeSymbol("symbol")
reg22895 := __e.Call(__defun__shen_4bindv, V2486, reg22894, V2801)
_ = reg22895
reg22896 := __e.Call(__defun__shen_4incinfs)
_ = reg22896
reg22897 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22898 := PrimIsSymbol(reg22897)
reg22899 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg22900 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
reg22901 := __e.Call(__defun__shen_4ue_2, reg22900)
reg22902 := PrimNot(reg22901)
__ctx.TailApply(__defun__fwhen, reg22902, V2801, V2802)
return
}, 0)
reg22904 := __e.Call(__defun__fwhen, reg22898, V2801, reg22899)
Result := reg22904
_ = Result
reg22905 := __e.Call(__defun__shen_4unbindv, V2486, V2801)
_ = reg22905
reg22907 = Result
} else {
reg22906 := False;
reg22907 = reg22906
}
reg22908 = reg22907
}
Case := reg22908
_ = Case
reg22909 := False;
reg22910 := PrimEqual(Case, reg22909)
if reg22910 == True {
reg22911 := __e.Call(__defun__shen_4lazyderef, V2799, V2801)
V2487 := reg22911
_ = V2487
reg22912 := Nil;
reg22913 := PrimEqual(reg22912, V2487)
if reg22913 == True {
reg22914 := __e.Call(__defun__shen_4lazyderef, V2800, V2801)
V2488 := reg22914
_ = V2488
reg22915 := PrimIsPair(V2488)
if reg22915 == True {
reg22916 := PrimHead(V2488)
reg22917 := __e.Call(__defun__shen_4lazyderef, reg22916, V2801)
V2489 := reg22917
_ = V2489
reg22918 := MakeSymbol("list")
reg22919 := PrimEqual(reg22918, V2489)
if reg22919 == True {
reg22920 := PrimTail(V2488)
reg22921 := __e.Call(__defun__shen_4lazyderef, reg22920, V2801)
V2490 := reg22921
_ = V2490
reg22922 := PrimIsPair(V2490)
if reg22922 == True {
reg22923 := PrimHead(V2490)
A := reg22923
_ = A
reg22924 := PrimTail(V2490)
reg22925 := __e.Call(__defun__shen_4lazyderef, reg22924, V2801)
V2491 := reg22925
_ = V2491
reg22926 := Nil;
reg22927 := PrimEqual(reg22926, V2491)
if reg22927 == True {
reg22928 := __e.Call(__defun__shen_4incinfs)
_ = reg22928
__ctx.TailApply(__defun__thaw, V2802)
return
} else {
reg22930 := __e.Call(__defun__shen_4pvar_2, V2491)
if reg22930 == True {
reg22931 := Nil;
reg22932 := __e.Call(__defun__shen_4bindv, V2491, reg22931, V2801)
_ = reg22932
reg22933 := __e.Call(__defun__shen_4incinfs)
_ = reg22933
reg22934 := __e.Call(__defun__thaw, V2802)
Result := reg22934
_ = Result
reg22935 := __e.Call(__defun__shen_4unbindv, V2491, V2801)
_ = reg22935
__ctx.Return(Result)
return
} else {
reg22936 := False;
__ctx.Return(reg22936)
return
}
}
} else {
reg22937 := __e.Call(__defun__shen_4pvar_2, V2490)
if reg22937 == True {
reg22938 := __e.Call(__defun__shen_4newpv, V2801)
A := reg22938
_ = A
reg22939 := Nil;
reg22940 := PrimCons(A, reg22939)
reg22941 := __e.Call(__defun__shen_4bindv, V2490, reg22940, V2801)
_ = reg22941
reg22942 := __e.Call(__defun__shen_4incinfs)
_ = reg22942
reg22943 := __e.Call(__defun__thaw, V2802)
Result := reg22943
_ = Result
reg22944 := __e.Call(__defun__shen_4unbindv, V2490, V2801)
_ = reg22944
__ctx.Return(Result)
return
} else {
reg22945 := False;
__ctx.Return(reg22945)
return
}
}
} else {
reg22946 := __e.Call(__defun__shen_4pvar_2, V2489)
if reg22946 == True {
reg22947 := MakeSymbol("list")
reg22948 := __e.Call(__defun__shen_4bindv, V2489, reg22947, V2801)
_ = reg22948
reg22949 := PrimTail(V2488)
reg22950 := __e.Call(__defun__shen_4lazyderef, reg22949, V2801)
V2492 := reg22950
_ = V2492
reg22951 := PrimIsPair(V2492)
var reg22978 Obj
if reg22951 == True {
reg22952 := PrimHead(V2492)
A := reg22952
_ = A
reg22953 := PrimTail(V2492)
reg22954 := __e.Call(__defun__shen_4lazyderef, reg22953, V2801)
V2493 := reg22954
_ = V2493
reg22955 := Nil;
reg22956 := PrimEqual(reg22955, V2493)
var reg22967 Obj
if reg22956 == True {
reg22957 := __e.Call(__defun__shen_4incinfs)
_ = reg22957
reg22958 := __e.Call(__defun__thaw, V2802)
reg22967 = reg22958
} else {
reg22959 := __e.Call(__defun__shen_4pvar_2, V2493)
var reg22966 Obj
if reg22959 == True {
reg22960 := Nil;
reg22961 := __e.Call(__defun__shen_4bindv, V2493, reg22960, V2801)
_ = reg22961
reg22962 := __e.Call(__defun__shen_4incinfs)
_ = reg22962
reg22963 := __e.Call(__defun__thaw, V2802)
Result := reg22963
_ = Result
reg22964 := __e.Call(__defun__shen_4unbindv, V2493, V2801)
_ = reg22964
reg22966 = Result
} else {
reg22965 := False;
reg22966 = reg22965
}
reg22967 = reg22966
}
reg22978 = reg22967
} else {
reg22968 := __e.Call(__defun__shen_4pvar_2, V2492)
var reg22977 Obj
if reg22968 == True {
reg22969 := __e.Call(__defun__shen_4newpv, V2801)
A := reg22969
_ = A
reg22970 := Nil;
reg22971 := PrimCons(A, reg22970)
reg22972 := __e.Call(__defun__shen_4bindv, V2492, reg22971, V2801)
_ = reg22972
reg22973 := __e.Call(__defun__shen_4incinfs)
_ = reg22973
reg22974 := __e.Call(__defun__thaw, V2802)
Result := reg22974
_ = Result
reg22975 := __e.Call(__defun__shen_4unbindv, V2492, V2801)
_ = reg22975
reg22977 = Result
} else {
reg22976 := False;
reg22977 = reg22976
}
reg22978 = reg22977
}
Result := reg22978
_ = Result
reg22979 := __e.Call(__defun__shen_4unbindv, V2489, V2801)
_ = reg22979
__ctx.Return(Result)
return
} else {
reg22980 := False;
__ctx.Return(reg22980)
return
}
}
} else {
reg22981 := __e.Call(__defun__shen_4pvar_2, V2488)
if reg22981 == True {
reg22982 := __e.Call(__defun__shen_4newpv, V2801)
A := reg22982
_ = A
reg22983 := MakeSymbol("list")
reg22984 := Nil;
reg22985 := PrimCons(A, reg22984)
reg22986 := PrimCons(reg22983, reg22985)
reg22987 := __e.Call(__defun__shen_4bindv, V2488, reg22986, V2801)
_ = reg22987
reg22988 := __e.Call(__defun__shen_4incinfs)
_ = reg22988
reg22989 := __e.Call(__defun__thaw, V2802)
Result := reg22989
_ = Result
reg22990 := __e.Call(__defun__shen_4unbindv, V2488, V2801)
_ = reg22990
__ctx.Return(Result)
return
} else {
reg22991 := False;
__ctx.Return(reg22991)
return
}
}
} else {
reg22992 := False;
__ctx.Return(reg22992)
return
}
} else {
__ctx.Return(Case)
return
}
} else {
__ctx.Return(Case)
return
}
} else {
__ctx.Return(Case)
return
}
} else {
__ctx.Return(Case)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.base", value: __defun__shen_4base})

__defun__shen_4by__hypothesis = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2808 := __args[0]
_ = V2808
V2809 := __args[1]
_ = V2809
V2810 := __args[2]
_ = V2810
V2811 := __args[3]
_ = V2811
V2812 := __args[4]
_ = V2812
reg22993 := __e.Call(__defun__shen_4lazyderef, V2810, V2811)
V2474 := reg22993
_ = V2474
reg22994 := PrimIsPair(V2474)
var reg23029 Obj
if reg22994 == True {
reg22995 := PrimHead(V2474)
reg22996 := __e.Call(__defun__shen_4lazyderef, reg22995, V2811)
V2475 := reg22996
_ = V2475
reg22997 := PrimIsPair(V2475)
var reg23027 Obj
if reg22997 == True {
reg22998 := PrimHead(V2475)
Y := reg22998
_ = Y
reg22999 := PrimTail(V2475)
reg23000 := __e.Call(__defun__shen_4lazyderef, reg22999, V2811)
V2476 := reg23000
_ = V2476
reg23001 := PrimIsPair(V2476)
var reg23025 Obj
if reg23001 == True {
reg23002 := PrimHead(V2476)
reg23003 := __e.Call(__defun__shen_4lazyderef, reg23002, V2811)
V2477 := reg23003
_ = V2477
reg23004 := MakeSymbol(":")
reg23005 := PrimEqual(reg23004, V2477)
var reg23023 Obj
if reg23005 == True {
reg23006 := PrimTail(V2476)
reg23007 := __e.Call(__defun__shen_4lazyderef, reg23006, V2811)
V2478 := reg23007
_ = V2478
reg23008 := PrimIsPair(V2478)
var reg23021 Obj
if reg23008 == True {
reg23009 := PrimHead(V2478)
B := reg23009
_ = B
reg23010 := PrimTail(V2478)
reg23011 := __e.Call(__defun__shen_4lazyderef, reg23010, V2811)
V2479 := reg23011
_ = V2479
reg23012 := Nil;
reg23013 := PrimEqual(reg23012, V2479)
var reg23019 Obj
if reg23013 == True {
reg23014 := __e.Call(__defun__shen_4incinfs)
_ = reg23014
reg23015 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__unify_b, V2809, B, V2811, V2812)
return
}, 0)
reg23017 := __e.Call(__defun__identical, V2808, Y, V2811, reg23015)
reg23019 = reg23017
} else {
reg23018 := False;
reg23019 = reg23018
}
reg23021 = reg23019
} else {
reg23020 := False;
reg23021 = reg23020
}
reg23023 = reg23021
} else {
reg23022 := False;
reg23023 = reg23022
}
reg23025 = reg23023
} else {
reg23024 := False;
reg23025 = reg23024
}
reg23027 = reg23025
} else {
reg23026 := False;
reg23027 = reg23026
}
reg23029 = reg23027
} else {
reg23028 := False;
reg23029 = reg23028
}
Case := reg23029
_ = Case
reg23030 := False;
reg23031 := PrimEqual(Case, reg23030)
if reg23031 == True {
reg23032 := __e.Call(__defun__shen_4lazyderef, V2810, V2811)
V2480 := reg23032
_ = V2480
reg23033 := PrimIsPair(V2480)
if reg23033 == True {
reg23034 := PrimTail(V2480)
Hyp := reg23034
_ = Hyp
reg23035 := __e.Call(__defun__shen_4incinfs)
_ = reg23035
__ctx.TailApply(__defun__shen_4by__hypothesis, V2808, V2809, Hyp, V2811, V2812)
return
} else {
reg23037 := False;
__ctx.Return(reg23037)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.by_hypothesis", value: __defun__shen_4by__hypothesis})

__defun__shen_4t_d_1def = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2818 := __args[0]
_ = V2818
V2819 := __args[1]
_ = V2819
V2820 := __args[2]
_ = V2820
V2821 := __args[3]
_ = V2821
V2822 := __args[4]
_ = V2822
reg23038 := __e.Call(__defun__shen_4lazyderef, V2818, V2821)
V2468 := reg23038
_ = V2468
reg23039 := PrimIsPair(V2468)
if reg23039 == True {
reg23040 := PrimHead(V2468)
reg23041 := __e.Call(__defun__shen_4lazyderef, reg23040, V2821)
V2469 := reg23041
_ = V2469
reg23042 := MakeSymbol("define")
reg23043 := PrimEqual(reg23042, V2469)
if reg23043 == True {
reg23044 := PrimTail(V2468)
reg23045 := __e.Call(__defun__shen_4lazyderef, reg23044, V2821)
V2470 := reg23045
_ = V2470
reg23046 := PrimIsPair(V2470)
if reg23046 == True {
reg23047 := PrimHead(V2470)
F := reg23047
_ = F
reg23048 := PrimTail(V2470)
X := reg23048
_ = X
reg23049 := __e.Call(__defun__shen_4newpv, V2821)
Y := reg23049
_ = Y
reg23050 := __e.Call(__defun__shen_4newpv, V2821)
E := reg23050
_ = E
reg23051 := __e.Call(__defun__shen_4incinfs)
_ = reg23051
reg23052 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4_5sig_7rules_6, Y)
return
}, 1)
reg23054 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg23055 := PrimIsPair(E)
if reg23055 == True {
reg23056 := MakeString("parse error here: ")
reg23057 := MakeString("\n")
reg23058 := MakeSymbol("shen.s")
reg23059 := __e.Call(__defun__shen_4app, E, reg23057, reg23058)
reg23060 := PrimStringConcat(reg23056, reg23059)
reg23061 := PrimSimpleError(reg23060)
__ctx.Return(reg23061)
return
} else {
reg23062 := MakeString("parse error\n")
reg23063 := PrimSimpleError(reg23062)
__ctx.Return(reg23063)
return
}
}, 1)
reg23064 := __e.Call(__defun__compile, reg23052, X, reg23054)
__ctx.TailApply(__defun__shen_4t_d_1defh, reg23064, F, V2819, V2820, V2821, V2822)
return
} else {
reg23066 := False;
__ctx.Return(reg23066)
return
}
} else {
reg23067 := False;
__ctx.Return(reg23067)
return
}
} else {
reg23068 := False;
__ctx.Return(reg23068)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-def", value: __defun__shen_4t_d_1def})

__defun__shen_4t_d_1defh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2829 := __args[0]
_ = V2829
V2830 := __args[1]
_ = V2830
V2831 := __args[2]
_ = V2831
V2832 := __args[3]
_ = V2832
V2833 := __args[4]
_ = V2833
V2834 := __args[5]
_ = V2834
reg23069 := __e.Call(__defun__shen_4lazyderef, V2829, V2833)
V2464 := reg23069
_ = V2464
reg23070 := PrimIsPair(V2464)
if reg23070 == True {
reg23071 := PrimHead(V2464)
Sig := reg23071
_ = Sig
reg23072 := PrimTail(V2464)
Rules := reg23072
_ = Rules
reg23073 := __e.Call(__defun__shen_4incinfs)
_ = reg23073
reg23074 := __e.Call(__defun__shen_4ue_1sig, Sig)
__ctx.TailApply(__defun__shen_4t_d_1defhh, Sig, reg23074, V2830, V2831, V2832, Rules, V2833, V2834)
return
} else {
reg23076 := False;
__ctx.Return(reg23076)
return
}
}, 6)
__initDefs = append(__initDefs, defType{name: "shen.t*-defh", value: __defun__shen_4t_d_1defh})

__defun__shen_4t_d_1defhh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2843 := __args[0]
_ = V2843
V2844 := __args[1]
_ = V2844
V2845 := __args[2]
_ = V2845
V2846 := __args[3]
_ = V2846
V2847 := __args[4]
_ = V2847
V2848 := __args[5]
_ = V2848
V2849 := __args[6]
_ = V2849
V2850 := __args[7]
_ = V2850
reg23077 := __e.Call(__defun__shen_4incinfs)
_ = reg23077
reg23078 := MakeNumber(1)
reg23079 := MakeSymbol(":")
reg23080 := Nil;
reg23081 := PrimCons(V2844, reg23080)
reg23082 := PrimCons(reg23079, reg23081)
reg23083 := PrimCons(V2845, reg23082)
reg23084 := PrimCons(reg23083, V2847)
reg23085 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4memo, V2845, V2843, V2846, V2849, V2850)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d_1rules, V2848, V2844, reg23078, V2845, reg23084, V2849, reg23085)
return
}, 8)
__initDefs = append(__initDefs, defType{name: "shen.t*-defhh", value: __defun__shen_4t_d_1defhh})

__defun__shen_4memo = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2856 := __args[0]
_ = V2856
V2857 := __args[1]
_ = V2857
V2858 := __args[2]
_ = V2858
V2859 := __args[3]
_ = V2859
V2860 := __args[4]
_ = V2860
reg23088 := __e.Call(__defun__shen_4newpv, V2859)
Jnk := reg23088
_ = Jnk
reg23089 := __e.Call(__defun__shen_4incinfs)
_ = reg23089
reg23090 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23091 := __e.Call(__defun__shen_4lazyderef, V2856, V2859)
reg23092 := __e.Call(__defun__shen_4lazyderef, V2858, V2859)
reg23093 := __e.Call(__defun__declare, reg23091, reg23092)
__ctx.TailApply(__defun__bind, Jnk, reg23093, V2859, V2860)
return
}, 0)
__ctx.TailApply(__defun__unify_b, V2858, V2857, V2859, reg23090)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.memo", value: __defun__shen_4memo})

__defun__shen_4_5sig_7rules_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2862 := __args[0]
_ = V2862
reg23096 := __e.Call(__defun__shen_4_5signature_6, V2862)
Parse__shen_4_5signature_6 := reg23096
_ = Parse__shen_4_5signature_6
reg23097 := __e.Call(__defun__fail)
reg23098 := PrimEqual(reg23097, Parse__shen_4_5signature_6)
reg23099 := PrimNot(reg23098)
if reg23099 == True {
reg23100 := __e.Call(__defun__shen_4_5non_1ll_1rules_6, Parse__shen_4_5signature_6)
Parse__shen_4_5non_1ll_1rules_6 := reg23100
_ = Parse__shen_4_5non_1ll_1rules_6
reg23101 := __e.Call(__defun__fail)
reg23102 := PrimEqual(reg23101, Parse__shen_4_5non_1ll_1rules_6)
reg23103 := PrimNot(reg23102)
if reg23103 == True {
reg23104 := PrimHead(Parse__shen_4_5non_1ll_1rules_6)
reg23105 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_6)
reg23106 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5non_1ll_1rules_6)
reg23107 := PrimCons(reg23105, reg23106)
__ctx.TailApply(__defun__shen_4pair, reg23104, reg23107)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<sig+rules>", value: __defun__shen_4_5sig_7rules_6})

__defun__shen_4_5non_1ll_1rules_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2864 := __args[0]
_ = V2864
reg23111 := __e.Call(__defun__shen_4_5rule_6, V2864)
Parse__shen_4_5rule_6 := reg23111
_ = Parse__shen_4_5rule_6
reg23112 := __e.Call(__defun__fail)
reg23113 := PrimEqual(reg23112, Parse__shen_4_5rule_6)
reg23114 := PrimNot(reg23113)
var reg23127 Obj
if reg23114 == True {
reg23115 := __e.Call(__defun__shen_4_5non_1ll_1rules_6, Parse__shen_4_5rule_6)
Parse__shen_4_5non_1ll_1rules_6 := reg23115
_ = Parse__shen_4_5non_1ll_1rules_6
reg23116 := __e.Call(__defun__fail)
reg23117 := PrimEqual(reg23116, Parse__shen_4_5non_1ll_1rules_6)
reg23118 := PrimNot(reg23117)
var reg23125 Obj
if reg23118 == True {
reg23119 := PrimHead(Parse__shen_4_5non_1ll_1rules_6)
reg23120 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg23121 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5non_1ll_1rules_6)
reg23122 := PrimCons(reg23120, reg23121)
reg23123 := __e.Call(__defun__shen_4pair, reg23119, reg23122)
reg23125 = reg23123
} else {
reg23124 := __e.Call(__defun__fail)
reg23125 = reg23124
}
reg23127 = reg23125
} else {
reg23126 := __e.Call(__defun__fail)
reg23127 = reg23126
}
YaccParse := reg23127
_ = YaccParse
reg23128 := __e.Call(__defun__fail)
reg23129 := PrimEqual(YaccParse, reg23128)
if reg23129 == True {
reg23130 := __e.Call(__defun__shen_4_5rule_6, V2864)
Parse__shen_4_5rule_6 := reg23130
_ = Parse__shen_4_5rule_6
reg23131 := __e.Call(__defun__fail)
reg23132 := PrimEqual(reg23131, Parse__shen_4_5rule_6)
reg23133 := PrimNot(reg23132)
if reg23133 == True {
reg23134 := PrimHead(Parse__shen_4_5rule_6)
reg23135 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg23136 := Nil;
reg23137 := PrimCons(reg23135, reg23136)
__ctx.TailApply(__defun__shen_4pair, reg23134, reg23137)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<non-ll-rules>", value: __defun__shen_4_5non_1ll_1rules_6})

__defun__shen_4ue = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2866 := __args[0]
_ = V2866
reg23140 := PrimIsPair(V2866)
var reg23164 Obj
if reg23140 == True {
reg23141 := PrimTail(V2866)
reg23142 := PrimIsPair(reg23141)
var reg23159 Obj
if reg23142 == True {
reg23143 := Nil;
reg23144 := PrimTail(V2866)
reg23145 := PrimTail(reg23144)
reg23146 := PrimEqual(reg23143, reg23145)
var reg23154 Obj
if reg23146 == True {
reg23147 := PrimHead(V2866)
reg23148 := MakeSymbol("protect")
reg23149 := PrimEqual(reg23147, reg23148)
var reg23152 Obj
if reg23149 == True {
reg23150 := True;
reg23152 = reg23150
} else {
reg23151 := False;
reg23152 = reg23151
}
reg23154 = reg23152
} else {
reg23153 := False;
reg23154 = reg23153
}
var reg23157 Obj
if reg23154 == True {
reg23155 := True;
reg23157 = reg23155
} else {
reg23156 := False;
reg23157 = reg23156
}
reg23159 = reg23157
} else {
reg23158 := False;
reg23159 = reg23158
}
var reg23162 Obj
if reg23159 == True {
reg23160 := True;
reg23162 = reg23160
} else {
reg23161 := False;
reg23162 = reg23161
}
reg23164 = reg23162
} else {
reg23163 := False;
reg23164 = reg23163
}
if reg23164 == True {
__ctx.Return(V2866)
return
} else {
reg23165 := PrimIsPair(V2866)
if reg23165 == True {
reg23166 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4ue, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg23166, V2866)
return
} else {
reg23169 := PrimIsVariable(V2866)
if reg23169 == True {
reg23170 := MakeSymbol("&&")
__ctx.TailApply(__defun__concat, reg23170, V2866)
return
} else {
__ctx.Return(V2866)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue", value: __defun__shen_4ue})

__defun__shen_4ue_1sig = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2868 := __args[0]
_ = V2868
reg23172 := PrimIsPair(V2868)
if reg23172 == True {
reg23173 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4ue_1sig, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg23173, V2868)
return
} else {
reg23176 := PrimIsVariable(V2868)
if reg23176 == True {
reg23177 := MakeSymbol("&&&")
__ctx.TailApply(__defun__concat, reg23177, V2868)
return
} else {
__ctx.Return(V2868)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue-sig", value: __defun__shen_4ue_1sig})

__defun__shen_4ues = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2874 := __args[0]
_ = V2874
reg23179 := __e.Call(__defun__shen_4ue_2, V2874)
if reg23179 == True {
reg23180 := Nil;
reg23181 := PrimCons(V2874, reg23180)
__ctx.Return(reg23181)
return
} else {
reg23182 := PrimIsPair(V2874)
if reg23182 == True {
reg23183 := PrimHead(V2874)
reg23184 := __e.Call(__defun__shen_4ues, reg23183)
reg23185 := PrimTail(V2874)
reg23186 := __e.Call(__defun__shen_4ues, reg23185)
__ctx.TailApply(__defun__union, reg23184, reg23186)
return
} else {
reg23188 := Nil;
__ctx.Return(reg23188)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ues", value: __defun__shen_4ues})

__defun__shen_4ue_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2876 := __args[0]
_ = V2876
reg23189 := PrimIsSymbol(V2876)
if reg23189 == True {
reg23190 := PrimStr(V2876)
reg23191 := __e.Call(__defun__shen_4ue_1h_2, reg23190)
if reg23191 == True {
reg23192 := True;
__ctx.Return(reg23192)
return
} else {
reg23193 := False;
__ctx.Return(reg23193)
return
}
} else {
reg23194 := False;
__ctx.Return(reg23194)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue?", value: __defun__shen_4ue_2})

__defun__shen_4ue_1h_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2884 := __args[0]
_ = V2884
reg23195 := __e.Call(__defun__shen_4_7string_2, V2884)
var reg23221 Obj
if reg23195 == True {
reg23196 := MakeString("&")
reg23197 := MakeNumber(0)
reg23198 := PrimPos(V2884, reg23197)
reg23199 := PrimEqual(reg23196, reg23198)
var reg23216 Obj
if reg23199 == True {
reg23200 := PrimTailString(V2884)
reg23201 := __e.Call(__defun__shen_4_7string_2, reg23200)
var reg23211 Obj
if reg23201 == True {
reg23202 := MakeString("&")
reg23203 := PrimTailString(V2884)
reg23204 := MakeNumber(0)
reg23205 := PrimPos(reg23203, reg23204)
reg23206 := PrimEqual(reg23202, reg23205)
var reg23209 Obj
if reg23206 == True {
reg23207 := True;
reg23209 = reg23207
} else {
reg23208 := False;
reg23209 = reg23208
}
reg23211 = reg23209
} else {
reg23210 := False;
reg23211 = reg23210
}
var reg23214 Obj
if reg23211 == True {
reg23212 := True;
reg23214 = reg23212
} else {
reg23213 := False;
reg23214 = reg23213
}
reg23216 = reg23214
} else {
reg23215 := False;
reg23216 = reg23215
}
var reg23219 Obj
if reg23216 == True {
reg23217 := True;
reg23219 = reg23217
} else {
reg23218 := False;
reg23219 = reg23218
}
reg23221 = reg23219
} else {
reg23220 := False;
reg23221 = reg23220
}
if reg23221 == True {
reg23222 := True;
__ctx.Return(reg23222)
return
} else {
reg23223 := False;
__ctx.Return(reg23223)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue-h?", value: __defun__shen_4ue_1h_2})

__defun__shen_4t_d_1rules = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2892 := __args[0]
_ = V2892
V2893 := __args[1]
_ = V2893
V2894 := __args[2]
_ = V2894
V2895 := __args[3]
_ = V2895
V2896 := __args[4]
_ = V2896
V2897 := __args[5]
_ = V2897
V2898 := __args[6]
_ = V2898
reg23224 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg23224
_ = Throwcontrol
reg23225 := __e.Call(__defun__shen_4lazyderef, V2892, V2897)
V2448 := reg23225
_ = V2448
reg23226 := Nil;
reg23227 := PrimEqual(reg23226, V2448)
var reg23231 Obj
if reg23227 == True {
reg23228 := __e.Call(__defun__shen_4incinfs)
_ = reg23228
reg23229 := __e.Call(__defun__thaw, V2898)
reg23231 = reg23229
} else {
reg23230 := False;
reg23231 = reg23230
}
Case := reg23231
_ = Case
reg23232 := False;
reg23233 := PrimEqual(Case, reg23232)
var reg23267 Obj
if reg23233 == True {
reg23234 := __e.Call(__defun__shen_4lazyderef, V2892, V2897)
V2449 := reg23234
_ = V2449
reg23235 := PrimIsPair(V2449)
var reg23248 Obj
if reg23235 == True {
reg23236 := PrimHead(V2449)
Rule := reg23236
_ = Rule
reg23237 := PrimTail(V2449)
Rules := reg23237
_ = Rules
reg23238 := __e.Call(__defun__shen_4incinfs)
_ = reg23238
reg23239 := __e.Call(__defun__shen_4ue, Rule)
reg23240 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23241 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23242 := MakeNumber(1)
reg23243 := PrimNumberAdd(V2894, reg23242)
__ctx.TailApply(__defun__shen_4t_d_1rules, Rules, V2893, reg23243, V2895, V2896, V2897, V2898)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2897, reg23241)
return
}, 0)
reg23246 := __e.Call(__defun__shen_4t_d_1rule, reg23239, V2893, V2896, V2897, reg23240)
reg23248 = reg23246
} else {
reg23247 := False;
reg23248 = reg23247
}
Case := reg23248
_ = Case
reg23249 := False;
reg23250 := PrimEqual(Case, reg23249)
var reg23266 Obj
if reg23250 == True {
reg23251 := __e.Call(__defun__shen_4newpv, V2897)
Err := reg23251
_ = Err
reg23252 := __e.Call(__defun__shen_4incinfs)
_ = reg23252
reg23253 := MakeString("type error in rule ")
reg23254 := __e.Call(__defun__shen_4lazyderef, V2894, V2897)
reg23255 := MakeString(" of ")
reg23256 := __e.Call(__defun__shen_4lazyderef, V2895, V2897)
reg23257 := MakeString("")
reg23258 := MakeSymbol("shen.a")
reg23259 := __e.Call(__defun__shen_4app, reg23256, reg23257, reg23258)
reg23260 := PrimStringConcat(reg23255, reg23259)
reg23261 := MakeSymbol("shen.a")
reg23262 := __e.Call(__defun__shen_4app, reg23254, reg23260, reg23261)
reg23263 := PrimStringConcat(reg23253, reg23262)
reg23264 := PrimSimpleError(reg23263)
reg23265 := __e.Call(__defun__bind, Err, reg23264, V2897, V2898)
reg23266 = reg23265
} else {
reg23266 = Case
}
reg23267 = reg23266
} else {
reg23267 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg23267)
return
}, 7)
__initDefs = append(__initDefs, defType{name: "shen.t*-rules", value: __defun__shen_4t_d_1rules})

__defun__shen_4t_d_1rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2904 := __args[0]
_ = V2904
V2905 := __args[1]
_ = V2905
V2906 := __args[2]
_ = V2906
V2907 := __args[3]
_ = V2907
V2908 := __args[4]
_ = V2908
reg23269 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg23269
_ = Throwcontrol
reg23270 := __e.Call(__defun__shen_4lazyderef, V2904, V2907)
V2440 := reg23270
_ = V2440
reg23271 := PrimIsPair(V2440)
var reg23300 Obj
if reg23271 == True {
reg23272 := PrimHead(V2440)
Patterns := reg23272
_ = Patterns
reg23273 := PrimTail(V2440)
reg23274 := __e.Call(__defun__shen_4lazyderef, reg23273, V2907)
V2441 := reg23274
_ = V2441
reg23275 := PrimIsPair(V2441)
var reg23298 Obj
if reg23275 == True {
reg23276 := PrimHead(V2441)
Action := reg23276
_ = Action
reg23277 := PrimTail(V2441)
reg23278 := __e.Call(__defun__shen_4lazyderef, reg23277, V2907)
V2442 := reg23278
_ = V2442
reg23279 := Nil;
reg23280 := PrimEqual(reg23279, V2442)
var reg23296 Obj
if reg23280 == True {
reg23281 := __e.Call(__defun__shen_4newpv, V2907)
NewHyps := reg23281
_ = NewHyps
reg23282 := __e.Call(__defun__shen_4incinfs)
_ = reg23282
reg23283 := __e.Call(__defun__shen_4placeholders, Patterns)
reg23284 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23285 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23286 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23287 := __e.Call(__defun__shen_4ue, Action)
reg23288 := __e.Call(__defun__shen_4curry, reg23287)
reg23289 := __e.Call(__defun__shen_4result_1type, Patterns, V2905)
reg23290 := __e.Call(__defun__shen_4patthyps, Patterns, V2905, V2906)
__ctx.TailApply(__defun__shen_4t_d_1action, reg23288, reg23289, reg23290, V2907, V2908)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2907, reg23286)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d_1patterns, Patterns, V2905, NewHyps, V2907, reg23285)
return
}, 0)
reg23294 := __e.Call(__defun__shen_4newhyps, reg23283, V2906, NewHyps, V2907, reg23284)
reg23296 = reg23294
} else {
reg23295 := False;
reg23296 = reg23295
}
reg23298 = reg23296
} else {
reg23297 := False;
reg23298 = reg23297
}
reg23300 = reg23298
} else {
reg23299 := False;
reg23300 = reg23299
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg23300)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-rule", value: __defun__shen_4t_d_1rule})

__defun__shen_4placeholders = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2914 := __args[0]
_ = V2914
reg23302 := __e.Call(__defun__shen_4ue_2, V2914)
if reg23302 == True {
reg23303 := Nil;
reg23304 := PrimCons(V2914, reg23303)
__ctx.Return(reg23304)
return
} else {
reg23305 := PrimIsPair(V2914)
if reg23305 == True {
reg23306 := PrimHead(V2914)
reg23307 := __e.Call(__defun__shen_4placeholders, reg23306)
reg23308 := PrimTail(V2914)
reg23309 := __e.Call(__defun__shen_4placeholders, reg23308)
__ctx.TailApply(__defun__union, reg23307, reg23309)
return
} else {
reg23311 := Nil;
__ctx.Return(reg23311)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.placeholders", value: __defun__shen_4placeholders})

__defun__shen_4newhyps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2920 := __args[0]
_ = V2920
V2921 := __args[1]
_ = V2921
V2922 := __args[2]
_ = V2922
V2923 := __args[3]
_ = V2923
V2924 := __args[4]
_ = V2924
reg23312 := __e.Call(__defun__shen_4lazyderef, V2920, V2923)
V2427 := reg23312
_ = V2427
reg23313 := Nil;
reg23314 := PrimEqual(reg23313, V2427)
var reg23318 Obj
if reg23314 == True {
reg23315 := __e.Call(__defun__shen_4incinfs)
_ = reg23315
reg23316 := __e.Call(__defun__unify_b, V2922, V2921, V2923, V2924)
reg23318 = reg23316
} else {
reg23317 := False;
reg23318 = reg23317
}
Case := reg23318
_ = Case
reg23319 := False;
reg23320 := PrimEqual(Case, reg23319)
if reg23320 == True {
reg23321 := __e.Call(__defun__shen_4lazyderef, V2920, V2923)
V2428 := reg23321
_ = V2428
reg23322 := PrimIsPair(V2428)
if reg23322 == True {
reg23323 := PrimHead(V2428)
V2423 := reg23323
_ = V2423
reg23324 := PrimTail(V2428)
Vs := reg23324
_ = Vs
reg23325 := __e.Call(__defun__shen_4lazyderef, V2922, V2923)
V2429 := reg23325
_ = V2429
reg23326 := PrimIsPair(V2429)
if reg23326 == True {
reg23327 := PrimHead(V2429)
reg23328 := __e.Call(__defun__shen_4lazyderef, reg23327, V2923)
V2430 := reg23328
_ = V2430
reg23329 := PrimIsPair(V2430)
if reg23329 == True {
reg23330 := PrimHead(V2430)
V := reg23330
_ = V
reg23331 := PrimTail(V2430)
reg23332 := __e.Call(__defun__shen_4lazyderef, reg23331, V2923)
V2431 := reg23332
_ = V2431
reg23333 := PrimIsPair(V2431)
if reg23333 == True {
reg23334 := PrimHead(V2431)
reg23335 := __e.Call(__defun__shen_4lazyderef, reg23334, V2923)
V2432 := reg23335
_ = V2432
reg23336 := MakeSymbol(":")
reg23337 := PrimEqual(reg23336, V2432)
if reg23337 == True {
reg23338 := PrimTail(V2431)
reg23339 := __e.Call(__defun__shen_4lazyderef, reg23338, V2923)
V2433 := reg23339
_ = V2433
reg23340 := PrimIsPair(V2433)
if reg23340 == True {
reg23341 := PrimHead(V2433)
A := reg23341
_ = A
reg23342 := PrimTail(V2433)
reg23343 := __e.Call(__defun__shen_4lazyderef, reg23342, V2923)
V2434 := reg23343
_ = V2434
reg23344 := Nil;
reg23345 := PrimEqual(reg23344, V2434)
if reg23345 == True {
reg23346 := PrimTail(V2429)
NewHyp := reg23346
_ = NewHyp
reg23347 := __e.Call(__defun__shen_4incinfs)
_ = reg23347
reg23348 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
__ctx.TailApply(__defun__unify_b, V, V2423, V2923, reg23348)
return
} else {
reg23351 := __e.Call(__defun__shen_4pvar_2, V2434)
if reg23351 == True {
reg23352 := Nil;
reg23353 := __e.Call(__defun__shen_4bindv, V2434, reg23352, V2923)
_ = reg23353
reg23354 := PrimTail(V2429)
NewHyp := reg23354
_ = NewHyp
reg23355 := __e.Call(__defun__shen_4incinfs)
_ = reg23355
reg23356 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23358 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23356)
Result := reg23358
_ = Result
reg23359 := __e.Call(__defun__shen_4unbindv, V2434, V2923)
_ = reg23359
__ctx.Return(Result)
return
} else {
reg23360 := False;
__ctx.Return(reg23360)
return
}
}
} else {
reg23361 := __e.Call(__defun__shen_4pvar_2, V2433)
if reg23361 == True {
reg23362 := __e.Call(__defun__shen_4newpv, V2923)
A := reg23362
_ = A
reg23363 := Nil;
reg23364 := PrimCons(A, reg23363)
reg23365 := __e.Call(__defun__shen_4bindv, V2433, reg23364, V2923)
_ = reg23365
reg23366 := PrimTail(V2429)
NewHyp := reg23366
_ = NewHyp
reg23367 := __e.Call(__defun__shen_4incinfs)
_ = reg23367
reg23368 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23370 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23368)
Result := reg23370
_ = Result
reg23371 := __e.Call(__defun__shen_4unbindv, V2433, V2923)
_ = reg23371
__ctx.Return(Result)
return
} else {
reg23372 := False;
__ctx.Return(reg23372)
return
}
}
} else {
reg23373 := __e.Call(__defun__shen_4pvar_2, V2432)
if reg23373 == True {
reg23374 := MakeSymbol(":")
reg23375 := __e.Call(__defun__shen_4bindv, V2432, reg23374, V2923)
_ = reg23375
reg23376 := PrimTail(V2431)
reg23377 := __e.Call(__defun__shen_4lazyderef, reg23376, V2923)
V2435 := reg23377
_ = V2435
reg23378 := PrimIsPair(V2435)
var reg23414 Obj
if reg23378 == True {
reg23379 := PrimHead(V2435)
A := reg23379
_ = A
reg23380 := PrimTail(V2435)
reg23381 := __e.Call(__defun__shen_4lazyderef, reg23380, V2923)
V2436 := reg23381
_ = V2436
reg23382 := Nil;
reg23383 := PrimEqual(reg23382, V2436)
var reg23400 Obj
if reg23383 == True {
reg23384 := PrimTail(V2429)
NewHyp := reg23384
_ = NewHyp
reg23385 := __e.Call(__defun__shen_4incinfs)
_ = reg23385
reg23386 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23388 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23386)
reg23400 = reg23388
} else {
reg23389 := __e.Call(__defun__shen_4pvar_2, V2436)
var reg23399 Obj
if reg23389 == True {
reg23390 := Nil;
reg23391 := __e.Call(__defun__shen_4bindv, V2436, reg23390, V2923)
_ = reg23391
reg23392 := PrimTail(V2429)
NewHyp := reg23392
_ = NewHyp
reg23393 := __e.Call(__defun__shen_4incinfs)
_ = reg23393
reg23394 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23396 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23394)
Result := reg23396
_ = Result
reg23397 := __e.Call(__defun__shen_4unbindv, V2436, V2923)
_ = reg23397
reg23399 = Result
} else {
reg23398 := False;
reg23399 = reg23398
}
reg23400 = reg23399
}
reg23414 = reg23400
} else {
reg23401 := __e.Call(__defun__shen_4pvar_2, V2435)
var reg23413 Obj
if reg23401 == True {
reg23402 := __e.Call(__defun__shen_4newpv, V2923)
A := reg23402
_ = A
reg23403 := Nil;
reg23404 := PrimCons(A, reg23403)
reg23405 := __e.Call(__defun__shen_4bindv, V2435, reg23404, V2923)
_ = reg23405
reg23406 := PrimTail(V2429)
NewHyp := reg23406
_ = NewHyp
reg23407 := __e.Call(__defun__shen_4incinfs)
_ = reg23407
reg23408 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23410 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23408)
Result := reg23410
_ = Result
reg23411 := __e.Call(__defun__shen_4unbindv, V2435, V2923)
_ = reg23411
reg23413 = Result
} else {
reg23412 := False;
reg23413 = reg23412
}
reg23414 = reg23413
}
Result := reg23414
_ = Result
reg23415 := __e.Call(__defun__shen_4unbindv, V2432, V2923)
_ = reg23415
__ctx.Return(Result)
return
} else {
reg23416 := False;
__ctx.Return(reg23416)
return
}
}
} else {
reg23417 := __e.Call(__defun__shen_4pvar_2, V2431)
if reg23417 == True {
reg23418 := __e.Call(__defun__shen_4newpv, V2923)
A := reg23418
_ = A
reg23419 := MakeSymbol(":")
reg23420 := Nil;
reg23421 := PrimCons(A, reg23420)
reg23422 := PrimCons(reg23419, reg23421)
reg23423 := __e.Call(__defun__shen_4bindv, V2431, reg23422, V2923)
_ = reg23423
reg23424 := PrimTail(V2429)
NewHyp := reg23424
_ = NewHyp
reg23425 := __e.Call(__defun__shen_4incinfs)
_ = reg23425
reg23426 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23428 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23426)
Result := reg23428
_ = Result
reg23429 := __e.Call(__defun__shen_4unbindv, V2431, V2923)
_ = reg23429
__ctx.Return(Result)
return
} else {
reg23430 := False;
__ctx.Return(reg23430)
return
}
}
} else {
reg23431 := __e.Call(__defun__shen_4pvar_2, V2430)
if reg23431 == True {
reg23432 := __e.Call(__defun__shen_4newpv, V2923)
V := reg23432
_ = V
reg23433 := __e.Call(__defun__shen_4newpv, V2923)
A := reg23433
_ = A
reg23434 := MakeSymbol(":")
reg23435 := Nil;
reg23436 := PrimCons(A, reg23435)
reg23437 := PrimCons(reg23434, reg23436)
reg23438 := PrimCons(V, reg23437)
reg23439 := __e.Call(__defun__shen_4bindv, V2430, reg23438, V2923)
_ = reg23439
reg23440 := PrimTail(V2429)
NewHyp := reg23440
_ = NewHyp
reg23441 := __e.Call(__defun__shen_4incinfs)
_ = reg23441
reg23442 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23444 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23442)
Result := reg23444
_ = Result
reg23445 := __e.Call(__defun__shen_4unbindv, V2430, V2923)
_ = reg23445
__ctx.Return(Result)
return
} else {
reg23446 := False;
__ctx.Return(reg23446)
return
}
}
} else {
reg23447 := __e.Call(__defun__shen_4pvar_2, V2429)
if reg23447 == True {
reg23448 := __e.Call(__defun__shen_4newpv, V2923)
V := reg23448
_ = V
reg23449 := __e.Call(__defun__shen_4newpv, V2923)
A := reg23449
_ = A
reg23450 := __e.Call(__defun__shen_4newpv, V2923)
NewHyp := reg23450
_ = NewHyp
reg23451 := MakeSymbol(":")
reg23452 := Nil;
reg23453 := PrimCons(A, reg23452)
reg23454 := PrimCons(reg23451, reg23453)
reg23455 := PrimCons(V, reg23454)
reg23456 := PrimCons(reg23455, NewHyp)
reg23457 := __e.Call(__defun__shen_4bindv, V2429, reg23456, V2923)
_ = reg23457
reg23458 := __e.Call(__defun__shen_4incinfs)
_ = reg23458
reg23459 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V2921, NewHyp, V2923, V2924)
return
}, 0)
reg23461 := __e.Call(__defun__unify_b, V, V2423, V2923, reg23459)
Result := reg23461
_ = Result
reg23462 := __e.Call(__defun__shen_4unbindv, V2429, V2923)
_ = reg23462
__ctx.Return(Result)
return
} else {
reg23463 := False;
__ctx.Return(reg23463)
return
}
}
} else {
reg23464 := False;
__ctx.Return(reg23464)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.newhyps", value: __defun__shen_4newhyps})

__defun__shen_4patthyps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2930 := __args[0]
_ = V2930
V2931 := __args[1]
_ = V2931
V2932 := __args[2]
_ = V2932
reg23465 := Nil;
reg23466 := PrimEqual(reg23465, V2930)
if reg23466 == True {
__ctx.Return(V2932)
return
} else {
reg23467 := PrimIsPair(V2930)
var reg23507 Obj
if reg23467 == True {
reg23468 := PrimIsPair(V2931)
var reg23502 Obj
if reg23468 == True {
reg23469 := PrimTail(V2931)
reg23470 := PrimIsPair(reg23469)
var reg23497 Obj
if reg23470 == True {
reg23471 := MakeSymbol("-->")
reg23472 := PrimTail(V2931)
reg23473 := PrimHead(reg23472)
reg23474 := PrimEqual(reg23471, reg23473)
var reg23492 Obj
if reg23474 == True {
reg23475 := PrimTail(V2931)
reg23476 := PrimTail(reg23475)
reg23477 := PrimIsPair(reg23476)
var reg23487 Obj
if reg23477 == True {
reg23478 := Nil;
reg23479 := PrimTail(V2931)
reg23480 := PrimTail(reg23479)
reg23481 := PrimTail(reg23480)
reg23482 := PrimEqual(reg23478, reg23481)
var reg23485 Obj
if reg23482 == True {
reg23483 := True;
reg23485 = reg23483
} else {
reg23484 := False;
reg23485 = reg23484
}
reg23487 = reg23485
} else {
reg23486 := False;
reg23487 = reg23486
}
var reg23490 Obj
if reg23487 == True {
reg23488 := True;
reg23490 = reg23488
} else {
reg23489 := False;
reg23490 = reg23489
}
reg23492 = reg23490
} else {
reg23491 := False;
reg23492 = reg23491
}
var reg23495 Obj
if reg23492 == True {
reg23493 := True;
reg23495 = reg23493
} else {
reg23494 := False;
reg23495 = reg23494
}
reg23497 = reg23495
} else {
reg23496 := False;
reg23497 = reg23496
}
var reg23500 Obj
if reg23497 == True {
reg23498 := True;
reg23500 = reg23498
} else {
reg23499 := False;
reg23500 = reg23499
}
reg23502 = reg23500
} else {
reg23501 := False;
reg23502 = reg23501
}
var reg23505 Obj
if reg23502 == True {
reg23503 := True;
reg23505 = reg23503
} else {
reg23504 := False;
reg23505 = reg23504
}
reg23507 = reg23505
} else {
reg23506 := False;
reg23507 = reg23506
}
if reg23507 == True {
reg23508 := PrimHead(V2930)
reg23509 := MakeSymbol(":")
reg23510 := PrimHead(V2931)
reg23511 := Nil;
reg23512 := PrimCons(reg23510, reg23511)
reg23513 := PrimCons(reg23509, reg23512)
reg23514 := PrimCons(reg23508, reg23513)
reg23515 := PrimTail(V2930)
reg23516 := PrimTail(V2931)
reg23517 := PrimTail(reg23516)
reg23518 := PrimHead(reg23517)
reg23519 := __e.Call(__defun__shen_4patthyps, reg23515, reg23518, V2932)
__ctx.TailApply(__defun__adjoin, reg23514, reg23519)
return
} else {
reg23521 := MakeSymbol("shen.patthyps")
__ctx.TailApply(__defun__shen_4f__error, reg23521)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.patthyps", value: __defun__shen_4patthyps})

__defun__shen_4result_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2939 := __args[0]
_ = V2939
V2940 := __args[1]
_ = V2940
reg23523 := Nil;
reg23524 := PrimEqual(reg23523, V2939)
var reg23554 Obj
if reg23524 == True {
reg23525 := PrimIsPair(V2940)
var reg23549 Obj
if reg23525 == True {
reg23526 := MakeSymbol("-->")
reg23527 := PrimHead(V2940)
reg23528 := PrimEqual(reg23526, reg23527)
var reg23544 Obj
if reg23528 == True {
reg23529 := PrimTail(V2940)
reg23530 := PrimIsPair(reg23529)
var reg23539 Obj
if reg23530 == True {
reg23531 := Nil;
reg23532 := PrimTail(V2940)
reg23533 := PrimTail(reg23532)
reg23534 := PrimEqual(reg23531, reg23533)
var reg23537 Obj
if reg23534 == True {
reg23535 := True;
reg23537 = reg23535
} else {
reg23536 := False;
reg23537 = reg23536
}
reg23539 = reg23537
} else {
reg23538 := False;
reg23539 = reg23538
}
var reg23542 Obj
if reg23539 == True {
reg23540 := True;
reg23542 = reg23540
} else {
reg23541 := False;
reg23542 = reg23541
}
reg23544 = reg23542
} else {
reg23543 := False;
reg23544 = reg23543
}
var reg23547 Obj
if reg23544 == True {
reg23545 := True;
reg23547 = reg23545
} else {
reg23546 := False;
reg23547 = reg23546
}
reg23549 = reg23547
} else {
reg23548 := False;
reg23549 = reg23548
}
var reg23552 Obj
if reg23549 == True {
reg23550 := True;
reg23552 = reg23550
} else {
reg23551 := False;
reg23552 = reg23551
}
reg23554 = reg23552
} else {
reg23553 := False;
reg23554 = reg23553
}
if reg23554 == True {
reg23555 := PrimTail(V2940)
reg23556 := PrimHead(reg23555)
__ctx.Return(reg23556)
return
} else {
reg23557 := Nil;
reg23558 := PrimEqual(reg23557, V2939)
if reg23558 == True {
__ctx.Return(V2940)
return
} else {
reg23559 := PrimIsPair(V2939)
var reg23599 Obj
if reg23559 == True {
reg23560 := PrimIsPair(V2940)
var reg23594 Obj
if reg23560 == True {
reg23561 := PrimTail(V2940)
reg23562 := PrimIsPair(reg23561)
var reg23589 Obj
if reg23562 == True {
reg23563 := MakeSymbol("-->")
reg23564 := PrimTail(V2940)
reg23565 := PrimHead(reg23564)
reg23566 := PrimEqual(reg23563, reg23565)
var reg23584 Obj
if reg23566 == True {
reg23567 := PrimTail(V2940)
reg23568 := PrimTail(reg23567)
reg23569 := PrimIsPair(reg23568)
var reg23579 Obj
if reg23569 == True {
reg23570 := Nil;
reg23571 := PrimTail(V2940)
reg23572 := PrimTail(reg23571)
reg23573 := PrimTail(reg23572)
reg23574 := PrimEqual(reg23570, reg23573)
var reg23577 Obj
if reg23574 == True {
reg23575 := True;
reg23577 = reg23575
} else {
reg23576 := False;
reg23577 = reg23576
}
reg23579 = reg23577
} else {
reg23578 := False;
reg23579 = reg23578
}
var reg23582 Obj
if reg23579 == True {
reg23580 := True;
reg23582 = reg23580
} else {
reg23581 := False;
reg23582 = reg23581
}
reg23584 = reg23582
} else {
reg23583 := False;
reg23584 = reg23583
}
var reg23587 Obj
if reg23584 == True {
reg23585 := True;
reg23587 = reg23585
} else {
reg23586 := False;
reg23587 = reg23586
}
reg23589 = reg23587
} else {
reg23588 := False;
reg23589 = reg23588
}
var reg23592 Obj
if reg23589 == True {
reg23590 := True;
reg23592 = reg23590
} else {
reg23591 := False;
reg23592 = reg23591
}
reg23594 = reg23592
} else {
reg23593 := False;
reg23594 = reg23593
}
var reg23597 Obj
if reg23594 == True {
reg23595 := True;
reg23597 = reg23595
} else {
reg23596 := False;
reg23597 = reg23596
}
reg23599 = reg23597
} else {
reg23598 := False;
reg23599 = reg23598
}
if reg23599 == True {
reg23600 := PrimTail(V2939)
reg23601 := PrimTail(V2940)
reg23602 := PrimTail(reg23601)
reg23603 := PrimHead(reg23602)
__ctx.TailApply(__defun__shen_4result_1type, reg23600, reg23603)
return
} else {
reg23605 := MakeSymbol("shen.result-type")
__ctx.TailApply(__defun__shen_4f__error, reg23605)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.result-type", value: __defun__shen_4result_1type})

__defun__shen_4t_d_1patterns = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2946 := __args[0]
_ = V2946
V2947 := __args[1]
_ = V2947
V2948 := __args[2]
_ = V2948
V2949 := __args[3]
_ = V2949
V2950 := __args[4]
_ = V2950
reg23607 := __e.Call(__defun__shen_4lazyderef, V2946, V2949)
V2415 := reg23607
_ = V2415
reg23608 := Nil;
reg23609 := PrimEqual(reg23608, V2415)
var reg23613 Obj
if reg23609 == True {
reg23610 := __e.Call(__defun__shen_4incinfs)
_ = reg23610
reg23611 := __e.Call(__defun__thaw, V2950)
reg23613 = reg23611
} else {
reg23612 := False;
reg23613 = reg23612
}
Case := reg23613
_ = Case
reg23614 := False;
reg23615 := PrimEqual(Case, reg23614)
if reg23615 == True {
reg23616 := __e.Call(__defun__shen_4lazyderef, V2946, V2949)
V2416 := reg23616
_ = V2416
reg23617 := PrimIsPair(V2416)
if reg23617 == True {
reg23618 := PrimHead(V2416)
Pattern := reg23618
_ = Pattern
reg23619 := PrimTail(V2416)
Patterns := reg23619
_ = Patterns
reg23620 := __e.Call(__defun__shen_4lazyderef, V2947, V2949)
V2417 := reg23620
_ = V2417
reg23621 := PrimIsPair(V2417)
if reg23621 == True {
reg23622 := PrimHead(V2417)
A := reg23622
_ = A
reg23623 := PrimTail(V2417)
reg23624 := __e.Call(__defun__shen_4lazyderef, reg23623, V2949)
V2418 := reg23624
_ = V2418
reg23625 := PrimIsPair(V2418)
if reg23625 == True {
reg23626 := PrimHead(V2418)
reg23627 := __e.Call(__defun__shen_4lazyderef, reg23626, V2949)
V2419 := reg23627
_ = V2419
reg23628 := MakeSymbol("-->")
reg23629 := PrimEqual(reg23628, V2419)
if reg23629 == True {
reg23630 := PrimTail(V2418)
reg23631 := __e.Call(__defun__shen_4lazyderef, reg23630, V2949)
V2420 := reg23631
_ = V2420
reg23632 := PrimIsPair(V2420)
if reg23632 == True {
reg23633 := PrimHead(V2420)
B := reg23633
_ = B
reg23634 := PrimTail(V2420)
reg23635 := __e.Call(__defun__shen_4lazyderef, reg23634, V2949)
V2421 := reg23635
_ = V2421
reg23636 := Nil;
reg23637 := PrimEqual(reg23636, V2421)
if reg23637 == True {
reg23638 := __e.Call(__defun__shen_4incinfs)
_ = reg23638
reg23639 := MakeSymbol(":")
reg23640 := Nil;
reg23641 := PrimCons(A, reg23640)
reg23642 := PrimCons(reg23639, reg23641)
reg23643 := PrimCons(Pattern, reg23642)
reg23644 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4t_d_1patterns, Patterns, B, V2948, V2949, V2950)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d, reg23643, V2948, V2949, reg23644)
return
} else {
reg23647 := False;
__ctx.Return(reg23647)
return
}
} else {
reg23648 := False;
__ctx.Return(reg23648)
return
}
} else {
reg23649 := False;
__ctx.Return(reg23649)
return
}
} else {
reg23650 := False;
__ctx.Return(reg23650)
return
}
} else {
reg23651 := False;
__ctx.Return(reg23651)
return
}
} else {
reg23652 := False;
__ctx.Return(reg23652)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-patterns", value: __defun__shen_4t_d_1patterns})

__defun__shen_4t_d_1action = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2956 := __args[0]
_ = V2956
V2957 := __args[1]
_ = V2957
V2958 := __args[2]
_ = V2958
V2959 := __args[3]
_ = V2959
V2960 := __args[4]
_ = V2960
reg23653 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg23653
_ = Throwcontrol
reg23654 := __e.Call(__defun__shen_4lazyderef, V2956, V2959)
V2392 := reg23654
_ = V2392
reg23655 := PrimIsPair(V2392)
var reg23702 Obj
if reg23655 == True {
reg23656 := PrimHead(V2392)
reg23657 := __e.Call(__defun__shen_4lazyderef, reg23656, V2959)
V2393 := reg23657
_ = V2393
reg23658 := MakeSymbol("where")
reg23659 := PrimEqual(reg23658, V2393)
var reg23700 Obj
if reg23659 == True {
reg23660 := PrimTail(V2392)
reg23661 := __e.Call(__defun__shen_4lazyderef, reg23660, V2959)
V2394 := reg23661
_ = V2394
reg23662 := PrimIsPair(V2394)
var reg23698 Obj
if reg23662 == True {
reg23663 := PrimHead(V2394)
P := reg23663
_ = P
reg23664 := PrimTail(V2394)
reg23665 := __e.Call(__defun__shen_4lazyderef, reg23664, V2959)
V2395 := reg23665
_ = V2395
reg23666 := PrimIsPair(V2395)
var reg23696 Obj
if reg23666 == True {
reg23667 := PrimHead(V2395)
Action := reg23667
_ = Action
reg23668 := PrimTail(V2395)
reg23669 := __e.Call(__defun__shen_4lazyderef, reg23668, V2959)
V2396 := reg23669
_ = V2396
reg23670 := Nil;
reg23671 := PrimEqual(reg23670, V2396)
var reg23694 Obj
if reg23671 == True {
reg23672 := __e.Call(__defun__shen_4incinfs)
_ = reg23672
reg23673 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23674 := MakeSymbol(":")
reg23675 := MakeSymbol("boolean")
reg23676 := Nil;
reg23677 := PrimCons(reg23675, reg23676)
reg23678 := PrimCons(reg23674, reg23677)
reg23679 := PrimCons(P, reg23678)
reg23680 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23681 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23682 := MakeSymbol(":")
reg23683 := MakeSymbol("verified")
reg23684 := Nil;
reg23685 := PrimCons(reg23683, reg23684)
reg23686 := PrimCons(reg23682, reg23685)
reg23687 := PrimCons(P, reg23686)
reg23688 := PrimCons(reg23687, V2958)
__ctx.TailApply(__defun__shen_4t_d_1action, Action, V2957, reg23688, V2959, V2960)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V2959, reg23681)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d, reg23679, V2958, V2959, reg23680)
return
}, 0)
reg23692 := __e.Call(__defun__cut, Throwcontrol, V2959, reg23673)
reg23694 = reg23692
} else {
reg23693 := False;
reg23694 = reg23693
}
reg23696 = reg23694
} else {
reg23695 := False;
reg23696 = reg23695
}
reg23698 = reg23696
} else {
reg23697 := False;
reg23698 = reg23697
}
reg23700 = reg23698
} else {
reg23699 := False;
reg23700 = reg23699
}
reg23702 = reg23700
} else {
reg23701 := False;
reg23702 = reg23701
}
Case := reg23702
_ = Case
reg23703 := False;
reg23704 := PrimEqual(Case, reg23703)
var reg23840 Obj
if reg23704 == True {
reg23705 := __e.Call(__defun__shen_4lazyderef, V2956, V2959)
V2397 := reg23705
_ = V2397
reg23706 := PrimIsPair(V2397)
var reg23781 Obj
if reg23706 == True {
reg23707 := PrimHead(V2397)
reg23708 := __e.Call(__defun__shen_4lazyderef, reg23707, V2959)
V2398 := reg23708
_ = V2398
reg23709 := MakeSymbol("shen.choicepoint!")
reg23710 := PrimEqual(reg23709, V2398)
var reg23779 Obj
if reg23710 == True {
reg23711 := PrimTail(V2397)
reg23712 := __e.Call(__defun__shen_4lazyderef, reg23711, V2959)
V2399 := reg23712
_ = V2399
reg23713 := PrimIsPair(V2399)
var reg23777 Obj
if reg23713 == True {
reg23714 := PrimHead(V2399)
reg23715 := __e.Call(__defun__shen_4lazyderef, reg23714, V2959)
V2400 := reg23715
_ = V2400
reg23716 := PrimIsPair(V2400)
var reg23775 Obj
if reg23716 == True {
reg23717 := PrimHead(V2400)
reg23718 := __e.Call(__defun__shen_4lazyderef, reg23717, V2959)
V2401 := reg23718
_ = V2401
reg23719 := PrimIsPair(V2401)
var reg23773 Obj
if reg23719 == True {
reg23720 := PrimHead(V2401)
reg23721 := __e.Call(__defun__shen_4lazyderef, reg23720, V2959)
V2402 := reg23721
_ = V2402
reg23722 := MakeSymbol("fail-if")
reg23723 := PrimEqual(reg23722, V2402)
var reg23771 Obj
if reg23723 == True {
reg23724 := PrimTail(V2401)
reg23725 := __e.Call(__defun__shen_4lazyderef, reg23724, V2959)
V2403 := reg23725
_ = V2403
reg23726 := PrimIsPair(V2403)
var reg23769 Obj
if reg23726 == True {
reg23727 := PrimHead(V2403)
F := reg23727
_ = F
reg23728 := PrimTail(V2403)
reg23729 := __e.Call(__defun__shen_4lazyderef, reg23728, V2959)
V2404 := reg23729
_ = V2404
reg23730 := Nil;
reg23731 := PrimEqual(reg23730, V2404)
var reg23767 Obj
if reg23731 == True {
reg23732 := PrimTail(V2400)
reg23733 := __e.Call(__defun__shen_4lazyderef, reg23732, V2959)
V2405 := reg23733
_ = V2405
reg23734 := PrimIsPair(V2405)
var reg23765 Obj
if reg23734 == True {
reg23735 := PrimHead(V2405)
Action := reg23735
_ = Action
reg23736 := PrimTail(V2405)
reg23737 := __e.Call(__defun__shen_4lazyderef, reg23736, V2959)
V2406 := reg23737
_ = V2406
reg23738 := Nil;
reg23739 := PrimEqual(reg23738, V2406)
var reg23763 Obj
if reg23739 == True {
reg23740 := PrimTail(V2399)
reg23741 := __e.Call(__defun__shen_4lazyderef, reg23740, V2959)
V2407 := reg23741
_ = V2407
reg23742 := Nil;
reg23743 := PrimEqual(reg23742, V2407)
var reg23761 Obj
if reg23743 == True {
reg23744 := __e.Call(__defun__shen_4incinfs)
_ = reg23744
reg23745 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23746 := MakeSymbol("where")
reg23747 := MakeSymbol("not")
reg23748 := Nil;
reg23749 := PrimCons(Action, reg23748)
reg23750 := PrimCons(F, reg23749)
reg23751 := Nil;
reg23752 := PrimCons(reg23750, reg23751)
reg23753 := PrimCons(reg23747, reg23752)
reg23754 := Nil;
reg23755 := PrimCons(Action, reg23754)
reg23756 := PrimCons(reg23753, reg23755)
reg23757 := PrimCons(reg23746, reg23756)
__ctx.TailApply(__defun__shen_4t_d_1action, reg23757, V2957, V2958, V2959, V2960)
return
}, 0)
reg23759 := __e.Call(__defun__cut, Throwcontrol, V2959, reg23745)
reg23761 = reg23759
} else {
reg23760 := False;
reg23761 = reg23760
}
reg23763 = reg23761
} else {
reg23762 := False;
reg23763 = reg23762
}
reg23765 = reg23763
} else {
reg23764 := False;
reg23765 = reg23764
}
reg23767 = reg23765
} else {
reg23766 := False;
reg23767 = reg23766
}
reg23769 = reg23767
} else {
reg23768 := False;
reg23769 = reg23768
}
reg23771 = reg23769
} else {
reg23770 := False;
reg23771 = reg23770
}
reg23773 = reg23771
} else {
reg23772 := False;
reg23773 = reg23772
}
reg23775 = reg23773
} else {
reg23774 := False;
reg23775 = reg23774
}
reg23777 = reg23775
} else {
reg23776 := False;
reg23777 = reg23776
}
reg23779 = reg23777
} else {
reg23778 := False;
reg23779 = reg23778
}
reg23781 = reg23779
} else {
reg23780 := False;
reg23781 = reg23780
}
Case := reg23781
_ = Case
reg23782 := False;
reg23783 := PrimEqual(Case, reg23782)
var reg23839 Obj
if reg23783 == True {
reg23784 := __e.Call(__defun__shen_4lazyderef, V2956, V2959)
V2408 := reg23784
_ = V2408
reg23785 := PrimIsPair(V2408)
var reg23828 Obj
if reg23785 == True {
reg23786 := PrimHead(V2408)
reg23787 := __e.Call(__defun__shen_4lazyderef, reg23786, V2959)
V2409 := reg23787
_ = V2409
reg23788 := MakeSymbol("shen.choicepoint!")
reg23789 := PrimEqual(reg23788, V2409)
var reg23826 Obj
if reg23789 == True {
reg23790 := PrimTail(V2408)
reg23791 := __e.Call(__defun__shen_4lazyderef, reg23790, V2959)
V2410 := reg23791
_ = V2410
reg23792 := PrimIsPair(V2410)
var reg23824 Obj
if reg23792 == True {
reg23793 := PrimHead(V2410)
Action := reg23793
_ = Action
reg23794 := PrimTail(V2410)
reg23795 := __e.Call(__defun__shen_4lazyderef, reg23794, V2959)
V2411 := reg23795
_ = V2411
reg23796 := Nil;
reg23797 := PrimEqual(reg23796, V2411)
var reg23822 Obj
if reg23797 == True {
reg23798 := __e.Call(__defun__shen_4incinfs)
_ = reg23798
reg23799 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23800 := MakeSymbol("where")
reg23801 := MakeSymbol("not")
reg23802 := MakeSymbol("=")
reg23803 := Nil;
reg23804 := PrimCons(Action, reg23803)
reg23805 := PrimCons(reg23802, reg23804)
reg23806 := MakeSymbol("fail")
reg23807 := Nil;
reg23808 := PrimCons(reg23806, reg23807)
reg23809 := Nil;
reg23810 := PrimCons(reg23808, reg23809)
reg23811 := PrimCons(reg23805, reg23810)
reg23812 := Nil;
reg23813 := PrimCons(reg23811, reg23812)
reg23814 := PrimCons(reg23801, reg23813)
reg23815 := Nil;
reg23816 := PrimCons(Action, reg23815)
reg23817 := PrimCons(reg23814, reg23816)
reg23818 := PrimCons(reg23800, reg23817)
__ctx.TailApply(__defun__shen_4t_d_1action, reg23818, V2957, V2958, V2959, V2960)
return
}, 0)
reg23820 := __e.Call(__defun__cut, Throwcontrol, V2959, reg23799)
reg23822 = reg23820
} else {
reg23821 := False;
reg23822 = reg23821
}
reg23824 = reg23822
} else {
reg23823 := False;
reg23824 = reg23823
}
reg23826 = reg23824
} else {
reg23825 := False;
reg23826 = reg23825
}
reg23828 = reg23826
} else {
reg23827 := False;
reg23828 = reg23827
}
Case := reg23828
_ = Case
reg23829 := False;
reg23830 := PrimEqual(Case, reg23829)
var reg23838 Obj
if reg23830 == True {
reg23831 := __e.Call(__defun__shen_4incinfs)
_ = reg23831
reg23832 := MakeSymbol(":")
reg23833 := Nil;
reg23834 := PrimCons(V2957, reg23833)
reg23835 := PrimCons(reg23832, reg23834)
reg23836 := PrimCons(V2956, reg23835)
reg23837 := __e.Call(__defun__shen_4t_d, reg23836, V2958, V2959, V2960)
reg23838 = reg23837
} else {
reg23838 = Case
}
reg23839 = reg23838
} else {
reg23839 = Case
}
reg23840 = reg23839
} else {
reg23840 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg23840)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-action", value: __defun__shen_4t_d_1action})

__defun__findall = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2966 := __args[0]
_ = V2966
V2967 := __args[1]
_ = V2967
V2968 := __args[2]
_ = V2968
V2969 := __args[3]
_ = V2969
V2970 := __args[4]
_ = V2970
reg23842 := __e.Call(__defun__shen_4newpv, V2969)
B := reg23842
_ = B
reg23843 := __e.Call(__defun__shen_4newpv, V2969)
A := reg23843
_ = A
reg23844 := __e.Call(__defun__shen_4incinfs)
_ = reg23844
reg23845 := MakeSymbol("shen.a")
reg23846 := __e.Call(__defun__gensym, reg23845)
reg23847 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23848 := __e.Call(__defun__shen_4lazyderef, A, V2969)
reg23849 := Nil;
reg23850 := PrimSet(reg23848, reg23849)
reg23851 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4findallhelp, V2966, V2967, V2968, A, V2969, V2970)
return
}, 0)
__ctx.TailApply(__defun__bind, B, reg23850, V2969, reg23851)
return
}, 0)
__ctx.TailApply(__defun__bind, A, reg23846, V2969, reg23847)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "findall", value: __defun__findall})

__defun__shen_4findallhelp = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2977 := __args[0]
_ = V2977
V2978 := __args[1]
_ = V2978
V2979 := __args[2]
_ = V2979
V2980 := __args[3]
_ = V2980
V2981 := __args[4]
_ = V2981
V2982 := __args[5]
_ = V2982
reg23855 := __e.Call(__defun__shen_4incinfs)
_ = reg23855
reg23856 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23857 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg23858 := False;
__ctx.TailApply(__defun__fwhen, reg23858, V2981, V2982)
return
}, 0)
__ctx.TailApply(__defun__shen_4remember, V2980, V2977, V2981, reg23857)
return
}, 0)
reg23861 := __e.Call(__defun__call, V2978, V2981, reg23856)
Case := reg23861
_ = Case
reg23862 := False;
reg23863 := PrimEqual(Case, reg23862)
if reg23863 == True {
reg23864 := __e.Call(__defun__shen_4incinfs)
_ = reg23864
reg23865 := __e.Call(__defun__shen_4lazyderef, V2980, V2981)
reg23866 := PrimValue(reg23865)
__ctx.TailApply(__defun__bind, V2979, reg23866, V2981, V2982)
return
} else {
__ctx.Return(Case)
return
}
}, 6)
__initDefs = append(__initDefs, defType{name: "shen.findallhelp", value: __defun__shen_4findallhelp})

__defun__shen_4remember = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2987 := __args[0]
_ = V2987
V2988 := __args[1]
_ = V2988
V2989 := __args[2]
_ = V2989
V2990 := __args[3]
_ = V2990
reg23868 := __e.Call(__defun__shen_4newpv, V2989)
B := reg23868
_ = B
reg23869 := __e.Call(__defun__shen_4incinfs)
_ = reg23869
reg23870 := __e.Call(__defun__shen_4deref, V2987, V2989)
reg23871 := __e.Call(__defun__shen_4deref, V2988, V2989)
reg23872 := __e.Call(__defun__shen_4deref, V2987, V2989)
reg23873 := PrimValue(reg23872)
reg23874 := PrimCons(reg23871, reg23873)
reg23875 := PrimSet(reg23870, reg23874)
__ctx.TailApply(__defun__bind, B, reg23875, V2989, V2990)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.remember", value: __defun__shen_4remember})

}
