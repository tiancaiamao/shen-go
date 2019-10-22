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
reg311795 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg311795)
return
}, 0))
__defun__shen_4typecheck = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3644 := __args[0]
_ = V3644
V3645 := __args[1]
_ = V3645
reg311796 := __e.Call(__defun__shen_4curry, V3644)
Curry := reg311796
_ = Curry
reg311797 := __e.Call(__defun__shen_4start_1new_1prolog_1process)
ProcessN := reg311797
_ = ProcessN
reg311798 := __e.Call(__defun__shen_4curry_1type, V3645)
reg311799 := __e.Call(__defun__shen_4demodulate, reg311798)
reg311800 := __e.Call(__defun__shen_4insert_1prolog_1variables, reg311799, ProcessN)
Type := reg311800
_ = Type
reg311801 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg311802 := MakeSymbol("shen.void")
__ctx.TailApply(__defun__return, Type, ProcessN, reg311802)
return
}, 0)
Continuation := reg311801
_ = Continuation
reg311804 := MakeSymbol(":")
reg311805 := Nil;
reg311806 := PrimCons(Type, reg311805)
reg311807 := PrimCons(reg311804, reg311806)
reg311808 := PrimCons(Curry, reg311807)
reg311809 := Nil;
__ctx.TailApply(__defun__shen_4t_d, reg311808, reg311809, ProcessN, Continuation)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.typecheck", value: __defun__shen_4typecheck})

__defun__shen_4curry = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3647 := __args[0]
_ = V3647
reg311811 := PrimIsPair(V3647)
var reg311818 Obj
if reg311811 == True {
reg311812 := PrimHead(V3647)
reg311813 := __e.Call(__defun__shen_4special_2, reg311812)
var reg311816 Obj
if reg311813 == True {
reg311814 := True;
reg311816 = reg311814
} else {
reg311815 := False;
reg311816 = reg311815
}
reg311818 = reg311816
} else {
reg311817 := False;
reg311818 = reg311817
}
if reg311818 == True {
reg311819 := PrimHead(V3647)
reg311820 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4curry, Y)
return
}, 1)
reg311822 := PrimTail(V3647)
reg311823 := __e.Call(__defun__map, reg311820, reg311822)
reg311824 := PrimCons(reg311819, reg311823)
__ctx.Return(reg311824)
return
} else {
reg311825 := PrimIsPair(V3647)
var reg311839 Obj
if reg311825 == True {
reg311826 := PrimTail(V3647)
reg311827 := PrimIsPair(reg311826)
var reg311834 Obj
if reg311827 == True {
reg311828 := PrimHead(V3647)
reg311829 := __e.Call(__defun__shen_4extraspecial_2, reg311828)
var reg311832 Obj
if reg311829 == True {
reg311830 := True;
reg311832 = reg311830
} else {
reg311831 := False;
reg311832 = reg311831
}
reg311834 = reg311832
} else {
reg311833 := False;
reg311834 = reg311833
}
var reg311837 Obj
if reg311834 == True {
reg311835 := True;
reg311837 = reg311835
} else {
reg311836 := False;
reg311837 = reg311836
}
reg311839 = reg311837
} else {
reg311838 := False;
reg311839 = reg311838
}
if reg311839 == True {
__ctx.Return(V3647)
return
} else {
reg311840 := PrimIsPair(V3647)
var reg311873 Obj
if reg311840 == True {
reg311841 := MakeSymbol("type")
reg311842 := PrimHead(V3647)
reg311843 := PrimEqual(reg311841, reg311842)
var reg311868 Obj
if reg311843 == True {
reg311844 := PrimTail(V3647)
reg311845 := PrimIsPair(reg311844)
var reg311863 Obj
if reg311845 == True {
reg311846 := PrimTail(V3647)
reg311847 := PrimTail(reg311846)
reg311848 := PrimIsPair(reg311847)
var reg311858 Obj
if reg311848 == True {
reg311849 := Nil;
reg311850 := PrimTail(V3647)
reg311851 := PrimTail(reg311850)
reg311852 := PrimTail(reg311851)
reg311853 := PrimEqual(reg311849, reg311852)
var reg311856 Obj
if reg311853 == True {
reg311854 := True;
reg311856 = reg311854
} else {
reg311855 := False;
reg311856 = reg311855
}
reg311858 = reg311856
} else {
reg311857 := False;
reg311858 = reg311857
}
var reg311861 Obj
if reg311858 == True {
reg311859 := True;
reg311861 = reg311859
} else {
reg311860 := False;
reg311861 = reg311860
}
reg311863 = reg311861
} else {
reg311862 := False;
reg311863 = reg311862
}
var reg311866 Obj
if reg311863 == True {
reg311864 := True;
reg311866 = reg311864
} else {
reg311865 := False;
reg311866 = reg311865
}
reg311868 = reg311866
} else {
reg311867 := False;
reg311868 = reg311867
}
var reg311871 Obj
if reg311868 == True {
reg311869 := True;
reg311871 = reg311869
} else {
reg311870 := False;
reg311871 = reg311870
}
reg311873 = reg311871
} else {
reg311872 := False;
reg311873 = reg311872
}
if reg311873 == True {
reg311874 := MakeSymbol("type")
reg311875 := PrimTail(V3647)
reg311876 := PrimHead(reg311875)
reg311877 := __e.Call(__defun__shen_4curry, reg311876)
reg311878 := PrimTail(V3647)
reg311879 := PrimTail(reg311878)
reg311880 := PrimCons(reg311877, reg311879)
reg311881 := PrimCons(reg311874, reg311880)
__ctx.Return(reg311881)
return
} else {
reg311882 := PrimIsPair(V3647)
var reg311897 Obj
if reg311882 == True {
reg311883 := PrimTail(V3647)
reg311884 := PrimIsPair(reg311883)
var reg311892 Obj
if reg311884 == True {
reg311885 := PrimTail(V3647)
reg311886 := PrimTail(reg311885)
reg311887 := PrimIsPair(reg311886)
var reg311890 Obj
if reg311887 == True {
reg311888 := True;
reg311890 = reg311888
} else {
reg311889 := False;
reg311890 = reg311889
}
reg311892 = reg311890
} else {
reg311891 := False;
reg311892 = reg311891
}
var reg311895 Obj
if reg311892 == True {
reg311893 := True;
reg311895 = reg311893
} else {
reg311894 := False;
reg311895 = reg311894
}
reg311897 = reg311895
} else {
reg311896 := False;
reg311897 = reg311896
}
if reg311897 == True {
reg311898 := PrimHead(V3647)
reg311899 := PrimTail(V3647)
reg311900 := PrimHead(reg311899)
reg311901 := Nil;
reg311902 := PrimCons(reg311900, reg311901)
reg311903 := PrimCons(reg311898, reg311902)
reg311904 := PrimTail(V3647)
reg311905 := PrimTail(reg311904)
reg311906 := PrimCons(reg311903, reg311905)
__ctx.TailApply(__defun__shen_4curry, reg311906)
return
} else {
reg311908 := PrimIsPair(V3647)
var reg311924 Obj
if reg311908 == True {
reg311909 := PrimTail(V3647)
reg311910 := PrimIsPair(reg311909)
var reg311919 Obj
if reg311910 == True {
reg311911 := Nil;
reg311912 := PrimTail(V3647)
reg311913 := PrimTail(reg311912)
reg311914 := PrimEqual(reg311911, reg311913)
var reg311917 Obj
if reg311914 == True {
reg311915 := True;
reg311917 = reg311915
} else {
reg311916 := False;
reg311917 = reg311916
}
reg311919 = reg311917
} else {
reg311918 := False;
reg311919 = reg311918
}
var reg311922 Obj
if reg311919 == True {
reg311920 := True;
reg311922 = reg311920
} else {
reg311921 := False;
reg311922 = reg311921
}
reg311924 = reg311922
} else {
reg311923 := False;
reg311924 = reg311923
}
if reg311924 == True {
reg311925 := PrimHead(V3647)
reg311926 := __e.Call(__defun__shen_4curry, reg311925)
reg311927 := PrimTail(V3647)
reg311928 := PrimHead(reg311927)
reg311929 := __e.Call(__defun__shen_4curry, reg311928)
reg311930 := Nil;
reg311931 := PrimCons(reg311929, reg311930)
reg311932 := PrimCons(reg311926, reg311931)
__ctx.Return(reg311932)
return
} else {
__ctx.Return(V3647)
return
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.curry", value: __defun__shen_4curry})

__defun__shen_4special_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3649 := __args[0]
_ = V3649
reg311933 := MakeSymbol("shen.*special*")
reg311934 := PrimValue(reg311933)
__ctx.TailApply(__defun__element_2, V3649, reg311934)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.special?", value: __defun__shen_4special_2})

__defun__shen_4extraspecial_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3651 := __args[0]
_ = V3651
reg311936 := MakeSymbol("shen.*extraspecial*")
reg311937 := PrimValue(reg311936)
__ctx.TailApply(__defun__element_2, V3651, reg311937)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.extraspecial?", value: __defun__shen_4extraspecial_2})

__defun__shen_4t_d = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3656 := __args[0]
_ = V3656
V3657 := __args[1]
_ = V3657
V3658 := __args[2]
_ = V3658
V3659 := __args[3]
_ = V3659
reg311939 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg311939
_ = Throwcontrol
reg311940 := __e.Call(__defun__shen_4newpv, V3658)
Error := reg311940
_ = Error
reg311941 := __e.Call(__defun__shen_4incinfs)
_ = reg311941
reg311942 := __e.Call(__defun__shen_4maxinfexceeded_2)
reg311943 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg311944 := __e.Call(__defun__shen_4errormaxinfs)
__ctx.TailApply(__defun__bind, Error, reg311944, V3658, V3659)
return
}, 0)
reg311946 := __e.Call(__defun__fwhen, reg311942, V3658, reg311943)
Case := reg311946
_ = Case
reg311947 := False;
reg311948 := PrimEqual(Case, reg311947)
var reg312008 Obj
if reg311948 == True {
reg311949 := __e.Call(__defun__shen_4lazyderef, V3656, V3658)
V3636 := reg311949
_ = V3636
reg311950 := MakeSymbol("fail")
reg311951 := PrimEqual(reg311950, V3636)
var reg311957 Obj
if reg311951 == True {
reg311952 := __e.Call(__defun__shen_4incinfs)
_ = reg311952
reg311953 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4prolog_1failure, V3658, V3659)
return
}, 0)
reg311955 := __e.Call(__defun__cut, Throwcontrol, V3658, reg311953)
reg311957 = reg311955
} else {
reg311956 := False;
reg311957 = reg311956
}
Case := reg311957
_ = Case
reg311958 := False;
reg311959 := PrimEqual(Case, reg311958)
var reg312007 Obj
if reg311959 == True {
reg311960 := __e.Call(__defun__shen_4lazyderef, V3656, V3658)
V3637 := reg311960
_ = V3637
reg311961 := PrimIsPair(V3637)
var reg311994 Obj
if reg311961 == True {
reg311962 := PrimHead(V3637)
X := reg311962
_ = X
reg311963 := PrimTail(V3637)
reg311964 := __e.Call(__defun__shen_4lazyderef, reg311963, V3658)
V3638 := reg311964
_ = V3638
reg311965 := PrimIsPair(V3638)
var reg311992 Obj
if reg311965 == True {
reg311966 := PrimHead(V3638)
reg311967 := __e.Call(__defun__shen_4lazyderef, reg311966, V3658)
V3639 := reg311967
_ = V3639
reg311968 := MakeSymbol(":")
reg311969 := PrimEqual(reg311968, V3639)
var reg311990 Obj
if reg311969 == True {
reg311970 := PrimTail(V3638)
reg311971 := __e.Call(__defun__shen_4lazyderef, reg311970, V3658)
V3640 := reg311971
_ = V3640
reg311972 := PrimIsPair(V3640)
var reg311988 Obj
if reg311972 == True {
reg311973 := PrimHead(V3640)
A := reg311973
_ = A
reg311974 := PrimTail(V3640)
reg311975 := __e.Call(__defun__shen_4lazyderef, reg311974, V3658)
V3641 := reg311975
_ = V3641
reg311976 := Nil;
reg311977 := PrimEqual(reg311976, V3641)
var reg311986 Obj
if reg311977 == True {
reg311978 := __e.Call(__defun__shen_4incinfs)
_ = reg311978
reg311979 := __e.Call(__defun__shen_4type_1theory_1enabled_2)
reg311980 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg311981 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, X, A, V3657, V3658, V3659)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3658, reg311981)
return
}, 0)
reg311984 := __e.Call(__defun__fwhen, reg311979, V3658, reg311980)
reg311986 = reg311984
} else {
reg311985 := False;
reg311986 = reg311985
}
reg311988 = reg311986
} else {
reg311987 := False;
reg311988 = reg311987
}
reg311990 = reg311988
} else {
reg311989 := False;
reg311990 = reg311989
}
reg311992 = reg311990
} else {
reg311991 := False;
reg311992 = reg311991
}
reg311994 = reg311992
} else {
reg311993 := False;
reg311994 = reg311993
}
Case := reg311994
_ = Case
reg311995 := False;
reg311996 := PrimEqual(Case, reg311995)
var reg312006 Obj
if reg311996 == True {
reg311997 := __e.Call(__defun__shen_4newpv, V3658)
Datatypes := reg311997
_ = Datatypes
reg311998 := __e.Call(__defun__shen_4incinfs)
_ = reg311998
reg311999 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312000 := MakeSymbol("shen.*datatypes*")
reg312001 := PrimValue(reg312000)
reg312002 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4udefs_d, V3656, V3657, Datatypes, V3658, V3659)
return
}, 0)
__ctx.TailApply(__defun__bind, Datatypes, reg312001, V3658, reg312002)
return
}, 0)
reg312005 := __e.Call(__defun__shen_4show, V3656, V3657, V3658, reg311999)
reg312006 = reg312005
} else {
reg312006 = Case
}
reg312007 = reg312006
} else {
reg312007 = Case
}
reg312008 = reg312007
} else {
reg312008 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg312008)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.t*", value: __defun__shen_4t_d})

__defun__shen_4type_1theory_1enabled_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312010 := MakeSymbol("shen.*shen-type-theory-enabled?*")
reg312011 := PrimValue(reg312010)
__ctx.Return(reg312011)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.type-theory-enabled?", value: __defun__shen_4type_1theory_1enabled_2})

__defun__enable_1type_1theory = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3665 := __args[0]
_ = V3665
reg312012 := MakeSymbol("+")
reg312013 := PrimEqual(reg312012, V3665)
if reg312013 == True {
reg312014 := MakeSymbol("shen.*shen-type-theory-enabled?*")
reg312015 := True;
reg312016 := PrimSet(reg312014, reg312015)
__ctx.Return(reg312016)
return
} else {
reg312017 := MakeSymbol("-")
reg312018 := PrimEqual(reg312017, V3665)
if reg312018 == True {
reg312019 := MakeSymbol("shen.*shen-type-theory-enabled?*")
reg312020 := False;
reg312021 := PrimSet(reg312019, reg312020)
__ctx.Return(reg312021)
return
} else {
reg312022 := MakeString("enable-type-theory expects a + or a -\n")
reg312023 := PrimSimpleError(reg312022)
__ctx.Return(reg312023)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "enable-type-theory", value: __defun__enable_1type_1theory})

__defun__shen_4prolog_1failure = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3676 := __args[0]
_ = V3676
V3677 := __args[1]
_ = V3677
reg312024 := False;
__ctx.Return(reg312024)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.prolog-failure", value: __defun__shen_4prolog_1failure})

__defun__shen_4maxinfexceeded_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312025 := __e.Call(__defun__inferences)
reg312026 := MakeSymbol("shen.*maxinferences*")
reg312027 := PrimValue(reg312026)
reg312028 := PrimGreatThan(reg312025, reg312027)
__ctx.Return(reg312028)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.maxinfexceeded?", value: __defun__shen_4maxinfexceeded_2})

__defun__shen_4errormaxinfs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312029 := MakeString("maximum inferences exceeded~%")
reg312030 := PrimSimpleError(reg312029)
__ctx.Return(reg312030)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.errormaxinfs", value: __defun__shen_4errormaxinfs})

