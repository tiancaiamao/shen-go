package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8773 := MakeNative(func(__e *ControlFlow) {
V5839 := __e.Get(1)
_ = V5839
tmp8774 := MakeNative(func(__e *ControlFlow) {
W5840 := __e.Get(1)
_ = W5840
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5839, W5840, W5840)
return
}, 1)

tmp8775 := MakeNative(func(__e *ControlFlow) {
Z5841 := __e.Get(1)
_ = Z5841
__e.Return(PrimTail(Z5841))
return
}, 1)

tmp8776 := PrimValue(sym_dmacros_d)

tmp8777 := Call(__e, PrimFunc(symmap), tmp8775, tmp8776)


__e.TailApply(tmp8774, tmp8777)
return


}, 1)

tmp8778 := Call(__e, ns2_1set, symmacroexpand, tmp8773)


_ = tmp8778

tmp8779 := MakeNative(func(__e *ControlFlow) {
V5850 := __e.Get(1)
_ = V5850
V5851 := __e.Get(2)
_ = V5851
V5852 := __e.Get(3)
_ = V5852
tmp8789 := PrimEqual(Nil, V5851)

if True == tmp8789 {
__e.Return(V5850)
return
} else {
tmp8787 := PrimIsPair(V5851)

if True == tmp8787 {
tmp8780 := MakeNative(func(__e *ControlFlow) {
W5853 := __e.Get(1)
_ = W5853
tmp8783 := PrimEqual(V5850, W5853)

if True == tmp8783 {
tmp8781 := PrimTail(V5851)

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5850, tmp8781, V5852)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W5853, V5852, V5852)
return
}


}, 1)

tmp8784 := PrimHead(V5851)

tmp8785 := Call(__e, PrimFunc(symshen_4walk), tmp8784, V5850)


__e.TailApply(tmp8780, tmp8785)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
return
}


}


}, 3)

tmp8790 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8779)


_ = tmp8790

tmp8791 := MakeNative(func(__e *ControlFlow) {
V5854 := __e.Get(1)
_ = V5854
V5855 := __e.Get(2)
_ = V5855
tmp8795 := PrimIsPair(V5855)

if True == tmp8795 {
tmp8792 := MakeNative(func(__e *ControlFlow) {
Z5856 := __e.Get(1)
_ = Z5856
__e.TailApply(PrimFunc(symshen_4walk), V5854, Z5856)
return
}, 1)

tmp8793 := Call(__e, PrimFunc(symmap), tmp8792, V5855)


__e.TailApply(V5854, tmp8793)
return


} else {
__e.TailApply(V5854, V5855)
return
}


}, 2)

tmp8796 := Call(__e, ns2_1set, symshen_4walk, tmp8791)


_ = tmp8796

