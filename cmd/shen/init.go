package main

import . "github.com/tiancaiamao/shen-go/cora"

var InitMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp4963 := MakeNative(func(__e *ControlFlow) {
		tmp4964 := PrimNS3Set(symshen_4_dinstalling_1kl_d, False)

		_ = tmp4964

		tmp4965 := PrimNS3Set(symshen_4_dhistory_d, Nil)

		_ = tmp4965

		tmp4966 := PrimNS3Set(symshen_4_dtc_d, False)

		_ = tmp4966

		tmp4967 := Call(__e, PrimNS2Value(symshen_4dict), MakeNumber(20000))

		tmp4968 := PrimNS3Set(sym_dproperty_1vector_d, tmp4967)

		_ = tmp4968

		tmp4969 := PrimNS3Set(symshen_4_dprocess_1counter_d, MakeNumber(0))

		_ = tmp4969

		tmp4970 := Call(__e, PrimNS2Value(symvector), MakeNumber(10000))

		tmp4971 := PrimNS3Set(symshen_4_dvarcounter_d, tmp4970)

		_ = tmp4971

		tmp4972 := Call(__e, PrimNS2Value(symvector), MakeNumber(10000))

		tmp4973 := PrimNS3Set(symshen_4_dprologvectors_d, tmp4972)

		_ = tmp4973

		tmp4974 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(X)
			return
		}, 1)

		tmp4975 := PrimNS3Set(symshen_4_ddemodulation_1function_d, tmp4974)

		_ = tmp4975

		tmp4976 := MakeNative(func(__e *ControlFlow) {
			Arg := __e.Get(1)
			_ = Arg
			__e.Return(MakeNative(func(__e *ControlFlow) {
				OnFail := __e.Get(1)
				_ = OnFail
				__e.TailApply(PrimNS2Value(symthaw), OnFail)
				return
			}, 1))
			return
		}, 1)

		tmp4977 := PrimNS3Set(symshen_4_dcustom_1pattern_1compiler_d, tmp4976)

		_ = tmp4977

		tmp4978 := MakeNative(func(__e *ControlFlow) {
			Arg := __e.Get(1)
			_ = Arg
			__e.Return(Arg)
			return
		}, 1)

		tmp4979 := PrimNS3Set(symshen_4_dcustom_1pattern_1reducer_d, tmp4978)

		_ = tmp4979

		tmp4980 := PrimCons(symshen_4function_1macro, Nil)

		tmp4981 := PrimCons(symshen_4defprolog_1macro, tmp4980)

		tmp4982 := PrimCons(symshen_4_8s_1macro, tmp4981)

		tmp4983 := PrimCons(symshen_4nl_1macro, tmp4982)

		tmp4984 := PrimCons(symshen_4synonyms_1macro, tmp4983)

		tmp4985 := PrimCons(symshen_4prolog_1macro, tmp4984)

		tmp4986 := PrimCons(symshen_4error_1macro, tmp4985)

		tmp4987 := PrimCons(symshen_4input_1macro, tmp4986)

		tmp4988 := PrimCons(symshen_4output_1macro, tmp4987)

		tmp4989 := PrimCons(symshen_4make_1string_1macro, tmp4988)

		tmp4990 := PrimCons(symshen_4assoc_1macro, tmp4989)

		tmp4991 := PrimCons(symshen_4let_1macro, tmp4990)

		tmp4992 := PrimCons(symshen_4datatype_1macro, tmp4991)

		tmp4993 := PrimCons(symshen_4compile_1macro, tmp4992)

		tmp4994 := PrimCons(symshen_4put_cget_1macro, tmp4993)

		tmp4995 := PrimCons(symshen_4abs_1macro, tmp4994)

		tmp4996 := PrimCons(symshen_4cases_1macro, tmp4995)

		tmp4997 := PrimCons(symshen_4timer_1macro, tmp4996)

		tmp4998 := PrimNS3Set(symshen_4_dmacroreg_d, tmp4997)

		_ = tmp4998

		tmp4999 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4timer_1macro), X)
			return
		}, 1)

		tmp5000 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4cases_1macro), X)
			return
		}, 1)

		tmp5001 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4abs_1macro), X)
			return
		}, 1)

		tmp5002 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4put_cget_1macro), X)
			return
		}, 1)

		tmp5003 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4compile_1macro), X)
			return
		}, 1)

		tmp5004 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4datatype_1macro), X)
			return
		}, 1)

		tmp5005 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4let_1macro), X)
			return
		}, 1)

		tmp5006 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4assoc_1macro), X)
			return
		}, 1)

		tmp5007 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4make_1string_1macro), X)
			return
		}, 1)

		tmp5008 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4output_1macro), X)
			return
		}, 1)

		tmp5009 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4input_1macro), X)
			return
		}, 1)

		tmp5010 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4error_1macro), X)
			return
		}, 1)

		tmp5011 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4prolog_1macro), X)
			return
		}, 1)

		tmp5012 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4synonyms_1macro), X)
			return
		}, 1)

		tmp5013 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4nl_1macro), X)
			return
		}, 1)

		tmp5014 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_8s_1macro), X)
			return
		}, 1)

		tmp5015 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4defprolog_1macro), X)
			return
		}, 1)

		tmp5016 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4function_1macro), X)
			return
		}, 1)

		tmp5017 := PrimCons(tmp5016, Nil)

		tmp5018 := PrimCons(tmp5015, tmp5017)

		tmp5019 := PrimCons(tmp5014, tmp5018)

		tmp5020 := PrimCons(tmp5013, tmp5019)

		tmp5021 := PrimCons(tmp5012, tmp5020)

		tmp5022 := PrimCons(tmp5011, tmp5021)

		tmp5023 := PrimCons(tmp5010, tmp5022)

		tmp5024 := PrimCons(tmp5009, tmp5023)

		tmp5025 := PrimCons(tmp5008, tmp5024)

		tmp5026 := PrimCons(tmp5007, tmp5025)

		tmp5027 := PrimCons(tmp5006, tmp5026)

		tmp5028 := PrimCons(tmp5005, tmp5027)

		tmp5029 := PrimCons(tmp5004, tmp5028)

		tmp5030 := PrimCons(tmp5003, tmp5029)

		tmp5031 := PrimCons(tmp5002, tmp5030)

		tmp5032 := PrimCons(tmp5001, tmp5031)

		tmp5033 := PrimCons(tmp5000, tmp5032)

		tmp5034 := PrimCons(tmp4999, tmp5033)

		tmp5035 := PrimNS3Set(sym_dmacros_d, tmp5034)

		_ = tmp5035

		tmp5036 := PrimNS3Set(symshen_4_dgensym_d, MakeNumber(0))

		_ = tmp5036

		tmp5037 := PrimNS3Set(symshen_4_dtracking_d, Nil)

		_ = tmp5037

		tmp5038 := PrimCons(symZ, Nil)

		tmp5039 := PrimCons(symY, tmp5038)

		tmp5040 := PrimCons(symX, tmp5039)

		tmp5041 := PrimCons(symW, tmp5040)

		tmp5042 := PrimCons(symV, tmp5041)

		tmp5043 := PrimCons(symU, tmp5042)

		tmp5044 := PrimCons(symT, tmp5043)

		tmp5045 := PrimCons(symS, tmp5044)

		tmp5046 := PrimCons(symR, tmp5045)

		tmp5047 := PrimCons(symQ, tmp5046)

		tmp5048 := PrimCons(symP, tmp5047)

		tmp5049 := PrimCons(symO, tmp5048)

		tmp5050 := PrimCons(symN, tmp5049)

		tmp5051 := PrimCons(symM, tmp5050)

		tmp5052 := PrimCons(symL, tmp5051)

		tmp5053 := PrimCons(symK, tmp5052)

		tmp5054 := PrimCons(symJ, tmp5053)

		tmp5055 := PrimCons(symI, tmp5054)

		tmp5056 := PrimCons(symH, tmp5055)

		tmp5057 := PrimCons(symG, tmp5056)

		tmp5058 := PrimCons(symF, tmp5057)

		tmp5059 := PrimCons(symE, tmp5058)

		tmp5060 := PrimCons(symD, tmp5059)

		tmp5061 := PrimCons(symC, tmp5060)

		tmp5062 := PrimCons(symB, tmp5061)

		tmp5063 := PrimCons(symA, tmp5062)

		tmp5064 := PrimNS3Set(symshen_4_dalphabet_d, tmp5063)

		_ = tmp5064

		tmp5065 := PrimCons(symopen, Nil)

		tmp5066 := PrimCons(symset, tmp5065)

		tmp5067 := PrimCons(symwhere, tmp5066)

		tmp5068 := PrimCons(symlet, tmp5067)

		tmp5069 := PrimCons(symlambda, tmp5068)

		tmp5070 := PrimCons(symcons, tmp5069)

		tmp5071 := PrimCons(sym_8v, tmp5070)

		tmp5072 := PrimCons(sym_8s, tmp5071)

		tmp5073 := PrimCons(sym_8p, tmp5072)

		tmp5074 := PrimNS3Set(symshen_4_dspecial_d, tmp5073)

		_ = tmp5074

		tmp5075 := PrimCons(symdefmacro, Nil)

		tmp5076 := PrimCons(symshen_4read_7, tmp5075)

		tmp5077 := PrimCons(symdefcc, tmp5076)

		tmp5078 := PrimCons(syminput_7, tmp5077)

		tmp5079 := PrimCons(symshen_4process_1datatype, tmp5078)

		tmp5080 := PrimCons(symdefine, tmp5079)

		tmp5081 := PrimNS3Set(symshen_4_dextraspecial_d, tmp5080)

		_ = tmp5081

		tmp5082 := PrimNS3Set(symshen_4_dspy_d, False)

		_ = tmp5082

		tmp5083 := PrimNS3Set(symshen_4_ddatatypes_d, Nil)

		_ = tmp5083

		tmp5084 := PrimNS3Set(symshen_4_dalldatatypes_d, Nil)

		_ = tmp5084

		tmp5085 := PrimNS3Set(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

		_ = tmp5085

		tmp5086 := PrimNS3Set(symshen_4_dsynonyms_d, Nil)

		_ = tmp5086

		tmp5087 := PrimNS3Set(symshen_4_dsystem_d, Nil)

		_ = tmp5087

		tmp5088 := PrimNS3Set(symshen_4_dmaxcomplexity_d, MakeNumber(128))

		_ = tmp5088

		tmp5089 := PrimNS3Set(symshen_4_doccurs_d, True)

		_ = tmp5089

		tmp5090 := PrimNS3Set(symshen_4_dmaxinferences_d, MakeNumber(1000000))

		_ = tmp5090

		tmp5091 := PrimNS3Set(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

		_ = tmp5091

		tmp5092 := PrimNS3Set(symshen_4_dcatch_d, MakeNumber(0))

		_ = tmp5092

		tmp5093 := PrimNS3Set(symshen_4_dcall_d, MakeNumber(0))

		_ = tmp5093

		tmp5094 := PrimNS3Set(symshen_4_dinfs_d, MakeNumber(0))

		_ = tmp5094

		tmp5095 := PrimNS3Set(sym_dhush_d, False)

		_ = tmp5095

		tmp5096 := PrimNS3Set(symshen_4_doptimise_d, False)

		_ = tmp5096

		tmp5097 := PrimNS3Set(sym_dversion_d, MakeString("Shen 22.3"))

		_ = tmp5097

		tmp5100 := Call(__e, PrimNS2Value(symbound_2), sym_dhome_1directory_d)

		tmp5101 := PrimNot(tmp5100)

		var ifres5098 Obj

		if True == tmp5101 {
			tmp5099 := PrimNS3Set(sym_dhome_1directory_d, MakeString(""))

			ifres5098 = tmp5099

		} else {
			ifres5098 = symshen_4skip

		}

		_ = ifres5098

		tmp5105 := Call(__e, PrimNS2Value(symbound_2), sym_dsterror_d)

		tmp5106 := PrimNot(tmp5105)

		var ifres5102 Obj

		if True == tmp5106 {
			tmp5103 := PrimNS3Value(sym_dstoutput_d)

			tmp5104 := PrimNS3Set(sym_dsterror_d, tmp5103)

			ifres5102 = tmp5104

		} else {
			ifres5102 = symshen_4skip

		}

		_ = ifres5102

		tmp5107 := PrimCons(MakeNumber(1), Nil)

		tmp5108 := PrimCons(syminclude_1all_1but, tmp5107)

		tmp5109 := PrimCons(MakeNumber(1), tmp5108)

		tmp5110 := PrimCons(sympreclude_1all_1but, tmp5109)

		tmp5111 := PrimCons(MakeNumber(1), tmp5110)

		tmp5112 := PrimCons(syminclude, tmp5111)

		tmp5113 := PrimCons(MakeNumber(1), tmp5112)

		tmp5114 := PrimCons(sympreclude, tmp5113)

		tmp5115 := PrimCons(MakeNumber(2), tmp5114)

		tmp5116 := PrimCons(sym_8s, tmp5115)

		tmp5117 := PrimCons(MakeNumber(2), tmp5116)

		tmp5118 := PrimCons(sym_8v, tmp5117)

		tmp5119 := PrimCons(MakeNumber(2), tmp5118)

		tmp5120 := PrimCons(sym_8p, tmp5119)

		tmp5121 := PrimCons(MakeNumber(1), tmp5120)

		tmp5122 := PrimCons(sym_5_b_6, tmp5121)

		tmp5123 := PrimCons(MakeNumber(1), tmp5122)

		tmp5124 := PrimCons(sym_5e_6, tmp5123)

		tmp5125 := PrimCons(MakeNumber(2), tmp5124)

		tmp5126 := PrimCons(sym_a_a, tmp5125)

		tmp5127 := PrimCons(MakeNumber(2), tmp5126)

		tmp5128 := PrimCons(sym_1, tmp5127)

		tmp5129 := PrimCons(MakeNumber(2), tmp5128)

		tmp5130 := PrimCons(sym_c, tmp5129)

		tmp5131 := PrimCons(MakeNumber(2), tmp5130)

		tmp5132 := PrimCons(sym_d, tmp5131)

		tmp5133 := PrimCons(MakeNumber(2), tmp5132)

		tmp5134 := PrimCons(sym_7, tmp5133)

		tmp5135 := PrimCons(MakeNumber(1), tmp5134)

		tmp5136 := PrimCons(symy_1or_1n_2, tmp5135)

		tmp5137 := PrimCons(MakeNumber(2), tmp5136)

		tmp5138 := PrimCons(symwrite_1to_1file, tmp5137)

		tmp5139 := PrimCons(MakeNumber(2), tmp5138)

		tmp5140 := PrimCons(symwrite_1byte, tmp5139)

		tmp5141 := PrimCons(MakeNumber(0), tmp5140)

		tmp5142 := PrimCons(symversion, tmp5141)

		tmp5143 := PrimCons(MakeNumber(1), tmp5142)

		tmp5144 := PrimCons(symvariable_2, tmp5143)

		tmp5145 := PrimCons(MakeNumber(1), tmp5144)

		tmp5146 := PrimCons(symvalue, tmp5145)

		tmp5147 := PrimCons(MakeNumber(3), tmp5146)

		tmp5148 := PrimCons(symvector_1_6, tmp5147)

		tmp5149 := PrimCons(MakeNumber(1), tmp5148)

		tmp5150 := PrimCons(symvector_2, tmp5149)

		tmp5151 := PrimCons(MakeNumber(1), tmp5150)

		tmp5152 := PrimCons(symvector, tmp5151)

		tmp5153 := PrimCons(MakeNumber(1), tmp5152)

		tmp5154 := PrimCons(symundefmacro, tmp5153)

		tmp5155 := PrimCons(MakeNumber(1), tmp5154)

		tmp5156 := PrimCons(symunspecialise, tmp5155)

		tmp5157 := PrimCons(MakeNumber(1), tmp5156)

		tmp5158 := PrimCons(symuntrack, tmp5157)

		tmp5159 := PrimCons(MakeNumber(2), tmp5158)

		tmp5160 := PrimCons(symunion, tmp5159)

		tmp5161 := PrimCons(MakeNumber(4), tmp5160)

		tmp5162 := PrimCons(symunify_b, tmp5161)

		tmp5163 := PrimCons(MakeNumber(4), tmp5162)

		tmp5164 := PrimCons(symunify, tmp5163)

		tmp5165 := PrimCons(MakeNumber(1), tmp5164)

		tmp5166 := PrimCons(symunprofile, tmp5165)

		tmp5167 := PrimCons(MakeNumber(3), tmp5166)

		tmp5168 := PrimCons(symunput, tmp5167)

		tmp5169 := PrimCons(MakeNumber(1), tmp5168)

		tmp5170 := PrimCons(symundefmacro, tmp5169)

		tmp5171 := PrimCons(MakeNumber(3), tmp5170)

		tmp5172 := PrimCons(symreturn, tmp5171)

		tmp5173 := PrimCons(MakeNumber(2), tmp5172)

		tmp5174 := PrimCons(symtype, tmp5173)

		tmp5175 := PrimCons(MakeNumber(1), tmp5174)

		tmp5176 := PrimCons(symtuple_2, tmp5175)

		tmp5177 := PrimCons(MakeNumber(2), tmp5176)

		tmp5178 := PrimCons(symtrap_1error, tmp5177)

		tmp5179 := PrimCons(MakeNumber(1), tmp5178)

		tmp5180 := PrimCons(symtrack, tmp5179)

		tmp5181 := PrimCons(MakeNumber(1), tmp5180)

		tmp5182 := PrimCons(symtlstr, tmp5181)

		tmp5183 := PrimCons(MakeNumber(1), tmp5182)

		tmp5184 := PrimCons(symthaw, tmp5183)

		tmp5185 := PrimCons(MakeNumber(0), tmp5184)

		tmp5186 := PrimCons(symtc_2, tmp5185)

		tmp5187 := PrimCons(MakeNumber(1), tmp5186)

		tmp5188 := PrimCons(symtc, tmp5187)

		tmp5189 := PrimCons(MakeNumber(1), tmp5188)

		tmp5190 := PrimCons(symtl, tmp5189)

		tmp5191 := PrimCons(MakeNumber(1), tmp5190)

		tmp5192 := PrimCons(symtail, tmp5191)

		tmp5193 := PrimCons(MakeNumber(1), tmp5192)

		tmp5194 := PrimCons(symsystemf, tmp5193)

		tmp5195 := PrimCons(MakeNumber(1), tmp5194)

		tmp5196 := PrimCons(symsymbol_2, tmp5195)

		tmp5197 := PrimCons(MakeNumber(1), tmp5196)

		tmp5198 := PrimCons(symsum, tmp5197)

		tmp5199 := PrimCons(MakeNumber(3), tmp5198)

		tmp5200 := PrimCons(symsubst, tmp5199)

		tmp5201 := PrimCons(MakeNumber(1), tmp5200)

		tmp5202 := PrimCons(symstr, tmp5201)

		tmp5203 := PrimCons(MakeNumber(1), tmp5202)

		tmp5204 := PrimCons(symstring_2, tmp5203)

		tmp5205 := PrimCons(MakeNumber(1), tmp5204)

		tmp5206 := PrimCons(symstring_1_6symbol, tmp5205)

		tmp5207 := PrimCons(MakeNumber(1), tmp5206)

		tmp5208 := PrimCons(symstring_1_6n, tmp5207)

		tmp5209 := PrimCons(MakeNumber(0), tmp5208)

		tmp5210 := PrimCons(symsterror, tmp5209)

		tmp5211 := PrimCons(MakeNumber(0), tmp5210)

		tmp5212 := PrimCons(symstoutput, tmp5211)

		tmp5213 := PrimCons(MakeNumber(0), tmp5212)

		tmp5214 := PrimCons(symstinput, tmp5213)

		tmp5215 := PrimCons(MakeNumber(1), tmp5214)

		tmp5216 := PrimCons(symstep, tmp5215)

		tmp5217 := PrimCons(MakeNumber(1), tmp5216)

		tmp5218 := PrimCons(symspy, tmp5217)

		tmp5219 := PrimCons(MakeNumber(1), tmp5218)

		tmp5220 := PrimCons(symspecialise, tmp5219)

		tmp5221 := PrimCons(MakeNumber(1), tmp5220)

		tmp5222 := PrimCons(symsnd, tmp5221)

		tmp5223 := PrimCons(MakeNumber(1), tmp5222)

		tmp5224 := PrimCons(symsimple_1error, tmp5223)

		tmp5225 := PrimCons(MakeNumber(2), tmp5224)

		tmp5226 := PrimCons(symset, tmp5225)

		tmp5227 := PrimCons(MakeNumber(1), tmp5226)

		tmp5228 := PrimCons(symreverse, tmp5227)

		tmp5229 := PrimCons(MakeNumber(3), tmp5228)

		tmp5230 := PrimCons(symshen_4require, tmp5229)

		tmp5231 := PrimCons(MakeNumber(2), tmp5230)

		tmp5232 := PrimCons(symremove, tmp5231)

		tmp5233 := PrimCons(MakeNumber(0), tmp5232)

		tmp5234 := PrimCons(symrelease, tmp5233)

		tmp5235 := PrimCons(MakeNumber(1), tmp5234)

		tmp5236 := PrimCons(symreceive, tmp5235)

		tmp5237 := PrimCons(MakeNumber(1), tmp5236)

		tmp5238 := PrimCons(symread_1from_1string, tmp5237)

		tmp5239 := PrimCons(MakeNumber(1), tmp5238)

		tmp5240 := PrimCons(symread_1byte, tmp5239)

		tmp5241 := PrimCons(MakeNumber(1), tmp5240)

		tmp5242 := PrimCons(symread, tmp5241)

		tmp5243 := PrimCons(MakeNumber(1), tmp5242)

		tmp5244 := PrimCons(symread_1file_1as_1bytelist, tmp5243)

		tmp5245 := PrimCons(MakeNumber(1), tmp5244)

		tmp5246 := PrimCons(symread_1file, tmp5245)

		tmp5247 := PrimCons(MakeNumber(1), tmp5246)

		tmp5248 := PrimCons(symread_1file_1as_1string, tmp5247)

		tmp5249 := PrimCons(MakeNumber(2), tmp5248)

		tmp5250 := PrimCons(symshen_4reassemble, tmp5249)

		tmp5251 := PrimCons(MakeNumber(4), tmp5250)

		tmp5252 := PrimCons(symput, tmp5251)

		tmp5253 := PrimCons(MakeNumber(3), tmp5252)

		tmp5254 := PrimCons(symaddress_1_6, tmp5253)

		tmp5255 := PrimCons(MakeNumber(1), tmp5254)

		tmp5256 := PrimCons(symprotect, tmp5255)

		tmp5257 := PrimCons(MakeNumber(1), tmp5256)

		tmp5258 := PrimCons(sympreclude_1all_1but, tmp5257)

		tmp5259 := PrimCons(MakeNumber(1), tmp5258)

		tmp5260 := PrimCons(sympreclude, tmp5259)

		tmp5261 := PrimCons(MakeNumber(1), tmp5260)

		tmp5262 := PrimCons(symps, tmp5261)

		tmp5263 := PrimCons(MakeNumber(2), tmp5262)

		tmp5264 := PrimCons(sympr, tmp5263)

		tmp5265 := PrimCons(MakeNumber(1), tmp5264)

		tmp5266 := PrimCons(symprofile_1results, tmp5265)

		tmp5267 := PrimCons(MakeNumber(1), tmp5266)

		tmp5268 := PrimCons(symprofile, tmp5267)

		tmp5269 := PrimCons(MakeNumber(1), tmp5268)

		tmp5270 := PrimCons(symprint, tmp5269)

		tmp5271 := PrimCons(MakeNumber(2), tmp5270)

		tmp5272 := PrimCons(sympos, tmp5271)

		tmp5273 := PrimCons(MakeNumber(0), tmp5272)

		tmp5274 := PrimCons(symporters, tmp5273)

		tmp5275 := PrimCons(MakeNumber(0), tmp5274)

		tmp5276 := PrimCons(symport, tmp5275)

		tmp5277 := PrimCons(MakeNumber(1), tmp5276)

		tmp5278 := PrimCons(sympackage_2, tmp5277)

		tmp5279 := PrimCons(MakeNumber(3), tmp5278)

		tmp5280 := PrimCons(sympackage, tmp5279)

		tmp5281 := PrimCons(MakeNumber(0), tmp5280)

		tmp5282 := PrimCons(symos, tmp5281)

		tmp5283 := PrimCons(MakeNumber(2), tmp5282)

		tmp5284 := PrimCons(symor, tmp5283)

		tmp5285 := PrimCons(MakeNumber(1), tmp5284)

		tmp5286 := PrimCons(symoptimise, tmp5285)

		tmp5287 := PrimCons(MakeNumber(2), tmp5286)

		tmp5288 := PrimCons(symopen, tmp5287)

		tmp5289 := PrimCons(MakeNumber(1), tmp5288)

		tmp5290 := PrimCons(symoccurs_1check, tmp5289)

		tmp5291 := PrimCons(MakeNumber(2), tmp5290)

		tmp5292 := PrimCons(symoccurrences, tmp5291)

		tmp5293 := PrimCons(MakeNumber(1), tmp5292)

		tmp5294 := PrimCons(symoccurs_1check, tmp5293)

		tmp5295 := PrimCons(MakeNumber(1), tmp5294)

		tmp5296 := PrimCons(symnumber_2, tmp5295)

		tmp5297 := PrimCons(MakeNumber(1), tmp5296)

		tmp5298 := PrimCons(symn_1_6string, tmp5297)

		tmp5299 := PrimCons(MakeNumber(2), tmp5298)

		tmp5300 := PrimCons(symnth, tmp5299)

		tmp5301 := PrimCons(MakeNumber(1), tmp5300)

		tmp5302 := PrimCons(symnot, tmp5301)

		tmp5303 := PrimCons(MakeNumber(1), tmp5302)

		tmp5304 := PrimCons(symnl, tmp5303)

		tmp5305 := PrimCons(MakeNumber(1), tmp5304)

		tmp5306 := PrimCons(symmaxinferences, tmp5305)

		tmp5307 := PrimCons(MakeNumber(2), tmp5306)

		tmp5308 := PrimCons(symmapcan, tmp5307)

		tmp5309 := PrimCons(MakeNumber(2), tmp5308)

		tmp5310 := PrimCons(symmap, tmp5309)

		tmp5311 := PrimCons(MakeNumber(1), tmp5310)

		tmp5312 := PrimCons(symmacroexpand, tmp5311)

		tmp5313 := PrimCons(MakeNumber(1), tmp5312)

		tmp5314 := PrimCons(symvector, tmp5313)

		tmp5315 := PrimCons(MakeNumber(2), tmp5314)

		tmp5316 := PrimCons(sym_5_a, tmp5315)

		tmp5317 := PrimCons(MakeNumber(2), tmp5316)

		tmp5318 := PrimCons(sym_5, tmp5317)

		tmp5319 := PrimCons(MakeNumber(1), tmp5318)

		tmp5320 := PrimCons(symload, tmp5319)

		tmp5321 := PrimCons(MakeNumber(1), tmp5320)

		tmp5322 := PrimCons(symlineread, tmp5321)

		tmp5323 := PrimCons(MakeNumber(1), tmp5322)

		tmp5324 := PrimCons(symlimit, tmp5323)

		tmp5325 := PrimCons(MakeNumber(1), tmp5324)

		tmp5326 := PrimCons(symlength, tmp5325)

		tmp5327 := PrimCons(MakeNumber(0), tmp5326)

		tmp5328 := PrimCons(symlanguage, tmp5327)

		tmp5329 := PrimCons(MakeNumber(0), tmp5328)

		tmp5330 := PrimCons(symkill, tmp5329)

		tmp5331 := PrimCons(MakeNumber(0), tmp5330)

		tmp5332 := PrimCons(symit, tmp5331)

		tmp5333 := PrimCons(MakeNumber(1), tmp5332)

		tmp5334 := PrimCons(syminternal, tmp5333)

		tmp5335 := PrimCons(MakeNumber(2), tmp5334)

		tmp5336 := PrimCons(symintersection, tmp5335)

		tmp5337 := PrimCons(MakeNumber(0), tmp5336)

		tmp5338 := PrimCons(symimplementation, tmp5337)

		tmp5339 := PrimCons(MakeNumber(2), tmp5338)

		tmp5340 := PrimCons(syminput_7, tmp5339)

		tmp5341 := PrimCons(MakeNumber(1), tmp5340)

		tmp5342 := PrimCons(syminput, tmp5341)

		tmp5343 := PrimCons(MakeNumber(0), tmp5342)

		tmp5344 := PrimCons(syminferences, tmp5343)

		tmp5345 := PrimCons(MakeNumber(4), tmp5344)

		tmp5346 := PrimCons(symidentical, tmp5345)

		tmp5347 := PrimCons(MakeNumber(1), tmp5346)

		tmp5348 := PrimCons(symintern, tmp5347)

		tmp5349 := PrimCons(MakeNumber(1), tmp5348)

		tmp5350 := PrimCons(syminteger_2, tmp5349)

		tmp5351 := PrimCons(MakeNumber(3), tmp5350)

		tmp5352 := PrimCons(symif, tmp5351)

		tmp5353 := PrimCons(MakeNumber(1), tmp5352)

		tmp5354 := PrimCons(symhead, tmp5353)

		tmp5355 := PrimCons(MakeNumber(1), tmp5354)

		tmp5356 := PrimCons(symhdstr, tmp5355)

		tmp5357 := PrimCons(MakeNumber(1), tmp5356)

		tmp5358 := PrimCons(symhdv, tmp5357)

		tmp5359 := PrimCons(MakeNumber(1), tmp5358)

		tmp5360 := PrimCons(symhd, tmp5359)

		tmp5361 := PrimCons(MakeNumber(2), tmp5360)

		tmp5362 := PrimCons(symhash, tmp5361)

		tmp5363 := PrimCons(MakeNumber(2), tmp5362)

		tmp5364 := PrimCons(sym_a, tmp5363)

		tmp5365 := PrimCons(MakeNumber(2), tmp5364)

		tmp5366 := PrimCons(sym_6_a, tmp5365)

		tmp5367 := PrimCons(MakeNumber(2), tmp5366)

		tmp5368 := PrimCons(sym_6, tmp5367)

		tmp5369 := PrimCons(MakeNumber(2), tmp5368)

		tmp5370 := PrimCons(sym_5_1vector, tmp5369)

		tmp5371 := PrimCons(MakeNumber(2), tmp5370)

		tmp5372 := PrimCons(sym_5_1address, tmp5371)

		tmp5373 := PrimCons(MakeNumber(3), tmp5372)

		tmp5374 := PrimCons(symaddress_1_6, tmp5373)

		tmp5375 := PrimCons(MakeNumber(1), tmp5374)

		tmp5376 := PrimCons(symget_1time, tmp5375)

		tmp5377 := PrimCons(MakeNumber(3), tmp5376)

		tmp5378 := PrimCons(symget, tmp5377)

		tmp5379 := PrimCons(MakeNumber(1), tmp5378)

		tmp5380 := PrimCons(symgensym, tmp5379)

		tmp5381 := PrimCons(MakeNumber(1), tmp5380)

		tmp5382 := PrimCons(symfst, tmp5381)

		tmp5383 := PrimCons(MakeNumber(1), tmp5382)

		tmp5384 := PrimCons(symfreeze, tmp5383)

		tmp5385 := PrimCons(MakeNumber(5), tmp5384)

		tmp5386 := PrimCons(symfindall, tmp5385)

		tmp5387 := PrimCons(MakeNumber(2), tmp5386)

		tmp5388 := PrimCons(symfix, tmp5387)

		tmp5389 := PrimCons(MakeNumber(0), tmp5388)

		tmp5390 := PrimCons(symfail, tmp5389)

		tmp5391 := PrimCons(MakeNumber(2), tmp5390)

		tmp5392 := PrimCons(symfail_1if, tmp5391)

		tmp5393 := PrimCons(MakeNumber(1), tmp5392)

		tmp5394 := PrimCons(symexternal, tmp5393)

		tmp5395 := PrimCons(MakeNumber(1), tmp5394)

		tmp5396 := PrimCons(symexplode, tmp5395)

		tmp5397 := PrimCons(MakeNumber(1), tmp5396)

		tmp5398 := PrimCons(symeval_1kl, tmp5397)

		tmp5399 := PrimCons(MakeNumber(1), tmp5398)

		tmp5400 := PrimCons(symeval, tmp5399)

		tmp5401 := PrimCons(MakeNumber(2), tmp5400)

		tmp5402 := PrimCons(symshen_4interror, tmp5401)

		tmp5403 := PrimCons(MakeNumber(1), tmp5402)

		tmp5404 := PrimCons(symerror_1to_1string, tmp5403)

		tmp5405 := PrimCons(MakeNumber(1), tmp5404)

		tmp5406 := PrimCons(symenable_1type_1theory, tmp5405)

		tmp5407 := PrimCons(MakeNumber(1), tmp5406)

		tmp5408 := PrimCons(symempty_2, tmp5407)

		tmp5409 := PrimCons(MakeNumber(2), tmp5408)

		tmp5410 := PrimCons(symelement_2, tmp5409)

		tmp5411 := PrimCons(MakeNumber(2), tmp5410)

		tmp5412 := PrimCons(symdo, tmp5411)

		tmp5413 := PrimCons(MakeNumber(2), tmp5412)

		tmp5414 := PrimCons(symdifference, tmp5413)

		tmp5415 := PrimCons(MakeNumber(1), tmp5414)

		tmp5416 := PrimCons(symdestroy, tmp5415)

		tmp5417 := PrimCons(MakeNumber(2), tmp5416)

		tmp5418 := PrimCons(symdeclare, tmp5417)

		tmp5419 := PrimCons(MakeNumber(2), tmp5418)

		tmp5420 := PrimCons(symcn, tmp5419)

		tmp5421 := PrimCons(MakeNumber(1), tmp5420)

		tmp5422 := PrimCons(symcons_2, tmp5421)

		tmp5423 := PrimCons(MakeNumber(2), tmp5422)

		tmp5424 := PrimCons(symcons, tmp5423)

		tmp5425 := PrimCons(MakeNumber(2), tmp5424)

		tmp5426 := PrimCons(symconcat, tmp5425)

		tmp5427 := PrimCons(MakeNumber(3), tmp5426)

		tmp5428 := PrimCons(symcompile, tmp5427)

		tmp5429 := PrimCons(MakeNumber(1), tmp5428)

		tmp5430 := PrimCons(symclose, tmp5429)

		tmp5431 := PrimCons(MakeNumber(1), tmp5430)

		tmp5432 := PrimCons(symcd, tmp5431)

		tmp5433 := PrimCons(MakeNumber(1), tmp5432)

		tmp5434 := PrimCons(symbound_2, tmp5433)

		tmp5435 := PrimCons(MakeNumber(1), tmp5434)

		tmp5436 := PrimCons(symboolean_2, tmp5435)

		tmp5437 := PrimCons(MakeNumber(2), tmp5436)

		tmp5438 := PrimCons(symassoc, tmp5437)

		tmp5439 := PrimCons(MakeNumber(1), tmp5438)

		tmp5440 := PrimCons(symarity, tmp5439)

		tmp5441 := PrimCons(MakeNumber(2), tmp5440)

		tmp5442 := PrimCons(symappend, tmp5441)

		tmp5443 := PrimCons(MakeNumber(2), tmp5442)

		tmp5444 := PrimCons(symand, tmp5443)

		tmp5445 := PrimCons(MakeNumber(2), tmp5444)

		tmp5446 := PrimCons(symadjoin, tmp5445)

		tmp5447 := PrimCons(MakeNumber(1), tmp5446)

		tmp5448 := PrimCons(symabsvector, tmp5447)

		tmp5449 := PrimCons(MakeNumber(1), tmp5448)

		tmp5450 := PrimCons(symabsvector_2, tmp5449)

		tmp5451 := PrimCons(MakeNumber(0), tmp5450)

		tmp5452 := PrimCons(symabort, tmp5451)

		tmp5453 := Call(__e, PrimNS2Value(symshen_4initialise__arity__table), tmp5452)

		_ = tmp5453

		tmp5454 := PrimIntern(MakeString("shen"))

		tmp5455 := PrimCons(symabsvector, Nil)

		tmp5456 := PrimCons(symabsvector_2, tmp5455)

		tmp5457 := PrimCons(symaddress_1_6, tmp5456)

		tmp5458 := PrimCons(sym_5_1address, tmp5457)

		tmp5459 := PrimCons(symadjoin, tmp5458)

		tmp5460 := PrimCons(symand, tmp5459)

		tmp5461 := PrimCons(symappend, tmp5460)

		tmp5462 := PrimCons(symabort, tmp5461)

		tmp5463 := PrimCons(symarity, tmp5462)

		tmp5464 := PrimCons(symassoc, tmp5463)

		tmp5465 := PrimCons(symbar_b, tmp5464)

		tmp5466 := PrimCons(symboolean, tmp5465)

		tmp5467 := PrimCons(symboolean_2, tmp5466)

		tmp5468 := PrimCons(symbound_2, tmp5467)

		tmp5469 := PrimCons(symbind, tmp5468)

		tmp5470 := PrimCons(symclose, tmp5469)

		tmp5471 := PrimCons(symcall, tmp5470)

		tmp5472 := PrimCons(symcases, tmp5471)

		tmp5473 := PrimCons(symcd, tmp5472)

		tmp5474 := PrimCons(symcompile, tmp5473)

		tmp5475 := PrimCons(symconcat, tmp5474)

		tmp5476 := PrimCons(symcond, tmp5475)

		tmp5477 := PrimCons(symcons, tmp5476)

		tmp5478 := PrimCons(symcons_2, tmp5477)

		tmp5479 := PrimCons(symcn, tmp5478)

		tmp5480 := PrimCons(symcut, tmp5479)

		tmp5481 := PrimCons(symdatatype, tmp5480)

		tmp5482 := PrimCons(symdeclare, tmp5481)

		tmp5483 := PrimCons(symdefprolog, tmp5482)

		tmp5484 := PrimCons(symdefcc, tmp5483)

		tmp5485 := PrimCons(symdefmacro, tmp5484)

		tmp5486 := PrimCons(symdefine, tmp5485)

		tmp5487 := PrimCons(symdefun, tmp5486)

		tmp5488 := PrimCons(symdestroy, tmp5487)

		tmp5489 := PrimCons(symdifference, tmp5488)

		tmp5490 := PrimCons(symdo, tmp5489)

		tmp5491 := PrimCons(symelement_2, tmp5490)

		tmp5492 := PrimCons(symempty_2, tmp5491)

		tmp5493 := PrimCons(symerror, tmp5492)

		tmp5494 := PrimCons(symerror_1to_1string, tmp5493)

		tmp5495 := PrimCons(symeval, tmp5494)

		tmp5496 := PrimCons(symeval_1kl, tmp5495)

		tmp5497 := PrimCons(symexception, tmp5496)

		tmp5498 := PrimCons(symexternal, tmp5497)

		tmp5499 := PrimCons(symexplode, tmp5498)

		tmp5500 := PrimCons(symenable_1type_1theory, tmp5499)

		tmp5501 := PrimCons(False, tmp5500)

		tmp5502 := PrimCons(symfindall, tmp5501)

		tmp5503 := PrimCons(symfwhen, tmp5502)

		tmp5504 := PrimCons(symfail_1if, tmp5503)

		tmp5505 := PrimCons(symfail, tmp5504)

		tmp5506 := PrimCons(symfile, tmp5505)

		tmp5507 := PrimCons(symfix, tmp5506)

		tmp5508 := PrimCons(symfreeze, tmp5507)

		tmp5509 := PrimCons(symfst, tmp5508)

		tmp5510 := PrimCons(symfunction, tmp5509)

		tmp5511 := PrimCons(symgensym, tmp5510)

		tmp5512 := PrimCons(symget_1time, tmp5511)

		tmp5513 := PrimCons(symget, tmp5512)

		tmp5514 := PrimCons(symhash, tmp5513)

		tmp5515 := PrimCons(symhdstr, tmp5514)

		tmp5516 := PrimCons(symhdv, tmp5515)

		tmp5517 := PrimCons(symhd, tmp5516)

		tmp5518 := PrimCons(symhead, tmp5517)

		tmp5519 := PrimCons(symidentical, tmp5518)

		tmp5520 := PrimCons(symif, tmp5519)

		tmp5521 := PrimCons(symimplementation, tmp5520)

		tmp5522 := PrimCons(syminternal, tmp5521)

		tmp5523 := PrimCons(symin, tmp5522)

		tmp5524 := PrimCons(symit, tmp5523)

		tmp5525 := PrimCons(syminclude_1all_1but, tmp5524)

		tmp5526 := PrimCons(syminclude, tmp5525)

		tmp5527 := PrimCons(syminput_7, tmp5526)

		tmp5528 := PrimCons(syminput, tmp5527)

		tmp5529 := PrimCons(syminteger_2, tmp5528)

		tmp5530 := PrimCons(symintern, tmp5529)

		tmp5531 := PrimCons(syminferences, tmp5530)

		tmp5532 := PrimCons(symintersection, tmp5531)

		tmp5533 := PrimCons(symis, tmp5532)

		tmp5534 := PrimCons(symkill, tmp5533)

		tmp5535 := PrimCons(symlanguage, tmp5534)

		tmp5536 := PrimCons(symlambda, tmp5535)

		tmp5537 := PrimCons(symlazy, tmp5536)

		tmp5538 := PrimCons(symlet, tmp5537)

		tmp5539 := PrimCons(symlength, tmp5538)

		tmp5540 := PrimCons(symlimit, tmp5539)

		tmp5541 := PrimCons(symlineread, tmp5540)

		tmp5542 := PrimCons(symlist, tmp5541)

		tmp5543 := PrimCons(symloaded, tmp5542)

		tmp5544 := PrimCons(symload, tmp5543)

		tmp5545 := PrimCons(symmake_1string, tmp5544)

		tmp5546 := PrimCons(symmap, tmp5545)

		tmp5547 := PrimCons(symmapcan, tmp5546)

		tmp5548 := PrimCons(symmaxinferences, tmp5547)

		tmp5549 := PrimCons(symmacroexpand, tmp5548)

		tmp5550 := PrimCons(symmode, tmp5549)

		tmp5551 := PrimCons(symnl, tmp5550)

		tmp5552 := PrimCons(symnot, tmp5551)

		tmp5553 := PrimCons(symnth, tmp5552)

		tmp5554 := PrimCons(symnull, tmp5553)

		tmp5555 := PrimCons(symnumber, tmp5554)

		tmp5556 := PrimCons(symnumber_2, tmp5555)

		tmp5557 := PrimCons(symn_1_6string, tmp5556)

		tmp5558 := PrimCons(symoccurs_1check, tmp5557)

		tmp5559 := PrimCons(symoccurrences, tmp5558)

		tmp5560 := PrimCons(symopen, tmp5559)

		tmp5561 := PrimCons(symoptimise, tmp5560)

		tmp5562 := PrimCons(symor, tmp5561)

		tmp5563 := PrimCons(symos, tmp5562)

		tmp5564 := PrimCons(symout, tmp5563)

		tmp5565 := PrimCons(symoutput, tmp5564)

		tmp5566 := PrimCons(sympackage, tmp5565)

		tmp5567 := PrimCons(symport, tmp5566)

		tmp5568 := PrimCons(symporters, tmp5567)

		tmp5569 := PrimCons(sympos, tmp5568)

		tmp5570 := PrimCons(sympr, tmp5569)

		tmp5571 := PrimCons(symprint, tmp5570)

		tmp5572 := PrimCons(symprofile, tmp5571)

		tmp5573 := PrimCons(symprofile_1results, tmp5572)

		tmp5574 := PrimCons(symprotect, tmp5573)

		tmp5575 := PrimCons(symprolog_2, tmp5574)

		tmp5576 := PrimCons(symps, tmp5575)

		tmp5577 := PrimCons(sympreclude_1all_1but, tmp5576)

		tmp5578 := PrimCons(sympreclude, tmp5577)

		tmp5579 := PrimCons(symput, tmp5578)

		tmp5580 := PrimCons(sympackage_2, tmp5579)

		tmp5581 := PrimCons(symread_1from_1string, tmp5580)

		tmp5582 := PrimCons(symread_1byte, tmp5581)

		tmp5583 := PrimCons(symread_1file_1as_1string, tmp5582)

		tmp5584 := PrimCons(symread_1file_1as_1bytelist, tmp5583)

		tmp5585 := PrimCons(symread_1file, tmp5584)

		tmp5586 := PrimCons(symreceive, tmp5585)

		tmp5587 := PrimCons(symread, tmp5586)

		tmp5588 := PrimCons(symrelease, tmp5587)

		tmp5589 := PrimCons(symremove, tmp5588)

		tmp5590 := PrimCons(symreverse, tmp5589)

		tmp5591 := PrimCons(symrun, tmp5590)

		tmp5592 := PrimCons(symstr, tmp5591)

		tmp5593 := PrimCons(symsave, tmp5592)

		tmp5594 := PrimCons(symset, tmp5593)

		tmp5595 := PrimCons(symsimple_1error, tmp5594)

		tmp5596 := PrimCons(symsnd, tmp5595)

		tmp5597 := PrimCons(symspecialise, tmp5596)

		tmp5598 := PrimCons(symspy, tmp5597)

		tmp5599 := PrimCons(symstep, tmp5598)

		tmp5600 := PrimCons(symstoutput, tmp5599)

		tmp5601 := PrimCons(symsterror, tmp5600)

		tmp5602 := PrimCons(symstinput, tmp5601)

		tmp5603 := PrimCons(symstring, tmp5602)

		tmp5604 := PrimCons(symstream, tmp5603)

		tmp5605 := PrimCons(symstring_1_6n, tmp5604)

		tmp5606 := PrimCons(symstring_2, tmp5605)

		tmp5607 := PrimCons(symsubst, tmp5606)

		tmp5608 := PrimCons(symsum, tmp5607)

		tmp5609 := PrimCons(symstring_1_6symbol, tmp5608)

		tmp5610 := PrimCons(symsymbol_2, tmp5609)

		tmp5611 := PrimCons(symsymbol, tmp5610)

		tmp5612 := PrimCons(symsynonyms, tmp5611)

		tmp5613 := PrimCons(symsystemf, tmp5612)

		tmp5614 := PrimCons(symtail, tmp5613)

		tmp5615 := PrimCons(symtlv, tmp5614)

		tmp5616 := PrimCons(symtlstr, tmp5615)

		tmp5617 := PrimCons(symtl, tmp5616)

		tmp5618 := PrimCons(symtc, tmp5617)

		tmp5619 := PrimCons(symtc_2, tmp5618)

		tmp5620 := PrimCons(symthaw, tmp5619)

		tmp5621 := PrimCons(symtime, tmp5620)

		tmp5622 := PrimCons(symtrack, tmp5621)

		tmp5623 := PrimCons(symtrap_1error, tmp5622)

		tmp5624 := PrimCons(True, tmp5623)

		tmp5625 := PrimCons(symtuple_2, tmp5624)

		tmp5626 := PrimCons(symtype, tmp5625)

		tmp5627 := PrimCons(symreturn, tmp5626)

		tmp5628 := PrimCons(symundefmacro, tmp5627)

		tmp5629 := PrimCons(symunprofile, tmp5628)

		tmp5630 := PrimCons(symunput, tmp5629)

		tmp5631 := PrimCons(symunify_b, tmp5630)

		tmp5632 := PrimCons(symunify, tmp5631)

		tmp5633 := PrimCons(symunion, tmp5632)

		tmp5634 := PrimCons(symshen_4unix, tmp5633)

		tmp5635 := PrimCons(symunit, tmp5634)

		tmp5636 := PrimCons(symuntrack, tmp5635)

		tmp5637 := PrimCons(symunspecialise, tmp5636)

		tmp5638 := PrimCons(symvector_2, tmp5637)

		tmp5639 := PrimCons(symvector, tmp5638)

		tmp5640 := PrimCons(sym_5_1vector, tmp5639)

		tmp5641 := PrimCons(symvector_1_6, tmp5640)

		tmp5642 := PrimCons(symvalue, tmp5641)

		tmp5643 := PrimCons(symvariable_2, tmp5642)

		tmp5644 := PrimCons(symverified, tmp5643)

		tmp5645 := PrimCons(symversion, tmp5644)

		tmp5646 := PrimCons(symwarn, tmp5645)

		tmp5647 := PrimCons(symwhen, tmp5646)

		tmp5648 := PrimCons(symwhere, tmp5647)

		tmp5649 := PrimCons(symwrite_1byte, tmp5648)

		tmp5650 := PrimCons(symwrite_1to_1file, tmp5649)

		tmp5651 := PrimCons(symy_1or_1n_2, tmp5650)

		tmp5652 := PrimCons(sym_6_6, tmp5651)

		tmp5653 := PrimCons(sym_5, tmp5652)

		tmp5654 := PrimCons(sym_5_a, tmp5653)

		tmp5655 := PrimCons(sym_7, tmp5654)

		tmp5656 := PrimCons(sym_d, tmp5655)

		tmp5657 := PrimCons(sym_c, tmp5656)

		tmp5658 := PrimCons(sym_1, tmp5657)

		tmp5659 := PrimCons(sym_3, tmp5658)

		tmp5660 := PrimCons(sym_a_b, tmp5659)

		tmp5661 := PrimCons(sym_c_4, tmp5660)

		tmp5662 := PrimCons(sym_6, tmp5661)

		tmp5663 := PrimCons(sym_6_a, tmp5662)

		tmp5664 := PrimCons(sym_a, tmp5663)

		tmp5665 := PrimCons(sym_a_a, tmp5664)

		tmp5666 := PrimCons(sym_5_b_6, tmp5665)

		tmp5667 := PrimCons(sym_5e_6, tmp5666)

		tmp5668 := PrimCons(sym_1_6, tmp5667)

		tmp5669 := PrimCons(sym_5_1, tmp5668)

		tmp5670 := PrimCons(sym_8s, tmp5669)

		tmp5671 := PrimCons(sym_8p, tmp5670)

		tmp5672 := PrimCons(sym_8v, tmp5671)

		tmp5673 := PrimCons(sym_dhush_d, tmp5672)

		tmp5674 := PrimCons(sym_dporters_d, tmp5673)

		tmp5675 := PrimCons(sym_dport_d, tmp5674)

		tmp5676 := PrimCons(sym_dproperty_1vector_d, tmp5675)

		tmp5677 := PrimCons(sym_drelease_d, tmp5676)

		tmp5678 := PrimCons(sym_dos_d, tmp5677)

		tmp5679 := PrimCons(sym_dmacros_d, tmp5678)

		tmp5680 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp5679)

		tmp5681 := PrimCons(sym_dversion_d, tmp5680)

		tmp5682 := PrimCons(sym_dhome_1directory_d, tmp5681)

		tmp5683 := PrimCons(sym_dsterror_d, tmp5682)

		tmp5684 := PrimCons(sym_dstoutput_d, tmp5683)

		tmp5685 := PrimCons(sym_dstinput_d, tmp5684)

		tmp5686 := PrimCons(sym_dimplementation_d, tmp5685)

		tmp5687 := PrimCons(sym_dlanguage_d, tmp5686)

		tmp5688 := PrimCons(sym_l, tmp5687)

		tmp5689 := PrimCons(sym__, tmp5688)

		tmp5690 := PrimCons(sym_h_a, tmp5689)

		tmp5691 := PrimCons(sym_h_1, tmp5690)

		tmp5692 := PrimCons(sym_k, tmp5691)

		tmp5693 := PrimCons(sym_h, tmp5692)

		tmp5694 := PrimCons(sym_e_e, tmp5693)

		tmp5695 := PrimCons(sym_5_1_1, tmp5694)

		tmp5696 := PrimCons(sym_1_1_6, tmp5695)

		tmp5697 := PrimCons(sym_i, tmp5696)

		tmp5698 := PrimCons(sym_j, tmp5697)

		tmp5699 := PrimCons(sym_b, tmp5698)

		tmp5700 := PrimNS3Value(sym_dproperty_1vector_d)

		tmp5701 := Call(__e, PrimNS2Value(symput), tmp5454, symshen_4external_1symbols, tmp5699, tmp5700)

		_ = tmp5701

		tmp5702 := PrimNS3Set(symshen_4_dhistory_d, Nil)

		_ = tmp5702

		tmp5703 := PrimNS3Set(symshen_4_dstep_d, False)

		_ = tmp5703

		tmp5704 := PrimAbsvector(MakeNumber(0))

		__e.Return(PrimNS3Set(symshen_4_dempty_1absvector_d, tmp5704))
		return

	}, 0)

	tmp5705 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise_1environment, tmp4963)

	_ = tmp5705

	tmp5706 := MakeNative(func(__e *ControlFlow) {
		tmp5707 := PrimNS3Set(symshen_4_dsignedfuncs_d, Nil)

		_ = tmp5707

		tmp5708 := PrimCons(symboolean, Nil)

		tmp5709 := PrimCons(sym_1_1_6, tmp5708)

		tmp5710 := PrimCons(symA, tmp5709)

		tmp5711 := PrimCons(symabsvector_2, tmp5710)

		tmp5712 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5713 := PrimCons(tmp5711, tmp5712)

		tmp5714 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5713)

		_ = tmp5714

		tmp5715 := PrimCons(symA, Nil)

		tmp5716 := PrimCons(symlist, tmp5715)

		tmp5717 := PrimCons(symA, Nil)

		tmp5718 := PrimCons(symlist, tmp5717)

		tmp5719 := PrimCons(tmp5718, Nil)

		tmp5720 := PrimCons(sym_1_1_6, tmp5719)

		tmp5721 := PrimCons(tmp5716, tmp5720)

		tmp5722 := PrimCons(tmp5721, Nil)

		tmp5723 := PrimCons(sym_1_1_6, tmp5722)

		tmp5724 := PrimCons(symA, tmp5723)

		tmp5725 := PrimCons(symadjoin, tmp5724)

		tmp5726 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5727 := PrimCons(tmp5725, tmp5726)

		tmp5728 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5727)

		_ = tmp5728

		tmp5729 := PrimCons(symboolean, Nil)

		tmp5730 := PrimCons(sym_1_1_6, tmp5729)

		tmp5731 := PrimCons(symboolean, tmp5730)

		tmp5732 := PrimCons(tmp5731, Nil)

		tmp5733 := PrimCons(sym_1_1_6, tmp5732)

		tmp5734 := PrimCons(symboolean, tmp5733)

		tmp5735 := PrimCons(symand, tmp5734)

		tmp5736 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5737 := PrimCons(tmp5735, tmp5736)

		tmp5738 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5737)

		_ = tmp5738

		tmp5739 := PrimCons(symstring, Nil)

		tmp5740 := PrimCons(sym_1_1_6, tmp5739)

		tmp5741 := PrimCons(symsymbol, tmp5740)

		tmp5742 := PrimCons(tmp5741, Nil)

		tmp5743 := PrimCons(sym_1_1_6, tmp5742)

		tmp5744 := PrimCons(symstring, tmp5743)

		tmp5745 := PrimCons(tmp5744, Nil)

		tmp5746 := PrimCons(sym_1_1_6, tmp5745)

		tmp5747 := PrimCons(symA, tmp5746)

		tmp5748 := PrimCons(symshen_4app, tmp5747)

		tmp5749 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5750 := PrimCons(tmp5748, tmp5749)

		tmp5751 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5750)

		_ = tmp5751

		tmp5752 := PrimCons(symA, Nil)

		tmp5753 := PrimCons(symlist, tmp5752)

		tmp5754 := PrimCons(symA, Nil)

		tmp5755 := PrimCons(symlist, tmp5754)

		tmp5756 := PrimCons(symA, Nil)

		tmp5757 := PrimCons(symlist, tmp5756)

		tmp5758 := PrimCons(tmp5757, Nil)

		tmp5759 := PrimCons(sym_1_1_6, tmp5758)

		tmp5760 := PrimCons(tmp5755, tmp5759)

		tmp5761 := PrimCons(tmp5760, Nil)

		tmp5762 := PrimCons(sym_1_1_6, tmp5761)

		tmp5763 := PrimCons(tmp5753, tmp5762)

		tmp5764 := PrimCons(symappend, tmp5763)

		tmp5765 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5766 := PrimCons(tmp5764, tmp5765)

		tmp5767 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5766)

		_ = tmp5767

		tmp5768 := PrimCons(symnumber, Nil)

		tmp5769 := PrimCons(sym_1_1_6, tmp5768)

		tmp5770 := PrimCons(symA, tmp5769)

		tmp5771 := PrimCons(symarity, tmp5770)

		tmp5772 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5773 := PrimCons(tmp5771, tmp5772)

		tmp5774 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5773)

		_ = tmp5774

		tmp5775 := PrimCons(symA, Nil)

		tmp5776 := PrimCons(symlist, tmp5775)

		tmp5777 := PrimCons(tmp5776, Nil)

		tmp5778 := PrimCons(symlist, tmp5777)

		tmp5779 := PrimCons(symA, Nil)

		tmp5780 := PrimCons(symlist, tmp5779)

		tmp5781 := PrimCons(tmp5780, Nil)

		tmp5782 := PrimCons(sym_1_1_6, tmp5781)

		tmp5783 := PrimCons(tmp5778, tmp5782)

		tmp5784 := PrimCons(tmp5783, Nil)

		tmp5785 := PrimCons(sym_1_1_6, tmp5784)

		tmp5786 := PrimCons(symA, tmp5785)

		tmp5787 := PrimCons(symassoc, tmp5786)

		tmp5788 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5789 := PrimCons(tmp5787, tmp5788)

		tmp5790 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5789)

		_ = tmp5790

		tmp5791 := PrimCons(symboolean, Nil)

		tmp5792 := PrimCons(sym_1_1_6, tmp5791)

		tmp5793 := PrimCons(symA, tmp5792)

		tmp5794 := PrimCons(symboolean_2, tmp5793)

		tmp5795 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5796 := PrimCons(tmp5794, tmp5795)

		tmp5797 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5796)

		_ = tmp5797

		tmp5798 := PrimCons(symboolean, Nil)

		tmp5799 := PrimCons(sym_1_1_6, tmp5798)

		tmp5800 := PrimCons(symsymbol, tmp5799)

		tmp5801 := PrimCons(symbound_2, tmp5800)

		tmp5802 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5803 := PrimCons(tmp5801, tmp5802)

		tmp5804 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5803)

		_ = tmp5804

		tmp5805 := PrimCons(symstring, Nil)

		tmp5806 := PrimCons(sym_1_1_6, tmp5805)

		tmp5807 := PrimCons(symstring, tmp5806)

		tmp5808 := PrimCons(symcd, tmp5807)

		tmp5809 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5810 := PrimCons(tmp5808, tmp5809)

		tmp5811 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5810)

		_ = tmp5811

		tmp5812 := PrimCons(symA, Nil)

		tmp5813 := PrimCons(symstream, tmp5812)

		tmp5814 := PrimCons(symB, Nil)

		tmp5815 := PrimCons(symlist, tmp5814)

		tmp5816 := PrimCons(tmp5815, Nil)

		tmp5817 := PrimCons(sym_1_1_6, tmp5816)

		tmp5818 := PrimCons(tmp5813, tmp5817)

		tmp5819 := PrimCons(symclose, tmp5818)

		tmp5820 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5821 := PrimCons(tmp5819, tmp5820)

		tmp5822 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5821)

		_ = tmp5822

		tmp5823 := PrimCons(symstring, Nil)

		tmp5824 := PrimCons(sym_1_1_6, tmp5823)

		tmp5825 := PrimCons(symstring, tmp5824)

		tmp5826 := PrimCons(tmp5825, Nil)

		tmp5827 := PrimCons(sym_1_1_6, tmp5826)

		tmp5828 := PrimCons(symstring, tmp5827)

		tmp5829 := PrimCons(symcn, tmp5828)

		tmp5830 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5831 := PrimCons(tmp5829, tmp5830)

		tmp5832 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5831)

		_ = tmp5832

		tmp5833 := PrimCons(symB, Nil)

		tmp5834 := PrimCons(symshen_4_a_a_6, tmp5833)

		tmp5835 := PrimCons(symA, tmp5834)

		tmp5836 := PrimCons(symB, Nil)

		tmp5837 := PrimCons(sym_1_1_6, tmp5836)

		tmp5838 := PrimCons(symA, tmp5837)

		tmp5839 := PrimCons(symB, Nil)

		tmp5840 := PrimCons(sym_1_1_6, tmp5839)

		tmp5841 := PrimCons(tmp5838, tmp5840)

		tmp5842 := PrimCons(tmp5841, Nil)

		tmp5843 := PrimCons(sym_1_1_6, tmp5842)

		tmp5844 := PrimCons(symA, tmp5843)

		tmp5845 := PrimCons(tmp5844, Nil)

		tmp5846 := PrimCons(sym_1_1_6, tmp5845)

		tmp5847 := PrimCons(tmp5835, tmp5846)

		tmp5848 := PrimCons(symcompile, tmp5847)

		tmp5849 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5850 := PrimCons(tmp5848, tmp5849)

		tmp5851 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5850)

		_ = tmp5851

		tmp5852 := PrimCons(symboolean, Nil)

		tmp5853 := PrimCons(sym_1_1_6, tmp5852)

		tmp5854 := PrimCons(symA, tmp5853)

		tmp5855 := PrimCons(symcons_2, tmp5854)

		tmp5856 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5857 := PrimCons(tmp5855, tmp5856)

		tmp5858 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5857)

		_ = tmp5858

		tmp5859 := PrimCons(symB, Nil)

		tmp5860 := PrimCons(sym_1_1_6, tmp5859)

		tmp5861 := PrimCons(symA, tmp5860)

		tmp5862 := PrimCons(symsymbol, Nil)

		tmp5863 := PrimCons(sym_1_1_6, tmp5862)

		tmp5864 := PrimCons(tmp5861, tmp5863)

		tmp5865 := PrimCons(symdestroy, tmp5864)

		tmp5866 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5867 := PrimCons(tmp5865, tmp5866)

		tmp5868 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5867)

		_ = tmp5868

		tmp5869 := PrimCons(symA, Nil)

		tmp5870 := PrimCons(symlist, tmp5869)

		tmp5871 := PrimCons(symA, Nil)

		tmp5872 := PrimCons(symlist, tmp5871)

		tmp5873 := PrimCons(symA, Nil)

		tmp5874 := PrimCons(symlist, tmp5873)

		tmp5875 := PrimCons(tmp5874, Nil)

		tmp5876 := PrimCons(sym_1_1_6, tmp5875)

		tmp5877 := PrimCons(tmp5872, tmp5876)

		tmp5878 := PrimCons(tmp5877, Nil)

		tmp5879 := PrimCons(sym_1_1_6, tmp5878)

		tmp5880 := PrimCons(tmp5870, tmp5879)

		tmp5881 := PrimCons(symdifference, tmp5880)

		tmp5882 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5883 := PrimCons(tmp5881, tmp5882)

		tmp5884 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5883)

		_ = tmp5884

		tmp5885 := PrimCons(symB, Nil)

		tmp5886 := PrimCons(sym_1_1_6, tmp5885)

		tmp5887 := PrimCons(symB, tmp5886)

		tmp5888 := PrimCons(tmp5887, Nil)

		tmp5889 := PrimCons(sym_1_1_6, tmp5888)

		tmp5890 := PrimCons(symA, tmp5889)

		tmp5891 := PrimCons(symdo, tmp5890)

		tmp5892 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5893 := PrimCons(tmp5891, tmp5892)

		tmp5894 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5893)

		_ = tmp5894

		tmp5895 := PrimCons(symA, Nil)

		tmp5896 := PrimCons(symlist, tmp5895)

		tmp5897 := PrimCons(symB, Nil)

		tmp5898 := PrimCons(symlist, tmp5897)

		tmp5899 := PrimCons(tmp5898, Nil)

		tmp5900 := PrimCons(symshen_4_a_a_6, tmp5899)

		tmp5901 := PrimCons(tmp5896, tmp5900)

		tmp5902 := PrimCons(sym_5e_6, tmp5901)

		tmp5903 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5904 := PrimCons(tmp5902, tmp5903)

		tmp5905 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5904)

		_ = tmp5905

		tmp5906 := PrimCons(symA, Nil)

		tmp5907 := PrimCons(symlist, tmp5906)

		tmp5908 := PrimCons(symA, Nil)

		tmp5909 := PrimCons(symlist, tmp5908)

		tmp5910 := PrimCons(tmp5909, Nil)

		tmp5911 := PrimCons(symshen_4_a_a_6, tmp5910)

		tmp5912 := PrimCons(tmp5907, tmp5911)

		tmp5913 := PrimCons(sym_5_b_6, tmp5912)

		tmp5914 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5915 := PrimCons(tmp5913, tmp5914)

		tmp5916 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5915)

		_ = tmp5916

		tmp5917 := PrimCons(symA, Nil)

		tmp5918 := PrimCons(symlist, tmp5917)

		tmp5919 := PrimCons(symboolean, Nil)

		tmp5920 := PrimCons(sym_1_1_6, tmp5919)

		tmp5921 := PrimCons(tmp5918, tmp5920)

		tmp5922 := PrimCons(tmp5921, Nil)

		tmp5923 := PrimCons(sym_1_1_6, tmp5922)

		tmp5924 := PrimCons(symA, tmp5923)

		tmp5925 := PrimCons(symelement_2, tmp5924)

		tmp5926 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5927 := PrimCons(tmp5925, tmp5926)

		tmp5928 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5927)

		_ = tmp5928

		tmp5929 := PrimCons(symboolean, Nil)

		tmp5930 := PrimCons(sym_1_1_6, tmp5929)

		tmp5931 := PrimCons(symA, tmp5930)

		tmp5932 := PrimCons(symempty_2, tmp5931)

		tmp5933 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5934 := PrimCons(tmp5932, tmp5933)

		tmp5935 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5934)

		_ = tmp5935

		tmp5936 := PrimCons(symboolean, Nil)

		tmp5937 := PrimCons(sym_1_1_6, tmp5936)

		tmp5938 := PrimCons(symsymbol, tmp5937)

		tmp5939 := PrimCons(symenable_1type_1theory, tmp5938)

		tmp5940 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5941 := PrimCons(tmp5939, tmp5940)

		tmp5942 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5941)

		_ = tmp5942

		tmp5943 := PrimCons(symsymbol, Nil)

		tmp5944 := PrimCons(symlist, tmp5943)

		tmp5945 := PrimCons(tmp5944, Nil)

		tmp5946 := PrimCons(sym_1_1_6, tmp5945)

		tmp5947 := PrimCons(symsymbol, tmp5946)

		tmp5948 := PrimCons(symexternal, tmp5947)

		tmp5949 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5950 := PrimCons(tmp5948, tmp5949)

		tmp5951 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5950)

		_ = tmp5951

		tmp5952 := PrimCons(symstring, Nil)

		tmp5953 := PrimCons(sym_1_1_6, tmp5952)

		tmp5954 := PrimCons(symexception, tmp5953)

		tmp5955 := PrimCons(symerror_1to_1string, tmp5954)

		tmp5956 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5957 := PrimCons(tmp5955, tmp5956)

		tmp5958 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5957)

		_ = tmp5958

		tmp5959 := PrimCons(symstring, Nil)

		tmp5960 := PrimCons(symlist, tmp5959)

		tmp5961 := PrimCons(tmp5960, Nil)

		tmp5962 := PrimCons(sym_1_1_6, tmp5961)

		tmp5963 := PrimCons(symA, tmp5962)

		tmp5964 := PrimCons(symexplode, tmp5963)

		tmp5965 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5966 := PrimCons(tmp5964, tmp5965)

		tmp5967 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5966)

		_ = tmp5967

		tmp5968 := PrimCons(symsymbol, Nil)

		tmp5969 := PrimCons(sym_1_1_6, tmp5968)

		tmp5970 := PrimCons(symfail, tmp5969)

		tmp5971 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5972 := PrimCons(tmp5970, tmp5971)

		tmp5973 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5972)

		_ = tmp5973

		tmp5974 := PrimCons(symboolean, Nil)

		tmp5975 := PrimCons(sym_1_1_6, tmp5974)

		tmp5976 := PrimCons(symsymbol, tmp5975)

		tmp5977 := PrimCons(symsymbol, Nil)

		tmp5978 := PrimCons(sym_1_1_6, tmp5977)

		tmp5979 := PrimCons(symsymbol, tmp5978)

		tmp5980 := PrimCons(tmp5979, Nil)

		tmp5981 := PrimCons(sym_1_1_6, tmp5980)

		tmp5982 := PrimCons(tmp5976, tmp5981)

		tmp5983 := PrimCons(symfail_1if, tmp5982)

		tmp5984 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5985 := PrimCons(tmp5983, tmp5984)

		tmp5986 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5985)

		_ = tmp5986

		tmp5987 := PrimCons(symA, Nil)

		tmp5988 := PrimCons(sym_1_1_6, tmp5987)

		tmp5989 := PrimCons(symA, tmp5988)

		tmp5990 := PrimCons(symA, Nil)

		tmp5991 := PrimCons(sym_1_1_6, tmp5990)

		tmp5992 := PrimCons(symA, tmp5991)

		tmp5993 := PrimCons(tmp5992, Nil)

		tmp5994 := PrimCons(sym_1_1_6, tmp5993)

		tmp5995 := PrimCons(tmp5989, tmp5994)

		tmp5996 := PrimCons(symfix, tmp5995)

		tmp5997 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp5998 := PrimCons(tmp5996, tmp5997)

		tmp5999 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp5998)

		_ = tmp5999

		tmp6000 := PrimCons(symA, Nil)

		tmp6001 := PrimCons(symlazy, tmp6000)

		tmp6002 := PrimCons(tmp6001, Nil)

		tmp6003 := PrimCons(sym_1_1_6, tmp6002)

		tmp6004 := PrimCons(symA, tmp6003)

		tmp6005 := PrimCons(symfreeze, tmp6004)

		tmp6006 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6007 := PrimCons(tmp6005, tmp6006)

		tmp6008 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6007)

		_ = tmp6008

		tmp6009 := PrimCons(symB, Nil)

		tmp6010 := PrimCons(sym_d, tmp6009)

		tmp6011 := PrimCons(symA, tmp6010)

		tmp6012 := PrimCons(symA, Nil)

		tmp6013 := PrimCons(sym_1_1_6, tmp6012)

		tmp6014 := PrimCons(tmp6011, tmp6013)

		tmp6015 := PrimCons(symfst, tmp6014)

		tmp6016 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6017 := PrimCons(tmp6015, tmp6016)

		tmp6018 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6017)

		_ = tmp6018

		tmp6019 := PrimCons(symB, Nil)

		tmp6020 := PrimCons(sym_1_1_6, tmp6019)

		tmp6021 := PrimCons(symA, tmp6020)

		tmp6022 := PrimCons(symB, Nil)

		tmp6023 := PrimCons(sym_1_1_6, tmp6022)

		tmp6024 := PrimCons(symA, tmp6023)

		tmp6025 := PrimCons(tmp6024, Nil)

		tmp6026 := PrimCons(sym_1_1_6, tmp6025)

		tmp6027 := PrimCons(tmp6021, tmp6026)

		tmp6028 := PrimCons(symfunction, tmp6027)

		tmp6029 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6030 := PrimCons(tmp6028, tmp6029)

		tmp6031 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6030)

		_ = tmp6031

		tmp6032 := PrimCons(symsymbol, Nil)

		tmp6033 := PrimCons(sym_1_1_6, tmp6032)

		tmp6034 := PrimCons(symsymbol, tmp6033)

		tmp6035 := PrimCons(symgensym, tmp6034)

		tmp6036 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6037 := PrimCons(tmp6035, tmp6036)

		tmp6038 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6037)

		_ = tmp6038

		tmp6039 := PrimCons(symA, Nil)

		tmp6040 := PrimCons(symvector, tmp6039)

		tmp6041 := PrimCons(symA, Nil)

		tmp6042 := PrimCons(sym_1_1_6, tmp6041)

		tmp6043 := PrimCons(symnumber, tmp6042)

		tmp6044 := PrimCons(tmp6043, Nil)

		tmp6045 := PrimCons(sym_1_1_6, tmp6044)

		tmp6046 := PrimCons(tmp6040, tmp6045)

		tmp6047 := PrimCons(sym_5_1vector, tmp6046)

		tmp6048 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6049 := PrimCons(tmp6047, tmp6048)

		tmp6050 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6049)

		_ = tmp6050

		tmp6051 := PrimCons(symA, Nil)

		tmp6052 := PrimCons(symvector, tmp6051)

		tmp6053 := PrimCons(symA, Nil)

		tmp6054 := PrimCons(symvector, tmp6053)

		tmp6055 := PrimCons(tmp6054, Nil)

		tmp6056 := PrimCons(sym_1_1_6, tmp6055)

		tmp6057 := PrimCons(symA, tmp6056)

		tmp6058 := PrimCons(tmp6057, Nil)

		tmp6059 := PrimCons(sym_1_1_6, tmp6058)

		tmp6060 := PrimCons(symnumber, tmp6059)

		tmp6061 := PrimCons(tmp6060, Nil)

		tmp6062 := PrimCons(sym_1_1_6, tmp6061)

		tmp6063 := PrimCons(tmp6052, tmp6062)

		tmp6064 := PrimCons(symvector_1_6, tmp6063)

		tmp6065 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6066 := PrimCons(tmp6064, tmp6065)

		tmp6067 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6066)

		_ = tmp6067

		tmp6068 := PrimCons(symA, Nil)

		tmp6069 := PrimCons(symvector, tmp6068)

		tmp6070 := PrimCons(tmp6069, Nil)

		tmp6071 := PrimCons(sym_1_1_6, tmp6070)

		tmp6072 := PrimCons(symnumber, tmp6071)

		tmp6073 := PrimCons(symvector, tmp6072)

		tmp6074 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6075 := PrimCons(tmp6073, tmp6074)

		tmp6076 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6075)

		_ = tmp6076

		tmp6077 := PrimCons(symnumber, Nil)

		tmp6078 := PrimCons(sym_1_1_6, tmp6077)

		tmp6079 := PrimCons(symsymbol, tmp6078)

		tmp6080 := PrimCons(symget_1time, tmp6079)

		tmp6081 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6082 := PrimCons(tmp6080, tmp6081)

		tmp6083 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6082)

		_ = tmp6083

		tmp6084 := PrimCons(symnumber, Nil)

		tmp6085 := PrimCons(sym_1_1_6, tmp6084)

		tmp6086 := PrimCons(symnumber, tmp6085)

		tmp6087 := PrimCons(tmp6086, Nil)

		tmp6088 := PrimCons(sym_1_1_6, tmp6087)

		tmp6089 := PrimCons(symA, tmp6088)

		tmp6090 := PrimCons(symhash, tmp6089)

		tmp6091 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6092 := PrimCons(tmp6090, tmp6091)

		tmp6093 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6092)

		_ = tmp6093

		tmp6094 := PrimCons(symA, Nil)

		tmp6095 := PrimCons(symlist, tmp6094)

		tmp6096 := PrimCons(symA, Nil)

		tmp6097 := PrimCons(sym_1_1_6, tmp6096)

		tmp6098 := PrimCons(tmp6095, tmp6097)

		tmp6099 := PrimCons(symhead, tmp6098)

		tmp6100 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6101 := PrimCons(tmp6099, tmp6100)

		tmp6102 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6101)

		_ = tmp6102

		tmp6103 := PrimCons(symA, Nil)

		tmp6104 := PrimCons(symvector, tmp6103)

		tmp6105 := PrimCons(symA, Nil)

		tmp6106 := PrimCons(sym_1_1_6, tmp6105)

		tmp6107 := PrimCons(tmp6104, tmp6106)

		tmp6108 := PrimCons(symhdv, tmp6107)

		tmp6109 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6110 := PrimCons(tmp6108, tmp6109)

		tmp6111 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6110)

		_ = tmp6111

		tmp6112 := PrimCons(symstring, Nil)

		tmp6113 := PrimCons(sym_1_1_6, tmp6112)

		tmp6114 := PrimCons(symstring, tmp6113)

		tmp6115 := PrimCons(symhdstr, tmp6114)

		tmp6116 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6117 := PrimCons(tmp6115, tmp6116)

		tmp6118 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6117)

		_ = tmp6118

		tmp6119 := PrimCons(symA, Nil)

		tmp6120 := PrimCons(sym_1_1_6, tmp6119)

		tmp6121 := PrimCons(symA, tmp6120)

		tmp6122 := PrimCons(tmp6121, Nil)

		tmp6123 := PrimCons(sym_1_1_6, tmp6122)

		tmp6124 := PrimCons(symA, tmp6123)

		tmp6125 := PrimCons(tmp6124, Nil)

		tmp6126 := PrimCons(sym_1_1_6, tmp6125)

		tmp6127 := PrimCons(symboolean, tmp6126)

		tmp6128 := PrimCons(symif, tmp6127)

		tmp6129 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6130 := PrimCons(tmp6128, tmp6129)

		tmp6131 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6130)

		_ = tmp6131

		tmp6132 := PrimCons(symstring, Nil)

		tmp6133 := PrimCons(sym_1_1_6, tmp6132)

		tmp6134 := PrimCons(symit, tmp6133)

		tmp6135 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6136 := PrimCons(tmp6134, tmp6135)

		tmp6137 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6136)

		_ = tmp6137

		tmp6138 := PrimCons(symstring, Nil)

		tmp6139 := PrimCons(sym_1_1_6, tmp6138)

		tmp6140 := PrimCons(symimplementation, tmp6139)

		tmp6141 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6142 := PrimCons(tmp6140, tmp6141)

		tmp6143 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6142)

		_ = tmp6143

		tmp6144 := PrimCons(symsymbol, Nil)

		tmp6145 := PrimCons(symlist, tmp6144)

		tmp6146 := PrimCons(symsymbol, Nil)

		tmp6147 := PrimCons(symlist, tmp6146)

		tmp6148 := PrimCons(tmp6147, Nil)

		tmp6149 := PrimCons(sym_1_1_6, tmp6148)

		tmp6150 := PrimCons(tmp6145, tmp6149)

		tmp6151 := PrimCons(syminclude, tmp6150)

		tmp6152 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6153 := PrimCons(tmp6151, tmp6152)

		tmp6154 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6153)

		_ = tmp6154

		tmp6155 := PrimCons(symsymbol, Nil)

		tmp6156 := PrimCons(symlist, tmp6155)

		tmp6157 := PrimCons(symsymbol, Nil)

		tmp6158 := PrimCons(symlist, tmp6157)

		tmp6159 := PrimCons(tmp6158, Nil)

		tmp6160 := PrimCons(sym_1_1_6, tmp6159)

		tmp6161 := PrimCons(tmp6156, tmp6160)

		tmp6162 := PrimCons(syminclude_1all_1but, tmp6161)

		tmp6163 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6164 := PrimCons(tmp6162, tmp6163)

		tmp6165 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6164)

		_ = tmp6165

		tmp6166 := PrimCons(symnumber, Nil)

		tmp6167 := PrimCons(sym_1_1_6, tmp6166)

		tmp6168 := PrimCons(syminferences, tmp6167)

		tmp6169 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6170 := PrimCons(tmp6168, tmp6169)

		tmp6171 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6170)

		_ = tmp6171

		tmp6172 := PrimCons(symstring, Nil)

		tmp6173 := PrimCons(sym_1_1_6, tmp6172)

		tmp6174 := PrimCons(symstring, tmp6173)

		tmp6175 := PrimCons(tmp6174, Nil)

		tmp6176 := PrimCons(sym_1_1_6, tmp6175)

		tmp6177 := PrimCons(symA, tmp6176)

		tmp6178 := PrimCons(symshen_4insert, tmp6177)

		tmp6179 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6180 := PrimCons(tmp6178, tmp6179)

		tmp6181 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6180)

		_ = tmp6181

		tmp6182 := PrimCons(symboolean, Nil)

		tmp6183 := PrimCons(sym_1_1_6, tmp6182)

		tmp6184 := PrimCons(symA, tmp6183)

		tmp6185 := PrimCons(syminteger_2, tmp6184)

		tmp6186 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6187 := PrimCons(tmp6185, tmp6186)

		tmp6188 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6187)

		_ = tmp6188

		tmp6189 := PrimCons(symsymbol, Nil)

		tmp6190 := PrimCons(symlist, tmp6189)

		tmp6191 := PrimCons(tmp6190, Nil)

		tmp6192 := PrimCons(sym_1_1_6, tmp6191)

		tmp6193 := PrimCons(symsymbol, tmp6192)

		tmp6194 := PrimCons(syminternal, tmp6193)

		tmp6195 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6196 := PrimCons(tmp6194, tmp6195)

		tmp6197 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6196)

		_ = tmp6197

		tmp6198 := PrimCons(symA, Nil)

		tmp6199 := PrimCons(symlist, tmp6198)

		tmp6200 := PrimCons(symA, Nil)

		tmp6201 := PrimCons(symlist, tmp6200)

		tmp6202 := PrimCons(symA, Nil)

		tmp6203 := PrimCons(symlist, tmp6202)

		tmp6204 := PrimCons(tmp6203, Nil)

		tmp6205 := PrimCons(sym_1_1_6, tmp6204)

		tmp6206 := PrimCons(tmp6201, tmp6205)

		tmp6207 := PrimCons(tmp6206, Nil)

		tmp6208 := PrimCons(sym_1_1_6, tmp6207)

		tmp6209 := PrimCons(tmp6199, tmp6208)

		tmp6210 := PrimCons(symintersection, tmp6209)

		tmp6211 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6212 := PrimCons(tmp6210, tmp6211)

		tmp6213 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6212)

		_ = tmp6213

		tmp6214 := PrimCons(symA, Nil)

		tmp6215 := PrimCons(sym_1_1_6, tmp6214)

		tmp6216 := PrimCons(symkill, tmp6215)

		tmp6217 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6218 := PrimCons(tmp6216, tmp6217)

		tmp6219 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6218)

		_ = tmp6219

		tmp6220 := PrimCons(symstring, Nil)

		tmp6221 := PrimCons(sym_1_1_6, tmp6220)

		tmp6222 := PrimCons(symlanguage, tmp6221)

		tmp6223 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6224 := PrimCons(tmp6222, tmp6223)

		tmp6225 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6224)

		_ = tmp6225

		tmp6226 := PrimCons(symA, Nil)

		tmp6227 := PrimCons(symlist, tmp6226)

		tmp6228 := PrimCons(symnumber, Nil)

		tmp6229 := PrimCons(sym_1_1_6, tmp6228)

		tmp6230 := PrimCons(tmp6227, tmp6229)

		tmp6231 := PrimCons(symlength, tmp6230)

		tmp6232 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6233 := PrimCons(tmp6231, tmp6232)

		tmp6234 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6233)

		_ = tmp6234

		tmp6235 := PrimCons(symA, Nil)

		tmp6236 := PrimCons(symvector, tmp6235)

		tmp6237 := PrimCons(symnumber, Nil)

		tmp6238 := PrimCons(sym_1_1_6, tmp6237)

		tmp6239 := PrimCons(tmp6236, tmp6238)

		tmp6240 := PrimCons(symlimit, tmp6239)

		tmp6241 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6242 := PrimCons(tmp6240, tmp6241)

		tmp6243 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6242)

		_ = tmp6243

		tmp6244 := PrimCons(symsymbol, Nil)

		tmp6245 := PrimCons(sym_1_1_6, tmp6244)

		tmp6246 := PrimCons(symstring, tmp6245)

		tmp6247 := PrimCons(symload, tmp6246)

		tmp6248 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6249 := PrimCons(tmp6247, tmp6248)

		tmp6250 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6249)

		_ = tmp6250

		tmp6251 := PrimCons(symB, Nil)

		tmp6252 := PrimCons(sym_1_1_6, tmp6251)

		tmp6253 := PrimCons(symA, tmp6252)

		tmp6254 := PrimCons(symA, Nil)

		tmp6255 := PrimCons(symlist, tmp6254)

		tmp6256 := PrimCons(symB, Nil)

		tmp6257 := PrimCons(symlist, tmp6256)

		tmp6258 := PrimCons(tmp6257, Nil)

		tmp6259 := PrimCons(sym_1_1_6, tmp6258)

		tmp6260 := PrimCons(tmp6255, tmp6259)

		tmp6261 := PrimCons(tmp6260, Nil)

		tmp6262 := PrimCons(sym_1_1_6, tmp6261)

		tmp6263 := PrimCons(tmp6253, tmp6262)

		tmp6264 := PrimCons(symmap, tmp6263)

		tmp6265 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6266 := PrimCons(tmp6264, tmp6265)

		tmp6267 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6266)

		_ = tmp6267

		tmp6268 := PrimCons(symB, Nil)

		tmp6269 := PrimCons(symlist, tmp6268)

		tmp6270 := PrimCons(tmp6269, Nil)

		tmp6271 := PrimCons(sym_1_1_6, tmp6270)

		tmp6272 := PrimCons(symA, tmp6271)

		tmp6273 := PrimCons(symA, Nil)

		tmp6274 := PrimCons(symlist, tmp6273)

		tmp6275 := PrimCons(symB, Nil)

		tmp6276 := PrimCons(symlist, tmp6275)

		tmp6277 := PrimCons(tmp6276, Nil)

		tmp6278 := PrimCons(sym_1_1_6, tmp6277)

		tmp6279 := PrimCons(tmp6274, tmp6278)

		tmp6280 := PrimCons(tmp6279, Nil)

		tmp6281 := PrimCons(sym_1_1_6, tmp6280)

		tmp6282 := PrimCons(tmp6272, tmp6281)

		tmp6283 := PrimCons(symmapcan, tmp6282)

		tmp6284 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6285 := PrimCons(tmp6283, tmp6284)

		tmp6286 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6285)

		_ = tmp6286

		tmp6287 := PrimCons(symnumber, Nil)

		tmp6288 := PrimCons(sym_1_1_6, tmp6287)

		tmp6289 := PrimCons(symnumber, tmp6288)

		tmp6290 := PrimCons(symmaxinferences, tmp6289)

		tmp6291 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6292 := PrimCons(tmp6290, tmp6291)

		tmp6293 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6292)

		_ = tmp6293

		tmp6294 := PrimCons(symstring, Nil)

		tmp6295 := PrimCons(sym_1_1_6, tmp6294)

		tmp6296 := PrimCons(symnumber, tmp6295)

		tmp6297 := PrimCons(symn_1_6string, tmp6296)

		tmp6298 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6299 := PrimCons(tmp6297, tmp6298)

		tmp6300 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6299)

		_ = tmp6300

		tmp6301 := PrimCons(symnumber, Nil)

		tmp6302 := PrimCons(sym_1_1_6, tmp6301)

		tmp6303 := PrimCons(symnumber, tmp6302)

		tmp6304 := PrimCons(symnl, tmp6303)

		tmp6305 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6306 := PrimCons(tmp6304, tmp6305)

		tmp6307 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6306)

		_ = tmp6307

		tmp6308 := PrimCons(symboolean, Nil)

		tmp6309 := PrimCons(sym_1_1_6, tmp6308)

		tmp6310 := PrimCons(symboolean, tmp6309)

		tmp6311 := PrimCons(symnot, tmp6310)

		tmp6312 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6313 := PrimCons(tmp6311, tmp6312)

		tmp6314 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6313)

		_ = tmp6314

		tmp6315 := PrimCons(symA, Nil)

		tmp6316 := PrimCons(symlist, tmp6315)

		tmp6317 := PrimCons(symA, Nil)

		tmp6318 := PrimCons(sym_1_1_6, tmp6317)

		tmp6319 := PrimCons(tmp6316, tmp6318)

		tmp6320 := PrimCons(tmp6319, Nil)

		tmp6321 := PrimCons(sym_1_1_6, tmp6320)

		tmp6322 := PrimCons(symnumber, tmp6321)

		tmp6323 := PrimCons(symnth, tmp6322)

		tmp6324 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6325 := PrimCons(tmp6323, tmp6324)

		tmp6326 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6325)

		_ = tmp6326

		tmp6327 := PrimCons(symboolean, Nil)

		tmp6328 := PrimCons(sym_1_1_6, tmp6327)

		tmp6329 := PrimCons(symA, tmp6328)

		tmp6330 := PrimCons(symnumber_2, tmp6329)

		tmp6331 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6332 := PrimCons(tmp6330, tmp6331)

		tmp6333 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6332)

		_ = tmp6333

		tmp6334 := PrimCons(symnumber, Nil)

		tmp6335 := PrimCons(sym_1_1_6, tmp6334)

		tmp6336 := PrimCons(symB, tmp6335)

		tmp6337 := PrimCons(tmp6336, Nil)

		tmp6338 := PrimCons(sym_1_1_6, tmp6337)

		tmp6339 := PrimCons(symA, tmp6338)

		tmp6340 := PrimCons(symoccurrences, tmp6339)

		tmp6341 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6342 := PrimCons(tmp6340, tmp6341)

		tmp6343 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6342)

		_ = tmp6343

		tmp6344 := PrimCons(symboolean, Nil)

		tmp6345 := PrimCons(sym_1_1_6, tmp6344)

		tmp6346 := PrimCons(symsymbol, tmp6345)

		tmp6347 := PrimCons(symoccurs_1check, tmp6346)

		tmp6348 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6349 := PrimCons(tmp6347, tmp6348)

		tmp6350 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6349)

		_ = tmp6350

		tmp6351 := PrimCons(symboolean, Nil)

		tmp6352 := PrimCons(sym_1_1_6, tmp6351)

		tmp6353 := PrimCons(symsymbol, tmp6352)

		tmp6354 := PrimCons(symoptimise, tmp6353)

		tmp6355 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6356 := PrimCons(tmp6354, tmp6355)

		tmp6357 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6356)

		_ = tmp6357

		tmp6358 := PrimCons(symboolean, Nil)

		tmp6359 := PrimCons(sym_1_1_6, tmp6358)

		tmp6360 := PrimCons(symboolean, tmp6359)

		tmp6361 := PrimCons(tmp6360, Nil)

		tmp6362 := PrimCons(sym_1_1_6, tmp6361)

		tmp6363 := PrimCons(symboolean, tmp6362)

		tmp6364 := PrimCons(symor, tmp6363)

		tmp6365 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6366 := PrimCons(tmp6364, tmp6365)

		tmp6367 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6366)

		_ = tmp6367

		tmp6368 := PrimCons(symstring, Nil)

		tmp6369 := PrimCons(sym_1_1_6, tmp6368)

		tmp6370 := PrimCons(symos, tmp6369)

		tmp6371 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6372 := PrimCons(tmp6370, tmp6371)

		tmp6373 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6372)

		_ = tmp6373

		tmp6374 := PrimCons(symboolean, Nil)

		tmp6375 := PrimCons(sym_1_1_6, tmp6374)

		tmp6376 := PrimCons(symsymbol, tmp6375)

		tmp6377 := PrimCons(sympackage_2, tmp6376)

		tmp6378 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6379 := PrimCons(tmp6377, tmp6378)

		tmp6380 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6379)

		_ = tmp6380

		tmp6381 := PrimCons(symstring, Nil)

		tmp6382 := PrimCons(sym_1_1_6, tmp6381)

		tmp6383 := PrimCons(symport, tmp6382)

		tmp6384 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6385 := PrimCons(tmp6383, tmp6384)

		tmp6386 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6385)

		_ = tmp6386

		tmp6387 := PrimCons(symstring, Nil)

		tmp6388 := PrimCons(sym_1_1_6, tmp6387)

		tmp6389 := PrimCons(symporters, tmp6388)

		tmp6390 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6391 := PrimCons(tmp6389, tmp6390)

		tmp6392 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6391)

		_ = tmp6392

		tmp6393 := PrimCons(symstring, Nil)

		tmp6394 := PrimCons(sym_1_1_6, tmp6393)

		tmp6395 := PrimCons(symnumber, tmp6394)

		tmp6396 := PrimCons(tmp6395, Nil)

		tmp6397 := PrimCons(sym_1_1_6, tmp6396)

		tmp6398 := PrimCons(symstring, tmp6397)

		tmp6399 := PrimCons(sympos, tmp6398)

		tmp6400 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6401 := PrimCons(tmp6399, tmp6400)

		tmp6402 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6401)

		_ = tmp6402

		tmp6403 := PrimCons(symout, Nil)

		tmp6404 := PrimCons(symstream, tmp6403)

		tmp6405 := PrimCons(symstring, Nil)

		tmp6406 := PrimCons(sym_1_1_6, tmp6405)

		tmp6407 := PrimCons(tmp6404, tmp6406)

		tmp6408 := PrimCons(tmp6407, Nil)

		tmp6409 := PrimCons(sym_1_1_6, tmp6408)

		tmp6410 := PrimCons(symstring, tmp6409)

		tmp6411 := PrimCons(sympr, tmp6410)

		tmp6412 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6413 := PrimCons(tmp6411, tmp6412)

		tmp6414 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6413)

		_ = tmp6414

		tmp6415 := PrimCons(symA, Nil)

		tmp6416 := PrimCons(sym_1_1_6, tmp6415)

		tmp6417 := PrimCons(symA, tmp6416)

		tmp6418 := PrimCons(symprint, tmp6417)

		tmp6419 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6420 := PrimCons(tmp6418, tmp6419)

		tmp6421 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6420)

		_ = tmp6421

		tmp6422 := PrimCons(symB, Nil)

		tmp6423 := PrimCons(sym_1_1_6, tmp6422)

		tmp6424 := PrimCons(symA, tmp6423)

		tmp6425 := PrimCons(symB, Nil)

		tmp6426 := PrimCons(sym_1_1_6, tmp6425)

		tmp6427 := PrimCons(symA, tmp6426)

		tmp6428 := PrimCons(tmp6427, Nil)

		tmp6429 := PrimCons(sym_1_1_6, tmp6428)

		tmp6430 := PrimCons(tmp6424, tmp6429)

		tmp6431 := PrimCons(symprofile, tmp6430)

		tmp6432 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6433 := PrimCons(tmp6431, tmp6432)

		tmp6434 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6433)

		_ = tmp6434

		tmp6435 := PrimCons(symsymbol, Nil)

		tmp6436 := PrimCons(symlist, tmp6435)

		tmp6437 := PrimCons(symsymbol, Nil)

		tmp6438 := PrimCons(symlist, tmp6437)

		tmp6439 := PrimCons(tmp6438, Nil)

		tmp6440 := PrimCons(sym_1_1_6, tmp6439)

		tmp6441 := PrimCons(tmp6436, tmp6440)

		tmp6442 := PrimCons(sympreclude, tmp6441)

		tmp6443 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6444 := PrimCons(tmp6442, tmp6443)

		tmp6445 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6444)

		_ = tmp6445

		tmp6446 := PrimCons(symstring, Nil)

		tmp6447 := PrimCons(sym_1_1_6, tmp6446)

		tmp6448 := PrimCons(symstring, tmp6447)

		tmp6449 := PrimCons(symshen_4proc_1nl, tmp6448)

		tmp6450 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6451 := PrimCons(tmp6449, tmp6450)

		tmp6452 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6451)

		_ = tmp6452

		tmp6453 := PrimCons(symB, Nil)

		tmp6454 := PrimCons(sym_1_1_6, tmp6453)

		tmp6455 := PrimCons(symA, tmp6454)

		tmp6456 := PrimCons(symB, Nil)

		tmp6457 := PrimCons(sym_1_1_6, tmp6456)

		tmp6458 := PrimCons(symA, tmp6457)

		tmp6459 := PrimCons(symnumber, Nil)

		tmp6460 := PrimCons(sym_d, tmp6459)

		tmp6461 := PrimCons(tmp6458, tmp6460)

		tmp6462 := PrimCons(tmp6461, Nil)

		tmp6463 := PrimCons(sym_1_1_6, tmp6462)

		tmp6464 := PrimCons(tmp6455, tmp6463)

		tmp6465 := PrimCons(symprofile_1results, tmp6464)

		tmp6466 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6467 := PrimCons(tmp6465, tmp6466)

		tmp6468 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6467)

		_ = tmp6468

		tmp6469 := PrimCons(symsymbol, Nil)

		tmp6470 := PrimCons(sym_1_1_6, tmp6469)

		tmp6471 := PrimCons(symsymbol, tmp6470)

		tmp6472 := PrimCons(symprotect, tmp6471)

		tmp6473 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6474 := PrimCons(tmp6472, tmp6473)

		tmp6475 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6474)

		_ = tmp6475

		tmp6476 := PrimCons(symsymbol, Nil)

		tmp6477 := PrimCons(symlist, tmp6476)

		tmp6478 := PrimCons(symsymbol, Nil)

		tmp6479 := PrimCons(symlist, tmp6478)

		tmp6480 := PrimCons(tmp6479, Nil)

		tmp6481 := PrimCons(sym_1_1_6, tmp6480)

		tmp6482 := PrimCons(tmp6477, tmp6481)

		tmp6483 := PrimCons(sympreclude_1all_1but, tmp6482)

		tmp6484 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6485 := PrimCons(tmp6483, tmp6484)

		tmp6486 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6485)

		_ = tmp6486

		tmp6487 := PrimCons(symout, Nil)

		tmp6488 := PrimCons(symstream, tmp6487)

		tmp6489 := PrimCons(symstring, Nil)

		tmp6490 := PrimCons(sym_1_1_6, tmp6489)

		tmp6491 := PrimCons(tmp6488, tmp6490)

		tmp6492 := PrimCons(tmp6491, Nil)

		tmp6493 := PrimCons(sym_1_1_6, tmp6492)

		tmp6494 := PrimCons(symstring, tmp6493)

		tmp6495 := PrimCons(symshen_4prhush, tmp6494)

		tmp6496 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6497 := PrimCons(tmp6495, tmp6496)

		tmp6498 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6497)

		_ = tmp6498

		tmp6499 := PrimCons(symunit, Nil)

		tmp6500 := PrimCons(symlist, tmp6499)

		tmp6501 := PrimCons(tmp6500, Nil)

		tmp6502 := PrimCons(sym_1_1_6, tmp6501)

		tmp6503 := PrimCons(symsymbol, tmp6502)

		tmp6504 := PrimCons(symps, tmp6503)

		tmp6505 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6506 := PrimCons(tmp6504, tmp6505)

		tmp6507 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6506)

		_ = tmp6507

		tmp6508 := PrimCons(symin, Nil)

		tmp6509 := PrimCons(symstream, tmp6508)

		tmp6510 := PrimCons(symunit, Nil)

		tmp6511 := PrimCons(sym_1_1_6, tmp6510)

		tmp6512 := PrimCons(tmp6509, tmp6511)

		tmp6513 := PrimCons(symread, tmp6512)

		tmp6514 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6515 := PrimCons(tmp6513, tmp6514)

		tmp6516 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6515)

		_ = tmp6516

		tmp6517 := PrimCons(symin, Nil)

		tmp6518 := PrimCons(symstream, tmp6517)

		tmp6519 := PrimCons(symnumber, Nil)

		tmp6520 := PrimCons(sym_1_1_6, tmp6519)

		tmp6521 := PrimCons(tmp6518, tmp6520)

		tmp6522 := PrimCons(symread_1byte, tmp6521)

		tmp6523 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6524 := PrimCons(tmp6522, tmp6523)

		tmp6525 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6524)

		_ = tmp6525

		tmp6526 := PrimCons(symnumber, Nil)

		tmp6527 := PrimCons(symlist, tmp6526)

		tmp6528 := PrimCons(tmp6527, Nil)

		tmp6529 := PrimCons(sym_1_1_6, tmp6528)

		tmp6530 := PrimCons(symstring, tmp6529)

		tmp6531 := PrimCons(symread_1file_1as_1bytelist, tmp6530)

		tmp6532 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6533 := PrimCons(tmp6531, tmp6532)

		tmp6534 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6533)

		_ = tmp6534

		tmp6535 := PrimCons(symstring, Nil)

		tmp6536 := PrimCons(sym_1_1_6, tmp6535)

		tmp6537 := PrimCons(symstring, tmp6536)

		tmp6538 := PrimCons(symread_1file_1as_1string, tmp6537)

		tmp6539 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6540 := PrimCons(tmp6538, tmp6539)

		tmp6541 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6540)

		_ = tmp6541

		tmp6542 := PrimCons(symunit, Nil)

		tmp6543 := PrimCons(symlist, tmp6542)

		tmp6544 := PrimCons(tmp6543, Nil)

		tmp6545 := PrimCons(sym_1_1_6, tmp6544)

		tmp6546 := PrimCons(symstring, tmp6545)

		tmp6547 := PrimCons(symread_1file, tmp6546)

		tmp6548 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6549 := PrimCons(tmp6547, tmp6548)

		tmp6550 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6549)

		_ = tmp6550

		tmp6551 := PrimCons(symunit, Nil)

		tmp6552 := PrimCons(symlist, tmp6551)

		tmp6553 := PrimCons(tmp6552, Nil)

		tmp6554 := PrimCons(sym_1_1_6, tmp6553)

		tmp6555 := PrimCons(symstring, tmp6554)

		tmp6556 := PrimCons(symread_1from_1string, tmp6555)

		tmp6557 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6558 := PrimCons(tmp6556, tmp6557)

		tmp6559 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6558)

		_ = tmp6559

		tmp6560 := PrimCons(symstring, Nil)

		tmp6561 := PrimCons(sym_1_1_6, tmp6560)

		tmp6562 := PrimCons(symrelease, tmp6561)

		tmp6563 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6564 := PrimCons(tmp6562, tmp6563)

		tmp6565 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6564)

		_ = tmp6565

		tmp6566 := PrimCons(symA, Nil)

		tmp6567 := PrimCons(symlist, tmp6566)

		tmp6568 := PrimCons(symA, Nil)

		tmp6569 := PrimCons(symlist, tmp6568)

		tmp6570 := PrimCons(tmp6569, Nil)

		tmp6571 := PrimCons(sym_1_1_6, tmp6570)

		tmp6572 := PrimCons(tmp6567, tmp6571)

		tmp6573 := PrimCons(tmp6572, Nil)

		tmp6574 := PrimCons(sym_1_1_6, tmp6573)

		tmp6575 := PrimCons(symA, tmp6574)

		tmp6576 := PrimCons(symremove, tmp6575)

		tmp6577 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6578 := PrimCons(tmp6576, tmp6577)

		tmp6579 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6578)

		_ = tmp6579

		tmp6580 := PrimCons(symA, Nil)

		tmp6581 := PrimCons(symlist, tmp6580)

		tmp6582 := PrimCons(symA, Nil)

		tmp6583 := PrimCons(symlist, tmp6582)

		tmp6584 := PrimCons(tmp6583, Nil)

		tmp6585 := PrimCons(sym_1_1_6, tmp6584)

		tmp6586 := PrimCons(tmp6581, tmp6585)

		tmp6587 := PrimCons(symreverse, tmp6586)

		tmp6588 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6589 := PrimCons(tmp6587, tmp6588)

		tmp6590 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6589)

		_ = tmp6590

		tmp6591 := PrimCons(symA, Nil)

		tmp6592 := PrimCons(sym_1_1_6, tmp6591)

		tmp6593 := PrimCons(symstring, tmp6592)

		tmp6594 := PrimCons(symsimple_1error, tmp6593)

		tmp6595 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6596 := PrimCons(tmp6594, tmp6595)

		tmp6597 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6596)

		_ = tmp6597

		tmp6598 := PrimCons(symB, Nil)

		tmp6599 := PrimCons(sym_d, tmp6598)

		tmp6600 := PrimCons(symA, tmp6599)

		tmp6601 := PrimCons(symB, Nil)

		tmp6602 := PrimCons(sym_1_1_6, tmp6601)

		tmp6603 := PrimCons(tmp6600, tmp6602)

		tmp6604 := PrimCons(symsnd, tmp6603)

		tmp6605 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6606 := PrimCons(tmp6604, tmp6605)

		tmp6607 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6606)

		_ = tmp6607

		tmp6608 := PrimCons(symsymbol, Nil)

		tmp6609 := PrimCons(sym_1_1_6, tmp6608)

		tmp6610 := PrimCons(symsymbol, tmp6609)

		tmp6611 := PrimCons(symspecialise, tmp6610)

		tmp6612 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6613 := PrimCons(tmp6611, tmp6612)

		tmp6614 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6613)

		_ = tmp6614

		tmp6615 := PrimCons(symboolean, Nil)

		tmp6616 := PrimCons(sym_1_1_6, tmp6615)

		tmp6617 := PrimCons(symsymbol, tmp6616)

		tmp6618 := PrimCons(symspy, tmp6617)

		tmp6619 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6620 := PrimCons(tmp6618, tmp6619)

		tmp6621 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6620)

		_ = tmp6621

		tmp6622 := PrimCons(symboolean, Nil)

		tmp6623 := PrimCons(sym_1_1_6, tmp6622)

		tmp6624 := PrimCons(symsymbol, tmp6623)

		tmp6625 := PrimCons(symstep, tmp6624)

		tmp6626 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6627 := PrimCons(tmp6625, tmp6626)

		tmp6628 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6627)

		_ = tmp6628

		tmp6629 := PrimCons(symin, Nil)

		tmp6630 := PrimCons(symstream, tmp6629)

		tmp6631 := PrimCons(tmp6630, Nil)

		tmp6632 := PrimCons(sym_1_1_6, tmp6631)

		tmp6633 := PrimCons(symstinput, tmp6632)

		tmp6634 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6635 := PrimCons(tmp6633, tmp6634)

		tmp6636 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6635)

		_ = tmp6636

		tmp6637 := PrimCons(symout, Nil)

		tmp6638 := PrimCons(symstream, tmp6637)

		tmp6639 := PrimCons(tmp6638, Nil)

		tmp6640 := PrimCons(sym_1_1_6, tmp6639)

		tmp6641 := PrimCons(symsterror, tmp6640)

		tmp6642 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6643 := PrimCons(tmp6641, tmp6642)

		tmp6644 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6643)

		_ = tmp6644

		tmp6645 := PrimCons(symout, Nil)

		tmp6646 := PrimCons(symstream, tmp6645)

		tmp6647 := PrimCons(tmp6646, Nil)

		tmp6648 := PrimCons(sym_1_1_6, tmp6647)

		tmp6649 := PrimCons(symstoutput, tmp6648)

		tmp6650 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6651 := PrimCons(tmp6649, tmp6650)

		tmp6652 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6651)

		_ = tmp6652

		tmp6653 := PrimCons(symboolean, Nil)

		tmp6654 := PrimCons(sym_1_1_6, tmp6653)

		tmp6655 := PrimCons(symA, tmp6654)

		tmp6656 := PrimCons(symstring_2, tmp6655)

		tmp6657 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6658 := PrimCons(tmp6656, tmp6657)

		tmp6659 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6658)

		_ = tmp6659

		tmp6660 := PrimCons(symstring, Nil)

		tmp6661 := PrimCons(sym_1_1_6, tmp6660)

		tmp6662 := PrimCons(symA, tmp6661)

		tmp6663 := PrimCons(symstr, tmp6662)

		tmp6664 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6665 := PrimCons(tmp6663, tmp6664)

		tmp6666 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6665)

		_ = tmp6666

		tmp6667 := PrimCons(symnumber, Nil)

		tmp6668 := PrimCons(sym_1_1_6, tmp6667)

		tmp6669 := PrimCons(symstring, tmp6668)

		tmp6670 := PrimCons(symstring_1_6n, tmp6669)

		tmp6671 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6672 := PrimCons(tmp6670, tmp6671)

		tmp6673 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6672)

		_ = tmp6673

		tmp6674 := PrimCons(symsymbol, Nil)

		tmp6675 := PrimCons(sym_1_1_6, tmp6674)

		tmp6676 := PrimCons(symstring, tmp6675)

		tmp6677 := PrimCons(symstring_1_6symbol, tmp6676)

		tmp6678 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6679 := PrimCons(tmp6677, tmp6678)

		tmp6680 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6679)

		_ = tmp6680

		tmp6681 := PrimCons(symnumber, Nil)

		tmp6682 := PrimCons(symlist, tmp6681)

		tmp6683 := PrimCons(symnumber, Nil)

		tmp6684 := PrimCons(sym_1_1_6, tmp6683)

		tmp6685 := PrimCons(tmp6682, tmp6684)

		tmp6686 := PrimCons(symsum, tmp6685)

		tmp6687 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6688 := PrimCons(tmp6686, tmp6687)

		tmp6689 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6688)

		_ = tmp6689

		tmp6690 := PrimCons(symboolean, Nil)

		tmp6691 := PrimCons(sym_1_1_6, tmp6690)

		tmp6692 := PrimCons(symA, tmp6691)

		tmp6693 := PrimCons(symsymbol_2, tmp6692)

		tmp6694 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6695 := PrimCons(tmp6693, tmp6694)

		tmp6696 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6695)

		_ = tmp6696

		tmp6697 := PrimCons(symsymbol, Nil)

		tmp6698 := PrimCons(sym_1_1_6, tmp6697)

		tmp6699 := PrimCons(symsymbol, tmp6698)

		tmp6700 := PrimCons(symsystemf, tmp6699)

		tmp6701 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6702 := PrimCons(tmp6700, tmp6701)

		tmp6703 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6702)

		_ = tmp6703

		tmp6704 := PrimCons(symA, Nil)

		tmp6705 := PrimCons(symlist, tmp6704)

		tmp6706 := PrimCons(symA, Nil)

		tmp6707 := PrimCons(symlist, tmp6706)

		tmp6708 := PrimCons(tmp6707, Nil)

		tmp6709 := PrimCons(sym_1_1_6, tmp6708)

		tmp6710 := PrimCons(tmp6705, tmp6709)

		tmp6711 := PrimCons(symtail, tmp6710)

		tmp6712 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6713 := PrimCons(tmp6711, tmp6712)

		tmp6714 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6713)

		_ = tmp6714

		tmp6715 := PrimCons(symstring, Nil)

		tmp6716 := PrimCons(sym_1_1_6, tmp6715)

		tmp6717 := PrimCons(symstring, tmp6716)

		tmp6718 := PrimCons(symtlstr, tmp6717)

		tmp6719 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6720 := PrimCons(tmp6718, tmp6719)

		tmp6721 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6720)

		_ = tmp6721

		tmp6722 := PrimCons(symA, Nil)

		tmp6723 := PrimCons(symvector, tmp6722)

		tmp6724 := PrimCons(symA, Nil)

		tmp6725 := PrimCons(symvector, tmp6724)

		tmp6726 := PrimCons(tmp6725, Nil)

		tmp6727 := PrimCons(sym_1_1_6, tmp6726)

		tmp6728 := PrimCons(tmp6723, tmp6727)

		tmp6729 := PrimCons(symtlv, tmp6728)

		tmp6730 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6731 := PrimCons(tmp6729, tmp6730)

		tmp6732 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6731)

		_ = tmp6732

		tmp6733 := PrimCons(symboolean, Nil)

		tmp6734 := PrimCons(sym_1_1_6, tmp6733)

		tmp6735 := PrimCons(symsymbol, tmp6734)

		tmp6736 := PrimCons(symtc, tmp6735)

		tmp6737 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6738 := PrimCons(tmp6736, tmp6737)

		tmp6739 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6738)

		_ = tmp6739

		tmp6740 := PrimCons(symboolean, Nil)

		tmp6741 := PrimCons(sym_1_1_6, tmp6740)

		tmp6742 := PrimCons(symtc_2, tmp6741)

		tmp6743 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6744 := PrimCons(tmp6742, tmp6743)

		tmp6745 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6744)

		_ = tmp6745

		tmp6746 := PrimCons(symA, Nil)

		tmp6747 := PrimCons(symlazy, tmp6746)

		tmp6748 := PrimCons(symA, Nil)

		tmp6749 := PrimCons(sym_1_1_6, tmp6748)

		tmp6750 := PrimCons(tmp6747, tmp6749)

		tmp6751 := PrimCons(symthaw, tmp6750)

		tmp6752 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6753 := PrimCons(tmp6751, tmp6752)

		tmp6754 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6753)

		_ = tmp6754

		tmp6755 := PrimCons(symsymbol, Nil)

		tmp6756 := PrimCons(sym_1_1_6, tmp6755)

		tmp6757 := PrimCons(symsymbol, tmp6756)

		tmp6758 := PrimCons(symtrack, tmp6757)

		tmp6759 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6760 := PrimCons(tmp6758, tmp6759)

		tmp6761 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6760)

		_ = tmp6761

		tmp6762 := PrimCons(symA, Nil)

		tmp6763 := PrimCons(sym_1_1_6, tmp6762)

		tmp6764 := PrimCons(symexception, tmp6763)

		tmp6765 := PrimCons(symA, Nil)

		tmp6766 := PrimCons(sym_1_1_6, tmp6765)

		tmp6767 := PrimCons(tmp6764, tmp6766)

		tmp6768 := PrimCons(tmp6767, Nil)

		tmp6769 := PrimCons(sym_1_1_6, tmp6768)

		tmp6770 := PrimCons(symA, tmp6769)

		tmp6771 := PrimCons(symtrap_1error, tmp6770)

		tmp6772 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6773 := PrimCons(tmp6771, tmp6772)

		tmp6774 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6773)

		_ = tmp6774

		tmp6775 := PrimCons(symboolean, Nil)

		tmp6776 := PrimCons(sym_1_1_6, tmp6775)

		tmp6777 := PrimCons(symA, tmp6776)

		tmp6778 := PrimCons(symtuple_2, tmp6777)

		tmp6779 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6780 := PrimCons(tmp6778, tmp6779)

		tmp6781 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6780)

		_ = tmp6781

		tmp6782 := PrimCons(symsymbol, Nil)

		tmp6783 := PrimCons(sym_1_1_6, tmp6782)

		tmp6784 := PrimCons(symsymbol, tmp6783)

		tmp6785 := PrimCons(symundefmacro, tmp6784)

		tmp6786 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6787 := PrimCons(tmp6785, tmp6786)

		tmp6788 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6787)

		_ = tmp6788

		tmp6789 := PrimCons(symA, Nil)

		tmp6790 := PrimCons(symlist, tmp6789)

		tmp6791 := PrimCons(symA, Nil)

		tmp6792 := PrimCons(symlist, tmp6791)

		tmp6793 := PrimCons(symA, Nil)

		tmp6794 := PrimCons(symlist, tmp6793)

		tmp6795 := PrimCons(tmp6794, Nil)

		tmp6796 := PrimCons(sym_1_1_6, tmp6795)

		tmp6797 := PrimCons(tmp6792, tmp6796)

		tmp6798 := PrimCons(tmp6797, Nil)

		tmp6799 := PrimCons(sym_1_1_6, tmp6798)

		tmp6800 := PrimCons(tmp6790, tmp6799)

		tmp6801 := PrimCons(symunion, tmp6800)

		tmp6802 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6803 := PrimCons(tmp6801, tmp6802)

		tmp6804 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6803)

		_ = tmp6804

		tmp6805 := PrimCons(symB, Nil)

		tmp6806 := PrimCons(sym_1_1_6, tmp6805)

		tmp6807 := PrimCons(symA, tmp6806)

		tmp6808 := PrimCons(symB, Nil)

		tmp6809 := PrimCons(sym_1_1_6, tmp6808)

		tmp6810 := PrimCons(symA, tmp6809)

		tmp6811 := PrimCons(tmp6810, Nil)

		tmp6812 := PrimCons(sym_1_1_6, tmp6811)

		tmp6813 := PrimCons(tmp6807, tmp6812)

		tmp6814 := PrimCons(symunprofile, tmp6813)

		tmp6815 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6816 := PrimCons(tmp6814, tmp6815)

		tmp6817 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6816)

		_ = tmp6817

		tmp6818 := PrimCons(symsymbol, Nil)

		tmp6819 := PrimCons(sym_1_1_6, tmp6818)

		tmp6820 := PrimCons(symsymbol, tmp6819)

		tmp6821 := PrimCons(symuntrack, tmp6820)

		tmp6822 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6823 := PrimCons(tmp6821, tmp6822)

		tmp6824 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6823)

		_ = tmp6824

		tmp6825 := PrimCons(symsymbol, Nil)

		tmp6826 := PrimCons(sym_1_1_6, tmp6825)

		tmp6827 := PrimCons(symsymbol, tmp6826)

		tmp6828 := PrimCons(symunspecialise, tmp6827)

		tmp6829 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6830 := PrimCons(tmp6828, tmp6829)

		tmp6831 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6830)

		_ = tmp6831

		tmp6832 := PrimCons(symboolean, Nil)

		tmp6833 := PrimCons(sym_1_1_6, tmp6832)

		tmp6834 := PrimCons(symA, tmp6833)

		tmp6835 := PrimCons(symvariable_2, tmp6834)

		tmp6836 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6837 := PrimCons(tmp6835, tmp6836)

		tmp6838 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6837)

		_ = tmp6838

		tmp6839 := PrimCons(symboolean, Nil)

		tmp6840 := PrimCons(sym_1_1_6, tmp6839)

		tmp6841 := PrimCons(symA, tmp6840)

		tmp6842 := PrimCons(symvector_2, tmp6841)

		tmp6843 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6844 := PrimCons(tmp6842, tmp6843)

		tmp6845 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6844)

		_ = tmp6845

		tmp6846 := PrimCons(symstring, Nil)

		tmp6847 := PrimCons(sym_1_1_6, tmp6846)

		tmp6848 := PrimCons(symversion, tmp6847)

		tmp6849 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6850 := PrimCons(tmp6848, tmp6849)

		tmp6851 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6850)

		_ = tmp6851

		tmp6852 := PrimCons(symA, Nil)

		tmp6853 := PrimCons(sym_1_1_6, tmp6852)

		tmp6854 := PrimCons(symA, tmp6853)

		tmp6855 := PrimCons(tmp6854, Nil)

		tmp6856 := PrimCons(sym_1_1_6, tmp6855)

		tmp6857 := PrimCons(symstring, tmp6856)

		tmp6858 := PrimCons(symwrite_1to_1file, tmp6857)

		tmp6859 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6860 := PrimCons(tmp6858, tmp6859)

		tmp6861 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6860)

		_ = tmp6861

		tmp6862 := PrimCons(symout, Nil)

		tmp6863 := PrimCons(symstream, tmp6862)

		tmp6864 := PrimCons(symnumber, Nil)

		tmp6865 := PrimCons(sym_1_1_6, tmp6864)

		tmp6866 := PrimCons(tmp6863, tmp6865)

		tmp6867 := PrimCons(tmp6866, Nil)

		tmp6868 := PrimCons(sym_1_1_6, tmp6867)

		tmp6869 := PrimCons(symnumber, tmp6868)

		tmp6870 := PrimCons(symwrite_1byte, tmp6869)

		tmp6871 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6872 := PrimCons(tmp6870, tmp6871)

		tmp6873 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6872)

		_ = tmp6873

		tmp6874 := PrimCons(symboolean, Nil)

		tmp6875 := PrimCons(sym_1_1_6, tmp6874)

		tmp6876 := PrimCons(symstring, tmp6875)

		tmp6877 := PrimCons(symy_1or_1n_2, tmp6876)

		tmp6878 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6879 := PrimCons(tmp6877, tmp6878)

		tmp6880 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6879)

		_ = tmp6880

		tmp6881 := PrimCons(symboolean, Nil)

		tmp6882 := PrimCons(sym_1_1_6, tmp6881)

		tmp6883 := PrimCons(symnumber, tmp6882)

		tmp6884 := PrimCons(tmp6883, Nil)

		tmp6885 := PrimCons(sym_1_1_6, tmp6884)

		tmp6886 := PrimCons(symnumber, tmp6885)

		tmp6887 := PrimCons(sym_6, tmp6886)

		tmp6888 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6889 := PrimCons(tmp6887, tmp6888)

		tmp6890 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6889)

		_ = tmp6890

		tmp6891 := PrimCons(symboolean, Nil)

		tmp6892 := PrimCons(sym_1_1_6, tmp6891)

		tmp6893 := PrimCons(symnumber, tmp6892)

		tmp6894 := PrimCons(tmp6893, Nil)

		tmp6895 := PrimCons(sym_1_1_6, tmp6894)

		tmp6896 := PrimCons(symnumber, tmp6895)

		tmp6897 := PrimCons(sym_5, tmp6896)

		tmp6898 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6899 := PrimCons(tmp6897, tmp6898)

		tmp6900 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6899)

		_ = tmp6900

		tmp6901 := PrimCons(symboolean, Nil)

		tmp6902 := PrimCons(sym_1_1_6, tmp6901)

		tmp6903 := PrimCons(symnumber, tmp6902)

		tmp6904 := PrimCons(tmp6903, Nil)

		tmp6905 := PrimCons(sym_1_1_6, tmp6904)

		tmp6906 := PrimCons(symnumber, tmp6905)

		tmp6907 := PrimCons(sym_6_a, tmp6906)

		tmp6908 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6909 := PrimCons(tmp6907, tmp6908)

		tmp6910 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6909)

		_ = tmp6910

		tmp6911 := PrimCons(symboolean, Nil)

		tmp6912 := PrimCons(sym_1_1_6, tmp6911)

		tmp6913 := PrimCons(symnumber, tmp6912)

		tmp6914 := PrimCons(tmp6913, Nil)

		tmp6915 := PrimCons(sym_1_1_6, tmp6914)

		tmp6916 := PrimCons(symnumber, tmp6915)

		tmp6917 := PrimCons(sym_5_a, tmp6916)

		tmp6918 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6919 := PrimCons(tmp6917, tmp6918)

		tmp6920 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6919)

		_ = tmp6920

		tmp6921 := PrimCons(symboolean, Nil)

		tmp6922 := PrimCons(sym_1_1_6, tmp6921)

		tmp6923 := PrimCons(symA, tmp6922)

		tmp6924 := PrimCons(tmp6923, Nil)

		tmp6925 := PrimCons(sym_1_1_6, tmp6924)

		tmp6926 := PrimCons(symA, tmp6925)

		tmp6927 := PrimCons(sym_a, tmp6926)

		tmp6928 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6929 := PrimCons(tmp6927, tmp6928)

		tmp6930 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6929)

		_ = tmp6930

		tmp6931 := PrimCons(symnumber, Nil)

		tmp6932 := PrimCons(sym_1_1_6, tmp6931)

		tmp6933 := PrimCons(symnumber, tmp6932)

		tmp6934 := PrimCons(tmp6933, Nil)

		tmp6935 := PrimCons(sym_1_1_6, tmp6934)

		tmp6936 := PrimCons(symnumber, tmp6935)

		tmp6937 := PrimCons(sym_7, tmp6936)

		tmp6938 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6939 := PrimCons(tmp6937, tmp6938)

		tmp6940 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6939)

		_ = tmp6940

		tmp6941 := PrimCons(symnumber, Nil)

		tmp6942 := PrimCons(sym_1_1_6, tmp6941)

		tmp6943 := PrimCons(symnumber, tmp6942)

		tmp6944 := PrimCons(tmp6943, Nil)

		tmp6945 := PrimCons(sym_1_1_6, tmp6944)

		tmp6946 := PrimCons(symnumber, tmp6945)

		tmp6947 := PrimCons(sym_c, tmp6946)

		tmp6948 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6949 := PrimCons(tmp6947, tmp6948)

		tmp6950 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6949)

		_ = tmp6950

		tmp6951 := PrimCons(symnumber, Nil)

		tmp6952 := PrimCons(sym_1_1_6, tmp6951)

		tmp6953 := PrimCons(symnumber, tmp6952)

		tmp6954 := PrimCons(tmp6953, Nil)

		tmp6955 := PrimCons(sym_1_1_6, tmp6954)

		tmp6956 := PrimCons(symnumber, tmp6955)

		tmp6957 := PrimCons(sym_1, tmp6956)

		tmp6958 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6959 := PrimCons(tmp6957, tmp6958)

		tmp6960 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6959)

		_ = tmp6960

		tmp6961 := PrimCons(symnumber, Nil)

		tmp6962 := PrimCons(sym_1_1_6, tmp6961)

		tmp6963 := PrimCons(symnumber, tmp6962)

		tmp6964 := PrimCons(tmp6963, Nil)

		tmp6965 := PrimCons(sym_1_1_6, tmp6964)

		tmp6966 := PrimCons(symnumber, tmp6965)

		tmp6967 := PrimCons(sym_d, tmp6966)

		tmp6968 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6969 := PrimCons(tmp6967, tmp6968)

		tmp6970 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6969)

		_ = tmp6970

		tmp6971 := PrimCons(symboolean, Nil)

		tmp6972 := PrimCons(sym_1_1_6, tmp6971)

		tmp6973 := PrimCons(symB, tmp6972)

		tmp6974 := PrimCons(tmp6973, Nil)

		tmp6975 := PrimCons(sym_1_1_6, tmp6974)

		tmp6976 := PrimCons(symA, tmp6975)

		tmp6977 := PrimCons(sym_a_a, tmp6976)

		tmp6978 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp6979 := PrimCons(tmp6977, tmp6978)

		__e.Return(PrimNS3Set(symshen_4_dsignedfuncs_d, tmp6979))
		return

	}, 0)

	tmp6980 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise_1signedfuncs, tmp5706)

	_ = tmp6980

	tmp6981 := MakeNative(func(__e *ControlFlow) {
		tmp6982 := MakeNative(func(__e *ControlFlow) {
			V3239 := __e.Get(1)
			_ = V3239
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3240 := __e.Get(1)
				_ = V3240
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3241 := __e.Get(1)
					_ = V3241
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1absvector_2), V3239, V3240, V3241)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp6983 := PrimCons(symshen_4type_1signature_1of_1absvector_2, tmp6982)

		tmp6984 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp6983)

		_ = tmp6984

		tmp6985 := MakeNative(func(__e *ControlFlow) {
			V3249 := __e.Get(1)
			_ = V3249
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3250 := __e.Get(1)
				_ = V3250
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3251 := __e.Get(1)
					_ = V3251
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1adjoin), V3249, V3250, V3251)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp6986 := PrimCons(symshen_4type_1signature_1of_1adjoin, tmp6985)

		tmp6987 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp6986)

		_ = tmp6987

		tmp6988 := MakeNative(func(__e *ControlFlow) {
			V3259 := __e.Get(1)
			_ = V3259
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3260 := __e.Get(1)
				_ = V3260
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3261 := __e.Get(1)
					_ = V3261
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1and), V3259, V3260, V3261)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp6989 := PrimCons(symshen_4type_1signature_1of_1and, tmp6988)

		tmp6990 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp6989)

		_ = tmp6990

		tmp6991 := MakeNative(func(__e *ControlFlow) {
			V3269 := __e.Get(1)
			_ = V3269
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3270 := __e.Get(1)
				_ = V3270
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3271 := __e.Get(1)
					_ = V3271
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1shen_4app), V3269, V3270, V3271)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp6992 := PrimCons(symshen_4type_1signature_1of_1shen_4app, tmp6991)

		tmp6993 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp6992)

		_ = tmp6993

		tmp6994 := MakeNative(func(__e *ControlFlow) {
			V3279 := __e.Get(1)
			_ = V3279
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3280 := __e.Get(1)
				_ = V3280
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3281 := __e.Get(1)
					_ = V3281
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1append), V3279, V3280, V3281)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp6995 := PrimCons(symshen_4type_1signature_1of_1append, tmp6994)

		tmp6996 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp6995)

		_ = tmp6996

		tmp6997 := MakeNative(func(__e *ControlFlow) {
			V3289 := __e.Get(1)
			_ = V3289
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3290 := __e.Get(1)
				_ = V3290
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3291 := __e.Get(1)
					_ = V3291
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1arity), V3289, V3290, V3291)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp6998 := PrimCons(symshen_4type_1signature_1of_1arity, tmp6997)

		tmp6999 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp6998)

		_ = tmp6999

		tmp7000 := MakeNative(func(__e *ControlFlow) {
			V3299 := __e.Get(1)
			_ = V3299
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3300 := __e.Get(1)
				_ = V3300
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3301 := __e.Get(1)
					_ = V3301
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1assoc), V3299, V3300, V3301)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7001 := PrimCons(symshen_4type_1signature_1of_1assoc, tmp7000)

		tmp7002 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7001)

		_ = tmp7002

		tmp7003 := MakeNative(func(__e *ControlFlow) {
			V3309 := __e.Get(1)
			_ = V3309
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3310 := __e.Get(1)
				_ = V3310
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3311 := __e.Get(1)
					_ = V3311
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1boolean_2), V3309, V3310, V3311)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7004 := PrimCons(symshen_4type_1signature_1of_1boolean_2, tmp7003)

		tmp7005 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7004)

		_ = tmp7005

		tmp7006 := MakeNative(func(__e *ControlFlow) {
			V3319 := __e.Get(1)
			_ = V3319
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3320 := __e.Get(1)
				_ = V3320
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3321 := __e.Get(1)
					_ = V3321
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1bound_2), V3319, V3320, V3321)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7007 := PrimCons(symshen_4type_1signature_1of_1bound_2, tmp7006)

		tmp7008 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7007)

		_ = tmp7008

		tmp7009 := MakeNative(func(__e *ControlFlow) {
			V3329 := __e.Get(1)
			_ = V3329
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3330 := __e.Get(1)
				_ = V3330
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3331 := __e.Get(1)
					_ = V3331
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1cd), V3329, V3330, V3331)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7010 := PrimCons(symshen_4type_1signature_1of_1cd, tmp7009)

		tmp7011 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7010)

		_ = tmp7011

		tmp7012 := MakeNative(func(__e *ControlFlow) {
			V3339 := __e.Get(1)
			_ = V3339
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3340 := __e.Get(1)
				_ = V3340
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3341 := __e.Get(1)
					_ = V3341
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1close), V3339, V3340, V3341)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7013 := PrimCons(symshen_4type_1signature_1of_1close, tmp7012)

		tmp7014 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7013)

		_ = tmp7014

		tmp7015 := MakeNative(func(__e *ControlFlow) {
			V3349 := __e.Get(1)
			_ = V3349
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3350 := __e.Get(1)
				_ = V3350
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3351 := __e.Get(1)
					_ = V3351
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1cn), V3349, V3350, V3351)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7016 := PrimCons(symshen_4type_1signature_1of_1cn, tmp7015)

		tmp7017 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7016)

		_ = tmp7017

		tmp7018 := MakeNative(func(__e *ControlFlow) {
			V3359 := __e.Get(1)
			_ = V3359
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3360 := __e.Get(1)
				_ = V3360
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3361 := __e.Get(1)
					_ = V3361
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1compile), V3359, V3360, V3361)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7019 := PrimCons(symshen_4type_1signature_1of_1compile, tmp7018)

		tmp7020 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7019)

		_ = tmp7020

		tmp7021 := MakeNative(func(__e *ControlFlow) {
			V3369 := __e.Get(1)
			_ = V3369
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3370 := __e.Get(1)
				_ = V3370
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3371 := __e.Get(1)
					_ = V3371
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1cons_2), V3369, V3370, V3371)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7022 := PrimCons(symshen_4type_1signature_1of_1cons_2, tmp7021)

		tmp7023 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7022)

		_ = tmp7023

		tmp7024 := MakeNative(func(__e *ControlFlow) {
			V3379 := __e.Get(1)
			_ = V3379
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3380 := __e.Get(1)
				_ = V3380
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3381 := __e.Get(1)
					_ = V3381
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1destroy), V3379, V3380, V3381)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7025 := PrimCons(symshen_4type_1signature_1of_1destroy, tmp7024)

		tmp7026 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7025)

		_ = tmp7026

		tmp7027 := MakeNative(func(__e *ControlFlow) {
			V3389 := __e.Get(1)
			_ = V3389
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3390 := __e.Get(1)
				_ = V3390
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3391 := __e.Get(1)
					_ = V3391
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1difference), V3389, V3390, V3391)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7028 := PrimCons(symshen_4type_1signature_1of_1difference, tmp7027)

		tmp7029 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7028)

		_ = tmp7029

		tmp7030 := MakeNative(func(__e *ControlFlow) {
			V3399 := __e.Get(1)
			_ = V3399
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3400 := __e.Get(1)
				_ = V3400
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3401 := __e.Get(1)
					_ = V3401
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1do), V3399, V3400, V3401)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7031 := PrimCons(symshen_4type_1signature_1of_1do, tmp7030)

		tmp7032 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7031)

		_ = tmp7032

		tmp7033 := MakeNative(func(__e *ControlFlow) {
			V3409 := __e.Get(1)
			_ = V3409
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3410 := __e.Get(1)
				_ = V3410
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3411 := __e.Get(1)
					_ = V3411
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_5e_6), V3409, V3410, V3411)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7034 := PrimCons(symshen_4type_1signature_1of_1_5e_6, tmp7033)

		tmp7035 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7034)

		_ = tmp7035

		tmp7036 := MakeNative(func(__e *ControlFlow) {
			V3419 := __e.Get(1)
			_ = V3419
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3420 := __e.Get(1)
				_ = V3420
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3421 := __e.Get(1)
					_ = V3421
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_5_b_6), V3419, V3420, V3421)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7037 := PrimCons(symshen_4type_1signature_1of_1_5_b_6, tmp7036)

		tmp7038 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7037)

		_ = tmp7038

		tmp7039 := MakeNative(func(__e *ControlFlow) {
			V3429 := __e.Get(1)
			_ = V3429
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3430 := __e.Get(1)
				_ = V3430
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3431 := __e.Get(1)
					_ = V3431
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1element_2), V3429, V3430, V3431)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7040 := PrimCons(symshen_4type_1signature_1of_1element_2, tmp7039)

		tmp7041 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7040)

		_ = tmp7041

		tmp7042 := MakeNative(func(__e *ControlFlow) {
			V3439 := __e.Get(1)
			_ = V3439
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3440 := __e.Get(1)
				_ = V3440
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3441 := __e.Get(1)
					_ = V3441
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1empty_2), V3439, V3440, V3441)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7043 := PrimCons(symshen_4type_1signature_1of_1empty_2, tmp7042)

		tmp7044 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7043)

		_ = tmp7044

		tmp7045 := MakeNative(func(__e *ControlFlow) {
			V3449 := __e.Get(1)
			_ = V3449
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3450 := __e.Get(1)
				_ = V3450
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3451 := __e.Get(1)
					_ = V3451
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1enable_1type_1theory), V3449, V3450, V3451)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7046 := PrimCons(symshen_4type_1signature_1of_1enable_1type_1theory, tmp7045)

		tmp7047 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7046)

		_ = tmp7047

		tmp7048 := MakeNative(func(__e *ControlFlow) {
			V3459 := __e.Get(1)
			_ = V3459
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3460 := __e.Get(1)
				_ = V3460
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3461 := __e.Get(1)
					_ = V3461
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1external), V3459, V3460, V3461)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7049 := PrimCons(symshen_4type_1signature_1of_1external, tmp7048)

		tmp7050 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7049)

		_ = tmp7050

		tmp7051 := MakeNative(func(__e *ControlFlow) {
			V3469 := __e.Get(1)
			_ = V3469
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3470 := __e.Get(1)
				_ = V3470
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3471 := __e.Get(1)
					_ = V3471
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1error_1to_1string), V3469, V3470, V3471)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7052 := PrimCons(symshen_4type_1signature_1of_1error_1to_1string, tmp7051)

		tmp7053 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7052)

		_ = tmp7053

		tmp7054 := MakeNative(func(__e *ControlFlow) {
			V3479 := __e.Get(1)
			_ = V3479
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3480 := __e.Get(1)
				_ = V3480
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3481 := __e.Get(1)
					_ = V3481
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1explode), V3479, V3480, V3481)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7055 := PrimCons(symshen_4type_1signature_1of_1explode, tmp7054)

		tmp7056 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7055)

		_ = tmp7056

		tmp7057 := MakeNative(func(__e *ControlFlow) {
			V3489 := __e.Get(1)
			_ = V3489
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3490 := __e.Get(1)
				_ = V3490
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3491 := __e.Get(1)
					_ = V3491
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1fail), V3489, V3490, V3491)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7058 := PrimCons(symshen_4type_1signature_1of_1fail, tmp7057)

		tmp7059 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7058)

		_ = tmp7059

		tmp7060 := MakeNative(func(__e *ControlFlow) {
			V3499 := __e.Get(1)
			_ = V3499
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3500 := __e.Get(1)
				_ = V3500
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3501 := __e.Get(1)
					_ = V3501
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1fail_1if), V3499, V3500, V3501)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7061 := PrimCons(symshen_4type_1signature_1of_1fail_1if, tmp7060)

		tmp7062 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7061)

		_ = tmp7062

		tmp7063 := MakeNative(func(__e *ControlFlow) {
			V3509 := __e.Get(1)
			_ = V3509
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3510 := __e.Get(1)
				_ = V3510
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3511 := __e.Get(1)
					_ = V3511
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1fix), V3509, V3510, V3511)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7064 := PrimCons(symshen_4type_1signature_1of_1fix, tmp7063)

		tmp7065 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7064)

		_ = tmp7065

		tmp7066 := MakeNative(func(__e *ControlFlow) {
			V3519 := __e.Get(1)
			_ = V3519
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3520 := __e.Get(1)
				_ = V3520
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3521 := __e.Get(1)
					_ = V3521
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1freeze), V3519, V3520, V3521)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7067 := PrimCons(symshen_4type_1signature_1of_1freeze, tmp7066)

		tmp7068 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7067)

		_ = tmp7068

		tmp7069 := MakeNative(func(__e *ControlFlow) {
			V3529 := __e.Get(1)
			_ = V3529
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3530 := __e.Get(1)
				_ = V3530
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3531 := __e.Get(1)
					_ = V3531
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1fst), V3529, V3530, V3531)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7070 := PrimCons(symshen_4type_1signature_1of_1fst, tmp7069)

		tmp7071 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7070)

		_ = tmp7071

		tmp7072 := MakeNative(func(__e *ControlFlow) {
			V3539 := __e.Get(1)
			_ = V3539
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3540 := __e.Get(1)
				_ = V3540
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3541 := __e.Get(1)
					_ = V3541
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1function), V3539, V3540, V3541)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7073 := PrimCons(symshen_4type_1signature_1of_1function, tmp7072)

		tmp7074 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7073)

		_ = tmp7074

		tmp7075 := MakeNative(func(__e *ControlFlow) {
			V3549 := __e.Get(1)
			_ = V3549
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3550 := __e.Get(1)
				_ = V3550
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3551 := __e.Get(1)
					_ = V3551
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1gensym), V3549, V3550, V3551)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7076 := PrimCons(symshen_4type_1signature_1of_1gensym, tmp7075)

		tmp7077 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7076)

		_ = tmp7077

		tmp7078 := MakeNative(func(__e *ControlFlow) {
			V3559 := __e.Get(1)
			_ = V3559
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3560 := __e.Get(1)
				_ = V3560
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3561 := __e.Get(1)
					_ = V3561
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_5_1vector), V3559, V3560, V3561)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7079 := PrimCons(symshen_4type_1signature_1of_1_5_1vector, tmp7078)

		tmp7080 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7079)

		_ = tmp7080

		tmp7081 := MakeNative(func(__e *ControlFlow) {
			V3569 := __e.Get(1)
			_ = V3569
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3570 := __e.Get(1)
				_ = V3570
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3571 := __e.Get(1)
					_ = V3571
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1vector_1_6), V3569, V3570, V3571)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7082 := PrimCons(symshen_4type_1signature_1of_1vector_1_6, tmp7081)

		tmp7083 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7082)

		_ = tmp7083

		tmp7084 := MakeNative(func(__e *ControlFlow) {
			V3579 := __e.Get(1)
			_ = V3579
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3580 := __e.Get(1)
				_ = V3580
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3581 := __e.Get(1)
					_ = V3581
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1vector), V3579, V3580, V3581)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7085 := PrimCons(symshen_4type_1signature_1of_1vector, tmp7084)

		tmp7086 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7085)

		_ = tmp7086

		tmp7087 := MakeNative(func(__e *ControlFlow) {
			V3589 := __e.Get(1)
			_ = V3589
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3590 := __e.Get(1)
				_ = V3590
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3591 := __e.Get(1)
					_ = V3591
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1get_1time), V3589, V3590, V3591)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7088 := PrimCons(symshen_4type_1signature_1of_1get_1time, tmp7087)

		tmp7089 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7088)

		_ = tmp7089

		tmp7090 := MakeNative(func(__e *ControlFlow) {
			V3599 := __e.Get(1)
			_ = V3599
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3600 := __e.Get(1)
				_ = V3600
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3601 := __e.Get(1)
					_ = V3601
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1hash), V3599, V3600, V3601)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7091 := PrimCons(symshen_4type_1signature_1of_1hash, tmp7090)

		tmp7092 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7091)

		_ = tmp7092

		tmp7093 := MakeNative(func(__e *ControlFlow) {
			V3609 := __e.Get(1)
			_ = V3609
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3610 := __e.Get(1)
				_ = V3610
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3611 := __e.Get(1)
					_ = V3611
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1head), V3609, V3610, V3611)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7094 := PrimCons(symshen_4type_1signature_1of_1head, tmp7093)

		tmp7095 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7094)

		_ = tmp7095

		tmp7096 := MakeNative(func(__e *ControlFlow) {
			V3619 := __e.Get(1)
			_ = V3619
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3620 := __e.Get(1)
				_ = V3620
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3621 := __e.Get(1)
					_ = V3621
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1hdv), V3619, V3620, V3621)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7097 := PrimCons(symshen_4type_1signature_1of_1hdv, tmp7096)

		tmp7098 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7097)

		_ = tmp7098

		tmp7099 := MakeNative(func(__e *ControlFlow) {
			V3629 := __e.Get(1)
			_ = V3629
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3630 := __e.Get(1)
				_ = V3630
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3631 := __e.Get(1)
					_ = V3631
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1hdstr), V3629, V3630, V3631)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7100 := PrimCons(symshen_4type_1signature_1of_1hdstr, tmp7099)

		tmp7101 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7100)

		_ = tmp7101

		tmp7102 := MakeNative(func(__e *ControlFlow) {
			V3639 := __e.Get(1)
			_ = V3639
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3640 := __e.Get(1)
				_ = V3640
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3641 := __e.Get(1)
					_ = V3641
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1if), V3639, V3640, V3641)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7103 := PrimCons(symshen_4type_1signature_1of_1if, tmp7102)

		tmp7104 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7103)

		_ = tmp7104

		tmp7105 := MakeNative(func(__e *ControlFlow) {
			V3649 := __e.Get(1)
			_ = V3649
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3650 := __e.Get(1)
				_ = V3650
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3651 := __e.Get(1)
					_ = V3651
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1it), V3649, V3650, V3651)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7106 := PrimCons(symshen_4type_1signature_1of_1it, tmp7105)

		tmp7107 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7106)

		_ = tmp7107

		tmp7108 := MakeNative(func(__e *ControlFlow) {
			V3659 := __e.Get(1)
			_ = V3659
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3660 := __e.Get(1)
				_ = V3660
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3661 := __e.Get(1)
					_ = V3661
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1implementation), V3659, V3660, V3661)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7109 := PrimCons(symshen_4type_1signature_1of_1implementation, tmp7108)

		tmp7110 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7109)

		_ = tmp7110

		tmp7111 := MakeNative(func(__e *ControlFlow) {
			V3669 := __e.Get(1)
			_ = V3669
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3670 := __e.Get(1)
				_ = V3670
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3671 := __e.Get(1)
					_ = V3671
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1include), V3669, V3670, V3671)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7112 := PrimCons(symshen_4type_1signature_1of_1include, tmp7111)

		tmp7113 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7112)

		_ = tmp7113

		tmp7114 := MakeNative(func(__e *ControlFlow) {
			V3679 := __e.Get(1)
			_ = V3679
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3680 := __e.Get(1)
				_ = V3680
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3681 := __e.Get(1)
					_ = V3681
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1include_1all_1but), V3679, V3680, V3681)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7115 := PrimCons(symshen_4type_1signature_1of_1include_1all_1but, tmp7114)

		tmp7116 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7115)

		_ = tmp7116

		tmp7117 := MakeNative(func(__e *ControlFlow) {
			V3689 := __e.Get(1)
			_ = V3689
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3690 := __e.Get(1)
				_ = V3690
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3691 := __e.Get(1)
					_ = V3691
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1inferences), V3689, V3690, V3691)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7118 := PrimCons(symshen_4type_1signature_1of_1inferences, tmp7117)

		tmp7119 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7118)

		_ = tmp7119

		tmp7120 := MakeNative(func(__e *ControlFlow) {
			V3699 := __e.Get(1)
			_ = V3699
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3700 := __e.Get(1)
				_ = V3700
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3701 := __e.Get(1)
					_ = V3701
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1shen_4insert), V3699, V3700, V3701)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7121 := PrimCons(symshen_4type_1signature_1of_1shen_4insert, tmp7120)

		tmp7122 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7121)

		_ = tmp7122

		tmp7123 := MakeNative(func(__e *ControlFlow) {
			V3709 := __e.Get(1)
			_ = V3709
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3710 := __e.Get(1)
				_ = V3710
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3711 := __e.Get(1)
					_ = V3711
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1integer_2), V3709, V3710, V3711)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7124 := PrimCons(symshen_4type_1signature_1of_1integer_2, tmp7123)

		tmp7125 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7124)

		_ = tmp7125

		tmp7126 := MakeNative(func(__e *ControlFlow) {
			V3719 := __e.Get(1)
			_ = V3719
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3720 := __e.Get(1)
				_ = V3720
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3721 := __e.Get(1)
					_ = V3721
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1internal), V3719, V3720, V3721)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7127 := PrimCons(symshen_4type_1signature_1of_1internal, tmp7126)

		tmp7128 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7127)

		_ = tmp7128

		tmp7129 := MakeNative(func(__e *ControlFlow) {
			V3729 := __e.Get(1)
			_ = V3729
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3730 := __e.Get(1)
				_ = V3730
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3731 := __e.Get(1)
					_ = V3731
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1intersection), V3729, V3730, V3731)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7130 := PrimCons(symshen_4type_1signature_1of_1intersection, tmp7129)

		tmp7131 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7130)

		_ = tmp7131

		tmp7132 := MakeNative(func(__e *ControlFlow) {
			V3739 := __e.Get(1)
			_ = V3739
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3740 := __e.Get(1)
				_ = V3740
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3741 := __e.Get(1)
					_ = V3741
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1kill), V3739, V3740, V3741)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7133 := PrimCons(symshen_4type_1signature_1of_1kill, tmp7132)

		tmp7134 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7133)

		_ = tmp7134

		tmp7135 := MakeNative(func(__e *ControlFlow) {
			V3749 := __e.Get(1)
			_ = V3749
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3750 := __e.Get(1)
				_ = V3750
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3751 := __e.Get(1)
					_ = V3751
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1language), V3749, V3750, V3751)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7136 := PrimCons(symshen_4type_1signature_1of_1language, tmp7135)

		tmp7137 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7136)

		_ = tmp7137

		tmp7138 := MakeNative(func(__e *ControlFlow) {
			V3759 := __e.Get(1)
			_ = V3759
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3760 := __e.Get(1)
				_ = V3760
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3761 := __e.Get(1)
					_ = V3761
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1length), V3759, V3760, V3761)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7139 := PrimCons(symshen_4type_1signature_1of_1length, tmp7138)

		tmp7140 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7139)

		_ = tmp7140

		tmp7141 := MakeNative(func(__e *ControlFlow) {
			V3769 := __e.Get(1)
			_ = V3769
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3770 := __e.Get(1)
				_ = V3770
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3771 := __e.Get(1)
					_ = V3771
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1limit), V3769, V3770, V3771)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7142 := PrimCons(symshen_4type_1signature_1of_1limit, tmp7141)

		tmp7143 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7142)

		_ = tmp7143

		tmp7144 := MakeNative(func(__e *ControlFlow) {
			V3779 := __e.Get(1)
			_ = V3779
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3780 := __e.Get(1)
				_ = V3780
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3781 := __e.Get(1)
					_ = V3781
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1load), V3779, V3780, V3781)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7145 := PrimCons(symshen_4type_1signature_1of_1load, tmp7144)

		tmp7146 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7145)

		_ = tmp7146

		tmp7147 := MakeNative(func(__e *ControlFlow) {
			V3789 := __e.Get(1)
			_ = V3789
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3790 := __e.Get(1)
				_ = V3790
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3791 := __e.Get(1)
					_ = V3791
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1map), V3789, V3790, V3791)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7148 := PrimCons(symshen_4type_1signature_1of_1map, tmp7147)

		tmp7149 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7148)

		_ = tmp7149

		tmp7150 := MakeNative(func(__e *ControlFlow) {
			V3799 := __e.Get(1)
			_ = V3799
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3800 := __e.Get(1)
				_ = V3800
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3801 := __e.Get(1)
					_ = V3801
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1mapcan), V3799, V3800, V3801)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7151 := PrimCons(symshen_4type_1signature_1of_1mapcan, tmp7150)

		tmp7152 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7151)

		_ = tmp7152

		tmp7153 := MakeNative(func(__e *ControlFlow) {
			V3809 := __e.Get(1)
			_ = V3809
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3810 := __e.Get(1)
				_ = V3810
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3811 := __e.Get(1)
					_ = V3811
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1maxinferences), V3809, V3810, V3811)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7154 := PrimCons(symshen_4type_1signature_1of_1maxinferences, tmp7153)

		tmp7155 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7154)

		_ = tmp7155

		tmp7156 := MakeNative(func(__e *ControlFlow) {
			V3819 := __e.Get(1)
			_ = V3819
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3820 := __e.Get(1)
				_ = V3820
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3821 := __e.Get(1)
					_ = V3821
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1n_1_6string), V3819, V3820, V3821)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7157 := PrimCons(symshen_4type_1signature_1of_1n_1_6string, tmp7156)

		tmp7158 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7157)

		_ = tmp7158

		tmp7159 := MakeNative(func(__e *ControlFlow) {
			V3829 := __e.Get(1)
			_ = V3829
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3830 := __e.Get(1)
				_ = V3830
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3831 := __e.Get(1)
					_ = V3831
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1nl), V3829, V3830, V3831)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7160 := PrimCons(symshen_4type_1signature_1of_1nl, tmp7159)

		tmp7161 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7160)

		_ = tmp7161

		tmp7162 := MakeNative(func(__e *ControlFlow) {
			V3839 := __e.Get(1)
			_ = V3839
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3840 := __e.Get(1)
				_ = V3840
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3841 := __e.Get(1)
					_ = V3841
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1not), V3839, V3840, V3841)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7163 := PrimCons(symshen_4type_1signature_1of_1not, tmp7162)

		tmp7164 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7163)

		_ = tmp7164

		tmp7165 := MakeNative(func(__e *ControlFlow) {
			V3849 := __e.Get(1)
			_ = V3849
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3850 := __e.Get(1)
				_ = V3850
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3851 := __e.Get(1)
					_ = V3851
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1nth), V3849, V3850, V3851)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7166 := PrimCons(symshen_4type_1signature_1of_1nth, tmp7165)

		tmp7167 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7166)

		_ = tmp7167

		tmp7168 := MakeNative(func(__e *ControlFlow) {
			V3859 := __e.Get(1)
			_ = V3859
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3860 := __e.Get(1)
				_ = V3860
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3861 := __e.Get(1)
					_ = V3861
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1number_2), V3859, V3860, V3861)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7169 := PrimCons(symshen_4type_1signature_1of_1number_2, tmp7168)

		tmp7170 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7169)

		_ = tmp7170

		tmp7171 := MakeNative(func(__e *ControlFlow) {
			V3869 := __e.Get(1)
			_ = V3869
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3870 := __e.Get(1)
				_ = V3870
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3871 := __e.Get(1)
					_ = V3871
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1occurrences), V3869, V3870, V3871)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7172 := PrimCons(symshen_4type_1signature_1of_1occurrences, tmp7171)

		tmp7173 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7172)

		_ = tmp7173

		tmp7174 := MakeNative(func(__e *ControlFlow) {
			V3879 := __e.Get(1)
			_ = V3879
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3880 := __e.Get(1)
				_ = V3880
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3881 := __e.Get(1)
					_ = V3881
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1occurs_1check), V3879, V3880, V3881)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7175 := PrimCons(symshen_4type_1signature_1of_1occurs_1check, tmp7174)

		tmp7176 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7175)

		_ = tmp7176

		tmp7177 := MakeNative(func(__e *ControlFlow) {
			V3889 := __e.Get(1)
			_ = V3889
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3890 := __e.Get(1)
				_ = V3890
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3891 := __e.Get(1)
					_ = V3891
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1optimise), V3889, V3890, V3891)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7178 := PrimCons(symshen_4type_1signature_1of_1optimise, tmp7177)

		tmp7179 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7178)

		_ = tmp7179

		tmp7180 := MakeNative(func(__e *ControlFlow) {
			V3899 := __e.Get(1)
			_ = V3899
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3900 := __e.Get(1)
				_ = V3900
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3901 := __e.Get(1)
					_ = V3901
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1or), V3899, V3900, V3901)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7181 := PrimCons(symshen_4type_1signature_1of_1or, tmp7180)

		tmp7182 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7181)

		_ = tmp7182

		tmp7183 := MakeNative(func(__e *ControlFlow) {
			V3909 := __e.Get(1)
			_ = V3909
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3910 := __e.Get(1)
				_ = V3910
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3911 := __e.Get(1)
					_ = V3911
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1os), V3909, V3910, V3911)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7184 := PrimCons(symshen_4type_1signature_1of_1os, tmp7183)

		tmp7185 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7184)

		_ = tmp7185

		tmp7186 := MakeNative(func(__e *ControlFlow) {
			V3919 := __e.Get(1)
			_ = V3919
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3920 := __e.Get(1)
				_ = V3920
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3921 := __e.Get(1)
					_ = V3921
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1package_2), V3919, V3920, V3921)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7187 := PrimCons(symshen_4type_1signature_1of_1package_2, tmp7186)

		tmp7188 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7187)

		_ = tmp7188

		tmp7189 := MakeNative(func(__e *ControlFlow) {
			V3929 := __e.Get(1)
			_ = V3929
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3930 := __e.Get(1)
				_ = V3930
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3931 := __e.Get(1)
					_ = V3931
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1port), V3929, V3930, V3931)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7190 := PrimCons(symshen_4type_1signature_1of_1port, tmp7189)

		tmp7191 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7190)

		_ = tmp7191

		tmp7192 := MakeNative(func(__e *ControlFlow) {
			V3939 := __e.Get(1)
			_ = V3939
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3940 := __e.Get(1)
				_ = V3940
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3941 := __e.Get(1)
					_ = V3941
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1porters), V3939, V3940, V3941)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7193 := PrimCons(symshen_4type_1signature_1of_1porters, tmp7192)

		tmp7194 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7193)

		_ = tmp7194

		tmp7195 := MakeNative(func(__e *ControlFlow) {
			V3949 := __e.Get(1)
			_ = V3949
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3950 := __e.Get(1)
				_ = V3950
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3951 := __e.Get(1)
					_ = V3951
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1pos), V3949, V3950, V3951)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7196 := PrimCons(symshen_4type_1signature_1of_1pos, tmp7195)

		tmp7197 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7196)

		_ = tmp7197

		tmp7198 := MakeNative(func(__e *ControlFlow) {
			V3959 := __e.Get(1)
			_ = V3959
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3960 := __e.Get(1)
				_ = V3960
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3961 := __e.Get(1)
					_ = V3961
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1pr), V3959, V3960, V3961)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7199 := PrimCons(symshen_4type_1signature_1of_1pr, tmp7198)

		tmp7200 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7199)

		_ = tmp7200

		tmp7201 := MakeNative(func(__e *ControlFlow) {
			V3969 := __e.Get(1)
			_ = V3969
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3970 := __e.Get(1)
				_ = V3970
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3971 := __e.Get(1)
					_ = V3971
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1print), V3969, V3970, V3971)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7202 := PrimCons(symshen_4type_1signature_1of_1print, tmp7201)

		tmp7203 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7202)

		_ = tmp7203

		tmp7204 := MakeNative(func(__e *ControlFlow) {
			V3979 := __e.Get(1)
			_ = V3979
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3980 := __e.Get(1)
				_ = V3980
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3981 := __e.Get(1)
					_ = V3981
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1profile), V3979, V3980, V3981)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7205 := PrimCons(symshen_4type_1signature_1of_1profile, tmp7204)

		tmp7206 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7205)

		_ = tmp7206

		tmp7207 := MakeNative(func(__e *ControlFlow) {
			V3989 := __e.Get(1)
			_ = V3989
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V3990 := __e.Get(1)
				_ = V3990
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V3991 := __e.Get(1)
					_ = V3991
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1preclude), V3989, V3990, V3991)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7208 := PrimCons(symshen_4type_1signature_1of_1preclude, tmp7207)

		tmp7209 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7208)

		_ = tmp7209

		tmp7210 := MakeNative(func(__e *ControlFlow) {
			V3999 := __e.Get(1)
			_ = V3999
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4000 := __e.Get(1)
				_ = V4000
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4001 := __e.Get(1)
					_ = V4001
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1shen_4proc_1nl), V3999, V4000, V4001)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7211 := PrimCons(symshen_4type_1signature_1of_1shen_4proc_1nl, tmp7210)

		tmp7212 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7211)

		_ = tmp7212

		tmp7213 := MakeNative(func(__e *ControlFlow) {
			V4009 := __e.Get(1)
			_ = V4009
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4010 := __e.Get(1)
				_ = V4010
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4011 := __e.Get(1)
					_ = V4011
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1profile_1results), V4009, V4010, V4011)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7214 := PrimCons(symshen_4type_1signature_1of_1profile_1results, tmp7213)

		tmp7215 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7214)

		_ = tmp7215

		tmp7216 := MakeNative(func(__e *ControlFlow) {
			V4019 := __e.Get(1)
			_ = V4019
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4020 := __e.Get(1)
				_ = V4020
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4021 := __e.Get(1)
					_ = V4021
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1protect), V4019, V4020, V4021)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7217 := PrimCons(symshen_4type_1signature_1of_1protect, tmp7216)

		tmp7218 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7217)

		_ = tmp7218

		tmp7219 := MakeNative(func(__e *ControlFlow) {
			V4029 := __e.Get(1)
			_ = V4029
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4030 := __e.Get(1)
				_ = V4030
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4031 := __e.Get(1)
					_ = V4031
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1preclude_1all_1but), V4029, V4030, V4031)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7220 := PrimCons(symshen_4type_1signature_1of_1preclude_1all_1but, tmp7219)

		tmp7221 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7220)

		_ = tmp7221

		tmp7222 := MakeNative(func(__e *ControlFlow) {
			V4039 := __e.Get(1)
			_ = V4039
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4040 := __e.Get(1)
				_ = V4040
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4041 := __e.Get(1)
					_ = V4041
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1shen_4prhush), V4039, V4040, V4041)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7223 := PrimCons(symshen_4type_1signature_1of_1shen_4prhush, tmp7222)

		tmp7224 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7223)

		_ = tmp7224

		tmp7225 := MakeNative(func(__e *ControlFlow) {
			V4049 := __e.Get(1)
			_ = V4049
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4050 := __e.Get(1)
				_ = V4050
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4051 := __e.Get(1)
					_ = V4051
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1ps), V4049, V4050, V4051)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7226 := PrimCons(symshen_4type_1signature_1of_1ps, tmp7225)

		tmp7227 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7226)

		_ = tmp7227

		tmp7228 := MakeNative(func(__e *ControlFlow) {
			V4059 := __e.Get(1)
			_ = V4059
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4060 := __e.Get(1)
				_ = V4060
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4061 := __e.Get(1)
					_ = V4061
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1read), V4059, V4060, V4061)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7229 := PrimCons(symshen_4type_1signature_1of_1read, tmp7228)

		tmp7230 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7229)

		_ = tmp7230

		tmp7231 := MakeNative(func(__e *ControlFlow) {
			V4069 := __e.Get(1)
			_ = V4069
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4070 := __e.Get(1)
				_ = V4070
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4071 := __e.Get(1)
					_ = V4071
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1read_1byte), V4069, V4070, V4071)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7232 := PrimCons(symshen_4type_1signature_1of_1read_1byte, tmp7231)

		tmp7233 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7232)

		_ = tmp7233

		tmp7234 := MakeNative(func(__e *ControlFlow) {
			V4079 := __e.Get(1)
			_ = V4079
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4080 := __e.Get(1)
				_ = V4080
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4081 := __e.Get(1)
					_ = V4081
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1read_1file_1as_1bytelist), V4079, V4080, V4081)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7235 := PrimCons(symshen_4type_1signature_1of_1read_1file_1as_1bytelist, tmp7234)

		tmp7236 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7235)

		_ = tmp7236

		tmp7237 := MakeNative(func(__e *ControlFlow) {
			V4089 := __e.Get(1)
			_ = V4089
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4090 := __e.Get(1)
				_ = V4090
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4091 := __e.Get(1)
					_ = V4091
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1read_1file_1as_1string), V4089, V4090, V4091)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7238 := PrimCons(symshen_4type_1signature_1of_1read_1file_1as_1string, tmp7237)

		tmp7239 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7238)

		_ = tmp7239

		tmp7240 := MakeNative(func(__e *ControlFlow) {
			V4099 := __e.Get(1)
			_ = V4099
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4100 := __e.Get(1)
				_ = V4100
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4101 := __e.Get(1)
					_ = V4101
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1read_1file), V4099, V4100, V4101)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7241 := PrimCons(symshen_4type_1signature_1of_1read_1file, tmp7240)

		tmp7242 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7241)

		_ = tmp7242

		tmp7243 := MakeNative(func(__e *ControlFlow) {
			V4109 := __e.Get(1)
			_ = V4109
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4110 := __e.Get(1)
				_ = V4110
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4111 := __e.Get(1)
					_ = V4111
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1read_1from_1string), V4109, V4110, V4111)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7244 := PrimCons(symshen_4type_1signature_1of_1read_1from_1string, tmp7243)

		tmp7245 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7244)

		_ = tmp7245

		tmp7246 := MakeNative(func(__e *ControlFlow) {
			V4119 := __e.Get(1)
			_ = V4119
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4120 := __e.Get(1)
				_ = V4120
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4121 := __e.Get(1)
					_ = V4121
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1release), V4119, V4120, V4121)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7247 := PrimCons(symshen_4type_1signature_1of_1release, tmp7246)

		tmp7248 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7247)

		_ = tmp7248

		tmp7249 := MakeNative(func(__e *ControlFlow) {
			V4129 := __e.Get(1)
			_ = V4129
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4130 := __e.Get(1)
				_ = V4130
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4131 := __e.Get(1)
					_ = V4131
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1remove), V4129, V4130, V4131)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7250 := PrimCons(symshen_4type_1signature_1of_1remove, tmp7249)

		tmp7251 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7250)

		_ = tmp7251

		tmp7252 := MakeNative(func(__e *ControlFlow) {
			V4139 := __e.Get(1)
			_ = V4139
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4140 := __e.Get(1)
				_ = V4140
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4141 := __e.Get(1)
					_ = V4141
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1reverse), V4139, V4140, V4141)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7253 := PrimCons(symshen_4type_1signature_1of_1reverse, tmp7252)

		tmp7254 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7253)

		_ = tmp7254

		tmp7255 := MakeNative(func(__e *ControlFlow) {
			V4149 := __e.Get(1)
			_ = V4149
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4150 := __e.Get(1)
				_ = V4150
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4151 := __e.Get(1)
					_ = V4151
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1simple_1error), V4149, V4150, V4151)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7256 := PrimCons(symshen_4type_1signature_1of_1simple_1error, tmp7255)

		tmp7257 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7256)

		_ = tmp7257

		tmp7258 := MakeNative(func(__e *ControlFlow) {
			V4159 := __e.Get(1)
			_ = V4159
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4160 := __e.Get(1)
				_ = V4160
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4161 := __e.Get(1)
					_ = V4161
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1snd), V4159, V4160, V4161)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7259 := PrimCons(symshen_4type_1signature_1of_1snd, tmp7258)

		tmp7260 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7259)

		_ = tmp7260

		tmp7261 := MakeNative(func(__e *ControlFlow) {
			V4169 := __e.Get(1)
			_ = V4169
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4170 := __e.Get(1)
				_ = V4170
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4171 := __e.Get(1)
					_ = V4171
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1specialise), V4169, V4170, V4171)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7262 := PrimCons(symshen_4type_1signature_1of_1specialise, tmp7261)

		tmp7263 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7262)

		_ = tmp7263

		tmp7264 := MakeNative(func(__e *ControlFlow) {
			V4179 := __e.Get(1)
			_ = V4179
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4180 := __e.Get(1)
				_ = V4180
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4181 := __e.Get(1)
					_ = V4181
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1spy), V4179, V4180, V4181)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7265 := PrimCons(symshen_4type_1signature_1of_1spy, tmp7264)

		tmp7266 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7265)

		_ = tmp7266

		tmp7267 := MakeNative(func(__e *ControlFlow) {
			V4189 := __e.Get(1)
			_ = V4189
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4190 := __e.Get(1)
				_ = V4190
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4191 := __e.Get(1)
					_ = V4191
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1step), V4189, V4190, V4191)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7268 := PrimCons(symshen_4type_1signature_1of_1step, tmp7267)

		tmp7269 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7268)

		_ = tmp7269

		tmp7270 := MakeNative(func(__e *ControlFlow) {
			V4199 := __e.Get(1)
			_ = V4199
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4200 := __e.Get(1)
				_ = V4200
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4201 := __e.Get(1)
					_ = V4201
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1stinput), V4199, V4200, V4201)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7271 := PrimCons(symshen_4type_1signature_1of_1stinput, tmp7270)

		tmp7272 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7271)

		_ = tmp7272

		tmp7273 := MakeNative(func(__e *ControlFlow) {
			V4209 := __e.Get(1)
			_ = V4209
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4210 := __e.Get(1)
				_ = V4210
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4211 := __e.Get(1)
					_ = V4211
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1sterror), V4209, V4210, V4211)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7274 := PrimCons(symshen_4type_1signature_1of_1sterror, tmp7273)

		tmp7275 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7274)

		_ = tmp7275

		tmp7276 := MakeNative(func(__e *ControlFlow) {
			V4219 := __e.Get(1)
			_ = V4219
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4220 := __e.Get(1)
				_ = V4220
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4221 := __e.Get(1)
					_ = V4221
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1stoutput), V4219, V4220, V4221)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7277 := PrimCons(symshen_4type_1signature_1of_1stoutput, tmp7276)

		tmp7278 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7277)

		_ = tmp7278

		tmp7279 := MakeNative(func(__e *ControlFlow) {
			V4229 := __e.Get(1)
			_ = V4229
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4230 := __e.Get(1)
				_ = V4230
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4231 := __e.Get(1)
					_ = V4231
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1string_2), V4229, V4230, V4231)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7280 := PrimCons(symshen_4type_1signature_1of_1string_2, tmp7279)

		tmp7281 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7280)

		_ = tmp7281

		tmp7282 := MakeNative(func(__e *ControlFlow) {
			V4239 := __e.Get(1)
			_ = V4239
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4240 := __e.Get(1)
				_ = V4240
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4241 := __e.Get(1)
					_ = V4241
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1str), V4239, V4240, V4241)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7283 := PrimCons(symshen_4type_1signature_1of_1str, tmp7282)

		tmp7284 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7283)

		_ = tmp7284

		tmp7285 := MakeNative(func(__e *ControlFlow) {
			V4249 := __e.Get(1)
			_ = V4249
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4250 := __e.Get(1)
				_ = V4250
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4251 := __e.Get(1)
					_ = V4251
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1string_1_6n), V4249, V4250, V4251)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7286 := PrimCons(symshen_4type_1signature_1of_1string_1_6n, tmp7285)

		tmp7287 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7286)

		_ = tmp7287

		tmp7288 := MakeNative(func(__e *ControlFlow) {
			V4259 := __e.Get(1)
			_ = V4259
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4260 := __e.Get(1)
				_ = V4260
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4261 := __e.Get(1)
					_ = V4261
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1string_1_6symbol), V4259, V4260, V4261)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7289 := PrimCons(symshen_4type_1signature_1of_1string_1_6symbol, tmp7288)

		tmp7290 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7289)

		_ = tmp7290

		tmp7291 := MakeNative(func(__e *ControlFlow) {
			V4269 := __e.Get(1)
			_ = V4269
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4270 := __e.Get(1)
				_ = V4270
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4271 := __e.Get(1)
					_ = V4271
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1sum), V4269, V4270, V4271)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7292 := PrimCons(symshen_4type_1signature_1of_1sum, tmp7291)

		tmp7293 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7292)

		_ = tmp7293

		tmp7294 := MakeNative(func(__e *ControlFlow) {
			V4279 := __e.Get(1)
			_ = V4279
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4280 := __e.Get(1)
				_ = V4280
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4281 := __e.Get(1)
					_ = V4281
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1symbol_2), V4279, V4280, V4281)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7295 := PrimCons(symshen_4type_1signature_1of_1symbol_2, tmp7294)

		tmp7296 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7295)

		_ = tmp7296

		tmp7297 := MakeNative(func(__e *ControlFlow) {
			V4289 := __e.Get(1)
			_ = V4289
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4290 := __e.Get(1)
				_ = V4290
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4291 := __e.Get(1)
					_ = V4291
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1systemf), V4289, V4290, V4291)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7298 := PrimCons(symshen_4type_1signature_1of_1systemf, tmp7297)

		tmp7299 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7298)

		_ = tmp7299

		tmp7300 := MakeNative(func(__e *ControlFlow) {
			V4299 := __e.Get(1)
			_ = V4299
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4300 := __e.Get(1)
				_ = V4300
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4301 := __e.Get(1)
					_ = V4301
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1tail), V4299, V4300, V4301)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7301 := PrimCons(symshen_4type_1signature_1of_1tail, tmp7300)

		tmp7302 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7301)

		_ = tmp7302

		tmp7303 := MakeNative(func(__e *ControlFlow) {
			V4309 := __e.Get(1)
			_ = V4309
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4310 := __e.Get(1)
				_ = V4310
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4311 := __e.Get(1)
					_ = V4311
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1tlstr), V4309, V4310, V4311)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7304 := PrimCons(symshen_4type_1signature_1of_1tlstr, tmp7303)

		tmp7305 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7304)

		_ = tmp7305

		tmp7306 := MakeNative(func(__e *ControlFlow) {
			V4319 := __e.Get(1)
			_ = V4319
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4320 := __e.Get(1)
				_ = V4320
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4321 := __e.Get(1)
					_ = V4321
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1tlv), V4319, V4320, V4321)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7307 := PrimCons(symshen_4type_1signature_1of_1tlv, tmp7306)

		tmp7308 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7307)

		_ = tmp7308

		tmp7309 := MakeNative(func(__e *ControlFlow) {
			V4329 := __e.Get(1)
			_ = V4329
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4330 := __e.Get(1)
				_ = V4330
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4331 := __e.Get(1)
					_ = V4331
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1tc), V4329, V4330, V4331)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7310 := PrimCons(symshen_4type_1signature_1of_1tc, tmp7309)

		tmp7311 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7310)

		_ = tmp7311

		tmp7312 := MakeNative(func(__e *ControlFlow) {
			V4339 := __e.Get(1)
			_ = V4339
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4340 := __e.Get(1)
				_ = V4340
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4341 := __e.Get(1)
					_ = V4341
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1tc_2), V4339, V4340, V4341)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7313 := PrimCons(symshen_4type_1signature_1of_1tc_2, tmp7312)

		tmp7314 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7313)

		_ = tmp7314

		tmp7315 := MakeNative(func(__e *ControlFlow) {
			V4349 := __e.Get(1)
			_ = V4349
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4350 := __e.Get(1)
				_ = V4350
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4351 := __e.Get(1)
					_ = V4351
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1thaw), V4349, V4350, V4351)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7316 := PrimCons(symshen_4type_1signature_1of_1thaw, tmp7315)

		tmp7317 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7316)

		_ = tmp7317

		tmp7318 := MakeNative(func(__e *ControlFlow) {
			V4359 := __e.Get(1)
			_ = V4359
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4360 := __e.Get(1)
				_ = V4360
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4361 := __e.Get(1)
					_ = V4361
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1track), V4359, V4360, V4361)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7319 := PrimCons(symshen_4type_1signature_1of_1track, tmp7318)

		tmp7320 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7319)

		_ = tmp7320

		tmp7321 := MakeNative(func(__e *ControlFlow) {
			V4369 := __e.Get(1)
			_ = V4369
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4370 := __e.Get(1)
				_ = V4370
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4371 := __e.Get(1)
					_ = V4371
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1trap_1error), V4369, V4370, V4371)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7322 := PrimCons(symshen_4type_1signature_1of_1trap_1error, tmp7321)

		tmp7323 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7322)

		_ = tmp7323

		tmp7324 := MakeNative(func(__e *ControlFlow) {
			V4379 := __e.Get(1)
			_ = V4379
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4380 := __e.Get(1)
				_ = V4380
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4381 := __e.Get(1)
					_ = V4381
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1tuple_2), V4379, V4380, V4381)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7325 := PrimCons(symshen_4type_1signature_1of_1tuple_2, tmp7324)

		tmp7326 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7325)

		_ = tmp7326

		tmp7327 := MakeNative(func(__e *ControlFlow) {
			V4389 := __e.Get(1)
			_ = V4389
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4390 := __e.Get(1)
				_ = V4390
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4391 := __e.Get(1)
					_ = V4391
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1undefmacro), V4389, V4390, V4391)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7328 := PrimCons(symshen_4type_1signature_1of_1undefmacro, tmp7327)

		tmp7329 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7328)

		_ = tmp7329

		tmp7330 := MakeNative(func(__e *ControlFlow) {
			V4399 := __e.Get(1)
			_ = V4399
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4400 := __e.Get(1)
				_ = V4400
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4401 := __e.Get(1)
					_ = V4401
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1union), V4399, V4400, V4401)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7331 := PrimCons(symshen_4type_1signature_1of_1union, tmp7330)

		tmp7332 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7331)

		_ = tmp7332

		tmp7333 := MakeNative(func(__e *ControlFlow) {
			V4409 := __e.Get(1)
			_ = V4409
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4410 := __e.Get(1)
				_ = V4410
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4411 := __e.Get(1)
					_ = V4411
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1unprofile), V4409, V4410, V4411)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7334 := PrimCons(symshen_4type_1signature_1of_1unprofile, tmp7333)

		tmp7335 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7334)

		_ = tmp7335

		tmp7336 := MakeNative(func(__e *ControlFlow) {
			V4419 := __e.Get(1)
			_ = V4419
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4420 := __e.Get(1)
				_ = V4420
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4421 := __e.Get(1)
					_ = V4421
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1untrack), V4419, V4420, V4421)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7337 := PrimCons(symshen_4type_1signature_1of_1untrack, tmp7336)

		tmp7338 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7337)

		_ = tmp7338

		tmp7339 := MakeNative(func(__e *ControlFlow) {
			V4429 := __e.Get(1)
			_ = V4429
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4430 := __e.Get(1)
				_ = V4430
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4431 := __e.Get(1)
					_ = V4431
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1unspecialise), V4429, V4430, V4431)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7340 := PrimCons(symshen_4type_1signature_1of_1unspecialise, tmp7339)

		tmp7341 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7340)

		_ = tmp7341

		tmp7342 := MakeNative(func(__e *ControlFlow) {
			V4439 := __e.Get(1)
			_ = V4439
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4440 := __e.Get(1)
				_ = V4440
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4441 := __e.Get(1)
					_ = V4441
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1variable_2), V4439, V4440, V4441)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7343 := PrimCons(symshen_4type_1signature_1of_1variable_2, tmp7342)

		tmp7344 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7343)

		_ = tmp7344

		tmp7345 := MakeNative(func(__e *ControlFlow) {
			V4449 := __e.Get(1)
			_ = V4449
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4450 := __e.Get(1)
				_ = V4450
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4451 := __e.Get(1)
					_ = V4451
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1vector_2), V4449, V4450, V4451)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7346 := PrimCons(symshen_4type_1signature_1of_1vector_2, tmp7345)

		tmp7347 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7346)

		_ = tmp7347

		tmp7348 := MakeNative(func(__e *ControlFlow) {
			V4459 := __e.Get(1)
			_ = V4459
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4460 := __e.Get(1)
				_ = V4460
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4461 := __e.Get(1)
					_ = V4461
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1version), V4459, V4460, V4461)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7349 := PrimCons(symshen_4type_1signature_1of_1version, tmp7348)

		tmp7350 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7349)

		_ = tmp7350

		tmp7351 := MakeNative(func(__e *ControlFlow) {
			V4469 := __e.Get(1)
			_ = V4469
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4470 := __e.Get(1)
				_ = V4470
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4471 := __e.Get(1)
					_ = V4471
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1write_1to_1file), V4469, V4470, V4471)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7352 := PrimCons(symshen_4type_1signature_1of_1write_1to_1file, tmp7351)

		tmp7353 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7352)

		_ = tmp7353

		tmp7354 := MakeNative(func(__e *ControlFlow) {
			V4479 := __e.Get(1)
			_ = V4479
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4480 := __e.Get(1)
				_ = V4480
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4481 := __e.Get(1)
					_ = V4481
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1write_1byte), V4479, V4480, V4481)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7355 := PrimCons(symshen_4type_1signature_1of_1write_1byte, tmp7354)

		tmp7356 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7355)

		_ = tmp7356

		tmp7357 := MakeNative(func(__e *ControlFlow) {
			V4489 := __e.Get(1)
			_ = V4489
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4490 := __e.Get(1)
				_ = V4490
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4491 := __e.Get(1)
					_ = V4491
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1y_1or_1n_2), V4489, V4490, V4491)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7358 := PrimCons(symshen_4type_1signature_1of_1y_1or_1n_2, tmp7357)

		tmp7359 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7358)

		_ = tmp7359

		tmp7360 := MakeNative(func(__e *ControlFlow) {
			V4499 := __e.Get(1)
			_ = V4499
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4500 := __e.Get(1)
				_ = V4500
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4501 := __e.Get(1)
					_ = V4501
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_6), V4499, V4500, V4501)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7361 := PrimCons(symshen_4type_1signature_1of_1_6, tmp7360)

		tmp7362 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7361)

		_ = tmp7362

		tmp7363 := MakeNative(func(__e *ControlFlow) {
			V4509 := __e.Get(1)
			_ = V4509
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4510 := __e.Get(1)
				_ = V4510
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4511 := __e.Get(1)
					_ = V4511
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_5), V4509, V4510, V4511)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7364 := PrimCons(symshen_4type_1signature_1of_1_5, tmp7363)

		tmp7365 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7364)

		_ = tmp7365

		tmp7366 := MakeNative(func(__e *ControlFlow) {
			V4519 := __e.Get(1)
			_ = V4519
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4520 := __e.Get(1)
				_ = V4520
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4521 := __e.Get(1)
					_ = V4521
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_6_a), V4519, V4520, V4521)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7367 := PrimCons(symshen_4type_1signature_1of_1_6_a, tmp7366)

		tmp7368 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7367)

		_ = tmp7368

		tmp7369 := MakeNative(func(__e *ControlFlow) {
			V4529 := __e.Get(1)
			_ = V4529
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4530 := __e.Get(1)
				_ = V4530
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4531 := __e.Get(1)
					_ = V4531
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_5_a), V4529, V4530, V4531)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7370 := PrimCons(symshen_4type_1signature_1of_1_5_a, tmp7369)

		tmp7371 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7370)

		_ = tmp7371

		tmp7372 := MakeNative(func(__e *ControlFlow) {
			V4539 := __e.Get(1)
			_ = V4539
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4540 := __e.Get(1)
				_ = V4540
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4541 := __e.Get(1)
					_ = V4541
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_a), V4539, V4540, V4541)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7373 := PrimCons(symshen_4type_1signature_1of_1_a, tmp7372)

		tmp7374 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7373)

		_ = tmp7374

		tmp7375 := MakeNative(func(__e *ControlFlow) {
			V4549 := __e.Get(1)
			_ = V4549
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4550 := __e.Get(1)
				_ = V4550
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4551 := __e.Get(1)
					_ = V4551
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_7), V4549, V4550, V4551)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7376 := PrimCons(symshen_4type_1signature_1of_1_7, tmp7375)

		tmp7377 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7376)

		_ = tmp7377

		tmp7378 := MakeNative(func(__e *ControlFlow) {
			V4559 := __e.Get(1)
			_ = V4559
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4560 := __e.Get(1)
				_ = V4560
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4561 := __e.Get(1)
					_ = V4561
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_c), V4559, V4560, V4561)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7379 := PrimCons(symshen_4type_1signature_1of_1_c, tmp7378)

		tmp7380 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7379)

		_ = tmp7380

		tmp7381 := MakeNative(func(__e *ControlFlow) {
			V4569 := __e.Get(1)
			_ = V4569
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4570 := __e.Get(1)
				_ = V4570
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4571 := __e.Get(1)
					_ = V4571
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_1), V4569, V4570, V4571)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7382 := PrimCons(symshen_4type_1signature_1of_1_1, tmp7381)

		tmp7383 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7382)

		_ = tmp7383

		tmp7384 := MakeNative(func(__e *ControlFlow) {
			V4579 := __e.Get(1)
			_ = V4579
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4580 := __e.Get(1)
				_ = V4580
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4581 := __e.Get(1)
					_ = V4581
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_d), V4579, V4580, V4581)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7385 := PrimCons(symshen_4type_1signature_1of_1_d, tmp7384)

		tmp7386 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7385)

		_ = tmp7386

		tmp7387 := MakeNative(func(__e *ControlFlow) {
			V4589 := __e.Get(1)
			_ = V4589
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V4590 := __e.Get(1)
				_ = V4590
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V4591 := __e.Get(1)
					_ = V4591
					__e.TailApply(PrimNS2Value(symshen_4type_1signature_1of_1_a_a), V4589, V4590, V4591)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7388 := PrimCons(symshen_4type_1signature_1of_1_a_a, tmp7387)

		__e.TailApply(PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7388)
		return

	}, 0)

	tmp7389 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise_1signedfunc_1lambda_1forms, tmp6981)

	_ = tmp7389

	tmp7390 := MakeNative(func(__e *ControlFlow) {
		tmp7391 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4datatype_1error), X)
			return
		}, 1)

		tmp7392 := PrimCons(symshen_4datatype_1error, tmp7391)

		tmp7393 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7392)

		_ = tmp7393

		tmp7394 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4tuple), X)
			return
		}, 1)

		tmp7395 := PrimCons(symshen_4tuple, tmp7394)

		tmp7396 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7395)

		_ = tmp7396

		tmp7397 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4pvar), X)
			return
		}, 1)

		tmp7398 := PrimCons(symshen_4pvar, tmp7397)

		tmp7399 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7398)

		_ = tmp7399

		tmp7400 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4dictionary), X)
			return
		}, 1)

		tmp7401 := PrimCons(symshen_4dictionary, tmp7400)

		tmp7402 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7401)

		_ = tmp7402

		tmp7403 := MakeNative(func(__e *ControlFlow) {
			V476 := __e.Get(1)
			_ = V476
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V477 := __e.Get(1)
				_ = V477
				__e.TailApply(PrimNS2Value(sym_8v), V476, V477)
				return
			}, 1))
			return
		}, 1)

		tmp7404 := PrimCons(sym_8v, tmp7403)

		tmp7405 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7404)

		_ = tmp7405

		tmp7406 := MakeNative(func(__e *ControlFlow) {
			V478 := __e.Get(1)
			_ = V478
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V479 := __e.Get(1)
				_ = V479
				__e.TailApply(PrimNS2Value(sym_8p), V478, V479)
				return
			}, 1))
			return
		}, 1)

		tmp7407 := PrimCons(sym_8p, tmp7406)

		tmp7408 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7407)

		_ = tmp7408

		tmp7409 := MakeNative(func(__e *ControlFlow) {
			V480 := __e.Get(1)
			_ = V480
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V481 := __e.Get(1)
				_ = V481
				__e.TailApply(PrimNS2Value(sym_8s), V480, V481)
				return
			}, 1))
			return
		}, 1)

		tmp7410 := PrimCons(sym_8s, tmp7409)

		tmp7411 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7410)

		_ = tmp7411

		tmp7412 := MakeNative(func(__e *ControlFlow) {
			V482 := __e.Get(1)
			_ = V482
			__e.TailApply(PrimNS2Value(sym_5e_6), V482)
			return
		}, 1)

		tmp7413 := PrimCons(sym_5e_6, tmp7412)

		tmp7414 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7413)

		_ = tmp7414

		tmp7415 := MakeNative(func(__e *ControlFlow) {
			V483 := __e.Get(1)
			_ = V483
			__e.TailApply(PrimNS2Value(sym_5_b_6), V483)
			return
		}, 1)

		tmp7416 := PrimCons(sym_5_b_6, tmp7415)

		tmp7417 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7416)

		_ = tmp7417

		tmp7418 := MakeNative(func(__e *ControlFlow) {
			V484 := __e.Get(1)
			_ = V484
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V485 := __e.Get(1)
				_ = V485
				__e.TailApply(PrimNS2Value(sym_a_a), V484, V485)
				return
			}, 1))
			return
		}, 1)

		tmp7419 := PrimCons(sym_a_a, tmp7418)

		tmp7420 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7419)

		_ = tmp7420

		tmp7421 := MakeNative(func(__e *ControlFlow) {
			V486 := __e.Get(1)
			_ = V486
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V487 := __e.Get(1)
				_ = V487
				__e.Return(PrimEqual(V486, V487))
				return
			}, 1))
			return
		}, 1)

		tmp7422 := PrimCons(sym_a, tmp7421)

		tmp7423 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7422)

		_ = tmp7423

		tmp7424 := MakeNative(func(__e *ControlFlow) {
			V488 := __e.Get(1)
			_ = V488
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V489 := __e.Get(1)
				_ = V489
				__e.Return(PrimGreatEqual(V488, V489))
				return
			}, 1))
			return
		}, 1)

		tmp7425 := PrimCons(sym_6_a, tmp7424)

		tmp7426 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7425)

		_ = tmp7426

		tmp7427 := MakeNative(func(__e *ControlFlow) {
			V490 := __e.Get(1)
			_ = V490
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V491 := __e.Get(1)
				_ = V491
				__e.Return(PrimGreatThan(V490, V491))
				return
			}, 1))
			return
		}, 1)

		tmp7428 := PrimCons(sym_6, tmp7427)

		tmp7429 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7428)

		_ = tmp7429

		tmp7430 := MakeNative(func(__e *ControlFlow) {
			V492 := __e.Get(1)
			_ = V492
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V493 := __e.Get(1)
				_ = V493
				__e.Return(PrimNumberSubtract(V492, V493))
				return
			}, 1))
			return
		}, 1)

		tmp7431 := PrimCons(sym_1, tmp7430)

		tmp7432 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7431)

		_ = tmp7432

		tmp7433 := MakeNative(func(__e *ControlFlow) {
			V494 := __e.Get(1)
			_ = V494
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V495 := __e.Get(1)
				_ = V495
				__e.Return(PrimNumberDivide(V494, V495))
				return
			}, 1))
			return
		}, 1)

		tmp7434 := PrimCons(sym_c, tmp7433)

		tmp7435 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7434)

		_ = tmp7435

		tmp7436 := MakeNative(func(__e *ControlFlow) {
			V496 := __e.Get(1)
			_ = V496
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V497 := __e.Get(1)
				_ = V497
				__e.Return(PrimNumberMultiply(V496, V497))
				return
			}, 1))
			return
		}, 1)

		tmp7437 := PrimCons(sym_d, tmp7436)

		tmp7438 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7437)

		_ = tmp7438

		tmp7439 := MakeNative(func(__e *ControlFlow) {
			V498 := __e.Get(1)
			_ = V498
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V499 := __e.Get(1)
				_ = V499
				__e.Return(PrimNumberAdd(V498, V499))
				return
			}, 1))
			return
		}, 1)

		tmp7440 := PrimCons(sym_7, tmp7439)

		tmp7441 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7440)

		_ = tmp7441

		tmp7442 := MakeNative(func(__e *ControlFlow) {
			V500 := __e.Get(1)
			_ = V500
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V501 := __e.Get(1)
				_ = V501
				__e.Return(PrimLessEqual(V500, V501))
				return
			}, 1))
			return
		}, 1)

		tmp7443 := PrimCons(sym_5_a, tmp7442)

		tmp7444 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7443)

		_ = tmp7444

		tmp7445 := MakeNative(func(__e *ControlFlow) {
			V502 := __e.Get(1)
			_ = V502
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V503 := __e.Get(1)
				_ = V503
				__e.Return(PrimLessThan(V502, V503))
				return
			}, 1))
			return
		}, 1)

		tmp7446 := PrimCons(sym_5, tmp7445)

		tmp7447 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7446)

		_ = tmp7447

		tmp7448 := MakeNative(func(__e *ControlFlow) {
			V504 := __e.Get(1)
			_ = V504
			__e.TailApply(PrimNS2Value(symy_1or_1n_2), V504)
			return
		}, 1)

		tmp7449 := PrimCons(symy_1or_1n_2, tmp7448)

		tmp7450 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7449)

		_ = tmp7450

		tmp7451 := MakeNative(func(__e *ControlFlow) {
			V505 := __e.Get(1)
			_ = V505
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V506 := __e.Get(1)
				_ = V506
				__e.TailApply(PrimNS2Value(symwrite_1to_1file), V505, V506)
				return
			}, 1))
			return
		}, 1)

		tmp7452 := PrimCons(symwrite_1to_1file, tmp7451)

		tmp7453 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7452)

		_ = tmp7453

		tmp7454 := MakeNative(func(__e *ControlFlow) {
			V507 := __e.Get(1)
			_ = V507
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V508 := __e.Get(1)
				_ = V508
				__e.Return(PrimWriteByte(V507, V508))
				return
			}, 1))
			return
		}, 1)

		tmp7455 := PrimCons(symwrite_1byte, tmp7454)

		tmp7456 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7455)

		_ = tmp7456

		tmp7457 := MakeNative(func(__e *ControlFlow) {
			V509 := __e.Get(1)
			_ = V509
			__e.Return(PrimIsVariable(V509))
			return
		}, 1)

		tmp7458 := PrimCons(symvariable_2, tmp7457)

		tmp7459 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7458)

		_ = tmp7459

		tmp7460 := MakeNative(func(__e *ControlFlow) {
			V510 := __e.Get(1)
			_ = V510
			__e.Return(PrimNS3Value(V510))
			return
		}, 1)

		tmp7461 := PrimCons(symvalue, tmp7460)

		tmp7462 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7461)

		_ = tmp7462

		tmp7463 := MakeNative(func(__e *ControlFlow) {
			V511 := __e.Get(1)
			_ = V511
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V512 := __e.Get(1)
				_ = V512
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V513 := __e.Get(1)
					_ = V513
					__e.TailApply(PrimNS2Value(symvector_1_6), V511, V512, V513)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7464 := PrimCons(symvector_1_6, tmp7463)

		tmp7465 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7464)

		_ = tmp7465

		tmp7466 := MakeNative(func(__e *ControlFlow) {
			V514 := __e.Get(1)
			_ = V514
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V515 := __e.Get(1)
				_ = V515
				__e.TailApply(PrimNS2Value(sym_5_1vector), V514, V515)
				return
			}, 1))
			return
		}, 1)

		tmp7467 := PrimCons(sym_5_1vector, tmp7466)

		tmp7468 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7467)

		_ = tmp7468

		tmp7469 := MakeNative(func(__e *ControlFlow) {
			V516 := __e.Get(1)
			_ = V516
			__e.TailApply(PrimNS2Value(symvector), V516)
			return
		}, 1)

		tmp7470 := PrimCons(symvector, tmp7469)

		tmp7471 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7470)

		_ = tmp7471

		tmp7472 := MakeNative(func(__e *ControlFlow) {
			V517 := __e.Get(1)
			_ = V517
			__e.TailApply(PrimNS2Value(symvector_2), V517)
			return
		}, 1)

		tmp7473 := PrimCons(symvector_2, tmp7472)

		tmp7474 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7473)

		_ = tmp7474

		tmp7475 := MakeNative(func(__e *ControlFlow) {
			V518 := __e.Get(1)
			_ = V518
			__e.TailApply(PrimNS2Value(symunspecialise), V518)
			return
		}, 1)

		tmp7476 := PrimCons(symunspecialise, tmp7475)

		tmp7477 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7476)

		_ = tmp7477

		tmp7478 := MakeNative(func(__e *ControlFlow) {
			V519 := __e.Get(1)
			_ = V519
			__e.TailApply(PrimNS2Value(symuntrack), V519)
			return
		}, 1)

		tmp7479 := PrimCons(symuntrack, tmp7478)

		tmp7480 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7479)

		_ = tmp7480

		tmp7481 := MakeNative(func(__e *ControlFlow) {
			V520 := __e.Get(1)
			_ = V520
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V521 := __e.Get(1)
				_ = V521
				__e.TailApply(PrimNS2Value(symunion), V520, V521)
				return
			}, 1))
			return
		}, 1)

		tmp7482 := PrimCons(symunion, tmp7481)

		tmp7483 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7482)

		_ = tmp7483

		tmp7484 := MakeNative(func(__e *ControlFlow) {
			V522 := __e.Get(1)
			_ = V522
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V523 := __e.Get(1)
				_ = V523
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V524 := __e.Get(1)
					_ = V524
					__e.Return(MakeNative(func(__e *ControlFlow) {
						V525 := __e.Get(1)
						_ = V525
						__e.TailApply(PrimNS2Value(symunify), V522, V523, V524, V525)
						return
					}, 1))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7485 := PrimCons(symunify, tmp7484)

		tmp7486 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7485)

		_ = tmp7486

		tmp7487 := MakeNative(func(__e *ControlFlow) {
			V526 := __e.Get(1)
			_ = V526
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V527 := __e.Get(1)
				_ = V527
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V528 := __e.Get(1)
					_ = V528
					__e.Return(MakeNative(func(__e *ControlFlow) {
						V529 := __e.Get(1)
						_ = V529
						__e.TailApply(PrimNS2Value(symunify_b), V526, V527, V528, V529)
						return
					}, 1))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7488 := PrimCons(symunify_b, tmp7487)

		tmp7489 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7488)

		_ = tmp7489

		tmp7490 := MakeNative(func(__e *ControlFlow) {
			V530 := __e.Get(1)
			_ = V530
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V531 := __e.Get(1)
				_ = V531
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V532 := __e.Get(1)
					_ = V532
					__e.TailApply(PrimNS2Value(symunput), V530, V531, V532)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7491 := PrimCons(symunput, tmp7490)

		tmp7492 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7491)

		_ = tmp7492

		tmp7493 := MakeNative(func(__e *ControlFlow) {
			V533 := __e.Get(1)
			_ = V533
			__e.TailApply(PrimNS2Value(symunprofile), V533)
			return
		}, 1)

		tmp7494 := PrimCons(symunprofile, tmp7493)

		tmp7495 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7494)

		_ = tmp7495

		tmp7496 := MakeNative(func(__e *ControlFlow) {
			V534 := __e.Get(1)
			_ = V534
			__e.TailApply(PrimNS2Value(symundefmacro), V534)
			return
		}, 1)

		tmp7497 := PrimCons(symundefmacro, tmp7496)

		tmp7498 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7497)

		_ = tmp7498

		tmp7499 := MakeNative(func(__e *ControlFlow) {
			V535 := __e.Get(1)
			_ = V535
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V536 := __e.Get(1)
				_ = V536
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V537 := __e.Get(1)
					_ = V537
					__e.TailApply(PrimNS2Value(symreturn), V535, V536, V537)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7500 := PrimCons(symreturn, tmp7499)

		tmp7501 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7500)

		_ = tmp7501

		tmp7502 := MakeNative(func(__e *ControlFlow) {
			V538 := __e.Get(1)
			_ = V538
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V539 := __e.Get(1)
				_ = V539
				__e.TailApply(PrimNS2Value(symtype), V538, V539)
				return
			}, 1))
			return
		}, 1)

		tmp7503 := PrimCons(symtype, tmp7502)

		tmp7504 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7503)

		_ = tmp7504

		tmp7505 := MakeNative(func(__e *ControlFlow) {
			V540 := __e.Get(1)
			_ = V540
			__e.TailApply(PrimNS2Value(symtuple_2), V540)
			return
		}, 1)

		tmp7506 := PrimCons(symtuple_2, tmp7505)

		tmp7507 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7506)

		_ = tmp7507

		tmp7508 := MakeNative(func(__e *ControlFlow) {
			V541 := __e.Get(1)
			_ = V541
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V542 := __e.Get(1)
				_ = V542
				tmp7509 := MakeNative(func(__e *ControlFlow) {
					__e.Return(V541)
					return
				}, 0)

				__e.TailApply(PrimNS1Value(symtry_1catch), tmp7509, V542)
				return

			}, 1))
			return
		}, 1)

		tmp7510 := PrimCons(symtrap_1error, tmp7508)

		tmp7511 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7510)

		_ = tmp7511

		tmp7512 := MakeNative(func(__e *ControlFlow) {
			V543 := __e.Get(1)
			_ = V543
			__e.TailApply(PrimNS2Value(symtrack), V543)
			return
		}, 1)

		tmp7513 := PrimCons(symtrack, tmp7512)

		tmp7514 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7513)

		_ = tmp7514

		tmp7515 := MakeNative(func(__e *ControlFlow) {
			V544 := __e.Get(1)
			_ = V544
			__e.TailApply(PrimNS2Value(symthaw), V544)
			return
		}, 1)

		tmp7516 := PrimCons(symthaw, tmp7515)

		tmp7517 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7516)

		_ = tmp7517

		tmp7518 := MakeNative(func(__e *ControlFlow) {
			V545 := __e.Get(1)
			_ = V545
			__e.TailApply(PrimNS2Value(symtc), V545)
			return
		}, 1)

		tmp7519 := PrimCons(symtc, tmp7518)

		tmp7520 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7519)

		_ = tmp7520

		tmp7521 := MakeNative(func(__e *ControlFlow) {
			V546 := __e.Get(1)
			_ = V546
			__e.Return(PrimTail(V546))
			return
		}, 1)

		tmp7522 := PrimCons(symtl, tmp7521)

		tmp7523 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7522)

		_ = tmp7523

		tmp7524 := MakeNative(func(__e *ControlFlow) {
			V547 := __e.Get(1)
			_ = V547
			__e.Return(PrimTailString(V547))
			return
		}, 1)

		tmp7525 := PrimCons(symtlstr, tmp7524)

		tmp7526 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7525)

		_ = tmp7526

		tmp7527 := MakeNative(func(__e *ControlFlow) {
			V548 := __e.Get(1)
			_ = V548
			__e.TailApply(PrimNS2Value(symtail), V548)
			return
		}, 1)

		tmp7528 := PrimCons(symtail, tmp7527)

		tmp7529 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7528)

		_ = tmp7529

		tmp7530 := MakeNative(func(__e *ControlFlow) {
			V549 := __e.Get(1)
			_ = V549
			__e.TailApply(PrimNS2Value(symsystemf), V549)
			return
		}, 1)

		tmp7531 := PrimCons(symsystemf, tmp7530)

		tmp7532 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7531)

		_ = tmp7532

		tmp7533 := MakeNative(func(__e *ControlFlow) {
			V550 := __e.Get(1)
			_ = V550
			__e.Return(PrimIsSymbol(V550))
			return
		}, 1)

		tmp7534 := PrimCons(symsymbol_2, tmp7533)

		tmp7535 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7534)

		_ = tmp7535

		tmp7536 := MakeNative(func(__e *ControlFlow) {
			V551 := __e.Get(1)
			_ = V551
			__e.TailApply(PrimNS2Value(symstring_1_6symbol), V551)
			return
		}, 1)

		tmp7537 := PrimCons(symstring_1_6symbol, tmp7536)

		tmp7538 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7537)

		_ = tmp7538

		tmp7539 := MakeNative(func(__e *ControlFlow) {
			V552 := __e.Get(1)
			_ = V552
			__e.TailApply(PrimNS2Value(symsum), V552)
			return
		}, 1)

		tmp7540 := PrimCons(symsum, tmp7539)

		tmp7541 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7540)

		_ = tmp7541

		tmp7542 := MakeNative(func(__e *ControlFlow) {
			V553 := __e.Get(1)
			_ = V553
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V554 := __e.Get(1)
				_ = V554
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V555 := __e.Get(1)
					_ = V555
					__e.TailApply(PrimNS2Value(symsubst), V553, V554, V555)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7543 := PrimCons(symsubst, tmp7542)

		tmp7544 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7543)

		_ = tmp7544

		tmp7545 := MakeNative(func(__e *ControlFlow) {
			V556 := __e.Get(1)
			_ = V556
			__e.Return(PrimIsString(V556))
			return
		}, 1)

		tmp7546 := PrimCons(symstring_2, tmp7545)

		tmp7547 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7546)

		_ = tmp7547

		tmp7548 := MakeNative(func(__e *ControlFlow) {
			V557 := __e.Get(1)
			_ = V557
			__e.Return(PrimStringToNumber(V557))
			return
		}, 1)

		tmp7549 := PrimCons(symstring_1_6n, tmp7548)

		tmp7550 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7549)

		_ = tmp7550

		tmp7551 := MakeNative(func(__e *ControlFlow) {
			V558 := __e.Get(1)
			_ = V558
			__e.TailApply(PrimNS2Value(symstep), V558)
			return
		}, 1)

		tmp7552 := PrimCons(symstep, tmp7551)

		tmp7553 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7552)

		_ = tmp7553

		tmp7554 := MakeNative(func(__e *ControlFlow) {
			V559 := __e.Get(1)
			_ = V559
			__e.TailApply(PrimNS2Value(symspy), V559)
			return
		}, 1)

		tmp7555 := PrimCons(symspy, tmp7554)

		tmp7556 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7555)

		_ = tmp7556

		tmp7557 := MakeNative(func(__e *ControlFlow) {
			V560 := __e.Get(1)
			_ = V560
			__e.TailApply(PrimNS2Value(symspecialise), V560)
			return
		}, 1)

		tmp7558 := PrimCons(symspecialise, tmp7557)

		tmp7559 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7558)

		_ = tmp7559

		tmp7560 := MakeNative(func(__e *ControlFlow) {
			V561 := __e.Get(1)
			_ = V561
			__e.TailApply(PrimNS2Value(symsnd), V561)
			return
		}, 1)

		tmp7561 := PrimCons(symsnd, tmp7560)

		tmp7562 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7561)

		_ = tmp7562

		tmp7563 := MakeNative(func(__e *ControlFlow) {
			V562 := __e.Get(1)
			_ = V562
			__e.Return(PrimSimpleError(V562))
			return
		}, 1)

		tmp7564 := PrimCons(symsimple_1error, tmp7563)

		tmp7565 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7564)

		_ = tmp7565

		tmp7566 := MakeNative(func(__e *ControlFlow) {
			V563 := __e.Get(1)
			_ = V563
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V564 := __e.Get(1)
				_ = V564
				__e.Return(PrimNS3Set(V563, V564))
				return
			}, 1))
			return
		}, 1)

		tmp7567 := PrimCons(symset, tmp7566)

		tmp7568 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7567)

		_ = tmp7568

		tmp7569 := MakeNative(func(__e *ControlFlow) {
			V565 := __e.Get(1)
			_ = V565
			__e.Return(PrimStr(V565))
			return
		}, 1)

		tmp7570 := PrimCons(symstr, tmp7569)

		tmp7571 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7570)

		_ = tmp7571

		tmp7572 := MakeNative(func(__e *ControlFlow) {
			V566 := __e.Get(1)
			_ = V566
			__e.TailApply(PrimNS2Value(symreverse), V566)
			return
		}, 1)

		tmp7573 := PrimCons(symreverse, tmp7572)

		tmp7574 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7573)

		_ = tmp7574

		tmp7575 := MakeNative(func(__e *ControlFlow) {
			V567 := __e.Get(1)
			_ = V567
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V568 := __e.Get(1)
				_ = V568
				__e.TailApply(PrimNS2Value(symremove), V567, V568)
				return
			}, 1))
			return
		}, 1)

		tmp7576 := PrimCons(symremove, tmp7575)

		tmp7577 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7576)

		_ = tmp7577

		tmp7578 := MakeNative(func(__e *ControlFlow) {
			V569 := __e.Get(1)
			_ = V569
			__e.TailApply(PrimNS2Value(symread), V569)
			return
		}, 1)

		tmp7579 := PrimCons(symread, tmp7578)

		tmp7580 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7579)

		_ = tmp7580

		tmp7581 := MakeNative(func(__e *ControlFlow) {
			V570 := __e.Get(1)
			_ = V570
			__e.TailApply(PrimNS2Value(symread_1file), V570)
			return
		}, 1)

		tmp7582 := PrimCons(symread_1file, tmp7581)

		tmp7583 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7582)

		_ = tmp7583

		tmp7584 := MakeNative(func(__e *ControlFlow) {
			V571 := __e.Get(1)
			_ = V571
			__e.Return(PrimReadFileAsByteList(V571))
			return
		}, 1)

		tmp7585 := PrimCons(symread_1file_1as_1bytelist, tmp7584)

		tmp7586 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7585)

		_ = tmp7586

		tmp7587 := MakeNative(func(__e *ControlFlow) {
			V572 := __e.Get(1)
			_ = V572
			__e.Return(PrimReadFileAsString(V572))
			return
		}, 1)

		tmp7588 := PrimCons(symread_1file_1as_1string, tmp7587)

		tmp7589 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7588)

		_ = tmp7589

		tmp7590 := MakeNative(func(__e *ControlFlow) {
			V573 := __e.Get(1)
			_ = V573
			__e.Return(PrimReadByte(V573))
			return
		}, 1)

		tmp7591 := PrimCons(symread_1byte, tmp7590)

		tmp7592 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7591)

		_ = tmp7592

		tmp7593 := MakeNative(func(__e *ControlFlow) {
			V574 := __e.Get(1)
			_ = V574
			__e.TailApply(PrimNS2Value(symread_1from_1string), V574)
			return
		}, 1)

		tmp7594 := PrimCons(symread_1from_1string, tmp7593)

		tmp7595 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7594)

		_ = tmp7595

		tmp7596 := MakeNative(func(__e *ControlFlow) {
			V575 := __e.Get(1)
			_ = V575
			__e.TailApply(PrimNS2Value(sympackage_2), V575)
			return
		}, 1)

		tmp7597 := PrimCons(sympackage_2, tmp7596)

		tmp7598 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7597)

		_ = tmp7598

		tmp7599 := MakeNative(func(__e *ControlFlow) {
			V576 := __e.Get(1)
			_ = V576
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V577 := __e.Get(1)
				_ = V577
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V578 := __e.Get(1)
					_ = V578
					__e.Return(MakeNative(func(__e *ControlFlow) {
						V579 := __e.Get(1)
						_ = V579
						__e.TailApply(PrimNS2Value(symput), V576, V577, V578, V579)
						return
					}, 1))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7600 := PrimCons(symput, tmp7599)

		tmp7601 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7600)

		_ = tmp7601

		tmp7602 := MakeNative(func(__e *ControlFlow) {
			V580 := __e.Get(1)
			_ = V580
			__e.TailApply(PrimNS2Value(sympreclude), V580)
			return
		}, 1)

		tmp7603 := PrimCons(sympreclude, tmp7602)

		tmp7604 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7603)

		_ = tmp7604

		tmp7605 := MakeNative(func(__e *ControlFlow) {
			V581 := __e.Get(1)
			_ = V581
			__e.TailApply(PrimNS2Value(sympreclude_1all_1but), V581)
			return
		}, 1)

		tmp7606 := PrimCons(sympreclude_1all_1but, tmp7605)

		tmp7607 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7606)

		_ = tmp7607

		tmp7608 := MakeNative(func(__e *ControlFlow) {
			V582 := __e.Get(1)
			_ = V582
			__e.TailApply(PrimNS2Value(symps), V582)
			return
		}, 1)

		tmp7609 := PrimCons(symps, tmp7608)

		tmp7610 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7609)

		_ = tmp7610

		tmp7611 := MakeNative(func(__e *ControlFlow) {
			V583 := __e.Get(1)
			_ = V583
			__e.TailApply(PrimNS2Value(symprotect), V583)
			return
		}, 1)

		tmp7612 := PrimCons(symprotect, tmp7611)

		tmp7613 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7612)

		_ = tmp7613

		tmp7614 := MakeNative(func(__e *ControlFlow) {
			V584 := __e.Get(1)
			_ = V584
			__e.TailApply(PrimNS2Value(symprofile_1results), V584)
			return
		}, 1)

		tmp7615 := PrimCons(symprofile_1results, tmp7614)

		tmp7616 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7615)

		_ = tmp7616

		tmp7617 := MakeNative(func(__e *ControlFlow) {
			V585 := __e.Get(1)
			_ = V585
			__e.TailApply(PrimNS2Value(symprofile), V585)
			return
		}, 1)

		tmp7618 := PrimCons(symprofile, tmp7617)

		tmp7619 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7618)

		_ = tmp7619

		tmp7620 := MakeNative(func(__e *ControlFlow) {
			V586 := __e.Get(1)
			_ = V586
			__e.TailApply(PrimNS2Value(symprint), V586)
			return
		}, 1)

		tmp7621 := PrimCons(symprint, tmp7620)

		tmp7622 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7621)

		_ = tmp7622

		tmp7623 := MakeNative(func(__e *ControlFlow) {
			V587 := __e.Get(1)
			_ = V587
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V588 := __e.Get(1)
				_ = V588
				__e.TailApply(PrimNS2Value(sympr), V587, V588)
				return
			}, 1))
			return
		}, 1)

		tmp7624 := PrimCons(sympr, tmp7623)

		tmp7625 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7624)

		_ = tmp7625

		tmp7626 := MakeNative(func(__e *ControlFlow) {
			V589 := __e.Get(1)
			_ = V589
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V590 := __e.Get(1)
				_ = V590
				__e.Return(PrimPos(V589, V590))
				return
			}, 1))
			return
		}, 1)

		tmp7627 := PrimCons(sympos, tmp7626)

		tmp7628 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7627)

		_ = tmp7628

		tmp7629 := MakeNative(func(__e *ControlFlow) {
			V591 := __e.Get(1)
			_ = V591
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V592 := __e.Get(1)
				_ = V592
				if True == V591 {
					__e.Return(True)
					return
				} else {
					if True == V592 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}
				}
			}, 1))
			return
		}, 1)

		tmp7632 := PrimCons(symor, tmp7629)

		tmp7633 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7632)

		_ = tmp7633

		tmp7634 := MakeNative(func(__e *ControlFlow) {
			V593 := __e.Get(1)
			_ = V593
			__e.TailApply(PrimNS2Value(symoptimise), V593)
			return
		}, 1)

		tmp7635 := PrimCons(symoptimise, tmp7634)

		tmp7636 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7635)

		_ = tmp7636

		tmp7637 := MakeNative(func(__e *ControlFlow) {
			V594 := __e.Get(1)
			_ = V594
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V595 := __e.Get(1)
				_ = V595
				__e.Return(PrimOpenStream(V594, V595))
				return
			}, 1))
			return
		}, 1)

		tmp7638 := PrimCons(symopen, tmp7637)

		tmp7639 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7638)

		_ = tmp7639

		tmp7640 := MakeNative(func(__e *ControlFlow) {
			V596 := __e.Get(1)
			_ = V596
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V597 := __e.Get(1)
				_ = V597
				__e.TailApply(PrimNS2Value(symoccurrences), V596, V597)
				return
			}, 1))
			return
		}, 1)

		tmp7641 := PrimCons(symoccurrences, tmp7640)

		tmp7642 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7641)

		_ = tmp7642

		tmp7643 := MakeNative(func(__e *ControlFlow) {
			V598 := __e.Get(1)
			_ = V598
			__e.TailApply(PrimNS2Value(symoccurs_1check), V598)
			return
		}, 1)

		tmp7644 := PrimCons(symoccurs_1check, tmp7643)

		tmp7645 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7644)

		_ = tmp7645

		tmp7646 := MakeNative(func(__e *ControlFlow) {
			V599 := __e.Get(1)
			_ = V599
			__e.Return(PrimNumberToString(V599))
			return
		}, 1)

		tmp7647 := PrimCons(symn_1_6string, tmp7646)

		tmp7648 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7647)

		_ = tmp7648

		tmp7649 := MakeNative(func(__e *ControlFlow) {
			V600 := __e.Get(1)
			_ = V600
			__e.Return(PrimIsNumber(V600))
			return
		}, 1)

		tmp7650 := PrimCons(symnumber_2, tmp7649)

		tmp7651 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7650)

		_ = tmp7651

		tmp7652 := MakeNative(func(__e *ControlFlow) {
			V601 := __e.Get(1)
			_ = V601
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V602 := __e.Get(1)
				_ = V602
				__e.TailApply(PrimNS2Value(symnth), V601, V602)
				return
			}, 1))
			return
		}, 1)

		tmp7653 := PrimCons(symnth, tmp7652)

		tmp7654 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7653)

		_ = tmp7654

		tmp7655 := MakeNative(func(__e *ControlFlow) {
			V603 := __e.Get(1)
			_ = V603
			__e.Return(PrimNot(V603))
			return
		}, 1)

		tmp7656 := PrimCons(symnot, tmp7655)

		tmp7657 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7656)

		_ = tmp7657

		tmp7658 := MakeNative(func(__e *ControlFlow) {
			V604 := __e.Get(1)
			_ = V604
			__e.TailApply(PrimNS2Value(symnl), V604)
			return
		}, 1)

		tmp7659 := PrimCons(symnl, tmp7658)

		tmp7660 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7659)

		_ = tmp7660

		tmp7661 := MakeNative(func(__e *ControlFlow) {
			V605 := __e.Get(1)
			_ = V605
			__e.TailApply(PrimNS2Value(symmacroexpand), V605)
			return
		}, 1)

		tmp7662 := PrimCons(symmacroexpand, tmp7661)

		tmp7663 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7662)

		_ = tmp7663

		tmp7664 := MakeNative(func(__e *ControlFlow) {
			V606 := __e.Get(1)
			_ = V606
			__e.TailApply(PrimNS2Value(symmaxinferences), V606)
			return
		}, 1)

		tmp7665 := PrimCons(symmaxinferences, tmp7664)

		tmp7666 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7665)

		_ = tmp7666

		tmp7667 := MakeNative(func(__e *ControlFlow) {
			V607 := __e.Get(1)
			_ = V607
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V608 := __e.Get(1)
				_ = V608
				__e.TailApply(PrimNS2Value(symmapcan), V607, V608)
				return
			}, 1))
			return
		}, 1)

		tmp7668 := PrimCons(symmapcan, tmp7667)

		tmp7669 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7668)

		_ = tmp7669

		tmp7670 := MakeNative(func(__e *ControlFlow) {
			V609 := __e.Get(1)
			_ = V609
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V610 := __e.Get(1)
				_ = V610
				__e.TailApply(PrimNS2Value(symmap), V609, V610)
				return
			}, 1))
			return
		}, 1)

		tmp7671 := PrimCons(symmap, tmp7670)

		tmp7672 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7671)

		_ = tmp7672

		tmp7673 := MakeNative(func(__e *ControlFlow) {
			V611 := __e.Get(1)
			_ = V611
			__e.TailApply(PrimNS2Value(symload), V611)
			return
		}, 1)

		tmp7674 := PrimCons(symload, tmp7673)

		tmp7675 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7674)

		_ = tmp7675

		tmp7676 := MakeNative(func(__e *ControlFlow) {
			V612 := __e.Get(1)
			_ = V612
			__e.TailApply(PrimNS2Value(symlineread), V612)
			return
		}, 1)

		tmp7677 := PrimCons(symlineread, tmp7676)

		tmp7678 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7677)

		_ = tmp7678

		tmp7679 := MakeNative(func(__e *ControlFlow) {
			V613 := __e.Get(1)
			_ = V613
			__e.TailApply(PrimNS2Value(symlimit), V613)
			return
		}, 1)

		tmp7680 := PrimCons(symlimit, tmp7679)

		tmp7681 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7680)

		_ = tmp7681

		tmp7682 := MakeNative(func(__e *ControlFlow) {
			V614 := __e.Get(1)
			_ = V614
			__e.TailApply(PrimNS2Value(symlength), V614)
			return
		}, 1)

		tmp7683 := PrimCons(symlength, tmp7682)

		tmp7684 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7683)

		_ = tmp7684

		tmp7685 := MakeNative(func(__e *ControlFlow) {
			V615 := __e.Get(1)
			_ = V615
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V616 := __e.Get(1)
				_ = V616
				__e.TailApply(PrimNS2Value(symintersection), V615, V616)
				return
			}, 1))
			return
		}, 1)

		tmp7686 := PrimCons(symintersection, tmp7685)

		tmp7687 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7686)

		_ = tmp7687

		tmp7688 := MakeNative(func(__e *ControlFlow) {
			V617 := __e.Get(1)
			_ = V617
			__e.Return(PrimIntern(V617))
			return
		}, 1)

		tmp7689 := PrimCons(symintern, tmp7688)

		tmp7690 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7689)

		_ = tmp7690

		tmp7691 := MakeNative(func(__e *ControlFlow) {
			V618 := __e.Get(1)
			_ = V618
			__e.Return(PrimIsInteger(V618))
			return
		}, 1)

		tmp7692 := PrimCons(syminteger_2, tmp7691)

		tmp7693 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7692)

		_ = tmp7693

		tmp7694 := MakeNative(func(__e *ControlFlow) {
			V619 := __e.Get(1)
			_ = V619
			__e.TailApply(PrimNS2Value(syminput), V619)
			return
		}, 1)

		tmp7695 := PrimCons(syminput, tmp7694)

		tmp7696 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7695)

		_ = tmp7696

		tmp7697 := MakeNative(func(__e *ControlFlow) {
			V620 := __e.Get(1)
			_ = V620
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V621 := __e.Get(1)
				_ = V621
				__e.TailApply(PrimNS2Value(syminput_7), V620, V621)
				return
			}, 1))
			return
		}, 1)

		tmp7698 := PrimCons(syminput_7, tmp7697)

		tmp7699 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7698)

		_ = tmp7699

		tmp7700 := MakeNative(func(__e *ControlFlow) {
			V622 := __e.Get(1)
			_ = V622
			__e.TailApply(PrimNS2Value(syminclude), V622)
			return
		}, 1)

		tmp7701 := PrimCons(syminclude, tmp7700)

		tmp7702 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7701)

		_ = tmp7702

		tmp7703 := MakeNative(func(__e *ControlFlow) {
			V623 := __e.Get(1)
			_ = V623
			__e.TailApply(PrimNS2Value(syminclude_1all_1but), V623)
			return
		}, 1)

		tmp7704 := PrimCons(syminclude_1all_1but, tmp7703)

		tmp7705 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7704)

		_ = tmp7705

		tmp7706 := MakeNative(func(__e *ControlFlow) {
			V624 := __e.Get(1)
			_ = V624
			__e.TailApply(PrimNS2Value(syminternal), V624)
			return
		}, 1)

		tmp7707 := PrimCons(syminternal, tmp7706)

		tmp7708 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7707)

		_ = tmp7708

		tmp7709 := MakeNative(func(__e *ControlFlow) {
			V625 := __e.Get(1)
			_ = V625
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V626 := __e.Get(1)
				_ = V626
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V627 := __e.Get(1)
					_ = V627
					if True == V625 {
						__e.Return(V626)
						return
					} else {
						__e.Return(V627)
						return
					}
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7711 := PrimCons(symif, tmp7709)

		tmp7712 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7711)

		_ = tmp7712

		tmp7713 := MakeNative(func(__e *ControlFlow) {
			V628 := __e.Get(1)
			_ = V628
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V629 := __e.Get(1)
				_ = V629
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V630 := __e.Get(1)
					_ = V630
					__e.Return(MakeNative(func(__e *ControlFlow) {
						V631 := __e.Get(1)
						_ = V631
						__e.TailApply(PrimNS2Value(symidentical), V628, V629, V630, V631)
						return
					}, 1))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7714 := PrimCons(symidentical, tmp7713)

		tmp7715 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7714)

		_ = tmp7715

		tmp7716 := MakeNative(func(__e *ControlFlow) {
			V632 := __e.Get(1)
			_ = V632
			__e.TailApply(PrimNS2Value(symhead), V632)
			return
		}, 1)

		tmp7717 := PrimCons(symhead, tmp7716)

		tmp7718 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7717)

		_ = tmp7718

		tmp7719 := MakeNative(func(__e *ControlFlow) {
			V633 := __e.Get(1)
			_ = V633
			__e.Return(PrimHead(V633))
			return
		}, 1)

		tmp7720 := PrimCons(symhd, tmp7719)

		tmp7721 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7720)

		_ = tmp7721

		tmp7722 := MakeNative(func(__e *ControlFlow) {
			V634 := __e.Get(1)
			_ = V634
			__e.TailApply(PrimNS2Value(symhdv), V634)
			return
		}, 1)

		tmp7723 := PrimCons(symhdv, tmp7722)

		tmp7724 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7723)

		_ = tmp7724

		tmp7725 := MakeNative(func(__e *ControlFlow) {
			V635 := __e.Get(1)
			_ = V635
			__e.TailApply(PrimNS2Value(symhdstr), V635)
			return
		}, 1)

		tmp7726 := PrimCons(symhdstr, tmp7725)

		tmp7727 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7726)

		_ = tmp7727

		tmp7728 := MakeNative(func(__e *ControlFlow) {
			V636 := __e.Get(1)
			_ = V636
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V637 := __e.Get(1)
				_ = V637
				__e.TailApply(PrimNS2Value(symhash), V636, V637)
				return
			}, 1))
			return
		}, 1)

		tmp7729 := PrimCons(symhash, tmp7728)

		tmp7730 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7729)

		_ = tmp7730

		tmp7731 := MakeNative(func(__e *ControlFlow) {
			V638 := __e.Get(1)
			_ = V638
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V639 := __e.Get(1)
				_ = V639
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V640 := __e.Get(1)
					_ = V640
					__e.TailApply(PrimNS2Value(symget), V638, V639, V640)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7732 := PrimCons(symget, tmp7731)

		tmp7733 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7732)

		_ = tmp7733

		tmp7734 := MakeNative(func(__e *ControlFlow) {
			V641 := __e.Get(1)
			_ = V641
			__e.Return(PrimGetTime(V641))
			return
		}, 1)

		tmp7735 := PrimCons(symget_1time, tmp7734)

		tmp7736 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7735)

		_ = tmp7736

		tmp7737 := MakeNative(func(__e *ControlFlow) {
			V642 := __e.Get(1)
			_ = V642
			__e.TailApply(PrimNS2Value(symgensym), V642)
			return
		}, 1)

		tmp7738 := PrimCons(symgensym, tmp7737)

		tmp7739 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7738)

		_ = tmp7739

		tmp7740 := MakeNative(func(__e *ControlFlow) {
			V643 := __e.Get(1)
			_ = V643
			__e.TailApply(PrimNS2Value(symfst), V643)
			return
		}, 1)

		tmp7741 := PrimCons(symfst, tmp7740)

		tmp7742 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7741)

		_ = tmp7742

		tmp7743 := MakeNative(func(__e *ControlFlow) {
			V644 := __e.Get(1)
			_ = V644
			__e.Return(MakeNative(func(__e *ControlFlow) {
				__e.Return(V644)
				return
			}, 0))
			return
		}, 1)

		tmp7744 := PrimCons(symfreeze, tmp7743)

		tmp7745 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7744)

		_ = tmp7745

		tmp7746 := MakeNative(func(__e *ControlFlow) {
			V645 := __e.Get(1)
			_ = V645
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V646 := __e.Get(1)
				_ = V646
				__e.TailApply(PrimNS2Value(symfix), V645, V646)
				return
			}, 1))
			return
		}, 1)

		tmp7747 := PrimCons(symfix, tmp7746)

		tmp7748 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7747)

		_ = tmp7748

		tmp7749 := MakeNative(func(__e *ControlFlow) {
			V647 := __e.Get(1)
			_ = V647
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V648 := __e.Get(1)
				_ = V648
				__e.TailApply(PrimNS2Value(symfail_1if), V647, V648)
				return
			}, 1))
			return
		}, 1)

		tmp7750 := PrimCons(symfail_1if, tmp7749)

		tmp7751 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7750)

		_ = tmp7751

		tmp7752 := MakeNative(func(__e *ControlFlow) {
			V649 := __e.Get(1)
			_ = V649
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V650 := __e.Get(1)
				_ = V650
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V651 := __e.Get(1)
					_ = V651
					__e.Return(MakeNative(func(__e *ControlFlow) {
						V652 := __e.Get(1)
						_ = V652
						__e.Return(MakeNative(func(__e *ControlFlow) {
							V653 := __e.Get(1)
							_ = V653
							__e.TailApply(PrimNS2Value(symfindall), V649, V650, V651, V652, V653)
							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7753 := PrimCons(symfindall, tmp7752)

		tmp7754 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7753)

		_ = tmp7754

		tmp7755 := MakeNative(func(__e *ControlFlow) {
			V654 := __e.Get(1)
			_ = V654
			__e.TailApply(PrimNS2Value(symenable_1type_1theory), V654)
			return
		}, 1)

		tmp7756 := PrimCons(symenable_1type_1theory, tmp7755)

		tmp7757 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7756)

		_ = tmp7757

		tmp7758 := MakeNative(func(__e *ControlFlow) {
			V655 := __e.Get(1)
			_ = V655
			__e.TailApply(PrimNS2Value(symexplode), V655)
			return
		}, 1)

		tmp7759 := PrimCons(symexplode, tmp7758)

		tmp7760 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7759)

		_ = tmp7760

		tmp7761 := MakeNative(func(__e *ControlFlow) {
			V656 := __e.Get(1)
			_ = V656
			__e.TailApply(PrimNS2Value(symexternal), V656)
			return
		}, 1)

		tmp7762 := PrimCons(symexternal, tmp7761)

		tmp7763 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7762)

		_ = tmp7763

		tmp7764 := MakeNative(func(__e *ControlFlow) {
			V657 := __e.Get(1)
			_ = V657
			__e.TailApply(PrimNS2Value(symeval_1kl), V657)
			return
		}, 1)

		tmp7765 := PrimCons(symeval_1kl, tmp7764)

		tmp7766 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7765)

		_ = tmp7766

		tmp7767 := MakeNative(func(__e *ControlFlow) {
			V658 := __e.Get(1)
			_ = V658
			__e.TailApply(PrimNS2Value(symeval), V658)
			return
		}, 1)

		tmp7768 := PrimCons(symeval, tmp7767)

		tmp7769 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7768)

		_ = tmp7769

		tmp7770 := MakeNative(func(__e *ControlFlow) {
			V659 := __e.Get(1)
			_ = V659
			__e.Return(PrimErrorToString(V659))
			return
		}, 1)

		tmp7771 := PrimCons(symerror_1to_1string, tmp7770)

		tmp7772 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7771)

		_ = tmp7772

		tmp7773 := MakeNative(func(__e *ControlFlow) {
			V660 := __e.Get(1)
			_ = V660
			__e.TailApply(PrimNS2Value(symempty_2), V660)
			return
		}, 1)

		tmp7774 := PrimCons(symempty_2, tmp7773)

		tmp7775 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7774)

		_ = tmp7775

		tmp7776 := MakeNative(func(__e *ControlFlow) {
			V661 := __e.Get(1)
			_ = V661
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V662 := __e.Get(1)
				_ = V662
				__e.TailApply(PrimNS2Value(symelement_2), V661, V662)
				return
			}, 1))
			return
		}, 1)

		tmp7777 := PrimCons(symelement_2, tmp7776)

		tmp7778 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7777)

		_ = tmp7778

		tmp7779 := MakeNative(func(__e *ControlFlow) {
			V663 := __e.Get(1)
			_ = V663
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V664 := __e.Get(1)
				_ = V664
				_ = V663

				__e.Return(V664)
				return

			}, 1))
			return
		}, 1)

		tmp7780 := PrimCons(symdo, tmp7779)

		tmp7781 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7780)

		_ = tmp7781

		tmp7782 := MakeNative(func(__e *ControlFlow) {
			V665 := __e.Get(1)
			_ = V665
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V666 := __e.Get(1)
				_ = V666
				__e.TailApply(PrimNS2Value(symdifference), V665, V666)
				return
			}, 1))
			return
		}, 1)

		tmp7783 := PrimCons(symdifference, tmp7782)

		tmp7784 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7783)

		_ = tmp7784

		tmp7785 := MakeNative(func(__e *ControlFlow) {
			V667 := __e.Get(1)
			_ = V667
			__e.TailApply(PrimNS2Value(symdestroy), V667)
			return
		}, 1)

		tmp7786 := PrimCons(symdestroy, tmp7785)

		tmp7787 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7786)

		_ = tmp7787

		tmp7788 := MakeNative(func(__e *ControlFlow) {
			V668 := __e.Get(1)
			_ = V668
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V669 := __e.Get(1)
				_ = V669
				__e.TailApply(PrimNS2Value(symdeclare), V668, V669)
				return
			}, 1))
			return
		}, 1)

		tmp7789 := PrimCons(symdeclare, tmp7788)

		tmp7790 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7789)

		_ = tmp7790

		tmp7791 := MakeNative(func(__e *ControlFlow) {
			V670 := __e.Get(1)
			_ = V670
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V671 := __e.Get(1)
				_ = V671
				__e.Return(PrimStringConcat(V670, V671))
				return
			}, 1))
			return
		}, 1)

		tmp7792 := PrimCons(symcn, tmp7791)

		tmp7793 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7792)

		_ = tmp7793

		tmp7794 := MakeNative(func(__e *ControlFlow) {
			V672 := __e.Get(1)
			_ = V672
			__e.Return(PrimIsPair(V672))
			return
		}, 1)

		tmp7795 := PrimCons(symcons_2, tmp7794)

		tmp7796 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7795)

		_ = tmp7796

		tmp7797 := MakeNative(func(__e *ControlFlow) {
			V673 := __e.Get(1)
			_ = V673
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V674 := __e.Get(1)
				_ = V674
				__e.Return(PrimCons(V673, V674))
				return
			}, 1))
			return
		}, 1)

		tmp7798 := PrimCons(symcons, tmp7797)

		tmp7799 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7798)

		_ = tmp7799

		tmp7800 := MakeNative(func(__e *ControlFlow) {
			V675 := __e.Get(1)
			_ = V675
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V676 := __e.Get(1)
				_ = V676
				__e.TailApply(PrimNS2Value(symconcat), V675, V676)
				return
			}, 1))
			return
		}, 1)

		tmp7801 := PrimCons(symconcat, tmp7800)

		tmp7802 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7801)

		_ = tmp7802

		tmp7803 := MakeNative(func(__e *ControlFlow) {
			V677 := __e.Get(1)
			_ = V677
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V678 := __e.Get(1)
				_ = V678
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V679 := __e.Get(1)
					_ = V679
					__e.TailApply(PrimNS2Value(symcompile), V677, V678, V679)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7804 := PrimCons(symcompile, tmp7803)

		tmp7805 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7804)

		_ = tmp7805

		tmp7806 := MakeNative(func(__e *ControlFlow) {
			V680 := __e.Get(1)
			_ = V680
			__e.TailApply(PrimNS2Value(symcd), V680)
			return
		}, 1)

		tmp7807 := PrimCons(symcd, tmp7806)

		tmp7808 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7807)

		_ = tmp7808

		tmp7809 := MakeNative(func(__e *ControlFlow) {
			V681 := __e.Get(1)
			_ = V681
			__e.Return(PrimCloseStream(V681))
			return
		}, 1)

		tmp7810 := PrimCons(symclose, tmp7809)

		tmp7811 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7810)

		_ = tmp7811

		tmp7812 := MakeNative(func(__e *ControlFlow) {
			V682 := __e.Get(1)
			_ = V682
			__e.TailApply(PrimNS2Value(symbound_2), V682)
			return
		}, 1)

		tmp7813 := PrimCons(symbound_2, tmp7812)

		tmp7814 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7813)

		_ = tmp7814

		tmp7815 := MakeNative(func(__e *ControlFlow) {
			V683 := __e.Get(1)
			_ = V683
			__e.TailApply(PrimNS2Value(symboolean_2), V683)
			return
		}, 1)

		tmp7816 := PrimCons(symboolean_2, tmp7815)

		tmp7817 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7816)

		_ = tmp7817

		tmp7818 := MakeNative(func(__e *ControlFlow) {
			V684 := __e.Get(1)
			_ = V684
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V685 := __e.Get(1)
				_ = V685
				__e.TailApply(PrimNS2Value(symassoc), V684, V685)
				return
			}, 1))
			return
		}, 1)

		tmp7819 := PrimCons(symassoc, tmp7818)

		tmp7820 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7819)

		_ = tmp7820

		tmp7821 := MakeNative(func(__e *ControlFlow) {
			V686 := __e.Get(1)
			_ = V686
			__e.TailApply(PrimNS2Value(symarity), V686)
			return
		}, 1)

		tmp7822 := PrimCons(symarity, tmp7821)

		tmp7823 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7822)

		_ = tmp7823

		tmp7824 := MakeNative(func(__e *ControlFlow) {
			V687 := __e.Get(1)
			_ = V687
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V688 := __e.Get(1)
				_ = V688
				__e.TailApply(PrimNS2Value(symappend), V687, V688)
				return
			}, 1))
			return
		}, 1)

		tmp7825 := PrimCons(symappend, tmp7824)

		tmp7826 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7825)

		_ = tmp7826

		tmp7827 := MakeNative(func(__e *ControlFlow) {
			V689 := __e.Get(1)
			_ = V689
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V690 := __e.Get(1)
				_ = V690
				if True == V689 {
					if True == V690 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}
				} else {
					__e.Return(False)
					return
				}
			}, 1))
			return
		}, 1)

		tmp7830 := PrimCons(symand, tmp7827)

		tmp7831 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7830)

		_ = tmp7831

		tmp7832 := MakeNative(func(__e *ControlFlow) {
			V691 := __e.Get(1)
			_ = V691
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V692 := __e.Get(1)
				_ = V692
				__e.TailApply(PrimNS2Value(symadjoin), V691, V692)
				return
			}, 1))
			return
		}, 1)

		tmp7833 := PrimCons(symadjoin, tmp7832)

		tmp7834 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7833)

		_ = tmp7834

		tmp7835 := MakeNative(func(__e *ControlFlow) {
			V693 := __e.Get(1)
			_ = V693
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V694 := __e.Get(1)
				_ = V694
				__e.Return(PrimVectorGet(V693, V694))
				return
			}, 1))
			return
		}, 1)

		tmp7836 := PrimCons(sym_5_1address, tmp7835)

		tmp7837 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7836)

		_ = tmp7837

		tmp7838 := MakeNative(func(__e *ControlFlow) {
			V695 := __e.Get(1)
			_ = V695
			__e.Return(MakeNative(func(__e *ControlFlow) {
				V696 := __e.Get(1)
				_ = V696
				__e.Return(MakeNative(func(__e *ControlFlow) {
					V697 := __e.Get(1)
					_ = V697
					__e.Return(PrimVectorSet(V695, V696, V697))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp7839 := PrimCons(symaddress_1_6, tmp7838)

		tmp7840 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7839)

		_ = tmp7840

		tmp7841 := MakeNative(func(__e *ControlFlow) {
			V698 := __e.Get(1)
			_ = V698
			__e.Return(PrimIsVector(V698))
			return
		}, 1)

		tmp7842 := PrimCons(symabsvector_2, tmp7841)

		tmp7843 := Call(__e, PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7842)

		_ = tmp7843

		tmp7844 := MakeNative(func(__e *ControlFlow) {
			V699 := __e.Get(1)
			_ = V699
			__e.Return(PrimAbsvector(V699))
			return
		}, 1)

		tmp7845 := PrimCons(symabsvector, tmp7844)

		__e.TailApply(PrimNS2Value(symshen_4set_1lambda_1form_1entry), tmp7845)
		return

	}, 0)

	tmp7846 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise_1lambda_1forms, tmp7390)

	_ = tmp7846

	tmp7847 := MakeNative(func(__e *ControlFlow) {
		tmp7848 := Call(__e, PrimNS2Value(symshen_4initialise_1environment))

		_ = tmp7848

		tmp7849 := Call(__e, PrimNS2Value(symshen_4initialise_1lambda_1forms))

		_ = tmp7849

		tmp7850 := Call(__e, PrimNS2Value(symshen_4initialise_1signedfunc_1lambda_1forms))

		_ = tmp7850

		__e.TailApply(PrimNS2Value(symshen_4initialise_1signedfuncs))
		return

	}, 0)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4initialise, tmp7847)
	return

}, 0)