__defun__shen_4udefs_d = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3683 := __args[0]
_ = V3683
V3684 := __args[1]
_ = V3684
V3685 := __args[2]
_ = V3685
V3686 := __args[3]
_ = V3686
V3687 := __args[4]
_ = V3687
reg312031 := __e.Call(__defun__shen_4lazyderef, V3685, V3686)
V3632 := reg312031
_ = V3632
reg312032 := PrimIsPair(V3632)
var reg312041 Obj
if reg312032 == True {
reg312033 := PrimHead(V3632)
D := reg312033
_ = D
reg312034 := __e.Call(__defun__shen_4incinfs)
_ = reg312034
reg312035 := Nil;
reg312036 := PrimCons(V3684, reg312035)
reg312037 := PrimCons(V3683, reg312036)
reg312038 := PrimCons(D, reg312037)
reg312039 := __e.Call(__defun__call, reg312038, V3686, V3687)
reg312041 = reg312039
} else {
reg312040 := False;
reg312041 = reg312040
}
Case := reg312041
_ = Case
reg312042 := False;
reg312043 := PrimEqual(Case, reg312042)
if reg312043 == True {
reg312044 := __e.Call(__defun__shen_4lazyderef, V3685, V3686)
V3633 := reg312044
_ = V3633
reg312045 := PrimIsPair(V3633)
if reg312045 == True {
reg312046 := PrimTail(V3633)
Ds := reg312046
_ = Ds
reg312047 := __e.Call(__defun__shen_4incinfs)
_ = reg312047
__ctx.TailApply(__defun__shen_4udefs_d, V3683, V3684, Ds, V3686, V3687)
return
} else {
reg312049 := False;
__ctx.Return(reg312049)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.udefs*", value: __defun__shen_4udefs_d})

__defun__shen_4th_d = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3693 := __args[0]
_ = V3693
V3694 := __args[1]
_ = V3694
V3695 := __args[2]
_ = V3695
V3696 := __args[3]
_ = V3696
V3697 := __args[4]
_ = V3697
reg312050 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg312050
_ = Throwcontrol
reg312051 := __e.Call(__defun__shen_4incinfs)
_ = reg312051
reg312052 := MakeSymbol(":")
reg312053 := Nil;
reg312054 := PrimCons(V3694, reg312053)
reg312055 := PrimCons(reg312052, reg312054)
reg312056 := PrimCons(V3693, reg312055)
reg312057 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312058 := False;
__ctx.TailApply(__defun__fwhen, reg312058, V3696, V3697)
return
}, 0)
reg312060 := __e.Call(__defun__shen_4show, reg312056, V3695, V3696, reg312057)
Case := reg312060
_ = Case
reg312061 := False;
reg312062 := PrimEqual(Case, reg312061)
var reg313481 Obj
if reg312062 == True {
reg312063 := __e.Call(__defun__shen_4newpv, V3696)
F := reg312063
_ = F
reg312064 := __e.Call(__defun__shen_4incinfs)
_ = reg312064
reg312065 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
reg312066 := __e.Call(__defun__shen_4typedf_2, reg312065)
reg312067 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312068 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
reg312069 := __e.Call(__defun__shen_4sigf, reg312068)
reg312070 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312071 := Nil;
reg312072 := PrimCons(V3694, reg312071)
reg312073 := PrimCons(F, reg312072)
__ctx.TailApply(__defun__call, reg312073, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, F, reg312069, V3696, reg312070)
return
}, 0)
reg312076 := __e.Call(__defun__fwhen, reg312066, V3696, reg312067)
Case := reg312076
_ = Case
reg312077 := False;
reg312078 := PrimEqual(Case, reg312077)
var reg313480 Obj
if reg312078 == True {
reg312079 := __e.Call(__defun__shen_4incinfs)
_ = reg312079
reg312080 := __e.Call(__defun__shen_4base, V3693, V3694, V3696, V3697)
Case := reg312080
_ = Case
reg312081 := False;
reg312082 := PrimEqual(Case, reg312081)
var reg313479 Obj
if reg312082 == True {
reg312083 := __e.Call(__defun__shen_4incinfs)
_ = reg312083
reg312084 := __e.Call(__defun__shen_4by__hypothesis, V3693, V3694, V3695, V3696, V3697)
Case := reg312084
_ = Case
reg312085 := False;
reg312086 := PrimEqual(Case, reg312085)
var reg313478 Obj
if reg312086 == True {
reg312087 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3528 := reg312087
_ = V3528
reg312088 := PrimIsPair(V3528)
var reg312103 Obj
if reg312088 == True {
reg312089 := PrimHead(V3528)
F := reg312089
_ = F
reg312090 := PrimTail(V3528)
reg312091 := __e.Call(__defun__shen_4lazyderef, reg312090, V3696)
V3529 := reg312091
_ = V3529
reg312092 := Nil;
reg312093 := PrimEqual(reg312092, V3529)
var reg312101 Obj
if reg312093 == True {
reg312094 := __e.Call(__defun__shen_4incinfs)
_ = reg312094
reg312095 := MakeSymbol("-->")
reg312096 := Nil;
reg312097 := PrimCons(V3694, reg312096)
reg312098 := PrimCons(reg312095, reg312097)
reg312099 := __e.Call(__defun__shen_4th_d, F, reg312098, V3695, V3696, V3697)
reg312101 = reg312099
} else {
reg312100 := False;
reg312101 = reg312100
}
reg312103 = reg312101
} else {
reg312102 := False;
reg312103 = reg312102
}
Case := reg312103
_ = Case
reg312104 := False;
reg312105 := PrimEqual(Case, reg312104)
var reg313477 Obj
if reg312105 == True {
reg312106 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3530 := reg312106
_ = V3530
reg312107 := PrimIsPair(V3530)
var reg312132 Obj
if reg312107 == True {
reg312108 := PrimHead(V3530)
F := reg312108
_ = F
reg312109 := PrimTail(V3530)
reg312110 := __e.Call(__defun__shen_4lazyderef, reg312109, V3696)
V3531 := reg312110
_ = V3531
reg312111 := PrimIsPair(V3531)
var reg312130 Obj
if reg312111 == True {
reg312112 := PrimHead(V3531)
X := reg312112
_ = X
reg312113 := PrimTail(V3531)
reg312114 := __e.Call(__defun__shen_4lazyderef, reg312113, V3696)
V3532 := reg312114
_ = V3532
reg312115 := Nil;
reg312116 := PrimEqual(reg312115, V3532)
var reg312128 Obj
if reg312116 == True {
reg312117 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312117
_ = B
reg312118 := __e.Call(__defun__shen_4incinfs)
_ = reg312118
reg312119 := MakeSymbol("-->")
reg312120 := Nil;
reg312121 := PrimCons(V3694, reg312120)
reg312122 := PrimCons(reg312119, reg312121)
reg312123 := PrimCons(B, reg312122)
reg312124 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, X, B, V3695, V3696, V3697)
return
}, 0)
reg312126 := __e.Call(__defun__shen_4th_d, F, reg312123, V3695, V3696, reg312124)
reg312128 = reg312126
} else {
reg312127 := False;
reg312128 = reg312127
}
reg312130 = reg312128
} else {
reg312129 := False;
reg312130 = reg312129
}
reg312132 = reg312130
} else {
reg312131 := False;
reg312132 = reg312131
}
Case := reg312132
_ = Case
reg312133 := False;
reg312134 := PrimEqual(Case, reg312133)
var reg313476 Obj
if reg312134 == True {
reg312135 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3533 := reg312135
_ = V3533
reg312136 := PrimIsPair(V3533)
var reg312290 Obj
if reg312136 == True {
reg312137 := PrimHead(V3533)
reg312138 := __e.Call(__defun__shen_4lazyderef, reg312137, V3696)
V3534 := reg312138
_ = V3534
reg312139 := MakeSymbol("cons")
reg312140 := PrimEqual(reg312139, V3534)
var reg312288 Obj
if reg312140 == True {
reg312141 := PrimTail(V3533)
reg312142 := __e.Call(__defun__shen_4lazyderef, reg312141, V3696)
V3535 := reg312142
_ = V3535
reg312143 := PrimIsPair(V3535)
var reg312286 Obj
if reg312143 == True {
reg312144 := PrimHead(V3535)
X := reg312144
_ = X
reg312145 := PrimTail(V3535)
reg312146 := __e.Call(__defun__shen_4lazyderef, reg312145, V3696)
V3536 := reg312146
_ = V3536
reg312147 := PrimIsPair(V3536)
var reg312284 Obj
if reg312147 == True {
reg312148 := PrimHead(V3536)
Y := reg312148
_ = Y
reg312149 := PrimTail(V3536)
reg312150 := __e.Call(__defun__shen_4lazyderef, reg312149, V3696)
V3537 := reg312150
_ = V3537
reg312151 := Nil;
reg312152 := PrimEqual(reg312151, V3537)
var reg312282 Obj
if reg312152 == True {
reg312153 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3538 := reg312153
_ = V3538
reg312154 := PrimIsPair(V3538)
var reg312280 Obj
if reg312154 == True {
reg312155 := PrimHead(V3538)
reg312156 := __e.Call(__defun__shen_4lazyderef, reg312155, V3696)
V3539 := reg312156
_ = V3539
reg312157 := MakeSymbol("list")
reg312158 := PrimEqual(reg312157, V3539)
var reg312261 Obj
if reg312158 == True {
reg312159 := PrimTail(V3538)
reg312160 := __e.Call(__defun__shen_4lazyderef, reg312159, V3696)
V3540 := reg312160
_ = V3540
reg312161 := PrimIsPair(V3540)
var reg312206 Obj
if reg312161 == True {
reg312162 := PrimHead(V3540)
A := reg312162
_ = A
reg312163 := PrimTail(V3540)
reg312164 := __e.Call(__defun__shen_4lazyderef, reg312163, V3696)
V3541 := reg312164
_ = V3541
reg312165 := Nil;
reg312166 := PrimEqual(reg312165, V3541)
var reg312189 Obj
if reg312166 == True {
reg312167 := __e.Call(__defun__shen_4incinfs)
_ = reg312167
reg312168 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312169 := MakeSymbol("list")
reg312170 := Nil;
reg312171 := PrimCons(A, reg312170)
reg312172 := PrimCons(reg312169, reg312171)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312172, V3695, V3696, V3697)
return
}, 0)
reg312174 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312168)
reg312189 = reg312174
} else {
reg312175 := __e.Call(__defun__shen_4pvar_2, V3541)
var reg312188 Obj
if reg312175 == True {
reg312176 := Nil;
reg312177 := __e.Call(__defun__shen_4bindv, V3541, reg312176, V3696)
_ = reg312177
reg312178 := __e.Call(__defun__shen_4incinfs)
_ = reg312178
reg312179 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312180 := MakeSymbol("list")
reg312181 := Nil;
reg312182 := PrimCons(A, reg312181)
reg312183 := PrimCons(reg312180, reg312182)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312183, V3695, V3696, V3697)
return
}, 0)
reg312185 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312179)
Result := reg312185
_ = Result
reg312186 := __e.Call(__defun__shen_4unbindv, V3541, V3696)
_ = reg312186
reg312188 = Result
} else {
reg312187 := False;
reg312188 = reg312187
}
reg312189 = reg312188
}
reg312206 = reg312189
} else {
reg312190 := __e.Call(__defun__shen_4pvar_2, V3540)
var reg312205 Obj
if reg312190 == True {
reg312191 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312191
_ = A
reg312192 := Nil;
reg312193 := PrimCons(A, reg312192)
reg312194 := __e.Call(__defun__shen_4bindv, V3540, reg312193, V3696)
_ = reg312194
reg312195 := __e.Call(__defun__shen_4incinfs)
_ = reg312195
reg312196 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312197 := MakeSymbol("list")
reg312198 := Nil;
reg312199 := PrimCons(A, reg312198)
reg312200 := PrimCons(reg312197, reg312199)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312200, V3695, V3696, V3697)
return
}, 0)
reg312202 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312196)
Result := reg312202
_ = Result
reg312203 := __e.Call(__defun__shen_4unbindv, V3540, V3696)
_ = reg312203
reg312205 = Result
} else {
reg312204 := False;
reg312205 = reg312204
}
reg312206 = reg312205
}
reg312261 = reg312206
} else {
reg312207 := __e.Call(__defun__shen_4pvar_2, V3539)
var reg312260 Obj
if reg312207 == True {
reg312208 := MakeSymbol("list")
reg312209 := __e.Call(__defun__shen_4bindv, V3539, reg312208, V3696)
_ = reg312209
reg312210 := PrimTail(V3538)
reg312211 := __e.Call(__defun__shen_4lazyderef, reg312210, V3696)
V3542 := reg312211
_ = V3542
reg312212 := PrimIsPair(V3542)
var reg312257 Obj
if reg312212 == True {
reg312213 := PrimHead(V3542)
A := reg312213
_ = A
reg312214 := PrimTail(V3542)
reg312215 := __e.Call(__defun__shen_4lazyderef, reg312214, V3696)
V3543 := reg312215
_ = V3543
reg312216 := Nil;
reg312217 := PrimEqual(reg312216, V3543)
var reg312240 Obj
if reg312217 == True {
reg312218 := __e.Call(__defun__shen_4incinfs)
_ = reg312218
reg312219 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312220 := MakeSymbol("list")
reg312221 := Nil;
reg312222 := PrimCons(A, reg312221)
reg312223 := PrimCons(reg312220, reg312222)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312223, V3695, V3696, V3697)
return
}, 0)
reg312225 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312219)
reg312240 = reg312225
} else {
reg312226 := __e.Call(__defun__shen_4pvar_2, V3543)
var reg312239 Obj
if reg312226 == True {
reg312227 := Nil;
reg312228 := __e.Call(__defun__shen_4bindv, V3543, reg312227, V3696)
_ = reg312228
reg312229 := __e.Call(__defun__shen_4incinfs)
_ = reg312229
reg312230 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312231 := MakeSymbol("list")
reg312232 := Nil;
reg312233 := PrimCons(A, reg312232)
reg312234 := PrimCons(reg312231, reg312233)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312234, V3695, V3696, V3697)
return
}, 0)
reg312236 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312230)
Result := reg312236
_ = Result
reg312237 := __e.Call(__defun__shen_4unbindv, V3543, V3696)
_ = reg312237
reg312239 = Result
} else {
reg312238 := False;
reg312239 = reg312238
}
reg312240 = reg312239
}
reg312257 = reg312240
} else {
reg312241 := __e.Call(__defun__shen_4pvar_2, V3542)
var reg312256 Obj
if reg312241 == True {
reg312242 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312242
_ = A
reg312243 := Nil;
reg312244 := PrimCons(A, reg312243)
reg312245 := __e.Call(__defun__shen_4bindv, V3542, reg312244, V3696)
_ = reg312245
reg312246 := __e.Call(__defun__shen_4incinfs)
_ = reg312246
reg312247 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312248 := MakeSymbol("list")
reg312249 := Nil;
reg312250 := PrimCons(A, reg312249)
reg312251 := PrimCons(reg312248, reg312250)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312251, V3695, V3696, V3697)
return
}, 0)
reg312253 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312247)
Result := reg312253
_ = Result
reg312254 := __e.Call(__defun__shen_4unbindv, V3542, V3696)
_ = reg312254
reg312256 = Result
} else {
reg312255 := False;
reg312256 = reg312255
}
reg312257 = reg312256
}
Result := reg312257
_ = Result
reg312258 := __e.Call(__defun__shen_4unbindv, V3539, V3696)
_ = reg312258
reg312260 = Result
} else {
reg312259 := False;
reg312260 = reg312259
}
reg312261 = reg312260
}
reg312280 = reg312261
} else {
reg312262 := __e.Call(__defun__shen_4pvar_2, V3538)
var reg312279 Obj
if reg312262 == True {
reg312263 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312263
_ = A
reg312264 := MakeSymbol("list")
reg312265 := Nil;
reg312266 := PrimCons(A, reg312265)
reg312267 := PrimCons(reg312264, reg312266)
reg312268 := __e.Call(__defun__shen_4bindv, V3538, reg312267, V3696)
_ = reg312268
reg312269 := __e.Call(__defun__shen_4incinfs)
_ = reg312269
reg312270 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312271 := MakeSymbol("list")
reg312272 := Nil;
reg312273 := PrimCons(A, reg312272)
reg312274 := PrimCons(reg312271, reg312273)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312274, V3695, V3696, V3697)
return
}, 0)
reg312276 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312270)
Result := reg312276
_ = Result
reg312277 := __e.Call(__defun__shen_4unbindv, V3538, V3696)
_ = reg312277
reg312279 = Result
} else {
reg312278 := False;
reg312279 = reg312278
}
reg312280 = reg312279
}
reg312282 = reg312280
} else {
reg312281 := False;
reg312282 = reg312281
}
reg312284 = reg312282
} else {
reg312283 := False;
reg312284 = reg312283
}
reg312286 = reg312284
} else {
reg312285 := False;
reg312286 = reg312285
}
reg312288 = reg312286
} else {
reg312287 := False;
reg312288 = reg312287
}
reg312290 = reg312288
} else {
reg312289 := False;
reg312290 = reg312289
}
Case := reg312290
_ = Case
reg312291 := False;
reg312292 := PrimEqual(Case, reg312291)
var reg313475 Obj
if reg312292 == True {
reg312293 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3544 := reg312293
_ = V3544
reg312294 := PrimIsPair(V3544)
var reg312441 Obj
if reg312294 == True {
reg312295 := PrimHead(V3544)
reg312296 := __e.Call(__defun__shen_4lazyderef, reg312295, V3696)
V3545 := reg312296
_ = V3545
reg312297 := MakeSymbol("@p")
reg312298 := PrimEqual(reg312297, V3545)
var reg312439 Obj
if reg312298 == True {
reg312299 := PrimTail(V3544)
reg312300 := __e.Call(__defun__shen_4lazyderef, reg312299, V3696)
V3546 := reg312300
_ = V3546
reg312301 := PrimIsPair(V3546)
var reg312437 Obj
if reg312301 == True {
reg312302 := PrimHead(V3546)
X := reg312302
_ = X
reg312303 := PrimTail(V3546)
reg312304 := __e.Call(__defun__shen_4lazyderef, reg312303, V3696)
V3547 := reg312304
_ = V3547
reg312305 := PrimIsPair(V3547)
var reg312435 Obj
if reg312305 == True {
reg312306 := PrimHead(V3547)
Y := reg312306
_ = Y
reg312307 := PrimTail(V3547)
reg312308 := __e.Call(__defun__shen_4lazyderef, reg312307, V3696)
V3548 := reg312308
_ = V3548
reg312309 := Nil;
reg312310 := PrimEqual(reg312309, V3548)
var reg312433 Obj
if reg312310 == True {
reg312311 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3549 := reg312311
_ = V3549
reg312312 := PrimIsPair(V3549)
var reg312431 Obj
if reg312312 == True {
reg312313 := PrimHead(V3549)
A := reg312313
_ = A
reg312314 := PrimTail(V3549)
reg312315 := __e.Call(__defun__shen_4lazyderef, reg312314, V3696)
V3550 := reg312315
_ = V3550
reg312316 := PrimIsPair(V3550)
var reg312414 Obj
if reg312316 == True {
reg312317 := PrimHead(V3550)
reg312318 := __e.Call(__defun__shen_4lazyderef, reg312317, V3696)
V3551 := reg312318
_ = V3551
reg312319 := MakeSymbol("*")
reg312320 := PrimEqual(reg312319, V3551)
var reg312399 Obj
if reg312320 == True {
reg312321 := PrimTail(V3550)
reg312322 := __e.Call(__defun__shen_4lazyderef, reg312321, V3696)
V3552 := reg312322
_ = V3552
reg312323 := PrimIsPair(V3552)
var reg312356 Obj
if reg312323 == True {
reg312324 := PrimHead(V3552)
B := reg312324
_ = B
reg312325 := PrimTail(V3552)
reg312326 := __e.Call(__defun__shen_4lazyderef, reg312325, V3696)
V3553 := reg312326
_ = V3553
reg312327 := Nil;
reg312328 := PrimEqual(reg312327, V3553)
var reg312343 Obj
if reg312328 == True {
reg312329 := __e.Call(__defun__shen_4incinfs)
_ = reg312329
reg312330 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312332 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312330)
reg312343 = reg312332
} else {
reg312333 := __e.Call(__defun__shen_4pvar_2, V3553)
var reg312342 Obj
if reg312333 == True {
reg312334 := Nil;
reg312335 := __e.Call(__defun__shen_4bindv, V3553, reg312334, V3696)
_ = reg312335
reg312336 := __e.Call(__defun__shen_4incinfs)
_ = reg312336
reg312337 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312339 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312337)
Result := reg312339
_ = Result
reg312340 := __e.Call(__defun__shen_4unbindv, V3553, V3696)
_ = reg312340
reg312342 = Result
} else {
reg312341 := False;
reg312342 = reg312341
}
reg312343 = reg312342
}
reg312356 = reg312343
} else {
reg312344 := __e.Call(__defun__shen_4pvar_2, V3552)
var reg312355 Obj
if reg312344 == True {
reg312345 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312345
_ = B
reg312346 := Nil;
reg312347 := PrimCons(B, reg312346)
reg312348 := __e.Call(__defun__shen_4bindv, V3552, reg312347, V3696)
_ = reg312348
reg312349 := __e.Call(__defun__shen_4incinfs)
_ = reg312349
reg312350 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312352 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312350)
Result := reg312352
_ = Result
reg312353 := __e.Call(__defun__shen_4unbindv, V3552, V3696)
_ = reg312353
reg312355 = Result
} else {
reg312354 := False;
reg312355 = reg312354
}
reg312356 = reg312355
}
reg312399 = reg312356
} else {
reg312357 := __e.Call(__defun__shen_4pvar_2, V3551)
var reg312398 Obj
if reg312357 == True {
reg312358 := MakeSymbol("*")
reg312359 := __e.Call(__defun__shen_4bindv, V3551, reg312358, V3696)
_ = reg312359
reg312360 := PrimTail(V3550)
reg312361 := __e.Call(__defun__shen_4lazyderef, reg312360, V3696)
V3554 := reg312361
_ = V3554
reg312362 := PrimIsPair(V3554)
var reg312395 Obj
if reg312362 == True {
reg312363 := PrimHead(V3554)
B := reg312363
_ = B
reg312364 := PrimTail(V3554)
reg312365 := __e.Call(__defun__shen_4lazyderef, reg312364, V3696)
V3555 := reg312365
_ = V3555
reg312366 := Nil;
reg312367 := PrimEqual(reg312366, V3555)
var reg312382 Obj
if reg312367 == True {
reg312368 := __e.Call(__defun__shen_4incinfs)
_ = reg312368
reg312369 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312371 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312369)
reg312382 = reg312371
} else {
reg312372 := __e.Call(__defun__shen_4pvar_2, V3555)
var reg312381 Obj
if reg312372 == True {
reg312373 := Nil;
reg312374 := __e.Call(__defun__shen_4bindv, V3555, reg312373, V3696)
_ = reg312374
reg312375 := __e.Call(__defun__shen_4incinfs)
_ = reg312375
reg312376 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312378 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312376)
Result := reg312378
_ = Result
reg312379 := __e.Call(__defun__shen_4unbindv, V3555, V3696)
_ = reg312379
reg312381 = Result
} else {
reg312380 := False;
reg312381 = reg312380
}
reg312382 = reg312381
}
reg312395 = reg312382
} else {
reg312383 := __e.Call(__defun__shen_4pvar_2, V3554)
var reg312394 Obj
if reg312383 == True {
reg312384 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312384
_ = B
reg312385 := Nil;
reg312386 := PrimCons(B, reg312385)
reg312387 := __e.Call(__defun__shen_4bindv, V3554, reg312386, V3696)
_ = reg312387
reg312388 := __e.Call(__defun__shen_4incinfs)
_ = reg312388
reg312389 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312391 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312389)
Result := reg312391
_ = Result
reg312392 := __e.Call(__defun__shen_4unbindv, V3554, V3696)
_ = reg312392
reg312394 = Result
} else {
reg312393 := False;
reg312394 = reg312393
}
reg312395 = reg312394
}
Result := reg312395
_ = Result
reg312396 := __e.Call(__defun__shen_4unbindv, V3551, V3696)
_ = reg312396
reg312398 = Result
} else {
reg312397 := False;
reg312398 = reg312397
}
reg312399 = reg312398
}
reg312414 = reg312399
} else {
reg312400 := __e.Call(__defun__shen_4pvar_2, V3550)
var reg312413 Obj
if reg312400 == True {
reg312401 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312401
_ = B
reg312402 := MakeSymbol("*")
reg312403 := Nil;
reg312404 := PrimCons(B, reg312403)
reg312405 := PrimCons(reg312402, reg312404)
reg312406 := __e.Call(__defun__shen_4bindv, V3550, reg312405, V3696)
_ = reg312406
reg312407 := __e.Call(__defun__shen_4incinfs)
_ = reg312407
reg312408 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312410 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312408)
Result := reg312410
_ = Result
reg312411 := __e.Call(__defun__shen_4unbindv, V3550, V3696)
_ = reg312411
reg312413 = Result
} else {
reg312412 := False;
reg312413 = reg312412
}
reg312414 = reg312413
}
reg312431 = reg312414
} else {
reg312415 := __e.Call(__defun__shen_4pvar_2, V3549)
var reg312430 Obj
if reg312415 == True {
reg312416 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312416
_ = A
reg312417 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312417
_ = B
reg312418 := MakeSymbol("*")
reg312419 := Nil;
reg312420 := PrimCons(B, reg312419)
reg312421 := PrimCons(reg312418, reg312420)
reg312422 := PrimCons(A, reg312421)
reg312423 := __e.Call(__defun__shen_4bindv, V3549, reg312422, V3696)
_ = reg312423
reg312424 := __e.Call(__defun__shen_4incinfs)
_ = reg312424
reg312425 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Y, B, V3695, V3696, V3697)
return
}, 0)
reg312427 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312425)
Result := reg312427
_ = Result
reg312428 := __e.Call(__defun__shen_4unbindv, V3549, V3696)
_ = reg312428
reg312430 = Result
} else {
reg312429 := False;
reg312430 = reg312429
}
reg312431 = reg312430
}
reg312433 = reg312431
} else {
reg312432 := False;
reg312433 = reg312432
}
reg312435 = reg312433
} else {
reg312434 := False;
reg312435 = reg312434
}
reg312437 = reg312435
} else {
reg312436 := False;
reg312437 = reg312436
}
reg312439 = reg312437
} else {
reg312438 := False;
reg312439 = reg312438
}
reg312441 = reg312439
} else {
reg312440 := False;
reg312441 = reg312440
}
Case := reg312441
_ = Case
reg312442 := False;
reg312443 := PrimEqual(Case, reg312442)
var reg313474 Obj
if reg312443 == True {
reg312444 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3556 := reg312444
_ = V3556
reg312445 := PrimIsPair(V3556)
var reg312599 Obj
if reg312445 == True {
reg312446 := PrimHead(V3556)
reg312447 := __e.Call(__defun__shen_4lazyderef, reg312446, V3696)
V3557 := reg312447
_ = V3557
reg312448 := MakeSymbol("@v")
reg312449 := PrimEqual(reg312448, V3557)
var reg312597 Obj
if reg312449 == True {
reg312450 := PrimTail(V3556)
reg312451 := __e.Call(__defun__shen_4lazyderef, reg312450, V3696)
V3558 := reg312451
_ = V3558
reg312452 := PrimIsPair(V3558)
var reg312595 Obj
if reg312452 == True {
reg312453 := PrimHead(V3558)
X := reg312453
_ = X
reg312454 := PrimTail(V3558)
reg312455 := __e.Call(__defun__shen_4lazyderef, reg312454, V3696)
V3559 := reg312455
_ = V3559
reg312456 := PrimIsPair(V3559)
var reg312593 Obj
if reg312456 == True {
reg312457 := PrimHead(V3559)
Y := reg312457
_ = Y
reg312458 := PrimTail(V3559)
reg312459 := __e.Call(__defun__shen_4lazyderef, reg312458, V3696)
V3560 := reg312459
_ = V3560
reg312460 := Nil;
reg312461 := PrimEqual(reg312460, V3560)
var reg312591 Obj
if reg312461 == True {
reg312462 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3561 := reg312462
_ = V3561
reg312463 := PrimIsPair(V3561)
var reg312589 Obj
if reg312463 == True {
reg312464 := PrimHead(V3561)
reg312465 := __e.Call(__defun__shen_4lazyderef, reg312464, V3696)
V3562 := reg312465
_ = V3562
reg312466 := MakeSymbol("vector")
reg312467 := PrimEqual(reg312466, V3562)
var reg312570 Obj
if reg312467 == True {
reg312468 := PrimTail(V3561)
reg312469 := __e.Call(__defun__shen_4lazyderef, reg312468, V3696)
V3563 := reg312469
_ = V3563
reg312470 := PrimIsPair(V3563)
var reg312515 Obj
if reg312470 == True {
reg312471 := PrimHead(V3563)
A := reg312471
_ = A
reg312472 := PrimTail(V3563)
reg312473 := __e.Call(__defun__shen_4lazyderef, reg312472, V3696)
V3564 := reg312473
_ = V3564
reg312474 := Nil;
reg312475 := PrimEqual(reg312474, V3564)
var reg312498 Obj
if reg312475 == True {
reg312476 := __e.Call(__defun__shen_4incinfs)
_ = reg312476
reg312477 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312478 := MakeSymbol("vector")
reg312479 := Nil;
reg312480 := PrimCons(A, reg312479)
reg312481 := PrimCons(reg312478, reg312480)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312481, V3695, V3696, V3697)
return
}, 0)
reg312483 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312477)
reg312498 = reg312483
} else {
reg312484 := __e.Call(__defun__shen_4pvar_2, V3564)
var reg312497 Obj
if reg312484 == True {
reg312485 := Nil;
reg312486 := __e.Call(__defun__shen_4bindv, V3564, reg312485, V3696)
_ = reg312486
reg312487 := __e.Call(__defun__shen_4incinfs)
_ = reg312487
reg312488 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312489 := MakeSymbol("vector")
reg312490 := Nil;
reg312491 := PrimCons(A, reg312490)
reg312492 := PrimCons(reg312489, reg312491)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312492, V3695, V3696, V3697)
return
}, 0)
reg312494 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312488)
Result := reg312494
_ = Result
reg312495 := __e.Call(__defun__shen_4unbindv, V3564, V3696)
_ = reg312495
reg312497 = Result
} else {
reg312496 := False;
reg312497 = reg312496
}
reg312498 = reg312497
}
reg312515 = reg312498
} else {
reg312499 := __e.Call(__defun__shen_4pvar_2, V3563)
var reg312514 Obj
if reg312499 == True {
reg312500 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312500
_ = A
reg312501 := Nil;
reg312502 := PrimCons(A, reg312501)
reg312503 := __e.Call(__defun__shen_4bindv, V3563, reg312502, V3696)
_ = reg312503
reg312504 := __e.Call(__defun__shen_4incinfs)
_ = reg312504
reg312505 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312506 := MakeSymbol("vector")
reg312507 := Nil;
reg312508 := PrimCons(A, reg312507)
reg312509 := PrimCons(reg312506, reg312508)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312509, V3695, V3696, V3697)
return
}, 0)
reg312511 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312505)
Result := reg312511
_ = Result
reg312512 := __e.Call(__defun__shen_4unbindv, V3563, V3696)
_ = reg312512
reg312514 = Result
} else {
reg312513 := False;
reg312514 = reg312513
}
reg312515 = reg312514
}
reg312570 = reg312515
} else {
reg312516 := __e.Call(__defun__shen_4pvar_2, V3562)
var reg312569 Obj
if reg312516 == True {
reg312517 := MakeSymbol("vector")
reg312518 := __e.Call(__defun__shen_4bindv, V3562, reg312517, V3696)
_ = reg312518
reg312519 := PrimTail(V3561)
reg312520 := __e.Call(__defun__shen_4lazyderef, reg312519, V3696)
V3565 := reg312520
_ = V3565
reg312521 := PrimIsPair(V3565)
var reg312566 Obj
if reg312521 == True {
reg312522 := PrimHead(V3565)
A := reg312522
_ = A
reg312523 := PrimTail(V3565)
reg312524 := __e.Call(__defun__shen_4lazyderef, reg312523, V3696)
V3566 := reg312524
_ = V3566
reg312525 := Nil;
reg312526 := PrimEqual(reg312525, V3566)
var reg312549 Obj
if reg312526 == True {
reg312527 := __e.Call(__defun__shen_4incinfs)
_ = reg312527
reg312528 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312529 := MakeSymbol("vector")
reg312530 := Nil;
reg312531 := PrimCons(A, reg312530)
reg312532 := PrimCons(reg312529, reg312531)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312532, V3695, V3696, V3697)
return
}, 0)
reg312534 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312528)
reg312549 = reg312534
} else {
reg312535 := __e.Call(__defun__shen_4pvar_2, V3566)
var reg312548 Obj
if reg312535 == True {
reg312536 := Nil;
reg312537 := __e.Call(__defun__shen_4bindv, V3566, reg312536, V3696)
_ = reg312537
reg312538 := __e.Call(__defun__shen_4incinfs)
_ = reg312538
reg312539 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312540 := MakeSymbol("vector")
reg312541 := Nil;
reg312542 := PrimCons(A, reg312541)
reg312543 := PrimCons(reg312540, reg312542)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312543, V3695, V3696, V3697)
return
}, 0)
reg312545 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312539)
Result := reg312545
_ = Result
reg312546 := __e.Call(__defun__shen_4unbindv, V3566, V3696)
_ = reg312546
reg312548 = Result
} else {
reg312547 := False;
reg312548 = reg312547
}
reg312549 = reg312548
}
reg312566 = reg312549
} else {
reg312550 := __e.Call(__defun__shen_4pvar_2, V3565)
var reg312565 Obj
if reg312550 == True {
reg312551 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312551
_ = A
reg312552 := Nil;
reg312553 := PrimCons(A, reg312552)
reg312554 := __e.Call(__defun__shen_4bindv, V3565, reg312553, V3696)
_ = reg312554
reg312555 := __e.Call(__defun__shen_4incinfs)
_ = reg312555
reg312556 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312557 := MakeSymbol("vector")
reg312558 := Nil;
reg312559 := PrimCons(A, reg312558)
reg312560 := PrimCons(reg312557, reg312559)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312560, V3695, V3696, V3697)
return
}, 0)
reg312562 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312556)
Result := reg312562
_ = Result
reg312563 := __e.Call(__defun__shen_4unbindv, V3565, V3696)
_ = reg312563
reg312565 = Result
} else {
reg312564 := False;
reg312565 = reg312564
}
reg312566 = reg312565
}
Result := reg312566
_ = Result
reg312567 := __e.Call(__defun__shen_4unbindv, V3562, V3696)
_ = reg312567
reg312569 = Result
} else {
reg312568 := False;
reg312569 = reg312568
}
reg312570 = reg312569
}
reg312589 = reg312570
} else {
reg312571 := __e.Call(__defun__shen_4pvar_2, V3561)
var reg312588 Obj
if reg312571 == True {
reg312572 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312572
_ = A
reg312573 := MakeSymbol("vector")
reg312574 := Nil;
reg312575 := PrimCons(A, reg312574)
reg312576 := PrimCons(reg312573, reg312575)
reg312577 := __e.Call(__defun__shen_4bindv, V3561, reg312576, V3696)
_ = reg312577
reg312578 := __e.Call(__defun__shen_4incinfs)
_ = reg312578
reg312579 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312580 := MakeSymbol("vector")
reg312581 := Nil;
reg312582 := PrimCons(A, reg312581)
reg312583 := PrimCons(reg312580, reg312582)
__ctx.TailApply(__defun__shen_4th_d, Y, reg312583, V3695, V3696, V3697)
return
}, 0)
reg312585 := __e.Call(__defun__shen_4th_d, X, A, V3695, V3696, reg312579)
Result := reg312585
_ = Result
reg312586 := __e.Call(__defun__shen_4unbindv, V3561, V3696)
_ = reg312586
reg312588 = Result
} else {
reg312587 := False;
reg312588 = reg312587
}
reg312589 = reg312588
}
reg312591 = reg312589
} else {
reg312590 := False;
reg312591 = reg312590
}
reg312593 = reg312591
} else {
reg312592 := False;
reg312593 = reg312592
}
reg312595 = reg312593
} else {
reg312594 := False;
reg312595 = reg312594
}
reg312597 = reg312595
} else {
reg312596 := False;
reg312597 = reg312596
}
reg312599 = reg312597
} else {
reg312598 := False;
reg312599 = reg312598
}
Case := reg312599
_ = Case
reg312600 := False;
reg312601 := PrimEqual(Case, reg312600)
var reg313473 Obj
if reg312601 == True {
reg312602 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3567 := reg312602
_ = V3567
reg312603 := PrimIsPair(V3567)
var reg312651 Obj
if reg312603 == True {
reg312604 := PrimHead(V3567)
reg312605 := __e.Call(__defun__shen_4lazyderef, reg312604, V3696)
V3568 := reg312605
_ = V3568
reg312606 := MakeSymbol("@s")
reg312607 := PrimEqual(reg312606, V3568)
var reg312649 Obj
if reg312607 == True {
reg312608 := PrimTail(V3567)
reg312609 := __e.Call(__defun__shen_4lazyderef, reg312608, V3696)
V3569 := reg312609
_ = V3569
reg312610 := PrimIsPair(V3569)
var reg312647 Obj
if reg312610 == True {
reg312611 := PrimHead(V3569)
X := reg312611
_ = X
reg312612 := PrimTail(V3569)
reg312613 := __e.Call(__defun__shen_4lazyderef, reg312612, V3696)
V3570 := reg312613
_ = V3570
reg312614 := PrimIsPair(V3570)
var reg312645 Obj
if reg312614 == True {
reg312615 := PrimHead(V3570)
Y := reg312615
_ = Y
reg312616 := PrimTail(V3570)
reg312617 := __e.Call(__defun__shen_4lazyderef, reg312616, V3696)
V3571 := reg312617
_ = V3571
reg312618 := Nil;
reg312619 := PrimEqual(reg312618, V3571)
var reg312643 Obj
if reg312619 == True {
reg312620 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3572 := reg312620
_ = V3572
reg312621 := MakeSymbol("string")
reg312622 := PrimEqual(reg312621, V3572)
var reg312641 Obj
if reg312622 == True {
reg312623 := __e.Call(__defun__shen_4incinfs)
_ = reg312623
reg312624 := MakeSymbol("string")
reg312625 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312626 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, Y, reg312626, V3695, V3696, V3697)
return
}, 0)
reg312628 := __e.Call(__defun__shen_4th_d, X, reg312624, V3695, V3696, reg312625)
reg312641 = reg312628
} else {
reg312629 := __e.Call(__defun__shen_4pvar_2, V3572)
var reg312640 Obj
if reg312629 == True {
reg312630 := MakeSymbol("string")
reg312631 := __e.Call(__defun__shen_4bindv, V3572, reg312630, V3696)
_ = reg312631
reg312632 := __e.Call(__defun__shen_4incinfs)
_ = reg312632
reg312633 := MakeSymbol("string")
reg312634 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312635 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, Y, reg312635, V3695, V3696, V3697)
return
}, 0)
reg312637 := __e.Call(__defun__shen_4th_d, X, reg312633, V3695, V3696, reg312634)
Result := reg312637
_ = Result
reg312638 := __e.Call(__defun__shen_4unbindv, V3572, V3696)
_ = reg312638
reg312640 = Result
} else {
reg312639 := False;
reg312640 = reg312639
}
reg312641 = reg312640
}
reg312643 = reg312641
} else {
reg312642 := False;
reg312643 = reg312642
}
reg312645 = reg312643
} else {
reg312644 := False;
reg312645 = reg312644
}
reg312647 = reg312645
} else {
reg312646 := False;
reg312647 = reg312646
}
reg312649 = reg312647
} else {
reg312648 := False;
reg312649 = reg312648
}
reg312651 = reg312649
} else {
reg312650 := False;
reg312651 = reg312650
}
Case := reg312651
_ = Case
reg312652 := False;
reg312653 := PrimEqual(Case, reg312652)
var reg313472 Obj
if reg312653 == True {
reg312654 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3573 := reg312654
_ = V3573
reg312655 := PrimIsPair(V3573)
var reg312938 Obj
if reg312655 == True {
reg312656 := PrimHead(V3573)
reg312657 := __e.Call(__defun__shen_4lazyderef, reg312656, V3696)
V3574 := reg312657
_ = V3574
reg312658 := MakeSymbol("lambda")
reg312659 := PrimEqual(reg312658, V3574)
var reg312936 Obj
if reg312659 == True {
reg312660 := PrimTail(V3573)
reg312661 := __e.Call(__defun__shen_4lazyderef, reg312660, V3696)
V3575 := reg312661
_ = V3575
reg312662 := PrimIsPair(V3575)
var reg312934 Obj
if reg312662 == True {
reg312663 := PrimHead(V3575)
X := reg312663
_ = X
reg312664 := PrimTail(V3575)
reg312665 := __e.Call(__defun__shen_4lazyderef, reg312664, V3696)
V3576 := reg312665
_ = V3576
reg312666 := PrimIsPair(V3576)
var reg312932 Obj
if reg312666 == True {
reg312667 := PrimHead(V3576)
Y := reg312667
_ = Y
reg312668 := PrimTail(V3576)
reg312669 := __e.Call(__defun__shen_4lazyderef, reg312668, V3696)
V3577 := reg312669
_ = V3577
reg312670 := Nil;
reg312671 := PrimEqual(reg312670, V3577)
var reg312930 Obj
if reg312671 == True {
reg312672 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3578 := reg312672
_ = V3578
reg312673 := PrimIsPair(V3578)
var reg312928 Obj
if reg312673 == True {
reg312674 := PrimHead(V3578)
A := reg312674
_ = A
reg312675 := PrimTail(V3578)
reg312676 := __e.Call(__defun__shen_4lazyderef, reg312675, V3696)
V3579 := reg312676
_ = V3579
reg312677 := PrimIsPair(V3579)
var reg312894 Obj
if reg312677 == True {
reg312678 := PrimHead(V3579)
reg312679 := __e.Call(__defun__shen_4lazyderef, reg312678, V3696)
V3580 := reg312679
_ = V3580
reg312680 := MakeSymbol("-->")
reg312681 := PrimEqual(reg312680, V3580)
var reg312862 Obj
if reg312681 == True {
reg312682 := PrimTail(V3579)
reg312683 := __e.Call(__defun__shen_4lazyderef, reg312682, V3696)
V3581 := reg312683
_ = V3581
reg312684 := PrimIsPair(V3581)
var reg312768 Obj
if reg312684 == True {
reg312685 := PrimHead(V3581)
B := reg312685
_ = B
reg312686 := PrimTail(V3581)
reg312687 := __e.Call(__defun__shen_4lazyderef, reg312686, V3696)
V3582 := reg312687
_ = V3582
reg312688 := Nil;
reg312689 := PrimEqual(reg312688, V3582)
var reg312738 Obj
if reg312689 == True {
reg312690 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312690
_ = Z
reg312691 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312691
_ = X_e_e
reg312692 := __e.Call(__defun__shen_4incinfs)
_ = reg312692
reg312693 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312694 := __e.Call(__defun__shen_4placeholder)
reg312695 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312696 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312697 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312698 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312699 := __e.Call(__defun__shen_4ebr, reg312696, reg312697, reg312698)
reg312700 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312701 := MakeSymbol(":")
reg312702 := Nil;
reg312703 := PrimCons(A, reg312702)
reg312704 := PrimCons(reg312701, reg312703)
reg312705 := PrimCons(X_e_e, reg312704)
reg312706 := PrimCons(reg312705, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312706, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312699, V3696, reg312700)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312694, V3696, reg312695)
return
}, 0)
reg312710 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312693)
reg312738 = reg312710
} else {
reg312711 := __e.Call(__defun__shen_4pvar_2, V3582)
var reg312737 Obj
if reg312711 == True {
reg312712 := Nil;
reg312713 := __e.Call(__defun__shen_4bindv, V3582, reg312712, V3696)
_ = reg312713
reg312714 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312714
_ = Z
reg312715 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312715
_ = X_e_e
reg312716 := __e.Call(__defun__shen_4incinfs)
_ = reg312716
reg312717 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312718 := __e.Call(__defun__shen_4placeholder)
reg312719 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312720 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312721 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312722 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312723 := __e.Call(__defun__shen_4ebr, reg312720, reg312721, reg312722)
reg312724 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312725 := MakeSymbol(":")
reg312726 := Nil;
reg312727 := PrimCons(A, reg312726)
reg312728 := PrimCons(reg312725, reg312727)
reg312729 := PrimCons(X_e_e, reg312728)
reg312730 := PrimCons(reg312729, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312730, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312723, V3696, reg312724)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312718, V3696, reg312719)
return
}, 0)
reg312734 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312717)
Result := reg312734
_ = Result
reg312735 := __e.Call(__defun__shen_4unbindv, V3582, V3696)
_ = reg312735
reg312737 = Result
} else {
reg312736 := False;
reg312737 = reg312736
}
reg312738 = reg312737
}
reg312768 = reg312738
} else {
reg312739 := __e.Call(__defun__shen_4pvar_2, V3581)
var reg312767 Obj
if reg312739 == True {
reg312740 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312740
_ = B
reg312741 := Nil;
reg312742 := PrimCons(B, reg312741)
reg312743 := __e.Call(__defun__shen_4bindv, V3581, reg312742, V3696)
_ = reg312743
reg312744 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312744
_ = Z
reg312745 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312745
_ = X_e_e
reg312746 := __e.Call(__defun__shen_4incinfs)
_ = reg312746
reg312747 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312748 := __e.Call(__defun__shen_4placeholder)
reg312749 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312750 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312751 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312752 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312753 := __e.Call(__defun__shen_4ebr, reg312750, reg312751, reg312752)
reg312754 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312755 := MakeSymbol(":")
reg312756 := Nil;
reg312757 := PrimCons(A, reg312756)
reg312758 := PrimCons(reg312755, reg312757)
reg312759 := PrimCons(X_e_e, reg312758)
reg312760 := PrimCons(reg312759, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312760, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312753, V3696, reg312754)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312748, V3696, reg312749)
return
}, 0)
reg312764 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312747)
Result := reg312764
_ = Result
reg312765 := __e.Call(__defun__shen_4unbindv, V3581, V3696)
_ = reg312765
reg312767 = Result
} else {
reg312766 := False;
reg312767 = reg312766
}
reg312768 = reg312767
}
reg312862 = reg312768
} else {
reg312769 := __e.Call(__defun__shen_4pvar_2, V3580)
var reg312861 Obj
if reg312769 == True {
reg312770 := MakeSymbol("-->")
reg312771 := __e.Call(__defun__shen_4bindv, V3580, reg312770, V3696)
_ = reg312771
reg312772 := PrimTail(V3579)
reg312773 := __e.Call(__defun__shen_4lazyderef, reg312772, V3696)
V3583 := reg312773
_ = V3583
reg312774 := PrimIsPair(V3583)
var reg312858 Obj
if reg312774 == True {
reg312775 := PrimHead(V3583)
B := reg312775
_ = B
reg312776 := PrimTail(V3583)
reg312777 := __e.Call(__defun__shen_4lazyderef, reg312776, V3696)
V3584 := reg312777
_ = V3584
reg312778 := Nil;
reg312779 := PrimEqual(reg312778, V3584)
var reg312828 Obj
if reg312779 == True {
reg312780 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312780
_ = Z
reg312781 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312781
_ = X_e_e
reg312782 := __e.Call(__defun__shen_4incinfs)
_ = reg312782
reg312783 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312784 := __e.Call(__defun__shen_4placeholder)
reg312785 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312786 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312787 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312788 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312789 := __e.Call(__defun__shen_4ebr, reg312786, reg312787, reg312788)
reg312790 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312791 := MakeSymbol(":")
reg312792 := Nil;
reg312793 := PrimCons(A, reg312792)
reg312794 := PrimCons(reg312791, reg312793)
reg312795 := PrimCons(X_e_e, reg312794)
reg312796 := PrimCons(reg312795, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312796, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312789, V3696, reg312790)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312784, V3696, reg312785)
return
}, 0)
reg312800 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312783)
reg312828 = reg312800
} else {
reg312801 := __e.Call(__defun__shen_4pvar_2, V3584)
var reg312827 Obj
if reg312801 == True {
reg312802 := Nil;
reg312803 := __e.Call(__defun__shen_4bindv, V3584, reg312802, V3696)
_ = reg312803
reg312804 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312804
_ = Z
reg312805 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312805
_ = X_e_e
reg312806 := __e.Call(__defun__shen_4incinfs)
_ = reg312806
reg312807 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312808 := __e.Call(__defun__shen_4placeholder)
reg312809 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312810 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312811 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312812 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312813 := __e.Call(__defun__shen_4ebr, reg312810, reg312811, reg312812)
reg312814 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312815 := MakeSymbol(":")
reg312816 := Nil;
reg312817 := PrimCons(A, reg312816)
reg312818 := PrimCons(reg312815, reg312817)
reg312819 := PrimCons(X_e_e, reg312818)
reg312820 := PrimCons(reg312819, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312820, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312813, V3696, reg312814)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312808, V3696, reg312809)
return
}, 0)
reg312824 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312807)
Result := reg312824
_ = Result
reg312825 := __e.Call(__defun__shen_4unbindv, V3584, V3696)
_ = reg312825
reg312827 = Result
} else {
reg312826 := False;
reg312827 = reg312826
}
reg312828 = reg312827
}
reg312858 = reg312828
} else {
reg312829 := __e.Call(__defun__shen_4pvar_2, V3583)
var reg312857 Obj
if reg312829 == True {
reg312830 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312830
_ = B
reg312831 := Nil;
reg312832 := PrimCons(B, reg312831)
reg312833 := __e.Call(__defun__shen_4bindv, V3583, reg312832, V3696)
_ = reg312833
reg312834 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312834
_ = Z
reg312835 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312835
_ = X_e_e
reg312836 := __e.Call(__defun__shen_4incinfs)
_ = reg312836
reg312837 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312838 := __e.Call(__defun__shen_4placeholder)
reg312839 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312840 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312841 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312842 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312843 := __e.Call(__defun__shen_4ebr, reg312840, reg312841, reg312842)
reg312844 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312845 := MakeSymbol(":")
reg312846 := Nil;
reg312847 := PrimCons(A, reg312846)
reg312848 := PrimCons(reg312845, reg312847)
reg312849 := PrimCons(X_e_e, reg312848)
reg312850 := PrimCons(reg312849, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312850, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312843, V3696, reg312844)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312838, V3696, reg312839)
return
}, 0)
reg312854 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312837)
Result := reg312854
_ = Result
reg312855 := __e.Call(__defun__shen_4unbindv, V3583, V3696)
_ = reg312855
reg312857 = Result
} else {
reg312856 := False;
reg312857 = reg312856
}
reg312858 = reg312857
}
Result := reg312858
_ = Result
reg312859 := __e.Call(__defun__shen_4unbindv, V3580, V3696)
_ = reg312859
reg312861 = Result
} else {
reg312860 := False;
reg312861 = reg312860
}
reg312862 = reg312861
}
reg312894 = reg312862
} else {
reg312863 := __e.Call(__defun__shen_4pvar_2, V3579)
var reg312893 Obj
if reg312863 == True {
reg312864 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312864
_ = B
reg312865 := MakeSymbol("-->")
reg312866 := Nil;
reg312867 := PrimCons(B, reg312866)
reg312868 := PrimCons(reg312865, reg312867)
reg312869 := __e.Call(__defun__shen_4bindv, V3579, reg312868, V3696)
_ = reg312869
reg312870 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312870
_ = Z
reg312871 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312871
_ = X_e_e
reg312872 := __e.Call(__defun__shen_4incinfs)
_ = reg312872
reg312873 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312874 := __e.Call(__defun__shen_4placeholder)
reg312875 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312876 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312877 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312878 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312879 := __e.Call(__defun__shen_4ebr, reg312876, reg312877, reg312878)
reg312880 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312881 := MakeSymbol(":")
reg312882 := Nil;
reg312883 := PrimCons(A, reg312882)
reg312884 := PrimCons(reg312881, reg312883)
reg312885 := PrimCons(X_e_e, reg312884)
reg312886 := PrimCons(reg312885, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312886, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312879, V3696, reg312880)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312874, V3696, reg312875)
return
}, 0)
reg312890 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312873)
Result := reg312890
_ = Result
reg312891 := __e.Call(__defun__shen_4unbindv, V3579, V3696)
_ = reg312891
reg312893 = Result
} else {
reg312892 := False;
reg312893 = reg312892
}
reg312894 = reg312893
}
reg312928 = reg312894
} else {
reg312895 := __e.Call(__defun__shen_4pvar_2, V3578)
var reg312927 Obj
if reg312895 == True {
reg312896 := __e.Call(__defun__shen_4newpv, V3696)
A := reg312896
_ = A
reg312897 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312897
_ = B
reg312898 := MakeSymbol("-->")
reg312899 := Nil;
reg312900 := PrimCons(B, reg312899)
reg312901 := PrimCons(reg312898, reg312900)
reg312902 := PrimCons(A, reg312901)
reg312903 := __e.Call(__defun__shen_4bindv, V3578, reg312902, V3696)
_ = reg312903
reg312904 := __e.Call(__defun__shen_4newpv, V3696)
Z := reg312904
_ = Z
reg312905 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312905
_ = X_e_e
reg312906 := __e.Call(__defun__shen_4incinfs)
_ = reg312906
reg312907 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312908 := __e.Call(__defun__shen_4placeholder)
reg312909 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312910 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312911 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312912 := __e.Call(__defun__shen_4lazyderef, Y, V3696)
reg312913 := __e.Call(__defun__shen_4ebr, reg312910, reg312911, reg312912)
reg312914 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312915 := MakeSymbol(":")
reg312916 := Nil;
reg312917 := PrimCons(A, reg312916)
reg312918 := PrimCons(reg312915, reg312917)
reg312919 := PrimCons(X_e_e, reg312918)
reg312920 := PrimCons(reg312919, V3695)
__ctx.TailApply(__defun__shen_4th_d, Z, B, reg312920, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, Z, reg312913, V3696, reg312914)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312908, V3696, reg312909)
return
}, 0)
reg312924 := __e.Call(__defun__cut, Throwcontrol, V3696, reg312907)
Result := reg312924
_ = Result
reg312925 := __e.Call(__defun__shen_4unbindv, V3578, V3696)
_ = reg312925
reg312927 = Result
} else {
reg312926 := False;
reg312927 = reg312926
}
reg312928 = reg312927
}
reg312930 = reg312928
} else {
reg312929 := False;
reg312930 = reg312929
}
reg312932 = reg312930
} else {
reg312931 := False;
reg312932 = reg312931
}
reg312934 = reg312932
} else {
reg312933 := False;
reg312934 = reg312933
}
reg312936 = reg312934
} else {
reg312935 := False;
reg312936 = reg312935
}
reg312938 = reg312936
} else {
reg312937 := False;
reg312938 = reg312937
}
Case := reg312938
_ = Case
reg312939 := False;
reg312940 := PrimEqual(Case, reg312939)
var reg313471 Obj
if reg312940 == True {
reg312941 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3585 := reg312941
_ = V3585
reg312942 := PrimIsPair(V3585)
var reg312996 Obj
if reg312942 == True {
reg312943 := PrimHead(V3585)
reg312944 := __e.Call(__defun__shen_4lazyderef, reg312943, V3696)
V3586 := reg312944
_ = V3586
reg312945 := MakeSymbol("let")
reg312946 := PrimEqual(reg312945, V3586)
var reg312994 Obj
if reg312946 == True {
reg312947 := PrimTail(V3585)
reg312948 := __e.Call(__defun__shen_4lazyderef, reg312947, V3696)
V3587 := reg312948
_ = V3587
reg312949 := PrimIsPair(V3587)
var reg312992 Obj
if reg312949 == True {
reg312950 := PrimHead(V3587)
X := reg312950
_ = X
reg312951 := PrimTail(V3587)
reg312952 := __e.Call(__defun__shen_4lazyderef, reg312951, V3696)
V3588 := reg312952
_ = V3588
reg312953 := PrimIsPair(V3588)
var reg312990 Obj
if reg312953 == True {
reg312954 := PrimHead(V3588)
Y := reg312954
_ = Y
reg312955 := PrimTail(V3588)
reg312956 := __e.Call(__defun__shen_4lazyderef, reg312955, V3696)
V3589 := reg312956
_ = V3589
reg312957 := PrimIsPair(V3589)
var reg312988 Obj
if reg312957 == True {
reg312958 := PrimHead(V3589)
Z := reg312958
_ = Z
reg312959 := PrimTail(V3589)
reg312960 := __e.Call(__defun__shen_4lazyderef, reg312959, V3696)
V3590 := reg312960
_ = V3590
reg312961 := Nil;
reg312962 := PrimEqual(reg312961, V3590)
var reg312986 Obj
if reg312962 == True {
reg312963 := __e.Call(__defun__shen_4newpv, V3696)
W := reg312963
_ = W
reg312964 := __e.Call(__defun__shen_4newpv, V3696)
X_e_e := reg312964
_ = X_e_e
reg312965 := __e.Call(__defun__shen_4newpv, V3696)
B := reg312965
_ = B
reg312966 := __e.Call(__defun__shen_4incinfs)
_ = reg312966
reg312967 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312968 := __e.Call(__defun__shen_4placeholder)
reg312969 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312970 := __e.Call(__defun__shen_4lazyderef, X_e_e, V3696)
reg312971 := __e.Call(__defun__shen_4lazyderef, X, V3696)
reg312972 := __e.Call(__defun__shen_4lazyderef, Z, V3696)
reg312973 := __e.Call(__defun__shen_4ebr, reg312970, reg312971, reg312972)
reg312974 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg312975 := MakeSymbol(":")
reg312976 := Nil;
reg312977 := PrimCons(B, reg312976)
reg312978 := PrimCons(reg312975, reg312977)
reg312979 := PrimCons(X_e_e, reg312978)
reg312980 := PrimCons(reg312979, V3695)
__ctx.TailApply(__defun__shen_4th_d, W, V3694, reg312980, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__bind, W, reg312973, V3696, reg312974)
return
}, 0)
__ctx.TailApply(__defun__bind, X_e_e, reg312968, V3696, reg312969)
return
}, 0)
reg312984 := __e.Call(__defun__shen_4th_d, Y, B, V3695, V3696, reg312967)
reg312986 = reg312984
} else {
reg312985 := False;
reg312986 = reg312985
}
reg312988 = reg312986
} else {
reg312987 := False;
reg312988 = reg312987
}
reg312990 = reg312988
} else {
reg312989 := False;
reg312990 = reg312989
}
reg312992 = reg312990
} else {
reg312991 := False;
reg312992 = reg312991
}
reg312994 = reg312992
} else {
reg312993 := False;
reg312994 = reg312993
}
reg312996 = reg312994
} else {
reg312995 := False;
reg312996 = reg312995
}
Case := reg312996
_ = Case
reg312997 := False;
reg312998 := PrimEqual(Case, reg312997)
var reg313470 Obj
if reg312998 == True {
reg312999 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3591 := reg312999
_ = V3591
reg313000 := PrimIsPair(V3591)
var reg313210 Obj
if reg313000 == True {
reg313001 := PrimHead(V3591)
reg313002 := __e.Call(__defun__shen_4lazyderef, reg313001, V3696)
V3592 := reg313002
_ = V3592
reg313003 := MakeSymbol("open")
reg313004 := PrimEqual(reg313003, V3592)
var reg313208 Obj
if reg313004 == True {
reg313005 := PrimTail(V3591)
reg313006 := __e.Call(__defun__shen_4lazyderef, reg313005, V3696)
V3593 := reg313006
_ = V3593
reg313007 := PrimIsPair(V3593)
var reg313206 Obj
if reg313007 == True {
reg313008 := PrimHead(V3593)
FileName := reg313008
_ = FileName
reg313009 := PrimTail(V3593)
reg313010 := __e.Call(__defun__shen_4lazyderef, reg313009, V3696)
V3594 := reg313010
_ = V3594
reg313011 := PrimIsPair(V3594)
var reg313204 Obj
if reg313011 == True {
reg313012 := PrimHead(V3594)
Direction3524 := reg313012
_ = Direction3524
reg313013 := PrimTail(V3594)
reg313014 := __e.Call(__defun__shen_4lazyderef, reg313013, V3696)
V3595 := reg313014
_ = V3595
reg313015 := Nil;
reg313016 := PrimEqual(reg313015, V3595)
var reg313202 Obj
if reg313016 == True {
reg313017 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3596 := reg313017
_ = V3596
reg313018 := PrimIsPair(V3596)
var reg313200 Obj
if reg313018 == True {
reg313019 := PrimHead(V3596)
reg313020 := __e.Call(__defun__shen_4lazyderef, reg313019, V3696)
V3597 := reg313020
_ = V3597
reg313021 := MakeSymbol("stream")
reg313022 := PrimEqual(reg313021, V3597)
var reg313173 Obj
if reg313022 == True {
reg313023 := PrimTail(V3596)
reg313024 := __e.Call(__defun__shen_4lazyderef, reg313023, V3696)
V3598 := reg313024
_ = V3598
reg313025 := PrimIsPair(V3598)
var reg313094 Obj
if reg313025 == True {
reg313026 := PrimHead(V3598)
Direction := reg313026
_ = Direction
reg313027 := PrimTail(V3598)
reg313028 := __e.Call(__defun__shen_4lazyderef, reg313027, V3696)
V3599 := reg313028
_ = V3599
reg313029 := Nil;
reg313030 := PrimEqual(reg313029, V3599)
var reg313069 Obj
if reg313030 == True {
reg313031 := __e.Call(__defun__shen_4incinfs)
_ = reg313031
reg313032 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313033 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313034 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313035 := MakeSymbol("in")
reg313036 := MakeSymbol("out")
reg313037 := Nil;
reg313038 := PrimCons(reg313036, reg313037)
reg313039 := PrimCons(reg313035, reg313038)
reg313040 := __e.Call(__defun__element_2, reg313034, reg313039)
reg313041 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313042 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313042, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313040, V3696, reg313041)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313033)
return
}, 0)
reg313046 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313032)
reg313069 = reg313046
} else {
reg313047 := __e.Call(__defun__shen_4pvar_2, V3599)
var reg313068 Obj
if reg313047 == True {
reg313048 := Nil;
reg313049 := __e.Call(__defun__shen_4bindv, V3599, reg313048, V3696)
_ = reg313049
reg313050 := __e.Call(__defun__shen_4incinfs)
_ = reg313050
reg313051 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313052 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313053 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313054 := MakeSymbol("in")
reg313055 := MakeSymbol("out")
reg313056 := Nil;
reg313057 := PrimCons(reg313055, reg313056)
reg313058 := PrimCons(reg313054, reg313057)
reg313059 := __e.Call(__defun__element_2, reg313053, reg313058)
reg313060 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313061 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313061, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313059, V3696, reg313060)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313052)
return
}, 0)
reg313065 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313051)
Result := reg313065
_ = Result
reg313066 := __e.Call(__defun__shen_4unbindv, V3599, V3696)
_ = reg313066
reg313068 = Result
} else {
reg313067 := False;
reg313068 = reg313067
}
reg313069 = reg313068
}
reg313094 = reg313069
} else {
reg313070 := __e.Call(__defun__shen_4pvar_2, V3598)
var reg313093 Obj
if reg313070 == True {
reg313071 := __e.Call(__defun__shen_4newpv, V3696)
Direction := reg313071
_ = Direction
reg313072 := Nil;
reg313073 := PrimCons(Direction, reg313072)
reg313074 := __e.Call(__defun__shen_4bindv, V3598, reg313073, V3696)
_ = reg313074
reg313075 := __e.Call(__defun__shen_4incinfs)
_ = reg313075
reg313076 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313077 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313078 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313079 := MakeSymbol("in")
reg313080 := MakeSymbol("out")
reg313081 := Nil;
reg313082 := PrimCons(reg313080, reg313081)
reg313083 := PrimCons(reg313079, reg313082)
reg313084 := __e.Call(__defun__element_2, reg313078, reg313083)
reg313085 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313086 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313086, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313084, V3696, reg313085)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313077)
return
}, 0)
reg313090 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313076)
Result := reg313090
_ = Result
reg313091 := __e.Call(__defun__shen_4unbindv, V3598, V3696)
_ = reg313091
reg313093 = Result
} else {
reg313092 := False;
reg313093 = reg313092
}
reg313094 = reg313093
}
reg313173 = reg313094
} else {
reg313095 := __e.Call(__defun__shen_4pvar_2, V3597)
var reg313172 Obj
if reg313095 == True {
reg313096 := MakeSymbol("stream")
reg313097 := __e.Call(__defun__shen_4bindv, V3597, reg313096, V3696)
_ = reg313097
reg313098 := PrimTail(V3596)
reg313099 := __e.Call(__defun__shen_4lazyderef, reg313098, V3696)
V3600 := reg313099
_ = V3600
reg313100 := PrimIsPair(V3600)
var reg313169 Obj
if reg313100 == True {
reg313101 := PrimHead(V3600)
Direction := reg313101
_ = Direction
reg313102 := PrimTail(V3600)
reg313103 := __e.Call(__defun__shen_4lazyderef, reg313102, V3696)
V3601 := reg313103
_ = V3601
reg313104 := Nil;
reg313105 := PrimEqual(reg313104, V3601)
var reg313144 Obj
if reg313105 == True {
reg313106 := __e.Call(__defun__shen_4incinfs)
_ = reg313106
reg313107 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313108 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313109 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313110 := MakeSymbol("in")
reg313111 := MakeSymbol("out")
reg313112 := Nil;
reg313113 := PrimCons(reg313111, reg313112)
reg313114 := PrimCons(reg313110, reg313113)
reg313115 := __e.Call(__defun__element_2, reg313109, reg313114)
reg313116 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313117 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313117, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313115, V3696, reg313116)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313108)
return
}, 0)
reg313121 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313107)
reg313144 = reg313121
} else {
reg313122 := __e.Call(__defun__shen_4pvar_2, V3601)
var reg313143 Obj
if reg313122 == True {
reg313123 := Nil;
reg313124 := __e.Call(__defun__shen_4bindv, V3601, reg313123, V3696)
_ = reg313124
reg313125 := __e.Call(__defun__shen_4incinfs)
_ = reg313125
reg313126 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313127 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313128 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313129 := MakeSymbol("in")
reg313130 := MakeSymbol("out")
reg313131 := Nil;
reg313132 := PrimCons(reg313130, reg313131)
reg313133 := PrimCons(reg313129, reg313132)
reg313134 := __e.Call(__defun__element_2, reg313128, reg313133)
reg313135 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313136 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313136, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313134, V3696, reg313135)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313127)
return
}, 0)
reg313140 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313126)
Result := reg313140
_ = Result
reg313141 := __e.Call(__defun__shen_4unbindv, V3601, V3696)
_ = reg313141
reg313143 = Result
} else {
reg313142 := False;
reg313143 = reg313142
}
reg313144 = reg313143
}
reg313169 = reg313144
} else {
reg313145 := __e.Call(__defun__shen_4pvar_2, V3600)
var reg313168 Obj
if reg313145 == True {
reg313146 := __e.Call(__defun__shen_4newpv, V3696)
Direction := reg313146
_ = Direction
reg313147 := Nil;
reg313148 := PrimCons(Direction, reg313147)
reg313149 := __e.Call(__defun__shen_4bindv, V3600, reg313148, V3696)
_ = reg313149
reg313150 := __e.Call(__defun__shen_4incinfs)
_ = reg313150
reg313151 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313152 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313153 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313154 := MakeSymbol("in")
reg313155 := MakeSymbol("out")
reg313156 := Nil;
reg313157 := PrimCons(reg313155, reg313156)
reg313158 := PrimCons(reg313154, reg313157)
reg313159 := __e.Call(__defun__element_2, reg313153, reg313158)
reg313160 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313161 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313161, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313159, V3696, reg313160)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313152)
return
}, 0)
reg313165 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313151)
Result := reg313165
_ = Result
reg313166 := __e.Call(__defun__shen_4unbindv, V3600, V3696)
_ = reg313166
reg313168 = Result
} else {
reg313167 := False;
reg313168 = reg313167
}
reg313169 = reg313168
}
Result := reg313169
_ = Result
reg313170 := __e.Call(__defun__shen_4unbindv, V3597, V3696)
_ = reg313170
reg313172 = Result
} else {
reg313171 := False;
reg313172 = reg313171
}
reg313173 = reg313172
}
reg313200 = reg313173
} else {
reg313174 := __e.Call(__defun__shen_4pvar_2, V3596)
var reg313199 Obj
if reg313174 == True {
reg313175 := __e.Call(__defun__shen_4newpv, V3696)
Direction := reg313175
_ = Direction
reg313176 := MakeSymbol("stream")
reg313177 := Nil;
reg313178 := PrimCons(Direction, reg313177)
reg313179 := PrimCons(reg313176, reg313178)
reg313180 := __e.Call(__defun__shen_4bindv, V3596, reg313179, V3696)
_ = reg313180
reg313181 := __e.Call(__defun__shen_4incinfs)
_ = reg313181
reg313182 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313183 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313184 := __e.Call(__defun__shen_4lazyderef, Direction, V3696)
reg313185 := MakeSymbol("in")
reg313186 := MakeSymbol("out")
reg313187 := Nil;
reg313188 := PrimCons(reg313186, reg313187)
reg313189 := PrimCons(reg313185, reg313188)
reg313190 := __e.Call(__defun__element_2, reg313184, reg313189)
reg313191 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313192 := MakeSymbol("string")
__ctx.TailApply(__defun__shen_4th_d, FileName, reg313192, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__fwhen, reg313190, V3696, reg313191)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313183)
return
}, 0)
reg313196 := __e.Call(__defun__unify_b, Direction, Direction3524, V3696, reg313182)
Result := reg313196
_ = Result
reg313197 := __e.Call(__defun__shen_4unbindv, V3596, V3696)
_ = reg313197
reg313199 = Result
} else {
reg313198 := False;
reg313199 = reg313198
}
reg313200 = reg313199
}
reg313202 = reg313200
} else {
reg313201 := False;
reg313202 = reg313201
}
reg313204 = reg313202
} else {
reg313203 := False;
reg313204 = reg313203
}
reg313206 = reg313204
} else {
reg313205 := False;
reg313206 = reg313205
}
reg313208 = reg313206
} else {
reg313207 := False;
reg313208 = reg313207
}
reg313210 = reg313208
} else {
reg313209 := False;
reg313210 = reg313209
}
Case := reg313210
_ = Case
reg313211 := False;
reg313212 := PrimEqual(Case, reg313211)
var reg313469 Obj
if reg313212 == True {
reg313213 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3602 := reg313213
_ = V3602
reg313214 := PrimIsPair(V3602)
var reg313246 Obj
if reg313214 == True {
reg313215 := PrimHead(V3602)
reg313216 := __e.Call(__defun__shen_4lazyderef, reg313215, V3696)
V3603 := reg313216
_ = V3603
reg313217 := MakeSymbol("type")
reg313218 := PrimEqual(reg313217, V3603)
var reg313244 Obj
if reg313218 == True {
reg313219 := PrimTail(V3602)
reg313220 := __e.Call(__defun__shen_4lazyderef, reg313219, V3696)
V3604 := reg313220
_ = V3604
reg313221 := PrimIsPair(V3604)
var reg313242 Obj
if reg313221 == True {
reg313222 := PrimHead(V3604)
X := reg313222
_ = X
reg313223 := PrimTail(V3604)
reg313224 := __e.Call(__defun__shen_4lazyderef, reg313223, V3696)
V3605 := reg313224
_ = V3605
reg313225 := PrimIsPair(V3605)
var reg313240 Obj
if reg313225 == True {
reg313226 := PrimHead(V3605)
A := reg313226
_ = A
reg313227 := PrimTail(V3605)
reg313228 := __e.Call(__defun__shen_4lazyderef, reg313227, V3696)
V3606 := reg313228
_ = V3606
reg313229 := Nil;
reg313230 := PrimEqual(reg313229, V3606)
var reg313238 Obj
if reg313230 == True {
reg313231 := __e.Call(__defun__shen_4incinfs)
_ = reg313231
reg313232 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313233 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, X, A, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__unify, A, V3694, V3696, reg313233)
return
}, 0)
reg313236 := __e.Call(__defun__cut, Throwcontrol, V3696, reg313232)
reg313238 = reg313236
} else {
reg313237 := False;
reg313238 = reg313237
}
reg313240 = reg313238
} else {
reg313239 := False;
reg313240 = reg313239
}
reg313242 = reg313240
} else {
reg313241 := False;
reg313242 = reg313241
}
reg313244 = reg313242
} else {
reg313243 := False;
reg313244 = reg313243
}
reg313246 = reg313244
} else {
reg313245 := False;
reg313246 = reg313245
}
Case := reg313246
_ = Case
reg313247 := False;
reg313248 := PrimEqual(Case, reg313247)
var reg313468 Obj
if reg313248 == True {
reg313249 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3607 := reg313249
_ = V3607
reg313250 := PrimIsPair(V3607)
var reg313290 Obj
if reg313250 == True {
reg313251 := PrimHead(V3607)
reg313252 := __e.Call(__defun__shen_4lazyderef, reg313251, V3696)
V3608 := reg313252
_ = V3608
reg313253 := MakeSymbol("input+")
reg313254 := PrimEqual(reg313253, V3608)
var reg313288 Obj
if reg313254 == True {
reg313255 := PrimTail(V3607)
reg313256 := __e.Call(__defun__shen_4lazyderef, reg313255, V3696)
V3609 := reg313256
_ = V3609
reg313257 := PrimIsPair(V3609)
var reg313286 Obj
if reg313257 == True {
reg313258 := PrimHead(V3609)
A := reg313258
_ = A
reg313259 := PrimTail(V3609)
reg313260 := __e.Call(__defun__shen_4lazyderef, reg313259, V3696)
V3610 := reg313260
_ = V3610
reg313261 := PrimIsPair(V3610)
var reg313284 Obj
if reg313261 == True {
reg313262 := PrimHead(V3610)
Stream := reg313262
_ = Stream
reg313263 := PrimTail(V3610)
reg313264 := __e.Call(__defun__shen_4lazyderef, reg313263, V3696)
V3611 := reg313264
_ = V3611
reg313265 := Nil;
reg313266 := PrimEqual(reg313265, V3611)
var reg313282 Obj
if reg313266 == True {
reg313267 := __e.Call(__defun__shen_4newpv, V3696)
C := reg313267
_ = C
reg313268 := __e.Call(__defun__shen_4incinfs)
_ = reg313268
reg313269 := __e.Call(__defun__shen_4lazyderef, A, V3696)
reg313270 := __e.Call(__defun__shen_4demodulate, reg313269)
reg313271 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313272 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313273 := MakeSymbol("stream")
reg313274 := MakeSymbol("in")
reg313275 := Nil;
reg313276 := PrimCons(reg313274, reg313275)
reg313277 := PrimCons(reg313273, reg313276)
__ctx.TailApply(__defun__shen_4th_d, Stream, reg313277, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__unify, V3694, C, V3696, reg313272)
return
}, 0)
reg313280 := __e.Call(__defun__bind, C, reg313270, V3696, reg313271)
reg313282 = reg313280
} else {
reg313281 := False;
reg313282 = reg313281
}
reg313284 = reg313282
} else {
reg313283 := False;
reg313284 = reg313283
}
reg313286 = reg313284
} else {
reg313285 := False;
reg313286 = reg313285
}
reg313288 = reg313286
} else {
reg313287 := False;
reg313288 = reg313287
}
reg313290 = reg313288
} else {
reg313289 := False;
reg313290 = reg313289
}
Case := reg313290
_ = Case
reg313291 := False;
reg313292 := PrimEqual(Case, reg313291)
var reg313467 Obj
if reg313292 == True {
reg313293 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3612 := reg313293
_ = V3612
reg313294 := PrimIsPair(V3612)
var reg313335 Obj
if reg313294 == True {
reg313295 := PrimHead(V3612)
reg313296 := __e.Call(__defun__shen_4lazyderef, reg313295, V3696)
V3613 := reg313296
_ = V3613
reg313297 := MakeSymbol("set")
reg313298 := PrimEqual(reg313297, V3613)
var reg313333 Obj
if reg313298 == True {
reg313299 := PrimTail(V3612)
reg313300 := __e.Call(__defun__shen_4lazyderef, reg313299, V3696)
V3614 := reg313300
_ = V3614
reg313301 := PrimIsPair(V3614)
var reg313331 Obj
if reg313301 == True {
reg313302 := PrimHead(V3614)
Var := reg313302
_ = Var
reg313303 := PrimTail(V3614)
reg313304 := __e.Call(__defun__shen_4lazyderef, reg313303, V3696)
V3615 := reg313304
_ = V3615
reg313305 := PrimIsPair(V3615)
var reg313329 Obj
if reg313305 == True {
reg313306 := PrimHead(V3615)
Val := reg313306
_ = Val
reg313307 := PrimTail(V3615)
reg313308 := __e.Call(__defun__shen_4lazyderef, reg313307, V3696)
V3616 := reg313308
_ = V3616
reg313309 := Nil;
reg313310 := PrimEqual(reg313309, V3616)
var reg313327 Obj
if reg313310 == True {
reg313311 := __e.Call(__defun__shen_4incinfs)
_ = reg313311
reg313312 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313313 := MakeSymbol("symbol")
reg313314 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313315 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313316 := MakeSymbol("value")
reg313317 := Nil;
reg313318 := PrimCons(Var, reg313317)
reg313319 := PrimCons(reg313316, reg313318)
reg313320 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, Val, V3694, V3695, V3696, V3697)
return
}, 0)
__ctx.TailApply(__defun__shen_4th_d, reg313319, V3694, V3695, V3696, reg313320)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3696, reg313315)
return
}, 0)
__ctx.TailApply(__defun__shen_4th_d, Var, reg313313, V3695, V3696, reg313314)
return
}, 0)
reg313325 := __e.Call(__defun__cut, Throwcontrol, V3696, reg313312)
reg313327 = reg313325
} else {
reg313326 := False;
reg313327 = reg313326
}
reg313329 = reg313327
} else {
reg313328 := False;
reg313329 = reg313328
}
reg313331 = reg313329
} else {
reg313330 := False;
reg313331 = reg313330
}
reg313333 = reg313331
} else {
reg313332 := False;
reg313333 = reg313332
}
reg313335 = reg313333
} else {
reg313334 := False;
reg313335 = reg313334
}
Case := reg313335
_ = Case
reg313336 := False;
reg313337 := PrimEqual(Case, reg313336)
var reg313466 Obj
if reg313337 == True {
reg313338 := __e.Call(__defun__shen_4newpv, V3696)
NewHyp := reg313338
_ = NewHyp
reg313339 := __e.Call(__defun__shen_4incinfs)
_ = reg313339
reg313340 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4th_d, V3693, V3694, NewHyp, V3696, V3697)
return
}, 0)
reg313342 := __e.Call(__defun__shen_4t_d_1hyps, V3695, NewHyp, V3696, reg313340)
Case := reg313342
_ = Case
reg313343 := False;
reg313344 := PrimEqual(Case, reg313343)
var reg313465 Obj
if reg313344 == True {
reg313345 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3617 := reg313345
_ = V3617
reg313346 := PrimIsPair(V3617)
var reg313368 Obj
if reg313346 == True {
reg313347 := PrimHead(V3617)
reg313348 := __e.Call(__defun__shen_4lazyderef, reg313347, V3696)
V3618 := reg313348
_ = V3618
reg313349 := MakeSymbol("define")
reg313350 := PrimEqual(reg313349, V3618)
var reg313366 Obj
if reg313350 == True {
reg313351 := PrimTail(V3617)
reg313352 := __e.Call(__defun__shen_4lazyderef, reg313351, V3696)
V3619 := reg313352
_ = V3619
reg313353 := PrimIsPair(V3619)
var reg313364 Obj
if reg313353 == True {
reg313354 := PrimHead(V3619)
F := reg313354
_ = F
reg313355 := PrimTail(V3619)
X := reg313355
_ = X
reg313356 := __e.Call(__defun__shen_4incinfs)
_ = reg313356
reg313357 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313358 := MakeSymbol("define")
reg313359 := PrimCons(F, X)
reg313360 := PrimCons(reg313358, reg313359)
__ctx.TailApply(__defun__shen_4t_d_1def, reg313360, V3694, V3695, V3696, V3697)
return
}, 0)
reg313362 := __e.Call(__defun__cut, Throwcontrol, V3696, reg313357)
reg313364 = reg313362
} else {
reg313363 := False;
reg313364 = reg313363
}
reg313366 = reg313364
} else {
reg313365 := False;
reg313366 = reg313365
}
reg313368 = reg313366
} else {
reg313367 := False;
reg313368 = reg313367
}
Case := reg313368
_ = Case
reg313369 := False;
reg313370 := PrimEqual(Case, reg313369)
var reg313464 Obj
if reg313370 == True {
reg313371 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3620 := reg313371
_ = V3620
reg313372 := PrimIsPair(V3620)
var reg313394 Obj
if reg313372 == True {
reg313373 := PrimHead(V3620)
reg313374 := __e.Call(__defun__shen_4lazyderef, reg313373, V3696)
V3621 := reg313374
_ = V3621
reg313375 := MakeSymbol("defmacro")
reg313376 := PrimEqual(reg313375, V3621)
var reg313392 Obj
if reg313376 == True {
reg313377 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3622 := reg313377
_ = V3622
reg313378 := MakeSymbol("unit")
reg313379 := PrimEqual(reg313378, V3622)
var reg313390 Obj
if reg313379 == True {
reg313380 := __e.Call(__defun__shen_4incinfs)
_ = reg313380
reg313381 := __e.Call(__defun__cut, Throwcontrol, V3696, V3697)
reg313390 = reg313381
} else {
reg313382 := __e.Call(__defun__shen_4pvar_2, V3622)
var reg313389 Obj
if reg313382 == True {
reg313383 := MakeSymbol("unit")
reg313384 := __e.Call(__defun__shen_4bindv, V3622, reg313383, V3696)
_ = reg313384
reg313385 := __e.Call(__defun__shen_4incinfs)
_ = reg313385
reg313386 := __e.Call(__defun__cut, Throwcontrol, V3696, V3697)
Result := reg313386
_ = Result
reg313387 := __e.Call(__defun__shen_4unbindv, V3622, V3696)
_ = reg313387
reg313389 = Result
} else {
reg313388 := False;
reg313389 = reg313388
}
reg313390 = reg313389
}
reg313392 = reg313390
} else {
reg313391 := False;
reg313392 = reg313391
}
reg313394 = reg313392
} else {
reg313393 := False;
reg313394 = reg313393
}
Case := reg313394
_ = Case
reg313395 := False;
reg313396 := PrimEqual(Case, reg313395)
var reg313463 Obj
if reg313396 == True {
reg313397 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3623 := reg313397
_ = V3623
reg313398 := PrimIsPair(V3623)
var reg313420 Obj
if reg313398 == True {
reg313399 := PrimHead(V3623)
reg313400 := __e.Call(__defun__shen_4lazyderef, reg313399, V3696)
V3624 := reg313400
_ = V3624
reg313401 := MakeSymbol("shen.process-datatype")
reg313402 := PrimEqual(reg313401, V3624)
var reg313418 Obj
if reg313402 == True {
reg313403 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3625 := reg313403
_ = V3625
reg313404 := MakeSymbol("symbol")
reg313405 := PrimEqual(reg313404, V3625)
var reg313416 Obj
if reg313405 == True {
reg313406 := __e.Call(__defun__shen_4incinfs)
_ = reg313406
reg313407 := __e.Call(__defun__thaw, V3697)
reg313416 = reg313407
} else {
reg313408 := __e.Call(__defun__shen_4pvar_2, V3625)
var reg313415 Obj
if reg313408 == True {
reg313409 := MakeSymbol("symbol")
reg313410 := __e.Call(__defun__shen_4bindv, V3625, reg313409, V3696)
_ = reg313410
reg313411 := __e.Call(__defun__shen_4incinfs)
_ = reg313411
reg313412 := __e.Call(__defun__thaw, V3697)
Result := reg313412
_ = Result
reg313413 := __e.Call(__defun__shen_4unbindv, V3625, V3696)
_ = reg313413
reg313415 = Result
} else {
reg313414 := False;
reg313415 = reg313414
}
reg313416 = reg313415
}
reg313418 = reg313416
} else {
reg313417 := False;
reg313418 = reg313417
}
reg313420 = reg313418
} else {
reg313419 := False;
reg313420 = reg313419
}
Case := reg313420
_ = Case
reg313421 := False;
reg313422 := PrimEqual(Case, reg313421)
var reg313462 Obj
if reg313422 == True {
reg313423 := __e.Call(__defun__shen_4lazyderef, V3693, V3696)
V3626 := reg313423
_ = V3626
reg313424 := PrimIsPair(V3626)
var reg313446 Obj
if reg313424 == True {
reg313425 := PrimHead(V3626)
reg313426 := __e.Call(__defun__shen_4lazyderef, reg313425, V3696)
V3627 := reg313426
_ = V3627
reg313427 := MakeSymbol("shen.synonyms-help")
reg313428 := PrimEqual(reg313427, V3627)
var reg313444 Obj
if reg313428 == True {
reg313429 := __e.Call(__defun__shen_4lazyderef, V3694, V3696)
V3628 := reg313429
_ = V3628
reg313430 := MakeSymbol("symbol")
reg313431 := PrimEqual(reg313430, V3628)
var reg313442 Obj
if reg313431 == True {
reg313432 := __e.Call(__defun__shen_4incinfs)
_ = reg313432
reg313433 := __e.Call(__defun__thaw, V3697)
reg313442 = reg313433
} else {
reg313434 := __e.Call(__defun__shen_4pvar_2, V3628)
var reg313441 Obj
if reg313434 == True {
reg313435 := MakeSymbol("symbol")
reg313436 := __e.Call(__defun__shen_4bindv, V3628, reg313435, V3696)
_ = reg313436
reg313437 := __e.Call(__defun__shen_4incinfs)
_ = reg313437
reg313438 := __e.Call(__defun__thaw, V3697)
Result := reg313438
_ = Result
reg313439 := __e.Call(__defun__shen_4unbindv, V3628, V3696)
_ = reg313439
reg313441 = Result
} else {
reg313440 := False;
reg313441 = reg313440
}
reg313442 = reg313441
}
reg313444 = reg313442
} else {
reg313443 := False;
reg313444 = reg313443
}
reg313446 = reg313444
} else {
reg313445 := False;
reg313446 = reg313445
}
Case := reg313446
_ = Case
reg313447 := False;
reg313448 := PrimEqual(Case, reg313447)
var reg313461 Obj
if reg313448 == True {
reg313449 := __e.Call(__defun__shen_4newpv, V3696)
Datatypes := reg313449
_ = Datatypes
reg313450 := __e.Call(__defun__shen_4incinfs)
_ = reg313450
reg313451 := MakeSymbol("shen.*datatypes*")
reg313452 := PrimValue(reg313451)
reg313453 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg313454 := MakeSymbol(":")
reg313455 := Nil;
reg313456 := PrimCons(V3694, reg313455)
reg313457 := PrimCons(reg313454, reg313456)
reg313458 := PrimCons(V3693, reg313457)
__ctx.TailApply(__defun__shen_4udefs_d, reg313458, V3695, Datatypes, V3696, V3697)
return
}, 0)
reg313460 := __e.Call(__defun__bind, Datatypes, reg313452, V3696, reg313453)
reg313461 = reg313460
} else {
reg313461 = Case
}
reg313462 = reg313461
} else {
reg313462 = Case
}
reg313463 = reg313462
} else {
reg313463 = Case
}
reg313464 = reg313463
} else {
reg313464 = Case
}
reg313465 = reg313464
} else {
reg313465 = Case
}
reg313466 = reg313465
} else {
reg313466 = Case
}
reg313467 = reg313466
} else {
reg313467 = Case
}
reg313468 = reg313467
} else {
reg313468 = Case
}
reg313469 = reg313468
} else {
reg313469 = Case
}
reg313470 = reg313469
} else {
reg313470 = Case
}
reg313471 = reg313470
} else {
reg313471 = Case
}
reg313472 = reg313471
} else {
reg313472 = Case
}
reg313473 = reg313472
} else {
reg313473 = Case
}
reg313474 = reg313473
} else {
reg313474 = Case
}
reg313475 = reg313474
} else {
reg313475 = Case
}
reg313476 = reg313475
} else {
reg313476 = Case
}
reg313477 = reg313476
} else {
reg313477 = Case
}
reg313478 = reg313477
} else {
reg313478 = Case
}
reg313479 = reg313478
} else {
reg313479 = Case
}
reg313480 = reg313479
} else {
reg313480 = Case
}
reg313481 = reg313480
} else {
reg313481 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg313481)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.th*", value: __defun__shen_4th_d})

