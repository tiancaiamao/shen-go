package main

import . "github.com/tiancaiamao/cora/cora_go"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
	tmp5976 := PrimNS3Set(symshen_4_dhistory_d, Nil)

	_ = tmp5976

	tmp5977 := PrimNS3Set(symshen_4_dtc_d, False)

	_ = tmp5977

	tmp5978 := Call(__e, PrimNS2Value(symvector), MakeNumber(20000))

	tmp5979 := PrimNS3Set(sym_dproperty_1vector_d, tmp5978)

	_ = tmp5979

	tmp5980 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4defmacro_1macro), X)
		return
	}, 1)

	tmp5981 := PrimCons(symshen_4defmacro_1macro, tmp5980)

	tmp5982 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4timer_1macro), X)
		return
	}, 1)

	tmp5983 := PrimCons(symshen_4timer_1macro, tmp5982)

	tmp5984 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4cases_1macro), X)
		return
	}, 1)

	tmp5985 := PrimCons(symshen_4cases_1macro, tmp5984)

	tmp5986 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4abs_1macro), X)
		return
	}, 1)

	tmp5987 := PrimCons(symshen_4abs_1macro, tmp5986)

	tmp5988 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4put_cget_1macro), X)
		return
	}, 1)

	tmp5989 := PrimCons(symshen_4put_cget_1macro, tmp5988)

	tmp5990 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4datatype_1macro), X)
		return
	}, 1)

	tmp5991 := PrimCons(symshen_4datatype_1macro, tmp5990)

	tmp5992 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4let_1macro), X)
		return
	}, 1)

	tmp5993 := PrimCons(symshen_4let_1macro, tmp5992)

	tmp5994 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4assoc_1macro), X)
		return
	}, 1)

	tmp5995 := PrimCons(symshen_4assoc_1macro, tmp5994)

	tmp5996 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4make_1string_1macro), X)
		return
	}, 1)

	tmp5997 := PrimCons(symmake_1string, tmp5996)

	tmp5998 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4output_1macro), X)
		return
	}, 1)

	tmp5999 := PrimCons(symshen_4output_1macro, tmp5998)

	tmp6000 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4input_1macro), X)
		return
	}, 1)

	tmp6001 := PrimCons(symshen_4input_1macro, tmp6000)

	tmp6002 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4error_1macro), X)
		return
	}, 1)

	tmp6003 := PrimCons(symshen_4error_1macro, tmp6002)

	tmp6004 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4prolog_1macro), X)
		return
	}, 1)

	tmp6005 := PrimCons(symshen_4prolog_1macro, tmp6004)

	tmp6006 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4synonyms_1macro), X)
		return
	}, 1)

	tmp6007 := PrimCons(symshen_4synonyms_1macro, tmp6006)

	tmp6008 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4nl_1macro), X)
		return
	}, 1)

	tmp6009 := PrimCons(symshen_4nl_1macro, tmp6008)

	tmp6010 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4_8s_1macro), X)
		return
	}, 1)

	tmp6011 := PrimCons(symshen_4_8s_1macro, tmp6010)

	tmp6012 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4defprolog_1macro), X)
		return
	}, 1)

	tmp6013 := PrimCons(symdefprolog, tmp6012)

	tmp6014 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4defcc_1macro), X)
		return
	}, 1)

	tmp6015 := PrimCons(symshen_4defcc_1macro, tmp6014)

	tmp6016 := MakeNative(func(__e *ControlFlow) {
		X := __e.Get(1)
		_ = X
		__e.TailApply(PrimNS2Value(symshen_4u_b_1macro), X)
		return
	}, 1)

	tmp6017 := PrimCons(symshen_4u_b_1macro, tmp6016)

	tmp6018 := PrimCons(tmp6017, Nil)

	tmp6019 := PrimCons(tmp6015, tmp6018)

	tmp6020 := PrimCons(tmp6013, tmp6019)

	tmp6021 := PrimCons(tmp6011, tmp6020)

	tmp6022 := PrimCons(tmp6009, tmp6021)

	tmp6023 := PrimCons(tmp6007, tmp6022)

	tmp6024 := PrimCons(tmp6005, tmp6023)

	tmp6025 := PrimCons(tmp6003, tmp6024)

	tmp6026 := PrimCons(tmp6001, tmp6025)

	tmp6027 := PrimCons(tmp5999, tmp6026)

	tmp6028 := PrimCons(tmp5997, tmp6027)

	tmp6029 := PrimCons(tmp5995, tmp6028)

	tmp6030 := PrimCons(tmp5993, tmp6029)

	tmp6031 := PrimCons(tmp5991, tmp6030)

	tmp6032 := PrimCons(tmp5989, tmp6031)

	tmp6033 := PrimCons(tmp5987, tmp6032)

	tmp6034 := PrimCons(tmp5985, tmp6033)

	tmp6035 := PrimCons(tmp5983, tmp6034)

	tmp6036 := PrimCons(tmp5981, tmp6035)

	tmp6037 := PrimNS3Set(sym_dmacros_d, tmp6036)

	_ = tmp6037

	tmp6038 := PrimNS3Set(symshen_4_dgensym_d, MakeNumber(0))

	_ = tmp6038

	tmp6039 := PrimNS3Set(symshen_4_dtracking_d, Nil)

	_ = tmp6039

	tmp6040 := PrimNS3Set(symshen_4_dprofiled_d, Nil)

	_ = tmp6040

	tmp6041 := PrimNS3Set(sym_dhome_1directory_d, MakeString(""))

	_ = tmp6041

	tmp6042 := PrimCons(symtype, Nil)

	tmp6043 := PrimCons(syminput_7, tmp6042)

	tmp6044 := PrimCons(symopen, tmp6043)

	tmp6045 := PrimCons(symset, tmp6044)

	tmp6046 := PrimCons(symwhere, tmp6045)

	tmp6047 := PrimCons(symlet, tmp6046)

	tmp6048 := PrimCons(symlambda, tmp6047)

	tmp6049 := PrimCons(symcons, tmp6048)

	tmp6050 := PrimCons(sym_8v, tmp6049)

	tmp6051 := PrimCons(sym_8s, tmp6050)

	tmp6052 := PrimCons(sym_8p, tmp6051)

	tmp6053 := PrimNS3Set(symshen_4_dspecial_d, tmp6052)

	_ = tmp6053

	tmp6054 := PrimNS3Set(symshen_4_dextraspecial_d, Nil)

	_ = tmp6054

	tmp6055 := PrimNS3Set(symshen_4_dspy_d, False)

	_ = tmp6055

	tmp6056 := PrimNS3Set(symshen_4_ddatatypes_d, Nil)

	_ = tmp6056

	tmp6057 := PrimNS3Set(symshen_4_dalldatatypes_d, Nil)

	_ = tmp6057

	tmp6058 := PrimNS3Set(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

	_ = tmp6058

	tmp6059 := PrimNS3Set(symshen_4_dpackage_d, symnull)

	_ = tmp6059

	tmp6060 := PrimNS3Set(symshen_4_dsynonyms_d, Nil)

	_ = tmp6060

	tmp6061 := PrimNS3Set(symshen_4_dsystem_d, Nil)

	_ = tmp6061

	tmp6062 := PrimNS3Set(symshen_4_dsigf_d, Nil)

	_ = tmp6062

	tmp6063 := PrimNS3Set(symshen_4_doccurs_d, True)

	_ = tmp6063

	tmp6064 := PrimNS3Set(symshen_4_dfactorise_2_d, False)

	_ = tmp6064

	tmp6065 := PrimNS3Set(symshen_4_dmaxinferences_d, MakeNumber(1000000))

	_ = tmp6065

	tmp6066 := PrimNS3Set(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

	_ = tmp6066

	tmp6067 := PrimNS3Set(symshen_4_dcall_d, MakeNumber(0))

	_ = tmp6067

	tmp6068 := PrimNS3Set(symshen_4_dinfs_d, MakeNumber(0))

	_ = tmp6068

	tmp6069 := PrimNS3Set(sym_dhush_d, False)

	_ = tmp6069

	tmp6070 := PrimNS3Set(symshen_4_doptimise_d, False)

	_ = tmp6070

	tmp6071 := PrimNS3Set(sym_dversion_d, MakeString("31"))

	_ = tmp6071

	tmp6072 := PrimNS3Set(symshen_4_dstep_d, False)

	_ = tmp6072

	tmp6073 := PrimNS3Set(symshen_4_dit_d, MakeString(""))

	_ = tmp6073

	tmp6074 := PrimNS3Set(symshen_4_dresidue_d, Nil)

	_ = tmp6074

	tmp6075 := MakeNative(func(__e *ControlFlow) {
		V1674 := __e.Get(1)
		_ = V1674
		tmp6076 := MakeNative(func(__e *ControlFlow) {
			Bindings := __e.Get(1)
			_ = Bindings
			tmp6077 := MakeNative(func(__e *ControlFlow) {
				PrintNamed := __e.Get(1)
				_ = PrintNamed
				tmp6078 := MakeNative(func(__e *ControlFlow) {
					Ticketed := __e.Get(1)
					_ = Ticketed
					tmp6079 := MakeNative(func(__e *ControlFlow) {
						Assign := __e.Get(1)
						_ = Assign
						__e.Return(V1674)
						return
					}, 1)

					tmp6080 := PrimNS3Set(symshen_4_dprolog_1vector_d, Ticketed)

					__e.TailApply(tmp6079, tmp6080)
					return

				}, 1)

				tmp6081 := PrimVectorSet(Bindings, MakeNumber(1), MakeNumber(2))

				__e.TailApply(tmp6078, tmp6081)
				return

			}, 1)

			tmp6082 := PrimVectorSet(Bindings, MakeNumber(0), symshen_4print_1prolog_1vector)

			__e.TailApply(tmp6077, tmp6082)
			return

		}, 1)

		tmp6083 := PrimAbsvector(V1674)

		__e.TailApply(tmp6076, tmp6083)
		return

	}, 1)

	tmp6084 := Call(__e, PrimNS2Value(symdef), symprolog_1memory, tmp6075)

	_ = tmp6084

	tmp6085 := Call(__e, PrimNS2Value(symprolog_1memory), MakeNumber(10000))

	_ = tmp6085

	tmp6086 := PrimNS3Set(symshen_4_dloading_2_d, False)

	_ = tmp6086

	tmp6087 := MakeNative(func(__e *ControlFlow) {
		V1677 := __e.Get(1)
		_ = V1677
		tmp6103 := PrimEqual(Nil, V1677)

		if True == tmp6103 {
			__e.Return(Nil)
			return
		} else {
			tmp6102 := PrimIsPair(V1677)

			var ifres6098 Obj

			if True == tmp6102 {
				tmp6100 := PrimTail(V1677)

				tmp6101 := PrimIsPair(tmp6100)

				var ifres6099 Obj

				if True == tmp6101 {
					ifres6099 = True

				} else {
					ifres6099 = False

				}

				ifres6098 = ifres6099

			} else {
				ifres6098 = False

			}

			if True == ifres6098 {
				tmp6090 := MakeNative(func(__e *ControlFlow) {
					DecArity := __e.Get(1)
					_ = DecArity
					tmp6091 := PrimTail(V1677)

					tmp6092 := PrimTail(tmp6091)

					__e.TailApply(PrimNS2Value(symshen_4initialise_1arity_1table), tmp6092)
					return

				}, 1)

				tmp6093 := PrimHead(V1677)

				tmp6094 := PrimTail(V1677)

				tmp6095 := PrimHead(tmp6094)

				tmp6096 := PrimNS3Value(sym_dproperty_1vector_d)

				tmp6097 := Call(__e, PrimNS2Value(symput), tmp6093, symarity, tmp6095, tmp6096)

				__e.TailApply(tmp6090, tmp6097)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
				return
			}

		}

	}, 1)

	tmp6104 := Call(__e, PrimNS2Value(symdef), symshen_4initialise_1lambda_1tables, tmp6087)

	_ = tmp6104

	tmp6105 := MakeNative(func(__e *ControlFlow) {
		V1678 := __e.Get(1)
		_ = V1678
		tmp6106 := MakeNative(func(__e *ControlFlow) {
			tmp6107 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V1678, symarity, tmp6107)
			return

		}, 0)

		tmp6108 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(-1))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp6106, tmp6108)
		return

	}, 1)

	tmp6109 := Call(__e, PrimNS2Value(symdef), symarity, tmp6105)

	_ = tmp6109

	tmp6110 := MakeNative(func(__e *ControlFlow) {
		V1681 := __e.Get(1)
		_ = V1681
		tmp6126 := PrimEqual(Nil, V1681)

		if True == tmp6126 {
			__e.Return(Nil)
			return
		} else {
			tmp6125 := PrimIsPair(V1681)

			var ifres6121 Obj

			if True == tmp6125 {
				tmp6123 := PrimTail(V1681)

				tmp6124 := PrimIsPair(tmp6123)

				var ifres6122 Obj

				if True == tmp6124 {
					ifres6122 = True

				} else {
					ifres6122 = False

				}

				ifres6121 = ifres6122

			} else {
				ifres6121 = False

			}

			if True == ifres6121 {
				tmp6113 := MakeNative(func(__e *ControlFlow) {
					DecArity := __e.Get(1)
					_ = DecArity
					tmp6114 := PrimTail(V1681)

					tmp6115 := PrimTail(tmp6114)

					__e.TailApply(PrimNS2Value(symshen_4initialise_1arity_1table), tmp6115)
					return

				}, 1)

				tmp6116 := PrimHead(V1681)

				tmp6117 := PrimTail(V1681)

				tmp6118 := PrimHead(tmp6117)

				tmp6119 := PrimNS3Value(sym_dproperty_1vector_d)

				tmp6120 := Call(__e, PrimNS2Value(symput), tmp6116, symarity, tmp6118, tmp6119)

				__e.TailApply(tmp6113, tmp6120)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise_arity_table")))
				return
			}

		}

	}, 1)

	tmp6127 := Call(__e, PrimNS2Value(symdef), symshen_4initialise_1arity_1table, tmp6110)

	_ = tmp6127

	tmp6128 := PrimCons(MakeNumber(2), Nil)

	tmp6129 := PrimCons(sym_8s, tmp6128)

	tmp6130 := PrimCons(MakeNumber(2), tmp6129)

	tmp6131 := PrimCons(sym_8v, tmp6130)

	tmp6132 := PrimCons(MakeNumber(2), tmp6131)

	tmp6133 := PrimCons(sym_8p, tmp6132)

	tmp6134 := PrimCons(MakeNumber(1), tmp6133)

	tmp6135 := PrimCons(sym_5_b_6, tmp6134)

	tmp6136 := PrimCons(MakeNumber(1), tmp6135)

	tmp6137 := PrimCons(symshen_4_5end_6, tmp6136)

	tmp6138 := PrimCons(MakeNumber(1), tmp6137)

	tmp6139 := PrimCons(sym_5e_6, tmp6138)

	tmp6140 := PrimCons(MakeNumber(2), tmp6139)

	tmp6141 := PrimCons(sym_a_a, tmp6140)

	tmp6142 := PrimCons(MakeNumber(2), tmp6141)

	tmp6143 := PrimCons(sym_1, tmp6142)

	tmp6144 := PrimCons(MakeNumber(2), tmp6143)

	tmp6145 := PrimCons(sym_c, tmp6144)

	tmp6146 := PrimCons(MakeNumber(2), tmp6145)

	tmp6147 := PrimCons(sym_d, tmp6146)

	tmp6148 := PrimCons(MakeNumber(2), tmp6147)

	tmp6149 := PrimCons(sym_7, tmp6148)

	tmp6150 := PrimCons(MakeNumber(1), tmp6149)

	tmp6151 := PrimCons(symy_1or_1n_2, tmp6150)

	tmp6152 := PrimCons(MakeNumber(2), tmp6151)

	tmp6153 := PrimCons(symwrite_1to_1file, tmp6152)

	tmp6154 := PrimCons(MakeNumber(2), tmp6153)

	tmp6155 := PrimCons(symwrite_1byte, tmp6154)

	tmp6156 := PrimCons(MakeNumber(5), tmp6155)

	tmp6157 := PrimCons(symwhen, tmp6156)

	tmp6158 := PrimCons(MakeNumber(0), tmp6157)

	tmp6159 := PrimCons(symversion, tmp6158)

	tmp6160 := PrimCons(MakeNumber(5), tmp6159)

	tmp6161 := PrimCons(symvar_2, tmp6160)

	tmp6162 := PrimCons(MakeNumber(1), tmp6161)

	tmp6163 := PrimCons(symvariable_2, tmp6162)

	tmp6164 := PrimCons(MakeNumber(1), tmp6163)

	tmp6165 := PrimCons(symvalue, tmp6164)

	tmp6166 := PrimCons(MakeNumber(3), tmp6165)

	tmp6167 := PrimCons(symvector_1_6, tmp6166)

	tmp6168 := PrimCons(MakeNumber(1), tmp6167)

	tmp6169 := PrimCons(symvector_2, tmp6168)

	tmp6170 := PrimCons(MakeNumber(1), tmp6169)

	tmp6171 := PrimCons(symvector, tmp6170)

	tmp6172 := PrimCons(MakeNumber(2), tmp6171)

	tmp6173 := PrimCons(symupdate_1lambda_1table, tmp6172)

	tmp6174 := PrimCons(MakeNumber(1), tmp6173)

	tmp6175 := PrimCons(symundefmacro, tmp6174)

	tmp6176 := PrimCons(MakeNumber(1), tmp6175)

	tmp6177 := PrimCons(symuntrack, tmp6176)

	tmp6178 := PrimCons(MakeNumber(2), tmp6177)

	tmp6179 := PrimCons(symunion, tmp6178)

	tmp6180 := PrimCons(MakeNumber(1), tmp6179)

	tmp6181 := PrimCons(symunprofile, tmp6180)

	tmp6182 := PrimCons(MakeNumber(3), tmp6181)

	tmp6183 := PrimCons(symunput, tmp6182)

	tmp6184 := PrimCons(MakeNumber(1), tmp6183)

	tmp6185 := PrimCons(symundefmacro, tmp6184)

	tmp6186 := PrimCons(MakeNumber(5), tmp6185)

	tmp6187 := PrimCons(symreturn, tmp6186)

	tmp6188 := PrimCons(MakeNumber(2), tmp6187)

	tmp6189 := PrimCons(symtype, tmp6188)

	tmp6190 := PrimCons(MakeNumber(1), tmp6189)

	tmp6191 := PrimCons(symtuple_2, tmp6190)

	tmp6192 := PrimCons(MakeNumber(2), tmp6191)

	tmp6193 := PrimCons(symtrap_1error, tmp6192)

	tmp6194 := PrimCons(MakeNumber(1), tmp6193)

	tmp6195 := PrimCons(symtrack, tmp6194)

	tmp6196 := PrimCons(MakeNumber(1), tmp6195)

	tmp6197 := PrimCons(symtlstr, tmp6196)

	tmp6198 := PrimCons(MakeNumber(1), tmp6197)

	tmp6199 := PrimCons(symthaw, tmp6198)

	tmp6200 := PrimCons(MakeNumber(0), tmp6199)

	tmp6201 := PrimCons(symtc_2, tmp6200)

	tmp6202 := PrimCons(MakeNumber(1), tmp6201)

	tmp6203 := PrimCons(symtc, tmp6202)

	tmp6204 := PrimCons(MakeNumber(1), tmp6203)

	tmp6205 := PrimCons(symtl, tmp6204)

	tmp6206 := PrimCons(MakeNumber(1), tmp6205)

	tmp6207 := PrimCons(symtail, tmp6206)

	tmp6208 := PrimCons(MakeNumber(1), tmp6207)

	tmp6209 := PrimCons(symsystemf, tmp6208)

	tmp6210 := PrimCons(MakeNumber(1), tmp6209)

	tmp6211 := PrimCons(symsymbol_2, tmp6210)

	tmp6212 := PrimCons(MakeNumber(1), tmp6211)

	tmp6213 := PrimCons(symsum, tmp6212)

	tmp6214 := PrimCons(MakeNumber(3), tmp6213)

	tmp6215 := PrimCons(symsubst, tmp6214)

	tmp6216 := PrimCons(MakeNumber(1), tmp6215)

	tmp6217 := PrimCons(symstring_2, tmp6216)

	tmp6218 := PrimCons(MakeNumber(1), tmp6217)

	tmp6219 := PrimCons(symstring_1_6symbol, tmp6218)

	tmp6220 := PrimCons(MakeNumber(1), tmp6219)

	tmp6221 := PrimCons(symstring_1_6n, tmp6220)

	tmp6222 := PrimCons(MakeNumber(1), tmp6221)

	tmp6223 := PrimCons(symstr, tmp6222)

	tmp6224 := PrimCons(MakeNumber(0), tmp6223)

	tmp6225 := PrimCons(symstoutput, tmp6224)

	tmp6226 := PrimCons(MakeNumber(0), tmp6225)

	tmp6227 := PrimCons(symstinput, tmp6226)

	tmp6228 := PrimCons(MakeNumber(1), tmp6227)

	tmp6229 := PrimCons(symstep, tmp6228)

	tmp6230 := PrimCons(MakeNumber(1), tmp6229)

	tmp6231 := PrimCons(symspy, tmp6230)

	tmp6232 := PrimCons(MakeNumber(2), tmp6231)

	tmp6233 := PrimCons(symspecialise, tmp6232)

	tmp6234 := PrimCons(MakeNumber(1), tmp6233)

	tmp6235 := PrimCons(symsnd, tmp6234)

	tmp6236 := PrimCons(MakeNumber(1), tmp6235)

	tmp6237 := PrimCons(symsimple_1error, tmp6236)

	tmp6238 := PrimCons(MakeNumber(2), tmp6237)

	tmp6239 := PrimCons(symset, tmp6238)

	tmp6240 := PrimCons(MakeNumber(1), tmp6239)

	tmp6241 := PrimCons(symreverse, tmp6240)

	tmp6242 := PrimCons(MakeNumber(2), tmp6241)

	tmp6243 := PrimCons(symremove, tmp6242)

	tmp6244 := PrimCons(MakeNumber(0), tmp6243)

	tmp6245 := PrimCons(symrelease, tmp6244)

	tmp6246 := PrimCons(MakeNumber(1), tmp6245)

	tmp6247 := PrimCons(symreceive, tmp6246)

	tmp6248 := PrimCons(MakeNumber(1), tmp6247)

	tmp6249 := PrimCons(symshen_4read_1unit_1string, tmp6248)

	tmp6250 := PrimCons(MakeNumber(1), tmp6249)

	tmp6251 := PrimCons(symread_1from_1string_1unprocessed, tmp6250)

	tmp6252 := PrimCons(MakeNumber(1), tmp6251)

	tmp6253 := PrimCons(symread_1from_1string, tmp6252)

	tmp6254 := PrimCons(MakeNumber(1), tmp6253)

	tmp6255 := PrimCons(symread_1byte, tmp6254)

	tmp6256 := PrimCons(MakeNumber(1), tmp6255)

	tmp6257 := PrimCons(symread, tmp6256)

	tmp6258 := PrimCons(MakeNumber(1), tmp6257)

	tmp6259 := PrimCons(symread_1file, tmp6258)

	tmp6260 := PrimCons(MakeNumber(1), tmp6259)

	tmp6261 := PrimCons(symread_1file_1as_1bytelist, tmp6260)

	tmp6262 := PrimCons(MakeNumber(1), tmp6261)

	tmp6263 := PrimCons(symread_1file_1as_1string, tmp6262)

	tmp6264 := PrimCons(MakeNumber(4), tmp6263)

	tmp6265 := PrimCons(symput, tmp6264)

	tmp6266 := PrimCons(MakeNumber(1), tmp6265)

	tmp6267 := PrimCons(symprotect, tmp6266)

	tmp6268 := PrimCons(MakeNumber(1), tmp6267)

	tmp6269 := PrimCons(sympreclude_1all_1but, tmp6268)

	tmp6270 := PrimCons(MakeNumber(1), tmp6269)

	tmp6271 := PrimCons(sympreclude, tmp6270)

	tmp6272 := PrimCons(MakeNumber(1), tmp6271)

	tmp6273 := PrimCons(symps, tmp6272)

	tmp6274 := PrimCons(MakeNumber(2), tmp6273)

	tmp6275 := PrimCons(sympr, tmp6274)

	tmp6276 := PrimCons(MakeNumber(1), tmp6275)

	tmp6277 := PrimCons(symprofile_1results, tmp6276)

	tmp6278 := PrimCons(MakeNumber(1), tmp6277)

	tmp6279 := PrimCons(symprolog_1memory, tmp6278)

	tmp6280 := PrimCons(MakeNumber(1), tmp6279)

	tmp6281 := PrimCons(symshen_4printF, tmp6280)

	tmp6282 := PrimCons(MakeNumber(1), tmp6281)

	tmp6283 := PrimCons(symshen_4print_1freshterm, tmp6282)

	tmp6284 := PrimCons(MakeNumber(1), tmp6283)

	tmp6285 := PrimCons(symshen_4print_1prolog_1vector, tmp6284)

	tmp6286 := PrimCons(MakeNumber(1), tmp6285)

	tmp6287 := PrimCons(symprofile, tmp6286)

	tmp6288 := PrimCons(MakeNumber(1), tmp6287)

	tmp6289 := PrimCons(symshen_4pprint, tmp6288)

	tmp6290 := PrimCons(MakeNumber(1), tmp6289)

	tmp6291 := PrimCons(symprint, tmp6290)

	tmp6292 := PrimCons(MakeNumber(1), tmp6291)

	tmp6293 := PrimCons(sympreclude_1all_1but, tmp6292)

	tmp6294 := PrimCons(MakeNumber(2), tmp6293)

	tmp6295 := PrimCons(sympos, tmp6294)

	tmp6296 := PrimCons(MakeNumber(0), tmp6295)

	tmp6297 := PrimCons(symporters, tmp6296)

	tmp6298 := PrimCons(MakeNumber(0), tmp6297)

	tmp6299 := PrimCons(symport, tmp6298)

	tmp6300 := PrimCons(MakeNumber(1), tmp6299)

	tmp6301 := PrimCons(sympackage_2, tmp6300)

	tmp6302 := PrimCons(MakeNumber(3), tmp6301)

	tmp6303 := PrimCons(sympackage, tmp6302)

	tmp6304 := PrimCons(MakeNumber(0), tmp6303)

	tmp6305 := PrimCons(symos, tmp6304)

	tmp6306 := PrimCons(MakeNumber(2), tmp6305)

	tmp6307 := PrimCons(symor, tmp6306)

	tmp6308 := PrimCons(MakeNumber(1), tmp6307)

	tmp6309 := PrimCons(symoptimise, tmp6308)

	tmp6310 := PrimCons(MakeNumber(2), tmp6309)

	tmp6311 := PrimCons(symopen, tmp6310)

	tmp6312 := PrimCons(MakeNumber(1), tmp6311)

	tmp6313 := PrimCons(symoccurs_1check, tmp6312)

	tmp6314 := PrimCons(MakeNumber(2), tmp6313)

	tmp6315 := PrimCons(symoccurrences, tmp6314)

	tmp6316 := PrimCons(MakeNumber(1), tmp6315)

	tmp6317 := PrimCons(symoccurs_1check, tmp6316)

	tmp6318 := PrimCons(MakeNumber(1), tmp6317)

	tmp6319 := PrimCons(symnumber_2, tmp6318)

	tmp6320 := PrimCons(MakeNumber(1), tmp6319)

	tmp6321 := PrimCons(symn_1_6string, tmp6320)

	tmp6322 := PrimCons(MakeNumber(2), tmp6321)

	tmp6323 := PrimCons(symnth, tmp6322)

	tmp6324 := PrimCons(MakeNumber(1), tmp6323)

	tmp6325 := PrimCons(symnot, tmp6324)

	tmp6326 := PrimCons(MakeNumber(1), tmp6325)

	tmp6327 := PrimCons(symnl, tmp6326)

	tmp6328 := PrimCons(MakeNumber(1), tmp6327)

	tmp6329 := PrimCons(symmaxinferences, tmp6328)

	tmp6330 := PrimCons(MakeNumber(2), tmp6329)

	tmp6331 := PrimCons(symmapcan, tmp6330)

	tmp6332 := PrimCons(MakeNumber(2), tmp6331)

	tmp6333 := PrimCons(symmap, tmp6332)

	tmp6334 := PrimCons(MakeNumber(1), tmp6333)

	tmp6335 := PrimCons(symmacroexpand, tmp6334)

	tmp6336 := PrimCons(MakeNumber(1), tmp6335)

	tmp6337 := PrimCons(symvector, tmp6336)

	tmp6338 := PrimCons(MakeNumber(2), tmp6337)

	tmp6339 := PrimCons(sym_5_a, tmp6338)

	tmp6340 := PrimCons(MakeNumber(2), tmp6339)

	tmp6341 := PrimCons(sym_5, tmp6340)

	tmp6342 := PrimCons(MakeNumber(1), tmp6341)

	tmp6343 := PrimCons(symload, tmp6342)

	tmp6344 := PrimCons(MakeNumber(1), tmp6343)

	tmp6345 := PrimCons(symlist, tmp6344)

	tmp6346 := PrimCons(MakeNumber(1), tmp6345)

	tmp6347 := PrimCons(symlineread, tmp6346)

	tmp6348 := PrimCons(MakeNumber(1), tmp6347)

	tmp6349 := PrimCons(symlimit, tmp6348)

	tmp6350 := PrimCons(MakeNumber(1), tmp6349)

	tmp6351 := PrimCons(symlength, tmp6350)

	tmp6352 := PrimCons(MakeNumber(0), tmp6351)

	tmp6353 := PrimCons(symlanguage, tmp6352)

	tmp6354 := PrimCons(MakeNumber(6), tmp6353)

	tmp6355 := PrimCons(symis_b, tmp6354)

	tmp6356 := PrimCons(MakeNumber(6), tmp6355)

	tmp6357 := PrimCons(symis, tmp6356)

	tmp6358 := PrimCons(MakeNumber(0), tmp6357)

	tmp6359 := PrimCons(symit, tmp6358)

	tmp6360 := PrimCons(MakeNumber(1), tmp6359)

	tmp6361 := PrimCons(syminternal, tmp6360)

	tmp6362 := PrimCons(MakeNumber(2), tmp6361)

	tmp6363 := PrimCons(symintersection, tmp6362)

	tmp6364 := PrimCons(MakeNumber(1), tmp6363)

	tmp6365 := PrimCons(syminclude_1all_1but, tmp6364)

	tmp6366 := PrimCons(MakeNumber(0), tmp6365)

	tmp6367 := PrimCons(symimplementation, tmp6366)

	tmp6368 := PrimCons(MakeNumber(2), tmp6367)

	tmp6369 := PrimCons(syminput_7, tmp6368)

	tmp6370 := PrimCons(MakeNumber(1), tmp6369)

	tmp6371 := PrimCons(syminput, tmp6370)

	tmp6372 := PrimCons(MakeNumber(0), tmp6371)

	tmp6373 := PrimCons(syminferences, tmp6372)

	tmp6374 := PrimCons(MakeNumber(1), tmp6373)

	tmp6375 := PrimCons(symintern, tmp6374)

	tmp6376 := PrimCons(MakeNumber(1), tmp6375)

	tmp6377 := PrimCons(syminternal, tmp6376)

	tmp6378 := PrimCons(MakeNumber(1), tmp6377)

	tmp6379 := PrimCons(syminteger_2, tmp6378)

	tmp6380 := PrimCons(MakeNumber(1), tmp6379)

	tmp6381 := PrimCons(symin_1package, tmp6380)

	tmp6382 := PrimCons(MakeNumber(1), tmp6381)

	tmp6383 := PrimCons(syminclude, tmp6382)

	tmp6384 := PrimCons(MakeNumber(3), tmp6383)

	tmp6385 := PrimCons(symif, tmp6384)

	tmp6386 := PrimCons(MakeNumber(1), tmp6385)

	tmp6387 := PrimCons(symhead, tmp6386)

	tmp6388 := PrimCons(MakeNumber(1), tmp6387)

	tmp6389 := PrimCons(symhdstr, tmp6388)

	tmp6390 := PrimCons(MakeNumber(1), tmp6389)

	tmp6391 := PrimCons(symhdv, tmp6390)

	tmp6392 := PrimCons(MakeNumber(1), tmp6391)

	tmp6393 := PrimCons(symhd, tmp6392)

	tmp6394 := PrimCons(MakeNumber(2), tmp6393)

	tmp6395 := PrimCons(symhash, tmp6394)

	tmp6396 := PrimCons(MakeNumber(2), tmp6395)

	tmp6397 := PrimCons(sym_a, tmp6396)

	tmp6398 := PrimCons(MakeNumber(2), tmp6397)

	tmp6399 := PrimCons(sym_6_a, tmp6398)

	tmp6400 := PrimCons(MakeNumber(2), tmp6399)

	tmp6401 := PrimCons(sym_6, tmp6400)

	tmp6402 := PrimCons(MakeNumber(2), tmp6401)

	tmp6403 := PrimCons(sym_5_1vector, tmp6402)

	tmp6404 := PrimCons(MakeNumber(2), tmp6403)

	tmp6405 := PrimCons(sym_5_1address, tmp6404)

	tmp6406 := PrimCons(MakeNumber(3), tmp6405)

	tmp6407 := PrimCons(symaddress_1_6, tmp6406)

	tmp6408 := PrimCons(MakeNumber(1), tmp6407)

	tmp6409 := PrimCons(symget_1time, tmp6408)

	tmp6410 := PrimCons(MakeNumber(3), tmp6409)

	tmp6411 := PrimCons(symget, tmp6410)

	tmp6412 := PrimCons(MakeNumber(1), tmp6411)

	tmp6413 := PrimCons(symgensym, tmp6412)

	tmp6414 := PrimCons(MakeNumber(1), tmp6413)

	tmp6415 := PrimCons(symfunction, tmp6414)

	tmp6416 := PrimCons(MakeNumber(1), tmp6415)

	tmp6417 := PrimCons(symfn, tmp6416)

	tmp6418 := PrimCons(MakeNumber(1), tmp6417)

	tmp6419 := PrimCons(symfst, tmp6418)

	tmp6420 := PrimCons(MakeNumber(0), tmp6419)

	tmp6421 := PrimCons(symfresh, tmp6420)

	tmp6422 := PrimCons(MakeNumber(1), tmp6421)

	tmp6423 := PrimCons(symfreeze, tmp6422)

	tmp6424 := PrimCons(MakeNumber(5), tmp6423)

	tmp6425 := PrimCons(symfork, tmp6424)

	tmp6426 := PrimCons(MakeNumber(7), tmp6425)

	tmp6427 := PrimCons(symfindall, tmp6426)

	tmp6428 := PrimCons(MakeNumber(2), tmp6427)

	tmp6429 := PrimCons(symfix, tmp6428)

	tmp6430 := PrimCons(MakeNumber(0), tmp6429)

	tmp6431 := PrimCons(symfail, tmp6430)

	tmp6432 := PrimCons(MakeNumber(2), tmp6431)

	tmp6433 := PrimCons(symfail_1if, tmp6432)

	tmp6434 := PrimCons(MakeNumber(1), tmp6433)

	tmp6435 := PrimCons(symfactorise, tmp6434)

	tmp6436 := PrimCons(MakeNumber(1), tmp6435)

	tmp6437 := PrimCons(symexternal, tmp6436)

	tmp6438 := PrimCons(MakeNumber(1), tmp6437)

	tmp6439 := PrimCons(symexplode, tmp6438)

	tmp6440 := PrimCons(MakeNumber(1), tmp6439)

	tmp6441 := PrimCons(symeval_1kl, tmp6440)

	tmp6442 := PrimCons(MakeNumber(1), tmp6441)

	tmp6443 := PrimCons(symeval, tmp6442)

	tmp6444 := PrimCons(MakeNumber(2), tmp6443)

	tmp6445 := PrimCons(symshen_4interror, tmp6444)

	tmp6446 := PrimCons(MakeNumber(1), tmp6445)

	tmp6447 := PrimCons(symerror_1to_1string, tmp6446)

	tmp6448 := PrimCons(MakeNumber(1), tmp6447)

	tmp6449 := PrimCons(symexternal, tmp6448)

	tmp6450 := PrimCons(MakeNumber(1), tmp6449)

	tmp6451 := PrimCons(symenable_1type_1theory, tmp6450)

	tmp6452 := PrimCons(MakeNumber(1), tmp6451)

	tmp6453 := PrimCons(symempty_2, tmp6452)

	tmp6454 := PrimCons(MakeNumber(2), tmp6453)

	tmp6455 := PrimCons(symelement_2, tmp6454)

	tmp6456 := PrimCons(MakeNumber(2), tmp6455)

	tmp6457 := PrimCons(symdo, tmp6456)

	tmp6458 := PrimCons(MakeNumber(2), tmp6457)

	tmp6459 := PrimCons(symdifference, tmp6458)

	tmp6460 := PrimCons(MakeNumber(1), tmp6459)

	tmp6461 := PrimCons(symdestroy, tmp6460)

	tmp6462 := PrimCons(MakeNumber(2), tmp6461)

	tmp6463 := PrimCons(symdeclare, tmp6462)

	tmp6464 := PrimCons(MakeNumber(1), tmp6463)

	tmp6465 := PrimCons(symclose, tmp6464)

	tmp6466 := PrimCons(MakeNumber(2), tmp6465)

	tmp6467 := PrimCons(symcn, tmp6466)

	tmp6468 := PrimCons(MakeNumber(1), tmp6467)

	tmp6469 := PrimCons(symcons_2, tmp6468)

	tmp6470 := PrimCons(MakeNumber(2), tmp6469)

	tmp6471 := PrimCons(symcons, tmp6470)

	tmp6472 := PrimCons(MakeNumber(2), tmp6471)

	tmp6473 := PrimCons(symconcat, tmp6472)

	tmp6474 := PrimCons(MakeNumber(2), tmp6473)

	tmp6475 := PrimCons(symcompile, tmp6474)

	tmp6476 := PrimCons(MakeNumber(1), tmp6475)

	tmp6477 := PrimCons(symcd, tmp6476)

	tmp6478 := PrimCons(MakeNumber(5), tmp6477)

	tmp6479 := PrimCons(symcall, tmp6478)

	tmp6480 := PrimCons(MakeNumber(6), tmp6479)

	tmp6481 := PrimCons(symbind, tmp6480)

	tmp6482 := PrimCons(MakeNumber(1), tmp6481)

	tmp6483 := PrimCons(symbound_2, tmp6482)

	tmp6484 := PrimCons(MakeNumber(1), tmp6483)

	tmp6485 := PrimCons(symbootstrap, tmp6484)

	tmp6486 := PrimCons(MakeNumber(1), tmp6485)

	tmp6487 := PrimCons(symboolean_2, tmp6486)

	tmp6488 := PrimCons(MakeNumber(1), tmp6487)

	tmp6489 := PrimCons(symatom_2, tmp6488)

	tmp6490 := PrimCons(MakeNumber(2), tmp6489)

	tmp6491 := PrimCons(symassoc, tmp6490)

	tmp6492 := PrimCons(MakeNumber(1), tmp6491)

	tmp6493 := PrimCons(symarity, tmp6492)

	tmp6494 := PrimCons(MakeNumber(2), tmp6493)

	tmp6495 := PrimCons(symappend, tmp6494)

	tmp6496 := PrimCons(MakeNumber(2), tmp6495)

	tmp6497 := PrimCons(symand, tmp6496)

	tmp6498 := PrimCons(MakeNumber(2), tmp6497)

	tmp6499 := PrimCons(symadjoin, tmp6498)

	tmp6500 := PrimCons(MakeNumber(3), tmp6499)

	tmp6501 := PrimCons(symaddress_1_6, tmp6500)

	tmp6502 := PrimCons(MakeNumber(1), tmp6501)

	tmp6503 := PrimCons(symabsvector, tmp6502)

	tmp6504 := PrimCons(MakeNumber(1), tmp6503)

	tmp6505 := PrimCons(symabsvector_2, tmp6504)

	tmp6506 := PrimCons(MakeNumber(0), tmp6505)

	tmp6507 := PrimCons(symabort, tmp6506)

	tmp6508 := Call(__e, PrimNS2Value(symshen_4initialise_1arity_1table), tmp6507)

	_ = tmp6508

	tmp6509 := MakeNative(func(__e *ControlFlow) {
		V1682 := __e.Get(1)
		_ = V1682
		tmp6510 := MakeNative(func(__e *ControlFlow) {
			External := __e.Get(1)
			_ = External
			tmp6511 := MakeNative(func(__e *ControlFlow) {
				Place := __e.Get(1)
				_ = Place
				__e.Return(V1682)
				return
			}, 1)

			tmp6512 := Call(__e, PrimNS2Value(symadjoin), V1682, External)

			tmp6513 := PrimNS3Value(sym_dproperty_1vector_d)

			tmp6514 := Call(__e, PrimNS2Value(symput), symshen, symshen_4external_1symbols, tmp6512, tmp6513)

			__e.TailApply(tmp6511, tmp6514)
			return

		}, 1)

		tmp6515 := PrimNS3Value(sym_dproperty_1vector_d)

		tmp6516 := Call(__e, PrimNS2Value(symget), symshen, symshen_4external_1symbols, tmp6515)

		__e.TailApply(tmp6510, tmp6516)
		return

	}, 1)

	tmp6517 := Call(__e, PrimNS2Value(symdef), symsystemf, tmp6509)

	_ = tmp6517

	tmp6518 := MakeNative(func(__e *ControlFlow) {
		V1683 := __e.Get(1)
		_ = V1683
		V1684 := __e.Get(2)
		_ = V1684
		tmp6520 := Call(__e, PrimNS2Value(symelement_2), V1683, V1684)

		if True == tmp6520 {
			__e.Return(V1684)
			return
		} else {
			__e.Return(PrimCons(V1683, V1684))
			return
		}

	}, 2)

	tmp6521 := Call(__e, PrimNS2Value(symdef), symadjoin, tmp6518)

	_ = tmp6521

	tmp6522 := PrimIntern(MakeString(":"))

	tmp6523 := PrimIntern(MakeString(";"))

	tmp6524 := PrimIntern(MakeString(":="))

	tmp6525 := PrimIntern(MakeString(","))

	tmp6526 := Call(__e, PrimNS2Value(symvector), MakeNumber(0))

	tmp6527 := PrimIntern(MakeString("bar!"))

	tmp6528 := PrimCons(symabort, Nil)

	tmp6529 := PrimCons(symabsvector, tmp6528)

	tmp6530 := PrimCons(symabsvector_2, tmp6529)

	tmp6531 := PrimCons(symaddress_1_6, tmp6530)

	tmp6532 := PrimCons(sym_5_1address, tmp6531)

	tmp6533 := PrimCons(symadjoin, tmp6532)

	tmp6534 := PrimCons(symand, tmp6533)

	tmp6535 := PrimCons(symappend, tmp6534)

	tmp6536 := PrimCons(symarity, tmp6535)

	tmp6537 := PrimCons(symassoc, tmp6536)

	tmp6538 := PrimCons(symatom_2, tmp6537)

	tmp6539 := PrimCons(tmp6527, tmp6538)

	tmp6540 := PrimCons(symbootstrap, tmp6539)

	tmp6541 := PrimCons(symboolean, tmp6540)

	tmp6542 := PrimCons(symboolean_2, tmp6541)

	tmp6543 := PrimCons(symbound_2, tmp6542)

	tmp6544 := PrimCons(symbind, tmp6543)

	tmp6545 := PrimCons(symclose, tmp6544)

	tmp6546 := PrimCons(symcall, tmp6545)

	tmp6547 := PrimCons(symcases, tmp6546)

	tmp6548 := PrimCons(symcd, tmp6547)

	tmp6549 := PrimCons(symcompile, tmp6548)

	tmp6550 := PrimCons(symconcat, tmp6549)

	tmp6551 := PrimCons(symcond, tmp6550)

	tmp6552 := PrimCons(symcons, tmp6551)

	tmp6553 := PrimCons(symcons_2, tmp6552)

	tmp6554 := PrimCons(symcn, tmp6553)

	tmp6555 := PrimCons(symdatatype, tmp6554)

	tmp6556 := PrimCons(symdeclare, tmp6555)

	tmp6557 := PrimCons(symdefprolog, tmp6556)

	tmp6558 := PrimCons(symdefcc, tmp6557)

	tmp6559 := PrimCons(symdefmacro, tmp6558)

	tmp6560 := PrimCons(symdefine, tmp6559)

	tmp6561 := PrimCons(symdefun, tmp6560)

	tmp6562 := PrimCons(symdestroy, tmp6561)

	tmp6563 := PrimCons(symdifference, tmp6562)

	tmp6564 := PrimCons(symdo, tmp6563)

	tmp6565 := PrimCons(symelement_2, tmp6564)

	tmp6566 := PrimCons(symempty_2, tmp6565)

	tmp6567 := PrimCons(symerror, tmp6566)

	tmp6568 := PrimCons(symerror_1to_1string, tmp6567)

	tmp6569 := PrimCons(symeval, tmp6568)

	tmp6570 := PrimCons(symeval_1kl, tmp6569)

	tmp6571 := PrimCons(symexception, tmp6570)

	tmp6572 := PrimCons(symexternal, tmp6571)

	tmp6573 := PrimCons(symexplode, tmp6572)

	tmp6574 := PrimCons(symenable_1type_1theory, tmp6573)

	tmp6575 := PrimCons(False, tmp6574)

	tmp6576 := PrimCons(symfindall, tmp6575)

	tmp6577 := PrimCons(symfactorise, tmp6576)

	tmp6578 := PrimCons(symfail_1if, tmp6577)

	tmp6579 := PrimCons(symfail, tmp6578)

	tmp6580 := PrimCons(symfile, tmp6579)

	tmp6581 := PrimCons(symfix, tmp6580)

	tmp6582 := PrimCons(symfork, tmp6581)

	tmp6583 := PrimCons(symfresh, tmp6582)

	tmp6584 := PrimCons(symfreeze, tmp6583)

	tmp6585 := PrimCons(symfst, tmp6584)

	tmp6586 := PrimCons(symfunction, tmp6585)

	tmp6587 := PrimCons(symfn, tmp6586)

	tmp6588 := PrimCons(symgensym, tmp6587)

	tmp6589 := PrimCons(symget_1time, tmp6588)

	tmp6590 := PrimCons(symget, tmp6589)

	tmp6591 := PrimCons(symhash, tmp6590)

	tmp6592 := PrimCons(symhdstr, tmp6591)

	tmp6593 := PrimCons(symhdv, tmp6592)

	tmp6594 := PrimCons(symhd, tmp6593)

	tmp6595 := PrimCons(symhead, tmp6594)

	tmp6596 := PrimCons(symif, tmp6595)

	tmp6597 := PrimCons(symimplementation, tmp6596)

	tmp6598 := PrimCons(syminternal, tmp6597)

	tmp6599 := PrimCons(symin_1package, tmp6598)

	tmp6600 := PrimCons(symin, tmp6599)

	tmp6601 := PrimCons(symis_b, tmp6600)

	tmp6602 := PrimCons(symis, tmp6601)

	tmp6603 := PrimCons(symit, tmp6602)

	tmp6604 := PrimCons(syminclude_1all_1but, tmp6603)

	tmp6605 := PrimCons(syminclude, tmp6604)

	tmp6606 := PrimCons(syminline, tmp6605)

	tmp6607 := PrimCons(syminput_7, tmp6606)

	tmp6608 := PrimCons(syminput, tmp6607)

	tmp6609 := PrimCons(syminteger_2, tmp6608)

	tmp6610 := PrimCons(symintern, tmp6609)

	tmp6611 := PrimCons(syminferences, tmp6610)

	tmp6612 := PrimCons(symintersection, tmp6611)

	tmp6613 := PrimCons(symis, tmp6612)

	tmp6614 := PrimCons(symlanguage, tmp6613)

	tmp6615 := PrimCons(symlambda, tmp6614)

	tmp6616 := PrimCons(symlazy, tmp6615)

	tmp6617 := PrimCons(symlet, tmp6616)

	tmp6618 := PrimCons(symlength, tmp6617)

	tmp6619 := PrimCons(symlimit, tmp6618)

	tmp6620 := PrimCons(symlineread, tmp6619)

	tmp6621 := PrimCons(symlist, tmp6620)

	tmp6622 := PrimCons(symloaded, tmp6621)

	tmp6623 := PrimCons(symload, tmp6622)

	tmp6624 := PrimCons(symmake_1string, tmp6623)

	tmp6625 := PrimCons(symmap, tmp6624)

	tmp6626 := PrimCons(symmapcan, tmp6625)

	tmp6627 := PrimCons(symmaxinferences, tmp6626)

	tmp6628 := PrimCons(symmacroexpand, tmp6627)

	tmp6629 := PrimCons(symmode, tmp6628)

	tmp6630 := PrimCons(symnl, tmp6629)

	tmp6631 := PrimCons(symnot, tmp6630)

	tmp6632 := PrimCons(symnth, tmp6631)

	tmp6633 := PrimCons(symnull, tmp6632)

	tmp6634 := PrimCons(symnumber, tmp6633)

	tmp6635 := PrimCons(symnumber_2, tmp6634)

	tmp6636 := PrimCons(symn_1_6string, tmp6635)

	tmp6637 := PrimCons(symoccurs_1check, tmp6636)

	tmp6638 := PrimCons(symoccurrences, tmp6637)

	tmp6639 := PrimCons(symopen, tmp6638)

	tmp6640 := PrimCons(symoptimise, tmp6639)

	tmp6641 := PrimCons(symor, tmp6640)

	tmp6642 := PrimCons(symos, tmp6641)

	tmp6643 := PrimCons(symout, tmp6642)

	tmp6644 := PrimCons(symoutput, tmp6643)

	tmp6645 := PrimCons(sympackage, tmp6644)

	tmp6646 := PrimCons(symport, tmp6645)

	tmp6647 := PrimCons(symporters, tmp6646)

	tmp6648 := PrimCons(sympos, tmp6647)

	tmp6649 := PrimCons(sympr, tmp6648)

	tmp6650 := PrimCons(symshen_4pprint, tmp6649)

	tmp6651 := PrimCons(symprint, tmp6650)

	tmp6652 := PrimCons(symprolog_1memory, tmp6651)

	tmp6653 := PrimCons(symprofile, tmp6652)

	tmp6654 := PrimCons(symprofile_1results, tmp6653)

	tmp6655 := PrimCons(symprotect, tmp6654)

	tmp6656 := PrimCons(symprolog_2, tmp6655)

	tmp6657 := PrimCons(symps, tmp6656)

	tmp6658 := PrimCons(sympreclude_1all_1but, tmp6657)

	tmp6659 := PrimCons(sympreclude, tmp6658)

	tmp6660 := PrimCons(symput, tmp6659)

	tmp6661 := PrimCons(sympackage_2, tmp6660)

	tmp6662 := PrimCons(symread_1from_1string_1unprocessed, tmp6661)

	tmp6663 := PrimCons(symread_1from_1string, tmp6662)

	tmp6664 := PrimCons(symread_1byte, tmp6663)

	tmp6665 := PrimCons(symread_1file_1as_1string, tmp6664)

	tmp6666 := PrimCons(symread_1file_1as_1bytelist, tmp6665)

	tmp6667 := PrimCons(symread_1file, tmp6666)

	tmp6668 := PrimCons(symreceive, tmp6667)

	tmp6669 := PrimCons(symread, tmp6668)

	tmp6670 := PrimCons(symrelease, tmp6669)

	tmp6671 := PrimCons(symremove, tmp6670)

	tmp6672 := PrimCons(symreverse, tmp6671)

	tmp6673 := PrimCons(symrun, tmp6672)

	tmp6674 := PrimCons(symstr, tmp6673)

	tmp6675 := PrimCons(symsave, tmp6674)

	tmp6676 := PrimCons(symset, tmp6675)

	tmp6677 := PrimCons(symsimple_1error, tmp6676)

	tmp6678 := PrimCons(symsnd, tmp6677)

	tmp6679 := PrimCons(symspecialise, tmp6678)

	tmp6680 := PrimCons(symspy, tmp6679)

	tmp6681 := PrimCons(symstep, tmp6680)

	tmp6682 := PrimCons(symstoutput, tmp6681)

	tmp6683 := PrimCons(symstinput, tmp6682)

	tmp6684 := PrimCons(symstring, tmp6683)

	tmp6685 := PrimCons(symstream, tmp6684)

	tmp6686 := PrimCons(symstring_1_6n, tmp6685)

	tmp6687 := PrimCons(symstring_2, tmp6686)

	tmp6688 := PrimCons(symsubst, tmp6687)

	tmp6689 := PrimCons(symsum, tmp6688)

	tmp6690 := PrimCons(symstring_1_6symbol, tmp6689)

	tmp6691 := PrimCons(symsymbol_2, tmp6690)

	tmp6692 := PrimCons(symsymbol, tmp6691)

	tmp6693 := PrimCons(symsynonyms, tmp6692)

	tmp6694 := PrimCons(symsystemf, tmp6693)

	tmp6695 := PrimCons(symtail, tmp6694)

	tmp6696 := PrimCons(symtlv, tmp6695)

	tmp6697 := PrimCons(symtlstr, tmp6696)

	tmp6698 := PrimCons(symtl, tmp6697)

	tmp6699 := PrimCons(symtc, tmp6698)

	tmp6700 := PrimCons(symtc_2, tmp6699)

	tmp6701 := PrimCons(symthaw, tmp6700)

	tmp6702 := PrimCons(symtime, tmp6701)

	tmp6703 := PrimCons(symtrack, tmp6702)

	tmp6704 := PrimCons(symtrap_1error, tmp6703)

	tmp6705 := PrimCons(True, tmp6704)

	tmp6706 := PrimCons(symtuple_2, tmp6705)

	tmp6707 := PrimCons(symtype, tmp6706)

	tmp6708 := PrimCons(symreturn, tmp6707)

	tmp6709 := PrimCons(symundefmacro, tmp6708)

	tmp6710 := PrimCons(symunprofile, tmp6709)

	tmp6711 := PrimCons(symunput, tmp6710)

	tmp6712 := PrimCons(symunion, tmp6711)

	tmp6713 := PrimCons(symshen_4unix, tmp6712)

	tmp6714 := PrimCons(symunit, tmp6713)

	tmp6715 := PrimCons(symuntrack, tmp6714)

	tmp6716 := PrimCons(symunspecialise, tmp6715)

	tmp6717 := PrimCons(symupdate_1lambda_1table, tmp6716)

	tmp6718 := PrimCons(symu_b, tmp6717)

	tmp6719 := PrimCons(symvector_2, tmp6718)

	tmp6720 := PrimCons(symvector, tmp6719)

	tmp6721 := PrimCons(sym_5_1vector, tmp6720)

	tmp6722 := PrimCons(symvector_1_6, tmp6721)

	tmp6723 := PrimCons(symvalue, tmp6722)

	tmp6724 := PrimCons(symvar_2, tmp6723)

	tmp6725 := PrimCons(symvariable_2, tmp6724)

	tmp6726 := PrimCons(symverified, tmp6725)

	tmp6727 := PrimCons(symversion, tmp6726)

	tmp6728 := PrimCons(symwarn, tmp6727)

	tmp6729 := PrimCons(symwhen, tmp6728)

	tmp6730 := PrimCons(symwhere, tmp6729)

	tmp6731 := PrimCons(symwrite_1byte, tmp6730)

	tmp6732 := PrimCons(symwrite_1to_1file, tmp6731)

	tmp6733 := PrimCons(symy_1or_1n_2, tmp6732)

	tmp6734 := PrimCons(tmp6526, tmp6733)

	tmp6735 := PrimCons(sym_6_6, tmp6734)

	tmp6736 := PrimCons(sym_5, tmp6735)

	tmp6737 := PrimCons(sym_5_a, tmp6736)

	tmp6738 := PrimCons(sym_7, tmp6737)

	tmp6739 := PrimCons(sym_d, tmp6738)

	tmp6740 := PrimCons(sym_c, tmp6739)

	tmp6741 := PrimCons(sym_1, tmp6740)

	tmp6742 := PrimCons(sym_3, tmp6741)

	tmp6743 := PrimCons(symshen_4_5end_6, tmp6742)

	tmp6744 := PrimCons(sym_5_b_6, tmp6743)

	tmp6745 := PrimCons(sym_c_4, tmp6744)

	tmp6746 := PrimCons(sym_a_a_6, tmp6745)

	tmp6747 := PrimCons(sym_6, tmp6746)

	tmp6748 := PrimCons(sym_6_a, tmp6747)

	tmp6749 := PrimCons(sym_a, tmp6748)

	tmp6750 := PrimCons(sym_a_a, tmp6749)

	tmp6751 := PrimCons(sym_5e_6, tmp6750)

	tmp6752 := PrimCons(sym_1_6, tmp6751)

	tmp6753 := PrimCons(sym_5_1, tmp6752)

	tmp6754 := PrimCons(sym_dhush_d, tmp6753)

	tmp6755 := PrimCons(sym_dporters_d, tmp6754)

	tmp6756 := PrimCons(sym_dport_d, tmp6755)

	tmp6757 := PrimCons(sym_8s, tmp6756)

	tmp6758 := PrimCons(sym_8p, tmp6757)

	tmp6759 := PrimCons(sym_8v, tmp6758)

	tmp6760 := PrimCons(sym_dproperty_1vector_d, tmp6759)

	tmp6761 := PrimCons(sym_drelease_d, tmp6760)

	tmp6762 := PrimCons(sym_dos_d, tmp6761)

	tmp6763 := PrimCons(sym_dmacros_d, tmp6762)

	tmp6764 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp6763)

	tmp6765 := PrimCons(sym_dversion_d, tmp6764)

	tmp6766 := PrimCons(sym_dhome_1directory_d, tmp6765)

	tmp6767 := PrimCons(sym_dstoutput_d, tmp6766)

	tmp6768 := PrimCons(sym_dstinput_d, tmp6767)

	tmp6769 := PrimCons(sym_dimplementation_d, tmp6768)

	tmp6770 := PrimCons(sym_dlanguage_d, tmp6769)

	tmp6771 := PrimCons(sym__, tmp6770)

	tmp6772 := PrimCons(tmp6525, tmp6771)

	tmp6773 := PrimCons(tmp6524, tmp6772)

	tmp6774 := PrimCons(tmp6523, tmp6773)

	tmp6775 := PrimCons(tmp6522, tmp6774)

	tmp6776 := PrimCons(sym_e_e, tmp6775)

	tmp6777 := PrimCons(sym_5_1_1, tmp6776)

	tmp6778 := PrimCons(sym_1_1_6, tmp6777)

	tmp6779 := PrimCons(sym_i, tmp6778)

	tmp6780 := PrimCons(sym_j, tmp6779)

	tmp6781 := PrimCons(sym_b, tmp6780)

	tmp6782 := PrimNS3Value(sym_dproperty_1vector_d)

	tmp6783 := Call(__e, PrimNS2Value(symput), symshen, symshen_4external_1symbols, tmp6781, tmp6782)

	_ = tmp6783

	tmp6784 := MakeNative(func(__e *ControlFlow) {
		V1685 := __e.Get(1)
		_ = V1685
		tmp6785 := MakeNative(func(__e *ControlFlow) {
			ArityF := __e.Get(1)
			_ = ArityF
			tmp6793 := PrimEqual(ArityF, MakeNumber(-1))

			var ifres6790 Obj

			if True == tmp6793 {
				ifres6790 = True

			} else {
				tmp6792 := PrimEqual(ArityF, MakeNumber(0))

				var ifres6791 Obj

				if True == tmp6792 {
					ifres6791 = True

				} else {
					ifres6791 = False

				}

				ifres6790 = ifres6791

			}

			if True == ifres6790 {
				__e.Return(Nil)
				return
			} else {
				tmp6787 := PrimCons(V1685, Nil)

				tmp6788 := Call(__e, PrimNS2Value(symshen_4lambda_1function), tmp6787, ArityF)

				tmp6789 := Call(__e, PrimNS2Value(symeval_1kl), tmp6788)

				__e.Return(PrimCons(V1685, tmp6789))
				return

			}

		}, 1)

		tmp6794 := Call(__e, PrimNS2Value(symarity), V1685)

		__e.TailApply(tmp6785, tmp6794)
		return

	}, 1)

	tmp6795 := Call(__e, PrimNS2Value(symdef), symshen_4lambda_1entry, tmp6784)

	_ = tmp6795

	tmp6796 := MakeNative(func(__e *ControlFlow) {
		V1686 := __e.Get(1)
		_ = V1686
		tmp6797 := MakeNative(func(__e *ControlFlow) {
			LambdaEntries := __e.Get(1)
			_ = LambdaEntries
			tmp6798 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4tuple), X)
				return
			}, 1)

			tmp6799 := PrimCons(symshen_4tuple, tmp6798)

			tmp6800 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4pvar), X)
				return
			}, 1)

			tmp6801 := PrimCons(symshen_4pvar, tmp6800)

			tmp6802 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4print_1prolog_1vector), X)
				return
			}, 1)

			tmp6803 := PrimCons(symshen_4print_1prolog_1vector, tmp6802)

			tmp6804 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4print_1freshterm), X)
				return
			}, 1)

			tmp6805 := PrimCons(symshen_4print_1freshterm, tmp6804)

			tmp6806 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4printF), X)
				return
			}, 1)

			tmp6807 := PrimCons(symshen_4printF, tmp6806)

			tmp6808 := PrimCons(tmp6807, LambdaEntries)

			tmp6809 := PrimCons(tmp6805, tmp6808)

			tmp6810 := PrimCons(tmp6803, tmp6809)

			tmp6811 := PrimCons(tmp6801, tmp6810)

			tmp6812 := PrimCons(tmp6799, tmp6811)

			__e.Return(PrimNS3Set(symshen_4_dlambdatable_d, tmp6812))
			return

		}, 1)

		tmp6813 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4lambda_1entry), X)
			return
		}, 1)

		tmp6814 := Call(__e, PrimNS2Value(symmap), tmp6813, V1686)

		__e.TailApply(tmp6797, tmp6814)
		return

	}, 1)

	tmp6815 := Call(__e, PrimNS2Value(symdef), symshen_4build_1lambda_1table, tmp6796)

	_ = tmp6815

	tmp6816 := Call(__e, PrimNS2Value(symexternal), symshen)

	__e.TailApply(PrimNS2Value(symshen_4build_1lambda_1table), tmp6816)
	return

}, 0)
