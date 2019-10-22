package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__thaw Obj // thaw
var __defun__eval Obj // eval
var __defun__shen_4eval_1without_1macros Obj // shen.eval-without-macros
var __defun__shen_4proc_1input_7 Obj // shen.proc-input+
var __defun__shen_4elim_1def Obj // shen.elim-def
var __defun__shen_4add_1macro Obj // shen.add-macro
var __defun__shen_4packaged_2 Obj // shen.packaged?
var __defun__external Obj // external
var __defun__internal Obj // internal
var __defun__shen_4package_1contents Obj // shen.package-contents
var __defun__shen_4walk Obj // shen.walk
var __defun__compile Obj // compile
var __defun__fail_1if Obj // fail-if
var __defun___8s Obj // @s
var __defun__tc_2 Obj // tc?
var __defun__ps Obj // ps
var __defun__stinput Obj // stinput
var __defun__vector Obj // vector
var __defun__shen_4fillvector Obj // shen.fillvector
var __defun__vector_2 Obj // vector?
var __defun__vector_1_6 Obj // vector->
var __defun___5_1vector Obj // <-vector
var __defun__shen_4posint_2 Obj // shen.posint?
var __defun__limit Obj // limit
var __defun__symbol_2 Obj // symbol?
var __defun__shen_4analyse_1symbol_2 Obj // shen.analyse-symbol?
var __defun__shen_4alpha_2 Obj // shen.alpha?
var __defun__shen_4alphanums_2 Obj // shen.alphanums?
var __defun__shen_4alphanum_2 Obj // shen.alphanum?
var __defun__shen_4digit_2 Obj // shen.digit?
var __defun__variable_2 Obj // variable?
var __defun__shen_4analyse_1variable_2 Obj // shen.analyse-variable?
var __defun__shen_4uppercase_2 Obj // shen.uppercase?
var __defun__gensym Obj // gensym
var __defun__concat Obj // concat
var __defun___8p Obj // @p
var __defun__fst Obj // fst
var __defun__snd Obj // snd
var __defun__tuple_2 Obj // tuple?
var __defun__append Obj // append
var __defun___8v Obj // @v
var __defun__shen_4_8v_1help Obj // shen.@v-help
var __defun__shen_4copyfromvector Obj // shen.copyfromvector
var __defun__hdv Obj // hdv
var __defun__tlv Obj // tlv
var __defun__shen_4tlv_1help Obj // shen.tlv-help
var __defun__assoc Obj // assoc
var __defun__shen_4assoc_1set Obj // shen.assoc-set
var __defun__shen_4assoc_1rm Obj // shen.assoc-rm
var __defun__boolean_2 Obj // boolean?
var __defun__nl Obj // nl
var __defun__difference Obj // difference
var __defun__do Obj // do
var __defun__element_2 Obj // element?
var __defun__empty_2 Obj // empty?
var __defun__fix Obj // fix
var __defun__shen_4fix_1help Obj // shen.fix-help
var __defun__put Obj // put
var __defun__unput Obj // unput
var __defun__get Obj // get
var __defun__hash Obj // hash
var __defun__shen_4mod Obj // shen.mod
var __defun__shen_4multiples Obj // shen.multiples
var __defun__shen_4modh Obj // shen.modh
var __defun__sum Obj // sum
var __defun__head Obj // head
var __defun__tail Obj // tail
var __defun__hdstr Obj // hdstr
var __defun__intersection Obj // intersection
var __defun__reverse Obj // reverse
var __defun__shen_4reverse__help Obj // shen.reverse_help
var __defun__union Obj // union
var __defun__y_1or_1n_2 Obj // y-or-n?
var __defun__not Obj // not
var __defun__subst Obj // subst
var __defun__explode Obj // explode
var __defun__shen_4explode_1h Obj // shen.explode-h
var __defun__cd Obj // cd
var __defun__shen_4for_1each Obj // shen.for-each
var __defun__map Obj // map
var __defun__shen_4map_1h Obj // shen.map-h
var __defun__length Obj // length
var __defun__shen_4length_1h Obj // shen.length-h
var __defun__occurrences Obj // occurrences
var __defun__nth Obj // nth
var __defun__integer_2 Obj // integer?
var __defun__shen_4abs Obj // shen.abs
var __defun__shen_4magless Obj // shen.magless
var __defun__shen_4integer_1test_2 Obj // shen.integer-test?
var __defun__mapcan Obj // mapcan
var __defun___a_a Obj // ==
var __defun__abort Obj // abort
var __defun__bound_2 Obj // bound?
var __defun__shen_4string_1_6bytes Obj // shen.string->bytes
var __defun__maxinferences Obj // maxinferences
var __defun__inferences Obj // inferences
var __defun__protect Obj // protect
var __defun__stoutput Obj // stoutput
var __defun__sterror Obj // sterror
var __defun__string_1_6symbol Obj // string->symbol
var __defun__optimise Obj // optimise
var __defun__os Obj // os
var __defun__language Obj // language
var __defun__version Obj // version
var __defun__port Obj // port
var __defun__porters Obj // porters
var __defun__implementation Obj // implementation
var __defun__release Obj // release
var __defun__package_2 Obj // package?
var __defun__function Obj // function
var __defun__shen_4lookup_1func Obj // shen.lookup-func

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg295967 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg295967)
return
}, 0))
__defun__thaw = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2827 := __args[0]
_ = V2827
__ctx.TailApply(V2827)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "thaw", value: __defun__thaw})

__defun__eval = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2829 := __args[0]
_ = V2829
reg295969 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__macroexpand, Y)
return
}, 1)
reg295971 := __e.Call(__defun__shen_4walk, reg295969, V2829)
Macroexpand := reg295971
_ = Macroexpand
reg295972 := __e.Call(__defun__shen_4packaged_2, Macroexpand)
if reg295972 == True {
reg295973 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4eval_1without_1macros, Z)
return
}, 1)
reg295975 := __e.Call(__defun__shen_4package_1contents, Macroexpand)
__ctx.TailApply(__defun__map, reg295973, reg295975)
return
} else {
__ctx.TailApply(__defun__shen_4eval_1without_1macros, Macroexpand)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "eval", value: __defun__eval})

__defun__shen_4eval_1without_1macros = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2831 := __args[0]
_ = V2831
reg295978 := __e.Call(__defun__shen_4proc_1input_7, V2831)
reg295979 := __e.Call(__defun__shen_4elim_1def, reg295978)
reg295980 := PrimEvalKL(__e, reg295979)
__ctx.Return(reg295980)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.eval-without-macros", value: __defun__shen_4eval_1without_1macros})

