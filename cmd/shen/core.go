package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4shen_1_6kl Obj // shen.shen->kl
var __defun__shen_4shen_1syntax_1error Obj // shen.shen-syntax-error
var __defun__shen_4_5define_6 Obj // shen.<define>
var __defun__shen_4_5name_6 Obj // shen.<name>
var __defun__shen_4sysfunc_2 Obj // shen.sysfunc?
var __defun__shen_4_5signature_6 Obj // shen.<signature>
var __defun__shen_4curry_1type Obj // shen.curry-type
var __defun__shen_4active_1cons Obj // shen.active-cons
var __defun__shen_4curry_1type_1h Obj // shen.curry-type-h
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
var __defun__shen_4clash_2 Obj // shen.clash?
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
reg744 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
__ctx.Return(reg744)
return
}, 0))
__defun__shen_4shen_1_6kl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V128 := __args[0]
_ = V128
V129 := __args[1]
_ = V129
reg745 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5define_6, X)
return
}, 1)
reg747 := PrimCons(V128, V129)
reg748 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4shen_1syntax_1error, V128, X)
return
}, 1)
__ctx.TailApply(__defun__compile, reg745, reg747, reg748)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.shen->kl", value: __defun__shen_4shen_1_6kl})

__defun__shen_4shen_1syntax_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V136 := __args[0]
_ = V136
V137 := __args[1]
_ = V137
reg751 := PrimIsPair(V137)
if reg751 == True {
reg752 := MakeString("syntax error in ")
reg753 := MakeString(" here:\n\n ")
reg754 := MakeNumber(50)
reg755 := PrimHead(V137)
reg756 := __e.Call(__defun__shen_4next_150, reg754, reg755)
reg757 := MakeString("\n")
reg758 := MakeSymbol("shen.a")
reg759 := __e.Call(__defun__shen_4app, reg756, reg757, reg758)
reg760 := PrimStringConcat(reg753, reg759)
reg761 := MakeSymbol("shen.a")
reg762 := __e.Call(__defun__shen_4app, V136, reg760, reg761)
reg763 := PrimStringConcat(reg752, reg762)
reg764 := PrimSimpleError(reg763)
__ctx.Return(reg764)
return
} else {
reg765 := MakeString("syntax error in ")
reg766 := MakeString("\n")
reg767 := MakeSymbol("shen.a")
reg768 := __e.Call(__defun__shen_4app, V136, reg766, reg767)
reg769 := PrimStringConcat(reg765, reg768)
reg770 := PrimSimpleError(reg769)
__ctx.Return(reg770)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.shen-syntax-error", value: __defun__shen_4shen_1syntax_1error})

__defun__shen_4_5define_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V139 := __args[0]
_ = V139
reg771 := __e.Call(__defun__shen_4_5name_6, V139)
Parse__shen_4_5name_6 := reg771
_ = Parse__shen_4_5name_6
reg772 := __e.Call(__defun__fail)
reg773 := PrimEqual(reg772, Parse__shen_4_5name_6)
reg774 := PrimNot(reg773)
var reg793 Obj
if reg774 == True {
reg775 := __e.Call(__defun__shen_4_5signature_6, Parse__shen_4_5name_6)
Parse__shen_4_5signature_6 := reg775
_ = Parse__shen_4_5signature_6
reg776 := __e.Call(__defun__fail)
reg777 := PrimEqual(reg776, Parse__shen_4_5signature_6)
reg778 := PrimNot(reg777)
var reg791 Obj
if reg778 == True {
reg779 := __e.Call(__defun__shen_4_5rules_6, Parse__shen_4_5signature_6)
Parse__shen_4_5rules_6 := reg779
_ = Parse__shen_4_5rules_6
reg780 := __e.Call(__defun__fail)
reg781 := PrimEqual(reg780, Parse__shen_4_5rules_6)
reg782 := PrimNot(reg781)
var reg789 Obj
if reg782 == True {
reg783 := PrimHead(Parse__shen_4_5rules_6)
reg784 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5name_6)
reg785 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rules_6)
reg786 := __e.Call(__defun__shen_4compile__to__machine__code, reg784, reg785)
reg787 := __e.Call(__defun__shen_4pair, reg783, reg786)
reg789 = reg787
} else {
reg788 := __e.Call(__defun__fail)
reg789 = reg788
}
reg791 = reg789
} else {
reg790 := __e.Call(__defun__fail)
reg791 = reg790
}
reg793 = reg791
} else {
reg792 := __e.Call(__defun__fail)
reg793 = reg792
}
YaccParse := reg793
_ = YaccParse
reg794 := __e.Call(__defun__fail)
reg795 := PrimEqual(YaccParse, reg794)
if reg795 == True {
reg796 := __e.Call(__defun__shen_4_5name_6, V139)
Parse__shen_4_5name_6 := reg796
_ = Parse__shen_4_5name_6
reg797 := __e.Call(__defun__fail)
reg798 := PrimEqual(reg797, Parse__shen_4_5name_6)
reg799 := PrimNot(reg798)
if reg799 == True {
reg800 := __e.Call(__defun__shen_4_5rules_6, Parse__shen_4_5name_6)
Parse__shen_4_5rules_6 := reg800
_ = Parse__shen_4_5rules_6
reg801 := __e.Call(__defun__fail)
reg802 := PrimEqual(reg801, Parse__shen_4_5rules_6)
reg803 := PrimNot(reg802)
if reg803 == True {
reg804 := PrimHead(Parse__shen_4_5rules_6)
reg805 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5name_6)
reg806 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rules_6)
reg807 := __e.Call(__defun__shen_4compile__to__machine__code, reg805, reg806)
__ctx.TailApply(__defun__shen_4pair, reg804, reg807)
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
V141 := __args[0]
_ = V141
reg811 := PrimHead(V141)
reg812 := PrimIsPair(reg811)
if reg812 == True {
reg813 := __e.Call(__defun__shen_4hdhd, V141)
Parse__X := reg813
_ = Parse__X
reg814 := __e.Call(__defun__shen_4tlhd, V141)
reg815 := __e.Call(__defun__shen_4hdtl, V141)
reg816 := __e.Call(__defun__shen_4pair, reg814, reg815)
reg817 := PrimHead(reg816)
reg818 := PrimIsSymbol(Parse__X)
var reg825 Obj
if reg818 == True {
reg819 := __e.Call(__defun__shen_4sysfunc_2, Parse__X)
reg820 := PrimNot(reg819)
var reg823 Obj
if reg820 == True {
reg821 := True;
reg823 = reg821
} else {
reg822 := False;
reg823 = reg822
}
reg825 = reg823
} else {
reg824 := False;
reg825 = reg824
}
var reg830 Obj
if reg825 == True {
reg830 = Parse__X
} else {
reg826 := MakeString(" is not a legitimate function name.\n")
reg827 := MakeSymbol("shen.a")
reg828 := __e.Call(__defun__shen_4app, Parse__X, reg826, reg827)
reg829 := PrimSimpleError(reg828)
reg830 = reg829
}
__ctx.TailApply(__defun__shen_4pair, reg817, reg830)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<name>", value: __defun__shen_4_5name_6})

__defun__shen_4sysfunc_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V143 := __args[0]
_ = V143
reg833 := MakeString("shen")
reg834 := PrimIntern(reg833)
reg835 := MakeSymbol("shen.external-symbols")
reg836 := MakeSymbol("*property-vector*")
reg837 := PrimValue(reg836)
reg838 := __e.Call(__defun__get, reg834, reg835, reg837)
__ctx.TailApply(__defun__element_2, V143, reg838)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sysfunc?", value: __defun__shen_4sysfunc_2})

__defun__shen_4_5signature_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V147 := __args[0]
_ = V147
reg840 := PrimHead(V147)
reg841 := PrimIsPair(reg840)
var reg849 Obj
if reg841 == True {
reg842 := MakeSymbol("{")
reg843 := __e.Call(__defun__shen_4hdhd, V147)
reg844 := PrimEqual(reg842, reg843)
var reg847 Obj
if reg844 == True {
reg845 := True;
reg847 = reg845
} else {
reg846 := False;
reg847 = reg846
}
reg849 = reg847
} else {
reg848 := False;
reg849 = reg848
}
if reg849 == True {
reg850 := __e.Call(__defun__shen_4tlhd, V147)
reg851 := __e.Call(__defun__shen_4hdtl, V147)
reg852 := __e.Call(__defun__shen_4pair, reg850, reg851)
NewStream144 := reg852
_ = NewStream144
reg853 := __e.Call(__defun__shen_4_5signature_1help_6, NewStream144)
Parse__shen_4_5signature_1help_6 := reg853
_ = Parse__shen_4_5signature_1help_6
reg854 := __e.Call(__defun__fail)
reg855 := PrimEqual(reg854, Parse__shen_4_5signature_1help_6)
reg856 := PrimNot(reg855)
if reg856 == True {
reg857 := PrimHead(Parse__shen_4_5signature_1help_6)
reg858 := PrimIsPair(reg857)
var reg866 Obj
if reg858 == True {
reg859 := MakeSymbol("}")
reg860 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5signature_1help_6)
reg861 := PrimEqual(reg859, reg860)
var reg864 Obj
if reg861 == True {
reg862 := True;
reg864 = reg862
} else {
reg863 := False;
reg864 = reg863
}
reg866 = reg864
} else {
reg865 := False;
reg866 = reg865
}
if reg866 == True {
reg867 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5signature_1help_6)
reg868 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_1help_6)
reg869 := __e.Call(__defun__shen_4pair, reg867, reg868)
NewStream145 := reg869
_ = NewStream145
reg870 := PrimHead(NewStream145)
reg871 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_1help_6)
reg872 := __e.Call(__defun__shen_4curry_1type, reg871)
reg873 := __e.Call(__defun__shen_4demodulate, reg872)
__ctx.TailApply(__defun__shen_4pair, reg870, reg873)
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
V149 := __args[0]
_ = V149
reg878 := __e.Call(__defun__shen_4curry_1type_1h, V149)
__ctx.TailApply(__defun__shen_4active_1cons, reg878)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.curry-type", value: __defun__shen_4curry_1type})

__defun__shen_4active_1cons = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V151 := __args[0]
_ = V151
reg880 := PrimIsPair(V151)
var reg914 Obj
if reg880 == True {
reg881 := PrimTail(V151)
reg882 := PrimIsPair(reg881)
var reg909 Obj
if reg882 == True {
reg883 := PrimTail(V151)
reg884 := PrimTail(reg883)
reg885 := PrimIsPair(reg884)
var reg904 Obj
if reg885 == True {
reg886 := Nil;
reg887 := PrimTail(V151)
reg888 := PrimTail(reg887)
reg889 := PrimTail(reg888)
reg890 := PrimEqual(reg886, reg889)
var reg899 Obj
if reg890 == True {
reg891 := PrimTail(V151)
reg892 := PrimHead(reg891)
reg893 := MakeSymbol("bar!")
reg894 := PrimEqual(reg892, reg893)
var reg897 Obj
if reg894 == True {
reg895 := True;
reg897 = reg895
} else {
reg896 := False;
reg897 = reg896
}
reg899 = reg897
} else {
reg898 := False;
reg899 = reg898
}
var reg902 Obj
if reg899 == True {
reg900 := True;
reg902 = reg900
} else {
reg901 := False;
reg902 = reg901
}
reg904 = reg902
} else {
reg903 := False;
reg904 = reg903
}
var reg907 Obj
if reg904 == True {
reg905 := True;
reg907 = reg905
} else {
reg906 := False;
reg907 = reg906
}
reg909 = reg907
} else {
reg908 := False;
reg909 = reg908
}
var reg912 Obj
if reg909 == True {
reg910 := True;
reg912 = reg910
} else {
reg911 := False;
reg912 = reg911
}
reg914 = reg912
} else {
reg913 := False;
reg914 = reg913
}
if reg914 == True {
reg915 := PrimHead(V151)
reg916 := __e.Call(__defun__shen_4active_1cons, reg915)
reg917 := PrimTail(V151)
reg918 := PrimTail(reg917)
reg919 := PrimHead(reg918)
reg920 := __e.Call(__defun__shen_4active_1cons, reg919)
reg921 := PrimCons(reg916, reg920)
__ctx.Return(reg921)
return
} else {
reg922 := PrimIsPair(V151)
if reg922 == True {
reg923 := PrimHead(V151)
reg924 := __e.Call(__defun__shen_4active_1cons, reg923)
reg925 := PrimTail(V151)
reg926 := __e.Call(__defun__shen_4active_1cons, reg925)
reg927 := PrimCons(reg924, reg926)
__ctx.Return(reg927)
return
} else {
__ctx.Return(V151)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.active-cons", value: __defun__shen_4active_1cons})

__defun__shen_4curry_1type_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V153 := __args[0]
_ = V153
reg928 := PrimIsPair(V153)
var reg972 Obj
if reg928 == True {
reg929 := PrimTail(V153)
reg930 := PrimIsPair(reg929)
var reg967 Obj
if reg930 == True {
reg931 := MakeSymbol("-->")
reg932 := PrimTail(V153)
reg933 := PrimHead(reg932)
reg934 := PrimEqual(reg931, reg933)
var reg962 Obj
if reg934 == True {
reg935 := PrimTail(V153)
reg936 := PrimTail(reg935)
reg937 := PrimIsPair(reg936)
var reg957 Obj
if reg937 == True {
reg938 := PrimTail(V153)
reg939 := PrimTail(reg938)
reg940 := PrimTail(reg939)
reg941 := PrimIsPair(reg940)
var reg952 Obj
if reg941 == True {
reg942 := MakeSymbol("-->")
reg943 := PrimTail(V153)
reg944 := PrimTail(reg943)
reg945 := PrimTail(reg944)
reg946 := PrimHead(reg945)
reg947 := PrimEqual(reg942, reg946)
var reg950 Obj
if reg947 == True {
reg948 := True;
reg950 = reg948
} else {
reg949 := False;
reg950 = reg949
}
reg952 = reg950
} else {
reg951 := False;
reg952 = reg951
}
var reg955 Obj
if reg952 == True {
reg953 := True;
reg955 = reg953
} else {
reg954 := False;
reg955 = reg954
}
reg957 = reg955
} else {
reg956 := False;
reg957 = reg956
}
var reg960 Obj
if reg957 == True {
reg958 := True;
reg960 = reg958
} else {
reg959 := False;
reg960 = reg959
}
reg962 = reg960
} else {
reg961 := False;
reg962 = reg961
}
var reg965 Obj
if reg962 == True {
reg963 := True;
reg965 = reg963
} else {
reg964 := False;
reg965 = reg964
}
reg967 = reg965
} else {
reg966 := False;
reg967 = reg966
}
var reg970 Obj
if reg967 == True {
reg968 := True;
reg970 = reg968
} else {
reg969 := False;
reg970 = reg969
}
reg972 = reg970
} else {
reg971 := False;
reg972 = reg971
}
if reg972 == True {
reg973 := PrimHead(V153)
reg974 := MakeSymbol("-->")
reg975 := PrimTail(V153)
reg976 := PrimTail(reg975)
reg977 := Nil;
reg978 := PrimCons(reg976, reg977)
reg979 := PrimCons(reg974, reg978)
reg980 := PrimCons(reg973, reg979)
__ctx.TailApply(__defun__shen_4curry_1type_1h, reg980)
return
} else {
reg982 := PrimIsPair(V153)
var reg1026 Obj
if reg982 == True {
reg983 := PrimTail(V153)
reg984 := PrimIsPair(reg983)
var reg1021 Obj
if reg984 == True {
reg985 := MakeSymbol("*")
reg986 := PrimTail(V153)
reg987 := PrimHead(reg986)
reg988 := PrimEqual(reg985, reg987)
var reg1016 Obj
if reg988 == True {
reg989 := PrimTail(V153)
reg990 := PrimTail(reg989)
reg991 := PrimIsPair(reg990)
var reg1011 Obj
if reg991 == True {
reg992 := PrimTail(V153)
reg993 := PrimTail(reg992)
reg994 := PrimTail(reg993)
reg995 := PrimIsPair(reg994)
var reg1006 Obj
if reg995 == True {
reg996 := MakeSymbol("*")
reg997 := PrimTail(V153)
reg998 := PrimTail(reg997)
reg999 := PrimTail(reg998)
reg1000 := PrimHead(reg999)
reg1001 := PrimEqual(reg996, reg1000)
var reg1004 Obj
if reg1001 == True {
reg1002 := True;
reg1004 = reg1002
} else {
reg1003 := False;
reg1004 = reg1003
}
reg1006 = reg1004
} else {
reg1005 := False;
reg1006 = reg1005
}
var reg1009 Obj
if reg1006 == True {
reg1007 := True;
reg1009 = reg1007
} else {
reg1008 := False;
reg1009 = reg1008
}
reg1011 = reg1009
} else {
reg1010 := False;
reg1011 = reg1010
}
var reg1014 Obj
if reg1011 == True {
reg1012 := True;
reg1014 = reg1012
} else {
reg1013 := False;
reg1014 = reg1013
}
reg1016 = reg1014
} else {
reg1015 := False;
reg1016 = reg1015
}
var reg1019 Obj
if reg1016 == True {
reg1017 := True;
reg1019 = reg1017
} else {
reg1018 := False;
reg1019 = reg1018
}
reg1021 = reg1019
} else {
reg1020 := False;
reg1021 = reg1020
}
var reg1024 Obj
if reg1021 == True {
reg1022 := True;
reg1024 = reg1022
} else {
reg1023 := False;
reg1024 = reg1023
}
reg1026 = reg1024
} else {
reg1025 := False;
reg1026 = reg1025
}
if reg1026 == True {
reg1027 := PrimHead(V153)
reg1028 := MakeSymbol("*")
reg1029 := PrimTail(V153)
reg1030 := PrimTail(reg1029)
reg1031 := Nil;
reg1032 := PrimCons(reg1030, reg1031)
reg1033 := PrimCons(reg1028, reg1032)
reg1034 := PrimCons(reg1027, reg1033)
__ctx.TailApply(__defun__shen_4curry_1type_1h, reg1034)
return
} else {
reg1036 := PrimIsPair(V153)
if reg1036 == True {
reg1037 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4curry_1type_1h, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg1037, V153)
return
} else {
__ctx.Return(V153)
return
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.curry-type-h", value: __defun__shen_4curry_1type_1h})

__defun__shen_4_5signature_1help_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V155 := __args[0]
_ = V155
reg1040 := PrimHead(V155)
reg1041 := PrimIsPair(reg1040)
var reg1066 Obj
if reg1041 == True {
reg1042 := __e.Call(__defun__shen_4hdhd, V155)
Parse__X := reg1042
_ = Parse__X
reg1043 := __e.Call(__defun__shen_4tlhd, V155)
reg1044 := __e.Call(__defun__shen_4hdtl, V155)
reg1045 := __e.Call(__defun__shen_4pair, reg1043, reg1044)
reg1046 := __e.Call(__defun__shen_4_5signature_1help_6, reg1045)
Parse__shen_4_5signature_1help_6 := reg1046
_ = Parse__shen_4_5signature_1help_6
reg1047 := __e.Call(__defun__fail)
reg1048 := PrimEqual(reg1047, Parse__shen_4_5signature_1help_6)
reg1049 := PrimNot(reg1048)
var reg1064 Obj
if reg1049 == True {
reg1050 := MakeSymbol("{")
reg1051 := MakeSymbol("}")
reg1052 := Nil;
reg1053 := PrimCons(reg1051, reg1052)
reg1054 := PrimCons(reg1050, reg1053)
reg1055 := __e.Call(__defun__element_2, Parse__X, reg1054)
reg1056 := PrimNot(reg1055)
var reg1062 Obj
if reg1056 == True {
reg1057 := PrimHead(Parse__shen_4_5signature_1help_6)
reg1058 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5signature_1help_6)
reg1059 := PrimCons(Parse__X, reg1058)
reg1060 := __e.Call(__defun__shen_4pair, reg1057, reg1059)
reg1062 = reg1060
} else {
reg1061 := __e.Call(__defun__fail)
reg1062 = reg1061
}
reg1064 = reg1062
} else {
reg1063 := __e.Call(__defun__fail)
reg1064 = reg1063
}
reg1066 = reg1064
} else {
reg1065 := __e.Call(__defun__fail)
reg1066 = reg1065
}
YaccParse := reg1066
_ = YaccParse
reg1067 := __e.Call(__defun__fail)
reg1068 := PrimEqual(YaccParse, reg1067)
if reg1068 == True {
reg1069 := __e.Call(__defun___5e_6, V155)
Parse___5e_6 := reg1069
_ = Parse___5e_6
reg1070 := __e.Call(__defun__fail)
reg1071 := PrimEqual(reg1070, Parse___5e_6)
reg1072 := PrimNot(reg1071)
if reg1072 == True {
reg1073 := PrimHead(Parse___5e_6)
reg1074 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg1073, reg1074)
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
V157 := __args[0]
_ = V157
reg1077 := __e.Call(__defun__shen_4_5rule_6, V157)
Parse__shen_4_5rule_6 := reg1077
_ = Parse__shen_4_5rule_6
reg1078 := __e.Call(__defun__fail)
reg1079 := PrimEqual(reg1078, Parse__shen_4_5rule_6)
reg1080 := PrimNot(reg1079)
var reg1094 Obj
if reg1080 == True {
reg1081 := __e.Call(__defun__shen_4_5rules_6, Parse__shen_4_5rule_6)
Parse__shen_4_5rules_6 := reg1081
_ = Parse__shen_4_5rules_6
reg1082 := __e.Call(__defun__fail)
reg1083 := PrimEqual(reg1082, Parse__shen_4_5rules_6)
reg1084 := PrimNot(reg1083)
var reg1092 Obj
if reg1084 == True {
reg1085 := PrimHead(Parse__shen_4_5rules_6)
reg1086 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg1087 := __e.Call(__defun__shen_4linearise, reg1086)
reg1088 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rules_6)
reg1089 := PrimCons(reg1087, reg1088)
reg1090 := __e.Call(__defun__shen_4pair, reg1085, reg1089)
reg1092 = reg1090
} else {
reg1091 := __e.Call(__defun__fail)
reg1092 = reg1091
}
reg1094 = reg1092
} else {
reg1093 := __e.Call(__defun__fail)
reg1094 = reg1093
}
YaccParse := reg1094
_ = YaccParse
reg1095 := __e.Call(__defun__fail)
reg1096 := PrimEqual(YaccParse, reg1095)
if reg1096 == True {
reg1097 := __e.Call(__defun__shen_4_5rule_6, V157)
Parse__shen_4_5rule_6 := reg1097
_ = Parse__shen_4_5rule_6
reg1098 := __e.Call(__defun__fail)
reg1099 := PrimEqual(reg1098, Parse__shen_4_5rule_6)
reg1100 := PrimNot(reg1099)
if reg1100 == True {
reg1101 := PrimHead(Parse__shen_4_5rule_6)
reg1102 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5rule_6)
reg1103 := __e.Call(__defun__shen_4linearise, reg1102)
reg1104 := Nil;
reg1105 := PrimCons(reg1103, reg1104)
__ctx.TailApply(__defun__shen_4pair, reg1101, reg1105)
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
V165 := __args[0]
_ = V165
reg1108 := __e.Call(__defun__shen_4_5patterns_6, V165)
Parse__shen_4_5patterns_6 := reg1108
_ = Parse__shen_4_5patterns_6
reg1109 := __e.Call(__defun__fail)
reg1110 := PrimEqual(reg1109, Parse__shen_4_5patterns_6)
reg1111 := PrimNot(reg1110)
var reg1168 Obj
if reg1111 == True {
reg1112 := PrimHead(Parse__shen_4_5patterns_6)
reg1113 := PrimIsPair(reg1112)
var reg1121 Obj
if reg1113 == True {
reg1114 := MakeSymbol("->")
reg1115 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5patterns_6)
reg1116 := PrimEqual(reg1114, reg1115)
var reg1119 Obj
if reg1116 == True {
reg1117 := True;
reg1119 = reg1117
} else {
reg1118 := False;
reg1119 = reg1118
}
reg1121 = reg1119
} else {
reg1120 := False;
reg1121 = reg1120
}
var reg1166 Obj
if reg1121 == True {
reg1122 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5patterns_6)
reg1123 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1124 := __e.Call(__defun__shen_4pair, reg1122, reg1123)
NewStream158 := reg1124
_ = NewStream158
reg1125 := __e.Call(__defun__shen_4_5action_6, NewStream158)
Parse__shen_4_5action_6 := reg1125
_ = Parse__shen_4_5action_6
reg1126 := __e.Call(__defun__fail)
reg1127 := PrimEqual(reg1126, Parse__shen_4_5action_6)
reg1128 := PrimNot(reg1127)
var reg1164 Obj
if reg1128 == True {
reg1129 := PrimHead(Parse__shen_4_5action_6)
reg1130 := PrimIsPair(reg1129)
var reg1138 Obj
if reg1130 == True {
reg1131 := MakeSymbol("where")
reg1132 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5action_6)
reg1133 := PrimEqual(reg1131, reg1132)
var reg1136 Obj
if reg1133 == True {
reg1134 := True;
reg1136 = reg1134
} else {
reg1135 := False;
reg1136 = reg1135
}
reg1138 = reg1136
} else {
reg1137 := False;
reg1138 = reg1137
}
var reg1162 Obj
if reg1138 == True {
reg1139 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5action_6)
reg1140 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg1141 := __e.Call(__defun__shen_4pair, reg1139, reg1140)
NewStream159 := reg1141
_ = NewStream159
reg1142 := __e.Call(__defun__shen_4_5guard_6, NewStream159)
Parse__shen_4_5guard_6 := reg1142
_ = Parse__shen_4_5guard_6
reg1143 := __e.Call(__defun__fail)
reg1144 := PrimEqual(reg1143, Parse__shen_4_5guard_6)
reg1145 := PrimNot(reg1144)
var reg1160 Obj
if reg1145 == True {
reg1146 := PrimHead(Parse__shen_4_5guard_6)
reg1147 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1148 := MakeSymbol("where")
reg1149 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5guard_6)
reg1150 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg1151 := Nil;
reg1152 := PrimCons(reg1150, reg1151)
reg1153 := PrimCons(reg1149, reg1152)
reg1154 := PrimCons(reg1148, reg1153)
reg1155 := Nil;
reg1156 := PrimCons(reg1154, reg1155)
reg1157 := PrimCons(reg1147, reg1156)
reg1158 := __e.Call(__defun__shen_4pair, reg1146, reg1157)
reg1160 = reg1158
} else {
reg1159 := __e.Call(__defun__fail)
reg1160 = reg1159
}
reg1162 = reg1160
} else {
reg1161 := __e.Call(__defun__fail)
reg1162 = reg1161
}
reg1164 = reg1162
} else {
reg1163 := __e.Call(__defun__fail)
reg1164 = reg1163
}
reg1166 = reg1164
} else {
reg1165 := __e.Call(__defun__fail)
reg1166 = reg1165
}
reg1168 = reg1166
} else {
reg1167 := __e.Call(__defun__fail)
reg1168 = reg1167
}
YaccParse := reg1168
_ = YaccParse
reg1169 := __e.Call(__defun__fail)
reg1170 := PrimEqual(YaccParse, reg1169)
if reg1170 == True {
reg1171 := __e.Call(__defun__shen_4_5patterns_6, V165)
Parse__shen_4_5patterns_6 := reg1171
_ = Parse__shen_4_5patterns_6
reg1172 := __e.Call(__defun__fail)
reg1173 := PrimEqual(reg1172, Parse__shen_4_5patterns_6)
reg1174 := PrimNot(reg1173)
var reg1204 Obj
if reg1174 == True {
reg1175 := PrimHead(Parse__shen_4_5patterns_6)
reg1176 := PrimIsPair(reg1175)
var reg1184 Obj
if reg1176 == True {
reg1177 := MakeSymbol("->")
reg1178 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5patterns_6)
reg1179 := PrimEqual(reg1177, reg1178)
var reg1182 Obj
if reg1179 == True {
reg1180 := True;
reg1182 = reg1180
} else {
reg1181 := False;
reg1182 = reg1181
}
reg1184 = reg1182
} else {
reg1183 := False;
reg1184 = reg1183
}
var reg1202 Obj
if reg1184 == True {
reg1185 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5patterns_6)
reg1186 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1187 := __e.Call(__defun__shen_4pair, reg1185, reg1186)
NewStream160 := reg1187
_ = NewStream160
reg1188 := __e.Call(__defun__shen_4_5action_6, NewStream160)
Parse__shen_4_5action_6 := reg1188
_ = Parse__shen_4_5action_6
reg1189 := __e.Call(__defun__fail)
reg1190 := PrimEqual(reg1189, Parse__shen_4_5action_6)
reg1191 := PrimNot(reg1190)
var reg1200 Obj
if reg1191 == True {
reg1192 := PrimHead(Parse__shen_4_5action_6)
reg1193 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1194 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg1195 := Nil;
reg1196 := PrimCons(reg1194, reg1195)
reg1197 := PrimCons(reg1193, reg1196)
reg1198 := __e.Call(__defun__shen_4pair, reg1192, reg1197)
reg1200 = reg1198
} else {
reg1199 := __e.Call(__defun__fail)
reg1200 = reg1199
}
reg1202 = reg1200
} else {
reg1201 := __e.Call(__defun__fail)
reg1202 = reg1201
}
reg1204 = reg1202
} else {
reg1203 := __e.Call(__defun__fail)
reg1204 = reg1203
}
YaccParse := reg1204
_ = YaccParse
reg1205 := __e.Call(__defun__fail)
reg1206 := PrimEqual(YaccParse, reg1205)
if reg1206 == True {
reg1207 := __e.Call(__defun__shen_4_5patterns_6, V165)
Parse__shen_4_5patterns_6 := reg1207
_ = Parse__shen_4_5patterns_6
reg1208 := __e.Call(__defun__fail)
reg1209 := PrimEqual(reg1208, Parse__shen_4_5patterns_6)
reg1210 := PrimNot(reg1209)
var reg1271 Obj
if reg1210 == True {
reg1211 := PrimHead(Parse__shen_4_5patterns_6)
reg1212 := PrimIsPair(reg1211)
var reg1220 Obj
if reg1212 == True {
reg1213 := MakeSymbol("<-")
reg1214 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5patterns_6)
reg1215 := PrimEqual(reg1213, reg1214)
var reg1218 Obj
if reg1215 == True {
reg1216 := True;
reg1218 = reg1216
} else {
reg1217 := False;
reg1218 = reg1217
}
reg1220 = reg1218
} else {
reg1219 := False;
reg1220 = reg1219
}
var reg1269 Obj
if reg1220 == True {
reg1221 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5patterns_6)
reg1222 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1223 := __e.Call(__defun__shen_4pair, reg1221, reg1222)
NewStream161 := reg1223
_ = NewStream161
reg1224 := __e.Call(__defun__shen_4_5action_6, NewStream161)
Parse__shen_4_5action_6 := reg1224
_ = Parse__shen_4_5action_6
reg1225 := __e.Call(__defun__fail)
reg1226 := PrimEqual(reg1225, Parse__shen_4_5action_6)
reg1227 := PrimNot(reg1226)
var reg1267 Obj
if reg1227 == True {
reg1228 := PrimHead(Parse__shen_4_5action_6)
reg1229 := PrimIsPair(reg1228)
var reg1237 Obj
if reg1229 == True {
reg1230 := MakeSymbol("where")
reg1231 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5action_6)
reg1232 := PrimEqual(reg1230, reg1231)
var reg1235 Obj
if reg1232 == True {
reg1233 := True;
reg1235 = reg1233
} else {
reg1234 := False;
reg1235 = reg1234
}
reg1237 = reg1235
} else {
reg1236 := False;
reg1237 = reg1236
}
var reg1265 Obj
if reg1237 == True {
reg1238 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5action_6)
reg1239 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg1240 := __e.Call(__defun__shen_4pair, reg1238, reg1239)
NewStream162 := reg1240
_ = NewStream162
reg1241 := __e.Call(__defun__shen_4_5guard_6, NewStream162)
Parse__shen_4_5guard_6 := reg1241
_ = Parse__shen_4_5guard_6
reg1242 := __e.Call(__defun__fail)
reg1243 := PrimEqual(reg1242, Parse__shen_4_5guard_6)
reg1244 := PrimNot(reg1243)
var reg1263 Obj
if reg1244 == True {
reg1245 := PrimHead(Parse__shen_4_5guard_6)
reg1246 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1247 := MakeSymbol("where")
reg1248 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5guard_6)
reg1249 := MakeSymbol("shen.choicepoint!")
reg1250 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg1251 := Nil;
reg1252 := PrimCons(reg1250, reg1251)
reg1253 := PrimCons(reg1249, reg1252)
reg1254 := Nil;
reg1255 := PrimCons(reg1253, reg1254)
reg1256 := PrimCons(reg1248, reg1255)
reg1257 := PrimCons(reg1247, reg1256)
reg1258 := Nil;
reg1259 := PrimCons(reg1257, reg1258)
reg1260 := PrimCons(reg1246, reg1259)
reg1261 := __e.Call(__defun__shen_4pair, reg1245, reg1260)
reg1263 = reg1261
} else {
reg1262 := __e.Call(__defun__fail)
reg1263 = reg1262
}
reg1265 = reg1263
} else {
reg1264 := __e.Call(__defun__fail)
reg1265 = reg1264
}
reg1267 = reg1265
} else {
reg1266 := __e.Call(__defun__fail)
reg1267 = reg1266
}
reg1269 = reg1267
} else {
reg1268 := __e.Call(__defun__fail)
reg1269 = reg1268
}
reg1271 = reg1269
} else {
reg1270 := __e.Call(__defun__fail)
reg1271 = reg1270
}
YaccParse := reg1271
_ = YaccParse
reg1272 := __e.Call(__defun__fail)
reg1273 := PrimEqual(YaccParse, reg1272)
if reg1273 == True {
reg1274 := __e.Call(__defun__shen_4_5patterns_6, V165)
Parse__shen_4_5patterns_6 := reg1274
_ = Parse__shen_4_5patterns_6
reg1275 := __e.Call(__defun__fail)
reg1276 := PrimEqual(reg1275, Parse__shen_4_5patterns_6)
reg1277 := PrimNot(reg1276)
if reg1277 == True {
reg1278 := PrimHead(Parse__shen_4_5patterns_6)
reg1279 := PrimIsPair(reg1278)
var reg1287 Obj
if reg1279 == True {
reg1280 := MakeSymbol("<-")
reg1281 := __e.Call(__defun__shen_4hdhd, Parse__shen_4_5patterns_6)
reg1282 := PrimEqual(reg1280, reg1281)
var reg1285 Obj
if reg1282 == True {
reg1283 := True;
reg1285 = reg1283
} else {
reg1284 := False;
reg1285 = reg1284
}
reg1287 = reg1285
} else {
reg1286 := False;
reg1287 = reg1286
}
if reg1287 == True {
reg1288 := __e.Call(__defun__shen_4tlhd, Parse__shen_4_5patterns_6)
reg1289 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1290 := __e.Call(__defun__shen_4pair, reg1288, reg1289)
NewStream163 := reg1290
_ = NewStream163
reg1291 := __e.Call(__defun__shen_4_5action_6, NewStream163)
Parse__shen_4_5action_6 := reg1291
_ = Parse__shen_4_5action_6
reg1292 := __e.Call(__defun__fail)
reg1293 := PrimEqual(reg1292, Parse__shen_4_5action_6)
reg1294 := PrimNot(reg1293)
if reg1294 == True {
reg1295 := PrimHead(Parse__shen_4_5action_6)
reg1296 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1297 := MakeSymbol("shen.choicepoint!")
reg1298 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5action_6)
reg1299 := Nil;
reg1300 := PrimCons(reg1298, reg1299)
reg1301 := PrimCons(reg1297, reg1300)
reg1302 := Nil;
reg1303 := PrimCons(reg1301, reg1302)
reg1304 := PrimCons(reg1296, reg1303)
__ctx.TailApply(__defun__shen_4pair, reg1295, reg1304)
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
V168 := __args[0]
_ = V168
V169 := __args[1]
_ = V169
reg1309 := __e.Call(V168, V169)
if reg1309 == True {
__ctx.TailApply(__defun__fail)
return
} else {
__ctx.Return(V169)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.fail_if", value: __defun__shen_4fail__if})

__defun__shen_4succeeds_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V175 := __args[0]
_ = V175
reg1311 := __e.Call(__defun__fail)
reg1312 := PrimEqual(V175, reg1311)
if reg1312 == True {
reg1313 := False;
__ctx.Return(reg1313)
return
} else {
reg1314 := True;
__ctx.Return(reg1314)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.succeeds?", value: __defun__shen_4succeeds_2})

__defun__shen_4_5patterns_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V177 := __args[0]
_ = V177
reg1315 := __e.Call(__defun__shen_4_5pattern_6, V177)
Parse__shen_4_5pattern_6 := reg1315
_ = Parse__shen_4_5pattern_6
reg1316 := __e.Call(__defun__fail)
reg1317 := PrimEqual(reg1316, Parse__shen_4_5pattern_6)
reg1318 := PrimNot(reg1317)
var reg1331 Obj
if reg1318 == True {
reg1319 := __e.Call(__defun__shen_4_5patterns_6, Parse__shen_4_5pattern_6)
Parse__shen_4_5patterns_6 := reg1319
_ = Parse__shen_4_5patterns_6
reg1320 := __e.Call(__defun__fail)
reg1321 := PrimEqual(reg1320, Parse__shen_4_5patterns_6)
reg1322 := PrimNot(reg1321)
var reg1329 Obj
if reg1322 == True {
reg1323 := PrimHead(Parse__shen_4_5patterns_6)
reg1324 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern_6)
reg1325 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5patterns_6)
reg1326 := PrimCons(reg1324, reg1325)
reg1327 := __e.Call(__defun__shen_4pair, reg1323, reg1326)
reg1329 = reg1327
} else {
reg1328 := __e.Call(__defun__fail)
reg1329 = reg1328
}
reg1331 = reg1329
} else {
reg1330 := __e.Call(__defun__fail)
reg1331 = reg1330
}
YaccParse := reg1331
_ = YaccParse
reg1332 := __e.Call(__defun__fail)
reg1333 := PrimEqual(YaccParse, reg1332)
if reg1333 == True {
reg1334 := __e.Call(__defun___5e_6, V177)
Parse___5e_6 := reg1334
_ = Parse___5e_6
reg1335 := __e.Call(__defun__fail)
reg1336 := PrimEqual(reg1335, Parse___5e_6)
reg1337 := PrimNot(reg1336)
if reg1337 == True {
reg1338 := PrimHead(Parse___5e_6)
reg1339 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg1338, reg1339)
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
V190 := __args[0]
_ = V190
reg1342 := PrimHead(V190)
reg1343 := PrimIsPair(reg1342)
var reg1350 Obj
if reg1343 == True {
reg1344 := __e.Call(__defun__shen_4hdhd, V190)
reg1345 := PrimIsPair(reg1344)
var reg1348 Obj
if reg1345 == True {
reg1346 := True;
reg1348 = reg1346
} else {
reg1347 := False;
reg1348 = reg1347
}
reg1350 = reg1348
} else {
reg1349 := False;
reg1350 = reg1349
}
var reg1403 Obj
if reg1350 == True {
reg1351 := __e.Call(__defun__shen_4hdhd, V190)
reg1352 := __e.Call(__defun__shen_4hdtl, V190)
reg1353 := __e.Call(__defun__shen_4pair, reg1351, reg1352)
reg1354 := PrimHead(reg1353)
reg1355 := PrimIsPair(reg1354)
var reg1366 Obj
if reg1355 == True {
reg1356 := MakeSymbol("@p")
reg1357 := __e.Call(__defun__shen_4hdhd, V190)
reg1358 := __e.Call(__defun__shen_4hdtl, V190)
reg1359 := __e.Call(__defun__shen_4pair, reg1357, reg1358)
reg1360 := __e.Call(__defun__shen_4hdhd, reg1359)
reg1361 := PrimEqual(reg1356, reg1360)
var reg1364 Obj
if reg1361 == True {
reg1362 := True;
reg1364 = reg1362
} else {
reg1363 := False;
reg1364 = reg1363
}
reg1366 = reg1364
} else {
reg1365 := False;
reg1366 = reg1365
}
var reg1401 Obj
if reg1366 == True {
reg1367 := __e.Call(__defun__shen_4hdhd, V190)
reg1368 := __e.Call(__defun__shen_4hdtl, V190)
reg1369 := __e.Call(__defun__shen_4pair, reg1367, reg1368)
reg1370 := __e.Call(__defun__shen_4tlhd, reg1369)
reg1371 := __e.Call(__defun__shen_4hdhd, V190)
reg1372 := __e.Call(__defun__shen_4hdtl, V190)
reg1373 := __e.Call(__defun__shen_4pair, reg1371, reg1372)
reg1374 := __e.Call(__defun__shen_4hdtl, reg1373)
reg1375 := __e.Call(__defun__shen_4pair, reg1370, reg1374)
NewStream179 := reg1375
_ = NewStream179
reg1376 := __e.Call(__defun__shen_4_5pattern1_6, NewStream179)
Parse__shen_4_5pattern1_6 := reg1376
_ = Parse__shen_4_5pattern1_6
reg1377 := __e.Call(__defun__fail)
reg1378 := PrimEqual(reg1377, Parse__shen_4_5pattern1_6)
reg1379 := PrimNot(reg1378)
var reg1399 Obj
if reg1379 == True {
reg1380 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg1380
_ = Parse__shen_4_5pattern2_6
reg1381 := __e.Call(__defun__fail)
reg1382 := PrimEqual(reg1381, Parse__shen_4_5pattern2_6)
reg1383 := PrimNot(reg1382)
var reg1397 Obj
if reg1383 == True {
reg1384 := __e.Call(__defun__shen_4tlhd, V190)
reg1385 := __e.Call(__defun__shen_4hdtl, V190)
reg1386 := __e.Call(__defun__shen_4pair, reg1384, reg1385)
reg1387 := PrimHead(reg1386)
reg1388 := MakeSymbol("@p")
reg1389 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg1390 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg1391 := Nil;
reg1392 := PrimCons(reg1390, reg1391)
reg1393 := PrimCons(reg1389, reg1392)
reg1394 := PrimCons(reg1388, reg1393)
reg1395 := __e.Call(__defun__shen_4pair, reg1387, reg1394)
reg1397 = reg1395
} else {
reg1396 := __e.Call(__defun__fail)
reg1397 = reg1396
}
reg1399 = reg1397
} else {
reg1398 := __e.Call(__defun__fail)
reg1399 = reg1398
}
reg1401 = reg1399
} else {
reg1400 := __e.Call(__defun__fail)
reg1401 = reg1400
}
reg1403 = reg1401
} else {
reg1402 := __e.Call(__defun__fail)
reg1403 = reg1402
}
YaccParse := reg1403
_ = YaccParse
reg1404 := __e.Call(__defun__fail)
reg1405 := PrimEqual(YaccParse, reg1404)
if reg1405 == True {
reg1406 := PrimHead(V190)
reg1407 := PrimIsPair(reg1406)
var reg1414 Obj
if reg1407 == True {
reg1408 := __e.Call(__defun__shen_4hdhd, V190)
reg1409 := PrimIsPair(reg1408)
var reg1412 Obj
if reg1409 == True {
reg1410 := True;
reg1412 = reg1410
} else {
reg1411 := False;
reg1412 = reg1411
}
reg1414 = reg1412
} else {
reg1413 := False;
reg1414 = reg1413
}
var reg1467 Obj
if reg1414 == True {
reg1415 := __e.Call(__defun__shen_4hdhd, V190)
reg1416 := __e.Call(__defun__shen_4hdtl, V190)
reg1417 := __e.Call(__defun__shen_4pair, reg1415, reg1416)
reg1418 := PrimHead(reg1417)
reg1419 := PrimIsPair(reg1418)
var reg1430 Obj
if reg1419 == True {
reg1420 := MakeSymbol("cons")
reg1421 := __e.Call(__defun__shen_4hdhd, V190)
reg1422 := __e.Call(__defun__shen_4hdtl, V190)
reg1423 := __e.Call(__defun__shen_4pair, reg1421, reg1422)
reg1424 := __e.Call(__defun__shen_4hdhd, reg1423)
reg1425 := PrimEqual(reg1420, reg1424)
var reg1428 Obj
if reg1425 == True {
reg1426 := True;
reg1428 = reg1426
} else {
reg1427 := False;
reg1428 = reg1427
}
reg1430 = reg1428
} else {
reg1429 := False;
reg1430 = reg1429
}
var reg1465 Obj
if reg1430 == True {
reg1431 := __e.Call(__defun__shen_4hdhd, V190)
reg1432 := __e.Call(__defun__shen_4hdtl, V190)
reg1433 := __e.Call(__defun__shen_4pair, reg1431, reg1432)
reg1434 := __e.Call(__defun__shen_4tlhd, reg1433)
reg1435 := __e.Call(__defun__shen_4hdhd, V190)
reg1436 := __e.Call(__defun__shen_4hdtl, V190)
reg1437 := __e.Call(__defun__shen_4pair, reg1435, reg1436)
reg1438 := __e.Call(__defun__shen_4hdtl, reg1437)
reg1439 := __e.Call(__defun__shen_4pair, reg1434, reg1438)
NewStream181 := reg1439
_ = NewStream181
reg1440 := __e.Call(__defun__shen_4_5pattern1_6, NewStream181)
Parse__shen_4_5pattern1_6 := reg1440
_ = Parse__shen_4_5pattern1_6
reg1441 := __e.Call(__defun__fail)
reg1442 := PrimEqual(reg1441, Parse__shen_4_5pattern1_6)
reg1443 := PrimNot(reg1442)
var reg1463 Obj
if reg1443 == True {
reg1444 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg1444
_ = Parse__shen_4_5pattern2_6
reg1445 := __e.Call(__defun__fail)
reg1446 := PrimEqual(reg1445, Parse__shen_4_5pattern2_6)
reg1447 := PrimNot(reg1446)
var reg1461 Obj
if reg1447 == True {
reg1448 := __e.Call(__defun__shen_4tlhd, V190)
reg1449 := __e.Call(__defun__shen_4hdtl, V190)
reg1450 := __e.Call(__defun__shen_4pair, reg1448, reg1449)
reg1451 := PrimHead(reg1450)
reg1452 := MakeSymbol("cons")
reg1453 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg1454 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg1455 := Nil;
reg1456 := PrimCons(reg1454, reg1455)
reg1457 := PrimCons(reg1453, reg1456)
reg1458 := PrimCons(reg1452, reg1457)
reg1459 := __e.Call(__defun__shen_4pair, reg1451, reg1458)
reg1461 = reg1459
} else {
reg1460 := __e.Call(__defun__fail)
reg1461 = reg1460
}
reg1463 = reg1461
} else {
reg1462 := __e.Call(__defun__fail)
reg1463 = reg1462
}
reg1465 = reg1463
} else {
reg1464 := __e.Call(__defun__fail)
reg1465 = reg1464
}
reg1467 = reg1465
} else {
reg1466 := __e.Call(__defun__fail)
reg1467 = reg1466
}
YaccParse := reg1467
_ = YaccParse
reg1468 := __e.Call(__defun__fail)
reg1469 := PrimEqual(YaccParse, reg1468)
if reg1469 == True {
reg1470 := PrimHead(V190)
reg1471 := PrimIsPair(reg1470)
var reg1478 Obj
if reg1471 == True {
reg1472 := __e.Call(__defun__shen_4hdhd, V190)
reg1473 := PrimIsPair(reg1472)
var reg1476 Obj
if reg1473 == True {
reg1474 := True;
reg1476 = reg1474
} else {
reg1475 := False;
reg1476 = reg1475
}
reg1478 = reg1476
} else {
reg1477 := False;
reg1478 = reg1477
}
var reg1531 Obj
if reg1478 == True {
reg1479 := __e.Call(__defun__shen_4hdhd, V190)
reg1480 := __e.Call(__defun__shen_4hdtl, V190)
reg1481 := __e.Call(__defun__shen_4pair, reg1479, reg1480)
reg1482 := PrimHead(reg1481)
reg1483 := PrimIsPair(reg1482)
var reg1494 Obj
if reg1483 == True {
reg1484 := MakeSymbol("@v")
reg1485 := __e.Call(__defun__shen_4hdhd, V190)
reg1486 := __e.Call(__defun__shen_4hdtl, V190)
reg1487 := __e.Call(__defun__shen_4pair, reg1485, reg1486)
reg1488 := __e.Call(__defun__shen_4hdhd, reg1487)
reg1489 := PrimEqual(reg1484, reg1488)
var reg1492 Obj
if reg1489 == True {
reg1490 := True;
reg1492 = reg1490
} else {
reg1491 := False;
reg1492 = reg1491
}
reg1494 = reg1492
} else {
reg1493 := False;
reg1494 = reg1493
}
var reg1529 Obj
if reg1494 == True {
reg1495 := __e.Call(__defun__shen_4hdhd, V190)
reg1496 := __e.Call(__defun__shen_4hdtl, V190)
reg1497 := __e.Call(__defun__shen_4pair, reg1495, reg1496)
reg1498 := __e.Call(__defun__shen_4tlhd, reg1497)
reg1499 := __e.Call(__defun__shen_4hdhd, V190)
reg1500 := __e.Call(__defun__shen_4hdtl, V190)
reg1501 := __e.Call(__defun__shen_4pair, reg1499, reg1500)
reg1502 := __e.Call(__defun__shen_4hdtl, reg1501)
reg1503 := __e.Call(__defun__shen_4pair, reg1498, reg1502)
NewStream183 := reg1503
_ = NewStream183
reg1504 := __e.Call(__defun__shen_4_5pattern1_6, NewStream183)
Parse__shen_4_5pattern1_6 := reg1504
_ = Parse__shen_4_5pattern1_6
reg1505 := __e.Call(__defun__fail)
reg1506 := PrimEqual(reg1505, Parse__shen_4_5pattern1_6)
reg1507 := PrimNot(reg1506)
var reg1527 Obj
if reg1507 == True {
reg1508 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg1508
_ = Parse__shen_4_5pattern2_6
reg1509 := __e.Call(__defun__fail)
reg1510 := PrimEqual(reg1509, Parse__shen_4_5pattern2_6)
reg1511 := PrimNot(reg1510)
var reg1525 Obj
if reg1511 == True {
reg1512 := __e.Call(__defun__shen_4tlhd, V190)
reg1513 := __e.Call(__defun__shen_4hdtl, V190)
reg1514 := __e.Call(__defun__shen_4pair, reg1512, reg1513)
reg1515 := PrimHead(reg1514)
reg1516 := MakeSymbol("@v")
reg1517 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg1518 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg1519 := Nil;
reg1520 := PrimCons(reg1518, reg1519)
reg1521 := PrimCons(reg1517, reg1520)
reg1522 := PrimCons(reg1516, reg1521)
reg1523 := __e.Call(__defun__shen_4pair, reg1515, reg1522)
reg1525 = reg1523
} else {
reg1524 := __e.Call(__defun__fail)
reg1525 = reg1524
}
reg1527 = reg1525
} else {
reg1526 := __e.Call(__defun__fail)
reg1527 = reg1526
}
reg1529 = reg1527
} else {
reg1528 := __e.Call(__defun__fail)
reg1529 = reg1528
}
reg1531 = reg1529
} else {
reg1530 := __e.Call(__defun__fail)
reg1531 = reg1530
}
YaccParse := reg1531
_ = YaccParse
reg1532 := __e.Call(__defun__fail)
reg1533 := PrimEqual(YaccParse, reg1532)
if reg1533 == True {
reg1534 := PrimHead(V190)
reg1535 := PrimIsPair(reg1534)
var reg1542 Obj
if reg1535 == True {
reg1536 := __e.Call(__defun__shen_4hdhd, V190)
reg1537 := PrimIsPair(reg1536)
var reg1540 Obj
if reg1537 == True {
reg1538 := True;
reg1540 = reg1538
} else {
reg1539 := False;
reg1540 = reg1539
}
reg1542 = reg1540
} else {
reg1541 := False;
reg1542 = reg1541
}
var reg1595 Obj
if reg1542 == True {
reg1543 := __e.Call(__defun__shen_4hdhd, V190)
reg1544 := __e.Call(__defun__shen_4hdtl, V190)
reg1545 := __e.Call(__defun__shen_4pair, reg1543, reg1544)
reg1546 := PrimHead(reg1545)
reg1547 := PrimIsPair(reg1546)
var reg1558 Obj
if reg1547 == True {
reg1548 := MakeSymbol("@s")
reg1549 := __e.Call(__defun__shen_4hdhd, V190)
reg1550 := __e.Call(__defun__shen_4hdtl, V190)
reg1551 := __e.Call(__defun__shen_4pair, reg1549, reg1550)
reg1552 := __e.Call(__defun__shen_4hdhd, reg1551)
reg1553 := PrimEqual(reg1548, reg1552)
var reg1556 Obj
if reg1553 == True {
reg1554 := True;
reg1556 = reg1554
} else {
reg1555 := False;
reg1556 = reg1555
}
reg1558 = reg1556
} else {
reg1557 := False;
reg1558 = reg1557
}
var reg1593 Obj
if reg1558 == True {
reg1559 := __e.Call(__defun__shen_4hdhd, V190)
reg1560 := __e.Call(__defun__shen_4hdtl, V190)
reg1561 := __e.Call(__defun__shen_4pair, reg1559, reg1560)
reg1562 := __e.Call(__defun__shen_4tlhd, reg1561)
reg1563 := __e.Call(__defun__shen_4hdhd, V190)
reg1564 := __e.Call(__defun__shen_4hdtl, V190)
reg1565 := __e.Call(__defun__shen_4pair, reg1563, reg1564)
reg1566 := __e.Call(__defun__shen_4hdtl, reg1565)
reg1567 := __e.Call(__defun__shen_4pair, reg1562, reg1566)
NewStream185 := reg1567
_ = NewStream185
reg1568 := __e.Call(__defun__shen_4_5pattern1_6, NewStream185)
Parse__shen_4_5pattern1_6 := reg1568
_ = Parse__shen_4_5pattern1_6
reg1569 := __e.Call(__defun__fail)
reg1570 := PrimEqual(reg1569, Parse__shen_4_5pattern1_6)
reg1571 := PrimNot(reg1570)
var reg1591 Obj
if reg1571 == True {
reg1572 := __e.Call(__defun__shen_4_5pattern2_6, Parse__shen_4_5pattern1_6)
Parse__shen_4_5pattern2_6 := reg1572
_ = Parse__shen_4_5pattern2_6
reg1573 := __e.Call(__defun__fail)
reg1574 := PrimEqual(reg1573, Parse__shen_4_5pattern2_6)
reg1575 := PrimNot(reg1574)
var reg1589 Obj
if reg1575 == True {
reg1576 := __e.Call(__defun__shen_4tlhd, V190)
reg1577 := __e.Call(__defun__shen_4hdtl, V190)
reg1578 := __e.Call(__defun__shen_4pair, reg1576, reg1577)
reg1579 := PrimHead(reg1578)
reg1580 := MakeSymbol("@s")
reg1581 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern1_6)
reg1582 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern2_6)
reg1583 := Nil;
reg1584 := PrimCons(reg1582, reg1583)
reg1585 := PrimCons(reg1581, reg1584)
reg1586 := PrimCons(reg1580, reg1585)
reg1587 := __e.Call(__defun__shen_4pair, reg1579, reg1586)
reg1589 = reg1587
} else {
reg1588 := __e.Call(__defun__fail)
reg1589 = reg1588
}
reg1591 = reg1589
} else {
reg1590 := __e.Call(__defun__fail)
reg1591 = reg1590
}
reg1593 = reg1591
} else {
reg1592 := __e.Call(__defun__fail)
reg1593 = reg1592
}
reg1595 = reg1593
} else {
reg1594 := __e.Call(__defun__fail)
reg1595 = reg1594
}
YaccParse := reg1595
_ = YaccParse
reg1596 := __e.Call(__defun__fail)
reg1597 := PrimEqual(YaccParse, reg1596)
if reg1597 == True {
reg1598 := PrimHead(V190)
reg1599 := PrimIsPair(reg1598)
var reg1606 Obj
if reg1599 == True {
reg1600 := __e.Call(__defun__shen_4hdhd, V190)
reg1601 := PrimIsPair(reg1600)
var reg1604 Obj
if reg1601 == True {
reg1602 := True;
reg1604 = reg1602
} else {
reg1603 := False;
reg1604 = reg1603
}
reg1606 = reg1604
} else {
reg1605 := False;
reg1606 = reg1605
}
var reg1660 Obj
if reg1606 == True {
reg1607 := __e.Call(__defun__shen_4hdhd, V190)
reg1608 := __e.Call(__defun__shen_4hdtl, V190)
reg1609 := __e.Call(__defun__shen_4pair, reg1607, reg1608)
reg1610 := PrimHead(reg1609)
reg1611 := PrimIsPair(reg1610)
var reg1622 Obj
if reg1611 == True {
reg1612 := MakeSymbol("vector")
reg1613 := __e.Call(__defun__shen_4hdhd, V190)
reg1614 := __e.Call(__defun__shen_4hdtl, V190)
reg1615 := __e.Call(__defun__shen_4pair, reg1613, reg1614)
reg1616 := __e.Call(__defun__shen_4hdhd, reg1615)
reg1617 := PrimEqual(reg1612, reg1616)
var reg1620 Obj
if reg1617 == True {
reg1618 := True;
reg1620 = reg1618
} else {
reg1619 := False;
reg1620 = reg1619
}
reg1622 = reg1620
} else {
reg1621 := False;
reg1622 = reg1621
}
var reg1658 Obj
if reg1622 == True {
reg1623 := __e.Call(__defun__shen_4hdhd, V190)
reg1624 := __e.Call(__defun__shen_4hdtl, V190)
reg1625 := __e.Call(__defun__shen_4pair, reg1623, reg1624)
reg1626 := __e.Call(__defun__shen_4tlhd, reg1625)
reg1627 := __e.Call(__defun__shen_4hdhd, V190)
reg1628 := __e.Call(__defun__shen_4hdtl, V190)
reg1629 := __e.Call(__defun__shen_4pair, reg1627, reg1628)
reg1630 := __e.Call(__defun__shen_4hdtl, reg1629)
reg1631 := __e.Call(__defun__shen_4pair, reg1626, reg1630)
NewStream187 := reg1631
_ = NewStream187
reg1632 := PrimHead(NewStream187)
reg1633 := PrimIsPair(reg1632)
var reg1641 Obj
if reg1633 == True {
reg1634 := MakeNumber(0)
reg1635 := __e.Call(__defun__shen_4hdhd, NewStream187)
reg1636 := PrimEqual(reg1634, reg1635)
var reg1639 Obj
if reg1636 == True {
reg1637 := True;
reg1639 = reg1637
} else {
reg1638 := False;
reg1639 = reg1638
}
reg1641 = reg1639
} else {
reg1640 := False;
reg1641 = reg1640
}
var reg1656 Obj
if reg1641 == True {
reg1642 := __e.Call(__defun__shen_4tlhd, NewStream187)
reg1643 := __e.Call(__defun__shen_4hdtl, NewStream187)
reg1644 := __e.Call(__defun__shen_4pair, reg1642, reg1643)
NewStream188 := reg1644
_ = NewStream188
reg1645 := __e.Call(__defun__shen_4tlhd, V190)
reg1646 := __e.Call(__defun__shen_4hdtl, V190)
reg1647 := __e.Call(__defun__shen_4pair, reg1645, reg1646)
reg1648 := PrimHead(reg1647)
reg1649 := MakeSymbol("vector")
reg1650 := MakeNumber(0)
reg1651 := Nil;
reg1652 := PrimCons(reg1650, reg1651)
reg1653 := PrimCons(reg1649, reg1652)
reg1654 := __e.Call(__defun__shen_4pair, reg1648, reg1653)
reg1656 = reg1654
} else {
reg1655 := __e.Call(__defun__fail)
reg1656 = reg1655
}
reg1658 = reg1656
} else {
reg1657 := __e.Call(__defun__fail)
reg1658 = reg1657
}
reg1660 = reg1658
} else {
reg1659 := __e.Call(__defun__fail)
reg1660 = reg1659
}
YaccParse := reg1660
_ = YaccParse
reg1661 := __e.Call(__defun__fail)
reg1662 := PrimEqual(YaccParse, reg1661)
if reg1662 == True {
reg1663 := PrimHead(V190)
reg1664 := PrimIsPair(reg1663)
var reg1676 Obj
if reg1664 == True {
reg1665 := __e.Call(__defun__shen_4hdhd, V190)
Parse__X := reg1665
_ = Parse__X
reg1666 := PrimIsPair(Parse__X)
var reg1674 Obj
if reg1666 == True {
reg1667 := __e.Call(__defun__shen_4tlhd, V190)
reg1668 := __e.Call(__defun__shen_4hdtl, V190)
reg1669 := __e.Call(__defun__shen_4pair, reg1667, reg1668)
reg1670 := PrimHead(reg1669)
reg1671 := __e.Call(__defun__shen_4constructor_1error, Parse__X)
reg1672 := __e.Call(__defun__shen_4pair, reg1670, reg1671)
reg1674 = reg1672
} else {
reg1673 := __e.Call(__defun__fail)
reg1674 = reg1673
}
reg1676 = reg1674
} else {
reg1675 := __e.Call(__defun__fail)
reg1676 = reg1675
}
YaccParse := reg1676
_ = YaccParse
reg1677 := __e.Call(__defun__fail)
reg1678 := PrimEqual(YaccParse, reg1677)
if reg1678 == True {
reg1679 := __e.Call(__defun__shen_4_5simple__pattern_6, V190)
Parse__shen_4_5simple__pattern_6 := reg1679
_ = Parse__shen_4_5simple__pattern_6
reg1680 := __e.Call(__defun__fail)
reg1681 := PrimEqual(reg1680, Parse__shen_4_5simple__pattern_6)
reg1682 := PrimNot(reg1681)
if reg1682 == True {
reg1683 := PrimHead(Parse__shen_4_5simple__pattern_6)
reg1684 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5simple__pattern_6)
__ctx.TailApply(__defun__shen_4pair, reg1683, reg1684)
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
V192 := __args[0]
_ = V192
reg1687 := MakeString(" is not a legitimate constructor\n")
reg1688 := MakeSymbol("shen.a")
reg1689 := __e.Call(__defun__shen_4app, V192, reg1687, reg1688)
reg1690 := PrimSimpleError(reg1689)
__ctx.Return(reg1690)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.constructor-error", value: __defun__shen_4constructor_1error})

__defun__shen_4_5simple__pattern_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V194 := __args[0]
_ = V194
reg1691 := PrimHead(V194)
reg1692 := PrimIsPair(reg1691)
var reg1706 Obj
if reg1692 == True {
reg1693 := __e.Call(__defun__shen_4hdhd, V194)
Parse__X := reg1693
_ = Parse__X
reg1694 := MakeSymbol("_")
reg1695 := PrimEqual(Parse__X, reg1694)
var reg1704 Obj
if reg1695 == True {
reg1696 := __e.Call(__defun__shen_4tlhd, V194)
reg1697 := __e.Call(__defun__shen_4hdtl, V194)
reg1698 := __e.Call(__defun__shen_4pair, reg1696, reg1697)
reg1699 := PrimHead(reg1698)
reg1700 := MakeSymbol("Parse_Y")
reg1701 := __e.Call(__defun__gensym, reg1700)
reg1702 := __e.Call(__defun__shen_4pair, reg1699, reg1701)
reg1704 = reg1702
} else {
reg1703 := __e.Call(__defun__fail)
reg1704 = reg1703
}
reg1706 = reg1704
} else {
reg1705 := __e.Call(__defun__fail)
reg1706 = reg1705
}
YaccParse := reg1706
_ = YaccParse
reg1707 := __e.Call(__defun__fail)
reg1708 := PrimEqual(YaccParse, reg1707)
if reg1708 == True {
reg1709 := PrimHead(V194)
reg1710 := PrimIsPair(reg1709)
if reg1710 == True {
reg1711 := __e.Call(__defun__shen_4hdhd, V194)
Parse__X := reg1711
_ = Parse__X
reg1712 := MakeSymbol("->")
reg1713 := MakeSymbol("<-")
reg1714 := Nil;
reg1715 := PrimCons(reg1713, reg1714)
reg1716 := PrimCons(reg1712, reg1715)
reg1717 := __e.Call(__defun__element_2, Parse__X, reg1716)
reg1718 := PrimNot(reg1717)
if reg1718 == True {
reg1719 := __e.Call(__defun__shen_4tlhd, V194)
reg1720 := __e.Call(__defun__shen_4hdtl, V194)
reg1721 := __e.Call(__defun__shen_4pair, reg1719, reg1720)
reg1722 := PrimHead(reg1721)
__ctx.TailApply(__defun__shen_4pair, reg1722, Parse__X)
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
V196 := __args[0]
_ = V196
reg1726 := __e.Call(__defun__shen_4_5pattern_6, V196)
Parse__shen_4_5pattern_6 := reg1726
_ = Parse__shen_4_5pattern_6
reg1727 := __e.Call(__defun__fail)
reg1728 := PrimEqual(reg1727, Parse__shen_4_5pattern_6)
reg1729 := PrimNot(reg1728)
if reg1729 == True {
reg1730 := PrimHead(Parse__shen_4_5pattern_6)
reg1731 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern_6)
__ctx.TailApply(__defun__shen_4pair, reg1730, reg1731)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<pattern1>", value: __defun__shen_4_5pattern1_6})

__defun__shen_4_5pattern2_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V198 := __args[0]
_ = V198
reg1734 := __e.Call(__defun__shen_4_5pattern_6, V198)
Parse__shen_4_5pattern_6 := reg1734
_ = Parse__shen_4_5pattern_6
reg1735 := __e.Call(__defun__fail)
reg1736 := PrimEqual(reg1735, Parse__shen_4_5pattern_6)
reg1737 := PrimNot(reg1736)
if reg1737 == True {
reg1738 := PrimHead(Parse__shen_4_5pattern_6)
reg1739 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5pattern_6)
__ctx.TailApply(__defun__shen_4pair, reg1738, reg1739)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<pattern2>", value: __defun__shen_4_5pattern2_6})

__defun__shen_4_5action_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V200 := __args[0]
_ = V200
reg1742 := PrimHead(V200)
reg1743 := PrimIsPair(reg1742)
if reg1743 == True {
reg1744 := __e.Call(__defun__shen_4hdhd, V200)
Parse__X := reg1744
_ = Parse__X
reg1745 := __e.Call(__defun__shen_4tlhd, V200)
reg1746 := __e.Call(__defun__shen_4hdtl, V200)
reg1747 := __e.Call(__defun__shen_4pair, reg1745, reg1746)
reg1748 := PrimHead(reg1747)
__ctx.TailApply(__defun__shen_4pair, reg1748, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<action>", value: __defun__shen_4_5action_6})

__defun__shen_4_5guard_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V202 := __args[0]
_ = V202
reg1751 := PrimHead(V202)
reg1752 := PrimIsPair(reg1751)
if reg1752 == True {
reg1753 := __e.Call(__defun__shen_4hdhd, V202)
Parse__X := reg1753
_ = Parse__X
reg1754 := __e.Call(__defun__shen_4tlhd, V202)
reg1755 := __e.Call(__defun__shen_4hdtl, V202)
reg1756 := __e.Call(__defun__shen_4pair, reg1754, reg1755)
reg1757 := PrimHead(reg1756)
__ctx.TailApply(__defun__shen_4pair, reg1757, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<guard>", value: __defun__shen_4_5guard_6})

__defun__shen_4compile__to__machine__code = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V205 := __args[0]
_ = V205
V206 := __args[1]
_ = V206
reg1760 := __e.Call(__defun__shen_4compile__to__lambda_7, V205, V206)
Lambda_7 := reg1760
_ = Lambda_7
reg1761 := __e.Call(__defun__shen_4compile__to__kl, V205, Lambda_7)
KL := reg1761
_ = KL
reg1762 := __e.Call(__defun__shen_4record_1source, V205, KL)
Record := reg1762
_ = Record
__ctx.Return(KL)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compile_to_machine_code", value: __defun__shen_4compile__to__machine__code})

__defun__shen_4record_1source = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V211 := __args[0]
_ = V211
V212 := __args[1]
_ = V212
reg1763 := MakeSymbol("shen.*installing-kl*")
reg1764 := PrimValue(reg1763)
if reg1764 == True {
reg1765 := MakeSymbol("shen.skip")
__ctx.Return(reg1765)
return
} else {
reg1766 := MakeSymbol("shen.source")
reg1767 := MakeSymbol("*property-vector*")
reg1768 := PrimValue(reg1767)
__ctx.TailApply(__defun__put, V211, reg1766, V212, reg1768)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.record-source", value: __defun__shen_4record_1source})

__defun__shen_4compile__to__lambda_7 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V215 := __args[0]
_ = V215
V216 := __args[1]
_ = V216
reg1770 := __e.Call(__defun__shen_4aritycheck, V215, V216)
Arity := reg1770
_ = Arity
reg1771 := __e.Call(__defun__shen_4update_1symbol_1table, V215, Arity)
UpDateSymbolTable := reg1771
_ = UpDateSymbolTable
reg1772 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Rule := __args[0]
_ = Rule
__ctx.TailApply(__defun__shen_4free__variable__check, V215, Rule)
return
}, 1)
reg1774 := __e.Call(__defun__shen_4for_1each, reg1772, V216)
Free := reg1774
_ = Free
reg1775 := __e.Call(__defun__shen_4parameters, Arity)
Variables := reg1775
_ = Variables
reg1776 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4strip_1protect, X)
return
}, 1)
reg1778 := __e.Call(__defun__map, reg1776, V216)
Strip := reg1778
_ = Strip
reg1779 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4abstract__rule, X)
return
}, 1)
reg1781 := __e.Call(__defun__map, reg1779, Strip)
Abstractions := reg1781
_ = Abstractions
reg1782 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4application__build, Variables, X)
return
}, 1)
reg1784 := __e.Call(__defun__map, reg1782, Abstractions)
Applications := reg1784
_ = Applications
reg1785 := Nil;
reg1786 := PrimCons(Applications, reg1785)
reg1787 := PrimCons(Variables, reg1786)
__ctx.Return(reg1787)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compile_to_lambda+", value: __defun__shen_4compile__to__lambda_7})

__defun__shen_4update_1symbol_1table = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V219 := __args[0]
_ = V219
V220 := __args[1]
_ = V220
reg1788 := MakeNumber(0)
reg1789 := PrimEqual(reg1788, V220)
if reg1789 == True {
reg1790 := MakeSymbol("shen.skip")
__ctx.Return(reg1790)
return
} else {
reg1791 := MakeSymbol("shen.lambda-form")
reg1792 := __e.Call(__defun__shen_4lambda_1form, V219, V220)
reg1793 := PrimEvalKL(__e, reg1792)
reg1794 := MakeSymbol("*property-vector*")
reg1795 := PrimValue(reg1794)
__ctx.TailApply(__defun__put, V219, reg1791, reg1793, reg1795)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.update-symbol-table", value: __defun__shen_4update_1symbol_1table})

__defun__shen_4free__variable__check = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V223 := __args[0]
_ = V223
V224 := __args[1]
_ = V224
reg1797 := PrimIsPair(V224)
var reg1813 Obj
if reg1797 == True {
reg1798 := PrimTail(V224)
reg1799 := PrimIsPair(reg1798)
var reg1808 Obj
if reg1799 == True {
reg1800 := Nil;
reg1801 := PrimTail(V224)
reg1802 := PrimTail(reg1801)
reg1803 := PrimEqual(reg1800, reg1802)
var reg1806 Obj
if reg1803 == True {
reg1804 := True;
reg1806 = reg1804
} else {
reg1805 := False;
reg1806 = reg1805
}
reg1808 = reg1806
} else {
reg1807 := False;
reg1808 = reg1807
}
var reg1811 Obj
if reg1808 == True {
reg1809 := True;
reg1811 = reg1809
} else {
reg1810 := False;
reg1811 = reg1810
}
reg1813 = reg1811
} else {
reg1812 := False;
reg1813 = reg1812
}
if reg1813 == True {
reg1814 := PrimHead(V224)
reg1815 := __e.Call(__defun__shen_4extract__vars, reg1814)
Bound := reg1815
_ = Bound
reg1816 := PrimTail(V224)
reg1817 := PrimHead(reg1816)
reg1818 := __e.Call(__defun__shen_4extract__free__vars, Bound, reg1817)
Free := reg1818
_ = Free
__ctx.TailApply(__defun__shen_4free__variable__warnings, V223, Free)
return
} else {
reg1820 := MakeSymbol("shen.free_variable_check")
__ctx.TailApply(__defun__shen_4f__error, reg1820)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.free_variable_check", value: __defun__shen_4free__variable__check})

__defun__shen_4extract__vars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V226 := __args[0]
_ = V226
reg1822 := PrimIsVariable(V226)
if reg1822 == True {
reg1823 := Nil;
reg1824 := PrimCons(V226, reg1823)
__ctx.Return(reg1824)
return
} else {
reg1825 := PrimIsPair(V226)
if reg1825 == True {
reg1826 := PrimHead(V226)
reg1827 := __e.Call(__defun__shen_4extract__vars, reg1826)
reg1828 := PrimTail(V226)
reg1829 := __e.Call(__defun__shen_4extract__vars, reg1828)
__ctx.TailApply(__defun__union, reg1827, reg1829)
return
} else {
reg1831 := Nil;
__ctx.Return(reg1831)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.extract_vars", value: __defun__shen_4extract__vars})

__defun__shen_4extract__free__vars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V238 := __args[0]
_ = V238
V239 := __args[1]
_ = V239
reg1832 := PrimIsPair(V239)
var reg1856 Obj
if reg1832 == True {
reg1833 := PrimTail(V239)
reg1834 := PrimIsPair(reg1833)
var reg1851 Obj
if reg1834 == True {
reg1835 := Nil;
reg1836 := PrimTail(V239)
reg1837 := PrimTail(reg1836)
reg1838 := PrimEqual(reg1835, reg1837)
var reg1846 Obj
if reg1838 == True {
reg1839 := PrimHead(V239)
reg1840 := MakeSymbol("protect")
reg1841 := PrimEqual(reg1839, reg1840)
var reg1844 Obj
if reg1841 == True {
reg1842 := True;
reg1844 = reg1842
} else {
reg1843 := False;
reg1844 = reg1843
}
reg1846 = reg1844
} else {
reg1845 := False;
reg1846 = reg1845
}
var reg1849 Obj
if reg1846 == True {
reg1847 := True;
reg1849 = reg1847
} else {
reg1848 := False;
reg1849 = reg1848
}
reg1851 = reg1849
} else {
reg1850 := False;
reg1851 = reg1850
}
var reg1854 Obj
if reg1851 == True {
reg1852 := True;
reg1854 = reg1852
} else {
reg1853 := False;
reg1854 = reg1853
}
reg1856 = reg1854
} else {
reg1855 := False;
reg1856 = reg1855
}
if reg1856 == True {
reg1857 := Nil;
__ctx.Return(reg1857)
return
} else {
reg1858 := PrimIsVariable(V239)
var reg1865 Obj
if reg1858 == True {
reg1859 := __e.Call(__defun__element_2, V239, V238)
reg1860 := PrimNot(reg1859)
var reg1863 Obj
if reg1860 == True {
reg1861 := True;
reg1863 = reg1861
} else {
reg1862 := False;
reg1863 = reg1862
}
reg1865 = reg1863
} else {
reg1864 := False;
reg1865 = reg1864
}
if reg1865 == True {
reg1866 := Nil;
reg1867 := PrimCons(V239, reg1866)
__ctx.Return(reg1867)
return
} else {
reg1868 := PrimIsPair(V239)
var reg1901 Obj
if reg1868 == True {
reg1869 := MakeSymbol("lambda")
reg1870 := PrimHead(V239)
reg1871 := PrimEqual(reg1869, reg1870)
var reg1896 Obj
if reg1871 == True {
reg1872 := PrimTail(V239)
reg1873 := PrimIsPair(reg1872)
var reg1891 Obj
if reg1873 == True {
reg1874 := PrimTail(V239)
reg1875 := PrimTail(reg1874)
reg1876 := PrimIsPair(reg1875)
var reg1886 Obj
if reg1876 == True {
reg1877 := Nil;
reg1878 := PrimTail(V239)
reg1879 := PrimTail(reg1878)
reg1880 := PrimTail(reg1879)
reg1881 := PrimEqual(reg1877, reg1880)
var reg1884 Obj
if reg1881 == True {
reg1882 := True;
reg1884 = reg1882
} else {
reg1883 := False;
reg1884 = reg1883
}
reg1886 = reg1884
} else {
reg1885 := False;
reg1886 = reg1885
}
var reg1889 Obj
if reg1886 == True {
reg1887 := True;
reg1889 = reg1887
} else {
reg1888 := False;
reg1889 = reg1888
}
reg1891 = reg1889
} else {
reg1890 := False;
reg1891 = reg1890
}
var reg1894 Obj
if reg1891 == True {
reg1892 := True;
reg1894 = reg1892
} else {
reg1893 := False;
reg1894 = reg1893
}
reg1896 = reg1894
} else {
reg1895 := False;
reg1896 = reg1895
}
var reg1899 Obj
if reg1896 == True {
reg1897 := True;
reg1899 = reg1897
} else {
reg1898 := False;
reg1899 = reg1898
}
reg1901 = reg1899
} else {
reg1900 := False;
reg1901 = reg1900
}
if reg1901 == True {
reg1902 := PrimTail(V239)
reg1903 := PrimHead(reg1902)
reg1904 := PrimCons(reg1903, V238)
reg1905 := PrimTail(V239)
reg1906 := PrimTail(reg1905)
reg1907 := PrimHead(reg1906)
__ctx.TailApply(__defun__shen_4extract__free__vars, reg1904, reg1907)
return
} else {
reg1909 := PrimIsPair(V239)
var reg1952 Obj
if reg1909 == True {
reg1910 := MakeSymbol("let")
reg1911 := PrimHead(V239)
reg1912 := PrimEqual(reg1910, reg1911)
var reg1947 Obj
if reg1912 == True {
reg1913 := PrimTail(V239)
reg1914 := PrimIsPair(reg1913)
var reg1942 Obj
if reg1914 == True {
reg1915 := PrimTail(V239)
reg1916 := PrimTail(reg1915)
reg1917 := PrimIsPair(reg1916)
var reg1937 Obj
if reg1917 == True {
reg1918 := PrimTail(V239)
reg1919 := PrimTail(reg1918)
reg1920 := PrimTail(reg1919)
reg1921 := PrimIsPair(reg1920)
var reg1932 Obj
if reg1921 == True {
reg1922 := Nil;
reg1923 := PrimTail(V239)
reg1924 := PrimTail(reg1923)
reg1925 := PrimTail(reg1924)
reg1926 := PrimTail(reg1925)
reg1927 := PrimEqual(reg1922, reg1926)
var reg1930 Obj
if reg1927 == True {
reg1928 := True;
reg1930 = reg1928
} else {
reg1929 := False;
reg1930 = reg1929
}
reg1932 = reg1930
} else {
reg1931 := False;
reg1932 = reg1931
}
var reg1935 Obj
if reg1932 == True {
reg1933 := True;
reg1935 = reg1933
} else {
reg1934 := False;
reg1935 = reg1934
}
reg1937 = reg1935
} else {
reg1936 := False;
reg1937 = reg1936
}
var reg1940 Obj
if reg1937 == True {
reg1938 := True;
reg1940 = reg1938
} else {
reg1939 := False;
reg1940 = reg1939
}
reg1942 = reg1940
} else {
reg1941 := False;
reg1942 = reg1941
}
var reg1945 Obj
if reg1942 == True {
reg1943 := True;
reg1945 = reg1943
} else {
reg1944 := False;
reg1945 = reg1944
}
reg1947 = reg1945
} else {
reg1946 := False;
reg1947 = reg1946
}
var reg1950 Obj
if reg1947 == True {
reg1948 := True;
reg1950 = reg1948
} else {
reg1949 := False;
reg1950 = reg1949
}
reg1952 = reg1950
} else {
reg1951 := False;
reg1952 = reg1951
}
if reg1952 == True {
reg1953 := PrimTail(V239)
reg1954 := PrimTail(reg1953)
reg1955 := PrimHead(reg1954)
reg1956 := __e.Call(__defun__shen_4extract__free__vars, V238, reg1955)
reg1957 := PrimTail(V239)
reg1958 := PrimHead(reg1957)
reg1959 := PrimCons(reg1958, V238)
reg1960 := PrimTail(V239)
reg1961 := PrimTail(reg1960)
reg1962 := PrimTail(reg1961)
reg1963 := PrimHead(reg1962)
reg1964 := __e.Call(__defun__shen_4extract__free__vars, reg1959, reg1963)
__ctx.TailApply(__defun__union, reg1956, reg1964)
return
} else {
reg1966 := PrimIsPair(V239)
if reg1966 == True {
reg1967 := PrimHead(V239)
reg1968 := __e.Call(__defun__shen_4extract__free__vars, V238, reg1967)
reg1969 := PrimTail(V239)
reg1970 := __e.Call(__defun__shen_4extract__free__vars, V238, reg1969)
__ctx.TailApply(__defun__union, reg1968, reg1970)
return
} else {
reg1972 := Nil;
__ctx.Return(reg1972)
return
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.extract_free_vars", value: __defun__shen_4extract__free__vars})

__defun__shen_4free__variable__warnings = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V244 := __args[0]
_ = V244
V245 := __args[1]
_ = V245
reg1973 := Nil;
reg1974 := PrimEqual(reg1973, V245)
if reg1974 == True {
reg1975 := MakeSymbol("_")
__ctx.Return(reg1975)
return
} else {
reg1976 := MakeString("error: the following variables are free in ")
reg1977 := MakeString(": ")
reg1978 := __e.Call(__defun__shen_4list__variables, V245)
reg1979 := MakeString("")
reg1980 := MakeSymbol("shen.a")
reg1981 := __e.Call(__defun__shen_4app, reg1978, reg1979, reg1980)
reg1982 := PrimStringConcat(reg1977, reg1981)
reg1983 := MakeSymbol("shen.a")
reg1984 := __e.Call(__defun__shen_4app, V244, reg1982, reg1983)
reg1985 := PrimStringConcat(reg1976, reg1984)
reg1986 := PrimSimpleError(reg1985)
__ctx.Return(reg1986)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.free_variable_warnings", value: __defun__shen_4free__variable__warnings})

__defun__shen_4list__variables = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V247 := __args[0]
_ = V247
reg1987 := PrimIsPair(V247)
var reg1995 Obj
if reg1987 == True {
reg1988 := Nil;
reg1989 := PrimTail(V247)
reg1990 := PrimEqual(reg1988, reg1989)
var reg1993 Obj
if reg1990 == True {
reg1991 := True;
reg1993 = reg1991
} else {
reg1992 := False;
reg1993 = reg1992
}
reg1995 = reg1993
} else {
reg1994 := False;
reg1995 = reg1994
}
if reg1995 == True {
reg1996 := PrimHead(V247)
reg1997 := PrimStr(reg1996)
reg1998 := MakeString(".")
reg1999 := PrimStringConcat(reg1997, reg1998)
__ctx.Return(reg1999)
return
} else {
reg2000 := PrimIsPair(V247)
if reg2000 == True {
reg2001 := PrimHead(V247)
reg2002 := PrimStr(reg2001)
reg2003 := MakeString(", ")
reg2004 := PrimTail(V247)
reg2005 := __e.Call(__defun__shen_4list__variables, reg2004)
reg2006 := PrimStringConcat(reg2003, reg2005)
reg2007 := PrimStringConcat(reg2002, reg2006)
__ctx.Return(reg2007)
return
} else {
reg2008 := MakeSymbol("shen.list_variables")
__ctx.TailApply(__defun__shen_4f__error, reg2008)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.list_variables", value: __defun__shen_4list__variables})

__defun__shen_4strip_1protect = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V249 := __args[0]
_ = V249
reg2010 := PrimIsPair(V249)
var reg2034 Obj
if reg2010 == True {
reg2011 := PrimTail(V249)
reg2012 := PrimIsPair(reg2011)
var reg2029 Obj
if reg2012 == True {
reg2013 := Nil;
reg2014 := PrimTail(V249)
reg2015 := PrimTail(reg2014)
reg2016 := PrimEqual(reg2013, reg2015)
var reg2024 Obj
if reg2016 == True {
reg2017 := PrimHead(V249)
reg2018 := MakeSymbol("protect")
reg2019 := PrimEqual(reg2017, reg2018)
var reg2022 Obj
if reg2019 == True {
reg2020 := True;
reg2022 = reg2020
} else {
reg2021 := False;
reg2022 = reg2021
}
reg2024 = reg2022
} else {
reg2023 := False;
reg2024 = reg2023
}
var reg2027 Obj
if reg2024 == True {
reg2025 := True;
reg2027 = reg2025
} else {
reg2026 := False;
reg2027 = reg2026
}
reg2029 = reg2027
} else {
reg2028 := False;
reg2029 = reg2028
}
var reg2032 Obj
if reg2029 == True {
reg2030 := True;
reg2032 = reg2030
} else {
reg2031 := False;
reg2032 = reg2031
}
reg2034 = reg2032
} else {
reg2033 := False;
reg2034 = reg2033
}
if reg2034 == True {
reg2035 := PrimTail(V249)
reg2036 := PrimHead(reg2035)
__ctx.TailApply(__defun__shen_4strip_1protect, reg2036)
return
} else {
reg2038 := PrimIsPair(V249)
if reg2038 == True {
reg2039 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Z := __args[0]
_ = Z
__ctx.TailApply(__defun__shen_4strip_1protect, Z)
return
}, 1)
__ctx.TailApply(__defun__map, reg2039, V249)
return
} else {
__ctx.Return(V249)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.strip-protect", value: __defun__shen_4strip_1protect})

__defun__shen_4linearise = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V251 := __args[0]
_ = V251
reg2042 := PrimIsPair(V251)
var reg2058 Obj
if reg2042 == True {
reg2043 := PrimTail(V251)
reg2044 := PrimIsPair(reg2043)
var reg2053 Obj
if reg2044 == True {
reg2045 := Nil;
reg2046 := PrimTail(V251)
reg2047 := PrimTail(reg2046)
reg2048 := PrimEqual(reg2045, reg2047)
var reg2051 Obj
if reg2048 == True {
reg2049 := True;
reg2051 = reg2049
} else {
reg2050 := False;
reg2051 = reg2050
}
reg2053 = reg2051
} else {
reg2052 := False;
reg2053 = reg2052
}
var reg2056 Obj
if reg2053 == True {
reg2054 := True;
reg2056 = reg2054
} else {
reg2055 := False;
reg2056 = reg2055
}
reg2058 = reg2056
} else {
reg2057 := False;
reg2058 = reg2057
}
if reg2058 == True {
reg2059 := PrimHead(V251)
reg2060 := __e.Call(__defun__shen_4flatten, reg2059)
reg2061 := PrimHead(V251)
reg2062 := PrimTail(V251)
reg2063 := PrimHead(reg2062)
__ctx.TailApply(__defun__shen_4linearise__help, reg2060, reg2061, reg2063)
return
} else {
reg2065 := MakeSymbol("shen.linearise")
__ctx.TailApply(__defun__shen_4f__error, reg2065)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.linearise", value: __defun__shen_4linearise})

__defun__shen_4flatten = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V253 := __args[0]
_ = V253
reg2067 := Nil;
reg2068 := PrimEqual(reg2067, V253)
if reg2068 == True {
reg2069 := Nil;
__ctx.Return(reg2069)
return
} else {
reg2070 := PrimIsPair(V253)
if reg2070 == True {
reg2071 := PrimHead(V253)
reg2072 := __e.Call(__defun__shen_4flatten, reg2071)
reg2073 := PrimTail(V253)
reg2074 := __e.Call(__defun__shen_4flatten, reg2073)
__ctx.TailApply(__defun__append, reg2072, reg2074)
return
} else {
reg2076 := Nil;
reg2077 := PrimCons(V253, reg2076)
__ctx.Return(reg2077)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.flatten", value: __defun__shen_4flatten})

__defun__shen_4linearise__help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V257 := __args[0]
_ = V257
V258 := __args[1]
_ = V258
V259 := __args[2]
_ = V259
reg2078 := Nil;
reg2079 := PrimEqual(reg2078, V257)
if reg2079 == True {
reg2080 := Nil;
reg2081 := PrimCons(V259, reg2080)
reg2082 := PrimCons(V258, reg2081)
__ctx.Return(reg2082)
return
} else {
reg2083 := PrimIsPair(V257)
if reg2083 == True {
reg2084 := PrimHead(V257)
reg2085 := PrimIsVariable(reg2084)
var reg2093 Obj
if reg2085 == True {
reg2086 := PrimHead(V257)
reg2087 := PrimTail(V257)
reg2088 := __e.Call(__defun__element_2, reg2086, reg2087)
var reg2091 Obj
if reg2088 == True {
reg2089 := True;
reg2091 = reg2089
} else {
reg2090 := False;
reg2091 = reg2090
}
reg2093 = reg2091
} else {
reg2092 := False;
reg2093 = reg2092
}
if reg2093 == True {
reg2094 := PrimHead(V257)
reg2095 := __e.Call(__defun__gensym, reg2094)
Var := reg2095
_ = Var
reg2096 := MakeSymbol("where")
reg2097 := MakeSymbol("=")
reg2098 := PrimHead(V257)
reg2099 := Nil;
reg2100 := PrimCons(Var, reg2099)
reg2101 := PrimCons(reg2098, reg2100)
reg2102 := PrimCons(reg2097, reg2101)
reg2103 := Nil;
reg2104 := PrimCons(V259, reg2103)
reg2105 := PrimCons(reg2102, reg2104)
reg2106 := PrimCons(reg2096, reg2105)
NewAction := reg2106
_ = NewAction
reg2107 := PrimHead(V257)
reg2108 := __e.Call(__defun__shen_4linearise__X, reg2107, Var, V258)
NewPatts := reg2108
_ = NewPatts
reg2109 := PrimTail(V257)
__ctx.TailApply(__defun__shen_4linearise__help, reg2109, NewPatts, NewAction)
return
} else {
reg2111 := PrimTail(V257)
__ctx.TailApply(__defun__shen_4linearise__help, reg2111, V258, V259)
return
}
} else {
reg2113 := MakeSymbol("shen.linearise_help")
__ctx.TailApply(__defun__shen_4f__error, reg2113)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.linearise_help", value: __defun__shen_4linearise__help})

__defun__shen_4linearise__X = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V272 := __args[0]
_ = V272
V273 := __args[1]
_ = V273
V274 := __args[2]
_ = V274
reg2115 := PrimEqual(V274, V272)
if reg2115 == True {
__ctx.Return(V273)
return
} else {
reg2116 := PrimIsPair(V274)
if reg2116 == True {
reg2117 := PrimHead(V274)
reg2118 := __e.Call(__defun__shen_4linearise__X, V272, V273, reg2117)
L := reg2118
_ = L
reg2119 := PrimHead(V274)
reg2120 := PrimEqual(L, reg2119)
if reg2120 == True {
reg2121 := PrimHead(V274)
reg2122 := PrimTail(V274)
reg2123 := __e.Call(__defun__shen_4linearise__X, V272, V273, reg2122)
reg2124 := PrimCons(reg2121, reg2123)
__ctx.Return(reg2124)
return
} else {
reg2125 := PrimTail(V274)
reg2126 := PrimCons(L, reg2125)
__ctx.Return(reg2126)
return
}
} else {
__ctx.Return(V274)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.linearise_X", value: __defun__shen_4linearise__X})

__defun__shen_4aritycheck = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V277 := __args[0]
_ = V277
V278 := __args[1]
_ = V278
reg2127 := PrimIsPair(V278)
var reg2160 Obj
if reg2127 == True {
reg2128 := PrimHead(V278)
reg2129 := PrimIsPair(reg2128)
var reg2155 Obj
if reg2129 == True {
reg2130 := PrimHead(V278)
reg2131 := PrimTail(reg2130)
reg2132 := PrimIsPair(reg2131)
var reg2150 Obj
if reg2132 == True {
reg2133 := Nil;
reg2134 := PrimHead(V278)
reg2135 := PrimTail(reg2134)
reg2136 := PrimTail(reg2135)
reg2137 := PrimEqual(reg2133, reg2136)
var reg2145 Obj
if reg2137 == True {
reg2138 := Nil;
reg2139 := PrimTail(V278)
reg2140 := PrimEqual(reg2138, reg2139)
var reg2143 Obj
if reg2140 == True {
reg2141 := True;
reg2143 = reg2141
} else {
reg2142 := False;
reg2143 = reg2142
}
reg2145 = reg2143
} else {
reg2144 := False;
reg2145 = reg2144
}
var reg2148 Obj
if reg2145 == True {
reg2146 := True;
reg2148 = reg2146
} else {
reg2147 := False;
reg2148 = reg2147
}
reg2150 = reg2148
} else {
reg2149 := False;
reg2150 = reg2149
}
var reg2153 Obj
if reg2150 == True {
reg2151 := True;
reg2153 = reg2151
} else {
reg2152 := False;
reg2153 = reg2152
}
reg2155 = reg2153
} else {
reg2154 := False;
reg2155 = reg2154
}
var reg2158 Obj
if reg2155 == True {
reg2156 := True;
reg2158 = reg2156
} else {
reg2157 := False;
reg2158 = reg2157
}
reg2160 = reg2158
} else {
reg2159 := False;
reg2160 = reg2159
}
if reg2160 == True {
reg2161 := PrimHead(V278)
reg2162 := PrimTail(reg2161)
reg2163 := PrimHead(reg2162)
reg2164 := __e.Call(__defun__shen_4aritycheck_1action, reg2163)
_ = reg2164
reg2165 := __e.Call(__defun__arity, V277)
reg2166 := PrimHead(V278)
reg2167 := PrimHead(reg2166)
reg2168 := __e.Call(__defun__length, reg2167)
__ctx.TailApply(__defun__shen_4aritycheck_1name, V277, reg2165, reg2168)
return
} else {
reg2170 := PrimIsPair(V278)
var reg2230 Obj
if reg2170 == True {
reg2171 := PrimHead(V278)
reg2172 := PrimIsPair(reg2171)
var reg2225 Obj
if reg2172 == True {
reg2173 := PrimHead(V278)
reg2174 := PrimTail(reg2173)
reg2175 := PrimIsPair(reg2174)
var reg2220 Obj
if reg2175 == True {
reg2176 := Nil;
reg2177 := PrimHead(V278)
reg2178 := PrimTail(reg2177)
reg2179 := PrimTail(reg2178)
reg2180 := PrimEqual(reg2176, reg2179)
var reg2215 Obj
if reg2180 == True {
reg2181 := PrimTail(V278)
reg2182 := PrimIsPair(reg2181)
var reg2210 Obj
if reg2182 == True {
reg2183 := PrimTail(V278)
reg2184 := PrimHead(reg2183)
reg2185 := PrimIsPair(reg2184)
var reg2205 Obj
if reg2185 == True {
reg2186 := PrimTail(V278)
reg2187 := PrimHead(reg2186)
reg2188 := PrimTail(reg2187)
reg2189 := PrimIsPair(reg2188)
var reg2200 Obj
if reg2189 == True {
reg2190 := Nil;
reg2191 := PrimTail(V278)
reg2192 := PrimHead(reg2191)
reg2193 := PrimTail(reg2192)
reg2194 := PrimTail(reg2193)
reg2195 := PrimEqual(reg2190, reg2194)
var reg2198 Obj
if reg2195 == True {
reg2196 := True;
reg2198 = reg2196
} else {
reg2197 := False;
reg2198 = reg2197
}
reg2200 = reg2198
} else {
reg2199 := False;
reg2200 = reg2199
}
var reg2203 Obj
if reg2200 == True {
reg2201 := True;
reg2203 = reg2201
} else {
reg2202 := False;
reg2203 = reg2202
}
reg2205 = reg2203
} else {
reg2204 := False;
reg2205 = reg2204
}
var reg2208 Obj
if reg2205 == True {
reg2206 := True;
reg2208 = reg2206
} else {
reg2207 := False;
reg2208 = reg2207
}
reg2210 = reg2208
} else {
reg2209 := False;
reg2210 = reg2209
}
var reg2213 Obj
if reg2210 == True {
reg2211 := True;
reg2213 = reg2211
} else {
reg2212 := False;
reg2213 = reg2212
}
reg2215 = reg2213
} else {
reg2214 := False;
reg2215 = reg2214
}
var reg2218 Obj
if reg2215 == True {
reg2216 := True;
reg2218 = reg2216
} else {
reg2217 := False;
reg2218 = reg2217
}
reg2220 = reg2218
} else {
reg2219 := False;
reg2220 = reg2219
}
var reg2223 Obj
if reg2220 == True {
reg2221 := True;
reg2223 = reg2221
} else {
reg2222 := False;
reg2223 = reg2222
}
reg2225 = reg2223
} else {
reg2224 := False;
reg2225 = reg2224
}
var reg2228 Obj
if reg2225 == True {
reg2226 := True;
reg2228 = reg2226
} else {
reg2227 := False;
reg2228 = reg2227
}
reg2230 = reg2228
} else {
reg2229 := False;
reg2230 = reg2229
}
if reg2230 == True {
reg2231 := PrimHead(V278)
reg2232 := PrimHead(reg2231)
reg2233 := __e.Call(__defun__length, reg2232)
reg2234 := PrimTail(V278)
reg2235 := PrimHead(reg2234)
reg2236 := PrimHead(reg2235)
reg2237 := __e.Call(__defun__length, reg2236)
reg2238 := PrimEqual(reg2233, reg2237)
if reg2238 == True {
reg2239 := PrimHead(V278)
reg2240 := PrimTail(reg2239)
reg2241 := PrimHead(reg2240)
reg2242 := __e.Call(__defun__shen_4aritycheck_1action, reg2241)
_ = reg2242
reg2243 := PrimTail(V278)
__ctx.TailApply(__defun__shen_4aritycheck, V277, reg2243)
return
} else {
reg2245 := MakeString("arity error in ")
reg2246 := MakeString("\n")
reg2247 := MakeSymbol("shen.a")
reg2248 := __e.Call(__defun__shen_4app, V277, reg2246, reg2247)
reg2249 := PrimStringConcat(reg2245, reg2248)
reg2250 := PrimSimpleError(reg2249)
__ctx.Return(reg2250)
return
}
} else {
reg2251 := MakeSymbol("shen.aritycheck")
__ctx.TailApply(__defun__shen_4f__error, reg2251)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.aritycheck", value: __defun__shen_4aritycheck})

__defun__shen_4aritycheck_1name = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V291 := __args[0]
_ = V291
V292 := __args[1]
_ = V292
V293 := __args[2]
_ = V293
reg2253 := MakeNumber(-1)
reg2254 := PrimEqual(reg2253, V292)
if reg2254 == True {
__ctx.Return(V293)
return
} else {
reg2255 := PrimEqual(V293, V292)
if reg2255 == True {
__ctx.Return(V293)
return
} else {
reg2256 := MakeString("\nwarning: changing the arity of ")
reg2257 := MakeString(" can cause errors.\n")
reg2258 := MakeSymbol("shen.a")
reg2259 := __e.Call(__defun__shen_4app, V291, reg2257, reg2258)
reg2260 := PrimStringConcat(reg2256, reg2259)
reg2261 := __e.Call(__defun__stoutput)
reg2262 := __e.Call(__defun__shen_4prhush, reg2260, reg2261)
_ = reg2262
__ctx.Return(V293)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.aritycheck-name", value: __defun__shen_4aritycheck_1name})

__defun__shen_4aritycheck_1action = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V299 := __args[0]
_ = V299
reg2263 := PrimIsPair(V299)
if reg2263 == True {
reg2264 := PrimHead(V299)
reg2265 := PrimTail(V299)
reg2266 := __e.Call(__defun__shen_4aah, reg2264, reg2265)
_ = reg2266
reg2267 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
__ctx.TailApply(__defun__shen_4aritycheck_1action, Y)
return
}, 1)
__ctx.TailApply(__defun__shen_4for_1each, reg2267, V299)
return
} else {
reg2270 := MakeSymbol("shen.skip")
__ctx.Return(reg2270)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.aritycheck-action", value: __defun__shen_4aritycheck_1action})

__defun__shen_4aah = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V302 := __args[0]
_ = V302
V303 := __args[1]
_ = V303
reg2271 := __e.Call(__defun__arity, V302)
Arity := reg2271
_ = Arity
reg2272 := __e.Call(__defun__length, V303)
Len := reg2272
_ = Len
reg2273 := MakeNumber(-1)
reg2274 := PrimGreatThan(Arity, reg2273)
var reg2280 Obj
if reg2274 == True {
reg2275 := PrimGreatThan(Len, Arity)
var reg2278 Obj
if reg2275 == True {
reg2276 := True;
reg2278 = reg2276
} else {
reg2277 := False;
reg2278 = reg2277
}
reg2280 = reg2278
} else {
reg2279 := False;
reg2280 = reg2279
}
if reg2280 == True {
reg2281 := MakeString("warning: ")
reg2282 := MakeString(" might not like ")
reg2283 := MakeString(" argument")
reg2284 := MakeNumber(1)
reg2285 := PrimGreatThan(Len, reg2284)
var reg2288 Obj
if reg2285 == True {
reg2286 := MakeString("s")
reg2288 = reg2286
} else {
reg2287 := MakeString("")
reg2288 = reg2287
}
reg2289 := MakeString(".\n")
reg2290 := MakeSymbol("shen.a")
reg2291 := __e.Call(__defun__shen_4app, reg2288, reg2289, reg2290)
reg2292 := PrimStringConcat(reg2283, reg2291)
reg2293 := MakeSymbol("shen.a")
reg2294 := __e.Call(__defun__shen_4app, Len, reg2292, reg2293)
reg2295 := PrimStringConcat(reg2282, reg2294)
reg2296 := MakeSymbol("shen.a")
reg2297 := __e.Call(__defun__shen_4app, V302, reg2295, reg2296)
reg2298 := PrimStringConcat(reg2281, reg2297)
reg2299 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg2298, reg2299)
return
} else {
reg2301 := MakeSymbol("shen.skip")
__ctx.Return(reg2301)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.aah", value: __defun__shen_4aah})

__defun__shen_4abstract__rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V305 := __args[0]
_ = V305
reg2302 := PrimIsPair(V305)
var reg2318 Obj
if reg2302 == True {
reg2303 := PrimTail(V305)
reg2304 := PrimIsPair(reg2303)
var reg2313 Obj
if reg2304 == True {
reg2305 := Nil;
reg2306 := PrimTail(V305)
reg2307 := PrimTail(reg2306)
reg2308 := PrimEqual(reg2305, reg2307)
var reg2311 Obj
if reg2308 == True {
reg2309 := True;
reg2311 = reg2309
} else {
reg2310 := False;
reg2311 = reg2310
}
reg2313 = reg2311
} else {
reg2312 := False;
reg2313 = reg2312
}
var reg2316 Obj
if reg2313 == True {
reg2314 := True;
reg2316 = reg2314
} else {
reg2315 := False;
reg2316 = reg2315
}
reg2318 = reg2316
} else {
reg2317 := False;
reg2318 = reg2317
}
if reg2318 == True {
reg2319 := PrimHead(V305)
reg2320 := PrimTail(V305)
reg2321 := PrimHead(reg2320)
__ctx.TailApply(__defun__shen_4abstraction__build, reg2319, reg2321)
return
} else {
reg2323 := MakeSymbol("shen.abstract_rule")
__ctx.TailApply(__defun__shen_4f__error, reg2323)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.abstract_rule", value: __defun__shen_4abstract__rule})

__defun__shen_4abstraction__build = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V308 := __args[0]
_ = V308
V309 := __args[1]
_ = V309
reg2325 := Nil;
reg2326 := PrimEqual(reg2325, V308)
if reg2326 == True {
__ctx.Return(V309)
return
} else {
reg2327 := PrimIsPair(V308)
if reg2327 == True {
reg2328 := MakeSymbol("/.")
reg2329 := PrimHead(V308)
reg2330 := PrimTail(V308)
reg2331 := __e.Call(__defun__shen_4abstraction__build, reg2330, V309)
reg2332 := Nil;
reg2333 := PrimCons(reg2331, reg2332)
reg2334 := PrimCons(reg2329, reg2333)
reg2335 := PrimCons(reg2328, reg2334)
__ctx.Return(reg2335)
return
} else {
reg2336 := MakeSymbol("shen.abstraction_build")
__ctx.TailApply(__defun__shen_4f__error, reg2336)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.abstraction_build", value: __defun__shen_4abstraction__build})

__defun__shen_4parameters = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V311 := __args[0]
_ = V311
reg2338 := MakeNumber(0)
reg2339 := PrimEqual(reg2338, V311)
if reg2339 == True {
reg2340 := Nil;
__ctx.Return(reg2340)
return
} else {
reg2341 := MakeSymbol("V")
reg2342 := __e.Call(__defun__gensym, reg2341)
reg2343 := MakeNumber(1)
reg2344 := PrimNumberSubtract(V311, reg2343)
reg2345 := __e.Call(__defun__shen_4parameters, reg2344)
reg2346 := PrimCons(reg2342, reg2345)
__ctx.Return(reg2346)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.parameters", value: __defun__shen_4parameters})

__defun__shen_4application__build = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V314 := __args[0]
_ = V314
V315 := __args[1]
_ = V315
reg2347 := Nil;
reg2348 := PrimEqual(reg2347, V314)
if reg2348 == True {
__ctx.Return(V315)
return
} else {
reg2349 := PrimIsPair(V314)
if reg2349 == True {
reg2350 := PrimTail(V314)
reg2351 := PrimHead(V314)
reg2352 := Nil;
reg2353 := PrimCons(reg2351, reg2352)
reg2354 := PrimCons(V315, reg2353)
__ctx.TailApply(__defun__shen_4application__build, reg2350, reg2354)
return
} else {
reg2356 := MakeSymbol("shen.application_build")
__ctx.TailApply(__defun__shen_4f__error, reg2356)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.application_build", value: __defun__shen_4application__build})

__defun__shen_4compile__to__kl = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V318 := __args[0]
_ = V318
V319 := __args[1]
_ = V319
reg2358 := PrimIsPair(V319)
var reg2374 Obj
if reg2358 == True {
reg2359 := PrimTail(V319)
reg2360 := PrimIsPair(reg2359)
var reg2369 Obj
if reg2360 == True {
reg2361 := Nil;
reg2362 := PrimTail(V319)
reg2363 := PrimTail(reg2362)
reg2364 := PrimEqual(reg2361, reg2363)
var reg2367 Obj
if reg2364 == True {
reg2365 := True;
reg2367 = reg2365
} else {
reg2366 := False;
reg2367 = reg2366
}
reg2369 = reg2367
} else {
reg2368 := False;
reg2369 = reg2368
}
var reg2372 Obj
if reg2369 == True {
reg2370 := True;
reg2372 = reg2370
} else {
reg2371 := False;
reg2372 = reg2371
}
reg2374 = reg2372
} else {
reg2373 := False;
reg2374 = reg2373
}
if reg2374 == True {
reg2375 := PrimHead(V319)
reg2376 := __e.Call(__defun__length, reg2375)
reg2377 := __e.Call(__defun__shen_4store_1arity, V318, reg2376)
Arity := reg2377
_ = Arity
reg2378 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4reduce, X)
return
}, 1)
reg2380 := PrimTail(V319)
reg2381 := PrimHead(reg2380)
reg2382 := __e.Call(__defun__map, reg2378, reg2381)
Reduce := reg2382
_ = Reduce
reg2383 := PrimHead(V319)
reg2384 := __e.Call(__defun__shen_4cond_1expression, V318, reg2383, Reduce)
CondExpression := reg2384
_ = CondExpression
reg2385 := MakeSymbol("shen.*optimise*")
reg2386 := PrimValue(reg2385)
var reg2391 Obj
if reg2386 == True {
reg2387 := __e.Call(__defun__shen_4get_1type, V318)
reg2388 := PrimHead(V319)
reg2389 := __e.Call(__defun__shen_4typextable, reg2387, reg2388)
reg2391 = reg2389
} else {
reg2390 := MakeSymbol("shen.skip")
reg2391 = reg2390
}
TypeTable := reg2391
_ = TypeTable
reg2392 := MakeSymbol("shen.*optimise*")
reg2393 := PrimValue(reg2392)
var reg2396 Obj
if reg2393 == True {
reg2394 := PrimHead(V319)
reg2395 := __e.Call(__defun__shen_4assign_1types, reg2394, TypeTable, CondExpression)
reg2396 = reg2395
} else {
reg2396 = CondExpression
}
TypedCondExpression := reg2396
_ = TypedCondExpression
reg2397 := MakeSymbol("defun")
reg2398 := PrimHead(V319)
reg2399 := Nil;
reg2400 := PrimCons(TypedCondExpression, reg2399)
reg2401 := PrimCons(reg2398, reg2400)
reg2402 := PrimCons(V318, reg2401)
reg2403 := PrimCons(reg2397, reg2402)
__ctx.Return(reg2403)
return
} else {
reg2404 := MakeSymbol("shen.compile_to_kl")
__ctx.TailApply(__defun__shen_4f__error, reg2404)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.compile_to_kl", value: __defun__shen_4compile__to__kl})

__defun__shen_4get_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V325 := __args[0]
_ = V325
reg2406 := PrimIsPair(V325)
if reg2406 == True {
reg2407 := MakeSymbol("shen.skip")
__ctx.Return(reg2407)
return
} else {
reg2408 := MakeSymbol("shen.*signedfuncs*")
reg2409 := PrimValue(reg2408)
reg2410 := __e.Call(__defun__assoc, V325, reg2409)
FType := reg2410
_ = FType
reg2411 := __e.Call(__defun__empty_2, FType)
if reg2411 == True {
reg2412 := MakeSymbol("shen.skip")
__ctx.Return(reg2412)
return
} else {
reg2413 := PrimTail(FType)
__ctx.Return(reg2413)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.get-type", value: __defun__shen_4get_1type})

__defun__shen_4typextable = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V336 := __args[0]
_ = V336
V337 := __args[1]
_ = V337
reg2414 := PrimIsPair(V336)
var reg2454 Obj
if reg2414 == True {
reg2415 := PrimTail(V336)
reg2416 := PrimIsPair(reg2415)
var reg2449 Obj
if reg2416 == True {
reg2417 := MakeSymbol("-->")
reg2418 := PrimTail(V336)
reg2419 := PrimHead(reg2418)
reg2420 := PrimEqual(reg2417, reg2419)
var reg2444 Obj
if reg2420 == True {
reg2421 := PrimTail(V336)
reg2422 := PrimTail(reg2421)
reg2423 := PrimIsPair(reg2422)
var reg2439 Obj
if reg2423 == True {
reg2424 := Nil;
reg2425 := PrimTail(V336)
reg2426 := PrimTail(reg2425)
reg2427 := PrimTail(reg2426)
reg2428 := PrimEqual(reg2424, reg2427)
var reg2434 Obj
if reg2428 == True {
reg2429 := PrimIsPair(V337)
var reg2432 Obj
if reg2429 == True {
reg2430 := True;
reg2432 = reg2430
} else {
reg2431 := False;
reg2432 = reg2431
}
reg2434 = reg2432
} else {
reg2433 := False;
reg2434 = reg2433
}
var reg2437 Obj
if reg2434 == True {
reg2435 := True;
reg2437 = reg2435
} else {
reg2436 := False;
reg2437 = reg2436
}
reg2439 = reg2437
} else {
reg2438 := False;
reg2439 = reg2438
}
var reg2442 Obj
if reg2439 == True {
reg2440 := True;
reg2442 = reg2440
} else {
reg2441 := False;
reg2442 = reg2441
}
reg2444 = reg2442
} else {
reg2443 := False;
reg2444 = reg2443
}
var reg2447 Obj
if reg2444 == True {
reg2445 := True;
reg2447 = reg2445
} else {
reg2446 := False;
reg2447 = reg2446
}
reg2449 = reg2447
} else {
reg2448 := False;
reg2449 = reg2448
}
var reg2452 Obj
if reg2449 == True {
reg2450 := True;
reg2452 = reg2450
} else {
reg2451 := False;
reg2452 = reg2451
}
reg2454 = reg2452
} else {
reg2453 := False;
reg2454 = reg2453
}
if reg2454 == True {
reg2455 := PrimHead(V336)
reg2456 := PrimIsVariable(reg2455)
if reg2456 == True {
reg2457 := PrimTail(V336)
reg2458 := PrimTail(reg2457)
reg2459 := PrimHead(reg2458)
reg2460 := PrimTail(V337)
__ctx.TailApply(__defun__shen_4typextable, reg2459, reg2460)
return
} else {
reg2462 := PrimHead(V337)
reg2463 := PrimHead(V336)
reg2464 := PrimCons(reg2462, reg2463)
reg2465 := PrimTail(V336)
reg2466 := PrimTail(reg2465)
reg2467 := PrimHead(reg2466)
reg2468 := PrimTail(V337)
reg2469 := __e.Call(__defun__shen_4typextable, reg2467, reg2468)
reg2470 := PrimCons(reg2464, reg2469)
__ctx.Return(reg2470)
return
}
} else {
reg2471 := Nil;
__ctx.Return(reg2471)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.typextable", value: __defun__shen_4typextable})

__defun__shen_4assign_1types = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V341 := __args[0]
_ = V341
V342 := __args[1]
_ = V342
V343 := __args[2]
_ = V343
reg2472 := PrimIsPair(V343)
var reg2515 Obj
if reg2472 == True {
reg2473 := MakeSymbol("let")
reg2474 := PrimHead(V343)
reg2475 := PrimEqual(reg2473, reg2474)
var reg2510 Obj
if reg2475 == True {
reg2476 := PrimTail(V343)
reg2477 := PrimIsPair(reg2476)
var reg2505 Obj
if reg2477 == True {
reg2478 := PrimTail(V343)
reg2479 := PrimTail(reg2478)
reg2480 := PrimIsPair(reg2479)
var reg2500 Obj
if reg2480 == True {
reg2481 := PrimTail(V343)
reg2482 := PrimTail(reg2481)
reg2483 := PrimTail(reg2482)
reg2484 := PrimIsPair(reg2483)
var reg2495 Obj
if reg2484 == True {
reg2485 := Nil;
reg2486 := PrimTail(V343)
reg2487 := PrimTail(reg2486)
reg2488 := PrimTail(reg2487)
reg2489 := PrimTail(reg2488)
reg2490 := PrimEqual(reg2485, reg2489)
var reg2493 Obj
if reg2490 == True {
reg2491 := True;
reg2493 = reg2491
} else {
reg2492 := False;
reg2493 = reg2492
}
reg2495 = reg2493
} else {
reg2494 := False;
reg2495 = reg2494
}
var reg2498 Obj
if reg2495 == True {
reg2496 := True;
reg2498 = reg2496
} else {
reg2497 := False;
reg2498 = reg2497
}
reg2500 = reg2498
} else {
reg2499 := False;
reg2500 = reg2499
}
var reg2503 Obj
if reg2500 == True {
reg2501 := True;
reg2503 = reg2501
} else {
reg2502 := False;
reg2503 = reg2502
}
reg2505 = reg2503
} else {
reg2504 := False;
reg2505 = reg2504
}
var reg2508 Obj
if reg2505 == True {
reg2506 := True;
reg2508 = reg2506
} else {
reg2507 := False;
reg2508 = reg2507
}
reg2510 = reg2508
} else {
reg2509 := False;
reg2510 = reg2509
}
var reg2513 Obj
if reg2510 == True {
reg2511 := True;
reg2513 = reg2511
} else {
reg2512 := False;
reg2513 = reg2512
}
reg2515 = reg2513
} else {
reg2514 := False;
reg2515 = reg2514
}
if reg2515 == True {
reg2516 := MakeSymbol("let")
reg2517 := PrimTail(V343)
reg2518 := PrimHead(reg2517)
reg2519 := PrimTail(V343)
reg2520 := PrimTail(reg2519)
reg2521 := PrimHead(reg2520)
reg2522 := __e.Call(__defun__shen_4assign_1types, V341, V342, reg2521)
reg2523 := PrimTail(V343)
reg2524 := PrimHead(reg2523)
reg2525 := PrimCons(reg2524, V341)
reg2526 := PrimTail(V343)
reg2527 := PrimTail(reg2526)
reg2528 := PrimTail(reg2527)
reg2529 := PrimHead(reg2528)
reg2530 := __e.Call(__defun__shen_4assign_1types, reg2525, V342, reg2529)
reg2531 := Nil;
reg2532 := PrimCons(reg2530, reg2531)
reg2533 := PrimCons(reg2522, reg2532)
reg2534 := PrimCons(reg2518, reg2533)
reg2535 := PrimCons(reg2516, reg2534)
__ctx.Return(reg2535)
return
} else {
reg2536 := PrimIsPair(V343)
var reg2569 Obj
if reg2536 == True {
reg2537 := MakeSymbol("lambda")
reg2538 := PrimHead(V343)
reg2539 := PrimEqual(reg2537, reg2538)
var reg2564 Obj
if reg2539 == True {
reg2540 := PrimTail(V343)
reg2541 := PrimIsPair(reg2540)
var reg2559 Obj
if reg2541 == True {
reg2542 := PrimTail(V343)
reg2543 := PrimTail(reg2542)
reg2544 := PrimIsPair(reg2543)
var reg2554 Obj
if reg2544 == True {
reg2545 := Nil;
reg2546 := PrimTail(V343)
reg2547 := PrimTail(reg2546)
reg2548 := PrimTail(reg2547)
reg2549 := PrimEqual(reg2545, reg2548)
var reg2552 Obj
if reg2549 == True {
reg2550 := True;
reg2552 = reg2550
} else {
reg2551 := False;
reg2552 = reg2551
}
reg2554 = reg2552
} else {
reg2553 := False;
reg2554 = reg2553
}
var reg2557 Obj
if reg2554 == True {
reg2555 := True;
reg2557 = reg2555
} else {
reg2556 := False;
reg2557 = reg2556
}
reg2559 = reg2557
} else {
reg2558 := False;
reg2559 = reg2558
}
var reg2562 Obj
if reg2559 == True {
reg2560 := True;
reg2562 = reg2560
} else {
reg2561 := False;
reg2562 = reg2561
}
reg2564 = reg2562
} else {
reg2563 := False;
reg2564 = reg2563
}
var reg2567 Obj
if reg2564 == True {
reg2565 := True;
reg2567 = reg2565
} else {
reg2566 := False;
reg2567 = reg2566
}
reg2569 = reg2567
} else {
reg2568 := False;
reg2569 = reg2568
}
if reg2569 == True {
reg2570 := MakeSymbol("lambda")
reg2571 := PrimTail(V343)
reg2572 := PrimHead(reg2571)
reg2573 := PrimTail(V343)
reg2574 := PrimHead(reg2573)
reg2575 := PrimCons(reg2574, V341)
reg2576 := PrimTail(V343)
reg2577 := PrimTail(reg2576)
reg2578 := PrimHead(reg2577)
reg2579 := __e.Call(__defun__shen_4assign_1types, reg2575, V342, reg2578)
reg2580 := Nil;
reg2581 := PrimCons(reg2579, reg2580)
reg2582 := PrimCons(reg2572, reg2581)
reg2583 := PrimCons(reg2570, reg2582)
__ctx.Return(reg2583)
return
} else {
reg2584 := PrimIsPair(V343)
var reg2592 Obj
if reg2584 == True {
reg2585 := MakeSymbol("cond")
reg2586 := PrimHead(V343)
reg2587 := PrimEqual(reg2585, reg2586)
var reg2590 Obj
if reg2587 == True {
reg2588 := True;
reg2590 = reg2588
} else {
reg2589 := False;
reg2590 = reg2589
}
reg2592 = reg2590
} else {
reg2591 := False;
reg2592 = reg2591
}
if reg2592 == True {
reg2593 := MakeSymbol("cond")
reg2594 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
reg2595 := PrimHead(Y)
reg2596 := __e.Call(__defun__shen_4assign_1types, V341, V342, reg2595)
reg2597 := PrimTail(Y)
reg2598 := PrimHead(reg2597)
reg2599 := __e.Call(__defun__shen_4assign_1types, V341, V342, reg2598)
reg2600 := Nil;
reg2601 := PrimCons(reg2599, reg2600)
reg2602 := PrimCons(reg2596, reg2601)
__ctx.Return(reg2602)
return
}, 1)
reg2603 := PrimTail(V343)
reg2604 := __e.Call(__defun__map, reg2594, reg2603)
reg2605 := PrimCons(reg2593, reg2604)
__ctx.Return(reg2605)
return
} else {
reg2606 := PrimIsPair(V343)
if reg2606 == True {
reg2607 := PrimHead(V343)
reg2608 := __e.Call(__defun__shen_4get_1type, reg2607)
reg2609 := PrimTail(V343)
reg2610 := __e.Call(__defun__shen_4typextable, reg2608, reg2609)
NewTable := reg2610
_ = NewTable
reg2611 := PrimHead(V343)
reg2612 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Y := __args[0]
_ = Y
reg2613 := __e.Call(__defun__append, V342, NewTable)
__ctx.TailApply(__defun__shen_4assign_1types, V341, reg2613, Y)
return
}, 1)
reg2615 := PrimTail(V343)
reg2616 := __e.Call(__defun__map, reg2612, reg2615)
reg2617 := PrimCons(reg2611, reg2616)
__ctx.Return(reg2617)
return
} else {
reg2618 := __e.Call(__defun__assoc, V343, V342)
AtomType := reg2618
_ = AtomType
reg2619 := PrimIsPair(AtomType)
if reg2619 == True {
reg2620 := MakeSymbol("type")
reg2621 := PrimTail(AtomType)
reg2622 := Nil;
reg2623 := PrimCons(reg2621, reg2622)
reg2624 := PrimCons(V343, reg2623)
reg2625 := PrimCons(reg2620, reg2624)
__ctx.Return(reg2625)
return
} else {
reg2626 := __e.Call(__defun__element_2, V343, V341)
if reg2626 == True {
__ctx.Return(V343)
return
} else {
__ctx.TailApply(__defun__shen_4atom_1type, V343)
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
V345 := __args[0]
_ = V345
reg2628 := PrimIsString(V345)
if reg2628 == True {
reg2629 := MakeSymbol("type")
reg2630 := MakeSymbol("string")
reg2631 := Nil;
reg2632 := PrimCons(reg2630, reg2631)
reg2633 := PrimCons(V345, reg2632)
reg2634 := PrimCons(reg2629, reg2633)
__ctx.Return(reg2634)
return
} else {
reg2635 := PrimIsNumber(V345)
if reg2635 == True {
reg2636 := MakeSymbol("type")
reg2637 := MakeSymbol("number")
reg2638 := Nil;
reg2639 := PrimCons(reg2637, reg2638)
reg2640 := PrimCons(V345, reg2639)
reg2641 := PrimCons(reg2636, reg2640)
__ctx.Return(reg2641)
return
} else {
reg2642 := __e.Call(__defun__boolean_2, V345)
if reg2642 == True {
reg2643 := MakeSymbol("type")
reg2644 := MakeSymbol("boolean")
reg2645 := Nil;
reg2646 := PrimCons(reg2644, reg2645)
reg2647 := PrimCons(V345, reg2646)
reg2648 := PrimCons(reg2643, reg2647)
__ctx.Return(reg2648)
return
} else {
reg2649 := PrimIsSymbol(V345)
if reg2649 == True {
reg2650 := MakeSymbol("type")
reg2651 := MakeSymbol("symbol")
reg2652 := Nil;
reg2653 := PrimCons(reg2651, reg2652)
reg2654 := PrimCons(V345, reg2653)
reg2655 := PrimCons(reg2650, reg2654)
__ctx.Return(reg2655)
return
} else {
__ctx.Return(V345)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.atom-type", value: __defun__shen_4atom_1type})

__defun__shen_4store_1arity = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V350 := __args[0]
_ = V350
V351 := __args[1]
_ = V351
reg2656 := MakeSymbol("shen.*installing-kl*")
reg2657 := PrimValue(reg2656)
if reg2657 == True {
reg2658 := MakeSymbol("shen.skip")
__ctx.Return(reg2658)
return
} else {
reg2659 := MakeSymbol("arity")
reg2660 := MakeSymbol("*property-vector*")
reg2661 := PrimValue(reg2660)
__ctx.TailApply(__defun__put, V350, reg2659, V351, reg2661)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.store-arity", value: __defun__shen_4store_1arity})

__defun__shen_4reduce = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V353 := __args[0]
_ = V353
reg2663 := MakeSymbol("shen.*teststack*")
reg2664 := Nil;
reg2665 := PrimSet(reg2663, reg2664)
_ = reg2665
reg2666 := __e.Call(__defun__shen_4reduce__help, V353)
Result := reg2666
_ = Result
reg2667 := MakeSymbol(":")
reg2668 := MakeSymbol("shen.tests")
reg2669 := MakeSymbol("shen.*teststack*")
reg2670 := PrimValue(reg2669)
reg2671 := __e.Call(__defun__reverse, reg2670)
reg2672 := PrimCons(reg2668, reg2671)
reg2673 := PrimCons(reg2667, reg2672)
reg2674 := Nil;
reg2675 := PrimCons(Result, reg2674)
reg2676 := PrimCons(reg2673, reg2675)
__ctx.Return(reg2676)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.reduce", value: __defun__shen_4reduce})

__defun__shen_4reduce__help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V355 := __args[0]
_ = V355
reg2677 := PrimIsPair(V355)
var reg2791 Obj
if reg2677 == True {
reg2678 := PrimHead(V355)
reg2679 := PrimIsPair(reg2678)
var reg2786 Obj
if reg2679 == True {
reg2680 := MakeSymbol("/.")
reg2681 := PrimHead(V355)
reg2682 := PrimHead(reg2681)
reg2683 := PrimEqual(reg2680, reg2682)
var reg2781 Obj
if reg2683 == True {
reg2684 := PrimHead(V355)
reg2685 := PrimTail(reg2684)
reg2686 := PrimIsPair(reg2685)
var reg2776 Obj
if reg2686 == True {
reg2687 := PrimHead(V355)
reg2688 := PrimTail(reg2687)
reg2689 := PrimHead(reg2688)
reg2690 := PrimIsPair(reg2689)
var reg2771 Obj
if reg2690 == True {
reg2691 := MakeSymbol("cons")
reg2692 := PrimHead(V355)
reg2693 := PrimTail(reg2692)
reg2694 := PrimHead(reg2693)
reg2695 := PrimHead(reg2694)
reg2696 := PrimEqual(reg2691, reg2695)
var reg2766 Obj
if reg2696 == True {
reg2697 := PrimHead(V355)
reg2698 := PrimTail(reg2697)
reg2699 := PrimHead(reg2698)
reg2700 := PrimTail(reg2699)
reg2701 := PrimIsPair(reg2700)
var reg2761 Obj
if reg2701 == True {
reg2702 := PrimHead(V355)
reg2703 := PrimTail(reg2702)
reg2704 := PrimHead(reg2703)
reg2705 := PrimTail(reg2704)
reg2706 := PrimTail(reg2705)
reg2707 := PrimIsPair(reg2706)
var reg2756 Obj
if reg2707 == True {
reg2708 := Nil;
reg2709 := PrimHead(V355)
reg2710 := PrimTail(reg2709)
reg2711 := PrimHead(reg2710)
reg2712 := PrimTail(reg2711)
reg2713 := PrimTail(reg2712)
reg2714 := PrimTail(reg2713)
reg2715 := PrimEqual(reg2708, reg2714)
var reg2751 Obj
if reg2715 == True {
reg2716 := PrimHead(V355)
reg2717 := PrimTail(reg2716)
reg2718 := PrimTail(reg2717)
reg2719 := PrimIsPair(reg2718)
var reg2746 Obj
if reg2719 == True {
reg2720 := Nil;
reg2721 := PrimHead(V355)
reg2722 := PrimTail(reg2721)
reg2723 := PrimTail(reg2722)
reg2724 := PrimTail(reg2723)
reg2725 := PrimEqual(reg2720, reg2724)
var reg2741 Obj
if reg2725 == True {
reg2726 := PrimTail(V355)
reg2727 := PrimIsPair(reg2726)
var reg2736 Obj
if reg2727 == True {
reg2728 := Nil;
reg2729 := PrimTail(V355)
reg2730 := PrimTail(reg2729)
reg2731 := PrimEqual(reg2728, reg2730)
var reg2734 Obj
if reg2731 == True {
reg2732 := True;
reg2734 = reg2732
} else {
reg2733 := False;
reg2734 = reg2733
}
reg2736 = reg2734
} else {
reg2735 := False;
reg2736 = reg2735
}
var reg2739 Obj
if reg2736 == True {
reg2737 := True;
reg2739 = reg2737
} else {
reg2738 := False;
reg2739 = reg2738
}
reg2741 = reg2739
} else {
reg2740 := False;
reg2741 = reg2740
}
var reg2744 Obj
if reg2741 == True {
reg2742 := True;
reg2744 = reg2742
} else {
reg2743 := False;
reg2744 = reg2743
}
reg2746 = reg2744
} else {
reg2745 := False;
reg2746 = reg2745
}
var reg2749 Obj
if reg2746 == True {
reg2747 := True;
reg2749 = reg2747
} else {
reg2748 := False;
reg2749 = reg2748
}
reg2751 = reg2749
} else {
reg2750 := False;
reg2751 = reg2750
}
var reg2754 Obj
if reg2751 == True {
reg2752 := True;
reg2754 = reg2752
} else {
reg2753 := False;
reg2754 = reg2753
}
reg2756 = reg2754
} else {
reg2755 := False;
reg2756 = reg2755
}
var reg2759 Obj
if reg2756 == True {
reg2757 := True;
reg2759 = reg2757
} else {
reg2758 := False;
reg2759 = reg2758
}
reg2761 = reg2759
} else {
reg2760 := False;
reg2761 = reg2760
}
var reg2764 Obj
if reg2761 == True {
reg2762 := True;
reg2764 = reg2762
} else {
reg2763 := False;
reg2764 = reg2763
}
reg2766 = reg2764
} else {
reg2765 := False;
reg2766 = reg2765
}
var reg2769 Obj
if reg2766 == True {
reg2767 := True;
reg2769 = reg2767
} else {
reg2768 := False;
reg2769 = reg2768
}
reg2771 = reg2769
} else {
reg2770 := False;
reg2771 = reg2770
}
var reg2774 Obj
if reg2771 == True {
reg2772 := True;
reg2774 = reg2772
} else {
reg2773 := False;
reg2774 = reg2773
}
reg2776 = reg2774
} else {
reg2775 := False;
reg2776 = reg2775
}
var reg2779 Obj
if reg2776 == True {
reg2777 := True;
reg2779 = reg2777
} else {
reg2778 := False;
reg2779 = reg2778
}
reg2781 = reg2779
} else {
reg2780 := False;
reg2781 = reg2780
}
var reg2784 Obj
if reg2781 == True {
reg2782 := True;
reg2784 = reg2782
} else {
reg2783 := False;
reg2784 = reg2783
}
reg2786 = reg2784
} else {
reg2785 := False;
reg2786 = reg2785
}
var reg2789 Obj
if reg2786 == True {
reg2787 := True;
reg2789 = reg2787
} else {
reg2788 := False;
reg2789 = reg2788
}
reg2791 = reg2789
} else {
reg2790 := False;
reg2791 = reg2790
}
if reg2791 == True {
reg2792 := MakeSymbol("cons?")
reg2793 := PrimTail(V355)
reg2794 := PrimCons(reg2792, reg2793)
reg2795 := __e.Call(__defun__shen_4add__test, reg2794)
_ = reg2795
reg2796 := MakeSymbol("/.")
reg2797 := PrimHead(V355)
reg2798 := PrimTail(reg2797)
reg2799 := PrimHead(reg2798)
reg2800 := PrimTail(reg2799)
reg2801 := PrimHead(reg2800)
reg2802 := MakeSymbol("/.")
reg2803 := PrimHead(V355)
reg2804 := PrimTail(reg2803)
reg2805 := PrimHead(reg2804)
reg2806 := PrimTail(reg2805)
reg2807 := PrimTail(reg2806)
reg2808 := PrimHead(reg2807)
reg2809 := PrimTail(V355)
reg2810 := PrimHead(reg2809)
reg2811 := PrimHead(V355)
reg2812 := PrimTail(reg2811)
reg2813 := PrimHead(reg2812)
reg2814 := PrimHead(V355)
reg2815 := PrimTail(reg2814)
reg2816 := PrimTail(reg2815)
reg2817 := PrimHead(reg2816)
reg2818 := __e.Call(__defun__shen_4ebr, reg2810, reg2813, reg2817)
reg2819 := Nil;
reg2820 := PrimCons(reg2818, reg2819)
reg2821 := PrimCons(reg2808, reg2820)
reg2822 := PrimCons(reg2802, reg2821)
reg2823 := Nil;
reg2824 := PrimCons(reg2822, reg2823)
reg2825 := PrimCons(reg2801, reg2824)
reg2826 := PrimCons(reg2796, reg2825)
Abstraction := reg2826
_ = Abstraction
reg2827 := MakeSymbol("hd")
reg2828 := PrimTail(V355)
reg2829 := PrimCons(reg2827, reg2828)
reg2830 := Nil;
reg2831 := PrimCons(reg2829, reg2830)
reg2832 := PrimCons(Abstraction, reg2831)
reg2833 := MakeSymbol("tl")
reg2834 := PrimTail(V355)
reg2835 := PrimCons(reg2833, reg2834)
reg2836 := Nil;
reg2837 := PrimCons(reg2835, reg2836)
reg2838 := PrimCons(reg2832, reg2837)
Application := reg2838
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg2840 := PrimIsPair(V355)
var reg2954 Obj
if reg2840 == True {
reg2841 := PrimHead(V355)
reg2842 := PrimIsPair(reg2841)
var reg2949 Obj
if reg2842 == True {
reg2843 := MakeSymbol("/.")
reg2844 := PrimHead(V355)
reg2845 := PrimHead(reg2844)
reg2846 := PrimEqual(reg2843, reg2845)
var reg2944 Obj
if reg2846 == True {
reg2847 := PrimHead(V355)
reg2848 := PrimTail(reg2847)
reg2849 := PrimIsPair(reg2848)
var reg2939 Obj
if reg2849 == True {
reg2850 := PrimHead(V355)
reg2851 := PrimTail(reg2850)
reg2852 := PrimHead(reg2851)
reg2853 := PrimIsPair(reg2852)
var reg2934 Obj
if reg2853 == True {
reg2854 := MakeSymbol("@p")
reg2855 := PrimHead(V355)
reg2856 := PrimTail(reg2855)
reg2857 := PrimHead(reg2856)
reg2858 := PrimHead(reg2857)
reg2859 := PrimEqual(reg2854, reg2858)
var reg2929 Obj
if reg2859 == True {
reg2860 := PrimHead(V355)
reg2861 := PrimTail(reg2860)
reg2862 := PrimHead(reg2861)
reg2863 := PrimTail(reg2862)
reg2864 := PrimIsPair(reg2863)
var reg2924 Obj
if reg2864 == True {
reg2865 := PrimHead(V355)
reg2866 := PrimTail(reg2865)
reg2867 := PrimHead(reg2866)
reg2868 := PrimTail(reg2867)
reg2869 := PrimTail(reg2868)
reg2870 := PrimIsPair(reg2869)
var reg2919 Obj
if reg2870 == True {
reg2871 := Nil;
reg2872 := PrimHead(V355)
reg2873 := PrimTail(reg2872)
reg2874 := PrimHead(reg2873)
reg2875 := PrimTail(reg2874)
reg2876 := PrimTail(reg2875)
reg2877 := PrimTail(reg2876)
reg2878 := PrimEqual(reg2871, reg2877)
var reg2914 Obj
if reg2878 == True {
reg2879 := PrimHead(V355)
reg2880 := PrimTail(reg2879)
reg2881 := PrimTail(reg2880)
reg2882 := PrimIsPair(reg2881)
var reg2909 Obj
if reg2882 == True {
reg2883 := Nil;
reg2884 := PrimHead(V355)
reg2885 := PrimTail(reg2884)
reg2886 := PrimTail(reg2885)
reg2887 := PrimTail(reg2886)
reg2888 := PrimEqual(reg2883, reg2887)
var reg2904 Obj
if reg2888 == True {
reg2889 := PrimTail(V355)
reg2890 := PrimIsPair(reg2889)
var reg2899 Obj
if reg2890 == True {
reg2891 := Nil;
reg2892 := PrimTail(V355)
reg2893 := PrimTail(reg2892)
reg2894 := PrimEqual(reg2891, reg2893)
var reg2897 Obj
if reg2894 == True {
reg2895 := True;
reg2897 = reg2895
} else {
reg2896 := False;
reg2897 = reg2896
}
reg2899 = reg2897
} else {
reg2898 := False;
reg2899 = reg2898
}
var reg2902 Obj
if reg2899 == True {
reg2900 := True;
reg2902 = reg2900
} else {
reg2901 := False;
reg2902 = reg2901
}
reg2904 = reg2902
} else {
reg2903 := False;
reg2904 = reg2903
}
var reg2907 Obj
if reg2904 == True {
reg2905 := True;
reg2907 = reg2905
} else {
reg2906 := False;
reg2907 = reg2906
}
reg2909 = reg2907
} else {
reg2908 := False;
reg2909 = reg2908
}
var reg2912 Obj
if reg2909 == True {
reg2910 := True;
reg2912 = reg2910
} else {
reg2911 := False;
reg2912 = reg2911
}
reg2914 = reg2912
} else {
reg2913 := False;
reg2914 = reg2913
}
var reg2917 Obj
if reg2914 == True {
reg2915 := True;
reg2917 = reg2915
} else {
reg2916 := False;
reg2917 = reg2916
}
reg2919 = reg2917
} else {
reg2918 := False;
reg2919 = reg2918
}
var reg2922 Obj
if reg2919 == True {
reg2920 := True;
reg2922 = reg2920
} else {
reg2921 := False;
reg2922 = reg2921
}
reg2924 = reg2922
} else {
reg2923 := False;
reg2924 = reg2923
}
var reg2927 Obj
if reg2924 == True {
reg2925 := True;
reg2927 = reg2925
} else {
reg2926 := False;
reg2927 = reg2926
}
reg2929 = reg2927
} else {
reg2928 := False;
reg2929 = reg2928
}
var reg2932 Obj
if reg2929 == True {
reg2930 := True;
reg2932 = reg2930
} else {
reg2931 := False;
reg2932 = reg2931
}
reg2934 = reg2932
} else {
reg2933 := False;
reg2934 = reg2933
}
var reg2937 Obj
if reg2934 == True {
reg2935 := True;
reg2937 = reg2935
} else {
reg2936 := False;
reg2937 = reg2936
}
reg2939 = reg2937
} else {
reg2938 := False;
reg2939 = reg2938
}
var reg2942 Obj
if reg2939 == True {
reg2940 := True;
reg2942 = reg2940
} else {
reg2941 := False;
reg2942 = reg2941
}
reg2944 = reg2942
} else {
reg2943 := False;
reg2944 = reg2943
}
var reg2947 Obj
if reg2944 == True {
reg2945 := True;
reg2947 = reg2945
} else {
reg2946 := False;
reg2947 = reg2946
}
reg2949 = reg2947
} else {
reg2948 := False;
reg2949 = reg2948
}
var reg2952 Obj
if reg2949 == True {
reg2950 := True;
reg2952 = reg2950
} else {
reg2951 := False;
reg2952 = reg2951
}
reg2954 = reg2952
} else {
reg2953 := False;
reg2954 = reg2953
}
if reg2954 == True {
reg2955 := MakeSymbol("tuple?")
reg2956 := PrimTail(V355)
reg2957 := PrimCons(reg2955, reg2956)
reg2958 := __e.Call(__defun__shen_4add__test, reg2957)
_ = reg2958
reg2959 := MakeSymbol("/.")
reg2960 := PrimHead(V355)
reg2961 := PrimTail(reg2960)
reg2962 := PrimHead(reg2961)
reg2963 := PrimTail(reg2962)
reg2964 := PrimHead(reg2963)
reg2965 := MakeSymbol("/.")
reg2966 := PrimHead(V355)
reg2967 := PrimTail(reg2966)
reg2968 := PrimHead(reg2967)
reg2969 := PrimTail(reg2968)
reg2970 := PrimTail(reg2969)
reg2971 := PrimHead(reg2970)
reg2972 := PrimTail(V355)
reg2973 := PrimHead(reg2972)
reg2974 := PrimHead(V355)
reg2975 := PrimTail(reg2974)
reg2976 := PrimHead(reg2975)
reg2977 := PrimHead(V355)
reg2978 := PrimTail(reg2977)
reg2979 := PrimTail(reg2978)
reg2980 := PrimHead(reg2979)
reg2981 := __e.Call(__defun__shen_4ebr, reg2973, reg2976, reg2980)
reg2982 := Nil;
reg2983 := PrimCons(reg2981, reg2982)
reg2984 := PrimCons(reg2971, reg2983)
reg2985 := PrimCons(reg2965, reg2984)
reg2986 := Nil;
reg2987 := PrimCons(reg2985, reg2986)
reg2988 := PrimCons(reg2964, reg2987)
reg2989 := PrimCons(reg2959, reg2988)
Abstraction := reg2989
_ = Abstraction
reg2990 := MakeSymbol("fst")
reg2991 := PrimTail(V355)
reg2992 := PrimCons(reg2990, reg2991)
reg2993 := Nil;
reg2994 := PrimCons(reg2992, reg2993)
reg2995 := PrimCons(Abstraction, reg2994)
reg2996 := MakeSymbol("snd")
reg2997 := PrimTail(V355)
reg2998 := PrimCons(reg2996, reg2997)
reg2999 := Nil;
reg3000 := PrimCons(reg2998, reg2999)
reg3001 := PrimCons(reg2995, reg3000)
Application := reg3001
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg3003 := PrimIsPair(V355)
var reg3117 Obj
if reg3003 == True {
reg3004 := PrimHead(V355)
reg3005 := PrimIsPair(reg3004)
var reg3112 Obj
if reg3005 == True {
reg3006 := MakeSymbol("/.")
reg3007 := PrimHead(V355)
reg3008 := PrimHead(reg3007)
reg3009 := PrimEqual(reg3006, reg3008)
var reg3107 Obj
if reg3009 == True {
reg3010 := PrimHead(V355)
reg3011 := PrimTail(reg3010)
reg3012 := PrimIsPair(reg3011)
var reg3102 Obj
if reg3012 == True {
reg3013 := PrimHead(V355)
reg3014 := PrimTail(reg3013)
reg3015 := PrimHead(reg3014)
reg3016 := PrimIsPair(reg3015)
var reg3097 Obj
if reg3016 == True {
reg3017 := MakeSymbol("@v")
reg3018 := PrimHead(V355)
reg3019 := PrimTail(reg3018)
reg3020 := PrimHead(reg3019)
reg3021 := PrimHead(reg3020)
reg3022 := PrimEqual(reg3017, reg3021)
var reg3092 Obj
if reg3022 == True {
reg3023 := PrimHead(V355)
reg3024 := PrimTail(reg3023)
reg3025 := PrimHead(reg3024)
reg3026 := PrimTail(reg3025)
reg3027 := PrimIsPair(reg3026)
var reg3087 Obj
if reg3027 == True {
reg3028 := PrimHead(V355)
reg3029 := PrimTail(reg3028)
reg3030 := PrimHead(reg3029)
reg3031 := PrimTail(reg3030)
reg3032 := PrimTail(reg3031)
reg3033 := PrimIsPair(reg3032)
var reg3082 Obj
if reg3033 == True {
reg3034 := Nil;
reg3035 := PrimHead(V355)
reg3036 := PrimTail(reg3035)
reg3037 := PrimHead(reg3036)
reg3038 := PrimTail(reg3037)
reg3039 := PrimTail(reg3038)
reg3040 := PrimTail(reg3039)
reg3041 := PrimEqual(reg3034, reg3040)
var reg3077 Obj
if reg3041 == True {
reg3042 := PrimHead(V355)
reg3043 := PrimTail(reg3042)
reg3044 := PrimTail(reg3043)
reg3045 := PrimIsPair(reg3044)
var reg3072 Obj
if reg3045 == True {
reg3046 := Nil;
reg3047 := PrimHead(V355)
reg3048 := PrimTail(reg3047)
reg3049 := PrimTail(reg3048)
reg3050 := PrimTail(reg3049)
reg3051 := PrimEqual(reg3046, reg3050)
var reg3067 Obj
if reg3051 == True {
reg3052 := PrimTail(V355)
reg3053 := PrimIsPair(reg3052)
var reg3062 Obj
if reg3053 == True {
reg3054 := Nil;
reg3055 := PrimTail(V355)
reg3056 := PrimTail(reg3055)
reg3057 := PrimEqual(reg3054, reg3056)
var reg3060 Obj
if reg3057 == True {
reg3058 := True;
reg3060 = reg3058
} else {
reg3059 := False;
reg3060 = reg3059
}
reg3062 = reg3060
} else {
reg3061 := False;
reg3062 = reg3061
}
var reg3065 Obj
if reg3062 == True {
reg3063 := True;
reg3065 = reg3063
} else {
reg3064 := False;
reg3065 = reg3064
}
reg3067 = reg3065
} else {
reg3066 := False;
reg3067 = reg3066
}
var reg3070 Obj
if reg3067 == True {
reg3068 := True;
reg3070 = reg3068
} else {
reg3069 := False;
reg3070 = reg3069
}
reg3072 = reg3070
} else {
reg3071 := False;
reg3072 = reg3071
}
var reg3075 Obj
if reg3072 == True {
reg3073 := True;
reg3075 = reg3073
} else {
reg3074 := False;
reg3075 = reg3074
}
reg3077 = reg3075
} else {
reg3076 := False;
reg3077 = reg3076
}
var reg3080 Obj
if reg3077 == True {
reg3078 := True;
reg3080 = reg3078
} else {
reg3079 := False;
reg3080 = reg3079
}
reg3082 = reg3080
} else {
reg3081 := False;
reg3082 = reg3081
}
var reg3085 Obj
if reg3082 == True {
reg3083 := True;
reg3085 = reg3083
} else {
reg3084 := False;
reg3085 = reg3084
}
reg3087 = reg3085
} else {
reg3086 := False;
reg3087 = reg3086
}
var reg3090 Obj
if reg3087 == True {
reg3088 := True;
reg3090 = reg3088
} else {
reg3089 := False;
reg3090 = reg3089
}
reg3092 = reg3090
} else {
reg3091 := False;
reg3092 = reg3091
}
var reg3095 Obj
if reg3092 == True {
reg3093 := True;
reg3095 = reg3093
} else {
reg3094 := False;
reg3095 = reg3094
}
reg3097 = reg3095
} else {
reg3096 := False;
reg3097 = reg3096
}
var reg3100 Obj
if reg3097 == True {
reg3098 := True;
reg3100 = reg3098
} else {
reg3099 := False;
reg3100 = reg3099
}
reg3102 = reg3100
} else {
reg3101 := False;
reg3102 = reg3101
}
var reg3105 Obj
if reg3102 == True {
reg3103 := True;
reg3105 = reg3103
} else {
reg3104 := False;
reg3105 = reg3104
}
reg3107 = reg3105
} else {
reg3106 := False;
reg3107 = reg3106
}
var reg3110 Obj
if reg3107 == True {
reg3108 := True;
reg3110 = reg3108
} else {
reg3109 := False;
reg3110 = reg3109
}
reg3112 = reg3110
} else {
reg3111 := False;
reg3112 = reg3111
}
var reg3115 Obj
if reg3112 == True {
reg3113 := True;
reg3115 = reg3113
} else {
reg3114 := False;
reg3115 = reg3114
}
reg3117 = reg3115
} else {
reg3116 := False;
reg3117 = reg3116
}
if reg3117 == True {
reg3118 := MakeSymbol("shen.+vector?")
reg3119 := PrimTail(V355)
reg3120 := PrimCons(reg3118, reg3119)
reg3121 := __e.Call(__defun__shen_4add__test, reg3120)
_ = reg3121
reg3122 := MakeSymbol("/.")
reg3123 := PrimHead(V355)
reg3124 := PrimTail(reg3123)
reg3125 := PrimHead(reg3124)
reg3126 := PrimTail(reg3125)
reg3127 := PrimHead(reg3126)
reg3128 := MakeSymbol("/.")
reg3129 := PrimHead(V355)
reg3130 := PrimTail(reg3129)
reg3131 := PrimHead(reg3130)
reg3132 := PrimTail(reg3131)
reg3133 := PrimTail(reg3132)
reg3134 := PrimHead(reg3133)
reg3135 := PrimTail(V355)
reg3136 := PrimHead(reg3135)
reg3137 := PrimHead(V355)
reg3138 := PrimTail(reg3137)
reg3139 := PrimHead(reg3138)
reg3140 := PrimHead(V355)
reg3141 := PrimTail(reg3140)
reg3142 := PrimTail(reg3141)
reg3143 := PrimHead(reg3142)
reg3144 := __e.Call(__defun__shen_4ebr, reg3136, reg3139, reg3143)
reg3145 := Nil;
reg3146 := PrimCons(reg3144, reg3145)
reg3147 := PrimCons(reg3134, reg3146)
reg3148 := PrimCons(reg3128, reg3147)
reg3149 := Nil;
reg3150 := PrimCons(reg3148, reg3149)
reg3151 := PrimCons(reg3127, reg3150)
reg3152 := PrimCons(reg3122, reg3151)
Abstraction := reg3152
_ = Abstraction
reg3153 := MakeSymbol("hdv")
reg3154 := PrimTail(V355)
reg3155 := PrimCons(reg3153, reg3154)
reg3156 := Nil;
reg3157 := PrimCons(reg3155, reg3156)
reg3158 := PrimCons(Abstraction, reg3157)
reg3159 := MakeSymbol("tlv")
reg3160 := PrimTail(V355)
reg3161 := PrimCons(reg3159, reg3160)
reg3162 := Nil;
reg3163 := PrimCons(reg3161, reg3162)
reg3164 := PrimCons(reg3158, reg3163)
Application := reg3164
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg3166 := PrimIsPair(V355)
var reg3280 Obj
if reg3166 == True {
reg3167 := PrimHead(V355)
reg3168 := PrimIsPair(reg3167)
var reg3275 Obj
if reg3168 == True {
reg3169 := MakeSymbol("/.")
reg3170 := PrimHead(V355)
reg3171 := PrimHead(reg3170)
reg3172 := PrimEqual(reg3169, reg3171)
var reg3270 Obj
if reg3172 == True {
reg3173 := PrimHead(V355)
reg3174 := PrimTail(reg3173)
reg3175 := PrimIsPair(reg3174)
var reg3265 Obj
if reg3175 == True {
reg3176 := PrimHead(V355)
reg3177 := PrimTail(reg3176)
reg3178 := PrimHead(reg3177)
reg3179 := PrimIsPair(reg3178)
var reg3260 Obj
if reg3179 == True {
reg3180 := MakeSymbol("@s")
reg3181 := PrimHead(V355)
reg3182 := PrimTail(reg3181)
reg3183 := PrimHead(reg3182)
reg3184 := PrimHead(reg3183)
reg3185 := PrimEqual(reg3180, reg3184)
var reg3255 Obj
if reg3185 == True {
reg3186 := PrimHead(V355)
reg3187 := PrimTail(reg3186)
reg3188 := PrimHead(reg3187)
reg3189 := PrimTail(reg3188)
reg3190 := PrimIsPair(reg3189)
var reg3250 Obj
if reg3190 == True {
reg3191 := PrimHead(V355)
reg3192 := PrimTail(reg3191)
reg3193 := PrimHead(reg3192)
reg3194 := PrimTail(reg3193)
reg3195 := PrimTail(reg3194)
reg3196 := PrimIsPair(reg3195)
var reg3245 Obj
if reg3196 == True {
reg3197 := Nil;
reg3198 := PrimHead(V355)
reg3199 := PrimTail(reg3198)
reg3200 := PrimHead(reg3199)
reg3201 := PrimTail(reg3200)
reg3202 := PrimTail(reg3201)
reg3203 := PrimTail(reg3202)
reg3204 := PrimEqual(reg3197, reg3203)
var reg3240 Obj
if reg3204 == True {
reg3205 := PrimHead(V355)
reg3206 := PrimTail(reg3205)
reg3207 := PrimTail(reg3206)
reg3208 := PrimIsPair(reg3207)
var reg3235 Obj
if reg3208 == True {
reg3209 := Nil;
reg3210 := PrimHead(V355)
reg3211 := PrimTail(reg3210)
reg3212 := PrimTail(reg3211)
reg3213 := PrimTail(reg3212)
reg3214 := PrimEqual(reg3209, reg3213)
var reg3230 Obj
if reg3214 == True {
reg3215 := PrimTail(V355)
reg3216 := PrimIsPair(reg3215)
var reg3225 Obj
if reg3216 == True {
reg3217 := Nil;
reg3218 := PrimTail(V355)
reg3219 := PrimTail(reg3218)
reg3220 := PrimEqual(reg3217, reg3219)
var reg3223 Obj
if reg3220 == True {
reg3221 := True;
reg3223 = reg3221
} else {
reg3222 := False;
reg3223 = reg3222
}
reg3225 = reg3223
} else {
reg3224 := False;
reg3225 = reg3224
}
var reg3228 Obj
if reg3225 == True {
reg3226 := True;
reg3228 = reg3226
} else {
reg3227 := False;
reg3228 = reg3227
}
reg3230 = reg3228
} else {
reg3229 := False;
reg3230 = reg3229
}
var reg3233 Obj
if reg3230 == True {
reg3231 := True;
reg3233 = reg3231
} else {
reg3232 := False;
reg3233 = reg3232
}
reg3235 = reg3233
} else {
reg3234 := False;
reg3235 = reg3234
}
var reg3238 Obj
if reg3235 == True {
reg3236 := True;
reg3238 = reg3236
} else {
reg3237 := False;
reg3238 = reg3237
}
reg3240 = reg3238
} else {
reg3239 := False;
reg3240 = reg3239
}
var reg3243 Obj
if reg3240 == True {
reg3241 := True;
reg3243 = reg3241
} else {
reg3242 := False;
reg3243 = reg3242
}
reg3245 = reg3243
} else {
reg3244 := False;
reg3245 = reg3244
}
var reg3248 Obj
if reg3245 == True {
reg3246 := True;
reg3248 = reg3246
} else {
reg3247 := False;
reg3248 = reg3247
}
reg3250 = reg3248
} else {
reg3249 := False;
reg3250 = reg3249
}
var reg3253 Obj
if reg3250 == True {
reg3251 := True;
reg3253 = reg3251
} else {
reg3252 := False;
reg3253 = reg3252
}
reg3255 = reg3253
} else {
reg3254 := False;
reg3255 = reg3254
}
var reg3258 Obj
if reg3255 == True {
reg3256 := True;
reg3258 = reg3256
} else {
reg3257 := False;
reg3258 = reg3257
}
reg3260 = reg3258
} else {
reg3259 := False;
reg3260 = reg3259
}
var reg3263 Obj
if reg3260 == True {
reg3261 := True;
reg3263 = reg3261
} else {
reg3262 := False;
reg3263 = reg3262
}
reg3265 = reg3263
} else {
reg3264 := False;
reg3265 = reg3264
}
var reg3268 Obj
if reg3265 == True {
reg3266 := True;
reg3268 = reg3266
} else {
reg3267 := False;
reg3268 = reg3267
}
reg3270 = reg3268
} else {
reg3269 := False;
reg3270 = reg3269
}
var reg3273 Obj
if reg3270 == True {
reg3271 := True;
reg3273 = reg3271
} else {
reg3272 := False;
reg3273 = reg3272
}
reg3275 = reg3273
} else {
reg3274 := False;
reg3275 = reg3274
}
var reg3278 Obj
if reg3275 == True {
reg3276 := True;
reg3278 = reg3276
} else {
reg3277 := False;
reg3278 = reg3277
}
reg3280 = reg3278
} else {
reg3279 := False;
reg3280 = reg3279
}
if reg3280 == True {
reg3281 := MakeSymbol("shen.+string?")
reg3282 := PrimTail(V355)
reg3283 := PrimCons(reg3281, reg3282)
reg3284 := __e.Call(__defun__shen_4add__test, reg3283)
_ = reg3284
reg3285 := MakeSymbol("/.")
reg3286 := PrimHead(V355)
reg3287 := PrimTail(reg3286)
reg3288 := PrimHead(reg3287)
reg3289 := PrimTail(reg3288)
reg3290 := PrimHead(reg3289)
reg3291 := MakeSymbol("/.")
reg3292 := PrimHead(V355)
reg3293 := PrimTail(reg3292)
reg3294 := PrimHead(reg3293)
reg3295 := PrimTail(reg3294)
reg3296 := PrimTail(reg3295)
reg3297 := PrimHead(reg3296)
reg3298 := PrimTail(V355)
reg3299 := PrimHead(reg3298)
reg3300 := PrimHead(V355)
reg3301 := PrimTail(reg3300)
reg3302 := PrimHead(reg3301)
reg3303 := PrimHead(V355)
reg3304 := PrimTail(reg3303)
reg3305 := PrimTail(reg3304)
reg3306 := PrimHead(reg3305)
reg3307 := __e.Call(__defun__shen_4ebr, reg3299, reg3302, reg3306)
reg3308 := Nil;
reg3309 := PrimCons(reg3307, reg3308)
reg3310 := PrimCons(reg3297, reg3309)
reg3311 := PrimCons(reg3291, reg3310)
reg3312 := Nil;
reg3313 := PrimCons(reg3311, reg3312)
reg3314 := PrimCons(reg3290, reg3313)
reg3315 := PrimCons(reg3285, reg3314)
Abstraction := reg3315
_ = Abstraction
reg3316 := MakeSymbol("pos")
reg3317 := PrimTail(V355)
reg3318 := PrimHead(reg3317)
reg3319 := MakeNumber(0)
reg3320 := Nil;
reg3321 := PrimCons(reg3319, reg3320)
reg3322 := PrimCons(reg3318, reg3321)
reg3323 := PrimCons(reg3316, reg3322)
reg3324 := Nil;
reg3325 := PrimCons(reg3323, reg3324)
reg3326 := PrimCons(Abstraction, reg3325)
reg3327 := MakeSymbol("tlstr")
reg3328 := PrimTail(V355)
reg3329 := PrimCons(reg3327, reg3328)
reg3330 := Nil;
reg3331 := PrimCons(reg3329, reg3330)
reg3332 := PrimCons(reg3326, reg3331)
Application := reg3332
_ = Application
__ctx.TailApply(__defun__shen_4reduce__help, Application)
return
} else {
reg3334 := PrimIsPair(V355)
var reg3404 Obj
if reg3334 == True {
reg3335 := PrimHead(V355)
reg3336 := PrimIsPair(reg3335)
var reg3399 Obj
if reg3336 == True {
reg3337 := MakeSymbol("/.")
reg3338 := PrimHead(V355)
reg3339 := PrimHead(reg3338)
reg3340 := PrimEqual(reg3337, reg3339)
var reg3394 Obj
if reg3340 == True {
reg3341 := PrimHead(V355)
reg3342 := PrimTail(reg3341)
reg3343 := PrimIsPair(reg3342)
var reg3389 Obj
if reg3343 == True {
reg3344 := PrimHead(V355)
reg3345 := PrimTail(reg3344)
reg3346 := PrimTail(reg3345)
reg3347 := PrimIsPair(reg3346)
var reg3384 Obj
if reg3347 == True {
reg3348 := Nil;
reg3349 := PrimHead(V355)
reg3350 := PrimTail(reg3349)
reg3351 := PrimTail(reg3350)
reg3352 := PrimTail(reg3351)
reg3353 := PrimEqual(reg3348, reg3352)
var reg3379 Obj
if reg3353 == True {
reg3354 := PrimTail(V355)
reg3355 := PrimIsPair(reg3354)
var reg3374 Obj
if reg3355 == True {
reg3356 := Nil;
reg3357 := PrimTail(V355)
reg3358 := PrimTail(reg3357)
reg3359 := PrimEqual(reg3356, reg3358)
var reg3369 Obj
if reg3359 == True {
reg3360 := PrimHead(V355)
reg3361 := PrimTail(reg3360)
reg3362 := PrimHead(reg3361)
reg3363 := PrimIsVariable(reg3362)
reg3364 := PrimNot(reg3363)
var reg3367 Obj
if reg3364 == True {
reg3365 := True;
reg3367 = reg3365
} else {
reg3366 := False;
reg3367 = reg3366
}
reg3369 = reg3367
} else {
reg3368 := False;
reg3369 = reg3368
}
var reg3372 Obj
if reg3369 == True {
reg3370 := True;
reg3372 = reg3370
} else {
reg3371 := False;
reg3372 = reg3371
}
reg3374 = reg3372
} else {
reg3373 := False;
reg3374 = reg3373
}
var reg3377 Obj
if reg3374 == True {
reg3375 := True;
reg3377 = reg3375
} else {
reg3376 := False;
reg3377 = reg3376
}
reg3379 = reg3377
} else {
reg3378 := False;
reg3379 = reg3378
}
var reg3382 Obj
if reg3379 == True {
reg3380 := True;
reg3382 = reg3380
} else {
reg3381 := False;
reg3382 = reg3381
}
reg3384 = reg3382
} else {
reg3383 := False;
reg3384 = reg3383
}
var reg3387 Obj
if reg3384 == True {
reg3385 := True;
reg3387 = reg3385
} else {
reg3386 := False;
reg3387 = reg3386
}
reg3389 = reg3387
} else {
reg3388 := False;
reg3389 = reg3388
}
var reg3392 Obj
if reg3389 == True {
reg3390 := True;
reg3392 = reg3390
} else {
reg3391 := False;
reg3392 = reg3391
}
reg3394 = reg3392
} else {
reg3393 := False;
reg3394 = reg3393
}
var reg3397 Obj
if reg3394 == True {
reg3395 := True;
reg3397 = reg3395
} else {
reg3396 := False;
reg3397 = reg3396
}
reg3399 = reg3397
} else {
reg3398 := False;
reg3399 = reg3398
}
var reg3402 Obj
if reg3399 == True {
reg3400 := True;
reg3402 = reg3400
} else {
reg3401 := False;
reg3402 = reg3401
}
reg3404 = reg3402
} else {
reg3403 := False;
reg3404 = reg3403
}
if reg3404 == True {
reg3405 := MakeSymbol("=")
reg3406 := PrimHead(V355)
reg3407 := PrimTail(reg3406)
reg3408 := PrimHead(reg3407)
reg3409 := PrimTail(V355)
reg3410 := PrimCons(reg3408, reg3409)
reg3411 := PrimCons(reg3405, reg3410)
reg3412 := __e.Call(__defun__shen_4add__test, reg3411)
_ = reg3412
reg3413 := PrimHead(V355)
reg3414 := PrimTail(reg3413)
reg3415 := PrimTail(reg3414)
reg3416 := PrimHead(reg3415)
__ctx.TailApply(__defun__shen_4reduce__help, reg3416)
return
} else {
reg3418 := PrimIsPair(V355)
var reg3478 Obj
if reg3418 == True {
reg3419 := PrimHead(V355)
reg3420 := PrimIsPair(reg3419)
var reg3473 Obj
if reg3420 == True {
reg3421 := MakeSymbol("/.")
reg3422 := PrimHead(V355)
reg3423 := PrimHead(reg3422)
reg3424 := PrimEqual(reg3421, reg3423)
var reg3468 Obj
if reg3424 == True {
reg3425 := PrimHead(V355)
reg3426 := PrimTail(reg3425)
reg3427 := PrimIsPair(reg3426)
var reg3463 Obj
if reg3427 == True {
reg3428 := PrimHead(V355)
reg3429 := PrimTail(reg3428)
reg3430 := PrimTail(reg3429)
reg3431 := PrimIsPair(reg3430)
var reg3458 Obj
if reg3431 == True {
reg3432 := Nil;
reg3433 := PrimHead(V355)
reg3434 := PrimTail(reg3433)
reg3435 := PrimTail(reg3434)
reg3436 := PrimTail(reg3435)
reg3437 := PrimEqual(reg3432, reg3436)
var reg3453 Obj
if reg3437 == True {
reg3438 := PrimTail(V355)
reg3439 := PrimIsPair(reg3438)
var reg3448 Obj
if reg3439 == True {
reg3440 := Nil;
reg3441 := PrimTail(V355)
reg3442 := PrimTail(reg3441)
reg3443 := PrimEqual(reg3440, reg3442)
var reg3446 Obj
if reg3443 == True {
reg3444 := True;
reg3446 = reg3444
} else {
reg3445 := False;
reg3446 = reg3445
}
reg3448 = reg3446
} else {
reg3447 := False;
reg3448 = reg3447
}
var reg3451 Obj
if reg3448 == True {
reg3449 := True;
reg3451 = reg3449
} else {
reg3450 := False;
reg3451 = reg3450
}
reg3453 = reg3451
} else {
reg3452 := False;
reg3453 = reg3452
}
var reg3456 Obj
if reg3453 == True {
reg3454 := True;
reg3456 = reg3454
} else {
reg3455 := False;
reg3456 = reg3455
}
reg3458 = reg3456
} else {
reg3457 := False;
reg3458 = reg3457
}
var reg3461 Obj
if reg3458 == True {
reg3459 := True;
reg3461 = reg3459
} else {
reg3460 := False;
reg3461 = reg3460
}
reg3463 = reg3461
} else {
reg3462 := False;
reg3463 = reg3462
}
var reg3466 Obj
if reg3463 == True {
reg3464 := True;
reg3466 = reg3464
} else {
reg3465 := False;
reg3466 = reg3465
}
reg3468 = reg3466
} else {
reg3467 := False;
reg3468 = reg3467
}
var reg3471 Obj
if reg3468 == True {
reg3469 := True;
reg3471 = reg3469
} else {
reg3470 := False;
reg3471 = reg3470
}
reg3473 = reg3471
} else {
reg3472 := False;
reg3473 = reg3472
}
var reg3476 Obj
if reg3473 == True {
reg3474 := True;
reg3476 = reg3474
} else {
reg3475 := False;
reg3476 = reg3475
}
reg3478 = reg3476
} else {
reg3477 := False;
reg3478 = reg3477
}
if reg3478 == True {
reg3479 := PrimTail(V355)
reg3480 := PrimHead(reg3479)
reg3481 := PrimHead(V355)
reg3482 := PrimTail(reg3481)
reg3483 := PrimHead(reg3482)
reg3484 := PrimHead(V355)
reg3485 := PrimTail(reg3484)
reg3486 := PrimTail(reg3485)
reg3487 := PrimHead(reg3486)
reg3488 := __e.Call(__defun__shen_4ebr, reg3480, reg3483, reg3487)
__ctx.TailApply(__defun__shen_4reduce__help, reg3488)
return
} else {
reg3490 := PrimIsPair(V355)
var reg3523 Obj
if reg3490 == True {
reg3491 := MakeSymbol("where")
reg3492 := PrimHead(V355)
reg3493 := PrimEqual(reg3491, reg3492)
var reg3518 Obj
if reg3493 == True {
reg3494 := PrimTail(V355)
reg3495 := PrimIsPair(reg3494)
var reg3513 Obj
if reg3495 == True {
reg3496 := PrimTail(V355)
reg3497 := PrimTail(reg3496)
reg3498 := PrimIsPair(reg3497)
var reg3508 Obj
if reg3498 == True {
reg3499 := Nil;
reg3500 := PrimTail(V355)
reg3501 := PrimTail(reg3500)
reg3502 := PrimTail(reg3501)
reg3503 := PrimEqual(reg3499, reg3502)
var reg3506 Obj
if reg3503 == True {
reg3504 := True;
reg3506 = reg3504
} else {
reg3505 := False;
reg3506 = reg3505
}
reg3508 = reg3506
} else {
reg3507 := False;
reg3508 = reg3507
}
var reg3511 Obj
if reg3508 == True {
reg3509 := True;
reg3511 = reg3509
} else {
reg3510 := False;
reg3511 = reg3510
}
reg3513 = reg3511
} else {
reg3512 := False;
reg3513 = reg3512
}
var reg3516 Obj
if reg3513 == True {
reg3514 := True;
reg3516 = reg3514
} else {
reg3515 := False;
reg3516 = reg3515
}
reg3518 = reg3516
} else {
reg3517 := False;
reg3518 = reg3517
}
var reg3521 Obj
if reg3518 == True {
reg3519 := True;
reg3521 = reg3519
} else {
reg3520 := False;
reg3521 = reg3520
}
reg3523 = reg3521
} else {
reg3522 := False;
reg3523 = reg3522
}
if reg3523 == True {
reg3524 := PrimTail(V355)
reg3525 := PrimHead(reg3524)
reg3526 := __e.Call(__defun__shen_4add__test, reg3525)
_ = reg3526
reg3527 := PrimTail(V355)
reg3528 := PrimTail(reg3527)
reg3529 := PrimHead(reg3528)
__ctx.TailApply(__defun__shen_4reduce__help, reg3529)
return
} else {
reg3531 := PrimIsPair(V355)
var reg3547 Obj
if reg3531 == True {
reg3532 := PrimTail(V355)
reg3533 := PrimIsPair(reg3532)
var reg3542 Obj
if reg3533 == True {
reg3534 := Nil;
reg3535 := PrimTail(V355)
reg3536 := PrimTail(reg3535)
reg3537 := PrimEqual(reg3534, reg3536)
var reg3540 Obj
if reg3537 == True {
reg3538 := True;
reg3540 = reg3538
} else {
reg3539 := False;
reg3540 = reg3539
}
reg3542 = reg3540
} else {
reg3541 := False;
reg3542 = reg3541
}
var reg3545 Obj
if reg3542 == True {
reg3543 := True;
reg3545 = reg3543
} else {
reg3544 := False;
reg3545 = reg3544
}
reg3547 = reg3545
} else {
reg3546 := False;
reg3547 = reg3546
}
if reg3547 == True {
reg3548 := PrimHead(V355)
reg3549 := __e.Call(__defun__shen_4reduce__help, reg3548)
Z := reg3549
_ = Z
reg3550 := PrimHead(V355)
reg3551 := PrimEqual(reg3550, Z)
if reg3551 == True {
__ctx.Return(V355)
return
} else {
reg3552 := PrimTail(V355)
reg3553 := PrimCons(Z, reg3552)
__ctx.TailApply(__defun__shen_4reduce__help, reg3553)
return
}
} else {
__ctx.Return(V355)
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
V357 := __args[0]
_ = V357
reg3555 := MakeString("")
reg3556 := PrimEqual(reg3555, V357)
if reg3556 == True {
reg3557 := False;
__ctx.Return(reg3557)
return
} else {
reg3558 := PrimIsString(V357)
__ctx.Return(reg3558)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.+string?", value: __defun__shen_4_7string_2})

__defun__shen_4_7vector_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V359 := __args[0]
_ = V359
reg3559 := PrimIsVector(V359)
if reg3559 == True {
reg3560 := MakeNumber(0)
reg3561 := PrimVectorGet(V359, reg3560)
reg3562 := MakeNumber(0)
reg3563 := PrimGreatThan(reg3561, reg3562)
if reg3563 == True {
reg3564 := True;
__ctx.Return(reg3564)
return
} else {
reg3565 := False;
__ctx.Return(reg3565)
return
}
} else {
reg3566 := False;
__ctx.Return(reg3566)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.+vector?", value: __defun__shen_4_7vector_2})

__defun__shen_4ebr = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V372 := __args[0]
_ = V372
V373 := __args[1]
_ = V373
V374 := __args[2]
_ = V374
reg3567 := PrimEqual(V374, V373)
if reg3567 == True {
__ctx.Return(V372)
return
} else {
reg3568 := PrimIsPair(V374)
var reg3609 Obj
if reg3568 == True {
reg3569 := MakeSymbol("lambda")
reg3570 := PrimHead(V374)
reg3571 := PrimEqual(reg3569, reg3570)
var reg3604 Obj
if reg3571 == True {
reg3572 := PrimTail(V374)
reg3573 := PrimIsPair(reg3572)
var reg3599 Obj
if reg3573 == True {
reg3574 := PrimTail(V374)
reg3575 := PrimTail(reg3574)
reg3576 := PrimIsPair(reg3575)
var reg3594 Obj
if reg3576 == True {
reg3577 := Nil;
reg3578 := PrimTail(V374)
reg3579 := PrimTail(reg3578)
reg3580 := PrimTail(reg3579)
reg3581 := PrimEqual(reg3577, reg3580)
var reg3589 Obj
if reg3581 == True {
reg3582 := PrimTail(V374)
reg3583 := PrimHead(reg3582)
reg3584 := __e.Call(__defun__shen_4clash_2, reg3583, V373)
var reg3587 Obj
if reg3584 == True {
reg3585 := True;
reg3587 = reg3585
} else {
reg3586 := False;
reg3587 = reg3586
}
reg3589 = reg3587
} else {
reg3588 := False;
reg3589 = reg3588
}
var reg3592 Obj
if reg3589 == True {
reg3590 := True;
reg3592 = reg3590
} else {
reg3591 := False;
reg3592 = reg3591
}
reg3594 = reg3592
} else {
reg3593 := False;
reg3594 = reg3593
}
var reg3597 Obj
if reg3594 == True {
reg3595 := True;
reg3597 = reg3595
} else {
reg3596 := False;
reg3597 = reg3596
}
reg3599 = reg3597
} else {
reg3598 := False;
reg3599 = reg3598
}
var reg3602 Obj
if reg3599 == True {
reg3600 := True;
reg3602 = reg3600
} else {
reg3601 := False;
reg3602 = reg3601
}
reg3604 = reg3602
} else {
reg3603 := False;
reg3604 = reg3603
}
var reg3607 Obj
if reg3604 == True {
reg3605 := True;
reg3607 = reg3605
} else {
reg3606 := False;
reg3607 = reg3606
}
reg3609 = reg3607
} else {
reg3608 := False;
reg3609 = reg3608
}
if reg3609 == True {
__ctx.Return(V374)
return
} else {
reg3610 := PrimIsPair(V374)
var reg3661 Obj
if reg3610 == True {
reg3611 := MakeSymbol("let")
reg3612 := PrimHead(V374)
reg3613 := PrimEqual(reg3611, reg3612)
var reg3656 Obj
if reg3613 == True {
reg3614 := PrimTail(V374)
reg3615 := PrimIsPair(reg3614)
var reg3651 Obj
if reg3615 == True {
reg3616 := PrimTail(V374)
reg3617 := PrimTail(reg3616)
reg3618 := PrimIsPair(reg3617)
var reg3646 Obj
if reg3618 == True {
reg3619 := PrimTail(V374)
reg3620 := PrimTail(reg3619)
reg3621 := PrimTail(reg3620)
reg3622 := PrimIsPair(reg3621)
var reg3641 Obj
if reg3622 == True {
reg3623 := Nil;
reg3624 := PrimTail(V374)
reg3625 := PrimTail(reg3624)
reg3626 := PrimTail(reg3625)
reg3627 := PrimTail(reg3626)
reg3628 := PrimEqual(reg3623, reg3627)
var reg3636 Obj
if reg3628 == True {
reg3629 := PrimTail(V374)
reg3630 := PrimHead(reg3629)
reg3631 := __e.Call(__defun__shen_4clash_2, reg3630, V373)
var reg3634 Obj
if reg3631 == True {
reg3632 := True;
reg3634 = reg3632
} else {
reg3633 := False;
reg3634 = reg3633
}
reg3636 = reg3634
} else {
reg3635 := False;
reg3636 = reg3635
}
var reg3639 Obj
if reg3636 == True {
reg3637 := True;
reg3639 = reg3637
} else {
reg3638 := False;
reg3639 = reg3638
}
reg3641 = reg3639
} else {
reg3640 := False;
reg3641 = reg3640
}
var reg3644 Obj
if reg3641 == True {
reg3642 := True;
reg3644 = reg3642
} else {
reg3643 := False;
reg3644 = reg3643
}
reg3646 = reg3644
} else {
reg3645 := False;
reg3646 = reg3645
}
var reg3649 Obj
if reg3646 == True {
reg3647 := True;
reg3649 = reg3647
} else {
reg3648 := False;
reg3649 = reg3648
}
reg3651 = reg3649
} else {
reg3650 := False;
reg3651 = reg3650
}
var reg3654 Obj
if reg3651 == True {
reg3652 := True;
reg3654 = reg3652
} else {
reg3653 := False;
reg3654 = reg3653
}
reg3656 = reg3654
} else {
reg3655 := False;
reg3656 = reg3655
}
var reg3659 Obj
if reg3656 == True {
reg3657 := True;
reg3659 = reg3657
} else {
reg3658 := False;
reg3659 = reg3658
}
reg3661 = reg3659
} else {
reg3660 := False;
reg3661 = reg3660
}
if reg3661 == True {
reg3662 := MakeSymbol("let")
reg3663 := PrimTail(V374)
reg3664 := PrimHead(reg3663)
reg3665 := PrimTail(V374)
reg3666 := PrimTail(reg3665)
reg3667 := PrimHead(reg3666)
reg3668 := __e.Call(__defun__shen_4ebr, V372, V373, reg3667)
reg3669 := PrimTail(V374)
reg3670 := PrimTail(reg3669)
reg3671 := PrimTail(reg3670)
reg3672 := PrimCons(reg3668, reg3671)
reg3673 := PrimCons(reg3664, reg3672)
reg3674 := PrimCons(reg3662, reg3673)
__ctx.Return(reg3674)
return
} else {
reg3675 := PrimIsPair(V374)
if reg3675 == True {
reg3676 := PrimHead(V374)
reg3677 := __e.Call(__defun__shen_4ebr, V372, V373, reg3676)
reg3678 := PrimTail(V374)
reg3679 := __e.Call(__defun__shen_4ebr, V372, V373, reg3678)
reg3680 := PrimCons(reg3677, reg3679)
__ctx.Return(reg3680)
return
} else {
__ctx.Return(V374)
return
}
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.ebr", value: __defun__shen_4ebr})

__defun__shen_4clash_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V386 := __args[0]
_ = V386
V387 := __args[1]
_ = V387
reg3681 := PrimEqual(V387, V386)
if reg3681 == True {
reg3682 := True;
__ctx.Return(reg3682)
return
} else {
reg3683 := PrimIsPair(V387)
if reg3683 == True {
reg3684 := PrimHead(V387)
reg3685 := __e.Call(__defun__shen_4clash_2, V386, reg3684)
if reg3685 == True {
reg3686 := True;
__ctx.Return(reg3686)
return
} else {
reg3687 := PrimTail(V387)
reg3688 := __e.Call(__defun__shen_4clash_2, V386, reg3687)
if reg3688 == True {
reg3689 := True;
__ctx.Return(reg3689)
return
} else {
reg3690 := False;
__ctx.Return(reg3690)
return
}
}
} else {
reg3691 := False;
__ctx.Return(reg3691)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.clash?", value: __defun__shen_4clash_2})

__defun__shen_4add__test = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V389 := __args[0]
_ = V389
reg3692 := MakeSymbol("shen.*teststack*")
reg3693 := MakeSymbol("shen.*teststack*")
reg3694 := PrimValue(reg3693)
reg3695 := PrimCons(V389, reg3694)
reg3696 := PrimSet(reg3692, reg3695)
__ctx.Return(reg3696)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.add_test", value: __defun__shen_4add__test})

__defun__shen_4cond_1expression = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V393 := __args[0]
_ = V393
V394 := __args[1]
_ = V394
V395 := __args[2]
_ = V395
reg3697 := __e.Call(__defun__shen_4err_1condition, V393)
Err := reg3697
_ = Err
reg3698 := __e.Call(__defun__shen_4case_1form, V395, Err)
Cases := reg3698
_ = Cases
reg3699 := __e.Call(__defun__shen_4encode_1choices, Cases, V393)
EncodeChoices := reg3699
_ = EncodeChoices
__ctx.TailApply(__defun__shen_4cond_1form, EncodeChoices)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.cond-expression", value: __defun__shen_4cond_1expression})

__defun__shen_4cond_1form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V399 := __args[0]
_ = V399
reg3701 := PrimIsPair(V399)
var reg3735 Obj
if reg3701 == True {
reg3702 := PrimHead(V399)
reg3703 := PrimIsPair(reg3702)
var reg3730 Obj
if reg3703 == True {
reg3704 := True;
reg3705 := PrimHead(V399)
reg3706 := PrimHead(reg3705)
reg3707 := PrimEqual(reg3704, reg3706)
var reg3725 Obj
if reg3707 == True {
reg3708 := PrimHead(V399)
reg3709 := PrimTail(reg3708)
reg3710 := PrimIsPair(reg3709)
var reg3720 Obj
if reg3710 == True {
reg3711 := Nil;
reg3712 := PrimHead(V399)
reg3713 := PrimTail(reg3712)
reg3714 := PrimTail(reg3713)
reg3715 := PrimEqual(reg3711, reg3714)
var reg3718 Obj
if reg3715 == True {
reg3716 := True;
reg3718 = reg3716
} else {
reg3717 := False;
reg3718 = reg3717
}
reg3720 = reg3718
} else {
reg3719 := False;
reg3720 = reg3719
}
var reg3723 Obj
if reg3720 == True {
reg3721 := True;
reg3723 = reg3721
} else {
reg3722 := False;
reg3723 = reg3722
}
reg3725 = reg3723
} else {
reg3724 := False;
reg3725 = reg3724
}
var reg3728 Obj
if reg3725 == True {
reg3726 := True;
reg3728 = reg3726
} else {
reg3727 := False;
reg3728 = reg3727
}
reg3730 = reg3728
} else {
reg3729 := False;
reg3730 = reg3729
}
var reg3733 Obj
if reg3730 == True {
reg3731 := True;
reg3733 = reg3731
} else {
reg3732 := False;
reg3733 = reg3732
}
reg3735 = reg3733
} else {
reg3734 := False;
reg3735 = reg3734
}
if reg3735 == True {
reg3736 := PrimHead(V399)
reg3737 := PrimTail(reg3736)
reg3738 := PrimHead(reg3737)
__ctx.Return(reg3738)
return
} else {
reg3739 := MakeSymbol("cond")
reg3740 := PrimCons(reg3739, V399)
__ctx.Return(reg3740)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.cond-form", value: __defun__shen_4cond_1form})

__defun__shen_4encode_1choices = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V404 := __args[0]
_ = V404
V405 := __args[1]
_ = V405
reg3741 := Nil;
reg3742 := PrimEqual(reg3741, V404)
if reg3742 == True {
reg3743 := Nil;
__ctx.Return(reg3743)
return
} else {
reg3744 := PrimIsPair(V404)
var reg3828 Obj
if reg3744 == True {
reg3745 := PrimHead(V404)
reg3746 := PrimIsPair(reg3745)
var reg3823 Obj
if reg3746 == True {
reg3747 := True;
reg3748 := PrimHead(V404)
reg3749 := PrimHead(reg3748)
reg3750 := PrimEqual(reg3747, reg3749)
var reg3818 Obj
if reg3750 == True {
reg3751 := PrimHead(V404)
reg3752 := PrimTail(reg3751)
reg3753 := PrimIsPair(reg3752)
var reg3813 Obj
if reg3753 == True {
reg3754 := PrimHead(V404)
reg3755 := PrimTail(reg3754)
reg3756 := PrimHead(reg3755)
reg3757 := PrimIsPair(reg3756)
var reg3808 Obj
if reg3757 == True {
reg3758 := MakeSymbol("shen.choicepoint!")
reg3759 := PrimHead(V404)
reg3760 := PrimTail(reg3759)
reg3761 := PrimHead(reg3760)
reg3762 := PrimHead(reg3761)
reg3763 := PrimEqual(reg3758, reg3762)
var reg3803 Obj
if reg3763 == True {
reg3764 := PrimHead(V404)
reg3765 := PrimTail(reg3764)
reg3766 := PrimHead(reg3765)
reg3767 := PrimTail(reg3766)
reg3768 := PrimIsPair(reg3767)
var reg3798 Obj
if reg3768 == True {
reg3769 := Nil;
reg3770 := PrimHead(V404)
reg3771 := PrimTail(reg3770)
reg3772 := PrimHead(reg3771)
reg3773 := PrimTail(reg3772)
reg3774 := PrimTail(reg3773)
reg3775 := PrimEqual(reg3769, reg3774)
var reg3793 Obj
if reg3775 == True {
reg3776 := Nil;
reg3777 := PrimHead(V404)
reg3778 := PrimTail(reg3777)
reg3779 := PrimTail(reg3778)
reg3780 := PrimEqual(reg3776, reg3779)
var reg3788 Obj
if reg3780 == True {
reg3781 := Nil;
reg3782 := PrimTail(V404)
reg3783 := PrimEqual(reg3781, reg3782)
var reg3786 Obj
if reg3783 == True {
reg3784 := True;
reg3786 = reg3784
} else {
reg3785 := False;
reg3786 = reg3785
}
reg3788 = reg3786
} else {
reg3787 := False;
reg3788 = reg3787
}
var reg3791 Obj
if reg3788 == True {
reg3789 := True;
reg3791 = reg3789
} else {
reg3790 := False;
reg3791 = reg3790
}
reg3793 = reg3791
} else {
reg3792 := False;
reg3793 = reg3792
}
var reg3796 Obj
if reg3793 == True {
reg3794 := True;
reg3796 = reg3794
} else {
reg3795 := False;
reg3796 = reg3795
}
reg3798 = reg3796
} else {
reg3797 := False;
reg3798 = reg3797
}
var reg3801 Obj
if reg3798 == True {
reg3799 := True;
reg3801 = reg3799
} else {
reg3800 := False;
reg3801 = reg3800
}
reg3803 = reg3801
} else {
reg3802 := False;
reg3803 = reg3802
}
var reg3806 Obj
if reg3803 == True {
reg3804 := True;
reg3806 = reg3804
} else {
reg3805 := False;
reg3806 = reg3805
}
reg3808 = reg3806
} else {
reg3807 := False;
reg3808 = reg3807
}
var reg3811 Obj
if reg3808 == True {
reg3809 := True;
reg3811 = reg3809
} else {
reg3810 := False;
reg3811 = reg3810
}
reg3813 = reg3811
} else {
reg3812 := False;
reg3813 = reg3812
}
var reg3816 Obj
if reg3813 == True {
reg3814 := True;
reg3816 = reg3814
} else {
reg3815 := False;
reg3816 = reg3815
}
reg3818 = reg3816
} else {
reg3817 := False;
reg3818 = reg3817
}
var reg3821 Obj
if reg3818 == True {
reg3819 := True;
reg3821 = reg3819
} else {
reg3820 := False;
reg3821 = reg3820
}
reg3823 = reg3821
} else {
reg3822 := False;
reg3823 = reg3822
}
var reg3826 Obj
if reg3823 == True {
reg3824 := True;
reg3826 = reg3824
} else {
reg3825 := False;
reg3826 = reg3825
}
reg3828 = reg3826
} else {
reg3827 := False;
reg3828 = reg3827
}
if reg3828 == True {
reg3829 := True;
reg3830 := MakeSymbol("let")
reg3831 := MakeSymbol("Result")
reg3832 := PrimHead(V404)
reg3833 := PrimTail(reg3832)
reg3834 := PrimHead(reg3833)
reg3835 := PrimTail(reg3834)
reg3836 := PrimHead(reg3835)
reg3837 := MakeSymbol("if")
reg3838 := MakeSymbol("=")
reg3839 := MakeSymbol("Result")
reg3840 := MakeSymbol("fail")
reg3841 := Nil;
reg3842 := PrimCons(reg3840, reg3841)
reg3843 := Nil;
reg3844 := PrimCons(reg3842, reg3843)
reg3845 := PrimCons(reg3839, reg3844)
reg3846 := PrimCons(reg3838, reg3845)
reg3847 := MakeSymbol("shen.*installing-kl*")
reg3848 := PrimValue(reg3847)
var reg3857 Obj
if reg3848 == True {
reg3849 := MakeSymbol("shen.sys-error")
reg3850 := Nil;
reg3851 := PrimCons(V405, reg3850)
reg3852 := PrimCons(reg3849, reg3851)
reg3857 = reg3852
} else {
reg3853 := MakeSymbol("shen.f_error")
reg3854 := Nil;
reg3855 := PrimCons(V405, reg3854)
reg3856 := PrimCons(reg3853, reg3855)
reg3857 = reg3856
}
reg3858 := MakeSymbol("Result")
reg3859 := Nil;
reg3860 := PrimCons(reg3858, reg3859)
reg3861 := PrimCons(reg3857, reg3860)
reg3862 := PrimCons(reg3846, reg3861)
reg3863 := PrimCons(reg3837, reg3862)
reg3864 := Nil;
reg3865 := PrimCons(reg3863, reg3864)
reg3866 := PrimCons(reg3836, reg3865)
reg3867 := PrimCons(reg3831, reg3866)
reg3868 := PrimCons(reg3830, reg3867)
reg3869 := Nil;
reg3870 := PrimCons(reg3868, reg3869)
reg3871 := PrimCons(reg3829, reg3870)
reg3872 := Nil;
reg3873 := PrimCons(reg3871, reg3872)
__ctx.Return(reg3873)
return
} else {
reg3874 := PrimIsPair(V404)
var reg3950 Obj
if reg3874 == True {
reg3875 := PrimHead(V404)
reg3876 := PrimIsPair(reg3875)
var reg3945 Obj
if reg3876 == True {
reg3877 := True;
reg3878 := PrimHead(V404)
reg3879 := PrimHead(reg3878)
reg3880 := PrimEqual(reg3877, reg3879)
var reg3940 Obj
if reg3880 == True {
reg3881 := PrimHead(V404)
reg3882 := PrimTail(reg3881)
reg3883 := PrimIsPair(reg3882)
var reg3935 Obj
if reg3883 == True {
reg3884 := PrimHead(V404)
reg3885 := PrimTail(reg3884)
reg3886 := PrimHead(reg3885)
reg3887 := PrimIsPair(reg3886)
var reg3930 Obj
if reg3887 == True {
reg3888 := MakeSymbol("shen.choicepoint!")
reg3889 := PrimHead(V404)
reg3890 := PrimTail(reg3889)
reg3891 := PrimHead(reg3890)
reg3892 := PrimHead(reg3891)
reg3893 := PrimEqual(reg3888, reg3892)
var reg3925 Obj
if reg3893 == True {
reg3894 := PrimHead(V404)
reg3895 := PrimTail(reg3894)
reg3896 := PrimHead(reg3895)
reg3897 := PrimTail(reg3896)
reg3898 := PrimIsPair(reg3897)
var reg3920 Obj
if reg3898 == True {
reg3899 := Nil;
reg3900 := PrimHead(V404)
reg3901 := PrimTail(reg3900)
reg3902 := PrimHead(reg3901)
reg3903 := PrimTail(reg3902)
reg3904 := PrimTail(reg3903)
reg3905 := PrimEqual(reg3899, reg3904)
var reg3915 Obj
if reg3905 == True {
reg3906 := Nil;
reg3907 := PrimHead(V404)
reg3908 := PrimTail(reg3907)
reg3909 := PrimTail(reg3908)
reg3910 := PrimEqual(reg3906, reg3909)
var reg3913 Obj
if reg3910 == True {
reg3911 := True;
reg3913 = reg3911
} else {
reg3912 := False;
reg3913 = reg3912
}
reg3915 = reg3913
} else {
reg3914 := False;
reg3915 = reg3914
}
var reg3918 Obj
if reg3915 == True {
reg3916 := True;
reg3918 = reg3916
} else {
reg3917 := False;
reg3918 = reg3917
}
reg3920 = reg3918
} else {
reg3919 := False;
reg3920 = reg3919
}
var reg3923 Obj
if reg3920 == True {
reg3921 := True;
reg3923 = reg3921
} else {
reg3922 := False;
reg3923 = reg3922
}
reg3925 = reg3923
} else {
reg3924 := False;
reg3925 = reg3924
}
var reg3928 Obj
if reg3925 == True {
reg3926 := True;
reg3928 = reg3926
} else {
reg3927 := False;
reg3928 = reg3927
}
reg3930 = reg3928
} else {
reg3929 := False;
reg3930 = reg3929
}
var reg3933 Obj
if reg3930 == True {
reg3931 := True;
reg3933 = reg3931
} else {
reg3932 := False;
reg3933 = reg3932
}
reg3935 = reg3933
} else {
reg3934 := False;
reg3935 = reg3934
}
var reg3938 Obj
if reg3935 == True {
reg3936 := True;
reg3938 = reg3936
} else {
reg3937 := False;
reg3938 = reg3937
}
reg3940 = reg3938
} else {
reg3939 := False;
reg3940 = reg3939
}
var reg3943 Obj
if reg3940 == True {
reg3941 := True;
reg3943 = reg3941
} else {
reg3942 := False;
reg3943 = reg3942
}
reg3945 = reg3943
} else {
reg3944 := False;
reg3945 = reg3944
}
var reg3948 Obj
if reg3945 == True {
reg3946 := True;
reg3948 = reg3946
} else {
reg3947 := False;
reg3948 = reg3947
}
reg3950 = reg3948
} else {
reg3949 := False;
reg3950 = reg3949
}
if reg3950 == True {
reg3951 := True;
reg3952 := MakeSymbol("let")
reg3953 := MakeSymbol("Result")
reg3954 := PrimHead(V404)
reg3955 := PrimTail(reg3954)
reg3956 := PrimHead(reg3955)
reg3957 := PrimTail(reg3956)
reg3958 := PrimHead(reg3957)
reg3959 := MakeSymbol("if")
reg3960 := MakeSymbol("=")
reg3961 := MakeSymbol("Result")
reg3962 := MakeSymbol("fail")
reg3963 := Nil;
reg3964 := PrimCons(reg3962, reg3963)
reg3965 := Nil;
reg3966 := PrimCons(reg3964, reg3965)
reg3967 := PrimCons(reg3961, reg3966)
reg3968 := PrimCons(reg3960, reg3967)
reg3969 := PrimTail(V404)
reg3970 := __e.Call(__defun__shen_4encode_1choices, reg3969, V405)
reg3971 := __e.Call(__defun__shen_4cond_1form, reg3970)
reg3972 := MakeSymbol("Result")
reg3973 := Nil;
reg3974 := PrimCons(reg3972, reg3973)
reg3975 := PrimCons(reg3971, reg3974)
reg3976 := PrimCons(reg3968, reg3975)
reg3977 := PrimCons(reg3959, reg3976)
reg3978 := Nil;
reg3979 := PrimCons(reg3977, reg3978)
reg3980 := PrimCons(reg3958, reg3979)
reg3981 := PrimCons(reg3953, reg3980)
reg3982 := PrimCons(reg3952, reg3981)
reg3983 := Nil;
reg3984 := PrimCons(reg3982, reg3983)
reg3985 := PrimCons(reg3951, reg3984)
reg3986 := Nil;
reg3987 := PrimCons(reg3985, reg3986)
__ctx.Return(reg3987)
return
} else {
reg3988 := PrimIsPair(V404)
var reg4055 Obj
if reg3988 == True {
reg3989 := PrimHead(V404)
reg3990 := PrimIsPair(reg3989)
var reg4050 Obj
if reg3990 == True {
reg3991 := PrimHead(V404)
reg3992 := PrimTail(reg3991)
reg3993 := PrimIsPair(reg3992)
var reg4045 Obj
if reg3993 == True {
reg3994 := PrimHead(V404)
reg3995 := PrimTail(reg3994)
reg3996 := PrimHead(reg3995)
reg3997 := PrimIsPair(reg3996)
var reg4040 Obj
if reg3997 == True {
reg3998 := MakeSymbol("shen.choicepoint!")
reg3999 := PrimHead(V404)
reg4000 := PrimTail(reg3999)
reg4001 := PrimHead(reg4000)
reg4002 := PrimHead(reg4001)
reg4003 := PrimEqual(reg3998, reg4002)
var reg4035 Obj
if reg4003 == True {
reg4004 := PrimHead(V404)
reg4005 := PrimTail(reg4004)
reg4006 := PrimHead(reg4005)
reg4007 := PrimTail(reg4006)
reg4008 := PrimIsPair(reg4007)
var reg4030 Obj
if reg4008 == True {
reg4009 := Nil;
reg4010 := PrimHead(V404)
reg4011 := PrimTail(reg4010)
reg4012 := PrimHead(reg4011)
reg4013 := PrimTail(reg4012)
reg4014 := PrimTail(reg4013)
reg4015 := PrimEqual(reg4009, reg4014)
var reg4025 Obj
if reg4015 == True {
reg4016 := Nil;
reg4017 := PrimHead(V404)
reg4018 := PrimTail(reg4017)
reg4019 := PrimTail(reg4018)
reg4020 := PrimEqual(reg4016, reg4019)
var reg4023 Obj
if reg4020 == True {
reg4021 := True;
reg4023 = reg4021
} else {
reg4022 := False;
reg4023 = reg4022
}
reg4025 = reg4023
} else {
reg4024 := False;
reg4025 = reg4024
}
var reg4028 Obj
if reg4025 == True {
reg4026 := True;
reg4028 = reg4026
} else {
reg4027 := False;
reg4028 = reg4027
}
reg4030 = reg4028
} else {
reg4029 := False;
reg4030 = reg4029
}
var reg4033 Obj
if reg4030 == True {
reg4031 := True;
reg4033 = reg4031
} else {
reg4032 := False;
reg4033 = reg4032
}
reg4035 = reg4033
} else {
reg4034 := False;
reg4035 = reg4034
}
var reg4038 Obj
if reg4035 == True {
reg4036 := True;
reg4038 = reg4036
} else {
reg4037 := False;
reg4038 = reg4037
}
reg4040 = reg4038
} else {
reg4039 := False;
reg4040 = reg4039
}
var reg4043 Obj
if reg4040 == True {
reg4041 := True;
reg4043 = reg4041
} else {
reg4042 := False;
reg4043 = reg4042
}
reg4045 = reg4043
} else {
reg4044 := False;
reg4045 = reg4044
}
var reg4048 Obj
if reg4045 == True {
reg4046 := True;
reg4048 = reg4046
} else {
reg4047 := False;
reg4048 = reg4047
}
reg4050 = reg4048
} else {
reg4049 := False;
reg4050 = reg4049
}
var reg4053 Obj
if reg4050 == True {
reg4051 := True;
reg4053 = reg4051
} else {
reg4052 := False;
reg4053 = reg4052
}
reg4055 = reg4053
} else {
reg4054 := False;
reg4055 = reg4054
}
if reg4055 == True {
reg4056 := True;
reg4057 := MakeSymbol("let")
reg4058 := MakeSymbol("Freeze")
reg4059 := MakeSymbol("freeze")
reg4060 := PrimTail(V404)
reg4061 := __e.Call(__defun__shen_4encode_1choices, reg4060, V405)
reg4062 := __e.Call(__defun__shen_4cond_1form, reg4061)
reg4063 := Nil;
reg4064 := PrimCons(reg4062, reg4063)
reg4065 := PrimCons(reg4059, reg4064)
reg4066 := MakeSymbol("if")
reg4067 := PrimHead(V404)
reg4068 := PrimHead(reg4067)
reg4069 := MakeSymbol("let")
reg4070 := MakeSymbol("Result")
reg4071 := PrimHead(V404)
reg4072 := PrimTail(reg4071)
reg4073 := PrimHead(reg4072)
reg4074 := PrimTail(reg4073)
reg4075 := PrimHead(reg4074)
reg4076 := MakeSymbol("if")
reg4077 := MakeSymbol("=")
reg4078 := MakeSymbol("Result")
reg4079 := MakeSymbol("fail")
reg4080 := Nil;
reg4081 := PrimCons(reg4079, reg4080)
reg4082 := Nil;
reg4083 := PrimCons(reg4081, reg4082)
reg4084 := PrimCons(reg4078, reg4083)
reg4085 := PrimCons(reg4077, reg4084)
reg4086 := MakeSymbol("thaw")
reg4087 := MakeSymbol("Freeze")
reg4088 := Nil;
reg4089 := PrimCons(reg4087, reg4088)
reg4090 := PrimCons(reg4086, reg4089)
reg4091 := MakeSymbol("Result")
reg4092 := Nil;
reg4093 := PrimCons(reg4091, reg4092)
reg4094 := PrimCons(reg4090, reg4093)
reg4095 := PrimCons(reg4085, reg4094)
reg4096 := PrimCons(reg4076, reg4095)
reg4097 := Nil;
reg4098 := PrimCons(reg4096, reg4097)
reg4099 := PrimCons(reg4075, reg4098)
reg4100 := PrimCons(reg4070, reg4099)
reg4101 := PrimCons(reg4069, reg4100)
reg4102 := MakeSymbol("thaw")
reg4103 := MakeSymbol("Freeze")
reg4104 := Nil;
reg4105 := PrimCons(reg4103, reg4104)
reg4106 := PrimCons(reg4102, reg4105)
reg4107 := Nil;
reg4108 := PrimCons(reg4106, reg4107)
reg4109 := PrimCons(reg4101, reg4108)
reg4110 := PrimCons(reg4068, reg4109)
reg4111 := PrimCons(reg4066, reg4110)
reg4112 := Nil;
reg4113 := PrimCons(reg4111, reg4112)
reg4114 := PrimCons(reg4065, reg4113)
reg4115 := PrimCons(reg4058, reg4114)
reg4116 := PrimCons(reg4057, reg4115)
reg4117 := Nil;
reg4118 := PrimCons(reg4116, reg4117)
reg4119 := PrimCons(reg4056, reg4118)
reg4120 := Nil;
reg4121 := PrimCons(reg4119, reg4120)
__ctx.Return(reg4121)
return
} else {
reg4122 := PrimIsPair(V404)
var reg4147 Obj
if reg4122 == True {
reg4123 := PrimHead(V404)
reg4124 := PrimIsPair(reg4123)
var reg4142 Obj
if reg4124 == True {
reg4125 := PrimHead(V404)
reg4126 := PrimTail(reg4125)
reg4127 := PrimIsPair(reg4126)
var reg4137 Obj
if reg4127 == True {
reg4128 := Nil;
reg4129 := PrimHead(V404)
reg4130 := PrimTail(reg4129)
reg4131 := PrimTail(reg4130)
reg4132 := PrimEqual(reg4128, reg4131)
var reg4135 Obj
if reg4132 == True {
reg4133 := True;
reg4135 = reg4133
} else {
reg4134 := False;
reg4135 = reg4134
}
reg4137 = reg4135
} else {
reg4136 := False;
reg4137 = reg4136
}
var reg4140 Obj
if reg4137 == True {
reg4138 := True;
reg4140 = reg4138
} else {
reg4139 := False;
reg4140 = reg4139
}
reg4142 = reg4140
} else {
reg4141 := False;
reg4142 = reg4141
}
var reg4145 Obj
if reg4142 == True {
reg4143 := True;
reg4145 = reg4143
} else {
reg4144 := False;
reg4145 = reg4144
}
reg4147 = reg4145
} else {
reg4146 := False;
reg4147 = reg4146
}
if reg4147 == True {
reg4148 := PrimHead(V404)
reg4149 := PrimTail(V404)
reg4150 := __e.Call(__defun__shen_4encode_1choices, reg4149, V405)
reg4151 := PrimCons(reg4148, reg4150)
__ctx.Return(reg4151)
return
} else {
reg4152 := MakeSymbol("shen.encode-choices")
__ctx.TailApply(__defun__shen_4f__error, reg4152)
return
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.encode-choices", value: __defun__shen_4encode_1choices})

__defun__shen_4case_1form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V412 := __args[0]
_ = V412
V413 := __args[1]
_ = V413
reg4154 := Nil;
reg4155 := PrimEqual(reg4154, V412)
if reg4155 == True {
reg4156 := Nil;
reg4157 := PrimCons(V413, reg4156)
__ctx.Return(reg4157)
return
} else {
reg4158 := PrimIsPair(V412)
var reg4274 Obj
if reg4158 == True {
reg4159 := PrimHead(V412)
reg4160 := PrimIsPair(reg4159)
var reg4269 Obj
if reg4160 == True {
reg4161 := PrimHead(V412)
reg4162 := PrimHead(reg4161)
reg4163 := PrimIsPair(reg4162)
var reg4264 Obj
if reg4163 == True {
reg4164 := MakeSymbol(":")
reg4165 := PrimHead(V412)
reg4166 := PrimHead(reg4165)
reg4167 := PrimHead(reg4166)
reg4168 := PrimEqual(reg4164, reg4167)
var reg4259 Obj
if reg4168 == True {
reg4169 := PrimHead(V412)
reg4170 := PrimHead(reg4169)
reg4171 := PrimTail(reg4170)
reg4172 := PrimIsPair(reg4171)
var reg4254 Obj
if reg4172 == True {
reg4173 := MakeSymbol("shen.tests")
reg4174 := PrimHead(V412)
reg4175 := PrimHead(reg4174)
reg4176 := PrimTail(reg4175)
reg4177 := PrimHead(reg4176)
reg4178 := PrimEqual(reg4173, reg4177)
var reg4249 Obj
if reg4178 == True {
reg4179 := Nil;
reg4180 := PrimHead(V412)
reg4181 := PrimHead(reg4180)
reg4182 := PrimTail(reg4181)
reg4183 := PrimTail(reg4182)
reg4184 := PrimEqual(reg4179, reg4183)
var reg4244 Obj
if reg4184 == True {
reg4185 := PrimHead(V412)
reg4186 := PrimTail(reg4185)
reg4187 := PrimIsPair(reg4186)
var reg4239 Obj
if reg4187 == True {
reg4188 := PrimHead(V412)
reg4189 := PrimTail(reg4188)
reg4190 := PrimHead(reg4189)
reg4191 := PrimIsPair(reg4190)
var reg4234 Obj
if reg4191 == True {
reg4192 := MakeSymbol("shen.choicepoint!")
reg4193 := PrimHead(V412)
reg4194 := PrimTail(reg4193)
reg4195 := PrimHead(reg4194)
reg4196 := PrimHead(reg4195)
reg4197 := PrimEqual(reg4192, reg4196)
var reg4229 Obj
if reg4197 == True {
reg4198 := PrimHead(V412)
reg4199 := PrimTail(reg4198)
reg4200 := PrimHead(reg4199)
reg4201 := PrimTail(reg4200)
reg4202 := PrimIsPair(reg4201)
var reg4224 Obj
if reg4202 == True {
reg4203 := Nil;
reg4204 := PrimHead(V412)
reg4205 := PrimTail(reg4204)
reg4206 := PrimHead(reg4205)
reg4207 := PrimTail(reg4206)
reg4208 := PrimTail(reg4207)
reg4209 := PrimEqual(reg4203, reg4208)
var reg4219 Obj
if reg4209 == True {
reg4210 := Nil;
reg4211 := PrimHead(V412)
reg4212 := PrimTail(reg4211)
reg4213 := PrimTail(reg4212)
reg4214 := PrimEqual(reg4210, reg4213)
var reg4217 Obj
if reg4214 == True {
reg4215 := True;
reg4217 = reg4215
} else {
reg4216 := False;
reg4217 = reg4216
}
reg4219 = reg4217
} else {
reg4218 := False;
reg4219 = reg4218
}
var reg4222 Obj
if reg4219 == True {
reg4220 := True;
reg4222 = reg4220
} else {
reg4221 := False;
reg4222 = reg4221
}
reg4224 = reg4222
} else {
reg4223 := False;
reg4224 = reg4223
}
var reg4227 Obj
if reg4224 == True {
reg4225 := True;
reg4227 = reg4225
} else {
reg4226 := False;
reg4227 = reg4226
}
reg4229 = reg4227
} else {
reg4228 := False;
reg4229 = reg4228
}
var reg4232 Obj
if reg4229 == True {
reg4230 := True;
reg4232 = reg4230
} else {
reg4231 := False;
reg4232 = reg4231
}
reg4234 = reg4232
} else {
reg4233 := False;
reg4234 = reg4233
}
var reg4237 Obj
if reg4234 == True {
reg4235 := True;
reg4237 = reg4235
} else {
reg4236 := False;
reg4237 = reg4236
}
reg4239 = reg4237
} else {
reg4238 := False;
reg4239 = reg4238
}
var reg4242 Obj
if reg4239 == True {
reg4240 := True;
reg4242 = reg4240
} else {
reg4241 := False;
reg4242 = reg4241
}
reg4244 = reg4242
} else {
reg4243 := False;
reg4244 = reg4243
}
var reg4247 Obj
if reg4244 == True {
reg4245 := True;
reg4247 = reg4245
} else {
reg4246 := False;
reg4247 = reg4246
}
reg4249 = reg4247
} else {
reg4248 := False;
reg4249 = reg4248
}
var reg4252 Obj
if reg4249 == True {
reg4250 := True;
reg4252 = reg4250
} else {
reg4251 := False;
reg4252 = reg4251
}
reg4254 = reg4252
} else {
reg4253 := False;
reg4254 = reg4253
}
var reg4257 Obj
if reg4254 == True {
reg4255 := True;
reg4257 = reg4255
} else {
reg4256 := False;
reg4257 = reg4256
}
reg4259 = reg4257
} else {
reg4258 := False;
reg4259 = reg4258
}
var reg4262 Obj
if reg4259 == True {
reg4260 := True;
reg4262 = reg4260
} else {
reg4261 := False;
reg4262 = reg4261
}
reg4264 = reg4262
} else {
reg4263 := False;
reg4264 = reg4263
}
var reg4267 Obj
if reg4264 == True {
reg4265 := True;
reg4267 = reg4265
} else {
reg4266 := False;
reg4267 = reg4266
}
reg4269 = reg4267
} else {
reg4268 := False;
reg4269 = reg4268
}
var reg4272 Obj
if reg4269 == True {
reg4270 := True;
reg4272 = reg4270
} else {
reg4271 := False;
reg4272 = reg4271
}
reg4274 = reg4272
} else {
reg4273 := False;
reg4274 = reg4273
}
if reg4274 == True {
reg4275 := True;
reg4276 := PrimHead(V412)
reg4277 := PrimTail(reg4276)
reg4278 := PrimCons(reg4275, reg4277)
reg4279 := PrimTail(V412)
reg4280 := __e.Call(__defun__shen_4case_1form, reg4279, V413)
reg4281 := PrimCons(reg4278, reg4280)
__ctx.Return(reg4281)
return
} else {
reg4282 := PrimIsPair(V412)
var reg4356 Obj
if reg4282 == True {
reg4283 := PrimHead(V412)
reg4284 := PrimIsPair(reg4283)
var reg4351 Obj
if reg4284 == True {
reg4285 := PrimHead(V412)
reg4286 := PrimHead(reg4285)
reg4287 := PrimIsPair(reg4286)
var reg4346 Obj
if reg4287 == True {
reg4288 := MakeSymbol(":")
reg4289 := PrimHead(V412)
reg4290 := PrimHead(reg4289)
reg4291 := PrimHead(reg4290)
reg4292 := PrimEqual(reg4288, reg4291)
var reg4341 Obj
if reg4292 == True {
reg4293 := PrimHead(V412)
reg4294 := PrimHead(reg4293)
reg4295 := PrimTail(reg4294)
reg4296 := PrimIsPair(reg4295)
var reg4336 Obj
if reg4296 == True {
reg4297 := MakeSymbol("shen.tests")
reg4298 := PrimHead(V412)
reg4299 := PrimHead(reg4298)
reg4300 := PrimTail(reg4299)
reg4301 := PrimHead(reg4300)
reg4302 := PrimEqual(reg4297, reg4301)
var reg4331 Obj
if reg4302 == True {
reg4303 := Nil;
reg4304 := PrimHead(V412)
reg4305 := PrimHead(reg4304)
reg4306 := PrimTail(reg4305)
reg4307 := PrimTail(reg4306)
reg4308 := PrimEqual(reg4303, reg4307)
var reg4326 Obj
if reg4308 == True {
reg4309 := PrimHead(V412)
reg4310 := PrimTail(reg4309)
reg4311 := PrimIsPair(reg4310)
var reg4321 Obj
if reg4311 == True {
reg4312 := Nil;
reg4313 := PrimHead(V412)
reg4314 := PrimTail(reg4313)
reg4315 := PrimTail(reg4314)
reg4316 := PrimEqual(reg4312, reg4315)
var reg4319 Obj
if reg4316 == True {
reg4317 := True;
reg4319 = reg4317
} else {
reg4318 := False;
reg4319 = reg4318
}
reg4321 = reg4319
} else {
reg4320 := False;
reg4321 = reg4320
}
var reg4324 Obj
if reg4321 == True {
reg4322 := True;
reg4324 = reg4322
} else {
reg4323 := False;
reg4324 = reg4323
}
reg4326 = reg4324
} else {
reg4325 := False;
reg4326 = reg4325
}
var reg4329 Obj
if reg4326 == True {
reg4327 := True;
reg4329 = reg4327
} else {
reg4328 := False;
reg4329 = reg4328
}
reg4331 = reg4329
} else {
reg4330 := False;
reg4331 = reg4330
}
var reg4334 Obj
if reg4331 == True {
reg4332 := True;
reg4334 = reg4332
} else {
reg4333 := False;
reg4334 = reg4333
}
reg4336 = reg4334
} else {
reg4335 := False;
reg4336 = reg4335
}
var reg4339 Obj
if reg4336 == True {
reg4337 := True;
reg4339 = reg4337
} else {
reg4338 := False;
reg4339 = reg4338
}
reg4341 = reg4339
} else {
reg4340 := False;
reg4341 = reg4340
}
var reg4344 Obj
if reg4341 == True {
reg4342 := True;
reg4344 = reg4342
} else {
reg4343 := False;
reg4344 = reg4343
}
reg4346 = reg4344
} else {
reg4345 := False;
reg4346 = reg4345
}
var reg4349 Obj
if reg4346 == True {
reg4347 := True;
reg4349 = reg4347
} else {
reg4348 := False;
reg4349 = reg4348
}
reg4351 = reg4349
} else {
reg4350 := False;
reg4351 = reg4350
}
var reg4354 Obj
if reg4351 == True {
reg4352 := True;
reg4354 = reg4352
} else {
reg4353 := False;
reg4354 = reg4353
}
reg4356 = reg4354
} else {
reg4355 := False;
reg4356 = reg4355
}
if reg4356 == True {
reg4357 := True;
reg4358 := PrimHead(V412)
reg4359 := PrimTail(reg4358)
reg4360 := PrimCons(reg4357, reg4359)
reg4361 := Nil;
reg4362 := PrimCons(reg4360, reg4361)
__ctx.Return(reg4362)
return
} else {
reg4363 := PrimIsPair(V412)
var reg4426 Obj
if reg4363 == True {
reg4364 := PrimHead(V412)
reg4365 := PrimIsPair(reg4364)
var reg4421 Obj
if reg4365 == True {
reg4366 := PrimHead(V412)
reg4367 := PrimHead(reg4366)
reg4368 := PrimIsPair(reg4367)
var reg4416 Obj
if reg4368 == True {
reg4369 := MakeSymbol(":")
reg4370 := PrimHead(V412)
reg4371 := PrimHead(reg4370)
reg4372 := PrimHead(reg4371)
reg4373 := PrimEqual(reg4369, reg4372)
var reg4411 Obj
if reg4373 == True {
reg4374 := PrimHead(V412)
reg4375 := PrimHead(reg4374)
reg4376 := PrimTail(reg4375)
reg4377 := PrimIsPair(reg4376)
var reg4406 Obj
if reg4377 == True {
reg4378 := MakeSymbol("shen.tests")
reg4379 := PrimHead(V412)
reg4380 := PrimHead(reg4379)
reg4381 := PrimTail(reg4380)
reg4382 := PrimHead(reg4381)
reg4383 := PrimEqual(reg4378, reg4382)
var reg4401 Obj
if reg4383 == True {
reg4384 := PrimHead(V412)
reg4385 := PrimTail(reg4384)
reg4386 := PrimIsPair(reg4385)
var reg4396 Obj
if reg4386 == True {
reg4387 := Nil;
reg4388 := PrimHead(V412)
reg4389 := PrimTail(reg4388)
reg4390 := PrimTail(reg4389)
reg4391 := PrimEqual(reg4387, reg4390)
var reg4394 Obj
if reg4391 == True {
reg4392 := True;
reg4394 = reg4392
} else {
reg4393 := False;
reg4394 = reg4393
}
reg4396 = reg4394
} else {
reg4395 := False;
reg4396 = reg4395
}
var reg4399 Obj
if reg4396 == True {
reg4397 := True;
reg4399 = reg4397
} else {
reg4398 := False;
reg4399 = reg4398
}
reg4401 = reg4399
} else {
reg4400 := False;
reg4401 = reg4400
}
var reg4404 Obj
if reg4401 == True {
reg4402 := True;
reg4404 = reg4402
} else {
reg4403 := False;
reg4404 = reg4403
}
reg4406 = reg4404
} else {
reg4405 := False;
reg4406 = reg4405
}
var reg4409 Obj
if reg4406 == True {
reg4407 := True;
reg4409 = reg4407
} else {
reg4408 := False;
reg4409 = reg4408
}
reg4411 = reg4409
} else {
reg4410 := False;
reg4411 = reg4410
}
var reg4414 Obj
if reg4411 == True {
reg4412 := True;
reg4414 = reg4412
} else {
reg4413 := False;
reg4414 = reg4413
}
reg4416 = reg4414
} else {
reg4415 := False;
reg4416 = reg4415
}
var reg4419 Obj
if reg4416 == True {
reg4417 := True;
reg4419 = reg4417
} else {
reg4418 := False;
reg4419 = reg4418
}
reg4421 = reg4419
} else {
reg4420 := False;
reg4421 = reg4420
}
var reg4424 Obj
if reg4421 == True {
reg4422 := True;
reg4424 = reg4422
} else {
reg4423 := False;
reg4424 = reg4423
}
reg4426 = reg4424
} else {
reg4425 := False;
reg4426 = reg4425
}
if reg4426 == True {
reg4427 := PrimHead(V412)
reg4428 := PrimHead(reg4427)
reg4429 := PrimTail(reg4428)
reg4430 := PrimTail(reg4429)
reg4431 := __e.Call(__defun__shen_4embed_1and, reg4430)
reg4432 := PrimHead(V412)
reg4433 := PrimTail(reg4432)
reg4434 := PrimCons(reg4431, reg4433)
reg4435 := PrimTail(V412)
reg4436 := __e.Call(__defun__shen_4case_1form, reg4435, V413)
reg4437 := PrimCons(reg4434, reg4436)
__ctx.Return(reg4437)
return
} else {
reg4438 := MakeSymbol("shen.case-form")
__ctx.TailApply(__defun__shen_4f__error, reg4438)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.case-form", value: __defun__shen_4case_1form})

__defun__shen_4embed_1and = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V415 := __args[0]
_ = V415
reg4440 := PrimIsPair(V415)
var reg4448 Obj
if reg4440 == True {
reg4441 := Nil;
reg4442 := PrimTail(V415)
reg4443 := PrimEqual(reg4441, reg4442)
var reg4446 Obj
if reg4443 == True {
reg4444 := True;
reg4446 = reg4444
} else {
reg4445 := False;
reg4446 = reg4445
}
reg4448 = reg4446
} else {
reg4447 := False;
reg4448 = reg4447
}
if reg4448 == True {
reg4449 := PrimHead(V415)
__ctx.Return(reg4449)
return
} else {
reg4450 := PrimIsPair(V415)
if reg4450 == True {
reg4451 := MakeSymbol("and")
reg4452 := PrimHead(V415)
reg4453 := PrimTail(V415)
reg4454 := __e.Call(__defun__shen_4embed_1and, reg4453)
reg4455 := Nil;
reg4456 := PrimCons(reg4454, reg4455)
reg4457 := PrimCons(reg4452, reg4456)
reg4458 := PrimCons(reg4451, reg4457)
__ctx.Return(reg4458)
return
} else {
reg4459 := MakeSymbol("shen.embed-and")
__ctx.TailApply(__defun__shen_4f__error, reg4459)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.embed-and", value: __defun__shen_4embed_1and})

__defun__shen_4err_1condition = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V417 := __args[0]
_ = V417
reg4461 := True;
reg4462 := MakeSymbol("shen.f_error")
reg4463 := Nil;
reg4464 := PrimCons(V417, reg4463)
reg4465 := PrimCons(reg4462, reg4464)
reg4466 := Nil;
reg4467 := PrimCons(reg4465, reg4466)
reg4468 := PrimCons(reg4461, reg4467)
__ctx.Return(reg4468)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.err-condition", value: __defun__shen_4err_1condition})

__defun__shen_4sys_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V419 := __args[0]
_ = V419
reg4469 := MakeString("system function ")
reg4470 := MakeString(": unexpected argument\n")
reg4471 := MakeSymbol("shen.a")
reg4472 := __e.Call(__defun__shen_4app, V419, reg4470, reg4471)
reg4473 := PrimStringConcat(reg4469, reg4472)
reg4474 := PrimSimpleError(reg4473)
__ctx.Return(reg4474)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sys-error", value: __defun__shen_4sys_1error})

}