tmp8797 := MakeNative(func(__e *ControlFlow) {
V5857 := __e.Get(1)
_ = V5857
tmp8798 := MakeNative(func(__e *ControlFlow) {
GoTo5858 := __e.Get(1)
_ = GoTo5858
tmp9050 := PrimIsPair(V5857)

if True == tmp9050 {
tmp8799 := MakeNative(func(__e *ControlFlow) {
Select5859 := __e.Get(1)
_ = Select5859
tmp8800 := MakeNative(func(__e *ControlFlow) {
Select5860 := __e.Get(1)
_ = Select5860
tmp9046 := PrimEqual(symdefmacro, Select5859)

var ifres9043 Obj

if True == tmp9046 {
tmp9045 := PrimIsPair(Select5860)

var ifres9044 Obj

if True == tmp9045 {
ifres9044 = True


} else {
ifres9044 = False


}

ifres9043 = ifres9044


} else {
ifres9043 = False


}

if True == ifres9043 {
tmp8801 := PrimHead(Select5860)

tmp8802 := PrimTail(Select5860)

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8801, tmp8802)
return


} else {
tmp9041 := PrimEqual(symdefcc, Select5859)

if True == tmp9041 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select5860)
return
} else {
tmp9039 := PrimEqual(symu_b, Select5859)

var ifres9032 Obj

if True == tmp9039 {
tmp9038 := PrimIsPair(Select5860)

var ifres9034 Obj

if True == tmp9038 {
tmp9036 := PrimTail(Select5860)

tmp9037 := PrimEqual(Nil, tmp9036)

var ifres9035 Obj

if True == tmp9037 {
ifres9035 = True


} else {
ifres9035 = False


}

ifres9034 = ifres9035


} else {
ifres9034 = False


}

var ifres9033 Obj

if True == ifres9034 {
ifres9033 = True


} else {
ifres9033 = False


}

ifres9032 = ifres9033


} else {
ifres9032 = False


}

if True == ifres9032 {
tmp8803 := PrimHead(Select5860)

tmp8804 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8803)


tmp8805 := PrimCons(tmp8804, Nil)

__e.Return(PrimCons(symprotect, tmp8805))
return


} else {
tmp9030 := PrimEqual(symerror, Select5859)

var ifres9027 Obj

if True == tmp9030 {
tmp9029 := PrimIsPair(Select5860)

var ifres9028 Obj

if True == tmp9029 {
ifres9028 = True


} else {
ifres9028 = False


}

ifres9027 = ifres9028


} else {
ifres9027 = False


}

if True == ifres9027 {
tmp8806 := PrimHead(Select5860)

tmp8807 := PrimTail(Select5860)

tmp8808 := Call(__e, PrimFunc(symshen_4mkstr), tmp8806, tmp8807)


tmp8809 := PrimCons(tmp8808, Nil)

__e.Return(PrimCons(symsimple_1error, tmp8809))
return


} else {
tmp9025 := PrimEqual(symoutput, Select5859)

var ifres9022 Obj

if True == tmp9025 {
tmp9024 := PrimIsPair(Select5860)

var ifres9023 Obj

if True == tmp9024 {
ifres9023 = True


} else {
ifres9023 = False


}

ifres9022 = ifres9023


} else {
ifres9022 = False


}

if True == ifres9022 {
tmp8810 := PrimHead(Select5860)

tmp8811 := PrimTail(Select5860)

tmp8812 := Call(__e, PrimFunc(symshen_4mkstr), tmp8810, tmp8811)


tmp8813 := PrimCons(symstoutput, Nil)

tmp8814 := PrimCons(tmp8813, Nil)

tmp8815 := PrimCons(tmp8812, tmp8814)

__e.Return(PrimCons(sympr, tmp8815))
return


} else {
tmp9020 := PrimEqual(sympr, Select5859)

var ifres9013 Obj

if True == tmp9020 {
tmp9019 := PrimIsPair(Select5860)

var ifres9015 Obj

if True == tmp9019 {
tmp9017 := PrimTail(Select5860)

tmp9018 := PrimEqual(Nil, tmp9017)

var ifres9016 Obj

if True == tmp9018 {
ifres9016 = True


} else {
ifres9016 = False


}

ifres9015 = ifres9016


} else {
ifres9015 = False


}

var ifres9014 Obj

if True == ifres9015 {
ifres9014 = True


} else {
ifres9014 = False


}

ifres9013 = ifres9014


} else {
ifres9013 = False


}

if True == ifres9013 {
tmp8816 := PrimHead(Select5860)

tmp8817 := PrimCons(symstoutput, Nil)

tmp8818 := PrimCons(tmp8817, Nil)

tmp8819 := PrimCons(tmp8816, tmp8818)

__e.Return(PrimCons(sympr, tmp8819))
return


} else {
tmp9011 := PrimEqual(symmake_1string, Select5859)

var ifres9008 Obj

if True == tmp9011 {
tmp9010 := PrimIsPair(Select5860)

var ifres9009 Obj

if True == tmp9010 {
ifres9009 = True


} else {
ifres9009 = False


}

ifres9008 = ifres9009


} else {
ifres9008 = False


}

if True == ifres9008 {
tmp8820 := PrimHead(Select5860)

tmp8821 := PrimTail(Select5860)

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8820, tmp8821)
return


} else {
tmp9006 := PrimEqual(symlineread, Select5859)

var ifres9003 Obj

if True == tmp9006 {
tmp9005 := PrimEqual(Nil, Select5860)

var ifres9004 Obj

if True == tmp9005 {
ifres9004 = True


} else {
ifres9004 = False


}

ifres9003 = ifres9004


} else {
ifres9003 = False


}

if True == ifres9003 {
tmp8822 := PrimCons(symstinput, Nil)

tmp8823 := PrimCons(tmp8822, Nil)

__e.Return(PrimCons(symlineread, tmp8823))
return


} else {
tmp9001 := PrimEqual(syminput, Select5859)

var ifres8998 Obj

if True == tmp9001 {
tmp9000 := PrimEqual(Nil, Select5860)

var ifres8999 Obj

if True == tmp9000 {
ifres8999 = True


} else {
ifres8999 = False


}

ifres8998 = ifres8999


} else {
ifres8998 = False


}

if True == ifres8998 {
tmp8824 := PrimCons(symstinput, Nil)

tmp8825 := PrimCons(tmp8824, Nil)

__e.Return(PrimCons(syminput, tmp8825))
return


} else {
tmp8996 := PrimEqual(symread, Select5859)

var ifres8993 Obj

if True == tmp8996 {
tmp8995 := PrimEqual(Nil, Select5860)

var ifres8994 Obj

if True == tmp8995 {
ifres8994 = True


} else {
ifres8994 = False


}

ifres8993 = ifres8994


} else {
ifres8993 = False


}

if True == ifres8993 {
tmp8826 := PrimCons(symstinput, Nil)

tmp8827 := PrimCons(tmp8826, Nil)

__e.Return(PrimCons(symread, tmp8827))
return


} else {
tmp8991 := PrimEqual(syminput_7, Select5859)

var ifres8984 Obj

if True == tmp8991 {
tmp8990 := PrimIsPair(Select5860)

var ifres8986 Obj

if True == tmp8990 {
tmp8988 := PrimTail(Select5860)

tmp8989 := PrimEqual(Nil, tmp8988)

var ifres8987 Obj

if True == tmp8989 {
ifres8987 = True


} else {
ifres8987 = False


}

ifres8986 = ifres8987


} else {
ifres8986 = False


}

var ifres8985 Obj

if True == ifres8986 {
ifres8985 = True


} else {
ifres8985 = False


}

ifres8984 = ifres8985


} else {
ifres8984 = False


}

if True == ifres8984 {
tmp8828 := PrimHead(Select5860)

tmp8829 := PrimCons(symstinput, Nil)

tmp8830 := PrimCons(tmp8829, Nil)

tmp8831 := PrimCons(tmp8828, tmp8830)

__e.Return(PrimCons(syminput_7, tmp8831))
return


} else {
tmp8982 := PrimEqual(symread_1byte, Select5859)

var ifres8979 Obj

if True == tmp8982 {
tmp8981 := PrimEqual(Nil, Select5860)

var ifres8980 Obj

if True == tmp8981 {
ifres8980 = True


} else {
ifres8980 = False


}

ifres8979 = ifres8980


} else {
ifres8979 = False


}

if True == ifres8979 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp8977 := PrimEqual(symprolog_2, Select5859)

if True == tmp8977 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), Select5860)
return
} else {
tmp8975 := PrimEqual(symdefprolog, Select5859)

var ifres8972 Obj

if True == tmp8975 {
tmp8974 := PrimIsPair(Select5860)

var ifres8973 Obj

if True == tmp8974 {
ifres8973 = True


} else {
ifres8973 = False


}

ifres8972 = ifres8973


} else {
ifres8972 = False


}

if True == ifres8972 {
tmp8832 := PrimHead(Select5860)

tmp8833 := PrimTail(Select5860)

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8832, tmp8833)
return


} else {
tmp8970 := PrimEqual(symdatatype, Select5859)

var ifres8967 Obj

if True == tmp8970 {
tmp8969 := PrimIsPair(Select5860)

var ifres8968 Obj

if True == tmp8969 {
ifres8968 = True


} else {
ifres8968 = False


}

ifres8967 = ifres8968


} else {
ifres8967 = False


}

if True == ifres8967 {
tmp8834 := PrimHead(Select5860)

tmp8835 := PrimTail(Select5860)

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8834, tmp8835)
return


} else {
tmp8965 := PrimEqual(sym_8s, Select5859)

if True == tmp8965 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V5857)
return
} else {
tmp8963 := PrimEqual(symsynonyms, Select5859)

if True == tmp8963 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select5860)
return
} else {
tmp8961 := PrimEqual(symnl, Select5859)

var ifres8958 Obj

if True == tmp8961 {
tmp8960 := PrimEqual(Nil, Select5860)

var ifres8959 Obj

if True == tmp8960 {
ifres8959 = True


} else {
ifres8959 = False


}

ifres8958 = ifres8959


} else {
ifres8958 = False


}

if True == ifres8958 {
tmp8836 := PrimCons(MakeNumber(1), Nil)

__e.Return(PrimCons(symnl, tmp8836))
return


} else {
tmp8956 := PrimEqual(symlet, Select5859)

if True == tmp8956 {
__e.TailApply(PrimFunc(symshen_4process_1let), V5857)
return
} else {
tmp8954 := PrimEqual(sym_c_4, Select5859)

if True == tmp8954 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V5857)
return
} else {
tmp8952 := PrimEqual(symcases, Select5859)

if True == tmp8952 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V5857)
return
} else {
tmp8950 := PrimEqual(symtime, Select5859)

var ifres8943 Obj

if True == tmp8950 {
tmp8949 := PrimIsPair(Select5860)

var ifres8945 Obj

if True == tmp8949 {
tmp8947 := PrimTail(Select5860)

tmp8948 := PrimEqual(Nil, tmp8947)

var ifres8946 Obj

if True == tmp8948 {
ifres8946 = True


} else {
ifres8946 = False


}

ifres8945 = ifres8946


} else {
ifres8945 = False


}

var ifres8944 Obj

if True == ifres8945 {
ifres8944 = True


} else {
ifres8944 = False


}

ifres8943 = ifres8944


} else {
ifres8943 = False


}

if True == ifres8943 {
tmp8837 := PrimHead(Select5860)

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8837)
return


} else {
tmp8941 := PrimEqual(symput, Select5859)

var ifres8923 Obj

if True == tmp8941 {
tmp8940 := PrimIsPair(Select5860)

var ifres8925 Obj

if True == tmp8940 {
tmp8938 := PrimTail(Select5860)

tmp8939 := PrimIsPair(tmp8938)

var ifres8927 Obj

if True == tmp8939 {
tmp8935 := PrimTail(Select5860)

tmp8936 := PrimTail(tmp8935)

tmp8937 := PrimIsPair(tmp8936)

var ifres8929 Obj

if True == tmp8937 {
tmp8931 := PrimTail(Select5860)

tmp8932 := PrimTail(tmp8931)

tmp8933 := PrimTail(tmp8932)

tmp8934 := PrimEqual(Nil, tmp8933)

var ifres8930 Obj

if True == tmp8934 {
ifres8930 = True


} else {
ifres8930 = False


}

ifres8929 = ifres8930


} else {
ifres8929 = False


}

var ifres8928 Obj

if True == ifres8929 {
ifres8928 = True


} else {
ifres8928 = False


}

ifres8927 = ifres8928


} else {
ifres8927 = False


}

var ifres8926 Obj

if True == ifres8927 {
ifres8926 = True


} else {
ifres8926 = False


}

ifres8925 = ifres8926


} else {
ifres8925 = False


}

var ifres8924 Obj

if True == ifres8925 {
ifres8924 = True


} else {
ifres8924 = False


}

ifres8923 = ifres8924


} else {
ifres8923 = False


}

if True == ifres8923 {
tmp8838 := PrimHead(Select5860)

tmp8839 := PrimTail(Select5860)

tmp8840 := PrimHead(tmp8839)

tmp8841 := PrimTail(Select5860)

tmp8842 := PrimTail(tmp8841)

tmp8843 := PrimHead(tmp8842)

tmp8844 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8845 := PrimCons(symvalue, tmp8844)

tmp8846 := PrimCons(tmp8845, Nil)

tmp8847 := PrimCons(tmp8843, tmp8846)

tmp8848 := PrimCons(tmp8840, tmp8847)

tmp8849 := PrimCons(tmp8838, tmp8848)

__e.Return(PrimCons(symput, tmp8849))
return


} else {
tmp8921 := PrimEqual(symget, Select5859)

var ifres8909 Obj

if True == tmp8921 {
tmp8920 := PrimIsPair(Select5860)

var ifres8911 Obj

if True == tmp8920 {
tmp8918 := PrimTail(Select5860)

tmp8919 := PrimIsPair(tmp8918)

var ifres8913 Obj

if True == tmp8919 {
tmp8915 := PrimTail(Select5860)

tmp8916 := PrimTail(tmp8915)

tmp8917 := PrimEqual(Nil, tmp8916)

var ifres8914 Obj

if True == tmp8917 {
ifres8914 = True


} else {
ifres8914 = False


}

ifres8913 = ifres8914


} else {
ifres8913 = False


}

var ifres8912 Obj

if True == ifres8913 {
ifres8912 = True


} else {
ifres8912 = False


}

ifres8911 = ifres8912


} else {
ifres8911 = False


}

var ifres8910 Obj

if True == ifres8911 {
ifres8910 = True


} else {
ifres8910 = False


}

ifres8909 = ifres8910


} else {
ifres8909 = False


}

if True == ifres8909 {
tmp8850 := PrimHead(Select5860)

tmp8851 := PrimTail(Select5860)

tmp8852 := PrimHead(tmp8851)

tmp8853 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8854 := PrimCons(symvalue, tmp8853)

tmp8855 := PrimCons(tmp8854, Nil)

tmp8856 := PrimCons(tmp8852, tmp8855)

tmp8857 := PrimCons(tmp8850, tmp8856)

__e.Return(PrimCons(symget, tmp8857))
return


} else {
tmp8907 := PrimEqual(symunput, Select5859)

var ifres8895 Obj

if True == tmp8907 {
tmp8906 := PrimIsPair(Select5860)

var ifres8897 Obj

if True == tmp8906 {
tmp8904 := PrimTail(Select5860)

tmp8905 := PrimIsPair(tmp8904)

var ifres8899 Obj

if True == tmp8905 {
tmp8901 := PrimTail(Select5860)

tmp8902 := PrimTail(tmp8901)

tmp8903 := PrimEqual(Nil, tmp8902)

var ifres8900 Obj

if True == tmp8903 {
ifres8900 = True


} else {
ifres8900 = False


}

ifres8899 = ifres8900


} else {
ifres8899 = False


}

var ifres8898 Obj

if True == ifres8899 {
ifres8898 = True


} else {
ifres8898 = False


}

ifres8897 = ifres8898


} else {
ifres8897 = False


}

var ifres8896 Obj

if True == ifres8897 {
ifres8896 = True


} else {
ifres8896 = False


}

ifres8895 = ifres8896


} else {
ifres8895 = False


}

if True == ifres8895 {
tmp8858 := PrimHead(Select5860)

tmp8859 := PrimTail(Select5860)

tmp8860 := PrimHead(tmp8859)

tmp8861 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8862 := PrimCons(symvalue, tmp8861)

tmp8863 := PrimCons(tmp8862, Nil)

tmp8864 := PrimCons(tmp8860, tmp8863)

tmp8865 := PrimCons(tmp8858, tmp8864)

__e.Return(PrimCons(symunput, tmp8865))
return


} else {
tmp8893 := PrimIsPair(Select5860)

var ifres8873 Obj

if True == tmp8893 {
tmp8891 := PrimTail(Select5860)

tmp8892 := PrimIsPair(tmp8891)

var ifres8875 Obj

if True == tmp8892 {
tmp8888 := PrimTail(Select5860)

tmp8889 := PrimTail(tmp8888)

tmp8890 := PrimIsPair(tmp8889)

var ifres8877 Obj

if True == tmp8890 {
tmp8879 := PrimCons(symdo, Nil)

tmp8880 := PrimCons(sym_d, tmp8879)

tmp8881 := PrimCons(sym_7, tmp8880)

tmp8882 := PrimCons(symor, tmp8881)

tmp8883 := PrimCons(symand, tmp8882)

tmp8884 := PrimCons(symappend, tmp8883)

tmp8885 := PrimCons(sym_8v, tmp8884)

tmp8886 := PrimCons(sym_8p, tmp8885)

tmp8887 := Call(__e, PrimFunc(symelement_2), Select5859, tmp8886)


var ifres8878 Obj

if True == tmp8887 {
ifres8878 = True


} else {
ifres8878 = False


}

ifres8877 = ifres8878


} else {
ifres8877 = False


}

var ifres8876 Obj

if True == ifres8877 {
ifres8876 = True


} else {
ifres8876 = False


}

ifres8875 = ifres8876


} else {
ifres8875 = False


}

var ifres8874 Obj

if True == ifres8875 {
ifres8874 = True


} else {
ifres8874 = False


}

ifres8873 = ifres8874


} else {
ifres8873 = False


}

if True == ifres8873 {
tmp8866 := PrimHead(Select5860)

tmp8867 := PrimTail(Select5860)

tmp8868 := PrimCons(Select5859, tmp8867)

tmp8869 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp8868)


tmp8870 := PrimCons(tmp8869, Nil)

tmp8871 := PrimCons(tmp8866, tmp8870)

__e.Return(PrimCons(Select5859, tmp8871))
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5858)
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


}


}


}


}


}


}


}, 1)

tmp9047 := PrimTail(V5857)

__e.TailApply(tmp8800, tmp9047)
return


}, 1)

tmp9048 := PrimHead(V5857)

__e.TailApply(tmp8799, tmp9048)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5858)
return
}


}, 1)