__defun__shen_4proc_1input_7 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2833 := __args[0]
_ = V2833
reg295981 := PrimIsPair(V2833)
var reg296014 Obj
if reg295981 == True {
reg295982 := MakeSymbol("input+")
reg295983 := PrimHead(V2833)
reg295984 := PrimEqual(reg295982, reg295983)
var reg296009 Obj
if reg295984 == True {
reg295985 := PrimTail(V2833)
reg295986 := PrimIsPair(reg295985)
var reg296004 Obj
if reg295986 == True {
reg295987 := PrimTail(V2833)
reg295988 := PrimTail(reg295987)
reg295989 := PrimIsPair(reg295988)
var reg295999 Obj
if reg295989 == True {
reg295990 := Nil;
reg295991 := PrimTail(V2833)
reg295992 := PrimTail(reg295991)
reg295993 := PrimTail(reg295992)
reg295994 := PrimEqual(reg295990, reg295993)
var reg295997 Obj
if reg295994 == True {
reg295995 := True;
reg295997 = reg295995
} else {
reg295996 := False;
reg295997 = reg295996
}
reg295999 = reg295997
} else {
reg295998 := False;
reg295999 = reg295998
}
var reg296002 Obj
if reg295999 == True {
reg296000 := True;
reg296002 = reg296000
} else {
reg296001 := False;
reg296002 = reg296001
}
reg296004 = reg296002
} else {
reg296003 := False;
reg296004 = reg296003
}
var reg296007 Obj
if reg296004 == True {
reg296005 := True;
reg296007 = reg296005
} else {
reg296006 := False;
reg296007 = reg296006
}
reg296009 = reg296007
} else {
reg296008 := False;
reg296009 = reg296008
}
var reg296012 Obj
if reg296009 == True {
reg296010 := True;
reg296012 = reg296010
} else {
reg296011 := False;
reg296012 = reg296011
}
reg296014 = reg296012
} else {
reg296013 := False;
reg296014 = reg296013
}
if reg296014 == True {
reg296015 := MakeSymbol("input+")
reg296016 := PrimTail(V2833)
reg296017 := PrimHead(reg296016)
reg296018 := __e.Call(__defun__shen_4rcons__form, reg296017)
reg296019 := PrimTail(V2833)
reg296020 := PrimTail(reg296019)
reg296021 := PrimCons(reg296018, reg296020)
reg296022 := PrimCons(reg296015, reg296021)
__ctx.Return(reg296022)
return
} else {
reg296023 := PrimIsPair(V2833)
var reg296056 Obj
if reg296023 == True {
reg296024 := MakeSymbol("shen.read+")
reg296025 := PrimHead(V2833)
reg296026 := PrimEqual(reg296024, reg296025)
var reg296051 Obj
if reg296026 == True {
reg296027 := PrimTail(V2833)
reg296028 := PrimIsPair(reg296027)
var reg296046 Obj
if reg296028 == True {
reg296029 := PrimTail(V2833)
reg296030 := PrimTail(reg296029)
reg296031 := PrimIsPair(reg296030)
var reg296041 Obj
if reg296031 == True {
reg296032 := Nil;
reg296033 := PrimTail(V2833)
reg296034 := PrimTail(reg296033)
reg296035 := PrimTail(reg296034)
reg296036 := PrimEqual(reg296032, reg296035)
var reg296039 Obj
if reg296036 == True {
reg296037 := True;
reg296039 = reg296037
} else {
reg296038 := False;
reg296039 = reg296038
}
reg296041 = reg296039
} else {
reg296040 := False;
reg296041 = reg296040
}
var reg296044 Obj
if reg296041 == True {
reg296042 := True;
reg296044 = reg296042
} else {
reg296043 := False;
reg296044 = reg296043
}
reg296046 = reg296044
} else {
reg296045 := False;
reg296046 = reg296045
}
var reg296049 Obj
if reg296046 == True {
reg296047 := True;
reg296049 = reg296047
} else {
reg296048 := False;
reg296049 = reg296048
}
reg296051 = reg296049
} else {
reg296050 := False;
reg296051 = reg296050
}
var reg296054 Obj
if reg296051 == True {
reg296052 := True;
reg296054 = reg296052
} else {
reg296053 := False;
reg296054 = reg296053
}
reg296056 = reg296054
} else {
reg296055 := False;
reg296056 = reg296055
}
if reg296056 == True {
reg296057 := MakeSymbol("shen.read+")
reg296058 := PrimTail(V2833)
reg296059 := PrimHead(reg296058)
reg296060 := __e.Call(__defun__shen_4rcons__form, reg296059)
reg296061 := PrimTail(V2833)
reg296062 := PrimTail(reg296061)
reg296063 := PrimCons(reg296060, reg296062)
reg296064 := PrimCons(reg296057, reg296063)
__ctx.Return(reg296064)
return
} else {
reg296065 := PrimIsPair(V2833)
if reg296065 == True {
reg296066 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4proc_1input_7, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg296066, V2833)
return
} else {
__ctx.Return(V2833)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.proc-input+", value: __defun__shen_4proc_1input_7})

__defun__shen_4elim_1def = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2835 := __args[0]
_ = V2835
reg296069 := PrimIsPair(V2835)
var reg296084 Obj
if reg296069 == True {
reg296070 := MakeSymbol("define")
reg296071 := PrimHead(V2835)
reg296072 := PrimEqual(reg296070, reg296071)
var reg296079 Obj
if reg296072 == True {
reg296073 := PrimTail(V2835)
reg296074 := PrimIsPair(reg296073)
var reg296077 Obj
if reg296074 == True {
reg296075 := True;
reg296077 = reg296075
} else {
reg296076 := False;
reg296077 = reg296076
}
reg296079 = reg296077
} else {
reg296078 := False;
reg296079 = reg296078
}
var reg296082 Obj
if reg296079 == True {
reg296080 := True;
reg296082 = reg296080
} else {
reg296081 := False;
reg296082 = reg296081
}
reg296084 = reg296082
} else {
reg296083 := False;
reg296084 = reg296083
}
if reg296084 == True {
reg296085 := PrimTail(V2835)
reg296086 := PrimHead(reg296085)
reg296087 := PrimTail(V2835)
reg296088 := PrimTail(reg296087)
__ctx.TailApply(__defun__shen_4shen_1_6kl, reg296086, reg296088)
return
} else {
reg296090 := PrimIsPair(V2835)
var reg296105 Obj
if reg296090 == True {
reg296091 := MakeSymbol("defmacro")
reg296092 := PrimHead(V2835)
reg296093 := PrimEqual(reg296091, reg296092)
var reg296100 Obj
if reg296093 == True {
reg296094 := PrimTail(V2835)
reg296095 := PrimIsPair(reg296094)
var reg296098 Obj
if reg296095 == True {
reg296096 := True;
reg296098 = reg296096
} else {
reg296097 := False;
reg296098 = reg296097
}
reg296100 = reg296098
} else {
reg296099 := False;
reg296100 = reg296099
}
var reg296103 Obj
if reg296100 == True {
reg296101 := True;
reg296103 = reg296101
} else {
reg296102 := False;
reg296103 = reg296102
}
reg296105 = reg296103
} else {
reg296104 := False;
reg296105 = reg296104
}
if reg296105 == True {
reg296106 := MakeSymbol("X")
reg296107 := MakeSymbol("->")
reg296108 := MakeSymbol("X")
reg296109 := Nil;
reg296110 := PrimCons(reg296108, reg296109)
reg296111 := PrimCons(reg296107, reg296110)
reg296112 := PrimCons(reg296106, reg296111)
Default := reg296112
_ = Default
reg296113 := MakeSymbol("define")
reg296114 := PrimTail(V2835)
reg296115 := PrimHead(reg296114)
reg296116 := PrimTail(V2835)
reg296117 := PrimTail(reg296116)
reg296118 := __e.Call(__defun__append, reg296117, Default)
reg296119 := PrimCons(reg296115, reg296118)
reg296120 := PrimCons(reg296113, reg296119)
reg296121 := __e.Call(__defun__shen_4elim_1def, reg296120)
Def := reg296121
_ = Def
reg296122 := PrimTail(V2835)
reg296123 := PrimHead(reg296122)
reg296124 := __e.Call(__defun__shen_4add_1macro, reg296123)
MacroAdd := reg296124
_ = MacroAdd
__ctx.Return(Def)
return
} else {
reg296125 := PrimIsPair(V2835)
var reg296140 Obj
if reg296125 == True {
reg296126 := MakeSymbol("defcc")
reg296127 := PrimHead(V2835)
reg296128 := PrimEqual(reg296126, reg296127)
var reg296135 Obj
if reg296128 == True {
reg296129 := PrimTail(V2835)
reg296130 := PrimIsPair(reg296129)
var reg296133 Obj
if reg296130 == True {
reg296131 := True;
reg296133 = reg296131
} else {
reg296132 := False;
reg296133 = reg296132
}
reg296135 = reg296133
} else {
reg296134 := False;
reg296135 = reg296134
}
var reg296138 Obj
if reg296135 == True {
reg296136 := True;
reg296138 = reg296136
} else {
reg296137 := False;
reg296138 = reg296137
}
reg296140 = reg296138
} else {
reg296139 := False;
reg296140 = reg296139
}
if reg296140 == True {
reg296141 := __e.Call(__defun__shen_4yacc, V2835)
__ctx.TailApply(__defun__shen_4elim_1def, reg296141)
return
} else {
reg296143 := PrimIsPair(V2835)
if reg296143 == True {
reg296144 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4elim_1def, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg296144, V2835)
return
} else {
__ctx.Return(V2835)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.elim-def", value: __defun__shen_4elim_1def})

__defun__shen_4add_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2837 := __args[0]
_ = V2837
reg296147 := MakeSymbol("shen.*macroreg*")
reg296148 := PrimValue(reg296147)
MacroReg := reg296148
_ = MacroReg
reg296149 := MakeSymbol("shen.*macroreg*")
reg296150 := MakeSymbol("shen.*macroreg*")
reg296151 := PrimValue(reg296150)
reg296152 := __e.Call(__defun__adjoin, V2837, reg296151)
reg296153 := PrimSet(reg296149, reg296152)
NewMacroReg := reg296153
_ = NewMacroReg
reg296154 := PrimEqual(MacroReg, NewMacroReg)
if reg296154 == True {
reg296155 := MakeSymbol("shen.skip")
__ctx.Return(reg296155)
return
} else {
reg296156 := MakeSymbol("*macros*")
reg296157 := __e.Call(__defun__function, V2837)
reg296158 := MakeSymbol("*macros*")
reg296159 := PrimValue(reg296158)
reg296160 := PrimCons(reg296157, reg296159)
reg296161 := PrimSet(reg296156, reg296160)
__ctx.Return(reg296161)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.add-macro", value: __defun__shen_4add_1macro})

__defun__shen_4packaged_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2845 := __args[0]
_ = V2845
reg296162 := PrimIsPair(V2845)
var reg296185 Obj
if reg296162 == True {
reg296163 := MakeSymbol("package")
reg296164 := PrimHead(V2845)
reg296165 := PrimEqual(reg296163, reg296164)
var reg296180 Obj
if reg296165 == True {
reg296166 := PrimTail(V2845)
reg296167 := PrimIsPair(reg296166)
var reg296175 Obj
if reg296167 == True {
reg296168 := PrimTail(V2845)
reg296169 := PrimTail(reg296168)
reg296170 := PrimIsPair(reg296169)
var reg296173 Obj
if reg296170 == True {
reg296171 := True;
reg296173 = reg296171
} else {
reg296172 := False;
reg296173 = reg296172
}
reg296175 = reg296173
} else {
reg296174 := False;
reg296175 = reg296174
}
var reg296178 Obj
if reg296175 == True {
reg296176 := True;
reg296178 = reg296176
} else {
reg296177 := False;
reg296178 = reg296177
}
reg296180 = reg296178
} else {
reg296179 := False;
reg296180 = reg296179
}
var reg296183 Obj
if reg296180 == True {
reg296181 := True;
reg296183 = reg296181
} else {
reg296182 := False;
reg296183 = reg296182
}
reg296185 = reg296183
} else {
reg296184 := False;
reg296185 = reg296184
}
if reg296185 == True {
reg296186 := True;
__ctx.Return(reg296186)
return
} else {
reg296187 := False;
__ctx.Return(reg296187)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.packaged?", value: __defun__shen_4packaged_2})

__defun__external = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2847 := __args[0]
_ = V2847
reg296188 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296189 := MakeSymbol("shen.external-symbols")
reg296190 := MakeSymbol("*property-vector*")
reg296191 := PrimValue(reg296190)
__ctx.TailApply(__defun__get, V2847, reg296189, reg296191)
return
}, 0)
reg296193 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296194 := MakeString("package ")
reg296195 := MakeString(" has not been used.\n")
reg296196 := MakeSymbol("shen.a")
reg296197 := __e.Call(__defun__shen_4app, V2847, reg296195, reg296196)
reg296198 := PrimStringConcat(reg296194, reg296197)
reg296199 := PrimSimpleError(reg296198)
__ctx.Return(reg296199)
return
}, 1)
reg296200 := __e.Try(reg296188).Catch(reg296193)
__ctx.Return(reg296200)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "external", value: __defun__external})

__defun__internal = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2849 := __args[0]
_ = V2849
reg296201 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296202 := MakeSymbol("shen.internal-symbols")
reg296203 := MakeSymbol("*property-vector*")
reg296204 := PrimValue(reg296203)
__ctx.TailApply(__defun__get, V2849, reg296202, reg296204)
return
}, 0)
reg296206 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296207 := MakeString("package ")
reg296208 := MakeString(" has not been used.\n")
reg296209 := MakeSymbol("shen.a")
reg296210 := __e.Call(__defun__shen_4app, V2849, reg296208, reg296209)
reg296211 := PrimStringConcat(reg296207, reg296210)
reg296212 := PrimSimpleError(reg296211)
__ctx.Return(reg296212)
return
}, 1)
reg296213 := __e.Try(reg296201).Catch(reg296206)
__ctx.Return(reg296213)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "internal", value: __defun__internal})

__defun__shen_4package_1contents = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2853 := __args[0]
_ = V2853
reg296214 := PrimIsPair(V2853)
var reg296246 Obj
if reg296214 == True {
reg296215 := MakeSymbol("package")
reg296216 := PrimHead(V2853)
reg296217 := PrimEqual(reg296215, reg296216)
var reg296241 Obj
if reg296217 == True {
reg296218 := PrimTail(V2853)
reg296219 := PrimIsPair(reg296218)
var reg296236 Obj
if reg296219 == True {
reg296220 := MakeSymbol("null")
reg296221 := PrimTail(V2853)
reg296222 := PrimHead(reg296221)
reg296223 := PrimEqual(reg296220, reg296222)
var reg296231 Obj
if reg296223 == True {
reg296224 := PrimTail(V2853)
reg296225 := PrimTail(reg296224)
reg296226 := PrimIsPair(reg296225)
var reg296229 Obj
if reg296226 == True {
reg296227 := True;
reg296229 = reg296227
} else {
reg296228 := False;
reg296229 = reg296228
}
reg296231 = reg296229
} else {
reg296230 := False;
reg296231 = reg296230
}
var reg296234 Obj
if reg296231 == True {
reg296232 := True;
reg296234 = reg296232
} else {
reg296233 := False;
reg296234 = reg296233
}
reg296236 = reg296234
} else {
reg296235 := False;
reg296236 = reg296235
}
var reg296239 Obj
if reg296236 == True {
reg296237 := True;
reg296239 = reg296237
} else {
reg296238 := False;
reg296239 = reg296238
}
reg296241 = reg296239
} else {
reg296240 := False;
reg296241 = reg296240
}
var reg296244 Obj
if reg296241 == True {
reg296242 := True;
reg296244 = reg296242
} else {
reg296243 := False;
reg296244 = reg296243
}
reg296246 = reg296244
} else {
reg296245 := False;
reg296246 = reg296245
}
if reg296246 == True {
reg296247 := PrimTail(V2853)
reg296248 := PrimTail(reg296247)
reg296249 := PrimTail(reg296248)
__ctx.Return(reg296249)
return
} else {
reg296250 := PrimIsPair(V2853)
var reg296273 Obj
if reg296250 == True {
reg296251 := MakeSymbol("package")
reg296252 := PrimHead(V2853)
reg296253 := PrimEqual(reg296251, reg296252)
var reg296268 Obj
if reg296253 == True {
reg296254 := PrimTail(V2853)
reg296255 := PrimIsPair(reg296254)
var reg296263 Obj
if reg296255 == True {
reg296256 := PrimTail(V2853)
reg296257 := PrimTail(reg296256)
reg296258 := PrimIsPair(reg296257)
var reg296261 Obj
if reg296258 == True {
reg296259 := True;
reg296261 = reg296259
} else {
reg296260 := False;
reg296261 = reg296260
}
reg296263 = reg296261
} else {
reg296262 := False;
reg296263 = reg296262
}
var reg296266 Obj
if reg296263 == True {
reg296264 := True;
reg296266 = reg296264
} else {
reg296265 := False;
reg296266 = reg296265
}
reg296268 = reg296266
} else {
reg296267 := False;
reg296268 = reg296267
}
var reg296271 Obj
if reg296268 == True {
reg296269 := True;
reg296271 = reg296269
} else {
reg296270 := False;
reg296271 = reg296270
}
reg296273 = reg296271
} else {
reg296272 := False;
reg296273 = reg296272
}
if reg296273 == True {
reg296274 := PrimTail(V2853)
reg296275 := PrimHead(reg296274)
reg296276 := PrimStr(reg296275)
reg296277 := MakeString(".")
reg296278 := PrimStringConcat(reg296276, reg296277)
reg296279 := PrimIntern(reg296278)
PackageNameDot := reg296279
_ = PackageNameDot
reg296280 := __e.Call(__defun__explode, PackageNameDot)
ExpPackageNameDot := reg296280
_ = ExpPackageNameDot
reg296281 := PrimTail(V2853)
reg296282 := PrimHead(reg296281)
reg296283 := PrimTail(V2853)
reg296284 := PrimTail(reg296283)
reg296285 := PrimHead(reg296284)
reg296286 := PrimTail(V2853)
reg296287 := PrimTail(reg296286)
reg296288 := PrimTail(reg296287)
__ctx.TailApply(__defun__shen_4packageh, reg296282, reg296285, reg296288, ExpPackageNameDot)
return
} else {
reg296290 := MakeSymbol("shen.package-contents")
__ctx.TailApply(__defun__shen_4f__error, reg296290)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.package-contents", value: __defun__shen_4package_1contents})

__defun__shen_4walk = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2856 := __args[0]
_ = V2856
V2857 := __args[1]
_ = V2857
reg296292 := PrimIsPair(V2857)
if reg296292 == True {
reg296293 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4walk, V2856, Z)
return
}, 1)
reg296295 := __e.Call(__defun__map, reg296293, V2857)
__ctx.TailApply(V2856, reg296295)
return
} else {
__ctx.TailApply(V2856, V2857)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.walk", value: __defun__shen_4walk})

__defun__compile = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2861 := __args[0]
_ = V2861
V2862 := __args[1]
_ = V2862
V2863 := __args[2]
_ = V2863
reg296298 := Nil;
reg296299 := Nil;
reg296300 := PrimCons(reg296298, reg296299)
reg296301 := PrimCons(V2862, reg296300)
reg296302 := __e.Call(V2861, reg296301)
O := reg296302
_ = O
reg296303 := __e.Call(__defun__fail)
reg296304 := PrimEqual(reg296303, O)
var reg296312 Obj
if reg296304 == True {
reg296305 := True;
reg296312 = reg296305
} else {
reg296306 := PrimHead(O)
reg296307 := __e.Call(__defun__empty_2, reg296306)
reg296308 := PrimNot(reg296307)
var reg296311 Obj
if reg296308 == True {
reg296309 := True;
reg296311 = reg296309
} else {
reg296310 := False;
reg296311 = reg296310
}
reg296312 = reg296311
}
if reg296312 == True {
__ctx.TailApply(V2863, O)
return
} else {
__ctx.TailApply(__defun__shen_4hdtl, O)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "compile", value: __defun__compile})

