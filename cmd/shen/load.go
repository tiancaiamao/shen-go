package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__load Obj // load
var __defun__shen_4load_1help Obj // shen.load-help
var __defun__shen_4remove_1synonyms Obj // shen.remove-synonyms
var __defun__shen_4typecheck_1and_1load Obj // shen.typecheck-and-load
var __defun__shen_4typetable Obj // shen.typetable
var __defun__shen_4assumetype Obj // shen.assumetype
var __defun__shen_4unwind_1types Obj // shen.unwind-types
var __defun__shen_4remtype Obj // shen.remtype
var __defun__shen_4removetype Obj // shen.removetype
var __defun__shen_4_5sig_7rest_6 Obj // shen.<sig+rest>
var __defun__write_1to_1file Obj // write-to-file

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg307834 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg307834)
return
}, 0))
__defun__load = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1673 := __args[0]
_ = V1673
reg307835 := MakeSymbol("run")
reg307836 := PrimGetTime(reg307835)
Start := reg307836
_ = Start
reg307837 := MakeSymbol("shen.*tc*")
reg307838 := PrimValue(reg307837)
reg307839 := __e.Call(__defun__read_1file, V1673)
reg307840 := __e.Call(__defun__shen_4load_1help, reg307838, reg307839)
Result := reg307840
_ = Result
reg307841 := MakeSymbol("run")
reg307842 := PrimGetTime(reg307841)
Finish := reg307842
_ = Finish
reg307843 := PrimNumberSubtract(Finish, Start)
Time := reg307843
_ = Time
reg307844 := MakeString("\nrun time: ")
reg307845 := PrimStr(Time)
reg307846 := MakeString(" secs\n")
reg307847 := PrimStringConcat(reg307845, reg307846)
reg307848 := PrimStringConcat(reg307844, reg307847)
reg307849 := __e.Call(__defun__stoutput)
reg307850 := __e.Call(__defun__shen_4prhush, reg307848, reg307849)
Message := reg307850
_ = Message
Load := Result
_ = Load
reg307851 := MakeSymbol("shen.*tc*")
reg307852 := PrimValue(reg307851)
var reg307862 Obj
if reg307852 == True {
reg307853 := MakeString("\ntypechecked in ")
reg307854 := __e.Call(__defun__inferences)
reg307855 := MakeString(" inferences\n")
reg307856 := MakeSymbol("shen.a")
reg307857 := __e.Call(__defun__shen_4app, reg307854, reg307855, reg307856)
reg307858 := PrimStringConcat(reg307853, reg307857)
reg307859 := __e.Call(__defun__stoutput)
reg307860 := __e.Call(__defun__shen_4prhush, reg307858, reg307859)
reg307862 = reg307860
} else {
reg307861 := MakeSymbol("shen.skip")
reg307862 = reg307861
}
Infs := reg307862
_ = Infs
reg307863 := MakeSymbol("loaded")
__ctx.Return(reg307863)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "load", value: __defun__load})

__defun__shen_4load_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1680 := __args[0]
_ = V1680
V1681 := __args[1]
_ = V1681
reg307864 := False;
reg307865 := PrimEqual(reg307864, V1680)
if reg307865 == True {
reg307866 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg307867 := __e.Call(__defun__shen_4eval_1without_1macros, X)
reg307868 := MakeString("\n")
reg307869 := MakeSymbol("shen.s")
reg307870 := __e.Call(__defun__shen_4app, reg307867, reg307868, reg307869)
reg307871 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg307870, reg307871)
return
}, 1)
__ctx.TailApply(__defun__shen_4for_1each, reg307866, V1681)
return
} else {
reg307874 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4remove_1synonyms, X)
return
}, 1)
reg307876 := __e.Call(__defun__mapcan, reg307874, V1681)
RemoveSynonyms := reg307876
_ = RemoveSynonyms
reg307877 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4typetable, X)
return
}, 1)
reg307879 := __e.Call(__defun__mapcan, reg307877, RemoveSynonyms)
Table := reg307879
_ = Table
reg307880 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4assumetype, X)
return
}, 1)
reg307882 := __e.Call(__defun__shen_4for_1each, reg307880, Table)
Assume := reg307882
_ = Assume
reg307883 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg307884 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4typecheck_1and_1load, X)
return
}, 1)
__ctx.TailApply(__defun__shen_4for_1each, reg307884, RemoveSynonyms)
return
}, 0)
reg307887 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.TailApply(__defun__shen_4unwind_1types, E, Table)
return
}, 1)
reg307889 := __e.Try(reg307883).Catch(reg307887)
__ctx.Return(reg307889)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.load-help", value: __defun__shen_4load_1help})

