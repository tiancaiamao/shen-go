package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4shen_1_6kl Obj // shen.shen->kl
var __defun__shen_4shen_1syntax_1error Obj // shen.shen-syntax-error
var __defun__shen_4_5define_6 Obj // shen.<define>
var __defun__shen_4_5name_6 Obj // shen.<name>
var __defun__shen_4sysfunc_2 Obj // shen.sysfunc?
var __defun__shen_4_5signature_6 Obj // shen.<signature>
var __defun__shen_4curry_1type Obj // shen.curry-type
var __defun__shen_4_5signature_1help_6 Obj // shen.<signature-help>
var __defun__shen_4_5rules_6 Obj // shen.<rules>
var __defun__shen_4_5rule_6 Obj // shen.<rule>
var __defun__shen_4fail__if Obj // shen.fail_if
var __defun__shen_4succeeds_2 Obj // shen.succeeds?
var __defun__shen_4_5patterns_6 Obj // shen.<patterns>
var __defun__shen_4_5pattern_6 Obj // shen.<pattern>
var __defun__shen_4constructor_1error Obj // shen.constructor-error
var __defun__shen_4_5simple__pattern_6 Obj // shen.<simple_pattern>
var __defun__shen_4_5pattern1_6 Obj // shen.<pattern1>
var __defun__shen_4_5pattern2_6 Obj // shen.<pattern2>
var __defun__shen_4_5action_6 Obj // shen.<action>
var __defun__shen_4_5guard_6 Obj // shen.<guard>
var __defun__shen_4compile__to__machine__code Obj // shen.compile_to_machine_code
var __defun__shen_4record_1source Obj // shen.record-source
var __defun__shen_4compile__to__lambda_7 Obj // shen.compile_to_lambda+
var __defun__shen_4update_1symbol_1table Obj // shen.update-symbol-table
var __defun__shen_4free__variable__check Obj // shen.free_variable_check
var __defun__shen_4extract__vars Obj // shen.extract_vars
var __defun__shen_4extract__free__vars Obj // shen.extract_free_vars
var __defun__shen_4free__variable__warnings Obj // shen.free_variable_warnings
var __defun__shen_4list__variables Obj // shen.list_variables
var __defun__shen_4strip_1protect Obj // shen.strip-protect
var __defun__shen_4linearise Obj // shen.linearise
var __defun__shen_4flatten Obj // shen.flatten
var __defun__shen_4linearise__help Obj // shen.linearise_help
var __defun__shen_4linearise__X Obj // shen.linearise_X
var __defun__shen_4aritycheck Obj // shen.aritycheck
var __defun__shen_4aritycheck_1name Obj // shen.aritycheck-name
var __defun__shen_4aritycheck_1action Obj // shen.aritycheck-action
var __defun__shen_4aah Obj // shen.aah
var __defun__shen_4abstract__rule Obj // shen.abstract_rule
var __defun__shen_4abstraction__build Obj // shen.abstraction_build
var __defun__shen_4parameters Obj // shen.parameters
var __defun__shen_4application__build Obj // shen.application_build
var __defun__shen_4compile__to__kl Obj // shen.compile_to_kl
var __defun__shen_4get_1type Obj // shen.get-type
var __defun__shen_4typextable Obj // shen.typextable
var __defun__shen_4assign_1types Obj // shen.assign-types
var __defun__shen_4atom_1type Obj // shen.atom-type
var __defun__shen_4store_1arity Obj // shen.store-arity
var __defun__shen_4reduce Obj // shen.reduce
var __defun__shen_4reduce__help Obj // shen.reduce_help
var __defun__shen_4_7string_2 Obj // shen.+string?
var __defun__shen_4_7vector_2 Obj // shen.+vector?
var __defun__shen_4ebr Obj // shen.ebr
var __defun__shen_4add__test Obj // shen.add_test
var __defun__shen_4cond_1expression Obj // shen.cond-expression
var __defun__shen_4cond_1form Obj // shen.cond-form
var __defun__shen_4encode_1choices Obj // shen.encode-choices
var __defun__shen_4case_1form Obj // shen.case-form
var __defun__shen_4embed_1and Obj // shen.embed-and
var __defun__shen_4err_1condition Obj // shen.err-condition
var __defun__shen_4sys_1error Obj // shen.sys-error

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg292142 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg292142)
return
}, 0))
__defun__shen_4shen_1_6kl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1384 := __args[0]
_ = V1384
V1385 := __args[1]
_ = V1385
reg292143 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5define_6, X)
return
}, 1)
reg292145 := PrimCons(V1384, V1385)
reg292146 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4shen_1syntax_1error, V1384, X)
return
}, 1)
__ctx.TailApply(__defun__compile, reg292143, reg292145, reg292146)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.shen->kl", value: __defun__shen_4shen_1_6kl})