__defun__fail_1if = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2866 := __args[0]
_ = V2866
V2867 := __args[1]
_ = V2867
reg296315 := __e.Call(V2866, V2867)
if reg296315 == True {
__ctx.TailApply(__defun__fail)
return
} else {
__ctx.Return(V2867)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "fail-if", value: __defun__fail_1if})

__defun___8s = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2870 := __args[0]
_ = V2870
V2871 := __args[1]
_ = V2871
reg296317 := PrimStringConcat(V2870, V2871)
__ctx.Return(reg296317)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "@s", value: __defun___8s})

__defun__tc_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296318 := MakeSymbol("shen.*tc*")
reg296319 := PrimValue(reg296318)
__ctx.Return(reg296319)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "tc?", value: __defun__tc_2})

__defun__ps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2873 := __args[0]
_ = V2873
reg296320 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296321 := MakeSymbol("shen.source")
reg296322 := MakeSymbol("*property-vector*")
reg296323 := PrimValue(reg296322)
__ctx.TailApply(__defun__get, V2873, reg296321, reg296323)
return
}, 0)
reg296325 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296326 := MakeString(" not found.\n")
reg296327 := MakeSymbol("shen.a")
reg296328 := __e.Call(__defun__shen_4app, V2873, reg296326, reg296327)
reg296329 := PrimSimpleError(reg296328)
__ctx.Return(reg296329)
return
}, 1)
reg296330 := __e.Try(reg296320).Catch(reg296325)
__ctx.Return(reg296330)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "ps", value: __defun__ps})

__defun__stinput = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296331 := MakeSymbol("*stinput*")
reg296332 := PrimValue(reg296331)
__ctx.Return(reg296332)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "stinput", value: __defun__stinput})

__defun__vector = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2875 := __args[0]
_ = V2875
reg296333 := MakeNumber(1)
reg296334 := PrimNumberAdd(V2875, reg296333)
reg296335 := PrimAbsvector(reg296334)
Vector := reg296335
_ = Vector
reg296336 := MakeNumber(0)
reg296337 := PrimVectorSet(Vector, reg296336, V2875)
ZeroStamp := reg296337
_ = ZeroStamp
reg296338 := MakeNumber(0)
reg296339 := PrimEqual(V2875, reg296338)
var reg296343 Obj
if reg296339 == True {
reg296343 = ZeroStamp
} else {
reg296340 := MakeNumber(1)
reg296341 := __e.Call(__defun__fail)
reg296342 := __e.Call(__defun__shen_4fillvector, ZeroStamp, reg296340, V2875, reg296341)
reg296343 = reg296342
}
Standard := reg296343
_ = Standard
__ctx.Return(Standard)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "vector", value: __defun__vector})

__defun__shen_4fillvector = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2881 := __args[0]
_ = V2881
V2882 := __args[1]
_ = V2882
V2883 := __args[2]
_ = V2883
V2884 := __args[3]
_ = V2884
reg296344 := PrimEqual(V2883, V2882)
if reg296344 == True {
reg296345 := PrimVectorSet(V2881, V2883, V2884)
__ctx.Return(reg296345)
return
} else {
reg296346 := PrimVectorSet(V2881, V2882, V2884)
reg296347 := MakeNumber(1)
reg296348 := PrimNumberAdd(reg296347, V2882)
__ctx.TailApply(__defun__shen_4fillvector, reg296346, reg296348, V2883, V2884)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.fillvector", value: __defun__shen_4fillvector})

__defun__vector_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2886 := __args[0]
_ = V2886
reg296350 := PrimIsVector(V2886)
if reg296350 == True {
reg296351 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296352 := MakeNumber(0)
reg296353 := PrimVectorGet(V2886, reg296352)
__ctx.Return(reg296353)
return
}, 0)
reg296354 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296355 := MakeNumber(-1)
__ctx.Return(reg296355)
return
}, 1)
reg296356 := __e.Try(reg296351).Catch(reg296354)
X := reg296356
_ = X
reg296357 := PrimIsNumber(X)
var reg296364 Obj
if reg296357 == True {
reg296358 := MakeNumber(0)
reg296359 := PrimGreatEqual(X, reg296358)
var reg296362 Obj
if reg296359 == True {
reg296360 := True;
reg296362 = reg296360
} else {
reg296361 := False;
reg296362 = reg296361
}
reg296364 = reg296362
} else {
reg296363 := False;
reg296364 = reg296363
}
if reg296364 == True {
reg296365 := True;
__ctx.Return(reg296365)
return
} else {
reg296366 := False;
__ctx.Return(reg296366)
return
}
} else {
reg296367 := False;
__ctx.Return(reg296367)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "vector?", value: __defun__vector_2})

__defun__vector_1_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2890 := __args[0]
_ = V2890
V2891 := __args[1]
_ = V2891
V2892 := __args[2]
_ = V2892
reg296368 := MakeNumber(0)
reg296369 := PrimEqual(V2891, reg296368)
if reg296369 == True {
reg296370 := MakeString("cannot access 0th element of a vector\n")
reg296371 := PrimSimpleError(reg296370)
__ctx.Return(reg296371)
return
} else {
reg296372 := PrimVectorSet(V2890, V2891, V2892)
__ctx.Return(reg296372)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "vector->", value: __defun__vector_1_6})

__defun___5_1vector = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2895 := __args[0]
_ = V2895
V2896 := __args[1]
_ = V2896
reg296373 := MakeNumber(0)
reg296374 := PrimEqual(V2896, reg296373)
if reg296374 == True {
reg296375 := MakeString("cannot access 0th element of a vector\n")
reg296376 := PrimSimpleError(reg296375)
__ctx.Return(reg296376)
return
} else {
reg296377 := PrimVectorGet(V2895, V2896)
VectorElement := reg296377
_ = VectorElement
reg296378 := __e.Call(__defun__fail)
reg296379 := PrimEqual(VectorElement, reg296378)
if reg296379 == True {
reg296380 := MakeString("vector element not found\n")
reg296381 := PrimSimpleError(reg296380)
__ctx.Return(reg296381)
return
} else {
__ctx.Return(VectorElement)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "<-vector", value: __defun___5_1vector})

__defun__shen_4posint_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2898 := __args[0]
_ = V2898
reg296382 := PrimIsInteger(V2898)
if reg296382 == True {
reg296383 := MakeNumber(0)
reg296384 := PrimGreatEqual(V2898, reg296383)
if reg296384 == True {
reg296385 := True;
__ctx.Return(reg296385)
return
} else {
reg296386 := False;
__ctx.Return(reg296386)
return
}
} else {
reg296387 := False;
__ctx.Return(reg296387)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.posint?", value: __defun__shen_4posint_2})

__defun__limit = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2900 := __args[0]
_ = V2900
reg296388 := MakeNumber(0)
reg296389 := PrimVectorGet(V2900, reg296388)
__ctx.Return(reg296389)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "limit", value: __defun__limit})

__defun__symbol_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2902 := __args[0]
_ = V2902
reg296390 := __e.Call(__defun__boolean_2, V2902)
var reg296402 Obj
if reg296390 == True {
reg296391 := True;
reg296402 = reg296391
} else {
reg296392 := PrimIsNumber(V2902)
var reg296398 Obj
if reg296392 == True {
reg296393 := True;
reg296398 = reg296393
} else {
reg296394 := PrimIsString(V2902)
var reg296397 Obj
if reg296394 == True {
reg296395 := True;
reg296397 = reg296395
} else {
reg296396 := False;
reg296397 = reg296396
}
reg296398 = reg296397
}
var reg296401 Obj
if reg296398 == True {
reg296399 := True;
reg296401 = reg296399
} else {
reg296400 := False;
reg296401 = reg296400
}
reg296402 = reg296401
}
if reg296402 == True {
reg296403 := False;
__ctx.Return(reg296403)
return
} else {
reg296404 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296405 := PrimStr(V2902)
String := reg296405
_ = String
__ctx.TailApply(__defun__shen_4analyse_1symbol_2, String)
return
}, 0)
reg296407 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296408 := False;
__ctx.Return(reg296408)
return
}, 1)
reg296409 := __e.Try(reg296404).Catch(reg296407)
__ctx.Return(reg296409)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "symbol?", value: __defun__symbol_2})