__defun__shen_4t_d_1hyps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3702 := __args[0]
_ = V3702
V3703 := __args[1]
_ = V3703
V3704 := __args[2]
_ = V3704
V3705 := __args[3]
_ = V3705
reg313483 := __e.Call(__defun__shen_4lazyderef, V3702, V3704)
V3439 := reg313483
_ = V3439
reg313484 := PrimIsPair(V3439)
var reg314022 Obj
if reg313484 == True {
reg313485 := PrimHead(V3439)
reg313486 := __e.Call(__defun__shen_4lazyderef, reg313485, V3704)
V3440 := reg313486
_ = V3440
reg313487 := PrimIsPair(V3440)
var reg314020 Obj
if reg313487 == True {
reg313488 := PrimHead(V3440)
reg313489 := __e.Call(__defun__shen_4lazyderef, reg313488, V3704)
V3441 := reg313489
_ = V3441
reg313490 := PrimIsPair(V3441)
var reg314018 Obj
if reg313490 == True {
reg313491 := PrimHead(V3441)
reg313492 := __e.Call(__defun__shen_4lazyderef, reg313491, V3704)
V3442 := reg313492
_ = V3442
reg313493 := MakeSymbol("cons")
reg313494 := PrimEqual(reg313493, V3442)
var reg314016 Obj
if reg313494 == True {
reg313495 := PrimTail(V3441)
reg313496 := __e.Call(__defun__shen_4lazyderef, reg313495, V3704)
V3443 := reg313496
_ = V3443
reg313497 := PrimIsPair(V3443)
var reg314014 Obj
if reg313497 == True {
reg313498 := PrimHead(V3443)
X := reg313498
_ = X
reg313499 := PrimTail(V3443)
reg313500 := __e.Call(__defun__shen_4lazyderef, reg313499, V3704)
V3444 := reg313500
_ = V3444
reg313501 := PrimIsPair(V3444)
var reg314012 Obj
if reg313501 == True {
reg313502 := PrimHead(V3444)
Y := reg313502
_ = Y
reg313503 := PrimTail(V3444)
reg313504 := __e.Call(__defun__shen_4lazyderef, reg313503, V3704)
V3445 := reg313504
_ = V3445
reg313505 := Nil;
reg313506 := PrimEqual(reg313505, V3445)
var reg314010 Obj
if reg313506 == True {
reg313507 := PrimTail(V3440)
reg313508 := __e.Call(__defun__shen_4lazyderef, reg313507, V3704)
V3446 := reg313508
_ = V3446
reg313509 := PrimIsPair(V3446)
var reg314008 Obj
if reg313509 == True {
reg313510 := PrimHead(V3446)
reg313511 := __e.Call(__defun__shen_4lazyderef, reg313510, V3704)
V3447 := reg313511
_ = V3447
reg313512 := MakeSymbol(":")
reg313513 := PrimEqual(reg313512, V3447)
var reg314006 Obj
if reg313513 == True {
reg313514 := PrimTail(V3446)
reg313515 := __e.Call(__defun__shen_4lazyderef, reg313514, V3704)
V3448 := reg313515
_ = V3448
reg313516 := PrimIsPair(V3448)
var reg314004 Obj
if reg313516 == True {
reg313517 := PrimHead(V3448)
reg313518 := __e.Call(__defun__shen_4lazyderef, reg313517, V3704)
V3449 := reg313518
_ = V3449
reg313519 := PrimIsPair(V3449)
var reg314002 Obj
if reg313519 == True {
reg313520 := PrimHead(V3449)
reg313521 := __e.Call(__defun__shen_4lazyderef, reg313520, V3704)
V3450 := reg313521
_ = V3450
reg313522 := MakeSymbol("list")
reg313523 := PrimEqual(reg313522, V3450)
var reg313932 Obj
if reg313523 == True {
reg313524 := PrimTail(V3449)
reg313525 := __e.Call(__defun__shen_4lazyderef, reg313524, V3704)
V3451 := reg313525
_ = V3451
reg313526 := PrimIsPair(V3451)
var reg313724 Obj
if reg313526 == True {
reg313527 := PrimHead(V3451)
A := reg313527
_ = A
reg313528 := PrimTail(V3451)
reg313529 := __e.Call(__defun__shen_4lazyderef, reg313528, V3704)
V3452 := reg313529
_ = V3452
reg313530 := Nil;
reg313531 := PrimEqual(reg313530, V3452)
var reg313656 Obj
if reg313531 == True {
reg313532 := PrimTail(V3448)
reg313533 := __e.Call(__defun__shen_4lazyderef, reg313532, V3704)
V3453 := reg313533
_ = V3453
reg313534 := Nil;
reg313535 := PrimEqual(reg313534, V3453)
var reg313590 Obj
if reg313535 == True {
reg313536 := PrimTail(V3439)
Hyp := reg313536
_ = Hyp
reg313537 := __e.Call(__defun__shen_4incinfs)
_ = reg313537
reg313538 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313539 := MakeSymbol(":")
reg313540 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313541 := Nil;
reg313542 := PrimCons(reg313540, reg313541)
reg313543 := PrimCons(reg313539, reg313542)
reg313544 := PrimCons(reg313538, reg313543)
reg313545 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313546 := MakeSymbol(":")
reg313547 := MakeSymbol("list")
reg313548 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313549 := Nil;
reg313550 := PrimCons(reg313548, reg313549)
reg313551 := PrimCons(reg313547, reg313550)
reg313552 := Nil;
reg313553 := PrimCons(reg313551, reg313552)
reg313554 := PrimCons(reg313546, reg313553)
reg313555 := PrimCons(reg313545, reg313554)
reg313556 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313557 := PrimCons(reg313555, reg313556)
reg313558 := PrimCons(reg313544, reg313557)
reg313559 := __e.Call(__defun__bind, V3703, reg313558, V3704, V3705)
reg313590 = reg313559
} else {
reg313560 := __e.Call(__defun__shen_4pvar_2, V3453)
var reg313589 Obj
if reg313560 == True {
reg313561 := Nil;
reg313562 := __e.Call(__defun__shen_4bindv, V3453, reg313561, V3704)
_ = reg313562
reg313563 := PrimTail(V3439)
Hyp := reg313563
_ = Hyp
reg313564 := __e.Call(__defun__shen_4incinfs)
_ = reg313564
reg313565 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313566 := MakeSymbol(":")
reg313567 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313568 := Nil;
reg313569 := PrimCons(reg313567, reg313568)
reg313570 := PrimCons(reg313566, reg313569)
reg313571 := PrimCons(reg313565, reg313570)
reg313572 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313573 := MakeSymbol(":")
reg313574 := MakeSymbol("list")
reg313575 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313576 := Nil;
reg313577 := PrimCons(reg313575, reg313576)
reg313578 := PrimCons(reg313574, reg313577)
reg313579 := Nil;
reg313580 := PrimCons(reg313578, reg313579)
reg313581 := PrimCons(reg313573, reg313580)
reg313582 := PrimCons(reg313572, reg313581)
reg313583 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313584 := PrimCons(reg313582, reg313583)
reg313585 := PrimCons(reg313571, reg313584)
reg313586 := __e.Call(__defun__bind, V3703, reg313585, V3704, V3705)
Result := reg313586
_ = Result
reg313587 := __e.Call(__defun__shen_4unbindv, V3453, V3704)
_ = reg313587
reg313589 = Result
} else {
reg313588 := False;
reg313589 = reg313588
}
reg313590 = reg313589
}
reg313656 = reg313590
} else {
reg313591 := __e.Call(__defun__shen_4pvar_2, V3452)
var reg313655 Obj
if reg313591 == True {
reg313592 := Nil;
reg313593 := __e.Call(__defun__shen_4bindv, V3452, reg313592, V3704)
_ = reg313593
reg313594 := PrimTail(V3448)
reg313595 := __e.Call(__defun__shen_4lazyderef, reg313594, V3704)
V3454 := reg313595
_ = V3454
reg313596 := Nil;
reg313597 := PrimEqual(reg313596, V3454)
var reg313652 Obj
if reg313597 == True {
reg313598 := PrimTail(V3439)
Hyp := reg313598
_ = Hyp
reg313599 := __e.Call(__defun__shen_4incinfs)
_ = reg313599
reg313600 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313601 := MakeSymbol(":")
reg313602 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313603 := Nil;
reg313604 := PrimCons(reg313602, reg313603)
reg313605 := PrimCons(reg313601, reg313604)
reg313606 := PrimCons(reg313600, reg313605)
reg313607 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313608 := MakeSymbol(":")
reg313609 := MakeSymbol("list")
reg313610 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313611 := Nil;
reg313612 := PrimCons(reg313610, reg313611)
reg313613 := PrimCons(reg313609, reg313612)
reg313614 := Nil;
reg313615 := PrimCons(reg313613, reg313614)
reg313616 := PrimCons(reg313608, reg313615)
reg313617 := PrimCons(reg313607, reg313616)
reg313618 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313619 := PrimCons(reg313617, reg313618)
reg313620 := PrimCons(reg313606, reg313619)
reg313621 := __e.Call(__defun__bind, V3703, reg313620, V3704, V3705)
reg313652 = reg313621
} else {
reg313622 := __e.Call(__defun__shen_4pvar_2, V3454)
var reg313651 Obj
if reg313622 == True {
reg313623 := Nil;
reg313624 := __e.Call(__defun__shen_4bindv, V3454, reg313623, V3704)
_ = reg313624
reg313625 := PrimTail(V3439)
Hyp := reg313625
_ = Hyp
reg313626 := __e.Call(__defun__shen_4incinfs)
_ = reg313626
reg313627 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313628 := MakeSymbol(":")
reg313629 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313630 := Nil;
reg313631 := PrimCons(reg313629, reg313630)
reg313632 := PrimCons(reg313628, reg313631)
reg313633 := PrimCons(reg313627, reg313632)
reg313634 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313635 := MakeSymbol(":")
reg313636 := MakeSymbol("list")
reg313637 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313638 := Nil;
reg313639 := PrimCons(reg313637, reg313638)
reg313640 := PrimCons(reg313636, reg313639)
reg313641 := Nil;
reg313642 := PrimCons(reg313640, reg313641)
reg313643 := PrimCons(reg313635, reg313642)
reg313644 := PrimCons(reg313634, reg313643)
reg313645 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313646 := PrimCons(reg313644, reg313645)
reg313647 := PrimCons(reg313633, reg313646)
reg313648 := __e.Call(__defun__bind, V3703, reg313647, V3704, V3705)
Result := reg313648
_ = Result
reg313649 := __e.Call(__defun__shen_4unbindv, V3454, V3704)
_ = reg313649
reg313651 = Result
} else {
reg313650 := False;
reg313651 = reg313650
}
reg313652 = reg313651
}
Result := reg313652
_ = Result
reg313653 := __e.Call(__defun__shen_4unbindv, V3452, V3704)
_ = reg313653
reg313655 = Result
} else {
reg313654 := False;
reg313655 = reg313654
}
reg313656 = reg313655
}
reg313724 = reg313656
} else {
reg313657 := __e.Call(__defun__shen_4pvar_2, V3451)
var reg313723 Obj
if reg313657 == True {
reg313658 := __e.Call(__defun__shen_4newpv, V3704)
A := reg313658
_ = A
reg313659 := Nil;
reg313660 := PrimCons(A, reg313659)
reg313661 := __e.Call(__defun__shen_4bindv, V3451, reg313660, V3704)
_ = reg313661
reg313662 := PrimTail(V3448)
reg313663 := __e.Call(__defun__shen_4lazyderef, reg313662, V3704)
V3455 := reg313663
_ = V3455
reg313664 := Nil;
reg313665 := PrimEqual(reg313664, V3455)
var reg313720 Obj
if reg313665 == True {
reg313666 := PrimTail(V3439)
Hyp := reg313666
_ = Hyp
reg313667 := __e.Call(__defun__shen_4incinfs)
_ = reg313667
reg313668 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313669 := MakeSymbol(":")
reg313670 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313671 := Nil;
reg313672 := PrimCons(reg313670, reg313671)
reg313673 := PrimCons(reg313669, reg313672)
reg313674 := PrimCons(reg313668, reg313673)
reg313675 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313676 := MakeSymbol(":")
reg313677 := MakeSymbol("list")
reg313678 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313679 := Nil;
reg313680 := PrimCons(reg313678, reg313679)
reg313681 := PrimCons(reg313677, reg313680)
reg313682 := Nil;
reg313683 := PrimCons(reg313681, reg313682)
reg313684 := PrimCons(reg313676, reg313683)
reg313685 := PrimCons(reg313675, reg313684)
reg313686 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313687 := PrimCons(reg313685, reg313686)
reg313688 := PrimCons(reg313674, reg313687)
reg313689 := __e.Call(__defun__bind, V3703, reg313688, V3704, V3705)
reg313720 = reg313689
} else {
reg313690 := __e.Call(__defun__shen_4pvar_2, V3455)
var reg313719 Obj
if reg313690 == True {
reg313691 := Nil;
reg313692 := __e.Call(__defun__shen_4bindv, V3455, reg313691, V3704)
_ = reg313692
reg313693 := PrimTail(V3439)
Hyp := reg313693
_ = Hyp
reg313694 := __e.Call(__defun__shen_4incinfs)
_ = reg313694
reg313695 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313696 := MakeSymbol(":")
reg313697 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313698 := Nil;
reg313699 := PrimCons(reg313697, reg313698)
reg313700 := PrimCons(reg313696, reg313699)
reg313701 := PrimCons(reg313695, reg313700)
reg313702 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313703 := MakeSymbol(":")
reg313704 := MakeSymbol("list")
reg313705 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313706 := Nil;
reg313707 := PrimCons(reg313705, reg313706)
reg313708 := PrimCons(reg313704, reg313707)
reg313709 := Nil;
reg313710 := PrimCons(reg313708, reg313709)
reg313711 := PrimCons(reg313703, reg313710)
reg313712 := PrimCons(reg313702, reg313711)
reg313713 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313714 := PrimCons(reg313712, reg313713)
reg313715 := PrimCons(reg313701, reg313714)
reg313716 := __e.Call(__defun__bind, V3703, reg313715, V3704, V3705)
Result := reg313716
_ = Result
reg313717 := __e.Call(__defun__shen_4unbindv, V3455, V3704)
_ = reg313717
reg313719 = Result
} else {
reg313718 := False;
reg313719 = reg313718
}
reg313720 = reg313719
}
Result := reg313720
_ = Result
reg313721 := __e.Call(__defun__shen_4unbindv, V3451, V3704)
_ = reg313721
reg313723 = Result
} else {
reg313722 := False;
reg313723 = reg313722
}
reg313724 = reg313723
}
reg313932 = reg313724
} else {
reg313725 := __e.Call(__defun__shen_4pvar_2, V3450)
var reg313931 Obj
if reg313725 == True {
reg313726 := MakeSymbol("list")
reg313727 := __e.Call(__defun__shen_4bindv, V3450, reg313726, V3704)
_ = reg313727
reg313728 := PrimTail(V3449)
reg313729 := __e.Call(__defun__shen_4lazyderef, reg313728, V3704)
V3456 := reg313729
_ = V3456
reg313730 := PrimIsPair(V3456)
var reg313928 Obj
if reg313730 == True {
reg313731 := PrimHead(V3456)
A := reg313731
_ = A
reg313732 := PrimTail(V3456)
reg313733 := __e.Call(__defun__shen_4lazyderef, reg313732, V3704)
V3457 := reg313733
_ = V3457
reg313734 := Nil;
reg313735 := PrimEqual(reg313734, V3457)
var reg313860 Obj
if reg313735 == True {
reg313736 := PrimTail(V3448)
reg313737 := __e.Call(__defun__shen_4lazyderef, reg313736, V3704)
V3458 := reg313737
_ = V3458
reg313738 := Nil;
reg313739 := PrimEqual(reg313738, V3458)
var reg313794 Obj
if reg313739 == True {
reg313740 := PrimTail(V3439)
Hyp := reg313740
_ = Hyp
reg313741 := __e.Call(__defun__shen_4incinfs)
_ = reg313741
reg313742 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313743 := MakeSymbol(":")
reg313744 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313745 := Nil;
reg313746 := PrimCons(reg313744, reg313745)
reg313747 := PrimCons(reg313743, reg313746)
reg313748 := PrimCons(reg313742, reg313747)
reg313749 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313750 := MakeSymbol(":")
reg313751 := MakeSymbol("list")
reg313752 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313753 := Nil;
reg313754 := PrimCons(reg313752, reg313753)
reg313755 := PrimCons(reg313751, reg313754)
reg313756 := Nil;
reg313757 := PrimCons(reg313755, reg313756)
reg313758 := PrimCons(reg313750, reg313757)
reg313759 := PrimCons(reg313749, reg313758)
reg313760 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313761 := PrimCons(reg313759, reg313760)
reg313762 := PrimCons(reg313748, reg313761)
reg313763 := __e.Call(__defun__bind, V3703, reg313762, V3704, V3705)
reg313794 = reg313763
} else {
reg313764 := __e.Call(__defun__shen_4pvar_2, V3458)
var reg313793 Obj
if reg313764 == True {
reg313765 := Nil;
reg313766 := __e.Call(__defun__shen_4bindv, V3458, reg313765, V3704)
_ = reg313766
reg313767 := PrimTail(V3439)
Hyp := reg313767
_ = Hyp
reg313768 := __e.Call(__defun__shen_4incinfs)
_ = reg313768
reg313769 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313770 := MakeSymbol(":")
reg313771 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313772 := Nil;
reg313773 := PrimCons(reg313771, reg313772)
reg313774 := PrimCons(reg313770, reg313773)
reg313775 := PrimCons(reg313769, reg313774)
reg313776 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313777 := MakeSymbol(":")
reg313778 := MakeSymbol("list")
reg313779 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313780 := Nil;
reg313781 := PrimCons(reg313779, reg313780)
reg313782 := PrimCons(reg313778, reg313781)
reg313783 := Nil;
reg313784 := PrimCons(reg313782, reg313783)
reg313785 := PrimCons(reg313777, reg313784)
reg313786 := PrimCons(reg313776, reg313785)
reg313787 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313788 := PrimCons(reg313786, reg313787)
reg313789 := PrimCons(reg313775, reg313788)
reg313790 := __e.Call(__defun__bind, V3703, reg313789, V3704, V3705)
Result := reg313790
_ = Result
reg313791 := __e.Call(__defun__shen_4unbindv, V3458, V3704)
_ = reg313791
reg313793 = Result
} else {
reg313792 := False;
reg313793 = reg313792
}
reg313794 = reg313793
}
reg313860 = reg313794
} else {
reg313795 := __e.Call(__defun__shen_4pvar_2, V3457)
var reg313859 Obj
if reg313795 == True {
reg313796 := Nil;
reg313797 := __e.Call(__defun__shen_4bindv, V3457, reg313796, V3704)
_ = reg313797
reg313798 := PrimTail(V3448)
reg313799 := __e.Call(__defun__shen_4lazyderef, reg313798, V3704)
V3459 := reg313799
_ = V3459
reg313800 := Nil;
reg313801 := PrimEqual(reg313800, V3459)
var reg313856 Obj
if reg313801 == True {
reg313802 := PrimTail(V3439)
Hyp := reg313802
_ = Hyp
reg313803 := __e.Call(__defun__shen_4incinfs)
_ = reg313803
reg313804 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313805 := MakeSymbol(":")
reg313806 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313807 := Nil;
reg313808 := PrimCons(reg313806, reg313807)
reg313809 := PrimCons(reg313805, reg313808)
reg313810 := PrimCons(reg313804, reg313809)
reg313811 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313812 := MakeSymbol(":")
reg313813 := MakeSymbol("list")
reg313814 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313815 := Nil;
reg313816 := PrimCons(reg313814, reg313815)
reg313817 := PrimCons(reg313813, reg313816)
reg313818 := Nil;
reg313819 := PrimCons(reg313817, reg313818)
reg313820 := PrimCons(reg313812, reg313819)
reg313821 := PrimCons(reg313811, reg313820)
reg313822 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313823 := PrimCons(reg313821, reg313822)
reg313824 := PrimCons(reg313810, reg313823)
reg313825 := __e.Call(__defun__bind, V3703, reg313824, V3704, V3705)
reg313856 = reg313825
} else {
reg313826 := __e.Call(__defun__shen_4pvar_2, V3459)
var reg313855 Obj
if reg313826 == True {
reg313827 := Nil;
reg313828 := __e.Call(__defun__shen_4bindv, V3459, reg313827, V3704)
_ = reg313828
reg313829 := PrimTail(V3439)
Hyp := reg313829
_ = Hyp
reg313830 := __e.Call(__defun__shen_4incinfs)
_ = reg313830
reg313831 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313832 := MakeSymbol(":")
reg313833 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313834 := Nil;
reg313835 := PrimCons(reg313833, reg313834)
reg313836 := PrimCons(reg313832, reg313835)
reg313837 := PrimCons(reg313831, reg313836)
reg313838 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313839 := MakeSymbol(":")
reg313840 := MakeSymbol("list")
reg313841 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313842 := Nil;
reg313843 := PrimCons(reg313841, reg313842)
reg313844 := PrimCons(reg313840, reg313843)
reg313845 := Nil;
reg313846 := PrimCons(reg313844, reg313845)
reg313847 := PrimCons(reg313839, reg313846)
reg313848 := PrimCons(reg313838, reg313847)
reg313849 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313850 := PrimCons(reg313848, reg313849)
reg313851 := PrimCons(reg313837, reg313850)
reg313852 := __e.Call(__defun__bind, V3703, reg313851, V3704, V3705)
Result := reg313852
_ = Result
reg313853 := __e.Call(__defun__shen_4unbindv, V3459, V3704)
_ = reg313853
reg313855 = Result
} else {
reg313854 := False;
reg313855 = reg313854
}
reg313856 = reg313855
}
Result := reg313856
_ = Result
reg313857 := __e.Call(__defun__shen_4unbindv, V3457, V3704)
_ = reg313857
reg313859 = Result
} else {
reg313858 := False;
reg313859 = reg313858
}
reg313860 = reg313859
}
reg313928 = reg313860
} else {
reg313861 := __e.Call(__defun__shen_4pvar_2, V3456)
var reg313927 Obj
if reg313861 == True {
reg313862 := __e.Call(__defun__shen_4newpv, V3704)
A := reg313862
_ = A
reg313863 := Nil;
reg313864 := PrimCons(A, reg313863)
reg313865 := __e.Call(__defun__shen_4bindv, V3456, reg313864, V3704)
_ = reg313865
reg313866 := PrimTail(V3448)
reg313867 := __e.Call(__defun__shen_4lazyderef, reg313866, V3704)
V3460 := reg313867
_ = V3460
reg313868 := Nil;
reg313869 := PrimEqual(reg313868, V3460)
var reg313924 Obj
if reg313869 == True {
reg313870 := PrimTail(V3439)
Hyp := reg313870
_ = Hyp
reg313871 := __e.Call(__defun__shen_4incinfs)
_ = reg313871
reg313872 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313873 := MakeSymbol(":")
reg313874 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313875 := Nil;
reg313876 := PrimCons(reg313874, reg313875)
reg313877 := PrimCons(reg313873, reg313876)
reg313878 := PrimCons(reg313872, reg313877)
reg313879 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313880 := MakeSymbol(":")
reg313881 := MakeSymbol("list")
reg313882 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313883 := Nil;
reg313884 := PrimCons(reg313882, reg313883)
reg313885 := PrimCons(reg313881, reg313884)
reg313886 := Nil;
reg313887 := PrimCons(reg313885, reg313886)
reg313888 := PrimCons(reg313880, reg313887)
reg313889 := PrimCons(reg313879, reg313888)
reg313890 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313891 := PrimCons(reg313889, reg313890)
reg313892 := PrimCons(reg313878, reg313891)
reg313893 := __e.Call(__defun__bind, V3703, reg313892, V3704, V3705)
reg313924 = reg313893
} else {
reg313894 := __e.Call(__defun__shen_4pvar_2, V3460)
var reg313923 Obj
if reg313894 == True {
reg313895 := Nil;
reg313896 := __e.Call(__defun__shen_4bindv, V3460, reg313895, V3704)
_ = reg313896
reg313897 := PrimTail(V3439)
Hyp := reg313897
_ = Hyp
reg313898 := __e.Call(__defun__shen_4incinfs)
_ = reg313898
reg313899 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313900 := MakeSymbol(":")
reg313901 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313902 := Nil;
reg313903 := PrimCons(reg313901, reg313902)
reg313904 := PrimCons(reg313900, reg313903)
reg313905 := PrimCons(reg313899, reg313904)
reg313906 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313907 := MakeSymbol(":")
reg313908 := MakeSymbol("list")
reg313909 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313910 := Nil;
reg313911 := PrimCons(reg313909, reg313910)
reg313912 := PrimCons(reg313908, reg313911)
reg313913 := Nil;
reg313914 := PrimCons(reg313912, reg313913)
reg313915 := PrimCons(reg313907, reg313914)
reg313916 := PrimCons(reg313906, reg313915)
reg313917 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313918 := PrimCons(reg313916, reg313917)
reg313919 := PrimCons(reg313905, reg313918)
reg313920 := __e.Call(__defun__bind, V3703, reg313919, V3704, V3705)
Result := reg313920
_ = Result
reg313921 := __e.Call(__defun__shen_4unbindv, V3460, V3704)
_ = reg313921
reg313923 = Result
} else {
reg313922 := False;
reg313923 = reg313922
}
reg313924 = reg313923
}
Result := reg313924
_ = Result
reg313925 := __e.Call(__defun__shen_4unbindv, V3456, V3704)
_ = reg313925
reg313927 = Result
} else {
reg313926 := False;
reg313927 = reg313926
}
reg313928 = reg313927
}
Result := reg313928
_ = Result
reg313929 := __e.Call(__defun__shen_4unbindv, V3450, V3704)
_ = reg313929
reg313931 = Result
} else {
reg313930 := False;
reg313931 = reg313930
}
reg313932 = reg313931
}
reg314002 = reg313932
} else {
reg313933 := __e.Call(__defun__shen_4pvar_2, V3449)
var reg314001 Obj
if reg313933 == True {
reg313934 := __e.Call(__defun__shen_4newpv, V3704)
A := reg313934
_ = A
reg313935 := MakeSymbol("list")
reg313936 := Nil;
reg313937 := PrimCons(A, reg313936)
reg313938 := PrimCons(reg313935, reg313937)
reg313939 := __e.Call(__defun__shen_4bindv, V3449, reg313938, V3704)
_ = reg313939
reg313940 := PrimTail(V3448)
reg313941 := __e.Call(__defun__shen_4lazyderef, reg313940, V3704)
V3461 := reg313941
_ = V3461
reg313942 := Nil;
reg313943 := PrimEqual(reg313942, V3461)
var reg313998 Obj
if reg313943 == True {
reg313944 := PrimTail(V3439)
Hyp := reg313944
_ = Hyp
reg313945 := __e.Call(__defun__shen_4incinfs)
_ = reg313945
reg313946 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313947 := MakeSymbol(":")
reg313948 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313949 := Nil;
reg313950 := PrimCons(reg313948, reg313949)
reg313951 := PrimCons(reg313947, reg313950)
reg313952 := PrimCons(reg313946, reg313951)
reg313953 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313954 := MakeSymbol(":")
reg313955 := MakeSymbol("list")
reg313956 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313957 := Nil;
reg313958 := PrimCons(reg313956, reg313957)
reg313959 := PrimCons(reg313955, reg313958)
reg313960 := Nil;
reg313961 := PrimCons(reg313959, reg313960)
reg313962 := PrimCons(reg313954, reg313961)
reg313963 := PrimCons(reg313953, reg313962)
reg313964 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313965 := PrimCons(reg313963, reg313964)
reg313966 := PrimCons(reg313952, reg313965)
reg313967 := __e.Call(__defun__bind, V3703, reg313966, V3704, V3705)
reg313998 = reg313967
} else {
reg313968 := __e.Call(__defun__shen_4pvar_2, V3461)
var reg313997 Obj
if reg313968 == True {
reg313969 := Nil;
reg313970 := __e.Call(__defun__shen_4bindv, V3461, reg313969, V3704)
_ = reg313970
reg313971 := PrimTail(V3439)
Hyp := reg313971
_ = Hyp
reg313972 := __e.Call(__defun__shen_4incinfs)
_ = reg313972
reg313973 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg313974 := MakeSymbol(":")
reg313975 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313976 := Nil;
reg313977 := PrimCons(reg313975, reg313976)
reg313978 := PrimCons(reg313974, reg313977)
reg313979 := PrimCons(reg313973, reg313978)
reg313980 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg313981 := MakeSymbol(":")
reg313982 := MakeSymbol("list")
reg313983 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg313984 := Nil;
reg313985 := PrimCons(reg313983, reg313984)
reg313986 := PrimCons(reg313982, reg313985)
reg313987 := Nil;
reg313988 := PrimCons(reg313986, reg313987)
reg313989 := PrimCons(reg313981, reg313988)
reg313990 := PrimCons(reg313980, reg313989)
reg313991 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg313992 := PrimCons(reg313990, reg313991)
reg313993 := PrimCons(reg313979, reg313992)
reg313994 := __e.Call(__defun__bind, V3703, reg313993, V3704, V3705)
Result := reg313994
_ = Result
reg313995 := __e.Call(__defun__shen_4unbindv, V3461, V3704)
_ = reg313995
reg313997 = Result
} else {
reg313996 := False;
reg313997 = reg313996
}
reg313998 = reg313997
}
Result := reg313998
_ = Result
reg313999 := __e.Call(__defun__shen_4unbindv, V3449, V3704)
_ = reg313999
reg314001 = Result
} else {
reg314000 := False;
reg314001 = reg314000
}
reg314002 = reg314001
}
reg314004 = reg314002
} else {
reg314003 := False;
reg314004 = reg314003
}
reg314006 = reg314004
} else {
reg314005 := False;
reg314006 = reg314005
}
reg314008 = reg314006
} else {
reg314007 := False;
reg314008 = reg314007
}
reg314010 = reg314008
} else {
reg314009 := False;
reg314010 = reg314009
}
reg314012 = reg314010
} else {
reg314011 := False;
reg314012 = reg314011
}
reg314014 = reg314012
} else {
reg314013 := False;
reg314014 = reg314013
}
reg314016 = reg314014
} else {
reg314015 := False;
reg314016 = reg314015
}
reg314018 = reg314016
} else {
reg314017 := False;
reg314018 = reg314017
}
reg314020 = reg314018
} else {
reg314019 := False;
reg314020 = reg314019
}
reg314022 = reg314020
} else {
reg314021 := False;
reg314022 = reg314021
}
Case := reg314022
_ = Case
reg314023 := False;
reg314024 := PrimEqual(Case, reg314023)
if reg314024 == True {
reg314025 := __e.Call(__defun__shen_4lazyderef, V3702, V3704)
V3462 := reg314025
_ = V3462
reg314026 := PrimIsPair(V3462)
var reg314576 Obj
if reg314026 == True {
reg314027 := PrimHead(V3462)
reg314028 := __e.Call(__defun__shen_4lazyderef, reg314027, V3704)
V3463 := reg314028
_ = V3463
reg314029 := PrimIsPair(V3463)
var reg314574 Obj
if reg314029 == True {
reg314030 := PrimHead(V3463)
reg314031 := __e.Call(__defun__shen_4lazyderef, reg314030, V3704)
V3464 := reg314031
_ = V3464
reg314032 := PrimIsPair(V3464)
var reg314572 Obj
if reg314032 == True {
reg314033 := PrimHead(V3464)
reg314034 := __e.Call(__defun__shen_4lazyderef, reg314033, V3704)
V3465 := reg314034
_ = V3465
reg314035 := MakeSymbol("@p")
reg314036 := PrimEqual(reg314035, V3465)
var reg314570 Obj
if reg314036 == True {
reg314037 := PrimTail(V3464)
reg314038 := __e.Call(__defun__shen_4lazyderef, reg314037, V3704)
V3466 := reg314038
_ = V3466
reg314039 := PrimIsPair(V3466)
var reg314568 Obj
if reg314039 == True {
reg314040 := PrimHead(V3466)
X := reg314040
_ = X
reg314041 := PrimTail(V3466)
reg314042 := __e.Call(__defun__shen_4lazyderef, reg314041, V3704)
V3467 := reg314042
_ = V3467
reg314043 := PrimIsPair(V3467)
var reg314566 Obj
if reg314043 == True {
reg314044 := PrimHead(V3467)
Y := reg314044
_ = Y
reg314045 := PrimTail(V3467)
reg314046 := __e.Call(__defun__shen_4lazyderef, reg314045, V3704)
V3468 := reg314046
_ = V3468
reg314047 := Nil;
reg314048 := PrimEqual(reg314047, V3468)
var reg314564 Obj
if reg314048 == True {
reg314049 := PrimTail(V3463)
reg314050 := __e.Call(__defun__shen_4lazyderef, reg314049, V3704)
V3469 := reg314050
_ = V3469
reg314051 := PrimIsPair(V3469)
var reg314562 Obj
if reg314051 == True {
reg314052 := PrimHead(V3469)
reg314053 := __e.Call(__defun__shen_4lazyderef, reg314052, V3704)
V3470 := reg314053
_ = V3470
reg314054 := MakeSymbol(":")
reg314055 := PrimEqual(reg314054, V3470)
var reg314560 Obj
if reg314055 == True {
reg314056 := PrimTail(V3469)
reg314057 := __e.Call(__defun__shen_4lazyderef, reg314056, V3704)
V3471 := reg314057
_ = V3471
reg314058 := PrimIsPair(V3471)
var reg314558 Obj
if reg314058 == True {
reg314059 := PrimHead(V3471)
reg314060 := __e.Call(__defun__shen_4lazyderef, reg314059, V3704)
V3472 := reg314060
_ = V3472
reg314061 := PrimIsPair(V3472)
var reg314556 Obj
if reg314061 == True {
reg314062 := PrimHead(V3472)
A := reg314062
_ = A
reg314063 := PrimTail(V3472)
reg314064 := __e.Call(__defun__shen_4lazyderef, reg314063, V3704)
V3473 := reg314064
_ = V3473
reg314065 := PrimIsPair(V3473)
var reg314492 Obj
if reg314065 == True {
reg314066 := PrimHead(V3473)
reg314067 := __e.Call(__defun__shen_4lazyderef, reg314066, V3704)
V3474 := reg314067
_ = V3474
reg314068 := MakeSymbol("*")
reg314069 := PrimEqual(reg314068, V3474)
var reg314430 Obj
if reg314069 == True {
reg314070 := PrimTail(V3473)
reg314071 := __e.Call(__defun__shen_4lazyderef, reg314070, V3704)
V3475 := reg314071
_ = V3475
reg314072 := PrimIsPair(V3475)
var reg314246 Obj
if reg314072 == True {
reg314073 := PrimHead(V3475)
B := reg314073
_ = B
reg314074 := PrimTail(V3475)
reg314075 := __e.Call(__defun__shen_4lazyderef, reg314074, V3704)
V3476 := reg314075
_ = V3476
reg314076 := Nil;
reg314077 := PrimEqual(reg314076, V3476)
var reg314186 Obj
if reg314077 == True {
reg314078 := PrimTail(V3471)
reg314079 := __e.Call(__defun__shen_4lazyderef, reg314078, V3704)
V3477 := reg314079
_ = V3477
reg314080 := Nil;
reg314081 := PrimEqual(reg314080, V3477)
var reg314128 Obj
if reg314081 == True {
reg314082 := PrimTail(V3462)
Hyp := reg314082
_ = Hyp
reg314083 := __e.Call(__defun__shen_4incinfs)
_ = reg314083
reg314084 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314085 := MakeSymbol(":")
reg314086 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314087 := Nil;
reg314088 := PrimCons(reg314086, reg314087)
reg314089 := PrimCons(reg314085, reg314088)
reg314090 := PrimCons(reg314084, reg314089)
reg314091 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314092 := MakeSymbol(":")
reg314093 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314094 := Nil;
reg314095 := PrimCons(reg314093, reg314094)
reg314096 := PrimCons(reg314092, reg314095)
reg314097 := PrimCons(reg314091, reg314096)
reg314098 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314099 := PrimCons(reg314097, reg314098)
reg314100 := PrimCons(reg314090, reg314099)
reg314101 := __e.Call(__defun__bind, V3703, reg314100, V3704, V3705)
reg314128 = reg314101
} else {
reg314102 := __e.Call(__defun__shen_4pvar_2, V3477)
var reg314127 Obj
if reg314102 == True {
reg314103 := Nil;
reg314104 := __e.Call(__defun__shen_4bindv, V3477, reg314103, V3704)
_ = reg314104
reg314105 := PrimTail(V3462)
Hyp := reg314105
_ = Hyp
reg314106 := __e.Call(__defun__shen_4incinfs)
_ = reg314106
reg314107 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314108 := MakeSymbol(":")
reg314109 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314110 := Nil;
reg314111 := PrimCons(reg314109, reg314110)
reg314112 := PrimCons(reg314108, reg314111)
reg314113 := PrimCons(reg314107, reg314112)
reg314114 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314115 := MakeSymbol(":")
reg314116 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314117 := Nil;
reg314118 := PrimCons(reg314116, reg314117)
reg314119 := PrimCons(reg314115, reg314118)
reg314120 := PrimCons(reg314114, reg314119)
reg314121 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314122 := PrimCons(reg314120, reg314121)
reg314123 := PrimCons(reg314113, reg314122)
reg314124 := __e.Call(__defun__bind, V3703, reg314123, V3704, V3705)
Result := reg314124
_ = Result
reg314125 := __e.Call(__defun__shen_4unbindv, V3477, V3704)
_ = reg314125
reg314127 = Result
} else {
reg314126 := False;
reg314127 = reg314126
}
reg314128 = reg314127
}
reg314186 = reg314128
} else {
reg314129 := __e.Call(__defun__shen_4pvar_2, V3476)
var reg314185 Obj
if reg314129 == True {
reg314130 := Nil;
reg314131 := __e.Call(__defun__shen_4bindv, V3476, reg314130, V3704)
_ = reg314131
reg314132 := PrimTail(V3471)
reg314133 := __e.Call(__defun__shen_4lazyderef, reg314132, V3704)
V3478 := reg314133
_ = V3478
reg314134 := Nil;
reg314135 := PrimEqual(reg314134, V3478)
var reg314182 Obj
if reg314135 == True {
reg314136 := PrimTail(V3462)
Hyp := reg314136
_ = Hyp
reg314137 := __e.Call(__defun__shen_4incinfs)
_ = reg314137
reg314138 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314139 := MakeSymbol(":")
reg314140 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314141 := Nil;
reg314142 := PrimCons(reg314140, reg314141)
reg314143 := PrimCons(reg314139, reg314142)
reg314144 := PrimCons(reg314138, reg314143)
reg314145 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314146 := MakeSymbol(":")
reg314147 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314148 := Nil;
reg314149 := PrimCons(reg314147, reg314148)
reg314150 := PrimCons(reg314146, reg314149)
reg314151 := PrimCons(reg314145, reg314150)
reg314152 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314153 := PrimCons(reg314151, reg314152)
reg314154 := PrimCons(reg314144, reg314153)
reg314155 := __e.Call(__defun__bind, V3703, reg314154, V3704, V3705)
reg314182 = reg314155
} else {
reg314156 := __e.Call(__defun__shen_4pvar_2, V3478)
var reg314181 Obj
if reg314156 == True {
reg314157 := Nil;
reg314158 := __e.Call(__defun__shen_4bindv, V3478, reg314157, V3704)
_ = reg314158
reg314159 := PrimTail(V3462)
Hyp := reg314159
_ = Hyp
reg314160 := __e.Call(__defun__shen_4incinfs)
_ = reg314160
reg314161 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314162 := MakeSymbol(":")
reg314163 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314164 := Nil;
reg314165 := PrimCons(reg314163, reg314164)
reg314166 := PrimCons(reg314162, reg314165)
reg314167 := PrimCons(reg314161, reg314166)
reg314168 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314169 := MakeSymbol(":")
reg314170 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314171 := Nil;
reg314172 := PrimCons(reg314170, reg314171)
reg314173 := PrimCons(reg314169, reg314172)
reg314174 := PrimCons(reg314168, reg314173)
reg314175 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314176 := PrimCons(reg314174, reg314175)
reg314177 := PrimCons(reg314167, reg314176)
reg314178 := __e.Call(__defun__bind, V3703, reg314177, V3704, V3705)
Result := reg314178
_ = Result
reg314179 := __e.Call(__defun__shen_4unbindv, V3478, V3704)
_ = reg314179
reg314181 = Result
} else {
reg314180 := False;
reg314181 = reg314180
}
reg314182 = reg314181
}
Result := reg314182
_ = Result
reg314183 := __e.Call(__defun__shen_4unbindv, V3476, V3704)
_ = reg314183
reg314185 = Result
} else {
reg314184 := False;
reg314185 = reg314184
}
reg314186 = reg314185
}
reg314246 = reg314186
} else {
reg314187 := __e.Call(__defun__shen_4pvar_2, V3475)
var reg314245 Obj
if reg314187 == True {
reg314188 := __e.Call(__defun__shen_4newpv, V3704)
B := reg314188
_ = B
reg314189 := Nil;
reg314190 := PrimCons(B, reg314189)
reg314191 := __e.Call(__defun__shen_4bindv, V3475, reg314190, V3704)
_ = reg314191
reg314192 := PrimTail(V3471)
reg314193 := __e.Call(__defun__shen_4lazyderef, reg314192, V3704)
V3479 := reg314193
_ = V3479
reg314194 := Nil;
reg314195 := PrimEqual(reg314194, V3479)
var reg314242 Obj
if reg314195 == True {
reg314196 := PrimTail(V3462)
Hyp := reg314196
_ = Hyp
reg314197 := __e.Call(__defun__shen_4incinfs)
_ = reg314197
reg314198 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314199 := MakeSymbol(":")
reg314200 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314201 := Nil;
reg314202 := PrimCons(reg314200, reg314201)
reg314203 := PrimCons(reg314199, reg314202)
reg314204 := PrimCons(reg314198, reg314203)
reg314205 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314206 := MakeSymbol(":")
reg314207 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314208 := Nil;
reg314209 := PrimCons(reg314207, reg314208)
reg314210 := PrimCons(reg314206, reg314209)
reg314211 := PrimCons(reg314205, reg314210)
reg314212 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314213 := PrimCons(reg314211, reg314212)
reg314214 := PrimCons(reg314204, reg314213)
reg314215 := __e.Call(__defun__bind, V3703, reg314214, V3704, V3705)
reg314242 = reg314215
} else {
reg314216 := __e.Call(__defun__shen_4pvar_2, V3479)
var reg314241 Obj
if reg314216 == True {
reg314217 := Nil;
reg314218 := __e.Call(__defun__shen_4bindv, V3479, reg314217, V3704)
_ = reg314218
reg314219 := PrimTail(V3462)
Hyp := reg314219
_ = Hyp
reg314220 := __e.Call(__defun__shen_4incinfs)
_ = reg314220
reg314221 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314222 := MakeSymbol(":")
reg314223 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314224 := Nil;
reg314225 := PrimCons(reg314223, reg314224)
reg314226 := PrimCons(reg314222, reg314225)
reg314227 := PrimCons(reg314221, reg314226)
reg314228 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314229 := MakeSymbol(":")
reg314230 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314231 := Nil;
reg314232 := PrimCons(reg314230, reg314231)
reg314233 := PrimCons(reg314229, reg314232)
reg314234 := PrimCons(reg314228, reg314233)
reg314235 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314236 := PrimCons(reg314234, reg314235)
reg314237 := PrimCons(reg314227, reg314236)
reg314238 := __e.Call(__defun__bind, V3703, reg314237, V3704, V3705)
Result := reg314238
_ = Result
reg314239 := __e.Call(__defun__shen_4unbindv, V3479, V3704)
_ = reg314239
reg314241 = Result
} else {
reg314240 := False;
reg314241 = reg314240
}
reg314242 = reg314241
}
Result := reg314242
_ = Result
reg314243 := __e.Call(__defun__shen_4unbindv, V3475, V3704)
_ = reg314243
reg314245 = Result
} else {
reg314244 := False;
reg314245 = reg314244
}
reg314246 = reg314245
}
reg314430 = reg314246
} else {
reg314247 := __e.Call(__defun__shen_4pvar_2, V3474)
var reg314429 Obj
if reg314247 == True {
reg314248 := MakeSymbol("*")
reg314249 := __e.Call(__defun__shen_4bindv, V3474, reg314248, V3704)
_ = reg314249
reg314250 := PrimTail(V3473)
reg314251 := __e.Call(__defun__shen_4lazyderef, reg314250, V3704)
V3480 := reg314251
_ = V3480
reg314252 := PrimIsPair(V3480)
var reg314426 Obj
if reg314252 == True {
reg314253 := PrimHead(V3480)
B := reg314253
_ = B
reg314254 := PrimTail(V3480)
reg314255 := __e.Call(__defun__shen_4lazyderef, reg314254, V3704)
V3481 := reg314255
_ = V3481
reg314256 := Nil;
reg314257 := PrimEqual(reg314256, V3481)
var reg314366 Obj
if reg314257 == True {
reg314258 := PrimTail(V3471)
reg314259 := __e.Call(__defun__shen_4lazyderef, reg314258, V3704)
V3482 := reg314259
_ = V3482
reg314260 := Nil;
reg314261 := PrimEqual(reg314260, V3482)
var reg314308 Obj
if reg314261 == True {
reg314262 := PrimTail(V3462)
Hyp := reg314262
_ = Hyp
reg314263 := __e.Call(__defun__shen_4incinfs)
_ = reg314263
reg314264 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314265 := MakeSymbol(":")
reg314266 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314267 := Nil;
reg314268 := PrimCons(reg314266, reg314267)
reg314269 := PrimCons(reg314265, reg314268)
reg314270 := PrimCons(reg314264, reg314269)
reg314271 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314272 := MakeSymbol(":")
reg314273 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314274 := Nil;
reg314275 := PrimCons(reg314273, reg314274)
reg314276 := PrimCons(reg314272, reg314275)
reg314277 := PrimCons(reg314271, reg314276)
reg314278 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314279 := PrimCons(reg314277, reg314278)
reg314280 := PrimCons(reg314270, reg314279)
reg314281 := __e.Call(__defun__bind, V3703, reg314280, V3704, V3705)
reg314308 = reg314281
} else {
reg314282 := __e.Call(__defun__shen_4pvar_2, V3482)
var reg314307 Obj
if reg314282 == True {
reg314283 := Nil;
reg314284 := __e.Call(__defun__shen_4bindv, V3482, reg314283, V3704)
_ = reg314284
reg314285 := PrimTail(V3462)
Hyp := reg314285
_ = Hyp
reg314286 := __e.Call(__defun__shen_4incinfs)
_ = reg314286
reg314287 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314288 := MakeSymbol(":")
reg314289 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314290 := Nil;
reg314291 := PrimCons(reg314289, reg314290)
reg314292 := PrimCons(reg314288, reg314291)
reg314293 := PrimCons(reg314287, reg314292)
reg314294 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314295 := MakeSymbol(":")
reg314296 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314297 := Nil;
reg314298 := PrimCons(reg314296, reg314297)
reg314299 := PrimCons(reg314295, reg314298)
reg314300 := PrimCons(reg314294, reg314299)
reg314301 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314302 := PrimCons(reg314300, reg314301)
reg314303 := PrimCons(reg314293, reg314302)
reg314304 := __e.Call(__defun__bind, V3703, reg314303, V3704, V3705)
Result := reg314304
_ = Result
reg314305 := __e.Call(__defun__shen_4unbindv, V3482, V3704)
_ = reg314305
reg314307 = Result
} else {
reg314306 := False;
reg314307 = reg314306
}
reg314308 = reg314307
}
reg314366 = reg314308
} else {
reg314309 := __e.Call(__defun__shen_4pvar_2, V3481)
var reg314365 Obj
if reg314309 == True {
reg314310 := Nil;
reg314311 := __e.Call(__defun__shen_4bindv, V3481, reg314310, V3704)
_ = reg314311
reg314312 := PrimTail(V3471)
reg314313 := __e.Call(__defun__shen_4lazyderef, reg314312, V3704)
V3483 := reg314313
_ = V3483
reg314314 := Nil;
reg314315 := PrimEqual(reg314314, V3483)
var reg314362 Obj
if reg314315 == True {
reg314316 := PrimTail(V3462)
Hyp := reg314316
_ = Hyp
reg314317 := __e.Call(__defun__shen_4incinfs)
_ = reg314317
reg314318 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314319 := MakeSymbol(":")
reg314320 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314321 := Nil;
reg314322 := PrimCons(reg314320, reg314321)
reg314323 := PrimCons(reg314319, reg314322)
reg314324 := PrimCons(reg314318, reg314323)
reg314325 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314326 := MakeSymbol(":")
reg314327 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314328 := Nil;
reg314329 := PrimCons(reg314327, reg314328)
reg314330 := PrimCons(reg314326, reg314329)
reg314331 := PrimCons(reg314325, reg314330)
reg314332 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314333 := PrimCons(reg314331, reg314332)
reg314334 := PrimCons(reg314324, reg314333)
reg314335 := __e.Call(__defun__bind, V3703, reg314334, V3704, V3705)
reg314362 = reg314335
} else {
reg314336 := __e.Call(__defun__shen_4pvar_2, V3483)
var reg314361 Obj
if reg314336 == True {
reg314337 := Nil;
reg314338 := __e.Call(__defun__shen_4bindv, V3483, reg314337, V3704)
_ = reg314338
reg314339 := PrimTail(V3462)
Hyp := reg314339
_ = Hyp
reg314340 := __e.Call(__defun__shen_4incinfs)
_ = reg314340
reg314341 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314342 := MakeSymbol(":")
reg314343 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314344 := Nil;
reg314345 := PrimCons(reg314343, reg314344)
reg314346 := PrimCons(reg314342, reg314345)
reg314347 := PrimCons(reg314341, reg314346)
reg314348 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314349 := MakeSymbol(":")
reg314350 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314351 := Nil;
reg314352 := PrimCons(reg314350, reg314351)
reg314353 := PrimCons(reg314349, reg314352)
reg314354 := PrimCons(reg314348, reg314353)
reg314355 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314356 := PrimCons(reg314354, reg314355)
reg314357 := PrimCons(reg314347, reg314356)
reg314358 := __e.Call(__defun__bind, V3703, reg314357, V3704, V3705)
Result := reg314358
_ = Result
reg314359 := __e.Call(__defun__shen_4unbindv, V3483, V3704)
_ = reg314359
reg314361 = Result
} else {
reg314360 := False;
reg314361 = reg314360
}
reg314362 = reg314361
}
Result := reg314362
_ = Result
reg314363 := __e.Call(__defun__shen_4unbindv, V3481, V3704)
_ = reg314363
reg314365 = Result
} else {
reg314364 := False;
reg314365 = reg314364
}
reg314366 = reg314365
}
reg314426 = reg314366
} else {
reg314367 := __e.Call(__defun__shen_4pvar_2, V3480)
var reg314425 Obj
if reg314367 == True {
reg314368 := __e.Call(__defun__shen_4newpv, V3704)
B := reg314368
_ = B
reg314369 := Nil;
reg314370 := PrimCons(B, reg314369)
reg314371 := __e.Call(__defun__shen_4bindv, V3480, reg314370, V3704)
_ = reg314371
reg314372 := PrimTail(V3471)
reg314373 := __e.Call(__defun__shen_4lazyderef, reg314372, V3704)
V3484 := reg314373
_ = V3484
reg314374 := Nil;
reg314375 := PrimEqual(reg314374, V3484)
var reg314422 Obj
if reg314375 == True {
reg314376 := PrimTail(V3462)
Hyp := reg314376
_ = Hyp
reg314377 := __e.Call(__defun__shen_4incinfs)
_ = reg314377
reg314378 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314379 := MakeSymbol(":")
reg314380 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314381 := Nil;
reg314382 := PrimCons(reg314380, reg314381)
reg314383 := PrimCons(reg314379, reg314382)
reg314384 := PrimCons(reg314378, reg314383)
reg314385 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314386 := MakeSymbol(":")
reg314387 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314388 := Nil;
reg314389 := PrimCons(reg314387, reg314388)
reg314390 := PrimCons(reg314386, reg314389)
reg314391 := PrimCons(reg314385, reg314390)
reg314392 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314393 := PrimCons(reg314391, reg314392)
reg314394 := PrimCons(reg314384, reg314393)
reg314395 := __e.Call(__defun__bind, V3703, reg314394, V3704, V3705)
reg314422 = reg314395
} else {
reg314396 := __e.Call(__defun__shen_4pvar_2, V3484)
var reg314421 Obj
if reg314396 == True {
reg314397 := Nil;
reg314398 := __e.Call(__defun__shen_4bindv, V3484, reg314397, V3704)
_ = reg314398
reg314399 := PrimTail(V3462)
Hyp := reg314399
_ = Hyp
reg314400 := __e.Call(__defun__shen_4incinfs)
_ = reg314400
reg314401 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314402 := MakeSymbol(":")
reg314403 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314404 := Nil;
reg314405 := PrimCons(reg314403, reg314404)
reg314406 := PrimCons(reg314402, reg314405)
reg314407 := PrimCons(reg314401, reg314406)
reg314408 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314409 := MakeSymbol(":")
reg314410 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314411 := Nil;
reg314412 := PrimCons(reg314410, reg314411)
reg314413 := PrimCons(reg314409, reg314412)
reg314414 := PrimCons(reg314408, reg314413)
reg314415 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314416 := PrimCons(reg314414, reg314415)
reg314417 := PrimCons(reg314407, reg314416)
reg314418 := __e.Call(__defun__bind, V3703, reg314417, V3704, V3705)
Result := reg314418
_ = Result
reg314419 := __e.Call(__defun__shen_4unbindv, V3484, V3704)
_ = reg314419
reg314421 = Result
} else {
reg314420 := False;
reg314421 = reg314420
}
reg314422 = reg314421
}
Result := reg314422
_ = Result
reg314423 := __e.Call(__defun__shen_4unbindv, V3480, V3704)
_ = reg314423
reg314425 = Result
} else {
reg314424 := False;
reg314425 = reg314424
}
reg314426 = reg314425
}
Result := reg314426
_ = Result
reg314427 := __e.Call(__defun__shen_4unbindv, V3474, V3704)
_ = reg314427
reg314429 = Result
} else {
reg314428 := False;
reg314429 = reg314428
}
reg314430 = reg314429
}
reg314492 = reg314430
} else {
reg314431 := __e.Call(__defun__shen_4pvar_2, V3473)
var reg314491 Obj
if reg314431 == True {
reg314432 := __e.Call(__defun__shen_4newpv, V3704)
B := reg314432
_ = B
reg314433 := MakeSymbol("*")
reg314434 := Nil;
reg314435 := PrimCons(B, reg314434)
reg314436 := PrimCons(reg314433, reg314435)
reg314437 := __e.Call(__defun__shen_4bindv, V3473, reg314436, V3704)
_ = reg314437
reg314438 := PrimTail(V3471)
reg314439 := __e.Call(__defun__shen_4lazyderef, reg314438, V3704)
V3485 := reg314439
_ = V3485
reg314440 := Nil;
reg314441 := PrimEqual(reg314440, V3485)
var reg314488 Obj
if reg314441 == True {
reg314442 := PrimTail(V3462)
Hyp := reg314442
_ = Hyp
reg314443 := __e.Call(__defun__shen_4incinfs)
_ = reg314443
reg314444 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314445 := MakeSymbol(":")
reg314446 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314447 := Nil;
reg314448 := PrimCons(reg314446, reg314447)
reg314449 := PrimCons(reg314445, reg314448)
reg314450 := PrimCons(reg314444, reg314449)
reg314451 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314452 := MakeSymbol(":")
reg314453 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314454 := Nil;
reg314455 := PrimCons(reg314453, reg314454)
reg314456 := PrimCons(reg314452, reg314455)
reg314457 := PrimCons(reg314451, reg314456)
reg314458 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314459 := PrimCons(reg314457, reg314458)
reg314460 := PrimCons(reg314450, reg314459)
reg314461 := __e.Call(__defun__bind, V3703, reg314460, V3704, V3705)
reg314488 = reg314461
} else {
reg314462 := __e.Call(__defun__shen_4pvar_2, V3485)
var reg314487 Obj
if reg314462 == True {
reg314463 := Nil;
reg314464 := __e.Call(__defun__shen_4bindv, V3485, reg314463, V3704)
_ = reg314464
reg314465 := PrimTail(V3462)
Hyp := reg314465
_ = Hyp
reg314466 := __e.Call(__defun__shen_4incinfs)
_ = reg314466
reg314467 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314468 := MakeSymbol(":")
reg314469 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314470 := Nil;
reg314471 := PrimCons(reg314469, reg314470)
reg314472 := PrimCons(reg314468, reg314471)
reg314473 := PrimCons(reg314467, reg314472)
reg314474 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314475 := MakeSymbol(":")
reg314476 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314477 := Nil;
reg314478 := PrimCons(reg314476, reg314477)
reg314479 := PrimCons(reg314475, reg314478)
reg314480 := PrimCons(reg314474, reg314479)
reg314481 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314482 := PrimCons(reg314480, reg314481)
reg314483 := PrimCons(reg314473, reg314482)
reg314484 := __e.Call(__defun__bind, V3703, reg314483, V3704, V3705)
Result := reg314484
_ = Result
reg314485 := __e.Call(__defun__shen_4unbindv, V3485, V3704)
_ = reg314485
reg314487 = Result
} else {
reg314486 := False;
reg314487 = reg314486
}
reg314488 = reg314487
}
Result := reg314488
_ = Result
reg314489 := __e.Call(__defun__shen_4unbindv, V3473, V3704)
_ = reg314489
reg314491 = Result
} else {
reg314490 := False;
reg314491 = reg314490
}
reg314492 = reg314491
}
reg314556 = reg314492
} else {
reg314493 := __e.Call(__defun__shen_4pvar_2, V3472)
var reg314555 Obj
if reg314493 == True {
reg314494 := __e.Call(__defun__shen_4newpv, V3704)
A := reg314494
_ = A
reg314495 := __e.Call(__defun__shen_4newpv, V3704)
B := reg314495
_ = B
reg314496 := MakeSymbol("*")
reg314497 := Nil;
reg314498 := PrimCons(B, reg314497)
reg314499 := PrimCons(reg314496, reg314498)
reg314500 := PrimCons(A, reg314499)
reg314501 := __e.Call(__defun__shen_4bindv, V3472, reg314500, V3704)
_ = reg314501
reg314502 := PrimTail(V3471)
reg314503 := __e.Call(__defun__shen_4lazyderef, reg314502, V3704)
V3486 := reg314503
_ = V3486
reg314504 := Nil;
reg314505 := PrimEqual(reg314504, V3486)
var reg314552 Obj
if reg314505 == True {
reg314506 := PrimTail(V3462)
Hyp := reg314506
_ = Hyp
reg314507 := __e.Call(__defun__shen_4incinfs)
_ = reg314507
reg314508 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314509 := MakeSymbol(":")
reg314510 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314511 := Nil;
reg314512 := PrimCons(reg314510, reg314511)
reg314513 := PrimCons(reg314509, reg314512)
reg314514 := PrimCons(reg314508, reg314513)
reg314515 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314516 := MakeSymbol(":")
reg314517 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314518 := Nil;
reg314519 := PrimCons(reg314517, reg314518)
reg314520 := PrimCons(reg314516, reg314519)
reg314521 := PrimCons(reg314515, reg314520)
reg314522 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314523 := PrimCons(reg314521, reg314522)
reg314524 := PrimCons(reg314514, reg314523)
reg314525 := __e.Call(__defun__bind, V3703, reg314524, V3704, V3705)
reg314552 = reg314525
} else {
reg314526 := __e.Call(__defun__shen_4pvar_2, V3486)
var reg314551 Obj
if reg314526 == True {
reg314527 := Nil;
reg314528 := __e.Call(__defun__shen_4bindv, V3486, reg314527, V3704)
_ = reg314528
reg314529 := PrimTail(V3462)
Hyp := reg314529
_ = Hyp
reg314530 := __e.Call(__defun__shen_4incinfs)
_ = reg314530
reg314531 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314532 := MakeSymbol(":")
reg314533 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314534 := Nil;
reg314535 := PrimCons(reg314533, reg314534)
reg314536 := PrimCons(reg314532, reg314535)
reg314537 := PrimCons(reg314531, reg314536)
reg314538 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314539 := MakeSymbol(":")
reg314540 := __e.Call(__defun__shen_4lazyderef, B, V3704)
reg314541 := Nil;
reg314542 := PrimCons(reg314540, reg314541)
reg314543 := PrimCons(reg314539, reg314542)
reg314544 := PrimCons(reg314538, reg314543)
reg314545 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314546 := PrimCons(reg314544, reg314545)
reg314547 := PrimCons(reg314537, reg314546)
reg314548 := __e.Call(__defun__bind, V3703, reg314547, V3704, V3705)
Result := reg314548
_ = Result
reg314549 := __e.Call(__defun__shen_4unbindv, V3486, V3704)
_ = reg314549
reg314551 = Result
} else {
reg314550 := False;
reg314551 = reg314550
}
reg314552 = reg314551
}
Result := reg314552
_ = Result
reg314553 := __e.Call(__defun__shen_4unbindv, V3472, V3704)
_ = reg314553
reg314555 = Result
} else {
reg314554 := False;
reg314555 = reg314554
}
reg314556 = reg314555
}
reg314558 = reg314556
} else {
reg314557 := False;
reg314558 = reg314557
}
reg314560 = reg314558
} else {
reg314559 := False;
reg314560 = reg314559
}
reg314562 = reg314560
} else {
reg314561 := False;
reg314562 = reg314561
}
reg314564 = reg314562
} else {
reg314563 := False;
reg314564 = reg314563
}
reg314566 = reg314564
} else {
reg314565 := False;
reg314566 = reg314565
}
reg314568 = reg314566
} else {
reg314567 := False;
reg314568 = reg314567
}
reg314570 = reg314568
} else {
reg314569 := False;
reg314570 = reg314569
}
reg314572 = reg314570
} else {
reg314571 := False;
reg314572 = reg314571
}
reg314574 = reg314572
} else {
reg314573 := False;
reg314574 = reg314573
}
reg314576 = reg314574
} else {
reg314575 := False;
reg314576 = reg314575
}
Case := reg314576
_ = Case
reg314577 := False;
reg314578 := PrimEqual(Case, reg314577)
if reg314578 == True {
reg314579 := __e.Call(__defun__shen_4lazyderef, V3702, V3704)
V3487 := reg314579
_ = V3487
reg314580 := PrimIsPair(V3487)
var reg315118 Obj
if reg314580 == True {
reg314581 := PrimHead(V3487)
reg314582 := __e.Call(__defun__shen_4lazyderef, reg314581, V3704)
V3488 := reg314582
_ = V3488
reg314583 := PrimIsPair(V3488)
var reg315116 Obj
if reg314583 == True {
reg314584 := PrimHead(V3488)
reg314585 := __e.Call(__defun__shen_4lazyderef, reg314584, V3704)
V3489 := reg314585
_ = V3489
reg314586 := PrimIsPair(V3489)
var reg315114 Obj
if reg314586 == True {
reg314587 := PrimHead(V3489)
reg314588 := __e.Call(__defun__shen_4lazyderef, reg314587, V3704)
V3490 := reg314588
_ = V3490
reg314589 := MakeSymbol("@v")
reg314590 := PrimEqual(reg314589, V3490)
var reg315112 Obj
if reg314590 == True {
reg314591 := PrimTail(V3489)
reg314592 := __e.Call(__defun__shen_4lazyderef, reg314591, V3704)
V3491 := reg314592
_ = V3491
reg314593 := PrimIsPair(V3491)
var reg315110 Obj
if reg314593 == True {
reg314594 := PrimHead(V3491)
X := reg314594
_ = X
reg314595 := PrimTail(V3491)
reg314596 := __e.Call(__defun__shen_4lazyderef, reg314595, V3704)
V3492 := reg314596
_ = V3492
reg314597 := PrimIsPair(V3492)
var reg315108 Obj
if reg314597 == True {
reg314598 := PrimHead(V3492)
Y := reg314598
_ = Y
reg314599 := PrimTail(V3492)
reg314600 := __e.Call(__defun__shen_4lazyderef, reg314599, V3704)
V3493 := reg314600
_ = V3493
reg314601 := Nil;
reg314602 := PrimEqual(reg314601, V3493)
var reg315106 Obj
if reg314602 == True {
reg314603 := PrimTail(V3488)
reg314604 := __e.Call(__defun__shen_4lazyderef, reg314603, V3704)
V3494 := reg314604
_ = V3494
reg314605 := PrimIsPair(V3494)
var reg315104 Obj
if reg314605 == True {
reg314606 := PrimHead(V3494)
reg314607 := __e.Call(__defun__shen_4lazyderef, reg314606, V3704)
V3495 := reg314607
_ = V3495
reg314608 := MakeSymbol(":")
reg314609 := PrimEqual(reg314608, V3495)
var reg315102 Obj
if reg314609 == True {
reg314610 := PrimTail(V3494)
reg314611 := __e.Call(__defun__shen_4lazyderef, reg314610, V3704)
V3496 := reg314611
_ = V3496
reg314612 := PrimIsPair(V3496)
var reg315100 Obj
if reg314612 == True {
reg314613 := PrimHead(V3496)
reg314614 := __e.Call(__defun__shen_4lazyderef, reg314613, V3704)
V3497 := reg314614
_ = V3497
reg314615 := PrimIsPair(V3497)
var reg315098 Obj
if reg314615 == True {
reg314616 := PrimHead(V3497)
reg314617 := __e.Call(__defun__shen_4lazyderef, reg314616, V3704)
V3498 := reg314617
_ = V3498
reg314618 := MakeSymbol("vector")
reg314619 := PrimEqual(reg314618, V3498)
var reg315028 Obj
if reg314619 == True {
reg314620 := PrimTail(V3497)
reg314621 := __e.Call(__defun__shen_4lazyderef, reg314620, V3704)
V3499 := reg314621
_ = V3499
reg314622 := PrimIsPair(V3499)
var reg314820 Obj
if reg314622 == True {
reg314623 := PrimHead(V3499)
A := reg314623
_ = A
reg314624 := PrimTail(V3499)
reg314625 := __e.Call(__defun__shen_4lazyderef, reg314624, V3704)
V3500 := reg314625
_ = V3500
reg314626 := Nil;
reg314627 := PrimEqual(reg314626, V3500)
var reg314752 Obj
if reg314627 == True {
reg314628 := PrimTail(V3496)
reg314629 := __e.Call(__defun__shen_4lazyderef, reg314628, V3704)
V3501 := reg314629
_ = V3501
reg314630 := Nil;
reg314631 := PrimEqual(reg314630, V3501)
var reg314686 Obj
if reg314631 == True {
reg314632 := PrimTail(V3487)
Hyp := reg314632
_ = Hyp
reg314633 := __e.Call(__defun__shen_4incinfs)
_ = reg314633
reg314634 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314635 := MakeSymbol(":")
reg314636 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314637 := Nil;
reg314638 := PrimCons(reg314636, reg314637)
reg314639 := PrimCons(reg314635, reg314638)
reg314640 := PrimCons(reg314634, reg314639)
reg314641 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314642 := MakeSymbol(":")
reg314643 := MakeSymbol("vector")
reg314644 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314645 := Nil;
reg314646 := PrimCons(reg314644, reg314645)
reg314647 := PrimCons(reg314643, reg314646)
reg314648 := Nil;
reg314649 := PrimCons(reg314647, reg314648)
reg314650 := PrimCons(reg314642, reg314649)
reg314651 := PrimCons(reg314641, reg314650)
reg314652 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314653 := PrimCons(reg314651, reg314652)
reg314654 := PrimCons(reg314640, reg314653)
reg314655 := __e.Call(__defun__bind, V3703, reg314654, V3704, V3705)
reg314686 = reg314655
} else {
reg314656 := __e.Call(__defun__shen_4pvar_2, V3501)
var reg314685 Obj
if reg314656 == True {
reg314657 := Nil;
reg314658 := __e.Call(__defun__shen_4bindv, V3501, reg314657, V3704)
_ = reg314658
reg314659 := PrimTail(V3487)
Hyp := reg314659
_ = Hyp
reg314660 := __e.Call(__defun__shen_4incinfs)
_ = reg314660
reg314661 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314662 := MakeSymbol(":")
reg314663 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314664 := Nil;
reg314665 := PrimCons(reg314663, reg314664)
reg314666 := PrimCons(reg314662, reg314665)
reg314667 := PrimCons(reg314661, reg314666)
reg314668 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314669 := MakeSymbol(":")
reg314670 := MakeSymbol("vector")
reg314671 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314672 := Nil;
reg314673 := PrimCons(reg314671, reg314672)
reg314674 := PrimCons(reg314670, reg314673)
reg314675 := Nil;
reg314676 := PrimCons(reg314674, reg314675)
reg314677 := PrimCons(reg314669, reg314676)
reg314678 := PrimCons(reg314668, reg314677)
reg314679 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314680 := PrimCons(reg314678, reg314679)
reg314681 := PrimCons(reg314667, reg314680)
reg314682 := __e.Call(__defun__bind, V3703, reg314681, V3704, V3705)
Result := reg314682
_ = Result
reg314683 := __e.Call(__defun__shen_4unbindv, V3501, V3704)
_ = reg314683
reg314685 = Result
} else {
reg314684 := False;
reg314685 = reg314684
}
reg314686 = reg314685
}
reg314752 = reg314686
} else {
reg314687 := __e.Call(__defun__shen_4pvar_2, V3500)
var reg314751 Obj
if reg314687 == True {
reg314688 := Nil;
reg314689 := __e.Call(__defun__shen_4bindv, V3500, reg314688, V3704)
_ = reg314689
reg314690 := PrimTail(V3496)
reg314691 := __e.Call(__defun__shen_4lazyderef, reg314690, V3704)
V3502 := reg314691
_ = V3502
reg314692 := Nil;
reg314693 := PrimEqual(reg314692, V3502)
var reg314748 Obj
if reg314693 == True {
reg314694 := PrimTail(V3487)
Hyp := reg314694
_ = Hyp
reg314695 := __e.Call(__defun__shen_4incinfs)
_ = reg314695
reg314696 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314697 := MakeSymbol(":")
reg314698 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314699 := Nil;
reg314700 := PrimCons(reg314698, reg314699)
reg314701 := PrimCons(reg314697, reg314700)
reg314702 := PrimCons(reg314696, reg314701)
reg314703 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314704 := MakeSymbol(":")
reg314705 := MakeSymbol("vector")
reg314706 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314707 := Nil;
reg314708 := PrimCons(reg314706, reg314707)
reg314709 := PrimCons(reg314705, reg314708)
reg314710 := Nil;
reg314711 := PrimCons(reg314709, reg314710)
reg314712 := PrimCons(reg314704, reg314711)
reg314713 := PrimCons(reg314703, reg314712)
reg314714 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314715 := PrimCons(reg314713, reg314714)
reg314716 := PrimCons(reg314702, reg314715)
reg314717 := __e.Call(__defun__bind, V3703, reg314716, V3704, V3705)
reg314748 = reg314717
} else {
reg314718 := __e.Call(__defun__shen_4pvar_2, V3502)
var reg314747 Obj
if reg314718 == True {
reg314719 := Nil;
reg314720 := __e.Call(__defun__shen_4bindv, V3502, reg314719, V3704)
_ = reg314720
reg314721 := PrimTail(V3487)
Hyp := reg314721
_ = Hyp
reg314722 := __e.Call(__defun__shen_4incinfs)
_ = reg314722
reg314723 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314724 := MakeSymbol(":")
reg314725 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314726 := Nil;
reg314727 := PrimCons(reg314725, reg314726)
reg314728 := PrimCons(reg314724, reg314727)
reg314729 := PrimCons(reg314723, reg314728)
reg314730 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314731 := MakeSymbol(":")
reg314732 := MakeSymbol("vector")
reg314733 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314734 := Nil;
reg314735 := PrimCons(reg314733, reg314734)
reg314736 := PrimCons(reg314732, reg314735)
reg314737 := Nil;
reg314738 := PrimCons(reg314736, reg314737)
reg314739 := PrimCons(reg314731, reg314738)
reg314740 := PrimCons(reg314730, reg314739)
reg314741 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314742 := PrimCons(reg314740, reg314741)
reg314743 := PrimCons(reg314729, reg314742)
reg314744 := __e.Call(__defun__bind, V3703, reg314743, V3704, V3705)
Result := reg314744
_ = Result
reg314745 := __e.Call(__defun__shen_4unbindv, V3502, V3704)
_ = reg314745
reg314747 = Result
} else {
reg314746 := False;
reg314747 = reg314746
}
reg314748 = reg314747
}
Result := reg314748
_ = Result
reg314749 := __e.Call(__defun__shen_4unbindv, V3500, V3704)
_ = reg314749
reg314751 = Result
} else {
reg314750 := False;
reg314751 = reg314750
}
reg314752 = reg314751
}
reg314820 = reg314752
} else {
reg314753 := __e.Call(__defun__shen_4pvar_2, V3499)
var reg314819 Obj
if reg314753 == True {
reg314754 := __e.Call(__defun__shen_4newpv, V3704)
A := reg314754
_ = A
reg314755 := Nil;
reg314756 := PrimCons(A, reg314755)
reg314757 := __e.Call(__defun__shen_4bindv, V3499, reg314756, V3704)
_ = reg314757
reg314758 := PrimTail(V3496)
reg314759 := __e.Call(__defun__shen_4lazyderef, reg314758, V3704)
V3503 := reg314759
_ = V3503
reg314760 := Nil;
reg314761 := PrimEqual(reg314760, V3503)
var reg314816 Obj
if reg314761 == True {
reg314762 := PrimTail(V3487)
Hyp := reg314762
_ = Hyp
reg314763 := __e.Call(__defun__shen_4incinfs)
_ = reg314763
reg314764 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314765 := MakeSymbol(":")
reg314766 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314767 := Nil;
reg314768 := PrimCons(reg314766, reg314767)
reg314769 := PrimCons(reg314765, reg314768)
reg314770 := PrimCons(reg314764, reg314769)
reg314771 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314772 := MakeSymbol(":")
reg314773 := MakeSymbol("vector")
reg314774 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314775 := Nil;
reg314776 := PrimCons(reg314774, reg314775)
reg314777 := PrimCons(reg314773, reg314776)
reg314778 := Nil;
reg314779 := PrimCons(reg314777, reg314778)
reg314780 := PrimCons(reg314772, reg314779)
reg314781 := PrimCons(reg314771, reg314780)
reg314782 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314783 := PrimCons(reg314781, reg314782)
reg314784 := PrimCons(reg314770, reg314783)
reg314785 := __e.Call(__defun__bind, V3703, reg314784, V3704, V3705)
reg314816 = reg314785
} else {
reg314786 := __e.Call(__defun__shen_4pvar_2, V3503)
var reg314815 Obj
if reg314786 == True {
reg314787 := Nil;
reg314788 := __e.Call(__defun__shen_4bindv, V3503, reg314787, V3704)
_ = reg314788
reg314789 := PrimTail(V3487)
Hyp := reg314789
_ = Hyp
reg314790 := __e.Call(__defun__shen_4incinfs)
_ = reg314790
reg314791 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314792 := MakeSymbol(":")
reg314793 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314794 := Nil;
reg314795 := PrimCons(reg314793, reg314794)
reg314796 := PrimCons(reg314792, reg314795)
reg314797 := PrimCons(reg314791, reg314796)
reg314798 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314799 := MakeSymbol(":")
reg314800 := MakeSymbol("vector")
reg314801 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314802 := Nil;
reg314803 := PrimCons(reg314801, reg314802)
reg314804 := PrimCons(reg314800, reg314803)
reg314805 := Nil;
reg314806 := PrimCons(reg314804, reg314805)
reg314807 := PrimCons(reg314799, reg314806)
reg314808 := PrimCons(reg314798, reg314807)
reg314809 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314810 := PrimCons(reg314808, reg314809)
reg314811 := PrimCons(reg314797, reg314810)
reg314812 := __e.Call(__defun__bind, V3703, reg314811, V3704, V3705)
Result := reg314812
_ = Result
reg314813 := __e.Call(__defun__shen_4unbindv, V3503, V3704)
_ = reg314813
reg314815 = Result
} else {
reg314814 := False;
reg314815 = reg314814
}
reg314816 = reg314815
}
Result := reg314816
_ = Result
reg314817 := __e.Call(__defun__shen_4unbindv, V3499, V3704)
_ = reg314817
reg314819 = Result
} else {
reg314818 := False;
reg314819 = reg314818
}
reg314820 = reg314819
}
reg315028 = reg314820
} else {
reg314821 := __e.Call(__defun__shen_4pvar_2, V3498)
var reg315027 Obj
if reg314821 == True {
reg314822 := MakeSymbol("vector")
reg314823 := __e.Call(__defun__shen_4bindv, V3498, reg314822, V3704)
_ = reg314823
reg314824 := PrimTail(V3497)
reg314825 := __e.Call(__defun__shen_4lazyderef, reg314824, V3704)
V3504 := reg314825
_ = V3504
reg314826 := PrimIsPair(V3504)
var reg315024 Obj
if reg314826 == True {
reg314827 := PrimHead(V3504)
A := reg314827
_ = A
reg314828 := PrimTail(V3504)
reg314829 := __e.Call(__defun__shen_4lazyderef, reg314828, V3704)
V3505 := reg314829
_ = V3505
reg314830 := Nil;
reg314831 := PrimEqual(reg314830, V3505)
var reg314956 Obj
if reg314831 == True {
reg314832 := PrimTail(V3496)
reg314833 := __e.Call(__defun__shen_4lazyderef, reg314832, V3704)
V3506 := reg314833
_ = V3506
reg314834 := Nil;
reg314835 := PrimEqual(reg314834, V3506)
var reg314890 Obj
if reg314835 == True {
reg314836 := PrimTail(V3487)
Hyp := reg314836
_ = Hyp
reg314837 := __e.Call(__defun__shen_4incinfs)
_ = reg314837
reg314838 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314839 := MakeSymbol(":")
reg314840 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314841 := Nil;
reg314842 := PrimCons(reg314840, reg314841)
reg314843 := PrimCons(reg314839, reg314842)
reg314844 := PrimCons(reg314838, reg314843)
reg314845 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314846 := MakeSymbol(":")
reg314847 := MakeSymbol("vector")
reg314848 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314849 := Nil;
reg314850 := PrimCons(reg314848, reg314849)
reg314851 := PrimCons(reg314847, reg314850)
reg314852 := Nil;
reg314853 := PrimCons(reg314851, reg314852)
reg314854 := PrimCons(reg314846, reg314853)
reg314855 := PrimCons(reg314845, reg314854)
reg314856 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314857 := PrimCons(reg314855, reg314856)
reg314858 := PrimCons(reg314844, reg314857)
reg314859 := __e.Call(__defun__bind, V3703, reg314858, V3704, V3705)
reg314890 = reg314859
} else {
reg314860 := __e.Call(__defun__shen_4pvar_2, V3506)
var reg314889 Obj
if reg314860 == True {
reg314861 := Nil;
reg314862 := __e.Call(__defun__shen_4bindv, V3506, reg314861, V3704)
_ = reg314862
reg314863 := PrimTail(V3487)
Hyp := reg314863
_ = Hyp
reg314864 := __e.Call(__defun__shen_4incinfs)
_ = reg314864
reg314865 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314866 := MakeSymbol(":")
reg314867 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314868 := Nil;
reg314869 := PrimCons(reg314867, reg314868)
reg314870 := PrimCons(reg314866, reg314869)
reg314871 := PrimCons(reg314865, reg314870)
reg314872 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314873 := MakeSymbol(":")
reg314874 := MakeSymbol("vector")
reg314875 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314876 := Nil;
reg314877 := PrimCons(reg314875, reg314876)
reg314878 := PrimCons(reg314874, reg314877)
reg314879 := Nil;
reg314880 := PrimCons(reg314878, reg314879)
reg314881 := PrimCons(reg314873, reg314880)
reg314882 := PrimCons(reg314872, reg314881)
reg314883 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314884 := PrimCons(reg314882, reg314883)
reg314885 := PrimCons(reg314871, reg314884)
reg314886 := __e.Call(__defun__bind, V3703, reg314885, V3704, V3705)
Result := reg314886
_ = Result
reg314887 := __e.Call(__defun__shen_4unbindv, V3506, V3704)
_ = reg314887
reg314889 = Result
} else {
reg314888 := False;
reg314889 = reg314888
}
reg314890 = reg314889
}
reg314956 = reg314890
} else {
reg314891 := __e.Call(__defun__shen_4pvar_2, V3505)
var reg314955 Obj
if reg314891 == True {
reg314892 := Nil;
reg314893 := __e.Call(__defun__shen_4bindv, V3505, reg314892, V3704)
_ = reg314893
reg314894 := PrimTail(V3496)
reg314895 := __e.Call(__defun__shen_4lazyderef, reg314894, V3704)
V3507 := reg314895
_ = V3507
reg314896 := Nil;
reg314897 := PrimEqual(reg314896, V3507)
var reg314952 Obj
if reg314897 == True {
reg314898 := PrimTail(V3487)
Hyp := reg314898
_ = Hyp
reg314899 := __e.Call(__defun__shen_4incinfs)
_ = reg314899
reg314900 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314901 := MakeSymbol(":")
reg314902 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314903 := Nil;
reg314904 := PrimCons(reg314902, reg314903)
reg314905 := PrimCons(reg314901, reg314904)
reg314906 := PrimCons(reg314900, reg314905)
reg314907 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314908 := MakeSymbol(":")
reg314909 := MakeSymbol("vector")
reg314910 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314911 := Nil;
reg314912 := PrimCons(reg314910, reg314911)
reg314913 := PrimCons(reg314909, reg314912)
reg314914 := Nil;
reg314915 := PrimCons(reg314913, reg314914)
reg314916 := PrimCons(reg314908, reg314915)
reg314917 := PrimCons(reg314907, reg314916)
reg314918 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314919 := PrimCons(reg314917, reg314918)
reg314920 := PrimCons(reg314906, reg314919)
reg314921 := __e.Call(__defun__bind, V3703, reg314920, V3704, V3705)
reg314952 = reg314921
} else {
reg314922 := __e.Call(__defun__shen_4pvar_2, V3507)
var reg314951 Obj
if reg314922 == True {
reg314923 := Nil;
reg314924 := __e.Call(__defun__shen_4bindv, V3507, reg314923, V3704)
_ = reg314924
reg314925 := PrimTail(V3487)
Hyp := reg314925
_ = Hyp
reg314926 := __e.Call(__defun__shen_4incinfs)
_ = reg314926
reg314927 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314928 := MakeSymbol(":")
reg314929 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314930 := Nil;
reg314931 := PrimCons(reg314929, reg314930)
reg314932 := PrimCons(reg314928, reg314931)
reg314933 := PrimCons(reg314927, reg314932)
reg314934 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314935 := MakeSymbol(":")
reg314936 := MakeSymbol("vector")
reg314937 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314938 := Nil;
reg314939 := PrimCons(reg314937, reg314938)
reg314940 := PrimCons(reg314936, reg314939)
reg314941 := Nil;
reg314942 := PrimCons(reg314940, reg314941)
reg314943 := PrimCons(reg314935, reg314942)
reg314944 := PrimCons(reg314934, reg314943)
reg314945 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314946 := PrimCons(reg314944, reg314945)
reg314947 := PrimCons(reg314933, reg314946)
reg314948 := __e.Call(__defun__bind, V3703, reg314947, V3704, V3705)
Result := reg314948
_ = Result
reg314949 := __e.Call(__defun__shen_4unbindv, V3507, V3704)
_ = reg314949
reg314951 = Result
} else {
reg314950 := False;
reg314951 = reg314950
}
reg314952 = reg314951
}
Result := reg314952
_ = Result
reg314953 := __e.Call(__defun__shen_4unbindv, V3505, V3704)
_ = reg314953
reg314955 = Result
} else {
reg314954 := False;
reg314955 = reg314954
}
reg314956 = reg314955
}
reg315024 = reg314956
} else {
reg314957 := __e.Call(__defun__shen_4pvar_2, V3504)
var reg315023 Obj
if reg314957 == True {
reg314958 := __e.Call(__defun__shen_4newpv, V3704)
A := reg314958
_ = A
reg314959 := Nil;
reg314960 := PrimCons(A, reg314959)
reg314961 := __e.Call(__defun__shen_4bindv, V3504, reg314960, V3704)
_ = reg314961
reg314962 := PrimTail(V3496)
reg314963 := __e.Call(__defun__shen_4lazyderef, reg314962, V3704)
V3508 := reg314963
_ = V3508
reg314964 := Nil;
reg314965 := PrimEqual(reg314964, V3508)
var reg315020 Obj
if reg314965 == True {
reg314966 := PrimTail(V3487)
Hyp := reg314966
_ = Hyp
reg314967 := __e.Call(__defun__shen_4incinfs)
_ = reg314967
reg314968 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314969 := MakeSymbol(":")
reg314970 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314971 := Nil;
reg314972 := PrimCons(reg314970, reg314971)
reg314973 := PrimCons(reg314969, reg314972)
reg314974 := PrimCons(reg314968, reg314973)
reg314975 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg314976 := MakeSymbol(":")
reg314977 := MakeSymbol("vector")
reg314978 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314979 := Nil;
reg314980 := PrimCons(reg314978, reg314979)
reg314981 := PrimCons(reg314977, reg314980)
reg314982 := Nil;
reg314983 := PrimCons(reg314981, reg314982)
reg314984 := PrimCons(reg314976, reg314983)
reg314985 := PrimCons(reg314975, reg314984)
reg314986 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg314987 := PrimCons(reg314985, reg314986)
reg314988 := PrimCons(reg314974, reg314987)
reg314989 := __e.Call(__defun__bind, V3703, reg314988, V3704, V3705)
reg315020 = reg314989
} else {
reg314990 := __e.Call(__defun__shen_4pvar_2, V3508)
var reg315019 Obj
if reg314990 == True {
reg314991 := Nil;
reg314992 := __e.Call(__defun__shen_4bindv, V3508, reg314991, V3704)
_ = reg314992
reg314993 := PrimTail(V3487)
Hyp := reg314993
_ = Hyp
reg314994 := __e.Call(__defun__shen_4incinfs)
_ = reg314994
reg314995 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg314996 := MakeSymbol(":")
reg314997 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg314998 := Nil;
reg314999 := PrimCons(reg314997, reg314998)
reg315000 := PrimCons(reg314996, reg314999)
reg315001 := PrimCons(reg314995, reg315000)
reg315002 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315003 := MakeSymbol(":")
reg315004 := MakeSymbol("vector")
reg315005 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg315006 := Nil;
reg315007 := PrimCons(reg315005, reg315006)
reg315008 := PrimCons(reg315004, reg315007)
reg315009 := Nil;
reg315010 := PrimCons(reg315008, reg315009)
reg315011 := PrimCons(reg315003, reg315010)
reg315012 := PrimCons(reg315002, reg315011)
reg315013 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315014 := PrimCons(reg315012, reg315013)
reg315015 := PrimCons(reg315001, reg315014)
reg315016 := __e.Call(__defun__bind, V3703, reg315015, V3704, V3705)
Result := reg315016
_ = Result
reg315017 := __e.Call(__defun__shen_4unbindv, V3508, V3704)
_ = reg315017
reg315019 = Result
} else {
reg315018 := False;
reg315019 = reg315018
}
reg315020 = reg315019
}
Result := reg315020
_ = Result
reg315021 := __e.Call(__defun__shen_4unbindv, V3504, V3704)
_ = reg315021
reg315023 = Result
} else {
reg315022 := False;
reg315023 = reg315022
}
reg315024 = reg315023
}
Result := reg315024
_ = Result
reg315025 := __e.Call(__defun__shen_4unbindv, V3498, V3704)
_ = reg315025
reg315027 = Result
} else {
reg315026 := False;
reg315027 = reg315026
}
reg315028 = reg315027
}
reg315098 = reg315028
} else {
reg315029 := __e.Call(__defun__shen_4pvar_2, V3497)
var reg315097 Obj
if reg315029 == True {
reg315030 := __e.Call(__defun__shen_4newpv, V3704)
A := reg315030
_ = A
reg315031 := MakeSymbol("vector")
reg315032 := Nil;
reg315033 := PrimCons(A, reg315032)
reg315034 := PrimCons(reg315031, reg315033)
reg315035 := __e.Call(__defun__shen_4bindv, V3497, reg315034, V3704)
_ = reg315035
reg315036 := PrimTail(V3496)
reg315037 := __e.Call(__defun__shen_4lazyderef, reg315036, V3704)
V3509 := reg315037
_ = V3509
reg315038 := Nil;
reg315039 := PrimEqual(reg315038, V3509)
var reg315094 Obj
if reg315039 == True {
reg315040 := PrimTail(V3487)
Hyp := reg315040
_ = Hyp
reg315041 := __e.Call(__defun__shen_4incinfs)
_ = reg315041
reg315042 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315043 := MakeSymbol(":")
reg315044 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg315045 := Nil;
reg315046 := PrimCons(reg315044, reg315045)
reg315047 := PrimCons(reg315043, reg315046)
reg315048 := PrimCons(reg315042, reg315047)
reg315049 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315050 := MakeSymbol(":")
reg315051 := MakeSymbol("vector")
reg315052 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg315053 := Nil;
reg315054 := PrimCons(reg315052, reg315053)
reg315055 := PrimCons(reg315051, reg315054)
reg315056 := Nil;
reg315057 := PrimCons(reg315055, reg315056)
reg315058 := PrimCons(reg315050, reg315057)
reg315059 := PrimCons(reg315049, reg315058)
reg315060 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315061 := PrimCons(reg315059, reg315060)
reg315062 := PrimCons(reg315048, reg315061)
reg315063 := __e.Call(__defun__bind, V3703, reg315062, V3704, V3705)
reg315094 = reg315063
} else {
reg315064 := __e.Call(__defun__shen_4pvar_2, V3509)
var reg315093 Obj
if reg315064 == True {
reg315065 := Nil;
reg315066 := __e.Call(__defun__shen_4bindv, V3509, reg315065, V3704)
_ = reg315066
reg315067 := PrimTail(V3487)
Hyp := reg315067
_ = Hyp
reg315068 := __e.Call(__defun__shen_4incinfs)
_ = reg315068
reg315069 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315070 := MakeSymbol(":")
reg315071 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg315072 := Nil;
reg315073 := PrimCons(reg315071, reg315072)
reg315074 := PrimCons(reg315070, reg315073)
reg315075 := PrimCons(reg315069, reg315074)
reg315076 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315077 := MakeSymbol(":")
reg315078 := MakeSymbol("vector")
reg315079 := __e.Call(__defun__shen_4lazyderef, A, V3704)
reg315080 := Nil;
reg315081 := PrimCons(reg315079, reg315080)
reg315082 := PrimCons(reg315078, reg315081)
reg315083 := Nil;
reg315084 := PrimCons(reg315082, reg315083)
reg315085 := PrimCons(reg315077, reg315084)
reg315086 := PrimCons(reg315076, reg315085)
reg315087 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315088 := PrimCons(reg315086, reg315087)
reg315089 := PrimCons(reg315075, reg315088)
reg315090 := __e.Call(__defun__bind, V3703, reg315089, V3704, V3705)
Result := reg315090
_ = Result
reg315091 := __e.Call(__defun__shen_4unbindv, V3509, V3704)
_ = reg315091
reg315093 = Result
} else {
reg315092 := False;
reg315093 = reg315092
}
reg315094 = reg315093
}
Result := reg315094
_ = Result
reg315095 := __e.Call(__defun__shen_4unbindv, V3497, V3704)
_ = reg315095
reg315097 = Result
} else {
reg315096 := False;
reg315097 = reg315096
}
reg315098 = reg315097
}
reg315100 = reg315098
} else {
reg315099 := False;
reg315100 = reg315099
}
reg315102 = reg315100
} else {
reg315101 := False;
reg315102 = reg315101
}
reg315104 = reg315102
} else {
reg315103 := False;
reg315104 = reg315103
}
reg315106 = reg315104
} else {
reg315105 := False;
reg315106 = reg315105
}
reg315108 = reg315106
} else {
reg315107 := False;
reg315108 = reg315107
}
reg315110 = reg315108
} else {
reg315109 := False;
reg315110 = reg315109
}
reg315112 = reg315110
} else {
reg315111 := False;
reg315112 = reg315111
}
reg315114 = reg315112
} else {
reg315113 := False;
reg315114 = reg315113
}
reg315116 = reg315114
} else {
reg315115 := False;
reg315116 = reg315115
}
reg315118 = reg315116
} else {
reg315117 := False;
reg315118 = reg315117
}
Case := reg315118
_ = Case
reg315119 := False;
reg315120 := PrimEqual(Case, reg315119)
if reg315120 == True {
reg315121 := __e.Call(__defun__shen_4lazyderef, V3702, V3704)
V3510 := reg315121
_ = V3510
reg315122 := PrimIsPair(V3510)
var reg315287 Obj
if reg315122 == True {
reg315123 := PrimHead(V3510)
reg315124 := __e.Call(__defun__shen_4lazyderef, reg315123, V3704)
V3511 := reg315124
_ = V3511
reg315125 := PrimIsPair(V3511)
var reg315285 Obj
if reg315125 == True {
reg315126 := PrimHead(V3511)
reg315127 := __e.Call(__defun__shen_4lazyderef, reg315126, V3704)
V3512 := reg315127
_ = V3512
reg315128 := PrimIsPair(V3512)
var reg315283 Obj
if reg315128 == True {
reg315129 := PrimHead(V3512)
reg315130 := __e.Call(__defun__shen_4lazyderef, reg315129, V3704)
V3513 := reg315130
_ = V3513
reg315131 := MakeSymbol("@s")
reg315132 := PrimEqual(reg315131, V3513)
var reg315281 Obj
if reg315132 == True {
reg315133 := PrimTail(V3512)
reg315134 := __e.Call(__defun__shen_4lazyderef, reg315133, V3704)
V3514 := reg315134
_ = V3514
reg315135 := PrimIsPair(V3514)
var reg315279 Obj
if reg315135 == True {
reg315136 := PrimHead(V3514)
X := reg315136
_ = X
reg315137 := PrimTail(V3514)
reg315138 := __e.Call(__defun__shen_4lazyderef, reg315137, V3704)
V3515 := reg315138
_ = V3515
reg315139 := PrimIsPair(V3515)
var reg315277 Obj
if reg315139 == True {
reg315140 := PrimHead(V3515)
Y := reg315140
_ = Y
reg315141 := PrimTail(V3515)
reg315142 := __e.Call(__defun__shen_4lazyderef, reg315141, V3704)
V3516 := reg315142
_ = V3516
reg315143 := Nil;
reg315144 := PrimEqual(reg315143, V3516)
var reg315275 Obj
if reg315144 == True {
reg315145 := PrimTail(V3511)
reg315146 := __e.Call(__defun__shen_4lazyderef, reg315145, V3704)
V3517 := reg315146
_ = V3517
reg315147 := PrimIsPair(V3517)
var reg315273 Obj
if reg315147 == True {
reg315148 := PrimHead(V3517)
reg315149 := __e.Call(__defun__shen_4lazyderef, reg315148, V3704)
V3518 := reg315149
_ = V3518
reg315150 := MakeSymbol(":")
reg315151 := PrimEqual(reg315150, V3518)
var reg315271 Obj
if reg315151 == True {
reg315152 := PrimTail(V3517)
reg315153 := __e.Call(__defun__shen_4lazyderef, reg315152, V3704)
V3519 := reg315153
_ = V3519
reg315154 := PrimIsPair(V3519)
var reg315269 Obj
if reg315154 == True {
reg315155 := PrimHead(V3519)
reg315156 := __e.Call(__defun__shen_4lazyderef, reg315155, V3704)
V3520 := reg315156
_ = V3520
reg315157 := MakeSymbol("string")
reg315158 := PrimEqual(reg315157, V3520)
var reg315267 Obj
if reg315158 == True {
reg315159 := PrimTail(V3519)
reg315160 := __e.Call(__defun__shen_4lazyderef, reg315159, V3704)
V3521 := reg315160
_ = V3521
reg315161 := Nil;
reg315162 := PrimEqual(reg315161, V3521)
var reg315209 Obj
if reg315162 == True {
reg315163 := PrimTail(V3510)
Hyp := reg315163
_ = Hyp
reg315164 := __e.Call(__defun__shen_4incinfs)
_ = reg315164
reg315165 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315166 := MakeSymbol(":")
reg315167 := MakeSymbol("string")
reg315168 := Nil;
reg315169 := PrimCons(reg315167, reg315168)
reg315170 := PrimCons(reg315166, reg315169)
reg315171 := PrimCons(reg315165, reg315170)
reg315172 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315173 := MakeSymbol(":")
reg315174 := MakeSymbol("string")
reg315175 := Nil;
reg315176 := PrimCons(reg315174, reg315175)
reg315177 := PrimCons(reg315173, reg315176)
reg315178 := PrimCons(reg315172, reg315177)
reg315179 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315180 := PrimCons(reg315178, reg315179)
reg315181 := PrimCons(reg315171, reg315180)
reg315182 := __e.Call(__defun__bind, V3703, reg315181, V3704, V3705)
reg315209 = reg315182
} else {
reg315183 := __e.Call(__defun__shen_4pvar_2, V3521)
var reg315208 Obj
if reg315183 == True {
reg315184 := Nil;
reg315185 := __e.Call(__defun__shen_4bindv, V3521, reg315184, V3704)
_ = reg315185
reg315186 := PrimTail(V3510)
Hyp := reg315186
_ = Hyp
reg315187 := __e.Call(__defun__shen_4incinfs)
_ = reg315187
reg315188 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315189 := MakeSymbol(":")
reg315190 := MakeSymbol("string")
reg315191 := Nil;
reg315192 := PrimCons(reg315190, reg315191)
reg315193 := PrimCons(reg315189, reg315192)
reg315194 := PrimCons(reg315188, reg315193)
reg315195 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315196 := MakeSymbol(":")
reg315197 := MakeSymbol("string")
reg315198 := Nil;
reg315199 := PrimCons(reg315197, reg315198)
reg315200 := PrimCons(reg315196, reg315199)
reg315201 := PrimCons(reg315195, reg315200)
reg315202 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315203 := PrimCons(reg315201, reg315202)
reg315204 := PrimCons(reg315194, reg315203)
reg315205 := __e.Call(__defun__bind, V3703, reg315204, V3704, V3705)
Result := reg315205
_ = Result
reg315206 := __e.Call(__defun__shen_4unbindv, V3521, V3704)
_ = reg315206
reg315208 = Result
} else {
reg315207 := False;
reg315208 = reg315207
}
reg315209 = reg315208
}
reg315267 = reg315209
} else {
reg315210 := __e.Call(__defun__shen_4pvar_2, V3520)
var reg315266 Obj
if reg315210 == True {
reg315211 := MakeSymbol("string")
reg315212 := __e.Call(__defun__shen_4bindv, V3520, reg315211, V3704)
_ = reg315212
reg315213 := PrimTail(V3519)
reg315214 := __e.Call(__defun__shen_4lazyderef, reg315213, V3704)
V3522 := reg315214
_ = V3522
reg315215 := Nil;
reg315216 := PrimEqual(reg315215, V3522)
var reg315263 Obj
if reg315216 == True {
reg315217 := PrimTail(V3510)
Hyp := reg315217
_ = Hyp
reg315218 := __e.Call(__defun__shen_4incinfs)
_ = reg315218
reg315219 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315220 := MakeSymbol(":")
reg315221 := MakeSymbol("string")
reg315222 := Nil;
reg315223 := PrimCons(reg315221, reg315222)
reg315224 := PrimCons(reg315220, reg315223)
reg315225 := PrimCons(reg315219, reg315224)
reg315226 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315227 := MakeSymbol(":")
reg315228 := MakeSymbol("string")
reg315229 := Nil;
reg315230 := PrimCons(reg315228, reg315229)
reg315231 := PrimCons(reg315227, reg315230)
reg315232 := PrimCons(reg315226, reg315231)
reg315233 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315234 := PrimCons(reg315232, reg315233)
reg315235 := PrimCons(reg315225, reg315234)
reg315236 := __e.Call(__defun__bind, V3703, reg315235, V3704, V3705)
reg315263 = reg315236
} else {
reg315237 := __e.Call(__defun__shen_4pvar_2, V3522)
var reg315262 Obj
if reg315237 == True {
reg315238 := Nil;
reg315239 := __e.Call(__defun__shen_4bindv, V3522, reg315238, V3704)
_ = reg315239
reg315240 := PrimTail(V3510)
Hyp := reg315240
_ = Hyp
reg315241 := __e.Call(__defun__shen_4incinfs)
_ = reg315241
reg315242 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315243 := MakeSymbol(":")
reg315244 := MakeSymbol("string")
reg315245 := Nil;
reg315246 := PrimCons(reg315244, reg315245)
reg315247 := PrimCons(reg315243, reg315246)
reg315248 := PrimCons(reg315242, reg315247)
reg315249 := __e.Call(__defun__shen_4lazyderef, Y, V3704)
reg315250 := MakeSymbol(":")
reg315251 := MakeSymbol("string")
reg315252 := Nil;
reg315253 := PrimCons(reg315251, reg315252)
reg315254 := PrimCons(reg315250, reg315253)
reg315255 := PrimCons(reg315249, reg315254)
reg315256 := __e.Call(__defun__shen_4lazyderef, Hyp, V3704)
reg315257 := PrimCons(reg315255, reg315256)
reg315258 := PrimCons(reg315248, reg315257)
reg315259 := __e.Call(__defun__bind, V3703, reg315258, V3704, V3705)
Result := reg315259
_ = Result
reg315260 := __e.Call(__defun__shen_4unbindv, V3522, V3704)
_ = reg315260
reg315262 = Result
} else {
reg315261 := False;
reg315262 = reg315261
}
reg315263 = reg315262
}
Result := reg315263
_ = Result
reg315264 := __e.Call(__defun__shen_4unbindv, V3520, V3704)
_ = reg315264
reg315266 = Result
} else {
reg315265 := False;
reg315266 = reg315265
}
reg315267 = reg315266
}
reg315269 = reg315267
} else {
reg315268 := False;
reg315269 = reg315268
}
reg315271 = reg315269
} else {
reg315270 := False;
reg315271 = reg315270
}
reg315273 = reg315271
} else {
reg315272 := False;
reg315273 = reg315272
}
reg315275 = reg315273
} else {
reg315274 := False;
reg315275 = reg315274
}
reg315277 = reg315275
} else {
reg315276 := False;
reg315277 = reg315276
}
reg315279 = reg315277
} else {
reg315278 := False;
reg315279 = reg315278
}
reg315281 = reg315279
} else {
reg315280 := False;
reg315281 = reg315280
}
reg315283 = reg315281
} else {
reg315282 := False;
reg315283 = reg315282
}
reg315285 = reg315283
} else {
reg315284 := False;
reg315285 = reg315284
}
reg315287 = reg315285
} else {
reg315286 := False;
reg315287 = reg315286
}
Case := reg315287
_ = Case
reg315288 := False;
reg315289 := PrimEqual(Case, reg315288)
if reg315289 == True {
reg315290 := __e.Call(__defun__shen_4lazyderef, V3702, V3704)
V3523 := reg315290
_ = V3523
reg315291 := PrimIsPair(V3523)
if reg315291 == True {
reg315292 := PrimHead(V3523)
X := reg315292
_ = X
reg315293 := PrimTail(V3523)
Hyp := reg315293
_ = Hyp
reg315294 := __e.Call(__defun__shen_4newpv, V3704)
NewHyps := reg315294
_ = NewHyps
reg315295 := __e.Call(__defun__shen_4incinfs)
_ = reg315295
reg315296 := __e.Call(__defun__shen_4lazyderef, X, V3704)
reg315297 := __e.Call(__defun__shen_4lazyderef, NewHyps, V3704)
reg315298 := PrimCons(reg315296, reg315297)
reg315299 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4t_d_1hyps, Hyp, NewHyps, V3704, V3705)
return
}, 0)
__ctx.TailApply(__defun__bind, V3703, reg315298, V3704, reg315299)
return
} else {
reg315302 := False;
__ctx.Return(reg315302)
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
V3722 := __args[0]
_ = V3722
V3723 := __args[1]
_ = V3723
V3724 := __args[2]
_ = V3724
V3725 := __args[3]
_ = V3725
reg315303 := MakeSymbol("shen.*spy*")
reg315304 := PrimValue(reg315303)
if reg315304 == True {
reg315305 := __e.Call(__defun__shen_4line)
_ = reg315305
reg315306 := __e.Call(__defun__shen_4deref, V3722, V3724)
reg315307 := __e.Call(__defun__shen_4show_1p, reg315306)
_ = reg315307
reg315308 := MakeNumber(1)
reg315309 := __e.Call(__defun__nl, reg315308)
_ = reg315309
reg315310 := MakeNumber(1)
reg315311 := __e.Call(__defun__nl, reg315310)
_ = reg315311
reg315312 := __e.Call(__defun__shen_4deref, V3723, V3724)
reg315313 := MakeNumber(1)
reg315314 := __e.Call(__defun__shen_4show_1assumptions, reg315312, reg315313)
_ = reg315314
reg315315 := MakeString("\n> ")
reg315316 := __e.Call(__defun__stoutput)
reg315317 := __e.Call(__defun__shen_4prhush, reg315315, reg315316)
_ = reg315317
reg315318 := __e.Call(__defun__shen_4pause_1for_1user)
_ = reg315318
__ctx.TailApply(__defun__thaw, V3725)
return
} else {
__ctx.TailApply(__defun__thaw, V3725)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.show", value: __defun__shen_4show})

__defun__shen_4line = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315321 := __e.Call(__defun__inferences)
Infs := reg315321
_ = Infs
reg315322 := MakeString("____________________________________________________________ ")
reg315323 := MakeString(" inference")
reg315324 := MakeNumber(1)
reg315325 := PrimEqual(reg315324, Infs)
var reg315328 Obj
if reg315325 == True {
reg315326 := MakeString("")
reg315328 = reg315326
} else {
reg315327 := MakeString("s")
reg315328 = reg315327
}
reg315329 := MakeString(" \n?- ")
reg315330 := MakeSymbol("shen.a")
reg315331 := __e.Call(__defun__shen_4app, reg315328, reg315329, reg315330)
reg315332 := PrimStringConcat(reg315323, reg315331)
reg315333 := MakeSymbol("shen.a")
reg315334 := __e.Call(__defun__shen_4app, Infs, reg315332, reg315333)
reg315335 := PrimStringConcat(reg315322, reg315334)
reg315336 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg315335, reg315336)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.line", value: __defun__shen_4line})

__defun__shen_4show_1p = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3727 := __args[0]
_ = V3727
reg315338 := PrimIsPair(V3727)
var reg315372 Obj
if reg315338 == True {
reg315339 := PrimTail(V3727)
reg315340 := PrimIsPair(reg315339)
var reg315367 Obj
if reg315340 == True {
reg315341 := MakeSymbol(":")
reg315342 := PrimTail(V3727)
reg315343 := PrimHead(reg315342)
reg315344 := PrimEqual(reg315341, reg315343)
var reg315362 Obj
if reg315344 == True {
reg315345 := PrimTail(V3727)
reg315346 := PrimTail(reg315345)
reg315347 := PrimIsPair(reg315346)
var reg315357 Obj
if reg315347 == True {
reg315348 := Nil;
reg315349 := PrimTail(V3727)
reg315350 := PrimTail(reg315349)
reg315351 := PrimTail(reg315350)
reg315352 := PrimEqual(reg315348, reg315351)
var reg315355 Obj
if reg315352 == True {
reg315353 := True;
reg315355 = reg315353
} else {
reg315354 := False;
reg315355 = reg315354
}
reg315357 = reg315355
} else {
reg315356 := False;
reg315357 = reg315356
}
var reg315360 Obj
if reg315357 == True {
reg315358 := True;
reg315360 = reg315358
} else {
reg315359 := False;
reg315360 = reg315359
}
reg315362 = reg315360
} else {
reg315361 := False;
reg315362 = reg315361
}
var reg315365 Obj
if reg315362 == True {
reg315363 := True;
reg315365 = reg315363
} else {
reg315364 := False;
reg315365 = reg315364
}
reg315367 = reg315365
} else {
reg315366 := False;
reg315367 = reg315366
}
var reg315370 Obj
if reg315367 == True {
reg315368 := True;
reg315370 = reg315368
} else {
reg315369 := False;
reg315370 = reg315369
}
reg315372 = reg315370
} else {
reg315371 := False;
reg315372 = reg315371
}
if reg315372 == True {
reg315373 := PrimHead(V3727)
reg315374 := MakeString(" : ")
reg315375 := PrimTail(V3727)
reg315376 := PrimTail(reg315375)
reg315377 := PrimHead(reg315376)
reg315378 := MakeString("")
reg315379 := MakeSymbol("shen.r")
reg315380 := __e.Call(__defun__shen_4app, reg315377, reg315378, reg315379)
reg315381 := PrimStringConcat(reg315374, reg315380)
reg315382 := MakeSymbol("shen.r")
reg315383 := __e.Call(__defun__shen_4app, reg315373, reg315381, reg315382)
reg315384 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg315383, reg315384)
return
} else {
reg315386 := MakeString("")
reg315387 := MakeSymbol("shen.r")
reg315388 := __e.Call(__defun__shen_4app, V3727, reg315386, reg315387)
reg315389 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg315388, reg315389)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.show-p", value: __defun__shen_4show_1p})

__defun__shen_4show_1assumptions = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3732 := __args[0]
_ = V3732
V3733 := __args[1]
_ = V3733
reg315391 := Nil;
reg315392 := PrimEqual(reg315391, V3732)
if reg315392 == True {
reg315393 := MakeSymbol("shen.skip")
__ctx.Return(reg315393)
return
} else {
reg315394 := PrimIsPair(V3732)
if reg315394 == True {
reg315395 := MakeString(". ")
reg315396 := MakeSymbol("shen.a")
reg315397 := __e.Call(__defun__shen_4app, V3733, reg315395, reg315396)
reg315398 := __e.Call(__defun__stoutput)
reg315399 := __e.Call(__defun__shen_4prhush, reg315397, reg315398)
_ = reg315399
reg315400 := PrimHead(V3732)
reg315401 := __e.Call(__defun__shen_4show_1p, reg315400)
_ = reg315401
reg315402 := MakeNumber(1)
reg315403 := __e.Call(__defun__nl, reg315402)
_ = reg315403
reg315404 := PrimTail(V3732)
reg315405 := MakeNumber(1)
reg315406 := PrimNumberAdd(V3733, reg315405)
__ctx.TailApply(__defun__shen_4show_1assumptions, reg315404, reg315406)
return
} else {
reg315408 := MakeSymbol("shen.show-assumptions")
__ctx.TailApply(__defun__shen_4f__error, reg315408)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.show-assumptions", value: __defun__shen_4show_1assumptions})

__defun__shen_4pause_1for_1user = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315410 := __e.Call(__defun__stinput)
reg315411 := PrimReadByte(reg315410)
Byte := reg315411
_ = Byte
reg315412 := MakeNumber(94)
reg315413 := PrimEqual(Byte, reg315412)
if reg315413 == True {
reg315414 := MakeString("input aborted\n")
reg315415 := PrimSimpleError(reg315414)
__ctx.Return(reg315415)
return
} else {
reg315416 := MakeNumber(1)
__ctx.TailApply(__defun__nl, reg315416)
return
}
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.pause-for-user", value: __defun__shen_4pause_1for_1user})

__defun__shen_4typedf_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3735 := __args[0]
_ = V3735
reg315418 := MakeSymbol("shen.*signedfuncs*")
reg315419 := PrimValue(reg315418)
reg315420 := __e.Call(__defun__assoc, V3735, reg315419)
reg315421 := PrimIsPair(reg315420)
__ctx.Return(reg315421)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.typedf?", value: __defun__shen_4typedf_2})

__defun__shen_4sigf = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3737 := __args[0]
_ = V3737
reg315422 := MakeSymbol("shen.type-signature-of-")
__ctx.TailApply(__defun__concat, reg315422, V3737)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sigf", value: __defun__shen_4sigf})

__defun__shen_4placeholder = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315424 := MakeSymbol("&&")
__ctx.TailApply(__defun__gensym, reg315424)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.placeholder", value: __defun__shen_4placeholder})

__defun__shen_4base = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3742 := __args[0]
_ = V3742
V3743 := __args[1]
_ = V3743
V3744 := __args[2]
_ = V3744
V3745 := __args[3]
_ = V3745
reg315426 := __e.Call(__defun__shen_4lazyderef, V3743, V3744)
V3426 := reg315426
_ = V3426
reg315427 := MakeSymbol("number")
reg315428 := PrimEqual(reg315427, V3426)
var reg315443 Obj
if reg315428 == True {
reg315429 := __e.Call(__defun__shen_4incinfs)
_ = reg315429
reg315430 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315431 := PrimIsNumber(reg315430)
reg315432 := __e.Call(__defun__fwhen, reg315431, V3744, V3745)
reg315443 = reg315432
} else {
reg315433 := __e.Call(__defun__shen_4pvar_2, V3426)
var reg315442 Obj
if reg315433 == True {
reg315434 := MakeSymbol("number")
reg315435 := __e.Call(__defun__shen_4bindv, V3426, reg315434, V3744)
_ = reg315435
reg315436 := __e.Call(__defun__shen_4incinfs)
_ = reg315436
reg315437 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315438 := PrimIsNumber(reg315437)
reg315439 := __e.Call(__defun__fwhen, reg315438, V3744, V3745)
Result := reg315439
_ = Result
reg315440 := __e.Call(__defun__shen_4unbindv, V3426, V3744)
_ = reg315440
reg315442 = Result
} else {
reg315441 := False;
reg315442 = reg315441
}
reg315443 = reg315442
}
Case := reg315443
_ = Case
reg315444 := False;
reg315445 := PrimEqual(Case, reg315444)
if reg315445 == True {
reg315446 := __e.Call(__defun__shen_4lazyderef, V3743, V3744)
V3427 := reg315446
_ = V3427
reg315447 := MakeSymbol("boolean")
reg315448 := PrimEqual(reg315447, V3427)
var reg315463 Obj
if reg315448 == True {
reg315449 := __e.Call(__defun__shen_4incinfs)
_ = reg315449
reg315450 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315451 := __e.Call(__defun__boolean_2, reg315450)
reg315452 := __e.Call(__defun__fwhen, reg315451, V3744, V3745)
reg315463 = reg315452
} else {
reg315453 := __e.Call(__defun__shen_4pvar_2, V3427)
var reg315462 Obj
if reg315453 == True {
reg315454 := MakeSymbol("boolean")
reg315455 := __e.Call(__defun__shen_4bindv, V3427, reg315454, V3744)
_ = reg315455
reg315456 := __e.Call(__defun__shen_4incinfs)
_ = reg315456
reg315457 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315458 := __e.Call(__defun__boolean_2, reg315457)
reg315459 := __e.Call(__defun__fwhen, reg315458, V3744, V3745)
Result := reg315459
_ = Result
reg315460 := __e.Call(__defun__shen_4unbindv, V3427, V3744)
_ = reg315460
reg315462 = Result
} else {
reg315461 := False;
reg315462 = reg315461
}
reg315463 = reg315462
}
Case := reg315463
_ = Case
reg315464 := False;
reg315465 := PrimEqual(Case, reg315464)
if reg315465 == True {
reg315466 := __e.Call(__defun__shen_4lazyderef, V3743, V3744)
V3428 := reg315466
_ = V3428
reg315467 := MakeSymbol("string")
reg315468 := PrimEqual(reg315467, V3428)
var reg315483 Obj
if reg315468 == True {
reg315469 := __e.Call(__defun__shen_4incinfs)
_ = reg315469
reg315470 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315471 := PrimIsString(reg315470)
reg315472 := __e.Call(__defun__fwhen, reg315471, V3744, V3745)
reg315483 = reg315472
} else {
reg315473 := __e.Call(__defun__shen_4pvar_2, V3428)
var reg315482 Obj
if reg315473 == True {
reg315474 := MakeSymbol("string")
reg315475 := __e.Call(__defun__shen_4bindv, V3428, reg315474, V3744)
_ = reg315475
reg315476 := __e.Call(__defun__shen_4incinfs)
_ = reg315476
reg315477 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315478 := PrimIsString(reg315477)
reg315479 := __e.Call(__defun__fwhen, reg315478, V3744, V3745)
Result := reg315479
_ = Result
reg315480 := __e.Call(__defun__shen_4unbindv, V3428, V3744)
_ = reg315480
reg315482 = Result
} else {
reg315481 := False;
reg315482 = reg315481
}
reg315483 = reg315482
}
Case := reg315483
_ = Case
reg315484 := False;
reg315485 := PrimEqual(Case, reg315484)
if reg315485 == True {
reg315486 := __e.Call(__defun__shen_4lazyderef, V3743, V3744)
V3429 := reg315486
_ = V3429
reg315487 := MakeSymbol("symbol")
reg315488 := PrimEqual(reg315487, V3429)
var reg315513 Obj
if reg315488 == True {
reg315489 := __e.Call(__defun__shen_4incinfs)
_ = reg315489
reg315490 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315491 := PrimIsSymbol(reg315490)
reg315492 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315493 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315494 := __e.Call(__defun__shen_4ue_2, reg315493)
reg315495 := PrimNot(reg315494)
__ctx.TailApply(__defun__fwhen, reg315495, V3744, V3745)
return
}, 0)
reg315497 := __e.Call(__defun__fwhen, reg315491, V3744, reg315492)
reg315513 = reg315497
} else {
reg315498 := __e.Call(__defun__shen_4pvar_2, V3429)
var reg315512 Obj
if reg315498 == True {
reg315499 := MakeSymbol("symbol")
reg315500 := __e.Call(__defun__shen_4bindv, V3429, reg315499, V3744)
_ = reg315500
reg315501 := __e.Call(__defun__shen_4incinfs)
_ = reg315501
reg315502 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315503 := PrimIsSymbol(reg315502)
reg315504 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315505 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
reg315506 := __e.Call(__defun__shen_4ue_2, reg315505)
reg315507 := PrimNot(reg315506)
__ctx.TailApply(__defun__fwhen, reg315507, V3744, V3745)
return
}, 0)
reg315509 := __e.Call(__defun__fwhen, reg315503, V3744, reg315504)
Result := reg315509
_ = Result
reg315510 := __e.Call(__defun__shen_4unbindv, V3429, V3744)
_ = reg315510
reg315512 = Result
} else {
reg315511 := False;
reg315512 = reg315511
}
reg315513 = reg315512
}
Case := reg315513
_ = Case
reg315514 := False;
reg315515 := PrimEqual(Case, reg315514)
if reg315515 == True {
reg315516 := __e.Call(__defun__shen_4lazyderef, V3742, V3744)
V3430 := reg315516
_ = V3430
reg315517 := Nil;
reg315518 := PrimEqual(reg315517, V3430)
if reg315518 == True {
reg315519 := __e.Call(__defun__shen_4lazyderef, V3743, V3744)
V3431 := reg315519
_ = V3431
reg315520 := PrimIsPair(V3431)
if reg315520 == True {
reg315521 := PrimHead(V3431)
reg315522 := __e.Call(__defun__shen_4lazyderef, reg315521, V3744)
V3432 := reg315522
_ = V3432
reg315523 := MakeSymbol("list")
reg315524 := PrimEqual(reg315523, V3432)
if reg315524 == True {
reg315525 := PrimTail(V3431)
reg315526 := __e.Call(__defun__shen_4lazyderef, reg315525, V3744)
V3433 := reg315526
_ = V3433
reg315527 := PrimIsPair(V3433)
if reg315527 == True {
reg315528 := PrimHead(V3433)
A := reg315528
_ = A
reg315529 := PrimTail(V3433)
reg315530 := __e.Call(__defun__shen_4lazyderef, reg315529, V3744)
V3434 := reg315530
_ = V3434
reg315531 := Nil;
reg315532 := PrimEqual(reg315531, V3434)
if reg315532 == True {
reg315533 := __e.Call(__defun__shen_4incinfs)
_ = reg315533
__ctx.TailApply(__defun__thaw, V3745)
return
} else {
reg315535 := __e.Call(__defun__shen_4pvar_2, V3434)
if reg315535 == True {
reg315536 := Nil;
reg315537 := __e.Call(__defun__shen_4bindv, V3434, reg315536, V3744)
_ = reg315537
reg315538 := __e.Call(__defun__shen_4incinfs)
_ = reg315538
reg315539 := __e.Call(__defun__thaw, V3745)
Result := reg315539
_ = Result
reg315540 := __e.Call(__defun__shen_4unbindv, V3434, V3744)
_ = reg315540
__ctx.Return(Result)
return
} else {
reg315541 := False;
__ctx.Return(reg315541)
return
}
}
} else {
reg315542 := __e.Call(__defun__shen_4pvar_2, V3433)
if reg315542 == True {
reg315543 := __e.Call(__defun__shen_4newpv, V3744)
A := reg315543
_ = A
reg315544 := Nil;
reg315545 := PrimCons(A, reg315544)
reg315546 := __e.Call(__defun__shen_4bindv, V3433, reg315545, V3744)
_ = reg315546
reg315547 := __e.Call(__defun__shen_4incinfs)
_ = reg315547
reg315548 := __e.Call(__defun__thaw, V3745)
Result := reg315548
_ = Result
reg315549 := __e.Call(__defun__shen_4unbindv, V3433, V3744)
_ = reg315549
__ctx.Return(Result)
return
} else {
reg315550 := False;
__ctx.Return(reg315550)
return
}
}
} else {
reg315551 := __e.Call(__defun__shen_4pvar_2, V3432)
if reg315551 == True {
reg315552 := MakeSymbol("list")
reg315553 := __e.Call(__defun__shen_4bindv, V3432, reg315552, V3744)
_ = reg315553
reg315554 := PrimTail(V3431)
reg315555 := __e.Call(__defun__shen_4lazyderef, reg315554, V3744)
V3435 := reg315555
_ = V3435
reg315556 := PrimIsPair(V3435)
var reg315583 Obj
if reg315556 == True {
reg315557 := PrimHead(V3435)
A := reg315557
_ = A
reg315558 := PrimTail(V3435)
reg315559 := __e.Call(__defun__shen_4lazyderef, reg315558, V3744)
V3436 := reg315559
_ = V3436
reg315560 := Nil;
reg315561 := PrimEqual(reg315560, V3436)
var reg315572 Obj
if reg315561 == True {
reg315562 := __e.Call(__defun__shen_4incinfs)
_ = reg315562
reg315563 := __e.Call(__defun__thaw, V3745)
reg315572 = reg315563
} else {
reg315564 := __e.Call(__defun__shen_4pvar_2, V3436)
var reg315571 Obj
if reg315564 == True {
reg315565 := Nil;
reg315566 := __e.Call(__defun__shen_4bindv, V3436, reg315565, V3744)
_ = reg315566
reg315567 := __e.Call(__defun__shen_4incinfs)
_ = reg315567
reg315568 := __e.Call(__defun__thaw, V3745)
Result := reg315568
_ = Result
reg315569 := __e.Call(__defun__shen_4unbindv, V3436, V3744)
_ = reg315569
reg315571 = Result
} else {
reg315570 := False;
reg315571 = reg315570
}
reg315572 = reg315571
}
reg315583 = reg315572
} else {
reg315573 := __e.Call(__defun__shen_4pvar_2, V3435)
var reg315582 Obj
if reg315573 == True {
reg315574 := __e.Call(__defun__shen_4newpv, V3744)
A := reg315574
_ = A
reg315575 := Nil;
reg315576 := PrimCons(A, reg315575)
reg315577 := __e.Call(__defun__shen_4bindv, V3435, reg315576, V3744)
_ = reg315577
reg315578 := __e.Call(__defun__shen_4incinfs)
_ = reg315578
reg315579 := __e.Call(__defun__thaw, V3745)
Result := reg315579
_ = Result
reg315580 := __e.Call(__defun__shen_4unbindv, V3435, V3744)
_ = reg315580
reg315582 = Result
} else {
reg315581 := False;
reg315582 = reg315581
}
reg315583 = reg315582
}
Result := reg315583
_ = Result
reg315584 := __e.Call(__defun__shen_4unbindv, V3432, V3744)
_ = reg315584
__ctx.Return(Result)
return
} else {
reg315585 := False;
__ctx.Return(reg315585)
return
}
}
} else {
reg315586 := __e.Call(__defun__shen_4pvar_2, V3431)
if reg315586 == True {
reg315587 := __e.Call(__defun__shen_4newpv, V3744)
A := reg315587
_ = A
reg315588 := MakeSymbol("list")
reg315589 := Nil;
reg315590 := PrimCons(A, reg315589)
reg315591 := PrimCons(reg315588, reg315590)
reg315592 := __e.Call(__defun__shen_4bindv, V3431, reg315591, V3744)
_ = reg315592
reg315593 := __e.Call(__defun__shen_4incinfs)
_ = reg315593
reg315594 := __e.Call(__defun__thaw, V3745)
Result := reg315594
_ = Result
reg315595 := __e.Call(__defun__shen_4unbindv, V3431, V3744)
_ = reg315595
__ctx.Return(Result)
return
} else {
reg315596 := False;
__ctx.Return(reg315596)
return
}
}
} else {
reg315597 := False;
__ctx.Return(reg315597)
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
V3751 := __args[0]
_ = V3751
V3752 := __args[1]
_ = V3752
V3753 := __args[2]
_ = V3753
V3754 := __args[3]
_ = V3754
V3755 := __args[4]
_ = V3755
reg315598 := __e.Call(__defun__shen_4lazyderef, V3753, V3754)
V3417 := reg315598
_ = V3417
reg315599 := PrimIsPair(V3417)
var reg315634 Obj
if reg315599 == True {
reg315600 := PrimHead(V3417)
reg315601 := __e.Call(__defun__shen_4lazyderef, reg315600, V3754)
V3418 := reg315601
_ = V3418
reg315602 := PrimIsPair(V3418)
var reg315632 Obj
if reg315602 == True {
reg315603 := PrimHead(V3418)
Y := reg315603
_ = Y
reg315604 := PrimTail(V3418)
reg315605 := __e.Call(__defun__shen_4lazyderef, reg315604, V3754)
V3419 := reg315605
_ = V3419
reg315606 := PrimIsPair(V3419)
var reg315630 Obj
if reg315606 == True {
reg315607 := PrimHead(V3419)
reg315608 := __e.Call(__defun__shen_4lazyderef, reg315607, V3754)
V3420 := reg315608
_ = V3420
reg315609 := MakeSymbol(":")
reg315610 := PrimEqual(reg315609, V3420)
var reg315628 Obj
if reg315610 == True {
reg315611 := PrimTail(V3419)
reg315612 := __e.Call(__defun__shen_4lazyderef, reg315611, V3754)
V3421 := reg315612
_ = V3421
reg315613 := PrimIsPair(V3421)
var reg315626 Obj
if reg315613 == True {
reg315614 := PrimHead(V3421)
B := reg315614
_ = B
reg315615 := PrimTail(V3421)
reg315616 := __e.Call(__defun__shen_4lazyderef, reg315615, V3754)
V3422 := reg315616
_ = V3422
reg315617 := Nil;
reg315618 := PrimEqual(reg315617, V3422)
var reg315624 Obj
if reg315618 == True {
reg315619 := __e.Call(__defun__shen_4incinfs)
_ = reg315619
reg315620 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__unify_b, V3752, B, V3754, V3755)
return
}, 0)
reg315622 := __e.Call(__defun__identical, V3751, Y, V3754, reg315620)
reg315624 = reg315622
} else {
reg315623 := False;
reg315624 = reg315623
}
reg315626 = reg315624
} else {
reg315625 := False;
reg315626 = reg315625
}
reg315628 = reg315626
} else {
reg315627 := False;
reg315628 = reg315627
}
reg315630 = reg315628
} else {
reg315629 := False;
reg315630 = reg315629
}
reg315632 = reg315630
} else {
reg315631 := False;
reg315632 = reg315631
}
reg315634 = reg315632
} else {
reg315633 := False;
reg315634 = reg315633
}
Case := reg315634
_ = Case
reg315635 := False;
reg315636 := PrimEqual(Case, reg315635)
if reg315636 == True {
reg315637 := __e.Call(__defun__shen_4lazyderef, V3753, V3754)
V3423 := reg315637
_ = V3423
reg315638 := PrimIsPair(V3423)
if reg315638 == True {
reg315639 := PrimTail(V3423)
Hyp := reg315639
_ = Hyp
reg315640 := __e.Call(__defun__shen_4incinfs)
_ = reg315640
__ctx.TailApply(__defun__shen_4by__hypothesis, V3751, V3752, Hyp, V3754, V3755)
return
} else {
reg315642 := False;
__ctx.Return(reg315642)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.by_hypothesis", value: __defun__shen_4by__hypothesis})

__defun__shen_4t_d_1def = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3761 := __args[0]
_ = V3761
V3762 := __args[1]
_ = V3762
V3763 := __args[2]
_ = V3763
V3764 := __args[3]
_ = V3764
V3765 := __args[4]
_ = V3765
reg315643 := __e.Call(__defun__shen_4lazyderef, V3761, V3764)
V3411 := reg315643
_ = V3411
reg315644 := PrimIsPair(V3411)
if reg315644 == True {
reg315645 := PrimHead(V3411)
reg315646 := __e.Call(__defun__shen_4lazyderef, reg315645, V3764)
V3412 := reg315646
_ = V3412
reg315647 := MakeSymbol("define")
reg315648 := PrimEqual(reg315647, V3412)
if reg315648 == True {
reg315649 := PrimTail(V3411)
reg315650 := __e.Call(__defun__shen_4lazyderef, reg315649, V3764)
V3413 := reg315650
_ = V3413
reg315651 := PrimIsPair(V3413)
if reg315651 == True {
reg315652 := PrimHead(V3413)
F := reg315652
_ = F
reg315653 := PrimTail(V3413)
X := reg315653
_ = X
reg315654 := __e.Call(__defun__shen_4newpv, V3764)
Y := reg315654
_ = Y
reg315655 := __e.Call(__defun__shen_4newpv, V3764)
E := reg315655
_ = E
reg315656 := __e.Call(__defun__shen_4incinfs)
_ = reg315656
reg315657 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4_5sig_7rules_6, Y)
return
}, 1)
reg315659 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg315660 := PrimIsPair(E)
if reg315660 == True {
reg315661 := MakeString("parse error here: ")
reg315662 := MakeString("\n")
reg315663 := MakeSymbol("shen.s")
reg315664 := __e.Call(__defun__shen_4app, E, reg315662, reg315663)
reg315665 := PrimStringConcat(reg315661, reg315664)
reg315666 := PrimSimpleError(reg315665)
__ctx.Return(reg315666)
return
} else {
reg315667 := MakeString("parse error\n")
reg315668 := PrimSimpleError(reg315667)
__ctx.Return(reg315668)
return
}
}, 1)
reg315669 := __e.Call(__defun__compile, reg315657, X, reg315659)
__ctx.TailApply(__defun__shen_4t_d_1defh, reg315669, F, V3762, V3763, V3764, V3765)
return
} else {
reg315671 := False;
__ctx.Return(reg315671)
return
}
} else {
reg315672 := False;
__ctx.Return(reg315672)
return
}
} else {
reg315673 := False;
__ctx.Return(reg315673)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-def", value: __defun__shen_4t_d_1def})

__defun__shen_4t_d_1defh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3772 := __args[0]
_ = V3772
V3773 := __args[1]
_ = V3773
V3774 := __args[2]
_ = V3774
V3775 := __args[3]
_ = V3775
V3776 := __args[4]
_ = V3776
V3777 := __args[5]
_ = V3777
reg315674 := __e.Call(__defun__shen_4lazyderef, V3772, V3776)
V3407 := reg315674
_ = V3407
reg315675 := PrimIsPair(V3407)
if reg315675 == True {
reg315676 := PrimHead(V3407)
Sig := reg315676
_ = Sig
reg315677 := PrimTail(V3407)
Rules := reg315677
_ = Rules
reg315678 := __e.Call(__defun__shen_4incinfs)
_ = reg315678
reg315679 := __e.Call(__defun__shen_4ue_1sig, Sig)
__ctx.TailApply(__defun__shen_4t_d_1defhh, Sig, reg315679, V3773, V3774, V3775, Rules, V3776, V3777)
return
} else {
reg315681 := False;
__ctx.Return(reg315681)
return
}
}, 6)
__initDefs = append(__initDefs, defType{name: "shen.t*-defh", value: __defun__shen_4t_d_1defh})

