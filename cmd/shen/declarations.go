package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp6954 := PrimSet(symshen_4_dhistory_d, Nil)

_ = tmp6954

tmp6955 := PrimSet(symshen_4_dtc_d, False)

_ = tmp6955

tmp6956 := Call(__e, PrimFunc(symvector), MakeNumber(20000))


tmp6957 := PrimSet(sym_dproperty_1vector_d, tmp6956)

_ = tmp6957

tmp6958 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4defmacro_1macro), X)
return
}, 1)

tmp6959 := PrimCons(symshen_4defmacro_1macro, tmp6958)

tmp6960 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4timer_1macro), X)
return
}, 1)

tmp6961 := PrimCons(symshen_4timer_1macro, tmp6960)

tmp6962 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4cases_1macro), X)
return
}, 1)

tmp6963 := PrimCons(symshen_4cases_1macro, tmp6962)

tmp6964 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4abs_1macro), X)
return
}, 1)

tmp6965 := PrimCons(symshen_4abs_1macro, tmp6964)

tmp6966 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4put_cget_1macro), X)
return
}, 1)

tmp6967 := PrimCons(symshen_4put_cget_1macro, tmp6966)

tmp6968 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4datatype_1macro), X)
return
}, 1)

tmp6969 := PrimCons(symshen_4datatype_1macro, tmp6968)

tmp6970 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4let_1macro), X)
return
}, 1)

tmp6971 := PrimCons(symshen_4let_1macro, tmp6970)

tmp6972 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4assoc_1macro), X)
return
}, 1)

tmp6973 := PrimCons(symshen_4assoc_1macro, tmp6972)

tmp6974 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4make_1string_1macro), X)
return
}, 1)

tmp6975 := PrimCons(symmake_1string, tmp6974)

tmp6976 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4output_1macro), X)
return
}, 1)

tmp6977 := PrimCons(symshen_4output_1macro, tmp6976)

tmp6978 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4input_1macro), X)
return
}, 1)

tmp6979 := PrimCons(symshen_4input_1macro, tmp6978)

tmp6980 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4error_1macro), X)
return
}, 1)

tmp6981 := PrimCons(symshen_4error_1macro, tmp6980)

tmp6982 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4prolog_1macro), X)
return
}, 1)

tmp6983 := PrimCons(symshen_4prolog_1macro, tmp6982)

tmp6984 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4synonyms_1macro), X)
return
}, 1)

tmp6985 := PrimCons(symshen_4synonyms_1macro, tmp6984)

tmp6986 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4nl_1macro), X)
return
}, 1)

tmp6987 := PrimCons(symshen_4nl_1macro, tmp6986)

tmp6988 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_8s_1macro), X)
return
}, 1)

tmp6989 := PrimCons(symshen_4_8s_1macro, tmp6988)

tmp6990 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4defprolog_1macro), X)
return
}, 1)

tmp6991 := PrimCons(symdefprolog, tmp6990)

tmp6992 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4defcc_1macro), X)
return
}, 1)

tmp6993 := PrimCons(symshen_4defcc_1macro, tmp6992)

tmp6994 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4u_b_1macro), X)
return
}, 1)

tmp6995 := PrimCons(symshen_4u_b_1macro, tmp6994)

tmp6996 := PrimCons(tmp6995, Nil)

tmp6997 := PrimCons(tmp6993, tmp6996)

tmp6998 := PrimCons(tmp6991, tmp6997)

tmp6999 := PrimCons(tmp6989, tmp6998)

tmp7000 := PrimCons(tmp6987, tmp6999)

tmp7001 := PrimCons(tmp6985, tmp7000)

tmp7002 := PrimCons(tmp6983, tmp7001)

tmp7003 := PrimCons(tmp6981, tmp7002)

tmp7004 := PrimCons(tmp6979, tmp7003)

tmp7005 := PrimCons(tmp6977, tmp7004)

tmp7006 := PrimCons(tmp6975, tmp7005)

tmp7007 := PrimCons(tmp6973, tmp7006)

tmp7008 := PrimCons(tmp6971, tmp7007)

tmp7009 := PrimCons(tmp6969, tmp7008)

tmp7010 := PrimCons(tmp6967, tmp7009)

tmp7011 := PrimCons(tmp6965, tmp7010)

tmp7012 := PrimCons(tmp6963, tmp7011)

tmp7013 := PrimCons(tmp6961, tmp7012)

tmp7014 := PrimCons(tmp6959, tmp7013)

tmp7015 := PrimSet(sym_dmacros_d, tmp7014)

_ = tmp7015

tmp7016 := PrimSet(symshen_4_dgensym_d, MakeNumber(0))

_ = tmp7016

tmp7017 := PrimSet(symshen_4_dtracking_d, Nil)

_ = tmp7017

tmp7018 := PrimSet(symshen_4_dprofiled_d, Nil)

_ = tmp7018

tmp7019 := PrimSet(sym_dhome_1directory_d, MakeString(""))

_ = tmp7019

tmp7020 := PrimCons(symtype, Nil)

tmp7021 := PrimCons(syminput_7, tmp7020)

tmp7022 := PrimCons(symopen, tmp7021)

tmp7023 := PrimCons(symset, tmp7022)

tmp7024 := PrimCons(symwhere, tmp7023)

tmp7025 := PrimCons(symlet, tmp7024)

tmp7026 := PrimCons(symlambda, tmp7025)

tmp7027 := PrimCons(symcons, tmp7026)

tmp7028 := PrimCons(sym_8v, tmp7027)

tmp7029 := PrimCons(sym_8s, tmp7028)

tmp7030 := PrimCons(sym_8p, tmp7029)

tmp7031 := PrimSet(symshen_4_dspecial_d, tmp7030)

_ = tmp7031

tmp7032 := PrimSet(symshen_4_dextraspecial_d, Nil)

_ = tmp7032

tmp7033 := PrimSet(symshen_4_dspy_d, False)

_ = tmp7033

tmp7034 := PrimSet(symshen_4_ddatatypes_d, Nil)

_ = tmp7034

tmp7035 := PrimSet(symshen_4_dalldatatypes_d, Nil)

_ = tmp7035

tmp7036 := PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp7036

tmp7037 := PrimSet(symshen_4_dpackage_d, symnull)

_ = tmp7037

tmp7038 := PrimSet(symshen_4_dsynonyms_d, Nil)

_ = tmp7038

tmp7039 := PrimSet(symshen_4_dsystem_d, Nil)

_ = tmp7039

tmp7040 := PrimSet(symshen_4_dsigf_d, Nil)

_ = tmp7040

tmp7041 := PrimSet(symshen_4_doccurs_d, True)

_ = tmp7041

tmp7042 := PrimSet(symshen_4_dfactorise_2_d, False)

_ = tmp7042

tmp7043 := PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1000000))

_ = tmp7043

tmp7044 := PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

_ = tmp7044

tmp7045 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp7045

tmp7046 := PrimSet(symshen_4_dinfs_d, MakeNumber(0))

_ = tmp7046

tmp7047 := PrimSet(sym_dhush_d, False)

_ = tmp7047

tmp7048 := PrimSet(symshen_4_doptimise_d, False)

_ = tmp7048

tmp7049 := PrimSet(sym_dversion_d, MakeString("31"))

_ = tmp7049

tmp7050 := PrimSet(symshen_4_dstep_d, False)

_ = tmp7050

tmp7051 := PrimSet(symshen_4_dit_d, MakeString(""))

_ = tmp7051

