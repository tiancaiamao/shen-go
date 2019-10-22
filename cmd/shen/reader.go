package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4read_1char_1code Obj // shen.read-char-code
var __defun__read_1file_1as_1bytelist Obj // read-file-as-bytelist
var __defun__shen_4read_1file_1as_1charlist Obj // shen.read-file-as-charlist
var __defun__shen_4read_1file_1as_1Xlist Obj // shen.read-file-as-Xlist
var __defun__shen_4read_1file_1as_1Xlist_1help Obj // shen.read-file-as-Xlist-help
var __defun__read_1file_1as_1string Obj // read-file-as-string
var __defun__shen_4rfas_1h Obj // shen.rfas-h
var __defun__input Obj // input
var __defun__input_7 Obj // input+
var __defun__shen_4monotype Obj // shen.monotype
var __defun__read Obj // read
var __defun__it Obj // it
var __defun__shen_4read_1loop Obj // shen.read-loop
var __defun__shen_4terminator_2 Obj // shen.terminator?
var __defun__lineread Obj // lineread
var __defun__shen_4lineread_1loop Obj // shen.lineread-loop
var __defun__shen_4record_1it Obj // shen.record-it
var __defun__shen_4trim_1whitespace Obj // shen.trim-whitespace
var __defun__shen_4record_1it_1h Obj // shen.record-it-h
var __defun__shen_4cn_1all Obj // shen.cn-all
var __defun__read_1file Obj // read-file
var __defun__read_1from_1string Obj // read-from-string
var __defun__shen_4read_1error Obj // shen.read-error
var __defun__shen_4compress_150 Obj // shen.compress-50
var __defun__shen_4_5st__input_6 Obj // shen.<st_input>
var __defun__shen_4_5lsb_6 Obj // shen.<lsb>
var __defun__shen_4_5rsb_6 Obj // shen.<rsb>
var __defun__shen_4_5lcurly_6 Obj // shen.<lcurly>
var __defun__shen_4_5rcurly_6 Obj // shen.<rcurly>
var __defun__shen_4_5bar_6 Obj // shen.<bar>
var __defun__shen_4_5semicolon_6 Obj // shen.<semicolon>
var __defun__shen_4_5colon_6 Obj // shen.<colon>
var __defun__shen_4_5comma_6 Obj // shen.<comma>
var __defun__shen_4_5equal_6 Obj // shen.<equal>
var __defun__shen_4_5minus_6 Obj // shen.<minus>
var __defun__shen_4_5lrb_6 Obj // shen.<lrb>
var __defun__shen_4_5rrb_6 Obj // shen.<rrb>
var __defun__shen_4_5atom_6 Obj // shen.<atom>
var __defun__shen_4control_1chars Obj // shen.control-chars
var __defun__shen_4code_1point Obj // shen.code-point
var __defun__shen_4after_1codepoint Obj // shen.after-codepoint
var __defun__shen_4decimalise Obj // shen.decimalise
var __defun__shen_4digits_1_6integers Obj // shen.digits->integers
var __defun__shen_4_5sym_6 Obj // shen.<sym>
var __defun__shen_4_5alphanums_6 Obj // shen.<alphanums>
var __defun__shen_4_5alphanum_6 Obj // shen.<alphanum>
var __defun__shen_4_5num_6 Obj // shen.<num>
var __defun__shen_4numbyte_2 Obj // shen.numbyte?
var __defun__shen_4_5alpha_6 Obj // shen.<alpha>
var __defun__shen_4symbol_1code_2 Obj // shen.symbol-code?
var __defun__shen_4_5str_6 Obj // shen.<str>
var __defun__shen_4_5dbq_6 Obj // shen.<dbq>
var __defun__shen_4_5strcontents_6 Obj // shen.<strcontents>
var __defun__shen_4_5byte_6 Obj // shen.<byte>
var __defun__shen_4_5strc_6 Obj // shen.<strc>
var __defun__shen_4_5number_6 Obj // shen.<number>
var __defun__shen_4_5E_6 Obj // shen.<E>
var __defun__shen_4_5log10_6 Obj // shen.<log10>
var __defun__shen_4_5plus_6 Obj // shen.<plus>
var __defun__shen_4_5stop_6 Obj // shen.<stop>
var __defun__shen_4_5predigits_6 Obj // shen.<predigits>
var __defun__shen_4_5postdigits_6 Obj // shen.<postdigits>
var __defun__shen_4_5digits_6 Obj // shen.<digits>
var __defun__shen_4_5digit_6 Obj // shen.<digit>
var __defun__shen_4byte_1_6digit Obj // shen.byte->digit
var __defun__shen_4pre Obj // shen.pre
var __defun__shen_4post Obj // shen.post
var __defun__shen_4expt Obj // shen.expt
var __defun__shen_4_5st__input1_6 Obj // shen.<st_input1>
var __defun__shen_4_5st__input2_6 Obj // shen.<st_input2>
var __defun__shen_4_5comment_6 Obj // shen.<comment>
var __defun__shen_4_5singleline_6 Obj // shen.<singleline>
var __defun__shen_4_5backslash_6 Obj // shen.<backslash>
var __defun__shen_4_5anysingle_6 Obj // shen.<anysingle>
var __defun__shen_4_5non_1return_6 Obj // shen.<non-return>
var __defun__shen_4_5return_6 Obj // shen.<return>
var __defun__shen_4_5multiline_6 Obj // shen.<multiline>
var __defun__shen_4_5times_6 Obj // shen.<times>
var __defun__shen_4_5anymulti_6 Obj // shen.<anymulti>
var __defun__shen_4_5whitespaces_6 Obj // shen.<whitespaces>
var __defun__shen_4_5whitespace_6 Obj // shen.<whitespace>
var __defun__shen_4cons__form Obj // shen.cons_form
var __defun__shen_4package_1macro Obj // shen.package-macro
var __defun__shen_4record_1exceptions Obj // shen.record-exceptions
var __defun__shen_4record_1internal Obj // shen.record-internal
var __defun__shen_4internal_1symbols Obj // shen.internal-symbols
var __defun__shen_4packageh Obj // shen.packageh

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg299875 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg299875)
return
}, 0))
__defun__shen_4read_1char_1code = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2358 := __args[0]
_ = V2358
reg299876 := PrimReadByte(V2358)
__ctx.Return(reg299876)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.read-char-code", value: __defun__shen_4read_1char_1code})

__defun__read_1file_1as_1bytelist = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2360 := __args[0]
_ = V2360
reg299877 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
S := __args[0]
_ = S
reg299878 := PrimReadByte(S)
__ctx.Return(reg299878)
return
}, 1)
__ctx.TailApply(__defun__shen_4read_1file_1as_1Xlist, V2360, reg299877)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "read-file-as-bytelist", value: __defun__read_1file_1as_1bytelist})

__defun__shen_4read_1file_1as_1charlist = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2362 := __args[0]
_ = V2362
reg299880 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
S := __args[0]
_ = S
__ctx.TailApply(__defun__shen_4read_1char_1code, S)
return
}, 1)
__ctx.TailApply(__defun__shen_4read_1file_1as_1Xlist, V2362, reg299880)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.read-file-as-charlist", value: __defun__shen_4read_1file_1as_1charlist})

__defun__shen_4read_1file_1as_1Xlist = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2365 := __args[0]
_ = V2365
V2366 := __args[1]
_ = V2366
reg299883 := MakeSymbol("in")
reg299884 := PrimOpenStream(V2365, reg299883)
Stream := reg299884
_ = Stream
reg299885 := __e.Call(V2366, Stream)
X := reg299885
_ = X
reg299886 := Nil;
reg299887 := __e.Call(__defun__shen_4read_1file_1as_1Xlist_1help, Stream, V2366, X, reg299886)
Xs := reg299887
_ = Xs
reg299888 := PrimCloseStream(Stream)
Close := reg299888
_ = Close
__ctx.TailApply(__defun__reverse, Xs)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.read-file-as-Xlist", value: __defun__shen_4read_1file_1as_1Xlist})

__defun__shen_4read_1file_1as_1Xlist_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2371 := __args[0]
_ = V2371
V2372 := __args[1]
_ = V2372
V2373 := __args[2]
_ = V2373
V2374 := __args[3]
_ = V2374
reg299890 := MakeNumber(-1)
reg299891 := PrimEqual(reg299890, V2373)
if reg299891 == True {
__ctx.Return(V2374)
return
} else {
reg299892 := __e.Call(V2372, V2371)
reg299893 := PrimCons(V2373, V2374)
__ctx.TailApply(__defun__shen_4read_1file_1as_1Xlist_1help, V2371, V2372, reg299892, reg299893)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.read-file-as-Xlist-help", value: __defun__shen_4read_1file_1as_1Xlist_1help})

__defun__read_1file_1as_1string = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2376 := __args[0]
_ = V2376
reg299895 := MakeSymbol("in")
reg299896 := PrimOpenStream(V2376, reg299895)
Stream := reg299896
_ = Stream
reg299897 := __e.Call(__defun__shen_4read_1char_1code, Stream)
reg299898 := MakeString("")
__ctx.TailApply(__defun__shen_4rfas_1h, Stream, reg299897, reg299898)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "read-file-as-string", value: __defun__read_1file_1as_1string})

__defun__shen_4rfas_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2380 := __args[0]
_ = V2380
V2381 := __args[1]
_ = V2381
V2382 := __args[2]
_ = V2382
reg299900 := MakeNumber(-1)
reg299901 := PrimEqual(reg299900, V2381)
if reg299901 == True {
reg299902 := PrimCloseStream(V2380)
_ = reg299902
__ctx.Return(V2382)
return
} else {
reg299903 := __e.Call(__defun__shen_4read_1char_1code, V2380)
reg299904 := PrimNumberToString(V2381)
reg299905 := PrimStringConcat(V2382, reg299904)
__ctx.TailApply(__defun__shen_4rfas_1h, V2380, reg299903, reg299905)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.rfas-h", value: __defun__shen_4rfas_1h})

__defun__input = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2384 := __args[0]
_ = V2384
reg299907 := __e.Call(__defun__read, V2384)
reg299908 := PrimEvalKL(__e, reg299907)
__ctx.Return(reg299908)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "input", value: __defun__input})

__defun__input_7 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2387 := __args[0]
_ = V2387
V2388 := __args[1]
_ = V2388
reg299909 := __e.Call(__defun__shen_4monotype, V2387)
Mono_2 := reg299909
_ = Mono_2
reg299910 := __e.Call(__defun__read, V2388)
Input := reg299910
_ = Input
reg299911 := False;
reg299912 := __e.Call(__defun__shen_4demodulate, V2387)
reg299913 := __e.Call(__defun__shen_4typecheck, Input, reg299912)
reg299914 := PrimEqual(reg299911, reg299913)
if reg299914 == True {
reg299915 := MakeString("type error: ")
reg299916 := MakeString(" is not of type ")
reg299917 := MakeString("\n")
reg299918 := MakeSymbol("shen.r")
reg299919 := __e.Call(__defun__shen_4app, V2387, reg299917, reg299918)
reg299920 := PrimStringConcat(reg299916, reg299919)
reg299921 := MakeSymbol("shen.r")
reg299922 := __e.Call(__defun__shen_4app, Input, reg299920, reg299921)
reg299923 := PrimStringConcat(reg299915, reg299922)
reg299924 := PrimSimpleError(reg299923)
__ctx.Return(reg299924)
return
} else {
reg299925 := PrimEvalKL(__e, Input)
__ctx.Return(reg299925)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "input+", value: __defun__input_7})

__defun__shen_4monotype = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2390 := __args[0]
_ = V2390
reg299926 := PrimIsPair(V2390)
if reg299926 == True {
reg299927 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4monotype, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg299927, V2390)
return
} else {
reg299930 := PrimIsVariable(V2390)
if reg299930 == True {
reg299931 := MakeString("input+ expects a monotype: not ")
reg299932 := MakeString("\n")
reg299933 := MakeSymbol("shen.a")
reg299934 := __e.Call(__defun__shen_4app, V2390, reg299932, reg299933)
reg299935 := PrimStringConcat(reg299931, reg299934)
reg299936 := PrimSimpleError(reg299935)
__ctx.Return(reg299936)
return
} else {
__ctx.Return(V2390)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.monotype", value: __defun__shen_4monotype})

__defun__read = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2392 := __args[0]
_ = V2392
reg299937 := __e.Call(__defun__shen_4read_1char_1code, V2392)
reg299938 := Nil;
reg299939 := __e.Call(__defun__shen_4read_1loop, V2392, reg299937, reg299938)
reg299940 := PrimHead(reg299939)
__ctx.Return(reg299940)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "read", value: __defun__read})

__defun__it = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg299941 := MakeSymbol("shen.*it*")
reg299942 := PrimValue(reg299941)
__ctx.Return(reg299942)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "it", value: __defun__it})

__defun__shen_4read_1loop = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2400 := __args[0]
_ = V2400
V2401 := __args[1]
_ = V2401
V2402 := __args[2]
_ = V2402
reg299943 := MakeNumber(94)
reg299944 := PrimEqual(reg299943, V2401)
if reg299944 == True {
reg299945 := MakeString("read aborted")
reg299946 := PrimSimpleError(reg299945)
__ctx.Return(reg299946)
return
} else {
reg299947 := MakeNumber(-1)
reg299948 := PrimEqual(reg299947, V2401)
if reg299948 == True {
reg299949 := __e.Call(__defun__empty_2, V2402)
if reg299949 == True {
reg299950 := MakeString("error: empty stream")
reg299951 := PrimSimpleError(reg299950)
__ctx.Return(reg299951)
return
} else {
reg299952 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg299954 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.Return(E)
return
}, 1)
__ctx.TailApply(__defun__compile, reg299952, V2402, reg299954)
return
}
} else {
reg299956 := __e.Call(__defun__shen_4terminator_2, V2401)
if reg299956 == True {
reg299957 := Nil;
reg299958 := PrimCons(V2401, reg299957)
reg299959 := __e.Call(__defun__append, V2402, reg299958)
AllChars := reg299959
_ = AllChars
reg299960 := __e.Call(__defun__shen_4record_1it, AllChars)
It := reg299960
_ = It
reg299961 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg299963 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg299964 := MakeSymbol("shen.nextbyte")
__ctx.Return(reg299964)
return
}, 1)
reg299965 := __e.Call(__defun__compile, reg299961, AllChars, reg299963)
Read := reg299965
_ = Read
reg299966 := MakeSymbol("shen.nextbyte")
reg299967 := PrimEqual(Read, reg299966)
var reg299973 Obj
if reg299967 == True {
reg299968 := True;
reg299973 = reg299968
} else {
reg299969 := __e.Call(__defun__empty_2, Read)
var reg299972 Obj
if reg299969 == True {
reg299970 := True;
reg299972 = reg299970
} else {
reg299971 := False;
reg299972 = reg299971
}
reg299973 = reg299972
}
if reg299973 == True {
reg299974 := __e.Call(__defun__shen_4read_1char_1code, V2400)
__ctx.TailApply(__defun__shen_4read_1loop, V2400, reg299974, AllChars)
return
} else {
__ctx.Return(Read)
return
}
} else {
reg299976 := __e.Call(__defun__shen_4read_1char_1code, V2400)
reg299977 := Nil;
reg299978 := PrimCons(V2401, reg299977)
reg299979 := __e.Call(__defun__append, V2402, reg299978)
__ctx.TailApply(__defun__shen_4read_1loop, V2400, reg299976, reg299979)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.read-loop", value: __defun__shen_4read_1loop})

__defun__shen_4terminator_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2404 := __args[0]
_ = V2404
reg299981 := MakeNumber(9)
reg299982 := MakeNumber(10)
reg299983 := MakeNumber(13)
reg299984 := MakeNumber(32)
reg299985 := MakeNumber(34)
reg299986 := MakeNumber(41)
reg299987 := MakeNumber(93)
reg299988 := Nil;
reg299989 := PrimCons(reg299987, reg299988)
reg299990 := PrimCons(reg299986, reg299989)
reg299991 := PrimCons(reg299985, reg299990)
reg299992 := PrimCons(reg299984, reg299991)
reg299993 := PrimCons(reg299983, reg299992)
reg299994 := PrimCons(reg299982, reg299993)
reg299995 := PrimCons(reg299981, reg299994)
__ctx.TailApply(__defun__element_2, V2404, reg299995)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.terminator?", value: __defun__shen_4terminator_2})

__defun__lineread = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2406 := __args[0]
_ = V2406
reg299997 := __e.Call(__defun__shen_4read_1char_1code, V2406)
reg299998 := Nil;
__ctx.TailApply(__defun__shen_4lineread_1loop, reg299997, reg299998, V2406)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "lineread", value: __defun__lineread})

__defun__shen_4lineread_1loop = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2411 := __args[0]
_ = V2411
V2412 := __args[1]
_ = V2412
V2413 := __args[2]
_ = V2413
reg300000 := MakeNumber(-1)
reg300001 := PrimEqual(reg300000, V2411)
if reg300001 == True {
reg300002 := __e.Call(__defun__empty_2, V2412)
if reg300002 == True {
reg300003 := MakeString("empty stream")
reg300004 := PrimSimpleError(reg300003)
__ctx.Return(reg300004)
return
} else {
reg300005 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg300007 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.Return(E)
return
}, 1)
__ctx.TailApply(__defun__compile, reg300005, V2412, reg300007)
return
}
} else {
reg300009 := __e.Call(__defun__shen_4hat)
reg300010 := PrimEqual(V2411, reg300009)
if reg300010 == True {
reg300011 := MakeString("line read aborted")
reg300012 := PrimSimpleError(reg300011)
__ctx.Return(reg300012)
return
} else {
reg300013 := __e.Call(__defun__shen_4newline)
reg300014 := __e.Call(__defun__shen_4carriage_1return)
reg300015 := Nil;
reg300016 := PrimCons(reg300014, reg300015)
reg300017 := PrimCons(reg300013, reg300016)
reg300018 := __e.Call(__defun__element_2, V2411, reg300017)
if reg300018 == True {
reg300019 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg300021 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg300022 := MakeSymbol("shen.nextline")
__ctx.Return(reg300022)
return
}, 1)
reg300023 := __e.Call(__defun__compile, reg300019, V2412, reg300021)
Line := reg300023
_ = Line
reg300024 := __e.Call(__defun__shen_4record_1it, V2412)
It := reg300024
_ = It
reg300025 := MakeSymbol("shen.nextline")
reg300026 := PrimEqual(Line, reg300025)
var reg300032 Obj
if reg300026 == True {
reg300027 := True;
reg300032 = reg300027
} else {
reg300028 := __e.Call(__defun__empty_2, Line)
var reg300031 Obj
if reg300028 == True {
reg300029 := True;
reg300031 = reg300029
} else {
reg300030 := False;
reg300031 = reg300030
}
reg300032 = reg300031
}
if reg300032 == True {
reg300033 := __e.Call(__defun__shen_4read_1char_1code, V2413)
reg300034 := Nil;
reg300035 := PrimCons(V2411, reg300034)
reg300036 := __e.Call(__defun__append, V2412, reg300035)
__ctx.TailApply(__defun__shen_4lineread_1loop, reg300033, reg300036, V2413)
return
} else {
__ctx.Return(Line)
return
}
} else {
reg300038 := __e.Call(__defun__shen_4read_1char_1code, V2413)
reg300039 := Nil;
reg300040 := PrimCons(V2411, reg300039)
reg300041 := __e.Call(__defun__append, V2412, reg300040)
__ctx.TailApply(__defun__shen_4lineread_1loop, reg300038, reg300041, V2413)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.lineread-loop", value: __defun__shen_4lineread_1loop})

__defun__shen_4record_1it = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2415 := __args[0]
_ = V2415
reg300043 := __e.Call(__defun__shen_4trim_1whitespace, V2415)
TrimLeft := reg300043
_ = TrimLeft
reg300044 := __e.Call(__defun__reverse, TrimLeft)
reg300045 := __e.Call(__defun__shen_4trim_1whitespace, reg300044)
TrimRight := reg300045
_ = TrimRight
reg300046 := __e.Call(__defun__reverse, TrimRight)
Trimmed := reg300046
_ = Trimmed
__ctx.TailApply(__defun__shen_4record_1it_1h, Trimmed)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.record-it", value: __defun__shen_4record_1it})

__defun__shen_4trim_1whitespace = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2417 := __args[0]
_ = V2417
reg300048 := PrimIsPair(V2417)
var reg300064 Obj
if reg300048 == True {
reg300049 := PrimHead(V2417)
reg300050 := MakeNumber(9)
reg300051 := MakeNumber(10)
reg300052 := MakeNumber(13)
reg300053 := MakeNumber(32)
reg300054 := Nil;
reg300055 := PrimCons(reg300053, reg300054)
reg300056 := PrimCons(reg300052, reg300055)
reg300057 := PrimCons(reg300051, reg300056)
reg300058 := PrimCons(reg300050, reg300057)
reg300059 := __e.Call(__defun__element_2, reg300049, reg300058)
var reg300062 Obj
if reg300059 == True {
reg300060 := True;
reg300062 = reg300060
} else {
reg300061 := False;
reg300062 = reg300061
}
reg300064 = reg300062
} else {
reg300063 := False;
reg300064 = reg300063
}
if reg300064 == True {
reg300065 := PrimTail(V2417)
__ctx.TailApply(__defun__shen_4trim_1whitespace, reg300065)
return
} else {
__ctx.Return(V2417)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.trim-whitespace", value: __defun__shen_4trim_1whitespace})

__defun__shen_4record_1it_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2419 := __args[0]
_ = V2419
reg300067 := MakeSymbol("shen.*it*")
reg300068 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg300069 := PrimNumberToString(X)
__ctx.Return(reg300069)
return
}, 1)
reg300070 := __e.Call(__defun__map, reg300068, V2419)
reg300071 := __e.Call(__defun__shen_4cn_1all, reg300070)
reg300072 := PrimSet(reg300067, reg300071)
_ = reg300072
__ctx.Return(V2419)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.record-it-h", value: __defun__shen_4record_1it_1h})