__defun__shen_4analyse_1symbol_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2904 := __args[0]
_ = V2904
reg296410 := MakeString("")
reg296411 := PrimEqual(reg296410, V2904)
if reg296411 == True {
reg296412 := False;
__ctx.Return(reg296412)
return
} else {
reg296413 := __e.Call(__defun__shen_4_7string_2, V2904)
if reg296413 == True {
reg296414 := MakeNumber(0)
reg296415 := PrimPos(V2904, reg296414)
reg296416 := __e.Call(__defun__shen_4alpha_2, reg296415)
if reg296416 == True {
reg296417 := PrimTailString(V2904)
reg296418 := __e.Call(__defun__shen_4alphanums_2, reg296417)
if reg296418 == True {
reg296419 := True;
__ctx.Return(reg296419)
return
} else {
reg296420 := False;
__ctx.Return(reg296420)
return
}
} else {
reg296421 := False;
__ctx.Return(reg296421)
return
}
} else {
reg296422 := MakeSymbol("shen.analyse-symbol?")
__ctx.TailApply(__defun__shen_4f__error, reg296422)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.analyse-symbol?", value: __defun__shen_4analyse_1symbol_2})

__defun__shen_4alpha_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2906 := __args[0]
_ = V2906
reg296424 := MakeString("A")
reg296425 := MakeString("B")
reg296426 := MakeString("C")
reg296427 := MakeString("D")
reg296428 := MakeString("E")
reg296429 := MakeString("F")
reg296430 := MakeString("G")
reg296431 := MakeString("H")
reg296432 := MakeString("I")
reg296433 := MakeString("J")
reg296434 := MakeString("K")
reg296435 := MakeString("L")
reg296436 := MakeString("M")
reg296437 := MakeString("N")
reg296438 := MakeString("O")
reg296439 := MakeString("P")
reg296440 := MakeString("Q")
reg296441 := MakeString("R")
reg296442 := MakeString("S")
reg296443 := MakeString("T")
reg296444 := MakeString("U")
reg296445 := MakeString("V")
reg296446 := MakeString("W")
reg296447 := MakeString("X")
reg296448 := MakeString("Y")
reg296449 := MakeString("Z")
reg296450 := MakeString("a")
reg296451 := MakeString("b")
reg296452 := MakeString("c")
reg296453 := MakeString("d")
reg296454 := MakeString("e")
reg296455 := MakeString("f")
reg296456 := MakeString("g")
reg296457 := MakeString("h")
reg296458 := MakeString("i")
reg296459 := MakeString("j")
reg296460 := MakeString("k")
reg296461 := MakeString("l")
reg296462 := MakeString("m")
reg296463 := MakeString("n")
reg296464 := MakeString("o")
reg296465 := MakeString("p")
reg296466 := MakeString("q")
reg296467 := MakeString("r")
reg296468 := MakeString("s")
reg296469 := MakeString("t")
reg296470 := MakeString("u")
reg296471 := MakeString("v")
reg296472 := MakeString("w")
reg296473 := MakeString("x")
reg296474 := MakeString("y")
reg296475 := MakeString("z")
reg296476 := MakeString("=")
reg296477 := MakeString("*")
reg296478 := MakeString("/")
reg296479 := MakeString("+")
reg296480 := MakeString("-")
reg296481 := MakeString("_")
reg296482 := MakeString("?")
reg296483 := MakeString("$")
reg296484 := MakeString("!")
reg296485 := MakeString("@")
reg296486 := MakeString("~")
reg296487 := MakeString(">")
reg296488 := MakeString("<")
reg296489 := MakeString("&")
reg296490 := MakeString("%")
reg296491 := MakeString("{")
reg296492 := MakeString("}")
reg296493 := MakeString(":")
reg296494 := MakeString(";")
reg296495 := MakeString("`")
reg296496 := MakeString("#")
reg296497 := MakeString("'")
reg296498 := MakeString(".")
reg296499 := Nil;
reg296500 := PrimCons(reg296498, reg296499)
reg296501 := PrimCons(reg296497, reg296500)
reg296502 := PrimCons(reg296496, reg296501)
reg296503 := PrimCons(reg296495, reg296502)
reg296504 := PrimCons(reg296494, reg296503)
reg296505 := PrimCons(reg296493, reg296504)
reg296506 := PrimCons(reg296492, reg296505)
reg296507 := PrimCons(reg296491, reg296506)
reg296508 := PrimCons(reg296490, reg296507)
reg296509 := PrimCons(reg296489, reg296508)
reg296510 := PrimCons(reg296488, reg296509)
reg296511 := PrimCons(reg296487, reg296510)
reg296512 := PrimCons(reg296486, reg296511)
reg296513 := PrimCons(reg296485, reg296512)
reg296514 := PrimCons(reg296484, reg296513)
reg296515 := PrimCons(reg296483, reg296514)
reg296516 := PrimCons(reg296482, reg296515)
reg296517 := PrimCons(reg296481, reg296516)
reg296518 := PrimCons(reg296480, reg296517)
reg296519 := PrimCons(reg296479, reg296518)
reg296520 := PrimCons(reg296478, reg296519)
reg296521 := PrimCons(reg296477, reg296520)
reg296522 := PrimCons(reg296476, reg296521)
reg296523 := PrimCons(reg296475, reg296522)
reg296524 := PrimCons(reg296474, reg296523)
reg296525 := PrimCons(reg296473, reg296524)
reg296526 := PrimCons(reg296472, reg296525)
reg296527 := PrimCons(reg296471, reg296526)
reg296528 := PrimCons(reg296470, reg296527)
reg296529 := PrimCons(reg296469, reg296528)
reg296530 := PrimCons(reg296468, reg296529)
reg296531 := PrimCons(reg296467, reg296530)
reg296532 := PrimCons(reg296466, reg296531)
reg296533 := PrimCons(reg296465, reg296532)
reg296534 := PrimCons(reg296464, reg296533)
reg296535 := PrimCons(reg296463, reg296534)
reg296536 := PrimCons(reg296462, reg296535)
reg296537 := PrimCons(reg296461, reg296536)
reg296538 := PrimCons(reg296460, reg296537)
reg296539 := PrimCons(reg296459, reg296538)
reg296540 := PrimCons(reg296458, reg296539)
reg296541 := PrimCons(reg296457, reg296540)
reg296542 := PrimCons(reg296456, reg296541)
reg296543 := PrimCons(reg296455, reg296542)
reg296544 := PrimCons(reg296454, reg296543)
reg296545 := PrimCons(reg296453, reg296544)
reg296546 := PrimCons(reg296452, reg296545)
reg296547 := PrimCons(reg296451, reg296546)
reg296548 := PrimCons(reg296450, reg296547)
reg296549 := PrimCons(reg296449, reg296548)
reg296550 := PrimCons(reg296448, reg296549)
reg296551 := PrimCons(reg296447, reg296550)
reg296552 := PrimCons(reg296446, reg296551)
reg296553 := PrimCons(reg296445, reg296552)
reg296554 := PrimCons(reg296444, reg296553)
reg296555 := PrimCons(reg296443, reg296554)
reg296556 := PrimCons(reg296442, reg296555)
reg296557 := PrimCons(reg296441, reg296556)
reg296558 := PrimCons(reg296440, reg296557)
reg296559 := PrimCons(reg296439, reg296558)
reg296560 := PrimCons(reg296438, reg296559)
reg296561 := PrimCons(reg296437, reg296560)
reg296562 := PrimCons(reg296436, reg296561)
reg296563 := PrimCons(reg296435, reg296562)
reg296564 := PrimCons(reg296434, reg296563)
reg296565 := PrimCons(reg296433, reg296564)
reg296566 := PrimCons(reg296432, reg296565)
reg296567 := PrimCons(reg296431, reg296566)
reg296568 := PrimCons(reg296430, reg296567)
reg296569 := PrimCons(reg296429, reg296568)
reg296570 := PrimCons(reg296428, reg296569)
reg296571 := PrimCons(reg296427, reg296570)
reg296572 := PrimCons(reg296426, reg296571)
reg296573 := PrimCons(reg296425, reg296572)
reg296574 := PrimCons(reg296424, reg296573)
__ctx.TailApply(__defun__element_2, V2906, reg296574)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.alpha?", value: __defun__shen_4alpha_2})

__defun__shen_4alphanums_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2908 := __args[0]
_ = V2908
reg296576 := MakeString("")
reg296577 := PrimEqual(reg296576, V2908)
if reg296577 == True {
reg296578 := True;
__ctx.Return(reg296578)
return
} else {
reg296579 := __e.Call(__defun__shen_4_7string_2, V2908)
if reg296579 == True {
reg296580 := MakeNumber(0)
reg296581 := PrimPos(V2908, reg296580)
reg296582 := __e.Call(__defun__shen_4alphanum_2, reg296581)
if reg296582 == True {
reg296583 := PrimTailString(V2908)
reg296584 := __e.Call(__defun__shen_4alphanums_2, reg296583)
if reg296584 == True {
reg296585 := True;
__ctx.Return(reg296585)
return
} else {
reg296586 := False;
__ctx.Return(reg296586)
return
}
} else {
reg296587 := False;
__ctx.Return(reg296587)
return
}
} else {
reg296588 := MakeSymbol("shen.alphanums?")
__ctx.TailApply(__defun__shen_4f__error, reg296588)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.alphanums?", value: __defun__shen_4alphanums_2})

__defun__shen_4alphanum_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2910 := __args[0]
_ = V2910
reg296590 := __e.Call(__defun__shen_4alpha_2, V2910)
if reg296590 == True {
reg296591 := True;
__ctx.Return(reg296591)
return
} else {
reg296592 := __e.Call(__defun__shen_4digit_2, V2910)
if reg296592 == True {
reg296593 := True;
__ctx.Return(reg296593)
return
} else {
reg296594 := False;
__ctx.Return(reg296594)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.alphanum?", value: __defun__shen_4alphanum_2})

__defun__shen_4digit_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2912 := __args[0]
_ = V2912
reg296595 := MakeString("1")
reg296596 := MakeString("2")
reg296597 := MakeString("3")
reg296598 := MakeString("4")
reg296599 := MakeString("5")
reg296600 := MakeString("6")
reg296601 := MakeString("7")
reg296602 := MakeString("8")
reg296603 := MakeString("9")
reg296604 := MakeString("0")
reg296605 := Nil;
reg296606 := PrimCons(reg296604, reg296605)
reg296607 := PrimCons(reg296603, reg296606)
reg296608 := PrimCons(reg296602, reg296607)
reg296609 := PrimCons(reg296601, reg296608)
reg296610 := PrimCons(reg296600, reg296609)
reg296611 := PrimCons(reg296599, reg296610)
reg296612 := PrimCons(reg296598, reg296611)
reg296613 := PrimCons(reg296597, reg296612)
reg296614 := PrimCons(reg296596, reg296613)
reg296615 := PrimCons(reg296595, reg296614)
__ctx.TailApply(__defun__element_2, V2912, reg296615)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.digit?", value: __defun__shen_4digit_2})

__defun__variable_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2914 := __args[0]
_ = V2914
reg296617 := __e.Call(__defun__boolean_2, V2914)
var reg296629 Obj
if reg296617 == True {
reg296618 := True;
reg296629 = reg296618
} else {
reg296619 := PrimIsNumber(V2914)
var reg296625 Obj
if reg296619 == True {
reg296620 := True;
reg296625 = reg296620
} else {
reg296621 := PrimIsString(V2914)
var reg296624 Obj
if reg296621 == True {
reg296622 := True;
reg296624 = reg296622
} else {
reg296623 := False;
reg296624 = reg296623
}
reg296625 = reg296624
}
var reg296628 Obj
if reg296625 == True {
reg296626 := True;
reg296628 = reg296626
} else {
reg296627 := False;
reg296628 = reg296627
}
reg296629 = reg296628
}
if reg296629 == True {
reg296630 := False;
__ctx.Return(reg296630)
return
} else {
reg296631 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296632 := PrimStr(V2914)
String := reg296632
_ = String
__ctx.TailApply(__defun__shen_4analyse_1variable_2, String)
return
}, 0)
reg296634 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296635 := False;
__ctx.Return(reg296635)
return
}, 1)
reg296636 := __e.Try(reg296631).Catch(reg296634)
__ctx.Return(reg296636)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "variable?", value: __defun__variable_2})

__defun__shen_4analyse_1variable_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2916 := __args[0]
_ = V2916
reg296637 := __e.Call(__defun__shen_4_7string_2, V2916)
if reg296637 == True {
reg296638 := MakeNumber(0)
reg296639 := PrimPos(V2916, reg296638)
reg296640 := __e.Call(__defun__shen_4uppercase_2, reg296639)
if reg296640 == True {
reg296641 := PrimTailString(V2916)
reg296642 := __e.Call(__defun__shen_4alphanums_2, reg296641)
if reg296642 == True {
reg296643 := True;
__ctx.Return(reg296643)
return
} else {
reg296644 := False;
__ctx.Return(reg296644)
return
}
} else {
reg296645 := False;
__ctx.Return(reg296645)
return
}
} else {
reg296646 := MakeSymbol("shen.analyse-variable?")
__ctx.TailApply(__defun__shen_4f__error, reg296646)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.analyse-variable?", value: __defun__shen_4analyse_1variable_2})

__defun__shen_4uppercase_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2918 := __args[0]
_ = V2918
reg296648 := MakeString("A")
reg296649 := MakeString("B")
reg296650 := MakeString("C")
reg296651 := MakeString("D")
reg296652 := MakeString("E")
reg296653 := MakeString("F")
reg296654 := MakeString("G")
reg296655 := MakeString("H")
reg296656 := MakeString("I")
reg296657 := MakeString("J")
reg296658 := MakeString("K")
reg296659 := MakeString("L")
reg296660 := MakeString("M")
reg296661 := MakeString("N")
reg296662 := MakeString("O")
reg296663 := MakeString("P")
reg296664 := MakeString("Q")
reg296665 := MakeString("R")
reg296666 := MakeString("S")
reg296667 := MakeString("T")
reg296668 := MakeString("U")
reg296669 := MakeString("V")
reg296670 := MakeString("W")
reg296671 := MakeString("X")
reg296672 := MakeString("Y")
reg296673 := MakeString("Z")
reg296674 := Nil;
reg296675 := PrimCons(reg296673, reg296674)
reg296676 := PrimCons(reg296672, reg296675)
reg296677 := PrimCons(reg296671, reg296676)
reg296678 := PrimCons(reg296670, reg296677)
reg296679 := PrimCons(reg296669, reg296678)
reg296680 := PrimCons(reg296668, reg296679)
reg296681 := PrimCons(reg296667, reg296680)
reg296682 := PrimCons(reg296666, reg296681)
reg296683 := PrimCons(reg296665, reg296682)
reg296684 := PrimCons(reg296664, reg296683)
reg296685 := PrimCons(reg296663, reg296684)
reg296686 := PrimCons(reg296662, reg296685)
reg296687 := PrimCons(reg296661, reg296686)
reg296688 := PrimCons(reg296660, reg296687)
reg296689 := PrimCons(reg296659, reg296688)
reg296690 := PrimCons(reg296658, reg296689)
reg296691 := PrimCons(reg296657, reg296690)
reg296692 := PrimCons(reg296656, reg296691)
reg296693 := PrimCons(reg296655, reg296692)
reg296694 := PrimCons(reg296654, reg296693)
reg296695 := PrimCons(reg296653, reg296694)
reg296696 := PrimCons(reg296652, reg296695)
reg296697 := PrimCons(reg296651, reg296696)
reg296698 := PrimCons(reg296650, reg296697)
reg296699 := PrimCons(reg296649, reg296698)
reg296700 := PrimCons(reg296648, reg296699)
__ctx.TailApply(__defun__element_2, V2918, reg296700)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.uppercase?", value: __defun__shen_4uppercase_2})

__defun__gensym = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2920 := __args[0]
_ = V2920
reg296702 := MakeSymbol("shen.*gensym*")
reg296703 := MakeNumber(1)
reg296704 := MakeSymbol("shen.*gensym*")
reg296705 := PrimValue(reg296704)
reg296706 := PrimNumberAdd(reg296703, reg296705)
reg296707 := PrimSet(reg296702, reg296706)
__ctx.TailApply(__defun__concat, V2920, reg296707)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "gensym", value: __defun__gensym})