tmp9051 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5857)
return
}, 0)

__e.TailApply(tmp8798, tmp9051)
return


}, 1)

tmp9052 := Call(__e, ns2_1set, symshen_4macros, tmp8797)


_ = tmp9052

tmp9053 := MakeNative(func(__e *ControlFlow) {
V5861 := __e.Get(1)
_ = V5861
V5862 := __e.Get(2)
_ = V5862
tmp9054 := MakeNative(func(__e *ControlFlow) {
W5863 := __e.Get(1)
_ = W5863
tmp9055 := MakeNative(func(__e *ControlFlow) {
W5864 := __e.Get(1)
_ = W5864
tmp9056 := MakeNative(func(__e *ControlFlow) {
W5865 := __e.Get(1)
_ = W5865
__e.Return(V5861)
return
}, 1)

tmp9057 := Call(__e, PrimFunc(symfn), V5861)


tmp9058 := Call(__e, PrimFunc(symshen_4record_1macro), V5861, tmp9057)


__e.TailApply(tmp9056, tmp9058)
return


}, 1)

tmp9059 := Call(__e, PrimFunc(symappend), V5862, W5863)


tmp9060 := PrimCons(V5861, tmp9059)

tmp9061 := PrimCons(symdefine, tmp9060)

tmp9062 := Call(__e, PrimFunc(symeval), tmp9061)


__e.TailApply(tmp9055, tmp9062)
return


}, 1)

tmp9063 := PrimCons(symX, Nil)

tmp9064 := PrimCons(sym_1_6, tmp9063)

tmp9065 := PrimCons(symX, tmp9064)

__e.TailApply(tmp9054, tmp9065)
return


}, 2)