__defun__shen_4t_d_1defhh = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3786 := __args[0]
_ = V3786
V3787 := __args[1]
_ = V3787
V3788 := __args[2]
_ = V3788
V3789 := __args[3]
_ = V3789
V3790 := __args[4]
_ = V3790
V3791 := __args[5]
_ = V3791
V3792 := __args[6]
_ = V3792
V3793 := __args[7]
_ = V3793
reg315682 := __e.Call(__defun__shen_4incinfs)
_ = reg315682
reg315683 := MakeNumber(1)
reg315684 := MakeSymbol(":")
reg315685 := Nil;
reg315686 := PrimCons(V3787, reg315685)
reg315687 := PrimCons(reg315684, reg315686)
reg315688 := PrimCons(V3788, reg315687)
reg315689 := PrimCons(reg315688, V3790)
reg315690 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4memo, V3788, V3786, V3789, V3792, V3793)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d_1rules, V3791, V3787, reg315683, V3788, reg315689, V3792, reg315690)
return
}, 8)
__initDefs = append(__initDefs, defType{name: "shen.t*-defhh", value: __defun__shen_4t_d_1defhh})

__defun__shen_4memo = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3799 := __args[0]
_ = V3799
V3800 := __args[1]
_ = V3800
V3801 := __args[2]
_ = V3801
V3802 := __args[3]
_ = V3802
V3803 := __args[4]
_ = V3803
reg315693 := __e.Call(__defun__shen_4newpv, V3802)
Jnk := reg315693
_ = Jnk
reg315694 := __e.Call(__defun__shen_4incinfs)
_ = reg315694
reg315695 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315696 := __e.Call(__defun__shen_4lazyderef, V3799, V3802)
reg315697 := __e.Call(__defun__shen_4lazyderef, V3801, V3802)
reg315698 := __e.Call(__defun__declare, reg315696, reg315697)
__ctx.TailApply(__defun__bind, Jnk, reg315698, V3802, V3803)
return
}, 0)
__ctx.TailApply(__defun__unify_b, V3801, V3800, V3802, reg315695)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.memo", value: __defun__shen_4memo})