__defun__concat = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2923 := __args[0]
_ = V2923
V2924 := __args[1]
_ = V2924
reg296709 := PrimStr(V2923)
reg296710 := PrimStr(V2924)
reg296711 := PrimStringConcat(reg296709, reg296710)
reg296712 := PrimIntern(reg296711)
__ctx.Return(reg296712)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "concat", value: __defun__concat})

__defun___8p = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2927 := __args[0]
_ = V2927
V2928 := __args[1]
_ = V2928
reg296713 := MakeNumber(3)
reg296714 := PrimAbsvector(reg296713)
Vector := reg296714
_ = Vector
reg296715 := MakeNumber(0)
reg296716 := MakeSymbol("shen.tuple")
reg296717 := PrimVectorSet(Vector, reg296715, reg296716)
Tag := reg296717
_ = Tag
reg296718 := MakeNumber(1)
reg296719 := PrimVectorSet(Vector, reg296718, V2927)
Fst := reg296719
_ = Fst
reg296720 := MakeNumber(2)
reg296721 := PrimVectorSet(Vector, reg296720, V2928)
Snd := reg296721
_ = Snd
__ctx.Return(Vector)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "@p", value: __defun___8p})

__defun__fst = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2930 := __args[0]
_ = V2930
reg296722 := MakeNumber(1)
reg296723 := PrimVectorGet(V2930, reg296722)
__ctx.Return(reg296723)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "fst", value: __defun__fst})

__defun__snd = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2932 := __args[0]
_ = V2932
reg296724 := MakeNumber(2)
reg296725 := PrimVectorGet(V2932, reg296724)
__ctx.Return(reg296725)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "snd", value: __defun__snd})

__defun__tuple_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2934 := __args[0]
_ = V2934
reg296726 := PrimIsVector(V2934)
if reg296726 == True {
reg296727 := MakeSymbol("shen.tuple")
reg296728 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296729 := MakeNumber(0)
reg296730 := PrimVectorGet(V2934, reg296729)
__ctx.Return(reg296730)
return
}, 0)
reg296731 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296732 := MakeSymbol("shen.not-tuple")
__ctx.Return(reg296732)
return
}, 1)
reg296733 := __e.Try(reg296728).Catch(reg296731)
reg296734 := PrimEqual(reg296727, reg296733)
if reg296734 == True {
reg296735 := True;
__ctx.Return(reg296735)
return
} else {
reg296736 := False;
__ctx.Return(reg296736)
return
}
} else {
reg296737 := False;
__ctx.Return(reg296737)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "tuple?", value: __defun__tuple_2})

__defun__append = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2937 := __args[0]
_ = V2937
V2938 := __args[1]
_ = V2938
reg296738 := Nil;
reg296739 := PrimEqual(reg296738, V2937)
if reg296739 == True {
__ctx.Return(V2938)
return
} else {
reg296740 := PrimIsPair(V2937)
if reg296740 == True {
reg296741 := PrimHead(V2937)
reg296742 := PrimTail(V2937)
reg296743 := __e.Call(__defun__append, reg296742, V2938)
reg296744 := PrimCons(reg296741, reg296743)
__ctx.Return(reg296744)
return
} else {
reg296745 := MakeSymbol("append")
__ctx.TailApply(__defun__shen_4f__error, reg296745)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "append", value: __defun__append})

__defun___8v = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2941 := __args[0]
_ = V2941
V2942 := __args[1]
_ = V2942
reg296747 := __e.Call(__defun__limit, V2942)
Limit := reg296747
_ = Limit
reg296748 := MakeNumber(1)
reg296749 := PrimNumberAdd(Limit, reg296748)
reg296750 := __e.Call(__defun__vector, reg296749)
NewVector := reg296750
_ = NewVector
reg296751 := MakeNumber(1)
reg296752 := __e.Call(__defun__vector_1_6, NewVector, reg296751, V2941)
X_7NewVector := reg296752
_ = X_7NewVector
reg296753 := MakeNumber(0)
reg296754 := PrimEqual(Limit, reg296753)
if reg296754 == True {
__ctx.Return(X_7NewVector)
return
} else {
reg296755 := MakeNumber(1)
__ctx.TailApply(__defun__shen_4_8v_1help, V2942, reg296755, Limit, X_7NewVector)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "@v", value: __defun___8v})

__defun__shen_4_8v_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2948 := __args[0]
_ = V2948
V2949 := __args[1]
_ = V2949
V2950 := __args[2]
_ = V2950
V2951 := __args[3]
_ = V2951
reg296757 := PrimEqual(V2950, V2949)
if reg296757 == True {
reg296758 := MakeNumber(1)
reg296759 := PrimNumberAdd(V2950, reg296758)
__ctx.TailApply(__defun__shen_4copyfromvector, V2948, V2951, V2950, reg296759)
return
} else {
reg296761 := MakeNumber(1)
reg296762 := PrimNumberAdd(V2949, reg296761)
reg296763 := MakeNumber(1)
reg296764 := PrimNumberAdd(V2949, reg296763)
reg296765 := __e.Call(__defun__shen_4copyfromvector, V2948, V2951, V2949, reg296764)
__ctx.TailApply(__defun__shen_4_8v_1help, V2948, reg296762, V2950, reg296765)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.@v-help", value: __defun__shen_4_8v_1help})

__defun__shen_4copyfromvector = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2956 := __args[0]
_ = V2956
V2957 := __args[1]
_ = V2957
V2958 := __args[2]
_ = V2958
V2959 := __args[3]
_ = V2959
reg296767 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296768 := __e.Call(__defun___5_1vector, V2956, V2958)
__ctx.TailApply(__defun__vector_1_6, V2957, V2959, reg296768)
return
}, 0)
reg296770 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.Return(V2957)
return
}, 1)
reg296771 := __e.Try(reg296767).Catch(reg296770)
__ctx.Return(reg296771)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.copyfromvector", value: __defun__shen_4copyfromvector})

__defun__hdv = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2961 := __args[0]
_ = V2961
reg296772 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg296773 := MakeNumber(1)
__ctx.TailApply(__defun___5_1vector, V2961, reg296773)
return
}, 0)
reg296775 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296776 := MakeString("hdv needs a non-empty vector as an argument; not ")
reg296777 := MakeString("\n")
reg296778 := MakeSymbol("shen.s")
reg296779 := __e.Call(__defun__shen_4app, V2961, reg296777, reg296778)
reg296780 := PrimStringConcat(reg296776, reg296779)
reg296781 := PrimSimpleError(reg296780)
__ctx.Return(reg296781)
return
}, 1)
reg296782 := __e.Try(reg296772).Catch(reg296775)
__ctx.Return(reg296782)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "hdv", value: __defun__hdv})

__defun__tlv = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2963 := __args[0]
_ = V2963
reg296783 := __e.Call(__defun__limit, V2963)
Limit := reg296783
_ = Limit
reg296784 := MakeNumber(0)
reg296785 := PrimEqual(Limit, reg296784)
if reg296785 == True {
reg296786 := MakeString("cannot take the tail of the empty vector\n")
reg296787 := PrimSimpleError(reg296786)
__ctx.Return(reg296787)
return
} else {
reg296788 := MakeNumber(1)
reg296789 := PrimEqual(Limit, reg296788)
if reg296789 == True {
reg296790 := MakeNumber(0)
__ctx.TailApply(__defun__vector, reg296790)
return
} else {
reg296792 := MakeNumber(1)
reg296793 := PrimNumberSubtract(Limit, reg296792)
reg296794 := __e.Call(__defun__vector, reg296793)
NewVector := reg296794
_ = NewVector
reg296795 := MakeNumber(2)
reg296796 := MakeNumber(1)
reg296797 := PrimNumberSubtract(Limit, reg296796)
reg296798 := __e.Call(__defun__vector, reg296797)
__ctx.TailApply(__defun__shen_4tlv_1help, V2963, reg296795, Limit, reg296798)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "tlv", value: __defun__tlv})