tmp9066 := Call(__e, ns2_1set, symshen_4process_1def, tmp9053)


_ = tmp9066

tmp9067 := MakeNative(func(__e *ControlFlow) {
V5866 := __e.Get(1)
_ = V5866
tmp9107 := PrimIsPair(V5866)

var ifres9081 Obj

if True == tmp9107 {
tmp9105 := PrimHead(V5866)

tmp9106 := PrimEqual(symlet, tmp9105)

var ifres9083 Obj

if True == tmp9106 {
tmp9103 := PrimTail(V5866)

tmp9104 := PrimIsPair(tmp9103)

var ifres9085 Obj

if True == tmp9104 {
tmp9100 := PrimTail(V5866)

tmp9101 := PrimTail(tmp9100)

tmp9102 := PrimIsPair(tmp9101)

var ifres9087 Obj

if True == tmp9102 {
tmp9096 := PrimTail(V5866)

tmp9097 := PrimTail(tmp9096)

tmp9098 := PrimTail(tmp9097)

tmp9099 := PrimIsPair(tmp9098)

var ifres9089 Obj

if True == tmp9099 {
tmp9091 := PrimTail(V5866)

tmp9092 := PrimTail(tmp9091)

tmp9093 := PrimTail(tmp9092)

tmp9094 := PrimTail(tmp9093)

tmp9095 := PrimIsPair(tmp9094)

var ifres9090 Obj

if True == tmp9095 {
ifres9090 = True


} else {
ifres9090 = False


}

ifres9089 = ifres9090


} else {
ifres9089 = False


}

var ifres9088 Obj

if True == ifres9089 {
ifres9088 = True


} else {
ifres9088 = False


}

ifres9087 = ifres9088


} else {
ifres9087 = False


}

var ifres9086 Obj

if True == ifres9087 {
ifres9086 = True


} else {
ifres9086 = False


}

ifres9085 = ifres9086


} else {
ifres9085 = False


}

var ifres9084 Obj

if True == ifres9085 {
ifres9084 = True


} else {
ifres9084 = False


}

ifres9083 = ifres9084


} else {
ifres9083 = False


}

var ifres9082 Obj

if True == ifres9083 {
ifres9082 = True


} else {
ifres9082 = False


}

ifres9081 = ifres9082


} else {
ifres9081 = False


}

if True == ifres9081 {
tmp9068 := PrimTail(V5866)

tmp9069 := PrimHead(tmp9068)

tmp9070 := PrimTail(V5866)

tmp9071 := PrimTail(tmp9070)

tmp9072 := PrimHead(tmp9071)

tmp9073 := PrimTail(V5866)

tmp9074 := PrimTail(tmp9073)

tmp9075 := PrimTail(tmp9074)

tmp9076 := PrimCons(symlet, tmp9075)

tmp9077 := PrimCons(tmp9076, Nil)

tmp9078 := PrimCons(tmp9072, tmp9077)

tmp9079 := PrimCons(tmp9069, tmp9078)

__e.Return(PrimCons(symlet, tmp9079))
return


} else {
__e.Return(V5866)
return
}


}, 1)

tmp9108 := Call(__e, ns2_1set, symshen_4process_1let, tmp9067)


_ = tmp9108

tmp9109 := MakeNative(func(__e *ControlFlow) {
V5867 := __e.Get(1)
_ = V5867
tmp9110 := MakeNative(func(__e *ControlFlow) {
GoTo5869 := __e.Get(1)
_ = GoTo5869
tmp9145 := PrimIsPair(V5867)

if True == tmp9145 {
tmp9111 := MakeNative(func(__e *ControlFlow) {
Select5876 := __e.Get(1)
_ = Select5876
tmp9141 := PrimHead(V5867)

tmp9142 := PrimEqual(sym_8s, tmp9141)

if True == tmp9142 {
tmp9139 := PrimIsPair(Select5876)

if True == tmp9139 {
tmp9112 := MakeNative(func(__e *ControlFlow) {
Select5874 := __e.Get(1)
_ = Select5874
tmp9113 := MakeNative(func(__e *ControlFlow) {
Select5875 := __e.Get(1)
_ = Select5875
tmp9135 := PrimIsPair(Select5875)

if True == tmp9135 {
tmp9114 := MakeNative(func(__e *ControlFlow) {
Select5873 := __e.Get(1)
_ = Select5873
tmp9132 := PrimIsPair(Select5873)

if True == tmp9132 {
tmp9115 := PrimCons(sym_8s, Select5875)

tmp9116 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9115)


tmp9117 := PrimCons(tmp9116, Nil)

tmp9118 := PrimCons(Select5874, tmp9117)

__e.Return(PrimCons(sym_8s, tmp9118))
return


} else {
tmp9130 := PrimEqual(Nil, Select5873)

var ifres9127 Obj

if True == tmp9130 {
tmp9129 := PrimIsString(Select5874)

var ifres9128 Obj

if True == tmp9129 {
ifres9128 = True


} else {
ifres9128 = False


}

ifres9127 = ifres9128


} else {
ifres9127 = False


}

if True == ifres9127 {
tmp9119 := MakeNative(func(__e *ControlFlow) {
W5868 := __e.Get(1)
_ = W5868
tmp9123 := Call(__e, PrimFunc(symlength), W5868)


tmp9124 := PrimGreatThan(tmp9123, MakeNumber(1))

if True == tmp9124 {
tmp9120 := Call(__e, PrimFunc(symappend), W5868, Select5875)


tmp9121 := PrimCons(sym_8s, tmp9120)

__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9121)
return


} else {
__e.Return(V5867)
return
}


}, 1)

tmp9125 := Call(__e, PrimFunc(symexplode), Select5874)


__e.TailApply(tmp9119, tmp9125)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}


}, 1)

tmp9133 := PrimTail(Select5875)

__e.TailApply(tmp9114, tmp9133)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}, 1)

tmp9136 := PrimTail(Select5876)

__e.TailApply(tmp9113, tmp9136)
return


}, 1)

tmp9137 := PrimHead(Select5876)

__e.TailApply(tmp9112, tmp9137)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}, 1)

tmp9143 := PrimTail(V5867)

__e.TailApply(tmp9111, tmp9143)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}, 1)

tmp9146 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5867)
return
}, 0)

__e.TailApply(tmp9110, tmp9146)
return


}, 1)

tmp9147 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9109)


_ = tmp9147