__defun__shen_4_5sig_7rules_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3805 := __args[0]
_ = V3805
reg315701 := __e.Call(__defun__shen_4_5signature_6, V3805)
Parse__shen_4_5signature_6 := reg315701
_ = Parse__shen_4_5signature_6
reg315702 := __e.Call(__defun__fail)
reg315703 := PrimEqual(reg315702, Parse__shen_4_5signature_6)
reg315704 := PrimNot(reg315703)
if reg315704 == True {
reg315705 := __e.Call(__defun__shen_4_5non_1ll_1rules_6, Parse__shen_4_5signature_6)
Parse__shen_4_5non_1ll_1rules_6 := reg315705
_ = Parse__shen_4_5non_1ll_1rules_6
reg315706 := __e.Call(__defun__fail)
reg315707 := PrimEqual(reg315706, Parse__shen_4_5non_1ll_1rules_6)
reg315708 := PrimNot(reg315707)
if reg315708 == True {
reg315709 := PrimHead(Parse__shen_4_5non_1ll_1rules_6)
reg315710 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_6)
reg315711 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5non_1ll_1rules_6)
reg315712 := PrimCons(reg315710, reg315711)
__ctx.TailApply(__defun__shen_4pair, reg315709, reg315712)
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
V3807 := __args[0]
_ = V3807
reg315716 := __e.Call(__defun__shen_4_5rule_6, V3807)
Parse__shen_4_5rule_6 := reg315716
_ = Parse__shen_4_5rule_6
reg315717 := __e.Call(__defun__fail)
reg315718 := PrimEqual(reg315717, Parse__shen_4_5rule_6)
reg315719 := PrimNot(reg315718)
var reg315732 Obj
if reg315719 == True {
reg315720 := __e.Call(__defun__shen_4_5non_1ll_1rules_6, Parse__shen_4_5rule_6)
Parse__shen_4_5non_1ll_1rules_6 := reg315720
_ = Parse__shen_4_5non_1ll_1rules_6
reg315721 := __e.Call(__defun__fail)
reg315722 := PrimEqual(reg315721, Parse__shen_4_5non_1ll_1rules_6)
reg315723 := PrimNot(reg315722)
var reg315730 Obj
if reg315723 == True {
reg315724 := PrimHead(Parse__shen_4_5non_1ll_1rules_6)
reg315725 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg315726 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5non_1ll_1rules_6)
reg315727 := PrimCons(reg315725, reg315726)
reg315728 := __e.Call(__defun__shen_4pair, reg315724, reg315727)
reg315730 = reg315728
} else {
reg315729 := __e.Call(__defun__fail)
reg315730 = reg315729
}
reg315732 = reg315730
} else {
reg315731 := __e.Call(__defun__fail)
reg315732 = reg315731
}
YaccParse := reg315732
_ = YaccParse
reg315733 := __e.Call(__defun__fail)
reg315734 := PrimEqual(YaccParse, reg315733)
if reg315734 == True {
reg315735 := __e.Call(__defun__shen_4_5rule_6, V3807)
Parse__shen_4_5rule_6 := reg315735
_ = Parse__shen_4_5rule_6
reg315736 := __e.Call(__defun__fail)
reg315737 := PrimEqual(reg315736, Parse__shen_4_5rule_6)
reg315738 := PrimNot(reg315737)
if reg315738 == True {
reg315739 := PrimHead(Parse__shen_4_5rule_6)
reg315740 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg315741 := Nil;
reg315742 := PrimCons(reg315740, reg315741)
__ctx.TailApply(__defun__shen_4pair, reg315739, reg315742)
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
V3809 := __args[0]
_ = V3809
reg315745 := PrimIsPair(V3809)
var reg315769 Obj
if reg315745 == True {
reg315746 := PrimTail(V3809)
reg315747 := PrimIsPair(reg315746)
var reg315764 Obj
if reg315747 == True {
reg315748 := Nil;
reg315749 := PrimTail(V3809)
reg315750 := PrimTail(reg315749)
reg315751 := PrimEqual(reg315748, reg315750)
var reg315759 Obj
if reg315751 == True {
reg315752 := PrimHead(V3809)
reg315753 := MakeSymbol("protect")
reg315754 := PrimEqual(reg315752, reg315753)
var reg315757 Obj
if reg315754 == True {
reg315755 := True;
reg315757 = reg315755
} else {
reg315756 := False;
reg315757 = reg315756
}
reg315759 = reg315757
} else {
reg315758 := False;
reg315759 = reg315758
}
var reg315762 Obj
if reg315759 == True {
reg315760 := True;
reg315762 = reg315760
} else {
reg315761 := False;
reg315762 = reg315761
}
reg315764 = reg315762
} else {
reg315763 := False;
reg315764 = reg315763
}
var reg315767 Obj
if reg315764 == True {
reg315765 := True;
reg315767 = reg315765
} else {
reg315766 := False;
reg315767 = reg315766
}
reg315769 = reg315767
} else {
reg315768 := False;
reg315769 = reg315768
}
if reg315769 == True {
__ctx.Return(V3809)
return
} else {
reg315770 := PrimIsPair(V3809)
if reg315770 == True {
reg315771 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4ue, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg315771, V3809)
return
} else {
reg315774 := PrimIsVariable(V3809)
if reg315774 == True {
reg315775 := MakeSymbol("&&")
__ctx.TailApply(__defun__concat, reg315775, V3809)
return
} else {
__ctx.Return(V3809)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue", value: __defun__shen_4ue})

__defun__shen_4ue_1sig = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3811 := __args[0]
_ = V3811
reg315777 := PrimIsPair(V3811)
if reg315777 == True {
reg315778 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4ue_1sig, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg315778, V3811)
return
} else {
reg315781 := PrimIsVariable(V3811)
if reg315781 == True {
reg315782 := MakeSymbol("&&&")
__ctx.TailApply(__defun__concat, reg315782, V3811)
return
} else {
__ctx.Return(V3811)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue-sig", value: __defun__shen_4ue_1sig})

__defun__shen_4ues = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3817 := __args[0]
_ = V3817
reg315784 := __e.Call(__defun__shen_4ue_2, V3817)
if reg315784 == True {
reg315785 := Nil;
reg315786 := PrimCons(V3817, reg315785)
__ctx.Return(reg315786)
return
} else {
reg315787 := PrimIsPair(V3817)
if reg315787 == True {
reg315788 := PrimHead(V3817)
reg315789 := __e.Call(__defun__shen_4ues, reg315788)
reg315790 := PrimTail(V3817)
reg315791 := __e.Call(__defun__shen_4ues, reg315790)
__ctx.TailApply(__defun__union, reg315789, reg315791)
return
} else {
reg315793 := Nil;
__ctx.Return(reg315793)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ues", value: __defun__shen_4ues})

__defun__shen_4ue_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3819 := __args[0]
_ = V3819
reg315794 := PrimIsSymbol(V3819)
if reg315794 == True {
reg315795 := PrimStr(V3819)
reg315796 := __e.Call(__defun__shen_4ue_1h_2, reg315795)
if reg315796 == True {
reg315797 := True;
__ctx.Return(reg315797)
return
} else {
reg315798 := False;
__ctx.Return(reg315798)
return
}
} else {
reg315799 := False;
__ctx.Return(reg315799)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue?", value: __defun__shen_4ue_2})

__defun__shen_4ue_1h_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3827 := __args[0]
_ = V3827
reg315800 := __e.Call(__defun__shen_4_7string_2, V3827)
var reg315826 Obj
if reg315800 == True {
reg315801 := MakeString("&")
reg315802 := MakeNumber(0)
reg315803 := PrimPos(V3827, reg315802)
reg315804 := PrimEqual(reg315801, reg315803)
var reg315821 Obj
if reg315804 == True {
reg315805 := PrimTailString(V3827)
reg315806 := __e.Call(__defun__shen_4_7string_2, reg315805)
var reg315816 Obj
if reg315806 == True {
reg315807 := MakeString("&")
reg315808 := PrimTailString(V3827)
reg315809 := MakeNumber(0)
reg315810 := PrimPos(reg315808, reg315809)
reg315811 := PrimEqual(reg315807, reg315810)
var reg315814 Obj
if reg315811 == True {
reg315812 := True;
reg315814 = reg315812
} else {
reg315813 := False;
reg315814 = reg315813
}
reg315816 = reg315814
} else {
reg315815 := False;
reg315816 = reg315815
}
var reg315819 Obj
if reg315816 == True {
reg315817 := True;
reg315819 = reg315817
} else {
reg315818 := False;
reg315819 = reg315818
}
reg315821 = reg315819
} else {
reg315820 := False;
reg315821 = reg315820
}
var reg315824 Obj
if reg315821 == True {
reg315822 := True;
reg315824 = reg315822
} else {
reg315823 := False;
reg315824 = reg315823
}
reg315826 = reg315824
} else {
reg315825 := False;
reg315826 = reg315825
}
if reg315826 == True {
reg315827 := True;
__ctx.Return(reg315827)
return
} else {
reg315828 := False;
__ctx.Return(reg315828)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.ue-h?", value: __defun__shen_4ue_1h_2})

__defun__shen_4t_d_1rules = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3835 := __args[0]
_ = V3835
V3836 := __args[1]
_ = V3836
V3837 := __args[2]
_ = V3837
V3838 := __args[3]
_ = V3838
V3839 := __args[4]
_ = V3839
V3840 := __args[5]
_ = V3840
V3841 := __args[6]
_ = V3841
reg315829 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg315829
_ = Throwcontrol
reg315830 := __e.Call(__defun__shen_4lazyderef, V3835, V3840)
V3391 := reg315830
_ = V3391
reg315831 := Nil;
reg315832 := PrimEqual(reg315831, V3391)
var reg315836 Obj
if reg315832 == True {
reg315833 := __e.Call(__defun__shen_4incinfs)
_ = reg315833
reg315834 := __e.Call(__defun__thaw, V3841)
reg315836 = reg315834
} else {
reg315835 := False;
reg315836 = reg315835
}
Case := reg315836
_ = Case
reg315837 := False;
reg315838 := PrimEqual(Case, reg315837)
var reg315872 Obj
if reg315838 == True {
reg315839 := __e.Call(__defun__shen_4lazyderef, V3835, V3840)
V3392 := reg315839
_ = V3392
reg315840 := PrimIsPair(V3392)
var reg315853 Obj
if reg315840 == True {
reg315841 := PrimHead(V3392)
Rule := reg315841
_ = Rule
reg315842 := PrimTail(V3392)
Rules := reg315842
_ = Rules
reg315843 := __e.Call(__defun__shen_4incinfs)
_ = reg315843
reg315844 := __e.Call(__defun__shen_4ue, Rule)
reg315845 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315846 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315847 := MakeNumber(1)
reg315848 := PrimNumberAdd(V3837, reg315847)
__ctx.TailApply(__defun__shen_4t_d_1rules, Rules, V3836, reg315848, V3838, V3839, V3840, V3841)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3840, reg315846)
return
}, 0)
reg315851 := __e.Call(__defun__shen_4t_d_1rule, reg315844, V3836, V3839, V3840, reg315845)
reg315853 = reg315851
} else {
reg315852 := False;
reg315853 = reg315852
}
Case := reg315853
_ = Case
reg315854 := False;
reg315855 := PrimEqual(Case, reg315854)
var reg315871 Obj
if reg315855 == True {
reg315856 := __e.Call(__defun__shen_4newpv, V3840)
Err := reg315856
_ = Err
reg315857 := __e.Call(__defun__shen_4incinfs)
_ = reg315857
reg315858 := MakeString("type error in rule ")
reg315859 := __e.Call(__defun__shen_4lazyderef, V3837, V3840)
reg315860 := MakeString(" of ")
reg315861 := __e.Call(__defun__shen_4lazyderef, V3838, V3840)
reg315862 := MakeString("")
reg315863 := MakeSymbol("shen.a")
reg315864 := __e.Call(__defun__shen_4app, reg315861, reg315862, reg315863)
reg315865 := PrimStringConcat(reg315860, reg315864)
reg315866 := MakeSymbol("shen.a")
reg315867 := __e.Call(__defun__shen_4app, reg315859, reg315865, reg315866)
reg315868 := PrimStringConcat(reg315858, reg315867)
reg315869 := PrimSimpleError(reg315868)
reg315870 := __e.Call(__defun__bind, Err, reg315869, V3840, V3841)
reg315871 = reg315870
} else {
reg315871 = Case
}
reg315872 = reg315871
} else {
reg315872 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg315872)
return
}, 7)
__initDefs = append(__initDefs, defType{name: "shen.t*-rules", value: __defun__shen_4t_d_1rules})

__defun__shen_4t_d_1rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3847 := __args[0]
_ = V3847
V3848 := __args[1]
_ = V3848
V3849 := __args[2]
_ = V3849
V3850 := __args[3]
_ = V3850
V3851 := __args[4]
_ = V3851
reg315874 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg315874
_ = Throwcontrol
reg315875 := __e.Call(__defun__shen_4lazyderef, V3847, V3850)
V3383 := reg315875
_ = V3383
reg315876 := PrimIsPair(V3383)
var reg315905 Obj
if reg315876 == True {
reg315877 := PrimHead(V3383)
Patterns := reg315877
_ = Patterns
reg315878 := PrimTail(V3383)
reg315879 := __e.Call(__defun__shen_4lazyderef, reg315878, V3850)
V3384 := reg315879
_ = V3384
reg315880 := PrimIsPair(V3384)
var reg315903 Obj
if reg315880 == True {
reg315881 := PrimHead(V3384)
Action := reg315881
_ = Action
reg315882 := PrimTail(V3384)
reg315883 := __e.Call(__defun__shen_4lazyderef, reg315882, V3850)
V3385 := reg315883
_ = V3385
reg315884 := Nil;
reg315885 := PrimEqual(reg315884, V3385)
var reg315901 Obj
if reg315885 == True {
reg315886 := __e.Call(__defun__shen_4newpv, V3850)
NewHyps := reg315886
_ = NewHyps
reg315887 := __e.Call(__defun__shen_4incinfs)
_ = reg315887
reg315888 := __e.Call(__defun__shen_4placeholders, Patterns)
reg315889 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315890 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315891 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg315892 := __e.Call(__defun__shen_4ue, Action)
reg315893 := __e.Call(__defun__shen_4curry, reg315892)
reg315894 := __e.Call(__defun__shen_4result_1type, Patterns, V3848)
reg315895 := __e.Call(__defun__shen_4patthyps, Patterns, V3848, V3849)
__ctx.TailApply(__defun__shen_4t_d_1action, reg315893, reg315894, reg315895, V3850, V3851)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3850, reg315891)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d_1patterns, Patterns, V3848, NewHyps, V3850, reg315890)
return
}, 0)
reg315899 := __e.Call(__defun__shen_4newhyps, reg315888, V3849, NewHyps, V3850, reg315889)
reg315901 = reg315899
} else {
reg315900 := False;
reg315901 = reg315900
}
reg315903 = reg315901
} else {
reg315902 := False;
reg315903 = reg315902
}
reg315905 = reg315903
} else {
reg315904 := False;
reg315905 = reg315904
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg315905)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-rule", value: __defun__shen_4t_d_1rule})