tmp7052 := PrimSet(symshen_4_dresidue_d, Nil)

_ = tmp7052

tmp7053 := MakeNative(func(__e *ControlFlow) {
V1674 := __e.Get(1)
_ = V1674
tmp7054 := MakeNative(func(__e *ControlFlow) {
Bindings := __e.Get(1)
_ = Bindings
tmp7055 := MakeNative(func(__e *ControlFlow) {
PrintNamed := __e.Get(1)
_ = PrintNamed
tmp7056 := MakeNative(func(__e *ControlFlow) {
Ticketed := __e.Get(1)
_ = Ticketed
tmp7057 := MakeNative(func(__e *ControlFlow) {
Assign := __e.Get(1)
_ = Assign
__e.Return(V1674)
return
}, 1)

tmp7058 := PrimSet(symshen_4_dprolog_1vector_d, Ticketed)

__e.TailApply(tmp7057, tmp7058)
return


}, 1)

tmp7059 := PrimVectorSet(Bindings, MakeNumber(1), MakeNumber(2))

__e.TailApply(tmp7056, tmp7059)
return


}, 1)

tmp7060 := PrimVectorSet(Bindings, MakeNumber(0), symshen_4print_1prolog_1vector)

__e.TailApply(tmp7055, tmp7060)
return


}, 1)

tmp7061 := PrimAbsvector(V1674)

__e.TailApply(tmp7054, tmp7061)
return


}, 1)

tmp7062 := Call(__e, ns2_1set, symprolog_1memory, tmp7053)


_ = tmp7062

tmp7063 := Call(__e, PrimFunc(symprolog_1memory), MakeNumber(10000))


_ = tmp7063

tmp7064 := PrimSet(symshen_4_dloading_2_d, False)

_ = tmp7064

tmp7065 := MakeNative(func(__e *ControlFlow) {
V1677 := __e.Get(1)
_ = V1677
tmp7081 := PrimEqual(Nil, V1677)

if True == tmp7081 {
__e.Return(Nil)
return
} else {
tmp7079 := PrimIsPair(V1677)

var ifres7075 Obj

if True == tmp7079 {
tmp7077 := PrimTail(V1677)

tmp7078 := PrimIsPair(tmp7077)

var ifres7076 Obj

if True == tmp7078 {
ifres7076 = True


} else {
ifres7076 = False


}

ifres7075 = ifres7076


} else {
ifres7075 = False


}

if True == ifres7075 {
tmp7066 := MakeNative(func(__e *ControlFlow) {
DecArity := __e.Get(1)
_ = DecArity
tmp7067 := PrimTail(V1677)

tmp7068 := PrimTail(tmp7067)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7068)
return


}, 1)

tmp7069 := PrimHead(V1677)

tmp7070 := PrimTail(V1677)

tmp7071 := PrimHead(tmp7070)

tmp7072 := PrimValue(sym_dproperty_1vector_d)

tmp7073 := Call(__e, PrimFunc(symput), tmp7069, symarity, tmp7071, tmp7072)


__e.TailApply(tmp7066, tmp7073)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7082 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1tables, tmp7065)


_ = tmp7082

tmp7083 := MakeNative(func(__e *ControlFlow) {
V1678 := __e.Get(1)
_ = V1678
tmp7084 := MakeNative(func(__e *ControlFlow) {
tmp7085 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V1678, symarity, tmp7085)
return


}, 0)

tmp7086 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(MakeNumber(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp7084, tmp7086)
return


}, 1)

tmp7087 := Call(__e, ns2_1set, symarity, tmp7083)


_ = tmp7087

tmp7088 := MakeNative(func(__e *ControlFlow) {
V1681 := __e.Get(1)
_ = V1681
tmp7104 := PrimEqual(Nil, V1681)

if True == tmp7104 {
__e.Return(Nil)
return
} else {
tmp7102 := PrimIsPair(V1681)

var ifres7098 Obj

if True == tmp7102 {
tmp7100 := PrimTail(V1681)

tmp7101 := PrimIsPair(tmp7100)

var ifres7099 Obj

if True == tmp7101 {
ifres7099 = True


} else {
ifres7099 = False


}

ifres7098 = ifres7099


} else {
ifres7098 = False


}

if True == ifres7098 {
tmp7089 := MakeNative(func(__e *ControlFlow) {
DecArity := __e.Get(1)
_ = DecArity
tmp7090 := PrimTail(V1681)

tmp7091 := PrimTail(tmp7090)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7091)
return


}, 1)

tmp7092 := PrimHead(V1681)

tmp7093 := PrimTail(V1681)

tmp7094 := PrimHead(tmp7093)

tmp7095 := PrimValue(sym_dproperty_1vector_d)

tmp7096 := Call(__e, PrimFunc(symput), tmp7092, symarity, tmp7094, tmp7095)


__e.TailApply(tmp7089, tmp7096)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise_arity_table")))
return
}


}


}, 1)

tmp7105 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7088)


_ = tmp7105

tmp7106 := PrimCons(MakeNumber(2), Nil)

tmp7107 := PrimCons(sym_8s, tmp7106)

tmp7108 := PrimCons(MakeNumber(2), tmp7107)

tmp7109 := PrimCons(sym_8v, tmp7108)

tmp7110 := PrimCons(MakeNumber(2), tmp7109)

tmp7111 := PrimCons(sym_8p, tmp7110)

tmp7112 := PrimCons(MakeNumber(1), tmp7111)

tmp7113 := PrimCons(sym_5_b_6, tmp7112)

tmp7114 := PrimCons(MakeNumber(1), tmp7113)

tmp7115 := PrimCons(symshen_4_5end_6, tmp7114)

tmp7116 := PrimCons(MakeNumber(1), tmp7115)

tmp7117 := PrimCons(sym_5e_6, tmp7116)

tmp7118 := PrimCons(MakeNumber(2), tmp7117)

tmp7119 := PrimCons(sym_a_a, tmp7118)

tmp7120 := PrimCons(MakeNumber(2), tmp7119)

tmp7121 := PrimCons(sym_1, tmp7120)

tmp7122 := PrimCons(MakeNumber(2), tmp7121)

tmp7123 := PrimCons(sym_c, tmp7122)

tmp7124 := PrimCons(MakeNumber(2), tmp7123)

tmp7125 := PrimCons(sym_d, tmp7124)

tmp7126 := PrimCons(MakeNumber(2), tmp7125)

tmp7127 := PrimCons(sym_7, tmp7126)

tmp7128 := PrimCons(MakeNumber(1), tmp7127)

tmp7129 := PrimCons(symy_1or_1n_2, tmp7128)

tmp7130 := PrimCons(MakeNumber(2), tmp7129)

tmp7131 := PrimCons(symwrite_1to_1file, tmp7130)

tmp7132 := PrimCons(MakeNumber(2), tmp7131)

tmp7133 := PrimCons(symwrite_1byte, tmp7132)

tmp7134 := PrimCons(MakeNumber(5), tmp7133)

tmp7135 := PrimCons(symwhen, tmp7134)

tmp7136 := PrimCons(MakeNumber(0), tmp7135)

tmp7137 := PrimCons(symversion, tmp7136)

tmp7138 := PrimCons(MakeNumber(5), tmp7137)

tmp7139 := PrimCons(symvar_2, tmp7138)

tmp7140 := PrimCons(MakeNumber(1), tmp7139)

tmp7141 := PrimCons(symvariable_2, tmp7140)

tmp7142 := PrimCons(MakeNumber(1), tmp7141)