tmp9148 := MakeNative(func(__e *ControlFlow) {
V5877 := __e.Get(1)
_ = V5877
V5878 := __e.Get(2)
_ = V5878
tmp9149 := MakeNative(func(__e *ControlFlow) {
W5879 := __e.Get(1)
_ = W5879
tmp9150 := MakeNative(func(__e *ControlFlow) {
W5880 := __e.Get(1)
_ = W5880
__e.Return(W5879)
return
}, 1)

tmp9151 := MakeNative(func(__e *ControlFlow) {
Z5881 := __e.Get(1)
_ = Z5881
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z5881)
return
}, 1)

tmp9152 := PrimCons(W5879, V5878)

tmp9153 := Call(__e, PrimFunc(symcompile), tmp9151, tmp9152)


__e.TailApply(tmp9150, tmp9153)
return


}, 1)

tmp9154 := Call(__e, PrimFunc(symshen_4intern_1type), V5877)


__e.TailApply(tmp9149, tmp9154)
return


}, 2)

tmp9155 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9148)


_ = tmp9155

tmp9156 := MakeNative(func(__e *ControlFlow) {
V5882 := __e.Get(1)
_ = V5882
tmp9157 := PrimStr(V5882)

tmp9158 := PrimStringConcat(tmp9157, MakeString("#type"))

__e.Return(PrimIntern(tmp9158))
return


}, 1)

tmp9159 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9156)


_ = tmp9159

tmp9160 := MakeNative(func(__e *ControlFlow) {
V5883 := __e.Get(1)
_ = V5883
tmp9161 := PrimValue(symshen_4_dsynonyms_d)

tmp9162 := Call(__e, PrimFunc(symappend), V5883, tmp9161)


tmp9163 := PrimSet(symshen_4_dsynonyms_d, tmp9162)

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9163)
return


}, 1)

tmp9164 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9160)


_ = tmp9164

tmp9165 := MakeNative(func(__e *ControlFlow) {
V5884 := __e.Get(1)
_ = V5884
tmp9166 := MakeNative(func(__e *ControlFlow) {
W5885 := __e.Get(1)
_ = W5885
tmp9167 := MakeNative(func(__e *ControlFlow) {
W5887 := __e.Get(1)
_ = W5887
__e.Return(symsynonyms)
return
}, 1)

tmp9168 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W5885)


tmp9169 := PrimCons(symshen_4demod, tmp9168)

tmp9170 := PrimCons(symdefine, tmp9169)

tmp9171 := Call(__e, PrimFunc(symeval), tmp9170)


__e.TailApply(tmp9167, tmp9171)
return


}, 1)

tmp9172 := MakeNative(func(__e *ControlFlow) {
Z5886 := __e.Get(1)
_ = Z5886
__e.TailApply(PrimFunc(symshen_4curry_1type), Z5886)
return
}, 1)

tmp9173 := Call(__e, PrimFunc(symmap), tmp9172, V5884)


__e.TailApply(tmp9166, tmp9173)
return


}, 1)

tmp9174 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9165)


_ = tmp9174

tmp9175 := MakeNative(func(__e *ControlFlow) {
V5890 := __e.Get(1)
_ = V5890
tmp9197 := PrimEqual(Nil, V5890)

if True == tmp9197 {
tmp9176 := MakeNative(func(__e *ControlFlow) {
W5891 := __e.Get(1)
_ = W5891
tmp9177 := PrimCons(W5891, Nil)

tmp9178 := PrimCons(sym_1_6, tmp9177)

__e.Return(PrimCons(W5891, tmp9178))
return


}, 1)

tmp9179 := Call(__e, PrimFunc(symgensym), symX)


__e.TailApply(tmp9176, tmp9179)
return


} else {
tmp9195 := PrimIsPair(V5890)

var ifres9191 Obj

if True == tmp9195 {
tmp9193 := PrimTail(V5890)

tmp9194 := PrimIsPair(tmp9193)

var ifres9192 Obj

if True == tmp9194 {
ifres9192 = True


} else {
ifres9192 = False


}

ifres9191 = ifres9192


} else {
ifres9191 = False


}

if True == ifres9191 {
tmp9180 := PrimHead(V5890)

tmp9181 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9180)


tmp9182 := PrimTail(V5890)

tmp9183 := PrimHead(tmp9182)

tmp9184 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9183)


tmp9185 := PrimTail(V5890)

tmp9186 := PrimTail(tmp9185)

tmp9187 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9186)


tmp9188 := PrimCons(tmp9184, tmp9187)

tmp9189 := PrimCons(sym_1_6, tmp9188)

__e.Return(PrimCons(tmp9181, tmp9189))
return


} else {
__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
return
}


}


}, 1)

tmp9198 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9175)


_ = tmp9198

