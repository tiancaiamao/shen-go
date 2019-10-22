package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__macroexpand Obj // macroexpand
var __defun__shen_4error_1macro Obj // shen.error-macro
var __defun__shen_4output_1macro Obj // shen.output-macro
var __defun__shen_4make_1string_1macro Obj // shen.make-string-macro
var __defun__shen_4input_1macro Obj // shen.input-macro
var __defun__shen_4compose Obj // shen.compose
var __defun__shen_4compile_1macro Obj // shen.compile-macro
var __defun__shen_4prolog_1macro Obj // shen.prolog-macro
var __defun__shen_4receive_1terms Obj // shen.receive-terms
var __defun__shen_4pass_1literals Obj // shen.pass-literals
var __defun__shen_4defprolog_1macro Obj // shen.defprolog-macro
var __defun__shen_4datatype_1macro Obj // shen.datatype-macro
var __defun__shen_4intern_1type Obj // shen.intern-type
var __defun__shen_4_8s_1macro Obj // shen.@s-macro
var __defun__shen_4synonyms_1macro Obj // shen.synonyms-macro
var __defun__shen_4curry_1synonyms Obj // shen.curry-synonyms
var __defun__shen_4nl_1macro Obj // shen.nl-macro
var __defun__shen_4assoc_1macro Obj // shen.assoc-macro
var __defun__shen_4let_1macro Obj // shen.let-macro
var __defun__shen_4abs_1macro Obj // shen.abs-macro
var __defun__shen_4cases_1macro Obj // shen.cases-macro
var __defun__shen_4timer_1macro Obj // shen.timer-macro
var __defun__shen_4tuple_1up Obj // shen.tuple-up
var __defun__shen_4put_cget_1macro Obj // shen.put/get-macro
var __defun__shen_4function_1macro Obj // shen.function-macro
var __defun__shen_4function_1abstraction Obj // shen.function-abstraction
var __defun__shen_4function_1abstraction_1help Obj // shen.function-abstraction-help
var __defun__undefmacro Obj // undefmacro
var __defun__shen_4findpos Obj // shen.findpos
var __defun__shen_4remove_1nth Obj // shen.remove-nth

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg308762 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg308762)
return
}, 0))
__defun__macroexpand = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1720 := __args[0]
_ = V1720
reg308763 := MakeSymbol("*macros*")
reg308764 := PrimValue(reg308763)
reg308765 := __e.Call(__defun__shen_4compose, reg308764, V1720)
Y := reg308765
_ = Y
reg308766 := PrimEqual(V1720, Y)
if reg308766 == True {
__ctx.Return(V1720)
return
} else {
reg308767 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__macroexpand, Z)
return
}, 1)
__ctx.TailApply(__defun__shen_4walk, reg308767, Y)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "macroexpand", value: __defun__macroexpand})