__defun__shen_4remove_1synonyms = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1683 := __args[0]
_ = V1683
reg307890 := PrimIsPair(V1683)
var reg307898 Obj
if reg307890 == True {
reg307891 := MakeSymbol("shen.synonyms-help")
reg307892 := PrimHead(V1683)
reg307893 := PrimEqual(reg307891, reg307892)
var reg307896 Obj
if reg307893 == True {
reg307894 := True;
reg307896 = reg307894
} else {
reg307895 := False;
reg307896 = reg307895
}
reg307898 = reg307896
} else {
reg307897 := False;
reg307898 = reg307897
}
if reg307898 == True {
reg307899 := __e.Call(__defun__eval, V1683)
_ = reg307899
reg307900 := Nil;
__ctx.Return(reg307900)
return
} else {
reg307901 := Nil;
reg307902 := PrimCons(V1683, reg307901)
__ctx.Return(reg307902)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.remove-synonyms", value: __defun__shen_4remove_1synonyms})

__defun__shen_4typecheck_1and_1load = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1685 := __args[0]
_ = V1685
reg307903 := MakeNumber(1)
reg307904 := __e.Call(__defun__nl, reg307903)
_ = reg307904
reg307905 := MakeSymbol("A")
reg307906 := __e.Call(__defun__gensym, reg307905)
__ctx.TailApply(__defun__shen_4typecheck_1and_1evaluate, V1685, reg307906)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.typecheck-and-load", value: __defun__shen_4typecheck_1and_1load})

__defun__shen_4typetable = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1691 := __args[0]
_ = V1691
reg307908 := PrimIsPair(V1691)
var reg307923 Obj
if reg307908 == True {
reg307909 := MakeSymbol("define")
reg307910 := PrimHead(V1691)
reg307911 := PrimEqual(reg307909, reg307910)
var reg307918 Obj
if reg307911 == True {
reg307912 := PrimTail(V1691)
reg307913 := PrimIsPair(reg307912)
var reg307916 Obj
if reg307913 == True {
reg307914 := True;
reg307916 = reg307914
} else {
reg307915 := False;
reg307916 = reg307915
}
reg307918 = reg307916
} else {
reg307917 := False;
reg307918 = reg307917
}
var reg307921 Obj
if reg307918 == True {
reg307919 := True;
reg307921 = reg307919
} else {
reg307920 := False;
reg307921 = reg307920
}
reg307923 = reg307921
} else {
reg307922 := False;
reg307923 = reg307922
}
if reg307923 == True {
reg307924 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4_5sig_7rest_6, Y)
return
}, 1)
reg307926 := PrimTail(V1691)
reg307927 := PrimTail(reg307926)
reg307928 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg307929 := PrimTail(V1691)
reg307930 := PrimHead(reg307929)
reg307931 := MakeString(" lacks a proper signature.\n")
reg307932 := MakeSymbol("shen.a")
reg307933 := __e.Call(__defun__shen_4app, reg307930, reg307931, reg307932)
reg307934 := PrimSimpleError(reg307933)
__ctx.Return(reg307934)
return
}, 1)
reg307935 := __e.Call(__defun__compile, reg307924, reg307927, reg307928)
Sig := reg307935
_ = Sig
reg307936 := PrimTail(V1691)
reg307937 := PrimHead(reg307936)
reg307938 := PrimCons(reg307937, Sig)
reg307939 := Nil;
reg307940 := PrimCons(reg307938, reg307939)
__ctx.Return(reg307940)
return
} else {
reg307941 := Nil;
__ctx.Return(reg307941)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.typetable", value: __defun__shen_4typetable})

__defun__shen_4assumetype = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1693 := __args[0]
_ = V1693
reg307942 := PrimIsPair(V1693)
if reg307942 == True {
reg307943 := PrimHead(V1693)
reg307944 := PrimTail(V1693)
__ctx.TailApply(__defun__declare, reg307943, reg307944)
return
} else {
reg307946 := MakeSymbol("shen.assumetype")
__ctx.TailApply(__defun__shen_4f__error, reg307946)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.assumetype", value: __defun__shen_4assumetype})

__defun__shen_4unwind_1types = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1700 := __args[0]
_ = V1700
V1701 := __args[1]
_ = V1701
reg307948 := Nil;
reg307949 := PrimEqual(reg307948, V1701)
if reg307949 == True {
reg307950 := PrimErrorToString(V1700)
reg307951 := PrimSimpleError(reg307950)
__ctx.Return(reg307951)
return
} else {
reg307952 := PrimIsPair(V1701)
var reg307959 Obj
if reg307952 == True {
reg307953 := PrimHead(V1701)
reg307954 := PrimIsPair(reg307953)
var reg307957 Obj
if reg307954 == True {
reg307955 := True;
reg307957 = reg307955
} else {
reg307956 := False;
reg307957 = reg307956
}
reg307959 = reg307957
} else {
reg307958 := False;
reg307959 = reg307958
}
if reg307959 == True {
reg307960 := PrimHead(V1701)
reg307961 := PrimHead(reg307960)
reg307962 := __e.Call(__defun__shen_4remtype, reg307961)
_ = reg307962
reg307963 := PrimTail(V1701)
__ctx.TailApply(__defun__shen_4unwind_1types, V1700, reg307963)
return
} else {
reg307965 := MakeSymbol("shen.unwind-types")
__ctx.TailApply(__defun__shen_4f__error, reg307965)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.unwind-types", value: __defun__shen_4unwind_1types})

__defun__shen_4remtype = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1703 := __args[0]
_ = V1703
reg307967 := MakeSymbol("shen.*signedfuncs*")
reg307968 := MakeSymbol("shen.*signedfuncs*")
reg307969 := PrimValue(reg307968)
reg307970 := __e.Call(__defun__shen_4removetype, V1703, reg307969)
reg307971 := PrimSet(reg307967, reg307970)
__ctx.Return(reg307971)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.remtype", value: __defun__shen_4remtype})

__defun__shen_4removetype = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1711 := __args[0]
_ = V1711
V1712 := __args[1]
_ = V1712
reg307972 := Nil;
reg307973 := PrimEqual(reg307972, V1712)
if reg307973 == True {
reg307974 := Nil;
__ctx.Return(reg307974)
return
} else {
reg307975 := PrimIsPair(V1712)
var reg307990 Obj
if reg307975 == True {
reg307976 := PrimHead(V1712)
reg307977 := PrimIsPair(reg307976)
var reg307985 Obj
if reg307977 == True {
reg307978 := PrimHead(V1712)
reg307979 := PrimHead(reg307978)
reg307980 := PrimEqual(reg307979, V1711)
var reg307983 Obj
if reg307980 == True {
reg307981 := True;
reg307983 = reg307981
} else {
reg307982 := False;
reg307983 = reg307982
}
reg307985 = reg307983
} else {
reg307984 := False;
reg307985 = reg307984
}
var reg307988 Obj
if reg307985 == True {
reg307986 := True;
reg307988 = reg307986
} else {
reg307987 := False;
reg307988 = reg307987
}
reg307990 = reg307988
} else {
reg307989 := False;
reg307990 = reg307989
}
if reg307990 == True {
reg307991 := PrimHead(V1712)
reg307992 := PrimHead(reg307991)
reg307993 := PrimTail(V1712)
__ctx.TailApply(__defun__shen_4removetype, reg307992, reg307993)
return
} else {
reg307995 := PrimIsPair(V1712)
if reg307995 == True {
reg307996 := PrimHead(V1712)
reg307997 := PrimTail(V1712)
reg307998 := __e.Call(__defun__shen_4removetype, V1711, reg307997)
reg307999 := PrimCons(reg307996, reg307998)
__ctx.Return(reg307999)
return
} else {
reg308000 := MakeSymbol("shen.removetype")
__ctx.TailApply(__defun__shen_4f__error, reg308000)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.removetype", value: __defun__shen_4removetype})

__defun__shen_4_5sig_7rest_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1714 := __args[0]
_ = V1714
reg308002 := __e.Call(__defun__shen_4_5signature_6, V1714)
Parse__shen_4_5signature_6 := reg308002
_ = Parse__shen_4_5signature_6
reg308003 := __e.Call(__defun__fail)
reg308004 := PrimEqual(reg308003, Parse__shen_4_5signature_6)
reg308005 := PrimNot(reg308004)
if reg308005 == True {
reg308006 := __e.Call(__defun___5_b_6, Parse__shen_4_5signature_6)
Parse___5_b_6 := reg308006
_ = Parse___5_b_6
reg308007 := __e.Call(__defun__fail)
reg308008 := PrimEqual(reg308007, Parse___5_b_6)
reg308009 := PrimNot(reg308008)
if reg308009 == True {
reg308010 := PrimHead(Parse___5_b_6)
reg308011 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_6)
__ctx.TailApply(__defun__shen_4pair, reg308010, reg308011)
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
__initDefs = append(__initDefs, defType{name: "shen.<sig+rest>", value: __defun__shen_4_5sig_7rest_6})

__defun__write_1to_1file = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1717 := __args[0]
_ = V1717
V1718 := __args[1]
_ = V1718
reg308015 := MakeSymbol("out")
reg308016 := PrimOpenStream(V1717, reg308015)
Stream := reg308016
_ = Stream
reg308017 := PrimIsString(V1718)
var reg308024 Obj
if reg308017 == True {
reg308018 := MakeString("\n\n")
reg308019 := MakeSymbol("shen.a")
reg308020 := __e.Call(__defun__shen_4app, V1718, reg308018, reg308019)
reg308024 = reg308020
} else {
reg308021 := MakeString("\n\n")
reg308022 := MakeSymbol("shen.s")
reg308023 := __e.Call(__defun__shen_4app, V1718, reg308021, reg308022)
reg308024 = reg308023
}
String := reg308024
_ = String
reg308025 := __e.Call(__defun__pr, String, Stream)
Write := reg308025
_ = Write
reg308026 := PrimCloseStream(Stream)
Close := reg308026
_ = Close
__ctx.Return(V1718)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "write-to-file", value: __defun__write_1to_1file})

}