tmp9199 := MakeNative(func(__e *ControlFlow) {
V5892 := __e.Get(1)
_ = V5892
tmp9200 := MakeNative(func(__e *ControlFlow) {
GoTo5893 := __e.Get(1)
_ = GoTo5893
tmp9228 := PrimIsPair(V5892)

if True == tmp9228 {
tmp9201 := MakeNative(func(__e *ControlFlow) {
Select5900 := __e.Get(1)
_ = Select5900
tmp9224 := PrimHead(V5892)

tmp9225 := PrimEqual(sym_c_4, tmp9224)

if True == tmp9225 {
tmp9222 := PrimIsPair(Select5900)

if True == tmp9222 {
tmp9202 := MakeNative(func(__e *ControlFlow) {
Select5898 := __e.Get(1)
_ = Select5898
tmp9203 := MakeNative(func(__e *ControlFlow) {
Select5899 := __e.Get(1)
_ = Select5899
tmp9218 := PrimIsPair(Select5899)

if True == tmp9218 {
tmp9204 := MakeNative(func(__e *ControlFlow) {
Select5897 := __e.Get(1)
_ = Select5897
tmp9215 := PrimIsPair(Select5897)

if True == tmp9215 {
tmp9205 := PrimCons(sym_c_4, Select5899)

tmp9206 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9205)


tmp9207 := PrimCons(tmp9206, Nil)

tmp9208 := PrimCons(Select5898, tmp9207)

__e.Return(PrimCons(symlambda, tmp9208))
return


} else {
tmp9213 := PrimEqual(Nil, Select5897)

if True == tmp9213 {
tmp9211 := PrimIsVariable(Select5898)

if True == tmp9211 {
__e.Return(PrimCons(symlambda, Select5900))
return
} else {
tmp9209 := Call(__e, PrimFunc(symshen_4app), Select5898, MakeString(" is not a variable\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp9209))
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}


}, 1)

tmp9216 := PrimTail(Select5899)

__e.TailApply(tmp9204, tmp9216)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9219 := PrimTail(Select5900)

__e.TailApply(tmp9203, tmp9219)
return


}, 1)

tmp9220 := PrimHead(Select5900)

__e.TailApply(tmp9202, tmp9220)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9226 := PrimTail(V5892)

__e.TailApply(tmp9201, tmp9226)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9229 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5892)
return
}, 0)

__e.TailApply(tmp9200, tmp9229)
return


}, 1)

tmp9230 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9199)


_ = tmp9230

tmp9231 := MakeNative(func(__e *ControlFlow) {
V5903 := __e.Get(1)
_ = V5903
tmp9232 := MakeNative(func(__e *ControlFlow) {
GoTo5904 := __e.Get(1)
_ = GoTo5904
tmp9272 := PrimIsPair(V5903)

if True == tmp9272 {
tmp9233 := MakeNative(func(__e *ControlFlow) {
Select5912 := __e.Get(1)
_ = Select5912
tmp9268 := PrimHead(V5903)

tmp9269 := PrimEqual(symcases, tmp9268)

if True == tmp9269 {
tmp9266 := PrimIsPair(Select5912)

if True == tmp9266 {
tmp9234 := MakeNative(func(__e *ControlFlow) {
Select5910 := __e.Get(1)
_ = Select5910
tmp9235 := MakeNative(func(__e *ControlFlow) {
Select5911 := __e.Get(1)
_ = Select5911
tmp9262 := PrimEqual(True, Select5910)

var ifres9259 Obj

if True == tmp9262 {
tmp9261 := PrimIsPair(Select5911)

var ifres9260 Obj

if True == tmp9261 {
ifres9260 = True


} else {
ifres9260 = False


}

ifres9259 = ifres9260


} else {
ifres9259 = False


}

if True == ifres9259 {
__e.Return(PrimHead(Select5911))
return
} else {
tmp9236 := MakeNative(func(__e *ControlFlow) {
GoTo5907 := __e.Get(1)
_ = GoTo5907
tmp9254 := PrimIsPair(Select5911)

if True == tmp9254 {
tmp9237 := MakeNative(func(__e *ControlFlow) {
Select5908 := __e.Get(1)
_ = Select5908
tmp9238 := MakeNative(func(__e *ControlFlow) {
Select5909 := __e.Get(1)
_ = Select5909
tmp9250 := PrimEqual(Nil, Select5909)

if True == tmp9250 {
tmp9239 := PrimCons(MakeString("error: cases exhausted"), Nil)

tmp9240 := PrimCons(symsimple_1error, tmp9239)

tmp9241 := PrimCons(tmp9240, Nil)

tmp9242 := PrimCons(Select5908, tmp9241)

tmp9243 := PrimCons(Select5910, tmp9242)

__e.Return(PrimCons(symif, tmp9243))
return


} else {
tmp9244 := PrimCons(symcases, Select5909)

tmp9245 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9244)


tmp9246 := PrimCons(tmp9245, Nil)

tmp9247 := PrimCons(Select5908, tmp9246)

tmp9248 := PrimCons(Select5910, tmp9247)

__e.Return(PrimCons(symif, tmp9248))
return


}


}, 1)

tmp9251 := PrimTail(Select5911)

__e.TailApply(tmp9238, tmp9251)
return


}, 1)

tmp9252 := PrimHead(Select5911)

__e.TailApply(tmp9237, tmp9252)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5907)
return
}


}, 1)

tmp9255 := MakeNative(func(__e *ControlFlow) {
tmp9257 := PrimEqual(Nil, Select5911)

if True == tmp9257 {
__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


}, 0)

__e.TailApply(tmp9236, tmp9255)
return


}


}, 1)

tmp9263 := PrimTail(Select5912)

__e.TailApply(tmp9235, tmp9263)
return


}, 1)

tmp9264 := PrimHead(Select5912)

__e.TailApply(tmp9234, tmp9264)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


}, 1)

tmp9270 := PrimTail(V5903)

__e.TailApply(tmp9233, tmp9270)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


}, 1)

tmp9273 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5903)
return
}, 0)

__e.TailApply(tmp9232, tmp9273)
return


}, 1)

tmp9274 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9231)


_ = tmp9274

tmp9275 := MakeNative(func(__e *ControlFlow) {
V5913 := __e.Get(1)
_ = V5913
tmp9276 := PrimCons(symrun, Nil)

tmp9277 := PrimCons(symget_1time, tmp9276)

tmp9278 := PrimCons(symrun, Nil)

tmp9279 := PrimCons(symget_1time, tmp9278)

tmp9280 := PrimCons(symStart, Nil)

tmp9281 := PrimCons(symFinish, tmp9280)

tmp9282 := PrimCons(sym_1, tmp9281)

tmp9283 := PrimCons(symTime, Nil)

tmp9284 := PrimCons(symstr, tmp9283)

tmp9285 := PrimCons(MakeString(" secs\n"), Nil)

tmp9286 := PrimCons(tmp9284, tmp9285)

tmp9287 := PrimCons(symcn, tmp9286)

tmp9288 := PrimCons(tmp9287, Nil)

tmp9289 := PrimCons(MakeString("\nrun time: "), tmp9288)

tmp9290 := PrimCons(symcn, tmp9289)

tmp9291 := PrimCons(symstoutput, Nil)

tmp9292 := PrimCons(tmp9291, Nil)

tmp9293 := PrimCons(tmp9290, tmp9292)

tmp9294 := PrimCons(sympr, tmp9293)

tmp9295 := PrimCons(symResult, Nil)

tmp9296 := PrimCons(tmp9294, tmp9295)

tmp9297 := PrimCons(symMessage, tmp9296)

tmp9298 := PrimCons(tmp9282, tmp9297)

tmp9299 := PrimCons(symTime, tmp9298)

tmp9300 := PrimCons(tmp9279, tmp9299)

tmp9301 := PrimCons(symFinish, tmp9300)

tmp9302 := PrimCons(V5913, tmp9301)

tmp9303 := PrimCons(symResult, tmp9302)

tmp9304 := PrimCons(tmp9277, tmp9303)

tmp9305 := PrimCons(symStart, tmp9304)

__e.Return(PrimCons(symlet, tmp9305))
return


}, 1)

tmp9306 := Call(__e, ns2_1set, symshen_4process_1time, tmp9275)


_ = tmp9306

tmp9307 := MakeNative(func(__e *ControlFlow) {
V5914 := __e.Get(1)
_ = V5914
tmp9333 := PrimIsPair(V5914)

var ifres9318 Obj

if True == tmp9333 {
tmp9331 := PrimTail(V5914)

tmp9332 := PrimIsPair(tmp9331)

var ifres9320 Obj

if True == tmp9332 {
tmp9328 := PrimTail(V5914)

tmp9329 := PrimTail(tmp9328)

tmp9330 := PrimIsPair(tmp9329)

var ifres9322 Obj

if True == tmp9330 {
tmp9324 := PrimTail(V5914)

tmp9325 := PrimTail(tmp9324)

tmp9326 := PrimTail(tmp9325)

tmp9327 := PrimIsPair(tmp9326)

var ifres9323 Obj

if True == tmp9327 {
ifres9323 = True


} else {
ifres9323 = False


}

ifres9322 = ifres9323


} else {
ifres9322 = False


}

var ifres9321 Obj

if True == ifres9322 {
ifres9321 = True


} else {
ifres9321 = False


}

ifres9320 = ifres9321


} else {
ifres9320 = False


}

var ifres9319 Obj

if True == ifres9320 {
ifres9319 = True


} else {
ifres9319 = False


}

ifres9318 = ifres9319


} else {
ifres9318 = False


}

if True == ifres9318 {
tmp9308 := PrimHead(V5914)

tmp9309 := PrimTail(V5914)

tmp9310 := PrimHead(tmp9309)

tmp9311 := PrimHead(V5914)

tmp9312 := PrimTail(V5914)

tmp9313 := PrimTail(tmp9312)

tmp9314 := PrimCons(tmp9311, tmp9313)

tmp9315 := PrimCons(tmp9314, Nil)

tmp9316 := PrimCons(tmp9310, tmp9315)

__e.Return(PrimCons(tmp9308, tmp9316))
return


} else {
__e.Return(V5914)
return
}


}, 1)

tmp9334 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9307)


_ = tmp9334

tmp9335 := MakeNative(func(__e *ControlFlow) {
V5915 := __e.Get(1)
_ = V5915
tmp9336 := PrimStr(V5915)

tmp9337 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9336)


__e.Return(PrimIntern(tmp9337))
return


}, 1)

tmp9338 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9335)


_ = tmp9338

tmp9339 := MakeNative(func(__e *ControlFlow) {
V5916 := __e.Get(1)
_ = V5916
tmp9358 := PrimEqual(MakeString(""), V5916)

if True == tmp9358 {
__e.Return(MakeString(""))
return
} else {
tmp9356 := Call(__e, PrimFunc(symshen_4_7string_2), V5916)


if True == tmp9356 {
tmp9340 := MakeNative(func(__e *ControlFlow) {
W5917 := __e.Get(1)
_ = W5917
tmp9341 := MakeNative(func(__e *ControlFlow) {
W5918 := __e.Get(1)
_ = W5918
tmp9342 := MakeNative(func(__e *ControlFlow) {
W5919 := __e.Get(1)
_ = W5919
tmp9343 := PrimTailString(V5916)

tmp9344 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9343)


__e.TailApply(PrimFunc(sym_8s), W5919, tmp9344)
return


}, 1)

tmp9351 := PrimGreatEqual(W5917, MakeNumber(97))

var ifres9348 Obj

if True == tmp9351 {
tmp9350 := PrimLessEqual(W5917, MakeNumber(122))

var ifres9349 Obj

if True == tmp9350 {
ifres9349 = True


} else {
ifres9349 = False


}

ifres9348 = ifres9349


} else {
ifres9348 = False


}

var ifres9345 Obj

if True == ifres9348 {
tmp9346 := PrimNumberToString(W5918)

ifres9345 = tmp9346


} else {
tmp9347 := Call(__e, PrimFunc(symhdstr), V5916)


ifres9345 = tmp9347


}

__e.TailApply(tmp9342, ifres9345)
return


}, 1)

tmp9352 := PrimNumberSubtract(W5917, MakeNumber(32))

__e.TailApply(tmp9341, tmp9352)
return


}, 1)

tmp9353 := Call(__e, PrimFunc(symhdstr), V5916)


tmp9354 := PrimStringToNumber(tmp9353)

__e.TailApply(tmp9340, tmp9354)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.mu-h")))
return
}


}


}, 1)

tmp9359 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9339)


_ = tmp9359

tmp9360 := MakeNative(func(__e *ControlFlow) {
V5920 := __e.Get(1)
_ = V5920
V5921 := __e.Get(2)
_ = V5921
tmp9361 := PrimValue(sym_dmacros_d)

tmp9362 := Call(__e, PrimFunc(symshen_4update_1assoc), V5920, V5921, tmp9361)


__e.Return(PrimSet(sym_dmacros_d, tmp9362))
return


}, 2)

tmp9363 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9360)


_ = tmp9363

tmp9364 := MakeNative(func(__e *ControlFlow) {
V5931 := __e.Get(1)
_ = V5931
V5932 := __e.Get(2)
_ = V5932
V5933 := __e.Get(3)
_ = V5933
tmp9384 := PrimEqual(Nil, V5933)

if True == tmp9384 {
tmp9365 := PrimCons(V5931, V5932)

__e.Return(PrimCons(tmp9365, Nil))
return


} else {
tmp9366 := MakeNative(func(__e *ControlFlow) {
GoTo5934 := __e.Get(1)
_ = GoTo5934
tmp9381 := PrimIsPair(V5933)

if True == tmp9381 {
tmp9367 := MakeNative(func(__e *ControlFlow) {
Select5935 := __e.Get(1)
_ = Select5935
tmp9368 := MakeNative(func(__e *ControlFlow) {
Select5936 := __e.Get(1)
_ = Select5936
tmp9377 := PrimIsPair(Select5935)

var ifres9373 Obj

if True == tmp9377 {
tmp9375 := PrimHead(Select5935)

tmp9376 := PrimEqual(V5931, tmp9375)

var ifres9374 Obj

if True == tmp9376 {
ifres9374 = True


} else {
ifres9374 = False


}

ifres9373 = ifres9374


} else {
ifres9373 = False


}

if True == ifres9373 {
tmp9369 := PrimHead(Select5935)

tmp9370 := PrimCons(tmp9369, V5932)

__e.Return(PrimCons(tmp9370, Select5936))
return


} else {
tmp9371 := Call(__e, PrimFunc(symshen_4update_1assoc), V5931, V5932, Select5936)


__e.Return(PrimCons(Select5935, tmp9371))
return


}


}, 1)

tmp9378 := PrimTail(V5933)

__e.TailApply(tmp9368, tmp9378)
return


}, 1)

tmp9379 := PrimHead(V5933)

__e.TailApply(tmp9367, tmp9379)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5934)
return
}


}, 1)

tmp9382 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
return
}, 0)

__e.TailApply(tmp9366, tmp9382)
return


}


}, 3)