__defun__shen_4shen_1syntax_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1392 := __args[0]
_ = V1392
V1393 := __args[1]
_ = V1393
reg292149 := PrimIsPair(V1393)
if reg292149 == True {
reg292150 := MakeString("syntax error in ")
reg292151 := MakeString(" here:\n\n ")
reg292152 := MakeNumber(50)
reg292153 := PrimHead(V1393)
reg292154 := __e.Call(__defun__shen_4next_150, reg292152, reg292153)
reg292155 := MakeString("\n")
reg292156 := MakeSymbol("shen.a")
reg292157 := __e.Call(__defun__shen_4app, reg292154, reg292155, reg292156)
reg292158 := PrimStringConcat(reg292151, reg292157)
reg292159 := MakeSymbol("shen.a")
reg292160 := __e.Call(__defun__shen_4app, V1392, reg292158, reg292159)
reg292161 := PrimStringConcat(reg292150, reg292160)
reg292162 := PrimSimpleError(reg292161)
__ctx.Return(reg292162)
return
} else {
reg292163 := MakeString("syntax error in ")
reg292164 := MakeString("\n")
reg292165 := MakeSymbol("shen.a")
reg292166 := __e.Call(__defun__shen_4app, V1392, reg292164, reg292165)
reg292167 := PrimStringConcat(reg292163, reg292166)
reg292168 := PrimSimpleError(reg292167)
__ctx.Return(reg292168)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.shen-syntax-error", value: __defun__shen_4shen_1syntax_1error})

__defun__shen_4_5define_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1395 := __args[0]
_ = V1395
reg292169 := __e.Call(__defun__shen_4_5name_6, V1395)
Parse__shen_4_5name_6 := reg292169
_ = Parse__shen_4_5name_6
reg292170 := __e.Call(__defun__fail)
reg292171 := PrimEqual(reg292170, Parse__shen_4_5name_6)
reg292172 := PrimNot(reg292171)
var reg292191 Obj
if reg292172 == True {
reg292173 := __e.Call(__defun__shen_4_5signature_6, Parse__shen_4_5name_6)
Parse__shen_4_5signature_6 := reg292173
_ = Parse__shen_4_5signature_6
reg292174 := __e.Call(__defun__fail)
reg292175 := PrimEqual(reg292174, Parse__shen_4_5signature_6)
reg292176 := PrimNot(reg292175)
var reg292189 Obj
if reg292176 == True {
reg292177 := __e.Call(__defun__shen_4_5rules_6, Parse__shen_4_5signature_6)
Parse__shen_4_5rules_6 := reg292177
_ = Parse__shen_4_5rules_6
reg292178 := __e.Call(__defun__fail)
reg292179 := PrimEqual(reg292178, Parse__shen_4_5rules_6)
reg292180 := PrimNot(reg292179)
var reg292187 Obj
if reg292180 == True {
reg292181 := PrimHead(Parse__shen_4_5rules_6)
reg292182 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5name_6)
reg292183 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rules_6)
reg292184 := __e.Call(__defun__shen_4compile__to__machine__code, reg292182, reg292183)
reg292185 := __e.Call(__defun__shen_4pair, reg292181, reg292184)
reg292187 = reg292185
} else {
reg292186 := __e.Call(__defun__fail)
reg292187 = reg292186
}
reg292189 = reg292187
} else {
reg292188 := __e.Call(__defun__fail)
reg292189 = reg292188
}
reg292191 = reg292189
} else {
reg292190 := __e.Call(__defun__fail)
reg292191 = reg292190
}
YaccParse := reg292191
_ = YaccParse
reg292192 := __e.Call(__defun__fail)
reg292193 := PrimEqual(YaccParse, reg292192)
if reg292193 == True {
reg292194 := __e.Call(__defun__shen_4_5name_6, V1395)
Parse__shen_4_5name_6 := reg292194
_ = Parse__shen_4_5name_6
reg292195 := __e.Call(__defun__fail)
reg292196 := PrimEqual(reg292195, Parse__shen_4_5name_6)
reg292197 := PrimNot(reg292196)
if reg292197 == True {
reg292198 := __e.Call(__defun__shen_4_5rules_6, Parse__shen_4_5name_6)
Parse__shen_4_5rules_6 := reg292198
_ = Parse__shen_4_5rules_6
reg292199 := __e.Call(__defun__fail)
reg292200 := PrimEqual(reg292199, Parse__shen_4_5rules_6)
reg292201 := PrimNot(reg292200)
if reg292201 == True {
reg292202 := PrimHead(Parse__shen_4_5rules_6)
reg292203 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5name_6)
reg292204 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rules_6)
reg292205 := __e.Call(__defun__shen_4compile__to__machine__code, reg292203, reg292204)
__ctx.TailApply(__defun__shen_4pair, reg292202, reg292205)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<define>", value: __defun__shen_4_5define_6})

__defun__shen_4_5name_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1397 := __args[0]
_ = V1397
reg292209 := PrimHead(V1397)
reg292210 := PrimIsPair(reg292209)
if reg292210 == True {
reg292211 := PrimHead(V1397)
reg292212 := PrimHead(reg292211)
Parse__X := reg292212
_ = Parse__X
reg292213 := PrimHead(V1397)
reg292214 := PrimTail(reg292213)
reg292215 := __e.Call(__defun__shen_4hdtl, V1397)
reg292216 := __e.Call(__defun__shen_4pair, reg292214, reg292215)
reg292217 := PrimHead(reg292216)
reg292218 := PrimIsSymbol(Parse__X)
var reg292225 Obj
if reg292218 == True {
reg292219 := __e.Call(__defun__shen_4sysfunc_2, Parse__X)
reg292220 := PrimNot(reg292219)
var reg292223 Obj
if reg292220 == True {
reg292221 := True;
reg292223 = reg292221
} else {
reg292222 := False;
reg292223 = reg292222
}
reg292225 = reg292223
} else {
reg292224 := False;
reg292225 = reg292224
}
var reg292230 Obj
if reg292225 == True {
reg292230 = Parse__X
} else {
reg292226 := MakeString(" is not a legitimate function name.\n")
reg292227 := MakeSymbol("shen.a")
reg292228 := __e.Call(__defun__shen_4app, Parse__X, reg292226, reg292227)
reg292229 := PrimSimpleError(reg292228)
reg292230 = reg292229
}
__ctx.TailApply(__defun__shen_4pair, reg292217, reg292230)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<name>", value: __defun__shen_4_5name_6})

__defun__shen_4sysfunc_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1399 := __args[0]
_ = V1399
reg292233 := MakeString("shen")
reg292234 := PrimIntern(reg292233)
reg292235 := MakeSymbol("shen.external-symbols")
reg292236 := MakeSymbol("*property-vector*")
reg292237 := PrimValue(reg292236)
reg292238 := __e.Call(__defun__get, reg292234, reg292235, reg292237)
__ctx.TailApply(__defun__element_2, V1399, reg292238)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sysfunc?", value: __defun__shen_4sysfunc_2})

__defun__shen_4_5signature_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1401 := __args[0]
_ = V1401
reg292240 := PrimHead(V1401)
reg292241 := PrimIsPair(reg292240)
var reg292250 Obj
if reg292241 == True {
reg292242 := MakeSymbol("{")
reg292243 := PrimHead(V1401)
reg292244 := PrimHead(reg292243)
reg292245 := PrimEqual(reg292242, reg292244)
var reg292248 Obj
if reg292245 == True {
reg292246 := True;
reg292248 = reg292246
} else {
reg292247 := False;
reg292248 = reg292247
}
reg292250 = reg292248
} else {
reg292249 := False;
reg292250 = reg292249
}
if reg292250 == True {
reg292251 := PrimHead(V1401)
reg292252 := PrimTail(reg292251)
reg292253 := __e.Call(__defun__shen_4hdtl, V1401)
reg292254 := __e.Call(__defun__shen_4pair, reg292252, reg292253)
reg292255 := __e.Call(__defun__shen_4_5signature_1help_6, reg292254)
Parse__shen_4_5signature_1help_6 := reg292255
_ = Parse__shen_4_5signature_1help_6
reg292256 := __e.Call(__defun__fail)
reg292257 := PrimEqual(reg292256, Parse__shen_4_5signature_1help_6)
reg292258 := PrimNot(reg292257)
if reg292258 == True {
reg292259 := PrimHead(Parse__shen_4_5signature_1help_6)
reg292260 := PrimIsPair(reg292259)
var reg292269 Obj
if reg292260 == True {
reg292261 := MakeSymbol("}")
reg292262 := PrimHead(Parse__shen_4_5signature_1help_6)
reg292263 := PrimHead(reg292262)
reg292264 := PrimEqual(reg292261, reg292263)
var reg292267 Obj
if reg292264 == True {
reg292265 := True;
reg292267 = reg292265
} else {
reg292266 := False;
reg292267 = reg292266
}
reg292269 = reg292267
} else {
reg292268 := False;
reg292269 = reg292268
}
if reg292269 == True {
reg292270 := PrimHead(Parse__shen_4_5signature_1help_6)
reg292271 := PrimTail(reg292270)
reg292272 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_1help_6)
reg292273 := __e.Call(__defun__shen_4pair, reg292271, reg292272)
reg292274 := PrimHead(reg292273)
reg292275 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_1help_6)
reg292276 := __e.Call(__defun__shen_4curry_1type, reg292275)
reg292277 := __e.Call(__defun__shen_4demodulate, reg292276)
__ctx.TailApply(__defun__shen_4pair, reg292274, reg292277)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<signature>", value: __defun__shen_4_5signature_6})

__defun__shen_4curry_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1403 := __args[0]
_ = V1403
reg292282 := PrimIsPair(V1403)
var reg292326 Obj
if reg292282 == True {
reg292283 := PrimTail(V1403)
reg292284 := PrimIsPair(reg292283)
var reg292321 Obj
if reg292284 == True {
reg292285 := MakeSymbol("-->")
reg292286 := PrimTail(V1403)
reg292287 := PrimHead(reg292286)
reg292288 := PrimEqual(reg292285, reg292287)
var reg292316 Obj
if reg292288 == True {
reg292289 := PrimTail(V1403)
reg292290 := PrimTail(reg292289)
reg292291 := PrimIsPair(reg292290)
var reg292311 Obj
if reg292291 == True {
reg292292 := PrimTail(V1403)
reg292293 := PrimTail(reg292292)
reg292294 := PrimTail(reg292293)
reg292295 := PrimIsPair(reg292294)
var reg292306 Obj
if reg292295 == True {
reg292296 := MakeSymbol("-->")
reg292297 := PrimTail(V1403)
reg292298 := PrimTail(reg292297)
reg292299 := PrimTail(reg292298)
reg292300 := PrimHead(reg292299)
reg292301 := PrimEqual(reg292296, reg292300)
var reg292304 Obj
if reg292301 == True {
reg292302 := True;
reg292304 = reg292302
} else {
reg292303 := False;
reg292304 = reg292303
}
reg292306 = reg292304
} else {
reg292305 := False;
reg292306 = reg292305
}
var reg292309 Obj
if reg292306 == True {
reg292307 := True;
reg292309 = reg292307
} else {
reg292308 := False;
reg292309 = reg292308
}
reg292311 = reg292309
} else {
reg292310 := False;
reg292311 = reg292310
}
var reg292314 Obj
if reg292311 == True {
reg292312 := True;
reg292314 = reg292312
} else {
reg292313 := False;
reg292314 = reg292313
}
reg292316 = reg292314
} else {
reg292315 := False;
reg292316 = reg292315
}
var reg292319 Obj
if reg292316 == True {
reg292317 := True;
reg292319 = reg292317
} else {
reg292318 := False;
reg292319 = reg292318
}
reg292321 = reg292319
} else {
reg292320 := False;
reg292321 = reg292320
}
var reg292324 Obj
if reg292321 == True {
reg292322 := True;
reg292324 = reg292322
} else {
reg292323 := False;
reg292324 = reg292323
}
reg292326 = reg292324
} else {
reg292325 := False;
reg292326 = reg292325
}
if reg292326 == True {
reg292327 := PrimHead(V1403)
reg292328 := MakeSymbol("-->")
reg292329 := PrimTail(V1403)
reg292330 := PrimTail(reg292329)
reg292331 := Nil;
reg292332 := PrimCons(reg292330, reg292331)
reg292333 := PrimCons(reg292328, reg292332)
reg292334 := PrimCons(reg292327, reg292333)
__ctx.TailApply(__defun__shen_4curry_1type, reg292334)
return
} else {
reg292336 := PrimIsPair(V1403)
var reg292380 Obj
if reg292336 == True {
reg292337 := PrimTail(V1403)
reg292338 := PrimIsPair(reg292337)
var reg292375 Obj
if reg292338 == True {
reg292339 := MakeSymbol("*")
reg292340 := PrimTail(V1403)
reg292341 := PrimHead(reg292340)
reg292342 := PrimEqual(reg292339, reg292341)
var reg292370 Obj
if reg292342 == True {
reg292343 := PrimTail(V1403)
reg292344 := PrimTail(reg292343)
reg292345 := PrimIsPair(reg292344)
var reg292365 Obj
if reg292345 == True {
reg292346 := PrimTail(V1403)
reg292347 := PrimTail(reg292346)
reg292348 := PrimTail(reg292347)
reg292349 := PrimIsPair(reg292348)
var reg292360 Obj
if reg292349 == True {
reg292350 := MakeSymbol("*")
reg292351 := PrimTail(V1403)
reg292352 := PrimTail(reg292351)
reg292353 := PrimTail(reg292352)
reg292354 := PrimHead(reg292353)
reg292355 := PrimEqual(reg292350, reg292354)
var reg292358 Obj
if reg292355 == True {
reg292356 := True;
reg292358 = reg292356
} else {
reg292357 := False;
reg292358 = reg292357
}
reg292360 = reg292358
} else {
reg292359 := False;
reg292360 = reg292359
}
var reg292363 Obj
if reg292360 == True {
reg292361 := True;
reg292363 = reg292361
} else {
reg292362 := False;
reg292363 = reg292362
}
reg292365 = reg292363
} else {
reg292364 := False;
reg292365 = reg292364
}
var reg292368 Obj
if reg292365 == True {
reg292366 := True;
reg292368 = reg292366
} else {
reg292367 := False;
reg292368 = reg292367
}
reg292370 = reg292368
} else {
reg292369 := False;
reg292370 = reg292369
}
var reg292373 Obj
if reg292370 == True {
reg292371 := True;
reg292373 = reg292371
} else {
reg292372 := False;
reg292373 = reg292372
}
reg292375 = reg292373
} else {
reg292374 := False;
reg292375 = reg292374
}
var reg292378 Obj
if reg292375 == True {
reg292376 := True;
reg292378 = reg292376
} else {
reg292377 := False;
reg292378 = reg292377
}
reg292380 = reg292378
} else {
reg292379 := False;
reg292380 = reg292379
}
if reg292380 == True {
reg292381 := PrimHead(V1403)
reg292382 := MakeSymbol("*")
reg292383 := PrimTail(V1403)
reg292384 := PrimTail(reg292383)
reg292385 := Nil;
reg292386 := PrimCons(reg292384, reg292385)
reg292387 := PrimCons(reg292382, reg292386)
reg292388 := PrimCons(reg292381, reg292387)
__ctx.TailApply(__defun__shen_4curry_1type, reg292388)
return
} else {
reg292390 := PrimIsPair(V1403)
if reg292390 == True {
reg292391 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4curry_1type, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg292391, V1403)
return
} else {
__ctx.Return(V1403)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.curry-type", value: __defun__shen_4curry_1type})

__defun__shen_4_5signature_1help_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1405 := __args[0]
_ = V1405
reg292394 := PrimHead(V1405)
reg292395 := PrimIsPair(reg292394)
var reg292422 Obj
if reg292395 == True {
reg292396 := PrimHead(V1405)
reg292397 := PrimHead(reg292396)
Parse__X := reg292397
_ = Parse__X
reg292398 := PrimHead(V1405)
reg292399 := PrimTail(reg292398)
reg292400 := __e.Call(__defun__shen_4hdtl, V1405)
reg292401 := __e.Call(__defun__shen_4pair, reg292399, reg292400)
reg292402 := __e.Call(__defun__shen_4_5signature_1help_6, reg292401)
Parse__shen_4_5signature_1help_6 := reg292402
_ = Parse__shen_4_5signature_1help_6
reg292403 := __e.Call(__defun__fail)
reg292404 := PrimEqual(reg292403, Parse__shen_4_5signature_1help_6)
reg292405 := PrimNot(reg292404)
var reg292420 Obj
if reg292405 == True {
reg292406 := MakeSymbol("{")
reg292407 := MakeSymbol("}")
reg292408 := Nil;
reg292409 := PrimCons(reg292407, reg292408)
reg292410 := PrimCons(reg292406, reg292409)
reg292411 := __e.Call(__defun__element_2, Parse__X, reg292410)
reg292412 := PrimNot(reg292411)
var reg292418 Obj
if reg292412 == True {
reg292413 := PrimHead(Parse__shen_4_5signature_1help_6)
reg292414 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_1help_6)
reg292415 := PrimCons(Parse__X, reg292414)
reg292416 := __e.Call(__defun__shen_4pair, reg292413, reg292415)
reg292418 = reg292416
} else {
reg292417 := __e.Call(__defun__fail)
reg292418 = reg292417
}
reg292420 = reg292418
} else {
reg292419 := __e.Call(__defun__fail)
reg292420 = reg292419
}
reg292422 = reg292420
} else {
reg292421 := __e.Call(__defun__fail)
reg292422 = reg292421
}
YaccParse := reg292422
_ = YaccParse
reg292423 := __e.Call(__defun__fail)
reg292424 := PrimEqual(YaccParse, reg292423)
if reg292424 == True {
reg292425 := __e.Call(__defun___5e_6, V1405)
Parse___5e_6 := reg292425
_ = Parse___5e_6
reg292426 := __e.Call(__defun__fail)
reg292427 := PrimEqual(reg292426, Parse___5e_6)
reg292428 := PrimNot(reg292427)
if reg292428 == True {
reg292429 := PrimHead(Parse___5e_6)
reg292430 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg292429, reg292430)
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
__initDefs = append(__initDefs, defType{name: "shen.<signature-help>", value: __defun__shen_4_5signature_1help_6})

__defun__shen_4_5rules_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1407 := __args[0]
_ = V1407
reg292433 := __e.Call(__defun__shen_4_5rule_6, V1407)
Parse__shen_4_5rule_6 := reg292433
_ = Parse__shen_4_5rule_6
reg292434 := __e.Call(__defun__fail)
reg292435 := PrimEqual(reg292434, Parse__shen_4_5rule_6)
reg292436 := PrimNot(reg292435)
var reg292450 Obj
if reg292436 == True {
reg292437 := __e.Call(__defun__shen_4_5rules_6, Parse__shen_4_5rule_6)
Parse__shen_4_5rules_6 := reg292437
_ = Parse__shen_4_5rules_6
reg292438 := __e.Call(__defun__fail)
reg292439 := PrimEqual(reg292438, Parse__shen_4_5rules_6)
reg292440 := PrimNot(reg292439)
var reg292448 Obj
if reg292440 == True {
reg292441 := PrimHead(Parse__shen_4_5rules_6)
reg292442 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg292443 := __e.Call(__defun__shen_4linearise, reg292442)
reg292444 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rules_6)
reg292445 := PrimCons(reg292443, reg292444)
reg292446 := __e.Call(__defun__shen_4pair, reg292441, reg292445)
reg292448 = reg292446
} else {
reg292447 := __e.Call(__defun__fail)
reg292448 = reg292447
}
reg292450 = reg292448
} else {
reg292449 := __e.Call(__defun__fail)
reg292450 = reg292449
}
YaccParse := reg292450
_ = YaccParse
reg292451 := __e.Call(__defun__fail)
reg292452 := PrimEqual(YaccParse, reg292451)
if reg292452 == True {
reg292453 := __e.Call(__defun__shen_4_5rule_6, V1407)
Parse__shen_4_5rule_6 := reg292453
_ = Parse__shen_4_5rule_6
reg292454 := __e.Call(__defun__fail)
reg292455 := PrimEqual(reg292454, Parse__shen_4_5rule_6)
reg292456 := PrimNot(reg292455)
if reg292456 == True {
reg292457 := PrimHead(Parse__shen_4_5rule_6)
reg292458 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg292459 := __e.Call(__defun__shen_4linearise, reg292458)
reg292460 := Nil;
reg292461 := PrimCons(reg292459, reg292460)
__ctx.TailApply(__defun__shen_4pair, reg292457, reg292461)
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
__initDefs = append(__initDefs, defType{name: "shen.<rules>", value: __defun__shen_4_5rules_6})

__defun__shen_4_5rule_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1409 := __args[0]
_ = V1409
reg292464 := __e.Call(__defun__shen_4_5patterns_6, V1409)
Parse__shen_4_5patterns_6 := reg292464
_ = Parse__shen_4_5patterns_6
reg292465 := __e.Call(__defun__fail)
reg292466 := PrimEqual(reg292465, Parse__shen_4_5patterns_6)
reg292467 := PrimNot(reg292466)
var reg292528 Obj
if reg292467 == True {
reg292468 := PrimHead(Parse__shen_4_5patterns_6)
reg292469 := PrimIsPair(reg292468)
var reg292478 Obj
if reg292469 == True {
reg292470 := MakeSymbol("->")
reg292471 := PrimHead(Parse__shen_4_5patterns_6)
reg292472 := PrimHead(reg292471)
reg292473 := PrimEqual(reg292470, reg292472)
var reg292476 Obj
if reg292473 == True {
reg292474 := True;
reg292476 = reg292474
} else {
reg292475 := False;
reg292476 = reg292475
}
reg292478 = reg292476
} else {
reg292477 := False;
reg292478 = reg292477
}
var reg292526 Obj
if reg292478 == True {
reg292479 := PrimHead(Parse__shen_4_5patterns_6)
reg292480 := PrimTail(reg292479)
reg292481 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292482 := __e.Call(__defun__shen_4pair, reg292480, reg292481)
reg292483 := __e.Call(__defun__shen_4_5action_6, reg292482)
Parse__shen_4_5action_6 := reg292483
_ = Parse__shen_4_5action_6
reg292484 := __e.Call(__defun__fail)
reg292485 := PrimEqual(reg292484, Parse__shen_4_5action_6)
reg292486 := PrimNot(reg292485)
var reg292524 Obj
if reg292486 == True {
reg292487 := PrimHead(Parse__shen_4_5action_6)
reg292488 := PrimIsPair(reg292487)
var reg292497 Obj
if reg292488 == True {
reg292489 := MakeSymbol("where")
reg292490 := PrimHead(Parse__shen_4_5action_6)
reg292491 := PrimHead(reg292490)
reg292492 := PrimEqual(reg292489, reg292491)
var reg292495 Obj
if reg292492 == True {
reg292493 := True;
reg292495 = reg292493
} else {
reg292494 := False;
reg292495 = reg292494
}
reg292497 = reg292495
} else {
reg292496 := False;
reg292497 = reg292496
}
var reg292522 Obj
if reg292497 == True {
reg292498 := PrimHead(Parse__shen_4_5action_6)
reg292499 := PrimTail(reg292498)
reg292500 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg292501 := __e.Call(__defun__shen_4pair, reg292499, reg292500)
reg292502 := __e.Call(__defun__shen_4_5guard_6, reg292501)
Parse__shen_4_5guard_6 := reg292502
_ = Parse__shen_4_5guard_6
reg292503 := __e.Call(__defun__fail)
reg292504 := PrimEqual(reg292503, Parse__shen_4_5guard_6)
reg292505 := PrimNot(reg292504)
var reg292520 Obj
if reg292505 == True {
reg292506 := PrimHead(Parse__shen_4_5guard_6)
reg292507 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292508 := MakeSymbol("where")
reg292509 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5guard_6)
reg292510 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg292511 := Nil;
reg292512 := PrimCons(reg292510, reg292511)
reg292513 := PrimCons(reg292509, reg292512)
reg292514 := PrimCons(reg292508, reg292513)
reg292515 := Nil;
reg292516 := PrimCons(reg292514, reg292515)
reg292517 := PrimCons(reg292507, reg292516)
reg292518 := __e.Call(__defun__shen_4pair, reg292506, reg292517)
reg292520 = reg292518
} else {
reg292519 := __e.Call(__defun__fail)
reg292520 = reg292519
}
reg292522 = reg292520
} else {
reg292521 := __e.Call(__defun__fail)
reg292522 = reg292521
}
reg292524 = reg292522
} else {
reg292523 := __e.Call(__defun__fail)
reg292524 = reg292523
}
reg292526 = reg292524
} else {
reg292525 := __e.Call(__defun__fail)
reg292526 = reg292525
}
reg292528 = reg292526
} else {
reg292527 := __e.Call(__defun__fail)
reg292528 = reg292527
}
YaccParse := reg292528
_ = YaccParse
reg292529 := __e.Call(__defun__fail)
reg292530 := PrimEqual(YaccParse, reg292529)
if reg292530 == True {
reg292531 := __e.Call(__defun__shen_4_5patterns_6, V1409)
Parse__shen_4_5patterns_6 := reg292531
_ = Parse__shen_4_5patterns_6
reg292532 := __e.Call(__defun__fail)
reg292533 := PrimEqual(reg292532, Parse__shen_4_5patterns_6)
reg292534 := PrimNot(reg292533)
var reg292566 Obj
if reg292534 == True {
reg292535 := PrimHead(Parse__shen_4_5patterns_6)
reg292536 := PrimIsPair(reg292535)
var reg292545 Obj
if reg292536 == True {
reg292537 := MakeSymbol("->")
reg292538 := PrimHead(Parse__shen_4_5patterns_6)
reg292539 := PrimHead(reg292538)
reg292540 := PrimEqual(reg292537, reg292539)
var reg292543 Obj
if reg292540 == True {
reg292541 := True;
reg292543 = reg292541
} else {
reg292542 := False;
reg292543 = reg292542
}
reg292545 = reg292543
} else {
reg292544 := False;
reg292545 = reg292544
}
var reg292564 Obj
if reg292545 == True {
reg292546 := PrimHead(Parse__shen_4_5patterns_6)
reg292547 := PrimTail(reg292546)
reg292548 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292549 := __e.Call(__defun__shen_4pair, reg292547, reg292548)
reg292550 := __e.Call(__defun__shen_4_5action_6, reg292549)
Parse__shen_4_5action_6 := reg292550
_ = Parse__shen_4_5action_6
reg292551 := __e.Call(__defun__fail)
reg292552 := PrimEqual(reg292551, Parse__shen_4_5action_6)
reg292553 := PrimNot(reg292552)
var reg292562 Obj
if reg292553 == True {
reg292554 := PrimHead(Parse__shen_4_5action_6)
reg292555 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292556 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg292557 := Nil;
reg292558 := PrimCons(reg292556, reg292557)
reg292559 := PrimCons(reg292555, reg292558)
reg292560 := __e.Call(__defun__shen_4pair, reg292554, reg292559)
reg292562 = reg292560
} else {
reg292561 := __e.Call(__defun__fail)
reg292562 = reg292561
}
reg292564 = reg292562
} else {
reg292563 := __e.Call(__defun__fail)
reg292564 = reg292563
}
reg292566 = reg292564
} else {
reg292565 := __e.Call(__defun__fail)
reg292566 = reg292565
}
YaccParse := reg292566
_ = YaccParse
reg292567 := __e.Call(__defun__fail)
reg292568 := PrimEqual(YaccParse, reg292567)
if reg292568 == True {
reg292569 := __e.Call(__defun__shen_4_5patterns_6, V1409)
Parse__shen_4_5patterns_6 := reg292569
_ = Parse__shen_4_5patterns_6
reg292570 := __e.Call(__defun__fail)
reg292571 := PrimEqual(reg292570, Parse__shen_4_5patterns_6)
reg292572 := PrimNot(reg292571)
var reg292637 Obj
if reg292572 == True {
reg292573 := PrimHead(Parse__shen_4_5patterns_6)
reg292574 := PrimIsPair(reg292573)
var reg292583 Obj
if reg292574 == True {
reg292575 := MakeSymbol("<-")
reg292576 := PrimHead(Parse__shen_4_5patterns_6)
reg292577 := PrimHead(reg292576)
reg292578 := PrimEqual(reg292575, reg292577)
var reg292581 Obj
if reg292578 == True {
reg292579 := True;
reg292581 = reg292579
} else {
reg292580 := False;
reg292581 = reg292580
}
reg292583 = reg292581
} else {
reg292582 := False;
reg292583 = reg292582
}
var reg292635 Obj
if reg292583 == True {
reg292584 := PrimHead(Parse__shen_4_5patterns_6)
reg292585 := PrimTail(reg292584)
reg292586 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292587 := __e.Call(__defun__shen_4pair, reg292585, reg292586)
reg292588 := __e.Call(__defun__shen_4_5action_6, reg292587)
Parse__shen_4_5action_6 := reg292588
_ = Parse__shen_4_5action_6
reg292589 := __e.Call(__defun__fail)
reg292590 := PrimEqual(reg292589, Parse__shen_4_5action_6)
reg292591 := PrimNot(reg292590)
var reg292633 Obj
if reg292591 == True {
reg292592 := PrimHead(Parse__shen_4_5action_6)
reg292593 := PrimIsPair(reg292592)
var reg292602 Obj
if reg292593 == True {
reg292594 := MakeSymbol("where")
reg292595 := PrimHead(Parse__shen_4_5action_6)
reg292596 := PrimHead(reg292595)
reg292597 := PrimEqual(reg292594, reg292596)
var reg292600 Obj
if reg292597 == True {
reg292598 := True;
reg292600 = reg292598
} else {
reg292599 := False;
reg292600 = reg292599
}
reg292602 = reg292600
} else {
reg292601 := False;
reg292602 = reg292601
}
var reg292631 Obj
if reg292602 == True {
reg292603 := PrimHead(Parse__shen_4_5action_6)
reg292604 := PrimTail(reg292603)
reg292605 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg292606 := __e.Call(__defun__shen_4pair, reg292604, reg292605)
reg292607 := __e.Call(__defun__shen_4_5guard_6, reg292606)
Parse__shen_4_5guard_6 := reg292607
_ = Parse__shen_4_5guard_6
reg292608 := __e.Call(__defun__fail)
reg292609 := PrimEqual(reg292608, Parse__shen_4_5guard_6)
reg292610 := PrimNot(reg292609)
var reg292629 Obj
if reg292610 == True {
reg292611 := PrimHead(Parse__shen_4_5guard_6)
reg292612 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292613 := MakeSymbol("where")
reg292614 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5guard_6)
reg292615 := MakeSymbol("shen.choicepoint!")
reg292616 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg292617 := Nil;
reg292618 := PrimCons(reg292616, reg292617)
reg292619 := PrimCons(reg292615, reg292618)
reg292620 := Nil;
reg292621 := PrimCons(reg292619, reg292620)
reg292622 := PrimCons(reg292614, reg292621)
reg292623 := PrimCons(reg292613, reg292622)
reg292624 := Nil;
reg292625 := PrimCons(reg292623, reg292624)
reg292626 := PrimCons(reg292612, reg292625)
reg292627 := __e.Call(__defun__shen_4pair, reg292611, reg292626)
reg292629 = reg292627
} else {
reg292628 := __e.Call(__defun__fail)
reg292629 = reg292628
}
reg292631 = reg292629
} else {
reg292630 := __e.Call(__defun__fail)
reg292631 = reg292630
}
reg292633 = reg292631
} else {
reg292632 := __e.Call(__defun__fail)
reg292633 = reg292632
}
reg292635 = reg292633
} else {
reg292634 := __e.Call(__defun__fail)
reg292635 = reg292634
}
reg292637 = reg292635
} else {
reg292636 := __e.Call(__defun__fail)
reg292637 = reg292636
}
YaccParse := reg292637
_ = YaccParse
reg292638 := __e.Call(__defun__fail)
reg292639 := PrimEqual(YaccParse, reg292638)
if reg292639 == True {
reg292640 := __e.Call(__defun__shen_4_5patterns_6, V1409)
Parse__shen_4_5patterns_6 := reg292640
_ = Parse__shen_4_5patterns_6
reg292641 := __e.Call(__defun__fail)
reg292642 := PrimEqual(reg292641, Parse__shen_4_5patterns_6)
reg292643 := PrimNot(reg292642)
if reg292643 == True {
reg292644 := PrimHead(Parse__shen_4_5patterns_6)
reg292645 := PrimIsPair(reg292644)
var reg292654 Obj
if reg292645 == True {
reg292646 := MakeSymbol("<-")
reg292647 := PrimHead(Parse__shen_4_5patterns_6)
reg292648 := PrimHead(reg292647)
reg292649 := PrimEqual(reg292646, reg292648)
var reg292652 Obj
if reg292649 == True {
reg292650 := True;
reg292652 = reg292650
} else {
reg292651 := False;
reg292652 = reg292651
}
reg292654 = reg292652
} else {
reg292653 := False;
reg292654 = reg292653
}
if reg292654 == True {
reg292655 := PrimHead(Parse__shen_4_5patterns_6)
reg292656 := PrimTail(reg292655)
reg292657 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292658 := __e.Call(__defun__shen_4pair, reg292656, reg292657)
reg292659 := __e.Call(__defun__shen_4_5action_6, reg292658)
Parse__shen_4_5action_6 := reg292659
_ = Parse__shen_4_5action_6
reg292660 := __e.Call(__defun__fail)
reg292661 := PrimEqual(reg292660, Parse__shen_4_5action_6)
reg292662 := PrimNot(reg292661)
if reg292662 == True {
reg292663 := PrimHead(Parse__shen_4_5action_6)
reg292664 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292665 := MakeSymbol("shen.choicepoint!")
reg292666 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg292667 := Nil;
reg292668 := PrimCons(reg292666, reg292667)
reg292669 := PrimCons(reg292665, reg292668)
reg292670 := Nil;
reg292671 := PrimCons(reg292669, reg292670)
reg292672 := PrimCons(reg292664, reg292671)
__ctx.TailApply(__defun__shen_4pair, reg292663, reg292672)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<rule>", value: __defun__shen_4_5rule_6})

__defun__shen_4fail__if = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1412 := __args[0]
_ = V1412
V1413 := __args[1]
_ = V1413
reg292677 := __e.Call(V1412, V1413)
if reg292677 == True {
__ctx.TailApply(__defun__fail)
return
} else {
__ctx.Return(V1413)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.fail_if", value: __defun__shen_4fail__if})

__defun__shen_4succeeds_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1419 := __args[0]
_ = V1419
reg292679 := __e.Call(__defun__fail)
reg292680 := PrimEqual(V1419, reg292679)
if reg292680 == True {
reg292681 := False;
__ctx.Return(reg292681)
return
} else {
reg292682 := True;
__ctx.Return(reg292682)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.succeeds?", value: __defun__shen_4succeeds_2})

__defun__shen_4_5patterns_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1421 := __args[0]
_ = V1421
reg292683 := __e.Call(__defun__shen_4_5pattern_6, V1421)
Parse__shen_4_5pattern_6 := reg292683
_ = Parse__shen_4_5pattern_6
reg292684 := __e.Call(__defun__fail)
reg292685 := PrimEqual(reg292684, Parse__shen_4_5pattern_6)
reg292686 := PrimNot(reg292685)
var reg292699 Obj
if reg292686 == True {
reg292687 := __e.Call(__defun__shen_4_5patterns_6, Parse__shen_4_5pattern_6)
Parse__shen_4_5patterns_6 := reg292687
_ = Parse__shen_4_5patterns_6
reg292688 := __e.Call(__defun__fail)
reg292689 := PrimEqual(reg292688, Parse__shen_4_5patterns_6)
reg292690 := PrimNot(reg292689)
var reg292697 Obj
if reg292690 == True {
reg292691 := PrimHead(Parse__shen_4_5patterns_6)
reg292692 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern_6)
reg292693 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg292694 := PrimCons(reg292692, reg292693)
reg292695 := __e.Call(__defun__shen_4pair, reg292691, reg292694)
reg292697 = reg292695
} else {
reg292696 := __e.Call(__defun__fail)
reg292697 = reg292696
}
reg292699 = reg292697
} else {
reg292698 := __e.Call(__defun__fail)
reg292699 = reg292698
}
YaccParse := reg292699
_ = YaccParse
reg292700 := __e.Call(__defun__fail)
reg292701 := PrimEqual(YaccParse, reg292700)
if reg292701 == True {
reg292702 := __e.Call(__defun___5e_6, V1421)
Parse___5e_6 := reg292702
_ = Parse___5e_6
reg292703 := __e.Call(__defun__fail)
reg292704 := PrimEqual(reg292703, Parse___5e_6)
reg292705 := PrimNot(reg292704)
if reg292705 == True {
reg292706 := PrimHead(Parse___5e_6)
reg292707 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg292706, reg292707)
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
__initDefs = append(__initDefs, defType{name: "shen.<patterns>", value: __defun__shen_4_5patterns_6})

__defun__shen_4_5pattern_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1428 := __args[0]
_ = V1428
reg292710 := PrimHead(V1428)
reg292711 := PrimIsPair(reg292710)
var reg292719 Obj
if reg292711 == True {
reg292712 := PrimHead(V1428)
reg292713 := PrimHead(reg292712)
reg292714 := PrimIsPair(reg292713)
var reg292717 Obj
if reg292714 == True {
reg292715 := True;
reg292717 = reg292715
} else {
reg292716 := False;
reg292717 = reg292716
}
reg292719 = reg292717
} else {
reg292718 := False;
reg292719 = reg292718
}
var reg292784 Obj
if reg292719 == True {
reg292720 := PrimHead(V1428)
reg292721 := PrimHead(reg292720)
reg292722 := PrimTail(V1428)
reg292723 := PrimHead(reg292722)
reg292724 := __e.Call(__defun__shen_4pair, reg292721, reg292723)
reg292725 := PrimHead(reg292724)
reg292726 := PrimIsPair(reg292725)
var reg292740 Obj
if reg292726 == True {
reg292727 := MakeSymbol("@p")
reg292728 := PrimHead(V1428)
reg292729 := PrimHead(reg292728)
reg292730 := PrimTail(V1428)
reg292731 := PrimHead(reg292730)
reg292732 := __e.Call(__defun__shen_4pair, reg292729, reg292731)
reg292733 := PrimHead(reg292732)
reg292734 := PrimHead(reg292733)
reg292735 := PrimEqual(reg292727, reg292734)
var reg292738 Obj
if reg292735 == True {
reg292736 := True;
reg292738 = reg292736
} else {
reg292737 := False;
reg292738 = reg292737
}
reg292740 = reg292738
} else {
reg292739 := False;
reg292740 = reg292739
}
var reg292782 Obj
if reg292740 == True {
reg292741 := PrimHead(V1428)
reg292742 := PrimHead(reg292741)
reg292743 := PrimTail(V1428)
reg292744 := PrimHead(reg292743)
reg292745 := __e.Call(__defun__shen_4pair, reg292742, reg292744)
reg292746 := PrimHead(reg292745)
reg292747 := PrimTail(reg292746)
reg292748 := PrimHead(V1428)
reg292749 := PrimHead(reg292748)
reg292750 := PrimTail(V1428)
reg292751 := PrimHead(reg292750)
reg292752 := __e.Call(__defun__shen_4pair, reg292749, reg292751)
reg292753 := __e.Call(__defun__shen_4hdtl, reg292752)
reg292754 := __e.Call(__defun__shen_4pair, reg292747, reg292753)
reg292755 := __e.Call(__defun__shen_4_5pattern1_6, reg292754)
Parse__shen_4_5pattern1_6 := reg292755
_ = Parse__shen_4_5pattern1_6
reg292756 := __e.Call(__defun__fail)
reg292757 := PrimEqual(reg292756, Parse__shen_4_5pattern1_6)
reg292758 := PrimNot(reg292757)
var reg292780 Obj
if reg292758 == True {
reg292759 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg292759
_ = Parse__shen_4_5pattern2_6
reg292760 := __e.Call(__defun__fail)
reg292761 := PrimEqual(reg292760, Parse__shen_4_5pattern2_6)
reg292762 := PrimNot(reg292761)
var reg292778 Obj
if reg292762 == True {
reg292763 := PrimHead(V1428)
reg292764 := PrimTail(reg292763)
reg292765 := PrimTail(V1428)
reg292766 := PrimHead(reg292765)
reg292767 := __e.Call(__defun__shen_4pair, reg292764, reg292766)
reg292768 := PrimHead(reg292767)
reg292769 := MakeSymbol("@p")
reg292770 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg292771 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg292772 := Nil;
reg292773 := PrimCons(reg292771, reg292772)
reg292774 := PrimCons(reg292770, reg292773)
reg292775 := PrimCons(reg292769, reg292774)
reg292776 := __e.Call(__defun__shen_4pair, reg292768, reg292775)
reg292778 = reg292776
} else {
reg292777 := __e.Call(__defun__fail)
reg292778 = reg292777
}
reg292780 = reg292778
} else {
reg292779 := __e.Call(__defun__fail)
reg292780 = reg292779
}
reg292782 = reg292780
} else {
reg292781 := __e.Call(__defun__fail)
reg292782 = reg292781
}
reg292784 = reg292782
} else {
reg292783 := __e.Call(__defun__fail)
reg292784 = reg292783
}
YaccParse := reg292784
_ = YaccParse
reg292785 := __e.Call(__defun__fail)
reg292786 := PrimEqual(YaccParse, reg292785)
if reg292786 == True {
reg292787 := PrimHead(V1428)
reg292788 := PrimIsPair(reg292787)
var reg292796 Obj
if reg292788 == True {
reg292789 := PrimHead(V1428)
reg292790 := PrimHead(reg292789)
reg292791 := PrimIsPair(reg292790)
var reg292794 Obj
if reg292791 == True {
reg292792 := True;
reg292794 = reg292792
} else {
reg292793 := False;
reg292794 = reg292793
}
reg292796 = reg292794
} else {
reg292795 := False;
reg292796 = reg292795
}
var reg292861 Obj
if reg292796 == True {
reg292797 := PrimHead(V1428)
reg292798 := PrimHead(reg292797)
reg292799 := PrimTail(V1428)
reg292800 := PrimHead(reg292799)
reg292801 := __e.Call(__defun__shen_4pair, reg292798, reg292800)
reg292802 := PrimHead(reg292801)
reg292803 := PrimIsPair(reg292802)
var reg292817 Obj
if reg292803 == True {
reg292804 := MakeSymbol("cons")
reg292805 := PrimHead(V1428)
reg292806 := PrimHead(reg292805)
reg292807 := PrimTail(V1428)
reg292808 := PrimHead(reg292807)
reg292809 := __e.Call(__defun__shen_4pair, reg292806, reg292808)
reg292810 := PrimHead(reg292809)
reg292811 := PrimHead(reg292810)
reg292812 := PrimEqual(reg292804, reg292811)
var reg292815 Obj
if reg292812 == True {
reg292813 := True;
reg292815 = reg292813
} else {
reg292814 := False;
reg292815 = reg292814
}
reg292817 = reg292815
} else {
reg292816 := False;
reg292817 = reg292816
}
var reg292859 Obj
if reg292817 == True {
reg292818 := PrimHead(V1428)
reg292819 := PrimHead(reg292818)
reg292820 := PrimTail(V1428)
reg292821 := PrimHead(reg292820)
reg292822 := __e.Call(__defun__shen_4pair, reg292819, reg292821)
reg292823 := PrimHead(reg292822)
reg292824 := PrimTail(reg292823)
reg292825 := PrimHead(V1428)
reg292826 := PrimHead(reg292825)
reg292827 := PrimTail(V1428)
reg292828 := PrimHead(reg292827)
reg292829 := __e.Call(__defun__shen_4pair, reg292826, reg292828)
reg292830 := __e.Call(__defun__shen_4hdtl, reg292829)
reg292831 := __e.Call(__defun__shen_4pair, reg292824, reg292830)
reg292832 := __e.Call(__defun__shen_4_5pattern1_6, reg292831)
Parse__shen_4_5pattern1_6 := reg292832
_ = Parse__shen_4_5pattern1_6
reg292833 := __e.Call(__defun__fail)
reg292834 := PrimEqual(reg292833, Parse__shen_4_5pattern1_6)
reg292835 := PrimNot(reg292834)
var reg292857 Obj
if reg292835 == True {
reg292836 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg292836
_ = Parse__shen_4_5pattern2_6
reg292837 := __e.Call(__defun__fail)
reg292838 := PrimEqual(reg292837, Parse__shen_4_5pattern2_6)
reg292839 := PrimNot(reg292838)
var reg292855 Obj
if reg292839 == True {
reg292840 := PrimHead(V1428)
reg292841 := PrimTail(reg292840)
reg292842 := PrimTail(V1428)
reg292843 := PrimHead(reg292842)
reg292844 := __e.Call(__defun__shen_4pair, reg292841, reg292843)
reg292845 := PrimHead(reg292844)
reg292846 := MakeSymbol("cons")
reg292847 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg292848 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg292849 := Nil;
reg292850 := PrimCons(reg292848, reg292849)
reg292851 := PrimCons(reg292847, reg292850)
reg292852 := PrimCons(reg292846, reg292851)
reg292853 := __e.Call(__defun__shen_4pair, reg292845, reg292852)
reg292855 = reg292853
} else {
reg292854 := __e.Call(__defun__fail)
reg292855 = reg292854
}
reg292857 = reg292855
} else {
reg292856 := __e.Call(__defun__fail)
reg292857 = reg292856
}
reg292859 = reg292857
} else {
reg292858 := __e.Call(__defun__fail)
reg292859 = reg292858
}
reg292861 = reg292859
} else {
reg292860 := __e.Call(__defun__fail)
reg292861 = reg292860
}
YaccParse := reg292861
_ = YaccParse
reg292862 := __e.Call(__defun__fail)
reg292863 := PrimEqual(YaccParse, reg292862)
if reg292863 == True {
reg292864 := PrimHead(V1428)
reg292865 := PrimIsPair(reg292864)
var reg292873 Obj
if reg292865 == True {
reg292866 := PrimHead(V1428)
reg292867 := PrimHead(reg292866)
reg292868 := PrimIsPair(reg292867)
var reg292871 Obj
if reg292868 == True {
reg292869 := True;
reg292871 = reg292869
} else {
reg292870 := False;
reg292871 = reg292870
}
reg292873 = reg292871
} else {
reg292872 := False;
reg292873 = reg292872
}
var reg292938 Obj
if reg292873 == True {
reg292874 := PrimHead(V1428)
reg292875 := PrimHead(reg292874)
reg292876 := PrimTail(V1428)
reg292877 := PrimHead(reg292876)
reg292878 := __e.Call(__defun__shen_4pair, reg292875, reg292877)
reg292879 := PrimHead(reg292878)
reg292880 := PrimIsPair(reg292879)
var reg292894 Obj
if reg292880 == True {
reg292881 := MakeSymbol("@v")
reg292882 := PrimHead(V1428)
reg292883 := PrimHead(reg292882)
reg292884 := PrimTail(V1428)
reg292885 := PrimHead(reg292884)
reg292886 := __e.Call(__defun__shen_4pair, reg292883, reg292885)
reg292887 := PrimHead(reg292886)
reg292888 := PrimHead(reg292887)
reg292889 := PrimEqual(reg292881, reg292888)
var reg292892 Obj
if reg292889 == True {
reg292890 := True;
reg292892 = reg292890
} else {
reg292891 := False;
reg292892 = reg292891
}
reg292894 = reg292892
} else {
reg292893 := False;
reg292894 = reg292893
}
var reg292936 Obj
if reg292894 == True {
reg292895 := PrimHead(V1428)
reg292896 := PrimHead(reg292895)
reg292897 := PrimTail(V1428)
reg292898 := PrimHead(reg292897)
reg292899 := __e.Call(__defun__shen_4pair, reg292896, reg292898)
reg292900 := PrimHead(reg292899)
reg292901 := PrimTail(reg292900)
reg292902 := PrimHead(V1428)
reg292903 := PrimHead(reg292902)
reg292904 := PrimTail(V1428)
reg292905 := PrimHead(reg292904)
reg292906 := __e.Call(__defun__shen_4pair, reg292903, reg292905)
reg292907 := __e.Call(__defun__shen_4hdtl, reg292906)
reg292908 := __e.Call(__defun__shen_4pair, reg292901, reg292907)
reg292909 := __e.Call(__defun__shen_4_5pattern1_6, reg292908)
Parse__shen_4_5pattern1_6 := reg292909
_ = Parse__shen_4_5pattern1_6
reg292910 := __e.Call(__defun__fail)
reg292911 := PrimEqual(reg292910, Parse__shen_4_5pattern1_6)
reg292912 := PrimNot(reg292911)
var reg292934 Obj
if reg292912 == True {
reg292913 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg292913
_ = Parse__shen_4_5pattern2_6
reg292914 := __e.Call(__defun__fail)
reg292915 := PrimEqual(reg292914, Parse__shen_4_5pattern2_6)
reg292916 := PrimNot(reg292915)
var reg292932 Obj
if reg292916 == True {
reg292917 := PrimHead(V1428)
reg292918 := PrimTail(reg292917)
reg292919 := PrimTail(V1428)
reg292920 := PrimHead(reg292919)
reg292921 := __e.Call(__defun__shen_4pair, reg292918, reg292920)
reg292922 := PrimHead(reg292921)
reg292923 := MakeSymbol("@v")
reg292924 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg292925 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg292926 := Nil;
reg292927 := PrimCons(reg292925, reg292926)
reg292928 := PrimCons(reg292924, reg292927)
reg292929 := PrimCons(reg292923, reg292928)
reg292930 := __e.Call(__defun__shen_4pair, reg292922, reg292929)
reg292932 = reg292930
} else {
reg292931 := __e.Call(__defun__fail)
reg292932 = reg292931
}
reg292934 = reg292932
} else {
reg292933 := __e.Call(__defun__fail)
reg292934 = reg292933
}
reg292936 = reg292934
} else {
reg292935 := __e.Call(__defun__fail)
reg292936 = reg292935
}
reg292938 = reg292936
} else {
reg292937 := __e.Call(__defun__fail)
reg292938 = reg292937
}
YaccParse := reg292938
_ = YaccParse
reg292939 := __e.Call(__defun__fail)
reg292940 := PrimEqual(YaccParse, reg292939)
if reg292940 == True {
reg292941 := PrimHead(V1428)
reg292942 := PrimIsPair(reg292941)
var reg292950 Obj
if reg292942 == True {
reg292943 := PrimHead(V1428)
reg292944 := PrimHead(reg292943)
reg292945 := PrimIsPair(reg292944)
var reg292948 Obj
if reg292945 == True {
reg292946 := True;
reg292948 = reg292946
} else {
reg292947 := False;
reg292948 = reg292947
}
reg292950 = reg292948
} else {
reg292949 := False;
reg292950 = reg292949
}
var reg293015 Obj
if reg292950 == True {
reg292951 := PrimHead(V1428)
reg292952 := PrimHead(reg292951)
reg292953 := PrimTail(V1428)
reg292954 := PrimHead(reg292953)
reg292955 := __e.Call(__defun__shen_4pair, reg292952, reg292954)
reg292956 := PrimHead(reg292955)
reg292957 := PrimIsPair(reg292956)
var reg292971 Obj
if reg292957 == True {
reg292958 := MakeSymbol("@s")
reg292959 := PrimHead(V1428)
reg292960 := PrimHead(reg292959)
reg292961 := PrimTail(V1428)
reg292962 := PrimHead(reg292961)
reg292963 := __e.Call(__defun__shen_4pair, reg292960, reg292962)
reg292964 := PrimHead(reg292963)
reg292965 := PrimHead(reg292964)
reg292966 := PrimEqual(reg292958, reg292965)
var reg292969 Obj
if reg292966 == True {
reg292967 := True;
reg292969 = reg292967
} else {
reg292968 := False;
reg292969 = reg292968
}
reg292971 = reg292969
} else {
reg292970 := False;
reg292971 = reg292970
}
var reg293013 Obj
if reg292971 == True {
reg292972 := PrimHead(V1428)
reg292973 := PrimHead(reg292972)
reg292974 := PrimTail(V1428)
reg292975 := PrimHead(reg292974)
reg292976 := __e.Call(__defun__shen_4pair, reg292973, reg292975)
reg292977 := PrimHead(reg292976)
reg292978 := PrimTail(reg292977)
reg292979 := PrimHead(V1428)
reg292980 := PrimHead(reg292979)
reg292981 := PrimTail(V1428)
reg292982 := PrimHead(reg292981)
reg292983 := __e.Call(__defun__shen_4pair, reg292980, reg292982)
reg292984 := __e.Call(__defun__shen_4hdtl, reg292983)
reg292985 := __e.Call(__defun__shen_4pair, reg292978, reg292984)
reg292986 := __e.Call(__defun__shen_4_5pattern1_6, reg292985)
Parse__shen_4_5pattern1_6 := reg292986
_ = Parse__shen_4_5pattern1_6
reg292987 := __e.Call(__defun__fail)
reg292988 := PrimEqual(reg292987, Parse__shen_4_5pattern1_6)
reg292989 := PrimNot(reg292988)
var reg293011 Obj
if reg292989 == True {
reg292990 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg292990
_ = Parse__shen_4_5pattern2_6
reg292991 := __e.Call(__defun__fail)
reg292992 := PrimEqual(reg292991, Parse__shen_4_5pattern2_6)
reg292993 := PrimNot(reg292992)
var reg293009 Obj
if reg292993 == True {
reg292994 := PrimHead(V1428)
reg292995 := PrimTail(reg292994)
reg292996 := PrimTail(V1428)
reg292997 := PrimHead(reg292996)
reg292998 := __e.Call(__defun__shen_4pair, reg292995, reg292997)
reg292999 := PrimHead(reg292998)
reg293000 := MakeSymbol("@s")
reg293001 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg293002 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg293003 := Nil;
reg293004 := PrimCons(reg293002, reg293003)
reg293005 := PrimCons(reg293001, reg293004)
reg293006 := PrimCons(reg293000, reg293005)
reg293007 := __e.Call(__defun__shen_4pair, reg292999, reg293006)
reg293009 = reg293007
} else {
reg293008 := __e.Call(__defun__fail)
reg293009 = reg293008
}
reg293011 = reg293009
} else {
reg293010 := __e.Call(__defun__fail)
reg293011 = reg293010
}
reg293013 = reg293011
} else {
reg293012 := __e.Call(__defun__fail)
reg293013 = reg293012
}
reg293015 = reg293013
} else {
reg293014 := __e.Call(__defun__fail)
reg293015 = reg293014
}
YaccParse := reg293015
_ = YaccParse
reg293016 := __e.Call(__defun__fail)
reg293017 := PrimEqual(YaccParse, reg293016)
if reg293017 == True {
reg293018 := PrimHead(V1428)
reg293019 := PrimIsPair(reg293018)
var reg293027 Obj
if reg293019 == True {
reg293020 := PrimHead(V1428)
reg293021 := PrimHead(reg293020)
reg293022 := PrimIsPair(reg293021)
var reg293025 Obj
if reg293022 == True {
reg293023 := True;
reg293025 = reg293023
} else {
reg293024 := False;
reg293025 = reg293024
}
reg293027 = reg293025
} else {
reg293026 := False;
reg293027 = reg293026
}
var reg293105 Obj
if reg293027 == True {
reg293028 := PrimHead(V1428)
reg293029 := PrimHead(reg293028)
reg293030 := PrimTail(V1428)
reg293031 := PrimHead(reg293030)
reg293032 := __e.Call(__defun__shen_4pair, reg293029, reg293031)
reg293033 := PrimHead(reg293032)
reg293034 := PrimIsPair(reg293033)
var reg293048 Obj
if reg293034 == True {
reg293035 := MakeSymbol("vector")
reg293036 := PrimHead(V1428)
reg293037 := PrimHead(reg293036)
reg293038 := PrimTail(V1428)
reg293039 := PrimHead(reg293038)
reg293040 := __e.Call(__defun__shen_4pair, reg293037, reg293039)
reg293041 := PrimHead(reg293040)
reg293042 := PrimHead(reg293041)
reg293043 := PrimEqual(reg293035, reg293042)
var reg293046 Obj
if reg293043 == True {
reg293044 := True;
reg293046 = reg293044
} else {
reg293045 := False;
reg293046 = reg293045
}
reg293048 = reg293046
} else {
reg293047 := False;
reg293048 = reg293047
}
var reg293103 Obj
if reg293048 == True {
reg293049 := PrimHead(V1428)
reg293050 := PrimHead(reg293049)
reg293051 := PrimTail(V1428)
reg293052 := PrimHead(reg293051)
reg293053 := __e.Call(__defun__shen_4pair, reg293050, reg293052)
reg293054 := PrimHead(reg293053)
reg293055 := PrimTail(reg293054)
reg293056 := PrimHead(V1428)
reg293057 := PrimHead(reg293056)
reg293058 := PrimTail(V1428)
reg293059 := PrimHead(reg293058)
reg293060 := __e.Call(__defun__shen_4pair, reg293057, reg293059)
reg293061 := __e.Call(__defun__shen_4hdtl, reg293060)
reg293062 := __e.Call(__defun__shen_4pair, reg293055, reg293061)
reg293063 := PrimHead(reg293062)
reg293064 := PrimIsPair(reg293063)
var reg293087 Obj
if reg293064 == True {
reg293065 := MakeNumber(0)
reg293066 := PrimHead(V1428)
reg293067 := PrimHead(reg293066)
reg293068 := PrimTail(V1428)
reg293069 := PrimHead(reg293068)
reg293070 := __e.Call(__defun__shen_4pair, reg293067, reg293069)
reg293071 := PrimHead(reg293070)
reg293072 := PrimTail(reg293071)
reg293073 := PrimHead(V1428)
reg293074 := PrimHead(reg293073)
reg293075 := PrimTail(V1428)
reg293076 := PrimHead(reg293075)
reg293077 := __e.Call(__defun__shen_4pair, reg293074, reg293076)
reg293078 := __e.Call(__defun__shen_4hdtl, reg293077)
reg293079 := __e.Call(__defun__shen_4pair, reg293072, reg293078)
reg293080 := PrimHead(reg293079)
reg293081 := PrimHead(reg293080)
reg293082 := PrimEqual(reg293065, reg293081)
var reg293085 Obj
if reg293082 == True {
reg293083 := True;
reg293085 = reg293083
} else {
reg293084 := False;
reg293085 = reg293084
}
reg293087 = reg293085
} else {
reg293086 := False;
reg293087 = reg293086
}
var reg293101 Obj
if reg293087 == True {
reg293088 := PrimHead(V1428)
reg293089 := PrimTail(reg293088)
reg293090 := PrimTail(V1428)
reg293091 := PrimHead(reg293090)
reg293092 := __e.Call(__defun__shen_4pair, reg293089, reg293091)
reg293093 := PrimHead(reg293092)
reg293094 := MakeSymbol("vector")
reg293095 := MakeNumber(0)
reg293096 := Nil;
reg293097 := PrimCons(reg293095, reg293096)
reg293098 := PrimCons(reg293094, reg293097)
reg293099 := __e.Call(__defun__shen_4pair, reg293093, reg293098)
reg293101 = reg293099
} else {
reg293100 := __e.Call(__defun__fail)
reg293101 = reg293100
}
reg293103 = reg293101
} else {
reg293102 := __e.Call(__defun__fail)
reg293103 = reg293102
}
reg293105 = reg293103
} else {
reg293104 := __e.Call(__defun__fail)
reg293105 = reg293104
}
YaccParse := reg293105
_ = YaccParse
reg293106 := __e.Call(__defun__fail)
reg293107 := PrimEqual(YaccParse, reg293106)
if reg293107 == True {
reg293108 := PrimHead(V1428)
reg293109 := PrimIsPair(reg293108)
var reg293123 Obj
if reg293109 == True {
reg293110 := PrimHead(V1428)
reg293111 := PrimHead(reg293110)
Parse__X := reg293111
_ = Parse__X
reg293112 := PrimIsPair(Parse__X)
var reg293121 Obj
if reg293112 == True {
reg293113 := PrimHead(V1428)
reg293114 := PrimTail(reg293113)
reg293115 := __e.Call(__defun__shen_4hdtl, V1428)
reg293116 := __e.Call(__defun__shen_4pair, reg293114, reg293115)
reg293117 := PrimHead(reg293116)
reg293118 := __e.Call(__defun__shen_4constructor_1error, Parse__X)
reg293119 := __e.Call(__defun__shen_4pair, reg293117, reg293118)
reg293121 = reg293119
} else {
reg293120 := __e.Call(__defun__fail)
reg293121 = reg293120
}
reg293123 = reg293121
} else {
reg293122 := __e.Call(__defun__fail)
reg293123 = reg293122
}
YaccParse := reg293123
_ = YaccParse
reg293124 := __e.Call(__defun__fail)
reg293125 := PrimEqual(YaccParse, reg293124)
if reg293125 == True {
reg293126 := __e.Call(__defun__shen_4_5simple__pattern_6, V1428)
Parse__shen_4_5simple__pattern_6 := reg293126
_ = Parse__shen_4_5simple__pattern_6
reg293127 := __e.Call(__defun__fail)
reg293128 := PrimEqual(reg293127, Parse__shen_4_5simple__pattern_6)
reg293129 := PrimNot(reg293128)
if reg293129 == True {
reg293130 := PrimHead(Parse__shen_4_5simple__pattern_6)
reg293131 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5simple__pattern_6)
__ctx.TailApply(__defun__shen_4pair, reg293130, reg293131)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<pattern>", value: __defun__shen_4_5pattern_6})

__defun__shen_4constructor_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1430 := __args[0]
_ = V1430
reg293134 := MakeString(" is not a legitimate constructor\n")
reg293135 := MakeSymbol("shen.a")
reg293136 := __e.Call(__defun__shen_4app, V1430, reg293134, reg293135)
reg293137 := PrimSimpleError(reg293136)
__ctx.Return(reg293137)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.constructor-error", value: __defun__shen_4constructor_1error})

__defun__shen_4_5simple__pattern_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1432 := __args[0]
_ = V1432
reg293138 := PrimHead(V1432)
reg293139 := PrimIsPair(reg293138)
var reg293155 Obj
if reg293139 == True {
reg293140 := PrimHead(V1432)
reg293141 := PrimHead(reg293140)
Parse__X := reg293141
_ = Parse__X
reg293142 := MakeSymbol("_")
reg293143 := PrimEqual(Parse__X, reg293142)
var reg293153 Obj
if reg293143 == True {
reg293144 := PrimHead(V1432)
reg293145 := PrimTail(reg293144)
reg293146 := __e.Call(__defun__shen_4hdtl, V1432)
reg293147 := __e.Call(__defun__shen_4pair, reg293145, reg293146)
reg293148 := PrimHead(reg293147)
reg293149 := MakeSymbol("Parse_Y")
reg293150 := __e.Call(__defun__gensym, reg293149)
reg293151 := __e.Call(__defun__shen_4pair, reg293148, reg293150)
reg293153 = reg293151
} else {
reg293152 := __e.Call(__defun__fail)
reg293153 = reg293152
}
reg293155 = reg293153
} else {
reg293154 := __e.Call(__defun__fail)
reg293155 = reg293154
}
YaccParse := reg293155
_ = YaccParse
reg293156 := __e.Call(__defun__fail)
reg293157 := PrimEqual(YaccParse, reg293156)
if reg293157 == True {
reg293158 := PrimHead(V1432)
reg293159 := PrimIsPair(reg293158)
if reg293159 == True {
reg293160 := PrimHead(V1432)
reg293161 := PrimHead(reg293160)
Parse__X := reg293161
_ = Parse__X
reg293162 := MakeSymbol("->")
reg293163 := MakeSymbol("<-")
reg293164 := Nil;
reg293165 := PrimCons(reg293163, reg293164)
reg293166 := PrimCons(reg293162, reg293165)
reg293167 := __e.Call(__defun__element_2, Parse__X, reg293166)
reg293168 := PrimNot(reg293167)
if reg293168 == True {
reg293169 := PrimHead(V1432)
reg293170 := PrimTail(reg293169)
reg293171 := __e.Call(__defun__shen_4hdtl, V1432)
reg293172 := __e.Call(__defun__shen_4pair, reg293170, reg293171)
reg293173 := PrimHead(reg293172)
__ctx.TailApply(__defun__shen_4pair, reg293173, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<simple_pattern>", value: __defun__shen_4_5simple__pattern_6})

__defun__shen_4_5pattern1_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1434 := __args[0]
_ = V1434
reg293177 := __e.Call(__defun__shen_4_5pattern_6, V1434)
Parse__shen_4_5pattern_6 := reg293177
_ = Parse__shen_4_5pattern_6
reg293178 := __e.Call(__defun__fail)
reg293179 := PrimEqual(reg293178, Parse__shen_4_5pattern_6)
reg293180 := PrimNot(reg293179)
if reg293180 == True {
reg293181 := PrimHead(Parse__shen_4_5pattern_6)
reg293182 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern_6)
__ctx.TailApply(__defun__shen_4pair, reg293181, reg293182)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<pattern1>", value: __defun__shen_4_5pattern1_6})

__defun__shen_4_5pattern2_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1436 := __args[0]
_ = V1436
reg293185 := __e.Call(__defun__shen_4_5pattern_6, V1436)
Parse__shen_4_5pattern_6 := reg293185
_ = Parse__shen_4_5pattern_6
reg293186 := __e.Call(__defun__fail)
reg293187 := PrimEqual(reg293186, Parse__shen_4_5pattern_6)
reg293188 := PrimNot(reg293187)
if reg293188 == True {
reg293189 := PrimHead(Parse__shen_4_5pattern_6)
reg293190 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern_6)
__ctx.TailApply(__defun__shen_4pair, reg293189, reg293190)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<pattern2>", value: __defun__shen_4_5pattern2_6})

__defun__shen_4_5action_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1438 := __args[0]
_ = V1438
reg293193 := PrimHead(V1438)
reg293194 := PrimIsPair(reg293193)
if reg293194 == True {
reg293195 := PrimHead(V1438)
reg293196 := PrimHead(reg293195)
Parse__X := reg293196
_ = Parse__X
reg293197 := PrimHead(V1438)
reg293198 := PrimTail(reg293197)
reg293199 := __e.Call(__defun__shen_4hdtl, V1438)
reg293200 := __e.Call(__defun__shen_4pair, reg293198, reg293199)
reg293201 := PrimHead(reg293200)
__ctx.TailApply(__defun__shen_4pair, reg293201, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<action>", value: __defun__shen_4_5action_6})

__defun__shen_4_5guard_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1440 := __args[0]
_ = V1440
reg293204 := PrimHead(V1440)
reg293205 := PrimIsPair(reg293204)
if reg293205 == True {
reg293206 := PrimHead(V1440)
reg293207 := PrimHead(reg293206)
Parse__X := reg293207
_ = Parse__X
reg293208 := PrimHead(V1440)
reg293209 := PrimTail(reg293208)
reg293210 := __e.Call(__defun__shen_4hdtl, V1440)
reg293211 := __e.Call(__defun__shen_4pair, reg293209, reg293210)
reg293212 := PrimHead(reg293211)
__ctx.TailApply(__defun__shen_4pair, reg293212, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<guard>", value: __defun__shen_4_5guard_6})

__defun__shen_4compile__to__machine__code = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1443 := __args[0]
_ = V1443
V1444 := __args[1]
_ = V1444
reg293215 := __e.Call(__defun__shen_4compile__to__lambda_7, V1443, V1444)
Lambda_7 := reg293215
_ = Lambda_7
reg293216 := __e.Call(__defun__shen_4compile__to__kl, V1443, Lambda_7)
KL := reg293216
_ = KL
reg293217 := __e.Call(__defun__shen_4record_1source, V1443, KL)
Record := reg293217
_ = Record
__ctx.Return(KL)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compile_to_machine_code", value: __defun__shen_4compile__to__machine__code})

__defun__shen_4record_1source = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1449 := __args[0]
_ = V1449
V1450 := __args[1]
_ = V1450
reg293218 := MakeSymbol("shen.*installing-kl*")
reg293219 := PrimValue(reg293218)
if reg293219 == True {
reg293220 := MakeSymbol("shen.skip")
__ctx.Return(reg293220)
return
} else {
reg293221 := MakeSymbol("shen.source")
reg293222 := MakeSymbol("*property-vector*")
reg293223 := PrimValue(reg293222)
__ctx.TailApply(__defun__put, V1449, reg293221, V1450, reg293223)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.record-source", value: __defun__shen_4record_1source})

__defun__shen_4compile__to__lambda_7 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1453 := __args[0]
_ = V1453
V1454 := __args[1]
_ = V1454
reg293225 := __e.Call(__defun__shen_4aritycheck, V1453, V1454)
Arity := reg293225
_ = Arity
reg293226 := __e.Call(__defun__shen_4update_1symbol_1table, V1453, Arity)
UpDateSymbolTable := reg293226
_ = UpDateSymbolTable
reg293227 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Rule := __args[0]
_ = Rule
__ctx.TailApply(__defun__shen_4free__variable__check, V1453, Rule)
return
}, 1)
reg293229 := __e.Call(__defun__shen_4for_1each, reg293227, V1454)
Free := reg293229
_ = Free
reg293230 := __e.Call(__defun__shen_4parameters, Arity)
Variables := reg293230
_ = Variables
reg293231 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4strip_1protect, X)
return
}, 1)
reg293233 := __e.Call(__defun__map, reg293231, V1454)
Strip := reg293233
_ = Strip
reg293234 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4abstract__rule, X)
return
}, 1)
reg293236 := __e.Call(__defun__map, reg293234, Strip)
Abstractions := reg293236
_ = Abstractions
reg293237 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4application__build, Variables, X)
return
}, 1)
reg293239 := __e.Call(__defun__map, reg293237, Abstractions)
Applications := reg293239
_ = Applications
reg293240 := Nil;
reg293241 := PrimCons(Applications, reg293240)
reg293242 := PrimCons(Variables, reg293241)
__ctx.Return(reg293242)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compile_to_lambda+", value: __defun__shen_4compile__to__lambda_7})

__defun__shen_4update_1symbol_1table = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1457 := __args[0]
_ = V1457
V1458 := __args[1]
_ = V1458
reg293243 := MakeNumber(0)
reg293244 := PrimEqual(reg293243, V1458)
if reg293244 == True {
reg293245 := MakeSymbol("shen.skip")
__ctx.Return(reg293245)
return
} else {
reg293246 := MakeSymbol("shen.lambda-form")
reg293247 := __e.Call(__defun__shen_4lambda_1form, V1457, V1458)
reg293248 := PrimEvalKL(__e, reg293247)
reg293249 := MakeSymbol("*property-vector*")
reg293250 := PrimValue(reg293249)
__ctx.TailApply(__defun__put, V1457, reg293246, reg293248, reg293250)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.update-symbol-table", value: __defun__shen_4update_1symbol_1table})

__defun__shen_4free__variable__check = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1461 := __args[0]
_ = V1461
V1462 := __args[1]
_ = V1462
reg293252 := PrimIsPair(V1462)
var reg293268 Obj
if reg293252 == True {
reg293253 := PrimTail(V1462)
reg293254 := PrimIsPair(reg293253)
var reg293263 Obj
if reg293254 == True {
reg293255 := Nil;
reg293256 := PrimTail(V1462)
reg293257 := PrimTail(reg293256)
reg293258 := PrimEqual(reg293255, reg293257)
var reg293261 Obj
if reg293258 == True {
reg293259 := True;
reg293261 = reg293259
} else {
reg293260 := False;
reg293261 = reg293260
}
reg293263 = reg293261
} else {
reg293262 := False;
reg293263 = reg293262
}
var reg293266 Obj
if reg293263 == True {
reg293264 := True;
reg293266 = reg293264
} else {
reg293265 := False;
reg293266 = reg293265
}
reg293268 = reg293266
} else {
reg293267 := False;
reg293268 = reg293267
}
if reg293268 == True {
reg293269 := PrimHead(V1462)
reg293270 := __e.Call(__defun__shen_4extract__vars, reg293269)
Bound := reg293270
_ = Bound
reg293271 := PrimTail(V1462)
reg293272 := PrimHead(reg293271)
reg293273 := __e.Call(__defun__shen_4extract__free__vars, Bound, reg293272)
Free := reg293273
_ = Free
__ctx.TailApply(__defun__shen_4free__variable__warnings, V1461, Free)
return
} else {
reg293275 := MakeSymbol("shen.free_variable_check")
__ctx.TailApply(__defun__shen_4f__error, reg293275)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.free_variable_check", value: __defun__shen_4free__variable__check})

__defun__shen_4extract__vars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1464 := __args[0]
_ = V1464
reg293277 := PrimIsVariable(V1464)
if reg293277 == True {
reg293278 := Nil;
reg293279 := PrimCons(V1464, reg293278)
__ctx.Return(reg293279)
return
} else {
reg293280 := PrimIsPair(V1464)
if reg293280 == True {
reg293281 := PrimHead(V1464)
reg293282 := __e.Call(__defun__shen_4extract__vars, reg293281)
reg293283 := PrimTail(V1464)
reg293284 := __e.Call(__defun__shen_4extract__vars, reg293283)
__ctx.TailApply(__defun__union, reg293282, reg293284)
return
} else {
reg293286 := Nil;
__ctx.Return(reg293286)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.extract_vars", value: __defun__shen_4extract__vars})

__defun__shen_4extract__free__vars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1476 := __args[0]
_ = V1476
V1477 := __args[1]
_ = V1477
reg293287 := PrimIsPair(V1477)
var reg293311 Obj
if reg293287 == True {
reg293288 := PrimTail(V1477)
reg293289 := PrimIsPair(reg293288)
var reg293306 Obj
if reg293289 == True {
reg293290 := Nil;
reg293291 := PrimTail(V1477)
reg293292 := PrimTail(reg293291)
reg293293 := PrimEqual(reg293290, reg293292)
var reg293301 Obj
if reg293293 == True {
reg293294 := PrimHead(V1477)
reg293295 := MakeSymbol("protect")
reg293296 := PrimEqual(reg293294, reg293295)
var reg293299 Obj
if reg293296 == True {
reg293297 := True;
reg293299 = reg293297
} else {
reg293298 := False;
reg293299 = reg293298
}
reg293301 = reg293299
} else {
reg293300 := False;
reg293301 = reg293300
}
var reg293304 Obj
if reg293301 == True {
reg293302 := True;
reg293304 = reg293302
} else {
reg293303 := False;
reg293304 = reg293303
}
reg293306 = reg293304
} else {
reg293305 := False;
reg293306 = reg293305
}
var reg293309 Obj
if reg293306 == True {
reg293307 := True;
reg293309 = reg293307
} else {
reg293308 := False;
reg293309 = reg293308
}
reg293311 = reg293309
} else {
reg293310 := False;
reg293311 = reg293310
}
if reg293311 == True {
reg293312 := Nil;
__ctx.Return(reg293312)
return
} else {
reg293313 := PrimIsVariable(V1477)
var reg293320 Obj
if reg293313 == True {
reg293314 := __e.Call(__defun__element_2, V1477, V1476)
reg293315 := PrimNot(reg293314)
var reg293318 Obj
if reg293315 == True {
reg293316 := True;
reg293318 = reg293316
} else {
reg293317 := False;
reg293318 = reg293317
}
reg293320 = reg293318
} else {
reg293319 := False;
reg293320 = reg293319
}
if reg293320 == True {
reg293321 := Nil;
reg293322 := PrimCons(V1477, reg293321)
__ctx.Return(reg293322)
return
} else {
reg293323 := PrimIsPair(V1477)
var reg293356 Obj
if reg293323 == True {
reg293324 := MakeSymbol("lambda")
reg293325 := PrimHead(V1477)
reg293326 := PrimEqual(reg293324, reg293325)
var reg293351 Obj
if reg293326 == True {
reg293327 := PrimTail(V1477)
reg293328 := PrimIsPair(reg293327)
var reg293346 Obj
if reg293328 == True {
reg293329 := PrimTail(V1477)
reg293330 := PrimTail(reg293329)
reg293331 := PrimIsPair(reg293330)
var reg293341 Obj
if reg293331 == True {
reg293332 := Nil;
reg293333 := PrimTail(V1477)
reg293334 := PrimTail(reg293333)
reg293335 := PrimTail(reg293334)
reg293336 := PrimEqual(reg293332, reg293335)
var reg293339 Obj
if reg293336 == True {
reg293337 := True;
reg293339 = reg293337
} else {
reg293338 := False;
reg293339 = reg293338
}
reg293341 = reg293339
} else {
reg293340 := False;
reg293341 = reg293340
}
var reg293344 Obj
if reg293341 == True {
reg293342 := True;
reg293344 = reg293342
} else {
reg293343 := False;
reg293344 = reg293343
}
reg293346 = reg293344
} else {
reg293345 := False;
reg293346 = reg293345
}
var reg293349 Obj
if reg293346 == True {
reg293347 := True;
reg293349 = reg293347
} else {
reg293348 := False;
reg293349 = reg293348
}
reg293351 = reg293349
} else {
reg293350 := False;
reg293351 = reg293350
}
var reg293354 Obj
if reg293351 == True {
reg293352 := True;
reg293354 = reg293352
} else {
reg293353 := False;
reg293354 = reg293353
}
reg293356 = reg293354
} else {
reg293355 := False;
reg293356 = reg293355
}
if reg293356 == True {
reg293357 := PrimTail(V1477)
reg293358 := PrimHead(reg293357)
reg293359 := PrimCons(reg293358, V1476)
reg293360 := PrimTail(V1477)
reg293361 := PrimTail(reg293360)
reg293362 := PrimHead(reg293361)
__ctx.TailApply(__defun__shen_4extract__free__vars, reg293359, reg293362)
return
} else {
reg293364 := PrimIsPair(V1477)
var reg293407 Obj
if reg293364 == True {
reg293365 := MakeSymbol("let")
reg293366 := PrimHead(V1477)
reg293367 := PrimEqual(reg293365, reg293366)
var reg293402 Obj
if reg293367 == True {
reg293368 := PrimTail(V1477)
reg293369 := PrimIsPair(reg293368)
var reg293397 Obj
if reg293369 == True {
reg293370 := PrimTail(V1477)
reg293371 := PrimTail(reg293370)
reg293372 := PrimIsPair(reg293371)
var reg293392 Obj
if reg293372 == True {
reg293373 := PrimTail(V1477)
reg293374 := PrimTail(reg293373)
reg293375 := PrimTail(reg293374)
reg293376 := PrimIsPair(reg293375)
var reg293387 Obj
if reg293376 == True {
reg293377 := Nil;
reg293378 := PrimTail(V1477)
reg293379 := PrimTail(reg293378)
reg293380 := PrimTail(reg293379)
reg293381 := PrimTail(reg293380)
reg293382 := PrimEqual(reg293377, reg293381)
var reg293385 Obj
if reg293382 == True {
reg293383 := True;
reg293385 = reg293383
} else {
reg293384 := False;
reg293385 = reg293384
}
reg293387 = reg293385
} else {
reg293386 := False;
reg293387 = reg293386
}
var reg293390 Obj
if reg293387 == True {
reg293388 := True;
reg293390 = reg293388
} else {
reg293389 := False;
reg293390 = reg293389
}
reg293392 = reg293390
} else {
reg293391 := False;
reg293392 = reg293391
}
var reg293395 Obj
if reg293392 == True {
reg293393 := True;
reg293395 = reg293393
} else {
reg293394 := False;
reg293395 = reg293394
}
reg293397 = reg293395
} else {
reg293396 := False;
reg293397 = reg293396
}
var reg293400 Obj
if reg293397 == True {
reg293398 := True;
reg293400 = reg293398
} else {
reg293399 := False;
reg293400 = reg293399
}
reg293402 = reg293400
} else {
reg293401 := False;
reg293402 = reg293401
}
var reg293405 Obj
if reg293402 == True {
reg293403 := True;
reg293405 = reg293403
} else {
reg293404 := False;
reg293405 = reg293404
}
reg293407 = reg293405
} else {
reg293406 := False;
reg293407 = reg293406
}
if reg293407 == True {
reg293408 := PrimTail(V1477)
reg293409 := PrimTail(reg293408)
reg293410 := PrimHead(reg293409)
reg293411 := __e.Call(__defun__shen_4extract__free__vars, V1476, reg293410)
reg293412 := PrimTail(V1477)
reg293413 := PrimHead(reg293412)
reg293414 := PrimCons(reg293413, V1476)
reg293415 := PrimTail(V1477)
reg293416 := PrimTail(reg293415)
reg293417 := PrimTail(reg293416)
reg293418 := PrimHead(reg293417)
reg293419 := __e.Call(__defun__shen_4extract__free__vars, reg293414, reg293418)
__ctx.TailApply(__defun__union, reg293411, reg293419)
return
} else {
reg293421 := PrimIsPair(V1477)
if reg293421 == True {
reg293422 := PrimHead(V1477)
reg293423 := __e.Call(__defun__shen_4extract__free__vars, V1476, reg293422)
reg293424 := PrimTail(V1477)
reg293425 := __e.Call(__defun__shen_4extract__free__vars, V1476, reg293424)
__ctx.TailApply(__defun__union, reg293423, reg293425)
return
} else {
reg293427 := Nil;
__ctx.Return(reg293427)
return
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.extract_free_vars", value: __defun__shen_4extract__free__vars})

__defun__shen_4free__variable__warnings = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1482 := __args[0]
_ = V1482
V1483 := __args[1]
_ = V1483
reg293428 := Nil;
reg293429 := PrimEqual(reg293428, V1483)
if reg293429 == True {
reg293430 := MakeSymbol("_")
__ctx.Return(reg293430)
return
} else {
reg293431 := MakeString("error: the following variables are free in ")
reg293432 := MakeString(": ")
reg293433 := __e.Call(__defun__shen_4list__variables, V1483)
reg293434 := MakeString("")
reg293435 := MakeSymbol("shen.a")
reg293436 := __e.Call(__defun__shen_4app, reg293433, reg293434, reg293435)
reg293437 := PrimStringConcat(reg293432, reg293436)
reg293438 := MakeSymbol("shen.a")
reg293439 := __e.Call(__defun__shen_4app, V1482, reg293437, reg293438)
reg293440 := PrimStringConcat(reg293431, reg293439)
reg293441 := PrimSimpleError(reg293440)
__ctx.Return(reg293441)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.free_variable_warnings", value: __defun__shen_4free__variable__warnings})

__defun__shen_4list__variables = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1485 := __args[0]
_ = V1485
reg293442 := PrimIsPair(V1485)
var reg293450 Obj
if reg293442 == True {
reg293443 := Nil;
reg293444 := PrimTail(V1485)
reg293445 := PrimEqual(reg293443, reg293444)
var reg293448 Obj
if reg293445 == True {
reg293446 := True;
reg293448 = reg293446
} else {
reg293447 := False;
reg293448 = reg293447
}
reg293450 = reg293448
} else {
reg293449 := False;
reg293450 = reg293449
}
if reg293450 == True {
reg293451 := PrimHead(V1485)
reg293452 := PrimStr(reg293451)
reg293453 := MakeString(".")
reg293454 := PrimStringConcat(reg293452, reg293453)
__ctx.Return(reg293454)
return
} else {
reg293455 := PrimIsPair(V1485)
if reg293455 == True {
reg293456 := PrimHead(V1485)
reg293457 := PrimStr(reg293456)
reg293458 := MakeString(", ")
reg293459 := PrimTail(V1485)
reg293460 := __e.Call(__defun__shen_4list__variables, reg293459)
reg293461 := PrimStringConcat(reg293458, reg293460)
reg293462 := PrimStringConcat(reg293457, reg293461)
__ctx.Return(reg293462)
return
} else {
reg293463 := MakeSymbol("shen.list_variables")
__ctx.TailApply(__defun__shen_4f__error, reg293463)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.list_variables", value: __defun__shen_4list__variables})

__defun__shen_4strip_1protect = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1487 := __args[0]
_ = V1487
reg293465 := PrimIsPair(V1487)
var reg293489 Obj
if reg293465 == True {
reg293466 := PrimTail(V1487)
reg293467 := PrimIsPair(reg293466)
var reg293484 Obj
if reg293467 == True {
reg293468 := Nil;
reg293469 := PrimTail(V1487)
reg293470 := PrimTail(reg293469)
reg293471 := PrimEqual(reg293468, reg293470)
var reg293479 Obj
if reg293471 == True {
reg293472 := PrimHead(V1487)
reg293473 := MakeSymbol("protect")
reg293474 := PrimEqual(reg293472, reg293473)
var reg293477 Obj
if reg293474 == True {
reg293475 := True;
reg293477 = reg293475
} else {
reg293476 := False;
reg293477 = reg293476
}
reg293479 = reg293477
} else {
reg293478 := False;
reg293479 = reg293478
}
var reg293482 Obj
if reg293479 == True {
reg293480 := True;
reg293482 = reg293480
} else {
reg293481 := False;
reg293482 = reg293481
}
reg293484 = reg293482
} else {
reg293483 := False;
reg293484 = reg293483
}
var reg293487 Obj
if reg293484 == True {
reg293485 := True;
reg293487 = reg293485
} else {
reg293486 := False;
reg293487 = reg293486
}
reg293489 = reg293487
} else {
reg293488 := False;
reg293489 = reg293488
}
if reg293489 == True {
reg293490 := PrimTail(V1487)
reg293491 := PrimHead(reg293490)
__ctx.TailApply(__defun__shen_4strip_1protect, reg293491)
return
} else {
reg293493 := PrimIsPair(V1487)
if reg293493 == True {
reg293494 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4strip_1protect, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg293494, V1487)
return
} else {
__ctx.Return(V1487)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.strip-protect", value: __defun__shen_4strip_1protect})

__defun__shen_4linearise = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1489 := __args[0]
_ = V1489
reg293497 := PrimIsPair(V1489)
var reg293513 Obj
if reg293497 == True {
reg293498 := PrimTail(V1489)
reg293499 := PrimIsPair(reg293498)
var reg293508 Obj
if reg293499 == True {
reg293500 := Nil;
reg293501 := PrimTail(V1489)
reg293502 := PrimTail(reg293501)
reg293503 := PrimEqual(reg293500, reg293502)
var reg293506 Obj
if reg293503 == True {
reg293504 := True;
reg293506 = reg293504
} else {
reg293505 := False;
reg293506 = reg293505
}
reg293508 = reg293506
} else {
reg293507 := False;
reg293508 = reg293507
}
var reg293511 Obj
if reg293508 == True {
reg293509 := True;
reg293511 = reg293509
} else {
reg293510 := False;
reg293511 = reg293510
}
reg293513 = reg293511
} else {
reg293512 := False;
reg293513 = reg293512
}
if reg293513 == True {
reg293514 := PrimHead(V1489)
reg293515 := __e.Call(__defun__shen_4flatten, reg293514)
reg293516 := PrimHead(V1489)
reg293517 := PrimTail(V1489)
reg293518 := PrimHead(reg293517)
__ctx.TailApply(__defun__shen_4linearise__help, reg293515, reg293516, reg293518)
return
} else {
reg293520 := MakeSymbol("shen.linearise")
__ctx.TailApply(__defun__shen_4f__error, reg293520)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.linearise", value: __defun__shen_4linearise})

__defun__shen_4flatten = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1491 := __args[0]
_ = V1491
reg293522 := Nil;
reg293523 := PrimEqual(reg293522, V1491)
if reg293523 == True {
reg293524 := Nil;
__ctx.Return(reg293524)
return
} else {
reg293525 := PrimIsPair(V1491)
if reg293525 == True {
reg293526 := PrimHead(V1491)
reg293527 := __e.Call(__defun__shen_4flatten, reg293526)
reg293528 := PrimTail(V1491)
reg293529 := __e.Call(__defun__shen_4flatten, reg293528)
__ctx.TailApply(__defun__append, reg293527, reg293529)
return
} else {
reg293531 := Nil;
reg293532 := PrimCons(V1491, reg293531)
__ctx.Return(reg293532)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.flatten", value: __defun__shen_4flatten})

__defun__shen_4linearise__help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1495 := __args[0]
_ = V1495
V1496 := __args[1]
_ = V1496
V1497 := __args[2]
_ = V1497
reg293533 := Nil;
reg293534 := PrimEqual(reg293533, V1495)
if reg293534 == True {
reg293535 := Nil;
reg293536 := PrimCons(V1497, reg293535)
reg293537 := PrimCons(V1496, reg293536)
__ctx.Return(reg293537)
return
} else {
reg293538 := PrimIsPair(V1495)
if reg293538 == True {
reg293539 := PrimHead(V1495)
reg293540 := PrimIsVariable(reg293539)
var reg293548 Obj
if reg293540 == True {
reg293541 := PrimHead(V1495)
reg293542 := PrimTail(V1495)
reg293543 := __e.Call(__defun__element_2, reg293541, reg293542)
var reg293546 Obj
if reg293543 == True {
reg293544 := True;
reg293546 = reg293544
} else {
reg293545 := False;
reg293546 = reg293545
}
reg293548 = reg293546
} else {
reg293547 := False;
reg293548 = reg293547
}
if reg293548 == True {
reg293549 := PrimHead(V1495)
reg293550 := __e.Call(__defun__gensym, reg293549)
Var := reg293550
_ = Var
reg293551 := MakeSymbol("where")
reg293552 := MakeSymbol("=")
reg293553 := PrimHead(V1495)
reg293554 := Nil;
reg293555 := PrimCons(Var, reg293554)
reg293556 := PrimCons(reg293553, reg293555)
reg293557 := PrimCons(reg293552, reg293556)
reg293558 := Nil;
reg293559 := PrimCons(V1497, reg293558)
reg293560 := PrimCons(reg293557, reg293559)
reg293561 := PrimCons(reg293551, reg293560)
NewAction := reg293561
_ = NewAction
reg293562 := PrimHead(V1495)
reg293563 := __e.Call(__defun__shen_4linearise__X, reg293562, Var, V1496)
NewPatts := reg293563
_ = NewPatts
reg293564 := PrimTail(V1495)
__ctx.TailApply(__defun__shen_4linearise__help, reg293564, NewPatts, NewAction)
return
} else {
reg293566 := PrimTail(V1495)
__ctx.TailApply(__defun__shen_4linearise__help, reg293566, V1496, V1497)
return
}
} else {
reg293568 := MakeSymbol("shen.linearise_help")
__ctx.TailApply(__defun__shen_4f__error, reg293568)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.linearise_help", value: __defun__shen_4linearise__help})

__defun__shen_4linearise__X = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1510 := __args[0]
_ = V1510
V1511 := __args[1]
_ = V1511
V1512 := __args[2]
_ = V1512
reg293570 := PrimEqual(V1512, V1510)
if reg293570 == True {
__ctx.Return(V1511)
return
} else {
reg293571 := PrimIsPair(V1512)
if reg293571 == True {
reg293572 := PrimHead(V1512)
reg293573 := __e.Call(__defun__shen_4linearise__X, V1510, V1511, reg293572)
L := reg293573
_ = L
reg293574 := PrimHead(V1512)
reg293575 := PrimEqual(L, reg293574)
if reg293575 == True {
reg293576 := PrimHead(V1512)
reg293577 := PrimTail(V1512)
reg293578 := __e.Call(__defun__shen_4linearise__X, V1510, V1511, reg293577)
reg293579 := PrimCons(reg293576, reg293578)
__ctx.Return(reg293579)
return
} else {
reg293580 := PrimTail(V1512)
reg293581 := PrimCons(L, reg293580)
__ctx.Return(reg293581)
return
}
} else {
__ctx.Return(V1512)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.linearise_X", value: __defun__shen_4linearise__X})

__defun__shen_4aritycheck = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1515 := __args[0]
_ = V1515
V1516 := __args[1]
_ = V1516
reg293582 := PrimIsPair(V1516)
var reg293615 Obj
if reg293582 == True {
reg293583 := PrimHead(V1516)
reg293584 := PrimIsPair(reg293583)
var reg293610 Obj
if reg293584 == True {
reg293585 := PrimHead(V1516)
reg293586 := PrimTail(reg293585)
reg293587 := PrimIsPair(reg293586)
var reg293605 Obj
if reg293587 == True {
reg293588 := Nil;
reg293589 := PrimHead(V1516)
reg293590 := PrimTail(reg293589)
reg293591 := PrimTail(reg293590)
reg293592 := PrimEqual(reg293588, reg293591)
var reg293600 Obj
if reg293592 == True {
reg293593 := Nil;
reg293594 := PrimTail(V1516)
reg293595 := PrimEqual(reg293593, reg293594)
var reg293598 Obj
if reg293595 == True {
reg293596 := True;
reg293598 = reg293596
} else {
reg293597 := False;
reg293598 = reg293597
}
reg293600 = reg293598
} else {
reg293599 := False;
reg293600 = reg293599
}
var reg293603 Obj
if reg293600 == True {
reg293601 := True;
reg293603 = reg293601
} else {
reg293602 := False;
reg293603 = reg293602
}
reg293605 = reg293603
} else {
reg293604 := False;
reg293605 = reg293604
}
var reg293608 Obj
if reg293605 == True {
reg293606 := True;
reg293608 = reg293606
} else {
reg293607 := False;
reg293608 = reg293607
}
reg293610 = reg293608
} else {
reg293609 := False;
reg293610 = reg293609
}
var reg293613 Obj
if reg293610 == True {
reg293611 := True;
reg293613 = reg293611
} else {
reg293612 := False;
reg293613 = reg293612
}
reg293615 = reg293613
} else {
reg293614 := False;
reg293615 = reg293614
}
if reg293615 == True {
reg293616 := PrimHead(V1516)
reg293617 := PrimTail(reg293616)
reg293618 := PrimHead(reg293617)
reg293619 := __e.Call(__defun__shen_4aritycheck_1action, reg293618)
_ = reg293619
reg293620 := __e.Call(__defun__arity, V1515)
reg293621 := PrimHead(V1516)
reg293622 := PrimHead(reg293621)
reg293623 := __e.Call(__defun__length, reg293622)
__ctx.TailApply(__defun__shen_4aritycheck_1name, V1515, reg293620, reg293623)
return
} else {
reg293625 := PrimIsPair(V1516)
var reg293685 Obj
if reg293625 == True {
reg293626 := PrimHead(V1516)
reg293627 := PrimIsPair(reg293626)
var reg293680 Obj
if reg293627 == True {
reg293628 := PrimHead(V1516)
reg293629 := PrimTail(reg293628)
reg293630 := PrimIsPair(reg293629)
var reg293675 Obj
if reg293630 == True {
reg293631 := Nil;
reg293632 := PrimHead(V1516)
reg293633 := PrimTail(reg293632)
reg293634 := PrimTail(reg293633)
reg293635 := PrimEqual(reg293631, reg293634)
var reg293670 Obj
if reg293635 == True {
reg293636 := PrimTail(V1516)
reg293637 := PrimIsPair(reg293636)
var reg293665 Obj
if reg293637 == True {
reg293638 := PrimTail(V1516)
reg293639 := PrimHead(reg293638)
reg293640 := PrimIsPair(reg293639)
var reg293660 Obj
if reg293640 == True {
reg293641 := PrimTail(V1516)
reg293642 := PrimHead(reg293641)
reg293643 := PrimTail(reg293642)
reg293644 := PrimIsPair(reg293643)
var reg293655 Obj
if reg293644 == True {
reg293645 := Nil;
reg293646 := PrimTail(V1516)
reg293647 := PrimHead(reg293646)
reg293648 := PrimTail(reg293647)
reg293649 := PrimTail(reg293648)
reg293650 := PrimEqual(reg293645, reg293649)
var reg293653 Obj
if reg293650 == True {
reg293651 := True;
reg293653 = reg293651
} else {
reg293652 := False;
reg293653 = reg293652
}
reg293655 = reg293653
} else {
reg293654 := False;
reg293655 = reg293654
}
var reg293658 Obj
if reg293655 == True {
reg293656 := True;
reg293658 = reg293656
} else {
reg293657 := False;
reg293658 = reg293657
}
reg293660 = reg293658
} else {
reg293659 := False;
reg293660 = reg293659
}
var reg293663 Obj
if reg293660 == True {
reg293661 := True;
reg293663 = reg293661
} else {
reg293662 := False;
reg293663 = reg293662
}
reg293665 = reg293663
} else {
reg293664 := False;
reg293665 = reg293664
}
var reg293668 Obj
if reg293665 == True {
reg293666 := True;
reg293668 = reg293666
} else {
reg293667 := False;
reg293668 = reg293667
}
reg293670 = reg293668
} else {
reg293669 := False;
reg293670 = reg293669
}
var reg293673 Obj
if reg293670 == True {
reg293671 := True;
reg293673 = reg293671
} else {
reg293672 := False;
reg293673 = reg293672
}
reg293675 = reg293673
} else {
reg293674 := False;
reg293675 = reg293674
}
var reg293678 Obj
if reg293675 == True {
reg293676 := True;
reg293678 = reg293676
} else {
reg293677 := False;
reg293678 = reg293677
}
reg293680 = reg293678
} else {
reg293679 := False;
reg293680 = reg293679
}
var reg293683 Obj
if reg293680 == True {
reg293681 := True;
reg293683 = reg293681
} else {
reg293682 := False;
reg293683 = reg293682
}
reg293685 = reg293683
} else {
reg293684 := False;
reg293685 = reg293684
}
if reg293685 == True {
reg293686 := PrimHead(V1516)
reg293687 := PrimHead(reg293686)
reg293688 := __e.Call(__defun__length, reg293687)
reg293689 := PrimTail(V1516)
reg293690 := PrimHead(reg293689)
reg293691 := PrimHead(reg293690)
reg293692 := __e.Call(__defun__length, reg293691)
reg293693 := PrimEqual(reg293688, reg293692)
if reg293693 == True {
reg293694 := PrimHead(V1516)
reg293695 := PrimTail(reg293694)
reg293696 := PrimHead(reg293695)
reg293697 := __e.Call(__defun__shen_4aritycheck_1action, reg293696)
_ = reg293697
reg293698 := PrimTail(V1516)
__ctx.TailApply(__defun__shen_4aritycheck, V1515, reg293698)
return
} else {
reg293700 := MakeString("arity error in ")
reg293701 := MakeString("\n")
reg293702 := MakeSymbol("shen.a")
reg293703 := __e.Call(__defun__shen_4app, V1515, reg293701, reg293702)
reg293704 := PrimStringConcat(reg293700, reg293703)
reg293705 := PrimSimpleError(reg293704)
__ctx.Return(reg293705)
return
}
} else {
reg293706 := MakeSymbol("shen.aritycheck")
__ctx.TailApply(__defun__shen_4f__error, reg293706)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.aritycheck", value: __defun__shen_4aritycheck})

__defun__shen_4aritycheck_1name = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1529 := __args[0]
_ = V1529
V1530 := __args[1]
_ = V1530
V1531 := __args[2]
_ = V1531
reg293708 := MakeNumber(-1)
reg293709 := PrimEqual(reg293708, V1530)
if reg293709 == True {
__ctx.Return(V1531)
return
} else {
reg293710 := PrimEqual(V1531, V1530)
if reg293710 == True {
__ctx.Return(V1531)
return
} else {
reg293711 := MakeString("\nwarning: changing the arity of ")
reg293712 := MakeString(" can cause errors.\n")
reg293713 := MakeSymbol("shen.a")
reg293714 := __e.Call(__defun__shen_4app, V1529, reg293712, reg293713)
reg293715 := PrimStringConcat(reg293711, reg293714)
reg293716 := __e.Call(__defun__stoutput)
reg293717 := __e.Call(__defun__shen_4prhush, reg293715, reg293716)
_ = reg293717
__ctx.Return(V1531)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.aritycheck-name", value: __defun__shen_4aritycheck_1name})

__defun__shen_4aritycheck_1action = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1537 := __args[0]
_ = V1537
reg293718 := PrimIsPair(V1537)
if reg293718 == True {
reg293719 := PrimHead(V1537)
reg293720 := PrimTail(V1537)
reg293721 := __e.Call(__defun__shen_4aah, reg293719, reg293720)
_ = reg293721
reg293722 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4aritycheck_1action, Y)
return
}, 1)
__ctx.TailApply(__defun__shen_4for_1each, reg293722, V1537)
return
} else {
reg293725 := MakeSymbol("shen.skip")
__ctx.Return(reg293725)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.aritycheck-action", value: __defun__shen_4aritycheck_1action})

__defun__shen_4aah = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1540 := __args[0]
_ = V1540
V1541 := __args[1]
_ = V1541
reg293726 := __e.Call(__defun__arity, V1540)
Arity := reg293726
_ = Arity
reg293727 := __e.Call(__defun__length, V1541)
Len := reg293727
_ = Len
reg293728 := MakeNumber(-1)
reg293729 := PrimGreatThan(Arity, reg293728)
var reg293735 Obj
if reg293729 == True {
reg293730 := PrimGreatThan(Len, Arity)
var reg293733 Obj
if reg293730 == True {
reg293731 := True;
reg293733 = reg293731
} else {
reg293732 := False;
reg293733 = reg293732
}
reg293735 = reg293733
} else {
reg293734 := False;
reg293735 = reg293734
}
if reg293735 == True {
reg293736 := MakeString("warning: ")
reg293737 := MakeString(" might not like ")
reg293738 := MakeString(" argument")
reg293739 := MakeNumber(1)
reg293740 := PrimGreatThan(Len, reg293739)
var reg293743 Obj
if reg293740 == True {
reg293741 := MakeString("s")
reg293743 = reg293741
} else {
reg293742 := MakeString("")
reg293743 = reg293742
}
reg293744 := MakeString(".\n")
reg293745 := MakeSymbol("shen.a")
reg293746 := __e.Call(__defun__shen_4app, reg293743, reg293744, reg293745)
reg293747 := PrimStringConcat(reg293738, reg293746)
reg293748 := MakeSymbol("shen.a")
reg293749 := __e.Call(__defun__shen_4app, Len, reg293747, reg293748)
reg293750 := PrimStringConcat(reg293737, reg293749)
reg293751 := MakeSymbol("shen.a")
reg293752 := __e.Call(__defun__shen_4app, V1540, reg293750, reg293751)
reg293753 := PrimStringConcat(reg293736, reg293752)
reg293754 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg293753, reg293754)
return
} else {
reg293756 := MakeSymbol("shen.skip")
__ctx.Return(reg293756)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.aah", value: __defun__shen_4aah})

__defun__shen_4abstract__rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1543 := __args[0]
_ = V1543
reg293757 := PrimIsPair(V1543)
var reg293773 Obj
if reg293757 == True {
reg293758 := PrimTail(V1543)
reg293759 := PrimIsPair(reg293758)
var reg293768 Obj
if reg293759 == True {
reg293760 := Nil;
reg293761 := PrimTail(V1543)
reg293762 := PrimTail(reg293761)
reg293763 := PrimEqual(reg293760, reg293762)
var reg293766 Obj
if reg293763 == True {
reg293764 := True;
reg293766 = reg293764
} else {
reg293765 := False;
reg293766 = reg293765
}
reg293768 = reg293766
} else {
reg293767 := False;
reg293768 = reg293767
}
var reg293771 Obj
if reg293768 == True {
reg293769 := True;
reg293771 = reg293769
} else {
reg293770 := False;
reg293771 = reg293770
}
reg293773 = reg293771
} else {
reg293772 := False;
reg293773 = reg293772
}
if reg293773 == True {
reg293774 := PrimHead(V1543)
reg293775 := PrimTail(V1543)
reg293776 := PrimHead(reg293775)
__ctx.TailApply(__defun__shen_4abstraction__build, reg293774, reg293776)
return
} else {
reg293778 := MakeSymbol("shen.abstract_rule")
__ctx.TailApply(__defun__shen_4f__error, reg293778)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.abstract_rule", value: __defun__shen_4abstract__rule})

__defun__shen_4abstraction__build = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1546 := __args[0]
_ = V1546
V1547 := __args[1]
_ = V1547
reg293780 := Nil;
reg293781 := PrimEqual(reg293780, V1546)
if reg293781 == True {
__ctx.Return(V1547)
return
} else {
reg293782 := PrimIsPair(V1546)
if reg293782 == True {
reg293783 := MakeSymbol("/.")
reg293784 := PrimHead(V1546)
reg293785 := PrimTail(V1546)
reg293786 := __e.Call(__defun__shen_4abstraction__build, reg293785, V1547)
reg293787 := Nil;
reg293788 := PrimCons(reg293786, reg293787)
reg293789 := PrimCons(reg293784, reg293788)
reg293790 := PrimCons(reg293783, reg293789)
__ctx.Return(reg293790)
return
} else {
reg293791 := MakeSymbol("shen.abstraction_build")
__ctx.TailApply(__defun__shen_4f__error, reg293791)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.abstraction_build", value: __defun__shen_4abstraction__build})

__defun__shen_4parameters = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1549 := __args[0]
_ = V1549
reg293793 := MakeNumber(0)
reg293794 := PrimEqual(reg293793, V1549)
if reg293794 == True {
reg293795 := Nil;
__ctx.Return(reg293795)
return
} else {
reg293796 := MakeSymbol("V")
reg293797 := __e.Call(__defun__gensym, reg293796)
reg293798 := MakeNumber(1)
reg293799 := PrimNumberSubtract(V1549, reg293798)
reg293800 := __e.Call(__defun__shen_4parameters, reg293799)
reg293801 := PrimCons(reg293797, reg293800)
__ctx.Return(reg293801)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.parameters", value: __defun__shen_4parameters})

__defun__shen_4application__build = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1552 := __args[0]
_ = V1552
V1553 := __args[1]
_ = V1553
reg293802 := Nil;
reg293803 := PrimEqual(reg293802, V1552)
if reg293803 == True {
__ctx.Return(V1553)
return
} else {
reg293804 := PrimIsPair(V1552)
if reg293804 == True {
reg293805 := PrimTail(V1552)
reg293806 := PrimHead(V1552)
reg293807 := Nil;
reg293808 := PrimCons(reg293806, reg293807)
reg293809 := PrimCons(V1553, reg293808)
__ctx.TailApply(__defun__shen_4application__build, reg293805, reg293809)
return
} else {
reg293811 := MakeSymbol("shen.application_build")
__ctx.TailApply(__defun__shen_4f__error, reg293811)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.application_build", value: __defun__shen_4application__build})

__defun__shen_4compile__to__kl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1556 := __args[0]
_ = V1556
V1557 := __args[1]
_ = V1557
reg293813 := PrimIsPair(V1557)
var reg293829 Obj
if reg293813 == True {
reg293814 := PrimTail(V1557)
reg293815 := PrimIsPair(reg293814)
var reg293824 Obj
if reg293815 == True {
reg293816 := Nil;
reg293817 := PrimTail(V1557)
reg293818 := PrimTail(reg293817)
reg293819 := PrimEqual(reg293816, reg293818)
var reg293822 Obj
if reg293819 == True {
reg293820 := True;
reg293822 = reg293820
} else {
reg293821 := False;
reg293822 = reg293821
}
reg293824 = reg293822
} else {
reg293823 := False;
reg293824 = reg293823
}
var reg293827 Obj
if reg293824 == True {
reg293825 := True;
reg293827 = reg293825
} else {
reg293826 := False;
reg293827 = reg293826
}
reg293829 = reg293827
} else {
reg293828 := False;
reg293829 = reg293828
}
if reg293829 == True {
reg293830 := PrimHead(V1557)
reg293831 := __e.Call(__defun__length, reg293830)
reg293832 := __e.Call(__defun__shen_4store_1arity, V1556, reg293831)
Arity := reg293832
_ = Arity
reg293833 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4reduce, X)
return
}, 1)
reg293835 := PrimTail(V1557)
reg293836 := PrimHead(reg293835)
reg293837 := __e.Call(__defun__map, reg293833, reg293836)
Reduce := reg293837
_ = Reduce
reg293838 := PrimHead(V1557)
reg293839 := __e.Call(__defun__shen_4cond_1expression, V1556, reg293838, Reduce)
CondExpression := reg293839
_ = CondExpression
reg293840 := MakeSymbol("shen.*optimise*")
reg293841 := PrimValue(reg293840)
var reg293846 Obj
if reg293841 == True {
reg293842 := __e.Call(__defun__shen_4get_1type, V1556)
reg293843 := PrimHead(V1557)
reg293844 := __e.Call(__defun__shen_4typextable, reg293842, reg293843)
reg293846 = reg293844
} else {
reg293845 := MakeSymbol("shen.skip")
reg293846 = reg293845
}
TypeTable := reg293846
_ = TypeTable
reg293847 := MakeSymbol("shen.*optimise*")
reg293848 := PrimValue(reg293847)
var reg293851 Obj
if reg293848 == True {
reg293849 := PrimHead(V1557)
reg293850 := __e.Call(__defun__shen_4assign_1types, reg293849, TypeTable, CondExpression)
reg293851 = reg293850
} else {
reg293851 = CondExpression
}
TypedCondExpression := reg293851
_ = TypedCondExpression
reg293852 := MakeSymbol("defun")
reg293853 := PrimHead(V1557)
reg293854 := Nil;
reg293855 := PrimCons(TypedCondExpression, reg293854)
reg293856 := PrimCons(reg293853, reg293855)
reg293857 := PrimCons(V1556, reg293856)
reg293858 := PrimCons(reg293852, reg293857)
__ctx.Return(reg293858)
return
} else {
reg293859 := MakeSymbol("shen.compile_to_kl")
__ctx.TailApply(__defun__shen_4f__error, reg293859)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compile_to_kl", value: __defun__shen_4compile__to__kl})

__defun__shen_4get_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1563 := __args[0]
_ = V1563
reg293861 := PrimIsPair(V1563)
if reg293861 == True {
reg293862 := MakeSymbol("shen.skip")
__ctx.Return(reg293862)
return
} else {
reg293863 := MakeSymbol("shen.*signedfuncs*")
reg293864 := PrimValue(reg293863)
reg293865 := __e.Call(__defun__assoc, V1563, reg293864)
FType := reg293865
_ = FType
reg293866 := __e.Call(__defun__empty_2, FType)
if reg293866 == True {
reg293867 := MakeSymbol("shen.skip")
__ctx.Return(reg293867)
return
} else {
reg293868 := PrimTail(FType)
__ctx.Return(reg293868)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.get-type", value: __defun__shen_4get_1type})

__defun__shen_4typextable = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1574 := __args[0]
_ = V1574
V1575 := __args[1]
_ = V1575
reg293869 := PrimIsPair(V1574)
var reg293909 Obj
if reg293869 == True {
reg293870 := PrimTail(V1574)
reg293871 := PrimIsPair(reg293870)
var reg293904 Obj
if reg293871 == True {
reg293872 := MakeSymbol("-->")
reg293873 := PrimTail(V1574)
reg293874 := PrimHead(reg293873)
reg293875 := PrimEqual(reg293872, reg293874)
var reg293899 Obj
if reg293875 == True {
reg293876 := PrimTail(V1574)
reg293877 := PrimTail(reg293876)
reg293878 := PrimIsPair(reg293877)
var reg293894 Obj
if reg293878 == True {
reg293879 := Nil;
reg293880 := PrimTail(V1574)
reg293881 := PrimTail(reg293880)
reg293882 := PrimTail(reg293881)
reg293883 := PrimEqual(reg293879, reg293882)
var reg293889 Obj
if reg293883 == True {
reg293884 := PrimIsPair(V1575)
var reg293887 Obj
if reg293884 == True {
reg293885 := True;
reg293887 = reg293885
} else {
reg293886 := False;
reg293887 = reg293886
}
reg293889 = reg293887
} else {
reg293888 := False;
reg293889 = reg293888
}
var reg293892 Obj
if reg293889 == True {
reg293890 := True;
reg293892 = reg293890
} else {
reg293891 := False;
reg293892 = reg293891
}
reg293894 = reg293892
} else {
reg293893 := False;
reg293894 = reg293893
}
var reg293897 Obj
if reg293894 == True {
reg293895 := True;
reg293897 = reg293895
} else {
reg293896 := False;
reg293897 = reg293896
}
reg293899 = reg293897
} else {
reg293898 := False;
reg293899 = reg293898
}
var reg293902 Obj
if reg293899 == True {
reg293900 := True;
reg293902 = reg293900
} else {
reg293901 := False;
reg293902 = reg293901
}
reg293904 = reg293902
} else {
reg293903 := False;
reg293904 = reg293903
}
var reg293907 Obj
if reg293904 == True {
reg293905 := True;
reg293907 = reg293905
} else {
reg293906 := False;
reg293907 = reg293906
}
reg293909 = reg293907
} else {
reg293908 := False;
reg293909 = reg293908
}
if reg293909 == True {
reg293910 := PrimHead(V1574)
reg293911 := PrimIsVariable(reg293910)
if reg293911 == True {
reg293912 := PrimTail(V1574)
reg293913 := PrimTail(reg293912)
reg293914 := PrimHead(reg293913)
reg293915 := PrimTail(V1575)
__ctx.TailApply(__defun__shen_4typextable, reg293914, reg293915)
return
} else {
reg293917 := PrimHead(V1575)
reg293918 := PrimHead(V1574)
reg293919 := PrimCons(reg293917, reg293918)
reg293920 := PrimTail(V1574)
reg293921 := PrimTail(reg293920)
reg293922 := PrimHead(reg293921)
reg293923 := PrimTail(V1575)
reg293924 := __e.Call(__defun__shen_4typextable, reg293922, reg293923)
reg293925 := PrimCons(reg293919, reg293924)
__ctx.Return(reg293925)
return
}
} else {
reg293926 := Nil;
__ctx.Return(reg293926)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.typextable", value: __defun__shen_4typextable})

__defun__shen_4assign_1types = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1579 := __args[0]
_ = V1579
V1580 := __args[1]
_ = V1580
V1581 := __args[2]
_ = V1581
reg293927 := PrimIsPair(V1581)
var reg293970 Obj
if reg293927 == True {
reg293928 := MakeSymbol("let")
reg293929 := PrimHead(V1581)
reg293930 := PrimEqual(reg293928, reg293929)
var reg293965 Obj
if reg293930 == True {
reg293931 := PrimTail(V1581)
reg293932 := PrimIsPair(reg293931)
var reg293960 Obj
if reg293932 == True {
reg293933 := PrimTail(V1581)
reg293934 := PrimTail(reg293933)
reg293935 := PrimIsPair(reg293934)
var reg293955 Obj
if reg293935 == True {
reg293936 := PrimTail(V1581)
reg293937 := PrimTail(reg293936)
reg293938 := PrimTail(reg293937)
reg293939 := PrimIsPair(reg293938)
var reg293950 Obj
if reg293939 == True {
reg293940 := Nil;
reg293941 := PrimTail(V1581)
reg293942 := PrimTail(reg293941)
reg293943 := PrimTail(reg293942)
reg293944 := PrimTail(reg293943)
reg293945 := PrimEqual(reg293940, reg293944)
var reg293948 Obj
if reg293945 == True {
reg293946 := True;
reg293948 = reg293946
} else {
reg293947 := False;
reg293948 = reg293947
}
reg293950 = reg293948
} else {
reg293949 := False;
reg293950 = reg293949
}
var reg293953 Obj
if reg293950 == True {
reg293951 := True;
reg293953 = reg293951
} else {
reg293952 := False;
reg293953 = reg293952
}
reg293955 = reg293953
} else {
reg293954 := False;
reg293955 = reg293954
}
var reg293958 Obj
if reg293955 == True {
reg293956 := True;
reg293958 = reg293956
} else {
reg293957 := False;
reg293958 = reg293957
}
reg293960 = reg293958
} else {
reg293959 := False;
reg293960 = reg293959
}
var reg293963 Obj
if reg293960 == True {
reg293961 := True;
reg293963 = reg293961
} else {
reg293962 := False;
reg293963 = reg293962
}
reg293965 = reg293963
} else {
reg293964 := False;
reg293965 = reg293964
}
var reg293968 Obj
if reg293965 == True {
reg293966 := True;
reg293968 = reg293966
} else {
reg293967 := False;
reg293968 = reg293967
}
reg293970 = reg293968
} else {
reg293969 := False;
reg293970 = reg293969
}
if reg293970 == True {
reg293971 := MakeSymbol("let")
reg293972 := PrimTail(V1581)
reg293973 := PrimHead(reg293972)
reg293974 := PrimTail(V1581)
reg293975 := PrimTail(reg293974)
reg293976 := PrimHead(reg293975)
reg293977 := __e.Call(__defun__shen_4assign_1types, V1579, V1580, reg293976)
reg293978 := PrimTail(V1581)
reg293979 := PrimHead(reg293978)
reg293980 := PrimCons(reg293979, V1579)
reg293981 := PrimTail(V1581)
reg293982 := PrimTail(reg293981)
reg293983 := PrimTail(reg293982)
reg293984 := PrimHead(reg293983)
reg293985 := __e.Call(__defun__shen_4assign_1types, reg293980, V1580, reg293984)
reg293986 := Nil;
reg293987 := PrimCons(reg293985, reg293986)
reg293988 := PrimCons(reg293977, reg293987)
reg293989 := PrimCons(reg293973, reg293988)
reg293990 := PrimCons(reg293971, reg293989)
__ctx.Return(reg293990)
return
} else {
reg293991 := PrimIsPair(V1581)
var reg294024 Obj
if reg293991 == True {
reg293992 := MakeSymbol("lambda")
reg293993 := PrimHead(V1581)
reg293994 := PrimEqual(reg293992, reg293993)
var reg294019 Obj
if reg293994 == True {
reg293995 := PrimTail(V1581)
reg293996 := PrimIsPair(reg293995)
var reg294014 Obj
if reg293996 == True {
reg293997 := PrimTail(V1581)
reg293998 := PrimTail(reg293997)
reg293999 := PrimIsPair(reg293998)
var reg294009 Obj
if reg293999 == True {
reg294000 := Nil;
reg294001 := PrimTail(V1581)
reg294002 := PrimTail(reg294001)
reg294003 := PrimTail(reg294002)
reg294004 := PrimEqual(reg294000, reg294003)
var reg294007 Obj
if reg294004 == True {
reg294005 := True;
reg294007 = reg294005
} else {
reg294006 := False;
reg294007 = reg294006
}
reg294009 = reg294007
} else {
reg294008 := False;
reg294009 = reg294008
}
var reg294012 Obj
if reg294009 == True {
reg294010 := True;
reg294012 = reg294010
} else {
reg294011 := False;
reg294012 = reg294011
}
reg294014 = reg294012
} else {
reg294013 := False;
reg294014 = reg294013
}
var reg294017 Obj
if reg294014 == True {
reg294015 := True;
reg294017 = reg294015
} else {
reg294016 := False;
reg294017 = reg294016
}
reg294019 = reg294017
} else {
reg294018 := False;
reg294019 = reg294018
}
var reg294022 Obj
if reg294019 == True {
reg294020 := True;
reg294022 = reg294020
} else {
reg294021 := False;
reg294022 = reg294021
}
reg294024 = reg294022
} else {
reg294023 := False;
reg294024 = reg294023
}
if reg294024 == True {
reg294025 := MakeSymbol("lambda")
reg294026 := PrimTail(V1581)
reg294027 := PrimHead(reg294026)
reg294028 := PrimTail(V1581)
reg294029 := PrimHead(reg294028)
reg294030 := PrimCons(reg294029, V1579)
reg294031 := PrimTail(V1581)
reg294032 := PrimTail(reg294031)
reg294033 := PrimHead(reg294032)
reg294034 := __e.Call(__defun__shen_4assign_1types, reg294030, V1580, reg294033)
reg294035 := Nil;
reg294036 := PrimCons(reg294034, reg294035)
reg294037 := PrimCons(reg294027, reg294036)
reg294038 := PrimCons(reg294025, reg294037)
__ctx.Return(reg294038)
return
} else {
reg294039 := PrimIsPair(V1581)
var reg294047 Obj
if reg294039 == True {
reg294040 := MakeSymbol("cond")
reg294041 := PrimHead(V1581)
reg294042 := PrimEqual(reg294040, reg294041)
var reg294045 Obj
if reg294042 == True {
reg294043 := True;
reg294045 = reg294043
} else {
reg294044 := False;
reg294045 = reg294044
}
reg294047 = reg294045
} else {
reg294046 := False;
reg294047 = reg294046
}
if reg294047 == True {
reg294048 := MakeSymbol("cond")
reg294049 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
reg294050 := PrimHead(Y)
reg294051 := __e.Call(__defun__shen_4assign_1types, V1579, V1580, reg294050)
reg294052 := PrimTail(Y)
reg294053 := PrimHead(reg294052)
reg294054 := __e.Call(__defun__shen_4assign_1types, V1579, V1580, reg294053)
reg294055 := Nil;
reg294056 := PrimCons(reg294054, reg294055)
reg294057 := PrimCons(reg294051, reg294056)
__ctx.Return(reg294057)
return
}, 1)
reg294058 := PrimTail(V1581)
reg294059 := __e.Call(__defun__map, reg294049, reg294058)
reg294060 := PrimCons(reg294048, reg294059)
__ctx.Return(reg294060)
return
} else {
reg294061 := PrimIsPair(V1581)
if reg294061 == True {
reg294062 := PrimHead(V1581)
reg294063 := __e.Call(__defun__shen_4get_1type, reg294062)
reg294064 := PrimTail(V1581)
reg294065 := __e.Call(__defun__shen_4typextable, reg294063, reg294064)
NewTable := reg294065
_ = NewTable
reg294066 := PrimHead(V1581)
reg294067 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
reg294068 := __e.Call(__defun__append, V1580, NewTable)
__ctx.TailApply(__defun__shen_4assign_1types, V1579, reg294068, Y)
return
}, 1)
reg294070 := PrimTail(V1581)
reg294071 := __e.Call(__defun__map, reg294067, reg294070)
reg294072 := PrimCons(reg294066, reg294071)
__ctx.Return(reg294072)
return
} else {
reg294073 := __e.Call(__defun__assoc, V1581, V1580)
AtomType := reg294073
_ = AtomType
reg294074 := PrimIsPair(AtomType)
if reg294074 == True {
reg294075 := MakeSymbol("type")
reg294076 := PrimTail(AtomType)
reg294077 := Nil;
reg294078 := PrimCons(reg294076, reg294077)
reg294079 := PrimCons(V1581, reg294078)
reg294080 := PrimCons(reg294075, reg294079)
__ctx.Return(reg294080)
return
} else {
reg294081 := __e.Call(__defun__element_2, V1581, V1579)
if reg294081 == True {
__ctx.Return(V1581)
return
} else {
__ctx.TailApply(__defun__shen_4atom_1type, V1581)
return
}
}
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.assign-types", value: __defun__shen_4assign_1types})

__defun__shen_4atom_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1583 := __args[0]
_ = V1583
reg294083 := PrimIsString(V1583)
if reg294083 == True {
reg294084 := MakeSymbol("type")
reg294085 := MakeSymbol("string")
reg294086 := Nil;
reg294087 := PrimCons(reg294085, reg294086)
reg294088 := PrimCons(V1583, reg294087)
reg294089 := PrimCons(reg294084, reg294088)
__ctx.Return(reg294089)
return
} else {
reg294090 := PrimIsNumber(V1583)
if reg294090 == True {
reg294091 := MakeSymbol("type")
reg294092 := MakeSymbol("number")
reg294093 := Nil;
reg294094 := PrimCons(reg294092, reg294093)
reg294095 := PrimCons(V1583, reg294094)
reg294096 := PrimCons(reg294091, reg294095)
__ctx.Return(reg294096)
return
} else {
reg294097 := __e.Call(__defun__boolean_2, V1583)
if reg294097 == True {
reg294098 := MakeSymbol("type")
reg294099 := MakeSymbol("boolean")
reg294100 := Nil;
reg294101 := PrimCons(reg294099, reg294100)
reg294102 := PrimCons(V1583, reg294101)
reg294103 := PrimCons(reg294098, reg294102)
__ctx.Return(reg294103)
return
} else {
reg294104 := PrimIsSymbol(V1583)
if reg294104 == True {
reg294105 := MakeSymbol("type")
reg294106 := MakeSymbol("symbol")
reg294107 := Nil;
reg294108 := PrimCons(reg294106, reg294107)
reg294109 := PrimCons(V1583, reg294108)
reg294110 := PrimCons(reg294105, reg294109)
__ctx.Return(reg294110)
return
} else {
__ctx.Return(V1583)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.atom-type", value: __defun__shen_4atom_1type})

__defun__shen_4store_1arity = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1588 := __args[0]
_ = V1588
V1589 := __args[1]
_ = V1589
reg294111 := MakeSymbol("shen.*installing-kl*")
reg294112 := PrimValue(reg294111)
if reg294112 == True {
reg294113 := MakeSymbol("shen.skip")
__ctx.Return(reg294113)
return
} else {
reg294114 := MakeSymbol("arity")
reg294115 := MakeSymbol("*property-vector*")
reg294116 := PrimValue(reg294115)
__ctx.TailApply(__defun__put, V1588, reg294114, V1589, reg294116)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.store-arity", value: __defun__shen_4store_1arity})

__defun__shen_4reduce = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1591 := __args[0]
_ = V1591
reg294118 := MakeSymbol("shen.*teststack*")
reg294119 := Nil;
reg294120 := PrimSet(reg294118, reg294119)
_ = reg294120
reg294121 := __e.Call(__defun__shen_4reduce__help, V1591)
Result := reg294121
_ = Result
reg294122 := MakeSymbol(":")
reg294123 := MakeSymbol("shen.tests")
reg294124 := MakeSymbol("shen.*teststack*")
reg294125 := PrimValue(reg294124)
reg294126 := __e.Call(__defun__reverse, reg294125)
reg294127 := PrimCons(reg294123, reg294126)
reg294128 := PrimCons(reg294122, reg294127)
reg294129 := Nil;
reg294130 := PrimCons(Result, reg294129)
reg294131 := PrimCons(reg294128, reg294130)
__ctx.Return(reg294131)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.reduce", value: __defun__shen_4reduce})

__defun__shen_4reduce__help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1593 := __args[0]
_ = V1593
reg294132 := PrimIsPair(V1593)
var reg294246 Obj
if reg294132 == True {
reg294133 := PrimHead(V1593)
reg294134 := PrimIsPair(reg294133)
var reg294241 Obj
if reg294134 == True {
reg294135 := MakeSymbol("/.")
reg294136 := PrimHead(V1593)
reg294137 := PrimHead(reg294136)
reg294138 := PrimEqual(reg294135, reg294137)
var reg294236 Obj
if reg294138 == True {
reg294139 := PrimHead(V1593)
reg294140 := PrimTail(reg294139)
reg294141 := PrimIsPair(reg294140)
var reg294231 Obj
if reg294141 == True {
reg294142 := PrimHead(V1593)
reg294143 := PrimTail(reg294142)
reg294144 := PrimHead(reg294143)
reg294145 := PrimIsPair(reg294144)
var reg294226 Obj
if reg294145 == True {
reg294146 := MakeSymbol("cons")
reg294147 := PrimHead(V1593)
reg294148 := PrimTail(reg294147)
reg294149 := PrimHead(reg294148)
reg294150 := PrimHead(reg294149)
reg294151 := PrimEqual(reg294146, reg294150)
var reg294221 Obj
if reg294151 == True {
reg294152 := PrimHead(V1593)
reg294153 := PrimTail(reg294152)
reg294154 := PrimHead(reg294153)
reg294155 := PrimTail(reg294154)
reg294156 := PrimIsPair(reg294155)
var reg294216 Obj
if reg294156 == True {
reg294157 := PrimHead(V1593)
reg294158 := PrimTail(reg294157)
reg294159 := PrimHead(reg294158)
reg294160 := PrimTail(reg294159)
reg294161 := PrimTail(reg294160)
reg294162 := PrimIsPair(reg294161)
var reg294211 Obj
if reg294162 == True {
reg294163 := Nil;
reg294164 := PrimHead(V1593)
reg294165 := PrimTail(reg294164)
reg294166 := PrimHead(reg294165)
reg294167 := PrimTail(reg294166)
reg294168 := PrimTail(reg294167)
reg294169 := PrimTail(reg294168)
reg294170 := PrimEqual(reg294163, reg294169)
var reg294206 Obj
if reg294170 == True {
reg294171 := PrimHead(V1593)
reg294172 := PrimTail(reg294171)
reg294173 := PrimTail(reg294172)
reg294174 := PrimIsPair(reg294173)
var reg294201 Obj
if reg294174 == True {
reg294175 := Nil;
reg294176 := PrimHead(V1593)
reg294177 := PrimTail(reg294176)
reg294178 := PrimTail(reg294177)
reg294179 := PrimTail(reg294178)
reg294180 := PrimEqual(reg294175, reg294179)
var reg294196 Obj
if reg294180 == True {
reg294181 := PrimTail(V1593)
reg294182 := PrimIsPair(reg294181)
var reg294191 Obj
if reg294182 == True {
reg294183 := Nil;
reg294184 := PrimTail(V1593)
reg294185 := PrimTail(reg294184)
reg294186 := PrimEqual(reg294183, reg294185)
var reg294189 Obj
if reg294186 == True {
reg294187 := True;
reg294189 = reg294187
} else {
reg294188 := False;
reg294189 = reg294188
}
reg294191 = reg294189
} else {
reg294190 := False;
reg294191 = reg294190
}
var reg294194 Obj
if reg294191 == True {
reg294192 := True;
reg294194 = reg294192
} else {
reg294193 := False;
reg294194 = reg294193
}
reg294196 = reg294194
} else {
reg294195 := False;
reg294196 = reg294195
}
var reg294199 Obj
if reg294196 == True {
reg294197 := True;
reg294199 = reg294197
} else {
reg294198 := False;
reg294199 = reg294198
}
reg294201 = reg294199
} else {
reg294200 := False;
reg294201 = reg294200
}
var reg294204 Obj
if reg294201 == True {
reg294202 := True;
reg294204 = reg294202
} else {
reg294203 := False;
reg294204 = reg294203
}
reg294206 = reg294204
} else {
reg294205 := False;
reg294206 = reg294205
}
var reg294209 Obj
if reg294206 == True {
reg294207 := True;
reg294209 = reg294207
} else {
reg294208 := False;
reg294209 = reg294208
}
reg294211 = reg294209
} else {
reg294210 := False;
reg294211 = reg294210
}
var reg294214 Obj
if reg294211 == True {
reg294212 := True;
reg294214 = reg294212
} else {
reg294213 := False;
reg294214 = reg294213
}
reg294216 = reg294214
} else {
reg294215 := False;
reg294216 = reg294215
}
var reg294219 Obj
if reg294216 == True {
reg294217 := True;
reg294219 = reg294217
} else {
reg294218 := False;
reg294219 = reg294218
}
reg294221 = reg294219
} else {
reg294220 := False;
reg294221 = reg294220
}
var reg294224 Obj
if reg294221 == True {
reg294222 := True;
reg294224 = reg294222
} else {
reg294223 := False;
reg294224 = reg294223
}
reg294226 = reg294224
} else {
reg294225 := False;
reg294226 = reg294225
}
var reg294229 Obj
if reg294226 == True {
reg294227 := True;
reg294229 = reg294227
} else {
reg294228 := False;
reg294229 = reg294228
}
reg294231 = reg294229
} else {
reg294230 := False;
reg294231 = reg294230
}
var reg294234 Obj
if reg294231 == True {
reg294232 := True;
reg294234 = reg294232
} else {
reg294233 := False;
reg294234 = reg294233
}
reg294236 = reg294234
} else {
reg294235 := False;
reg294236 = reg294235
}
var reg294239 Obj
if reg294236 == True {
reg294237 := True;
reg294239 = reg294237
} else {
reg294238 := False;
reg294239 = reg294238
}
reg294241 = reg294239
} else {
reg294240 := False;
reg294241 = reg294240
}
var reg294244 Obj
if reg294241 == True {
reg294242 := True;
reg294244 = reg294242
} else {
reg294243 := False;
reg294244 = reg294243
}
reg294246 = reg294244
} else {
reg294245 := False;
reg294246 = reg294245
}
if reg294246 == True {
reg294247 := MakeSymbol("cons?")
reg294248 := PrimTail(V1593)
reg294249 := PrimCons(reg294247, reg294248)
reg294250 := __e.Call(__defun__shen_4add__test, reg294249)
_ = reg294250
reg294251 := MakeSymbol("/.")
reg294252 := PrimHead(V1593)
reg294253 := PrimTail(reg294252)
reg294254 := PrimHead(reg294253)
reg294255 := PrimTail(reg294254)
reg294256 := PrimHead(reg294255)
reg294257 := MakeSymbol("/.")
reg294258 := PrimHead(V1593)
reg294259 := PrimTail(reg294258)
reg294260 := PrimHead(reg294259)
reg294261 := PrimTail(reg294260)
reg294262 := PrimTail(reg294261)
reg294263 := PrimHead(reg294262)
reg294264 := PrimTail(V1593)
reg294265 := PrimHead(reg294264)
reg294266 := PrimHead(V1593)
reg294267 := PrimTail(reg294266)
reg294268 := PrimHead(reg294267)
reg294269 := PrimHead(V1593)
reg294270 := PrimTail(reg294269)
reg294271 := PrimTail(reg294270)
reg294272 := PrimHead(reg294271)
reg294273 := __e.Call(__defun__shen_4ebr, reg294265, reg294268, reg294272)
reg294274 := Nil;
reg294275 := PrimCons(reg294273, reg294274)
reg294276 := PrimCons(reg294263, reg294275)
reg294277 := PrimCons(reg294257, reg294276)
reg294278 := Nil;
reg294279 := PrimCons(reg294277, reg294278)
reg294280 := PrimCons(reg294256, reg294279)
reg294281 := PrimCons(reg294251, reg294280)
Abstraction := reg294281
_ = Abstraction
reg294282 := MakeSymbol("hd")
reg294283 := PrimTail(V1593)
reg294284 := PrimCons(reg294282, reg294283)
reg294285 := Nil;
reg294286 := PrimCons(reg294284, reg294285)
reg294287 := PrimCons(Abstraction, reg294286)
reg294288 := MakeSymbol("tl")
reg294289 := PrimTail(V1593)
reg294290 := PrimCons(reg294288, reg294289)
reg294291 := Nil;
reg294292 := PrimCons(reg294290, reg294291)
reg294293 := PrimCons(reg294287, reg294292)
Application := reg294293
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg294295 := PrimIsPair(V1593)
var reg294409 Obj
if reg294295 == True {
reg294296 := PrimHead(V1593)
reg294297 := PrimIsPair(reg294296)
var reg294404 Obj
if reg294297 == True {
reg294298 := MakeSymbol("/.")
reg294299 := PrimHead(V1593)
reg294300 := PrimHead(reg294299)
reg294301 := PrimEqual(reg294298, reg294300)
var reg294399 Obj
if reg294301 == True {
reg294302 := PrimHead(V1593)
reg294303 := PrimTail(reg294302)
reg294304 := PrimIsPair(reg294303)
var reg294394 Obj
if reg294304 == True {
reg294305 := PrimHead(V1593)
reg294306 := PrimTail(reg294305)
reg294307 := PrimHead(reg294306)
reg294308 := PrimIsPair(reg294307)
var reg294389 Obj
if reg294308 == True {
reg294309 := MakeSymbol("@p")
reg294310 := PrimHead(V1593)
reg294311 := PrimTail(reg294310)
reg294312 := PrimHead(reg294311)
reg294313 := PrimHead(reg294312)
reg294314 := PrimEqual(reg294309, reg294313)
var reg294384 Obj
if reg294314 == True {
reg294315 := PrimHead(V1593)
reg294316 := PrimTail(reg294315)
reg294317 := PrimHead(reg294316)
reg294318 := PrimTail(reg294317)
reg294319 := PrimIsPair(reg294318)
var reg294379 Obj
if reg294319 == True {
reg294320 := PrimHead(V1593)
reg294321 := PrimTail(reg294320)
reg294322 := PrimHead(reg294321)
reg294323 := PrimTail(reg294322)
reg294324 := PrimTail(reg294323)
reg294325 := PrimIsPair(reg294324)
var reg294374 Obj
if reg294325 == True {
reg294326 := Nil;
reg294327 := PrimHead(V1593)
reg294328 := PrimTail(reg294327)
reg294329 := PrimHead(reg294328)
reg294330 := PrimTail(reg294329)
reg294331 := PrimTail(reg294330)
reg294332 := PrimTail(reg294331)
reg294333 := PrimEqual(reg294326, reg294332)
var reg294369 Obj
if reg294333 == True {
reg294334 := PrimHead(V1593)
reg294335 := PrimTail(reg294334)
reg294336 := PrimTail(reg294335)
reg294337 := PrimIsPair(reg294336)
var reg294364 Obj
if reg294337 == True {
reg294338 := Nil;
reg294339 := PrimHead(V1593)
reg294340 := PrimTail(reg294339)
reg294341 := PrimTail(reg294340)
reg294342 := PrimTail(reg294341)
reg294343 := PrimEqual(reg294338, reg294342)
var reg294359 Obj
if reg294343 == True {
reg294344 := PrimTail(V1593)
reg294345 := PrimIsPair(reg294344)
var reg294354 Obj
if reg294345 == True {
reg294346 := Nil;
reg294347 := PrimTail(V1593)
reg294348 := PrimTail(reg294347)
reg294349 := PrimEqual(reg294346, reg294348)
var reg294352 Obj
if reg294349 == True {
reg294350 := True;
reg294352 = reg294350
} else {
reg294351 := False;
reg294352 = reg294351
}
reg294354 = reg294352
} else {
reg294353 := False;
reg294354 = reg294353
}
var reg294357 Obj
if reg294354 == True {
reg294355 := True;
reg294357 = reg294355
} else {
reg294356 := False;
reg294357 = reg294356
}
reg294359 = reg294357
} else {
reg294358 := False;
reg294359 = reg294358
}
var reg294362 Obj
if reg294359 == True {
reg294360 := True;
reg294362 = reg294360
} else {
reg294361 := False;
reg294362 = reg294361
}
reg294364 = reg294362
} else {
reg294363 := False;
reg294364 = reg294363
}
var reg294367 Obj
if reg294364 == True {
reg294365 := True;
reg294367 = reg294365
} else {
reg294366 := False;
reg294367 = reg294366
}
reg294369 = reg294367
} else {
reg294368 := False;
reg294369 = reg294368
}
var reg294372 Obj
if reg294369 == True {
reg294370 := True;
reg294372 = reg294370
} else {
reg294371 := False;
reg294372 = reg294371
}
reg294374 = reg294372
} else {
reg294373 := False;
reg294374 = reg294373
}
var reg294377 Obj
if reg294374 == True {
reg294375 := True;
reg294377 = reg294375
} else {
reg294376 := False;
reg294377 = reg294376
}
reg294379 = reg294377
} else {
reg294378 := False;
reg294379 = reg294378
}
var reg294382 Obj
if reg294379 == True {
reg294380 := True;
reg294382 = reg294380
} else {
reg294381 := False;
reg294382 = reg294381
}
reg294384 = reg294382
} else {
reg294383 := False;
reg294384 = reg294383
}
var reg294387 Obj
if reg294384 == True {
reg294385 := True;
reg294387 = reg294385
} else {
reg294386 := False;
reg294387 = reg294386
}
reg294389 = reg294387
} else {
reg294388 := False;
reg294389 = reg294388
}
var reg294392 Obj
if reg294389 == True {
reg294390 := True;
reg294392 = reg294390
} else {
reg294391 := False;
reg294392 = reg294391
}
reg294394 = reg294392
} else {
reg294393 := False;
reg294394 = reg294393
}
var reg294397 Obj
if reg294394 == True {
reg294395 := True;
reg294397 = reg294395
} else {
reg294396 := False;
reg294397 = reg294396
}
reg294399 = reg294397
} else {
reg294398 := False;
reg294399 = reg294398
}
var reg294402 Obj
if reg294399 == True {
reg294400 := True;
reg294402 = reg294400
} else {
reg294401 := False;
reg294402 = reg294401
}
reg294404 = reg294402
} else {
reg294403 := False;
reg294404 = reg294403
}
var reg294407 Obj
if reg294404 == True {
reg294405 := True;
reg294407 = reg294405
} else {
reg294406 := False;
reg294407 = reg294406
}
reg294409 = reg294407
} else {
reg294408 := False;
reg294409 = reg294408
}
if reg294409 == True {
reg294410 := MakeSymbol("tuple?")
reg294411 := PrimTail(V1593)
reg294412 := PrimCons(reg294410, reg294411)
reg294413 := __e.Call(__defun__shen_4add__test, reg294412)
_ = reg294413
reg294414 := MakeSymbol("/.")
reg294415 := PrimHead(V1593)
reg294416 := PrimTail(reg294415)
reg294417 := PrimHead(reg294416)
reg294418 := PrimTail(reg294417)
reg294419 := PrimHead(reg294418)
reg294420 := MakeSymbol("/.")
reg294421 := PrimHead(V1593)
reg294422 := PrimTail(reg294421)
reg294423 := PrimHead(reg294422)
reg294424 := PrimTail(reg294423)
reg294425 := PrimTail(reg294424)
reg294426 := PrimHead(reg294425)
reg294427 := PrimTail(V1593)
reg294428 := PrimHead(reg294427)
reg294429 := PrimHead(V1593)
reg294430 := PrimTail(reg294429)
reg294431 := PrimHead(reg294430)
reg294432 := PrimHead(V1593)
reg294433 := PrimTail(reg294432)
reg294434 := PrimTail(reg294433)
reg294435 := PrimHead(reg294434)
reg294436 := __e.Call(__defun__shen_4ebr, reg294428, reg294431, reg294435)
reg294437 := Nil;
reg294438 := PrimCons(reg294436, reg294437)
reg294439 := PrimCons(reg294426, reg294438)
reg294440 := PrimCons(reg294420, reg294439)
reg294441 := Nil;
reg294442 := PrimCons(reg294440, reg294441)
reg294443 := PrimCons(reg294419, reg294442)
reg294444 := PrimCons(reg294414, reg294443)
Abstraction := reg294444
_ = Abstraction
reg294445 := MakeSymbol("fst")
reg294446 := PrimTail(V1593)
reg294447 := PrimCons(reg294445, reg294446)
reg294448 := Nil;
reg294449 := PrimCons(reg294447, reg294448)
reg294450 := PrimCons(Abstraction, reg294449)
reg294451 := MakeSymbol("snd")
reg294452 := PrimTail(V1593)
reg294453 := PrimCons(reg294451, reg294452)
reg294454 := Nil;
reg294455 := PrimCons(reg294453, reg294454)
reg294456 := PrimCons(reg294450, reg294455)
Application := reg294456
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg294458 := PrimIsPair(V1593)
var reg294572 Obj
if reg294458 == True {
reg294459 := PrimHead(V1593)
reg294460 := PrimIsPair(reg294459)
var reg294567 Obj
if reg294460 == True {
reg294461 := MakeSymbol("/.")
reg294462 := PrimHead(V1593)
reg294463 := PrimHead(reg294462)
reg294464 := PrimEqual(reg294461, reg294463)
var reg294562 Obj
if reg294464 == True {
reg294465 := PrimHead(V1593)
reg294466 := PrimTail(reg294465)
reg294467 := PrimIsPair(reg294466)
var reg294557 Obj
if reg294467 == True {
reg294468 := PrimHead(V1593)
reg294469 := PrimTail(reg294468)
reg294470 := PrimHead(reg294469)
reg294471 := PrimIsPair(reg294470)
var reg294552 Obj
if reg294471 == True {
reg294472 := MakeSymbol("@v")
reg294473 := PrimHead(V1593)
reg294474 := PrimTail(reg294473)
reg294475 := PrimHead(reg294474)
reg294476 := PrimHead(reg294475)
reg294477 := PrimEqual(reg294472, reg294476)
var reg294547 Obj
if reg294477 == True {
reg294478 := PrimHead(V1593)
reg294479 := PrimTail(reg294478)
reg294480 := PrimHead(reg294479)
reg294481 := PrimTail(reg294480)
reg294482 := PrimIsPair(reg294481)
var reg294542 Obj
if reg294482 == True {
reg294483 := PrimHead(V1593)
reg294484 := PrimTail(reg294483)
reg294485 := PrimHead(reg294484)
reg294486 := PrimTail(reg294485)
reg294487 := PrimTail(reg294486)
reg294488 := PrimIsPair(reg294487)
var reg294537 Obj
if reg294488 == True {
reg294489 := Nil;
reg294490 := PrimHead(V1593)
reg294491 := PrimTail(reg294490)
reg294492 := PrimHead(reg294491)
reg294493 := PrimTail(reg294492)
reg294494 := PrimTail(reg294493)
reg294495 := PrimTail(reg294494)
reg294496 := PrimEqual(reg294489, reg294495)
var reg294532 Obj
if reg294496 == True {
reg294497 := PrimHead(V1593)
reg294498 := PrimTail(reg294497)
reg294499 := PrimTail(reg294498)
reg294500 := PrimIsPair(reg294499)
var reg294527 Obj
if reg294500 == True {
reg294501 := Nil;
reg294502 := PrimHead(V1593)
reg294503 := PrimTail(reg294502)
reg294504 := PrimTail(reg294503)
reg294505 := PrimTail(reg294504)
reg294506 := PrimEqual(reg294501, reg294505)
var reg294522 Obj
if reg294506 == True {
reg294507 := PrimTail(V1593)
reg294508 := PrimIsPair(reg294507)
var reg294517 Obj
if reg294508 == True {
reg294509 := Nil;
reg294510 := PrimTail(V1593)
reg294511 := PrimTail(reg294510)
reg294512 := PrimEqual(reg294509, reg294511)
var reg294515 Obj
if reg294512 == True {
reg294513 := True;
reg294515 = reg294513
} else {
reg294514 := False;
reg294515 = reg294514
}
reg294517 = reg294515
} else {
reg294516 := False;
reg294517 = reg294516
}
var reg294520 Obj
if reg294517 == True {
reg294518 := True;
reg294520 = reg294518
} else {
reg294519 := False;
reg294520 = reg294519
}
reg294522 = reg294520
} else {
reg294521 := False;
reg294522 = reg294521
}
var reg294525 Obj
if reg294522 == True {
reg294523 := True;
reg294525 = reg294523
} else {
reg294524 := False;
reg294525 = reg294524
}
reg294527 = reg294525
} else {
reg294526 := False;
reg294527 = reg294526
}
var reg294530 Obj
if reg294527 == True {
reg294528 := True;
reg294530 = reg294528
} else {
reg294529 := False;
reg294530 = reg294529
}
reg294532 = reg294530
} else {
reg294531 := False;
reg294532 = reg294531
}
var reg294535 Obj
if reg294532 == True {
reg294533 := True;
reg294535 = reg294533
} else {
reg294534 := False;
reg294535 = reg294534
}
reg294537 = reg294535
} else {
reg294536 := False;
reg294537 = reg294536
}
var reg294540 Obj
if reg294537 == True {
reg294538 := True;
reg294540 = reg294538
} else {
reg294539 := False;
reg294540 = reg294539
}
reg294542 = reg294540
} else {
reg294541 := False;
reg294542 = reg294541
}
var reg294545 Obj
if reg294542 == True {
reg294543 := True;
reg294545 = reg294543
} else {
reg294544 := False;
reg294545 = reg294544
}
reg294547 = reg294545
} else {
reg294546 := False;
reg294547 = reg294546
}
var reg294550 Obj
if reg294547 == True {
reg294548 := True;
reg294550 = reg294548
} else {
reg294549 := False;
reg294550 = reg294549
}
reg294552 = reg294550
} else {
reg294551 := False;
reg294552 = reg294551
}
var reg294555 Obj
if reg294552 == True {
reg294553 := True;
reg294555 = reg294553
} else {
reg294554 := False;
reg294555 = reg294554
}
reg294557 = reg294555
} else {
reg294556 := False;
reg294557 = reg294556
}
var reg294560 Obj
if reg294557 == True {
reg294558 := True;
reg294560 = reg294558
} else {
reg294559 := False;
reg294560 = reg294559
}
reg294562 = reg294560
} else {
reg294561 := False;
reg294562 = reg294561
}
var reg294565 Obj
if reg294562 == True {
reg294563 := True;
reg294565 = reg294563
} else {
reg294564 := False;
reg294565 = reg294564
}
reg294567 = reg294565
} else {
reg294566 := False;
reg294567 = reg294566
}
var reg294570 Obj
if reg294567 == True {
reg294568 := True;
reg294570 = reg294568
} else {
reg294569 := False;
reg294570 = reg294569
}
reg294572 = reg294570
} else {
reg294571 := False;
reg294572 = reg294571
}
if reg294572 == True {
reg294573 := MakeSymbol("shen.+vector?")
reg294574 := PrimTail(V1593)
reg294575 := PrimCons(reg294573, reg294574)
reg294576 := __e.Call(__defun__shen_4add__test, reg294575)
_ = reg294576
reg294577 := MakeSymbol("/.")
reg294578 := PrimHead(V1593)
reg294579 := PrimTail(reg294578)
reg294580 := PrimHead(reg294579)
reg294581 := PrimTail(reg294580)
reg294582 := PrimHead(reg294581)
reg294583 := MakeSymbol("/.")
reg294584 := PrimHead(V1593)
reg294585 := PrimTail(reg294584)
reg294586 := PrimHead(reg294585)
reg294587 := PrimTail(reg294586)
reg294588 := PrimTail(reg294587)
reg294589 := PrimHead(reg294588)
reg294590 := PrimTail(V1593)
reg294591 := PrimHead(reg294590)
reg294592 := PrimHead(V1593)
reg294593 := PrimTail(reg294592)
reg294594 := PrimHead(reg294593)
reg294595 := PrimHead(V1593)
reg294596 := PrimTail(reg294595)
reg294597 := PrimTail(reg294596)
reg294598 := PrimHead(reg294597)
reg294599 := __e.Call(__defun__shen_4ebr, reg294591, reg294594, reg294598)
reg294600 := Nil;
reg294601 := PrimCons(reg294599, reg294600)
reg294602 := PrimCons(reg294589, reg294601)
reg294603 := PrimCons(reg294583, reg294602)
reg294604 := Nil;
reg294605 := PrimCons(reg294603, reg294604)
reg294606 := PrimCons(reg294582, reg294605)
reg294607 := PrimCons(reg294577, reg294606)
Abstraction := reg294607
_ = Abstraction
reg294608 := MakeSymbol("hdv")
reg294609 := PrimTail(V1593)
reg294610 := PrimCons(reg294608, reg294609)
reg294611 := Nil;
reg294612 := PrimCons(reg294610, reg294611)
reg294613 := PrimCons(Abstraction, reg294612)
reg294614 := MakeSymbol("tlv")
reg294615 := PrimTail(V1593)
reg294616 := PrimCons(reg294614, reg294615)
reg294617 := Nil;
reg294618 := PrimCons(reg294616, reg294617)
reg294619 := PrimCons(reg294613, reg294618)
Application := reg294619
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg294621 := PrimIsPair(V1593)
var reg294735 Obj
if reg294621 == True {
reg294622 := PrimHead(V1593)
reg294623 := PrimIsPair(reg294622)
var reg294730 Obj
if reg294623 == True {
reg294624 := MakeSymbol("/.")
reg294625 := PrimHead(V1593)
reg294626 := PrimHead(reg294625)
reg294627 := PrimEqual(reg294624, reg294626)
var reg294725 Obj
if reg294627 == True {
reg294628 := PrimHead(V1593)
reg294629 := PrimTail(reg294628)
reg294630 := PrimIsPair(reg294629)
var reg294720 Obj
if reg294630 == True {
reg294631 := PrimHead(V1593)
reg294632 := PrimTail(reg294631)
reg294633 := PrimHead(reg294632)
reg294634 := PrimIsPair(reg294633)
var reg294715 Obj
if reg294634 == True {
reg294635 := MakeSymbol("@s")
reg294636 := PrimHead(V1593)
reg294637 := PrimTail(reg294636)
reg294638 := PrimHead(reg294637)
reg294639 := PrimHead(reg294638)
reg294640 := PrimEqual(reg294635, reg294639)
var reg294710 Obj
if reg294640 == True {
reg294641 := PrimHead(V1593)
reg294642 := PrimTail(reg294641)
reg294643 := PrimHead(reg294642)
reg294644 := PrimTail(reg294643)
reg294645 := PrimIsPair(reg294644)
var reg294705 Obj
if reg294645 == True {
reg294646 := PrimHead(V1593)
reg294647 := PrimTail(reg294646)
reg294648 := PrimHead(reg294647)
reg294649 := PrimTail(reg294648)
reg294650 := PrimTail(reg294649)
reg294651 := PrimIsPair(reg294650)
var reg294700 Obj
if reg294651 == True {
reg294652 := Nil;
reg294653 := PrimHead(V1593)
reg294654 := PrimTail(reg294653)
reg294655 := PrimHead(reg294654)
reg294656 := PrimTail(reg294655)
reg294657 := PrimTail(reg294656)
reg294658 := PrimTail(reg294657)
reg294659 := PrimEqual(reg294652, reg294658)
var reg294695 Obj
if reg294659 == True {
reg294660 := PrimHead(V1593)
reg294661 := PrimTail(reg294660)
reg294662 := PrimTail(reg294661)
reg294663 := PrimIsPair(reg294662)
var reg294690 Obj
if reg294663 == True {
reg294664 := Nil;
reg294665 := PrimHead(V1593)
reg294666 := PrimTail(reg294665)
reg294667 := PrimTail(reg294666)
reg294668 := PrimTail(reg294667)
reg294669 := PrimEqual(reg294664, reg294668)
var reg294685 Obj
if reg294669 == True {
reg294670 := PrimTail(V1593)
reg294671 := PrimIsPair(reg294670)
var reg294680 Obj
if reg294671 == True {
reg294672 := Nil;
reg294673 := PrimTail(V1593)
reg294674 := PrimTail(reg294673)
reg294675 := PrimEqual(reg294672, reg294674)
var reg294678 Obj
if reg294675 == True {
reg294676 := True;
reg294678 = reg294676
} else {
reg294677 := False;
reg294678 = reg294677
}
reg294680 = reg294678
} else {
reg294679 := False;
reg294680 = reg294679
}
var reg294683 Obj
if reg294680 == True {
reg294681 := True;
reg294683 = reg294681
} else {
reg294682 := False;
reg294683 = reg294682
}
reg294685 = reg294683
} else {
reg294684 := False;
reg294685 = reg294684
}
var reg294688 Obj
if reg294685 == True {
reg294686 := True;
reg294688 = reg294686
} else {
reg294687 := False;
reg294688 = reg294687
}
reg294690 = reg294688
} else {
reg294689 := False;
reg294690 = reg294689
}
var reg294693 Obj
if reg294690 == True {
reg294691 := True;
reg294693 = reg294691
} else {
reg294692 := False;
reg294693 = reg294692
}
reg294695 = reg294693
} else {
reg294694 := False;
reg294695 = reg294694
}
var reg294698 Obj
if reg294695 == True {
reg294696 := True;
reg294698 = reg294696
} else {
reg294697 := False;
reg294698 = reg294697
}
reg294700 = reg294698
} else {
reg294699 := False;
reg294700 = reg294699
}
var reg294703 Obj
if reg294700 == True {
reg294701 := True;
reg294703 = reg294701
} else {
reg294702 := False;
reg294703 = reg294702
}
reg294705 = reg294703
} else {
reg294704 := False;
reg294705 = reg294704
}
var reg294708 Obj
if reg294705 == True {
reg294706 := True;
reg294708 = reg294706
} else {
reg294707 := False;
reg294708 = reg294707
}
reg294710 = reg294708
} else {
reg294709 := False;
reg294710 = reg294709
}
var reg294713 Obj
if reg294710 == True {
reg294711 := True;
reg294713 = reg294711
} else {
reg294712 := False;
reg294713 = reg294712
}
reg294715 = reg294713
} else {
reg294714 := False;
reg294715 = reg294714
}
var reg294718 Obj
if reg294715 == True {
reg294716 := True;
reg294718 = reg294716
} else {
reg294717 := False;
reg294718 = reg294717
}
reg294720 = reg294718
} else {
reg294719 := False;
reg294720 = reg294719
}
var reg294723 Obj
if reg294720 == True {
reg294721 := True;
reg294723 = reg294721
} else {
reg294722 := False;
reg294723 = reg294722
}
reg294725 = reg294723
} else {
reg294724 := False;
reg294725 = reg294724
}
var reg294728 Obj
if reg294725 == True {
reg294726 := True;
reg294728 = reg294726
} else {
reg294727 := False;
reg294728 = reg294727
}
reg294730 = reg294728
} else {
reg294729 := False;
reg294730 = reg294729
}
var reg294733 Obj
if reg294730 == True {
reg294731 := True;
reg294733 = reg294731
} else {
reg294732 := False;
reg294733 = reg294732
}
reg294735 = reg294733
} else {
reg294734 := False;
reg294735 = reg294734
}
if reg294735 == True {
reg294736 := MakeSymbol("shen.+string?")
reg294737 := PrimTail(V1593)
reg294738 := PrimCons(reg294736, reg294737)
reg294739 := __e.Call(__defun__shen_4add__test, reg294738)
_ = reg294739
reg294740 := MakeSymbol("/.")
reg294741 := PrimHead(V1593)
reg294742 := PrimTail(reg294741)
reg294743 := PrimHead(reg294742)
reg294744 := PrimTail(reg294743)
reg294745 := PrimHead(reg294744)
reg294746 := MakeSymbol("/.")
reg294747 := PrimHead(V1593)
reg294748 := PrimTail(reg294747)
reg294749 := PrimHead(reg294748)
reg294750 := PrimTail(reg294749)
reg294751 := PrimTail(reg294750)
reg294752 := PrimHead(reg294751)
reg294753 := PrimTail(V1593)
reg294754 := PrimHead(reg294753)
reg294755 := PrimHead(V1593)
reg294756 := PrimTail(reg294755)
reg294757 := PrimHead(reg294756)
reg294758 := PrimHead(V1593)
reg294759 := PrimTail(reg294758)
reg294760 := PrimTail(reg294759)
reg294761 := PrimHead(reg294760)
reg294762 := __e.Call(__defun__shen_4ebr, reg294754, reg294757, reg294761)
reg294763 := Nil;
reg294764 := PrimCons(reg294762, reg294763)
reg294765 := PrimCons(reg294752, reg294764)
reg294766 := PrimCons(reg294746, reg294765)
reg294767 := Nil;
reg294768 := PrimCons(reg294766, reg294767)
reg294769 := PrimCons(reg294745, reg294768)
reg294770 := PrimCons(reg294740, reg294769)
Abstraction := reg294770
_ = Abstraction
reg294771 := MakeSymbol("pos")
reg294772 := PrimTail(V1593)
reg294773 := PrimHead(reg294772)
reg294774 := MakeNumber(0)
reg294775 := Nil;
reg294776 := PrimCons(reg294774, reg294775)
reg294777 := PrimCons(reg294773, reg294776)
reg294778 := PrimCons(reg294771, reg294777)
reg294779 := Nil;
reg294780 := PrimCons(reg294778, reg294779)
reg294781 := PrimCons(Abstraction, reg294780)
reg294782 := MakeSymbol("tlstr")
reg294783 := PrimTail(V1593)
reg294784 := PrimCons(reg294782, reg294783)
reg294785 := Nil;
reg294786 := PrimCons(reg294784, reg294785)
reg294787 := PrimCons(reg294781, reg294786)
Application := reg294787
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg294789 := PrimIsPair(V1593)
var reg294859 Obj
if reg294789 == True {
reg294790 := PrimHead(V1593)
reg294791 := PrimIsPair(reg294790)
var reg294854 Obj
if reg294791 == True {
reg294792 := MakeSymbol("/.")
reg294793 := PrimHead(V1593)
reg294794 := PrimHead(reg294793)
reg294795 := PrimEqual(reg294792, reg294794)
var reg294849 Obj
if reg294795 == True {
reg294796 := PrimHead(V1593)
reg294797 := PrimTail(reg294796)
reg294798 := PrimIsPair(reg294797)
var reg294844 Obj
if reg294798 == True {
reg294799 := PrimHead(V1593)
reg294800 := PrimTail(reg294799)
reg294801 := PrimTail(reg294800)
reg294802 := PrimIsPair(reg294801)
var reg294839 Obj
if reg294802 == True {
reg294803 := Nil;
reg294804 := PrimHead(V1593)
reg294805 := PrimTail(reg294804)
reg294806 := PrimTail(reg294805)
reg294807 := PrimTail(reg294806)
reg294808 := PrimEqual(reg294803, reg294807)
var reg294834 Obj
if reg294808 == True {
reg294809 := PrimTail(V1593)
reg294810 := PrimIsPair(reg294809)
var reg294829 Obj
if reg294810 == True {
reg294811 := Nil;
reg294812 := PrimTail(V1593)
reg294813 := PrimTail(reg294812)
reg294814 := PrimEqual(reg294811, reg294813)
var reg294824 Obj
if reg294814 == True {
reg294815 := PrimHead(V1593)
reg294816 := PrimTail(reg294815)
reg294817 := PrimHead(reg294816)
reg294818 := PrimIsVariable(reg294817)
reg294819 := PrimNot(reg294818)
var reg294822 Obj
if reg294819 == True {
reg294820 := True;
reg294822 = reg294820
} else {
reg294821 := False;
reg294822 = reg294821
}
reg294824 = reg294822
} else {
reg294823 := False;
reg294824 = reg294823
}
var reg294827 Obj
if reg294824 == True {
reg294825 := True;
reg294827 = reg294825
} else {
reg294826 := False;
reg294827 = reg294826
}
reg294829 = reg294827
} else {
reg294828 := False;
reg294829 = reg294828
}
var reg294832 Obj
if reg294829 == True {
reg294830 := True;
reg294832 = reg294830
} else {
reg294831 := False;
reg294832 = reg294831
}
reg294834 = reg294832
} else {
reg294833 := False;
reg294834 = reg294833
}
var reg294837 Obj
if reg294834 == True {
reg294835 := True;
reg294837 = reg294835
} else {
reg294836 := False;
reg294837 = reg294836
}
reg294839 = reg294837
} else {
reg294838 := False;
reg294839 = reg294838
}
var reg294842 Obj
if reg294839 == True {
reg294840 := True;
reg294842 = reg294840
} else {
reg294841 := False;
reg294842 = reg294841
}
reg294844 = reg294842
} else {
reg294843 := False;
reg294844 = reg294843
}
var reg294847 Obj
if reg294844 == True {
reg294845 := True;
reg294847 = reg294845
} else {
reg294846 := False;
reg294847 = reg294846
}
reg294849 = reg294847
} else {
reg294848 := False;
reg294849 = reg294848
}
var reg294852 Obj
if reg294849 == True {
reg294850 := True;
reg294852 = reg294850
} else {
reg294851 := False;
reg294852 = reg294851
}
reg294854 = reg294852
} else {
reg294853 := False;
reg294854 = reg294853
}
var reg294857 Obj
if reg294854 == True {
reg294855 := True;
reg294857 = reg294855
} else {
reg294856 := False;
reg294857 = reg294856
}
reg294859 = reg294857
} else {
reg294858 := False;
reg294859 = reg294858
}
if reg294859 == True {
reg294860 := MakeSymbol("=")
reg294861 := PrimHead(V1593)
reg294862 := PrimTail(reg294861)
reg294863 := PrimHead(reg294862)
reg294864 := PrimTail(V1593)
reg294865 := PrimCons(reg294863, reg294864)
reg294866 := PrimCons(reg294860, reg294865)
reg294867 := __e.Call(__defun__shen_4add__test, reg294866)
_ = reg294867
reg294868 := PrimHead(V1593)
reg294869 := PrimTail(reg294868)
reg294870 := PrimTail(reg294869)
reg294871 := PrimHead(reg294870)
__ctx.TailApply(__defun__shen_4reduce__help, reg294871)
return
} else {
reg294873 := PrimIsPair(V1593)
var reg294933 Obj
if reg294873 == True {
reg294874 := PrimHead(V1593)
reg294875 := PrimIsPair(reg294874)
var reg294928 Obj
if reg294875 == True {
reg294876 := MakeSymbol("/.")
reg294877 := PrimHead(V1593)
reg294878 := PrimHead(reg294877)
reg294879 := PrimEqual(reg294876, reg294878)
var reg294923 Obj
if reg294879 == True {
reg294880 := PrimHead(V1593)
reg294881 := PrimTail(reg294880)
reg294882 := PrimIsPair(reg294881)
var reg294918 Obj
if reg294882 == True {
reg294883 := PrimHead(V1593)
reg294884 := PrimTail(reg294883)
reg294885 := PrimTail(reg294884)
reg294886 := PrimIsPair(reg294885)
var reg294913 Obj
if reg294886 == True {
reg294887 := Nil;
reg294888 := PrimHead(V1593)
reg294889 := PrimTail(reg294888)
reg294890 := PrimTail(reg294889)
reg294891 := PrimTail(reg294890)
reg294892 := PrimEqual(reg294887, reg294891)
var reg294908 Obj
if reg294892 == True {
reg294893 := PrimTail(V1593)
reg294894 := PrimIsPair(reg294893)
var reg294903 Obj
if reg294894 == True {
reg294895 := Nil;
reg294896 := PrimTail(V1593)
reg294897 := PrimTail(reg294896)
reg294898 := PrimEqual(reg294895, reg294897)
var reg294901 Obj
if reg294898 == True {
reg294899 := True;
reg294901 = reg294899
} else {
reg294900 := False;
reg294901 = reg294900
}
reg294903 = reg294901
} else {
reg294902 := False;
reg294903 = reg294902
}
var reg294906 Obj
if reg294903 == True {
reg294904 := True;
reg294906 = reg294904
} else {
reg294905 := False;
reg294906 = reg294905
}
reg294908 = reg294906
} else {
reg294907 := False;
reg294908 = reg294907
}
var reg294911 Obj
if reg294908 == True {
reg294909 := True;
reg294911 = reg294909
} else {
reg294910 := False;
reg294911 = reg294910
}
reg294913 = reg294911
} else {
reg294912 := False;
reg294913 = reg294912
}
var reg294916 Obj
if reg294913 == True {
reg294914 := True;
reg294916 = reg294914
} else {
reg294915 := False;
reg294916 = reg294915
}
reg294918 = reg294916
} else {
reg294917 := False;
reg294918 = reg294917
}
var reg294921 Obj
if reg294918 == True {
reg294919 := True;
reg294921 = reg294919
} else {
reg294920 := False;
reg294921 = reg294920
}
reg294923 = reg294921
} else {
reg294922 := False;
reg294923 = reg294922
}
var reg294926 Obj
if reg294923 == True {
reg294924 := True;
reg294926 = reg294924
} else {
reg294925 := False;
reg294926 = reg294925
}
reg294928 = reg294926
} else {
reg294927 := False;
reg294928 = reg294927
}
var reg294931 Obj
if reg294928 == True {
reg294929 := True;
reg294931 = reg294929
} else {
reg294930 := False;
reg294931 = reg294930
}
reg294933 = reg294931
} else {
reg294932 := False;
reg294933 = reg294932
}
if reg294933 == True {
reg294934 := PrimTail(V1593)
reg294935 := PrimHead(reg294934)
reg294936 := PrimHead(V1593)
reg294937 := PrimTail(reg294936)
reg294938 := PrimHead(reg294937)
reg294939 := PrimHead(V1593)
reg294940 := PrimTail(reg294939)
reg294941 := PrimTail(reg294940)
reg294942 := PrimHead(reg294941)
reg294943 := __e.Call(__defun__shen_4ebr, reg294935, reg294938, reg294942)
__ctx.TailApply(__defun__shen_4reduce__help, reg294943)
return
} else {
reg294945 := PrimIsPair(V1593)
var reg294978 Obj
if reg294945 == True {
reg294946 := MakeSymbol("where")
reg294947 := PrimHead(V1593)
reg294948 := PrimEqual(reg294946, reg294947)
var reg294973 Obj
if reg294948 == True {
reg294949 := PrimTail(V1593)
reg294950 := PrimIsPair(reg294949)
var reg294968 Obj
if reg294950 == True {
reg294951 := PrimTail(V1593)
reg294952 := PrimTail(reg294951)
reg294953 := PrimIsPair(reg294952)
var reg294963 Obj
if reg294953 == True {
reg294954 := Nil;
reg294955 := PrimTail(V1593)
reg294956 := PrimTail(reg294955)
reg294957 := PrimTail(reg294956)
reg294958 := PrimEqual(reg294954, reg294957)
var reg294961 Obj
if reg294958 == True {
reg294959 := True;
reg294961 = reg294959
} else {
reg294960 := False;
reg294961 = reg294960
}
reg294963 = reg294961
} else {
reg294962 := False;
reg294963 = reg294962
}
var reg294966 Obj
if reg294963 == True {
reg294964 := True;
reg294966 = reg294964
} else {
reg294965 := False;
reg294966 = reg294965
}
reg294968 = reg294966
} else {
reg294967 := False;
reg294968 = reg294967
}
var reg294971 Obj
if reg294968 == True {
reg294969 := True;
reg294971 = reg294969
} else {
reg294970 := False;
reg294971 = reg294970
}
reg294973 = reg294971
} else {
reg294972 := False;
reg294973 = reg294972
}
var reg294976 Obj
if reg294973 == True {
reg294974 := True;
reg294976 = reg294974
} else {
reg294975 := False;
reg294976 = reg294975
}
reg294978 = reg294976
} else {
reg294977 := False;
reg294978 = reg294977
}
if reg294978 == True {
reg294979 := PrimTail(V1593)
reg294980 := PrimHead(reg294979)
reg294981 := __e.Call(__defun__shen_4add__test, reg294980)
_ = reg294981
reg294982 := PrimTail(V1593)
reg294983 := PrimTail(reg294982)
reg294984 := PrimHead(reg294983)
__ctx.TailApply(__defun__shen_4reduce__help, reg294984)
return
} else {
reg294986 := PrimIsPair(V1593)
var reg295002 Obj
if reg294986 == True {
reg294987 := PrimTail(V1593)
reg294988 := PrimIsPair(reg294987)
var reg294997 Obj
if reg294988 == True {
reg294989 := Nil;
reg294990 := PrimTail(V1593)
reg294991 := PrimTail(reg294990)
reg294992 := PrimEqual(reg294989, reg294991)
var reg294995 Obj
if reg294992 == True {
reg294993 := True;
reg294995 = reg294993
} else {
reg294994 := False;
reg294995 = reg294994
}
reg294997 = reg294995
} else {
reg294996 := False;
reg294997 = reg294996
}
var reg295000 Obj
if reg294997 == True {
reg294998 := True;
reg295000 = reg294998
} else {
reg294999 := False;
reg295000 = reg294999
}
reg295002 = reg295000
} else {
reg295001 := False;
reg295002 = reg295001
}
if reg295002 == True {
reg295003 := PrimHead(V1593)
reg295004 := __e.Call(__defun__shen_4reduce__help, reg295003)
Z := reg295004
_ = Z
reg295005 := PrimHead(V1593)
reg295006 := PrimEqual(reg295005, Z)
if reg295006 == True {
__ctx.Return(V1593)
return
} else {
reg295007 := PrimTail(V1593)
reg295008 := PrimCons(Z, reg295007)
__ctx.TailApply(__defun__shen_4reduce__help, reg295008)
return
}
} else {
__ctx.Return(V1593)
return
}
}
}
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.reduce_help", value: __defun__shen_4reduce__help})

__defun__shen_4_7string_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1595 := __args[0]
_ = V1595
reg295010 := MakeString("")
reg295011 := PrimEqual(reg295010, V1595)
if reg295011 == True {
reg295012 := False;
__ctx.Return(reg295012)
return
} else {
reg295013 := PrimIsString(V1595)
__ctx.Return(reg295013)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.+string?", value: __defun__shen_4_7string_2})

__defun__shen_4_7vector_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1597 := __args[0]
_ = V1597
reg295014 := PrimIsVector(V1597)
if reg295014 == True {
reg295015 := MakeNumber(0)
reg295016 := PrimVectorGet(V1597, reg295015)
reg295017 := MakeNumber(0)
reg295018 := PrimGreatThan(reg295016, reg295017)
if reg295018 == True {
reg295019 := True;
__ctx.Return(reg295019)
return
} else {
reg295020 := False;
__ctx.Return(reg295020)
return
}
} else {
reg295021 := False;
__ctx.Return(reg295021)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.+vector?", value: __defun__shen_4_7vector_2})

__defun__shen_4ebr = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1611 := __args[0]
_ = V1611
V1612 := __args[1]
_ = V1612
V1613 := __args[2]
_ = V1613
reg295022 := PrimEqual(V1613, V1612)
if reg295022 == True {
__ctx.Return(V1611)
return
} else {
reg295023 := PrimIsPair(V1613)
var reg295066 Obj
if reg295023 == True {
reg295024 := MakeSymbol("/.")
reg295025 := PrimHead(V1613)
reg295026 := PrimEqual(reg295024, reg295025)
var reg295061 Obj
if reg295026 == True {
reg295027 := PrimTail(V1613)
reg295028 := PrimIsPair(reg295027)
var reg295056 Obj
if reg295028 == True {
reg295029 := PrimTail(V1613)
reg295030 := PrimTail(reg295029)
reg295031 := PrimIsPair(reg295030)
var reg295051 Obj
if reg295031 == True {
reg295032 := Nil;
reg295033 := PrimTail(V1613)
reg295034 := PrimTail(reg295033)
reg295035 := PrimTail(reg295034)
reg295036 := PrimEqual(reg295032, reg295035)
var reg295046 Obj
if reg295036 == True {
reg295037 := PrimTail(V1613)
reg295038 := PrimHead(reg295037)
reg295039 := __e.Call(__defun__occurrences, V1612, reg295038)
reg295040 := MakeNumber(0)
reg295041 := PrimGreatThan(reg295039, reg295040)
var reg295044 Obj
if reg295041 == True {
reg295042 := True;
reg295044 = reg295042
} else {
reg295043 := False;
reg295044 = reg295043
}
reg295046 = reg295044
} else {
reg295045 := False;
reg295046 = reg295045
}
var reg295049 Obj
if reg295046 == True {
reg295047 := True;
reg295049 = reg295047
} else {
reg295048 := False;
reg295049 = reg295048
}
reg295051 = reg295049
} else {
reg295050 := False;
reg295051 = reg295050
}
var reg295054 Obj
if reg295051 == True {
reg295052 := True;
reg295054 = reg295052
} else {
reg295053 := False;
reg295054 = reg295053
}
reg295056 = reg295054
} else {
reg295055 := False;
reg295056 = reg295055
}
var reg295059 Obj
if reg295056 == True {
reg295057 := True;
reg295059 = reg295057
} else {
reg295058 := False;
reg295059 = reg295058
}
reg295061 = reg295059
} else {
reg295060 := False;
reg295061 = reg295060
}
var reg295064 Obj
if reg295061 == True {
reg295062 := True;
reg295064 = reg295062
} else {
reg295063 := False;
reg295064 = reg295063
}
reg295066 = reg295064
} else {
reg295065 := False;
reg295066 = reg295065
}
if reg295066 == True {
__ctx.Return(V1613)
return
} else {
reg295067 := PrimIsPair(V1613)
var reg295110 Obj
if reg295067 == True {
reg295068 := MakeSymbol("lambda")
reg295069 := PrimHead(V1613)
reg295070 := PrimEqual(reg295068, reg295069)
var reg295105 Obj
if reg295070 == True {
reg295071 := PrimTail(V1613)
reg295072 := PrimIsPair(reg295071)
var reg295100 Obj
if reg295072 == True {
reg295073 := PrimTail(V1613)
reg295074 := PrimTail(reg295073)
reg295075 := PrimIsPair(reg295074)
var reg295095 Obj
if reg295075 == True {
reg295076 := Nil;
reg295077 := PrimTail(V1613)
reg295078 := PrimTail(reg295077)
reg295079 := PrimTail(reg295078)
reg295080 := PrimEqual(reg295076, reg295079)
var reg295090 Obj
if reg295080 == True {
reg295081 := PrimTail(V1613)
reg295082 := PrimHead(reg295081)
reg295083 := __e.Call(__defun__occurrences, V1612, reg295082)
reg295084 := MakeNumber(0)
reg295085 := PrimGreatThan(reg295083, reg295084)
var reg295088 Obj
if reg295085 == True {
reg295086 := True;
reg295088 = reg295086
} else {
reg295087 := False;
reg295088 = reg295087
}
reg295090 = reg295088
} else {
reg295089 := False;
reg295090 = reg295089
}
var reg295093 Obj
if reg295090 == True {
reg295091 := True;
reg295093 = reg295091
} else {
reg295092 := False;
reg295093 = reg295092
}
reg295095 = reg295093
} else {
reg295094 := False;
reg295095 = reg295094
}
var reg295098 Obj
if reg295095 == True {
reg295096 := True;
reg295098 = reg295096
} else {
reg295097 := False;
reg295098 = reg295097
}
reg295100 = reg295098
} else {
reg295099 := False;
reg295100 = reg295099
}
var reg295103 Obj
if reg295100 == True {
reg295101 := True;
reg295103 = reg295101
} else {
reg295102 := False;
reg295103 = reg295102
}
reg295105 = reg295103
} else {
reg295104 := False;
reg295105 = reg295104
}
var reg295108 Obj
if reg295105 == True {
reg295106 := True;
reg295108 = reg295106
} else {
reg295107 := False;
reg295108 = reg295107
}
reg295110 = reg295108
} else {
reg295109 := False;
reg295110 = reg295109
}
if reg295110 == True {
__ctx.Return(V1613)
return
} else {
reg295111 := PrimIsPair(V1613)
var reg295162 Obj
if reg295111 == True {
reg295112 := MakeSymbol("let")
reg295113 := PrimHead(V1613)
reg295114 := PrimEqual(reg295112, reg295113)
var reg295157 Obj
if reg295114 == True {
reg295115 := PrimTail(V1613)
reg295116 := PrimIsPair(reg295115)
var reg295152 Obj
if reg295116 == True {
reg295117 := PrimTail(V1613)
reg295118 := PrimTail(reg295117)
reg295119 := PrimIsPair(reg295118)
var reg295147 Obj
if reg295119 == True {
reg295120 := PrimTail(V1613)
reg295121 := PrimTail(reg295120)
reg295122 := PrimTail(reg295121)
reg295123 := PrimIsPair(reg295122)
var reg295142 Obj
if reg295123 == True {
reg295124 := Nil;
reg295125 := PrimTail(V1613)
reg295126 := PrimTail(reg295125)
reg295127 := PrimTail(reg295126)
reg295128 := PrimTail(reg295127)
reg295129 := PrimEqual(reg295124, reg295128)
var reg295137 Obj
if reg295129 == True {
reg295130 := PrimTail(V1613)
reg295131 := PrimHead(reg295130)
reg295132 := PrimEqual(reg295131, V1612)
var reg295135 Obj
if reg295132 == True {
reg295133 := True;
reg295135 = reg295133
} else {
reg295134 := False;
reg295135 = reg295134
}
reg295137 = reg295135
} else {
reg295136 := False;
reg295137 = reg295136
}
var reg295140 Obj
if reg295137 == True {
reg295138 := True;
reg295140 = reg295138
} else {
reg295139 := False;
reg295140 = reg295139
}
reg295142 = reg295140
} else {
reg295141 := False;
reg295142 = reg295141
}
var reg295145 Obj
if reg295142 == True {
reg295143 := True;
reg295145 = reg295143
} else {
reg295144 := False;
reg295145 = reg295144
}
reg295147 = reg295145
} else {
reg295146 := False;
reg295147 = reg295146
}
var reg295150 Obj
if reg295147 == True {
reg295148 := True;
reg295150 = reg295148
} else {
reg295149 := False;
reg295150 = reg295149
}
reg295152 = reg295150
} else {
reg295151 := False;
reg295152 = reg295151
}
var reg295155 Obj
if reg295152 == True {
reg295153 := True;
reg295155 = reg295153
} else {
reg295154 := False;
reg295155 = reg295154
}
reg295157 = reg295155
} else {
reg295156 := False;
reg295157 = reg295156
}
var reg295160 Obj
if reg295157 == True {
reg295158 := True;
reg295160 = reg295158
} else {
reg295159 := False;
reg295160 = reg295159
}
reg295162 = reg295160
} else {
reg295161 := False;
reg295162 = reg295161
}
if reg295162 == True {
reg295163 := MakeSymbol("let")
reg295164 := PrimTail(V1613)
reg295165 := PrimHead(reg295164)
reg295166 := PrimTail(V1613)
reg295167 := PrimHead(reg295166)
reg295168 := PrimTail(V1613)
reg295169 := PrimTail(reg295168)
reg295170 := PrimHead(reg295169)
reg295171 := __e.Call(__defun__shen_4ebr, V1611, reg295167, reg295170)
reg295172 := PrimTail(V1613)
reg295173 := PrimTail(reg295172)
reg295174 := PrimTail(reg295173)
reg295175 := PrimCons(reg295171, reg295174)
reg295176 := PrimCons(reg295165, reg295175)
reg295177 := PrimCons(reg295163, reg295176)
__ctx.Return(reg295177)
return
} else {
reg295178 := PrimIsPair(V1613)
if reg295178 == True {
reg295179 := PrimHead(V1613)
reg295180 := __e.Call(__defun__shen_4ebr, V1611, V1612, reg295179)
reg295181 := PrimTail(V1613)
reg295182 := __e.Call(__defun__shen_4ebr, V1611, V1612, reg295181)
reg295183 := PrimCons(reg295180, reg295182)
__ctx.Return(reg295183)
return
} else {
__ctx.Return(V1613)
return
}
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.ebr", value: __defun__shen_4ebr})

__defun__shen_4add__test = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1615 := __args[0]
_ = V1615
reg295184 := MakeSymbol("shen.*teststack*")
reg295185 := MakeSymbol("shen.*teststack*")
reg295186 := PrimValue(reg295185)
reg295187 := PrimCons(V1615, reg295186)
reg295188 := PrimSet(reg295184, reg295187)
__ctx.Return(reg295188)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.add_test", value: __defun__shen_4add__test})

__defun__shen_4cond_1expression = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1619 := __args[0]
_ = V1619
V1620 := __args[1]
_ = V1620
V1621 := __args[2]
_ = V1621
reg295189 := __e.Call(__defun__shen_4err_1condition, V1619)
Err := reg295189
_ = Err
reg295190 := __e.Call(__defun__shen_4case_1form, V1621, Err)
Cases := reg295190
_ = Cases
reg295191 := __e.Call(__defun__shen_4encode_1choices, Cases, V1619)
EncodeChoices := reg295191
_ = EncodeChoices
__ctx.TailApply(__defun__shen_4cond_1form, EncodeChoices)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.cond-expression", value: __defun__shen_4cond_1expression})

__defun__shen_4cond_1form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1625 := __args[0]
_ = V1625
reg295193 := PrimIsPair(V1625)
var reg295227 Obj
if reg295193 == True {
reg295194 := PrimHead(V1625)
reg295195 := PrimIsPair(reg295194)
var reg295222 Obj
if reg295195 == True {
reg295196 := True;
reg295197 := PrimHead(V1625)
reg295198 := PrimHead(reg295197)
reg295199 := PrimEqual(reg295196, reg295198)
var reg295217 Obj
if reg295199 == True {
reg295200 := PrimHead(V1625)
reg295201 := PrimTail(reg295200)
reg295202 := PrimIsPair(reg295201)
var reg295212 Obj
if reg295202 == True {
reg295203 := Nil;
reg295204 := PrimHead(V1625)
reg295205 := PrimTail(reg295204)
reg295206 := PrimTail(reg295205)
reg295207 := PrimEqual(reg295203, reg295206)
var reg295210 Obj
if reg295207 == True {
reg295208 := True;
reg295210 = reg295208
} else {
reg295209 := False;
reg295210 = reg295209
}
reg295212 = reg295210
} else {
reg295211 := False;
reg295212 = reg295211
}
var reg295215 Obj
if reg295212 == True {
reg295213 := True;
reg295215 = reg295213
} else {
reg295214 := False;
reg295215 = reg295214
}
reg295217 = reg295215
} else {
reg295216 := False;
reg295217 = reg295216
}
var reg295220 Obj
if reg295217 == True {
reg295218 := True;
reg295220 = reg295218
} else {
reg295219 := False;
reg295220 = reg295219
}
reg295222 = reg295220
} else {
reg295221 := False;
reg295222 = reg295221
}
var reg295225 Obj
if reg295222 == True {
reg295223 := True;
reg295225 = reg295223
} else {
reg295224 := False;
reg295225 = reg295224
}
reg295227 = reg295225
} else {
reg295226 := False;
reg295227 = reg295226
}
if reg295227 == True {
reg295228 := PrimHead(V1625)
reg295229 := PrimTail(reg295228)
reg295230 := PrimHead(reg295229)
__ctx.Return(reg295230)
return
} else {
reg295231 := MakeSymbol("cond")
reg295232 := PrimCons(reg295231, V1625)
__ctx.Return(reg295232)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.cond-form", value: __defun__shen_4cond_1form})

__defun__shen_4encode_1choices = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1630 := __args[0]
_ = V1630
V1631 := __args[1]
_ = V1631
reg295233 := Nil;
reg295234 := PrimEqual(reg295233, V1630)
if reg295234 == True {
reg295235 := Nil;
__ctx.Return(reg295235)
return
} else {
reg295236 := PrimIsPair(V1630)
var reg295320 Obj
if reg295236 == True {
reg295237 := PrimHead(V1630)
reg295238 := PrimIsPair(reg295237)
var reg295315 Obj
if reg295238 == True {
reg295239 := True;
reg295240 := PrimHead(V1630)
reg295241 := PrimHead(reg295240)
reg295242 := PrimEqual(reg295239, reg295241)
var reg295310 Obj
if reg295242 == True {
reg295243 := PrimHead(V1630)
reg295244 := PrimTail(reg295243)
reg295245 := PrimIsPair(reg295244)
var reg295305 Obj
if reg295245 == True {
reg295246 := PrimHead(V1630)
reg295247 := PrimTail(reg295246)
reg295248 := PrimHead(reg295247)
reg295249 := PrimIsPair(reg295248)
var reg295300 Obj
if reg295249 == True {
reg295250 := MakeSymbol("shen.choicepoint!")
reg295251 := PrimHead(V1630)
reg295252 := PrimTail(reg295251)
reg295253 := PrimHead(reg295252)
reg295254 := PrimHead(reg295253)
reg295255 := PrimEqual(reg295250, reg295254)
var reg295295 Obj
if reg295255 == True {
reg295256 := PrimHead(V1630)
reg295257 := PrimTail(reg295256)
reg295258 := PrimHead(reg295257)
reg295259 := PrimTail(reg295258)
reg295260 := PrimIsPair(reg295259)
var reg295290 Obj
if reg295260 == True {
reg295261 := Nil;
reg295262 := PrimHead(V1630)
reg295263 := PrimTail(reg295262)
reg295264 := PrimHead(reg295263)
reg295265 := PrimTail(reg295264)
reg295266 := PrimTail(reg295265)
reg295267 := PrimEqual(reg295261, reg295266)
var reg295285 Obj
if reg295267 == True {
reg295268 := Nil;
reg295269 := PrimHead(V1630)
reg295270 := PrimTail(reg295269)
reg295271 := PrimTail(reg295270)
reg295272 := PrimEqual(reg295268, reg295271)
var reg295280 Obj
if reg295272 == True {
reg295273 := Nil;
reg295274 := PrimTail(V1630)
reg295275 := PrimEqual(reg295273, reg295274)
var reg295278 Obj
if reg295275 == True {
reg295276 := True;
reg295278 = reg295276
} else {
reg295277 := False;
reg295278 = reg295277
}
reg295280 = reg295278
} else {
reg295279 := False;
reg295280 = reg295279
}
var reg295283 Obj
if reg295280 == True {
reg295281 := True;
reg295283 = reg295281
} else {
reg295282 := False;
reg295283 = reg295282
}
reg295285 = reg295283
} else {
reg295284 := False;
reg295285 = reg295284
}
var reg295288 Obj
if reg295285 == True {
reg295286 := True;
reg295288 = reg295286
} else {
reg295287 := False;
reg295288 = reg295287
}
reg295290 = reg295288
} else {
reg295289 := False;
reg295290 = reg295289
}
var reg295293 Obj
if reg295290 == True {
reg295291 := True;
reg295293 = reg295291
} else {
reg295292 := False;
reg295293 = reg295292
}
reg295295 = reg295293
} else {
reg295294 := False;
reg295295 = reg295294
}
var reg295298 Obj
if reg295295 == True {
reg295296 := True;
reg295298 = reg295296
} else {
reg295297 := False;
reg295298 = reg295297
}
reg295300 = reg295298
} else {
reg295299 := False;
reg295300 = reg295299
}
var reg295303 Obj
if reg295300 == True {
reg295301 := True;
reg295303 = reg295301
} else {
reg295302 := False;
reg295303 = reg295302
}
reg295305 = reg295303
} else {
reg295304 := False;
reg295305 = reg295304
}
var reg295308 Obj
if reg295305 == True {
reg295306 := True;
reg295308 = reg295306
} else {
reg295307 := False;
reg295308 = reg295307
}
reg295310 = reg295308
} else {
reg295309 := False;
reg295310 = reg295309
}
var reg295313 Obj
if reg295310 == True {
reg295311 := True;
reg295313 = reg295311
} else {
reg295312 := False;
reg295313 = reg295312
}
reg295315 = reg295313
} else {
reg295314 := False;
reg295315 = reg295314
}
var reg295318 Obj
if reg295315 == True {
reg295316 := True;
reg295318 = reg295316
} else {
reg295317 := False;
reg295318 = reg295317
}
reg295320 = reg295318
} else {
reg295319 := False;
reg295320 = reg295319
}
if reg295320 == True {
reg295321 := True;
reg295322 := MakeSymbol("let")
reg295323 := MakeSymbol("Result")
reg295324 := PrimHead(V1630)
reg295325 := PrimTail(reg295324)
reg295326 := PrimHead(reg295325)
reg295327 := PrimTail(reg295326)
reg295328 := PrimHead(reg295327)
reg295329 := MakeSymbol("if")
reg295330 := MakeSymbol("=")
reg295331 := MakeSymbol("Result")
reg295332 := MakeSymbol("fail")
reg295333 := Nil;
reg295334 := PrimCons(reg295332, reg295333)
reg295335 := Nil;
reg295336 := PrimCons(reg295334, reg295335)
reg295337 := PrimCons(reg295331, reg295336)
reg295338 := PrimCons(reg295330, reg295337)
reg295339 := MakeSymbol("shen.*installing-kl*")
reg295340 := PrimValue(reg295339)
var reg295349 Obj
if reg295340 == True {
reg295341 := MakeSymbol("shen.sys-error")
reg295342 := Nil;
reg295343 := PrimCons(V1631, reg295342)
reg295344 := PrimCons(reg295341, reg295343)
reg295349 = reg295344
} else {
reg295345 := MakeSymbol("shen.f_error")
reg295346 := Nil;
reg295347 := PrimCons(V1631, reg295346)
reg295348 := PrimCons(reg295345, reg295347)
reg295349 = reg295348
}
reg295350 := MakeSymbol("Result")
reg295351 := Nil;
reg295352 := PrimCons(reg295350, reg295351)
reg295353 := PrimCons(reg295349, reg295352)
reg295354 := PrimCons(reg295338, reg295353)
reg295355 := PrimCons(reg295329, reg295354)
reg295356 := Nil;
reg295357 := PrimCons(reg295355, reg295356)
reg295358 := PrimCons(reg295328, reg295357)
reg295359 := PrimCons(reg295323, reg295358)
reg295360 := PrimCons(reg295322, reg295359)
reg295361 := Nil;
reg295362 := PrimCons(reg295360, reg295361)
reg295363 := PrimCons(reg295321, reg295362)
reg295364 := Nil;
reg295365 := PrimCons(reg295363, reg295364)
__ctx.Return(reg295365)
return
} else {
reg295366 := PrimIsPair(V1630)
var reg295442 Obj
if reg295366 == True {
reg295367 := PrimHead(V1630)
reg295368 := PrimIsPair(reg295367)
var reg295437 Obj
if reg295368 == True {
reg295369 := True;
reg295370 := PrimHead(V1630)
reg295371 := PrimHead(reg295370)
reg295372 := PrimEqual(reg295369, reg295371)
var reg295432 Obj
if reg295372 == True {
reg295373 := PrimHead(V1630)
reg295374 := PrimTail(reg295373)
reg295375 := PrimIsPair(reg295374)
var reg295427 Obj
if reg295375 == True {
reg295376 := PrimHead(V1630)
reg295377 := PrimTail(reg295376)
reg295378 := PrimHead(reg295377)
reg295379 := PrimIsPair(reg295378)
var reg295422 Obj
if reg295379 == True {
reg295380 := MakeSymbol("shen.choicepoint!")
reg295381 := PrimHead(V1630)
reg295382 := PrimTail(reg295381)
reg295383 := PrimHead(reg295382)
reg295384 := PrimHead(reg295383)
reg295385 := PrimEqual(reg295380, reg295384)
var reg295417 Obj
if reg295385 == True {
reg295386 := PrimHead(V1630)
reg295387 := PrimTail(reg295386)
reg295388 := PrimHead(reg295387)
reg295389 := PrimTail(reg295388)
reg295390 := PrimIsPair(reg295389)
var reg295412 Obj
if reg295390 == True {
reg295391 := Nil;
reg295392 := PrimHead(V1630)
reg295393 := PrimTail(reg295392)
reg295394 := PrimHead(reg295393)
reg295395 := PrimTail(reg295394)
reg295396 := PrimTail(reg295395)
reg295397 := PrimEqual(reg295391, reg295396)
var reg295407 Obj
if reg295397 == True {
reg295398 := Nil;
reg295399 := PrimHead(V1630)
reg295400 := PrimTail(reg295399)
reg295401 := PrimTail(reg295400)
reg295402 := PrimEqual(reg295398, reg295401)
var reg295405 Obj
if reg295402 == True {
reg295403 := True;
reg295405 = reg295403
} else {
reg295404 := False;
reg295405 = reg295404
}
reg295407 = reg295405
} else {
reg295406 := False;
reg295407 = reg295406
}
var reg295410 Obj
if reg295407 == True {
reg295408 := True;
reg295410 = reg295408
} else {
reg295409 := False;
reg295410 = reg295409
}
reg295412 = reg295410
} else {
reg295411 := False;
reg295412 = reg295411
}
var reg295415 Obj
if reg295412 == True {
reg295413 := True;
reg295415 = reg295413
} else {
reg295414 := False;
reg295415 = reg295414
}
reg295417 = reg295415
} else {
reg295416 := False;
reg295417 = reg295416
}
var reg295420 Obj
if reg295417 == True {
reg295418 := True;
reg295420 = reg295418
} else {
reg295419 := False;
reg295420 = reg295419
}
reg295422 = reg295420
} else {
reg295421 := False;
reg295422 = reg295421
}
var reg295425 Obj
if reg295422 == True {
reg295423 := True;
reg295425 = reg295423
} else {
reg295424 := False;
reg295425 = reg295424
}
reg295427 = reg295425
} else {
reg295426 := False;
reg295427 = reg295426
}
var reg295430 Obj
if reg295427 == True {
reg295428 := True;
reg295430 = reg295428
} else {
reg295429 := False;
reg295430 = reg295429
}
reg295432 = reg295430
} else {
reg295431 := False;
reg295432 = reg295431
}
var reg295435 Obj
if reg295432 == True {
reg295433 := True;
reg295435 = reg295433
} else {
reg295434 := False;
reg295435 = reg295434
}
reg295437 = reg295435
} else {
reg295436 := False;
reg295437 = reg295436
}
var reg295440 Obj
if reg295437 == True {
reg295438 := True;
reg295440 = reg295438
} else {
reg295439 := False;
reg295440 = reg295439
}
reg295442 = reg295440
} else {
reg295441 := False;
reg295442 = reg295441
}
if reg295442 == True {
reg295443 := True;
reg295444 := MakeSymbol("let")
reg295445 := MakeSymbol("Result")
reg295446 := PrimHead(V1630)
reg295447 := PrimTail(reg295446)
reg295448 := PrimHead(reg295447)
reg295449 := PrimTail(reg295448)
reg295450 := PrimHead(reg295449)
reg295451 := MakeSymbol("if")
reg295452 := MakeSymbol("=")
reg295453 := MakeSymbol("Result")
reg295454 := MakeSymbol("fail")
reg295455 := Nil;
reg295456 := PrimCons(reg295454, reg295455)
reg295457 := Nil;
reg295458 := PrimCons(reg295456, reg295457)
reg295459 := PrimCons(reg295453, reg295458)
reg295460 := PrimCons(reg295452, reg295459)
reg295461 := PrimTail(V1630)
reg295462 := __e.Call(__defun__shen_4encode_1choices, reg295461, V1631)
reg295463 := __e.Call(__defun__shen_4cond_1form, reg295462)
reg295464 := MakeSymbol("Result")
reg295465 := Nil;
reg295466 := PrimCons(reg295464, reg295465)
reg295467 := PrimCons(reg295463, reg295466)
reg295468 := PrimCons(reg295460, reg295467)
reg295469 := PrimCons(reg295451, reg295468)
reg295470 := Nil;
reg295471 := PrimCons(reg295469, reg295470)
reg295472 := PrimCons(reg295450, reg295471)
reg295473 := PrimCons(reg295445, reg295472)
reg295474 := PrimCons(reg295444, reg295473)
reg295475 := Nil;
reg295476 := PrimCons(reg295474, reg295475)
reg295477 := PrimCons(reg295443, reg295476)
reg295478 := Nil;
reg295479 := PrimCons(reg295477, reg295478)
__ctx.Return(reg295479)
return
} else {
reg295480 := PrimIsPair(V1630)
var reg295547 Obj
if reg295480 == True {
reg295481 := PrimHead(V1630)
reg295482 := PrimIsPair(reg295481)
var reg295542 Obj
if reg295482 == True {
reg295483 := PrimHead(V1630)
reg295484 := PrimTail(reg295483)
reg295485 := PrimIsPair(reg295484)
var reg295537 Obj
if reg295485 == True {
reg295486 := PrimHead(V1630)
reg295487 := PrimTail(reg295486)
reg295488 := PrimHead(reg295487)
reg295489 := PrimIsPair(reg295488)
var reg295532 Obj
if reg295489 == True {
reg295490 := MakeSymbol("shen.choicepoint!")
reg295491 := PrimHead(V1630)
reg295492 := PrimTail(reg295491)
reg295493 := PrimHead(reg295492)
reg295494 := PrimHead(reg295493)
reg295495 := PrimEqual(reg295490, reg295494)
var reg295527 Obj
if reg295495 == True {
reg295496 := PrimHead(V1630)
reg295497 := PrimTail(reg295496)
reg295498 := PrimHead(reg295497)
reg295499 := PrimTail(reg295498)
reg295500 := PrimIsPair(reg295499)
var reg295522 Obj
if reg295500 == True {
reg295501 := Nil;
reg295502 := PrimHead(V1630)
reg295503 := PrimTail(reg295502)
reg295504 := PrimHead(reg295503)
reg295505 := PrimTail(reg295504)
reg295506 := PrimTail(reg295505)
reg295507 := PrimEqual(reg295501, reg295506)
var reg295517 Obj
if reg295507 == True {
reg295508 := Nil;
reg295509 := PrimHead(V1630)
reg295510 := PrimTail(reg295509)
reg295511 := PrimTail(reg295510)
reg295512 := PrimEqual(reg295508, reg295511)
var reg295515 Obj
if reg295512 == True {
reg295513 := True;
reg295515 = reg295513
} else {
reg295514 := False;
reg295515 = reg295514
}
reg295517 = reg295515
} else {
reg295516 := False;
reg295517 = reg295516
}
var reg295520 Obj
if reg295517 == True {
reg295518 := True;
reg295520 = reg295518
} else {
reg295519 := False;
reg295520 = reg295519
}
reg295522 = reg295520
} else {
reg295521 := False;
reg295522 = reg295521
}
var reg295525 Obj
if reg295522 == True {
reg295523 := True;
reg295525 = reg295523
} else {
reg295524 := False;
reg295525 = reg295524
}
reg295527 = reg295525
} else {
reg295526 := False;
reg295527 = reg295526
}
var reg295530 Obj
if reg295527 == True {
reg295528 := True;
reg295530 = reg295528
} else {
reg295529 := False;
reg295530 = reg295529
}
reg295532 = reg295530
} else {
reg295531 := False;
reg295532 = reg295531
}
var reg295535 Obj
if reg295532 == True {
reg295533 := True;
reg295535 = reg295533
} else {
reg295534 := False;
reg295535 = reg295534
}
reg295537 = reg295535
} else {
reg295536 := False;
reg295537 = reg295536
}
var reg295540 Obj
if reg295537 == True {
reg295538 := True;
reg295540 = reg295538
} else {
reg295539 := False;
reg295540 = reg295539
}
reg295542 = reg295540
} else {
reg295541 := False;
reg295542 = reg295541
}
var reg295545 Obj
if reg295542 == True {
reg295543 := True;
reg295545 = reg295543
} else {
reg295544 := False;
reg295545 = reg295544
}
reg295547 = reg295545
} else {
reg295546 := False;
reg295547 = reg295546
}
if reg295547 == True {
reg295548 := True;
reg295549 := MakeSymbol("let")
reg295550 := MakeSymbol("Freeze")
reg295551 := MakeSymbol("freeze")
reg295552 := PrimTail(V1630)
reg295553 := __e.Call(__defun__shen_4encode_1choices, reg295552, V1631)
reg295554 := __e.Call(__defun__shen_4cond_1form, reg295553)
reg295555 := Nil;
reg295556 := PrimCons(reg295554, reg295555)
reg295557 := PrimCons(reg295551, reg295556)
reg295558 := MakeSymbol("if")
reg295559 := PrimHead(V1630)
reg295560 := PrimHead(reg295559)
reg295561 := MakeSymbol("let")
reg295562 := MakeSymbol("Result")
reg295563 := PrimHead(V1630)
reg295564 := PrimTail(reg295563)
reg295565 := PrimHead(reg295564)
reg295566 := PrimTail(reg295565)
reg295567 := PrimHead(reg295566)
reg295568 := MakeSymbol("if")
reg295569 := MakeSymbol("=")
reg295570 := MakeSymbol("Result")
reg295571 := MakeSymbol("fail")
reg295572 := Nil;
reg295573 := PrimCons(reg295571, reg295572)
reg295574 := Nil;
reg295575 := PrimCons(reg295573, reg295574)
reg295576 := PrimCons(reg295570, reg295575)
reg295577 := PrimCons(reg295569, reg295576)
reg295578 := MakeSymbol("thaw")
reg295579 := MakeSymbol("Freeze")
reg295580 := Nil;
reg295581 := PrimCons(reg295579, reg295580)
reg295582 := PrimCons(reg295578, reg295581)
reg295583 := MakeSymbol("Result")
reg295584 := Nil;
reg295585 := PrimCons(reg295583, reg295584)
reg295586 := PrimCons(reg295582, reg295585)
reg295587 := PrimCons(reg295577, reg295586)
reg295588 := PrimCons(reg295568, reg295587)
reg295589 := Nil;
reg295590 := PrimCons(reg295588, reg295589)
reg295591 := PrimCons(reg295567, reg295590)
reg295592 := PrimCons(reg295562, reg295591)
reg295593 := PrimCons(reg295561, reg295592)
reg295594 := MakeSymbol("thaw")
reg295595 := MakeSymbol("Freeze")
reg295596 := Nil;
reg295597 := PrimCons(reg295595, reg295596)
reg295598 := PrimCons(reg295594, reg295597)
reg295599 := Nil;
reg295600 := PrimCons(reg295598, reg295599)
reg295601 := PrimCons(reg295593, reg295600)
reg295602 := PrimCons(reg295560, reg295601)
reg295603 := PrimCons(reg295558, reg295602)
reg295604 := Nil;
reg295605 := PrimCons(reg295603, reg295604)
reg295606 := PrimCons(reg295557, reg295605)
reg295607 := PrimCons(reg295550, reg295606)
reg295608 := PrimCons(reg295549, reg295607)
reg295609 := Nil;
reg295610 := PrimCons(reg295608, reg295609)
reg295611 := PrimCons(reg295548, reg295610)
reg295612 := Nil;
reg295613 := PrimCons(reg295611, reg295612)
__ctx.Return(reg295613)
return
} else {
reg295614 := PrimIsPair(V1630)
var reg295639 Obj
if reg295614 == True {
reg295615 := PrimHead(V1630)
reg295616 := PrimIsPair(reg295615)
var reg295634 Obj
if reg295616 == True {
reg295617 := PrimHead(V1630)
reg295618 := PrimTail(reg295617)
reg295619 := PrimIsPair(reg295618)
var reg295629 Obj
if reg295619 == True {
reg295620 := Nil;
reg295621 := PrimHead(V1630)
reg295622 := PrimTail(reg295621)
reg295623 := PrimTail(reg295622)
reg295624 := PrimEqual(reg295620, reg295623)
var reg295627 Obj
if reg295624 == True {
reg295625 := True;
reg295627 = reg295625
} else {
reg295626 := False;
reg295627 = reg295626
}
reg295629 = reg295627
} else {
reg295628 := False;
reg295629 = reg295628
}
var reg295632 Obj
if reg295629 == True {
reg295630 := True;
reg295632 = reg295630
} else {
reg295631 := False;
reg295632 = reg295631
}
reg295634 = reg295632
} else {
reg295633 := False;
reg295634 = reg295633
}
var reg295637 Obj
if reg295634 == True {
reg295635 := True;
reg295637 = reg295635
} else {
reg295636 := False;
reg295637 = reg295636
}
reg295639 = reg295637
} else {
reg295638 := False;
reg295639 = reg295638
}
if reg295639 == True {
reg295640 := PrimHead(V1630)
reg295641 := PrimTail(V1630)
reg295642 := __e.Call(__defun__shen_4encode_1choices, reg295641, V1631)
reg295643 := PrimCons(reg295640, reg295642)
__ctx.Return(reg295643)
return
} else {
reg295644 := MakeSymbol("shen.encode-choices")
__ctx.TailApply(__defun__shen_4f__error, reg295644)
return
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.encode-choices", value: __defun__shen_4encode_1choices})

__defun__shen_4case_1form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1638 := __args[0]
_ = V1638
V1639 := __args[1]
_ = V1639
reg295646 := Nil;
reg295647 := PrimEqual(reg295646, V1638)
if reg295647 == True {
reg295648 := Nil;
reg295649 := PrimCons(V1639, reg295648)
__ctx.Return(reg295649)
return
} else {
reg295650 := PrimIsPair(V1638)
var reg295766 Obj
if reg295650 == True {
reg295651 := PrimHead(V1638)
reg295652 := PrimIsPair(reg295651)
var reg295761 Obj
if reg295652 == True {
reg295653 := PrimHead(V1638)
reg295654 := PrimHead(reg295653)
reg295655 := PrimIsPair(reg295654)
var reg295756 Obj
if reg295655 == True {
reg295656 := MakeSymbol(":")
reg295657 := PrimHead(V1638)
reg295658 := PrimHead(reg295657)
reg295659 := PrimHead(reg295658)
reg295660 := PrimEqual(reg295656, reg295659)
var reg295751 Obj
if reg295660 == True {
reg295661 := PrimHead(V1638)
reg295662 := PrimHead(reg295661)
reg295663 := PrimTail(reg295662)
reg295664 := PrimIsPair(reg295663)
var reg295746 Obj
if reg295664 == True {
reg295665 := MakeSymbol("shen.tests")
reg295666 := PrimHead(V1638)
reg295667 := PrimHead(reg295666)
reg295668 := PrimTail(reg295667)
reg295669 := PrimHead(reg295668)
reg295670 := PrimEqual(reg295665, reg295669)
var reg295741 Obj
if reg295670 == True {
reg295671 := Nil;
reg295672 := PrimHead(V1638)
reg295673 := PrimHead(reg295672)
reg295674 := PrimTail(reg295673)
reg295675 := PrimTail(reg295674)
reg295676 := PrimEqual(reg295671, reg295675)
var reg295736 Obj
if reg295676 == True {
reg295677 := PrimHead(V1638)
reg295678 := PrimTail(reg295677)
reg295679 := PrimIsPair(reg295678)
var reg295731 Obj
if reg295679 == True {
reg295680 := PrimHead(V1638)
reg295681 := PrimTail(reg295680)
reg295682 := PrimHead(reg295681)
reg295683 := PrimIsPair(reg295682)
var reg295726 Obj
if reg295683 == True {
reg295684 := MakeSymbol("shen.choicepoint!")
reg295685 := PrimHead(V1638)
reg295686 := PrimTail(reg295685)
reg295687 := PrimHead(reg295686)
reg295688 := PrimHead(reg295687)
reg295689 := PrimEqual(reg295684, reg295688)
var reg295721 Obj
if reg295689 == True {
reg295690 := PrimHead(V1638)
reg295691 := PrimTail(reg295690)
reg295692 := PrimHead(reg295691)
reg295693 := PrimTail(reg295692)
reg295694 := PrimIsPair(reg295693)
var reg295716 Obj
if reg295694 == True {
reg295695 := Nil;
reg295696 := PrimHead(V1638)
reg295697 := PrimTail(reg295696)
reg295698 := PrimHead(reg295697)
reg295699 := PrimTail(reg295698)
reg295700 := PrimTail(reg295699)
reg295701 := PrimEqual(reg295695, reg295700)
var reg295711 Obj
if reg295701 == True {
reg295702 := Nil;
reg295703 := PrimHead(V1638)
reg295704 := PrimTail(reg295703)
reg295705 := PrimTail(reg295704)
reg295706 := PrimEqual(reg295702, reg295705)
var reg295709 Obj
if reg295706 == True {
reg295707 := True;
reg295709 = reg295707
} else {
reg295708 := False;
reg295709 = reg295708
}
reg295711 = reg295709
} else {
reg295710 := False;
reg295711 = reg295710
}
var reg295714 Obj
if reg295711 == True {
reg295712 := True;
reg295714 = reg295712
} else {
reg295713 := False;
reg295714 = reg295713
}
reg295716 = reg295714
} else {
reg295715 := False;
reg295716 = reg295715
}
var reg295719 Obj
if reg295716 == True {
reg295717 := True;
reg295719 = reg295717
} else {
reg295718 := False;
reg295719 = reg295718
}
reg295721 = reg295719
} else {
reg295720 := False;
reg295721 = reg295720
}
var reg295724 Obj
if reg295721 == True {
reg295722 := True;
reg295724 = reg295722
} else {
reg295723 := False;
reg295724 = reg295723
}
reg295726 = reg295724
} else {
reg295725 := False;
reg295726 = reg295725
}
var reg295729 Obj
if reg295726 == True {
reg295727 := True;
reg295729 = reg295727
} else {
reg295728 := False;
reg295729 = reg295728
}
reg295731 = reg295729
} else {
reg295730 := False;
reg295731 = reg295730
}
var reg295734 Obj
if reg295731 == True {
reg295732 := True;
reg295734 = reg295732
} else {
reg295733 := False;
reg295734 = reg295733
}
reg295736 = reg295734
} else {
reg295735 := False;
reg295736 = reg295735
}
var reg295739 Obj
if reg295736 == True {
reg295737 := True;
reg295739 = reg295737
} else {
reg295738 := False;
reg295739 = reg295738
}
reg295741 = reg295739
} else {
reg295740 := False;
reg295741 = reg295740
}
var reg295744 Obj
if reg295741 == True {
reg295742 := True;
reg295744 = reg295742
} else {
reg295743 := False;
reg295744 = reg295743
}
reg295746 = reg295744
} else {
reg295745 := False;
reg295746 = reg295745
}
var reg295749 Obj
if reg295746 == True {
reg295747 := True;
reg295749 = reg295747
} else {
reg295748 := False;
reg295749 = reg295748
}
reg295751 = reg295749
} else {
reg295750 := False;
reg295751 = reg295750
}
var reg295754 Obj
if reg295751 == True {
reg295752 := True;
reg295754 = reg295752
} else {
reg295753 := False;
reg295754 = reg295753
}
reg295756 = reg295754
} else {
reg295755 := False;
reg295756 = reg295755
}
var reg295759 Obj
if reg295756 == True {
reg295757 := True;
reg295759 = reg295757
} else {
reg295758 := False;
reg295759 = reg295758
}
reg295761 = reg295759
} else {
reg295760 := False;
reg295761 = reg295760
}
var reg295764 Obj
if reg295761 == True {
reg295762 := True;
reg295764 = reg295762
} else {
reg295763 := False;
reg295764 = reg295763
}
reg295766 = reg295764
} else {
reg295765 := False;
reg295766 = reg295765
}
if reg295766 == True {
reg295767 := True;
reg295768 := PrimHead(V1638)
reg295769 := PrimTail(reg295768)
reg295770 := PrimCons(reg295767, reg295769)
reg295771 := PrimTail(V1638)
reg295772 := __e.Call(__defun__shen_4case_1form, reg295771, V1639)
reg295773 := PrimCons(reg295770, reg295772)
__ctx.Return(reg295773)
return
} else {
reg295774 := PrimIsPair(V1638)
var reg295848 Obj
if reg295774 == True {
reg295775 := PrimHead(V1638)
reg295776 := PrimIsPair(reg295775)
var reg295843 Obj
if reg295776 == True {
reg295777 := PrimHead(V1638)
reg295778 := PrimHead(reg295777)
reg295779 := PrimIsPair(reg295778)
var reg295838 Obj
if reg295779 == True {
reg295780 := MakeSymbol(":")
reg295781 := PrimHead(V1638)
reg295782 := PrimHead(reg295781)
reg295783 := PrimHead(reg295782)
reg295784 := PrimEqual(reg295780, reg295783)
var reg295833 Obj
if reg295784 == True {
reg295785 := PrimHead(V1638)
reg295786 := PrimHead(reg295785)
reg295787 := PrimTail(reg295786)
reg295788 := PrimIsPair(reg295787)
var reg295828 Obj
if reg295788 == True {
reg295789 := MakeSymbol("shen.tests")
reg295790 := PrimHead(V1638)
reg295791 := PrimHead(reg295790)
reg295792 := PrimTail(reg295791)
reg295793 := PrimHead(reg295792)
reg295794 := PrimEqual(reg295789, reg295793)
var reg295823 Obj
if reg295794 == True {
reg295795 := Nil;
reg295796 := PrimHead(V1638)
reg295797 := PrimHead(reg295796)
reg295798 := PrimTail(reg295797)
reg295799 := PrimTail(reg295798)
reg295800 := PrimEqual(reg295795, reg295799)
var reg295818 Obj
if reg295800 == True {
reg295801 := PrimHead(V1638)
reg295802 := PrimTail(reg295801)
reg295803 := PrimIsPair(reg295802)
var reg295813 Obj
if reg295803 == True {
reg295804 := Nil;
reg295805 := PrimHead(V1638)
reg295806 := PrimTail(reg295805)
reg295807 := PrimTail(reg295806)
reg295808 := PrimEqual(reg295804, reg295807)
var reg295811 Obj
if reg295808 == True {
reg295809 := True;
reg295811 = reg295809
} else {
reg295810 := False;
reg295811 = reg295810
}
reg295813 = reg295811
} else {
reg295812 := False;
reg295813 = reg295812
}
var reg295816 Obj
if reg295813 == True {
reg295814 := True;
reg295816 = reg295814
} else {
reg295815 := False;
reg295816 = reg295815
}
reg295818 = reg295816
} else {
reg295817 := False;
reg295818 = reg295817
}
var reg295821 Obj
if reg295818 == True {
reg295819 := True;
reg295821 = reg295819
} else {
reg295820 := False;
reg295821 = reg295820
}
reg295823 = reg295821
} else {
reg295822 := False;
reg295823 = reg295822
}
var reg295826 Obj
if reg295823 == True {
reg295824 := True;
reg295826 = reg295824
} else {
reg295825 := False;
reg295826 = reg295825
}
reg295828 = reg295826
} else {
reg295827 := False;
reg295828 = reg295827
}
var reg295831 Obj
if reg295828 == True {
reg295829 := True;
reg295831 = reg295829
} else {
reg295830 := False;
reg295831 = reg295830
}
reg295833 = reg295831
} else {
reg295832 := False;
reg295833 = reg295832
}
var reg295836 Obj
if reg295833 == True {
reg295834 := True;
reg295836 = reg295834
} else {
reg295835 := False;
reg295836 = reg295835
}
reg295838 = reg295836
} else {
reg295837 := False;
reg295838 = reg295837
}
var reg295841 Obj
if reg295838 == True {
reg295839 := True;
reg295841 = reg295839
} else {
reg295840 := False;
reg295841 = reg295840
}
reg295843 = reg295841
} else {
reg295842 := False;
reg295843 = reg295842
}
var reg295846 Obj
if reg295843 == True {
reg295844 := True;
reg295846 = reg295844
} else {
reg295845 := False;
reg295846 = reg295845
}
reg295848 = reg295846
} else {
reg295847 := False;
reg295848 = reg295847
}
if reg295848 == True {
reg295849 := True;
reg295850 := PrimHead(V1638)
reg295851 := PrimTail(reg295850)
reg295852 := PrimCons(reg295849, reg295851)
reg295853 := Nil;
reg295854 := PrimCons(reg295852, reg295853)
__ctx.Return(reg295854)
return
} else {
reg295855 := PrimIsPair(V1638)
var reg295918 Obj
if reg295855 == True {
reg295856 := PrimHead(V1638)
reg295857 := PrimIsPair(reg295856)
var reg295913 Obj
if reg295857 == True {
reg295858 := PrimHead(V1638)
reg295859 := PrimHead(reg295858)
reg295860 := PrimIsPair(reg295859)
var reg295908 Obj
if reg295860 == True {
reg295861 := MakeSymbol(":")
reg295862 := PrimHead(V1638)
reg295863 := PrimHead(reg295862)
reg295864 := PrimHead(reg295863)
reg295865 := PrimEqual(reg295861, reg295864)
var reg295903 Obj
if reg295865 == True {
reg295866 := PrimHead(V1638)
reg295867 := PrimHead(reg295866)
reg295868 := PrimTail(reg295867)
reg295869 := PrimIsPair(reg295868)
var reg295898 Obj
if reg295869 == True {
reg295870 := MakeSymbol("shen.tests")
reg295871 := PrimHead(V1638)
reg295872 := PrimHead(reg295871)
reg295873 := PrimTail(reg295872)
reg295874 := PrimHead(reg295873)
reg295875 := PrimEqual(reg295870, reg295874)
var reg295893 Obj
if reg295875 == True {
reg295876 := PrimHead(V1638)
reg295877 := PrimTail(reg295876)
reg295878 := PrimIsPair(reg295877)
var reg295888 Obj
if reg295878 == True {
reg295879 := Nil;
reg295880 := PrimHead(V1638)
reg295881 := PrimTail(reg295880)
reg295882 := PrimTail(reg295881)
reg295883 := PrimEqual(reg295879, reg295882)
var reg295886 Obj
if reg295883 == True {
reg295884 := True;
reg295886 = reg295884
} else {
reg295885 := False;
reg295886 = reg295885
}
reg295888 = reg295886
} else {
reg295887 := False;
reg295888 = reg295887
}
var reg295891 Obj
if reg295888 == True {
reg295889 := True;
reg295891 = reg295889
} else {
reg295890 := False;
reg295891 = reg295890
}
reg295893 = reg295891
} else {
reg295892 := False;
reg295893 = reg295892
}
var reg295896 Obj
if reg295893 == True {
reg295894 := True;
reg295896 = reg295894
} else {
reg295895 := False;
reg295896 = reg295895
}
reg295898 = reg295896
} else {
reg295897 := False;
reg295898 = reg295897
}
var reg295901 Obj
if reg295898 == True {
reg295899 := True;
reg295901 = reg295899
} else {
reg295900 := False;
reg295901 = reg295900
}
reg295903 = reg295901
} else {
reg295902 := False;
reg295903 = reg295902
}
var reg295906 Obj
if reg295903 == True {
reg295904 := True;
reg295906 = reg295904
} else {
reg295905 := False;
reg295906 = reg295905
}
reg295908 = reg295906
} else {
reg295907 := False;
reg295908 = reg295907
}
var reg295911 Obj
if reg295908 == True {
reg295909 := True;
reg295911 = reg295909
} else {
reg295910 := False;
reg295911 = reg295910
}
reg295913 = reg295911
} else {
reg295912 := False;
reg295913 = reg295912
}
var reg295916 Obj
if reg295913 == True {
reg295914 := True;
reg295916 = reg295914
} else {
reg295915 := False;
reg295916 = reg295915
}
reg295918 = reg295916
} else {
reg295917 := False;
reg295918 = reg295917
}
if reg295918 == True {
reg295919 := PrimHead(V1638)
reg295920 := PrimHead(reg295919)
reg295921 := PrimTail(reg295920)
reg295922 := PrimTail(reg295921)
reg295923 := __e.Call(__defun__shen_4embed_1and, reg295922)
reg295924 := PrimHead(V1638)
reg295925 := PrimTail(reg295924)
reg295926 := PrimCons(reg295923, reg295925)
reg295927 := PrimTail(V1638)
reg295928 := __e.Call(__defun__shen_4case_1form, reg295927, V1639)
reg295929 := PrimCons(reg295926, reg295928)
__ctx.Return(reg295929)
return
} else {
reg295930 := MakeSymbol("shen.case-form")
__ctx.TailApply(__defun__shen_4f__error, reg295930)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.case-form", value: __defun__shen_4case_1form})

__defun__shen_4embed_1and = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1641 := __args[0]
_ = V1641
reg295932 := PrimIsPair(V1641)
var reg295940 Obj
if reg295932 == True {
reg295933 := Nil;
reg295934 := PrimTail(V1641)
reg295935 := PrimEqual(reg295933, reg295934)
var reg295938 Obj
if reg295935 == True {
reg295936 := True;
reg295938 = reg295936
} else {
reg295937 := False;
reg295938 = reg295937
}
reg295940 = reg295938
} else {
reg295939 := False;
reg295940 = reg295939
}
if reg295940 == True {
reg295941 := PrimHead(V1641)
__ctx.Return(reg295941)
return
} else {
reg295942 := PrimIsPair(V1641)
if reg295942 == True {
reg295943 := MakeSymbol("and")
reg295944 := PrimHead(V1641)
reg295945 := PrimTail(V1641)
reg295946 := __e.Call(__defun__shen_4embed_1and, reg295945)
reg295947 := Nil;
reg295948 := PrimCons(reg295946, reg295947)
reg295949 := PrimCons(reg295944, reg295948)
reg295950 := PrimCons(reg295943, reg295949)
__ctx.Return(reg295950)
return
} else {
reg295951 := MakeSymbol("shen.embed-and")
__ctx.TailApply(__defun__shen_4f__error, reg295951)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.embed-and", value: __defun__shen_4embed_1and})

__defun__shen_4err_1condition = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1643 := __args[0]
_ = V1643
reg295953 := True;
reg295954 := MakeSymbol("shen.f_error")
reg295955 := Nil;
reg295956 := PrimCons(V1643, reg295955)
reg295957 := PrimCons(reg295954, reg295956)
reg295958 := Nil;
reg295959 := PrimCons(reg295957, reg295958)
reg295960 := PrimCons(reg295953, reg295959)
__ctx.Return(reg295960)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.err-condition", value: __defun__shen_4err_1condition})

__defun__shen_4sys_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1645 := __args[0]
_ = V1645
reg295961 := MakeString("system function ")
reg295962 := MakeString(": unexpected argument\n")
reg295963 := MakeSymbol("shen.a")
reg295964 := __e.Call(__defun__shen_4app, V1645, reg295962, reg295963)
reg295965 := PrimStringConcat(reg295961, reg295964)
reg295966 := PrimSimpleError(reg295965)
__ctx.Return(reg295966)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sys-error", value: __defun__shen_4sys_1error})

}