__defun__shen_4error_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1722 := __args[0]
_ = V1722
reg308770 := PrimIsPair(V1722)
var reg308785 Obj
if reg308770 == True {
reg308771 := MakeSymbol("error")
reg308772 := PrimHead(V1722)
reg308773 := PrimEqual(reg308771, reg308772)
var reg308780 Obj
if reg308773 == True {
reg308774 := PrimTail(V1722)
reg308775 := PrimIsPair(reg308774)
var reg308778 Obj
if reg308775 == True {
reg308776 := True;
reg308778 = reg308776
} else {
reg308777 := False;
reg308778 = reg308777
}
reg308780 = reg308778
} else {
reg308779 := False;
reg308780 = reg308779
}
var reg308783 Obj
if reg308780 == True {
reg308781 := True;
reg308783 = reg308781
} else {
reg308782 := False;
reg308783 = reg308782
}
reg308785 = reg308783
} else {
reg308784 := False;
reg308785 = reg308784
}
if reg308785 == True {
reg308786 := MakeSymbol("simple-error")
reg308787 := PrimTail(V1722)
reg308788 := PrimHead(reg308787)
reg308789 := PrimTail(V1722)
reg308790 := PrimTail(reg308789)
reg308791 := __e.Call(__defun__shen_4mkstr, reg308788, reg308790)
reg308792 := Nil;
reg308793 := PrimCons(reg308791, reg308792)
reg308794 := PrimCons(reg308786, reg308793)
__ctx.Return(reg308794)
return
} else {
__ctx.Return(V1722)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.error-macro", value: __defun__shen_4error_1macro})

__defun__shen_4output_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1724 := __args[0]
_ = V1724
reg308795 := PrimIsPair(V1724)
var reg308810 Obj
if reg308795 == True {
reg308796 := MakeSymbol("output")
reg308797 := PrimHead(V1724)
reg308798 := PrimEqual(reg308796, reg308797)
var reg308805 Obj
if reg308798 == True {
reg308799 := PrimTail(V1724)
reg308800 := PrimIsPair(reg308799)
var reg308803 Obj
if reg308800 == True {
reg308801 := True;
reg308803 = reg308801
} else {
reg308802 := False;
reg308803 = reg308802
}
reg308805 = reg308803
} else {
reg308804 := False;
reg308805 = reg308804
}
var reg308808 Obj
if reg308805 == True {
reg308806 := True;
reg308808 = reg308806
} else {
reg308807 := False;
reg308808 = reg308807
}
reg308810 = reg308808
} else {
reg308809 := False;
reg308810 = reg308809
}
if reg308810 == True {
reg308811 := MakeSymbol("shen.prhush")
reg308812 := PrimTail(V1724)
reg308813 := PrimHead(reg308812)
reg308814 := PrimTail(V1724)
reg308815 := PrimTail(reg308814)
reg308816 := __e.Call(__defun__shen_4mkstr, reg308813, reg308815)
reg308817 := MakeSymbol("stoutput")
reg308818 := Nil;
reg308819 := PrimCons(reg308817, reg308818)
reg308820 := Nil;
reg308821 := PrimCons(reg308819, reg308820)
reg308822 := PrimCons(reg308816, reg308821)
reg308823 := PrimCons(reg308811, reg308822)
__ctx.Return(reg308823)
return
} else {
reg308824 := PrimIsPair(V1724)
var reg308848 Obj
if reg308824 == True {
reg308825 := MakeSymbol("pr")
reg308826 := PrimHead(V1724)
reg308827 := PrimEqual(reg308825, reg308826)
var reg308843 Obj
if reg308827 == True {
reg308828 := PrimTail(V1724)
reg308829 := PrimIsPair(reg308828)
var reg308838 Obj
if reg308829 == True {
reg308830 := Nil;
reg308831 := PrimTail(V1724)
reg308832 := PrimTail(reg308831)
reg308833 := PrimEqual(reg308830, reg308832)
var reg308836 Obj
if reg308833 == True {
reg308834 := True;
reg308836 = reg308834
} else {
reg308835 := False;
reg308836 = reg308835
}
reg308838 = reg308836
} else {
reg308837 := False;
reg308838 = reg308837
}
var reg308841 Obj
if reg308838 == True {
reg308839 := True;
reg308841 = reg308839
} else {
reg308840 := False;
reg308841 = reg308840
}
reg308843 = reg308841
} else {
reg308842 := False;
reg308843 = reg308842
}
var reg308846 Obj
if reg308843 == True {
reg308844 := True;
reg308846 = reg308844
} else {
reg308845 := False;
reg308846 = reg308845
}
reg308848 = reg308846
} else {
reg308847 := False;
reg308848 = reg308847
}
if reg308848 == True {
reg308849 := MakeSymbol("pr")
reg308850 := PrimTail(V1724)
reg308851 := PrimHead(reg308850)
reg308852 := MakeSymbol("stoutput")
reg308853 := Nil;
reg308854 := PrimCons(reg308852, reg308853)
reg308855 := Nil;
reg308856 := PrimCons(reg308854, reg308855)
reg308857 := PrimCons(reg308851, reg308856)
reg308858 := PrimCons(reg308849, reg308857)
__ctx.Return(reg308858)
return
} else {
__ctx.Return(V1724)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.output-macro", value: __defun__shen_4output_1macro})

__defun__shen_4make_1string_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1726 := __args[0]
_ = V1726
reg308859 := PrimIsPair(V1726)
var reg308874 Obj
if reg308859 == True {
reg308860 := MakeSymbol("make-string")
reg308861 := PrimHead(V1726)
reg308862 := PrimEqual(reg308860, reg308861)
var reg308869 Obj
if reg308862 == True {
reg308863 := PrimTail(V1726)
reg308864 := PrimIsPair(reg308863)
var reg308867 Obj
if reg308864 == True {
reg308865 := True;
reg308867 = reg308865
} else {
reg308866 := False;
reg308867 = reg308866
}
reg308869 = reg308867
} else {
reg308868 := False;
reg308869 = reg308868
}
var reg308872 Obj
if reg308869 == True {
reg308870 := True;
reg308872 = reg308870
} else {
reg308871 := False;
reg308872 = reg308871
}
reg308874 = reg308872
} else {
reg308873 := False;
reg308874 = reg308873
}
if reg308874 == True {
reg308875 := PrimTail(V1726)
reg308876 := PrimHead(reg308875)
reg308877 := PrimTail(V1726)
reg308878 := PrimTail(reg308877)
__ctx.TailApply(__defun__shen_4mkstr, reg308876, reg308878)
return
} else {
__ctx.Return(V1726)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.make-string-macro", value: __defun__shen_4make_1string_1macro})

__defun__shen_4input_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1728 := __args[0]
_ = V1728
reg308880 := PrimIsPair(V1728)
var reg308896 Obj
if reg308880 == True {
reg308881 := MakeSymbol("lineread")
reg308882 := PrimHead(V1728)
reg308883 := PrimEqual(reg308881, reg308882)
var reg308891 Obj
if reg308883 == True {
reg308884 := Nil;
reg308885 := PrimTail(V1728)
reg308886 := PrimEqual(reg308884, reg308885)
var reg308889 Obj
if reg308886 == True {
reg308887 := True;
reg308889 = reg308887
} else {
reg308888 := False;
reg308889 = reg308888
}
reg308891 = reg308889
} else {
reg308890 := False;
reg308891 = reg308890
}
var reg308894 Obj
if reg308891 == True {
reg308892 := True;
reg308894 = reg308892
} else {
reg308893 := False;
reg308894 = reg308893
}
reg308896 = reg308894
} else {
reg308895 := False;
reg308896 = reg308895
}
if reg308896 == True {
reg308897 := MakeSymbol("lineread")
reg308898 := MakeSymbol("stinput")
reg308899 := Nil;
reg308900 := PrimCons(reg308898, reg308899)
reg308901 := Nil;
reg308902 := PrimCons(reg308900, reg308901)
reg308903 := PrimCons(reg308897, reg308902)
__ctx.Return(reg308903)
return
} else {
reg308904 := PrimIsPair(V1728)
var reg308920 Obj
if reg308904 == True {
reg308905 := MakeSymbol("input")
reg308906 := PrimHead(V1728)
reg308907 := PrimEqual(reg308905, reg308906)
var reg308915 Obj
if reg308907 == True {
reg308908 := Nil;
reg308909 := PrimTail(V1728)
reg308910 := PrimEqual(reg308908, reg308909)
var reg308913 Obj
if reg308910 == True {
reg308911 := True;
reg308913 = reg308911
} else {
reg308912 := False;
reg308913 = reg308912
}
reg308915 = reg308913
} else {
reg308914 := False;
reg308915 = reg308914
}
var reg308918 Obj
if reg308915 == True {
reg308916 := True;
reg308918 = reg308916
} else {
reg308917 := False;
reg308918 = reg308917
}
reg308920 = reg308918
} else {
reg308919 := False;
reg308920 = reg308919
}
if reg308920 == True {
reg308921 := MakeSymbol("input")
reg308922 := MakeSymbol("stinput")
reg308923 := Nil;
reg308924 := PrimCons(reg308922, reg308923)
reg308925 := Nil;
reg308926 := PrimCons(reg308924, reg308925)
reg308927 := PrimCons(reg308921, reg308926)
__ctx.Return(reg308927)
return
} else {
reg308928 := PrimIsPair(V1728)
var reg308944 Obj
if reg308928 == True {
reg308929 := MakeSymbol("read")
reg308930 := PrimHead(V1728)
reg308931 := PrimEqual(reg308929, reg308930)
var reg308939 Obj
if reg308931 == True {
reg308932 := Nil;
reg308933 := PrimTail(V1728)
reg308934 := PrimEqual(reg308932, reg308933)
var reg308937 Obj
if reg308934 == True {
reg308935 := True;
reg308937 = reg308935
} else {
reg308936 := False;
reg308937 = reg308936
}
reg308939 = reg308937
} else {
reg308938 := False;
reg308939 = reg308938
}
var reg308942 Obj
if reg308939 == True {
reg308940 := True;
reg308942 = reg308940
} else {
reg308941 := False;
reg308942 = reg308941
}
reg308944 = reg308942
} else {
reg308943 := False;
reg308944 = reg308943
}
if reg308944 == True {
reg308945 := MakeSymbol("read")
reg308946 := MakeSymbol("stinput")
reg308947 := Nil;
reg308948 := PrimCons(reg308946, reg308947)
reg308949 := Nil;
reg308950 := PrimCons(reg308948, reg308949)
reg308951 := PrimCons(reg308945, reg308950)
__ctx.Return(reg308951)
return
} else {
reg308952 := PrimIsPair(V1728)
var reg308976 Obj
if reg308952 == True {
reg308953 := MakeSymbol("input+")
reg308954 := PrimHead(V1728)
reg308955 := PrimEqual(reg308953, reg308954)
var reg308971 Obj
if reg308955 == True {
reg308956 := PrimTail(V1728)
reg308957 := PrimIsPair(reg308956)
var reg308966 Obj
if reg308957 == True {
reg308958 := Nil;
reg308959 := PrimTail(V1728)
reg308960 := PrimTail(reg308959)
reg308961 := PrimEqual(reg308958, reg308960)
var reg308964 Obj
if reg308961 == True {
reg308962 := True;
reg308964 = reg308962
} else {
reg308963 := False;
reg308964 = reg308963
}
reg308966 = reg308964
} else {
reg308965 := False;
reg308966 = reg308965
}
var reg308969 Obj
if reg308966 == True {
reg308967 := True;
reg308969 = reg308967
} else {
reg308968 := False;
reg308969 = reg308968
}
reg308971 = reg308969
} else {
reg308970 := False;
reg308971 = reg308970
}
var reg308974 Obj
if reg308971 == True {
reg308972 := True;
reg308974 = reg308972
} else {
reg308973 := False;
reg308974 = reg308973
}
reg308976 = reg308974
} else {
reg308975 := False;
reg308976 = reg308975
}
if reg308976 == True {
reg308977 := MakeSymbol("input+")
reg308978 := PrimTail(V1728)
reg308979 := PrimHead(reg308978)
reg308980 := MakeSymbol("stinput")
reg308981 := Nil;
reg308982 := PrimCons(reg308980, reg308981)
reg308983 := Nil;
reg308984 := PrimCons(reg308982, reg308983)
reg308985 := PrimCons(reg308979, reg308984)
reg308986 := PrimCons(reg308977, reg308985)
__ctx.Return(reg308986)
return
} else {
reg308987 := PrimIsPair(V1728)
var reg309003 Obj
if reg308987 == True {
reg308988 := MakeSymbol("read-byte")
reg308989 := PrimHead(V1728)
reg308990 := PrimEqual(reg308988, reg308989)
var reg308998 Obj
if reg308990 == True {
reg308991 := Nil;
reg308992 := PrimTail(V1728)
reg308993 := PrimEqual(reg308991, reg308992)
var reg308996 Obj
if reg308993 == True {
reg308994 := True;
reg308996 = reg308994
} else {
reg308995 := False;
reg308996 = reg308995
}
reg308998 = reg308996
} else {
reg308997 := False;
reg308998 = reg308997
}
var reg309001 Obj
if reg308998 == True {
reg308999 := True;
reg309001 = reg308999
} else {
reg309000 := False;
reg309001 = reg309000
}
reg309003 = reg309001
} else {
reg309002 := False;
reg309003 = reg309002
}
if reg309003 == True {
reg309004 := MakeSymbol("read-byte")
reg309005 := MakeSymbol("stinput")
reg309006 := Nil;
reg309007 := PrimCons(reg309005, reg309006)
reg309008 := Nil;
reg309009 := PrimCons(reg309007, reg309008)
reg309010 := PrimCons(reg309004, reg309009)
__ctx.Return(reg309010)
return
} else {
__ctx.Return(V1728)
return
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.input-macro", value: __defun__shen_4input_1macro})

__defun__shen_4compose = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1731 := __args[0]
_ = V1731
V1732 := __args[1]
_ = V1732
reg309012 := Nil;
reg309013 := PrimEqual(reg309012, V1731)
if reg309013 == True {
__ctx.Return(V1732)
return
} else {
reg309014 := PrimIsPair(V1731)
if reg309014 == True {
reg309015 := PrimTail(V1731)
reg309016 := PrimHead(V1731)
f309011 := reg309016
_ = f309011
reg309017 := __e.Call(f309011, V1732)
__ctx.TailApply(__defun__shen_4compose, reg309015, reg309017)
return
} else {
reg309019 := MakeSymbol("shen.compose")
__ctx.TailApply(__defun__shen_4f__error, reg309019)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compose", value: __defun__shen_4compose})

__defun__shen_4compile_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1734 := __args[0]
_ = V1734
reg309021 := PrimIsPair(V1734)
var reg309054 Obj
if reg309021 == True {
reg309022 := MakeSymbol("compile")
reg309023 := PrimHead(V1734)
reg309024 := PrimEqual(reg309022, reg309023)
var reg309049 Obj
if reg309024 == True {
reg309025 := PrimTail(V1734)
reg309026 := PrimIsPair(reg309025)
var reg309044 Obj
if reg309026 == True {
reg309027 := PrimTail(V1734)
reg309028 := PrimTail(reg309027)
reg309029 := PrimIsPair(reg309028)
var reg309039 Obj
if reg309029 == True {
reg309030 := Nil;
reg309031 := PrimTail(V1734)
reg309032 := PrimTail(reg309031)
reg309033 := PrimTail(reg309032)
reg309034 := PrimEqual(reg309030, reg309033)
var reg309037 Obj
if reg309034 == True {
reg309035 := True;
reg309037 = reg309035
} else {
reg309036 := False;
reg309037 = reg309036
}
reg309039 = reg309037
} else {
reg309038 := False;
reg309039 = reg309038
}
var reg309042 Obj
if reg309039 == True {
reg309040 := True;
reg309042 = reg309040
} else {
reg309041 := False;
reg309042 = reg309041
}
reg309044 = reg309042
} else {
reg309043 := False;
reg309044 = reg309043
}
var reg309047 Obj
if reg309044 == True {
reg309045 := True;
reg309047 = reg309045
} else {
reg309046 := False;
reg309047 = reg309046
}
reg309049 = reg309047
} else {
reg309048 := False;
reg309049 = reg309048
}
var reg309052 Obj
if reg309049 == True {
reg309050 := True;
reg309052 = reg309050
} else {
reg309051 := False;
reg309052 = reg309051
}
reg309054 = reg309052
} else {
reg309053 := False;
reg309054 = reg309053
}
if reg309054 == True {
reg309055 := MakeSymbol("compile")
reg309056 := PrimTail(V1734)
reg309057 := PrimHead(reg309056)
reg309058 := PrimTail(V1734)
reg309059 := PrimTail(reg309058)
reg309060 := PrimHead(reg309059)
reg309061 := MakeSymbol("lambda")
reg309062 := MakeSymbol("E")
reg309063 := MakeSymbol("if")
reg309064 := MakeSymbol("cons?")
reg309065 := MakeSymbol("E")
reg309066 := Nil;
reg309067 := PrimCons(reg309065, reg309066)
reg309068 := PrimCons(reg309064, reg309067)
reg309069 := MakeSymbol("error")
reg309070 := MakeString("parse error here: ~S~%")
reg309071 := MakeSymbol("E")
reg309072 := Nil;
reg309073 := PrimCons(reg309071, reg309072)
reg309074 := PrimCons(reg309070, reg309073)
reg309075 := PrimCons(reg309069, reg309074)
reg309076 := MakeSymbol("error")
reg309077 := MakeString("parse error~%")
reg309078 := Nil;
reg309079 := PrimCons(reg309077, reg309078)
reg309080 := PrimCons(reg309076, reg309079)
reg309081 := Nil;
reg309082 := PrimCons(reg309080, reg309081)
reg309083 := PrimCons(reg309075, reg309082)
reg309084 := PrimCons(reg309068, reg309083)
reg309085 := PrimCons(reg309063, reg309084)
reg309086 := Nil;
reg309087 := PrimCons(reg309085, reg309086)
reg309088 := PrimCons(reg309062, reg309087)
reg309089 := PrimCons(reg309061, reg309088)
reg309090 := Nil;
reg309091 := PrimCons(reg309089, reg309090)
reg309092 := PrimCons(reg309060, reg309091)
reg309093 := PrimCons(reg309057, reg309092)
reg309094 := PrimCons(reg309055, reg309093)
__ctx.Return(reg309094)
return
} else {
__ctx.Return(V1734)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.compile-macro", value: __defun__shen_4compile_1macro})

__defun__shen_4prolog_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1736 := __args[0]
_ = V1736
reg309095 := PrimIsPair(V1736)
var reg309103 Obj
if reg309095 == True {
reg309096 := MakeSymbol("prolog?")
reg309097 := PrimHead(V1736)
reg309098 := PrimEqual(reg309096, reg309097)
var reg309101 Obj
if reg309098 == True {
reg309099 := True;
reg309101 = reg309099
} else {
reg309100 := False;
reg309101 = reg309100
}
reg309103 = reg309101
} else {
reg309102 := False;
reg309103 = reg309102
}
if reg309103 == True {
reg309104 := MakeSymbol("shen.f")
reg309105 := __e.Call(__defun__gensym, reg309104)
F := reg309105
_ = F
reg309106 := PrimTail(V1736)
reg309107 := __e.Call(__defun__shen_4receive_1terms, reg309106)
Receive := reg309107
_ = Receive
reg309108 := MakeSymbol("defprolog")
reg309109 := Nil;
reg309110 := PrimCons(F, reg309109)
reg309111 := PrimCons(reg309108, reg309110)
reg309112 := MakeSymbol("<--")
reg309113 := Nil;
reg309114 := PrimCons(reg309112, reg309113)
reg309115 := PrimTail(V1736)
reg309116 := __e.Call(__defun__shen_4pass_1literals, reg309115)
reg309117 := MakeSymbol(";")
reg309118 := Nil;
reg309119 := PrimCons(reg309117, reg309118)
reg309120 := __e.Call(__defun__append, reg309116, reg309119)
reg309121 := __e.Call(__defun__append, reg309114, reg309120)
reg309122 := __e.Call(__defun__append, Receive, reg309121)
reg309123 := __e.Call(__defun__append, reg309111, reg309122)
reg309124 := __e.Call(__defun__eval, reg309123)
PrologDef := reg309124
_ = PrologDef
reg309125 := MakeSymbol("shen.start-new-prolog-process")
reg309126 := Nil;
reg309127 := PrimCons(reg309125, reg309126)
reg309128 := MakeSymbol("freeze")
reg309129 := True;
reg309130 := Nil;
reg309131 := PrimCons(reg309129, reg309130)
reg309132 := PrimCons(reg309128, reg309131)
reg309133 := Nil;
reg309134 := PrimCons(reg309132, reg309133)
reg309135 := PrimCons(reg309127, reg309134)
reg309136 := __e.Call(__defun__append, Receive, reg309135)
reg309137 := PrimCons(F, reg309136)
Query := reg309137
_ = Query
__ctx.Return(Query)
return
} else {
__ctx.Return(V1736)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.prolog-macro", value: __defun__shen_4prolog_1macro})

__defun__shen_4receive_1terms = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1742 := __args[0]
_ = V1742
reg309138 := Nil;
reg309139 := PrimEqual(reg309138, V1742)
if reg309139 == True {
reg309140 := Nil;
__ctx.Return(reg309140)
return
} else {
reg309141 := PrimIsPair(V1742)
var reg309175 Obj
if reg309141 == True {
reg309142 := PrimHead(V1742)
reg309143 := PrimIsPair(reg309142)
var reg309170 Obj
if reg309143 == True {
reg309144 := MakeSymbol("receive")
reg309145 := PrimHead(V1742)
reg309146 := PrimHead(reg309145)
reg309147 := PrimEqual(reg309144, reg309146)
var reg309165 Obj
if reg309147 == True {
reg309148 := PrimHead(V1742)
reg309149 := PrimTail(reg309148)
reg309150 := PrimIsPair(reg309149)
var reg309160 Obj
if reg309150 == True {
reg309151 := Nil;
reg309152 := PrimHead(V1742)
reg309153 := PrimTail(reg309152)
reg309154 := PrimTail(reg309153)
reg309155 := PrimEqual(reg309151, reg309154)
var reg309158 Obj
if reg309155 == True {
reg309156 := True;
reg309158 = reg309156
} else {
reg309157 := False;
reg309158 = reg309157
}
reg309160 = reg309158
} else {
reg309159 := False;
reg309160 = reg309159
}
var reg309163 Obj
if reg309160 == True {
reg309161 := True;
reg309163 = reg309161
} else {
reg309162 := False;
reg309163 = reg309162
}
reg309165 = reg309163
} else {
reg309164 := False;
reg309165 = reg309164
}
var reg309168 Obj
if reg309165 == True {
reg309166 := True;
reg309168 = reg309166
} else {
reg309167 := False;
reg309168 = reg309167
}
reg309170 = reg309168
} else {
reg309169 := False;
reg309170 = reg309169
}
var reg309173 Obj
if reg309170 == True {
reg309171 := True;
reg309173 = reg309171
} else {
reg309172 := False;
reg309173 = reg309172
}
reg309175 = reg309173
} else {
reg309174 := False;
reg309175 = reg309174
}
if reg309175 == True {
reg309176 := PrimHead(V1742)
reg309177 := PrimTail(reg309176)
reg309178 := PrimHead(reg309177)
reg309179 := PrimTail(V1742)
reg309180 := __e.Call(__defun__shen_4receive_1terms, reg309179)
reg309181 := PrimCons(reg309178, reg309180)
__ctx.Return(reg309181)
return
} else {
reg309182 := PrimIsPair(V1742)
if reg309182 == True {
reg309183 := PrimTail(V1742)
__ctx.TailApply(__defun__shen_4receive_1terms, reg309183)
return
} else {
reg309185 := MakeSymbol("shen.receive-terms")
__ctx.TailApply(__defun__shen_4f__error, reg309185)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.receive-terms", value: __defun__shen_4receive_1terms})

__defun__shen_4pass_1literals = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1746 := __args[0]
_ = V1746
reg309187 := Nil;
reg309188 := PrimEqual(reg309187, V1746)
if reg309188 == True {
reg309189 := Nil;
__ctx.Return(reg309189)
return
} else {
reg309190 := PrimIsPair(V1746)
var reg309224 Obj
if reg309190 == True {
reg309191 := PrimHead(V1746)
reg309192 := PrimIsPair(reg309191)
var reg309219 Obj
if reg309192 == True {
reg309193 := MakeSymbol("receive")
reg309194 := PrimHead(V1746)
reg309195 := PrimHead(reg309194)
reg309196 := PrimEqual(reg309193, reg309195)
var reg309214 Obj
if reg309196 == True {
reg309197 := PrimHead(V1746)
reg309198 := PrimTail(reg309197)
reg309199 := PrimIsPair(reg309198)
var reg309209 Obj
if reg309199 == True {
reg309200 := Nil;
reg309201 := PrimHead(V1746)
reg309202 := PrimTail(reg309201)
reg309203 := PrimTail(reg309202)
reg309204 := PrimEqual(reg309200, reg309203)
var reg309207 Obj
if reg309204 == True {
reg309205 := True;
reg309207 = reg309205
} else {
reg309206 := False;
reg309207 = reg309206
}
reg309209 = reg309207
} else {
reg309208 := False;
reg309209 = reg309208
}
var reg309212 Obj
if reg309209 == True {
reg309210 := True;
reg309212 = reg309210
} else {
reg309211 := False;
reg309212 = reg309211
}
reg309214 = reg309212
} else {
reg309213 := False;
reg309214 = reg309213
}
var reg309217 Obj
if reg309214 == True {
reg309215 := True;
reg309217 = reg309215
} else {
reg309216 := False;
reg309217 = reg309216
}
reg309219 = reg309217
} else {
reg309218 := False;
reg309219 = reg309218
}
var reg309222 Obj
if reg309219 == True {
reg309220 := True;
reg309222 = reg309220
} else {
reg309221 := False;
reg309222 = reg309221
}
reg309224 = reg309222
} else {
reg309223 := False;
reg309224 = reg309223
}
if reg309224 == True {
reg309225 := PrimTail(V1746)
__ctx.TailApply(__defun__shen_4pass_1literals, reg309225)
return
} else {
reg309227 := PrimIsPair(V1746)
if reg309227 == True {
reg309228 := PrimHead(V1746)
reg309229 := PrimTail(V1746)
reg309230 := __e.Call(__defun__shen_4pass_1literals, reg309229)
reg309231 := PrimCons(reg309228, reg309230)
__ctx.Return(reg309231)
return
} else {
reg309232 := MakeSymbol("shen.pass-literals")
__ctx.TailApply(__defun__shen_4f__error, reg309232)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.pass-literals", value: __defun__shen_4pass_1literals})

__defun__shen_4defprolog_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1748 := __args[0]
_ = V1748
reg309234 := PrimIsPair(V1748)
var reg309249 Obj
if reg309234 == True {
reg309235 := MakeSymbol("defprolog")
reg309236 := PrimHead(V1748)
reg309237 := PrimEqual(reg309235, reg309236)
var reg309244 Obj
if reg309237 == True {
reg309238 := PrimTail(V1748)
reg309239 := PrimIsPair(reg309238)
var reg309242 Obj
if reg309239 == True {
reg309240 := True;
reg309242 = reg309240
} else {
reg309241 := False;
reg309242 = reg309241
}
reg309244 = reg309242
} else {
reg309243 := False;
reg309244 = reg309243
}
var reg309247 Obj
if reg309244 == True {
reg309245 := True;
reg309247 = reg309245
} else {
reg309246 := False;
reg309247 = reg309246
}
reg309249 = reg309247
} else {
reg309248 := False;
reg309249 = reg309248
}
if reg309249 == True {
reg309250 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4_5defprolog_6, Y)
return
}, 1)
reg309252 := PrimTail(V1748)
reg309253 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
reg309254 := PrimTail(V1748)
reg309255 := PrimHead(reg309254)
__ctx.TailApply(__defun__shen_4prolog_1error, reg309255, Y)
return
}, 1)
__ctx.TailApply(__defun__compile, reg309250, reg309252, reg309253)
return
} else {
__ctx.Return(V1748)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.defprolog-macro", value: __defun__shen_4defprolog_1macro})

__defun__shen_4datatype_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1750 := __args[0]
_ = V1750
reg309258 := PrimIsPair(V1750)
var reg309273 Obj
if reg309258 == True {
reg309259 := MakeSymbol("datatype")
reg309260 := PrimHead(V1750)
reg309261 := PrimEqual(reg309259, reg309260)
var reg309268 Obj
if reg309261 == True {
reg309262 := PrimTail(V1750)
reg309263 := PrimIsPair(reg309262)
var reg309266 Obj
if reg309263 == True {
reg309264 := True;
reg309266 = reg309264
} else {
reg309265 := False;
reg309266 = reg309265
}
reg309268 = reg309266
} else {
reg309267 := False;
reg309268 = reg309267
}
var reg309271 Obj
if reg309268 == True {
reg309269 := True;
reg309271 = reg309269
} else {
reg309270 := False;
reg309271 = reg309270
}
reg309273 = reg309271
} else {
reg309272 := False;
reg309273 = reg309272
}
if reg309273 == True {
reg309274 := MakeSymbol("shen.process-datatype")
reg309275 := PrimTail(V1750)
reg309276 := PrimHead(reg309275)
reg309277 := __e.Call(__defun__shen_4intern_1type, reg309276)
reg309278 := MakeSymbol("compile")
reg309279 := MakeSymbol("lambda")
reg309280 := MakeSymbol("X")
reg309281 := MakeSymbol("shen.<datatype-rules>")
reg309282 := MakeSymbol("X")
reg309283 := Nil;
reg309284 := PrimCons(reg309282, reg309283)
reg309285 := PrimCons(reg309281, reg309284)
reg309286 := Nil;
reg309287 := PrimCons(reg309285, reg309286)
reg309288 := PrimCons(reg309280, reg309287)
reg309289 := PrimCons(reg309279, reg309288)
reg309290 := PrimTail(V1750)
reg309291 := PrimTail(reg309290)
reg309292 := __e.Call(__defun__shen_4rcons__form, reg309291)
reg309293 := MakeSymbol("function")
reg309294 := MakeSymbol("shen.datatype-error")
reg309295 := Nil;
reg309296 := PrimCons(reg309294, reg309295)
reg309297 := PrimCons(reg309293, reg309296)
reg309298 := Nil;
reg309299 := PrimCons(reg309297, reg309298)
reg309300 := PrimCons(reg309292, reg309299)
reg309301 := PrimCons(reg309289, reg309300)
reg309302 := PrimCons(reg309278, reg309301)
reg309303 := Nil;
reg309304 := PrimCons(reg309302, reg309303)
reg309305 := PrimCons(reg309277, reg309304)
reg309306 := PrimCons(reg309274, reg309305)
__ctx.Return(reg309306)
return
} else {
__ctx.Return(V1750)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.datatype-macro", value: __defun__shen_4datatype_1macro})

__defun__shen_4intern_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1752 := __args[0]
_ = V1752
reg309307 := MakeString("type#")
reg309308 := PrimStr(V1752)
reg309309 := PrimStringConcat(reg309307, reg309308)
reg309310 := PrimIntern(reg309309)
__ctx.Return(reg309310)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.intern-type", value: __defun__shen_4intern_1type})

__defun__shen_4_8s_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1754 := __args[0]
_ = V1754
reg309311 := PrimIsPair(V1754)
var reg309343 Obj
if reg309311 == True {
reg309312 := MakeSymbol("@s")
reg309313 := PrimHead(V1754)
reg309314 := PrimEqual(reg309312, reg309313)
var reg309338 Obj
if reg309314 == True {
reg309315 := PrimTail(V1754)
reg309316 := PrimIsPair(reg309315)
var reg309333 Obj
if reg309316 == True {
reg309317 := PrimTail(V1754)
reg309318 := PrimTail(reg309317)
reg309319 := PrimIsPair(reg309318)
var reg309328 Obj
if reg309319 == True {
reg309320 := PrimTail(V1754)
reg309321 := PrimTail(reg309320)
reg309322 := PrimTail(reg309321)
reg309323 := PrimIsPair(reg309322)
var reg309326 Obj
if reg309323 == True {
reg309324 := True;
reg309326 = reg309324
} else {
reg309325 := False;
reg309326 = reg309325
}
reg309328 = reg309326
} else {
reg309327 := False;
reg309328 = reg309327
}
var reg309331 Obj
if reg309328 == True {
reg309329 := True;
reg309331 = reg309329
} else {
reg309330 := False;
reg309331 = reg309330
}
reg309333 = reg309331
} else {
reg309332 := False;
reg309333 = reg309332
}
var reg309336 Obj
if reg309333 == True {
reg309334 := True;
reg309336 = reg309334
} else {
reg309335 := False;
reg309336 = reg309335
}
reg309338 = reg309336
} else {
reg309337 := False;
reg309338 = reg309337
}
var reg309341 Obj
if reg309338 == True {
reg309339 := True;
reg309341 = reg309339
} else {
reg309340 := False;
reg309341 = reg309340
}
reg309343 = reg309341
} else {
reg309342 := False;
reg309343 = reg309342
}
if reg309343 == True {
reg309344 := MakeSymbol("@s")
reg309345 := PrimTail(V1754)
reg309346 := PrimHead(reg309345)
reg309347 := MakeSymbol("@s")
reg309348 := PrimTail(V1754)
reg309349 := PrimTail(reg309348)
reg309350 := PrimCons(reg309347, reg309349)
reg309351 := __e.Call(__defun__shen_4_8s_1macro, reg309350)
reg309352 := Nil;
reg309353 := PrimCons(reg309351, reg309352)
reg309354 := PrimCons(reg309346, reg309353)
reg309355 := PrimCons(reg309344, reg309354)
__ctx.Return(reg309355)
return
} else {
reg309356 := PrimIsPair(V1754)
var reg309397 Obj
if reg309356 == True {
reg309357 := MakeSymbol("@s")
reg309358 := PrimHead(V1754)
reg309359 := PrimEqual(reg309357, reg309358)
var reg309392 Obj
if reg309359 == True {
reg309360 := PrimTail(V1754)
reg309361 := PrimIsPair(reg309360)
var reg309387 Obj
if reg309361 == True {
reg309362 := PrimTail(V1754)
reg309363 := PrimTail(reg309362)
reg309364 := PrimIsPair(reg309363)
var reg309382 Obj
if reg309364 == True {
reg309365 := Nil;
reg309366 := PrimTail(V1754)
reg309367 := PrimTail(reg309366)
reg309368 := PrimTail(reg309367)
reg309369 := PrimEqual(reg309365, reg309368)
var reg309377 Obj
if reg309369 == True {
reg309370 := PrimTail(V1754)
reg309371 := PrimHead(reg309370)
reg309372 := PrimIsString(reg309371)
var reg309375 Obj
if reg309372 == True {
reg309373 := True;
reg309375 = reg309373
} else {
reg309374 := False;
reg309375 = reg309374
}
reg309377 = reg309375
} else {
reg309376 := False;
reg309377 = reg309376
}
var reg309380 Obj
if reg309377 == True {
reg309378 := True;
reg309380 = reg309378
} else {
reg309379 := False;
reg309380 = reg309379
}
reg309382 = reg309380
} else {
reg309381 := False;
reg309382 = reg309381
}
var reg309385 Obj
if reg309382 == True {
reg309383 := True;
reg309385 = reg309383
} else {
reg309384 := False;
reg309385 = reg309384
}
reg309387 = reg309385
} else {
reg309386 := False;
reg309387 = reg309386
}
var reg309390 Obj
if reg309387 == True {
reg309388 := True;
reg309390 = reg309388
} else {
reg309389 := False;
reg309390 = reg309389
}
reg309392 = reg309390
} else {
reg309391 := False;
reg309392 = reg309391
}
var reg309395 Obj
if reg309392 == True {
reg309393 := True;
reg309395 = reg309393
} else {
reg309394 := False;
reg309395 = reg309394
}
reg309397 = reg309395
} else {
reg309396 := False;
reg309397 = reg309396
}
if reg309397 == True {
reg309398 := PrimTail(V1754)
reg309399 := PrimHead(reg309398)
reg309400 := __e.Call(__defun__explode, reg309399)
E := reg309400
_ = E
reg309401 := __e.Call(__defun__length, E)
reg309402 := MakeNumber(1)
reg309403 := PrimGreatThan(reg309401, reg309402)
if reg309403 == True {
reg309404 := MakeSymbol("@s")
reg309405 := PrimTail(V1754)
reg309406 := PrimTail(reg309405)
reg309407 := __e.Call(__defun__append, E, reg309406)
reg309408 := PrimCons(reg309404, reg309407)
__ctx.TailApply(__defun__shen_4_8s_1macro, reg309408)
return
} else {
__ctx.Return(V1754)
return
}
} else {
__ctx.Return(V1754)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.@s-macro", value: __defun__shen_4_8s_1macro})

__defun__shen_4synonyms_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1756 := __args[0]
_ = V1756
reg309410 := PrimIsPair(V1756)
var reg309418 Obj
if reg309410 == True {
reg309411 := MakeSymbol("synonyms")
reg309412 := PrimHead(V1756)
reg309413 := PrimEqual(reg309411, reg309412)
var reg309416 Obj
if reg309413 == True {
reg309414 := True;
reg309416 = reg309414
} else {
reg309415 := False;
reg309416 = reg309415
}
reg309418 = reg309416
} else {
reg309417 := False;
reg309418 = reg309417
}
if reg309418 == True {
reg309419 := MakeSymbol("shen.synonyms-help")
reg309420 := PrimTail(V1756)
reg309421 := __e.Call(__defun__shen_4curry_1synonyms, reg309420)
reg309422 := __e.Call(__defun__shen_4rcons__form, reg309421)
reg309423 := Nil;
reg309424 := PrimCons(reg309422, reg309423)
reg309425 := PrimCons(reg309419, reg309424)
__ctx.Return(reg309425)
return
} else {
__ctx.Return(V1756)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.synonyms-macro", value: __defun__shen_4synonyms_1macro})

__defun__shen_4curry_1synonyms = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1758 := __args[0]
_ = V1758
reg309426 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4curry_1type, X)
return
}, 1)
__ctx.TailApply(__defun__map, reg309426, V1758)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.curry-synonyms", value: __defun__shen_4curry_1synonyms})

__defun__shen_4nl_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1760 := __args[0]
_ = V1760
reg309429 := PrimIsPair(V1760)
var reg309445 Obj
if reg309429 == True {
reg309430 := MakeSymbol("nl")
reg309431 := PrimHead(V1760)
reg309432 := PrimEqual(reg309430, reg309431)
var reg309440 Obj
if reg309432 == True {
reg309433 := Nil;
reg309434 := PrimTail(V1760)
reg309435 := PrimEqual(reg309433, reg309434)
var reg309438 Obj
if reg309435 == True {
reg309436 := True;
reg309438 = reg309436
} else {
reg309437 := False;
reg309438 = reg309437
}
reg309440 = reg309438
} else {
reg309439 := False;
reg309440 = reg309439
}
var reg309443 Obj
if reg309440 == True {
reg309441 := True;
reg309443 = reg309441
} else {
reg309442 := False;
reg309443 = reg309442
}
reg309445 = reg309443
} else {
reg309444 := False;
reg309445 = reg309444
}
if reg309445 == True {
reg309446 := MakeSymbol("nl")
reg309447 := MakeNumber(1)
reg309448 := Nil;
reg309449 := PrimCons(reg309447, reg309448)
reg309450 := PrimCons(reg309446, reg309449)
__ctx.Return(reg309450)
return
} else {
__ctx.Return(V1760)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.nl-macro", value: __defun__shen_4nl_1macro})

__defun__shen_4assoc_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1762 := __args[0]
_ = V1762
reg309451 := PrimIsPair(V1762)
var reg309499 Obj
if reg309451 == True {
reg309452 := PrimTail(V1762)
reg309453 := PrimIsPair(reg309452)
var reg309494 Obj
if reg309453 == True {
reg309454 := PrimTail(V1762)
reg309455 := PrimTail(reg309454)
reg309456 := PrimIsPair(reg309455)
var reg309489 Obj
if reg309456 == True {
reg309457 := PrimTail(V1762)
reg309458 := PrimTail(reg309457)
reg309459 := PrimTail(reg309458)
reg309460 := PrimIsPair(reg309459)
var reg309484 Obj
if reg309460 == True {
reg309461 := PrimHead(V1762)
reg309462 := MakeSymbol("@p")
reg309463 := MakeSymbol("@v")
reg309464 := MakeSymbol("append")
reg309465 := MakeSymbol("and")
reg309466 := MakeSymbol("or")
reg309467 := MakeSymbol("+")
reg309468 := MakeSymbol("*")
reg309469 := MakeSymbol("do")
reg309470 := Nil;
reg309471 := PrimCons(reg309469, reg309470)
reg309472 := PrimCons(reg309468, reg309471)
reg309473 := PrimCons(reg309467, reg309472)
reg309474 := PrimCons(reg309466, reg309473)
reg309475 := PrimCons(reg309465, reg309474)
reg309476 := PrimCons(reg309464, reg309475)
reg309477 := PrimCons(reg309463, reg309476)
reg309478 := PrimCons(reg309462, reg309477)
reg309479 := __e.Call(__defun__element_2, reg309461, reg309478)
var reg309482 Obj
if reg309479 == True {
reg309480 := True;
reg309482 = reg309480
} else {
reg309481 := False;
reg309482 = reg309481
}
reg309484 = reg309482
} else {
reg309483 := False;
reg309484 = reg309483
}
var reg309487 Obj
if reg309484 == True {
reg309485 := True;
reg309487 = reg309485
} else {
reg309486 := False;
reg309487 = reg309486
}
reg309489 = reg309487
} else {
reg309488 := False;
reg309489 = reg309488
}
var reg309492 Obj
if reg309489 == True {
reg309490 := True;
reg309492 = reg309490
} else {
reg309491 := False;
reg309492 = reg309491
}
reg309494 = reg309492
} else {
reg309493 := False;
reg309494 = reg309493
}
var reg309497 Obj
if reg309494 == True {
reg309495 := True;
reg309497 = reg309495
} else {
reg309496 := False;
reg309497 = reg309496
}
reg309499 = reg309497
} else {
reg309498 := False;
reg309499 = reg309498
}
if reg309499 == True {
reg309500 := PrimHead(V1762)
reg309501 := PrimTail(V1762)
reg309502 := PrimHead(reg309501)
reg309503 := PrimHead(V1762)
reg309504 := PrimTail(V1762)
reg309505 := PrimTail(reg309504)
reg309506 := PrimCons(reg309503, reg309505)
reg309507 := __e.Call(__defun__shen_4assoc_1macro, reg309506)
reg309508 := Nil;
reg309509 := PrimCons(reg309507, reg309508)
reg309510 := PrimCons(reg309502, reg309509)
reg309511 := PrimCons(reg309500, reg309510)
__ctx.Return(reg309511)
return
} else {
__ctx.Return(V1762)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.assoc-macro", value: __defun__shen_4assoc_1macro})

__defun__shen_4let_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1764 := __args[0]
_ = V1764
reg309512 := PrimIsPair(V1764)
var reg309554 Obj
if reg309512 == True {
reg309513 := MakeSymbol("let")
reg309514 := PrimHead(V1764)
reg309515 := PrimEqual(reg309513, reg309514)
var reg309549 Obj
if reg309515 == True {
reg309516 := PrimTail(V1764)
reg309517 := PrimIsPair(reg309516)
var reg309544 Obj
if reg309517 == True {
reg309518 := PrimTail(V1764)
reg309519 := PrimTail(reg309518)
reg309520 := PrimIsPair(reg309519)
var reg309539 Obj
if reg309520 == True {
reg309521 := PrimTail(V1764)
reg309522 := PrimTail(reg309521)
reg309523 := PrimTail(reg309522)
reg309524 := PrimIsPair(reg309523)
var reg309534 Obj
if reg309524 == True {
reg309525 := PrimTail(V1764)
reg309526 := PrimTail(reg309525)
reg309527 := PrimTail(reg309526)
reg309528 := PrimTail(reg309527)
reg309529 := PrimIsPair(reg309528)
var reg309532 Obj
if reg309529 == True {
reg309530 := True;
reg309532 = reg309530
} else {
reg309531 := False;
reg309532 = reg309531
}
reg309534 = reg309532
} else {
reg309533 := False;
reg309534 = reg309533
}
var reg309537 Obj
if reg309534 == True {
reg309535 := True;
reg309537 = reg309535
} else {
reg309536 := False;
reg309537 = reg309536
}
reg309539 = reg309537
} else {
reg309538 := False;
reg309539 = reg309538
}
var reg309542 Obj
if reg309539 == True {
reg309540 := True;
reg309542 = reg309540
} else {
reg309541 := False;
reg309542 = reg309541
}
reg309544 = reg309542
} else {
reg309543 := False;
reg309544 = reg309543
}
var reg309547 Obj
if reg309544 == True {
reg309545 := True;
reg309547 = reg309545
} else {
reg309546 := False;
reg309547 = reg309546
}
reg309549 = reg309547
} else {
reg309548 := False;
reg309549 = reg309548
}
var reg309552 Obj
if reg309549 == True {
reg309550 := True;
reg309552 = reg309550
} else {
reg309551 := False;
reg309552 = reg309551
}
reg309554 = reg309552
} else {
reg309553 := False;
reg309554 = reg309553
}
if reg309554 == True {
reg309555 := MakeSymbol("let")
reg309556 := PrimTail(V1764)
reg309557 := PrimHead(reg309556)
reg309558 := PrimTail(V1764)
reg309559 := PrimTail(reg309558)
reg309560 := PrimHead(reg309559)
reg309561 := MakeSymbol("let")
reg309562 := PrimTail(V1764)
reg309563 := PrimTail(reg309562)
reg309564 := PrimTail(reg309563)
reg309565 := PrimCons(reg309561, reg309564)
reg309566 := __e.Call(__defun__shen_4let_1macro, reg309565)
reg309567 := Nil;
reg309568 := PrimCons(reg309566, reg309567)
reg309569 := PrimCons(reg309560, reg309568)
reg309570 := PrimCons(reg309557, reg309569)
reg309571 := PrimCons(reg309555, reg309570)
__ctx.Return(reg309571)
return
} else {
__ctx.Return(V1764)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.let-macro", value: __defun__shen_4let_1macro})

__defun__shen_4abs_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1766 := __args[0]
_ = V1766
reg309572 := PrimIsPair(V1766)
var reg309604 Obj
if reg309572 == True {
reg309573 := MakeSymbol("/.")
reg309574 := PrimHead(V1766)
reg309575 := PrimEqual(reg309573, reg309574)
var reg309599 Obj
if reg309575 == True {
reg309576 := PrimTail(V1766)
reg309577 := PrimIsPair(reg309576)
var reg309594 Obj
if reg309577 == True {
reg309578 := PrimTail(V1766)
reg309579 := PrimTail(reg309578)
reg309580 := PrimIsPair(reg309579)
var reg309589 Obj
if reg309580 == True {
reg309581 := PrimTail(V1766)
reg309582 := PrimTail(reg309581)
reg309583 := PrimTail(reg309582)
reg309584 := PrimIsPair(reg309583)
var reg309587 Obj
if reg309584 == True {
reg309585 := True;
reg309587 = reg309585
} else {
reg309586 := False;
reg309587 = reg309586
}
reg309589 = reg309587
} else {
reg309588 := False;
reg309589 = reg309588
}
var reg309592 Obj
if reg309589 == True {
reg309590 := True;
reg309592 = reg309590
} else {
reg309591 := False;
reg309592 = reg309591
}
reg309594 = reg309592
} else {
reg309593 := False;
reg309594 = reg309593
}
var reg309597 Obj
if reg309594 == True {
reg309595 := True;
reg309597 = reg309595
} else {
reg309596 := False;
reg309597 = reg309596
}
reg309599 = reg309597
} else {
reg309598 := False;
reg309599 = reg309598
}
var reg309602 Obj
if reg309599 == True {
reg309600 := True;
reg309602 = reg309600
} else {
reg309601 := False;
reg309602 = reg309601
}
reg309604 = reg309602
} else {
reg309603 := False;
reg309604 = reg309603
}
if reg309604 == True {
reg309605 := MakeSymbol("lambda")
reg309606 := PrimTail(V1766)
reg309607 := PrimHead(reg309606)
reg309608 := MakeSymbol("/.")
reg309609 := PrimTail(V1766)
reg309610 := PrimTail(reg309609)
reg309611 := PrimCons(reg309608, reg309610)
reg309612 := __e.Call(__defun__shen_4abs_1macro, reg309611)
reg309613 := Nil;
reg309614 := PrimCons(reg309612, reg309613)
reg309615 := PrimCons(reg309607, reg309614)
reg309616 := PrimCons(reg309605, reg309615)
__ctx.Return(reg309616)
return
} else {
reg309617 := PrimIsPair(V1766)
var reg309650 Obj
if reg309617 == True {
reg309618 := MakeSymbol("/.")
reg309619 := PrimHead(V1766)
reg309620 := PrimEqual(reg309618, reg309619)
var reg309645 Obj
if reg309620 == True {
reg309621 := PrimTail(V1766)
reg309622 := PrimIsPair(reg309621)
var reg309640 Obj
if reg309622 == True {
reg309623 := PrimTail(V1766)
reg309624 := PrimTail(reg309623)
reg309625 := PrimIsPair(reg309624)
var reg309635 Obj
if reg309625 == True {
reg309626 := Nil;
reg309627 := PrimTail(V1766)
reg309628 := PrimTail(reg309627)
reg309629 := PrimTail(reg309628)
reg309630 := PrimEqual(reg309626, reg309629)
var reg309633 Obj
if reg309630 == True {
reg309631 := True;
reg309633 = reg309631
} else {
reg309632 := False;
reg309633 = reg309632
}
reg309635 = reg309633
} else {
reg309634 := False;
reg309635 = reg309634
}
var reg309638 Obj
if reg309635 == True {
reg309636 := True;
reg309638 = reg309636
} else {
reg309637 := False;
reg309638 = reg309637
}
reg309640 = reg309638
} else {
reg309639 := False;
reg309640 = reg309639
}
var reg309643 Obj
if reg309640 == True {
reg309641 := True;
reg309643 = reg309641
} else {
reg309642 := False;
reg309643 = reg309642
}
reg309645 = reg309643
} else {
reg309644 := False;
reg309645 = reg309644
}
var reg309648 Obj
if reg309645 == True {
reg309646 := True;
reg309648 = reg309646
} else {
reg309647 := False;
reg309648 = reg309647
}
reg309650 = reg309648
} else {
reg309649 := False;
reg309650 = reg309649
}
if reg309650 == True {
reg309651 := MakeSymbol("lambda")
reg309652 := PrimTail(V1766)
reg309653 := PrimCons(reg309651, reg309652)
__ctx.Return(reg309653)
return
} else {
__ctx.Return(V1766)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.abs-macro", value: __defun__shen_4abs_1macro})

__defun__shen_4cases_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1770 := __args[0]
_ = V1770
reg309654 := PrimIsPair(V1770)
var reg309686 Obj
if reg309654 == True {
reg309655 := MakeSymbol("cases")
reg309656 := PrimHead(V1770)
reg309657 := PrimEqual(reg309655, reg309656)
var reg309681 Obj
if reg309657 == True {
reg309658 := PrimTail(V1770)
reg309659 := PrimIsPair(reg309658)
var reg309676 Obj
if reg309659 == True {
reg309660 := True;
reg309661 := PrimTail(V1770)
reg309662 := PrimHead(reg309661)
reg309663 := PrimEqual(reg309660, reg309662)
var reg309671 Obj
if reg309663 == True {
reg309664 := PrimTail(V1770)
reg309665 := PrimTail(reg309664)
reg309666 := PrimIsPair(reg309665)
var reg309669 Obj
if reg309666 == True {
reg309667 := True;
reg309669 = reg309667
} else {
reg309668 := False;
reg309669 = reg309668
}
reg309671 = reg309669
} else {
reg309670 := False;
reg309671 = reg309670
}
var reg309674 Obj
if reg309671 == True {
reg309672 := True;
reg309674 = reg309672
} else {
reg309673 := False;
reg309674 = reg309673
}
reg309676 = reg309674
} else {
reg309675 := False;
reg309676 = reg309675
}
var reg309679 Obj
if reg309676 == True {
reg309677 := True;
reg309679 = reg309677
} else {
reg309678 := False;
reg309679 = reg309678
}
reg309681 = reg309679
} else {
reg309680 := False;
reg309681 = reg309680
}
var reg309684 Obj
if reg309681 == True {
reg309682 := True;
reg309684 = reg309682
} else {
reg309683 := False;
reg309684 = reg309683
}
reg309686 = reg309684
} else {
reg309685 := False;
reg309686 = reg309685
}
if reg309686 == True {
reg309687 := PrimTail(V1770)
reg309688 := PrimTail(reg309687)
reg309689 := PrimHead(reg309688)
__ctx.Return(reg309689)
return
} else {
reg309690 := PrimIsPair(V1770)
var reg309723 Obj
if reg309690 == True {
reg309691 := MakeSymbol("cases")
reg309692 := PrimHead(V1770)
reg309693 := PrimEqual(reg309691, reg309692)
var reg309718 Obj
if reg309693 == True {
reg309694 := PrimTail(V1770)
reg309695 := PrimIsPair(reg309694)
var reg309713 Obj
if reg309695 == True {
reg309696 := PrimTail(V1770)
reg309697 := PrimTail(reg309696)
reg309698 := PrimIsPair(reg309697)
var reg309708 Obj
if reg309698 == True {
reg309699 := Nil;
reg309700 := PrimTail(V1770)
reg309701 := PrimTail(reg309700)
reg309702 := PrimTail(reg309701)
reg309703 := PrimEqual(reg309699, reg309702)
var reg309706 Obj
if reg309703 == True {
reg309704 := True;
reg309706 = reg309704
} else {
reg309705 := False;
reg309706 = reg309705
}
reg309708 = reg309706
} else {
reg309707 := False;
reg309708 = reg309707
}
var reg309711 Obj
if reg309708 == True {
reg309709 := True;
reg309711 = reg309709
} else {
reg309710 := False;
reg309711 = reg309710
}
reg309713 = reg309711
} else {
reg309712 := False;
reg309713 = reg309712
}
var reg309716 Obj
if reg309713 == True {
reg309714 := True;
reg309716 = reg309714
} else {
reg309715 := False;
reg309716 = reg309715
}
reg309718 = reg309716
} else {
reg309717 := False;
reg309718 = reg309717
}
var reg309721 Obj
if reg309718 == True {
reg309719 := True;
reg309721 = reg309719
} else {
reg309720 := False;
reg309721 = reg309720
}
reg309723 = reg309721
} else {
reg309722 := False;
reg309723 = reg309722
}
if reg309723 == True {
reg309724 := MakeSymbol("if")
reg309725 := PrimTail(V1770)
reg309726 := PrimHead(reg309725)
reg309727 := PrimTail(V1770)
reg309728 := PrimTail(reg309727)
reg309729 := PrimHead(reg309728)
reg309730 := MakeSymbol("simple-error")
reg309731 := MakeString("error: cases exhausted")
reg309732 := Nil;
reg309733 := PrimCons(reg309731, reg309732)
reg309734 := PrimCons(reg309730, reg309733)
reg309735 := Nil;
reg309736 := PrimCons(reg309734, reg309735)
reg309737 := PrimCons(reg309729, reg309736)
reg309738 := PrimCons(reg309726, reg309737)
reg309739 := PrimCons(reg309724, reg309738)
__ctx.Return(reg309739)
return
} else {
reg309740 := PrimIsPair(V1770)
var reg309763 Obj
if reg309740 == True {
reg309741 := MakeSymbol("cases")
reg309742 := PrimHead(V1770)
reg309743 := PrimEqual(reg309741, reg309742)
var reg309758 Obj
if reg309743 == True {
reg309744 := PrimTail(V1770)
reg309745 := PrimIsPair(reg309744)
var reg309753 Obj
if reg309745 == True {
reg309746 := PrimTail(V1770)
reg309747 := PrimTail(reg309746)
reg309748 := PrimIsPair(reg309747)
var reg309751 Obj
if reg309748 == True {
reg309749 := True;
reg309751 = reg309749
} else {
reg309750 := False;
reg309751 = reg309750
}
reg309753 = reg309751
} else {
reg309752 := False;
reg309753 = reg309752
}
var reg309756 Obj
if reg309753 == True {
reg309754 := True;
reg309756 = reg309754
} else {
reg309755 := False;
reg309756 = reg309755
}
reg309758 = reg309756
} else {
reg309757 := False;
reg309758 = reg309757
}
var reg309761 Obj
if reg309758 == True {
reg309759 := True;
reg309761 = reg309759
} else {
reg309760 := False;
reg309761 = reg309760
}
reg309763 = reg309761
} else {
reg309762 := False;
reg309763 = reg309762
}
if reg309763 == True {
reg309764 := MakeSymbol("if")
reg309765 := PrimTail(V1770)
reg309766 := PrimHead(reg309765)
reg309767 := PrimTail(V1770)
reg309768 := PrimTail(reg309767)
reg309769 := PrimHead(reg309768)
reg309770 := MakeSymbol("cases")
reg309771 := PrimTail(V1770)
reg309772 := PrimTail(reg309771)
reg309773 := PrimTail(reg309772)
reg309774 := PrimCons(reg309770, reg309773)
reg309775 := __e.Call(__defun__shen_4cases_1macro, reg309774)
reg309776 := Nil;
reg309777 := PrimCons(reg309775, reg309776)
reg309778 := PrimCons(reg309769, reg309777)
reg309779 := PrimCons(reg309766, reg309778)
reg309780 := PrimCons(reg309764, reg309779)
__ctx.Return(reg309780)
return
} else {
reg309781 := PrimIsPair(V1770)
var reg309805 Obj
if reg309781 == True {
reg309782 := MakeSymbol("cases")
reg309783 := PrimHead(V1770)
reg309784 := PrimEqual(reg309782, reg309783)
var reg309800 Obj
if reg309784 == True {
reg309785 := PrimTail(V1770)
reg309786 := PrimIsPair(reg309785)
var reg309795 Obj
if reg309786 == True {
reg309787 := Nil;
reg309788 := PrimTail(V1770)
reg309789 := PrimTail(reg309788)
reg309790 := PrimEqual(reg309787, reg309789)
var reg309793 Obj
if reg309790 == True {
reg309791 := True;
reg309793 = reg309791
} else {
reg309792 := False;
reg309793 = reg309792
}
reg309795 = reg309793
} else {
reg309794 := False;
reg309795 = reg309794
}
var reg309798 Obj
if reg309795 == True {
reg309796 := True;
reg309798 = reg309796
} else {
reg309797 := False;
reg309798 = reg309797
}
reg309800 = reg309798
} else {
reg309799 := False;
reg309800 = reg309799
}
var reg309803 Obj
if reg309800 == True {
reg309801 := True;
reg309803 = reg309801
} else {
reg309802 := False;
reg309803 = reg309802
}
reg309805 = reg309803
} else {
reg309804 := False;
reg309805 = reg309804
}
if reg309805 == True {
reg309806 := MakeString("error: odd number of case elements\n")
reg309807 := PrimSimpleError(reg309806)
__ctx.Return(reg309807)
return
} else {
__ctx.Return(V1770)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.cases-macro", value: __defun__shen_4cases_1macro})

__defun__shen_4timer_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1772 := __args[0]
_ = V1772
reg309808 := PrimIsPair(V1772)
var reg309832 Obj
if reg309808 == True {
reg309809 := MakeSymbol("time")
reg309810 := PrimHead(V1772)
reg309811 := PrimEqual(reg309809, reg309810)
var reg309827 Obj
if reg309811 == True {
reg309812 := PrimTail(V1772)
reg309813 := PrimIsPair(reg309812)
var reg309822 Obj
if reg309813 == True {
reg309814 := Nil;
reg309815 := PrimTail(V1772)
reg309816 := PrimTail(reg309815)
reg309817 := PrimEqual(reg309814, reg309816)
var reg309820 Obj
if reg309817 == True {
reg309818 := True;
reg309820 = reg309818
} else {
reg309819 := False;
reg309820 = reg309819
}
reg309822 = reg309820
} else {
reg309821 := False;
reg309822 = reg309821
}
var reg309825 Obj
if reg309822 == True {
reg309823 := True;
reg309825 = reg309823
} else {
reg309824 := False;
reg309825 = reg309824
}
reg309827 = reg309825
} else {
reg309826 := False;
reg309827 = reg309826
}
var reg309830 Obj
if reg309827 == True {
reg309828 := True;
reg309830 = reg309828
} else {
reg309829 := False;
reg309830 = reg309829
}
reg309832 = reg309830
} else {
reg309831 := False;
reg309832 = reg309831
}
if reg309832 == True {
reg309833 := MakeSymbol("let")
reg309834 := MakeSymbol("Start")
reg309835 := MakeSymbol("get-time")
reg309836 := MakeSymbol("run")
reg309837 := Nil;
reg309838 := PrimCons(reg309836, reg309837)
reg309839 := PrimCons(reg309835, reg309838)
reg309840 := MakeSymbol("Result")
reg309841 := PrimTail(V1772)
reg309842 := PrimHead(reg309841)
reg309843 := MakeSymbol("Finish")
reg309844 := MakeSymbol("get-time")
reg309845 := MakeSymbol("run")
reg309846 := Nil;
reg309847 := PrimCons(reg309845, reg309846)
reg309848 := PrimCons(reg309844, reg309847)
reg309849 := MakeSymbol("Time")
reg309850 := MakeSymbol("-")
reg309851 := MakeSymbol("Finish")
reg309852 := MakeSymbol("Start")
reg309853 := Nil;
reg309854 := PrimCons(reg309852, reg309853)
reg309855 := PrimCons(reg309851, reg309854)
reg309856 := PrimCons(reg309850, reg309855)
reg309857 := MakeSymbol("Message")
reg309858 := MakeSymbol("shen.prhush")
reg309859 := MakeSymbol("cn")
reg309860 := MakeString("\nrun time: ")
reg309861 := MakeSymbol("cn")
reg309862 := MakeSymbol("str")
reg309863 := MakeSymbol("Time")
reg309864 := Nil;
reg309865 := PrimCons(reg309863, reg309864)
reg309866 := PrimCons(reg309862, reg309865)
reg309867 := MakeString(" secs\n")
reg309868 := Nil;
reg309869 := PrimCons(reg309867, reg309868)
reg309870 := PrimCons(reg309866, reg309869)
reg309871 := PrimCons(reg309861, reg309870)
reg309872 := Nil;
reg309873 := PrimCons(reg309871, reg309872)
reg309874 := PrimCons(reg309860, reg309873)
reg309875 := PrimCons(reg309859, reg309874)
reg309876 := MakeSymbol("stoutput")
reg309877 := Nil;
reg309878 := PrimCons(reg309876, reg309877)
reg309879 := Nil;
reg309880 := PrimCons(reg309878, reg309879)
reg309881 := PrimCons(reg309875, reg309880)
reg309882 := PrimCons(reg309858, reg309881)
reg309883 := MakeSymbol("Result")
reg309884 := Nil;
reg309885 := PrimCons(reg309883, reg309884)
reg309886 := PrimCons(reg309882, reg309885)
reg309887 := PrimCons(reg309857, reg309886)
reg309888 := PrimCons(reg309856, reg309887)
reg309889 := PrimCons(reg309849, reg309888)
reg309890 := PrimCons(reg309848, reg309889)
reg309891 := PrimCons(reg309843, reg309890)
reg309892 := PrimCons(reg309842, reg309891)
reg309893 := PrimCons(reg309840, reg309892)
reg309894 := PrimCons(reg309839, reg309893)
reg309895 := PrimCons(reg309834, reg309894)
reg309896 := PrimCons(reg309833, reg309895)
__ctx.TailApply(__defun__shen_4let_1macro, reg309896)
return
} else {
__ctx.Return(V1772)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.timer-macro", value: __defun__shen_4timer_1macro})

__defun__shen_4tuple_1up = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1774 := __args[0]
_ = V1774
reg309898 := PrimIsPair(V1774)
if reg309898 == True {
reg309899 := MakeSymbol("@p")
reg309900 := PrimHead(V1774)
reg309901 := PrimTail(V1774)
reg309902 := __e.Call(__defun__shen_4tuple_1up, reg309901)
reg309903 := Nil;
reg309904 := PrimCons(reg309902, reg309903)
reg309905 := PrimCons(reg309900, reg309904)
reg309906 := PrimCons(reg309899, reg309905)
__ctx.Return(reg309906)
return
} else {
__ctx.Return(V1774)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.tuple-up", value: __defun__shen_4tuple_1up})

__defun__shen_4put_cget_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1776 := __args[0]
_ = V1776
reg309907 := PrimIsPair(V1776)
var reg309950 Obj
if reg309907 == True {
reg309908 := MakeSymbol("put")
reg309909 := PrimHead(V1776)
reg309910 := PrimEqual(reg309908, reg309909)
var reg309945 Obj
if reg309910 == True {
reg309911 := PrimTail(V1776)
reg309912 := PrimIsPair(reg309911)
var reg309940 Obj
if reg309912 == True {
reg309913 := PrimTail(V1776)
reg309914 := PrimTail(reg309913)
reg309915 := PrimIsPair(reg309914)
var reg309935 Obj
if reg309915 == True {
reg309916 := PrimTail(V1776)
reg309917 := PrimTail(reg309916)
reg309918 := PrimTail(reg309917)
reg309919 := PrimIsPair(reg309918)
var reg309930 Obj
if reg309919 == True {
reg309920 := Nil;
reg309921 := PrimTail(V1776)
reg309922 := PrimTail(reg309921)
reg309923 := PrimTail(reg309922)
reg309924 := PrimTail(reg309923)
reg309925 := PrimEqual(reg309920, reg309924)
var reg309928 Obj
if reg309925 == True {
reg309926 := True;
reg309928 = reg309926
} else {
reg309927 := False;
reg309928 = reg309927
}
reg309930 = reg309928
} else {
reg309929 := False;
reg309930 = reg309929
}
var reg309933 Obj
if reg309930 == True {
reg309931 := True;
reg309933 = reg309931
} else {
reg309932 := False;
reg309933 = reg309932
}
reg309935 = reg309933
} else {
reg309934 := False;
reg309935 = reg309934
}
var reg309938 Obj
if reg309935 == True {
reg309936 := True;
reg309938 = reg309936
} else {
reg309937 := False;
reg309938 = reg309937
}
reg309940 = reg309938
} else {
reg309939 := False;
reg309940 = reg309939
}
var reg309943 Obj
if reg309940 == True {
reg309941 := True;
reg309943 = reg309941
} else {
reg309942 := False;
reg309943 = reg309942
}
reg309945 = reg309943
} else {
reg309944 := False;
reg309945 = reg309944
}
var reg309948 Obj
if reg309945 == True {
reg309946 := True;
reg309948 = reg309946
} else {
reg309947 := False;
reg309948 = reg309947
}
reg309950 = reg309948
} else {
reg309949 := False;
reg309950 = reg309949
}
if reg309950 == True {
reg309951 := MakeSymbol("put")
reg309952 := PrimTail(V1776)
reg309953 := PrimHead(reg309952)
reg309954 := PrimTail(V1776)
reg309955 := PrimTail(reg309954)
reg309956 := PrimHead(reg309955)
reg309957 := PrimTail(V1776)
reg309958 := PrimTail(reg309957)
reg309959 := PrimTail(reg309958)
reg309960 := PrimHead(reg309959)
reg309961 := MakeSymbol("value")
reg309962 := MakeSymbol("*property-vector*")
reg309963 := Nil;
reg309964 := PrimCons(reg309962, reg309963)
reg309965 := PrimCons(reg309961, reg309964)
reg309966 := Nil;
reg309967 := PrimCons(reg309965, reg309966)
reg309968 := PrimCons(reg309960, reg309967)
reg309969 := PrimCons(reg309956, reg309968)
reg309970 := PrimCons(reg309953, reg309969)
reg309971 := PrimCons(reg309951, reg309970)
__ctx.Return(reg309971)
return
} else {
reg309972 := PrimIsPair(V1776)
var reg310005 Obj
if reg309972 == True {
reg309973 := MakeSymbol("get")
reg309974 := PrimHead(V1776)
reg309975 := PrimEqual(reg309973, reg309974)
var reg310000 Obj
if reg309975 == True {
reg309976 := PrimTail(V1776)
reg309977 := PrimIsPair(reg309976)
var reg309995 Obj
if reg309977 == True {
reg309978 := PrimTail(V1776)
reg309979 := PrimTail(reg309978)
reg309980 := PrimIsPair(reg309979)
var reg309990 Obj
if reg309980 == True {
reg309981 := Nil;
reg309982 := PrimTail(V1776)
reg309983 := PrimTail(reg309982)
reg309984 := PrimTail(reg309983)
reg309985 := PrimEqual(reg309981, reg309984)
var reg309988 Obj
if reg309985 == True {
reg309986 := True;
reg309988 = reg309986
} else {
reg309987 := False;
reg309988 = reg309987
}
reg309990 = reg309988
} else {
reg309989 := False;
reg309990 = reg309989
}
var reg309993 Obj
if reg309990 == True {
reg309991 := True;
reg309993 = reg309991
} else {
reg309992 := False;
reg309993 = reg309992
}
reg309995 = reg309993
} else {
reg309994 := False;
reg309995 = reg309994
}
var reg309998 Obj
if reg309995 == True {
reg309996 := True;
reg309998 = reg309996
} else {
reg309997 := False;
reg309998 = reg309997
}
reg310000 = reg309998
} else {
reg309999 := False;
reg310000 = reg309999
}
var reg310003 Obj
if reg310000 == True {
reg310001 := True;
reg310003 = reg310001
} else {
reg310002 := False;
reg310003 = reg310002
}
reg310005 = reg310003
} else {
reg310004 := False;
reg310005 = reg310004
}
if reg310005 == True {
reg310006 := MakeSymbol("get")
reg310007 := PrimTail(V1776)
reg310008 := PrimHead(reg310007)
reg310009 := PrimTail(V1776)
reg310010 := PrimTail(reg310009)
reg310011 := PrimHead(reg310010)
reg310012 := MakeSymbol("value")
reg310013 := MakeSymbol("*property-vector*")
reg310014 := Nil;
reg310015 := PrimCons(reg310013, reg310014)
reg310016 := PrimCons(reg310012, reg310015)
reg310017 := Nil;
reg310018 := PrimCons(reg310016, reg310017)
reg310019 := PrimCons(reg310011, reg310018)
reg310020 := PrimCons(reg310008, reg310019)
reg310021 := PrimCons(reg310006, reg310020)
__ctx.Return(reg310021)
return
} else {
reg310022 := PrimIsPair(V1776)
var reg310055 Obj
if reg310022 == True {
reg310023 := MakeSymbol("unput")
reg310024 := PrimHead(V1776)
reg310025 := PrimEqual(reg310023, reg310024)
var reg310050 Obj
if reg310025 == True {
reg310026 := PrimTail(V1776)
reg310027 := PrimIsPair(reg310026)
var reg310045 Obj
if reg310027 == True {
reg310028 := PrimTail(V1776)
reg310029 := PrimTail(reg310028)
reg310030 := PrimIsPair(reg310029)
var reg310040 Obj
if reg310030 == True {
reg310031 := Nil;
reg310032 := PrimTail(V1776)
reg310033 := PrimTail(reg310032)
reg310034 := PrimTail(reg310033)
reg310035 := PrimEqual(reg310031, reg310034)
var reg310038 Obj
if reg310035 == True {
reg310036 := True;
reg310038 = reg310036
} else {
reg310037 := False;
reg310038 = reg310037
}
reg310040 = reg310038
} else {
reg310039 := False;
reg310040 = reg310039
}
var reg310043 Obj
if reg310040 == True {
reg310041 := True;
reg310043 = reg310041
} else {
reg310042 := False;
reg310043 = reg310042
}
reg310045 = reg310043
} else {
reg310044 := False;
reg310045 = reg310044
}
var reg310048 Obj
if reg310045 == True {
reg310046 := True;
reg310048 = reg310046
} else {
reg310047 := False;
reg310048 = reg310047
}
reg310050 = reg310048
} else {
reg310049 := False;
reg310050 = reg310049
}
var reg310053 Obj
if reg310050 == True {
reg310051 := True;
reg310053 = reg310051
} else {
reg310052 := False;
reg310053 = reg310052
}
reg310055 = reg310053
} else {
reg310054 := False;
reg310055 = reg310054
}
if reg310055 == True {
reg310056 := MakeSymbol("unput")
reg310057 := PrimTail(V1776)
reg310058 := PrimHead(reg310057)
reg310059 := PrimTail(V1776)
reg310060 := PrimTail(reg310059)
reg310061 := PrimHead(reg310060)
reg310062 := MakeSymbol("value")
reg310063 := MakeSymbol("*property-vector*")
reg310064 := Nil;
reg310065 := PrimCons(reg310063, reg310064)
reg310066 := PrimCons(reg310062, reg310065)
reg310067 := Nil;
reg310068 := PrimCons(reg310066, reg310067)
reg310069 := PrimCons(reg310061, reg310068)
reg310070 := PrimCons(reg310058, reg310069)
reg310071 := PrimCons(reg310056, reg310070)
__ctx.Return(reg310071)
return
} else {
__ctx.Return(V1776)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.put/get-macro", value: __defun__shen_4put_cget_1macro})

__defun__shen_4function_1macro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1778 := __args[0]
_ = V1778
reg310072 := PrimIsPair(V1778)
var reg310096 Obj
if reg310072 == True {
reg310073 := MakeSymbol("function")
reg310074 := PrimHead(V1778)
reg310075 := PrimEqual(reg310073, reg310074)
var reg310091 Obj
if reg310075 == True {
reg310076 := PrimTail(V1778)
reg310077 := PrimIsPair(reg310076)
var reg310086 Obj
if reg310077 == True {
reg310078 := Nil;
reg310079 := PrimTail(V1778)
reg310080 := PrimTail(reg310079)
reg310081 := PrimEqual(reg310078, reg310080)
var reg310084 Obj
if reg310081 == True {
reg310082 := True;
reg310084 = reg310082
} else {
reg310083 := False;
reg310084 = reg310083
}
reg310086 = reg310084
} else {
reg310085 := False;
reg310086 = reg310085
}
var reg310089 Obj
if reg310086 == True {
reg310087 := True;
reg310089 = reg310087
} else {
reg310088 := False;
reg310089 = reg310088
}
reg310091 = reg310089
} else {
reg310090 := False;
reg310091 = reg310090
}
var reg310094 Obj
if reg310091 == True {
reg310092 := True;
reg310094 = reg310092
} else {
reg310093 := False;
reg310094 = reg310093
}
reg310096 = reg310094
} else {
reg310095 := False;
reg310096 = reg310095
}
if reg310096 == True {
reg310097 := PrimTail(V1778)
reg310098 := PrimHead(reg310097)
reg310099 := PrimTail(V1778)
reg310100 := PrimHead(reg310099)
reg310101 := __e.Call(__defun__arity, reg310100)
__ctx.TailApply(__defun__shen_4function_1abstraction, reg310098, reg310101)
return
} else {
__ctx.Return(V1778)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.function-macro", value: __defun__shen_4function_1macro})

__defun__shen_4function_1abstraction = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1781 := __args[0]
_ = V1781
V1782 := __args[1]
_ = V1782
reg310103 := MakeNumber(0)
reg310104 := PrimEqual(reg310103, V1782)
if reg310104 == True {
reg310105 := MakeString(" has no lambda form\n")
reg310106 := MakeSymbol("shen.a")
reg310107 := __e.Call(__defun__shen_4app, V1781, reg310105, reg310106)
reg310108 := PrimSimpleError(reg310107)
__ctx.Return(reg310108)
return
} else {
reg310109 := MakeNumber(-1)
reg310110 := PrimEqual(reg310109, V1782)
if reg310110 == True {
reg310111 := MakeSymbol("function")
reg310112 := Nil;
reg310113 := PrimCons(V1781, reg310112)
reg310114 := PrimCons(reg310111, reg310113)
__ctx.Return(reg310114)
return
} else {
reg310115 := Nil;
__ctx.TailApply(__defun__shen_4function_1abstraction_1help, V1781, V1782, reg310115)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.function-abstraction", value: __defun__shen_4function_1abstraction})

__defun__shen_4function_1abstraction_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1786 := __args[0]
_ = V1786
V1787 := __args[1]
_ = V1787
V1788 := __args[2]
_ = V1788
reg310117 := MakeNumber(0)
reg310118 := PrimEqual(reg310117, V1787)
if reg310118 == True {
reg310119 := PrimCons(V1786, V1788)
__ctx.Return(reg310119)
return
} else {
reg310120 := MakeSymbol("V")
reg310121 := __e.Call(__defun__gensym, reg310120)
X := reg310121
_ = X
reg310122 := MakeSymbol("/.")
reg310123 := MakeNumber(1)
reg310124 := PrimNumberSubtract(V1787, reg310123)
reg310125 := Nil;
reg310126 := PrimCons(X, reg310125)
reg310127 := __e.Call(__defun__append, V1788, reg310126)
reg310128 := __e.Call(__defun__shen_4function_1abstraction_1help, V1786, reg310124, reg310127)
reg310129 := Nil;
reg310130 := PrimCons(reg310128, reg310129)
reg310131 := PrimCons(X, reg310130)
reg310132 := PrimCons(reg310122, reg310131)
__ctx.Return(reg310132)
return
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.function-abstraction-help", value: __defun__shen_4function_1abstraction_1help})

__defun__undefmacro = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1790 := __args[0]
_ = V1790
reg310133 := MakeSymbol("shen.*macroreg*")
reg310134 := PrimValue(reg310133)
MacroReg := reg310134
_ = MacroReg
reg310135 := __e.Call(__defun__shen_4findpos, V1790, MacroReg)
Pos := reg310135
_ = Pos
reg310136 := MakeSymbol("shen.*macroreg*")
reg310137 := __e.Call(__defun__remove, V1790, MacroReg)
reg310138 := PrimSet(reg310136, reg310137)
Remove1 := reg310138
_ = Remove1
reg310139 := MakeSymbol("*macros*")
reg310140 := MakeSymbol("*macros*")
reg310141 := PrimValue(reg310140)
reg310142 := __e.Call(__defun__shen_4remove_1nth, Pos, reg310141)
reg310143 := PrimSet(reg310139, reg310142)
Remove2 := reg310143
_ = Remove2
__ctx.Return(V1790)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "undefmacro", value: __defun__undefmacro})

__defun__shen_4findpos = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1800 := __args[0]
_ = V1800
V1801 := __args[1]
_ = V1801
reg310144 := Nil;
reg310145 := PrimEqual(reg310144, V1801)
if reg310145 == True {
reg310146 := MakeString(" is not a macro\n")
reg310147 := MakeSymbol("shen.a")
reg310148 := __e.Call(__defun__shen_4app, V1800, reg310146, reg310147)
reg310149 := PrimSimpleError(reg310148)
__ctx.Return(reg310149)
return
} else {
reg310150 := PrimIsPair(V1801)
var reg310157 Obj
if reg310150 == True {
reg310151 := PrimHead(V1801)
reg310152 := PrimEqual(reg310151, V1800)
var reg310155 Obj
if reg310152 == True {
reg310153 := True;
reg310155 = reg310153
} else {
reg310154 := False;
reg310155 = reg310154
}
reg310157 = reg310155
} else {
reg310156 := False;
reg310157 = reg310156
}
if reg310157 == True {
reg310158 := MakeNumber(1)
__ctx.Return(reg310158)
return
} else {
reg310159 := PrimIsPair(V1801)
if reg310159 == True {
reg310160 := MakeNumber(1)
reg310161 := PrimTail(V1801)
reg310162 := __e.Call(__defun__shen_4findpos, V1800, reg310161)
reg310163 := PrimNumberAdd(reg310160, reg310162)
__ctx.Return(reg310163)
return
} else {
reg310164 := MakeSymbol("shen.findpos")
__ctx.TailApply(__defun__shen_4f__error, reg310164)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.findpos", value: __defun__shen_4findpos})

__defun__shen_4remove_1nth = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V1806 := __args[0]
_ = V1806
V1807 := __args[1]
_ = V1807
reg310166 := MakeNumber(1)
reg310167 := PrimEqual(reg310166, V1806)
var reg310173 Obj
if reg310167 == True {
reg310168 := PrimIsPair(V1807)
var reg310171 Obj
if reg310168 == True {
reg310169 := True;
reg310171 = reg310169
} else {
reg310170 := False;
reg310171 = reg310170
}
reg310173 = reg310171
} else {
reg310172 := False;
reg310173 = reg310172
}
if reg310173 == True {
reg310174 := PrimTail(V1807)
__ctx.Return(reg310174)
return
} else {
reg310175 := PrimIsPair(V1807)
if reg310175 == True {
reg310176 := PrimHead(V1807)
reg310177 := MakeNumber(1)
reg310178 := PrimNumberSubtract(V1806, reg310177)
reg310179 := PrimTail(V1807)
reg310180 := __e.Call(__defun__shen_4remove_1nth, reg310178, reg310179)
reg310181 := PrimCons(reg310176, reg310180)
__ctx.Return(reg310181)
return
} else {
reg310182 := MakeSymbol("shen.remove-nth")
__ctx.TailApply(__defun__shen_4f__error, reg310182)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.remove-nth", value: __defun__shen_4remove_1nth})

}
