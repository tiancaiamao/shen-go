package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4initialise__arity__table Obj // shen.initialise_arity_table
var __defun__arity Obj // arity
var __defun__systemf Obj // systemf
var __defun__adjoin Obj // adjoin
var __defun__shen_4lambda_1form_1entry Obj // shen.lambda-form-entry
var __defun__shen_4lambda_1form Obj // shen.lambda-form
var __defun__shen_4add_1end Obj // shen.add-end
var __defun__shen_4set_1lambda_1form_1entry Obj // shen.set-lambda-form-entry
var __defun__specialise Obj // specialise
var __defun__unspecialise Obj // unspecialise

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19088 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
__ctx.Return(reg19088)
return
}, 0))
__defun__shen_4initialise__arity__table = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V421 := __args[0]
_ = V421
reg19089 := Nil;
reg19090 := PrimEqual(reg19089, V421)
if reg19090 == True {
reg19091 := Nil;
__ctx.Return(reg19091)
return
} else {
reg19092 := PrimIsPair(V421)
var reg19099 Obj
if reg19092 == True {
reg19093 := PrimTail(V421)
reg19094 := PrimIsPair(reg19093)
var reg19097 Obj
if reg19094 == True {
reg19095 := True;
reg19097 = reg19095
} else {
reg19096 := False;
reg19097 = reg19096
}
reg19099 = reg19097
} else {
reg19098 := False;
reg19099 = reg19098
}
if reg19099 == True {
reg19100 := PrimHead(V421)
reg19101 := MakeSymbol("arity")
reg19102 := PrimTail(V421)
reg19103 := PrimHead(reg19102)
reg19104 := MakeSymbol("*property-vector*")
reg19105 := PrimValue(reg19104)
reg19106 := __e.Call(__defun__put, reg19100, reg19101, reg19103, reg19105)
DecArity := reg19106
_ = DecArity
reg19107 := PrimTail(V421)
reg19108 := PrimTail(reg19107)
__ctx.TailApply(__defun__shen_4initialise__arity__table, reg19108)
return
} else {
reg19110 := MakeSymbol("shen.initialise_arity_table")
__ctx.TailApply(__defun__shen_4f__error, reg19110)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.initialise_arity_table", value: __defun__shen_4initialise__arity__table})

__defun__arity = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V423 := __args[0]
_ = V423
reg19112 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg19113 := MakeSymbol("arity")
reg19114 := MakeSymbol("*property-vector*")
reg19115 := PrimValue(reg19114)
__ctx.TailApply(__defun__get, V423, reg19113, reg19115)
return
}, 0)
reg19117 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg19118 := MakeNumber(-1)
__ctx.Return(reg19118)
return
}, 1)
reg19119 := __e.Try(reg19112).Catch(reg19117)
__ctx.Return(reg19119)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "arity", value: __defun__arity})

__defun__systemf = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V425 := __args[0]
_ = V425
reg19120 := MakeString("shen")
reg19121 := PrimIntern(reg19120)
Shen := reg19121
_ = Shen
reg19122 := MakeSymbol("shen.external-symbols")
reg19123 := MakeSymbol("*property-vector*")
reg19124 := PrimValue(reg19123)
reg19125 := __e.Call(__defun__get, Shen, reg19122, reg19124)
External := reg19125
_ = External
reg19126 := MakeSymbol("shen.external-symbols")
reg19127 := __e.Call(__defun__adjoin, V425, External)
reg19128 := MakeSymbol("*property-vector*")
reg19129 := PrimValue(reg19128)
reg19130 := __e.Call(__defun__put, Shen, reg19126, reg19127, reg19129)
Place := reg19130
_ = Place
__ctx.Return(V425)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "systemf", value: __defun__systemf})

__defun__adjoin = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V428 := __args[0]
_ = V428
V429 := __args[1]
_ = V429
reg19131 := __e.Call(__defun__element_2, V428, V429)
if reg19131 == True {
__ctx.Return(V429)
return
} else {
reg19132 := PrimCons(V428, V429)
__ctx.Return(reg19132)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "adjoin", value: __defun__adjoin})