tmp7143 := PrimCons(symvalue, tmp7142)

tmp7144 := PrimCons(MakeNumber(3), tmp7143)

tmp7145 := PrimCons(symvector_1_6, tmp7144)

tmp7146 := PrimCons(MakeNumber(1), tmp7145)

tmp7147 := PrimCons(symvector_2, tmp7146)

tmp7148 := PrimCons(MakeNumber(1), tmp7147)

tmp7149 := PrimCons(symvector, tmp7148)

tmp7150 := PrimCons(MakeNumber(2), tmp7149)

tmp7151 := PrimCons(symupdate_1lambda_1table, tmp7150)

tmp7152 := PrimCons(MakeNumber(1), tmp7151)

tmp7153 := PrimCons(symundefmacro, tmp7152)

tmp7154 := PrimCons(MakeNumber(1), tmp7153)

tmp7155 := PrimCons(symuntrack, tmp7154)

tmp7156 := PrimCons(MakeNumber(2), tmp7155)

tmp7157 := PrimCons(symunion, tmp7156)

tmp7158 := PrimCons(MakeNumber(1), tmp7157)

tmp7159 := PrimCons(symunprofile, tmp7158)

tmp7160 := PrimCons(MakeNumber(3), tmp7159)

tmp7161 := PrimCons(symunput, tmp7160)

tmp7162 := PrimCons(MakeNumber(1), tmp7161)

tmp7163 := PrimCons(symundefmacro, tmp7162)

tmp7164 := PrimCons(MakeNumber(5), tmp7163)

tmp7165 := PrimCons(symreturn, tmp7164)

tmp7166 := PrimCons(MakeNumber(2), tmp7165)

tmp7167 := PrimCons(symtype, tmp7166)

tmp7168 := PrimCons(MakeNumber(1), tmp7167)

tmp7169 := PrimCons(symtuple_2, tmp7168)

tmp7170 := PrimCons(MakeNumber(2), tmp7169)

tmp7171 := PrimCons(symtrap_1error, tmp7170)

tmp7172 := PrimCons(MakeNumber(1), tmp7171)

tmp7173 := PrimCons(symtrack, tmp7172)

tmp7174 := PrimCons(MakeNumber(1), tmp7173)

tmp7175 := PrimCons(symtlstr, tmp7174)

tmp7176 := PrimCons(MakeNumber(1), tmp7175)

tmp7177 := PrimCons(symthaw, tmp7176)

tmp7178 := PrimCons(MakeNumber(0), tmp7177)

tmp7179 := PrimCons(symtc_2, tmp7178)

tmp7180 := PrimCons(MakeNumber(1), tmp7179)

tmp7181 := PrimCons(symtc, tmp7180)

tmp7182 := PrimCons(MakeNumber(1), tmp7181)

tmp7183 := PrimCons(symtl, tmp7182)

tmp7184 := PrimCons(MakeNumber(1), tmp7183)

tmp7185 := PrimCons(symtail, tmp7184)

tmp7186 := PrimCons(MakeNumber(1), tmp7185)

tmp7187 := PrimCons(symsystemf, tmp7186)

tmp7188 := PrimCons(MakeNumber(1), tmp7187)

tmp7189 := PrimCons(symsymbol_2, tmp7188)

tmp7190 := PrimCons(MakeNumber(1), tmp7189)

tmp7191 := PrimCons(symsum, tmp7190)

tmp7192 := PrimCons(MakeNumber(3), tmp7191)

tmp7193 := PrimCons(symsubst, tmp7192)

tmp7194 := PrimCons(MakeNumber(1), tmp7193)

tmp7195 := PrimCons(symstring_2, tmp7194)

tmp7196 := PrimCons(MakeNumber(1), tmp7195)

tmp7197 := PrimCons(symstring_1_6symbol, tmp7196)

tmp7198 := PrimCons(MakeNumber(1), tmp7197)

tmp7199 := PrimCons(symstring_1_6n, tmp7198)

tmp7200 := PrimCons(MakeNumber(1), tmp7199)

tmp7201 := PrimCons(symstr, tmp7200)

tmp7202 := PrimCons(MakeNumber(0), tmp7201)

tmp7203 := PrimCons(symstoutput, tmp7202)

tmp7204 := PrimCons(MakeNumber(0), tmp7203)

tmp7205 := PrimCons(symstinput, tmp7204)

tmp7206 := PrimCons(MakeNumber(1), tmp7205)

tmp7207 := PrimCons(symstep, tmp7206)

tmp7208 := PrimCons(MakeNumber(1), tmp7207)

tmp7209 := PrimCons(symspy, tmp7208)

tmp7210 := PrimCons(MakeNumber(2), tmp7209)

tmp7211 := PrimCons(symspecialise, tmp7210)

tmp7212 := PrimCons(MakeNumber(1), tmp7211)

tmp7213 := PrimCons(symsnd, tmp7212)

tmp7214 := PrimCons(MakeNumber(1), tmp7213)

tmp7215 := PrimCons(symsimple_1error, tmp7214)

tmp7216 := PrimCons(MakeNumber(2), tmp7215)

tmp7217 := PrimCons(symset, tmp7216)

tmp7218 := PrimCons(MakeNumber(1), tmp7217)

tmp7219 := PrimCons(symreverse, tmp7218)

tmp7220 := PrimCons(MakeNumber(2), tmp7219)

tmp7221 := PrimCons(symremove, tmp7220)

tmp7222 := PrimCons(MakeNumber(0), tmp7221)

tmp7223 := PrimCons(symrelease, tmp7222)

tmp7224 := PrimCons(MakeNumber(1), tmp7223)

tmp7225 := PrimCons(symreceive, tmp7224)

tmp7226 := PrimCons(MakeNumber(1), tmp7225)

tmp7227 := PrimCons(symshen_4read_1unit_1string, tmp7226)

tmp7228 := PrimCons(MakeNumber(1), tmp7227)

tmp7229 := PrimCons(symread_1from_1string_1unprocessed, tmp7228)

tmp7230 := PrimCons(MakeNumber(1), tmp7229)

tmp7231 := PrimCons(symread_1from_1string, tmp7230)

tmp7232 := PrimCons(MakeNumber(1), tmp7231)

tmp7233 := PrimCons(symread_1byte, tmp7232)

tmp7234 := PrimCons(MakeNumber(1), tmp7233)

tmp7235 := PrimCons(symread, tmp7234)

tmp7236 := PrimCons(MakeNumber(1), tmp7235)

tmp7237 := PrimCons(symread_1file, tmp7236)

tmp7238 := PrimCons(MakeNumber(1), tmp7237)

tmp7239 := PrimCons(symread_1file_1as_1bytelist, tmp7238)

tmp7240 := PrimCons(MakeNumber(1), tmp7239)

tmp7241 := PrimCons(symread_1file_1as_1string, tmp7240)

tmp7242 := PrimCons(MakeNumber(4), tmp7241)

tmp7243 := PrimCons(symput, tmp7242)

tmp7244 := PrimCons(MakeNumber(1), tmp7243)

tmp7245 := PrimCons(symprotect, tmp7244)

tmp7246 := PrimCons(MakeNumber(1), tmp7245)

tmp7247 := PrimCons(sympreclude_1all_1but, tmp7246)

tmp7248 := PrimCons(MakeNumber(1), tmp7247)

tmp7249 := PrimCons(sympreclude, tmp7248)

tmp7250 := PrimCons(MakeNumber(1), tmp7249)

