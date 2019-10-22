package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4dict Obj // shen.dict
var __defun__shen_4dict_2 Obj // shen.dict?
var __defun__shen_4dict_1capacity Obj // shen.dict-capacity
var __defun__shen_4dict_1count Obj // shen.dict-count
var __defun__shen_4dict_1count_1_6 Obj // shen.dict-count->
var __defun__shen_4_5_1dict_1bucket Obj // shen.<-dict-bucket
var __defun__shen_4dict_1bucket_1_6 Obj // shen.dict-bucket->
var __defun__shen_4dict_1update_1count Obj // shen.dict-update-count
var __defun__shen_4dict_1_6 Obj // shen.dict->
var __defun__shen_4_5_1dict Obj // shen.<-dict
var __defun__shen_4dict_1rm Obj // shen.dict-rm
var __defun__shen_4dict_1fold Obj // shen.dict-fold
var __defun__shen_4dict_1fold_1h Obj // shen.dict-fold-h
var __defun__shen_4bucket_1fold Obj // shen.bucket-fold
var __defun__shen_4dict_1keys Obj // shen.dict-keys
var __defun__shen_4dict_1values Obj // shen.dict-values

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297325 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg297325)
return
}, 0))
__defun__shen_4dict = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3255 := __args[0]
_ = V3255
reg297326 := MakeNumber(1)
reg297327 := PrimLessThan(V3255, reg297326)
if reg297327 == True {
reg297328 := MakeString("invalid initial dict size: ")
reg297329 := MakeString("")
reg297330 := MakeSymbol("shen.s")
reg297331 := __e.Call(__defun__shen_4app, V3255, reg297329, reg297330)
reg297332 := PrimStringConcat(reg297328, reg297331)
reg297333 := PrimSimpleError(reg297332)
__ctx.Return(reg297333)
return
} else {
reg297334 := MakeNumber(3)
reg297335 := PrimNumberAdd(reg297334, V3255)
reg297336 := PrimAbsvector(reg297335)
D := reg297336
_ = D
reg297337 := MakeNumber(0)
reg297338 := MakeSymbol("shen.dictionary")
reg297339 := PrimVectorSet(D, reg297337, reg297338)
Tag := reg297339
_ = Tag
reg297340 := MakeNumber(1)
reg297341 := PrimVectorSet(D, reg297340, V3255)
Capacity := reg297341
_ = Capacity
reg297342 := MakeNumber(2)
reg297343 := MakeNumber(0)
reg297344 := PrimVectorSet(D, reg297342, reg297343)
Count := reg297344
_ = Count
reg297345 := MakeNumber(3)
reg297346 := MakeNumber(2)
reg297347 := PrimNumberAdd(reg297346, V3255)
reg297348 := Nil;
reg297349 := __e.Call(__defun__shen_4fillvector, D, reg297345, reg297347, reg297348)
Fill := reg297349
_ = Fill
__ctx.Return(D)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dict", value: __defun__shen_4dict})

__defun__shen_4dict_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3257 := __args[0]
_ = V3257
reg297350 := PrimIsVector(V3257)
if reg297350 == True {
reg297351 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297352 := MakeNumber(0)
reg297353 := PrimVectorGet(V3257, reg297352)
__ctx.Return(reg297353)
return
}, 0)
reg297354 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg297355 := MakeSymbol("shen.not-dictionary")
__ctx.Return(reg297355)
return
}, 1)
reg297356 := __e.Try(reg297351).Catch(reg297354)
reg297357 := MakeSymbol("shen.dictionary")
reg297358 := PrimEqual(reg297356, reg297357)
if reg297358 == True {
reg297359 := True;
__ctx.Return(reg297359)
return
} else {
reg297360 := False;
__ctx.Return(reg297360)
return
}
} else {
reg297361 := False;
__ctx.Return(reg297361)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dict?", value: __defun__shen_4dict_2})

__defun__shen_4dict_1capacity = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3259 := __args[0]
_ = V3259
reg297362 := MakeNumber(1)
reg297363 := PrimVectorGet(V3259, reg297362)
__ctx.Return(reg297363)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dict-capacity", value: __defun__shen_4dict_1capacity})

__defun__shen_4dict_1count = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3261 := __args[0]
_ = V3261
reg297364 := MakeNumber(2)
reg297365 := PrimVectorGet(V3261, reg297364)
__ctx.Return(reg297365)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dict-count", value: __defun__shen_4dict_1count})

__defun__shen_4dict_1count_1_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3264 := __args[0]
_ = V3264
V3265 := __args[1]
_ = V3265
reg297366 := MakeNumber(2)
reg297367 := PrimVectorSet(V3264, reg297366, V3265)
__ctx.Return(reg297367)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.dict-count->", value: __defun__shen_4dict_1count_1_6})

__defun__shen_4_5_1dict_1bucket = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3268 := __args[0]
_ = V3268
V3269 := __args[1]
_ = V3269
reg297368 := MakeNumber(3)
reg297369 := PrimNumberAdd(reg297368, V3269)
reg297370 := PrimVectorGet(V3268, reg297369)
__ctx.Return(reg297370)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.<-dict-bucket", value: __defun__shen_4_5_1dict_1bucket})

__defun__shen_4dict_1bucket_1_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3273 := __args[0]
_ = V3273
V3274 := __args[1]
_ = V3274
V3275 := __args[2]
_ = V3275
reg297371 := MakeNumber(3)
reg297372 := PrimNumberAdd(reg297371, V3274)
reg297373 := PrimVectorSet(V3273, reg297372, V3275)
__ctx.Return(reg297373)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.dict-bucket->", value: __defun__shen_4dict_1bucket_1_6})

__defun__shen_4dict_1update_1count = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3279 := __args[0]
_ = V3279
V3280 := __args[1]
_ = V3280
V3281 := __args[2]
_ = V3281
reg297374 := __e.Call(__defun__length, V3281)
reg297375 := __e.Call(__defun__length, V3280)
reg297376 := PrimNumberSubtract(reg297374, reg297375)
Diff := reg297376
_ = Diff
reg297377 := __e.Call(__defun__shen_4dict_1count, V3279)
reg297378 := PrimNumberAdd(Diff, reg297377)
__ctx.TailApply(__defun__shen_4dict_1count_1_6, V3279, reg297378)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.dict-update-count", value: __defun__shen_4dict_1update_1count})

__defun__shen_4dict_1_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3285 := __args[0]
_ = V3285
V3286 := __args[1]
_ = V3286
V3287 := __args[2]
_ = V3287
reg297380 := __e.Call(__defun__shen_4dict_1capacity, V3285)
reg297381 := __e.Call(__defun__hash, V3286, reg297380)
N := reg297381
_ = N
reg297382 := __e.Call(__defun__shen_4_5_1dict_1bucket, V3285, N)
Bucket := reg297382
_ = Bucket
reg297383 := __e.Call(__defun__shen_4assoc_1set, V3286, V3287, Bucket)
NewBucket := reg297383
_ = NewBucket
reg297384 := __e.Call(__defun__shen_4dict_1bucket_1_6, V3285, N, NewBucket)
Change := reg297384
_ = Change
reg297385 := __e.Call(__defun__shen_4dict_1update_1count, V3285, Bucket, NewBucket)
Count := reg297385
_ = Count
__ctx.Return(V3287)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.dict->", value: __defun__shen_4dict_1_6})

__defun__shen_4_5_1dict = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3290 := __args[0]
_ = V3290
V3291 := __args[1]
_ = V3291
reg297386 := __e.Call(__defun__shen_4dict_1capacity, V3290)
reg297387 := __e.Call(__defun__hash, V3291, reg297386)
N := reg297387
_ = N
reg297388 := __e.Call(__defun__shen_4_5_1dict_1bucket, V3290, N)
Bucket := reg297388
_ = Bucket
reg297389 := __e.Call(__defun__assoc, V3291, Bucket)
Result := reg297389
_ = Result
reg297390 := __e.Call(__defun__empty_2, Result)
if reg297390 == True {
reg297391 := MakeString("value ")
reg297392 := MakeString(" not found in dict\n")
reg297393 := MakeSymbol("shen.a")
reg297394 := __e.Call(__defun__shen_4app, V3291, reg297392, reg297393)
reg297395 := PrimStringConcat(reg297391, reg297394)
reg297396 := PrimSimpleError(reg297395)
__ctx.Return(reg297396)
return
} else {
reg297397 := PrimTail(Result)
__ctx.Return(reg297397)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.<-dict", value: __defun__shen_4_5_1dict})

__defun__shen_4dict_1rm = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3294 := __args[0]
_ = V3294
V3295 := __args[1]
_ = V3295
reg297398 := __e.Call(__defun__shen_4dict_1capacity, V3294)
reg297399 := __e.Call(__defun__hash, V3295, reg297398)
N := reg297399
_ = N
reg297400 := __e.Call(__defun__shen_4_5_1dict_1bucket, V3294, N)
Bucket := reg297400
_ = Bucket
reg297401 := __e.Call(__defun__shen_4assoc_1rm, V3295, Bucket)
NewBucket := reg297401
_ = NewBucket
reg297402 := __e.Call(__defun__shen_4dict_1bucket_1_6, V3294, N, NewBucket)
Change := reg297402
_ = Change
reg297403 := __e.Call(__defun__shen_4dict_1update_1count, V3294, Bucket, NewBucket)
Count := reg297403
_ = Count
__ctx.Return(V3295)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.dict-rm", value: __defun__shen_4dict_1rm})

__defun__shen_4dict_1fold = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3299 := __args[0]
_ = V3299
V3300 := __args[1]
_ = V3300
V3301 := __args[2]
_ = V3301
reg297404 := __e.Call(__defun__shen_4dict_1capacity, V3300)
Limit := reg297404
_ = Limit
reg297405 := MakeNumber(0)
__ctx.TailApply(__defun__shen_4dict_1fold_1h, V3299, V3300, V3301, reg297405, Limit)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.dict-fold", value: __defun__shen_4dict_1fold})

__defun__shen_4dict_1fold_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3308 := __args[0]
_ = V3308
V3309 := __args[1]
_ = V3309
V3310 := __args[2]
_ = V3310
V3311 := __args[3]
_ = V3311
V3312 := __args[4]
_ = V3312
reg297407 := PrimEqual(V3312, V3311)
if reg297407 == True {
__ctx.Return(V3310)
return
} else {
reg297408 := __e.Call(__defun__shen_4_5_1dict_1bucket, V3309, V3311)
B := reg297408
_ = B
reg297409 := __e.Call(__defun__shen_4bucket_1fold, V3308, B, V3310)
Acc := reg297409
_ = Acc
reg297410 := MakeNumber(1)
reg297411 := PrimNumberAdd(reg297410, V3311)
__ctx.TailApply(__defun__shen_4dict_1fold_1h, V3308, V3309, Acc, reg297411, V3312)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.dict-fold-h", value: __defun__shen_4dict_1fold_1h})

__defun__shen_4bucket_1fold = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3316 := __args[0]
_ = V3316
V3317 := __args[1]
_ = V3317
V3318 := __args[2]
_ = V3318
reg297413 := Nil;
reg297414 := PrimEqual(reg297413, V3317)
if reg297414 == True {
__ctx.Return(V3318)
return
} else {
reg297415 := PrimIsPair(V3317)
var reg297422 Obj
if reg297415 == True {
reg297416 := PrimHead(V3317)
reg297417 := PrimIsPair(reg297416)
var reg297420 Obj
if reg297417 == True {
reg297418 := True;
reg297420 = reg297418
} else {
reg297419 := False;
reg297420 = reg297419
}
reg297422 = reg297420
} else {
reg297421 := False;
reg297422 = reg297421
}
if reg297422 == True {
reg297423 := PrimHead(V3317)
reg297424 := PrimHead(reg297423)
reg297425 := PrimHead(V3317)
reg297426 := PrimTail(reg297425)
reg297427 := PrimTail(V3317)
reg297428 := __e.Call(__defun__shen_4bucket_1fold, V3316, reg297427, V3318)
__ctx.TailApply(V3316, reg297424, reg297426, reg297428)
return
} else {
reg297430 := MakeSymbol("shen.bucket-fold")
__ctx.TailApply(__defun__shen_4f__error, reg297430)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.bucket-fold", value: __defun__shen_4bucket_1fold})

__defun__shen_4dict_1keys = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3320 := __args[0]
_ = V3320
reg297432 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
K := __args[0]
_ = K
reg297433 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ := __args[0]
_ = __
reg297434 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Acc := __args[0]
_ = Acc
reg297435 := PrimCons(K, Acc)
__ctx.Return(reg297435)
return
}, 1)
__ctx.Return(reg297434)
return
}, 1)
__ctx.Return(reg297433)
return
}, 1)
reg297436 := Nil;
__ctx.TailApply(__defun__shen_4dict_1fold, reg297432, V3320, reg297436)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dict-keys", value: __defun__shen_4dict_1keys})

__defun__shen_4dict_1values = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3322 := __args[0]
_ = V3322
reg297438 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ := __args[0]
_ = __
reg297439 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V := __args[0]
_ = V
reg297440 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Acc := __args[0]
_ = Acc
reg297441 := PrimCons(V, Acc)
__ctx.Return(reg297441)
return
}, 1)
__ctx.Return(reg297440)
return
}, 1)
__ctx.Return(reg297439)
return
}, 1)
reg297442 := Nil;
__ctx.TailApply(__defun__shen_4dict_1fold, reg297438, V3322, reg297442)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dict-values", value: __defun__shen_4dict_1values})

}