tmp9385 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9364)


_ = tmp9385

tmp9386 := MakeNative(func(__e *ControlFlow) {
tmp9394 := Call(__e, PrimFunc(symstinput))


tmp9395 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp9394)


if True == tmp9395 {
tmp9387 := PrimCons(symstinput, Nil)

tmp9388 := PrimCons(tmp9387, Nil)

tmp9389 := PrimCons(symshen_4read_1unit_1string, tmp9388)

tmp9390 := PrimCons(tmp9389, Nil)

__e.Return(PrimCons(symstring_1_6n, tmp9390))
return


} else {
tmp9391 := PrimCons(symstinput, Nil)

tmp9392 := PrimCons(tmp9391, Nil)

__e.Return(PrimCons(symread_1byte, tmp9392))
return


}


}, 0)

tmp9396 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9386)


_ = tmp9396

tmp9397 := MakeNative(func(__e *ControlFlow) {
V5937 := __e.Get(1)
_ = V5937
tmp9398 := MakeNative(func(__e *ControlFlow) {
W5938 := __e.Get(1)
_ = W5938
tmp9399 := MakeNative(func(__e *ControlFlow) {
W5939 := __e.Get(1)
_ = W5939
tmp9400 := MakeNative(func(__e *ControlFlow) {
W5940 := __e.Get(1)
_ = W5940
tmp9401 := MakeNative(func(__e *ControlFlow) {
W5941 := __e.Get(1)
_ = W5941
tmp9402 := MakeNative(func(__e *ControlFlow) {
W5942 := __e.Get(1)
_ = W5942
tmp9403 := MakeNative(func(__e *ControlFlow) {
W5944 := __e.Get(1)
_ = W5944
tmp9404 := MakeNative(func(__e *ControlFlow) {
W5945 := __e.Get(1)
_ = W5945
tmp9405 := MakeNative(func(__e *ControlFlow) {
W5946 := __e.Get(1)
_ = W5946
tmp9406 := MakeNative(func(__e *ControlFlow) {
W5947 := __e.Get(1)
_ = W5947
tmp9407 := MakeNative(func(__e *ControlFlow) {
W5948 := __e.Get(1)
_ = W5948
tmp9408 := MakeNative(func(__e *ControlFlow) {
W5949 := __e.Get(1)
_ = W5949
tmp9409 := PrimCons(W5941, Nil)

tmp9410 := PrimCons(W5940, tmp9409)

tmp9411 := PrimCons(W5939, tmp9410)

tmp9412 := PrimCons(W5938, tmp9411)

__e.Return(PrimCons(W5949, tmp9412))
return


}, 1)

tmp9413 := Call(__e, PrimFunc(symshen_4continue), W5944, W5942, W5945, W5946, W5947, W5948)


tmp9414 := PrimCons(tmp9413, Nil)

tmp9415 := PrimCons(W5948, tmp9414)

tmp9416 := PrimCons(symlambda, tmp9415)

tmp9417 := PrimCons(tmp9416, Nil)

tmp9418 := PrimCons(W5947, tmp9417)

tmp9419 := PrimCons(symlambda, tmp9418)

tmp9420 := PrimCons(tmp9419, Nil)

tmp9421 := PrimCons(W5946, tmp9420)

tmp9422 := PrimCons(symlambda, tmp9421)

tmp9423 := PrimCons(tmp9422, Nil)

tmp9424 := PrimCons(W5945, tmp9423)

tmp9425 := PrimCons(symlambda, tmp9424)

__e.TailApply(tmp9408, tmp9425)
return


}, 1)

tmp9426 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp9407, tmp9426)
return


}, 1)

tmp9427 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp9406, tmp9427)
return


}, 1)

tmp9428 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp9405, tmp9428)
return


}, 1)

tmp9429 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp9404, tmp9429)
return


}, 1)

tmp9430 := Call(__e, PrimFunc(symshen_4received), V5937)


__e.TailApply(tmp9403, tmp9430)
return


}, 1)

tmp9431 := MakeNative(func(__e *ControlFlow) {
Z5943 := __e.Get(1)
_ = Z5943
__e.TailApply(PrimFunc(symshen_4_5body_6), Z5943)
return
}, 1)

tmp9432 := Call(__e, PrimFunc(symcompile), tmp9431, V5937)


__e.TailApply(tmp9402, tmp9432)
return


}, 1)

tmp9433 := PrimCons(True, Nil)

tmp9434 := PrimCons(symfreeze, tmp9433)

__e.TailApply(tmp9401, tmp9434)
return


}, 1)

__e.TailApply(tmp9400, MakeNumber(0))
return


}, 1)