__defun__shen_4cn_1all = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2421 := __args[0]
_ = V2421
reg300073 := Nil;
reg300074 := PrimEqual(reg300073, V2421)
if reg300074 == True {
reg300075 := MakeString("")
__ctx.Return(reg300075)
return
} else {
reg300076 := PrimIsPair(V2421)
if reg300076 == True {
reg300077 := PrimHead(V2421)
reg300078 := PrimTail(V2421)
reg300079 := __e.Call(__defun__shen_4cn_1all, reg300078)
reg300080 := PrimStringConcat(reg300077, reg300079)
__ctx.Return(reg300080)
return
} else {
reg300081 := MakeSymbol("shen.cn-all")
__ctx.TailApply(__defun__shen_4f__error, reg300081)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.cn-all", value: __defun__shen_4cn_1all})

__defun__read_1file = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2423 := __args[0]
_ = V2423
reg300083 := __e.Call(__defun__shen_4read_1file_1as_1charlist, V2423)
Charlist := reg300083
_ = Charlist
reg300084 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg300086 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4read_1error, X)
return
}, 1)
__ctx.TailApply(__defun__compile, reg300084, Charlist, reg300086)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "read-file", value: __defun__read_1file})

__defun__read_1from_1string = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2425 := __args[0]
_ = V2425
reg300089 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg300090 := PrimStringToNumber(X)
__ctx.Return(reg300090)
return
}, 1)
reg300091 := __e.Call(__defun__explode, V2425)
reg300092 := __e.Call(__defun__map, reg300089, reg300091)
Ns := reg300092
_ = Ns
reg300093 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg300095 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4read_1error, X)
return
}, 1)
__ctx.TailApply(__defun__compile, reg300093, Ns, reg300095)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "read-from-string", value: __defun__read_1from_1string})