__defun__shen_4placeholders = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3857 := __args[0]
_ = V3857
reg315907 := __e.Call(__defun__shen_4ue_2, V3857)
if reg315907 == True {
reg315908 := Nil;
reg315909 := PrimCons(V3857, reg315908)
__ctx.Return(reg315909)
return
} else {
reg315910 := PrimIsPair(V3857)
if reg315910 == True {
reg315911 := PrimHead(V3857)
reg315912 := __e.Call(__defun__shen_4placeholders, reg315911)
reg315913 := PrimTail(V3857)
reg315914 := __e.Call(__defun__shen_4placeholders, reg315913)
__ctx.TailApply(__defun__union, reg315912, reg315914)
return
} else {
reg315916 := Nil;
__ctx.Return(reg315916)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.placeholders", value: __defun__shen_4placeholders})

__defun__shen_4newhyps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3863 := __args[0]
_ = V3863
V3864 := __args[1]
_ = V3864
V3865 := __args[2]
_ = V3865
V3866 := __args[3]
_ = V3866
V3867 := __args[4]
_ = V3867
reg315917 := __e.Call(__defun__shen_4lazyderef, V3863, V3866)
V3370 := reg315917
_ = V3370
reg315918 := Nil;
reg315919 := PrimEqual(reg315918, V3370)
var reg315923 Obj
if reg315919 == True {
reg315920 := __e.Call(__defun__shen_4incinfs)
_ = reg315920
reg315921 := __e.Call(__defun__unify_b, V3865, V3864, V3866, V3867)
reg315923 = reg315921
} else {
reg315922 := False;
reg315923 = reg315922
}
Case := reg315923
_ = Case
reg315924 := False;
reg315925 := PrimEqual(Case, reg315924)
if reg315925 == True {
reg315926 := __e.Call(__defun__shen_4lazyderef, V3863, V3866)
V3371 := reg315926
_ = V3371
reg315927 := PrimIsPair(V3371)
if reg315927 == True {
reg315928 := PrimHead(V3371)
V3366 := reg315928
_ = V3366
reg315929 := PrimTail(V3371)
Vs := reg315929
_ = Vs
reg315930 := __e.Call(__defun__shen_4lazyderef, V3865, V3866)
V3372 := reg315930
_ = V3372
reg315931 := PrimIsPair(V3372)
if reg315931 == True {
reg315932 := PrimHead(V3372)
reg315933 := __e.Call(__defun__shen_4lazyderef, reg315932, V3866)
V3373 := reg315933
_ = V3373
reg315934 := PrimIsPair(V3373)
if reg315934 == True {
reg315935 := PrimHead(V3373)
V := reg315935
_ = V
reg315936 := PrimTail(V3373)
reg315937 := __e.Call(__defun__shen_4lazyderef, reg315936, V3866)
V3374 := reg315937
_ = V3374
reg315938 := PrimIsPair(V3374)
if reg315938 == True {
reg315939 := PrimHead(V3374)
reg315940 := __e.Call(__defun__shen_4lazyderef, reg315939, V3866)
V3375 := reg315940
_ = V3375
reg315941 := MakeSymbol(":")
reg315942 := PrimEqual(reg315941, V3375)
if reg315942 == True {
reg315943 := PrimTail(V3374)
reg315944 := __e.Call(__defun__shen_4lazyderef, reg315943, V3866)
V3376 := reg315944
_ = V3376
reg315945 := PrimIsPair(V3376)
if reg315945 == True {
reg315946 := PrimHead(V3376)
A := reg315946
_ = A
reg315947 := PrimTail(V3376)
reg315948 := __e.Call(__defun__shen_4lazyderef, reg315947, V3866)
V3377 := reg315948
_ = V3377
reg315949 := Nil;
reg315950 := PrimEqual(reg315949, V3377)
if reg315950 == True {
reg315951 := PrimTail(V3372)
NewHyp := reg315951
_ = NewHyp
reg315952 := __e.Call(__defun__shen_4incinfs)
_ = reg315952
reg315953 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
__ctx.TailApply(__defun__unify_b, V, V3366, V3866, reg315953)
return
} else {
reg315956 := __e.Call(__defun__shen_4pvar_2, V3377)
if reg315956 == True {
reg315957 := Nil;
reg315958 := __e.Call(__defun__shen_4bindv, V3377, reg315957, V3866)
_ = reg315958
reg315959 := PrimTail(V3372)
NewHyp := reg315959
_ = NewHyp
reg315960 := __e.Call(__defun__shen_4incinfs)
_ = reg315960
reg315961 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg315963 := __e.Call(__defun__unify_b, V, V3366, V3866, reg315961)
Result := reg315963
_ = Result
reg315964 := __e.Call(__defun__shen_4unbindv, V3377, V3866)
_ = reg315964
__ctx.Return(Result)
return
} else {
reg315965 := False;
__ctx.Return(reg315965)
return
}
}
} else {
reg315966 := __e.Call(__defun__shen_4pvar_2, V3376)
if reg315966 == True {
reg315967 := __e.Call(__defun__shen_4newpv, V3866)
A := reg315967
_ = A
reg315968 := Nil;
reg315969 := PrimCons(A, reg315968)
reg315970 := __e.Call(__defun__shen_4bindv, V3376, reg315969, V3866)
_ = reg315970
reg315971 := PrimTail(V3372)
NewHyp := reg315971
_ = NewHyp
reg315972 := __e.Call(__defun__shen_4incinfs)
_ = reg315972
reg315973 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg315975 := __e.Call(__defun__unify_b, V, V3366, V3866, reg315973)
Result := reg315975
_ = Result
reg315976 := __e.Call(__defun__shen_4unbindv, V3376, V3866)
_ = reg315976
__ctx.Return(Result)
return
} else {
reg315977 := False;
__ctx.Return(reg315977)
return
}
}
} else {
reg315978 := __e.Call(__defun__shen_4pvar_2, V3375)
if reg315978 == True {
reg315979 := MakeSymbol(":")
reg315980 := __e.Call(__defun__shen_4bindv, V3375, reg315979, V3866)
_ = reg315980
reg315981 := PrimTail(V3374)
reg315982 := __e.Call(__defun__shen_4lazyderef, reg315981, V3866)
V3378 := reg315982
_ = V3378
reg315983 := PrimIsPair(V3378)
var reg316019 Obj
if reg315983 == True {
reg315984 := PrimHead(V3378)
A := reg315984
_ = A
reg315985 := PrimTail(V3378)
reg315986 := __e.Call(__defun__shen_4lazyderef, reg315985, V3866)
V3379 := reg315986
_ = V3379
reg315987 := Nil;
reg315988 := PrimEqual(reg315987, V3379)
var reg316005 Obj
if reg315988 == True {
reg315989 := PrimTail(V3372)
NewHyp := reg315989
_ = NewHyp
reg315990 := __e.Call(__defun__shen_4incinfs)
_ = reg315990
reg315991 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg315993 := __e.Call(__defun__unify_b, V, V3366, V3866, reg315991)
reg316005 = reg315993
} else {
reg315994 := __e.Call(__defun__shen_4pvar_2, V3379)
var reg316004 Obj
if reg315994 == True {
reg315995 := Nil;
reg315996 := __e.Call(__defun__shen_4bindv, V3379, reg315995, V3866)
_ = reg315996
reg315997 := PrimTail(V3372)
NewHyp := reg315997
_ = NewHyp
reg315998 := __e.Call(__defun__shen_4incinfs)
_ = reg315998
reg315999 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg316001 := __e.Call(__defun__unify_b, V, V3366, V3866, reg315999)
Result := reg316001
_ = Result
reg316002 := __e.Call(__defun__shen_4unbindv, V3379, V3866)
_ = reg316002
reg316004 = Result
} else {
reg316003 := False;
reg316004 = reg316003
}
reg316005 = reg316004
}
reg316019 = reg316005
} else {
reg316006 := __e.Call(__defun__shen_4pvar_2, V3378)
var reg316018 Obj
if reg316006 == True {
reg316007 := __e.Call(__defun__shen_4newpv, V3866)
A := reg316007
_ = A
reg316008 := Nil;
reg316009 := PrimCons(A, reg316008)
reg316010 := __e.Call(__defun__shen_4bindv, V3378, reg316009, V3866)
_ = reg316010
reg316011 := PrimTail(V3372)
NewHyp := reg316011
_ = NewHyp
reg316012 := __e.Call(__defun__shen_4incinfs)
_ = reg316012
reg316013 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg316015 := __e.Call(__defun__unify_b, V, V3366, V3866, reg316013)
Result := reg316015
_ = Result
reg316016 := __e.Call(__defun__shen_4unbindv, V3378, V3866)
_ = reg316016
reg316018 = Result
} else {
reg316017 := False;
reg316018 = reg316017
}
reg316019 = reg316018
}
Result := reg316019
_ = Result
reg316020 := __e.Call(__defun__shen_4unbindv, V3375, V3866)
_ = reg316020
__ctx.Return(Result)
return
} else {
reg316021 := False;
__ctx.Return(reg316021)
return
}
}
} else {
reg316022 := __e.Call(__defun__shen_4pvar_2, V3374)
if reg316022 == True {
reg316023 := __e.Call(__defun__shen_4newpv, V3866)
A := reg316023
_ = A
reg316024 := MakeSymbol(":")
reg316025 := Nil;
reg316026 := PrimCons(A, reg316025)
reg316027 := PrimCons(reg316024, reg316026)
reg316028 := __e.Call(__defun__shen_4bindv, V3374, reg316027, V3866)
_ = reg316028
reg316029 := PrimTail(V3372)
NewHyp := reg316029
_ = NewHyp
reg316030 := __e.Call(__defun__shen_4incinfs)
_ = reg316030
reg316031 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg316033 := __e.Call(__defun__unify_b, V, V3366, V3866, reg316031)
Result := reg316033
_ = Result
reg316034 := __e.Call(__defun__shen_4unbindv, V3374, V3866)
_ = reg316034
__ctx.Return(Result)
return
} else {
reg316035 := False;
__ctx.Return(reg316035)
return
}
}
} else {
reg316036 := __e.Call(__defun__shen_4pvar_2, V3373)
if reg316036 == True {
reg316037 := __e.Call(__defun__shen_4newpv, V3866)
V := reg316037
_ = V
reg316038 := __e.Call(__defun__shen_4newpv, V3866)
A := reg316038
_ = A
reg316039 := MakeSymbol(":")
reg316040 := Nil;
reg316041 := PrimCons(A, reg316040)
reg316042 := PrimCons(reg316039, reg316041)
reg316043 := PrimCons(V, reg316042)
reg316044 := __e.Call(__defun__shen_4bindv, V3373, reg316043, V3866)
_ = reg316044
reg316045 := PrimTail(V3372)
NewHyp := reg316045
_ = NewHyp
reg316046 := __e.Call(__defun__shen_4incinfs)
_ = reg316046
reg316047 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg316049 := __e.Call(__defun__unify_b, V, V3366, V3866, reg316047)
Result := reg316049
_ = Result
reg316050 := __e.Call(__defun__shen_4unbindv, V3373, V3866)
_ = reg316050
__ctx.Return(Result)
return
} else {
reg316051 := False;
__ctx.Return(reg316051)
return
}
}
} else {
reg316052 := __e.Call(__defun__shen_4pvar_2, V3372)
if reg316052 == True {
reg316053 := __e.Call(__defun__shen_4newpv, V3866)
V := reg316053
_ = V
reg316054 := __e.Call(__defun__shen_4newpv, V3866)
A := reg316054
_ = A
reg316055 := __e.Call(__defun__shen_4newpv, V3866)
NewHyp := reg316055
_ = NewHyp
reg316056 := MakeSymbol(":")
reg316057 := Nil;
reg316058 := PrimCons(A, reg316057)
reg316059 := PrimCons(reg316056, reg316058)
reg316060 := PrimCons(V, reg316059)
reg316061 := PrimCons(reg316060, NewHyp)
reg316062 := __e.Call(__defun__shen_4bindv, V3372, reg316061, V3866)
_ = reg316062
reg316063 := __e.Call(__defun__shen_4incinfs)
_ = reg316063
reg316064 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4newhyps, Vs, V3864, NewHyp, V3866, V3867)
return
}, 0)
reg316066 := __e.Call(__defun__unify_b, V, V3366, V3866, reg316064)
Result := reg316066
_ = Result
reg316067 := __e.Call(__defun__shen_4unbindv, V3372, V3866)
_ = reg316067
__ctx.Return(Result)
return
} else {
reg316068 := False;
__ctx.Return(reg316068)
return
}
}
} else {
reg316069 := False;
__ctx.Return(reg316069)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.newhyps", value: __defun__shen_4newhyps})

__defun__shen_4patthyps = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3873 := __args[0]
_ = V3873
V3874 := __args[1]
_ = V3874
V3875 := __args[2]
_ = V3875
reg316070 := Nil;
reg316071 := PrimEqual(reg316070, V3873)
if reg316071 == True {
__ctx.Return(V3875)
return
} else {
reg316072 := PrimIsPair(V3873)
var reg316112 Obj
if reg316072 == True {
reg316073 := PrimIsPair(V3874)
var reg316107 Obj
if reg316073 == True {
reg316074 := PrimTail(V3874)
reg316075 := PrimIsPair(reg316074)
var reg316102 Obj
if reg316075 == True {
reg316076 := MakeSymbol("-->")
reg316077 := PrimTail(V3874)
reg316078 := PrimHead(reg316077)
reg316079 := PrimEqual(reg316076, reg316078)
var reg316097 Obj
if reg316079 == True {
reg316080 := PrimTail(V3874)
reg316081 := PrimTail(reg316080)
reg316082 := PrimIsPair(reg316081)
var reg316092 Obj
if reg316082 == True {
reg316083 := Nil;
reg316084 := PrimTail(V3874)
reg316085 := PrimTail(reg316084)
reg316086 := PrimTail(reg316085)
reg316087 := PrimEqual(reg316083, reg316086)
var reg316090 Obj
if reg316087 == True {
reg316088 := True;
reg316090 = reg316088
} else {
reg316089 := False;
reg316090 = reg316089
}
reg316092 = reg316090
} else {
reg316091 := False;
reg316092 = reg316091
}
var reg316095 Obj
if reg316092 == True {
reg316093 := True;
reg316095 = reg316093
} else {
reg316094 := False;
reg316095 = reg316094
}
reg316097 = reg316095
} else {
reg316096 := False;
reg316097 = reg316096
}
var reg316100 Obj
if reg316097 == True {
reg316098 := True;
reg316100 = reg316098
} else {
reg316099 := False;
reg316100 = reg316099
}
reg316102 = reg316100
} else {
reg316101 := False;
reg316102 = reg316101
}
var reg316105 Obj
if reg316102 == True {
reg316103 := True;
reg316105 = reg316103
} else {
reg316104 := False;
reg316105 = reg316104
}
reg316107 = reg316105
} else {
reg316106 := False;
reg316107 = reg316106
}
var reg316110 Obj
if reg316107 == True {
reg316108 := True;
reg316110 = reg316108
} else {
reg316109 := False;
reg316110 = reg316109
}
reg316112 = reg316110
} else {
reg316111 := False;
reg316112 = reg316111
}
if reg316112 == True {
reg316113 := PrimHead(V3873)
reg316114 := MakeSymbol(":")
reg316115 := PrimHead(V3874)
reg316116 := Nil;
reg316117 := PrimCons(reg316115, reg316116)
reg316118 := PrimCons(reg316114, reg316117)
reg316119 := PrimCons(reg316113, reg316118)
reg316120 := PrimTail(V3873)
reg316121 := PrimTail(V3874)
reg316122 := PrimTail(reg316121)
reg316123 := PrimHead(reg316122)
reg316124 := __e.Call(__defun__shen_4patthyps, reg316120, reg316123, V3875)
__ctx.TailApply(__defun__adjoin, reg316119, reg316124)
return
} else {
reg316126 := MakeSymbol("shen.patthyps")
__ctx.TailApply(__defun__shen_4f__error, reg316126)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.patthyps", value: __defun__shen_4patthyps})

__defun__shen_4result_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3882 := __args[0]
_ = V3882
V3883 := __args[1]
_ = V3883
reg316128 := Nil;
reg316129 := PrimEqual(reg316128, V3882)
var reg316159 Obj
if reg316129 == True {
reg316130 := PrimIsPair(V3883)
var reg316154 Obj
if reg316130 == True {
reg316131 := MakeSymbol("-->")
reg316132 := PrimHead(V3883)
reg316133 := PrimEqual(reg316131, reg316132)
var reg316149 Obj
if reg316133 == True {
reg316134 := PrimTail(V3883)
reg316135 := PrimIsPair(reg316134)
var reg316144 Obj
if reg316135 == True {
reg316136 := Nil;
reg316137 := PrimTail(V3883)
reg316138 := PrimTail(reg316137)
reg316139 := PrimEqual(reg316136, reg316138)
var reg316142 Obj
if reg316139 == True {
reg316140 := True;
reg316142 = reg316140
} else {
reg316141 := False;
reg316142 = reg316141
}
reg316144 = reg316142
} else {
reg316143 := False;
reg316144 = reg316143
}
var reg316147 Obj
if reg316144 == True {
reg316145 := True;
reg316147 = reg316145
} else {
reg316146 := False;
reg316147 = reg316146
}
reg316149 = reg316147
} else {
reg316148 := False;
reg316149 = reg316148
}
var reg316152 Obj
if reg316149 == True {
reg316150 := True;
reg316152 = reg316150
} else {
reg316151 := False;
reg316152 = reg316151
}
reg316154 = reg316152
} else {
reg316153 := False;
reg316154 = reg316153
}
var reg316157 Obj
if reg316154 == True {
reg316155 := True;
reg316157 = reg316155
} else {
reg316156 := False;
reg316157 = reg316156
}
reg316159 = reg316157
} else {
reg316158 := False;
reg316159 = reg316158
}
if reg316159 == True {
reg316160 := PrimTail(V3883)
reg316161 := PrimHead(reg316160)
__ctx.Return(reg316161)
return
} else {
reg316162 := Nil;
reg316163 := PrimEqual(reg316162, V3882)
if reg316163 == True {
__ctx.Return(V3883)
return
} else {
reg316164 := PrimIsPair(V3882)
var reg316204 Obj
if reg316164 == True {
reg316165 := PrimIsPair(V3883)
var reg316199 Obj
if reg316165 == True {
reg316166 := PrimTail(V3883)
reg316167 := PrimIsPair(reg316166)
var reg316194 Obj
if reg316167 == True {
reg316168 := MakeSymbol("-->")
reg316169 := PrimTail(V3883)
reg316170 := PrimHead(reg316169)
reg316171 := PrimEqual(reg316168, reg316170)
var reg316189 Obj
if reg316171 == True {
reg316172 := PrimTail(V3883)
reg316173 := PrimTail(reg316172)
reg316174 := PrimIsPair(reg316173)
var reg316184 Obj
if reg316174 == True {
reg316175 := Nil;
reg316176 := PrimTail(V3883)
reg316177 := PrimTail(reg316176)
reg316178 := PrimTail(reg316177)
reg316179 := PrimEqual(reg316175, reg316178)
var reg316182 Obj
if reg316179 == True {
reg316180 := True;
reg316182 = reg316180
} else {
reg316181 := False;
reg316182 = reg316181
}
reg316184 = reg316182
} else {
reg316183 := False;
reg316184 = reg316183
}
var reg316187 Obj
if reg316184 == True {
reg316185 := True;
reg316187 = reg316185
} else {
reg316186 := False;
reg316187 = reg316186
}
reg316189 = reg316187
} else {
reg316188 := False;
reg316189 = reg316188
}
var reg316192 Obj
if reg316189 == True {
reg316190 := True;
reg316192 = reg316190
} else {
reg316191 := False;
reg316192 = reg316191
}
reg316194 = reg316192
} else {
reg316193 := False;
reg316194 = reg316193
}
var reg316197 Obj
if reg316194 == True {
reg316195 := True;
reg316197 = reg316195
} else {
reg316196 := False;
reg316197 = reg316196
}
reg316199 = reg316197
} else {
reg316198 := False;
reg316199 = reg316198
}
var reg316202 Obj
if reg316199 == True {
reg316200 := True;
reg316202 = reg316200
} else {
reg316201 := False;
reg316202 = reg316201
}
reg316204 = reg316202
} else {
reg316203 := False;
reg316204 = reg316203
}
if reg316204 == True {
reg316205 := PrimTail(V3882)
reg316206 := PrimTail(V3883)
reg316207 := PrimTail(reg316206)
reg316208 := PrimHead(reg316207)
__ctx.TailApply(__defun__shen_4result_1type, reg316205, reg316208)
return
} else {
reg316210 := MakeSymbol("shen.result-type")
__ctx.TailApply(__defun__shen_4f__error, reg316210)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.result-type", value: __defun__shen_4result_1type})

__defun__shen_4t_d_1patterns = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3889 := __args[0]
_ = V3889
V3890 := __args[1]
_ = V3890
V3891 := __args[2]
_ = V3891
V3892 := __args[3]
_ = V3892
V3893 := __args[4]
_ = V3893
reg316212 := __e.Call(__defun__shen_4lazyderef, V3889, V3892)
V3358 := reg316212
_ = V3358
reg316213 := Nil;
reg316214 := PrimEqual(reg316213, V3358)
var reg316218 Obj
if reg316214 == True {
reg316215 := __e.Call(__defun__shen_4incinfs)
_ = reg316215
reg316216 := __e.Call(__defun__thaw, V3893)
reg316218 = reg316216
} else {
reg316217 := False;
reg316218 = reg316217
}
Case := reg316218
_ = Case
reg316219 := False;
reg316220 := PrimEqual(Case, reg316219)
if reg316220 == True {
reg316221 := __e.Call(__defun__shen_4lazyderef, V3889, V3892)
V3359 := reg316221
_ = V3359
reg316222 := PrimIsPair(V3359)
if reg316222 == True {
reg316223 := PrimHead(V3359)
Pattern := reg316223
_ = Pattern
reg316224 := PrimTail(V3359)
Patterns := reg316224
_ = Patterns
reg316225 := __e.Call(__defun__shen_4lazyderef, V3890, V3892)
V3360 := reg316225
_ = V3360
reg316226 := PrimIsPair(V3360)
if reg316226 == True {
reg316227 := PrimHead(V3360)
A := reg316227
_ = A
reg316228 := PrimTail(V3360)
reg316229 := __e.Call(__defun__shen_4lazyderef, reg316228, V3892)
V3361 := reg316229
_ = V3361
reg316230 := PrimIsPair(V3361)
if reg316230 == True {
reg316231 := PrimHead(V3361)
reg316232 := __e.Call(__defun__shen_4lazyderef, reg316231, V3892)
V3362 := reg316232
_ = V3362
reg316233 := MakeSymbol("-->")
reg316234 := PrimEqual(reg316233, V3362)
if reg316234 == True {
reg316235 := PrimTail(V3361)
reg316236 := __e.Call(__defun__shen_4lazyderef, reg316235, V3892)
V3363 := reg316236
_ = V3363
reg316237 := PrimIsPair(V3363)
if reg316237 == True {
reg316238 := PrimHead(V3363)
B := reg316238
_ = B
reg316239 := PrimTail(V3363)
reg316240 := __e.Call(__defun__shen_4lazyderef, reg316239, V3892)
V3364 := reg316240
_ = V3364
reg316241 := Nil;
reg316242 := PrimEqual(reg316241, V3364)
if reg316242 == True {
reg316243 := __e.Call(__defun__shen_4incinfs)
_ = reg316243
reg316244 := MakeSymbol(":")
reg316245 := Nil;
reg316246 := PrimCons(A, reg316245)
reg316247 := PrimCons(reg316244, reg316246)
reg316248 := PrimCons(Pattern, reg316247)
reg316249 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4t_d_1patterns, Patterns, B, V3891, V3892, V3893)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d, reg316248, V3891, V3892, reg316249)
return
} else {
reg316252 := False;
__ctx.Return(reg316252)
return
}
} else {
reg316253 := False;
__ctx.Return(reg316253)
return
}
} else {
reg316254 := False;
__ctx.Return(reg316254)
return
}
} else {
reg316255 := False;
__ctx.Return(reg316255)
return
}
} else {
reg316256 := False;
__ctx.Return(reg316256)
return
}
} else {
reg316257 := False;
__ctx.Return(reg316257)
return
}
} else {
__ctx.Return(Case)
return
}
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-patterns", value: __defun__shen_4t_d_1patterns})

__defun__shen_4t_d_1action = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3899 := __args[0]
_ = V3899
V3900 := __args[1]
_ = V3900
V3901 := __args[2]
_ = V3901
V3902 := __args[3]
_ = V3902
V3903 := __args[4]
_ = V3903
reg316258 := __e.Call(__defun__shen_4catchpoint)
Throwcontrol := reg316258
_ = Throwcontrol
reg316259 := __e.Call(__defun__shen_4lazyderef, V3899, V3902)
V3335 := reg316259
_ = V3335
reg316260 := PrimIsPair(V3335)
var reg316307 Obj
if reg316260 == True {
reg316261 := PrimHead(V3335)
reg316262 := __e.Call(__defun__shen_4lazyderef, reg316261, V3902)
V3336 := reg316262
_ = V3336
reg316263 := MakeSymbol("where")
reg316264 := PrimEqual(reg316263, V3336)
var reg316305 Obj
if reg316264 == True {
reg316265 := PrimTail(V3335)
reg316266 := __e.Call(__defun__shen_4lazyderef, reg316265, V3902)
V3337 := reg316266
_ = V3337
reg316267 := PrimIsPair(V3337)
var reg316303 Obj
if reg316267 == True {
reg316268 := PrimHead(V3337)
P := reg316268
_ = P
reg316269 := PrimTail(V3337)
reg316270 := __e.Call(__defun__shen_4lazyderef, reg316269, V3902)
V3338 := reg316270
_ = V3338
reg316271 := PrimIsPair(V3338)
var reg316301 Obj
if reg316271 == True {
reg316272 := PrimHead(V3338)
Action := reg316272
_ = Action
reg316273 := PrimTail(V3338)
reg316274 := __e.Call(__defun__shen_4lazyderef, reg316273, V3902)
V3339 := reg316274
_ = V3339
reg316275 := Nil;
reg316276 := PrimEqual(reg316275, V3339)
var reg316299 Obj
if reg316276 == True {
reg316277 := __e.Call(__defun__shen_4incinfs)
_ = reg316277
reg316278 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316279 := MakeSymbol(":")
reg316280 := MakeSymbol("boolean")
reg316281 := Nil;
reg316282 := PrimCons(reg316280, reg316281)
reg316283 := PrimCons(reg316279, reg316282)
reg316284 := PrimCons(P, reg316283)
reg316285 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316286 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316287 := MakeSymbol(":")
reg316288 := MakeSymbol("verified")
reg316289 := Nil;
reg316290 := PrimCons(reg316288, reg316289)
reg316291 := PrimCons(reg316287, reg316290)
reg316292 := PrimCons(P, reg316291)
reg316293 := PrimCons(reg316292, V3901)
__ctx.TailApply(__defun__shen_4t_d_1action, Action, V3900, reg316293, V3902, V3903)
return
}, 0)
__ctx.TailApply(__defun__cut, Throwcontrol, V3902, reg316286)
return
}, 0)
__ctx.TailApply(__defun__shen_4t_d, reg316284, V3901, V3902, reg316285)
return
}, 0)
reg316297 := __e.Call(__defun__cut, Throwcontrol, V3902, reg316278)
reg316299 = reg316297
} else {
reg316298 := False;
reg316299 = reg316298
}
reg316301 = reg316299
} else {
reg316300 := False;
reg316301 = reg316300
}
reg316303 = reg316301
} else {
reg316302 := False;
reg316303 = reg316302
}
reg316305 = reg316303
} else {
reg316304 := False;
reg316305 = reg316304
}
reg316307 = reg316305
} else {
reg316306 := False;
reg316307 = reg316306
}
Case := reg316307
_ = Case
reg316308 := False;
reg316309 := PrimEqual(Case, reg316308)
var reg316445 Obj
if reg316309 == True {
reg316310 := __e.Call(__defun__shen_4lazyderef, V3899, V3902)
V3340 := reg316310
_ = V3340
reg316311 := PrimIsPair(V3340)
var reg316386 Obj
if reg316311 == True {
reg316312 := PrimHead(V3340)
reg316313 := __e.Call(__defun__shen_4lazyderef, reg316312, V3902)
V3341 := reg316313
_ = V3341
reg316314 := MakeSymbol("shen.choicepoint!")
reg316315 := PrimEqual(reg316314, V3341)
var reg316384 Obj
if reg316315 == True {
reg316316 := PrimTail(V3340)
reg316317 := __e.Call(__defun__shen_4lazyderef, reg316316, V3902)
V3342 := reg316317
_ = V3342
reg316318 := PrimIsPair(V3342)
var reg316382 Obj
if reg316318 == True {
reg316319 := PrimHead(V3342)
reg316320 := __e.Call(__defun__shen_4lazyderef, reg316319, V3902)
V3343 := reg316320
_ = V3343
reg316321 := PrimIsPair(V3343)
var reg316380 Obj
if reg316321 == True {
reg316322 := PrimHead(V3343)
reg316323 := __e.Call(__defun__shen_4lazyderef, reg316322, V3902)
V3344 := reg316323
_ = V3344
reg316324 := PrimIsPair(V3344)
var reg316378 Obj
if reg316324 == True {
reg316325 := PrimHead(V3344)
reg316326 := __e.Call(__defun__shen_4lazyderef, reg316325, V3902)
V3345 := reg316326
_ = V3345
reg316327 := MakeSymbol("fail-if")
reg316328 := PrimEqual(reg316327, V3345)
var reg316376 Obj
if reg316328 == True {
reg316329 := PrimTail(V3344)
reg316330 := __e.Call(__defun__shen_4lazyderef, reg316329, V3902)
V3346 := reg316330
_ = V3346
reg316331 := PrimIsPair(V3346)
var reg316374 Obj
if reg316331 == True {
reg316332 := PrimHead(V3346)
F := reg316332
_ = F
reg316333 := PrimTail(V3346)
reg316334 := __e.Call(__defun__shen_4lazyderef, reg316333, V3902)
V3347 := reg316334
_ = V3347
reg316335 := Nil;
reg316336 := PrimEqual(reg316335, V3347)
var reg316372 Obj
if reg316336 == True {
reg316337 := PrimTail(V3343)
reg316338 := __e.Call(__defun__shen_4lazyderef, reg316337, V3902)
V3348 := reg316338
_ = V3348
reg316339 := PrimIsPair(V3348)
var reg316370 Obj
if reg316339 == True {
reg316340 := PrimHead(V3348)
Action := reg316340
_ = Action
reg316341 := PrimTail(V3348)
reg316342 := __e.Call(__defun__shen_4lazyderef, reg316341, V3902)
V3349 := reg316342
_ = V3349
reg316343 := Nil;
reg316344 := PrimEqual(reg316343, V3349)
var reg316368 Obj
if reg316344 == True {
reg316345 := PrimTail(V3342)
reg316346 := __e.Call(__defun__shen_4lazyderef, reg316345, V3902)
V3350 := reg316346
_ = V3350
reg316347 := Nil;
reg316348 := PrimEqual(reg316347, V3350)
var reg316366 Obj
if reg316348 == True {
reg316349 := __e.Call(__defun__shen_4incinfs)
_ = reg316349
reg316350 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316351 := MakeSymbol("where")
reg316352 := MakeSymbol("not")
reg316353 := Nil;
reg316354 := PrimCons(Action, reg316353)
reg316355 := PrimCons(F, reg316354)
reg316356 := Nil;
reg316357 := PrimCons(reg316355, reg316356)
reg316358 := PrimCons(reg316352, reg316357)
reg316359 := Nil;
reg316360 := PrimCons(Action, reg316359)
reg316361 := PrimCons(reg316358, reg316360)
reg316362 := PrimCons(reg316351, reg316361)
__ctx.TailApply(__defun__shen_4t_d_1action, reg316362, V3900, V3901, V3902, V3903)
return
}, 0)
reg316364 := __e.Call(__defun__cut, Throwcontrol, V3902, reg316350)
reg316366 = reg316364
} else {
reg316365 := False;
reg316366 = reg316365
}
reg316368 = reg316366
} else {
reg316367 := False;
reg316368 = reg316367
}
reg316370 = reg316368
} else {
reg316369 := False;
reg316370 = reg316369
}
reg316372 = reg316370
} else {
reg316371 := False;
reg316372 = reg316371
}
reg316374 = reg316372
} else {
reg316373 := False;
reg316374 = reg316373
}
reg316376 = reg316374
} else {
reg316375 := False;
reg316376 = reg316375
}
reg316378 = reg316376
} else {
reg316377 := False;
reg316378 = reg316377
}
reg316380 = reg316378
} else {
reg316379 := False;
reg316380 = reg316379
}
reg316382 = reg316380
} else {
reg316381 := False;
reg316382 = reg316381
}
reg316384 = reg316382
} else {
reg316383 := False;
reg316384 = reg316383
}
reg316386 = reg316384
} else {
reg316385 := False;
reg316386 = reg316385
}
Case := reg316386
_ = Case
reg316387 := False;
reg316388 := PrimEqual(Case, reg316387)
var reg316444 Obj
if reg316388 == True {
reg316389 := __e.Call(__defun__shen_4lazyderef, V3899, V3902)
V3351 := reg316389
_ = V3351
reg316390 := PrimIsPair(V3351)
var reg316433 Obj
if reg316390 == True {
reg316391 := PrimHead(V3351)
reg316392 := __e.Call(__defun__shen_4lazyderef, reg316391, V3902)
V3352 := reg316392
_ = V3352
reg316393 := MakeSymbol("shen.choicepoint!")
reg316394 := PrimEqual(reg316393, V3352)
var reg316431 Obj
if reg316394 == True {
reg316395 := PrimTail(V3351)
reg316396 := __e.Call(__defun__shen_4lazyderef, reg316395, V3902)
V3353 := reg316396
_ = V3353
reg316397 := PrimIsPair(V3353)
var reg316429 Obj
if reg316397 == True {
reg316398 := PrimHead(V3353)
Action := reg316398
_ = Action
reg316399 := PrimTail(V3353)
reg316400 := __e.Call(__defun__shen_4lazyderef, reg316399, V3902)
V3354 := reg316400
_ = V3354
reg316401 := Nil;
reg316402 := PrimEqual(reg316401, V3354)
var reg316427 Obj
if reg316402 == True {
reg316403 := __e.Call(__defun__shen_4incinfs)
_ = reg316403
reg316404 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316405 := MakeSymbol("where")
reg316406 := MakeSymbol("not")
reg316407 := MakeSymbol("=")
reg316408 := Nil;
reg316409 := PrimCons(Action, reg316408)
reg316410 := PrimCons(reg316407, reg316409)
reg316411 := MakeSymbol("fail")
reg316412 := Nil;
reg316413 := PrimCons(reg316411, reg316412)
reg316414 := Nil;
reg316415 := PrimCons(reg316413, reg316414)
reg316416 := PrimCons(reg316410, reg316415)
reg316417 := Nil;
reg316418 := PrimCons(reg316416, reg316417)
reg316419 := PrimCons(reg316406, reg316418)
reg316420 := Nil;
reg316421 := PrimCons(Action, reg316420)
reg316422 := PrimCons(reg316419, reg316421)
reg316423 := PrimCons(reg316405, reg316422)
__ctx.TailApply(__defun__shen_4t_d_1action, reg316423, V3900, V3901, V3902, V3903)
return
}, 0)
reg316425 := __e.Call(__defun__cut, Throwcontrol, V3902, reg316404)
reg316427 = reg316425
} else {
reg316426 := False;
reg316427 = reg316426
}
reg316429 = reg316427
} else {
reg316428 := False;
reg316429 = reg316428
}
reg316431 = reg316429
} else {
reg316430 := False;
reg316431 = reg316430
}
reg316433 = reg316431
} else {
reg316432 := False;
reg316433 = reg316432
}
Case := reg316433
_ = Case
reg316434 := False;
reg316435 := PrimEqual(Case, reg316434)
var reg316443 Obj
if reg316435 == True {
reg316436 := __e.Call(__defun__shen_4incinfs)
_ = reg316436
reg316437 := MakeSymbol(":")
reg316438 := Nil;
reg316439 := PrimCons(V3900, reg316438)
reg316440 := PrimCons(reg316437, reg316439)
reg316441 := PrimCons(V3899, reg316440)
reg316442 := __e.Call(__defun__shen_4t_d, reg316441, V3901, V3902, V3903)
reg316443 = reg316442
} else {
reg316443 = Case
}
reg316444 = reg316443
} else {
reg316444 = Case
}
reg316445 = reg316444
} else {
reg316445 = Case
}
__ctx.TailApply(__defun__shen_4cutpoint, Throwcontrol, reg316445)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "shen.t*-action", value: __defun__shen_4t_d_1action})