tmp9435 := PrimCons(MakeNumber(0), Nil)

tmp9436 := PrimCons(symvector, tmp9435)

tmp9437 := PrimCons(tmp9436, Nil)

tmp9438 := PrimCons(MakeNumber(0), tmp9437)

tmp9439 := PrimCons(True, tmp9438)

tmp9440 := PrimCons(sym_8v, tmp9439)

__e.TailApply(tmp9399, tmp9440)
return


}, 1)

tmp9441 := PrimCons(symshen_4prolog_1vector, Nil)

__e.TailApply(tmp9398, tmp9441)
return


}, 1)

tmp9442 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9397)


_ = tmp9442

tmp9443 := MakeNative(func(__e *ControlFlow) {
V5952 := __e.Get(1)
_ = V5952
tmp9444 := MakeNative(func(__e *ControlFlow) {
GoTo5953 := __e.Get(1)
_ = GoTo5953
tmp9461 := PrimIsPair(V5952)

if True == tmp9461 {
tmp9445 := MakeNative(func(__e *ControlFlow) {
Select5954 := __e.Get(1)
_ = Select5954
tmp9446 := MakeNative(func(__e *ControlFlow) {
Select5955 := __e.Get(1)
_ = Select5955
tmp9457 := PrimEqual(symreceive, Select5954)

var ifres9450 Obj

if True == tmp9457 {
tmp9456 := PrimIsPair(Select5955)

var ifres9452 Obj

if True == tmp9456 {
tmp9454 := PrimTail(Select5955)

tmp9455 := PrimEqual(Nil, tmp9454)

var ifres9453 Obj

if True == tmp9455 {
ifres9453 = True


} else {
ifres9453 = False


}

ifres9452 = ifres9453


} else {
ifres9452 = False


}

var ifres9451 Obj

if True == ifres9452 {
ifres9451 = True


} else {
ifres9451 = False


}

ifres9450 = ifres9451


} else {
ifres9450 = False


}

if True == ifres9450 {
__e.Return(Select5955)
return
} else {
tmp9447 := Call(__e, PrimFunc(symshen_4received), Select5954)


tmp9448 := Call(__e, PrimFunc(symshen_4received), Select5955)


__e.TailApply(PrimFunc(symunion), tmp9447, tmp9448)
return


}


}, 1)

tmp9458 := PrimTail(V5952)

__e.TailApply(tmp9446, tmp9458)
return


}, 1)

tmp9459 := PrimHead(V5952)

__e.TailApply(tmp9445, tmp9459)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5953)
return
}


}, 1)

tmp9462 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

__e.TailApply(tmp9444, tmp9462)
return


}, 1)

tmp9463 := Call(__e, ns2_1set, symshen_4received, tmp9443)


_ = tmp9463

tmp9464 := MakeNative(func(__e *ControlFlow) {
tmp9465 := MakeNative(func(__e *ControlFlow) {
W5956 := __e.Get(1)
_ = W5956
tmp9466 := MakeNative(func(__e *ControlFlow) {
W5957 := __e.Get(1)
_ = W5957
tmp9467 := MakeNative(func(__e *ControlFlow) {
W5958 := __e.Get(1)
_ = W5958
__e.Return(W5958)
return
}, 1)

tmp9468 := PrimVectorSet(W5956, MakeNumber(1), MakeNumber(2))

__e.TailApply(tmp9467, tmp9468)
return


}, 1)

tmp9469 := PrimVectorSet(W5956, MakeNumber(0), symshen_4print_1prolog_1vector)

__e.TailApply(tmp9466, tmp9469)
return


}, 1)

tmp9470 := PrimValue(symshen_4_dprolog_1memory_d)

tmp9471 := PrimAbsvector(tmp9470)

__e.TailApply(tmp9465, tmp9471)
return


}, 0)

tmp9472 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9464)


_ = tmp9472

tmp9473 := MakeNative(func(__e *ControlFlow) {
V5959 := __e.Get(1)
_ = V5959
__e.Return(V5959)
return
}, 1)

tmp9474 := Call(__e, ns2_1set, symreceive, tmp9473)


_ = tmp9474

tmp9475 := MakeNative(func(__e *ControlFlow) {
V5960 := __e.Get(1)
_ = V5960
tmp9483 := PrimIsPair(V5960)

if True == tmp9483 {
tmp9476 := PrimHead(V5960)

tmp9477 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9476)


tmp9478 := PrimTail(V5960)

tmp9479 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9478)


tmp9480 := PrimCons(tmp9479, Nil)

tmp9481 := PrimCons(tmp9477, tmp9480)

__e.Return(PrimCons(symcons, tmp9481))
return


} else {
__e.Return(V5960)
return
}


}, 1)

tmp9484 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9475)


_ = tmp9484

tmp9485 := MakeNative(func(__e *ControlFlow) {
V5961 := __e.Get(1)
_ = V5961
tmp9492 := PrimIsPair(V5961)

if True == tmp9492 {
tmp9486 := PrimHead(V5961)

tmp9487 := PrimTail(V5961)

tmp9488 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9487)


tmp9489 := PrimCons(tmp9488, Nil)

tmp9490 := PrimCons(tmp9486, tmp9489)

__e.Return(PrimCons(sym_8p, tmp9490))
return


} else {
__e.Return(V5961)
return
}


}, 1)

tmp9493 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9485)


_ = tmp9493

tmp9494 := MakeNative(func(__e *ControlFlow) {
V5962 := __e.Get(1)
_ = V5962
tmp9495 := PrimValue(sym_dmacros_d)

tmp9496 := Call(__e, PrimFunc(symassoc), V5962, tmp9495)


tmp9497 := PrimValue(sym_dmacros_d)

tmp9498 := Call(__e, PrimFunc(symremove), tmp9496, tmp9497)


tmp9499 := PrimSet(sym_dmacros_d, tmp9498)

_ = tmp9499

__e.Return(V5962)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9494)
return




}, 0)