tmp7251 := PrimCons(symps, tmp7250)

tmp7252 := PrimCons(MakeNumber(2), tmp7251)

tmp7253 := PrimCons(sympr, tmp7252)

tmp7254 := PrimCons(MakeNumber(1), tmp7253)

tmp7255 := PrimCons(symprofile_1results, tmp7254)

tmp7256 := PrimCons(MakeNumber(1), tmp7255)

tmp7257 := PrimCons(symprolog_1memory, tmp7256)

tmp7258 := PrimCons(MakeNumber(1), tmp7257)

tmp7259 := PrimCons(symshen_4printF, tmp7258)

tmp7260 := PrimCons(MakeNumber(1), tmp7259)

tmp7261 := PrimCons(symshen_4print_1freshterm, tmp7260)

tmp7262 := PrimCons(MakeNumber(1), tmp7261)

tmp7263 := PrimCons(symshen_4print_1prolog_1vector, tmp7262)

tmp7264 := PrimCons(MakeNumber(1), tmp7263)

tmp7265 := PrimCons(symprofile, tmp7264)

tmp7266 := PrimCons(MakeNumber(1), tmp7265)

tmp7267 := PrimCons(symshen_4pprint, tmp7266)

tmp7268 := PrimCons(MakeNumber(1), tmp7267)

tmp7269 := PrimCons(symprint, tmp7268)

tmp7270 := PrimCons(MakeNumber(1), tmp7269)

tmp7271 := PrimCons(sympreclude_1all_1but, tmp7270)

tmp7272 := PrimCons(MakeNumber(2), tmp7271)

tmp7273 := PrimCons(sympos, tmp7272)

tmp7274 := PrimCons(MakeNumber(0), tmp7273)

tmp7275 := PrimCons(symporters, tmp7274)

tmp7276 := PrimCons(MakeNumber(0), tmp7275)

tmp7277 := PrimCons(symport, tmp7276)

tmp7278 := PrimCons(MakeNumber(1), tmp7277)

tmp7279 := PrimCons(sympackage_2, tmp7278)

tmp7280 := PrimCons(MakeNumber(3), tmp7279)

tmp7281 := PrimCons(sympackage, tmp7280)

tmp7282 := PrimCons(MakeNumber(0), tmp7281)

tmp7283 := PrimCons(symos, tmp7282)

tmp7284 := PrimCons(MakeNumber(2), tmp7283)

tmp7285 := PrimCons(symor, tmp7284)

tmp7286 := PrimCons(MakeNumber(1), tmp7285)

tmp7287 := PrimCons(symoptimise, tmp7286)

tmp7288 := PrimCons(MakeNumber(2), tmp7287)

tmp7289 := PrimCons(symopen, tmp7288)

tmp7290 := PrimCons(MakeNumber(1), tmp7289)

tmp7291 := PrimCons(symoccurs_1check, tmp7290)

tmp7292 := PrimCons(MakeNumber(2), tmp7291)

tmp7293 := PrimCons(symoccurrences, tmp7292)

tmp7294 := PrimCons(MakeNumber(1), tmp7293)

tmp7295 := PrimCons(symoccurs_1check, tmp7294)

tmp7296 := PrimCons(MakeNumber(1), tmp7295)

tmp7297 := PrimCons(symnumber_2, tmp7296)

tmp7298 := PrimCons(MakeNumber(1), tmp7297)

tmp7299 := PrimCons(symn_1_6string, tmp7298)

tmp7300 := PrimCons(MakeNumber(2), tmp7299)

tmp7301 := PrimCons(symnth, tmp7300)

tmp7302 := PrimCons(MakeNumber(1), tmp7301)

tmp7303 := PrimCons(symnot, tmp7302)

tmp7304 := PrimCons(MakeNumber(1), tmp7303)

tmp7305 := PrimCons(symnl, tmp7304)

tmp7306 := PrimCons(MakeNumber(1), tmp7305)

tmp7307 := PrimCons(symmaxinferences, tmp7306)

tmp7308 := PrimCons(MakeNumber(2), tmp7307)

tmp7309 := PrimCons(symmapcan, tmp7308)

tmp7310 := PrimCons(MakeNumber(2), tmp7309)

tmp7311 := PrimCons(symmap, tmp7310)

tmp7312 := PrimCons(MakeNumber(1), tmp7311)

tmp7313 := PrimCons(symmacroexpand, tmp7312)

tmp7314 := PrimCons(MakeNumber(1), tmp7313)

tmp7315 := PrimCons(symvector, tmp7314)

tmp7316 := PrimCons(MakeNumber(2), tmp7315)

tmp7317 := PrimCons(sym_5_a, tmp7316)

tmp7318 := PrimCons(MakeNumber(2), tmp7317)

tmp7319 := PrimCons(sym_5, tmp7318)

tmp7320 := PrimCons(MakeNumber(1), tmp7319)

tmp7321 := PrimCons(symload, tmp7320)

tmp7322 := PrimCons(MakeNumber(1), tmp7321)

tmp7323 := PrimCons(symlist, tmp7322)

tmp7324 := PrimCons(MakeNumber(1), tmp7323)

tmp7325 := PrimCons(symlineread, tmp7324)

tmp7326 := PrimCons(MakeNumber(1), tmp7325)

tmp7327 := PrimCons(symlimit, tmp7326)

tmp7328 := PrimCons(MakeNumber(1), tmp7327)

tmp7329 := PrimCons(symlength, tmp7328)

tmp7330 := PrimCons(MakeNumber(0), tmp7329)

tmp7331 := PrimCons(symlanguage, tmp7330)

tmp7332 := PrimCons(MakeNumber(6), tmp7331)

tmp7333 := PrimCons(symis_b, tmp7332)

tmp7334 := PrimCons(MakeNumber(6), tmp7333)

tmp7335 := PrimCons(symis, tmp7334)

tmp7336 := PrimCons(MakeNumber(0), tmp7335)

tmp7337 := PrimCons(symit, tmp7336)

tmp7338 := PrimCons(MakeNumber(1), tmp7337)

tmp7339 := PrimCons(syminternal, tmp7338)

tmp7340 := PrimCons(MakeNumber(2), tmp7339)

tmp7341 := PrimCons(symintersection, tmp7340)

tmp7342 := PrimCons(MakeNumber(1), tmp7341)

tmp7343 := PrimCons(syminclude_1all_1but, tmp7342)

tmp7344 := PrimCons(MakeNumber(0), tmp7343)

tmp7345 := PrimCons(symimplementation, tmp7344)

tmp7346 := PrimCons(MakeNumber(2), tmp7345)

tmp7347 := PrimCons(syminput_7, tmp7346)

tmp7348 := PrimCons(MakeNumber(1), tmp7347)

tmp7349 := PrimCons(syminput, tmp7348)

tmp7350 := PrimCons(MakeNumber(0), tmp7349)

tmp7351 := PrimCons(syminferences, tmp7350)

tmp7352 := PrimCons(MakeNumber(1), tmp7351)

tmp7353 := PrimCons(symintern, tmp7352)

tmp7354 := PrimCons(MakeNumber(1), tmp7353)

tmp7355 := PrimCons(syminternal, tmp7354)

tmp7356 := PrimCons(MakeNumber(1), tmp7355)

tmp7357 := PrimCons(syminteger_2, tmp7356)

tmp7358 := PrimCons(MakeNumber(1), tmp7357)

tmp7359 := PrimCons(symin_1package, tmp7358)