__defun__shen_4tlv_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2969 := __args[0]
_ = V2969
V2970 := __args[1]
_ = V2970
V2971 := __args[2]
_ = V2971
V2972 := __args[3]
_ = V2972
reg296800 := PrimEqual(V2971, V2970)
if reg296800 == True {
reg296801 := MakeNumber(1)
reg296802 := PrimNumberSubtract(V2971, reg296801)
__ctx.TailApply(__defun__shen_4copyfromvector, V2969, V2972, V2971, reg296802)
return
} else {
reg296804 := MakeNumber(1)
reg296805 := PrimNumberAdd(V2970, reg296804)
reg296806 := MakeNumber(1)
reg296807 := PrimNumberSubtract(V2970, reg296806)
reg296808 := __e.Call(__defun__shen_4copyfromvector, V2969, V2972, V2970, reg296807)
__ctx.TailApply(__defun__shen_4tlv_1help, V2969, reg296805, V2971, reg296808)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.tlv-help", value: __defun__shen_4tlv_1help})

__defun__assoc = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2984 := __args[0]
_ = V2984
V2985 := __args[1]
_ = V2985
reg296810 := Nil;
reg296811 := PrimEqual(reg296810, V2985)
if reg296811 == True {
reg296812 := Nil;
__ctx.Return(reg296812)
return
} else {
reg296813 := PrimIsPair(V2985)
var reg296828 Obj
if reg296813 == True {
reg296814 := PrimHead(V2985)
reg296815 := PrimIsPair(reg296814)
var reg296823 Obj
if reg296815 == True {
reg296816 := PrimHead(V2985)
reg296817 := PrimHead(reg296816)
reg296818 := PrimEqual(reg296817, V2984)
var reg296821 Obj
if reg296818 == True {
reg296819 := True;
reg296821 = reg296819
} else {
reg296820 := False;
reg296821 = reg296820
}
reg296823 = reg296821
} else {
reg296822 := False;
reg296823 = reg296822
}
var reg296826 Obj
if reg296823 == True {
reg296824 := True;
reg296826 = reg296824
} else {
reg296825 := False;
reg296826 = reg296825
}
reg296828 = reg296826
} else {
reg296827 := False;
reg296828 = reg296827
}
if reg296828 == True {
reg296829 := PrimHead(V2985)
__ctx.Return(reg296829)
return
} else {
reg296830 := PrimIsPair(V2985)
if reg296830 == True {
reg296831 := PrimTail(V2985)
__ctx.TailApply(__defun__assoc, V2984, reg296831)
return
} else {
reg296833 := MakeSymbol("assoc")
__ctx.TailApply(__defun__shen_4f__error, reg296833)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "assoc", value: __defun__assoc})

__defun__shen_4assoc_1set = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2992 := __args[0]
_ = V2992
V2993 := __args[1]
_ = V2993
V2994 := __args[2]
_ = V2994
reg296835 := Nil;
reg296836 := PrimEqual(reg296835, V2994)
if reg296836 == True {
reg296837 := PrimCons(V2992, V2993)
reg296838 := Nil;
reg296839 := PrimCons(reg296837, reg296838)
__ctx.Return(reg296839)
return
} else {
reg296840 := PrimIsPair(V2994)
var reg296855 Obj
if reg296840 == True {
reg296841 := PrimHead(V2994)
reg296842 := PrimIsPair(reg296841)
var reg296850 Obj
if reg296842 == True {
reg296843 := PrimHead(V2994)
reg296844 := PrimHead(reg296843)
reg296845 := PrimEqual(reg296844, V2992)
var reg296848 Obj
if reg296845 == True {
reg296846 := True;
reg296848 = reg296846
} else {
reg296847 := False;
reg296848 = reg296847
}
reg296850 = reg296848
} else {
reg296849 := False;
reg296850 = reg296849
}
var reg296853 Obj
if reg296850 == True {
reg296851 := True;
reg296853 = reg296851
} else {
reg296852 := False;
reg296853 = reg296852
}
reg296855 = reg296853
} else {
reg296854 := False;
reg296855 = reg296854
}
if reg296855 == True {
reg296856 := PrimHead(V2994)
reg296857 := PrimHead(reg296856)
reg296858 := PrimCons(reg296857, V2993)
reg296859 := PrimTail(V2994)
reg296860 := PrimCons(reg296858, reg296859)
__ctx.Return(reg296860)
return
} else {
reg296861 := PrimIsPair(V2994)
if reg296861 == True {
reg296862 := PrimHead(V2994)
reg296863 := PrimTail(V2994)
reg296864 := __e.Call(__defun__shen_4assoc_1set, V2992, V2993, reg296863)
reg296865 := PrimCons(reg296862, reg296864)
__ctx.Return(reg296865)
return
} else {
reg296866 := MakeSymbol("shen.assoc-set")
__ctx.TailApply(__defun__shen_4f__error, reg296866)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.assoc-set", value: __defun__shen_4assoc_1set})

__defun__shen_4assoc_1rm = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3000 := __args[0]
_ = V3000
V3001 := __args[1]
_ = V3001
reg296868 := Nil;
reg296869 := PrimEqual(reg296868, V3001)
if reg296869 == True {
reg296870 := Nil;
__ctx.Return(reg296870)
return
} else {
reg296871 := PrimIsPair(V3001)
var reg296886 Obj
if reg296871 == True {
reg296872 := PrimHead(V3001)
reg296873 := PrimIsPair(reg296872)
var reg296881 Obj
if reg296873 == True {
reg296874 := PrimHead(V3001)
reg296875 := PrimHead(reg296874)
reg296876 := PrimEqual(reg296875, V3000)
var reg296879 Obj
if reg296876 == True {
reg296877 := True;
reg296879 = reg296877
} else {
reg296878 := False;
reg296879 = reg296878
}
reg296881 = reg296879
} else {
reg296880 := False;
reg296881 = reg296880
}
var reg296884 Obj
if reg296881 == True {
reg296882 := True;
reg296884 = reg296882
} else {
reg296883 := False;
reg296884 = reg296883
}
reg296886 = reg296884
} else {
reg296885 := False;
reg296886 = reg296885
}
if reg296886 == True {
reg296887 := PrimTail(V3001)
__ctx.Return(reg296887)
return
} else {
reg296888 := PrimIsPair(V3001)
if reg296888 == True {
reg296889 := PrimHead(V3001)
reg296890 := PrimTail(V3001)
reg296891 := __e.Call(__defun__shen_4assoc_1rm, V3000, reg296890)
reg296892 := PrimCons(reg296889, reg296891)
__ctx.Return(reg296892)
return
} else {
reg296893 := MakeSymbol("shen.assoc-rm")
__ctx.TailApply(__defun__shen_4f__error, reg296893)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.assoc-rm", value: __defun__shen_4assoc_1rm})

__defun__boolean_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3007 := __args[0]
_ = V3007
reg296895 := True;
reg296896 := PrimEqual(reg296895, V3007)
if reg296896 == True {
reg296897 := True;
__ctx.Return(reg296897)
return
} else {
reg296898 := False;
reg296899 := PrimEqual(reg296898, V3007)
if reg296899 == True {
reg296900 := True;
__ctx.Return(reg296900)
return
} else {
reg296901 := False;
__ctx.Return(reg296901)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "boolean?", value: __defun__boolean_2})

__defun__nl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3009 := __args[0]
_ = V3009
reg296902 := MakeNumber(0)
reg296903 := PrimEqual(reg296902, V3009)
if reg296903 == True {
reg296904 := MakeNumber(0)
__ctx.Return(reg296904)
return
} else {
reg296905 := MakeString("\n")
reg296906 := __e.Call(__defun__stoutput)
reg296907 := __e.Call(__defun__shen_4prhush, reg296905, reg296906)
_ = reg296907
reg296908 := MakeNumber(1)
reg296909 := PrimNumberSubtract(V3009, reg296908)
__ctx.TailApply(__defun__nl, reg296909)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "nl", value: __defun__nl})

__defun__difference = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3014 := __args[0]
_ = V3014
V3015 := __args[1]
_ = V3015
reg296911 := Nil;
reg296912 := PrimEqual(reg296911, V3014)
if reg296912 == True {
reg296913 := Nil;
__ctx.Return(reg296913)
return
} else {
reg296914 := PrimIsPair(V3014)
if reg296914 == True {
reg296915 := PrimHead(V3014)
reg296916 := __e.Call(__defun__element_2, reg296915, V3015)
if reg296916 == True {
reg296917 := PrimTail(V3014)
__ctx.TailApply(__defun__difference, reg296917, V3015)
return
} else {
reg296919 := PrimHead(V3014)
reg296920 := PrimTail(V3014)
reg296921 := __e.Call(__defun__difference, reg296920, V3015)
reg296922 := PrimCons(reg296919, reg296921)
__ctx.Return(reg296922)
return
}
} else {
reg296923 := MakeSymbol("difference")
__ctx.TailApply(__defun__shen_4f__error, reg296923)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "difference", value: __defun__difference})

__defun__do = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3018 := __args[0]
_ = V3018
V3019 := __args[1]
_ = V3019
__ctx.Return(V3019)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "do", value: __defun__do})

__defun__element_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3031 := __args[0]
_ = V3031
V3032 := __args[1]
_ = V3032
reg296925 := Nil;
reg296926 := PrimEqual(reg296925, V3032)
if reg296926 == True {
reg296927 := False;
__ctx.Return(reg296927)
return
} else {
reg296928 := PrimIsPair(V3032)
var reg296935 Obj
if reg296928 == True {
reg296929 := PrimHead(V3032)
reg296930 := PrimEqual(reg296929, V3031)
var reg296933 Obj
if reg296930 == True {
reg296931 := True;
reg296933 = reg296931
} else {
reg296932 := False;
reg296933 = reg296932
}
reg296935 = reg296933
} else {
reg296934 := False;
reg296935 = reg296934
}
if reg296935 == True {
reg296936 := True;
__ctx.Return(reg296936)
return
} else {
reg296937 := PrimIsPair(V3032)
if reg296937 == True {
reg296938 := PrimTail(V3032)
__ctx.TailApply(__defun__element_2, V3031, reg296938)
return
} else {
reg296940 := MakeSymbol("element?")
__ctx.TailApply(__defun__shen_4f__error, reg296940)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "element?", value: __defun__element_2})

__defun__empty_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3038 := __args[0]
_ = V3038
reg296942 := Nil;
reg296943 := PrimEqual(reg296942, V3038)
if reg296943 == True {
reg296944 := True;
__ctx.Return(reg296944)
return
} else {
reg296945 := False;
__ctx.Return(reg296945)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "empty?", value: __defun__empty_2})

__defun__fix = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3041 := __args[0]
_ = V3041
V3042 := __args[1]
_ = V3042
reg296946 := __e.Call(V3041, V3042)
__ctx.TailApply(__defun__shen_4fix_1help, V3041, V3042, reg296946)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "fix", value: __defun__fix})

__defun__shen_4fix_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3053 := __args[0]
_ = V3053
V3054 := __args[1]
_ = V3054
V3055 := __args[2]
_ = V3055
reg296948 := PrimEqual(V3055, V3054)
if reg296948 == True {
__ctx.Return(V3055)
return
} else {
reg296949 := __e.Call(V3053, V3055)
__ctx.TailApply(__defun__shen_4fix_1help, V3053, V3055, reg296949)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.fix-help", value: __defun__shen_4fix_1help})

__defun__put = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3060 := __args[0]
_ = V3060
V3061 := __args[1]
_ = V3061
V3062 := __args[2]
_ = V3062
V3063 := __args[3]
_ = V3063
reg296951 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4_5_1dict, V3063, V3060)
return
}, 0)
reg296953 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296954 := Nil;
__ctx.Return(reg296954)
return
}, 1)
reg296955 := __e.Try(reg296951).Catch(reg296953)
Curr := reg296955
_ = Curr
reg296956 := __e.Call(__defun__shen_4assoc_1set, V3061, V3062, Curr)
Added := reg296956
_ = Added
reg296957 := __e.Call(__defun__shen_4dict_1_6, V3063, V3060, Added)
Update := reg296957
_ = Update
__ctx.Return(V3062)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "put", value: __defun__put})

__defun__unput = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3067 := __args[0]
_ = V3067
V3068 := __args[1]
_ = V3068
V3069 := __args[2]
_ = V3069
reg296958 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4_5_1dict, V3069, V3067)
return
}, 0)
reg296960 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296961 := Nil;
__ctx.Return(reg296961)
return
}, 1)
reg296962 := __e.Try(reg296958).Catch(reg296960)
Curr := reg296962
_ = Curr
reg296963 := __e.Call(__defun__shen_4assoc_1rm, V3068, Curr)
Removed := reg296963
_ = Removed
reg296964 := __e.Call(__defun__shen_4dict_1_6, V3069, V3067, Removed)
Update := reg296964
_ = Update
__ctx.Return(V3067)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "unput", value: __defun__unput})

__defun__get = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3073 := __args[0]
_ = V3073
V3074 := __args[1]
_ = V3074
V3075 := __args[2]
_ = V3075
reg296965 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4_5_1dict, V3075, V3073)
return
}, 0)
reg296967 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg296968 := Nil;
__ctx.Return(reg296968)
return
}, 1)
reg296969 := __e.Try(reg296965).Catch(reg296967)
Entry := reg296969
_ = Entry
reg296970 := __e.Call(__defun__assoc, V3074, Entry)
Result := reg296970
_ = Result
reg296971 := __e.Call(__defun__empty_2, Result)
if reg296971 == True {
reg296972 := MakeString("value not found\n")
reg296973 := PrimSimpleError(reg296972)
__ctx.Return(reg296973)
return
} else {
reg296974 := PrimTail(Result)
__ctx.Return(reg296974)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "get", value: __defun__get})

__defun__hash = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3078 := __args[0]
_ = V3078
V3079 := __args[1]
_ = V3079
reg296975 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg296976 := PrimStringToNumber(X)
__ctx.Return(reg296976)
return
}, 1)
reg296977 := __e.Call(__defun__explode, V3078)
reg296978 := __e.Call(__defun__map, reg296975, reg296977)
reg296979 := __e.Call(__defun__sum, reg296978)
__ctx.TailApply(__defun__shen_4mod, reg296979, V3079)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "hash", value: __defun__hash})

__defun__shen_4mod = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3082 := __args[0]
_ = V3082
V3083 := __args[1]
_ = V3083
reg296981 := Nil;
reg296982 := PrimCons(V3083, reg296981)
reg296983 := __e.Call(__defun__shen_4multiples, V3082, reg296982)
__ctx.TailApply(__defun__shen_4modh, V3082, reg296983)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.mod", value: __defun__shen_4mod})