__defun__shen_4lambda_1form_1entry = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V431 := __args[0]
_ = V431
reg19133 := MakeSymbol("package")
reg19134 := PrimEqual(reg19133, V431)
if reg19134 == True {
reg19135 := Nil;
__ctx.Return(reg19135)
return
} else {
reg19136 := MakeSymbol("receive")
reg19137 := PrimEqual(reg19136, V431)
if reg19137 == True {
reg19138 := Nil;
__ctx.Return(reg19138)
return
} else {
reg19139 := __e.Call(__defun__arity, V431)
ArityF := reg19139
_ = ArityF
reg19140 := MakeNumber(-1)
reg19141 := PrimEqual(ArityF, reg19140)
if reg19141 == True {
reg19142 := Nil;
__ctx.Return(reg19142)
return
} else {
reg19143 := MakeNumber(0)
reg19144 := PrimEqual(ArityF, reg19143)
if reg19144 == True {
reg19145 := Nil;
__ctx.Return(reg19145)
return
} else {
reg19146 := __e.Call(__defun__shen_4lambda_1form, V431, ArityF)
reg19147 := PrimEvalKL(__e, reg19146)
reg19148 := PrimCons(V431, reg19147)
reg19149 := Nil;
reg19150 := PrimCons(reg19148, reg19149)
__ctx.Return(reg19150)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.lambda-form-entry", value: __defun__shen_4lambda_1form_1entry})

__defun__shen_4lambda_1form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V434 := __args[0]
_ = V434
V435 := __args[1]
_ = V435
reg19151 := MakeNumber(0)
reg19152 := PrimEqual(reg19151, V435)
if reg19152 == True {
__ctx.Return(V434)
return
} else {
reg19153 := MakeSymbol("V")
reg19154 := __e.Call(__defun__gensym, reg19153)
X := reg19154
_ = X
reg19155 := MakeSymbol("lambda")
reg19156 := __e.Call(__defun__shen_4add_1end, V434, X)
reg19157 := MakeNumber(1)
reg19158 := PrimNumberSubtract(V435, reg19157)
reg19159 := __e.Call(__defun__shen_4lambda_1form, reg19156, reg19158)
reg19160 := Nil;
reg19161 := PrimCons(reg19159, reg19160)
reg19162 := PrimCons(X, reg19161)
reg19163 := PrimCons(reg19155, reg19162)
__ctx.Return(reg19163)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.lambda-form", value: __defun__shen_4lambda_1form})

__defun__shen_4add_1end = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V438 := __args[0]
_ = V438
V439 := __args[1]
_ = V439
reg19164 := PrimIsPair(V438)
if reg19164 == True {
reg19165 := Nil;
reg19166 := PrimCons(V439, reg19165)
__ctx.TailApply(__defun__append, V438, reg19166)
return
} else {
reg19168 := Nil;
reg19169 := PrimCons(V439, reg19168)
reg19170 := PrimCons(V438, reg19169)
__ctx.Return(reg19170)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.add-end", value: __defun__shen_4add_1end})

__defun__shen_4set_1lambda_1form_1entry = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V441 := __args[0]
_ = V441
reg19171 := PrimIsPair(V441)
if reg19171 == True {
reg19172 := PrimHead(V441)
reg19173 := MakeSymbol("shen.lambda-form")
reg19174 := PrimTail(V441)
reg19175 := MakeSymbol("*property-vector*")
reg19176 := PrimValue(reg19175)
__ctx.TailApply(__defun__put, reg19172, reg19173, reg19174, reg19176)
return
} else {
reg19178 := MakeSymbol("shen.set-lambda-form-entry")
__ctx.TailApply(__defun__shen_4f__error, reg19178)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.set-lambda-form-entry", value: __defun__shen_4set_1lambda_1form_1entry})

__defun__specialise = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V443 := __args[0]
_ = V443
reg19180 := MakeSymbol("shen.*special*")
reg19181 := MakeSymbol("shen.*special*")
reg19182 := PrimValue(reg19181)
reg19183 := PrimCons(V443, reg19182)
reg19184 := PrimSet(reg19180, reg19183)
_ = reg19184
__ctx.Return(V443)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "specialise", value: __defun__specialise})

__defun__unspecialise = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V445 := __args[0]
_ = V445
reg19185 := MakeSymbol("shen.*special*")
reg19186 := MakeSymbol("shen.*special*")
reg19187 := PrimValue(reg19186)
reg19188 := __e.Call(__defun__remove, V445, reg19187)
reg19189 := PrimSet(reg19185, reg19188)
_ = reg19189
__ctx.Return(V445)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "unspecialise", value: __defun__unspecialise})

}