tmp7360 := PrimCons(MakeNumber(1), tmp7359)

tmp7361 := PrimCons(syminclude, tmp7360)

tmp7362 := PrimCons(MakeNumber(3), tmp7361)

tmp7363 := PrimCons(symif, tmp7362)

tmp7364 := PrimCons(MakeNumber(1), tmp7363)

tmp7365 := PrimCons(symhead, tmp7364)

tmp7366 := PrimCons(MakeNumber(1), tmp7365)

tmp7367 := PrimCons(symhdstr, tmp7366)

tmp7368 := PrimCons(MakeNumber(1), tmp7367)

tmp7369 := PrimCons(symhdv, tmp7368)

tmp7370 := PrimCons(MakeNumber(1), tmp7369)

tmp7371 := PrimCons(symhd, tmp7370)

tmp7372 := PrimCons(MakeNumber(2), tmp7371)

tmp7373 := PrimCons(symhash, tmp7372)

tmp7374 := PrimCons(MakeNumber(2), tmp7373)

tmp7375 := PrimCons(sym_a, tmp7374)

tmp7376 := PrimCons(MakeNumber(2), tmp7375)

tmp7377 := PrimCons(sym_6_a, tmp7376)

tmp7378 := PrimCons(MakeNumber(2), tmp7377)

tmp7379 := PrimCons(sym_6, tmp7378)

tmp7380 := PrimCons(MakeNumber(2), tmp7379)

tmp7381 := PrimCons(sym_5_1vector, tmp7380)

tmp7382 := PrimCons(MakeNumber(2), tmp7381)

tmp7383 := PrimCons(sym_5_1address, tmp7382)

tmp7384 := PrimCons(MakeNumber(3), tmp7383)

tmp7385 := PrimCons(symaddress_1_6, tmp7384)

tmp7386 := PrimCons(MakeNumber(1), tmp7385)

tmp7387 := PrimCons(symget_1time, tmp7386)

tmp7388 := PrimCons(MakeNumber(3), tmp7387)

tmp7389 := PrimCons(symget, tmp7388)

tmp7390 := PrimCons(MakeNumber(1), tmp7389)

tmp7391 := PrimCons(symgensym, tmp7390)

tmp7392 := PrimCons(MakeNumber(1), tmp7391)

tmp7393 := PrimCons(symfunction, tmp7392)

tmp7394 := PrimCons(MakeNumber(1), tmp7393)

tmp7395 := PrimCons(symfn, tmp7394)

tmp7396 := PrimCons(MakeNumber(1), tmp7395)

tmp7397 := PrimCons(symfst, tmp7396)

tmp7398 := PrimCons(MakeNumber(0), tmp7397)

tmp7399 := PrimCons(symfresh, tmp7398)

tmp7400 := PrimCons(MakeNumber(1), tmp7399)

tmp7401 := PrimCons(symfreeze, tmp7400)

tmp7402 := PrimCons(MakeNumber(5), tmp7401)

tmp7403 := PrimCons(symfork, tmp7402)

tmp7404 := PrimCons(MakeNumber(7), tmp7403)

tmp7405 := PrimCons(symfindall, tmp7404)

tmp7406 := PrimCons(MakeNumber(2), tmp7405)

tmp7407 := PrimCons(symfix, tmp7406)

tmp7408 := PrimCons(MakeNumber(0), tmp7407)

tmp7409 := PrimCons(symfail, tmp7408)

tmp7410 := PrimCons(MakeNumber(2), tmp7409)

tmp7411 := PrimCons(symfail_1if, tmp7410)

tmp7412 := PrimCons(MakeNumber(1), tmp7411)

tmp7413 := PrimCons(symfactorise, tmp7412)

tmp7414 := PrimCons(MakeNumber(1), tmp7413)

tmp7415 := PrimCons(symexternal, tmp7414)

tmp7416 := PrimCons(MakeNumber(1), tmp7415)

tmp7417 := PrimCons(symexplode, tmp7416)

tmp7418 := PrimCons(MakeNumber(1), tmp7417)

tmp7419 := PrimCons(symeval_1kl, tmp7418)

tmp7420 := PrimCons(MakeNumber(1), tmp7419)

tmp7421 := PrimCons(symeval, tmp7420)

tmp7422 := PrimCons(MakeNumber(2), tmp7421)

tmp7423 := PrimCons(symshen_4interror, tmp7422)

tmp7424 := PrimCons(MakeNumber(1), tmp7423)

tmp7425 := PrimCons(symerror_1to_1string, tmp7424)

tmp7426 := PrimCons(MakeNumber(1), tmp7425)

tmp7427 := PrimCons(symexternal, tmp7426)

tmp7428 := PrimCons(MakeNumber(1), tmp7427)

tmp7429 := PrimCons(symenable_1type_1theory, tmp7428)

tmp7430 := PrimCons(MakeNumber(1), tmp7429)

tmp7431 := PrimCons(symempty_2, tmp7430)

tmp7432 := PrimCons(MakeNumber(2), tmp7431)

tmp7433 := PrimCons(symelement_2, tmp7432)

tmp7434 := PrimCons(MakeNumber(2), tmp7433)

tmp7435 := PrimCons(symdo, tmp7434)

tmp7436 := PrimCons(MakeNumber(2), tmp7435)

tmp7437 := PrimCons(symdifference, tmp7436)

tmp7438 := PrimCons(MakeNumber(1), tmp7437)

tmp7439 := PrimCons(symdestroy, tmp7438)

tmp7440 := PrimCons(MakeNumber(2), tmp7439)

tmp7441 := PrimCons(symdeclare, tmp7440)

tmp7442 := PrimCons(MakeNumber(1), tmp7441)

tmp7443 := PrimCons(symclose, tmp7442)

tmp7444 := PrimCons(MakeNumber(2), tmp7443)

tmp7445 := PrimCons(symcn, tmp7444)

tmp7446 := PrimCons(MakeNumber(1), tmp7445)

tmp7447 := PrimCons(symcons_2, tmp7446)

tmp7448 := PrimCons(MakeNumber(2), tmp7447)

tmp7449 := PrimCons(symcons, tmp7448)

tmp7450 := PrimCons(MakeNumber(2), tmp7449)

tmp7451 := PrimCons(symconcat, tmp7450)

tmp7452 := PrimCons(MakeNumber(2), tmp7451)

tmp7453 := PrimCons(symcompile, tmp7452)

tmp7454 := PrimCons(MakeNumber(1), tmp7453)

tmp7455 := PrimCons(symcd, tmp7454)

tmp7456 := PrimCons(MakeNumber(5), tmp7455)

tmp7457 := PrimCons(symcall, tmp7456)

tmp7458 := PrimCons(MakeNumber(6), tmp7457)

tmp7459 := PrimCons(symbind, tmp7458)

tmp7460 := PrimCons(MakeNumber(1), tmp7459)

tmp7461 := PrimCons(symbound_2, tmp7460)

tmp7462 := PrimCons(MakeNumber(1), tmp7461)

tmp7463 := PrimCons(symbootstrap, tmp7462)

tmp7464 := PrimCons(MakeNumber(1), tmp7463)

tmp7465 := PrimCons(symboolean_2, tmp7464)

tmp7466 := PrimCons(MakeNumber(1), tmp7465)

tmp7467 := PrimCons(symatom_2, tmp7466)

tmp7468 := PrimCons(MakeNumber(2), tmp7467)

tmp7469 := PrimCons(symassoc, tmp7468)