__defun__shen_4multiples = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3086 := __args[0]
_ = V3086
V3087 := __args[1]
_ = V3087
reg296985 := PrimIsPair(V3087)
var reg296992 Obj
if reg296985 == True {
reg296986 := PrimHead(V3087)
reg296987 := PrimGreatThan(reg296986, V3086)
var reg296990 Obj
if reg296987 == True {
reg296988 := True;
reg296990 = reg296988
} else {
reg296989 := False;
reg296990 = reg296989
}
reg296992 = reg296990
} else {
reg296991 := False;
reg296992 = reg296991
}
if reg296992 == True {
reg296993 := PrimTail(V3087)
__ctx.Return(reg296993)
return
} else {
reg296994 := PrimIsPair(V3087)
if reg296994 == True {
reg296995 := MakeNumber(2)
reg296996 := PrimHead(V3087)
reg296997 := PrimNumberMultiply(reg296995, reg296996)
reg296998 := PrimCons(reg296997, V3087)
__ctx.TailApply(__defun__shen_4multiples, V3086, reg296998)
return
} else {
reg297000 := MakeSymbol("shen.multiples")
__ctx.TailApply(__defun__shen_4f__error, reg297000)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.multiples", value: __defun__shen_4multiples})

__defun__shen_4modh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3092 := __args[0]
_ = V3092
V3093 := __args[1]
_ = V3093
reg297002 := MakeNumber(0)
reg297003 := PrimEqual(reg297002, V3092)
if reg297003 == True {
reg297004 := MakeNumber(0)
__ctx.Return(reg297004)
return
} else {
reg297005 := Nil;
reg297006 := PrimEqual(reg297005, V3093)
if reg297006 == True {
__ctx.Return(V3092)
return
} else {
reg297007 := PrimIsPair(V3093)
var reg297014 Obj
if reg297007 == True {
reg297008 := PrimHead(V3093)
reg297009 := PrimGreatThan(reg297008, V3092)
var reg297012 Obj
if reg297009 == True {
reg297010 := True;
reg297012 = reg297010
} else {
reg297011 := False;
reg297012 = reg297011
}
reg297014 = reg297012
} else {
reg297013 := False;
reg297014 = reg297013
}
if reg297014 == True {
reg297015 := PrimTail(V3093)
reg297016 := __e.Call(__defun__empty_2, reg297015)
if reg297016 == True {
__ctx.Return(V3092)
return
} else {
reg297017 := PrimTail(V3093)
__ctx.TailApply(__defun__shen_4modh, V3092, reg297017)
return
}
} else {
reg297019 := PrimIsPair(V3093)
if reg297019 == True {
reg297020 := PrimHead(V3093)
reg297021 := PrimNumberSubtract(V3092, reg297020)
__ctx.TailApply(__defun__shen_4modh, reg297021, V3093)
return
} else {
reg297023 := MakeSymbol("shen.modh")
__ctx.TailApply(__defun__shen_4f__error, reg297023)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.modh", value: __defun__shen_4modh})

__defun__sum = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3095 := __args[0]
_ = V3095
reg297025 := Nil;
reg297026 := PrimEqual(reg297025, V3095)
if reg297026 == True {
reg297027 := MakeNumber(0)
__ctx.Return(reg297027)
return
} else {
reg297028 := PrimIsPair(V3095)
if reg297028 == True {
reg297029 := PrimHead(V3095)
reg297030 := PrimTail(V3095)
reg297031 := __e.Call(__defun__sum, reg297030)
reg297032 := PrimNumberAdd(reg297029, reg297031)
__ctx.Return(reg297032)
return
} else {
reg297033 := MakeSymbol("sum")
__ctx.TailApply(__defun__shen_4f__error, reg297033)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "sum", value: __defun__sum})

__defun__head = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3103 := __args[0]
_ = V3103
reg297035 := PrimIsPair(V3103)
if reg297035 == True {
reg297036 := PrimHead(V3103)
__ctx.Return(reg297036)
return
} else {
reg297037 := MakeString("head expects a non-empty list")
reg297038 := PrimSimpleError(reg297037)
__ctx.Return(reg297038)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "head", value: __defun__head})

__defun__tail = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3111 := __args[0]
_ = V3111
reg297039 := PrimIsPair(V3111)
if reg297039 == True {
reg297040 := PrimTail(V3111)
__ctx.Return(reg297040)
return
} else {
reg297041 := MakeString("tail expects a non-empty list")
reg297042 := PrimSimpleError(reg297041)
__ctx.Return(reg297042)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "tail", value: __defun__tail})

__defun__hdstr = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3113 := __args[0]
_ = V3113
reg297043 := MakeNumber(0)
reg297044 := PrimPos(V3113, reg297043)
__ctx.Return(reg297044)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "hdstr", value: __defun__hdstr})

__defun__intersection = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3118 := __args[0]
_ = V3118
V3119 := __args[1]
_ = V3119
reg297045 := Nil;
reg297046 := PrimEqual(reg297045, V3118)
if reg297046 == True {
reg297047 := Nil;
__ctx.Return(reg297047)
return
} else {
reg297048 := PrimIsPair(V3118)
if reg297048 == True {
reg297049 := PrimHead(V3118)
reg297050 := __e.Call(__defun__element_2, reg297049, V3119)
if reg297050 == True {
reg297051 := PrimHead(V3118)
reg297052 := PrimTail(V3118)
reg297053 := __e.Call(__defun__intersection, reg297052, V3119)
reg297054 := PrimCons(reg297051, reg297053)
__ctx.Return(reg297054)
return
} else {
reg297055 := PrimTail(V3118)
__ctx.TailApply(__defun__intersection, reg297055, V3119)
return
}
} else {
reg297057 := MakeSymbol("intersection")
__ctx.TailApply(__defun__shen_4f__error, reg297057)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "intersection", value: __defun__intersection})

__defun__reverse = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3121 := __args[0]
_ = V3121
reg297059 := Nil;
__ctx.TailApply(__defun__shen_4reverse__help, V3121, reg297059)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "reverse", value: __defun__reverse})

__defun__shen_4reverse__help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3124 := __args[0]
_ = V3124
V3125 := __args[1]
_ = V3125
reg297061 := Nil;
reg297062 := PrimEqual(reg297061, V3124)
if reg297062 == True {
__ctx.Return(V3125)
return
} else {
reg297063 := PrimIsPair(V3124)
if reg297063 == True {
reg297064 := PrimTail(V3124)
reg297065 := PrimHead(V3124)
reg297066 := PrimCons(reg297065, V3125)
__ctx.TailApply(__defun__shen_4reverse__help, reg297064, reg297066)
return
} else {
reg297068 := MakeSymbol("shen.reverse_help")
__ctx.TailApply(__defun__shen_4f__error, reg297068)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.reverse_help", value: __defun__shen_4reverse__help})

__defun__union = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3128 := __args[0]
_ = V3128
V3129 := __args[1]
_ = V3129
reg297070 := Nil;
reg297071 := PrimEqual(reg297070, V3128)
if reg297071 == True {
__ctx.Return(V3129)
return
} else {
reg297072 := PrimIsPair(V3128)
if reg297072 == True {
reg297073 := PrimHead(V3128)
reg297074 := __e.Call(__defun__element_2, reg297073, V3129)
if reg297074 == True {
reg297075 := PrimTail(V3128)
__ctx.TailApply(__defun__union, reg297075, V3129)
return
} else {
reg297077 := PrimHead(V3128)
reg297078 := PrimTail(V3128)
reg297079 := __e.Call(__defun__union, reg297078, V3129)
reg297080 := PrimCons(reg297077, reg297079)
__ctx.Return(reg297080)
return
}
} else {
reg297081 := MakeSymbol("union")
__ctx.TailApply(__defun__shen_4f__error, reg297081)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "union", value: __defun__union})

__defun__y_1or_1n_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3131 := __args[0]
_ = V3131
reg297083 := __e.Call(__defun__shen_4proc_1nl, V3131)
reg297084 := __e.Call(__defun__stoutput)
reg297085 := __e.Call(__defun__shen_4prhush, reg297083, reg297084)
Message := reg297085
_ = Message
reg297086 := MakeString(" (y/n) ")
reg297087 := __e.Call(__defun__stoutput)
reg297088 := __e.Call(__defun__shen_4prhush, reg297086, reg297087)
Y_1or_1N := reg297088
_ = Y_1or_1N
reg297089 := __e.Call(__defun__stinput)
reg297090 := __e.Call(__defun__read, reg297089)
reg297091 := MakeString("")
reg297092 := MakeSymbol("shen.s")
reg297093 := __e.Call(__defun__shen_4app, reg297090, reg297091, reg297092)
Input := reg297093
_ = Input
reg297094 := MakeString("y")
reg297095 := PrimEqual(reg297094, Input)
if reg297095 == True {
reg297096 := True;
__ctx.Return(reg297096)
return
} else {
reg297097 := MakeString("n")
reg297098 := PrimEqual(reg297097, Input)
if reg297098 == True {
reg297099 := False;
__ctx.Return(reg297099)
return
} else {
reg297100 := MakeString("please answer y or n\n")
reg297101 := __e.Call(__defun__stoutput)
reg297102 := __e.Call(__defun__shen_4prhush, reg297100, reg297101)
_ = reg297102
__ctx.TailApply(__defun__y_1or_1n_2, V3131)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "y-or-n?", value: __defun__y_1or_1n_2})

__defun__not = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3133 := __args[0]
_ = V3133
if V3133 == True {
reg297104 := False;
__ctx.Return(reg297104)
return
} else {
reg297105 := True;
__ctx.Return(reg297105)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "not", value: __defun__not})

__defun__subst = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3146 := __args[0]
_ = V3146
V3147 := __args[1]
_ = V3147
V3148 := __args[2]
_ = V3148
reg297106 := PrimEqual(V3148, V3147)
if reg297106 == True {
__ctx.Return(V3146)
return
} else {
reg297107 := PrimIsPair(V3148)
if reg297107 == True {
reg297108 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
W := __args[0]
_ = W
__ctx.TailApply(__defun__subst, V3146, V3147, W)
return
}, 1)
__ctx.TailApply(__defun__map, reg297108, V3148)
return
} else {
__ctx.Return(V3148)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "subst", value: __defun__subst})

__defun__explode = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3150 := __args[0]
_ = V3150
reg297111 := MakeString("")
reg297112 := MakeSymbol("shen.a")
reg297113 := __e.Call(__defun__shen_4app, V3150, reg297111, reg297112)
__ctx.TailApply(__defun__shen_4explode_1h, reg297113)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "explode", value: __defun__explode})

__defun__shen_4explode_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3152 := __args[0]
_ = V3152
reg297115 := MakeString("")
reg297116 := PrimEqual(reg297115, V3152)
if reg297116 == True {
reg297117 := Nil;
__ctx.Return(reg297117)
return
} else {
reg297118 := __e.Call(__defun__shen_4_7string_2, V3152)
if reg297118 == True {
reg297119 := MakeNumber(0)
reg297120 := PrimPos(V3152, reg297119)
reg297121 := PrimTailString(V3152)
reg297122 := __e.Call(__defun__shen_4explode_1h, reg297121)
reg297123 := PrimCons(reg297120, reg297122)
__ctx.Return(reg297123)
return
} else {
reg297124 := MakeSymbol("shen.explode-h")
__ctx.TailApply(__defun__shen_4f__error, reg297124)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.explode-h", value: __defun__shen_4explode_1h})

__defun__cd = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3154 := __args[0]
_ = V3154
reg297126 := MakeSymbol("*home-directory*")
reg297127 := MakeString("")
reg297128 := PrimEqual(V3154, reg297127)
var reg297133 Obj
if reg297128 == True {
reg297129 := MakeString("")
reg297133 = reg297129
} else {
reg297130 := MakeString("/")
reg297131 := MakeSymbol("shen.a")
reg297132 := __e.Call(__defun__shen_4app, V3154, reg297130, reg297131)
reg297133 = reg297132
}
reg297134 := PrimSet(reg297126, reg297133)
__ctx.Return(reg297134)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "cd", value: __defun__cd})

__defun__shen_4for_1each = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3157 := __args[0]
_ = V3157
V3158 := __args[1]
_ = V3158
reg297135 := Nil;
reg297136 := PrimEqual(reg297135, V3158)
if reg297136 == True {
reg297137 := True;
__ctx.Return(reg297137)
return
} else {
reg297138 := PrimIsPair(V3158)
if reg297138 == True {
reg297139 := PrimHead(V3158)
reg297140 := __e.Call(V3157, reg297139)
__ := reg297140
_ = __
reg297141 := PrimTail(V3158)
__ctx.TailApply(__defun__shen_4for_1each, V3157, reg297141)
return
} else {
reg297143 := MakeSymbol("shen.for-each")
__ctx.TailApply(__defun__shen_4f__error, reg297143)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.for-each", value: __defun__shen_4for_1each})

__defun__map = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3161 := __args[0]
_ = V3161
V3162 := __args[1]
_ = V3162
reg297145 := Nil;
__ctx.TailApply(__defun__shen_4map_1h, V3161, V3162, reg297145)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "map", value: __defun__map})

__defun__shen_4map_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3168 := __args[0]
_ = V3168
V3169 := __args[1]
_ = V3169
V3170 := __args[2]
_ = V3170
reg297147 := Nil;
reg297148 := PrimEqual(reg297147, V3169)
if reg297148 == True {
__ctx.TailApply(__defun__reverse, V3170)
return
} else {
reg297150 := PrimIsPair(V3169)
if reg297150 == True {
reg297151 := PrimTail(V3169)
reg297152 := PrimHead(V3169)
reg297153 := __e.Call(V3168, reg297152)
reg297154 := PrimCons(reg297153, V3170)
__ctx.TailApply(__defun__shen_4map_1h, V3168, reg297151, reg297154)
return
} else {
reg297156 := MakeSymbol("shen.map-h")
__ctx.TailApply(__defun__shen_4f__error, reg297156)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.map-h", value: __defun__shen_4map_1h})

__defun__length = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3172 := __args[0]
_ = V3172
reg297158 := MakeNumber(0)
__ctx.TailApply(__defun__shen_4length_1h, V3172, reg297158)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "length", value: __defun__length})