__defun__findall = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3909 := __args[0]
_ = V3909
V3910 := __args[1]
_ = V3910
V3911 := __args[2]
_ = V3911
V3912 := __args[3]
_ = V3912
V3913 := __args[4]
_ = V3913
reg316447 := __e.Call(__defun__shen_4newpv, V3912)
B := reg316447
_ = B
reg316448 := __e.Call(__defun__shen_4newpv, V3912)
A := reg316448
_ = A
reg316449 := __e.Call(__defun__shen_4incinfs)
_ = reg316449
reg316450 := MakeSymbol("shen.a")
reg316451 := __e.Call(__defun__gensym, reg316450)
reg316452 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316453 := __e.Call(__defun__shen_4lazyderef, A, V3912)
reg316454 := Nil;
reg316455 := PrimSet(reg316453, reg316454)
reg316456 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4findallhelp, V3909, V3910, V3911, A, V3912, V3913)
return
}, 0)
__ctx.TailApply(__defun__bind, B, reg316455, V3912, reg316456)
return
}, 0)
__ctx.TailApply(__defun__bind, A, reg316451, V3912, reg316452)
return
}, 5)
__initDefs = append(__initDefs, defType{name: "findall", value: __defun__findall})

__defun__shen_4findallhelp = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3920 := __args[0]
_ = V3920
V3921 := __args[1]
_ = V3921
V3922 := __args[2]
_ = V3922
V3923 := __args[3]
_ = V3923
V3924 := __args[4]
_ = V3924
V3925 := __args[5]
_ = V3925
reg316460 := __e.Call(__defun__shen_4incinfs)
_ = reg316460
reg316461 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316462 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316463 := False;
__ctx.TailApply(__defun__fwhen, reg316463, V3924, V3925)
return
}, 0)
__ctx.TailApply(__defun__shen_4remember, V3923, V3920, V3924, reg316462)
return
}, 0)
reg316466 := __e.Call(__defun__call, V3921, V3924, reg316461)
Case := reg316466
_ = Case
reg316467 := False;
reg316468 := PrimEqual(Case, reg316467)
if reg316468 == True {
reg316469 := __e.Call(__defun__shen_4incinfs)
_ = reg316469
reg316470 := __e.Call(__defun__shen_4lazyderef, V3923, V3924)
reg316471 := PrimValue(reg316470)
__ctx.TailApply(__defun__bind, V3922, reg316471, V3924, V3925)
return
} else {
__ctx.Return(Case)
return
}
}, 6)
__initDefs = append(__initDefs, defType{name: "shen.findallhelp", value: __defun__shen_4findallhelp})

__defun__shen_4remember = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3930 := __args[0]
_ = V3930
V3931 := __args[1]
_ = V3931
V3932 := __args[2]
_ = V3932
V3933 := __args[3]
_ = V3933
reg316473 := __e.Call(__defun__shen_4newpv, V3932)
B := reg316473
_ = B
reg316474 := __e.Call(__defun__shen_4incinfs)
_ = reg316474
reg316475 := __e.Call(__defun__shen_4deref, V3930, V3932)
reg316476 := __e.Call(__defun__shen_4deref, V3931, V3932)
reg316477 := __e.Call(__defun__shen_4deref, V3930, V3932)
reg316478 := PrimValue(reg316477)
reg316479 := PrimCons(reg316476, reg316478)
reg316480 := PrimSet(reg316475, reg316479)
__ctx.TailApply(__defun__bind, B, reg316480, V3932, V3933)
return
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.remember", value: __defun__shen_4remember})

}