tmp7470 := PrimCons(MakeNumber(1), tmp7469)

tmp7471 := PrimCons(symarity, tmp7470)

tmp7472 := PrimCons(MakeNumber(2), tmp7471)

tmp7473 := PrimCons(symappend, tmp7472)

tmp7474 := PrimCons(MakeNumber(2), tmp7473)

tmp7475 := PrimCons(symand, tmp7474)

tmp7476 := PrimCons(MakeNumber(2), tmp7475)

tmp7477 := PrimCons(symadjoin, tmp7476)

tmp7478 := PrimCons(MakeNumber(3), tmp7477)

tmp7479 := PrimCons(symaddress_1_6, tmp7478)

tmp7480 := PrimCons(MakeNumber(1), tmp7479)

tmp7481 := PrimCons(symabsvector, tmp7480)

tmp7482 := PrimCons(MakeNumber(1), tmp7481)

tmp7483 := PrimCons(symabsvector_2, tmp7482)

tmp7484 := PrimCons(MakeNumber(0), tmp7483)

tmp7485 := PrimCons(symabort, tmp7484)

tmp7486 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp7485)


_ = tmp7486

tmp7487 := MakeNative(func(__e *ControlFlow) {
V1682 := __e.Get(1)
_ = V1682
tmp7488 := MakeNative(func(__e *ControlFlow) {
External := __e.Get(1)
_ = External
tmp7489 := MakeNative(func(__e *ControlFlow) {
Place := __e.Get(1)
_ = Place
__e.Return(V1682)
return
}, 1)

tmp7490 := Call(__e, PrimFunc(symadjoin), V1682, External)


tmp7491 := PrimValue(sym_dproperty_1vector_d)

tmp7492 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp7490, tmp7491)


__e.TailApply(tmp7489, tmp7492)
return


}, 1)

tmp7493 := PrimValue(sym_dproperty_1vector_d)

tmp7494 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp7493)


__e.TailApply(tmp7488, tmp7494)
return


}, 1)

tmp7495 := Call(__e, ns2_1set, symsystemf, tmp7487)


_ = tmp7495

tmp7496 := MakeNative(func(__e *ControlFlow) {
V1683 := __e.Get(1)
_ = V1683
V1684 := __e.Get(2)
_ = V1684
tmp7498 := Call(__e, PrimFunc(symelement_2), V1683, V1684)


if True == tmp7498 {
__e.Return(V1684)
return
} else {
__e.Return(PrimCons(V1683, V1684))
return
}


}, 2)

tmp7499 := Call(__e, ns2_1set, symadjoin, tmp7496)


_ = tmp7499

tmp7500 := PrimIntern(MakeString(":"))

tmp7501 := PrimIntern(MakeString(";"))

tmp7502 := PrimIntern(MakeString(":="))

tmp7503 := PrimIntern(MakeString(","))

tmp7504 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp7505 := PrimIntern(MakeString("bar!"))

tmp7506 := PrimCons(symabort, Nil)

tmp7507 := PrimCons(symabsvector, tmp7506)

tmp7508 := PrimCons(symabsvector_2, tmp7507)

tmp7509 := PrimCons(symaddress_1_6, tmp7508)

tmp7510 := PrimCons(sym_5_1address, tmp7509)

tmp7511 := PrimCons(symadjoin, tmp7510)

tmp7512 := PrimCons(symand, tmp7511)

tmp7513 := PrimCons(symappend, tmp7512)

tmp7514 := PrimCons(symarity, tmp7513)

tmp7515 := PrimCons(symassoc, tmp7514)

tmp7516 := PrimCons(symatom_2, tmp7515)

tmp7517 := PrimCons(tmp7505, tmp7516)

tmp7518 := PrimCons(symbootstrap, tmp7517)

tmp7519 := PrimCons(symboolean, tmp7518)

tmp7520 := PrimCons(symboolean_2, tmp7519)

tmp7521 := PrimCons(symbound_2, tmp7520)

tmp7522 := PrimCons(symbind, tmp7521)

tmp7523 := PrimCons(symclose, tmp7522)

tmp7524 := PrimCons(symcall, tmp7523)

tmp7525 := PrimCons(symcases, tmp7524)

tmp7526 := PrimCons(symcd, tmp7525)

tmp7527 := PrimCons(symcompile, tmp7526)

tmp7528 := PrimCons(symconcat, tmp7527)

tmp7529 := PrimCons(symcond, tmp7528)

tmp7530 := PrimCons(symcons, tmp7529)

tmp7531 := PrimCons(symcons_2, tmp7530)

tmp7532 := PrimCons(symcn, tmp7531)

tmp7533 := PrimCons(symdatatype, tmp7532)

tmp7534 := PrimCons(symdeclare, tmp7533)

tmp7535 := PrimCons(symdefprolog, tmp7534)

tmp7536 := PrimCons(symdefcc, tmp7535)

tmp7537 := PrimCons(symdefmacro, tmp7536)

tmp7538 := PrimCons(symdefine, tmp7537)

tmp7539 := PrimCons(symdefun, tmp7538)

tmp7540 := PrimCons(symdestroy, tmp7539)

tmp7541 := PrimCons(symdifference, tmp7540)

tmp7542 := PrimCons(symdo, tmp7541)

tmp7543 := PrimCons(symelement_2, tmp7542)

tmp7544 := PrimCons(symempty_2, tmp7543)

tmp7545 := PrimCons(symerror, tmp7544)

tmp7546 := PrimCons(symerror_1to_1string, tmp7545)

tmp7547 := PrimCons(symeval, tmp7546)

tmp7548 := PrimCons(symeval_1kl, tmp7547)

tmp7549 := PrimCons(symexception, tmp7548)

tmp7550 := PrimCons(symexternal, tmp7549)

tmp7551 := PrimCons(symexplode, tmp7550)

tmp7552 := PrimCons(symenable_1type_1theory, tmp7551)

tmp7553 := PrimCons(False, tmp7552)

tmp7554 := PrimCons(symfindall, tmp7553)

tmp7555 := PrimCons(symfactorise, tmp7554)

tmp7556 := PrimCons(symfail_1if, tmp7555)

tmp7557 := PrimCons(symfail, tmp7556)

tmp7558 := PrimCons(symfile, tmp7557)

tmp7559 := PrimCons(symfix, tmp7558)

tmp7560 := PrimCons(symfork, tmp7559)

tmp7561 := PrimCons(symfresh, tmp7560)

tmp7562 := PrimCons(symfreeze, tmp7561)

tmp7563 := PrimCons(symfst, tmp7562)

tmp7564 := PrimCons(symfunction, tmp7563)

tmp7565 := PrimCons(symfn, tmp7564)

tmp7566 := PrimCons(symgensym, tmp7565)

tmp7567 := PrimCons(symget_1time, tmp7566)

tmp7568 := PrimCons(symget, tmp7567)

tmp7569 := PrimCons(symhash, tmp7568)

tmp7570 := PrimCons(symhdstr, tmp7569)

tmp7571 := PrimCons(symhdv, tmp7570)

tmp7572 := PrimCons(symhd, tmp7571)

tmp7573 := PrimCons(symhead, tmp7572)

tmp7574 := PrimCons(symif, tmp7573)

tmp7575 := PrimCons(symimplementation, tmp7574)

tmp7576 := PrimCons(syminternal, tmp7575)

tmp7577 := PrimCons(symin_1package, tmp7576)

tmp7578 := PrimCons(symin, tmp7577)