__defun__shen_4length_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3175 := __args[0]
_ = V3175
V3176 := __args[1]
_ = V3176
reg297160 := Nil;
reg297161 := PrimEqual(reg297160, V3175)
if reg297161 == True {
__ctx.Return(V3176)
return
} else {
reg297162 := PrimTail(V3175)
reg297163 := MakeNumber(1)
reg297164 := PrimNumberAdd(V3176, reg297163)
__ctx.TailApply(__defun__shen_4length_1h, reg297162, reg297164)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.length-h", value: __defun__shen_4length_1h})

__defun__occurrences = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3188 := __args[0]
_ = V3188
V3189 := __args[1]
_ = V3189
reg297166 := PrimEqual(V3189, V3188)
if reg297166 == True {
reg297167 := MakeNumber(1)
__ctx.Return(reg297167)
return
} else {
reg297168 := PrimIsPair(V3189)
if reg297168 == True {
reg297169 := PrimHead(V3189)
reg297170 := __e.Call(__defun__occurrences, V3188, reg297169)
reg297171 := PrimTail(V3189)
reg297172 := __e.Call(__defun__occurrences, V3188, reg297171)
reg297173 := PrimNumberAdd(reg297170, reg297172)
__ctx.Return(reg297173)
return
} else {
reg297174 := MakeNumber(0)
__ctx.Return(reg297174)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "occurrences", value: __defun__occurrences})

__defun__nth = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3196 := __args[0]
_ = V3196
V3197 := __args[1]
_ = V3197
reg297175 := MakeNumber(1)
reg297176 := PrimEqual(reg297175, V3196)
var reg297182 Obj
if reg297176 == True {
reg297177 := PrimIsPair(V3197)
var reg297180 Obj
if reg297177 == True {
reg297178 := True;
reg297180 = reg297178
} else {
reg297179 := False;
reg297180 = reg297179
}
reg297182 = reg297180
} else {
reg297181 := False;
reg297182 = reg297181
}
if reg297182 == True {
reg297183 := PrimHead(V3197)
__ctx.Return(reg297183)
return
} else {
reg297184 := PrimIsPair(V3197)
if reg297184 == True {
reg297185 := MakeNumber(1)
reg297186 := PrimNumberSubtract(V3196, reg297185)
reg297187 := PrimTail(V3197)
__ctx.TailApply(__defun__nth, reg297186, reg297187)
return
} else {
reg297189 := MakeString("nth applied to ")
reg297190 := MakeString(", ")
reg297191 := MakeString("\n")
reg297192 := MakeSymbol("shen.a")
reg297193 := __e.Call(__defun__shen_4app, V3197, reg297191, reg297192)
reg297194 := PrimStringConcat(reg297190, reg297193)
reg297195 := MakeSymbol("shen.a")
reg297196 := __e.Call(__defun__shen_4app, V3196, reg297194, reg297195)
reg297197 := PrimStringConcat(reg297189, reg297196)
reg297198 := PrimSimpleError(reg297197)
__ctx.Return(reg297198)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "nth", value: __defun__nth})

__defun__integer_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3199 := __args[0]
_ = V3199
reg297199 := PrimIsNumber(V3199)
if reg297199 == True {
reg297200 := __e.Call(__defun__shen_4abs, V3199)
Abs := reg297200
_ = Abs
reg297201 := MakeNumber(1)
reg297202 := __e.Call(__defun__shen_4magless, Abs, reg297201)
reg297203 := __e.Call(__defun__shen_4integer_1test_2, Abs, reg297202)
if reg297203 == True {
reg297204 := True;
__ctx.Return(reg297204)
return
} else {
reg297205 := False;
__ctx.Return(reg297205)
return
}
} else {
reg297206 := False;
__ctx.Return(reg297206)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "integer?", value: __defun__integer_2})

__defun__shen_4abs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3201 := __args[0]
_ = V3201
reg297207 := MakeNumber(0)
reg297208 := PrimGreatThan(V3201, reg297207)
if reg297208 == True {
__ctx.Return(V3201)
return
} else {
reg297209 := MakeNumber(0)
reg297210 := PrimNumberSubtract(reg297209, V3201)
__ctx.Return(reg297210)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.abs", value: __defun__shen_4abs})

__defun__shen_4magless = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3204 := __args[0]
_ = V3204
V3205 := __args[1]
_ = V3205
reg297211 := MakeNumber(2)
reg297212 := PrimNumberMultiply(V3205, reg297211)
Nx2 := reg297212
_ = Nx2
reg297213 := PrimGreatThan(Nx2, V3204)
if reg297213 == True {
__ctx.Return(V3205)
return
} else {
__ctx.TailApply(__defun__shen_4magless, V3204, Nx2)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.magless", value: __defun__shen_4magless})

__defun__shen_4integer_1test_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3211 := __args[0]
_ = V3211
V3212 := __args[1]
_ = V3212
reg297215 := MakeNumber(0)
reg297216 := PrimEqual(reg297215, V3211)
if reg297216 == True {
reg297217 := True;
__ctx.Return(reg297217)
return
} else {
reg297218 := MakeNumber(1)
reg297219 := PrimGreatThan(reg297218, V3211)
if reg297219 == True {
reg297220 := False;
__ctx.Return(reg297220)
return
} else {
reg297221 := PrimNumberSubtract(V3211, V3212)
Abs_1N := reg297221
_ = Abs_1N
reg297222 := MakeNumber(0)
reg297223 := PrimGreatThan(reg297222, Abs_1N)
if reg297223 == True {
reg297224 := PrimIsInteger(V3211)
__ctx.Return(reg297224)
return
} else {
__ctx.TailApply(__defun__shen_4integer_1test_2, Abs_1N, V3212)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.integer-test?", value: __defun__shen_4integer_1test_2})

__defun__mapcan = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3217 := __args[0]
_ = V3217
V3218 := __args[1]
_ = V3218
reg297226 := Nil;
reg297227 := PrimEqual(reg297226, V3218)
if reg297227 == True {
reg297228 := Nil;
__ctx.Return(reg297228)
return
} else {
reg297229 := PrimIsPair(V3218)
if reg297229 == True {
reg297230 := PrimHead(V3218)
reg297231 := __e.Call(V3217, reg297230)
reg297232 := PrimTail(V3218)
reg297233 := __e.Call(__defun__mapcan, V3217, reg297232)
__ctx.TailApply(__defun__append, reg297231, reg297233)
return
} else {
reg297235 := MakeSymbol("mapcan")
__ctx.TailApply(__defun__shen_4f__error, reg297235)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "mapcan", value: __defun__mapcan})

__defun___a_a = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3230 := __args[0]
_ = V3230
V3231 := __args[1]
_ = V3231
reg297237 := PrimEqual(V3231, V3230)
if reg297237 == True {
reg297238 := True;
__ctx.Return(reg297238)
return
} else {
reg297239 := False;
__ctx.Return(reg297239)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "==", value: __defun___a_a})

__defun__abort = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297240 := MakeString("")
reg297241 := PrimSimpleError(reg297240)
__ctx.Return(reg297241)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "abort", value: __defun__abort})

__defun__bound_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3233 := __args[0]
_ = V3233
reg297242 := PrimIsSymbol(V3233)
if reg297242 == True {
reg297243 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297244 := PrimValue(V3233)
__ctx.Return(reg297244)
return
}, 0)
reg297245 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg297246 := MakeSymbol("shen.this-symbol-is-unbound")
__ctx.Return(reg297246)
return
}, 1)
reg297247 := __e.Try(reg297243).Catch(reg297245)
Val := reg297247
_ = Val
reg297248 := MakeSymbol("shen.this-symbol-is-unbound")
reg297249 := PrimEqual(Val, reg297248)
var reg297252 Obj
if reg297249 == True {
reg297250 := False;
reg297252 = reg297250
} else {
reg297251 := True;
reg297252 = reg297251
}
if reg297252 == True {
reg297253 := True;
__ctx.Return(reg297253)
return
} else {
reg297254 := False;
__ctx.Return(reg297254)
return
}
} else {
reg297255 := False;
__ctx.Return(reg297255)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "bound?", value: __defun__bound_2})

__defun__shen_4string_1_6bytes = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3235 := __args[0]
_ = V3235
reg297256 := MakeString("")
reg297257 := PrimEqual(reg297256, V3235)
if reg297257 == True {
reg297258 := Nil;
__ctx.Return(reg297258)
return
} else {
reg297259 := MakeNumber(0)
reg297260 := PrimPos(V3235, reg297259)
reg297261 := PrimStringToNumber(reg297260)
reg297262 := PrimTailString(V3235)
reg297263 := __e.Call(__defun__shen_4string_1_6bytes, reg297262)
reg297264 := PrimCons(reg297261, reg297263)
__ctx.Return(reg297264)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.string->bytes", value: __defun__shen_4string_1_6bytes})

__defun__maxinferences = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3237 := __args[0]
_ = V3237
reg297265 := MakeSymbol("shen.*maxinferences*")
reg297266 := PrimSet(reg297265, V3237)
__ctx.Return(reg297266)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "maxinferences", value: __defun__maxinferences})

__defun__inferences = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297267 := MakeSymbol("shen.*infs*")
reg297268 := PrimValue(reg297267)
__ctx.Return(reg297268)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "inferences", value: __defun__inferences})

__defun__protect = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3239 := __args[0]
_ = V3239
__ctx.Return(V3239)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "protect", value: __defun__protect})

__defun__stoutput = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297269 := MakeSymbol("*stoutput*")
reg297270 := PrimValue(reg297269)
__ctx.Return(reg297270)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "stoutput", value: __defun__stoutput})

__defun__sterror = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297271 := MakeSymbol("*sterror*")
reg297272 := PrimValue(reg297271)
__ctx.Return(reg297272)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "sterror", value: __defun__sterror})

__defun__string_1_6symbol = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3241 := __args[0]
_ = V3241
reg297273 := PrimIntern(V3241)
Symbol := reg297273
_ = Symbol
reg297274 := PrimIsSymbol(Symbol)
if reg297274 == True {
__ctx.Return(Symbol)
return
} else {
reg297275 := MakeString("cannot intern ")
reg297276 := MakeString(" to a symbol")
reg297277 := MakeSymbol("shen.s")
reg297278 := __e.Call(__defun__shen_4app, V3241, reg297276, reg297277)
reg297279 := PrimStringConcat(reg297275, reg297278)
reg297280 := PrimSimpleError(reg297279)
__ctx.Return(reg297280)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "string->symbol", value: __defun__string_1_6symbol})

__defun__optimise = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3247 := __args[0]
_ = V3247
reg297281 := MakeSymbol("+")
reg297282 := PrimEqual(reg297281, V3247)
if reg297282 == True {
reg297283 := MakeSymbol("shen.*optimise*")
reg297284 := True;
reg297285 := PrimSet(reg297283, reg297284)
__ctx.Return(reg297285)
return
} else {
reg297286 := MakeSymbol("-")
reg297287 := PrimEqual(reg297286, V3247)
if reg297287 == True {
reg297288 := MakeSymbol("shen.*optimise*")
reg297289 := False;
reg297290 := PrimSet(reg297288, reg297289)
__ctx.Return(reg297290)
return
} else {
reg297291 := MakeString("optimise expects a + or a -.\n")
reg297292 := PrimSimpleError(reg297291)
__ctx.Return(reg297292)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "optimise", value: __defun__optimise})

__defun__os = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297293 := MakeSymbol("*os*")
reg297294 := PrimValue(reg297293)
__ctx.Return(reg297294)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "os", value: __defun__os})

__defun__language = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297295 := MakeSymbol("*language*")
reg297296 := PrimValue(reg297295)
__ctx.Return(reg297296)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "language", value: __defun__language})

__defun__version = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297297 := MakeSymbol("*version*")
reg297298 := PrimValue(reg297297)
__ctx.Return(reg297298)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "version", value: __defun__version})

__defun__port = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297299 := MakeSymbol("*port*")
reg297300 := PrimValue(reg297299)
__ctx.Return(reg297300)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "port", value: __defun__port})

__defun__porters = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297301 := MakeSymbol("*porters*")
reg297302 := PrimValue(reg297301)
__ctx.Return(reg297302)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "porters", value: __defun__porters})

__defun__implementation = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297303 := MakeSymbol("*implementation*")
reg297304 := PrimValue(reg297303)
__ctx.Return(reg297304)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "implementation", value: __defun__implementation})

__defun__release = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297305 := MakeSymbol("*release*")
reg297306 := PrimValue(reg297305)
__ctx.Return(reg297306)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "release", value: __defun__release})

__defun__package_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3249 := __args[0]
_ = V3249
reg297307 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297308 := __e.Call(__defun__external, V3249)
_ = reg297308
reg297309 := True;
__ctx.Return(reg297309)
return
}, 0)
reg297310 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg297311 := False;
__ctx.Return(reg297311)
return
}, 1)
reg297312 := __e.Try(reg297307).Catch(reg297310)
__ctx.Return(reg297312)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "package?", value: __defun__package_2})

__defun__function = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3251 := __args[0]
_ = V3251
__ctx.TailApply(__defun__shen_4lookup_1func, V3251)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "function", value: __defun__function})

__defun__shen_4lookup_1func = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3253 := __args[0]
_ = V3253
reg297314 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297315 := MakeSymbol("shen.lambda-form")
reg297316 := MakeSymbol("*property-vector*")
reg297317 := PrimValue(reg297316)
__ctx.TailApply(__defun__get, V3253, reg297315, reg297317)
return
}, 0)
reg297319 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg297320 := MakeString(" has no lambda expansion\n")
reg297321 := MakeSymbol("shen.a")
reg297322 := __e.Call(__defun__shen_4app, V3253, reg297320, reg297321)
reg297323 := PrimSimpleError(reg297322)
__ctx.Return(reg297323)
return
}, 1)
reg297324 := __e.Try(reg297314).Catch(reg297319)
__ctx.Return(reg297324)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.lookup-func", value: __defun__shen_4lookup_1func})

}