__defun__shen_4read_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2433 := __args[0]
_ = V2433
reg300098 := PrimIsPair(V2433)
var reg300121 Obj
if reg300098 == True {
reg300099 := PrimHead(V2433)
reg300100 := PrimIsPair(reg300099)
var reg300116 Obj
if reg300100 == True {
reg300101 := PrimTail(V2433)
reg300102 := PrimIsPair(reg300101)
var reg300111 Obj
if reg300102 == True {
reg300103 := Nil;
reg300104 := PrimTail(V2433)
reg300105 := PrimTail(reg300104)
reg300106 := PrimEqual(reg300103, reg300105)
var reg300109 Obj
if reg300106 == True {
reg300107 := True;
reg300109 = reg300107
} else {
reg300108 := False;
reg300109 = reg300108
}
reg300111 = reg300109
} else {
reg300110 := False;
reg300111 = reg300110
}
var reg300114 Obj
if reg300111 == True {
reg300112 := True;
reg300114 = reg300112
} else {
reg300113 := False;
reg300114 = reg300113
}
reg300116 = reg300114
} else {
reg300115 := False;
reg300116 = reg300115
}
var reg300119 Obj
if reg300116 == True {
reg300117 := True;
reg300119 = reg300117
} else {
reg300118 := False;
reg300119 = reg300118
}
reg300121 = reg300119
} else {
reg300120 := False;
reg300121 = reg300120
}
if reg300121 == True {
reg300122 := MakeString("read error here:\n\n ")
reg300123 := MakeNumber(50)
reg300124 := PrimHead(V2433)
reg300125 := __e.Call(__defun__shen_4compress_150, reg300123, reg300124)
reg300126 := MakeString("\n")
reg300127 := MakeSymbol("shen.a")
reg300128 := __e.Call(__defun__shen_4app, reg300125, reg300126, reg300127)
reg300129 := PrimStringConcat(reg300122, reg300128)
reg300130 := PrimSimpleError(reg300129)
__ctx.Return(reg300130)
return
} else {
reg300131 := MakeString("read error\n")
reg300132 := PrimSimpleError(reg300131)
__ctx.Return(reg300132)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.read-error", value: __defun__shen_4read_1error})

__defun__shen_4compress_150 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2440 := __args[0]
_ = V2440
V2441 := __args[1]
_ = V2441
reg300133 := Nil;
reg300134 := PrimEqual(reg300133, V2441)
if reg300134 == True {
reg300135 := MakeString("")
__ctx.Return(reg300135)
return
} else {
reg300136 := MakeNumber(0)
reg300137 := PrimEqual(reg300136, V2440)
if reg300137 == True {
reg300138 := MakeString("")
__ctx.Return(reg300138)
return
} else {
reg300139 := PrimIsPair(V2441)
if reg300139 == True {
reg300140 := PrimHead(V2441)
reg300141 := PrimNumberToString(reg300140)
reg300142 := MakeNumber(1)
reg300143 := PrimNumberSubtract(V2440, reg300142)
reg300144 := PrimTail(V2441)
reg300145 := __e.Call(__defun__shen_4compress_150, reg300143, reg300144)
reg300146 := PrimStringConcat(reg300141, reg300145)
__ctx.Return(reg300146)
return
} else {
reg300147 := MakeSymbol("shen.compress-50")
__ctx.TailApply(__defun__shen_4f__error, reg300147)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compress-50", value: __defun__shen_4compress_150})

__defun__shen_4_5st__input_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2443 := __args[0]
_ = V2443
reg300149 := __e.Call(__defun__shen_4_5lsb_6, V2443)
Parse__shen_4_5lsb_6 := reg300149
_ = Parse__shen_4_5lsb_6
reg300150 := __e.Call(__defun__fail)
reg300151 := PrimEqual(reg300150, Parse__shen_4_5lsb_6)
reg300152 := PrimNot(reg300151)
var reg300179 Obj
if reg300152 == True {
reg300153 := __e.Call(__defun__shen_4_5st__input1_6, Parse__shen_4_5lsb_6)
Parse__shen_4_5st__input1_6 := reg300153
_ = Parse__shen_4_5st__input1_6
reg300154 := __e.Call(__defun__fail)
reg300155 := PrimEqual(reg300154, Parse__shen_4_5st__input1_6)
reg300156 := PrimNot(reg300155)
var reg300177 Obj
if reg300156 == True {
reg300157 := __e.Call(__defun__shen_4_5rsb_6, Parse__shen_4_5st__input1_6)
Parse__shen_4_5rsb_6 := reg300157
_ = Parse__shen_4_5rsb_6
reg300158 := __e.Call(__defun__fail)
reg300159 := PrimEqual(reg300158, Parse__shen_4_5rsb_6)
reg300160 := PrimNot(reg300159)
var reg300175 Obj
if reg300160 == True {
reg300161 := __e.Call(__defun__shen_4_5st__input2_6, Parse__shen_4_5rsb_6)
Parse__shen_4_5st__input2_6 := reg300161
_ = Parse__shen_4_5st__input2_6
reg300162 := __e.Call(__defun__fail)
reg300163 := PrimEqual(reg300162, Parse__shen_4_5st__input2_6)
reg300164 := PrimNot(reg300163)
var reg300173 Obj
if reg300164 == True {
reg300165 := PrimHead(Parse__shen_4_5st__input2_6)
reg300166 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input1_6)
reg300167 := __e.Call(__defun__shen_4cons__form, reg300166)
reg300168 := __e.Call(__defun__macroexpand, reg300167)
reg300169 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input2_6)
reg300170 := PrimCons(reg300168, reg300169)
reg300171 := __e.Call(__defun__shen_4pair, reg300165, reg300170)
reg300173 = reg300171
} else {
reg300172 := __e.Call(__defun__fail)
reg300173 = reg300172
}
reg300175 = reg300173
} else {
reg300174 := __e.Call(__defun__fail)
reg300175 = reg300174
}
reg300177 = reg300175
} else {
reg300176 := __e.Call(__defun__fail)
reg300177 = reg300176
}
reg300179 = reg300177
} else {
reg300178 := __e.Call(__defun__fail)
reg300179 = reg300178
}
YaccParse := reg300179
_ = YaccParse
reg300180 := __e.Call(__defun__fail)
reg300181 := PrimEqual(YaccParse, reg300180)
if reg300181 == True {
reg300182 := __e.Call(__defun__shen_4_5lrb_6, V2443)
Parse__shen_4_5lrb_6 := reg300182
_ = Parse__shen_4_5lrb_6
reg300183 := __e.Call(__defun__fail)
reg300184 := PrimEqual(reg300183, Parse__shen_4_5lrb_6)
reg300185 := PrimNot(reg300184)
var reg300211 Obj
if reg300185 == True {
reg300186 := __e.Call(__defun__shen_4_5st__input1_6, Parse__shen_4_5lrb_6)
Parse__shen_4_5st__input1_6 := reg300186
_ = Parse__shen_4_5st__input1_6
reg300187 := __e.Call(__defun__fail)
reg300188 := PrimEqual(reg300187, Parse__shen_4_5st__input1_6)
reg300189 := PrimNot(reg300188)
var reg300209 Obj
if reg300189 == True {
reg300190 := __e.Call(__defun__shen_4_5rrb_6, Parse__shen_4_5st__input1_6)
Parse__shen_4_5rrb_6 := reg300190
_ = Parse__shen_4_5rrb_6
reg300191 := __e.Call(__defun__fail)
reg300192 := PrimEqual(reg300191, Parse__shen_4_5rrb_6)
reg300193 := PrimNot(reg300192)
var reg300207 Obj
if reg300193 == True {
reg300194 := __e.Call(__defun__shen_4_5st__input2_6, Parse__shen_4_5rrb_6)
Parse__shen_4_5st__input2_6 := reg300194
_ = Parse__shen_4_5st__input2_6
reg300195 := __e.Call(__defun__fail)
reg300196 := PrimEqual(reg300195, Parse__shen_4_5st__input2_6)
reg300197 := PrimNot(reg300196)
var reg300205 Obj
if reg300197 == True {
reg300198 := PrimHead(Parse__shen_4_5st__input2_6)
reg300199 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input1_6)
reg300200 := __e.Call(__defun__macroexpand, reg300199)
reg300201 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input2_6)
reg300202 := __e.Call(__defun__shen_4package_1macro, reg300200, reg300201)
reg300203 := __e.Call(__defun__shen_4pair, reg300198, reg300202)
reg300205 = reg300203
} else {
reg300204 := __e.Call(__defun__fail)
reg300205 = reg300204
}
reg300207 = reg300205
} else {
reg300206 := __e.Call(__defun__fail)
reg300207 = reg300206
}
reg300209 = reg300207
} else {
reg300208 := __e.Call(__defun__fail)
reg300209 = reg300208
}
reg300211 = reg300209
} else {
reg300210 := __e.Call(__defun__fail)
reg300211 = reg300210
}
YaccParse := reg300211
_ = YaccParse
reg300212 := __e.Call(__defun__fail)
reg300213 := PrimEqual(YaccParse, reg300212)
if reg300213 == True {
reg300214 := __e.Call(__defun__shen_4_5lcurly_6, V2443)
Parse__shen_4_5lcurly_6 := reg300214
_ = Parse__shen_4_5lcurly_6
reg300215 := __e.Call(__defun__fail)
reg300216 := PrimEqual(reg300215, Parse__shen_4_5lcurly_6)
reg300217 := PrimNot(reg300216)
var reg300230 Obj
if reg300217 == True {
reg300218 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5lcurly_6)
Parse__shen_4_5st__input_6 := reg300218
_ = Parse__shen_4_5st__input_6
reg300219 := __e.Call(__defun__fail)
reg300220 := PrimEqual(reg300219, Parse__shen_4_5st__input_6)
reg300221 := PrimNot(reg300220)
var reg300228 Obj
if reg300221 == True {
reg300222 := PrimHead(Parse__shen_4_5st__input_6)
reg300223 := MakeSymbol("{")
reg300224 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300225 := PrimCons(reg300223, reg300224)
reg300226 := __e.Call(__defun__shen_4pair, reg300222, reg300225)
reg300228 = reg300226
} else {
reg300227 := __e.Call(__defun__fail)
reg300228 = reg300227
}
reg300230 = reg300228
} else {
reg300229 := __e.Call(__defun__fail)
reg300230 = reg300229
}
YaccParse := reg300230
_ = YaccParse
reg300231 := __e.Call(__defun__fail)
reg300232 := PrimEqual(YaccParse, reg300231)
if reg300232 == True {
reg300233 := __e.Call(__defun__shen_4_5rcurly_6, V2443)
Parse__shen_4_5rcurly_6 := reg300233
_ = Parse__shen_4_5rcurly_6
reg300234 := __e.Call(__defun__fail)
reg300235 := PrimEqual(reg300234, Parse__shen_4_5rcurly_6)
reg300236 := PrimNot(reg300235)
var reg300249 Obj
if reg300236 == True {
reg300237 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5rcurly_6)
Parse__shen_4_5st__input_6 := reg300237
_ = Parse__shen_4_5st__input_6
reg300238 := __e.Call(__defun__fail)
reg300239 := PrimEqual(reg300238, Parse__shen_4_5st__input_6)
reg300240 := PrimNot(reg300239)
var reg300247 Obj
if reg300240 == True {
reg300241 := PrimHead(Parse__shen_4_5st__input_6)
reg300242 := MakeSymbol("}")
reg300243 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300244 := PrimCons(reg300242, reg300243)
reg300245 := __e.Call(__defun__shen_4pair, reg300241, reg300244)
reg300247 = reg300245
} else {
reg300246 := __e.Call(__defun__fail)
reg300247 = reg300246
}
reg300249 = reg300247
} else {
reg300248 := __e.Call(__defun__fail)
reg300249 = reg300248
}
YaccParse := reg300249
_ = YaccParse
reg300250 := __e.Call(__defun__fail)
reg300251 := PrimEqual(YaccParse, reg300250)
if reg300251 == True {
reg300252 := __e.Call(__defun__shen_4_5bar_6, V2443)
Parse__shen_4_5bar_6 := reg300252
_ = Parse__shen_4_5bar_6
reg300253 := __e.Call(__defun__fail)
reg300254 := PrimEqual(reg300253, Parse__shen_4_5bar_6)
reg300255 := PrimNot(reg300254)
var reg300268 Obj
if reg300255 == True {
reg300256 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5bar_6)
Parse__shen_4_5st__input_6 := reg300256
_ = Parse__shen_4_5st__input_6
reg300257 := __e.Call(__defun__fail)
reg300258 := PrimEqual(reg300257, Parse__shen_4_5st__input_6)
reg300259 := PrimNot(reg300258)
var reg300266 Obj
if reg300259 == True {
reg300260 := PrimHead(Parse__shen_4_5st__input_6)
reg300261 := MakeSymbol("bar!")
reg300262 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300263 := PrimCons(reg300261, reg300262)
reg300264 := __e.Call(__defun__shen_4pair, reg300260, reg300263)
reg300266 = reg300264
} else {
reg300265 := __e.Call(__defun__fail)
reg300266 = reg300265
}
reg300268 = reg300266
} else {
reg300267 := __e.Call(__defun__fail)
reg300268 = reg300267
}
YaccParse := reg300268
_ = YaccParse
reg300269 := __e.Call(__defun__fail)
reg300270 := PrimEqual(YaccParse, reg300269)
if reg300270 == True {
reg300271 := __e.Call(__defun__shen_4_5semicolon_6, V2443)
Parse__shen_4_5semicolon_6 := reg300271
_ = Parse__shen_4_5semicolon_6
reg300272 := __e.Call(__defun__fail)
reg300273 := PrimEqual(reg300272, Parse__shen_4_5semicolon_6)
reg300274 := PrimNot(reg300273)
var reg300287 Obj
if reg300274 == True {
reg300275 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5semicolon_6)
Parse__shen_4_5st__input_6 := reg300275
_ = Parse__shen_4_5st__input_6
reg300276 := __e.Call(__defun__fail)
reg300277 := PrimEqual(reg300276, Parse__shen_4_5st__input_6)
reg300278 := PrimNot(reg300277)
var reg300285 Obj
if reg300278 == True {
reg300279 := PrimHead(Parse__shen_4_5st__input_6)
reg300280 := MakeSymbol(";")
reg300281 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300282 := PrimCons(reg300280, reg300281)
reg300283 := __e.Call(__defun__shen_4pair, reg300279, reg300282)
reg300285 = reg300283
} else {
reg300284 := __e.Call(__defun__fail)
reg300285 = reg300284
}
reg300287 = reg300285
} else {
reg300286 := __e.Call(__defun__fail)
reg300287 = reg300286
}
YaccParse := reg300287
_ = YaccParse
reg300288 := __e.Call(__defun__fail)
reg300289 := PrimEqual(YaccParse, reg300288)
if reg300289 == True {
reg300290 := __e.Call(__defun__shen_4_5colon_6, V2443)
Parse__shen_4_5colon_6 := reg300290
_ = Parse__shen_4_5colon_6
reg300291 := __e.Call(__defun__fail)
reg300292 := PrimEqual(reg300291, Parse__shen_4_5colon_6)
reg300293 := PrimNot(reg300292)
var reg300312 Obj
if reg300293 == True {
reg300294 := __e.Call(__defun__shen_4_5equal_6, Parse__shen_4_5colon_6)
Parse__shen_4_5equal_6 := reg300294
_ = Parse__shen_4_5equal_6
reg300295 := __e.Call(__defun__fail)
reg300296 := PrimEqual(reg300295, Parse__shen_4_5equal_6)
reg300297 := PrimNot(reg300296)
var reg300310 Obj
if reg300297 == True {
reg300298 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5equal_6)
Parse__shen_4_5st__input_6 := reg300298
_ = Parse__shen_4_5st__input_6
reg300299 := __e.Call(__defun__fail)
reg300300 := PrimEqual(reg300299, Parse__shen_4_5st__input_6)
reg300301 := PrimNot(reg300300)
var reg300308 Obj
if reg300301 == True {
reg300302 := PrimHead(Parse__shen_4_5st__input_6)
reg300303 := MakeSymbol(":=")
reg300304 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300305 := PrimCons(reg300303, reg300304)
reg300306 := __e.Call(__defun__shen_4pair, reg300302, reg300305)
reg300308 = reg300306
} else {
reg300307 := __e.Call(__defun__fail)
reg300308 = reg300307
}
reg300310 = reg300308
} else {
reg300309 := __e.Call(__defun__fail)
reg300310 = reg300309
}
reg300312 = reg300310
} else {
reg300311 := __e.Call(__defun__fail)
reg300312 = reg300311
}
YaccParse := reg300312
_ = YaccParse
reg300313 := __e.Call(__defun__fail)
reg300314 := PrimEqual(YaccParse, reg300313)
if reg300314 == True {
reg300315 := __e.Call(__defun__shen_4_5colon_6, V2443)
Parse__shen_4_5colon_6 := reg300315
_ = Parse__shen_4_5colon_6
reg300316 := __e.Call(__defun__fail)
reg300317 := PrimEqual(reg300316, Parse__shen_4_5colon_6)
reg300318 := PrimNot(reg300317)
var reg300337 Obj
if reg300318 == True {
reg300319 := __e.Call(__defun__shen_4_5minus_6, Parse__shen_4_5colon_6)
Parse__shen_4_5minus_6 := reg300319
_ = Parse__shen_4_5minus_6
reg300320 := __e.Call(__defun__fail)
reg300321 := PrimEqual(reg300320, Parse__shen_4_5minus_6)
reg300322 := PrimNot(reg300321)
var reg300335 Obj
if reg300322 == True {
reg300323 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5minus_6)
Parse__shen_4_5st__input_6 := reg300323
_ = Parse__shen_4_5st__input_6
reg300324 := __e.Call(__defun__fail)
reg300325 := PrimEqual(reg300324, Parse__shen_4_5st__input_6)
reg300326 := PrimNot(reg300325)
var reg300333 Obj
if reg300326 == True {
reg300327 := PrimHead(Parse__shen_4_5st__input_6)
reg300328 := MakeSymbol(":-")
reg300329 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300330 := PrimCons(reg300328, reg300329)
reg300331 := __e.Call(__defun__shen_4pair, reg300327, reg300330)
reg300333 = reg300331
} else {
reg300332 := __e.Call(__defun__fail)
reg300333 = reg300332
}
reg300335 = reg300333
} else {
reg300334 := __e.Call(__defun__fail)
reg300335 = reg300334
}
reg300337 = reg300335
} else {
reg300336 := __e.Call(__defun__fail)
reg300337 = reg300336
}
YaccParse := reg300337
_ = YaccParse
reg300338 := __e.Call(__defun__fail)
reg300339 := PrimEqual(YaccParse, reg300338)
if reg300339 == True {
reg300340 := __e.Call(__defun__shen_4_5colon_6, V2443)
Parse__shen_4_5colon_6 := reg300340
_ = Parse__shen_4_5colon_6
reg300341 := __e.Call(__defun__fail)
reg300342 := PrimEqual(reg300341, Parse__shen_4_5colon_6)
reg300343 := PrimNot(reg300342)
var reg300356 Obj
if reg300343 == True {
reg300344 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5colon_6)
Parse__shen_4_5st__input_6 := reg300344
_ = Parse__shen_4_5st__input_6
reg300345 := __e.Call(__defun__fail)
reg300346 := PrimEqual(reg300345, Parse__shen_4_5st__input_6)
reg300347 := PrimNot(reg300346)
var reg300354 Obj
if reg300347 == True {
reg300348 := PrimHead(Parse__shen_4_5st__input_6)
reg300349 := MakeSymbol(":")
reg300350 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300351 := PrimCons(reg300349, reg300350)
reg300352 := __e.Call(__defun__shen_4pair, reg300348, reg300351)
reg300354 = reg300352
} else {
reg300353 := __e.Call(__defun__fail)
reg300354 = reg300353
}
reg300356 = reg300354
} else {
reg300355 := __e.Call(__defun__fail)
reg300356 = reg300355
}
YaccParse := reg300356
_ = YaccParse
reg300357 := __e.Call(__defun__fail)
reg300358 := PrimEqual(YaccParse, reg300357)
if reg300358 == True {
reg300359 := __e.Call(__defun__shen_4_5comma_6, V2443)
Parse__shen_4_5comma_6 := reg300359
_ = Parse__shen_4_5comma_6
reg300360 := __e.Call(__defun__fail)
reg300361 := PrimEqual(reg300360, Parse__shen_4_5comma_6)
reg300362 := PrimNot(reg300361)
var reg300376 Obj
if reg300362 == True {
reg300363 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5comma_6)
Parse__shen_4_5st__input_6 := reg300363
_ = Parse__shen_4_5st__input_6
reg300364 := __e.Call(__defun__fail)
reg300365 := PrimEqual(reg300364, Parse__shen_4_5st__input_6)
reg300366 := PrimNot(reg300365)
var reg300374 Obj
if reg300366 == True {
reg300367 := PrimHead(Parse__shen_4_5st__input_6)
reg300368 := MakeString(",")
reg300369 := PrimIntern(reg300368)
reg300370 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300371 := PrimCons(reg300369, reg300370)
reg300372 := __e.Call(__defun__shen_4pair, reg300367, reg300371)
reg300374 = reg300372
} else {
reg300373 := __e.Call(__defun__fail)
reg300374 = reg300373
}
reg300376 = reg300374
} else {
reg300375 := __e.Call(__defun__fail)
reg300376 = reg300375
}
YaccParse := reg300376
_ = YaccParse
reg300377 := __e.Call(__defun__fail)
reg300378 := PrimEqual(YaccParse, reg300377)
if reg300378 == True {
reg300379 := __e.Call(__defun__shen_4_5comment_6, V2443)
Parse__shen_4_5comment_6 := reg300379
_ = Parse__shen_4_5comment_6
reg300380 := __e.Call(__defun__fail)
reg300381 := PrimEqual(reg300380, Parse__shen_4_5comment_6)
reg300382 := PrimNot(reg300381)
var reg300393 Obj
if reg300382 == True {
reg300383 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5comment_6)
Parse__shen_4_5st__input_6 := reg300383
_ = Parse__shen_4_5st__input_6
reg300384 := __e.Call(__defun__fail)
reg300385 := PrimEqual(reg300384, Parse__shen_4_5st__input_6)
reg300386 := PrimNot(reg300385)
var reg300391 Obj
if reg300386 == True {
reg300387 := PrimHead(Parse__shen_4_5st__input_6)
reg300388 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300389 := __e.Call(__defun__shen_4pair, reg300387, reg300388)
reg300391 = reg300389
} else {
reg300390 := __e.Call(__defun__fail)
reg300391 = reg300390
}
reg300393 = reg300391
} else {
reg300392 := __e.Call(__defun__fail)
reg300393 = reg300392
}
YaccParse := reg300393
_ = YaccParse
reg300394 := __e.Call(__defun__fail)
reg300395 := PrimEqual(YaccParse, reg300394)
if reg300395 == True {
reg300396 := __e.Call(__defun__shen_4_5atom_6, V2443)
Parse__shen_4_5atom_6 := reg300396
_ = Parse__shen_4_5atom_6
reg300397 := __e.Call(__defun__fail)
reg300398 := PrimEqual(reg300397, Parse__shen_4_5atom_6)
reg300399 := PrimNot(reg300398)
var reg300413 Obj
if reg300399 == True {
reg300400 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5atom_6)
Parse__shen_4_5st__input_6 := reg300400
_ = Parse__shen_4_5st__input_6
reg300401 := __e.Call(__defun__fail)
reg300402 := PrimEqual(reg300401, Parse__shen_4_5st__input_6)
reg300403 := PrimNot(reg300402)
var reg300411 Obj
if reg300403 == True {
reg300404 := PrimHead(Parse__shen_4_5st__input_6)
reg300405 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5atom_6)
reg300406 := __e.Call(__defun__macroexpand, reg300405)
reg300407 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300408 := PrimCons(reg300406, reg300407)
reg300409 := __e.Call(__defun__shen_4pair, reg300404, reg300408)
reg300411 = reg300409
} else {
reg300410 := __e.Call(__defun__fail)
reg300411 = reg300410
}
reg300413 = reg300411
} else {
reg300412 := __e.Call(__defun__fail)
reg300413 = reg300412
}
YaccParse := reg300413
_ = YaccParse
reg300414 := __e.Call(__defun__fail)
reg300415 := PrimEqual(YaccParse, reg300414)
if reg300415 == True {
reg300416 := __e.Call(__defun__shen_4_5whitespaces_6, V2443)
Parse__shen_4_5whitespaces_6 := reg300416
_ = Parse__shen_4_5whitespaces_6
reg300417 := __e.Call(__defun__fail)
reg300418 := PrimEqual(reg300417, Parse__shen_4_5whitespaces_6)
reg300419 := PrimNot(reg300418)
var reg300430 Obj
if reg300419 == True {
reg300420 := __e.Call(__defun__shen_4_5st__input_6, Parse__shen_4_5whitespaces_6)
Parse__shen_4_5st__input_6 := reg300420
_ = Parse__shen_4_5st__input_6
reg300421 := __e.Call(__defun__fail)
reg300422 := PrimEqual(reg300421, Parse__shen_4_5st__input_6)
reg300423 := PrimNot(reg300422)
var reg300428 Obj
if reg300423 == True {
reg300424 := PrimHead(Parse__shen_4_5st__input_6)
reg300425 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
reg300426 := __e.Call(__defun__shen_4pair, reg300424, reg300425)
reg300428 = reg300426
} else {
reg300427 := __e.Call(__defun__fail)
reg300428 = reg300427
}
reg300430 = reg300428
} else {
reg300429 := __e.Call(__defun__fail)
reg300430 = reg300429
}
YaccParse := reg300430
_ = YaccParse
reg300431 := __e.Call(__defun__fail)
reg300432 := PrimEqual(YaccParse, reg300431)
if reg300432 == True {
reg300433 := __e.Call(__defun___5e_6, V2443)
Parse___5e_6 := reg300433
_ = Parse___5e_6
reg300434 := __e.Call(__defun__fail)
reg300435 := PrimEqual(reg300434, Parse___5e_6)
reg300436 := PrimNot(reg300435)
if reg300436 == True {
reg300437 := PrimHead(Parse___5e_6)
reg300438 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg300437, reg300438)
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
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<st_input>", value: __defun__shen_4_5st__input_6})

__defun__shen_4_5lsb_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2445 := __args[0]
_ = V2445
reg300441 := PrimHead(V2445)
reg300442 := PrimIsPair(reg300441)
var reg300451 Obj
if reg300442 == True {
reg300443 := MakeNumber(91)
reg300444 := PrimHead(V2445)
reg300445 := PrimHead(reg300444)
reg300446 := PrimEqual(reg300443, reg300445)
var reg300449 Obj
if reg300446 == True {
reg300447 := True;
reg300449 = reg300447
} else {
reg300448 := False;
reg300449 = reg300448
}
reg300451 = reg300449
} else {
reg300450 := False;
reg300451 = reg300450
}
if reg300451 == True {
reg300452 := PrimHead(V2445)
reg300453 := PrimTail(reg300452)
reg300454 := __e.Call(__defun__shen_4hdtl, V2445)
reg300455 := __e.Call(__defun__shen_4pair, reg300453, reg300454)
reg300456 := PrimHead(reg300455)
reg300457 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300456, reg300457)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<lsb>", value: __defun__shen_4_5lsb_6})

__defun__shen_4_5rsb_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2447 := __args[0]
_ = V2447
reg300460 := PrimHead(V2447)
reg300461 := PrimIsPair(reg300460)
var reg300470 Obj
if reg300461 == True {
reg300462 := MakeNumber(93)
reg300463 := PrimHead(V2447)
reg300464 := PrimHead(reg300463)
reg300465 := PrimEqual(reg300462, reg300464)
var reg300468 Obj
if reg300465 == True {
reg300466 := True;
reg300468 = reg300466
} else {
reg300467 := False;
reg300468 = reg300467
}
reg300470 = reg300468
} else {
reg300469 := False;
reg300470 = reg300469
}
if reg300470 == True {
reg300471 := PrimHead(V2447)
reg300472 := PrimTail(reg300471)
reg300473 := __e.Call(__defun__shen_4hdtl, V2447)
reg300474 := __e.Call(__defun__shen_4pair, reg300472, reg300473)
reg300475 := PrimHead(reg300474)
reg300476 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300475, reg300476)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<rsb>", value: __defun__shen_4_5rsb_6})

__defun__shen_4_5lcurly_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2449 := __args[0]
_ = V2449
reg300479 := PrimHead(V2449)
reg300480 := PrimIsPair(reg300479)
var reg300489 Obj
if reg300480 == True {
reg300481 := MakeNumber(123)
reg300482 := PrimHead(V2449)
reg300483 := PrimHead(reg300482)
reg300484 := PrimEqual(reg300481, reg300483)
var reg300487 Obj
if reg300484 == True {
reg300485 := True;
reg300487 = reg300485
} else {
reg300486 := False;
reg300487 = reg300486
}
reg300489 = reg300487
} else {
reg300488 := False;
reg300489 = reg300488
}
if reg300489 == True {
reg300490 := PrimHead(V2449)
reg300491 := PrimTail(reg300490)
reg300492 := __e.Call(__defun__shen_4hdtl, V2449)
reg300493 := __e.Call(__defun__shen_4pair, reg300491, reg300492)
reg300494 := PrimHead(reg300493)
reg300495 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300494, reg300495)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<lcurly>", value: __defun__shen_4_5lcurly_6})

__defun__shen_4_5rcurly_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2451 := __args[0]
_ = V2451
reg300498 := PrimHead(V2451)
reg300499 := PrimIsPair(reg300498)
var reg300508 Obj
if reg300499 == True {
reg300500 := MakeNumber(125)
reg300501 := PrimHead(V2451)
reg300502 := PrimHead(reg300501)
reg300503 := PrimEqual(reg300500, reg300502)
var reg300506 Obj
if reg300503 == True {
reg300504 := True;
reg300506 = reg300504
} else {
reg300505 := False;
reg300506 = reg300505
}
reg300508 = reg300506
} else {
reg300507 := False;
reg300508 = reg300507
}
if reg300508 == True {
reg300509 := PrimHead(V2451)
reg300510 := PrimTail(reg300509)
reg300511 := __e.Call(__defun__shen_4hdtl, V2451)
reg300512 := __e.Call(__defun__shen_4pair, reg300510, reg300511)
reg300513 := PrimHead(reg300512)
reg300514 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300513, reg300514)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<rcurly>", value: __defun__shen_4_5rcurly_6})

__defun__shen_4_5bar_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2453 := __args[0]
_ = V2453
reg300517 := PrimHead(V2453)
reg300518 := PrimIsPair(reg300517)
var reg300527 Obj
if reg300518 == True {
reg300519 := MakeNumber(124)
reg300520 := PrimHead(V2453)
reg300521 := PrimHead(reg300520)
reg300522 := PrimEqual(reg300519, reg300521)
var reg300525 Obj
if reg300522 == True {
reg300523 := True;
reg300525 = reg300523
} else {
reg300524 := False;
reg300525 = reg300524
}
reg300527 = reg300525
} else {
reg300526 := False;
reg300527 = reg300526
}
if reg300527 == True {
reg300528 := PrimHead(V2453)
reg300529 := PrimTail(reg300528)
reg300530 := __e.Call(__defun__shen_4hdtl, V2453)
reg300531 := __e.Call(__defun__shen_4pair, reg300529, reg300530)
reg300532 := PrimHead(reg300531)
reg300533 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300532, reg300533)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<bar>", value: __defun__shen_4_5bar_6})

__defun__shen_4_5semicolon_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2455 := __args[0]
_ = V2455
reg300536 := PrimHead(V2455)
reg300537 := PrimIsPair(reg300536)
var reg300546 Obj
if reg300537 == True {
reg300538 := MakeNumber(59)
reg300539 := PrimHead(V2455)
reg300540 := PrimHead(reg300539)
reg300541 := PrimEqual(reg300538, reg300540)
var reg300544 Obj
if reg300541 == True {
reg300542 := True;
reg300544 = reg300542
} else {
reg300543 := False;
reg300544 = reg300543
}
reg300546 = reg300544
} else {
reg300545 := False;
reg300546 = reg300545
}
if reg300546 == True {
reg300547 := PrimHead(V2455)
reg300548 := PrimTail(reg300547)
reg300549 := __e.Call(__defun__shen_4hdtl, V2455)
reg300550 := __e.Call(__defun__shen_4pair, reg300548, reg300549)
reg300551 := PrimHead(reg300550)
reg300552 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300551, reg300552)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<semicolon>", value: __defun__shen_4_5semicolon_6})

__defun__shen_4_5colon_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2457 := __args[0]
_ = V2457
reg300555 := PrimHead(V2457)
reg300556 := PrimIsPair(reg300555)
var reg300565 Obj
if reg300556 == True {
reg300557 := MakeNumber(58)
reg300558 := PrimHead(V2457)
reg300559 := PrimHead(reg300558)
reg300560 := PrimEqual(reg300557, reg300559)
var reg300563 Obj
if reg300560 == True {
reg300561 := True;
reg300563 = reg300561
} else {
reg300562 := False;
reg300563 = reg300562
}
reg300565 = reg300563
} else {
reg300564 := False;
reg300565 = reg300564
}
if reg300565 == True {
reg300566 := PrimHead(V2457)
reg300567 := PrimTail(reg300566)
reg300568 := __e.Call(__defun__shen_4hdtl, V2457)
reg300569 := __e.Call(__defun__shen_4pair, reg300567, reg300568)
reg300570 := PrimHead(reg300569)
reg300571 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300570, reg300571)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<colon>", value: __defun__shen_4_5colon_6})

__defun__shen_4_5comma_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2459 := __args[0]
_ = V2459
reg300574 := PrimHead(V2459)
reg300575 := PrimIsPair(reg300574)
var reg300584 Obj
if reg300575 == True {
reg300576 := MakeNumber(44)
reg300577 := PrimHead(V2459)
reg300578 := PrimHead(reg300577)
reg300579 := PrimEqual(reg300576, reg300578)
var reg300582 Obj
if reg300579 == True {
reg300580 := True;
reg300582 = reg300580
} else {
reg300581 := False;
reg300582 = reg300581
}
reg300584 = reg300582
} else {
reg300583 := False;
reg300584 = reg300583
}
if reg300584 == True {
reg300585 := PrimHead(V2459)
reg300586 := PrimTail(reg300585)
reg300587 := __e.Call(__defun__shen_4hdtl, V2459)
reg300588 := __e.Call(__defun__shen_4pair, reg300586, reg300587)
reg300589 := PrimHead(reg300588)
reg300590 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300589, reg300590)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<comma>", value: __defun__shen_4_5comma_6})

__defun__shen_4_5equal_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2461 := __args[0]
_ = V2461
reg300593 := PrimHead(V2461)
reg300594 := PrimIsPair(reg300593)
var reg300603 Obj
if reg300594 == True {
reg300595 := MakeNumber(61)
reg300596 := PrimHead(V2461)
reg300597 := PrimHead(reg300596)
reg300598 := PrimEqual(reg300595, reg300597)
var reg300601 Obj
if reg300598 == True {
reg300599 := True;
reg300601 = reg300599
} else {
reg300600 := False;
reg300601 = reg300600
}
reg300603 = reg300601
} else {
reg300602 := False;
reg300603 = reg300602
}
if reg300603 == True {
reg300604 := PrimHead(V2461)
reg300605 := PrimTail(reg300604)
reg300606 := __e.Call(__defun__shen_4hdtl, V2461)
reg300607 := __e.Call(__defun__shen_4pair, reg300605, reg300606)
reg300608 := PrimHead(reg300607)
reg300609 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300608, reg300609)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<equal>", value: __defun__shen_4_5equal_6})

__defun__shen_4_5minus_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2463 := __args[0]
_ = V2463
reg300612 := PrimHead(V2463)
reg300613 := PrimIsPair(reg300612)
var reg300622 Obj
if reg300613 == True {
reg300614 := MakeNumber(45)
reg300615 := PrimHead(V2463)
reg300616 := PrimHead(reg300615)
reg300617 := PrimEqual(reg300614, reg300616)
var reg300620 Obj
if reg300617 == True {
reg300618 := True;
reg300620 = reg300618
} else {
reg300619 := False;
reg300620 = reg300619
}
reg300622 = reg300620
} else {
reg300621 := False;
reg300622 = reg300621
}
if reg300622 == True {
reg300623 := PrimHead(V2463)
reg300624 := PrimTail(reg300623)
reg300625 := __e.Call(__defun__shen_4hdtl, V2463)
reg300626 := __e.Call(__defun__shen_4pair, reg300624, reg300625)
reg300627 := PrimHead(reg300626)
reg300628 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300627, reg300628)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<minus>", value: __defun__shen_4_5minus_6})

__defun__shen_4_5lrb_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2465 := __args[0]
_ = V2465
reg300631 := PrimHead(V2465)
reg300632 := PrimIsPair(reg300631)
var reg300641 Obj
if reg300632 == True {
reg300633 := MakeNumber(40)
reg300634 := PrimHead(V2465)
reg300635 := PrimHead(reg300634)
reg300636 := PrimEqual(reg300633, reg300635)
var reg300639 Obj
if reg300636 == True {
reg300637 := True;
reg300639 = reg300637
} else {
reg300638 := False;
reg300639 = reg300638
}
reg300641 = reg300639
} else {
reg300640 := False;
reg300641 = reg300640
}
if reg300641 == True {
reg300642 := PrimHead(V2465)
reg300643 := PrimTail(reg300642)
reg300644 := __e.Call(__defun__shen_4hdtl, V2465)
reg300645 := __e.Call(__defun__shen_4pair, reg300643, reg300644)
reg300646 := PrimHead(reg300645)
reg300647 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300646, reg300647)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<lrb>", value: __defun__shen_4_5lrb_6})

__defun__shen_4_5rrb_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2467 := __args[0]
_ = V2467
reg300650 := PrimHead(V2467)
reg300651 := PrimIsPair(reg300650)
var reg300660 Obj
if reg300651 == True {
reg300652 := MakeNumber(41)
reg300653 := PrimHead(V2467)
reg300654 := PrimHead(reg300653)
reg300655 := PrimEqual(reg300652, reg300654)
var reg300658 Obj
if reg300655 == True {
reg300656 := True;
reg300658 = reg300656
} else {
reg300657 := False;
reg300658 = reg300657
}
reg300660 = reg300658
} else {
reg300659 := False;
reg300660 = reg300659
}
if reg300660 == True {
reg300661 := PrimHead(V2467)
reg300662 := PrimTail(reg300661)
reg300663 := __e.Call(__defun__shen_4hdtl, V2467)
reg300664 := __e.Call(__defun__shen_4pair, reg300662, reg300663)
reg300665 := PrimHead(reg300664)
reg300666 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg300665, reg300666)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<rrb>", value: __defun__shen_4_5rrb_6})

__defun__shen_4_5atom_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2469 := __args[0]
_ = V2469
reg300669 := __e.Call(__defun__shen_4_5str_6, V2469)
Parse__shen_4_5str_6 := reg300669
_ = Parse__shen_4_5str_6
reg300670 := __e.Call(__defun__fail)
reg300671 := PrimEqual(reg300670, Parse__shen_4_5str_6)
reg300672 := PrimNot(reg300671)
var reg300678 Obj
if reg300672 == True {
reg300673 := PrimHead(Parse__shen_4_5str_6)
reg300674 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5str_6)
reg300675 := __e.Call(__defun__shen_4control_1chars, reg300674)
reg300676 := __e.Call(__defun__shen_4pair, reg300673, reg300675)
reg300678 = reg300676
} else {
reg300677 := __e.Call(__defun__fail)
reg300678 = reg300677
}
YaccParse := reg300678
_ = YaccParse
reg300679 := __e.Call(__defun__fail)
reg300680 := PrimEqual(YaccParse, reg300679)
if reg300680 == True {
reg300681 := __e.Call(__defun__shen_4_5number_6, V2469)
Parse__shen_4_5number_6 := reg300681
_ = Parse__shen_4_5number_6
reg300682 := __e.Call(__defun__fail)
reg300683 := PrimEqual(reg300682, Parse__shen_4_5number_6)
reg300684 := PrimNot(reg300683)
var reg300689 Obj
if reg300684 == True {
reg300685 := PrimHead(Parse__shen_4_5number_6)
reg300686 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5number_6)
reg300687 := __e.Call(__defun__shen_4pair, reg300685, reg300686)
reg300689 = reg300687
} else {
reg300688 := __e.Call(__defun__fail)
reg300689 = reg300688
}
YaccParse := reg300689
_ = YaccParse
reg300690 := __e.Call(__defun__fail)
reg300691 := PrimEqual(YaccParse, reg300690)
if reg300691 == True {
reg300692 := __e.Call(__defun__shen_4_5sym_6, V2469)
Parse__shen_4_5sym_6 := reg300692
_ = Parse__shen_4_5sym_6
reg300693 := __e.Call(__defun__fail)
reg300694 := PrimEqual(reg300693, Parse__shen_4_5sym_6)
reg300695 := PrimNot(reg300694)
if reg300695 == True {
reg300696 := PrimHead(Parse__shen_4_5sym_6)
reg300697 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5sym_6)
reg300698 := MakeString("<>")
reg300699 := PrimEqual(reg300697, reg300698)
var reg300707 Obj
if reg300699 == True {
reg300700 := MakeSymbol("vector")
reg300701 := MakeNumber(0)
reg300702 := Nil;
reg300703 := PrimCons(reg300701, reg300702)
reg300704 := PrimCons(reg300700, reg300703)
reg300707 = reg300704
} else {
reg300705 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5sym_6)
reg300706 := PrimIntern(reg300705)
reg300707 = reg300706
}
__ctx.TailApply(__defun__shen_4pair, reg300696, reg300707)
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
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<atom>", value: __defun__shen_4_5atom_6})

__defun__shen_4control_1chars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2471 := __args[0]
_ = V2471
reg300710 := Nil;
reg300711 := PrimEqual(reg300710, V2471)
if reg300711 == True {
reg300712 := MakeString("")
__ctx.Return(reg300712)
return
} else {
reg300713 := PrimIsPair(V2471)
var reg300737 Obj
if reg300713 == True {
reg300714 := MakeString("c")
reg300715 := PrimHead(V2471)
reg300716 := PrimEqual(reg300714, reg300715)
var reg300732 Obj
if reg300716 == True {
reg300717 := PrimTail(V2471)
reg300718 := PrimIsPair(reg300717)
var reg300727 Obj
if reg300718 == True {
reg300719 := MakeString("#")
reg300720 := PrimTail(V2471)
reg300721 := PrimHead(reg300720)
reg300722 := PrimEqual(reg300719, reg300721)
var reg300725 Obj
if reg300722 == True {
reg300723 := True;
reg300725 = reg300723
} else {
reg300724 := False;
reg300725 = reg300724
}
reg300727 = reg300725
} else {
reg300726 := False;
reg300727 = reg300726
}
var reg300730 Obj
if reg300727 == True {
reg300728 := True;
reg300730 = reg300728
} else {
reg300729 := False;
reg300730 = reg300729
}
reg300732 = reg300730
} else {
reg300731 := False;
reg300732 = reg300731
}
var reg300735 Obj
if reg300732 == True {
reg300733 := True;
reg300735 = reg300733
} else {
reg300734 := False;
reg300735 = reg300734
}
reg300737 = reg300735
} else {
reg300736 := False;
reg300737 = reg300736
}
if reg300737 == True {
reg300738 := PrimTail(V2471)
reg300739 := PrimTail(reg300738)
reg300740 := __e.Call(__defun__shen_4code_1point, reg300739)
CodePoint := reg300740
_ = CodePoint
reg300741 := PrimTail(V2471)
reg300742 := PrimTail(reg300741)
reg300743 := __e.Call(__defun__shen_4after_1codepoint, reg300742)
AfterCodePoint := reg300743
_ = AfterCodePoint
reg300744 := __e.Call(__defun__shen_4decimalise, CodePoint)
reg300745 := PrimNumberToString(reg300744)
reg300746 := __e.Call(__defun__shen_4control_1chars, AfterCodePoint)
__ctx.TailApply(__defun___8s, reg300745, reg300746)
return
} else {
reg300748 := PrimIsPair(V2471)
if reg300748 == True {
reg300749 := PrimHead(V2471)
reg300750 := PrimTail(V2471)
reg300751 := __e.Call(__defun__shen_4control_1chars, reg300750)
__ctx.TailApply(__defun___8s, reg300749, reg300751)
return
} else {
reg300753 := MakeSymbol("shen.control-chars")
__ctx.TailApply(__defun__shen_4f__error, reg300753)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.control-chars", value: __defun__shen_4control_1chars})

__defun__shen_4code_1point = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2475 := __args[0]
_ = V2475
reg300755 := PrimIsPair(V2475)
var reg300763 Obj
if reg300755 == True {
reg300756 := MakeString(";")
reg300757 := PrimHead(V2475)
reg300758 := PrimEqual(reg300756, reg300757)
var reg300761 Obj
if reg300758 == True {
reg300759 := True;
reg300761 = reg300759
} else {
reg300760 := False;
reg300761 = reg300760
}
reg300763 = reg300761
} else {
reg300762 := False;
reg300763 = reg300762
}
if reg300763 == True {
reg300764 := MakeString("")
__ctx.Return(reg300764)
return
} else {
reg300765 := PrimIsPair(V2475)
var reg300795 Obj
if reg300765 == True {
reg300766 := PrimHead(V2475)
reg300767 := MakeString("0")
reg300768 := MakeString("1")
reg300769 := MakeString("2")
reg300770 := MakeString("3")
reg300771 := MakeString("4")
reg300772 := MakeString("5")
reg300773 := MakeString("6")
reg300774 := MakeString("7")
reg300775 := MakeString("8")
reg300776 := MakeString("9")
reg300777 := MakeString("0")
reg300778 := Nil;
reg300779 := PrimCons(reg300777, reg300778)
reg300780 := PrimCons(reg300776, reg300779)
reg300781 := PrimCons(reg300775, reg300780)
reg300782 := PrimCons(reg300774, reg300781)
reg300783 := PrimCons(reg300773, reg300782)
reg300784 := PrimCons(reg300772, reg300783)
reg300785 := PrimCons(reg300771, reg300784)
reg300786 := PrimCons(reg300770, reg300785)
reg300787 := PrimCons(reg300769, reg300786)
reg300788 := PrimCons(reg300768, reg300787)
reg300789 := PrimCons(reg300767, reg300788)
reg300790 := __e.Call(__defun__element_2, reg300766, reg300789)
var reg300793 Obj
if reg300790 == True {
reg300791 := True;
reg300793 = reg300791
} else {
reg300792 := False;
reg300793 = reg300792
}
reg300795 = reg300793
} else {
reg300794 := False;
reg300795 = reg300794
}
if reg300795 == True {
reg300796 := PrimHead(V2475)
reg300797 := PrimTail(V2475)
reg300798 := __e.Call(__defun__shen_4code_1point, reg300797)
reg300799 := PrimCons(reg300796, reg300798)
__ctx.Return(reg300799)
return
} else {
reg300800 := MakeString("code point parse error ")
reg300801 := MakeString("\n")
reg300802 := MakeSymbol("shen.a")
reg300803 := __e.Call(__defun__shen_4app, V2475, reg300801, reg300802)
reg300804 := PrimStringConcat(reg300800, reg300803)
reg300805 := PrimSimpleError(reg300804)
__ctx.Return(reg300805)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.code-point", value: __defun__shen_4code_1point})

__defun__shen_4after_1codepoint = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2481 := __args[0]
_ = V2481
reg300806 := Nil;
reg300807 := PrimEqual(reg300806, V2481)
if reg300807 == True {
reg300808 := Nil;
__ctx.Return(reg300808)
return
} else {
reg300809 := PrimIsPair(V2481)
var reg300817 Obj
if reg300809 == True {
reg300810 := MakeString(";")
reg300811 := PrimHead(V2481)
reg300812 := PrimEqual(reg300810, reg300811)
var reg300815 Obj
if reg300812 == True {
reg300813 := True;
reg300815 = reg300813
} else {
reg300814 := False;
reg300815 = reg300814
}
reg300817 = reg300815
} else {
reg300816 := False;
reg300817 = reg300816
}
if reg300817 == True {
reg300818 := PrimTail(V2481)
__ctx.Return(reg300818)
return
} else {
reg300819 := PrimIsPair(V2481)
if reg300819 == True {
reg300820 := PrimTail(V2481)
__ctx.TailApply(__defun__shen_4after_1codepoint, reg300820)
return
} else {
reg300822 := MakeSymbol("shen.after-codepoint")
__ctx.TailApply(__defun__shen_4f__error, reg300822)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.after-codepoint", value: __defun__shen_4after_1codepoint})

__defun__shen_4decimalise = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2483 := __args[0]
_ = V2483
reg300824 := __e.Call(__defun__shen_4digits_1_6integers, V2483)
reg300825 := __e.Call(__defun__reverse, reg300824)
reg300826 := MakeNumber(0)
__ctx.TailApply(__defun__shen_4pre, reg300825, reg300826)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.decimalise", value: __defun__shen_4decimalise})

__defun__shen_4digits_1_6integers = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2489 := __args[0]
_ = V2489
reg300828 := PrimIsPair(V2489)
var reg300836 Obj
if reg300828 == True {
reg300829 := MakeString("0")
reg300830 := PrimHead(V2489)
reg300831 := PrimEqual(reg300829, reg300830)
var reg300834 Obj
if reg300831 == True {
reg300832 := True;
reg300834 = reg300832
} else {
reg300833 := False;
reg300834 = reg300833
}
reg300836 = reg300834
} else {
reg300835 := False;
reg300836 = reg300835
}
if reg300836 == True {
reg300837 := MakeNumber(0)
reg300838 := PrimTail(V2489)
reg300839 := __e.Call(__defun__shen_4digits_1_6integers, reg300838)
reg300840 := PrimCons(reg300837, reg300839)
__ctx.Return(reg300840)
return
} else {
reg300841 := PrimIsPair(V2489)
var reg300849 Obj
if reg300841 == True {
reg300842 := MakeString("1")
reg300843 := PrimHead(V2489)
reg300844 := PrimEqual(reg300842, reg300843)
var reg300847 Obj
if reg300844 == True {
reg300845 := True;
reg300847 = reg300845
} else {
reg300846 := False;
reg300847 = reg300846
}
reg300849 = reg300847
} else {
reg300848 := False;
reg300849 = reg300848
}
if reg300849 == True {
reg300850 := MakeNumber(1)
reg300851 := PrimTail(V2489)
reg300852 := __e.Call(__defun__shen_4digits_1_6integers, reg300851)
reg300853 := PrimCons(reg300850, reg300852)
__ctx.Return(reg300853)
return
} else {
reg300854 := PrimIsPair(V2489)
var reg300862 Obj
if reg300854 == True {
reg300855 := MakeString("2")
reg300856 := PrimHead(V2489)
reg300857 := PrimEqual(reg300855, reg300856)
var reg300860 Obj
if reg300857 == True {
reg300858 := True;
reg300860 = reg300858
} else {
reg300859 := False;
reg300860 = reg300859
}
reg300862 = reg300860
} else {
reg300861 := False;
reg300862 = reg300861
}
if reg300862 == True {
reg300863 := MakeNumber(2)
reg300864 := PrimTail(V2489)
reg300865 := __e.Call(__defun__shen_4digits_1_6integers, reg300864)
reg300866 := PrimCons(reg300863, reg300865)
__ctx.Return(reg300866)
return
} else {
reg300867 := PrimIsPair(V2489)
var reg300875 Obj
if reg300867 == True {
reg300868 := MakeString("3")
reg300869 := PrimHead(V2489)
reg300870 := PrimEqual(reg300868, reg300869)
var reg300873 Obj
if reg300870 == True {
reg300871 := True;
reg300873 = reg300871
} else {
reg300872 := False;
reg300873 = reg300872
}
reg300875 = reg300873
} else {
reg300874 := False;
reg300875 = reg300874
}
if reg300875 == True {
reg300876 := MakeNumber(3)
reg300877 := PrimTail(V2489)
reg300878 := __e.Call(__defun__shen_4digits_1_6integers, reg300877)
reg300879 := PrimCons(reg300876, reg300878)
__ctx.Return(reg300879)
return
} else {
reg300880 := PrimIsPair(V2489)
var reg300888 Obj
if reg300880 == True {
reg300881 := MakeString("4")
reg300882 := PrimHead(V2489)
reg300883 := PrimEqual(reg300881, reg300882)
var reg300886 Obj
if reg300883 == True {
reg300884 := True;
reg300886 = reg300884
} else {
reg300885 := False;
reg300886 = reg300885
}
reg300888 = reg300886
} else {
reg300887 := False;
reg300888 = reg300887
}
if reg300888 == True {
reg300889 := MakeNumber(4)
reg300890 := PrimTail(V2489)
reg300891 := __e.Call(__defun__shen_4digits_1_6integers, reg300890)
reg300892 := PrimCons(reg300889, reg300891)
__ctx.Return(reg300892)
return
} else {
reg300893 := PrimIsPair(V2489)
var reg300901 Obj
if reg300893 == True {
reg300894 := MakeString("5")
reg300895 := PrimHead(V2489)
reg300896 := PrimEqual(reg300894, reg300895)
var reg300899 Obj
if reg300896 == True {
reg300897 := True;
reg300899 = reg300897
} else {
reg300898 := False;
reg300899 = reg300898
}
reg300901 = reg300899
} else {
reg300900 := False;
reg300901 = reg300900
}
if reg300901 == True {
reg300902 := MakeNumber(5)
reg300903 := PrimTail(V2489)
reg300904 := __e.Call(__defun__shen_4digits_1_6integers, reg300903)
reg300905 := PrimCons(reg300902, reg300904)
__ctx.Return(reg300905)
return
} else {
reg300906 := PrimIsPair(V2489)
var reg300914 Obj
if reg300906 == True {
reg300907 := MakeString("6")
reg300908 := PrimHead(V2489)
reg300909 := PrimEqual(reg300907, reg300908)
var reg300912 Obj
if reg300909 == True {
reg300910 := True;
reg300912 = reg300910
} else {
reg300911 := False;
reg300912 = reg300911
}
reg300914 = reg300912
} else {
reg300913 := False;
reg300914 = reg300913
}
if reg300914 == True {
reg300915 := MakeNumber(6)
reg300916 := PrimTail(V2489)
reg300917 := __e.Call(__defun__shen_4digits_1_6integers, reg300916)
reg300918 := PrimCons(reg300915, reg300917)
__ctx.Return(reg300918)
return
} else {
reg300919 := PrimIsPair(V2489)
var reg300927 Obj
if reg300919 == True {
reg300920 := MakeString("7")
reg300921 := PrimHead(V2489)
reg300922 := PrimEqual(reg300920, reg300921)
var reg300925 Obj
if reg300922 == True {
reg300923 := True;
reg300925 = reg300923
} else {
reg300924 := False;
reg300925 = reg300924
}
reg300927 = reg300925
} else {
reg300926 := False;
reg300927 = reg300926
}
if reg300927 == True {
reg300928 := MakeNumber(7)
reg300929 := PrimTail(V2489)
reg300930 := __e.Call(__defun__shen_4digits_1_6integers, reg300929)
reg300931 := PrimCons(reg300928, reg300930)
__ctx.Return(reg300931)
return
} else {
reg300932 := PrimIsPair(V2489)
var reg300940 Obj
if reg300932 == True {
reg300933 := MakeString("8")
reg300934 := PrimHead(V2489)
reg300935 := PrimEqual(reg300933, reg300934)
var reg300938 Obj
if reg300935 == True {
reg300936 := True;
reg300938 = reg300936
} else {
reg300937 := False;
reg300938 = reg300937
}
reg300940 = reg300938
} else {
reg300939 := False;
reg300940 = reg300939
}
if reg300940 == True {
reg300941 := MakeNumber(8)
reg300942 := PrimTail(V2489)
reg300943 := __e.Call(__defun__shen_4digits_1_6integers, reg300942)
reg300944 := PrimCons(reg300941, reg300943)
__ctx.Return(reg300944)
return
} else {
reg300945 := PrimIsPair(V2489)
var reg300953 Obj
if reg300945 == True {
reg300946 := MakeString("9")
reg300947 := PrimHead(V2489)
reg300948 := PrimEqual(reg300946, reg300947)
var reg300951 Obj
if reg300948 == True {
reg300949 := True;
reg300951 = reg300949
} else {
reg300950 := False;
reg300951 = reg300950
}
reg300953 = reg300951
} else {
reg300952 := False;
reg300953 = reg300952
}
if reg300953 == True {
reg300954 := MakeNumber(9)
reg300955 := PrimTail(V2489)
reg300956 := __e.Call(__defun__shen_4digits_1_6integers, reg300955)
reg300957 := PrimCons(reg300954, reg300956)
__ctx.Return(reg300957)
return
} else {
reg300958 := Nil;
__ctx.Return(reg300958)
return
}
}
}
}
}
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.digits->integers", value: __defun__shen_4digits_1_6integers})

__defun__shen_4_5sym_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2491 := __args[0]
_ = V2491
reg300959 := __e.Call(__defun__shen_4_5alpha_6, V2491)
Parse__shen_4_5alpha_6 := reg300959
_ = Parse__shen_4_5alpha_6
reg300960 := __e.Call(__defun__fail)
reg300961 := PrimEqual(reg300960, Parse__shen_4_5alpha_6)
reg300962 := PrimNot(reg300961)
if reg300962 == True {
reg300963 := __e.Call(__defun__shen_4_5alphanums_6, Parse__shen_4_5alpha_6)
Parse__shen_4_5alphanums_6 := reg300963
_ = Parse__shen_4_5alphanums_6
reg300964 := __e.Call(__defun__fail)
reg300965 := PrimEqual(reg300964, Parse__shen_4_5alphanums_6)
reg300966 := PrimNot(reg300965)
if reg300966 == True {
reg300967 := PrimHead(Parse__shen_4_5alphanums_6)
reg300968 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5alpha_6)
reg300969 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5alphanums_6)
reg300970 := __e.Call(__defun___8s, reg300968, reg300969)
__ctx.TailApply(__defun__shen_4pair, reg300967, reg300970)
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
__initDefs = append(__initDefs, defType{name: "shen.<sym>", value: __defun__shen_4_5sym_6})

__defun__shen_4_5alphanums_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2493 := __args[0]
_ = V2493
reg300974 := __e.Call(__defun__shen_4_5alphanum_6, V2493)
Parse__shen_4_5alphanum_6 := reg300974
_ = Parse__shen_4_5alphanum_6
reg300975 := __e.Call(__defun__fail)
reg300976 := PrimEqual(reg300975, Parse__shen_4_5alphanum_6)
reg300977 := PrimNot(reg300976)
var reg300990 Obj
if reg300977 == True {
reg300978 := __e.Call(__defun__shen_4_5alphanums_6, Parse__shen_4_5alphanum_6)
Parse__shen_4_5alphanums_6 := reg300978
_ = Parse__shen_4_5alphanums_6
reg300979 := __e.Call(__defun__fail)
reg300980 := PrimEqual(reg300979, Parse__shen_4_5alphanums_6)
reg300981 := PrimNot(reg300980)
var reg300988 Obj
if reg300981 == True {
reg300982 := PrimHead(Parse__shen_4_5alphanums_6)
reg300983 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5alphanum_6)
reg300984 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5alphanums_6)
reg300985 := __e.Call(__defun___8s, reg300983, reg300984)
reg300986 := __e.Call(__defun__shen_4pair, reg300982, reg300985)
reg300988 = reg300986
} else {
reg300987 := __e.Call(__defun__fail)
reg300988 = reg300987
}
reg300990 = reg300988
} else {
reg300989 := __e.Call(__defun__fail)
reg300990 = reg300989
}
YaccParse := reg300990
_ = YaccParse
reg300991 := __e.Call(__defun__fail)
reg300992 := PrimEqual(YaccParse, reg300991)
if reg300992 == True {
reg300993 := __e.Call(__defun___5e_6, V2493)
Parse___5e_6 := reg300993
_ = Parse___5e_6
reg300994 := __e.Call(__defun__fail)
reg300995 := PrimEqual(reg300994, Parse___5e_6)
reg300996 := PrimNot(reg300995)
if reg300996 == True {
reg300997 := PrimHead(Parse___5e_6)
reg300998 := MakeString("")
__ctx.TailApply(__defun__shen_4pair, reg300997, reg300998)
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
__initDefs = append(__initDefs, defType{name: "shen.<alphanums>", value: __defun__shen_4_5alphanums_6})

__defun__shen_4_5alphanum_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2495 := __args[0]
_ = V2495
reg301001 := __e.Call(__defun__shen_4_5alpha_6, V2495)
Parse__shen_4_5alpha_6 := reg301001
_ = Parse__shen_4_5alpha_6
reg301002 := __e.Call(__defun__fail)
reg301003 := PrimEqual(reg301002, Parse__shen_4_5alpha_6)
reg301004 := PrimNot(reg301003)
var reg301009 Obj
if reg301004 == True {
reg301005 := PrimHead(Parse__shen_4_5alpha_6)
reg301006 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5alpha_6)
reg301007 := __e.Call(__defun__shen_4pair, reg301005, reg301006)
reg301009 = reg301007
} else {
reg301008 := __e.Call(__defun__fail)
reg301009 = reg301008
}
YaccParse := reg301009
_ = YaccParse
reg301010 := __e.Call(__defun__fail)
reg301011 := PrimEqual(YaccParse, reg301010)
if reg301011 == True {
reg301012 := __e.Call(__defun__shen_4_5num_6, V2495)
Parse__shen_4_5num_6 := reg301012
_ = Parse__shen_4_5num_6
reg301013 := __e.Call(__defun__fail)
reg301014 := PrimEqual(reg301013, Parse__shen_4_5num_6)
reg301015 := PrimNot(reg301014)
if reg301015 == True {
reg301016 := PrimHead(Parse__shen_4_5num_6)
reg301017 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5num_6)
__ctx.TailApply(__defun__shen_4pair, reg301016, reg301017)
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
__initDefs = append(__initDefs, defType{name: "shen.<alphanum>", value: __defun__shen_4_5alphanum_6})

__defun__shen_4_5num_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2497 := __args[0]
_ = V2497
reg301020 := PrimHead(V2497)
reg301021 := PrimIsPair(reg301020)
if reg301021 == True {
reg301022 := PrimHead(V2497)
reg301023 := PrimHead(reg301022)
Parse__Char := reg301023
_ = Parse__Char
reg301024 := __e.Call(__defun__shen_4numbyte_2, Parse__Char)
if reg301024 == True {
reg301025 := PrimHead(V2497)
reg301026 := PrimTail(reg301025)
reg301027 := __e.Call(__defun__shen_4hdtl, V2497)
reg301028 := __e.Call(__defun__shen_4pair, reg301026, reg301027)
reg301029 := PrimHead(reg301028)
reg301030 := PrimNumberToString(Parse__Char)
__ctx.TailApply(__defun__shen_4pair, reg301029, reg301030)
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
__initDefs = append(__initDefs, defType{name: "shen.<num>", value: __defun__shen_4_5num_6})

__defun__shen_4numbyte_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2503 := __args[0]
_ = V2503
reg301034 := MakeNumber(48)
reg301035 := PrimEqual(reg301034, V2503)
if reg301035 == True {
reg301036 := True;
__ctx.Return(reg301036)
return
} else {
reg301037 := MakeNumber(49)
reg301038 := PrimEqual(reg301037, V2503)
if reg301038 == True {
reg301039 := True;
__ctx.Return(reg301039)
return
} else {
reg301040 := MakeNumber(50)
reg301041 := PrimEqual(reg301040, V2503)
if reg301041 == True {
reg301042 := True;
__ctx.Return(reg301042)
return
} else {
reg301043 := MakeNumber(51)
reg301044 := PrimEqual(reg301043, V2503)
if reg301044 == True {
reg301045 := True;
__ctx.Return(reg301045)
return
} else {
reg301046 := MakeNumber(52)
reg301047 := PrimEqual(reg301046, V2503)
if reg301047 == True {
reg301048 := True;
__ctx.Return(reg301048)
return
} else {
reg301049 := MakeNumber(53)
reg301050 := PrimEqual(reg301049, V2503)
if reg301050 == True {
reg301051 := True;
__ctx.Return(reg301051)
return
} else {
reg301052 := MakeNumber(54)
reg301053 := PrimEqual(reg301052, V2503)
if reg301053 == True {
reg301054 := True;
__ctx.Return(reg301054)
return
} else {
reg301055 := MakeNumber(55)
reg301056 := PrimEqual(reg301055, V2503)
if reg301056 == True {
reg301057 := True;
__ctx.Return(reg301057)
return
} else {
reg301058 := MakeNumber(56)
reg301059 := PrimEqual(reg301058, V2503)
if reg301059 == True {
reg301060 := True;
__ctx.Return(reg301060)
return
} else {
reg301061 := MakeNumber(57)
reg301062 := PrimEqual(reg301061, V2503)
if reg301062 == True {
reg301063 := True;
__ctx.Return(reg301063)
return
} else {
reg301064 := False;
__ctx.Return(reg301064)
return
}
}
}
}
}
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.numbyte?", value: __defun__shen_4numbyte_2})

__defun__shen_4_5alpha_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2505 := __args[0]
_ = V2505
reg301065 := PrimHead(V2505)
reg301066 := PrimIsPair(reg301065)
if reg301066 == True {
reg301067 := PrimHead(V2505)
reg301068 := PrimHead(reg301067)
Parse__Char := reg301068
_ = Parse__Char
reg301069 := __e.Call(__defun__shen_4symbol_1code_2, Parse__Char)
if reg301069 == True {
reg301070 := PrimHead(V2505)
reg301071 := PrimTail(reg301070)
reg301072 := __e.Call(__defun__shen_4hdtl, V2505)
reg301073 := __e.Call(__defun__shen_4pair, reg301071, reg301072)
reg301074 := PrimHead(reg301073)
reg301075 := PrimNumberToString(Parse__Char)
__ctx.TailApply(__defun__shen_4pair, reg301074, reg301075)
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
__initDefs = append(__initDefs, defType{name: "shen.<alpha>", value: __defun__shen_4_5alpha_6})

__defun__shen_4symbol_1code_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2507 := __args[0]
_ = V2507
reg301079 := MakeNumber(126)
reg301080 := PrimEqual(V2507, reg301079)
if reg301080 == True {
reg301081 := True;
__ctx.Return(reg301081)
return
} else {
reg301082 := MakeNumber(94)
reg301083 := PrimGreatThan(V2507, reg301082)
var reg301090 Obj
if reg301083 == True {
reg301084 := MakeNumber(123)
reg301085 := PrimLessThan(V2507, reg301084)
var reg301088 Obj
if reg301085 == True {
reg301086 := True;
reg301088 = reg301086
} else {
reg301087 := False;
reg301088 = reg301087
}
reg301090 = reg301088
} else {
reg301089 := False;
reg301090 = reg301089
}
var reg301147 Obj
if reg301090 == True {
reg301091 := True;
reg301147 = reg301091
} else {
reg301092 := MakeNumber(59)
reg301093 := PrimGreatThan(V2507, reg301092)
var reg301100 Obj
if reg301093 == True {
reg301094 := MakeNumber(91)
reg301095 := PrimLessThan(V2507, reg301094)
var reg301098 Obj
if reg301095 == True {
reg301096 := True;
reg301098 = reg301096
} else {
reg301097 := False;
reg301098 = reg301097
}
reg301100 = reg301098
} else {
reg301099 := False;
reg301100 = reg301099
}
var reg301143 Obj
if reg301100 == True {
reg301101 := True;
reg301143 = reg301101
} else {
reg301102 := MakeNumber(41)
reg301103 := PrimGreatThan(V2507, reg301102)
var reg301118 Obj
if reg301103 == True {
reg301104 := MakeNumber(58)
reg301105 := PrimLessThan(V2507, reg301104)
var reg301113 Obj
if reg301105 == True {
reg301106 := MakeNumber(44)
reg301107 := PrimEqual(V2507, reg301106)
reg301108 := PrimNot(reg301107)
var reg301111 Obj
if reg301108 == True {
reg301109 := True;
reg301111 = reg301109
} else {
reg301110 := False;
reg301111 = reg301110
}
reg301113 = reg301111
} else {
reg301112 := False;
reg301113 = reg301112
}
var reg301116 Obj
if reg301113 == True {
reg301114 := True;
reg301116 = reg301114
} else {
reg301115 := False;
reg301116 = reg301115
}
reg301118 = reg301116
} else {
reg301117 := False;
reg301118 = reg301117
}
var reg301139 Obj
if reg301118 == True {
reg301119 := True;
reg301139 = reg301119
} else {
reg301120 := MakeNumber(34)
reg301121 := PrimGreatThan(V2507, reg301120)
var reg301128 Obj
if reg301121 == True {
reg301122 := MakeNumber(40)
reg301123 := PrimLessThan(V2507, reg301122)
var reg301126 Obj
if reg301123 == True {
reg301124 := True;
reg301126 = reg301124
} else {
reg301125 := False;
reg301126 = reg301125
}
reg301128 = reg301126
} else {
reg301127 := False;
reg301128 = reg301127
}
var reg301135 Obj
if reg301128 == True {
reg301129 := True;
reg301135 = reg301129
} else {
reg301130 := MakeNumber(33)
reg301131 := PrimEqual(V2507, reg301130)
var reg301134 Obj
if reg301131 == True {
reg301132 := True;
reg301134 = reg301132
} else {
reg301133 := False;
reg301134 = reg301133
}
reg301135 = reg301134
}
var reg301138 Obj
if reg301135 == True {
reg301136 := True;
reg301138 = reg301136
} else {
reg301137 := False;
reg301138 = reg301137
}
reg301139 = reg301138
}
var reg301142 Obj
if reg301139 == True {
reg301140 := True;
reg301142 = reg301140
} else {
reg301141 := False;
reg301142 = reg301141
}
reg301143 = reg301142
}
var reg301146 Obj
if reg301143 == True {
reg301144 := True;
reg301146 = reg301144
} else {
reg301145 := False;
reg301146 = reg301145
}
reg301147 = reg301146
}
if reg301147 == True {
reg301148 := True;
__ctx.Return(reg301148)
return
} else {
reg301149 := False;
__ctx.Return(reg301149)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.symbol-code?", value: __defun__shen_4symbol_1code_2})

__defun__shen_4_5str_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2509 := __args[0]
_ = V2509
reg301150 := __e.Call(__defun__shen_4_5dbq_6, V2509)
Parse__shen_4_5dbq_6 := reg301150
_ = Parse__shen_4_5dbq_6
reg301151 := __e.Call(__defun__fail)
reg301152 := PrimEqual(reg301151, Parse__shen_4_5dbq_6)
reg301153 := PrimNot(reg301152)
if reg301153 == True {
reg301154 := __e.Call(__defun__shen_4_5strcontents_6, Parse__shen_4_5dbq_6)
Parse__shen_4_5strcontents_6 := reg301154
_ = Parse__shen_4_5strcontents_6
reg301155 := __e.Call(__defun__fail)
reg301156 := PrimEqual(reg301155, Parse__shen_4_5strcontents_6)
reg301157 := PrimNot(reg301156)
if reg301157 == True {
reg301158 := __e.Call(__defun__shen_4_5dbq_6, Parse__shen_4_5strcontents_6)
Parse__shen_4_5dbq_6 := reg301158
_ = Parse__shen_4_5dbq_6
reg301159 := __e.Call(__defun__fail)
reg301160 := PrimEqual(reg301159, Parse__shen_4_5dbq_6)
reg301161 := PrimNot(reg301160)
if reg301161 == True {
reg301162 := PrimHead(Parse__shen_4_5dbq_6)
reg301163 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5strcontents_6)
__ctx.TailApply(__defun__shen_4pair, reg301162, reg301163)
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
__initDefs = append(__initDefs, defType{name: "shen.<str>", value: __defun__shen_4_5str_6})

__defun__shen_4_5dbq_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2511 := __args[0]
_ = V2511
reg301168 := PrimHead(V2511)
reg301169 := PrimIsPair(reg301168)
if reg301169 == True {
reg301170 := PrimHead(V2511)
reg301171 := PrimHead(reg301170)
Parse__Char := reg301171
_ = Parse__Char
reg301172 := MakeNumber(34)
reg301173 := PrimEqual(Parse__Char, reg301172)
if reg301173 == True {
reg301174 := PrimHead(V2511)
reg301175 := PrimTail(reg301174)
reg301176 := __e.Call(__defun__shen_4hdtl, V2511)
reg301177 := __e.Call(__defun__shen_4pair, reg301175, reg301176)
reg301178 := PrimHead(reg301177)
__ctx.TailApply(__defun__shen_4pair, reg301178, Parse__Char)
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
__initDefs = append(__initDefs, defType{name: "shen.<dbq>", value: __defun__shen_4_5dbq_6})

__defun__shen_4_5strcontents_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2513 := __args[0]
_ = V2513
reg301182 := __e.Call(__defun__shen_4_5strc_6, V2513)
Parse__shen_4_5strc_6 := reg301182
_ = Parse__shen_4_5strc_6
reg301183 := __e.Call(__defun__fail)
reg301184 := PrimEqual(reg301183, Parse__shen_4_5strc_6)
reg301185 := PrimNot(reg301184)
var reg301198 Obj
if reg301185 == True {
reg301186 := __e.Call(__defun__shen_4_5strcontents_6, Parse__shen_4_5strc_6)
Parse__shen_4_5strcontents_6 := reg301186
_ = Parse__shen_4_5strcontents_6
reg301187 := __e.Call(__defun__fail)
reg301188 := PrimEqual(reg301187, Parse__shen_4_5strcontents_6)
reg301189 := PrimNot(reg301188)
var reg301196 Obj
if reg301189 == True {
reg301190 := PrimHead(Parse__shen_4_5strcontents_6)
reg301191 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5strc_6)
reg301192 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5strcontents_6)
reg301193 := PrimCons(reg301191, reg301192)
reg301194 := __e.Call(__defun__shen_4pair, reg301190, reg301193)
reg301196 = reg301194
} else {
reg301195 := __e.Call(__defun__fail)
reg301196 = reg301195
}
reg301198 = reg301196
} else {
reg301197 := __e.Call(__defun__fail)
reg301198 = reg301197
}
YaccParse := reg301198
_ = YaccParse
reg301199 := __e.Call(__defun__fail)
reg301200 := PrimEqual(YaccParse, reg301199)
if reg301200 == True {
reg301201 := __e.Call(__defun___5e_6, V2513)
Parse___5e_6 := reg301201
_ = Parse___5e_6
reg301202 := __e.Call(__defun__fail)
reg301203 := PrimEqual(reg301202, Parse___5e_6)
reg301204 := PrimNot(reg301203)
if reg301204 == True {
reg301205 := PrimHead(Parse___5e_6)
reg301206 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg301205, reg301206)
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
__initDefs = append(__initDefs, defType{name: "shen.<strcontents>", value: __defun__shen_4_5strcontents_6})

__defun__shen_4_5byte_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2515 := __args[0]
_ = V2515
reg301209 := PrimHead(V2515)
reg301210 := PrimIsPair(reg301209)
if reg301210 == True {
reg301211 := PrimHead(V2515)
reg301212 := PrimHead(reg301211)
Parse__Char := reg301212
_ = Parse__Char
reg301213 := PrimHead(V2515)
reg301214 := PrimTail(reg301213)
reg301215 := __e.Call(__defun__shen_4hdtl, V2515)
reg301216 := __e.Call(__defun__shen_4pair, reg301214, reg301215)
reg301217 := PrimHead(reg301216)
reg301218 := PrimNumberToString(Parse__Char)
__ctx.TailApply(__defun__shen_4pair, reg301217, reg301218)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<byte>", value: __defun__shen_4_5byte_6})

__defun__shen_4_5strc_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2517 := __args[0]
_ = V2517
reg301221 := PrimHead(V2517)
reg301222 := PrimIsPair(reg301221)
if reg301222 == True {
reg301223 := PrimHead(V2517)
reg301224 := PrimHead(reg301223)
Parse__Char := reg301224
_ = Parse__Char
reg301225 := MakeNumber(34)
reg301226 := PrimEqual(Parse__Char, reg301225)
reg301227 := PrimNot(reg301226)
if reg301227 == True {
reg301228 := PrimHead(V2517)
reg301229 := PrimTail(reg301228)
reg301230 := __e.Call(__defun__shen_4hdtl, V2517)
reg301231 := __e.Call(__defun__shen_4pair, reg301229, reg301230)
reg301232 := PrimHead(reg301231)
reg301233 := PrimNumberToString(Parse__Char)
__ctx.TailApply(__defun__shen_4pair, reg301232, reg301233)
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
__initDefs = append(__initDefs, defType{name: "shen.<strc>", value: __defun__shen_4_5strc_6})

__defun__shen_4_5number_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2519 := __args[0]
_ = V2519
reg301237 := __e.Call(__defun__shen_4_5minus_6, V2519)
Parse__shen_4_5minus_6 := reg301237
_ = Parse__shen_4_5minus_6
reg301238 := __e.Call(__defun__fail)
reg301239 := PrimEqual(reg301238, Parse__shen_4_5minus_6)
reg301240 := PrimNot(reg301239)
var reg301253 Obj
if reg301240 == True {
reg301241 := __e.Call(__defun__shen_4_5number_6, Parse__shen_4_5minus_6)
Parse__shen_4_5number_6 := reg301241
_ = Parse__shen_4_5number_6
reg301242 := __e.Call(__defun__fail)
reg301243 := PrimEqual(reg301242, Parse__shen_4_5number_6)
reg301244 := PrimNot(reg301243)
var reg301251 Obj
if reg301244 == True {
reg301245 := PrimHead(Parse__shen_4_5number_6)
reg301246 := MakeNumber(0)
reg301247 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5number_6)
reg301248 := PrimNumberSubtract(reg301246, reg301247)
reg301249 := __e.Call(__defun__shen_4pair, reg301245, reg301248)
reg301251 = reg301249
} else {
reg301250 := __e.Call(__defun__fail)
reg301251 = reg301250
}
reg301253 = reg301251
} else {
reg301252 := __e.Call(__defun__fail)
reg301253 = reg301252
}
YaccParse := reg301253
_ = YaccParse
reg301254 := __e.Call(__defun__fail)
reg301255 := PrimEqual(YaccParse, reg301254)
if reg301255 == True {
reg301256 := __e.Call(__defun__shen_4_5plus_6, V2519)
Parse__shen_4_5plus_6 := reg301256
_ = Parse__shen_4_5plus_6
reg301257 := __e.Call(__defun__fail)
reg301258 := PrimEqual(reg301257, Parse__shen_4_5plus_6)
reg301259 := PrimNot(reg301258)
var reg301270 Obj
if reg301259 == True {
reg301260 := __e.Call(__defun__shen_4_5number_6, Parse__shen_4_5plus_6)
Parse__shen_4_5number_6 := reg301260
_ = Parse__shen_4_5number_6
reg301261 := __e.Call(__defun__fail)
reg301262 := PrimEqual(reg301261, Parse__shen_4_5number_6)
reg301263 := PrimNot(reg301262)
var reg301268 Obj
if reg301263 == True {
reg301264 := PrimHead(Parse__shen_4_5number_6)
reg301265 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5number_6)
reg301266 := __e.Call(__defun__shen_4pair, reg301264, reg301265)
reg301268 = reg301266
} else {
reg301267 := __e.Call(__defun__fail)
reg301268 = reg301267
}
reg301270 = reg301268
} else {
reg301269 := __e.Call(__defun__fail)
reg301270 = reg301269
}
YaccParse := reg301270
_ = YaccParse
reg301271 := __e.Call(__defun__fail)
reg301272 := PrimEqual(YaccParse, reg301271)
if reg301272 == True {
reg301273 := __e.Call(__defun__shen_4_5predigits_6, V2519)
Parse__shen_4_5predigits_6 := reg301273
_ = Parse__shen_4_5predigits_6
reg301274 := __e.Call(__defun__fail)
reg301275 := PrimEqual(reg301274, Parse__shen_4_5predigits_6)
reg301276 := PrimNot(reg301275)
var reg301316 Obj
if reg301276 == True {
reg301277 := __e.Call(__defun__shen_4_5stop_6, Parse__shen_4_5predigits_6)
Parse__shen_4_5stop_6 := reg301277
_ = Parse__shen_4_5stop_6
reg301278 := __e.Call(__defun__fail)
reg301279 := PrimEqual(reg301278, Parse__shen_4_5stop_6)
reg301280 := PrimNot(reg301279)
var reg301314 Obj
if reg301280 == True {
reg301281 := __e.Call(__defun__shen_4_5postdigits_6, Parse__shen_4_5stop_6)
Parse__shen_4_5postdigits_6 := reg301281
_ = Parse__shen_4_5postdigits_6
reg301282 := __e.Call(__defun__fail)
reg301283 := PrimEqual(reg301282, Parse__shen_4_5postdigits_6)
reg301284 := PrimNot(reg301283)
var reg301312 Obj
if reg301284 == True {
reg301285 := __e.Call(__defun__shen_4_5E_6, Parse__shen_4_5postdigits_6)
Parse__shen_4_5E_6 := reg301285
_ = Parse__shen_4_5E_6
reg301286 := __e.Call(__defun__fail)
reg301287 := PrimEqual(reg301286, Parse__shen_4_5E_6)
reg301288 := PrimNot(reg301287)
var reg301310 Obj
if reg301288 == True {
reg301289 := __e.Call(__defun__shen_4_5log10_6, Parse__shen_4_5E_6)
Parse__shen_4_5log10_6 := reg301289
_ = Parse__shen_4_5log10_6
reg301290 := __e.Call(__defun__fail)
reg301291 := PrimEqual(reg301290, Parse__shen_4_5log10_6)
reg301292 := PrimNot(reg301291)
var reg301308 Obj
if reg301292 == True {
reg301293 := PrimHead(Parse__shen_4_5log10_6)
reg301294 := MakeNumber(10)
reg301295 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5log10_6)
reg301296 := __e.Call(__defun__shen_4expt, reg301294, reg301295)
reg301297 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5predigits_6)
reg301298 := __e.Call(__defun__reverse, reg301297)
reg301299 := MakeNumber(0)
reg301300 := __e.Call(__defun__shen_4pre, reg301298, reg301299)
reg301301 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5postdigits_6)
reg301302 := MakeNumber(1)
reg301303 := __e.Call(__defun__shen_4post, reg301301, reg301302)
reg301304 := PrimNumberAdd(reg301300, reg301303)
reg301305 := PrimNumberMultiply(reg301296, reg301304)
reg301306 := __e.Call(__defun__shen_4pair, reg301293, reg301305)
reg301308 = reg301306
} else {
reg301307 := __e.Call(__defun__fail)
reg301308 = reg301307
}
reg301310 = reg301308
} else {
reg301309 := __e.Call(__defun__fail)
reg301310 = reg301309
}
reg301312 = reg301310
} else {
reg301311 := __e.Call(__defun__fail)
reg301312 = reg301311
}
reg301314 = reg301312
} else {
reg301313 := __e.Call(__defun__fail)
reg301314 = reg301313
}
reg301316 = reg301314
} else {
reg301315 := __e.Call(__defun__fail)
reg301316 = reg301315
}
YaccParse := reg301316
_ = YaccParse
reg301317 := __e.Call(__defun__fail)
reg301318 := PrimEqual(YaccParse, reg301317)
if reg301318 == True {
reg301319 := __e.Call(__defun__shen_4_5digits_6, V2519)
Parse__shen_4_5digits_6 := reg301319
_ = Parse__shen_4_5digits_6
reg301320 := __e.Call(__defun__fail)
reg301321 := PrimEqual(reg301320, Parse__shen_4_5digits_6)
reg301322 := PrimNot(reg301321)
var reg301346 Obj
if reg301322 == True {
reg301323 := __e.Call(__defun__shen_4_5E_6, Parse__shen_4_5digits_6)
Parse__shen_4_5E_6 := reg301323
_ = Parse__shen_4_5E_6
reg301324 := __e.Call(__defun__fail)
reg301325 := PrimEqual(reg301324, Parse__shen_4_5E_6)
reg301326 := PrimNot(reg301325)
var reg301344 Obj
if reg301326 == True {
reg301327 := __e.Call(__defun__shen_4_5log10_6, Parse__shen_4_5E_6)
Parse__shen_4_5log10_6 := reg301327
_ = Parse__shen_4_5log10_6
reg301328 := __e.Call(__defun__fail)
reg301329 := PrimEqual(reg301328, Parse__shen_4_5log10_6)
reg301330 := PrimNot(reg301329)
var reg301342 Obj
if reg301330 == True {
reg301331 := PrimHead(Parse__shen_4_5log10_6)
reg301332 := MakeNumber(10)
reg301333 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5log10_6)
reg301334 := __e.Call(__defun__shen_4expt, reg301332, reg301333)
reg301335 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
reg301336 := __e.Call(__defun__reverse, reg301335)
reg301337 := MakeNumber(0)
reg301338 := __e.Call(__defun__shen_4pre, reg301336, reg301337)
reg301339 := PrimNumberMultiply(reg301334, reg301338)
reg301340 := __e.Call(__defun__shen_4pair, reg301331, reg301339)
reg301342 = reg301340
} else {
reg301341 := __e.Call(__defun__fail)
reg301342 = reg301341
}
reg301344 = reg301342
} else {
reg301343 := __e.Call(__defun__fail)
reg301344 = reg301343
}
reg301346 = reg301344
} else {
reg301345 := __e.Call(__defun__fail)
reg301346 = reg301345
}
YaccParse := reg301346
_ = YaccParse
reg301347 := __e.Call(__defun__fail)
reg301348 := PrimEqual(YaccParse, reg301347)
if reg301348 == True {
reg301349 := __e.Call(__defun__shen_4_5predigits_6, V2519)
Parse__shen_4_5predigits_6 := reg301349
_ = Parse__shen_4_5predigits_6
reg301350 := __e.Call(__defun__fail)
reg301351 := PrimEqual(reg301350, Parse__shen_4_5predigits_6)
reg301352 := PrimNot(reg301351)
var reg301376 Obj
if reg301352 == True {
reg301353 := __e.Call(__defun__shen_4_5stop_6, Parse__shen_4_5predigits_6)
Parse__shen_4_5stop_6 := reg301353
_ = Parse__shen_4_5stop_6
reg301354 := __e.Call(__defun__fail)
reg301355 := PrimEqual(reg301354, Parse__shen_4_5stop_6)
reg301356 := PrimNot(reg301355)
var reg301374 Obj
if reg301356 == True {
reg301357 := __e.Call(__defun__shen_4_5postdigits_6, Parse__shen_4_5stop_6)
Parse__shen_4_5postdigits_6 := reg301357
_ = Parse__shen_4_5postdigits_6
reg301358 := __e.Call(__defun__fail)
reg301359 := PrimEqual(reg301358, Parse__shen_4_5postdigits_6)
reg301360 := PrimNot(reg301359)
var reg301372 Obj
if reg301360 == True {
reg301361 := PrimHead(Parse__shen_4_5postdigits_6)
reg301362 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5predigits_6)
reg301363 := __e.Call(__defun__reverse, reg301362)
reg301364 := MakeNumber(0)
reg301365 := __e.Call(__defun__shen_4pre, reg301363, reg301364)
reg301366 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5postdigits_6)
reg301367 := MakeNumber(1)
reg301368 := __e.Call(__defun__shen_4post, reg301366, reg301367)
reg301369 := PrimNumberAdd(reg301365, reg301368)
reg301370 := __e.Call(__defun__shen_4pair, reg301361, reg301369)
reg301372 = reg301370
} else {
reg301371 := __e.Call(__defun__fail)
reg301372 = reg301371
}
reg301374 = reg301372
} else {
reg301373 := __e.Call(__defun__fail)
reg301374 = reg301373
}
reg301376 = reg301374
} else {
reg301375 := __e.Call(__defun__fail)
reg301376 = reg301375
}
YaccParse := reg301376
_ = YaccParse
reg301377 := __e.Call(__defun__fail)
reg301378 := PrimEqual(YaccParse, reg301377)
if reg301378 == True {
reg301379 := __e.Call(__defun__shen_4_5digits_6, V2519)
Parse__shen_4_5digits_6 := reg301379
_ = Parse__shen_4_5digits_6
reg301380 := __e.Call(__defun__fail)
reg301381 := PrimEqual(reg301380, Parse__shen_4_5digits_6)
reg301382 := PrimNot(reg301381)
if reg301382 == True {
reg301383 := PrimHead(Parse__shen_4_5digits_6)
reg301384 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
reg301385 := __e.Call(__defun__reverse, reg301384)
reg301386 := MakeNumber(0)
reg301387 := __e.Call(__defun__shen_4pre, reg301385, reg301386)
__ctx.TailApply(__defun__shen_4pair, reg301383, reg301387)
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
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<number>", value: __defun__shen_4_5number_6})

__defun__shen_4_5E_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2521 := __args[0]
_ = V2521
reg301390 := PrimHead(V2521)
reg301391 := PrimIsPair(reg301390)
var reg301400 Obj
if reg301391 == True {
reg301392 := MakeNumber(101)
reg301393 := PrimHead(V2521)
reg301394 := PrimHead(reg301393)
reg301395 := PrimEqual(reg301392, reg301394)
var reg301398 Obj
if reg301395 == True {
reg301396 := True;
reg301398 = reg301396
} else {
reg301397 := False;
reg301398 = reg301397
}
reg301400 = reg301398
} else {
reg301399 := False;
reg301400 = reg301399
}
if reg301400 == True {
reg301401 := PrimHead(V2521)
reg301402 := PrimTail(reg301401)
reg301403 := __e.Call(__defun__shen_4hdtl, V2521)
reg301404 := __e.Call(__defun__shen_4pair, reg301402, reg301403)
reg301405 := PrimHead(reg301404)
reg301406 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301405, reg301406)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<E>", value: __defun__shen_4_5E_6})

__defun__shen_4_5log10_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2523 := __args[0]
_ = V2523
reg301409 := __e.Call(__defun__shen_4_5minus_6, V2523)
Parse__shen_4_5minus_6 := reg301409
_ = Parse__shen_4_5minus_6
reg301410 := __e.Call(__defun__fail)
reg301411 := PrimEqual(reg301410, Parse__shen_4_5minus_6)
reg301412 := PrimNot(reg301411)
var reg301428 Obj
if reg301412 == True {
reg301413 := __e.Call(__defun__shen_4_5digits_6, Parse__shen_4_5minus_6)
Parse__shen_4_5digits_6 := reg301413
_ = Parse__shen_4_5digits_6
reg301414 := __e.Call(__defun__fail)
reg301415 := PrimEqual(reg301414, Parse__shen_4_5digits_6)
reg301416 := PrimNot(reg301415)
var reg301426 Obj
if reg301416 == True {
reg301417 := PrimHead(Parse__shen_4_5digits_6)
reg301418 := MakeNumber(0)
reg301419 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
reg301420 := __e.Call(__defun__reverse, reg301419)
reg301421 := MakeNumber(0)
reg301422 := __e.Call(__defun__shen_4pre, reg301420, reg301421)
reg301423 := PrimNumberSubtract(reg301418, reg301422)
reg301424 := __e.Call(__defun__shen_4pair, reg301417, reg301423)
reg301426 = reg301424
} else {
reg301425 := __e.Call(__defun__fail)
reg301426 = reg301425
}
reg301428 = reg301426
} else {
reg301427 := __e.Call(__defun__fail)
reg301428 = reg301427
}
YaccParse := reg301428
_ = YaccParse
reg301429 := __e.Call(__defun__fail)
reg301430 := PrimEqual(YaccParse, reg301429)
if reg301430 == True {
reg301431 := __e.Call(__defun__shen_4_5digits_6, V2523)
Parse__shen_4_5digits_6 := reg301431
_ = Parse__shen_4_5digits_6
reg301432 := __e.Call(__defun__fail)
reg301433 := PrimEqual(reg301432, Parse__shen_4_5digits_6)
reg301434 := PrimNot(reg301433)
if reg301434 == True {
reg301435 := PrimHead(Parse__shen_4_5digits_6)
reg301436 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
reg301437 := __e.Call(__defun__reverse, reg301436)
reg301438 := MakeNumber(0)
reg301439 := __e.Call(__defun__shen_4pre, reg301437, reg301438)
__ctx.TailApply(__defun__shen_4pair, reg301435, reg301439)
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
__initDefs = append(__initDefs, defType{name: "shen.<log10>", value: __defun__shen_4_5log10_6})

__defun__shen_4_5plus_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2525 := __args[0]
_ = V2525
reg301442 := PrimHead(V2525)
reg301443 := PrimIsPair(reg301442)
if reg301443 == True {
reg301444 := PrimHead(V2525)
reg301445 := PrimHead(reg301444)
Parse__Char := reg301445
_ = Parse__Char
reg301446 := MakeNumber(43)
reg301447 := PrimEqual(Parse__Char, reg301446)
if reg301447 == True {
reg301448 := PrimHead(V2525)
reg301449 := PrimTail(reg301448)
reg301450 := __e.Call(__defun__shen_4hdtl, V2525)
reg301451 := __e.Call(__defun__shen_4pair, reg301449, reg301450)
reg301452 := PrimHead(reg301451)
__ctx.TailApply(__defun__shen_4pair, reg301452, Parse__Char)
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
__initDefs = append(__initDefs, defType{name: "shen.<plus>", value: __defun__shen_4_5plus_6})

__defun__shen_4_5stop_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2527 := __args[0]
_ = V2527
reg301456 := PrimHead(V2527)
reg301457 := PrimIsPair(reg301456)
if reg301457 == True {
reg301458 := PrimHead(V2527)
reg301459 := PrimHead(reg301458)
Parse__Char := reg301459
_ = Parse__Char
reg301460 := MakeNumber(46)
reg301461 := PrimEqual(Parse__Char, reg301460)
if reg301461 == True {
reg301462 := PrimHead(V2527)
reg301463 := PrimTail(reg301462)
reg301464 := __e.Call(__defun__shen_4hdtl, V2527)
reg301465 := __e.Call(__defun__shen_4pair, reg301463, reg301464)
reg301466 := PrimHead(reg301465)
__ctx.TailApply(__defun__shen_4pair, reg301466, Parse__Char)
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
__initDefs = append(__initDefs, defType{name: "shen.<stop>", value: __defun__shen_4_5stop_6})

__defun__shen_4_5predigits_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2529 := __args[0]
_ = V2529
reg301470 := __e.Call(__defun__shen_4_5digits_6, V2529)
Parse__shen_4_5digits_6 := reg301470
_ = Parse__shen_4_5digits_6
reg301471 := __e.Call(__defun__fail)
reg301472 := PrimEqual(reg301471, Parse__shen_4_5digits_6)
reg301473 := PrimNot(reg301472)
var reg301478 Obj
if reg301473 == True {
reg301474 := PrimHead(Parse__shen_4_5digits_6)
reg301475 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
reg301476 := __e.Call(__defun__shen_4pair, reg301474, reg301475)
reg301478 = reg301476
} else {
reg301477 := __e.Call(__defun__fail)
reg301478 = reg301477
}
YaccParse := reg301478
_ = YaccParse
reg301479 := __e.Call(__defun__fail)
reg301480 := PrimEqual(YaccParse, reg301479)
if reg301480 == True {
reg301481 := __e.Call(__defun___5e_6, V2529)
Parse___5e_6 := reg301481
_ = Parse___5e_6
reg301482 := __e.Call(__defun__fail)
reg301483 := PrimEqual(reg301482, Parse___5e_6)
reg301484 := PrimNot(reg301483)
if reg301484 == True {
reg301485 := PrimHead(Parse___5e_6)
reg301486 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg301485, reg301486)
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
__initDefs = append(__initDefs, defType{name: "shen.<predigits>", value: __defun__shen_4_5predigits_6})

__defun__shen_4_5postdigits_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2531 := __args[0]
_ = V2531
reg301489 := __e.Call(__defun__shen_4_5digits_6, V2531)
Parse__shen_4_5digits_6 := reg301489
_ = Parse__shen_4_5digits_6
reg301490 := __e.Call(__defun__fail)
reg301491 := PrimEqual(reg301490, Parse__shen_4_5digits_6)
reg301492 := PrimNot(reg301491)
if reg301492 == True {
reg301493 := PrimHead(Parse__shen_4_5digits_6)
reg301494 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
__ctx.TailApply(__defun__shen_4pair, reg301493, reg301494)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<postdigits>", value: __defun__shen_4_5postdigits_6})

__defun__shen_4_5digits_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2533 := __args[0]
_ = V2533
reg301497 := __e.Call(__defun__shen_4_5digit_6, V2533)
Parse__shen_4_5digit_6 := reg301497
_ = Parse__shen_4_5digit_6
reg301498 := __e.Call(__defun__fail)
reg301499 := PrimEqual(reg301498, Parse__shen_4_5digit_6)
reg301500 := PrimNot(reg301499)
var reg301513 Obj
if reg301500 == True {
reg301501 := __e.Call(__defun__shen_4_5digits_6, Parse__shen_4_5digit_6)
Parse__shen_4_5digits_6 := reg301501
_ = Parse__shen_4_5digits_6
reg301502 := __e.Call(__defun__fail)
reg301503 := PrimEqual(reg301502, Parse__shen_4_5digits_6)
reg301504 := PrimNot(reg301503)
var reg301511 Obj
if reg301504 == True {
reg301505 := PrimHead(Parse__shen_4_5digits_6)
reg301506 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digit_6)
reg301507 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digits_6)
reg301508 := PrimCons(reg301506, reg301507)
reg301509 := __e.Call(__defun__shen_4pair, reg301505, reg301508)
reg301511 = reg301509
} else {
reg301510 := __e.Call(__defun__fail)
reg301511 = reg301510
}
reg301513 = reg301511
} else {
reg301512 := __e.Call(__defun__fail)
reg301513 = reg301512
}
YaccParse := reg301513
_ = YaccParse
reg301514 := __e.Call(__defun__fail)
reg301515 := PrimEqual(YaccParse, reg301514)
if reg301515 == True {
reg301516 := __e.Call(__defun__shen_4_5digit_6, V2533)
Parse__shen_4_5digit_6 := reg301516
_ = Parse__shen_4_5digit_6
reg301517 := __e.Call(__defun__fail)
reg301518 := PrimEqual(reg301517, Parse__shen_4_5digit_6)
reg301519 := PrimNot(reg301518)
if reg301519 == True {
reg301520 := PrimHead(Parse__shen_4_5digit_6)
reg301521 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5digit_6)
reg301522 := Nil;
reg301523 := PrimCons(reg301521, reg301522)
__ctx.TailApply(__defun__shen_4pair, reg301520, reg301523)
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
__initDefs = append(__initDefs, defType{name: "shen.<digits>", value: __defun__shen_4_5digits_6})

__defun__shen_4_5digit_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2535 := __args[0]
_ = V2535
reg301526 := PrimHead(V2535)
reg301527 := PrimIsPair(reg301526)
if reg301527 == True {
reg301528 := PrimHead(V2535)
reg301529 := PrimHead(reg301528)
Parse__X := reg301529
_ = Parse__X
reg301530 := __e.Call(__defun__shen_4numbyte_2, Parse__X)
if reg301530 == True {
reg301531 := PrimHead(V2535)
reg301532 := PrimTail(reg301531)
reg301533 := __e.Call(__defun__shen_4hdtl, V2535)
reg301534 := __e.Call(__defun__shen_4pair, reg301532, reg301533)
reg301535 := PrimHead(reg301534)
reg301536 := __e.Call(__defun__shen_4byte_1_6digit, Parse__X)
__ctx.TailApply(__defun__shen_4pair, reg301535, reg301536)
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
__initDefs = append(__initDefs, defType{name: "shen.<digit>", value: __defun__shen_4_5digit_6})

__defun__shen_4byte_1_6digit = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2537 := __args[0]
_ = V2537
reg301540 := MakeNumber(48)
reg301541 := PrimEqual(reg301540, V2537)
if reg301541 == True {
reg301542 := MakeNumber(0)
__ctx.Return(reg301542)
return
} else {
reg301543 := MakeNumber(49)
reg301544 := PrimEqual(reg301543, V2537)
if reg301544 == True {
reg301545 := MakeNumber(1)
__ctx.Return(reg301545)
return
} else {
reg301546 := MakeNumber(50)
reg301547 := PrimEqual(reg301546, V2537)
if reg301547 == True {
reg301548 := MakeNumber(2)
__ctx.Return(reg301548)
return
} else {
reg301549 := MakeNumber(51)
reg301550 := PrimEqual(reg301549, V2537)
if reg301550 == True {
reg301551 := MakeNumber(3)
__ctx.Return(reg301551)
return
} else {
reg301552 := MakeNumber(52)
reg301553 := PrimEqual(reg301552, V2537)
if reg301553 == True {
reg301554 := MakeNumber(4)
__ctx.Return(reg301554)
return
} else {
reg301555 := MakeNumber(53)
reg301556 := PrimEqual(reg301555, V2537)
if reg301556 == True {
reg301557 := MakeNumber(5)
__ctx.Return(reg301557)
return
} else {
reg301558 := MakeNumber(54)
reg301559 := PrimEqual(reg301558, V2537)
if reg301559 == True {
reg301560 := MakeNumber(6)
__ctx.Return(reg301560)
return
} else {
reg301561 := MakeNumber(55)
reg301562 := PrimEqual(reg301561, V2537)
if reg301562 == True {
reg301563 := MakeNumber(7)
__ctx.Return(reg301563)
return
} else {
reg301564 := MakeNumber(56)
reg301565 := PrimEqual(reg301564, V2537)
if reg301565 == True {
reg301566 := MakeNumber(8)
__ctx.Return(reg301566)
return
} else {
reg301567 := MakeNumber(57)
reg301568 := PrimEqual(reg301567, V2537)
if reg301568 == True {
reg301569 := MakeNumber(9)
__ctx.Return(reg301569)
return
} else {
reg301570 := MakeSymbol("shen.byte->digit")
__ctx.TailApply(__defun__shen_4f__error, reg301570)
return
}
}
}
}
}
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.byte->digit", value: __defun__shen_4byte_1_6digit})

__defun__shen_4pre = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2542 := __args[0]
_ = V2542
V2543 := __args[1]
_ = V2543
reg301572 := Nil;
reg301573 := PrimEqual(reg301572, V2542)
if reg301573 == True {
reg301574 := MakeNumber(0)
__ctx.Return(reg301574)
return
} else {
reg301575 := PrimIsPair(V2542)
if reg301575 == True {
reg301576 := MakeNumber(10)
reg301577 := __e.Call(__defun__shen_4expt, reg301576, V2543)
reg301578 := PrimHead(V2542)
reg301579 := PrimNumberMultiply(reg301577, reg301578)
reg301580 := PrimTail(V2542)
reg301581 := MakeNumber(1)
reg301582 := PrimNumberAdd(V2543, reg301581)
reg301583 := __e.Call(__defun__shen_4pre, reg301580, reg301582)
reg301584 := PrimNumberAdd(reg301579, reg301583)
__ctx.Return(reg301584)
return
} else {
reg301585 := MakeSymbol("shen.pre")
__ctx.TailApply(__defun__shen_4f__error, reg301585)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.pre", value: __defun__shen_4pre})

__defun__shen_4post = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2548 := __args[0]
_ = V2548
V2549 := __args[1]
_ = V2549
reg301587 := Nil;
reg301588 := PrimEqual(reg301587, V2548)
if reg301588 == True {
reg301589 := MakeNumber(0)
__ctx.Return(reg301589)
return
} else {
reg301590 := PrimIsPair(V2548)
if reg301590 == True {
reg301591 := MakeNumber(10)
reg301592 := MakeNumber(0)
reg301593 := PrimNumberSubtract(reg301592, V2549)
reg301594 := __e.Call(__defun__shen_4expt, reg301591, reg301593)
reg301595 := PrimHead(V2548)
reg301596 := PrimNumberMultiply(reg301594, reg301595)
reg301597 := PrimTail(V2548)
reg301598 := MakeNumber(1)
reg301599 := PrimNumberAdd(V2549, reg301598)
reg301600 := __e.Call(__defun__shen_4post, reg301597, reg301599)
reg301601 := PrimNumberAdd(reg301596, reg301600)
__ctx.Return(reg301601)
return
} else {
reg301602 := MakeSymbol("shen.post")
__ctx.TailApply(__defun__shen_4f__error, reg301602)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.post", value: __defun__shen_4post})

__defun__shen_4expt = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2554 := __args[0]
_ = V2554
V2555 := __args[1]
_ = V2555
reg301604 := MakeNumber(0)
reg301605 := PrimEqual(reg301604, V2555)
if reg301605 == True {
reg301606 := MakeNumber(1)
__ctx.Return(reg301606)
return
} else {
reg301607 := MakeNumber(0)
reg301608 := PrimGreatThan(V2555, reg301607)
if reg301608 == True {
reg301609 := MakeNumber(1)
reg301610 := PrimNumberSubtract(V2555, reg301609)
reg301611 := __e.Call(__defun__shen_4expt, V2554, reg301610)
reg301612 := PrimNumberMultiply(V2554, reg301611)
__ctx.Return(reg301612)
return
} else {
reg301613 := MakeNumber(1)
reg301614 := MakeNumber(1)
reg301615 := PrimNumberAdd(V2555, reg301614)
reg301616 := __e.Call(__defun__shen_4expt, V2554, reg301615)
reg301617 := PrimNumberDivide(reg301616, V2554)
reg301618 := PrimNumberMultiply(reg301613, reg301617)
__ctx.Return(reg301618)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.expt", value: __defun__shen_4expt})

__defun__shen_4_5st__input1_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2557 := __args[0]
_ = V2557
reg301619 := __e.Call(__defun__shen_4_5st__input_6, V2557)
Parse__shen_4_5st__input_6 := reg301619
_ = Parse__shen_4_5st__input_6
reg301620 := __e.Call(__defun__fail)
reg301621 := PrimEqual(reg301620, Parse__shen_4_5st__input_6)
reg301622 := PrimNot(reg301621)
if reg301622 == True {
reg301623 := PrimHead(Parse__shen_4_5st__input_6)
reg301624 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
__ctx.TailApply(__defun__shen_4pair, reg301623, reg301624)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<st_input1>", value: __defun__shen_4_5st__input1_6})

__defun__shen_4_5st__input2_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2559 := __args[0]
_ = V2559
reg301627 := __e.Call(__defun__shen_4_5st__input_6, V2559)
Parse__shen_4_5st__input_6 := reg301627
_ = Parse__shen_4_5st__input_6
reg301628 := __e.Call(__defun__fail)
reg301629 := PrimEqual(reg301628, Parse__shen_4_5st__input_6)
reg301630 := PrimNot(reg301629)
if reg301630 == True {
reg301631 := PrimHead(Parse__shen_4_5st__input_6)
reg301632 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5st__input_6)
__ctx.TailApply(__defun__shen_4pair, reg301631, reg301632)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<st_input2>", value: __defun__shen_4_5st__input2_6})

__defun__shen_4_5comment_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2561 := __args[0]
_ = V2561
reg301635 := __e.Call(__defun__shen_4_5singleline_6, V2561)
Parse__shen_4_5singleline_6 := reg301635
_ = Parse__shen_4_5singleline_6
reg301636 := __e.Call(__defun__fail)
reg301637 := PrimEqual(reg301636, Parse__shen_4_5singleline_6)
reg301638 := PrimNot(reg301637)
var reg301643 Obj
if reg301638 == True {
reg301639 := PrimHead(Parse__shen_4_5singleline_6)
reg301640 := MakeSymbol("shen.skip")
reg301641 := __e.Call(__defun__shen_4pair, reg301639, reg301640)
reg301643 = reg301641
} else {
reg301642 := __e.Call(__defun__fail)
reg301643 = reg301642
}
YaccParse := reg301643
_ = YaccParse
reg301644 := __e.Call(__defun__fail)
reg301645 := PrimEqual(YaccParse, reg301644)
if reg301645 == True {
reg301646 := __e.Call(__defun__shen_4_5multiline_6, V2561)
Parse__shen_4_5multiline_6 := reg301646
_ = Parse__shen_4_5multiline_6
reg301647 := __e.Call(__defun__fail)
reg301648 := PrimEqual(reg301647, Parse__shen_4_5multiline_6)
reg301649 := PrimNot(reg301648)
if reg301649 == True {
reg301650 := PrimHead(Parse__shen_4_5multiline_6)
reg301651 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301650, reg301651)
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
__initDefs = append(__initDefs, defType{name: "shen.<comment>", value: __defun__shen_4_5comment_6})

__defun__shen_4_5singleline_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2563 := __args[0]
_ = V2563
reg301654 := __e.Call(__defun__shen_4_5backslash_6, V2563)
Parse__shen_4_5backslash_6 := reg301654
_ = Parse__shen_4_5backslash_6
reg301655 := __e.Call(__defun__fail)
reg301656 := PrimEqual(reg301655, Parse__shen_4_5backslash_6)
reg301657 := PrimNot(reg301656)
if reg301657 == True {
reg301658 := __e.Call(__defun__shen_4_5backslash_6, Parse__shen_4_5backslash_6)
Parse__shen_4_5backslash_6 := reg301658
_ = Parse__shen_4_5backslash_6
reg301659 := __e.Call(__defun__fail)
reg301660 := PrimEqual(reg301659, Parse__shen_4_5backslash_6)
reg301661 := PrimNot(reg301660)
if reg301661 == True {
reg301662 := __e.Call(__defun__shen_4_5anysingle_6, Parse__shen_4_5backslash_6)
Parse__shen_4_5anysingle_6 := reg301662
_ = Parse__shen_4_5anysingle_6
reg301663 := __e.Call(__defun__fail)
reg301664 := PrimEqual(reg301663, Parse__shen_4_5anysingle_6)
reg301665 := PrimNot(reg301664)
if reg301665 == True {
reg301666 := __e.Call(__defun__shen_4_5return_6, Parse__shen_4_5anysingle_6)
Parse__shen_4_5return_6 := reg301666
_ = Parse__shen_4_5return_6
reg301667 := __e.Call(__defun__fail)
reg301668 := PrimEqual(reg301667, Parse__shen_4_5return_6)
reg301669 := PrimNot(reg301668)
if reg301669 == True {
reg301670 := PrimHead(Parse__shen_4_5return_6)
reg301671 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301670, reg301671)
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
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<singleline>", value: __defun__shen_4_5singleline_6})

__defun__shen_4_5backslash_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2565 := __args[0]
_ = V2565
reg301677 := PrimHead(V2565)
reg301678 := PrimIsPair(reg301677)
var reg301687 Obj
if reg301678 == True {
reg301679 := MakeNumber(92)
reg301680 := PrimHead(V2565)
reg301681 := PrimHead(reg301680)
reg301682 := PrimEqual(reg301679, reg301681)
var reg301685 Obj
if reg301682 == True {
reg301683 := True;
reg301685 = reg301683
} else {
reg301684 := False;
reg301685 = reg301684
}
reg301687 = reg301685
} else {
reg301686 := False;
reg301687 = reg301686
}
if reg301687 == True {
reg301688 := PrimHead(V2565)
reg301689 := PrimTail(reg301688)
reg301690 := __e.Call(__defun__shen_4hdtl, V2565)
reg301691 := __e.Call(__defun__shen_4pair, reg301689, reg301690)
reg301692 := PrimHead(reg301691)
reg301693 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301692, reg301693)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<backslash>", value: __defun__shen_4_5backslash_6})

__defun__shen_4_5anysingle_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2567 := __args[0]
_ = V2567
reg301696 := __e.Call(__defun__shen_4_5non_1return_6, V2567)
Parse__shen_4_5non_1return_6 := reg301696
_ = Parse__shen_4_5non_1return_6
reg301697 := __e.Call(__defun__fail)
reg301698 := PrimEqual(reg301697, Parse__shen_4_5non_1return_6)
reg301699 := PrimNot(reg301698)
var reg301710 Obj
if reg301699 == True {
reg301700 := __e.Call(__defun__shen_4_5anysingle_6, Parse__shen_4_5non_1return_6)
Parse__shen_4_5anysingle_6 := reg301700
_ = Parse__shen_4_5anysingle_6
reg301701 := __e.Call(__defun__fail)
reg301702 := PrimEqual(reg301701, Parse__shen_4_5anysingle_6)
reg301703 := PrimNot(reg301702)
var reg301708 Obj
if reg301703 == True {
reg301704 := PrimHead(Parse__shen_4_5anysingle_6)
reg301705 := MakeSymbol("shen.skip")
reg301706 := __e.Call(__defun__shen_4pair, reg301704, reg301705)
reg301708 = reg301706
} else {
reg301707 := __e.Call(__defun__fail)
reg301708 = reg301707
}
reg301710 = reg301708
} else {
reg301709 := __e.Call(__defun__fail)
reg301710 = reg301709
}
YaccParse := reg301710
_ = YaccParse
reg301711 := __e.Call(__defun__fail)
reg301712 := PrimEqual(YaccParse, reg301711)
if reg301712 == True {
reg301713 := __e.Call(__defun___5e_6, V2567)
Parse___5e_6 := reg301713
_ = Parse___5e_6
reg301714 := __e.Call(__defun__fail)
reg301715 := PrimEqual(reg301714, Parse___5e_6)
reg301716 := PrimNot(reg301715)
if reg301716 == True {
reg301717 := PrimHead(Parse___5e_6)
reg301718 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301717, reg301718)
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
__initDefs = append(__initDefs, defType{name: "shen.<anysingle>", value: __defun__shen_4_5anysingle_6})

__defun__shen_4_5non_1return_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2569 := __args[0]
_ = V2569
reg301721 := PrimHead(V2569)
reg301722 := PrimIsPair(reg301721)
if reg301722 == True {
reg301723 := PrimHead(V2569)
reg301724 := PrimHead(reg301723)
Parse__X := reg301724
_ = Parse__X
reg301725 := MakeNumber(10)
reg301726 := MakeNumber(13)
reg301727 := Nil;
reg301728 := PrimCons(reg301726, reg301727)
reg301729 := PrimCons(reg301725, reg301728)
reg301730 := __e.Call(__defun__element_2, Parse__X, reg301729)
reg301731 := PrimNot(reg301730)
if reg301731 == True {
reg301732 := PrimHead(V2569)
reg301733 := PrimTail(reg301732)
reg301734 := __e.Call(__defun__shen_4hdtl, V2569)
reg301735 := __e.Call(__defun__shen_4pair, reg301733, reg301734)
reg301736 := PrimHead(reg301735)
reg301737 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301736, reg301737)
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
__initDefs = append(__initDefs, defType{name: "shen.<non-return>", value: __defun__shen_4_5non_1return_6})

__defun__shen_4_5return_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2571 := __args[0]
_ = V2571
reg301741 := PrimHead(V2571)
reg301742 := PrimIsPair(reg301741)
if reg301742 == True {
reg301743 := PrimHead(V2571)
reg301744 := PrimHead(reg301743)
Parse__X := reg301744
_ = Parse__X
reg301745 := MakeNumber(10)
reg301746 := MakeNumber(13)
reg301747 := Nil;
reg301748 := PrimCons(reg301746, reg301747)
reg301749 := PrimCons(reg301745, reg301748)
reg301750 := __e.Call(__defun__element_2, Parse__X, reg301749)
if reg301750 == True {
reg301751 := PrimHead(V2571)
reg301752 := PrimTail(reg301751)
reg301753 := __e.Call(__defun__shen_4hdtl, V2571)
reg301754 := __e.Call(__defun__shen_4pair, reg301752, reg301753)
reg301755 := PrimHead(reg301754)
reg301756 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301755, reg301756)
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
__initDefs = append(__initDefs, defType{name: "shen.<return>", value: __defun__shen_4_5return_6})

__defun__shen_4_5multiline_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2573 := __args[0]
_ = V2573
reg301760 := __e.Call(__defun__shen_4_5backslash_6, V2573)
Parse__shen_4_5backslash_6 := reg301760
_ = Parse__shen_4_5backslash_6
reg301761 := __e.Call(__defun__fail)
reg301762 := PrimEqual(reg301761, Parse__shen_4_5backslash_6)
reg301763 := PrimNot(reg301762)
if reg301763 == True {
reg301764 := __e.Call(__defun__shen_4_5times_6, Parse__shen_4_5backslash_6)
Parse__shen_4_5times_6 := reg301764
_ = Parse__shen_4_5times_6
reg301765 := __e.Call(__defun__fail)
reg301766 := PrimEqual(reg301765, Parse__shen_4_5times_6)
reg301767 := PrimNot(reg301766)
if reg301767 == True {
reg301768 := __e.Call(__defun__shen_4_5anymulti_6, Parse__shen_4_5times_6)
Parse__shen_4_5anymulti_6 := reg301768
_ = Parse__shen_4_5anymulti_6
reg301769 := __e.Call(__defun__fail)
reg301770 := PrimEqual(reg301769, Parse__shen_4_5anymulti_6)
reg301771 := PrimNot(reg301770)
if reg301771 == True {
reg301772 := PrimHead(Parse__shen_4_5anymulti_6)
reg301773 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301772, reg301773)
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
__initDefs = append(__initDefs, defType{name: "shen.<multiline>", value: __defun__shen_4_5multiline_6})

__defun__shen_4_5times_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2575 := __args[0]
_ = V2575
reg301778 := PrimHead(V2575)
reg301779 := PrimIsPair(reg301778)
var reg301788 Obj
if reg301779 == True {
reg301780 := MakeNumber(42)
reg301781 := PrimHead(V2575)
reg301782 := PrimHead(reg301781)
reg301783 := PrimEqual(reg301780, reg301782)
var reg301786 Obj
if reg301783 == True {
reg301784 := True;
reg301786 = reg301784
} else {
reg301785 := False;
reg301786 = reg301785
}
reg301788 = reg301786
} else {
reg301787 := False;
reg301788 = reg301787
}
if reg301788 == True {
reg301789 := PrimHead(V2575)
reg301790 := PrimTail(reg301789)
reg301791 := __e.Call(__defun__shen_4hdtl, V2575)
reg301792 := __e.Call(__defun__shen_4pair, reg301790, reg301791)
reg301793 := PrimHead(reg301792)
reg301794 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301793, reg301794)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<times>", value: __defun__shen_4_5times_6})

__defun__shen_4_5anymulti_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2577 := __args[0]
_ = V2577
reg301797 := __e.Call(__defun__shen_4_5comment_6, V2577)
Parse__shen_4_5comment_6 := reg301797
_ = Parse__shen_4_5comment_6
reg301798 := __e.Call(__defun__fail)
reg301799 := PrimEqual(reg301798, Parse__shen_4_5comment_6)
reg301800 := PrimNot(reg301799)
var reg301811 Obj
if reg301800 == True {
reg301801 := __e.Call(__defun__shen_4_5anymulti_6, Parse__shen_4_5comment_6)
Parse__shen_4_5anymulti_6 := reg301801
_ = Parse__shen_4_5anymulti_6
reg301802 := __e.Call(__defun__fail)
reg301803 := PrimEqual(reg301802, Parse__shen_4_5anymulti_6)
reg301804 := PrimNot(reg301803)
var reg301809 Obj
if reg301804 == True {
reg301805 := PrimHead(Parse__shen_4_5anymulti_6)
reg301806 := MakeSymbol("shen.skip")
reg301807 := __e.Call(__defun__shen_4pair, reg301805, reg301806)
reg301809 = reg301807
} else {
reg301808 := __e.Call(__defun__fail)
reg301809 = reg301808
}
reg301811 = reg301809
} else {
reg301810 := __e.Call(__defun__fail)
reg301811 = reg301810
}
YaccParse := reg301811
_ = YaccParse
reg301812 := __e.Call(__defun__fail)
reg301813 := PrimEqual(YaccParse, reg301812)
if reg301813 == True {
reg301814 := __e.Call(__defun__shen_4_5times_6, V2577)
Parse__shen_4_5times_6 := reg301814
_ = Parse__shen_4_5times_6
reg301815 := __e.Call(__defun__fail)
reg301816 := PrimEqual(reg301815, Parse__shen_4_5times_6)
reg301817 := PrimNot(reg301816)
var reg301828 Obj
if reg301817 == True {
reg301818 := __e.Call(__defun__shen_4_5backslash_6, Parse__shen_4_5times_6)
Parse__shen_4_5backslash_6 := reg301818
_ = Parse__shen_4_5backslash_6
reg301819 := __e.Call(__defun__fail)
reg301820 := PrimEqual(reg301819, Parse__shen_4_5backslash_6)
reg301821 := PrimNot(reg301820)
var reg301826 Obj
if reg301821 == True {
reg301822 := PrimHead(Parse__shen_4_5backslash_6)
reg301823 := MakeSymbol("shen.skip")
reg301824 := __e.Call(__defun__shen_4pair, reg301822, reg301823)
reg301826 = reg301824
} else {
reg301825 := __e.Call(__defun__fail)
reg301826 = reg301825
}
reg301828 = reg301826
} else {
reg301827 := __e.Call(__defun__fail)
reg301828 = reg301827
}
YaccParse := reg301828
_ = YaccParse
reg301829 := __e.Call(__defun__fail)
reg301830 := PrimEqual(YaccParse, reg301829)
if reg301830 == True {
reg301831 := PrimHead(V2577)
reg301832 := PrimIsPair(reg301831)
if reg301832 == True {
reg301833 := PrimHead(V2577)
reg301834 := PrimHead(reg301833)
Parse__X := reg301834
_ = Parse__X
reg301835 := PrimHead(V2577)
reg301836 := PrimTail(reg301835)
reg301837 := __e.Call(__defun__shen_4hdtl, V2577)
reg301838 := __e.Call(__defun__shen_4pair, reg301836, reg301837)
reg301839 := __e.Call(__defun__shen_4_5anymulti_6, reg301838)
Parse__shen_4_5anymulti_6 := reg301839
_ = Parse__shen_4_5anymulti_6
reg301840 := __e.Call(__defun__fail)
reg301841 := PrimEqual(reg301840, Parse__shen_4_5anymulti_6)
reg301842 := PrimNot(reg301841)
if reg301842 == True {
reg301843 := PrimHead(Parse__shen_4_5anymulti_6)
reg301844 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301843, reg301844)
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
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<anymulti>", value: __defun__shen_4_5anymulti_6})

__defun__shen_4_5whitespaces_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2579 := __args[0]
_ = V2579
reg301848 := __e.Call(__defun__shen_4_5whitespace_6, V2579)
Parse__shen_4_5whitespace_6 := reg301848
_ = Parse__shen_4_5whitespace_6
reg301849 := __e.Call(__defun__fail)
reg301850 := PrimEqual(reg301849, Parse__shen_4_5whitespace_6)
reg301851 := PrimNot(reg301850)
var reg301862 Obj
if reg301851 == True {
reg301852 := __e.Call(__defun__shen_4_5whitespaces_6, Parse__shen_4_5whitespace_6)
Parse__shen_4_5whitespaces_6 := reg301852
_ = Parse__shen_4_5whitespaces_6
reg301853 := __e.Call(__defun__fail)
reg301854 := PrimEqual(reg301853, Parse__shen_4_5whitespaces_6)
reg301855 := PrimNot(reg301854)
var reg301860 Obj
if reg301855 == True {
reg301856 := PrimHead(Parse__shen_4_5whitespaces_6)
reg301857 := MakeSymbol("shen.skip")
reg301858 := __e.Call(__defun__shen_4pair, reg301856, reg301857)
reg301860 = reg301858
} else {
reg301859 := __e.Call(__defun__fail)
reg301860 = reg301859
}
reg301862 = reg301860
} else {
reg301861 := __e.Call(__defun__fail)
reg301862 = reg301861
}
YaccParse := reg301862
_ = YaccParse
reg301863 := __e.Call(__defun__fail)
reg301864 := PrimEqual(YaccParse, reg301863)
if reg301864 == True {
reg301865 := __e.Call(__defun__shen_4_5whitespace_6, V2579)
Parse__shen_4_5whitespace_6 := reg301865
_ = Parse__shen_4_5whitespace_6
reg301866 := __e.Call(__defun__fail)
reg301867 := PrimEqual(reg301866, Parse__shen_4_5whitespace_6)
reg301868 := PrimNot(reg301867)
if reg301868 == True {
reg301869 := PrimHead(Parse__shen_4_5whitespace_6)
reg301870 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301869, reg301870)
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
__initDefs = append(__initDefs, defType{name: "shen.<whitespaces>", value: __defun__shen_4_5whitespaces_6})

__defun__shen_4_5whitespace_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2581 := __args[0]
_ = V2581
reg301873 := PrimHead(V2581)
reg301874 := PrimIsPair(reg301873)
if reg301874 == True {
reg301875 := PrimHead(V2581)
reg301876 := PrimHead(reg301875)
Parse__X := reg301876
_ = Parse__X
Parse__Case := Parse__X
_ = Parse__Case
reg301877 := MakeNumber(32)
reg301878 := PrimEqual(Parse__Case, reg301877)
var reg301899 Obj
if reg301878 == True {
reg301879 := True;
reg301899 = reg301879
} else {
reg301880 := MakeNumber(13)
reg301881 := PrimEqual(Parse__Case, reg301880)
var reg301895 Obj
if reg301881 == True {
reg301882 := True;
reg301895 = reg301882
} else {
reg301883 := MakeNumber(10)
reg301884 := PrimEqual(Parse__Case, reg301883)
var reg301891 Obj
if reg301884 == True {
reg301885 := True;
reg301891 = reg301885
} else {
reg301886 := MakeNumber(9)
reg301887 := PrimEqual(Parse__Case, reg301886)
var reg301890 Obj
if reg301887 == True {
reg301888 := True;
reg301890 = reg301888
} else {
reg301889 := False;
reg301890 = reg301889
}
reg301891 = reg301890
}
var reg301894 Obj
if reg301891 == True {
reg301892 := True;
reg301894 = reg301892
} else {
reg301893 := False;
reg301894 = reg301893
}
reg301895 = reg301894
}
var reg301898 Obj
if reg301895 == True {
reg301896 := True;
reg301898 = reg301896
} else {
reg301897 := False;
reg301898 = reg301897
}
reg301899 = reg301898
}
if reg301899 == True {
reg301900 := PrimHead(V2581)
reg301901 := PrimTail(reg301900)
reg301902 := __e.Call(__defun__shen_4hdtl, V2581)
reg301903 := __e.Call(__defun__shen_4pair, reg301901, reg301902)
reg301904 := PrimHead(reg301903)
reg301905 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg301904, reg301905)
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
__initDefs = append(__initDefs, defType{name: "shen.<whitespace>", value: __defun__shen_4_5whitespace_6})

__defun__shen_4cons__form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2583 := __args[0]
_ = V2583
reg301909 := Nil;
reg301910 := PrimEqual(reg301909, V2583)
if reg301910 == True {
reg301911 := Nil;
__ctx.Return(reg301911)
return
} else {
reg301912 := PrimIsPair(V2583)
var reg301946 Obj
if reg301912 == True {
reg301913 := PrimTail(V2583)
reg301914 := PrimIsPair(reg301913)
var reg301941 Obj
if reg301914 == True {
reg301915 := PrimTail(V2583)
reg301916 := PrimTail(reg301915)
reg301917 := PrimIsPair(reg301916)
var reg301936 Obj
if reg301917 == True {
reg301918 := Nil;
reg301919 := PrimTail(V2583)
reg301920 := PrimTail(reg301919)
reg301921 := PrimTail(reg301920)
reg301922 := PrimEqual(reg301918, reg301921)
var reg301931 Obj
if reg301922 == True {
reg301923 := PrimTail(V2583)
reg301924 := PrimHead(reg301923)
reg301925 := MakeSymbol("bar!")
reg301926 := PrimEqual(reg301924, reg301925)
var reg301929 Obj
if reg301926 == True {
reg301927 := True;
reg301929 = reg301927
} else {
reg301928 := False;
reg301929 = reg301928
}
reg301931 = reg301929
} else {
reg301930 := False;
reg301931 = reg301930
}
var reg301934 Obj
if reg301931 == True {
reg301932 := True;
reg301934 = reg301932
} else {
reg301933 := False;
reg301934 = reg301933
}
reg301936 = reg301934
} else {
reg301935 := False;
reg301936 = reg301935
}
var reg301939 Obj
if reg301936 == True {
reg301937 := True;
reg301939 = reg301937
} else {
reg301938 := False;
reg301939 = reg301938
}
reg301941 = reg301939
} else {
reg301940 := False;
reg301941 = reg301940
}
var reg301944 Obj
if reg301941 == True {
reg301942 := True;
reg301944 = reg301942
} else {
reg301943 := False;
reg301944 = reg301943
}
reg301946 = reg301944
} else {
reg301945 := False;
reg301946 = reg301945
}
if reg301946 == True {
reg301947 := MakeSymbol("cons")
reg301948 := PrimHead(V2583)
reg301949 := PrimTail(V2583)
reg301950 := PrimTail(reg301949)
reg301951 := PrimCons(reg301948, reg301950)
reg301952 := PrimCons(reg301947, reg301951)
__ctx.Return(reg301952)
return
} else {
reg301953 := PrimIsPair(V2583)
if reg301953 == True {
reg301954 := MakeSymbol("cons")
reg301955 := PrimHead(V2583)
reg301956 := PrimTail(V2583)
reg301957 := __e.Call(__defun__shen_4cons__form, reg301956)
reg301958 := Nil;
reg301959 := PrimCons(reg301957, reg301958)
reg301960 := PrimCons(reg301955, reg301959)
reg301961 := PrimCons(reg301954, reg301960)
__ctx.Return(reg301961)
return
} else {
reg301962 := MakeSymbol("shen.cons_form")
__ctx.TailApply(__defun__shen_4f__error, reg301962)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.cons_form", value: __defun__shen_4cons__form})

__defun__shen_4package_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2588 := __args[0]
_ = V2588
V2589 := __args[1]
_ = V2589
reg301964 := PrimIsPair(V2588)
var reg301988 Obj
if reg301964 == True {
reg301965 := MakeSymbol("$")
reg301966 := PrimHead(V2588)
reg301967 := PrimEqual(reg301965, reg301966)
var reg301983 Obj
if reg301967 == True {
reg301968 := PrimTail(V2588)
reg301969 := PrimIsPair(reg301968)
var reg301978 Obj
if reg301969 == True {
reg301970 := Nil;
reg301971 := PrimTail(V2588)
reg301972 := PrimTail(reg301971)
reg301973 := PrimEqual(reg301970, reg301972)
var reg301976 Obj
if reg301973 == True {
reg301974 := True;
reg301976 = reg301974
} else {
reg301975 := False;
reg301976 = reg301975
}
reg301978 = reg301976
} else {
reg301977 := False;
reg301978 = reg301977
}
var reg301981 Obj
if reg301978 == True {
reg301979 := True;
reg301981 = reg301979
} else {
reg301980 := False;
reg301981 = reg301980
}
reg301983 = reg301981
} else {
reg301982 := False;
reg301983 = reg301982
}
var reg301986 Obj
if reg301983 == True {
reg301984 := True;
reg301986 = reg301984
} else {
reg301985 := False;
reg301986 = reg301985
}
reg301988 = reg301986
} else {
reg301987 := False;
reg301988 = reg301987
}
if reg301988 == True {
reg301989 := PrimTail(V2588)
reg301990 := PrimHead(reg301989)
reg301991 := __e.Call(__defun__explode, reg301990)
__ctx.TailApply(__defun__append, reg301991, V2589)
return
} else {
reg301993 := PrimIsPair(V2588)
var reg302025 Obj
if reg301993 == True {
reg301994 := MakeSymbol("package")
reg301995 := PrimHead(V2588)
reg301996 := PrimEqual(reg301994, reg301995)
var reg302020 Obj
if reg301996 == True {
reg301997 := PrimTail(V2588)
reg301998 := PrimIsPair(reg301997)
var reg302015 Obj
if reg301998 == True {
reg301999 := MakeSymbol("null")
reg302000 := PrimTail(V2588)
reg302001 := PrimHead(reg302000)
reg302002 := PrimEqual(reg301999, reg302001)
var reg302010 Obj
if reg302002 == True {
reg302003 := PrimTail(V2588)
reg302004 := PrimTail(reg302003)
reg302005 := PrimIsPair(reg302004)
var reg302008 Obj
if reg302005 == True {
reg302006 := True;
reg302008 = reg302006
} else {
reg302007 := False;
reg302008 = reg302007
}
reg302010 = reg302008
} else {
reg302009 := False;
reg302010 = reg302009
}
var reg302013 Obj
if reg302010 == True {
reg302011 := True;
reg302013 = reg302011
} else {
reg302012 := False;
reg302013 = reg302012
}
reg302015 = reg302013
} else {
reg302014 := False;
reg302015 = reg302014
}
var reg302018 Obj
if reg302015 == True {
reg302016 := True;
reg302018 = reg302016
} else {
reg302017 := False;
reg302018 = reg302017
}
reg302020 = reg302018
} else {
reg302019 := False;
reg302020 = reg302019
}
var reg302023 Obj
if reg302020 == True {
reg302021 := True;
reg302023 = reg302021
} else {
reg302022 := False;
reg302023 = reg302022
}
reg302025 = reg302023
} else {
reg302024 := False;
reg302025 = reg302024
}
if reg302025 == True {
reg302026 := PrimTail(V2588)
reg302027 := PrimTail(reg302026)
reg302028 := PrimTail(reg302027)
__ctx.TailApply(__defun__append, reg302028, V2589)
return
} else {
reg302030 := PrimIsPair(V2588)
var reg302053 Obj
if reg302030 == True {
reg302031 := MakeSymbol("package")
reg302032 := PrimHead(V2588)
reg302033 := PrimEqual(reg302031, reg302032)
var reg302048 Obj
if reg302033 == True {
reg302034 := PrimTail(V2588)
reg302035 := PrimIsPair(reg302034)
var reg302043 Obj
if reg302035 == True {
reg302036 := PrimTail(V2588)
reg302037 := PrimTail(reg302036)
reg302038 := PrimIsPair(reg302037)
var reg302041 Obj
if reg302038 == True {
reg302039 := True;
reg302041 = reg302039
} else {
reg302040 := False;
reg302041 = reg302040
}
reg302043 = reg302041
} else {
reg302042 := False;
reg302043 = reg302042
}
var reg302046 Obj
if reg302043 == True {
reg302044 := True;
reg302046 = reg302044
} else {
reg302045 := False;
reg302046 = reg302045
}
reg302048 = reg302046
} else {
reg302047 := False;
reg302048 = reg302047
}
var reg302051 Obj
if reg302048 == True {
reg302049 := True;
reg302051 = reg302049
} else {
reg302050 := False;
reg302051 = reg302050
}
reg302053 = reg302051
} else {
reg302052 := False;
reg302053 = reg302052
}
if reg302053 == True {
reg302054 := PrimTail(V2588)
reg302055 := PrimTail(reg302054)
reg302056 := PrimHead(reg302055)
reg302057 := __e.Call(__defun__shen_4eval_1without_1macros, reg302056)
ListofExceptions := reg302057
_ = ListofExceptions
reg302058 := PrimTail(V2588)
reg302059 := PrimHead(reg302058)
reg302060 := __e.Call(__defun__shen_4record_1exceptions, ListofExceptions, reg302059)
External := reg302060
_ = External
reg302061 := PrimTail(V2588)
reg302062 := PrimHead(reg302061)
reg302063 := PrimStr(reg302062)
reg302064 := MakeString(".")
reg302065 := PrimStringConcat(reg302063, reg302064)
reg302066 := PrimIntern(reg302065)
PackageNameDot := reg302066
_ = PackageNameDot
reg302067 := __e.Call(__defun__explode, PackageNameDot)
ExpPackageNameDot := reg302067
_ = ExpPackageNameDot
reg302068 := PrimTail(V2588)
reg302069 := PrimTail(reg302068)
reg302070 := PrimTail(reg302069)
reg302071 := __e.Call(__defun__shen_4packageh, PackageNameDot, ListofExceptions, reg302070, ExpPackageNameDot)
Packaged := reg302071
_ = Packaged
reg302072 := PrimTail(V2588)
reg302073 := PrimHead(reg302072)
reg302074 := __e.Call(__defun__shen_4internal_1symbols, ExpPackageNameDot, Packaged)
reg302075 := __e.Call(__defun__shen_4record_1internal, reg302073, reg302074)
Internal := reg302075
_ = Internal
__ctx.TailApply(__defun__append, Packaged, V2589)
return
} else {
reg302077 := PrimCons(V2588, V2589)
__ctx.Return(reg302077)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.package-macro", value: __defun__shen_4package_1macro})

__defun__shen_4record_1exceptions = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2592 := __args[0]
_ = V2592
V2593 := __args[1]
_ = V2593
reg302078 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg302079 := MakeSymbol("shen.external-symbols")
reg302080 := MakeSymbol("*property-vector*")
reg302081 := PrimValue(reg302080)
__ctx.TailApply(__defun__get, V2593, reg302079, reg302081)
return
}, 0)
reg302083 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg302084 := Nil;
__ctx.Return(reg302084)
return
}, 1)
reg302085 := __e.Try(reg302078).Catch(reg302083)
CurrExceptions := reg302085
_ = CurrExceptions
reg302086 := __e.Call(__defun__union, V2592, CurrExceptions)
AllExceptions := reg302086
_ = AllExceptions
reg302087 := MakeSymbol("shen.external-symbols")
reg302088 := MakeSymbol("*property-vector*")
reg302089 := PrimValue(reg302088)
__ctx.TailApply(__defun__put, V2593, reg302087, AllExceptions, reg302089)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.record-exceptions", value: __defun__shen_4record_1exceptions})

__defun__shen_4record_1internal = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2596 := __args[0]
_ = V2596
V2597 := __args[1]
_ = V2597
reg302091 := MakeSymbol("shen.internal-symbols")
reg302092 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg302093 := MakeSymbol("shen.internal-symbols")
reg302094 := MakeSymbol("*property-vector*")
reg302095 := PrimValue(reg302094)
__ctx.TailApply(__defun__get, V2596, reg302093, reg302095)
return
}, 0)
reg302097 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg302098 := Nil;
__ctx.Return(reg302098)
return
}, 1)
reg302099 := __e.Try(reg302092).Catch(reg302097)
reg302100 := __e.Call(__defun__union, V2597, reg302099)
reg302101 := MakeSymbol("*property-vector*")
reg302102 := PrimValue(reg302101)
__ctx.TailApply(__defun__put, V2596, reg302091, reg302100, reg302102)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.record-internal", value: __defun__shen_4record_1internal})

__defun__shen_4internal_1symbols = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2608 := __args[0]
_ = V2608
V2609 := __args[1]
_ = V2609
reg302104 := PrimIsSymbol(V2609)
var reg302111 Obj
if reg302104 == True {
reg302105 := __e.Call(__defun__explode, V2609)
reg302106 := __e.Call(__defun__shen_4prefix_2, V2608, reg302105)
var reg302109 Obj
if reg302106 == True {
reg302107 := True;
reg302109 = reg302107
} else {
reg302108 := False;
reg302109 = reg302108
}
reg302111 = reg302109
} else {
reg302110 := False;
reg302111 = reg302110
}
if reg302111 == True {
reg302112 := Nil;
reg302113 := PrimCons(V2609, reg302112)
__ctx.Return(reg302113)
return
} else {
reg302114 := PrimIsPair(V2609)
if reg302114 == True {
reg302115 := PrimHead(V2609)
reg302116 := __e.Call(__defun__shen_4internal_1symbols, V2608, reg302115)
reg302117 := PrimTail(V2609)
reg302118 := __e.Call(__defun__shen_4internal_1symbols, V2608, reg302117)
__ctx.TailApply(__defun__union, reg302116, reg302118)
return
} else {
reg302120 := Nil;
__ctx.Return(reg302120)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.internal-symbols", value: __defun__shen_4internal_1symbols})

__defun__shen_4packageh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2626 := __args[0]
_ = V2626
V2627 := __args[1]
_ = V2627
V2628 := __args[2]
_ = V2628
V2629 := __args[3]
_ = V2629
reg302121 := PrimIsPair(V2628)
if reg302121 == True {
reg302122 := PrimHead(V2628)
reg302123 := __e.Call(__defun__shen_4packageh, V2626, V2627, reg302122, V2629)
reg302124 := PrimTail(V2628)
reg302125 := __e.Call(__defun__shen_4packageh, V2626, V2627, reg302124, V2629)
reg302126 := PrimCons(reg302123, reg302125)
__ctx.Return(reg302126)
return
} else {
reg302127 := __e.Call(__defun__shen_4sysfunc_2, V2628)
var reg302151 Obj
if reg302127 == True {
reg302128 := True;
reg302151 = reg302128
} else {
reg302129 := PrimIsVariable(V2628)
var reg302147 Obj
if reg302129 == True {
reg302130 := True;
reg302147 = reg302130
} else {
reg302131 := __e.Call(__defun__element_2, V2628, V2627)
var reg302143 Obj
if reg302131 == True {
reg302132 := True;
reg302143 = reg302132
} else {
reg302133 := __e.Call(__defun__shen_4doubleunderline_2, V2628)
var reg302139 Obj
if reg302133 == True {
reg302134 := True;
reg302139 = reg302134
} else {
reg302135 := __e.Call(__defun__shen_4singleunderline_2, V2628)
var reg302138 Obj
if reg302135 == True {
reg302136 := True;
reg302138 = reg302136
} else {
reg302137 := False;
reg302138 = reg302137
}
reg302139 = reg302138
}
var reg302142 Obj
if reg302139 == True {
reg302140 := True;
reg302142 = reg302140
} else {
reg302141 := False;
reg302142 = reg302141
}
reg302143 = reg302142
}
var reg302146 Obj
if reg302143 == True {
reg302144 := True;
reg302146 = reg302144
} else {
reg302145 := False;
reg302146 = reg302145
}
reg302147 = reg302146
}
var reg302150 Obj
if reg302147 == True {
reg302148 := True;
reg302150 = reg302148
} else {
reg302149 := False;
reg302150 = reg302149
}
reg302151 = reg302150
}
if reg302151 == True {
__ctx.Return(V2628)
return
} else {
reg302152 := PrimIsSymbol(V2628)
var reg302178 Obj
if reg302152 == True {
reg302153 := __e.Call(__defun__explode, V2628)
ExplodeX := reg302153
_ = ExplodeX
reg302154 := MakeString("s")
reg302155 := MakeString("h")
reg302156 := MakeString("e")
reg302157 := MakeString("n")
reg302158 := MakeString(".")
reg302159 := Nil;
reg302160 := PrimCons(reg302158, reg302159)
reg302161 := PrimCons(reg302157, reg302160)
reg302162 := PrimCons(reg302156, reg302161)
reg302163 := PrimCons(reg302155, reg302162)
reg302164 := PrimCons(reg302154, reg302163)
reg302165 := __e.Call(__defun__shen_4prefix_2, reg302164, ExplodeX)
reg302166 := PrimNot(reg302165)
var reg302173 Obj
if reg302166 == True {
reg302167 := __e.Call(__defun__shen_4prefix_2, V2629, ExplodeX)
reg302168 := PrimNot(reg302167)
var reg302171 Obj
if reg302168 == True {
reg302169 := True;
reg302171 = reg302169
} else {
reg302170 := False;
reg302171 = reg302170
}
reg302173 = reg302171
} else {
reg302172 := False;
reg302173 = reg302172
}
var reg302176 Obj
if reg302173 == True {
reg302174 := True;
reg302176 = reg302174
} else {
reg302175 := False;
reg302176 = reg302175
}
reg302178 = reg302176
} else {
reg302177 := False;
reg302178 = reg302177
}
if reg302178 == True {
__ctx.TailApply(__defun__concat, V2626, V2628)
return
} else {
__ctx.Return(V2628)
return
}
}
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.packageh", value: __defun__shen_4packageh})

}