tmp7579 := PrimCons(symis_b, tmp7578)

tmp7580 := PrimCons(symis, tmp7579)

tmp7581 := PrimCons(symit, tmp7580)

tmp7582 := PrimCons(syminclude_1all_1but, tmp7581)

tmp7583 := PrimCons(syminclude, tmp7582)

tmp7584 := PrimCons(syminline, tmp7583)

tmp7585 := PrimCons(syminput_7, tmp7584)

tmp7586 := PrimCons(syminput, tmp7585)

tmp7587 := PrimCons(syminteger_2, tmp7586)

tmp7588 := PrimCons(symintern, tmp7587)

tmp7589 := PrimCons(syminferences, tmp7588)

tmp7590 := PrimCons(symintersection, tmp7589)

tmp7591 := PrimCons(symis, tmp7590)

tmp7592 := PrimCons(symlanguage, tmp7591)

tmp7593 := PrimCons(symlambda, tmp7592)

tmp7594 := PrimCons(symlazy, tmp7593)

tmp7595 := PrimCons(symlet, tmp7594)

tmp7596 := PrimCons(symlength, tmp7595)

tmp7597 := PrimCons(symlimit, tmp7596)

tmp7598 := PrimCons(symlineread, tmp7597)

tmp7599 := PrimCons(symlist, tmp7598)

tmp7600 := PrimCons(symloaded, tmp7599)

tmp7601 := PrimCons(symload, tmp7600)

tmp7602 := PrimCons(symmake_1string, tmp7601)

tmp7603 := PrimCons(symmap, tmp7602)

tmp7604 := PrimCons(symmapcan, tmp7603)

tmp7605 := PrimCons(symmaxinferences, tmp7604)

tmp7606 := PrimCons(symmacroexpand, tmp7605)

tmp7607 := PrimCons(symmode, tmp7606)

tmp7608 := PrimCons(symnl, tmp7607)

tmp7609 := PrimCons(symnot, tmp7608)

tmp7610 := PrimCons(symnth, tmp7609)

tmp7611 := PrimCons(symnull, tmp7610)

tmp7612 := PrimCons(symnumber, tmp7611)

tmp7613 := PrimCons(symnumber_2, tmp7612)

tmp7614 := PrimCons(symn_1_6string, tmp7613)

tmp7615 := PrimCons(symoccurs_1check, tmp7614)

tmp7616 := PrimCons(symoccurrences, tmp7615)

tmp7617 := PrimCons(symopen, tmp7616)

tmp7618 := PrimCons(symoptimise, tmp7617)

tmp7619 := PrimCons(symor, tmp7618)

tmp7620 := PrimCons(symos, tmp7619)

tmp7621 := PrimCons(symout, tmp7620)

tmp7622 := PrimCons(symoutput, tmp7621)

tmp7623 := PrimCons(sympackage, tmp7622)

tmp7624 := PrimCons(symport, tmp7623)

tmp7625 := PrimCons(symporters, tmp7624)

tmp7626 := PrimCons(sympos, tmp7625)

tmp7627 := PrimCons(sympr, tmp7626)

tmp7628 := PrimCons(symshen_4pprint, tmp7627)

tmp7629 := PrimCons(symprint, tmp7628)

tmp7630 := PrimCons(symprolog_1memory, tmp7629)

tmp7631 := PrimCons(symprofile, tmp7630)

tmp7632 := PrimCons(symprofile_1results, tmp7631)

tmp7633 := PrimCons(symprotect, tmp7632)

tmp7634 := PrimCons(symprolog_2, tmp7633)

tmp7635 := PrimCons(symps, tmp7634)

tmp7636 := PrimCons(sympreclude_1all_1but, tmp7635)

tmp7637 := PrimCons(sympreclude, tmp7636)

tmp7638 := PrimCons(symput, tmp7637)

tmp7639 := PrimCons(sympackage_2, tmp7638)

tmp7640 := PrimCons(symread_1from_1string_1unprocessed, tmp7639)

tmp7641 := PrimCons(symread_1from_1string, tmp7640)

tmp7642 := PrimCons(symread_1byte, tmp7641)

tmp7643 := PrimCons(symread_1file_1as_1string, tmp7642)

tmp7644 := PrimCons(symread_1file_1as_1bytelist, tmp7643)

tmp7645 := PrimCons(symread_1file, tmp7644)

tmp7646 := PrimCons(symreceive, tmp7645)

tmp7647 := PrimCons(symread, tmp7646)

tmp7648 := PrimCons(symrelease, tmp7647)

tmp7649 := PrimCons(symremove, tmp7648)

tmp7650 := PrimCons(symreverse, tmp7649)

tmp7651 := PrimCons(symrun, tmp7650)

tmp7652 := PrimCons(symstr, tmp7651)

tmp7653 := PrimCons(symsave, tmp7652)

tmp7654 := PrimCons(symset, tmp7653)

tmp7655 := PrimCons(symsimple_1error, tmp7654)

tmp7656 := PrimCons(symsnd, tmp7655)

tmp7657 := PrimCons(symspecialise, tmp7656)

tmp7658 := PrimCons(symspy, tmp7657)

tmp7659 := PrimCons(symstep, tmp7658)

tmp7660 := PrimCons(symstoutput, tmp7659)

tmp7661 := PrimCons(symstinput, tmp7660)

tmp7662 := PrimCons(symstring, tmp7661)

tmp7663 := PrimCons(symstream, tmp7662)

tmp7664 := PrimCons(symstring_1_6n, tmp7663)

tmp7665 := PrimCons(symstring_2, tmp7664)

tmp7666 := PrimCons(symsubst, tmp7665)

tmp7667 := PrimCons(symsum, tmp7666)

tmp7668 := PrimCons(symstring_1_6symbol, tmp7667)

tmp7669 := PrimCons(symsymbol_2, tmp7668)

tmp7670 := PrimCons(symsymbol, tmp7669)

tmp7671 := PrimCons(symsynonyms, tmp7670)

tmp7672 := PrimCons(symsystemf, tmp7671)

tmp7673 := PrimCons(symtail, tmp7672)

tmp7674 := PrimCons(symtlv, tmp7673)

tmp7675 := PrimCons(symtlstr, tmp7674)

tmp7676 := PrimCons(symtl, tmp7675)

tmp7677 := PrimCons(symtc, tmp7676)

tmp7678 := PrimCons(symtc_2, tmp7677)

tmp7679 := PrimCons(symthaw, tmp7678)

tmp7680 := PrimCons(symtime, tmp7679)

tmp7681 := PrimCons(symtrack, tmp7680)

tmp7682 := PrimCons(symtrap_1error, tmp7681)

tmp7683 := PrimCons(True, tmp7682)

tmp7684 := PrimCons(symtuple_2, tmp7683)

tmp7685 := PrimCons(symtype, tmp7684)

tmp7686 := PrimCons(symreturn, tmp7685)

tmp7687 := PrimCons(symundefmacro, tmp7686)

tmp7688 := PrimCons(symunprofile, tmp7687)

tmp7689 := PrimCons(symunput, tmp7688)

tmp7690 := PrimCons(symunion, tmp7689)

tmp7691 := PrimCons(symshen_4unix, tmp7690)

tmp7692 := PrimCons(symunit, tmp7691)

tmp7693 := PrimCons(symuntrack, tmp7692)

tmp7694 := PrimCons(symunspecialise, tmp7693)

tmp7695 := PrimCons(symupdate_1lambda_1table, tmp7694)

tmp7696 := PrimCons(symu_b, tmp7695)

tmp7697 := PrimCons(symvector_2, tmp7696)

tmp7698 := PrimCons(symvector, tmp7697)

tmp7699 := PrimCons(sym_5_1vector, tmp7698)

tmp7700 := PrimCons(symvector_1_6, tmp7699)

tmp7701 := PrimCons(symvalue, tmp7700)

tmp7702 := PrimCons(symvar_2, tmp7701)

tmp7703 := PrimCons(symvariable_2, tmp7702)

tmp7704 := PrimCons(symverified, tmp7703)

tmp7705 := PrimCons(symversion, tmp7704)

tmp7706 := PrimCons(symwarn, tmp7705)

tmp7707 := PrimCons(symwhen, tmp7706)

tmp7708 := PrimCons(symwhere, tmp7707)

tmp7709 := PrimCons(symwrite_1byte, tmp7708)

tmp7710 := PrimCons(symwrite_1to_1file, tmp7709)

tmp7711 := PrimCons(symy_1or_1n_2, tmp7710)

tmp7712 := PrimCons(tmp7504, tmp7711)

tmp7713 := PrimCons(sym_6_6, tmp7712)

tmp7714 := PrimCons(sym_5, tmp7713)

tmp7715 := PrimCons(sym_5_a, tmp7714)

tmp7716 := PrimCons(sym_7, tmp7715)

tmp7717 := PrimCons(sym_d, tmp7716)

tmp7718 := PrimCons(sym_c, tmp7717)

tmp7719 := PrimCons(sym_1, tmp7718)

tmp7720 := PrimCons(sym_3, tmp7719)

tmp7721 := PrimCons(symshen_4_5end_6, tmp7720)

tmp7722 := PrimCons(sym_5_b_6, tmp7721)

tmp7723 := PrimCons(sym_c_4, tmp7722)

tmp7724 := PrimCons(sym_a_a_6, tmp7723)

tmp7725 := PrimCons(sym_6, tmp7724)

tmp7726 := PrimCons(sym_6_a, tmp7725)

tmp7727 := PrimCons(sym_a, tmp7726)

tmp7728 := PrimCons(sym_a_a, tmp7727)

tmp7729 := PrimCons(sym_5e_6, tmp7728)

tmp7730 := PrimCons(sym_1_6, tmp7729)

tmp7731 := PrimCons(sym_5_1, tmp7730)

tmp7732 := PrimCons(sym_dhush_d, tmp7731)

tmp7733 := PrimCons(sym_dporters_d, tmp7732)

tmp7734 := PrimCons(sym_dport_d, tmp7733)

tmp7735 := PrimCons(sym_8s, tmp7734)

tmp7736 := PrimCons(sym_8p, tmp7735)

tmp7737 := PrimCons(sym_8v, tmp7736)

tmp7738 := PrimCons(sym_dproperty_1vector_d, tmp7737)

tmp7739 := PrimCons(sym_drelease_d, tmp7738)

tmp7740 := PrimCons(sym_dos_d, tmp7739)

tmp7741 := PrimCons(sym_dmacros_d, tmp7740)

tmp7742 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp7741)

tmp7743 := PrimCons(sym_dversion_d, tmp7742)

tmp7744 := PrimCons(sym_dhome_1directory_d, tmp7743)

tmp7745 := PrimCons(sym_dstoutput_d, tmp7744)

tmp7746 := PrimCons(sym_dstinput_d, tmp7745)

tmp7747 := PrimCons(sym_dimplementation_d, tmp7746)

tmp7748 := PrimCons(sym_dlanguage_d, tmp7747)

tmp7749 := PrimCons(sym__, tmp7748)

tmp7750 := PrimCons(tmp7503, tmp7749)

tmp7751 := PrimCons(tmp7502, tmp7750)

tmp7752 := PrimCons(tmp7501, tmp7751)

tmp7753 := PrimCons(tmp7500, tmp7752)

tmp7754 := PrimCons(sym_e_e, tmp7753)

tmp7755 := PrimCons(sym_5_1_1, tmp7754)

tmp7756 := PrimCons(sym_1_1_6, tmp7755)

tmp7757 := PrimCons(sym_i, tmp7756)

tmp7758 := PrimCons(sym_j, tmp7757)

tmp7759 := PrimCons(sym_b, tmp7758)

tmp7760 := PrimValue(sym_dproperty_1vector_d)

tmp7761 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp7759, tmp7760)


_ = tmp7761

tmp7762 := MakeNative(func(__e *ControlFlow) {
V1685 := __e.Get(1)
_ = V1685
tmp7763 := MakeNative(func(__e *ControlFlow) {
ArityF := __e.Get(1)
_ = ArityF
tmp7771 := PrimEqual(ArityF, MakeNumber(-1))

var ifres7768 Obj

if True == tmp7771 {
ifres7768 = True


} else {
tmp7770 := PrimEqual(ArityF, MakeNumber(0))

var ifres7769 Obj

if True == tmp7770 {
ifres7769 = True


} else {
ifres7769 = False


}

ifres7768 = ifres7769


}

if True == ifres7768 {
__e.Return(Nil)
return
} else {
tmp7764 := PrimCons(V1685, Nil)

tmp7765 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp7764, ArityF)


tmp7766 := Call(__e, PrimFunc(symeval_1kl), tmp7765)


__e.Return(PrimCons(V1685, tmp7766))
return


}


}, 1)

tmp7772 := Call(__e, PrimFunc(symarity), V1685)


__e.TailApply(tmp7763, tmp7772)
return


}, 1)

tmp7773 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp7762)


_ = tmp7773

tmp7774 := MakeNative(func(__e *ControlFlow) {
V1686 := __e.Get(1)
_ = V1686
tmp7775 := MakeNative(func(__e *ControlFlow) {
LambdaEntries := __e.Get(1)
_ = LambdaEntries
tmp7776 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4tuple), X)
return
}, 1)

tmp7777 := PrimCons(symshen_4tuple, tmp7776)

tmp7778 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4pvar), X)
return
}, 1)

tmp7779 := PrimCons(symshen_4pvar, tmp7778)

tmp7780 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), X)
return
}, 1)

tmp7781 := PrimCons(symshen_4print_1prolog_1vector, tmp7780)

tmp7782 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4print_1freshterm), X)
return
}, 1)

tmp7783 := PrimCons(symshen_4print_1freshterm, tmp7782)

tmp7784 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4printF), X)
return
}, 1)

tmp7785 := PrimCons(symshen_4printF, tmp7784)

tmp7786 := PrimCons(tmp7785, LambdaEntries)

tmp7787 := PrimCons(tmp7783, tmp7786)

tmp7788 := PrimCons(tmp7781, tmp7787)

tmp7789 := PrimCons(tmp7779, tmp7788)

tmp7790 := PrimCons(tmp7777, tmp7789)

__e.Return(PrimSet(symshen_4_dlambdatable_d, tmp7790))
return


}, 1)

tmp7791 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4lambda_1entry), X)
return
}, 1)

tmp7792 := Call(__e, PrimFunc(symmap), tmp7791, V1686)


__e.TailApply(tmp7775, tmp7792)
return


}, 1)

tmp7793 := Call(__e, ns2_1set, symshen_4build_1lambda_1table, tmp7774)


_ = tmp7793

tmp7794 := Call(__e, PrimFunc(symexternal), symshen)


__e.TailApply(PrimFunc(symshen_4build_1lambda_1table), tmp7794)
return




}, 0)

